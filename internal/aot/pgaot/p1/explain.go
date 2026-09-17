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
		F_appendStringInfoString(m, v9, int32(_a_F_ExplainEndOutput_0))
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
		F_appendStringInfoString(m, v17, int32(_a_F_ExplainEndOutput_1))
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
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
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
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v187 int64
	_ = v187
	var v188 int64
	_ = v188
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v226 int64
	_ = v226
	var v228 int64
	_ = v228
	var v234 int32
	_ = v234
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v247 int32
	_ = v247
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
	var v366 int32
	_ = v366
	var v369 int32
	_ = v369
	var v377 int32
	_ = v377
	var v389 int32
	_ = v389
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v397 int32
	_ = v397
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v414 int32
	_ = v414
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v424 int32
	_ = v424
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v446 int32
	_ = v446
	var v448 int32
	_ = v448
	var v450 int32
	_ = v450
	var v460 int32
	_ = v460
	var v463 int32
	_ = v463
	var v467 int32
	_ = v467
	var v472 int32
	_ = v472
	var v476 int32
	_ = v476
	var v480 int32
	_ = v480
	var v485 int32
	_ = v485
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v492 int32
	_ = v492
	var v495 int32
	_ = v495
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v502 int32
	_ = v502
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
	v502 = *(*int32)(unsafe.Add(mBase, uint32(v500)+28))
	if v502 != 0 {
		v23 = v502
		v24 = v499
		goto L3
	} else {
		goto L134
	}
L6:
	;
	v489 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	if v489 == int32(0) {
		goto L129
	} else {
		goto L130
	}
L7:
	;
	F_ExplainDummyGroup(m, int32(_a_F_ExplainOneUtility_0), l2)
	mBase = m.M
	v488 = m.ExcPending
	if v488 != 0 {
		goto L15
	} else {
		goto L128
	}
L8:
	;
	v149 = m.G0
	v151 = v149 - int32(288)
	m.G0 = v151
	v153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+11)))
	if v153 == int32(1) {
		goto L65
	} else {
		goto L66
	}
L9:
	;
	v113 = int32(0)
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
	v115 = F_copyObjectImpl(m, v114)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L15
	} else {
		goto L48
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
	v55 = v53 - int32(23)
	if v55 != 0 {
		goto L22
	} else {
		goto L23
	}
L19:
	;
	goto L20
L20:
	;
	v78 = int32(0)
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	v80 = F_copyObjectImpl(m, v79)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L15
	} else {
		goto L31
	}
L21:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L15
	} else {
		goto L28
	}
L22:
	;
	if v55 != int32(18) {
		goto L21
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	F_ExplainDummyGroup(m, int32(_a_F_ExplainOneUtility_2), l2)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L15
	} else {
		goto L27
	}
L25:
	;
	F_ExplainDummyGroup(m, int32(_a_F_ExplainOneUtility_3), l2)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L15
	} else {
		goto L26
	}
L26:
	;
	goto L1
L27:
	;
	goto L1
L28:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v68
	F_errmsg_internal(m, int32(_a_F_ExplainOneUtility_4), v19)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L15
	} else {
		goto L29
	}
L29:
	;
	F_errfinish(m, int32(_a_F_ExplainOneUtility_5), int32(419), int32(_a_F_ExplainOneUtility_6))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L15
	} else {
		goto L30
	}
L30:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L31:
	;
	v83 = *(*int32)(unsafe.Add(mBase, _c_F_ExplainOneUtility[0]))
	switch v83 {
	case 0:
		v90 = v78
		goto L32
	case 1:
		goto L33
	default:
		goto L34
	}
L32:
	;
	v92 = *(*int32)(unsafe.Add(mBase, _c_F_ExplainOneUtility[1]))
	if v92 != 0 {
		goto L37
	} else {
		goto L38
	}
L33:
	;
	v88 = F_JumbleQuery(m, v80)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L15
	} else {
		goto L36
	}
L34:
	;
	v85 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ExplainOneUtility[2])))
	if v85 != int32(1) {
		v90 = v78
		goto L32
	} else {
		goto L35
	}
L35:
	;
	goto L33
L36:
	;
	v90 = v88
	goto L32
L37:
	;
	m.T0[v92].(func(*base.Module, int32, int32, int32))(m, l3, v80, v90)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L15
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	v95 = F_QueryRewrite(m, v80)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L15
	} else {
		goto L41
	}
L40:
	;
	goto L39
L41:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v95)+12))
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v98)))
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v99)+4))
	if v100 == int32(6) {
		v499 = v97
		v500 = v99
		goto L5
	} else {
		goto L42
	}
L42:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(l3)+88))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	v106 = *(*int32)(unsafe.Add(mBase, _c_F_ExplainOneUtility[3]))
	if v106 != 0 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	m.T0[v106].(func(*base.Module, int32, int32, int32, int32, int32, int32, int32))(m, v99, int32(2048), v97, l2, v104, l4, v103)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L15
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	F_standard_ExplainOneQuery(m, v99, int32(2048), v97, l2, v104, l4, v103)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L15
	} else {
		goto L47
	}
L46:
	;
	goto L1
L47:
	;
	goto L1
L48:
	;
	v118 = *(*int32)(unsafe.Add(mBase, _c_F_ExplainOneUtility[0]))
	switch v118 {
	case 0:
		v125 = v113
		goto L49
	case 1:
		goto L50
	default:
		goto L51
	}
L49:
	;
	v127 = *(*int32)(unsafe.Add(mBase, _c_F_ExplainOneUtility[1]))
	if v127 != 0 {
		goto L54
	} else {
		goto L55
	}
L50:
	;
	v123 = F_JumbleQuery(m, v115)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L15
	} else {
		goto L53
	}
L51:
	;
	v120 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ExplainOneUtility[2])))
	if v120 != int32(1) {
		v125 = v113
		goto L49
	} else {
		goto L52
	}
