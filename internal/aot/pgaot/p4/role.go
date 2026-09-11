package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F__equalRoleSpec(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v42 int32
	_ = v42
	v3 = int32(0)
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v4 != v5 {
		v42 = v3
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v42
L2:
	;
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v8 != 0 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v42 = int32(1)
	goto L1
L4:
	;
	if v7 == int32(0) {
		v42 = v3
		goto L1
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	if v8 != v7 {
		v42 = v3
		goto L1
	} else {
		goto L17
	}
L7:
	;
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8))))
	if v14 == int32(0) {
		v33 = v13
		v34 = v14
		goto L9
	} else {
		goto L10
	}
L8:
	;
	if v34-v33 == int32(0) {
		goto L3
	} else {
		goto L16
	}
L9:
	;
	goto L8
L10:
	;
	if v13 != v14 {
		v33 = v13
		v34 = v14
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v18 = v8
	v19 = v7
	goto L12
L12:
	;
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+1)))
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+1)))
	if v23 == int32(0) {
		v33 = v22
		v34 = v23
		goto L9
	} else {
		goto L14
	}
L13:
	;
	v33 = v22
	v34 = v23
	goto L9
L14:
	;
	v26 = int32(1)
	if v22 == v23 {
		v18 = v18 + v26
		v19 = v19 + v26
		goto L12
	} else {
		goto L15
	}
L15:
	;
	goto L13
L16:
	;
	v42 = v3
	goto L1
L17:
	;
	goto L3
}
func F_assign_role(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v4 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+4)))
	F_SetCurrentRoleId(m, v3, v4)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return
	} else {
		return
	}
}
