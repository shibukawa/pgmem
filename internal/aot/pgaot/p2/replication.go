package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"sync/atomic"
	"unsafe"
)

func F_ReplicationOriginExitCleanup(m *base.Module, l0 int32, l1 int64) {
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
	var v20 int32
	_ = v20
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
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
			if v14 == v16 {
				*(*int32)(unsafe.Add(mBase, uint32(v13)+24)) = int32(0)
			} else {
			}
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v13)+28))
			*(*int32)(unsafe.Add(mBase, uint32(v13)+28)) = v20 - int32(1)
			*(*int32)(unsafe.Add(mBase, _c_F_ReplicationOriginExitCleanup[0])) = int32(0)
			v28 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationOriginExitCleanup[1]))
			F_LWLockRelease(m, v28+int32(_a_F_ReplicationOriginExitCleanup_0))
			mBase = m.M
			v32 = m.ExcPending
			if v32 != 0 {
				return
			} else {
				F_ConditionVariableBroadcast(m, v13+int32(32))
				mBase = m.M
				v36 = m.ExcPending
				if v36 != 0 {
					return
				} else {
					return
				}
			}
		}
	} else {
		return
	}
}
func F_ReplicationOriginShmemInit(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	v2 = int32(0)
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationOriginShmemInit[0]))
	if v5 == v2 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v10 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationOriginShmemInit[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_ReplicationOriginShmemInit[2])) = v10 + int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = int32(67)
	v17 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationOriginShmemInit[0]))
	if v17 <= int32(0) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v20 = v2
	goto L4
L4:
	;
	v23 = v20 << (uint(int32(6)) % 32)
	v25 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationOriginShmemInit[2]))
	v30 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationOriginShmemInit[1]))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	F_LWLockInitialize(m, v23+v25+int32(44), v31)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L1
L6:
	;
	return
L7:
	;
	v35 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationOriginShmemInit[2]))
	v38 = v35 + v23 + int32(32)
	v39 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v38))), uint32(v39))
	*(*int64)(unsafe.Add(mBase, uint32(v38)+4)) = int64(-1)
	goto L8
L8:
	;
	v45 = v20 + int32(1)
	v47 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationOriginShmemInit[0]))
	if v45 < v47 {
		v20 = v45
		goto L4
	} else {
		goto L9
	}
L9:
	;
	goto L5
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
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
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
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v221 int32
	_ = v221
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v276 int32
	_ = v276
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v293 int32
	_ = v293
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v302 int32
	_ = v302
	var v306 int32
	_ = v306
	var v309 int32
	_ = v309
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v344 int32
	_ = v344
	var v355 int32
	_ = v355
	var v359 int32
	_ = v359
	var v363 int32
	_ = v363
	var v366 int32
	_ = v366
	var v372 int32
	_ = v372
	var v377 int32
	_ = v377
	v10 = m.G0
	v12 = v10 - int32(96)
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
	v24 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotAcquire[2]))
	v25 = v22 + v24
	if v25 <= int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v355 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotAcquire[0]))
	F_LWLockRelease(m, v355+int32(_a_F_ReplicationSlotAcquire_0))
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L1
	} else {
		goto L112
	}
L4:
	;
	v34 = v25
	goto L8
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L1
	} else {
		goto L103
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L1
	} else {
		goto L98
	}
L7:
	;
	if l1 == int32(0) {
		goto L72
	} else {
		goto L73
	}
L8:
	;
	v39 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotAcquire[3]))
	v45 = int32(0)
	goto L10
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L1
	} else {
		goto L68
	}
L10:
	;
	v51 = v39 + v45*int32(296)
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51)+4)))
	if v52 == int32(1) {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	v90 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotAcquire[4]))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v90)))
	v93 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotAcquire[5]))
	goto L25
L12:
	;
	goto L11
L13:
	;
	v56 = v51 + int32(24)
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56))))
	if base.B2i32(v59 == int32(0))|base.B2i32(v59 != v62) != 0 {
		v80 = v59
		v81 = v62
		goto L17
	} else {
		goto L18
	}
L14:
	;
	goto L15
L15:
	;
	v87 = v45 + int32(1)
	if v87 != v34 {
		v45 = v87
		goto L10
	} else {
		goto L24
	}
L16:
	;
	if v80-v81 == int32(0) {
		goto L12
	} else {
		goto L23
	}
L17:
	;
	goto L16
L18:
	;
	v65 = l0
	v66 = v56
	goto L19
L19:
	;
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+1)))
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65)+1)))
	if v70 == int32(0) {
		v80 = v70
		v81 = v69
		goto L17
	} else {
		goto L21
	}
L20:
	;
	v80 = v70
	v81 = v69
	goto L17
L21:
	;
	v73 = int32(1)
	if v70 == v69 {
		v65 = v65 + v73
		v66 = v66 + v73
		goto L19
	} else {
		goto L22
	}
L22:
	;
	goto L20
L23:
	;
	goto L15
L24:
	;
	goto L3
L25:
	;
	if base.B2i32(v91 == v93) == int32(0) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v97 = int32(_a_F_ReplicationSlotAcquire_1)
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v103 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ReplicationSlotAcquire[6])))
	if base.B2i32(v100 == int32(0))|base.B2i32(v100 != v103) != 0 {
		v121 = v100
		v122 = v103
		goto L30
	} else {
		goto L31
	}
L27:
	;
	goto L28
L28:
	;
	v127 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ReplicationSlotAcquire[7])))
	if v127 == int32(1) {
		goto L39
	} else {
		goto L40
	}
L29:
	;
	if v121-v122 == int32(0) {
		goto L6
	} else {
		goto L36
	}
L30:
	;
	goto L29
L31:
	;
	v106 = l0
	v107 = v97
	goto L32
L32:
	;
	v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+1)))
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106)+1)))
	if v111 == int32(0) {
		v121 = v111
		v122 = v110
		goto L30
	} else {
		goto L34
	}
L33:
	;
	v121 = v111
	v122 = v110
	goto L30
L34:
	;
	v114 = int32(1)
	if v111 == v110 {
		v106 = v106 + v114
		v107 = v107 + v114
		goto L32
	} else {
		goto L35
	}
L35:
	;
	goto L33
L36:
	;
	goto L28
L37:
	;
	v166 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v51))), uint32(v166))
	v170 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotAcquire[8]))
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v170)))
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v171+v165*int32(768))+12))
	v177 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotAcquire[0]))
	F_LWLockRelease(m, v177+int32(_a_F_ReplicationSlotAcquire_0))
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L1
	} else {
		goto L59
	}
L38:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v51)+272)) = int64(0)
	v165 = v162
	goto L37
L39:
	;
	if l1 == int32(0) {
		goto L42
	} else {
		goto L43
	}
L40:
	;
	goto L41
L41:
	;
	v153 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotAcquire[9]))
	*(*int32)(unsafe.Add(mBase, uint32(v51)+8)) = v153
	v157 = base.AtomicRmwXchg32(m, v51, int32(0), int32(1))
	if v157 != 0 {
		goto L54
	} else {
		goto L55
	}
L42:
	;
	F_ConditionVariablePrepareToSleep(m, v51+int32(224))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L1
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	v138 = base.AtomicRmwXchg32(m, v51, int32(0), int32(1))
	if v138 != 0 {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	goto L44
L46:
	;
	F_s_lock(m, v51, int32(_a_F_ReplicationSlotAcquire_2))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L1
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v51)+8))
	if v142 == int32(-1) {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	goto L48
L50:
	;
	v146 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotAcquire[9]))
	*(*int32)(unsafe.Add(mBase, uint32(v51)+8)) = v146
	v148 = v146
	goto L52
L51:
	;
	v148 = v142
	goto L52
L52:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v51)+112))
	if v149 == int32(0) {
		v162 = v148
		goto L38
	} else {
		goto L53
	}
L53:
	;
	v165 = v148
	goto L37
L54:
	;
	F_s_lock(m, v51, int32(_a_F_ReplicationSlotAcquire_2))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L1
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v51)+112))
	if v161 != 0 {
		v165 = v153
		goto L37
	} else {
		goto L58
	}
L57:
	;
	goto L56
L58:
	;
	v162 = v153
	goto L38
L59:
	;
	v183 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotAcquire[9]))
	if v165 == v183 {
		goto L7
	} else {
		goto L60
	}
L60:
	;
	if l1 == int32(0) {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	F_ConditionVariableSleep(m, v51+int32(224), int32(134217778))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L1
	} else {
		goto L64
	}
L62:
	;
	goto L63
L63:
	;
	goto L9
L64:
	;
	F_ConditionVariableCancelSleep(m)
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	v195 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotAcquire[0]))
	v199 = F_LWLockAcquire(m, v195+int32(_a_F_ReplicationSlotAcquire_0), int32(1))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	v202 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotAcquire[1]))
	v204 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotAcquire[2]))
	v205 = v202 + v204
	if int32(0) < v205 {
		v34 = v205
		goto L8
	} else {
		goto L67
	}
L67:
	;
	goto L3
L68:
	;
	F_errcode(m, int32(100663621))
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L1
	} else {
		goto L69
	}
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+52)) = v175
	*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = v56
	F_errmsg(m, int32(_a_F_ReplicationSlotAcquire_3), v12+int32(48))
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L1
	} else {
		goto L70
	}
L70:
	;
	F_errfinish(m, int32(_a_F_ReplicationSlotAcquire_4), int32(720), int32(_a_F_ReplicationSlotAcquire_5))
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L1
	} else {
		goto L71
	}
L71:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L72:
	;
	F_ConditionVariableCancelSleep(m)
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L1
	} else {
		goto L75
	}
L73:
	;
	goto L74
L74:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotAcquire[10])) = v51
	if l2 != 0 {
		goto L76
	} else {
		goto L77
	}
L75:
	;
	goto L74
L76:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v51)+112))
	if v233 != 0 {
		goto L5
	} else {
		goto L79
	}
L77:
	;
	goto L78
L78:
	;
	F_ConditionVariableBroadcast(m, v51+int32(224))
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L1
	} else {
		goto L80
	}
L79:
	;
	goto L78
L80:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v51)+88))
	if v238 != 0 {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	v242 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotAcquire[3]))
	v245 = base.I32_div_s(v51-v242, int32(296))
	goto L84
L82:
	;
	goto L83
L83:
	;
	v252 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ReplicationSlotAcquire[11])))
	if v252 != int32(1) {
		goto L86
	} else {
		goto L87
	}
L84:
	;
	v249 = F_pgstat_get_entry_ref(m, int32(4), int32(0), base.I64_extend_i32_s(v245), int32(1), int32(0))
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	goto L83
L86:
	;
	m.G0 = v12 + int32(96)
	return
