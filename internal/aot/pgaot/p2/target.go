package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_resolveTargetListUnknowns(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	v3 = int32(0)
	if l1 == v3 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v7 <= int32(0) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v12 = v3
	goto L4
L4:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v14+v12<<(uint(int32(2))%32))))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	v20 = F_exprType(m, v19)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L1
L6:
	;
	return
L7:
	;
	if v20 == int32(705) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	v27 = int32(-1)
	v31 = F_coerce_type(m, l0, v24, int32(705), int32(25), v27, int32(0), int32(2), v27)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L6
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	v35 = v12 + int32(1)
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v35 < v36 {
		v12 = v35
		goto L4
	} else {
		goto L12
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+4)) = v31
	goto L10
L12:
	;
	goto L5
}
