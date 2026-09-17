package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_FreePageManagerGetInternal(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v104 int32
	_ = v104
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
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
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v276 int32
	_ = v276
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v319 int32
	_ = v319
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v342 int32
	_ = v342
	var v348 int32
	_ = v348
	var v353 int32
	_ = v353
	var v356 int32
	_ = v356
	var v358 int32
	_ = v358
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v377 int32
	_ = v377
	var v389 int32
	_ = v389
	var v394 int32
	_ = v394
	var v398 int32
	_ = v398
	var v401 int32
	_ = v401
	var v404 int32
	_ = v404
	var v409 int32
	_ = v409
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v418 int32
	_ = v418
	var v428 int32
	_ = v428
	var v431 int32
	_ = v431
	var v433 int32
	_ = v433
	var v435 int32
	_ = v435
	var v437 int32
	_ = v437
	var v439 int32
	_ = v439
	var v442 int32
	_ = v442
	var v445 int32
	_ = v445
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v469 int32
	_ = v469
	var v475 int32
	_ = v475
	var v479 int32
	_ = v479
	var v484 int32
	_ = v484
	var v486 int32
	_ = v486
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v535 int32
	_ = v535
	var v538 int32
	_ = v538
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v550 int32
	_ = v550
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v566 int32
	_ = v566
	var v568 int32
	_ = v568
	var v570 int32
	_ = v570
	var v572 int32
	_ = v572
	var v576 int32
	_ = v576
	var v580 int32
	_ = v580
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v600 int32
	_ = v600
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v606 int32
	_ = v606
	var v612 int32
	_ = v612
	var v616 int32
	_ = v616
	var v621 int32
	_ = v621
	var v623 int32
	_ = v623
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v652 int32
	_ = v652
	var v660 int32
	_ = v660
	var v671 int32
	_ = v671
	var v672 int32
	_ = v672
	var v673 int32
	_ = v673
	var v676 int32
	_ = v676
	var v681 int32
	_ = v681
	var v682 int32
	_ = v682
	var v684 int32
	_ = v684
	var v685 int32
	_ = v685
	var v688 int32
	_ = v688
	var v689 int32
	_ = v689
	var v693 int32
	_ = v693
	var v695 int32
	_ = v695
	var v699 int32
	_ = v699
	var v704 int32
	_ = v704
	var v707 int32
	_ = v707
	var v736 int32
	_ = v736
	v4 = int32(0)
	v19 = m.G0
	v21 = v19 - int32(16)
	m.G0 = v21
	v23 = int32(129)
	if base.Ui32(v23) <= base.Ui32(l1) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v21 + int32(16)
	return v736
L2:
	;
	v26 = v23
	goto L4
L3:
	;
	v26 = l1
	goto L4
L4:
	;
	v28 = v26 - int32(1)
	if base.Ui32(int32(128)) < base.Ui32(v28) {
		v736 = v4
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v34 = l0 - v31 + int32(1)
	v40 = v28
	goto L7
L6:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v104)+12))
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v104)+8))
	if v115 != 0 {
		goto L27
	} else {
		goto L28
	}
L7:
	;
	v57 = l0 + int32(36) + v40<<(uint(int32(2))%32)
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
	if v58 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v104 = v58 + v34 - int32(1)
	goto L6
L9:
	;
	goto L8
L10:
	;
	if v40 != int32(128) {
		goto L9
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	v90 = v40 + int32(1)
	if v90 != int32(129) {
		v40 = v90
		goto L7
	} else {
		goto L25
	}
L13:
	;
	v68 = v58
	v69 = v4
	goto L14
L14:
	;
	v79 = v68 + v34
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)+3))
	if base.Ui32(v80) < base.Ui32(l1) {
		v87 = v69
		goto L16
	} else {
		goto L17
	}
L15:
	;
	if v87 != 0 {
		v104 = v87
		goto L6
	} else {
		goto L24
	}
L16:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v79)+11))
	if v88 != 0 {
		v68 = v88
		v69 = v87
		goto L14
	} else {
		goto L23
	}
L17:
	;
	if v69 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v69)+4))
	if base.Ui32(v82) <= base.Ui32(v80) {
		v87 = v69
		goto L16
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v85 = v79 - int32(1)
	if l1 == v80 {
		v104 = v85
		goto L6
	} else {
		goto L22
	}
L21:
	;
	goto L20
L22:
	;
	v87 = v85
	goto L16
L23:
	;
	goto L15
L24:
	;
	v736 = v4
	goto L1
L25:
	;
	v736 = v4
	goto L1
L26:
	;
	if v114 != 0 {
		goto L30
	} else {
		goto L31
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v115+v34)+11)) = v114
	goto L26
L28:
	;
	goto L29
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57))) = v114
	goto L26
L30:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v104)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v114+v34)+7)) = v120
	goto L32
L31:
	;
	goto L32
L32:
	;
	if v40 != int32(128) {
		goto L36
	} else {
		goto L37
	}
L33:
	;
	v139 = int32(base.Ui32(v104-v34) >> (uint(int32(12)) % 32))
	v140 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v140 == int32(0) {
		goto L43
	} else {
		goto L44
	}
L34:
	;
	v135 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)) = uint8(v135)
	goto L33
L35:
	;
	if v40+int32(1) != v129 {
		goto L33
	} else {
		goto L40
	}
L36:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v129 = v125
	goto L35
L37:
	;
	goto L38
L38:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v104)+4))
	if v126 == v127 {
		goto L34
	} else {
		goto L39
	}
L39:
	;
	v129 = v126
	goto L35
L40:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
	if v133 != 0 {
		goto L33
	} else {
		goto L41
	}
L41:
	;
	goto L34
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v139
	v736 = int32(1)
	goto L1
L43:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v144 = v143 + l1
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v144
	v146 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v147 = v146 - l1
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v147
	if l1 == v146 {
		goto L42
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v197 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v198 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+12)) = v198
	if v196 != 0 {
		goto L60
	} else {
		goto L61
	}
L46:
	;
	v150 = int32(129)
	if base.Ui32(v150) <= base.Ui32(v147) {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v153 = v150
	goto L49
L48:
	;
	v153 = v147
	goto L49
L49:
	;
	v158 = l0 + v153<<(uint(int32(2))%32) + int32(32)
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v158)))
	v160 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v162 = int32(1)
	v163 = l0 - v160 + v162
	v165 = v144 << (uint(int32(12)) % 32)
	v166 = v163 + v165
	*(*int32)(unsafe.Add(mBase, uint32(v166)+4)) = v147
	*(*int32)(unsafe.Add(mBase, uint32(v166))) = int32(-364896016)
	v170 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v166)+8)) = v170
	v172 = v163 + v159
	if v159 != 0 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v176 = v172 - v162
	goto L52
L51:
	;
	v176 = v170
	goto L52
L52:
	;
	if v159 != 0 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v181 = v176 - v163 + int32(1)
	goto L55
L54:
	;
	v181 = int32(0)
	goto L55
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v166)+12)) = v181
	v184 = v165 | int32(1)
	if v159 != 0 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v172)+7)) = v184
	goto L58
L57:
	;
	goto L58
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v158))) = v184
	goto L42
L59:
	;
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v104)+4))
	if l1 == v375 {
		goto L110
	} else {
		goto L111
	}
