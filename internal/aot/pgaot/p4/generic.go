package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_GenericXLogFinish(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v30 int32
	_ = v30
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v151 int32
	_ = v151
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v238 int32
	_ = v238
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v258 int32
	_ = v258
	var v262 int32
	_ = v262
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v274 int32
	_ = v274
	var v278 int32
	_ = v278
	var v288 int32
	_ = v288
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v299 int32
	_ = v299
	var v303 int32
	_ = v303
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v317 int32
	_ = v317
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v347 int32
	_ = v347
	var v356 int32
	_ = v356
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v392 int32
	_ = v392
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v409 int32
	_ = v409
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v421 int32
	_ = v421
	var v424 int32
	_ = v424
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v433 int32
	_ = v433
	var v436 int32
	_ = v436
	var v439 int32
	_ = v439
	var v441 int32
	_ = v441
	var v449 int32
	_ = v449
	var v454 int64
	_ = v454
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v464 int32
	_ = v464
	var v470 int32
	_ = v470
	var v472 int32
	_ = v472
	var v478 int32
	_ = v478
	var v482 int32
	_ = v482
	var v486 int32
	_ = v486
	var v493 int32
	_ = v493
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v504 int32
	_ = v504
	var v508 int32
	_ = v508
	var v515 int32
	_ = v515
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v526 int32
	_ = v526
	var v532 int32
	_ = v532
	var v539 int32
	_ = v539
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v549 int32
	_ = v549
	var v551 int32
	_ = v551
	var v555 int32
	_ = v555
	var v559 int32
	_ = v559
	var v565 int32
	_ = v565
	var v567 int32
	_ = v567
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v577 int32
	_ = v577
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v584 int32
	_ = v584
	var v591 int32
	_ = v591
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v602 int32
	_ = v602
	var v604 int32
	_ = v604
	var v605 int32
	_ = v605
	var v609 int32
	_ = v609
	var v616 int32
	_ = v616
	var v622 int32
	_ = v622
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v627 int32
	_ = v627
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v636 int32
	_ = v636
	var v643 int32
	_ = v643
	var v649 int32
	_ = v649
	var v650 int32
	_ = v650
	var v651 int32
	_ = v651
	var v654 int32
	_ = v654
	var v656 int32
	_ = v656
	var v668 int32
	_ = v668
	var v670 int32
	_ = v670
	var v675 int32
	_ = v675
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_GenericXLogFinish[0]))))
	if v12 == int32(1) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v668 = int32(_a_F_GenericXLogFinish_0)
	v670 = *(*int32)(unsafe.Add(mBase, _c_F_GenericXLogFinish[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_GenericXLogFinish[1])) = v670 - int32(1)
	F_pfree(m, l0)
	mBase = m.M
	v675 = m.ExcPending
	if v675 != 0 {
		goto L5
	} else {
		goto L218
	}
L2:
	;
	F_XLogBeginInsert(m)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	v549 = int32(_a_F_GenericXLogFinish_0)
	v551 = *(*int32)(unsafe.Add(mBase, _c_F_GenericXLogFinish[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_GenericXLogFinish[1])) = v551 + int32(1)
	v555 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_GenericXLogFinish[2])))
	if v555 != 0 {
		goto L188
	} else {
		goto L189
	}
L5:
	;
	return
L6:
	;
	v17 = int32(_a_F_GenericXLogFinish_0)
	v19 = *(*int32)(unsafe.Add(mBase, _c_F_GenericXLogFinish[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_GenericXLogFinish[1])) = v19 + int32(1)
	v30 = int32(0)
	goto L7
L7:
	;
	v38 = l0 + int32(_a_F_GenericXLogFinish_1) + v30*int32(_a_F_GenericXLogFinish_2)
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
	if v39 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v454 = F_XLogInsert(m, int32(20), int32(0))
	mBase = m.M
	v455 = m.ExcPending
	if v455 != 0 {
		goto L5
	} else {
		goto L161
	}
L9:
	;
	v449 = v30 + int32(1)
	if v449 != int32(4) {
		v30 = v449
		goto L7
	} else {
		goto L160
	}
L10:
	;
	if v39 < int32(0) {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v38)+12))
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+4)))
	if v61&int32(1) == int32(0) {
		goto L15
	} else {
		goto L16
	}
L12:
	;
	v45 = *(*int32)(unsafe.Add(mBase, _c_F_GenericXLogFinish[3]))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v45+(v39^int32(-1))<<(uint(int32(2))%32))))
	v59 = v51
	goto L11
L13:
	;
	goto L14
L14:
	;
	v53 = *(*int32)(unsafe.Add(mBase, _c_F_GenericXLogFinish[4]))
	v59 = v53 + v39<<(uint(int32(13))%32) + int32(-8192)
	goto L11
L15:
	;
	v66 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v59)+14)))
	v67 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v60)+14)))
	v68 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v59)+12)))
	v69 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v60)+12)))
	v70 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v38)+8)) = v70
	v79 = int32(-1)
	goto L20
L16:
	;
	v405 = v60
	goto L17
L17:
	;
	v409 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v60)+12)))
	if v409 != 0 {
		goto L144
	} else {
		goto L145
	}
L18:
	;
	v238 = int32(_a_F_GenericXLogFinish_3)
	v245 = int32(-1)
	v247 = base.B2i32(base.Ui32(v67) < base.Ui32(v66))
	if base.Ui32(v67) < base.Ui32(v66) {
		goto L82
	} else {
		goto L83
	}
L20:
	;
	goto L21
L21:
	;
	goto L24
L22:
	;
	if v204 < int32(0) {
		goto L63
	} else {
		goto L64
	}
L24:
	;
	goto L25
L25:
	;
	if base.Ui32(v69) < base.Ui32(v68) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v85 = v69
	goto L28
L27:
	;
	v85 = v68
	goto L28
L28:
	;
	if base.Ui32(v85) <= base.Ui32(v70) {
		v204 = v79
		v205 = v79
		goto L22
	} else {
		goto L29
	}
L29:
	;
	v92 = v70
	v96 = v79
	goto L30
