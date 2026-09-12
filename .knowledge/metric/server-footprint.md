---
id: metric:server-footprint
type: metric
title: Server Footprint
---
Measured startup, latency, memory and size of one pgmem server (docs/benchmarks.md, 2026-09-11, Apple M-series arm64, Go 1.27).

```yaml
metric:
  start_first_in_process_s: 0.13
  start_further_in_process_s: 0.1
  simple_select_tcp_us: 26.5      # pgx QueryRow, 1000-row table, primary key
  simple_select_in_process_us: 8.5  # api:in-process-dialer
  sort_count_200k_rows_ms: 95
  rss_default_server_mb: 152      # shared_buffers=32MB; 65 MB linear memory, 39 MB vfs data dir, rest Go heap and binary
  rss_with_setup_child_mb: 161    # -database app -user tester
  rss_shared_buffers_128mb_mb: 230
  binary_stripped_mb: 36.2        # minimal program, -ldflags="-s -w"
  generated_go_source_mb: 104   # backend; 110 MB shipped in total with the aot glue
  compile_once_s: 35
  fork: metric:fork-cost
  unmeasured: [amd64, Linux, Windows]
```
