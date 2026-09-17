package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_pgp_armor_decode(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v175 int32
	_ = v175
	var v186 int32
	_ = v186
	var v192 int32
	_ = v192
	var v198 int32
	_ = v198
	var v203 int32
	_ = v203
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v245 int32
	_ = v245
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v260 int32
	_ = v260
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v304 int32
	_ = v304
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v323 int32
	_ = v323
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v333 int32
	_ = v333
	var v335 int32
	_ = v335
	var v340 int32
	_ = v340
	var v342 int32
	_ = v342
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v352 int32
	_ = v352
	var v354 int32
	_ = v354
	var v357 int32
	_ = v357
	var v369 int32
	_ = v369
	var v373 int32
	_ = v373
	var v375 int32
	_ = v375
	var v386 int32
	_ = v386
	var v389 int32
	_ = v389
	var v393 int32
	_ = v393
	var v400 int32
	_ = v400
	var v402 int32
	_ = v402
	var v409 int32
	_ = v409
	var v412 int32
	_ = v412
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v436 int32
	_ = v436
	var v440 int32
	_ = v440
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v453 int32
	_ = v453
	var v477 int32
	_ = v477
	var v488 int32
	_ = v488
	var v490 int32
	_ = v490
	var v493 int32
	_ = v493
	var v498 int32
	_ = v498
	var v502 int32
	_ = v502
	var v504 int32
	_ = v504
	var v508 int32
	_ = v508
	var v510 int32
	_ = v510
	var v513 int32
	_ = v513
	var v517 int32
	_ = v517
	var v527 int32
	_ = v527
	var v530 int32
	_ = v530
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v547 int32
	_ = v547
	var v550 int32
	_ = v550
	var v572 int32
	_ = v572
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v582 int32
	_ = v582
	var v586 int32
	_ = v586
	var v589 int32
	_ = v589
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v606 int32
	_ = v606
	var v610 int32
	_ = v610
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v623 int32
	_ = v623
	var v647 int32
	_ = v647
	var v658 int32
	_ = v658
	var v660 int32
	_ = v660
	var v663 int32
	_ = v663
	var v668 int32
	_ = v668
	var v672 int32
	_ = v672
	var v674 int32
	_ = v674
	var v678 int32
	_ = v678
	var v680 int32
	_ = v680
	var v683 int32
	_ = v683
	var v687 int32
	_ = v687
	var v697 int32
	_ = v697
	var v700 int32
	_ = v700
	var v703 int32
	_ = v703
	var v704 int32
	_ = v704
	var v705 int32
	_ = v705
	var v710 int32
	_ = v710
	var v711 int32
	_ = v711
	var v712 int32
	_ = v712
	var v713 int32
	_ = v713
	var v717 int32
	_ = v717
	var v720 int32
	_ = v720
	var v726 int32
	_ = v726
	var v742 int32
	_ = v742
	var v753 int32
	_ = v753
	var v755 int32
	_ = v755
	var v758 int32
	_ = v758
	var v761 int32
	_ = v761
	var v764 int32
	_ = v764
	var v767 int32
	_ = v767
	var v769 int32
	_ = v769
	var v774 int32
	_ = v774
	var v776 int32
	_ = v776
	var v781 int32
	_ = v781
	var v783 int32
	_ = v783
	var v788 int32
	_ = v788
	var v790 int32
	_ = v790
	var v795 int32
	_ = v795
	var v797 int32
	_ = v797
	var v802 int32
	_ = v802
	var v804 int32
	_ = v804
	var v809 int32
	_ = v809
	var v811 int32
	_ = v811
	var v816 int32
	_ = v816
	var v818 int32
	_ = v818
	var v823 int32
	_ = v823
	var v824 int32
	_ = v824
	var v827 int32
	_ = v827
	var v840 int32
	_ = v840
	var v842 int32
	_ = v842
	var v862 int32
	_ = v862
	var v870 int32
	_ = v870
	var v874 int32
	_ = v874
	var v879 int32
	_ = v879
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = l0
	v15 = int32(-101)
	v16 = l0 + l1
	if base.Ui32(v16) <= base.Ui32(l0) {
		v128 = v15
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v870 = m.ExcPending
	if v870 != 0 {
		goto L168
	} else {
		goto L246
	}
L2:
	;
	m.G0 = v12 + int32(16)
	return v862
L3:
	;
	if v128 <= int32(0) {
		v862 = v15
		goto L2
	} else {
		goto L42
	}
L4:
	;
	goto L3
L5:
	;
	v28 = int32(10)
	goto L7
L7:
	;
	goto L8
L8:
	;
	if v16-l0 < v28 {
		v128 = v15
		goto L4
	} else {
		goto L9
	}
L9:
	;
	v33 = int32(_a_F_pgp_armor_decode_0)
	goto L11
L11:
	;
	goto L12
L12:
	;
	v35 = int32(*(*int8)(unsafe.Add(mBase, _c_F_pgp_armor_decode[0])))
	v39 = l0
	goto L13
L13:
	;
	v46 = F_memchr(m, v39, v35, v16-v39)
	mBase = m.M
	if v46 == int32(0) {
		v128 = v15
		goto L4
	} else {
		goto L15
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12+int32(12)))) = v46
	if base.Ui32(v16) <= base.Ui32(v49) {
		v86 = v49
		goto L26
	} else {
		goto L27
	}
L15:
	;
	v49 = v46 + v28
	if base.Ui32(v16) < base.Ui32(v49) {
		v128 = v15
		goto L4
	} else {
		goto L16
	}
L16:
	;
	v51 = F_memcmp(m, v46, v33, v28)
	mBase = m.M
	if v51 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v53 = v46 + int32(1)
	if base.Ui32(v53) < base.Ui32(v16) {
		v39 = v53
		goto L13
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	if l0 == v46 {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	v128 = v15
	goto L4
L21:
	;
	goto L14
L22:
	;
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46-int32(1)))))
	if v58 == int32(10) {
		goto L21
	} else {
		goto L23
	}
L23:
	;
	if base.Ui32(v16) <= base.Ui32(v49) {
		v128 = v15
		goto L4
	} else {
		goto L24
	}
L24:
	;
	if v28 <= v16-v49 {
		v39 = v49
		goto L13
	} else {
		goto L25
	}
L25:
	;
	v128 = v15
	goto L4
L26:
	;
	if v16-v86 < int32(5) {
		v128 = v15
		goto L4
	} else {
		goto L33
	}
L27:
	;
	v69 = v49
	goto L28
L28:
	;
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69))))
	if v75 == int32(45) {
		v86 = v69
		goto L26
	} else {
		goto L30
	}
L29:
	;
	v86 = v16
	goto L26
L30:
	;
	if base.Ui32(v75) < base.Ui32(int32(32)) {
		v128 = v15
		goto L4
	} else {
		goto L31
	}
L31:
	;
	v81 = v69 + int32(1)
	if base.Ui32(v81) < base.Ui32(v16) {
		v69 = v81
		goto L28
	} else {
		goto L32
	}
L32:
	;
	goto L29
L33:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
	v96 = *(*int32)(unsafe.Add(mBase, _c_F_pgp_armor_decode[0]))
	v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86)+4)))
	v99 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pgp_armor_decode[1])))
	if v95^v96|(v98^v99) != 0 {
		v128 = v15
		goto L4
	} else {
		goto L34
	}
L34:
	;
	v103 = v86 + int32(5)
	if base.Ui32(v16) <= base.Ui32(v103) {
		v118 = v103
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v128 = v118 - v46
	goto L4
L36:
	;
	v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103))))
	switch v105 - int32(10) {
	case 0, 3:
		goto L37
	default:
		v128 = v15
		goto L4
	}
L37:
	;
	if v105 == int32(13) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v112 = v86 + int32(6)
	goto L40
L39:
	;
	v112 = v103
	goto L40