L30:
	;
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59+v92))))
	v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60+v92))))
	if v102 != v104 {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	v204 = v193
	v205 = v194
	goto L22
L32:
	;
	if v96 < int32(0) {
		goto L35
	} else {
		goto L36
	}
L33:
	;
	v133 = v92
	v137 = v96
	goto L34
L34:
	;
	v143 = v133 + int32(1)
	if v143 < v85 {
		goto L44
	} else {
		goto L45
	}
L35:
	;
	v108 = v92
	goto L37
L36:
	;
	v108 = v96
	goto L37
L37:
	;
	v112 = v92
	goto L38
L38:
	;
	v122 = v112 + int32(1)
	if v85 <= v122 {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	v133 = v122
	v137 = v108
	goto L34
L40:
	;
	v204 = v108
	v205 = int32(-1)
	goto L22
L41:
	;
	goto L42
L42:
	;
	v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59+v122))))
	v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60+v122))))
	if v126 != v128 {
		v112 = v122
		goto L38
	} else {
		goto L43
	}
L43:
	;
	goto L39
L44:
	;
	v145 = v85
	goto L46
L45:
	;
	v145 = v143
	goto L46
L46:
	;
	v151 = v133
	goto L47
L47:
	;
	if v151 == v145-int32(1) {
		goto L50
	} else {
		goto L51
	}
L48:
	;
	if v137 < int32(0) {
		goto L55
	} else {
		goto L56
	}
L49:
	;
	goto L48
L50:
	;
	v168 = v145
	goto L49
L51:
	;
	goto L52
L52:
	;
	v162 = v151 + int32(1)
	v164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59+v162))))
	v166 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60+v162))))
	if v164 == v166 {
		v151 = v162
		goto L47
	} else {
		goto L53
	}
L53:
	;
	v168 = v162
	goto L49
L54:
	;
	if v168 < v85 {
		v92 = v168
		v96 = v193
		goto L30
	} else {
		goto L62
	}
L55:
	;
	v193 = int32(-1)
	v194 = v133
	goto L54
L56:
	;
	goto L57
L57:
	;
	if base.Ui32(v168-v133) < base.Ui32(int32(5)) {
		v193 = v137
		v194 = v133
		goto L54
	} else {
		goto L58
	}
L58:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v38)+8))
	v176 = v38 + int32(16) + v175
	v177 = v133 - v137
	*(*uint16)(unsafe.Add(mBase, uint32(v176)+2)) = uint16(v177)
	*(*uint16)(unsafe.Add(mBase, uint32(v176))) = uint16(v137)
	v181 = v177 & int32(_a_F_GenericXLogFinish_4)
	if v181 != 0 {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	base.MemoryCopy(m, v176+int32(4), v60+v137, v181)
	goto L61
L60:
	;
	goto L61
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+8)) = v181 + v175 + int32(4)
	v190 = int32(-1)
	v193 = v190
	v194 = v190
	goto L54
L62:
	;
	goto L31
L63:
	;
	v211 = v85
	goto L65
L64:
	;
	v211 = v204
	goto L65
L65:
	;
	v212 = base.B2i32(base.Ui32(v68) < base.Ui32(v69))
	if base.Ui32(v68) < base.Ui32(v69) {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v213 = v211
	goto L68
L67:
	;
	v213 = v204
	goto L68
L68:
	;
	if int32(0) <= v213 {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v38)+8))
	v217 = v38 + v216
	*(*uint16)(unsafe.Add(mBase, uint32(v217)+16)) = uint16(v213)
	if base.Ui32(v68) < base.Ui32(v69) {
		goto L72
	} else {
		goto L73
	}
L70:
	;
	goto L71
L71:
	;
	goto L18
L72:
	;
	v219 = v69
	goto L74
L73:
	;
	v219 = v205
	goto L74
L74:
	;
	if v219 < int32(0) {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v222 = v69
	goto L77
L76:
	;
	v222 = v219
	goto L77
L77:
	;
	v223 = v222 - v213
	*(*uint16)(unsafe.Add(mBase, uint32(v217)+18)) = uint16(v223)
	v226 = v223 & int32(_a_F_GenericXLogFinish_4)
	if v226 != 0 {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	base.MemoryCopy(m, v217+int32(20), v213+v60, v226)
	goto L80
L79:
	;
	goto L80
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+8)) = v226 + v216 + int32(4)
	goto L71
L81:
	;
	v404 = *(*int32)(unsafe.Add(mBase, uint32(v38)+12))
	v405 = v404
	goto L17
L82:
	;
	v248 = v67
	goto L84
L83:
	;
	v248 = v245
	goto L84
L84:
	;
	if base.Ui32(v67) < base.Ui32(v66) {
		goto L86
	} else {
		goto L87
	}
L85:
	;
	goto L126
L86:
	;
	v249 = v66
	goto L88
L87:
	;
	v249 = v67
	goto L88
L88:
	;
	goto L90
L90:
	;
	goto L91
L91:
	;
	if base.Ui32(v238) <= base.Ui32(v249) {
		v370 = v248
		v371 = v245
		goto L85
	} else {
		goto L92
	}
L92:
	;
	v258 = v249
	v262 = v248
	goto L93
L93:
	;
	v268 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59+v258))))
	v270 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60+v258))))
	if v268 != v270 {
		goto L95
	} else {
		goto L96
	}
L94:
	;
	v370 = v359
	v371 = v360
	goto L85
L95:
	;
	if v262 < int32(0) {
		goto L98
	} else {
		goto L99
	}
L96:
	;
	v299 = v258
	v303 = v262
	goto L97
L97:
	;
	v309 = v299 + int32(1)
	if v309 < v238 {
		goto L107
	} else {
		goto L108
	}
L98:
	;
	v274 = v258
	goto L100
L99:
	;
	v274 = v262
	goto L100
L100:
	;
	v278 = v258
	goto L101
L101:
	;
	v288 = v278 + int32(1)
	if v238 <= v288 {
		goto L103
	} else {
		goto L104
	}
L102:
	;
	v299 = v288
	v303 = v274
	goto L97
L103:
	;
	v370 = v274
	v371 = int32(-1)
	goto L85
L104:
	;
	goto L105
L105:
	;
	v292 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59+v288))))
	v294 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60+v288))))
	if v292 != v294 {
		v278 = v288
		goto L101
	} else {
		goto L106
	}
