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
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v199 int32
	_ = v199
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v210 int32
	_ = v210
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v226 int32
	_ = v226
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
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
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v276 int32
	_ = v276
	var v280 int32
	_ = v280
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v321 int32
	_ = v321
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v348 int32
	_ = v348
	var v350 int32
	_ = v350
	var v357 int32
	_ = v357
	var v365 int32
	_ = v365
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v381 int32
	_ = v381
	var __phi381 int32
	_ = __phi381
	var v386 int32
	_ = v386
	var __phi386 int32
	_ = __phi386
	var v394 int32
	_ = v394
	var v396 int32
	_ = v396
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v408 int32
	_ = v408
	var v415 int32
	_ = v415
	var v420 int32
	_ = v420
	var v428 int32
	_ = v428
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v436 int32
	_ = v436
	var v443 int32
	_ = v443
	var v446 int32
	_ = v446
	var v450 int32
	_ = v450
	var v455 int32
	_ = v455
	var v459 int32
	_ = v459
	var v462 int32
	_ = v462
	var v466 int32
	_ = v466
	var v471 int32
	_ = v471
	var v475 int32
	_ = v475
	var v478 int32
	_ = v478
	var v482 int32
	_ = v482
	var v487 int32
	_ = v487
	var v495 int32
	_ = v495
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v507 int32
	_ = v507
	var v511 int32
	_ = v511
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v517 int32
	_ = v517
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v545 int32
	_ = v545
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v557 int32
	_ = v557
	var v563 int32
	_ = v563
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v579 int32
	_ = v579
	var v584 int32
	_ = v584
	var v595 int32
	_ = v595
	v6 = int32(0)
	if l3 != int32(1) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v22 = base.B2i32(int32(0) < l1)
	F_check_stack_depth(m)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	if v16 != int32(37) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	return int32(1)
L4:
	;
	return int32(0)
L5:
	;
	if l1 <= int32(0) {
		v552 = l2
		v553 = l3
		v557 = v22
		goto L7
	} else {
		goto L8
	}
L6:
	;
	return v595
L7:
	;
	if v557 != 0 {
		v595 = v6
		goto L6
	} else {
		goto L165
	}
L8:
	;
	if l3 <= int32(0) {
		v552 = l2
		v553 = l3
		v557 = v22
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v31 = l0
	v32 = l1
	v33 = l2
	v34 = l3
	goto L10
L10:
	;
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33))))
	switch v44 - int32(92) {
	case 0:
		goto L23
	case 1, 2:
		goto L20
	case 3:
		goto L21
	default:
		goto L24
	}
L11:
	;
	v552 = v545
	v553 = v543
	v557 = v541
	goto L7
L12:
	;
	v540 = int32(0)
	v541 = base.B2i32(v540 < v528)
	v542 = int32(1)
	v543 = v530 - v542
	v545 = v529 + v542
	if v528 <= v540 {
		v552 = v545
		v553 = v543
		v557 = v541
		goto L7
	} else {
		goto L163
	}
L13:
	;
	v523 = int32(1)
	v527 = v31 + v523
	v528 = v32 - v523
	v529 = v521
	v530 = v522
	goto L12
L14:
	;
	v514 = F_pg_strncoll(m, v260, v365, v31, v32, l4)
	mBase = m.M
	v515 = m.ExcPending
	if v515 != 0 {
		goto L4
	} else {
		goto L158
	}
L15:
	;
	v507 = int32(0)
	if v375 == v507 {
		v595 = v507
		goto L6
	} else {
		goto L156
	}
L16:
	;
	v502 = F_pg_strncoll(m, v33, v495-v33, v31, v32, l4)
	mBase = m.M
	v503 = m.ExcPending
	if v503 != 0 {
		goto L4
	} else {
		goto L155
	}
L17:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v475 = m.ExcPending
	if v475 != 0 {
		goto L4
	} else {
		goto L151
	}
L18:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v459 = m.ExcPending
	if v459 != 0 {
		goto L4
	} else {
		goto L147
	}
L19:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v443 = m.ExcPending
	if v443 != 0 {
		goto L4
	} else {
		goto L143
	}
L20:
	;
	if l4 == int32(0) {
		goto L74
	} else {
		goto L75
	}
L21:
	;
	v185 = v31
	v186 = v32
	goto L68
L22:
	;
	v63 = v31
	v64 = v32
	v65 = v33
	v66 = v34
	goto L32
