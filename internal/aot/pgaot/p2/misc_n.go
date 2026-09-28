package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_NIAddAffix(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) {
	mBase := m.M
	_ = mBase
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v242 int32
	_ = v242
	var v248 int32
	_ = v248
	var v253 int32
	_ = v253
	var v258 int32
	_ = v258
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v285 int32
	_ = v285
	var v291 int32
	_ = v291
	var v296 int32
	_ = v296
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v314 int32
	_ = v314
	var v325 int32
	_ = v325
	var v331 int32
	_ = v331
	var v337 int32
	_ = v337
	var v342 int32
	_ = v342
	var v346 int32
	_ = v346
	var v353 int32
	_ = v353
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v391 int32
	_ = v391
	var v395 int32
	_ = v395
	var v411 int32
	_ = v411
	var v413 int32
	_ = v413
	var v437 int32
	_ = v437
	var v441 int32
	_ = v441
	var v446 int32
	_ = v446
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v457 int32
	_ = v457
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v497 int32
	_ = v497
	var v501 int32
	_ = v501
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v509 int32
	_ = v509
	var v512 int32
	_ = v512
	var v514 int32
	_ = v514
	var v516 int32
	_ = v516
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v539 int32
	_ = v539
	var v546 int32
	_ = v546
	var v550 int32
	_ = v550
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v560 int32
	_ = v560
	var v564 int32
	_ = v564
	var v566 int32
	_ = v566
	var v568 int32
	_ = v568
	var v571 int32
	_ = v571
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v583 int32
	_ = v583
	var v585 int32
	_ = v585
	var v588 int32
	_ = v588
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v602 int32
	_ = v602
	var v604 int32
	_ = v604
	var v605 int32
	_ = v605
	var v607 int32
	_ = v607
	var v616 int32
	_ = v616
	var v619 int32
	_ = v619
	var v625 int32
	_ = v625
	var v628 int32
	_ = v628
	var v630 int32
	_ = v630
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v638 int32
	_ = v638
	var v639 int32
	_ = v639
	var v641 int32
	_ = v641
	var v642 int32
	_ = v642
	var v644 int32
	_ = v644
	var v645 int32
	_ = v645
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v652 int32
	_ = v652
	var v660 int32
	_ = v660
	var v664 int32
	_ = v664
	var v666 int32
	_ = v666
	var v667 int32
	_ = v667
	var v671 int32
	_ = v671
	var v672 int32
	_ = v672
	var v674 int32
	_ = v674
	var v678 int32
	_ = v678
	var v680 int32
	_ = v680
	var v682 int32
	_ = v682
	var v685 int32
	_ = v685
	var v690 int32
	_ = v690
	var v691 int32
	_ = v691
	var v692 int32
	_ = v692
	var v694 int32
	_ = v694
	var v695 int32
	_ = v695
	var v697 int32
	_ = v697
	var v699 int32
	_ = v699
	var v702 int32
	_ = v702
	var v707 int32
	_ = v707
	var v708 int32
	_ = v708
	var v709 int32
	_ = v709
	var v716 int32
	_ = v716
	var v718 int32
	_ = v718
	var v719 int32
	_ = v719
	var v721 int32
	_ = v721
	var v729 int32
	_ = v729
	var v730 int32
	_ = v730
	var v731 int32
	_ = v731
	var v736 int32
	_ = v736
	var v738 int32
	_ = v738
	var v743 int32
	_ = v743
	var v745 int32
	_ = v745
	var v748 int32
	_ = v748
	var v749 int32
	_ = v749
	var v753 int32
	_ = v753
	var v754 int32
	_ = v754
	var v756 int32
	_ = v756
	var v757 int32
	_ = v757
	var v759 int32
	_ = v759
	var v760 int32
	_ = v760
	var v761 int32
	_ = v761
	var v762 int32
	_ = v762
	var v769 int32
	_ = v769
	var v775 int32
	_ = v775
	var v779 int32
	_ = v779
	var v781 int32
	_ = v781
	var v782 int32
	_ = v782
	var v786 int32
	_ = v786
	var v787 int32
	_ = v787
	var v789 int32
	_ = v789
	var v793 int32
	_ = v793
	var v795 int32
	_ = v795
	var v797 int32
	_ = v797
	var v800 int32
	_ = v800
	var v805 int32
	_ = v805
	var v806 int32
	_ = v806
	var v807 int32
	_ = v807
	var v809 int32
	_ = v809
	var v810 int32
	_ = v810
	var v812 int32
	_ = v812
	var v814 int32
	_ = v814
	var v817 int32
	_ = v817
	var v822 int32
	_ = v822
	var v823 int32
	_ = v823
	var v824 int32
	_ = v824
	var v831 int32
	_ = v831
	var v833 int32
	_ = v833
	var v834 int32
	_ = v834
	var v836 int32
	_ = v836
	var v848 int32
	_ = v848
	var v850 int32
	_ = v850
	var v859 int32
	_ = v859
	var v861 int32
	_ = v861
	var v865 int32
	_ = v865
	var v868 int32
	_ = v868
	var v872 int32
	_ = v872
	var v877 int32
	_ = v877
	v18 = m.G0
	v20 = v18 - int32(144)
	m.G0 = v20
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v23 < v22 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v46 = v42 + v43*int32(24)
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3))))
	if v47 != 0 {
		goto L15
	} else {
		goto L16
	}
L2:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v42 = v25
	goto L1
L3:
	;
	goto L4
L4:
	;
	if v22 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v40
	v42 = v40
	goto L1
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v22 << (uint(int32(1)) % 32)
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v32 = F_repalloc(m, v29, v22*int32(48))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	v34 = int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v34
	v38 = F_palloc_mul(m, int32(24), v34)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L9
	} else {
		goto L11
	}
L9:
	;
	return
L10:
	;
	v40 = v32
	goto L5
L11:
	;
	v40 = v38
	goto L5
L12:
	;
	v859 = v20 + int32(32)
	F_pg_regerror(m, v478, v859)
	mBase = m.M
	v861 = m.ExcPending
	if v861 != 0 {
		goto L9
	} else {
		goto L254
	}
L13:
	;
	v497 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	v501 = l2 << (uint(int32(1)) % 32)
	v504 = v497&int32(-255) | v501&int32(254)
	v505 = int32(28)
	if v501&v505 != 0 {
		goto L147
	} else {
		goto L148
	}
L14:
	;
	v57 = m.G0
	v59 = v57 - int32(16)
	m.G0 = v59
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3))))
	if v62 == int32(0) {
		v149 = int32(1)
		goto L20
	} else {
		goto L21
	}
L15:
	;
	if v47 != int32(46) {
		goto L14
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v46)+4)) = v51&int32(-769) | int32(256)
	goto L13
L18:
	;
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+1)))
	if v50 != 0 {
		goto L14
	} else {
		goto L19
	}
L19:
	;
	goto L17
L20:
	;
	m.G0 = v59 + int32(16)
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	v155 = v153 & int32(-769)
	if v149 != 0 {
		goto L52
	} else {
		goto L53
	}
L21:
	;
	v73 = int32(4)
	v74 = l3
	v76 = v62
	goto L22
L22:
	;
	switch v73 - int32(1) {
	case 0:
		goto L27
	default:
		goto L26
	case 3:
		goto L28
	}
L23:
	;
	v149 = base.B2i32(v125 == int32(4))
	goto L20
L24:
	;
	v126 = F_pg_mblen_cstr(m, v74)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L9
	} else {
		goto L50
	}
L25:
	;
	v125 = int32(1)
	goto L24
L26:
	;
	if v73&int32(-2) == int32(2) {
		goto L41
	} else {
		goto L42
	}
L27:
	;
	if v76 == int32(94) {
		goto L34
	} else {
		goto L35
	}
L28:
	;
	v85 = F_t_isalpha_cstr(m, v74)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L9
	} else {
		goto L29
	}
L29:
	;
	if v85 != 0 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v125 = int32(4)
	goto L24
L31:
	;
	goto L32
L32:
	;
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74))))
	if v88 == int32(91) {
		goto L25
	} else {
		goto L33
	}
L33:
	;
	v149 = int32(0)
	goto L20
L34:
	;
	v125 = int32(3)
	goto L24
L35:
	;
	goto L36
L36:
	;
	v95 = F_t_isalpha_cstr(m, v74)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L9
	} else {
		goto L37
	}
L37:
	;
	if v95 != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v125 = int32(2)
	goto L24
L39:
	;
	goto L40
L40:
	;
	v149 = int32(0)
	goto L20
L41:
	;
	v103 = F_t_isalpha_cstr(m, v74)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L9
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L9
	} else {
		goto L47
	}
L44:
	;
	if v103 != 0 {
		v125 = v73
		goto L24
	} else {
		goto L45
	}
L45:
	;
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74))))
	if v106 != int32(93) {
		v149 = int32(0)
		goto L20
	} else {
		goto L46
	}
L46:
	;
	v125 = int32(4)
	goto L24
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v59))) = int32(3)
	F_errmsg_internal(m, int32(_a_F_NIAddAffix_0), v59)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L9
	} else {
		goto L48
	}
L48:
	;
	F_errfinish(m, int32(_a_F_NIAddAffix_1), int32(66), int32(_a_F_NIAddAffix_2))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L9
	} else {
		goto L49
	}
L49:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L50:
	;
	v128 = v126 + v74
	v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128))))
	if v129 != 0 {
		v73 = v125
		v74 = v128
		v76 = v129
		goto L22
	} else {
		goto L51
	}
L51:
	;
	goto L23
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+4)) = v155 | int32(512)
	v160 = v46 + int32(16)
	v164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3))))
	if v164 != 0 {
		goto L55
	} else {
		goto L56
	}
L53:
	;
	goto L54
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+4)) = v155
	v448 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v449 = F_strlen(m, l3)
	mBase = m.M
	v452 = F_MemoryContextAlloc(m, v448, v449+int32(3))
	mBase = m.M
	v453 = m.ExcPending
	if v453 != 0 {
		goto L9
	} else {
		goto L137
	}
L55:
	;
	v165 = l3
	goto L57
L56:
	;
	v165 = int32(_a_F_NIAddAffix_3)
	goto L57
L57:
	;
	v166 = m.G0
	v168 = v166 - int32(80)
	m.G0 = v168
	v170 = F_strlen(m, v165)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v160))) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v160)+4)) = base.B2i32(l6 != int32(0))
	v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165))))
	if v174 != 0 {
		goto L60
	} else {
		goto L61
	}
L58:
	;
	goto L13
L59:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v437 = m.ExcPending
	if v437 != 0 {
		goto L9
	} else {
		goto L134
	}
L60:
	;
	v176 = v170 + int32(9)
	v182 = int32(0)
	v186 = v165
	v187 = v174
	v191 = int32(4)
	goto L63
L61:
	;
	goto L62
L62:
	;
	v386 = *(*int32)(unsafe.Add(mBase, uint32(v160)))
	if v386 != 0 {
		goto L128
	} else {
		goto L129
	}
L63:
	;
	switch v191 - int32(1) {
	case 0:
		goto L67
	default:
		goto L66
	case 3:
		goto L68
	}
L64:
	;
	if v361 != int32(4) {
		goto L59
	} else {
		goto L127
	}
L65:
	;
	v363 = F_pg_mblen_cstr(m, v186)
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L9
	} else {
		goto L125
	}
L66:
	;
	if v191&int32(-2) == int32(2) {
		goto L107
	} else {
		goto L108
	}
L67:
	;
	if v187&int32(255) == int32(94) {
		goto L93
	} else {
		goto L94
	}
L68:
	;
	v198 = F_t_isalpha_cstr(m, v186)
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L9
	} else {
		goto L69
	}
L69:
	;
	if v198 != 0 {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v200 = F_palloc0(m, v176)
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L9
	} else {
		goto L73
	}
L71:
	;
	goto L72
L72:
	;
	v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v186))))
	if v225 == int32(91) {
		goto L82
	} else {
		goto L83
	}
L73:
	;
	if v182 != 0 {
		goto L75
	} else {
		goto L76
	}
L74:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v200)))
	*(*int32)(unsafe.Add(mBase, uint32(v200))) = v204&int32(-4) | int32(1)
	v210 = F_pg_mblen_cstr(m, v186)
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L9
	} else {
		goto L78
	}
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v182)+4)) = v200
	goto L74
L76:
	;
	goto L77
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v160))) = v200
	goto L74
L78:
	;
	if v210 != 0 {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	base.MemoryCopy(m, v200+int32(8), v186, v210)
	goto L81
L80:
	;
	goto L81
L81:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v200)))
	*(*int32)(unsafe.Add(mBase, uint32(v200))) = v215&int32(-262141) | v210<<(uint(int32(2))%32)&int32(_a_F_NIAddAffix_4)
	v359 = v200
	v361 = int32(4)
	goto L65
L82:
	;
	v228 = F_palloc0(m, v176)
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L9
	} else {
		goto L85
	}
L83:
	;
	goto L84
L84:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L9
	} else {
		goto L90
	}
L85:
	;
	if v182 != 0 {
		goto L87
	} else {
		goto L88
	}
L86:
	;
	v232 = int32(1)
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v228)))
	*(*int32)(unsafe.Add(mBase, uint32(v228))) = v233&int32(-4) | v232
	v359 = v228
	v361 = v232
	goto L65
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v182)+4)) = v228
	goto L86
L88:
	;
	goto L89
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v160))) = v228
	goto L86
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v168)+48)) = v165
	F_errmsg_internal(m, int32(_a_F_NIAddAffix_5), v168+int32(48))
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L9
	} else {
		goto L91
	}
L91:
	;
	F_errfinish(m, int32(_a_F_NIAddAffix_1), int32(118), int32(_a_F_NIAddAffix_6))
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L9
	} else {
		goto L92
	}
L92:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L93:
	;
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v182)))
	*(*int32)(unsafe.Add(mBase, uint32(v182))) = v258&int32(-4) | int32(2)
	v359 = v182
	v361 = int32(3)
	goto L65
L94:
	;
	goto L95
L95:
	;
	v265 = F_t_isalpha_cstr(m, v186)
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L9
	} else {
		goto L96
	}
L96:
	;
	if v265 != 0 {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	v267 = F_pg_mblen_cstr(m, v186)
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L9
	} else {
		goto L100
	}
L98:
	;
	goto L99
L99:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L9
	} else {
		goto L104
	}
L100:
	;
	if v267 != 0 {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	base.MemoryCopy(m, v182+int32(8), v186, v267)
	goto L103
L102:
	;
	goto L103
L103:
	;
	v272 = int32(2)
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v182)))
	*(*int32)(unsafe.Add(mBase, uint32(v182))) = v273&int32(-262141) | v267<<(uint(v272)%32)&int32(_a_F_NIAddAffix_4)
	v359 = v182
	v361 = v272
	goto L65
L104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v168)+64)) = v165
	F_errmsg_internal(m, int32(_a_F_NIAddAffix_5), v168-int32(-64))
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L9
	} else {
		goto L105
	}
L105:
	;
	F_errfinish(m, int32(_a_F_NIAddAffix_1), int32(133), int32(_a_F_NIAddAffix_6))
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L9
	} else {
		goto L106
	}
L106:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L107:
	;
	v301 = F_t_isalpha_cstr(m, v186)
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L9
	} else {
		goto L110
	}
L108:
	;
	goto L109
L109:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L9
	} else {
		goto L122
	}
L110:
	;
	if v301 != 0 {
		goto L111
	} else {
		goto L112
	}
L111:
	;
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v182)))
	v304 = F_pg_mblen_cstr(m, v186)
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L9
	} else {
		goto L114
	}
L112:
	;
	goto L113
