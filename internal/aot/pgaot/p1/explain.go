package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ExplainEndOutput(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	switch v2 - int32(1) {
	case 0:
		v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v5 - int32(1)
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		F_appendStringInfoString(m, v9, int32(574279))
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return
		} else {
			return
		}
	case 1:
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v13 - int32(1)
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		F_appendStringInfoString(m, v17, int32(535147))
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return
		} else {
			v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
			v22 = F_list_delete_first(m, v21)
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v22
				return
			}
		}
	case 2:
		v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v26 = F_list_delete_first(m, v25)
		mBase = m.M
		v27 = m.ExcPending
		if v27 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v26
			return
		}
	default:
		return
	}
}
func F_ExplainOneUtility(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v39 int32
	_ = v39
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
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
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v175 int32
	_ = v175
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v186 int64
	_ = v186
	var v187 int64
	_ = v187
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v225 int64
	_ = v225
	var v227 int64
	_ = v227
	var v233 int32
	_ = v233
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v255 int64
	_ = v255
	var v257 int64
	_ = v257
	var v258 int64
	_ = v258
	var v262 int64
	_ = v262
	var v264 int64
	_ = v264
	var v265 int64
	_ = v265
	var v269 int64
	_ = v269
	var v271 int64
	_ = v271
	var v272 int64
	_ = v272
	var v276 int64
	_ = v276
	var v278 int64
	_ = v278
	var v279 int64
	_ = v279
	var v283 int64
	_ = v283
	var v285 int64
	_ = v285
	var v286 int64
	_ = v286
	var v290 int64
	_ = v290
	var v292 int64
	_ = v292
	var v293 int64
	_ = v293
	var v297 int64
	_ = v297
	var v299 int64
	_ = v299
	var v300 int64
	_ = v300
	var v304 int64
	_ = v304
	var v306 int64
	_ = v306
	var v307 int64
	_ = v307
	var v311 int64
	_ = v311
	var v313 int64
	_ = v313
	var v314 int64
	_ = v314
	var v318 int64
	_ = v318
	var v320 int64
	_ = v320
	var v321 int64
	_ = v321
	var v325 int64
	_ = v325
	var v327 int64
	_ = v327
	var v328 int64
	_ = v328
	var v332 int64
	_ = v332
	var v334 int64
	_ = v334
	var v335 int64
	_ = v335
	var v339 int64
	_ = v339
	var v341 int64
	_ = v341
	var v342 int64
	_ = v342
	var v346 int64
	_ = v346
	var v348 int64
	_ = v348
	var v349 int64
	_ = v349
	var v353 int64
	_ = v353
	var v355 int64
	_ = v355
	var v356 int64
	_ = v356
	var v360 int64
	_ = v360
	var v362 int64
	_ = v362
	var v363 int64
	_ = v363
	var v367 int32
	_ = v367
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v381 int32
	_ = v381
	var v390 int32
	_ = v390
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v398 int32
	_ = v398
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v415 int32
	_ = v415
	var v417 int32
	_ = v417
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v427 int32
	_ = v427
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v449 int32
	_ = v449
	var v451 int32
	_ = v451
	var v453 int32
	_ = v453
	var v463 int32
	_ = v463
	var v466 int32
	_ = v466
	var v470 int32
	_ = v470
	var v475 int32
	_ = v475
	var v479 int32
	_ = v479
	var v483 int32
	_ = v483
	var v488 int32
	_ = v488
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v495 int32
	_ = v495
	var v498 int32
	_ = v498
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v505 int32
	_ = v505
	v6 = int32(0)
	v17 = m.G0
	v19 = v17 - int32(16)
	m.G0 = v19
	if l0 == v6 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v19 + int32(16)
	return
L2:
	;
	v23 = l0
	v24 = l1
	goto L3
L3:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	switch v39 - int32(242) {
	case 0:
		goto L10
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10:
		goto L6
	case 11:
		goto L8
	default:
		goto L11
	}
L4:
	;
	goto L1
L5:
	;
	v505 = *(*int32)(unsafe.Add(mBase, uint32(v503)+28))
	if v505 != 0 {
		v23 = v505
		v24 = v502
		goto L3
	} else {
		goto L137
	}
L6:
	;
	v492 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	if v492 == int32(0) {
		goto L132
	} else {
		goto L133
	}
L7:
	;
	F_ExplainDummyGroup(m, int32(21258), l2)
	mBase = m.M
	v491 = m.ExcPending
	if v491 != 0 {
		goto L15
	} else {
		goto L131
	}
L8:
	;
	v147 = m.G0
	v149 = v147 - int32(288)
	m.G0 = v149
	v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+11)))
	if v151 == int32(1) {
		goto L63
	} else {
		goto L64
	}
L9:
	;
	v111 = int32(0)
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
	v113 = F_copyObjectImpl(m, v112)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L15
	} else {
		goto L46
	}
L10:
	;
	v51 = F_CreateTableAsRelExists(m, v23)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L15
	} else {
		goto L17
	}
L11:
	;
	if v39 == int32(201) {
		goto L9
	} else {
		goto L12
	}
L12:
	;
	if v39 != int32(222) {
		goto L6
	} else {
		goto L13
	}
L13:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	if v46 != 0 {
		goto L7
	} else {
		goto L14
	}
L14:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	F_appendStringInfoString(m, v47, int32(789142))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	return
L16:
	;
	goto L1
L17:
	;
	if v51 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
	switch v53 - int32(23) {
	case 0:
		goto L22
	default:
		goto L21
	case 18:
		goto L23
	}
L19:
	;
	goto L20
L20:
	;
	v76 = int32(0)
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	v78 = F_copyObjectImpl(m, v77)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L15
	} else {
		goto L29
	}
L21:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L15
	} else {
		goto L26
	}
L22:
	;
	F_ExplainDummyGroup(m, int32(543804), l2)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L15
	} else {
		goto L25
	}
L23:
	;
	F_ExplainDummyGroup(m, int32(551641), l2)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L15
	} else {
		goto L24
	}
L24:
	;
	goto L1
L25:
	;
	goto L1
L26:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v66
	F_errmsg_internal(m, int32(509175), v19)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L15
	} else {
		goto L27
	}
L27:
	;
	F_errfinish(m, int32(522511), int32(419), int32(11632))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L15
	} else {
		goto L28
	}
L28:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L29:
	;
	v81 = *(*int32)(unsafe.Add(mBase, _consts[334]))
	switch v81 {
	case 0:
		v88 = v76
		goto L30
	case 1:
		goto L31
	default:
		goto L32
	}