L52:
	;
	goto L50
L53:
	;
	v125 = v123
	goto L49
L54:
	;
	m.T0[v127].(func(*base.Module, int32, int32, int32))(m, l3, v115, v125)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L15
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	v131 = F_QueryRewrite(m, v115)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L15
	} else {
		goto L58
	}
L57:
	;
	goto L56
L58:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v131)+12))
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v133)))
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v134)+4))
	if v135 == int32(6) {
		v499 = int32(0)
		v500 = v134
		goto L5
	} else {
		goto L59
	}
L59:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	v139 = *(*int32)(unsafe.Add(mBase, uint32(l3)+88))
	v140 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	v142 = *(*int32)(unsafe.Add(mBase, _c_F_ExplainOneUtility[3]))
	if v142 != 0 {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	m.T0[v142].(func(*base.Module, int32, int32, int32, int32, int32, int32, int32))(m, v134, v138, int32(0), l2, v140, l4, v139)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L15
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	F_standard_ExplainOneQuery(m, v134, v138, int32(0), l2, v140, l4, v139)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L15
	} else {
		goto L64
	}
L63:
	;
	goto L1
L64:
	;
	goto L1
L65:
	;
	v157 = *(*int32)(unsafe.Add(mBase, _c_F_ExplainOneUtility[4]))
	v162 = F_AllocSetContextCreateInternal(m, v157, int32(_a_F_ExplainOneUtility_7), int32(0), int32(_a_F_ExplainOneUtility_8), int32(_a_F_ExplainOneUtility_9))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L15
	} else {
		goto L68
	}
L66:
	;
	v168 = v6
	v169 = v6
	goto L67
L67:
	;
	v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+7)))
	if v170 == int32(1) {
		goto L69
	} else {
		goto L70
	}
L68:
	;
	v164 = int32(_a_F_ExplainOneUtility_10)
	v165 = *(*int32)(unsafe.Add(mBase, _c_F_ExplainOneUtility[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_ExplainOneUtility[4])) = v162
	v168 = v162
	v169 = v165
	goto L67
L69:
	;
	base.MemoryCopy(m, v151+int32(152), int32(_a_F_ExplainOneUtility_11), int32(128))
	goto L71
L70:
	;
	goto L71
L71:
	;
	F___clock_gettime(m, int32(1), v151+int32(24))
	mBase = m.M
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	v184 = *(*int32)(unsafe.Add(mBase, _c_F_ExplainOneUtility[5]))
	if v184 == int32(0) {
		goto L74
	} else {
		goto L75
	}
L72:
	;
	goto L1
L73:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v476 = m.ExcPending
	if v476 != 0 {
		goto L15
	} else {
		goto L125
	}
L74:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v460 = m.ExcPending
	if v460 != 0 {
		goto L15
	} else {
		goto L121
	}
L75:
	;
	v187 = int64(*(*int32)(unsafe.Add(mBase, uint32(v151)+32)))
	v188 = *(*int64)(unsafe.Add(mBase, uint32(v151)+24))
	v189 = int32(0)
	v191 = F_hash_search(m, v184, v182, v189, v189)
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L15
	} else {
		goto L76
	}
L76:
	;
	if v191 == int32(0) {
		goto L74
	} else {
		goto L77
	}
L77:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v191)+64))
	v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v195)+48)))
	if v196 == int32(0) {
		goto L73
	} else {
		goto L78
	}
L78:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v195)+12))
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v195)+24))
	if v200 != 0 {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v202 = F_make_parsestate(m, int32(0))
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L15
	} else {
		goto L82
	}
L80:
	;
	v214 = v6
	v215 = v6
	v216 = v195
	goto L81
L81:
	;
	v218 = *(*int32)(unsafe.Add(mBase, _c_F_ExplainOneUtility[6]))
	v219 = *(*int32)(unsafe.Add(mBase, uint32(l3)+88))
	v220 = F_GetCachedPlan(m, v216, v215, v218, v219)
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L15
	} else {
		goto L85
	}
L82:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v202)+4)) = v204
	v206 = F_CreateExecutorState(m)
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L15
	} else {
		goto L83
	}
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v206)+88)) = l4
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	v210 = F_EvaluateParams(m, v202, v191, v209, v206)
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L15
	} else {
		goto L84
	}
L84:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v191)+64))
	v214 = v206
	v215 = v210
	v216 = v212
	goto L81
L85:
	;
	v222 = int32(1)
	F___clock_gettime(m, v222, v151+int32(24))
	mBase = m.M
	v226 = int64(*(*int32)(unsafe.Add(mBase, uint32(v151)+32)))
	v228 = *(*int64)(unsafe.Add(mBase, uint32(v151)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v151)+280)) = v226 - v187 + (v228-v188)*int64(1000000000)
	v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+11)))
	if v234 == v222 {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExplainOneUtility[4])) = v169
	F_MemoryContextMemConsumed(m, v168, v151+int32(8))
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L15
	} else {
		goto L89
	}
L87:
	;
	goto L88
L88:
	;
	v243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+7)))
	if v243 == int32(1) {
		goto L90
	} else {
		goto L91
	}
L89:
	;
	goto L88
