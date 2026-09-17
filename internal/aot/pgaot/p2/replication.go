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
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationOriginExitCleanup[0]))
	if v4 != 0 {
		v6 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationOriginExitCleanup[1]))
		v10 = F_LWLockAcquire(m, v6+int32(_a_F_ReplicationOriginExitCleanup_0), int32(0))
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationOriginExitCleanup[0]))
			v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+24))
			v16 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationOriginExitCleanup[2]))
			if v14 != v16 {
				v19 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationOriginExitCleanup[1]))
				F_LWLockRelease(m, v19+int32(_a_F_ReplicationOriginExitCleanup_0))
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
				*(*int32)(unsafe.Add(mBase, _c_F_ReplicationOriginExitCleanup[0])) = v24
				v30 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationOriginExitCleanup[1]))
				F_LWLockRelease(m, v30+int32(_a_F_ReplicationOriginExitCleanup_0))
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
	var v31 int32
	_ = v31
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
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v174 int32
	_ = v174
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v229 int32
	_ = v229
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v270 int32
	_ = v270
	var v275 int32
	_ = v275
	var v286 int32
	_ = v286
	var v290 int32
	_ = v290
	var v294 int32
	_ = v294
	var v297 int32
	_ = v297
	var v303 int32
	_ = v303
	var v308 int32
	_ = v308
	v10 = m.G0
	v12 = v10 - int32(80)
	m.G0 = v12
	v15 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotAcquire[0]))
	v19 = F_LWLockAcquire(m, v15+int32(_a_F_ReplicationSlotAcquire_0), int32(1))
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
	v22 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotAcquire[1]))
	if v22 <= int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v286 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotAcquire[0]))
	F_LWLockRelease(m, v286+int32(_a_F_ReplicationSlotAcquire_0))
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L1
	} else {
		goto L94
	}
L4:
	;
	v31 = v22
	goto L6
L5:
	;
	if l1 == int32(0) {
		goto L58
	} else {
		goto L59
	}
L6:
	;
	v36 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotAcquire[2]))
	v42 = int32(0)
	goto L8
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L1
	} else {
		goto L54
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
	v87 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ReplicationSlotAcquire[3])))
	if v87 == int32(1) {
		goto L25
	} else {
		goto L26
	}
L10:
	;
	goto L9
L11:
	;
	v53 = v48 + int32(24)
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53))))
	if base.B2i32(v56 == int32(0))|base.B2i32(v56 != v59) != 0 {
		v77 = v56
		v78 = v59
		goto L15
	} else {
		goto L16
	}
L12:
	;
	goto L13
L13:
	;
	v84 = v42 + int32(1)
	if v84 != v31 {
		v42 = v84
		goto L8
	} else {
		goto L22
	}
L14:
	;
	if v77-v78 == int32(0) {
		goto L10
	} else {
		goto L21
	}
L15:
	;
	goto L14
L16:
	;
	v62 = l0
	v63 = v53
	goto L17
L17:
	;
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63)+1)))
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62)+1)))
	if v67 == int32(0) {
		v77 = v67
		v78 = v66
		goto L15
	} else {
		goto L19
	}
L18:
	;
	v77 = v67
	v78 = v66
	goto L15
L19:
	;
	v70 = int32(1)
	if v67 == v66 {
		v62 = v62 + v70
		v63 = v63 + v70
		goto L17
	} else {
		goto L20
	}
L20:
	;
	goto L18
L21:
	;
	goto L13
L22:
	;
	goto L3
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48))) = int32(0)
	v133 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotAcquire[0]))
	F_LWLockRelease(m, v133+int32(_a_F_ReplicationSlotAcquire_0))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L1
	} else {
		goto L45
	}
L24:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v48)+272)) = int64(0)
	v129 = v126
	goto L23
L25:
	;
	if l1 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L26:
	;
	goto L27
L27:
	;
	v115 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotAcquire[4]))
	*(*int32)(unsafe.Add(mBase, uint32(v48)+8)) = v115
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
	*(*int32)(unsafe.Add(mBase, uint32(v48))) = int32(1)
	if v117 != 0 {
		goto L40
	} else {
		goto L41
	}
L28:
	;
	F_ConditionVariablePrepareToSleep(m, v48+int32(224))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L1
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
	*(*int32)(unsafe.Add(mBase, uint32(v48))) = int32(1)
	if v96 != 0 {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	goto L30
L32:
	;
	F_s_lock(m, v48, int32(_a_F_ReplicationSlotAcquire_1), int32(632), int32(_a_F_ReplicationSlotAcquire_2))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L1
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v48)+8))
	if v104 == int32(0) {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	goto L34
L36:
	;
	v108 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotAcquire[4]))
	*(*int32)(unsafe.Add(mBase, uint32(v48)+8)) = v108
	v110 = v108
	goto L38
L37:
	;
	v110 = v104
	goto L38
L38:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v48)+112))
	if v111 == int32(0) {
		v126 = v110
		goto L24
	} else {
		goto L39
	}
L39:
	;
	v129 = v110
	goto L23
L40:
	;
	F_s_lock(m, v48, int32(_a_F_ReplicationSlotAcquire_3), int32(251), int32(_a_F_ReplicationSlotAcquire_4))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L1
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v48)+112))
	if v125 != 0 {
		v129 = v115
		goto L23
	} else {
		goto L44
	}
L43:
	;
	goto L42
L44:
	;
	v126 = v115
	goto L24
L45:
	;
	v139 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotAcquire[4]))
	if v129 == v139 {
		goto L5
	} else {
		goto L46
	}
L46:
	;
	if l1 == int32(0) {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	F_ConditionVariableSleep(m, v48+int32(224), int32(134217777))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L1
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	goto L7
L50:
	;
	F_ConditionVariableCancelSleep(m)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	v151 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotAcquire[0]))
	v155 = F_LWLockAcquire(m, v151+int32(_a_F_ReplicationSlotAcquire_0), int32(1))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	v158 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotAcquire[1]))
	if int32(0) < v158 {
		v31 = v158
		goto L6
	} else {
		goto L53
	}
L53:
	;
	goto L3
L54:
	;
	F_errcode(m, int32(100663621))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+52)) = v129
	*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = v53
	F_errmsg(m, int32(_a_F_ReplicationSlotAcquire_5), v12+int32(48))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	F_errfinish(m, int32(_a_F_ReplicationSlotAcquire_1), int32(665), int32(_a_F_ReplicationSlotAcquire_2))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L58:
	;
	F_ConditionVariableCancelSleep(m)
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L1
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotAcquire[5])) = v48
	if l2 != 0 {
		goto L63
	} else {
		goto L64
	}
L61:
	;
	goto L60
L62:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L1
	} else {
		goto L85
	}
L63:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v48)+112))
	if v186 != 0 {
		goto L62
	} else {
		goto L66
	}
L64:
	;
	goto L65
L65:
	;
	F_ConditionVariableBroadcast(m, v48+int32(224))
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L1
	} else {
		goto L67
	}
L66:
	;
	goto L65
L67:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v48)+88))
	if v191 != 0 {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v195 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotAcquire[2]))
	v198 = base.I32_div_s(v48-v195, int32(288))
	goto L71
L69:
	;
	goto L70
L70:
	;
	v205 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ReplicationSlotAcquire[6])))
	if v205 != int32(1) {
		goto L73
	} else {
		goto L74
	}
L71:
	;
	v202 = F_pgstat_get_entry_ref(m, int32(4), int32(0), base.I64_extend_i32_s(v198), int32(1), int32(0))
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	goto L70
L73:
	;
	m.G0 = v12 + int32(80)
	return
L74:
	;
	v211 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ReplicationSlotAcquire[7])))
	if v211 != 0 {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v212 = int32(15)
	goto L77
L76:
	;
	v212 = int32(14)
	goto L77
L77:
	;
	v214 = F_errstart(m, v212, int32(0))
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L1
	} else {
		goto L78
	}
L78:
	;
	if v214 == int32(0) {
		goto L73
	} else {
		goto L79
	}
L79:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v48)+88))
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v53
	if v218 != 0 {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v222 = int32(_a_F_ReplicationSlotAcquire_6)
	goto L82
L81:
	;
	v222 = int32(_a_F_ReplicationSlotAcquire_7)
	goto L82
L82:
	;
	F_errmsg(m, v222, v12)
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L1
	} else {
		goto L83
	}
L83:
	;
	F_errfinish(m, int32(_a_F_ReplicationSlotAcquire_1), int32(705), int32(_a_F_ReplicationSlotAcquire_2))
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L1
	} else {
		goto L84
	}
L84:
	;
	goto L73
L85:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L1
	} else {
		goto L86
	}
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v53
	F_errmsg(m, int32(_a_F_ReplicationSlotAcquire_8), v12+int32(32))
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L1
	} else {
		goto L87
	}
L87:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v48)+112))
	if base.B2i32(int32(base.Ui32(int32(279))>>(uint(v247)%32))&int32(1) == int32(0))|base.B2i32(base.Ui32(int32(8)) < base.Ui32(v247)) != 0 {
		goto L89
	} else {
		goto L90
	}
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v264
	F_errdetail(m, int32(_a_F_ReplicationSlotAcquire_9), v12+int32(16))
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L1
	} else {
		goto L92
	}
L89:
	;
	v264 = int32(_a_F_ReplicationSlotAcquire_10)
	goto L91
L90:
	;
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v247<<(uint(int32(2))%32))+uint32(_c_F_ReplicationSlotAcquire[8])))
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v262)+4))
	v264 = v263
	goto L91
L91:
	;
	goto L88
L92:
	;
	F_errfinish(m, int32(_a_F_ReplicationSlotAcquire_1), int32(684), int32(_a_F_ReplicationSlotAcquire_2))
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L1
	} else {
		goto L93
	}
L93:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L94:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L1
	} else {
		goto L95
	}
L95:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L1
	} else {
		goto L96
	}
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+64)) = l0
	F_errmsg(m, int32(_a_F_ReplicationSlotAcquire_11), v12-int32(-64))
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L1
	} else {
		goto L97
	}
L97:
	;
	F_errfinish(m, int32(_a_F_ReplicationSlotAcquire_1), int32(610), int32(_a_F_ReplicationSlotAcquire_2))
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L1
	} else {
		goto L98
	}
