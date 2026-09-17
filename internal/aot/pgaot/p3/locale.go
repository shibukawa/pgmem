package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_check_locale(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v127 int32
	_ = v127
	var v132 int32
	_ = v132
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v213 int64
	_ = v213
	var v216 int64
	_ = v216
	var v219 int64
	_ = v219
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v254 int32
	_ = v254
	var v258 int64
	_ = v258
	var v261 int64
	_ = v261
	var v264 int64
	_ = v264
	var v267 int32
	_ = v267
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v280 int32
	_ = v280
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v297 int32
	_ = v297
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v321 int32
	_ = v321
	var v326 int32
	_ = v326
	var v330 int32
	_ = v330
	var v338 int32
	_ = v338
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v352 int32
	_ = v352
	var v354 int32
	_ = v354
	var v363 int64
	_ = v363
	var v366 int64
	_ = v366
	var v369 int64
	_ = v369
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v385 int32
	_ = v385
	var v387 int32
	_ = v387
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v404 int32
	_ = v404
	var v408 int64
	_ = v408
	var v411 int64
	_ = v411
	var v414 int64
	_ = v414
	var v417 int32
	_ = v417
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v430 int32
	_ = v430
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v447 int32
	_ = v447
	var v450 int32
	_ = v450
	var v453 int32
	_ = v453
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v463 int32
	_ = v463
	var v466 int32
	_ = v466
	var v468 int32
	_ = v468
	var v471 int32
	_ = v471
	var v476 int32
	_ = v476
	var v480 int32
	_ = v480
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v499 int32
	_ = v499
	var v504 int32
	_ = v504
	var v506 int32
	_ = v506
	var v508 int32
	_ = v508
	var v511 int32
	_ = v511
	var v515 int32
	_ = v515
	var v517 int32
	_ = v517
	var v524 int32
	_ = v524
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v536 int32
	_ = v536
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v544 int32
	_ = v544
	var v550 int32
	_ = v550
	v4 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(48)
	m.G0 = v9
	v12 = l1
	goto L3
L1:
	;
	m.G0 = v9 + int32(48)
	return v550
L2:
	;
	if base.B2i32(v14 == int32(0)) == int32(0) {
		goto L6
	} else {
		goto L7
	}
L3:
	;
	v14 = int32(*(*int8)(unsafe.Add(mBase, uint32(v12))))
	if int32(0) < v14 {
		v12 = v12 + int32(1)
		goto L3
	} else {
		goto L5
	}
L4:
	;
	goto L2
L5:
	;
	goto L4
L6:
	;
	v25 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	if l2 != 0 {
		goto L15
	} else {
		goto L16
	}
L9:
	;
	return int32(0)
L10:
	;
	if v25 == int32(0) {
		v550 = v4
		goto L1
	} else {
		goto L11
	}
L11:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L9
	} else {
		goto L12
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = l1
	F_errmsg(m, int32(_a_F_check_locale_0), v9+int32(32))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L9
	} else {
		goto L13
	}
L13:
	;
	F_errfinish(m, int32(_a_F_check_locale_1), int32(312), int32(_a_F_check_locale_2))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L9
	} else {
		goto L14
	}
L14:
	;
	v550 = v4
	goto L1
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(0)
	goto L17
L16:
	;
	goto L17
L17:
	;
	v47 = int32(0)
	v54 = m.G0
	v56 = v54 - int32(48)
	m.G0 = v56
	if base.Ui32(int32(6)) < base.Ui32(l0) {
		v182 = v47
		goto L19
	} else {
		goto L20
	}
L18:
	;
	if v182 == int32(0) {
		v550 = v4
		goto L1
	} else {
		goto L54
	}
L19:
	;
	m.G0 = v56 + int32(48)
	goto L18
L20:
	;
	if l0 == int32(6) {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v143 = int32(0)
	v144 = int32(_a_F_check_locale_3)
	v149 = v47
	goto L45
L22:
	;
	goto L21
L23:
	;
	goto L24
L24:
	;
	goto L39
L37:
	;
	if v127 != 0 {
		goto L42
	} else {
		goto L43
	}
L39:
	;
	goto L40
L40:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(l0<<(uint(int32(2))%32))+uint32(_c_F_check_locale[0])))
	goto L37
