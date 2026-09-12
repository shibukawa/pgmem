package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_StartupSUBTRANS(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int64
	_ = v7
	var v16 int32
	_ = v16
	var v17 int64
	_ = v17
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v46 int64
	_ = v46
	var v48 int32
	_ = v48
	v6 = *(*int32)(unsafe.Add(mBase, _consts[87]))
	v7 = *(*int64)(unsafe.Add(mBase, uint32(v6)+8))
	v16 = int32(0)
	v17 = base.I64_extend_i32_u(int32(base.Ui32(l0) >> (uint(int32(11)) % 32)))
	goto L1
L1:
	;
	v21 = *(*int32)(unsafe.Add(mBase, _consts[152]))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+28))
	v25 = int32(*(*uint16)(unsafe.Add(mBase, _consts[153])))
	v26 = base.I32_rem_u_s(base.I32_wrap_i64(v17), v25)
	v29 = v22 + v26<<(uint(int32(7))%32)
	if v29 != v16 {
		goto L3
	} else {
		goto L4
	}
L2:
	;
	F_LWLockRelease(m, v29)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L9
	} else {
		goto L19
	}
L3:
	;
	if v16 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	v36 = v16
	goto L5
L5:
	;
	v38 = F_SimpleLruZeroPage(m, int32(4350120), v17)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L9
	} else {
		goto L12
	}
L6:
	;
	F_LWLockRelease(m, v16)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	v34 = F_LWLockAcquire(m, v29, int32(0))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L9
	} else {
		goto L11
	}
L9:
	;
	return
L10:
	;
	goto L8
L11:
	;
	v36 = v29
	goto L5
L12:
	;
	if v17 != int64(base.Ui64(v7)>>(uint(int64(11))%64))&int64(2097151) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	if v17 < int64(2097151) {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	goto L15
L15:
	;
	goto L2
L16:
	;
	v46 = v17 + int64(1)
	goto L18
L17:
	;
	v46 = int64(0)
	goto L18
L18:
	;
	v16 = v36
	v17 = v46
	goto L1
L19:
	;
	return
}
