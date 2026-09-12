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
	v4 = *(*int32)(unsafe.Add(mBase, _consts[646]))
	if v4 != 0 {
		v6 = *(*int32)(unsafe.Add(mBase, _consts[2]))
		v10 = F_LWLockAcquire(m, v6+int32(5120), int32(0))
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, _consts[646]))
			v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+24))
			v16 = *(*int32)(unsafe.Add(mBase, _consts[350]))
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
				*(*int32)(unsafe.Add(mBase, _consts[646])) = v24
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
	v22 = *(*int32)(unsafe.Add(mBase, _consts[676]))
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
	v36 = *(*int32)(unsafe.Add(mBase, _consts[677]))
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
	v88 = int32(*(*uint8)(unsafe.Add(mBase, _consts[92])))
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
	v116 = *(*int32)(unsafe.Add(mBase, _consts[350]))
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
	F_s_lock(m, v48, int32(473983), int32(632), int32(350372))
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
	v109 = *(*int32)(unsafe.Add(mBase, _consts[350]))
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
	F_s_lock(m, v48, int32(313054), int32(251), int32(399609))
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
	v142 = *(*int32)(unsafe.Add(mBase, _consts[350]))
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
	v161 = *(*int32)(unsafe.Add(mBase, _consts[676]))
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
	F_errmsg(m, int32(461186), v12-int32(-64))
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	F_errfinish(m, int32(473983), int32(665), int32(350372))
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
	*(*int32)(unsafe.Add(mBase, _consts[643])) = v48
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
	v198 = *(*int32)(unsafe.Add(mBase, _consts[677]))
	v201 = base.I32_div_s(v48-v198, int32(288))
	goto L73
L71:
	;
	goto L72
L72:
	;
	v208 = int32(*(*uint8)(unsafe.Add(mBase, _consts[678])))
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
	v214 = int32(*(*uint8)(unsafe.Add(mBase, _consts[679])))
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
	v225 = int32(660599)
	goto L84
L83:
	;
	v225 = int32(660520)
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
	F_errfinish(m, int32(473983), int32(705), int32(350372))
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
	F_errmsg(m, int32(660368), v12+int32(48))
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
	v254 = int32(358077)
	if base.Ui32(int32(8)) < base.Ui32(v252) {
		v269 = v254
		goto L91
	} else {
		goto L92
	}
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v269
	F_errdetail(m, int32(628625), v12+int32(32))
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
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v252<<(uint(int32(2))%32))+uint32(_consts[680])))
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v267)+4))
	v269 = v268
	goto L91
L94:
	;
	F_errfinish(m, int32(473983), int32(684), int32(350372))
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
	F_errmsg(m, int32(67967), v12)
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L1
	} else {
		goto L99
	}
L99:
	;
	F_errfinish(m, int32(473983), int32(610), int32(350372))
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
	v14 = *(*int32)(unsafe.Add(mBase, _consts[676]))
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
	v24 = *(*int32)(unsafe.Add(mBase, _consts[677]))
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
	F_s_lock(m, v32, int32(473983), int32(820), int32(223004))
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
	v46 = *(*int32)(unsafe.Add(mBase, _consts[350]))
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
	v54 = *(*int32)(unsafe.Add(mBase, _consts[677]))
	v56 = *(*int32)(unsafe.Add(mBase, _consts[676]))
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
	v84 = *(*int32)(unsafe.Add(mBase, _consts[676]))
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
	v9 = *(*int32)(unsafe.Add(mBase, _consts[643]))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = int32(1)
	if v10 != 0 {
		F_s_lock(m, v9, int32(473983), int32(1125), int32(70472))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return
		} else {
			v18 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v9))) = v18
			*(*int32)(unsafe.Add(mBase, uint32(v9)+92)) = v18
			v23 = *(*int32)(unsafe.Add(mBase, _consts[643]))
			v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
			*(*int32)(unsafe.Add(mBase, uint32(v23))) = int32(1)
			if v24 != 0 {
				F_s_lock(m, v23, int32(473983), int32(1107), int32(8157))
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return
				} else {
					v32 = int32(4359764)
					v33 = *(*int32)(unsafe.Add(mBase, _consts[643]))
					v34 = int32(257)
					*(*uint16)(unsafe.Add(mBase, uint32(v33)+12)) = uint16(v34)
					*(*int32)(unsafe.Add(mBase, uint32(v23))) = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v6))) = int32(81023)
					v41 = *(*int32)(unsafe.Add(mBase, _consts[643]))
					*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v41 + int32(24)
					v48 = F_pg_sprintf(m, v6+int32(16), int32(168545), v6)
					mBase = m.M
					v49 = m.ExcPending
					if v49 != 0 {
						return
					} else {
						v51 = *(*int32)(unsafe.Add(mBase, _consts[643]))
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
				v32 = int32(4359764)
				v33 = *(*int32)(unsafe.Add(mBase, _consts[643]))
				v34 = int32(257)
				*(*uint16)(unsafe.Add(mBase, uint32(v33)+12)) = uint16(v34)
				*(*int32)(unsafe.Add(mBase, uint32(v23))) = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v6))) = int32(81023)
				v41 = *(*int32)(unsafe.Add(mBase, _consts[643]))
				*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v41 + int32(24)
				v48 = F_pg_sprintf(m, v6+int32(16), int32(168545), v6)
				mBase = m.M
				v49 = m.ExcPending
				if v49 != 0 {
					return
				} else {
					v51 = *(*int32)(unsafe.Add(mBase, _consts[643]))
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
		v23 = *(*int32)(unsafe.Add(mBase, _consts[643]))
		v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
		*(*int32)(unsafe.Add(mBase, uint32(v23))) = int32(1)
		if v24 != 0 {
			F_s_lock(m, v23, int32(473983), int32(1107), int32(8157))
			mBase = m.M
			v31 = m.ExcPending
			if v31 != 0 {
				return
			} else {
				v32 = int32(4359764)
				v33 = *(*int32)(unsafe.Add(mBase, _consts[643]))
				v34 = int32(257)
				*(*uint16)(unsafe.Add(mBase, uint32(v33)+12)) = uint16(v34)
				*(*int32)(unsafe.Add(mBase, uint32(v23))) = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v6))) = int32(81023)
				v41 = *(*int32)(unsafe.Add(mBase, _consts[643]))
				*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v41 + int32(24)
				v48 = F_pg_sprintf(m, v6+int32(16), int32(168545), v6)
				mBase = m.M
				v49 = m.ExcPending
				if v49 != 0 {
					return
				} else {
					v51 = *(*int32)(unsafe.Add(mBase, _consts[643]))
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
			v32 = int32(4359764)
			v33 = *(*int32)(unsafe.Add(mBase, _consts[643]))
			v34 = int32(257)
			*(*uint16)(unsafe.Add(mBase, uint32(v33)+12)) = uint16(v34)
			*(*int32)(unsafe.Add(mBase, uint32(v23))) = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v6))) = int32(81023)
			v41 = *(*int32)(unsafe.Add(mBase, _consts[643]))
			*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v41 + int32(24)
			v48 = F_pg_sprintf(m, v6+int32(16), int32(168545), v6)
			mBase = m.M
			v49 = m.ExcPending
			if v49 != 0 {
				return
			} else {
				v51 = *(*int32)(unsafe.Add(mBase, _consts[643]))
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
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v68 int32
	_ = v68
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v143 int32
	_ = v143
	var v152 int32
	_ = v152
	v8 = m.G0
	v10 = v8 - int32(48)
	m.G0 = v10
	if l0&int32(3) == int32(0) {
		v35 = l0
		goto L6
	} else {
		goto L7
	}
L1:
	;
	m.G0 = v10 + int32(48)
	return v152
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v143
	v152 = int32(0)
	goto L1
L3:
	;
	v143 = int32(0)
	goto L2
L4:
	;
	if v68 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L5:
	;
	v68 = v60 - l0
	goto L4
L6:
	;
	v39 = v35
	goto L15
L7:
	;
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v19 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v68 = int32(0)
	goto L4
L9:
	;
	goto L10
L10:
	;
	v24 = l0
	goto L11
L11:
	;
	v28 = v24 + int32(1)
	if v28&int32(3) == int32(0) {
		v35 = v28
		goto L6
	} else {
		goto L13
	}
L12:
	;
	v60 = v28
	goto L5
L13:
	;
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28))))
	if v33 != 0 {
		v24 = v28
		goto L11
	} else {
		goto L14
	}
L14:
	;
	goto L12
L15:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
	v48 = int32(-2139062144)
	if (int32(16843008)-v45|v45)&v48 == v48 {
		v39 = v39 + int32(4)
		goto L15
	} else {
		goto L17
	}
L16:
	;
	v54 = v39
	goto L18
L17:
	;
	goto L16
L18:
	;
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54))))
	if v58 != 0 {
		v54 = v54 + int32(1)
		goto L18
	} else {
		goto L20
	}
L19:
	;
	v60 = v54
	goto L5
L20:
	;
	goto L19
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(33579140)
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = l0
	v75 = F_psprintf(m, int32(77432), v10)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L24
	} else {
		goto L25
	}
L22:
	;
	goto L23
L23:
	;
	if base.Ui32(v68) <= base.Ui32(int32(63)) {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	return int32(0)
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v75
	goto L3
L26:
	;
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v82 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L27:
	;
	goto L28
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(34103428)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = l0
	v132 = F_psprintf(m, int32(314326), v10+int32(16))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L24
	} else {
		goto L41
	}
