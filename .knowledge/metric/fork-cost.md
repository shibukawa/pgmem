---
id: metric:fork-cost
type: metric
title: Fork Cost
---
Time and memory of creating one fork; drives defaults in policy:fork-pool-limit.

```yaml
metric:
  vfs_data_dir_bytes: 40000000  # measured, 16 MB of it is one WAL segment
  vfs_clone_time: a few ms (memcpy)
  backend_start_time_s: 0.1-0.2  # README numbers for additional servers
  wasm_memory_per_backend: unmeasured; expected dominant term; measure before fixing default MaxForks
  targets:
    fork_latency_s: 0.3
    forks_in_512mb: to be determined
```