L113:
	;
	v325 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v186))))
	if v325 == int32(93) {
		v359 = v182
		v361 = int32(4)
		goto L65
	} else {
		goto L118
	}
L114:
	;
	if v304 != 0 {
		goto L115
	} else {
		goto L116
	}
L115:
	;
	base.MemoryCopy(m, v182+int32(base.Ui32(v303)>>(uint(int32(2))%32))&int32(_a_F_NIAddAffix_7)+int32(8), v186, v304)
	goto L117
L116:
	;
	goto L117
L117:
	;
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v182)))
	*(*int32)(unsafe.Add(mBase, uint32(v182))) = (v314+v304<<(uint(int32(2))%32))&int32(_a_F_NIAddAffix_4) | v314&int32(-262141)
	v359 = v182
	v361 = v191
	goto L65
L118:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L9
	} else {
		goto L119
	}
L119:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v168)+16)) = v165
	F_errmsg_internal(m, int32(_a_F_NIAddAffix_5), v168+int32(16))
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L9
	} else {
		goto L120
	}
L120:
	;
	F_errfinish(m, int32(_a_F_NIAddAffix_1), int32(142), int32(_a_F_NIAddAffix_6))
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L9
	} else {
		goto L121
	}
L121:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v168)+32)) = int32(3)
	F_errmsg_internal(m, int32(_a_F_NIAddAffix_8), v168+int32(32))
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L9
	} else {
		goto L123
	}
L123:
	;
	F_errfinish(m, int32(_a_F_NIAddAffix_1), int32(145), int32(_a_F_NIAddAffix_6))
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L9
	} else {
		goto L124
	}
L124:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L125:
	;
	v365 = v363 + v186
	v366 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v365))))
	if v366 != 0 {
		v182 = v359
		v186 = v365
		v187 = v366
		v191 = v361
		goto L63
	} else {
		goto L126
	}
L126:
	;
	goto L64
L127:
	;
	goto L62
L128:
	;
	v387 = *(*int32)(unsafe.Add(mBase, uint32(v160)+4))
	v391 = v386
	v395 = v387
	goto L131
L129:
	;
	goto L130
L130:
	;
	m.G0 = v168 + int32(80)
	goto L58
L131:
	;
	v411 = (v395+int32(2))&int32(_a_F_NIAddAffix_9) | v395&int32(-131071)
	*(*int32)(unsafe.Add(mBase, uint32(v160)+4)) = v411
	v413 = *(*int32)(unsafe.Add(mBase, uint32(v391)+4))
	if v413 != 0 {
		v391 = v413
		v395 = v411
		goto L131
	} else {
		goto L133
	}
L132:
	;
	goto L130
L133:
	;
	goto L132
L134:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v168))) = v165
	F_errmsg_internal(m, int32(_a_F_NIAddAffix_5), v168)
	mBase = m.M
	v441 = m.ExcPending
	if v441 != 0 {
		goto L9
	} else {
		goto L135
	}
L135:
	;
	F_errfinish(m, int32(_a_F_NIAddAffix_1), int32(150), int32(_a_F_NIAddAffix_6))
	mBase = m.M
	v446 = m.ExcPending
	if v446 != 0 {
		goto L9
	} else {
		goto L136
	}
L136:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L137:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = l3
	if l6 != 0 {
		goto L138
	} else {
		goto L139
	}
L138:
	;
	v457 = int32(_a_F_NIAddAffix_10)
	goto L140
L139:
	;
	v457 = int32(_a_F_NIAddAffix_11)
	goto L140
L140:
	;
	v460 = F_pg_sprintf(m, v452, v457, v20+int32(16))
	mBase = m.M
	v461 = m.ExcPending
	if v461 != 0 {
		goto L9
	} else {
		goto L141
	}
L141:
	;
	v462 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v463 = F_strlen(m, v452)
	mBase = m.M
	v468 = F_MemoryContextAlloc(m, v462, v463<<(uint(int32(2))%32)+int32(4))
	mBase = m.M
	v469 = m.ExcPending
	if v469 != 0 {
		goto L9
	} else {
		goto L142
	}
L142:
	;
	v470 = F_pg_mb2wchar_with_len(m, v452, v468, v463)
	mBase = m.M
	v471 = m.ExcPending
	if v471 != 0 {
		goto L9
	} else {
		goto L143
	}
L143:
	;
	v473 = F_palloc(m, int32(32))
	mBase = m.M
	v474 = m.ExcPending
	if v474 != 0 {
		goto L9
	} else {
		goto L144
	}
L144:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+16)) = v473
	v478 = F_pg_regcomp(m, v473, v468, v470, int32(19), int32(100))
	mBase = m.M
	v479 = m.ExcPending
	if v479 != 0 {
		goto L9
	} else {
		goto L145
	}
L145:
	;
	if v478 != 0 {
		goto L12
	} else {
		goto L146
	}
L146:
	;
	goto L13
L147:
	;
	v509 = v504
	goto L149
L148:
	;
	v509 = v504 | v505
	goto L149
L149:
	;
	if v501&int32(34) != 0 {
		goto L150
	} else {
		goto L151
	}
L150:
	;
	v512 = v509
	goto L152
L151:
	;
	v512 = v504
	goto L152
L152:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+4)) = v512
	v514 = F_strlen(m, l1)
	mBase = m.M
	v516 = v514 + int32(1)
	if base.Ui32(int32(1025)) <= base.Ui32(v516) {
		goto L154
	} else {
		goto L155
	}
L153:
	;
	if (l1^v539)&int32(3) != 0 {
		goto L166
	} else {
		goto L167
	}
L154:
	;
	v519 = F_palloc0(m, v516)
	mBase = m.M
	v520 = m.ExcPending
	if v520 != 0 {
		goto L9
	} else {
		goto L157
	}
L155:
	;
	goto L156
L156:
	;
	v524 = (v514 + int32(8)) & int32(4088)
	v525 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	if base.Ui32(v524) <= base.Ui32(v525) {
		goto L159
	} else {
		goto L160
	}
L157:
	;
	v539 = v519
	goto L153
L158:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = v532 - v524
	*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = v524 + v533
	v539 = v533
	goto L153
L159:
	;
	v527 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	v532 = v525
	v533 = v527
	goto L158
L160:
	;
	goto L161
L161:
	;
	v528 = int32(_a_F_NIAddAffix_12)
	v530 = F_palloc0(m, v528)
	mBase = m.M
	v531 = m.ExcPending
	if v531 != 0 {
		goto L9
	} else {
		goto L162
	}
L162:
	;
	v532 = v528
	v533 = v530
	goto L158
L163:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46))) = v539
	v616 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	v619 = v616&int32(-2) | l6
	*(*int32)(unsafe.Add(mBase, uint32(v46)+4)) = v619
	if l4 == int32(0) {
		goto L185
	} else {
		goto L186
	}
L164:
	;
	goto L163
L165:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v595))) = uint8(v594)
	if v594&int32(255) == int32(0) {
		goto L164
	} else {
		goto L180
	}
L166:
	;
	v546 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	v593 = l1
	v594 = v546
	v595 = v539
	goto L165
L167:
	;
	goto L168
L168:
	;
	if l1&int32(3) != 0 {
		goto L169
	} else {
		goto L170
	}
L169:
	;
	v550 = l1
	v552 = v539
	goto L172
L170:
	;
	v564 = l1
	v566 = v539
	goto L171
L171:
	;
	v568 = *(*int32)(unsafe.Add(mBase, uint32(v564)))
	v571 = int32(-2139062144)
	if (int32(16843008)-v568|v568)&v571 != v571 {
		v593 = v564
		v594 = v568
		v595 = v566
		goto L165
	} else {
		goto L176
	}
L172:
	;
	v553 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v550))))
	*(*uint8)(unsafe.Add(mBase, uint32(v552))) = uint8(v553)
	if v553 == int32(0) {
		goto L164
	} else {
		goto L174
	}
L173:
	;
	v564 = v560
	v566 = v558
	goto L171
L174:
	;
	v557 = int32(1)
	v558 = v552 + v557
	v560 = v550 + v557
	if v560&int32(3) != 0 {
		v550 = v560
		v552 = v558
		goto L172
	} else {
		goto L175
	}
L175:
	;
	goto L173
L176:
	;
	v576 = v564
	v577 = v568
	v578 = v566
	goto L177
L177:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v578))) = v577
	v580 = int32(4)
	v581 = v578 + v580
	v583 = v576 + v580
	v585 = *(*int32)(unsafe.Add(mBase, uint32(v576)+4))
	v588 = int32(-2139062144)
	if (int32(16843008)-v585|v585)&v588 == v588 {
		v576 = v583
		v577 = v585
		v578 = v581
		goto L177
	} else {
		goto L179
	}
L178:
	;
	v593 = v583
	v594 = v585
	v595 = v581
	goto L165
L179:
	;
	goto L178
L180:
	;
	v602 = v593
	v604 = v595
	goto L181
L181:
	;
	v605 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v602)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v604)+1)) = uint8(v605)
	v607 = int32(1)
	if v605 != 0 {
		v602 = v602 + v607
		v604 = v604 + v607
		goto L181
	} else {
		goto L183
	}
L182:
	;
	goto L164
L183:
	;
	goto L182
L184:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+8)) = v731
	v736 = F_strlen(m, l5)
	mBase = m.M
	v738 = v736 & int32(_a_F_NIAddAffix_13)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+4)) = v730&int32(-16776193) | v738<<(uint(int32(10))%32)
	if v738 != 0 {
		goto L220
	} else {
		goto L221
	}
L185:
	;
	v730 = v619
	v731 = int32(_a_F_NIAddAffix_3)
	goto L184
L186:
	;
	goto L187
L187:
	;
	v625 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4))))
	if v625 == int32(0) {
		v730 = v619
		v731 = int32(_a_F_NIAddAffix_3)
		goto L184
	} else {
		goto L188
	}
L188:
	;
	v628 = F_strlen(m, l4)
	mBase = m.M
	v630 = v628 + int32(1)
	if base.Ui32(int32(1025)) <= base.Ui32(v630) {
		goto L190
	} else {
		goto L191
	}
L189:
	;
	if (l4^v652)&int32(3) != 0 {
		goto L202
	} else {
		goto L203
	}
L190:
	;
	v633 = F_palloc0(m, v630)
	mBase = m.M
	v634 = m.ExcPending
	if v634 != 0 {
		goto L9
	} else {
		goto L193
	}
L191:
	;
	goto L192
L192:
	;
	v638 = (v628 + int32(8)) & int32(4088)
	v639 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	if base.Ui32(v638) <= base.Ui32(v639) {
		goto L195
	} else {
		goto L196
	}
L193:
	;
	v652 = v633
	goto L189
L194:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = v646 - v638
	*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = v647 + v638
	v652 = v647
	goto L189
L195:
	;
	v641 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	v646 = v639
	v647 = v641
	goto L194
L196:
	;
	goto L197
L197:
	;
	v642 = int32(_a_F_NIAddAffix_12)
	v644 = F_palloc0(m, v642)
	mBase = m.M
	v645 = m.ExcPending
	if v645 != 0 {
		goto L9
	} else {
		goto L198
	}
L198:
	;
	v646 = v642
	v647 = v644
	goto L194
L199:
	;
	v729 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	v730 = v729
	v731 = v652
	goto L184
L200:
	;
	goto L199
L201:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v709))) = uint8(v708)
	if v708&int32(255) == int32(0) {
		goto L200
	} else {
		goto L216
	}
L202:
	;
	v660 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4))))
	v707 = l4
	v708 = v660
	v709 = v652
	goto L201
L203:
	;
	goto L204
L204:
	;
	if l4&int32(3) != 0 {
		goto L205
	} else {
		goto L206
	}
L205:
	;
	v664 = l4
	v666 = v652
	goto L208
L206:
	;
	v678 = l4
	v680 = v652
	goto L207
L207:
	;
	v682 = *(*int32)(unsafe.Add(mBase, uint32(v678)))
	v685 = int32(-2139062144)
	if (int32(16843008)-v682|v682)&v685 != v685 {
		v707 = v678
		v708 = v682
		v709 = v680
		goto L201
	} else {
		goto L212
	}
L208:
	;
	v667 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v664))))
	*(*uint8)(unsafe.Add(mBase, uint32(v666))) = uint8(v667)
	if v667 == int32(0) {
		goto L200
	} else {
		goto L210
	}
L209:
	;
	v678 = v674
	v680 = v672
	goto L207
L210:
	;
	v671 = int32(1)
	v672 = v666 + v671
	v674 = v664 + v671
	if v674&int32(3) != 0 {
		v664 = v674
		v666 = v672
		goto L208
	} else {
		goto L211
	}
L211:
	;
	goto L209
L212:
	;
	v690 = v678
	v691 = v682
	v692 = v680
	goto L213
L213:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v692))) = v691
	v694 = int32(4)
	v695 = v692 + v694
	v697 = v690 + v694
	v699 = *(*int32)(unsafe.Add(mBase, uint32(v690)+4))
	v702 = int32(-2139062144)
	if (int32(16843008)-v699|v699)&v702 == v702 {
		v690 = v697
		v691 = v699
		v692 = v695
		goto L213
	} else {
		goto L215
	}
L214:
	;
	v707 = v697
	v708 = v699
	v709 = v695
	goto L201
L215:
	;
	goto L214
L216:
	;
	v716 = v707
	v718 = v709
	goto L217
L217:
	;
	v719 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v716)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v718)+1)) = uint8(v719)
	v721 = int32(1)
	if v719 != 0 {
		v716 = v716 + v721
		v718 = v718 + v721
		goto L217
	} else {
		goto L219
	}
L218:
	;
	goto L200
L219:
	;
	goto L218
L220:
	;
	v743 = F_strlen(m, l5)
	mBase = m.M
	v745 = v743 + int32(1)
	if base.Ui32(int32(1025)) <= base.Ui32(v745) {
		goto L224
	} else {
		goto L225
	}
L221:
	;
	v848 = int32(_a_F_NIAddAffix_3)
	goto L222
L222:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+12)) = v848
	v850 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v850 + int32(1)
	m.G0 = v20 + int32(144)
	return
L223:
	;
	if (l5^v769)&int32(3) != 0 {
		goto L236
	} else {
		goto L237
	}
L224:
	;
	v748 = F_palloc0(m, v745)
	mBase = m.M
	v749 = m.ExcPending
	if v749 != 0 {
		goto L9
	} else {
		goto L227
	}
L225:
	;
	goto L226
L226:
	;
	v753 = (v743 + int32(8)) & int32(4088)
	v754 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	if base.Ui32(v753) <= base.Ui32(v754) {
		goto L229
	} else {
		goto L230
	}
L227:
	;
	v769 = v748
	goto L223
L228:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = v761 - v753
	*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = v753 + v762
	v769 = v762
	goto L223
L229:
	;
	v756 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	v761 = v754
	v762 = v756
	goto L228
L230:
	;
	goto L231
L231:
	;
	v757 = int32(_a_F_NIAddAffix_12)
	v759 = F_palloc0(m, v757)
	mBase = m.M
	v760 = m.ExcPending
	if v760 != 0 {
		goto L9
	} else {
		goto L232
	}
L232:
	;
	v761 = v757
	v762 = v759
	goto L228
L233:
	;
	v848 = v769
	goto L222
L234:
	;
	goto L233
L235:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v824))) = uint8(v823)
	if v823&int32(255) == int32(0) {
		goto L234
	} else {
		goto L250
	}
L236:
	;
	v775 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l5))))
	v822 = l5
	v823 = v775
	v824 = v769
	goto L235
L237:
	;
	goto L238
L238:
	;
	if l5&int32(3) != 0 {
		goto L239
	} else {
		goto L240
	}
L239:
	;
	v779 = l5
	v781 = v769
	goto L242
