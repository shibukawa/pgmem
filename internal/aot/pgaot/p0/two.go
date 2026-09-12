package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_TwoPhaseGetDummyProcNumber(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	v3 = F_TwoPhaseGetGXact(m, l0, l1)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(v3)+4))
		return v7
	}
}
func F_TwoPhaseGetGXact(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v39 int32
	_ = v39
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v76 int32
	_ = v76
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v13 = *(*int32)(unsafe.Add(mBase, _consts[145]))
	if v13 == l0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L9
	} else {
		goto L22
	}
L2:
	;
	m.G0 = v10 + int32(16)
	return v76
L3:
	;
	v16 = *(*int32)(unsafe.Add(mBase, _consts[146]))
	v76 = v16
	goto L2
L4:
	;
	goto L5
L5:
	;
	if l1 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v20 = *(*int32)(unsafe.Add(mBase, _consts[29]))
	v24 = F_LWLockAcquire(m, v20+int32(2304), int32(1))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	v29 = *(*int32)(unsafe.Add(mBase, _consts[147]))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
	if v30 <= int32(0) {
		v55 = int32(0)
		goto L11
	} else {
		goto L12
	}
L9:
	;
	return int32(0)
L10:
	;
	goto L8
L11:
	;
	if l1 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L12:
	;
	v39 = int32(0)
	goto L13
L13:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v29+int32(8)+v39<<(uint(int32(2))%32))))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)+32))
	if v47 == l0 {
		v55 = v46
		goto L11
	} else {
		goto L15
	}
L14:
	;
	v55 = int32(0)
	goto L11
L15:
	;
	v50 = v39 + int32(1)
	if v50 != v30 {
		v39 = v50
		goto L13
	} else {
		goto L16
	}
L16:
	;
	goto L14
L17:
	;
	v63 = *(*int32)(unsafe.Add(mBase, _consts[29]))
	F_LWLockRelease(m, v63+int32(2304))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L9
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	if v55 == int32(0) {
		goto L1
	} else {
		goto L21
	}
L20:
	;
	goto L19
L21:
	;
	*(*int32)(unsafe.Add(mBase, _consts[146])) = v55
	*(*int32)(unsafe.Add(mBase, _consts[145])) = l0
	v76 = v55
	goto L2
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = l0
	F_errmsg_internal(m, int32(59010), v10)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L9
	} else {
		goto L23
	}
L23:
	;
	F_errfinish(m, int32(524564), int32(835), int32(119557))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L9
	} else {
		goto L24
	}
L24:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