L23:
	;
	if v34 <= int32(1) {
		goto L19
	} else {
		goto L27
	}
L24:
	;
	if v44 != int32(37) {
		goto L20
	} else {
		goto L25
	}
L25:
	;
	v49 = int32(1)
	if base.Ui32(v34) <= base.Ui32(v49) {
		v595 = v49
		goto L6
	} else {
		goto L26
	}
L26:
	;
	goto L22
L27:
	;
	v55 = v33 + int32(1)
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55))))
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31))))
	if v56 == v57 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v521 = v55
	v522 = v34 - int32(1)
	goto L13
L29:
	;
	goto L30
L30:
	;
	return int32(0)
L31:
	;
	if v64 <= int32(0) {
		goto L50
	} else {
		goto L51
	}
L32:
	;
	v76 = int32(1)
	v77 = v66 - v76
	v79 = v65 + v76
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79))))
	switch v80 - int32(92) {
	case 0:
		goto L34
	case 1, 2:
		v132 = v80
		goto L31
	case 3:
		goto L36
	default:
		goto L37
	}
L33:
	;
	if v77 == int32(1) {
		goto L18
	} else {
		goto L49
	}
L34:
	;
	goto L33
L35:
	;
	if int32(2) < v66 {
		v63 = v113
		v64 = v114
		v65 = v79
		v66 = v77
		goto L32
	} else {
		goto L48
	}
L36:
	;
	if v64 <= int32(0) {
		goto L39
	} else {
		goto L40
	}
L37:
	;
	if v80 == int32(37) {
		v113 = v63
		v114 = v64
		goto L35
	} else {
		goto L38
	}
L38:
	;
	v132 = v80
	goto L31
L39:
	;
	return int32(-1)
L40:
	;
	goto L41
L41:
	;
	v90 = v63
	v91 = v64
	goto L42
L42:
	;
	if base.Ui32(v91) < base.Ui32(int32(2)) {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	v113 = v109
	v114 = v107
	goto L35
L44:
	;
	v113 = v63 + v64
	v114 = int32(0)
	goto L35
L45:
	;
	goto L46
L46:
	;
	v106 = int32(1)
	v107 = v91 - v106
	v109 = v90 + v106
	v110 = int32(*(*int8)(unsafe.Add(mBase, uint32(v109))))
	if v110 < int32(-64) {
		v90 = v109
		v91 = v107
		goto L42
	} else {
		goto L47
	}
L47:
	;
	goto L43
L48:
	;
	v595 = int32(1)
	goto L6
L49:
	;
	v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65)+2)))
	v132 = v131
	goto L31
L50:
	;
	return int32(-1)
L51:
	;
	goto L52
L52:
	;
	v139 = v63
	v140 = v64
	goto L53
L53:
	;
	v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v139))))
	if v132&int32(255) != v152 {
		goto L56
	} else {
		goto L57
	}
L54:
	;
	v595 = v173
	goto L6
L55:
	;
	v160 = v139
	v161 = v140
	goto L63
L56:
	;
	if l4 == int32(0) {
		goto L55
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	v157 = F_UTF8_MatchText(m, v139, v140, v79, v77, l4)
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L4
	} else {
		goto L61
	}
L59:
	;
	v156 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+1)))
	if v156 != 0 {
		goto L55
	} else {
		goto L60
	}
L60:
	;
	goto L58
L61:
	;
	if v157 != 0 {
		v595 = v157
		goto L6
	} else {
		goto L62
	}
L62:
	;
	goto L55
L63:
	;
	v173 = int32(-1)
	if base.Ui32(v161) < base.Ui32(int32(2)) {
		v595 = v173
		goto L6
	} else {
		goto L65
	}
L64:
	;
	if int32(1) < v161 {
		v139 = v179
		v140 = v177
		goto L53
	} else {
		goto L67
	}
L65:
	;
	v176 = int32(1)
	v177 = v161 - v176
	v179 = v160 + v176
	v180 = int32(*(*int8)(unsafe.Add(mBase, uint32(v179))))
	if v180 < int32(-64) {
		v160 = v179
		v161 = v177
		goto L63
	} else {
		goto L66
	}
L66:
	;
	goto L64
L67:
	;
	goto L54
L68:
	;
	v199 = v185 + int32(1)
	if v186 < int32(2) {
		goto L70
	} else {
		goto L71
	}
L69:
	;
	v527 = v199
	v528 = v204
	v529 = v33
	v530 = v34
	goto L12
