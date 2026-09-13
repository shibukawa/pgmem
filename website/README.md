# pgmem website

Documentation site for pgmem, built with [Astro](https://astro.build) and [Starlight](https://starlight.astro.build). English pages live in `src/content/docs/`, Japanese translations in `src/content/docs/ja/` with the same file names.

```sh
npm install
npm run dev      # http://localhost:4321/pgmem/
npm run build    # static site in dist/
```

## Benchmark numbers

Every number in the comparison tables comes from `src/data/benchmarks.json`, rendered by `src/components/BenchmarkTable.astro`. Regenerate it with the harness:

```sh
cd ../bench/alternatives
RUNS=5 ./run.sh   # needs a Docker daemon and devbox
```

## Source of truth

Facts on these pages are organized in the `.knowledge` catalog at the repository root. Update the catalog when behaviour changes, then the pages.
