package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_FindLockCycleRecurseMember(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v67 int32
	_ = v67
	var v74 int32
	_ = v74
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v117 int32
	_ = v117
	var v124 int32
	_ = v124
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v166 int32
	_ = v166
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v182 int32
	_ = v182
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v205 int64
	_ = v205
	var v207 int64
	_ = v207
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v240 int32
	_ = v240
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	var v269 int32
	_ = v269
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v339 int32
	_ = v339
	var v348 int32
	_ = v348
	var v354 int32
	_ = v354
	var v360 int32
	_ = v360
	var v371 int32
	_ = v371
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v387 int32
	_ = v387
	var v389 int32
	_ = v389
	var v392 int32
	_ = v392
	var v402 int32
	_ = v402
	var v409 int32
	_ = v409
	var v426 int32
	_ = v426
	var v429 int32
	_ = v429
	var v432 int32
	_ = v432
	var v436 int32
	_ = v436
	var v440 int32
	_ = v440
	var v442 int32
	_ = v442
	var v451 int32
	_ = v451
	var v456 int32
	_ = v456
	var v460 int32
	_ = v460
	var v462 int32
	_ = v462
	var v467 int32
	_ = v467
	var v484 int32
	_ = v484
	var v486 int32
	_ = v486
	var v488 int32
	_ = v488
	var v493 int32
	_ = v493
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v505 int32
	_ = v505
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v520 int32
	_ = v520
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v531 int32
	_ = v531
	var v533 int32
	_ = v533
	var v536 int32
	_ = v536
	var v546 int32
	_ = v546
	var v553 int32
	_ = v553
	var v570 int32
	_ = v570
	var v573 int32
	_ = v573
	var v576 int32
	_ = v576
	var v580 int32
	_ = v580
	var v584 int32
	_ = v584
	var v586 int32
	_ = v586
	var v595 int32
	_ = v595
	var v600 int32
	_ = v600
	var v604 int32
	_ = v604
	var v606 int32
	_ = v606
	var v611 int32
	_ = v611
	var v628 int32
	_ = v628
	var v630 int32
	_ = v630
	var v633 int32
	_ = v633
	var v636 int32
	_ = v636
	var v637 int64
	_ = v637
	var v639 int64
	_ = v639
	var v641 int32
	_ = v641
	var v643 int32
	_ = v643
	var v645 int32
	_ = v645
	var v646 int32
	_ = v646
	var v650 int32
	_ = v650
	var v655 int32
	_ = v655
	var v660 int32
	_ = v660
	var v661 int32
	_ = v661
	var v667 int32
	_ = v667
	var v670 int32
	_ = v670
	var v671 int64
	_ = v671
	var v673 int64
	_ = v673
	var v675 int32
	_ = v675
	var v677 int32
	_ = v677
	var v679 int32
	_ = v679
	var v680 int32
	_ = v680
	var v684 int32
	_ = v684
	var v689 int32
	_ = v689
	var v694 int32
	_ = v694
	var v695 int32
	_ = v695
	var v710 int32
	_ = v710
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+14)))
	if v17 == int32(1) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	goto L3
L3:
	;
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+15)))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v22<<(uint(int32(2))%32))+uint32(_c_F_FindLockCycleRecurseMember[0])))
	goto L4
L4:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v26+v27<<(uint(int32(2))%32))))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v16)+28))
	if v32 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v257 = int32(0)
	v259 = *(*int32)(unsafe.Add(mBase, _c_F_FindLockCycleRecurseMember[1]))
	if v257 < v259 {
		goto L64
	} else {
		goto L65
	}
L6:
	;
	v36 = v16 + int32(24)
	if v32 == v36 {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	v40 = l2 + int32(1)
	v54 = v32
	goto L8
L8:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v54-int32(16))))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)+616))
	if v59 != 0 {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L5
L10:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v54)+4))
	if v240 != v36 {
		v54 = v240
		goto L8
	} else {
		goto L60
	}
L11:
	;
	v60 = v59
	goto L13
L12:
	;
	v60 = v58
	goto L13
L13:
	;
	if base.B2i32(v60 == l1)|base.B2i32(v38 <= int32(0)) != 0 {
		goto L10
	} else {
		goto L14
	}
L14:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v54-int32(8))))
	v74 = int32(1)
	goto L15
L15:
	;
	v85 = int32(1) << (uint(v74) % 32)
	if v85&v31 != 0 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v58)+616))
	if v98 != 0 {
		goto L25
	} else {
		goto L26
	}
L17:
	;
	v89 = v67 & v85
	goto L19
L18:
	;
	v89 = int32(0)
	goto L19
L19:
	;
	if v89 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v93 = v74 + int32(1)
	if v93 <= v38 {
		v74 = v93
		goto L15
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	goto L16
L23:
	;
	goto L10
L24:
	;
	if v199 != 0 {
		goto L55
	} else {
		goto L56
	}
L25:
	;
	v99 = v98
	goto L27
L26:
	;
	v99 = v58
	goto L27
L27:
	;
	v100 = int32(0)
	v102 = *(*int32)(unsafe.Add(mBase, _c_F_FindLockCycleRecurseMember[2]))
	v104 = *(*int32)(unsafe.Add(mBase, _c_F_FindLockCycleRecurseMember[3]))
	if v100 < v104 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v107 = v100
	goto L31
L29:
	;
	goto L30
L30:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_FindLockCycleRecurseMember[3])) = v104 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v102+v104<<(uint(int32(2))%32)))) = v99
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v99)+4))
	if v141 == int32(0) {
		goto L40
	} else {
		goto L41
	}
L31:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v102+v107<<(uint(int32(2))%32))))
	if v99 == v117 {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	goto L30
L33:
	;
	if v107 != 0 {
		goto L36
	} else {
		goto L37
	}
L34:
	;
	goto L35
L35:
	;
	v124 = v107 + int32(1)
	if v124 != v104 {
		v107 = v124
		goto L31
	} else {
		goto L39
	}
L36:
	;
	v199 = int32(0)
	goto L24
L37:
	;
	goto L38
L38:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_FindLockCycleRecurseMember[4])) = v40
	v199 = int32(1)
	goto L24
L39:
	;
	goto L32
L40:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v99)+624))
	if v151 == int32(0) {
		goto L44
	} else {
		goto L45
	}
L41:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v99)+92))
	if v144 == int32(0) {
		goto L40
	} else {
		goto L42
	}
L42:
	;
	v147 = F_FindLockCycleRecurseMember(m, v99, v99, v40, l3, l4)
	mBase = m.M
	if v147 == int32(0) {
		goto L40
	} else {
		goto L43
	}
L43:
	;
	v199 = int32(1)
	goto L24
L44:
	;
	v199 = int32(0)
	goto L24
L45:
	;
	v155 = v99 + int32(620)
	if v151 == v155 {
		goto L44
	} else {
		goto L46
	}
L46:
	;
	v157 = v151
	goto L47
L47:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v157-int32(624))))
	if v166 == int32(0) {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	goto L44
L49:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v157)+4))
	if v182 != v155 {
		v157 = v182
		goto L47
	} else {
		goto L54
	}
L50:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v157-int32(536))))
	if v171 == int32(0) {
		goto L49
	} else {
		goto L51
	}
L51:
	;
	v175 = v157 - int32(628)
	if v175 == v99 {
		goto L49
	} else {
		goto L52
	}
L52:
	;
	v177 = F_FindLockCycleRecurseMember(m, v175, v99, v40, l3, l4)
	mBase = m.M
	if v177 == int32(0) {
		goto L49
	} else {
		goto L53
	}
L53:
	;
	v199 = int32(1)
	goto L24
L54:
	;
	goto L48
L55:
	;
	v201 = *(*int32)(unsafe.Add(mBase, _c_F_FindLockCycleRecurseMember[5]))
	v204 = v201 + l2*int32(24)
	v205 = *(*int64)(unsafe.Add(mBase, uint32(v16)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v204)+8)) = v205
	v207 = *(*int64)(unsafe.Add(mBase, uint32(v16)))
	*(*int64)(unsafe.Add(mBase, uint32(v204))) = v207
	v209 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	*(*int32)(unsafe.Add(mBase, uint32(v204)+16)) = v209
	v211 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v204)+20)) = v211
	return int32(1)
L56:
	;
	goto L57
L57:
	;
	v216 = *(*int32)(unsafe.Add(mBase, _c_F_FindLockCycleRecurseMember[6]))
	if l0 != v216 {
		goto L10
	} else {
		goto L58
	}
L58:
	;
	v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+124)))
	if v218&int32(1) == int32(0) {
		goto L10
	} else {
		goto L59
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_FindLockCycleRecurseMember[7])) = v58
	goto L10
L60:
	;
	goto L9
L61:
	;
	return v710
L62:
	;
	v667 = *(*int32)(unsafe.Add(mBase, _c_F_FindLockCycleRecurseMember[5]))
	v670 = v667 + l2*int32(24)
	v671 = *(*int64)(unsafe.Add(mBase, uint32(v16)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v670)+8)) = v671
	v673 = *(*int64)(unsafe.Add(mBase, uint32(v16)))
	*(*int64)(unsafe.Add(mBase, uint32(v670))) = v673
	v675 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	*(*int32)(unsafe.Add(mBase, uint32(v670)+16)) = v675
	v677 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v670)+20)) = v677
	v679 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	v680 = int32(20)
	*(*int32)(unsafe.Add(mBase, uint32(l3+v679*v680))) = l1
	v684 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	*(*int32)(unsafe.Add(mBase, uint32(l3+v684*v680)+4)) = v378
	v689 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	*(*int32)(unsafe.Add(mBase, uint32(l3+v689*v680)+8)) = v16
	v694 = int32(1)
	v695 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v695 + v694
	v710 = v694
	goto L61
L63:
	;
	v488 = *(*int32)(unsafe.Add(mBase, uint32(v281)+8))
	if v488 <= int32(0) {
		goto L123
	} else {
		goto L124
	}
L64:
	;
	v263 = *(*int32)(unsafe.Add(mBase, _c_F_FindLockCycleRecurseMember[8]))
	v269 = v257
	goto L67
L65:
	;
	goto L66
L66:
	;
	v303 = v16 + int32(32)
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v16)+36))
	v305 = *(*int32)(unsafe.Add(mBase, uint32(l0)+616))
	if v305 == int32(0) {
		v339 = l0
		goto L71
	} else {
		goto L72
	}
L67:
	;
	v281 = v263 + v269*int32(12)
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v281)))
	if v282 == v16 {
		goto L63
	} else {
		goto L69
	}
L68:
	;
	goto L66
L69:
	;
	v285 = v269 + int32(1)
	if v285 != v259 {
		v269 = v285
		goto L67
	} else {
		goto L70
	}
L70:
	;
	goto L68
L71:
	;
	v348 = int32(0)
	if base.B2i32(v304 == v348)|base.B2i32(v304 == v303) != 0 {
		v710 = v348
		goto L61
	} else {
		goto L80
	}
L72:
	;
	v308 = int32(0)
	if base.B2i32(v304 == v308)|base.B2i32(v304 == v303) != 0 {
		v339 = v308
		goto L71
	} else {
		goto L73
	}
L73:
	;
	v319 = v308
	v320 = v304
	goto L74
L74:
	;
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v320)+616))
	if v328 == l1 {
		goto L76
	} else {
		goto L77
	}
L75:
	;
	v339 = v330
	goto L71
L76:
	;
	v330 = v320
	goto L78
L77:
	;
	v330 = v319
	goto L78
L78:
	;
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v320)+4))
	if v331 != v303 {
		v319 = v330
		v320 = v331
		goto L74
	} else {
		goto L79
	}
L79:
	;
	goto L75
L80:
	;
	v354 = l2 + int32(1)
	v360 = v304
	goto L81
L81:
	;
	if v360 == v339 {
		v710 = v348
		goto L61
	} else {
		goto L83
	}
L82:
	;
	v710 = v348
	goto L61
L83:
	;
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v360)+100))
	if int32(base.Ui32(v31)>>(uint(v371)%32))&int32(1) == int32(0) {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	v486 = *(*int32)(unsafe.Add(mBase, uint32(v360)+4))
	if v486 != v303 {
		v360 = v486
		goto L81
	} else {
		goto L122
	}
L85:
	;
	v377 = *(*int32)(unsafe.Add(mBase, uint32(v360)+616))
	if v377 != 0 {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v378 = v377
	goto L88
L87:
	;
	v378 = v360
	goto L88
L88:
	;
	if v378 == l1 {
		goto L84
	} else {
		goto L89
	}
L89:
	;
	v383 = *(*int32)(unsafe.Add(mBase, uint32(v360)+616))
	if v383 != 0 {
		goto L91
	} else {
		goto L92
	}
L90:
	;
	if v484 != 0 {
		goto L62
	} else {
		goto L121
	}
L91:
	;
	v384 = v383
	goto L93
L92:
	;
	v384 = v360
	goto L93
L93:
	;
	v385 = int32(0)
	v387 = *(*int32)(unsafe.Add(mBase, _c_F_FindLockCycleRecurseMember[2]))
	v389 = *(*int32)(unsafe.Add(mBase, _c_F_FindLockCycleRecurseMember[3]))
	if v385 < v389 {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	v392 = v385
	goto L97
L95:
	;
	goto L96
L96:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_FindLockCycleRecurseMember[3])) = v389 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v387+v389<<(uint(int32(2))%32)))) = v384
	v426 = *(*int32)(unsafe.Add(mBase, uint32(v384)+4))
	if v426 == int32(0) {
		goto L106
	} else {
		goto L107
	}
L97:
	;
	v402 = *(*int32)(unsafe.Add(mBase, uint32(v387+v392<<(uint(int32(2))%32))))
	if v384 == v402 {
		goto L99
	} else {
		goto L100
	}
L98:
	;
	goto L96
L99:
	;
	if v392 != 0 {
		goto L102
	} else {
		goto L103
	}
L100:
	;
	goto L101
L101:
	;
	v409 = v392 + int32(1)
	if v409 != v389 {
		v392 = v409
		goto L97
	} else {
		goto L105
	}
L102:
	;
	v484 = int32(0)
	goto L90
L103:
	;
	goto L104
L104:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_FindLockCycleRecurseMember[4])) = v354
	v484 = int32(1)
	goto L90
L105:
	;
	goto L98
L106:
	;
	v436 = *(*int32)(unsafe.Add(mBase, uint32(v384)+624))
	if v436 == int32(0) {
		goto L110
	} else {
		goto L111
	}
L107:
	;
	v429 = *(*int32)(unsafe.Add(mBase, uint32(v384)+92))
	if v429 == int32(0) {
		goto L106
	} else {
		goto L108
	}
L108:
	;
	v432 = F_FindLockCycleRecurseMember(m, v384, v384, v354, l3, l4)
	mBase = m.M
	if v432 == int32(0) {
		goto L106
	} else {
		goto L109
	}
L109:
	;
	v484 = int32(1)
	goto L90
L110:
	;
	v484 = int32(0)
	goto L90
L111:
	;
	v440 = v384 + int32(620)
	if v436 == v440 {
		goto L110
	} else {
		goto L112
	}
L112:
	;
	v442 = v436
	goto L113
L113:
	;
	v451 = *(*int32)(unsafe.Add(mBase, uint32(v442-int32(624))))
	if v451 == int32(0) {
		goto L115
	} else {
		goto L116
	}
L114:
	;
	goto L110
L115:
	;
	v467 = *(*int32)(unsafe.Add(mBase, uint32(v442)+4))
	if v467 != v440 {
		v442 = v467
		goto L113
	} else {
		goto L120
	}
L116:
	;
	v456 = *(*int32)(unsafe.Add(mBase, uint32(v442-int32(536))))
	if v456 == int32(0) {
		goto L115
	} else {
		goto L117
	}
L117:
	;
	v460 = v442 - int32(628)
	if v460 == v384 {
		goto L115
	} else {
		goto L118
	}
L118:
	;
	v462 = F_FindLockCycleRecurseMember(m, v460, v384, v354, l3, l4)
	mBase = m.M
	if v462 == int32(0) {
		goto L115
	} else {
		goto L119
	}
L119:
	;
	v484 = int32(1)
	goto L90
L120:
	;
	goto L114
L121:
	;
	goto L84
L122:
	;
	goto L82
L123:
	;
	return int32(0)
L124:
	;
	goto L125
L125:
	;
	v493 = *(*int32)(unsafe.Add(mBase, uint32(v281)+4))
	v495 = l2 + int32(1)
	v496 = int32(0)
	v505 = v496
	goto L126
L126:
	;
	v516 = *(*int32)(unsafe.Add(mBase, uint32(v493+v505<<(uint(int32(2))%32))))
	v517 = *(*int32)(unsafe.Add(mBase, uint32(v516)+616))
	if v517 != 0 {
		goto L128
	} else {
		goto L129
	}
L127:
	;
	v633 = *(*int32)(unsafe.Add(mBase, _c_F_FindLockCycleRecurseMember[5]))
	v636 = v633 + l2*int32(24)
	v637 = *(*int64)(unsafe.Add(mBase, uint32(v16)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v636)+8)) = v637
	v639 = *(*int64)(unsafe.Add(mBase, uint32(v16)))
	*(*int64)(unsafe.Add(mBase, uint32(v636))) = v639
	v641 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	*(*int32)(unsafe.Add(mBase, uint32(v636)+16)) = v641
	v643 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v636)+20)) = v643
	v645 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	v646 = int32(20)
	*(*int32)(unsafe.Add(mBase, uint32(l3+v645*v646))) = l1
	v650 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	*(*int32)(unsafe.Add(mBase, uint32(l3+v650*v646)+4)) = v518
	v655 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	*(*int32)(unsafe.Add(mBase, uint32(l3+v655*v646)+8)) = v16
	v660 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	v661 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v660 + v661
	return v661
L128:
	;
	v518 = v517
	goto L130
L129:
	;
	v518 = v516
	goto L130
L130:
	;
	if v518 == l1 {
		v710 = v496
		goto L61
	} else {
		goto L131
	}
L131:
	;
	v520 = *(*int32)(unsafe.Add(mBase, uint32(v516)+100))
	if int32(base.Ui32(v31)>>(uint(v520)%32))&int32(1) != 0 {
		goto L133
	} else {
		goto L134
	}
L132:
	;
	goto L127
L133:
	;
	v527 = *(*int32)(unsafe.Add(mBase, uint32(v516)+616))
	if v527 != 0 {
		goto L137
	} else {
		goto L138
	}
L134:
	;
	goto L135
L135:
	;
	v630 = v505 + int32(1)
	if v630 != v488 {
		v505 = v630
		goto L126
	} else {
		goto L168
	}
L136:
	;
	if v628 != 0 {
		goto L132
	} else {
		goto L167
	}
L137:
	;
	v528 = v527
	goto L139
L138:
	;
	v528 = v516
	goto L139
L139:
	;
	v529 = int32(0)
	v531 = *(*int32)(unsafe.Add(mBase, _c_F_FindLockCycleRecurseMember[2]))
	v533 = *(*int32)(unsafe.Add(mBase, _c_F_FindLockCycleRecurseMember[3]))
	if v529 < v533 {
		goto L140
	} else {
		goto L141
	}
L140:
	;
	v536 = v529
	goto L143
L141:
	;
	goto L142
L142:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_FindLockCycleRecurseMember[3])) = v533 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v531+v533<<(uint(int32(2))%32)))) = v528
	v570 = *(*int32)(unsafe.Add(mBase, uint32(v528)+4))
	if v570 == int32(0) {
		goto L152
	} else {
		goto L153
	}
L143:
	;
	v546 = *(*int32)(unsafe.Add(mBase, uint32(v531+v536<<(uint(int32(2))%32))))
	if v528 == v546 {
		goto L145
	} else {
		goto L146
	}
L144:
	;
	goto L142
L145:
	;
	if v536 != 0 {
		goto L148
	} else {
		goto L149
	}
L146:
	;
	goto L147
L147:
	;
	v553 = v536 + int32(1)
	if v553 != v533 {
		v536 = v553
		goto L143
	} else {
		goto L151
	}
L148:
	;
	v628 = int32(0)
	goto L136
L149:
	;
	goto L150
L150:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_FindLockCycleRecurseMember[4])) = v495
	v628 = int32(1)
	goto L136
L151:
	;
	goto L144
L152:
	;
	v580 = *(*int32)(unsafe.Add(mBase, uint32(v528)+624))
	if v580 == int32(0) {
		goto L156
	} else {
		goto L157
	}
L153:
	;
	v573 = *(*int32)(unsafe.Add(mBase, uint32(v528)+92))
	if v573 == int32(0) {
		goto L152
	} else {
		goto L154
	}
L154:
	;
	v576 = F_FindLockCycleRecurseMember(m, v528, v528, v495, l3, l4)
	mBase = m.M
	if v576 == int32(0) {
		goto L152
	} else {
		goto L155
	}
L155:
	;
	v628 = int32(1)
	goto L136
L156:
	;
	v628 = int32(0)
	goto L136