L30:
	;
	v90 = *(*int32)(unsafe.Add(mBase, _consts[335]))
	if v90 != 0 {
		goto L35
	} else {
		goto L36
	}
L31:
	;
	v86 = F_JumbleQuery(m, v78)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L15
	} else {
		goto L34
	}
L32:
	;
	v83 = int32(*(*uint8)(unsafe.Add(mBase, _consts[336])))
	if v83 != int32(1) {
		v88 = v76
		goto L30
	} else {
		goto L33
	}
L33:
	;
	goto L31
L34:
	;
	v88 = v86
	goto L30
L35:
	;
	m.T0[v90].(func(*base.Module, int32, int32, int32))(m, l3, v78, v88)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L15
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	v93 = F_QueryRewrite(m, v78)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L15
	} else {
		goto L39
	}
L38:
	;
	goto L37
L39:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v93)+12))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v96)))
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v97)+4))
	if v98 == int32(6) {
		v502 = v95
		v503 = v97
		goto L5
	} else {
		goto L40
	}
L40:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(l3)+88))
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	v104 = *(*int32)(unsafe.Add(mBase, _consts[366]))
	if v104 != 0 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	m.T0[v104].(func(*base.Module, int32, int32, int32, int32, int32, int32, int32))(m, v97, int32(2048), v95, l2, v102, l4, v101)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L15
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	F_standard_ExplainOneQuery(m, v97, int32(2048), v95, l2, v102, l4, v101)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L15
	} else {
		goto L45
	}
L44:
	;
	goto L1
L45:
	;
	goto L1
L46:
	;
	v116 = *(*int32)(unsafe.Add(mBase, _consts[334]))
	switch v116 {
	case 0:
		v123 = v111
		goto L47
	case 1:
		goto L48
	default:
		goto L49
	}
L47:
	;
	v125 = *(*int32)(unsafe.Add(mBase, _consts[335]))
	if v125 != 0 {
		goto L52
	} else {
		goto L53
	}
L48:
	;
	v121 = F_JumbleQuery(m, v113)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L15
	} else {
		goto L51
	}
L49:
	;
	v118 = int32(*(*uint8)(unsafe.Add(mBase, _consts[336])))
	if v118 != int32(1) {
		v123 = v111
		goto L47
	} else {
		goto L50
	}
L50:
	;
	goto L48
L51:
	;
	v123 = v121
	goto L47
L52:
	;
	m.T0[v125].(func(*base.Module, int32, int32, int32))(m, l3, v113, v123)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L15
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	v129 = F_QueryRewrite(m, v113)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L15
	} else {
		goto L56
	}
L55:
	;
	goto L54
L56:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v129)+12))
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v131)))
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v132)+4))
	if v133 == int32(6) {
		v502 = int32(0)
		v503 = v132
		goto L5
	} else {
		goto L57
	}
L57:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	v137 = *(*int32)(unsafe.Add(mBase, uint32(l3)+88))
	v138 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	v140 = *(*int32)(unsafe.Add(mBase, _consts[366]))
	if v140 != 0 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	m.T0[v140].(func(*base.Module, int32, int32, int32, int32, int32, int32, int32))(m, v132, v136, int32(0), l2, v138, l4, v137)
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L15
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	F_standard_ExplainOneQuery(m, v132, v136, int32(0), l2, v138, l4, v137)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L15
	} else {
		goto L62
	}
L61:
	;
	goto L1
L62:
	;
	goto L1
L63:
	;
	v155 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	v160 = F_AllocSetContextCreateInternal(m, v155, int32(66423), int32(0), int32(8192), int32(8388608))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L15
	} else {
		goto L66
	}
L64:
	;
	v166 = v6
	v167 = v6
	goto L65
L65:
	;
	v168 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+7)))
	if v168 == int32(1) {
		goto L67
	} else {
		goto L68
	}
L66:
	;
	v162 = int32(4562096)
	v163 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v160
	v166 = v160
	v167 = v163
	goto L65
L67:
	;
	goto L71
L68:
	;
	goto L69
L69:
	;
	F___clock_gettime(m, int32(1), v149+int32(24))
	mBase = m.M
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	v183 = *(*int32)(unsafe.Add(mBase, _consts[367]))
	if v183 == int32(0) {
		goto L76
	} else {
		goto L77
	}
L70:
	;
	goto L69
L71:
	;
	v175 = F__emscripten_memcpy_bulkmem(m, v149+int32(152), int32(4460328), int32(128))
	mBase = m.M
	goto L73
L73:
	;
	goto L70
L74:
	;
	goto L1
L75:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v479 = m.ExcPending
	if v479 != 0 {
		goto L15
	} else {
		goto L128
	}
L76:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v463 = m.ExcPending
	if v463 != 0 {
		goto L15
	} else {
		goto L124
	}
L77:
	;
	v186 = int64(*(*int32)(unsafe.Add(mBase, uint32(v149)+32)))
	v187 = *(*int64)(unsafe.Add(mBase, uint32(v149)+24))
	v188 = int32(0)
	v190 = F_hash_search(m, v183, v181, v188, v188)
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L15
	} else {
		goto L78
	}
L78:
	;
	if v190 == int32(0) {
		goto L76
	} else {
		goto L79
	}
L79:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v190)+64))
	v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v194)+48)))
	if v195 == int32(0) {
		goto L75
	} else {
		goto L80
	}
L80:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v194)+12))
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v194)+24))
	if v199 != 0 {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	v201 = F_make_parsestate(m, int32(0))
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L15
	} else {
		goto L84
	}
L82:
	;
	v213 = v6
	v214 = v6
	v215 = v194
	goto L83
L83:
	;
	v217 = *(*int32)(unsafe.Add(mBase, _consts[175]))
	v218 = *(*int32)(unsafe.Add(mBase, uint32(l3)+88))
	v219 = F_GetCachedPlan(m, v215, v214, v217, v218)
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L15
	} else {
		goto L87
	}
L84:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v201)+4)) = v203
	v205 = F_CreateExecutorState(m)
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L15
	} else {
		goto L85
	}
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v205)+88)) = l4
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	v209 = F_EvaluateParams(m, v201, v190, v208, v205)
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L15
	} else {
		goto L86
	}
L86:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v190)+64))
	v213 = v205
	v214 = v209
	v215 = v211
	goto L83
