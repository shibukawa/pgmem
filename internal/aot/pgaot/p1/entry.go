package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_entryGetItem(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v61 int32
	_ = v61
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v175 int32
	_ = v175
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
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
	var v236 int32
	_ = v236
	var v242 int32
	_ = v242
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
	var v275 int32
	_ = v275
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v310 int32
	_ = v310
	var v318 int32
	_ = v318
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v335 int32
	_ = v335
	var v339 int32
	_ = v339
	var v343 int32
	_ = v343
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v363 int32
	_ = v363
	var v366 int64
	_ = v366
	var v367 int64
	_ = v367
	var v368 int64
	_ = v368
	var v389 float64
	_ = v389
	var v391 int32
	_ = v391
	var v393 int32
	_ = v393
	var v398 int32
	_ = v398
	var v404 int32
	_ = v404
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v423 int32
	_ = v423
	var v425 int32
	_ = v425
	var v427 int32
	_ = v427
	var v430 int32
	_ = v430
	var v432 int64
	_ = v432
	var v433 int64
	_ = v433
	var v434 int64
	_ = v434
	var v436 int64
	_ = v436
	var v437 int64
	_ = v437
	var v440 int64
	_ = v440
	var v441 int32
	_ = v441
	var v443 int32
	_ = v443
	var v447 int32
	_ = v447
	var v452 int64
	_ = v452
	var v453 int32
	_ = v453
	var v458 int32
	_ = v458
	var v460 int32
	_ = v460
	var v465 int32
	_ = v465
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v471 int32
	_ = v471
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v482 int32
	_ = v482
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v491 int32
	_ = v491
	var v493 int32
	_ = v493
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v502 int32
	_ = v502
	var v510 int32
	_ = v510
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v520 int32
	_ = v520
	var v526 int32
	_ = v526
	var v528 int32
	_ = v528
	var v534 int32
	_ = v534
	var v539 int32
	_ = v539
	var v542 int32
	_ = v542
	var v551 int32
	_ = v551
	var v553 int32
	_ = v553
	var v555 int32
	_ = v555
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v565 int32
	_ = v565
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v574 int32
	_ = v574
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v582 int32
	_ = v582
	var v588 int32
	_ = v588
	var v590 int32
	_ = v590
	var v596 int32
	_ = v596
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v600 int32
	_ = v600
	var v601 int32
	_ = v601
	var v604 int32
	_ = v604
	var v607 int64
	_ = v607
	var v608 int64
	_ = v608
	var v609 int64
	_ = v609
	var v611 int64
	_ = v611
	var v612 int64
	_ = v612
	var v616 int64
	_ = v616
	var v617 int64
	_ = v617
	var v620 int64
	_ = v620
	var v626 int32
	_ = v626
	var v628 int32
	_ = v628
	var v632 int32
	_ = v632
	var v633 int32
	_ = v633
	var v635 int32
	_ = v635
	var v638 int64
	_ = v638
	var v639 int64
	_ = v639
	var v642 int64
	_ = v642
	var v652 int32
	_ = v652
	var v665 int32
	_ = v665
	var v666 int64
	_ = v666
	var v669 int64
	_ = v669
	var v673 int64
	_ = v673
	var v677 int32
	_ = v677
	var v678 int32
	_ = v678
	var v680 int32
	_ = v680
	var v684 int32
	_ = v684
	var v687 int32
	_ = v687
	var v689 int32
	_ = v689
	var v693 int32
	_ = v693
	var v709 int32
	_ = v709
	var v712 int32
	_ = v712
	var v717 int32
	_ = v717
	var v719 int32
	_ = v719
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v725 int32
	_ = v725
	var v727 int64
	_ = v727
	var v728 int64
	_ = v728
	var v729 int64
	_ = v729
	var v731 int64
	_ = v731
	var v732 int64
	_ = v732
	var v736 int64
	_ = v736
	var v737 int64
	_ = v737
	var v740 int64
	_ = v740
	var v746 int32
	_ = v746
	var v749 int32
	_ = v749
	var v752 int64
	_ = v752
	var v753 int64
	_ = v753
	var v754 int64
	_ = v754
	var v775 float64
	_ = v775
	var v777 int32
	_ = v777
	var v779 int32
	_ = v779
	var v785 int32
	_ = v785
	var v787 int32
	_ = v787
	var v789 int32
	_ = v789
	var v790 int32
	_ = v790
	var v792 int64
	_ = v792
	var v793 int64
	_ = v793
	var v796 int64
	_ = v796
	var v802 int32
	_ = v802
	var v803 int32
	_ = v803
	var v806 int32
	_ = v806
	var v807 int32
	_ = v807
	var v819 int32
	_ = v819
	var v821 int32
	_ = v821
	var v824 int32
	_ = v824
	var v825 int32
	_ = v825
	var v827 int32
	_ = v827
	var v829 int64
	_ = v829
	var v830 int64
	_ = v830
	var v833 int64
	_ = v833
	var v839 int32
	_ = v839
	var v842 int32
	_ = v842
	var v845 int64
	_ = v845
	var v846 int64
	_ = v846
	var v847 int64
	_ = v847
	var v868 float64
	_ = v868
	var v870 int32
	_ = v870
	var v872 int32
	_ = v872
	var v878 int32
	_ = v878
	var v879 int32
	_ = v879
	var v880 int32
	_ = v880
	var v881 int32
	_ = v881
	var v883 int32
	_ = v883
	var v900 int32
	_ = v900
	var v902 int32
	_ = v902
	v16 = m.G0
	v18 = v16 - int32(48)
	m.G0 = v18
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	if v20 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v18 + int32(48)
	return
L2:
	;
	v22 = l1 + int32(56)
	v24 = l1 + int32(44)
	v25 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+2)))
	v26 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2))))
	v29 = v25 | v26<<(uint(int32(16))%32)
	v30 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+4)))
	goto L5