L90:
	;
	v247 = v151 + int32(24)
	base.MemoryFill(m, v247, int32(0), int32(128))
	v252 = v151 + int32(152)
	v253 = *(*int64)(unsafe.Add(mBase, uint32(v247)))
	v255 = *(*int64)(unsafe.Add(mBase, _c_F_ExplainOneUtility[7]))
	v256 = *(*int64)(unsafe.Add(mBase, uint32(v252)))
	*(*int64)(unsafe.Add(mBase, uint32(v247))) = v253 + (v255 - v256)
	v260 = *(*int64)(unsafe.Add(mBase, uint32(v247)+8))
	v262 = *(*int64)(unsafe.Add(mBase, _c_F_ExplainOneUtility[8]))
	v263 = *(*int64)(unsafe.Add(mBase, uint32(v252)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v247)+8)) = v260 + (v262 - v263)
	v267 = *(*int64)(unsafe.Add(mBase, uint32(v247)+16))
	v269 = *(*int64)(unsafe.Add(mBase, _c_F_ExplainOneUtility[9]))
	v270 = *(*int64)(unsafe.Add(mBase, uint32(v252)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v247)+16)) = v267 + (v269 - v270)
	v274 = *(*int64)(unsafe.Add(mBase, uint32(v247)+24))
	v276 = *(*int64)(unsafe.Add(mBase, _c_F_ExplainOneUtility[10]))
	v277 = *(*int64)(unsafe.Add(mBase, uint32(v252)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v247)+24)) = v274 + (v276 - v277)
	v281 = *(*int64)(unsafe.Add(mBase, uint32(v247)+32))
	v283 = *(*int64)(unsafe.Add(mBase, _c_F_ExplainOneUtility[11]))
	v284 = *(*int64)(unsafe.Add(mBase, uint32(v252)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v247)+32)) = v281 + (v283 - v284)
	v288 = *(*int64)(unsafe.Add(mBase, uint32(v247)+40))
	v290 = *(*int64)(unsafe.Add(mBase, _c_F_ExplainOneUtility[12]))
	v291 = *(*int64)(unsafe.Add(mBase, uint32(v252)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v247)+40)) = v288 + (v290 - v291)
	v295 = *(*int64)(unsafe.Add(mBase, uint32(v247)+48))
	v297 = *(*int64)(unsafe.Add(mBase, _c_F_ExplainOneUtility[13]))
	v298 = *(*int64)(unsafe.Add(mBase, uint32(v252)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v247)+48)) = v295 + (v297 - v298)
	v302 = *(*int64)(unsafe.Add(mBase, uint32(v247)+56))
	v304 = *(*int64)(unsafe.Add(mBase, _c_F_ExplainOneUtility[14]))
	v305 = *(*int64)(unsafe.Add(mBase, uint32(v252)+56))
	*(*int64)(unsafe.Add(mBase, uint32(v247)+56)) = v302 + (v304 - v305)
	v309 = *(*int64)(unsafe.Add(mBase, uint32(v247)+64))
	v311 = *(*int64)(unsafe.Add(mBase, _c_F_ExplainOneUtility[15]))
	v312 = *(*int64)(unsafe.Add(mBase, uint32(v252)+64))
	*(*int64)(unsafe.Add(mBase, uint32(v247)+64)) = v309 + (v311 - v312)
	v316 = *(*int64)(unsafe.Add(mBase, uint32(v247)+72))
	v318 = *(*int64)(unsafe.Add(mBase, _c_F_ExplainOneUtility[16]))
	v319 = *(*int64)(unsafe.Add(mBase, uint32(v252)+72))
	*(*int64)(unsafe.Add(mBase, uint32(v247)+72)) = v316 + (v318 - v319)
	v323 = *(*int64)(unsafe.Add(mBase, uint32(v247)+80))
	v325 = *(*int64)(unsafe.Add(mBase, _c_F_ExplainOneUtility[17]))
	v326 = *(*int64)(unsafe.Add(mBase, uint32(v252)+80))
	*(*int64)(unsafe.Add(mBase, uint32(v247)+80)) = v323 + (v325 - v326)
	v330 = *(*int64)(unsafe.Add(mBase, uint32(v247)+88))
	v332 = *(*int64)(unsafe.Add(mBase, _c_F_ExplainOneUtility[18]))
	v333 = *(*int64)(unsafe.Add(mBase, uint32(v252)+88))
	*(*int64)(unsafe.Add(mBase, uint32(v247)+88)) = v330 + (v332 - v333)
	v337 = *(*int64)(unsafe.Add(mBase, uint32(v247)+96))
	v339 = *(*int64)(unsafe.Add(mBase, _c_F_ExplainOneUtility[19]))
	v340 = *(*int64)(unsafe.Add(mBase, uint32(v252)+96))
	*(*int64)(unsafe.Add(mBase, uint32(v247)+96)) = v337 + (v339 - v340)
	v344 = *(*int64)(unsafe.Add(mBase, uint32(v247)+104))
	v346 = *(*int64)(unsafe.Add(mBase, _c_F_ExplainOneUtility[20]))
	v347 = *(*int64)(unsafe.Add(mBase, uint32(v252)+104))
	*(*int64)(unsafe.Add(mBase, uint32(v247)+104)) = v344 + (v346 - v347)
	v351 = *(*int64)(unsafe.Add(mBase, uint32(v247)+112))
	v353 = *(*int64)(unsafe.Add(mBase, _c_F_ExplainOneUtility[21]))
	v354 = *(*int64)(unsafe.Add(mBase, uint32(v252)+112))
	*(*int64)(unsafe.Add(mBase, uint32(v247)+112)) = v351 + (v353 - v354)
	v358 = *(*int64)(unsafe.Add(mBase, uint32(v247)+120))
	v360 = *(*int64)(unsafe.Add(mBase, _c_F_ExplainOneUtility[22]))
	v361 = *(*int64)(unsafe.Add(mBase, uint32(v252)+120))
	*(*int64)(unsafe.Add(mBase, uint32(v247)+120)) = v358 + (v360 - v361)
	goto L93
L91:
	;
	goto L92
L92:
	;
	v366 = *(*int32)(unsafe.Add(mBase, uint32(v220)+4))
	if v366 == int32(0) {
		goto L94
	} else {
		goto L95
	}
L93:
	;
	goto L92
L94:
	;
	if v214 != 0 {
		goto L116
	} else {
		goto L117
	}
L95:
	;
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v366)+4))
	if v369 <= int32(0) {
		goto L94
	} else {
		goto L96
	}
L96:
	;
	v377 = int32(0)
	goto L97
