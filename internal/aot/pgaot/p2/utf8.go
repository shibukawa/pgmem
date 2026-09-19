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
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v213 int32
	_ = v213
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v230 int32
	_ = v230
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v284 int32
	_ = v284
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v325 int32
	_ = v325
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v352 int32
	_ = v352
	var v354 int32
	_ = v354
	var v362 int32
	_ = v362
	var v369 int32
	_ = v369
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v385 int32
	_ = v385
	var __phi385 int32
	_ = __phi385
	var v391 int32
	_ = v391
	var __phi391 int32
	_ = __phi391
	var v398 int32
	_ = v398
	var v400 int32
	_ = v400
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v412 int32
	_ = v412
	var v419 int32
	_ = v419
	var v423 int32
	_ = v423
	var v432 int32
	_ = v432
	var v435 int32
	_ = v435
	var v437 int32
	_ = v437
	var v440 int32
	_ = v440
	var v447 int32
	_ = v447
	var v450 int32
	_ = v450
	var v454 int32
	_ = v454
	var v459 int32
	_ = v459
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
	var v482 int32
	_ = v482
	var v486 int32
	_ = v486
	var v491 int32
	_ = v491
	var v500 int32
	_ = v500
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v511 int32
	_ = v511
	var v515 int32
	_ = v515
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v521 int32
	_ = v521
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v536 int32
	_ = v536
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v550 int32
	_ = v550
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v561 int32
	_ = v561
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v581 int32
	_ = v581
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v597 int32
	_ = v597
	var v602 int32
	_ = v602
	var v615 int32
	_ = v615
	v6 = int32(0)
	if l3 != int32(1) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
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
	v27 = int32(0)
	if base.B2i32(l1 <= v27)|base.B2i32(l3 <= v27) != 0 {
		v557 = l2
		v558 = l3
		v561 = base.B2i32(int32(0) < l1)
		goto L8
	} else {
		goto L9
	}
L6:
	;
	return v615
L7:
	;
	v581 = int32(1)
	if v571 <= int32(0) {
		v615 = v581
		goto L6
	} else {
		goto L167
	}
L8:
	;
	if v561 != 0 {
		v615 = v6
		goto L6
	} else {
		goto L166
	}
L9:
	;
	v32 = l0
	v33 = l1
	v34 = l2
	v35 = l3
	goto L10
L10:
	;
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34))))
	switch v45 - int32(92) {
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
	v557 = v550
	v558 = v548
	v561 = v546
	goto L8
L12:
	;
	v545 = int32(0)
	v546 = base.B2i32(v545 < v532)
	v547 = int32(1)
	v548 = v536 - v547
	v550 = v533 + v547
	if v532 <= v545 {
		v557 = v550
		v558 = v548
		v561 = v546
		goto L8
	} else {
		goto L164
	}
L13:
	;
	v527 = int32(1)
	v532 = v33 - v527
	v533 = v525
	v536 = v526
	v544 = v32 + v527
	goto L12
L14:
	;
	v518 = F_pg_strncoll(m, v264, v369, v32, v33, l4)
	mBase = m.M
	v519 = m.ExcPending
	if v519 != 0 {
		goto L4
	} else {
		goto L159
	}
L15:
	;
	v511 = int32(0)
	if v380 == v511 {
		v615 = v511
		goto L6
	} else {
		goto L157
	}
L16:
	;
	v506 = F_pg_strncoll(m, v34, v500-v34, v32, v33, l4)
	mBase = m.M
	v507 = m.ExcPending
	if v507 != 0 {
		goto L4
	} else {
		goto L156
	}
L17:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v479 = m.ExcPending
	if v479 != 0 {
		goto L4
	} else {
		goto L152
	}
L18:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v463 = m.ExcPending
	if v463 != 0 {
		goto L4
	} else {
		goto L148
	}
L19:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v447 = m.ExcPending
	if v447 != 0 {
		goto L4
	} else {
		goto L144
	}
L20:
	;
	if l4 == int32(0) {
		goto L75
	} else {
		goto L76
	}
L21:
	;
	v185 = v32
	v186 = v33
	goto L69
L22:
	;
	v64 = v32
	v65 = v33
	v66 = v34
	v67 = v35
	goto L32
L23:
	;
	if v35 <= int32(1) {
		goto L19
	} else {
		goto L27
	}
L24:
	;
	if v45 != int32(37) {
		goto L20
	} else {
		goto L25
	}
L25:
	;
	v50 = int32(1)
	if base.Ui32(v35) <= base.Ui32(v50) {
		v615 = v50
		goto L6
	} else {
		goto L26
	}
L26:
	;
	goto L22
L27:
	;
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+1)))
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32))))
	if v55 == v56 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v58 = int32(1)
	v525 = v34 + v58
	v526 = v35 - v58
	goto L13