L240:
	;
	v793 = l5
	v795 = v769
	goto L241
L241:
	;
	v797 = *(*int32)(unsafe.Add(mBase, uint32(v793)))
	v800 = int32(-2139062144)
	if (int32(16843008)-v797|v797)&v800 != v800 {
		v822 = v793
		v823 = v797
		v824 = v795
		goto L235
	} else {
		goto L246
	}
L242:
	;
	v782 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v779))))
	*(*uint8)(unsafe.Add(mBase, uint32(v781))) = uint8(v782)
	if v782 == int32(0) {
		goto L234
	} else {
		goto L244
	}
L243:
	;
	v793 = v789
	v795 = v787
	goto L241
L244:
	;
	v786 = int32(1)
	v787 = v781 + v786
	v789 = v779 + v786
	if v789&int32(3) != 0 {
		v779 = v789
		v781 = v787
		goto L242
	} else {
		goto L245
	}
L245:
	;
	goto L243
L246:
	;
	v805 = v793
	v806 = v797
	v807 = v795
	goto L247
L247:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v807))) = v806
	v809 = int32(4)
	v810 = v807 + v809
	v812 = v805 + v809
	v814 = *(*int32)(unsafe.Add(mBase, uint32(v805)+4))
	v817 = int32(-2139062144)
	if (int32(16843008)-v814|v814)&v817 == v817 {
		v805 = v812
		v806 = v814
		v807 = v810
		goto L247
	} else {
		goto L249
	}
L248:
	;
	v822 = v812
	v823 = v814
	v824 = v810
	goto L235
L249:
	;
	goto L248
L250:
	;
	v831 = v822
	v833 = v824
	goto L251
L251:
	;
	v834 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v831)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v833)+1)) = uint8(v834)
	v836 = int32(1)
	if v834 != 0 {
		v831 = v831 + v836
		v833 = v833 + v836
		goto L251
	} else {
		goto L253
	}
L252:
	;
	goto L234
L253:
	;
	goto L252
L254:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v865 = m.ExcPending
	if v865 != 0 {
		goto L9
	} else {
		goto L255
	}
L255:
	;
	F_errcode(m, int32(302252162))
	mBase = m.M
	v868 = m.ExcPending
	if v868 != 0 {
		goto L9
	} else {
		goto L256
	}
L256:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v859
	F_errmsg(m, int32(_a_F_NIAddAffix_14), v20)
	mBase = m.M
	v872 = m.ExcPending
	if v872 != 0 {
		goto L9
	} else {
		goto L257
	}
L257:
	;
	F_errfinish(m, int32(_a_F_NIAddAffix_15), int32(752), int32(_a_F_NIAddAffix_16))
	mBase = m.M
	v877 = m.ExcPending
	if v877 != 0 {
		goto L9
	} else {
		goto L258
	}
L258:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_NeedsUpdated(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v88 int64
	_ = v88
	var v89 int64
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int64
	_ = v97
	var v98 int64
	_ = v98
	var v103 int64
	_ = v103
	var v108 int64
	_ = v108
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v144 int32
	_ = v144
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	v3 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)+84))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v18 = F_ReadBufferExtended(m, v13, v3, v15, v3, v17)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	F_LockBufferInternal(m, v18, int32(1))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v25 = int32(0)
	if v18 < v25 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	F_UnlockReleaseBuffer(m, v18)
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L1
	} else {
		goto L28
	}
L5:
	;
	v44 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+82)))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v43+v44<<(uint(int32(2))%32))+20))
	v51 = v48&int32(_a_F_NeedsUpdated_0) + v43
	v52 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v51)+2)))
	if v52 == int32(0) {
		v179 = v25
		goto L4
	} else {
		goto L9
	}
L6:
	;
	v29 = *(*int32)(unsafe.Add(mBase, _c_F_NeedsUpdated[0]))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v29+(v18^int32(-1))<<(uint(int32(2))%32))))
	v43 = v35
	goto L5
L7:
	;
	goto L8
L8:
	;
	v37 = *(*int32)(unsafe.Add(mBase, _c_F_NeedsUpdated[1]))
	v43 = v37 + v18<<(uint(int32(13))%32) + int32(-8192)
	goto L5
L9:
	;
	v59 = int32(0)
	v60 = v52
	goto L10
L10:
	;
	v68 = v51 + int32(4) + v59*int32(6)
	v69 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v68)+4)))
	if v69 != 0 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v157 = int32(0)
	if v153 == v157 {
		v179 = v157
		goto L4
	} else {
		goto L26
	}
L12:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v71 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v68)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v11)+12)) = uint16(v71)
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v73
	v77 = v11 + int32(8)
	v82 = m.G0
	v84 = v82 - int32(16)
	m.G0 = v84
	v87 = v11 + int32(12)
	v88 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v87))))
	v89 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v77))))
	v90 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v87))))
	*(*uint16)(unsafe.Add(mBase, uint32(v84)+12)) = uint16(v90)
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v77)))
	*(*int32)(unsafe.Add(mBase, uint32(v84)+8)) = v92
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v70)+20))
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v70)+12))
	v97 = v88 << (uint(int64(32)) % 64)
	v98 = int64(33)
	v103 = (int64(base.Ui64(v97)>>(uint(v98)%64)) ^ (v97 | v89)) * int64(-49064778989728563)
	v108 = (int64(base.Ui64(v103)>>(uint(v98)%64)) ^ v103) * int64(-4265267296055464877)
	v113 = v95 & base.I32_wrap_i64(int64(base.Ui64(v108)>>(uint(v98)%64))^v108)
	v116 = v94 + v113<<(uint(int32(3))%32)
	v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116)+6)))
	if v117 != 0 {
		goto L17
	} else {
		goto L18
	}
L13:
	;
	v153 = v60
	goto L14
L14:
	;
	v155 = v59 + int32(1)
	if base.Ui32(v155) < base.Ui32(v153) {
		v59 = v155
		v60 = v153
		goto L10
	} else {
		goto L25
	}
L15:
	;
	if v144 != 0 {
		v179 = int32(1)
		goto L4
	} else {
		goto L24
	}
L16:
	;
	m.G0 = v84 + int32(16)
	goto L15
L17:
	;
	v119 = v116
	v123 = v113
	goto L20
L18:
	;
	goto L19
L19:
	;
	v144 = int32(0)
	goto L16
L20:
	;
	v126 = F_ItemPointerEquals(m, v119, v84+int32(8))
	mBase = m.M
	if v126 != 0 {
		v144 = v119
		goto L16
	} else {
		goto L22
	}
L21:
	;
	goto L19
L22:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v70)+20))
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v70)+12))
	v131 = v128 & (v123 + int32(1))
	v134 = v127 + v131<<(uint(int32(3))%32)
	v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v134)+6)))
	if v135 != 0 {
		v119 = v134
		v123 = v131
		goto L20
	} else {
		goto L23
	}
L23:
	;
	goto L21
L24:
	;
	v152 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v51)+2)))
	v153 = v152
	goto L14
L25:
	;
	goto L11
L26:
	;
	v165 = v51 + v153*int32(6) - int32(2)
	if v165 == int32(0) {
		v179 = int32(1)
		goto L4
	} else {
		goto L27
	}
L27:
	;
	v168 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v165)+4)))
	v179 = base.B2i32(v168 == int32(0))
	goto L4
L28:
	;
	m.G0 = v11 + int32(16)
	return v179
}
func F_NullCommand(m *base.Module, l0 int32) {
	var v2 int32
	_ = v2
	var v8 int32
	_ = v8
	v2 = int32(2)
	if base.Ui32(l0-v2) <= base.Ui32(v2) {
		F_pq_putemptymessage(m, int32(73))
		v8 = m.ExcPending
		if v8 != 0 {
			return
		} else {
			return
		}
	} else {
		return
	}
}
func F___nl_langinfo_l(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
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
	if l0 == int32(14) {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
		if v9 != 0 {
			v10 = int32(_a_F___nl_langinfo_l_0)
		} else {
			v10 = int32(_a_F___nl_langinfo_l_1)
		}
		return v10
	} else {
		v12 = int32(_a_F___nl_langinfo_l_2)
		v13 = l0 & v12
		v17 = l0 >> (uint(int32(16)) % 32)
		if base.B2i32(v13 != v12)|base.B2i32(int32(5) < v17) == int32(0) {
			v26 = *(*int32)(unsafe.Add(mBase, uint32(l1+v17<<(uint(int32(2))%32))))
			if v26 != 0 {
				v30 = v26 + int32(8)
			} else {
				v30 = int32(_a_F___nl_langinfo_l_3)
			}
			return v30
		} else {
			v32 = int32(_a_F___nl_langinfo_l_4)
			switch v17 - int32(1) {
			case 0:
				if base.Ui32(int32(1)) < base.Ui32(v13) {
					v56 = v32
				} else {
					v44 = int32(_a_F___nl_langinfo_l_5)
					if v13 == int32(0) {
						v56 = v44
					} else {
						v47 = v44
						v49 = v13
						for {
							v52 = v47 + int32(1)
							v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47))))
							if v53 != 0 {
								v47 = v52
								continue
							} else {
							}
							v55 = v49 - int32(1)
							if v55 != 0 {
								v47 = v52
								v49 = v55
								continue
							} else {
								break
							}
							break
						}
						v56 = v52
					}
				}
			case 1:
				if base.Ui32(int32(49)) < base.Ui32(v13) {
					v56 = v32
				} else {
					v44 = int32(_a_F___nl_langinfo_l_6)
					if v13 == int32(0) {
						v56 = v44
					} else {
						v47 = v44
						v49 = v13
						for {
							v52 = v47 + int32(1)
							v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47))))
							if v53 != 0 {
								v47 = v52
								continue
							} else {
							}
							v55 = v49 - int32(1)
							if v55 != 0 {
								v47 = v52
								v49 = v55
								continue
							} else {
								break
							}
							break
						}
						v56 = v52
					}
				}
			default:
				v56 = v32
			case 4:
				if base.Ui32(int32(3)) < base.Ui32(v13) {
					v56 = v32
				} else {
					v44 = int32(_a_F___nl_langinfo_l_7)
					if v13 == int32(0) {
						v56 = v44
					} else {
						v47 = v44
						v49 = v13
						for {
							v52 = v47 + int32(1)
							v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47))))
							if v53 != 0 {
								v47 = v52
								continue
							} else {
							}
							v55 = v49 - int32(1)
							if v55 != 0 {
								v47 = v52
								v49 = v55
								continue
							} else {
								break
							}
							break
						}
						v56 = v52
					}
				}
			}
			return v56
		}
	}
}
func F_namefastcmp_locale(m *base.Module, l0 int64, l1 int64, l2 int32) int32 {
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	v6 = base.I32_wrap_i64(l1)
	v7 = base.I32_wrap_i64(l0)
	v8 = F_strlen(m, v7)
	v9 = F_strlen(m, v6)
	v10 = F_varstrfastcmp_locale(m, v7, v8, v6, v9, l2)
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		return v10
	}
}
func F_namelike(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
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
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v9 = F_pg_detoast_datum_packed(m, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int64(0)
	} else {
		v13 = F_strlen(m, v7)
		mBase = m.M
		v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9))))
		if v14 == int32(1) {
			v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+1)))
			if v20 == int32(18) {
				v23 = int32(16)
			} else {
				v23 = int32(0)
			}
			if base.Ui32((v20-int32(1))&int32(255)) < base.Ui32(int32(3)) {
				v30 = int32(4)
			} else {
				v30 = v23
			}
			v43 = v30
		} else {
			v31 = int32(1)
			if v14&v31 != 0 {
				v43 = int32(base.Ui32(v14)>>(uint(v31)%32)) - v31
			} else {
				v37 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
				v43 = int32(base.Ui32(v37)>>(uint(int32(2))%32)) - int32(4)
			}
		}
		v44 = int32(1)
		if v14&v44 != 0 {
			v48 = v44
		} else {
			v48 = int32(4)
		}
		v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v51 = F_GenericMatchText(m, v7, v13, v9+v48, v43, v50)
		mBase = m.M
		v52 = m.ExcPending
		if v52 != 0 {
			return int64(0)
		} else {
			return base.I64_extend_i32_u(base.B2i32(v51 == int32(1)))
		}
	}
}
func F_nameregexeq(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14330(m, l0, int32(19))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return v3
	}
}
func F_namestrcpy(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
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
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	v3 = int32(64)
	if (l1^l0)&int32(3) != 0 {
		v74 = l1
		v75 = v3
		v76 = l0
		goto L5
	} else {
		goto L6
	}
L1:
	;
	v113 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+63)) = uint8(v113)
	return
L2:
	;
	F___memset(m, v109, int32(0), v108)
	mBase = m.M
	goto L1
L3:
	;
	v108 = int32(0)
	v109 = v103
	goto L2
L4:
	;
	v86 = v81
	v87 = v82
	v88 = v83
	goto L22
L5:
	;
	if v75 == int32(0) {
		v103 = v76
		goto L3
	} else {
		goto L21
	}
L6:
	;
	if base.B2i32(l1&int32(3) == int32(0))|int32(0) != 0 {
		v40 = l1
		v41 = v3
		v42 = l0
		v43 = int32(1)
		goto L7
	} else {
		goto L8
	}
L7:
	;
	if v43 == int32(0) {
		v103 = v42
		goto L3
	} else {
		goto L14
	}
L8:
	;
	v19 = l1
	v20 = v3
	v21 = l0
	goto L9
L9:
	;
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
	*(*uint8)(unsafe.Add(mBase, uint32(v21))) = uint8(v23)
	if v23 == int32(0) {
		v108 = v20
		v109 = v21
		goto L2
	} else {
		goto L11
	}
L10:
	;
	v40 = v34
	v41 = v30
	v42 = v28
	v43 = v32
	goto L7
L11:
	;
	v27 = int32(1)
	v28 = v21 + v27
	v30 = v20 - v27
	v31 = int32(0)
	v32 = base.B2i32(v30 != v31)
	v34 = v19 + v27
	if v34&int32(3) == v31 {
		v40 = v34
		v41 = v30
		v42 = v28
		v43 = v32
		goto L7
	} else {
		goto L12
	}
L12:
	;
	if v30 != 0 {
		v19 = v34
		v20 = v30
		v21 = v28
		goto L9
	} else {
		goto L13
	}
L13:
	;
	goto L10
L14:
	;
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40))))
	if v46 == int32(0) {
		v108 = v41
		v109 = v42
		goto L2
	} else {
		goto L15
	}
L15:
	;
	if base.Ui32(v41) < base.Ui32(int32(4)) {
		v74 = v40
		v75 = v41
		v76 = v42
		goto L5
	} else {
		goto L16
	}
L16:
	;
	v52 = v40
	v53 = v41
	v54 = v42
	goto L17
L17:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
	v60 = int32(-2139062144)
	if (int32(16843008)-v57|v57)&v60 != v60 {
		v81 = v52
		v82 = v53
		v83 = v54
		goto L4
	} else {
		goto L19
	}
L18:
	;
	v74 = v68
	v75 = v70
	v76 = v66
	goto L5
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54))) = v57
	v65 = int32(4)
	v66 = v54 + v65
	v68 = v52 + v65
	v70 = v53 - v65
	if base.Ui32(int32(3)) < base.Ui32(v70) {
		v52 = v68
		v53 = v70
		v54 = v66
		goto L17
	} else {
		goto L20
	}
L20:
	;
	goto L18
L21:
	;
	v81 = v74
	v82 = v75
	v83 = v76
	goto L4
L22:
	;
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86))))
	*(*uint8)(unsafe.Add(mBase, uint32(v88))) = uint8(v90)
	if v90 == int32(0) {
		v108 = v87
		v109 = v88
		goto L2
	} else {
		goto L24
	}
