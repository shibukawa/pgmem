---
id: metric:alternative-comparison
type: metric
title: Alternative Comparison Numbers
---
pgmem measured next to docker run, Testcontainers and devbox on one machine with one harness; the website tables render the same data.

```yaml
metric:
  measured: 2026-09-13, Apple M3 8 cores 16 GB, macOS (Darwin 27), OrbStack 2.2.3 Docker 29.4.0, devbox 0.17.5, Go 1.27, median of 5 runs
  harness: bench/alternatives (run.sh -> results/summary.json -> website/src/data/benchmarks.json)
  workload: bench 1000 rows + users 1000 + orders 10000; isolate = copy + connect + join, median of 20
  servers: pgmem 18.3 in-process; docker/testcontainers postgres:18-alpine 18.6; devbox nix postgresql 18.6
  start_to_first_query_ms: {pgmem: 48, pgmem_binary_ready: 129, devbox_initdb_plus_start: 760, docker_run: 1523, testcontainers_with_ryuk: 1530}
  fresh_copy_per_test_ms: {pgmem_fork: 13.6, testcontainers_snapshot_restore: 30.5, docker_template_db: 56.3, devbox_template_db: 60.9}
  indexed_select_us: {pgmem_tcp: 33.1, pgmem_in_process: 9.3, devbox: 28.7, docker: 84.8, testcontainers: 88.1}
  sort_200k_rows_ms: {pgmem: 113, devbox: 134, docker: 250, testcontainers: 248}
  memory_idle_mb: {pgmem_in_process_growth: 164, pgmem_binary_rss: 166, devbox_rss_sum: 66, container_docker_stats: 24}
  memory_after_load_mb: {pgmem: 486, devbox: 81, container: 39-48}  # pgmem linear memory and freed forks are not returned while the process runs
  download_mb: {go_link_added_stripped: 37.9, go_link_added_unstripped: 58.0, binary_zip: 16.8, binary_raw: 41.5, image_arm64: 117.9, ryuk: 2.1, devbox_nix_closure: 112.7, devbox_nix_unpacked: 463}
  reading:
    - pgmem wins start, per-test copy and download; it loses idle memory (container figure excludes the Docker VM)
    - native devbox PostgreSQL matches pgmem on TCP latency; pgmem's in-process dialer is ~3x faster
    - container latency includes OrbStack port forwarding
    - defaults differ: containers and devbox use shared_buffers=128MB and a disk; pgmem 32MB and memory
  first_samples_noisy: first run per target is slower (cold caches); medians hide it
```
