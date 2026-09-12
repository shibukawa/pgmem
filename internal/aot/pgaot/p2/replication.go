package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ReplicationOriginExitCleanup(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	v4 = *(*int32)(unsafe.Add(mBase, _consts[648]))
	if v4 != 0 {
		v6 = *(*int32)(unsafe.Add(mBase, _consts[2]))
		v10 = F_LWLockAcquire(m, v6+int32(5120), int32(0))
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, _consts[648]))
			v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+24))
			v16 = *(*int32)(unsafe.Add(mBase, _consts[353]))
			if v14 != v16 {
				v19 = *(*int32)(unsafe.Add(mBase, _consts[2]))
				F_LWLockRelease(m, v19+int32(5120))
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return
				} else {
					return
				}
			} else {
				v24 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v13)+24)) = v24
				*(*int32)(unsafe.Add(mBase, _consts[648])) = v24
				v30 = *(*int32)(unsafe.Add(mBase, _consts[2]))
				F_LWLockRelease(m, v30+int32(5120))
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return
				} else {
					F_ConditionVariableBroadcast(m, v13+int32(28))
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
						return
					} else {
						return
					}
				}
			}
		}
	} else {
		return
	}
}
func F_ReplicationSlotAcquire(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
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
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v177 int32
	_ = v177
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v234 int32
	_ = v234
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v275 int32
	_ = v275
	var v280 int32
	_ = v280
	var v291 int32
	_ = v291
	var v295 int32
	_ = v295
	var v299 int32
	_ = v299
	var v302 int32
	_ = v302
	var v306 int32
	_ = v306
	var v311 int32
	_ = v311
	v10 = m.G0
	v12 = v10 - int32(80)
	m.G0 = v12
	v15 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v19 = F_LWLockAcquire(m, v15+int32(4736), int32(1))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v22 = *(*int32)(unsafe.Add(mBase, _consts[678]))
	if v22 <= int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v291 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	F_LWLockRelease(m, v291+int32(4736))
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L1
	} else {
		goto L96
	}
L4:
	;
	v33 = v22
	goto L6
L5:
	;
	if l1 == int32(0) {
		goto L60
	} else {
		goto L61
	}
L6:
	;
	v36 = *(*int32)(unsafe.Add(mBase, _consts[679]))
	v42 = int32(0)
	goto L8
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L1
	} else {
		goto L56
	}
L8:
	;
	v48 = v36 + v42*int32(288)
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+4)))
	if v49 == int32(1) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	if v48 == int32(0) {
		goto L3
	} else {
		goto L24
	}
L10:
	;
	goto L9
L11:
	;
	v53 = v48 + int32(24)
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53))))
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v57 == int32(0) {
		v76 = v56
		v77 = v57
		goto L15
	} else {
		goto L16
	}
L12:
	;
	goto L13
L13:
	;
	v83 = v42 + int32(1)
	if v83 != v33 {
		v42 = v83
		goto L8
	} else {
		goto L23
	}
L14:
	;
	if v77-v76 == int32(0) {
		goto L10
	} else {
		goto L22
	}
L15:
	;
	goto L14
L16:
	;
	if v56 != v57 {
		v76 = v56
		v77 = v57
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v61 = l0
	v62 = v53
	goto L18
L18:
	;
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62)+1)))
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61)+1)))
	if v66 == int32(0) {
		v76 = v65
		v77 = v66
		goto L15
	} else {
		goto L20
	}
L19:
	;
	v76 = v65
	v77 = v66
	goto L15
L20:
	;
	v69 = int32(1)
	if v65 == v66 {
		v61 = v61 + v69
		v62 = v62 + v69
		goto L18
	} else {
		goto L21
	}
L21:
	;
	goto L19
L22:
	;
	goto L13
L23:
	;
	goto L3
L24:
	;
	v88 = int32(*(*uint8)(unsafe.Add(mBase, _consts[96])))
	if v88 == int32(1) {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48))) = int32(0)
	v136 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	F_LWLockRelease(m, v136+int32(4736))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L1
	} else {
		goto L47
	}
L26:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v48)+272)) = int64(0)
	v131 = v127
	goto L25
L27:
	;
	if l1 == int32(0) {
		goto L30
	} else {
		goto L31
	}
L28:
	;
	goto L29
L29:
	;
	v116 = *(*int32)(unsafe.Add(mBase, _consts[353]))
	*(*int32)(unsafe.Add(mBase, uint32(v48)+8)) = v116
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
	*(*int32)(unsafe.Add(mBase, uint32(v48))) = int32(1)
	if v118 != 0 {
		goto L42
	} else {
		goto L43
	}
L30:
	;
	F_ConditionVariablePrepareToSleep(m, v48+int32(224))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L1
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
	*(*int32)(unsafe.Add(mBase, uint32(v48))) = int32(1)
	if v97 != 0 {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	goto L32
L34:
	;
	F_s_lock(m, v48, int32(514781), int32(632), int32(381430))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L1
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v48)+8))
	if v105 == int32(0) {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	goto L36
L38:
	;
	v109 = *(*int32)(unsafe.Add(mBase, _consts[353]))
	*(*int32)(unsafe.Add(mBase, uint32(v48)+8)) = v109
	v111 = v109
	goto L40
L39:
	;
	v111 = v105
	goto L40
L40:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v48)+112))
	if v112 == int32(0) {
		v127 = v111
		goto L26
	} else {
		goto L41
	}
L41:
	;
	v131 = v111
	goto L25
L42:
	;
	F_s_lock(m, v48, int32(341339), int32(251), int32(434334))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L1
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v48)+112))
	if v126 != 0 {
		v131 = v116
		goto L25
	} else {
		goto L46
	}
L45:
	;
	goto L44
L46:
	;
	v127 = v116
	goto L26
L47:
	;
	v142 = *(*int32)(unsafe.Add(mBase, _consts[353]))
	if v131 == v142 {
		goto L5
	} else {
		goto L48
	}
L48:
	;
	if l1 == int32(0) {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	F_ConditionVariableSleep(m, v48+int32(224), int32(134217777))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L1
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	goto L7
L52:
	;
	F_ConditionVariableCancelSleep(m)
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	v154 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v158 = F_LWLockAcquire(m, v154+int32(4736), int32(1))
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	v161 = *(*int32)(unsafe.Add(mBase, _consts[678]))
	if int32(0) < v161 {
		v33 = v161
		goto L6
	} else {
		goto L55
	}
L55:
	;
	goto L3
L56:
	;
	F_errcode(m, int32(100663621))
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+68)) = v131
	*(*int32)(unsafe.Add(mBase, uint32(v12)+64)) = v53
	F_errmsg(m, int32(500615), v12-int32(-64))
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	F_errfinish(m, int32(514781), int32(665), int32(381430))
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L60:
	;
	F_ConditionVariableCancelSleep(m)
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L1
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	*(*int32)(unsafe.Add(mBase, _consts[645])) = v48
	if l2 != 0 {
		goto L65
	} else {
		goto L66
	}
L63:
	;
	goto L62
L64:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L1
	} else {
		goto L87
	}
L65:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v48)+112))
	if v189 != 0 {
		goto L64
	} else {
		goto L68
	}
L66:
	;
	goto L67
L67:
	;
	F_ConditionVariableBroadcast(m, v48+int32(224))
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L1
	} else {
		goto L69
	}
L68:
	;
	goto L67
L69:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v48)+88))
	if v194 != 0 {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v198 = *(*int32)(unsafe.Add(mBase, _consts[679]))
	v201 = base.I32_div_s(v48-v198, int32(288))
	goto L73
L71:
	;
	goto L72
L72:
	;
	v208 = int32(*(*uint8)(unsafe.Add(mBase, _consts[680])))
	if v208 != int32(1) {
		goto L75
	} else {
		goto L76
	}
L73:
	;
	v205 = F_pgstat_get_entry_ref(m, int32(4), int32(0), base.I64_extend_i32_s(v201), int32(1), int32(0))
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L1
	} else {
		goto L74
	}
L74:
	;
	goto L72
L75:
	;
	m.G0 = v12 + int32(80)
	return
L76:
	;
	v214 = int32(*(*uint8)(unsafe.Add(mBase, _consts[681])))
	if v214 != 0 {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v215 = int32(15)
	goto L79
L78:
	;
	v215 = int32(14)
	goto L79
L79:
	;
	v217 = F_errstart(m, v215, int32(0))
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L1
	} else {
		goto L80
	}
L80:
	;
	if v217 == int32(0) {
		goto L75
	} else {
		goto L81
	}
L81:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v48)+88))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v53
	if v221 != 0 {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v225 = int32(729632)
	goto L84
L83:
	;
	v225 = int32(729553)
	goto L84
L84:
	;
	F_errmsg(m, v225, v12+int32(16))
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	F_errfinish(m, int32(514781), int32(705), int32(381430))
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L1
	} else {
		goto L86
	}
L86:
	;
	goto L75
L87:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L1
	} else {
		goto L88
	}
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = v53
	F_errmsg(m, int32(729401), v12+int32(48))
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L1
	} else {
		goto L89
	}
L89:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v48)+112))
	v254 = int32(389545)
	if base.Ui32(int32(8)) < base.Ui32(v252) {
		v269 = v254
		goto L91
	} else {
		goto L92
	}
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v269
	F_errdetail(m, int32(694083), v12+int32(32))
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L1
	} else {
		goto L94
	}
L91:
	;
	goto L90
L92:
	;
	if int32(base.Ui32(int32(279))>>(uint(v252)%32))&int32(1) == int32(0) {
		v269 = v254
		goto L91
	} else {
		goto L93
	}
L93:
	;
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v252<<(uint(int32(2))%32))+uint32(_consts[682])))
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v267)+4))
	v269 = v268
	goto L91
L94:
	;
	F_errfinish(m, int32(514781), int32(684), int32(381430))
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L1
	} else {
		goto L95
	}
L95:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L96:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L1
	} else {
		goto L97
	}
L97:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L1
	} else {
		goto L98
	}
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = l0
	F_errmsg(m, int32(76362), v12)
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L1
	} else {
		goto L99
	}
L99:
	;
	F_errfinish(m, int32(514781), int32(610), int32(381430))
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L1
	} else {
		goto L100
	}
L100:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_ReplicationSlotCleanup(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
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
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	v7 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v11 = F_LWLockAcquire(m, v7+int32(4736), int32(1))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v14 = *(*int32)(unsafe.Add(mBase, _consts[678]))
	if v14 <= int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v93 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	F_LWLockRelease(m, v93+int32(4736))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L1
	} else {
		goto L28
	}
L4:
	;
	v19 = v14
	goto L5
L5:
	;
	v24 = *(*int32)(unsafe.Add(mBase, _consts[679]))
	v27 = v19
	v28 = int32(0)
	v29 = v24
	goto L7
L6:
	;
	goto L3
L7:
	;
	v32 = v29 + v28*int32(288)
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+4)))
	if v33 == int32(1) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32))) = int32(0)
	v65 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	F_LWLockRelease(m, v65+int32(4736))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L1
	} else {
		goto L23
	}