L157:
	;
	v584 = v528 + int32(620)
	if v580 == v584 {
		goto L156
	} else {
		goto L158
	}
L158:
	;
	v586 = v580
	goto L159
L159:
	;
	v595 = *(*int32)(unsafe.Add(mBase, uint32(v586-int32(624))))
	if v595 == int32(0) {
		goto L161
	} else {
		goto L162
	}
L160:
	;
	goto L156
L161:
	;
	v611 = *(*int32)(unsafe.Add(mBase, uint32(v586)+4))
	if v611 != v584 {
		v586 = v611
		goto L159
	} else {
		goto L166
	}
L162:
	;
	v600 = *(*int32)(unsafe.Add(mBase, uint32(v586-int32(536))))
	if v600 == int32(0) {
		goto L161
	} else {
		goto L163
	}
L163:
	;
	v604 = v586 - int32(628)
	if v604 == v528 {
		goto L161
	} else {
		goto L164
	}
L164:
	;
	v606 = F_FindLockCycleRecurseMember(m, v604, v528, v495, l3, l4)
	mBase = m.M
	if v606 == int32(0) {
		goto L161
	} else {
		goto L165
	}
L165:
	;
	v628 = int32(1)
	goto L136
L166:
	;
	goto L160
L167:
	;
	goto L135
L168:
	;
	v710 = v496
	goto L61
}
func F_ForceSyncCommit(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_ForceSyncCommit[0])) = uint8(v2)
	return
}
func F_ForgetBackgroundWorker(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	v5 = m.G0
	v6 = int32(16)
	v7 = v5 - v6
	m.G0 = v7
	v10 = *(*int32)(unsafe.Add(mBase, _c_F_ForgetBackgroundWorker[0]))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1472))
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+192)))
	if v17&v6 != 0 {
		v20 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
		*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v20 + int32(1)
	} else {
	}
	v24 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v10+v11*int32(1480)+v6))) = uint8(v24)
	v28 = F_errstart(m, int32(14), v24)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		return
	} else {
		if v28 != 0 {
			*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
			F_errmsg_internal(m, int32(_a_F_ForgetBackgroundWorker_0), v7)
			mBase = m.M
			v33 = m.ExcPending
			if v33 != 0 {
				return
			} else {
				F_errfinish(m, int32(_a_F_ForgetBackgroundWorker_1), int32(449), int32(_a_F_ForgetBackgroundWorker_2))
				mBase = m.M
				v38 = m.ExcPending
				if v38 != 0 {
					return
				} else {
					v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1480))
					v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1484))
					*(*int32)(unsafe.Add(mBase, uint32(v39)+4)) = v40
					v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1480))
					*(*int32)(unsafe.Add(mBase, uint32(v40))) = v42
					F_pfree(m, l0)
					mBase = m.M
					v45 = m.ExcPending
					if v45 != 0 {
						return
					} else {
						m.G0 = v7 + int32(16)
						return
					}
				}
			}
		} else {
			v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1480))
			v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1484))
			*(*int32)(unsafe.Add(mBase, uint32(v39)+4)) = v40
			v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1480))
			*(*int32)(unsafe.Add(mBase, uint32(v40))) = v42
			F_pfree(m, l0)
			mBase = m.M
			v45 = m.ExcPending
			if v45 != 0 {
				return
			} else {
				m.G0 = v7 + int32(16)
				return
			}
		}
	}
}
func F_ForgetDatabaseSyncRequests(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	v3 = m.G0
	v5 = v3 - int32(32)
	m.G0 = v5
	*(*int64)(unsafe.Add(mBase, uint32(v5)+8)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v5)+20)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v5)+16)) = l0
	*(*int64)(unsafe.Add(mBase, uint32(v5)+24)) = int64(4294967295)
	v14 = int32(_a_F_ForgetDatabaseSyncRequests_0)
	*(*uint16)(unsafe.Add(mBase, uint32(v5)+10)) = uint16(v14)
	v20 = F_RegisterSyncRequest(m, v5+int32(8), int32(3), int32(1))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		return
	} else {
		m.G0 = v5 + int32(32)
		return
	}
}
func F_FreeDesc(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
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
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int64
	_ = v50
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch v4 {
	case 0:
		v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v34 = F_fclose(m, v33)
		mBase = m.M
		v35 = m.ExcPending
		if v35 != 0 {
			return int32(0)
		} else {
			v36 = v34
			v37 = int32(_a_F_FreeDesc_0)
			v39 = *(*int32)(unsafe.Add(mBase, _c_F_FreeDesc[0]))
			v41 = v39 - int32(1)
			*(*int32)(unsafe.Add(mBase, _c_F_FreeDesc[0])) = v41
			v44 = *(*int32)(unsafe.Add(mBase, _c_F_FreeDesc[1]))
			v47 = v44 + v41*int32(12)
			v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)+8))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v48
			v50 = *(*int64)(unsafe.Add(mBase, uint32(v47)))
			*(*int64)(unsafe.Add(mBase, uint32(l0))) = v50
			return v36
		}
	case 1:
		v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v6 = F_pgl_pclose(m, v5)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			v36 = v6
			v37 = int32(_a_F_FreeDesc_0)
			v39 = *(*int32)(unsafe.Add(mBase, _c_F_FreeDesc[0]))
			v41 = v39 - int32(1)
			*(*int32)(unsafe.Add(mBase, _c_F_FreeDesc[0])) = v41
			v44 = *(*int32)(unsafe.Add(mBase, _c_F_FreeDesc[1]))
			v47 = v44 + v41*int32(12)
			v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)+8))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v48
			v50 = *(*int64)(unsafe.Add(mBase, uint32(v47)))
			*(*int64)(unsafe.Add(mBase, uint32(l0))) = v50
			return v36
		}
	case 2:
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v12 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
		v13 = F_close(m, v12)
		mBase = m.M
		F_emscripten_builtin_free(m, v10)
		mBase = m.M
		v36 = v13
		v37 = int32(_a_F_FreeDesc_0)
		v39 = *(*int32)(unsafe.Add(mBase, _c_F_FreeDesc[0]))
		v41 = v39 - int32(1)
		*(*int32)(unsafe.Add(mBase, _c_F_FreeDesc[0])) = v41
		v44 = *(*int32)(unsafe.Add(mBase, _c_F_FreeDesc[1]))
		v47 = v44 + v41*int32(12)
		v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)+8))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v48
		v50 = *(*int64)(unsafe.Add(mBase, uint32(v47)))
		*(*int64)(unsafe.Add(mBase, uint32(l0))) = v50
		return v36
	case 3:
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		F_pgaio_closing_fd(m, v15)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v19 = F_close(m, v18)
			mBase = m.M
			v36 = v19
			v37 = int32(_a_F_FreeDesc_0)
			v39 = *(*int32)(unsafe.Add(mBase, _c_F_FreeDesc[0]))
			v41 = v39 - int32(1)
			*(*int32)(unsafe.Add(mBase, _c_F_FreeDesc[0])) = v41
			v44 = *(*int32)(unsafe.Add(mBase, _c_F_FreeDesc[1]))
			v47 = v44 + v41*int32(12)
			v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)+8))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v48
			v50 = *(*int64)(unsafe.Add(mBase, uint32(v47)))
			*(*int64)(unsafe.Add(mBase, uint32(l0))) = v50
			return v36
		}
	default:
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return int32(0)
		} else {
			F_errmsg_internal(m, int32(_a_F_FreeDesc_1), int32(0))
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(_a_F_FreeDesc_2), int32(2829), int32(_a_F_FreeDesc_3))
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
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
func F_FreeSnapshotBuilder(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v6 != 0 {
		v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+30)))
		if v7 == int32(1) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return
			} else {
				F_errmsg_internal(m, int32(_a_F_FreeSnapshotBuilder_0), int32(0))
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_FreeSnapshotBuilder_1), int32(344), int32(_a_F_FreeSnapshotBuilder_2))
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v10 = *(*int32)(unsafe.Add(mBase, uint32(v6)+44))
			v12 = v10 - int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v6)+44)) = v12
			if v12 == int32(0) {
				F_pfree(m, v6)
				mBase = m.M
				v17 = m.ExcPending
				if v17 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = int32(0)
					F_MemoryContextDelete(m, v5)
					mBase = m.M
					v22 = m.ExcPending
					if v22 != 0 {
						return
					} else {
						return
					}
				}
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = int32(0)
				F_MemoryContextDelete(m, v5)
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return
				} else {
					return
				}
			}
		}
	} else {
		F_MemoryContextDelete(m, v5)
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return
		} else {
			return
		}
	}
}
func F_FreeSpaceMapVacuum(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int64
	_ = v8
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v8 = *(*int64)(unsafe.Add(mBase, _c_F_FreeSpaceMapVacuum[0]))
	*(*int64)(unsafe.Add(mBase, uint32(v5))) = v8
	v14 = F_fsm_vacuum_page(m, l0, v5, int32(0), int32(-1), v5+int32(15))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return
	} else {
		m.G0 = v5 + int32(16)
		return
	}
}
func F_FuncnameGetCandidates(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v69 int32
	_ = v69
	var v82 int32
	_ = v82
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
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
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
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
	var v240 int32
	_ = v240
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v303 int32
	_ = v303
	var v308 int32
	_ = v308
	var v314 int32
	_ = v314
	var v320 int32
	_ = v320
	var v326 int32
	_ = v326
	var v332 int32
	_ = v332
	var v338 int32
	_ = v338
	var v344 int32
	_ = v344
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v352 int32
	_ = v352
	var v362 int32
	_ = v362
	var v393 int32
	_ = v393
	var v403 int32
	_ = v403
	var v422 int32
	_ = v422
	var v425 int32
	_ = v425
	var v434 int32
	_ = v434
	var v458 int32
	_ = v458
	var v461 int32
	_ = v461
	var v471 int32
	_ = v471
	var v479 int32
	_ = v479
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v499 int32
	_ = v499
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v513 int32
	_ = v513
	var v523 int32
	_ = v523
	var v541 int32
	_ = v541
	var v543 int32
	_ = v543
	var v557 int32
	_ = v557
	var v562 int32
	_ = v562
	var v565 int32
	_ = v565
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v576 int32
	_ = v576
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	var v590 int32
	_ = v590
	var v597 int32
	_ = v597
	var v599 int32
	_ = v599
	var v600 int32
	_ = v600
	var v604 int32
	_ = v604
	var v607 int32
	_ = v607
	var v609 int32
	_ = v609
	var v614 int32
	_ = v614
	var v616 int32
	_ = v616
	var v622 int32
	_ = v622
	var v626 int32
	_ = v626
	var v631 int32
	_ = v631
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v637 int32
	_ = v637
	var v641 int32
	_ = v641
	var v644 int32
	_ = v644
	var v645 int32
	_ = v645
	var v659 int32
	_ = v659
	var v689 int32
	_ = v689
	var v698 int32
	_ = v698
	var v700 int32
	_ = v700
	var v725 int32
	_ = v725
	var v735 int32
	_ = v735
	var v737 int32
	_ = v737
	var v771 int32
	_ = v771
	var v779 int32
	_ = v779
	var v781 int32
	_ = v781
	var v789 int32
	_ = v789
	var v793 int32
	_ = v793
	var v799 int32
	_ = v799
	var v805 int32
	_ = v805
	var v807 int32
	_ = v807
	var v810 int32
	_ = v810
	var v811 int32
	_ = v811
	var v813 int32
	_ = v813
	var v821 int32
	_ = v821
	var v823 int32
	_ = v823
	var v824 int32
	_ = v824
	var v840 int32
	_ = v840
	var v855 int32
	_ = v855
	var v862 int32
	_ = v862
	var v863 int32
	_ = v863
	var v866 int32
	_ = v866
	var v870 int32
	_ = v870
	var v872 int32
	_ = v872
	var v873 int32
	_ = v873
	var v876 int32
	_ = v876
	var v880 int32
	_ = v880
	var v883 int32
	_ = v883
	var v886 int32
	_ = v886
	var v890 int32
	_ = v890
	var v893 int32
	_ = v893
	var v896 int32
	_ = v896
	var v900 int32
	_ = v900
	var v903 int32
	_ = v903
	var v905 int32
	_ = v905
	var v918 int32
	_ = v918
	var v949 int32
	_ = v949
	var v963 int32
	_ = v963
	var v971 int32
	_ = v971
	var v972 int32
	_ = v972
	var v975 int32
	_ = v975
	var v979 int32
	_ = v979
	var v981 int32
	_ = v981
	var v984 int32
	_ = v984
	var v987 int32
	_ = v987
	var v1024 int32
	_ = v1024
	var v1025 int32
	_ = v1025
	var v1026 int32
	_ = v1026
	var v1029 int32
	_ = v1029
	var v1031 int32
	_ = v1031
	var v1034 int32
	_ = v1034
	var v1043 int32
	_ = v1043
	var v1044 int32
	_ = v1044
	var v1070 int32
	_ = v1070
	var v1071 int32
	_ = v1071
	var v1073 int32
	_ = v1073
	var v1084 int32
	_ = v1084
	var v1117 int32
	_ = v1117
	var v1141 int32
	_ = v1141
	var v1151 int32
	_ = v1151
	var v1186 int32
	_ = v1186
	var v1189 int32
	_ = v1189
	var v1197 int32
	_ = v1197
	var v1205 int32
	_ = v1205
	var v1208 int32
	_ = v1208
	var v1209 int32
	_ = v1209
	var v1211 int32
	_ = v1211
	var v1219 int32
	_ = v1219
	var v1220 int32
	_ = v1220
	var v1221 int32
	_ = v1221
	var v1224 int32
	_ = v1224
	var v1225 int32
	_ = v1225
	var v1227 int32
	_ = v1227
	var v1228 int32
	_ = v1228
	var v1230 int32
	_ = v1230
	var v1232 int32
	_ = v1232
	var v1235 int32
	_ = v1235
	var v1236 int32
	_ = v1236
	var v1237 int32
	_ = v1237
	var v1242 int32
	_ = v1242
	var v1243 int32
	_ = v1243
	var v1244 int32
	_ = v1244
	var v1247 int32
	_ = v1247
	var v1248 int32
	_ = v1248
	var v1249 int32
	_ = v1249
	var v1252 int32
	_ = v1252
	var v1253 int32
	_ = v1253
	var v1255 int32
	_ = v1255
	var v1260 int32
	_ = v1260
	var v1273 int32
	_ = v1273
	var v1278 int32
	_ = v1278
	var v1279 int32
	_ = v1279
	var v1280 int32
	_ = v1280
	var v1282 int32
	_ = v1282
	var v1289 int32
	_ = v1289
	var v1314 int32
	_ = v1314
	var v1315 int32
	_ = v1315
	var v1319 int32
	_ = v1319
	var v1327 int32
	_ = v1327
	var v1328 int32
	_ = v1328
	var v1329 int32
	_ = v1329
	var v1332 int32
	_ = v1332
	var v1333 int32
	_ = v1333
	var v1335 int32
	_ = v1335
	var v1336 int32
	_ = v1336
	var v1338 int32
	_ = v1338
	var v1340 int32
	_ = v1340
	var v1343 int32
	_ = v1343
	var v1344 int32
	_ = v1344
	var v1345 int32
	_ = v1345
	var v1350 int32
	_ = v1350
	var v1351 int32
	_ = v1351
	var v1352 int32
	_ = v1352
	var v1355 int32
	_ = v1355
	var v1356 int32
	_ = v1356
	var v1357 int32
	_ = v1357
	var v1360 int32
	_ = v1360
	var v1361 int32
	_ = v1361
	var v1363 int32
	_ = v1363
	var v1368 int32
	_ = v1368
	var v1381 int32
	_ = v1381
	var v1384 int32
	_ = v1384
	var v1392 int32
	_ = v1392
	var v1417 int32
	_ = v1417
	var v1419 int32
	_ = v1419
	var v1424 int32
	_ = v1424
	var v1429 int32
	_ = v1429
	var v1438 int32
	_ = v1438
	var v1441 int32
	_ = v1441
	var v1442 int32
	_ = v1442
	var v1475 int32
	_ = v1475
	var v1477 int32
	_ = v1477
	var v1511 int32
	_ = v1511
	var v1513 int32
	_ = v1513
	var v1522 int32
	_ = v1522
	var v1566 int32
	_ = v1566
	var v1584 int32
	_ = v1584
	var v1597 int32
	_ = v1597
	var v1609 int32
	_ = v1609
	var v1610 int32
	_ = v1610
	var v1619 int32
	_ = v1619
	var v1644 int32
	_ = v1644
	var v1652 int32
	_ = v1652
	v8 = int32(0)
	v32 = m.G0
	v34 = v32 - int32(128)
	m.G0 = v34
	F_DeconstructQualifiedName(m, l0, v34+int32(8), v34+int32(4))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v34)+8))
	if v44 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	m.G0 = v34 + int32(128)
	return v1652
L4:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v34)+4))
	v53 = int32(0)
	v55 = F_SearchSysCacheList(m, int32(46), int32(1), v52, v53, v53)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L1
	} else {
		goto L11
	}
L5:
	;
	v45 = F_LookupExplicitNamespace(m, v44, l6)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L1
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	F_recomputeNamespacePath(m)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L1
	} else {
		goto L10
	}
L8:
	;
	if v45 != 0 {
		v49 = v45
		goto L4
	} else {
		goto L9
	}
L9:
	;
	v1652 = v8
	goto L3
L10:
	;
	v49 = v8
	goto L4
L11:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v55)+40))
	if int32(0) < v57 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v69 = v8
	v82 = v8
	v91 = v8
	goto L15
L13:
	;
	v1619 = v8
	goto L14
L14:
	;
	F_ReleaseCatCacheList(m, v55)
	mBase = m.M
	v1644 = m.ExcPending
	if v1644 != 0 {
		goto L1
	} else {
		goto L246
	}
L15:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v55+int32(48)+v91<<(uint(int32(2))%32))))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v96)+56))
	v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97)+22)))
	v99 = v97 + v98
	v100 = int32(*(*int16)(unsafe.Add(mBase, uint32(v99)+104)))
	if v49 != 0 {
		goto L20
	} else {
		goto L21
	}
L16:
	;
	v1619 = v1584
	goto L14
L17:
	;
	v1609 = v91 + int32(1)
	v1610 = *(*int32)(unsafe.Add(mBase, uint32(v55)+40))
	if v1609 < v1610 {
		v69 = v1584
		v82 = v1597
		v91 = v1609
		goto L15
	} else {
		goto L245
	}
L18:
	;
	v1584 = v69
	v1597 = v1566
	goto L17
L19:
	;
	v193 = v99 + int32(136)
	v195 = v96 + int32(40)
	if l5 == int32(0) {
		v220 = v100
		v221 = v193
		goto L38
	} else {
		goto L39
	}
L20:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v99)+68))
	if v102 != v49 {
		v1566 = v82
		goto L18
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v105 = *(*int32)(unsafe.Add(mBase, _c_F_FuncnameGetCandidates[0]))
	if v105 == int32(0) {
		v1566 = v82
		goto L18
	} else {
		goto L24
	}
L23:
	;
	v161 = int32(0)
	goto L19
L24:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v105)+4))
	if v108 <= int32(0) {
		v1566 = v82
		goto L18
	} else {
		goto L25
	}
L25:
	;
	v111 = int32(0)
	if v111 < v108 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v114 = v108
	goto L28
L27:
	;
	v114 = v111
	goto L28
L28:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v105)+12))
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v99)+68))
	v119 = *(*int32)(unsafe.Add(mBase, _c_F_FuncnameGetCandidates[1]))
	v121 = int32(0)
	goto L29
L29:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v115+v121<<(uint(int32(2))%32))))
	if base.B2i32(v155 == v117)&base.B2i32(v117 != v119) != 0 {
		v161 = v121
		goto L19
	} else {
		goto L31
	}
L30:
	;
	v1566 = v82
	goto L18
L31:
	;
	v159 = v121 + int32(1)
	if v159 != v114 {
		v121 = v159
		goto L29
	} else {
		goto L32
	}
L32:
	;
	goto L30
L33:
	;
	if l1 < v220 {
		goto L126
	} else {
		goto L127
	}
L34:
	;
	if base.B2i32(v220 <= v239)|base.B2i32(v220 <= l1) == int32(0) {
		goto L116
	} else {
		goto L117
	}
L35:
	;
	v637 = l4 & base.B2i32(l1 < v220)
	if v637 == int32(0) {
		goto L111
	} else {
		goto L112
	}
L36:
	;
	v633 = int32(0)
	v634 = v609
	v635 = v82
	goto L35
L37:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v622 = m.ExcPending
	if v622 != 0 {
		goto L1
	} else {
		goto L107
	}
L38:
	;
	if l2 != 0 {
		goto L47
	} else {
		goto L48
	}
L39:
	;
	v202 = F_SysCacheGetAttr(m, int32(46), v195, int32(21), v34+int32(16))
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	v204 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+16)))
	if v204 != 0 {
		v220 = v100
		v221 = v193
		goto L38
	} else {
		goto L41
	}
