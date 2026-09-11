---
id: metric:fork-cost
type: metric
title: Fork Cost
---
Time and memory of creating one fork; drives defaults in policy:fork-pool-limit.

```yaml
metric:
  vfs_data_dir_bytes: 40000000  # measured, 16 MB of it is one WAL segment
  snapshot_time_ms: 10  # measured: CHECKPOINT + vfs clone
  fork_time_ms: 20  # measured: vfs clone + backend start + listen
  wasm_memory_per_backend: unmeasured; expected dominant term; measure before fixing default MaxForks
  targets:
    fork_latency_s: 0.3
    forks_in_512mb: to be determined
```