L87:
	;
	v221 = int32(1)
	F___clock_gettime(m, v221, v149+int32(24))
	mBase = m.M
	v225 = int64(*(*int32)(unsafe.Add(mBase, uint32(v149)+32)))
	v227 = *(*int64)(unsafe.Add(mBase, uint32(v149)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v149)+280)) = v225 - v186 + (v227-v187)*int64(1000000000)
	v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+11)))
	if v233 == v221 {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v167
	F_MemoryContextMemConsumed(m, v166, v149+int32(8))
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L15
	} else {
		goto L91
	}
L89:
	;
	goto L90
L90:
	;
	v242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+7)))
	if v242 == int32(1) {
		goto L92
	} else {
		goto L93
	}
L91:
	;
	goto L90
L92:
	;
	v250 = F__emscripten_memset_bulkmem(m, v149+int32(24), base.I32_extend8_s(int32(0)), int32(128))
	mBase = m.M
	goto L95
L93:
	;
	goto L94
L94:
	;
	v367 = *(*int32)(unsafe.Add(mBase, uint32(v219)+4))
	if v367 == int32(0) {
		goto L97
	} else {
		goto L98
	}
L95:
	;
	v252 = v149 + int32(24)
	v254 = v149 + int32(152)
	v255 = *(*int64)(unsafe.Add(mBase, uint32(v252)))
	v257 = *(*int64)(unsafe.Add(mBase, _consts[277]))
	v258 = *(*int64)(unsafe.Add(mBase, uint32(v254)))
	*(*int64)(unsafe.Add(mBase, uint32(v252))) = v255 + (v257 - v258)
	v262 = *(*int64)(unsafe.Add(mBase, uint32(v252)+8))
	v264 = *(*int64)(unsafe.Add(mBase, _consts[368]))
	v265 = *(*int64)(unsafe.Add(mBase, uint32(v254)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v252)+8)) = v262 + (v264 - v265)
	v269 = *(*int64)(unsafe.Add(mBase, uint32(v252)+16))
	v271 = *(*int64)(unsafe.Add(mBase, _consts[369]))
	v272 = *(*int64)(unsafe.Add(mBase, uint32(v254)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v252)+16)) = v269 + (v271 - v272)
	v276 = *(*int64)(unsafe.Add(mBase, uint32(v252)+24))
	v278 = *(*int64)(unsafe.Add(mBase, _consts[370]))
	v279 = *(*int64)(unsafe.Add(mBase, uint32(v254)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v252)+24)) = v276 + (v278 - v279)
	v283 = *(*int64)(unsafe.Add(mBase, uint32(v252)+32))
	v285 = *(*int64)(unsafe.Add(mBase, _consts[266]))
	v286 = *(*int64)(unsafe.Add(mBase, uint32(v254)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v252)+32)) = v283 + (v285 - v286)
	v290 = *(*int64)(unsafe.Add(mBase, uint32(v252)+40))
	v292 = *(*int64)(unsafe.Add(mBase, _consts[371]))
	v293 = *(*int64)(unsafe.Add(mBase, uint32(v254)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v252)+40)) = v290 + (v292 - v293)
	v297 = *(*int64)(unsafe.Add(mBase, uint32(v252)+48))
	v299 = *(*int64)(unsafe.Add(mBase, _consts[372]))
	v300 = *(*int64)(unsafe.Add(mBase, uint32(v254)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v252)+48)) = v297 + (v299 - v300)
	v304 = *(*int64)(unsafe.Add(mBase, uint32(v252)+56))
	v306 = *(*int64)(unsafe.Add(mBase, _consts[373]))
	v307 = *(*int64)(unsafe.Add(mBase, uint32(v254)+56))
	*(*int64)(unsafe.Add(mBase, uint32(v252)+56)) = v304 + (v306 - v307)
	v311 = *(*int64)(unsafe.Add(mBase, uint32(v252)+64))
	v313 = *(*int64)(unsafe.Add(mBase, _consts[374]))
	v314 = *(*int64)(unsafe.Add(mBase, uint32(v254)+64))
	*(*int64)(unsafe.Add(mBase, uint32(v252)+64)) = v311 + (v313 - v314)
	v318 = *(*int64)(unsafe.Add(mBase, uint32(v252)+72))
	v320 = *(*int64)(unsafe.Add(mBase, _consts[375]))
	v321 = *(*int64)(unsafe.Add(mBase, uint32(v254)+72))
	*(*int64)(unsafe.Add(mBase, uint32(v252)+72)) = v318 + (v320 - v321)
	v325 = *(*int64)(unsafe.Add(mBase, uint32(v252)+80))
	v327 = *(*int64)(unsafe.Add(mBase, _consts[376]))
	v328 = *(*int64)(unsafe.Add(mBase, uint32(v254)+80))
	*(*int64)(unsafe.Add(mBase, uint32(v252)+80)) = v325 + (v327 - v328)
	v332 = *(*int64)(unsafe.Add(mBase, uint32(v252)+88))
	v334 = *(*int64)(unsafe.Add(mBase, _consts[377]))
	v335 = *(*int64)(unsafe.Add(mBase, uint32(v254)+88))
	*(*int64)(unsafe.Add(mBase, uint32(v252)+88)) = v332 + (v334 - v335)
	v339 = *(*int64)(unsafe.Add(mBase, uint32(v252)+96))
	v341 = *(*int64)(unsafe.Add(mBase, _consts[378]))
	v342 = *(*int64)(unsafe.Add(mBase, uint32(v254)+96))
	*(*int64)(unsafe.Add(mBase, uint32(v252)+96)) = v339 + (v341 - v342)
	v346 = *(*int64)(unsafe.Add(mBase, uint32(v252)+104))
	v348 = *(*int64)(unsafe.Add(mBase, _consts[379]))
	v349 = *(*int64)(unsafe.Add(mBase, uint32(v254)+104))
	*(*int64)(unsafe.Add(mBase, uint32(v252)+104)) = v346 + (v348 - v349)
	v353 = *(*int64)(unsafe.Add(mBase, uint32(v252)+112))
	v355 = *(*int64)(unsafe.Add(mBase, _consts[380]))
	v356 = *(*int64)(unsafe.Add(mBase, uint32(v254)+112))
	*(*int64)(unsafe.Add(mBase, uint32(v252)+112)) = v353 + (v355 - v356)
	v360 = *(*int64)(unsafe.Add(mBase, uint32(v252)+120))
	v362 = *(*int64)(unsafe.Add(mBase, _consts[381]))
	v363 = *(*int64)(unsafe.Add(mBase, uint32(v254)+120))
	*(*int64)(unsafe.Add(mBase, uint32(v252)+120)) = v360 + (v362 - v363)
	goto L96
L96:
	;
	goto L94