L29:
	;
	goto L30
L30:
	;
	return int32(0)
L31:
	;
	if v65 <= int32(0) {
		goto L50
	} else {
		goto L51
	}
L32:
	;
	v77 = int32(1)
	v78 = v67 - v77
	v80 = v66 + v77
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+1)))
	switch v81 - int32(92) {
	case 0:
		goto L34
	case 1, 2:
		v133 = v81
		goto L31
	case 3:
		goto L36
	default:
		goto L37
	}
L33:
	;
	if v78 == int32(1) {
		goto L18
	} else {
		goto L49
	}
L34:
	;
	goto L33
L35:
	;
	if int32(2) < v67 {
		v64 = v114
		v65 = v115
		v66 = v80
		v67 = v78
		goto L32
	} else {
		goto L48
	}
L36:
	;
	if v65 <= int32(0) {
		goto L39
	} else {
		goto L40
	}
L37:
	;
	if v81 == int32(37) {
		v114 = v64
		v115 = v65
		goto L35
	} else {
		goto L38
	}
L38:
	;
	v133 = v81
	goto L31
L39:
	;
	return int32(-1)
L40:
	;
	goto L41
L41:
	;
	v92 = v65
	v93 = v64
	goto L42
L42:
	;
	if base.Ui32(v92) < base.Ui32(int32(2)) {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	v114 = v111
	v115 = v108
	goto L35
L44:
	;
	v114 = v64 + v65
	v115 = int32(0)
	goto L35
L45:
	;
	goto L46
L46:
	;
	v107 = int32(1)
	v108 = v92 - v107
	v109 = int32(*(*int8)(unsafe.Add(mBase, uint32(v93)+1)))
	v111 = v93 + v107
	if v109 < int32(-64) {
		v92 = v108
		v93 = v111
		goto L42
	} else {
		goto L47
	}
L47:
	;
	goto L43
L48:
	;
	v615 = int32(1)
	goto L6
L49:
	;
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+2)))
	v133 = v132
	goto L31
L50:
	;
	return int32(-1)
L51:
	;
	goto L52
L52:
	;
	v140 = v64
	v141 = v65
	goto L53
L53:
	;
	v153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v140))))
	if v133&int32(255) != v153 {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	v162 = v141
	v164 = v140
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
	v158 = F_UTF8_MatchText(m, v140, v141, v80, v78, l4)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L4
	} else {
		goto L61
	}
L59:
	;
	v157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+1)))
	if v157 != 0 {
		goto L55
	} else {
		goto L60
	}
L60:
	;
	goto L58
L61:
	;
	if v158 != 0 {
		v615 = v158
		goto L6
	} else {
		goto L62
	}
L62:
	;
	goto L55
L63:
	;
	if base.Ui32(v162) < base.Ui32(int32(2)) {
		goto L65
	} else {
		goto L66
	}
L64:
	;
	v140 = v182
	v141 = v179
	goto L53
L65:
	;
	return int32(-1)
L66:
	;
	goto L67
L67:
	;
	v178 = int32(1)
	v179 = v162 - v178
	v180 = int32(*(*int8)(unsafe.Add(mBase, uint32(v164)+1)))
	v182 = v164 + v178
	if v180 < int32(-64) {
		v162 = v179
		v164 = v182
		goto L63
	} else {
		goto L68
	}
L68:
	;
	goto L64
L69:
	;
	if base.Ui32(v186) <= base.Ui32(int32(1)) {
		goto L71
	} else {
		goto L72
	}
L70:
	;
	v532 = v205
	v533 = v34
	v536 = v35
	v544 = v208
	goto L12
L71:
	;
	v200 = int32(1)
	v570 = v34 + v200
	v571 = v35 - v200
	goto L7
