package p2

import base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"

func F_emscripten_stack_get_current(m *base.Module) int32 {
	var v1 int32
	_ = v1
	v1 = m.G0
	return v1
}
