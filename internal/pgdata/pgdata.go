// Package pgdata embeds a data directory produced by initdb (see
// cmd/pgmem-mkdata) so servers can start without running initdb.
package pgdata

import _ "embed"

// TarZst is the initdb output (zstd-compressed tar), superuser
// "postgres", encoding UTF8, locale C.UTF-8, auth trust.
//
//go:embed pgdata.tar.zst
var TarZst []byte
