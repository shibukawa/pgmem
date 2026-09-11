package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

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
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v173 int32
	_ = v173
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v184 int64
	_ = v184
	var v185 int64
	_ = v185
	var v186 int32
	_ = v186
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
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v223 int64
	_ = v223
	var v225 int64
	_ = v225
	var v231 int32
	_ = v231
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v253 int64
	_ = v253
	var v255 int64
	_ = v255
	var v256 int64
	_ = v256
	var v260 int64
	_ = v260
	var v262 int64
	_ = v262
	var v263 int64
	_ = v263
	var v267 int64
	_ = v267
	var v269 int64
	_ = v269
	var v270 int64
	_ = v270
	var v274 int64
	_ = v274
	var v276 int64
	_ = v276
	var v277 int64
	_ = v277
	var v281 int64
	_ = v281
	var v283 int64
	_ = v283
	var v284 int64
	_ = v284
	var v288 int64
	_ = v288
	var v290 int64
	_ = v290
	var v291 int64
	_ = v291
	var v295 int64
	_ = v295
	var v297 int64
	_ = v297
	var v298 int64
	_ = v298
	var v302 int64
	_ = v302
	var v304 int64
	_ = v304
	var v305 int64
	_ = v305
	var v309 int64
	_ = v309
	var v311 int64
	_ = v311
	var v312 int64
	_ = v312
	var v316 int64
	_ = v316
	var v318 int64
	_ = v318
	var v319 int64
	_ = v319
	var v323 int64
	_ = v323
	var v325 int64
	_ = v325
	var v326 int64
	_ = v326
	var v330 int64
	_ = v330
	var v332 int64
	_ = v332
	var v333 int64
	_ = v333
	var v337 int64
	_ = v337
	var v339 int64
	_ = v339
	var v340 int64
	_ = v340
	var v344 int64
	_ = v344
	var v346 int64
	_ = v346
	var v347 int64
	_ = v347
	var v351 int64
	_ = v351
	var v353 int64
	_ = v353
	var v354 int64
	_ = v354
	var v358 int64
	_ = v358
	var v360 int64
	_ = v360
	var v361 int64
	_ = v361
	var v365 int32
	_ = v365
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v379 int32
	_ = v379
	var v388 int32
	_ = v388
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v396 int32
	_ = v396
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v413 int32
	_ = v413
	var v415 int32
	_ = v415
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v425 int32
	_ = v425
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v447 int32
	_ = v447
	var v449 int32
	_ = v449
	var v451 int32
	_ = v451
	var v461 int32
	_ = v461
	var v464 int32
	_ = v464
	var v468 int32
	_ = v468
	var v473 int32
	_ = v473
	var v477 int32
	_ = v477
	var v481 int32
	_ = v481
	var v486 int32
	_ = v486
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v493 int32
	_ = v493
	var v496 int32
	_ = v496
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v503 int32
	_ = v503
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
	v503 = *(*int32)(unsafe.Add(mBase, uint32(v501)+28))
	if v503 != 0 {
		v23 = v503
		v24 = v500
		goto L3
	} else {
		goto L135
	}
L6:
	;
	v490 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	if v490 == int32(0) {
		goto L130
	} else {
		goto L131
	}
L7:
	;
	F_ExplainDummyGroup(m, int32(_a_F_ExplainOneUtility_0), l2)
	mBase = m.M
	v489 = m.ExcPending
	if v489 != 0 {
		goto L15
	} else {
		goto L129
	}
L8:
	;
	v145 = m.G0
	v147 = v145 - int32(288)
	m.G0 = v147
	v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+11)))
	if v149 == int32(1) {
		goto L61
	} else {
		goto L62
	}
L9:
	;
	v110 = int32(0)
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
	v112 = F_copyObjectImpl(m, v111)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L15
	} else {
		goto L45
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
	F_appendStringInfoString(m, v47, int32(_a_F_ExplainOneUtility_1))
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
	F_ExplainDummyGroup(m, int32(_a_F_ExplainOneUtility_2), l2)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L15
	} else {
		goto L25
	}
L23:
	;
	F_ExplainDummyGroup(m, int32(_a_F_ExplainOneUtility_3), l2)
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
	F_errmsg_internal(m, int32(_a_F_ExplainOneUtility_4), v19)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L15
	} else {
		goto L27
	}