L9:
	;
	goto L8
L10:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	*(*int32)(unsafe.Add(mBase, uint32(v32))) = int32(1)
	if v36 != 0 {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	v57 = v27
	v58 = v29
	goto L12
L12:
	;
	v60 = v28 + int32(1)
	if v60 < v57 {
		v27 = v57
		v28 = v60
		v29 = v58
		goto L7
	} else {
		goto L22
	}
L13:
	;
	F_s_lock(m, v32, int32(514781), int32(820), int32(243865))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L1
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v32)+8))
	v46 = *(*int32)(unsafe.Add(mBase, _consts[353]))
	if v44 == v46 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	goto L15
L17:
	;
	if l0 == int32(0) {
		goto L9
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32))) = int32(0)
	v54 = *(*int32)(unsafe.Add(mBase, _consts[679]))
	v56 = *(*int32)(unsafe.Add(mBase, _consts[678]))
	v57 = v56
	v58 = v54
	goto L12
L20:
	;
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+201)))
	if v50 != 0 {
		goto L9
	} else {
		goto L21
	}
L21:
	;
	goto L19
L22:
	;
	goto L3
L23:
	;
	F_ReplicationSlotDropPtr(m, v32)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	F_ConditionVariableBroadcast(m, v32+int32(224))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	v77 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v81 = F_LWLockAcquire(m, v77+int32(4736), int32(1))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	v84 = *(*int32)(unsafe.Add(mBase, _consts[678]))
	if int32(0) < v84 {
		v19 = v84
		goto L5
	} else {
		goto L27
	}
L27:
	;
	goto L6
L28:
	;
	return
}
func F_ReplicationSlotPersist(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v41 int32
	_ = v41
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	v4 = m.G0
	v6 = v4 - int32(1040)
	m.G0 = v6
	v9 = *(*int32)(unsafe.Add(mBase, _consts[645]))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = int32(1)
	if v10 != 0 {
		F_s_lock(m, v9, int32(514781), int32(1125), int32(78898))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return
		} else {
			v18 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v9))) = v18
			*(*int32)(unsafe.Add(mBase, uint32(v9)+92)) = v18
			v23 = *(*int32)(unsafe.Add(mBase, _consts[645]))
			v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
			*(*int32)(unsafe.Add(mBase, uint32(v23))) = int32(1)
			if v24 != 0 {
				F_s_lock(m, v23, int32(514781), int32(1107), int32(8431))
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return
				} else {
					v32 = int32(4464132)
					v33 = *(*int32)(unsafe.Add(mBase, _consts[645]))
					v34 = int32(257)
					*(*uint16)(unsafe.Add(mBase, uint32(v33)+12)) = uint16(v34)
					*(*int32)(unsafe.Add(mBase, uint32(v23))) = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v6))) = int32(90762)
					v41 = *(*int32)(unsafe.Add(mBase, _consts[645]))
					*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v41 + int32(24)
					v48 = F_pg_sprintf(m, v6+int32(16), int32(186923), v6)
					mBase = m.M
					v49 = m.ExcPending
					if v49 != 0 {
						return
					} else {
						v51 = *(*int32)(unsafe.Add(mBase, _consts[645]))
						F_SaveSlotToPath(m, v51, v6+int32(16), int32(21))
						mBase = m.M
						v56 = m.ExcPending
						if v56 != 0 {
							return
						} else {
							m.G0 = v6 + int32(1040)
							return
						}
					}
				}
			} else {
				v32 = int32(4464132)
				v33 = *(*int32)(unsafe.Add(mBase, _consts[645]))
				v34 = int32(257)
				*(*uint16)(unsafe.Add(mBase, uint32(v33)+12)) = uint16(v34)
				*(*int32)(unsafe.Add(mBase, uint32(v23))) = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v6))) = int32(90762)
				v41 = *(*int32)(unsafe.Add(mBase, _consts[645]))
				*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v41 + int32(24)
				v48 = F_pg_sprintf(m, v6+int32(16), int32(186923), v6)
				mBase = m.M
				v49 = m.ExcPending
				if v49 != 0 {
					return
				} else {
					v51 = *(*int32)(unsafe.Add(mBase, _consts[645]))
					F_SaveSlotToPath(m, v51, v6+int32(16), int32(21))
					mBase = m.M
					v56 = m.ExcPending
					if v56 != 0 {
						return
					} else {
						m.G0 = v6 + int32(1040)
						return
					}
				}
			}
		}
	} else {
		v18 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v9))) = v18
		*(*int32)(unsafe.Add(mBase, uint32(v9)+92)) = v18
		v23 = *(*int32)(unsafe.Add(mBase, _consts[645]))
		v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
		*(*int32)(unsafe.Add(mBase, uint32(v23))) = int32(1)
		if v24 != 0 {
			F_s_lock(m, v23, int32(514781), int32(1107), int32(8431))
			mBase = m.M
			v31 = m.ExcPending
			if v31 != 0 {
				return
			} else {
				v32 = int32(4464132)
				v33 = *(*int32)(unsafe.Add(mBase, _consts[645]))
				v34 = int32(257)
				*(*uint16)(unsafe.Add(mBase, uint32(v33)+12)) = uint16(v34)
				*(*int32)(unsafe.Add(mBase, uint32(v23))) = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v6))) = int32(90762)
				v41 = *(*int32)(unsafe.Add(mBase, _consts[645]))
				*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v41 + int32(24)
				v48 = F_pg_sprintf(m, v6+int32(16), int32(186923), v6)
				mBase = m.M
				v49 = m.ExcPending
				if v49 != 0 {
					return
				} else {
					v51 = *(*int32)(unsafe.Add(mBase, _consts[645]))
					F_SaveSlotToPath(m, v51, v6+int32(16), int32(21))
					mBase = m.M
					v56 = m.ExcPending
					if v56 != 0 {
						return
					} else {
						m.G0 = v6 + int32(1040)
						return
					}
				}
			}
		} else {
			v32 = int32(4464132)
			v33 = *(*int32)(unsafe.Add(mBase, _consts[645]))
			v34 = int32(257)
			*(*uint16)(unsafe.Add(mBase, uint32(v33)+12)) = uint16(v34)
			*(*int32)(unsafe.Add(mBase, uint32(v23))) = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v6))) = int32(90762)
			v41 = *(*int32)(unsafe.Add(mBase, _consts[645]))
			*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v41 + int32(24)
			v48 = F_pg_sprintf(m, v6+int32(16), int32(186923), v6)
			mBase = m.M
			v49 = m.ExcPending
			if v49 != 0 {
				return
			} else {
				v51 = *(*int32)(unsafe.Add(mBase, _consts[645]))
				F_SaveSlotToPath(m, v51, v6+int32(16), int32(21))
				mBase = m.M
				v56 = m.ExcPending
				if v56 != 0 {
					return
				} else {
					m.G0 = v6 + int32(1040)
					return
				}
			}
		}
	}
}
func F_ReplicationSlotValidateNameInternal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v87 int32
	_ = v87
	var v96 int32
	_ = v96
	v8 = m.G0
	v10 = v8 - int32(48)
	m.G0 = v10
	v12 = F_strlen(m, l0)
	mBase = m.M
	if v12 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	m.G0 = v10 + int32(48)
	return v96
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v87
	v96 = int32(0)
	goto L1
L3:
	;
	v87 = int32(0)
	goto L2
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(33579140)
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = l0
	v19 = F_psprintf(m, int32(86955), v10)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	if base.Ui32(v12) <= base.Ui32(int32(63)) {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	return int32(0)
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v19
	goto L3
L9:
	;
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v26 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	goto L11
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(34103428)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = l0
	v76 = F_psprintf(m, int32(343425), v10+int32(16))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L7
	} else {
		goto L24
	}
L12:
	;
	v96 = int32(1)
	goto L1
L13:
	;
	goto L14
L14:
	;
	v33 = l0
	v35 = v26
	goto L15
L15:
	;
	if base.Ui32((v35-int32(97))&int32(255)) < base.Ui32(int32(26)) {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v96 = v66
	goto L1
L17:
	;
	v66 = int32(1)
	v68 = v33 + v66
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68))))
	if v69 != 0 {
		v33 = v68
		v35 = v69
		goto L15
	} else {
		goto L23
	}
L18:
	;
	if v35&int32(255) == int32(95) {
		goto L17
	} else {
		goto L19
	}