L98:
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
	v7 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotCleanup[0]))
	v11 = F_LWLockAcquire(m, v7+int32(_a_F_ReplicationSlotCleanup_0), int32(1))
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
	v14 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotCleanup[1]))
	if v14 <= int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v93 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotCleanup[0]))
	F_LWLockRelease(m, v93+int32(_a_F_ReplicationSlotCleanup_0))
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
	v24 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotCleanup[2]))
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
	v65 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotCleanup[0]))
	F_LWLockRelease(m, v65+int32(_a_F_ReplicationSlotCleanup_0))
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
	F_s_lock(m, v32, int32(_a_F_ReplicationSlotCleanup_1), int32(820), int32(_a_F_ReplicationSlotCleanup_2))
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
	v46 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotCleanup[3]))
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
	v54 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotCleanup[2]))
	v56 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotCleanup[1]))
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
	v77 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotCleanup[0]))
	v81 = F_LWLockAcquire(m, v77+int32(_a_F_ReplicationSlotCleanup_0), int32(1))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	v84 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotCleanup[1]))
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
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	v4 = m.G0
	v6 = v4 - int32(1040)
	m.G0 = v6
	v9 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotPersist[0]))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = int32(1)
	if v10 != 0 {
		F_s_lock(m, v9, int32(_a_F_ReplicationSlotPersist_0), int32(1125), int32(_a_F_ReplicationSlotPersist_1))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return
		} else {
			v18 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v9))) = v18
			*(*int32)(unsafe.Add(mBase, uint32(v9)+92)) = v18
			v23 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotPersist[0]))
			v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
			*(*int32)(unsafe.Add(mBase, uint32(v23))) = int32(1)
			if v24 != 0 {
				F_s_lock(m, v23, int32(_a_F_ReplicationSlotPersist_0), int32(1107), int32(_a_F_ReplicationSlotPersist_2))
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return
				} else {
					v32 = int32(_a_F_ReplicationSlotPersist_3)
					v33 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotPersist[0]))
					v34 = int32(257)
					*(*uint16)(unsafe.Add(mBase, uint32(v33)+12)) = uint16(v34)
					*(*int32)(unsafe.Add(mBase, uint32(v23))) = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v6))) = int32(_a_F_ReplicationSlotPersist_4)
					v41 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotPersist[0]))
					*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v41 + int32(24)
					v46 = v6 + int32(16)
					v48 = F_pg_sprintf(m, v46, int32(_a_F_ReplicationSlotPersist_5), v6)
					mBase = m.M
					v49 = m.ExcPending
					if v49 != 0 {
						return
					} else {
						v51 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotPersist[0]))
						F_SaveSlotToPath(m, v51, v46, int32(21))
						mBase = m.M
						v54 = m.ExcPending
						if v54 != 0 {
							return
						} else {
							m.G0 = v6 + int32(1040)
							return
						}
					}
				}
			} else {
				v32 = int32(_a_F_ReplicationSlotPersist_3)
				v33 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotPersist[0]))
				v34 = int32(257)
				*(*uint16)(unsafe.Add(mBase, uint32(v33)+12)) = uint16(v34)
				*(*int32)(unsafe.Add(mBase, uint32(v23))) = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v6))) = int32(_a_F_ReplicationSlotPersist_4)
				v41 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotPersist[0]))
				*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v41 + int32(24)
				v46 = v6 + int32(16)
				v48 = F_pg_sprintf(m, v46, int32(_a_F_ReplicationSlotPersist_5), v6)
				mBase = m.M
				v49 = m.ExcPending
				if v49 != 0 {
					return
				} else {
					v51 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotPersist[0]))
					F_SaveSlotToPath(m, v51, v46, int32(21))
					mBase = m.M
					v54 = m.ExcPending
					if v54 != 0 {
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
		v23 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotPersist[0]))
		v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
		*(*int32)(unsafe.Add(mBase, uint32(v23))) = int32(1)
		if v24 != 0 {
			F_s_lock(m, v23, int32(_a_F_ReplicationSlotPersist_0), int32(1107), int32(_a_F_ReplicationSlotPersist_2))
			mBase = m.M
			v31 = m.ExcPending
			if v31 != 0 {
				return
			} else {
				v32 = int32(_a_F_ReplicationSlotPersist_3)
				v33 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotPersist[0]))
				v34 = int32(257)
				*(*uint16)(unsafe.Add(mBase, uint32(v33)+12)) = uint16(v34)
				*(*int32)(unsafe.Add(mBase, uint32(v23))) = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v6))) = int32(_a_F_ReplicationSlotPersist_4)
				v41 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotPersist[0]))
				*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v41 + int32(24)
				v46 = v6 + int32(16)
				v48 = F_pg_sprintf(m, v46, int32(_a_F_ReplicationSlotPersist_5), v6)
				mBase = m.M
				v49 = m.ExcPending
				if v49 != 0 {
					return
				} else {
					v51 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotPersist[0]))
					F_SaveSlotToPath(m, v51, v46, int32(21))
					mBase = m.M
					v54 = m.ExcPending
					if v54 != 0 {
						return
					} else {
						m.G0 = v6 + int32(1040)
						return
					}
				}
			}
		} else {
			v32 = int32(_a_F_ReplicationSlotPersist_3)
			v33 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotPersist[0]))
			v34 = int32(257)
			*(*uint16)(unsafe.Add(mBase, uint32(v33)+12)) = uint16(v34)
			*(*int32)(unsafe.Add(mBase, uint32(v23))) = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v6))) = int32(_a_F_ReplicationSlotPersist_4)
			v41 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotPersist[0]))
			*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v41 + int32(24)
			v46 = v6 + int32(16)
			v48 = F_pg_sprintf(m, v46, int32(_a_F_ReplicationSlotPersist_5), v6)
			mBase = m.M
			v49 = m.ExcPending
			if v49 != 0 {
				return
			} else {
				v51 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotPersist[0]))
				F_SaveSlotToPath(m, v51, v46, int32(21))
				mBase = m.M
				v54 = m.ExcPending
				if v54 != 0 {
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
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v43 int32
	_ = v43
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v92 int32
	_ = v92
	var v101 int32
	_ = v101
	v9 = m.G0
	v11 = v9 - int32(48)
	m.G0 = v11
	v13 = F_strlen(m, l0)
	mBase = m.M
	if v13 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	m.G0 = v11 + int32(48)
	return v101
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v92
	v101 = int32(0)
	goto L1
L3:
	;
	v92 = int32(0)
	goto L2
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(33579140)
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = l0
	v20 = F_psprintf(m, int32(_a_F_ReplicationSlotValidateNameInternal_0), v11)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	if base.Ui32(v13) <= base.Ui32(int32(63)) {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	return int32(0)
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v20
	goto L3
L9:
	;
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v27 == int32(0) {
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
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = l0
	v80 = F_psprintf(m, int32(_a_F_ReplicationSlotValidateNameInternal_1), v11+int32(16))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L7
	} else {
		goto L23
	}
L12:
	;
	v101 = int32(1)
	goto L1
L13:
	;
	goto L14
L14:
	;
	v34 = l0
	v36 = v27
	goto L15
L15:
	;
	v43 = int32(255)
	if base.B2i32(v36 == int32(95))|base.B2i32(base.Ui32((v36-int32(97))&v43) < base.Ui32(int32(26)))|base.B2i32(base.Ui32((v36-int32(48))&v43) < base.Ui32(int32(10))) == int32(0) {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v101 = v70
	goto L1
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(33579140)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = l0
	v63 = F_psprintf(m, int32(_a_F_ReplicationSlotValidateNameInternal_2), v11+int32(32))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L7
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v70 = int32(1)
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+1)))
	if v73 != 0 {
		v34 = v34 + v70
		v36 = v73
		goto L15
	} else {
		goto L22
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v63
	v68 = F_psprintf(m, int32(_a_F_ReplicationSlotValidateNameInternal_3), int32(0))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L7
	} else {
		goto L21
	}
L21:
	;
	v92 = v68
	goto L2
