package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_grouping_is_sortable(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	if l0 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(1)
L2:
	;
	goto L3
L3:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v11 <= int32(0) {
		v39 = int32(1)
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return v39
L5:
	;
	v14 = int32(0)
	if v14 < v11 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v17 = v11
	goto L8
L7:
	;
	v17 = v14
	goto L8
L8:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v20 = int32(0)
	goto L9
L9:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v18+v20<<(uint(int32(2))%32))))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+12))
	v30 = int32(0)
	v31 = base.B2i32(v29 != v30)
	if v29 == v30 {
		v39 = v31
		goto L4
	} else {
		goto L11
	}
L10:
	;
	v39 = v31
	goto L4
L11:
	;
	v35 = v20 + int32(1)
	if v35 != v17 {
		v20 = v35
		goto L9
	} else {
		goto L12
	}
L12:
	;
	goto L10
}
