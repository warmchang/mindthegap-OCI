# Changelog

<!-- Release notes generated using configuration in .github/release.yaml at main -->

## What's Changed
### Exciting New Features 🎉
* feat: Serve multiple bundles by @jimmidyson in https://github.com/mesosphere/mindthegap/pull/165
* feat: Import multiple image bundles by @jimmidyson in https://github.com/mesosphere/mindthegap/pull/167
* feat: Push multiple image bundles by @jimmidyson in https://github.com/mesosphere/mindthegap/pull/169
### Fixes 🔧
* fix: Write out merged image bundle config when serving by @jimmidyson in https://github.com/mesosphere/mindthegap/pull/168
### Other Changes
* build(deps): bump helm.sh/helm/v3 from 3.9.2 to 3.9.3 by @dependabot in https://github.com/mesosphere/mindthegap/pull/163
* build(deps): bump github.com/aws/aws-sdk-go-v2/service/ecr from 1.17.11 to 1.17.12 by @dependabot in https://github.com/mesosphere/mindthegap/pull/162
* build(deps): bump github.com/aws/aws-sdk-go-v2/config from 1.15.17 to 1.17.1 by @dependabot in https://github.com/mesosphere/mindthegap/pull/164


**Full Changelog**: https://github.com/mesosphere/mindthegap/compare/v0.18.0...v0.19.0

<!-- Release notes generated using configuration in .github/release.yaml at main -->

## What's Changed
### Breaking Changes 🛠
* fix!: Remove unnecessary gzip compression of bundles by default by @jimmidyson in https://github.com/mesosphere/mindthegap/pull/174
### Exciting New Features 🎉
* feat: Use OCI storage for helm chart bundle by @jimmidyson in https://github.com/mesosphere/mindthegap/pull/175
* feat: support globs for image and helm bundles by @dkoshkin in https://github.com/mesosphere/mindthegap/pull/183
### Fixes 🔧
* fix: Ensure that image IDs do not change on import by @jimmidyson in https://github.com/mesosphere/mindthegap/pull/185
### Other Changes
* refactor: Move mindthegap cmd to own package by @jimmidyson in https://github.com/mesosphere/mindthegap/pull/170
* build(deps): bump helm.sh/helm/v3 from 3.9.3 to 3.9.4 by @dependabot in https://github.com/mesosphere/mindthegap/pull/173
* build: Upgrade all tools by @jimmidyson in https://github.com/mesosphere/mindthegap/pull/176
* build(deps): bump github.com/aws/aws-sdk-go-v2/service/ecr from 1.17.12 to 1.17.16 by @dependabot in https://github.com/mesosphere/mindthegap/pull/178
* build(deps): bump github.com/docker/cli from 20.10.17+incompatible to 20.10.18+incompatible by @dependabot in https://github.com/mesosphere/mindthegap/pull/181
* build(deps): bump k8s.io/klog/v2 from 2.70.1 to 2.80.1 by @dependabot in https://github.com/mesosphere/mindthegap/pull/182
* build(deps): bump github.com/aws/aws-sdk-go-v2/config from 1.17.1 to 1.17.5 by @dependabot in https://github.com/mesosphere/mindthegap/pull/179
* build(deps): bump k8s.io/apimachinery from 0.24.2 to 0.25.0 by @dependabot in https://github.com/mesosphere/mindthegap/pull/180
* test: Add e2e tests for helm bundle functionality by @jimmidyson in https://github.com/mesosphere/mindthegap/pull/184
* test: Add e2e tests for image bundle serve and push by @jimmidyson in https://github.com/mesosphere/mindthegap/pull/186


**Full Changelog**: https://github.com/mesosphere/mindthegap/compare/v0.19.0...v1.0.0-rc.1

<!-- Release notes generated using configuration in .github/release.yaml at main -->

## What's Changed
### Exciting New Features 🎉
* feat: Hide helm bundle commands from DKP CLI by @jimmidyson in https://github.com/mesosphere/mindthegap/pull/187
### Other Changes
* docs: Add mindthegap image for fun by @jimmidyson in https://github.com/mesosphere/mindthegap/pull/189


**Full Changelog**: https://github.com/mesosphere/mindthegap/compare/v1.0.0-rc.1...v1.0.0

<!-- Release notes generated using configuration in .github/release.yaml at main -->

## What's Changed
### Exciting New Features 🎉
* feat: Add Docker image by @jimmidyson in https://github.com/mesosphere/mindthegap/pull/198
### Fixes 🔧
* fix: Ensure file permissions are preserved when copying files by @jimmidyson in https://github.com/mesosphere/mindthegap/pull/191
* test: Add e2e test for import image-bundle  by @jimmidyson in https://github.com/mesosphere/mindthegap/pull/192
* fix: Properly handle friendly image names by @jimmidyson in https://github.com/mesosphere/mindthegap/pull/195
### Other Changes
* build(deps): Upgrade direct dependencies by @jimmidyson in https://github.com/mesosphere/mindthegap/pull/193
* build: Upgrade tools by @jimmidyson in https://github.com/mesosphere/mindthegap/pull/194
* build(deps): bump k8s.io/apimachinery from 0.25.0 to 0.25.1 by @dependabot in https://github.com/mesosphere/mindthegap/pull/196
* build(deps): bump github.com/onsi/ginkgo/v2 from 2.1.6 to 2.2.0 by @dependabot in https://github.com/mesosphere/mindthegap/pull/197


**Full Changelog**: https://github.com/mesosphere/mindthegap/compare/v1.0.0...v1.1.0

<!-- Release notes generated using configuration in .github/release.yaml at main -->