L97:
	;
	if v213 != 0 {
		goto L119
	} else {
		goto L120
	}
L98:
	;
	v370 = int32(0)
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v367)+4))
	if v371 <= v370 {
		goto L97
	} else {
		goto L99
	}
L99:
	;
	v381 = v370
	goto L100
L100:
	;
	v390 = *(*int32)(unsafe.Add(mBase, uint32(v367)+12))
	v393 = v390 + v381<<(uint(int32(2))%32)
	v394 = *(*int32)(unsafe.Add(mBase, uint32(v393)))
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v394)+4))
	if v395 != int32(6) {
		goto L103
	} else {
		goto L104
	}
L101:
	;
	goto L97
L102:
	;
	v417 = v393 + int32(4)
	if v417 == int32(0) {
		goto L114
	} else {
		goto L115
	}
L103:
	;
	v398 = *(*int32)(unsafe.Add(mBase, uint32(l3)+88))
	v404 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+7)))
	if v404 != 0 {
		goto L106
	} else {
		goto L107
	}
L104:
	;
	goto L105
L105:
	;
	v413 = *(*int32)(unsafe.Add(mBase, uint32(v394)+88))
	F_ExplainOneUtility(m, v413, v24, l2, l3, v214)
	mBase = m.M
	v415 = m.ExcPending
	if v415 != 0 {
		goto L15
	} else {
		goto L113
	}
L106:
	;
	v405 = v149 + int32(24)
	goto L108
L107:
	;
	v405 = int32(0)
	goto L108
L108:
	;
	v409 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+11)))
	if v409 != 0 {
		goto L109
	} else {
		goto L110
	}
L109:
	;
	v410 = v149 + int32(8)
	goto L111
L110:
	;
	v410 = int32(0)
	goto L111
L111:
	;
	F_ExplainOnePlan(m, v394, v24, l2, v198, v214, v398, v149+int32(280), v405, v410)
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L15
	} else {
		goto L112
	}
L112:
	;
	goto L102
L113:
	;
	goto L102
L114:
	;
	v429 = v381 + int32(1)
	v430 = *(*int32)(unsafe.Add(mBase, uint32(v367)+4))
	if v429 < v430 {
		v381 = v429
		goto L100
	} else {
		goto L118
	}
L115:
	;
	v420 = *(*int32)(unsafe.Add(mBase, uint32(v367)+12))
	v421 = *(*int32)(unsafe.Add(mBase, uint32(v367)+4))
	if base.Ui32(v420+v421<<(uint(int32(2))%32)) <= base.Ui32(v417) {
		goto L114
	} else {
		goto L116
	}
L116:
	;
	F_ExplainSeparatePlans(m, l2)
	mBase = m.M
	v427 = m.ExcPending
	if v427 != 0 {
		goto L15
	} else {
		goto L117
	}
L117:
	;
	goto L114
L118:
	;
	goto L101
L119:
	;
	F_FreeExecutorState(m, v213)
	mBase = m.M
	v449 = m.ExcPending
	if v449 != 0 {
		goto L15
	} else {
		goto L122
	}
L120:
	;
	goto L121
L121:
	;
	v451 = *(*int32)(unsafe.Add(mBase, _consts[175]))
	F_ReleaseCachedPlan(m, v219, v451)
	mBase = m.M
	v453 = m.ExcPending
	if v453 != 0 {
		goto L15
	} else {
		goto L123
	}
L122:
	;
	goto L121
L123:
	;
	m.G0 = v149 + int32(288)
	goto L74
L124:
	;
	F_errcode(m, int32(386))
	mBase = m.M
	v466 = m.ExcPending
	if v466 != 0 {
		goto L15
	} else {
		goto L125
	}
L125:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v149))) = v181
	F_errmsg(m, int32(76859), v149)
	mBase = m.M
	v470 = m.ExcPending
	if v470 != 0 {
		goto L15
	} else {
		goto L126
	}
L126:
	;
	F_errfinish(m, int32(525081), int32(454), int32(102558))
	mBase = m.M
	v475 = m.ExcPending
	if v475 != 0 {
		goto L15
	} else {
		goto L127
	}
L127:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L128:
	;
	F_errmsg_internal(m, int32(159443), int32(0))
	mBase = m.M
	v483 = m.ExcPending
	if v483 != 0 {
		goto L15
	} else {
		goto L129
	}
L129:
	;
	F_errfinish(m, int32(525081), int32(608), int32(17384))
	mBase = m.M
	v488 = m.ExcPending
	if v488 != 0 {
		goto L15
	} else {
		goto L130
	}
L130:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L131:
	;
	goto L1
L132:
	;
	v495 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	F_appendStringInfoString(m, v495, int32(786921))
	mBase = m.M
	v498 = m.ExcPending
	if v498 != 0 {
		goto L15
	} else {
		goto L135
	}
L133:
	;
	goto L134
L134:
	;
	F_ExplainDummyGroup(m, int32(102604), l2)
	mBase = m.M
	v501 = m.ExcPending
	if v501 != 0 {
		goto L15
	} else {
		goto L136
	}
L135:
	;
	goto L1
L136:
	;
	goto L1
L137:
	;
	goto L4
}
func F_ExplainPrintPlan(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v86 int32
	_ = v86
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v123 int32
	_ = v123
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v286 int32
	_ = v286
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v304 int32
	_ = v304
	var v310 int32
	_ = v310
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v332 float64
	_ = v332
	var v333 int32
	_ = v333
	var v334 float64
	_ = v334
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v351 int32
	_ = v351
	var v356 int32
	_ = v356
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v374 int32
	_ = v374
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v388 int32
	_ = v388
	var v393 int32
	_ = v393
	var v400 int32
	_ = v400
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v429 int32
	_ = v429
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v449 int32
	_ = v449
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v494 int32
	_ = v494
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v509 int32
	_ = v509
	var v521 int32
	_ = v521
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v541 int32
	_ = v541
	var v547 int32
	_ = v547
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v564 int32
	_ = v564
	var v566 int32
	_ = v566
	var v578 int32
	_ = v578
	var v581 int32
	_ = v581
	var v582 int64
	_ = v582
	var v586 int32
	_ = v586
	var v592 int32
	_ = v592
	v12 = m.G0
	v14 = v12 - int32(80)
	m.G0 = v14
	*(*int32)(unsafe.Add(mBase, uint32(v14)+56)) = int32(0)
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v18
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v18)+44))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v20
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
	v25 = F_ExplainPreScanNode(m, v22, v14+int32(56))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v14)+56))
	v29 = m.G0
	v31 = v29 - int32(80)
	m.G0 = v31
	v38 = F__emscripten_memset_bulkmem(m, v31+int32(4), base.I32_extend8_s(int32(0)), int32(76))
	mBase = m.M
	goto L3
