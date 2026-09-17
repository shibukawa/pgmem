package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_HandleRecoveryConflictInterrupt(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	v4 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0<<(uint(int32(2))%32))+uint32(_c_F_HandleRecoveryConflictInterrupt[0]))) = v4
	*(*int32)(unsafe.Add(mBase, _c_F_HandleRecoveryConflictInterrupt[1])) = v4
	*(*int32)(unsafe.Add(mBase, _c_F_HandleRecoveryConflictInterrupt[2])) = v4
	return
}
func F_recoveryPausesHere(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	v4 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_recoveryPausesHere[0])))
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
	v8 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_recoveryPausesHere[1])))
	if v8&int32(1) != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v13 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return
L5:
	;
	if v13 != 0 {
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
	v17 = int32(_a_F_recoveryPausesHere_0)
	goto L11
L10:
	;
	v17 = int32(_a_F_recoveryPausesHere_1)
	goto L11
L11:
	;
	F_errmsg(m, v17, int32(0))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
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
	v23 = int32(_a_F_recoveryPausesHere_2)
	goto L15
L14:
	;
	v23 = int32(_a_F_recoveryPausesHere_3)
	goto L15
L15:
	;
	F_errhint(m, v23, int32(0))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
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
	v30 = int32(2960)
	goto L19
L18:
	;
	v30 = int32(2964)
	goto L19
L19:
	;
	F_errfinish(m, int32(_a_F_recoveryPausesHere_4), v30, int32(_a_F_recoveryPausesHere_5))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L4
	} else {
		goto L20
	}
L20:
	;
	goto L8
L21:
	;
	v37 = *(*int32)(unsafe.Add(mBase, _c_F_recoveryPausesHere[2]))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)+96))
	*(*int32)(unsafe.Add(mBase, uint32(v37)+96)) = int32(1)
	if v38 != 0 {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	F_ConditionVariableCancelSleep(m)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L4
	} else {
		goto L41
	}
L23:
	;
	v42 = *(*int32)(unsafe.Add(mBase, _c_F_recoveryPausesHere[2]))
	F_s_lock(m, v42+int32(96), int32(_a_F_recoveryPausesHere_4), int32(3096), int32(_a_F_recoveryPausesHere_6))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L4
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	v51 = *(*int32)(unsafe.Add(mBase, _c_F_recoveryPausesHere[2]))
	*(*int32)(unsafe.Add(mBase, uint32(v51)+96)) = int32(0)
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v51)+80))
	if v54 != 0 {
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
	v56 = m.ExcPending
	if v56 != 0 {
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
	v57 = F_CheckForStandbyTrigger(m)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L4
	} else {
		goto L31
	}
L31:
	;
	if v57 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	v60 = *(*int32)(unsafe.Add(mBase, _c_F_recoveryPausesHere[2]))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)+96))
	*(*int32)(unsafe.Add(mBase, uint32(v60)+96)) = int32(1)
	if v61 != 0 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v65 = *(*int32)(unsafe.Add(mBase, _c_F_recoveryPausesHere[2]))
	F_s_lock(m, v65+int32(96), int32(_a_F_recoveryPausesHere_4), int32(3135), int32(_a_F_recoveryPausesHere_7))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L4
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	v74 = *(*int32)(unsafe.Add(mBase, _c_F_recoveryPausesHere[2]))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)+80))
	if v75 == int32(1) {
		goto L37
	} else {
		goto L38
	}
L36:
	;
	goto L35
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v74)+80)) = int32(2)
	goto L39
L38:
	;
	goto L39
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v74)+96)) = int32(0)
	v86 = F_ConditionVariableTimedSleep(m, v74+int32(84), int32(1000), int32(134217775))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
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
