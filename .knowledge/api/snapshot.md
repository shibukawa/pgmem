---
id: api:snapshot
type: api
title: Snapshot API
---
Freezes the current state of a concept:template-server into an immutable image that api:clone forks from.

```yaml
api:
  go: func (s *Server) Snapshot(ctx context.Context, opts SnapshotOptions) (*Snapshot, error)
  file: snapshot.go
  behavior:
    - acquires the backend (waits for open transactions; ctx bounds the wait via acquireCtx), runs CHECKPOINT
    - deep-copies the vfs (concept:vfs-snapshot); template keeps running
    - allocates the fork slot pool (policy:fork-pool-limit)
  methods:
    - Fork(ctx) (*Server, error): api:clone
    - Close(): rejects new forks; live forks keep their own copies
    - Wait(): blocks until every fork is closed
  options:
    - MaxForks int (0 = memory limit / 4 / (shared_buffers + 32 MB); GOMAXPROCS when memory is unknown)
  measured: snapshot ~10ms, fork ~20ms (metric:fork-cost)
  wrappers: exposed as op snapshot in api:control-protocol
  later: Export(w io.Writer) tar for cross-process reuse; Prewarm
```
