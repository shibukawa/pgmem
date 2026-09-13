---
id: metric:alternative-comparison
type: metric
title: Alternative Comparison Numbers
---
pgmem measured next to docker run, Testcontainers and devbox on one machine with one harness; the website tables render the same data.

```yaml
metric:
  measured: 2026-09-13, Apple M3 8 cores 16 GB, macOS (Darwin 27), OrbStack 2.2.3 Docker 29.4.0, devbox 0.17.5, Go 1.27
  harness: bench/alternatives (RUNS=5 MEM_RUNS=3 run.sh -> results/summary.json -> website/src/data/benchmarks.json)
  passes:
    timing: 5 runs per target, Docker VM left running between runs as on a developer machine; medians
    memory: docker and testcontainers 3 runs each after orbctl stop/start; growth of the OrbStack VM process from idle to schema loaded; refuses when non-testcontainers containers run
  workload: bench 1000 rows + users 1000 + orders 10000; isolate = copy + connect + join, median of 20
  servers: pgmem 18.3 in-process; docker/testcontainers postgres:18-alpine 18.6; devbox nix postgresql 18.6
  memory_metric: macOS phys_footprint via footprint(1), summed per process; RSS kept as rss_idle_mb
  start_to_first_query_ms: {pgmem: 73, pgmem_binary_ready: 150, devbox_initdb_plus_start: 784, docker_run: 1676, testcontainers_with_ryuk: 2008}
  fresh_copy_per_test_ms: {pgmem_fork: 14.7, testcontainers_snapshot_restore: 23.2, docker_template_db: 54.6, devbox_template_db: 61.2}
  indexed_select_us: {pgmem_tcp: 32.5, pgmem_in_process: 9.7, devbox: 27.8, docker: 86, testcontainers: 84}
  sort_200k_rows_ms: {pgmem: 116, devbox: 135, docker: 255, testcontainers: 246}
  memory_host_idle_mb:
    pgmem_test_process_growth: 140   # samples split 110 / 140 by GC timing; RSS growth 164
    pgmem_binary_process: 106        # RSS 166
    devbox_postmaster_and_children: 45  # RSS sum 66 double-counts shared buffers
    docker_vm_growth: 421            # container cgroup alone 76; VM total 1322
    testcontainers_vm_growth: 414    # postgres + ryuk containers 87; VM total 1310
  memory_after_load_mb: {pgmem: 440, devbox: 59, docker_vm_growth: 575, testcontainers_vm_growth: 515}
  vm_idle_baseline_mb: 893   # freshly started OrbStack VM before any container; held on top of the growth
  vm_memory_not_returned: VM phys_footprint stays at its peak after containers stop (observed +459 MB persisting)
  download_mb: {go_link_added_stripped: 37.9, go_link_added_unstripped: 58.0, binary_zip: 16.8, binary_raw: 41.5, image_arm64: 117.9, ryuk: 2.1, devbox_nix_closure: 112.7, devbox_nix_unpacked: 463}
  reading:
    - counting the VM process, containers cost about 3x pgmem's host memory at idle, before the ~0.9 GB VM baseline
    - devbox's native processes are the smallest at idle; pgmem wins start, per-test copy and download
    - native devbox PostgreSQL matches pgmem on TCP latency; pgmem's in-process dialer is ~3x faster
    - container latency includes OrbStack port forwarding
    - defaults differ: containers and devbox use shared_buffers=128MB and a disk; pgmem 32MB and memory
  history: an earlier run reported container memory from docker stats only (24 MB); replaced 2026-09-13 because it hid the VM
```
