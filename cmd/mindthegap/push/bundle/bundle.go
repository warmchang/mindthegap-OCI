// Copyright 2021 D2iQ, Inc. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package bundle

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
	"sync"

	"github.com/containers/image/v5/types"
	"github.com/google/go-containerregistry/pkg/authn"
	"github.com/google/go-containerregistry/pkg/logs"
	"github.com/google/go-containerregistry/pkg/name"
	v1 "github.com/google/go-containerregistry/pkg/v1"
	"github.com/google/go-containerregistry/pkg/v1/empty"
	"github.com/google/go-containerregistry/pkg/v1/mutate"
	"github.com/google/go-containerregistry/pkg/v1/remote"
	"github.com/google/go-containerregistry/pkg/v1/remote/transport"
	mediatypes "github.com/google/go-containerregistry/pkg/v1/types"
	"github.com/spf13/cobra"
	"github.com/thediveo/enumflag/v2"
	"golang.org/x/sync/errgroup"

	"github.com/mesosphere/dkp-cli-runtime/core/output"

	"github.com/mesosphere/mindthegap/cleanup"
	"github.com/mesosphere/mindthegap/cmd/mindthegap/flags"
	"github.com/mesosphere/mindthegap/cmd/mindthegap/utils"
	"github.com/mesosphere/mindthegap/config"
	"github.com/mesosphere/mindthegap/docker/ecr"
	"github.com/mesosphere/mindthegap/docker/registry"
	"github.com/mesosphere/mindthegap/images"
	"github.com/mesosphere/mindthegap/images/authnhelpers"
	"github.com/mesosphere/mindthegap/images/httputils"
)

type onExistingTagMode enumflag.Flag

const (
	Overwrite onExistingTagMode = iota
	Error
	Skip
	MergeWithRetain
	MergeWithOverwrite
)

var onExistingTagModes = map[onExistingTagMode][]string{
	Overwrite:          {"overwrite"},
	Error:              {"error"},
	Skip:               {"skip"},
	MergeWithRetain:    {"merge-with-retain"},
	MergeWithOverwrite: {"merge-with-overwrite"},
}

