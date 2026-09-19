package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"sync/atomic"
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
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
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
	v40 = base.AtomicRmwXchg32(m, v37, int32(96), int32(1))
	if v40 != 0 {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	F_ConditionVariableCancelSleep(m)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
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
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)+80))
	v53 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v51)+96)), uint32(v53))
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
	v57 = m.ExcPending
	if v57 != 0 {
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
	v58 = F_CheckForStandbyTrigger(m)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L4
	} else {
		goto L31
	}
L31:
	;
	if v58 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	v61 = *(*int32)(unsafe.Add(mBase, _c_F_recoveryPausesHere[2]))
	v64 = base.AtomicRmwXchg32(m, v61, int32(96), int32(1))
	if v64 != 0 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v66 = *(*int32)(unsafe.Add(mBase, _c_F_recoveryPausesHere[2]))
	F_s_lock(m, v66+int32(96), int32(_a_F_recoveryPausesHere_4), int32(3135), int32(_a_F_recoveryPausesHere_7))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L4
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	v75 = *(*int32)(unsafe.Add(mBase, _c_F_recoveryPausesHere[2]))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v75)+80))
	if v76 == int32(1) {
		goto L37
	} else {
		goto L38
	}
L36:
	;
	goto L35
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v75)+80)) = int32(2)
	goto L39
L38:
	;
	goto L39
L39:
	;
	v81 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v75)+96)), uint32(v81))
	v88 = F_ConditionVariableTimedSleep(m, v75+int32(84), int32(1000), int32(134217775))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
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
