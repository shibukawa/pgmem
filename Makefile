GO_PKGS := $(shell go list ./... | grep -v /internal/aot)

.PHONY: test test-aot bench vet wasm aot mkdata python java

test:        ## run the test suite
	go test ./... -count=1

bench:       ## query benchmarks (TCP, in-process, engine-bound)
	go test . -run xxx -bench 'SimpleQueries|CPUHeavy' -benchtime 3s

vet:         ## vet everything; the generated code only gets the cheap passes
	go vet $(GO_PKGS)
	go vet -unreachable=false ./internal/aot
	@unformatted=$$(gofmt -l $$(git ls-files '*.go' | grep -v '^internal/aot/pgaot/')); \
	if [ -n "$$unformatted" ]; then echo "needs gofmt:"; echo "$$unformatted"; exit 1; fi

wasm:        ## rebuild postgres.wasm / initdb.wasm / share.tar.gz (needs emsdk)
	./wasm/build.sh

aot:         ## regenerate internal/aot/pgaot with the forked wasm2go
	./wasm/gen-aot.sh

mkdata:      ## regenerate the embedded initdb data directory
	go run ./cmd/pgmem-mkdata

python:      ## test and build the Python package (packages/python)
	cd packages/python && uv run pytest -q && uv build --wheel

java:        ## test and build the Java artifacts (packages/java)
	cd packages/java && ./gradlew build