func NewCommand(out output.Output, bundleCmdName string) *cobra.Command {
	var (
		bundleFiles                   []string
		destRegistryURI               flags.RegistryURI
		destRegistryCACertificateFile string
		destRegistrySkipTLSVerify     bool
		destRegistryUsername          string
		destRegistryPassword          string
		ecrLifecyclePolicy            string
		onExistingTag                 = Overwrite
		imagePushConcurrency          int
		forceOCIMediaTypes            bool
	)

	cmd := &cobra.Command{
		Use:   bundleCmdName,
		Short: "Push from bundles into an existing OCI registry",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if err := cmd.ValidateRequiredFlags(); err != nil {
				return err
			}

			if err := flags.ValidateFlagsThatRequireValues(cmd, bundleCmdName, "to-registry"); err != nil {
				return err
			}

			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			cleaner := cleanup.NewCleaner()
			defer cleaner.Cleanup()

			out.StartOperation("Creating temporary directory")
			tempDir, err := os.MkdirTemp("", ".bundle-*")
			if err != nil {
				out.EndOperationWithStatus(output.Failure())
				return fmt.Errorf("failed to create temporary directory: %w", err)
			}
			cleaner.AddCleanupFn(func() { _ = os.RemoveAll(tempDir) })
			out.EndOperationWithStatus(output.Success())

			bundleFiles, err = utils.FilesWithGlobs(bundleFiles)
			if err != nil {
				return err
			}
			imagesCfg, chartsCfg, err := utils.ExtractConfigs(tempDir, out, bundleFiles...)
			if err != nil {
				return err
			}

			out.StartOperation("Starting temporary Docker registry")
			storage, err := registry.ArchiveStorage("", bundleFiles...)
			if err != nil {
				out.EndOperationWithStatus(output.Failure())
				return fmt.Errorf("failed to create storage for Docker registry from supplied bundles: %w", err)
			}
			reg, err := registry.NewRegistry(
				registry.Config{Storage: storage},
			)
			if err != nil {
				out.EndOperationWithStatus(output.Failure())
				return fmt.Errorf("failed to create local Docker registry: %w", err)
			}
			go func() {
				if err := reg.ListenAndServe(output.NewOutputLogr(out)); err != nil {
					out.Error(err, "error serving Docker registry")
					os.Exit(2)
				}
			}()
			out.EndOperationWithStatus(output.Success())

			logs.Debug.SetOutput(out.V(4).InfoWriter())
			logs.Warn.SetOutput(out.V(2).InfoWriter())

			sourceTLSRoundTripper, err := httputils.InsecureTLSRoundTripper(remote.DefaultTransport)
			if err != nil {
				out.Error(err, "error configuring TLS for source registry")
				os.Exit(2)
			}
			sourceRemoteOpts := []remote.Option{
				remote.WithTransport(sourceTLSRoundTripper),
				remote.WithUserAgent(utils.Useragent()),
			}

			destTLSRoundTripper, err := httputils.TLSConfiguredRoundTripper(
				remote.DefaultTransport,
				destRegistryURI.Host(),
				flags.SkipTLSVerify(destRegistrySkipTLSVerify, &destRegistryURI),
				destRegistryCACertificateFile,
			)
			if err != nil {
				out.Error(err, "error configuring TLS for destination registry")
				os.Exit(2)
			}
			destRemoteOpts := []remote.Option{
				remote.WithTransport(destTLSRoundTripper),
				remote.WithUserAgent(utils.Useragent()),
			}

			var destNameOpts []name.Option
			if flags.SkipTLSVerify(destRegistrySkipTLSVerify, &destRegistryURI) {
				destNameOpts = append(destNameOpts, name.Insecure)
			}

			// Determine type of destination registry.
			var prePushFuncs []prePushFunc
			if ecr.IsECRRegistry(destRegistryURI.Host()) {
				ecrClient, err := ecr.ClientForRegistry(destRegistryURI.Host())
				if err != nil {
					return err
				}

				prePushFuncs = append(
					prePushFuncs,
					ecr.EnsureRepositoryExistsFunc(ecrClient, ecrLifecyclePolicy),
				)

				// If a password hasn't been specified, then try to retrieve a token.
				if destRegistryPassword == "" {
					out.StartOperation("Retrieving ECR credentials")
					destRegistryUsername, destRegistryPassword, err = ecr.RetrieveUsernameAndToken(
						ecrClient,
					)
					if err != nil {
						out.EndOperationWithStatus(output.Failure())
						return fmt.Errorf(
							"failed to retrieve ECR credentials: %w\n\nPlease ensure you have authenticated to AWS and try again",
							err,
						)
					}
					out.EndOperationWithStatus(output.Success())
				}
			}

			var keychain authn.Keychain = authn.DefaultKeychain
			if destRegistryUsername != "" && destRegistryPassword != "" {
				keychain = authn.NewMultiKeychain(
					authn.NewKeychainFromHelper(
						authnhelpers.NewStaticHelper(
							destRegistryURI.Host(),
							&types.DockerAuthConfig{
								Username: destRegistryUsername,
								Password: destRegistryPassword,
							},
						),
					),
					keychain,
				)
			}
			destRemoteOpts = append(destRemoteOpts, remote.WithAuthFromKeychain(keychain))

			srcRegistry, err := name.NewRegistry(
				reg.Address(),
				name.Insecure,
				name.StrictValidation,
			)
			if err != nil {
				return err
			}
			destRegistry, err := name.NewRegistry(
				destRegistryURI.Host(),
				append(destNameOpts, name.StrictValidation)...)
			if err != nil {
				return err
			}

			if imagesCfg != nil {
				err := pushImages(
					*imagesCfg,
					srcRegistry,
					sourceRemoteOpts,
					destRegistry,
					destRegistryURI.Path(),
					destRemoteOpts,
					onExistingTag,
					imagePushConcurrency,
					out,
					forceOCIMediaTypes,
					prePushFuncs...,
				)
				if err != nil {
					return err
				}
			}

			chartsSrcRegistry, err := name.NewRegistry(
				reg.Address(),
				name.Insecure,
			)
			if err != nil {
				return err
			}

			if chartsCfg != nil {
				err := pushOCIArtifacts(
					*chartsCfg,
					chartsSrcRegistry,
					"/charts",
					sourceRemoteOpts,
					destRegistry,
					destRegistryURI.Path(),
					destRemoteOpts,
					out,
					prePushFuncs...,
				)
				if err != nil {
					return err
				}
			}

			return nil
		},
	}

	cmd.Flags().StringSliceVar(&bundleFiles, bundleCmdName, nil,
		"Tarball containing list of images to push. Can also be a glob pattern.")
	_ = cmd.MarkFlagRequired(bundleCmdName)
	cmd.Flags().Var(&destRegistryURI, "to-registry", "Registry to push images to. "+
		"TLS verification will be skipped when using an http:// registry.")
	_ = cmd.MarkFlagRequired("to-registry")
	cmd.Flags().StringVar(&destRegistryCACertificateFile, "to-registry-ca-cert-file", "",
		"CA certificate file used to verify TLS verification of registry to push images to")
	cmd.Flags().BoolVar(&destRegistrySkipTLSVerify, "to-registry-insecure-skip-tls-verify", false,
		"Skip TLS verification of registry to push images to (also use for non-TLS http registries)")
	cmd.MarkFlagsMutuallyExclusive(
		"to-registry-ca-cert-file",
		"to-registry-insecure-skip-tls-verify",
	)
	cmd.Flags().StringVar(&destRegistryUsername, "to-registry-username", "",
		"Username to use to log in to destination registry")
	cmd.Flags().StringVar(&destRegistryPassword, "to-registry-password", "",
		"Password to use to log in to destination registry")
	cmd.MarkFlagsRequiredTogether(
		"to-registry-username",
		"to-registry-password",
	)
	cmd.Flags().StringVar(&ecrLifecyclePolicy, "ecr-lifecycle-policy-file", "",
		"File containing ECR lifecycle policy for newly created repositories "+
			"(only applies if target registry is hosted on ECR, ignored otherwise)")

	cmd.Flags().Var(
		enumflag.New(&onExistingTag, "string", onExistingTagModes, enumflag.EnumCaseSensitive),
		"on-existing-tag",
		`how to handle existing tags: one of "overwrite", "error", or "skip"`,
	)
	cmd.Flags().
		IntVar(&imagePushConcurrency, "image-push-concurrency", 1, "Image push concurrency")

	cmd.Flags().
		BoolVar(&forceOCIMediaTypes, "force-oci-media-types", false, "force OCI media types")

	return cmd
}

