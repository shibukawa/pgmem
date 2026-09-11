package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_QTNClearFlags(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	F_check_stack_depth(m)
	mBase = m.M
	v5 = m.ExcPending
	if v5 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v6 & (l1 ^ int32(-1))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
	if v12 == int32(1) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return
L4:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v15 <= int32(0) {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v20 = int32(0)
	goto L6
L6:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v21+v20<<(uint(int32(2))%32))))
	F_QTNClearFlags(m, v25, l1)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L8
	}
L7:
	;
	goto L3
L8:
	;
	v29 = v20 + int32(1)
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v29 < v30 {
		v20 = v29
		goto L6
	} else {
		goto L9
	}
L9:
	;
	goto L7
}
func F_cmpQTN(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v5 = F_QTNodeCompare(m, v3, v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
