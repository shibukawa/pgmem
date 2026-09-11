package p0

import base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"

func F__emscripten_stack_alloc(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	v2 = m.G0
	v5 = (v2 - l0) & int32(-16)
	m.G0 = v5
	return v5
}