L40:
	;
	if base.Ui32(v16) <= base.Ui32(v112) {
		v118 = v112
		goto L35
	} else {
		goto L41
	}
L41:
	;
	v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v112))))
	v118 = v112 + base.B2i32(v114 == int32(10))
	goto L35
L42:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	v133 = v132 + v128
	v142 = int32(-101)
	if base.Ui32(v16) <= base.Ui32(v133) {
		v245 = v142
		goto L44
	} else {
		goto L45
	}
L43:
	;
	if v245 <= int32(0) {
		v862 = v15
		goto L2
	} else {
		goto L82
	}
L44:
	;
	goto L43
L45:
	;
	v144 = int32(8)
	goto L46
L46:
	;
	goto L48
L48:
	;
	if v16-v133 < v144 {
		v245 = v142
		goto L44
	} else {
		goto L49
	}
L49:
	;
	v149 = int32(_a_F_pgp_armor_decode_1)
	goto L50
L50:
	;
	goto L52
L52:
	;
	v152 = int32(*(*int8)(unsafe.Add(mBase, _c_F_pgp_armor_decode[2])))
	v156 = v133
	goto L53
L53:
	;
	v163 = F_memchr(m, v156, v152, v16-v156)
	mBase = m.M
	if v163 == int32(0) {
		v245 = v142
		goto L44
	} else {
		goto L55
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12+int32(8)))) = v163
	if base.Ui32(v16) <= base.Ui32(v166) {
		v203 = v166
		goto L66
	} else {
		goto L67
	}
L55:
	;
	v166 = v163 + v144
	if base.Ui32(v16) < base.Ui32(v166) {
		v245 = v142
		goto L44
	} else {
		goto L56
	}
L56:
	;
	v168 = F_memcmp(m, v163, v149, v144)
	mBase = m.M
	if v168 != 0 {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v170 = v163 + int32(1)
	if base.Ui32(v170) < base.Ui32(v16) {
		v156 = v170
		goto L53
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	if v133 == v163 {
		goto L61
	} else {
		goto L62
	}
L60:
	;
	v245 = v142
	goto L44
L61:
	;
	goto L54
L62:
	;
	v175 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v163-int32(1)))))
	if v175 == int32(10) {
		goto L61
	} else {
		goto L63
	}
L63:
	;
	if base.Ui32(v16) <= base.Ui32(v166) {
		v245 = v142
		goto L44
	} else {
		goto L64
	}
L64:
	;
	if v144 <= v16-v166 {
		v156 = v166
		goto L53
	} else {
		goto L65
	}
L65:
	;
	v245 = v142
	goto L44
L66:
	;
	if v16-v203 < int32(5) {
		v245 = v142
		goto L44
	} else {
		goto L73
	}
L67:
	;
	v186 = v166
	goto L68
L68:
	;
	v192 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v186))))
	if v192 == int32(45) {
		v203 = v186
		goto L66
	} else {
		goto L70
	}
L69:
	;
	v203 = v16
	goto L66
L70:
	;
	if base.Ui32(v192) < base.Ui32(int32(32)) {
		v245 = v142
		goto L44
	} else {
		goto L71
	}
L71:
	;
	v198 = v186 + int32(1)
	if base.Ui32(v198) < base.Ui32(v16) {
		v186 = v198
		goto L68
	} else {
		goto L72
	}
L72:
	;
	goto L69
L73:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v203)))
	v213 = *(*int32)(unsafe.Add(mBase, _c_F_pgp_armor_decode[2]))
	v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v203)+4)))
	v216 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pgp_armor_decode[3])))
	if v212^v213|(v215^v216) != 0 {
		v245 = v142
		goto L44
	} else {
		goto L74
	}
L74:
	;
	v220 = v203 + int32(5)
	if base.Ui32(v16) <= base.Ui32(v220) {
		v235 = v220
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v245 = v235 - v163
	goto L44
L76:
	;
	v222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v220))))
	switch v222 - int32(10) {
	case 0, 3:
		goto L77
	default:
		v245 = v142
		goto L44
	}
L77:
	;
	if v222 == int32(13) {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v229 = v203 + int32(6)
	goto L80
L79:
	;
	v229 = v220
	goto L80
L80:
	;
	if base.Ui32(v16) <= base.Ui32(v229) {
		v235 = v229
		goto L75
	} else {
		goto L81
	}
L81:
	;
	v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v229))))
	v235 = v229 + base.B2i32(v231 == int32(10))
	goto L75
L82:
	;
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
	if base.Ui32(v249) <= base.Ui32(v133) {
		v375 = v133
		goto L83
	} else {
		goto L84
	}
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v249
	if base.Ui32(v249) < base.Ui32(v375) {
		v862 = v15
		goto L2
	} else {
		goto L115
	}
L84:
	;
	v251 = v133
	goto L85
L85:
	;
	v260 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v251))))
	switch v260 - int32(10) {
	case 0, 3:
		v375 = v251
		goto L83
	default:
		goto L87
	}
L86:
	;
	v375 = v373
	goto L83
L87:
	;
	v264 = v249 - v251
	v265 = int32(0)
	if base.B2i32(v251&int32(3) == v265)|base.B2i32(v264 == v265) != 0 {
		v295 = v251
		v297 = v264
		v298 = base.B2i32(v264 != v265)
		goto L91
	} else {
		goto L92
	}
L88:
	;
	if v369 == int32(0) {
		v862 = v15
		goto L2
	} else {
		goto L113
	}
L89:
	;
	v369 = int32(0)
	goto L88
L90:
	;
	v347 = v340
	v349 = v342
	goto L107
L91:
	;
	if v298 == int32(0) {
		goto L89
	} else {
		goto L98
	}
L92:
	;
	v278 = v251
	v280 = v264
	goto L93
L93:
	;
	v283 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v278))))
	if v283 == int32(10) {
		v340 = v278
		v342 = v280
		goto L90
	} else {
		goto L95
	}
L94:
	;
	v295 = v290
	v297 = v286
	v298 = v288
	goto L91
L95:
	;
	v285 = int32(1)
	v286 = v280 - v285
	v287 = int32(0)
	v288 = base.B2i32(v286 != v287)
	v290 = v278 + v285
	if v290&int32(3) == v287 {
		v295 = v290
		v297 = v286
		v298 = v288
		goto L91
	} else {
		goto L96
	}
L96:
	;
	if v286 != 0 {
		v278 = v290
		v280 = v286
		goto L93
	} else {
		goto L97
	}
L97:
	;
	goto L94
L98:
	;
	v304 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v295))))
	if base.B2i32(int32(10) == v304)|base.B2i32(base.Ui32(v297) < base.Ui32(int32(4))) == int32(0) {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	v313 = v295
	v315 = v297
	goto L102
L100:
	;
	v333 = v295
	v335 = v297
	goto L101
L101:
	;
	if v335 == int32(0) {
		goto L89
	} else {
		goto L106
	}
L102:
	;
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v313)))
	v320 = v319 ^ int32(168430090)
	v323 = int32(-2139062144)
	if (int32(16843008)-v320|v320)&v323 != v323 {
		v340 = v313
		v342 = v315
		goto L90
	} else {
		goto L104
	}
L103:
	;
	v333 = v328
	v335 = v330
	goto L101
L104:
	;
	v327 = int32(4)
	v328 = v313 + v327
	v330 = v315 - v327
	if base.Ui32(int32(3)) < base.Ui32(v330) {
		v313 = v328
		v315 = v330
		goto L102
	} else {
		goto L105
	}
L105:
	;
	goto L103
L106:
	;
	v340 = v333
	v342 = v335
	goto L90
L107:
	;
	v352 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v347))))
	if int32(10) == v352 {
		goto L109
	} else {
		goto L110
	}
L108:
	;
	goto L89
L109:
	;
	v369 = v347
	goto L88
L110:
	;
	goto L111
