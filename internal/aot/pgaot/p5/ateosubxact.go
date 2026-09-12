package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_AtEOSubXact_Files(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	v9 = *(*int32)(unsafe.Add(mBase, _consts[647]))
	if v9 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v11 = *(*int32)(unsafe.Add(mBase, _consts[648]))
	v15 = int32(0)
	v16 = v9
	v17 = v11
	goto L4
L2:
	;
	goto L3
L3:
	;
	return
L4:
	;
	v21 = v17 + v15*int32(12)
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	if v22 != l1 {
		v33 = v15
		v34 = v16
		v35 = v17
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L3
L6:
	;
	v37 = v33 + int32(1)
	if base.Ui32(v37) < base.Ui32(v34) {
		v15 = v37
		v16 = v34
		v17 = v35
		goto L4
	} else {
		goto L13
	}
L7:
	;
	if l0 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+4)) = l2
	v33 = v15
	v34 = v16
	v35 = v17
	goto L6
L9:
	;
	goto L10
L10:
	;
	v25 = F_FreeDesc(m, v21)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	return
L12:
	;
	v30 = *(*int32)(unsafe.Add(mBase, _consts[647]))
	v32 = *(*int32)(unsafe.Add(mBase, _consts[648]))
	v33 = v15 - int32(1)
	v34 = v30
	v35 = v32
	goto L6
L13:
	;
	goto L5
}
func F_AtEOSubXact_on_commit_actions(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	v4 = int32(0)
	v10 = *(*int32)(unsafe.Add(mBase, _consts[223]))
	if v10 == v4 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	if l0 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v14 = l2
	goto L5
L4:
	;
	v14 = int32(0)
	goto L5
L5:
	;
	v18 = v4
	v19 = v10
	goto L6
L6:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	if v23 <= v18 {
		goto L1
	} else {
		goto L8
	}
L7:
	;
	goto L1
L8:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v25+v18<<(uint(int32(2))%32))))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+8))
	if l0 != 0 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	if v48 != 0 {
		v18 = v47 + int32(1)
		v19 = v48
		goto L6
	} else {
		goto L20
	}
L10:
	;
	if l1 == v30 {
		goto L16
	} else {
		goto L17
	}
L11:
	;
	if l1 != v30 {
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v32 = int32(4412640)
	v34 = *(*int32)(unsafe.Add(mBase, _consts[223]))
	v35 = F_list_delete_nth_cell(m, v34, v18)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	return
L14:
	;
	*(*int32)(unsafe.Add(mBase, _consts[223])) = v35
	F_pfree(m, v29)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v47 = v18 - int32(1)
	v48 = v35
	goto L9
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+8)) = l2
	goto L18
L17:
	;
	goto L18
L18:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v29)+12))
	if v44 != l1 {
		v47 = v18
		v48 = v19
		goto L9
	} else {
		goto L19
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+12)) = v14
	v47 = v18
	v48 = v19
	goto L9
L20:
	;
	goto L7
}
