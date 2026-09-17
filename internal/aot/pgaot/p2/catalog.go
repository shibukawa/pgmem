package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_CatalogIndexInsert(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v95 int32
	_ = v95
	v13 = m.G0
	v15 = v13 - int32(160)
	m.G0 = v15
	if l2 != int32(2) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v15 + int32(160)
	return
L2:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v20 = int32(*(*int16)(unsafe.Add(mBase, uint32(v19)+18)))
	if v20 < int32(0) {
		goto L1
	} else {
		goto L5
	}
L3:
	;
	goto L4
L4:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v23 == int32(0) {
		goto L1
	} else {
		goto L6
	}
L5:
	;
	goto L4
L6:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+52))
	v31 = F_MakeTupleTableSlot(m, v29, int32(_a_F_CatalogIndexInsert_0))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	return
L8:
	;
	v34 = F_ExecStoreHeapTuple(m, l1, v31, int32(0))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	if int32(0) < v23 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v41 = int32(0)
	goto L13
L11:
	;
	goto L12
L12:
	;
	F_ExecDropSingleTupleTableSlot(m, v31)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L7
	} else {
		goto L24
	}
L13:
	;
	v54 = v41 << (uint(int32(2)) % 32)
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v26+v54)))
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+118)))
	if v57 != int32(1) {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	goto L12
L15:
	;
	v80 = v41 + int32(1)
	if v80 != v23 {
		v41 = v80
		goto L13
	} else {
		goto L23
	}
L16:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v54+v27)))
	if l2 == int32(2) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+123)))
	if v64 != int32(1) {
		goto L15
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v69 = v15 + int32(32)
	F_FormIndexDatum(m, v56, v31, int32(0), v69, v15)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L7
	} else {
		goto L21
	}
L20:
	;
	goto L19
L21:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v61)+192))
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72)+12)))
	v75 = F_index_insert(m, v61, v69, v15, l1+int32(4), v28, v73, int32(0), v56)
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
