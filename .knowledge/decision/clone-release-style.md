---
id: decision:clone-release-style
type: decision
title: Clone Release Style
---
How a fork created by api:clone is released: callback scope, explicit Close with defer, or testing.TB cleanup.

```yaml
decision:
  resolution: defer_close primitive + testing_tb helpers in pgmemtest (2026-09-12)
  options:
    callback:
      shape: Clone(ctx, func(ctx, db *sql.DB))
      pros: [release guaranteed on return and on panic, no leaked slot in policy:fork-pool-limit]
      cons: [test body nested in closure, awkward with t.Run subtests and table tests, not usable outside tests]
    defer_close:
      shape: srv, err := snap.Fork(ctx); defer srv.Close()
      pros: [simplest primitive, works in non-test code, other helpers build on it]
      cons: [forgotten Close leaks a slot and blocks other tests forever under the pool limit]
    testing_tb:
      shape: db := fx.DB(t); cleanup via t.Cleanup
      pros: [idiomatic Go test helper, release guaranteed at test end, t.Fatal on failure, parallel-safe]
      cons: [Go-test only, lives in pgmemtest so the library does not import testing]
  rationale: matches Go standard library conventions; callback form deferred until a non-test caller needs scoped forks
```
