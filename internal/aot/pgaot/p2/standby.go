package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_StandbyDeadLockHandler(m *base.Module) {
	mBase := m.M
	_ = mBase
	*(*int32)(unsafe.Add(mBase, _consts[795])) = int32(1)
	return
}
func F_StandbyTimeoutHandler(m *base.Module) {
	mBase := m.M
	_ = mBase
	*(*int32)(unsafe.Add(mBase, _consts[796])) = int32(1)
	return
}
func F_standby_redo(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	v10 = m.G0
	v12 = v10 - int32(48)
	m.G0 = v12
	v15 = *(*int32)(unsafe.Add(mBase, _consts[51]))
	if v15 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v12 + int32(48)
	return
L2:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+48)))
	v21 = v19 & int32(240)
	switch v21 - int32(16) {
	case 0:
		goto L3
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15:
		goto L4
	case 16:
		goto L5
	default:
		goto L6
	}
L3:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v18)+64))
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v144)))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v145
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v144)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v147
	v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v144)+8)))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = v149
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v144)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+28)) = v151
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v144)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+40)) = v153
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v144)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+44)) = v144 + int32(24)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v155
	F_ProcArrayApplyRecoveryInfo(m, v12+int32(16))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L13
	} else {
		goto L36
	}
L4:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L13
	} else {
		goto L33
	}
L5:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v18)+64))
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v122)+12))
	v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122)+8)))
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v122)))
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v122)+4))
	F_ProcessCommittedInvalidationMessages(m, v122+int32(16), v125, v126, v127, v128)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L13
	} else {
		goto L32
	}
L6:
	;
	if v21 != 0 {
		goto L4
	} else {
		goto L7
	}
L7:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v18)+64))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	if v25 <= int32(0) {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v31 = int32(0)
	goto L9
L9:
	;
	v42 = v24 + int32(4) + v31*int32(12)
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v42)+8))
	v46 = m.G0
	v48 = v46 - int32(48)
	m.G0 = v48
	*(*int32)(unsafe.Add(mBase, uint32(v48)+44)) = v43
	if v43 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	goto L1
L11:
	;
	m.G0 = v48 + int32(48)
	v119 = v31 + int32(1)
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	if v119 < v120 {
		v31 = v119
		goto L9
	} else {
		goto L31
	}
L12:
	;
	v53 = F_TransactionIdDidCommit(m, v43)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	return
L14:
	;
	if v53 != 0 {
		goto L11
	} else {
		goto L15
	}
L15:
	;
	v55 = F_TransactionIdDidAbort(m, v43)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L13
	} else {
		goto L16
	}
L16:
	;
	if v55 != 0 {
		goto L11
	} else {
		goto L17
	}
L17:
	;
	v59 = F_errstart(m, int32(11), int32(0))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L13
	} else {
		goto L18
	}
L18:
	;
	if v59 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+4)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v48))) = v44
	F_errmsg_internal(m, int32(45049), v48)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L13
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v72 = *(*int32)(unsafe.Add(mBase, _consts[773]))
	v78 = F_hash_search(m, v72, v48+int32(44), int32(1), v48+int32(15))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L13
	} else {
		goto L24
	}
L22:
	;
	F_errfinish(m, int32(471083), int32(1000), int32(303362))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L13
	} else {
		goto L23
	}
L23:
	;
	goto L21
L24:
	;
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+15)))
	if v80 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v78)+4)) = int32(0)
	goto L27
L26:
	;
	goto L27
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+40)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v48)+36)) = v44
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v48)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v48)+32)) = v87
	v90 = *(*int32)(unsafe.Add(mBase, _consts[794]))
	v96 = F_hash_search(m, v90, v48+int32(32), int32(1), v48+int32(15))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L13
	} else {
		goto L28
	}
L28:
	;
	v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+15)))
	if v98 != 0 {
		goto L11
	} else {
		goto L29
	}
L29:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v78)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v96)+12)) = v99
	*(*int32)(unsafe.Add(mBase, uint32(v78)+4)) = v96
	*(*int64)(unsafe.Add(mBase, uint32(v48)+24)) = int64(72057594037927936)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+20)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v48)+16)) = v44
	v111 = F_LockAcquire(m, v48+int32(16), int32(8), int32(1), int32(0))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L13
	} else {
		goto L30
	}
L30:
	;
	goto L11
L31:
	;
	goto L10
L32:
	;
	goto L1
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v21
	F_errmsg_internal(m, int32(49951), v12)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L13
	} else {
		goto L34
	}
L34:
	;
	F_errfinish(m, int32(471083), int32(1219), int32(231105))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L13
	} else {
		goto L35
	}
L35:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L36:
	;
	v165 = F_pgstat_report_stat(m, int32(1))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L13
	} else {
		goto L37
	}
L37:
	;
	goto L1
}
