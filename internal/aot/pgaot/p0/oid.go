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
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v40 int32
	_ = v40
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	v9 = m.G0
	v11 = v9 - int32(112)
	m.G0 = v11
	v16 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	F_fmgr_info_cxt_security(m, l0, v11+int32(16), v16, int32(0))
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
		*(*int32)(unsafe.Add(mBase, uint32(v11)+44)) = v11 + int32(16)
		v53 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
		v54 = m.T0[v53].(func(*base.Module, int32) int32)(m, v11+int32(44))
		mBase = m.M
		v55 = m.ExcPending
		if v55 != 0 {
			return int32(0)
		} else {
			v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+60)))
			if v56 == int32(1) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v62 = m.ExcPending
				if v62 != 0 {
					return int32(0)
				} else {
					v63 = *(*int32)(unsafe.Add(mBase, uint32(v11)+20))
					*(*int32)(unsafe.Add(mBase, uint32(v11))) = v63
					F_errmsg_internal(m, int32(556509), v11)
					mBase = m.M
					v67 = m.ExcPending
					if v67 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(517506), int32(1278), int32(317378))
						mBase = m.M
						v72 = m.ExcPending
						if v72 != 0 {
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
				return v54
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