L87:
	;
	v258 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ReplicationSlotAcquire[12])))
	if v258 != 0 {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v259 = int32(15)
	goto L90
L89:
	;
	v259 = int32(14)
	goto L90
L90:
	;
	v261 = F_errstart(m, v259, int32(0))
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L1
	} else {
		goto L91
	}
L91:
	;
	if v261 == int32(0) {
		goto L86
	} else {
		goto L92
	}
L92:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v51)+88))
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v56
	if v265 != 0 {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v269 = int32(_a_F_ReplicationSlotAcquire_6)
	goto L95
L94:
	;
	v269 = int32(_a_F_ReplicationSlotAcquire_7)
	goto L95
L95:
	;
	F_errmsg(m, v269, v12)
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L1
	} else {
		goto L96
	}
L96:
	;
	F_errfinish(m, int32(_a_F_ReplicationSlotAcquire_4), int32(760), int32(_a_F_ReplicationSlotAcquire_5))
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L1
	} else {
		goto L97
	}
L97:
	;
	goto L86
L98:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L1
	} else {
		goto L99
	}
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+80)) = l0
	F_errmsg(m, int32(_a_F_ReplicationSlotAcquire_8), v12+int32(80))
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L1
	} else {
		goto L100
	}
L100:
	;
	v296 = F_errdetail(m, int32(_a_F_ReplicationSlotAcquire_9), int32(0))
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L1
	} else {
		goto L101
	}
L101:
	;
	F_errfinish(m, int32(_a_F_ReplicationSlotAcquire_4), int32(665), int32(_a_F_ReplicationSlotAcquire_5))
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L1
	} else {
		goto L102
	}
L102:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L103:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L1
	} else {
		goto L104
	}
L104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v56
	F_errmsg(m, int32(_a_F_ReplicationSlotAcquire_10), v12+int32(32))
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L1
	} else {
		goto L105
	}
L105:
	;
	v316 = *(*int32)(unsafe.Add(mBase, uint32(v51)+112))
	if base.B2i32(int32(base.Ui32(int32(279))>>(uint(v316)%32))&int32(1) == int32(0))|base.B2i32(base.Ui32(int32(8)) < base.Ui32(v316)) != 0 {
		goto L107
	} else {
		goto L108
	}
L106:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v333
	v338 = F_errdetail(m, int32(_a_F_ReplicationSlotAcquire_11), v12+int32(16))
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L1
	} else {
		goto L110
	}
L107:
	;
	v333 = int32(_a_F_ReplicationSlotAcquire_12)
	goto L109
L108:
	;
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v316<<(uint(int32(2))%32))+uint32(_c_F_ReplicationSlotAcquire[13])))
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v331)+4))
	v333 = v332
	goto L109
L109:
	;
	goto L106
L110:
	;
	F_errfinish(m, int32(_a_F_ReplicationSlotAcquire_4), int32(739), int32(_a_F_ReplicationSlotAcquire_5))
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L1
	} else {
		goto L111
	}
L111:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L112:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v363 = m.ExcPending
	if v363 != 0 {
		goto L1
	} else {
		goto L113
	}
L113:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L1
	} else {
		goto L114
	}
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+64)) = l0
	F_errmsg(m, int32(_a_F_ReplicationSlotAcquire_13), v12-int32(-64))
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L1
	} else {
		goto L115
	}
L115:
	;
	F_errfinish(m, int32(_a_F_ReplicationSlotAcquire_4), int32(653), int32(_a_F_ReplicationSlotAcquire_5))
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L1
	} else {
		goto L116
	}
L116:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_ReplicationSlotCleanup(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
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
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v145 int32
	_ = v145
	v2 = int32(0)
	v10 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotCleanup[0]))
	v14 = F_LWLockAcquire(m, v10+int32(_a_F_ReplicationSlotCleanup_0), int32(1))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v17 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotCleanup[1]))
	v19 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotCleanup[2]))
	if v17+v19 <= int32(0) {
		v126 = v2
		v128 = v2
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v130 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotCleanup[0]))
	F_LWLockRelease(m, v130+int32(_a_F_ReplicationSlotCleanup_0))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L1
	} else {
		goto L31
	}
L4:
	;
	v25 = v17
	v26 = v19
	v28 = v2
	v30 = v2
	goto L5
L5:
	;
	v32 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotCleanup[3]))
	v36 = v25
	v37 = v26
	v38 = int32(0)
	v39 = v28
	v40 = v32
	goto L7
L6:
	;
	v126 = v87
	v128 = v113
	goto L3
L7:
	;
	v44 = v40 + v38*int32(296)
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+4)))
	if v45 == int32(1) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v87 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v44))), uint32(v87))
	v92 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotCleanup[0]))
	F_LWLockRelease(m, v92+int32(_a_F_ReplicationSlotCleanup_0))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L1
	} else {
		goto L26
	}
L9:
	;
	goto L8
L10:
	;
	v50 = base.AtomicRmwXchg32(m, v44, int32(0), int32(1))
	if v50 != 0 {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	v79 = v36
	v80 = v37
	v81 = v39
	v82 = v40
	goto L12
L12:
	;
	v84 = v38 + int32(1)
	if v84 < v79+v80 {
		v36 = v79
		v37 = v80
		v38 = v84
		v39 = v81
		v40 = v82
		goto L7
	} else {
		goto L25
	}
L13:
	;
	F_s_lock(m, v44, int32(_a_F_ReplicationSlotCleanup_1))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L1
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v44)+88))
	if v54 != 0 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	goto L15
L17:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v44)+112))
	v57 = v55
	goto L19
L18:
	;
	v57 = int32(1)
	goto L19
L19:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v44)+8))
	v62 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotCleanup[4]))
	if v60 == v62 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	if l0 == int32(0) {
		goto L9
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v70 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v44))), uint32(v70))
	v74 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotCleanup[2]))
	v76 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotCleanup[3]))
	v78 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotCleanup[1]))
	v79 = v78
	v80 = v74
	v81 = base.B2i32(v57 == int32(0)) | v39&int32(1)
	v82 = v76
	goto L12
L23:
	;
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+201)))
	if v66 != 0 {
		goto L9
	} else {
		goto L24
	}
L24:
	;
	goto L22
L25:
	;
	v126 = v81
	v128 = v30
	goto L3
L26:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v44)+88))
	F_ReplicationSlotDropPtr(m, v44)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	F_ConditionVariableBroadcast(m, v44+int32(224))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	v105 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotCleanup[0]))
	v109 = F_LWLockAcquire(m, v105+int32(_a_F_ReplicationSlotCleanup_0), int32(1))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	v111 = int32(0)
	v113 = base.B2i32(v97 != v111) | v30
	v115 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotCleanup[1]))
	v117 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotCleanup[2]))
	if v111 < v115+v117 {
		v25 = v115
		v26 = v117
		v28 = v87
		v30 = v113
		goto L5
	} else {
		goto L30
	}
L30:
	;
	goto L6
L31:
	;
	v135 = int32(1)
	v137 = int32(0)
	if base.B2i32(v128&v135 == v137)|v126&v135 == v137 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	F_RequestDisableLogicalDecoding(m)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L1
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	return
L35:
	;
	goto L34
}
func F_ReplicationSlotPersist(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	v3 = m.G0
	v5 = v3 - int32(1040)
	m.G0 = v5
	v8 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotPersist[0]))
	v11 = base.AtomicRmwXchg32(m, v8, int32(0), int32(1))
	if v11 != 0 {
		F_s_lock(m, v8, int32(_a_F_ReplicationSlotPersist_0))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return
		} else {
			v15 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v8)+92)) = v15
			atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v8))), uint32(v15))
			v21 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotPersist[0]))
			v24 = base.AtomicRmwXchg32(m, v21, v15, int32(1))
			if v24 != 0 {
				F_s_lock(m, v21, int32(_a_F_ReplicationSlotPersist_0))
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return
				} else {
					v28 = int32(_a_F_ReplicationSlotPersist_1)
					v29 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotPersist[0]))
					v30 = int32(257)
					*(*uint16)(unsafe.Add(mBase, uint32(v29)+12)) = uint16(v30)
					v32 = int32(0)
					atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v21))), uint32(v32))
					*(*int32)(unsafe.Add(mBase, uint32(v5))) = int32(_a_F_ReplicationSlotPersist_2)
					v38 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotPersist[0]))
					*(*int32)(unsafe.Add(mBase, uint32(v5)+4)) = v38 + int32(24)
					v43 = v5 + int32(16)
					v45 = F_pg_sprintf(m, v43, int32(_a_F_ReplicationSlotPersist_3), v5)
					mBase = m.M
					v46 = m.ExcPending
					if v46 != 0 {
						return
					} else {
						v48 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotPersist[0]))
						F_SaveSlotToPath(m, v48, v43, int32(21))
						mBase = m.M
						v51 = m.ExcPending
						if v51 != 0 {
							return
						} else {
							m.G0 = v5 + int32(1040)
							return
						}
					}
				}
			} else {
				v28 = int32(_a_F_ReplicationSlotPersist_1)
				v29 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotPersist[0]))
				v30 = int32(257)
				*(*uint16)(unsafe.Add(mBase, uint32(v29)+12)) = uint16(v30)
				v32 = int32(0)
				atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v21))), uint32(v32))
				*(*int32)(unsafe.Add(mBase, uint32(v5))) = int32(_a_F_ReplicationSlotPersist_2)
				v38 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotPersist[0]))
				*(*int32)(unsafe.Add(mBase, uint32(v5)+4)) = v38 + int32(24)
				v43 = v5 + int32(16)
				v45 = F_pg_sprintf(m, v43, int32(_a_F_ReplicationSlotPersist_3), v5)
				mBase = m.M
				v46 = m.ExcPending
				if v46 != 0 {
					return
				} else {
					v48 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotPersist[0]))
					F_SaveSlotToPath(m, v48, v43, int32(21))
					mBase = m.M
					v51 = m.ExcPending
					if v51 != 0 {
						return
					} else {
						m.G0 = v5 + int32(1040)
						return
					}
				}
			}
		}
	} else {
		v15 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v8)+92)) = v15
		atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v8))), uint32(v15))
		v21 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotPersist[0]))
		v24 = base.AtomicRmwXchg32(m, v21, v15, int32(1))
		if v24 != 0 {
			F_s_lock(m, v21, int32(_a_F_ReplicationSlotPersist_0))
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return
			} else {
				v28 = int32(_a_F_ReplicationSlotPersist_1)
				v29 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotPersist[0]))
				v30 = int32(257)
				*(*uint16)(unsafe.Add(mBase, uint32(v29)+12)) = uint16(v30)
				v32 = int32(0)
				atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v21))), uint32(v32))
				*(*int32)(unsafe.Add(mBase, uint32(v5))) = int32(_a_F_ReplicationSlotPersist_2)
				v38 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotPersist[0]))
				*(*int32)(unsafe.Add(mBase, uint32(v5)+4)) = v38 + int32(24)
				v43 = v5 + int32(16)
				v45 = F_pg_sprintf(m, v43, int32(_a_F_ReplicationSlotPersist_3), v5)
				mBase = m.M
				v46 = m.ExcPending
				if v46 != 0 {
					return
				} else {
					v48 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotPersist[0]))
					F_SaveSlotToPath(m, v48, v43, int32(21))
					mBase = m.M
					v51 = m.ExcPending
					if v51 != 0 {
						return
					} else {
						m.G0 = v5 + int32(1040)
						return
					}
				}
			}
		} else {
			v28 = int32(_a_F_ReplicationSlotPersist_1)
			v29 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotPersist[0]))
			v30 = int32(257)
			*(*uint16)(unsafe.Add(mBase, uint32(v29)+12)) = uint16(v30)
			v32 = int32(0)
			atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v21))), uint32(v32))
			*(*int32)(unsafe.Add(mBase, uint32(v5))) = int32(_a_F_ReplicationSlotPersist_2)
			v38 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotPersist[0]))
			*(*int32)(unsafe.Add(mBase, uint32(v5)+4)) = v38 + int32(24)
			v43 = v5 + int32(16)
			v45 = F_pg_sprintf(m, v43, int32(_a_F_ReplicationSlotPersist_3), v5)
			mBase = m.M
			v46 = m.ExcPending
			if v46 != 0 {
				return
			} else {
				v48 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotPersist[0]))
				F_SaveSlotToPath(m, v48, v43, int32(21))
				mBase = m.M
				v51 = m.ExcPending
				if v51 != 0 {
					return
				} else {
					m.G0 = v5 + int32(1040)
					return
				}
			}
		}
	}
}
func F_ReplicationSlotValidateNameInternal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
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
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v141 int32
	_ = v141
	var v148 int32
	_ = v148
	v9 = m.G0
	v11 = v9 - int32(80)
	m.G0 = v11
	v13 = F_strlen(m, l0)
	mBase = m.M
	if v13 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v11 + int32(80)
	return v148
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v141
	v148 = int32(0)
	goto L1
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(33579140)
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = l0
	v20 = F_psprintf(m, int32(_a_F_ReplicationSlotValidateNameInternal_0), v11)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	if base.Ui32(v13) <= base.Ui32(int32(63)) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	return int32(0)
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v20
	v141 = int32(0)
	goto L2
