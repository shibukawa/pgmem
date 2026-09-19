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
    pgmem_total_ms: 209   # 2026-09-19 after COW vfs + anonymous-mmap change: boot 39, import 48, snapshot 3.5, 20 forks ~11 ms each (8 in flight), sample ~17 each; range 206-215 (was 312 on 2026-09-18)
    testcontainers_total_ms: 1726   # 2026-09-19, testcontainers-go postgres module, in-memory mode (PGDATA tmpfs, fsync/synchronous_commit/full_page_writes off), one driver process per run so Ryuk is in the boot: boot 761, 20 x (reset ~7 + sample ~30) sequential; the top page compares pgmem with this one (user, 2026-09-19)
    devbox_native_postgres_total_ms: 658   # boot 200, 20 x (reset ~6 + sample ~12) sequential, import 53, reset ~6 + sample ~10-20 per write test, sequential
    docker_run_total_ms: 2471   # boot 1601 (disk-backed, fresh container), rest as devbox but slower over the VM
  observations:
    - a fork alone is ~6 ms (was 15); with eight in flight ~10 ms each (was 30-40): see metric:fork-cost (vfs clone + backend start contend) but overlap
    - read tests on one shared fork serialize at transaction boundaries (rule:single-session-per-backend); they are short either way
    - the first execution of a freshly built test binary is several hundred ms slower (page-in, code signing): driver warms up once, untimed
  chart: tests drawn in go test -parallel slots (first free slot), not one row per test, so 70 tests fit in 8 rows; every timeline shares one time axis (user, 2026-09-19); hero variant = pgmem vs Testcontainers, full = pgmem, Testcontainers, devbox, docker run
  outputs: bench/modelcase/results/{raw.jsonl,summary.json}; website/src/data/modelcase.json rendered by ModelCaseChart.astro (hero variant on index, full with per-node table on benchmarks)
```
