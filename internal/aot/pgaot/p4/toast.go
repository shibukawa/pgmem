package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_toast_raw_datum_size(m *base.Module, l0 int64) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int64
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v57 int32
	_ = v57
	var v64 int32
	_ = v64
	v5 = base.I32_wrap_i64(l0)
	v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5))))
	if v6 == int32(1) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
	return int32(base.Ui32(v64) >> (uint(int32(2)) % 32))
L2:
	;
	return int32(base.Ui32(v57)>>(uint(int32(1))%32)) + int32(3)
L3:
	;
	v9 = l0
	v10 = v5
	goto L6
L4:
	;
	v38 = v5
	v39 = v6
	goto L5
L5:
	;
	if v39&int32(3) == int32(2) {
		goto L19
	} else {
		goto L20
	}
L6:
	;
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+1)))
	if v13 != int32(1) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v38 = v32
	v39 = v34
	goto L5
L8:
	;
	if v13 == int32(18) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L10
L10:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v10)+2))
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32))))
	if v34 == int32(1) {
		v9 = base.I64_extend_i32_u(v32)
		v10 = v32
		goto L6
	} else {
		goto L18
	}
L11:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v10)+2))
	return v18
L12:
	;
	goto L13
L13:
	;
	if v13&int32(254) != int32(2) {
		v57 = int32(1)
		goto L2
	} else {
		goto L14
	}
L14:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v9))+2))
	goto L15
L15:
	;
	v27 = F_EOH_get_flat_size(m, v26)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	return int32(0)
L17:
	;
	return v27
L18:
	;
	goto L7
L19:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
	return v45&int32(1073741823) + int32(4)
L20:
	;
	goto L21
L21:
	;
	if v39&int32(1) == int32(0) {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	v57 = v39
	goto L2
}
