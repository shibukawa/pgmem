package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_y_cmp(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 float64
	_ = v8
	var v9 int32
	_ = v9
	var v10 float64
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v8 = *(*float64)(unsafe.Add(mBase, uint32(v7)+8))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v10 = *(*float64)(unsafe.Add(mBase, uint32(v9)+8))
	if base.F64_gt(v8, v10) != 0 {
		v12 = int32(1)
	} else {
		v12 = int32(-1)
	}
	if base.F64_ne(v8, v10) != 0 {
		v15 = v12
	} else {
		v15 = int32(0)
	}
	return v15
}
func F_yy_fatal_error_3(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v5))) = l0
		F_errmsg_internal(m, int32(195849), v5)
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return
		} else {
			F_errfinish(m, int32(299138), int32(38), int32(76551))
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return
			} else {
				base.Wasm_trap_unreachable()
				for {
				}
			}
		}
	}
}
