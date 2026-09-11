package main

import (
	"github.com/shibukawa/pgmem/internal/aot"
	"github.com/shibukawa/pgmem/internal/engine"
)

// initdb's postgres children run on the default (ahead-of-time) backend;
// initdb itself always runs under wazero.
func init() {
	engine.PostgresFactory = aot.Factory{}
}
