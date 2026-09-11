---
id: decision:clone-release-style
type: decision
title: Clone Release Style
---
How a fork created by api:clone is released: callback scope, explicit Close with defer, or testing.TB cleanup.

```yaml
decision:
  resolution: pending
  options:
    callback:
      shape: Clone(ctx, func(ctx, db *sql.DB))
      pros: [release guaranteed on return and on panic, no leaked slot in policy:fork-pool-limit]
      cons: [test body nested in closure, awkward with t.Run subtests and table tests, not usable outside tests]
    defer_close:
      shape: f, err := Fork(ctx); defer f.Close()
      pros: [simplest primitive, works in non-test code, other helpers build on it]
      cons: [forgotten Close leaks a slot and blocks other tests forever under the pool limit]
    testing_tb:
      shape: db := CloneT(t); cleanup via t.Cleanup
      pros: [idiomatic Go test helper, release guaranteed at test end, t.Fatal on failure, parallel-safe]
      cons: [Go-test only, needs a test-only package to avoid importing testing in the library]
  recommendation:
    - ship defer_close as the primitive
    - ship testing_tb in a pgmemtest subpackage as the primary test API
    - add callback only if a non-test caller needs scoped forks
```
