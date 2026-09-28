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
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v13 = F_SearchSysCache1(m, int32(47), base.I64_extend_i32_u(l0))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return
	} else {
		if v13 != 0 {
			v15 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
			v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+22)))
			v17 = v15 + v16
			v19 = int32(*(*int16)(unsafe.Add(mBase, uint32(v17)+104)))
			*(*int32)(unsafe.Add(mBase, uint32(l2))) = v19
			v23 = F_palloc(m, v19<<(uint(int32(2))%32))
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l1))) = v23
				v26 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
				v28 = v26 << (uint(int32(2)) % 32)
				if v28 != 0 {
					base.MemoryCopy(m, v23, v17+int32(136), v28)
				} else {
				}
				F_ReleaseCatCache(m, v13)
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return
				} else {
					m.G0 = v9 + int32(16)
					return
				}
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v40 = m.ExcPending
			if v40 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
				F_errmsg_internal(m, int32(_a_F_get_func_signature_0), v9)
				mBase = m.M
				v44 = m.ExcPending
				if v44 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_get_func_signature_1), int32(2016), int32(_a_F_get_func_signature_2))
					mBase = m.M
					v49 = m.ExcPending
					if v49 != 0 {
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