L23:
	;
	v103 = v95
	goto L3
L24:
	;
	v94 = int32(1)
	v95 = v88 + v94
	v99 = v87 - v94
	if v99 != 0 {
		v86 = v86 + v94
		v87 = v99
		v88 = v95
		goto L22
	} else {
		goto L25
	}
L25:
	;
	goto L23
}
func F_ndistinct_array_end(m *base.Module, l0 int32) int32 {
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	v10 = Fn14254(m, l0, int32(_a_F_ndistinct_array_end_0), int32(286), int32(_a_F_ndistinct_array_end_1), int32(_a_F_ndistinct_array_end_2), int32(275), int32(_a_F_ndistinct_array_end_3), int32(6), int32(261))
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		return v10
	}
}
func F_negate_clause(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int64
	_ = v21
	var v25 int32
	_ = v25
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
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
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
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v215 int32
	_ = v215
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v226 int32
	_ = v226
	var v231 int32
	_ = v231
	v2 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	if l0 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L16
	} else {
		goto L67
	}
L2:
	;
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch v9 - int32(7) {
	case 0:
		goto L12
	default:
		goto L6
	case 10:
		goto L11
	case 13:
		goto L10
	case 14:
		goto L9
	case 45:
		goto L8
	case 46:
		goto L7
	}
L3:
	;
	goto L4
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L16
	} else {
		goto L64
	}
L5:
	;
	m.G0 = v7 + int32(32)
	return v196
L6:
	;
	v193 = F_make_notclause(m, l0)
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L16
	} else {
		goto L63
	}
L7:
	;
	v177 = F_palloc0(m, int32(16))
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L16
	} else {
		goto L61
	}
L8:
	;
	v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)))
	if v160 != 0 {
		goto L6
	} else {
		goto L59
	}
L9:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	switch v77 {
	case 0:
		goto L28
	case 1:
		goto L27
	case 2:
		goto L26
	default:
		goto L25
	}
L10:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v53 = F_get_negator(m, v52)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L16
	} else {
		goto L22
	}
L11:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v28 = F_get_negator(m, v27)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L16
	} else {
		goto L19
	}
L12:
	;
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v12 == int32(1) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v17 = F_makeBoolConst(m, int32(0), int32(1))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	goto L15
L15:
	;
	v21 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v25 = F_makeBoolConst(m, base.B2i32(v21 == int64(0)), int32(0))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L16
	} else {
		goto L18
	}
L16:
	;
	return int32(0)
L17:
	;
	v196 = v17
	goto L5
L18:
	;
	v196 = v25
	goto L5
L19:
	;
	if v28 == int32(0) {
		goto L6
	} else {
		goto L20
	}
L20:
	;
	v33 = F_palloc0(m, int32(36))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L16
	} else {
		goto L21
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v33)+4)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v33))) = int32(17)
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v33)+12)) = v40
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)))
	*(*uint8)(unsafe.Add(mBase, uint32(v33)+16)) = uint8(v42)
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v33)+20)) = v44
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v33)+24)) = v46
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v33)+28)) = v48
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v33)+32)) = v50
	v196 = v33
	goto L5
L22:
	;
	if v53 == int32(0) {
		goto L6
	} else {
		goto L23
	}
L23:
	;
	v58 = F_palloc0(m, int32(36))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L16
	} else {
		goto L24
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+16)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v58)+8)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v58)+4)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v58))) = int32(20)
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
	v69 = v67 ^ int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v58)+20)) = uint8(v69)
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v58)+24)) = v71
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v58)+28)) = v73
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v58)+32)) = v75
	v196 = v58
	goto L5
L25:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L16
	} else {
		goto L56
	}
L26:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v143)+12))
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v144)))
	v196 = v145
	goto L5
L27:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v110 == int32(0) {
		goto L43
	} else {
		goto L44
	}
L28:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v78 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v82 = F_make_orclause(m, int32(0))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L16
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	v84 = int32(0)
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v78)+4))
	if v85 <= v84 {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	v196 = v82
	goto L5
L33:
	;
	v89 = F_make_orclause(m, int32(0))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L16
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	v91 = v84
	v92 = v2
	goto L37
L36:
	;
	v196 = v89
	goto L5
L37:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v78)+12))
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v95+v91<<(uint(int32(2))%32))))
	v100 = F_negate_clause(m, v99)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L16
	} else {
		goto L39
	}
L38:
	;
	v108 = F_make_orclause(m, v102)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L16
	} else {
		goto L42
	}
L39:
	;
	v102 = F_lappend(m, v92, v100)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L16
	} else {
		goto L40
	}
L40:
	;
	v105 = v91 + int32(1)
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v78)+4))
	if v105 < v106 {
		v91 = v105
		v92 = v102
		goto L37
	} else {
		goto L41
	}
L41:
	;
	goto L38
L42:
	;
	v196 = v108
	goto L5
L43:
	;
	v114 = F_make_andclause(m, int32(0))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L16
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v110)+4))
	if int32(0) < v116 {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	v196 = v114
	goto L5
L47:
	;
	v120 = int32(0)
	v121 = v2
	goto L50
L48:
	;
	v138 = v2
	goto L49
L49:
	;
	v141 = F_make_andclause(m, v138)
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L16
	} else {
		goto L55
	}
L50:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v110)+12))
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v124+v120<<(uint(int32(2))%32))))
	v129 = F_negate_clause(m, v128)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L16
	} else {
		goto L52
	}
L51:
	;
	v138 = v131
	goto L49
L52:
	;
	v131 = F_lappend(m, v121, v129)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L16
	} else {
		goto L53
	}
L53:
	;
	v134 = v120 + int32(1)
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v110)+4))
	if v134 < v135 {
		v120 = v134
		v121 = v131
		goto L50
	} else {
		goto L54
	}
L54:
	;
	goto L51
L55:
	;
	v196 = v141
	goto L5
L56:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v150
	F_errmsg_internal(m, int32(_a_F_negate_clause_0), v7)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L16
	} else {
		goto L57
	}
L57:
	;
	F_errfinish(m, int32(_a_F_negate_clause_1), int32(195), int32(_a_F_negate_clause_2))
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L16
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
	v162 = F_palloc0(m, int32(20))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L16
	} else {
		goto L60
	}
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v162))) = int32(52)
	v166 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v162)+4)) = v166
	v168 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v162)+8)) = base.B2i32(v168 == int32(0))
	v172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)))
	*(*uint8)(unsafe.Add(mBase, uint32(v162)+12)) = uint8(v172)
	v174 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v162)+16)) = v174
	v196 = v162
	goto L5
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v177))) = int32(53)
	v181 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v177)+4)) = v181
	v183 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if base.Ui32(int32(6)) <= base.Ui32(v183) {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v183<<(uint(int32(2))%32))+uint32(_c_F_negate_clause[0])))
	*(*int32)(unsafe.Add(mBase, uint32(v177)+8)) = v188
	v190 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v177)+12)) = v190
	v196 = v177
	goto L5
L63:
	;
	v196 = v193
	goto L5
L64:
	;
	F_errmsg_internal(m, int32(_a_F_negate_clause_3), int32(0))
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L16
	} else {
		goto L65
	}
L65:
	;
	F_errfinish(m, int32(_a_F_negate_clause_1), int32(76), int32(_a_F_negate_clause_2))
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L16
	} else {
		goto L66
	}
L66:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L67:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v220
	F_errmsg_internal(m, int32(_a_F_negate_clause_4), v7+int32(16))
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L16
	} else {
		goto L68
	}
L68:
	;
	F_errfinish(m, int32(_a_F_negate_clause_1), int32(250), int32(_a_F_negate_clause_2))
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L16
	} else {
		goto L69
	}