L8:
	;
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v28 != 0 {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L10
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(34103428)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = l0
	v129 = F_psprintf(m, int32(_a_F_ReplicationSlotValidateNameInternal_1), v11+int32(16))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L6
	} else {
		goto L33
	}
L11:
	;
	v33 = l0
	v35 = v28
	goto L14
L12:
	;
	goto L13
L13:
	;
	v79 = int32(1)
	if l1 != 0 {
		v148 = v79
		goto L1
	} else {
		goto L22
	}
L14:
	;
	v41 = int32(255)
	if base.B2i32(v35 == int32(95))|base.B2i32(base.Ui32((v35-int32(97))&v41) < base.Ui32(int32(26)))|base.B2i32(base.Ui32((v35-int32(48))&v41) < base.Ui32(int32(10))) == int32(0) {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	goto L13
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(33579140)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+64)) = l0
	v61 = F_psprintf(m, int32(_a_F_ReplicationSlotValidateNameInternal_2), v11-int32(-64))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L6
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+1)))
	if v70 != 0 {
		v33 = v33 + int32(1)
		v35 = v70
		goto L14
	} else {
		goto L21
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v61
	v66 = F_psprintf(m, int32(_a_F_ReplicationSlotValidateNameInternal_3), int32(0))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L6
	} else {
		goto L20
	}
L20:
	;
	v141 = v66
	goto L2
L21:
	;
	goto L15
L22:
	;
	v80 = int32(_a_F_ReplicationSlotValidateNameInternal_4)
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v86 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ReplicationSlotValidateNameInternal[0])))
	if base.B2i32(v83 == int32(0))|base.B2i32(v83 != v86) != 0 {
		v104 = v83
		v105 = v86
		goto L24
	} else {
		goto L25
	}
L23:
	;
	if v104-v105 != 0 {
		v148 = v79
		goto L1
	} else {
		goto L30
	}
L24:
	;
	goto L23
L25:
	;
	v89 = l0
	v90 = v80
	goto L26
L26:
	;
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90)+1)))
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89)+1)))
	if v94 == int32(0) {
		v104 = v94
		v105 = v93
		goto L24
	} else {
		goto L28
	}
L27:
	;
	v104 = v94
	v105 = v93
	goto L24
L28:
	;
	v97 = int32(1)
	if v94 == v93 {
		v89 = v89 + v97
		v90 = v90 + v97
		goto L26
	} else {
		goto L29
	}
L29:
	;
	goto L27
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(151818372)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = l0
	v113 = F_psprintf(m, int32(_a_F_ReplicationSlotValidateNameInternal_5), v11+int32(48))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L6
	} else {
		goto L31
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v113
	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = int32(_a_F_ReplicationSlotValidateNameInternal_4)
	v121 = F_psprintf(m, int32(_a_F_ReplicationSlotValidateNameInternal_6), v11+int32(32))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L6
	} else {
		goto L32
	}
