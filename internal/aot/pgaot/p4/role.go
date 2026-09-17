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
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v43 int32
	_ = v43
	v3 = int32(0)
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v4 != v5 {
		v43 = v3
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v43
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
	v43 = int32(1)
	goto L1
L4:
	;
	if v7 == int32(0) {
		v43 = v3
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
		v43 = v3
		goto L1
	} else {
		goto L16
	}
L7:
	;
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8))))
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
	if base.B2i32(v13 == int32(0))|base.B2i32(v13 != v16) != 0 {
		v34 = v13
		v35 = v16
		goto L9
	} else {
		goto L10
	}
L8:
	;
	if v34-v35 == int32(0) {
		goto L3
	} else {
		goto L15
	}
L9:
	;
	goto L8
L10:
	;
	v19 = v8
	v20 = v7
	goto L11
L11:
	;
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+1)))
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+1)))
	if v24 == int32(0) {
		v34 = v24
		v35 = v23
		goto L9
	} else {
		goto L13
	}
L12:
	;
	v34 = v24
	v35 = v23
	goto L9
L13:
	;
	v27 = int32(1)
	if v24 == v23 {
		v19 = v19 + v27
		v20 = v20 + v27
		goto L11
	} else {
		goto L14
	}
L14:
	;
	goto L12
L15:
	;
	v43 = v3
	goto L1
L16:
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