L27:
	;
	F_errfinish(m, int32(_a_F_ExplainOneUtility_5), int32(419), int32(_a_F_ExplainOneUtility_6))
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
	v81 = *(*int32)(unsafe.Add(mBase, _c_F_ExplainOneUtility[0]))
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
	v90 = *(*int32)(unsafe.Add(mBase, _c_F_ExplainOneUtility[1]))
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
	v83 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ExplainOneUtility[2])))
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
		v500 = v95
		v501 = v97
		goto L5
	} else {
		goto L40
	}
L40:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(l3)+88))
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	v104 = *(*int32)(unsafe.Add(mBase, _c_F_ExplainOneUtility[3]))
	if v104 != 0 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	m.T0[v104].(func(*base.Module, int32, int32, int32, int32, int32, int32, int32))(m, v97, int32(2048), v95, l2, v102, l4, v101)
	mBase = m.M
	goto L1
L42:
	;
	goto L43
L43:
	;
	F_standard_ExplainOneQuery(m, v97, int32(2048), v95, l2, v102, l4, v101)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L15
	} else {
		goto L44
	}
L44:
	;
	goto L1
L45:
	;
	v115 = *(*int32)(unsafe.Add(mBase, _c_F_ExplainOneUtility[0]))
	switch v115 {
	case 0:
		v122 = v110
		goto L46
	case 1:
		goto L47
	default:
		goto L48
	}
L46:
	;
	v124 = *(*int32)(unsafe.Add(mBase, _c_F_ExplainOneUtility[1]))
	if v124 != 0 {
		goto L51
	} else {
		goto L52
	}
L47:
	;
	v120 = F_JumbleQuery(m, v112)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L15
	} else {
		goto L50
	}
L48:
	;
	v117 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ExplainOneUtility[2])))
	if v117 != int32(1) {
		v122 = v110
		goto L46
	} else {
		goto L49
	}
L49:
	;
	goto L47
L50:
	;
	v122 = v120
	goto L46
L51:
	;
	m.T0[v124].(func(*base.Module, int32, int32, int32))(m, l3, v112, v122)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L15
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	v128 = F_QueryRewrite(m, v112)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L15
	} else {
		goto L55
	}
L54:
	;
	goto L53
L55:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v128)+12))
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v130)))
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v131)+4))
	if v132 == int32(6) {
		v500 = int32(0)
		v501 = v131
		goto L5
	} else {
		goto L56
	}
L56:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	v136 = *(*int32)(unsafe.Add(mBase, uint32(l3)+88))
	v137 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	v139 = *(*int32)(unsafe.Add(mBase, _c_F_ExplainOneUtility[3]))
	if v139 != 0 {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	m.T0[v139].(func(*base.Module, int32, int32, int32, int32, int32, int32, int32))(m, v131, v135, int32(0), l2, v137, l4, v136)
	mBase = m.M
	goto L1
L58:
	;
	goto L59
L59:
	;
	F_standard_ExplainOneQuery(m, v131, v135, int32(0), l2, v137, l4, v136)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L15
	} else {
		goto L60
	}
L60:
	;
	goto L1
L61:
	;
	v153 = *(*int32)(unsafe.Add(mBase, _c_F_ExplainOneUtility[4]))
	v158 = F_AllocSetContextCreateInternal(m, v153, int32(_a_F_ExplainOneUtility_7), int32(0), int32(_a_F_ExplainOneUtility_8), int32(_a_F_ExplainOneUtility_9))
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L15
	} else {
		goto L64
	}
L62:
	;
	v164 = v6
	v165 = v6
	goto L63
L63:
	;
	v166 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+7)))
	if v166 == int32(1) {
		goto L65
	} else {
		goto L66
	}