L41:
	;
	v205 = F_pg_detoast_datum(m, v202)
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v205)+4))
	if v207 != int32(1) {
		goto L37
	} else {
		goto L43
	}
L43:
	;
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v205)+16))
	if v210 < int32(0) {
		goto L37
	} else {
		goto L44
	}
L44:
	;
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v205)+8))
	if v213 != 0 {
		goto L37
	} else {
		goto L45
	}
L45:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v205)+12))
	if v214 != int32(26) {
		goto L37
	} else {
		goto L46
	}
L46:
	;
	v220 = v210
	v221 = v205 + int32(24)
	goto L38
L47:
	;
	if l3 != 0 {
		goto L50
	} else {
		goto L51
	}
L48:
	;
	goto L49
L49:
	;
	v609 = int32(0)
	if base.B2i32(l3 == v609)|base.B2i32(l1 < v220) != 0 {
		goto L36
	} else {
		goto L106
	}
L50:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v99)+88))
	if v222 != 0 {
		v1566 = v82
		goto L18
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	v224 = l4 & base.B2i32(l1 < v220)
	if v224 != 0 {
		goto L55
	} else {
		goto L56
	}
L53:
	;
	goto L52
L54:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v96)+56))
	v230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v229)+22)))
	v231 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v236 = F_SysCacheGetAttr(m, int32(47), v195, int32(23), v34+int32(15))
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L1
	} else {
		goto L60
	}
L55:
	;
	v225 = int32(*(*int16)(unsafe.Add(mBase, uint32(v99)+106)))
	if v220 <= l1+v225 {
		goto L54
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	if l1 != v220 {
		v1566 = v82
		goto L18
	} else {
		goto L59
	}
L58:
	;
	v1566 = v82
	goto L18
L59:
	;
	goto L54
L60:
	;
	v238 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+15)))
	if v238 != 0 {
		v1566 = v82
		goto L18
	} else {
		goto L61
	}
L61:
	;
	v239 = l1 - v231
	v240 = int32(0)
	v247 = F_get_func_arg_info(m, v195, v34+int32(124), v34+int32(120), v34+int32(116))
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	v251 = F_palloc(m, v220<<(uint(int32(2))%32))
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	if v220 != 0 {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	base.MemoryFill(m, v34+int32(16), int32(0), v220)
	goto L66
L65:
	;
	goto L66
L66:
	;
	if v239 <= int32(0) {
		v434 = v240
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v458 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v458 <= int32(0) {
		v659 = v434
		goto L34
	} else {
		goto L82
	}
L68:
	;
	if v239 != 0 {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	base.MemoryFill(m, v34+int32(16), int32(1), v239)
	goto L71
L70:
	;
	goto L71
L71:
	;
	v264 = v239 & int32(7)
	v265 = int32(0)
	if base.Ui32(v231-l1) <= base.Ui32(int32(-8)) {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v278 = v265
	v279 = v240
	goto L75
L73:
	;
	v362 = v265
	goto L74
L74:
	;
	v393 = v362
	v403 = v265
	goto L79
L75:
	;
	v303 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v251+v278<<(uint(v303)%32)))) = v278
	v308 = v278 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v251+v308<<(uint(v303)%32)))) = v308
	v314 = v278 | v303
	*(*int32)(unsafe.Add(mBase, uint32(v251+v314<<(uint(v303)%32)))) = v314
	v320 = v278 | int32(3)
	*(*int32)(unsafe.Add(mBase, uint32(v251+v320<<(uint(v303)%32)))) = v320
	v326 = v278 | int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v251+v326<<(uint(v303)%32)))) = v326
	v332 = v278 | int32(5)
	*(*int32)(unsafe.Add(mBase, uint32(v251+v332<<(uint(v303)%32)))) = v332
	v338 = v278 | int32(6)
	*(*int32)(unsafe.Add(mBase, uint32(v251+v338<<(uint(v303)%32)))) = v338
	v344 = v278 | int32(7)
	*(*int32)(unsafe.Add(mBase, uint32(v251+v344<<(uint(v303)%32)))) = v344
	v349 = int32(8)
	v350 = v278 + v349
	v352 = v279 + v349
	if v352 != v239&int32(2147483640) {
		v278 = v350
		v279 = v352
		goto L75
	} else {
		goto L77
	}
L76:
	;
	if v264 == int32(0) {
		v434 = v239
		goto L67
	} else {
		goto L78
	}
L77:
	;
	goto L76
L78:
	;
	v362 = v350
	goto L74
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v251+v393<<(uint(int32(2))%32)))) = v393
	v422 = int32(1)
	v425 = v403 + v422
	if v425 != v264 {
		v393 = v393 + v422
		v403 = v425
		goto L79
	} else {
		goto L81
	}
L80:
	;
	v434 = v239
	goto L67
L81:
	;
	goto L80
L82:
	;
	v461 = int32(0)
	if v247 <= v461 {
		v1566 = v82
		goto L18
	} else {
		goto L83
	}
L83:
	;
	v471 = v434
	v479 = v461
	goto L84
L84:
	;
	v495 = *(*int32)(unsafe.Add(mBase, uint32(v34)+116))
	v496 = int32(0)
	v499 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v503 = *(*int32)(unsafe.Add(mBase, uint32(v499+v479<<(uint(int32(2))%32))))
	v504 = *(*int32)(unsafe.Add(mBase, uint32(v34)+120))
	v513 = v496
	v523 = v496
	goto L86
L85:
	;
	v1566 = v82
	goto L18
L86:
	;
	if l5|base.B2i32(v495 == v496) == int32(0) {
		goto L89
	} else {
		goto L90
	}
L87:
	;
	goto L85
L88:
	;
	v607 = v513 + int32(1)
	if v607 != v247 {
		v513 = v607
		v523 = v604
		goto L86
	} else {
		goto L105
	}
L89:
	;
	v541 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v513+v495))))
	v543 = v541 - int32(98)
	if base.B2i32(base.Ui32(int32(20)) < base.Ui32(v543))|base.B2i32(int32(1)<<(uint(v543)%32)&int32(_a_F_FuncnameGetCandidates_0) == int32(0)) != 0 {
		v604 = v523
		goto L88
	} else {
		goto L92
	}
L90:
	;
	goto L91
L91:
	;
	v557 = *(*int32)(unsafe.Add(mBase, uint32(v504+v513<<(uint(int32(2))%32))))
	if v557 == int32(0) {
		goto L93
	} else {
		goto L94
	}
L92:
	;
	goto L91
L93:
	;
	v604 = v523 + int32(1)
	goto L88
L94:
	;
	v562 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v557))))
	v565 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v503))))
	if base.B2i32(v562 == int32(0))|base.B2i32(v562 != v565) != 0 {
		v583 = v562
		v584 = v565
		goto L96
	} else {
		goto L97
	}
L95:
	;
	if v583-v584 != 0 {
		goto L93
	} else {
		goto L102
	}
L96:
	;
	goto L95
L97:
	;
	v568 = v557
	v569 = v503
	goto L98
L98:
	;
	v572 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v569)+1)))
	v573 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v568)+1)))
	if v573 == int32(0) {
		v583 = v573
		v584 = v572
		goto L96
	} else {
		goto L100
	}
L99:
	;
	v583 = v573
	v584 = v572
	goto L96
L100:
	;
	v576 = int32(1)
	if v573 == v572 {
		v568 = v568 + v576
		v569 = v569 + v576
		goto L98
	} else {
		goto L101
	}
L101:
	;
	goto L99
L102:
	;
	v588 = v34 + int32(16) + v523
	v589 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v588))))
	if v589 != 0 {
		v1566 = v82
		goto L18
	} else {
		goto L103
	}
L103:
	;
	v590 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v588))) = uint8(v590)
	*(*int32)(unsafe.Add(mBase, uint32(v251+v471<<(uint(int32(2))%32)))) = v523
	v597 = v471 + v590
	v599 = v479 + v590
	v600 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v599 < v600 {
		v471 = v597
		v479 = v599
		goto L84
	} else {
		goto L104
	}
L104:
	;
	v659 = v597
	goto L34
L105:
	;
	goto L87
L106:
	;
	v614 = *(*int32)(unsafe.Add(mBase, uint32(v99)+88))
	v616 = base.B2i32(v614 != int32(0))
	v633 = v614
	v634 = v616
	v635 = v82 | v616
	goto L35
L107:
	;
	F_errmsg_internal(m, int32(_a_F_FuncnameGetCandidates_1), int32(0))
	mBase = m.M
	v626 = m.ExcPending
	if v626 != 0 {
		goto L1
	} else {
		goto L108
	}
L108:
	;
	F_errfinish(m, int32(_a_F_FuncnameGetCandidates_2), int32(1289), int32(_a_F_FuncnameGetCandidates_3))
	mBase = m.M
	v631 = m.ExcPending
	if v631 != 0 {
		goto L1
	} else {
		goto L109
	}
L109:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L110:
	;
	v645 = int32(0)
	if base.B2i32(l1 == v220)|v634|v637|base.B2i32(l1 < v645) != 0 {
		v779 = v633
		v781 = v645
		v789 = v634
		v793 = v644
		v799 = v637
		goto L33
	} else {
		goto L115
	}
L111:
	;
	v644 = v635
	goto L110
L112:
	;
	goto L113
L113:
	;
	v641 = int32(*(*int16)(unsafe.Add(mBase, uint32(v99)+106)))
	if v220 <= l1+v641 {
		v644 = int32(1)
		goto L110
	} else {
		goto L114
	}
L114:
	;
	v1566 = v635
	goto L18
L115:
	;
	v1566 = v644
	goto L18
L116:
	;
	v689 = int32(*(*int16)(unsafe.Add(mBase, uint32(v229+v230)+106)))
	v698 = v659
	v700 = v239
	goto L119
L117:
	;
	goto L118
L118:
	;
	v771 = int32(0)
	v779 = v771
	v781 = v251
	v789 = v771
	v793 = int32(1)
	v799 = v224
	goto L33
L119:
	;
	v725 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34+int32(16)+v700))))
	if v725 == int32(0) {
		goto L121
	} else {
		goto L122
	}
L120:
	;
	goto L118
L121:
	;
	if v700 < v220-v689 {
		v1566 = v82
		goto L18
	} else {
		goto L124
	}
L122:
	;
	v735 = v698
	goto L123
L123:
	;
	v737 = v700 + int32(1)
	if v737 != v220 {
		v698 = v735
		v700 = v737
		goto L119
	} else {
		goto L125
	}
L124:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v251+v698<<(uint(int32(2))%32)))) = v700
	v735 = v698 + int32(1)
	goto L123
L125:
	;
	goto L120
L126:
	;
	v805 = v220
	goto L128
L127:
	;
	v805 = l1
	goto L128
L128:
	;
	v807 = v805 << (uint(int32(2)) % 32)
	v810 = F_palloc(m, v807+int32(32))
	mBase = m.M
	v811 = m.ExcPending
	if v811 != 0 {
		goto L1
	} else {
		goto L129
	}
L129:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v810)+4)) = v161
	v813 = *(*int32)(unsafe.Add(mBase, uint32(v99)))
	*(*int32)(unsafe.Add(mBase, uint32(v810)+28)) = v781
	*(*int32)(unsafe.Add(mBase, uint32(v810)+16)) = v805
	*(*int32)(unsafe.Add(mBase, uint32(v810)+12)) = v220
	*(*int32)(unsafe.Add(mBase, uint32(v810)+8)) = v813
	if v781 != 0 {
		goto L131
	} else {
		goto L132
	}
L130:
	;
	if v789 != 0 {
		goto L147
	} else {
		goto L148
	}
L131:
	;
	if v220 <= int32(0) {
		goto L130
	} else {
		goto L134
	}
L132:
	;
	goto L133
L133:
	;
	v987 = v220 << (uint(int32(2)) % 32)
	if v987 == int32(0) {
		goto L130
	} else {
		goto L145
	}
L134:
	;
	v821 = v220 & int32(3)
	v823 = v810 + int32(32)
	v824 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v220) {
		goto L135
	} else {
		goto L136
	}
L135:
	;
	v840 = v824
	v855 = int32(0)
	goto L138
L136:
	;
	v918 = v824
	goto L137
L137:
	;
	v949 = v918
	v963 = v824
	goto L142
L138:
	;
	v862 = int32(2)
	v863 = v840 << (uint(v862) % 32)
	v866 = *(*int32)(unsafe.Add(mBase, uint32(v781+v863)))
	v870 = *(*int32)(unsafe.Add(mBase, uint32(v221+v866<<(uint(v862)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v823+v863))) = v870
	v872 = int32(4)
	v873 = v863 | v872
	v876 = *(*int32)(unsafe.Add(mBase, uint32(v781+v873)))
	v880 = *(*int32)(unsafe.Add(mBase, uint32(v221+v876<<(uint(v862)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v823+v873))) = v880
	v883 = v863 | int32(8)
	v886 = *(*int32)(unsafe.Add(mBase, uint32(v781+v883)))
	v890 = *(*int32)(unsafe.Add(mBase, uint32(v221+v886<<(uint(v862)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v823+v883))) = v890
	v893 = v863 | int32(12)
	v896 = *(*int32)(unsafe.Add(mBase, uint32(v781+v893)))
	v900 = *(*int32)(unsafe.Add(mBase, uint32(v221+v896<<(uint(v862)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v823+v893))) = v900
	v903 = v840 + v872
	v905 = v855 + v872
	if v905 != v220&int32(2147483644) {
		v840 = v903
		v855 = v905
		goto L138
	} else {
		goto L140
	}
L139:
	;
	if v821 == int32(0) {
		goto L130
	} else {
		goto L141
	}
L140:
	;
	goto L139
L141:
	;
	v918 = v903
	goto L137
L142:
	;
	v971 = int32(2)
	v972 = v949 << (uint(v971) % 32)
	v975 = *(*int32)(unsafe.Add(mBase, uint32(v781+v972)))
	v979 = *(*int32)(unsafe.Add(mBase, uint32(v221+v975<<(uint(v971)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v823+v972))) = v979
	v981 = int32(1)
	v984 = v963 + v981
	if v984 != v821 {
		v949 = v949 + v981
		v963 = v984
		goto L142
	} else {
		goto L144
	}
L143:
	;
	goto L130
L144:
	;
	goto L143
L145:
	;
	base.MemoryCopy(m, v810+int32(32), v221, v987)
	goto L130
L146:
	;
	v1186 = int32(0)
	if v799 != 0 {
		goto L160
	} else {
		goto L161
	}
L147:
	;
	v1024 = v805 - v220
	v1025 = int32(1)
	v1026 = v1024 + v1025
	*(*int32)(unsafe.Add(mBase, uint32(v810)+20)) = v1026
	v1029 = v810 + int32(32)
	v1031 = v220 - v1025
	v1034 = v1026 & int32(7)
	if v1034 != 0 {
		goto L150
	} else {
		goto L151
	}
L148:
	;
	goto L149
L149:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v810)+20)) = int32(0)
	goto L146
L150:
	;
	v1043 = int32(0)
	v1044 = v1031
	goto L153
L151:
	;
	v1084 = v1031
	goto L152
L152:
	;
	if base.Ui32(v1024) < base.Ui32(int32(7)) {
		goto L146
	} else {
		goto L156
	}
L153:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1029+v1044<<(uint(int32(2))%32)))) = v779
	v1070 = int32(1)
	v1071 = v1044 + v1070
	v1073 = v1043 + v1070
	if v1073 != v1034 {
		v1043 = v1073
		v1044 = v1071
		goto L153
	} else {
		goto L155
	}
L154:
	;
	v1084 = v1071
	goto L152
L155:
	;
	goto L154
L156:
	;
	v1117 = v1084
	goto L157
L157:
	;
	v1141 = v1029 + v1117<<(uint(int32(2))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v1141))) = v779
	*(*int32)(unsafe.Add(mBase, uint32(v1141)+28)) = v779
	*(*int32)(unsafe.Add(mBase, uint32(v1141)+24)) = v779
	*(*int32)(unsafe.Add(mBase, uint32(v1141)+20)) = v779
	*(*int32)(unsafe.Add(mBase, uint32(v1141)+16)) = v779
	*(*int32)(unsafe.Add(mBase, uint32(v1141)+12)) = v779
	*(*int32)(unsafe.Add(mBase, uint32(v1141)+8)) = v779
	*(*int32)(unsafe.Add(mBase, uint32(v1141)+4)) = v779
	v1151 = v1117 + int32(8)
	if v1151 != v805 {
		v1117 = v1151
		goto L157
	} else {
		goto L159
	}
L158:
	;
	goto L146
L159:
	;
	goto L158
L160:
	;
	v1189 = v220 - l1
	goto L162
L161:
	;
	v1189 = v1186
	goto L162
L162:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v810)+24)) = v1189
	if v69 != 0 {
		goto L163
	} else {
		goto L164
	}
L163:
	;
	if (v793^int32(1))&base.B2i32(v49 != int32(0)) != 0 {
		goto L166
	} else {
		goto L167
	}
L164:
	;
	v1522 = v1186
	goto L165
L165:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v810))) = v1522
	v1584 = v810
	v1597 = v793
	goto L17
L166:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v810))) = v69
	v1584 = v810
	v1597 = v793
	goto L17
L167:
	;
	goto L168
L168:
	;
	v1197 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+37)))
	if (v1197^int32(-1)|v793)&int32(1) == int32(0) {
		goto L170
	} else {
		goto L171
	}
L169:
	;
	v1417 = *(*int32)(unsafe.Add(mBase, uint32(v1392)+4))
	if v1417 == v161 {
		goto L225
	} else {
		goto L226
	}
L170:
	;
	v1205 = *(*int32)(unsafe.Add(mBase, uint32(v69)+16))
	if v1205 != v805 {
		goto L173
	} else {
		goto L174
	}
L171:
	;
	goto L172
L172:
	;
	v1278 = v810 + int32(32)
	v1279 = *(*int32)(unsafe.Add(mBase, uint32(v810)+16))
	v1280 = v1279 - v1189
	v1282 = v1280 << (uint(int32(2)) % 32)
	v1289 = v69
	goto L195
L173:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v810))) = v69
	v1584 = v810
	v1597 = v793
	goto L17
L174:
	;
	goto L175
L175:
	;
	v1208 = int32(32)
	v1209 = v810 + v1208
	v1211 = v69 + v1208
	if base.Ui32(int32(4)) <= base.Ui32(v807) {
		goto L179
	} else {
		goto L180
	}
L176:
	;
	if v1273 == int32(0) {
		v1392 = v69
		goto L169
	} else {
		goto L194
	}
L177:
	;
	v1273 = int32(0)
	goto L176
L178:
	;
	v1247 = v1242
	v1248 = v1243
	v1249 = v1244
	goto L188
L179:
	;
	if (v1209|v1211)&int32(3) != 0 {
		v1242 = v1209
		v1243 = v1211
		v1244 = v807
		goto L178
	} else {
		goto L182
	}
L180:
	;
	v1235 = v1209
	v1236 = v1211
	v1237 = v807
	goto L181
L181:
	;
	if v1237 == int32(0) {
		goto L177
	} else {
		goto L187
	}
L182:
	;
	v1219 = v1209
	v1220 = v1211
	v1221 = v807
	goto L183
L183:
	;
	v1224 = *(*int32)(unsafe.Add(mBase, uint32(v1219)))
	v1225 = *(*int32)(unsafe.Add(mBase, uint32(v1220)))
	if v1224 != v1225 {
		v1242 = v1219
		v1243 = v1220
		v1244 = v1221
		goto L178
	} else {
		goto L185
	}
L184:
	;
	v1235 = v1230
	v1236 = v1228
	v1237 = v1232
	goto L181
L185:
	;
	v1227 = int32(4)
	v1228 = v1220 + v1227
	v1230 = v1219 + v1227
	v1232 = v1221 - v1227
	if base.Ui32(int32(3)) < base.Ui32(v1232) {
		v1219 = v1230
		v1220 = v1228
		v1221 = v1232
		goto L183
	} else {
		goto L186
	}
L186:
	;
	goto L184
L187:
	;
	v1242 = v1235
	v1243 = v1236
	v1244 = v1237
	goto L178
L188:
	;
	v1252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1247))))
	v1253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1248))))
	if v1252 == v1253 {
		goto L190
	} else {
		goto L191
	}
L189:
	;
	v1273 = v1252 - v1253
	goto L176
L190:
	;
	v1255 = int32(1)
	v1260 = v1249 - v1255
	if v1260 != 0 {
		v1247 = v1247 + v1255
		v1248 = v1248 + v1255
		v1249 = v1260
		goto L188
	} else {
		goto L193
	}
L191:
	;
	goto L192
L192:
	;
	goto L189
L193:
	;
	goto L177
L194:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v810))) = v69
	v1584 = v810
	v1597 = v793
	goto L17
