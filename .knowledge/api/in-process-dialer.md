---
id: api:in-process-dialer
type: api
title: In-Process Dialer
---
Server.Dial connects a Go client to the backend through net.Pipe instead of loopback TCP, which removes the kernel socket path from small queries.

```yaml
api:
  signature: 'func (s *Server) Dial(ctx context.Context, network, addr string) (net.Conn, error)  # fits pgx ConnConfig.DialFunc'
  pgx: 'cfg, _ := pgx.ParseConfig(s.DSN()); cfg.DialFunc = s.Dial; pgx.ConnectConfig(ctx, cfg)'
  database_sql: 'stdlib.RegisterConnConfig(cfg) then sql.Open("pgx", name); pgmemtest Fixture.DB uses this'
  gain: simple SELECT 8.5 us versus 26.5 us over TCP (metric:server-footprint); a TCP query spends ~5% in PostgreSQL and the rest in syscalls, poller and goroutine wake-ups
  scope: Go only; wrappers in other languages use TCP (decision:data-plane-transport)
  same_semantics: the connection gets its own backend process like a TCP one (rule:process-per-connection)
```