L64:
	;
	v160 = int32(_a_F_ExplainOneUtility_10)
	v161 = *(*int32)(unsafe.Add(mBase, _c_F_ExplainOneUtility[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_ExplainOneUtility[4])) = v158
	v164 = v158
	v165 = v161
	goto L63
L65:
	;
	goto L69
L66:
	;
	goto L67
L67:
	;
	F___clock_gettime(m, int32(1), v147+int32(24))
	mBase = m.M
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	v181 = *(*int32)(unsafe.Add(mBase, _c_F_ExplainOneUtility[5]))
	if v181 == int32(0) {
		goto L74
	} else {
		goto L75
	}
L68:
	;
	goto L67
L69:
	;
	v173 = F__emscripten_memcpy_bulkmem(m, v147+int32(152), int32(_a_F_ExplainOneUtility_11), int32(128))
	mBase = m.M
	goto L71
L71:
	;
	goto L68
L72:
	;
	goto L1
L73:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v477 = m.ExcPending
	if v477 != 0 {
		goto L15
	} else {
		goto L126
	}
L74:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v461 = m.ExcPending
	if v461 != 0 {
		goto L15
	} else {
		goto L122
	}
L75:
	;
	v184 = int64(*(*int32)(unsafe.Add(mBase, uint32(v147)+32)))
	v185 = *(*int64)(unsafe.Add(mBase, uint32(v147)+24))
	v186 = int32(0)
	v188 = F_hash_search(m, v181, v179, v186, v186)
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L15
	} else {
		goto L76
	}
L76:
	;
	if v188 == int32(0) {
		goto L74
	} else {
		goto L77
	}
L77:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v188)+64))
	v193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192)+48)))
	if v193 == int32(0) {
		goto L73
	} else {
		goto L78
	}
L78:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v192)+12))
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v192)+24))
	if v197 != 0 {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v199 = F_make_parsestate(m, int32(0))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L15
	} else {
		goto L82
	}
L80:
	;
	v211 = v6
	v212 = v6
	v213 = v192
	goto L81
L81:
	;
	v215 = *(*int32)(unsafe.Add(mBase, _c_F_ExplainOneUtility[6]))
	v216 = *(*int32)(unsafe.Add(mBase, uint32(l3)+88))
	v217 = F_GetCachedPlan(m, v213, v212, v215, v216)
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L15
	} else {
		goto L85
	}
L82:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v199)+4)) = v201
	v203 = F_CreateExecutorState(m)
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L15
	} else {
		goto L83
	}
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v203)+88)) = l4
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	v207 = F_EvaluateParams(m, v199, v188, v206, v203)
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L15
	} else {
		goto L84
	}
L84:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v188)+64))
	v211 = v203
	v212 = v207
	v213 = v209
	goto L81
L85:
	;
	v219 = int32(1)
	F___clock_gettime(m, v219, v147+int32(24))
	mBase = m.M
	v223 = int64(*(*int32)(unsafe.Add(mBase, uint32(v147)+32)))
	v225 = *(*int64)(unsafe.Add(mBase, uint32(v147)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v147)+280)) = v223 - v184 + (v225-v185)*int64(1000000000)
	v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+11)))
	if v231 == v219 {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExplainOneUtility[4])) = v165
	F_MemoryContextMemConsumed(m, v164, v147+int32(8))
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L15
	} else {
		goto L89
	}
L87:
	;
	goto L88
L88:
	;
	v240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+7)))
	if v240 == int32(1) {
		goto L90
	} else {
		goto L91
	}
L89:
	;
	goto L88
L90:
	;
	v248 = F__emscripten_memset_bulkmem(m, v147+int32(24), base.I32_extend8_s(int32(0)), int32(128))
	mBase = m.M
	goto L93
L91:
	;
	goto L92
L92:
	;
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v217)+4))
	if v365 == int32(0) {
		goto L95
	} else {
		goto L96
	}