L70:
	;
	v527 = v199
	v528 = int32(0)
	v529 = v33
	v530 = v34
	goto L12
L71:
	;
	goto L72
L72:
	;
	v204 = v186 - int32(1)
	v205 = int32(*(*int8)(unsafe.Add(mBase, uint32(v199))))
	if v205 < int32(-64) {
		v185 = v199
		v186 = v204
		goto L68
	} else {
		goto L73
	}
L73:
	;
	goto L69
L74:
	;
	v436 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31))))
	if v44 == v436 {
		v521 = v33
		v522 = v34
		goto L13
	} else {
		goto L142
	}
L75:
	;
	v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+1)))
	if v210 != 0 {
		goto L74
	} else {
		goto L76
	}
L76:
	;
	if v34 == int32(0) {
		v495 = v33
		goto L16
	} else {
		goto L77
	}
L77:
	;
	v216 = v34
	v218 = v6
	v220 = v33
	goto L81
L78:
	;
	__phi381 = v32
	__phi386 = v31
	v381 = __phi381
	v386 = __phi386
	goto L122
L79:
	;
	v369 = v33
	v370 = v216
	v374 = v220
	v375 = v6
	v376 = v220 - v33
	goto L78
L80:
	;
	v259 = v257 - v33
	v260 = F_palloc(m, v259)
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L4
	} else {
		goto L92
	}
L81:
	;
	v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v220))))
	switch v226 - int32(92) {
	case 0:
		goto L84
	case 1, 2:
		v242 = v216
		v243 = v218
		v244 = v220
		goto L83
	case 3:
		goto L85
	default:
		goto L86
	}
L82:
	;
	v249 = int32(1)
	if v243&v249 == int32(0) {
		v495 = v246
		goto L16
	} else {
		goto L91
	}
L83:
	;
	v245 = int32(1)
	v246 = v244 + v245
	v248 = v242 - v245
	if v248 != 0 {
		v216 = v248
		v218 = v243
		v220 = v246
		goto L81
	} else {
		goto L90
	}
L84:
	;
	v236 = v216 - int32(1)
	if v236 == int32(0) {
		goto L17
	} else {
		goto L89
	}
L85:
	;
	if v218&int32(1) == int32(0) {
		goto L79
	} else {
		goto L88
	}
L86:
	;
	if v226 != int32(37) {
		v242 = v216
		v243 = v218
		v244 = v220
		goto L83
	} else {
		goto L87
	}
L87:
	;
	goto L85
L88:
	;
	v255 = v216
	v257 = v220
	v258 = v6
	goto L80
L89:
	;
	v239 = int32(1)
	v242 = v236
	v243 = v239
	v244 = v220 + v239
	goto L83
L90:
	;
	goto L82
L91:
	;
	v255 = int32(0)
	v257 = v246
	v258 = v249
	goto L80
L92:
	;
	if base.Ui32(v257) <= base.Ui32(v33) {
		v357 = v260
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v365 = v357 - v260
	if v258 != 0 {
		goto L14
	} else {
		goto L121
	}
L94:
	;
	v264 = v259 & int32(3)
	if v264 == int32(0) {
		goto L96
	} else {
		goto L97
	}
L95:
	;
	if base.Ui32(int32(-4)) < base.Ui32(v33-v257) {
		v357 = v297
		goto L93
	} else {
		goto L105
	}
L96:
	;
	v297 = v260
	v298 = v33
	goto L95
L97:
	;
	goto L98
L98:
	;
	v272 = v260
	v273 = v33
	v276 = v6
	goto L99
L99:
	;
	v280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v273))))
	if v280 != int32(92) {
		goto L101
	} else {
		goto L102
	}
L100:
	;
	v297 = v286
	v298 = v288
	goto L95
L101:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v272))) = uint8(v280)
	v286 = v272 + int32(1)
	goto L103
L102:
	;
	v286 = v272
	goto L103
L103:
	;
	v287 = int32(1)
	v288 = v273 + v287
	v290 = v276 + v287
	if v290 != v264 {
		v272 = v286
		v273 = v288
		v276 = v290
		goto L99
	} else {
		goto L104
	}
L104:
	;
	goto L100
L105:
	;
	v313 = v297
	v314 = v298
	goto L106
L106:
	;
	v321 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v314))))
	if v321 != int32(92) {
		goto L108
	} else {
		goto L109
	}
L107:
	;
	v357 = v348
	goto L93