L97:
	;
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v366)+12))
	v392 = v389 + v377<<(uint(int32(2))%32)
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v392)))
	v394 = *(*int32)(unsafe.Add(mBase, uint32(v393)+4))
	if v394 != int32(6) {
		goto L100
	} else {
		goto L101
	}
L98:
	;
	goto L94
L99:
	;
	v417 = *(*int32)(unsafe.Add(mBase, uint32(v366)+12))
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v366)+4))
	if base.Ui32(v392+int32(4)) < base.Ui32(v417+v418<<(uint(int32(2))%32)) {
		goto L111
	} else {
		goto L112
	}
L100:
	;
	v397 = *(*int32)(unsafe.Add(mBase, uint32(l3)+88))
	v403 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+7)))
	if v403 != 0 {
		goto L103
	} else {
		goto L104
	}
L101:
	;
	goto L102
L102:
	;
	v412 = *(*int32)(unsafe.Add(mBase, uint32(v393)+88))
	F_ExplainOneUtility(m, v412, v24, l2, l3, v215)
	mBase = m.M
	v414 = m.ExcPending
	if v414 != 0 {
		goto L15
	} else {
		goto L110
	}
L103:
	;
	v404 = v151 + int32(24)
	goto L105
L104:
	;
	v404 = int32(0)
	goto L105
L105:
	;
	v408 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+11)))
	if v408 != 0 {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	v409 = v151 + int32(8)
	goto L108
L107:
	;
	v409 = int32(0)
	goto L108
L108:
	;
	F_ExplainOnePlan(m, v393, v24, l2, v199, v215, v397, v151+int32(280), v404, v409)
	mBase = m.M
	v411 = m.ExcPending
	if v411 != 0 {
		goto L15
	} else {
		goto L109
	}
L109:
	;
	goto L99
L110:
	;
	goto L99
L111:
	;
	F_ExplainSeparatePlans(m, l2)
	mBase = m.M
	v424 = m.ExcPending
	if v424 != 0 {
		goto L15
	} else {
		goto L114
	}
L112:
	;
	goto L113
L113:
	;
	v426 = v377 + int32(1)
	v427 = *(*int32)(unsafe.Add(mBase, uint32(v366)+4))
	if v426 < v427 {
		v377 = v426
		goto L97
	} else {
		goto L115
	}
L114:
	;
	goto L113
L115:
	;
	goto L98
L116:
	;
	F_FreeExecutorState(m, v214)
	mBase = m.M
	v446 = m.ExcPending
	if v446 != 0 {
		goto L15
	} else {
		goto L119
	}
L117:
	;
	goto L118
L118:
	;
	v448 = *(*int32)(unsafe.Add(mBase, _c_F_ExplainOneUtility[6]))
	F_ReleaseCachedPlan(m, v220, v448)
	mBase = m.M
	v450 = m.ExcPending
	if v450 != 0 {
		goto L15
	} else {
		goto L120
	}
L119:
	;
	goto L118
L120:
	;
	m.G0 = v151 + int32(288)
	goto L72
L121:
	;
	F_errcode(m, int32(386))
	mBase = m.M
	v463 = m.ExcPending
	if v463 != 0 {
		goto L15
	} else {
		goto L122
	}
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v151))) = v182
	F_errmsg(m, int32(_a_F_ExplainOneUtility_12), v151)
	mBase = m.M
	v467 = m.ExcPending
	if v467 != 0 {
		goto L15
	} else {
		goto L123
	}
L123:
	;
	F_errfinish(m, int32(_a_F_ExplainOneUtility_13), int32(454), int32(_a_F_ExplainOneUtility_14))
	mBase = m.M
	v472 = m.ExcPending
	if v472 != 0 {
		goto L15
	} else {
		goto L124
	}
L124:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L125:
	;
	F_errmsg_internal(m, int32(_a_F_ExplainOneUtility_15), int32(0))
	mBase = m.M
	v480 = m.ExcPending
	if v480 != 0 {
		goto L15
	} else {
		goto L126
	}
L126:
	;
	F_errfinish(m, int32(_a_F_ExplainOneUtility_13), int32(608), int32(_a_F_ExplainOneUtility_16))
	mBase = m.M
	v485 = m.ExcPending
	if v485 != 0 {
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
	goto L1
L129:
	;
	v492 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	F_appendStringInfoString(m, v492, int32(_a_F_ExplainOneUtility_17))
	mBase = m.M
	v495 = m.ExcPending
	if v495 != 0 {
		goto L15
	} else {
		goto L132
	}
L130:
	;
	goto L131
L131:
	;
	F_ExplainDummyGroup(m, int32(_a_F_ExplainOneUtility_18), l2)
	mBase = m.M
	v498 = m.ExcPending
	if v498 != 0 {
		goto L15
	} else {
		goto L133
	}
L132:
	;
	goto L1
L133:
	;
	goto L1
L134:
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
	var v35 int32
	_ = v35
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v122 int32
	_ = v122
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
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
	var v248 int32
	_ = v248
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v296 int32
	_ = v296
	var v307 int32
	_ = v307
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v325 float64
	_ = v325
	var v326 int32
	_ = v326
	var v327 float64
	_ = v327
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v338 int32
	_ = v338
	var v341 int32
	_ = v341
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v352 int32
	_ = v352
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v374 int32
	_ = v374
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v386 int32
	_ = v386
	var v391 int32
	_ = v391
	var v398 int32
	_ = v398
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v427 int32
	_ = v427
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v447 int32
	_ = v447
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v471 int32
	_ = v471
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v490 int32
	_ = v490
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v505 int32
	_ = v505
	var v517 int32
	_ = v517
	var v519 int32
	_ = v519
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v535 int32
	_ = v535
	var v541 int32
	_ = v541
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v558 int32
	_ = v558
	var v560 int32
	_ = v560
	var v572 int32
	_ = v572
	var v575 int32
	_ = v575
	var v576 int64
	_ = v576
	var v580 int32
	_ = v580
	var v586 int32
	_ = v586
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
	v35 = int32(0)
	base.MemoryFill(m, v31+int32(4), v35, int32(76))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+16)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v31))) = v27
	F_set_rtable_names(m, v31, v35, v28)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	v45 = int32(80)
	m.G0 = v31 + v45
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v44
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v50 = m.G0
	v52 = v50 - int32(16)
	m.G0 = v52
	v55 = F_palloc0(m, v45)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v49)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v55)+4)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v55))) = v57
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v49)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v55)+16)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v55)+12)) = v60
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v49)+60))
	if v64 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	F_set_simple_column_names(m, v55)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L1
	} else {
		goto L18
	}
