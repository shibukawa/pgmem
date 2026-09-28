package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_UTF8_MatchText(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v48 int32
	_ = v48
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v142 int32
	_ = v142
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v216 int32
	_ = v216
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
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
	var v289 int32
	_ = v289
	var v296 int32
	_ = v296
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v317 int32
	_ = v317
	var __phi317 int32
	_ = __phi317
	var v322 int32
	_ = v322
	var __phi322 int32
	_ = __phi322
	var v329 int32
	_ = v329
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v343 int32
	_ = v343
	var v350 int32
	_ = v350
	var v354 int32
	_ = v354
	var v362 int32
	_ = v362
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v379 int32
	_ = v379
	var v382 int32
	_ = v382
	var v386 int32
	_ = v386
	var v391 int32
	_ = v391
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
	var v405 int32
	_ = v405
	var v412 int32
	_ = v412
	var v415 int32
	_ = v415
	var v419 int32
	_ = v419
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v431 int32
	_ = v431
	var v435 int32
	_ = v435
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v441 int32
	_ = v441
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v462 int32
	_ = v462
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v474 int32
	_ = v474
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v491 int32
	_ = v491
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v506 int32
	_ = v506
	var v511 int32
	_ = v511
	var v523 int32
	_ = v523
	v6 = int32(0)
	if l4 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4))))
	v17 = v13 ^ int32(1)
	goto L3
L2:
	;
	v17 = int32(0)
	goto L3
L3:
	;
	if l3 != int32(1) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	F_check_stack_depth(m)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	if v20 != int32(37) {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	return int32(1)
L7:
	;
	return int32(0)
L8:
	;
	v31 = int32(0)
	if base.B2i32(l1 <= v31)|base.B2i32(l3 <= v31) != 0 {
		v469 = l2
		v470 = l3
		v474 = base.B2i32(int32(0) < l1)
		goto L11
	} else {
		goto L12
	}
L9:
	;
	return v523
L10:
	;
	v491 = int32(1)
	if v482 <= int32(0) {
		v523 = v491
		goto L9
	} else {
		goto L152
	}
L11:
	;
	if v474 != 0 {
		v523 = v6
		goto L9
	} else {
		goto L151
	}
L12:
	;
	v36 = l0
	v37 = l1
	v38 = l2
	v39 = l3
	goto L13
L13:
	;
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38))))
	if v48 != int32(95) {
		goto L21
	} else {
		goto L22
	}
L14:
	;
	v469 = v462
	v470 = v460
	v474 = v458
	goto L11
L15:
	;
	v457 = int32(0)
	v458 = base.B2i32(v457 < v446)
	v459 = int32(1)
	v460 = v448 - v459
	v462 = v447 + v459
	if v446 <= v457 {
		v469 = v462
		v470 = v460
		v474 = v458
		goto L11
	} else {
		goto L149
	}
L16:
	;
	v438 = F_pg_strncoll(m, v265, v303, v36, v37, l4)
	mBase = m.M
	v439 = m.ExcPending
	if v439 != 0 {
		goto L7
	} else {
		goto L144
	}
L17:
	;
	v431 = int32(0)
	if v313 == v431 {
		v523 = v431
		goto L9
	} else {
		goto L142
	}
L18:
	;
	if v48 == int32(92) {
		goto L130
	} else {
		goto L131
	}
L19:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L7
	} else {
		goto L124
	}
L20:
	;
	if v17&int32(1) == int32(0) {
		goto L18
	} else {
		goto L69
	}
L21:
	;
	if v48 != int32(37) {
		goto L20
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v173 = v36
	v174 = v37
	goto L63
L24:
	;
	if base.Ui32(v39) < base.Ui32(int32(2)) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	return int32(1)
L26:
	;
	goto L27
L27:
	;
	v57 = v36
	v58 = v37
	v59 = v38
	v60 = v39
	goto L29
L28:
	;
	if v58 <= int32(0) {
		goto L47
	} else {
		goto L48
	}
L29:
	;
	v69 = int32(1)
	v70 = v60 - v69
	v72 = v59 + v69
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+1)))
	switch v73 - int32(92) {
	case 0:
		goto L31
	case 1, 2:
		v123 = v73
		goto L28
	case 3:
		goto L33
	default:
		goto L34
	}