L111:
	;
	v354 = int32(1)
	v357 = v349 - v354
	if v357 != 0 {
		v347 = v347 + v354
		v349 = v357
		goto L107
	} else {
		goto L112
	}
L112:
	;
	goto L108
L113:
	;
	v373 = v369 + int32(1)
	if base.Ui32(v373) < base.Ui32(v249) {
		v251 = v373
		goto L85
	} else {
		goto L114
	}
L114:
	;
	goto L86
L115:
	;
	v386 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v249))))
	if v386 == int32(61) {
		goto L117
	} else {
		goto L118
	}
L116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v409
	v420 = v12 + int32(4)
	v421 = int32(0)
	goto L128
L117:
	;
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	v409 = v389
	v412 = v249
	goto L116
L118:
	;
	goto L119
L119:
	;
	v393 = v249
	goto L120
L120:
	;
	v400 = v393 - int32(1)
	if base.Ui32(v375) <= base.Ui32(v400) {
		goto L122
	} else {
		goto L123
	}
L121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v400
	v862 = v15
	goto L2
L122:
	;
	v402 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v400))))
	if v402 != int32(61) {
		v393 = v400
		goto L120
	} else {
		goto L125
	}
L123:
	;
	goto L124
L124:
	;
	goto L121
L125:
	;
	v409 = v400
	v412 = v400
	goto L116
L126:
	;
	if v572 != int32(3) {
		v862 = v15
		goto L2
	} else {
		goto L167
	}
L127:
	;
	goto L126
L128:
	;
	v428 = v412 + int32(5)
	v429 = v412 + int32(1)
	v432 = v420
	v433 = v421
	v434 = v421
	v436 = v421
	goto L131
L130:
	;
	v572 = v547 - v420
	goto L127
L131:
	;
	v440 = v429
	goto L139
L132:
	;
	if v550 != 0 {
		v572 = int32(-101)
		goto L127
	} else {
		goto L166
	}
L133:
	;
	goto L132
L134:
	;
	if base.Ui32(v450) < base.Ui32(v428) {
		v429 = v450
		v432 = v540
		v433 = v541
		v434 = v542
		v436 = v543
		goto L131
	} else {
		goto L165
	}
L135:
	;
	v540 = v533
	v541 = v534
	v542 = v535
	v543 = int32(0)
	goto L134
L136:
	;
	v533 = v530
	v534 = int32(0)
	v535 = v527
	goto L135
L137:
	;
	v502 = v498 + v433<<(uint(int32(6))%32)
	v504 = v436 + int32(1)
	if v504 != int32(4) {
		goto L156
	} else {
		goto L157
	}
L138:
	;
	v498 = int32(62)
	goto L137
L139:
	;
	v450 = v440 + int32(1)
	v451 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v440))))
	v453 = v451 - int32(65)
	if base.Ui32(v453&int32(255)) <= base.Ui32(int32(25)) {
		v498 = v453
		goto L137
	} else {
		goto L141
	}
L140:
	;
	if v434 != 0 {
		goto L151
	} else {
		goto L152
	}
L141:
	;
	if base.Ui32((v451-int32(97))&int32(255)) <= base.Ui32(int32(25)) {
		goto L142
	} else {
		goto L143
	}
L142:
	;
	v498 = v451 - int32(71)
	goto L137
L143:
	;
	goto L144
L144:
	;
	if base.Ui32((v451-int32(48))&int32(255)) <= base.Ui32(int32(9)) {
		goto L145
	} else {
		goto L146
	}
L145:
	;
	v498 = (v451 + int32(4)) & int32(255)
	goto L137
L146:
	;
	goto L147
L147:
	;
	v477 = int32(-101)
	switch v451 - int32(9) {
	case 0, 1, 4, 23:
		goto L149
	default:
		v572 = v477
		goto L127
	case 34:
		goto L138
	case 38:
		v498 = int32(63)
		goto L137
	case 52:
		goto L148
	}
L148:
	;
	goto L140
L149:
	;
	if base.Ui32(v450) < base.Ui32(v428) {
		v440 = v450
		goto L139
	} else {
		goto L150
	}
L150:
	;
	v547 = v432
	v550 = v436
	goto L133
L151:
	;
	v498 = int32(0)
	goto L137
L152:
	;
	goto L153
L153:
	;
	switch v436 - int32(2) {
	case 0:
		goto L155
	case 1:
		goto L154
	default:
		v572 = v477
		goto L127
	}
L154:
	;
	v488 = int32(2)
	v490 = int32(base.Ui32(v433) >> (uint(v488) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v432)+1)) = uint8(v490)
	v493 = int32(base.Ui32(v433) >> (uint(int32(10)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v432))) = uint8(v493)
	v527 = v488
	v530 = v432 + v488
	goto L136
L155:
	;
	v540 = v432
	v541 = v433 << (uint(int32(6)) % 32)
	v542 = int32(1)
	v543 = int32(3)
	goto L134
L156:
	;
	v540 = v432
	v541 = v502
	v542 = v434
	v543 = v504
	goto L134
L157:
	;
	goto L158
L158:
	;
	v508 = int32(base.Ui32(v502) >> (uint(int32(16)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v432))) = uint8(v508)
	v510 = int32(0)
	if v434 == int32(1) {
		goto L159
	} else {
		goto L160
	}
L159:
	;
	v513 = int32(1)
	v533 = v432 + v513
	v534 = v510
	v535 = v513
	goto L135
L160:
	;
	goto L161
L161:
	;
	v517 = int32(base.Ui32(v502) >> (uint(int32(8)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v432)+1)) = uint8(v517)
	if v434 != 0 {
		goto L162
	} else {
		goto L163
	}
L162:
	;
	v533 = v432 + int32(2)
	v534 = v510
	v535 = v434
	goto L135
L163:
	;
	goto L164
L164:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v432)+2)) = uint8(v502)
	v527 = int32(0)
	v530 = v432 + int32(3)
	goto L136
L165:
	;
	v547 = v540
	v550 = v543
	goto L133
L166:
	;
	goto L130
L167:
	;
	v576 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+6)))
	v577 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+5)))
	v578 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+4)))
	v582 = int32(base.Ui32(l1*int32(3)) >> (uint(int32(2)) % 32))
	F_enlargeStringInfo(m, l2, v582)
	mBase = m.M
	v586 = m.ExcPending
	if v586 != 0 {
		goto L168
	} else {
		goto L169
	}
L168:
	;
	return int32(0)
L169:
	;
	v589 = v375 ^ int32(-1) + v412
	v590 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v591 = int32(0)
	if v589 != 0 {
		goto L172
	} else {
		goto L173
	}
L170:
	;
	if v582 < v742 {
		goto L1
	} else {
		goto L211
	}
L171:
	;
	goto L170
L172:
	;
	v598 = v375 + v589
	v599 = v375
	v602 = v590
	v603 = v591
	v604 = v591
	v606 = v591
	goto L175
L173:
	;
	v726 = v590
	goto L174
L174:
	;
	v742 = v726 - v590
	goto L171
L175:
	;
	v610 = v599
	goto L183
L176:
	;
	if v720 != 0 {
		v742 = int32(-101)
		goto L171
	} else {
		goto L210
	}
L177:
	;
	goto L176
L178:
	;
	if base.Ui32(v620) < base.Ui32(v598) {
		v599 = v620
		v602 = v710
		v603 = v711
		v604 = v712
		v606 = v713
		goto L175
	} else {
		goto L209
	}
L179:
	;
	v710 = v703
	v711 = v704
	v712 = v705
	v713 = int32(0)
	goto L178
L180:
	;
	v703 = v700
	v704 = int32(0)
	v705 = v697
	goto L179
L181:
	;
	v672 = v668 + v603<<(uint(int32(6))%32)
	v674 = v606 + int32(1)
	if v674 != int32(4) {
		goto L200
	} else {
		goto L201
	}
L182:
	;
	v668 = int32(62)
	goto L181
