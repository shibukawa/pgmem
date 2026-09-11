package p4

import base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"

func F_StartupPacketTimeoutHandler(m *base.Module) {
	F__Exit(m, int32(1))
	base.Wasm_trap_unreachable()
	for {
	}
}
