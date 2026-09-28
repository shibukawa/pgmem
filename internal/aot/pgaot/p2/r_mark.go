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
	v5 = Fn14363(m, l0, int32(2), int32(_a_F_r_mark_lAr_0), int32(114))
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
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v142 int32
	_ = v142
	var v158 int32
	_ = v158
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v261 int32
	_ = v261
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v273 int32
	_ = v273
	var v290 int32
	_ = v290
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v342 int32
	_ = v342
	var v349 int32
	_ = v349
	var v351 int32
	_ = v351
	var v353 int32
	_ = v353
	var v355 int32
	_ = v355
	var v368 int32
	_ = v368
	var v370 int32
	_ = v370
	var v372 int32
	_ = v372
	var v390 int32
	_ = v390
	var v392 int32
	_ = v392
	var v400 int32
	_ = v400
	var v404 int32
	_ = v404
	var v406 int32
	_ = v406
	var v412 int32
	_ = v412
	var v428 int32
	_ = v428
	var v435 int32
	_ = v435
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v449 int32
	_ = v449
	var v451 int32
	_ = v451
	var v456 int32
	_ = v456
	var v458 int32
	_ = v458
	var v464 int32
	_ = v464
	var v469 int32
	_ = v469
	var v473 int32
	_ = v473
	var v476 int32
	_ = v476
	var v480 int32
	_ = v480
	var v494 int32
	_ = v494
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v531 int32
	_ = v531
	var v533 int32
	_ = v533
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
	var v583 int32
	_ = v583
	var v591 int32
	_ = v591
	var v595 int32
	_ = v595
	var v597 int32
	_ = v597
	var v603 int32
	_ = v603
	var v620 int32
	_ = v620
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v634 int32
	_ = v634
	var v641 int32
	_ = v641
	v2 = int32(0)
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v7 <= v8 {
		v641 = v2
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v641
L2:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v12 = int32(1)
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10+v7-v12))))
	if base.B2i32(v14&int32(224) != int32(96))|base.B2i32(v12<<(uint(v14)%32)&int32(67133440) == int32(0)) != 0 {
		v641 = v2
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v29 = F_find_among_b(m, l0, int32(_a_F_r_mark_possessives_0), int32(10), int32(0))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return int32(0)
L5:
	;
	if v29 == int32(0) {
		v641 = v2
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L11
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v634
	v641 = int32(1)
	goto L1
L8:
	;
	v303 = v36 - v35
	v304 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v305 = v303 + v304
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v305
	v320 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v321 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L54
L9:
	;
	if v165 != 0 {
		goto L8
	} else {
		goto L32
	}
L10:
	;
	v165 = v158
	goto L9
L11:
	;
	if v36 <= v50 {
		v158 = int32(-1)
		goto L10
	} else {
		goto L13
	}
L12:
	;
	v158 = int32(0)
	goto L10
L13:
	;
	v67 = int32(1)
	v68 = v36 - v67
	v70 = int32(*(*int8)(unsafe.Add(mBase, uint32(v51+v68))))
	v72 = v70 & int32(255)
	if base.B2i32(v68 == v50)|base.B2i32(int32(0) <= v70) != 0 {
		v130 = v72
		v134 = v67
		goto L14
	} else {
		goto L15
	}
L14:
	;
	if int32(305) < v130 {
		goto L22
	} else {
		goto L23
	}
L15:
	;
	v79 = v72 & int32(63)
	v81 = v36 - int32(2)
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51+v81))))
	v85 = v83 << (uint(int32(6)) % 32)
	if base.B2i32(v81 != v50)&base.B2i32(base.Ui32(v83) < base.Ui32(int32(192))) == int32(0) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v130 = v85&int32(1984) | v79
	v134 = int32(2)
	goto L14
L17:
	;
	goto L18
L18:
	;
	v98 = v85&int32(4032) | v79
	v100 = v36 - int32(3)
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51+v100))))
	if base.B2i32(v100 != v50)&base.B2i32(base.Ui32(v102) < base.Ui32(int32(224))) == int32(0) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v130 = v102<<(uint(int32(12))%32)&int32(_a_F_r_mark_possessives_1) | v98
	v134 = int32(3)
	goto L14
L20:
	;
	goto L21
L21:
	;
	v120 = int32(4)
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36+v51-v120))))
	v130 = v102<<(uint(int32(12))%32)&int32(_a_F_r_mark_possessives_2) | v122&int32(7)<<(uint(int32(18))%32) | v98
	v134 = v120
	goto L14
L22:
	;
	v165 = v134
	goto L9
L23:
	;
	goto L24
L24:
	;
	v136 = v130 - int32(105)
	if v136 < int32(0) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v165 = v134
	goto L9
L26:
	;
	goto L27
