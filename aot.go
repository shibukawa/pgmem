package pgmem

import (
	"github.com/shibukawa/pgmem/internal/aot"
	"github.com/shibukawa/pgmem/internal/engine"
)

// The backend is Go code generated ahead of time by wasm2go from the
// PostgreSQL wasm module; no wasm runtime is involved at run time.
func init() {
	engine.PostgresFactory = &aot.Factory{}
}