L69:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_neqjoinsel(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int64
	_ = v14
	var v15 int64
	_ = v15
	var v16 int64
	_ = v16
	var v17 int64
	_ = v17
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 float32
	_ = v42
	var v46 float64
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
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
	var v67 int64
	_ = v67
	var v68 int32
	_ = v68
	var v73 float64
	_ = v73
	v10 = m.G0
	v12 = v10 - int32(80)
	m.G0 = v12
	v14 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+88)))
	v15 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+56)))
	v16 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+24)))
	v17 = *(*int64)(unsafe.Add(mBase, uint32(l0)+72))
	if v17&int64(65534) == int64(4) {
		F_get_join_variables(m, base.I32_wrap_i64(v16), base.I32_wrap_i64(v15), base.I32_wrap_i64(v14), v12+int32(48), v12+int32(16), v12+int32(15))
		mBase = m.M
		v34 = m.ExcPending
		if v34 != 0 {
			return int64(0)
		} else {
			v35 = *(*int32)(unsafe.Add(mBase, uint32(v12)+24))
			v36 = *(*int32)(unsafe.Add(mBase, uint32(v12)+56))
			v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+15)))
			if v37 != 0 {
				v38 = v35
			} else {
				v38 = v36
			}
			if v38 != 0 {
				v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+16))
				v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+22)))
				v42 = *(*float32)(unsafe.Add(mBase, uint32(v39+v40)+8))
				v46 = base.F64_promote_f32(v42)
			} else {
				v46 = float64(0)
			}
			if v36 != 0 {
				v47 = *(*int32)(unsafe.Add(mBase, uint32(v12)+60))
				m.T0[v47].(func(*base.Module, int32))(m, v36)
				mBase = m.M
				v49 = m.ExcPending
				if v49 != 0 {
					return int64(0)
				} else {
					v50 = *(*int32)(unsafe.Add(mBase, uint32(v12)+24))
					v51 = v50
					if v51 == int32(0) {
						v73 = v46
						m.G0 = v12 + int32(80)
						return base.I64_reinterpret_f64(base.F64_sub(float64(1), v73))
					} else {
						v54 = *(*int32)(unsafe.Add(mBase, uint32(v12)+28))
						m.T0[v54].(func(*base.Module, int32))(m, v51)
						mBase = m.M
						v56 = m.ExcPending
						if v56 != 0 {
							return int64(0)
						} else {
							v73 = v46
							m.G0 = v12 + int32(80)
							return base.I64_reinterpret_f64(base.F64_sub(float64(1), v73))
						}
					}
				}
			} else {
				v51 = v35
				if v51 == int32(0) {
					v73 = v46
					m.G0 = v12 + int32(80)
					return base.I64_reinterpret_f64(base.F64_sub(float64(1), v73))
				} else {
					v54 = *(*int32)(unsafe.Add(mBase, uint32(v12)+28))
					m.T0[v54].(func(*base.Module, int32))(m, v51)
					mBase = m.M
					v56 = m.ExcPending
					if v56 != 0 {
						return int64(0)
					} else {
						v73 = v46
						m.G0 = v12 + int32(80)
						return base.I64_reinterpret_f64(base.F64_sub(float64(1), v73))
					}
				}
			}
		}
	} else {
		v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v59 = F_get_negator(m, v58)
		mBase = m.M
		v60 = m.ExcPending
		if v60 != 0 {
			return int64(0)
		} else {
			if v59 == int32(0) {
				v73 = float64(0.005)
				m.G0 = v12 + int32(80)
				return base.I64_reinterpret_f64(base.F64_sub(float64(1), v73))
			} else {
				v67 = F_DirectFunctionCall5Coll(m, int32(1707), v57, v16, base.I64_extend_i32_u(v59), v15, base.I64_extend16_s(v17), v14)
				mBase = m.M
				v68 = m.ExcPending
				if v68 != 0 {
					return int64(0)
				} else {
					v73 = base.F64_reinterpret_i64(v67)
					m.G0 = v12 + int32(80)
					return base.I64_reinterpret_f64(base.F64_sub(float64(1), v73))
				}
			}
		}
	}
}
func F_neqsel(m *base.Module, l0 int32) int64 {
	var v3 float64
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_eqsel_internal(m, l0, int32(1))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return base.I64_reinterpret_f64(v3)
	}
}
func F_networksel(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v7 int64
	_ = v7
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
	var v20 int32
	_ = v20
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v59 int64
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v72 int64
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v89 int64
	_ = v89
	var v90 int64
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 float32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v108 float64
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
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 float64
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v132 float64
	_ = v132
	var v133 float64
	_ = v133
	var v134 float64
	_ = v134
	var v137 float64
	_ = v137
	var v140 float64
	_ = v140
	var v148 float64
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v158 int64
	_ = v158
	v7 = int64(0)
	v12 = m.G0
	v14 = v12 - int32(128)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	switch v20 - int32(931) {
	case 0:
		v44 = int32(2)
		v51 = F_get_restriction_variable(m, v18, v17, v16, v14+int32(96), v14+int32(92), v14+int32(91))
		mBase = m.M
		v52 = m.ExcPending
		if v52 != 0 {
			return int64(0)
		} else {
			if v51 == int32(0) {
				if v20 == int32(3552) {
					v59 = int64(4576918229304087675)
				} else {
					v59 = int64(4572414629676717179)
				}
				v158 = v59
				m.G0 = v14 + int32(128)
				return v158
			} else {
				v60 = *(*int32)(unsafe.Add(mBase, uint32(v14)+92))
				v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
				if v61 != int32(7) {
					v64 = *(*int32)(unsafe.Add(mBase, uint32(v14)+104))
					if v64 != 0 {
						v65 = *(*int32)(unsafe.Add(mBase, uint32(v14)+108))
						m.T0[v65].(func(*base.Module, int32))(m, v64)
						mBase = m.M
						v67 = m.ExcPending
						if v67 != 0 {
							return int64(0)
						} else {
							if v20 == int32(3552) {
								v72 = int64(4576918229304087675)
							} else {
								v72 = int64(4572414629676717179)
							}
							v158 = v72
							m.G0 = v14 + int32(128)
							return v158
						}
					} else {
						if v20 == int32(3552) {
							v72 = int64(4576918229304087675)
						} else {
							v72 = int64(4572414629676717179)
						}
						v158 = v72
						m.G0 = v14 + int32(128)
						return v158
					}
				} else {
					v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+32)))
					if v73 == int32(1) {
						v76 = *(*int32)(unsafe.Add(mBase, uint32(v14)+104))
						if v76 == int32(0) {
							v158 = v7
							m.G0 = v14 + int32(128)
							return v158
						} else {
							v79 = *(*int32)(unsafe.Add(mBase, uint32(v14)+108))
							m.T0[v79].(func(*base.Module, int32))(m, v76)
							mBase = m.M
							v81 = m.ExcPending
							if v81 != 0 {
								return int64(0)
							} else {
								v158 = v7
								m.G0 = v14 + int32(128)
								return v158
							}
						}
					} else {
						v82 = *(*int32)(unsafe.Add(mBase, uint32(v14)+104))
						if v82 == int32(0) {
							if v20 == int32(3552) {
								v89 = int64(4576918229304087675)
							} else {
								v89 = int64(4572414629676717179)
							}
							v158 = v89
							m.G0 = v14 + int32(128)
							return v158
						} else {
							v90 = *(*int64)(unsafe.Add(mBase, uint32(v60)+24))
							v91 = *(*int32)(unsafe.Add(mBase, uint32(v82)+16))
							v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+22)))
							v94 = *(*float32)(unsafe.Add(mBase, uint32(v91+v92)+8))
							v95 = F_get_opcode(m, v20)
							mBase = m.M
							v96 = m.ExcPending
							if v96 != 0 {
								return int64(0)
							} else {
								v98 = v14 + int32(12)
								F_fmgr_info(m, v95, v98)
								mBase = m.M
								v100 = m.ExcPending
								if v100 != 0 {
									return int64(0)
								} else {
									v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+91)))
									v108 = F_mcv_selectivity(m, v14+int32(96), v98, int32(0), v90, v105, v14+int32(40))
									mBase = m.M
									v109 = m.ExcPending
									if v109 != 0 {
										return int64(0)
									} else {
										v111 = v14 + int32(52)
										v112 = *(*int32)(unsafe.Add(mBase, uint32(v14)+104))
										v116 = F_get_attstatsslot(m, v111, v112, int32(2), int32(0), int32(1))
										mBase = m.M
										v117 = m.ExcPending
										if v117 != 0 {
											return int64(0)
										} else {
											if v116 != 0 {
												v118 = *(*int32)(unsafe.Add(mBase, uint32(v14)+64))
												v119 = *(*int32)(unsafe.Add(mBase, uint32(v14)+68))
												v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+91)))
												if v122 != 0 {
													v123 = v44
												} else {
													v123 = int32(0) - v44
												}
												v124 = F_inet_hist_value_sel(m, v118, v119, v90, v123)
												mBase = m.M
												v125 = m.ExcPending
												if v125 != 0 {
													return int64(0)
												} else {
													F_free_attstatsslot(m, v111)
													mBase = m.M
													v127 = m.ExcPending
													if v127 != 0 {
														return int64(0)
													} else {
														v133 = v124
														v134 = float64(0)
														v137 = *(*float64)(unsafe.Add(mBase, uint32(v14)+40))
														v140 = base.F64_add(base.F64_mul(base.F64_sub(base.F64_sub(float64(1), base.F64_promote_f32(v94)), v137), v133), v108)
														if base.F64_lt(v140, v134) != 0 {
															v148 = v134
														} else {
															if base.F64_gt(v140, float64(1)) == int32(0) {
																v148 = v140
															} else {
																v148 = float64(1)
															}
														}
														v149 = *(*int32)(unsafe.Add(mBase, uint32(v14)+104))
														if v149 != 0 {
															v150 = *(*int32)(unsafe.Add(mBase, uint32(v14)+108))
															m.T0[v150].(func(*base.Module, int32))(m, v149)
															mBase = m.M
															v152 = m.ExcPending
															if v152 != 0 {
																return int64(0)
															} else {
																v158 = base.I64_reinterpret_f64(v148)
																m.G0 = v14 + int32(128)
																return v158
															}
														} else {
															v158 = base.I64_reinterpret_f64(v148)
															m.G0 = v14 + int32(128)
															return v158
														}
													}
												}
											} else {
												if v20 == int32(3552) {
													v132 = float64(0.01)
												} else {
													v132 = float64(0.005)
												}
												v133 = v132
												v134 = float64(0)
												v137 = *(*float64)(unsafe.Add(mBase, uint32(v14)+40))
												v140 = base.F64_add(base.F64_mul(base.F64_sub(base.F64_sub(float64(1), base.F64_promote_f32(v94)), v137), v133), v108)
												if base.F64_lt(v140, v134) != 0 {
													v148 = v134
												} else {
													if base.F64_gt(v140, float64(1)) == int32(0) {
														v148 = v140
													} else {
														v148 = float64(1)
													}
												}
												v149 = *(*int32)(unsafe.Add(mBase, uint32(v14)+104))
												if v149 != 0 {
													v150 = *(*int32)(unsafe.Add(mBase, uint32(v14)+108))
													m.T0[v150].(func(*base.Module, int32))(m, v149)
													mBase = m.M
													v152 = m.ExcPending
													if v152 != 0 {
														return int64(0)
													} else {
														v158 = base.I64_reinterpret_f64(v148)
														m.G0 = v14 + int32(128)
														return v158
													}
												} else {
													v158 = base.I64_reinterpret_f64(v148)
													m.G0 = v14 + int32(128)
													return v158
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
	case 1:
		v44 = int32(1)
		v51 = F_get_restriction_variable(m, v18, v17, v16, v14+int32(96), v14+int32(92), v14+int32(91))
		mBase = m.M
		v52 = m.ExcPending
		if v52 != 0 {
			return int64(0)
		} else {
			if v51 == int32(0) {
				if v20 == int32(3552) {
					v59 = int64(4576918229304087675)
				} else {
					v59 = int64(4572414629676717179)
				}
				v158 = v59
				m.G0 = v14 + int32(128)
				return v158
			} else {
				v60 = *(*int32)(unsafe.Add(mBase, uint32(v14)+92))
				v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
				if v61 != int32(7) {
					v64 = *(*int32)(unsafe.Add(mBase, uint32(v14)+104))
					if v64 != 0 {
						v65 = *(*int32)(unsafe.Add(mBase, uint32(v14)+108))
						m.T0[v65].(func(*base.Module, int32))(m, v64)
						mBase = m.M
						v67 = m.ExcPending
						if v67 != 0 {
							return int64(0)
						} else {
							if v20 == int32(3552) {
								v72 = int64(4576918229304087675)
							} else {
								v72 = int64(4572414629676717179)
							}
							v158 = v72
							m.G0 = v14 + int32(128)
							return v158
						}
					} else {
						if v20 == int32(3552) {
							v72 = int64(4576918229304087675)
						} else {
							v72 = int64(4572414629676717179)
						}
						v158 = v72
						m.G0 = v14 + int32(128)
						return v158
					}
				} else {
					v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+32)))
					if v73 == int32(1) {
						v76 = *(*int32)(unsafe.Add(mBase, uint32(v14)+104))
						if v76 == int32(0) {
							v158 = v7
							m.G0 = v14 + int32(128)
							return v158
						} else {
							v79 = *(*int32)(unsafe.Add(mBase, uint32(v14)+108))
							m.T0[v79].(func(*base.Module, int32))(m, v76)
							mBase = m.M
							v81 = m.ExcPending
							if v81 != 0 {
								return int64(0)
							} else {
								v158 = v7
								m.G0 = v14 + int32(128)
								return v158
							}
						}
					} else {
						v82 = *(*int32)(unsafe.Add(mBase, uint32(v14)+104))
						if v82 == int32(0) {
							if v20 == int32(3552) {
								v89 = int64(4576918229304087675)
							} else {
								v89 = int64(4572414629676717179)
							}
							v158 = v89
							m.G0 = v14 + int32(128)
							return v158
						} else {
							v90 = *(*int64)(unsafe.Add(mBase, uint32(v60)+24))
							v91 = *(*int32)(unsafe.Add(mBase, uint32(v82)+16))
							v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+22)))
							v94 = *(*float32)(unsafe.Add(mBase, uint32(v91+v92)+8))
							v95 = F_get_opcode(m, v20)
							mBase = m.M
							v96 = m.ExcPending
							if v96 != 0 {
								return int64(0)
							} else {
								v98 = v14 + int32(12)
								F_fmgr_info(m, v95, v98)
								mBase = m.M
								v100 = m.ExcPending
								if v100 != 0 {
									return int64(0)
								} else {
									v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+91)))
									v108 = F_mcv_selectivity(m, v14+int32(96), v98, int32(0), v90, v105, v14+int32(40))
									mBase = m.M
									v109 = m.ExcPending
									if v109 != 0 {
										return int64(0)
									} else {
										v111 = v14 + int32(52)
										v112 = *(*int32)(unsafe.Add(mBase, uint32(v14)+104))
										v116 = F_get_attstatsslot(m, v111, v112, int32(2), int32(0), int32(1))
										mBase = m.M
										v117 = m.ExcPending
										if v117 != 0 {
											return int64(0)
										} else {
											if v116 != 0 {
												v118 = *(*int32)(unsafe.Add(mBase, uint32(v14)+64))
												v119 = *(*int32)(unsafe.Add(mBase, uint32(v14)+68))
												v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+91)))
												if v122 != 0 {
													v123 = v44
												} else {
													v123 = int32(0) - v44
												}
												v124 = F_inet_hist_value_sel(m, v118, v119, v90, v123)
												mBase = m.M
												v125 = m.ExcPending
												if v125 != 0 {
													return int64(0)
												} else {
													F_free_attstatsslot(m, v111)
													mBase = m.M
													v127 = m.ExcPending
													if v127 != 0 {
														return int64(0)
													} else {
														v133 = v124
														v134 = float64(0)
														v137 = *(*float64)(unsafe.Add(mBase, uint32(v14)+40))
														v140 = base.F64_add(base.F64_mul(base.F64_sub(base.F64_sub(float64(1), base.F64_promote_f32(v94)), v137), v133), v108)
														if base.F64_lt(v140, v134) != 0 {
															v148 = v134
														} else {
															if base.F64_gt(v140, float64(1)) == int32(0) {
																v148 = v140
															} else {
																v148 = float64(1)
															}
														}
														v149 = *(*int32)(unsafe.Add(mBase, uint32(v14)+104))
														if v149 != 0 {
															v150 = *(*int32)(unsafe.Add(mBase, uint32(v14)+108))
															m.T0[v150].(func(*base.Module, int32))(m, v149)
															mBase = m.M
															v152 = m.ExcPending
															if v152 != 0 {
																return int64(0)
															} else {
																v158 = base.I64_reinterpret_f64(v148)
																m.G0 = v14 + int32(128)
																return v158
															}
														} else {
															v158 = base.I64_reinterpret_f64(v148)
															m.G0 = v14 + int32(128)
															return v158
														}
													}
												}
											} else {
												if v20 == int32(3552) {
													v132 = float64(0.01)
												} else {
													v132 = float64(0.005)
												}
												v133 = v132
												v134 = float64(0)
												v137 = *(*float64)(unsafe.Add(mBase, uint32(v14)+40))
												v140 = base.F64_add(base.F64_mul(base.F64_sub(base.F64_sub(float64(1), base.F64_promote_f32(v94)), v137), v133), v108)
												if base.F64_lt(v140, v134) != 0 {
													v148 = v134
												} else {
													if base.F64_gt(v140, float64(1)) == int32(0) {
														v148 = v140
													} else {
														v148 = float64(1)
													}
												}
												v149 = *(*int32)(unsafe.Add(mBase, uint32(v14)+104))
												if v149 != 0 {
													v150 = *(*int32)(unsafe.Add(mBase, uint32(v14)+108))
													m.T0[v150].(func(*base.Module, int32))(m, v149)
													mBase = m.M
													v152 = m.ExcPending
													if v152 != 0 {
														return int64(0)
													} else {
														v158 = base.I64_reinterpret_f64(v148)
														m.G0 = v14 + int32(128)
														return v158
													}
												} else {
													v158 = base.I64_reinterpret_f64(v148)
													m.G0 = v14 + int32(128)
													return v158
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
	case 2:
		v44 = int32(-2)
		v51 = F_get_restriction_variable(m, v18, v17, v16, v14+int32(96), v14+int32(92), v14+int32(91))
		mBase = m.M
		v52 = m.ExcPending
		if v52 != 0 {
			return int64(0)
		} else {
			if v51 == int32(0) {
				if v20 == int32(3552) {
					v59 = int64(4576918229304087675)
				} else {
					v59 = int64(4572414629676717179)
				}
				v158 = v59
				m.G0 = v14 + int32(128)
				return v158
			} else {
				v60 = *(*int32)(unsafe.Add(mBase, uint32(v14)+92))
				v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
				if v61 != int32(7) {
					v64 = *(*int32)(unsafe.Add(mBase, uint32(v14)+104))
					if v64 != 0 {
						v65 = *(*int32)(unsafe.Add(mBase, uint32(v14)+108))
						m.T0[v65].(func(*base.Module, int32))(m, v64)
						mBase = m.M
						v67 = m.ExcPending
						if v67 != 0 {
							return int64(0)
						} else {
							if v20 == int32(3552) {
								v72 = int64(4576918229304087675)
							} else {
								v72 = int64(4572414629676717179)
							}
							v158 = v72
							m.G0 = v14 + int32(128)
							return v158
						}
					} else {
						if v20 == int32(3552) {
							v72 = int64(4576918229304087675)
						} else {
							v72 = int64(4572414629676717179)
						}
						v158 = v72
						m.G0 = v14 + int32(128)
						return v158
					}
				} else {
					v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+32)))
					if v73 == int32(1) {
						v76 = *(*int32)(unsafe.Add(mBase, uint32(v14)+104))
						if v76 == int32(0) {
							v158 = v7
							m.G0 = v14 + int32(128)
							return v158
						} else {
							v79 = *(*int32)(unsafe.Add(mBase, uint32(v14)+108))
							m.T0[v79].(func(*base.Module, int32))(m, v76)
							mBase = m.M
							v81 = m.ExcPending
							if v81 != 0 {
								return int64(0)
							} else {
								v158 = v7
								m.G0 = v14 + int32(128)
								return v158
							}
						}
					} else {
						v82 = *(*int32)(unsafe.Add(mBase, uint32(v14)+104))
						if v82 == int32(0) {
							if v20 == int32(3552) {
								v89 = int64(4576918229304087675)
							} else {
								v89 = int64(4572414629676717179)
							}
							v158 = v89
							m.G0 = v14 + int32(128)
							return v158
						} else {
							v90 = *(*int64)(unsafe.Add(mBase, uint32(v60)+24))
							v91 = *(*int32)(unsafe.Add(mBase, uint32(v82)+16))
							v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+22)))
							v94 = *(*float32)(unsafe.Add(mBase, uint32(v91+v92)+8))
							v95 = F_get_opcode(m, v20)
							mBase = m.M
							v96 = m.ExcPending
							if v96 != 0 {
								return int64(0)
							} else {
								v98 = v14 + int32(12)
								F_fmgr_info(m, v95, v98)
								mBase = m.M
								v100 = m.ExcPending
								if v100 != 0 {
									return int64(0)
								} else {
									v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+91)))
									v108 = F_mcv_selectivity(m, v14+int32(96), v98, int32(0), v90, v105, v14+int32(40))
									mBase = m.M
									v109 = m.ExcPending
									if v109 != 0 {
										return int64(0)
									} else {
										v111 = v14 + int32(52)
										v112 = *(*int32)(unsafe.Add(mBase, uint32(v14)+104))
										v116 = F_get_attstatsslot(m, v111, v112, int32(2), int32(0), int32(1))
										mBase = m.M
										v117 = m.ExcPending
										if v117 != 0 {
											return int64(0)
										} else {
											if v116 != 0 {
												v118 = *(*int32)(unsafe.Add(mBase, uint32(v14)+64))
												v119 = *(*int32)(unsafe.Add(mBase, uint32(v14)+68))
												v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+91)))
												if v122 != 0 {
													v123 = v44
												} else {
													v123 = int32(0) - v44
												}
												v124 = F_inet_hist_value_sel(m, v118, v119, v90, v123)
												mBase = m.M
												v125 = m.ExcPending
												if v125 != 0 {
													return int64(0)
												} else {
													F_free_attstatsslot(m, v111)
													mBase = m.M
													v127 = m.ExcPending
													if v127 != 0 {
														return int64(0)
													} else {
														v133 = v124
														v134 = float64(0)
														v137 = *(*float64)(unsafe.Add(mBase, uint32(v14)+40))
														v140 = base.F64_add(base.F64_mul(base.F64_sub(base.F64_sub(float64(1), base.F64_promote_f32(v94)), v137), v133), v108)
														if base.F64_lt(v140, v134) != 0 {
															v148 = v134
														} else {
															if base.F64_gt(v140, float64(1)) == int32(0) {
																v148 = v140
															} else {
																v148 = float64(1)
															}
														}
														v149 = *(*int32)(unsafe.Add(mBase, uint32(v14)+104))
														if v149 != 0 {
															v150 = *(*int32)(unsafe.Add(mBase, uint32(v14)+108))
															m.T0[v150].(func(*base.Module, int32))(m, v149)
															mBase = m.M
															v152 = m.ExcPending
															if v152 != 0 {
																return int64(0)
															} else {
																v158 = base.I64_reinterpret_f64(v148)
																m.G0 = v14 + int32(128)
																return v158
															}
														} else {
															v158 = base.I64_reinterpret_f64(v148)
															m.G0 = v14 + int32(128)
															return v158
														}
													}
												}
											} else {
												if v20 == int32(3552) {
													v132 = float64(0.01)
												} else {
													v132 = float64(0.005)
												}
												v133 = v132
												v134 = float64(0)
												v137 = *(*float64)(unsafe.Add(mBase, uint32(v14)+40))
												v140 = base.F64_add(base.F64_mul(base.F64_sub(base.F64_sub(float64(1), base.F64_promote_f32(v94)), v137), v133), v108)
												if base.F64_lt(v140, v134) != 0 {
													v148 = v134
												} else {
													if base.F64_gt(v140, float64(1)) == int32(0) {
														v148 = v140
													} else {
														v148 = float64(1)
													}
												}
												v149 = *(*int32)(unsafe.Add(mBase, uint32(v14)+104))
												if v149 != 0 {
													v150 = *(*int32)(unsafe.Add(mBase, uint32(v14)+108))
													m.T0[v150].(func(*base.Module, int32))(m, v149)
													mBase = m.M
													v152 = m.ExcPending
													if v152 != 0 {
														return int64(0)
													} else {
														v158 = base.I64_reinterpret_f64(v148)
														m.G0 = v14 + int32(128)
														return v158
													}
												} else {
													v158 = base.I64_reinterpret_f64(v148)
													m.G0 = v14 + int32(128)
													return v158
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
	case 3:
		v44 = int32(-1)
		v51 = F_get_restriction_variable(m, v18, v17, v16, v14+int32(96), v14+int32(92), v14+int32(91))
		mBase = m.M
		v52 = m.ExcPending
		if v52 != 0 {
			return int64(0)
		} else {
			if v51 == int32(0) {
				if v20 == int32(3552) {
					v59 = int64(4576918229304087675)
				} else {
					v59 = int64(4572414629676717179)
				}
				v158 = v59
				m.G0 = v14 + int32(128)
				return v158
			} else {
				v60 = *(*int32)(unsafe.Add(mBase, uint32(v14)+92))
				v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
				if v61 != int32(7) {
					v64 = *(*int32)(unsafe.Add(mBase, uint32(v14)+104))
					if v64 != 0 {
						v65 = *(*int32)(unsafe.Add(mBase, uint32(v14)+108))
						m.T0[v65].(func(*base.Module, int32))(m, v64)
						mBase = m.M
						v67 = m.ExcPending
						if v67 != 0 {
							return int64(0)
						} else {
							if v20 == int32(3552) {
								v72 = int64(4576918229304087675)
							} else {
								v72 = int64(4572414629676717179)
							}
							v158 = v72
							m.G0 = v14 + int32(128)
							return v158
						}
					} else {
						if v20 == int32(3552) {
							v72 = int64(4576918229304087675)
						} else {
							v72 = int64(4572414629676717179)
						}
						v158 = v72
						m.G0 = v14 + int32(128)
						return v158
					}
				} else {
					v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+32)))
					if v73 == int32(1) {
						v76 = *(*int32)(unsafe.Add(mBase, uint32(v14)+104))
						if v76 == int32(0) {
							v158 = v7
							m.G0 = v14 + int32(128)
							return v158
						} else {
							v79 = *(*int32)(unsafe.Add(mBase, uint32(v14)+108))
							m.T0[v79].(func(*base.Module, int32))(m, v76)
							mBase = m.M
							v81 = m.ExcPending
							if v81 != 0 {
								return int64(0)
							} else {
								v158 = v7
								m.G0 = v14 + int32(128)
								return v158
							}
						}
					} else {
						v82 = *(*int32)(unsafe.Add(mBase, uint32(v14)+104))
						if v82 == int32(0) {
							if v20 == int32(3552) {
								v89 = int64(4576918229304087675)
							} else {
								v89 = int64(4572414629676717179)
							}
							v158 = v89
							m.G0 = v14 + int32(128)
							return v158
						} else {
							v90 = *(*int64)(unsafe.Add(mBase, uint32(v60)+24))
							v91 = *(*int32)(unsafe.Add(mBase, uint32(v82)+16))
							v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+22)))
							v94 = *(*float32)(unsafe.Add(mBase, uint32(v91+v92)+8))
							v95 = F_get_opcode(m, v20)
							mBase = m.M
							v96 = m.ExcPending
							if v96 != 0 {
								return int64(0)
							} else {
								v98 = v14 + int32(12)
								F_fmgr_info(m, v95, v98)
								mBase = m.M
								v100 = m.ExcPending
								if v100 != 0 {
									return int64(0)
								} else {
									v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+91)))
									v108 = F_mcv_selectivity(m, v14+int32(96), v98, int32(0), v90, v105, v14+int32(40))
									mBase = m.M
									v109 = m.ExcPending
									if v109 != 0 {
										return int64(0)
									} else {
										v111 = v14 + int32(52)
										v112 = *(*int32)(unsafe.Add(mBase, uint32(v14)+104))
										v116 = F_get_attstatsslot(m, v111, v112, int32(2), int32(0), int32(1))
										mBase = m.M
										v117 = m.ExcPending
										if v117 != 0 {
											return int64(0)
										} else {
											if v116 != 0 {
												v118 = *(*int32)(unsafe.Add(mBase, uint32(v14)+64))
												v119 = *(*int32)(unsafe.Add(mBase, uint32(v14)+68))
												v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+91)))
												if v122 != 0 {
													v123 = v44
												} else {
													v123 = int32(0) - v44
												}
												v124 = F_inet_hist_value_sel(m, v118, v119, v90, v123)
												mBase = m.M
												v125 = m.ExcPending
												if v125 != 0 {
													return int64(0)
												} else {
													F_free_attstatsslot(m, v111)
													mBase = m.M
													v127 = m.ExcPending
													if v127 != 0 {
														return int64(0)
													} else {
														v133 = v124
														v134 = float64(0)
														v137 = *(*float64)(unsafe.Add(mBase, uint32(v14)+40))
														v140 = base.F64_add(base.F64_mul(base.F64_sub(base.F64_sub(float64(1), base.F64_promote_f32(v94)), v137), v133), v108)
														if base.F64_lt(v140, v134) != 0 {
															v148 = v134
														} else {
															if base.F64_gt(v140, float64(1)) == int32(0) {
																v148 = v140
															} else {
																v148 = float64(1)
															}
														}
														v149 = *(*int32)(unsafe.Add(mBase, uint32(v14)+104))
														if v149 != 0 {
															v150 = *(*int32)(unsafe.Add(mBase, uint32(v14)+108))
															m.T0[v150].(func(*base.Module, int32))(m, v149)
															mBase = m.M
															v152 = m.ExcPending
															if v152 != 0 {
																return int64(0)
															} else {
																v158 = base.I64_reinterpret_f64(v148)
																m.G0 = v14 + int32(128)
																return v158
															}
														} else {
															v158 = base.I64_reinterpret_f64(v148)
															m.G0 = v14 + int32(128)
															return v158
														}
													}
												}
											} else {
												if v20 == int32(3552) {
													v132 = float64(0.01)
												} else {
													v132 = float64(0.005)
												}
												v133 = v132
												v134 = float64(0)
												v137 = *(*float64)(unsafe.Add(mBase, uint32(v14)+40))
												v140 = base.F64_add(base.F64_mul(base.F64_sub(base.F64_sub(float64(1), base.F64_promote_f32(v94)), v137), v133), v108)
												if base.F64_lt(v140, v134) != 0 {
													v148 = v134
												} else {
													if base.F64_gt(v140, float64(1)) == int32(0) {
														v148 = v140
													} else {
														v148 = float64(1)
													}
												}
												v149 = *(*int32)(unsafe.Add(mBase, uint32(v14)+104))
												if v149 != 0 {
													v150 = *(*int32)(unsafe.Add(mBase, uint32(v14)+108))
													m.T0[v150].(func(*base.Module, int32))(m, v149)
													mBase = m.M
													v152 = m.ExcPending
													if v152 != 0 {
														return int64(0)
													} else {
														v158 = base.I64_reinterpret_f64(v148)
														m.G0 = v14 + int32(128)
														return v158
													}
												} else {
													v158 = base.I64_reinterpret_f64(v148)
													m.G0 = v14 + int32(128)
													return v158
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
	default:
		if v20 == int32(3552) {
			v44 = int32(0)
			v51 = F_get_restriction_variable(m, v18, v17, v16, v14+int32(96), v14+int32(92), v14+int32(91))
			mBase = m.M
			v52 = m.ExcPending
			if v52 != 0 {
				return int64(0)
			} else {
				if v51 == int32(0) {
					if v20 == int32(3552) {
						v59 = int64(4576918229304087675)
					} else {
						v59 = int64(4572414629676717179)
					}
					v158 = v59
					m.G0 = v14 + int32(128)
					return v158
				} else {
					v60 = *(*int32)(unsafe.Add(mBase, uint32(v14)+92))
					v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
					if v61 != int32(7) {
						v64 = *(*int32)(unsafe.Add(mBase, uint32(v14)+104))
						if v64 != 0 {
							v65 = *(*int32)(unsafe.Add(mBase, uint32(v14)+108))
							m.T0[v65].(func(*base.Module, int32))(m, v64)
							mBase = m.M
							v67 = m.ExcPending
							if v67 != 0 {
								return int64(0)
							} else {
								if v20 == int32(3552) {
									v72 = int64(4576918229304087675)
								} else {
									v72 = int64(4572414629676717179)
								}
								v158 = v72
								m.G0 = v14 + int32(128)
								return v158
							}
						} else {
							if v20 == int32(3552) {
								v72 = int64(4576918229304087675)
							} else {
								v72 = int64(4572414629676717179)
							}
							v158 = v72
							m.G0 = v14 + int32(128)
							return v158
						}
					} else {
						v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+32)))
						if v73 == int32(1) {
							v76 = *(*int32)(unsafe.Add(mBase, uint32(v14)+104))
							if v76 == int32(0) {
								v158 = v7
								m.G0 = v14 + int32(128)
								return v158
							} else {
								v79 = *(*int32)(unsafe.Add(mBase, uint32(v14)+108))
								m.T0[v79].(func(*base.Module, int32))(m, v76)
								mBase = m.M
								v81 = m.ExcPending
								if v81 != 0 {
									return int64(0)
								} else {
									v158 = v7
									m.G0 = v14 + int32(128)
									return v158
								}
							}
						} else {
							v82 = *(*int32)(unsafe.Add(mBase, uint32(v14)+104))
							if v82 == int32(0) {
								if v20 == int32(3552) {
									v89 = int64(4576918229304087675)
								} else {
									v89 = int64(4572414629676717179)
								}
								v158 = v89
								m.G0 = v14 + int32(128)
								return v158
							} else {
								v90 = *(*int64)(unsafe.Add(mBase, uint32(v60)+24))
								v91 = *(*int32)(unsafe.Add(mBase, uint32(v82)+16))
								v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+22)))
								v94 = *(*float32)(unsafe.Add(mBase, uint32(v91+v92)+8))
								v95 = F_get_opcode(m, v20)
								mBase = m.M
								v96 = m.ExcPending
								if v96 != 0 {
									return int64(0)
								} else {
									v98 = v14 + int32(12)
									F_fmgr_info(m, v95, v98)
									mBase = m.M
									v100 = m.ExcPending
									if v100 != 0 {
										return int64(0)
									} else {
										v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+91)))
										v108 = F_mcv_selectivity(m, v14+int32(96), v98, int32(0), v90, v105, v14+int32(40))
										mBase = m.M
										v109 = m.ExcPending
										if v109 != 0 {
											return int64(0)
										} else {
											v111 = v14 + int32(52)
											v112 = *(*int32)(unsafe.Add(mBase, uint32(v14)+104))
											v116 = F_get_attstatsslot(m, v111, v112, int32(2), int32(0), int32(1))
											mBase = m.M
											v117 = m.ExcPending
											if v117 != 0 {
												return int64(0)
											} else {
												if v116 != 0 {
													v118 = *(*int32)(unsafe.Add(mBase, uint32(v14)+64))
													v119 = *(*int32)(unsafe.Add(mBase, uint32(v14)+68))
													v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+91)))
													if v122 != 0 {
														v123 = v44
													} else {
														v123 = int32(0) - v44
													}
													v124 = F_inet_hist_value_sel(m, v118, v119, v90, v123)
													mBase = m.M
													v125 = m.ExcPending
													if v125 != 0 {
														return int64(0)
													} else {
														F_free_attstatsslot(m, v111)
														mBase = m.M
														v127 = m.ExcPending
														if v127 != 0 {
															return int64(0)
														} else {
															v133 = v124
															v134 = float64(0)
															v137 = *(*float64)(unsafe.Add(mBase, uint32(v14)+40))
															v140 = base.F64_add(base.F64_mul(base.F64_sub(base.F64_sub(float64(1), base.F64_promote_f32(v94)), v137), v133), v108)
															if base.F64_lt(v140, v134) != 0 {
																v148 = v134
															} else {
																if base.F64_gt(v140, float64(1)) == int32(0) {
																	v148 = v140
																} else {
																	v148 = float64(1)
																}
															}
															v149 = *(*int32)(unsafe.Add(mBase, uint32(v14)+104))
															if v149 != 0 {
																v150 = *(*int32)(unsafe.Add(mBase, uint32(v14)+108))
																m.T0[v150].(func(*base.Module, int32))(m, v149)
																mBase = m.M
																v152 = m.ExcPending
																if v152 != 0 {
																	return int64(0)
																} else {
																	v158 = base.I64_reinterpret_f64(v148)
																	m.G0 = v14 + int32(128)
																	return v158
																}
															} else {
																v158 = base.I64_reinterpret_f64(v148)
																m.G0 = v14 + int32(128)
																return v158
															}
														}
													}
												} else {
													if v20 == int32(3552) {
														v132 = float64(0.01)
													} else {
														v132 = float64(0.005)
													}
													v133 = v132
													v134 = float64(0)
													v137 = *(*float64)(unsafe.Add(mBase, uint32(v14)+40))
													v140 = base.F64_add(base.F64_mul(base.F64_sub(base.F64_sub(float64(1), base.F64_promote_f32(v94)), v137), v133), v108)
													if base.F64_lt(v140, v134) != 0 {
														v148 = v134
													} else {
														if base.F64_gt(v140, float64(1)) == int32(0) {
															v148 = v140
														} else {
															v148 = float64(1)
														}
													}
													v149 = *(*int32)(unsafe.Add(mBase, uint32(v14)+104))
													if v149 != 0 {
														v150 = *(*int32)(unsafe.Add(mBase, uint32(v14)+108))
														m.T0[v150].(func(*base.Module, int32))(m, v149)
														mBase = m.M
														v152 = m.ExcPending
														if v152 != 0 {
															return int64(0)
														} else {
															v158 = base.I64_reinterpret_f64(v148)
															m.G0 = v14 + int32(128)
															return v158
														}
													} else {
														v158 = base.I64_reinterpret_f64(v148)
														m.G0 = v14 + int32(128)
														return v158
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
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v32 = m.ExcPending
			if v32 != 0 {
				return int64(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v14))) = v20
				F_errmsg_internal(m, int32(_a_F_networksel_0), v14)
				mBase = m.M
				v36 = m.ExcPending
				if v36 != 0 {
					return int64(0)
				} else {
					F_errfinish(m, int32(_a_F_networksel_1), int32(870), int32(_a_F_networksel_2))
					mBase = m.M
					v41 = m.ExcPending
					if v41 != 0 {
						return int64(0)
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
func F_newarc(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v47 int32
	_ = v47
	v6 = *(*int32)(unsafe.Add(mBase, _c_F_newarc[0]))
	if v6 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	if v9 <= v10 {
		goto L8
	} else {
		goto L9
	}
L4:
	;
	return
L5:
	;
	goto L3
L6:
	;
	return
L7:
	;
	F_createarc(m, l0, int32(110), int32(0), l1, l2)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L4
	} else {
		goto L27
	}
L8:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v12 == int32(0) {
		goto L7
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	if v26 == int32(0) {
		goto L7
	} else {
		goto L19
	}
L11:
	;
	v18 = v12
	goto L12
L12:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	if v19 != l2 {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	goto L7
L14:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v18)+16))
	if v25 != 0 {
		v18 = v25
		goto L12
	} else {
		goto L18
	}
L15:
	;
	v21 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v18)+4)))
	if v21 != 0 {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	if v22 == int32(110) {
		goto L6
	} else {
		goto L17
	}
L17:
	;
	goto L14
L18:
	;
	goto L13
L19:
	;
	v32 = v26
	goto L20
L20:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+8))
	if v33 != l1 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	goto L7
