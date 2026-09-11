package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_uuid_gt(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v66 int32
	_ = v66
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v4 = int32(16)
	goto L4
L1:
	;
	return base.B2i32(int32(0) < v66)
L2:
	;
	v66 = int32(0)
	goto L1
L3:
	;
	v40 = v35
	v41 = v36
	v42 = v37
	goto L13
L4:
	;
	if (v2|v3)&int32(3) != 0 {
		v35 = v2
		v36 = v3
		v37 = v4
		goto L3
	} else {
		goto L7
	}
L6:
	;
	if v25 == int32(0) {
		goto L2
	} else {
		goto L12
	}
L7:
	;
	v12 = v2
	v13 = v3
	v14 = v4
	goto L8
L8:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	if v17 != v18 {
		v35 = v12
		v36 = v13
		v37 = v14
		goto L3
	} else {
		goto L10
	}
L9:
	;
	goto L6
L10:
	;
	v20 = int32(4)
	v21 = v13 + v20
	v23 = v12 + v20
	v25 = v14 - v20
	if base.Ui32(int32(3)) < base.Ui32(v25) {
		v12 = v23
		v13 = v21
		v14 = v25
		goto L8
	} else {
		goto L11
	}
L11:
	;
	goto L9
L12:
	;
	v35 = v23
	v36 = v21
	v37 = v25
	goto L3
L13:
	;
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40))))
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41))))
	if v45 == v46 {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v66 = v45 - v46
	goto L1
L15:
	;
	v48 = int32(1)
	v53 = v42 - v48
	if v53 != 0 {
		v40 = v40 + v48
		v41 = v41 + v48
		v42 = v53
		goto L13
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	goto L14
L18:
	;
	goto L2
}