L3:
	;
	goto L4
L4:
	;
	v398 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	if v398 != 0 {
		goto L85
	} else {
		goto L86
	}
L5:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	if v46 == int32(-1) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v281 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+654)))
	if v281 != 0 {
		goto L1
	} else {
		goto L68
	}
L8:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v61)+8))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)+28))
	if v70 <= v68 {
		goto L21
	} else {
		goto L22
	}
L9:
	;
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+48)))
	if v49 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	if base.B2i32(v30 != int32(_a_F_entryGetItem_0))|base.B2i32(v46 != v29) != 0 {
		goto L7
	} else {
		goto L17
	}
L11:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l1)+640))
	v53 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+652)))
	if v52 <= v53 {
		goto L8
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	if base.Ui32(v46) < base.Ui32(v29) {
		goto L8
	} else {
		goto L16
	}
L14:
	;
	if base.Ui32(v29) <= base.Ui32(v46) {
		goto L10
	} else {
		goto L15
	}
L15:
	;
	goto L8
L16:
	;
	goto L10
L17:
	;
	goto L8
L18:
	;
	if v209 == int32(0) {
		goto L44
	} else {
		goto L45
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24))) = int32(-1)
	v209 = int32(0)
	goto L18
L20:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v69)+8))
	if v175 == int32(1) {
		goto L41
	} else {
		goto L42
	}
L21:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v61)+4))
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v69)+24))
	if v165 <= v164 {
		goto L19
	} else {
		goto L40
	}
L22:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v61)+12))
	v75 = v72
	v78 = v68
	goto L23
L23:
	;
	v81 = int32(256)
	if v75 <= v81 {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	goto L21
L25:
	;
	v84 = v81
	goto L27
L26:
	;
	v84 = v75
	goto L27
L27:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v69)+92))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v85+v78<<(uint(int32(2))%32))))
	v94 = v75
	goto L29
L28:
	;
	v148 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v61)+12)) = v148
	v152 = v78 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v61)+8)) = v152
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v69)+28))
	if v152 < v154 {
		v75 = v148
		v78 = v152
		goto L23
	} else {
		goto L39
	}
L29:
	;
	if v94 == v84 {
		goto L28
	} else {
		goto L31
	}
L30:
	;
	if int32(255) < v94 {
		goto L28
	} else {
		goto L33
	}
L31:
	;
	v101 = int32(1)
	v104 = base.I32_div_s(v94, int32(32))
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v89+int32(8)+v104<<(uint(int32(2))%32))))
	if int32(base.Ui32(v108)>>(uint(v94)%32))&v101 == int32(0) {
		v94 = v94 + v101
		goto L29
	} else {
		goto L32
	}
L32:
	;
	goto L30
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v61)+12)) = v94
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v69)+28))
	if v117 <= v78 {
		goto L21
	} else {
		goto L34
	}
