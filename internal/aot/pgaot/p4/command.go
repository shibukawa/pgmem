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
	var v10 int32
	_ = v10
	v2 = F_CreateCommandTag(m, l0)
	mBase = m.M
	v5 = m.ExcPending
	if v5 != 0 {
		return int32(0)
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(v2<<(uint(int32(3))%32))+uint32(_consts[505])))
		return v10
	}
}
func F_GetCommandTagEnum(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v9 int32
	_ = v9
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v44 int32
	_ = v44
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v67 int32
	_ = v67
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	v2 = int32(0)
	if l0 == v2 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v2
L2:
	;
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v9 == int32(0) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v15 = int32(1638736)
	v16 = int32(1640272)
	goto L4
L4:
	;
	v25 = v15 + (v16-v15)>>(uint(int32(4))%32)<<(uint(int32(3))%32)
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	v29 = l0
	v30 = v26
	goto L7
L5:
	;
	goto L1
L6:
	;
	if v67 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L7:
	;
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29))))
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30))))
	if v33 == v34 {
		v56 = v33
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v67 = int32(0)
	goto L6
L9:
	;
	v58 = int32(1)
	if v56 != 0 {
		v29 = v29 + v58
		v30 = v30 + v58
		goto L7
	} else {
		goto L18
	}
L10:
	;
	if base.Ui32((v33-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v44 = v33 | int32(32)
	goto L13
L12:
	;
	v44 = v33
	goto L13
L13:
	;
	if base.Ui32((v34-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v53 = v34 | int32(32)
	goto L16
L15:
	;
	v53 = v34
	goto L16
L16:
	;
	if v44 == v53 {
		v56 = v44
		goto L9
	} else {
		goto L17
	}
L17:
	;
	v67 = v44 - v53
	goto L6
L18:
	;
	goto L8
L19:
	;
	return (v25 - int32(1638736)) >> (uint(int32(3)) % 32)
L20:
	;
	goto L21
L21:
	;
	v78 = base.B2i32(v67 < int32(0))
	if v67 < int32(0) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v79 = v25 - int32(8)
	goto L24
L23:
	;
	v79 = v16
	goto L24
L24:
	;
	if v67 < int32(0) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v82 = v15
	goto L27
L26:
	;
	v82 = v25 + int32(8)
	goto L27
L27:
	;
	if base.Ui32(v82) <= base.Ui32(v79) {
		v15 = v82
		v16 = v79
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
	var v6 int32
	_ = v6
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0<<(uint(int32(3))%32))+uint32(_consts[505])))
	return v6
}