type prePushFunc func(destRepositoryName name.Repository, imageTags ...string) error

type pushConfig struct {
	forceOCIMediaTypes bool
	onExistingTag      onExistingTagMode
}

type pushOpt func(*pushConfig)

func withForceOCIMediaTypes(force bool) pushOpt {
	return func(cfg *pushConfig) {
		cfg.forceOCIMediaTypes = force
	}
}

func withOnExistingTagMode(mode onExistingTagMode) pushOpt {
	return func(cfg *pushConfig) {
		cfg.onExistingTag = mode
	}
}

type pushFunc func(
	srcImage name.Reference,
	sourceRemoteOpts []remote.Option,
	destImage name.Reference,
	destRemoteOpts []remote.Option,
	pushOpts ...pushOpt,
) error

func pushImages(
	cfg config.ImagesConfig,
	sourceRegistry name.Registry, sourceRemoteOpts []remote.Option,
	destRegistry name.Registry, destRegistryPath string, destRemoteOpts []remote.Option,
	onExistingTag onExistingTagMode,
	imagePushConcurrency int,
	out output.Output,
	forceOCIMediaTypes bool,
	prePushFuncs ...prePushFunc,
) error {
	puller, err := remote.NewPuller(destRemoteOpts...)
	if err != nil {
		return err
	}

	// Sort registries for deterministic ordering.
	regNames := cfg.SortedRegistryNames()

	eg, egCtx := errgroup.WithContext(context.Background())
	eg.SetLimit(imagePushConcurrency)

	sourceRemoteOpts = append(sourceRemoteOpts, remote.WithContext(egCtx))
	destRemoteOpts = append(destRemoteOpts, remote.WithContext(egCtx))

	pushGauge := &output.ProgressGauge{}
	pushGauge.SetCapacity(cfg.TotalImages())
	pushGauge.SetStatus("Pushing bundled images")

	out.StartOperationWithProgress(pushGauge)

	for registryIdx := range regNames {
		registryName := regNames[registryIdx]

		registryConfig := cfg[registryName]

		// Sort images for deterministic ordering.
		imageNames := registryConfig.SortedImageNames()

		for imageIdx := range imageNames {
			imageName := imageNames[imageIdx]

			srcRepository := sourceRegistry.Repo(imageName)
			destRepository := destRegistry.Repo(strings.TrimLeft(destRegistryPath, "/"), imageName)

			imageTags := registryConfig.Images[imageName]

			var (
				imageTagPrePushSync sync.Once
				imageTagPrePushErr  error
				existingImageTags   map[string]struct{}
			)

			for tagIdx := range imageTags {
				imageTag := imageTags[tagIdx]

				eg.Go(func() error {
					imageTagPrePushSync.Do(func() {
						for _, prePush := range prePushFuncs {
							if err := prePush(destRepository, imageTags...); err != nil {
								imageTagPrePushErr = fmt.Errorf("pre-push func failed: %w", err)
							}
						}

						existingImageTags, imageTagPrePushErr = getExistingImages(
							context.Background(),
							onExistingTag,
							puller,
							destRepository,
						)
					})

					if imageTagPrePushErr != nil {
						return imageTagPrePushErr
					}

					srcImage := srcRepository.Tag(imageTag)
					destImage := destRepository.Tag(imageTag)

					var pushFn pushFunc = pushTag

					switch onExistingTag {
					case Overwrite, MergeWithRetain, MergeWithOverwrite:
						// Nothing to do here, pushFn is already set to pushTag above.
					case Skip:
						// If tag exists already then do nothing.
						if _, exists := existingImageTags[imageTag]; exists {
							pushFn = func(
								_ name.Reference, _ []remote.Option, _ name.Reference, _ []remote.Option, _ ...pushOpt,
							) error {
								return nil
							}
						}
					case Error:
						if _, exists := existingImageTags[imageTag]; exists {
							return fmt.Errorf(
								"image tag already exists in destination registry",
							)
						}
					}

					opts := []pushOpt{withOnExistingTagMode(onExistingTag)}
					if forceOCIMediaTypes {
						opts = append(opts, withForceOCIMediaTypes(forceOCIMediaTypes))
					}

					if err := pushFn(srcImage, sourceRemoteOpts, destImage, destRemoteOpts, opts...); err != nil {
						return err
					}

					pushGauge.Inc()

					return nil
				})
			}
		}
	}

	if err := eg.Wait(); err != nil {
		out.EndOperationWithStatus(output.Failure())
		return err
	}

	out.EndOperationWithStatus(output.Success())

	return nil
}