L22:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v32)+24))
	if v39 != 0 {
		v32 = v39
		goto L20
	} else {
		goto L26
	}
L23:
	;
	v35 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v32)+4)))
	if v35 != 0 {
		goto L22
	} else {
		goto L24
	}
L24:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	if v36 == int32(110) {
		goto L6
	} else {
		goto L25
	}
L25:
	;
	goto L22
L26:
	;
	goto L21
L27:
	;
	goto L6
}
func F_nocache_index_getattr(m *base.Module, l0 int32, l1 int32, l2 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v219 int32
	_ = v219
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v275 int32
	_ = v275
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v341 int32
	_ = v341
	var v343 int32
	_ = v343
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v353 int32
	_ = v353
	var v358 int32
	_ = v358
	var v368 int32
	_ = v368
	var v370 int32
	_ = v370
	var v383 int32
	_ = v383
	var v385 int32
	_ = v385
	var v397 int32
	_ = v397
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v413 int32
	_ = v413
	var v415 int32
	_ = v415
	var v419 int32
	_ = v419
	var v424 int32
	_ = v424
	var v435 int32
	_ = v435
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v460 int64
	_ = v460
	var v461 int64
	_ = v461
	var v462 int64
	_ = v462
	var v463 int64
	_ = v463
	var v469 int32
	_ = v469
	var v473 int32
	_ = v473
	var v478 int32
	_ = v478
	var v480 int64
	_ = v480
	v4 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(16)
	m.G0 = v15
	v18 = l1 - int32(1)
	v19 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+6)))
	v21 = base.B2i32(v4 <= v19)
	if v21 == v4 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v25 = l0 + int32(8)
	v26 = int32(0)
	v28 = v18 >> (uint(int32(3)) % 32)
	if v28 <= v26 {
		v51 = v26
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v76 = v18
	v81 = v4
	goto L3
L3:
	;
	if v4 <= v19 {
		goto L13
	} else {
		goto L14
	}
L4:
	;
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51+v25))))
	v69 = base.I32_ctz(v63^int32(-1)) + v51<<(uint(int32(3))%32)
	if v69 < v18 {
		goto L10
	} else {
		goto L11
	}