L195:
	;
	v1314 = *(*int32)(unsafe.Add(mBase, uint32(v1289)+16))
	v1315 = *(*int32)(unsafe.Add(mBase, uint32(v1289)+24))
	if v1314-v1315 == v1280 {
		goto L197
	} else {
		goto L198
	}
L196:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v810))) = v69
	v1584 = v810
	v1597 = v793
	goto L17
L197:
	;
	v1319 = v1289 + int32(32)
	if base.Ui32(int32(4)) <= base.Ui32(v1282) {
		goto L203
	} else {
		goto L204
	}
L198:
	;
	goto L199
L199:
	;
	v1384 = *(*int32)(unsafe.Add(mBase, uint32(v1289)))
	if v1384 != 0 {
		v1289 = v1384
		goto L195
	} else {
		goto L219
	}
L200:
	;
	if v1381 == int32(0) {
		v1392 = v1289
		goto L169
	} else {
		goto L218
	}
L201:
	;
	v1381 = int32(0)
	goto L200
L202:
	;
	v1355 = v1350
	v1356 = v1351
	v1357 = v1352
	goto L212
L203:
	;
	if (v1278|v1319)&int32(3) != 0 {
		v1350 = v1278
		v1351 = v1319
		v1352 = v1282
		goto L202
	} else {
		goto L206
	}
L204:
	;
	v1343 = v1278
	v1344 = v1319
	v1345 = v1282
	goto L205
L205:
	;
	if v1345 == int32(0) {
		goto L201
	} else {
		goto L211
	}
L206:
	;
	v1327 = v1278
	v1328 = v1319
	v1329 = v1282
	goto L207
L207:
	;
	v1332 = *(*int32)(unsafe.Add(mBase, uint32(v1327)))
	v1333 = *(*int32)(unsafe.Add(mBase, uint32(v1328)))
	if v1332 != v1333 {
		v1350 = v1327
		v1351 = v1328
		v1352 = v1329
		goto L202
	} else {
		goto L209
	}
L208:
	;
	v1343 = v1338
	v1344 = v1336
	v1345 = v1340
	goto L205
L209:
	;
	v1335 = int32(4)
	v1336 = v1328 + v1335
	v1338 = v1327 + v1335
	v1340 = v1329 - v1335
	if base.Ui32(int32(3)) < base.Ui32(v1340) {
		v1327 = v1338
		v1328 = v1336
		v1329 = v1340
		goto L207
	} else {
		goto L210
	}
L210:
	;
	goto L208
L211:
	;
	v1350 = v1343
	v1351 = v1344
	v1352 = v1345
	goto L202
L212:
	;
	v1360 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1355))))
	v1361 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1356))))
	if v1360 == v1361 {
		goto L214
	} else {
		goto L215
	}
L213:
	;
	v1381 = v1360 - v1361
	goto L200
L214:
	;
	v1363 = int32(1)
	v1368 = v1357 - v1363
	if v1368 != 0 {
		v1355 = v1355 + v1363
		v1356 = v1356 + v1363
		v1357 = v1368
		goto L212
	} else {
		goto L217
	}
L215:
	;
	goto L216
L216:
	;
	goto L213
L217:
	;
	goto L201
L218:
	;
	goto L199
L219:
	;
	goto L196
L220:
	;
	if v69 == v1392 {
		goto L236
	} else {
		goto L237
	}
L221:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1392)+8)) = int32(0)
	F_pfree(m, v810)
	mBase = m.M
	v1438 = m.ExcPending
	if v1438 != 0 {
		goto L1
	} else {
		goto L234
	}
L222:
	;
	if int32(0) < v1419 {
		goto L220
	} else {
		goto L233
	}
L223:
	;
	if int32(0) <= v1424 {
		goto L221
	} else {
		goto L232
	}
L224:
	;
	F_pfree(m, v810)
	mBase = m.M
	v1429 = m.ExcPending
	if v1429 != 0 {
		goto L1
	} else {
		goto L231
	}
L225:
	;
	v1419 = *(*int32)(unsafe.Add(mBase, uint32(v1392)+20))
	if v789 == int32(0) {
		goto L222
	} else {
		goto L228
	}
L226:
	;
	goto L227
L227:
	;
	v1424 = v161 - v1417
	if v1424 <= int32(0) {
		goto L223
	} else {
		goto L230
	}
L228:
	;
	if v1419 == int32(0) {
		goto L224
	} else {
		goto L229
	}
L229:
	;
	goto L221
L230:
	;
	goto L224
L231:
	;
	v1566 = v793
	goto L18
L232:
	;
	goto L220
L233:
	;
	goto L221
L234:
	;
	v1566 = v793
	goto L18
L235:
	;
	F_pfree(m, v1392)
	mBase = m.M
	v1513 = m.ExcPending
	if v1513 != 0 {
		goto L1
	} else {
		goto L244
	}
L236:
	;
	v1441 = *(*int32)(unsafe.Add(mBase, uint32(v1392)))
	v1511 = v1441
	goto L235
L237:
	;
	goto L238
L238:
	;
	v1442 = v69
	goto L240
L239:
	;
	v1511 = v69
	goto L235
L240:
	;
	if v1442 == int32(0) {
		goto L239
	} else {
		goto L242
	}
L241:
	;
	v1477 = *(*int32)(unsafe.Add(mBase, uint32(v1392)))
	*(*int32)(unsafe.Add(mBase, uint32(v1442))) = v1477
	goto L239
L242:
	;
	v1475 = *(*int32)(unsafe.Add(mBase, uint32(v1442)))
	if v1392 != v1475 {
		v1442 = v1475
		goto L240
	} else {
		goto L243
	}
L243:
	;
	goto L241
L244:
	;
	v1522 = v1511
	goto L165
L245:
	;
	goto L16
L246:
	;
	v1652 = v1619
	goto L3
}
func F___fstatat(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v45 int32
	_ = v45
	v5 = int32(0)
	if base.B2i32(l3 != int32(_a_F___fstatat_0))|base.B2i32(l0 < v5) == v5 {
		v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
		if v13 != 0 {
			v33 = m.Env.X__syscall_newfstatat(m, l0, l1, l2, l3)
			mBase = m.M
			v37 = v33
		} else {
			v14 = m.Env.X__syscall_fstat64(m, l0, l2)
			mBase = m.M
			v37 = v14
		}
	} else {
		if l0 != int32(-100) {
			v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
			if base.B2i32(l3 == int32(0))&base.B2i32(v19 == int32(47)) != 0 {
				v31 = m.Env.X__syscall_stat64(m, l1, l2)
				mBase = m.M
				v37 = v31
			} else {
				if base.B2i32(l3 != int32(256))|base.B2i32(v19 != int32(47)) != 0 {
					v33 = m.Env.X__syscall_newfstatat(m, l0, l1, l2, l3)
					mBase = m.M
					v37 = v33
				} else {
					v35 = m.Env.X__syscall_lstat64(m, l1, l2)
					mBase = m.M
					v37 = v35
				}
			}
		} else {
			if l3 == int32(256) {
				v35 = m.Env.X__syscall_lstat64(m, l1, l2)
				mBase = m.M
				v37 = v35
			} else {
				if l3 != 0 {
					v33 = m.Env.X__syscall_newfstatat(m, l0, l1, l2, l3)
					mBase = m.M
					v37 = v33
				} else {
					v31 = m.Env.X__syscall_stat64(m, l1, l2)
					mBase = m.M
					v37 = v31
				}
			}
		}
	}
	if base.Ui32(int32(-4095)) <= base.Ui32(v37) {
		*(*int32)(unsafe.Add(mBase, _c_F___fstatat[0])) = int32(0) - v37
		v45 = int32(-1)
	} else {
		v45 = v37
	}
	return v45
}
func F_fastgetattr_5(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v6 = Fn13877(m, l0, l1, l2, l3, int32(_a_F_fastgetattr_5_0))
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return v6
	}
}
func F_fdw_handler_out(m *base.Module, l0 int32) int32 {
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v7 = Fn13854(m, l0, int32(_a_F_fdw_handler_out_0), int32(369), int32(_a_F_fdw_handler_out_1), int32(_a_F_fdw_handler_out_2), int32(_a_F_fdw_handler_out_3))
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
	}
}
func F_finalize_primnode(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	v3 = int32(0)
	if l0 == v3 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch v10 - int32(8) {
	case 0:
		goto L6
	case 1:
		goto L5
	default:
		goto L4
	case 15:
		goto L3
	}
L3:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v100)+8))
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v101)+8))
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v102)+12))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v103+v104<<(uint(int32(2))%32)-int32(4))))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v112 = F_finalize_primnode(m, v111, l1)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L8
	} else {
		goto L28
	}
L4:
	;
	v97 = F_expression_tree_walker_impl(m, l0, int32(848), l1)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L8
	} else {
		goto L27
	}
L5:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+276))
	if v26 == int32(0) {
		v70 = v3
		goto L11
	} else {
		goto L12
	}
L6:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v13 != int32(1) {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v18 = F_bms_add_member(m, v16, v17)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	return int32(0)
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v18
	return int32(0)
L10:
	;
	if v81 == int32(0) {
		goto L4
	} else {
		goto L25
	}
L11:
	;
	v81 = v70
	goto L10
L12:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v29 == int32(0) {
		v70 = v3
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
	if v32 != int32(1) {
		v70 = v3
		goto L11
	} else {
		goto L14
	}
L14:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	if v35 <= int32(0) {
		v70 = v3
		goto L11
	} else {
		goto L15
	}
L15:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v29)+12))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
	v43 = int32(0)
	v46 = v35
	goto L17
L16:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v52)+32))
	v70 = v66
	goto L11
L17:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v26)+12))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v48+v43<<(uint(int32(2))%32))))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)+4))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v53 == v54 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v81 = int32(0)
	goto L10
L19:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v52)+12))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
	v58 = F_equal(m, v56, v57)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L8
	} else {
		goto L22
	}
L20:
	;
	v61 = v46
	goto L21
L21:
	;
	v63 = v43 + int32(1)
	if v63 < v61 {
		v43 = v63
		v46 = v61
		goto L17
	} else {
		goto L24
	}
L22:
	;
	if v58 != 0 {
		goto L16
	} else {
		goto L23
	}
L23:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	v61 = v60
	goto L21
L24:
	;
	goto L18
L25:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v81)+8))
	v86 = F_bms_add_member(m, v84, v85)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L8
	} else {
		goto L26
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v86
	goto L4
L27:
	;
	return v97
L28:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v114 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v148 = F_finalize_primnode(m, v147, l1)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L8
	} else {
		goto L36
	}
L30:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v114)+4))
	if v117 <= int32(0) {
		goto L29
	} else {
		goto L31
	}
L31:
	;
	v122 = v3
	goto L32
L32:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v114)+12))
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v128+v122<<(uint(int32(2))%32))))
	v133 = F_bms_del_member(m, v127, v132)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L8
	} else {
		goto L34
	}
L33:
	;
	goto L29
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v133
	v137 = v122 + int32(1)
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v114)+4))
	if v137 < v138 {
		v122 = v137
		goto L32
	} else {
		goto L35
	}
L35:
	;
	goto L33
L36:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v110)+64))
	v151 = F_bms_copy(m, v150)
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L8
	} else {
		goto L37
	}
L37:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v153 == int32(0) {
		v181 = v151
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v186 = F_bms_join(m, v185, v181)
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L8
	} else {
		goto L45
	}
L39:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v153)+4))
	if v156 <= int32(0) {
		v181 = v151
		goto L38
	} else {
		goto L40
	}
L40:
	;
	v162 = int32(0)
	v163 = v151
	goto L41
L41:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v153)+12))
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v167+v162<<(uint(int32(2))%32))))
	v172 = F_bms_del_member(m, v163, v171)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L8
	} else {
		goto L43
	}
L42:
	;
	v181 = v172
	goto L38
L43:
	;
	v175 = v162 + int32(1)
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v153)+4))
	if v175 < v176 {
		v162 = v175
		v163 = v172
		goto L41
	} else {
		goto L44
	}
L44:
	;
	goto L42
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v186
	goto L1
}
func F_find_among_b(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v160 int32
	_ = v160
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	v4 = int32(0)
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v27 = l2
	v29 = v4
	v30 = v4
	v33 = v4
	v36 = v4
	goto L1
L1:
	;
	v44 = v27
	v47 = v30
	v50 = v33
	v53 = v36
	goto L3
L2:
	;
	v160 = v131
	goto L28
L3:
	;
	v62 = (v44-v47)>>(uint(int32(1))%32) + v47
	v65 = l1 + v62*int32(20)
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
	if v50 < v53 {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	v149 = base.B2i32(v128 != v131) & base.B2i32(v131 <= int32(0))
	if v149 != 0 {
		goto L20
	} else {
		goto L21
	}
L5:
	;
	if int32(1) < v128-v131 {
		v44 = v128
		v47 = v131
		v50 = v134
		v53 = v137
		goto L3
	} else {
		goto L19
	}
L6:
	;
	v128 = v44
	v131 = v62
	v134 = v113
	v137 = v53
	goto L5
L7:
	;
	v68 = v50
	goto L9
L8:
	;
	v68 = v53
	goto L9
L9:
	;
	v71 = v66 + (v68 ^ int32(-1))
	if v71 < int32(0) {
		v113 = v68
		goto L6
	} else {
		goto L10
	}
L10:
	;
	v78 = v68
	v80 = v71
	goto L13
L11:
	;
	v128 = v62
	v131 = v47
	v134 = v50
	v137 = v108
	goto L5
L12:
	;
	if int32(0) <= v98 {
		v113 = v78
		goto L6
	} else {
		goto L18
	}
L13:
	;
	if v19 == v18-v78 {
		v108 = v18 - v19
		goto L11
	} else {
		goto L15
	}
L14:
	;
	v128 = v44
	v131 = v62
	v134 = v66
	v137 = v53
	goto L5
L15:
	;
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21+v18-int32(1)-v78))))
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v65)+4))
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95+v80))))
	v98 = v94 - v97
	if v98 != 0 {
		goto L12
	} else {
		goto L16
	}
L16:
	;
	v99 = int32(1)
	if int32(0) < v80 {
		v78 = v78 + v99
		v80 = v80 - v99
		goto L13
	} else {
		goto L17
	}
L17:
	;
	goto L14
L18:
	;
	v108 = v78
	goto L11
L19:
	;
	goto L4
L20:
	;
	if v149 != 0 {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	goto L22
L22:
	;
	goto L2
L23:
	;
	v151 = int32(1)
	goto L25
L24:
	;
	v151 = v29
	goto L25
L25:
	;
	if v29 == int32(0) {
		v27 = v128
		v29 = v151
		v30 = v131
		v33 = v134
		v36 = v137
		goto L1
	} else {
		goto L26
	}
L26:
	;
	goto L22
L27:
	;
	return v197
L28:
	;
	v174 = l1 + v160*int32(20)
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v174)))
	if v175 <= v134 {
		goto L31
	} else {
		goto L32
	}
L29:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v174)+12))
	v197 = v194
	goto L27
L30:
	;
	goto L29
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v18 - v175
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v174)+16))
	if v179 == int32(0) {
		goto L30
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	v190 = int32(0)
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v174)+8))
	if v190 <= v191 {
		v160 = v191
		goto L28
	} else {
		goto L38
	}
L34:
	;
	v182 = m.T0[v179].(func(*base.Module, int32) int32)(m, l0)
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	return int32(0)
L36:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v174)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v18 - v186
	if v182 != 0 {
		goto L30
	} else {
		goto L37
	}
L37:
	;
	goto L33
L38:
	;
	v197 = v190
	goto L27
}
func F_find_coercion_pathway(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v105 int32
	_ = v105
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	v5 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v5
	if l1 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v14 = F_getBaseType(m, l1)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v18 = v5
	goto L3
L3:
	;
	if l0 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	return int32(0)
L5:
	;
	v18 = v14
	goto L3
L6:
	;
	v19 = F_getBaseType(m, l0)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L4
	} else {
		goto L9
	}
L7:
	;
	v21 = v5
	goto L8
L8:
	;
	if v18 != v21 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v21 = v19
	goto L8
L10:
	;
	v25 = F_SearchSysCache2(m, int32(12), v18, v21)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L4
	} else {
		goto L14
	}
L11:
	;
	v142 = int32(2)
	goto L12
L12:
	;
	m.G0 = v10 + int32(32)
	return v142
L13:
	;
	if v133 != 0 {
		goto L54
	} else {
		goto L55
	}
L14:
	;
	if v25 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v25)+16))
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+22)))
	v30 = v28 + v29
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+16)))
	switch v31 - int32(97) {
	case 0:
		v49 = int32(1)
		goto L19
	default:
		goto L21
	case 4:
		goto L20
	case 8:
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	if v21&int32(-9) == int32(22) {
		goto L36
	} else {
		goto L37
	}
L18:
	;
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+17)))
	switch v56 - int32(98) {
	case 0:
		goto L30
	default:
		goto L29
	case 4:
		goto L28
	case 7:
		v81 = int32(4)
		goto L27
	}
L19:
	;
	if base.Ui32(v49) <= base.Ui32(l2) {
		goto L18
	} else {
		goto L25
	}
L20:
	;
	v49 = int32(3)
	goto L19
L21:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L4
	} else {
		goto L22
	}
L22:
	;
	v38 = int32(*(*int8)(unsafe.Add(mBase, uint32(v30)+16)))
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v38
	F_errmsg_internal(m, int32(_a_F_find_coercion_pathway_3), v10)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L4
	} else {
		goto L23
	}
L23:
	;
	F_errfinish(m, int32(_a_F_find_coercion_pathway_1), int32(3196), int32(_a_F_find_coercion_pathway_2))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L4
	} else {
		goto L24
	}
L24:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L25:
	;
	F_ReleaseCatCache(m, v25)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L4
	} else {
		goto L26
	}
L26:
	;
	v133 = int32(0)
	goto L13
L27:
	;
	F_ReleaseCatCache(m, v25)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L4
	} else {
		goto L35
	}
L28:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v30)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v78
	v81 = int32(1)
	goto L27
L29:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L4
	} else {
		goto L32
	}
L30:
	;
	F_ReleaseCatCache(m, v25)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L4
	} else {
		goto L31
	}
L31:
	;
	v133 = int32(2)
	goto L13
L32:
	;
	v66 = int32(*(*int8)(unsafe.Add(mBase, uint32(v30)+17)))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v66
	F_errmsg_internal(m, int32(_a_F_find_coercion_pathway_0), v10+int32(16))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L4
	} else {
		goto L33
	}
L33:
	;
	F_errfinish(m, int32(_a_F_find_coercion_pathway_1), int32(3218), int32(_a_F_find_coercion_pathway_2))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L4
	} else {
		goto L34
	}
L34:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L35:
	;
	v133 = v81
	goto L13
L36:
	;
	v105 = int32(0)
	if l2 == v105 {
		v133 = v105
		goto L13
	} else {
		goto L44
	}
L37:
	;
	v88 = F_get_element_type(m, v21)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L4
	} else {
		goto L38
	}
L38:
	;
	if v88 == int32(0) {
		goto L36
	} else {
		goto L39
	}
L39:
	;
	v92 = F_get_element_type(m, v18)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L4
	} else {
		goto L40
	}
L40:
	;
	if v92 == int32(0) {
		goto L36
	} else {
		goto L41
	}
L41:
	;
	v98 = F_find_coercion_pathway(m, v88, v92, l2, v10+int32(24))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L4
	} else {
		goto L42
	}
L42:
	;
	if v98 == int32(0) {
		goto L36
	} else {
		goto L43
	}
L43:
	;
	v133 = int32(3)
	goto L13
L44:
	;
	F_get_type_category_preferred(m, v21, v10+int32(24), v10+int32(31))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L4
	} else {
		goto L45
	}
L45:
	;
	v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+24)))
	if v114 == int32(83) {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v133 = int32(4)
	goto L13
L47:
	;
	goto L48
L48:
	;
	if base.Ui32(l2) < base.Ui32(int32(3)) {
		v133 = v105
		goto L13
	} else {
		goto L49
	}
L49:
	;
	F_get_type_category_preferred(m, v18, v10+int32(24), v10+int32(31))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L4
	} else {
		goto L50
	}
L50:
	;
	v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+24)))
	if v128 == int32(83) {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v131 = int32(4)
	goto L53
L52:
	;
	v131 = int32(0)
	goto L53
L53:
	;
	v133 = v131
	goto L13
L54:
	;
	v137 = v133
	goto L56
L55:
	;
	v137 = int32(4)
	goto L56
L56:
	;
	if l2 == int32(2) {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v140 = v137
	goto L59
L58:
	;
	v140 = v133
	goto L59
L59:
	;
	v142 = v140
	goto L12
}
func F_find_nonnullable_rels(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_find_nonnullable_rels_walker(m, l0, int32(1))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_find_option(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
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
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v94 int32
	_ = v94
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v184 int32
	_ = v184
	var v193 int32
	_ = v193
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v231 int32
	_ = v231
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v261 int32
	_ = v261
	var v265 int32
	_ = v265
	var v270 int32
	_ = v270
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v289 int32
	_ = v289
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v300 int32
	_ = v300
	var v304 int32
	_ = v304
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v328 int32
	_ = v328
	var v333 int32
	_ = v333
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v348 int32
	_ = v348
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	v5 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = l0
	v16 = *(*int32)(unsafe.Add(mBase, _c_F_find_option[0]))
	v21 = F_hash_search(m, v16, v12+int32(8), v5, v5)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L6
	} else {
		goto L7
	}
L1:
	;
	m.G0 = v12 + int32(16)
	return v352
L2:
	;
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v339)+4))
	v350 = F_find_option(m, v348, int32(0), l2, l3)
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L6
	} else {
		goto L118
	}
