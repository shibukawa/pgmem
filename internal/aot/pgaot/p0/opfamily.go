package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_get_opfamily_member_for_cmptype(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int64
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v14 = base.I64_extend_i32_u(l0)
	v15 = F_SearchSysCache1(m, int32(42), v14)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		if v15 != 0 {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v15)+16))
			v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+22)))
			v22 = *(*int32)(unsafe.Add(mBase, uint32(v19+v20)+4))
			F_ReleaseCatCache(m, v15)
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int32(0)
			} else {
				v25 = int32(0)
				v27 = F_IndexAmTranslateCompareType(m, l3, v22, l0, int32(1))
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return int32(0)
				} else {
					if v27 == int32(0) {
						v48 = v25
						m.G0 = v11 + int32(16)
						return v48
					} else {
						v36 = F_SearchSysCache4(m, int32(4), v14, base.I64_extend_i32_u(l1), base.I64_extend_i32_u(l2), base.I64_extend16_s(base.I64_extend_i32_u(v27)))
						mBase = m.M
						v37 = m.ExcPending
						if v37 != 0 {
							return int32(0)
						} else {
							if v36 == int32(0) {
								v48 = v25
								m.G0 = v11 + int32(16)
								return v48
							} else {
								v40 = *(*int32)(unsafe.Add(mBase, uint32(v36)+16))
								v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+22)))
								v43 = *(*int32)(unsafe.Add(mBase, uint32(v40+v41)+20))
								F_ReleaseCatCache(m, v36)
								mBase = m.M
								v45 = m.ExcPending
								if v45 != 0 {
									return int32(0)
								} else {
									v48 = v43
									m.G0 = v11 + int32(16)
									return v48
								}
							}
						}
					}
				}
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v56 = m.ExcPending
			if v56 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v11))) = l0
				F_errmsg_internal(m, int32(_a_F_get_opfamily_member_for_cmptype_0), v11)
				mBase = m.M
				v60 = m.ExcPending
				if v60 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_get_opfamily_member_for_cmptype_1), int32(1542), int32(_a_F_get_opfamily_member_for_cmptype_2))
					mBase = m.M
					v65 = m.ExcPending
					if v65 != 0 {
						return int32(0)
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
