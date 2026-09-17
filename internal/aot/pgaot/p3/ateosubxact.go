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
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_AtEOSubXact_Parallel[0]))
	if base.B2i32(v5 == int32(0))|base.B2i32(v5 == int32(_a_F_AtEOSubXact_Parallel_0)) != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v13 = v5
	goto L3
L3:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	if v14 != l1 {
		goto L1
	} else {
		goto L5
	}
L4:
	;
	goto L1
L5:
	;
	if l0 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	F_DestroyParallelContext(m, v13)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L8
	} else {
		goto L13
	}
L7:
	;
	v20 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	return
L9:
	;
	if v20 == int32(0) {
		goto L6
	} else {
		goto L10
	}
L10:
	;
	F_errmsg_internal(m, int32(_a_F_AtEOSubXact_Parallel_1), int32(0))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L8
	} else {
		goto L11
	}
L11:
	;
	F_errfinish(m, int32(_a_F_AtEOSubXact_Parallel_2), int32(1271), int32(_a_F_AtEOSubXact_Parallel_3))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L8
	} else {
		goto L12
	}
L12:
	;
	goto L6
L13:
	;
	v36 = *(*int32)(unsafe.Add(mBase, _c_F_AtEOSubXact_Parallel[0]))
	if v36 == int32(0) {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	if v36 != int32(_a_F_AtEOSubXact_Parallel_0) {
		v13 = v36
		goto L3
	} else {
		goto L15
	}
L15:
	;
	goto L4
}
