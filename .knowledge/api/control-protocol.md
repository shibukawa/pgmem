---
id: api:control-protocol
type: api
title: Control Protocol
---
Newline-delimited JSON over the concept:server-process stdin/stdout that lets a wrapper create snapshots and forks and close them (decision:control-plane-transport).

```yaml
api:
  status: implemented 2026-09-12 in cmd/pgmem/control.go; tests in control_test.go; spec in docs/subprocess.md
  framing: one JSON object per line, UTF-8, no pretty printing; requests on stdin, responses and events on stdout
  version: ready event carries an integer protocol; the wrapper refuses a mismatch
  ready_event: '{"event":"ready","protocol":1,"version":"<binary version>","pid":123,"server":<endpoint>}'
  server_ids: template (default, from CLI flags), t2.. (start op), f1.. (fork op)
  endpoint: '{"id":"template","host":"127.0.0.1","port":54321,"user":"postgres","database":"app","dsn":"postgres://..."}'
  request: '{"id":<client int>,"op":"<name>",...}'
  response: '{"id":<same>,"ok":true,...}' or '{"id":<same>,"ok":false,"error":{"code":"<snake_case>","message":"..."}}'
  ops:
    snapshot: 'in {"server":"template","max_forks":N?,"timeout_ms":N?}  out {"snapshot":"s1"}; api:snapshot; waits for open transactions, busy after timeout'
    fork: 'in {"snapshot":"s1","timeout_ms":N?}  out {"server":<endpoint with id "f1">}; api:clone, blocks until a slot frees'
    close: 'in {"server":"f1"} or {"snapshot":"s1"}  out {}; idempotent, unknown ids succeed'
    shutdown: 'in {}  out {} then the process exits 0'
    start: 'in {"database":..,"user":..,"params":["k=v",..]}  out {"server":<endpoint with id "t2">}; extra template server for suites with several seed sets; closable like a fork'
  ordering: rule:non-blocking-control-channel
  error_codes: [unknown_op, unknown_id, snapshot_closed, pool_timeout, busy, protocol, internal]
  busy: a client connection idle in a transaction blocks snapshot forever (rule:single-session-per-backend); wrappers default timeout_ms to 30000 and tell the user to commit or close connections first
  malformed_line: response with id null and code protocol; the process keeps running
  events_without_id: ready, fatal (written before an abnormal exit)
  legacy: the ready line stays the first stdout line so pre-protocol wrappers keep working (concept:server-process)
```