L32:
	;
	v141 = v121
	goto L2
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v129
	v141 = int32(0)
	goto L2
}
func F_ReplicationSlotsShmemRequest(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v26 int32
	_ = v26
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v8 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotsShmemRequest[0]))
	v10 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotsShmemRequest[1]))
	v11 = v8 + v10
	if v11 != 0 {
		v14 = F_mul_size(m, v11, int32(296))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return
		} else {
			v16 = F_add_size(m, int32(0), v14)
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v5)+12)) = int32(_a_F_ReplicationSlotsShmemRequest_0)
				*(*int32)(unsafe.Add(mBase, uint32(v5)+8)) = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v5)+4)) = v16
				*(*int32)(unsafe.Add(mBase, uint32(v5))) = int32(_a_F_ReplicationSlotsShmemRequest_1)
				F_ShmemRequestStructWithOpts(m, v5)
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return
				} else {
					m.G0 = v5 + int32(16)
					return
				}
			}
		}
	} else {
		m.G0 = v5 + int32(16)
		return
	}
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
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v401 int32
	_ = v401
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v418 int32
	_ = v418
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v428 int32
	_ = v428
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
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
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v445 int32
	_ = v445
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v451 int32
	_ = v451
	var v454 int32
	_ = v454
	var v461 int32
	_ = v461
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v472 int32
	_ = v472
	var v475 int32
	_ = v475
	var v479 int32
	_ = v479
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
	var v495 int32
	_ = v495
	var v507 int32
	_ = v507
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v524 int32
	_ = v524
	var v534 int32
	_ = v534
	var v541 int32
	_ = v541
	var v545 int32
	_ = v545
	var v547 int32
	_ = v547
	var v549 int32
	_ = v549
	var v566 int32
	_ = v566
	var v569 int32
	_ = v569
	var v573 int32
	_ = v573
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
	var v633 int32
	_ = v633
	var v637 int32
	_ = v637
	var v642 int32
	_ = v642
	var v643 int32
	_ = v643
	var v651 int32
	_ = v651
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v664 int32
	_ = v664
	var v666 int32
	_ = v666
	var v669 int32
	_ = v669
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
	var v702 int32
	_ = v702
	var v704 int32
	_ = v704
	var v716 int32
	_ = v716
	var v718 int32
	_ = v718
	var v721 int32
	_ = v721
	var v729 int32
	_ = v729
	var v731 int32
	_ = v731
	var v736 int32
	_ = v736
	var v751 int32
	_ = v751
	var v752 int32
	_ = v752
	var v756 int32
	_ = v756
	var v759 int32
	_ = v759
	var v762 int32
	_ = v762
	var v763 int32
	_ = v763
	var v770 int32
	_ = v770
	var v773 int32
	_ = v773
	var v779 int32
	_ = v779
	var v780 int32
	_ = v780
	var v781 int32
	_ = v781
	var v782 int32
	_ = v782
	var v785 int32
	_ = v785
	var v788 int32
	_ = v788
	var v789 int32
	_ = v789
	var v792 int32
	_ = v792
	var v799 int32
	_ = v799
	var v804 int32
	_ = v804
	var v806 int32
	_ = v806
	var v810 int32
	_ = v810
	var v812 int32
	_ = v812
	var v815 int32
	_ = v815
	var v816 int32
	_ = v816
	var v817 int32
	_ = v817
	var v818 int32
	_ = v818
	var v819 int32
	_ = v819
	var v820 int32
	_ = v820
	var v825 int32
	_ = v825
	var v827 int32
	_ = v827
	var v828 int32
	_ = v828
	var v832 int32
	_ = v832
	var v833 int32
	_ = v833
	var v834 int32
	_ = v834
	var v837 int32
	_ = v837
	var v839 int32
	_ = v839
	var v850 int32
	_ = v850
	var v853 int32
	_ = v853
	var v854 int32
	_ = v854
	var v859 int32
	_ = v859
	var v868 int32
	_ = v868
	var v869 int32
	_ = v869
	var v870 int32
	_ = v870
	var v872 int32
	_ = v872
	var v873 int32
	_ = v873
	var v874 int32
	_ = v874
	var v878 int32
	_ = v878
	var v879 int32
	_ = v879
	var v884 int32
	_ = v884
	var v886 int32
	_ = v886
	var v887 int32
	_ = v887
	var v888 int32
	_ = v888
	var v897 int32
	_ = v897
	var v898 int32
	_ = v898
	var v899 int32
	_ = v899
	var v903 int32
	_ = v903
	var v904 int32
	_ = v904
	var v907 int32
	_ = v907
	var v911 int32
	_ = v911
	var v916 int32
	_ = v916
	var v917 int32
	_ = v917
	var v921 int32
	_ = v921
	var v922 int32
	_ = v922
	var v925 int32
	_ = v925
	var v926 int32
	_ = v926
	var v927 int32
	_ = v927
	var v932 int32
	_ = v932
	var v942 int32
	_ = v942
	var v943 int32
	_ = v943
	var v952 int32
	_ = v952
	var v958 int32
	_ = v958
	var v959 int32
	_ = v959
	var v963 int32
	_ = v963
	var v964 int32
	_ = v964
	var v968 int32
	_ = v968
	var v969 int32
	_ = v969
	var v972 int32
	_ = v972
	var v973 int32
	_ = v973
	var v974 int32
	_ = v974
	var v977 int32
	_ = v977
	var v985 int32
	_ = v985
	var v988 int32
	_ = v988
	var v991 int32
	_ = v991
	var v996 int32
	_ = v996
	var v1005 int32
	_ = v1005
	var v1006 int32
	_ = v1006
	var v1014 int32
	_ = v1014
	var v1023 int32
	_ = v1023
	var v1024 int32
	_ = v1024
	var v1027 int32
	_ = v1027
	var v1030 int32
	_ = v1030
	var v1031 int32
	_ = v1031
	var v1032 int32
	_ = v1032
	var v1036 int32
	_ = v1036
	var v1038 int32
	_ = v1038
	var v1039 int32
	_ = v1039
	var v1040 int32
	_ = v1040
	var v1041 int32
	_ = v1041
	var v1042 int32
	_ = v1042
	var v1043 int32
	_ = v1043
	var v1048 int32
	_ = v1048
	var v1050 int32
	_ = v1050
	var v1052 int32
	_ = v1052
	var v1054 int32
	_ = v1054
	var v1055 int32
	_ = v1055
	var v1059 int32
	_ = v1059
	var v1063 int32
	_ = v1063
	var v1065 int32
	_ = v1065
	var v1068 int32
	_ = v1068
	var v1075 int32
	_ = v1075
	var v1077 int32
	_ = v1077
	var v1080 int32
	_ = v1080
	var v1083 int32
	_ = v1083
	var v1084 int32
	_ = v1084
	var v1086 int32
	_ = v1086
	var v1088 int32
	_ = v1088
	var v1089 int32
	_ = v1089
	var v1092 int32
	_ = v1092
	var v1093 int32
	_ = v1093
	var v1095 int32
	_ = v1095
	var v1097 int32
	_ = v1097
	var v1100 int32
	_ = v1100
	var v1104 int32
	_ = v1104
	var v1105 int32
	_ = v1105
	var v1109 int32
	_ = v1109
	var v1117 int32
	_ = v1117
	var v1118 int32
	_ = v1118
	var v1121 int32
	_ = v1121
	var v1122 int32
	_ = v1122
	var v1123 int32
	_ = v1123
	var v1125 int32
	_ = v1125
	var v1126 int32
	_ = v1126
	var v1129 int32
	_ = v1129
	var v1130 int32
	_ = v1130
	var v1132 int32
	_ = v1132
	var v1135 int32
	_ = v1135
	var v1136 int32
	_ = v1136
	var v1137 int32
	_ = v1137
	var v1140 int32
	_ = v1140
	var v1149 int32
	_ = v1149
	var v1150 int32
	_ = v1150
	var v1151 int32
	_ = v1151
	var v1152 int32
	_ = v1152
	var v1153 int32
	_ = v1153
	var v1157 int32
	_ = v1157
	var v1158 int32
	_ = v1158
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
	var v1167 int32
	_ = v1167
	var v1168 int32
	_ = v1168
	var v1169 int32
	_ = v1169
	var v1170 int32
	_ = v1170
	var v1171 int32
	_ = v1171
	var v1174 int32
	_ = v1174
	var v1176 int32
	_ = v1176
	var v1177 int32
	_ = v1177
	var v1181 int32
	_ = v1181
	var v1182 int32
	_ = v1182
	var v1188 int32
	_ = v1188
	var v1189 int32
	_ = v1189
	var v1190 int32
	_ = v1190
	var v1192 int32
	_ = v1192
	var v1193 int32
	_ = v1193
	var v1195 int32
	_ = v1195
	var v1197 int32
	_ = v1197
	var v1200 int32
	_ = v1200
	var v1201 int32
	_ = v1201
	var v1203 int32
	_ = v1203
	var v1205 int32
	_ = v1205
	var v1206 int32
	_ = v1206
	var v1210 int32
	_ = v1210
	var v1211 int32
	_ = v1211
	var v1212 int32
	_ = v1212
	var v1216 int32
	_ = v1216
	var v1217 int32
	_ = v1217
	var v1220 int32
	_ = v1220
	var v1221 int32
	_ = v1221
	var v1222 int32
	_ = v1222
	var v1230 int32
	_ = v1230
	var v1231 int32
	_ = v1231
	var v1233 int32
	_ = v1233
	var v1236 int32
	_ = v1236
	var v1240 int32
	_ = v1240
	var v1250 int32
	_ = v1250
	var v1251 int32
	_ = v1251
	var v1252 int32
	_ = v1252
	var v1254 int32
	_ = v1254
	var v1257 int32
	_ = v1257
	var v1261 int32
	_ = v1261
	var v1264 int32
	_ = v1264
	var v1265 int32
	_ = v1265
	var v1270 int32
	_ = v1270
	var v1272 int32
	_ = v1272
	var v1274 int32
	_ = v1274
	var v1277 int32
	_ = v1277
	var v1289 int32
	_ = v1289
	var v1292 int32
	_ = v1292
	var v1293 int32
	_ = v1293
	var v1295 int32
	_ = v1295
	var v1296 int32
	_ = v1296
	var v1300 int32
	_ = v1300
	var v1301 int32
	_ = v1301
	var v1306 int32
	_ = v1306
	var v1316 int32
	_ = v1316
	var v1323 int32
	_ = v1323
	var v1327 int32
	_ = v1327
	var v1329 int32
	_ = v1329
	var v1331 int32
	_ = v1331
	var v1332 int32
	_ = v1332
	var v1333 int32
	_ = v1333
	var v1334 int32
	_ = v1334
	var v1337 int32
	_ = v1337
	var v1338 int32
	_ = v1338
	var v1347 int32
	_ = v1347
	var v1349 int32
	_ = v1349
	var v1351 int32
	_ = v1351
	var v1355 int32
	_ = v1355
	var v1365 int32
	_ = v1365
	var v1366 int32
	_ = v1366
	var v1367 int32
	_ = v1367
	var v1372 int32
	_ = v1372
	var v1376 int32
	_ = v1376
	var v1377 int32
	_ = v1377
	var v1381 int32
	_ = v1381
	var v1382 int32
	_ = v1382
	var v1387 int32
	_ = v1387
	var v1389 int32
	_ = v1389
	var v1391 int32
	_ = v1391
	var v1394 int32
	_ = v1394
	var v1406 int32
	_ = v1406
	var v1409 int32
	_ = v1409
	var v1410 int32
	_ = v1410
	var v1412 int32
	_ = v1412
	var v1413 int32
	_ = v1413
	var v1417 int32
	_ = v1417
	var v1418 int32
	_ = v1418
	var v1423 int32
	_ = v1423
	var v1433 int32
	_ = v1433
	var v1440 int32
	_ = v1440
	var v1444 int32
	_ = v1444
	var v1446 int32
	_ = v1446
	var v1450 int32
	_ = v1450
	var v1455 int32
	_ = v1455
	var v1459 int32
	_ = v1459
	var v1472 int32
	_ = v1472
	var v1476 int32
	_ = v1476
	var v1481 int32
	_ = v1481
	var v1491 int32
	_ = v1491
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
	m.G0 = v1491 + int32(16)
	return v1481
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
	v1481 = v90
	v1491 = v16
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
	if v199 != int32(285) {
		v110 = v199
		v117 = v117 + v193
		goto L27
	} else {
		goto L41
	}
L35:
	;
	v159 = int32(*(*int16)(unsafe.Add(mBase, uint32(v142<<(uint(int32(1))%32))+uint32(_c_F_replication_yylex[7]))))
	if int32(286) <= v159 {
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
		goto L95
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
		v1481 = int32(266)
		v1491 = v120
		goto L21
	case 7:
		goto L94
	case 8:
		goto L93
	case 9:
		goto L92
	case 10:
		goto L91
	case 11:
		goto L90
	case 12:
		goto L89
	case 13:
		goto L88
	case 14:
		goto L87
	case 15:
		goto L86
	case 16:
		goto L85
	case 17:
		goto L84
	case 18:
		goto L83
	case 19:
		goto L82
	case 20:
		goto L81
	case 21:
		goto L80
	case 22:
		v94 = v111
		v103 = v120
		goto L25
	case 23:
		goto L79
	case 24:
		goto L78
	case 25:
		goto L77
	case 26:
		goto L76
	case 27:
		goto L75
	case 28:
		goto L74
	case 29:
		goto L73
	case 30:
		goto L72
	case 31:
		goto L71
	case 32:
		goto L70
	case 33:
		goto L69
	case 34:
		goto L68
	case 35:
		goto L65
	case 36:
		goto L64
	case 37:
		goto L66
	case 38, 39:
		goto L67
	default:
		goto L63
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v111)+36)) = v1459
	*(*int32)(unsafe.Add(mBase, uint32(v111)+48)) = int32(0)
	v1472 = *(*int32)(unsafe.Add(mBase, uint32(v111)+44))
	v1476 = base.I32_div_s(v1472-int32(1), int32(2))
	v245 = v1476 + int32(37)
	goto L46
L49:
	;
	F_yy_fatal_error_3(m, int32(_a_F_replication_yylex_1))
	mBase = m.M
	v1455 = m.ExcPending
	if v1455 != 0 {
		goto L18
	} else {
		goto L298
	}
L50:
	;
	F_yy_fatal_error_3(m, int32(_a_F_replication_yylex_2))
	mBase = m.M
	v1450 = m.ExcPending
	if v1450 != 0 {
		goto L18
	} else {
		goto L297
	}
L51:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L52:
	;
	v1347 = v1334 + v1338
	*(*int32)(unsafe.Add(mBase, uint32(v111)+36)) = v1347
	v1349 = *(*int32)(unsafe.Add(mBase, uint32(v111)+44))
	if base.Ui32(v1347) <= base.Ui32(v1337) {
		v217 = v1349
		v224 = v1347
		v225 = v1337
		goto L44
	} else {
		goto L278
	}
L53:
	;
	v1024 = *(*int32)(unsafe.Add(mBase, uint32(v1023)))
	*(*int32)(unsafe.Add(mBase, uint32(v1024)+16)) = v1014
	v1027 = *(*int32)(unsafe.Add(mBase, uint32(v111)+28))
	if v1027 != 0 {
		v1149 = int32(0)
		goto L222
	} else {
		goto L223
	}
L54:
	;
	v1005 = *(*int32)(unsafe.Add(mBase, uint32(v111)+20))
	v1006 = *(*int32)(unsafe.Add(mBase, uint32(v111)+12))
	v1014 = v996
	v1023 = v1005 + v1006<<(uint(int32(2))%32)
	goto L53
L55:
	;
	F_yy_fatal_error_3(m, int32(_a_F_replication_yylex_3))
	mBase = m.M
	v991 = m.ExcPending
	if v991 != 0 {
		goto L18
	} else {
		goto L221
	}
L56:
	;
	F_yy_fatal_error_3(m, int32(_a_F_replication_yylex_4))
	mBase = m.M
	v988 = m.ExcPending
	if v988 != 0 {
		goto L18
	} else {
		goto L220
	}