L93:
	;
	v250 = v147 + int32(24)
	v252 = v147 + int32(152)
	v253 = *(*int64)(unsafe.Add(mBase, uint32(v250)))
	v255 = *(*int64)(unsafe.Add(mBase, _c_F_ExplainOneUtility[7]))
	v256 = *(*int64)(unsafe.Add(mBase, uint32(v252)))
	*(*int64)(unsafe.Add(mBase, uint32(v250))) = v253 + (v255 - v256)
	v260 = *(*int64)(unsafe.Add(mBase, uint32(v250)+8))
	v262 = *(*int64)(unsafe.Add(mBase, _c_F_ExplainOneUtility[8]))
	v263 = *(*int64)(unsafe.Add(mBase, uint32(v252)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v250)+8)) = v260 + (v262 - v263)
	v267 = *(*int64)(unsafe.Add(mBase, uint32(v250)+16))
	v269 = *(*int64)(unsafe.Add(mBase, _c_F_ExplainOneUtility[9]))
	v270 = *(*int64)(unsafe.Add(mBase, uint32(v252)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v250)+16)) = v267 + (v269 - v270)
	v274 = *(*int64)(unsafe.Add(mBase, uint32(v250)+24))
	v276 = *(*int64)(unsafe.Add(mBase, _c_F_ExplainOneUtility[10]))
	v277 = *(*int64)(unsafe.Add(mBase, uint32(v252)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v250)+24)) = v274 + (v276 - v277)
	v281 = *(*int64)(unsafe.Add(mBase, uint32(v250)+32))
	v283 = *(*int64)(unsafe.Add(mBase, _c_F_ExplainOneUtility[11]))
	v284 = *(*int64)(unsafe.Add(mBase, uint32(v252)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v250)+32)) = v281 + (v283 - v284)
	v288 = *(*int64)(unsafe.Add(mBase, uint32(v250)+40))
	v290 = *(*int64)(unsafe.Add(mBase, _c_F_ExplainOneUtility[12]))
	v291 = *(*int64)(unsafe.Add(mBase, uint32(v252)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v250)+40)) = v288 + (v290 - v291)
	v295 = *(*int64)(unsafe.Add(mBase, uint32(v250)+48))
	v297 = *(*int64)(unsafe.Add(mBase, _c_F_ExplainOneUtility[13]))
	v298 = *(*int64)(unsafe.Add(mBase, uint32(v252)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v250)+48)) = v295 + (v297 - v298)
	v302 = *(*int64)(unsafe.Add(mBase, uint32(v250)+56))
	v304 = *(*int64)(unsafe.Add(mBase, _c_F_ExplainOneUtility[14]))
	v305 = *(*int64)(unsafe.Add(mBase, uint32(v252)+56))
	*(*int64)(unsafe.Add(mBase, uint32(v250)+56)) = v302 + (v304 - v305)
	v309 = *(*int64)(unsafe.Add(mBase, uint32(v250)+64))
	v311 = *(*int64)(unsafe.Add(mBase, _c_F_ExplainOneUtility[15]))
	v312 = *(*int64)(unsafe.Add(mBase, uint32(v252)+64))
	*(*int64)(unsafe.Add(mBase, uint32(v250)+64)) = v309 + (v311 - v312)
	v316 = *(*int64)(unsafe.Add(mBase, uint32(v250)+72))
	v318 = *(*int64)(unsafe.Add(mBase, _c_F_ExplainOneUtility[16]))
	v319 = *(*int64)(unsafe.Add(mBase, uint32(v252)+72))
	*(*int64)(unsafe.Add(mBase, uint32(v250)+72)) = v316 + (v318 - v319)
	v323 = *(*int64)(unsafe.Add(mBase, uint32(v250)+80))
	v325 = *(*int64)(unsafe.Add(mBase, _c_F_ExplainOneUtility[17]))
	v326 = *(*int64)(unsafe.Add(mBase, uint32(v252)+80))
	*(*int64)(unsafe.Add(mBase, uint32(v250)+80)) = v323 + (v325 - v326)
	v330 = *(*int64)(unsafe.Add(mBase, uint32(v250)+88))
	v332 = *(*int64)(unsafe.Add(mBase, _c_F_ExplainOneUtility[18]))
	v333 = *(*int64)(unsafe.Add(mBase, uint32(v252)+88))
	*(*int64)(unsafe.Add(mBase, uint32(v250)+88)) = v330 + (v332 - v333)
	v337 = *(*int64)(unsafe.Add(mBase, uint32(v250)+96))
	v339 = *(*int64)(unsafe.Add(mBase, _c_F_ExplainOneUtility[19]))
	v340 = *(*int64)(unsafe.Add(mBase, uint32(v252)+96))
	*(*int64)(unsafe.Add(mBase, uint32(v250)+96)) = v337 + (v339 - v340)
	v344 = *(*int64)(unsafe.Add(mBase, uint32(v250)+104))
	v346 = *(*int64)(unsafe.Add(mBase, _c_F_ExplainOneUtility[20]))
	v347 = *(*int64)(unsafe.Add(mBase, uint32(v252)+104))
	*(*int64)(unsafe.Add(mBase, uint32(v250)+104)) = v344 + (v346 - v347)
	v351 = *(*int64)(unsafe.Add(mBase, uint32(v250)+112))
	v353 = *(*int64)(unsafe.Add(mBase, _c_F_ExplainOneUtility[21]))
	v354 = *(*int64)(unsafe.Add(mBase, uint32(v252)+112))
	*(*int64)(unsafe.Add(mBase, uint32(v250)+112)) = v351 + (v353 - v354)
	v358 = *(*int64)(unsafe.Add(mBase, uint32(v250)+120))
	v360 = *(*int64)(unsafe.Add(mBase, _c_F_ExplainOneUtility[22]))
	v361 = *(*int64)(unsafe.Add(mBase, uint32(v252)+120))
	*(*int64)(unsafe.Add(mBase, uint32(v250)+120)) = v358 + (v360 - v361)
	goto L94
