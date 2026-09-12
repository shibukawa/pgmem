---
id: api:clone
type: api
title: Clone API
---
Creates an isolated fork of an api:snapshot for one test case and returns a ready database handle.

```yaml
api:
  primitive: func (sn *Snapshot) Fork(ctx context.Context) (*Server, error)  # Close() frees the slot
  test_helpers: package pgmemtest (decision:clone-release-style)
    - Run(m *testing.M, opts Options, setup func(*Fixture)) int  # TestMain helper
    - New(ctx, Options) (*Fixture, error); Fixture.Close()
    - Fixture.Fork(t testing.TB) *pgmem.Server
    - Fixture.DB(t) *sql.DB; PgxConn(t) *pgx.Conn; PgxPool(t) *pgxpool.Pool; DSN(t) string  # decision:clone-return-type
    - Options.Prepare func(ctx, db *sql.DB, dsn string) error  # runs once on the template
  blocking: waits for a free slot when policy:fork-pool-limit is reached; t.Context() cancel aborts the wait
  cost: metric:fork-cost
  wrappers: exposed as op fork / close in api:control-protocol
  guarantees:
    - fork state equals snapshot state at creation
    - closing a fork never affects the snapshot or other forks
    - close is idempotent
```
