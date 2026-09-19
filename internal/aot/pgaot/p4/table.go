package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_OpenTableList(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v129 int32
	_ = v129
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v169 int32
	_ = v169
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v272 int32
	_ = v272
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v313 int32
	_ = v313
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v323 int32
	_ = v323
	var v329 int32
	_ = v329
	var v332 int32
	_ = v332
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v355 int32
	_ = v355
	var v361 int32
	_ = v361
	var v367 int32
	_ = v367
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v379 int32
	_ = v379
	var v384 int32
	_ = v384
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v392 int32
	_ = v392
	var v394 int32
	_ = v394
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
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
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v424 int32
	_ = v424
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v429 int32
	_ = v429
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v442 int32
	_ = v442
	var v444 int32
	_ = v444
	var v447 int32
	_ = v447
	var v452 int32
	_ = v452
	var v454 int32
	_ = v454
	var v462 int32
	_ = v462
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v472 int32
	_ = v472
	var v477 int32
	_ = v477
	var v481 int32
	_ = v481
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v493 int32
	_ = v493
	var v498 int32
	_ = v498
	var v502 int32
	_ = v502
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v514 int32
	_ = v514
	var v519 int32
	_ = v519
	v2 = int32(0)
	v15 = m.G0
	v17 = v15 + int32(-64)
	m.G0 = v17
	if l0 == v2 {
		v442 = v2
		v444 = v2
		v447 = v2
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v502 = m.ExcPending
	if v502 != 0 {
		goto L13
	} else {
		goto L161
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v481 = m.ExcPending
	if v481 != 0 {
		goto L13
	} else {
		goto L157
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v462 = m.ExcPending
	if v462 != 0 {
		goto L13
	} else {
		goto L153
	}
L4:
	;
	F_list_free(m, v444)
	mBase = m.M
	v452 = m.ExcPending
	if v452 != 0 {
		goto L13
	} else {
		goto L151
	}
L5:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v21 <= int32(0) {
		v442 = v2
		v444 = v2
		v447 = v2
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v29 = v2
	v31 = v2
	v32 = v2
	v34 = v2
	v37 = v2
	goto L7
L7:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v38+v37<<(uint(int32(2))%32))))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+16)))
	v46 = *(*int32)(unsafe.Add(mBase, _c_F_OpenTableList[0]))
	if v46 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v442 = v424
	v444 = v426
	v447 = v429
	goto L4
L9:
	;
	v434 = v37 + int32(1)
	v435 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v434 < v435 {
		v29 = v424
		v31 = v426
		v32 = v427
		v34 = v429
		v37 = v434
		goto L7
	} else {
		goto L150
	}
L10:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	v52 = v43
	goto L12
L12:
	;
	v54 = F_table_openrv(m, v52, int32(4))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L13
	} else {
		goto L15
	}
L13:
	;
	return int32(0)
L14:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
	v52 = v51
	goto L12
L15:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v54)+56))
	v57 = int32(0)
	if v31 == v57 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	if v95 != 0 {
		goto L29
	} else {
		goto L30
	}
L17:
	;
	v95 = int32(0)
	goto L16
L18:
	;
	goto L19
L19:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	if v63 <= int32(0) {
		v89 = v57
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v95 = v89
	goto L16
L21:
	;
	v66 = int32(0)
	if v66 < v63 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v69 = v63
	goto L24
L23:
	;
	v69 = v66
	goto L24
L24:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v31)+12))
	v72 = int32(0)
	goto L25
L25:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v70+v72<<(uint(int32(2))%32))))
	v81 = base.B2i32(v80 == v56)
	if v80 == v56 {
		v89 = v81
		goto L20
	} else {
		goto L27
	}
L26:
	;
	v89 = v81
	goto L20
L27:
	;
	v83 = v72 + int32(1)
	if v83 != v69 {
		v72 = v83
		goto L25
	} else {
		goto L28
	}
L28:
	;
	goto L26
L29:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v42)+8))
	if v96 != 0 {
		goto L3
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	v180 = F_palloc(m, int32(12))
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L13
	} else {
		goto L63
	}
