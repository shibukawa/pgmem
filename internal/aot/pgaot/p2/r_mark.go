package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_r_mark_lAr(m *base.Module, l0 int32) int32 {
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v5 = Fn13980(m, l0, int32(2), int32(_a_F_r_mark_lAr_0), int32(114))
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
func F_r_mark_possessives(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v141 int32
	_ = v141
	var v157 int32
	_ = v157
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v260 int32
	_ = v260
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v272 int32
	_ = v272
	var v289 int32
	_ = v289
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v348 int32
	_ = v348
	var v350 int32
	_ = v350
	var v352 int32
	_ = v352
	var v354 int32
	_ = v354
	var v367 int32
	_ = v367
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v389 int32
	_ = v389
	var v391 int32
	_ = v391
	var v399 int32
	_ = v399
	var v403 int32
	_ = v403
	var v405 int32
	_ = v405
	var v411 int32
	_ = v411
	var v427 int32
	_ = v427
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v448 int32
	_ = v448
	var v450 int32
	_ = v450
	var v455 int32
	_ = v455
	var v457 int32
	_ = v457
	var v463 int32
	_ = v463
	var v468 int32
	_ = v468
	var v472 int32
	_ = v472
	var v475 int32
	_ = v475
	var v479 int32
	_ = v479
	var v493 int32
	_ = v493
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v530 int32
	_ = v530
	var v532 int32
	_ = v532
	var v539 int32
	_ = v539
	var v541 int32
	_ = v541
	var v543 int32
	_ = v543
	var v545 int32
	_ = v545
	var v558 int32
	_ = v558
	var v560 int32
	_ = v560
	var v562 int32
	_ = v562
	var v580 int32
	_ = v580
	var v582 int32
	_ = v582
	var v590 int32
	_ = v590
	var v594 int32
	_ = v594
	var v596 int32
	_ = v596
	var v602 int32
	_ = v602
	var v619 int32
	_ = v619
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v633 int32
	_ = v633
	var v640 int32
	_ = v640
	v2 = int32(0)
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v7 <= v8 {
		v640 = v2
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v640
L2:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v12 = int32(1)
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10+v7-v12))))
	if base.B2i32(v14&int32(224) != int32(96))|base.B2i32(v12<<(uint(v14)%32)&int32(67133440) == int32(0)) != 0 {
		v640 = v2
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v28 = F_find_among_b(m, l0, int32(_a_F_r_mark_possessives_0), int32(10))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return int32(0)
L5:
	;
	if v28 == int32(0) {
		v640 = v2
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L11
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v633
	v640 = int32(1)
	goto L1
L8:
	;
	v302 = v35 - v34
	v303 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v304 = v302 + v303
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v304
	v319 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v320 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L54
L9:
	;
	if v164 != 0 {
		goto L8
	} else {
		goto L32
	}
L10:
	;
	v164 = v157
	goto L9
L11:
	;
	if v35 <= v49 {
		v157 = int32(-1)
		goto L10
	} else {
		goto L13
	}
L12:
	;
	v157 = int32(0)
	goto L10
L13:
	;
	v66 = int32(1)
	v67 = v35 - v66
	v69 = int32(*(*int8)(unsafe.Add(mBase, uint32(v50+v67))))
	v71 = v69 & int32(255)
	if base.B2i32(v67 == v49)|base.B2i32(int32(0) <= v69) != 0 {
		v129 = v71
		v133 = v66
		goto L14
	} else {
		goto L15
	}
L14:
	;
	if int32(305) < v129 {
		goto L22
	} else {
		goto L23
	}
L15:
	;
	v78 = v71 & int32(63)
	v80 = v35 - int32(2)
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50+v80))))
	v84 = v82 << (uint(int32(6)) % 32)
	if base.B2i32(v80 != v49)&base.B2i32(base.Ui32(v82) < base.Ui32(int32(192))) == int32(0) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v129 = v84&int32(1984) | v78
	v133 = int32(2)
	goto L14
L17:
	;
	goto L18
L18:
	;
	v97 = v84&int32(4032) | v78
	v99 = v35 - int32(3)
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50+v99))))
	if base.B2i32(v99 != v49)&base.B2i32(base.Ui32(v101) < base.Ui32(int32(224))) == int32(0) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v129 = v101<<(uint(int32(12))%32)&int32(_a_F_r_mark_possessives_1) | v97
	v133 = int32(3)
	goto L14
L20:
	;
	goto L21
L21:
	;
	v119 = int32(4)
	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35+v50-v119))))
	v129 = v101<<(uint(int32(12))%32)&int32(_a_F_r_mark_possessives_2) | v121&int32(7)<<(uint(int32(18))%32) | v97
	v133 = v119
	goto L14
L22:
	;
	v164 = v133
	goto L9
L23:
	;
	goto L24
L24:
	;
	v135 = v129 - int32(105)
	if v135 < int32(0) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v164 = v133
	goto L9
L26:
	;
	goto L27