L60:
	;
	v203 = int32(1)
	v204 = l0 - v197 + v203
	v207 = v204 + v196 - v203
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v207)))
	if v208 == int32(430584521) {
		goto L63
	} else {
		goto L64
	}
L61:
	;
	goto L62
L62:
	;
	v358 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+8)) = uint8(v358)
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = v358
	goto L59
L63:
	;
	v216 = v207
	v217 = v198
	goto L66
L64:
	;
	v297 = int32(2)
	v298 = v207
	goto L65
L65:
	;
	v305 = int32(0)
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v298)+4))
	if base.Ui32(int32(509)) < base.Ui32(v307) {
		goto L91
	} else {
		goto L92
	}
L66:
	;
	v224 = v216 + int32(12)
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v216)+4))
	v227 = int32(0)
	v231 = v225
	goto L68
L67:
	;
	v297 = v271 + int32(1)
	v298 = v287
	goto L65
L68:
	;
	if base.Ui32(v231) <= base.Ui32(v227) {
		goto L71
	} else {
		goto L72
	}
L69:
	;
	if base.Ui32(v254) < base.Ui32(v225) {
		goto L81
	} else {
		goto L82
	}
L70:
	;
	goto L69
L71:
	;
	v254 = v227
	goto L70
L72:
	;
	goto L73
L73:
	;
	v241 = int32(1)
	v242 = int32(base.Ui32(v227+v231) >> (uint(v241) % 32))
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v224+v242<<(uint(int32(3))%32))))
	v249 = base.B2i32(base.Ui32(v139) < base.Ui32(v248))
	if base.Ui32(v139) < base.Ui32(v248) {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v250 = v227
	goto L76
L75:
	;
	v250 = v242 + v241
	goto L76
L76:
	;
	if base.Ui32(v139) < base.Ui32(v248) {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v251 = v242
	goto L79
L78:
	;
	v251 = v231
	goto L79
L79:
	;
	if v139 != v248 {
		v227 = v250
		v231 = v251
		goto L68
	} else {
		goto L80
	}
L80:
	;
	v254 = v242
	goto L70
L81:
	;
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v224+v254<<(uint(int32(3))%32))))
	v265 = base.B2i32(v263 != v139)
	goto L83
L82:
	;
	v265 = int32(1)
	goto L83
L83:
	;
	if base.Ui32(int32(509)) < base.Ui32(v225) {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	v271 = v217 + int32(1)
	goto L86
L85:
	;
	v271 = int32(0)
	goto L86
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+12)) = v271
	v273 = int32(3)
	v276 = int32(0)
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v216+v254<<(uint(v273)%32)-base.B2i32(v254 != v276)&v265<<(uint(v273)%32))+16))
	v285 = v204 + v282 - int32(1)
	if v282 != 0 {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v287 = v285
	goto L89
L88:
	;
	v287 = v276
	goto L89
L89:
	;
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v285)))
	if v288 == int32(430584521) {
		v216 = v287
		v217 = v271
		goto L66
	} else {
		goto L90
	}
L90:
	;
	goto L67
L91:
	;
	v310 = v297
	goto L93
L92:
	;
	v310 = v305
	goto L93
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+12)) = v310
	v313 = v298 + int32(12)
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v298)+4))
	v315 = v305
	v319 = v314
	goto L94
L94:
	;
	if base.Ui32(v319) <= base.Ui32(v315) {
		goto L97
	} else {
		goto L98
	}
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+4)) = v342
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = v298
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v298)+4))
	if base.Ui32(v342) < base.Ui32(v348) {
		goto L107
	} else {
		goto L108
	}
L96:
	;
	goto L95
L97:
	;
	v342 = v315
	goto L96
L98:
	;
	goto L99
L99:
	;
	v329 = int32(1)
	v330 = int32(base.Ui32(v315+v319) >> (uint(v329) % 32))
	v336 = *(*int32)(unsafe.Add(mBase, uint32(v313+v330<<(uint(int32(3))%32))))
	v337 = base.B2i32(base.Ui32(v139) < base.Ui32(v336))
	if base.Ui32(v139) < base.Ui32(v336) {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	v338 = v315
	goto L102
L101:
	;
	v338 = v330 + v329
	goto L102
L102:
	;
	if base.Ui32(v139) < base.Ui32(v336) {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	v339 = v330
	goto L105
L104:
	;
	v339 = v319
	goto L105
L105:
	;
	if v139 != v336 {
		v315 = v338
		v319 = v339
		goto L94
	} else {
		goto L106
	}
L106:
	;
	v342 = v330
	goto L96
L107:
	;
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v313+v342<<(uint(int32(3))%32))))
	v356 = base.B2i32(v139 == v353)
	goto L109
L108:
	;
	v356 = int32(0)
	goto L109
L109:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+8)) = uint8(v356)
	goto L59
L110:
	;
	v377 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v374)+4))
	if v389 == int32(1) {
		goto L114
	} else {
		goto L115
	}
L111:
	;
	goto L112
L112:
	;
	v527 = v374 + int32(12)
	v528 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	v531 = v527 + v528<<(uint(int32(3))%32)
	v532 = *(*int32)(unsafe.Add(mBase, uint32(v531)))
	*(*int32)(unsafe.Add(mBase, uint32(v531))) = v532 + l1
	v535 = *(*int32)(unsafe.Add(mBase, uint32(v531)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v531)+4)) = v535 - l1
	v538 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v528 == int32(0) {
		goto L152
	} else {
		goto L153
	}
L113:
	;
	goto L42
L114:
	;
	F_FreePageBtreeRemovePage(m, l0, v374)
	mBase = m.M
	goto L113
L115:
	;
	goto L116
L116:
	;
	v394 = v389 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v374)+4)) = v394
	if base.Ui32(v394) <= base.Ui32(v377) {
		goto L117
	} else {
		goto L118
	}
L117:
	;
	F_FreePageBtreeConsolidate(m, l0, v374)
	mBase = m.M
	goto L113
L118:
	;
	v398 = v374 + int32(12)
	v401 = (v394 - v377) << (uint(int32(3)) % 32)
	if v401 != 0 {
		goto L119
	} else {
		goto L120
	}
L119:
	;
	v404 = v398 + v377<<(uint(int32(3))%32)
	base.MemoryCopy(m, v404, v404+int32(8), v401)
	goto L121
L120:
	;
	goto L121
L121:
	;
	if v377 != 0 {
		goto L117
	} else {
		goto L122
	}
L122:
	;
	v409 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v412 = l0 - v409 + int32(1)
	v413 = *(*int32)(unsafe.Add(mBase, uint32(v398)))
	v418 = v374
	goto L123
L123:
	;
	v428 = *(*int32)(unsafe.Add(mBase, uint32(v418)+8))
	if v428 == int32(0) {
		goto L117
	} else {
		goto L125
	}
L124:
	;
	goto L117
L125:
	;
	v431 = v428 + v412
	v433 = v431 - int32(1)
	if v428 != 0 {
		goto L126
	} else {
		goto L127
	}
L126:
	;
	v435 = v433
	goto L128
L127:
	;
	v435 = int32(0)
	goto L128