L32:
	;
	v97 = int32(0)
	if v29 == v97 {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	if v135 != 0 {
		goto L3
	} else {
		goto L46
	}
L34:
	;
	v135 = int32(0)
	goto L33
L35:
	;
	goto L36
L36:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
	if v103 <= int32(0) {
		v129 = v97
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v135 = v129
	goto L33
L38:
	;
	v106 = int32(0)
	if v106 < v103 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v109 = v103
	goto L41
L40:
	;
	v109 = v106
	goto L41
L41:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v29)+12))
	v112 = int32(0)
	goto L42
L42:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v110+v112<<(uint(int32(2))%32))))
	v121 = base.B2i32(v120 == v56)
	if v120 == v56 {
		v129 = v121
		goto L37
	} else {
		goto L44
	}
L43:
	;
	v129 = v121
	goto L37
L44:
	;
	v123 = v112 + int32(1)
	if v123 != v109 {
		v112 = v123
		goto L42
	} else {
		goto L45
	}
L45:
	;
	goto L43
L46:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v42)+12))
	if v136 != 0 {
		goto L2
	} else {
		goto L47
	}
L47:
	;
	v137 = int32(0)
	if v32 == v137 {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	if v175 != 0 {
		goto L2
	} else {
		goto L61
	}
L49:
	;
	v175 = int32(0)
	goto L48
L50:
	;
	goto L51
L51:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v32)+4))
	if v143 <= int32(0) {
		v169 = v137
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v175 = v169
	goto L48
L53:
	;
	v146 = int32(0)
	if v146 < v143 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v149 = v143
	goto L56
L55:
	;
	v149 = v146
	goto L56
L56:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v32)+12))
	v152 = int32(0)
	goto L57
L57:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v150+v152<<(uint(int32(2))%32))))
	v161 = base.B2i32(v160 == v56)
	if v160 == v56 {
		v169 = v161
		goto L52
	} else {
		goto L59
	}
L58:
	;
	v169 = v161
	goto L52
L59:
	;
	v163 = v152 + int32(1)
	if v163 != v149 {
		v152 = v163
		goto L57
	} else {
		goto L60
	}
L60:
	;
	goto L58
L61:
	;
	F_relation_close(m, v54, int32(4))
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L13
	} else {
		goto L62
	}
L62:
	;
	v424 = v29
	v426 = v31
	v427 = v32
	v429 = v34
	goto L9
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v180))) = v54
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v42)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v180)+4)) = v183
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v42)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v180)+8)) = v185
	v187 = F_lappend(m, v34, v180)
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L13
	} else {
		goto L64
	}
L64:
	;
	v189 = F_lappend_oid(m, v31, v56)
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L13
	} else {
		goto L65
	}
L65:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v42)+8))
	if v191 != 0 {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v192 = F_lappend_oid(m, v29, v56)
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L13
	} else {
		goto L69
	}
L67:
	;
	v194 = v29
	goto L68
L68:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v42)+12))
	if v195 != 0 {
		goto L70
	} else {
		goto L71
	}
L69:
	;
	v194 = v192
	goto L68
L70:
	;
	v196 = F_lappend_oid(m, v32, v56)
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L13
	} else {
		goto L73
	}
L71:
	;
	v198 = v32
	goto L72
L72:
	;
	if v44&int32(1) == int32(0) {
		v424 = v194
		v426 = v189
		v427 = v198
		v429 = v187
		goto L9
	} else {
		goto L74
	}
L73:
	;
	v198 = v196
	goto L72
L74:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v54)+48))
	v204 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v203)+119)))
	if v204 == int32(112) {
		v424 = v194
		v426 = v189
		v427 = v198
		v429 = v187
		goto L9
	} else {
		goto L75
	}
L75:
	;
	v209 = F_find_all_inheritors(m, v56, int32(4), int32(0))
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L13
	} else {
		goto L76
	}
L76:
	;
	if v209 == int32(0) {
		v424 = v194
		v426 = v189
		v427 = v198
		v429 = v187
		goto L9
	} else {
		goto L77
	}
L77:
	;
	v213 = int32(0)
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v209)+4))
	if v214 <= v213 {
		v424 = v194
		v426 = v189
		v427 = v198
		v429 = v187
		goto L9
	} else {
		goto L78
	}
L78:
	;
	v220 = v54
	v222 = v194
	v224 = v189
	v225 = v198
	v227 = v187
	v228 = v213
	goto L79
