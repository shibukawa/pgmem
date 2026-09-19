package engine

import "github.com/shibukawa/pgmem/internal/aot"

// The root package registers the ahead-of-time backend; this package's
// tests have to do it themselves.
func init() {
	PostgresFactory = &aot.Factory{}
}