L128:
	;
	v437 = v431 + int32(11)
	v439 = *(*int32)(unsafe.Add(mBase, uint32(v431)+3))
	v442 = int32(0)
	v445 = v439
	goto L129
L129:
	;
	if base.Ui32(v445) <= base.Ui32(v442) {
		goto L132
	} else {
		goto L133
	}
L130:
	;
	if base.Ui32(v469) < base.Ui32(v439) {
		goto L142
	} else {
		goto L143
	}
L131:
	;
	goto L130
L132:
	;
	v469 = v442
	goto L131
L133:
	;
	goto L134
L134:
	;
	v456 = int32(1)
	v457 = int32(base.Ui32(v442+v445) >> (uint(v456) % 32))
	v463 = *(*int32)(unsafe.Add(mBase, uint32(v437+v457<<(uint(int32(3))%32))))
	v464 = base.B2i32(base.Ui32(v413) < base.Ui32(v463))
	if base.Ui32(v413) < base.Ui32(v463) {
		goto L135
	} else {
		goto L136
	}
L135:
	;
	v465 = v442
	goto L137
L136:
	;
	v465 = v457 + v456
	goto L137
L137:
	;
	if base.Ui32(v413) < base.Ui32(v463) {
		goto L138
	} else {
		goto L139
	}
L138:
	;
	v466 = v457
	goto L140
L139:
	;
	v466 = v445
	goto L140
L140:
	;
	if v413 != v463 {
		v442 = v465
		v445 = v466
		goto L129
	} else {
		goto L141
	}
L141:
	;
	v469 = v457
	goto L131
L142:
	;
	v475 = int32(0)
	v479 = *(*int32)(unsafe.Add(mBase, uint32(v433+v469<<(uint(int32(3))%32))+16))
	if v479 != 0 {
		goto L145
	} else {
		goto L146
	}
L143:
	;
	v489 = int32(-1)
	goto L144
L144:
	;
	v490 = v489 + v469
	*(*int32)(unsafe.Add(mBase, uint32(v437+v490<<(uint(int32(3))%32)))) = v413
	if v490 == int32(0) {
		v418 = v435
		goto L123
	} else {
		goto L151
	}
L145:
	;
	v484 = v412 + v479 - int32(1)
	goto L147
L146:
	;
	v484 = v475
	goto L147
L147:
	;
	if v484 != v418 {
		goto L148
	} else {
		goto L149
	}
L148:
	;
	v486 = int32(-1)
	goto L150
L149:
	;
	v486 = v475
	goto L150
L150:
	;
	v489 = v486
	goto L144
L151:
	;
	goto L124
L152:
	;
	v543 = l0 - v538 + int32(1)
	v544 = *(*int32)(unsafe.Add(mBase, uint32(v527)))
	v550 = v374
	goto L155
L153:
	;
	v660 = v538
	goto L154
L154:
	;
	v671 = int32(129)
	v672 = *(*int32)(unsafe.Add(mBase, uint32(v104)+4))
	v673 = v672 - l1
	if base.Ui32(v671) <= base.Ui32(v673) {
		goto L186
	} else {
		goto L187
	}
L155:
	;
	v563 = *(*int32)(unsafe.Add(mBase, uint32(v550)+8))
	if v563 != 0 {
		goto L157
	} else {
		goto L158
	}
L156:
	;
	v652 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v660 = v652
	goto L154
L157:
	;
	v564 = v563 + v543
	v566 = v564 - int32(1)
	if v563 != 0 {
		goto L160
	} else {
		goto L161
	}
L158:
	;
	goto L159
L159:
	;
	goto L156
L160:
	;
	v568 = v566
	goto L162
L161:
	;
	v568 = int32(0)
	goto L162
L162:
	;
	v570 = v564 + int32(11)
	v572 = *(*int32)(unsafe.Add(mBase, uint32(v564)+3))
	v576 = int32(0)
	v580 = v572
	goto L163
L163:
	;
	if base.Ui32(v580) <= base.Ui32(v576) {
		goto L166
	} else {
		goto L167
	}
L164:
	;
	if base.Ui32(v606) < base.Ui32(v572) {
		goto L176
	} else {
		goto L177
	}
L165:
	;
	goto L164
L166:
	;
	v606 = v576
	goto L165
L167:
	;
	goto L168
L168:
	;
	v593 = int32(1)
	v594 = int32(base.Ui32(v576+v580) >> (uint(v593) % 32))
	v600 = *(*int32)(unsafe.Add(mBase, uint32(v570+v594<<(uint(int32(3))%32))))
	v601 = base.B2i32(base.Ui32(v544) < base.Ui32(v600))
	if base.Ui32(v544) < base.Ui32(v600) {
		goto L169
	} else {
		goto L170
	}
L169:
	;
	v602 = v576
	goto L171
L170:
	;
	v602 = v594 + v593
	goto L171
L171:
	;
	if base.Ui32(v544) < base.Ui32(v600) {
		goto L172
	} else {
		goto L173
	}
L172:
	;
	v603 = v594
	goto L174
L173:
	;
	v603 = v580
	goto L174
L174:
	;
	if v600 != v544 {
		v576 = v602
		v580 = v603
		goto L163
	} else {
		goto L175
	}
L175:
	;
	v606 = v594
	goto L165
L176:
	;
	v612 = int32(0)
	v616 = *(*int32)(unsafe.Add(mBase, uint32(v566+v606<<(uint(int32(3))%32))+16))
	if v616 != 0 {
		goto L179
	} else {
		goto L180
	}
L177:
	;
	v626 = int32(-1)
	goto L178
L178:
	;
	v627 = v626 + v606
	*(*int32)(unsafe.Add(mBase, uint32(v570+v627<<(uint(int32(3))%32)))) = v544
	if v627 == int32(0) {
		v550 = v568
		goto L155
	} else {
		goto L185
	}
L179:
	;
	v621 = v543 + v616 - int32(1)
	goto L181
L180:
	;
	v621 = v612
	goto L181
L181:
	;
	if v550 != v621 {
		goto L182
	} else {
		goto L183
	}
L182:
	;
	v623 = int32(-1)
	goto L184
L183:
	;
	v623 = v612
	goto L184
L184:
	;
	v626 = v623
	goto L178
L185:
	;
	goto L159
L186:
	;
	v676 = v671
	goto L188
L187:
	;
	v676 = v673
	goto L188
L188:
	;
	v681 = l0 + v676<<(uint(int32(2))%32) + int32(32)
	v682 = *(*int32)(unsafe.Add(mBase, uint32(v681)))
	v684 = int32(1)
	v685 = l0 - v660 + v684
	v688 = (l1 + v139) << (uint(int32(12)) % 32)
	v689 = v685 + v688
	*(*int32)(unsafe.Add(mBase, uint32(v689)+4)) = v673
	*(*int32)(unsafe.Add(mBase, uint32(v689))) = int32(-364896016)
	v693 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v689)+8)) = v693
	v695 = v682 + v685
	if v682 != 0 {
		goto L189
	} else {
		goto L190
	}
L189:
	;
	v699 = v695 - v684
	goto L191
L190:
	;
	v699 = v693
	goto L191
L191:
	;
	if v682 != 0 {
		goto L192
	} else {
		goto L193
	}
L192:
	;
	v704 = v699 - v685 + int32(1)
	goto L194
