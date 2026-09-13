---
id: policy:binary-distribution
type: policy
title: Binary Distribution
---
Each wrapper ships the platform pgmem binary inside its own package format; nothing is downloaded at run time by default.

```yaml
policy:
  build: scripts/build-binaries.sh cross-compiles cmd/pgmem with -trimpath -ldflags "-s -w -X main.version=vX.Y.Z" (~41MB) in the six-job matrix of flow:release; one tag produces every package at the same version
  platforms: [linux-x86_64, linux-arm64, darwin-x86_64, darwin-arm64, windows-x86_64, windows-arm64]  # classifier = BinaryLocator.classifier(); npm uses process.platform-process.arch names
  archives: GitHub Release assets pgmem-<version>-<goos>-<goarch>.tar.gz (zip for Windows, ~17MB) with SHA256SUMS (scripts/package-archives.sh)
  python:
    layout: platform wheels (ruff/uv style); binary at pgmem/_bin/pgmem or pgmem.exe; hatch_build.py deletes a binary left by another platform's build, since hatchling ignores the nested _bin/.gitignore
    tags: manylinux_2_17 and musllinux_1_1 in one tag set (static binary, so Alpine too), macosx_13_0 (the Go 1.27 floor), win_amd64, win_arm64
    size: ~17MB per wheel (41MB unpacked), under the PyPI 100MB file limit
    fallback: PGMEM_BINARY env for unsupported platforms and local builds
  java:
    layout: pgmem-native with packaging pom and only classifier jars (~17MB each, as osmem-server-binaries on Maven Central); resource /io/github/shibukawa/pgmem/native/<classifier>/pgmem; -PnativeDist=<dist dir> packs all six, otherwise the host binary
    select: user adds the classifier via os-maven-plugin or Gradle osdetector; depending on all six (~100MB) is discouraged
    extract: first use copies to PGMEM_CACHE_DIR, else ${user.home}/.cache/pgmem/<Implementation-Version>/, else java.io.tmpdir; sets the exec bit; atomic rename so concurrent JVMs are safe
    fallback: PGMEM_BINARY env, then pgmem on PATH
  node:
    layout: main package plus per-platform packages in optionalDependencies with os and cpu (esbuild and @osmem/core layout); no postinstall because package managers hold install scripts (system:node-test-runners)
    name: scoped (api:node-wrapper); unscoped pgmem collides with pg-mem under npm's name similarity rule
    fallback: PGMEM_BINARY env
    packing: scripts/build-npm.sh copies the dist binaries into platforms/<platform>/bin with mode 755 and packs 7 tarballs (~17MB per platform)
  notices: policy:third-party-notices
  version_check: wrapper version equals binary version (policy:versioning); protocol integer from the ready event (api:control-protocol) checked before use
  why_no_download: offline CI, supply-chain review, reproducibility; an opt-in downloader can come later
  exercised: darwin/arm64 wheel and native jar end to end 2026-09-12 in a clean venv and a bare classpath; 2026-09-14 all six binaries, archives, npm tarballs, wheels and a Maven bundle signed with throwaway ed25519 and RSA keys were built locally, and the darwin/arm64 wheel (venv), npm tarballs (npm install) and native jar (bare classpath) each started a server reporting v0.1.0; CI (2026-09-14) runs the Go tests on linux/amd64, darwin/arm64 and windows/amd64 and the wrapper tests on linux/amd64, while linux/arm64, windows/arm64 and darwin/amd64 only build
```