L183:
	;
	v620 = v610 + int32(1)
	v621 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v610))))
	v623 = v621 - int32(65)
	if base.Ui32(v623&int32(255)) <= base.Ui32(int32(25)) {
		v668 = v623
		goto L181
	} else {
		goto L185
	}
L184:
	;
	if v604 != 0 {
		goto L195
	} else {
		goto L196
	}
L185:
	;
	if base.Ui32((v621-int32(97))&int32(255)) <= base.Ui32(int32(25)) {
		goto L186
	} else {
		goto L187
	}
L186:
	;
	v668 = v621 - int32(71)
	goto L181
L187:
	;
	goto L188
L188:
	;
	if base.Ui32((v621-int32(48))&int32(255)) <= base.Ui32(int32(9)) {
		goto L189
	} else {
		goto L190
	}
L189:
	;
	v668 = (v621 + int32(4)) & int32(255)
	goto L181
L190:
	;
	goto L191
L191:
	;
	v647 = int32(-101)
	switch v621 - int32(9) {
	case 0, 1, 4, 23:
		goto L193
	default:
		v742 = v647
		goto L171
	case 34:
		goto L182
	case 38:
		v668 = int32(63)
		goto L181
	case 52:
		goto L192
	}
L192:
	;
	goto L184
L193:
	;
	if base.Ui32(v620) < base.Ui32(v598) {
		v610 = v620
		goto L183
	} else {
		goto L194
	}
L194:
	;
	v717 = v602
	v720 = v606
	goto L177
L195:
	;
	v668 = int32(0)
	goto L181
L196:
	;
	goto L197
L197:
	;
	switch v606 - int32(2) {
	case 0:
		goto L199
	case 1:
		goto L198
	default:
		v742 = v647
		goto L171
	}
L198:
	;
	v658 = int32(2)
	v660 = int32(base.Ui32(v603) >> (uint(v658) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v602)+1)) = uint8(v660)
	v663 = int32(base.Ui32(v603) >> (uint(int32(10)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v602))) = uint8(v663)
	v697 = v658
	v700 = v602 + v658
	goto L180
L199:
	;
	v710 = v602
	v711 = v603 << (uint(int32(6)) % 32)
	v712 = int32(1)
	v713 = int32(3)
	goto L178
L200:
	;
	v710 = v602
	v711 = v672
	v712 = v604
	v713 = v674
	goto L178
L201:
	;
	goto L202
L202:
	;
	v678 = int32(base.Ui32(v672) >> (uint(int32(16)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v602))) = uint8(v678)
	v680 = int32(0)
	if v604 == int32(1) {
		goto L203
	} else {
		goto L204
	}
L203:
	;
	v683 = int32(1)
	v703 = v602 + v683
	v704 = v680
	v705 = v683
	goto L179
L204:
	;
	goto L205
L205:
	;
	v687 = int32(base.Ui32(v672) >> (uint(int32(8)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v602)+1)) = uint8(v687)
	if v604 != 0 {
		goto L206
	} else {
		goto L207
	}
L206:
	;
	v703 = v602 + int32(2)
	v704 = v680
	v705 = v604
	goto L179
L207:
	;
	goto L208
L208:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v602)+2)) = uint8(v672)
	v697 = int32(0)
	v700 = v602 + int32(3)
	goto L180
L209:
	;
	v717 = v710
	v720 = v713
	goto L177
L210:
	;
	v726 = v717
	goto L174
L211:
	;
	if int32(0) <= v742 {
		goto L212
	} else {
		goto L213
	}
L212:
	;
	if v742 != 0 {
		goto L215
	} else {
		goto L216
	}
L213:
	;
	goto L214
L214:
	;
	v862 = v742
	goto L2
L215:
	;
	v753 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v755 = v753
	v758 = int32(_a_F_pgp_armor_decode_2)
	v761 = v742
	goto L218
L216:
	;
	v840 = int32(_a_F_pgp_armor_decode_2)
	goto L217
L217:
	;
	if v840 != v577<<(uint(int32(8))%32)|v578<<(uint(int32(16))%32)|v576 {
		v862 = v15
		goto L2
	} else {
		goto L245
	}
L218:
	;
	v764 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v755))))
	v767 = v764<<(uint(int32(16))%32) ^ v758
	v769 = v767 << (uint(int32(1)) % 32)
	if v767&int32(_a_F_pgp_armor_decode_3) != 0 {
		goto L220
	} else {
		goto L221
	}
L219:
	;
	v840 = v823 & int32(16777215)
	goto L217
L220:
	;
	v774 = v769 ^ int32(25578747)
	goto L222
L221:
	;
	v774 = v769
	goto L222
L222:
	;
	v776 = v774 << (uint(int32(1)) % 32)
	if v774&int32(_a_F_pgp_armor_decode_3) != 0 {
		goto L223
	} else {
		goto L224
	}
L223:
	;
	v781 = v776 ^ int32(25578747)
	goto L225
L224:
	;
	v781 = v776
	goto L225
L225:
	;
	v783 = v781 << (uint(int32(1)) % 32)
	if v781&int32(_a_F_pgp_armor_decode_3) != 0 {
		goto L226
	} else {
		goto L227
	}
L226:
	;
	v788 = v783 ^ int32(25578747)
	goto L228
L227:
	;
	v788 = v783
	goto L228
L228:
	;
	v790 = v788 << (uint(int32(1)) % 32)
	if v788&int32(_a_F_pgp_armor_decode_3) != 0 {
		goto L229
	} else {
		goto L230
	}
L229:
	;
	v795 = v790 ^ int32(25578747)
	goto L231
L230:
	;
	v795 = v790
	goto L231
L231:
	;
	v797 = v795 << (uint(int32(1)) % 32)
	if v795&int32(_a_F_pgp_armor_decode_3) != 0 {
		goto L232
	} else {
		goto L233
	}
L232:
	;
	v802 = v797 ^ int32(25578747)
	goto L234
L233:
	;
	v802 = v797
	goto L234
L234:
	;
	v804 = v802 << (uint(int32(1)) % 32)
	if v802&int32(_a_F_pgp_armor_decode_3) != 0 {
		goto L235
	} else {
		goto L236
	}
L235:
	;
	v809 = v804 ^ int32(25578747)
	goto L237
L236:
	;
	v809 = v804
	goto L237
L237:
	;
	v811 = v809 << (uint(int32(1)) % 32)
	if v809&int32(_a_F_pgp_armor_decode_3) != 0 {
		goto L238
	} else {
		goto L239
	}
L238:
	;
	v816 = v811 ^ int32(25578747)
	goto L240
L239:
	;
	v816 = v811
	goto L240
L240:
	;
	v818 = v816 << (uint(int32(1)) % 32)
	if v816&int32(_a_F_pgp_armor_decode_3) != 0 {
		goto L241
	} else {
		goto L242
	}
L241:
	;
	v823 = v818 ^ int32(25578747)
	goto L243
L242:
	;
	v823 = v818
	goto L243
L243:
	;
	v824 = int32(1)
	v827 = v761 - v824
	if v827 != 0 {
		v755 = v755 + v824
		v758 = v823
		v761 = v827
		goto L218
	} else {
		goto L244
	}
L244:
	;
	goto L219
L245:
	;
	v842 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v842 + v742
	goto L214
L246:
	;
	F_errmsg_internal(m, int32(_a_F_pgp_armor_decode_4), int32(0))
	mBase = m.M
	v874 = m.ExcPending
	if v874 != 0 {
		goto L168
	} else {
		goto L247
	}
L247:
	;
	F_errfinish(m, int32(_a_F_pgp_armor_decode_5), int32(370), int32(_a_F_pgp_armor_decode_6))
	mBase = m.M
	v879 = m.ExcPending
	if v879 != 0 {
		goto L168
	} else {
		goto L248
	}
