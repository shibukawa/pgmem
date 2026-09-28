package p5

import base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"

func F_ghstore_out(m *base.Module, l0 int32) int64 {
	var v7 int64
	_ = v7
	var v10 int32
	_ = v10
	v7 = Fn14235(m, l0, int32(_a_F_ghstore_out_0), int32(112), int32(_a_F_ghstore_out_1), int32(_a_F_ghstore_out_2), int32(_a_F_ghstore_out_3))
	v10 = m.ExcPending
	if v10 != 0 {
		return int64(0)
	} else {
		return v7
	}
}
func F_ghstore_same(m *base.Module, l0 int32) int64 {
	var v5 int64
	_ = v5
	var v8 int32
	_ = v8
	v5 = Fn14295(m, l0, int32(2), int32(4), int32(16))
	v8 = m.ExcPending
	if v8 != 0 {
		return int64(0)
	} else {
		return v5
	}
}