## What's Changed
### Fixes 🔧
* fix: Correctly identify ECR registries by @jimmidyson in https://github.com/mesosphere/mindthegap/pull/206
### Other Changes
* build: Upgrade go-tuf to fix security issue by @jimmidyson in https://github.com/mesosphere/mindthegap/pull/199
* test: Update sha for import bundle check by @jimmidyson in https://github.com/mesosphere/mindthegap/pull/207
* build(deps): bump github.com/aws/aws-sdk-go-v2/config from 1.17.6 to 1.17.7 by @dependabot in https://github.com/mesosphere/mindthegap/pull/201
* build(deps): bump k8s.io/apimachinery from 0.25.1 to 0.25.2 by @dependabot in https://github.com/mesosphere/mindthegap/pull/203
* build(deps): bump github.com/mesosphere/dkp-cli-runtime/core from 0.6.0 to 0.7.0 by @dependabot in https://github.com/mesosphere/mindthegap/pull/204
* build(deps): bump github.com/aws/aws-sdk-go-v2/service/ecr from 1.17.17 to 1.17.18 by @dependabot in https://github.com/mesosphere/mindthegap/pull/202
* build(deps): bump helm.sh/helm/v3 from 3.9.4 to 3.10.0 by @dependabot in https://github.com/mesosphere/mindthegap/pull/205
* ci: Login to Docker Hub before running make release by @jimmidyson in https://github.com/mesosphere/mindthegap/pull/208
* ci: Fix up Docker login by @jimmidyson in https://github.com/mesosphere/mindthegap/pull/209


**Full Changelog**: https://github.com/mesosphere/mindthegap/compare/v1.1.0...v1.1.1

<!-- Release notes generated using configuration in .github/release.yaml at main -->

## What's Changed
### Other Changes
* build(deps): bump github.com/containers/image/v5 from 5.22.0 to 5.23.0 by @dependabot in https://github.com/mesosphere/mindthegap/pull/212
* build(deps): bump github.com/aws/aws-sdk-go-v2/config from 1.17.7 to 1.17.8 by @dependabot in https://github.com/mesosphere/mindthegap/pull/211
* build(deps): bump github.com/containers/skopeo from 1.9.2 to 1.10.0 in /skopeo-static by @dependabot in https://github.com/mesosphere/mindthegap/pull/210
* ci: Specify release-please action version by @jimmidyson in https://github.com/mesosphere/mindthegap/pull/215


**Full Changelog**: https://github.com/mesosphere/mindthegap/compare/v1.1.1...v1.2.0

<!-- Release notes generated using configuration in .github/release.yaml at main -->

## What's Changed
### Other Changes
* build(deps): bump github.com/onsi/gomega from 1.20.2 to 1.21.1 by @dependabot in https://github.com/mesosphere/mindthegap/pull/216
* build(deps): bump github.com/opencontainers/image-spec from 1.1.0-rc1 to 1.1.0-rc2 by @dependabot in https://github.com/mesosphere/mindthegap/pull/217
* test: Fix changed image SHA for import e2e test by @jimmidyson in https://github.com/mesosphere/mindthegap/pull/224
* build: Upgrade tools by @jimmidyson in https://github.com/mesosphere/mindthegap/pull/225
* build(deps): bump k8s.io/apimachinery from 0.25.2 to 0.25.3 by @dependabot in https://github.com/mesosphere/mindthegap/pull/222
* build(deps): bump github.com/docker/docker from 20.10.18+incompatible to 20.10.19+incompatible by @dependabot in https://github.com/mesosphere/mindthegap/pull/220
* build(deps): bump github.com/onsi/ginkgo/v2 from 2.2.0 to 2.3.1 by @dependabot in https://github.com/mesosphere/mindthegap/pull/221
* build(deps): bump helm.sh/helm/v3 from 3.10.0 to 3.10.1 by @dependabot in https://github.com/mesosphere/mindthegap/pull/219
* build(deps): bump github.com/docker/cli from 20.10.18+incompatible to 20.10.19+incompatible by @dependabot in https://github.com/mesosphere/mindthegap/pull/223
* build: Latest distroless image by @jimmidyson in https://github.com/mesosphere/mindthegap/pull/226


**Full Changelog**: https://github.com/mesosphere/mindthegap/compare/v1.2.0...v1.2.1

<!-- Release notes generated using configuration in .github/release.yaml at main -->

## What's Changed
### Fixes 🔧
* fix: Upgrade github.com/sylabs/sif/v2@v2.8.1 by @jimmidyson in https://github.com/mesosphere/mindthegap/pull/227
### Other Changes
* build: Reinstate upx packing of linux binaries by @jimmidyson in https://github.com/mesosphere/mindthegap/pull/228


**Full Changelog**: https://github.com/mesosphere/mindthegap/compare/v1.2.1...v1.2.2

<!-- Release notes generated using configuration in .github/release.yaml at main -->



**Full Changelog**: https://github.com/mesosphere/mindthegap/compare/v1.2.2...v1.2.3

<!-- Release notes generated using configuration in .github/release.yaml at main -->

## What's Changed
### Fixes 🔧
* fix: Ensure docker daemon skopeo logging is output on error by @jimmidyson in https://github.com/mesosphere/mindthegap/pull/231


**Full Changelog**: https://github.com/mesosphere/mindthegap/compare/v1.2.3...v1.2.4

<!-- Release notes generated using configuration in .github/release.yaml at main -->

## What's Changed
### Fixes 🔧
* fix: Specify exe suffix for skopeo on Windows by @jimmidyson in https://github.com/mesosphere/mindthegap/pull/233


**Full Changelog**: https://github.com/mesosphere/mindthegap/compare/v1.2.4...v1.3.0

<!-- Release notes generated using configuration in .github/release.yaml at main -->

## What's Changed
### Fixes 🔧
* fix: Clarify help text for CA certs by @jimmidyson in https://github.com/mesosphere/mindthegap/pull/236


**Full Changelog**: https://github.com/mesosphere/mindthegap/compare/v1.3.0...v1.3.1

<!-- Release notes generated using configuration in .github/release.yaml at main -->

