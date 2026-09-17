package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_check_valid_internal_signature(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	if l0 != int32(2281) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	goto L3
L3:
	;
	v8 = int32(0)
	if l2 <= v8 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v29 = F_pstrdup(m, int32(_a_F_check_valid_internal_signature_0))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L12
	} else {
		goto L13
	}
L5:
	;
	v11 = v8
	goto L6
L6:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l1+v11<<(uint(int32(2))%32))))
	if v17 != int32(2281) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	return int32(0)
L8:
	;
	v21 = v11 + int32(1)
	if l2 != v21 {
		v11 = v21
		goto L6
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	goto L7
L11:
	;
	goto L4
L12:
	;
	return int32(0)
L13:
	;
	return v29
}
func F_check_valid_version_name(m *base.Module, l0 int32) {
	var v13 int32
	_ = v13
	Fn13864(m, l0, int32(_a_F_check_valid_version_name_0), int32(447), int32(_a_F_check_valid_version_name_1), int32(_a_F_check_valid_version_name_2), int32(437), int32(_a_F_check_valid_version_name_3), int32(428), int32(_a_F_check_valid_version_name_4), int32(419), int32(_a_F_check_valid_version_name_5))
	v13 = m.ExcPending
	if v13 != 0 {
		return
	} else {
		return
	}
}
