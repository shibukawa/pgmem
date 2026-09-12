package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_CreateTupleDescCopy(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v11 = F_palloc(m, v6*int32(116)+int32(20))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v6
	*(*int64)(unsafe.Add(mBase, uint32(v11)+12)) = int64(4294967295)
	*(*int64)(unsafe.Add(mBase, uint32(v11)+4)) = int64(-4294965047)
	v20 = int32(20)
	v21 = v11 + v20
	v22 = int32(4)
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v32 = v6 * int32(100)
	if v32 != 0 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	if int32(0) < v35 {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	v33 = F__emscripten_memcpy_bulkmem(m, v21+v6<<(uint(v22)%32), l0+v25<<(uint(v22)%32)+v20, v32)
	mBase = m.M
	goto L6
L5:
	;
	goto L6
L6:
	;
	goto L3
L7:
	;
	v41 = int32(0)
	v42 = v35
	goto L10
L8:
	;
	goto L9
L9:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v65
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v67
	return v11
L10:
	;
	v49 = v21 + v42<<(uint(int32(4))%32) + v41*int32(100)
	v50 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v49)+90)) = uint8(v50)
	*(*int32)(unsafe.Add(mBase, uint32(v49)+86)) = v50
	F_populate_compact_attribute(m, v11, v41)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L1
	} else {
		goto L12
	}
L11:
	;
	goto L9
L12:
	;
	v57 = v41 + int32(1)
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	if v57 < v58 {
		v41 = v57
		v42 = v58
		goto L10
	} else {
		goto L13
	}
