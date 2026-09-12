package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_pg_wc_isupper(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v74 int32
	_ = v74
	v3 = *(*int32)(unsafe.Add(mBase, _consts[666]))
	switch v3 - int32(1) {
	case 0:
		goto L3
	case 1:
		goto L2
	case 2:
		goto L1
	default:
		goto L4
	}
L1:
	;
	if base.Ui32(l0) <= base.Ui32(int32(255)) {
		goto L20
	} else {
		goto L21
	}
L2:
	;
	v59 = F_towlower(m, l0)
	mBase = m.M
	goto L19
L3:
	;
	if base.Ui32(int32(127)) < base.Ui32(l0) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	return base.B2i32(base.Ui32(l0-int32(65)) < base.Ui32(int32(26)))
L5:
	;
	return v54
L6:
	;
	v19 = int32(0)
	v20 = int32(655)
	goto L9
L7:
	;
	goto L8
L8:
	;
	v54 = base.B2i32(base.Ui32(l0-int32(65)) < base.Ui32(int32(26)))
	goto L5
L9:
	;
	v25 = base.I32_div_s(v19+v20, int32(2))
	v27 = v25 << (uint(int32(3)) % 32)
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[672])))
	if base.Ui32(v30) < base.Ui32(l0) {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	v54 = int32(0)
	goto L5
L11:
	;
	if v41 <= v42 {
		v19 = v41
		v20 = v42
		goto L9
	} else {
		goto L18
	}
L12:
	;
	v41 = v25 + int32(1)
	v42 = v20
	goto L11
L13:
	;
	goto L14
L14:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[673])))
	if base.Ui32(v36) <= base.Ui32(l0) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v54 = int32(1)
	goto L5
L16:
	;
	goto L17
L17:
	;
	v41 = v19
	v42 = v25 - int32(1)
	goto L11
L18:
	;
	goto L10
L19:
	;
	return base.B2i32(v59 != l0)
L20:
	;
	goto L23
L21:
	;
	v74 = int32(0)
	goto L22
L22:
	;
	return v74
L23:
	;
	v74 = base.B2i32(base.B2i32(base.Ui32(l0-int32(65)) < base.Ui32(int32(26))) != int32(0))
	goto L22
}
