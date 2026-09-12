package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_cmpLexemeQ(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
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
	var v29 int32
	_ = v29
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v4 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return base.B2i32(v3 != int32(0))
L2:
	;
	goto L3
L3:
	;
	if v3 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return int32(-1)
L5:
	;
	goto L6
L6:
	;
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3))))
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4))))
	if v17 == int32(0) {
		v36 = v16
		v37 = v17
		goto L8
	} else {
		goto L9
	}
L7:
	;
	return v37 - v36
L8:
	;
	goto L7
L9:
	;
	if v16 != v17 {
		v36 = v16
		v37 = v17
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v21 = v4
	v22 = v3
	goto L11
L11:
	;
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+1)))
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+1)))
	if v26 == int32(0) {
		v36 = v25
		v37 = v26
		goto L8
	} else {
		goto L13
	}
L12:
	;
	v36 = v25
	v37 = v26
	goto L8
L13:
	;
	v29 = int32(1)
	if v25 == v26 {
		v21 = v21 + v29
		v22 = v22 + v29
		goto L11
	} else {
		goto L14
	}
L14:
	;
	goto L12
}
