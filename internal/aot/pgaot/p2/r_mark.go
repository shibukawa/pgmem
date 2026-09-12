package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_r_mark_lAr(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	v2 = int32(0)
	v4 = F_r_check_vowel_harmony(m, l0)
	mBase = m.M
	if v4 == v2 {
		v27 = v2
		return v27
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		if v8-int32(2) <= v7 {
			v27 = v2
			return v27
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12+v8-int32(1)))))
			if v16 != int32(114) {
				v27 = v2
				return v27
			} else {
				v21 = F_find_among_b(m, l0, int32(4383424), int32(2))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					v27 = base.B2i32(v21 != int32(0))
					return v27
				}
			}
		}
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
	var v14 int32
	_ = v14
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
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
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v121 int32
	_ = v121
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v140 int32
	_ = v140
	var v156 int32
	_ = v156
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v252 int32
	_ = v252
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v265 int32
	_ = v265
	var v271 int32
	_ = v271
	var v288 int32
	_ = v288
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v318 int32
	_ = v318
	var v322 int32
	_ = v322
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v342 int32
	_ = v342
	var v344 int32
	_ = v344
	var v350 int32
	_ = v350
	var v352 int32
	_ = v352
	var v354 int32
	_ = v354
	var v356 int32
	_ = v356
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v373 int32
	_ = v373
	var v391 int32
	_ = v391
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v404 int32
	_ = v404
	var v410 int32
	_ = v410
	var v426 int32
	_ = v426
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v447 int32
	_ = v447
	var v449 int32
	_ = v449
	var v454 int32
	_ = v454
	var v456 int32
	_ = v456
	var v461 int32
	_ = v461
	var v466 int32
	_ = v466
	var v470 int32
	_ = v470
	var v473 int32
	_ = v473
	var v477 int32
	_ = v477
	var v491 int32
	_ = v491
	var v508 int32
	_ = v508
	var v512 int32
	_ = v512
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v532 int32
	_ = v532
	var v534 int32
	_ = v534
	var v540 int32
	_ = v540
	var v542 int32
	_ = v542
	var v544 int32
	_ = v544
	var v546 int32
	_ = v546
	var v559 int32
	_ = v559
	var v561 int32
	_ = v561
	var v563 int32
	_ = v563
	var v581 int32
	_ = v581
	var v589 int32
	_ = v589
	var v590 int32
	_ = v590
	var v594 int32
	_ = v594
	var v600 int32
	_ = v600
	var v617 int32
	_ = v617
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v631 int32
	_ = v631
	var v635 int32
	_ = v635
	v2 = int32(0)
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v7 <= v8 {
		v635 = v2
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v635
L2:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10+v7-int32(1)))))
	if v14&int32(224) != int32(96) {
		v635 = v2
		goto L1
	} else {
		goto L3
	}
L3:
	;
	if int32(1)<<(uint(v14)%32)&int32(67133440) == int32(0) {
		v635 = v2
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v27 = F_find_among_b(m, l0, int32(4384048), int32(10))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	return int32(0)
L6:
	;
	if v27 == int32(0) {
		v635 = v2
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L12
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v631
	v635 = int32(1)
	goto L1
L9:
	;
	v301 = v34 - v33
	v302 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v303 = v301 + v302
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v303
	v318 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v322 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L57
L10:
	;
	if v163 != 0 {
		goto L9
	} else {
		goto L34
	}
L11:
	;
	v163 = v156
	goto L10
L12:
	;
	if v34 <= v52 {
		v156 = int32(-1)
		goto L11
	} else {
		goto L14
	}
L13:
	;
	v156 = int32(0)
	goto L11
L14:
	;
	v69 = int32(1)
	v70 = v34 - v69
	v72 = int32(*(*int8)(unsafe.Add(mBase, uint32(v48+v70))))
	v74 = v72 & int32(255)
	if v70 == v52 {
		v129 = v74
		v130 = v69
		goto L15
	} else {
		goto L16
	}
L15:
	;
	if int32(305) < v129 {
		goto L24
	} else {
		goto L25
	}
L16:
	;
	if int32(0) <= v72 {
		v129 = v74
		v130 = v69
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v80 = v74 & int32(63)
	v82 = v34 - int32(2)
	v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48+v82))))
	v86 = v84 << (uint(int32(6)) % 32)
	if base.B2i32(v82 != v52)&base.B2i32(base.Ui32(v84) < base.Ui32(int32(192))) == int32(0) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v129 = v86&int32(1984) | v80
	v130 = int32(2)
	goto L15