L106:
	;
	goto L102
L107:
	;
	v311 = v238
	goto L109
L108:
	;
	v311 = v309
	goto L109
L109:
	;
	v317 = v299
	goto L110
L110:
	;
	if v317 == v311-int32(1) {
		goto L113
	} else {
		goto L114
	}
L111:
	;
	if v303 < int32(0) {
		goto L118
	} else {
		goto L119
	}
L112:
	;
	goto L111
L113:
	;
	v334 = v311
	goto L112
L114:
	;
	goto L115
L115:
	;
	v328 = v317 + int32(1)
	v330 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59+v328))))
	v332 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60+v328))))
	if v330 == v332 {
		v317 = v328
		goto L110
	} else {
		goto L116
	}
L116:
	;
	v334 = v328
	goto L112
L117:
	;
	if v334 < v238 {
		v258 = v334
		v262 = v359
		goto L93
	} else {
		goto L125
	}
L118:
	;
	v359 = int32(-1)
	v360 = v299
	goto L117
L119:
	;
	goto L120
L120:
	;
	if base.Ui32(v334-v299) < base.Ui32(int32(5)) {
		v359 = v303
		v360 = v299
		goto L117
	} else {
		goto L121
	}
L121:
	;
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v38)+8))
	v342 = v38 + int32(16) + v341
	v343 = v299 - v303
	*(*uint16)(unsafe.Add(mBase, uint32(v342)+2)) = uint16(v343)
	*(*uint16)(unsafe.Add(mBase, uint32(v342))) = uint16(v303)
	v347 = v343 & int32(_a_F_GenericXLogFinish_4)
	if v347 != 0 {
		goto L122
	} else {
		goto L123
	}
L122:
	;
	base.MemoryCopy(m, v342+int32(4), v60+v303, v347)
	goto L124
L123:
	;
	goto L124
L124:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+8)) = v347 + v341 + int32(4)
	v356 = int32(-1)
	v359 = v356
	v360 = v356
	goto L117
L125:
	;
	goto L94
L126:
	;
	goto L128
L128:
	;
	goto L130
L130:
	;
	goto L131
L131:
	;
	if int32(0) <= v370 {
		goto L132
	} else {
		goto L133
	}
L132:
	;
	v382 = *(*int32)(unsafe.Add(mBase, uint32(v38)+8))
	v383 = v38 + v382
	*(*uint16)(unsafe.Add(mBase, uint32(v383)+16)) = uint16(v370)
	goto L136
L133:
	;
	goto L134
L134:
	;
	goto L81
L136:
	;
	goto L137
L137:
	;
	if v371 < int32(0) {
		goto L138
	} else {
		goto L139
	}
L138:
	;
	v388 = v238
	goto L140
L139:
	;
	v388 = v371
	goto L140
L140:
	;
	v389 = v388 - v370
	*(*uint16)(unsafe.Add(mBase, uint32(v383)+18)) = uint16(v389)
	v392 = v389 & int32(_a_F_GenericXLogFinish_4)
	if v392 != 0 {
		goto L141
	} else {
		goto L142
	}
L141:
	;
	base.MemoryCopy(m, v383+int32(20), v370+v60, v392)
	goto L143
L142:
	;
	goto L143
L143:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+8)) = v392 + v382 + int32(4)
	goto L134
L144:
	;
	base.MemoryCopy(m, v59, v405, v409)
	goto L146
L145:
	;
	goto L146
L146:
	;
	v411 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v60)+14)))
	v412 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v60)+12)))
	v413 = v411 - v412
	if v413 != 0 {
		goto L147
	} else {
		goto L148
	}
L147:
	;
	base.MemoryFill(m, v59+v412, int32(0), v413)
	goto L149
L148:
	;
	goto L149
L149:
	;
	v418 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v60)+14)))
	v419 = int32(_a_F_GenericXLogFinish_3) - v418
	if v419 != 0 {
		goto L150
	} else {
		goto L151
	}
L150:
	;
	v421 = *(*int32)(unsafe.Add(mBase, uint32(v38)+12))
	base.MemoryCopy(m, v59+v418, v421+v418, v419)
	goto L152
L151:
	;
	goto L152
L152:
	;
	v424 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
	F_MarkBufferDirty(m, v424)
	mBase = m.M
	v426 = m.ExcPending
	if v426 != 0 {
		goto L5
	} else {
		goto L153
	}
L153:
	;
	v427 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
	v428 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+4)))
	if v428&int32(1) != 0 {
		goto L154
	} else {
		goto L155
	}
L154:
	;
	F_XLogRegisterBuffer(m, v30, v427, int32(9))
	mBase = m.M
	v433 = m.ExcPending
	if v433 != 0 {
		goto L5
	} else {
		goto L157
	}
L155:
	;
	goto L156
L156:
	;
	F_XLogRegisterBuffer(m, v30, v427, int32(8))
	mBase = m.M
	v436 = m.ExcPending
	if v436 != 0 {
		goto L5
	} else {
		goto L158
	}
L157:
	;
	goto L9
L158:
	;
	v439 = *(*int32)(unsafe.Add(mBase, uint32(v38)+8))
	F_XLogRegisterBufData(m, v30, v38+int32(16), v439)
	mBase = m.M
	v441 = m.ExcPending
	if v441 != 0 {
		goto L5
	} else {
		goto L159
	}
L159:
	;
	goto L9
L160:
	;
	goto L8
L161:
	;
	v456 = base.I32_wrap_i64(v454)
	v459 = base.I32_wrap_i64(int64(base.Ui64(v454) >> (uint(int64(32)) % 64)))
	v460 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_GenericXLogFinish[2])))
	if v460 != 0 {
		goto L162
	} else {
		goto L163
	}
L162:
	;
	if v460 < int32(0) {
		goto L166
	} else {
		goto L167
	}
L163:
	;
	goto L164
L164:
	;
	v482 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_GenericXLogFinish[5])))
	if v482 != 0 {
		goto L169
	} else {
		goto L170
	}
L165:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v478)+4)) = v456
	*(*int32)(unsafe.Add(mBase, uint32(v478))) = v459
	goto L164