L193:
	;
	v704 = int32(0)
	goto L194
L194:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v689)+12)) = v704
	v707 = v688 | int32(1)
	if v682 != 0 {
		goto L195
	} else {
		goto L196
	}
L195:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v695)+7)) = v707
	goto L197
L196:
	;
	goto L197
L197:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v681))) = v707
	goto L42
}
func F_PageGetFreeSpaceForMultipleTuples(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	v3 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+14)))
	v4 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
	v5 = v3 - v4
	v7 = l1 << (uint(int32(2)) % 32)
	if v7 <= v5 {
		v11 = v5 - v7
	} else {
		v11 = int32(0)
	}
	return v11
}
func F_PageRepairFragmentation(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v102 int32
	_ = v102
	var v110 int32
	_ = v110
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v228 int32
	_ = v228
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v243 int32
	_ = v243
	var v249 int32
	_ = v249
	var v259 int32
	_ = v259
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v276 int32
	_ = v276
	var v282 int32
	_ = v282
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v311 int32
	_ = v311
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v327 int32
	_ = v327
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v357 int32
	_ = v357
	var v360 int32
	_ = v360
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v374 int32
	_ = v374
	var v376 int32
	_ = v376
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v390 int32
	_ = v390
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v431 int32
	_ = v431
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v444 int32
	_ = v444
	var v450 int32
	_ = v450
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v465 int32
	_ = v465
	var v474 int32
	_ = v474
	var v495 int32
	_ = v495
	var v497 int32
	_ = v497
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v514 int32
	_ = v514
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v524 int32
	_ = v524
	var v532 int32
	_ = v532
	var v535 int32
	_ = v535
	var v541 int32
	_ = v541
	var v546 int32
	_ = v546
	var v550 int32
	_ = v550
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v560 int32
	_ = v560
	var v565 int32
	_ = v565
	var v569 int32
	_ = v569
	var v572 int32
	_ = v572
	var v579 int32
	_ = v579
	var v584 int32
	_ = v584
	v2 = int32(0)
	v18 = m.G0
	v20 = v18 - int32(1808)
	m.G0 = v20
	v22 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
	v25 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+14)))
	v28 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+16)))
	if base.B2i32(base.Ui32(v22) < base.Ui32(int32(24)))|base.B2i32(base.Ui32(v25) < base.Ui32(v22))|(base.B2i32(base.Ui32(v28) < base.Ui32(v25))|base.B2i32(base.Ui32(int32(_a_F_PageRepairFragmentation_0)) < base.Ui32(v28)))|base.B2i32((v28+int32(7))&int32(_a_F_PageRepairFragmentation_1) != v28) == v2 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v569 = m.ExcPending
	if v569 != 0 {
		goto L112
	} else {
		goto L121
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v550 = m.ExcPending
	if v550 != 0 {
		goto L112
	} else {
		goto L117
	}
L3:
	;
	if base.Ui32(int32(25)) <= base.Ui32(v22) {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	goto L5
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v532 = m.ExcPending
	if v532 != 0 {
		goto L112
	} else {
		goto L113
	}
L6:
	;
	v505 = int32(_a_F_PageRepairFragmentation_2)
	v506 = v497 & v505
	if v506 != v49&v505 {
		goto L109
	} else {
		goto L110
	}
L7:
	;
	v49 = int32(base.Ui32(v22+int32(_a_F_PageRepairFragmentation_3)) >> (uint(int32(2)) % 32))
	goto L9
L8:
	;
	v49 = int32(0)
	goto L9
L9:
	;
	v51 = v49 & int32(_a_F_PageRepairFragmentation_2)
	if v51 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+14)) = uint16(v28)
	v495 = v2
	v497 = v2
	goto L6
L11:
	;
	goto L12
L12:
	;
	v59 = int32(1)
	v64 = v59
	v66 = v20 + int32(48)
	v68 = v2
	v69 = v28
	v70 = v2
	v72 = v59
	v74 = v2
	goto L13
L13:
	;
	v80 = l0 + int32(20) + v64<<(uint(int32(2))%32)
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
	if v81&int32(_a_F_PageRepairFragmentation_4) != 0 {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	if v20+int32(48) == v116 {
		goto L27
	} else {
		goto L28
	}
L15:
	;
	if v64 != v51 {
		v64 = v64 + int32(1)
		v66 = v116
		v68 = v117
		v69 = v118
		v70 = v119
		v72 = v120
		v74 = v121
		goto L13
	} else {
		goto L26
	}
L16:
	;
	if base.Ui32(v81) < base.Ui32(int32(_a_F_PageRepairFragmentation_5)) {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	goto L18
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v80))) = int32(0)
	v116 = v66
	v117 = v68 + int32(1)
	v118 = v69
	v119 = v70
	v120 = v72
	v121 = v74
	goto L15
L19:
	;
	v116 = v66
	v117 = v68
	v118 = v69
	v119 = v64
	v120 = v72
	v121 = v74
	goto L15
L20:
	;
	goto L21
L21:
	;
	v87 = v64 - int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v66))) = uint16(v87)
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
	v91 = v89 & int32(_a_F_PageRepairFragmentation_6)
	*(*uint16)(unsafe.Add(mBase, uint32(v66)+2)) = uint16(v91)
	if base.B2i32(base.Ui32(v91) < base.Ui32(v25))|base.B2i32(base.Ui32(v28) <= base.Ui32(v91)) != 0 {
		goto L2
	} else {
		goto L22
	}
L22:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
	v102 = (int32(base.Ui32(v96)>>(uint(int32(17))%32)) + int32(7)) & int32(_a_F_PageRepairFragmentation_7)
	*(*uint16)(unsafe.Add(mBase, uint32(v66)+4)) = uint16(v102)
	if v69 < v91 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v110 = v69
	goto L25
L24:
	;
	v110 = v91
	goto L25
L25:
	;
	v116 = v66 + int32(6)
	v117 = v68
	v118 = v110
	v119 = v64
	v120 = base.B2i32(v91 < v69) & v72
	v121 = v102 + v74
	goto L15
L26:
	;
	goto L14
L27:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+14)) = uint16(v28)
	v495 = v117
	v497 = v119
	goto L6
L28:
	;
	goto L29
L29:
	;
	v129 = v28 - v22
	if base.Ui32(v129) < base.Ui32(v121) {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	v132 = v20 + int32(48)
	v135 = base.I32_div_s(v116-v132, int32(6))
	v136 = int32(0)
	v145 = m.G0
	v147 = v145 + int32(-8192)
	m.G0 = v147
	if v120 != 0 {
		goto L33
	} else {
		goto L34
	}
L31:
	;
	v495 = v117
	v497 = v119
	goto L6
L32:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+14)) = uint16(v474)
	m.G0 = v147 - int32(-8192)
	goto L31
L33:
	;
	v149 = int32(1)
	if v135 <= v149 {
		goto L36
	} else {
		goto L37
	}
L34:
	;
	goto L35
L35:
	;
	v249 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v249) {
		goto L61
	} else {
		goto L62
	}
L36:
	;
	v152 = v149
	goto L38
L37:
	;
	v152 = v135
	goto L38
L38:
	;
	v153 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+16)))
	v157 = v153
	v160 = v136
	goto L40
