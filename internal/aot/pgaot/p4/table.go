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
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v274 int32
	_ = v274
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v289 int32
	_ = v289
	var v292 int32
	_ = v292
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v315 int32
	_ = v315
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v325 int32
	_ = v325
	var v331 int32
	_ = v331
	var v334 int32
	_ = v334
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v351 int32
	_ = v351
	var v357 int32
	_ = v357
	var v363 int32
	_ = v363
	var v369 int32
	_ = v369
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v381 int32
	_ = v381
	var v386 int32
	_ = v386
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v394 int32
	_ = v394
	var v396 int32
	_ = v396
	var v398 int32
	_ = v398
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
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v428 int32
	_ = v428
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v446 int32
	_ = v446
	var v449 int32
	_ = v449
	var v451 int32
	_ = v451
	var v456 int32
	_ = v456
	var v458 int32
	_ = v458
	var v466 int32
	_ = v466
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v476 int32
	_ = v476
	var v481 int32
	_ = v481
	var v485 int32
	_ = v485
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v497 int32
	_ = v497
	var v502 int32
	_ = v502
	var v506 int32
	_ = v506
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v518 int32
	_ = v518
	var v523 int32
	_ = v523
	v2 = int32(0)
	v15 = m.G0
	v17 = v15 + int32(-64)
	m.G0 = v17
	if l0 == v2 {
		v446 = v2
		v449 = v2
		v451 = v2
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v506 = m.ExcPending
	if v506 != 0 {
		goto L13
	} else {
		goto L161
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v485 = m.ExcPending
	if v485 != 0 {
		goto L13
	} else {
		goto L157
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v466 = m.ExcPending
	if v466 != 0 {
		goto L13
	} else {
		goto L153
	}
L4:
	;
	F_list_free(m, v449)
	mBase = m.M
	v456 = m.ExcPending
	if v456 != 0 {
		goto L13
	} else {
		goto L151
	}
L5:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v21 <= int32(0) {
		v446 = v2
		v449 = v2
		v451 = v2
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v29 = v2
	v32 = v2
	v33 = v2
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
	v446 = v428
	v449 = v431
	v451 = v433
	goto L4
L9:
	;
	v438 = v37 + int32(1)
	v439 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v438 < v439 {
		v29 = v428
		v32 = v431
		v33 = v432
		v34 = v433
		v37 = v438
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
	if v32 == v57 {
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
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v32)+4))
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
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v32)+12))
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
	v180 = F_palloc(m, int32(16))
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
	if v33 == v137 {
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
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
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
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v33)+12))
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
	v428 = v29
	v431 = v32
	v432 = v33
	v433 = v34
	goto L9
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v180))) = v54
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v42)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v180)+4)) = v183
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v42)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v180)+8)) = v185
	v187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+16)))
	*(*uint8)(unsafe.Add(mBase, uint32(v180)+12)) = uint8(v187)
	v189 = F_lappend(m, v34, v180)
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L13
	} else {
		goto L64
	}
L64:
	;
	v191 = F_lappend_oid(m, v32, v56)
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L13
	} else {
		goto L65
	}
L65:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v42)+8))
	if v193 != 0 {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v194 = F_lappend_oid(m, v29, v56)
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L13
	} else {
		goto L69
	}
L67:
	;
	v196 = v29
	goto L68
L68:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v42)+12))
	if v197 != 0 {
		goto L70
	} else {
		goto L71
	}
L69:
	;
	v196 = v194
	goto L68
L70:
	;
	v198 = F_lappend_oid(m, v33, v56)
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L13
	} else {
		goto L73
	}
L71:
	;
	v200 = v33
	goto L72
L72:
	;
	if v44&int32(1) == int32(0) {
		v428 = v196
		v431 = v191
		v432 = v200
		v433 = v189
		goto L9
	} else {
		goto L74
	}
L73:
	;
	v200 = v198
	goto L72
