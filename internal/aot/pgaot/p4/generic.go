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
	var v457 int64
	_ = v457
	var v458 int32
	_ = v458
	var v462 int32
	_ = v462
	var v468 int32
	_ = v468
	var v470 int32
	_ = v470
	var v476 int32
	_ = v476
	var v478 int32
	_ = v478
	var v482 int32
	_ = v482
	var v489 int32
	_ = v489
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v498 int32
	_ = v498
	var v502 int32
	_ = v502
	var v509 int32
	_ = v509
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v518 int32
	_ = v518
	var v524 int32
	_ = v524
	var v532 int32
	_ = v532
	var v538 int32
	_ = v538
	var v540 int32
	_ = v540
	var v542 int32
	_ = v542
	var v546 int32
	_ = v546
	var v550 int32
	_ = v550
	var v556 int32
	_ = v556
	var v558 int32
	_ = v558
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v568 int32
	_ = v568
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v575 int32
	_ = v575
	var v582 int32
	_ = v582
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	var v590 int32
	_ = v590
	var v593 int32
	_ = v593
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v600 int32
	_ = v600
	var v607 int32
	_ = v607
	var v613 int32
	_ = v613
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v618 int32
	_ = v618
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v627 int32
	_ = v627
	var v634 int32
	_ = v634
	var v640 int32
	_ = v640
	var v641 int32
	_ = v641
	var v642 int32
	_ = v642
	var v645 int32
	_ = v645
	var v647 int32
	_ = v647
	var v659 int32
	_ = v659
	var v661 int32
	_ = v661
	var v666 int32
	_ = v666
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_GenericXLogFinish[0]))))
	if v12 == int32(1) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v659 = int32(_a_F_GenericXLogFinish_0)
	v661 = *(*int32)(unsafe.Add(mBase, _c_F_GenericXLogFinish[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_GenericXLogFinish[1])) = v661 - int32(1)
	F_pfree(m, l0)
	mBase = m.M
	v666 = m.ExcPending
	if v666 != 0 {
		goto L5
	} else {
		goto L217
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
	v540 = int32(_a_F_GenericXLogFinish_0)
	v542 = *(*int32)(unsafe.Add(mBase, _c_F_GenericXLogFinish[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_GenericXLogFinish[1])) = v542 + int32(1)
	v546 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_GenericXLogFinish[2])))
	if v546 != 0 {
		goto L187
	} else {
		goto L188
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
	base.MemoryFill(m, v412+v59, int32(0), v413)
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
	base.MemoryCopy(m, v418+v59, v421+v418, v419)
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
	v457 = base.I64_rotl(v454, int64(32))
	v458 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_GenericXLogFinish[2])))
	if v458 != 0 {
		goto L162
	} else {
		goto L163
	}
L162:
	;
	if v458 < int32(0) {
		goto L166
	} else {
		goto L167
	}
L163:
	;
	goto L164
L164:
	;
	v478 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_GenericXLogFinish[5])))
	if v478 != 0 {
		goto L169
	} else {
		goto L170
	}
L165:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v476))) = v457
	goto L164
L166:
	;
	v462 = *(*int32)(unsafe.Add(mBase, _c_F_GenericXLogFinish[3]))
	v468 = *(*int32)(unsafe.Add(mBase, uint32(v462+(v458^int32(-1))<<(uint(int32(2))%32))))
	v476 = v468
	goto L165
L167:
	;
	goto L168
L168:
	;
	v470 = *(*int32)(unsafe.Add(mBase, _c_F_GenericXLogFinish[4]))
	v476 = v470 + v458<<(uint(int32(13))%32) + int32(-8192)
	goto L165
L169:
	;
	if int32(0) <= v478 {
		goto L173
	} else {
		goto L174
	}
L170:
	;
	goto L171
L171:
	;
	v498 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_GenericXLogFinish[6])))
	if v498 != 0 {
		goto L176
	} else {
		goto L177
	}
L172:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v496))) = v457
	goto L171
L173:
	;
	v482 = *(*int32)(unsafe.Add(mBase, _c_F_GenericXLogFinish[4]))
	v496 = v482 + v478<<(uint(int32(13))%32) + int32(-8192)
	goto L172