L39:
	;
	v243 = v237 - v234
	if v243 == int32(0) {
		v474 = v233
		goto L32
	} else {
		goto L59
	}
L40:
	;
	v169 = v132 + v160*int32(6)
	v170 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v169)+4)))
	v171 = int32(*(*int16)(unsafe.Add(mBase, uint32(v169)+2)))
	v172 = v170 + v171
	if v157 == v172 {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	if v135 <= v160 {
		goto L46
	} else {
		goto L47
	}
L42:
	;
	v174 = v157 - v170
	v176 = v160 + int32(1)
	if v176 != v152 {
		v157 = v174
		v160 = v176
		goto L40
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	goto L41
L45:
	;
	v233 = v174
	v234 = v157
	v237 = v157
	goto L39
L46:
	;
	v233 = v157
	v234 = v172
	v237 = v172
	goto L39
L47:
	;
	goto L48
L48:
	;
	v184 = v157
	v185 = v172
	v187 = v160
	v188 = v172
	goto L49
L49:
	;
	v196 = v132 + v187*int32(6)
	v197 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v196))))
	v204 = l0 + int32(20) + (v197+int32(1))&int32(_a_F_PageRepairFragmentation_2)<<(uint(int32(2))%32)
	v205 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v196)+4)))
	v206 = int32(*(*int16)(unsafe.Add(mBase, uint32(v196)+2)))
	if v205+v206 == v185 {
		goto L52
	} else {
		goto L53
	}
L50:
	;
	v233 = v222
	v234 = v216
	v237 = v218
	goto L39
L51:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v204)))
	v222 = v184 - v217
	*(*int32)(unsafe.Add(mBase, uint32(v204))) = v219&int32(-32768) | v222&int32(_a_F_PageRepairFragmentation_6)
	v228 = v187 + int32(1)
	if v228 != v135 {
		v184 = v222
		v185 = v216
		v187 = v228
		v188 = v218
		goto L49
	} else {
		goto L58
	}
L52:
	;
	v216 = v206
	v217 = v205
	v218 = v188
	goto L51
L53:
	;
	goto L54
L54:
	;
	v209 = v188 - v185
	if v209 != 0 {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	base.MemoryCopy(m, l0+v184, l0+v185, v209)
	goto L57
L56:
	;
	goto L57
L57:
	;
	v213 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v196)+4)))
	v214 = int32(*(*int16)(unsafe.Add(mBase, uint32(v196)+2)))
	v216 = v214
	v217 = v213
	v218 = v213 + v214
	goto L51
L58:
	;
	goto L50
L59:
	;
	base.MemoryCopy(m, l0+v233, l0+v234, v243)
	v474 = v233
	goto L32
L60:
	;
	if v135 <= v393 {
		goto L95
	} else {
		goto L96
	}
L61:
	;
	v259 = int32(base.Ui32(v249+int32(_a_F_PageRepairFragmentation_3))>>(uint(int32(4))%32)) & int32(_a_F_PageRepairFragmentation_8)
	goto L63
L62:
	;
	v259 = int32(0)
	goto L63
L63:
	;
	if v135 < v259 {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	if int32(2) <= v135 {
		goto L68
	} else {
		goto L69
	}
L65:
	;
	goto L66
L66:
	;
	v349 = int32(1)
	if base.Ui32(v135) <= base.Ui32(v349) {
		goto L85
	} else {
		goto L86
	}
L67:
	;
	v344 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v132)+4)))
	v345 = int32(*(*int16)(unsafe.Add(mBase, uint32(v132)+2)))
	v347 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+16)))
	v390 = v347
	v392 = v344 + v345
	v393 = int32(0)
	goto L60
L68:
	;
	v264 = int32(1)
	if v135 <= v264 {
		goto L71
	} else {
		goto L72
	}
L69:
	;
	v311 = int32(0)
	goto L70
L70:
	;
	v323 = v132 + v311*int32(6)
	v324 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v323)+4)))
	if v324 == int32(0) {
		goto L67
	} else {
		goto L84
	}
L71:
	;
	v267 = v264
	goto L73
L72:
	;
	v267 = v135
	goto L73
L73:
	;
	v276 = int32(0)
	v282 = v136
	goto L74
L74:
	;
	v288 = v132 + v276*int32(6)
	v289 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v288)+4)))
	if v289 != 0 {
		goto L76
	} else {
		goto L77
	}
L75:
	;
	if v267&int32(1) == int32(0) {
		goto L67
	} else {
		goto L83
	}
L76:
	;
	v290 = int32(*(*int16)(unsafe.Add(mBase, uint32(v288)+2)))
	base.MemoryCopy(m, v147+v290, l0+v290, v289)
	goto L78
L77:
	;
	goto L78
L78:
	;
	v295 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v288)+10)))
	if v295 != 0 {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v296 = int32(*(*int16)(unsafe.Add(mBase, uint32(v288)+8)))
	base.MemoryCopy(m, v147+v296, l0+v296, v295)
	goto L81
L80:
	;
	goto L81
L81:
	;
	v301 = int32(2)
	v302 = v276 + v301
	v304 = v282 + v301
	if v304 != v267&int32(2147483646) {
		v276 = v302
		v282 = v304
		goto L74
	} else {
		goto L82
	}
L82:
	;
	goto L75
L83:
	;
	v311 = v302
	goto L70
L84:
	;
	v327 = int32(*(*int16)(unsafe.Add(mBase, uint32(v323)+2)))
	base.MemoryCopy(m, v147+v327, l0+v327, v324)
	goto L67
L85:
	;
	v352 = v349
	goto L87
L86:
	;
	v352 = v135
	goto L87
L87:
	;
	v353 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+16)))
	v357 = v353
	v360 = v136
	goto L89
L88:
	;
	v380 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+14)))
	v381 = v378 - v380
	if v381 == int32(0) {
		v390 = v378
		v392 = v372
		v393 = v379
		goto L60
	} else {
		goto L93
	}
L89:
	;
	v369 = v132 + v360*int32(6)
	v370 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v369)+4)))
	v371 = int32(*(*int16)(unsafe.Add(mBase, uint32(v369)+2)))
	v372 = v370 + v371
	if v357 != v372 {
		v378 = v357
		v379 = v360
		goto L88
	} else {
		goto L91
	}
L90:
	;
	v378 = v374
	v379 = v352
	goto L88
L91:
	;
	v374 = v357 - v370
	v376 = v360 + int32(1)
	if v376 != v352 {
		v357 = v374
		v360 = v376
		goto L89
	} else {
		goto L92
	}
L92:
	;
	goto L90
L93:
	;
	base.MemoryCopy(m, v380+v147, l0+v380, v381)
	v390 = v378
	v392 = v372
	v393 = v379
	goto L60
L94:
	;
	v465 = v457 - v456
	if v465 == int32(0) {
		v474 = v455
		goto L32
	} else {
		goto L108
	}
L95:
	;
	v455 = v390
	v456 = v392
	v457 = v392
	goto L94
L96:
	;
	goto L97
L97:
	;
	v406 = v390
	v407 = v392
	v408 = v392
	v409 = v393
	goto L98