L34:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v61)+12))
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v69)+92))
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v120+v78<<(uint(int32(2))%32))))
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v124)))
	v126 = v119 + v125
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v61)+4))
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v69)+24))
	if v127 < v128 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v69)+88))
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v130+v127<<(uint(int32(2))%32))))
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v134)))
	if base.Ui32(v135) <= base.Ui32(v126) {
		v169 = v127
		goto L20
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+8)) = int32(0)
	v139 = int32(257)
	*(*uint16)(unsafe.Add(mBase, uint32(v24)+4)) = uint16(v139)
	*(*int32)(unsafe.Add(mBase, uint32(v24))) = v126
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v61)+12))
	v143 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v61)+12)) = v142 + v143
	v209 = v143
	goto L18
L38:
	;
	goto L37
L39:
	;
	goto L24
L40:
	;
	v169 = v164
	goto L20
L41:
	;
	v185 = v69 + int32(40)
	goto L43
L42:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v69)+88))
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v180+v169<<(uint(int32(2))%32))))
	v185 = v184
	goto L43
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+8)) = v185
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v185)))
	v188 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+4)) = uint8(v188)
	*(*int32)(unsafe.Add(mBase, uint32(v24))) = v187
	v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v185)+6)))
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+5)) = uint8(v191)
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v61)+4))
	v194 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v61)+4)) = v193 + v194
	v209 = v194
	goto L18
L44:
	;
	v212 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(l1)+32)) = uint16(v212)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = int32(-1)
	v216 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	F_pfree(m, v216)
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L47
	} else {
		goto L48
	}
L45:
	;
	goto L46
L46:
	;
	v223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+48)))
	if v223 == int32(0) {
		goto L49
	} else {
		goto L50
	}
L47:
	;
	return
L48:
	;
	v219 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+654)) = uint8(v219)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+40)) = int32(0)
	goto L1
L49:
	;
	v226 = int32(0)
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
	v235 = v226
	v236 = v226
	goto L53
L50:
	;
	goto L51
L51:
	;
	v279 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(l1)+652)) = uint16(v279)
	goto L5
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+640)) = v270
	goto L51
L53:
	;
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v230+int32(8)+v236<<(uint(int32(2))%32))))
	if v242 != 0 {
		goto L55
	} else {
		goto L56
	}
L54:
	;
	goto L52
L55:
	;
	v247 = v242
	v249 = v235
	v251 = v236<<(uint(int32(5))%32) | int32(1)
	goto L58
L56:
	;
	v270 = v235
	goto L57
L57:
	;
	v275 = v236 + int32(1)
	if v275 != int32(10) {
		v235 = v270
		v236 = v275
		goto L53
	} else {
		goto L67
	}
L58:
	;
	if v247&int32(1) != 0 {
		goto L60
	} else {
		goto L61
	}
L59:
	;
	v270 = v263
	goto L57
L60:
	;
	if base.Ui32(v249) < base.Ui32(int32(291)) {
		goto L63
	} else {
		goto L64
	}
L61:
	;
	v263 = v249
	goto L62
L62:
	;
	v264 = int32(1)
	v267 = int32(base.Ui32(v247) >> (uint(v264) % 32))
	if v267 != 0 {
		v247 = v267
		v249 = v263
		v251 = v251 + v264
		goto L58
	} else {
		goto L66
	}
L63:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v22+v249<<(uint(int32(1))%32)))) = uint16(v251)
	goto L65
L64:
	;
	goto L65
L65:
	;
	v263 = v249 + int32(1)
	goto L62
L66:
	;
	goto L59
L67:
	;
	goto L54
L68:
	;
	v282 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
	v283 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+48)))
	if v283 == int32(1) {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v286 = int32(_a_F_entryGetItem_0)
	*(*uint16)(unsafe.Add(mBase, uint32(l1)+32)) = uint16(v286)
	*(*uint16)(unsafe.Add(mBase, uint32(l1)+30)) = uint16(v282)
	v290 = int32(base.Ui32(v282) >> (uint(int32(16)) % 32))
	*(*uint16)(unsafe.Add(mBase, uint32(l1)+28)) = uint16(v290)
	goto L1
L70:
	;
	goto L71
L71:
	;
	if v282 != v29 {
		goto L74
	} else {
		goto L75
	}
L72:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(l1)+652)) = uint16(v298)
	goto L5