L3:
	;
	v39 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+16)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v31))) = v27
	F_set_rtable_names(m, v31, v39, v28)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	v46 = int32(80)
	m.G0 = v31 + v46
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v45
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v51 = m.G0
	v53 = v51 - int32(16)
	m.G0 = v53
	v56 = F_palloc0(m, v46)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v50)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v56)+4)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v56))) = v58
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v50)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v56)+16)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v56)+12)) = v61
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v50)+60))
	if v65 != 0 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	F_set_simple_column_names(m, v56)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L1
	} else {
		goto L19
	}
L7:
	;
	if v58 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	goto L9
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v56)+20)) = int32(0)
	goto L6
L10:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v58)+4))
	v72 = v66<<(uint(int32(2))%32) + int32(4)
	goto L12
L11:
	;
	v72 = int32(4)
	goto L12
L12:
	;
	v73 = F_palloc0(m, v72)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v56)+20)) = v73
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v50)+60))
	if v76 == int32(0) {
		goto L6
	} else {
		goto L14
	}
L14:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v76)+4))
	if v79 <= int32(0) {
		goto L6
	} else {
		goto L15
	}
L15:
	;
	v86 = int32(0)
	goto L16
L16:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v56)+20))
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v76)+12))
	v96 = int32(2)
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v95+v86<<(uint(v96)%32))))
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v99)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v94+v100<<(uint(v96)%32)))) = v99
	v106 = v86 + int32(1)
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v76)+4))
	if v106 < v107 {
		v86 = v106
		goto L16
	} else {
		goto L18
	}
L17:
	;
	goto L6
L18:
	;
	goto L17
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v53)+8)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v53)+12)) = v56
	v129 = F_list_make1_impl(m, int32(1), v53+int32(8))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	m.G0 = v53 + int32(16)
	v134 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = v134
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v129
	v137 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v137 == v134 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v188)))
	if v189 != int32(432) {
		v199 = v188
		goto L35
	} else {
		goto L36
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = int32(0)
	goto L21
L23:
	;
	goto L24
L24:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v137)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = v142
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v137)+4))
	if v144 <= int32(0) {
		goto L21
	} else {
		goto L25
	}
L25:
	;
	v147 = int32(0)
	if v147 < v144 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v150 = v144
	goto L28
L27:
	;
	v150 = v147
	goto L28
L28:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v137)+12))
	v155 = int32(0)
	goto L29
L29:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v151+v155<<(uint(int32(2))%32))))
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v167)+12))
	if v168 != int32(9) {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = v142 - int32(1)
	goto L21
L31:
	;
	v172 = v155 + int32(1)
	if v150 != v172 {
		v155 = v172
		goto L29
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	goto L30
L34:
	;
	goto L21
L35:
	;
	v200 = int32(0)
	F_ExplainNode(m, v199, v200, v200, v200, l0)
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L1
	} else {
		goto L38
	}
L36:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v188)+4))
	v193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192)+81)))
	if v193 != int32(1) {
		v199 = v188
		goto L35
	} else {
		goto L37
	}
L37:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v188)+36))
	v197 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+52)) = uint8(v197)
	v199 = v196
	goto L35
L38:
	;
	v205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)))
	if v205 != int32(1) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v578 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)))
	if v578 != int32(1) {
		goto L120
	} else {
		goto L121
	}
L40:
	;
	v208 = m.G0
	v210 = v208 - int32(16)
	m.G0 = v210
	v213 = v14 + int32(76)
	*(*int32)(unsafe.Add(mBase, uint32(v213))) = int32(0)
	v217 = *(*int32)(unsafe.Add(mBase, _consts[382]))
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v217)))
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v219)+4))
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v219)+412))
	if v221 != 0 {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	v289 = F_palloc(m, v286<<(uint(int32(2))%32))
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L1
	} else {
		goto L45
	}
L42:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v219)+376))
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v219)+364))
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v219)+352))
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v219)+340))
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v219)+328))
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v219)+316))
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v219)+304))
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v219)+292))
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v219)+280))
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v219)+268))
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v219)+256))
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v219)+244))
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v219)+232))
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v219)+220))
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v219)+208))
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v219)+196))
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v219)+184))
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v219)+172))
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v219)+160))
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v219)+148))
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v219)+136))
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v219)+124))
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v219)+112))
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v219)+100))
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v219)+88))
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v219)+76))
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v219-int32(-64))))
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v219)+52))
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v219)+40))
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v219)+28))
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v219)+16))
	v286 = v222 + (v223 + (v224 + (v225 + (v226 + (v227 + (v228 + (v229 + (v230 + (v231 + (v232 + (v233 + (v234 + (v235 + (v236 + (v237 + (v238 + (v239 + (v240 + (v241 + (v242 + (v243 + (v244 + (v245 + (v246 + (v247 + (v250 + (v251 + (v252 + (v253 + (v254 + v220))))))))))))))))))))))))))))))
	goto L44
L43:
	;
	v286 = v220
	goto L44
L44:
	;
	goto L41
L45:
	;
	v292 = *(*int32)(unsafe.Add(mBase, _consts[383]))
	if v292 == int32(0) {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	m.G0 = v210 + int32(16)
	v417 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v417 != 0 {
		goto L83
	} else {
		goto L84
	}
L47:
	;
	if v292 == int32(4560124) {
		goto L46
	} else {
		goto L48
	}
L48:
	;
	v304 = v292
	goto L49
L49:
	;
	v310 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v304-int32(44)))))
	if v310&int32(32) == int32(0) {
		goto L51
	} else {
		goto L52
	}
L50:
	;
	goto L46
L51:
	;
	v400 = *(*int32)(unsafe.Add(mBase, uint32(v304)+4))
	if v400 != int32(4560124) {
		v304 = v400
		goto L49
	} else {
		goto L82
	}
L52:
	;
	v316 = v304 + int32(-64)
	v317 = F_ConfigOptionIsVisible(m, v316)
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	if v317 == int32(0) {
		goto L51
	} else {
		goto L54
	}
L54:
	;
	v322 = v304 - int32(40)
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v322)))
	switch v323 {
	case 0:
		goto L62
	case 1:
		goto L61
	case 2:
		goto L60
	case 3:
		goto L59
	case 4:
		goto L58
	default:
		goto L57
	}