L174:
	;
	goto L175
L175:
	;
	v489 = *(*int32)(unsafe.Add(mBase, _c_F_GenericXLogFinish[3]))
	v495 = *(*int32)(unsafe.Add(mBase, uint32(v489+(v478^int32(-1))<<(uint(int32(2))%32))))
	v496 = v495
	goto L172
L176:
	;
	if int32(0) <= v498 {
		goto L180
	} else {
		goto L181
	}
L177:
	;
	goto L178
L178:
	;
	v518 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_GenericXLogFinish[7])))
	if v518 == int32(0) {
		goto L1
	} else {
		goto L183
	}
L179:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v516))) = v457
	goto L178
L180:
	;
	v502 = *(*int32)(unsafe.Add(mBase, _c_F_GenericXLogFinish[4]))
	v516 = v502 + v498<<(uint(int32(13))%32) + int32(-8192)
	goto L179
L181:
	;
	goto L182
L182:
	;
	v509 = *(*int32)(unsafe.Add(mBase, _c_F_GenericXLogFinish[3]))
	v515 = *(*int32)(unsafe.Add(mBase, uint32(v509+(v498^int32(-1))<<(uint(int32(2))%32))))
	v516 = v515
	goto L179
L183:
	;
	if int32(0) <= v518 {
		goto L184
	} else {
		goto L185
	}
L184:
	;
	v524 = *(*int32)(unsafe.Add(mBase, _c_F_GenericXLogFinish[4]))
	*(*int64)(unsafe.Add(mBase, uint32(v524+v518<<(uint(int32(13))%32))+uint32(_c_F_GenericXLogFinish[8]))) = v457
	goto L1
L185:
	;
	goto L186
L186:
	;
	v532 = *(*int32)(unsafe.Add(mBase, _c_F_GenericXLogFinish[3]))
	v538 = *(*int32)(unsafe.Add(mBase, uint32(v532+(v518^int32(-1))<<(uint(int32(2))%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v538))) = v457
	goto L1
L187:
	;
	if v546 < int32(0) {
		goto L191
	} else {
		goto L192
	}
L188:
	;
	goto L189
L189:
	;
	v571 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_GenericXLogFinish[5])))
	if v571 != 0 {
		goto L195
	} else {
		goto L196
	}
L190:
	;
	v565 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_GenericXLogFinish[9])))
	base.MemoryCopy(m, v564, v565, int32(_a_F_GenericXLogFinish_3))
	v568 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_GenericXLogFinish[2])))
	F_MarkBufferDirty(m, v568)
	mBase = m.M
	v570 = m.ExcPending
	if v570 != 0 {
		goto L5
	} else {
		goto L194
	}
L191:
	;
	v550 = *(*int32)(unsafe.Add(mBase, _c_F_GenericXLogFinish[3]))
	v556 = *(*int32)(unsafe.Add(mBase, uint32(v550+(v546^int32(-1))<<(uint(int32(2))%32))))
	v564 = v556
	goto L190
L192:
	;
	goto L193
L193:
	;
	v558 = *(*int32)(unsafe.Add(mBase, _c_F_GenericXLogFinish[4]))
	v564 = v558 + v546<<(uint(int32(13))%32) + int32(-8192)
	goto L190
L194:
	;
	goto L189
L195:
	;
	if int32(0) <= v571 {
		goto L199
	} else {
		goto L200
	}
L196:
	;
	goto L197
L197:
	;
	v596 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_GenericXLogFinish[6])))
	if v596 != 0 {
		goto L203
	} else {
		goto L204
	}
L198:
	;
	v590 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_GenericXLogFinish[10])))
	base.MemoryCopy(m, v589, v590, int32(_a_F_GenericXLogFinish_3))
	v593 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_GenericXLogFinish[5])))
	F_MarkBufferDirty(m, v593)
	mBase = m.M
	v595 = m.ExcPending
	if v595 != 0 {
		goto L5
	} else {
		goto L202
	}
L199:
	;
	v575 = *(*int32)(unsafe.Add(mBase, _c_F_GenericXLogFinish[4]))
	v589 = v575 + v571<<(uint(int32(13))%32) + int32(-8192)
	goto L198
L200:
	;
	goto L201