L166:
	;
	v464 = *(*int32)(unsafe.Add(mBase, _c_F_GenericXLogFinish[3]))
	v470 = *(*int32)(unsafe.Add(mBase, uint32(v464+(v460^int32(-1))<<(uint(int32(2))%32))))
	v478 = v470
	goto L165
L167:
	;
	goto L168
L168:
	;
	v472 = *(*int32)(unsafe.Add(mBase, _c_F_GenericXLogFinish[4]))
	v478 = v472 + v460<<(uint(int32(13))%32) + int32(-8192)
	goto L165
L169:
	;
	if int32(0) <= v482 {
		goto L173
	} else {
		goto L174
	}
L170:
	;
	goto L171
L171:
	;
	v504 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_GenericXLogFinish[6])))
	if v504 != 0 {
		goto L176
	} else {
		goto L177
	}
L172:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v500)+4)) = v456
	*(*int32)(unsafe.Add(mBase, uint32(v500))) = v459
	goto L171
L173:
	;
	v486 = *(*int32)(unsafe.Add(mBase, _c_F_GenericXLogFinish[4]))
	v500 = v486 + v482<<(uint(int32(13))%32) + int32(-8192)
	goto L172
L174:
	;
	goto L175
L175:
	;
	v493 = *(*int32)(unsafe.Add(mBase, _c_F_GenericXLogFinish[3]))
	v499 = *(*int32)(unsafe.Add(mBase, uint32(v493+(v482^int32(-1))<<(uint(int32(2))%32))))
	v500 = v499
	goto L172
L176:
	;
	if int32(0) <= v504 {
		goto L180
	} else {
		goto L181
	}
L177:
	;
	goto L178
L178:
	;
	v526 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_GenericXLogFinish[7])))
	if v526 == int32(0) {
		goto L1
	} else {
		goto L183
	}
L179:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v522)+4)) = v456
	*(*int32)(unsafe.Add(mBase, uint32(v522))) = v459
	goto L178
L180:
	;
	v508 = *(*int32)(unsafe.Add(mBase, _c_F_GenericXLogFinish[4]))
	v522 = v508 + v504<<(uint(int32(13))%32) + int32(-8192)
	goto L179
L181:
	;
	goto L182
L182:
	;
	v515 = *(*int32)(unsafe.Add(mBase, _c_F_GenericXLogFinish[3]))
	v521 = *(*int32)(unsafe.Add(mBase, uint32(v515+(v504^int32(-1))<<(uint(int32(2))%32))))
	v522 = v521
	goto L179
L183:
	;
	if int32(0) <= v526 {
		goto L185
	} else {
		goto L186
	}
L184:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v546)+4)) = v456
	*(*int32)(unsafe.Add(mBase, uint32(v546))) = v459
	goto L1
L185:
	;
	v532 = *(*int32)(unsafe.Add(mBase, _c_F_GenericXLogFinish[4]))
	v546 = v532 + v526<<(uint(int32(13))%32) + int32(-8192)
	goto L184
L186:
	;
	goto L187
L187:
	;
	v539 = *(*int32)(unsafe.Add(mBase, _c_F_GenericXLogFinish[3]))
	v545 = *(*int32)(unsafe.Add(mBase, uint32(v539+(v526^int32(-1))<<(uint(int32(2))%32))))
	v546 = v545
	goto L184
L188:
	;
	if v555 < int32(0) {
		goto L192
	} else {
		goto L193
	}
L189:
	;
	goto L190
L190:
	;
	v580 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_GenericXLogFinish[5])))
	if v580 != 0 {
		goto L196
	} else {
		goto L197
	}
L191:
	;
	v574 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_GenericXLogFinish[8])))
	base.MemoryCopy(m, v573, v574, int32(_a_F_GenericXLogFinish_3))
	v577 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_GenericXLogFinish[2])))
	F_MarkBufferDirty(m, v577)
	mBase = m.M
	v579 = m.ExcPending
	if v579 != 0 {
		goto L5
	} else {
		goto L195
	}
L192:
	;
	v559 = *(*int32)(unsafe.Add(mBase, _c_F_GenericXLogFinish[3]))
	v565 = *(*int32)(unsafe.Add(mBase, uint32(v559+(v555^int32(-1))<<(uint(int32(2))%32))))
	v573 = v565
	goto L191
L193:
	;
	goto L194
L194:
	;
	v567 = *(*int32)(unsafe.Add(mBase, _c_F_GenericXLogFinish[4]))
	v573 = v567 + v555<<(uint(int32(13))%32) + int32(-8192)
	goto L191
L195:
	;
	goto L190
L196:
	;
	if int32(0) <= v580 {
		goto L200
	} else {
		goto L201
	}
L197:
	;
	goto L198
L198:
	;
	v605 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_GenericXLogFinish[6])))
	if v605 != 0 {
		goto L204
	} else {
		goto L205
	}
L199:
	;
	v599 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_GenericXLogFinish[9])))
	base.MemoryCopy(m, v598, v599, int32(_a_F_GenericXLogFinish_3))
	v602 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_GenericXLogFinish[5])))
	F_MarkBufferDirty(m, v602)
	mBase = m.M
	v604 = m.ExcPending
	if v604 != 0 {
		goto L5
	} else {
		goto L203
	}
L200:
	;
	v584 = *(*int32)(unsafe.Add(mBase, _c_F_GenericXLogFinish[4]))
	v598 = v584 + v580<<(uint(int32(13))%32) + int32(-8192)
	goto L199
L201:
	;
	goto L202
L202:
	;
	v591 = *(*int32)(unsafe.Add(mBase, _c_F_GenericXLogFinish[3]))
	v597 = *(*int32)(unsafe.Add(mBase, uint32(v591+(v580^int32(-1))<<(uint(int32(2))%32))))
	v598 = v597
	goto L199
L203:
	;
	goto L198
L204:
	;
	if int32(0) <= v605 {
		goto L208
	} else {
		goto L209
	}
L205:
	;
	goto L206
L206:
	;
	v630 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_GenericXLogFinish[7])))
	if v630 == int32(0) {
		goto L1
	} else {
		goto L212
	}
