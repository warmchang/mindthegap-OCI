// Copyright 2021 D2iQ, Inc. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package helmbundle

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/mesosphere/dkp-cli-runtime/core/output"

	"github.com/mesosphere/mindthegap/cleanup"
	"github.com/mesosphere/mindthegap/cmd/mindthegap/utils"
	"github.com/mesosphere/mindthegap/config"
	"github.com/mesosphere/mindthegap/docker/ecr"
	"github.com/mesosphere/mindthegap/docker/registry"
	"github.com/mesosphere/mindthegap/skopeo"
)

func NewCommand(out output.Output) *cobra.Command {
	var (
		helmBundleFiles           []string
		destRegistry              string
		destRegistrySkipTLSVerify bool
		destRegistryUsername      string
		destRegistryPassword      string
	)

	cmd := &cobra.Command{
		Use:   "helm-bundle",
		Short: "Push images from a Helm chart bundle into an existing OCI registry",
		RunE: func(cmd *cobra.Command, args []string) error {
			cleaner := cleanup.NewCleaner()
			defer cleaner.Cleanup()

			out.StartOperation("Creating temporary directory")
			tempDir, err := os.MkdirTemp("", ".helm-bundle-*")
			if err != nil {
				out.EndOperation(false)
				return fmt.Errorf("failed to create temporary directory: %w", err)
			}
			cleaner.AddCleanupFn(func() { _ = os.RemoveAll(tempDir) })
			out.EndOperation(true)

			helmBundleFiles, err = utils.FilesWithGlobs(helmBundleFiles)
			if err != nil {
				return err
			}
			_, cfg, err := utils.ExtractBundles(tempDir, out, helmBundleFiles...)
			if err != nil {
				return err
			}

			out.StartOperation("Starting temporary Docker registry")
			reg, err := registry.NewRegistry(
				registry.Config{StorageDirectory: tempDir, ReadOnly: true},
			)
			if err != nil {
				out.EndOperation(false)
				return fmt.Errorf("failed to create local Docker registry: %w", err)
			}
			go func() {
				if err := reg.ListenAndServe(); err != nil {
					out.Error(err, "error serving Docker registry")
					os.Exit(2)
				}
			}()
			out.EndOperation(true)

			skopeoRunner, skopeoCleanup := skopeo.NewRunner()
			cleaner.AddCleanupFn(func() { _ = skopeoCleanup() })

			skopeoOpts := []skopeo.SkopeoOption{
				skopeo.PreserveDigests(),
			}
			if destRegistryUsername != "" && destRegistryPassword != "" {
				skopeoOpts = append(
					skopeoOpts,
					skopeo.DestCredentials(
						destRegistryUsername,
						destRegistryPassword,
					),
				)
			} else {
				skopeoStdout, skopeoStderr, err := skopeoRunner.AttemptToLoginToRegistry(
					context.Background(),
					destRegistry,
				)
				if err != nil {
					out.Infof("---skopeo stdout---:\n%s", skopeoStdout)
					out.Infof("---skopeo stderr---:\n%s", skopeoStderr)
					return fmt.Errorf("error logging in to target registry: %w", err)
				}
				out.V(4).Infof("---skopeo stdout---:\n%s", skopeoStdout)
				out.V(4).Infof("---skopeo stderr---:\n%s", skopeoStderr)
			}

			// Determine type of destination registry.
			var prePushFuncs []prePushFunc
			if ecr.IsECRRegistry(destRegistry) {
				prePushFuncs = append(
					prePushFuncs,
					ecr.EnsureRepositoryExistsFunc(""),
				)
			}

			return pushOCIArtifacts(
				cfg,
				fmt.Sprintf("%s/charts", reg.Address()),
				destRegistry,
				skopeoOpts,
				destRegistrySkipTLSVerify,
				out,
				skopeoRunner,
				prePushFuncs...,
			)
		},
	}

	cmd.Flags().StringSliceVar(&helmBundleFiles, "helm-bundle", nil,
		"Tarball containing list of Helm charts to push. Can also be a glob pattern.")
	_ = cmd.MarkFlagRequired("helm-bundle")
	cmd.Flags().StringVar(&destRegistry, "to-registry", "", "Registry to push images to")
	_ = cmd.MarkFlagRequired("to-registry")
	cmd.Flags().BoolVar(&destRegistrySkipTLSVerify, "to-registry-insecure-skip-tls-verify", false,
		"Skip TLS verification of registry to push images to (use for http registries)")
	cmd.Flags().StringVar(&destRegistryUsername, "to-registry-username", "",
		"Username to use to log in to destination repository")
	cmd.Flags().StringVar(&destRegistryPassword, "to-registry-password", "",
		"Password to use to log in to destination registry")

	// TODO Unhide this from DKP CLI once DKP supports OCI registry for Helm charts.
	utils.AddCmdAnnotation(cmd, "exclude-from-dkp-cli", "true")

	return cmd
}

type prePushFunc func(destRegistry, imageName string, imageTags ...string) error

func pushOCIArtifacts(
	cfg config.HelmChartsConfig,
	sourceRegistry, destRegistry string,
	skopeoOpts []skopeo.SkopeoOption,
	destRegistrySkipTLSVerify bool,
	out output.Output,
	skopeoRunner *skopeo.Runner,
	prePushFuncs ...prePushFunc,
) error {
	skopeoOpts = append(skopeoOpts, skopeo.DisableSrcTLSVerify())
	if destRegistrySkipTLSVerify {
		skopeoOpts = append(skopeoOpts, skopeo.DisableDestTLSVerify())
	}

	// Sort repositories for deterministic ordering.
	repoNames := cfg.SortedRepositoryNames()

	for _, repoName := range repoNames {
		repoConfig := cfg.Repositories[repoName]

		// Sort charts for deterministic ordering.
		chartNames := repoConfig.SortedChartNames()

		for _, chartName := range chartNames {
			chartVersions := repoConfig.Charts[chartName]

			for _, prePush := range prePushFuncs {
				if err := prePush("", destRegistry); err != nil {
					return fmt.Errorf("pre-push func failed: %w", err)
				}
			}

			for _, chartVersion := range chartVersions {
				out.StartOperation(
					fmt.Sprintf("Copying %s:%s (from bundle) to %s/%s:%s",
						chartName, chartVersion,
						destRegistry, chartName, chartVersion,
					),
				)
				skopeoStdout, skopeoStderr, err := skopeoRunner.Copy(context.TODO(),
					fmt.Sprintf("docker://%s/%s:%s", sourceRegistry, chartName, chartVersion),
					fmt.Sprintf("docker://%s/%s:%s", destRegistry, chartName, chartVersion),
					append(
						skopeoOpts, skopeo.All(),
					)...,
				)
				if err != nil {
					out.EndOperation(false)
					out.Infof("---skopeo stdout---:\n%s", skopeoStdout)
					out.Infof("---skopeo stderr---:\n%s", skopeoStderr)
					return err
				}
				out.V(4).Infof("---skopeo stdout---:\n%s", skopeoStdout)
				out.V(4).Infof("---skopeo stderr---:\n%s", skopeoStderr)
				out.EndOperation(true)
			}
		}
	}

	return nil
}