L201:
	;
	v582 = *(*int32)(unsafe.Add(mBase, _c_F_GenericXLogFinish[3]))
	v588 = *(*int32)(unsafe.Add(mBase, uint32(v582+(v571^int32(-1))<<(uint(int32(2))%32))))
	v589 = v588
	goto L198
L202:
	;
	goto L197
L203:
	;
	if int32(0) <= v596 {
		goto L207
	} else {
		goto L208
	}
L204:
	;
	goto L205
L205:
	;
	v621 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_GenericXLogFinish[7])))
	if v621 == int32(0) {
		goto L1
	} else {
		goto L211
	}
L206:
	;
	v615 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_GenericXLogFinish[11])))
	base.MemoryCopy(m, v614, v615, int32(_a_F_GenericXLogFinish_3))
	v618 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_GenericXLogFinish[6])))
	F_MarkBufferDirty(m, v618)
	mBase = m.M
	v620 = m.ExcPending
	if v620 != 0 {
		goto L5
	} else {
		goto L210
	}
L207:
	;
	v600 = *(*int32)(unsafe.Add(mBase, _c_F_GenericXLogFinish[4]))
	v614 = v600 + v596<<(uint(int32(13))%32) + int32(-8192)
	goto L206
L208:
	;
	goto L209
L209:
	;
	v607 = *(*int32)(unsafe.Add(mBase, _c_F_GenericXLogFinish[3]))
	v613 = *(*int32)(unsafe.Add(mBase, uint32(v607+(v596^int32(-1))<<(uint(int32(2))%32))))
	v614 = v613
	goto L206
L210:
	;
	goto L205
L211:
	;
	if int32(0) <= v621 {
		goto L213
	} else {
		goto L214
	}
L212:
	;
	v642 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_GenericXLogFinish[12])))
	base.MemoryCopy(m, v641, v642, int32(_a_F_GenericXLogFinish_3))
	v645 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_GenericXLogFinish[7])))
	F_MarkBufferDirty(m, v645)
	mBase = m.M
	v647 = m.ExcPending
	if v647 != 0 {
		goto L5
	} else {
		goto L216
	}
L213:
	;
	v627 = *(*int32)(unsafe.Add(mBase, _c_F_GenericXLogFinish[4]))
	v641 = v627 + v621<<(uint(int32(13))%32) + int32(-8192)
	goto L212
L214:
	;
	goto L215
L215:
	;
	v634 = *(*int32)(unsafe.Add(mBase, _c_F_GenericXLogFinish[3]))
	v640 = *(*int32)(unsafe.Add(mBase, uint32(v634+(v621^int32(-1))<<(uint(int32(2))%32))))
	v641 = v640
	goto L212
L216:
	;
	goto L1