L57:
	;
	F_replication_yyerror(m, int32(_a_F_replication_yylex_5))
	mBase = m.M
	v985 = m.ExcPending
	if v985 != 0 {
		goto L18
	} else {
		goto L219
	}
L58:
	;
	v1481 = int32(272)
	v1491 = v120
	goto L21
L59:
	;
	v1481 = int32(265)
	v1491 = v120
	goto L21
L60:
	;
	v1481 = int32(264)
	v1491 = v120
	goto L21
L61:
	;
	v1481 = int32(263)
	v1491 = v120
	goto L21
L62:
	;
	v1481 = int32(262)
	v1491 = v120
	goto L21
L63:
	;
	F_yy_fatal_error_3(m, int32(_a_F_replication_yylex_6))
	mBase = m.M
	v977 = m.ExcPending
	if v977 != 0 {
		goto L18
	} else {
		goto L218
	}
L64:
	;
	v406 = *(*int32)(unsafe.Add(mBase, uint32(v111)+80))
	v407 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111)+24)))
	*(*uint8)(unsafe.Add(mBase, uint32(v224))) = uint8(v407)
	v409 = *(*int32)(unsafe.Add(mBase, uint32(v111)+20))
	v410 = *(*int32)(unsafe.Add(mBase, uint32(v111)+12))
	v413 = v409 + v410<<(uint(int32(2))%32)
	v414 = *(*int32)(unsafe.Add(mBase, uint32(v413)))
	v415 = *(*int32)(unsafe.Add(mBase, uint32(v414)+44))
	if v415 == int32(0) {
		goto L109
	} else {
		goto L110
	}
L65:
	;
	F_yy_fatal_error_3(m, int32(_a_F_replication_yylex_7))
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
		goto L18
	} else {
		goto L108
	}
L66:
	;
	v1481 = int32(0)
	v1491 = v120
	goto L21
L67:
	;
	F_replication_yyerror(m, int32(_a_F_replication_yylex_8))
	mBase = m.M
	v401 = m.ExcPending
	if v401 != 0 {
		goto L18
	} else {
		goto L107
	}
L68:
	;
	v397 = *(*int32)(unsafe.Add(mBase, uint32(v111)+80))
	v398 = int32(*(*int8)(unsafe.Add(mBase, uint32(v397))))
	v1481 = v398
	v1491 = v120
	goto L21
L69:
	;
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v111)+80))
	v390 = F_strlen(m, v389)
	mBase = m.M
	v392 = F_downcase_truncate_identifier(m, v389, v390, int32(1))
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L18
	} else {
		goto L106
	}
L70:
	;
	v382 = *(*int32)(unsafe.Add(mBase, uint32(v111)))
	v385 = *(*int32)(unsafe.Add(mBase, uint32(v111)+80))
	v386 = *(*int32)(unsafe.Add(mBase, uint32(v111)+32))
	F_appendBinaryStringInfo(m, v382+int32(4), v385, v386)
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L18
	} else {
		goto L105
	}
L71:
	;
	v376 = *(*int32)(unsafe.Add(mBase, uint32(v111)))
	F_appendStringInfoChar(m, v376+int32(4), int32(34))
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L18
	} else {
		goto L104
	}
L72:
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
		goto L103
	}
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v111)+44)) = int32(3)
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v111)))
	F_initStringInfo(m, v345+int32(4))
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L18
	} else {
		goto L102
	}
L74:
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
		goto L101
	}
L75:
	;
	v330 = *(*int32)(unsafe.Add(mBase, uint32(v111)))
	F_appendStringInfoChar(m, v330+int32(4), int32(39))
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L18
	} else {
		goto L100
	}
L76:
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
	v1481 = int32(258)
	v1491 = v120
	goto L21
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v111)+44)) = int32(5)
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v111)))
	F_initStringInfo(m, v305+int32(4))
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L18
	} else {
		goto L99
	}
L78:
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
		goto L97
	}
L79:
	;
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v111)+80))
	v278 = F_strtox_2(m, v274, int32(0), int32(10), int64(4294967295))
	mBase = m.M
	goto L96
L80:
	;
	v1481 = int32(282)
	v1491 = v120
	goto L21
L81:
	;
	v1481 = int32(271)
	v1491 = v120
	goto L21
L82:
	;
	v1481 = int32(281)
	v1491 = v120
	goto L21
L83:
	;
	v1481 = int32(280)
	v1491 = v120
	goto L21
L84:
	;
	v1481 = int32(279)
	v1491 = v120
	goto L21
L85:
	;
	v1481 = int32(278)
	v1491 = v120
	goto L21
L86:
	;
	v1481 = int32(277)
	v1491 = v120
	goto L21
L87:
	;
	v1481 = int32(275)
	v1491 = v120
	goto L21
L88:
	;
	v1481 = int32(274)
	v1491 = v120
	goto L21
L89:
	;
	v1481 = int32(276)
	v1491 = v120
	goto L21
L90:
	;
	v1481 = int32(273)
	v1491 = v120
	goto L21
L91:
	;
	v1481 = int32(270)
	v1491 = v120
	goto L21
L92:
	;
	v1481 = int32(269)
	v1491 = v120
	goto L21
L93:
	;
	v1481 = int32(268)
	v1491 = v120
	goto L21
L94:
	;
	v1481 = int32(267)
	v1491 = v120
	goto L21
L95:
	;
	v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111)+24)))
	*(*uint8)(unsafe.Add(mBase, uint32(v224))) = uint8(v257)
	v210 = v225
	goto L42
L96:
	;
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v111)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v280))) = base.I32_wrap_i64(v278)
	v1481 = int32(260)
	v1491 = v120
	goto L21
L97:
	;
	if v291 != int32(2) {
		goto L57
	} else {
		goto L98
	}
L98:
	;
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v111)+92))
	v296 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v120)+8)))
	v297 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v120)+12)))
	*(*int64)(unsafe.Add(mBase, uint32(v295))) = v296 | v297<<(uint(int64(32))%64)
	v1481 = int32(261)
	v1491 = v120
	goto L21
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
	v94 = v111
	v103 = v120
	goto L25
L103:
	;
	v1481 = int32(259)
	v1491 = v120
	goto L21
L104:
	;
	v94 = v111
	v103 = v120
	goto L25
L105:
	;
	v94 = v111
	v103 = v120
	goto L25
L106:
	;
	v394 = *(*int32)(unsafe.Add(mBase, uint32(v111)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v394))) = v392
	v1481 = int32(259)
	v1491 = v120
	goto L21
L107:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L108:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L109:
	;
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v414)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v111)+28)) = v418
	v420 = *(*int32)(unsafe.Add(mBase, uint32(v413)))
	v421 = *(*int32)(unsafe.Add(mBase, uint32(v111)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v420))) = v421
	v423 = *(*int32)(unsafe.Add(mBase, uint32(v111)+20))
	v424 = *(*int32)(unsafe.Add(mBase, uint32(v111)+12))
	v425 = int32(2)
	v428 = *(*int32)(unsafe.Add(mBase, uint32(v423+v424<<(uint(v425)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v428)+44)) = int32(1)
	v431 = *(*int32)(unsafe.Add(mBase, uint32(v111)+20))
	v432 = *(*int32)(unsafe.Add(mBase, uint32(v111)+12))
	v436 = *(*int32)(unsafe.Add(mBase, uint32(v431+v432<<(uint(v425)%32))))
	v437 = v436
	v438 = v431
	v439 = v432
	goto L111
L110:
	;
	v437 = v414
	v438 = v409
	v439 = v410
	goto L111
L111:
	;
	v440 = *(*int32)(unsafe.Add(mBase, uint32(v111)+36))
	v441 = *(*int32)(unsafe.Add(mBase, uint32(v437)+4))
	v442 = *(*int32)(unsafe.Add(mBase, uint32(v111)+28))
	v443 = v441 + v442
	if base.Ui32(v440) <= base.Ui32(v443) {
		goto L112
	} else {
		goto L113
	}
L112:
	;
	v445 = *(*int32)(unsafe.Add(mBase, uint32(v111)+80))
	v448 = v406 ^ int32(-1) + v224
	v449 = v445 + v448
	*(*int32)(unsafe.Add(mBase, uint32(v111)+36)) = v449
	v451 = *(*int32)(unsafe.Add(mBase, uint32(v111)+44))
	if int32(0) < v448 {
		goto L115
	} else {
		goto L116
	}
L113:
	;
	goto L114
L114:
	;
	if base.Ui32(v443+int32(1)) < base.Ui32(v440) {
		goto L56
	} else {
		goto L147
	}
L115:
	;
	v454 = v451
	v461 = v445
	goto L118
L116:
	;
	v549 = v451
	goto L117
L117:
	;
	v566 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v549<<(uint(int32(1))%32))+uint32(_c_F_replication_yylex[3]))))
	if v566 != 0 {
		goto L136
	} else {
		goto L137
	}
L118:
	;
	v468 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v461))))
	if v468 != 0 {
		goto L120
	} else {
		goto L121
	}
L119:
	;
	v549 = v545
	goto L117
L120:
	;
	v469 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v468)+uint32(_c_F_replication_yylex[2]))))
	v470 = v469
	goto L122
L121:
	;
	v470 = int32(1)
	goto L122
L122:
	;
	v472 = v454 << (uint(int32(1)) % 32)
	v475 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v472)+uint32(_c_F_replication_yylex[3]))))
	if v475 != 0 {
		goto L123
	} else {
		goto L124
	}
L123:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v111)+68)) = v461
	*(*int32)(unsafe.Add(mBase, uint32(v111)+64)) = v454
	goto L125
L124:
	;
	goto L125
L125:
	;
	v479 = v470 & int32(255)
	v482 = int32(*(*int16)(unsafe.Add(mBase, uint32(v472)+uint32(_c_F_replication_yylex[4]))))
	v483 = v479 + v482
	v488 = int32(*(*int16)(unsafe.Add(mBase, uint32(v483<<(uint(int32(1))%32))+uint32(_c_F_replication_yylex[5]))))
	if v488 != v454 {
		goto L126
	} else {
		goto L127
	}
L126:
	;
	v490 = v454
	v492 = v470
	v495 = v479
	goto L129
L127:
	;
	v534 = v483
	goto L128
L128:
	;
	v541 = int32(1)
	v545 = int32(*(*int16)(unsafe.Add(mBase, uint32(v534<<(uint(v541)%32))+uint32(_c_F_replication_yylex[6]))))
	v547 = v461 + v541
	if v547 != v449 {
		v454 = v545
		v461 = v547
		goto L118
	} else {
		goto L135
	}
L129:
	;
	v507 = int32(*(*int16)(unsafe.Add(mBase, uint32(v490<<(uint(int32(1))%32))+uint32(_c_F_replication_yylex[7]))))
	if int32(286) <= v507 {
		goto L131
	} else {
		goto L132
	}
L130:
	;
	v534 = v519
	goto L128
