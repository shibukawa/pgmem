package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_IsSharedRelation(m *base.Module, l0 int32) int32 {
	var v4 int32
	_ = v4
	var v16 int32
	_ = v16
	var v29 int32
	_ = v29
	var v46 int32
	_ = v46
	var v75 int32
	_ = v75
	v4 = int32(1)
	if l0 <= int32(3591) {
		if l0 <= int32(2670) {
			switch l0 - int32(1213) {
			case 0, 1, 19, 20, 47, 48, 49:
				v75 = v4
			case 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46:
				v75 = int32(0)
			default:
				if base.Ui32(int32(2)) <= base.Ui32(l0-int32(2396)) {
					v75 = int32(0)
				} else {
					v75 = v4
				}
			}
		} else {
			v16 = l0 - int32(2671)
			if base.B2i32(base.Ui32(int32(27)) < base.Ui32(v16))|base.B2i32(int32(1)<<(uint(v16)%32)&int32(226492515) == int32(0)) != 0 {
				if base.B2i32(base.Ui32(l0-int32(2964)) < base.Ui32(int32(4)))|base.B2i32(base.Ui32(l0-int32(2846)) < base.Ui32(int32(2))) != 0 {
					v75 = v4
				} else {
					v75 = int32(0)
				}
			} else {
				v75 = v4
			}
		}
	} else {
		if l0 <= int32(_a_F_IsSharedRelation_0) {
			v29 = l0 - int32(_a_F_IsSharedRelation_1)
			if base.B2i32(base.Ui32(int32(9)) < base.Ui32(v29))|base.B2i32(int32(1)<<(uint(v29)%32)&int32(963) == int32(0)) != 0 {
				if base.Ui32(l0-int32(3592)) < base.Ui32(int32(2)) {
					v75 = v4
				} else {
					if base.Ui32(int32(2)) <= base.Ui32(l0-int32(4060)) {
						v75 = int32(0)
					} else {
						v75 = v4
					}
				}
			} else {
				v75 = v4
			}
		} else {
			switch l0 - int32(_a_F_IsSharedRelation_2) {
			case 0, 1, 2, 3, 4, 59, 60:
				v75 = v4
			case 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58:
				v75 = int32(0)
			default:
				if base.Ui32(l0-int32(_a_F_IsSharedRelation_3)) < base.Ui32(int32(3)) {
					v75 = v4
				} else {
					v46 = l0 - int32(_a_F_IsSharedRelation_4)
					if base.Ui32(int32(15)) < base.Ui32(v46) {
						v75 = int32(0)
					} else {
						if int32(1)<<(uint(v46)%32)&int32(_a_F_IsSharedRelation_5) != 0 {
							v75 = v4
						} else {
							v75 = int32(0)
						}
					}
				}
			}
		}
	}
	return v75
}
func F_shared_record_table_hash(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	v4 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)))
	if v4 == int32(1) {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v8 = F_dsa_get_address(m, l2, v7)
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			v12 = int32(0)
			v15 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
			v16 = F_hash_bytes_uint32(m, v15)
			mBase = m.M
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
			v18 = F_hash_bytes_uint32(m, v17)
			mBase = m.M
			v19 = int32(1640531527)
			v20 = v16 - v19
			v29 = v18 + v20<<(uint(int32(6))%32) + int32(base.Ui32(v20)>>(uint(int32(2))%32)) - v19 ^ v20
			v30 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
			if v12 < v30 {
				v34 = v29
				v35 = v30
				v36 = v12
				for {
					v43 = *(*int32)(unsafe.Add(mBase, uint32(v8+v35<<(uint(int32(4))%32)+v36*int32(100))+88))
					v44 = F_hash_bytes_uint32(m, v43)
					mBase = m.M
					v53 = v44 + (v34<<(uint(int32(6))%32) + int32(base.Ui32(v34)>>(uint(int32(2))%32))) - int32(1640531527) ^ v34
					v55 = v36 + int32(1)
					v56 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
					if v55 < v56 {
						v34 = v53
						v35 = v56
						v36 = v55
						continue
					} else {
						break
					}
					break
				}
				v59 = v53
			} else {
				v59 = v29
			}
			return v59
		}
	} else {
		v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v64 = int32(0)
		v67 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
		v68 = F_hash_bytes_uint32(m, v67)
		mBase = m.M
		v69 = *(*int32)(unsafe.Add(mBase, uint32(v63)+4))
		v70 = F_hash_bytes_uint32(m, v69)
		mBase = m.M
		v71 = int32(1640531527)
		v72 = v68 - v71
		v81 = v70 + v72<<(uint(int32(6))%32) + int32(base.Ui32(v72)>>(uint(int32(2))%32)) - v71 ^ v72
		v82 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
		if v64 < v82 {
			v86 = v81
			v87 = v82
			v88 = v64
			for {
				v95 = *(*int32)(unsafe.Add(mBase, uint32(v63+v87<<(uint(int32(4))%32)+v88*int32(100))+88))
				v96 = F_hash_bytes_uint32(m, v95)
				mBase = m.M
				v105 = v96 + (v86<<(uint(int32(6))%32) + int32(base.Ui32(v86)>>(uint(int32(2))%32))) - int32(1640531527) ^ v86
				v107 = v88 + int32(1)
				v108 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
				if v107 < v108 {
					v86 = v105
					v87 = v108
					v88 = v107
					continue
				} else {
					break
				}
				break
			}
			v111 = v105
		} else {
			v111 = v81
		}
		return v111
	}
}