L3:
	;
	v146 = int32(0)
	v153 = v27
	goto L51
L4:
	;
	v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v141))))
	if v142 != 0 {
		goto L3
	} else {
		goto L49
	}
L5:
	;
	if base.Ui32((v28-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L12
	} else {
		goto L13
	}
L6:
	;
	return int32(0)
L7:
	;
	if v21 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27))))
	if v28 != 0 {
		goto L5
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	v352 = v30
	goto L1
L11:
	;
	v141 = int32(_a_F_find_option_0)
	goto L4
L12:
	;
	v39 = v28 | int32(32)
	goto L14
L13:
	;
	v39 = v28
	goto L14
L14:
	;
	if v39 != int32(115) {
		goto L3
	} else {
		goto L15
	}
L15:
	;
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+1)))
	if v43 == int32(0) {
		v141 = int32(_a_F_find_option_1)
		goto L4
	} else {
		goto L16
	}
L16:
	;
	if base.Ui32((v43-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v54 = v43 | int32(32)
	goto L19
L18:
	;
	v54 = v43
	goto L19
L19:
	;
	if v54 != int32(111) {
		goto L3
	} else {
		goto L20
	}
L20:
	;
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+2)))
	if v58 == int32(0) {
		v141 = int32(_a_F_find_option_2)
		goto L4
	} else {
		goto L21
	}
L21:
	;
	if base.Ui32((v58-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v69 = v58 | int32(32)
	goto L24
L23:
	;
	v69 = v58
	goto L24
L24:
	;
	if v69 != int32(114) {
		goto L3
	} else {
		goto L25
	}
L25:
	;
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+3)))
	if v73 == int32(0) {
		v141 = int32(_a_F_find_option_3)
		goto L4
	} else {
		goto L26
	}
L26:
	;
	if base.Ui32((v73-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v84 = v73 | int32(32)
	goto L29
L28:
	;
	v84 = v73
	goto L29
L29:
	;
	if v84 != int32(116) {
		goto L3
	} else {
		goto L30
	}
L30:
	;
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+4)))
	if v88 == int32(0) {
		v141 = int32(_a_F_find_option_4)
		goto L4
	} else {
		goto L31
	}
L31:
	;
	if v88 != int32(95) {
		goto L3
	} else {
		goto L32
	}
L32:
	;
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+5)))
	if v94 == int32(0) {
		v141 = int32(_a_F_find_option_5)
		goto L4
	} else {
		goto L33
	}
L33:
	;
	if base.Ui32((v94-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v105 = v94 | int32(32)
	goto L36
L35:
	;
	v105 = v94
	goto L36
L36:
	;
	if v105 != int32(109) {
		goto L3
	} else {
		goto L37
	}
L37:
	;
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+6)))
	if v109 == int32(0) {
		v141 = int32(_a_F_find_option_6)
		goto L4
	} else {
		goto L38
	}
L38:
	;
	if base.Ui32((v109-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v120 = v109 | int32(32)
	goto L41
L40:
	;
	v120 = v109
	goto L41
L41:
	;
	if v120 != int32(101) {
		goto L3
	} else {
		goto L42
	}
L42:
	;
	v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+7)))
	if v124 == int32(0) {
		v141 = int32(_a_F_find_option_7)
		goto L4
	} else {
		goto L43
	}
L43:
	;
	if base.Ui32((v124-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v135 = v124 | int32(32)
	goto L46
L45:
	;
	v135 = v124
	goto L46
L46:
	;
	if v135 != int32(109) {
		goto L3
	} else {
		goto L47
	}
L47:
	;
	v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+8)))
	if v138 != 0 {
		goto L3
	} else {
		goto L48
	}
L48:
	;
	v141 = int32(_a_F_find_option_8)
	goto L4
L49:
	;
	v339 = int32(_a_F_find_option_9)
	goto L2
L50:
	;
	v193 = int32(0)
	v200 = v27
	goto L66
L51:
	;
	v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v153))))
	if v155 != 0 {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	if v146 != int32(10) {
		goto L50
	} else {
		goto L64
	}
L53:
	;
	if v146 == int32(10) {
		goto L50
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	goto L52
L56:
	;
	v160 = int32(1)
	v164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146)+uint32(_c_F_find_option[1]))))
	if base.Ui32((v164-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v173 = v164 | int32(32)
	goto L59
L58:
	;
	v173 = v164
	goto L59
L59:
	;
	v174 = int32(255)
	if base.Ui32((v155-int32(65))&v174) < base.Ui32(int32(26)) {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v184 = v155 | int32(32)
	goto L62
L61:
	;
	v184 = v155
	goto L62
L62:
	;
	if v173&v174 == v184 {
		v146 = v146 + v160
		v153 = v153 + v160
		goto L51
	} else {
		goto L63
	}
L63:
	;
	goto L50
L64:
	;
	v339 = int32(_a_F_find_option_10)
	goto L2
L65:
	;
	if l1 != 0 {
		goto L82
	} else {
		goto L83
	}
L66:
	;
	v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v200))))
	if v202 != 0 {
		goto L68
	} else {
		goto L69
	}
L67:
	;
	if v193 != int32(14) {
		goto L65
	} else {
		goto L79
	}
L68:
	;
	if v193 == int32(14) {
		goto L65
	} else {
		goto L71
	}
L69:
	;
	goto L70
L70:
	;
	goto L67
L71:
	;
	v207 = int32(1)
	v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v193)+uint32(_c_F_find_option[2]))))
	if base.Ui32((v211-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v220 = v211 | int32(32)
	goto L74
L73:
	;
	v220 = v211
	goto L74
L74:
	;
	v221 = int32(255)
	if base.Ui32((v202-int32(65))&v221) < base.Ui32(int32(26)) {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v231 = v202 | int32(32)
	goto L77
L76:
	;
	v231 = v202
	goto L77
L77:
	;
	if v220&v221 == v231 {
		v193 = v193 + v207
		v200 = v200 + v207
		goto L66
	} else {
		goto L78
	}
L78:
	;
	goto L65
L79:
	;
	v339 = int32(_a_F_find_option_11)
	goto L2
L80:
	;
	F_pfree(m, v250)
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L6
	} else {
		goto L117
	}
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v293)+4)) = v250
	v352 = v250
	goto L1
L82:
	;
	v239 = F_assignable_custom_variable_name(m, v27, l2, l3)
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L6
	} else {
		goto L85
	}
L83:
	;
	goto L84
L84:
	;
	v315 = int32(0)
	if l2 != 0 {
		v352 = v315
		goto L1
	} else {
		goto L111
	}
L85:
	;
	if v239 == int32(0) {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v352 = int32(0)
	goto L1
L87:
	;
	goto L88
L88:
	;
	v244 = int32(0)
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
	v247 = *(*int32)(unsafe.Add(mBase, _c_F_find_option[3]))
	v250 = F_MemoryContextAllocExtended(m, v247, int32(124), int32(2))
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L6
	} else {
		goto L89
	}
L89:
	;
	if v250 == int32(0) {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v255 = F_errstart(m, l3, int32(0))
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L6
	} else {
		goto L93
	}
L91:
	;
	goto L92
L92:
	;
	base.MemoryFill(m, v250, int32(0), int32(124))
	v274 = F_guc_strdup(m, l3, v245)
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L6
	} else {
		goto L98
	}
L93:
	;
	if v255 == int32(0) {
		v352 = v244
		goto L1
	} else {
		goto L94
	}
L94:
	;
	F_errcode(m, int32(_a_F_find_option_12))
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L6
	} else {
		goto L95
	}
L95:
	;
	F_errmsg(m, int32(_a_F_find_option_13), int32(0))
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L6
	} else {
		goto L96
	}
L96:
	;
	F_errfinish(m, int32(_a_F_find_option_14), int32(647), int32(_a_F_find_option_15))
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L6
	} else {
		goto L97
	}
L97:
	;
	v352 = v244
	goto L1
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v250))) = v274
	if v274 == int32(0) {
		goto L80
	} else {
		goto L99
	}
L99:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v250)+20)) = int64(12884902532)
	*(*int32)(unsafe.Add(mBase, uint32(v250)+12)) = int32(_a_F_find_option_16)
	*(*int64)(unsafe.Add(mBase, uint32(v250)+4)) = int64(197568495622)
	*(*int32)(unsafe.Add(mBase, uint32(v250)+92)) = v250 + int32(120)
	v289 = *(*int32)(unsafe.Add(mBase, _c_F_find_option[0]))
	v293 = F_hash_search(m, v289, v250, int32(3), v12+int32(15))
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L6
	} else {
		goto L100
	}
L100:
	;
	if v293 != 0 {
		goto L81
	} else {
		goto L101
	}
L101:
	;
	v296 = F_errstart(m, l3, int32(0))
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L6
	} else {
		goto L102
	}
L102:
	;
	if v296 != 0 {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	F_errcode(m, int32(_a_F_find_option_12))
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L6
	} else {
		goto L106
	}
L104:
	;
	goto L105
L105:
	;
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v250)))
	if v310 == int32(0) {
		goto L80
	} else {
		goto L109
	}
L106:
	;
	F_errmsg(m, int32(_a_F_find_option_13), int32(0))
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L6
	} else {
		goto L107
	}
L107:
	;
	F_errfinish(m, int32(_a_F_find_option_14), int32(1060), int32(_a_F_find_option_17))
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L6
	} else {
		goto L108
	}
L108:
	;
	goto L105
L109:
	;
	F_pfree(m, v310)
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L6
	} else {
		goto L110
	}
L110:
	;
	goto L80
L111:
	;
	v317 = F_errstart(m, l3, int32(0))
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L6
	} else {
		goto L112
	}
L112:
	;
	if v317 == int32(0) {
		v352 = v315
		goto L1
	} else {
		goto L113
	}
L113:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L6
	} else {
		goto L114
	}
L114:
	;
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v324
	F_errmsg(m, int32(_a_F_find_option_18), v12)
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L6
	} else {
		goto L115
	}
L115:
	;
	F_errfinish(m, int32(_a_F_find_option_14), int32(1279), int32(_a_F_find_option_19))
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L6
	} else {
		goto L116
	}
L116:
	;
	v352 = v315
	goto L1
L117:
	;
	v352 = int32(0)
	goto L1
L118:
	;
	v352 = v350
	goto L1
}
func F_find_typed_table_dependencies(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	v8 = m.G0
	v10 = v8 + int32(-64)
	m.G0 = v10
	v14 = F_table_open(m, int32(1259), int32(1))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v19 = v8 + int32(-48)
	F_ScanKeyInit(m, v19, int32(5), int32(3), int32(184), l0)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v26 = F_table_beginscan_catalog(m, v14, int32(1), v19)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L6
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L1
	} else {
		goto L19
	}
L5:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)+188))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)+12))
	m.T0[v60].(func(*base.Module, int32))(m, v26)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L1
	} else {
		goto L17
	}
L6:
	;
	v28 = F_heap_getnext(m, v26)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	if v28 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v51 = int32(0)
	goto L5
L9:
	;
	goto L10
L10:
	;
	if l2 == int32(0) {
		goto L4
	} else {
		goto L11
	}
L11:
	;
	v36 = int32(0)
	v42 = v28
	goto L12
L12:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)+16))
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+22)))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v43+v44)))
	v47 = F_lappend_oid(m, v36, v46)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L1
	} else {
		goto L14
	}
L13:
	;
	v51 = v47
	goto L5
L14:
	;
	v49 = F_heap_getnext(m, v26)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	if v49 != 0 {
		v36 = v47
		v42 = v49
		goto L12
	} else {
		goto L16
	}
L16:
	;
	goto L13
L17:
	;
	F_relation_close(m, v14, int32(1))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	m.G0 = v10 - int32(-64)
	return v51