L29:
	;
	v152 = int32(1)
	goto L1
L30:
	;
	goto L31
L31:
	;
	v89 = l0
	v91 = v82
	goto L32
L32:
	;
	if base.Ui32((v91-int32(97))&int32(255)) < base.Ui32(int32(26)) {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	v152 = v122
	goto L1
L34:
	;
	v122 = int32(1)
	v124 = v89 + v122
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v124))))
	if v125 != 0 {
		v89 = v124
		v91 = v125
		goto L32
	} else {
		goto L40
	}
L35:
	;
	if v91&int32(255) == int32(95) {
		goto L34
	} else {
		goto L36
	}
L36:
	;
	if base.Ui32((v91-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		goto L34
	} else {
		goto L37
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(33579140)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = l0
	v115 = F_psprintf(m, int32(208261), v10+int32(32))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L24
	} else {
		goto L38
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v115
	v120 = F_psprintf(m, int32(572664), int32(0))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L24
	} else {
		goto L39
	}
L39:
	;
	v143 = v120
	goto L2
L40:
	;
	goto L33
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v132
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
	var v381 int32
	_ = v381
	var v386 int32
	_ = v386
	var v390 int32
	_ = v390
	var v395 int32
	_ = v395
	var v397 int32
	_ = v397
	var v401 int32
	_ = v401
	var v407 int32
	_ = v407
	var v410 int32
	_ = v410
	var v416 int32
	_ = v416
	var v420 int32
	_ = v420
	var v422 int32
	_ = v422
	var v430 int32
	_ = v430
	var v433 int32
	_ = v433
	var v435 int32
	_ = v435
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v450 int32
	_ = v450
	var v455 int32
	_ = v455
	var v459 int32
	_ = v459
	var v464 int32
	_ = v464
	var v466 int32
	_ = v466
	var v470 int32
	_ = v470
	var v476 int32
	_ = v476
	var v479 int32
	_ = v479
	var v485 int32
	_ = v485
	var v489 int32
	_ = v489
	var v491 int32
	_ = v491
	var v499 int32
	_ = v499
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v510 int32
	_ = v510
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v527 int32
	_ = v527
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v537 int32
	_ = v537
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v554 int32
	_ = v554
	var v558 int32
	_ = v558
	var v560 int32
	_ = v560
	var v562 int32
	_ = v562
	var v568 int32
	_ = v568
	var v576 int32
	_ = v576
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v585 int32
	_ = v585
	var v589 int32
	_ = v589
	var v590 int32
	_ = v590
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v600 int32
	_ = v600
	var v602 int32
	_ = v602
	var v604 int32
	_ = v604
	var v605 int32
	_ = v605
	var v619 int32
	_ = v619
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v632 int32
	_ = v632
	var v633 int32
	_ = v633
	var v638 int32
	_ = v638
	var v647 int32
	_ = v647
	var v655 int32
	_ = v655
	var v659 int32
	_ = v659
	var v661 int32
	_ = v661
	var v663 int32
	_ = v663
	var v677 int32
	_ = v677
	var v680 int32
	_ = v680
	var v685 int32
	_ = v685
	var v686 int32
	_ = v686
	var v687 int32
	_ = v687
	var v692 int32
	_ = v692
	var v694 int32
	_ = v694
	var v707 int32
	_ = v707
	var v711 int32
	_ = v711
	var v712 int32
	_ = v712
	var v717 int32
	_ = v717
	var v719 int32
	_ = v719
	var v724 int32
	_ = v724
	var v728 int32
	_ = v728
	var v743 int32
	_ = v743
	var v751 int32
	_ = v751
	var v756 int32
	_ = v756
	var v757 int32
	_ = v757
	var v765 int32
	_ = v765
	var v766 int32
	_ = v766
	var v767 int32
	_ = v767
	var v776 int32
	_ = v776
	var v778 int32
	_ = v778
	var v779 int32
	_ = v779
	var v789 int32
	_ = v789
	var v791 int32
	_ = v791
	var v793 int32
	_ = v793
	var v795 int32
	_ = v795
	var v797 int32
	_ = v797
	var v799 int32
	_ = v799
	var v801 int32
	_ = v801
	var v803 int32
	_ = v803
	var v805 int32
	_ = v805
	var v806 int32
	_ = v806
	var v808 int32
	_ = v808
	var v810 int32
	_ = v810
	var v812 int32
	_ = v812
	var v814 int32
	_ = v814
	var v826 int32
	_ = v826
	var v828 int32
	_ = v828
	var v829 int32
	_ = v829
	var v839 int32
	_ = v839
	var v841 int32
	_ = v841
	var v846 int32
	_ = v846
	var v861 int32
	_ = v861
	var v862 int32
	_ = v862
	var v866 int32
	_ = v866
	var v869 int32
	_ = v869
	var v870 int32
	_ = v870
	var v872 int32
	_ = v872
	var v880 int32
	_ = v880
	var v883 int32
	_ = v883
	var v889 int32
	_ = v889
	var v890 int32
	_ = v890
	var v891 int32
	_ = v891
	var v894 int32
	_ = v894
	var v895 int32
	_ = v895
	var v897 int32
	_ = v897
	var v900 int32
	_ = v900
	var v908 int32
	_ = v908
	var v913 int32
	_ = v913
	var v919 int32
	_ = v919
	var v922 int32
	_ = v922
	var v923 int32
	_ = v923
	var v924 int32
	_ = v924
	var v925 int32
	_ = v925
	var v926 int32
	_ = v926
	var v927 int32
	_ = v927
	var v932 int32
	_ = v932
	var v934 int32
	_ = v934
	var v935 int32
	_ = v935
	var v939 int32
	_ = v939
	var v940 int32
	_ = v940
	var v941 int32
	_ = v941
	var v946 int32
	_ = v946
	var v947 int32
	_ = v947
	var v957 int32
	_ = v957
	var v960 int32
	_ = v960
	var v962 int32
	_ = v962
	var v963 int32
	_ = v963
	var v976 int32
	_ = v976
	var v977 int32
	_ = v977
	var v978 int32
	_ = v978
	var v980 int32
	_ = v980
	var v981 int32
	_ = v981
	var v982 int32
	_ = v982
	var v986 int32
	_ = v986
	var v987 int32
	_ = v987
	var v992 int32
	_ = v992
	var v994 int32
	_ = v994
	var v995 int32
	_ = v995
	var v996 int32
	_ = v996
	var v999 int32
	_ = v999
	var v1000 int32
	_ = v1000
	var v1001 int32
	_ = v1001
	var v1010 int32
	_ = v1010
	var v1011 int32
	_ = v1011
	var v1012 int32
	_ = v1012
	var v1016 int32
	_ = v1016
	var v1017 int32
	_ = v1017
	var v1020 int32
	_ = v1020
	var v1024 int32
	_ = v1024
	var v1029 int32
	_ = v1029
	var v1030 int32
	_ = v1030
	var v1034 int32
	_ = v1034
	var v1035 int32
	_ = v1035
	var v1038 int32
	_ = v1038
	var v1039 int32
	_ = v1039
	var v1040 int32
	_ = v1040
	var v1041 int32
	_ = v1041
	var v1055 int32
	_ = v1055
	var v1056 int32
	_ = v1056
	var v1059 int32
	_ = v1059
	var v1060 int32
	_ = v1060
	var v1061 int32
	_ = v1061
	var v1070 int32
	_ = v1070
	var v1074 int32
	_ = v1074
	var v1076 int32
	_ = v1076
	var v1077 int32
	_ = v1077
	var v1080 int32
	_ = v1080
	var v1084 int32
	_ = v1084
	var v1088 int32
	_ = v1088
	var v1089 int32
	_ = v1089
	var v1093 int32
	_ = v1093
	var v1094 int32
	_ = v1094
	var v1097 int32
	_ = v1097
	var v1098 int32
	_ = v1098
	var v1099 int32
	_ = v1099
	var v1102 int32
	_ = v1102
	var v1110 int32
	_ = v1110
	var v1113 int32
	_ = v1113
	var v1116 int32
	_ = v1116
	var v1117 int32
	_ = v1117
	var v1130 int32
	_ = v1130
	var v1131 int32
	_ = v1131
	var v1135 int32
	_ = v1135
	var v1148 int32
	_ = v1148
	var v1149 int32
	_ = v1149
	var v1152 int32
	_ = v1152
	var v1155 int32
	_ = v1155
	var v1156 int32
	_ = v1156
	var v1157 int32
	_ = v1157
	var v1161 int32
	_ = v1161
	var v1163 int32
	_ = v1163
	var v1164 int32
	_ = v1164
	var v1165 int32
	_ = v1165
	var v1166 int32
	_ = v1166
	var v1167 int32
	_ = v1167
	var v1168 int32
	_ = v1168
	var v1173 int32
	_ = v1173
	var v1175 int32
	_ = v1175
	var v1177 int32
	_ = v1177
	var v1179 int32
	_ = v1179
	var v1180 int32
	_ = v1180
	var v1184 int32
	_ = v1184
	var v1188 int32
	_ = v1188
	var v1190 int32
	_ = v1190
	var v1193 int32
	_ = v1193
	var v1200 int32
	_ = v1200
	var v1202 int32
	_ = v1202
	var v1205 int32
	_ = v1205
	var v1208 int32
	_ = v1208
	var v1209 int32
	_ = v1209
	var v1211 int32
	_ = v1211
	var v1213 int32
	_ = v1213
	var v1214 int32
	_ = v1214
	var v1217 int32
	_ = v1217
	var v1218 int32
	_ = v1218
	var v1220 int32
	_ = v1220
	var v1222 int32
	_ = v1222
	var v1225 int32
	_ = v1225
	var v1229 int32
	_ = v1229
	var v1230 int32
	_ = v1230
	var v1234 int32
	_ = v1234
	var v1242 int32
	_ = v1242
	var v1243 int32
	_ = v1243
	var v1246 int32
	_ = v1246
	var v1247 int32
	_ = v1247
	var v1248 int32
	_ = v1248
	var v1250 int32
	_ = v1250
	var v1251 int32
	_ = v1251
	var v1254 int32
	_ = v1254
	var v1255 int32
	_ = v1255
	var v1257 int32
	_ = v1257
	var v1260 int32
	_ = v1260
	var v1261 int32
	_ = v1261
	var v1262 int32
	_ = v1262
	var v1265 int32
	_ = v1265
	var v1274 int32
	_ = v1274
	var v1275 int32
	_ = v1275
	var v1276 int32
	_ = v1276
	var v1277 int32
	_ = v1277
	var v1278 int32
	_ = v1278
	var v1282 int32
	_ = v1282
	var v1283 int32
	_ = v1283
	var v1287 int32
	_ = v1287
	var v1288 int32
	_ = v1288
	var v1289 int32
	_ = v1289
	var v1290 int32
	_ = v1290
	var v1291 int32
	_ = v1291
	var v1292 int32
	_ = v1292
	var v1293 int32
	_ = v1293
	var v1294 int32
	_ = v1294
	var v1295 int32
	_ = v1295
	var v1296 int32
	_ = v1296
	var v1299 int32
	_ = v1299
	var v1301 int32
	_ = v1301
	var v1302 int32
	_ = v1302
	var v1306 int32
	_ = v1306
	var v1307 int32
	_ = v1307
	var v1313 int32
	_ = v1313
	var v1314 int32
	_ = v1314
	var v1315 int32
	_ = v1315
	var v1317 int32
	_ = v1317
	var v1318 int32
	_ = v1318
	var v1319 int32
	_ = v1319
	var v1321 int32
	_ = v1321
	var v1324 int32
	_ = v1324
	var v1325 int32
	_ = v1325
	var v1327 int32
	_ = v1327
	var v1329 int32
	_ = v1329
	var v1330 int32
	_ = v1330
	var v1334 int32
	_ = v1334
	var v1335 int32
	_ = v1335
	var v1336 int32
	_ = v1336
	var v1340 int32
	_ = v1340
	var v1341 int32
	_ = v1341
	var v1344 int32
	_ = v1344
	var v1345 int32
	_ = v1345
	var v1346 int32
	_ = v1346
	var v1355 int32
	_ = v1355
	var v1357 int32
	_ = v1357
	var v1359 int32
	_ = v1359
	var v1366 int32
	_ = v1366
	var v1373 int32
	_ = v1373
	var v1376 int32
	_ = v1376
	var v1377 int32
	_ = v1377
	var v1382 int32
	_ = v1382
	var v1386 int32
	_ = v1386
	var v1387 int32
	_ = v1387
	var v1391 int32
	_ = v1391
	var v1392 int32
	_ = v1392
	var v1397 int32
	_ = v1397
	var v1399 int32
	_ = v1399
	var v1401 int32
	_ = v1401
	var v1402 int32
	_ = v1402
	var v1416 int32
	_ = v1416
	var v1421 int32
	_ = v1421
	var v1422 int32
	_ = v1422
	var v1424 int32
	_ = v1424
	var v1425 int32
	_ = v1425
	var v1429 int32
	_ = v1429
	var v1430 int32
	_ = v1430
	var v1435 int32
	_ = v1435
	var v1444 int32
	_ = v1444
	var v1452 int32
	_ = v1452
	var v1456 int32
	_ = v1456
	var v1458 int32
	_ = v1458
	var v1460 int32
	_ = v1460
	var v1461 int32
	_ = v1461
	var v1462 int32
	_ = v1462
	var v1463 int32
	_ = v1463
	var v1467 int32
	_ = v1467
	var v1470 int32
	_ = v1470
	var v1476 int32
	_ = v1476
	var v1478 int32
	_ = v1478
	var v1480 int32
	_ = v1480
	var v1487 int32
	_ = v1487
	var v1494 int32
	_ = v1494
	var v1497 int32
	_ = v1497
	var v1498 int32
	_ = v1498
	var v1503 int32
	_ = v1503
	var v1507 int32
	_ = v1507
	var v1508 int32
	_ = v1508
	var v1512 int32
	_ = v1512
	var v1513 int32
	_ = v1513
	var v1518 int32
	_ = v1518
	var v1520 int32
	_ = v1520
	var v1522 int32
	_ = v1522
	var v1523 int32
	_ = v1523
	var v1537 int32
	_ = v1537
	var v1542 int32
	_ = v1542
	var v1543 int32
	_ = v1543
	var v1545 int32
	_ = v1545
	var v1546 int32
	_ = v1546
	var v1550 int32
	_ = v1550
	var v1551 int32
	_ = v1551
	var v1556 int32
	_ = v1556
	var v1565 int32
	_ = v1565
	var v1573 int32
	_ = v1573
	var v1577 int32
	_ = v1577
	var v1579 int32
	_ = v1579
	var v1583 int32
	_ = v1583
	var v1588 int32
	_ = v1588
	var v1593 int32
	_ = v1593
	var v1605 int32
	_ = v1605
	var v1609 int32
	_ = v1609
	var v1614 int32
	_ = v1614
	var v1624 int32
	_ = v1624
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
	v33 = *(*int32)(unsafe.Add(mBase, _consts[361]))
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
	v39 = *(*int32)(unsafe.Add(mBase, _consts[668]))
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
	m.G0 = v1624 + int32(16)
	return v1614
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
	v1614 = v89
	v1624 = v16
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
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122)+uint32(_consts[669]))))
	v127 = v109 << (uint(int32(1)) % 32)
	v130 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v127)+uint32(_consts[670]))))
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
	v135 = int32(*(*int16)(unsafe.Add(mBase, uint32(v127)+uint32(_consts[671]))))
	v136 = v135 + v125
	v141 = int32(*(*int16)(unsafe.Add(mBase, uint32(v136<<(uint(int32(1))%32))+uint32(_consts[672]))))
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
	v202 = int32(*(*int16)(unsafe.Add(mBase, uint32(v188<<(uint(v196)%32))+uint32(_consts[673]))))
	if v202 != int32(284) {
		v109 = v202
		v115 = v115 + v196
		goto L27
	} else {
		goto L41
	}