L42:
	;
	v132 = v127 + int32(8)
	goto L44
L43:
	;
	v132 = int32(_a_F_check_locale_4)
	goto L44
L44:
	;
	v182 = v132
	goto L19
L45:
	;
	v152 = *(*int32)(unsafe.Add(mBase, _c_F_check_locale[0]))
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v143<<(uint(int32(2))%32))+uint32(_c_F_check_locale[0])))
	if v155 != 0 {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	v173 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v162))) = uint8(v173)
	if v168 != int32(6) {
		goto L51
	} else {
		goto L52
	}
L47:
	;
	v159 = v155 + int32(8)
	goto L49
L48:
	;
	v159 = int32(_a_F_check_locale_4)
	goto L49
L49:
	;
	v160 = F_strlen(m, v159)
	mBase = m.M
	v161 = F___memcpy(m, v144, v159, v160)
	mBase = m.M
	v162 = v144 + v160
	v163 = int32(59)
	*(*uint8)(unsafe.Add(mBase, uint32(v162))) = uint8(v163)
	v165 = int32(1)
	v168 = v149 + base.B2i32(v155 == v152)
	v170 = v143 + v165
	if v170 != int32(6) {
		v143 = v170
		v144 = v162 + v165
		v149 = v168
		goto L45
	} else {
		goto L50
	}
L50:
	;
	goto L46
L51:
	;
	v178 = int32(_a_F_check_locale_3)
	goto L53
L52:
	;
	v178 = v159
	goto L53
L53:
	;
	v182 = v178
	goto L19
L54:
	;
	v192 = F_pstrdup(m, v182)
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L9
	} else {
		goto L55
	}
L55:
	;
	v194 = int32(0)
	v202 = m.G0
	v204 = v202 - int32(48)
	m.G0 = v204
	if base.Ui32(int32(6)) < base.Ui32(l0) {
		v330 = v194
		goto L57
	} else {
		goto L58
	}
L56:
	;
	v338 = int32(0)
	if base.B2i32(l2 == v194)|base.B2i32(v330 == v338) == v338 {
		goto L92
	} else {
		goto L93
	}
L57:
	;
	m.G0 = v204 + int32(48)
	goto L56
L58:
	;
	if l0 == int32(6) {
		goto L60
	} else {
		goto L61
	}
L59:
	;
	v291 = int32(0)
	v292 = int32(_a_F_check_locale_3)
	v297 = v194
	goto L83
L60:
	;
	if l1 == int32(0) {
		goto L59
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	if l1 != 0 {
		goto L76
	} else {
		goto L77
	}
L63:
	;
	v213 = *(*int64)(unsafe.Add(mBase, _c_F_check_locale[1]))
	*(*int64)(unsafe.Add(mBase, uint32(v204)+16)) = v213
	v216 = *(*int64)(unsafe.Add(mBase, _c_F_check_locale[2]))
	*(*int64)(unsafe.Add(mBase, uint32(v204)+8)) = v216
	v219 = *(*int64)(unsafe.Add(mBase, _c_F_check_locale[3]))
	*(*int64)(unsafe.Add(mBase, uint32(v204))) = v219
	v222 = int32(0)
	v223 = l1
	goto L65
L64:
	;
	v330 = int32(0)
	goto L57
L65:
	;
	v231 = F___strchrnul(m, v223, int32(59))
	mBase = m.M
	v232 = v231 - v223
	if v232 <= int32(23) {
		goto L67
	} else {
		goto L68
	}
L66:
	;
	v258 = *(*int64)(unsafe.Add(mBase, uint32(v204)+40))
	*(*int64)(unsafe.Add(mBase, _c_F_check_locale[4])) = v258
	v261 = *(*int64)(unsafe.Add(mBase, uint32(v204)+32))
	*(*int64)(unsafe.Add(mBase, _c_F_check_locale[5])) = v261
	v264 = *(*int64)(unsafe.Add(mBase, uint32(v204)+24))
	*(*int64)(unsafe.Add(mBase, _c_F_check_locale[0])) = v264
	goto L59
L67:
	;
	v235 = F___memcpy(m, v204, v223, v232)
	mBase = m.M
	v237 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v204+v232))) = uint8(v237)
	v241 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v231))))
	if v241 != 0 {
		goto L70
	} else {
		goto L71
	}
