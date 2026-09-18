---
id: metric:fork-cost
type: metric
title: Fork Cost
---
Time and memory of creating one fork; drives defaults in policy:fork-pool-limit.

```yaml
metric:
  vfs_data_dir_bytes: 50000000  # measured with a small schema, 16 MB of it is one WAL segment; 2,100 nodes
  snapshot_time_ms: 10  # measured: CHECKPOINT + vfs clone
  fork_time_ms: 6  # measured 2026-09-19 (M3): vfs clone 0.15 ms (copy-on-write, was 5-9 ms) + backend start 6 ms (was 7-9 ms before the anonymous-mmap change) + listen 0.06 ms
  fork_time_ms_8_concurrent: 31  # per fork with 8 at once, wall 40 ms for all 8 (was 53 per fork, wall 86 ms): page faults on fresh memory serialize in the kernel
  remaining_cost_2026_09_19: bytes.Clone of the files startup writes (the WAL segment, ~1 ms per fork; chunked COW would remove it), guest memcpy and data-segment init, PostgreSQL's own startup (StartupXLOG, relcache)
  wasm_memory_per_backend: unmeasured; expected dominant term; measure before fixing default MaxForks
  targets:
    fork_latency_s: 0.3
    forks_in_512mb: to be determined
```