L108:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v313))) = uint8(v321)
	v327 = v313 + int32(1)
	goto L110
L109:
	;
	v327 = v313
	goto L110
L110:
	;
	v328 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v314)+1)))
	if v328 != int32(92) {
		goto L111
	} else {
		goto L112
	}
L111:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v327))) = uint8(v328)
	v334 = v327 + int32(1)
	goto L113
L112:
	;
	v334 = v327
	goto L113
L113:
	;
	v335 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v314)+2)))
	if v335 != int32(92) {
		goto L114
	} else {
		goto L115
	}
L114:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v334))) = uint8(v335)
	v341 = v334 + int32(1)
	goto L116
L115:
	;
	v341 = v334
	goto L116
L116:
	;
	v342 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v314)+3)))
	if v342 != int32(92) {
		goto L117
	} else {
		goto L118
	}
L117:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v341))) = uint8(v342)
	v348 = v341 + int32(1)
	goto L119
L118:
	;
	v348 = v341
	goto L119
L119:
	;
	v350 = v314 + int32(4)
	if v350 != v257 {
		v313 = v348
		v314 = v350
		goto L106
	} else {
		goto L120
	}
L120:
	;
	goto L107
L121:
	;
	v369 = v260
	v370 = v255
	v374 = v257
	v375 = v260
	v376 = v365
	goto L78
L122:
	;
	v394 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	if v394 != 0 {
		goto L124
	} else {
		goto L125
	}
L124:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L4
	} else {
		goto L127
	}
L125:
	;
	goto L126
L126:
	;
	v398 = F_pg_strncoll(m, v369, v376, v31, v386-v31, l4)
	mBase = m.M
	v399 = m.ExcPending
	if v399 != 0 {
		goto L4
	} else {
		goto L129
	}
L127:
	;
	goto L126
L128:
	;
	if v381 == int32(0) {
		goto L15
	} else {
		goto L135
	}
L129:
	;
	if v398 != 0 {
		goto L128
	} else {
		goto L130
	}
L130:
	;
	v400 = F_UTF8_MatchText(m, v386, v381, v374, v370, l4)
	mBase = m.M
	v401 = m.ExcPending
	if v401 != 0 {
		goto L4
	} else {
		goto L131
	}
L131:
	;
	if v400 != int32(1) {
		goto L128
	} else {
		goto L132
	}
L132:
	;
	if v375 == int32(0) {
		v595 = int32(1)
		goto L6
	} else {
		goto L133
	}
L133:
	;
	F_pfree(m, v375)
	mBase = m.M
	v408 = m.ExcPending
	if v408 != 0 {
		goto L4
	} else {
		goto L134
	}
L134:
	;
	return int32(1)
L135:
	;
	v415 = v381
	v420 = v386
	goto L136
L136:
	;
	v428 = v415 - int32(1)
	if v428 == int32(0) {
		goto L138
	} else {
		goto L139
	}
L137:
	;
	__phi381 = v428
	__phi386 = v432
	v381 = __phi381
	v386 = __phi386
	goto L122
L138:
	;
	__phi381 = v428
	__phi386 = v381 + v386
	v381 = __phi381
	v386 = __phi386
	goto L122
L139:
	;
	goto L140
L140:
	;
	v432 = v420 + int32(1)
	v433 = int32(*(*int8)(unsafe.Add(mBase, uint32(v432))))
	if v433 < int32(-64) {
		v415 = v428
		v420 = v432
		goto L136
	} else {
		goto L141
	}
L141:
	;
	goto L137
L142:
	;
	return int32(0)
L143:
	;
	F_errcode(m, int32(84410498))
	mBase = m.M
	v446 = m.ExcPending
	if v446 != 0 {
		goto L4
	} else {
		goto L144
	}
L144:
	;
	F_errmsg(m, int32(206559), int32(0))
	mBase = m.M
	v450 = m.ExcPending
	if v450 != 0 {
		goto L4
	} else {
		goto L145
	}
L145:
	;
	F_errfinish(m, int32(474619), int32(107), int32(61087))
	mBase = m.M
	v455 = m.ExcPending
	if v455 != 0 {
		goto L4
	} else {
		goto L146
	}
L146:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L147:
	;
	F_errcode(m, int32(84410498))
	mBase = m.M
	v462 = m.ExcPending
	if v462 != 0 {
		goto L4
	} else {
		goto L148
	}