L27:
	;
	v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v136)>>(uint(int32(3))%32)))+uint32(_c_F_r_mark_possessives[0]))))
	if int32(base.Ui32(v142)>>(uint(v136&int32(7))%32))&int32(1) == int32(0) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v165 = v134
	goto L9
L29:
	;
	goto L30
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v36 - v134
	goto L31
L31:
	;
	goto L12
L32:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v167 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v181 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v182 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L35
L33:
	;
	if v297 != 0 {
		goto L8
	} else {
		goto L51
	}
L34:
	;
	v297 = v290
	goto L33
L35:
	;
	if v166 <= v181 {
		v290 = int32(-1)
		goto L34
	} else {
		goto L37
	}
L36:
	;
	v290 = int32(0)
	goto L34
L37:
	;
	v198 = int32(1)
	v199 = v166 - v198
	v201 = int32(*(*int8)(unsafe.Add(mBase, uint32(v182+v199))))
	v203 = v201 & int32(255)
	if base.B2i32(v199 == v181)|base.B2i32(int32(0) <= v201) != 0 {
		v261 = v203
		v265 = v198
		goto L38
	} else {
		goto L39
	}
L38:
	;
	if int32(305) < v261 {
		goto L46
	} else {
		goto L47
	}
L39:
	;
	v210 = v203 & int32(63)
	v212 = v166 - int32(2)
	v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182+v212))))
	v216 = v214 << (uint(int32(6)) % 32)
	if base.B2i32(v212 != v181)&base.B2i32(base.Ui32(v214) < base.Ui32(int32(192))) == int32(0) {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v261 = v216&int32(1984) | v210
	v265 = int32(2)
	goto L38
L41:
	;
	goto L42
L42:
	;
	v229 = v216&int32(4032) | v210
	v231 = v166 - int32(3)
	v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182+v231))))
	if base.B2i32(v231 != v181)&base.B2i32(base.Ui32(v233) < base.Ui32(int32(224))) == int32(0) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v261 = v233<<(uint(int32(12))%32)&int32(_a_F_r_mark_possessives_1) | v229
	v265 = int32(3)
	goto L38
L44:
	;
	goto L45
L45:
	;
	v251 = int32(4)
	v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166+v182-v251))))
	v261 = v233<<(uint(int32(12))%32)&int32(_a_F_r_mark_possessives_2) | v253&int32(7)<<(uint(int32(18))%32) | v229
	v265 = v251
	goto L38
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v166 - v265
	goto L50
L47:
	;
	v267 = v261 - int32(97)
	if v267 < int32(0) {
		goto L46
	} else {
		goto L48
	}
L48:
	;
	v273 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v267)>>(uint(int32(3))%32)))+uint32(_c_F_r_mark_possessives[1]))))
	if int32(base.Ui32(v273)>>(uint(v267&int32(7))%32))&int32(1) == int32(0) {
		goto L46
	} else {
		goto L49
	}
L49:
	;
	v297 = v265
	goto L33
L50:
	;
	goto L36
L51:
	;
	v298 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v634 = v298 + (v166 - v167)
	goto L7
L52:
	;
	if v435 == int32(0) {
		v641 = v2
		goto L1
	} else {
		goto L75
	}
L53:
	;
	v435 = v428
	goto L52
L54:
	;
	if v305 <= v320 {
		v428 = int32(-1)
		goto L53
	} else {
		goto L56
	}
L55:
	;
	v428 = int32(0)
	goto L53
L56:
	;
	v337 = int32(1)
	v338 = v305 - v337
	v340 = int32(*(*int8)(unsafe.Add(mBase, uint32(v321+v338))))
	v342 = v340 & int32(255)
	if base.B2i32(v338 == v320)|base.B2i32(int32(0) <= v340) != 0 {
		v400 = v342
		v404 = v337
		goto L57
	} else {
		goto L58
	}
L57:
	;
	if int32(305) < v400 {
		goto L65
	} else {
		goto L66
	}
L58:
	;
	v349 = v342 & int32(63)
	v351 = v305 - int32(2)
	v353 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v321+v351))))
	v355 = v353 << (uint(int32(6)) % 32)
	if base.B2i32(v351 != v320)&base.B2i32(base.Ui32(v353) < base.Ui32(int32(192))) == int32(0) {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v400 = v355&int32(1984) | v349
	v404 = int32(2)
	goto L57
L60:
	;
	goto L61
L61:
	;
	v368 = v355&int32(4032) | v349
	v370 = v305 - int32(3)
	v372 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v321+v370))))
	if base.B2i32(v370 != v320)&base.B2i32(base.Ui32(v372) < base.Ui32(int32(224))) == int32(0) {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v400 = v372<<(uint(int32(12))%32)&int32(_a_F_r_mark_possessives_1) | v368
	v404 = int32(3)
	goto L57
L63:
	;
	goto L64
L64:
	;
	v390 = int32(4)
	v392 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v305+v321-v390))))
	v400 = v372<<(uint(int32(12))%32)&int32(_a_F_r_mark_possessives_2) | v392&int32(7)<<(uint(int32(18))%32) | v368
	v404 = v390
	goto L57
