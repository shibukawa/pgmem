package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_name_matches_visible_ENR(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	v4 = int32(0)
	if v3 == v4 {
		v36 = v4
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return base.B2i32(v36 != int32(0))
L2:
	;
	goto L1
L3:
	;
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v3)))
	if v9 == int32(0) {
		v36 = v4
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	if v12 <= int32(0) {
		v36 = v4
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
	v17 = int32(0)
	goto L6
L6:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v15+v17<<(uint(int32(2))%32))))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	v27 = F_strcmp(m, v26, l1)
	mBase = m.M
	if v27 == int32(0) {
		v36 = v25
		goto L2
	} else {
		goto L8
	}
L7:
	;
	v36 = int32(0)
	goto L2
L8:
	;
	v31 = v17 + int32(1)
	if v12 != v31 {
		v17 = v31
		goto L6
	} else {
		goto L9
	}
L9:
	;
	goto L7
}