## What's Changed
### Exciting New Features 🎉
* feat: support passing optional scheme in --to-registry by @dkoshkin in https://github.com/mesosphere/mindthegap/pull/252
* feat: Reenable upx for macos binaries by @jimmidyson in https://github.com/mesosphere/mindthegap/pull/262
### Fixes 🔧
* fix: Import all platforms from image bundle by @jimmidyson in https://github.com/mesosphere/mindthegap/pull/292
### Other Changes
* build: Upgrade tools by @jimmidyson in https://github.com/mesosphere/mindthegap/pull/244
* ci: Refactor checks to remove more asdf stuff by @jimmidyson in https://github.com/mesosphere/mindthegap/pull/253
* ci: Add gha actions to dependabot config by @jimmidyson in https://github.com/mesosphere/mindthegap/pull/263
* build: Upgrade tools to latest versions by @jimmidyson in https://github.com/mesosphere/mindthegap/pull/268
* ci: Add dependabot automation to auto-approve and enable auto-merge by @jimmidyson in https://github.com/mesosphere/mindthegap/pull/269
* build: Upgrade tools by @jimmidyson in https://github.com/mesosphere/mindthegap/pull/286
* build: Upgrade goreleaser and upx by @jimmidyson in https://github.com/mesosphere/mindthegap/pull/291
* build: Do not show dependency updates in release notes by @jimmidyson in https://github.com/mesosphere/mindthegap/pull/294
* build: Ensure that license and readme are contained in released archive by @jimmidyson in https://github.com/mesosphere/mindthegap/pull/295
* build: Update deprecated ginkgo flag: progress -> show-node-events by @jimmidyson in https://github.com/mesosphere/mindthegap/pull/296
* build: Ensure that e2e build tags are enabled in vscode by @jimmidyson in https://github.com/mesosphere/mindthegap/pull/297
* test: Add e2e for arm64 import by @jimmidyson in https://github.com/mesosphere/mindthegap/pull/293


**Full Changelog**: https://github.com/mesosphere/mindthegap/compare/v1.3.1...v1.4.0

<!-- Release notes generated using configuration in .github/release.yaml at main -->

## What's Changed
### Exciting New Features 🎉
* feat: Migrate to crane from skopeo by @jimmidyson in https://github.com/mesosphere/mindthegap/pull/308
### Other Changes
* build: Increase timeout for goreleaser to work on public GHA runners by @jimmidyson in https://github.com/mesosphere/mindthegap/pull/299
* build: Remove broken upx test after compression by @jimmidyson in https://github.com/mesosphere/mindthegap/pull/300
* build: Use lzma compression for better file size, and decrease upx level by @jimmidyson in https://github.com/mesosphere/mindthegap/pull/302
* build: Use crane instead of skopeo for updating distroless image by @jimmidyson in https://github.com/mesosphere/mindthegap/pull/307
* ci: Ensure github api token is set for asdf by @jimmidyson in https://github.com/mesosphere/mindthegap/pull/306


**Full Changelog**: https://github.com/mesosphere/mindthegap/compare/v1.4.0...v1.5.0

<!-- Release notes generated using configuration in .github/release.yaml at main -->

## What's Changed
### Fixes 🔧
* fix: Do not upx pack darwin/arm64 binaries by @jimmidyson in https://github.com/mesosphere/mindthegap/pull/311


**Full Changelog**: https://github.com/mesosphere/mindthegap/compare/v1.5.0...v1.5.1

<!-- Release notes generated using configuration in .github/release.yaml at main -->

## What's Changed
### Fixes 🔧
* fix: Support docker certs.d host certs again by @jimmidyson in https://github.com/mesosphere/mindthegap/pull/313


**Full Changelog**: https://github.com/mesosphere/mindthegap/compare/v1.5.1...v1.5.2

<!-- Release notes generated using configuration in .github/release.yaml at main -->

## What's Changed
### Fixes 🔧
* fix: Always use custom transport to ensure TLS config is correct by @jimmidyson in https://github.com/mesosphere/mindthegap/pull/315


**Full Changelog**: https://github.com/mesosphere/mindthegap/compare/v1.5.2...v1.5.3

<!-- Release notes generated using configuration in .github/release.yaml at main -->

## What's Changed
### Fixes 🔧
* fix: Do not upx pack darwin binaries by @jimmidyson in https://github.com/mesosphere/mindthegap/pull/342
### Other Changes
* ci: Remove redundant go version from golangci lint step by @jimmidyson in https://github.com/mesosphere/mindthegap/pull/336


**Full Changelog**: https://github.com/mesosphere/mindthegap/compare/v1.5.3...v1.5.4

<!-- Release notes generated using configuration in .github/release.yaml at main -->

## What's Changed
### Exciting New Features 🎉
* feat: Use ghcr.io rather than Docker Hub by @jimmidyson in https://github.com/mesosphere/mindthegap/pull/343


**Full Changelog**: https://github.com/mesosphere/mindthegap/compare/v1.5.4...v1.6.0

<!-- Release notes generated using configuration in .github/release.yaml at main -->

## What's Changed
### Fixes 🔧
* fix: remove hardcoded AWS region when creating ECR lifecycle policies by @supershal in https://github.com/mesosphere/mindthegap/pull/353

## New Contributors
* @supershal made their first contribution in https://github.com/mesosphere/mindthegap/pull/353

**Full Changelog**: https://github.com/mesosphere/mindthegap/compare/v1.6.0...v1.6.1

<!-- Release notes generated using configuration in .github/release.yaml at main -->

## What's Changed
### Exciting New Features 🎉
* feat: Introduce common serve and push bundle commands by @jimmidyson in https://github.com/mesosphere/mindthegap/pull/359
### Other Changes
* build: Upgrade all tools, dependencies, and fix build scripts by @jimmidyson in https://github.com/mesosphere/mindthegap/pull/356
* ci: Separate out release tag builds from release please workflow by @jimmidyson in https://github.com/mesosphere/mindthegap/pull/358


**Full Changelog**: https://github.com/mesosphere/mindthegap/compare/v1.6.1...v1.7.0

## 1.12.0 (2023-12-13)


### ⚠ BREAKING CHANGES

