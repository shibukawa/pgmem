package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_out(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	v4 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v4&int32(32) == int32(0) {
		v9 = F___fwritex(m, l1, l2, l0)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return
		} else {
			return
		}
	} else {
		return
	}
}
func F_out_grouping(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
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
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v10 < v9 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v12 = v9
	goto L3
L2:
	;
	v12 = v10
	goto L3
L3:
	;
	v19 = v9
	goto L5
L4:
	;
	return v48
L5:
	;
	if v19 == v12 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v48 = int32(0)
	goto L4
L7:
	;
	return int32(-1)
L8:
	;
	goto L9
L9:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24+v19))))
	if l3 < v26 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v43 = v19 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v43
	if l4 != 0 {
		v19 = v43
		goto L5
	} else {
		goto L14
	}
L11:
	;
	v28 = v26 - l2
	if v28 < int32(0) {
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v31 = int32(1)
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+int32(base.Ui32(v28)>>(uint(int32(3))%32))))))
	if int32(base.Ui32(v35)>>(uint(v28&int32(7))%32))&v31 != 0 {
		v48 = v31
		goto L4
	} else {
		goto L13
	}
L13:
	;
	goto L10
L14:
	;
	goto L6
}
