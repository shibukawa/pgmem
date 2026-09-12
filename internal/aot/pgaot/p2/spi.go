package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_SPI_cursor_open(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	v2 = int32(0)
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v9 <= v2 {
		v12 = int32(0)
		v15 = F_SPI_cursor_open_internal(m, v12, l0, v12, int32(1))
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			return v15
		}
	} else {
		v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v21 = F_makeParamList(m, v9)
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return int32(0)
		} else {
			v26 = v2
			for {
				v35 = v21 + int32(32) + v26*int32(12)
				v37 = v26 << (uint(int32(2)) % 32)
				v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
				*(*int32)(unsafe.Add(mBase, uint32(v35))) = v38
				v40 = int32(1)
				*(*uint16)(unsafe.Add(mBase, uint32(v35)+6)) = uint16(v40)
				v42 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v35)+4)) = uint8(v42)
				v45 = *(*int32)(unsafe.Add(mBase, uint32(v20+v37)))
				*(*int32)(unsafe.Add(mBase, uint32(v35)+8)) = v45
				v48 = v26 + v40
				if v48 != v9 {
					v26 = v48
					continue
				} else {
					break
				}
				break
			}
			v52 = F_SPI_cursor_open_internal(m, int32(0), l0, v21, int32(1))
			mBase = m.M
			v53 = m.ExcPending
			if v53 != 0 {
				return int32(0)
			} else {
				F_pfree(m, v21)
				mBase = m.M
				v55 = m.ExcPending
				if v55 != 0 {
					return int32(0)
				} else {
					return v52
				}
			}
		}
	}
}
func F_SPI_cursor_open_internal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v51 int64
	_ = v51
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v130 int32
	_ = v130
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v226 int32
	_ = v226
	var v231 int32
	_ = v231
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v295 int32
	_ = v295
	var v299 int32
	_ = v299
	var v304 int32
	_ = v304
	var v308 int32
	_ = v308
	var v311 int32
	_ = v311
	var v315 int32
	_ = v315
	var v319 int32
	_ = v319
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v328 int32
	_ = v328
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v346 int32
	_ = v346
	var v349 int32
	_ = v349
	var v353 int32
	_ = v353
	var v358 int32
	_ = v358
	var v365 int32
	_ = v365
	var v368 int32
	_ = v368
	var v372 int32
	_ = v372
	var v377 int32
	_ = v377
	v11 = m.G0
	v13 = v11 - int32(48)
	m.G0 = v13
	if l1 != 0 {
		goto L6
	} else {
		goto L7
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L27
	} else {
		goto L100
	}
L2:
	;
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v325)+4))
	if v328 != int32(1) {
		goto L1
	} else {
		goto L91
	}
L3:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	v35 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[459])) = v35
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v33)+52))
	if v37 == v35 {
		v325 = v22
		goto L2
	} else {
		goto L15
	}
L4:
	;
	if v29 == int32(0) {
		goto L1
	} else {
		goto L14
	}
L5:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v22 != 0 {
		goto L10
	} else {
		goto L11
	}
L6:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v15 == int32(569278163) {
		goto L5
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	*(*int32)(unsafe.Add(mBase, _consts[459])) = int32(-6)
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v29 = v21
	goto L4
L9:
	;
	goto L8
L10:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	if v23 == int32(1) {
		goto L3
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	*(*int32)(unsafe.Add(mBase, _consts[459])) = int32(0)
	v29 = v22
	goto L4
L13:
	;
	goto L12
L14:
	;
	v325 = v29
	goto L2
L15:
	;
	v41 = *(*int32)(unsafe.Add(mBase, _consts[456]))
	if v41 != 0 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L27
	} else {
		goto L86
	}
L17:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	v44 = *(*int32)(unsafe.Add(mBase, _consts[61]))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)+8))
	goto L20
L18:
	;
	goto L19
L19:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L27
	} else {
		goto L83
	}
L20:
	;
	v47 = *(*int32)(unsafe.Add(mBase, _consts[456]))
	*(*int32)(unsafe.Add(mBase, uint32(v47)+12)) = v45
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v47)+24))
	v51 = int64(0)
	*(*int64)(unsafe.Add(mBase, _consts[457])) = v51
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v49
	v56 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[458])) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v47)+8)) = v56
	*(*int64)(unsafe.Add(mBase, uint32(v47))) = v51
	if l0 != 0 {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v71)+8))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v42)+12))
	v74 = F_MemoryContextStrdup(m, v72, v73)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L27
	} else {
		goto L30
	}