L72:
	;
	goto L73
L73:
	;
	v204 = int32(1)
	v205 = v186 - v204
	v206 = int32(*(*int8)(unsafe.Add(mBase, uint32(v185)+1)))
	v208 = v185 + v204
	if v206 < int32(-64) {
		v185 = v208
		v186 = v205
		goto L69
	} else {
		goto L74
	}
L74:
	;
	goto L70
L75:
	;
	v440 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32))))
	if v45 == v440 {
		v525 = v34
		v526 = v35
		goto L13
	} else {
		goto L143
	}
L76:
	;
	v213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+1)))
	if v213 != 0 {
		goto L75
	} else {
		goto L77
	}
L77:
	;
	if v35 == int32(0) {
		v500 = v34
		goto L16
	} else {
		goto L78
	}
L78:
	;
	v220 = v35
	v223 = int32(0)
	v225 = v34
	goto L82
L79:
	;
	__phi385 = v33
	__phi391 = v32
	v385 = __phi385
	v391 = __phi391
	goto L123
L80:
	;
	v373 = v34
	v374 = v220
	v379 = v225
	v380 = v6
	v381 = v225 - v34
	goto L79
L81:
	;
	v263 = v261 - v34
	v264 = F_palloc(m, v263)
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L4
	} else {
		goto L93
	}
L82:
	;
	v230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v225))))
	switch v230 - int32(92) {
	case 0:
		goto L85
	case 1, 2:
		v246 = v220
		v247 = v223
		v248 = v225
		goto L84
	case 3:
		goto L86
	default:
		goto L87
	}
L83:
	;
	v253 = int32(1)
	if v247&v253 == int32(0) {
		v500 = v250
		goto L16
	} else {
		goto L92
	}
L84:
	;
	v249 = int32(1)
	v250 = v248 + v249
	v252 = v246 - v249
	if v252 != 0 {
		v220 = v252
		v223 = v247
		v225 = v250
		goto L82
	} else {
		goto L91
	}
L85:
	;
	v240 = v220 - int32(1)
	if v240 == int32(0) {
		goto L17
	} else {
		goto L90
	}
L86:
	;
	if v223&int32(1) == int32(0) {
		goto L80
	} else {
		goto L89
	}
L87:
	;
	if v230 != int32(37) {
		v246 = v220
		v247 = v223
		v248 = v225
		goto L84
	} else {
		goto L88
	}
L88:
	;
	goto L86
L89:
	;
	v259 = v220
	v261 = v225
	v262 = v6
	goto L81
L90:
	;
	v243 = int32(1)
	v246 = v240
	v247 = v243
	v248 = v225 + v243
	goto L84
L91:
	;
	goto L83
L92:
	;
	v259 = int32(0)
	v261 = v250
	v262 = v253
	goto L81
L93:
	;
	if base.Ui32(v261) <= base.Ui32(v34) {
		v362 = v264
		goto L94
	} else {
		goto L95
	}
L94:
	;
	v369 = v362 - v264
	if v262 != 0 {
		goto L14
	} else {
		goto L122
	}
L95:
	;
	v268 = v263 & int32(3)
	if v268 == int32(0) {
		goto L97
	} else {
		goto L98
	}
L96:
	;
	if base.Ui32(int32(-4)) < base.Ui32(v34-v261) {
		v362 = v302
		goto L94
	} else {
		goto L106
	}
L97:
	;
	v301 = v34
	v302 = v264
	goto L96
L98:
	;
	goto L99
L99:
	;
	v276 = v34
	v277 = v264
	v278 = v6
	goto L100
L100:
	;
	v284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v276))))
	if v284 != int32(92) {
		goto L102
	} else {
		goto L103
	}
L101:
	;
	v301 = v292
	v302 = v290
	goto L96
L102:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v277))) = uint8(v284)
	v290 = v277 + int32(1)
	goto L104
L103:
	;
	v290 = v277
	goto L104
L104:
	;
	v291 = int32(1)
	v292 = v276 + v291
	v294 = v278 + v291
	if v294 != v268 {
		v276 = v292
		v277 = v290
		v278 = v294
		goto L100
	} else {
		goto L105
	}
L105:
	;
	goto L101
L106:
	;
	v317 = v301
	v318 = v302
	goto L107