L35:
	;
	v160 = int32(*(*int16)(unsafe.Add(mBase, uint32(v143<<(uint(int32(1))%32))+uint32(_consts[674]))))
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
	v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v145)+uint32(_consts[675]))))
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
	v173 = int32(*(*int16)(unsafe.Add(mBase, uint32(v160<<(uint(v169)%32))+uint32(_consts[671]))))
	v174 = v168 + v173
	v179 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v174<<(uint(v169)%32))+uint32(_consts[672]))))
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
	v245 = int32(*(*int16)(unsafe.Add(mBase, uint32(v220<<(uint(int32(1))%32))+uint32(_consts[670]))))
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
		v1614 = int32(266)
		v1624 = v102
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
	*(*int32)(unsafe.Add(mBase, uint32(v93)+36)) = v1593
	*(*int32)(unsafe.Add(mBase, uint32(v93)+48)) = int32(0)
	v1605 = *(*int32)(unsafe.Add(mBase, uint32(v93)+44))
	v1609 = base.I32_div_s(v1605-int32(1), int32(2))
	v248 = v1609 + int32(36)
	goto L46
L49:
	;
	F_yy_fatal_error_3(m, int32(29516))
	mBase = m.M
	v1588 = m.ExcPending
	if v1588 != 0 {
		goto L18
	} else {
		goto L342
	}
L50:
	;
	F_yy_fatal_error_3(m, int32(646259))
	mBase = m.M
	v1583 = m.ExcPending
	if v1583 != 0 {
		goto L18
	} else {
		goto L341
	}
L51:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L52:
	;
	v1476 = v1463 + v1470
	*(*int32)(unsafe.Add(mBase, uint32(v93)+36)) = v1476
	v1478 = *(*int32)(unsafe.Add(mBase, uint32(v93)+44))
	if base.Ui32(v1476) <= base.Ui32(v1467) {
		v220 = v1478
		v226 = v1476
		v228 = v1467
		goto L44
	} else {
		goto L322
	}
L53:
	;
	v1149 = *(*int32)(unsafe.Add(mBase, uint32(v1148)))
	*(*int32)(unsafe.Add(mBase, uint32(v1149)+16)) = v1135
	v1152 = *(*int32)(unsafe.Add(mBase, uint32(v93)+28))
	if v1152 != 0 {
		v1274 = int32(0)
		goto L266
	} else {
		goto L267
	}
L54:
	;
	v1130 = *(*int32)(unsafe.Add(mBase, uint32(v93)+20))
	v1131 = *(*int32)(unsafe.Add(mBase, uint32(v93)+12))
	v1135 = v1117
	v1148 = v1130 + v1131<<(uint(int32(2))%32)
	goto L53
