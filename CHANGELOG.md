# Changelog

## [1.4.1](https://github.com/FelipeFuhr/ffreis-workflows-go/compare/v1.4.0...v1.4.1) (2026-08-09)


### Bug Fixes

* **ci:** remove redundant standalone scorecards.yml ([#104](https://github.com/FelipeFuhr/ffreis-workflows-go/issues/104)) ([3e34d2a](https://github.com/FelipeFuhr/ffreis-workflows-go/commit/3e34d2a00ddd0fc9a659a8b7a8c75107690f2209))
* **runner:** default the runner input to self-hosted, not ubuntu-latest ([#102](https://github.com/FelipeFuhr/ffreis-workflows-go/issues/102)) ([eb48820](https://github.com/FelipeFuhr/ffreis-workflows-go/commit/eb48820e594b2d1d387518eba957adee1906c826))

## [1.4.0](https://github.com/FelipeFuhr/ffreis-workflows-go/compare/v1.3.0...v1.4.0) (2026-08-07)


### Features

* **workflows:** add goprivate support to coverage and mutation jobs ([#100](https://github.com/FelipeFuhr/ffreis-workflows-go/issues/100)) ([63280b0](https://github.com/FelipeFuhr/ffreis-workflows-go/commit/63280b0e2572179d54d7d461098983d9d6c79611))

## [1.3.0](https://github.com/FelipeFuhr/ffreis-workflows-go/compare/v1.2.0...v1.3.0) (2026-08-07)


### Features

* add go-integration-coverage.yml reusable workflow ([#99](https://github.com/FelipeFuhr/ffreis-workflows-go/issues/99)) ([356d926](https://github.com/FelipeFuhr/ffreis-workflows-go/commit/356d9269b1d7db5b9c932681d7319e901826f742))
* **security:** add optional GOPRIVATE + private-module auth support ([#95](https://github.com/FelipeFuhr/ffreis-workflows-go/issues/95)) ([d520667](https://github.com/FelipeFuhr/ffreis-workflows-go/commit/d52066753ec83739bcf7a37a6274fb9bdf717710))

## [1.2.0](https://github.com/FelipeFuhr/ffreis-workflows-go/compare/v1.1.2...v1.2.0) (2026-08-01)


### Features

* **go-test:** two-tier S3 (L2) build cache — engaged only on a GH-cache (L1) miss ([#68](https://github.com/FelipeFuhr/ffreis-workflows-go/issues/68)) ([a7d3f9e](https://github.com/FelipeFuhr/ffreis-workflows-go/commit/a7d3f9e70ca8715152fc76d2c6acdd6e6fd31b7e))
* **go:** add optional GOPRIVATE + private-module auth support ([#90](https://github.com/FelipeFuhr/ffreis-workflows-go/issues/90)) ([600e3b4](https://github.com/FelipeFuhr/ffreis-workflows-go/commit/600e3b4802678747a7984ef575a46a1b19f531ff))
* **security:** add optional GOPRIVATE + private-module auth support ([#95](https://github.com/FelipeFuhr/ffreis-workflows-go/issues/95)) ([d520667](https://github.com/FelipeFuhr/ffreis-workflows-go/commit/d52066753ec83739bcf7a37a6274fb9bdf717710))
* **semgrep:** replace Docker-container semgrep-action with native CLI ([#89](https://github.com/FelipeFuhr/ffreis-workflows-go/issues/89)) ([db4a095](https://github.com/FelipeFuhr/ffreis-workflows-go/commit/db4a0956ceefaa7dbe59021f238340b9be694bb2))


### Bug Fixes

* **ci:** pin ffreis-workflows-general to v1.7.0 ([#80](https://github.com/FelipeFuhr/ffreis-workflows-go/issues/80)) ([237b32f](https://github.com/FelipeFuhr/ffreis-workflows-go/commit/237b32f397db402089d8e48f4983a23eca682671))
* **go-sonar:** correct corrupted runner input default ([#87](https://github.com/FelipeFuhr/ffreis-workflows-go/issues/87)) ([ba85750](https://github.com/FelipeFuhr/ffreis-workflows-go/commit/ba8575070827b3f4c74a7568d425c7d3e12d50a6))
* **shellcheck:** add SC2317 reason to common_die disable comment ([#79](https://github.com/FelipeFuhr/ffreis-workflows-go/issues/79)) ([c82cd30](https://github.com/FelipeFuhr/ffreis-workflows-go/commit/c82cd30ec1ed19b9219b73458257575a578f10f4))

## [1.1.2](https://github.com/FelipeFuhr/ffreis-workflows-go/compare/v1.1.1...v1.1.2) (2026-06-08)


### Bug Fixes

* correct sonar project key in examples/hello ([#63](https://github.com/FelipeFuhr/ffreis-workflows-go/issues/63)) ([22e7a90](https://github.com/FelipeFuhr/ffreis-workflows-go/commit/22e7a90bd761691e6831e4a94e3b1c22bb29e7c1))
* **go-sonar:** correct grep regex + add Mod function with test ([d707b48](https://github.com/FelipeFuhr/ffreis-workflows-go/commit/d707b489cae8f852d64a8c1bb6788a7a098ccac9))
* normalize SonarCloud project key to lowercase ([#61](https://github.com/FelipeFuhr/ffreis-workflows-go/issues/61)) ([cede25d](https://github.com/FelipeFuhr/ffreis-workflows-go/commit/cede25d081fba2cdb69428680614ce5f5de3d00d))
* resolve SonarQube issues ([#65](https://github.com/FelipeFuhr/ffreis-workflows-go/issues/65)) ([6547ede](https://github.com/FelipeFuhr/ffreis-workflows-go/commit/6547eded302de0f91b6465b546a6f5d482c54459))

## [1.1.1](https://github.com/FelipeFuhr/ffreis-workflows-go/compare/v1.1.0...v1.1.1) (2026-06-06)


### Bug Fixes

* **ci:** unique concurrency group in go-container reusable to avoid caller deadlock ([#57](https://github.com/FelipeFuhr/ffreis-workflows-go/issues/57)) ([8bbaa46](https://github.com/FelipeFuhr/ffreis-workflows-go/commit/8bbaa4626e530794be5b5aaf2d8f816d6bff20e2))
* **grype:** bump workflows-general SHA to prevent self-scan CVEs ([#53](https://github.com/FelipeFuhr/ffreis-workflows-go/issues/53)) ([124cd83](https://github.com/FelipeFuhr/ffreis-workflows-go/commit/124cd834ca0cbb9a4257ad20e7c6c35e0ea53d5e))

## [1.1.0](https://github.com/FelipeFuhr/ffreis-workflows-go/compare/v1.0.0...v1.1.0) (2026-05-23)


### Features

* migrate to platform standards ([#37](https://github.com/FelipeFuhr/ffreis-workflows-go/issues/37)) ([a543ef8](https://github.com/FelipeFuhr/ffreis-workflows-go/commit/a543ef895b5c2d6e4e7fe87245d440889c485b7c))

## 1.0.0 (2026-05-06)


### Features

* add ci and agents ([e7d00e0](https://github.com/FelipeFuhr/ffreis-workflows-go/commit/e7d00e09445d27de8e2b226ca02c6bd27450cdf8))
* add workflows ([f1fb8db](https://github.com/FelipeFuhr/ffreis-workflows-go/commit/f1fb8db3a39f959442b825fab795a3ef08f8f7dc))
* address review feedback - pin tool versions, fix whitespace trim, add fork guard and fail-on-issues input ([5343031](https://github.com/FelipeFuhr/ffreis-workflows-go/commit/53430319ec14ba5f2ed7de5fd9b41aa7014f1326))
* adjust examples, lefthook ([1d639c7](https://github.com/FelipeFuhr/ffreis-workflows-go/commit/1d639c7aa8e3a1c5884777e703c08318b60f682c))
* ci changes ([3ff10c4](https://github.com/FelipeFuhr/ffreis-workflows-go/commit/3ff10c474db2c927fb6a05da2ae117f790f609b6))
* ci changes ([b986ad4](https://github.com/FelipeFuhr/ffreis-workflows-go/commit/b986ad4aa35764ca23c827d631b0bee636602548))
* ci changes ([dac7612](https://github.com/FelipeFuhr/ffreis-workflows-go/commit/dac7612cf280bfe1feee936014ed670d21415316))
* ci changes ([86000d8](https://github.com/FelipeFuhr/ffreis-workflows-go/commit/86000d809165e9b1d350563f866f0ee302f1afdf))
* ci improvements ([ed2f7d2](https://github.com/FelipeFuhr/ffreis-workflows-go/commit/ed2f7d29467bf1fd6ec40df3953b3ee128ba3f04))
* **deps:** migrate to ffreis-platform-standards ([#30](https://github.com/FelipeFuhr/ffreis-workflows-go/issues/30)) ([e423173](https://github.com/FelipeFuhr/ffreis-workflows-go/commit/e4231730a8777ae7d4d963baee2f479361ef35e9))
* fix ([bb208a1](https://github.com/FelipeFuhr/ffreis-workflows-go/commit/bb208a1a7c63bd3d870fa562167b4cb382dbdfc7))
* fix ([4ce8a2e](https://github.com/FelipeFuhr/ffreis-workflows-go/commit/4ce8a2e954baa1e1a6e23d5c9ca7dc96df9263a1))
* improve ([829f07d](https://github.com/FelipeFuhr/ffreis-workflows-go/commit/829f07dc65febcbff9cd5311d0062fefacda9f4e))
* improvements ([500e6c9](https://github.com/FelipeFuhr/ffreis-workflows-go/commit/500e6c9ad3ad3da0ad08a418a62c71919663b7e4))
* more workflows ([94426ce](https://github.com/FelipeFuhr/ffreis-workflows-go/commit/94426ce78b830e782ce3db48d47a409b66f8fb5a))
* more workflows ([475fe17](https://github.com/FelipeFuhr/ffreis-workflows-go/commit/475fe1794f05716f4b53f65b56461c92fb85d705))
* sample repo for self test ([f6bcfec](https://github.com/FelipeFuhr/ffreis-workflows-go/commit/f6bcfec92a74c1a2b1a1130b4c66be23958ae286))
* setup repo ([fe63afb](https://github.com/FelipeFuhr/ffreis-workflows-go/commit/fe63afb1149a57b0b1b9776bd15eee50a1461ce5))


### Bug Fixes

* actionlint + pr labeler ([7aca16c](https://github.com/FelipeFuhr/ffreis-workflows-go/commit/7aca16cc32435ca9144438b2eede47eb610443e3))
* actionlint replace() + shellcheck warnings ([6facd86](https://github.com/FelipeFuhr/ffreis-workflows-go/commit/6facd86a318c73c097cc41a96938fea2893481b2))
* address all PR review comments ([03dbbe6](https://github.com/FelipeFuhr/ffreis-workflows-go/commit/03dbbe6a155f9de640f7a56379bb72a28f1b8c67))
* address go-container, go-build, go-fuzz, Makefile, and Containerfile review issues ([ed1eb63](https://github.com/FelipeFuhr/ffreis-workflows-go/commit/ed1eb63e4d14ba8a7b2069d62100e7c711f2202a))
* authorization token ([3dd69b0](https://github.com/FelipeFuhr/ffreis-workflows-go/commit/3dd69b0fad573588ef47d6e437ea216d20ddc3b9))
* ci ([b253b23](https://github.com/FelipeFuhr/ffreis-workflows-go/commit/b253b2329a11552b6276781c56f6dc6fd1718710))
* ci ([17ed746](https://github.com/FelipeFuhr/ffreis-workflows-go/commit/17ed7464b92df9e1b9ecaf5d8554ef00bc3f8cc1))
* ci ([6a9b12a](https://github.com/FelipeFuhr/ffreis-workflows-go/commit/6a9b12a1c8b8a6a8ba0002a8077a1713ddf31be8))
* ci ([319d8cb](https://github.com/FelipeFuhr/ffreis-workflows-go/commit/319d8cb0c2e6f63990eaad50722b3763d1f72a90))
* ci ([9b6b114](https://github.com/FelipeFuhr/ffreis-workflows-go/commit/9b6b114a140b663d8d271dab2868e10dc9373a3e))
* ci ([2b0442d](https://github.com/FelipeFuhr/ffreis-workflows-go/commit/2b0442dd47346ba6e32b8d600e4be72660832361))
* ci ([#27](https://github.com/FelipeFuhr/ffreis-workflows-go/issues/27)) ([af138c9](https://github.com/FelipeFuhr/ffreis-workflows-go/commit/af138c93d532b5f38aec48837c2e79af0dc46c91))
* detect fork PRs via repo match ([c6f09e4](https://github.com/FelipeFuhr/ffreis-workflows-go/commit/c6f09e408e122a06881d74cc674b5d9376468e5b))
* don\u2019t gate sonarqube on vars ([a8939d7](https://github.com/FelipeFuhr/ffreis-workflows-go/commit/a8939d710419f4abb0540a59c5b908f628ed9a0a))
* improve ([3f9db1b](https://github.com/FelipeFuhr/ffreis-workflows-go/commit/3f9db1bab72f3f4be9463329b6f4fe3a391afa6e))
* improve ([2f7aa10](https://github.com/FelipeFuhr/ffreis-workflows-go/commit/2f7aa10b58fa469893497fbdb0b3eeff594aa73d))
* improve ([b45eeb0](https://github.com/FelipeFuhr/ffreis-workflows-go/commit/b45eeb0c8189318f314acd029a1e2bae2b8c7878))
* improve ([25e60bd](https://github.com/FelipeFuhr/ffreis-workflows-go/commit/25e60bde9a9fec0efcbde9c0d48db0f31e0612da))
* improve ([cc24be2](https://github.com/FelipeFuhr/ffreis-workflows-go/commit/cc24be22af4d7139ddb19da99bb3eb043bca957a))
* improve ([2cabba3](https://github.com/FelipeFuhr/ffreis-workflows-go/commit/2cabba3e55a06c35a6e16cb1659e83d68674812f))
* lowercase sonar organization slug ([3eb85a4](https://github.com/FelipeFuhr/ffreis-workflows-go/commit/3eb85a4d92433f58ce845b926d51feba61656f96))
* make sonar opt-in + pin actions ([7841baa](https://github.com/FelipeFuhr/ffreis-workflows-go/commit/7841baa8f015ecb3683b1ec2e2f9a39446f756dd))
* more improvements ([7eeff74](https://github.com/FelipeFuhr/ffreis-workflows-go/commit/7eeff74ca019db2e207fa3c37ad5d35f27329857))
* more improvements ([88af8d5](https://github.com/FelipeFuhr/ffreis-workflows-go/commit/88af8d5b9efca1ca81756afb7e5f05891579acf3))
* never skip sonar when configured ([375fa71](https://github.com/FelipeFuhr/ffreis-workflows-go/commit/375fa71d542bc44a3b6d19f74d4635a0b4727721))
* pin tool versions, trim matrix whitespace, add fork guards and fail-on-issues input ([25608fa](https://github.com/FelipeFuhr/ffreis-workflows-go/commit/25608faad1fbb8d8abf5fa5827e5db951081ac84))
* run sonar jobs without secrets if ([cde8f72](https://github.com/FelipeFuhr/ffreis-workflows-go/commit/cde8f7289e7886a5d6c8df300e8671f4035e7ac3))
* run sonar token check after checkout ([6122af9](https://github.com/FelipeFuhr/ffreis-workflows-go/commit/6122af978fc7c782de6ff862d149c97cb9029e9e))
* sonarqube host url via var or input ([27ecee7](https://github.com/FelipeFuhr/ffreis-workflows-go/commit/27ecee7cefa34bddb4e3481476eec97afad3d2f5))
* upgrade golangci-lint-action to v7 and harden hook scripts ([2dc66b6](https://github.com/FelipeFuhr/ffreis-workflows-go/commit/2dc66b62be7f8685244fed807fcb160c7ba7af24))
* validate SONAR_HOST_URL scheme ([6a07eef](https://github.com/FelipeFuhr/ffreis-workflows-go/commit/6a07eef1a05a22a956eef4893ca3ade5a1be2fcd))