L19:
	;
	if base.Ui32((v35-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		goto L17
	} else {
		goto L20
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(33579140)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = l0
	v59 = F_psprintf(m, int32(227690), v10+int32(32))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L7
	} else {
		goto L21
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v59
	v64 = F_psprintf(m, int32(635478), int32(0))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L7
	} else {
		goto L22
	}
L22:
	;
	v87 = v64
	goto L2
L23:
	;
	goto L16
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v76
	goto L3
}
func F_replication_yylex(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v160 int32
	_ = v160
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v179 int32
	_ = v179
	var v188 int32
	_ = v188
	var v196 int32
	_ = v196
	var v202 int32
	_ = v202
	var v213 int32
	_ = v213
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v260 int32
	_ = v260
	var v277 int32
	_ = v277
	var v281 int64
	_ = v281
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v299 int64
	_ = v299
	var v300 int64
	_ = v300
	var v308 int32
	_ = v308
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v333 int32
	_ = v333
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v348 int32
	_ = v348
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v361 int32
	_ = v361
	var v363 int32
	_ = v363
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v377 int32
	_ = v377
	var v379 int32
	_ = v379
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v398 int32
	_ = v398
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v415 int32
	_ = v415
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v425 int32
	_ = v425
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v442 int32
	_ = v442
	var v446 int32
	_ = v446
	var v448 int32
	_ = v448
	var v450 int32
	_ = v450
	var v456 int32
	_ = v456
	var v464 int32
	_ = v464
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v473 int32
	_ = v473
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v488 int32
	_ = v488
	var v490 int32
	_ = v490
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v507 int32
	_ = v507
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v526 int32
	_ = v526
	var v535 int32
	_ = v535
	var v543 int32
	_ = v543
	var v547 int32
	_ = v547
	var v549 int32
	_ = v549
	var v551 int32
	_ = v551
	var v565 int32
	_ = v565
	var v568 int32
	_ = v568
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v580 int32
	_ = v580
	var v582 int32
	_ = v582
	var v595 int32
	_ = v595
	var v599 int32
	_ = v599
	var v600 int32
	_ = v600
	var v605 int32
	_ = v605
	var v607 int32
	_ = v607
	var v612 int32
	_ = v612
	var v616 int32
	_ = v616
	var v631 int32
	_ = v631
	var v639 int32
	_ = v639
	var v644 int32
	_ = v644
	var v645 int32
	_ = v645
	var v653 int32
	_ = v653
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v664 int32
	_ = v664
	var v666 int32
	_ = v666
	var v667 int32
	_ = v667
	var v677 int32
	_ = v677
	var v679 int32
	_ = v679
	var v681 int32
	_ = v681
	var v683 int32
	_ = v683
	var v685 int32
	_ = v685
	var v687 int32
	_ = v687
	var v689 int32
	_ = v689
	var v691 int32
	_ = v691
	var v693 int32
	_ = v693
	var v694 int32
	_ = v694
	var v696 int32
	_ = v696
	var v698 int32
	_ = v698
	var v700 int32
	_ = v700
	var v702 int32
	_ = v702
	var v714 int32
	_ = v714
	var v716 int32
	_ = v716
	var v717 int32
	_ = v717
	var v727 int32
	_ = v727
	var v729 int32
	_ = v729
	var v734 int32
	_ = v734
	var v749 int32
	_ = v749
	var v750 int32
	_ = v750
	var v754 int32
	_ = v754
	var v757 int32
	_ = v757
	var v758 int32
	_ = v758
	var v760 int32
	_ = v760
	var v768 int32
	_ = v768
	var v771 int32
	_ = v771
	var v777 int32
	_ = v777
	var v778 int32
	_ = v778
	var v779 int32
	_ = v779
	var v782 int32
	_ = v782
	var v783 int32
	_ = v783
	var v785 int32
	_ = v785
	var v788 int32
	_ = v788
	var v796 int32
	_ = v796
	var v801 int32
	_ = v801
	var v807 int32
	_ = v807
	var v810 int32
	_ = v810
	var v811 int32
	_ = v811
	var v812 int32
	_ = v812
	var v813 int32
	_ = v813
	var v814 int32
	_ = v814
	var v815 int32
	_ = v815
	var v820 int32
	_ = v820
	var v822 int32
	_ = v822
	var v823 int32
	_ = v823
	var v827 int32
	_ = v827
	var v828 int32
	_ = v828
	var v829 int32
	_ = v829
	var v834 int32
	_ = v834
	var v835 int32
	_ = v835
	var v845 int32
	_ = v845
	var v848 int32
	_ = v848
	var v850 int32
	_ = v850
	var v851 int32
	_ = v851
	var v864 int32
	_ = v864
	var v865 int32
	_ = v865
	var v866 int32
	_ = v866
	var v868 int32
	_ = v868
	var v869 int32
	_ = v869
	var v870 int32
	_ = v870
	var v874 int32
	_ = v874
	var v875 int32
	_ = v875
	var v880 int32
	_ = v880
	var v882 int32
	_ = v882
	var v883 int32
	_ = v883
	var v884 int32
	_ = v884
	var v887 int32
	_ = v887
	var v888 int32
	_ = v888
	var v889 int32
	_ = v889
	var v898 int32
	_ = v898
	var v899 int32
	_ = v899
	var v900 int32
	_ = v900
	var v904 int32
	_ = v904
	var v905 int32
	_ = v905
	var v908 int32
	_ = v908
	var v912 int32
	_ = v912
	var v917 int32
	_ = v917
	var v918 int32
	_ = v918
	var v922 int32
	_ = v922
	var v923 int32
	_ = v923
	var v926 int32
	_ = v926
	var v927 int32
	_ = v927
	var v928 int32
	_ = v928
	var v929 int32
	_ = v929
	var v943 int32
	_ = v943
	var v944 int32
	_ = v944
	var v947 int32
	_ = v947
	var v948 int32
	_ = v948
	var v949 int32
	_ = v949
	var v958 int32
	_ = v958
	var v962 int32
	_ = v962
	var v964 int32
	_ = v964
	var v965 int32
	_ = v965
	var v968 int32
	_ = v968
	var v972 int32
	_ = v972
	var v976 int32
	_ = v976
	var v977 int32
	_ = v977
	var v981 int32
	_ = v981
	var v982 int32
	_ = v982
	var v985 int32
	_ = v985
	var v986 int32
	_ = v986
	var v987 int32
	_ = v987
	var v990 int32
	_ = v990
	var v998 int32
	_ = v998
	var v1001 int32
	_ = v1001
	var v1004 int32
	_ = v1004
	var v1005 int32
	_ = v1005
	var v1018 int32
	_ = v1018
	var v1019 int32
	_ = v1019
	var v1023 int32
	_ = v1023
	var v1036 int32
	_ = v1036
	var v1037 int32
	_ = v1037
	var v1040 int32
	_ = v1040
	var v1043 int32
	_ = v1043
	var v1044 int32
	_ = v1044
	var v1045 int32
	_ = v1045
	var v1049 int32
	_ = v1049
	var v1051 int32
	_ = v1051
	var v1052 int32
	_ = v1052
	var v1053 int32
	_ = v1053
	var v1054 int32
	_ = v1054
	var v1055 int32
	_ = v1055
	var v1056 int32
	_ = v1056
	var v1061 int32
	_ = v1061
	var v1063 int32
	_ = v1063
	var v1065 int32
	_ = v1065
	var v1067 int32
	_ = v1067
	var v1068 int32
	_ = v1068
	var v1072 int32
	_ = v1072
	var v1076 int32
	_ = v1076
	var v1078 int32
	_ = v1078
	var v1081 int32
	_ = v1081
	var v1088 int32
	_ = v1088
	var v1090 int32
	_ = v1090
	var v1093 int32
	_ = v1093
	var v1096 int32
	_ = v1096
	var v1097 int32
	_ = v1097
	var v1099 int32
	_ = v1099
	var v1101 int32
	_ = v1101
	var v1102 int32
	_ = v1102
	var v1105 int32
	_ = v1105
	var v1106 int32
	_ = v1106
	var v1108 int32
	_ = v1108
	var v1110 int32
	_ = v1110
	var v1113 int32
	_ = v1113
	var v1117 int32
	_ = v1117
	var v1118 int32
	_ = v1118
	var v1122 int32
	_ = v1122
	var v1130 int32
	_ = v1130
	var v1131 int32
	_ = v1131
	var v1134 int32
	_ = v1134
	var v1135 int32
	_ = v1135
	var v1136 int32
	_ = v1136
	var v1138 int32
	_ = v1138
	var v1139 int32
	_ = v1139
	var v1142 int32
	_ = v1142
	var v1143 int32
	_ = v1143
	var v1145 int32
	_ = v1145
	var v1148 int32
	_ = v1148
	var v1149 int32
	_ = v1149
	var v1150 int32
	_ = v1150
	var v1153 int32
	_ = v1153
	var v1162 int32
	_ = v1162
	var v1163 int32
	_ = v1163
	var v1164 int32
	_ = v1164
	var v1165 int32
	_ = v1165
	var v1166 int32
	_ = v1166
	var v1170 int32
	_ = v1170
	var v1171 int32
	_ = v1171
	var v1175 int32
	_ = v1175
	var v1176 int32
	_ = v1176
	var v1177 int32
	_ = v1177
	var v1178 int32
	_ = v1178
	var v1179 int32
	_ = v1179
	var v1180 int32
	_ = v1180
	var v1181 int32
	_ = v1181
	var v1182 int32
	_ = v1182
	var v1183 int32
	_ = v1183
	var v1184 int32
	_ = v1184
	var v1187 int32
	_ = v1187
	var v1189 int32
	_ = v1189
	var v1190 int32
	_ = v1190
	var v1194 int32
	_ = v1194
	var v1195 int32
	_ = v1195
	var v1201 int32
	_ = v1201
	var v1202 int32
	_ = v1202
	var v1203 int32
	_ = v1203
	var v1205 int32
	_ = v1205
	var v1206 int32
	_ = v1206
	var v1207 int32
	_ = v1207
	var v1209 int32
	_ = v1209
	var v1212 int32
	_ = v1212
	var v1213 int32
	_ = v1213
	var v1215 int32
	_ = v1215
	var v1217 int32
	_ = v1217
	var v1218 int32
	_ = v1218
	var v1222 int32
	_ = v1222
	var v1223 int32
	_ = v1223
	var v1224 int32
	_ = v1224
	var v1228 int32
	_ = v1228
	var v1229 int32
	_ = v1229
	var v1232 int32
	_ = v1232
	var v1233 int32
	_ = v1233
	var v1234 int32
	_ = v1234
	var v1243 int32
	_ = v1243
	var v1245 int32
	_ = v1245
	var v1247 int32
	_ = v1247
	var v1254 int32
	_ = v1254
	var v1261 int32
	_ = v1261
	var v1264 int32
	_ = v1264
	var v1265 int32
	_ = v1265
	var v1270 int32
	_ = v1270
	var v1274 int32
	_ = v1274
	var v1275 int32
	_ = v1275
	var v1279 int32
	_ = v1279
	var v1280 int32
	_ = v1280
	var v1285 int32
	_ = v1285
	var v1287 int32
	_ = v1287
	var v1289 int32
	_ = v1289
	var v1290 int32
	_ = v1290
	var v1304 int32
	_ = v1304
	var v1309 int32
	_ = v1309
	var v1310 int32
	_ = v1310
	var v1312 int32
	_ = v1312
	var v1313 int32
	_ = v1313
	var v1317 int32
	_ = v1317
	var v1318 int32
	_ = v1318
	var v1323 int32
	_ = v1323
	var v1332 int32
	_ = v1332
	var v1340 int32
	_ = v1340
	var v1344 int32
	_ = v1344
	var v1346 int32
	_ = v1346
	var v1348 int32
	_ = v1348
	var v1349 int32
	_ = v1349
	var v1350 int32
	_ = v1350
	var v1351 int32
	_ = v1351
	var v1355 int32
	_ = v1355
	var v1358 int32
	_ = v1358
	var v1364 int32
	_ = v1364
	var v1366 int32
	_ = v1366
	var v1368 int32
	_ = v1368
	var v1375 int32
	_ = v1375
	var v1382 int32
	_ = v1382
	var v1385 int32
	_ = v1385
	var v1386 int32
	_ = v1386
	var v1391 int32
	_ = v1391
	var v1395 int32
	_ = v1395
	var v1396 int32
	_ = v1396
	var v1400 int32
	_ = v1400
	var v1401 int32
	_ = v1401
	var v1406 int32
	_ = v1406
	var v1408 int32
	_ = v1408
	var v1410 int32
	_ = v1410
	var v1411 int32
	_ = v1411
	var v1425 int32
	_ = v1425
	var v1430 int32
	_ = v1430
	var v1431 int32
	_ = v1431
	var v1433 int32
	_ = v1433
	var v1434 int32
	_ = v1434
	var v1438 int32
	_ = v1438
	var v1439 int32
	_ = v1439
	var v1444 int32
	_ = v1444
	var v1453 int32
	_ = v1453
	var v1461 int32
	_ = v1461
	var v1465 int32
	_ = v1465
	var v1467 int32
	_ = v1467
	var v1471 int32
	_ = v1471
	var v1476 int32
	_ = v1476
	var v1481 int32
	_ = v1481
	var v1493 int32
	_ = v1493
	var v1497 int32
	_ = v1497
	var v1502 int32
	_ = v1502
	var v1512 int32
	_ = v1512
	v14 = m.G0
	v16 = v14 - int32(16)
	m.G0 = v16
	*(*int32)(unsafe.Add(mBase, uint32(l1)+92)) = l0
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	if v19 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+40)) = int32(1)
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
	if v24 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v88)))
	if v89 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+44)) = int32(1)
	goto L6