L107:
	;
	v325 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v317))))
	if v325 != int32(92) {
		goto L109
	} else {
		goto L110
	}
L108:
	;
	v362 = v352
	goto L94
L109:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v318))) = uint8(v325)
	v331 = v318 + int32(1)
	goto L111
L110:
	;
	v331 = v318
	goto L111
L111:
	;
	v332 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v317)+1)))
	if v332 != int32(92) {
		goto L112
	} else {
		goto L113
	}
L112:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v331))) = uint8(v332)
	v338 = v331 + int32(1)
	goto L114
L113:
	;
	v338 = v331
	goto L114
L114:
	;
	v339 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v317)+2)))
	if v339 != int32(92) {
		goto L115
	} else {
		goto L116
	}
L115:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v338))) = uint8(v339)
	v345 = v338 + int32(1)
	goto L117
L116:
	;
	v345 = v338
	goto L117
L117:
	;
	v346 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v317)+3)))
	if v346 != int32(92) {
		goto L118
	} else {
		goto L119
	}
L118:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v345))) = uint8(v346)
	v352 = v345 + int32(1)
	goto L120
L119:
	;
	v352 = v345
	goto L120
L120:
	;
	v354 = v317 + int32(4)
	if v354 != v261 {
		v317 = v354
		v318 = v352
		goto L107
	} else {
		goto L121
	}
L121:
	;
	goto L108
L122:
	;
	v373 = v264
	v374 = v259
	v379 = v261
	v380 = v264
	v381 = v369
	goto L79
L123:
	;
	v398 = *(*int32)(unsafe.Add(mBase, _c_F_UTF8_MatchText[0]))
	if v398 != 0 {
		goto L125
	} else {
		goto L126
	}
L125:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L4
	} else {
		goto L128
	}
L126:
	;
	goto L127
L127:
	;
	v402 = F_pg_strncoll(m, v373, v381, v32, v391-v32, l4)
	mBase = m.M
	v403 = m.ExcPending
	if v403 != 0 {
		goto L4
	} else {
		goto L130
	}
L128:
	;
	goto L127
L129:
	;
	if v385 == int32(0) {
		goto L15
	} else {
		goto L136
	}
L130:
	;
	if v402 != 0 {
		goto L129
	} else {
		goto L131
	}
L131:
	;
	v404 = F_UTF8_MatchText(m, v391, v385, v379, v374, l4)
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
		goto L4
	} else {
		goto L132
	}
L132:
	;
	if v404 != int32(1) {
		goto L129
	} else {
		goto L133
	}
L133:
	;
	if v380 == int32(0) {
		v615 = int32(1)
		goto L6
	} else {
		goto L134
	}
L134:
	;
	F_pfree(m, v380)
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L4
	} else {
		goto L135
	}
L135:
	;
	return int32(1)
L136:
	;
	v419 = v385
	v423 = v391
	goto L137
L137:
	;
	v432 = v419 - int32(1)
	if v432 == int32(0) {
		goto L139
	} else {
		goto L140
	}
L138:
	;
	__phi385 = v432
	__phi391 = v437
	v385 = __phi385
	v391 = __phi391
	goto L123
L139:
	;
	__phi385 = v432
	__phi391 = v385 + v391
	v385 = __phi385
	v391 = __phi391
	goto L123
L140:
	;
	goto L141
L141:
	;
	v435 = int32(*(*int8)(unsafe.Add(mBase, uint32(v423)+1)))
	v437 = v423 + int32(1)
	if v435 < int32(-64) {
		v419 = v432
		v423 = v437
		goto L137
	} else {
		goto L142
	}
L142:
	;
	goto L138
L143:
	;
	return int32(0)
L144:
	;
	F_errcode(m, int32(84410498))
	mBase = m.M
	v450 = m.ExcPending
	if v450 != 0 {
		goto L4
	} else {
		goto L145
	}
L145:
	;
	F_errmsg(m, int32(_a_F_UTF8_MatchText_0), int32(0))
	mBase = m.M
	v454 = m.ExcPending
	if v454 != 0 {
		goto L4
	} else {
		goto L146
	}
L146:
	;
	F_errfinish(m, int32(_a_F_UTF8_MatchText_1), int32(107), int32(_a_F_UTF8_MatchText_2))
	mBase = m.M
	v459 = m.ExcPending
	if v459 != 0 {
		goto L4
	} else {
		goto L147
	}