L248:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pgp_armor_encode(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v101 int32
	_ = v101
	var v107 int32
	_ = v107
	var v117 int32
	_ = v117
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v201 int32
	_ = v201
	var v209 int32
	_ = v209
	var v217 int32
	_ = v217
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v244 int32
	_ = v244
	var v252 int32
	_ = v252
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v276 int32
	_ = v276
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v290 int32
	_ = v290
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v333 int32
	_ = v333
	var v340 int32
	_ = v340
	var v344 int32
	_ = v344
	var v349 int32
	_ = v349
	v12 = m.G0
	v14 = v12 - int32(16)
	m.G0 = v14
	v16 = int32(_a_F_pgp_armor_encode_0)
	if l1 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v23 = l0
	v24 = v16
	v26 = l1
	goto L4
L2:
	;
	v101 = v16
	goto L3
L3:
	;
	F_appendStringInfoString(m, l2, int32(_a_F_pgp_armor_encode_1))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L31
	} else {
		goto L32
	}
L4:
	;
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23))))
	v31 = v28<<(uint(int32(16))%32) ^ v24
	v33 = v31 << (uint(int32(1)) % 32)
	if v31&int32(_a_F_pgp_armor_encode_2) != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v101 = v87 & int32(16777215)
	goto L3
L6:
	;
	v38 = v33 ^ int32(25578747)
	goto L8
L7:
	;
	v38 = v33
	goto L8
L8:
	;
	v40 = v38 << (uint(int32(1)) % 32)
	if v38&int32(_a_F_pgp_armor_encode_2) != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v45 = v40 ^ int32(25578747)
	goto L11
L10:
	;
	v45 = v40
	goto L11
L11:
	;
	v47 = v45 << (uint(int32(1)) % 32)
	if v45&int32(_a_F_pgp_armor_encode_2) != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v52 = v47 ^ int32(25578747)
	goto L14
L13:
	;
	v52 = v47
	goto L14
L14:
	;
	v54 = v52 << (uint(int32(1)) % 32)
	if v52&int32(_a_F_pgp_armor_encode_2) != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v59 = v54 ^ int32(25578747)
	goto L17
L16:
	;
	v59 = v54
	goto L17
L17:
	;
	v61 = v59 << (uint(int32(1)) % 32)
	if v59&int32(_a_F_pgp_armor_encode_2) != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v66 = v61 ^ int32(25578747)
	goto L20
L19:
	;
	v66 = v61
	goto L20
L20:
	;
	v68 = v66 << (uint(int32(1)) % 32)
	if v66&int32(_a_F_pgp_armor_encode_2) != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v73 = v68 ^ int32(25578747)
	goto L23
L22:
	;
	v73 = v68
	goto L23
L23:
	;
	v75 = v73 << (uint(int32(1)) % 32)
	if v73&int32(_a_F_pgp_armor_encode_2) != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v80 = v75 ^ int32(25578747)
	goto L26
L25:
	;
	v80 = v75
	goto L26
L26:
	;
	v82 = v80 << (uint(int32(1)) % 32)
	if v80&int32(_a_F_pgp_armor_encode_2) != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v87 = v82 ^ int32(25578747)
	goto L29
L28:
	;
	v87 = v82
	goto L29
L29:
	;
	v88 = int32(1)
	v91 = v26 - v88
	if v91 != 0 {
		v23 = v23 + v88
		v24 = v87
		v26 = v91
		goto L4
	} else {
		goto L30
	}
L30:
	;
	goto L5
L31:
	;
	return
L32:
	;
	if int32(0) < l3 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v117 = int32(0)
	goto L36
L34:
	;
	goto L35
L35:
	;
	F_appendStringInfoChar(m, l2, int32(10))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L31
	} else {
		goto L40
	}
L36:
	;
	v123 = v117 << (uint(int32(2)) % 32)
	v125 = *(*int32)(unsafe.Add(mBase, uint32(l4+v123)))
	v127 = *(*int32)(unsafe.Add(mBase, uint32(l5+v123)))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v127
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v125
	F_appendStringInfo(m, l2, int32(_a_F_pgp_armor_encode_3), v14)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L31
	} else {
		goto L38
	}
L37:
	;
	goto L35
L38:
	;
	v134 = v117 + int32(1)
	if v134 != l3 {
		v117 = v134
		goto L36
	} else {
		goto L39
	}
L39:
	;
	goto L37
L40:
	;
	v150 = int32(2)
	v152 = base.I32_div_u_s(l1, int32(57))
	v156 = base.I32_div_u_s(l1+v150, int32(3))
	v159 = v152 + v156<<(uint(v150)%32)
	F_enlargeStringInfo(m, l2, v159)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L31
	} else {
		goto L41
	}
L41:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v163 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v164 = v162 + v163
	if l1 == int32(0) {
		v276 = v164
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v281 = v276 - v164
	if base.Ui32(v281) <= base.Ui32(v159) {
		goto L57
	} else {
		goto L58
	}
L43:
	;
	v171 = l0
	v172 = int32(0)
	v174 = v164 + int32(76)
	v177 = v164
	v180 = v150
	goto L44
L44:
	;
	v182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v171))))
	v186 = v182<<(uint(v180<<(uint(int32(3))%32))%32) | v172
	if int32(0) < v180 {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	if v225 == int32(2) {
		v276 = v234
		goto L42
	} else {
		goto L53
	}
L46:
	;
	v223 = v186
	v224 = v177
	v225 = v180 - int32(1)
	goto L48
L47:
	;
	v191 = int32(63)
	v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v186&v191)+uint32(_c_F_pgp_armor_encode[0]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v177)+3)) = uint8(v195)
	v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v186)>>(uint(int32(18))%32)))+uint32(_c_F_pgp_armor_encode[0]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v177))) = uint8(v201)
	v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v186)>>(uint(int32(6))%32))&v191)+uint32(_c_F_pgp_armor_encode[0]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v177)+2)) = uint8(v209)
	v217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v186)>>(uint(int32(12))%32))&v191)+uint32(_c_F_pgp_armor_encode[0]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v177)+1)) = uint8(v217)
	v223 = int32(0)
	v224 = v177 + int32(4)
	v225 = int32(2)
	goto L48
L48:
	;
	if base.Ui32(v174) <= base.Ui32(v224) {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v227 = int32(10)
	*(*uint8)(unsafe.Add(mBase, uint32(v224))) = uint8(v227)
	v233 = v224 + int32(77)
	v234 = v224 + int32(1)
	goto L51
L50:
	;
	v233 = v174
	v234 = v224
	goto L51
L51:
	;
	v236 = v171 + int32(1)
	if base.Ui32(v236) < base.Ui32(l0+l1) {
		v171 = v236
		v172 = v223
		v174 = v233
		v177 = v234
		v180 = v225
		goto L44
	} else {
		goto L52
	}
L52:
	;
	goto L45
L53:
	;
	v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v223)>>(uint(int32(18))%32)))+uint32(_c_F_pgp_armor_encode[0]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v234))) = uint8(v244)
	v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v223)>>(uint(int32(12))%32))&int32(63))+uint32(_c_F_pgp_armor_encode[0]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v234)+1)) = uint8(v252)
	if v225 == int32(0) {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v263 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v223)>>(uint(int32(6))%32))&int32(63))+uint32(_c_F_pgp_armor_encode[0]))))
	v264 = v263
	goto L56
L55:
	;
	v264 = int32(61)
	goto L56
L56:
	;
	v265 = int32(61)
	*(*uint8)(unsafe.Add(mBase, uint32(v234)+3)) = uint8(v265)
	*(*uint8)(unsafe.Add(mBase, uint32(v234)+2)) = uint8(v264)
	v276 = v234 + int32(4)
	goto L42
L57:
	;
	v283 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v284 = v283 + v281
	*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v284
	v286 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v290 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v286+v284-int32(1)))))
	if v290 != int32(10) {
		goto L60
	} else {
		goto L61
	}
L58:
	;
	goto L59