L5:
	;
	goto L6
L6:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v29 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v33 = *(*int32)(unsafe.Add(mBase, _consts[364]))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v33
	goto L9
L8:
	;
	goto L9
L9:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v35 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v39 = *(*int32)(unsafe.Add(mBase, _consts[670]))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v39
	goto L12
L11:
	;
	goto L12
L12:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v41 != 0 {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v71
	v75 = v68 + v69<<(uint(int32(2))%32)
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+80)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = v77
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v81
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77))))
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+24)) = uint8(v83)
	goto L3
L14:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v41+v42<<(uint(int32(2))%32))))
	if v46 != 0 {
		v68 = v41
		v69 = v42
		v70 = v46
		goto L13
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	F_replication_yyensure_buffer_stack(m, l1)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	goto L16
L18:
	;
	return int32(0)
L19:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v54 = F_replication_yy_create_buffer(m, v53, l1)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L18
	} else {
		goto L20
	}
L20:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v58 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v56+v57<<(uint(v58)%32)))) = v54
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v62+v63<<(uint(v58)%32))))
	v68 = v62
	v69 = v63
	v70 = v67
	goto L13
L21:
	;
	m.G0 = v1512 + int32(16)
	return v1502
L22:
	;
	v93 = l1
	v102 = v16
	goto L25
L23:
	;
	goto L24
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v88))) = int32(0)
	v1502 = v89
	v1512 = v16
	goto L21
L25:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v93)+36))
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v93)+24)))
	*(*uint8)(unsafe.Add(mBase, uint32(v105))) = uint8(v106)
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v93)+44))
	v109 = v108
	v113 = v105
	v115 = v105
	goto L27
L27:
	;
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115))))
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122)+uint32(_consts[671]))))
	v127 = v109 << (uint(int32(1)) % 32)
	v130 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v127)+uint32(_consts[672]))))
	if v130 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v93)+68)) = v115
	*(*int32)(unsafe.Add(mBase, uint32(v93)+64)) = v109
	goto L31
L30:
	;
	goto L31
L31:
	;
	v135 = int32(*(*int16)(unsafe.Add(mBase, uint32(v127)+uint32(_consts[673]))))
	v136 = v135 + v125
	v141 = int32(*(*int16)(unsafe.Add(mBase, uint32(v136<<(uint(int32(1))%32))+uint32(_consts[674]))))
	if v141 != v109 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v143 = v109
	v145 = v125
	v146 = v125
	goto L35
L33:
	;
	v188 = v136
	goto L34
L34:
	;
	v196 = int32(1)
	v202 = int32(*(*int16)(unsafe.Add(mBase, uint32(v188<<(uint(v196)%32))+uint32(_consts[675]))))
	if v202 != int32(284) {
		v109 = v202
		v115 = v115 + v196
		goto L27
	} else {
		goto L41
	}
L35:
	;
	v160 = int32(*(*int16)(unsafe.Add(mBase, uint32(v143<<(uint(int32(1))%32))+uint32(_consts[676]))))
	if int32(285) <= v160 {
		goto L37
	} else {
		goto L38
	}
L36:
	;
	v188 = v174
	goto L34
L37:
	;
	v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v145)+uint32(_consts[677]))))
	v166 = v165
	goto L39
L38:
	;
	v166 = v146
	goto L39
L39:
	;
	v168 = v166 & int32(255)
	v169 = int32(1)
	v173 = int32(*(*int16)(unsafe.Add(mBase, uint32(v160<<(uint(v169)%32))+uint32(_consts[673]))))
	v174 = v168 + v173
	v179 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v174<<(uint(v169)%32))+uint32(_consts[674]))))
	if v179 != v160&int32(65535) {
		v143 = v160
		v145 = v168
		v146 = v166
		goto L35
	} else {
		goto L40
	}
L40:
	;
	goto L36
L41:
	;
	v213 = v113
	goto L42
L42:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v93)+64))
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v93)+68))
	v220 = v218
	v226 = v219
	v228 = v213
	goto L44
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v93)+80)) = v228
	*(*int32)(unsafe.Add(mBase, uint32(v93)+32)) = v226 - v228
	v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v226))))
	*(*uint8)(unsafe.Add(mBase, uint32(v93)+24)) = uint8(v236)
	v238 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v226))) = uint8(v238)
	*(*int32)(unsafe.Add(mBase, uint32(v93)+36)) = v226
	v245 = int32(*(*int16)(unsafe.Add(mBase, uint32(v220<<(uint(int32(1))%32))+uint32(_consts[672]))))
	v248 = v245
	goto L46
L46:
	;
	switch v248 {
	case 0:
		goto L94
	case 1:
		goto L62
	case 2:
		goto L61
	case 3:
		goto L60
	case 4:
		goto L59
	case 5:
		goto L58
	case 6:
		v1502 = int32(266)
		v1512 = v102
		goto L21
	case 7:
		goto L93
	case 8:
		goto L92
	case 9:
		goto L91
	case 10:
		goto L90
	case 11:
		goto L89
	case 12:
		goto L88
	case 13:
		goto L87
	case 14:
		goto L86
	case 15:
		goto L85
	case 16:
		goto L84
	case 17:
		goto L83
	case 18:
		goto L82
	case 19:
		goto L81
	case 20:
		goto L80
	case 21:
		goto L79
	case 22:
		goto L25
	case 23:
		goto L78
	case 24:
		goto L77
	case 25:
		goto L76
	case 26:
		goto L75
	case 27:
		goto L74
	case 28:
		goto L73
	case 29:
		goto L72
	case 30:
		goto L71
	case 31:
		goto L70
	case 32:
		goto L69
	case 33:
		goto L68
	case 34:
		goto L65
	case 35:
		goto L64
	case 36:
		goto L66
	case 37, 38:
		goto L67
	default:
		goto L63
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v93)+36)) = v1481
	*(*int32)(unsafe.Add(mBase, uint32(v93)+48)) = int32(0)
	v1493 = *(*int32)(unsafe.Add(mBase, uint32(v93)+44))
	v1497 = base.I32_div_s(v1493-int32(1), int32(2))
	v248 = v1497 + int32(36)
	goto L46
L49:
	;
	F_yy_fatal_error_3(m, int32(32620))
	mBase = m.M
	v1476 = m.ExcPending
	if v1476 != 0 {
		goto L18
	} else {
		goto L308
	}
L50:
	;
	F_yy_fatal_error_3(m, int32(712990))
	mBase = m.M
	v1471 = m.ExcPending
	if v1471 != 0 {
		goto L18
	} else {
		goto L307
	}
L51:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L52:
	;
	v1364 = v1351 + v1358
	*(*int32)(unsafe.Add(mBase, uint32(v93)+36)) = v1364
	v1366 = *(*int32)(unsafe.Add(mBase, uint32(v93)+44))
	if base.Ui32(v1364) <= base.Ui32(v1355) {
		v220 = v1366
		v226 = v1364
		v228 = v1355
		goto L44
	} else {
		goto L288
	}
L53:
	;
	v1037 = *(*int32)(unsafe.Add(mBase, uint32(v1036)))
	*(*int32)(unsafe.Add(mBase, uint32(v1037)+16)) = v1023
	v1040 = *(*int32)(unsafe.Add(mBase, uint32(v93)+28))
	if v1040 != 0 {
		v1162 = int32(0)
		goto L232
	} else {
		goto L233
	}
L54:
	;
	v1018 = *(*int32)(unsafe.Add(mBase, uint32(v93)+20))
	v1019 = *(*int32)(unsafe.Add(mBase, uint32(v93)+12))
	v1023 = v1005
	v1036 = v1018 + v1019<<(uint(int32(2))%32)
	goto L53
L55:
	;
	F_yy_fatal_error_3(m, int32(474011))
	mBase = m.M
	v1004 = m.ExcPending
	if v1004 != 0 {
		goto L18
	} else {
		goto L231
	}
L56:
	;
	F_yy_fatal_error_3(m, int32(469667))
	mBase = m.M
	v1001 = m.ExcPending
	if v1001 != 0 {
		goto L18
	} else {
		goto L230
	}
L57:
	;
	F_replication_yyerror(m, int32(277147))
	mBase = m.M
	v998 = m.ExcPending
	if v998 != 0 {
		goto L18
	} else {
		goto L229
	}
L58:
	;
	v1502 = int32(272)
	v1512 = v102
	goto L21
L59:
	;
	v1502 = int32(265)
	v1512 = v102
	goto L21
L60:
	;
	v1502 = int32(264)
	v1512 = v102
	goto L21
L61:
	;
	v1502 = int32(263)
	v1512 = v102
	goto L21
L62:
	;
	v1502 = int32(262)
	v1512 = v102
	goto L21
L63:
	;
	F_yy_fatal_error_3(m, int32(443050))
	mBase = m.M
	v990 = m.ExcPending
	if v990 != 0 {
		goto L18
	} else {
		goto L228
	}
L64:
	;
	v403 = *(*int32)(unsafe.Add(mBase, uint32(v93)+80))
	v404 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v93)+24)))
	*(*uint8)(unsafe.Add(mBase, uint32(v226))) = uint8(v404)
	v406 = *(*int32)(unsafe.Add(mBase, uint32(v93)+20))
	v407 = *(*int32)(unsafe.Add(mBase, uint32(v93)+12))
	v410 = v406 + v407<<(uint(int32(2))%32)
	v411 = *(*int32)(unsafe.Add(mBase, uint32(v410)))
	v412 = *(*int32)(unsafe.Add(mBase, uint32(v411)+44))
	if v412 == int32(0) {
		goto L107
	} else {
		goto L108
	}
L65:
	;
	F_yy_fatal_error_3(m, int32(473363))
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		goto L18
	} else {
		goto L106
	}
L66:
	;
	v1502 = int32(0)
	v1512 = v102
	goto L21
L67:
	;
	F_replication_yyerror(m, int32(346015))
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L18
	} else {
		goto L105
	}
L68:
	;
	v394 = *(*int32)(unsafe.Add(mBase, uint32(v93)+80))
	v395 = int32(*(*int8)(unsafe.Add(mBase, uint32(v394))))
	v1502 = v395
	v1512 = v102
	goto L21
L69:
	;
	v386 = *(*int32)(unsafe.Add(mBase, uint32(v93)+80))
	v387 = F_strlen(m, v386)
	mBase = m.M
	v389 = F_downcase_truncate_identifier(m, v386, v387, int32(1))
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L18
	} else {
		goto L104
	}
L70:
	;
	v379 = *(*int32)(unsafe.Add(mBase, uint32(v93)))
	v382 = *(*int32)(unsafe.Add(mBase, uint32(v93)+80))
	v383 = *(*int32)(unsafe.Add(mBase, uint32(v93)+32))
	F_appendBinaryStringInfo(m, v379+int32(4), v382, v383)
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L18
	} else {
		goto L103
	}