L147:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L148:
	;
	F_errcode(m, int32(84410498))
	mBase = m.M
	v466 = m.ExcPending
	if v466 != 0 {
		goto L4
	} else {
		goto L149
	}
L149:
	;
	F_errmsg(m, int32(_a_F_UTF8_MatchText_0), int32(0))
	mBase = m.M
	v470 = m.ExcPending
	if v470 != 0 {
		goto L4
	} else {
		goto L150
	}
L150:
	;
	F_errfinish(m, int32(_a_F_UTF8_MatchText_1), int32(169), int32(_a_F_UTF8_MatchText_2))
	mBase = m.M
	v475 = m.ExcPending
	if v475 != 0 {
		goto L4
	} else {
		goto L151
	}
L151:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L152:
	;
	F_errcode(m, int32(84410498))
	mBase = m.M
	v482 = m.ExcPending
	if v482 != 0 {
		goto L4
	} else {
		goto L153
	}
L153:
	;
	F_errmsg(m, int32(_a_F_UTF8_MatchText_0), int32(0))
	mBase = m.M
	v486 = m.ExcPending
	if v486 != 0 {
		goto L4
	} else {
		goto L154
	}
L154:
	;
	F_errfinish(m, int32(_a_F_UTF8_MatchText_1), int32(237), int32(_a_F_UTF8_MatchText_2))
	mBase = m.M
	v491 = m.ExcPending
	if v491 != 0 {
		goto L4
	} else {
		goto L155
	}
L155:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L156:
	;
	return base.B2i32(v506 == int32(0))
L157:
	;
	F_pfree(m, v380)
	mBase = m.M
	v515 = m.ExcPending
	if v515 != 0 {
		goto L4
	} else {
		goto L158
	}
L158:
	;
	return int32(0)
L159:
	;
	if v264 != 0 {
		goto L160
	} else {
		goto L161
	}
L160:
	;
	F_pfree(m, v264)
	mBase = m.M
	v521 = m.ExcPending
	if v521 != 0 {
		goto L4
	} else {
		goto L163
	}
L161:
	;
	goto L162
L162:
	;
	return base.B2i32(v518 == int32(0))
L163:
	;
	goto L162
L164:
	;
	if int32(1) < v536 {
		v32 = v544
		v33 = v532
		v34 = v550
		v35 = v548
		goto L10
	} else {
		goto L165
	}
L165:
	;
	goto L11
L166:
	;
	v570 = v557
	v571 = v558
	goto L7
L167:
	;
	v586 = v570
	v587 = v571
	goto L168
L168:
	;
	v597 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v586))))
	if v597 != int32(37) {
		goto L170
	} else {
		goto L171
	}
L169:
	;
	v615 = v581
	goto L6
L170:
	;
	return int32(-1)
L171:
	;
	goto L172
L172:
	;
	v602 = int32(1)
	if v602 < v587 {
		v586 = v586 + v602
		v587 = v587 - v602
		goto L168
	} else {
		goto L173
	}
L173:
	;
	goto L169
}
func F_utf8_to_euc_jp(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v3 = int32(0)
	v7 = Fn14035(m, l0, int32(1), v3, v3, v3, int32(_a_F_utf8_to_euc_jp_0))
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
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
	var v128 int32
	_ = v128
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v159 int32
	_ = v159
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
		goto L5
	} else {
		goto L6
	}
L3:
	;
	F_report_untranslatable_char(m, int32(6), int32(8), v23, v24)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L1
	} else {
		goto L68
	}
L4:
	;
	v152 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v148))) = uint8(v152)
	return v144 - v11
L5:
	;
	v144 = v11
	v148 = v10
	goto L4
L6:
	;
	goto L7
L7:
	;
	v23 = v11
	v24 = v14
	v27 = v10
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
	v144 = v140
	v148 = v139
	goto L4
