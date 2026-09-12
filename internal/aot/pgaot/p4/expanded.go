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
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v60 int64
	_ = v60
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v6 == int32(0) {
		v9 = F_expanded_record_fetch_tupdesc(m, l0)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			v13 = v9
			v18 = F_AllocSetContextCreateInternal(m, l1, int32(420880), int32(0), int32(8192), int32(8388608))
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
					v30 = F__emscripten_memset_bulkmem(m, v25, base.I32_extend8_s(int32(0)), int32(120))
					mBase = m.M
					v32 = int32(513)
					*(*uint16)(unsafe.Add(mBase, uint32(v30)+18)) = uint16(v32)
					v34 = int32(769)
					*(*uint16)(unsafe.Add(mBase, uint32(v30)+12)) = uint16(v34)
					*(*int32)(unsafe.Add(mBase, uint32(v30)+8)) = v18
					*(*int32)(unsafe.Add(mBase, uint32(v30)+4)) = int32(1645168)
					*(*int32)(unsafe.Add(mBase, uint32(v30))) = int32(-1)
					*(*int32)(unsafe.Add(mBase, uint32(v30)+20)) = v30
					*(*int32)(unsafe.Add(mBase, uint32(v30)+14)) = v30
					v43 = v30 + int32(120)
					*(*int32)(unsafe.Add(mBase, uint32(v30)+56)) = v43
					*(*int32)(unsafe.Add(mBase, uint32(v30)+24)) = int32(1384727874)
					v47 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
					*(*int32)(unsafe.Add(mBase, uint32(v30)+60)) = v43 + v47<<(uint(int32(2))%32)
					v52 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
					*(*int32)(unsafe.Add(mBase, uint32(v30)+64)) = v52
					v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
					*(*int32)(unsafe.Add(mBase, uint32(v30)+32)) = v54
					v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
					*(*int32)(unsafe.Add(mBase, uint32(v30)+36)) = v56
					v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
					*(*int32)(unsafe.Add(mBase, uint32(v30)+40)) = v58
					v60 = *(*int64)(unsafe.Add(mBase, uint32(l0)+48))
					*(*int64)(unsafe.Add(mBase, uint32(v30)+48)) = v60
					v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					*(*int32)(unsafe.Add(mBase, uint32(v30)+28)) = v62 & int32(64)
					v66 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
					if int32(0) <= v66 {
						*(*int32)(unsafe.Add(mBase, uint32(v30)+108)) = int32(1308)
						*(*int32)(unsafe.Add(mBase, uint32(v30)+112)) = v30
						v72 = *(*int32)(unsafe.Add(mBase, uint32(v30)+8))
						v74 = v30 + int32(108)
						v75 = *(*int32)(unsafe.Add(mBase, uint32(v72)+40))
						*(*int32)(unsafe.Add(mBase, uint32(v74)+8)) = v75
						v77 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(v72)+4)) = uint8(v77)
						*(*int32)(unsafe.Add(mBase, uint32(v72)+40)) = v74
						*(*int32)(unsafe.Add(mBase, uint32(v30)+44)) = v13
						v81 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
						*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v81 + int32(1)
						return v30
					} else {
						v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
						if v86&int32(32) != 0 {
							v89 = int32(4489440)
							v90 = *(*int32)(unsafe.Add(mBase, _consts[28]))
							*(*int32)(unsafe.Add(mBase, _consts[28])) = v18
							v93 = F_CreateTupleDescCopy(m, v13)
							mBase = m.M
							v94 = m.ExcPending
							if v94 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v30)+44)) = v93
								v96 = *(*int32)(unsafe.Add(mBase, uint32(v30)+28))
								*(*int32)(unsafe.Add(mBase, uint32(v30)+28)) = v96 | int32(32)
								*(*int32)(unsafe.Add(mBase, _consts[28])) = v90
								return v30
							}
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v30)+44)) = v13
							return v30
						}
					}
				}
			}
		}
	} else {
		v13 = v6
		v18 = F_AllocSetContextCreateInternal(m, l1, int32(420880), int32(0), int32(8192), int32(8388608))
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
				v30 = F__emscripten_memset_bulkmem(m, v25, base.I32_extend8_s(int32(0)), int32(120))
				mBase = m.M
				v32 = int32(513)
				*(*uint16)(unsafe.Add(mBase, uint32(v30)+18)) = uint16(v32)
				v34 = int32(769)
				*(*uint16)(unsafe.Add(mBase, uint32(v30)+12)) = uint16(v34)
				*(*int32)(unsafe.Add(mBase, uint32(v30)+8)) = v18
				*(*int32)(unsafe.Add(mBase, uint32(v30)+4)) = int32(1645168)
				*(*int32)(unsafe.Add(mBase, uint32(v30))) = int32(-1)
				*(*int32)(unsafe.Add(mBase, uint32(v30)+20)) = v30
				*(*int32)(unsafe.Add(mBase, uint32(v30)+14)) = v30
				v43 = v30 + int32(120)
				*(*int32)(unsafe.Add(mBase, uint32(v30)+56)) = v43
				*(*int32)(unsafe.Add(mBase, uint32(v30)+24)) = int32(1384727874)
				v47 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
				*(*int32)(unsafe.Add(mBase, uint32(v30)+60)) = v43 + v47<<(uint(int32(2))%32)
				v52 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
				*(*int32)(unsafe.Add(mBase, uint32(v30)+64)) = v52
				v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
				*(*int32)(unsafe.Add(mBase, uint32(v30)+32)) = v54
				v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
				*(*int32)(unsafe.Add(mBase, uint32(v30)+36)) = v56
				v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
				*(*int32)(unsafe.Add(mBase, uint32(v30)+40)) = v58
				v60 = *(*int64)(unsafe.Add(mBase, uint32(l0)+48))
				*(*int64)(unsafe.Add(mBase, uint32(v30)+48)) = v60
				v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
				*(*int32)(unsafe.Add(mBase, uint32(v30)+28)) = v62 & int32(64)
				v66 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
				if int32(0) <= v66 {
					*(*int32)(unsafe.Add(mBase, uint32(v30)+108)) = int32(1308)
					*(*int32)(unsafe.Add(mBase, uint32(v30)+112)) = v30
					v72 = *(*int32)(unsafe.Add(mBase, uint32(v30)+8))
					v74 = v30 + int32(108)
					v75 = *(*int32)(unsafe.Add(mBase, uint32(v72)+40))
					*(*int32)(unsafe.Add(mBase, uint32(v74)+8)) = v75
					v77 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v72)+4)) = uint8(v77)
					*(*int32)(unsafe.Add(mBase, uint32(v72)+40)) = v74
					*(*int32)(unsafe.Add(mBase, uint32(v30)+44)) = v13
					v81 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
					*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v81 + int32(1)
					return v30
				} else {
					v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
					if v86&int32(32) != 0 {
						v89 = int32(4489440)
						v90 = *(*int32)(unsafe.Add(mBase, _consts[28]))
						*(*int32)(unsafe.Add(mBase, _consts[28])) = v18
						v93 = F_CreateTupleDescCopy(m, v13)
						mBase = m.M
						v94 = m.ExcPending
						if v94 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v30)+44)) = v93
							v96 = *(*int32)(unsafe.Add(mBase, uint32(v30)+28))
							*(*int32)(unsafe.Add(mBase, uint32(v30)+28)) = v96 | int32(32)
							*(*int32)(unsafe.Add(mBase, _consts[28])) = v90
							return v30
						}
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v30)+44)) = v13
						return v30
					}
				}
			}
		}
	}
}