L27:
	;
	v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v135)>>(uint(int32(3))%32)))+uint32(_c_F_r_mark_possessives[0]))))
	if int32(base.Ui32(v141)>>(uint(v135&int32(7))%32))&int32(1) == int32(0) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v164 = v133
	goto L9
L29:
	;
	goto L30
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v35 - v133
	goto L31
L31:
	;
	goto L12
L32:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v166 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v180 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v181 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L35
L33:
	;
	if v296 != 0 {
		goto L8
	} else {
		goto L51
	}
L34:
	;
	v296 = v289
	goto L33
L35:
	;
	if v165 <= v180 {
		v289 = int32(-1)
		goto L34
	} else {
		goto L37
	}
L36:
	;
	v289 = int32(0)
	goto L34
L37:
	;
	v197 = int32(1)
	v198 = v165 - v197
	v200 = int32(*(*int8)(unsafe.Add(mBase, uint32(v181+v198))))
	v202 = v200 & int32(255)
	if base.B2i32(v198 == v180)|base.B2i32(int32(0) <= v200) != 0 {
		v260 = v202
		v264 = v197
		goto L38
	} else {
		goto L39
	}
L38:
	;
	if int32(305) < v260 {
		goto L46
	} else {
		goto L47
	}
L39:
	;
	v209 = v202 & int32(63)
	v211 = v165 - int32(2)
	v213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v181+v211))))
	v215 = v213 << (uint(int32(6)) % 32)
	if base.B2i32(v211 != v180)&base.B2i32(base.Ui32(v213) < base.Ui32(int32(192))) == int32(0) {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v260 = v215&int32(1984) | v209
	v264 = int32(2)
	goto L38
L41:
	;
	goto L42
L42:
	;
	v228 = v215&int32(4032) | v209
	v230 = v165 - int32(3)
	v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v181+v230))))
	if base.B2i32(v230 != v180)&base.B2i32(base.Ui32(v232) < base.Ui32(int32(224))) == int32(0) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v260 = v232<<(uint(int32(12))%32)&int32(_a_F_r_mark_possessives_1) | v228
	v264 = int32(3)
	goto L38
L44:
	;
	goto L45
L45:
	;
	v250 = int32(4)
	v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165+v181-v250))))
	v260 = v232<<(uint(int32(12))%32)&int32(_a_F_r_mark_possessives_2) | v252&int32(7)<<(uint(int32(18))%32) | v228
	v264 = v250
	goto L38
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v165 - v264
	goto L50
L47:
	;
	v266 = v260 - int32(97)
	if v266 < int32(0) {
		goto L46
	} else {
		goto L48
	}
L48:
	;
	v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v266)>>(uint(int32(3))%32)))+uint32(_c_F_r_mark_possessives[1]))))
	if int32(base.Ui32(v272)>>(uint(v266&int32(7))%32))&int32(1) == int32(0) {
		goto L46
	} else {
		goto L49
	}
L49:
	;
	v296 = v264
	goto L33
L50:
	;
	goto L36
L51:
	;
	v297 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v633 = v297 + (v165 - v166)
	goto L7
L52:
	;
	v435 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v436 = v435 + v302
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v436
	if v434 == int32(0) {
		v640 = v2
		goto L1
	} else {
		goto L75
	}
L53:
	;
	v434 = v427
	goto L52
L54:
	;
	if v304 <= v319 {
		v427 = int32(-1)
		goto L53
	} else {
		goto L56
	}
L55:
	;
	v427 = int32(0)
	goto L53
L56:
	;
	v336 = int32(1)
	v337 = v304 - v336
	v339 = int32(*(*int8)(unsafe.Add(mBase, uint32(v320+v337))))
	v341 = v339 & int32(255)
	if base.B2i32(v337 == v319)|base.B2i32(int32(0) <= v339) != 0 {
		v399 = v341
		v403 = v336
		goto L57
	} else {
		goto L58
	}
L57:
	;
	if int32(305) < v399 {
		goto L65
	} else {
		goto L66
	}
L58:
	;
	v348 = v341 & int32(63)
	v350 = v304 - int32(2)
	v352 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v320+v350))))
	v354 = v352 << (uint(int32(6)) % 32)
	if base.B2i32(v350 != v319)&base.B2i32(base.Ui32(v352) < base.Ui32(int32(192))) == int32(0) {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v399 = v354&int32(1984) | v348
	v403 = int32(2)
	goto L57
L60:
	;
	goto L61
L61:
	;
	v367 = v354&int32(4032) | v348
	v369 = v304 - int32(3)
	v371 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v320+v369))))
	if base.B2i32(v369 != v319)&base.B2i32(base.Ui32(v371) < base.Ui32(int32(224))) == int32(0) {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v399 = v371<<(uint(int32(12))%32)&int32(_a_F_r_mark_possessives_1) | v367
	v403 = int32(3)
	goto L57
L63:
	;
	goto L64
L64:
	;
	v389 = int32(4)
	v391 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v304+v320-v389))))
	v399 = v371<<(uint(int32(12))%32)&int32(_a_F_r_mark_possessives_2) | v391&int32(7)<<(uint(int32(18))%32) | v367
	v403 = v389
	goto L57