L22:
	;
	v67 = int32(0)
	v69 = F_CreatePortal(m, l0, v67, v67)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L27
	} else {
		goto L29
	}
L23:
	;
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v62 != 0 {
		goto L22
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	v63 = F_CreateNewPortal(m)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	goto L25
L27:
	;
	return int32(0)
L28:
	;
	v71 = v63
	goto L21
L29:
	;
	v71 = v69
	goto L21
L30:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v42)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+40)) = v76
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+32)) = int32(779)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+44)) = v78
	v82 = int32(4482056)
	v83 = *(*int32)(unsafe.Add(mBase, _consts[84]))
	*(*int32)(unsafe.Add(mBase, _consts[84])) = v13 + int32(28)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+28)) = v83
	*(*int32)(unsafe.Add(mBase, uint32(v13)+36)) = v13 + int32(40)
	v94 = *(*int32)(unsafe.Add(mBase, _consts[456]))
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v94)+36))
	v96 = F_GetCachedPlan(m, v42, l2, int32(0), v95)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L27
	} else {
		goto L31
	}
L31:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v96)+4))
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+4)))
	if v99 == int32(0) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v102 = int32(4489152)
	v103 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v71)+8))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v105
	v107 = F_copyObjectImpl(m, v98)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L27
	} else {
		goto L35
	}
L33:
	;
	v115 = v98
	v116 = v96
	goto L34
L34:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v42)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v71)+48)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v71)+40)) = v119
	*(*int32)(unsafe.Add(mBase, uint32(v71)+32)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v71)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v71)+80)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v71)+60)) = v116
	*(*int32)(unsafe.Add(mBase, uint32(v71)+56)) = v115
	*(*int32)(unsafe.Add(mBase, uint32(v71)+36)) = v119
	goto L37
L35:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v103
	F_ReleaseCachedPlan(m, v96, int32(0))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L27
	} else {
		goto L36
	}
L36:
	;
	v115 = v107
	v116 = int32(0)
	goto L34
L37:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v71)+76)) = v130
	if v130&int32(6) == int32(0) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	if v115 == int32(0) {
		v155 = v130
		goto L42
	} else {
		goto L43
	}
L39:
	;
	v163 = v130
	goto L40
L40:
	;
	if v115 == int32(0) {
		goto L49
	} else {
		goto L50
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v71)+76)) = v161
	v163 = v161
	goto L40
L42:
	;
	v161 = v155 | int32(4)
	goto L41
L43:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v115)+4))
	if v138 != int32(1) {
		v155 = v130
		goto L42
	} else {
		goto L44
	}
L44:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v115)+12))
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v141)))
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v142)+4))
	if v143 == int32(6) {
		v155 = v130
		goto L42
	} else {
		goto L45
	}
L45:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v142)+72))
	if v146 != 0 {
		v155 = v130
		goto L42
	} else {
		goto L46
	}
L46:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v142)+36))
	v148 = F_ExecSupportsBackwardScan(m, v147)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L27
	} else {
		goto L47
	}
L47:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v71)+76))
	if v148 == int32(0) {
		v155 = v150
		goto L42
	} else {
		goto L48
	}
L48:
	;
	v161 = v150 | int32(2)
	goto L41
L49:
	;
	v182 = *(*int32)(unsafe.Add(mBase, _consts[456]))
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v182)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v71)+68)) = v183
	if l3 != 0 {
		goto L56
	} else {
		goto L57
	}
L50:
	;
	if v163&int32(2) == int32(0) {
		goto L49
	} else {
		goto L51
	}
L51:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v115)+4))
	if v171 != int32(1) {
		goto L49
	} else {
		goto L52
	}
L52:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v115)+12))
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v174)))
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v175)+4))
	if v176 == int32(6) {
		goto L49
	} else {
		goto L53
	}
L53:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v175)+72))
	if v179 != 0 {
		goto L16
	} else {
		goto L54
	}
L54:
	;
	goto L49
L55:
	;
	if l2 != 0 {
		goto L77
	} else {
		goto L78
	}
L56:
	;
	if v115 == int32(0) {
		goto L59
	} else {
		goto L60
	}
L57:
	;
	goto L58
L58:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L27
	} else {
		goto L75
	}
L59:
	;
	v243 = *(*int32)(unsafe.Add(mBase, _consts[131]))
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v243)))
	goto L74
L60:
	;
	v187 = int32(0)
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v115)+4))
	if v188 <= v187 {
		goto L59
	} else {
		goto L61
	}