L94:
	;
	goto L92
L95:
	;
	if v211 != 0 {
		goto L117
	} else {
		goto L118
	}
L96:
	;
	v368 = int32(0)
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v365)+4))
	if v369 <= v368 {
		goto L95
	} else {
		goto L97
	}
L97:
	;
	v379 = v368
	goto L98
L98:
	;
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v365)+12))
	v391 = v388 + v379<<(uint(int32(2))%32)
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v391)))
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v392)+4))
	if v393 != int32(6) {
		goto L101
	} else {
		goto L102
	}
L99:
	;
	goto L95
L100:
	;
	v415 = v391 + int32(4)
	if v415 == int32(0) {
		goto L112
	} else {
		goto L113
	}
L101:
	;
	v396 = *(*int32)(unsafe.Add(mBase, uint32(l3)+88))
	v402 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+7)))
	if v402 != 0 {
		goto L104
	} else {
		goto L105
	}
L102:
	;
	goto L103
L103:
	;
	v411 = *(*int32)(unsafe.Add(mBase, uint32(v392)+88))
	F_ExplainOneUtility(m, v411, v24, l2, l3, v212)
	mBase = m.M
	v413 = m.ExcPending
	if v413 != 0 {
		goto L15
	} else {
		goto L111
	}
L104:
	;
	v403 = v147 + int32(24)
	goto L106
L105:
	;
	v403 = int32(0)
	goto L106
L106:
	;
	v407 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+11)))
	if v407 != 0 {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	v408 = v147 + int32(8)
	goto L109
L108:
	;
	v408 = int32(0)
	goto L109
L109:
	;
	F_ExplainOnePlan(m, v392, v24, l2, v196, v212, v396, v147+int32(280), v403, v408)
	mBase = m.M
	v410 = m.ExcPending
	if v410 != 0 {
		goto L15
	} else {
		goto L110
	}
L110:
	;
	goto L100
L111:
	;
	goto L100
L112:
	;
	v427 = v379 + int32(1)
	v428 = *(*int32)(unsafe.Add(mBase, uint32(v365)+4))
	if v427 < v428 {
		v379 = v427
		goto L98
	} else {
		goto L116
	}
L113:
	;
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v365)+12))
	v419 = *(*int32)(unsafe.Add(mBase, uint32(v365)+4))
	if base.Ui32(v418+v419<<(uint(int32(2))%32)) <= base.Ui32(v415) {
		goto L112
	} else {
		goto L114
	}
L114:
	;
	F_ExplainSeparatePlans(m, l2)
	mBase = m.M
	v425 = m.ExcPending
	if v425 != 0 {
		goto L15
	} else {
		goto L115
	}
L115:
	;
	goto L112
L116:
	;
	goto L99
L117:
	;
	F_FreeExecutorState(m, v211)
	mBase = m.M
	v447 = m.ExcPending
	if v447 != 0 {
		goto L15
	} else {
		goto L120
	}
L118:
	;
	goto L119
L119:
	;
	v449 = *(*int32)(unsafe.Add(mBase, _c_F_ExplainOneUtility[6]))
	F_ReleaseCachedPlan(m, v217, v449)
	mBase = m.M
	v451 = m.ExcPending
	if v451 != 0 {
		goto L15
	} else {
		goto L121
	}
L120:
	;
	goto L119