L19:
	;
	goto L20
L20:
	;
	v99 = v86&int32(4032) | v80
	v101 = v34 - int32(3)
	v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48+v101))))
	if base.B2i32(v101 != v52)&base.B2i32(base.Ui32(v103) < base.Ui32(int32(224))) == int32(0) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v129 = v103<<(uint(int32(12))%32)&int32(61440) | v99
	v130 = int32(3)
	goto L15
L22:
	;
	goto L23
L23:
	;
	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34+(v48-int32(4))))))
	v129 = v103<<(uint(int32(12))%32)&int32(258048) | v121&int32(7)<<(uint(int32(18))%32) | v99
	v130 = int32(4)
	goto L15
L24:
	;
	v163 = v130
	goto L10
L25:
	;
	goto L26
L26:
	;
	v134 = v129 - int32(105)
	if v134 < int32(0) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v163 = v130
	goto L10
L28:
	;
	goto L29
L29:
	;
	v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v134)>>(uint(int32(3))%32)))+uint32(_consts[1340]))))
	if int32(base.Ui32(v140)>>(uint(v134&int32(7))%32))&int32(1) == int32(0) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v163 = v130
	goto L10
L31:
	;
	goto L32
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v34 - v130
	goto L33
L33:
	;
	goto L13
L34:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v165 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v179 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v183 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L37
L35:
	;
	if v295 != 0 {
		goto L9
	} else {
		goto L54
	}
L36:
	;
	v295 = v288
	goto L35
L37:
	;
	if v164 <= v183 {
		v288 = int32(-1)
		goto L36
	} else {
		goto L39
	}
L38:
	;
	v288 = int32(0)
	goto L36
L39:
	;
	v200 = int32(1)
	v201 = v164 - v200
	v203 = int32(*(*int8)(unsafe.Add(mBase, uint32(v179+v201))))
	v205 = v203 & int32(255)
	if v201 == v183 {
		v260 = v205
		v261 = v200
		goto L40
	} else {
		goto L41
	}
L40:
	;
	if int32(305) < v260 {
		goto L49
	} else {
		goto L50
	}
L41:
	;
	if int32(0) <= v203 {
		v260 = v205
		v261 = v200
		goto L40
	} else {
		goto L42
	}
L42:
	;
	v211 = v205 & int32(63)
	v213 = v164 - int32(2)
	v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v179+v213))))
	v217 = v215 << (uint(int32(6)) % 32)
	if base.B2i32(v213 != v183)&base.B2i32(base.Ui32(v215) < base.Ui32(int32(192))) == int32(0) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v260 = v217&int32(1984) | v211
	v261 = int32(2)
	goto L40
L44:
	;
	goto L45
L45:
	;
	v230 = v217&int32(4032) | v211
	v232 = v164 - int32(3)
	v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v179+v232))))
	if base.B2i32(v232 != v183)&base.B2i32(base.Ui32(v234) < base.Ui32(int32(224))) == int32(0) {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v260 = v234<<(uint(int32(12))%32)&int32(61440) | v230
	v261 = int32(3)
	goto L40
L47:
	;
	goto L48
L48:
	;
	v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v164+(v179-int32(4))))))
	v260 = v234<<(uint(int32(12))%32)&int32(258048) | v252&int32(7)<<(uint(int32(18))%32) | v230
	v261 = int32(4)
	goto L40
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v164 - v261
	goto L53