L59:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L31
	} else {
		goto L70
	}
L60:
	;
	F_appendStringInfoChar(m, l2, int32(10))
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L31
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	F_appendStringInfoChar(m, l2, int32(61))
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L31
	} else {
		goto L64
	}
L63:
	;
	goto L62
L64:
	;
	v303 = int32(*(*int8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v101)>>(uint(int32(18))%32)))+uint32(_c_F_pgp_armor_encode[0]))))
	F_appendStringInfoChar(m, l2, v303)
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L31
	} else {
		goto L65
	}
L65:
	;
	v312 = int32(*(*int8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v101)>>(uint(int32(12))%32))&int32(63))+uint32(_c_F_pgp_armor_encode[0]))))
	F_appendStringInfoChar(m, l2, v312)
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L31
	} else {
		goto L66
	}
L66:
	;
	v321 = int32(*(*int8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v101)>>(uint(int32(6))%32))&int32(63))+uint32(_c_F_pgp_armor_encode[0]))))
	F_appendStringInfoChar(m, l2, v321)
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L31
	} else {
		goto L67
	}
L67:
	;
	v328 = int32(*(*int8)(unsafe.Add(mBase, uint32(v101&int32(63))+uint32(_c_F_pgp_armor_encode[0]))))
	F_appendStringInfoChar(m, l2, v328)
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L31
	} else {
		goto L68
	}
L68:
	;
	F_appendStringInfoString(m, l2, int32(_a_F_pgp_armor_encode_4))
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L31
	} else {
		goto L69
	}
L69:
	;
	m.G0 = v14 + int32(16)
	return
L70:
	;
	F_errmsg_internal(m, int32(_a_F_pgp_armor_encode_5), int32(0))
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L31
	} else {
		goto L71
	}
L71:
	;
	F_errfinish(m, int32(_a_F_pgp_armor_encode_6), int32(227), int32(_a_F_pgp_armor_encode_7))
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L31
	} else {
		goto L72
	}