L55:
	;
	F_yy_fatal_error_3(m, int32(436751))
	mBase = m.M
	v1116 = m.ExcPending
	if v1116 != 0 {
		goto L18
	} else {
		goto L265
	}
L56:
	;
	F_yy_fatal_error_3(m, int32(432616))
	mBase = m.M
	v1113 = m.ExcPending
	if v1113 != 0 {
		goto L18
	} else {
		goto L264
	}
L57:
	;
	F_replication_yyerror(m, int32(254676))
	mBase = m.M
	v1110 = m.ExcPending
	if v1110 != 0 {
		goto L18
	} else {
		goto L263
	}
L58:
	;
	v1614 = int32(272)
	v1624 = v102
	goto L21
L59:
	;
	v1614 = int32(265)
	v1624 = v102
	goto L21
L60:
	;
	v1614 = int32(264)
	v1624 = v102
	goto L21
L61:
	;
	v1614 = int32(263)
	v1624 = v102
	goto L21
L62:
	;
	v1614 = int32(262)
	v1624 = v102
	goto L21
L63:
	;
	F_yy_fatal_error_3(m, int32(407247))
	mBase = m.M
	v1102 = m.ExcPending
	if v1102 != 0 {
		goto L18
	} else {
		goto L262
	}
L64:
	;
	v515 = *(*int32)(unsafe.Add(mBase, uint32(v93)+80))
	v516 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v93)+24)))
	*(*uint8)(unsafe.Add(mBase, uint32(v226))) = uint8(v516)
	v518 = *(*int32)(unsafe.Add(mBase, uint32(v93)+20))
	v519 = *(*int32)(unsafe.Add(mBase, uint32(v93)+12))
	v522 = v518 + v519<<(uint(int32(2))%32)
	v523 = *(*int32)(unsafe.Add(mBase, uint32(v522)))
	v524 = *(*int32)(unsafe.Add(mBase, uint32(v523)+44))
	if v524 == int32(0) {
		goto L141
	} else {
		goto L142
	}
L65:
	;
	F_yy_fatal_error_3(m, int32(436103))
	mBase = m.M
	v514 = m.ExcPending
	if v514 != 0 {
		goto L18
	} else {
		goto L140
	}
L66:
	;
	v1614 = int32(0)
	v1624 = v102
	goto L21
L67:
	;
	F_replication_yyerror(m, int32(316794))
	mBase = m.M
	v510 = m.ExcPending
	if v510 != 0 {
		goto L18
	} else {
		goto L139
	}
L68:
	;
	v506 = *(*int32)(unsafe.Add(mBase, uint32(v93)+80))
	v507 = int32(*(*int8)(unsafe.Add(mBase, uint32(v506))))
	v1614 = v507
	v1624 = v102
	goto L21
L69:
	;
	v442 = *(*int32)(unsafe.Add(mBase, uint32(v93)+80))
	if v442&int32(3) == int32(0) {
		v466 = v442
		goto L123
	} else {
		goto L124
	}
L70:
	;
	v435 = *(*int32)(unsafe.Add(mBase, uint32(v93)))
	v438 = *(*int32)(unsafe.Add(mBase, uint32(v93)+80))
	v439 = *(*int32)(unsafe.Add(mBase, uint32(v93)+32))
	F_appendBinaryStringInfo(m, v435+int32(4), v438, v439)
	mBase = m.M
	v441 = m.ExcPending
	if v441 != 0 {
		goto L18
	} else {
		goto L120
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
	if v373&int32(3) == v363 {
		v397 = v373
		goto L104
	} else {
		goto L105
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
	v1614 = int32(258)
	v1624 = v102
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
	v294 = F_sscanf(m, v286, int32(496235), v102)
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
	v1614 = int32(282)
	v1624 = v102
	goto L21
L80:
	;
	v1614 = int32(271)
	v1624 = v102
	goto L21
L81:
	;
	v1614 = int32(281)
	v1624 = v102
	goto L21
L82:
	;
	v1614 = int32(280)
	v1624 = v102
	goto L21
L83:
	;
	v1614 = int32(279)
	v1624 = v102
	goto L21
L84:
	;
	v1614 = int32(278)
	v1624 = v102
	goto L21
L85:
	;
	v1614 = int32(277)
	v1624 = v102
	goto L21
L86:
	;
	v1614 = int32(275)
	v1624 = v102
	goto L21
L87:
	;
	v1614 = int32(274)
	v1624 = v102
	goto L21
L88:
	;
	v1614 = int32(276)
	v1624 = v102
	goto L21
L89:
	;
	v1614 = int32(273)
	v1624 = v102
	goto L21
L90:
	;
	v1614 = int32(270)
	v1624 = v102
	goto L21
L91:
	;
	v1614 = int32(269)
	v1624 = v102
	goto L21
L92:
	;
	v1614 = int32(268)
	v1624 = v102
	goto L21
L93:
	;
	v1614 = int32(267)
	v1624 = v102
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
	v1614 = int32(260)
	v1624 = v102
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
	v1614 = int32(261)
	v1624 = v102
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
	F_truncate_identifier(m, v373, v430, int32(1))
	mBase = m.M
	v433 = m.ExcPending
	if v433 != 0 {
		goto L18
	} else {
		goto L119
	}
L103:
	;
	v430 = v422 - v373
	goto L102
L104:
	;
	v401 = v397
	goto L113
L105:
	;
	v381 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v373))))
	if v381 == int32(0) {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	v430 = int32(0)
	goto L102
L107:
	;
	goto L108
L108:
	;
	v386 = v373
	goto L109
L109:
	;
	v390 = v386 + int32(1)
	if v390&int32(3) == int32(0) {
		v397 = v390
		goto L104
	} else {
		goto L111
	}
L110:
	;
	v422 = v390
	goto L103
L111:
	;
	v395 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v390))))
	if v395 != 0 {
		v386 = v390
		goto L109
	} else {
		goto L112
	}
L112:
	;
	goto L110
L113:
	;
	v407 = *(*int32)(unsafe.Add(mBase, uint32(v401)))
	v410 = int32(-2139062144)
	if (int32(16843008)-v407|v407)&v410 == v410 {
		v401 = v401 + int32(4)
		goto L113
	} else {
		goto L115
	}
L114:
	;
	v416 = v401
	goto L116
L115:
	;
	goto L114
L116:
	;
	v420 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v416))))
	if v420 != 0 {
		v416 = v416 + int32(1)
		goto L116
	} else {
		goto L118
	}
L117:
	;
	v422 = v416
	goto L103
L118:
	;
	goto L117
L119:
	;
	v1614 = int32(259)
	v1624 = v102
	goto L21
L120:
	;
	goto L25
L121:
	;
	v501 = F_downcase_truncate_identifier(m, v442, v499, int32(1))
	mBase = m.M
	v502 = m.ExcPending
	if v502 != 0 {
		goto L18
	} else {
		goto L138
	}
L122:
	;
	v499 = v491 - v442
	goto L121
L123:
	;
	v470 = v466
	goto L132
L124:
	;
	v450 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v442))))
	if v450 == int32(0) {
		goto L125
	} else {
		goto L126
	}
L125:
	;
	v499 = int32(0)
	goto L121
L126:
	;
	goto L127
L127:
	;
	v455 = v442
	goto L128
L128:
	;
	v459 = v455 + int32(1)
	if v459&int32(3) == int32(0) {
		v466 = v459
		goto L123
	} else {
		goto L130
	}
L129:
	;
	v491 = v459
	goto L122
L130:
	;
	v464 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v459))))
	if v464 != 0 {
		v455 = v459
		goto L128
	} else {
		goto L131
	}
L131:
	;
	goto L129
L132:
	;
	v476 = *(*int32)(unsafe.Add(mBase, uint32(v470)))
	v479 = int32(-2139062144)
	if (int32(16843008)-v476|v476)&v479 == v479 {
		v470 = v470 + int32(4)
		goto L132
	} else {
		goto L134
	}
L133:
	;
	v485 = v470
	goto L135
L134:
	;
	goto L133
L135:
	;
	v489 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v485))))
	if v489 != 0 {
		v485 = v485 + int32(1)
		goto L135
	} else {
		goto L137
	}
L136:
	;
	v491 = v485
	goto L122
L137:
	;
	goto L136
L138:
	;
	v503 = *(*int32)(unsafe.Add(mBase, uint32(v93)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v503))) = v501
	v1614 = int32(259)
	v1624 = v102
	goto L21