L73:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(l1)+32)) = uint16(v339)
	*(*uint16)(unsafe.Add(mBase, uint32(l1)+30)) = uint16(v282)
	v354 = int32(1)
	v355 = v343 + v354
	*(*uint16)(unsafe.Add(mBase, uint32(l1)+652)) = uint16(v355)
	v358 = int32(base.Ui32(v282) >> (uint(int32(16)) % 32))
	*(*uint16)(unsafe.Add(mBase, uint32(l1)+28)) = uint16(v358)
	v360 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+655)))
	if v360 != v354 {
		goto L1
	} else {
		goto L82
	}
L74:
	;
	v293 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+652)))
	v297 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v22+v293<<(uint(int32(1))%32)))))
	v339 = v297
	v343 = v293
	goto L73
L75:
	;
	goto L76
L76:
	;
	v298 = *(*int32)(unsafe.Add(mBase, uint32(l1)+640))
	v304 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v22+v298<<(uint(int32(1))%32)-int32(2)))))
	if base.Ui32(v304) <= base.Ui32(v30) {
		goto L72
	} else {
		goto L77
	}
L77:
	;
	v306 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+652)))
	v310 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v22+v306<<(uint(int32(1))%32)))))
	if base.Ui32(v30) < base.Ui32(v310) {
		v339 = v310
		v343 = v306
		goto L73
	} else {
		goto L78
	}
L78:
	;
	v318 = v306
	goto L79
L79:
	;
	v327 = int32(1)
	v328 = v318 + v327
	*(*uint16)(unsafe.Add(mBase, uint32(l1)+652)) = uint16(v328)
	v335 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v22+v328&int32(_a_F_entryGetItem_0)<<(uint(v327)%32)))))
	if base.Ui32(v335) <= base.Ui32(v30) {
		v318 = v328
		goto L79
	} else {
		goto L81
	}
L80:
	;
	v339 = v335
	v343 = v328
	goto L73
L81:
	;
	goto L80
L82:
	;
	v363 = int32(_a_F_entryGetItem_1)
	v366 = *(*int64)(unsafe.Add(mBase, _c_F_entryGetItem[0]))
	v367 = *(*int64)(unsafe.Add(mBase, _c_F_entryGetItem[1]))
	v368 = v366 ^ v367
	*(*int64)(unsafe.Add(mBase, _c_F_entryGetItem[1])) = base.I64_rotl(v368, int64(37))
	*(*int64)(unsafe.Add(mBase, _c_F_entryGetItem[0])) = v368<<(uint(int64(16))%64) ^ base.I64_rotl(v366, int64(24)) ^ v368
	v389 = F_scalbn(m, base.F64_convert_i64_u(int64(base.Ui64(base.I64_rotl(v366*int64(5), int64(7))*int64(9))>>(uint(int64(12))%64))), int32(-52))
	mBase = m.M
	goto L83
L83:
	;
	v391 = *(*int32)(unsafe.Add(mBase, _c_F_entryGetItem[2]))
	v393 = *(*int32)(unsafe.Add(mBase, uint32(l1)+656))
	if base.F64_gt(v389, base.F64_div(base.F64_convert_i32_s(v391), base.F64_convert_i32_u(v393))) != 0 {
		goto L5
	} else {
		goto L84
	}
L84:
	;
	goto L1
L85:
	;
	v404 = l1 + int32(28)
	goto L88
L86:
	;
	goto L87
L87:
	;
	v789 = *(*int32)(unsafe.Add(mBase, uint32(l1)+648))
	v790 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+652)))
	if v790 < v789 {
		goto L158
	} else {
		goto L159
	}
L88:
	;
	v420 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+652)))
	v421 = *(*int32)(unsafe.Add(mBase, uint32(l1)+648))
	if v421 <= v420 {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v423 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v18)+44)) = uint16(v423)
	v425 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+40)) = v425
	v427 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	if v427 == int32(0) {
		goto L94
	} else {
		goto L95
	}
L91:
	;
	goto L92
