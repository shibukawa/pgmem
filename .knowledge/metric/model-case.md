---
id: metric:model-case
type: metric
title: Model-Case Suite Timing
---
One ordinary test suite run on pgmem and on a shared PostgreSQL server, every node of each execution model timed; the headline comparison on the top page (ui:website) and the first section of the benchmarks page.

```yaml
metric:
  project: bench/modelcase (store package: 5 migrations, master data 20 categories / 500 products / 1000 users, 40 sample orders, 5 read scenarios x 10 cases = 50 read tests + 5 write scenarios x 4 cases = 20 write tests; user asked for 10x/4x on 2026-09-18)
  models:
    pgmem: boot in process -> import -> snapshot -> read tests share the template server itself (sample data loaded into it after the snapshot; no fork, user 2026-09-18), each write test forks its own copy; all tests t.Parallel() (api:snapshot, decision:fork-or-not)
    shared_server: driver boots the server -> import into the shared database -> read tests parallel on it, write tests one at a time after reset (TRUNCATE + restore stock/prices) and sample load
  measured_2026_09_18:  # Apple M3, 8 cores, go test -parallel 8, median-total run of 5 after one warm-up
    pgmem_total_ms: 312   # boot 42, import 46, snapshot 9, 20 forks ~35 ms each (8 in flight), sample ~15 each, tests 0.3-8; range 296-435
    devbox_native_postgres_total_ms: 634   # boot 194, 20 x (reset ~5 + sample ~11) sequential, import 53, reset ~6 + sample ~10-20 per write test, sequential
    docker_run_total_ms: 2674   # boot 1647 (disk-backed, fresh container), rest as devbox but slower over the VM
  observations:
    - a fork alone is ~15 ms; with eight in flight they are 30-40 ms each (vfs clone + backend start contend) but overlap
    - read tests on one shared fork serialize at transaction boundaries (rule:single-session-per-backend); they are short either way
    - the first execution of a freshly built test binary is several hundred ms slower (page-in, code signing): driver warms up once, untimed
  chart: tests drawn in go test -parallel slots (first free slot), not one row per test, so 70 tests fit in 8 rows
  outputs: bench/modelcase/results/{raw.jsonl,summary.json}; website/src/data/modelcase.json rendered by ModelCaseChart.astro (hero variant on index, full with per-node table on benchmarks)
```