L72:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pgp_cfb_create(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
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
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v54 int32
	_ = v54
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v14 = F_pgp_load_cipher(m, l1, v10+int32(12))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		if v14 < int32(0) {
			v54 = v14
			m.G0 = v10 + int32(16)
			return v54
		} else {
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
			v22 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
			v23 = m.T0[v22].(func(*base.Module, int32, int32, int32, int32) int32)(m, v20, l2, l3, int32(0))
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int32(0)
			} else {
				if v23 < int32(0) {
					v27 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
					v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+24))
					m.T0[v28].(func(*base.Module, int32))(m, v27)
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
						return int32(0)
					} else {
						v54 = v23
						m.G0 = v10 + int32(16)
						return v54
					}
				} else {
					v32 = F_palloc0(m, int32(116))
					mBase = m.M
					v33 = m.ExcPending
					if v33 != 0 {
						return int32(0)
					} else {
						v34 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
						*(*int32)(unsafe.Add(mBase, uint32(v32))) = v34
						v36 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
						v37 = m.T0[v36].(func(*base.Module, int32) int32)(m, v34)
						mBase = m.M
						v38 = m.ExcPending
						if v38 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v32)+16)) = l4
							*(*int32)(unsafe.Add(mBase, uint32(v32)+4)) = v37
							v41 = int32(0)
							if base.B2i32(l5 == v41)|base.B2i32(v37 == v41) == v41 {
								base.MemoryCopy(m, v32+int32(20), l5, v37)
							} else {
							}
							*(*int32)(unsafe.Add(mBase, uint32(l0))) = v32
							v54 = int32(0)
							m.G0 = v10 + int32(16)
							return v54
						}
					}
				}
			}
		}
	}
}
func F_pgp_create_pkt_writer(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v10 = l1 | int32(192)
	*(*uint8)(unsafe.Add(mBase, uint32(v7)+15)) = uint8(v10)
	v15 = F_pushf_write(m, l0, v7+int32(15), int32(1))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		if int32(0) <= v15 {
			v23 = F_pushf_create(m, l2, int32(_a_F_pgp_create_pkt_writer_0), int32(0), l0)
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int32(0)
			} else {
				v25 = v23
				m.G0 = v7 + int32(16)
				return v25
			}
		} else {
			v25 = v15
			m.G0 = v7 + int32(16)
			return v25
		}
	}
}
func F_pgp_decompress_filter(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v5 = F_pullf_create(m, l0, int32(_a_F_pgp_decompress_filter_0), l1, l2)
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
func F_pgp_free(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	if v3 != 0 {
		F_pgp_key_free(m, v3)
		mBase = m.M
		v7 = m.ExcPending
		if v7 != 0 {
			return int32(0)
		} else {
			base.MemoryFill(m, l0, int32(0), int32(168))
			F_pfree(m, l0)
			mBase = m.M
			v12 = m.ExcPending
			if v12 != 0 {
				return int32(0)
			} else {
				return int32(0)
			}
		}
	} else {
		base.MemoryFill(m, l0, int32(0), int32(168))
		F_pfree(m, l0)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			return int32(0)
		}
	}
}
func F_pgp_get_cipher_block_size(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	v4 = l0 - int32(2)
	if base.B2i32(base.Ui32(int32(8)) < base.Ui32(v4))|base.B2i32(int32(base.Ui32(int32(487))>>(uint(v4)%32))&int32(1) == int32(0)) != 0 {
		v21 = int32(0)
	} else {
		v19 = *(*int32)(unsafe.Add(mBase, uint32(v4<<(uint(int32(2))%32))+uint32(_c_F_pgp_get_cipher_block_size[0])))
		v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+16))
		v21 = v20
	}
	return v21
}
func F_pgp_load_cipher(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	v6 = l0 - int32(2)
	if base.B2i32(base.Ui32(int32(8)) < base.Ui32(v6))|base.B2i32(int32(base.Ui32(int32(487))>>(uint(v6)%32))&int32(1) == int32(0)) != 0 {
		v30 = int32(-100)
		return v30
	} else {
		v23 = *(*int32)(unsafe.Add(mBase, uint32(v6<<(uint(int32(2))%32))+uint32(_c_F_pgp_load_cipher[0])))
		v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
		v25 = F_px_find_cipher(m, v24, l1)
		mBase = m.M
		v28 = m.ExcPending
		if v28 != 0 {
			return int32(0)
		} else {
			if v25 != 0 {
				v29 = int32(-103)
			} else {
				v29 = int32(0)
			}
			v30 = v29
			return v30
		}
	}
}
func F_pgp_mpi_alloc(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v36 int32
	_ = v36
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	if base.Ui32(int32(_a_F_pgp_mpi_alloc_0)) <= base.Ui32(l0) {
		*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
		F_px_debug(m, int32(_a_F_pgp_mpi_alloc_1), v8)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			v36 = int32(-100)
			m.G0 = v8 + int32(16)
			return v36
		}
	} else {
		v22 = int32(base.Ui32(l0+int32(7)) >> (uint(int32(3)) % 32))
		v25 = F_palloc(m, v22+int32(12))
		mBase = m.M
		v26 = m.ExcPending
		if v26 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v25)+8)) = v22
			*(*int32)(unsafe.Add(mBase, uint32(v25)+4)) = l0
			*(*int32)(unsafe.Add(mBase, uint32(v25))) = v25 + int32(12)
			*(*int32)(unsafe.Add(mBase, uint32(l1))) = v25
			v36 = int32(0)
			m.G0 = v8 + int32(16)
			return v36
		}
	}
}
func F_pgp_mpi_cksum(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v9 int32
	_ = v9
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
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
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	v3 = int32(0)
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v15 = l0 + v9>>(uint(int32(8))%32) + v9&int32(255)
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v16 <= v3 {
		v76 = v15
	} else {
		v20 = v16 & int32(3)
		v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
		if base.Ui32(v16) < base.Ui32(int32(4)) {
			v52 = v15
			v53 = int32(0)
			v60 = v52
			v61 = v53
			v67 = v3
			for {
				v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61+v21))))
				v70 = v60 + v69
				v71 = int32(1)
				v74 = v67 + v71
				if v74 != v20 {
					v60 = v70
					v61 = v61 + v71
					v67 = v74
					continue
				} else {
					break
				}
				break
			}
			v76 = v70
		} else {
			v28 = v15
			v29 = int32(0)
			v34 = v3
			for {
				v36 = v29 + v21
				v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36))))
				v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+1)))
				v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+2)))
				v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+3)))
				v44 = v28 + v37 + v39 + v41 + v43
				v45 = int32(4)
				v46 = v29 + v45
				v48 = v34 + v45
				if v48 != v16&int32(2147483644) {
					v28 = v44
					v29 = v46
					v34 = v48
					continue
				} else {
					break
				}
				break
			}
			if v20 == int32(0) {
				v76 = v44
			} else {
				v52 = v44
				v53 = v46
				v60 = v52
				v61 = v53
				v67 = v3
				for {
					v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61+v21))))
					v70 = v60 + v69
					v71 = int32(1)
					v74 = v67 + v71
					if v74 != v20 {
						v60 = v70
						v61 = v61 + v71
						v67 = v74
						continue
					} else {
						break
					}
					break
				}
				v76 = v70
			}
		}
	}
	return v76 & int32(_a_F_pgp_mpi_cksum_0)
}
func F_pgp_parse_pkt_hdr(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v104 int32
	_ = v104
	var v110 int32
	_ = v110
	var v116 int32
	_ = v116
	var v124 int32
	_ = v124
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v15 = F_pullf_read(m, l0, int32(1), v10+int32(4))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		if v15 < int32(0) {
			v124 = v15
			m.G0 = v10 + int32(16)
			return v124
		} else {
			if v15 == int32(0) {
				v124 = int32(0)
				m.G0 = v10 + int32(16)
				return v124
			} else {
				v23 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
				v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23))))
				v25 = base.I32_extend8_s(v24)
				if int32(0) <= v25 {
					F_px_debug(m, int32(_a_F_pgp_parse_pkt_hdr_0), int32(0))
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return int32(0)
					} else {
						v124 = int32(-100)
						m.G0 = v10 + int32(16)
						return v124
					}
				} else {
					if v24&int32(64) != 0 {
						v35 = v25 & int32(63)
						*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v35)
						v37 = F_parse_new_len(m, l0, l2)
						mBase = m.M
						v38 = m.ExcPending
						if v38 != 0 {
							return int32(0)
						} else {
							v124 = v37
							m.G0 = v10 + int32(16)
							return v124
						}
					} else {
						v42 = int32(base.Ui32(v25)>>(uint(int32(2))%32)) & int32(15)
						*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v42)
						v44 = int32(3)
						v45 = v24 & v44
						if v45 == v44 {
							if l3 != 0 {
								v50 = int32(3)
							} else {
								v50 = int32(-100)
							}
							v124 = v50
							m.G0 = v10 + int32(16)
							return v124
						} else {
							v54 = F_pullf_read_fixed(m, l0, int32(1), v10+int32(15))
							mBase = m.M
							v55 = m.ExcPending
							if v55 != 0 {
								return int32(0)
							} else {
								if v54 < int32(0) {
									v124 = v54
									m.G0 = v10 + int32(16)
									return v124
								} else {
									v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)))
									switch v45 - int32(1) {
									case 0:
										v64 = F_pullf_read_fixed(m, l0, int32(1), v10+int32(14))
										mBase = m.M
										v65 = m.ExcPending
										if v65 != 0 {
											return int32(0)
										} else {
											if v64 < int32(0) {
												v124 = v64
											} else {
												v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+14)))
												v116 = v68 | v58<<(uint(int32(8))%32)
												*(*int32)(unsafe.Add(mBase, uint32(l2))) = v116
												v124 = int32(1)
											}
											m.G0 = v10 + int32(16)
											return v124
										}
									case 1:
										v75 = F_pullf_read_fixed(m, l0, int32(1), v10+int32(13))
										mBase = m.M
										v76 = m.ExcPending
										if v76 != 0 {
											return int32(0)
										} else {
											if v75 < int32(0) {
												v124 = v75
												m.G0 = v10 + int32(16)
												return v124
											} else {
												v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+13)))
												v83 = F_pullf_read_fixed(m, l0, int32(1), v10+int32(12))
												mBase = m.M
												v84 = m.ExcPending
												if v84 != 0 {
													return int32(0)
												} else {
													if v83 < int32(0) {
														v124 = v83
														m.G0 = v10 + int32(16)
														return v124
													} else {
														v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+12)))
														v91 = F_pullf_read_fixed(m, l0, int32(1), v10+int32(11))
														mBase = m.M
														v92 = m.ExcPending
														if v92 != 0 {
															return int32(0)
														} else {
															if v91 < int32(0) {
																v124 = v91
																m.G0 = v10 + int32(16)
																return v124
															} else {
																v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+11)))
																v96 = int32(8)
																v104 = v95 | (v79<<(uint(v96)%32)|v58<<(uint(int32(16))%32)|v87)<<(uint(v96)%32)
																if base.Ui32(v104) < base.Ui32(int32(16777217)) {
																	v116 = v104
																	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v116
																	v124 = int32(1)
																	m.G0 = v10 + int32(16)
																	return v124
																} else {
																	F_px_debug(m, int32(_a_F_pgp_parse_pkt_hdr_1), int32(0))
																	mBase = m.M
																	v110 = m.ExcPending
																	if v110 != 0 {
																		return int32(0)
																	} else {
																		v124 = int32(-100)
																		m.G0 = v10 + int32(16)
																		return v124
																	}
																}
															}
														}
													}
												}
											}
										}
									default:
										v116 = v58
										*(*int32)(unsafe.Add(mBase, uint32(l2))) = v116
										v124 = int32(1)
										m.G0 = v10 + int32(16)
										return v124
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
func F_pgp_s2k_fill(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v87 int32
	_ = v87
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v142 int32
	_ = v142
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v163 int32
	_ = v163
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v175 int32
	_ = v175
	var v186 int32
	_ = v186
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v210 int32
	_ = v210
	var v215 int32
	_ = v215
	var v222 int32
	_ = v222
	var v229 int32
	_ = v229
	v2 = l1
	v3 = l2
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)) = uint8(v3)
	*(*uint8)(unsafe.Add(mBase, uint32(l0))) = uint8(v2)
	v14 = v2 & int32(255)
	switch v14 {
	case 0:
		v229 = v14
		goto L1
	case 1:
		goto L6
	default:
		goto L4
	case 3:
		goto L5
	}
L1:
	;
	m.G0 = v9 + int32(16)
	return v229
L2:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+10)) = uint8(v222)
	v229 = int32(0)
	goto L1
L3:
	;
	v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+15)))
	v222 = v215&int32(31) | int32(96)
	goto L2
L4:
	;
	v229 = int32(-121)
	goto L1
L5:
	;
	v71 = int32(-17)
	v75 = int32(0)
	v79 = m.G0
	v81 = v79 - int32(16)
	m.G0 = v81
	*(*int32)(unsafe.Add(mBase, uint32(v81))) = v75
	v87 = F_open(m, int32(_a_F_pgp_s2k_fill_0), v75, v81)
	mBase = m.M
	if v87 != int32(-1) {
		goto L24
	} else {
		goto L25
	}
L6:
	;
	v15 = int32(0)
	v24 = m.G0
	v26 = v24 - int32(16)
	m.G0 = v26
	*(*int32)(unsafe.Add(mBase, uint32(v26))) = v15
	v32 = F_open(m, int32(_a_F_pgp_s2k_fill_0), v15, v26)
	mBase = m.M
	if v32 != int32(-1) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	if v65 != 0 {
		goto L20
	} else {
		goto L21
	}