func pushTag(
	srcImage name.Reference,
	sourceRemoteOpts []remote.Option,
	destImage name.Reference,
	destRemoteOpts []remote.Option,
	pushOpts ...pushOpt,
) error {
	var pushCfg pushConfig
	for _, opt := range pushOpts {
		opt(&pushCfg)
	}

	desc, err := remote.Get(srcImage, sourceRemoteOpts...)
	if err != nil {
		return err
	}

	if !desc.MediaType.IsIndex() {
		image, err := desc.Image()
		if err != nil {
			return err
		}
		return remote.Write(destImage, image, destRemoteOpts...)
	}

	idx, err := desc.ImageIndex()
	if err != nil {
		return err
	}

	// Get the existing index from the destination registry if merging is enabled.
	if pushCfg.onExistingTag == MergeWithOverwrite || pushCfg.onExistingTag == MergeWithRetain {
		existingIdx, err := fetchExistingIndex(
			destImage,
			destRemoteOpts,
		)
		if err != nil {
			return fmt.Errorf("failed to fetch existing index: %w", err)
		}

		mergeFromIndex, mergeToIndex := idx, existingIdx
		if pushCfg.onExistingTag == MergeWithRetain {
			mergeFromIndex, mergeToIndex = existingIdx, idx
		}

		idx, err = mergeIndexesOverwriteExisting(mergeFromIndex, mergeToIndex)
		if err != nil {
			return fmt.Errorf("failed to merge indexes: %w", err)
		}
	}

	if pushCfg.forceOCIMediaTypes {
		idx, err = convertToOCIIndex(idx, srcImage, sourceRemoteOpts)
		if err != nil {
			return fmt.Errorf("failed to convert index to OCI format: %w", err)
		}
	}

	return remote.WriteIndex(destImage, idx, destRemoteOpts...)
}