L55:
	;
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v213)))
	*(*int32)(unsafe.Add(mBase, uint32(v289+v388<<(uint(int32(2))%32)))) = v316
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v213)))
	*(*int32)(unsafe.Add(mBase, uint32(v213))) = v393 + int32(1)
	goto L51
L56:
	;
	if v337 == int32(0) {
		goto L55
	} else {
		goto L72
	}
L57:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L1
	} else {
		goto L69
	}
L58:
	;
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v304)+32))
	v340 = *(*int32)(unsafe.Add(mBase, uint32(v304)+28))
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v340)))
	if v339 != v341 {
		goto L55
	} else {
		goto L68
	}
L59:
	;
	v336 = *(*int32)(unsafe.Add(mBase, uint32(v304)+28))
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v336)))
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v304)+32))
	if v338 != 0 {
		goto L56
	} else {
		goto L66
	}
L60:
	;
	v332 = *(*float64)(unsafe.Add(mBase, uint32(v304)+32))
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v304)+28))
	v334 = *(*float64)(unsafe.Add(mBase, uint32(v333)))
	if base.F64_ne(v332, v334) != 0 {
		goto L55
	} else {
		goto L65
	}
L61:
	;
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v304)+32))
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v304)+28))
	v330 = *(*int32)(unsafe.Add(mBase, uint32(v329)))
	if v328 != v330 {
		goto L55
	} else {
		goto L64
	}
L62:
	;
	v324 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v304)+32)))
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v304)+28))
	v326 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v325))))
	if v324 != v326 {
		goto L55
	} else {
		goto L63
	}
L63:
	;
	goto L51
L64:
	;
	goto L51
L65:
	;
	goto L51
L66:
	;
	if v337 != 0 {
		goto L55
	} else {
		goto L67
	}
L67:
	;
	goto L51
L68:
	;
	goto L51
L69:
	;
	v347 = *(*int32)(unsafe.Add(mBase, uint32(v322)))
	*(*int32)(unsafe.Add(mBase, uint32(v210))) = v347
	F_errmsg_internal(m, int32(510524), v210)
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L1
	} else {
		goto L70
	}
L70:
	;
	F_errfinish(m, int32(526378), int32(5420), int32(146603))
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
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
	v361 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v337))))
	v362 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v338))))
	if v362 == int32(0) {
		v381 = v361
		v382 = v362
		goto L74
	} else {
		goto L75
	}
L73:
	;
	if v382-v381 == int32(0) {
		goto L51
	} else {
		goto L81
	}
L74:
	;
	goto L73
L75:
	;
	if v361 != v362 {
		v381 = v361
		v382 = v362
		goto L74
	} else {
		goto L76
	}
L76:
	;
	v366 = v338
	v367 = v337
	goto L77
L77:
	;
	v370 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v367)+1)))
	v371 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v366)+1)))
	if v371 == int32(0) {
		v381 = v370
		v382 = v371
		goto L74
	} else {
		goto L79
	}
L78:
	;
	v381 = v370
	v382 = v371
	goto L74
L79:
	;
	v374 = int32(1)
	if v370 == v371 {
		v366 = v366 + v374
		v367 = v367 + v374
		goto L77
	} else {
		goto L80
	}
L80:
	;
	goto L78
L81:
	;
	goto L55
L82:
	;
	goto L50
L83:
	;
	v418 = int32(166420)
	F_ExplainOpenGroup(m, v418, v418, int32(1), l0)
	mBase = m.M
	v422 = m.ExcPending
	if v422 != 0 {
		goto L1
	} else {
		goto L86
	}
L84:
	;
	goto L85
L85:
	;
	v469 = *(*int32)(unsafe.Add(mBase, uint32(v14)+76))
	if v469 <= int32(0) {
		goto L39
	} else {
		goto L96
	}
L86:
	;
	v423 = int32(0)
	v424 = *(*int32)(unsafe.Add(mBase, uint32(v14)+76))
	if v423 < v424 {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v429 = v423
	goto L90
L88:
	;
	goto L89
L89:
	;
	F_ExplainCloseGroup(m, int32(166420), int32(1), l0)
	mBase = m.M
	v468 = m.ExcPending
	if v468 != 0 {
		goto L1
	} else {
		goto L95
	}
L90:
	;
	v441 = *(*int32)(unsafe.Add(mBase, uint32(v289+v429<<(uint(int32(2))%32))))
	v442 = *(*int32)(unsafe.Add(mBase, uint32(v441)))
	v445 = F_GetConfigOptionByName(m, v442, int32(0), int32(1))
	mBase = m.M
	v446 = m.ExcPending
	if v446 != 0 {
		goto L1
	} else {
		goto L92
	}
L91:
	;
	goto L89
L92:
	;
	v447 = *(*int32)(unsafe.Add(mBase, uint32(v441)))
	F_ExplainPropertyText(m, v447, v445, l0)
	mBase = m.M
	v449 = m.ExcPending
	if v449 != 0 {
		goto L1
	} else {
		goto L93
	}
L93:
	;
	v451 = v429 + int32(1)
	v452 = *(*int32)(unsafe.Add(mBase, uint32(v14)+76))
	if v451 < v452 {
		v429 = v451
		goto L90
	} else {
		goto L94
	}
L94:
	;
	goto L91
L95:
	;
	goto L39
L96:
	;
	F_initStringInfo(m, v14+int32(60))
	mBase = m.M
	v475 = m.ExcPending
	if v475 != 0 {
		goto L1
	} else {
		goto L97
	}
L97:
	;
	v476 = *(*int32)(unsafe.Add(mBase, uint32(v14)+76))
	if v476 <= int32(0) {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v564 = *(*int32)(unsafe.Add(mBase, uint32(v14)+60))
	F_ExplainPropertyText(m, int32(166420), v564, l0)
	mBase = m.M
	v566 = m.ExcPending
	if v566 != 0 {
		goto L1
	} else {
		goto L119
	}
L99:
	;
	v479 = *(*int32)(unsafe.Add(mBase, uint32(v289)))
	v480 = *(*int32)(unsafe.Add(mBase, uint32(v479)))
	v483 = F_GetConfigOptionByName(m, v480, int32(0), int32(1))
	mBase = m.M
	v484 = m.ExcPending
	if v484 != 0 {
		goto L1
	} else {
		goto L100
	}
L100:
	;
	v485 = *(*int32)(unsafe.Add(mBase, uint32(v479)))
	if v483 != 0 {
		goto L102
	} else {
		goto L103
	}
L101:
	;
	v503 = int32(1)
	v504 = *(*int32)(unsafe.Add(mBase, uint32(v14)+76))
	if v504 <= v503 {
		goto L98
	} else {
		goto L107
	}
L102:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+52)) = v483
	*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = v485
	F_appendStringInfo(m, v14+int32(60), int32(721353), v14+int32(48))
	mBase = m.M
	v494 = m.ExcPending
	if v494 != 0 {
		goto L1
	} else {
		goto L105
	}