L22:
	;
	goto L16
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v80
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
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	var v186 int32
	_ = v186
	var v193 int32
	_ = v193
	var v199 int32
	_ = v199
	var v210 int32
	_ = v210
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v257 int32
	_ = v257
	var v274 int32
	_ = v274
	var v278 int64
	_ = v278
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v295 int32
	_ = v295
	var v296 int64
	_ = v296
	var v297 int64
	_ = v297
	var v305 int32
	_ = v305
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v330 int32
	_ = v330
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v342 int32
	_ = v342
	var v345 int32
	_ = v345
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v374 int32
	_ = v374
	var v376 int32
	_ = v376
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v395 int32
	_ = v395
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v412 int32
	_ = v412
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v422 int32
	_ = v422
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
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
	var v439 int32
	_ = v439
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v445 int32
	_ = v445
	var v448 int32
	_ = v448
	var v455 int32
	_ = v455
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v466 int32
	_ = v466
	var v469 int32
	_ = v469
	var v473 int32
	_ = v473
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v482 int32
	_ = v482
	var v484 int32
	_ = v484
	var v486 int32
	_ = v486
	var v489 int32
	_ = v489
	var v501 int32
	_ = v501
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v518 int32
	_ = v518
	var v528 int32
	_ = v528
	var v535 int32
	_ = v535
	var v539 int32
	_ = v539
	var v541 int32
	_ = v541
	var v543 int32
	_ = v543
	var v560 int32
	_ = v560
	var v563 int32
	_ = v563
	var v567 int32
	_ = v567
	var v569 int32
	_ = v569
	var v574 int32
	_ = v574
	var v576 int32
	_ = v576
	var v589 int32
	_ = v589
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v599 int32
	_ = v599
	var v601 int32
	_ = v601
	var v606 int32
	_ = v606
	var v610 int32
	_ = v610
	var v627 int32
	_ = v627
	var v631 int32
	_ = v631
	var v636 int32
	_ = v636
	var v637 int32
	_ = v637
	var v645 int32
	_ = v645
	var v648 int32
	_ = v648
	var v649 int32
	_ = v649
	var v658 int32
	_ = v658
	var v660 int32
	_ = v660
	var v663 int32
	_ = v663
	var v671 int32
	_ = v671
	var v673 int32
	_ = v673
	var v675 int32
	_ = v675
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
	var v688 int32
	_ = v688
	var v690 int32
	_ = v690
	var v692 int32
	_ = v692
	var v696 int32
	_ = v696
	var v698 int32
	_ = v698
	var v710 int32
	_ = v710
	var v712 int32
	_ = v712
	var v715 int32
	_ = v715
	var v723 int32
	_ = v723
	var v725 int32
	_ = v725
	var v730 int32
	_ = v730
	var v745 int32
	_ = v745
	var v746 int32
	_ = v746
	var v750 int32
	_ = v750
	var v753 int32
	_ = v753
	var v756 int32
	_ = v756
	var v757 int32
	_ = v757
	var v764 int32
	_ = v764
	var v767 int32
	_ = v767
	var v773 int32
	_ = v773
	var v774 int32
	_ = v774
	var v775 int32
	_ = v775
	var v776 int32
	_ = v776
	var v779 int32
	_ = v779
	var v782 int32
	_ = v782
	var v783 int32
	_ = v783
	var v786 int32
	_ = v786
	var v793 int32
	_ = v793
	var v798 int32
	_ = v798
	var v800 int32
	_ = v800
	var v804 int32
	_ = v804
	var v806 int32
	_ = v806
	var v809 int32
	_ = v809
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
	var v819 int32
	_ = v819
	var v821 int32
	_ = v821
	var v822 int32
	_ = v822
	var v826 int32
	_ = v826
	var v827 int32
	_ = v827
	var v828 int32
	_ = v828
	var v831 int32
	_ = v831
	var v833 int32
	_ = v833
	var v844 int32
	_ = v844
	var v847 int32
	_ = v847
	var v848 int32
	_ = v848
	var v853 int32
	_ = v853
	var v862 int32
	_ = v862
	var v863 int32
	_ = v863
	var v864 int32
	_ = v864
	var v866 int32
	_ = v866
	var v867 int32
	_ = v867
	var v868 int32
	_ = v868
	var v872 int32
	_ = v872
	var v873 int32
	_ = v873
	var v878 int32
	_ = v878
	var v880 int32
	_ = v880
	var v881 int32
	_ = v881
	var v882 int32
	_ = v882
	var v891 int32
	_ = v891
	var v892 int32
	_ = v892
	var v893 int32
	_ = v893
	var v897 int32
	_ = v897
	var v898 int32
	_ = v898
	var v901 int32
	_ = v901
	var v905 int32
	_ = v905
	var v910 int32
	_ = v910
	var v911 int32
	_ = v911
	var v915 int32
	_ = v915
	var v916 int32
	_ = v916
	var v919 int32
	_ = v919
	var v920 int32
	_ = v920
	var v921 int32
	_ = v921
	var v926 int32
	_ = v926
	var v936 int32
	_ = v936
	var v937 int32
	_ = v937
	var v946 int32
	_ = v946
	var v952 int32
	_ = v952
	var v953 int32
	_ = v953
	var v957 int32
	_ = v957
	var v958 int32
	_ = v958
	var v962 int32
	_ = v962
	var v963 int32
	_ = v963
	var v966 int32
	_ = v966
	var v967 int32
	_ = v967
	var v968 int32
	_ = v968
	var v971 int32
	_ = v971
	var v979 int32
	_ = v979
	var v982 int32
	_ = v982
	var v985 int32
	_ = v985
	var v990 int32
	_ = v990
	var v999 int32
	_ = v999
	var v1000 int32
	_ = v1000
	var v1008 int32
	_ = v1008
	var v1017 int32
	_ = v1017
	var v1018 int32
	_ = v1018
	var v1021 int32
	_ = v1021
	var v1024 int32
	_ = v1024
	var v1025 int32
	_ = v1025
	var v1026 int32
	_ = v1026
	var v1030 int32
	_ = v1030
	var v1032 int32
	_ = v1032
	var v1033 int32
	_ = v1033
	var v1034 int32
	_ = v1034
	var v1035 int32
	_ = v1035
	var v1036 int32
	_ = v1036
	var v1037 int32
	_ = v1037
	var v1042 int32
	_ = v1042
	var v1044 int32
	_ = v1044
	var v1046 int32
	_ = v1046
	var v1048 int32
	_ = v1048
	var v1049 int32
	_ = v1049
	var v1053 int32
	_ = v1053
	var v1057 int32
	_ = v1057
	var v1059 int32
	_ = v1059
	var v1062 int32
	_ = v1062
	var v1069 int32
	_ = v1069
	var v1071 int32
	_ = v1071
	var v1074 int32
	_ = v1074
	var v1077 int32
	_ = v1077
	var v1078 int32
	_ = v1078
	var v1080 int32
	_ = v1080
	var v1082 int32
	_ = v1082
	var v1083 int32
	_ = v1083
	var v1086 int32
	_ = v1086
	var v1087 int32
	_ = v1087
	var v1089 int32
	_ = v1089
	var v1091 int32
	_ = v1091
	var v1094 int32
	_ = v1094
	var v1098 int32
	_ = v1098
	var v1099 int32
	_ = v1099
	var v1103 int32
	_ = v1103
	var v1111 int32
	_ = v1111
	var v1112 int32
	_ = v1112
	var v1115 int32
	_ = v1115
	var v1116 int32
	_ = v1116
	var v1117 int32
	_ = v1117
	var v1119 int32
	_ = v1119
	var v1120 int32
	_ = v1120
	var v1123 int32
	_ = v1123
	var v1124 int32
	_ = v1124
	var v1126 int32
	_ = v1126
	var v1129 int32
	_ = v1129
	var v1130 int32
	_ = v1130
	var v1131 int32
	_ = v1131
	var v1134 int32
	_ = v1134
	var v1143 int32
	_ = v1143
	var v1144 int32
	_ = v1144
	var v1145 int32
	_ = v1145
	var v1146 int32
	_ = v1146
	var v1147 int32
	_ = v1147
	var v1151 int32
	_ = v1151
	var v1152 int32
	_ = v1152
	var v1156 int32
	_ = v1156
	var v1157 int32
	_ = v1157
	var v1158 int32
	_ = v1158
	var v1159 int32
	_ = v1159
	var v1160 int32
	_ = v1160
	var v1161 int32
	_ = v1161
	var v1162 int32
	_ = v1162
	var v1163 int32
	_ = v1163
	var v1164 int32
	_ = v1164
	var v1165 int32
	_ = v1165
	var v1168 int32
	_ = v1168
	var v1170 int32
	_ = v1170
	var v1171 int32
	_ = v1171
	var v1175 int32
	_ = v1175
	var v1176 int32
	_ = v1176
	var v1182 int32
	_ = v1182
	var v1183 int32
	_ = v1183
	var v1184 int32
	_ = v1184
	var v1186 int32
	_ = v1186
	var v1187 int32
	_ = v1187
	var v1189 int32
	_ = v1189
	var v1191 int32
	_ = v1191
	var v1194 int32
	_ = v1194
	var v1195 int32
	_ = v1195
	var v1197 int32
	_ = v1197
	var v1199 int32
	_ = v1199
	var v1200 int32
	_ = v1200
	var v1204 int32
	_ = v1204
	var v1205 int32
	_ = v1205
	var v1206 int32
	_ = v1206
	var v1210 int32
	_ = v1210
	var v1211 int32
	_ = v1211
	var v1214 int32
	_ = v1214
	var v1215 int32
	_ = v1215
	var v1216 int32
	_ = v1216
	var v1224 int32
	_ = v1224
	var v1225 int32
	_ = v1225
	var v1227 int32
	_ = v1227
	var v1230 int32
	_ = v1230
	var v1234 int32
	_ = v1234
	var v1244 int32
	_ = v1244
	var v1245 int32
	_ = v1245
	var v1246 int32
	_ = v1246
	var v1248 int32
	_ = v1248
	var v1251 int32
	_ = v1251
	var v1255 int32
	_ = v1255
	var v1258 int32
	_ = v1258
	var v1259 int32
	_ = v1259
	var v1264 int32
	_ = v1264
	var v1266 int32
	_ = v1266
	var v1268 int32
	_ = v1268
	var v1271 int32
	_ = v1271
	var v1283 int32
	_ = v1283
	var v1286 int32
	_ = v1286
	var v1287 int32
	_ = v1287
	var v1289 int32
	_ = v1289
	var v1290 int32
	_ = v1290
	var v1294 int32
	_ = v1294
	var v1295 int32
	_ = v1295
	var v1300 int32
	_ = v1300
	var v1310 int32
	_ = v1310
	var v1317 int32
	_ = v1317
	var v1321 int32
	_ = v1321
	var v1323 int32
	_ = v1323
	var v1325 int32
	_ = v1325
	var v1326 int32
	_ = v1326
	var v1327 int32
	_ = v1327
	var v1328 int32
	_ = v1328
	var v1331 int32
	_ = v1331
	var v1332 int32
	_ = v1332
	var v1341 int32
	_ = v1341
	var v1343 int32
	_ = v1343
	var v1345 int32
	_ = v1345
	var v1349 int32
	_ = v1349
	var v1359 int32
	_ = v1359
	var v1360 int32
	_ = v1360
	var v1361 int32
	_ = v1361
	var v1366 int32
	_ = v1366
	var v1370 int32
	_ = v1370
	var v1371 int32
	_ = v1371
	var v1375 int32
	_ = v1375
	var v1376 int32
	_ = v1376
	var v1381 int32
	_ = v1381
	var v1383 int32
	_ = v1383
	var v1385 int32
	_ = v1385
	var v1388 int32
	_ = v1388
	var v1400 int32
	_ = v1400
	var v1403 int32
	_ = v1403
	var v1404 int32
	_ = v1404
	var v1406 int32
	_ = v1406
	var v1407 int32
	_ = v1407
	var v1411 int32
	_ = v1411
	var v1412 int32
	_ = v1412
	var v1417 int32
	_ = v1417
	var v1427 int32
	_ = v1427
	var v1434 int32
	_ = v1434
	var v1438 int32
	_ = v1438
	var v1440 int32
	_ = v1440
	var v1444 int32
	_ = v1444
	var v1449 int32
	_ = v1449
	var v1453 int32
	_ = v1453
	var v1466 int32
	_ = v1466
	var v1470 int32
	_ = v1470
	var v1475 int32
	_ = v1475
	var v1485 int32
	_ = v1485
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
	v89 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v89)))
	if v90 == int32(0) {
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
	v33 = *(*int32)(unsafe.Add(mBase, _c_F_replication_yylex[0]))
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
	v39 = *(*int32)(unsafe.Add(mBase, _c_F_replication_yylex[1]))
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
	m.G0 = v1485 + int32(16)
	return v1475
L22:
	;
	v94 = l1
	v103 = v16
	goto L25
L23:
	;
	goto L24
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v89))) = int32(0)
	v1475 = v90
	v1485 = v16
	goto L21
