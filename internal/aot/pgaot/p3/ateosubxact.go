package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_AtEOSubXact_Parallel(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	v5 = *(*int32)(unsafe.Add(mBase, _consts[153]))
	if v5 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	if v5 == int32(4146124) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v12 = v5
	goto L4
L4:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
	if v13 != l1 {
		goto L1
	} else {
		goto L6
	}
L5:
	;
	goto L1
L6:
	;
	if l0 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	F_DestroyParallelContext(m, v12)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L9
	} else {
		goto L14
	}
L8:
	;
	v19 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	return
L10:
	;
	if v19 == int32(0) {
		goto L7
	} else {
		goto L11
	}
L11:
	;
	F_errmsg_internal(m, int32(66120), int32(0))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L9
	} else {
		goto L12
	}
L12:
	;
	F_errfinish(m, int32(517893), int32(1271), int32(320314))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L9
	} else {
		goto L13
	}
L13:
	;
	goto L7
L14:
	;
	v35 = *(*int32)(unsafe.Add(mBase, _consts[153]))
	if v35 == int32(0) {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	if v35 != int32(4146124) {
		v12 = v35
		goto L4
	} else {
		goto L16
	}
L16:
	;
	goto L5
}