L121:
	;
	m.G0 = v147 + int32(288)
	goto L72
L122:
	;
	F_errcode(m, int32(386))
	mBase = m.M
	v464 = m.ExcPending
	if v464 != 0 {
		goto L15
	} else {
		goto L123
	}
L123:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v147))) = v179
	F_errmsg(m, int32(_a_F_ExplainOneUtility_12), v147)
	mBase = m.M
	v468 = m.ExcPending
	if v468 != 0 {
		goto L15
	} else {
		goto L124
	}
L124:
	;
	F_errfinish(m, int32(_a_F_ExplainOneUtility_13), int32(454), int32(_a_F_ExplainOneUtility_14))
	mBase = m.M
	v473 = m.ExcPending
	if v473 != 0 {
		goto L15
	} else {
		goto L125
	}
L125:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L126:
	;
	F_errmsg_internal(m, int32(_a_F_ExplainOneUtility_15), int32(0))
	mBase = m.M
	v481 = m.ExcPending
	if v481 != 0 {
		goto L15
	} else {
		goto L127
	}
L127:
	;
	F_errfinish(m, int32(_a_F_ExplainOneUtility_13), int32(608), int32(_a_F_ExplainOneUtility_16))
	mBase = m.M
	v486 = m.ExcPending
	if v486 != 0 {
		goto L15
	} else {
		goto L128
	}
L128:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L129:
	;
	goto L1
L130:
	;
	v493 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	F_appendStringInfoString(m, v493, int32(_a_F_ExplainOneUtility_17))
	mBase = m.M
	v496 = m.ExcPending
	if v496 != 0 {
		goto L15
	} else {
		goto L133
	}
L131:
	;
	goto L132
L132:
	;
	F_ExplainDummyGroup(m, int32(_a_F_ExplainOneUtility_18), l2)
	mBase = m.M
	v499 = m.ExcPending
	if v499 != 0 {
		goto L15
	} else {
		goto L134
	}
L133:
	;
	goto L1
L134:
	;
	goto L1
L135:
	;
	goto L4
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
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v152 int32
	_ = v152
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
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
		v116 = int32(0)
		v118 = v4
		v119 = v4
		v120 = v4
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
	v121 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	if v121 == int32(0) {
		goto L39
	} else {
		goto L40
	}
L5:
	;
	v116 = v113
	v118 = v4
	v119 = int32(1)
	v120 = v112
	goto L4
L6:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v21)+84))
	v112 = int32(_a_F_ExplainTargetRel_0)
	v113 = v111
	goto L5
L7:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v21)+108))
	v112 = int32(_a_F_ExplainTargetRel_1)
	v113 = v109
	goto L5
L8:
	;
	v86 = int32(1)
	v88 = int32(_a_F_ExplainTargetRel_2)
	v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v89)+4))
	switch v90 {
	case 0:
		v116 = int32(_a_F_ExplainTargetRel_3)
		v118 = v4
		v119 = v86
		v120 = v88
		goto L4
	case 1:
		goto L34
	default:
		goto L33
	}
L9:
	;
	v48 = int32(_a_F_ExplainTargetRel_4)
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
	v39 = int32(_a_F_ExplainTargetRel_5)
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)))
	if v40 != v38 {
		v116 = v36
		v118 = v4
		v119 = v38
		v120 = v39
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
	v116 = v36
	v118 = v46
	v119 = v38
	v120 = v39
	goto L4
L16:
	;
	v116 = int32(0)
	v118 = v4
	v119 = int32(1)
	v120 = v48
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
	v116 = int32(0)
	v118 = v4
	v119 = v54
	v120 = v48
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
	v116 = int32(0)
	v118 = v4
	v119 = v54
	v120 = v48
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
		v116 = v67
		v118 = v4
		v119 = v54
		v120 = v48
		goto L4
	} else {
		goto L26
	}
L26:
	;
	v73 = F_SearchSysCache1(m, int32(47), v66)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L11
	} else {
		goto L27
	}
L27:
	;
	if v73 != 0 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v73)+16))
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75)+22)))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v75+v76)+68))
	F_ReleaseCatCache(m, v73)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L11
	} else {
		goto L31
	}
L29:
	;
	v83 = int32(0)
	goto L30