L92:
	;
	v717 = v420 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(l1)+652)) = uint16(v717)
	v719 = *(*int32)(unsafe.Add(mBase, uint32(l1)+644))
	v722 = v719 + v420*int32(6)
	v723 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v722)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v404)+4)) = uint16(v723)
	v725 = *(*int32)(unsafe.Add(mBase, uint32(v722)))
	*(*int32)(unsafe.Add(mBase, uint32(v404))) = v725
	v727 = int64(*(*uint16)(unsafe.Add(mBase, uint32(l1)+32)))
	v728 = int64(*(*uint16)(unsafe.Add(mBase, uint32(l1)+30)))
	v729 = int64(32)
	v731 = int64(*(*uint16)(unsafe.Add(mBase, uint32(l1)+28)))
	v732 = int64(48)
	v736 = int64(*(*uint16)(unsafe.Add(mBase, uint32(l2)+4)))
	v737 = int64(*(*uint16)(unsafe.Add(mBase, uint32(l2)+2)))
	v740 = int64(*(*uint16)(unsafe.Add(mBase, uint32(l2))))
	if base.Ui64(v727|(v728<<(uint(v729)%64)|v731<<(uint(v732)%64))) <= base.Ui64(v736|(v737<<(uint(v729)%64)|v740<<(uint(v732)%64))) {
		goto L88
	} else {
		goto L154
	}
L93:
	;
	v709 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+654)))
	if v709 != int32(1) {
		goto L88
	} else {
		goto L153
	}
L94:
	;
	v430 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+654)) = uint8(v430)
	goto L93
L95:
	;
	goto L96
L96:
	;
	v432 = int64(*(*uint16)(unsafe.Add(mBase, uint32(l1)+32)))
	v433 = int64(*(*uint16)(unsafe.Add(mBase, uint32(l1)+30)))
	v434 = int64(32)
	v436 = int64(*(*uint16)(unsafe.Add(mBase, uint32(l1)+28)))
	v437 = int64(48)
	v440 = v432 | (v433<<(uint(v434)%64) | v436<<(uint(v437)%64))
	v441 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v18)+44)))
	v443 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v18)+42)))
	v447 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v18)+40)))
	v452 = base.I64_extend_i32_u(v441) | (base.I64_extend_i32_u(v443)<<(uint(v434)%64) | base.I64_extend_i32_u(v447)<<(uint(v437)%64))
	v453 = base.B2i32(v440 != v452)
	if v453 == int32(0) {
		goto L98
	} else {
		goto L99
	}
L97:
	;
	v498 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v499 = m.ExcPending
	if v499 != 0 {
		goto L47
	} else {
		goto L110
	}
L98:
	;
	F_LockBuffer(m, v427, int32(1))
	mBase = m.M
	v458 = m.ExcPending
	if v458 != 0 {
		goto L47
	} else {
		goto L101
	}
L99:
	;
	goto L100
L100:
	;
	F_ReleaseBuffer(m, v427)
	mBase = m.M
	v460 = m.ExcPending
	if v460 != 0 {
		goto L47
	} else {
		goto L102
	}
L101:
	;
	goto L97
L102:
	;
	if v441 != int32(_a_F_entryGetItem_0) {
		goto L104
	} else {
		goto L105
	}
L103:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(l1)+726)) = uint16(v479)
	*(*uint16)(unsafe.Add(mBase, uint32(l1)+724)) = uint16(v478)
	v482 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+712)) = uint8(v482)
	v486 = F_ginFindLeafPage(m, l1+int32(660), int32(1), v482)
	mBase = m.M
	v487 = m.ExcPending
	if v487 != 0 {
		goto L47
	} else {
		goto L107
	}
L104:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(l1)+722)) = uint16(v447)
	v478 = v443
	v479 = v441 + int32(1)
	goto L103
L105:
	;
	v465 = v447<<(uint(int32(16))%32) | v443
	if v465 == int32(-1) {
		goto L104
	} else {
		goto L106
	}
L106:
	;
	v468 = int32(1)
	v469 = v465 + v468
	v471 = int32(base.Ui32(v469) >> (uint(int32(16)) % 32))
	*(*uint16)(unsafe.Add(mBase, uint32(l1)+722)) = uint16(v471)
	v478 = v469
	v479 = v468
	goto L103
