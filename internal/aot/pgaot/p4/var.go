package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_makeVarFromTargetEntry(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	v7 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+8)))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v9 = F_exprType(m, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
		v14 = F_exprTypmod(m, v13)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			v17 = F_exprCollation(m, v16)
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				v20 = F_palloc0(m, int32(48))
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v20)+32)) = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v20)+20)) = v17
					*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = v14
					*(*int32)(unsafe.Add(mBase, uint32(v20)+12)) = v9
					*(*uint16)(unsafe.Add(mBase, uint32(v20)+8)) = uint16(v7)
					*(*int32)(unsafe.Add(mBase, uint32(v20)+4)) = l0
					*(*int32)(unsafe.Add(mBase, uint32(v20))) = int32(6)
					*(*int32)(unsafe.Add(mBase, uint32(v20)+44)) = int32(-1)
					*(*uint16)(unsafe.Add(mBase, uint32(v20)+40)) = uint16(v7)
					*(*int32)(unsafe.Add(mBase, uint32(v20)+36)) = l0
					*(*int64)(unsafe.Add(mBase, uint32(v20)+24)) = int64(0)
					return v20
				}
			}
		}
	}
}
func F_set_var_from_num(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v12 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+4)))
	if int32(0) <= v12 {
		v15 = int32(-8)
	} else {
		v15 = int32(-6)
	}
	v16 = int32(base.Ui32(v7)>>(uint(int32(2))%32)) + v15
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v19 != 0 {
		F_pfree(m, v19)
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return
		} else {
			v23 = v16 & int32(-2)
			v26 = F_palloc(m, v23+int32(2))
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = v26
				v29 = int32(0)
				*(*uint16)(unsafe.Add(mBase, uint32(v26))) = uint16(v29)
				*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(base.Ui32(v16) >> (uint(int32(1)) % 32))
				v32 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
				v34 = v32 + int32(2)
				*(*int32)(unsafe.Add(mBase, uint32(l1)+20)) = v34
				v36 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
				if base.I32_extend16_s(v36) < v29 {
					v50 = v36<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v36&int32(63)
				} else {
					v49 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+6)))
					v50 = v49
				}
				*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v50
				v52 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
				v53 = int32(49152)
				v54 = v52 & v53
				if v54 != v53 {
					if v54 != int32(32768) {
						v65 = v54
					} else {
						v65 = v52 << (uint(int32(1)) % 32) & int32(16384)
					}
				} else {
					v65 = v52 & int32(61440)
				}
				*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v65
				v67 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+4)))
				if v67 < int32(0) {
					v76 = int32(base.Ui32(v67)>>(uint(int32(7))%32)) & int32(63)
				} else {
					v76 = v67 & int32(16383)
				}
				*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v76
				v80 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+4)))
				if v80 < int32(0) {
					v83 = int32(6)
				} else {
					v83 = int32(8)
				}
				if v23 != 0 {
					v85 = F__emscripten_memcpy_bulkmem(m, v34, l0+v83, v23)
					mBase = m.M
				} else {
				}
				return
			}
		}
	} else {
		v23 = v16 & int32(-2)
		v26 = F_palloc(m, v23+int32(2))
		mBase = m.M
		v27 = m.ExcPending
		if v27 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = v26
			v29 = int32(0)
			*(*uint16)(unsafe.Add(mBase, uint32(v26))) = uint16(v29)
			*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(base.Ui32(v16) >> (uint(int32(1)) % 32))
			v32 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
			v34 = v32 + int32(2)
			*(*int32)(unsafe.Add(mBase, uint32(l1)+20)) = v34
			v36 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
			if base.I32_extend16_s(v36) < v29 {
				v50 = v36<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v36&int32(63)
			} else {
				v49 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+6)))
				v50 = v49
			}
			*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v50
			v52 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
			v53 = int32(49152)
			v54 = v52 & v53
			if v54 != v53 {
				if v54 != int32(32768) {
					v65 = v54
				} else {
					v65 = v52 << (uint(int32(1)) % 32) & int32(16384)
				}
			} else {
				v65 = v52 & int32(61440)
			}
			*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v65
			v67 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+4)))
			if v67 < int32(0) {
				v76 = int32(base.Ui32(v67)>>(uint(int32(7))%32)) & int32(63)
			} else {
				v76 = v67 & int32(16383)
			}
			*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v76
			v80 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+4)))
			if v80 < int32(0) {
				v83 = int32(6)
			} else {
				v83 = int32(8)
			}
			if v23 != 0 {
				v85 = F__emscripten_memcpy_bulkmem(m, v34, l0+v83, v23)
				mBase = m.M
			} else {
			}
			return
		}
	}
}
