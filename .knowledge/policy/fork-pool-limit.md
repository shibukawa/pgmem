---
id: policy:fork-pool-limit
type: policy
title: Fork Pool Limit
---
Concurrent live forks per api:snapshot are capped; api:clone blocks until a slot frees instead of allocating more backends.

```yaml
policy:
  reason: each fork holds a vfs copy plus a wasm instance; 100 parallel forks would exhaust memory (metric:fork-cost)
  limit_source: SnapshotOption MaxForks; default = memory limit / 4 / (shared_buffers + 32 MB) (metric:fork-cost), GOMAXPROCS only when memory is unknown; a slot costs nothing until used, so the cap is a ceiling not a reservation
  memory_limit: min(GOMEMLIMIT, cgroup memory.max or v1 limit_in_bytes, physical memory); linux/darwin/windows
  why_not_cpu: measured 2026-09-19 with GOMAXPROCS=2, fork cap 2 vs 8: CPU-bound tests +7%, tests that wait 3 ms per query 3.9x faster; the cap protects memory, not CPU
  node_vitest_2vcpu: measured 2026-09-19 in docker --cpus=2 --memory=4g, node 22 (libuv 1.51, availableParallelism=2), vitest 5, 16 files x 10 tests: at the default maxWorkers=2 old and new caps tie (about 3.0 s with 20 ms waits per test); at maxWorkers=4 or 8 the old cap of 2 stays at 2.8-3.1 s while the new cap (16) reaches 1.7-1.9 s; without waits every setting is 1.2-1.5 s (CPU-bound). Raising the cap only helps when the runner's workers exceed the CPU count
  wait: blocking acquire on a semaphore; honors ctx cancellation and testing deadlines
  deadlock_risk: a test holding two forks at once under a full pool waits forever; document and detect with a timeout warning
  prewarm: optional N forks started in advance to hide startup latency
  release: slot freed on Fork.Close; see decision:clone-release-style
```