L61:
	;
	v192 = v187
	goto L62
L62:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v115)+12))
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v201+v192<<(uint(int32(2))%32))))
	v206 = F_CommandIsReadOnly(m, v205)
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L27
	} else {
		goto L64
	}
L63:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L27
	} else {
		goto L69
	}
L64:
	;
	if v206 != 0 {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v209 = v192 + int32(1)
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v115)+4))
	if v209 < v210 {
		v192 = v209
		goto L62
	} else {
		goto L68
	}
L66:
	;
	goto L67
L67:
	;
	goto L63
L68:
	;
	goto L59
L69:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L27
	} else {
		goto L70
	}
L70:
	;
	v219 = F_CreateCommandName(m, v205)
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L27
	} else {
		goto L71
	}
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v219
	F_errmsg(m, int32(253117), v13+int32(16))
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L27
	} else {
		goto L72
	}
L72:
	;
	F_errfinish(m, int32(495923), int32(1745), int32(310520))
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L27
	} else {
		goto L73
	}
L73:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L74:
	;
	v259 = v244
	goto L55
L75:
	;
	v247 = F_GetTransactionSnapshot(m)
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L27
	} else {
		goto L76
	}
L76:
	;
	v259 = v247
	goto L55
L77:
	;
	v261 = int32(4489152)
	v262 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v71)+8))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v264
	v266 = F_copyParamList(m, l2)
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L27
	} else {
		goto L80
	}
L78:
	;
	v270 = int32(0)
	goto L79
L79:
	;
	F_PortalStart(m, v71, v270, int32(0), v259)
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L27
	} else {
		goto L81
	}
L80:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v262
	v270 = v266
	goto L79
L81:
	;
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v13)+28))
	*(*int32)(unsafe.Add(mBase, _consts[84])) = v276
	v280 = *(*int32)(unsafe.Add(mBase, _consts[456]))
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v280)+20))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v281
	*(*int32)(unsafe.Add(mBase, uint32(v280)+12)) = int32(0)
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v280)+24))
	F_MemoryContextReset(m, v285)
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L27
	} else {
		goto L82
	}
L82:
	;
	m.G0 = v13 + int32(48)
	return v71
L83:
	;
	F_errmsg_internal(m, int32(446089), int32(0))
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L27
	} else {
		goto L84
	}
L84:
	;
	F_errfinish(m, int32(495923), int32(1620), int32(310520))
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L27
	} else {
		goto L85
	}
L85:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L86:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L27
	} else {
		goto L87
	}
L87:
	;
	F_errmsg(m, int32(441382), int32(0))
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L27
	} else {
		goto L88
	}
L88:
	;
	F_errdetail(m, int32(638978), int32(0))
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L27
	} else {
		goto L89
	}
L89:
	;
	F_errfinish(m, int32(495923), int32(1719), int32(310520))
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L27
	} else {
		goto L90
	}
L90:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L91:
	;
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v325)+12))
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v332)))
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v333)+16))
	if v334 != int32(179) {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v334<<(uint(int32(3))%32))+uint32(_consts[462])))
	goto L95
L93:
	;
	v342 = int32(525026)
	goto L94
L94:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L27
	} else {
		goto L96
	}
L95:
	;
	v342 = v341
	goto L94
L96:
	;
	F_errcode(m, int32(17170564))
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L27
	} else {
		goto L97
	}
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v342
	F_errmsg(m, int32(209547), v13)
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L27
	} else {
		goto L98
	}
L98:
	;
	F_errfinish(m, int32(495923), int32(1612), int32(310520))
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L27
	} else {
		goto L99
	}
L99:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L100:
	;
	F_errcode(m, int32(17170564))
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L27
	} else {
		goto L101
	}
L101:
	;
	F_errmsg(m, int32(209578), int32(0))
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L27
	} else {
		goto L102
	}
L102:
	;
	F_errfinish(m, int32(495923), int32(1602), int32(310520))
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L27
	} else {
		goto L103
	}