L25:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v94)+36))
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94)+24)))
	*(*uint8)(unsafe.Add(mBase, uint32(v106))) = uint8(v107)
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v94)+44))
	v110 = v109
	v111 = v94
	v113 = v106
	v117 = v106
	v120 = v103
	goto L27
L27:
	;
	v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117))))
	v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123)+uint32(_c_F_replication_yylex[2]))))
	v126 = v110 << (uint(int32(1)) % 32)
	v129 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v126)+uint32(_c_F_replication_yylex[3]))))
	if v129 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v111)+68)) = v117
	*(*int32)(unsafe.Add(mBase, uint32(v111)+64)) = v110
	goto L31
L30:
	;
	goto L31
L31:
	;
	v134 = int32(*(*int16)(unsafe.Add(mBase, uint32(v126)+uint32(_c_F_replication_yylex[4]))))
	v135 = v134 + v124
	v140 = int32(*(*int16)(unsafe.Add(mBase, uint32(v135<<(uint(int32(1))%32))+uint32(_c_F_replication_yylex[5]))))
	if v140 != v110 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v142 = v110
	v144 = v124
	v147 = v124
	goto L35
L33:
	;
	v186 = v135
	goto L34
L34:
	;
	v193 = int32(1)
	v199 = int32(*(*int16)(unsafe.Add(mBase, uint32(v186<<(uint(v193)%32))+uint32(_c_F_replication_yylex[6]))))
	if v199 != int32(284) {
		v110 = v199
		v117 = v117 + v193
		goto L27
	} else {
		goto L41
	}
L35:
	;
	v159 = int32(*(*int16)(unsafe.Add(mBase, uint32(v142<<(uint(int32(1))%32))+uint32(_c_F_replication_yylex[7]))))
	if int32(285) <= v159 {
		goto L37
	} else {
		goto L38
	}
L36:
	;
	v186 = v171
	goto L34
L37:
	;
	v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v147)+uint32(_c_F_replication_yylex[8]))))
	v163 = v162
	goto L39
L38:
	;
	v163 = v144
	goto L39
L39:
	;
	v165 = v163 & int32(255)
	v166 = int32(1)
	v170 = int32(*(*int16)(unsafe.Add(mBase, uint32(v159<<(uint(v166)%32))+uint32(_c_F_replication_yylex[4]))))
	v171 = v165 + v170
	v176 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v171<<(uint(v166)%32))+uint32(_c_F_replication_yylex[5]))))
	if v176 != v159&int32(_a_F_replication_yylex_0) {
		v142 = v159
		v144 = v163
		v147 = v165
		goto L35
	} else {
		goto L40
	}
L40:
	;
	goto L36
L41:
	;
	v210 = v113
	goto L42
L42:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v111)+64))
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v111)+68))
	v217 = v215
	v224 = v216
	v225 = v210
	goto L44
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v111)+80)) = v225
	*(*int32)(unsafe.Add(mBase, uint32(v111)+32)) = v224 - v225
	v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v224))))
	*(*uint8)(unsafe.Add(mBase, uint32(v111)+24)) = uint8(v233)
	v235 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v224))) = uint8(v235)
	*(*int32)(unsafe.Add(mBase, uint32(v111)+36)) = v224
	v242 = int32(*(*int16)(unsafe.Add(mBase, uint32(v217<<(uint(int32(1))%32))+uint32(_c_F_replication_yylex[3]))))
	v245 = v242
	goto L46
L46:
	;
	switch v245 {
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
		v1475 = int32(266)
		v1485 = v120
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
		v94 = v111
		v103 = v120
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
	*(*int32)(unsafe.Add(mBase, uint32(v111)+36)) = v1453
	*(*int32)(unsafe.Add(mBase, uint32(v111)+48)) = int32(0)
	v1466 = *(*int32)(unsafe.Add(mBase, uint32(v111)+44))
	v1470 = base.I32_div_s(v1466-int32(1), int32(2))
	v245 = v1470 + int32(36)
	goto L46
L49:
	;
	F_yy_fatal_error_3(m, int32(_a_F_replication_yylex_1))
	mBase = m.M
	v1449 = m.ExcPending
	if v1449 != 0 {
		goto L18
	} else {
		goto L296
	}
L50:
	;
	F_yy_fatal_error_3(m, int32(_a_F_replication_yylex_2))
	mBase = m.M
	v1444 = m.ExcPending
	if v1444 != 0 {
		goto L18
	} else {
		goto L295
	}
L51:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L52:
	;
	v1341 = v1328 + v1332
	*(*int32)(unsafe.Add(mBase, uint32(v111)+36)) = v1341
	v1343 = *(*int32)(unsafe.Add(mBase, uint32(v111)+44))
	if base.Ui32(v1341) <= base.Ui32(v1331) {
		v217 = v1343
		v224 = v1341
		v225 = v1331
		goto L44
	} else {
		goto L276
	}
L53:
	;
	v1018 = *(*int32)(unsafe.Add(mBase, uint32(v1017)))
	*(*int32)(unsafe.Add(mBase, uint32(v1018)+16)) = v1008
	v1021 = *(*int32)(unsafe.Add(mBase, uint32(v111)+28))
	if v1021 != 0 {
		v1143 = int32(0)
		goto L220
	} else {
		goto L221
	}
L54:
	;
	v999 = *(*int32)(unsafe.Add(mBase, uint32(v111)+20))
	v1000 = *(*int32)(unsafe.Add(mBase, uint32(v111)+12))
	v1008 = v990
	v1017 = v999 + v1000<<(uint(int32(2))%32)
	goto L53
L55:
	;
	F_yy_fatal_error_3(m, int32(_a_F_replication_yylex_3))
	mBase = m.M
	v985 = m.ExcPending
	if v985 != 0 {
		goto L18
	} else {
		goto L219
	}
L56:
	;
	F_yy_fatal_error_3(m, int32(_a_F_replication_yylex_4))
	mBase = m.M
	v982 = m.ExcPending
	if v982 != 0 {
		goto L18
	} else {
		goto L218
	}
L57:
	;
	F_replication_yyerror(m, int32(_a_F_replication_yylex_5))
	mBase = m.M
	v979 = m.ExcPending
	if v979 != 0 {
		goto L18
	} else {
		goto L217
	}
L58:
	;
	v1475 = int32(272)
	v1485 = v120
	goto L21
L59:
	;
	v1475 = int32(265)
	v1485 = v120
	goto L21
L60:
	;
	v1475 = int32(264)
	v1485 = v120
	goto L21
L61:
	;
	v1475 = int32(263)
	v1485 = v120
	goto L21
L62:
	;
	v1475 = int32(262)
	v1485 = v120
	goto L21
L63:
	;
	F_yy_fatal_error_3(m, int32(_a_F_replication_yylex_6))
	mBase = m.M
	v971 = m.ExcPending
	if v971 != 0 {
		goto L18
	} else {
		goto L216
	}
L64:
	;
	v400 = *(*int32)(unsafe.Add(mBase, uint32(v111)+80))
	v401 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111)+24)))
	*(*uint8)(unsafe.Add(mBase, uint32(v224))) = uint8(v401)
	v403 = *(*int32)(unsafe.Add(mBase, uint32(v111)+20))
	v404 = *(*int32)(unsafe.Add(mBase, uint32(v111)+12))
	v407 = v403 + v404<<(uint(int32(2))%32)
	v408 = *(*int32)(unsafe.Add(mBase, uint32(v407)))
	v409 = *(*int32)(unsafe.Add(mBase, uint32(v408)+44))
	if v409 == int32(0) {
		goto L107
	} else {
		goto L108
	}
L65:
	;
	F_yy_fatal_error_3(m, int32(_a_F_replication_yylex_7))
	mBase = m.M
	v399 = m.ExcPending
	if v399 != 0 {
		goto L18
	} else {
		goto L106
	}
L66:
	;
	v1475 = int32(0)
	v1485 = v120
	goto L21
L67:
	;
	F_replication_yyerror(m, int32(_a_F_replication_yylex_8))
	mBase = m.M
	v395 = m.ExcPending
	if v395 != 0 {
		goto L18
	} else {
		goto L105
	}
L68:
	;
	v391 = *(*int32)(unsafe.Add(mBase, uint32(v111)+80))
	v392 = int32(*(*int8)(unsafe.Add(mBase, uint32(v391))))
	v1475 = v392
	v1485 = v120
	goto L21
L69:
	;
	v383 = *(*int32)(unsafe.Add(mBase, uint32(v111)+80))
	v384 = F_strlen(m, v383)
	mBase = m.M
	v386 = F_downcase_truncate_identifier(m, v383, v384, int32(1))
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L18
	} else {
		goto L104
	}
L70:
	;
	v376 = *(*int32)(unsafe.Add(mBase, uint32(v111)))
	v379 = *(*int32)(unsafe.Add(mBase, uint32(v111)+80))
	v380 = *(*int32)(unsafe.Add(mBase, uint32(v111)+32))
	F_appendBinaryStringInfo(m, v376+int32(4), v379, v380)
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L18
	} else {
		goto L103
	}
L71:
	;
	v350 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111)+24)))
	*(*uint8)(unsafe.Add(mBase, uint32(v224))) = uint8(v350)
	*(*int32)(unsafe.Add(mBase, uint32(v111)+80)) = v225
	v353 = int32(1)
	v354 = v225 + v353
	*(*int32)(unsafe.Add(mBase, uint32(v111)+36)) = v354
	*(*int32)(unsafe.Add(mBase, uint32(v111)+32)) = v353
	v358 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v225)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v111)+24)) = uint8(v358)
	v360 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v225)+1)) = uint8(v360)
	*(*int32)(unsafe.Add(mBase, uint32(v111)+44)) = v353
	*(*int32)(unsafe.Add(mBase, uint32(v111)+36)) = v354
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v111)+92))
	v366 = *(*int32)(unsafe.Add(mBase, uint32(v111)))
	v367 = *(*int32)(unsafe.Add(mBase, uint32(v366)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v365))) = v367
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v111)+92))
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v369)))
	v371 = F_strlen(m, v370)
	mBase = m.M
	F_truncate_identifier(m, v370, v371, v353)
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L18
	} else {
		goto L102
	}
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v111)+44)) = int32(3)
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v111)))
	F_initStringInfo(m, v345+int32(4))
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L18
	} else {
		goto L101
	}
L73:
	;
	v336 = *(*int32)(unsafe.Add(mBase, uint32(v111)))
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v111)+80))
	v340 = *(*int32)(unsafe.Add(mBase, uint32(v111)+32))
	F_appendBinaryStringInfo(m, v336+int32(4), v339, v340)
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L18
	} else {
		goto L100
	}
