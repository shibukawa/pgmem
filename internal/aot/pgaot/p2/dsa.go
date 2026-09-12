package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_dsa_attach_in_place(m *base.Module, l0 int32, l1 int32) int32 {
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	v5 = F_attach_internal(m, l0, int32(0))
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		if l1 != 0 {
			F_on_dsm_detach(m, l1, int32(1785), l0)
			v11 = m.ExcPending
			if v11 != 0 {
				return int32(0)
			} else {
				return v5
			}
		} else {
			return v5
		}
	}
}
func F_dsa_pin(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v7 = F_LWLockAcquire(m, v3+int32(1476), int32(0))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+1464)))
		if v10 == int32(1) {
			F_LWLockRelease(m, v9+int32(1476))
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return
				} else {
					F_errmsg_internal(m, int32(432975), int32(0))
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return
					} else {
						F_errfinish(m, int32(478368), int32(981), int32(262966))
						mBase = m.M
						v29 = m.ExcPending
						if v29 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		} else {
			v30 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(v9)+1464)) = uint8(v30)
			v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+1460))
			*(*int32)(unsafe.Add(mBase, uint32(v32)+1460)) = v33 + v30
			v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			F_LWLockRelease(m, v37+int32(1476))
			mBase = m.M
			v41 = m.ExcPending
			if v41 != 0 {
				return
			} else {
				return
			}
		}
	}
}
