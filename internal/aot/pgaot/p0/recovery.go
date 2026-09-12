package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_HandleRecoveryConflictInterrupt(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	v6 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0<<(uint(int32(2))%32))+uint32(_consts[666]))) = v6
	*(*int32)(unsafe.Add(mBase, _consts[615])) = v6
	*(*int32)(unsafe.Add(mBase, _consts[1])) = v6
	return
}
func F_recoveryPausesHere(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	v4 = int32(*(*uint8)(unsafe.Add(mBase, _consts[187])))
	if v4 != int32(1) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v8 = int32(*(*uint8)(unsafe.Add(mBase, _consts[184])))
	if v8 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v11 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return
L5:
	;
	if v11 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	if l0 != 0 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	goto L21
L9:
	;
	v15 = int32(14899)
	goto L11
L10:
	;
	v15 = int32(448838)
	goto L11
L11:
	;
	F_errmsg(m, v15, int32(0))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L4
	} else {
		goto L12
	}
L12:
	;
	if l0 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v21 = int32(630083)
	goto L15
L14:
	;
	v21 = int32(628385)
	goto L15
L15:
	;
	F_errhint(m, v21, int32(0))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L4
	} else {
		goto L16
	}
L16:
	;
	if l0 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v28 = int32(2960)
	goto L19
L18:
	;
	v28 = int32(2964)
	goto L19
L19:
	;
	F_errfinish(m, int32(492285), v28, int32(365128))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L4
	} else {
		goto L20
	}
L20:
	;
	goto L8
L21:
	;
	v35 = *(*int32)(unsafe.Add(mBase, _consts[186]))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)+96))
	*(*int32)(unsafe.Add(mBase, uint32(v35)+96)) = int32(1)
	if v36 != 0 {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	F_ConditionVariableCancelSleep(m)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L4
	} else {
		goto L41
	}
L23:
	;
	v40 = *(*int32)(unsafe.Add(mBase, _consts[186]))
	F_s_lock(m, v40+int32(96), int32(492285), int32(3096), int32(353994))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L4
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	v49 = *(*int32)(unsafe.Add(mBase, _consts[186]))
	*(*int32)(unsafe.Add(mBase, uint32(v49)+96)) = int32(0)
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v49)+80))
	if v52 != 0 {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	goto L25
L27:
	;
	F_ProcessStartupProcInterrupts(m)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L4
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	goto L22
L30:
	;
	v55 = F_CheckForStandbyTrigger(m)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L4
	} else {
		goto L31
	}
L31:
	;
	if v55 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	v58 = *(*int32)(unsafe.Add(mBase, _consts[186]))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)+96))
	*(*int32)(unsafe.Add(mBase, uint32(v58)+96)) = int32(1)
	if v59 != 0 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v63 = *(*int32)(unsafe.Add(mBase, _consts[186]))
	F_s_lock(m, v63+int32(96), int32(492285), int32(3135), int32(448858))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L4
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	v72 = *(*int32)(unsafe.Add(mBase, _consts[186]))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v72)+80))
	if v73 == int32(1) {
		goto L37
	} else {
		goto L38
	}
L36:
	;
	goto L35
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v72)+80)) = int32(2)
	goto L39
L38:
	;
	goto L39
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v72)+96)) = int32(0)
	v84 = F_ConditionVariableTimedSleep(m, v72+int32(84), int32(1000), int32(134217775))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L4
	} else {
		goto L40
	}
L40:
	;
	goto L21
L41:
	;
	goto L1
}
