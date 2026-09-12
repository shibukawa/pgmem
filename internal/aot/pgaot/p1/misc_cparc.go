package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_cparc(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v63 int32
	_ = v63
	v8 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+4)))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_cparc[0]))
	if v11 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	if v14 <= v15 {
		goto L8
	} else {
		goto L9
	}
L4:
	;
	return
L5:
	;
	goto L3
L6:
	;
	return
L7:
	;
	F_createarc(m, l0, v9, v8, l2, l3)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L4
	} else {
		goto L27
	}
L8:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	if v17 == int32(0) {
		goto L7
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	if v36 == int32(0) {
		goto L7
	} else {
		goto L19
	}
L11:
	;
	v23 = v17
	goto L12
L12:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
	if v29 != l3 {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	goto L7
L14:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v23)+16))
	if v35 != 0 {
		v23 = v35
		goto L12
	} else {
		goto L18
	}
L15:
	;
	v31 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+4)))
	if v31 != v8&int32(_a_F_cparc_0) {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	if v33 == v9 {
		goto L6
	} else {
		goto L17
	}
L17:
	;
	goto L14
L18:
	;
	goto L13
L19:
	;
	v42 = v36
	goto L20
L20:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v42)+8))
	if v48 != l2 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	goto L7
L22:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v42)+24))
	if v54 != 0 {
		v42 = v54
		goto L20
	} else {
		goto L26
	}
L23:
	;
	v50 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v42)+4)))
	if v50 != v8&int32(_a_F_cparc_0) {
		goto L22
	} else {
		goto L24
	}
L24:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
	if v52 == v9 {
		goto L6
	} else {
		goto L25
	}
L25:
	;
	goto L22
L26:
	;
	goto L21
L27:
	;
	goto L6
}
