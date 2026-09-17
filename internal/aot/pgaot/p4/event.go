package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_AddEventToPendingNotifies(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
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
	var v28 int32
	_ = v28
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	v2 = int32(0)
	v6 = m.G0
	v8 = v6 + int32(-64)
	m.G0 = v8
	*(*int32)(unsafe.Add(mBase, uint32(v8)+60)) = l0
	v13 = *(*int32)(unsafe.Add(mBase, _c_F_AddEventToPendingNotifies[0]))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if v14 == v2 {
		v88 = v2
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v89 = F_lappend(m, v88, l0)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L5
	} else {
		goto L17
	}
L2:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	if v17 < int32(16) {
		v88 = v14
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	if v20 != 0 {
		v88 = v14
		goto L1
	} else {
		goto L4
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+40)) = int32(512)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+36)) = int32(513)
	*(*int64)(unsafe.Add(mBase, uint32(v8)+28)) = int64(17179869188)
	v28 = *(*int32)(unsafe.Add(mBase, _c_F_AddEventToPendingNotifies[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+52)) = v28
	v35 = F_hash_create(m, int32(_a_F_AddEventToPendingNotifies_0), int32(256), v6+int32(-52), int32(1224))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	return
L6:
	;
	v38 = *(*int32)(unsafe.Add(mBase, _c_F_AddEventToPendingNotifies[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v38)+8)) = v35
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
	if v40 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
	if int32(0) < v41 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v81 = v38
	goto L9
L9:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v81)+4))
	v88 = v82
	goto L1
L10:
	;
	v48 = v2
	goto L13
L11:
	;
	goto L12
L12:
	;
	v75 = *(*int32)(unsafe.Add(mBase, _c_F_AddEventToPendingNotifies[0]))
	v81 = v75
	goto L9
L13:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v40)+12))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v49+v48<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v53
	v56 = *(*int32)(unsafe.Add(mBase, _c_F_AddEventToPendingNotifies[0]))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)+8))
	v63 = F_hash_search(m, v57, v6+int32(-56), int32(1), v6+int32(-57))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L5
	} else {
		goto L15
	}
L14:
	;
	goto L12
L15:
	;
	v66 = v48 + int32(1)
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
	if v66 < v67 {
		v48 = v66
		goto L13
	} else {
		goto L16
	}
L16:
	;
	goto L14
L17:
	;
	v92 = *(*int32)(unsafe.Add(mBase, _c_F_AddEventToPendingNotifies[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v92)+4)) = v89
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v92)+8))
	if v94 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v100 = F_hash_search(m, v94, v6+int32(-4), int32(1), v6+int32(-52))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L5
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	m.G0 = v8 - int32(-64)
	return
L21:
	;
	goto L20
}
func F_EventTriggerAlterTableStart(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_EventTriggerAlterTableStart[0]))
	if v5 == int32(0) {
		return
	} else {
		v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5)+20)))
		if v8 != 0 {
			return
		} else {
			v9 = int32(_a_F_EventTriggerAlterTableStart_0)
			v10 = *(*int32)(unsafe.Add(mBase, _c_F_EventTriggerAlterTableStart[1]))
			v12 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
			*(*int32)(unsafe.Add(mBase, _c_F_EventTriggerAlterTableStart[1])) = v12
			v15 = F_palloc(m, int32(40))
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v15))) = int32(1)
				v20 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_EventTriggerAlterTableStart[2])))
				*(*uint8)(unsafe.Add(mBase, uint32(v15)+4)) = uint8(v20)
				*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = int32(0)
				*(*int64)(unsafe.Add(mBase, uint32(v15)+12)) = int64(5407363825664)
				v26 = F_copyObjectImpl(m, l0)
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = v26
					v30 = *(*int32)(unsafe.Add(mBase, _c_F_EventTriggerAlterTableStart[0]))
					v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+24))
					*(*int32)(unsafe.Add(mBase, uint32(v15)+36)) = v31
					*(*int32)(unsafe.Add(mBase, uint32(v30)+24)) = v15
					*(*int32)(unsafe.Add(mBase, _c_F_EventTriggerAlterTableStart[1])) = v10
					return
				}
			}
		}
	}
}
func F_EventTriggerTableRewrite(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v112 int32
	_ = v112
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v125 int64
	_ = v125
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	v4 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(208)
	m.G0 = v12
	v19 = v4
	v20 = v4
	v21 = v4
	v22 = v4
	v23 = int32(-1)
	goto L1
L1:
	;
	goto L3
L2:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L3:
	;
	if v23 != int32(1) {
		goto L8
	} else {
		goto L9
	}
L4:
	;
	goto L2
L5:
	;
	v124 = int32(m.ExcTag)
	v125 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v124 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_EventTriggerTableRewrite[0])) = v68
	*(*int32)(unsafe.Add(mBase, _c_F_EventTriggerTableRewrite[1])) = v67
	v112 = *(*int32)(unsafe.Add(mBase, _c_F_EventTriggerTableRewrite[2]))
	*(*int64)(unsafe.Add(mBase, uint32(v112)+12)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+200)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v12)+196)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v12)+204)) = v69
	F_pg_re_throw(m)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L5
	} else {
		goto L24
	}