L74:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v54)+48))
	v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205)+119)))
	if v206 == int32(112) {
		v428 = v196
		v431 = v191
		v432 = v200
		v433 = v189
		goto L9
	} else {
		goto L75
	}
L75:
	;
	v211 = F_find_all_inheritors(m, v56, int32(4), int32(0))
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L13
	} else {
		goto L76
	}
L76:
	;
	if v211 == int32(0) {
		v428 = v196
		v431 = v191
		v432 = v200
		v433 = v189
		goto L9
	} else {
		goto L77
	}
L77:
	;
	v215 = int32(0)
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v211)+4))
	if v216 <= v215 {
		v428 = v196
		v431 = v191
		v432 = v200
		v433 = v189
		goto L9
	} else {
		goto L78
	}
L78:
	;
	v222 = v54
	v224 = v196
	v227 = v191
	v228 = v200
	v229 = v189
	v231 = v215
	goto L79
L79:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v211)+12))
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v233+v231<<(uint(int32(2))%32))))
	v239 = *(*int32)(unsafe.Add(mBase, _c_F_OpenTableList[0]))
	if v239 != 0 {
		goto L81
	} else {
		goto L82
	}
L80:
	;
	v428 = v414
	v431 = v416
	v432 = v417
	v433 = v418
	goto L9
L81:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L13
	} else {
		goto L84
	}
L82:
	;
	goto L83
L83:
	;
	v242 = int32(0)
	if v227 == v242 {
		goto L87
	} else {
		goto L88
	}
L84:
	;
	goto L83
L85:
	;
	v420 = v231 + int32(1)
	v421 = *(*int32)(unsafe.Add(mBase, uint32(v211)+4))
	if v420 < v421 {
		v222 = v413
		v224 = v414
		v227 = v416
		v228 = v417
		v229 = v418
		v231 = v420
		goto L79
	} else {
		goto L149
	}
L86:
	;
	if v280 != 0 {
		goto L99
	} else {
		goto L100
	}
L87:
	;
	v280 = int32(0)
	goto L86
L88:
	;
	goto L89
L89:
	;
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v227)+4))
	if v248 <= int32(0) {
		v274 = v242
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v280 = v274
	goto L86
L91:
	;
	v251 = int32(0)
	if v251 < v248 {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	v254 = v248
	goto L94
L93:
	;
	v254 = v251
	goto L94
L94:
	;
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v227)+12))
	v257 = int32(0)
	goto L95
L95:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v255+v257<<(uint(int32(2))%32))))
	v266 = base.B2i32(v265 == v237)
	if v265 == v237 {
		v274 = v266
		goto L90
	} else {
		goto L97
	}
L96:
	;
	v274 = v266
	goto L90
L97:
	;
	v268 = v257 + int32(1)
	if v268 != v254 {
		v257 = v268
		goto L95
	} else {
		goto L98
	}
L98:
	;
	goto L96
L99:
	;
	if v237 == v56 {
		v413 = v222
		v414 = v224
		v416 = v227
		v417 = v228
		v418 = v229
		goto L85
	} else {
		goto L102
	}
L100:
	;
	goto L101
L101:
	;
	v388 = F_table_open(m, v237, int32(0))
	mBase = m.M
	v389 = m.ExcPending
	if v389 != 0 {
		goto L13
	} else {
		goto L139
	}
L102:
	;
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v42)+8))
	if v282 != 0 {
		goto L1
	} else {
		goto L103
	}
L103:
	;
	v283 = int32(0)
	if v224 == v283 {
		goto L105
	} else {
		goto L106
	}
L104:
	;
	if v321 != 0 {
		goto L1
	} else {
		goto L117
	}
L105:
	;
	v321 = int32(0)
	goto L104
L106:
	;
	goto L107
L107:
	;
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v224)+4))
	if v289 <= int32(0) {
		v315 = v283
		goto L108
	} else {
		goto L109
	}