L50:
	;
	v265 = v260 - int32(97)
	if v265 < int32(0) {
		goto L49
	} else {
		goto L51
	}
L51:
	;
	v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v265)>>(uint(int32(3))%32)))+uint32(_consts[1333]))))
	if int32(base.Ui32(v271)>>(uint(v265&int32(7))%32))&int32(1) == int32(0) {
		goto L49
	} else {
		goto L52
	}
L52:
	;
	v295 = v261
	goto L35
L53:
	;
	goto L38
L54:
	;
	v296 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v631 = v296 + (v164 - v165)
	goto L8
L55:
	;
	v434 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v435 = v434 + v301
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v435
	if v433 == int32(0) {
		v635 = v2
		goto L1
	} else {
		goto L79
	}
L56:
	;
	v433 = v426
	goto L55
L57:
	;
	if v303 <= v322 {
		v426 = int32(-1)
		goto L56
	} else {
		goto L59
	}
L58:
	;
	v426 = int32(0)
	goto L56
L59:
	;
	v339 = int32(1)
	v340 = v303 - v339
	v342 = int32(*(*int8)(unsafe.Add(mBase, uint32(v318+v340))))
	v344 = v342 & int32(255)
	if v340 == v322 {
		v399 = v344
		v400 = v339
		goto L60
	} else {
		goto L61
	}
L60:
	;
	if int32(305) < v399 {
		goto L69
	} else {
		goto L70
	}
L61:
	;
	if int32(0) <= v342 {
		v399 = v344
		v400 = v339
		goto L60
	} else {
		goto L62
	}
L62:
	;
	v350 = v344 & int32(63)
	v352 = v303 - int32(2)
	v354 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v318+v352))))
	v356 = v354 << (uint(int32(6)) % 32)
	if base.B2i32(v352 != v322)&base.B2i32(base.Ui32(v354) < base.Ui32(int32(192))) == int32(0) {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v399 = v356&int32(1984) | v350
	v400 = int32(2)
	goto L60
L64:
	;
	goto L65
L65:
	;
	v369 = v356&int32(4032) | v350
	v371 = v303 - int32(3)
	v373 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v318+v371))))
	if base.B2i32(v371 != v322)&base.B2i32(base.Ui32(v373) < base.Ui32(int32(224))) == int32(0) {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v399 = v373<<(uint(int32(12))%32)&int32(61440) | v369
	v400 = int32(3)
	goto L60
L67:
	;
	goto L68
L68:
	;
	v391 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v303+(v318-int32(4))))))
	v399 = v373<<(uint(int32(12))%32)&int32(258048) | v391&int32(7)<<(uint(int32(18))%32) | v369
	v400 = int32(4)
	goto L60
L69:
	;
	v433 = v400
	goto L55
L70:
	;
	goto L71
L71:
	;
	v404 = v399 - int32(105)
	if v404 < int32(0) {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v433 = v400
	goto L55
L73:
	;
	goto L74
L74:
	;
	v410 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v404)>>(uint(int32(3))%32)))+uint32(_consts[1340]))))
	if int32(base.Ui32(v410)>>(uint(v404&int32(7))%32))&int32(1) == int32(0) {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v433 = v400
	goto L55
L76:
	;
	goto L77
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v303 - v400
	goto L78
L78:
	;
	goto L58
L79:
	;
	v439 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v440 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L82
L80:
	;
	if v491 < int32(0) {
		v635 = v2
		goto L1
	} else {
		goto L100
	}
L82:
	;
	goto L83
L83:
	;
	goto L84
L84:
	;
	v447 = v435
	v449 = int32(1)
	goto L87
L86:
	;
	v491 = v473
	goto L80
L87:
	;
	if v447 <= v440 {
		goto L89
	} else {
		goto L90
	}
L88:
	;
	goto L86
L89:
	;
	v491 = int32(-1)
	goto L80
L90:
	;
	goto L91
L91:
	;
	v454 = v447 - int32(1)
	v456 = int32(*(*int8)(unsafe.Add(mBase, uint32(v439+v454))))
	if int32(0) <= v456 {
		v473 = v454
		goto L92
	} else {
		goto L93
	}
L92:
	;
	v477 = int32(1)
	if v477 < v449 {
		v447 = v473
		v449 = v449 - v477
		goto L87
	} else {
		goto L99
	}
L93:
	;
	if v454 <= v440 {
		v473 = v454
		goto L92
	} else {
		goto L94
	}
L94:
	;
	v461 = v454
	goto L95
L95:
	;
	v466 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v439+v461))))
	if base.Ui32(int32(191)) < base.Ui32(v466) {
		v473 = v461
		goto L92
	} else {
		goto L97
	}
