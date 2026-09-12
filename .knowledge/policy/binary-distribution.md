---
id: policy:binary-distribution
type: policy
title: Binary Distribution
---
Each wrapper ships the platform pgmem binary inside its own package format; nothing is downloaded at run time by default.

```yaml
policy:
  build: GOOS/GOARCH cross-compile of cmd/pgmem with -ldflags "-s -w" (~37MB) in a CI matrix; one Go tag produces every package at the same version
  platforms: [linux-x86_64, linux-arm64, darwin-x86_64, darwin-arm64, windows-x86_64]  # classifier = BinaryLocator.classifier()
  python:
    layout: platform wheels (ruff/uv style); binary at pgmem/_bin/pgmem or pgmem.exe; tags manylinux_2_17, macosx, win_amd64
    size: ~37MB per wheel, under the PyPI 100MB file limit
    fallback: PGMEM_BINARY env for unsupported platforms and local builds
  java:
    layout: pgmem-native artifact with one classifier per platform (zonky style); resource /jp/shibu/pgmem/native/<classifier>/pgmem
    select: user adds the classifier via os-maven-plugin or Gradle osdetector; depending on all five (185MB) is discouraged
    extract: first use copies to PGMEM_CACHE_DIR, else ${user.home}/.cache/pgmem/<Implementation-Version>/, else java.io.tmpdir; sets the exec bit; atomic rename so concurrent JVMs are safe
    fallback: PGMEM_BINARY env, then pgmem on PATH
  version_check: wrapper version equals binary version; protocol integer from the ready event (api:control-protocol) checked before use
  why_no_download: offline CI, supply-chain review, reproducibility; an opt-in downloader can come later
  exercised: darwin/arm64 wheel and native jar verified end to end 2026-09-12 in a clean venv and a bare classpath; CI matrix for the other platforms is still to do
```
