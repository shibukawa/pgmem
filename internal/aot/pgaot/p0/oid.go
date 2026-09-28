package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_OidFunctionCall6Coll(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64) int64 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v40 int32
	_ = v40
	var v51 int32
	_ = v51
	var v52 int64
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	v9 = m.G0
	v11 = v9 - int32(160)
	m.G0 = v11
	v14 = v11 + int32(12)
	v16 = *(*int32)(unsafe.Add(mBase, _c_F_OidFunctionCall6Coll[0]))
	F_fmgr_info_cxt_security(m, l0, v14, v16, int32(0))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		return int64(0)
	} else {
		v22 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v11)+152)) = uint8(v22)
		*(*int64)(unsafe.Add(mBase, uint32(v11)+144)) = l6
		*(*uint8)(unsafe.Add(mBase, uint32(v11)+136)) = uint8(v22)
		*(*int64)(unsafe.Add(mBase, uint32(v11)+128)) = l5
		*(*uint8)(unsafe.Add(mBase, uint32(v11)+120)) = uint8(v22)
		*(*int64)(unsafe.Add(mBase, uint32(v11)+112)) = l4
		*(*uint8)(unsafe.Add(mBase, uint32(v11)+104)) = uint8(v22)
		*(*int64)(unsafe.Add(mBase, uint32(v11)+96)) = l3
		*(*uint8)(unsafe.Add(mBase, uint32(v11)+88)) = uint8(v22)
		*(*int64)(unsafe.Add(mBase, uint32(v11)+80)) = l2
		*(*uint8)(unsafe.Add(mBase, uint32(v11)+72)) = uint8(v22)
		*(*int64)(unsafe.Add(mBase, uint32(v11)+64)) = l1
		v40 = int32(6)
		*(*uint16)(unsafe.Add(mBase, uint32(v11)+58)) = uint16(v40)
		*(*uint8)(unsafe.Add(mBase, uint32(v11)+56)) = uint8(v22)
		*(*int32)(unsafe.Add(mBase, uint32(v11)+52)) = v22
		*(*int64)(unsafe.Add(mBase, uint32(v11)+44)) = int64(0)
		*(*int32)(unsafe.Add(mBase, uint32(v11)+40)) = v14
		v51 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
		v52 = m.T0[v51].(func(*base.Module, int32) int64)(m, v11+int32(40))
		mBase = m.M
		v53 = m.ExcPending
		if v53 != 0 {
			return int64(0)
		} else {
			v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+56)))
			if v54 == int32(1) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v60 = m.ExcPending
				if v60 != 0 {
					return int64(0)
				} else {
					v61 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
					*(*int32)(unsafe.Add(mBase, uint32(v11))) = v61
					F_errmsg_internal(m, int32(_a_F_OidFunctionCall6Coll_0), v11)
					mBase = m.M
					v65 = m.ExcPending
					if v65 != 0 {
						return int64(0)
					} else {
						F_errfinish(m, int32(_a_F_OidFunctionCall6Coll_1), int32(1280), int32(_a_F_OidFunctionCall6Coll_2))
						mBase = m.M
						v70 = m.ExcPending
						if v70 != 0 {
							return int64(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				m.G0 = v11 + int32(160)
				return v52
			}
		}
	}
}
func F_oid_increment(m *base.Module, l0 int32, l1 int64, l2 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v13 int64
	_ = v13
	v6 = base.B2i32(base.I32_wrap_i64(l1) == int32(-1))
	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v6)
	if base.I32_wrap_i64(l1) == int32(-1) {
		v13 = int64(0)
	} else {
		v13 = (l1 + int64(1)) & int64(4294967295)
	}
	return v13
}
