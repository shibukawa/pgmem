package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_cword_is_not_variable(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v35 int32
	_ = v35
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	F_errstart_cold(m, int32(21), int32(_a_F_cword_is_not_variable_0))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return
	} else {
		F_errcode(m, int32(16801924))
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v20 = F_NameListToString(m, v19)
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v8))) = v20
				F_errmsg(m, int32(_a_F_cword_is_not_variable_1), v8)
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return
				} else {
					v27 = F_plpgsql_scanner_errposition(m, l1, l2)
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_cword_is_not_variable_2), int32(2646), int32(_a_F_cword_is_not_variable_3))
						mBase = m.M
						v35 = m.ExcPending
						if v35 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		}
	}
}