L19:
	;
	F_errcode(m, int32(16909442))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = l1
	F_errmsg(m, int32(_a_F_find_typed_table_dependencies_0), v10)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	F_errhint(m, int32(_a_F_find_typed_table_dependencies_1), int32(0))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	F_errfinish(m, int32(_a_F_find_typed_table_dependencies_2), int32(_a_F_find_typed_table_dependencies_3), int32(_a_F_find_typed_table_dependencies_4))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_fix_opfuncids(m *base.Module, l0 int32) {
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	v3 = F_fix_opfuncids_walker(m, l0, int32(0))
	v4 = m.ExcPending
	if v4 != 0 {
		return
	} else {
		return
	}
}
func F_flatCopyTargetEntry(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v10 int64
	_ = v10
	var v12 int64
	_ = v12
	var v14 int32
	_ = v14
	var v16 int64
	_ = v16
	v4 = F_palloc0(m, int32(28))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v4))) = int32(62)
		v10 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
		*(*int64)(unsafe.Add(mBase, uint32(v4)+8)) = v10
		v12 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
		*(*int64)(unsafe.Add(mBase, uint32(v4)+16)) = v12
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
		*(*int32)(unsafe.Add(mBase, uint32(v4)+24)) = v14
		v16 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
		*(*int64)(unsafe.Add(mBase, uint32(v4))) = v16
		return v4
	}
}
func F_float48le(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 float64
	_ = v5
	var v11 float32
	_ = v11
	var v12 float64
	_ = v12
	var v22 int32
	_ = v22
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v5 = *(*float64)(unsafe.Add(mBase, uint32(v4)))
	if base.Ui64(base.I64_reinterpret_f64(v5)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
		v11 = *(*float32)(unsafe.Add(mBase, uint32(l0)+20))
		v12 = base.F64_promote_f32(v11)
		v22 = base.F64_ge(v5, v12) & base.B2i32(base.Ui64(base.I64_reinterpret_f64(v12)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)))
	} else {
		v22 = int32(1)
	}
	return v22
}
func F_float48ne(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 float64
	_ = v6
	var v8 int64
	_ = v8
	var v9 int64
	_ = v9
	var v10 float32
	_ = v10
	var v11 float64
	_ = v11
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v6 = *(*float64)(unsafe.Add(mBase, uint32(v5)))
	v8 = int64(9223372036854775807)
	v9 = base.I64_reinterpret_f64(v6) & v8
	v10 = *(*float32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = base.F64_promote_f32(v10)
	if base.Ui64(int64(9218868437227405313)) <= base.Ui64(base.I64_reinterpret_f64(v11)&v8) {
		return base.B2i32(base.Ui64(v9) < base.Ui64(int64(9218868437227405313)))
	} else {
		return base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v9)) | base.F64_ne(v6, v11)
	}
}
func F_float4_dist(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 float32
	_ = v5
	var v6 float32
	_ = v6
	var v8 float32
	_ = v8
	var v9 float32
	_ = v9
	var v24 int32
	_ = v24
	v5 = *(*float32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = *(*float32)(unsafe.Add(mBase, uint32(l0)+28))
	v8 = base.F32_abs(base.F32_sub(v5, v6))
	v9 = math.Float32frombits(uint32(0x7f800000))
	if base.F32_ne(v8, v9)|base.F32_eq(base.F32_abs(v5), v9)|base.F32_eq(base.F32_abs(v6), v9) == int32(0) {
		F_float_overflow_error(m)
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return int32(0)
		} else {
			base.Wasm_trap_unreachable()
			for {
			}
		}
	} else {
		return base.I32_reinterpret_f32(v8)
	}
}
func F_float4_to_char(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 float32
	_ = v59
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v77 float32
	_ = v77
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v150 float64
	_ = v150
	var v154 float32
	_ = v154
	var v155 float64
	_ = v155
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
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
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v230 int32
	_ = v230
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v250 int32
	_ = v250
	v2 = int32(0)
	v12 = m.G0
	v14 = v12 - int32(96)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v18 = F_pg_detoast_datum_packed(m, v17)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		return int32(0)
	} else {
		v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
		if v22 == int32(1) {
			v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+1)))
			if v28 == int32(18) {
				v31 = int32(16)
			} else {
				v31 = int32(0)
			}
			if base.Ui32((v28-int32(1))&int32(255)) < base.Ui32(int32(3)) {
				v38 = int32(4)
			} else {
				v38 = v31
			}
			v51 = v38
		} else {
			v39 = int32(1)
			if v22&v39 != 0 {
				v51 = int32(base.Ui32(v22)>>(uint(v39)%32)) - v39
			} else {
				v45 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
				v51 = int32(base.Ui32(v45)>>(uint(int32(2))%32)) - int32(4)
			}
		}
		if base.Ui32(int32(268435454)) <= base.Ui32(v51-int32(1)) {
			v57 = F_cstring_to_text(m, int32(_a_F_float4_to_char_0))
			mBase = m.M
			v58 = m.ExcPending
			if v58 != 0 {
				return int32(0)
			} else {
				v250 = v57
				m.G0 = v14 + int32(96)
				return v250
			}
		} else {
			v59 = base.F32_reinterpret_i32(v16)
			v64 = F_palloc0(m, v51<<(uint(int32(3))%32)|int32(5))
			mBase = m.M
			v65 = m.ExcPending
			if v65 != 0 {
				return int32(0)
			} else {
				v70 = F_NUM_cache(m, v51, v14+int32(60), v18, v14+int32(59))
				mBase = m.M
				v71 = m.ExcPending
				if v71 != 0 {
					return int32(0)
				} else {
					v72 = *(*int32)(unsafe.Add(mBase, uint32(v14)+72))
					if v72&int32(1024) != 0 {
						v76 = int32(2147483647)
						v77 = base.F32_nearest(v59)
						if base.Ui32(int32(2139095041)) <= base.Ui32(base.I32_reinterpret_f32(v77)&v76) {
							v84 = v76
						} else {
							v84 = base.I32_trunc_sat_f32_s(v77)
						}
						if base.F32_ge(v77, float32(-2.1474836e+09)) != 0 {
							v88 = v84
						} else {
							v88 = int32(2147483647)
						}
						if base.F32_lt(v77, float32(2.1474836e+09)) != 0 {
							v92 = v88
						} else {
							v92 = int32(2147483647)
						}
						v93 = F_int_to_roman(m, v92)
						mBase = m.M
						v94 = m.ExcPending
						if v94 != 0 {
							return int32(0)
						} else {
							v219 = v93
							v220 = int32(0)
							v224 = v2
							v230 = v64 + int32(4)
							F_NUM_processor(m, v70, v14+int32(60), v230, v219, int32(0), v220, v224, int32(1))
							mBase = m.M
							v234 = m.ExcPending
							if v234 != 0 {
								return int32(0)
							} else {
								v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+59)))
								if v235 == int32(1) {
									F_pfree(m, v70)
									mBase = m.M
									v239 = m.ExcPending
									if v239 != 0 {
										return int32(0)
									} else {
										v240 = F_strlen(m, v230)
										mBase = m.M
										*(*int32)(unsafe.Add(mBase, uint32(v64))) = v240<<(uint(int32(2))%32) + int32(16)
										v250 = v64
										m.G0 = v14 + int32(96)
										return v250
									}
								} else {
									v240 = F_strlen(m, v230)
									mBase = m.M
									*(*int32)(unsafe.Add(mBase, uint32(v64))) = v240<<(uint(int32(2))%32) + int32(16)
									v250 = v64
									m.G0 = v14 + int32(96)
									return v250
								}
							}
						}
					} else {
						if v72&int32(_a_F_float4_to_char_1) != 0 {
							if base.B2i32(base.Ui32(v16&int32(2147483647)) <= base.Ui32(int32(2139095040)))&base.F32_ne(base.F32_abs(v59), math.Float32frombits(uint32(0x7f800000))) == int32(0) {
								v107 = *(*int32)(unsafe.Add(mBase, uint32(v14)+60))
								v108 = *(*int32)(unsafe.Add(mBase, uint32(v14)+64))
								v109 = v107 + v108
								v112 = F_palloc(m, v109+int32(7))
								mBase = m.M
								v113 = m.ExcPending
								if v113 != 0 {
									return int32(0)
								} else {
									v115 = v109 + int32(6)
									if v115 != 0 {
										base.MemoryFill(m, v112, int32(35), v115)
									} else {
									}
									v118 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(v112+v115))) = uint8(v118)
									v122 = int32(32)
									*(*uint8)(unsafe.Add(mBase, uint32(v112))) = uint8(v122)
									v125 = int32(46)
									*(*uint8)(unsafe.Add(mBase, uint32(v112+v107)+1)) = uint8(v125)
									v219 = v112
									v220 = v118
									v224 = v2
									v230 = v64 + int32(4)
									F_NUM_processor(m, v70, v14+int32(60), v230, v219, int32(0), v220, v224, int32(1))
									mBase = m.M
									v234 = m.ExcPending
									if v234 != 0 {
										return int32(0)
									} else {
										v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+59)))
										if v235 == int32(1) {
											F_pfree(m, v70)
											mBase = m.M
											v239 = m.ExcPending
											if v239 != 0 {
												return int32(0)
											} else {
												v240 = F_strlen(m, v230)
												mBase = m.M
												*(*int32)(unsafe.Add(mBase, uint32(v64))) = v240<<(uint(int32(2))%32) + int32(16)
												v250 = v64
												m.G0 = v14 + int32(96)
												return v250
											}
										} else {
											v240 = F_strlen(m, v230)
											mBase = m.M
											*(*int32)(unsafe.Add(mBase, uint32(v64))) = v240<<(uint(int32(2))%32) + int32(16)
											v250 = v64
											m.G0 = v14 + int32(96)
											return v250
										}
									}
								}
							} else {
								v127 = *(*int32)(unsafe.Add(mBase, uint32(v14)+64))
								*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v127
								*(*float64)(unsafe.Add(mBase, uint32(v14)+40)) = base.F64_promote_f32(v59)
								v131 = int32(0)
								v135 = F_psprintf(m, int32(_a_F_float4_to_char_2), v14+int32(32))
								mBase = m.M
								v136 = m.ExcPending
								if v136 != 0 {
									return int32(0)
								} else {
									v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135))))
									if v137 != int32(43) {
										v219 = v135
										v220 = v131
										v224 = v2
									} else {
										v140 = int32(32)
										*(*uint8)(unsafe.Add(mBase, uint32(v135))) = uint8(v140)
										v219 = v135
										v220 = v131
										v224 = v2
									}
									v230 = v64 + int32(4)
									F_NUM_processor(m, v70, v14+int32(60), v230, v219, int32(0), v220, v224, int32(1))
									mBase = m.M
									v234 = m.ExcPending
									if v234 != 0 {
										return int32(0)
									} else {
										v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+59)))
										if v235 == int32(1) {
											F_pfree(m, v70)
											mBase = m.M
											v239 = m.ExcPending
											if v239 != 0 {
												return int32(0)
											} else {
												v240 = F_strlen(m, v230)
												mBase = m.M
												*(*int32)(unsafe.Add(mBase, uint32(v64))) = v240<<(uint(int32(2))%32) + int32(16)
												v250 = v64
												m.G0 = v14 + int32(96)
												return v250
											}
										} else {
											v240 = F_strlen(m, v230)
											mBase = m.M
											*(*int32)(unsafe.Add(mBase, uint32(v64))) = v240<<(uint(int32(2))%32) + int32(16)
											v250 = v64
											m.G0 = v14 + int32(96)
											return v250
										}
									}
								}
							}
						} else {
							if v72&int32(2048) != 0 {
								v144 = *(*int32)(unsafe.Add(mBase, uint32(v14)+80))
								v145 = *(*int32)(unsafe.Add(mBase, uint32(v14)+60))
								*(*int32)(unsafe.Add(mBase, uint32(v14)+60)) = v144 + v145
								v150 = F_pow(m, float64(10), base.F64_convert_i32_s(v144))
								mBase = m.M
								v154 = base.F32_mul(v59, base.F32_demote_f64(v150))
							} else {
								v154 = v59
							}
							v155 = base.F64_promote_f32(v154)
							*(*float64)(unsafe.Add(mBase, uint32(v14)+16)) = base.F64_abs(v155)
							v161 = F_psprintf(m, int32(_a_F_float4_to_char_3), v14+int32(16))
							mBase = m.M
							v162 = m.ExcPending
							if v162 != 0 {
								return int32(0)
							} else {
								v163 = F_strlen(m, v161)
								mBase = m.M
								if v163 <= int32(5) {
									v166 = *(*int32)(unsafe.Add(mBase, uint32(v14)+64))
									if v166+v163 < int32(7) {
										v174 = v166
									} else {
										v172 = int32(6) - v163
										*(*int32)(unsafe.Add(mBase, uint32(v14)+64)) = v172
										v174 = v172
									}
								} else {
									v172 = v2
									*(*int32)(unsafe.Add(mBase, uint32(v14)+64)) = v172
									v174 = v172
								}
								*(*float64)(unsafe.Add(mBase, uint32(v14)+8)) = v155
								*(*int32)(unsafe.Add(mBase, uint32(v14))) = v174
								v178 = F_psprintf(m, int32(_a_F_float4_to_char_4), v14)
								mBase = m.M
								v179 = m.ExcPending
								if v179 != 0 {
									return int32(0)
								} else {
									v180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178))))
									v182 = base.B2i32(v180 == int32(45))
									v183 = v178 + v182
									v184 = int32(46)
									v185 = F___strchrnul(m, v183, v184)
									mBase = m.M
									v187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v185))))
									if v187 == v184 {
										v191 = v185
									} else {
										v191 = int32(0)
									}
									if v191 != 0 {
										v194 = v191 - v183
									} else {
										v193 = F_strlen(m, v183)
										mBase = m.M
										v194 = v193
									}
									if v180 == int32(45) {
										v197 = int32(45)
									} else {
										v197 = int32(43)
									}
									v198 = *(*int32)(unsafe.Add(mBase, uint32(v14)+60))
									if v194 < v198 {
										v219 = v183
										v220 = v198 - v194
										v224 = v197
										v230 = v64 + int32(4)
										F_NUM_processor(m, v70, v14+int32(60), v230, v219, int32(0), v220, v224, int32(1))
										mBase = m.M
										v234 = m.ExcPending
										if v234 != 0 {
											return int32(0)
										} else {
											v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+59)))
											if v235 == int32(1) {
												F_pfree(m, v70)
												mBase = m.M
												v239 = m.ExcPending
												if v239 != 0 {
													return int32(0)
												} else {
													v240 = F_strlen(m, v230)
													mBase = m.M
													*(*int32)(unsafe.Add(mBase, uint32(v64))) = v240<<(uint(int32(2))%32) + int32(16)
													v250 = v64
													m.G0 = v14 + int32(96)
													return v250
												}
											} else {
												v240 = F_strlen(m, v230)
												mBase = m.M
												*(*int32)(unsafe.Add(mBase, uint32(v64))) = v240<<(uint(int32(2))%32) + int32(16)
												v250 = v64
												m.G0 = v14 + int32(96)
												return v250
											}
										}
									} else {
										if v194 <= v198 {
											v219 = v183
											v220 = int32(0)
											v224 = v197
											v230 = v64 + int32(4)
											F_NUM_processor(m, v70, v14+int32(60), v230, v219, int32(0), v220, v224, int32(1))
											mBase = m.M
											v234 = m.ExcPending
											if v234 != 0 {
												return int32(0)
											} else {
												v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+59)))
												if v235 == int32(1) {
													F_pfree(m, v70)
													mBase = m.M
													v239 = m.ExcPending
													if v239 != 0 {
														return int32(0)
													} else {
														v240 = F_strlen(m, v230)
														mBase = m.M
														*(*int32)(unsafe.Add(mBase, uint32(v64))) = v240<<(uint(int32(2))%32) + int32(16)
														v250 = v64
														m.G0 = v14 + int32(96)
														return v250
													}
												} else {
													v240 = F_strlen(m, v230)
													mBase = m.M
													*(*int32)(unsafe.Add(mBase, uint32(v64))) = v240<<(uint(int32(2))%32) + int32(16)
													v250 = v64
													m.G0 = v14 + int32(96)
													return v250
												}
											}
										} else {
											v203 = v174 + v198
											v206 = F_palloc(m, v203+int32(2))
											mBase = m.M
											v207 = m.ExcPending
											if v207 != 0 {
												return int32(0)
											} else {
												v209 = v203 + int32(1)
												if v209 != 0 {
													base.MemoryFill(m, v206, int32(35), v209)
												} else {
												}
												v212 = int32(0)
												*(*uint8)(unsafe.Add(mBase, uint32(v206+v209))) = uint8(v212)
												v217 = int32(46)
												*(*uint8)(unsafe.Add(mBase, uint32(v206+v198))) = uint8(v217)
												v219 = v206
												v220 = v212
												v224 = v197
												v230 = v64 + int32(4)
												F_NUM_processor(m, v70, v14+int32(60), v230, v219, int32(0), v220, v224, int32(1))
												mBase = m.M
												v234 = m.ExcPending
												if v234 != 0 {
													return int32(0)
												} else {
													v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+59)))
													if v235 == int32(1) {
														F_pfree(m, v70)
														mBase = m.M
														v239 = m.ExcPending
														if v239 != 0 {
															return int32(0)
														} else {
															v240 = F_strlen(m, v230)
															mBase = m.M
															*(*int32)(unsafe.Add(mBase, uint32(v64))) = v240<<(uint(int32(2))%32) + int32(16)
															v250 = v64
															m.G0 = v14 + int32(96)
															return v250
														}
													} else {
														v240 = F_strlen(m, v230)
														mBase = m.M
														*(*int32)(unsafe.Add(mBase, uint32(v64))) = v240<<(uint(int32(2))%32) + int32(16)
														v250 = v64
														m.G0 = v14 + int32(96)
														return v250
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
}
func F_float4div(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 float32
	_ = v5
	var v11 float32
	_ = v11
	var v17 float32
	_ = v17
	var v19 float32
	_ = v19
	var v25 float32
	_ = v25
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	v5 = *(*float32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = *(*float32)(unsafe.Add(mBase, uint32(l0)+28))
	if base.B2i32(base.Ui32(base.I32_reinterpret_f32(v5)&int32(2147483647)) <= base.Ui32(int32(2139095040)))&base.F32_eq(v11, float32(0)) == int32(0) {
		v17 = base.F32_div(v5, v11)
		v19 = math.Float32frombits(uint32(0x7f800000))
		if base.F32_eq(base.F32_abs(v17), v19)&base.F32_ne(base.F32_abs(v5), v19) != 0 {
			F_float_overflow_error(m)
			mBase = m.M
			v43 = m.ExcPending
			if v43 != 0 {
				return int32(0)
			} else {
				base.Wasm_trap_unreachable()
				for {
				}
			}
		} else {
			v25 = float32(0)
			if base.B2i32(base.F32_eq(v5, v25)|base.F32_ne(v17, v25) == int32(0))&base.F32_ne(base.F32_abs(v11), math.Float32frombits(uint32(0x7f800000))) != 0 {
				F_float_underflow_error(m)
				mBase = m.M
				v45 = m.ExcPending
				if v45 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			} else {
				return base.I32_reinterpret_f32(v17)
			}
		}
	} else {
		F_float_zero_divide_error(m)
		mBase = m.M
		v41 = m.ExcPending
		if v41 != 0 {
			return int32(0)
		} else {
			base.Wasm_trap_unreachable()
			for {
			}
		}
	}
}
func F_float4in_internal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) float32 {
	mBase := m.M
	_ = mBase
	var v9 float32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v25 int32
	_ = v25
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v63 float32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v93 int32
	_ = v93
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v124 int32
	_ = v124
	var v128 float32
	_ = v128
	var v129 int32
	_ = v129
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
	var v150 int32
	_ = v150
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v205 int32
	_ = v205
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v236 int32
	_ = v236
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v259 int32
	_ = v259
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v290 int32
	_ = v290
	var v294 int32
	_ = v294
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v315 int32
	_ = v315
	var v324 int32
	_ = v324
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v346 int32
	_ = v346
	var v349 int32
	_ = v349
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v370 int32
	_ = v370
	var v379 int32
	_ = v379
	var v382 int32
	_ = v382
	var v384 int32
	_ = v384
	var v401 int32
	_ = v401
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v425 int32
	_ = v425
	var v434 int32
	_ = v434
	var v437 int32
	_ = v437
	var v439 int32
	_ = v439
	var v456 int32
	_ = v456
	var v463 int32
	_ = v463
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v473 int32
	_ = v473
	var v475 float32
	_ = v475
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v482 int32
	_ = v482
	var v488 int32
	_ = v488
	var v493 int32
	_ = v493
	var v494 float32
	_ = v494
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v501 int32
	_ = v501
	var v508 int32
	_ = v508
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v515 float32
	_ = v515
	var v516 int32
	_ = v516
	var v518 int32
	_ = v518
	var v520 float32
	_ = v520
	var v526 int32
	_ = v526
	var v531 int32
	_ = v531
	var v543 float32
	_ = v543
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v550 int32
	_ = v550
	var v555 int32
	_ = v555
	var v560 int32
	_ = v560
	var v569 float32
	_ = v569
	v9 = float32(0)
	v11 = m.G0
	v13 = v11 + int32(-64)
	m.G0 = v13
	v15 = l0
	goto L3
L1:
	;
	m.G0 = v13 - int32(-64)
	return v569
L2:
	;
	v526 = v518
	goto L162
L3:
	;
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
	if base.B2i32(base.Ui32(v25-int32(9)) < base.Ui32(int32(5)))|base.B2i32(v25 == int32(32)) != 0 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v516 = v15 + v514
	*(*int32)(unsafe.Add(mBase, uint32(v13)+60)) = v516
	v518 = v516
	v520 = v515
	goto L2
L5:
	;
	v15 = v15 + int32(1)
	goto L3
L6:
	;
	if v25 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L4
L8:
	;
	goto L7
L9:
	;
	v37 = F_errsave_start(m, l3)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	goto L11
L11:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_float4in_internal[0])) = int32(0)
	v63 = F_strtof(m, v15, v11+int32(-4))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L12
	} else {
		goto L18
	}
L12:
	;
	return float32(0)
L13:
	;
	if v37 == int32(0) {
		v569 = v9
		goto L1
	} else {
		goto L14
	}
L14:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L12
	} else {
		goto L15
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+52)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v13)+48)) = l1
	F_errmsg(m, int32(_a_F_float4in_internal_0), v11+int32(-16))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L12
	} else {
		goto L16
	}
L16:
	;
	F_errsave_finish(m, l3, int32(_a_F_float4in_internal_1), int32(208), int32(_a_F_float4in_internal_2))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L12
	} else {
		goto L17
	}
L17:
	;
	v569 = v9
	goto L1
L18:
	;
	v66 = *(*int32)(unsafe.Add(mBase, _c_F_float4in_internal[0]))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v13)+60))
	if v66|base.B2i32(v67 == v15) == int32(0) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v518 = v67
	v520 = v63
	goto L2
L20:
	;
	goto L21
L21:
	;
	v72 = int32(3)
	v77 = v15
	v78 = int32(_a_F_float4in_internal_3)
	v79 = v72
	goto L23
L22:
	;
	if v124 == int32(0) {
		goto L38
	} else {
		goto L39
	}
L23:
	;
	if v79 != 0 {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	v124 = int32(0)
	goto L22
L25:
	;
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77))))
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78))))
	if v82 == v83 {
		v105 = v82
		goto L28
	} else {
		goto L29
	}
L26:
	;
	goto L27
L27:
	;
	goto L24
L28:
	;
	v107 = int32(1)
	if v105 != 0 {
		v77 = v77 + v107
		v78 = v78 + v107
		v79 = v79 - v107
		goto L23
	} else {
		goto L37
	}
L29:
	;
	if base.Ui32((v82-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v93 = v82 | int32(32)
	goto L32
L31:
	;
	v93 = v82
	goto L32
L32:
	;
	if base.Ui32((v83-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v102 = v83 | int32(32)
	goto L35
L34:
	;
	v102 = v83
	goto L35
L35:
	;
	if v93 == v102 {
		v105 = v93
		goto L28
	} else {
		goto L36
	}
L36:
	;
	v124 = v93 - v102
	goto L22
L37:
	;
	goto L27
L38:
	;
	v514 = v72
	v515 = math.Float32frombits(uint32(0x7fc00000))
	goto L8
L39:
	;
	goto L40
L40:
	;
	v128 = math.Float32frombits(uint32(0x7f800000))
	v129 = int32(8)
	v134 = v15
	v135 = int32(_a_F_float4in_internal_4)
	v136 = v129
	goto L42
L41:
	;
	if v181 == int32(0) {
		v514 = v129
		v515 = v128
		goto L8
	} else {
		goto L57
	}
L42:
	;
	if v136 != 0 {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	v181 = int32(0)
	goto L41
L44:
	;
	v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v134))))
	v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135))))
	if v139 == v140 {
		v162 = v139
		goto L47
	} else {
		goto L48
	}
L45:
	;
	goto L46
L46:
	;
	goto L43
L47:
	;
	v164 = int32(1)
	if v162 != 0 {
		v134 = v134 + v164
		v135 = v135 + v164
		v136 = v136 - v164
		goto L42
	} else {
		goto L56
	}
L48:
	;
	if base.Ui32((v139-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v150 = v139 | int32(32)
	goto L51
L50:
	;
	v150 = v139
	goto L51
L51:
	;
	if base.Ui32((v140-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v159 = v140 | int32(32)
	goto L54
L53:
	;
	v159 = v140
	goto L54
L54:
	;
	if v150 == v159 {
		v162 = v150
		goto L47
	} else {
		goto L55
	}
L55:
	;
	v181 = v150 - v159
	goto L41
L56:
	;
	goto L46
L57:
	;
	v184 = int32(9)
	v189 = v15
	v190 = int32(_a_F_float4in_internal_5)
	v191 = v184
	goto L59
L58:
	;
	if v236 == int32(0) {
		v514 = v184
		v515 = v128
		goto L8
	} else {
		goto L74
	}
L59:
	;
	if v191 != 0 {
		goto L61
	} else {
		goto L62
	}
L60:
	;
	v236 = int32(0)
	goto L58
L61:
	;
	v194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v189))))
	v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v190))))
	if v194 == v195 {
		v217 = v194
		goto L64
	} else {
		goto L65
	}
L62:
	;
	goto L63
L63:
	;
	goto L60
L64:
	;
	v219 = int32(1)
	if v217 != 0 {
		v189 = v189 + v219
		v190 = v190 + v219
		v191 = v191 - v219
		goto L59
	} else {
		goto L73
	}
L65:
	;
	if base.Ui32((v194-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v205 = v194 | int32(32)
	goto L68
L67:
	;
	v205 = v194
	goto L68
L68:
	;
	if base.Ui32((v195-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v214 = v195 | int32(32)
	goto L71
L70:
	;
	v214 = v195
	goto L71
L71:
	;
	if v205 == v214 {
		v217 = v205
		goto L64
	} else {
		goto L72
	}
L72:
	;
	v236 = v205 - v214
	goto L58
L73:
	;
	goto L63
L74:
	;
	v243 = v15
	v244 = int32(_a_F_float4in_internal_6)
	v245 = int32(9)
	goto L76
L75:
	;
	if v290 == int32(0) {
		goto L91
	} else {
		goto L92
	}
L76:
	;
	if v245 != 0 {
		goto L78
	} else {
		goto L79
	}
L77:
	;
	v290 = int32(0)
	goto L75
L78:
	;
	v248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v243))))
	v249 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v244))))
	if v248 == v249 {
		v271 = v248
		goto L81
	} else {
		goto L82
	}
L79:
	;
	goto L80
L80:
	;
	goto L77
L81:
	;
	v273 = int32(1)
	if v271 != 0 {
		v243 = v243 + v273
		v244 = v244 + v273
		v245 = v245 - v273
		goto L76
	} else {
		goto L90
	}
L82:
	;
	if base.Ui32((v248-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v259 = v248 | int32(32)
	goto L85
L84:
	;
	v259 = v248
	goto L85
L85:
	;
	if base.Ui32((v249-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v268 = v249 | int32(32)
	goto L88
L87:
	;
	v268 = v249
	goto L88
L88:
	;
	if v259 == v268 {
		v271 = v259
		goto L81
	} else {
		goto L89
	}
L89:
	;
	v290 = v259 - v268
	goto L75
L90:
	;
	goto L80
L91:
	;
	v514 = v184
	v515 = math.Float32frombits(uint32(0xff800000))
	goto L8
L92:
	;
	goto L93
L93:
	;
	v294 = int32(3)
	v299 = v15
	v300 = int32(_a_F_float4in_internal_7)
	v301 = v294
	goto L95
L94:
	;
	if v346 == int32(0) {
		v514 = v294
		v515 = v128
		goto L8
	} else {
		goto L110
	}
L95:
	;
	if v301 != 0 {
		goto L97
	} else {
		goto L98
	}
L96:
	;
	v346 = int32(0)
	goto L94
L97:
	;
	v304 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v299))))
	v305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v300))))
	if v304 == v305 {
		v327 = v304
		goto L100
	} else {
		goto L101
	}
L98:
	;
	goto L99
L99:
	;
	goto L96
L100:
	;
	v329 = int32(1)
	if v327 != 0 {
		v299 = v299 + v329
		v300 = v300 + v329
		v301 = v301 - v329
		goto L95
	} else {
		goto L109
	}
