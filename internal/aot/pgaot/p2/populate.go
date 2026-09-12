package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_populate_recordset_record(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
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
	var v87 int32
	_ = v87
	v10 = m.G0
	v12 = v10 - int32(32)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v14)+68))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v14)+52))
	if v18 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v14)+68))
	v57 = F_populate_record(m, v51, v14+int32(48), v54, v55, l1, int32(0))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L8
	} else {
		goto L17
	}
L2:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v14)+60))
	v31 = F_lookup_rowtype_tupdesc(m, v28, v30)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L8
	} else {
		goto L9
	}
L3:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v14)+56))
	v28 = v21
	goto L2
L4:
	;
	goto L5
L5:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v14)+56))
	if v22 != v23 {
		v28 = v23
		goto L2
	} else {
		goto L6
	}
L6:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v14)+60))
	if v25 == v26 {
		v51 = v18
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v28 = v22
	goto L2
L8:
	;
	return
L9:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v14)+52))
	if v33 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	F_FreeTupleDesc(m, v33)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L8
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	v36 = int32(4515712)
	v37 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v17
	v40 = F_CreateTupleDescCopy(m, v31)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L8
	} else {
		goto L14
	}
L13:
	;
	goto L12
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+52)) = v40
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v37
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v31)+12))
	if v45 < int32(0) {
		v51 = v40
		goto L1
	} else {
		goto L15
	}
L15:
	;
	F_DecrTupleDescRefCount(m, v31)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L8
	} else {
		goto L16
	}
L16:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v14)+52))
	v51 = v50
	goto L1
L17:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	if v59 == int32(67) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v62 = F_HeapTupleHeaderGetDatum(m, v57)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L8
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+28)) = v57
	v74 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = v74
	*(*uint16)(unsafe.Add(mBase, uint32(v12)+20)) = uint16(v74)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = int32(base.Ui32(v72) >> (uint(int32(2)) % 32))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	F_tuplestore_puttuple(m, v83, v12+int32(12))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L8
	} else {
		goto L23
	}
L21:
	;
	v64 = int32(0)
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v14)+68))
	v70 = F_domain_check_safe(m, v62, v64, v65, v14-int32(-64), v68, v64)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L8
	} else {
		goto L22
	}
L22:
	;
	goto L20
L23:
	;
	m.G0 = v12 + int32(32)
	return
}