L148:
	;
	F_errmsg(m, int32(206559), int32(0))
	mBase = m.M
	v466 = m.ExcPending
	if v466 != 0 {
		goto L4
	} else {
		goto L149
	}
L149:
	;
	F_errfinish(m, int32(474619), int32(169), int32(61087))
	mBase = m.M
	v471 = m.ExcPending
	if v471 != 0 {
		goto L4
	} else {
		goto L150
	}
L150:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L151:
	;
	F_errcode(m, int32(84410498))
	mBase = m.M
	v478 = m.ExcPending
	if v478 != 0 {
		goto L4
	} else {
		goto L152
	}
L152:
	;
	F_errmsg(m, int32(206559), int32(0))
	mBase = m.M
	v482 = m.ExcPending
	if v482 != 0 {
		goto L4
	} else {
		goto L153
	}
L153:
	;
	F_errfinish(m, int32(474619), int32(237), int32(61087))
	mBase = m.M
	v487 = m.ExcPending
	if v487 != 0 {
		goto L4
	} else {
		goto L154
	}
L154:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L155:
	;
	return base.B2i32(v502 == int32(0))
L156:
	;
	F_pfree(m, v375)
	mBase = m.M
	v511 = m.ExcPending
	if v511 != 0 {
		goto L4
	} else {
		goto L157
	}
L157:
	;
	return int32(0)
L158:
	;
	if v260 != 0 {
		goto L159
	} else {
		goto L160
	}
L159:
	;
	F_pfree(m, v260)
	mBase = m.M
	v517 = m.ExcPending
	if v517 != 0 {
		goto L4
	} else {
		goto L162
	}
L160:
	;
	goto L161
L161:
	;
	return base.B2i32(v514 == int32(0))
L162:
	;
	goto L161
L163:
	;
	if int32(1) < v530 {
		v31 = v527
		v32 = v528
		v33 = v545
		v34 = v543
		goto L10
	} else {
		goto L164
	}
L164:
	;
	goto L11
L165:
	;
	v563 = int32(1)
	if v553 <= int32(0) {
		v595 = v563
		goto L6
	} else {
		goto L166
	}
L166:
	;
	v568 = v552
	v569 = v553
	goto L167
L167:
	;
	v579 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v568))))
	if v579 != int32(37) {
		goto L169
	} else {
		goto L170
	}
L168:
	;
	v595 = v563
	goto L6
L169:
	;
	return int32(-1)
L170:
	;
	goto L171
L171:
	;
	v584 = int32(1)
	if v584 < v569 {
		v568 = v568 + v584
		v569 = v569 - v584
		goto L167
	} else {
		goto L172
	}
L172:
	;
	goto L168
}
func F_utf8_to_euc_jp(m *base.Module, l0 int32) int32 {
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
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_check_encoding_conversion_args(m, v8, v9, v10, int32(6), int32(1))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		v18 = int32(0)
		v24 = F_UtfToLocal(m, v6, v10, v5, int32(4330068), v18, v18, v18, int32(1), base.B2i32(v7 != v18))
		mBase = m.M
		v25 = m.ExcPending
		if v25 != 0 {
			return int32(0)
		} else {
			return v24
		}
	}
}
func F_utf8_to_iso8859_2(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
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
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v125 int32
	_ = v125
	var v133 int32
	_ = v133
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
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
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
	return int32(0)
L2:
	;
	if v14 <= int32(0) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v160 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v156))) = uint8(v160)
	return v152 - v11
L4:
	;
	v152 = v11
	v156 = v10
	goto L3
L5:
	;
	goto L6
L6:
	;
	v23 = v11
	v24 = v14
	v27 = v10
	goto L7
L7:
	;
	v31 = int32(*(*int8)(unsafe.Add(mBase, uint32(v23))))
	if v31 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v152 = v148
	v156 = v147
	goto L3
L9:
	;
	if v9 != 0 {
		v152 = v23
		v156 = v27
		goto L3
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	if v31 < int32(0) {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	F_report_invalid_encoding(m, int32(6), v23, v24)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L14:
	;
	v41 = int32(*(*int8)(unsafe.Add(mBase, uint32(v23))))
	if int32(0) <= v41 {
		goto L19
	} else {
		goto L20
	}
L15:
	;
	v142 = int32(-1)
	v143 = int32(1)
	v144 = v31
	goto L16
L16:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v27))) = uint8(v144)
	v147 = v27 + int32(1)
	v148 = v23 + v143
	v149 = v24 + v142
	if int32(0) < v149 {
		v23 = v148
		v24 = v149
		v27 = v147
		goto L7
	} else {
		goto L68
	}