L108:
	;
	v321 = v315
	goto L104
L109:
	;
	v292 = int32(0)
	if v292 < v289 {
		goto L110
	} else {
		goto L111
	}
L110:
	;
	v295 = v289
	goto L112
L111:
	;
	v295 = v292
	goto L112
L112:
	;
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v224)+12))
	v298 = int32(0)
	goto L113
L113:
	;
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v296+v298<<(uint(int32(2))%32))))
	v307 = base.B2i32(v306 == v237)
	if v306 == v237 {
		v315 = v307
		goto L108
	} else {
		goto L115
	}
L114:
	;
	v315 = v307
	goto L108
L115:
	;
	v309 = v298 + int32(1)
	if v309 != v295 {
		v298 = v309
		goto L113
	} else {
		goto L116
	}
L116:
	;
	goto L114
L117:
	;
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v42)+12))
	if v322 == int32(0) {
		goto L118
	} else {
		goto L119
	}
L118:
	;
	v325 = int32(0)
	if v228 == v325 {
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
	v369 = m.ExcPending
	if v369 != 0 {
		goto L13
	} else {
		goto L135
	}
L121:
	;
	if v363 == int32(0) {
		v413 = v222
		v414 = v224
		v416 = v227
		v417 = v228
		v418 = v229
		goto L85
	} else {
		goto L134
	}
L122:
	;
	v363 = int32(0)
	goto L121
L123:
	;
	goto L124
L124:
	;
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v228)+4))
	if v331 <= int32(0) {
		v357 = v325
		goto L125
	} else {
		goto L126
	}
L125:
	;
	v363 = v357
	goto L121
L126:
	;
	v334 = int32(0)
	if v334 < v331 {
		goto L127
	} else {
		goto L128
	}
L127:
	;
	v337 = v331
	goto L129
L128:
	;
	v337 = v334
	goto L129
L129:
	;
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v228)+12))
	v340 = int32(0)
	goto L130
L130:
	;
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v338+v340<<(uint(int32(2))%32))))
	v349 = base.B2i32(v348 == v237)
	if v348 == v237 {
		v357 = v349
		goto L125
	} else {
		goto L132
	}
L131:
	;
	v357 = v349
	goto L125
L132:
	;
	v351 = v340 + int32(1)
	if v351 != v337 {
		v340 = v351
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
	v372 = m.ExcPending
	if v372 != 0 {
		goto L13
	} else {
		goto L136
	}
L136:
	;
	v373 = *(*int32)(unsafe.Add(mBase, uint32(v222)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+48)) = v373 + int32(4)
	F_errmsg(m, int32(_a_F_OpenTableList_1), v15+int32(-16))
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L13
	} else {
		goto L137
	}
L137:
	;
	F_errfinish(m, int32(_a_F_OpenTableList_2), int32(1958), int32(_a_F_OpenTableList_3))
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
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
	v391 = F_palloc(m, int32(16))
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L13
	} else {
		goto L140
	}
L140:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v391))) = v388
	v394 = *(*int32)(unsafe.Add(mBase, uint32(v42)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v391)+4)) = v394
	v396 = *(*int32)(unsafe.Add(mBase, uint32(v42)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v391)+8)) = v396
	v398 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+16)))
	*(*uint8)(unsafe.Add(mBase, uint32(v391)+12)) = uint8(v398)
	v400 = F_lappend(m, v229, v391)
	mBase = m.M
	v401 = m.ExcPending
	if v401 != 0 {
		goto L13
	} else {
		goto L141
	}
L141:
	;
	v402 = F_lappend_oid(m, v227, v237)
	mBase = m.M
	v403 = m.ExcPending
	if v403 != 0 {
		goto L13
	} else {
		goto L142
	}
L142:
	;
	v404 = *(*int32)(unsafe.Add(mBase, uint32(v42)+8))
	if v404 != 0 {
		goto L143
	} else {
		goto L144
	}