L71:
	;
	v353 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v93)+24)))
	*(*uint8)(unsafe.Add(mBase, uint32(v226))) = uint8(v353)
	*(*int32)(unsafe.Add(mBase, uint32(v93)+80)) = v228
	v356 = int32(1)
	v357 = v228 + v356
	*(*int32)(unsafe.Add(mBase, uint32(v93)+36)) = v357
	*(*int32)(unsafe.Add(mBase, uint32(v93)+32)) = v356
	v361 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v357))))
	*(*uint8)(unsafe.Add(mBase, uint32(v93)+24)) = uint8(v361)
	v363 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v357))) = uint8(v363)
	*(*int32)(unsafe.Add(mBase, uint32(v93)+44)) = v356
	*(*int32)(unsafe.Add(mBase, uint32(v93)+36)) = v357
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v93)+92))
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v93)))
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v369)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v368))) = v370
	v372 = *(*int32)(unsafe.Add(mBase, uint32(v93)+92))
	v373 = *(*int32)(unsafe.Add(mBase, uint32(v372)))
	v374 = F_strlen(m, v373)
	mBase = m.M
	F_truncate_identifier(m, v373, v374, v356)
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L18
	} else {
		goto L102
	}
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v93)+44)) = int32(3)
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v93)))
	F_initStringInfo(m, v348+int32(4))
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L18
	} else {
		goto L101
	}
L73:
	;
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v93)))
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v93)+80))
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v93)+32))
	F_appendBinaryStringInfo(m, v339+int32(4), v342, v343)
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L18
	} else {
		goto L100
	}
L74:
	;
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v93)))
	F_appendStringInfoChar(m, v333+int32(4), int32(39))
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L18
	} else {
		goto L99
	}
L75:
	;
	v313 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v93)+24)))
	*(*uint8)(unsafe.Add(mBase, uint32(v226))) = uint8(v313)
	*(*int32)(unsafe.Add(mBase, uint32(v93)+80)) = v228
	v316 = int32(1)
	v317 = v228 + v316
	*(*int32)(unsafe.Add(mBase, uint32(v93)+36)) = v317
	*(*int32)(unsafe.Add(mBase, uint32(v93)+32)) = v316
	v321 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v317))))
	*(*uint8)(unsafe.Add(mBase, uint32(v93)+24)) = uint8(v321)
	v323 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v317))) = uint8(v323)
	*(*int32)(unsafe.Add(mBase, uint32(v93)+44)) = v316
	*(*int32)(unsafe.Add(mBase, uint32(v93)+36)) = v317
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v93)+92))
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v93)))
	v330 = *(*int32)(unsafe.Add(mBase, uint32(v329)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v328))) = v330
	v1502 = int32(258)
	v1512 = v102
	goto L21
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v93)+44)) = int32(5)
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v93)))
	F_initStringInfo(m, v308+int32(4))
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L18
	} else {
		goto L98
	}
L77:
	;
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v93)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v102)+4)) = v102 + int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v102))) = v102 + int32(12)
	v294 = F_sscanf(m, v286, int32(539886), v102)
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L18
	} else {
		goto L96
	}
L78:
	;
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v93)+80))
	v281 = F_strtox_2(m, v277, int32(0), int32(10), int64(4294967295))
	mBase = m.M
	goto L95
L79:
	;
	v1502 = int32(282)
	v1512 = v102
	goto L21
L80:
	;
	v1502 = int32(271)
	v1512 = v102
	goto L21
L81:
	;
	v1502 = int32(281)
	v1512 = v102
	goto L21
L82:
	;
	v1502 = int32(280)
	v1512 = v102
	goto L21
L83:
	;
	v1502 = int32(279)
	v1512 = v102
	goto L21
L84:
	;
	v1502 = int32(278)
	v1512 = v102
	goto L21
L85:
	;
	v1502 = int32(277)
	v1512 = v102
	goto L21
L86:
	;
	v1502 = int32(275)
	v1512 = v102
	goto L21
L87:
	;
	v1502 = int32(274)
	v1512 = v102
	goto L21
L88:
	;
	v1502 = int32(276)
	v1512 = v102
	goto L21
L89:
	;
	v1502 = int32(273)
	v1512 = v102
	goto L21
L90:
	;
	v1502 = int32(270)
	v1512 = v102
	goto L21
L91:
	;
	v1502 = int32(269)
	v1512 = v102
	goto L21
L92:
	;
	v1502 = int32(268)
	v1512 = v102
	goto L21
L93:
	;
	v1502 = int32(267)
	v1512 = v102
	goto L21
L94:
	;
	v260 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v93)+24)))
	*(*uint8)(unsafe.Add(mBase, uint32(v226))) = uint8(v260)
	v213 = v228
	goto L42
L95:
	;
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v93)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v283))) = base.I32_wrap_i64(v281)
	v1502 = int32(260)
	v1512 = v102
	goto L21
L96:
	;
	if v294 != int32(2) {
		goto L57
	} else {
		goto L97
	}
L97:
	;
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v93)+92))
	v299 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v102)+8)))
	v300 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v102)+12)))
	*(*int64)(unsafe.Add(mBase, uint32(v298))) = v299 | v300<<(uint(int64(32))%64)
	v1502 = int32(261)
	v1512 = v102
	goto L21
L98:
	;
	goto L25
L99:
	;
	goto L25
L100:
	;
	goto L25
L101:
	;
	goto L25
L102:
	;
	v1502 = int32(259)
	v1512 = v102
	goto L21
L103:
	;
	goto L25
L104:
	;
	v391 = *(*int32)(unsafe.Add(mBase, uint32(v93)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v391))) = v389
	v1502 = int32(259)
	v1512 = v102
	goto L21
L105:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L106:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L107:
	;
	v415 = *(*int32)(unsafe.Add(mBase, uint32(v411)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v93)+28)) = v415
	v417 = *(*int32)(unsafe.Add(mBase, uint32(v410)))
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v93)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v417))) = v418
	v420 = *(*int32)(unsafe.Add(mBase, uint32(v93)+20))
	v421 = *(*int32)(unsafe.Add(mBase, uint32(v93)+12))
	v422 = int32(2)
	v425 = *(*int32)(unsafe.Add(mBase, uint32(v420+v421<<(uint(v422)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v425)+44)) = int32(1)
	v428 = *(*int32)(unsafe.Add(mBase, uint32(v93)+20))
	v429 = *(*int32)(unsafe.Add(mBase, uint32(v93)+12))
	v433 = *(*int32)(unsafe.Add(mBase, uint32(v428+v429<<(uint(v422)%32))))
	v434 = v433
	v435 = v428
	v436 = v429
	goto L109
L108:
	;
	v434 = v411
	v435 = v406
	v436 = v407
	goto L109
L109:
	;
	v437 = *(*int32)(unsafe.Add(mBase, uint32(v93)+36))
	v438 = *(*int32)(unsafe.Add(mBase, uint32(v434)+4))
	v439 = *(*int32)(unsafe.Add(mBase, uint32(v93)+28))
	v440 = v438 + v439
	if base.Ui32(v437) <= base.Ui32(v440) {
		goto L110
	} else {
		goto L111
	}
L110:
	;
	v442 = *(*int32)(unsafe.Add(mBase, uint32(v93)+80))
	v446 = v442 + (v403 ^ int32(-1)) + v226
	*(*int32)(unsafe.Add(mBase, uint32(v93)+36)) = v446
	v448 = *(*int32)(unsafe.Add(mBase, uint32(v93)+44))
	if base.Ui32(v442) < base.Ui32(v446) {
		goto L113
	} else {
		goto L114
	}
L111:
	;
	goto L112
L112:
	;
	if base.Ui32(v440+int32(1)) < base.Ui32(v437) {
		goto L56
	} else {
		goto L145
	}
L113:
	;
	v450 = v448
	v456 = v442
	goto L116
L114:
	;
	v551 = v448
	goto L115
L115:
	;
	v565 = v551 << (uint(int32(1)) % 32)
	v568 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v565)+uint32(_consts[672]))))
	if v568 != 0 {
		goto L134
	} else {
		goto L135
	}
L116:
	;
	v464 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v456))))
	if v464 != 0 {
		goto L118
	} else {
		goto L119
	}
L117:
	;
	v551 = v547
	goto L115
L118:
	;
	v467 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v464)+uint32(_consts[671]))))
	v468 = v467
	goto L120
L119:
	;
	v468 = int32(1)
	goto L120
L120:
	;
	v473 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v450<<(uint(int32(1))%32))+uint32(_consts[672]))))
	if v473 != 0 {
		goto L121
	} else {
		goto L122
	}
L121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v93)+68)) = v456
	*(*int32)(unsafe.Add(mBase, uint32(v93)+64)) = v450
	goto L123
L122:
	;
	goto L123
L123:
	;
	v477 = v468 & int32(255)
	v478 = int32(1)
	v482 = int32(*(*int16)(unsafe.Add(mBase, uint32(v450<<(uint(v478)%32))+uint32(_consts[673]))))
	v483 = v477 + v482
	v488 = int32(*(*int16)(unsafe.Add(mBase, uint32(v483<<(uint(v478)%32))+uint32(_consts[674]))))
	if v488 != v450 {
		goto L124
	} else {
		goto L125
	}
L124:
	;
	v490 = v450
	v492 = v468
	v493 = v477
	goto L127
L125:
	;
	v535 = v483
	goto L126
L126:
	;
	v543 = int32(1)
	v547 = int32(*(*int16)(unsafe.Add(mBase, uint32(v535<<(uint(v543)%32))+uint32(_consts[675]))))
	v549 = v456 + v543
	if v549 != v446 {
		v450 = v547
		v456 = v549
		goto L116
	} else {
		goto L133
	}
L127:
	;
	v507 = int32(*(*int16)(unsafe.Add(mBase, uint32(v490<<(uint(int32(1))%32))+uint32(_consts[676]))))
	if int32(285) <= v507 {
		goto L129
	} else {
		goto L130
	}
L128:
	;
	v535 = v521
	goto L126
L129:
	;
	v512 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v493)+uint32(_consts[677]))))
	v513 = v512
	goto L131
L130:
	;
	v513 = v492
	goto L131
L131:
	;
	v515 = v513 & int32(255)
	v516 = int32(1)
	v520 = int32(*(*int16)(unsafe.Add(mBase, uint32(v507<<(uint(v516)%32))+uint32(_consts[673]))))
	v521 = v515 + v520
	v526 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v521<<(uint(v516)%32))+uint32(_consts[674]))))
	if v526 != v507&int32(65535) {
		v490 = v507
		v492 = v513
		v493 = v515
		goto L127
	} else {
		goto L132
	}
L132:
	;
	goto L128
L133:
	;
	goto L117
L134:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v93)+68)) = v446
	*(*int32)(unsafe.Add(mBase, uint32(v93)+64)) = v551
	goto L136
L135:
	;
	goto L136