L207:
	;
	v624 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_GenericXLogFinish[10])))
	base.MemoryCopy(m, v623, v624, int32(_a_F_GenericXLogFinish_3))
	v627 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_GenericXLogFinish[6])))
	F_MarkBufferDirty(m, v627)
	mBase = m.M
	v629 = m.ExcPending
	if v629 != 0 {
		goto L5
	} else {
		goto L211
	}
L208:
	;
	v609 = *(*int32)(unsafe.Add(mBase, _c_F_GenericXLogFinish[4]))
	v623 = v609 + v605<<(uint(int32(13))%32) + int32(-8192)
	goto L207
L209:
	;
	goto L210
L210:
	;
	v616 = *(*int32)(unsafe.Add(mBase, _c_F_GenericXLogFinish[3]))
	v622 = *(*int32)(unsafe.Add(mBase, uint32(v616+(v605^int32(-1))<<(uint(int32(2))%32))))
	v623 = v622
	goto L207
L211:
	;
	goto L206
L212:
	;
	if int32(0) <= v630 {
		goto L214
	} else {
		goto L215
	}
L213:
	;
	v651 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_GenericXLogFinish[11])))
	base.MemoryCopy(m, v650, v651, int32(_a_F_GenericXLogFinish_3))
	v654 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_GenericXLogFinish[7])))
	F_MarkBufferDirty(m, v654)
	mBase = m.M
	v656 = m.ExcPending
	if v656 != 0 {
		goto L5
	} else {
		goto L217
	}
L214:
	;
	v636 = *(*int32)(unsafe.Add(mBase, _c_F_GenericXLogFinish[4]))
	v650 = v636 + v630<<(uint(int32(13))%32) + int32(-8192)
	goto L213
L215:
	;
	goto L216
L216:
	;
	v643 = *(*int32)(unsafe.Add(mBase, _c_F_GenericXLogFinish[3]))
	v649 = *(*int32)(unsafe.Add(mBase, uint32(v643+(v630^int32(-1))<<(uint(int32(2))%32))))
	v650 = v649
	goto L213
L217:
	;
	goto L1