L101:
	;
	if base.Ui32((v304-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	v315 = v304 | int32(32)
	goto L104
L103:
	;
	v315 = v304
	goto L104
L104:
	;
	if base.Ui32((v305-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	v324 = v305 | int32(32)
	goto L107
L106:
	;
	v324 = v305
	goto L107
L107:
	;
	if v315 == v324 {
		v327 = v315
		goto L100
	} else {
		goto L108
	}
L108:
	;
	v346 = v315 - v324
	goto L94
L109:
	;
	goto L99
L110:
	;
	v349 = int32(4)
	v354 = v15
	v355 = int32(_a_F_float4in_internal_8)
	v356 = v349
	goto L112
L111:
	;
	if v401 == int32(0) {
		v514 = v349
		v515 = v128
		goto L8
	} else {
		goto L127
	}
L112:
	;
	if v356 != 0 {
		goto L114
	} else {
		goto L115
	}
L113:
	;
	v401 = int32(0)
	goto L111
L114:
	;
	v359 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v354))))
	v360 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v355))))
	if v359 == v360 {
		v382 = v359
		goto L117
	} else {
		goto L118
	}
L115:
	;
	goto L116
L116:
	;
	goto L113
L117:
	;
	v384 = int32(1)
	if v382 != 0 {
		v354 = v354 + v384
		v355 = v355 + v384
		v356 = v356 - v384
		goto L112
	} else {
		goto L126
	}
L118:
	;
	if base.Ui32((v359-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L119
	} else {
		goto L120
	}
L119:
	;
	v370 = v359 | int32(32)
	goto L121
L120:
	;
	v370 = v359
	goto L121
L121:
	;
	if base.Ui32((v360-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L122
	} else {
		goto L123
	}
L122:
	;
	v379 = v360 | int32(32)
	goto L124
L123:
	;
	v379 = v360
	goto L124
L124:
	;
	if v370 == v379 {
		v382 = v370
		goto L117
	} else {
		goto L125
	}
L125:
	;
	v401 = v370 - v379
	goto L111
L126:
	;
	goto L116
L127:
	;
	v409 = v15
	v410 = int32(_a_F_float4in_internal_9)
	v411 = int32(4)
	goto L129
L128:
	;
	if v456 == int32(0) {
		v514 = v349
		v515 = math.Float32frombits(uint32(0xff800000))
		goto L8
	} else {
		goto L144
	}
L129:
	;
	if v411 != 0 {
		goto L131
	} else {
		goto L132
	}
L130:
	;
	v456 = int32(0)
	goto L128
L131:
	;
	v414 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v409))))
	v415 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v410))))
	if v414 == v415 {
		v437 = v414
		goto L134
	} else {
		goto L135
	}
L132:
	;
	goto L133
L133:
	;
	goto L130
L134:
	;
	v439 = int32(1)
	if v437 != 0 {
		v409 = v409 + v439
		v410 = v410 + v439
		v411 = v411 - v439
		goto L129
	} else {
		goto L143
	}
L135:
	;
	if base.Ui32((v414-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L136
	} else {
		goto L137
	}
L136:
	;
	v425 = v414 | int32(32)
	goto L138
L137:
	;
	v425 = v414
	goto L138
L138:
	;
	if base.Ui32((v415-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L139
	} else {
		goto L140
	}
L139:
	;
	v434 = v415 | int32(32)
	goto L141
L140:
	;
	v434 = v415
	goto L141
L141:
	;
	if v425 == v434 {
		v437 = v425
		goto L134
	} else {
		goto L142
	}
L142:
	;
	v456 = v425 - v434
	goto L128
L143:
	;
	goto L133
L144:
	;
	if v66 == int32(68) {
		goto L145
	} else {
		goto L146
	}
L145:
	;
	v463 = base.I32_reinterpret_f32(v63) & int32(2147483647)
	if base.B2i32(v463 != int32(0))&base.B2i32(v463 != int32(2139095040)) != 0 {
		goto L148
	} else {
		goto L149
	}
L146:
	;
	goto L147
L147:
	;
	v494 = float32(0)
	v495 = F_errsave_start(m, l3)
	mBase = m.M
	v496 = m.ExcPending
	if v496 != 0 {
		goto L12
	} else {
		goto L157
	}
L148:
	;
	v518 = v67
	v520 = v63
	goto L2
L149:
	;
	goto L150
L150:
	;
	v469 = F_pstrdup(m, v15)
	mBase = m.M
	v470 = m.ExcPending
	if v470 != 0 {
		goto L12
	} else {
		goto L151
	}
L151:
	;
	v473 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v469+(v67-v15)))) = uint8(v473)
	v475 = float32(0)
	v476 = F_errsave_start(m, l3)
	mBase = m.M
	v477 = m.ExcPending
	if v477 != 0 {
		goto L12
	} else {
		goto L152
	}
L152:
	;
	if v476 == int32(0) {
		v569 = v475
		goto L1
	} else {
		goto L153
	}
L153:
	;
	F_errcode(m, int32(50331778))
	mBase = m.M
	v482 = m.ExcPending
	if v482 != 0 {
		goto L12
	} else {
		goto L154
	}
L154:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v469
	F_errmsg(m, int32(_a_F_float4in_internal_10), v11+int32(-48))
	mBase = m.M
	v488 = m.ExcPending
	if v488 != 0 {
		goto L12
	} else {
		goto L155
	}
L155:
	;
	F_errsave_finish(m, l3, int32(_a_F_float4in_internal_1), int32(288), int32(_a_F_float4in_internal_2))
	mBase = m.M
	v493 = m.ExcPending
	if v493 != 0 {
		goto L12
	} else {
		goto L156
	}
L156:
	;
	v569 = v475
	goto L1
L157:
	;
	if v495 == int32(0) {
		v569 = v494
		goto L1
	} else {
		goto L158
	}
L158:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v501 = m.ExcPending
	if v501 != 0 {
		goto L12
	} else {
		goto L159
	}
L159:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+36)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v13)+32)) = l1
	F_errmsg(m, int32(_a_F_float4in_internal_0), v11+int32(-32))
	mBase = m.M
	v508 = m.ExcPending
	if v508 != 0 {
		goto L12
	} else {
		goto L160
	}
L160:
	;
	F_errsave_finish(m, l3, int32(_a_F_float4in_internal_1), int32(295), int32(_a_F_float4in_internal_2))
	mBase = m.M
	v513 = m.ExcPending
	if v513 != 0 {
		goto L12
	} else {
		goto L161
	}
L161:
	;
	v569 = v494
	goto L1
L162:
	;
	v531 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v526))))
	if base.B2i32(base.Ui32(v531-int32(9)) < base.Ui32(int32(5)))|base.B2i32(v531 == int32(32)) != 0 {
		goto L164
	} else {
		goto L165
	}
L163:
	;
	v543 = float32(0)
	v544 = F_errsave_start(m, l3)
	mBase = m.M
	v545 = m.ExcPending
	if v545 != 0 {
		goto L12
	} else {
		goto L170
	}
L164:
	;
	v526 = v526 + int32(1)
	goto L162
L165:
	;
	if v531 == int32(0) {
		goto L167
	} else {
		goto L168
	}
L166:
	;
	goto L163
L167:
	;
	v569 = v520
	goto L1
L168:
	;
	goto L169
L169:
	;
	goto L166
L170:
	;
	if v544 == int32(0) {
		v569 = v543
		goto L1
	} else {
		goto L171
	}
L171:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v550 = m.ExcPending
	if v550 != 0 {
		goto L12
	} else {
		goto L172
	}
L172:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = l1
	F_errmsg(m, int32(_a_F_float4in_internal_0), v13)
	mBase = m.M
	v555 = m.ExcPending
	if v555 != 0 {
		goto L12
	} else {
		goto L173
	}
L173:
	;
	F_errsave_finish(m, l3, int32(_a_F_float4in_internal_1), int32(309), int32(_a_F_float4in_internal_2))
	mBase = m.M
	v560 = m.ExcPending
	if v560 != 0 {
		goto L12
	} else {
		goto L174
	}
L174:
	;
	v569 = v543
	goto L1
}
func F_float4up(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	return v2
}
func F_float84gt(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 float32
	_ = v4
	var v5 float64
	_ = v5
	var v11 int32
	_ = v11
	var v12 float64
	_ = v12
	var v22 int32
	_ = v22
	v4 = *(*float32)(unsafe.Add(mBase, uint32(l0)+28))
	v5 = base.F64_promote_f32(v4)
	if base.Ui64(base.I64_reinterpret_f64(v5)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v12 = *(*float64)(unsafe.Add(mBase, uint32(v11)))
		v22 = base.F64_lt(v5, v12) | base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v12)&int64(9223372036854775807)))
	} else {
		v22 = int32(0)
	}
	return v22
}
func F_float8div(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 float64
	_ = v6
	var v12 int32
	_ = v12
	var v13 float64
	_ = v13
	var v19 float64
	_ = v19
	var v21 float64
	_ = v21
	var v27 float64
	_ = v27
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = *(*float64)(unsafe.Add(mBase, uint32(v5)))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v13 = *(*float64)(unsafe.Add(mBase, uint32(v12)))
	if base.B2i32(base.Ui64(base.I64_reinterpret_f64(v6)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)))&base.F64_eq(v13, float64(0)) == int32(0) {
		v19 = base.F64_div(v6, v13)
		v21 = math.Float64frombits(uint64(0x7ff0000000000000))
		if base.F64_eq(base.F64_abs(v19), v21)&base.F64_ne(base.F64_abs(v6), v21) != 0 {
			F_float_overflow_error(m)
			mBase = m.M
			v46 = m.ExcPending
			if v46 != 0 {
				return int32(0)
			} else {
				base.Wasm_trap_unreachable()
				for {
				}
			}
		} else {
			v27 = float64(0)
			if base.B2i32(base.F64_eq(v6, v27)|base.F64_ne(v19, v27) == int32(0))&base.F64_ne(base.F64_abs(v13), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
				F_float_underflow_error(m)
				mBase = m.M
				v48 = m.ExcPending
				if v48 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			} else {
				v38 = F_Float8GetDatum(m, v19)
				mBase = m.M
				v41 = m.ExcPending
				if v41 != 0 {
					return int32(0)
				} else {
					return v38
				}
			}
		}
	} else {
		F_float_zero_divide_error(m)
		mBase = m.M
		v44 = m.ExcPending
		if v44 != 0 {
			return int32(0)
		} else {
			base.Wasm_trap_unreachable()
			for {
			}
		}
	}
}
func F_float8gt(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 float64
	_ = v5
	var v11 int32
	_ = v11
	var v12 float64
	_ = v12
	var v22 int32
	_ = v22
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v5 = *(*float64)(unsafe.Add(mBase, uint32(v4)))
	if base.Ui64(base.I64_reinterpret_f64(v5)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v12 = *(*float64)(unsafe.Add(mBase, uint32(v11)))
		v22 = base.F64_lt(v5, v12) | base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v12)&int64(9223372036854775807)))
	} else {
		v22 = int32(0)
	}
	return v22
}
func F_float8in_internal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) float64 {
	mBase := m.M
	_ = mBase
	var v10 float64
	_ = v10
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v29 int32
	_ = v29
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v67 float64
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v97 int32
	_ = v97
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v128 int32
	_ = v128
	var v132 float64
	_ = v132
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v154 int32
	_ = v154
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v209 int32
	_ = v209
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v240 int32
	_ = v240
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v263 int32
	_ = v263
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v294 int32
	_ = v294
	var v298 int32
	_ = v298
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v319 int32
	_ = v319
	var v328 int32
	_ = v328
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v350 int32
	_ = v350
	var v353 int32
	_ = v353
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v374 int32
	_ = v374
	var v383 int32
	_ = v383
	var v386 int32
	_ = v386
	var v388 int32
	_ = v388
	var v405 int32
	_ = v405
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v429 int32
	_ = v429
	var v438 int32
	_ = v438
	var v441 int32
	_ = v441
	var v443 int32
	_ = v443
	var v460 int32
	_ = v460
	var v467 int64
	_ = v467
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v477 int32
	_ = v477
	var v479 float64
	_ = v479
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v486 int32
	_ = v486
	var v492 int32
	_ = v492
	var v497 int32
	_ = v497
	var v498 float64
	_ = v498
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v505 int32
	_ = v505
	var v512 int32
	_ = v512
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v519 float64
	_ = v519
	var v520 int32
	_ = v520
	var v522 int32
	_ = v522
	var v524 float64
	_ = v524
	var v532 int32
	_ = v532
	var v538 int32
	_ = v538
	var v553 float64
	_ = v553
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v560 int32
	_ = v560
	var v565 int32
	_ = v565
	var v570 int32
	_ = v570
	var v580 float64
	_ = v580
	v10 = float64(0)
	v13 = m.G0
	v15 = v13 + int32(-64)
	m.G0 = v15
	v17 = l0
	goto L3
L1:
	;
	m.G0 = v15 - int32(-64)
	return v580
L2:
	;
	v532 = v522
	goto L164
L3:
	;
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
	if base.B2i32(base.Ui32(v29-int32(9)) < base.Ui32(int32(5)))|base.B2i32(v29 == int32(32)) != 0 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v520 = v17 + v518
	*(*int32)(unsafe.Add(mBase, uint32(v15)+60)) = v520
	v522 = v520
	v524 = v519
	goto L2
L5:
	;
	v17 = v17 + int32(1)
	goto L3
L6:
	;
	if v29 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L4
L8:
	;
	goto L7
L9:
	;
	v41 = F_errsave_start(m, l4)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	goto L11
L11:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_float8in_internal[0])) = int32(0)
	v67 = F_strtod(m, v17, v13+int32(-4))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L12
	} else {
		goto L18
	}
L12:
	;
	return float64(0)
L13:
	;
	if v41 == int32(0) {
		v580 = v10
		goto L1
	} else {
		goto L14
	}
L14:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L12
	} else {
		goto L15
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+52)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v15)+48)) = l2
	F_errmsg(m, int32(_a_F_float8in_internal_0), v13+int32(-16))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L12
	} else {
		goto L16
	}
L16:
	;
	F_errsave_finish(m, l4, int32(_a_F_float8in_internal_1), int32(414), int32(_a_F_float8in_internal_2))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L12
	} else {
		goto L17
	}
L17:
	;
	v580 = v10
	goto L1
L18:
	;
	v70 = *(*int32)(unsafe.Add(mBase, _c_F_float8in_internal[0]))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v15)+60))
	if v70|base.B2i32(v71 == v17) == int32(0) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v522 = v71
	v524 = v67
	goto L2
L20:
	;
	goto L21
L21:
	;
	v76 = int32(3)
	v81 = v17
	v82 = int32(_a_F_float8in_internal_3)
	v83 = v76
	goto L23
L22:
	;
	if v128 == int32(0) {
		goto L38
	} else {
		goto L39
	}
L23:
	;
	if v83 != 0 {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	v128 = int32(0)
	goto L22
L25:
	;
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81))))
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82))))
	if v86 == v87 {
		v109 = v86
		goto L28
	} else {
		goto L29
	}
L26:
	;
	goto L27
L27:
	;
	goto L24
L28:
	;
	v111 = int32(1)
	if v109 != 0 {
		v81 = v81 + v111
		v82 = v82 + v111
		v83 = v83 - v111
		goto L23
	} else {
		goto L37
	}
L29:
	;
	if base.Ui32((v86-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v97 = v86 | int32(32)
	goto L32
L31:
	;
	v97 = v86
	goto L32
L32:
	;
	if base.Ui32((v87-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v106 = v87 | int32(32)
	goto L35
L34:
	;
	v106 = v87
	goto L35
L35:
	;
	if v97 == v106 {
		v109 = v97
		goto L28
	} else {
		goto L36
	}
L36:
	;
	v128 = v97 - v106
	goto L22
L37:
	;
	goto L27
L38:
	;
	v518 = v76
	v519 = math.Float64frombits(uint64(0x7ff8000000000000))
	goto L8
L39:
	;
	goto L40
L40:
	;
	v132 = math.Float64frombits(uint64(0x7ff0000000000000))
	v133 = int32(8)
	v138 = v17
	v139 = int32(_a_F_float8in_internal_4)
	v140 = v133
	goto L42
L41:
	;
	if v185 == int32(0) {
		v518 = v133
		v519 = v132
		goto L8
	} else {
		goto L57
	}
L42:
	;
	if v140 != 0 {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	v185 = int32(0)
	goto L41
L44:
	;
	v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v138))))
	v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v139))))
	if v143 == v144 {
		v166 = v143
		goto L47
	} else {
		goto L48
	}
L45:
	;
	goto L46
L46:
	;
	goto L43
L47:
	;
	v168 = int32(1)
	if v166 != 0 {
		v138 = v138 + v168
		v139 = v139 + v168
		v140 = v140 - v168
		goto L42
	} else {
		goto L56
	}
L48:
	;
	if base.Ui32((v143-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v154 = v143 | int32(32)
	goto L51
L50:
	;
	v154 = v143
	goto L51
L51:
	;
	if base.Ui32((v144-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v163 = v144 | int32(32)
	goto L54
L53:
	;
	v163 = v144
	goto L54
L54:
	;
	if v154 == v163 {
		v166 = v154
		goto L47
	} else {
		goto L55
	}
L55:
	;
	v185 = v154 - v163
	goto L41
L56:
	;
	goto L46
L57:
	;
	v188 = int32(9)
	v193 = v17
	v194 = int32(_a_F_float8in_internal_5)
	v195 = v188
	goto L59
L58:
	;
	if v240 == int32(0) {
		v518 = v188
		v519 = v132
		goto L8
	} else {
		goto L74
	}
L59:
	;
	if v195 != 0 {
		goto L61
	} else {
		goto L62
	}
L60:
	;
	v240 = int32(0)
	goto L58
L61:
	;
	v198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v193))))
	v199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v194))))
	if v198 == v199 {
		v221 = v198
		goto L64
	} else {
		goto L65
	}
L62:
	;
	goto L63
L63:
	;
	goto L60
L64:
	;
	v223 = int32(1)
	if v221 != 0 {
		v193 = v193 + v223
		v194 = v194 + v223
		v195 = v195 - v223
		goto L59
	} else {
		goto L73
	}
L65:
	;
	if base.Ui32((v198-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v209 = v198 | int32(32)
	goto L68
L67:
	;
	v209 = v198
	goto L68
L68:
	;
	if base.Ui32((v199-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v218 = v199 | int32(32)
	goto L71
L70:
	;
	v218 = v199
	goto L71
L71:
	;
	if v209 == v218 {
		v221 = v209
		goto L64
	} else {
		goto L72
	}
L72:
	;
	v240 = v209 - v218
	goto L58
L73:
	;
	goto L63
L74:
	;
	v247 = v17
	v248 = int32(_a_F_float8in_internal_6)
	v249 = int32(9)
	goto L76
L75:
	;
	if v294 == int32(0) {
		goto L91
	} else {
		goto L92
	}
L76:
	;
	if v249 != 0 {
		goto L78
	} else {
		goto L79
	}
L77:
	;
	v294 = int32(0)
	goto L75
L78:
	;
	v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v247))))
	v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v248))))
	if v252 == v253 {
		v275 = v252
		goto L81
	} else {
		goto L82
	}
L79:
	;
	goto L80
L80:
	;
	goto L77
L81:
	;
	v277 = int32(1)
	if v275 != 0 {
		v247 = v247 + v277
		v248 = v248 + v277
		v249 = v249 - v277
		goto L76
	} else {
		goto L90
	}
L82:
	;
	if base.Ui32((v252-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v263 = v252 | int32(32)
	goto L85
L84:
	;
	v263 = v252
	goto L85
L85:
	;
	if base.Ui32((v253-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v272 = v253 | int32(32)
	goto L88
L87:
	;
	v272 = v253
	goto L88
L88:
	;
	if v263 == v272 {
		v275 = v263
		goto L81
	} else {
		goto L89
	}
L89:
	;
	v294 = v263 - v272
	goto L75
L90:
	;
	goto L80
L91:
	;
	v518 = v188
	v519 = math.Float64frombits(uint64(0xfff0000000000000))
	goto L8
L92:
	;
	goto L93
L93:
	;
	v298 = int32(3)
	v303 = v17
	v304 = int32(_a_F_float8in_internal_7)
	v305 = v298
	goto L95
L94:
	;
	if v350 == int32(0) {
		v518 = v298
		v519 = v132
		goto L8
	} else {
		goto L110
	}
L95:
	;
	if v305 != 0 {
		goto L97
	} else {
		goto L98
	}
L96:
	;
	v350 = int32(0)
	goto L94
L97:
	;
	v308 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v303))))
	v309 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v304))))
	if v308 == v309 {
		v331 = v308
		goto L100
	} else {
		goto L101
	}
L98:
	;
	goto L99
L99:
	;
	goto L96
L100:
	;
	v333 = int32(1)
	if v331 != 0 {
		v303 = v303 + v333
		v304 = v304 + v333
		v305 = v305 - v333
		goto L95
	} else {
		goto L109
	}