L139:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L140:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L141:
	;
	v527 = *(*int32)(unsafe.Add(mBase, uint32(v523)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v93)+28)) = v527
	v529 = *(*int32)(unsafe.Add(mBase, uint32(v522)))
	v530 = *(*int32)(unsafe.Add(mBase, uint32(v93)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v529))) = v530
	v532 = *(*int32)(unsafe.Add(mBase, uint32(v93)+20))
	v533 = *(*int32)(unsafe.Add(mBase, uint32(v93)+12))
	v534 = int32(2)
	v537 = *(*int32)(unsafe.Add(mBase, uint32(v532+v533<<(uint(v534)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v537)+44)) = int32(1)
	v540 = *(*int32)(unsafe.Add(mBase, uint32(v93)+20))
	v541 = *(*int32)(unsafe.Add(mBase, uint32(v93)+12))
	v545 = *(*int32)(unsafe.Add(mBase, uint32(v540+v541<<(uint(v534)%32))))
	v546 = v545
	v547 = v540
	v548 = v541
	goto L143
L142:
	;
	v546 = v523
	v547 = v518
	v548 = v519
	goto L143
L143:
	;
	v549 = *(*int32)(unsafe.Add(mBase, uint32(v93)+36))
	v550 = *(*int32)(unsafe.Add(mBase, uint32(v546)+4))
	v551 = *(*int32)(unsafe.Add(mBase, uint32(v93)+28))
	v552 = v550 + v551
	if base.Ui32(v549) <= base.Ui32(v552) {
		goto L144
	} else {
		goto L145
	}
L144:
	;
	v554 = *(*int32)(unsafe.Add(mBase, uint32(v93)+80))
	v558 = v554 + (v515 ^ int32(-1)) + v226
	*(*int32)(unsafe.Add(mBase, uint32(v93)+36)) = v558
	v560 = *(*int32)(unsafe.Add(mBase, uint32(v93)+44))
	if base.Ui32(v554) < base.Ui32(v558) {
		goto L147
	} else {
		goto L148
	}
L145:
	;
	goto L146
L146:
	;
	if base.Ui32(v552+int32(1)) < base.Ui32(v549) {
		goto L56
	} else {
		goto L179
	}
L147:
	;
	v562 = v560
	v568 = v554
	goto L150
L148:
	;
	v663 = v560
	goto L149
L149:
	;
	v677 = v663 << (uint(int32(1)) % 32)
	v680 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v677)+uint32(_consts[670]))))
	if v680 != 0 {
		goto L168
	} else {
		goto L169
	}
L150:
	;
	v576 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v568))))
	if v576 != 0 {
		goto L152
	} else {
		goto L153
	}
L151:
	;
	v663 = v659
	goto L149
L152:
	;
	v579 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v576)+uint32(_consts[669]))))
	v580 = v579
	goto L154
L153:
	;
	v580 = int32(1)
	goto L154
L154:
	;
	v585 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v562<<(uint(int32(1))%32))+uint32(_consts[670]))))
	if v585 != 0 {
		goto L155
	} else {
		goto L156
	}
L155:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v93)+68)) = v568
	*(*int32)(unsafe.Add(mBase, uint32(v93)+64)) = v562
	goto L157
L156:
	;
	goto L157
L157:
	;
	v589 = v580 & int32(255)
	v590 = int32(1)
	v594 = int32(*(*int16)(unsafe.Add(mBase, uint32(v562<<(uint(v590)%32))+uint32(_consts[671]))))
	v595 = v589 + v594
	v600 = int32(*(*int16)(unsafe.Add(mBase, uint32(v595<<(uint(v590)%32))+uint32(_consts[672]))))
	if v600 != v562 {
		goto L158
	} else {
		goto L159
	}
L158:
	;
	v602 = v562
	v604 = v580
	v605 = v589
	goto L161
L159:
	;
	v647 = v595
	goto L160
L160:
	;
	v655 = int32(1)
	v659 = int32(*(*int16)(unsafe.Add(mBase, uint32(v647<<(uint(v655)%32))+uint32(_consts[673]))))
	v661 = v568 + v655
	if v661 != v558 {
		v562 = v659
		v568 = v661
		goto L150
	} else {
		goto L167
	}
L161:
	;
	v619 = int32(*(*int16)(unsafe.Add(mBase, uint32(v602<<(uint(int32(1))%32))+uint32(_consts[674]))))
	if int32(285) <= v619 {
		goto L163
	} else {
		goto L164
	}
L162:
	;
	v647 = v633
	goto L160
L163:
	;
	v624 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v605)+uint32(_consts[675]))))
	v625 = v624
	goto L165
L164:
	;
	v625 = v604
	goto L165
L165:
	;
	v627 = v625 & int32(255)
	v628 = int32(1)
	v632 = int32(*(*int16)(unsafe.Add(mBase, uint32(v619<<(uint(v628)%32))+uint32(_consts[671]))))
	v633 = v627 + v632
	v638 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v633<<(uint(v628)%32))+uint32(_consts[672]))))
	if v638 != v619&int32(65535) {
		v602 = v619
		v604 = v625
		v605 = v627
		goto L161
	} else {
		goto L166
	}
L166:
	;
	goto L162
L167:
	;
	goto L151
L168:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v93)+68)) = v558
	*(*int32)(unsafe.Add(mBase, uint32(v93)+64)) = v663
	goto L170
L169:
	;
	goto L170
L170:
	;
	v685 = int32(*(*int16)(unsafe.Add(mBase, uint32(v677)+uint32(_consts[671]))))
	v686 = int32(1)
	v687 = v685 + v686
	v692 = int32(*(*int16)(unsafe.Add(mBase, uint32(v687<<(uint(v686)%32))+uint32(_consts[672]))))
	if v692 != v663 {
		goto L171
	} else {
		goto L172
	}
L171:
	;
	v694 = v663
	goto L174
L172:
	;
	v728 = v687
	goto L173
L173:
	;
	v743 = int32(*(*int16)(unsafe.Add(mBase, uint32(v728<<(uint(int32(1))%32))+uint32(_consts[673]))))
	if v743 == int32(284) {
		v213 = v554
		goto L42
	} else {
		goto L177
	}
L174:
	;
	v707 = int32(1)
	v711 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v694<<(uint(v707)%32))+uint32(_consts[674]))))
	v712 = base.I32_extend16_s(v711)
	v717 = int32(*(*int16)(unsafe.Add(mBase, uint32(v712<<(uint(v707)%32))+uint32(_consts[671]))))
	v719 = v717 + v707
	v724 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v719<<(uint(v707)%32))+uint32(_consts[672]))))
	if v711 != v724 {
		v694 = v712
		goto L174
	} else {
		goto L176
	}
L175:
	;
	v728 = v719
	goto L173
L176:
	;
	goto L175
L177:
	;
	if v728&int32(2147483647) == int32(0) {
		v213 = v554
		goto L42
	} else {
		goto L178
	}
L178:
	;
	v751 = v558 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v93)+36)) = v751
	v109 = v743
	v113 = v554
	v115 = v751
	goto L27
L179:
	;
	v756 = *(*int32)(unsafe.Add(mBase, uint32(v93)+80))
	v757 = *(*int32)(unsafe.Add(mBase, uint32(v546)+40))
	if v757 == int32(0) {
		goto L180
	} else {
		goto L181
	}
L180:
	;
	if v549-v756 != int32(1) {
		v1463 = v550
		v1467 = v756
		v1470 = v551
		goto L52
	} else {
		goto L183
	}
L181:
	;
	goto L182
L182:
	;
	v765 = v756 ^ int32(-1) + v549
	if v765 != 0 {
		goto L184
	} else {
		goto L185
	}
L183:
	;
	v1593 = v756
	goto L48
L184:
	;
	v766 = int32(7)
	v767 = v765 & v766
	if base.Ui32(v549-v756-int32(2)) < base.Ui32(v766) {
		goto L188
	} else {
		goto L189
	}
L185:
	;
	v869 = v546
	v870 = v547
	v872 = v548
	goto L186
L186:
	;
	v880 = *(*int32)(unsafe.Add(mBase, uint32(v869)+44))
	if v880 == int32(2) {
		goto L200
	} else {
		goto L201
	}
L187:
	;
	if v767 != 0 {
		goto L194
	} else {
		goto L195
	}
L188:
	;
	v812 = v550
	v814 = v756
	goto L187
L189:
	;
	goto L190
L190:
	;
	v776 = v550
	v778 = v756
	v779 = int32(0)
	goto L191
L191:
	;
	v789 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v778))))
	*(*uint8)(unsafe.Add(mBase, uint32(v776))) = uint8(v789)
	v791 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v778)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v776)+1)) = uint8(v791)
	v793 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v778)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v776)+2)) = uint8(v793)
	v795 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v778)+3)))
	*(*uint8)(unsafe.Add(mBase, uint32(v776)+3)) = uint8(v795)
	v797 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v778)+4)))
	*(*uint8)(unsafe.Add(mBase, uint32(v776)+4)) = uint8(v797)
	v799 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v778)+5)))
	*(*uint8)(unsafe.Add(mBase, uint32(v776)+5)) = uint8(v799)
	v801 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v778)+6)))
	*(*uint8)(unsafe.Add(mBase, uint32(v776)+6)) = uint8(v801)
	v803 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v778)+7)))
	*(*uint8)(unsafe.Add(mBase, uint32(v776)+7)) = uint8(v803)
	v805 = int32(8)
	v806 = v776 + v805
	v808 = v778 + v805
	v810 = v779 + v805
	if v810 != v765&int32(-8) {
		v776 = v806
		v778 = v808
		v779 = v810
		goto L191
	} else {
		goto L193
	}
L192:
	;
	v812 = v806
	v814 = v808
	goto L187
L193:
	;
	goto L192
L194:
	;
	v826 = v812
	v828 = v814
	v829 = int32(0)
	goto L197
L195:
	;
	goto L196
L196:
	;
	v861 = *(*int32)(unsafe.Add(mBase, uint32(v93)+20))
	v862 = *(*int32)(unsafe.Add(mBase, uint32(v93)+12))
	v866 = *(*int32)(unsafe.Add(mBase, uint32(v861+v862<<(uint(int32(2))%32))))
	v869 = v866
	v870 = v861
	v872 = v862
	goto L186
