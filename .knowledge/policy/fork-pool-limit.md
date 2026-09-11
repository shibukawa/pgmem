---
id: policy:fork-pool-limit
type: policy
title: Fork Pool Limit
---
Concurrent live forks per api:snapshot are capped; api:clone blocks until a slot frees instead of allocating more backends.

```yaml
policy:
  reason: each fork holds a vfs copy plus a wasm instance; 100 parallel forks would exhaust memory (metric:fork-cost)
  limit_source: SnapshotOption MaxForks; default GOMAXPROCS to match go test -parallel default
  wait: blocking acquire on a semaphore; honors ctx cancellation and testing deadlines
  deadlock_risk: a test holding two forks at once under a full pool waits forever; document and detect with a timeout warning
  prewarm: optional N forks started in advance to hide startup latency
  release: slot freed on Fork.Close; see decision:clone-release-style
```
