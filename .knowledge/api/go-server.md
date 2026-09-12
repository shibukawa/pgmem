---
id: api:go-server
type: api
title: Go Server API
---
The core Go entry point: start one in-memory PostgreSQL and get a DSN; everything else (api:snapshot, api:clone, pgmemtest) builds on it.

```yaml
api:
  module: github.com/shibukawa/pgmem; pure Go, no cgo, no download at run time
  start: 'func Start(ctx context.Context, opts Options) (*Server, error)'
  options:
    Database: served database name, created at startup by a standalone child (default postgres)
    User: initdb superuser name for the DSN (default postgres)
    Port: 0 picks a free loopback port; always 127.0.0.1
    Params: extra postgres -c arguments, e.g. shared_buffers=128MB or log_statement=all; shared_buffers defaults to 32MB, io_method forced to sync
    Log: 'func(format string, args ...any) receives server log and host diagnostics; nil discards'
  server:
    - 'DSN() string  # postgres://user@127.0.0.1:port/db?sslmode=disable; any driver'
    - 'Dial  # api:in-process-dialer'
    - 'Snapshot(ctx, SnapshotOptions) (*Snapshot, error)  # api:snapshot'
    - 'Close() error  # frees the listener, backend and linear memory at once'
  cost: metric:server-footprint
  isolation: every Server has its own vfs and database state; several per process share only the loaded engine (concept:server-process)
  debug: PGMEM_TRACE=1 prints every host call to stderr
```
