---
id: api:snapshot
type: api
title: Snapshot API
---
Freezes the current state of a concept:template-server into an immutable image that api:clone forks from.

```yaml
api:
  go: func (s *Server) Snapshot(ctx context.Context, opts ...SnapshotOption) (*Snapshot, error)
  behavior:
    - runs CHECKPOINT on the template backend
    - deep-copies the vfs (concept:vfs-snapshot)
    - allocates the fork pool for this snapshot (policy:fork-pool-limit)
  methods:
    - Clone (api:clone)
    - Close: waits for or rejects outstanding forks, frees image
    - Export(w io.Writer): optional tar.gz for cross-process reuse (requirement:in-memory-only allows explicit export)
  options:
    - MaxForks int
    - Prewarm int (forks started ahead of demand)
```