L68:
	;
	v243 = v223
	goto L69
L69:
	;
	v244 = F___get_locale(m, v222, v204)
	mBase = m.M
	if v244 == int32(-1) {
		goto L64
	} else {
		goto L73
	}
L70:
	;
	v242 = v231 + int32(1)
	goto L72
L71:
	;
	v242 = v223
	goto L72
L72:
	;
	v243 = v242
	goto L69
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v204+int32(24)+v222<<(uint(int32(2))%32)))) = v244
	v254 = v222 + int32(1)
	if v254 != int32(6) {
		v222 = v254
		v223 = v243
		goto L65
	} else {
		goto L74
	}
L74:
	;
	goto L66
L75:
	;
	if v276 != 0 {
		goto L80
	} else {
		goto L81
	}
L76:
	;
	v267 = F___get_locale(m, l0, l1)
	mBase = m.M
	if v267 == int32(-1) {
		v330 = v194
		goto L57
	} else {
		goto L79
	}
L77:
	;
	goto L78
L78:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(l0<<(uint(int32(2))%32))+uint32(_c_F_check_locale[0])))
	v276 = v275
	goto L75
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0<<(uint(int32(2))%32))+uint32(_c_F_check_locale[0]))) = v267
	v276 = v267
	goto L75
L80:
	;
	v280 = v276 + int32(8)
	goto L82
L81:
	;
	v280 = int32(_a_F_check_locale_4)
	goto L82
L82:
	;
	v330 = v280
	goto L57
L83:
	;
	v300 = *(*int32)(unsafe.Add(mBase, _c_F_check_locale[0]))
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v291<<(uint(int32(2))%32))+uint32(_c_F_check_locale[0])))
	if v303 != 0 {
		goto L85
	} else {
		goto L86
	}
L84:
	;
	v321 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v310))) = uint8(v321)
	if v316 != int32(6) {
		goto L89
	} else {
		goto L90
	}
L85:
	;
	v307 = v303 + int32(8)
	goto L87
L86:
	;
	v307 = int32(_a_F_check_locale_4)
	goto L87
L87:
	;
	v308 = F_strlen(m, v307)
	mBase = m.M
	v309 = F___memcpy(m, v292, v307, v308)
	mBase = m.M
	v310 = v292 + v308
	v311 = int32(59)
	*(*uint8)(unsafe.Add(mBase, uint32(v310))) = uint8(v311)
	v313 = int32(1)
	v316 = v297 + base.B2i32(v303 == v300)
	v318 = v291 + v313
	if v318 != int32(6) {
		v291 = v318
		v292 = v310 + v313
		v297 = v316
		goto L83
	} else {
		goto L88
	}
L88:
	;
	goto L84
L89:
	;
	v326 = int32(_a_F_check_locale_3)
	goto L91
L90:
	;
	v326 = v307
	goto L91
L91:
	;
	v330 = v326
	goto L57
L92:
	;
	v343 = F_pstrdup(m, v330)
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L9
	} else {
		goto L95
	}
L93:
	;
	goto L94
L94:
	;
	v346 = int32(0)
	v352 = m.G0
	v354 = v352 - int32(48)
	m.G0 = v354
	if base.Ui32(int32(6)) < base.Ui32(l0) {
		v480 = v346
		goto L98
	} else {
		goto L99
	}
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v343
	goto L94
L96:
	;
	v506 = base.B2i32(v330 != int32(0))
	F_pfree(m, v192)
	mBase = m.M
	v508 = m.ExcPending
	if v508 != 0 {
		goto L9
	} else {
		goto L138
	}
L97:
	;
	if v480 != 0 {
		goto L96
	} else {
		goto L133
	}