L30:
	;
	if v70 == int32(1) {
		goto L19
	} else {
		goto L46
	}
L31:
	;
	goto L30
L32:
	;
	if int32(2) < v60 {
		v57 = v105
		v58 = v106
		v59 = v72
		v60 = v70
		goto L29
	} else {
		goto L45
	}
L33:
	;
	if v58 <= int32(0) {
		goto L36
	} else {
		goto L37
	}
L34:
	;
	if v73 == int32(37) {
		v105 = v57
		v106 = v58
		goto L32
	} else {
		goto L35
	}
L35:
	;
	v123 = v73
	goto L28
L36:
	;
	return int32(-1)
L37:
	;
	goto L38
L38:
	;
	v84 = v58
	v85 = v57
	goto L39
L39:
	;
	if base.Ui32(v84) < base.Ui32(int32(2)) {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	v105 = v102
	v106 = v99
	goto L32
L41:
	;
	v105 = v57 + v58
	v106 = int32(0)
	goto L32
L42:
	;
	goto L43
L43:
	;
	v98 = int32(1)
	v99 = v84 - v98
	v100 = int32(*(*int8)(unsafe.Add(mBase, uint32(v85)+1)))
	v102 = v85 + v98
	if v100 < int32(-64) {
		v84 = v99
		v85 = v102
		goto L39
	} else {
		goto L44
	}
L44:
	;
	goto L40
L45:
	;
	v523 = int32(1)
	goto L9
L46:
	;
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+2)))
	v123 = v122
	goto L28
L47:
	;
	return int32(-1)
L48:
	;
	goto L49
L49:
	;
	v130 = v57
	v131 = v58
	goto L50
L50:
	;
	v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v130))))
	if (base.B2i32(v142 == v123&int32(255))|v17)&int32(1) != 0 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v147 = F_UTF8_MatchText(m, v130, v131, v72, v70, l4)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L7
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	v151 = v131
	v153 = v130
	goto L57
L55:
	;
	if v147 != 0 {
		v523 = v147
		goto L9
	} else {
		goto L56
	}
L56:
	;
	goto L54
L57:
	;
	if base.Ui32(v151) < base.Ui32(int32(2)) {
		goto L59
	} else {
		goto L60
	}
L58:
	;
	v130 = v170
	v131 = v167
	goto L50
L59:
	;
	return int32(-1)
L60:
	;
	goto L61
L61:
	;
	v166 = int32(1)
	v167 = v151 - v166
	v168 = int32(*(*int8)(unsafe.Add(mBase, uint32(v153)+1)))
	v170 = v153 + v166
	if v168 < int32(-64) {
		v151 = v167
		v153 = v170
		goto L57
	} else {
		goto L62
	}
L62:
	;
	goto L58
L63:
	;
	if base.Ui32(v174) <= base.Ui32(int32(1)) {
		goto L65
	} else {
		goto L66
	}
L64:
	;
	v445 = v195
	v446 = v192
	v447 = v38
	v448 = v39
	goto L15
L65:
	;
	v187 = int32(1)
	v481 = v38 + v187
	v482 = v39 - v187
	goto L10
L66:
	;
	goto L67
L67:
	;
	v191 = int32(1)
	v192 = v174 - v191
	v193 = int32(*(*int8)(unsafe.Add(mBase, uint32(v173)+1)))
	v195 = v173 + v191
	if v193 < int32(-64) {
		v173 = v195
		v174 = v192
		goto L63
	} else {
		goto L68
	}
L68:
	;
	goto L64
L69:
	;
	v206 = v39
	v208 = v48
	v210 = int32(0)
	v211 = v38
	goto L73
L70:
	;
	v371 = F_pg_strncoll(m, v38, v248-v38, v36, v37, l4)
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L7
	} else {
		goto L123
	}