L98:
	;
	v418 = v132 + v409*int32(6)
	v419 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v418))))
	v426 = l0 + int32(20) + (v419+int32(1))&int32(_a_F_PageRepairFragmentation_2)<<(uint(int32(2))%32)
	v427 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v418)+4)))
	v428 = int32(*(*int16)(unsafe.Add(mBase, uint32(v418)+2)))
	if v427+v428 == v407 {
		goto L101
	} else {
		goto L102
	}
L99:
	;
	v455 = v444
	v456 = v438
	v457 = v439
	goto L94
L100:
	;
	v441 = *(*int32)(unsafe.Add(mBase, uint32(v426)))
	v444 = v406 - v440
	*(*int32)(unsafe.Add(mBase, uint32(v426))) = v441&int32(-32768) | v444&int32(_a_F_PageRepairFragmentation_6)
	v450 = v409 + int32(1)
	if v450 != v135 {
		v406 = v444
		v407 = v438
		v408 = v439
		v409 = v450
		goto L98
	} else {
		goto L107
	}
L101:
	;
	v438 = v428
	v439 = v408
	v440 = v427
	goto L100
L102:
	;
	goto L103
L103:
	;
	v431 = v408 - v407
	if v431 != 0 {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	base.MemoryCopy(m, l0+v406, v407+v147, v431)
	goto L106
L105:
	;
	goto L106
L106:
	;
	v435 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v418)+4)))
	v436 = int32(*(*int16)(unsafe.Add(mBase, uint32(v418)+2)))
	v438 = v436
	v439 = v435 + v436
	v440 = v435
	goto L100
L107:
	;
	goto L99
L108:
	;
	base.MemoryCopy(m, l0+v455, v456+v147, v465)
	v474 = v455
	goto L32
L109:
	;
	v510 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
	v511 = v51 - v506
	v514 = v510 - v511<<(uint(int32(2))%32)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)) = uint16(v514)
	v518 = v495 - v511
	goto L111
L110:
	;
	v518 = v495
	goto L111
L111:
	;
	v519 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+10)))
	v524 = v519&int32(_a_F_PageRepairFragmentation_9) | base.B2i32(int32(0) < v518)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+10)) = uint16(v524)
	m.G0 = v20 + int32(1808)
	return
L112:
	;
	return
L113:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v535 = m.ExcPending
	if v535 != 0 {
		goto L112
	} else {
		goto L114
	}
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+8)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v20)+4)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v22
	F_errmsg(m, int32(_a_F_PageRepairFragmentation_10), v20)
	mBase = m.M
	v541 = m.ExcPending
	if v541 != 0 {
		goto L112
	} else {
		goto L115
	}
L115:
	;
	F_errfinish(m, int32(_a_F_PageRepairFragmentation_11), int32(730), int32(_a_F_PageRepairFragmentation_12))
	mBase = m.M
	v546 = m.ExcPending
	if v546 != 0 {
		goto L112
	} else {
		goto L116
	}
L116:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L117:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v553 = m.ExcPending
	if v553 != 0 {
		goto L112
	} else {
		goto L118
	}
L118:
	;
	v554 = int32(*(*int16)(unsafe.Add(mBase, uint32(v66)+2)))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+32)) = v554
	F_errmsg(m, int32(_a_F_PageRepairFragmentation_13), v20+int32(32))
	mBase = m.M
	v560 = m.ExcPending
	if v560 != 0 {
		goto L112
	} else {
		goto L119
	}
L119:
	;
	F_errfinish(m, int32(_a_F_PageRepairFragmentation_11), int32(759), int32(_a_F_PageRepairFragmentation_12))
	mBase = m.M
	v565 = m.ExcPending
	if v565 != 0 {
		goto L112
	} else {
		goto L120
	}
L120:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L121:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v572 = m.ExcPending
	if v572 != 0 {
		goto L112
	} else {
		goto L122
	}
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+20)) = v129
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = v121
	F_errmsg(m, int32(_a_F_PageRepairFragmentation_14), v20+int32(16))
	mBase = m.M
	v579 = m.ExcPending
	if v579 != 0 {
		goto L112
	} else {
		goto L123
	}
L123:
	;
	F_errfinish(m, int32(_a_F_PageRepairFragmentation_11), int32(789), int32(_a_F_PageRepairFragmentation_12))
	mBase = m.M
	v584 = m.ExcPending
	if v584 != 0 {
		goto L112
	} else {
		goto L124
	}