L103:
	;
	goto L104
L104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v485
	F_appendStringInfo(m, v14+int32(60), int32(560618), v14+int32(32))
	mBase = m.M
	v502 = m.ExcPending
	if v502 != 0 {
		goto L1
	} else {
		goto L106
	}
L105:
	;
	goto L101
L106:
	;
	goto L101
L107:
	;
	v509 = v503
	goto L108
L108:
	;
	v521 = *(*int32)(unsafe.Add(mBase, uint32(v289+v509<<(uint(int32(2))%32))))
	F_appendStringInfoString(m, v14+int32(60), int32(783295))
	mBase = m.M
	v526 = m.ExcPending
	if v526 != 0 {
		goto L1
	} else {
		goto L110
	}
L109:
	;
	goto L98
L110:
	;
	v527 = *(*int32)(unsafe.Add(mBase, uint32(v521)))
	v530 = F_GetConfigOptionByName(m, v527, int32(0), int32(1))
	mBase = m.M
	v531 = m.ExcPending
	if v531 != 0 {
		goto L1
	} else {
		goto L111
	}
L111:
	;
	v532 = *(*int32)(unsafe.Add(mBase, uint32(v521)))
	if v530 != 0 {
		goto L113
	} else {
		goto L114
	}
L112:
	;
	v549 = v509 + int32(1)
	v550 = *(*int32)(unsafe.Add(mBase, uint32(v14)+76))
	if v549 < v550 {
		v509 = v549
		goto L108
	} else {
		goto L118
	}
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = v530
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v532
	F_appendStringInfo(m, v14+int32(60), int32(721353), v14+int32(16))
	mBase = m.M
	v541 = m.ExcPending
	if v541 != 0 {
		goto L1
	} else {
		goto L116
	}
L114:
	;
	goto L115
L115:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v532
	F_appendStringInfo(m, v14+int32(60), int32(560618), v14)
	mBase = m.M
	v547 = m.ExcPending
	if v547 != 0 {
		goto L1
	} else {
		goto L117
	}
L116:
	;
	goto L112
L117:
	;
	goto L112
L118:
	;
	goto L109
L119:
	;
	goto L39
L120:
	;
	m.G0 = v14 + int32(80)
	return
L121:
	;
	v581 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v582 = *(*int64)(unsafe.Add(mBase, uint32(v581)+8))
	if v582 == int64(0) {
		goto L120
	} else {
		goto L122
	}
L122:
	;
	v586 = *(*int32)(unsafe.Add(mBase, _consts[334]))
	if v586 == int32(3) {
		goto L120
	} else {
		goto L123
	}
L123:
	;
	F_ExplainPropertyInteger(m, int32(234365), int32(0), v582, l0)
	mBase = m.M
	v592 = m.ExcPending
	if v592 != 0 {
		goto L1
	} else {
		goto L124
	}
L124:
	;
	goto L120
}
func F_ExplainPrintTriggers(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v128 int32
	_ = v128
	v3 = int32(0)
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+84))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v8)+80))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v8)+72))
	v12 = int32(143977)
	F_ExplainOpenGroup(m, v12, v12, v3, l0)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	if v11 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	if v9 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L4:
	;
	v65 = int32(0)
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	if v66 <= v65 {
		v88 = v59
		goto L3
	} else {
		goto L17
	}
L5:
	;
	if v10 != 0 {
		v59 = int32(1)
		goto L4
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	v23 = int32(0)
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	v28 = base.B2i32(v9|v10 != v23) | base.B2i32(int32(1) < v25)
	if v23 < v25 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v88 = base.B2i32(v9 != int32(0))
	goto L3
L9:
	;
	v33 = v3
	goto L12
L10:
	;
	goto L11
L11:
	;
	if v10 == int32(0) {
		v88 = v28
		goto L3
	} else {
		goto L16
	}
L12:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v38+v33<<(uint(int32(2))%32))))
	F_report_triggers(m, v42, v28, l0)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L1
	} else {
		goto L14
	}
L13:
	;
	goto L11
L14:
	;
	v46 = v33 + int32(1)
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	if v46 < v47 {
		v33 = v46
		goto L12
	} else {
		goto L15
	}
L15:
	;
	goto L13
L16:
	;
	v59 = v28
	goto L4
L17:
	;
	v71 = v65
	goto L18
L18:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v76+v71<<(uint(int32(2))%32))))
	F_report_triggers(m, v80, v59, l0)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L1
	} else {
		goto L20
	}
L19:
	;
	v88 = v59
	goto L3
L20:
	;
	v84 = v71 + int32(1)
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	if v84 < v85 {
		v71 = v84
		goto L18
	} else {
		goto L21
	}
L21:
	;
	goto L19
L22:
	;
	F_ExplainCloseGroup(m, int32(143977), int32(0), l0)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L1
	} else {
		goto L29
	}
L23:
	;
	v96 = int32(0)
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	if v97 <= v96 {
		goto L22
	} else {
		goto L24
	}
L24:
	;
	v102 = v96
	goto L25
L25:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v107+v102<<(uint(int32(2))%32))))
	F_report_triggers(m, v111, v88, l0)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L1
	} else {
		goto L27
	}
L26:
	;
	goto L22
L27:
	;
	v115 = v102 + int32(1)
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	if v115 < v116 {
		v102 = v115
		goto L25
	} else {
		goto L28
	}
L28:
	;
	goto L26
L29:
	;
	return
}
func F_ExplainPropertyText(m *base.Module, l0 int32, l1 int32, l2 int32) {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = int32(0)
	F_ExplainProperty(m, l0, v4, l1, v4, l2)
	v7 = m.ExcPending
	if v7 != 0 {
		return
	} else {
		return
	}
}
func F_ExplainTargetRel(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
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
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
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
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v141 int32
	_ = v141
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	v4 = int32(0)
	v10 = m.G0
	v12 = v10 + int32(-64)
	m.G0 = v12
	v17 = l1<<(uint(int32(2))%32) - int32(4)
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l2)+36))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v17+v19)))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l2)+40))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v23+v17)))
	if v25 == v4 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	v30 = v29
	goto L3
L2:
	;
	v30 = v25
	goto L3