L103:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_SPI_exec(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v2 = int32(0)
	v4 = F_SPI_execute(m, l0, v2, v2)
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_SPI_finish(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int64
	_ = v28
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v51 int32
	_ = v51
	v3 = *(*int32)(unsafe.Add(mBase, _consts[456]))
	if v3 == int32(0) {
		return int32(-4)
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(v3)+28))
		*(*int32)(unsafe.Add(mBase, _consts[0])) = v9
		v11 = *(*int32)(unsafe.Add(mBase, uint32(v3)+24))
		F_MemoryContextDelete(m, v11)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, _consts[456]))
			*(*int32)(unsafe.Add(mBase, uint32(v17)+24)) = int32(0)
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v17)+20))
			F_MemoryContextDelete(m, v20)
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int32(0)
			} else {
				v23 = int32(4387760)
				v24 = *(*int32)(unsafe.Add(mBase, _consts[456]))
				v25 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v24)+20)) = v25
				v28 = *(*int64)(unsafe.Add(mBase, uint32(v24)+48))
				*(*int64)(unsafe.Add(mBase, _consts[457])) = v28
				v31 = *(*int32)(unsafe.Add(mBase, uint32(v24)+56))
				*(*int32)(unsafe.Add(mBase, _consts[458])) = v31
				v34 = *(*int32)(unsafe.Add(mBase, uint32(v24)+60))
				*(*int32)(unsafe.Add(mBase, _consts[459])) = v34
				v36 = int32(4103748)
				v38 = *(*int32)(unsafe.Add(mBase, _consts[460]))
				v40 = v38 - int32(1)
				*(*int32)(unsafe.Add(mBase, _consts[460])) = v40
				v44 = *(*int32)(unsafe.Add(mBase, _consts[461]))
				if v25 <= v40 {
					v51 = v44 + v40<<(uint(int32(6))%32)
				} else {
					v51 = v25
				}
				*(*int32)(unsafe.Add(mBase, _consts[456])) = v51
				return int32(2)
			}
		}
	}
}
func F_SPI_freeplan(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	v2 = int32(0)
	if l0 == v2 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v6 != int32(569278163) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v9 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	F_MemoryContextDelete(m, v32)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L9
	} else {
		goto L12
	}
L5:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	if v12 <= int32(0) {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v17 = v2
	goto L7
L7:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v18+v17<<(uint(int32(2))%32))))
	F_DropCachedPlan(m, v22)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	goto L4
L9:
	;
	return
L10:
	;
	v26 = v17 + int32(1)
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	if v26 < v27 {
		v17 = v26
		goto L7
	} else {
		goto L11
	}
L11:
	;
	goto L8
