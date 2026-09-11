package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ResourceOwnerNewParent(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v5 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	if l1 != 0 {
		goto L10
	} else {
		goto L11
	}
L2:
	;
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v5)+4))
	if v8 == l0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v5)+4)) = v10
	goto L1
L4:
	;
	goto L5
L5:
	;
	v15 = v8
	goto L6
L6:
	;
	if v15 == int32(0) {
		goto L1
	} else {
		goto L8
	}
L7:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = v20
	goto L1
L8:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	if l0 != v18 {
		v15 = v18
		goto L6
	} else {
		goto L9
	}
L9:
	;
	goto L7
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = l1
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v27
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = l0
	return
L11:
	;
	goto L12
L12:
	;
	v30 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v30
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v30
	return
}
func F_ResourceOwnerRememberLock(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v14 int32
	_ = v14
	v4 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+18)))
	if base.Ui32(v4) <= base.Ui32(int32(15)) {
		if v4 != int32(15) {
			*(*int32)(unsafe.Add(mBase, uint32(l0+v4<<(uint(int32(2))%32))+292)) = l1
		} else {
		}
		v14 = v4 + int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+18)) = uint8(v14)
	} else {
	}
	return
}