L3:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch v32 - int32(333) {
	case 0, 6, 7, 8, 9, 11, 12, 13, 21, 22:
		goto L10
	default:
		v105 = int32(0)
		v106 = v4
		v108 = v4
		v109 = v4
		goto L4
	case 15:
		goto L9
	case 17:
		goto L8
	case 18, 20:
		goto L6
	case 19:
		goto L7
	}
L4:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	if v110 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L5:
	;
	v105 = v103
	v106 = v4
	v108 = int32(1)
	v109 = v102
	goto L4
L6:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v21)+84))
	v102 = int32(402136)
	v103 = v101
	goto L5
L7:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v21)+108))
	v102 = int32(402120)
	v103 = v99
	goto L5
L8:
	;
	v76 = int32(1)
	v78 = int32(402073)
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)+4))
	switch v80 {
	case 0:
		v105 = int32(411781)
		v106 = v4
		v108 = v76
		v109 = v78
		goto L4
	case 1:
		goto L30
	default:
		goto L29
	}
L9:
	;
	v48 = int32(402079)
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	if v49 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L10:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v21)+16))
	v36 = F_get_rel_name(m, v35)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	return
L12:
	;
	v38 = int32(1)
	v39 = int32(402093)
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)))
	if v40 != v38 {
		v105 = v36
		v106 = v4
		v108 = v38
		v109 = v39
		goto L4
	} else {
		goto L13
	}
L13:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v21)+16))
	v44 = F_get_rel_namespace(m, v43)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L11
	} else {
		goto L14
	}
L14:
	;
	v46 = F_get_namespace_name_or_temp(m, v44)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L11
	} else {
		goto L15
	}
L15:
	;
	v105 = v36
	v106 = v46
	v108 = v38
	v109 = v39
	goto L4
L16:
	;
	v105 = int32(0)
	v106 = v4
	v108 = int32(1)
	v109 = v48
	goto L4
L17:
	;
	goto L18
L18:
	;
	v54 = int32(1)
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v49)+4))
	if v55 != v54 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v105 = int32(0)
	v106 = v4
	v108 = v54
	v109 = v48
	goto L4
L20:
	;
	goto L21
L21:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v49)+12))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
	if v62 != int32(15) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v105 = int32(0)
	v106 = v4
	v108 = v54
	v109 = v48
	goto L4
L23:
	;
	goto L24
L24:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v61)+4))
	v67 = F_get_func_name(m, v66)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L11
	} else {
		goto L25
	}
L25:
	;
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)))
	if v69 != int32(1) {
		v105 = v67
		v106 = v4
		v108 = v54
		v109 = v48
		goto L4
	} else {
		goto L26
	}
L26:
	;
	v72 = F_get_func_namespace(m, v66)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L11
	} else {
		goto L27
	}
L27:
	;
	v74 = F_get_namespace_name_or_temp(m, v72)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L11
	} else {
		goto L28
	}
L28:
	;
	v105 = v67
	v106 = v74
	v108 = v54
	v109 = v48
	goto L4
L29:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L11
	} else {
		goto L31
	}
L30:
	;
	v105 = int32(412322)
	v106 = v4
	v108 = v76
	v109 = v78
	goto L4
L31:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v79)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = v86
	F_errmsg_internal(m, int32(500398), v10+int32(-16))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L11
	} else {
		goto L32
	}
L32:
	;
	F_errfinish(m, int32(522511), int32(4459), int32(324146))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L11
	} else {
		goto L33
	}
L33:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L34:
	;
	m.G0 = v12 - int32(-64)
	return
L35:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	F_appendStringInfoString(m, v113, int32(287118))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L11
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	if base.B2i32(v105 != int32(0))&v108 != 0 {
		goto L62
	} else {
		goto L63
	}
L38:
	;
	if v106 == int32(0) {
		goto L41
	} else {
		goto L42
	}
L39:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v176 = F_quote_identifier(m, v30)
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L11
	} else {
		goto L60
	}
L40:
	;
	v148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105))))
	v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30))))
	if v149 == int32(0) {
		v168 = v148
		v169 = v149
		goto L52
	} else {
		goto L53
	}
L41:
	;
	if v105 == int32(0) {
		goto L39
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v131 = F_quote_identifier(m, v106)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L11
	} else {
		goto L47
	}
L44:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v122 = F_quote_identifier(m, v105)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L11
	} else {
		goto L45
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v122
	F_appendStringInfo(m, v121, int32(217381), v10+int32(-48))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L11
	} else {
		goto L46
	}
L46:
	;
	goto L40
L47:
	;
	v133 = F_quote_identifier(m, v105)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L11
	} else {
		goto L48
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = v133
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v131
	F_appendStringInfo(m, v130, int32(188413), v10+int32(-32))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L11
	} else {
		goto L49
	}
L49:
	;
	if v105 == int32(0) {
		goto L39
	} else {
		goto L50
	}
L50:
	;
	goto L40
L51:
	;
	if v169-v168 == int32(0) {
		goto L34
	} else {
		goto L59
	}
L52:
	;
	goto L51
L53:
	;
	if v148 != v149 {
		v168 = v148
		v169 = v149
		goto L52
	} else {
		goto L54
	}
L54:
	;
	v153 = v30
	v154 = v105
	goto L55
L55:
	;
	v157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v154)+1)))
	v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v153)+1)))
	if v158 == int32(0) {
		v168 = v157
		v169 = v158
		goto L52
	} else {
		goto L57
	}
L56:
	;
	v168 = v157
	v169 = v158
	goto L52
L57:
	;
	v161 = int32(1)
	if v157 == v158 {
		v153 = v153 + v161
		v154 = v154 + v161
		goto L55
	} else {
		goto L58
	}
L58:
	;
	goto L56
L59:
	;
	goto L39
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v176
	F_appendStringInfo(m, v175, int32(217381), v12)
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L11
	} else {
		goto L61
	}
L61:
	;
	goto L34
L62:
	;
	F_ExplainPropertyText(m, v109, v105, l2)
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L11
	} else {
		goto L65
	}
L63:
	;
	goto L64
L64:
	;
	if v106 != 0 {
		goto L66
	} else {
		goto L67
	}
L65:
	;
	goto L64
L66:
	;
	F_ExplainPropertyText(m, int32(533921), v106, l2)
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L11
	} else {
		goto L69
	}
L67:
	;
	goto L68
L68:
	;
	F_ExplainPropertyText(m, int32(185598), v30, l2)
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L11
	} else {
		goto L70
	}
L69:
	;
	goto L68
L70:
	;
	goto L34
}