L218:
	;
	return
}
func F_GenericXLogRegisterBuffer(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v12 = l0 + int32(_a_F_GenericXLogRegisterBuffer_0)
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_GenericXLogRegisterBuffer[0])))
	if v13 == int32(0) {
		v52 = v12
		*(*int32)(unsafe.Add(mBase, uint32(v52)+4)) = l2
		*(*int32)(unsafe.Add(mBase, uint32(v52))) = l1
		v56 = *(*int32)(unsafe.Add(mBase, uint32(v52)+12))
		if l1 < int32(0) {
			v60 = *(*int32)(unsafe.Add(mBase, _c_F_GenericXLogRegisterBuffer[1]))
			v66 = *(*int32)(unsafe.Add(mBase, uint32(v60+(l1^int32(-1))<<(uint(int32(2))%32))))
			v74 = v66
		} else {
			v68 = *(*int32)(unsafe.Add(mBase, _c_F_GenericXLogRegisterBuffer[2]))
			v74 = v68 + l1<<(uint(int32(13))%32) + int32(-8192)
		}
		base.MemoryCopy(m, v56, v74, int32(_a_F_GenericXLogRegisterBuffer_1))
		v78 = v52
		v82 = *(*int32)(unsafe.Add(mBase, uint32(v78+int32(12))))
		m.G0 = v9 + int32(16)
		return v82
	} else {
		if l1 == v13 {
			v78 = v12
			v82 = *(*int32)(unsafe.Add(mBase, uint32(v78+int32(12))))
			m.G0 = v9 + int32(16)
			return v82
		} else {
			v18 = l0 + int32(_a_F_GenericXLogRegisterBuffer_2)
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_GenericXLogRegisterBuffer[3])))
			if v19 == int32(0) {
				v52 = v18
				*(*int32)(unsafe.Add(mBase, uint32(v52)+4)) = l2
				*(*int32)(unsafe.Add(mBase, uint32(v52))) = l1
				v56 = *(*int32)(unsafe.Add(mBase, uint32(v52)+12))
				if l1 < int32(0) {
					v60 = *(*int32)(unsafe.Add(mBase, _c_F_GenericXLogRegisterBuffer[1]))
					v66 = *(*int32)(unsafe.Add(mBase, uint32(v60+(l1^int32(-1))<<(uint(int32(2))%32))))
					v74 = v66
				} else {
					v68 = *(*int32)(unsafe.Add(mBase, _c_F_GenericXLogRegisterBuffer[2]))
					v74 = v68 + l1<<(uint(int32(13))%32) + int32(-8192)
				}
				base.MemoryCopy(m, v56, v74, int32(_a_F_GenericXLogRegisterBuffer_1))
				v78 = v52
				v82 = *(*int32)(unsafe.Add(mBase, uint32(v78+int32(12))))
				m.G0 = v9 + int32(16)
				return v82
			} else {
				if l1 == v19 {
					v78 = v18
					v82 = *(*int32)(unsafe.Add(mBase, uint32(v78+int32(12))))
					m.G0 = v9 + int32(16)
					return v82
				} else {
					v24 = l0 + int32(_a_F_GenericXLogRegisterBuffer_3)
					v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_GenericXLogRegisterBuffer[4])))
					if v25 == int32(0) {
						v52 = v24
						*(*int32)(unsafe.Add(mBase, uint32(v52)+4)) = l2
						*(*int32)(unsafe.Add(mBase, uint32(v52))) = l1
						v56 = *(*int32)(unsafe.Add(mBase, uint32(v52)+12))
						if l1 < int32(0) {
							v60 = *(*int32)(unsafe.Add(mBase, _c_F_GenericXLogRegisterBuffer[1]))
							v66 = *(*int32)(unsafe.Add(mBase, uint32(v60+(l1^int32(-1))<<(uint(int32(2))%32))))
							v74 = v66
						} else {
							v68 = *(*int32)(unsafe.Add(mBase, _c_F_GenericXLogRegisterBuffer[2]))
							v74 = v68 + l1<<(uint(int32(13))%32) + int32(-8192)
						}
						base.MemoryCopy(m, v56, v74, int32(_a_F_GenericXLogRegisterBuffer_1))
						v78 = v52
						v82 = *(*int32)(unsafe.Add(mBase, uint32(v78+int32(12))))
						m.G0 = v9 + int32(16)
						return v82
					} else {
						if l1 == v25 {
							v78 = v24
							v82 = *(*int32)(unsafe.Add(mBase, uint32(v78+int32(12))))
							m.G0 = v9 + int32(16)
							return v82
						} else {
							v30 = l0 + int32(_a_F_GenericXLogRegisterBuffer_4)
							v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_GenericXLogRegisterBuffer[5])))
							if v31 == int32(0) {
								v52 = v30
								*(*int32)(unsafe.Add(mBase, uint32(v52)+4)) = l2
								*(*int32)(unsafe.Add(mBase, uint32(v52))) = l1
								v56 = *(*int32)(unsafe.Add(mBase, uint32(v52)+12))
								if l1 < int32(0) {
									v60 = *(*int32)(unsafe.Add(mBase, _c_F_GenericXLogRegisterBuffer[1]))
									v66 = *(*int32)(unsafe.Add(mBase, uint32(v60+(l1^int32(-1))<<(uint(int32(2))%32))))
									v74 = v66
								} else {
									v68 = *(*int32)(unsafe.Add(mBase, _c_F_GenericXLogRegisterBuffer[2]))
									v74 = v68 + l1<<(uint(int32(13))%32) + int32(-8192)
								}
								base.MemoryCopy(m, v56, v74, int32(_a_F_GenericXLogRegisterBuffer_1))
								v78 = v52
								v82 = *(*int32)(unsafe.Add(mBase, uint32(v78+int32(12))))
								m.G0 = v9 + int32(16)
								return v82
							} else {
								if v31 == l1 {
									v78 = v30
									v82 = *(*int32)(unsafe.Add(mBase, uint32(v78+int32(12))))
									m.G0 = v9 + int32(16)
									return v82
								} else {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v40 = m.ExcPending
									if v40 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v9))) = int32(4)
										F_errmsg_internal(m, int32(_a_F_GenericXLogRegisterBuffer_5), v9)
										mBase = m.M
										v45 = m.ExcPending
										if v45 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_GenericXLogRegisterBuffer_6), int32(327), int32(_a_F_GenericXLogRegisterBuffer_7))
											mBase = m.M
											v50 = m.ExcPending
											if v50 != 0 {
												return int32(0)
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_Generic_Text_IC_like(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v156 int32
	_ = v156
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
	var v165 int32
	_ = v165
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v235 int32
	_ = v235
	if l2 != 0 {
		v10 = F_pg_newlocale_from_collation(m, l2)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int32(0)
		} else {
			v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+1)))
			if v14 == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v223 = m.ExcPending
				if v223 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(1088))
					mBase = m.M
					v226 = m.ExcPending
					if v226 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(_a_F_Generic_Text_IC_like_0), int32(0))
						mBase = m.M
						v230 = m.ExcPending
						if v230 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_Generic_Text_IC_like_1), int32(202), int32(_a_F_Generic_Text_IC_like_2))
							mBase = m.M
							v235 = m.ExcPending
							if v235 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			} else {
				v18 = *(*int32)(unsafe.Add(mBase, _c_F_Generic_Text_IC_like[0]))
				v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
				v24 = *(*int32)(unsafe.Add(mBase, uint32(v19*int32(28))+uint32(_c_F_Generic_Text_IC_like[1])))
				if v24 <= int32(1) {
					v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10))))
					if v27 != int32(105) {
						v125 = int32(1)
						v126 = l1 + v125
						v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
						v129 = v127 & v125
						if v127 == v125 {
							v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126))))
							if v135 == int32(18) {
								v138 = int32(16)
							} else {
								v138 = int32(0)
							}
							if base.Ui32((v135-int32(1))&int32(255)) < base.Ui32(int32(3)) {
								v145 = int32(4)
							} else {
								v145 = v138
							}
							v156 = v145
						} else {
							v146 = int32(1)
							if v129 != 0 {
								v156 = int32(base.Ui32(v127)>>(uint(v146)%32)) - v146
							} else {
								v150 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
								v156 = int32(base.Ui32(v150)>>(uint(int32(2))%32)) - int32(4)
							}
						}
						if v129 != 0 {
							v157 = v126
						} else {
							v157 = l1 + int32(4)
						}
						v158 = int32(1)
						v159 = l0 + v158
						v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
						v164 = v162 & v158
						if v164 != 0 {
							v165 = v159
						} else {
							v165 = l0 + int32(4)
						}
						if v162 == int32(1) {
							v171 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v159))))
							if v171 == int32(18) {
								v174 = int32(16)
							} else {
								v174 = int32(0)
							}
							if base.Ui32((v171-int32(1))&int32(255)) < base.Ui32(int32(3)) {
								v181 = int32(4)
							} else {
								v181 = v174
							}
							v182 = F_SB_IMatchText(m, v165, v181, v157, v156, v10)
							mBase = m.M
							v183 = m.ExcPending
							if v183 != 0 {
								return int32(0)
							} else {
								return v182
							}
						} else {
							if v164 != 0 {
								v185 = int32(1)
								v189 = F_SB_IMatchText(m, v165, int32(base.Ui32(v162)>>(uint(v185)%32))-v185, v157, v156, v10)
								mBase = m.M
								v190 = m.ExcPending
								if v190 != 0 {
									return int32(0)
								} else {
									return v189
								}
							} else {
								v192 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								v197 = F_SB_IMatchText(m, v165, int32(base.Ui32(v192)>>(uint(int32(2))%32))-int32(4), v157, v156, v10)
								mBase = m.M
								v198 = m.ExcPending
								if v198 != 0 {
									return int32(0)
								} else {
									return v197
								}
							}
						}
					} else {
						v31 = F_DirectFunctionCall1Coll(m, int32(1437), l2, l1)
						mBase = m.M
						v32 = m.ExcPending
						if v32 != 0 {
							return int32(0)
						} else {
							v33 = F_pg_detoast_datum_packed(m, v31)
							mBase = m.M
							v34 = m.ExcPending
							if v34 != 0 {
								return int32(0)
							} else {
								v35 = int32(1)
								v36 = v33 + v35
								v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33))))
								v39 = v37 & v35
								if v37 == v35 {
									v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36))))
									if v45 == int32(18) {
										v48 = int32(16)
									} else {
										v48 = int32(0)
									}
									if base.Ui32((v45-int32(1))&int32(255)) < base.Ui32(int32(3)) {
										v55 = int32(4)
									} else {
										v55 = v48
									}
									v66 = v55
								} else {
									v56 = int32(1)
									if v39 != 0 {
										v66 = int32(base.Ui32(v37)>>(uint(v56)%32)) - v56
									} else {
										v60 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
										v66 = int32(base.Ui32(v60)>>(uint(int32(2))%32)) - int32(4)
									}
								}
								v70 = F_DirectFunctionCall1Coll(m, int32(1437), l2, l0)
								mBase = m.M
								v71 = m.ExcPending
								if v71 != 0 {
									return int32(0)
								} else {
									v72 = F_pg_detoast_datum_packed(m, v70)
									mBase = m.M
									v73 = m.ExcPending
									if v73 != 0 {
										return int32(0)
									} else {
										v76 = int32(1)
										v77 = v72 + v76
										v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72))))
										v80 = v78 & v76
										if v78 == v76 {
											v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77))))
											if v86 == int32(18) {
												v89 = int32(16)
											} else {
												v89 = int32(0)
											}
											if base.Ui32((v86-int32(1))&int32(255)) < base.Ui32(int32(3)) {
												v96 = int32(4)
											} else {
												v96 = v89
											}
											v107 = v96
										} else {
											v97 = int32(1)
											if v80 != 0 {
												v107 = int32(base.Ui32(v78)>>(uint(v97)%32)) - v97
											} else {
												v101 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
												v107 = int32(base.Ui32(v101)>>(uint(int32(2))%32)) - int32(4)
											}
										}
										if v39 != 0 {
											v108 = v36
										} else {
											v108 = v33 + int32(4)
										}
										if v80 != 0 {
											v109 = v77
										} else {
											v109 = v72 + int32(4)
										}
										v111 = *(*int32)(unsafe.Add(mBase, _c_F_Generic_Text_IC_like[0]))
										v112 = *(*int32)(unsafe.Add(mBase, uint32(v111)+4))
										if v112 == int32(6) {
											v116 = F_UTF8_MatchText(m, v109, v107, v108, v66, int32(0))
											mBase = m.M
											v117 = m.ExcPending
											if v117 != 0 {
												return int32(0)
											} else {
												return v116
											}
										} else {
											v120 = F_MB_MatchText(m, v109, v107, v108, v66, int32(0))
											mBase = m.M
											v121 = m.ExcPending
											if v121 != 0 {
												return int32(0)
											} else {
												return v120
											}
										}
									}
								}
							}
						}
					}
				} else {
					v31 = F_DirectFunctionCall1Coll(m, int32(1437), l2, l1)
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return int32(0)
					} else {
						v33 = F_pg_detoast_datum_packed(m, v31)
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
							return int32(0)
						} else {
							v35 = int32(1)
							v36 = v33 + v35
							v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33))))
							v39 = v37 & v35
							if v37 == v35 {
								v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36))))
								if v45 == int32(18) {
									v48 = int32(16)
								} else {
									v48 = int32(0)
								}
								if base.Ui32((v45-int32(1))&int32(255)) < base.Ui32(int32(3)) {
									v55 = int32(4)
								} else {
									v55 = v48
								}
								v66 = v55
							} else {
								v56 = int32(1)
								if v39 != 0 {
									v66 = int32(base.Ui32(v37)>>(uint(v56)%32)) - v56
								} else {
									v60 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
									v66 = int32(base.Ui32(v60)>>(uint(int32(2))%32)) - int32(4)
								}
							}
							v70 = F_DirectFunctionCall1Coll(m, int32(1437), l2, l0)
							mBase = m.M
							v71 = m.ExcPending
							if v71 != 0 {
								return int32(0)
							} else {
								v72 = F_pg_detoast_datum_packed(m, v70)
								mBase = m.M
								v73 = m.ExcPending
								if v73 != 0 {
									return int32(0)
								} else {
									v76 = int32(1)
									v77 = v72 + v76
									v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72))))
									v80 = v78 & v76
									if v78 == v76 {
										v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77))))
										if v86 == int32(18) {
											v89 = int32(16)
										} else {
											v89 = int32(0)
										}
										if base.Ui32((v86-int32(1))&int32(255)) < base.Ui32(int32(3)) {
											v96 = int32(4)
										} else {
											v96 = v89
										}
										v107 = v96
									} else {
										v97 = int32(1)
										if v80 != 0 {
											v107 = int32(base.Ui32(v78)>>(uint(v97)%32)) - v97
										} else {
											v101 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
											v107 = int32(base.Ui32(v101)>>(uint(int32(2))%32)) - int32(4)
										}
									}
									if v39 != 0 {
										v108 = v36
									} else {
										v108 = v33 + int32(4)
									}
									if v80 != 0 {
										v109 = v77
									} else {
										v109 = v72 + int32(4)
									}
									v111 = *(*int32)(unsafe.Add(mBase, _c_F_Generic_Text_IC_like[0]))
									v112 = *(*int32)(unsafe.Add(mBase, uint32(v111)+4))
									if v112 == int32(6) {
										v116 = F_UTF8_MatchText(m, v109, v107, v108, v66, int32(0))
										mBase = m.M
										v117 = m.ExcPending
										if v117 != 0 {
											return int32(0)
										} else {
											return v116
										}
									} else {
										v120 = F_MB_MatchText(m, v109, v107, v108, v66, int32(0))
										mBase = m.M
										v121 = m.ExcPending
										if v121 != 0 {
											return int32(0)
										} else {
											return v120
										}
									}
								}
							}
						}
					}
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v203 = m.ExcPending
		if v203 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(34209924))
			mBase = m.M
			v206 = m.ExcPending
			if v206 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(_a_F_Generic_Text_IC_like_3), int32(0))
				mBase = m.M
				v210 = m.ExcPending
				if v210 != 0 {
					return int32(0)
				} else {
					F_errhint(m, int32(_a_F_Generic_Text_IC_like_4), int32(0))
					mBase = m.M
					v214 = m.ExcPending
					if v214 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_Generic_Text_IC_like_1), int32(194), int32(_a_F_Generic_Text_IC_like_2))
						mBase = m.M
						v219 = m.ExcPending
						if v219 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		}
	}
}
func F_generic_restriction_selectivity(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 float64) float64 {
	mBase := m.M
	_ = mBase
	var v8 float64
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 float64
	_ = v52
	var v53 int32
	_ = v53
	var v56 float64
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v65 float64
	_ = v65
	var v72 float64
	_ = v72
	var v74 float64
	_ = v74
	var v82 float64
	_ = v82
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 float32
	_ = v88
	var v92 float64
	_ = v92
	var v94 float64
	_ = v94
	var v101 int32
	_ = v101
	var v102 float64
	_ = v102
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v107 float64
	_ = v107
	var v120 float64
	_ = v120
	v8 = float64(0)
	v10 = m.G0
	v12 = v10 - int32(96)
	m.G0 = v12
	v20 = F_get_restriction_variable(m, l0, l3, l4, v12-int32(-64), v12+int32(60), v12+int32(59))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		return float64(0)
	} else {
		if v20 == int32(0) {
			v120 = l5
			m.G0 = v12 + int32(96)
			return v120
		} else {
			v26 = *(*int32)(unsafe.Add(mBase, uint32(v12)+60))
			v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
			if v27 != int32(7) {
				v30 = *(*int32)(unsafe.Add(mBase, uint32(v12)+72))
				v101 = v30
				v102 = l5
				if v101 != 0 {
					v104 = *(*int32)(unsafe.Add(mBase, uint32(v12)+76))
					m.T0[v104].(func(*base.Module, int32))(m, v101)
					mBase = m.M
					v106 = m.ExcPending
					if v106 != 0 {
						return float64(0)
					} else {
						v107 = float64(0)
						if base.F64_lt(v102, v107) != 0 {
							v120 = v107
						} else {
							if base.F64_gt(v102, float64(1)) == int32(0) {
								v120 = v102
							} else {
								v120 = float64(1)
							}
						}
						m.G0 = v12 + int32(96)
						return v120
					}
				} else {
					v107 = float64(0)
					if base.F64_lt(v102, v107) != 0 {
						v120 = v107
					} else {
						if base.F64_gt(v102, float64(1)) == int32(0) {
							v120 = v102
						} else {
							v120 = float64(1)
						}
					}
					m.G0 = v12 + int32(96)
					return v120
				}
			} else {
				v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+24)))
				if v31 == int32(1) {
					v34 = *(*int32)(unsafe.Add(mBase, uint32(v12)+72))
					if v34 == int32(0) {
						v120 = v8
						m.G0 = v12 + int32(96)
						return v120
					} else {
						v37 = *(*int32)(unsafe.Add(mBase, uint32(v12)+76))
						m.T0[v37].(func(*base.Module, int32))(m, v34)
						mBase = m.M
						v39 = m.ExcPending
						if v39 != 0 {
							return float64(0)
						} else {
							v120 = v8
							m.G0 = v12 + int32(96)
							return v120
						}
					}
				} else {
					v40 = *(*int32)(unsafe.Add(mBase, uint32(v26)+20))
					v41 = F_get_opcode(m, l1)
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return float64(0)
					} else {
						v44 = v12 + int32(28)
						F_fmgr_info(m, v41, v44)
						mBase = m.M
						v46 = m.ExcPending
						if v46 != 0 {
							return float64(0)
						} else {
							v48 = v12 - int32(-64)
							v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+59)))
							v52 = F_mcv_selectivity(m, v48, v44, l2, v40, v49, v12+int32(16))
							mBase = m.M
							v53 = m.ExcPending
							if v53 != 0 {
								return float64(0)
							} else {
								v56 = F_histogram_selectivity(m, v48, v44, l2, v40, v49, v12+int32(12))
								mBase = m.M
								v57 = m.ExcPending
								if v57 != 0 {
									return float64(0)
								} else {
									if base.F64_lt(v56, float64(0)) != 0 {
										v72 = l5
									} else {
										v60 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
										if int32(99) < v60 {
											v72 = v56
										} else {
											v65 = base.F64_div(base.F64_convert_i32_s(v60), float64(100))
											v72 = base.F64_add(base.F64_mul(v56, v65), base.F64_mul(l5, base.F64_sub(float64(1), v65)))
										}
									}
									v74 = float64(0.0001)
									if base.F64_lt(v72, v74) != 0 {
										v82 = v74
									} else {
										if base.F64_gt(v72, float64(0.9999)) == int32(0) {
											v82 = v72
										} else {
											v82 = float64(0.9999)
										}
									}
									v84 = *(*int32)(unsafe.Add(mBase, uint32(v12)+72))
									if v84 != 0 {
										v85 = *(*int32)(unsafe.Add(mBase, uint32(v84)+16))
										v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85)+22)))
										v88 = *(*float32)(unsafe.Add(mBase, uint32(v85+v86)+8))
										v92 = base.F64_promote_f32(v88)
									} else {
										v92 = float64(0)
									}
									v94 = *(*float64)(unsafe.Add(mBase, uint32(v12)+16))
									v101 = v84
									v102 = base.F64_add(v52, base.F64_mul(v82, base.F64_sub(base.F64_sub(float64(1), v92), v94)))
									if v101 != 0 {
										v104 = *(*int32)(unsafe.Add(mBase, uint32(v12)+76))
										m.T0[v104].(func(*base.Module, int32))(m, v101)
										mBase = m.M
										v106 = m.ExcPending
										if v106 != 0 {
											return float64(0)
										} else {
											v107 = float64(0)
											if base.F64_lt(v102, v107) != 0 {
												v120 = v107
											} else {
												if base.F64_gt(v102, float64(1)) == int32(0) {
													v120 = v102
												} else {
													v120 = float64(1)
												}
											}
											m.G0 = v12 + int32(96)
											return v120
										}
									} else {
										v107 = float64(0)
										if base.F64_lt(v102, v107) != 0 {
											v120 = v107
										} else {
											if base.F64_gt(v102, float64(1)) == int32(0) {
												v120 = v102
											} else {
												v120 = float64(1)
											}
										}
										m.G0 = v12 + int32(96)
										return v120
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