L30:
	;
	v84 = F_get_namespace_name_or_temp(m, v83)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L11
	} else {
		goto L32
	}
L31:
	;
	v83 = v78
	goto L30
L32:
	;
	v116 = v67
	v118 = v84
	v119 = v54
	v120 = v48
	goto L4
L33:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L11
	} else {
		goto L35
	}
L34:
	;
	v116 = int32(_a_F_ExplainTargetRel_6)
	v118 = v4
	v119 = v86
	v120 = v88
	goto L4
L35:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v89)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = v96
	F_errmsg_internal(m, int32(_a_F_ExplainTargetRel_7), v10+int32(-16))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L11
	} else {
		goto L36
	}
L36:
	;
	F_errfinish(m, int32(_a_F_ExplainTargetRel_8), int32(_a_F_ExplainTargetRel_9), int32(_a_F_ExplainTargetRel_10))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L11
	} else {
		goto L37
	}
L37:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L38:
	;
	m.G0 = v12 - int32(-64)
	return
L39:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	F_appendStringInfoString(m, v124, int32(_a_F_ExplainTargetRel_11))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L11
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	if base.B2i32(v116 != int32(0))&v119 != 0 {
		goto L66
	} else {
		goto L67
	}
L42:
	;
	if v118 == int32(0) {
		goto L45
	} else {
		goto L46
	}
L43:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v187 = F_quote_identifier(m, v30)
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L11
	} else {
		goto L64
	}
L44:
	;
	v159 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116))))
	v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30))))
	if v160 == int32(0) {
		v179 = v159
		v180 = v160
		goto L56
	} else {
		goto L57
	}
L45:
	;
	if v116 == int32(0) {
		goto L43
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v142 = F_quote_identifier(m, v118)
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L11
	} else {
		goto L51
	}
L48:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v133 = F_quote_identifier(m, v116)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L11
	} else {
		goto L49
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v133
	F_appendStringInfo(m, v132, int32(_a_F_ExplainTargetRel_12), v10+int32(-48))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L11
	} else {
		goto L50
	}
L50:
	;
	goto L44
L51:
	;
	v144 = F_quote_identifier(m, v116)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L11
	} else {
		goto L52
	}
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = v144
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v142
	F_appendStringInfo(m, v141, int32(_a_F_ExplainTargetRel_13), v10+int32(-32))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L11
	} else {
		goto L53
	}
L53:
	;
	if v116 == int32(0) {
		goto L43
	} else {
		goto L54
	}
L54:
	;
	goto L44
L55:
	;
	if v180-v179 == int32(0) {
		goto L38
	} else {
		goto L63
	}
L56:
	;
	goto L55
L57:
	;
	if v159 != v160 {
		v179 = v159
		v180 = v160
		goto L56
	} else {
		goto L58
	}
L58:
	;
	v164 = v30
	v165 = v116
	goto L59
L59:
	;
	v168 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+1)))
	v169 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v164)+1)))
	if v169 == int32(0) {
		v179 = v168
		v180 = v169
		goto L56
	} else {
		goto L61
	}
L60:
	;
	v179 = v168
	v180 = v169
	goto L56
L61:
	;
	v172 = int32(1)
	if v168 == v169 {
		v164 = v164 + v172
		v165 = v165 + v172
		goto L59
	} else {
		goto L62
	}
L62:
	;
	goto L60
L63:
	;
	goto L43
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v187
	F_appendStringInfo(m, v186, int32(_a_F_ExplainTargetRel_12), v12)
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L11
	} else {
		goto L65
	}
L65:
	;
	goto L38
L66:
	;
	F_ExplainPropertyText(m, v120, v116, l2)
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L11
	} else {
		goto L69
	}
L67:
	;
	goto L68
L68:
	;
	if v118 != 0 {
		goto L70
	} else {
		goto L71
	}
L69:
	;
	goto L68
L70:
	;
	F_ExplainPropertyText(m, int32(_a_F_ExplainTargetRel_14), v118, l2)
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L11
	} else {
		goto L73
	}
L71:
	;
	goto L72
L72:
	;
	F_ExplainPropertyText(m, int32(_a_F_ExplainTargetRel_15), v30, l2)
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L11
	} else {
		goto L74
	}
L73:
	;
	goto L72
L74:
	;
	goto L38
}
