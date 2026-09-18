package pgaot

import (
	"bytes"
	"compress/gzip"
	"io"
)

// wasm2go emits data.bin as raw bytes. gen-aot.sh replaces it with a
// deterministic gzip stream after generation, while keeping the generated
// embed declaration unchanged. Supporting raw bytes here as well makes
// the package safe to compile before the post-generation compression step.
func init() {
	if len(wasm2goData_data_bin) < 2 ||
		wasm2goData_data_bin[0] != 0x1f || wasm2goData_data_bin[1] != 0x8b {
		return
	}

	r, err := gzip.NewReader(bytes.NewReader(wasm2goData_data_bin))
	if err != nil {
		panic("pgaot: open compressed data.bin: " + err.Error())
	}
	data, err := io.ReadAll(r)
	closeErr := r.Close()
	if err != nil {
		panic("pgaot: decompress data.bin: " + err.Error())
	}
	if closeErr != nil {
		panic("pgaot: close compressed data.bin: " + closeErr.Error())
	}
	wasm2goData_data_bin = data
}