L71:
	;
	__phi317 = v37
	__phi322 = v36
	v317 = __phi317
	v322 = __phi322
	goto L103
L72:
	;
	v265 = F_palloc(m, v263-v38)
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L7
	} else {
		goto L95
	}
L73:
	;
	v216 = v208 & int32(255)
	switch v216 - int32(92) {
	case 0:
		goto L80
	case 1, 2:
		goto L77
	case 3:
		goto L78
	default:
		goto L79
	}
L74:
	;
	if v210 == int32(0) {
		goto L70
	} else {
		goto L94
	}
L75:
	;
	goto L74
L76:
	;
	v256 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v255))))
	v206 = v253
	v208 = v256
	v210 = v254
	v211 = v255
	goto L73
L77:
	;
	v247 = int32(1)
	v248 = v211 + v247
	v250 = v206 - v247
	if v250 == int32(0) {
		goto L75
	} else {
		goto L93
	}
L78:
	;
	if v210 != 0 {
		goto L90
	} else {
		goto L91
	}
L79:
	;
	if v216 != int32(37) {
		goto L77
	} else {
		goto L89
	}
L80:
	;
	if v206 == int32(1) {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L7
	} else {
		goto L84
	}
L82:
	;
	goto L83
L83:
	;
	v237 = int32(2)
	v238 = v211 + v237
	v239 = int32(1)
	v241 = v206 - v237
	if v241 != 0 {
		v253 = v241
		v254 = v239
		v255 = v238
		goto L76
	} else {
		goto L88
	}
L84:
	;
	F_errcode(m, int32(84410498))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L7
	} else {
		goto L85
	}
L85:
	;
	F_errmsg(m, int32(_a_F_UTF8_MatchText_0), int32(0))
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L7
	} else {
		goto L86
	}
L86:
	;
	F_errfinish(m, int32(_a_F_UTF8_MatchText_1), int32(236), int32(_a_F_UTF8_MatchText_2))
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L7
	} else {
		goto L87
	}
L87:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L88:
	;
	v261 = int32(0)
	v262 = v239
	v263 = v238
	goto L72
L89:
	;
	goto L78
L90:
	;
	v261 = v206
	v262 = int32(0)
	v263 = v211
	goto L72
L91:
	;
	goto L92
L92:
	;
	v306 = v38
	v307 = v206
	v312 = v211
	v313 = v6
	v314 = v211 - v38
	goto L71
L93:
	;
	v253 = v250
	v254 = v210
	v255 = v248
	goto L76
L94:
	;
	v261 = int32(0)
	v262 = int32(1)
	v263 = v248
	goto L72
L95:
	;
	if base.Ui32(v38) < base.Ui32(v263) {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v270 = v38
	v273 = v265
	goto L99
L97:
	;
	v296 = v265
	goto L98
L98:
	;
	v303 = v296 - v265
	if v262 != 0 {
		goto L16
	} else {
		goto L102
	}
L99:
	;
	v280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v270))))
	v283 = v270 + base.B2i32(v280 == int32(92))
	v284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v283))))
	*(*uint8)(unsafe.Add(mBase, uint32(v273))) = uint8(v284)
	v286 = int32(1)
	v287 = v273 + v286
	v289 = v283 + v286
	if base.Ui32(v289) < base.Ui32(v263) {
		v270 = v289
		v273 = v287
		goto L99
	} else {
		goto L101
	}
L100:
	;
	v296 = v287
	goto L98
L101:
	;
	goto L100
L102:
	;
	v306 = v265
	v307 = v261
	v312 = v263
	v313 = v265
	v314 = v303
	goto L71
L103:
	;
	v329 = *(*int32)(unsafe.Add(mBase, _c_F_UTF8_MatchText[0]))
	if v329 != 0 {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L7
	} else {
		goto L108
	}
L106:
	;
	goto L107
L107:
	;
	v333 = F_pg_strncoll(m, v306, v314, v36, v322-v36, l4)
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L7
	} else {
		goto L110
	}
