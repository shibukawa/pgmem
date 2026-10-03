---
id: requirement:go-external-test-backend
type: requirement
title: Go External Process Test Backend
---
Go tests may launch pgmem as an external process and connect through TCP or Unix domain sockets, alongside embedded operation. User requirement added and implemented 2026-10-03 in api:go-process and api:go-test-fixture.

```yaml
requirements:
  execution: select embedded or an owned subprocess independently of the SQL transport
  connections:
    embedded: existing in-process Dial or loopback TCP (api:go-server)
    subprocess: loopback TCP or explicit Unix domain socket mode
  test_interface: prepare once; snapshot; isolated test database; shared reads; reset; automatic cleanup with equivalent behavior across supported backends
  test_identity: a test database object supplies DB, pgx pool and DSN for the same fork; separate databases require explicit acquisition
  subprocess_client: do not import or link the embedded PostgreSQL engine; binary path can be configured; binary discovery must be documented
  process_control: use api:control-protocol on stdin/stdout; SQL sockets do not replace the control channel
  lifecycle: startup deadline, context cancellation, failed prepare cleanup, child exit detection and bounded shutdown
  uds:
    layout: PostgreSQL-compatible socket naming; independent endpoint per template and fork
    dsn: expose a driver-compatible connection string, including socket directory and port identifier
    stability: reset retains endpoint and existing connection behavior (api:reset)
    cleanup: private runtime directory; close removes only owned sockets; normal shutdown removes owned directory
    socket_type: AF_UNIX SOCK_STREAM, including Windows support; no SOCK_DGRAM or SOCK_SEQPACKET requirement
    temporary_parent: /tmp on Unix, os.TempDir() on Windows; explicit SocketDir overrides
    support: validate OS support and socket path length; unsupported explicit selection fails clearly instead of silently switching to TCP
  compatibility: retain existing embedded helper entry points; external connection helpers use network dialing instead of Server.Dial
acceptance:
  - SQL, fork isolation, reset and teardown verified for an owned TCP subprocess
  - same lifecycle verified for Unix sockets on supported platforms
  - template and concurrent forks receive distinct endpoints
  - external-client dependency graph excludes the embedded engine and generated PostgreSQL packages
  - subprocess, forks and owned socket paths are released on normal teardown and setup failure
```

Related: requirement:test-fixture-fork, decision:data-plane-transport, requirement:in-memory-only.