L98:
	;
	m.G0 = v354 + int32(48)
	goto L97
L99:
	;
	if l0 == int32(6) {
		goto L101
	} else {
		goto L102
	}
L100:
	;
	v441 = int32(0)
	v442 = int32(_a_F_check_locale_3)
	v447 = v346
	goto L124
L101:
	;
	if v192 == int32(0) {
		goto L100
	} else {
		goto L104
	}
L102:
	;
	goto L103
L103:
	;
	if v192 != 0 {
		goto L117
	} else {
		goto L118
	}
L104:
	;
	v363 = *(*int64)(unsafe.Add(mBase, _c_F_check_locale[1]))
	*(*int64)(unsafe.Add(mBase, uint32(v354)+16)) = v363
	v366 = *(*int64)(unsafe.Add(mBase, _c_F_check_locale[2]))
	*(*int64)(unsafe.Add(mBase, uint32(v354)+8)) = v366
	v369 = *(*int64)(unsafe.Add(mBase, _c_F_check_locale[3]))
	*(*int64)(unsafe.Add(mBase, uint32(v354))) = v369
	v372 = int32(0)
	v373 = v192
	goto L106
L105:
	;
	v480 = int32(0)
	goto L98
L106:
	;
	v381 = F___strchrnul(m, v373, int32(59))
	mBase = m.M
	v382 = v381 - v373
	if v382 <= int32(23) {
		goto L108
	} else {
		goto L109
	}
L107:
	;
	v408 = *(*int64)(unsafe.Add(mBase, uint32(v354)+40))
	*(*int64)(unsafe.Add(mBase, _c_F_check_locale[4])) = v408
	v411 = *(*int64)(unsafe.Add(mBase, uint32(v354)+32))
	*(*int64)(unsafe.Add(mBase, _c_F_check_locale[5])) = v411
	v414 = *(*int64)(unsafe.Add(mBase, uint32(v354)+24))
	*(*int64)(unsafe.Add(mBase, _c_F_check_locale[0])) = v414
	goto L100
L108:
	;
	v385 = F___memcpy(m, v354, v373, v382)
	mBase = m.M
	v387 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v354+v382))) = uint8(v387)
	v391 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v381))))
	if v391 != 0 {
		goto L111
	} else {
		goto L112
	}
L109:
	;
	v393 = v373
	goto L110
L110:
	;
	v394 = F___get_locale(m, v372, v354)
	mBase = m.M
	if v394 == int32(-1) {
		goto L105
	} else {
		goto L114
	}
L111:
	;
	v392 = v381 + int32(1)
	goto L113
L112:
	;
	v392 = v373
	goto L113
L113:
	;
	v393 = v392
	goto L110
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v354+int32(24)+v372<<(uint(int32(2))%32)))) = v394
	v404 = v372 + int32(1)
	if v404 != int32(6) {
		v372 = v404
		v373 = v393
		goto L106
	} else {
		goto L115
	}
L115:
	;
	goto L107
L116:
	;
	if v426 != 0 {
		goto L121
	} else {
		goto L122
	}
L117:
	;
	v417 = F___get_locale(m, l0, v192)
	mBase = m.M
	if v417 == int32(-1) {
		v480 = v346
		goto L98
	} else {
		goto L120
	}
L118:
	;
	goto L119
L119:
	;
	v425 = *(*int32)(unsafe.Add(mBase, uint32(l0<<(uint(int32(2))%32))+uint32(_c_F_check_locale[0])))
	v426 = v425
	goto L116
L120:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0<<(uint(int32(2))%32))+uint32(_c_F_check_locale[0]))) = v417
	v426 = v417
	goto L116
L121:
	;
	v430 = v426 + int32(8)
	goto L123
L122:
	;
	v430 = int32(_a_F_check_locale_4)
	goto L123
L123:
	;
	v480 = v430
	goto L98
L124:
	;
	v450 = *(*int32)(unsafe.Add(mBase, _c_F_check_locale[0]))
	v453 = *(*int32)(unsafe.Add(mBase, uint32(v441<<(uint(int32(2))%32))+uint32(_c_F_check_locale[0])))
	if v453 != 0 {
		goto L126
	} else {
		goto L127
	}
