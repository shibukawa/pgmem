package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_get_opfamily_member_for_cmptype(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v13 = F_SearchSysCache1(m, int32(42), l0)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		if v13 != 0 {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
			v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+22)))
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v17+v18)+4))
			F_ReleaseCatCache(m, v13)
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int32(0)
			} else {
				v23 = int32(0)
				v25 = F_IndexAmTranslateCompareType(m, l3, v20, l0, int32(1))
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int32(0)
				} else {
					if v25 == int32(0) {
						v42 = v23
						m.G0 = v10 + int32(16)
						return v42
					} else {
						v31 = F_SearchSysCache4(m, int32(4), l0, l1, l2, base.I32_extend16_s(v25))
						mBase = m.M
						v32 = m.ExcPending
						if v32 != 0 {
							return int32(0)
						} else {
							if v31 == int32(0) {
								v42 = v23
								m.G0 = v10 + int32(16)
								return v42
							} else {
								v35 = *(*int32)(unsafe.Add(mBase, uint32(v31)+16))
								v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+22)))
								v38 = *(*int32)(unsafe.Add(mBase, uint32(v35+v36)+20))
								F_ReleaseCatCache(m, v31)
								mBase = m.M
								v40 = m.ExcPending
								if v40 != 0 {
									return int32(0)
								} else {
									v42 = v38
									m.G0 = v10 + int32(16)
									return v42
								}
							}
						}
					}
				}
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v50 = m.ExcPending
			if v50 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v10))) = l0
				F_errmsg_internal(m, int32(42667), v10)
				mBase = m.M
				v54 = m.ExcPending
				if v54 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(524184), int32(1384), int32(443010))
					mBase = m.M
					v59 = m.ExcPending
					if v59 != 0 {
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