L101:
	;
	if base.Ui32((v308-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	v319 = v308 | int32(32)
	goto L104
L103:
	;
	v319 = v308
	goto L104
L104:
	;
	if base.Ui32((v309-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	v328 = v309 | int32(32)
	goto L107
L106:
	;
	v328 = v309
	goto L107
L107:
	;
	if v319 == v328 {
		v331 = v319
		goto L100
	} else {
		goto L108
	}
L108:
	;
	v350 = v319 - v328
	goto L94
L109:
	;
	goto L99
L110:
	;
	v353 = int32(4)
	v358 = v17
	v359 = int32(_a_F_float8in_internal_8)
	v360 = v353
	goto L112
L111:
	;
	if v405 == int32(0) {
		v518 = v353
		v519 = v132
		goto L8
	} else {
		goto L127
	}
L112:
	;
	if v360 != 0 {
		goto L114
	} else {
		goto L115
	}
L113:
	;
	v405 = int32(0)
	goto L111
L114:
	;
	v363 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v358))))
	v364 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v359))))
	if v363 == v364 {
		v386 = v363
		goto L117
	} else {
		goto L118
	}
L115:
	;
	goto L116
L116:
	;
	goto L113
L117:
	;
	v388 = int32(1)
	if v386 != 0 {
		v358 = v358 + v388
		v359 = v359 + v388
		v360 = v360 - v388
		goto L112
	} else {
		goto L126
	}
L118:
	;
	if base.Ui32((v363-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L119
	} else {
		goto L120
	}
L119:
	;
	v374 = v363 | int32(32)
	goto L121
L120:
	;
	v374 = v363
	goto L121
L121:
	;
	if base.Ui32((v364-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L122
	} else {
		goto L123
	}
L122:
	;
	v383 = v364 | int32(32)
	goto L124
L123:
	;
	v383 = v364
	goto L124
L124:
	;
	if v374 == v383 {
		v386 = v374
		goto L117
	} else {
		goto L125
	}
L125:
	;
	v405 = v374 - v383
	goto L111
L126:
	;
	goto L116
L127:
	;
	v413 = v17
	v414 = int32(_a_F_float8in_internal_9)
	v415 = int32(4)
	goto L129
L128:
	;
	if v460 == int32(0) {
		v518 = v353
		v519 = math.Float64frombits(uint64(0xfff0000000000000))
		goto L8
	} else {
		goto L144
	}
L129:
	;
	if v415 != 0 {
		goto L131
	} else {
		goto L132
	}
L130:
	;
	v460 = int32(0)
	goto L128
L131:
	;
	v418 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v413))))
	v419 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v414))))
	if v418 == v419 {
		v441 = v418
		goto L134
	} else {
		goto L135
	}
L132:
	;
	goto L133
L133:
	;
	goto L130
L134:
	;
	v443 = int32(1)
	if v441 != 0 {
		v413 = v413 + v443
		v414 = v414 + v443
		v415 = v415 - v443
		goto L129
	} else {
		goto L143
	}
L135:
	;
	if base.Ui32((v418-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L136
	} else {
		goto L137
	}
L136:
	;
	v429 = v418 | int32(32)
	goto L138
L137:
	;
	v429 = v418
	goto L138
L138:
	;
	if base.Ui32((v419-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L139
	} else {
		goto L140
	}
L139:
	;
	v438 = v419 | int32(32)
	goto L141
L140:
	;
	v438 = v419
	goto L141
L141:
	;
	if v429 == v438 {
		v441 = v429
		goto L134
	} else {
		goto L142
	}
L142:
	;
	v460 = v429 - v438
	goto L128
L143:
	;
	goto L133
L144:
	;
	if v70 == int32(68) {
		goto L145
	} else {
		goto L146
	}
L145:
	;
	v467 = base.I64_reinterpret_f64(v67) & int64(9223372036854775807)
	if base.B2i32(v467 != int64(0))&base.B2i32(v467 != int64(9218868437227405312)) != 0 {
		goto L148
	} else {
		goto L149
	}
L146:
	;
	goto L147
L147:
	;
	v498 = float64(0)
	v499 = F_errsave_start(m, l4)
	mBase = m.M
	v500 = m.ExcPending
	if v500 != 0 {
		goto L12
	} else {
		goto L157
	}
L148:
	;
	v522 = v71
	v524 = v67
	goto L2
L149:
	;
	goto L150
L150:
	;
	v473 = F_pstrdup(m, v17)
	mBase = m.M
	v474 = m.ExcPending
	if v474 != 0 {
		goto L12
	} else {
		goto L151
	}
L151:
	;
	v477 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v473+(v71-v17)))) = uint8(v477)
	v479 = float64(0)
	v480 = F_errsave_start(m, l4)
	mBase = m.M
	v481 = m.ExcPending
	if v481 != 0 {
		goto L12
	} else {
		goto L152
	}
L152:
	;
	if v480 == int32(0) {
		v580 = v479
		goto L1
	} else {
		goto L153
	}
L153:
	;
	F_errcode(m, int32(50331778))
	mBase = m.M
	v486 = m.ExcPending
	if v486 != 0 {
		goto L12
	} else {
		goto L154
	}
L154:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v473
	F_errmsg(m, int32(_a_F_float8in_internal_10), v13+int32(-48))
	mBase = m.M
	v492 = m.ExcPending
	if v492 != 0 {
		goto L12
	} else {
		goto L155
	}
L155:
	;
	F_errsave_finish(m, l4, int32(_a_F_float8in_internal_1), int32(490), int32(_a_F_float8in_internal_2))
	mBase = m.M
	v497 = m.ExcPending
	if v497 != 0 {
		goto L12
	} else {
		goto L156
	}
L156:
	;
	v580 = v479
	goto L1
L157:
	;
	if v499 == int32(0) {
		v580 = v498
		goto L1
	} else {
		goto L158
	}
L158:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v505 = m.ExcPending
	if v505 != 0 {
		goto L12
	} else {
		goto L159
	}
L159:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+36)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = l2
	F_errmsg(m, int32(_a_F_float8in_internal_0), v13+int32(-32))
	mBase = m.M
	v512 = m.ExcPending
	if v512 != 0 {
		goto L12
	} else {
		goto L160
	}
L160:
	;
	F_errsave_finish(m, l4, int32(_a_F_float8in_internal_1), int32(497), int32(_a_F_float8in_internal_2))
	mBase = m.M
	v517 = m.ExcPending
	if v517 != 0 {
		goto L12
	} else {
		goto L161
	}
L161:
	;
	v580 = v498
	goto L1
L162:
	;
	v553 = float64(0)
	v554 = F_errsave_start(m, l4)
	mBase = m.M
	v555 = m.ExcPending
	if v555 != 0 {
		goto L12
	} else {
		goto L174
	}
L163:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v532
	v580 = v524
	goto L1
L164:
	;
	v538 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v532))))
	if base.B2i32(base.Ui32(v538-int32(9)) < base.Ui32(int32(5)))|base.B2i32(v538 == int32(32)) != 0 {
		goto L166
	} else {
		goto L167
	}
L165:
	;
	if l1 == int32(0) {
		goto L162
	} else {
		goto L173
	}
L166:
	;
	v532 = v532 + int32(1)
	goto L164
L167:
	;
	if v538 == int32(0) {
		goto L169
	} else {
		goto L170
	}
L168:
	;
	goto L165
L169:
	;
	if l1 != 0 {
		goto L163
	} else {
		goto L172
	}
L170:
	;
	goto L171
L171:
	;
	goto L168
L172:
	;
	v580 = v524
	goto L1
L173:
	;
	goto L163
L174:
	;
	if v554 == int32(0) {
		v580 = v553
		goto L1
	} else {
		goto L175
	}
L175:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v560 = m.ExcPending
	if v560 != 0 {
		goto L12
	} else {
		goto L176
	}
L176:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = l2
	F_errmsg(m, int32(_a_F_float8in_internal_0), v15)
	mBase = m.M
	v565 = m.ExcPending
	if v565 != 0 {
		goto L12
	} else {
		goto L177
	}
L177:
	;
	F_errsave_finish(m, l4, int32(_a_F_float8in_internal_1), int32(511), int32(_a_F_float8in_internal_2))
	mBase = m.M
	v570 = m.ExcPending
	if v570 != 0 {
		goto L12
	} else {
		goto L178
	}
L178:
	;
	v580 = v553
	goto L1
}
func F_float8le(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 float64
	_ = v5
	var v11 int32
	_ = v11
	var v12 float64
	_ = v12
	var v22 int32
	_ = v22
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v5 = *(*float64)(unsafe.Add(mBase, uint32(v4)))
	if base.Ui64(base.I64_reinterpret_f64(v5)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v12 = *(*float64)(unsafe.Add(mBase, uint32(v11)))
		v22 = base.F64_ge(v5, v12) & base.B2i32(base.Ui64(base.I64_reinterpret_f64(v12)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)))
	} else {
		v22 = int32(1)
	}
	return v22
}
func F_flt4_mul_cash(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int64
	_ = v3
	var v4 float32
	_ = v4
	var v6 int64
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v3 = *(*int64)(unsafe.Add(mBase, uint32(v2)))
	v4 = *(*float32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = F_cash_mul_float8(m, v3, base.F64_promote_f32(v4))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		v10 = F_Int64GetDatum(m, v6)
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			return v10
		}
	}
}
func F_flt8_mul_cash(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int64
	_ = v3
	var v4 int32
	_ = v4
	var v5 float64
	_ = v5
	var v6 int64
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v3 = *(*int64)(unsafe.Add(mBase, uint32(v2)))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = *(*float64)(unsafe.Add(mBase, uint32(v4)))
	v6 = F_cash_mul_float8(m, v3, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		v10 = F_Int64GetDatum(m, v6)
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			return v10
		}
	}
}
func F_fmt_u(m *base.Module, l0 int64, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int64
	_ = v8
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v15 int64
	_ = v15
	var v16 int64
	_ = v16
	var v22 int32
	_ = v22
	var v26 int64
	_ = v26
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	if base.Ui64(int64(4294967296)) <= base.Ui64(l0) {
		v8 = l0
		v9 = l1
		for {
			v14 = v9 - int32(1)
			v15 = int64(10)
			v16 = base.I64_div_u_s(v8, v15)
			v22 = base.I32_wrap_i64(v8-v16*v15) | int32(48)
			*(*uint8)(unsafe.Add(mBase, uint32(v14))) = uint8(v22)
			if base.Ui64(int64(42949672959)) < base.Ui64(v8) {
				v8 = v16
				v9 = v14
				continue
			} else {
				break
			}
			break
		}
		v26 = v16
		v27 = v14
	} else {
		v26 = l0
		v27 = l1
	}
	v31 = base.I32_wrap_i64(v26)
	if base.Ui64(int64(10)) <= base.Ui64(v26) {
		v35 = v27
		v36 = v31
		for {
			v40 = v35 - int32(1)
			v41 = int32(10)
			v42 = base.I32_div_u_s(v36, v41)
			v47 = v36 - v42*v41 | int32(48)
			*(*uint8)(unsafe.Add(mBase, uint32(v40))) = uint8(v47)
			if base.Ui32(int32(99)) < base.Ui32(v36) {
				v35 = v40
				v36 = v42
				continue
			} else {
				break
			}
			break
		}
		v52 = v40
		v53 = v42
	} else {
		v52 = v27
		v53 = v31
	}
	if v53 != 0 {
		v57 = v52 - int32(1)
		v59 = v53 | int32(48)
		*(*uint8)(unsafe.Add(mBase, uint32(v57))) = uint8(v59)
		v61 = v57
	} else {
		v61 = v52
	}
	return v61
}
func F_forkname_chars(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	v5 = int32(3)
	v6 = int32(_a_F_forkname_chars_0)
	goto L6
L1:
	;
	return v169
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v167
	v169 = v166
	goto L1
L3:
	;
	if l1 == int32(0) {
		v169 = v162
		goto L1
	} else {
		goto L52
	}
L4:
	;
	if v44-v45 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L6:
	;
	goto L7
L7:
	;
	v13 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_forkname_chars[0])))
	if v13 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v14 = v6
	v15 = l0
	v16 = v5
	v17 = v13
	goto L12
L9:
	;
	v40 = l0
	v44 = int32(0)
	goto L10
L10:
	;
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40))))
	goto L4
L11:
	;
	v40 = v35
	v44 = v37
	goto L10
L12:
	;
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
	if base.B2i32(v17 != v19)|base.B2i32(v19 == int32(0)) != 0 {
		v35 = v15
		v37 = v17
		goto L11
	} else {
		goto L14
	}
L13:
	;
	v35 = v29
	v37 = int32(0)
	goto L11
L14:
	;
	v25 = v16 - int32(1)
	if v25 == int32(0) {
		v35 = v15
		v37 = v17
		goto L11
	} else {
		goto L15
	}
L15:
	;
	v28 = int32(1)
	v29 = v15 + v28
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+1)))
	if v30 != 0 {
		v14 = v14 + v28
		v15 = v29
		v16 = v25
		v17 = v30
		goto L12
	} else {
		goto L16
	}
L16:
	;
	goto L13
L17:
	;
	v162 = v5
	v163 = int32(1)
	goto L3
L18:
	;
	goto L19
L19:
	;
	v56 = int32(2)
	v58 = int32(_a_F_forkname_chars_1)
	goto L22
L20:
	;
	if v96-v97 == int32(0) {
		v162 = v56
		v163 = v56
		goto L3
	} else {
		goto L33
	}
L22:
	;
	goto L23
L23:
	;
	v65 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_forkname_chars[1])))
	if v65 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v66 = v58
	v67 = l0
	v68 = v56
	v69 = v65
	goto L28
L25:
	;
	v92 = l0
	v96 = int32(0)
	goto L26
L26:
	;
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92))))
	goto L20
L27:
	;
	v92 = v87
	v96 = v89
	goto L26
L28:
	;
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67))))
	if base.B2i32(v69 != v71)|base.B2i32(v71 == int32(0)) != 0 {
		v87 = v67
		v89 = v69
		goto L27
	} else {
		goto L30
	}
L29:
	;
	v87 = v81
	v89 = int32(0)
	goto L27
L30:
	;
	v77 = v68 - int32(1)
	if v77 == int32(0) {
		v87 = v67
		v89 = v69
		goto L27
	} else {
		goto L31
	}
L31:
	;
	v80 = int32(1)
	v81 = v67 + v80
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+1)))
	if v82 != 0 {
		v66 = v66 + v80
		v67 = v81
		v68 = v77
		v69 = v82
		goto L28
	} else {
		goto L32
	}
L32:
	;
	goto L29
L33:
	;
	v107 = int32(4)
	v108 = int32(_a_F_forkname_chars_2)
	goto L36
L34:
	;
	if v146-v147 == int32(0) {
		goto L47
	} else {
		goto L48
	}
L36:
	;
	goto L37
L37:
	;
	v115 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_forkname_chars[2])))
	if v115 != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v116 = v108
	v117 = l0
	v118 = v107
	v119 = v115
	goto L42
L39:
	;
	v142 = l0
	v146 = int32(0)
	goto L40
L40:
	;
	v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v142))))
	goto L34
L41:
	;
	v142 = v137
	v146 = v139
	goto L40
L42:
	;
	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117))))
	if base.B2i32(v119 != v121)|base.B2i32(v121 == int32(0)) != 0 {
		v137 = v117
		v139 = v119
		goto L41
	} else {
		goto L44
	}
L43:
	;
	v137 = v131
	v139 = int32(0)
	goto L41
L44:
	;
	v127 = v118 - int32(1)
	if v127 == int32(0) {
		v137 = v117
		v139 = v119
		goto L41
	} else {
		goto L45
	}
L45:
	;
	v130 = int32(1)
	v131 = v117 + v130
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116)+1)))
	if v132 != 0 {
		v116 = v116 + v130
		v117 = v131
		v118 = v127
		v119 = v132
		goto L42
	} else {
		goto L46
	}
L46:
	;
	goto L43
L47:
	;
	if l1 != 0 {
		v166 = v107
		v167 = int32(3)
		goto L2
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	v159 = int32(0)
	if l1 == v159 {
		v169 = v159
		goto L1
	} else {
		goto L51
	}
L50:
	;
	v169 = v107
	goto L1
L51:
	;
	v166 = v159
	v167 = int32(-1)
	goto L2
L52:
	;
	v166 = v162
	v167 = v163
	goto L2
}
func F_free_parsestate(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	if v8-int32(1) < int32(1665) {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
		if v13 != 0 {
			F_relation_close(m, v13, int32(0))
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return
			} else {
				F_pfree(m, l0)
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return
				} else {
					m.G0 = v6 + int32(16)
					return
				}
			}
		} else {
			F_pfree(m, l0)
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return
			} else {
				m.G0 = v6 + int32(16)
				return
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v25 = m.ExcPending
		if v25 != 0 {
			return
		} else {
			F_errcode(m, int32(17039621))
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v6))) = int32(1664)
				F_errmsg(m, int32(_a_F_free_parsestate_0), v6)
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_free_parsestate_1), int32(83), int32(_a_F_free_parsestate_2))
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
						return
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
func F_free_struct_lconv(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_emscripten_builtin_free(m, v2)
	mBase = m.M
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	F_emscripten_builtin_free(m, v4)
	mBase = m.M
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_emscripten_builtin_free(m, v6)
	mBase = m.M
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	F_emscripten_builtin_free(m, v8)
	mBase = m.M
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	F_emscripten_builtin_free(m, v10)
	mBase = m.M
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	F_emscripten_builtin_free(m, v12)
	mBase = m.M
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	F_emscripten_builtin_free(m, v14)
	mBase = m.M
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	F_emscripten_builtin_free(m, v16)
	mBase = m.M
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	F_emscripten_builtin_free(m, v18)
	mBase = m.M
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	F_emscripten_builtin_free(m, v20)
	mBase = m.M
	return
}
func F_freeifaddrs(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	if l0 != 0 {
		v3 = l0
		for {
			v5 = *(*int32)(unsafe.Add(mBase, uint32(v3)))
			F_emscripten_builtin_free(m, v3)
			mBase = m.M
			if v5 != 0 {
				v3 = v5
				continue
			} else {
				break
			}
			break
		}
	} else {
	}
	return
}
func F_frexp(m *base.Module, l0 float64, l1 int32) float64 {
	mBase := m.M
	_ = mBase
	var v5 int64
	_ = v5
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v19 float64
	_ = v19
	var v22 int64
	_ = v22
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v37 float64
	_ = v37
	var v38 int32
	_ = v38
	var v41 float64
	_ = v41
	var v42 int32
	_ = v42
	var v52 float64
	_ = v52
	var v54 float64
	_ = v54
	var v55 int32
	_ = v55
	var v58 float64
	_ = v58
	var v59 int32
	_ = v59
	var v70 float64
	_ = v70
	v5 = base.I64_reinterpret_f64(l0)
	v9 = int32(2047)
	v10 = base.I32_wrap_i64(int64(base.Ui64(v5)>>(uint(int64(52))%64))) & v9
	if v10 != v9 {
		if v10 == int32(0) {
			if base.F64_eq(l0, float64(0)) != 0 {
				v58 = l0
				v59 = int32(0)
			} else {
				v19 = base.F64_mul(l0, float64(1.8446744073709552e+19))
				v22 = base.I64_reinterpret_f64(v19)
				v26 = int32(2047)
				v27 = base.I32_wrap_i64(int64(base.Ui64(v22)>>(uint(int64(52))%64))) & v26
				if v27 != v26 {
					if v27 == int32(0) {
						if base.F64_eq(v19, float64(0)) != 0 {
							v41 = v19
							v42 = int32(0)
						} else {
							v37 = F_frexp(m, base.F64_mul(v19, float64(1.8446744073709552e+19)), l1)
							mBase = m.M
							v38 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
							v41 = v37
							v42 = v38 + int32(-64)
						}
						*(*int32)(unsafe.Add(mBase, uint32(l1))) = v42
						v54 = v41
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l1))) = v27 - int32(1022)
						v52 = base.F64_reinterpret_i64(v22&int64(-9218868437227405313) | int64(4602678819172646912))
						v54 = v52
					}
				} else {
					v52 = v19
					v54 = v52
				}
				v55 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				v58 = v54
				v59 = v55 + int32(-64)
			}
			*(*int32)(unsafe.Add(mBase, uint32(l1))) = v59
			return v58
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l1))) = v10 - int32(1022)
			v70 = base.F64_reinterpret_i64(v5&int64(-9218868437227405313) | int64(4602678819172646912))
			return v70
		}
	} else {
		v70 = l0
		return v70
	}
}
func F_fscanf(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = l2
	v10 = F_vfscanf(m, l0, l1, l2)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		m.G0 = v7 + int32(16)
		return v10
	}
}
func F_fsync_fname(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	v7 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_fsync_fname[0])))
	if v7 != 0 {
		v8 = int32(21)
	} else {
		v8 = int32(23)
	}
	v9 = F_fsync_fname_ext(m, l0, l1, int32(0), v8)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return
	} else {
		return
	}
}