L136:
	;
	v573 = int32(*(*int16)(unsafe.Add(mBase, uint32(v565)+uint32(_consts[673]))))
	v574 = int32(1)
	v575 = v573 + v574
	v580 = int32(*(*int16)(unsafe.Add(mBase, uint32(v575<<(uint(v574)%32))+uint32(_consts[674]))))
	if v580 != v551 {
		goto L137
	} else {
		goto L138
	}
L137:
	;
	v582 = v551
	goto L140
L138:
	;
	v616 = v575
	goto L139
L139:
	;
	v631 = int32(*(*int16)(unsafe.Add(mBase, uint32(v616<<(uint(int32(1))%32))+uint32(_consts[675]))))
	if v631 == int32(284) {
		v213 = v442
		goto L42
	} else {
		goto L143
	}
L140:
	;
	v595 = int32(1)
	v599 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v582<<(uint(v595)%32))+uint32(_consts[676]))))
	v600 = base.I32_extend16_s(v599)
	v605 = int32(*(*int16)(unsafe.Add(mBase, uint32(v600<<(uint(v595)%32))+uint32(_consts[673]))))
	v607 = v605 + v595
	v612 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v607<<(uint(v595)%32))+uint32(_consts[674]))))
	if v599 != v612 {
		v582 = v600
		goto L140
	} else {
		goto L142
	}
L141:
	;
	v616 = v607
	goto L139
L142:
	;
	goto L141
L143:
	;
	if v616&int32(2147483647) == int32(0) {
		v213 = v442
		goto L42
	} else {
		goto L144
	}
L144:
	;
	v639 = v446 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v93)+36)) = v639
	v109 = v631
	v113 = v442
	v115 = v639
	goto L27
L145:
	;
	v644 = *(*int32)(unsafe.Add(mBase, uint32(v93)+80))
	v645 = *(*int32)(unsafe.Add(mBase, uint32(v434)+40))
	if v645 == int32(0) {
		goto L146
	} else {
		goto L147
	}
L146:
	;
	if v437-v644 != int32(1) {
		v1351 = v438
		v1355 = v644
		v1358 = v439
		goto L52
	} else {
		goto L149
	}
L147:
	;
	goto L148
L148:
	;
	v653 = v644 ^ int32(-1) + v437
	if v653 != 0 {
		goto L150
	} else {
		goto L151
	}
L149:
	;
	v1481 = v644
	goto L48
L150:
	;
	v654 = int32(7)
	v655 = v653 & v654
	if base.Ui32(v437-v644-int32(2)) < base.Ui32(v654) {
		goto L154
	} else {
		goto L155
	}
L151:
	;
	v757 = v434
	v758 = v435
	v760 = v436
	goto L152
L152:
	;
	v768 = *(*int32)(unsafe.Add(mBase, uint32(v757)+44))
	if v768 == int32(2) {
		goto L166
	} else {
		goto L167
	}
L153:
	;
	if v655 != 0 {
		goto L160
	} else {
		goto L161
	}
L154:
	;
	v700 = v438
	v702 = v644
	goto L153
L155:
	;
	goto L156
L156:
	;
	v664 = v438
	v666 = v644
	v667 = int32(0)
	goto L157
L157:
	;
	v677 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v666))))
	*(*uint8)(unsafe.Add(mBase, uint32(v664))) = uint8(v677)
	v679 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v666)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v664)+1)) = uint8(v679)
	v681 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v666)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v664)+2)) = uint8(v681)
	v683 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v666)+3)))
	*(*uint8)(unsafe.Add(mBase, uint32(v664)+3)) = uint8(v683)
	v685 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v666)+4)))
	*(*uint8)(unsafe.Add(mBase, uint32(v664)+4)) = uint8(v685)
	v687 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v666)+5)))
	*(*uint8)(unsafe.Add(mBase, uint32(v664)+5)) = uint8(v687)
	v689 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v666)+6)))
	*(*uint8)(unsafe.Add(mBase, uint32(v664)+6)) = uint8(v689)
	v691 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v666)+7)))
	*(*uint8)(unsafe.Add(mBase, uint32(v664)+7)) = uint8(v691)
	v693 = int32(8)
	v694 = v664 + v693
	v696 = v666 + v693
	v698 = v667 + v693
	if v698 != v653&int32(-8) {
		v664 = v694
		v666 = v696
		v667 = v698
		goto L157
	} else {
		goto L159
	}
L158:
	;
	v700 = v694
	v702 = v696
	goto L153
L159:
	;
	goto L158
L160:
	;
	v714 = v700
	v716 = v702
	v717 = int32(0)
	goto L163
L161:
	;
	goto L162
L162:
	;
	v749 = *(*int32)(unsafe.Add(mBase, uint32(v93)+20))
	v750 = *(*int32)(unsafe.Add(mBase, uint32(v93)+12))
	v754 = *(*int32)(unsafe.Add(mBase, uint32(v749+v750<<(uint(int32(2))%32))))
	v757 = v754
	v758 = v749
	v760 = v750
	goto L152
L163:
	;
	v727 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v716))))
	*(*uint8)(unsafe.Add(mBase, uint32(v714))) = uint8(v727)
	v729 = int32(1)
	v734 = v717 + v729
	if v734 != v655 {
		v714 = v714 + v729
		v716 = v716 + v729
		v717 = v734
		goto L163
	} else {
		goto L165
	}
L164:
	;
	goto L162
L165:
	;
	goto L164
L166:
	;
	v771 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v93)+28)) = v771
	v1023 = v771
	v1036 = v758 + v760<<(uint(int32(2))%32)
	goto L53
L167:
	;
	goto L168
L168:
	;
	v777 = *(*int32)(unsafe.Add(mBase, uint32(v757)+12))
	v778 = v644 - v437
	v779 = v777 + v778
	if v779 == int32(0) {
		goto L169
	} else {
		goto L170
	}
L169:
	;
	v782 = *(*int32)(unsafe.Add(mBase, uint32(v93)+36))
	v783 = v777
	v785 = v757
	v788 = v782
	goto L172
L170:
	;
	v834 = v757
	v835 = v779
	goto L171
L171:
	;
	v845 = int32(8192)
	if base.Ui32(v845) <= base.Ui32(v835) {
		goto L188
	} else {
		goto L189
	}
L172:
	;
	v796 = *(*int32)(unsafe.Add(mBase, uint32(v785)+20))
	if v796 == int32(0) {
		goto L174
	} else {
		goto L175
	}
L173:
	;
	v834 = v827
	v835 = v829
	goto L171
L174:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v785)+4)) = int32(0)
	goto L49
L175:
	;
	goto L176
L176:
	;
	v801 = *(*int32)(unsafe.Add(mBase, uint32(v785)+4))
	if base.Ui32(int32(2147483646)) <= base.Ui32(v783) {
		goto L177
	} else {
		goto L178
	}
L177:
	;
	v807 = int32(-3)
	goto L179
L178:
	;
	v807 = v783 << (uint(int32(1)) % 32)
	goto L179
L179:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v785)+12)) = v807
	v810 = v807 + int32(2)
	if v801 != 0 {
		goto L181
	} else {
		goto L182
	}
L180:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v785)+4)) = v815
	if v815 == int32(0) {
		goto L49
	} else {
		goto L186
	}
L181:
	;
	v811 = F_repalloc(m, v801, v810)
	mBase = m.M
	v812 = m.ExcPending
	if v812 != 0 {
		goto L18
	} else {
		goto L184
	}
L182:
	;
	goto L183
L183:
	;
	v813 = F_palloc(m, v810)
	mBase = m.M
	v814 = m.ExcPending
	if v814 != 0 {
		goto L18
	} else {
		goto L185
	}
L184:
	;
	v815 = v811
	goto L180
L185:
	;
	v815 = v813
	goto L180
L186:
	;
	v820 = v815 + (v788 - v801)
	*(*int32)(unsafe.Add(mBase, uint32(v93)+36)) = v820
	v822 = *(*int32)(unsafe.Add(mBase, uint32(v93)+20))
	v823 = *(*int32)(unsafe.Add(mBase, uint32(v93)+12))
	v827 = *(*int32)(unsafe.Add(mBase, uint32(v822+v823<<(uint(int32(2))%32))))
	v828 = *(*int32)(unsafe.Add(mBase, uint32(v827)+12))
	v829 = v828 + v778
	if v829 == int32(0) {
		v783 = v828
		v785 = v827
		v788 = v820
		goto L172
	} else {
		goto L187
	}
L187:
	;
	goto L173
L188:
	;
	v848 = v845
	goto L190
L189:
	;
	v848 = v835
	goto L190
L190:
	;
	v850 = *(*int32)(unsafe.Add(mBase, uint32(v834)+24))
	if v850 != 0 {
		goto L191
	} else {
		goto L192
	}
L191:
	;
	v851 = int32(0)
	goto L195
L192:
	;
	goto L193
L193:
	;
	*(*int32)(unsafe.Add(mBase, _consts[158])) = int32(0)
	v917 = *(*int32)(unsafe.Add(mBase, uint32(v93)+20))
	v918 = *(*int32)(unsafe.Add(mBase, uint32(v93)+12))
	v922 = *(*int32)(unsafe.Add(mBase, uint32(v917+v918<<(uint(int32(2))%32))))
	v923 = *(*int32)(unsafe.Add(mBase, uint32(v922)+4))
	v926 = *(*int32)(unsafe.Add(mBase, uint32(v93)+4))
	v927 = F_fread(m, v923+v653, int32(1), v848, v926)
	mBase = m.M
	v928 = m.ExcPending
	if v928 != 0 {
		goto L18
	} else {
		goto L210
	}
L194:
	;
	switch v868 {
	case 0:
		goto L202
	default:
		v912 = v882
		goto L200
	case 11:
		goto L201
	}
L195:
	;
	v864 = *(*int32)(unsafe.Add(mBase, uint32(v93)+4))
	v865 = F_do_getc(m, v864)
	mBase = m.M
	v866 = m.ExcPending
	if v866 != 0 {
		goto L18
	} else {
		goto L198
	}
L196:
	;
	v882 = v848
	goto L194
L197:
	;
	v869 = *(*int32)(unsafe.Add(mBase, uint32(v93)+20))
	v870 = *(*int32)(unsafe.Add(mBase, uint32(v93)+12))
	v874 = *(*int32)(unsafe.Add(mBase, uint32(v869+v870<<(uint(int32(2))%32))))
	v875 = *(*int32)(unsafe.Add(mBase, uint32(v874)+4))
	*(*uint8)(unsafe.Add(mBase, uint32(v875+v653+v851))) = uint8(v865)
	v880 = v851 + int32(1)
	if v880 != v848 {
		v851 = v880
		goto L195
	} else {
		goto L199
	}
L198:
	;
	v868 = v865 + int32(1)
	switch v868 {
	case 0, 11:
		v882 = v851
		goto L194
	default:
		goto L197
	}
L199:
	;
	goto L196
L200:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v93)+28)) = v912
	v1005 = v912
	goto L54