L7:
	;
	m.G0 = v12 + int32(208)
	return
L8:
	;
	v27 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_EventTriggerTableRewrite[3])))
	if v27 != int32(1) {
		goto L7
	} else {
		goto L11
	}
L9:
	;
	v67 = v19
	v68 = v20
	v69 = v21
	v70 = v22
	goto L10
L10:
	;
	if v70 != 0 {
		goto L6
	} else {
		goto L20
	}
L11:
	;
	v31 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_EventTriggerTableRewrite[4])))
	if v31&int32(1) == int32(0) {
		goto L7
	} else {
		goto L12
	}
L12:
	;
	v37 = *(*int32)(unsafe.Add(mBase, _c_F_EventTriggerTableRewrite[2]))
	if v37 == int32(0) {
		goto L7
	} else {
		goto L13
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+200)) = v19
	*(*int32)(unsafe.Add(mBase, uint32(v12)+196)) = v20
	*(*int32)(unsafe.Add(mBase, uint32(v12)+204)) = v21
	v47 = F_EventTriggerCommonSetup(m, l0, int32(3), int32(_a_F_EventTriggerTableRewrite_0), v12+int32(180))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L5
	} else {
		goto L14
	}
L14:
	;
	if v47 == int32(0) {
		goto L7
	} else {
		goto L15
	}
L15:
	;
	v53 = *(*int32)(unsafe.Add(mBase, _c_F_EventTriggerTableRewrite[2]))
	*(*int32)(unsafe.Add(mBase, uint32(v53)+16)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v53)+12)) = l1
	v57 = *(*int32)(unsafe.Add(mBase, _c_F_EventTriggerTableRewrite[0]))
	v59 = *(*int32)(unsafe.Add(mBase, _c_F_EventTriggerTableRewrite[1]))
	goto L16
L16:
	;
	v61 = v12 + int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v61)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v61))) = v12 + int32(12)
	goto L19
L17:
	;
	v67 = v59
	v68 = v57
	v69 = v47
	v70 = int32(0)
	goto L10
L19:
	;
	goto L17
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+196)) = v68
	*(*int32)(unsafe.Add(mBase, _c_F_EventTriggerTableRewrite[1])) = v12 + int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+200)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v12)+204)) = v69
	F_EventTriggerInvoke(m, v69, v12+int32(180))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L5
	} else {
		goto L21
	}
L21:
	;
	v83 = *(*int32)(unsafe.Add(mBase, _c_F_EventTriggerTableRewrite[2]))
	*(*int64)(unsafe.Add(mBase, uint32(v83)+12)) = int64(0)
	*(*int32)(unsafe.Add(mBase, _c_F_EventTriggerTableRewrite[0])) = v68
	*(*int32)(unsafe.Add(mBase, _c_F_EventTriggerTableRewrite[1])) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v12)+196)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v12)+200)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v12)+204)) = v69
	F_list_free(m, v69)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L5
	} else {
		goto L22
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+200)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v12)+196)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v12)+204)) = v69
	F_CommandCounterIncrement(m)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L5
	} else {
		goto L23
	}
L23:
	;
	goto L7
L24:
	;
	goto L4
L25:
	;
	v129 = int32(v125)
	m.G0 = v12
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v129)+4))
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v129)))
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	if v12+int32(12) == v135 {
		goto L28
	} else {
		goto L29
	}
L26:
	;
	m.ExcPending = 1
	goto L34
L27:
	;
	if v139 != 0 {
		goto L31
	} else {
		goto L32
	}
L28:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v132)+4))
	v139 = v137
	goto L30
L29:
	;
	v139 = int32(0)
	goto L30
L30:
	;
	goto L27
L31:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v12)+204))
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v12)+200))
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v12)+196))
	v19 = v141
	v20 = v142
	v21 = v140
	v22 = v131
	v23 = v139
	goto L1
L32:
	;
	goto L33
L33:
	;
	F___wasm_longjmp(m, v132, v131)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	return
L35:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