L74:
	;
	v330 = *(*int32)(unsafe.Add(mBase, uint32(v111)))
	F_appendStringInfoChar(m, v330+int32(4), int32(39))
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L18
	} else {
		goto L99
	}
L75:
	;
	v310 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111)+24)))
	*(*uint8)(unsafe.Add(mBase, uint32(v224))) = uint8(v310)
	*(*int32)(unsafe.Add(mBase, uint32(v111)+80)) = v225
	v313 = int32(1)
	v314 = v225 + v313
	*(*int32)(unsafe.Add(mBase, uint32(v111)+36)) = v314
	*(*int32)(unsafe.Add(mBase, uint32(v111)+32)) = v313
	v318 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v225)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v111)+24)) = uint8(v318)
	v320 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v225)+1)) = uint8(v320)
	*(*int32)(unsafe.Add(mBase, uint32(v111)+44)) = v313
	*(*int32)(unsafe.Add(mBase, uint32(v111)+36)) = v314
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v111)+92))
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v111)))
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v326)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v325))) = v327
	v1475 = int32(258)
	v1485 = v120
	goto L21
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v111)+44)) = int32(5)
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v111)))
	F_initStringInfo(m, v305+int32(4))
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L18
	} else {
		goto L98
	}
L77:
	;
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v111)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v120)+4)) = v120 + int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v120))) = v120 + int32(12)
	v291 = F_sscanf(m, v283, int32(_a_F_replication_yylex_9), v120)
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L18
	} else {
		goto L96
	}
L78:
	;
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v111)+80))
	v278 = F_strtox_2(m, v274, int32(0), int32(10), int64(4294967295))
	mBase = m.M
	goto L95
L79:
	;
	v1475 = int32(282)
	v1485 = v120
	goto L21
L80:
	;
	v1475 = int32(271)
	v1485 = v120
	goto L21
L81:
	;
	v1475 = int32(281)
	v1485 = v120
	goto L21
L82:
	;
	v1475 = int32(280)
	v1485 = v120
	goto L21
L83:
	;
	v1475 = int32(279)
	v1485 = v120
	goto L21
L84:
	;
	v1475 = int32(278)
	v1485 = v120
	goto L21
L85:
	;
	v1475 = int32(277)
	v1485 = v120
	goto L21
L86:
	;
	v1475 = int32(275)
	v1485 = v120
	goto L21
L87:
	;
	v1475 = int32(274)
	v1485 = v120
	goto L21
L88:
	;
	v1475 = int32(276)
	v1485 = v120
	goto L21
L89:
	;
	v1475 = int32(273)
	v1485 = v120
	goto L21
L90:
	;
	v1475 = int32(270)
	v1485 = v120
	goto L21
L91:
	;
	v1475 = int32(269)
	v1485 = v120
	goto L21
L92:
	;
	v1475 = int32(268)
	v1485 = v120
	goto L21
L93:
	;
	v1475 = int32(267)
	v1485 = v120
	goto L21
L94:
	;
	v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111)+24)))
	*(*uint8)(unsafe.Add(mBase, uint32(v224))) = uint8(v257)
	v210 = v225
	goto L42
L95:
	;
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v111)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v280))) = base.I32_wrap_i64(v278)
	v1475 = int32(260)
	v1485 = v120
	goto L21
L96:
	;
	if v291 != int32(2) {
		goto L57
	} else {
		goto L97
	}
L97:
	;
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v111)+92))
	v296 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v120)+8)))
	v297 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v120)+12)))
	*(*int64)(unsafe.Add(mBase, uint32(v295))) = v296 | v297<<(uint(int64(32))%64)
	v1475 = int32(261)
	v1485 = v120
	goto L21
L98:
	;
	v94 = v111
	v103 = v120
	goto L25
L99:
	;
	v94 = v111
	v103 = v120
	goto L25
L100:
	;
	v94 = v111
	v103 = v120
	goto L25
L101:
	;
	v94 = v111
	v103 = v120
	goto L25
L102:
	;
	v1475 = int32(259)
	v1485 = v120
	goto L21
L103:
	;
	v94 = v111
	v103 = v120
	goto L25
L104:
	;
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v111)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v388))) = v386
	v1475 = int32(259)
	v1485 = v120
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
	v412 = *(*int32)(unsafe.Add(mBase, uint32(v408)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v111)+28)) = v412
	v414 = *(*int32)(unsafe.Add(mBase, uint32(v407)))
	v415 = *(*int32)(unsafe.Add(mBase, uint32(v111)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v414))) = v415
	v417 = *(*int32)(unsafe.Add(mBase, uint32(v111)+20))
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v111)+12))
	v419 = int32(2)
	v422 = *(*int32)(unsafe.Add(mBase, uint32(v417+v418<<(uint(v419)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v422)+44)) = int32(1)
	v425 = *(*int32)(unsafe.Add(mBase, uint32(v111)+20))
	v426 = *(*int32)(unsafe.Add(mBase, uint32(v111)+12))
	v430 = *(*int32)(unsafe.Add(mBase, uint32(v425+v426<<(uint(v419)%32))))
	v431 = v430
	v432 = v425
	v433 = v426
	goto L109
L108:
	;
	v431 = v408
	v432 = v403
	v433 = v404
	goto L109
L109:
	;
	v434 = *(*int32)(unsafe.Add(mBase, uint32(v111)+36))
	v435 = *(*int32)(unsafe.Add(mBase, uint32(v431)+4))
	v436 = *(*int32)(unsafe.Add(mBase, uint32(v111)+28))
	v437 = v435 + v436
	if base.Ui32(v434) <= base.Ui32(v437) {
		goto L110
	} else {
		goto L111
	}
L110:
	;
	v439 = *(*int32)(unsafe.Add(mBase, uint32(v111)+80))
	v442 = v400 ^ int32(-1) + v224
	v443 = v439 + v442
	*(*int32)(unsafe.Add(mBase, uint32(v111)+36)) = v443
	v445 = *(*int32)(unsafe.Add(mBase, uint32(v111)+44))
	if int32(0) < v442 {
		goto L113
	} else {
		goto L114
	}
L111:
	;
	goto L112
L112:
	;
	if base.Ui32(v437+int32(1)) < base.Ui32(v434) {
		goto L56
	} else {
		goto L145
	}
L113:
	;
	v448 = v445
	v455 = v439
	goto L116
L114:
	;
	v543 = v445
	goto L115
L115:
	;
	v560 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v543<<(uint(int32(1))%32))+uint32(_c_F_replication_yylex[3]))))
	if v560 != 0 {
		goto L134
	} else {
		goto L135
	}
L116:
	;
	v462 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v455))))
	if v462 != 0 {
		goto L118
	} else {
		goto L119
	}
L117:
	;
	v543 = v539
	goto L115
L118:
	;
	v463 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v462)+uint32(_c_F_replication_yylex[2]))))
	v464 = v463
	goto L120
L119:
	;
	v464 = int32(1)
	goto L120
L120:
	;
	v466 = v448 << (uint(int32(1)) % 32)
	v469 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v466)+uint32(_c_F_replication_yylex[3]))))
	if v469 != 0 {
		goto L121
	} else {
		goto L122
	}
L121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v111)+68)) = v455
	*(*int32)(unsafe.Add(mBase, uint32(v111)+64)) = v448
	goto L123
L122:
	;
	goto L123
L123:
	;
	v473 = v464 & int32(255)
	v476 = int32(*(*int16)(unsafe.Add(mBase, uint32(v466)+uint32(_c_F_replication_yylex[4]))))
	v477 = v473 + v476
	v482 = int32(*(*int16)(unsafe.Add(mBase, uint32(v477<<(uint(int32(1))%32))+uint32(_c_F_replication_yylex[5]))))
	if v482 != v448 {
		goto L124
	} else {
		goto L125
	}
L124:
	;
	v484 = v448
	v486 = v464
	v489 = v473
	goto L127
L125:
	;
	v528 = v477
	goto L126
L126:
	;
	v535 = int32(1)
	v539 = int32(*(*int16)(unsafe.Add(mBase, uint32(v528<<(uint(v535)%32))+uint32(_c_F_replication_yylex[6]))))
	v541 = v455 + v535
	if v541 != v443 {
		v448 = v539
		v455 = v541
		goto L116
	} else {
		goto L133
	}
L127:
	;
	v501 = int32(*(*int16)(unsafe.Add(mBase, uint32(v484<<(uint(int32(1))%32))+uint32(_c_F_replication_yylex[7]))))
	if int32(285) <= v501 {
		goto L129
	} else {
		goto L130
	}
L128:
	;
	v528 = v513
	goto L126
L129:
	;
	v504 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v489)+uint32(_c_F_replication_yylex[8]))))
	v505 = v504
	goto L131
L130:
	;
	v505 = v486
	goto L131
L131:
	;
	v507 = v505 & int32(255)
	v508 = int32(1)
	v512 = int32(*(*int16)(unsafe.Add(mBase, uint32(v501<<(uint(v508)%32))+uint32(_c_F_replication_yylex[4]))))
	v513 = v507 + v512
	v518 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v513<<(uint(v508)%32))+uint32(_c_F_replication_yylex[5]))))
	if v518 != v501&int32(_a_F_replication_yylex_0) {
		v484 = v501
		v486 = v505
		v489 = v507
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
	*(*int32)(unsafe.Add(mBase, uint32(v111)+68)) = v443
	*(*int32)(unsafe.Add(mBase, uint32(v111)+64)) = v543
	goto L136
L135:
	;
	goto L136
L136:
	;
	v563 = int32(1)
	v567 = int32(*(*int16)(unsafe.Add(mBase, uint32(v543<<(uint(v563)%32))+uint32(_c_F_replication_yylex[4]))))
	v569 = v567 + v563
	v574 = int32(*(*int16)(unsafe.Add(mBase, uint32(v569<<(uint(v563)%32))+uint32(_c_F_replication_yylex[5]))))
	if v574 != v543 {
		goto L137
	} else {
		goto L138
	}
L137:
	;
	v576 = v543
	goto L140
L138:
	;
	v610 = v569
	goto L139
L139:
	;
	if v610 == int32(0) {
		v210 = v439
		goto L42
	} else {
		goto L143
	}
L140:
	;
	v589 = int32(1)
	v593 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v576<<(uint(v589)%32))+uint32(_c_F_replication_yylex[7]))))
	v594 = base.I32_extend16_s(v593)
	v599 = int32(*(*int16)(unsafe.Add(mBase, uint32(v594<<(uint(v589)%32))+uint32(_c_F_replication_yylex[4]))))
	v601 = v599 + v589
	v606 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v601<<(uint(v589)%32))+uint32(_c_F_replication_yylex[5]))))
	if v593 != v606 {
		v576 = v594
		goto L140
	} else {
		goto L142
	}