L143:
	;
	v405 = F_lappend_oid(m, v224, v237)
	mBase = m.M
	v406 = m.ExcPending
	if v406 != 0 {
		goto L13
	} else {
		goto L146
	}
L144:
	;
	v407 = v224
	goto L145
L145:
	;
	v408 = *(*int32)(unsafe.Add(mBase, uint32(v42)+12))
	if v408 == int32(0) {
		v413 = v388
		v414 = v407
		v416 = v402
		v417 = v228
		v418 = v400
		goto L85
	} else {
		goto L147
	}
L146:
	;
	v407 = v405
	goto L145
L147:
	;
	v411 = F_lappend_oid(m, v228, v237)
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L13
	} else {
		goto L148
	}
L148:
	;
	v413 = v388
	v414 = v407
	v416 = v402
	v417 = v411
	v418 = v400
	goto L85
L149:
	;
	goto L80
L150:
	;
	goto L8
L151:
	;
	F_list_free(m, v446)
	mBase = m.M
	v458 = m.ExcPending
	if v458 != 0 {
		goto L13
	} else {
		goto L152
	}
L152:
	;
	m.G0 = v17 - int32(-64)
	return v451
L153:
	;
	F_errcode(m, int32(_a_F_OpenTableList_0))
	mBase = m.M
	v469 = m.ExcPending
	if v469 != 0 {
		goto L13
	} else {
		goto L154
	}
