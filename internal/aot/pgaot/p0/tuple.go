package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_CreateTupleDescCopy(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v10 = F_palloc(m, v5*int32(116)+int32(20))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v5
	*(*int64)(unsafe.Add(mBase, uint32(v10)+12)) = int64(4294967295)
	*(*int64)(unsafe.Add(mBase, uint32(v10)+4)) = int64(-4294965047)
	v20 = v5 * int32(100)
	if v20 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v21 = int32(4)
	v24 = int32(20)
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	base.MemoryCopy(m, v10+v5<<(uint(v21)%32)+v24, l0+v26<<(uint(v21)%32)+v24, v20)
	goto L5
L4:
	;
	goto L5
L5:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	if int32(0) < v33 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v39 = int32(0)
	v40 = v33
	goto L9
L7:
	;
	goto L8
L8:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v61
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v63
	return v10
L9:
	;
	v46 = v10 + v40<<(uint(int32(4))%32) + v39*int32(100)
	v47 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+110)) = uint8(v47)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+106)) = v47
	F_populate_compact_attribute(m, v10, v39)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L1
	} else {
		goto L11
	}
L10:
	;
	goto L8
L11:
	;
	v54 = v39 + int32(1)
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	if v54 < v55 {
		v39 = v54
		v40 = v55
		goto L9
	} else {
		goto L12
	}
L12:
	;
	goto L10
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
	var v55 int32
	_ = v55
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
	var v73 int32
	_ = v73
	var v82 int32
	_ = v82
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
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
	v49 = int32(_a_F_TupleDescInitBuiltinEntry_0)
	switch l3 - int32(16) {
	case 0:
		v55 = int32(1)
		v67 = int32(112)
		v68 = v55
		v69 = int32(0)
		v70 = int32(99)
		v71 = v55
		*(*int32)(unsafe.Add(mBase, uint32(v24)+96)) = v69
		v73 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v24)+85)) = uint8(v73)
		*(*uint8)(unsafe.Add(mBase, uint32(v24)+84)) = uint8(v67)
		*(*uint8)(unsafe.Add(mBase, uint32(v24)+83)) = uint8(v70)
		*(*uint8)(unsafe.Add(mBase, uint32(v24)+82)) = uint8(v71)
		*(*uint16)(unsafe.Add(mBase, uint32(v24)+72)) = uint16(v68)
		F_populate_compact_attribute(m, l0, v2-int32(1))
		mBase = m.M
		v82 = m.ExcPending
		if v82 != 0 {
			return
		} else {
			m.G0 = v13 + int32(16)
			return
		}
	case 1, 2, 3, 5, 6, 8:
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v89 = m.ExcPending
		if v89 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v13))) = l3
			F_errmsg_internal(m, int32(_a_F_TupleDescInitBuiltinEntry_1), v13)
			mBase = m.M
			v93 = m.ExcPending
			if v93 != 0 {
				return
			} else {
				F_errfinish(m, int32(_a_F_TupleDescInitBuiltinEntry_2), int32(1013), int32(_a_F_TupleDescInitBuiltinEntry_3))
				mBase = m.M
				v98 = m.ExcPending
				if v98 != 0 {
					return
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	case 4:
		v67 = int32(112)
		v68 = int32(8)
		v69 = int32(0)
		v70 = int32(100)
		v71 = v5
		*(*int32)(unsafe.Add(mBase, uint32(v24)+96)) = v69
		v73 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v24)+85)) = uint8(v73)
		*(*uint8)(unsafe.Add(mBase, uint32(v24)+84)) = uint8(v67)
		*(*uint8)(unsafe.Add(mBase, uint32(v24)+83)) = uint8(v70)
		*(*uint8)(unsafe.Add(mBase, uint32(v24)+82)) = uint8(v71)
		*(*uint16)(unsafe.Add(mBase, uint32(v24)+72)) = uint16(v68)
		F_populate_compact_attribute(m, l0, v2-int32(1))
		mBase = m.M
		v82 = m.ExcPending
		if v82 != 0 {
			return
		} else {
			m.G0 = v13 + int32(16)
			return
		}
	case 7, 10:
		v67 = int32(112)
		v68 = int32(4)
		v69 = int32(0)
		v70 = v48
		v71 = int32(1)
		*(*int32)(unsafe.Add(mBase, uint32(v24)+96)) = v69
		v73 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v24)+85)) = uint8(v73)
		*(*uint8)(unsafe.Add(mBase, uint32(v24)+84)) = uint8(v67)
		*(*uint8)(unsafe.Add(mBase, uint32(v24)+83)) = uint8(v70)
		*(*uint8)(unsafe.Add(mBase, uint32(v24)+82)) = uint8(v71)
		*(*uint16)(unsafe.Add(mBase, uint32(v24)+72)) = uint16(v68)
		F_populate_compact_attribute(m, l0, v2-int32(1))
		mBase = m.M
		v82 = m.ExcPending
		if v82 != 0 {
			return
		} else {
			m.G0 = v13 + int32(16)
			return
		}
	case 9:
		v67 = v47
		v68 = v49
		v69 = v15
		v70 = v48
		v71 = v5
		*(*int32)(unsafe.Add(mBase, uint32(v24)+96)) = v69
		v73 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v24)+85)) = uint8(v73)
		*(*uint8)(unsafe.Add(mBase, uint32(v24)+84)) = uint8(v67)
		*(*uint8)(unsafe.Add(mBase, uint32(v24)+83)) = uint8(v70)
		*(*uint8)(unsafe.Add(mBase, uint32(v24)+82)) = uint8(v71)
		*(*uint16)(unsafe.Add(mBase, uint32(v24)+72)) = uint16(v68)
		F_populate_compact_attribute(m, l0, v2-int32(1))
		mBase = m.M
		v82 = m.ExcPending
		if v82 != 0 {
			return
		} else {
			m.G0 = v13 + int32(16)
			return
		}
	default:
		if l3 != int32(1009) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v89 = m.ExcPending
			if v89 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v13))) = l3
				F_errmsg_internal(m, int32(_a_F_TupleDescInitBuiltinEntry_1), v13)
				mBase = m.M
				v93 = m.ExcPending
				if v93 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_TupleDescInitBuiltinEntry_2), int32(1013), int32(_a_F_TupleDescInitBuiltinEntry_3))
					mBase = m.M
					v98 = m.ExcPending
					if v98 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v67 = v47
			v68 = v49
			v69 = v15
			v70 = v48
			v71 = v5
			*(*int32)(unsafe.Add(mBase, uint32(v24)+96)) = v69
			v73 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v24)+85)) = uint8(v73)
			*(*uint8)(unsafe.Add(mBase, uint32(v24)+84)) = uint8(v67)
			*(*uint8)(unsafe.Add(mBase, uint32(v24)+83)) = uint8(v70)
			*(*uint8)(unsafe.Add(mBase, uint32(v24)+82)) = uint8(v71)
			*(*uint16)(unsafe.Add(mBase, uint32(v24)+72)) = uint16(v68)
			F_populate_compact_attribute(m, l0, v2-int32(1))
			mBase = m.M
			v82 = m.ExcPending
			if v82 != 0 {
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
