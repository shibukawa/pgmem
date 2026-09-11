---
id: concept:template-server
type: concept
title: Template Server
---
The server started in TestMain whose state after migrations and seeding becomes the origin of every fork.

```yaml
summary:
  lifetime: whole test binary
  after_snapshot: may stay running for suite-level checks, or be closed to save memory
  writes_after_snapshot: do not affect existing snapshot (concept:vfs-snapshot copies)
```