L79:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v209)+12))
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v231+v228<<(uint(int32(2))%32))))
	v237 = *(*int32)(unsafe.Add(mBase, _c_F_OpenTableList[0]))
	if v237 != 0 {
		goto L81
	} else {
		goto L82
	}
L80:
	;
	v424 = v410
	v426 = v411
	v427 = v412
	v429 = v413
	goto L9
L81:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L13
	} else {
		goto L84
	}
L82:
	;
	goto L83
L83:
	;
	v240 = int32(0)
	if v224 == v240 {
		goto L87
	} else {
		goto L88
	}
L84:
	;
	goto L83
L85:
	;
	v416 = v228 + int32(1)
	v417 = *(*int32)(unsafe.Add(mBase, uint32(v209)+4))
	if v416 < v417 {
		v220 = v409
		v222 = v410
		v224 = v411
		v225 = v412
		v227 = v413
		v228 = v416
		goto L79
	} else {
		goto L149
	}
L86:
	;
	if v278 != 0 {
		goto L99
	} else {
		goto L100
	}
L87:
	;
	v278 = int32(0)
	goto L86
L88:
	;
	goto L89
L89:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v224)+4))
	if v246 <= int32(0) {
		v272 = v240
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v278 = v272
	goto L86
L91:
	;
	v249 = int32(0)
	if v249 < v246 {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	v252 = v246
	goto L94
L93:
	;
	v252 = v249
	goto L94
L94:
	;
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v224)+12))
	v255 = int32(0)
	goto L95
L95:
	;
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v253+v255<<(uint(int32(2))%32))))
	v264 = base.B2i32(v263 == v235)
	if v263 == v235 {
		v272 = v264
		goto L90
	} else {
		goto L97
	}
L96:
	;
	v272 = v264
	goto L90
L97:
	;
	v266 = v255 + int32(1)
	if v266 != v252 {
		v255 = v266
		goto L95
	} else {
		goto L98
	}
L98:
	;
	goto L96
L99:
	;
	if v235 == v56 {
		v409 = v220
		v410 = v222
		v411 = v224
		v412 = v225
		v413 = v227
		goto L85
	} else {
		goto L102
	}
L100:
	;
	goto L101
L101:
	;
	v386 = F_table_open(m, v235, int32(0))
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L13
	} else {
		goto L139
	}
L102:
	;
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v42)+8))
	if v280 != 0 {
		goto L1
	} else {
		goto L103
	}
L103:
	;
	v281 = int32(0)
	if v222 == v281 {
		goto L105
	} else {
		goto L106
	}
L104:
	;
	if v319 != 0 {
		goto L1
	} else {
		goto L117
	}
L105:
	;
	v319 = int32(0)
	goto L104
L106:
	;
	goto L107
L107:
	;
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v222)+4))
	if v287 <= int32(0) {
		v313 = v281
		goto L108
	} else {
		goto L109
	}
L108:
	;
	v319 = v313
	goto L104
L109:
	;
	v290 = int32(0)
	if v290 < v287 {
		goto L110
	} else {
		goto L111
	}
L110:
	;
	v293 = v287
	goto L112
L111:
	;
	v293 = v290
	goto L112
L112:
	;
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v222)+12))
	v296 = int32(0)
	goto L113
L113:
	;
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v294+v296<<(uint(int32(2))%32))))
	v305 = base.B2i32(v304 == v235)
	if v304 == v235 {
		v313 = v305
		goto L108
	} else {
		goto L115
	}
L114:
	;
	v313 = v305
	goto L108
L115:
	;
	v307 = v296 + int32(1)
	if v307 != v293 {
		v296 = v307
		goto L113
	} else {
		goto L116
	}
L116:
	;
	goto L114
L117:
	;
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v42)+12))
	if v320 == int32(0) {
		goto L118
	} else {
		goto L119
	}
L118:
	;
	v323 = int32(0)
	if v225 == v323 {
		goto L122
	} else {
		goto L123
	}
L119:
	;
	goto L120
L120:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L13
	} else {
		goto L135
	}
L121:
	;
	if v361 == int32(0) {
		v409 = v220
		v410 = v222
		v411 = v224
		v412 = v225
		v413 = v227
		goto L85
	} else {
		goto L134
	}
L122:
	;
	v361 = int32(0)
	goto L121
L123:
	;
	goto L124
