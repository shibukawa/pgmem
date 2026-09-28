package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"sync/atomic"
	"unsafe"
)

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
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
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
	v30 = int32(2924)
	goto L19
L18:
	;
	v30 = int32(2928)
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
	v83 = m.ExcPending
	if v83 != 0 {
		goto L4
	} else {
		goto L41
	}
L23:
	;
	F_s_lock(m, v37+int32(96), int32(_a_F_recoveryPausesHere_6))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L4
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	v47 = *(*int32)(unsafe.Add(mBase, _c_F_recoveryPausesHere[2]))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)+80))
	v49 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v47)+96)), uint32(v49))
	if v48 != 0 {
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
	v53 = m.ExcPending
	if v53 != 0 {
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
	v54 = F_CheckForStandbyTrigger(m)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L4
	} else {
		goto L31
	}
L31:
	;
	if v54 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	v57 = *(*int32)(unsafe.Add(mBase, _c_F_recoveryPausesHere[2]))
	v60 = base.AtomicRmwXchg32(m, v57, int32(96), int32(1))
	if v60 != 0 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	F_s_lock(m, v57+int32(96), int32(_a_F_recoveryPausesHere_6))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L4
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	v67 = *(*int32)(unsafe.Add(mBase, _c_F_recoveryPausesHere[2]))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)+80))
	if v68 == int32(1) {
		goto L37
	} else {
		goto L38
	}
L36:
	;
	goto L35
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v67)+80)) = int32(2)
	goto L39
L38:
	;
	goto L39
L39:
	;
	v73 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v67)+96)), uint32(v73))
	v80 = F_ConditionVariableTimedSleep(m, v67+int32(84), int32(1000), int32(134217775))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
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
