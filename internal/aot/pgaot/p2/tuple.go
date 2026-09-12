package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_CreateTupleDescTruncatedCopy(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	v10 = F_palloc(m, l1*int32(116)+int32(20))
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
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = l1
	*(*int64)(unsafe.Add(mBase, uint32(v10)+12)) = int64(4294967295)
	*(*int64)(unsafe.Add(mBase, uint32(v10)+4)) = int64(-4294965047)
	v19 = int32(20)
	v20 = v10 + v19
	v21 = int32(4)
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v31 = l1 * int32(100)
	if v31 != 0 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	if int32(0) < v34 {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	v32 = F__emscripten_memcpy_bulkmem(m, v20+l1<<(uint(v21)%32), l0+v24<<(uint(v21)%32)+v19, v31)
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
	v39 = int32(0)
	v41 = v34
	goto L10
L8:
	;
	goto L9
L9:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v64
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v66
	return v10
L10:
	;
	v48 = v20 + v41<<(uint(int32(4))%32) + v39*int32(100)
	v49 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+90)) = uint8(v49)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+86)) = v49
	F_populate_compact_attribute(m, v10, v39)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L1
	} else {
		goto L12
	}
L11:
	;
	goto L9
L12:
	;
	v56 = v39 + int32(1)
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	if v56 < v57 {
		v39 = v56
		v41 = v57
		goto L10
	} else {
		goto L13
	}
L13:
	;
	goto L11
}
func F_CreateTupleQueueDestReceiver(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = F_palloc0(m, int32(24))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v4)+20)) = l0
		*(*int32)(unsafe.Add(mBase, uint32(v4)+16)) = int32(11)
		*(*int32)(unsafe.Add(mBase, uint32(v4)+12)) = int32(780)
		*(*int32)(unsafe.Add(mBase, uint32(v4)+8)) = int32(781)
		*(*int32)(unsafe.Add(mBase, uint32(v4)+4)) = int32(782)
		*(*int32)(unsafe.Add(mBase, uint32(v4))) = int32(783)
		return v4
	}
}
func F_ExecInitScanTupleSlot(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v17 int32
	_ = v17
	v6 = F_MakeTupleTableSlot(m, l2, l3)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
		v9 = F_lappend(m, v8, v6)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = v9
			*(*uint8)(unsafe.Add(mBase, uint32(l1)+96)) = uint8(base.B2i32(l2 != int32(0)))
			*(*int32)(unsafe.Add(mBase, uint32(l1)+76)) = l2
			*(*int32)(unsafe.Add(mBase, uint32(l1)+112)) = v6
			v17 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l1)+100)) = uint8(v17)
			*(*int32)(unsafe.Add(mBase, uint32(l1)+80)) = l3
			return
		}
	}
}
func F_GetTupleTransactionInfo(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
	v15 = m.T0[v14].(func(*base.Module, int32, int32, int32) int32)(m, l0, int32(-2), v8+int32(15))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l1))) = v15
		v20 = int32(0)
		v22 = int32(*(*uint8)(unsafe.Add(mBase, _consts[641])))
		if v22 == v20 {
			v25 = int32(0)
			*(*uint16)(unsafe.Add(mBase, uint32(l2))) = uint16(v25)
			*(*int64)(unsafe.Add(mBase, uint32(l3))) = int64(0)
			v31 = v20
			m.G0 = v8 + int32(16)
			return v31
		} else {
			v29 = F_TransactionIdGetCommitTsData(m, v15, l3, l2)
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return int32(0)
			} else {
				v31 = v29
				m.G0 = v8 + int32(16)
				return v31
			}
		}
	}
}
func F_TupleDescGetAttInMetadata(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v47 int32
	_ = v47
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	v12 = m.G0
	v13 = int32(16)
	v14 = v12 - v13
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v18 = F_palloc(m, v13)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v22 != int32(2249) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = l0
	v33 = F_palloc0(m, v16*int32(28))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L7
	}
L4:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if int32(0) <= v25 {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	F_assign_record_type_typmod(m, l0)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	goto L3
L7:
	;
	v36 = v16 << (uint(int32(2)) % 32)
	v37 = F_palloc0(m, v36)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v39 = F_palloc0(m, v36)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	if int32(0) < v16 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v47 = int32(0)
	goto L13
L11:
	;
	goto L12
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+12)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v18)+4)) = v33
	m.G0 = v14 + int32(16)
	return v18
L13:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v63 = l0 + int32(20) + v57<<(uint(int32(4))%32) + v47*int32(100)
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63)+91)))
	if v64 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	goto L12
L15:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v63)+68))
	v71 = v47 << (uint(int32(2)) % 32)
	F_getTypeInputInfo(m, v67, v14+int32(12), v37+v71)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L1
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	v86 = v47 + int32(1)
	if v86 != v16 {
		v47 = v86
		goto L13
	} else {
		goto L20
	}
L18:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	F_fmgr_info(m, v75, v33+v47*int32(28))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v63)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v39+v71))) = v82
	goto L17
L20:
	;
	goto L14
}
func F_TupleDescGetDefault(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v6 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	goto L3
L3:
	;
	v11 = int32(0)
	v12 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v6)+12)))
	if v12 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
	v16 = v11
	goto L8
L5:
	;
	v41 = v11
	goto L6
L6:
	;
	return v41
L7:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	v32 = F_stringToNode(m, v31)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L12
	} else {
		goto L13
	}
L8:
	;
	v23 = v13 + v16<<(uint(int32(3))%32)
	v24 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23))))
	if v24 == l1&int32(65535) {
		goto L7
	} else {
		goto L10
	}
L9:
	;
	return int32(0)
L10:
	;
	v27 = v16 + int32(1)
	if v27 != v12 {
		v16 = v27
		goto L8
	} else {
		goto L11
	}
L11:
	;
	goto L9
L12:
	;
	return int32(0)
L13:
	;
	v41 = v32
	goto L6
}
func F_TupleDescInitEntryCollation(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(l0+v4<<(uint(int32(4))%32)+l1*int32(100))+16)) = l2
	return
}
