package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_checkDomainOwner(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+22)))
	v9 = v7 + v8
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+79)))
	if v10 == int32(100) {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
		v16 = *(*int32)(unsafe.Add(mBase, _consts[31]))
		v17 = F_object_ownercheck(m, int32(1247), v14, v16)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return
		} else {
			if v17 == int32(0) {
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
				F_aclcheck_error_type(m, int32(2), v22)
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return
				} else {
					m.G0 = v5 + int32(16)
					return
				}
			} else {
				m.G0 = v5 + int32(16)
				return
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v31 = m.ExcPending
		if v31 != 0 {
			return
		} else {
			F_errcode(m, int32(151027844))
			mBase = m.M
			v34 = m.ExcPending
			if v34 != 0 {
				return
			} else {
				v35 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
				v36 = F_format_type_be(m, v35)
				mBase = m.M
				v37 = m.ExcPending
				if v37 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v5))) = v36
					F_errmsg(m, int32(273275), v5)
					mBase = m.M
					v41 = m.ExcPending
					if v41 != 0 {
						return
					} else {
						F_errfinish(m, int32(485608), int32(3494), int32(215187))
						mBase = m.M
						v46 = m.ExcPending
						if v46 != 0 {
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