L154:
	;
	v470 = *(*int32)(unsafe.Add(mBase, uint32(v54)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v470 + int32(4)
	F_errmsg(m, int32(_a_F_OpenTableList_4), v17)
	mBase = m.M
	v476 = m.ExcPending
	if v476 != 0 {
		goto L13
	} else {
		goto L155
	}
L155:
	;
	F_errfinish(m, int32(_a_F_OpenTableList_2), int32(1882), int32(_a_F_OpenTableList_3))
	mBase = m.M
	v481 = m.ExcPending
	if v481 != 0 {
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
	v488 = m.ExcPending
	if v488 != 0 {
		goto L13
	} else {
		goto L158
	}
L158:
	;
	v489 = *(*int32)(unsafe.Add(mBase, uint32(v54)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v489 + int32(4)
	F_errmsg(m, int32(_a_F_OpenTableList_1), v15+int32(-48))
	mBase = m.M
	v497 = m.ExcPending
	if v497 != 0 {
		goto L13
	} else {
		goto L159
	}
L159:
	;
	F_errfinish(m, int32(_a_F_OpenTableList_2), int32(1889), int32(_a_F_OpenTableList_3))
	mBase = m.M
	v502 = m.ExcPending
	if v502 != 0 {
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
	v509 = m.ExcPending
	if v509 != 0 {
		goto L13
	} else {
		goto L162
	}
L162:
	;
	v510 = *(*int32)(unsafe.Add(mBase, uint32(v222)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+32)) = v510 + int32(4)
	F_errmsg(m, int32(_a_F_OpenTableList_4), v15+int32(-32))
	mBase = m.M
	v518 = m.ExcPending
	if v518 != 0 {
		goto L13
	} else {
		goto L163
	}
L163:
	;
	F_errfinish(m, int32(_a_F_OpenTableList_2), int32(1946), int32(_a_F_OpenTableList_3))
	mBase = m.M
	v523 = m.ExcPending
	if v523 != 0 {
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
func F_table_am_handler_in(m *base.Module, l0 int32) int64 {
	var v7 int64
	_ = v7
	var v10 int32
	_ = v10
	v7 = Fn14235(m, l0, int32(_a_F_table_am_handler_in_0), int32(370), int32(_a_F_table_am_handler_in_1), int32(_a_F_table_am_handler_in_2), int32(_a_F_table_am_handler_in_3))
	v10 = m.ExcPending
	if v10 != 0 {
		return int64(0)
	} else {
		return v7
	}
}
func F_table_parallelscan_initialize(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int64
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)+36))
	v6 = m.T0[v5].(func(*base.Module, int32, int32) int32)(m, l0, l1)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = v6
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
		if v9 != 0 {
			v66 = int32(1)
		} else {
			v11 = v6 + l1
			v12 = int32(0)
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l2)+24))
			v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+29)))
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
			v21 = *(*int64)(unsafe.Add(mBase, uint32(l2)+4))
			v22 = *(*int32)(unsafe.Add(mBase, uint32(l2)+32))
			v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+28)))
			*(*uint8)(unsafe.Add(mBase, uint32(v11)+16)) = uint8(v23)
			*(*uint16)(unsafe.Add(mBase, uint32(v11)+18)) = uint16(v12)
			*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = v22
			*(*int64)(unsafe.Add(mBase, uint32(v11))) = v21
			*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v20
			*(*uint8)(unsafe.Add(mBase, uint32(v11)+17)) = uint8(v19)
			if v19&int32(1) != 0 {
				v34 = v18
			} else {
				v34 = v12
			}
			if v23 != 0 {
				v35 = v34
			} else {
				v35 = v18
			}
			*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v35
			v37 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
			if v37 == int32(0) {
			} else {
				v41 = v37 << (uint(int32(2)) % 32)
				if v41 == int32(0) {
				} else {
					v46 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
					base.MemoryCopy(m, v11+int32(24), v46, v41)
				}
			}
			if v35 <= int32(0) {
			} else {
				v51 = *(*int32)(unsafe.Add(mBase, uint32(l2)+24))
				v53 = v51 << (uint(int32(2)) % 32)
				if v53 == int32(0) {
				} else {
					v56 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
					v62 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
					base.MemoryCopy(m, v11+v56<<(uint(int32(2))%32)+int32(24), v62, v53)
				}
			}
			v66 = int32(0)
		}
		*(*uint8)(unsafe.Add(mBase, uint32(l1)+13)) = uint8(v66)
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
			v19 = F_MakeSingleTupleTableSlot(m, v18, v17)
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
		v19 = F_MakeSingleTupleTableSlot(m, v18, v17)
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
func F_table_to_xml(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int64
	_ = v10
	var v11 int64
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
	var v27 int64
	_ = v27
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v10 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v11 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	v13 = F_pg_detoast_datum_packed(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int64(0)
	} else {
		v17 = F_text_to_cstring(m, v13)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int64(0)
		} else {
			v20 = v8 + int32(16)
			F_initStringInfo(m, v20)
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int64(0)
			} else {
				v27 = F_DirectFunctionCall1Coll(m, int32(1760), int32(0), v10&int64(4294967295))
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return int64(0)
				} else {
					*(*uint32)(unsafe.Add(mBase, uint32(v8))) = uint32(v27)
					F_appendStringInfo(m, v20, int32(_a_F_table_to_xml_0), v8)
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return int64(0)
					} else {
						v33 = *(*int32)(unsafe.Add(mBase, uint32(v8)+16))
						v35 = F_get_rel_name(m, base.I32_wrap_i64(v10))
						mBase = m.M
						v36 = m.ExcPending
						if v36 != 0 {
							return int64(0)
						} else {
							v40 = F_query_to_xml_internal(m, v33, v35, int32(0), base.B2i32(v11 != int64(0)), v17)
							mBase = m.M
							v41 = m.ExcPending
							if v41 != 0 {
								return int64(0)
							} else {
								v42 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
								v43 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
								v44 = F_cstring_to_text_with_len(m, v42, v43)
								mBase = m.M
								v45 = m.ExcPending
								if v45 != 0 {
									return int64(0)
								} else {
									m.G0 = v8 + int32(32)
									return base.I64_extend_i32_u(v44)
								}
							}
						}
					}
				}
			}
		}
	}
}
