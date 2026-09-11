GO_PKGS := $(shell go list ./... | grep -v /internal/aot)

.PHONY: test test-aot bench vet wasm aot mkdata

test:        ## run the test suite
	go test ./... -count=1

bench:       ## query benchmarks (TCP, in-process, engine-bound)
	go test . -run xxx -bench 'SimpleQueries|CPUHeavy' -benchtime 3s

vet:         ## vet everything; the generated code only gets the cheap passes
	go vet $(GO_PKGS)
	go vet -unreachable=false ./internal/aot
	gofmt -l $(shell find . -name '*.go' -not -path './internal/aot/pgaot/*' -not -path './postgres-pglite/*' -not -path './toolchain/*')

wasm:        ## rebuild postgres.wasm / initdb.wasm / share.tar.gz (needs emsdk)
	./wasm/build.sh

aot:         ## regenerate internal/aot/pgaot with the forked wasm2go
	./wasm/gen-aot.sh

mkdata:      ## regenerate the embedded initdb data directory
	go run ./cmd/pgmem-mkdata