L124:
	;
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v225)+4))
	if v329 <= int32(0) {
		v355 = v323
		goto L125
	} else {
		goto L126
	}
L125:
	;
	v361 = v355
	goto L121
L126:
	;
	v332 = int32(0)
	if v332 < v329 {
		goto L127
	} else {
		goto L128
	}
L127:
	;
	v335 = v329
	goto L129
L128:
	;
	v335 = v332
	goto L129
L129:
	;
	v336 = *(*int32)(unsafe.Add(mBase, uint32(v225)+12))
	v338 = int32(0)
	goto L130
L130:
	;
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v336+v338<<(uint(int32(2))%32))))
	v347 = base.B2i32(v346 == v235)
	if v346 == v235 {
		v355 = v347
		goto L125
	} else {
		goto L132
	}
L131:
	;
	v355 = v347
	goto L125
L132:
	;
	v349 = v338 + int32(1)
	if v349 != v335 {
		v338 = v349
		goto L130
	} else {
		goto L133
	}
L133:
	;
	goto L131
L134:
	;
	goto L120
L135:
	;
	F_errcode(m, int32(_a_F_OpenTableList_0))
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
		goto L13
	} else {
		goto L136
	}
L136:
	;
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v220)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+48)) = v371 + int32(4)
	F_errmsg(m, int32(_a_F_OpenTableList_1), v15+int32(-16))
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L13
	} else {
		goto L137
	}
L137:
	;
	F_errfinish(m, int32(_a_F_OpenTableList_2), int32(1775), int32(_a_F_OpenTableList_3))
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L13
	} else {
		goto L138
	}
L138:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L139:
	;
	v389 = F_palloc(m, int32(12))
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L13
	} else {
		goto L140
	}
L140:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v389))) = v386
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v42)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v389)+4)) = v392
	v394 = *(*int32)(unsafe.Add(mBase, uint32(v42)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v389)+8)) = v394
	v396 = F_lappend(m, v227, v389)
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
		goto L13
	} else {
		goto L141
	}
L141:
	;
	v398 = F_lappend_oid(m, v224, v235)
	mBase = m.M
	v399 = m.ExcPending
	if v399 != 0 {
		goto L13
	} else {
		goto L142
	}
L142:
	;
	v400 = *(*int32)(unsafe.Add(mBase, uint32(v42)+8))
	if v400 != 0 {
		goto L143
	} else {
		goto L144
	}
L143:
	;
	v401 = F_lappend_oid(m, v222, v235)
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		goto L13
	} else {
		goto L146
	}
L144:
	;
	v403 = v222
	goto L145
L145:
	;
	v404 = *(*int32)(unsafe.Add(mBase, uint32(v42)+12))
	if v404 == int32(0) {
		v409 = v386
		v410 = v403
		v411 = v398
		v412 = v225
		v413 = v396
		goto L85
	} else {
		goto L147
	}
L146:
	;
	v403 = v401
	goto L145
L147:
	;
	v407 = F_lappend_oid(m, v225, v235)
	mBase = m.M
	v408 = m.ExcPending
	if v408 != 0 {
		goto L13
	} else {
		goto L148
	}
L148:
	;
	v409 = v386
	v410 = v403
	v411 = v398
	v412 = v407
	v413 = v396
	goto L85
L149:
	;
	goto L80
L150:
	;
	goto L8
L151:
	;
	F_list_free(m, v442)
	mBase = m.M
	v454 = m.ExcPending
	if v454 != 0 {
		goto L13
	} else {
		goto L152
	}
L152:
	;
	m.G0 = v17 - int32(-64)
	return v447
L153:
	;
	F_errcode(m, int32(_a_F_OpenTableList_0))
	mBase = m.M
	v465 = m.ExcPending
	if v465 != 0 {
		goto L13
	} else {
		goto L154
	}