L197:
	;
	v839 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v828))))
	*(*uint8)(unsafe.Add(mBase, uint32(v826))) = uint8(v839)
	v841 = int32(1)
	v846 = v829 + v841
	if v846 != v767 {
		v826 = v826 + v841
		v828 = v828 + v841
		v829 = v846
		goto L197
	} else {
		goto L199
	}
L198:
	;
	goto L196
L199:
	;
	goto L198
L200:
	;
	v883 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v93)+28)) = v883
	v1135 = v883
	v1148 = v870 + v872<<(uint(int32(2))%32)
	goto L53
L201:
	;
	goto L202
L202:
	;
	v889 = *(*int32)(unsafe.Add(mBase, uint32(v869)+12))
	v890 = v756 - v549
	v891 = v889 + v890
	if v891 == int32(0) {
		goto L203
	} else {
		goto L204
	}
L203:
	;
	v894 = *(*int32)(unsafe.Add(mBase, uint32(v93)+36))
	v895 = v889
	v897 = v869
	v900 = v894
	goto L206
L204:
	;
	v946 = v869
	v947 = v891
	goto L205
L205:
	;
	v957 = int32(8192)
	if base.Ui32(v957) <= base.Ui32(v947) {
		goto L222
	} else {
		goto L223
	}
L206:
	;
	v908 = *(*int32)(unsafe.Add(mBase, uint32(v897)+20))
	if v908 == int32(0) {
		goto L208
	} else {
		goto L209
	}
L207:
	;
	v946 = v939
	v947 = v941
	goto L205
L208:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v897)+4)) = int32(0)
	goto L49
L209:
	;
	goto L210
L210:
	;
	v913 = *(*int32)(unsafe.Add(mBase, uint32(v897)+4))
	if base.Ui32(int32(2147483646)) <= base.Ui32(v895) {
		goto L211
	} else {
		goto L212
	}
L211:
	;
	v919 = int32(-3)
	goto L213
L212:
	;
	v919 = v895 << (uint(int32(1)) % 32)
	goto L213
L213:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v897)+12)) = v919
	v922 = v919 + int32(2)
	if v913 != 0 {
		goto L215
	} else {
		goto L216
	}
L214:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v897)+4)) = v927
	if v927 == int32(0) {
		goto L49
	} else {
		goto L220
	}
L215:
	;
	v923 = F_repalloc(m, v913, v922)
	mBase = m.M
	v924 = m.ExcPending
	if v924 != 0 {
		goto L18
	} else {
		goto L218
	}
L216:
	;
	goto L217
L217:
	;
	v925 = F_palloc(m, v922)
	mBase = m.M
	v926 = m.ExcPending
	if v926 != 0 {
		goto L18
	} else {
		goto L219
	}
L218:
	;
	v927 = v923
	goto L214
L219:
	;
	v927 = v925
	goto L214
L220:
	;
	v932 = v927 + (v900 - v913)
	*(*int32)(unsafe.Add(mBase, uint32(v93)+36)) = v932
	v934 = *(*int32)(unsafe.Add(mBase, uint32(v93)+20))
	v935 = *(*int32)(unsafe.Add(mBase, uint32(v93)+12))
	v939 = *(*int32)(unsafe.Add(mBase, uint32(v934+v935<<(uint(int32(2))%32))))
	v940 = *(*int32)(unsafe.Add(mBase, uint32(v939)+12))
	v941 = v940 + v890
	if v941 == int32(0) {
		v895 = v940
		v897 = v939
		v900 = v932
		goto L206
	} else {
		goto L221
	}
L221:
	;
	goto L207
L222:
	;
	v960 = v957
	goto L224
L223:
	;
	v960 = v947
	goto L224
L224:
	;
	v962 = *(*int32)(unsafe.Add(mBase, uint32(v946)+24))
	if v962 != 0 {
		goto L225
	} else {
		goto L226
	}
L225:
	;
	v963 = int32(0)
	goto L229
L226:
	;
	goto L227
L227:
	;
	*(*int32)(unsafe.Add(mBase, _consts[155])) = int32(0)
	v1029 = *(*int32)(unsafe.Add(mBase, uint32(v93)+20))
	v1030 = *(*int32)(unsafe.Add(mBase, uint32(v93)+12))
	v1034 = *(*int32)(unsafe.Add(mBase, uint32(v1029+v1030<<(uint(int32(2))%32))))
	v1035 = *(*int32)(unsafe.Add(mBase, uint32(v1034)+4))
	v1038 = *(*int32)(unsafe.Add(mBase, uint32(v93)+4))
	v1039 = F_fread(m, v1035+v765, int32(1), v960, v1038)
	mBase = m.M
	v1040 = m.ExcPending
	if v1040 != 0 {
		goto L18
	} else {
		goto L244
	}
L228:
	;
	switch v980 {
	case 0:
		goto L236
	default:
		v1024 = v994
		goto L234
	case 11:
		goto L235
	}
L229:
	;
	v976 = *(*int32)(unsafe.Add(mBase, uint32(v93)+4))
	v977 = F_do_getc(m, v976)
	mBase = m.M
	v978 = m.ExcPending
	if v978 != 0 {
		goto L18
	} else {
		goto L232
	}
L230:
	;
	v994 = v960
	goto L228
L231:
	;
	v981 = *(*int32)(unsafe.Add(mBase, uint32(v93)+20))
	v982 = *(*int32)(unsafe.Add(mBase, uint32(v93)+12))
	v986 = *(*int32)(unsafe.Add(mBase, uint32(v981+v982<<(uint(int32(2))%32))))
	v987 = *(*int32)(unsafe.Add(mBase, uint32(v986)+4))
	*(*uint8)(unsafe.Add(mBase, uint32(v987+v765+v963))) = uint8(v977)
	v992 = v963 + int32(1)
	if v992 != v960 {
		v963 = v992
		goto L229
	} else {
		goto L233
	}
L232:
	;
	v980 = v977 + int32(1)
	switch v980 {
	case 0, 11:
		v994 = v963
		goto L228
	default:
		goto L231
	}
L233:
	;
	goto L230
L234:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v93)+28)) = v1024
	v1117 = v1024
	goto L54
L235:
	;
	v1011 = *(*int32)(unsafe.Add(mBase, uint32(v93)+20))
	v1012 = *(*int32)(unsafe.Add(mBase, uint32(v93)+12))
	v1016 = *(*int32)(unsafe.Add(mBase, uint32(v1011+v1012<<(uint(int32(2))%32))))
	v1017 = *(*int32)(unsafe.Add(mBase, uint32(v1016)+4))
	v1020 = int32(10)
	*(*uint8)(unsafe.Add(mBase, uint32(v1017+v765+v994))) = uint8(v1020)
	v1024 = v994 + int32(1)
	goto L234
L236:
	;
	v995 = *(*int32)(unsafe.Add(mBase, uint32(v93)+4))
	v996 = *(*int32)(unsafe.Add(mBase, uint32(v995)+76))
	if v996 < int32(0) {
		goto L239
	} else {
		goto L240
	}
L237:
	;
	if int32(base.Ui32(v1001)>>(uint(int32(5))%32))&int32(1) == int32(0) {
		v1024 = v994
		goto L234
	} else {
		goto L242
	}
L238:
	;
	goto L237
L239:
	;
	v999 = *(*int32)(unsafe.Add(mBase, uint32(v995)))
	v1001 = v999
	goto L238
L240:
	;
	goto L241
L241:
	;
	v1000 = *(*int32)(unsafe.Add(mBase, uint32(v995)))
	v1001 = v1000
	goto L238
L242:
	;
	F_yy_fatal_error_3(m, int32(436751))
	mBase = m.M
	v1010 = m.ExcPending
	if v1010 != 0 {
		goto L18
	} else {
		goto L243
	}
L243:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L244:
	;
	v1041 = v1039
	goto L245
L245:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v93)+28)) = v1041
	if v1041 != 0 {
		v1117 = v1041
		goto L54
	} else {
		goto L247
	}
L247:
	;
	v1055 = *(*int32)(unsafe.Add(mBase, uint32(v93)+4))
	v1056 = *(*int32)(unsafe.Add(mBase, uint32(v1055)+76))
	if v1056 < int32(0) {
		goto L250
	} else {
		goto L251
	}
L248:
	;
	if int32(base.Ui32(v1061)>>(uint(int32(5))%32))&int32(1) == int32(0) {
		goto L253
	} else {
		goto L254
	}
L249:
	;
	goto L248
L250:
	;
	v1059 = *(*int32)(unsafe.Add(mBase, uint32(v1055)))
	v1061 = v1059
	goto L249
L251:
	;
	goto L252
L252:
	;
	v1060 = *(*int32)(unsafe.Add(mBase, uint32(v1055)))
	v1061 = v1060
	goto L249
L253:
	;
	v1117 = int32(0)
	goto L54
L254:
	;
	goto L255
L255:
	;
	v1070 = *(*int32)(unsafe.Add(mBase, _consts[155]))
	if v1070 != int32(27) {
		goto L55
	} else {
		goto L256
	}
L256:
	;
	v1074 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[155])) = v1074
	v1076 = *(*int32)(unsafe.Add(mBase, uint32(v93)+4))
	v1077 = *(*int32)(unsafe.Add(mBase, uint32(v1076)+76))
	if v1074 <= v1077 {
		goto L258
	} else {
		goto L259
	}