L6:
	;
	if v57 != 0 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v55)+20)) = int32(0)
	goto L5
L9:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
	v71 = v65<<(uint(int32(2))%32) + int32(4)
	goto L11
L10:
	;
	v71 = int32(4)
	goto L11
L11:
	;
	v72 = F_palloc0(m, v71)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v55)+20)) = v72
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v49)+60))
	if v75 == int32(0) {
		goto L5
	} else {
		goto L13
	}
L13:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v75)+4))
	if v78 <= int32(0) {
		goto L5
	} else {
		goto L14
	}
L14:
	;
	v88 = int32(0)
	goto L15
L15:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v55)+20))
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v75)+12))
	v95 = int32(2)
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v94+v88<<(uint(v95)%32))))
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v98)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v93+v99<<(uint(v95)%32)))) = v98
	v105 = v88 + int32(1)
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v75)+4))
	if v105 < v106 {
		v88 = v105
		goto L15
	} else {
		goto L17
	}
L16:
	;
	goto L5
L17:
	;
	goto L16
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+8)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v52)+12)) = v55
	v128 = F_list_make1_impl(m, int32(1), v52+int32(8))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	m.G0 = v52 + int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v128
	v136 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v136 != 0 {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v184)))
	if v185 != int32(432) {
		v195 = v184
		goto L34
	} else {
		goto L35
	}
L21:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v136)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = v137
	if v137 <= int32(0) {
		goto L20
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = int32(0)
	goto L20
L24:
	;
	v141 = int32(0)
	if v141 < v137 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v144 = v137
	goto L27
L26:
	;
	v144 = v141
	goto L27
L27:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v136)+12))
	v149 = int32(0)
	goto L28
L28:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v145+v149<<(uint(int32(2))%32))))
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v161)+12))
	if v162 != int32(9) {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = v137 - int32(1)
	goto L20
L30:
	;
	v166 = v149 + int32(1)
	if v144 != v166 {
		v149 = v166
		goto L28
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	goto L29
L33:
	;
	goto L20
L34:
	;
	v196 = int32(0)
	F_ExplainNode(m, v195, v196, v196, v196, l0)
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L1
	} else {
		goto L37
	}
L35:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v184)+4))
	v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188)+81)))
	if v189 != int32(1) {
		v195 = v184
		goto L34
	} else {
		goto L36
	}
L36:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v184)+36))
	v193 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+52)) = uint8(v193)
	v195 = v192
	goto L34
L37:
	;
	v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)))
	if v201 != int32(1) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v572 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)))
	if v572 != int32(1) {
		goto L119
	} else {
		goto L120
	}
L39:
	;
	v204 = m.G0
	v206 = v204 - int32(16)
	m.G0 = v206
	v209 = v14 + int32(76)
	*(*int32)(unsafe.Add(mBase, uint32(v209))) = int32(0)
	v213 = *(*int32)(unsafe.Add(mBase, _c_F_ExplainPrintPlan[0]))
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v213)))
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v215)+4))
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v215)+412))
	if v217 != 0 {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	v283 = F_palloc(m, v280<<(uint(int32(2))%32))
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L1
	} else {
		goto L44
	}
L41:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v215)+376))
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v215)+364))
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v215)+352))
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v215)+340))
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v215)+328))
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v215)+316))
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v215)+304))
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v215)+292))
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v215)+280))
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v215)+268))
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v215)+256))
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v215)+244))
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v215)+232))
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v215)+220))
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v215)+208))
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v215)+196))
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v215)+184))
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v215)+172))
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v215)+160))
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v215)+148))
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v215)+136))
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v215)+124))
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v215)+112))
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v215)+100))
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v215)+88))
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v215)+76))
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v215)+64))
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v215)+52))
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v215)+40))
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v215)+28))
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v215)+16))
	v280 = v218 + (v219 + (v220 + (v221 + (v222 + (v223 + (v224 + (v225 + (v226 + (v227 + (v228 + (v229 + (v230 + (v231 + (v232 + (v233 + (v234 + (v235 + (v236 + (v237 + (v238 + (v239 + (v240 + (v241 + (v242 + (v243 + (v244 + (v245 + (v246 + (v247 + (v248 + v216))))))))))))))))))))))))))))))
	goto L43
L42:
	;
	v280 = v216
	goto L43
L43:
	;
	goto L40
L44:
	;
	v286 = *(*int32)(unsafe.Add(mBase, _c_F_ExplainPrintPlan[1]))
	v287 = int32(0)
	if base.B2i32(v286 == v287)|base.B2i32(v286 == int32(_a_F_ExplainPrintPlan_0)) == v287 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v296 = v286
	goto L48
L46:
	;
	goto L47
L47:
	;
	m.G0 = v206 + int32(16)
	v415 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v415 != 0 {
		goto L82
	} else {
		goto L83
	}
