package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_out_grouping_b(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v16 = v9
	goto L2
L1:
	;
	return v48
L2:
	;
	if v16 <= v10 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v48 = int32(0)
	goto L1
L4:
	;
	return int32(-1)
L5:
	;
	goto L6
L6:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22+v16-int32(1)))))
	if l3 < v26 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v43 = v16 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v43
	if l4 != 0 {
		v16 = v43
		goto L2
	} else {
		goto L11
	}
L8:
	;
	v28 = v26 - l2
	if v28 < int32(0) {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v31 = int32(1)
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+int32(base.Ui32(v28)>>(uint(int32(3))%32))))))
	if int32(base.Ui32(v35)>>(uint(v28&int32(7))%32))&v31 != 0 {
		v48 = v31
		goto L1
	} else {
		goto L10
	}
L10:
	;
	goto L7
L11:
	;
	goto L3
}