L12:
	;
	goto L1
}
func F_SPI_palloc(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	v4 = *(*int32)(unsafe.Add(mBase, _consts[456]))
	if v4 == int32(0) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			F_errmsg_internal(m, int32(532423), int32(0))
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(495923), int32(1341), int32(487129))
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	} else {
		v22 = *(*int32)(unsafe.Add(mBase, uint32(v4)+28))
		v23 = F_MemoryContextAlloc(m, v22, l0)
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return int32(0)
		} else {
			return v23
		}
	}
}
func F_SPI_result_code_string(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v9 = l0 + int32(13)
	if base.Ui32(int32(33)) <= base.Ui32(v9) {
		*(*int32)(unsafe.Add(mBase, uint32(v6))) = l0
		v26 = int32(4387776)
		v29 = F_pg_sprintf(m, v26, int32(476035), v6)
		mBase = m.M
		v32 = m.ExcPending
		if v32 != 0 {
			return int32(0)
		} else {
			v33 = v26
			m.G0 = v6 + int32(16)
			return v33
		}
	} else {
		if base.I32_wrap_i64(int64(base.Ui64(int64(8589926143))>>(uint(base.I64_extend_i32_u(v9))%64)))&int32(1) == int32(0) {
			*(*int32)(unsafe.Add(mBase, uint32(v6))) = l0
			v26 = int32(4387776)
			v29 = F_pg_sprintf(m, v26, int32(476035), v6)
			mBase = m.M
			v32 = m.ExcPending
			if v32 != 0 {
				return int32(0)
			} else {
				v33 = v26
				m.G0 = v6 + int32(16)
				return v33
			}
		} else {
			v24 = *(*int32)(unsafe.Add(mBase, uint32(v9<<(uint(int32(2))%32))+uint32(_consts[463])))
			v33 = v24
			m.G0 = v6 + int32(16)
			return v33
		}
	}
}
func F_SPI_sql_row_to_xmlelement(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	v5 = m.G0
	v7 = v5 - int32(48)
	m.G0 = v7
	if l1 != 0 {
		*(*int32)(unsafe.Add(mBase, uint32(v7)+32)) = int32(30851)
		F_appendStringInfo(m, l0, int32(176542), v7+int32(32))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return
		} else {
			F_appendStringInfoString(m, l0, int32(711945))
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return
			} else {
				v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
				if v19 != 0 {
					*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = l2
					F_appendStringInfo(m, l0, int32(674995), v7+int32(16))
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return
					} else {
						F_appendStringInfoString(m, l0, int32(736807))
						mBase = m.M
						v28 = m.ExcPending
						if v28 != 0 {
							return
						} else {
							v30 = *(*int32)(unsafe.Add(mBase, _consts[458]))
							v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
							v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
							if int32(0) < v32 {
								v59 = v31
								v61 = int32(0)
								*(*int32)(unsafe.Add(mBase, _consts[459])) = v61
								v63 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
								if v63 <= v61 {
									*(*int32)(unsafe.Add(mBase, _consts[459])) = int32(-9)
									v76 = F_map_sql_identifier_to_xml_name(m)
									mBase = m.M
									v77 = m.ExcPending
									if v77 != 0 {
										return
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								} else {
									v74 = F_pstrdup(m, v59+v63<<(uint(int32(4))%32)+int32(24))
									mBase = m.M
									v75 = m.ExcPending
									if v75 != 0 {
										return
									} else {
										v76 = F_map_sql_identifier_to_xml_name(m)
										mBase = m.M
										v77 = m.ExcPending
										if v77 != 0 {
											return
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(30851)
								F_appendStringInfo(m, l0, int32(735363), v7)
								mBase = m.M
								v39 = m.ExcPending
								if v39 != 0 {
									return
								} else {
									F_appendStringInfoChar(m, l0, int32(10))
									mBase = m.M
									v42 = m.ExcPending
									if v42 != 0 {
										return
									} else {
										m.G0 = v7 + int32(48)
										return
									}
								}
							}
						}
					}
				} else {
					F_appendStringInfoString(m, l0, int32(736807))
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return
					} else {
						v30 = *(*int32)(unsafe.Add(mBase, _consts[458]))
						v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
						v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
						if int32(0) < v32 {
							v59 = v31
							v61 = int32(0)
							*(*int32)(unsafe.Add(mBase, _consts[459])) = v61
							v63 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
							if v63 <= v61 {
								*(*int32)(unsafe.Add(mBase, _consts[459])) = int32(-9)
								v76 = F_map_sql_identifier_to_xml_name(m)
								mBase = m.M
								v77 = m.ExcPending
								if v77 != 0 {
									return
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							} else {
								v74 = F_pstrdup(m, v59+v63<<(uint(int32(4))%32)+int32(24))
								mBase = m.M
								v75 = m.ExcPending
								if v75 != 0 {
									return
								} else {
									v76 = F_map_sql_identifier_to_xml_name(m)
									mBase = m.M
									v77 = m.ExcPending
									if v77 != 0 {
										return
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(30851)
							F_appendStringInfo(m, l0, int32(735363), v7)
							mBase = m.M
							v39 = m.ExcPending
							if v39 != 0 {
								return
							} else {
								F_appendStringInfoChar(m, l0, int32(10))
								mBase = m.M
								v42 = m.ExcPending
								if v42 != 0 {
									return
								} else {
									m.G0 = v7 + int32(48)
									return
								}
							}
						}
					}
				}
			}
		}
	} else {
		F_appendStringInfoString(m, l0, int32(735336))
		mBase = m.M
		v45 = m.ExcPending
		if v45 != 0 {
			return
		} else {
			v47 = *(*int32)(unsafe.Add(mBase, _consts[458]))
			v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
			v49 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
			if int32(0) < v49 {
				v59 = v48
				v61 = int32(0)
				*(*int32)(unsafe.Add(mBase, _consts[459])) = v61
				v63 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
				if v63 <= v61 {
					*(*int32)(unsafe.Add(mBase, _consts[459])) = int32(-9)
					v76 = F_map_sql_identifier_to_xml_name(m)
					mBase = m.M
					v77 = m.ExcPending
					if v77 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				} else {
					v74 = F_pstrdup(m, v59+v63<<(uint(int32(4))%32)+int32(24))
					mBase = m.M
					v75 = m.ExcPending
					if v75 != 0 {
						return
					} else {
						v76 = F_map_sql_identifier_to_xml_name(m)
						mBase = m.M
						v77 = m.ExcPending
						if v77 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				F_appendStringInfoString(m, l0, int32(739718))
				mBase = m.M
				v54 = m.ExcPending
				if v54 != 0 {
					return
				} else {
					m.G0 = v7 + int32(48)
					return
				}
			}
		}
	}
}
func F__SPI_cursor_operation(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int64
	_ = v20
	var v25 int32
	_ = v25
	var v31 int64
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v42 int64
	_ = v42
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	if l0 != 0 {
		v9 = *(*int32)(unsafe.Add(mBase, _consts[456]))
		if v9 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v74 = m.ExcPending
			if v74 != 0 {
				return
			} else {
				F_errmsg_internal(m, int32(446041), int32(0))
				mBase = m.M
				v78 = m.ExcPending
				if v78 != 0 {
					return
				} else {
					F_errfinish(m, int32(495923), int32(3018), int32(258750))
					mBase = m.M
					v83 = m.ExcPending
					if v83 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, _consts[61]))
			v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
			v16 = *(*int32)(unsafe.Add(mBase, _consts[456]))
			*(*int32)(unsafe.Add(mBase, uint32(v16)+12)) = v14
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v16)+24))
			v20 = int64(0)
			*(*int64)(unsafe.Add(mBase, _consts[457])) = v20
			*(*int32)(unsafe.Add(mBase, _consts[0])) = v18
			v25 = int32(0)
			*(*int32)(unsafe.Add(mBase, _consts[458])) = v25
			*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = v25
			*(*int64)(unsafe.Add(mBase, uint32(v16))) = v20
			v31 = F_PortalRunFetch(m, l0, l1, l2, l3)
			mBase = m.M
			v32 = m.ExcPending
			if v32 != 0 {
				return
			} else {
				v34 = *(*int32)(unsafe.Add(mBase, _consts[456]))
				*(*int64)(unsafe.Add(mBase, uint32(v34))) = v31
				v36 = *(*int32)(unsafe.Add(mBase, uint32(v34)+8))
				v37 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
				if v37 == int32(5) {
					if v36 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v87 = m.ExcPending
						if v87 != 0 {
							return
						} else {
							F_errmsg_internal(m, int32(452719), int32(0))
							mBase = m.M
							v91 = m.ExcPending
							if v91 != 0 {
								return
							} else {
								F_errfinish(m, int32(495923), int32(3043), int32(258750))
								mBase = m.M
								v96 = m.ExcPending
								if v96 != 0 {
									return
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v42 = *(*int64)(unsafe.Add(mBase, uint32(v36)+8))
						if v31 != v42 {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v87 = m.ExcPending
							if v87 != 0 {
								return
							} else {
								F_errmsg_internal(m, int32(452719), int32(0))
								mBase = m.M
								v91 = m.ExcPending
								if v91 != 0 {
									return
								} else {
									F_errfinish(m, int32(495923), int32(3043), int32(258750))
									mBase = m.M
									v96 = m.ExcPending
									if v96 != 0 {
										return
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						} else {
							*(*int32)(unsafe.Add(mBase, _consts[458])) = v36
							*(*int64)(unsafe.Add(mBase, _consts[457])) = v31
							v48 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v34)+8)) = v48
							v51 = *(*int32)(unsafe.Add(mBase, uint32(v34)+20))
							*(*int32)(unsafe.Add(mBase, _consts[0])) = v51
							*(*int32)(unsafe.Add(mBase, uint32(v34)+12)) = v48
							v55 = *(*int32)(unsafe.Add(mBase, uint32(v34)+24))
							F_MemoryContextReset(m, v55)
							mBase = m.M
							v57 = m.ExcPending
							if v57 != 0 {
								return
							} else {
								return
							}
						}
					}
				} else {
					*(*int32)(unsafe.Add(mBase, _consts[458])) = v36
					*(*int64)(unsafe.Add(mBase, _consts[457])) = v31
					v48 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v34)+8)) = v48
					v51 = *(*int32)(unsafe.Add(mBase, uint32(v34)+20))
					*(*int32)(unsafe.Add(mBase, _consts[0])) = v51
					*(*int32)(unsafe.Add(mBase, uint32(v34)+12)) = v48
					v55 = *(*int32)(unsafe.Add(mBase, uint32(v34)+24))
					F_MemoryContextReset(m, v55)
					mBase = m.M
					v57 = m.ExcPending
					if v57 != 0 {
						return
					} else {
						return
					}
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v61 = m.ExcPending
		if v61 != 0 {
			return
		} else {
			F_errmsg_internal(m, int32(258938), int32(0))
			mBase = m.M
			v65 = m.ExcPending
			if v65 != 0 {
				return
			} else {
				F_errfinish(m, int32(495923), int32(3014), int32(258750))
				mBase = m.M
				v70 = m.ExcPending
				if v70 != 0 {
					return
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	}
}