L131:
	;
	v510 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v495)+uint32(_c_F_replication_yylex[8]))))
	v511 = v510
	goto L133
L132:
	;
	v511 = v492
	goto L133
L133:
	;
	v513 = v511 & int32(255)
	v514 = int32(1)
	v518 = int32(*(*int16)(unsafe.Add(mBase, uint32(v507<<(uint(v514)%32))+uint32(_c_F_replication_yylex[4]))))
	v519 = v513 + v518
	v524 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v519<<(uint(v514)%32))+uint32(_c_F_replication_yylex[5]))))
	if v524 != v507&int32(_a_F_replication_yylex_0) {
		v490 = v507
		v492 = v511
		v495 = v513
		goto L129
	} else {
		goto L134
	}
L134:
	;
	goto L130
L135:
	;
	goto L119
L136:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v111)+68)) = v449
	*(*int32)(unsafe.Add(mBase, uint32(v111)+64)) = v549
	goto L138
L137:
	;
	goto L138
L138:
	;
	v569 = int32(1)
	v573 = int32(*(*int16)(unsafe.Add(mBase, uint32(v549<<(uint(v569)%32))+uint32(_c_F_replication_yylex[4]))))
	v575 = v573 + v569
	v580 = int32(*(*int16)(unsafe.Add(mBase, uint32(v575<<(uint(v569)%32))+uint32(_c_F_replication_yylex[5]))))
	if v580 != v549 {
		goto L139
	} else {
		goto L140
	}
L139:
	;
	v582 = v549
	goto L142
L140:
	;
	v616 = v575
	goto L141
L141:
	;
	if v616 == int32(0) {
		v210 = v445
		goto L42
	} else {
		goto L145
	}
L142:
	;
	v595 = int32(1)
	v599 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v582<<(uint(v595)%32))+uint32(_c_F_replication_yylex[7]))))
	v600 = base.I32_extend16_s(v599)
	v605 = int32(*(*int16)(unsafe.Add(mBase, uint32(v600<<(uint(v595)%32))+uint32(_c_F_replication_yylex[4]))))
	v607 = v605 + v595
	v612 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v607<<(uint(v595)%32))+uint32(_c_F_replication_yylex[5]))))
	if v599 != v612 {
		v582 = v600
		goto L142
	} else {
		goto L144
	}
L143:
	;
	v616 = v607
	goto L141
L144:
	;
	goto L143
L145:
	;
	v633 = int32(*(*int16)(unsafe.Add(mBase, uint32(v616<<(uint(int32(1))%32))+uint32(_c_F_replication_yylex[6]))))
	if v633 == int32(285) {
		v210 = v445
		goto L42
	} else {
		goto L146
	}
L146:
	;
	v637 = v449 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v111)+36)) = v637
	v110 = v633
	v113 = v445
	v117 = v637
	goto L27
L147:
	;
	v642 = *(*int32)(unsafe.Add(mBase, uint32(v111)+80))
	v643 = *(*int32)(unsafe.Add(mBase, uint32(v437)+40))
	if v643 == int32(0) {
		goto L148
	} else {
		goto L149
	}
L148:
	;
	if v440-v642 != int32(1) {
		v1334 = v441
		v1337 = v642
		v1338 = v442
		goto L52
	} else {
		goto L151
	}
L149:
	;
	goto L150
L150:
	;
	v651 = v642 ^ int32(-1) + v440
	if int32(0) < v651 {
		goto L152
	} else {
		goto L153
	}
L151:
	;
	v1459 = v642
	goto L48
L152:
	;
	v654 = int32(7)
	v655 = v651 & v654
	if base.Ui32(v440-v642-int32(2)) < base.Ui32(v654) {
		goto L157
	} else {
		goto L158
	}
L153:
	;
	v759 = v437
	v762 = v438
	v763 = v439
	goto L154
L154:
	;
	v770 = *(*int32)(unsafe.Add(mBase, uint32(v759)+44))
	if v770 == int32(2) {
		goto L167
	} else {
		goto L168
	}
L155:
	;
	v751 = *(*int32)(unsafe.Add(mBase, uint32(v111)+20))
	v752 = *(*int32)(unsafe.Add(mBase, uint32(v111)+12))
	v756 = *(*int32)(unsafe.Add(mBase, uint32(v751+v752<<(uint(int32(2))%32))))
	v759 = v756
	v762 = v751
	v763 = v752
	goto L154
L156:
	;
	v716 = v702
	v718 = v704
	v721 = int32(0)
	goto L164
L157:
	;
	v702 = v441
	v704 = v642
	goto L156
L158:
	;
	goto L159
L159:
	;
	v664 = v441
	v666 = v642
	v669 = int32(0)
	goto L160
L160:
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
	v698 = v669 + v693
	if v698 != v651&int32(2147483640) {
		v664 = v694
		v666 = v696
		v669 = v698
		goto L160
	} else {
		goto L162
	}
L161:
	;
	if v655 == int32(0) {
		goto L155
	} else {
		goto L163
	}
L162:
	;
	goto L161
L163:
	;
	v702 = v694
	v704 = v696
	goto L156
L164:
	;
	v729 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v718))))
	*(*uint8)(unsafe.Add(mBase, uint32(v716))) = uint8(v729)
	v731 = int32(1)
	v736 = v721 + v731
	if v736 != v655 {
		v716 = v716 + v731
		v718 = v718 + v731
		v721 = v736
		goto L164
	} else {
		goto L166
	}
L165:
	;
	goto L155
L166:
	;
	goto L165
L167:
	;
	v773 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v111)+28)) = v773
	v1014 = v773
	v1023 = v762 + v763<<(uint(int32(2))%32)
	goto L53
L168:
	;
	goto L169
L169:
	;
	v779 = int32(0)
	v780 = *(*int32)(unsafe.Add(mBase, uint32(v759)+12))
	v781 = v642 - v440
	v782 = v780 + v781
	if v782 <= v779 {
		goto L170
	} else {
		goto L171
	}
L170:
	;
	v785 = *(*int32)(unsafe.Add(mBase, uint32(v111)+36))
	v788 = v759
	v789 = v785
	v792 = v780
	goto L173
L171:
	;
	v837 = v782
	v839 = v759
	goto L172
L172:
	;
	v850 = int32(_a_F_replication_yylex_10)
	if base.Ui32(v850) <= base.Ui32(v837) {
		goto L189
	} else {
		goto L190
	}
L173:
	;
	v799 = *(*int32)(unsafe.Add(mBase, uint32(v788)+20))
	if v799 == int32(0) {
		goto L175
	} else {
		goto L176
	}
L174:
	;
	v837 = v834
	v839 = v832
	goto L172
L175:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v788)+4)) = int32(0)
	goto L49
L176:
	;
	goto L177
L177:
	;
	v804 = *(*int32)(unsafe.Add(mBase, uint32(v788)+4))
	v806 = v792 << (uint(int32(1)) % 32)
	if v806 <= int32(0) {
		goto L178
	} else {
		goto L179
	}
L178:
	;
	v810 = base.I32_div_s(v792, int32(8))
	v812 = v810 + v792
	goto L180
L179:
	;
	v812 = v806
	goto L180
L180:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v788)+12)) = v812
	v815 = v812 + int32(2)
	if v804 != 0 {
		goto L182
	} else {
		goto L183
	}
L181:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v788)+4)) = v820
	if v820 == int32(0) {
		goto L49
	} else {
		goto L187
	}
L182:
	;
	v816 = F_repalloc(m, v804, v815)
	mBase = m.M
	v817 = m.ExcPending
	if v817 != 0 {
		goto L18
	} else {
		goto L185
	}
L183:
	;
	goto L184
L184:
	;
	v818 = F_palloc(m, v815)
	mBase = m.M
	v819 = m.ExcPending
	if v819 != 0 {
		goto L18
	} else {
		goto L186
	}
L185:
	;
	v820 = v816
	goto L181
L186:
	;
	v820 = v818
	goto L181
L187:
	;
	v825 = v820 + (v789 - v804)
	*(*int32)(unsafe.Add(mBase, uint32(v111)+36)) = v825
	v827 = *(*int32)(unsafe.Add(mBase, uint32(v111)+20))
	v828 = *(*int32)(unsafe.Add(mBase, uint32(v111)+12))
	v832 = *(*int32)(unsafe.Add(mBase, uint32(v827+v828<<(uint(int32(2))%32))))
	v833 = *(*int32)(unsafe.Add(mBase, uint32(v832)+12))
	v834 = v833 + v781
	if v834 <= int32(0) {
		v788 = v832
		v789 = v825
		v792 = v833
		goto L173
	} else {
		goto L188
	}
L188:
	;
	goto L174
L189:
	;
	v853 = v850
	goto L191
L190:
	;
	v853 = v837
	goto L191
L191:
	;
	v854 = *(*int32)(unsafe.Add(mBase, uint32(v839)+24))
	if v854 != 0 {
		goto L192
	} else {
		goto L193
	}
L192:
	;
	v859 = v779
	goto L196
L193:
	;
	goto L194
L194:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_replication_yylex[9])) = int32(0)
	v916 = *(*int32)(unsafe.Add(mBase, uint32(v111)+20))
	v917 = *(*int32)(unsafe.Add(mBase, uint32(v111)+12))
	v921 = *(*int32)(unsafe.Add(mBase, uint32(v916+v917<<(uint(int32(2))%32))))
	v922 = *(*int32)(unsafe.Add(mBase, uint32(v921)+4))
	v925 = *(*int32)(unsafe.Add(mBase, uint32(v111)+4))
	v926 = F_fread(m, v922+v651, int32(1), v853, v925)
	mBase = m.M
	v927 = m.ExcPending
	if v927 != 0 {
		goto L18
	} else {
		goto L207
	}
L195:
	;
	switch v872 {
	case 0:
		goto L203
	default:
		v911 = v886
		goto L201
	case 11:
		goto L202
	}
L196:
	;
	v868 = *(*int32)(unsafe.Add(mBase, uint32(v111)+4))
	v869 = F_do_getc(m, v868)
	mBase = m.M
	v870 = m.ExcPending
	if v870 != 0 {
		goto L18
	} else {
		goto L199
	}
L197:
	;
	v886 = v853
	goto L195
L198:
	;
	v873 = *(*int32)(unsafe.Add(mBase, uint32(v111)+20))
	v874 = *(*int32)(unsafe.Add(mBase, uint32(v111)+12))
	v878 = *(*int32)(unsafe.Add(mBase, uint32(v873+v874<<(uint(int32(2))%32))))
	v879 = *(*int32)(unsafe.Add(mBase, uint32(v878)+4))
	*(*uint8)(unsafe.Add(mBase, uint32(v879+v651+v859))) = uint8(v869)
	v884 = v859 + int32(1)
	if v884 != v853 {
		v859 = v884
		goto L196
	} else {
		goto L200
	}