L107:
	;
	v488 = *(*int32)(unsafe.Add(mBase, uint32(v486)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v488
	F_IncrBufferRefCount(m, v488)
	mBase = m.M
	v491 = m.ExcPending
	if v491 != 0 {
		goto L47
	} else {
		goto L108
	}
L108:
	;
	F_freeGinBtreeStack(m, v486)
	mBase = m.M
	v493 = m.ExcPending
	if v493 != 0 {
		goto L47
	} else {
		goto L109
	}
L109:
	;
	goto L97
L110:
	;
	if v498 != 0 {
		goto L111
	} else {
		goto L112
	}
L111:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+20)) = v441
	*(*int32)(unsafe.Add(mBase, uint32(v18)+24)) = v453
	v502 = int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v447<<(uint(v502)%32) | v443
	F_errmsg_internal(m, int32(_a_F_entryGetItem_2), v18+v502)
	mBase = m.M
	v510 = m.ExcPending
	if v510 != 0 {
		goto L47
	} else {
		goto L114
	}
L112:
	;
	goto L113
L113:
	;
	v516 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	if v516 < int32(0) {
		goto L117
	} else {
		goto L118
	}
L114:
	;
	F_errfinish(m, int32(_a_F_entryGetItem_3), int32(715), int32(_a_F_entryGetItem_4))
	mBase = m.M
	v515 = m.ExcPending
	if v515 != 0 {
		goto L47
	} else {
		goto L115
	}
L115:
	;
	goto L113
L116:
	;
	v539 = v534
	v542 = base.B2i32(v440 == v452)
	goto L120
L117:
	;
	v520 = *(*int32)(unsafe.Add(mBase, _c_F_entryGetItem[3]))
	v526 = *(*int32)(unsafe.Add(mBase, uint32(v520+(v516^int32(-1))<<(uint(int32(2))%32))))
	v534 = v526
	goto L116
L118:
	;
	goto L119
L119:
	;
	v528 = *(*int32)(unsafe.Add(mBase, _c_F_entryGetItem[4]))
	v534 = v528 + v516<<(uint(int32(13))%32) + int32(-8192)
	goto L116
L120:
	;
	v551 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(l1)+652)) = uint16(v551)
	v553 = *(*int32)(unsafe.Add(mBase, uint32(l1)+644))
	if v553 != 0 {
		goto L122
	} else {
		goto L123
	}
L121:
	;
	F_LockBuffer(m, v677, int32(0))
	mBase = m.M
	v693 = m.ExcPending
	if v693 != 0 {
		goto L47
	} else {
		goto L152
	}
L122:
	;
	F_pfree(m, v553)
	mBase = m.M
	v555 = m.ExcPending
	if v555 != 0 {
		goto L47
	} else {
		goto L125
	}
L123:
	;
	goto L124
L124:
	;
	if v542&int32(1) == int32(0) {
		v596 = v539
		goto L126
	} else {
		goto L127
	}
L125:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l1)+644)) = int64(0)
	goto L124
L126:
	;
	v598 = int32(1)
	v599 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v596)+16)))
	v600 = v596 + v599
	v601 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v600)+6)))
	if v601&int32(4) != 0 {
		v539 = v596
		v542 = v598
		goto L120
	} else {
		goto L136
	}
L127:
	;
	v562 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v563 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v539)+16)))
	v565 = *(*int32)(unsafe.Add(mBase, uint32(v539+v563)))
	if v565 == int32(-1) {
		goto L128
	} else {
		goto L129
	}
L128:
	;
	F_UnlockReleaseBuffer(m, v562)
	mBase = m.M
	v569 = m.ExcPending
	if v569 != 0 {
		goto L47
	} else {
		goto L131
	}
L129:
	;
	goto L130
L130:
	;
	v574 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v576 = F_ginStepRight(m, v562, v574, int32(1))
	mBase = m.M
	v577 = m.ExcPending
	if v577 != 0 {
		goto L47
	} else {
		goto L132
	}
L131:
	;
	v570 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+654)) = uint8(v570)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = int32(0)
	goto L93
L132:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v576
	if v576 < int32(0) {
		goto L133
	} else {
		goto L134
	}
L133:
	;
	v582 = *(*int32)(unsafe.Add(mBase, _c_F_entryGetItem[3]))
	v588 = *(*int32)(unsafe.Add(mBase, uint32(v582+(v576^int32(-1))<<(uint(int32(2))%32))))
	v596 = v588
	goto L126
L134:
	;
	goto L135
L135:
	;
	v590 = *(*int32)(unsafe.Add(mBase, _c_F_entryGetItem[4]))
	v596 = v590 + v576<<(uint(int32(13))%32) + int32(-8192)
	goto L126