L154:
	;
	v466 = *(*int32)(unsafe.Add(mBase, uint32(v54)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v466 + int32(4)
	F_errmsg(m, int32(_a_F_OpenTableList_4), v17)
	mBase = m.M
	v472 = m.ExcPending
	if v472 != 0 {
		goto L13
	} else {
		goto L155
	}
L155:
	;
	F_errfinish(m, int32(_a_F_OpenTableList_2), int32(1700), int32(_a_F_OpenTableList_3))
	mBase = m.M
	v477 = m.ExcPending
	if v477 != 0 {
		goto L13
	} else {
		goto L156
	}
L156:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L157:
	;
	F_errcode(m, int32(_a_F_OpenTableList_0))
	mBase = m.M
	v484 = m.ExcPending
	if v484 != 0 {
		goto L13
	} else {
		goto L158
	}
L158:
	;
	v485 = *(*int32)(unsafe.Add(mBase, uint32(v54)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v485 + int32(4)
	F_errmsg(m, int32(_a_F_OpenTableList_1), v15+int32(-48))
	mBase = m.M
	v493 = m.ExcPending
	if v493 != 0 {
		goto L13
	} else {
		goto L159
	}
L159:
	;
	F_errfinish(m, int32(_a_F_OpenTableList_2), int32(1707), int32(_a_F_OpenTableList_3))
	mBase = m.M
	v498 = m.ExcPending
	if v498 != 0 {
		goto L13
	} else {
		goto L160
	}
L160:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L161:
	;
	F_errcode(m, int32(_a_F_OpenTableList_0))
	mBase = m.M
	v505 = m.ExcPending
	if v505 != 0 {
		goto L13
	} else {
		goto L162
	}
L162:
	;
	v506 = *(*int32)(unsafe.Add(mBase, uint32(v220)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+32)) = v506 + int32(4)
	F_errmsg(m, int32(_a_F_OpenTableList_4), v15+int32(-32))
	mBase = m.M
	v514 = m.ExcPending
	if v514 != 0 {
		goto L13
	} else {
		goto L163
	}
L163:
	;
	F_errfinish(m, int32(_a_F_OpenTableList_2), int32(1763), int32(_a_F_OpenTableList_3))
	mBase = m.M
	v519 = m.ExcPending
	if v519 != 0 {
		goto L13
	} else {
		goto L164
	}
L164:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_fetch_table_list(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int64
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
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
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
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
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
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
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v180 int32
	_ = v180
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v197 int32
	_ = v197
	var v202 int32
	_ = v202
	v9 = m.G0
	v11 = v9 - int32(80)
	m.G0 = v11
	v14 = *(*int32)(unsafe.Add(mBase, _c_F_fetch_table_list[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+56)) = v14
	v17 = *(*int64)(unsafe.Add(mBase, _c_F_fetch_table_list[1]))
	*(*int64)(unsafe.Add(mBase, uint32(v11)+48)) = v17
	v20 = *(*int32)(unsafe.Add(mBase, _c_F_fetch_table_list[2]))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+24))
	v22 = m.T0[v21].(func(*base.Module, int32) int32)(m, l0)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v26 = F_makeStringInfo(m)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	F_initStringInfo(m, v11-int32(-64))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	F_GetPublicationsStr(m, l1, v26, int32(1))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	if int32(_a_F_fetch_table_list_0) <= v22 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v55
	F_appendStringInfo(m, v11-int32(-64), v53, v11+int32(32))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L1
	} else {
		goto L13
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+56)) = int32(22)
	v53 = int32(_a_F_fetch_table_list_1)
	goto L6
L8:
	;
	goto L9
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+56)) = int32(1003)
	v43 = v11 - int32(-64)
	F_appendStringInfoString(m, v43, int32(_a_F_fetch_table_list_2))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	v47 = int32(_a_F_fetch_table_list_3)
	if v22 < int32(_a_F_fetch_table_list_4) {
		v53 = v47
		goto L6
	} else {
		goto L11
	}
L11:
	;
	F_appendStringInfoString(m, v43, int32(_a_F_fetch_table_list_5))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v53 = v47
	goto L6
L13:
	;
	F_free_attrmap(m, v26)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	v65 = int32(0)
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v11)+64))
	if int32(_a_F_fetch_table_list_6) < v22 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v71 = int32(3)
	goto L17
L16:
	;
	v71 = int32(2)
	goto L17
L17:
	;
	v75 = *(*int32)(unsafe.Add(mBase, _c_F_fetch_table_list[2]))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v75)+60))
	v77 = m.T0[v76].(func(*base.Module, int32, int32, int32, int32) int32)(m, l0, v66, v71, v11+int32(48))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v11)+64))
	F_pfree(m, v79)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v77)))
	if v82 == int32(2) {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L1
	} else {
		goto L69
	}
