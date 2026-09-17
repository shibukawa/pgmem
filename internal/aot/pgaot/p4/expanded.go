package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_make_expanded_record_from_exprecord(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v59 int64
	_ = v59
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v6 == int32(0) {
		v9 = F_expanded_record_fetch_tupdesc(m, l0)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			v13 = v9
			v18 = F_AllocSetContextCreateInternal(m, l1, int32(_a_F_make_expanded_record_from_exprecord_0), int32(0), int32(_a_F_make_expanded_record_from_exprecord_1), int32(_a_F_make_expanded_record_from_exprecord_2))
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				v20 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
				v25 = F_MemoryContextAlloc(m, v18, v20*int32(5)+int32(120))
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int32(0)
				} else {
					base.MemoryFill(m, v25, int32(0), int32(120))
					v31 = int32(513)
					*(*uint16)(unsafe.Add(mBase, uint32(v25)+18)) = uint16(v31)
					v33 = int32(769)
					*(*uint16)(unsafe.Add(mBase, uint32(v25)+12)) = uint16(v33)
					*(*int32)(unsafe.Add(mBase, uint32(v25)+8)) = v18
					*(*int32)(unsafe.Add(mBase, uint32(v25)+4)) = int32(_a_F_make_expanded_record_from_exprecord_3)
					*(*int32)(unsafe.Add(mBase, uint32(v25))) = int32(-1)
					*(*int32)(unsafe.Add(mBase, uint32(v25)+20)) = v25
					*(*int32)(unsafe.Add(mBase, uint32(v25)+14)) = v25
					v42 = v25 + int32(120)
					*(*int32)(unsafe.Add(mBase, uint32(v25)+56)) = v42
					*(*int32)(unsafe.Add(mBase, uint32(v25)+24)) = int32(1384727874)
					v46 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
					*(*int32)(unsafe.Add(mBase, uint32(v25)+60)) = v42 + v46<<(uint(int32(2))%32)
					v51 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
					*(*int32)(unsafe.Add(mBase, uint32(v25)+64)) = v51
					v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
					*(*int32)(unsafe.Add(mBase, uint32(v25)+32)) = v53
					v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
					*(*int32)(unsafe.Add(mBase, uint32(v25)+36)) = v55
					v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
					*(*int32)(unsafe.Add(mBase, uint32(v25)+40)) = v57
					v59 = *(*int64)(unsafe.Add(mBase, uint32(l0)+48))
					*(*int64)(unsafe.Add(mBase, uint32(v25)+48)) = v59
					v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					*(*int32)(unsafe.Add(mBase, uint32(v25)+28)) = v61 & int32(64)
					v65 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
					if int32(0) <= v65 {
						*(*int32)(unsafe.Add(mBase, uint32(v25)+108)) = int32(1292)
						*(*int32)(unsafe.Add(mBase, uint32(v25)+112)) = v25
						v71 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
						v73 = v25 + int32(108)
						v74 = *(*int32)(unsafe.Add(mBase, uint32(v71)+40))
						*(*int32)(unsafe.Add(mBase, uint32(v73)+8)) = v74
						v76 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(v71)+4)) = uint8(v76)
						*(*int32)(unsafe.Add(mBase, uint32(v71)+40)) = v73
						*(*int32)(unsafe.Add(mBase, uint32(v25)+44)) = v13
						v80 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
						*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v80 + int32(1)
						return v25
					} else {
						v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
						if v85&int32(32) != 0 {
							v88 = int32(_a_F_make_expanded_record_from_exprecord_4)
							v89 = *(*int32)(unsafe.Add(mBase, _c_F_make_expanded_record_from_exprecord[0]))
							*(*int32)(unsafe.Add(mBase, _c_F_make_expanded_record_from_exprecord[0])) = v18
							v92 = F_CreateTupleDescCopy(m, v13)
							mBase = m.M
							v93 = m.ExcPending
							if v93 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v25)+44)) = v92
								v95 = *(*int32)(unsafe.Add(mBase, uint32(v25)+28))
								*(*int32)(unsafe.Add(mBase, uint32(v25)+28)) = v95 | int32(32)
								*(*int32)(unsafe.Add(mBase, _c_F_make_expanded_record_from_exprecord[0])) = v89
								return v25
							}
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v25)+44)) = v13
							return v25
						}
					}
				}
			}
		}
	} else {
		v13 = v6
		v18 = F_AllocSetContextCreateInternal(m, l1, int32(_a_F_make_expanded_record_from_exprecord_0), int32(0), int32(_a_F_make_expanded_record_from_exprecord_1), int32(_a_F_make_expanded_record_from_exprecord_2))
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int32(0)
		} else {
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
			v25 = F_MemoryContextAlloc(m, v18, v20*int32(5)+int32(120))
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return int32(0)
			} else {
				base.MemoryFill(m, v25, int32(0), int32(120))
				v31 = int32(513)
				*(*uint16)(unsafe.Add(mBase, uint32(v25)+18)) = uint16(v31)
				v33 = int32(769)
				*(*uint16)(unsafe.Add(mBase, uint32(v25)+12)) = uint16(v33)
				*(*int32)(unsafe.Add(mBase, uint32(v25)+8)) = v18
				*(*int32)(unsafe.Add(mBase, uint32(v25)+4)) = int32(_a_F_make_expanded_record_from_exprecord_3)
				*(*int32)(unsafe.Add(mBase, uint32(v25))) = int32(-1)
				*(*int32)(unsafe.Add(mBase, uint32(v25)+20)) = v25
				*(*int32)(unsafe.Add(mBase, uint32(v25)+14)) = v25
				v42 = v25 + int32(120)
				*(*int32)(unsafe.Add(mBase, uint32(v25)+56)) = v42
				*(*int32)(unsafe.Add(mBase, uint32(v25)+24)) = int32(1384727874)
				v46 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
				*(*int32)(unsafe.Add(mBase, uint32(v25)+60)) = v42 + v46<<(uint(int32(2))%32)
				v51 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
				*(*int32)(unsafe.Add(mBase, uint32(v25)+64)) = v51
				v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
				*(*int32)(unsafe.Add(mBase, uint32(v25)+32)) = v53
				v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
				*(*int32)(unsafe.Add(mBase, uint32(v25)+36)) = v55
				v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
				*(*int32)(unsafe.Add(mBase, uint32(v25)+40)) = v57
				v59 = *(*int64)(unsafe.Add(mBase, uint32(l0)+48))
				*(*int64)(unsafe.Add(mBase, uint32(v25)+48)) = v59
				v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
				*(*int32)(unsafe.Add(mBase, uint32(v25)+28)) = v61 & int32(64)
				v65 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
				if int32(0) <= v65 {
					*(*int32)(unsafe.Add(mBase, uint32(v25)+108)) = int32(1292)
					*(*int32)(unsafe.Add(mBase, uint32(v25)+112)) = v25
					v71 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
					v73 = v25 + int32(108)
					v74 = *(*int32)(unsafe.Add(mBase, uint32(v71)+40))
					*(*int32)(unsafe.Add(mBase, uint32(v73)+8)) = v74
					v76 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v71)+4)) = uint8(v76)
					*(*int32)(unsafe.Add(mBase, uint32(v71)+40)) = v73
					*(*int32)(unsafe.Add(mBase, uint32(v25)+44)) = v13
					v80 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
					*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v80 + int32(1)
					return v25
				} else {
					v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
					if v85&int32(32) != 0 {
						v88 = int32(_a_F_make_expanded_record_from_exprecord_4)
						v89 = *(*int32)(unsafe.Add(mBase, _c_F_make_expanded_record_from_exprecord[0]))
						*(*int32)(unsafe.Add(mBase, _c_F_make_expanded_record_from_exprecord[0])) = v18
						v92 = F_CreateTupleDescCopy(m, v13)
						mBase = m.M
						v93 = m.ExcPending
						if v93 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v25)+44)) = v92
							v95 = *(*int32)(unsafe.Add(mBase, uint32(v25)+28))
							*(*int32)(unsafe.Add(mBase, uint32(v25)+28)) = v95 | int32(32)
							*(*int32)(unsafe.Add(mBase, _c_F_make_expanded_record_from_exprecord[0])) = v89
							return v25
						}
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v25)+44)) = v13
						return v25
					}
				}
			}
		}
	}
}