L5:
	;
	v32 = v26
	goto L6
L6:
	;
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32+v25))))
	if v44 != int32(255) {
		v51 = v32
		goto L4
	} else {
		goto L8
	}
L7:
	;
	v51 = v28
	goto L4
L8:
	;
	v48 = v32 + int32(1)
	if v48 != v28 {
		v32 = v48
		goto L6
	} else {
		goto L9
	}
L9:
	;
	goto L7
L10:
	;
	v71 = v69
	goto L12
L11:
	;
	v71 = v18
	goto L12
L12:
	;
	v76 = v71
	v81 = v25
	goto L3
L13:
	;
	v86 = int32(8)
	goto L15
L14:
	;
	v86 = int32(16)
	goto L15
L15:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	if v87 <= int32(0) {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v104 = l0 + v86
	if v19&int32(_a_F_nocache_index_getattr_0) != 0 {
		goto L25
	} else {
		goto L26
	}
L17:
	;
	v102 = int32(0)
	v103 = v4
	goto L16
L18:
	;
	goto L19
L19:
	;
	v91 = int32(0)
	if v76 <= v91 {
		v102 = v91
		v103 = v4
		goto L16
	} else {
		goto L20
	}
L20:
	;
	if base.Ui32(v87) < base.Ui32(v76) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v95 = v87
	goto L23
L22:
	;
	v95 = v76
	goto L23
L23:
	;
	v97 = v95 - int32(1)
	v101 = int32(*(*int16)(unsafe.Add(mBase, uint32(l2+v97<<(uint(int32(3))%32))+28)))
	v102 = v97
	v103 = v101
	goto L16
L24:
	;
	v435 = l2 + v18<<(uint(int32(3))%32)
	v437 = v435 + int32(28)
	v438 = int32(*(*int16)(unsafe.Add(mBase, uint32(v435)+30)))
	if v438 == int32(-1) {
		goto L99
	} else {
		goto L100
	}
L25:
	;
	if v102 < v76 {
		goto L28
	} else {
		goto L29
	}
L26:
	;
	goto L27
L27:
	;
	if v102 < v76 {
		goto L79
	} else {
		goto L80
	}
L28:
	;
	v111 = v102
	v113 = v103
	goto L31
L29:
	;
	v190 = v102
	v192 = v103
	goto L30
L30:
	;
	if v18 <= v190 {
		v424 = v192
		goto L24
	} else {
		goto L53
	}
L31:
	;
	v124 = l2 + int32(28) + v111<<(uint(int32(3))%32)
	v125 = int32(*(*int16)(unsafe.Add(mBase, uint32(v124)+2)))
	if v125 == int32(-1) {
		goto L35
	} else {
		goto L36
	}
L32:
	;
	v190 = v76
	v192 = v185
	goto L30
L33:
	;
	v185 = v182 + v183
	v187 = v111 + int32(1)
	if v187 != v76 {
		v111 = v187
		v113 = v185
		goto L31
	} else {
		goto L52
	}
L34:
	;
	if v145&int32(1) != 0 {
		goto L49
	} else {
		goto L50
	}
L35:
	;
	v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v113+v104))))
	if v129 == int32(0) {
		goto L38
	} else {
		goto L39
	}
L36:
	;
	goto L37
L37:
	;
	v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v124)+5)))
	v166 = int32(0)
	v168 = (v113 + v162 - int32(1)) & (v166 - v162)
	if v166 < v125 {
		v182 = v125
		v183 = v168
		goto L33
	} else {
		goto L48
	}
L38:
	;
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v124)+5)))
	v138 = (v113 + v132 - int32(1)) & (int32(0) - v132)
	v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104+v138))))
	v141 = v140
	v142 = v138
	goto L40
L39:
	;
	v141 = v129
	v142 = v113
	goto L40
L40:
	;
	v143 = v142 + v104
	v145 = v141 & int32(255)
	if v145 != int32(1) {
		goto L34
	} else {
		goto L41
	}
L41:
	;
	v149 = int32(18)
	v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v143)+1)))
	if v151 == v149 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v154 = v149
	goto L44
L43:
	;
	v154 = int32(2)
	goto L44
L44:
	;
	if base.Ui32((v151-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v161 = int32(6)
	goto L47
L46:
	;
	v161 = v154
	goto L47
L47:
	;
	v182 = v161
	v183 = v142
	goto L33
L48:
	;
	v172 = F_strlen(m, v168+v104)
	mBase = m.M
	v182 = v172 + int32(1)
	v183 = v168
	goto L33
L49:
	;
	v182 = int32(base.Ui32(v145) >> (uint(int32(1)) % 32))
	v183 = v142
	goto L33
L50:
	;
	goto L51
L51:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v143)))
	v182 = int32(base.Ui32(v179) >> (uint(int32(2)) % 32))
	v183 = v142
	goto L33
L52:
	;
	goto L32
L53:
	;
	v205 = v190
	v207 = v192
	goto L54
L54:
	;
	v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81+int32(base.Ui32(v205)>>(uint(int32(3))%32))))))
	if int32(base.Ui32(v219)>>(uint(v205&int32(7))%32))&int32(1) != 0 {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	v424 = v290
	goto L24
L56:
	;
	v227 = l2 + int32(28) + v205<<(uint(int32(3))%32)
	v228 = int32(*(*int16)(unsafe.Add(mBase, uint32(v227)+2)))
	if v228 == int32(-1) {
		goto L61
	} else {
		goto L62
	}
L57:
	;
	v290 = v207
	goto L58
L58:
	;
	v293 = v205 + int32(1)
	if v293 != v18 {
		v205 = v293
		v207 = v290
		goto L54
	} else {
		goto L78
	}
L59:
	;
	v290 = v285 + v286
	goto L58
L60:
	;
	if v248&int32(1) != 0 {
		goto L75
	} else {
		goto L76
	}
L61:
	;
	v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207+v104))))
	if v232 == int32(0) {
		goto L64
	} else {
		goto L65
	}
L62:
	;
	goto L63
L63:
	;
	v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v227)+5)))
	v269 = int32(0)
	v271 = (v207 + v265 - int32(1)) & (v269 - v265)
	if v269 < v228 {
		v285 = v228
		v286 = v271
		goto L59
	} else {
		goto L74
	}
L64:
	;
	v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v227)+5)))
	v241 = (v207 + v235 - int32(1)) & (int32(0) - v235)
	v243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104+v241))))
	v244 = v243
	v245 = v241
	goto L66
L65:
	;
	v244 = v232
	v245 = v207
	goto L66
L66:
	;
	v246 = v245 + v104
	v248 = v244 & int32(255)
	if v248 != int32(1) {
		goto L60
	} else {
		goto L67
	}
L67:
	;
	v252 = int32(18)
	v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v246)+1)))
	if v254 == v252 {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v257 = v252
	goto L70
L69:
	;
	v257 = int32(2)
	goto L70
L70:
	;
	if base.Ui32((v254-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v264 = int32(6)
	goto L73
L72:
	;
	v264 = v257
	goto L73
L73:
	;
	v285 = v264
	v286 = v245
	goto L59
L74:
	;
	v275 = F_strlen(m, v271+v104)
	mBase = m.M
	v285 = v275 + int32(1)
	v286 = v271
	goto L59
L75:
	;
	v285 = int32(base.Ui32(v248) >> (uint(int32(1)) % 32))
	v286 = v245
	goto L59
L76:
	;
	goto L77
L77:
	;
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v246)))
	v285 = int32(base.Ui32(v282) >> (uint(int32(2)) % 32))
	v286 = v245
	goto L59
L78:
	;
	goto L55
L79:
	;
	v296 = int32(1)
	v297 = v102 + v296
	v299 = l2 + int32(28)
	if (v76-v102)&v296 != 0 {
		goto L82
	} else {
		goto L83
	}
L80:
	;
	v368 = v102
	v370 = v103
	goto L81
L81:
	;
	if v18 <= v368 {
		v424 = v370
		goto L24
	} else {
		goto L91
	}
L82:
	;
	v305 = v299 + v102<<(uint(int32(3))%32)
	v306 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v305)+5)))
	v313 = int32(*(*int16)(unsafe.Add(mBase, uint32(v305)+2)))
	v315 = v297
	v316 = (v103+v306-int32(1))&(int32(0)-v306) + v313
	goto L84
L83:
	;
	v315 = v102
	v316 = v103
	goto L84
L84:
	;
	if v297 != v76 {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v320 = v315
	v322 = v316
	goto L88
L86:
	;
	v358 = v316
	goto L87
L87:
	;
	v368 = v76
	v370 = v358
	goto L81
L88:
	;
	v333 = v299 + v320<<(uint(int32(3))%32)
	v334 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v333)+5)))
	v336 = int32(1)
	v338 = int32(0)
	v341 = int32(*(*int16)(unsafe.Add(mBase, uint32(v333)+2)))
	v343 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v333)+13)))
	v350 = int32(*(*int16)(unsafe.Add(mBase, uint32(v333)+10)))
	v351 = ((v322+v334-v336)&(v338-v334)+v341+v343-v336)&(v338-v343) + v350
	v353 = v320 + int32(2)
	if v353 != v76 {
		v320 = v353
		v322 = v351
		goto L88
	} else {
		goto L90
	}
L89:
	;
	v358 = v351
	goto L87
L90:
	;
	goto L89
L91:
	;
	v383 = v368
	v385 = v370
	goto L92
L92:
	;
	v397 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81+v383>>(uint(int32(3))%32)))))
	if int32(base.Ui32(v397)>>(uint(v383&int32(7))%32))&int32(1) != 0 {
		goto L94
	} else {
		goto L95
	}
L93:
	;
	v424 = v415
	goto L24
L94:
	;
	v405 = l2 + int32(28) + v383<<(uint(int32(3))%32)
	v406 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v405)+5)))
	v413 = int32(*(*int16)(unsafe.Add(mBase, uint32(v405)+2)))
	v415 = (v385+v406-int32(1))&(int32(0)-v406) + v413
	goto L96
L95:
	;
	v415 = v385
	goto L96
L96:
	;
	v419 = v383 + int32(1)
	if v419 != v18 {
		v383 = v419
		v385 = v415
		goto L92
	} else {
		goto L97
	}
L97:
	;
	goto L93
L98:
	;
	v452 = v451 + v104
	v453 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v437)+4)))
	if v453 == int32(1) {
		goto L104
	} else {
		goto L105
	}
L99:
	;
	v442 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v424+v104))))
	if v442 != 0 {
		v451 = v424
		goto L98
	} else {
		goto L102
	}
L100:
	;
	goto L101
L101:
	;
	v443 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v437)+5)))
	v451 = (v424 + v443 - int32(1)) & (int32(0) - v443)
	goto L98
L102:
	;
	goto L101
L103:
	;
	m.G0 = v15 + int32(16)
	return v480
L104:
	;
	if base.I32_popcnt(v438) != int32(1) {
		goto L107
	} else {
		goto L108
	}
L105:
	;
	goto L106
L106:
	;
	v480 = base.I64_extend_i32_u(v452)
	goto L103
L107:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v469 = m.ExcPending
	if v469 != 0 {
		goto L113
	} else {
		goto L114
	}
L108:
	;
	switch base.I32_ctz(v438) {
	case 0:
		goto L112
	case 1:
		goto L111
	case 2:
		goto L110
	case 3:
		goto L109
	default:
		goto L107
	}
L109:
	;
	v463 = *(*int64)(unsafe.Add(mBase, uint32(v452)))
	v480 = v463
	goto L103
L110:
	;
	v462 = int64(*(*int32)(unsafe.Add(mBase, uint32(v452))))
	v480 = v462
	goto L103
L111:
	;
	v461 = int64(*(*int16)(unsafe.Add(mBase, uint32(v452))))
	v480 = v461
	goto L103
L112:
	;
	v460 = int64(*(*int8)(unsafe.Add(mBase, uint32(v452))))
	v480 = v460
	goto L103
L113:
	;
	return int64(0)
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v438
	F_errmsg_internal(m, int32(_a_F_nocache_index_getattr_1), v15)
	mBase = m.M
	v473 = m.ExcPending
	if v473 != 0 {
		goto L113
	} else {
		goto L115
	}
L115:
	;
	F_errfinish(m, int32(_a_F_nocache_index_getattr_2), int32(123), int32(_a_F_nocache_index_getattr_3))
	mBase = m.M
	v478 = m.ExcPending
	if v478 != 0 {
		goto L113
	} else {
		goto L116
	}