L13:
	;
	goto L11
}
func F_ExecInitMergeTupleSlots(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v6 = v4 + int32(104)
	v7 = F_table_slot_create(m, v3, v6)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l1)+44)) = v7
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
		v11 = F_table_slot_create(m, v10, v6)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return
		} else {
			v13 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l1)+48)) = uint8(v13)
			*(*int32)(unsafe.Add(mBase, uint32(l1)+40)) = v11
			return
		}
	}
}
func F_TupleDescInitBuiltinEntry(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v56 int32
	_ = v56
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
	var v74 int32
	_ = v74
	var v83 int32
	_ = v83
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	v2 = l1
	v5 = int32(0)
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	v15 = int32(100)
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v22 = l0 + v16<<(uint(int32(4))%32) + v2*v15
	v24 = v22 - int32(80)
	*(*int32)(unsafe.Add(mBase, uint32(v24))) = v5
	v30 = F_strncpy(m, v22-int32(76), l2, int32(64))
	mBase = m.M
	*(*uint8)(unsafe.Add(mBase, uint32(v30)+63)) = uint8(v5)
	v33 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+6)) = v33
	*(*uint16)(unsafe.Add(mBase, uint32(v22))) = uint16(v33)
	*(*uint16)(unsafe.Add(mBase, uint32(v24)+74)) = uint16(v2)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+76)) = int32(-1)
	*(*uint16)(unsafe.Add(mBase, uint32(v22)+10)) = uint16(v33)
	*(*uint16)(unsafe.Add(mBase, uint32(v22)+14)) = uint16(v33)
	v44 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+12)) = uint8(v44)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+68)) = l3
	v47 = int32(120)
	v48 = int32(105)
	v49 = int32(65535)
	switch l3 - int32(16) {
	case 0:
		v56 = int32(1)
		v68 = int32(112)
		v69 = v56
		v70 = int32(0)
		v71 = v56
		v72 = int32(99)
		*(*int32)(unsafe.Add(mBase, uint32(v24)+96)) = v70
		v74 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v24)+85)) = uint8(v74)
		*(*uint8)(unsafe.Add(mBase, uint32(v24)+84)) = uint8(v68)
		*(*uint8)(unsafe.Add(mBase, uint32(v24)+83)) = uint8(v72)
		*(*uint8)(unsafe.Add(mBase, uint32(v24)+82)) = uint8(v69)
		*(*uint16)(unsafe.Add(mBase, uint32(v24)+72)) = uint16(v71)
		F_populate_compact_attribute(m, l0, v2-int32(1))
		mBase = m.M
		v83 = m.ExcPending
		if v83 != 0 {
			return
		} else {
			m.G0 = v13 + int32(16)
			return
		}
	case 1, 2, 3, 5, 6, 8:
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v90 = m.ExcPending
		if v90 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v13))) = l3
			F_errmsg_internal(m, int32(51574), v13)
			mBase = m.M
			v94 = m.ExcPending
			if v94 != 0 {
				return
			} else {
				F_errfinish(m, int32(512068), int32(1013), int32(12619))
				mBase = m.M
				v99 = m.ExcPending
				if v99 != 0 {
					return
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	case 4:
		v68 = int32(112)
		v69 = v33
		v70 = int32(0)
		v71 = int32(8)
		v72 = int32(100)
		*(*int32)(unsafe.Add(mBase, uint32(v24)+96)) = v70
		v74 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v24)+85)) = uint8(v74)
		*(*uint8)(unsafe.Add(mBase, uint32(v24)+84)) = uint8(v68)
		*(*uint8)(unsafe.Add(mBase, uint32(v24)+83)) = uint8(v72)
		*(*uint8)(unsafe.Add(mBase, uint32(v24)+82)) = uint8(v69)
		*(*uint16)(unsafe.Add(mBase, uint32(v24)+72)) = uint16(v71)
		F_populate_compact_attribute(m, l0, v2-int32(1))
		mBase = m.M
		v83 = m.ExcPending
		if v83 != 0 {
			return
		} else {
			m.G0 = v13 + int32(16)
			return
		}
	case 7, 10:
		v68 = int32(112)
		v69 = int32(1)
		v70 = int32(0)
		v71 = int32(4)
		v72 = v48
		*(*int32)(unsafe.Add(mBase, uint32(v24)+96)) = v70
		v74 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v24)+85)) = uint8(v74)
		*(*uint8)(unsafe.Add(mBase, uint32(v24)+84)) = uint8(v68)
		*(*uint8)(unsafe.Add(mBase, uint32(v24)+83)) = uint8(v72)
		*(*uint8)(unsafe.Add(mBase, uint32(v24)+82)) = uint8(v69)
		*(*uint16)(unsafe.Add(mBase, uint32(v24)+72)) = uint16(v71)
		F_populate_compact_attribute(m, l0, v2-int32(1))
		mBase = m.M
		v83 = m.ExcPending
		if v83 != 0 {
			return
		} else {
			m.G0 = v13 + int32(16)
			return
		}
	case 9:
		v68 = v47
		v69 = v33
		v70 = v15
		v71 = v49
		v72 = v48
		*(*int32)(unsafe.Add(mBase, uint32(v24)+96)) = v70
		v74 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v24)+85)) = uint8(v74)
		*(*uint8)(unsafe.Add(mBase, uint32(v24)+84)) = uint8(v68)
		*(*uint8)(unsafe.Add(mBase, uint32(v24)+83)) = uint8(v72)
		*(*uint8)(unsafe.Add(mBase, uint32(v24)+82)) = uint8(v69)
		*(*uint16)(unsafe.Add(mBase, uint32(v24)+72)) = uint16(v71)
		F_populate_compact_attribute(m, l0, v2-int32(1))
		mBase = m.M
		v83 = m.ExcPending
		if v83 != 0 {
			return
		} else {
			m.G0 = v13 + int32(16)
			return
		}
	default:
		if l3 != int32(1009) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v90 = m.ExcPending
			if v90 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v13))) = l3
				F_errmsg_internal(m, int32(51574), v13)
				mBase = m.M
				v94 = m.ExcPending
				if v94 != 0 {
					return
				} else {
					F_errfinish(m, int32(512068), int32(1013), int32(12619))
					mBase = m.M
					v99 = m.ExcPending
					if v99 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v68 = v47
			v69 = v33
			v70 = v15
			v71 = v49
			v72 = v48
			*(*int32)(unsafe.Add(mBase, uint32(v24)+96)) = v70
			v74 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v24)+85)) = uint8(v74)
			*(*uint8)(unsafe.Add(mBase, uint32(v24)+84)) = uint8(v68)
			*(*uint8)(unsafe.Add(mBase, uint32(v24)+83)) = uint8(v72)
			*(*uint8)(unsafe.Add(mBase, uint32(v24)+82)) = uint8(v69)
			*(*uint16)(unsafe.Add(mBase, uint32(v24)+72)) = uint16(v71)
			F_populate_compact_attribute(m, l0, v2-int32(1))
			mBase = m.M
			v83 = m.ExcPending
			if v83 != 0 {
				return
			} else {
				m.G0 = v13 + int32(16)
				return
			}
		}
	}
}
func F_UnlockTuple(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	v5 = m.G0
	v6 = int32(16)
	v7 = v5 - v6
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v11
	v13 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+2)))
	v14 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1))))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v13 | v14<<(uint(v6)%32)
	v19 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
	v20 = int32(260)
	*(*uint16)(unsafe.Add(mBase, uint32(v7)+14)) = uint16(v20)
	*(*uint16)(unsafe.Add(mBase, uint32(v7)+12)) = uint16(v19)
	v24 = F_LockRelease(m, v7, l2, int32(0))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		return
	} else {
		m.G0 = v7 + int32(16)
		return
	}
}