L8:
	;
	goto L12
L9:
	;
	v65 = v15
	goto L10
L10:
	;
	m.G0 = v26 + int32(16)
	goto L7
L11:
	;
	v60 = F_close(m, v32)
	mBase = m.M
	v65 = v58
	goto L10
L12:
	;
	v38 = l0 + int32(2)
	v39 = int32(8)
	goto L13
L13:
	;
	v44 = F_read(m, v32, v38, v39)
	mBase = m.M
	if v44 <= int32(0) {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v58 = int32(1)
	goto L11
L15:
	;
	v48 = *(*int32)(unsafe.Add(mBase, _c_F_pgp_s2k_fill[0]))
	if v48 == int32(27) {
		goto L13
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	v53 = v39 - v44
	if v53 != 0 {
		v38 = v38 + v44
		v39 = v53
		goto L13
	} else {
		goto L19
	}
L18:
	;
	v58 = int32(0)
	goto L11
L19:
	;
	goto L14
L20:
	;
	v70 = v15
	goto L22
L21:
	;
	v70 = int32(-17)
	goto L22
L22:
	;
	v229 = v70
	goto L1
L23:
	;
	if v120 == int32(0) {
		v229 = v71
		goto L1
	} else {
		goto L36
	}
L24:
	;
	goto L28
L25:
	;
	v120 = v75
	goto L26
L26:
	;
	m.G0 = v81 + int32(16)
	goto L23
L27:
	;
	v115 = F_close(m, v87)
	mBase = m.M
	v120 = v113
	goto L26
L28:
	;
	v93 = l0 + int32(2)
	v94 = int32(8)
	goto L29
L29:
	;
	v99 = F_read(m, v87, v93, v94)
	mBase = m.M
	if v99 <= int32(0) {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	v113 = int32(1)
	goto L27
L31:
	;
	v103 = *(*int32)(unsafe.Add(mBase, _c_F_pgp_s2k_fill[0]))
	if v103 == int32(27) {
		goto L29
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	v108 = v94 - v99
	if v108 != 0 {
		v93 = v93 + v99
		v94 = v108
		goto L29
	} else {
		goto L35
	}
L34:
	;
	v113 = int32(0)
	goto L27
L35:
	;
	goto L30
L36:
	;
	v130 = int32(0)
	v134 = m.G0
	v136 = v134 - int32(16)
	m.G0 = v136
	*(*int32)(unsafe.Add(mBase, uint32(v136))) = v130
	v142 = F_open(m, int32(_a_F_pgp_s2k_fill_0), v130, v136)
	mBase = m.M
	if v142 != int32(-1) {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	if v175 == int32(0) {
		v229 = v71
		goto L1
	} else {
		goto L50
	}
L38:
	;
	goto L42
L39:
	;
	v175 = v130
	goto L40
L40:
	;
	m.G0 = v136 + int32(16)
	goto L37
L41:
	;
	v170 = F_close(m, v142)
	mBase = m.M
	v175 = v168
	goto L40
L42:
	;
	v148 = v9 + int32(15)
	v149 = int32(1)
	goto L43
L43:
	;
	v154 = F_read(m, v142, v148, v149)
	mBase = m.M
	if v154 <= int32(0) {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	v168 = int32(1)
	goto L41
L45:
	;
	v158 = *(*int32)(unsafe.Add(mBase, _c_F_pgp_s2k_fill[0]))
	if v158 == int32(27) {
		goto L43
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	v163 = v149 - v154
	if v163 != 0 {
		v148 = v148 + v154
		v149 = v163
		goto L43
	} else {
		goto L49
	}
L48:
	;
	v168 = int32(0)
	goto L41
L49:
	;
	goto L44
L50:
	;
	if l3 == int32(-1) {
		goto L3
	} else {
		goto L51
	}
L51:
	;
	v186 = int32(0)
	goto L52
L52:
	;
	v198 = int32(base.Ui32(v186)>>(uint(int32(4))%32)) + int32(6)
	if base.Ui32(l3) <= base.Ui32((v186&int32(14)|int32(16))<<(uint(v198)%32)) {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	v222 = int32(255)
	goto L2
L54:
	;
	v222 = v186
	goto L2
L55:
	;
	goto L56
L56:
	;
	v202 = v186 | int32(1)
	if base.Ui32(l3) <= base.Ui32((v202&int32(15)|int32(16))<<(uint(v198)%32)) {
		v222 = v202
		goto L2
	} else {
		goto L57
	}
L57:
	;
	v210 = v186 + int32(2)
	if v210 != int32(256) {
		v186 = v210
		goto L52
	} else {
		goto L58
	}
L58:
	;
	goto L53
}
func F_pgp_set_cipher_algo(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	v6 = F_pg_strcasecmp(m, int32(_a_F_pgp_set_cipher_algo_0), l1)
	mBase = m.M
	if v6 == int32(0) {
		v49 = int32(_a_F_pgp_set_cipher_algo_1)
		v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)+4))
		v51 = v50
	} else {
		v11 = F_pg_strcasecmp(m, int32(_a_F_pgp_set_cipher_algo_2), l1)
		mBase = m.M
		if v11 == int32(0) {
			v49 = int32(_a_F_pgp_set_cipher_algo_3)
			v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)+4))
			v51 = v50
		} else {
			v16 = F_pg_strcasecmp(m, int32(_a_F_pgp_set_cipher_algo_4), l1)
			mBase = m.M
			if v16 == int32(0) {
				v49 = int32(_a_F_pgp_set_cipher_algo_5)
				v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)+4))
				v51 = v50
			} else {
				v21 = F_pg_strcasecmp(m, int32(_a_F_pgp_set_cipher_algo_6), l1)
				mBase = m.M
				if v21 == int32(0) {
					v49 = int32(_a_F_pgp_set_cipher_algo_7)
					v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)+4))
					v51 = v50
				} else {
					v26 = F_pg_strcasecmp(m, int32(_a_F_pgp_set_cipher_algo_8), l1)
					mBase = m.M
					if v26 == int32(0) {
						v49 = int32(_a_F_pgp_set_cipher_algo_9)
						v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)+4))
						v51 = v50
					} else {
						v31 = F_pg_strcasecmp(m, int32(_a_F_pgp_set_cipher_algo_10), l1)
						mBase = m.M
						if v31 == int32(0) {
							v49 = int32(_a_F_pgp_set_cipher_algo_11)
							v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)+4))
							v51 = v50
						} else {
							v36 = F_pg_strcasecmp(m, int32(_a_F_pgp_set_cipher_algo_12), l1)
							mBase = m.M
							if v36 == int32(0) {
								v49 = int32(_a_F_pgp_set_cipher_algo_13)
								v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)+4))
								v51 = v50
							} else {
								v41 = F_pg_strcasecmp(m, int32(_a_F_pgp_set_cipher_algo_14), l1)
								mBase = m.M
								if v41 == int32(0) {
									v49 = int32(_a_F_pgp_set_cipher_algo_15)
									v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)+4))
									v51 = v50
								} else {
									v46 = F_pg_strcasecmp(m, int32(_a_F_pgp_set_cipher_algo_16), l1)
									mBase = m.M
									if v46 != 0 {
										v51 = int32(-103)
									} else {
										v49 = int32(_a_F_pgp_set_cipher_algo_17)
										v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)+4))
										v51 = v50
									}
								}
							}
						}
					}
				}
			}
		}
	}
	if v51 < int32(0) {
		return v51
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v51
		return int32(0)
	}
}
func F_pgp_set_compress_level(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	if base.Ui32(l1) <= base.Ui32(int32(9)) {
		*(*int32)(unsafe.Add(mBase, uint32(l0)+68)) = l1
		v8 = int32(0)
	} else {
		v8 = int32(-13)
	}
	return v8
}
func F_pgp_set_convert_crlf(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	v3 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = base.B2i32(l1 != v3)
	return v3
}
