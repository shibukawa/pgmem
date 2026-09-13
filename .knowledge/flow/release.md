---
id: flow:release
type: flow
title: Release
---
How a vX.Y.Z tag becomes every published artifact at one version (policy:versioning, policy:binary-distribution, policy:third-party-notices); .github/workflows/release.yml, set up 2026-09-14 after the osmem release.yml that published osmem 0.1.1.

```yaml
flow:
  trigger: push of tag vX.Y.Z; workflow_dispatch rehearses (dry_run, the default) or, run on a tag, republishes to chosen registries
  steps:
    - id: prepare
      action: version from the tag or input; refuse to publish unless the ref is tag vX.Y.Z
    - id: binaries
      action: six ubuntu jobs, scripts/build-binaries.sh <goos>-<goarch> with Go 1.27.x (sets the macOS 13 wheel floor), -X main.version=vX.Y.Z; artifacts bin-<target>
    - id: packages
      action: scripts/set-version.sh; scripts/package-archives.sh (tar.gz, zip for Windows, SHA256SUMS); scripts/build-npm.sh (7 tarballs); scripts/build-python-wheels.sh then twine check
    - id: github
      action: GitHub Release with archives, wheels and npm tarballs
    - id: npm
      environment: release
      action: npm publish of each tarball by trusted publishing (OIDC), platform packages before @pgmem/core, versions already on npm skipped; rehearsal uses --dry-run
    - id: pypi
      environment: release
      action: pypa/gh-action-pypi-publish by trusted publishing with skip-existing; not run in rehearsal
    - id: maven
      environment: release
      action: scripts/build-maven-bundle.sh (Gradle publish into build/central-staging, in-memory GPG signing, zip without maven-metadata) then scripts/publish-maven-central.sh (Portal upload AUTOMATIC; rehearsal USER_MANAGED, wait for VALIDATED, drop); skipped when the POM is on repo1
    - id: go
      action: go list -m github.com/shibukawa/pgmem@vX.Y.Z through proxy.golang.org
  ci_guard: ci.yml job release-packaging runs the same scripts on every push with placeholder binaries and a throwaway ed25519 key
  registry_setup:
    pypi: pending trusted publisher for project pgmem (shibukawa/pgmem, release.yml, environment release); valid before the first upload
    npm: a trusted publisher needs an existing package, so the first version of the 7 packages is published by hand with 2FA from rehearsal tarballs; configurations need direct publishing enabled, since those created after 2026-09-03 otherwise allow only npm stage publish
    maven: groupId io.github.shibukawa.pgmem under the namespace io.github.shibukawa, which the Portal verifies through the GitHub account (the user switched from jp.shibu on 2026-09-14 before the first publish, as osmem did); CENTRAL_USERNAME/CENTRAL_PASSWORD and GPG_PRIVATE_KEY/GPG_PASSPHRASE secrets exist in environment release since 2026-09-13
    pages: build_type workflow; environment github-pages allows main; docs.yml deploys website/
  limits: Go module zip 26MB compressed, 129MB of the 500MB proxy limit (2026-09-14); a cold cross-compile of one target takes ~46s and ~2.2GB RSS on an M-series Mac
```