* Remove unnecessary gzip compression of bundles by default ([#174](https://github.com/warmchang/mindthegap-OCI/issues/174))

### Features

* Add --on-existing-tag flag to push bundle ([#449](https://github.com/warmchang/mindthegap-OCI/issues/449)) ([c655d74](https://github.com/warmchang/mindthegap-OCI/commit/c655d749fe0560ccd0c62acdd42c825176c01a27))
* Add Docker image ([#198](https://github.com/warmchang/mindthegap-OCI/issues/198)) ([41343a4](https://github.com/warmchang/mindthegap-OCI/commit/41343a40d627ed157c3f3adfdc142a7bb35fca91))
* Add mindthegap to useragent ([#476](https://github.com/warmchang/mindthegap-OCI/issues/476)) ([14f7d2c](https://github.com/warmchang/mindthegap-OCI/commit/14f7d2ce1819132d0bf369c7296b601b49a392dc))
* Auto-create ECR registries if they do not exist ([#79](https://github.com/warmchang/mindthegap-OCI/issues/79)) ([b595607](https://github.com/warmchang/mindthegap-OCI/commit/b595607e523497480998a4d35522c8d3553635ba))
* Automatic retrieval of ECR credentials if using ECR as registry ([#440](https://github.com/warmchang/mindthegap-OCI/issues/440)) ([19de7b0](https://github.com/warmchang/mindthegap-OCI/commit/19de7b0ff47044c741abb0e1e299c687f90e97e7))
* Build skopeo with go v1.18 ([9d04798](https://github.com/warmchang/mindthegap-OCI/commit/9d0479845fa34337ab56ae4a604b0147a6e061bb))
* go v1.18 ([#110](https://github.com/warmchang/mindthegap-OCI/issues/110)) ([f84ffa3](https://github.com/warmchang/mindthegap-OCI/commit/f84ffa3019590f441fdfc935f112e3c7f049b6fc))
* Hide helm bundle commands from DKP CLI ([#187](https://github.com/warmchang/mindthegap-OCI/issues/187)) ([0bfc696](https://github.com/warmchang/mindthegap-OCI/commit/0bfc6966e407d0e2cbb761cd139d5b6c9568a8ad))
* Image push concurrency ([#477](https://github.com/warmchang/mindthegap-OCI/issues/477)) ([dfa7be6](https://github.com/warmchang/mindthegap-OCI/commit/dfa7be6cc19f71657f871c89ee84b3be127b8047))
* Import multiple image bundles ([#167](https://github.com/warmchang/mindthegap-OCI/issues/167)) ([f94e4fe](https://github.com/warmchang/mindthegap-OCI/commit/f94e4fe35452e538ef6e2a165fc71f5e052bdb54))
* Introduce common serve and push bundle commands ([#359](https://github.com/warmchang/mindthegap-OCI/issues/359)) ([84ad6e3](https://github.com/warmchang/mindthegap-OCI/commit/84ad6e32b731c6752cbe3772b272e1f66ee45d33))
* Migrate to crane from skopeo ([#308](https://github.com/warmchang/mindthegap-OCI/issues/308)) ([a088f1e](https://github.com/warmchang/mindthegap-OCI/commit/a088f1e77796404c81b1359a6c14808355cf1495))
* Parallel image pulls ([#456](https://github.com/warmchang/mindthegap-OCI/issues/456)) ([4341db2](https://github.com/warmchang/mindthegap-OCI/commit/4341db2910d6b905fb16a31f3f2569ae9ab64431))
* Push multiple image bundles ([#169](https://github.com/warmchang/mindthegap-OCI/issues/169)) ([3c418fa](https://github.com/warmchang/mindthegap-OCI/commit/3c418fae7a72e764c090afdccbb011bdf68d19c7))
* Reenable upx for macos binaries ([#262](https://github.com/warmchang/mindthegap-OCI/issues/262)) ([f70eaf1](https://github.com/warmchang/mindthegap-OCI/commit/f70eaf19c6df973e786a5f23a5bd4c64206b7507))
* Serve multiple bundles ([#165](https://github.com/warmchang/mindthegap-OCI/issues/165)) ([f9fc263](https://github.com/warmchang/mindthegap-OCI/commit/f9fc263d8c8aa7ebede85a13b597c9c6f23c11f3))
* skopeo v1.7.0 ([9440c2b](https://github.com/warmchang/mindthegap-OCI/commit/9440c2b82ab10c6526fe34abcaec83deac0150f9))
* Skopeo v1.8.0 ([c8bdd81](https://github.com/warmchang/mindthegap-OCI/commit/c8bdd8135908de865eb8c07fdb7bd3320978ef0b))
* support globs for image and helm bundles ([#183](https://github.com/warmchang/mindthegap-OCI/issues/183)) ([2f7bafa](https://github.com/warmchang/mindthegap-OCI/commit/2f7bafa4e88a86f3a753a96cf661f9e14f744d2f))
* support passing optional scheme in --to-registry ([#252](https://github.com/warmchang/mindthegap-OCI/issues/252)) ([6983970](https://github.com/warmchang/mindthegap-OCI/commit/69839702f7b4935218db5e146d992e813e5254d2))
* Use ghcr.io rather than Docker Hub ([#343](https://github.com/warmchang/mindthegap-OCI/issues/343)) ([b3c7655](https://github.com/warmchang/mindthegap-OCI/commit/b3c7655f3e2248e02e25562a9a10b0cc1133ffaf))
* Use OCI storage for helm chart bundle ([#175](https://github.com/warmchang/mindthegap-OCI/issues/175)) ([dd3be37](https://github.com/warmchang/mindthegap-OCI/commit/dd3be371fb4e5528bd5e05c50ca869f7c622f57e))
* Validate empty values for required flags ([#541](https://github.com/warmchang/mindthegap-OCI/issues/541)) ([d1d8eee](https://github.com/warmchang/mindthegap-OCI/commit/d1d8eee987890c862fbab5349a2de3749f59b809))


### Bug Fixes

* Always use custom transport to ensure TLS config is correct ([#315](https://github.com/warmchang/mindthegap-OCI/issues/315)) ([962c092](https://github.com/warmchang/mindthegap-OCI/commit/962c0925532b1fa4241ac1435c9bd32cd6e30181))
* Clarify help text for CA certs ([#236](https://github.com/warmchang/mindthegap-OCI/issues/236)) ([8faea4f](https://github.com/warmchang/mindthegap-OCI/commit/8faea4fe67084c9d2e2f24d9ba5501ae598478ba))
* commitDate in goreleaser ([#92](https://github.com/warmchang/mindthegap-OCI/issues/92)) ([017d2f3](https://github.com/warmchang/mindthegap-OCI/commit/017d2f36da2eb61b556f034e3e3232bd84ce7796))
* Correctly handle multiple registries ([#459](https://github.com/warmchang/mindthegap-OCI/issues/459)) ([6113230](https://github.com/warmchang/mindthegap-OCI/commit/6113230bacc70b8dc0a388250ced0958a6d0df5d))
* Correctly identify ECR registries ([#206](https://github.com/warmchang/mindthegap-OCI/issues/206)) ([944e546](https://github.com/warmchang/mindthegap-OCI/commit/944e546270fba4dcad1c5b5f6cf20c08d714b97a))
* Disable upx for all platforms ([#129](https://github.com/warmchang/mindthegap-OCI/issues/129)) ([f5496c3](https://github.com/warmchang/mindthegap-OCI/commit/f5496c373f195e6e639d4d5911b68ad37c81ccc6))
* Do not log base64 encoded ECR token ([#450](https://github.com/warmchang/mindthegap-OCI/issues/450)) ([cc1d5b4](https://github.com/warmchang/mindthegap-OCI/commit/cc1d5b430ac7d06c01a0748f3267a13844f249ad))
* Do not upx pack darwin binaries ([#342](https://github.com/warmchang/mindthegap-OCI/issues/342)) ([e9d172d](https://github.com/warmchang/mindthegap-OCI/commit/e9d172d4a7ba3af37be5ef8234cdba11fe1bfb8c))
* Do not upx pack darwin/arm64 binaries ([#311](https://github.com/warmchang/mindthegap-OCI/issues/311)) ([9f2de10](https://github.com/warmchang/mindthegap-OCI/commit/9f2de10152e32aa0d252dbc2203178eb9cb8da7c))
* Ensure docker daemon skopeo logging is output on error ([#231](https://github.com/warmchang/mindthegap-OCI/issues/231)) ([5e061e4](https://github.com/warmchang/mindthegap-OCI/commit/5e061e417730f4a4539a6ba4144b3d42a65ccd91))
* Ensure file permissions are preserved when copying files ([#191](https://github.com/warmchang/mindthegap-OCI/issues/191)) ([25e68b4](https://github.com/warmchang/mindthegap-OCI/commit/25e68b4a70e86c0e8b7391ddf3e36c2095d46b44))
* Ensure that image IDs do not change on import ([#185](https://github.com/warmchang/mindthegap-OCI/issues/185)) ([80f7e5b](https://github.com/warmchang/mindthegap-OCI/commit/80f7e5b3dc02b89cf5646c05a3bf84f5ef2b63cc))
* Fix import image-bundle command ([#88](https://github.com/warmchang/mindthegap-OCI/issues/88)) ([73e7f6c](https://github.com/warmchang/mindthegap-OCI/commit/73e7f6c2e7d48fc12c6325607b0486cda810ad48))
* Hide crane warnings by default ([#507](https://github.com/warmchang/mindthegap-OCI/issues/507)) ([929c914](https://github.com/warmchang/mindthegap-OCI/commit/929c91415a50382886e3901001c2d317acd440ed))
* Import all platforms from image bundle ([#292](https://github.com/warmchang/mindthegap-OCI/issues/292)) ([696cf03](https://github.com/warmchang/mindthegap-OCI/commit/696cf03751c10a6f5caae4b275650a9d5ce4717f))
* prevent 404 on ECR during image creation ([#100](https://github.com/warmchang/mindthegap-OCI/issues/100)) ([e000ef7](https://github.com/warmchang/mindthegap-OCI/commit/e000ef7c83441c5e10a83eccfa5f4888fa5e310f))
* Properly handle friendly image names ([#195](https://github.com/warmchang/mindthegap-OCI/issues/195)) ([66e0bab](https://github.com/warmchang/mindthegap-OCI/commit/66e0babf8b184be9043e1dbe143858159b77a039))
* Remedy potential slowloris attack ([47a8362](https://github.com/warmchang/mindthegap-OCI/commit/47a836238d67b25738beb4a91bd0d56fa8f39c35))
* remove hardcoded AWS region when creating ECR lifecycle policies ([#353](https://github.com/warmchang/mindthegap-OCI/issues/353)) ([2fe144a](https://github.com/warmchang/mindthegap-OCI/commit/2fe144aa8cfc0f4d81d354b4a493a1ac96984701))
* Remove skopeo static builds from git-lfs ([703cfe0](https://github.com/warmchang/mindthegap-OCI/commit/703cfe0f6c9a437b127c37f8eff249965e4ebd9b))
* Remove unnecessary gzip compression of bundles by default ([#174](https://github.com/warmchang/mindthegap-OCI/issues/174)) ([06ee2d7](https://github.com/warmchang/mindthegap-OCI/commit/06ee2d78940ff4b16ed3ec5fe65549611194034d))
* Separate out source and destination remote options ([94087d3](https://github.com/warmchang/mindthegap-OCI/commit/94087d3002ab269a102c089a882e170d7646c485))
* Specify exe suffix for skopeo on Windows ([#233](https://github.com/warmchang/mindthegap-OCI/issues/233)) ([f8cb026](https://github.com/warmchang/mindthegap-OCI/commit/f8cb026add52468b7f17166be681943b2d7b7e84))
* support creating namespaced ECR repositories ([#386](https://github.com/warmchang/mindthegap-OCI/issues/386)) ([c79a8f2](https://github.com/warmchang/mindthegap-OCI/commit/c79a8f251961d1ec759de5c4d5c2694c6fe90375))
* Support docker certs.d host certs again ([#313](https://github.com/warmchang/mindthegap-OCI/issues/313)) ([6ab5849](https://github.com/warmchang/mindthegap-OCI/commit/6ab5849dcd5650290e691c6b95acd899db8fa287))
* Support http registries for pushing bundles ([#405](https://github.com/warmchang/mindthegap-OCI/issues/405)) ([6ef9359](https://github.com/warmchang/mindthegap-OCI/commit/6ef9359c5fca37a16c792c68e18382507a593eda))
* Unset REGISTRY_ prefixed env vars ([#424](https://github.com/warmchang/mindthegap-OCI/issues/424)) ([d48a2b0](https://github.com/warmchang/mindthegap-OCI/commit/d48a2b02ed3e7b09c451b0a35ac6dfd71747016f))
* Upgrade containerd dep to v1.6.6 to fix CVE ([#128](https://github.com/warmchang/mindthegap-OCI/issues/128)) ([638f80e](https://github.com/warmchang/mindthegap-OCI/commit/638f80ef5d991605fb4c0a369fee2f6aef46cd52))
* Upgrade github.com/sylabs/sif/v2@v2.8.1 ([#227](https://github.com/warmchang/mindthegap-OCI/issues/227)) ([757f1a2](https://github.com/warmchang/mindthegap-OCI/commit/757f1a26691231e2e6e791f6c6e89c7100353425))
* Use ocischema.Manifest for plain image details ([#145](https://github.com/warmchang/mindthegap-OCI/issues/145)) ([976a253](https://github.com/warmchang/mindthegap-OCI/commit/976a2539d0c9241689077faa1abebdef629ae161))
* using wrong argument when building ECR repository ([#387](https://github.com/warmchang/mindthegap-OCI/issues/387)) ([0c102f3](https://github.com/warmchang/mindthegap-OCI/commit/0c102f3864bbe1d5f6c1a79467dc958375a583db))
* Write out merged image bundle config when serving ([#168](https://github.com/warmchang/mindthegap-OCI/issues/168)) ([a3ab022](https://github.com/warmchang/mindthegap-OCI/commit/a3ab022f547bc4f8f65431966d5c4bcc16b7456b))


### Build System

* release v0.17.1 ([49eb8a2](https://github.com/warmchang/mindthegap-OCI/commit/49eb8a245c5b25c15a9ac0a06e5617761c690be8))
* release v0.18.0 ([29aed52](https://github.com/warmchang/mindthegap-OCI/commit/29aed52bd735057fbddb80d32df5ab6f05db3f04))
* release v1.0.0 ([b3546c3](https://github.com/warmchang/mindthegap-OCI/commit/b3546c32173e8df077a650d7469d592deda5133f))
* release v1.0.0-rc.1 ([a3079f5](https://github.com/warmchang/mindthegap-OCI/commit/a3079f5b0f011bd472dc17541733d6e23f00ded3))
* release v1.12.0 ([1802692](https://github.com/warmchang/mindthegap-OCI/commit/18026926556bb11c6cc7b971603511a523a748f8))
* release v1.2.0 ([e0d99dc](https://github.com/warmchang/mindthegap-OCI/commit/e0d99dce74e9feece9968b7c7a5dc67e2e1ca8b2))
* release v1.3.0 ([9ca6454](https://github.com/warmchang/mindthegap-OCI/commit/9ca64540120d85d9f9b83c9c30137028b0e0c8e1))

## 1.12.0 (2023-10-17)

<!-- Release notes generated using configuration in .github/release.yaml at main -->

## What's Changed
### Other Changes
* build: Add vscode settings for testing by @jimmidyson in https://github.com/mesosphere/mindthegap/pull/531
* build: Use ko for image building by @jimmidyson in https://github.com/mesosphere/mindthegap/pull/533
* build: go v1.21.1 by @jimmidyson in https://github.com/mesosphere/mindthegap/pull/530
* build: Update pre-commit hooks by @jimmidyson in https://github.com/mesosphere/mindthegap/pull/534
* build: Remove unused docker make stuff by @jimmidyson in https://github.com/mesosphere/mindthegap/pull/537


**Full Changelog**: https://github.com/mesosphere/mindthegap/compare/v1.11.1...v1.12.0

## 1.11.1 (2023-09-15)

<!-- Release notes generated using configuration in .github/release.yaml at main -->

## What's Changed
### Fixes 🔧
* fix: Hide crane warnings by default by @jimmidyson in https://github.com/mesosphere/mindthegap/pull/507
### Other Changes
* build(deps): Bump goproxy dependency for CVE by @jimmidyson in https://github.com/mesosphere/mindthegap/pull/479
* test: Try to fix proxy flaky tests by @jimmidyson in https://github.com/mesosphere/mindthegap/pull/480
* build: Upgrade tools by @jimmidyson in https://github.com/mesosphere/mindthegap/pull/481


**Full Changelog**: https://github.com/mesosphere/mindthegap/compare/v1.11.0...v1.11.1

## 1.11.0 (2023-07-31)

<!-- Release notes generated using configuration in .github/release.yaml at main -->

## What's Changed
### Exciting New Features 🎉
* feat: Add mindthegap to useragent by @jimmidyson in https://github.com/mesosphere/mindthegap/pull/476
* feat: Image push concurrency by @jimmidyson in https://github.com/mesosphere/mindthegap/pull/477
### Other Changes
* build: Try devbox instead of devenv by @jimmidyson in https://github.com/mesosphere/mindthegap/pull/465
* build(deps): Upgrade commit based deps by @jimmidyson in https://github.com/mesosphere/mindthegap/pull/470
* ci: Use latest devbox install action by @jimmidyson in https://github.com/mesosphere/mindthegap/pull/471


**Full Changelog**: https://github.com/mesosphere/mindthegap/compare/v1.10.1...v1.11.0

## 1.10.1 (2023-07-11)

<!-- Release notes generated using configuration in .github/release.yaml at main -->

## What's Changed
### Fixes 🔧
* fix: Correctly handle multiple registries by @jimmidyson in https://github.com/mesosphere/mindthegap/pull/459


**Full Changelog**: https://github.com/mesosphere/mindthegap/compare/v1.10.0...v1.10.1

## 1.10.0 (2023-07-11)

<!-- Release notes generated using configuration in .github/release.yaml at main -->

## What's Changed
### Exciting New Features 🎉
* feat: Parallel image pulls by @jimmidyson in https://github.com/mesosphere/mindthegap/pull/456
### Other Changes
* ci: Add write permissions for publishing test results by @jimmidyson in https://github.com/mesosphere/mindthegap/pull/454
* build: Upgrade devenv environment by @jimmidyson in https://github.com/mesosphere/mindthegap/pull/458
* build(deps): Upgrade dependencies with CVEs by @jimmidyson in https://github.com/mesosphere/mindthegap/pull/457


**Full Changelog**: https://github.com/mesosphere/mindthegap/compare/v1.9.0...v1.10.0

## 1.9.0 (2023-06-22)

<!-- Release notes generated using configuration in .github/release.yaml at main -->

## What's Changed
### Exciting New Features 🎉
* feat: Add --on-existing-tag flag to push bundle by @jimmidyson in https://github.com/mesosphere/mindthegap/pull/449
### Fixes 🔧
* fix: Do not log base64 encoded ECR token by @jimmidyson in https://github.com/mesosphere/mindthegap/pull/450
### Other Changes
* test: Use IPv4 localhost address instead of localhost hostname by @jimmidyson in https://github.com/mesosphere/mindthegap/pull/447
* ci: Run macOS tests on Ventura by @jimmidyson in https://github.com/mesosphere/mindthegap/pull/448
* build: Tweak golangci-lint configuration by @jimmidyson in https://github.com/mesosphere/mindthegap/pull/451


**Full Changelog**: https://github.com/mesosphere/mindthegap/compare/v1.8.0...v1.9.0

## 1.8.0 (2023-06-15)

<!-- Release notes generated using configuration in .github/release.yaml at main -->

## What's Changed
### Exciting New Features 🎉
* feat: Automatic retrieval of ECR credentials if using ECR as registry by @jimmidyson in https://github.com/mesosphere/mindthegap/pull/440
### Other Changes
* test: Ensure proxy settings do not affect bundle creation before test by @jimmidyson in https://github.com/mesosphere/mindthegap/pull/438
* build: Try using devenv.sh by @jimmidyson in https://github.com/mesosphere/mindthegap/pull/428


**Full Changelog**: https://github.com/mesosphere/mindthegap/compare/v1.7.7...v1.8.0

## 1.7.7 (2023-06-08)

<!-- Release notes generated using configuration in .github/release.yaml at main -->

## What's Changed
### Other Changes
* ci: Add write actions permissions for release please workflow by @jimmidyson in https://github.com/mesosphere/mindthegap/pull/433


**Full Changelog**: https://github.com/mesosphere/mindthegap/compare/v1.7.6...v1.7.7

## 1.7.6 (2023-06-08)

<!-- Release notes generated using configuration in .github/release.yaml at main -->

## What's Changed
### Other Changes
* test: Fix proxy env vars set up in push bundle by @jimmidyson in https://github.com/mesosphere/mindthegap/pull/429
* test: Ensure proxy tests are all run in serial by @jimmidyson in https://github.com/mesosphere/mindthegap/pull/431
* ci: Trigger release tag workflow on release please success by @jimmidyson in https://github.com/mesosphere/mindthegap/pull/432


**Full Changelog**: https://github.com/mesosphere/mindthegap/compare/v1.7.5...v1.7.6

## 1.7.5 (2023-06-05)

<!-- Release notes generated using configuration in .github/release.yaml at main -->

## What's Changed
### Fixes 🔧
* fix: Unset REGISTRY_ prefixed env vars by @jimmidyson in https://github.com/mesosphere/mindthegap/pull/424
### Other Changes
* test: Use table test for push image bundle by @jimmidyson in https://github.com/mesosphere/mindthegap/pull/414
* test: Add test using an http proxy by @jimmidyson in https://github.com/mesosphere/mindthegap/pull/415
* test: Add test for custom Docker config headers by @jimmidyson in https://github.com/mesosphere/mindthegap/pull/417
* build: Upgrade all tools by @jimmidyson in https://github.com/mesosphere/mindthegap/pull/423


**Full Changelog**: https://github.com/mesosphere/mindthegap/compare/v1.7.4...v1.7.5

## 1.7.4 (2023-05-22)

<!-- Release notes generated using configuration in .github/release.yaml at main -->

## What's Changed
### Fixes 🔧
* fix: Support http registries for pushing bundles by @jimmidyson in https://github.com/mesosphere/mindthegap/pull/405
### Other Changes
* ci: Use updated parse-asdf-versions by @jimmidyson in https://github.com/mesosphere/mindthegap/pull/407
* build: Upgrade tools and distroless base image by @jimmidyson in https://github.com/mesosphere/mindthegap/pull/413


**Full Changelog**: https://github.com/mesosphere/mindthegap/compare/v1.7.3...v1.7.4

## 1.7.3 (2023-04-24)

<!-- Release notes generated using configuration in .github/release.yaml at main -->

## What's Changed
### Fixes 🔧
* fix: Separate out source and destination remote options by @dkoshkin in https://github.com/mesosphere/mindthegap/pull/389
### Other Changes
* build: Upgrade tools and distroless base image by @jimmidyson in https://github.com/mesosphere/mindthegap/pull/390
* build: Insert flexible year into license header (#391 by @jimmidyson in https://github.com/mesosphere/mindthegap/pull/391


**Full Changelog**: https://github.com/mesosphere/mindthegap/compare/v1.7.2...v1.7.3

## 1.7.2 (2023-04-20)

<!-- Release notes generated using configuration in .github/release.yaml at main -->

## What's Changed
### Fixes 🔧
* fix: using wrong argument when building ECR repository by @dkoshkin in https://github.com/mesosphere/mindthegap/pull/387


**Full Changelog**: https://github.com/mesosphere/mindthegap/compare/v1.7.1...v1.7.2

## 1.7.1 (2023-04-19)

<!-- Release notes generated using configuration in .github/release.yaml at main -->

## What's Changed
### Fixes 🔧
* fix: support creating namespaced ECR repositories by @dkoshkin in https://github.com/mesosphere/mindthegap/pull/386
### Other Changes
* build: Upgrade all tools by @jimmidyson in https://github.com/mesosphere/mindthegap/pull/380
* build(deps): pin github.com/docker/* to v20.10.24+incompatible by @jimmidyson in https://github.com/mesosphere/mindthegap/pull/379


**Full Changelog**: https://github.com/mesosphere/mindthegap/compare/v1.7.0...v1.7.1

## [0.18.0](https://github.com/mesosphere/mindthegap/compare/v0.17.2...v0.18.0) (2022-08-10)


### Bug Fixes

* Remedy potential slowloris attack ([47a8362](https://github.com/mesosphere/mindthegap/commit/47a836238d67b25738beb4a91bd0d56fa8f39c35))
* Remove skopeo static builds from git-lfs ([703cfe0](https://github.com/mesosphere/mindthegap/commit/703cfe0f6c9a437b127c37f8eff249965e4ebd9b))


### Build System

* release v0.18.0 ([29aed52](https://github.com/mesosphere/mindthegap/commit/29aed52bd735057fbddb80d32df5ab6f05db3f04))

## [0.17.2](https://github.com/mesosphere/mindthegap/compare/v0.17.1...v0.17.2) (2022-07-27)


### Bug Fixes

* Use ocischema.Manifest for plain image details ([#145](https://github.com/mesosphere/mindthegap/issues/145)) ([976a253](https://github.com/mesosphere/mindthegap/commit/976a2539d0c9241689077faa1abebdef629ae161))

## [0.17.1](https://github.com/mesosphere/mindthegap/compare/v0.17.0...v0.17.1) (2022-07-01)


### Build System

* release v0.17.1 ([49eb8a2](https://github.com/mesosphere/mindthegap/commit/49eb8a245c5b25c15a9ac0a06e5617761c690be8))

## [0.17.0](https://github.com/mesosphere/mindthegap/compare/v0.16.0...v0.17.0) (2022-06-14)


### Features

* Build skopeo with go v1.18 ([9d04798](https://github.com/mesosphere/mindthegap/commit/9d0479845fa34337ab56ae4a604b0147a6e061bb))
* go v1.18 ([#110](https://github.com/mesosphere/mindthegap/issues/110)) ([f84ffa3](https://github.com/mesosphere/mindthegap/commit/f84ffa3019590f441fdfc935f112e3c7f049b6fc))


### Bug Fixes

* Disable upx for all platforms ([#129](https://github.com/mesosphere/mindthegap/issues/129)) ([f5496c3](https://github.com/mesosphere/mindthegap/commit/f5496c373f195e6e639d4d5911b68ad37c81ccc6))
* Upgrade containerd dep to v1.6.6 to fix CVE ([#128](https://github.com/mesosphere/mindthegap/issues/128)) ([638f80e](https://github.com/mesosphere/mindthegap/commit/638f80ef5d991605fb4c0a369fee2f6aef46cd52))

## [0.16.0](https://github.com/mesosphere/mindthegap/compare/v0.15.2...v0.16.0) (2022-05-09)


### Features

* Skopeo v1.8.0 ([c8bdd81](https://github.com/mesosphere/mindthegap/commit/c8bdd8135908de865eb8c07fdb7bd3320978ef0b))

### [0.15.2](https://github.com/mesosphere/mindthegap/compare/v0.15.1...v0.15.2) (2022-04-25)


### Bug Fixes

* prevent 404 on ECR during image creation ([#100](https://github.com/mesosphere/mindthegap/issues/100)) ([e000ef7](https://github.com/mesosphere/mindthegap/commit/e000ef7c83441c5e10a83eccfa5f4888fa5e310f))

### [0.15.1](https://github.com/mesosphere/mindthegap/compare/v0.15.0...v0.15.1) (2022-04-05)


### Bug Fixes

* commitDate in goreleaser ([#92](https://github.com/mesosphere/mindthegap/issues/92)) ([017d2f3](https://github.com/mesosphere/mindthegap/commit/017d2f36da2eb61b556f034e3e3232bd84ce7796))

## [0.15.0](https://github.com/mesosphere/mindthegap/compare/v0.14.0...v0.15.0) (2022-03-31)


### Features

* skopeo v1.7.0 ([9440c2b](https://github.com/mesosphere/mindthegap/commit/9440c2b82ab10c6526fe34abcaec83deac0150f9))


### Bug Fixes

* Fix import image-bundle command ([#88](https://github.com/mesosphere/mindthegap/issues/88)) ([73e7f6c](https://github.com/mesosphere/mindthegap/commit/73e7f6c2e7d48fc12c6325607b0486cda810ad48))

## [0.14.0](https://github.com/mesosphere/mindthegap/compare/v0.13.1...v0.14.0) (2022-03-22)


### Features

* Auto-create ECR registries if they do not exist ([#79](https://github.com/mesosphere/mindthegap/issues/79)) ([b595607](https://github.com/mesosphere/mindthegap/commit/b595607e523497480998a4d35522c8d3553635ba))

### [0.13.1](https://github.com/mesosphere/mindthegap/compare/v0.13.0...v0.13.1) (2022-03-18)


### Bug Fixes

* Do not upx pack darwin binaries ([503d726](https://github.com/mesosphere/mindthegap/commit/503d72631ef1220244c578c518bdb3110026eaa6))

## [0.13.0](https://github.com/mesosphere/mindthegap/compare/v0.12.0...v0.13.0) (2022-03-18)


### Build System

* release 0.13.0 ([f7bed9a](https://github.com/mesosphere/mindthegap/commit/f7bed9a41135e142ee146ba433fa7f0fdc2bad62))

## [0.12.0](https://github.com/mesosphere/mindthegap/compare/v0.11.0...v0.12.0) (2022-03-15)


### Features

* Add username and password auth to helm repository config ([#65](https://github.com/mesosphere/mindthegap/issues/65)) ([8f74681](https://github.com/mesosphere/mindthegap/commit/8f74681596f945c0571f8e8ddb8f6dd7a6ce6912))

## [0.11.0](https://github.com/mesosphere/mindthegap/compare/v0.10.0...v0.11.0) (2022-03-03)


### Features

* Sort images for deterministic ordering of create and push ([adb29f8](https://github.com/mesosphere/mindthegap/commit/adb29f8f96ccd3606c60acc2051239ba9890f09b))


### Bug Fixes

* Upgrade containerd dep to fix CVEs ([0079ee8](https://github.com/mesosphere/mindthegap/commit/0079ee8704af91d30109974babf9f61244daaaf0))

## [0.10.0](https://github.com/mesosphere/mindthegap/compare/v0.9.1...v0.10.0) (2022-03-03)


### Features

* Upgrade to skopeo v1.6.1 ([009def6](https://github.com/mesosphere/mindthegap/commit/009def69cdb2e0e1fe4721fa183d5c0be911c467))
