---
id: metric:alternative-comparison
type: metric
title: Alternative Comparison Numbers
---
Startup and memory of pgmem next to Docker-based and binary-based test databases; pgmem rows are measured, the others are typical public figures until the bench plan runs.

```yaml
metric:
  status: pgmem measured (metric:server-footprint, metric:fork-cost); Docker and embedded rows not measured in this repository as of 2026-09-12 (Docker daemon unavailable on the dev machine); treat them as order-of-magnitude
  rows:
    pgmem_start:
      time_s: 0.1
      memory_mb: 152
      source: measured
    pgmem_fork:
      time_s: 0.02
      memory_mb: ~100  # estimate: fresh linear memory ~65 MB touched plus ~39 MB vfs clone; measure before quoting
      source: derived
    testcontainers_postgres_container:
      time_s: 1-3    # image already pulled; ryuk container adds one more start
      first_pull_mb: ~100  # postgres alpine image size
      memory_mb: 50-200  # container RSS; excludes the Docker VM, 1-4 GB on macOS and Windows
      per_test_isolation: CREATE DATABASE from a template, ~0.1-0.3 s, or transaction rollback
      source: public docs and typical experience, unmeasured here
    embedded_postgres_binary:
      time_s: 1-2    # unpack cached, initdb plus postmaster start
      memory_mb: 30-150
      disk: temp data directory per server
      source: public docs, unmeasured here
  bench_plan:
    - start OrbStack or Docker, pull postgres:18-alpine once
    - Go: testcontainers-go postgres module, time from Run to first successful connect, 10 runs, report median
    - memory: docker stats for the container, ps for the Docker VM helper; compare with ps RSS of the pgmem binary
    - per test: time CREATE DATABASE ... TEMPLATE against the container versus api:clone; repeat with 8 parallel tests
    - Java: zonky embedded-postgres start time with a warm cache; Python: pytest-postgresql
    - record machine, versions and date; replace the unmeasured rows above
  publish_rule: the website must label unmeasured rows as such until this plan replaces them
```
