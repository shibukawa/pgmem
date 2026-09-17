package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_checkcondition_arr_2(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v38 int32
	_ = v38
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if base.Ui32(v8) <= base.Ui32(v7) {
		v38 = int32(0)
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v38
L2:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v11 = v8
	v13 = v7
	goto L3
L3:
	;
	v18 = int32(2)
	v21 = base.I32_div_s((v11-v13)>>(uint(v18)%32), v18)
	v24 = v13 + v21<<(uint(v18)%32)
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	v26 = base.B2i32(v25 == v10)
	if v25 == v10 {
		v38 = v26
		goto L1
	} else {
		goto L5
	}
L4:
	;
	v38 = v26
	goto L1
L5:
	;
	v29 = base.B2i32(v25 < v10)
	if v25 < v10 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v30 = v24 + int32(4)
	goto L8
L7:
	;
	v30 = v13
	goto L8
L8:
	;
	if v25 < v10 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v31 = v11
	goto L11
L10:
	;
	v31 = v24
	goto L11
L11:
	;
	if base.Ui32(v30) < base.Ui32(v31) {
		v11 = v31
		v13 = v30
		goto L3
	} else {
		goto L12
	}
L12:
	;
	goto L4
}
