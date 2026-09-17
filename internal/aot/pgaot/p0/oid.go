package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_OidFunctionCall6Coll(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32 {
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
	var v52 int32
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
	v11 = v9 - int32(112)
	m.G0 = v11
	v14 = v11 + int32(16)
	v16 = *(*int32)(unsafe.Add(mBase, _c_F_OidFunctionCall6Coll[0]))
	F_fmgr_info_cxt_security(m, l0, v14, v16, int32(0))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		return int32(0)
	} else {
		v22 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v11)+108)) = uint8(v22)
		*(*int32)(unsafe.Add(mBase, uint32(v11)+104)) = l6
		*(*uint8)(unsafe.Add(mBase, uint32(v11)+100)) = uint8(v22)
		*(*int32)(unsafe.Add(mBase, uint32(v11)+96)) = l5
		*(*uint8)(unsafe.Add(mBase, uint32(v11)+92)) = uint8(v22)
		*(*int32)(unsafe.Add(mBase, uint32(v11)+88)) = l4
		*(*uint8)(unsafe.Add(mBase, uint32(v11)+84)) = uint8(v22)
		*(*int32)(unsafe.Add(mBase, uint32(v11)+80)) = l3
		*(*uint8)(unsafe.Add(mBase, uint32(v11)+76)) = uint8(v22)
		*(*int32)(unsafe.Add(mBase, uint32(v11)+72)) = l2
		*(*uint8)(unsafe.Add(mBase, uint32(v11)+68)) = uint8(v22)
		*(*int32)(unsafe.Add(mBase, uint32(v11)+64)) = l1
		v40 = int32(6)
		*(*uint16)(unsafe.Add(mBase, uint32(v11)+62)) = uint16(v40)
		*(*uint8)(unsafe.Add(mBase, uint32(v11)+60)) = uint8(v22)
		*(*int32)(unsafe.Add(mBase, uint32(v11)+56)) = v22
		*(*int64)(unsafe.Add(mBase, uint32(v11)+48)) = int64(0)
		*(*int32)(unsafe.Add(mBase, uint32(v11)+44)) = v14
		v51 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
		v52 = m.T0[v51].(func(*base.Module, int32) int32)(m, v11+int32(44))
		mBase = m.M
		v53 = m.ExcPending
		if v53 != 0 {
			return int32(0)
		} else {
			v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+60)))
			if v54 == int32(1) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v60 = m.ExcPending
				if v60 != 0 {
					return int32(0)
				} else {
					v61 = *(*int32)(unsafe.Add(mBase, uint32(v11)+20))
					*(*int32)(unsafe.Add(mBase, uint32(v11))) = v61
					F_errmsg_internal(m, int32(_a_F_OidFunctionCall6Coll_0), v11)
					mBase = m.M
					v65 = m.ExcPending
					if v65 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_OidFunctionCall6Coll_1), int32(1278), int32(_a_F_OidFunctionCall6Coll_2))
						mBase = m.M
						v70 = m.ExcPending
						if v70 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				m.G0 = v11 + int32(112)
				return v52
			}
		}
	}
}
func F_oid_increment(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	v5 = l1 + int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(base.B2i32(v5 == int32(0)))
	return v5
}
