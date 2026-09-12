---
id: ui:website
type: ui
title: pgmem Website
---
Astro documentation site for pgmem; every page is generated from one primary concept plus its supporting concepts, never hand-written facts.

```yaml
ui:
  root:
    kind: site
    id: site.pgmem
    title: pgmem
    generator: Astro (content collections; one markdown entry per page, produced by export --profile review)
    languages: [ja, en]  # source concepts stay English; prose per language is generated
    children:
      - kind: page
        id: home
        target: vision:pgmem
        children: [metric:server-footprint, api:go-server]
      - kind: page
        id: why-pgmem
        target: concept:alternatives
        children: [metric:alternative-comparison, metric:server-footprint, metric:fork-cost]
      - kind: page
        id: architecture
        target: concept:architecture
        children: [concept:vfs-snapshot, concept:server-process, rule:single-session-per-backend, decision:data-plane-transport]
      - kind: page
        id: extensions
        target: concept:static-modules
        children: [flow:add-extension, requirement:cloud-common-extensions]
      - kind: page
        id: limits
        target: concept:limits
      - kind: page
        id: releases
        target: policy:versioning
        children: [policy:binary-distribution]
      - kind: section
        id: guides
        children:
          - {kind: page, id: guide-go, target: concept:go-guide, children: [api:go-server, api:clone, api:in-process-dialer, flow:test-lifecycle]}
          - {kind: page, id: guide-python, target: concept:python-guide, children: [api:python-wrapper, flow:wrapper-test-lifecycle]}
          - {kind: page, id: guide-java, target: concept:java-guide, children: [api:java-wrapper, flow:wrapper-test-lifecycle]}
          - {kind: page, id: guide-nodejs, target: requirement:nodejs-wrapper, state: planned}
          - {kind: page, id: guide-other-languages, target: api:control-protocol, children: [policy:binary-distribution]}
      - kind: section
        id: testing-patterns
        children:
          - {kind: page, id: fork-or-not, target: decision:fork-or-not}
          - {kind: page, id: migrations, target: concept:migration-tools}
          - {kind: page, id: seeding, target: concept:seeding}
  guide_page_shape:  # every language guide answers these in this order
    - install
    - basic_test          # start a server, connect, one assertion
    - initialization      # raw SQL migration and seed, then popular migration tools
    - test_case_styles    # decision:fork-or-not applied to the language's test framework
    - parallel_and_pools  # rule:single-session-per-backend, policy:fork-pool-limit
  build: export each target with --profile review --scope 1, then commit the generated markdown under the Astro project, not under .knowledge
```