L116:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_nodeToString(m *base.Module, l0 int32) int32 {
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
	var v9 int32
	_ = v9
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	v2 = int32(0)
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = int32(_a_F_nodeToString_0)
	v9 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_nodeToString[0])))
	*(*uint8)(unsafe.Add(mBase, _c_F_nodeToString[0])) = uint8(v2)
	F_initStringInfo(m, v6)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		F_outNode(m, v6, l0)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			v21 = v9 & int32(1)
			*(*uint8)(unsafe.Add(mBase, _c_F_nodeToString[0])) = uint8(v21)
			v23 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
			m.G0 = v6 + int32(16)
			return v23
		}
	}
}
func F_norwegian_ISO_8859_1_create_env(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_SN_new_env(m, int32(32))
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		if v3 != 0 {
			*(*int32)(unsafe.Add(mBase, uint32(v3)+28)) = int32(0)
		} else {
		}
		return v3
	}
}
func F_notification_hash(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v16 int32
	_ = v16
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
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
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
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
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v258 int32
	_ = v258
	var v262 int32
	_ = v262
	var v266 int32
	_ = v266
	var v270 int32
	_ = v270
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v5 = v3 + int32(4)
	v6 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3))))
	v7 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3)+2)))
	v10 = v6 + v7 + int32(1)
	v16 = v10 - int32(1636608432)
	if v5&int32(3) != 0 {
		if base.Ui32(int32(11)) < base.Ui32(v10) {
			v125 = v5
			v126 = v10
			v127 = v16
			v128 = v16
			v129 = v16
			for {
				v131 = *(*int32)(unsafe.Add(mBase, uint32(v125)+4))
				v132 = v131 + v128
				v133 = *(*int32)(unsafe.Add(mBase, uint32(v125)))
				v135 = *(*int32)(unsafe.Add(mBase, uint32(v125)+8))
				v136 = v135 + v129
				v138 = int32(4)
				v140 = v133 + v127 - v136 ^ base.I32_rotl(v136, v138)
				v144 = v132 - v140 ^ base.I32_rotl(v140, int32(6))
				v145 = v136 + v132
				v146 = v140 + v145
				v147 = v144 + v146
				v151 = v145 - v144 ^ base.I32_rotl(v144, int32(8))
				v155 = v146 - v151 ^ base.I32_rotl(v151, int32(16))
				v159 = v147 - v155 ^ base.I32_rotl(v155, int32(19))
				v160 = v151 + v147
				v161 = v155 + v160
				v162 = v159 + v161
				v166 = v160 - v159 ^ base.I32_rotl(v159, v138)
				v167 = int32(12)
				v168 = v125 + v167
				v170 = v126 - v167
				if base.Ui32(int32(11)) < base.Ui32(v170) {
					v125 = v168
					v126 = v170
					v127 = v161
					v128 = v162
					v129 = v166
					continue
				} else {
					break
				}
				break
			}
			v173 = v168
			v174 = v170
			v175 = v161
			v176 = v162
			v177 = v166
		} else {
			v173 = v5
			v174 = v10
			v175 = v16
			v176 = v16
			v177 = v16
		}
		switch v174 - int32(1) {
		case 0:
			v236 = v175
			v237 = v176
			v238 = v177
			v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173))))
			v243 = v236 + v239
			v244 = v237
			v245 = v238
		case 1:
			v229 = v175
			v230 = v176
			v231 = v177
			v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+1)))
			v236 = v232<<(uint(int32(8))%32) + v229
			v237 = v230
			v238 = v231
			v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173))))
			v243 = v236 + v239
			v244 = v237
			v245 = v238
		case 2:
			v222 = v175
			v223 = v176
			v224 = v177
			v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+2)))
			v229 = v225<<(uint(int32(16))%32) + v222
			v230 = v223
			v231 = v224
			v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+1)))
			v236 = v232<<(uint(int32(8))%32) + v229
			v237 = v230
			v238 = v231
			v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173))))
			v243 = v236 + v239
			v244 = v237
			v245 = v238
		case 3:
			v216 = v176
			v217 = v177
			v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+3)))
			v222 = v218<<(uint(int32(24))%32) + v175
			v223 = v216
			v224 = v217
			v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+2)))
			v229 = v225<<(uint(int32(16))%32) + v222
			v230 = v223
			v231 = v224
			v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+1)))
			v236 = v232<<(uint(int32(8))%32) + v229
			v237 = v230
			v238 = v231
			v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173))))
			v243 = v236 + v239
			v244 = v237
			v245 = v238
		case 4:
			v212 = v176
			v213 = v177
			v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+4)))
			v216 = v212 + v214
			v217 = v213
			v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+3)))
			v222 = v218<<(uint(int32(24))%32) + v175
			v223 = v216
			v224 = v217
			v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+2)))
			v229 = v225<<(uint(int32(16))%32) + v222
			v230 = v223
			v231 = v224
			v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+1)))
			v236 = v232<<(uint(int32(8))%32) + v229
			v237 = v230
			v238 = v231
			v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173))))
			v243 = v236 + v239
			v244 = v237
			v245 = v238
		case 5:
			v206 = v176
			v207 = v177
			v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+5)))
			v212 = v208<<(uint(int32(8))%32) + v206
			v213 = v207
			v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+4)))
			v216 = v212 + v214
			v217 = v213
			v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+3)))
			v222 = v218<<(uint(int32(24))%32) + v175
			v223 = v216
			v224 = v217
			v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+2)))
			v229 = v225<<(uint(int32(16))%32) + v222
			v230 = v223
			v231 = v224
			v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+1)))
			v236 = v232<<(uint(int32(8))%32) + v229
			v237 = v230
			v238 = v231
			v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173))))
			v243 = v236 + v239
			v244 = v237
			v245 = v238
		case 6:
			v200 = v176
			v201 = v177
			v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+6)))
			v206 = v202<<(uint(int32(16))%32) + v200
			v207 = v201
			v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+5)))
			v212 = v208<<(uint(int32(8))%32) + v206
			v213 = v207
			v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+4)))
			v216 = v212 + v214
			v217 = v213
			v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+3)))
			v222 = v218<<(uint(int32(24))%32) + v175
			v223 = v216
			v224 = v217
			v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+2)))
			v229 = v225<<(uint(int32(16))%32) + v222
			v230 = v223
			v231 = v224
			v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+1)))
			v236 = v232<<(uint(int32(8))%32) + v229
			v237 = v230
			v238 = v231
			v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173))))
			v243 = v236 + v239
			v244 = v237
			v245 = v238
		case 7:
			v195 = v177
			v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+7)))
			v200 = v196<<(uint(int32(24))%32) + v176
			v201 = v195
			v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+6)))
			v206 = v202<<(uint(int32(16))%32) + v200
			v207 = v201
			v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+5)))
			v212 = v208<<(uint(int32(8))%32) + v206
			v213 = v207
			v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+4)))
			v216 = v212 + v214
			v217 = v213
			v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+3)))
			v222 = v218<<(uint(int32(24))%32) + v175
			v223 = v216
			v224 = v217
			v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+2)))
			v229 = v225<<(uint(int32(16))%32) + v222
			v230 = v223
			v231 = v224
			v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+1)))
			v236 = v232<<(uint(int32(8))%32) + v229
			v237 = v230
			v238 = v231
			v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173))))
			v243 = v236 + v239
			v244 = v237
			v245 = v238
		case 8:
			v190 = v177
			v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+8)))
			v195 = v191<<(uint(int32(8))%32) + v190
			v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+7)))
			v200 = v196<<(uint(int32(24))%32) + v176
			v201 = v195
			v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+6)))
			v206 = v202<<(uint(int32(16))%32) + v200
			v207 = v201
			v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+5)))
			v212 = v208<<(uint(int32(8))%32) + v206
			v213 = v207
			v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+4)))
			v216 = v212 + v214
			v217 = v213
			v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+3)))
			v222 = v218<<(uint(int32(24))%32) + v175
			v223 = v216
			v224 = v217
			v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+2)))
			v229 = v225<<(uint(int32(16))%32) + v222
			v230 = v223
			v231 = v224
			v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+1)))
			v236 = v232<<(uint(int32(8))%32) + v229
			v237 = v230
			v238 = v231
			v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173))))
			v243 = v236 + v239
			v244 = v237
			v245 = v238
		case 9:
			v185 = v177
			v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+9)))
			v190 = v186<<(uint(int32(16))%32) + v185
			v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+8)))
			v195 = v191<<(uint(int32(8))%32) + v190
			v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+7)))
			v200 = v196<<(uint(int32(24))%32) + v176
			v201 = v195
			v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+6)))
			v206 = v202<<(uint(int32(16))%32) + v200
			v207 = v201
			v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+5)))
			v212 = v208<<(uint(int32(8))%32) + v206
			v213 = v207
			v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+4)))
			v216 = v212 + v214
			v217 = v213
			v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+3)))
			v222 = v218<<(uint(int32(24))%32) + v175
			v223 = v216
			v224 = v217
			v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+2)))
			v229 = v225<<(uint(int32(16))%32) + v222
			v230 = v223
			v231 = v224
			v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+1)))
			v236 = v232<<(uint(int32(8))%32) + v229
			v237 = v230
			v238 = v231
			v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173))))
			v243 = v236 + v239
			v244 = v237
			v245 = v238
		case 10:
			v181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+10)))
			v185 = v181<<(uint(int32(24))%32) + v177
			v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+9)))
			v190 = v186<<(uint(int32(16))%32) + v185
			v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+8)))
			v195 = v191<<(uint(int32(8))%32) + v190
			v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+7)))
			v200 = v196<<(uint(int32(24))%32) + v176
			v201 = v195
			v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+6)))
			v206 = v202<<(uint(int32(16))%32) + v200
			v207 = v201
			v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+5)))
			v212 = v208<<(uint(int32(8))%32) + v206
			v213 = v207
			v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+4)))
			v216 = v212 + v214
			v217 = v213
			v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+3)))
			v222 = v218<<(uint(int32(24))%32) + v175
			v223 = v216
			v224 = v217
			v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+2)))
			v229 = v225<<(uint(int32(16))%32) + v222
			v230 = v223
			v231 = v224
			v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+1)))
			v236 = v232<<(uint(int32(8))%32) + v229
			v237 = v230
			v238 = v231
			v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173))))
			v243 = v236 + v239
			v244 = v237
			v245 = v238
		default:
			v243 = v175
			v244 = v176
			v245 = v177
		}
	} else {
		if base.Ui32(v10) < base.Ui32(int32(12)) {
			v71 = v5
			v72 = v10
			v73 = v16
			v74 = v16
			v75 = v16
		} else {
			v23 = v5
			v24 = v10
			v25 = v16
			v26 = v16
			v27 = v16
			for {
				v29 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
				v30 = v29 + v26
				v31 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
				v33 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
				v34 = v33 + v27
				v36 = int32(4)
				v38 = v31 + v25 - v34 ^ base.I32_rotl(v34, v36)
				v42 = v30 - v38 ^ base.I32_rotl(v38, int32(6))
				v43 = v34 + v30
				v44 = v38 + v43
				v45 = v42 + v44
				v49 = v43 - v42 ^ base.I32_rotl(v42, int32(8))
				v53 = v44 - v49 ^ base.I32_rotl(v49, int32(16))
				v57 = v45 - v53 ^ base.I32_rotl(v53, int32(19))
				v58 = v49 + v45
				v59 = v53 + v58
				v60 = v57 + v59
				v64 = v58 - v57 ^ base.I32_rotl(v57, v36)
				v65 = int32(12)
				v66 = v23 + v65
				v68 = v24 - v65
				if base.Ui32(int32(11)) < base.Ui32(v68) {
					v23 = v66
					v24 = v68
					v25 = v59
					v26 = v60
					v27 = v64
					continue
				} else {
					break
				}
				break
			}
			v71 = v66
			v72 = v68
			v73 = v59
			v74 = v60
			v75 = v64
		}
		switch v72 - int32(1) {
		case 0:
			v122 = v73
			v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71))))
			v243 = v122 + v123
			v244 = v74
			v245 = v75
		case 1:
			v117 = v73
			v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+1)))
			v122 = v118<<(uint(int32(8))%32) + v117
			v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71))))
			v243 = v122 + v123
			v244 = v74
			v245 = v75
		case 2:
			v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+2)))
			v117 = v113<<(uint(int32(16))%32) + v73
			v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+1)))
			v122 = v118<<(uint(int32(8))%32) + v117
			v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71))))
			v243 = v122 + v123
			v244 = v74
			v245 = v75
		case 3:
			v110 = v74
			v111 = *(*int32)(unsafe.Add(mBase, uint32(v71)))
			v243 = v111 + v73
			v244 = v110
			v245 = v75
		case 4:
			v107 = v74
			v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+4)))
			v110 = v107 + v108
			v111 = *(*int32)(unsafe.Add(mBase, uint32(v71)))
			v243 = v111 + v73
			v244 = v110
			v245 = v75
		case 5:
			v102 = v74
			v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+5)))
			v107 = v103<<(uint(int32(8))%32) + v102
			v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+4)))
			v110 = v107 + v108
			v111 = *(*int32)(unsafe.Add(mBase, uint32(v71)))
			v243 = v111 + v73
			v244 = v110
			v245 = v75
		case 6:
			v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+6)))
			v102 = v98<<(uint(int32(16))%32) + v74
			v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+5)))
			v107 = v103<<(uint(int32(8))%32) + v102
			v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+4)))
			v110 = v107 + v108
			v111 = *(*int32)(unsafe.Add(mBase, uint32(v71)))
			v243 = v111 + v73
			v244 = v110
			v245 = v75
		case 7:
			v93 = v75
			v94 = *(*int32)(unsafe.Add(mBase, uint32(v71)))
			v96 = *(*int32)(unsafe.Add(mBase, uint32(v71)+4))
			v243 = v94 + v73
			v244 = v96 + v74
			v245 = v93
		case 8:
			v88 = v75
			v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+8)))
			v93 = v89<<(uint(int32(8))%32) + v88
			v94 = *(*int32)(unsafe.Add(mBase, uint32(v71)))
			v96 = *(*int32)(unsafe.Add(mBase, uint32(v71)+4))
			v243 = v94 + v73
			v244 = v96 + v74
			v245 = v93
		case 9:
			v83 = v75
			v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+9)))
			v88 = v84<<(uint(int32(16))%32) + v83
			v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+8)))
			v93 = v89<<(uint(int32(8))%32) + v88
			v94 = *(*int32)(unsafe.Add(mBase, uint32(v71)))
			v96 = *(*int32)(unsafe.Add(mBase, uint32(v71)+4))
			v243 = v94 + v73
			v244 = v96 + v74
			v245 = v93
		case 10:
			v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+10)))
			v83 = v79<<(uint(int32(24))%32) + v75
			v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+9)))
			v88 = v84<<(uint(int32(16))%32) + v83
			v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+8)))
			v93 = v89<<(uint(int32(8))%32) + v88
			v94 = *(*int32)(unsafe.Add(mBase, uint32(v71)))
			v96 = *(*int32)(unsafe.Add(mBase, uint32(v71)+4))
			v243 = v94 + v73
			v244 = v96 + v74
			v245 = v93
		default:
			v243 = v73
			v244 = v74
			v245 = v75
		}
	}
	v248 = int32(14)
	v250 = v244 ^ v245 - base.I32_rotl(v244, v248)
	v254 = v250 ^ v243 - base.I32_rotl(v250, int32(11))
	v258 = v254 ^ v244 - base.I32_rotl(v254, int32(25))
	v262 = v258 ^ v250 - base.I32_rotl(v258, int32(16))
	v266 = v262 ^ v254 - base.I32_rotl(v262, int32(4))
	v270 = v266 ^ v258 - base.I32_rotl(v266, v248)
	return v270 ^ v262 - base.I32_rotl(v270, int32(24))
}