L21:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v77)+16))
	v87 = F_MakeTupleTableSlot(m, v85, int32(_a_F_fetch_table_list_7))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L1
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L1
	} else {
		goto L65
	}
L24:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v77)+12))
	v92 = F_tuplestore_gettupleslot(m, v89, int32(1), int32(0), v87)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	if v92 != 0 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v97 = v65
	goto L29
L27:
	;
	v143 = v65
	goto L28
L28:
	;
	F_ExecDropSingleTupleTableSlot(m, v87)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L1
	} else {
		goto L51
	}
L29:
	;
	v104 = int32(*(*int16)(unsafe.Add(mBase, uint32(v87)+6)))
	if v104 <= int32(0) {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	v143 = v131
	goto L28
L31:
	;
	F_slot_getsomeattrs_int(m, v87, int32(1))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L1
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v87)+16))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v110)))
	v112 = F_text_to_cstring(m, v111)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L1
	} else {
		goto L35
	}
L34:
	;
	goto L33
L35:
	;
	v114 = int32(*(*int16)(unsafe.Add(mBase, uint32(v87)+6)))
	if v114 <= int32(1) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	F_slot_getsomeattrs_int(m, v87, int32(2))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L1
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v87)+16))
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v120)+4))
	v122 = F_text_to_cstring(m, v121)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L1
	} else {
		goto L40
	}
L39:
	;
	goto L38
L40:
	;
	v125 = F_makeRangeVar(m, v112, v122, int32(-1))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	if base.B2i32(v22 < int32(_a_F_fetch_table_list_4)) == int32(0) {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v129 = F_list_member(m, v97, v125)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L1
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	v131 = F_lappend(m, v97, v125)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L1
	} else {
		goto L47
	}
L45:
	;
	if v129 != 0 {
		goto L20
	} else {
		goto L46
	}
L46:
	;
	goto L44
L47:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v87)+8))
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v133)+12))
	m.T0[v134].(func(*base.Module, int32))(m, v87)
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v77)+12))
	v140 = F_tuplestore_gettupleslot(m, v137, int32(1), int32(0), v87)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	if v140 != 0 {
		v97 = v131
		goto L29
	} else {
		goto L50
	}
L50:
	;
	goto L30
L51:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v77)+8))
	if v152 != 0 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	F_pfree(m, v152)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L1
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v77)+12))
	if v155 != 0 {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	goto L54
L56:
	;
	F_tuplestore_end(m, v155)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L1
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v77)+16))
	if v158 != 0 {
		goto L60
	} else {
		goto L61
	}
L59:
	;
	goto L58
L60:
	;
	F_FreeTupleDesc(m, v158)
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L1
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	F_pfree(m, v77)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L1
	} else {
		goto L64
	}
L63:
	;
	goto L62
L64:
	;
	m.G0 = v11 + int32(80)
	return v143
L65:
	;
	F_errcode(m, int32(100663808))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v77)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v174
	F_errmsg(m, int32(_a_F_fetch_table_list_8), v11+int32(16))
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	F_errfinish(m, int32(_a_F_fetch_table_list_9), int32(2277), int32(_a_F_fetch_table_list_10))
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L1
	} else {
		goto L68
	}
L68:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L69:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L1
	} else {
		goto L70
	}
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v122
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v112
	F_errmsg(m, int32(_a_F_fetch_table_list_11), v11)
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L1
	} else {
		goto L71
	}
