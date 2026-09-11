package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_get_func_signature(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v12 = F_SearchSysCache1(m, int32(47), l0)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return
	} else {
		if v12 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
				F_errmsg_internal(m, int32(_a_F_get_func_signature_0), v9)
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_get_func_signature_1), int32(1844), int32(_a_F_get_func_signature_2))
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v29 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
			v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+22)))
			v31 = v29 + v30
			v33 = int32(*(*int16)(unsafe.Add(mBase, uint32(v31)+104)))
			*(*int32)(unsafe.Add(mBase, uint32(l2))) = v33
			v37 = F_palloc(m, v33<<(uint(int32(2))%32))
			mBase = m.M
			v38 = m.ExcPending
			if v38 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l1))) = v37
				v42 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
				v44 = v42 << (uint(int32(2)) % 32)
				if v44 != 0 {
					v45 = F__emscripten_memcpy_bulkmem(m, v37, v31+int32(136), v44)
					mBase = m.M
				} else {
				}
				F_ReleaseCatCache(m, v12)
				mBase = m.M
				v48 = m.ExcPending
				if v48 != 0 {
					return
				} else {
					m.G0 = v9 + int32(16)
					return
				}
			}
		}
	}
}