L108:
	;
	goto L107
L109:
	;
	if v317 == int32(0) {
		goto L17
	} else {
		goto L116
	}
L110:
	;
	if v333 != 0 {
		goto L109
	} else {
		goto L111
	}
L111:
	;
	v335 = F_UTF8_MatchText(m, v322, v317, v312, v307, l4)
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L7
	} else {
		goto L112
	}
L112:
	;
	if v335 != int32(1) {
		goto L109
	} else {
		goto L113
	}
L113:
	;
	if v313 == int32(0) {
		v523 = int32(1)
		goto L9
	} else {
		goto L114
	}
L114:
	;
	F_pfree(m, v313)
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L7
	} else {
		goto L115
	}
L115:
	;
	return int32(1)
L116:
	;
	v350 = v317
	v354 = v322
	goto L117
L117:
	;
	v362 = v350 - int32(1)
	if v362 == int32(0) {
		goto L119
	} else {
		goto L120
	}
L118:
	;
	__phi317 = v362
	__phi322 = v367
	v317 = __phi317
	v322 = __phi322
	goto L103
L119:
	;
	__phi317 = v362
	__phi322 = v317 + v322
	v317 = __phi317
	v322 = __phi322
	goto L103
L120:
	;
	goto L121
L121:
	;
	v365 = int32(*(*int8)(unsafe.Add(mBase, uint32(v354)+1)))
	v367 = v354 + int32(1)
	if v365 < int32(-64) {
		v350 = v362
		v354 = v367
		goto L117
	} else {
		goto L122
	}
L122:
	;
	goto L118
L123:
	;
	return base.B2i32(v371 == int32(0))
L124:
	;
	F_errcode(m, int32(84410498))
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L7
	} else {
		goto L125
	}
L125:
	;
	F_errmsg(m, int32(_a_F_UTF8_MatchText_0), int32(0))
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L7
	} else {
		goto L126
	}
L126:
	;
	F_errfinish(m, int32(_a_F_UTF8_MatchText_1), int32(168), int32(_a_F_UTF8_MatchText_2))
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L7
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
	v427 = int32(1)
	v445 = v36 + v427
	v446 = v37 - v427
	v447 = v425
	v448 = v426
	goto L15
L129:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L7
	} else {
		goto L138
	}
L130:
	;
	if base.Ui32(v39) <= base.Ui32(int32(1)) {
		goto L129
	} else {
		goto L133
	}
L131:
	;
	goto L132
L132:
	;
	v405 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36))))
	if v48 == v405 {
		v425 = v38
		v426 = v39
		goto L128
	} else {
		goto L137
	}
L133:
	;
	v396 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+1)))
	v397 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36))))
	if v396 == v397 {
		goto L134
	} else {
		goto L135
	}
L134:
	;
	v399 = int32(1)
	v425 = v38 + v399
	v426 = v39 - v399
	goto L128
L135:
	;
	goto L136
L136:
	;
	return int32(0)
L137:
	;
	return int32(0)
L138:
	;
	F_errcode(m, int32(84410498))
	mBase = m.M
	v415 = m.ExcPending
	if v415 != 0 {
		goto L7
	} else {
		goto L139
	}
L139:
	;
	F_errmsg(m, int32(_a_F_UTF8_MatchText_0), int32(0))
	mBase = m.M
	v419 = m.ExcPending
	if v419 != 0 {
		goto L7
	} else {
		goto L140
	}
L140:
	;
	F_errfinish(m, int32(_a_F_UTF8_MatchText_1), int32(356), int32(_a_F_UTF8_MatchText_2))
	mBase = m.M
	v424 = m.ExcPending
	if v424 != 0 {
		goto L7
	} else {
		goto L141
	}
L141:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L142:
	;
	F_pfree(m, v313)
	mBase = m.M
	v435 = m.ExcPending
	if v435 != 0 {
		goto L7
	} else {
		goto L143
	}
L143:
	;
	return int32(0)