L217:
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
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v105 int64
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v135 int32
	_ = v135
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v149 int64
	_ = v149
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v244 int32
	_ = v244
	var v249 int32
	_ = v249
	if l2 != 0 {
		v8 = F_pg_newlocale_from_collation(m, l2)
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8))))
			if v12 == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v237 = m.ExcPending
				if v237 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(1088))
					mBase = m.M
					v240 = m.ExcPending
					if v240 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(_a_F_Generic_Text_IC_like_0), int32(0))
						mBase = m.M
						v244 = m.ExcPending
						if v244 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_Generic_Text_IC_like_1), int32(190), int32(_a_F_Generic_Text_IC_like_2))
							mBase = m.M
							v249 = m.ExcPending
							if v249 != 0 {
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
				v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+2)))
				if v15 != int32(1) {
					v105 = F_DirectFunctionCall1Coll(m, int32(1557), l2, base.I64_extend_i32_u(l1))
					mBase = m.M
					v106 = m.ExcPending
					if v106 != 0 {
						return int32(0)
					} else {
						v108 = F_pg_detoast_datum_packed(m, base.I32_wrap_i64(v105))
						mBase = m.M
						v109 = m.ExcPending
						if v109 != 0 {
							return int32(0)
						} else {
							v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108))))
							v111 = int32(1)
							if v110 == v111 {
								v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108)+1)))
								if v118 == int32(18) {
									v121 = int32(16)
								} else {
									v121 = int32(0)
								}
								if base.Ui32((v118-int32(1))&int32(255)) < base.Ui32(int32(3)) {
									v128 = int32(4)
								} else {
									v128 = v121
								}
								v141 = v128
							} else {
								v129 = int32(1)
								if v110&v129 != 0 {
									v141 = int32(base.Ui32(v110)>>(uint(v129)%32)) - v129
								} else {
									v135 = *(*int32)(unsafe.Add(mBase, uint32(v108)))
									v141 = int32(base.Ui32(v135)>>(uint(int32(2))%32)) - int32(4)
								}
							}
							if v110&v111 != 0 {
								v144 = int32(1)
							} else {
								v144 = int32(4)
							}
							v149 = F_DirectFunctionCall1Coll(m, int32(1557), l2, base.I64_extend_i32_u(l0))
							mBase = m.M
							v150 = m.ExcPending
							if v150 != 0 {
								return int32(0)
							} else {
								v152 = F_pg_detoast_datum_packed(m, base.I32_wrap_i64(v149))
								mBase = m.M
								v153 = m.ExcPending
								if v153 != 0 {
									return int32(0)
								} else {
									v154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v152))))
									v156 = v154 & int32(1)
									if v156 != 0 {
										v157 = int32(1)
									} else {
										v157 = int32(4)
									}
									if v154 == int32(1) {
										v163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v152)+1)))
										if v163 == int32(18) {
											v166 = int32(16)
										} else {
											v166 = int32(0)
										}
										if base.Ui32((v163-int32(1))&int32(255)) < base.Ui32(int32(3)) {
											v173 = int32(4)
										} else {
											v173 = v166
										}
										v184 = v173
									} else {
										v174 = int32(1)
										if v156 != 0 {
											v184 = int32(base.Ui32(v154)>>(uint(v174)%32)) - v174
										} else {
											v178 = *(*int32)(unsafe.Add(mBase, uint32(v152)))
											v184 = int32(base.Ui32(v178)>>(uint(int32(2))%32)) - int32(4)
										}
									}
									v185 = v108 + v144
									v186 = v152 + v157
									v188 = *(*int32)(unsafe.Add(mBase, _c_F_Generic_Text_IC_like[0]))
									v189 = *(*int32)(unsafe.Add(mBase, uint32(v188)+4))
									if v189 == int32(6) {
										v193 = F_UTF8_MatchText(m, v186, v184, v185, v141, int32(0))
										mBase = m.M
										v194 = m.ExcPending
										if v194 != 0 {
											return int32(0)
										} else {
											return v193
										}
									} else {
										v197 = *(*int32)(unsafe.Add(mBase, _c_F_Generic_Text_IC_like[0]))
										v198 = *(*int32)(unsafe.Add(mBase, uint32(v197)+4))
										v203 = *(*int32)(unsafe.Add(mBase, uint32(v198*int32(28))+uint32(_c_F_Generic_Text_IC_like[1])))
										if int32(2) <= v203 {
											v207 = F_MB_MatchText(m, v186, v184, v185, v141, int32(0))
											mBase = m.M
											v208 = m.ExcPending
											if v208 != 0 {
												return int32(0)
											} else {
												return v207
											}
										} else {
											v211 = F_SB_MatchText(m, v186, v184, v185, v141, int32(0))
											mBase = m.M
											v212 = m.ExcPending
											if v212 != 0 {
												return int32(0)
											} else {
												return v211
											}
										}
									}
								}
							}
						}
					}
				} else {
					v19 = *(*int32)(unsafe.Add(mBase, _c_F_Generic_Text_IC_like[0]))
					v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
					v25 = *(*int32)(unsafe.Add(mBase, uint32(v20*int32(28))+uint32(_c_F_Generic_Text_IC_like[1])))
					if v25 != int32(1) {
						v105 = F_DirectFunctionCall1Coll(m, int32(1557), l2, base.I64_extend_i32_u(l1))
						mBase = m.M
						v106 = m.ExcPending
						if v106 != 0 {
							return int32(0)
						} else {
							v108 = F_pg_detoast_datum_packed(m, base.I32_wrap_i64(v105))
							mBase = m.M
							v109 = m.ExcPending
							if v109 != 0 {
								return int32(0)
							} else {
								v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108))))
								v111 = int32(1)
								if v110 == v111 {
									v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108)+1)))
									if v118 == int32(18) {
										v121 = int32(16)
									} else {
										v121 = int32(0)
									}
									if base.Ui32((v118-int32(1))&int32(255)) < base.Ui32(int32(3)) {
										v128 = int32(4)
									} else {
										v128 = v121
									}
									v141 = v128
								} else {
									v129 = int32(1)
									if v110&v129 != 0 {
										v141 = int32(base.Ui32(v110)>>(uint(v129)%32)) - v129
									} else {
										v135 = *(*int32)(unsafe.Add(mBase, uint32(v108)))
										v141 = int32(base.Ui32(v135)>>(uint(int32(2))%32)) - int32(4)
									}
								}
								if v110&v111 != 0 {
									v144 = int32(1)
								} else {
									v144 = int32(4)
								}
								v149 = F_DirectFunctionCall1Coll(m, int32(1557), l2, base.I64_extend_i32_u(l0))
								mBase = m.M
								v150 = m.ExcPending
								if v150 != 0 {
									return int32(0)
								} else {
									v152 = F_pg_detoast_datum_packed(m, base.I32_wrap_i64(v149))
									mBase = m.M
									v153 = m.ExcPending
									if v153 != 0 {
										return int32(0)
									} else {
										v154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v152))))
										v156 = v154 & int32(1)
										if v156 != 0 {
											v157 = int32(1)
										} else {
											v157 = int32(4)
										}
										if v154 == int32(1) {
											v163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v152)+1)))
											if v163 == int32(18) {
												v166 = int32(16)
											} else {
												v166 = int32(0)
											}
											if base.Ui32((v163-int32(1))&int32(255)) < base.Ui32(int32(3)) {
												v173 = int32(4)
											} else {
												v173 = v166
											}
											v184 = v173
										} else {
											v174 = int32(1)
											if v156 != 0 {
												v184 = int32(base.Ui32(v154)>>(uint(v174)%32)) - v174
											} else {
												v178 = *(*int32)(unsafe.Add(mBase, uint32(v152)))
												v184 = int32(base.Ui32(v178)>>(uint(int32(2))%32)) - int32(4)
											}
										}
										v185 = v108 + v144
										v186 = v152 + v157
										v188 = *(*int32)(unsafe.Add(mBase, _c_F_Generic_Text_IC_like[0]))
										v189 = *(*int32)(unsafe.Add(mBase, uint32(v188)+4))
										if v189 == int32(6) {
											v193 = F_UTF8_MatchText(m, v186, v184, v185, v141, int32(0))
											mBase = m.M
											v194 = m.ExcPending
											if v194 != 0 {
												return int32(0)
											} else {
												return v193
											}
										} else {
											v197 = *(*int32)(unsafe.Add(mBase, _c_F_Generic_Text_IC_like[0]))
											v198 = *(*int32)(unsafe.Add(mBase, uint32(v197)+4))
											v203 = *(*int32)(unsafe.Add(mBase, uint32(v198*int32(28))+uint32(_c_F_Generic_Text_IC_like[1])))
											if int32(2) <= v203 {
												v207 = F_MB_MatchText(m, v186, v184, v185, v141, int32(0))
												mBase = m.M
												v208 = m.ExcPending
												if v208 != 0 {
													return int32(0)
												} else {
													return v207
												}
											} else {
												v211 = F_SB_MatchText(m, v186, v184, v185, v141, int32(0))
												mBase = m.M
												v212 = m.ExcPending
												if v212 != 0 {
													return int32(0)
												} else {
													return v211
												}
											}
										}
									}
								}
							}
						}
					} else {
						v28 = int32(1)
						v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
						v32 = v30 & v28
						if v32 != 0 {
							v33 = v28
						} else {
							v33 = int32(4)
						}
						if v30 == int32(1) {
							v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
							if v39 == int32(18) {
								v42 = int32(16)
							} else {
								v42 = int32(0)
							}
							if base.Ui32((v39-int32(1))&int32(255)) < base.Ui32(int32(3)) {
								v49 = int32(4)
							} else {
								v49 = v42
							}
							v60 = v49
						} else {
							v50 = int32(1)
							if v32 != 0 {
								v60 = int32(base.Ui32(v30)>>(uint(v50)%32)) - v50
							} else {
								v54 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
								v60 = int32(base.Ui32(v54)>>(uint(int32(2))%32)) - int32(4)
							}
						}
						v61 = l1 + v33
						v62 = int32(1)
						v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
						v66 = v64 & v62
						if v66 != 0 {
							v67 = v62
						} else {
							v67 = int32(4)
						}
						v68 = l0 + v67
						if v64 == int32(1) {
							v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
							if v74 == int32(18) {
								v77 = int32(16)
							} else {
								v77 = int32(0)
							}
							if base.Ui32((v74-int32(1))&int32(255)) < base.Ui32(int32(3)) {
								v84 = int32(4)
							} else {
								v84 = v77
							}
							v85 = F_C_IMatchText(m, v68, v84, v61, v60, v8)
							mBase = m.M
							v86 = m.ExcPending
							if v86 != 0 {
								return int32(0)
							} else {
								return v85
							}
						} else {
							if v66 != 0 {
								v88 = int32(1)
								v92 = F_C_IMatchText(m, v68, int32(base.Ui32(v64)>>(uint(v88)%32))-v88, v61, v60, v8)
								mBase = m.M
								v93 = m.ExcPending
								if v93 != 0 {
									return int32(0)
								} else {
									return v92
								}
							} else {
								v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								v100 = F_C_IMatchText(m, v68, int32(base.Ui32(v95)>>(uint(int32(2))%32))-int32(4), v61, v60, v8)
								mBase = m.M
								v101 = m.ExcPending
								if v101 != 0 {
									return int32(0)
								} else {
									return v100
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
		v217 = m.ExcPending
		if v217 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(34209924))
			mBase = m.M
			v220 = m.ExcPending
			if v220 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(_a_F_Generic_Text_IC_like_3), int32(0))
				mBase = m.M
				v224 = m.ExcPending
				if v224 != 0 {
					return int32(0)
				} else {
					F_errhint(m, int32(_a_F_Generic_Text_IC_like_4), int32(0))
					mBase = m.M
					v228 = m.ExcPending
					if v228 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_Generic_Text_IC_like_1), int32(182), int32(_a_F_Generic_Text_IC_like_2))
						mBase = m.M
						v233 = m.ExcPending
						if v233 != 0 {
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
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int64
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 float64
	_ = v53
	var v54 int32
	_ = v54
	var v57 float64
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v66 float64
	_ = v66
	var v73 float64
	_ = v73
	var v75 float64
	_ = v75
	var v83 float64
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 float32
	_ = v89
	var v93 float64
	_ = v93
	var v95 float64
	_ = v95
	var v102 int32
	_ = v102
	var v103 float64
	_ = v103
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 float64
	_ = v109
	var v122 float64
	_ = v122
	v8 = float64(0)
	v11 = m.G0
	v13 = v11 - int32(96)
	m.G0 = v13
	v21 = F_get_restriction_variable(m, l0, l3, l4, v13-int32(-64), v13+int32(60), v13+int32(59))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		return float64(0)
	} else {
		if v21 == int32(0) {
			v122 = l5
			m.G0 = v13 + int32(96)
			return v122
		} else {
			v27 = *(*int32)(unsafe.Add(mBase, uint32(v13)+60))
			v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
			if v28 != int32(7) {
				v31 = *(*int32)(unsafe.Add(mBase, uint32(v13)+72))
				v102 = v31
				v103 = l5
				if v102 != 0 {
					v106 = *(*int32)(unsafe.Add(mBase, uint32(v13)+76))
					m.T0[v106].(func(*base.Module, int32))(m, v102)
					mBase = m.M
					v108 = m.ExcPending
					if v108 != 0 {
						return float64(0)
					} else {
						v109 = float64(0)
						if base.F64_lt(v103, v109) != 0 {
							v122 = v109
						} else {
							if base.F64_gt(v103, float64(1)) == int32(0) {
								v122 = v103
							} else {
								v122 = float64(1)
							}
						}
						m.G0 = v13 + int32(96)
						return v122
					}
				} else {
					v109 = float64(0)
					if base.F64_lt(v103, v109) != 0 {
						v122 = v109
					} else {
						if base.F64_gt(v103, float64(1)) == int32(0) {
							v122 = v103
						} else {
							v122 = float64(1)
						}
					}
					m.G0 = v13 + int32(96)
					return v122
				}
			} else {
				v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+32)))
				if v32 == int32(1) {
					v35 = *(*int32)(unsafe.Add(mBase, uint32(v13)+72))
					if v35 == int32(0) {
						v122 = v8
						m.G0 = v13 + int32(96)
						return v122
					} else {
						v38 = *(*int32)(unsafe.Add(mBase, uint32(v13)+76))
						m.T0[v38].(func(*base.Module, int32))(m, v35)
						mBase = m.M
						v40 = m.ExcPending
						if v40 != 0 {
							return float64(0)
						} else {
							v122 = v8
							m.G0 = v13 + int32(96)
							return v122
						}
					}
				} else {
					v41 = *(*int64)(unsafe.Add(mBase, uint32(v27)+24))
					v42 = F_get_opcode(m, l1)
					mBase = m.M
					v43 = m.ExcPending
					if v43 != 0 {
						return float64(0)
					} else {
						v45 = v13 + int32(28)
						F_fmgr_info(m, v42, v45)
						mBase = m.M
						v47 = m.ExcPending
						if v47 != 0 {
							return float64(0)
						} else {
							v49 = v13 - int32(-64)
							v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+59)))
							v53 = F_mcv_selectivity(m, v49, v45, l2, v41, v50, v13+int32(16))
							mBase = m.M
							v54 = m.ExcPending
							if v54 != 0 {
								return float64(0)
							} else {
								v57 = F_histogram_selectivity(m, v49, v45, l2, v41, v50, v13+int32(12))
								mBase = m.M
								v58 = m.ExcPending
								if v58 != 0 {
									return float64(0)
								} else {
									if base.F64_lt(v57, float64(0)) != 0 {
										v73 = l5
									} else {
										v61 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
										if int32(99) < v61 {
											v73 = v57
										} else {
											v66 = base.F64_div(base.F64_convert_i32_s(v61), float64(100))
											v73 = base.F64_add(base.F64_mul(v57, v66), base.F64_mul(l5, base.F64_sub(float64(1), v66)))
										}
									}
									v75 = float64(0.0001)
									if base.F64_lt(v73, v75) != 0 {
										v83 = v75
									} else {
										if base.F64_gt(v73, float64(0.9999)) == int32(0) {
											v83 = v73
										} else {
											v83 = float64(0.9999)
										}
									}
									v85 = *(*int32)(unsafe.Add(mBase, uint32(v13)+72))
									if v85 != 0 {
										v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)+16))
										v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86)+22)))
										v89 = *(*float32)(unsafe.Add(mBase, uint32(v86+v87)+8))
										v93 = base.F64_promote_f32(v89)
									} else {
										v93 = float64(0)
									}
									v95 = *(*float64)(unsafe.Add(mBase, uint32(v13)+16))
									v102 = v85
									v103 = base.F64_add(v53, base.F64_mul(v83, base.F64_sub(base.F64_sub(float64(1), v93), v95)))
									if v102 != 0 {
										v106 = *(*int32)(unsafe.Add(mBase, uint32(v13)+76))
										m.T0[v106].(func(*base.Module, int32))(m, v102)
										mBase = m.M
										v108 = m.ExcPending
										if v108 != 0 {
											return float64(0)
										} else {
											v109 = float64(0)
											if base.F64_lt(v103, v109) != 0 {
												v122 = v109
											} else {
												if base.F64_gt(v103, float64(1)) == int32(0) {
													v122 = v103
												} else {
													v122 = float64(1)
												}
											}
											m.G0 = v13 + int32(96)
											return v122
										}
									} else {
										v109 = float64(0)
										if base.F64_lt(v103, v109) != 0 {
											v122 = v109
										} else {
											if base.F64_gt(v103, float64(1)) == int32(0) {
												v122 = v103
											} else {
												v122 = float64(1)
											}
										}
										m.G0 = v13 + int32(96)
										return v122
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