L71:
	;
	F_errfinish(m, int32(_a_F_fetch_table_list_9), int32(2299), int32(_a_F_fetch_table_list_10))
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_table_am_handler_in(m *base.Module, l0 int32) int32 {
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v7 = Fn13874(m, l0, int32(_a_F_table_am_handler_in_0), int32(370), int32(_a_F_table_am_handler_in_1), int32(_a_F_table_am_handler_in_2), int32(_a_F_table_am_handler_in_3))
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
	}
}
func F_table_parallelscan_initialize(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int64
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)+36))
	v7 = m.T0[v6].(func(*base.Module, int32, int32) int32)(m, l0, l1)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = v7
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
		switch v11 {
		case 0, 5:
			v12 = l1 + v7
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l2)+24))
			v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+29)))
			v21 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
			v22 = *(*int64)(unsafe.Add(mBase, uint32(l2)+4))
			v23 = *(*int32)(unsafe.Add(mBase, uint32(l2)+32))
			v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+28)))
			*(*uint8)(unsafe.Add(mBase, uint32(v12)+16)) = uint8(v24)
			*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v23
			*(*int64)(unsafe.Add(mBase, uint32(v12))) = v22
			*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v21
			*(*uint8)(unsafe.Add(mBase, uint32(v12)+17)) = uint8(v20)
			if v20&int32(1) != 0 {
				v33 = v19
			} else {
				v33 = int32(0)
			}
			if v24 != 0 {
				v34 = v33
			} else {
				v34 = v19
			}
			*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v34
			v36 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
			if v36 == int32(0) {
			} else {
				v40 = v36 << (uint(int32(2)) % 32)
				if v40 == int32(0) {
				} else {
					v45 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
					base.MemoryCopy(m, v12+int32(24), v45, v40)
				}
			}
			if v34 <= int32(0) {
			} else {
				v50 = *(*int32)(unsafe.Add(mBase, uint32(l2)+24))
				v52 = v50 << (uint(int32(2)) % 32)
				if v52 == int32(0) {
				} else {
					v55 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
					v61 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
					base.MemoryCopy(m, v12+v55<<(uint(int32(2))%32)+int32(24), v61, v52)
				}
			}
			v65 = int32(0)
		default:
			v65 = int32(1)
		}
		*(*uint8)(unsafe.Add(mBase, uint32(l1)+13)) = uint8(v65)
		return
	}
}
func F_table_slot_callbacks(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	if v3 != 0 {
		v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)+4))
		v5 = m.T0[v4].(func(*base.Module, int32) int32)(m, l0)
		mBase = m.M
		v8 = m.ExcPending
		if v8 != 0 {
			return int32(0)
		} else {
			return v5
		}
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
		v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+119)))
		if v13 == int32(102) {
			v16 = int32(_a_F_table_slot_callbacks_0)
		} else {
			v16 = int32(_a_F_table_slot_callbacks_1)
		}
		return v16
	}
}
func F_table_slot_create(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
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
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	if v4 != 0 {
		v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)+4))
		v6 = m.T0[v5].(func(*base.Module, int32) int32)(m, l0)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			v17 = v6
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
			v19 = F_MakeTupleTableSlot(m, v18, v17)
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int32(0)
			} else {
				if l1 != 0 {
					v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
					v22 = F_lappend(m, v21, v19)
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l1))) = v22
						return v19
					}
				} else {
					return v19
				}
			}
		}
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
		v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+119)))
		if v13 == int32(102) {
			v16 = int32(_a_F_table_slot_create_0)
		} else {
			v16 = int32(_a_F_table_slot_create_1)
		}
		v17 = v16
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
		v19 = F_MakeTupleTableSlot(m, v18, v17)
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int32(0)
		} else {
			if l1 != 0 {
				v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				v22 = F_lappend(m, v21, v19)
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l1))) = v22
					return v19
				}
			} else {
				return v19
			}
		}
	}
}
func F_table_to_xml(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
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
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v13 = F_pg_detoast_datum_packed(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		v17 = F_text_to_cstring(m, v13)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			v20 = v8 + int32(16)
			F_initStringInfo(m, v20)
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int32(0)
			} else {
				v25 = F_DirectFunctionCall1Coll(m, int32(1547), int32(0), v11)
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v8))) = v25
					F_appendStringInfo(m, v20, int32(_a_F_table_to_xml_0), v8)
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
						return int32(0)
					} else {
						v31 = *(*int32)(unsafe.Add(mBase, uint32(v8)+16))
						v32 = F_get_rel_name(m, v11)
						mBase = m.M
						v33 = m.ExcPending
						if v33 != 0 {
							return int32(0)
						} else {
							v34 = int32(0)
							v37 = F_query_to_xml_internal(m, v31, v32, v34, base.B2i32(v10 != v34), v17)
							mBase = m.M
							v38 = m.ExcPending
							if v38 != 0 {
								return int32(0)
							} else {
								v39 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
								v40 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
								v41 = F_cstring_to_text_with_len(m, v39, v40)
								mBase = m.M
								v42 = m.ExcPending
								if v42 != 0 {
									return int32(0)
								} else {
									m.G0 = v8 + int32(32)
									return v41
								}
							}
						}
					}
				}
			}
		}
	}
}
