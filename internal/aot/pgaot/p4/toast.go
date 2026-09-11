package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_toast_raw_datum_size(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v50 int32
	_ = v50
	var v57 int32
	_ = v57
	v4 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v4 == int32(1) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	return int32(base.Ui32(v57) >> (uint(int32(2)) % 32))
L2:
	;
	return int32(base.Ui32(v50)>>(uint(int32(1))%32)) + int32(3)
L3:
	;
	v7 = l0
	goto L6
L4:
	;
	v32 = l0
	v33 = v4
	goto L5
L5:
	;
	if v33&int32(3) == int32(2) {
		goto L18
	} else {
		goto L19
	}
L6:
	;
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+1)))
	if v10 != int32(1) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v32 = v28
	v33 = v29
	goto L5
L8:
	;
	if v10 == int32(18) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L10
L10:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v7)+2))
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28))))
	if v29 == int32(1) {
		v7 = v28
		goto L6
	} else {
		goto L17
	}
L11:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v7)+2))
	return v15
L12:
	;
	goto L13
L13:
	;
	if v10&int32(254) != int32(2) {
		v50 = int32(1)
		goto L2
	} else {
		goto L14
	}
L14:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v7)+2))
	v23 = F_EOH_get_flat_size(m, v22)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	return int32(0)
L16:
	;
	return v23
L17:
	;
	goto L7
L18:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v32)+4))
	return v39&int32(1073741823) + int32(4)
L19:
	;
	goto L20
L20:
	;
	if v33&int32(1) == int32(0) {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v50 = v33
	goto L2
}