L10:
	;
	if v9 != 0 {
		v144 = v23
		v148 = v27
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
	v36 = m.ExcPending
	if v36 != 0 {
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
	v41 = int32(*(*int8)(unsafe.Add(mBase, uint32(v23))))
	if int32(0) <= v41 {
		goto L20
	} else {
		goto L21
	}
L16:
	;
	v134 = int32(-1)
	v135 = int32(1)
	v136 = v31
	goto L17
L17:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v27))) = uint8(v136)
	v139 = v27 + int32(1)
	v140 = v23 + v135
	v141 = v24 + v134
	if int32(0) < v141 {
		v23 = v140
		v24 = v141
		v27 = v139
		goto L8
	} else {
		goto L67
	}
L18:
	;
	if v65 != int32(2) {
		goto L59
	} else {
		goto L60
	}
L19:
	;
	if v65 <= v24 {
		goto L32
	} else {
		goto L33
	}
L20:
	;
	v65 = int32(1)
	goto L19
L21:
	;
	goto L22
L22:
	;
	v46 = v41 & int32(255)
	if v46&int32(224) == int32(192) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v65 = int32(2)
	goto L19
L24:
	;
	goto L25
L25:
	;
	if v46&int32(240) == int32(224) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v65 = int32(3)
	goto L19
L27:
	;
	goto L28
L28:
	;
	if v46&int32(248) == int32(240) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v63 = int32(4)
	goto L31
L30:
	;
	v63 = int32(1)
	goto L31
L31:
	;
	v65 = v63
	goto L19
L32:
	;
	v67 = int32(0)
	switch v65 - int32(1) {
	case 0:
		goto L39
	case 1:
		goto L40
	case 2:
		goto L41
	case 3:
		goto L42
	default:
		v116 = v67
		goto L36
	}
L33:
	;
	goto L34
L34:
	;
	if v9 != 0 {
		v144 = v23
		v148 = v27
		goto L4
	} else {
		goto L57
	}
L35:
	;
	if v116 != 0 {
		goto L18
	} else {
		goto L56
	}
L36:
	;
	goto L35
L37:
	;
	v116 = base.B2i32(base.Ui32(v108&int32(255)) < base.Ui32(int32(245)))
	goto L36
L38:
	;
	if base.I32_extend8_s(v103) < int32(-62) {
		v116 = v67
		goto L36
	} else {
		goto L55
	}
L39:
	;
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23))))
	v103 = v102
	goto L38
L40:
	;
	v76 = int32(*(*int8)(unsafe.Add(mBase, uint32(v23)+1)))
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23))))
	switch v77 - int32(224) {
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
	v73 = int32(*(*int8)(unsafe.Add(mBase, uint32(v23)+2)))
	if int32(-65) < v73 {
		v116 = v67
		goto L36
	} else {
		goto L44
	}
L42:
	;
	v70 = int32(*(*int8)(unsafe.Add(mBase, uint32(v23)+3)))
	if int32(-65) < v70 {
		v116 = v67
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
	if v76 <= int32(-65) {
		v103 = v77
		goto L38
	} else {
		goto L54
	}
L46:
	;
	if int32(-113) < v76 {
		v116 = v67
		goto L36
	} else {
		goto L53
	}
L47:
	;
	if base.Ui32((v76-int32(-64))&int32(255)) < base.Ui32(int32(208)) {
		v116 = v67
		goto L36
	} else {
		goto L52
	}
L48:
	;
	if int32(-97) < v76 {
		v116 = v67
		goto L36
	} else {
		goto L51
	}
L49:
	;
	v80 = int32(224)
	if base.Ui32(v80) <= base.Ui32((v76-int32(-64))&int32(255)) {
		v108 = v80
		goto L37
	} else {
		goto L50
	}
L50:
	;
	v116 = v67
	goto L36
L51:
	;
	v108 = int32(237)
	goto L37
L52:
	;
	v108 = int32(240)
	goto L37
L53:
	;
	v108 = int32(244)
	goto L37
L54:
	;
	v116 = v67
	goto L36
L55:
	;
	v108 = v103
	goto L37
L56:
	;
	goto L34
L57:
	;
	F_report_invalid_encoding(m, int32(6), v23, v24)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
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
	if v9 != 0 {
		v144 = v23
		v148 = v27
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
	if v9 != 0 {
		v144 = v23
		v148 = v27
		goto L4
	} else {
		goto L66
	}
L64:
	;
	goto L65
L65:
	;
	v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+1)))
	v134 = int32(-2)
	v135 = int32(2)
	v136 = v128&int32(63) | v31<<(uint(int32(6))%32)
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
