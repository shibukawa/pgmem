package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_date_eq(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	return base.I64_extend_i32_u(base.B2i32(v2 == v3))
}
func F_date_le_timestamptz(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int64
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int64
	_ = v15
	var v17 int64
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_date_le_timestamptz[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v12
	v15 = *(*int64)(unsafe.Add(mBase, _c_F_date_le_timestamptz[1]))
	*(*int64)(unsafe.Add(mBase, uint32(v7))) = v15
	v17 = F_date2timestamptz_safe(m, v10, v7)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		return int64(0)
	} else {
		v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+4)))
		if v21 != int32(1) {
			v41 = base.B2i32(v9 < v17) - base.B2i32(v17 < v9)
		} else {
			if v17 != int64(-9223372036854775807-1) {
				if v17 != int64(9223372036854775807) {
					v41 = base.B2i32(v9 < v17) - base.B2i32(v17 < v9)
				} else {
					if v9 == int64(9223372036854775807) {
						v32 = int32(-1)
					} else {
						v32 = int32(1)
					}
					v41 = v32
				}
			} else {
				if v9 == int64(-9223372036854775807-1) {
					v37 = int32(1)
				} else {
					v37 = int32(-1)
				}
				v41 = v37
			}
		}
		m.G0 = v7 + int32(16)
		return base.I64_extend_i32_u(base.B2i32(v41 <= int32(0)))
	}
}
func F_date_lt(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	return base.I64_extend_i32_u(base.B2i32(v2 < v3))
}
