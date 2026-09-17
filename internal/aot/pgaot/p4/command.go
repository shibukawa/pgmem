package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_CreateCommandName(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v2 = F_CreateCommandTag(m, l0)
	mBase = m.M
	v5 = m.ExcPending
	if v5 != 0 {
		return int32(0)
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(v2<<(uint(int32(3))%32))+uint32(_c_F_CreateCommandName[0])))
		return v8
	}
}
func F_GetCommandTagEnum(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v42 int32
	_ = v42
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v65 int32
	_ = v65
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	if l0 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v8 == int32(0) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v14 = int32(_a_F_GetCommandTagEnum_0)
	v15 = int32(_a_F_GetCommandTagEnum_1)
	goto L4
L4:
	;
	v23 = v14 + (v15-v14)>>(uint(int32(4))%32)<<(uint(int32(3))%32)
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	v27 = l0
	v28 = v24
	goto L7
L5:
	;
	goto L1
L6:
	;
	if v65 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L7:
	;
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27))))
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28))))
	if v31 == v32 {
		v54 = v31
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v65 = int32(0)
	goto L6
L9:
	;
	v56 = int32(1)
	if v54 != 0 {
		v27 = v27 + v56
		v28 = v28 + v56
		goto L7
	} else {
		goto L18
	}
L10:
	;
	if base.Ui32((v31-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v42 = v31 | int32(32)
	goto L13
L12:
	;
	v42 = v31
	goto L13
L13:
	;
	if base.Ui32((v32-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v51 = v32 | int32(32)
	goto L16
L15:
	;
	v51 = v32
	goto L16
L16:
	;
	if v42 == v51 {
		v54 = v42
		goto L9
	} else {
		goto L17
	}
L17:
	;
	v65 = v42 - v51
	goto L6
L18:
	;
	goto L8
L19:
	;
	return (v23 - int32(_a_F_GetCommandTagEnum_0)) >> (uint(int32(3)) % 32)
L20:
	;
	goto L21
L21:
	;
	v76 = base.B2i32(v65 < int32(0))
	if v65 < int32(0) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v77 = v23 - int32(8)
	goto L24
L23:
	;
	v77 = v15
	goto L24
L24:
	;
	if v65 < int32(0) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v80 = v14
	goto L27
L26:
	;
	v80 = v23 + int32(8)
	goto L27
L27:
	;
	if base.Ui32(v80) <= base.Ui32(v77) {
		v14 = v80
		v15 = v77
		goto L4
	} else {
		goto L28
	}
L28:
	;
	goto L5
}
func F_GetCommandTagName(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0<<(uint(int32(3))%32))+uint32(_c_F_GetCommandTagName[0])))
	return v4
}
