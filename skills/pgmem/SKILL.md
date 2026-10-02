---
name: pgmem
description: Set up and troubleshoot pgmem-backed unit, HTTP API, and browser E2E tests in Go, Python, Node.js, or Java, including shadow routing, prepared snapshots, framework integration, and safe parallel runs.
---

# pgmem test integration

Choose the language and test boundary, then read only its reference. For a named framework, follow the link in that reference. Use the project's existing driver, migration tool, test runner, and app startup convention. Prepare schema and seed data once, isolate writes with a fork or reset, and close transactions before snapshot or reset. Check the installed pgmem version against its API; these references describe the repository's current test APIs.

| Language | Unit | HTTP API | Browser E2E |
|---|---|---|---|
| Go | [unit](references/go-unit.md) | [API](references/go-api.md) | [E2E](references/go-e2e.md) |
| Python | [unit](references/python-unit.md) | [API](references/python-api.md) | [E2E](references/python-e2e.md) |
| Node.js | [unit](references/nodejs-unit.md) | [API](references/nodejs-api.md) | [E2E](references/nodejs-e2e.md) |
| Java | [unit](references/java-unit.md) | [API](references/java-api.md) | [E2E](references/java-e2e.md) |

Use `website/src/content/docs/guides/<language>/` for full runnable examples. The Node.js directory is `nodejs`. Verify integration against the repository's package tests and the user's application tests.
