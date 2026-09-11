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
	var v28 int32
	_ = v28
	var v44 int32
	_ = v44
	var v72 int32
	_ = v72
	v4 = int32(1)
	if l0 <= int32(3591) {
		if l0 <= int32(2670) {
			switch l0 - int32(1213) {
			case 0, 1, 19, 20, 47, 48, 49:
				v72 = v4
			case 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46:
				v72 = int32(0)
			default:
				if base.Ui32(int32(2)) <= base.Ui32(l0-int32(2396)) {
					v72 = int32(0)
				} else {
					v72 = v4
				}
			}
		} else {
			v16 = l0 - int32(2671)
			if base.Ui32(int32(27)) < base.Ui32(v16) {
				if base.Ui32(l0-int32(2964)) < base.Ui32(int32(4)) {
					v72 = v4
				} else {
					if base.Ui32(l0-int32(2846)) < base.Ui32(int32(2)) {
						v72 = v4
					} else {
						v72 = int32(0)
					}
				}
			} else {
				if int32(1)<<(uint(v16)%32)&int32(226492515) == int32(0) {
					if base.Ui32(l0-int32(2964)) < base.Ui32(int32(4)) {
						v72 = v4
					} else {
						if base.Ui32(l0-int32(2846)) < base.Ui32(int32(2)) {
							v72 = v4
						} else {
							v72 = int32(0)
						}
					}
				} else {
					v72 = v4
				}
			}
		}
	} else {
		if l0 <= int32(5999) {
			v28 = l0 - int32(4177)
			if base.Ui32(int32(9)) < base.Ui32(v28) {
				if base.Ui32(l0-int32(3592)) < base.Ui32(int32(2)) {
					v72 = v4
				} else {
					if base.Ui32(int32(2)) <= base.Ui32(l0-int32(4060)) {
						v72 = int32(0)
					} else {
						v72 = v4
					}
				}
			} else {
				if int32(1)<<(uint(v28)%32)&int32(963) == int32(0) {
					if base.Ui32(l0-int32(3592)) < base.Ui32(int32(2)) {
						v72 = v4
					} else {
						if base.Ui32(int32(2)) <= base.Ui32(l0-int32(4060)) {
							v72 = int32(0)
						} else {
							v72 = v4
						}
					}
				} else {
					v72 = v4
				}
			}
		} else {
			switch l0 - int32(6243) {
			case 0, 1, 2, 3, 4, 59, 60:
				v72 = v4
			case 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58:
				v72 = int32(0)
			default:
				if base.Ui32(l0-int32(6000)) < base.Ui32(int32(3)) {
					v72 = v4
				} else {
					v44 = l0 - int32(6100)
					if base.Ui32(int32(15)) < base.Ui32(v44) {
						v72 = int32(0)
					} else {
						if int32(1)<<(uint(v44)%32)&int32(49153) != 0 {
							v72 = v4
						} else {
							v72 = int32(0)
						}
					}
				}
			}
		}
	}
	return v72
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
	var v21 int32
	_ = v21
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
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
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
			v17 = F_hash_bytes_uint32(m, v16)
			mBase = m.M
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
			v19 = F_hash_bytes_uint32(m, v18)
			mBase = m.M
			v20 = int32(1640531527)
			v21 = v17 - v20
			v30 = v19 + v21<<(uint(int32(6))%32) + int32(base.Ui32(v21)>>(uint(int32(2))%32)) - v20 ^ v21
			v31 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
			if v12 < v31 {
				v37 = v30
				v38 = v12
				v39 = v31
				for {
					v47 = *(*int32)(unsafe.Add(mBase, uint32(v8+int32(88)+v39<<(uint(int32(4))%32)+v38*int32(100))))
					v48 = F_hash_bytes_uint32(m, v47)
					mBase = m.M
					v57 = v48 + (v37<<(uint(int32(6))%32) + int32(base.Ui32(v37)>>(uint(int32(2))%32))) - int32(1640531527) ^ v37
					v59 = v38 + int32(1)
					v60 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
					if v59 < v60 {
						v37 = v57
						v38 = v59
						v39 = v60
						continue
					} else {
						break
					}
					break
				}
				v63 = v57
			} else {
				v63 = v30
			}
			return v63
		}
	} else {
		v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v69 = int32(0)
		v73 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
		v74 = F_hash_bytes_uint32(m, v73)
		mBase = m.M
		v75 = *(*int32)(unsafe.Add(mBase, uint32(v68)+4))
		v76 = F_hash_bytes_uint32(m, v75)
		mBase = m.M
		v77 = int32(1640531527)
		v78 = v74 - v77
		v87 = v76 + v78<<(uint(int32(6))%32) + int32(base.Ui32(v78)>>(uint(int32(2))%32)) - v77 ^ v78
		v88 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
		if v69 < v88 {
			v94 = v87
			v95 = v69
			v96 = v88
			for {
				v104 = *(*int32)(unsafe.Add(mBase, uint32(v68+int32(88)+v96<<(uint(int32(4))%32)+v95*int32(100))))
				v105 = F_hash_bytes_uint32(m, v104)
				mBase = m.M
				v114 = v105 + (v94<<(uint(int32(6))%32) + int32(base.Ui32(v94)>>(uint(int32(2))%32))) - int32(1640531527) ^ v94
				v116 = v95 + int32(1)
				v117 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
				if v116 < v117 {
					v94 = v114
					v95 = v116
					v96 = v117
					continue
				} else {
					break
				}
				break
			}
			v120 = v114
		} else {
			v120 = v87
		}
		return v120
	}
}