L136:
	;
	v604 = *(*int32)(unsafe.Add(mBase, uint32(v600)))
	if v604 != int32(-1) {
		goto L137
	} else {
		goto L138
	}
L137:
	;
	v607 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v18)+44)))
	v608 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v18)+42)))
	v609 = int64(32)
	v611 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v18)+40)))
	v612 = int64(48)
	v616 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v596)+28)))
	v617 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v596)+26)))
	v620 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v596)+24)))
	if base.Ui64(v616|(v617<<(uint(v609)%64)|v620<<(uint(v612)%64))) <= base.Ui64(v607|(v608<<(uint(v609)%64)|v611<<(uint(v612)%64))) {
		v539 = v596
		v542 = v598
		goto L120
	} else {
		goto L140
	}
L138:
	;
	goto L139
L139:
	;
	v626 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v18)+44)))
	*(*uint16)(unsafe.Add(mBase, uint32(v18)+12)) = uint16(v626)
	v628 = *(*int32)(unsafe.Add(mBase, uint32(v18)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = v628
	v632 = F_GinDataLeafPageGetItems(m, v596, l1+int32(648), v18+int32(8))
	mBase = m.M
	v633 = m.ExcPending
	if v633 != 0 {
		goto L47
	} else {
		goto L141
	}
L140:
	;
	goto L139
L141:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+644)) = v632
	v635 = *(*int32)(unsafe.Add(mBase, uint32(l1)+648))
	if v635 <= int32(0) {
		v539 = v596
		v542 = v598
		goto L120
	} else {
		goto L142
	}
L142:
	;
	v638 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v18)+44)))
	v639 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v18)+42)))
	v642 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v18)+40)))
	v652 = int32(0)
	goto L143
L143:
	;
	v665 = v632 + v652*int32(6)
	v666 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v665)+2)))
	v669 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v665))))
	v673 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v665)+4)))
	if base.Ui64(v638|(v639<<(uint(int64(32))%64)|v642<<(uint(int64(48))%64))) < base.Ui64(v666<<(uint(int64(32))%64)|v669<<(uint(int64(48))%64)|v673) {
		goto L146
	} else {
		goto L147
	}
L144:
	;
	goto L121
L145:
	;
	goto L144
L146:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(l1)+652)) = uint16(v652)
	v677 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v678 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v596)+16)))
	v680 = *(*int32)(unsafe.Add(mBase, uint32(v596+v678)))
	if v680 != int32(-1) {
		goto L145
	} else {
		goto L149
	}
L147:
	;
	goto L148
L148:
	;
	v687 = int32(1)
	v689 = v652 + v687
	if v635 != v689 {
		v652 = v689
		goto L143
	} else {
		goto L151
	}
L149:
	;
	F_UnlockReleaseBuffer(m, v677)
	mBase = m.M
	v684 = m.ExcPending
	if v684 != 0 {
		goto L47
	} else {
		goto L150
	}
L150:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = int32(0)
	goto L93
L151:
	;
	v539 = v596
	v542 = v687
	goto L120
L152:
	;
	goto L93
L153:
	;
	v712 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(l1)+32)) = uint16(v712)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = int32(-1)
	goto L1
L154:
	;
	v746 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+655)))
	if v746 != int32(1) {
		goto L1
	} else {
		goto L155
	}
L155:
	;
	v749 = int32(_a_F_entryGetItem_1)
	v752 = *(*int64)(unsafe.Add(mBase, _c_F_entryGetItem[0]))
	v753 = *(*int64)(unsafe.Add(mBase, _c_F_entryGetItem[1]))
	v754 = v752 ^ v753
	*(*int64)(unsafe.Add(mBase, _c_F_entryGetItem[1])) = base.I64_rotl(v754, int64(37))
	*(*int64)(unsafe.Add(mBase, _c_F_entryGetItem[0])) = v754<<(uint(int64(16))%64) ^ base.I64_rotl(v752, int64(24)) ^ v754
	v775 = F_scalbn(m, base.F64_convert_i64_u(int64(base.Ui64(base.I64_rotl(v752*int64(5), int64(7))*int64(9))>>(uint(int64(12))%64))), int32(-52))
	mBase = m.M
	goto L156