func pushOCIArtifacts(
	cfg config.HelmChartsConfig,
	sourceRegistry name.Registry, sourceRegistryPath string, sourceRemoteOpts []remote.Option,
	destRegistry name.Registry, destRegistryPath string, destRemoteOpts []remote.Option,
	out output.Output,
	prePushFuncs ...prePushFunc,
) error {
	// Sort repositories for deterministic ordering.
	repoNames := cfg.SortedRepositoryNames()

	for _, repoName := range repoNames {
		repoConfig := cfg.Repositories[repoName]

		// Sort charts for deterministic ordering.
		chartNames := repoConfig.SortedChartNames()

		for _, chartName := range chartNames {
			srcRepository := sourceRegistry.Repo(
				strings.TrimLeft(sourceRegistryPath, "/"),
				chartName,
			)
			destRepository := destRegistry.Repo(strings.TrimLeft(destRegistryPath, "/"), chartName)

			chartVersions := repoConfig.Charts[chartName]

			for _, prePush := range prePushFuncs {
				if err := prePush(destRepository, chartVersions...); err != nil {
					return fmt.Errorf("pre-push func failed: %w", err)
				}
			}

			for _, chartVersion := range chartVersions {
				destChart := destRepository.Tag(chartVersion)

				out.StartOperation(
					fmt.Sprintf("Copying %s:%s (from bundle) to %s",
						chartName, chartVersion,
						destChart.Name(),
					),
				)

				srcChart := srcRepository.Tag(chartVersion)
				src, err := remote.Image(srcChart, sourceRemoteOpts...)
				if err != nil {
					out.EndOperationWithStatus(output.Failure())
					return err
				}

				if err := remote.Write(destChart, src, destRemoteOpts...); err != nil {
					out.EndOperationWithStatus(output.Failure())
					return err
				}

				out.EndOperationWithStatus(output.Success())
			}
		}
	}

	return nil
}

func getExistingImages(
	ctx context.Context,
	onExistingTag onExistingTagMode,
	puller *remote.Puller,
	repo name.Repository,
) (map[string]struct{}, error) {
	if onExistingTag == Overwrite {
		return nil, nil
	}

	tags, err := puller.List(ctx, repo)
	if err != nil {
		var terr *transport.Error
		if errors.As(err, &terr) {
			// Some registries create repository on first push, so listing tags will fail.
			// If we see 404 or 403, assume we failed because the repository hasn't been created yet.
			if terr.StatusCode == http.StatusNotFound || terr.StatusCode == http.StatusForbidden {
				return nil, nil
			}
		}
		return nil, fmt.Errorf("failed to list existing tags: %w", err)
	}

	existingTags := make(map[string]struct{}, len(tags))
	for _, t := range tags {
		existingTags[t] = struct{}{}
	}

	return existingTags, nil
}

func convertToOCIIndex(
	originalIndex v1.ImageIndex,
	srcImage name.Reference,
	sourceRemoteOpts []remote.Option,
) (v1.ImageIndex, error) {
	originalMediaType, err := originalIndex.MediaType()
	if err != nil {
		return nil, fmt.Errorf("failed to get media type of image index: %w", err)
	}

	if originalMediaType == mediatypes.OCIImageIndex {
		return originalIndex, nil
	}

	var ociIdx v1.ImageIndex = empty.Index
	ociIdx = mutate.IndexMediaType(ociIdx, mediatypes.OCIImageIndex)

	originalIdx, err := originalIndex.IndexManifest()
	if err != nil {
		return nil, fmt.Errorf("failed to read original image index manifest: %w", err)
	}

	for i := range originalIdx.Manifests {
		manifest := originalIdx.Manifests[i]
		manifest.MediaType = mediatypes.OCIManifestSchema1

		digestRef, err := name.NewDigest(
			fmt.Sprintf("%s@%s", srcImage.Context().Name(), manifest.Digest.String()),
		)
		if err != nil {
			return nil, fmt.Errorf("failed to create digest reference: %w", err)
		}

		imgDesc, err := remote.Get(digestRef, sourceRemoteOpts...)
		if err != nil {
			return nil, fmt.Errorf("failed to get image %q: %w", digestRef, err)
		}

		img, err := imgDesc.Image()
		if err != nil {
			return nil, fmt.Errorf(
				"failed to convert image descriptor for %q to image: %w",
				digestRef,
				err,
			)
		}

		ociImg := empty.Image
		ociImg = mutate.MediaType(ociImg, mediatypes.OCIManifestSchema1)
		ociImg = mutate.ConfigMediaType(ociImg, mediatypes.OCIConfigJSON)
		layers, err := img.Layers()
		if err != nil {
			return nil, fmt.Errorf("failed to get layers for image %q: %w", digestRef, err)
		}

		for _, layer := range layers {
			ociImg, err = mutate.Append(ociImg, mutate.Addendum{
				Layer:     layer,
				MediaType: mediatypes.OCILayer,
			})
			if err != nil {
				return nil, fmt.Errorf("failed to append layer to image %q: %w", digestRef, err)
			}
		}

		ociImgDigest, err := ociImg.Digest()
		if err != nil {
			return nil, fmt.Errorf("failed to get digest for image %q: %w", digestRef, err)
		}

		manifest.Digest = ociImgDigest

		ociIdx = mutate.AppendManifests(ociIdx, mutate.IndexAddendum{
			Add:        ociImg,
			Descriptor: manifest,
		})
	}

	return ociIdx, nil
}

