package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_in_grouping_b(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v17 int32
	_ = v17
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v17 = v10
	goto L2
L1:
	;
	return v51
L2:
	;
	if v17 <= v11 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v51 = int32(0)
	goto L1
L4:
	;
	return int32(-1)
L5:
	;
	goto L6
L6:
	;
	v24 = int32(1)
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25+v17-v24))))
	if l3 < v29 {
		v51 = v24
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v31 = v29 - l2
	if v31 < int32(0) {
		v51 = v24
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+int32(base.Ui32(v31)>>(uint(int32(3))%32))))))
	if int32(base.Ui32(v37)>>(uint(v31&int32(7))%32))&int32(1) == int32(0) {
		v51 = v24
		goto L1
	} else {
		goto L9
	}
L9:
	;
	v46 = v17 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v46
	if l4 != 0 {
		v17 = v46
		goto L2
	} else {
		goto L10
	}
L10:
	;
	goto L3
}