L141:
	;
	v610 = v601
	goto L139
L142:
	;
	goto L141
L143:
	;
	v627 = int32(*(*int16)(unsafe.Add(mBase, uint32(v610<<(uint(int32(1))%32))+uint32(_c_F_replication_yylex[6]))))
	if v627 == int32(284) {
		v210 = v439
		goto L42
	} else {
		goto L144
	}
L144:
	;
	v631 = v443 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v111)+36)) = v631
	v110 = v627
	v113 = v439
	v117 = v631
	goto L27
L145:
	;
	v636 = *(*int32)(unsafe.Add(mBase, uint32(v111)+80))
	v637 = *(*int32)(unsafe.Add(mBase, uint32(v431)+40))
	if v637 == int32(0) {
		goto L146
	} else {
		goto L147
	}
L146:
	;
	if v434-v636 != int32(1) {
		v1328 = v435
		v1331 = v636
		v1332 = v436
		goto L52
	} else {
		goto L149
	}
L147:
	;
	goto L148
L148:
	;
	v645 = v636 ^ int32(-1) + v434
	if int32(0) < v645 {
		goto L150
	} else {
		goto L151
	}
L149:
	;
	v1453 = v636
	goto L48
L150:
	;
	v648 = int32(7)
	v649 = v645 & v648
	if base.Ui32(v434-v636-int32(2)) < base.Ui32(v648) {
		goto L155
	} else {
		goto L156
	}
L151:
	;
	v753 = v431
	v756 = v432
	v757 = v433
	goto L152
L152:
	;
	v764 = *(*int32)(unsafe.Add(mBase, uint32(v753)+44))
	if v764 == int32(2) {
		goto L165
	} else {
		goto L166
	}
L153:
	;
	v745 = *(*int32)(unsafe.Add(mBase, uint32(v111)+20))
	v746 = *(*int32)(unsafe.Add(mBase, uint32(v111)+12))
	v750 = *(*int32)(unsafe.Add(mBase, uint32(v745+v746<<(uint(int32(2))%32))))
	v753 = v750
	v756 = v745
	v757 = v746
	goto L152
L154:
	;
	v710 = v696
	v712 = v698
	v715 = int32(0)
	goto L162
L155:
	;
	v696 = v435
	v698 = v636
	goto L154
L156:
	;
	goto L157
L157:
	;
	v658 = v435
	v660 = v636
	v663 = int32(0)
	goto L158
L158:
	;
	v671 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v660))))
	*(*uint8)(unsafe.Add(mBase, uint32(v658))) = uint8(v671)
	v673 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v660)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v658)+1)) = uint8(v673)
	v675 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v660)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v658)+2)) = uint8(v675)
	v677 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v660)+3)))
	*(*uint8)(unsafe.Add(mBase, uint32(v658)+3)) = uint8(v677)
	v679 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v660)+4)))
	*(*uint8)(unsafe.Add(mBase, uint32(v658)+4)) = uint8(v679)
	v681 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v660)+5)))
	*(*uint8)(unsafe.Add(mBase, uint32(v658)+5)) = uint8(v681)
	v683 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v660)+6)))
	*(*uint8)(unsafe.Add(mBase, uint32(v658)+6)) = uint8(v683)
	v685 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v660)+7)))
	*(*uint8)(unsafe.Add(mBase, uint32(v658)+7)) = uint8(v685)
	v687 = int32(8)
	v688 = v658 + v687
	v690 = v660 + v687
	v692 = v663 + v687
	if v692 != v645&int32(2147483640) {
		v658 = v688
		v660 = v690
		v663 = v692
		goto L158
	} else {
		goto L160
	}
L159:
	;
	if v649 == int32(0) {
		goto L153
	} else {
		goto L161
	}
L160:
	;
	goto L159
L161:
	;
	v696 = v688
	v698 = v690
	goto L154
L162:
	;
	v723 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v712))))
	*(*uint8)(unsafe.Add(mBase, uint32(v710))) = uint8(v723)
	v725 = int32(1)
	v730 = v715 + v725
	if v730 != v649 {
		v710 = v710 + v725
		v712 = v712 + v725
		v715 = v730
		goto L162
	} else {
		goto L164
	}
L163:
	;
	goto L153
L164:
	;
	goto L163
L165:
	;
	v767 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v111)+28)) = v767
	v1008 = v767
	v1017 = v756 + v757<<(uint(int32(2))%32)
	goto L53
L166:
	;
	goto L167
L167:
	;
	v773 = int32(0)
	v774 = *(*int32)(unsafe.Add(mBase, uint32(v753)+12))
	v775 = v636 - v434
	v776 = v774 + v775
	if v776 <= v773 {
		goto L168
	} else {
		goto L169
	}
L168:
	;
	v779 = *(*int32)(unsafe.Add(mBase, uint32(v111)+36))
	v782 = v753
	v783 = v779
	v786 = v774
	goto L171
L169:
	;
	v831 = v776
	v833 = v753
	goto L170
L170:
	;
	v844 = int32(_a_F_replication_yylex_10)
	if base.Ui32(v844) <= base.Ui32(v831) {
		goto L187
	} else {
		goto L188
	}
L171:
	;
	v793 = *(*int32)(unsafe.Add(mBase, uint32(v782)+20))
	if v793 == int32(0) {
		goto L173
	} else {
		goto L174
	}
L172:
	;
	v831 = v828
	v833 = v826
	goto L170
L173:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v782)+4)) = int32(0)
	goto L49
L174:
	;
	goto L175
L175:
	;
	v798 = *(*int32)(unsafe.Add(mBase, uint32(v782)+4))
	v800 = v786 << (uint(int32(1)) % 32)
	if v800 <= int32(0) {
		goto L176
	} else {
		goto L177
	}
L176:
	;
	v804 = base.I32_div_s(v786, int32(8))
	v806 = v804 + v786
	goto L178
L177:
	;
	v806 = v800
	goto L178
L178:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v782)+12)) = v806
	v809 = v806 + int32(2)
	if v798 != 0 {
		goto L180
	} else {
		goto L181
	}
L179:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v782)+4)) = v814
	if v814 == int32(0) {
		goto L49
	} else {
		goto L185
	}
L180:
	;
	v810 = F_repalloc(m, v798, v809)
	mBase = m.M
	v811 = m.ExcPending
	if v811 != 0 {
		goto L18
	} else {
		goto L183
	}
L181:
	;
	goto L182
L182:
	;
	v812 = F_palloc(m, v809)
	mBase = m.M
	v813 = m.ExcPending
	if v813 != 0 {
		goto L18
	} else {
		goto L184
	}
L183:
	;
	v814 = v810
	goto L179
L184:
	;
	v814 = v812
	goto L179
L185:
	;
	v819 = v814 + (v783 - v798)
	*(*int32)(unsafe.Add(mBase, uint32(v111)+36)) = v819
	v821 = *(*int32)(unsafe.Add(mBase, uint32(v111)+20))
	v822 = *(*int32)(unsafe.Add(mBase, uint32(v111)+12))
	v826 = *(*int32)(unsafe.Add(mBase, uint32(v821+v822<<(uint(int32(2))%32))))
	v827 = *(*int32)(unsafe.Add(mBase, uint32(v826)+12))
	v828 = v827 + v775
	if v828 <= int32(0) {
		v782 = v826
		v783 = v819
		v786 = v827
		goto L171
	} else {
		goto L186
	}
L186:
	;
	goto L172
L187:
	;
	v847 = v844
	goto L189
L188:
	;
	v847 = v831
	goto L189
L189:
	;
	v848 = *(*int32)(unsafe.Add(mBase, uint32(v833)+24))
	if v848 != 0 {
		goto L190
	} else {
		goto L191
	}
L190:
	;
	v853 = v773
	goto L194
L191:
	;
	goto L192
L192:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_replication_yylex[9])) = int32(0)
	v910 = *(*int32)(unsafe.Add(mBase, uint32(v111)+20))
	v911 = *(*int32)(unsafe.Add(mBase, uint32(v111)+12))
	v915 = *(*int32)(unsafe.Add(mBase, uint32(v910+v911<<(uint(int32(2))%32))))
	v916 = *(*int32)(unsafe.Add(mBase, uint32(v915)+4))
	v919 = *(*int32)(unsafe.Add(mBase, uint32(v111)+4))
	v920 = F_fread(m, v916+v645, int32(1), v847, v919)
	mBase = m.M
	v921 = m.ExcPending
	if v921 != 0 {
		goto L18
	} else {
		goto L205
	}
L193:
	;
	switch v866 {
	case 0:
		goto L201
	default:
		v905 = v880
		goto L199
	case 11:
		goto L200
	}
L194:
	;
	v862 = *(*int32)(unsafe.Add(mBase, uint32(v111)+4))
	v863 = F_do_getc(m, v862)
	mBase = m.M
	v864 = m.ExcPending
	if v864 != 0 {
		goto L18
	} else {
		goto L197
	}
L195:
	;
	v880 = v847
	goto L193
L196:
	;
	v867 = *(*int32)(unsafe.Add(mBase, uint32(v111)+20))
	v868 = *(*int32)(unsafe.Add(mBase, uint32(v111)+12))
	v872 = *(*int32)(unsafe.Add(mBase, uint32(v867+v868<<(uint(int32(2))%32))))
	v873 = *(*int32)(unsafe.Add(mBase, uint32(v872)+4))
	*(*uint8)(unsafe.Add(mBase, uint32(v873+v645+v853))) = uint8(v863)
	v878 = v853 + int32(1)
	if v878 != v847 {
		v853 = v878
		goto L194
	} else {
		goto L198
	}
L197:
	;
	v866 = v863 + int32(1)
	switch v866 {
	case 0, 11:
		v880 = v853
		goto L193
	default:
		goto L196
	}
L198:
	;
	goto L195
L199:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v111)+28)) = v905
	v990 = v905
	goto L54