L144:
	;
	if v265 != 0 {
		goto L145
	} else {
		goto L146
	}
L145:
	;
	F_pfree(m, v265)
	mBase = m.M
	v441 = m.ExcPending
	if v441 != 0 {
		goto L7
	} else {
		goto L148
	}
L146:
	;
	goto L147
L147:
	;
	return base.B2i32(v438 == int32(0))
L148:
	;
	goto L147
L149:
	;
	if int32(1) < v448 {
		v36 = v445
		v37 = v446
		v38 = v462
		v39 = v460
		goto L13
	} else {
		goto L150
	}
L150:
	;
	goto L14
L151:
	;
	v481 = v469
	v482 = v470
	goto L10
L152:
	;
	v496 = v481
	v497 = v482
	goto L153
L153:
	;
	v506 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v496))))
	if v506 != int32(37) {
		goto L155
	} else {
		goto L156
	}
L154:
	;
	v523 = v491
	goto L9
L155:
	;
	return int32(-1)
L156:
	;
	goto L157
L157:
	;
	v511 = int32(1)
	if v511 < v497 {
		v496 = v496 + v511
		v497 = v497 - v511
		goto L153
	} else {
		goto L158
	}
L158:
	;
	goto L154
}
func F_utf8_to_euc_jp(m *base.Module, l0 int32) int64 {
	var v3 int32
	_ = v3
	var v7 int64
	_ = v7
	var v10 int32
	_ = v10
	v3 = int32(0)
	v7 = Fn14398(m, l0, int32(1), v3, v3, v3, int32(_a_F_utf8_to_euc_jp_0))
	v10 = m.ExcPending
	if v10 != 0 {
		return int64(0)
	} else {
		return v7
	}
}
func F_utf8_to_iso8859_2(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v9 int64
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v136 int32
	_ = v136
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	var v168 int32
	_ = v168
	v9 = *(*int64)(unsafe.Add(mBase, uint32(l0)+104))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	F_check_encoding_conversion_args(m, v12, v13, v14, int32(6), int32(8))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	if v14 <= int32(0) {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	F_report_untranslatable_char(m, int32(6), int32(8), v23, v24)
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L1
	} else {
		goto L68
	}
L4:
	;
	v160 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v155))) = uint8(v160)
	return base.I64_extend_i32_s(v152 - v11)
L5:
	;
	v152 = v11
	v155 = v10
	goto L4
L6:
	;
	goto L7
L7:
	;
	v23 = v11
	v24 = v14
	v26 = v10
	goto L8
L8:
	;
	v31 = int32(*(*int8)(unsafe.Add(mBase, uint32(v23))))
	if v31 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v152 = v148
	v155 = v147
	goto L4
L10:
	;
	if v9 != int64(0) {
		v152 = v23
		v155 = v26
		goto L4
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	if v31 < int32(0) {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	F_report_invalid_encoding(m, int32(6), v23, v24)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L15:
	;
	v43 = int32(*(*int8)(unsafe.Add(mBase, uint32(v23))))
	if int32(0) <= v43 {
		goto L20
	} else {
		goto L21
	}
L16:
	;
	v142 = int32(-1)
	v143 = int32(1)
	v144 = v31
	goto L17
L17:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v26))) = uint8(v144)
	v147 = v26 + int32(1)
	v148 = v23 + v143
	v149 = v24 + v142
	if int32(0) < v149 {
		v23 = v148
		v24 = v149
		v26 = v147
		goto L8
	} else {
		goto L67
	}
L18:
	;
	if v67 != int32(2) {
		goto L59
	} else {
		goto L60
	}
L19:
	;
	if v67 <= v24 {
		goto L32
	} else {
		goto L33
	}
L20:
	;
	v67 = int32(1)
	goto L19
L21:
	;
	goto L22
L22:
	;
	v48 = v43 & int32(255)
	if v48&int32(224) == int32(192) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v67 = int32(2)
	goto L19
L24:
	;
	goto L25
L25:
	;
	if v48&int32(240) == int32(224) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v67 = int32(3)
	goto L19
