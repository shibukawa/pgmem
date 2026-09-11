---
id: api:clone
type: api
title: Clone API
---
Creates an isolated fork of an api:snapshot for one test case and returns a ready database handle.

```yaml
api:
  shapes_under_consideration: decision:clone-release-style
  handle_types: decision:clone-return-type
  candidate_go_signatures:
    low_level: func (s *Snapshot) Fork(ctx context.Context) (*Fork, error)  # Fork has DSN(), DB(), PgxConn(), Close()
    callback: func (s *Snapshot) Clone(ctx context.Context, fn func(ctx context.Context, db *sql.DB)) error
    testing_tb: func (s *Snapshot) CloneT(t testing.TB) *sql.DB  # registers t.Cleanup
  blocking: waits for a free slot when policy:fork-pool-limit is reached; ctx cancel or t deadline aborts the wait
  cost: metric:fork-cost
  guarantees:
    - fork state equals snapshot state at creation
    - closing a fork never affects the snapshot or other forks
    - close is idempotent
```
