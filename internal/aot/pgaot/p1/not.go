package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_findNotNullConstraintAttnum(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	v3 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(48)
	m.G0 = v10
	v14 = F_table_open(m, int32(2606), int32(1))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	F_ScanKeyInit(m, v10, int32(9), int32(3), int32(184), l0)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v24 = int32(1)
	v27 = F_systable_beginscan(m, v14, int32(2665), v24, int32(0), v24, v10)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L5
	}
L4:
	;
	F_systable_endscan(m, v27)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L1
	} else {
		goto L17
	}
L5:
	;
	v29 = F_systable_getnext(m, v27)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	if v29 == int32(0) {
		v57 = v3
		goto L4
	} else {
		goto L7
	}
L7:
	;
	v33 = v29
	goto L8
L8:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v33)+16))
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+22)))
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40+v41)+72)))
	if v43 != int32(110) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v57 = v3
	goto L4
L10:
	;
	v51 = F_systable_getnext(m, v27)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L1
	} else {
		goto L15
	}
L11:
	;
	v46 = F_extractNotNullColumn(m, v33)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	if v46 != l1 {
		goto L10
	} else {
		goto L13
	}
L13:
	;
	v49 = F_heap_copytuple(m, v33)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	v57 = v49
	goto L4
L15:
	;
	if v51 != 0 {
		v33 = v51
		goto L8
	} else {
		goto L16
	}
L16:
	;
	goto L9
L17:
	;
	F_sequence_close(m, v14, int32(1))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	m.G0 = v10 + int32(48)
	return v57
}