L199:
	;
	v872 = v869 + int32(1)
	switch v872 {
	case 0, 11:
		v886 = v859
		goto L195
	default:
		goto L198
	}
L200:
	;
	goto L197
L201:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v111)+28)) = v911
	v996 = v911
	goto L54
L202:
	;
	v898 = *(*int32)(unsafe.Add(mBase, uint32(v111)+20))
	v899 = *(*int32)(unsafe.Add(mBase, uint32(v111)+12))
	v903 = *(*int32)(unsafe.Add(mBase, uint32(v898+v899<<(uint(int32(2))%32))))
	v904 = *(*int32)(unsafe.Add(mBase, uint32(v903)+4))
	v907 = int32(10)
	*(*uint8)(unsafe.Add(mBase, uint32(v904+v651+v886))) = uint8(v907)
	v911 = v886 + int32(1)
	goto L201
L203:
	;
	v887 = *(*int32)(unsafe.Add(mBase, uint32(v111)+4))
	v888 = *(*int32)(unsafe.Add(mBase, uint32(v887)))
	goto L204
L204:
	;
	if int32(base.Ui32(v888)>>(uint(int32(5))%32))&int32(1) == int32(0) {
		v911 = v886
		goto L201
	} else {
		goto L205
	}
L205:
	;
	F_yy_fatal_error_3(m, int32(_a_F_replication_yylex_3))
	mBase = m.M
	v897 = m.ExcPending
	if v897 != 0 {
		goto L18
	} else {
		goto L206
	}
L206:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L207:
	;
	v932 = v926
	goto L208
L208:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v111)+28)) = v932
	if v932 != 0 {
		v996 = v932
		goto L54
	} else {
		goto L210
	}
L210:
	;
	v942 = *(*int32)(unsafe.Add(mBase, uint32(v111)+4))
	v943 = *(*int32)(unsafe.Add(mBase, uint32(v942)))
	goto L211
L211:
	;
	if int32(base.Ui32(v943)>>(uint(int32(5))%32))&int32(1) == int32(0) {
		goto L212
	} else {
		goto L213
	}
L212:
	;
	v996 = int32(0)
	goto L54
L213:
	;
	goto L214
L214:
	;
	v952 = *(*int32)(unsafe.Add(mBase, _c_F_replication_yylex[9]))
	if v952 != int32(27) {
		goto L55
	} else {
		goto L215
	}
L215:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_replication_yylex[9])) = int32(0)
	v958 = *(*int32)(unsafe.Add(mBase, uint32(v111)+4))
	v959 = *(*int32)(unsafe.Add(mBase, uint32(v958)))
	*(*int32)(unsafe.Add(mBase, uint32(v958))) = v959 & int32(-49)
	goto L216
L216:
	;
	v963 = *(*int32)(unsafe.Add(mBase, uint32(v111)+20))
	v964 = *(*int32)(unsafe.Add(mBase, uint32(v111)+12))
	v968 = *(*int32)(unsafe.Add(mBase, uint32(v963+v964<<(uint(int32(2))%32))))
	v969 = *(*int32)(unsafe.Add(mBase, uint32(v968)+4))
	v972 = *(*int32)(unsafe.Add(mBase, uint32(v111)+4))
	v973 = F_fread(m, v969+v651, int32(1), v853, v972)
	mBase = m.M
	v974 = m.ExcPending
	if v974 != 0 {
		goto L18
	} else {
		goto L217
	}
L217:
	;
	v932 = v973
	goto L208
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
	base.Wasm_trap_unreachable()
	for {
	}
L221:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L222:
	;
	v1150 = *(*int32)(unsafe.Add(mBase, uint32(v111)+28))
	v1151 = v1150 + v651
	v1152 = *(*int32)(unsafe.Add(mBase, uint32(v111)+20))
	v1153 = *(*int32)(unsafe.Add(mBase, uint32(v111)+12))
	v1157 = *(*int32)(unsafe.Add(mBase, uint32(v1152+v1153<<(uint(int32(2))%32))))
	v1158 = *(*int32)(unsafe.Add(mBase, uint32(v1157)+12))
	if v1158 < v1151 {
		goto L246
	} else {
		goto L247
	}
L223:
	;
	if v651 == int32(0) {
		goto L224
	} else {
		goto L225
	}
L224:
	;
	v1030 = *(*int32)(unsafe.Add(mBase, uint32(v111)+4))
	v1031 = *(*int32)(unsafe.Add(mBase, uint32(v111)+20))
	if v1031 != 0 {
		goto L229
	} else {
		goto L230
	}
L225:
	;
	goto L226
L226:
	;
	v1135 = *(*int32)(unsafe.Add(mBase, uint32(v111)+20))
	v1136 = *(*int32)(unsafe.Add(mBase, uint32(v111)+12))
	v1137 = int32(2)
	v1140 = *(*int32)(unsafe.Add(mBase, uint32(v1135+v1136<<(uint(v1137)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v1140)+44)) = v1137
	v1149 = v1137
	goto L222
L227:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1097)+40)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1097))) = v1030
	v1104 = *(*int32)(unsafe.Add(mBase, uint32(v111)+20))
	if v1104 != 0 {
		goto L242
	} else {
		goto L243
	}
L228:
	;
	v1054 = *(*int32)(unsafe.Add(mBase, _c_F_replication_yylex[9]))
	v1055 = *(*int32)(unsafe.Add(mBase, uint32(v111)+12))
	v1059 = *(*int32)(unsafe.Add(mBase, uint32(v1052+v1055<<(uint(int32(2))%32))))
	if v1059 == int32(0) {
		goto L236
	} else {
		goto L237
	}
L229:
	;
	v1032 = *(*int32)(unsafe.Add(mBase, uint32(v111)+12))
	v1036 = *(*int32)(unsafe.Add(mBase, uint32(v1031+v1032<<(uint(int32(2))%32))))
	if v1036 != 0 {
		v1052 = v1031
		goto L228
	} else {
		goto L232
	}
L230:
	;
	goto L231
L231:
	;
	F_replication_yyensure_buffer_stack(m, v111)
	mBase = m.M
	v1038 = m.ExcPending
	if v1038 != 0 {
		goto L18
	} else {
		goto L233
	}
L232:
	;
	goto L231
L233:
	;
	v1039 = *(*int32)(unsafe.Add(mBase, uint32(v111)+4))
	v1040 = F_replication_yy_create_buffer(m, v1039, v111)
	mBase = m.M
	v1041 = m.ExcPending
	if v1041 != 0 {
		goto L18
	} else {
		goto L234
	}
L234:
	;
	v1042 = *(*int32)(unsafe.Add(mBase, uint32(v111)+20))
	v1043 = *(*int32)(unsafe.Add(mBase, uint32(v111)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v1042+v1043<<(uint(int32(2))%32)))) = v1040
	v1048 = *(*int32)(unsafe.Add(mBase, uint32(v111)+20))
	if v1048 != 0 {
		v1052 = v1048
		goto L228
	} else {
		goto L235
	}
L235:
	;
	v1050 = *(*int32)(unsafe.Add(mBase, _c_F_replication_yylex[9]))
	v1097 = int32(0)
	v1100 = v1050
	goto L227
L236:
	;
	v1097 = int32(0)
	v1100 = v1054
	goto L227
L237:
	;
	goto L238
L238:
	;
	v1063 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1059)+16)) = v1063
	v1065 = *(*int32)(unsafe.Add(mBase, uint32(v1059)+4))
	*(*uint8)(unsafe.Add(mBase, uint32(v1065))) = uint8(v1063)
	v1068 = *(*int32)(unsafe.Add(mBase, uint32(v1059)+4))
	*(*uint8)(unsafe.Add(mBase, uint32(v1068)+1)) = uint8(v1063)
	*(*int32)(unsafe.Add(mBase, uint32(v1059)+44)) = v1063
	*(*int32)(unsafe.Add(mBase, uint32(v1059)+28)) = int32(1)
	v1075 = *(*int32)(unsafe.Add(mBase, uint32(v1059)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1059)+8)) = v1075
	v1077 = *(*int32)(unsafe.Add(mBase, uint32(v111)+20))
	if v1077 == v1063 {
		v1097 = v1059
		v1100 = v1054
		goto L227
	} else {
		goto L239
	}
L239:
	;
	v1080 = *(*int32)(unsafe.Add(mBase, uint32(v111)+12))
	v1083 = v1077 + v1080<<(uint(int32(2))%32)
	v1084 = *(*int32)(unsafe.Add(mBase, uint32(v1083)))
	if v1059 != v1084 {
		v1097 = v1059
		v1100 = v1054
		goto L227
	} else {
		goto L240
	}
L240:
	;
	v1086 = *(*int32)(unsafe.Add(mBase, uint32(v1084)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v111)+28)) = v1086
	v1088 = *(*int32)(unsafe.Add(mBase, uint32(v1083)))
	v1089 = *(*int32)(unsafe.Add(mBase, uint32(v1088)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v111)+80)) = v1089
	*(*int32)(unsafe.Add(mBase, uint32(v111)+36)) = v1089
	v1092 = *(*int32)(unsafe.Add(mBase, uint32(v1083)))
	v1093 = *(*int32)(unsafe.Add(mBase, uint32(v1092)))
	*(*int32)(unsafe.Add(mBase, uint32(v111)+4)) = v1093
	v1095 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1089))))
	*(*uint8)(unsafe.Add(mBase, uint32(v111)+24)) = uint8(v1095)
	v1097 = v1059
	v1100 = v1054
	goto L227
L241:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1097)+24)) = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_replication_yylex[9])) = v1100
	v1117 = *(*int32)(unsafe.Add(mBase, uint32(v111)+20))
	v1118 = *(*int32)(unsafe.Add(mBase, uint32(v111)+12))
	v1121 = v1117 + v1118<<(uint(int32(2))%32)
	v1122 = *(*int32)(unsafe.Add(mBase, uint32(v1121)))
	v1123 = *(*int32)(unsafe.Add(mBase, uint32(v1122)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v111)+28)) = v1123
	v1125 = *(*int32)(unsafe.Add(mBase, uint32(v1121)))
	v1126 = *(*int32)(unsafe.Add(mBase, uint32(v1125)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v111)+36)) = v1126
	*(*int32)(unsafe.Add(mBase, uint32(v111)+80)) = v1126
	v1129 = *(*int32)(unsafe.Add(mBase, uint32(v1121)))
	v1130 = *(*int32)(unsafe.Add(mBase, uint32(v1129)))
	*(*int32)(unsafe.Add(mBase, uint32(v111)+4)) = v1130
	v1132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1126))))
	*(*uint8)(unsafe.Add(mBase, uint32(v111)+24)) = uint8(v1132)
	v1149 = int32(1)
	goto L222
