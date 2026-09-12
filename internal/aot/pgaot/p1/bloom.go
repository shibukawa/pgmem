package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_BloomFormTuple(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1164))
	v9 = F_palloc0(m, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v13 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v9)+4)) = uint16(v13)
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1160))
	if int32(0) < v17 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v24 = int32(0)
	v28 = v17
	goto L6
L4:
	;
	goto L5
L5:
	;
	return v9
L6:
	;
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24+l3))))
	if v31 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	goto L5
L8:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l2+v24<<(uint(int32(2))%32))))
	F_signValue(m, l0, v9+int32(6), v37, v24)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L1
	} else {
		goto L11
	}
L9:
	;
	v41 = v28
	goto L10
L10:
	;
	v43 = v24 + int32(1)
	if v43 < v41 {
		v24 = v43
		v28 = v41
		goto L6
	} else {
		goto L12
	}
L11:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1160))
	v41 = v40
	goto L10
L12:
	;
	goto L7
}