L257:
	;
	v1088 = *(*int32)(unsafe.Add(mBase, uint32(v93)+20))
	v1089 = *(*int32)(unsafe.Add(mBase, uint32(v93)+12))
	v1093 = *(*int32)(unsafe.Add(mBase, uint32(v1088+v1089<<(uint(int32(2))%32))))
	v1094 = *(*int32)(unsafe.Add(mBase, uint32(v1093)+4))
	v1097 = *(*int32)(unsafe.Add(mBase, uint32(v93)+4))
	v1098 = F_fread(m, v1094+v765, int32(1), v960, v1097)
	mBase = m.M
	v1099 = m.ExcPending
	if v1099 != 0 {
		goto L18
	} else {
		goto L261
	}
L258:
	;
	v1080 = *(*int32)(unsafe.Add(mBase, uint32(v1076)))
	*(*int32)(unsafe.Add(mBase, uint32(v1076))) = v1080 & int32(-49)
	goto L257
L259:
	;
	goto L260
L260:
	;
	v1084 = *(*int32)(unsafe.Add(mBase, uint32(v1076)))
	*(*int32)(unsafe.Add(mBase, uint32(v1076))) = v1084 & int32(-49)
	goto L257
L261:
	;
	v1041 = v1098
	goto L245
L262:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L263:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L264:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L265:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L266:
	;
	v1275 = *(*int32)(unsafe.Add(mBase, uint32(v93)+28))
	v1276 = v1275 + v765
	v1277 = *(*int32)(unsafe.Add(mBase, uint32(v93)+20))
	v1278 = *(*int32)(unsafe.Add(mBase, uint32(v93)+12))
	v1282 = *(*int32)(unsafe.Add(mBase, uint32(v1277+v1278<<(uint(int32(2))%32))))
	v1283 = *(*int32)(unsafe.Add(mBase, uint32(v1282)+12))
	if base.Ui32(v1283) < base.Ui32(v1276) {
		goto L290
	} else {
		goto L291
	}
L267:
	;
	if v765 == int32(0) {
		goto L268
	} else {
		goto L269
	}
L268:
	;
	v1155 = *(*int32)(unsafe.Add(mBase, uint32(v93)+4))
	v1156 = *(*int32)(unsafe.Add(mBase, uint32(v93)+20))
	if v1156 != 0 {
		goto L273
	} else {
		goto L274
	}
L269:
	;
	goto L270
L270:
	;
	v1260 = *(*int32)(unsafe.Add(mBase, uint32(v93)+20))
	v1261 = *(*int32)(unsafe.Add(mBase, uint32(v93)+12))
	v1262 = int32(2)
	v1265 = *(*int32)(unsafe.Add(mBase, uint32(v1260+v1261<<(uint(v1262)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v1265)+44)) = v1262
	v1274 = v1262
	goto L266
L271:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1222)+40)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1222))) = v1155
	v1229 = *(*int32)(unsafe.Add(mBase, uint32(v93)+20))
	if v1229 != 0 {
		goto L286
	} else {
		goto L287
	}
L272:
	;
	v1179 = *(*int32)(unsafe.Add(mBase, _consts[155]))
	v1180 = *(*int32)(unsafe.Add(mBase, uint32(v93)+12))
	v1184 = *(*int32)(unsafe.Add(mBase, uint32(v1177+v1180<<(uint(int32(2))%32))))
	if v1184 == int32(0) {
		goto L280
	} else {
		goto L281
	}
L273:
	;
	v1157 = *(*int32)(unsafe.Add(mBase, uint32(v93)+12))
	v1161 = *(*int32)(unsafe.Add(mBase, uint32(v1156+v1157<<(uint(int32(2))%32))))
	if v1161 != 0 {
		v1177 = v1156
		goto L272
	} else {
		goto L276
	}
L274:
	;
	goto L275
L275:
	;
	F_replication_yyensure_buffer_stack(m, v93)
	mBase = m.M
	v1163 = m.ExcPending
	if v1163 != 0 {
		goto L18
	} else {
		goto L277
	}
L276:
	;
	goto L275
L277:
	;
	v1164 = *(*int32)(unsafe.Add(mBase, uint32(v93)+4))
	v1165 = F_replication_yy_create_buffer(m, v1164, v93)
	mBase = m.M
	v1166 = m.ExcPending
	if v1166 != 0 {
		goto L18
	} else {
		goto L278
	}
L278:
	;
	v1167 = *(*int32)(unsafe.Add(mBase, uint32(v93)+20))
	v1168 = *(*int32)(unsafe.Add(mBase, uint32(v93)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v1167+v1168<<(uint(int32(2))%32)))) = v1165
	v1173 = *(*int32)(unsafe.Add(mBase, uint32(v93)+20))
	if v1173 != 0 {
		v1177 = v1173
		goto L272
	} else {
		goto L279
	}
L279:
	;
	v1175 = *(*int32)(unsafe.Add(mBase, _consts[155]))
	v1222 = int32(0)
	v1225 = v1175
	goto L271
L280:
	;
	v1222 = int32(0)
	v1225 = v1179
	goto L271
L281:
	;
	goto L282
L282:
	;
	v1188 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1184)+16)) = v1188
	v1190 = *(*int32)(unsafe.Add(mBase, uint32(v1184)+4))
	*(*uint8)(unsafe.Add(mBase, uint32(v1190))) = uint8(v1188)
	v1193 = *(*int32)(unsafe.Add(mBase, uint32(v1184)+4))
	*(*uint8)(unsafe.Add(mBase, uint32(v1193)+1)) = uint8(v1188)
	*(*int32)(unsafe.Add(mBase, uint32(v1184)+44)) = v1188
	*(*int32)(unsafe.Add(mBase, uint32(v1184)+28)) = int32(1)
	v1200 = *(*int32)(unsafe.Add(mBase, uint32(v1184)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1184)+8)) = v1200
	v1202 = *(*int32)(unsafe.Add(mBase, uint32(v93)+20))
	if v1202 == v1188 {
		v1222 = v1184
		v1225 = v1179
		goto L271
	} else {
		goto L283
	}
L283:
	;
	v1205 = *(*int32)(unsafe.Add(mBase, uint32(v93)+12))
	v1208 = v1202 + v1205<<(uint(int32(2))%32)
	v1209 = *(*int32)(unsafe.Add(mBase, uint32(v1208)))
	if v1184 != v1209 {
		v1222 = v1184
		v1225 = v1179
		goto L271
	} else {
		goto L284
	}
L284:
	;
	v1211 = *(*int32)(unsafe.Add(mBase, uint32(v1209)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v93)+28)) = v1211
	v1213 = *(*int32)(unsafe.Add(mBase, uint32(v1208)))
	v1214 = *(*int32)(unsafe.Add(mBase, uint32(v1213)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v93)+80)) = v1214
	*(*int32)(unsafe.Add(mBase, uint32(v93)+36)) = v1214
	v1217 = *(*int32)(unsafe.Add(mBase, uint32(v1208)))
	v1218 = *(*int32)(unsafe.Add(mBase, uint32(v1217)))
	*(*int32)(unsafe.Add(mBase, uint32(v93)+4)) = v1218
	v1220 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1214))))
	*(*uint8)(unsafe.Add(mBase, uint32(v93)+24)) = uint8(v1220)
	v1222 = v1184
	v1225 = v1179
	goto L271
L285:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1222)+24)) = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[155])) = v1225
	v1242 = *(*int32)(unsafe.Add(mBase, uint32(v93)+20))
	v1243 = *(*int32)(unsafe.Add(mBase, uint32(v93)+12))
	v1246 = v1242 + v1243<<(uint(int32(2))%32)
	v1247 = *(*int32)(unsafe.Add(mBase, uint32(v1246)))
	v1248 = *(*int32)(unsafe.Add(mBase, uint32(v1247)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v93)+28)) = v1248
	v1250 = *(*int32)(unsafe.Add(mBase, uint32(v1246)))
	v1251 = *(*int32)(unsafe.Add(mBase, uint32(v1250)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v93)+36)) = v1251
	*(*int32)(unsafe.Add(mBase, uint32(v93)+80)) = v1251
	v1254 = *(*int32)(unsafe.Add(mBase, uint32(v1246)))
	v1255 = *(*int32)(unsafe.Add(mBase, uint32(v1254)))
	*(*int32)(unsafe.Add(mBase, uint32(v93)+4)) = v1255
	v1257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1251))))
	*(*uint8)(unsafe.Add(mBase, uint32(v93)+24)) = uint8(v1257)
	v1274 = int32(1)
	goto L266
L286:
	;
	v1230 = *(*int32)(unsafe.Add(mBase, uint32(v93)+12))
	v1234 = *(*int32)(unsafe.Add(mBase, uint32(v1229+v1230<<(uint(int32(2))%32))))
	if v1222 == v1234 {
		goto L285
	} else {
		goto L289
	}
L287:
	;
	goto L288
L288:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1222)+32)) = int64(1)
	goto L285
L289:
	;
	goto L288
L290:
	;
	v1287 = v1276 + int32(base.Ui32(v1275)>>(uint(int32(1))%32))
	v1288 = *(*int32)(unsafe.Add(mBase, uint32(v1282)+4))
	if v1288 != 0 {
		goto L294
	} else {
		goto L295
	}
L291:
	;
	v1317 = v1276
	v1318 = v1277
	v1319 = v1278
	goto L292
