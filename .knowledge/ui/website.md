---
id: ui:website
type: ui
title: pgmem Website
---
Astro Starlight site in website/, English at the root locale and Japanese under /ja with identical slugs; pages are written from the concepts below, and every benchmark number is rendered from one JSON file.

```yaml
ui:
  root:
    kind: site
    id: site.pgmem
    title: pgmem
    generator: Astro 7 + Starlight 0.42; base /pgmem; locales root=en, ja
    children:
      - kind: page
        id: index
        target: vision:pgmem
        children: [metric:model-case, metric:alternative-comparison]
        content: splash hero, ModelCaseChart (hero variant; totals and the pgmem vs native-server timelines) first for impact (user, 2026-09-18), feature cards, five-line start per language in synced tabs, BenchmarkOverview comparison cards last
      - {kind: page, id: why, target: concept:alternatives, children: [metric:alternative-comparison]}
      - {kind: page, id: benchmarks, target: metric:alternative-comparison, children: [metric:model-case, metric:server-footprint, metric:fork-cost], content: model-case section (ModelCaseChart full variant with per-node table) first, then the comparison table, memory chart and methodology}
      - {kind: page, id: architecture, target: concept:architecture, children: [concept:vfs-snapshot, concept:server-process]}
      - {kind: page, id: extensions, target: policy:bundled-extensions, children: [concept:static-modules, decision:pgcrypto-on-host]}
      - {kind: page, id: limits, target: concept:limits}
      - {kind: page, id: versioning, target: policy:versioning}
      - kind: section
        id: guides
        children:
          - {kind: page, id: guides/go/basics, target: concept:go-guide, children: [api:go-server, api:in-process-dialer, concept:migration-tools, concept:seeding]}
          - {kind: page, id: guides/go/testing, target: decision:fork-or-not, children: [api:clone, flow:test-lifecycle]}
          - {kind: page, id: guides/python/basics, target: concept:python-guide, children: [api:python-wrapper, concept:migration-tools, concept:seeding]}
          - {kind: page, id: guides/python/testing, target: decision:fork-or-not, children: [api:python-wrapper, flow:wrapper-test-lifecycle]}
          - {kind: page, id: guides/java/basics, target: concept:java-guide, children: [api:java-wrapper, concept:migration-tools, concept:seeding]}
          - {kind: page, id: guides/java/testing, target: decision:fork-or-not, children: [api:java-wrapper, flow:wrapper-test-lifecycle]}
          - {kind: page, id: guides/nodejs/basics, target: api:node-wrapper, children: [system:node-orms, concept:migration-tools]}
          - {kind: page, id: guides/nodejs/testing, target: decision:node-test-integration, children: [api:node-wrapper, api:reset, api:control-socket, system:node-test-runners]}
  guide_story:  # every language, in this order
    basics: [install with the official client, start, register schema (raw SQL then popular migration tools), seed data (raw SQL, COPY, dbtestify / DbUnit / factory_boy / ORM seeders), connect with official drivers, snapshot and fork by hand]
    testing: [fresh server per test, prepare once and share with read-only tests (class or file scope), fork or reset per test for writing tests, several templates, parallelism]
  tabs:
    syncKey: lang (Go, Python, Java, Node.js on the top page)
    jvm-build: [Gradle (Kotlin DSL), Gradle (Groovy DSL), Maven]
    jvm-lang: [Java, Kotlin]
    node-pm: [npm, pnpm, Yarn, Bun]
    per_guide: driver, migrate and seed tab groups share a syncKey per language
  numbers: website/src/data/benchmarks.json, copied by bench/alternatives/run.sh; BenchmarkTable.astro renders compact (top) and full (why, benchmarks) variants in both languages; website/src/data/modelcase.json, copied by bench/modelcase/run.sh, rendered by ModelCaseChart.astro
  deploy: .github/workflows/docs.yml builds website/ with withastro/action on pull requests and deploys to GitHub Pages (environment github-pages) on pushes to main that touch website/
  node_state: follows packages/node/core (index.d.ts, README) as merged 2026-09-13; install snippets assume the npm release made by flow:release
  versions_in_snippets: 0.1.0 while the 0.1.x releases bring flow:release up (user, 2026-09-14); switch to 1.18.0 when it ships (policy:versioning)
```