L242:
	;
	v1105 = *(*int32)(unsafe.Add(mBase, uint32(v111)+12))
	v1109 = *(*int32)(unsafe.Add(mBase, uint32(v1104+v1105<<(uint(int32(2))%32))))
	if v1097 == v1109 {
		goto L241
	} else {
		goto L245
	}
L243:
	;
	goto L244
L244:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1097)+32)) = int64(1)
	goto L241
L245:
	;
	goto L244
L246:
	;
	v1162 = v1151 + v1150>>(uint(int32(1))%32)
	v1163 = *(*int32)(unsafe.Add(mBase, uint32(v1157)+4))
	if v1163 != 0 {
		goto L250
	} else {
		goto L251
	}
L247:
	;
	v1192 = v1151
	v1193 = v1152
	v1195 = v1153
	goto L248
L248:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v111)+28)) = v1192
	v1197 = int32(2)
	v1200 = *(*int32)(unsafe.Add(mBase, uint32(v1193+v1195<<(uint(v1197)%32))))
	v1201 = *(*int32)(unsafe.Add(mBase, uint32(v1200)+4))
	v1203 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1201+v1192))) = uint8(v1203)
	v1205 = *(*int32)(unsafe.Add(mBase, uint32(v111)+20))
	v1206 = *(*int32)(unsafe.Add(mBase, uint32(v111)+12))
	v1210 = *(*int32)(unsafe.Add(mBase, uint32(v1205+v1206<<(uint(v1197)%32))))
	v1211 = *(*int32)(unsafe.Add(mBase, uint32(v1210)+4))
	v1212 = *(*int32)(unsafe.Add(mBase, uint32(v111)+28))
	*(*uint8)(unsafe.Add(mBase, uint32(v1211+v1212)+1)) = uint8(v1203)
	v1216 = *(*int32)(unsafe.Add(mBase, uint32(v111)+20))
	v1217 = *(*int32)(unsafe.Add(mBase, uint32(v111)+12))
	v1220 = v1216 + v1217<<(uint(v1197)%32)
	v1221 = *(*int32)(unsafe.Add(mBase, uint32(v1220)))
	v1222 = *(*int32)(unsafe.Add(mBase, uint32(v1221)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v111)+80)) = v1222
	if v1149 == int32(1) {
		v1459 = v1222
		goto L48
	} else {
		goto L256
	}
L249:
	;
	v1169 = *(*int32)(unsafe.Add(mBase, uint32(v111)+20))
	v1170 = *(*int32)(unsafe.Add(mBase, uint32(v111)+12))
	v1171 = int32(2)
	v1174 = *(*int32)(unsafe.Add(mBase, uint32(v1169+v1170<<(uint(v1171)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v1174)+4)) = v1168
	v1176 = *(*int32)(unsafe.Add(mBase, uint32(v111)+20))
	v1177 = *(*int32)(unsafe.Add(mBase, uint32(v111)+12))
	v1181 = *(*int32)(unsafe.Add(mBase, uint32(v1176+v1177<<(uint(v1171)%32))))
	v1182 = *(*int32)(unsafe.Add(mBase, uint32(v1181)+4))
	if v1182 == int32(0) {
		goto L50
	} else {
		goto L255
	}
L250:
	;
	v1164 = F_repalloc(m, v1163, v1162)
	mBase = m.M
	v1165 = m.ExcPending
	if v1165 != 0 {
		goto L18
	} else {
		goto L253
	}
L251:
	;
	goto L252
L252:
	;
	v1166 = F_palloc(m, v1162)
	mBase = m.M
	v1167 = m.ExcPending
	if v1167 != 0 {
		goto L18
	} else {
		goto L254
	}
L253:
	;
	v1168 = v1164
	goto L249
L254:
	;
	v1168 = v1166
	goto L249
L255:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1181)+12)) = v1162 - int32(2)
	v1188 = *(*int32)(unsafe.Add(mBase, uint32(v111)+12))
	v1189 = *(*int32)(unsafe.Add(mBase, uint32(v111)+20))
	v1190 = *(*int32)(unsafe.Add(mBase, uint32(v111)+28))
	v1192 = v1190 + v651
	v1193 = v1189
	v1195 = v1188
	goto L248
L256:
	;
	switch v1149 - int32(1) {
	case 0:
		goto L51
	case 1:
		goto L257
	default:
		goto L258
	}
L257:
	;
	v1331 = *(*int32)(unsafe.Add(mBase, uint32(v111)+28))
	v1332 = *(*int32)(unsafe.Add(mBase, uint32(v1220)))
	v1333 = *(*int32)(unsafe.Add(mBase, uint32(v1332)+4))
	v1334 = v1333
	v1337 = v1222
	v1338 = v1331
	goto L52
L258:
	;
	v1230 = v406 ^ int32(-1) + v224
	v1231 = v1222 + v1230
	*(*int32)(unsafe.Add(mBase, uint32(v111)+36)) = v1231
	v1233 = *(*int32)(unsafe.Add(mBase, uint32(v111)+44))
	if v1230 <= int32(0) {
		v110 = v1233
		v113 = v1222
		v117 = v1231
		goto L27
	} else {
		goto L259
	}
L259:
	;
	v1236 = v1233
	v1240 = v1222
	goto L260
L260:
	;
	v1250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1240))))
	if v1250 != 0 {
		goto L262
	} else {
		goto L263
	}
L261:
	;
	v110 = v1327
	v113 = v1222
	v117 = v1231
	goto L27
L262:
	;
	v1251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1250)+uint32(_c_F_replication_yylex[2]))))
	v1252 = v1251
	goto L264
L263:
	;
	v1252 = int32(1)
	goto L264
L264:
	;
	v1254 = v1236 << (uint(int32(1)) % 32)
	v1257 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1254)+uint32(_c_F_replication_yylex[3]))))
	if v1257 != 0 {
		goto L265
	} else {
		goto L266
	}
L265:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v111)+68)) = v1240
	*(*int32)(unsafe.Add(mBase, uint32(v111)+64)) = v1236
	goto L267
L266:
	;
	goto L267
L267:
	;
	v1261 = v1252 & int32(255)
	v1264 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1254)+uint32(_c_F_replication_yylex[4]))))
	v1265 = v1261 + v1264
	v1270 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1265<<(uint(int32(1))%32))+uint32(_c_F_replication_yylex[5]))))
	if v1270 != v1236 {
		goto L268
	} else {
		goto L269
	}
L268:
	;
	v1272 = v1236
	v1274 = v1252
	v1277 = v1261
	goto L271
L269:
	;
	v1316 = v1265
	goto L270
L270:
	;
	v1323 = int32(1)
	v1327 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1316<<(uint(v1323)%32))+uint32(_c_F_replication_yylex[6]))))
	v1329 = v1240 + v1323
	if v1231 != v1329 {
		v1236 = v1327
		v1240 = v1329
		goto L260
	} else {
		goto L277
	}
L271:
	;
	v1289 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1272<<(uint(int32(1))%32))+uint32(_c_F_replication_yylex[7]))))
	if int32(286) <= v1289 {
		goto L273
	} else {
		goto L274
	}
L272:
	;
	v1316 = v1301
	goto L270
L273:
	;
	v1292 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1277)+uint32(_c_F_replication_yylex[8]))))
	v1293 = v1292
	goto L275
L274:
	;
	v1293 = v1274
	goto L275
L275:
	;
	v1295 = v1293 & int32(255)
	v1296 = int32(1)
	v1300 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1289<<(uint(v1296)%32))+uint32(_c_F_replication_yylex[4]))))
	v1301 = v1295 + v1300
	v1306 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1301<<(uint(v1296)%32))+uint32(_c_F_replication_yylex[5]))))
	if v1306 != v1289&int32(_a_F_replication_yylex_0) {
		v1272 = v1289
		v1274 = v1293
		v1277 = v1295
		goto L271
	} else {
		goto L276
	}
L276:
	;
	goto L272
L277:
	;
	goto L261
L278:
	;
	v1351 = v1349
	v1355 = v1337
	goto L279
L279:
	;
	v1365 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1355))))
	if v1365 != 0 {
		goto L281
	} else {
		goto L282
	}
L280:
	;
	v217 = v1444
	v224 = v1347
	v225 = v1337
	goto L44
L281:
	;
	v1366 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1365)+uint32(_c_F_replication_yylex[2]))))
	v1367 = v1366
	goto L283
L282:
	;
	v1367 = int32(1)
	goto L283
L283:
	;
	v1372 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1351<<(uint(int32(1))%32))+uint32(_c_F_replication_yylex[3]))))
	if v1372 != 0 {
		goto L284
	} else {
		goto L285
	}
L284:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v111)+68)) = v1355
	*(*int32)(unsafe.Add(mBase, uint32(v111)+64)) = v1351
	goto L286
L285:
	;
	goto L286
L286:
	;
	v1376 = v1367 & int32(255)
	v1377 = int32(1)
	v1381 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1351<<(uint(v1377)%32))+uint32(_c_F_replication_yylex[4]))))
	v1382 = v1376 + v1381
	v1387 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1382<<(uint(v1377)%32))+uint32(_c_F_replication_yylex[5]))))
	if v1387 != v1351 {
		goto L287
	} else {
		goto L288
	}
L287:
	;
	v1389 = v1351
	v1391 = v1367
	v1394 = v1376
	goto L290
L288:
	;
	v1433 = v1382
	goto L289
L289:
	;
	v1440 = int32(1)
	v1444 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1433<<(uint(v1440)%32))+uint32(_c_F_replication_yylex[6]))))
	v1446 = v1355 + v1440
	if v1446 != v1347 {
		v1351 = v1444
		v1355 = v1446
		goto L279
	} else {
		goto L296
	}
L290:
	;
	v1406 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1389<<(uint(int32(1))%32))+uint32(_c_F_replication_yylex[7]))))
	if int32(286) <= v1406 {
		goto L292
	} else {
		goto L293
	}
L291:
	;
	v1433 = v1418
	goto L289
L292:
	;
	v1409 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1394)+uint32(_c_F_replication_yylex[8]))))
	v1410 = v1409
	goto L294
L293:
	;
	v1410 = v1391
	goto L294
L294:
	;
	v1412 = v1410 & int32(255)
	v1413 = int32(1)
	v1417 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1406<<(uint(v1413)%32))+uint32(_c_F_replication_yylex[4]))))
	v1418 = v1412 + v1417
	v1423 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1418<<(uint(v1413)%32))+uint32(_c_F_replication_yylex[5]))))
	if v1423 != v1406&int32(_a_F_replication_yylex_0) {
		v1389 = v1406
		v1391 = v1410
		v1394 = v1412
		goto L290
	} else {
		goto L295
	}
L295:
	;
	goto L291
L296:
	;
	goto L280
L297:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L298:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