L124:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_page_checksum(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_page_checksum_internal(m, l0, int32(0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_page_checksum_1_9(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_page_checksum_internal(m, l0, int32(1))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_page_checksum_internal(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int64
	_ = v12
	var v14 int64
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v74 int32
	_ = v74
	var v108 int32
	_ = v108
	var v115 int32
	_ = v115
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v185 int32
	_ = v185
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v232 int32
	_ = v232
	var v239 int32
	_ = v239
	var v247 int32
	_ = v247
	var v256 int32
	_ = v256
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v296 int32
	_ = v296
	var v303 int32
	_ = v303
	var v310 int32
	_ = v310
	var v318 int32
	_ = v318
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v397 int32
	_ = v397
	var v400 int32
	_ = v400
	var v404 int32
	_ = v404
	var v409 int32
	_ = v409
	var v413 int32
	_ = v413
	var v416 int32
	_ = v416
	var v420 int32
	_ = v420
	var v425 int32
	_ = v425
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = F_pg_detoast_datum(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		if l1 != 0 {
			v12 = *(*int64)(unsafe.Add(mBase, uint32(v11)))
			v14 = v12
		} else {
			v14 = base.I64_extend_i32_u(v11)
		}
		v15 = F_superuser(m)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			if v15 != 0 {
				if base.Ui64(int64(4294967295)) <= base.Ui64(v14) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v413 = m.ExcPending
					if v413 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(50856066))
						mBase = m.M
						v416 = m.ExcPending
						if v416 != 0 {
							return int32(0)
						} else {
							F_errmsg(m, int32(_a_F_page_checksum_internal_0), int32(0))
							mBase = m.M
							v420 = m.ExcPending
							if v420 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_page_checksum_internal_1), int32(356), int32(_a_F_page_checksum_internal_2))
								mBase = m.M
								v425 = m.ExcPending
								if v425 != 0 {
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
					v19 = F_get_page_from_raw(m, v7)
					mBase = m.M
					v20 = m.ExcPending
					if v20 != 0 {
						return int32(0)
					} else {
						v21 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v19)+14)))
						if v21 == int32(0) {
							v24 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v24)
							return int32(0)
						} else {
							v29 = int32(0)
							v62 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v19)+8)))
							*(*uint16)(unsafe.Add(mBase, uint32(v19)+8)) = uint16(v29)
							v65 = m.G0
							v66 = int32(128)
							v67 = v65 - v66
							base.MemoryCopy(m, v67, int32(_a_F_page_checksum_internal_3), v66)
							v74 = v29
							for {
								v108 = v19 + v74<<(uint(int32(7))%32)
								v115 = int32(0)
								for {
									v145 = int32(2)
									v146 = v115 << (uint(v145) % 32)
									v147 = v67 + v146
									v149 = *(*int32)(unsafe.Add(mBase, uint32(v108+v146)))
									v150 = *(*int32)(unsafe.Add(mBase, uint32(v147)))
									v151 = v149 ^ v150
									v152 = int32(16777619)
									v154 = int32(17)
									*(*int32)(unsafe.Add(mBase, uint32(v147))) = v151*v152 ^ int32(base.Ui32(v151)>>(uint(v154)%32))
									v159 = v146 | int32(4)
									v160 = v67 + v159
									v162 = *(*int32)(unsafe.Add(mBase, uint32(v108+v159)))
									v163 = *(*int32)(unsafe.Add(mBase, uint32(v160)))
									v164 = v162 ^ v163
									*(*int32)(unsafe.Add(mBase, uint32(v160))) = v164*v152 ^ int32(base.Ui32(v164)>>(uint(v154)%32))
									v172 = v115 + v145
									if v172 != int32(32) {
										v115 = v172
										continue
									} else {
										break
									}
									break
								}
								v176 = v74 + int32(1)
								if v176 != int32(64) {
									v74 = v176
									continue
								} else {
									break
								}
								break
							}
							v185 = int32(0)
							for {
								v217 = v67 + v185<<(uint(int32(2))%32)
								v218 = *(*int32)(unsafe.Add(mBase, uint32(v217)))
								v219 = int32(16777619)
								v221 = int32(17)
								*(*int32)(unsafe.Add(mBase, uint32(v217))) = v218*v219 ^ int32(base.Ui32(v218)>>(uint(v221)%32))
								v225 = *(*int32)(unsafe.Add(mBase, uint32(v217)+4))
								*(*int32)(unsafe.Add(mBase, uint32(v217)+4)) = v225*v219 ^ int32(base.Ui32(v225)>>(uint(v221)%32))
								v232 = *(*int32)(unsafe.Add(mBase, uint32(v217)+8))
								*(*int32)(unsafe.Add(mBase, uint32(v217)+8)) = v232*v219 ^ int32(base.Ui32(v232)>>(uint(v221)%32))
								v239 = *(*int32)(unsafe.Add(mBase, uint32(v217)+12))
								*(*int32)(unsafe.Add(mBase, uint32(v217)+12)) = v239*v219 ^ int32(base.Ui32(v239)>>(uint(v221)%32))
								v247 = v185 + int32(4)
								if v247 != int32(32) {
									v185 = v247
									continue
								} else {
									break
								}
								break
							}
							v256 = int32(0)
							for {
								v288 = v67 + v256<<(uint(int32(2))%32)
								v289 = *(*int32)(unsafe.Add(mBase, uint32(v288)))
								v290 = int32(16777619)
								v292 = int32(17)
								*(*int32)(unsafe.Add(mBase, uint32(v288))) = v289*v290 ^ int32(base.Ui32(v289)>>(uint(v292)%32))
								v296 = *(*int32)(unsafe.Add(mBase, uint32(v288)+4))
								*(*int32)(unsafe.Add(mBase, uint32(v288)+4)) = v296*v290 ^ int32(base.Ui32(v296)>>(uint(v292)%32))
								v303 = *(*int32)(unsafe.Add(mBase, uint32(v288)+8))
								*(*int32)(unsafe.Add(mBase, uint32(v288)+8)) = v303*v290 ^ int32(base.Ui32(v303)>>(uint(v292)%32))
								v310 = *(*int32)(unsafe.Add(mBase, uint32(v288)+12))
								*(*int32)(unsafe.Add(mBase, uint32(v288)+12)) = v310*v290 ^ int32(base.Ui32(v310)>>(uint(v292)%32))
								v318 = v256 + int32(4)
								if v318 != int32(32) {
									v256 = v318
									continue
								} else {
									break
								}
								break
							}
							v321 = *(*int32)(unsafe.Add(mBase, uint32(v67)+124))
							v322 = *(*int32)(unsafe.Add(mBase, uint32(v67)+120))
							v323 = *(*int32)(unsafe.Add(mBase, uint32(v67)+116))
							v324 = *(*int32)(unsafe.Add(mBase, uint32(v67)+112))
							v325 = *(*int32)(unsafe.Add(mBase, uint32(v67)+108))
							v326 = *(*int32)(unsafe.Add(mBase, uint32(v67)+104))
							v327 = *(*int32)(unsafe.Add(mBase, uint32(v67)+100))
							v328 = *(*int32)(unsafe.Add(mBase, uint32(v67)+96))
							v329 = *(*int32)(unsafe.Add(mBase, uint32(v67)+92))
							v330 = *(*int32)(unsafe.Add(mBase, uint32(v67)+88))
							v331 = *(*int32)(unsafe.Add(mBase, uint32(v67)+84))
							v332 = *(*int32)(unsafe.Add(mBase, uint32(v67)+80))
							v333 = *(*int32)(unsafe.Add(mBase, uint32(v67)+76))
							v334 = *(*int32)(unsafe.Add(mBase, uint32(v67)+72))
							v335 = *(*int32)(unsafe.Add(mBase, uint32(v67)+68))
							v336 = *(*int32)(unsafe.Add(mBase, uint32(v67)+64))
							v337 = *(*int32)(unsafe.Add(mBase, uint32(v67)+60))
							v338 = *(*int32)(unsafe.Add(mBase, uint32(v67)+56))
							v339 = *(*int32)(unsafe.Add(mBase, uint32(v67)+52))
							v340 = *(*int32)(unsafe.Add(mBase, uint32(v67)+48))
							v341 = *(*int32)(unsafe.Add(mBase, uint32(v67)+44))
							v342 = *(*int32)(unsafe.Add(mBase, uint32(v67)+40))
							v343 = *(*int32)(unsafe.Add(mBase, uint32(v67)+36))
							v344 = *(*int32)(unsafe.Add(mBase, uint32(v67)+32))
							v345 = *(*int32)(unsafe.Add(mBase, uint32(v67)+28))
							v346 = *(*int32)(unsafe.Add(mBase, uint32(v67)+24))
							v347 = *(*int32)(unsafe.Add(mBase, uint32(v67)+20))
							v348 = *(*int32)(unsafe.Add(mBase, uint32(v67)+16))
							v349 = *(*int32)(unsafe.Add(mBase, uint32(v67)+12))
							v350 = *(*int32)(unsafe.Add(mBase, uint32(v67)+8))
							v351 = *(*int32)(unsafe.Add(mBase, uint32(v67)+4))
							v352 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
							*(*uint16)(unsafe.Add(mBase, uint32(v19)+8)) = uint16(v62)
							v386 = int32(_a_F_page_checksum_internal_4)
							v387 = base.I32_rem_u_s(v321^(v322^(v323^(v324^(v325^(v326^(v327^(v328^(v329^(v330^(v331^(v332^(v333^(v334^(v335^(v336^(v337^(v338^(v339^(v340^(v341^(v342^(v343^(v344^(v345^(v346^(v347^(v348^(v349^(v350^(v351^(base.I32_wrap_i64(v14)^v352))))))))))))))))))))))))))))))), v386)
							return base.I32_extend16_s((v387 + int32(1)) & v386)
						}
					}
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v397 = m.ExcPending
				if v397 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(16797828))
					mBase = m.M
					v400 = m.ExcPending
					if v400 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(_a_F_page_checksum_internal_5), int32(0))
						mBase = m.M
						v404 = m.ExcPending
						if v404 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_page_checksum_internal_1), int32(351), int32(_a_F_page_checksum_internal_2))
							mBase = m.M
							v409 = m.ExcPending
							if v409 != 0 {
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
func F_page_header(m *base.Module, l0 int32) int32 {
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
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	v6 = m.G0
	v8 = v6 - int32(96)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = F_pg_detoast_datum(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		v15 = F_superuser(m)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			if v15 != 0 {
				v17 = F_get_page_from_raw(m, v11)
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return int32(0)
				} else {
					v22 = F_get_call_result_type(m, l0, int32(0), v8+int32(92))
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
						return int32(0)
					} else {
						if v22 != int32(1) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v132 = m.ExcPending
							if v132 != 0 {
								return int32(0)
							} else {
								F_errmsg_internal(m, int32(_a_F_page_header_0), int32(0))
								mBase = m.M
								v136 = m.ExcPending
								if v136 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_page_header_1), int32(274), int32(_a_F_page_header_2))
									mBase = m.M
									v141 = m.ExcPending
									if v141 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						} else {
							v26 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
							v27 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
							v28 = *(*int32)(unsafe.Add(mBase, uint32(v8)+92))
							v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
							v33 = *(*int32)(unsafe.Add(mBase, uint32(v28+v29<<(uint(int32(4))%32))+88))
							if v33 == int32(25) {
								*(*int32)(unsafe.Add(mBase, uint32(v8))) = v27
								*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v26
								v39 = v8 + int32(16)
								v42 = F_pg_snprintf(m, v39, int32(64), int32(_a_F_page_header_3), v8)
								mBase = m.M
								v43 = m.ExcPending
								if v43 != 0 {
									return int32(0)
								} else {
									v44 = F_cstring_to_text(m, v39)
									mBase = m.M
									v45 = m.ExcPending
									if v45 != 0 {
										return int32(0)
									} else {
										v54 = v44
										*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v54
										v56 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17)+8)))
										*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = v56
										v58 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17)+10)))
										*(*int32)(unsafe.Add(mBase, uint32(v8)+24)) = v58
										v60 = *(*int32)(unsafe.Add(mBase, uint32(v8)+92))
										v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
										v65 = *(*int32)(unsafe.Add(mBase, uint32(v60+v61<<(uint(int32(4))%32))+388))
										switch v65 - int32(21) {
										case 0, 2:
											v81 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17)+12)))
											*(*int32)(unsafe.Add(mBase, uint32(v8)+28)) = v81
											v83 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17)+14)))
											*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v83
											v85 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17)+16)))
											*(*int32)(unsafe.Add(mBase, uint32(v8)+36)) = v85
											v87 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17)+18)))
											*(*int32)(unsafe.Add(mBase, uint32(v8)+44)) = v87 & int32(255)
											*(*int32)(unsafe.Add(mBase, uint32(v8)+40)) = v87 & int32(_a_F_page_header_4)
											v94 = *(*int32)(unsafe.Add(mBase, uint32(v17)+20))
											*(*int64)(unsafe.Add(mBase, uint32(v8)+80)) = int64(0)
											*(*int32)(unsafe.Add(mBase, uint32(v8)+48)) = v94
											v98 = int32(0)
											*(*uint8)(unsafe.Add(mBase, uint32(v8)+88)) = uint8(v98)
											v104 = F_heap_form_tuple(m, v60, v8+int32(16), v8+int32(80))
											mBase = m.M
											v105 = m.ExcPending
											if v105 != 0 {
												return int32(0)
											} else {
												v106 = *(*int32)(unsafe.Add(mBase, uint32(v104)+16))
												v107 = F_HeapTupleHeaderGetDatum(m, v106)
												mBase = m.M
												v108 = m.ExcPending
												if v108 != 0 {
													return int32(0)
												} else {
													m.G0 = v8 + int32(96)
													return v107
												}
											}
										default:
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v71 = m.ExcPending
											if v71 != 0 {
												return int32(0)
											} else {
												F_errmsg_internal(m, int32(_a_F_page_header_5), int32(0))
												mBase = m.M
												v75 = m.ExcPending
												if v75 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(_a_F_page_header_1), int32(315), int32(_a_F_page_header_2))
													mBase = m.M
													v80 = m.ExcPending
													if v80 != 0 {
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
							} else {
								v51 = F_Int64GetDatum(m, base.I64_extend_i32_u(v26)|base.I64_extend_i32_u(v27)<<(uint(int64(32))%64))
								mBase = m.M
								v52 = m.ExcPending
								if v52 != 0 {
									return int32(0)
								} else {
									v54 = v51
									*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v54
									v56 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17)+8)))
									*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = v56
									v58 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17)+10)))
									*(*int32)(unsafe.Add(mBase, uint32(v8)+24)) = v58
									v60 = *(*int32)(unsafe.Add(mBase, uint32(v8)+92))
									v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
									v65 = *(*int32)(unsafe.Add(mBase, uint32(v60+v61<<(uint(int32(4))%32))+388))
									switch v65 - int32(21) {
									case 0, 2:
										v81 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17)+12)))
										*(*int32)(unsafe.Add(mBase, uint32(v8)+28)) = v81
										v83 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17)+14)))
										*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v83
										v85 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17)+16)))
										*(*int32)(unsafe.Add(mBase, uint32(v8)+36)) = v85
										v87 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17)+18)))
										*(*int32)(unsafe.Add(mBase, uint32(v8)+44)) = v87 & int32(255)
										*(*int32)(unsafe.Add(mBase, uint32(v8)+40)) = v87 & int32(_a_F_page_header_4)
										v94 = *(*int32)(unsafe.Add(mBase, uint32(v17)+20))
										*(*int64)(unsafe.Add(mBase, uint32(v8)+80)) = int64(0)
										*(*int32)(unsafe.Add(mBase, uint32(v8)+48)) = v94
										v98 = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(v8)+88)) = uint8(v98)
										v104 = F_heap_form_tuple(m, v60, v8+int32(16), v8+int32(80))
										mBase = m.M
										v105 = m.ExcPending
										if v105 != 0 {
											return int32(0)
										} else {
											v106 = *(*int32)(unsafe.Add(mBase, uint32(v104)+16))
											v107 = F_HeapTupleHeaderGetDatum(m, v106)
											mBase = m.M
											v108 = m.ExcPending
											if v108 != 0 {
												return int32(0)
											} else {
												m.G0 = v8 + int32(96)
												return v107
											}
										}
									default:
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v71 = m.ExcPending
										if v71 != 0 {
											return int32(0)
										} else {
											F_errmsg_internal(m, int32(_a_F_page_header_5), int32(0))
											mBase = m.M
											v75 = m.ExcPending
											if v75 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(_a_F_page_header_1), int32(315), int32(_a_F_page_header_2))
												mBase = m.M
												v80 = m.ExcPending
												if v80 != 0 {
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
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v116 = m.ExcPending
				if v116 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(16797828))
					mBase = m.M
					v119 = m.ExcPending
					if v119 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(_a_F_page_header_6), int32(0))
						mBase = m.M
						v123 = m.ExcPending
						if v123 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_page_header_1), int32(267), int32(_a_F_page_header_2))
							mBase = m.M
							v128 = m.ExcPending
							if v128 != 0 {
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