L48:
	;
	v307 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v296-int32(44)))))
	if v307&int32(32) == int32(0) {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	goto L47
L50:
	;
	v398 = *(*int32)(unsafe.Add(mBase, uint32(v296)+4))
	if v398 != int32(_a_F_ExplainPrintPlan_0) {
		v296 = v398
		goto L48
	} else {
		goto L81
	}
L51:
	;
	v313 = v296 + int32(-64)
	v314 = F_ConfigOptionIsVisible(m, v313)
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	if v314 == int32(0) {
		goto L50
	} else {
		goto L53
	}
L53:
	;
	v319 = v296 - int32(40)
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v319)))
	switch v320 {
	case 0:
		goto L60
	case 1:
		goto L55
	case 2:
		goto L59
	case 3:
		goto L58
	case 4:
		goto L57
	default:
		goto L56
	}
L54:
	;
	v386 = *(*int32)(unsafe.Add(mBase, uint32(v209)))
	*(*int32)(unsafe.Add(mBase, uint32(v283+v386<<(uint(int32(2))%32)))) = v313
	v391 = *(*int32)(unsafe.Add(mBase, uint32(v209)))
	*(*int32)(unsafe.Add(mBase, uint32(v209))) = v391 + int32(1)
	goto L50
L55:
	;
	v380 = *(*int32)(unsafe.Add(mBase, uint32(v296)+32))
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v296)+28))
	v382 = *(*int32)(unsafe.Add(mBase, uint32(v381)))
	if v380 == v382 {
		goto L50
	} else {
		goto L80
	}
L56:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L1
	} else {
		goto L77
	}
L57:
	;
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v296)+32))
	v363 = *(*int32)(unsafe.Add(mBase, uint32(v296)+28))
	v364 = *(*int32)(unsafe.Add(mBase, uint32(v363)))
	if v362 != v364 {
		goto L54
	} else {
		goto L76
	}
L58:
	;
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v296)+28))
	v330 = *(*int32)(unsafe.Add(mBase, uint32(v329)))
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v296)+32))
	if v331 == int32(0) {
		goto L63
	} else {
		goto L64
	}
L59:
	;
	v325 = *(*float64)(unsafe.Add(mBase, uint32(v296)+32))
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v296)+28))
	v327 = *(*float64)(unsafe.Add(mBase, uint32(v326)))
	if base.F64_ne(v325, v327) != 0 {
		goto L54
	} else {
		goto L62
	}
L60:
	;
	v321 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v296)+32)))
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v296)+28))
	v323 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v322))))
	if v321 != v323 {
		goto L54
	} else {
		goto L61
	}
L61:
	;
	goto L50
L62:
	;
	goto L50
L63:
	;
	if v330 != 0 {
		goto L54
	} else {
		goto L66
	}
L64:
	;
	goto L65
L65:
	;
	if v330 == int32(0) {
		goto L54
	} else {
		goto L67
	}
L66:
	;
	goto L50
L67:
	;
	v338 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v331))))
	v341 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v330))))
	if base.B2i32(v338 == int32(0))|base.B2i32(v338 != v341) != 0 {
		v359 = v338
		v360 = v341
		goto L69
	} else {
		goto L70
	}
L68:
	;
	if v359-v360 != 0 {
		goto L54
	} else {
		goto L75
	}
L69:
	;
	goto L68
L70:
	;
	v344 = v331
	v345 = v330
	goto L71
L71:
	;
	v348 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v345)+1)))
	v349 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v344)+1)))
	if v349 == int32(0) {
		v359 = v349
		v360 = v348
		goto L69
	} else {
		goto L73
	}
L72:
	;
	v359 = v349
	v360 = v348
	goto L69
L73:
	;
	v352 = int32(1)
	if v349 == v348 {
		v344 = v344 + v352
		v345 = v345 + v352
		goto L71
	} else {
		goto L74
	}
L74:
	;
	goto L72
L75:
	;
	goto L50
L76:
	;
	goto L50
L77:
	;
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v319)))
	*(*int32)(unsafe.Add(mBase, uint32(v206))) = v370
	F_errmsg_internal(m, int32(_a_F_ExplainPrintPlan_1), v206)
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L1
	} else {
		goto L78
	}
L78:
	;
	F_errfinish(m, int32(_a_F_ExplainPrintPlan_2), int32(_a_F_ExplainPrintPlan_3), int32(_a_F_ExplainPrintPlan_4))
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L1
	} else {
		goto L79
	}
L79:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L80:
	;
	goto L54
L81:
	;
	goto L49
L82:
	;
	v416 = int32(_a_F_ExplainPrintPlan_5)
	F_ExplainOpenGroup(m, v416, v416, int32(1), l0)
	mBase = m.M
	v420 = m.ExcPending
	if v420 != 0 {
		goto L1
	} else {
		goto L85
	}
L83:
	;
	goto L84
L84:
	;
	v467 = *(*int32)(unsafe.Add(mBase, uint32(v14)+76))
	if v467 <= int32(0) {
		goto L38
	} else {
		goto L95
	}