L201:
	;
	v899 = *(*int32)(unsafe.Add(mBase, uint32(v93)+20))
	v900 = *(*int32)(unsafe.Add(mBase, uint32(v93)+12))
	v904 = *(*int32)(unsafe.Add(mBase, uint32(v899+v900<<(uint(int32(2))%32))))
	v905 = *(*int32)(unsafe.Add(mBase, uint32(v904)+4))
	v908 = int32(10)
	*(*uint8)(unsafe.Add(mBase, uint32(v905+v653+v882))) = uint8(v908)
	v912 = v882 + int32(1)
	goto L200
L202:
	;
	v883 = *(*int32)(unsafe.Add(mBase, uint32(v93)+4))
	v884 = *(*int32)(unsafe.Add(mBase, uint32(v883)+76))
	if v884 < int32(0) {
		goto L205
	} else {
		goto L206
	}
L203:
	;
	if int32(base.Ui32(v889)>>(uint(int32(5))%32))&int32(1) == int32(0) {
		v912 = v882
		goto L200
	} else {
		goto L208
	}
L204:
	;
	goto L203
L205:
	;
	v887 = *(*int32)(unsafe.Add(mBase, uint32(v883)))
	v889 = v887
	goto L204
L206:
	;
	goto L207
L207:
	;
	v888 = *(*int32)(unsafe.Add(mBase, uint32(v883)))
	v889 = v888
	goto L204
L208:
	;
	F_yy_fatal_error_3(m, int32(474011))
	mBase = m.M
	v898 = m.ExcPending
	if v898 != 0 {
		goto L18
	} else {
		goto L209
	}
L209:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L210:
	;
	v929 = v927
	goto L211
L211:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v93)+28)) = v929
	if v929 != 0 {
		v1005 = v929
		goto L54
	} else {
		goto L213
	}
L213:
	;
	v943 = *(*int32)(unsafe.Add(mBase, uint32(v93)+4))
	v944 = *(*int32)(unsafe.Add(mBase, uint32(v943)+76))
	if v944 < int32(0) {
		goto L216
	} else {
		goto L217
	}
L214:
	;
	if int32(base.Ui32(v949)>>(uint(int32(5))%32))&int32(1) == int32(0) {
		goto L219
	} else {
		goto L220
	}
L215:
	;
	goto L214
L216:
	;
	v947 = *(*int32)(unsafe.Add(mBase, uint32(v943)))
	v949 = v947
	goto L215
L217:
	;
	goto L218
L218:
	;
	v948 = *(*int32)(unsafe.Add(mBase, uint32(v943)))
	v949 = v948
	goto L215
L219:
	;
	v1005 = int32(0)
	goto L54
L220:
	;
	goto L221
L221:
	;
	v958 = *(*int32)(unsafe.Add(mBase, _consts[158]))
	if v958 != int32(27) {
		goto L55
	} else {
		goto L222
	}
L222:
	;
	v962 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[158])) = v962
	v964 = *(*int32)(unsafe.Add(mBase, uint32(v93)+4))
	v965 = *(*int32)(unsafe.Add(mBase, uint32(v964)+76))
	if v962 <= v965 {
		goto L224
	} else {
		goto L225
	}
L223:
	;
	v976 = *(*int32)(unsafe.Add(mBase, uint32(v93)+20))
	v977 = *(*int32)(unsafe.Add(mBase, uint32(v93)+12))
	v981 = *(*int32)(unsafe.Add(mBase, uint32(v976+v977<<(uint(int32(2))%32))))
	v982 = *(*int32)(unsafe.Add(mBase, uint32(v981)+4))
	v985 = *(*int32)(unsafe.Add(mBase, uint32(v93)+4))
	v986 = F_fread(m, v982+v653, int32(1), v848, v985)
	mBase = m.M
	v987 = m.ExcPending
	if v987 != 0 {
		goto L18
	} else {
		goto L227
	}
L224:
	;
	v968 = *(*int32)(unsafe.Add(mBase, uint32(v964)))
	*(*int32)(unsafe.Add(mBase, uint32(v964))) = v968 & int32(-49)
	goto L223
L225:
	;
	goto L226
L226:
	;
	v972 = *(*int32)(unsafe.Add(mBase, uint32(v964)))
	*(*int32)(unsafe.Add(mBase, uint32(v964))) = v972 & int32(-49)
	goto L223
L227:
	;
	v929 = v986
	goto L211
L228:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L229:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L230:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L231:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L232:
	;
	v1163 = *(*int32)(unsafe.Add(mBase, uint32(v93)+28))
	v1164 = v1163 + v653
	v1165 = *(*int32)(unsafe.Add(mBase, uint32(v93)+20))
	v1166 = *(*int32)(unsafe.Add(mBase, uint32(v93)+12))
	v1170 = *(*int32)(unsafe.Add(mBase, uint32(v1165+v1166<<(uint(int32(2))%32))))
	v1171 = *(*int32)(unsafe.Add(mBase, uint32(v1170)+12))
	if base.Ui32(v1171) < base.Ui32(v1164) {
		goto L256
	} else {
		goto L257
	}
L233:
	;
	if v653 == int32(0) {
		goto L234
	} else {
		goto L235
	}
L234:
	;
	v1043 = *(*int32)(unsafe.Add(mBase, uint32(v93)+4))
	v1044 = *(*int32)(unsafe.Add(mBase, uint32(v93)+20))
	if v1044 != 0 {
		goto L239
	} else {
		goto L240
	}
L235:
	;
	goto L236
L236:
	;
	v1148 = *(*int32)(unsafe.Add(mBase, uint32(v93)+20))
	v1149 = *(*int32)(unsafe.Add(mBase, uint32(v93)+12))
	v1150 = int32(2)
	v1153 = *(*int32)(unsafe.Add(mBase, uint32(v1148+v1149<<(uint(v1150)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v1153)+44)) = v1150
	v1162 = v1150
	goto L232
L237:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1110)+40)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1110))) = v1043
	v1117 = *(*int32)(unsafe.Add(mBase, uint32(v93)+20))
	if v1117 != 0 {
		goto L252
	} else {
		goto L253
	}
L238:
	;
	v1067 = *(*int32)(unsafe.Add(mBase, _consts[158]))
	v1068 = *(*int32)(unsafe.Add(mBase, uint32(v93)+12))
	v1072 = *(*int32)(unsafe.Add(mBase, uint32(v1065+v1068<<(uint(int32(2))%32))))
	if v1072 == int32(0) {
		goto L246
	} else {
		goto L247
	}
L239:
	;
	v1045 = *(*int32)(unsafe.Add(mBase, uint32(v93)+12))
	v1049 = *(*int32)(unsafe.Add(mBase, uint32(v1044+v1045<<(uint(int32(2))%32))))
	if v1049 != 0 {
		v1065 = v1044
		goto L238
	} else {
		goto L242
	}
L240:
	;
	goto L241
L241:
	;
	F_replication_yyensure_buffer_stack(m, v93)
	mBase = m.M
	v1051 = m.ExcPending
	if v1051 != 0 {
		goto L18
	} else {
		goto L243
	}
L242:
	;
	goto L241
L243:
	;
	v1052 = *(*int32)(unsafe.Add(mBase, uint32(v93)+4))
	v1053 = F_replication_yy_create_buffer(m, v1052, v93)
	mBase = m.M
	v1054 = m.ExcPending
	if v1054 != 0 {
		goto L18
	} else {
		goto L244
	}
L244:
	;
	v1055 = *(*int32)(unsafe.Add(mBase, uint32(v93)+20))
	v1056 = *(*int32)(unsafe.Add(mBase, uint32(v93)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v1055+v1056<<(uint(int32(2))%32)))) = v1053
	v1061 = *(*int32)(unsafe.Add(mBase, uint32(v93)+20))
	if v1061 != 0 {
		v1065 = v1061
		goto L238
	} else {
		goto L245
	}
L245:
	;
	v1063 = *(*int32)(unsafe.Add(mBase, _consts[158]))
	v1110 = int32(0)
	v1113 = v1063
	goto L237
L246:
	;
	v1110 = int32(0)
	v1113 = v1067
	goto L237
L247:
	;
	goto L248
L248:
	;
	v1076 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1072)+16)) = v1076
	v1078 = *(*int32)(unsafe.Add(mBase, uint32(v1072)+4))
	*(*uint8)(unsafe.Add(mBase, uint32(v1078))) = uint8(v1076)
	v1081 = *(*int32)(unsafe.Add(mBase, uint32(v1072)+4))
	*(*uint8)(unsafe.Add(mBase, uint32(v1081)+1)) = uint8(v1076)
	*(*int32)(unsafe.Add(mBase, uint32(v1072)+44)) = v1076
	*(*int32)(unsafe.Add(mBase, uint32(v1072)+28)) = int32(1)
	v1088 = *(*int32)(unsafe.Add(mBase, uint32(v1072)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1072)+8)) = v1088
	v1090 = *(*int32)(unsafe.Add(mBase, uint32(v93)+20))
	if v1090 == v1076 {
		v1110 = v1072
		v1113 = v1067
		goto L237
	} else {
		goto L249
	}
L249:
	;
	v1093 = *(*int32)(unsafe.Add(mBase, uint32(v93)+12))
	v1096 = v1090 + v1093<<(uint(int32(2))%32)
	v1097 = *(*int32)(unsafe.Add(mBase, uint32(v1096)))
	if v1072 != v1097 {
		v1110 = v1072
		v1113 = v1067
		goto L237
	} else {
		goto L250
	}
L250:
	;
	v1099 = *(*int32)(unsafe.Add(mBase, uint32(v1097)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v93)+28)) = v1099
	v1101 = *(*int32)(unsafe.Add(mBase, uint32(v1096)))
	v1102 = *(*int32)(unsafe.Add(mBase, uint32(v1101)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v93)+80)) = v1102
	*(*int32)(unsafe.Add(mBase, uint32(v93)+36)) = v1102
	v1105 = *(*int32)(unsafe.Add(mBase, uint32(v1096)))
	v1106 = *(*int32)(unsafe.Add(mBase, uint32(v1105)))
	*(*int32)(unsafe.Add(mBase, uint32(v93)+4)) = v1106
	v1108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1102))))
	*(*uint8)(unsafe.Add(mBase, uint32(v93)+24)) = uint8(v1108)
	v1110 = v1072
	v1113 = v1067
	goto L237
L251:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1110)+24)) = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[158])) = v1113
	v1130 = *(*int32)(unsafe.Add(mBase, uint32(v93)+20))
	v1131 = *(*int32)(unsafe.Add(mBase, uint32(v93)+12))
	v1134 = v1130 + v1131<<(uint(int32(2))%32)
	v1135 = *(*int32)(unsafe.Add(mBase, uint32(v1134)))
	v1136 = *(*int32)(unsafe.Add(mBase, uint32(v1135)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v93)+28)) = v1136
	v1138 = *(*int32)(unsafe.Add(mBase, uint32(v1134)))
	v1139 = *(*int32)(unsafe.Add(mBase, uint32(v1138)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v93)+36)) = v1139
	*(*int32)(unsafe.Add(mBase, uint32(v93)+80)) = v1139
	v1142 = *(*int32)(unsafe.Add(mBase, uint32(v1134)))
	v1143 = *(*int32)(unsafe.Add(mBase, uint32(v1142)))
	*(*int32)(unsafe.Add(mBase, uint32(v93)+4)) = v1143
	v1145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1139))))
	*(*uint8)(unsafe.Add(mBase, uint32(v93)+24)) = uint8(v1145)
	v1162 = int32(1)
	goto L232