L156:
	;
	v777 = *(*int32)(unsafe.Add(mBase, _c_F_entryGetItem[2]))
	v779 = *(*int32)(unsafe.Add(mBase, uint32(l1)+656))
	if base.F64_gt(v775, base.F64_div(base.F64_convert_i32_s(v777), base.F64_convert_i32_u(v779))) == int32(0) {
		goto L1
	} else {
		goto L157
	}
L157:
	;
	v785 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v404)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(l2)+4)) = uint16(v785)
	v787 = *(*int32)(unsafe.Add(mBase, uint32(v404)))
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v787
	goto L88
L158:
	;
	v792 = int64(*(*uint16)(unsafe.Add(mBase, uint32(l2)+4)))
	v793 = int64(*(*uint16)(unsafe.Add(mBase, uint32(l2)+2)))
	v796 = int64(*(*uint16)(unsafe.Add(mBase, uint32(l2))))
	v802 = l1 + int32(28)
	v803 = v789
	v806 = v790
	v807 = v790
	goto L161
L159:
	;
	goto L160
L160:
	;
	v900 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+654)) = uint8(v900)
	v902 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(l1)+32)) = uint16(v902)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = int32(-1)
	goto L1
L161:
	;
	v819 = v807 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(l1)+652)) = uint16(v819)
	v821 = *(*int32)(unsafe.Add(mBase, uint32(l1)+644))
	v824 = v821 + v806*int32(6)
	v825 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v824)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v802)+4)) = uint16(v825)
	v827 = *(*int32)(unsafe.Add(mBase, uint32(v824)))
	*(*int32)(unsafe.Add(mBase, uint32(v802))) = v827
	v829 = int64(*(*uint16)(unsafe.Add(mBase, uint32(l1)+32)))
	v830 = int64(*(*uint16)(unsafe.Add(mBase, uint32(l1)+30)))
	v833 = int64(*(*uint16)(unsafe.Add(mBase, uint32(l1)+28)))
	if base.Ui64(v792|(v793<<(uint(int64(32))%64)|v796<<(uint(int64(48))%64))) < base.Ui64(v829|(v830<<(uint(int64(32))%64)|v833<<(uint(int64(48))%64))) {
		goto L163
	} else {
		goto L164
	}
L162:
	;
	goto L160
L163:
	;
	v839 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+655)))
	if v839 != int32(1) {
		goto L1
	} else {
		goto L166
	}
L164:
	;
	v880 = v803
	v881 = v819
	goto L165
L165:
	;
	v883 = v881 & int32(_a_F_entryGetItem_0)
	if v883 < v880 {
		v803 = v880
		v806 = v883
		v807 = v881
		goto L161
	} else {
		goto L169
	}
L166:
	;
	v842 = int32(_a_F_entryGetItem_1)
	v845 = *(*int64)(unsafe.Add(mBase, _c_F_entryGetItem[0]))
	v846 = *(*int64)(unsafe.Add(mBase, _c_F_entryGetItem[1]))
	v847 = v845 ^ v846
	*(*int64)(unsafe.Add(mBase, _c_F_entryGetItem[1])) = base.I64_rotl(v847, int64(37))
	*(*int64)(unsafe.Add(mBase, _c_F_entryGetItem[0])) = v847<<(uint(int64(16))%64) ^ base.I64_rotl(v845, int64(24)) ^ v847
	v868 = F_scalbn(m, base.F64_convert_i64_u(int64(base.Ui64(base.I64_rotl(v845*int64(5), int64(7))*int64(9))>>(uint(int64(12))%64))), int32(-52))
	mBase = m.M
	goto L167
L167:
	;
	v870 = *(*int32)(unsafe.Add(mBase, _c_F_entryGetItem[2]))
	v872 = *(*int32)(unsafe.Add(mBase, uint32(l1)+656))
	if base.F64_gt(v868, base.F64_div(base.F64_convert_i32_s(v870), base.F64_convert_i32_u(v872))) == int32(0) {
		goto L1
	} else {
		goto L168
	}
L168:
	;
	v878 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+652)))
	v879 = *(*int32)(unsafe.Add(mBase, uint32(l1)+648))
	v880 = v879
	v881 = v878
	goto L165
L169:
	;
	goto L162
}
func F_entryGetLeftMostPage(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v6 = l1 + v3&int32(_a_F_entryGetLeftMostPage_0)
	v7 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v6))))
	v10 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v6)+2)))
	return v7<<(uint(int32(16))%32) | v10
}