L200:
	;
	v892 = *(*int32)(unsafe.Add(mBase, uint32(v111)+20))
	v893 = *(*int32)(unsafe.Add(mBase, uint32(v111)+12))
	v897 = *(*int32)(unsafe.Add(mBase, uint32(v892+v893<<(uint(int32(2))%32))))
	v898 = *(*int32)(unsafe.Add(mBase, uint32(v897)+4))
	v901 = int32(10)
	*(*uint8)(unsafe.Add(mBase, uint32(v898+v645+v880))) = uint8(v901)
	v905 = v880 + int32(1)
	goto L199
L201:
	;
	v881 = *(*int32)(unsafe.Add(mBase, uint32(v111)+4))
	v882 = *(*int32)(unsafe.Add(mBase, uint32(v881)))
	goto L202
L202:
	;
	if int32(base.Ui32(v882)>>(uint(int32(5))%32))&int32(1) == int32(0) {
		v905 = v880
		goto L199
	} else {
		goto L203
	}
L203:
	;
	F_yy_fatal_error_3(m, int32(_a_F_replication_yylex_3))
	mBase = m.M
	v891 = m.ExcPending
	if v891 != 0 {
		goto L18
	} else {
		goto L204
	}
L204:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L205:
	;
	v926 = v920
	goto L206
L206:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v111)+28)) = v926
	if v926 != 0 {
		v990 = v926
		goto L54
	} else {
		goto L208
	}
L208:
	;
	v936 = *(*int32)(unsafe.Add(mBase, uint32(v111)+4))
	v937 = *(*int32)(unsafe.Add(mBase, uint32(v936)))
	goto L209
L209:
	;
	if int32(base.Ui32(v937)>>(uint(int32(5))%32))&int32(1) == int32(0) {
		goto L210
	} else {
		goto L211
	}
L210:
	;
	v990 = int32(0)
	goto L54
L211:
	;
	goto L212
L212:
	;
	v946 = *(*int32)(unsafe.Add(mBase, _c_F_replication_yylex[9]))
	if v946 != int32(27) {
		goto L55
	} else {
		goto L213
	}
L213:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_replication_yylex[9])) = int32(0)
	v952 = *(*int32)(unsafe.Add(mBase, uint32(v111)+4))
	v953 = *(*int32)(unsafe.Add(mBase, uint32(v952)))
	*(*int32)(unsafe.Add(mBase, uint32(v952))) = v953 & int32(-49)
	goto L214
L214:
	;
	v957 = *(*int32)(unsafe.Add(mBase, uint32(v111)+20))
	v958 = *(*int32)(unsafe.Add(mBase, uint32(v111)+12))
	v962 = *(*int32)(unsafe.Add(mBase, uint32(v957+v958<<(uint(int32(2))%32))))
	v963 = *(*int32)(unsafe.Add(mBase, uint32(v962)+4))
	v966 = *(*int32)(unsafe.Add(mBase, uint32(v111)+4))
	v967 = F_fread(m, v963+v645, int32(1), v847, v966)
	mBase = m.M
	v968 = m.ExcPending
	if v968 != 0 {
		goto L18
	} else {
		goto L215
	}
L215:
	;
	v926 = v967
	goto L206
L216:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L217:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L218:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L219:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L220:
	;
	v1144 = *(*int32)(unsafe.Add(mBase, uint32(v111)+28))
	v1145 = v1144 + v645
	v1146 = *(*int32)(unsafe.Add(mBase, uint32(v111)+20))
	v1147 = *(*int32)(unsafe.Add(mBase, uint32(v111)+12))
	v1151 = *(*int32)(unsafe.Add(mBase, uint32(v1146+v1147<<(uint(int32(2))%32))))
	v1152 = *(*int32)(unsafe.Add(mBase, uint32(v1151)+12))
	if v1152 < v1145 {
		goto L244
	} else {
		goto L245
	}
L221:
	;
	if v645 == int32(0) {
		goto L222
	} else {
		goto L223
	}
L222:
	;
	v1024 = *(*int32)(unsafe.Add(mBase, uint32(v111)+4))
	v1025 = *(*int32)(unsafe.Add(mBase, uint32(v111)+20))
	if v1025 != 0 {
		goto L227
	} else {
		goto L228
	}
L223:
	;
	goto L224
L224:
	;
	v1129 = *(*int32)(unsafe.Add(mBase, uint32(v111)+20))
	v1130 = *(*int32)(unsafe.Add(mBase, uint32(v111)+12))
	v1131 = int32(2)
	v1134 = *(*int32)(unsafe.Add(mBase, uint32(v1129+v1130<<(uint(v1131)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v1134)+44)) = v1131
	v1143 = v1131
	goto L220
L225:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1091)+40)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1091))) = v1024
	v1098 = *(*int32)(unsafe.Add(mBase, uint32(v111)+20))
	if v1098 != 0 {
		goto L240
	} else {
		goto L241
	}
L226:
	;
	v1048 = *(*int32)(unsafe.Add(mBase, _c_F_replication_yylex[9]))
	v1049 = *(*int32)(unsafe.Add(mBase, uint32(v111)+12))
	v1053 = *(*int32)(unsafe.Add(mBase, uint32(v1046+v1049<<(uint(int32(2))%32))))
	if v1053 == int32(0) {
		goto L234
	} else {
		goto L235
	}
L227:
	;
	v1026 = *(*int32)(unsafe.Add(mBase, uint32(v111)+12))
	v1030 = *(*int32)(unsafe.Add(mBase, uint32(v1025+v1026<<(uint(int32(2))%32))))
	if v1030 != 0 {
		v1046 = v1025
		goto L226
	} else {
		goto L230
	}
L228:
	;
	goto L229
L229:
	;
	F_replication_yyensure_buffer_stack(m, v111)
	mBase = m.M
	v1032 = m.ExcPending
	if v1032 != 0 {
		goto L18
	} else {
		goto L231
	}
L230:
	;
	goto L229
L231:
	;
	v1033 = *(*int32)(unsafe.Add(mBase, uint32(v111)+4))
	v1034 = F_replication_yy_create_buffer(m, v1033, v111)
	mBase = m.M
	v1035 = m.ExcPending
	if v1035 != 0 {
		goto L18
	} else {
		goto L232
	}
L232:
	;
	v1036 = *(*int32)(unsafe.Add(mBase, uint32(v111)+20))
	v1037 = *(*int32)(unsafe.Add(mBase, uint32(v111)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v1036+v1037<<(uint(int32(2))%32)))) = v1034
	v1042 = *(*int32)(unsafe.Add(mBase, uint32(v111)+20))
	if v1042 != 0 {
		v1046 = v1042
		goto L226
	} else {
		goto L233
	}
L233:
	;
	v1044 = *(*int32)(unsafe.Add(mBase, _c_F_replication_yylex[9]))
	v1091 = int32(0)
	v1094 = v1044
	goto L225
L234:
	;
	v1091 = int32(0)
	v1094 = v1048
	goto L225
L235:
	;
	goto L236
L236:
	;
	v1057 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1053)+16)) = v1057
	v1059 = *(*int32)(unsafe.Add(mBase, uint32(v1053)+4))
	*(*uint8)(unsafe.Add(mBase, uint32(v1059))) = uint8(v1057)
	v1062 = *(*int32)(unsafe.Add(mBase, uint32(v1053)+4))
	*(*uint8)(unsafe.Add(mBase, uint32(v1062)+1)) = uint8(v1057)
	*(*int32)(unsafe.Add(mBase, uint32(v1053)+44)) = v1057
	*(*int32)(unsafe.Add(mBase, uint32(v1053)+28)) = int32(1)
	v1069 = *(*int32)(unsafe.Add(mBase, uint32(v1053)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1053)+8)) = v1069
	v1071 = *(*int32)(unsafe.Add(mBase, uint32(v111)+20))
	if v1071 == v1057 {
		v1091 = v1053
		v1094 = v1048
		goto L225
	} else {
		goto L237
	}
L237:
	;
	v1074 = *(*int32)(unsafe.Add(mBase, uint32(v111)+12))
	v1077 = v1071 + v1074<<(uint(int32(2))%32)
	v1078 = *(*int32)(unsafe.Add(mBase, uint32(v1077)))
	if v1053 != v1078 {
		v1091 = v1053
		v1094 = v1048
		goto L225
	} else {
		goto L238
	}
L238:
	;
	v1080 = *(*int32)(unsafe.Add(mBase, uint32(v1078)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v111)+28)) = v1080
	v1082 = *(*int32)(unsafe.Add(mBase, uint32(v1077)))
	v1083 = *(*int32)(unsafe.Add(mBase, uint32(v1082)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v111)+80)) = v1083
	*(*int32)(unsafe.Add(mBase, uint32(v111)+36)) = v1083
	v1086 = *(*int32)(unsafe.Add(mBase, uint32(v1077)))
	v1087 = *(*int32)(unsafe.Add(mBase, uint32(v1086)))
	*(*int32)(unsafe.Add(mBase, uint32(v111)+4)) = v1087
	v1089 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1083))))
	*(*uint8)(unsafe.Add(mBase, uint32(v111)+24)) = uint8(v1089)
	v1091 = v1053
	v1094 = v1048
	goto L225
L239:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1091)+24)) = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_replication_yylex[9])) = v1094
	v1111 = *(*int32)(unsafe.Add(mBase, uint32(v111)+20))
	v1112 = *(*int32)(unsafe.Add(mBase, uint32(v111)+12))
	v1115 = v1111 + v1112<<(uint(int32(2))%32)
	v1116 = *(*int32)(unsafe.Add(mBase, uint32(v1115)))
	v1117 = *(*int32)(unsafe.Add(mBase, uint32(v1116)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v111)+28)) = v1117
	v1119 = *(*int32)(unsafe.Add(mBase, uint32(v1115)))
	v1120 = *(*int32)(unsafe.Add(mBase, uint32(v1119)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v111)+36)) = v1120
	*(*int32)(unsafe.Add(mBase, uint32(v111)+80)) = v1120
	v1123 = *(*int32)(unsafe.Add(mBase, uint32(v1115)))
	v1124 = *(*int32)(unsafe.Add(mBase, uint32(v1123)))
	*(*int32)(unsafe.Add(mBase, uint32(v111)+4)) = v1124
	v1126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1120))))
	*(*uint8)(unsafe.Add(mBase, uint32(v111)+24)) = uint8(v1126)
	v1143 = int32(1)
	goto L220
L240:
	;
	v1099 = *(*int32)(unsafe.Add(mBase, uint32(v111)+12))
	v1103 = *(*int32)(unsafe.Add(mBase, uint32(v1098+v1099<<(uint(int32(2))%32))))
	if v1091 == v1103 {
		goto L239
	} else {
		goto L243
	}