L252:
	;
	v1118 = *(*int32)(unsafe.Add(mBase, uint32(v93)+12))
	v1122 = *(*int32)(unsafe.Add(mBase, uint32(v1117+v1118<<(uint(int32(2))%32))))
	if v1110 == v1122 {
		goto L251
	} else {
		goto L255
	}
L253:
	;
	goto L254
L254:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1110)+32)) = int64(1)
	goto L251
L255:
	;
	goto L254
L256:
	;
	v1175 = v1164 + int32(base.Ui32(v1163)>>(uint(int32(1))%32))
	v1176 = *(*int32)(unsafe.Add(mBase, uint32(v1170)+4))
	if v1176 != 0 {
		goto L260
	} else {
		goto L261
	}
L257:
	;
	v1205 = v1164
	v1206 = v1165
	v1207 = v1166
	goto L258
L258:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v93)+28)) = v1205
	v1209 = int32(2)
	v1212 = *(*int32)(unsafe.Add(mBase, uint32(v1206+v1207<<(uint(v1209)%32))))
	v1213 = *(*int32)(unsafe.Add(mBase, uint32(v1212)+4))
	v1215 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1213+v1205))) = uint8(v1215)
	v1217 = *(*int32)(unsafe.Add(mBase, uint32(v93)+20))
	v1218 = *(*int32)(unsafe.Add(mBase, uint32(v93)+12))
	v1222 = *(*int32)(unsafe.Add(mBase, uint32(v1217+v1218<<(uint(v1209)%32))))
	v1223 = *(*int32)(unsafe.Add(mBase, uint32(v1222)+4))
	v1224 = *(*int32)(unsafe.Add(mBase, uint32(v93)+28))
	*(*uint8)(unsafe.Add(mBase, uint32(v1223+v1224)+1)) = uint8(v1215)
	v1228 = *(*int32)(unsafe.Add(mBase, uint32(v93)+20))
	v1229 = *(*int32)(unsafe.Add(mBase, uint32(v93)+12))
	v1232 = v1228 + v1229<<(uint(v1209)%32)
	v1233 = *(*int32)(unsafe.Add(mBase, uint32(v1232)))
	v1234 = *(*int32)(unsafe.Add(mBase, uint32(v1233)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v93)+80)) = v1234
	if v1162 == int32(1) {
		v1481 = v1234
		goto L48
	} else {
		goto L266
	}
L259:
	;
	v1182 = *(*int32)(unsafe.Add(mBase, uint32(v93)+20))
	v1183 = *(*int32)(unsafe.Add(mBase, uint32(v93)+12))
	v1184 = int32(2)
	v1187 = *(*int32)(unsafe.Add(mBase, uint32(v1182+v1183<<(uint(v1184)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v1187)+4)) = v1181
	v1189 = *(*int32)(unsafe.Add(mBase, uint32(v93)+20))
	v1190 = *(*int32)(unsafe.Add(mBase, uint32(v93)+12))
	v1194 = *(*int32)(unsafe.Add(mBase, uint32(v1189+v1190<<(uint(v1184)%32))))
	v1195 = *(*int32)(unsafe.Add(mBase, uint32(v1194)+4))
	if v1195 == int32(0) {
		goto L50
	} else {
		goto L265
	}
L260:
	;
	v1177 = F_repalloc(m, v1176, v1175)
	mBase = m.M
	v1178 = m.ExcPending
	if v1178 != 0 {
		goto L18
	} else {
		goto L263
	}
L261:
	;
	goto L262
L262:
	;
	v1179 = F_palloc(m, v1175)
	mBase = m.M
	v1180 = m.ExcPending
	if v1180 != 0 {
		goto L18
	} else {
		goto L264
	}
L263:
	;
	v1181 = v1177
	goto L259
L264:
	;
	v1181 = v1179
	goto L259
L265:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1194)+12)) = v1175 - int32(2)
	v1201 = *(*int32)(unsafe.Add(mBase, uint32(v93)+12))
	v1202 = *(*int32)(unsafe.Add(mBase, uint32(v93)+20))
	v1203 = *(*int32)(unsafe.Add(mBase, uint32(v93)+28))
	v1205 = v1203 + v653
	v1206 = v1202
	v1207 = v1201
	goto L258
L266:
	;
	switch v1162 - int32(1) {
	case 0:
		goto L51
	case 1:
		goto L267
	default:
		goto L268
	}
L267:
	;
	v1348 = *(*int32)(unsafe.Add(mBase, uint32(v93)+28))
	v1349 = *(*int32)(unsafe.Add(mBase, uint32(v1232)))
	v1350 = *(*int32)(unsafe.Add(mBase, uint32(v1349)+4))
	v1351 = v1350
	v1355 = v1234
	v1358 = v1348
	goto L52
L268:
	;
	v1243 = v1234 + (v403 ^ int32(-1)) + v226
	*(*int32)(unsafe.Add(mBase, uint32(v93)+36)) = v1243
	v1245 = *(*int32)(unsafe.Add(mBase, uint32(v93)+44))
	if base.Ui32(v1243) <= base.Ui32(v1234) {
		v109 = v1245
		v113 = v1234
		v115 = v1243
		goto L27
	} else {
		goto L269
	}
L269:
	;
	v1247 = v1245
	v1254 = v1234
	goto L270
L270:
	;
	v1261 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1254))))
	if v1261 != 0 {
		goto L272
	} else {
		goto L273
	}
L271:
	;
	v109 = v1344
	v113 = v1234
	v115 = v1243
	goto L27
L272:
	;
	v1264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1261)+uint32(_consts[671]))))
	v1265 = v1264
	goto L274
L273:
	;
	v1265 = int32(1)
	goto L274
L274:
	;
	v1270 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1247<<(uint(int32(1))%32))+uint32(_consts[672]))))
	if v1270 != 0 {
		goto L275
	} else {
		goto L276
	}
L275:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v93)+68)) = v1254
	*(*int32)(unsafe.Add(mBase, uint32(v93)+64)) = v1247
	goto L277
L276:
	;
	goto L277
L277:
	;
	v1274 = v1265 & int32(255)
	v1275 = int32(1)
	v1279 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1247<<(uint(v1275)%32))+uint32(_consts[673]))))
	v1280 = v1274 + v1279
	v1285 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1280<<(uint(v1275)%32))+uint32(_consts[674]))))
	if v1285 != v1247 {
		goto L278
	} else {
		goto L279
	}
L278:
	;
	v1287 = v1247
	v1289 = v1265
	v1290 = v1274
	goto L281
L279:
	;
	v1332 = v1280
	goto L280
L280:
	;
	v1340 = int32(1)
	v1344 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1332<<(uint(v1340)%32))+uint32(_consts[675]))))
	v1346 = v1254 + v1340
	if v1243 != v1346 {
		v1247 = v1344
		v1254 = v1346
		goto L270
	} else {
		goto L287
	}
L281:
	;
	v1304 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1287<<(uint(int32(1))%32))+uint32(_consts[676]))))
	if int32(285) <= v1304 {
		goto L283
	} else {
		goto L284
	}
L282:
	;
	v1332 = v1318
	goto L280
L283:
	;
	v1309 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1290)+uint32(_consts[677]))))
	v1310 = v1309
	goto L285
L284:
	;
	v1310 = v1289
	goto L285
L285:
	;
	v1312 = v1310 & int32(255)
	v1313 = int32(1)
	v1317 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1304<<(uint(v1313)%32))+uint32(_consts[673]))))
	v1318 = v1312 + v1317
	v1323 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1318<<(uint(v1313)%32))+uint32(_consts[674]))))
	if v1323 != v1304&int32(65535) {
		v1287 = v1304
		v1289 = v1310
		v1290 = v1312
		goto L281
	} else {
		goto L286
	}
L286:
	;
	goto L282
L287:
	;
	goto L271
L288:
	;
	v1368 = v1366
	v1375 = v1355
	goto L289
L289:
	;
	v1382 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1375))))
	if v1382 != 0 {
		goto L291
	} else {
		goto L292
	}
L290:
	;
	v220 = v1465
	v226 = v1364
	v228 = v1355
	goto L44
L291:
	;
	v1385 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1382)+uint32(_consts[671]))))
	v1386 = v1385
	goto L293
L292:
	;
	v1386 = int32(1)
	goto L293
L293:
	;
	v1391 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1368<<(uint(int32(1))%32))+uint32(_consts[672]))))
	if v1391 != 0 {
		goto L294
	} else {
		goto L295
	}
L294:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v93)+68)) = v1375
	*(*int32)(unsafe.Add(mBase, uint32(v93)+64)) = v1368
	goto L296
L295:
	;
	goto L296
L296:
	;
	v1395 = v1386 & int32(255)
	v1396 = int32(1)
	v1400 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1368<<(uint(v1396)%32))+uint32(_consts[673]))))
	v1401 = v1395 + v1400
	v1406 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1401<<(uint(v1396)%32))+uint32(_consts[674]))))
	if v1406 != v1368 {
		goto L297
	} else {
		goto L298
	}
L297:
	;
	v1408 = v1368
	v1410 = v1386
	v1411 = v1395
	goto L300
L298:
	;
	v1453 = v1401
	goto L299
L299:
	;
	v1461 = int32(1)
	v1465 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1453<<(uint(v1461)%32))+uint32(_consts[675]))))
	v1467 = v1375 + v1461
	if v1467 != v1364 {
		v1368 = v1465
		v1375 = v1467
		goto L289
	} else {
		goto L306
	}
L300:
	;
	v1425 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1408<<(uint(int32(1))%32))+uint32(_consts[676]))))
	if int32(285) <= v1425 {
		goto L302
	} else {
		goto L303
	}
L301:
	;
	v1453 = v1439
	goto L299
L302:
	;
	v1430 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1411)+uint32(_consts[677]))))
	v1431 = v1430
	goto L304
L303:
	;
	v1431 = v1410
	goto L304
L304:
	;
	v1433 = v1431 & int32(255)
	v1434 = int32(1)
	v1438 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1425<<(uint(v1434)%32))+uint32(_consts[673]))))
	v1439 = v1433 + v1438
	v1444 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1439<<(uint(v1434)%32))+uint32(_consts[674]))))
	if v1444 != v1425&int32(65535) {
		v1408 = v1425
		v1410 = v1431
		v1411 = v1433
		goto L300
	} else {
		goto L305
	}
L305:
	;
	goto L301
L306:
	;
	goto L290
L307:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L308:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
