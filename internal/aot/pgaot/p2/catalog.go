package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_CatalogIndexInsert(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v93 int32
	_ = v93
	v12 = m.G0
	v14 = v12 - int32(160)
	m.G0 = v14
	if l2 != int32(2) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v14 + int32(160)
	return
L2:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v19 = int32(*(*int16)(unsafe.Add(mBase, uint32(v18)+18)))
	if v19 < int32(0) {
		goto L1
	} else {
		goto L5
	}
L3:
	;
	goto L4
L4:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v22 == int32(0) {
		goto L1
	} else {
		goto L6
	}
L5:
	;
	goto L4
L6:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+52))
	v30 = F_MakeSingleTupleTableSlot(m, v28, int32(1618280))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	return
L8:
	;
	v33 = F_ExecStoreHeapTuple(m, l1, v30, int32(0))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	if int32(0) < v22 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v40 = int32(0)
	goto L13
L11:
	;
	goto L12
L12:
	;
	F_ExecDropSingleTupleTableSlot(m, v30)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L7
	} else {
		goto L24
	}
L13:
	;
	v52 = v40 << (uint(int32(2)) % 32)
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v25+v52)))
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54)+118)))
	if v55 != int32(1) {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	goto L12
L15:
	;
	v79 = v40 + int32(1)
	if v79 != v22 {
		v40 = v79
		goto L13
	} else {
		goto L23
	}
L16:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v52+v26)))
	if l2 == int32(2) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54)+123)))
	if v62 != int32(1) {
		goto L15
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	F_FormIndexDatum(m, v54, v30, int32(0), v14+int32(32), v14)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L7
	} else {
		goto L21
	}
L20:
	;
	goto L19
L21:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v59)+192))
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72)+12)))
	v75 = F_index_insert(m, v59, v14+int32(32), v14, l1+int32(4), v27, v73, int32(0), v54)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L7
	} else {
		goto L22
	}
L22:
	;
	goto L15
L23:
	;
	goto L14
L24:
	;
	goto L1
}
func F_CatalogTuplesMultiInsertWithInfo(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	v5 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	if v5 < l2 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v15 = F_GetCurrentCommandId(m, int32(1))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	m.G0 = v10 + int32(16)
	return
L4:
	;
	return
L5:
	;
	v17 = int32(0)
	F_heap_multi_insert(m, l0, l1, l2, v15, v17, v17)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v26 = v5
	goto L7
L7:
	;
	v30 = l1 + v26<<(uint(int32(2))%32)
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	v35 = F_ExecFetchSlotHeapTuple(m, v31, int32(1), v10+int32(15))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L4
	} else {
		goto L9
	}
L8:
	;
	goto L3
L9:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v35)+12)) = v38
	F_CatalogIndexInsert(m, l3, v35, int32(1))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L4
	} else {
		goto L10
	}
L10:
	;
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)))
	if v43 == int32(1) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	F_pfree(m, v35)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L4
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	v49 = v26 + int32(1)
	if v49 != l2 {
		v26 = v49
		goto L7
	} else {
		goto L15
	}
L14:
	;
	goto L13
L15:
	;
	goto L8
}