L17:
	;
	if v65 != int32(2) {
		goto L58
	} else {
		goto L59
	}
L18:
	;
	if v65 <= v24 {
		goto L31
	} else {
		goto L32
	}
L19:
	;
	v65 = int32(1)
	goto L18
L20:
	;
	goto L21
L21:
	;
	v46 = v41 & int32(255)
	if v46&int32(224) == int32(192) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v65 = int32(2)
	goto L18
L23:
	;
	goto L24
L24:
	;
	if v46&int32(240) == int32(224) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v65 = int32(3)
	goto L18
L26:
	;
	goto L27
L27:
	;
	if v46&int32(248) == int32(240) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v63 = int32(4)
	goto L30
L29:
	;
	v63 = int32(1)
	goto L30
L30:
	;
	v65 = v63
	goto L18
L31:
	;
	v67 = int32(0)
	switch v65 - int32(1) {
	case 0:
		goto L38
	case 1:
		goto L39
	case 2:
		goto L40
	case 3:
		goto L41
	default:
		v116 = v67
		goto L35
	}
L32:
	;
	goto L33
L33:
	;
	if v9 != 0 {
		v152 = v23
		v156 = v27
		goto L3
	} else {
		goto L56
	}
L34:
	;
	if v116 != 0 {
		goto L17
	} else {
		goto L55
	}
L35:
	;
	goto L34
L36:
	;
	v116 = base.B2i32(base.Ui32(v108&int32(255)) < base.Ui32(int32(245)))
	goto L35
L37:
	;
	if base.I32_extend8_s(v103) < int32(-62) {
		v116 = v67
		goto L35
	} else {
		goto L54
	}
L38:
	;
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23))))
	v103 = v102
	goto L37
L39:
	;
	v76 = int32(*(*int8)(unsafe.Add(mBase, uint32(v23)+1)))
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23))))
	switch v77 - int32(224) {
	case 0:
		goto L48
	default:
		goto L44
	case 13:
		goto L47
	case 16:
		goto L46
	case 20:
		goto L45
	}
L40:
	;
	v73 = int32(*(*int8)(unsafe.Add(mBase, uint32(v23)+2)))
	if int32(-65) < v73 {
		v116 = v67
		goto L35
	} else {
		goto L43
	}
L41:
	;
	v70 = int32(*(*int8)(unsafe.Add(mBase, uint32(v23)+3)))
	if int32(-65) < v70 {
		v116 = v67
		goto L35
	} else {
		goto L42
	}
L42:
	;
	goto L40
L43:
	;
	goto L39
L44:
	;
	if v76 <= int32(-65) {
		v103 = v77
		goto L37
	} else {
		goto L53
	}
L45:
	;
	if int32(-113) < v76 {
		v116 = v67
		goto L35
	} else {
		goto L52
	}
L46:
	;
	if base.Ui32((v76-int32(-64))&int32(255)) < base.Ui32(int32(208)) {
		v116 = v67
		goto L35
	} else {
		goto L51
	}
L47:
	;
	if int32(-97) < v76 {
		v116 = v67
		goto L35
	} else {
		goto L50
	}
L48:
	;
	v80 = int32(224)
	if base.Ui32(v80) <= base.Ui32((v76-int32(-64))&int32(255)) {
		v108 = v80
		goto L36
	} else {
		goto L49
	}
L49:
	;
	v116 = v67
	goto L35
L50:
	;
	v108 = int32(237)
	goto L36
L51:
	;
	v108 = int32(240)
	goto L36
L52:
	;
	v108 = int32(244)
	goto L36
L53:
	;
	v116 = v67
	goto L35
L54:
	;
	v108 = v103
	goto L36
L55:
	;
	goto L33
L56:
	;
	F_report_invalid_encoding(m, int32(6), v23, v24)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
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
	if v9 != 0 {
		v152 = v23
		v156 = v27
		goto L3
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	if v31&int32(30) != int32(2) {
		goto L63
	} else {
		goto L64
	}
L61:
	;
	F_report_untranslatable_char(m, int32(6), int32(8), v23, v24)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L63:
	;
	if v9 != 0 {
		v152 = v23
		v156 = v27
		goto L3
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
	goto L16
L66:
	;
	F_report_untranslatable_char(m, int32(6), int32(8), v23, v24)
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L68:
	;
	goto L8
}
