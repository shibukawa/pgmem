// Package assets embeds the PostgreSQL share directory (postgres.bki, sql
// scripts, timezone data). The server code itself is Go generated from the
// wasm module (internal/aot/pgaot); the wasm files are build inputs only.
package assets

import _ "embed"

// ShareTarGz is share/postgresql.
//
//go:embed share.tar.gz
var ShareTarGz []byte
