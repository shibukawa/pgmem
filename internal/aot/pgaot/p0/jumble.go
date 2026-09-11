package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F__jumbleCreateRoleStmt(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	F_AppendJumble32(m, l0, l1+int32(4))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v8 != 0 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	F__jumbleNode(m, l0, v74)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L1
	} else {
		goto L25
	}
L4:
	;
	if v8&int32(3) == int32(0) {
		v32 = v8
		goto L9
	} else {
		goto L10
	}
L5:
	;
	goto L6
L6:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v70 + int32(1)
	goto L3
L7:
	;
	F_AppendJumble(m, l0, v8, v65+int32(1))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L1
	} else {
		goto L24
	}
L8:
	;
	v65 = v57 - v8
	goto L7
L9:
	;
	v36 = v32
	goto L18
L10:
	;
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8))))
	if v16 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v65 = int32(0)
	goto L7
L12:
	;
	goto L13
L13:
	;
	v21 = v8
	goto L14
L14:
	;
	v25 = v21 + int32(1)
	if v25&int32(3) == int32(0) {
		v32 = v25
		goto L9
	} else {
		goto L16
	}
L15:
	;
	v57 = v25
	goto L8
L16:
	;
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25))))
	if v30 != 0 {
		v21 = v25
		goto L14
	} else {
		goto L17
	}
L17:
	;
	goto L15
L18:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
	v45 = int32(-2139062144)
	if (int32(16843008)-v42|v42)&v45 == v45 {
		v36 = v36 + int32(4)
		goto L18
	} else {
		goto L20
	}
L19:
	;
	v51 = v36
	goto L21
L20:
	;
	goto L19
L21:
	;
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51))))
	if v55 != 0 {
		v51 = v51 + int32(1)
		goto L21
	} else {
		goto L23
	}
L22:
	;
	v57 = v51
	goto L8
L23:
	;
	goto L22
L24:
	;
	goto L3
L25:
	;
	return
}
