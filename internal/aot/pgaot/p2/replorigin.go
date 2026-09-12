package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_replorigin_get_progress(m *base.Module, l0 int32, l1 int32) int64 {
	mBase := m.M
	_ = mBase
	var v7 int64
	_ = v7
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int64
	_ = v45
	var v46 int64
	_ = v46
	var v48 int32
	_ = v48
	var v55 int64
	_ = v55
	var v56 int64
	_ = v56
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	v7 = int64(0)
	v10 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v14 = F_LWLockAcquire(m, v10+int32(5120), int32(1))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v19 = *(*int32)(unsafe.Add(mBase, _consts[643]))
	if v19 <= int32(0) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v58 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	F_LWLockRelease(m, v58+int32(5120))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L1
	} else {
		goto L14
	}
L4:
	;
	v55 = v7
	v56 = v7
	goto L3
L5:
	;
	goto L6
L6:
	;
	v23 = *(*int32)(unsafe.Add(mBase, _consts[644]))
	v26 = int32(0)
	goto L8
L7:
	;
	v41 = v34 + int32(40)
	v43 = F_LWLockAcquire(m, v41, int32(1))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L1
	} else {
		goto L12
	}
L8:
	;
	v34 = v23 + v26*int32(56)
	v35 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v34))))
	if v35 == l0 {
		goto L7
	} else {
		goto L10
	}
L9:
	;
	v55 = v7
	v56 = v7
	goto L3
L10:
	;
	v38 = v26 + int32(1)
	if v38 != v19 {
		v26 = v38
		goto L8
	} else {
		goto L11
	}
L11:
	;
	goto L9
L12:
	;
	v45 = *(*int64)(unsafe.Add(mBase, uint32(v34)+16))
	v46 = *(*int64)(unsafe.Add(mBase, uint32(v34)+8))
	F_LWLockRelease(m, v41)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	v55 = v45
	v56 = v46
	goto L3
L14:
	;
	if l1 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	return v56
L16:
	;
	if v55 == int64(0) {
		goto L15
	} else {
		goto L17
	}
L17:
	;
	F_XLogFlush(m, v55)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	goto L15
}