L292:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v93)+28)) = v1317
	v1321 = int32(2)
	v1324 = *(*int32)(unsafe.Add(mBase, uint32(v1318+v1319<<(uint(v1321)%32))))
	v1325 = *(*int32)(unsafe.Add(mBase, uint32(v1324)+4))
	v1327 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1325+v1317))) = uint8(v1327)
	v1329 = *(*int32)(unsafe.Add(mBase, uint32(v93)+20))
	v1330 = *(*int32)(unsafe.Add(mBase, uint32(v93)+12))
	v1334 = *(*int32)(unsafe.Add(mBase, uint32(v1329+v1330<<(uint(v1321)%32))))
	v1335 = *(*int32)(unsafe.Add(mBase, uint32(v1334)+4))
	v1336 = *(*int32)(unsafe.Add(mBase, uint32(v93)+28))
	*(*uint8)(unsafe.Add(mBase, uint32(v1335+v1336)+1)) = uint8(v1327)
	v1340 = *(*int32)(unsafe.Add(mBase, uint32(v93)+20))
	v1341 = *(*int32)(unsafe.Add(mBase, uint32(v93)+12))
	v1344 = v1340 + v1341<<(uint(v1321)%32)
	v1345 = *(*int32)(unsafe.Add(mBase, uint32(v1344)))
	v1346 = *(*int32)(unsafe.Add(mBase, uint32(v1345)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v93)+80)) = v1346
	if v1274 == int32(1) {
		v1593 = v1346
		goto L48
	} else {
		goto L300
	}
L293:
	;
	v1294 = *(*int32)(unsafe.Add(mBase, uint32(v93)+20))
	v1295 = *(*int32)(unsafe.Add(mBase, uint32(v93)+12))
	v1296 = int32(2)
	v1299 = *(*int32)(unsafe.Add(mBase, uint32(v1294+v1295<<(uint(v1296)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v1299)+4)) = v1293
	v1301 = *(*int32)(unsafe.Add(mBase, uint32(v93)+20))
	v1302 = *(*int32)(unsafe.Add(mBase, uint32(v93)+12))
	v1306 = *(*int32)(unsafe.Add(mBase, uint32(v1301+v1302<<(uint(v1296)%32))))
	v1307 = *(*int32)(unsafe.Add(mBase, uint32(v1306)+4))
	if v1307 == int32(0) {
		goto L50
	} else {
		goto L299
	}
L294:
	;
	v1289 = F_repalloc(m, v1288, v1287)
	mBase = m.M
	v1290 = m.ExcPending
	if v1290 != 0 {
		goto L18
	} else {
		goto L297
	}
L295:
	;
	goto L296
L296:
	;
	v1291 = F_palloc(m, v1287)
	mBase = m.M
	v1292 = m.ExcPending
	if v1292 != 0 {
		goto L18
	} else {
		goto L298
	}
L297:
	;
	v1293 = v1289
	goto L293
L298:
	;
	v1293 = v1291
	goto L293
L299:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1306)+12)) = v1287 - int32(2)
	v1313 = *(*int32)(unsafe.Add(mBase, uint32(v93)+12))
	v1314 = *(*int32)(unsafe.Add(mBase, uint32(v93)+20))
	v1315 = *(*int32)(unsafe.Add(mBase, uint32(v93)+28))
	v1317 = v1315 + v765
	v1318 = v1314
	v1319 = v1313
	goto L292
L300:
	;
	switch v1274 - int32(1) {
	case 0:
		goto L51
	case 1:
		goto L301
	default:
		goto L302
	}
L301:
	;
	v1460 = *(*int32)(unsafe.Add(mBase, uint32(v93)+28))
	v1461 = *(*int32)(unsafe.Add(mBase, uint32(v1344)))
	v1462 = *(*int32)(unsafe.Add(mBase, uint32(v1461)+4))
	v1463 = v1462
	v1467 = v1346
	v1470 = v1460
	goto L52
L302:
	;
	v1355 = v1346 + (v515 ^ int32(-1)) + v226
	*(*int32)(unsafe.Add(mBase, uint32(v93)+36)) = v1355
	v1357 = *(*int32)(unsafe.Add(mBase, uint32(v93)+44))
	if base.Ui32(v1355) <= base.Ui32(v1346) {
		v109 = v1357
		v113 = v1346
		v115 = v1355
		goto L27
	} else {
		goto L303
	}
L303:
	;
	v1359 = v1357
	v1366 = v1346
	goto L304
L304:
	;
	v1373 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1366))))
	if v1373 != 0 {
		goto L306
	} else {
		goto L307
	}
L305:
	;
	v109 = v1456
	v113 = v1346
	v115 = v1355
	goto L27
L306:
	;
	v1376 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1373)+uint32(_consts[669]))))
	v1377 = v1376
	goto L308
L307:
	;
	v1377 = int32(1)
	goto L308
L308:
	;
	v1382 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1359<<(uint(int32(1))%32))+uint32(_consts[670]))))
	if v1382 != 0 {
		goto L309
	} else {
		goto L310
	}
L309:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v93)+68)) = v1366
	*(*int32)(unsafe.Add(mBase, uint32(v93)+64)) = v1359
	goto L311
L310:
	;
	goto L311
L311:
	;
	v1386 = v1377 & int32(255)
	v1387 = int32(1)
	v1391 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1359<<(uint(v1387)%32))+uint32(_consts[671]))))
	v1392 = v1386 + v1391
	v1397 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1392<<(uint(v1387)%32))+uint32(_consts[672]))))
	if v1397 != v1359 {
		goto L312
	} else {
		goto L313
	}
L312:
	;
	v1399 = v1359
	v1401 = v1377
	v1402 = v1386
	goto L315
L313:
	;
	v1444 = v1392
	goto L314
L314:
	;
	v1452 = int32(1)
	v1456 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1444<<(uint(v1452)%32))+uint32(_consts[673]))))
	v1458 = v1366 + v1452
	if v1355 != v1458 {
		v1359 = v1456
		v1366 = v1458
		goto L304
	} else {
		goto L321
	}
L315:
	;
	v1416 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1399<<(uint(int32(1))%32))+uint32(_consts[674]))))
	if int32(285) <= v1416 {
		goto L317
	} else {
		goto L318
	}
L316:
	;
	v1444 = v1430
	goto L314
L317:
	;
	v1421 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1402)+uint32(_consts[675]))))
	v1422 = v1421
	goto L319
L318:
	;
	v1422 = v1401
	goto L319
L319:
	;
	v1424 = v1422 & int32(255)
	v1425 = int32(1)
	v1429 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1416<<(uint(v1425)%32))+uint32(_consts[671]))))
	v1430 = v1424 + v1429
	v1435 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1430<<(uint(v1425)%32))+uint32(_consts[672]))))
	if v1435 != v1416&int32(65535) {
		v1399 = v1416
		v1401 = v1422
		v1402 = v1424
		goto L315
	} else {
		goto L320
	}
L320:
	;
	goto L316
L321:
	;
	goto L305
L322:
	;
	v1480 = v1478
	v1487 = v1467
	goto L323
L323:
	;
	v1494 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1487))))
	if v1494 != 0 {
		goto L325
	} else {
		goto L326
	}
L324:
	;
	v220 = v1577
	v226 = v1476
	v228 = v1467
	goto L44
L325:
	;
	v1497 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1494)+uint32(_consts[669]))))
	v1498 = v1497
	goto L327
L326:
	;
	v1498 = int32(1)
	goto L327
L327:
	;
	v1503 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1480<<(uint(int32(1))%32))+uint32(_consts[670]))))
	if v1503 != 0 {
		goto L328
	} else {
		goto L329
	}
L328:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v93)+68)) = v1487
	*(*int32)(unsafe.Add(mBase, uint32(v93)+64)) = v1480
	goto L330
L329:
	;
	goto L330
L330:
	;
	v1507 = v1498 & int32(255)
	v1508 = int32(1)
	v1512 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1480<<(uint(v1508)%32))+uint32(_consts[671]))))
	v1513 = v1507 + v1512
	v1518 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1513<<(uint(v1508)%32))+uint32(_consts[672]))))
	if v1518 != v1480 {
		goto L331
	} else {
		goto L332
	}
L331:
	;
	v1520 = v1480
	v1522 = v1498
	v1523 = v1507
	goto L334
L332:
	;
	v1565 = v1513
	goto L333
L333:
	;
	v1573 = int32(1)
	v1577 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1565<<(uint(v1573)%32))+uint32(_consts[673]))))
	v1579 = v1487 + v1573
	if v1579 != v1476 {
		v1480 = v1577
		v1487 = v1579
		goto L323
	} else {
		goto L340
	}
L334:
	;
	v1537 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1520<<(uint(int32(1))%32))+uint32(_consts[674]))))
	if int32(285) <= v1537 {
		goto L336
	} else {
		goto L337
	}
L335:
	;
	v1565 = v1551
	goto L333
L336:
	;
	v1542 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1523)+uint32(_consts[675]))))
	v1543 = v1542
	goto L338
L337:
	;
	v1543 = v1522
	goto L338
L338:
	;
	v1545 = v1543 & int32(255)
	v1546 = int32(1)
	v1550 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1537<<(uint(v1546)%32))+uint32(_consts[671]))))
	v1551 = v1545 + v1550
	v1556 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1551<<(uint(v1546)%32))+uint32(_consts[672]))))
	if v1556 != v1537&int32(65535) {
		v1520 = v1537
		v1522 = v1543
		v1523 = v1545
		goto L334
	} else {
		goto L339
	}
L339:
	;
	goto L335
L340:
	;
	goto L324
L341:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L342:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