L65:
	;
	v435 = v404
	goto L52
L66:
	;
	goto L67
L67:
	;
	v406 = v400 - int32(105)
	if v406 < int32(0) {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v435 = v404
	goto L52
L69:
	;
	goto L70
L70:
	;
	v412 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v406)>>(uint(int32(3))%32)))+uint32(_c_F_r_mark_possessives[0]))))
	if int32(base.Ui32(v412)>>(uint(v406&int32(7))%32))&int32(1) == int32(0) {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v435 = v404
	goto L52
L72:
	;
	goto L73
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v305 - v404
	goto L74
L74:
	;
	goto L55
L75:
	;
	v438 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v439 = v438 + v303
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v439
	v441 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v442 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L78
L76:
	;
	if v494 < int32(0) {
		v641 = v2
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
	v449 = v439
	v451 = int32(1)
	goto L83
L82:
	;
	v494 = v476
	goto L76
L83:
	;
	if v449 <= v442 {
		goto L85
	} else {
		goto L86
	}
L84:
	;
	goto L82
L85:
	;
	v494 = int32(-1)
	goto L76
L86:
	;
	goto L87
L87:
	;
	v456 = v449 - int32(1)
	v458 = int32(*(*int8)(unsafe.Add(mBase, uint32(v441+v456))))
	if base.B2i32(int32(0) <= v458)|base.B2i32(v456 <= v442) != 0 {
		v476 = v456
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v480 = int32(1)
	if v480 < v451 {
		v449 = v476
		v451 = v451 - v480
		goto L83
	} else {
		goto L94
	}
L89:
	;
	v464 = v456
	goto L90
L90:
	;
	v469 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v441+v464))))
	if base.Ui32(int32(191)) < base.Ui32(v469) {
		v476 = v464
		goto L88
	} else {
		goto L92
	}
L91:
	;
	v476 = v442
	goto L88
L92:
	;
	v473 = v464 - int32(1)
	if v442 < v473 {
		v464 = v473
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
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v494
	v511 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v512 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L98
L96:
	;
	if v627 != 0 {
		v641 = v2
		goto L1
	} else {
		goto L114
	}
L97:
	;
	v627 = v620
	goto L96
L98:
	;
	if v494 <= v511 {
		v620 = int32(-1)
		goto L97
	} else {
		goto L100
	}
L99:
	;
	v620 = int32(0)
	goto L97
L100:
	;
	v528 = int32(1)
	v529 = v494 - v528
	v531 = int32(*(*int8)(unsafe.Add(mBase, uint32(v512+v529))))
	v533 = v531 & int32(255)
	if base.B2i32(v529 == v511)|base.B2i32(int32(0) <= v531) != 0 {
		v591 = v533
		v595 = v528
		goto L101
	} else {
		goto L102
	}
L101:
	;
	if int32(305) < v591 {
		goto L109
	} else {
		goto L110
	}
L102:
	;
	v540 = v533 & int32(63)
	v542 = v494 - int32(2)
	v544 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v512+v542))))
	v546 = v544 << (uint(int32(6)) % 32)
	if base.B2i32(v542 != v511)&base.B2i32(base.Ui32(v544) < base.Ui32(int32(192))) == int32(0) {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	v591 = v546&int32(1984) | v540
	v595 = int32(2)
	goto L101
L104:
	;
	goto L105
L105:
	;
	v559 = v546&int32(4032) | v540
	v561 = v494 - int32(3)
	v563 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v512+v561))))
	if base.B2i32(v561 != v511)&base.B2i32(base.Ui32(v563) < base.Ui32(int32(224))) == int32(0) {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	v591 = v563<<(uint(int32(12))%32)&int32(_a_F_r_mark_possessives_1) | v559
	v595 = int32(3)
	goto L101
L107:
	;
	goto L108
L108:
	;
	v581 = int32(4)
	v583 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v494+v512-v581))))
	v591 = v563<<(uint(int32(12))%32)&int32(_a_F_r_mark_possessives_2) | v583&int32(7)<<(uint(int32(18))%32) | v559
	v595 = v581
	goto L101
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v494 - v595
	goto L113
L110:
	;
	v597 = v591 - int32(97)
	if v597 < int32(0) {
		goto L109
	} else {
		goto L111
	}
L111:
	;
	v603 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v597)>>(uint(int32(3))%32)))+uint32(_c_F_r_mark_possessives[1]))))
	if int32(base.Ui32(v603)>>(uint(v597&int32(7))%32))&int32(1) == int32(0) {
		goto L109
	} else {
		goto L112
	}
L112:
	;
	v627 = v595
	goto L96
L113:
	;
	goto L99
L114:
	;
	v628 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v634 = v628 + v303
	goto L7
}