L27:
	;
	goto L28
L28:
	;
	if v48&int32(248) == int32(240) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v65 = int32(4)
	goto L31
L30:
	;
	v65 = int32(1)
	goto L31
L31:
	;
	v67 = v65
	goto L19
L32:
	;
	v69 = int32(0)
	switch v67 - int32(1) {
	case 0:
		goto L39
	case 1:
		goto L40
	case 2:
		goto L41
	case 3:
		goto L42
	default:
		v118 = v69
		goto L36
	}
L33:
	;
	goto L34
L34:
	;
	if v9 != int64(0) {
		v152 = v23
		v155 = v26
		goto L4
	} else {
		goto L57
	}
L35:
	;
	if v118 != 0 {
		goto L18
	} else {
		goto L56
	}
L36:
	;
	goto L35
L37:
	;
	v118 = base.B2i32(base.Ui32(v110&int32(255)) < base.Ui32(int32(245)))
	goto L36
L38:
	;
	if base.I32_extend8_s(v105) < int32(-62) {
		v118 = v69
		goto L36
	} else {
		goto L55
	}
L39:
	;
	v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23))))
	v105 = v104
	goto L38
L40:
	;
	v78 = int32(*(*int8)(unsafe.Add(mBase, uint32(v23)+1)))
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23))))
	switch v79 - int32(224) {
	case 0:
		goto L49
	default:
		goto L45
	case 13:
		goto L48
	case 16:
		goto L47
	case 20:
		goto L46
	}
L41:
	;
	v75 = int32(*(*int8)(unsafe.Add(mBase, uint32(v23)+2)))
	if int32(-65) < v75 {
		v118 = v69
		goto L36
	} else {
		goto L44
	}
L42:
	;
	v72 = int32(*(*int8)(unsafe.Add(mBase, uint32(v23)+3)))
	if int32(-65) < v72 {
		v118 = v69
		goto L36
	} else {
		goto L43
	}
L43:
	;
	goto L41
L44:
	;
	goto L40
L45:
	;
	if v78 <= int32(-65) {
		v105 = v79
		goto L38
	} else {
		goto L54
	}
L46:
	;
	if int32(-113) < v78 {
		v118 = v69
		goto L36
	} else {
		goto L53
	}
L47:
	;
	if base.Ui32((v78-int32(-64))&int32(255)) < base.Ui32(int32(208)) {
		v118 = v69
		goto L36
	} else {
		goto L52
	}
L48:
	;
	if int32(-97) < v78 {
		v118 = v69
		goto L36
	} else {
		goto L51
	}
L49:
	;
	v82 = int32(224)
	if base.Ui32(v82) <= base.Ui32((v78-int32(-64))&int32(255)) {
		v110 = v82
		goto L37
	} else {
		goto L50
	}
L50:
	;
	v118 = v69
	goto L36
L51:
	;
	v110 = int32(237)
	goto L37
L52:
	;
	v110 = int32(240)
	goto L37
L53:
	;
	v110 = int32(244)
	goto L37
L54:
	;
	v118 = v69
	goto L36
L55:
	;
	v110 = v105
	goto L37
L56:
	;
	goto L34
L57:
	;
	F_report_invalid_encoding(m, int32(6), v23, v24)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L59:
	;
	if v9 != int64(0) {
		v152 = v23
		v155 = v26
		goto L4
	} else {
		goto L62
	}
L60:
	;
	goto L61
L61:
	;
	if v31&int32(30) != int32(2) {
		goto L63
	} else {
		goto L64
	}
L62:
	;
	goto L3
L63:
	;
	if v9 != int64(0) {
		v152 = v23
		v155 = v26
		goto L4
	} else {
		goto L66
	}
L64:
	;
	goto L65
L65:
	;
	v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+1)))
	v142 = int32(-2)
	v143 = int32(2)
	v144 = v136&int32(63) | v31<<(uint(int32(6))%32)
	goto L17
L66:
	;
	goto L3
L67:
	;
	goto L9
L68:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