func fetchExistingIndex(destImage name.Reference, destRemoteOpts []remote.Option) (v1.ImageIndex, error) {
	existingDesc, err := remote.Get(destImage, destRemoteOpts...)
	if err != nil {
		var terr *transport.Error
		if errors.As(err, &terr) {
			if terr.StatusCode == http.StatusNotFound {
				return empty.Index, nil
			}
		}
		return nil, fmt.Errorf("failed to fetch existing descriptor: %w", err)
	}

	switch {
	case existingDesc.MediaType.IsIndex():
		index, err := existingDesc.ImageIndex()
		if err != nil {
			return nil, fmt.Errorf("failed to read image index for %q: %w", destImage, err)
		}
		return index, nil
	case existingDesc.MediaType.IsImage():
		image, err := existingDesc.Image()
		if err != nil {
			return nil, fmt.Errorf("failed to read image for %q: %w", destImage, err)
		}
		return images.IndexForSinglePlatformImage(destImage, image)
	default:
		return nil, fmt.Errorf(
			"unexpected media type in descriptor for image %q: %v",
			destImage,
			existingDesc.MediaType,
		)
	}
}

func mergeIndexesOverwriteExisting(
	mergeFromIndex v1.ImageIndex,
	mergeToIndex v1.ImageIndex,
) (v1.ImageIndex, error) {
	mergeFromIndexManifest, err := mergeFromIndex.IndexManifest()
	if err != nil {
		return nil, fmt.Errorf("failed to read index manifest: %w", err)
	}

	// Collect all platforms from the mergeFromIndex.
	var fromPlatforms []v1.Platform
	for manifestIdx := range mergeFromIndexManifest.Manifests {
		child := mergeFromIndexManifest.Manifests[manifestIdx]
		if child.Platform != nil {
			fromPlatforms = append(fromPlatforms, *child.Platform)
		}
	}

	// Filter out all the platforms in mergeToIndex that were already in mergeFromIndex.
	mergeToIndex = mutate.RemoveManifests(mergeToIndex, func(manifest v1.Descriptor) bool {
		if manifest.Platform == nil {
			return false
		}

		for _, fromPlatform := range fromPlatforms {
			if manifest.Platform.Satisfies(fromPlatform) {
				return true
			}
		}

		return false
	})

	var adds []mutate.IndexAddendum
	for manifestIdx := range mergeFromIndexManifest.Manifests {
		child := mergeFromIndexManifest.Manifests[manifestIdx]
		switch {
		case child.MediaType.IsImage():
			img, err := mergeFromIndex.Image(child.Digest)
			if err != nil {
				return nil, err
			}
			adds = append(adds, mutate.IndexAddendum{
				Add:        img,
				Descriptor: child,
			})
		case child.MediaType.IsIndex():
			idx, err := mergeFromIndex.ImageIndex(child.Digest)
			if err != nil {
				return nil, err
			}
			adds = append(adds, mutate.IndexAddendum{
				Add:        idx,
				Descriptor: child,
			})
		default:
			return nil, fmt.Errorf("unexpected child %q with media type %q", child.Digest, child.MediaType)
		}
	}

	mergedIndex := mutate.AppendManifests(mergeToIndex, adds...)

	return mergedIndex, nil
}