L85:
	;
	v421 = int32(0)
	v422 = *(*int32)(unsafe.Add(mBase, uint32(v14)+76))
	if v421 < v422 {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v427 = v421
	goto L89
L87:
	;
	goto L88
L88:
	;
	F_ExplainCloseGroup(m, int32(_a_F_ExplainPrintPlan_5), int32(1), l0)
	mBase = m.M
	v466 = m.ExcPending
	if v466 != 0 {
		goto L1
	} else {
		goto L94
	}
L89:
	;
	v439 = *(*int32)(unsafe.Add(mBase, uint32(v283+v427<<(uint(int32(2))%32))))
	v440 = *(*int32)(unsafe.Add(mBase, uint32(v439)))
	v443 = F_GetConfigOptionByName(m, v440, int32(0), int32(1))
	mBase = m.M
	v444 = m.ExcPending
	if v444 != 0 {
		goto L1
	} else {
		goto L91
	}
L90:
	;
	goto L88
L91:
	;
	v445 = *(*int32)(unsafe.Add(mBase, uint32(v439)))
	F_ExplainPropertyText(m, v445, v443, l0)
	mBase = m.M
	v447 = m.ExcPending
	if v447 != 0 {
		goto L1
	} else {
		goto L92
	}
L92:
	;
	v449 = v427 + int32(1)
	v450 = *(*int32)(unsafe.Add(mBase, uint32(v14)+76))
	if v449 < v450 {
		v427 = v449
		goto L89
	} else {
		goto L93
	}
L93:
	;
	goto L90
L94:
	;
	goto L38
L95:
	;
	v471 = v14 + int32(60)
	F_initStringInfo(m, v471)
	mBase = m.M
	v473 = m.ExcPending
	if v473 != 0 {
		goto L1
	} else {
		goto L96
	}
L96:
	;
	v474 = *(*int32)(unsafe.Add(mBase, uint32(v14)+76))
	if v474 <= int32(0) {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	v558 = *(*int32)(unsafe.Add(mBase, uint32(v14)+60))
	F_ExplainPropertyText(m, int32(_a_F_ExplainPrintPlan_5), v558, l0)
	mBase = m.M
	v560 = m.ExcPending
	if v560 != 0 {
		goto L1
	} else {
		goto L118
	}
L98:
	;
	v477 = *(*int32)(unsafe.Add(mBase, uint32(v283)))
	v478 = *(*int32)(unsafe.Add(mBase, uint32(v477)))
	v481 = F_GetConfigOptionByName(m, v478, int32(0), int32(1))
	mBase = m.M
	v482 = m.ExcPending
	if v482 != 0 {
		goto L1
	} else {
		goto L99
	}
L99:
	;
	v483 = *(*int32)(unsafe.Add(mBase, uint32(v477)))
	if v481 != 0 {
		goto L101
	} else {
		goto L102
	}
L100:
	;
	v499 = int32(1)
	v500 = *(*int32)(unsafe.Add(mBase, uint32(v14)+76))
	if v500 <= v499 {
		goto L97
	} else {
		goto L106
	}
L101:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+52)) = v481
	*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = v483
	F_appendStringInfo(m, v471, int32(_a_F_ExplainPrintPlan_6), v14+int32(48))
	mBase = m.M
	v490 = m.ExcPending
	if v490 != 0 {
		goto L1
	} else {
		goto L104
	}
L102:
	;
	goto L103
L103:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v483
	F_appendStringInfo(m, v14+int32(60), int32(_a_F_ExplainPrintPlan_7), v14+int32(32))
	mBase = m.M
	v498 = m.ExcPending
	if v498 != 0 {
		goto L1
	} else {
		goto L105
	}
L104:
	;
	goto L100
L105:
	;
	goto L100
L106:
	;
	v505 = v499
	goto L107
L107:
	;
	v517 = *(*int32)(unsafe.Add(mBase, uint32(v283+v505<<(uint(int32(2))%32))))
	v519 = v14 + int32(60)
	F_appendStringInfoString(m, v519, int32(_a_F_ExplainPrintPlan_8))
	mBase = m.M
	v522 = m.ExcPending
	if v522 != 0 {
		goto L1
	} else {
		goto L109
	}
L108:
	;
	goto L97
L109:
	;
	v523 = *(*int32)(unsafe.Add(mBase, uint32(v517)))
	v526 = F_GetConfigOptionByName(m, v523, int32(0), int32(1))
	mBase = m.M
	v527 = m.ExcPending
	if v527 != 0 {
		goto L1
	} else {
		goto L110
	}
L110:
	;
	v528 = *(*int32)(unsafe.Add(mBase, uint32(v517)))
	if v526 != 0 {
		goto L112
	} else {
		goto L113
	}
L111:
	;
	v543 = v505 + int32(1)
	v544 = *(*int32)(unsafe.Add(mBase, uint32(v14)+76))
	if v543 < v544 {
		v505 = v543
		goto L107
	} else {
		goto L117
	}
L112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = v526
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v528
	F_appendStringInfo(m, v519, int32(_a_F_ExplainPrintPlan_6), v14+int32(16))
	mBase = m.M
	v535 = m.ExcPending
	if v535 != 0 {
		goto L1
	} else {
		goto L115
	}
L113:
	;
	goto L114
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v528
	F_appendStringInfo(m, v14+int32(60), int32(_a_F_ExplainPrintPlan_7), v14)
	mBase = m.M
	v541 = m.ExcPending
	if v541 != 0 {
		goto L1
	} else {
		goto L116
	}
L115:
	;
	goto L111
L116:
	;
	goto L111
L117:
	;
	goto L108
L118:
	;
	goto L38
L119:
	;
	m.G0 = v14 + int32(80)
	return
L120:
	;
	v575 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v576 = *(*int64)(unsafe.Add(mBase, uint32(v575)+8))
	if v576 == int64(0) {
		goto L119
	} else {
		goto L121
	}
L121:
	;
	v580 = *(*int32)(unsafe.Add(mBase, _c_F_ExplainPrintPlan[2]))
	if v580 == int32(3) {
		goto L119
	} else {
		goto L122
	}
L122:
	;
	F_ExplainPropertyInteger(m, int32(_a_F_ExplainPrintPlan_9), int32(0), v576, l0)
	mBase = m.M
	v586 = m.ExcPending
	if v586 != 0 {
		goto L1
	} else {
		goto L123
	}
L123:
	;
	goto L119
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
	v12 = int32(_a_F_ExplainPrintTriggers_0)
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
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	if v65 <= int32(0) {
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
	v71 = int32(0)
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
	F_ExplainCloseGroup(m, int32(_a_F_ExplainPrintTriggers_0), int32(0), l0)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L1
	} else {
		goto L29
	}
L23:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	if v96 <= int32(0) {
		goto L22
	} else {
		goto L24
	}
L24:
	;
	v102 = int32(0)
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
	var v31 int32
	_ = v31
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
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
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
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v138 int32
	_ = v138
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
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
	v31 = int32(0)
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch v32 - int32(333) {
	case 0, 6, 7, 8, 9, 11, 12, 13, 21, 22:
		goto L10
	default:
		v103 = v31
		v104 = v4
		v105 = v4
		v106 = v4
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
	v107 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	if v107 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L5:
	;
	v103 = v100
	v104 = int32(1)
	v105 = v99
	v106 = v4
	goto L4
L6:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v21)+84))
	v99 = int32(_a_F_ExplainTargetRel_0)
	v100 = v98
	goto L5