L65:
	;
	v434 = v403
	goto L52
L66:
	;
	goto L67
L67:
	;
	v405 = v399 - int32(105)
	if v405 < int32(0) {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v434 = v403
	goto L52
L69:
	;
	goto L70
L70:
	;
	v411 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v405)>>(uint(int32(3))%32)))+uint32(_c_F_r_mark_possessives[0]))))
	if int32(base.Ui32(v411)>>(uint(v405&int32(7))%32))&int32(1) == int32(0) {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v434 = v403
	goto L52
L72:
	;
	goto L73
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v304 - v403
	goto L74
L74:
	;
	goto L55
L75:
	;
	v440 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v441 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L78
L76:
	;
	if v493 < int32(0) {
		v640 = v2
		goto L1
	} else {
		goto L95
	}
L78:
	;
	goto L79
L79:
	;
	goto L80
L80:
	;
	v448 = v436
	v450 = int32(1)
	goto L83
L82:
	;
	v493 = v475
	goto L76
L83:
	;
	if v448 <= v441 {
		goto L85
	} else {
		goto L86
	}
L84:
	;
	goto L82
L85:
	;
	v493 = int32(-1)
	goto L76
L86:
	;
	goto L87
L87:
	;
	v455 = v448 - int32(1)
	v457 = int32(*(*int8)(unsafe.Add(mBase, uint32(v440+v455))))
	if base.B2i32(int32(0) <= v457)|base.B2i32(v455 <= v441) != 0 {
		v475 = v455
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v479 = int32(1)
	if v479 < v450 {
		v448 = v475
		v450 = v450 - v479
		goto L83
	} else {
		goto L94
	}
L89:
	;
	v463 = v455
	goto L90
L90:
	;
	v468 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v440+v463))))
	if base.Ui32(int32(191)) < base.Ui32(v468) {
		v475 = v463
		goto L88
	} else {
		goto L92
	}
L91:
	;
	v475 = v441
	goto L88
L92:
	;
	v472 = v463 - int32(1)
	if v441 < v472 {
		v463 = v472
		goto L90
	} else {
		goto L93
	}
L93:
	;
	goto L91
L94:
	;
	goto L84
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v493
	v510 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v511 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L98
L96:
	;
	if v626 != 0 {
		v640 = v2
		goto L1
	} else {
		goto L114
	}
L97:
	;
	v626 = v619
	goto L96
L98:
	;
	if v493 <= v510 {
		v619 = int32(-1)
		goto L97
	} else {
		goto L100
	}
L99:
	;
	v619 = int32(0)
	goto L97
L100:
	;
	v527 = int32(1)
	v528 = v493 - v527
	v530 = int32(*(*int8)(unsafe.Add(mBase, uint32(v511+v528))))
	v532 = v530 & int32(255)
	if base.B2i32(v528 == v510)|base.B2i32(int32(0) <= v530) != 0 {
		v590 = v532
		v594 = v527
		goto L101
	} else {
		goto L102
	}
L101:
	;
	if int32(305) < v590 {
		goto L109
	} else {
		goto L110
	}
L102:
	;
	v539 = v532 & int32(63)
	v541 = v493 - int32(2)
	v543 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v511+v541))))
	v545 = v543 << (uint(int32(6)) % 32)
	if base.B2i32(v541 != v510)&base.B2i32(base.Ui32(v543) < base.Ui32(int32(192))) == int32(0) {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	v590 = v545&int32(1984) | v539
	v594 = int32(2)
	goto L101
L104:
	;
	goto L105
L105:
	;
	v558 = v545&int32(4032) | v539
	v560 = v493 - int32(3)
	v562 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v511+v560))))
	if base.B2i32(v560 != v510)&base.B2i32(base.Ui32(v562) < base.Ui32(int32(224))) == int32(0) {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	v590 = v562<<(uint(int32(12))%32)&int32(_a_F_r_mark_possessives_1) | v558
	v594 = int32(3)
	goto L101
L107:
	;
	goto L108
L108:
	;
	v580 = int32(4)
	v582 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v493+v511-v580))))
	v590 = v562<<(uint(int32(12))%32)&int32(_a_F_r_mark_possessives_2) | v582&int32(7)<<(uint(int32(18))%32) | v558
	v594 = v580
	goto L101
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v493 - v594
	goto L113
L110:
	;
	v596 = v590 - int32(97)
	if v596 < int32(0) {
		goto L109
	} else {
		goto L111
	}
L111:
	;
	v602 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v596)>>(uint(int32(3))%32)))+uint32(_c_F_r_mark_possessives[1]))))
	if int32(base.Ui32(v602)>>(uint(v596&int32(7))%32))&int32(1) == int32(0) {
		goto L109
	} else {
		goto L112
	}
L112:
	;
	v626 = v594
	goto L96
L113:
	;
	goto L99
L114:
	;
	v627 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v633 = v627 + v302
	goto L7
}
