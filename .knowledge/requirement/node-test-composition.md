---
id: requirement:node-test-composition
type: requirement
title: Composable Node Test Integration
---
Node.js test integration must coexist with future Valkey and OpenSearch test services in the same runner and test file.

```yaml
requirements:
  ownership: each service owns only its server, file resource, environment variables and cleanup
  runner: user owns globalSetup and test command; multiple service setup entries can run together
  node_test: avoid replacing or monkey-patching the shared node:test export; any test wrapper must compose in either order with another service wrapper
  hooks: register cleanup/reset per service without replacing another service's hooks
  environment: inject each service endpoint before application imports; env names are configurable and collisions are explicit
  failure: partial setup failure releases already acquired resources; teardown remains reliable
  concurrency: document the boundary between sequential shared-file tests and concurrent cases requiring separate resources
  api: expose a small service-specific setup or fixture primitive usable by an application-owned composite helper
acceptance:
  - pgmem plus a second mock service can prepare, run and clean up in one node:test suite
  - reversing wrapper/setup order does not change isolation or cleanup
  - no service requires sole ownership of globalSetup, --import, or the test function
```

Related: decision:node-test-integration, api:node-wrapper, system:node-test-runners.