L7:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v21)+108))
	v99 = int32(_a_F_ExplainTargetRel_1)
	v100 = v96
	goto L5
L8:
	;
	v73 = int32(1)
	v75 = int32(_a_F_ExplainTargetRel_2)
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)+4))
	switch v77 {
	case 0:
		v103 = int32(_a_F_ExplainTargetRel_3)
		v104 = v73
		v105 = v75
		v106 = v4
		goto L4
	case 1:
		goto L30
	default:
		goto L29
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
		v103 = v36
		v104 = v38
		v105 = v39
		v106 = v4
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
	v103 = v36
	v104 = v38
	v105 = v39
	v106 = v46
	goto L4
L16:
	;
	v103 = v31
	v104 = int32(1)
	v105 = v48
	v106 = v4
	goto L4
L17:
	;
	goto L18
L18:
	;
	v53 = int32(1)
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v49)+4))
	if v54 != v53 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v103 = v31
	v104 = v53
	v105 = v48
	v106 = v4
	goto L4
L20:
	;
	goto L21
L21:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v49)+12))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)+4))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
	if v60 != int32(15) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v103 = v31
	v104 = v53
	v105 = v48
	v106 = v4
	goto L4
L23:
	;
	goto L24
L24:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v59)+4))
	v64 = F_get_func_name(m, v63)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L11
	} else {
		goto L25
	}
L25:
	;
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)))
	if v66 != int32(1) {
		v103 = v64
		v104 = v53
		v105 = v48
		v106 = v4
		goto L4
	} else {
		goto L26
	}
L26:
	;
	v69 = F_get_func_namespace(m, v63)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L11
	} else {
		goto L27
	}
L27:
	;
	v71 = F_get_namespace_name_or_temp(m, v69)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L11
	} else {
		goto L28
	}
L28:
	;
	v103 = v64
	v104 = v53
	v105 = v48
	v106 = v71
	goto L4
L29:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L11
	} else {
		goto L31
	}
L30:
	;
	v103 = int32(_a_F_ExplainTargetRel_6)
	v104 = v73
	v105 = v75
	v106 = v4
	goto L4
L31:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v76)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = v83
	F_errmsg_internal(m, int32(_a_F_ExplainTargetRel_7), v10+int32(-16))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L11
	} else {
		goto L32
	}
L32:
	;
	F_errfinish(m, int32(_a_F_ExplainTargetRel_8), int32(_a_F_ExplainTargetRel_9), int32(_a_F_ExplainTargetRel_10))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
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
	v110 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	F_appendStringInfoString(m, v110, int32(_a_F_ExplainTargetRel_11))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L11
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	if base.B2i32(v103 != int32(0))&v104 != 0 {
		goto L61
	} else {
		goto L62
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
	v171 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v172 = F_quote_identifier(m, v30)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L11
	} else {
		goto L59
	}
L40:
	;
	v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30))))
	v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103))))
	if base.B2i32(v144 == int32(0))|base.B2i32(v144 != v147) != 0 {
		v165 = v144
		v166 = v147
		goto L52
	} else {
		goto L53
	}
L41:
	;
	if v103 == int32(0) {
		goto L39
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v128 = F_quote_identifier(m, v106)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L11
	} else {
		goto L47
	}
L44:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v119 = F_quote_identifier(m, v103)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L11
	} else {
		goto L45
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v119
	F_appendStringInfo(m, v118, int32(_a_F_ExplainTargetRel_12), v10+int32(-48))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L11
	} else {
		goto L46
	}
L46:
	;
	goto L40
L47:
	;
	v130 = F_quote_identifier(m, v103)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L11
	} else {
		goto L48
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = v130
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v128
	F_appendStringInfo(m, v127, int32(_a_F_ExplainTargetRel_13), v10+int32(-32))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L11
	} else {
		goto L49
	}
L49:
	;
	if v103 == int32(0) {
		goto L39
	} else {
		goto L50
	}
L50:
	;
	goto L40
L51:
	;
	if v165-v166 == int32(0) {
		goto L34
	} else {
		goto L58
	}
L52:
	;
	goto L51
L53:
	;
	v150 = v30
	v151 = v103
	goto L54
L54:
	;
	v154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+1)))
	v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v150)+1)))
	if v155 == int32(0) {
		v165 = v155
		v166 = v154
		goto L52
	} else {
		goto L56
	}
L55:
	;
	v165 = v155
	v166 = v154
	goto L52
L56:
	;
	v158 = int32(1)
	if v155 == v154 {
		v150 = v150 + v158
		v151 = v151 + v158
		goto L54
	} else {
		goto L57
	}
L57:
	;
	goto L55
L58:
	;
	goto L39
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v172
	F_appendStringInfo(m, v171, int32(_a_F_ExplainTargetRel_12), v12)
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L11
	} else {
		goto L60
	}
L60:
	;
	goto L34
L61:
	;
	F_ExplainPropertyText(m, v105, v103, l2)
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L11
	} else {
		goto L64
	}
L62:
	;
	goto L63
L63:
	;
	if v106 != 0 {
		goto L65
	} else {
		goto L66
	}
L64:
	;
	goto L63
L65:
	;
	F_ExplainPropertyText(m, int32(_a_F_ExplainTargetRel_14), v106, l2)
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L11
	} else {
		goto L68
	}
L66:
	;
	goto L67
L67:
	;
	F_ExplainPropertyText(m, int32(_a_F_ExplainTargetRel_15), v30, l2)
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L11
	} else {
		goto L69
	}
L68:
	;
	goto L67
L69:
	;
	goto L34
}