L125:
	;
	v471 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v460))) = uint8(v471)
	if v466 != int32(6) {
		goto L130
	} else {
		goto L131
	}
L126:
	;
	v457 = v453 + int32(8)
	goto L128
L127:
	;
	v457 = int32(_a_F_check_locale_4)
	goto L128
L128:
	;
	v458 = F_strlen(m, v457)
	mBase = m.M
	v459 = F___memcpy(m, v442, v457, v458)
	mBase = m.M
	v460 = v442 + v458
	v461 = int32(59)
	*(*uint8)(unsafe.Add(mBase, uint32(v460))) = uint8(v461)
	v463 = int32(1)
	v466 = v447 + base.B2i32(v453 == v450)
	v468 = v441 + v463
	if v468 != int32(6) {
		v441 = v468
		v442 = v460 + v463
		v447 = v466
		goto L124
	} else {
		goto L129
	}
L129:
	;
	goto L125
L130:
	;
	v476 = int32(_a_F_check_locale_3)
	goto L132
L131:
	;
	v476 = v457
	goto L132
L132:
	;
	v480 = v476
	goto L98
L133:
	;
	v490 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v491 = m.ExcPending
	if v491 != 0 {
		goto L9
	} else {
		goto L134
	}
L134:
	;
	if v490 == int32(0) {
		goto L96
	} else {
		goto L135
	}
L135:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v192
	F_errmsg_internal(m, int32(_a_F_check_locale_5), v9+int32(16))
	mBase = m.M
	v499 = m.ExcPending
	if v499 != 0 {
		goto L9
	} else {
		goto L136
	}
L136:
	;
	F_errfinish(m, int32(_a_F_check_locale_1), int32(335), int32(_a_F_check_locale_2))
	mBase = m.M
	v504 = m.ExcPending
	if v504 != 0 {
		goto L9
	} else {
		goto L137
	}
L137:
	;
	goto L96
L138:
	;
	if l2 == int32(0) {
		v550 = v506
		goto L1
	} else {
		goto L139
	}
L139:
	;
	v511 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v511 == int32(0) {
		v550 = v506
		goto L1
	} else {
		goto L140
	}
L140:
	;
	v515 = v511
	goto L142
L141:
	;
	if v517 == int32(0) {
		v550 = v506
		goto L1
	} else {
		goto L145
	}
L142:
	;
	v517 = int32(*(*int8)(unsafe.Add(mBase, uint32(v515))))
	if int32(0) < v517 {
		v515 = v515 + int32(1)
		goto L142
	} else {
		goto L144
	}
L143:
	;
	goto L141
L144:
	;
	goto L143
L145:
	;
	v524 = int32(0)
	v527 = F_errstart(m, int32(19), v524)
	mBase = m.M
	v528 = m.ExcPending
	if v528 != 0 {
		goto L9
	} else {
		goto L146
	}
L146:
	;
	if v527 != 0 {
		goto L147
	} else {
		goto L148
	}
L147:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v531 = m.ExcPending
	if v531 != 0 {
		goto L9
	} else {
		goto L150
	}
L148:
	;
	goto L149
L149:
	;
	v542 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	F_pfree(m, v542)
	mBase = m.M
	v544 = m.ExcPending
	if v544 != 0 {
		goto L9
	} else {
		goto L153
	}
L150:
	;
	v532 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v532
	F_errmsg(m, int32(_a_F_check_locale_0), v9)
	mBase = m.M
	v536 = m.ExcPending
	if v536 != 0 {
		goto L9
	} else {
		goto L151
	}
L151:
	;
	F_errfinish(m, int32(_a_F_check_locale_1), int32(344), int32(_a_F_check_locale_2))
	mBase = m.M
	v541 = m.ExcPending
	if v541 != 0 {
		goto L9
	} else {
		goto L152
	}
L152:
	;
	goto L149
L153:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(0)
	v550 = v524
	goto L1
}
func F_check_locale_time(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v7 = F_check_locale(m, int32(2), v5, int32(0))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
	}
}