L241:
	;
	goto L242
L242:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1091)+32)) = int64(1)
	goto L239
L243:
	;
	goto L242
L244:
	;
	v1156 = v1145 + v1144>>(uint(int32(1))%32)
	v1157 = *(*int32)(unsafe.Add(mBase, uint32(v1151)+4))
	if v1157 != 0 {
		goto L248
	} else {
		goto L249
	}
L245:
	;
	v1186 = v1145
	v1187 = v1146
	v1189 = v1147
	goto L246
L246:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v111)+28)) = v1186
	v1191 = int32(2)
	v1194 = *(*int32)(unsafe.Add(mBase, uint32(v1187+v1189<<(uint(v1191)%32))))
	v1195 = *(*int32)(unsafe.Add(mBase, uint32(v1194)+4))
	v1197 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1195+v1186))) = uint8(v1197)
	v1199 = *(*int32)(unsafe.Add(mBase, uint32(v111)+20))
	v1200 = *(*int32)(unsafe.Add(mBase, uint32(v111)+12))
	v1204 = *(*int32)(unsafe.Add(mBase, uint32(v1199+v1200<<(uint(v1191)%32))))
	v1205 = *(*int32)(unsafe.Add(mBase, uint32(v1204)+4))
	v1206 = *(*int32)(unsafe.Add(mBase, uint32(v111)+28))
	*(*uint8)(unsafe.Add(mBase, uint32(v1205+v1206)+1)) = uint8(v1197)
	v1210 = *(*int32)(unsafe.Add(mBase, uint32(v111)+20))
	v1211 = *(*int32)(unsafe.Add(mBase, uint32(v111)+12))
	v1214 = v1210 + v1211<<(uint(v1191)%32)
	v1215 = *(*int32)(unsafe.Add(mBase, uint32(v1214)))
	v1216 = *(*int32)(unsafe.Add(mBase, uint32(v1215)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v111)+80)) = v1216
	if v1143 == int32(1) {
		v1453 = v1216
		goto L48
	} else {
		goto L254
	}
L247:
	;
	v1163 = *(*int32)(unsafe.Add(mBase, uint32(v111)+20))
	v1164 = *(*int32)(unsafe.Add(mBase, uint32(v111)+12))
	v1165 = int32(2)
	v1168 = *(*int32)(unsafe.Add(mBase, uint32(v1163+v1164<<(uint(v1165)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v1168)+4)) = v1162
	v1170 = *(*int32)(unsafe.Add(mBase, uint32(v111)+20))
	v1171 = *(*int32)(unsafe.Add(mBase, uint32(v111)+12))
	v1175 = *(*int32)(unsafe.Add(mBase, uint32(v1170+v1171<<(uint(v1165)%32))))
	v1176 = *(*int32)(unsafe.Add(mBase, uint32(v1175)+4))
	if v1176 == int32(0) {
		goto L50
	} else {
		goto L253
	}
L248:
	;
	v1158 = F_repalloc(m, v1157, v1156)
	mBase = m.M
	v1159 = m.ExcPending
	if v1159 != 0 {
		goto L18
	} else {
		goto L251
	}
L249:
	;
	goto L250
L250:
	;
	v1160 = F_palloc(m, v1156)
	mBase = m.M
	v1161 = m.ExcPending
	if v1161 != 0 {
		goto L18
	} else {
		goto L252
	}
L251:
	;
	v1162 = v1158
	goto L247
L252:
	;
	v1162 = v1160
	goto L247
L253:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1175)+12)) = v1156 - int32(2)
	v1182 = *(*int32)(unsafe.Add(mBase, uint32(v111)+12))
	v1183 = *(*int32)(unsafe.Add(mBase, uint32(v111)+20))
	v1184 = *(*int32)(unsafe.Add(mBase, uint32(v111)+28))
	v1186 = v1184 + v645
	v1187 = v1183
	v1189 = v1182
	goto L246
L254:
	;
	switch v1143 - int32(1) {
	case 0:
		goto L51
	case 1:
		goto L255
	default:
		goto L256
	}
L255:
	;
	v1325 = *(*int32)(unsafe.Add(mBase, uint32(v111)+28))
	v1326 = *(*int32)(unsafe.Add(mBase, uint32(v1214)))
	v1327 = *(*int32)(unsafe.Add(mBase, uint32(v1326)+4))
	v1328 = v1327
	v1331 = v1216
	v1332 = v1325
	goto L52
L256:
	;
	v1224 = v400 ^ int32(-1) + v224
	v1225 = v1216 + v1224
	*(*int32)(unsafe.Add(mBase, uint32(v111)+36)) = v1225
	v1227 = *(*int32)(unsafe.Add(mBase, uint32(v111)+44))
	if v1224 <= int32(0) {
		v110 = v1227
		v113 = v1216
		v117 = v1225
		goto L27
	} else {
		goto L257
	}
L257:
	;
	v1230 = v1227
	v1234 = v1216
	goto L258
L258:
	;
	v1244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1234))))
	if v1244 != 0 {
		goto L260
	} else {
		goto L261
	}
L259:
	;
	v110 = v1321
	v113 = v1216
	v117 = v1225
	goto L27
L260:
	;
	v1245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1244)+uint32(_c_F_replication_yylex[2]))))
	v1246 = v1245
	goto L262
L261:
	;
	v1246 = int32(1)
	goto L262
L262:
	;
	v1248 = v1230 << (uint(int32(1)) % 32)
	v1251 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1248)+uint32(_c_F_replication_yylex[3]))))
	if v1251 != 0 {
		goto L263
	} else {
		goto L264
	}
L263:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v111)+68)) = v1234
	*(*int32)(unsafe.Add(mBase, uint32(v111)+64)) = v1230
	goto L265
L264:
	;
	goto L265
L265:
	;
	v1255 = v1246 & int32(255)
	v1258 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1248)+uint32(_c_F_replication_yylex[4]))))
	v1259 = v1255 + v1258
	v1264 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1259<<(uint(int32(1))%32))+uint32(_c_F_replication_yylex[5]))))
	if v1264 != v1230 {
		goto L266
	} else {
		goto L267
	}
L266:
	;
	v1266 = v1230
	v1268 = v1246
	v1271 = v1255
	goto L269
L267:
	;
	v1310 = v1259
	goto L268
L268:
	;
	v1317 = int32(1)
	v1321 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1310<<(uint(v1317)%32))+uint32(_c_F_replication_yylex[6]))))
	v1323 = v1234 + v1317
	if v1225 != v1323 {
		v1230 = v1321
		v1234 = v1323
		goto L258
	} else {
		goto L275
	}
L269:
	;
	v1283 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1266<<(uint(int32(1))%32))+uint32(_c_F_replication_yylex[7]))))
	if int32(285) <= v1283 {
		goto L271
	} else {
		goto L272
	}
L270:
	;
	v1310 = v1295
	goto L268
L271:
	;
	v1286 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1271)+uint32(_c_F_replication_yylex[8]))))
	v1287 = v1286
	goto L273
L272:
	;
	v1287 = v1268
	goto L273
L273:
	;
	v1289 = v1287 & int32(255)
	v1290 = int32(1)
	v1294 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1283<<(uint(v1290)%32))+uint32(_c_F_replication_yylex[4]))))
	v1295 = v1289 + v1294
	v1300 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1295<<(uint(v1290)%32))+uint32(_c_F_replication_yylex[5]))))
	if v1300 != v1283&int32(_a_F_replication_yylex_0) {
		v1266 = v1283
		v1268 = v1287
		v1271 = v1289
		goto L269
	} else {
		goto L274
	}
L274:
	;
	goto L270
L275:
	;
	goto L259
L276:
	;
	v1345 = v1343
	v1349 = v1331
	goto L277
L277:
	;
	v1359 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1349))))
	if v1359 != 0 {
		goto L279
	} else {
		goto L280
	}
L278:
	;
	v217 = v1438
	v224 = v1341
	v225 = v1331
	goto L44
L279:
	;
	v1360 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1359)+uint32(_c_F_replication_yylex[2]))))
	v1361 = v1360
	goto L281
L280:
	;
	v1361 = int32(1)
	goto L281
L281:
	;
	v1366 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1345<<(uint(int32(1))%32))+uint32(_c_F_replication_yylex[3]))))
	if v1366 != 0 {
		goto L282
	} else {
		goto L283
	}
L282:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v111)+68)) = v1349
	*(*int32)(unsafe.Add(mBase, uint32(v111)+64)) = v1345
	goto L284
L283:
	;
	goto L284
L284:
	;
	v1370 = v1361 & int32(255)
	v1371 = int32(1)
	v1375 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1345<<(uint(v1371)%32))+uint32(_c_F_replication_yylex[4]))))
	v1376 = v1370 + v1375
	v1381 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1376<<(uint(v1371)%32))+uint32(_c_F_replication_yylex[5]))))
	if v1381 != v1345 {
		goto L285
	} else {
		goto L286
	}
L285:
	;
	v1383 = v1345
	v1385 = v1361
	v1388 = v1370
	goto L288
L286:
	;
	v1427 = v1376
	goto L287
L287:
	;
	v1434 = int32(1)
	v1438 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1427<<(uint(v1434)%32))+uint32(_c_F_replication_yylex[6]))))
	v1440 = v1349 + v1434
	if v1440 != v1341 {
		v1345 = v1438
		v1349 = v1440
		goto L277
	} else {
		goto L294
	}
L288:
	;
	v1400 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1383<<(uint(int32(1))%32))+uint32(_c_F_replication_yylex[7]))))
	if int32(285) <= v1400 {
		goto L290
	} else {
		goto L291
	}
L289:
	;
	v1427 = v1412
	goto L287
L290:
	;
	v1403 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1388)+uint32(_c_F_replication_yylex[8]))))
	v1404 = v1403
	goto L292
L291:
	;
	v1404 = v1385
	goto L292
L292:
	;
	v1406 = v1404 & int32(255)
	v1407 = int32(1)
	v1411 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1400<<(uint(v1407)%32))+uint32(_c_F_replication_yylex[4]))))
	v1412 = v1406 + v1411
	v1417 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1412<<(uint(v1407)%32))+uint32(_c_F_replication_yylex[5]))))
	if v1417 != v1400&int32(_a_F_replication_yylex_0) {
		v1383 = v1400
		v1385 = v1404
		v1388 = v1406
		goto L288
	} else {
		goto L293
	}
L293:
	;
	goto L289
L294:
	;
	goto L278
L295:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L296:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