L96:
	;
	v473 = v440
	goto L92
L97:
	;
	v470 = v461 - int32(1)
	if v440 < v470 {
		v461 = v470
		goto L95
	} else {
		goto L98
	}
L98:
	;
	goto L96
L99:
	;
	goto L88
L100:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v491
	v508 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v512 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L103
L101:
	;
	if v624 != 0 {
		v635 = v2
		goto L1
	} else {
		goto L120
	}
L102:
	;
	v624 = v617
	goto L101
L103:
	;
	if v491 <= v512 {
		v617 = int32(-1)
		goto L102
	} else {
		goto L105
	}
L104:
	;
	v617 = int32(0)
	goto L102
L105:
	;
	v529 = int32(1)
	v530 = v491 - v529
	v532 = int32(*(*int8)(unsafe.Add(mBase, uint32(v508+v530))))
	v534 = v532 & int32(255)
	if v530 == v512 {
		v589 = v534
		v590 = v529
		goto L106
	} else {
		goto L107
	}
L106:
	;
	if int32(305) < v589 {
		goto L115
	} else {
		goto L116
	}
L107:
	;
	if int32(0) <= v532 {
		v589 = v534
		v590 = v529
		goto L106
	} else {
		goto L108
	}
L108:
	;
	v540 = v534 & int32(63)
	v542 = v491 - int32(2)
	v544 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v508+v542))))
	v546 = v544 << (uint(int32(6)) % 32)
	if base.B2i32(v542 != v512)&base.B2i32(base.Ui32(v544) < base.Ui32(int32(192))) == int32(0) {
		goto L109
	} else {
		goto L110
	}
L109:
	;
	v589 = v546&int32(1984) | v540
	v590 = int32(2)
	goto L106
L110:
	;
	goto L111
L111:
	;
	v559 = v546&int32(4032) | v540
	v561 = v491 - int32(3)
	v563 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v508+v561))))
	if base.B2i32(v561 != v512)&base.B2i32(base.Ui32(v563) < base.Ui32(int32(224))) == int32(0) {
		goto L112
	} else {
		goto L113
	}
L112:
	;
	v589 = v563<<(uint(int32(12))%32)&int32(61440) | v559
	v590 = int32(3)
	goto L106
L113:
	;
	goto L114
L114:
	;
	v581 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v491+(v508-int32(4))))))
	v589 = v563<<(uint(int32(12))%32)&int32(258048) | v581&int32(7)<<(uint(int32(18))%32) | v559
	v590 = int32(4)
	goto L106
L115:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v491 - v590
	goto L119
L116:
	;
	v594 = v589 - int32(97)
	if v594 < int32(0) {
		goto L115
	} else {
		goto L117
	}
L117:
	;
	v600 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v594)>>(uint(int32(3))%32)))+uint32(_consts[1333]))))
	if int32(base.Ui32(v600)>>(uint(v594&int32(7))%32))&int32(1) == int32(0) {
		goto L115
	} else {
		goto L118
	}
L118:
	;
	v624 = v590
	goto L101
L119:
	;
	goto L104
L120:
	;
	v625 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v631 = v625 + v301
	goto L8
}
