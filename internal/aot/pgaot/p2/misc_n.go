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
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
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
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
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
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v286 int32
	_ = v286
	var v292 int32
	_ = v292
	var v297 int32
	_ = v297
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v316 int32
	_ = v316
	var v327 int32
	_ = v327
	var v333 int32
	_ = v333
	var v339 int32
	_ = v339
	var v344 int32
	_ = v344
	var v348 int32
	_ = v348
	var v355 int32
	_ = v355
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v363 int32
	_ = v363
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v393 int32
	_ = v393
	var v398 int32
	_ = v398
	var v413 int32
	_ = v413
	var v415 int32
	_ = v415
	var v439 int32
	_ = v439
	var v443 int32
	_ = v443
	var v448 int32
	_ = v448
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v459 int32
	_ = v459
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v499 int32
	_ = v499
	var v503 int32
	_ = v503
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v511 int32
	_ = v511
	var v514 int32
	_ = v514
	var v516 int32
	_ = v516
	var v518 int32
	_ = v518
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v541 int32
	_ = v541
	var v548 int32
	_ = v548
	var v552 int32
	_ = v552
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v562 int32
	_ = v562
	var v566 int32
	_ = v566
	var v568 int32
	_ = v568
	var v570 int32
	_ = v570
	var v573 int32
	_ = v573
	var v578 int32
	_ = v578
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v586 int32
	_ = v586
	var v590 int32
	_ = v590
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v604 int32
	_ = v604
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v609 int32
	_ = v609
	var v618 int32
	_ = v618
	var v621 int32
	_ = v621
	var v627 int32
	_ = v627
	var v630 int32
	_ = v630
	var v632 int32
	_ = v632
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v640 int32
	_ = v640
	var v641 int32
	_ = v641
	var v643 int32
	_ = v643
	var v644 int32
	_ = v644
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v648 int32
	_ = v648
	var v649 int32
	_ = v649
	var v654 int32
	_ = v654
	var v662 int32
	_ = v662
	var v666 int32
	_ = v666
	var v668 int32
	_ = v668
	var v669 int32
	_ = v669
	var v673 int32
	_ = v673
	var v674 int32
	_ = v674
	var v676 int32
	_ = v676
	var v680 int32
	_ = v680
	var v682 int32
	_ = v682
	var v684 int32
	_ = v684
	var v687 int32
	_ = v687
	var v692 int32
	_ = v692
	var v693 int32
	_ = v693
	var v694 int32
	_ = v694
	var v696 int32
	_ = v696
	var v697 int32
	_ = v697
	var v698 int32
	_ = v698
	var v700 int32
	_ = v700
	var v704 int32
	_ = v704
	var v709 int32
	_ = v709
	var v710 int32
	_ = v710
	var v711 int32
	_ = v711
	var v718 int32
	_ = v718
	var v720 int32
	_ = v720
	var v721 int32
	_ = v721
	var v723 int32
	_ = v723
	var v731 int32
	_ = v731
	var v732 int32
	_ = v732
	var v733 int32
	_ = v733
	var v738 int32
	_ = v738
	var v740 int32
	_ = v740
	var v745 int32
	_ = v745
	var v747 int32
	_ = v747
	var v750 int32
	_ = v750
	var v751 int32
	_ = v751
	var v755 int32
	_ = v755
	var v756 int32
	_ = v756
	var v758 int32
	_ = v758
	var v759 int32
	_ = v759
	var v761 int32
	_ = v761
	var v762 int32
	_ = v762
	var v763 int32
	_ = v763
	var v764 int32
	_ = v764
	var v770 int32
	_ = v770
	var v777 int32
	_ = v777
	var v781 int32
	_ = v781
	var v783 int32
	_ = v783
	var v784 int32
	_ = v784
	var v788 int32
	_ = v788
	var v789 int32
	_ = v789
	var v791 int32
	_ = v791
	var v795 int32
	_ = v795
	var v797 int32
	_ = v797
	var v799 int32
	_ = v799
	var v802 int32
	_ = v802
	var v807 int32
	_ = v807
	var v808 int32
	_ = v808
	var v809 int32
	_ = v809
	var v811 int32
	_ = v811
	var v812 int32
	_ = v812
	var v813 int32
	_ = v813
	var v815 int32
	_ = v815
	var v819 int32
	_ = v819
	var v824 int32
	_ = v824
	var v825 int32
	_ = v825
	var v826 int32
	_ = v826
	var v833 int32
	_ = v833
	var v835 int32
	_ = v835
	var v836 int32
	_ = v836
	var v838 int32
	_ = v838
	var v850 int32
	_ = v850
	var v852 int32
	_ = v852
	var v862 int32
	_ = v862
	var v863 int32
	_ = v863
	var v867 int32
	_ = v867
	var v870 int32
	_ = v870
	var v876 int32
	_ = v876
	var v881 int32
	_ = v881
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
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v45 = v41 + v42*int32(24)
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3))))
	if v46 != 0 {
		goto L15
	} else {
		goto L16
	}
L2:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v41 = v25
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
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v39
	v41 = v39
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
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(16)
	v37 = F_palloc(m, int32(384))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L9
	} else {
		goto L11
	}
L9:
	;
	return
L10:
	;
	v39 = v32
	goto L5
L11:
	;
	v39 = v37
	goto L5
L12:
	;
	v862 = F_pg_regerror(m, v480, v20+int32(32))
	mBase = m.M
	v863 = m.ExcPending
	if v863 != 0 {
		goto L9
	} else {
		goto L257
	}
L13:
	;
	v499 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
	v503 = l2 << (uint(int32(1)) % 32)
	v506 = v499&int32(-255) | v503&int32(254)
	v507 = int32(28)
	if v503&v507 != 0 {
		goto L150
	} else {
		goto L151
	}
L14:
	;
	v56 = m.G0
	v58 = v56 - int32(16)
	m.G0 = v58
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3))))
	if v61 == int32(0) {
		v148 = int32(1)
		goto L20
	} else {
		goto L21
	}
L15:
	;
	if v46 != int32(46) {
		goto L14
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v45)+4)) = v50&int32(-769) | int32(256)
	goto L13
L18:
	;
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+1)))
	if v49 != 0 {
		goto L14
	} else {
		goto L19
	}
L19:
	;
	goto L17
L20:
	;
	m.G0 = v58 + int32(16)
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
	v154 = v152 & int32(-769)
	if v148 != 0 {
		goto L52
	} else {
		goto L53
	}
L21:
	;
	v72 = l3
	v73 = int32(4)
	v75 = v61
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
	v148 = base.B2i32(v124 == int32(4))
	goto L20
L24:
	;
	v125 = F_pg_mblen_cstr(m, v72)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L9
	} else {
		goto L50
	}
L25:
	;
	v124 = int32(1)
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
	if v75 == int32(94) {
		goto L34
	} else {
		goto L35
	}
L28:
	;
	v84 = F_t_isalpha_cstr(m, v72)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L9
	} else {
		goto L29
	}
L29:
	;
	if v84 != 0 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v124 = int32(4)
	goto L24
L31:
	;
	goto L32
L32:
	;
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72))))
	if v87 == int32(91) {
		goto L25
	} else {
		goto L33
	}
L33:
	;
	v148 = int32(0)
	goto L20
L34:
	;
	v124 = int32(3)
	goto L24
L35:
	;
	goto L36
L36:
	;
	v94 = F_t_isalpha_cstr(m, v72)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L9
	} else {
		goto L37
	}
L37:
	;
	if v94 != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v124 = int32(2)
	goto L24
L39:
	;
	goto L40
L40:
	;
	v148 = int32(0)
	goto L20
L41:
	;
	v102 = F_t_isalpha_cstr(m, v72)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
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
	v112 = m.ExcPending
	if v112 != 0 {
		goto L9
	} else {
		goto L47
	}
L44:
	;
	if v102 != 0 {
		v124 = v73
		goto L24
	} else {
		goto L45
	}
L45:
	;
	v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72))))
	if v105 != int32(93) {
		v148 = int32(0)
		goto L20
	} else {
		goto L46
	}
L46:
	;
	v124 = int32(4)
	goto L24
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58))) = int32(3)
	F_errmsg_internal(m, int32(472186), v58)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L9
	} else {
		goto L48
	}
L48:
	;
	F_errfinish(m, int32(490897), int32(66), int32(154096))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
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
	v127 = v125 + v72
	v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127))))
	if v128 != 0 {
		v72 = v127
		v73 = v124
		v75 = v128
		goto L22
	} else {
		goto L51
	}
L51:
	;
	goto L23
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v45)+4)) = v154 | int32(512)
	v159 = v45 + int32(16)
	v163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3))))
	if v163 != 0 {
		goto L55
	} else {
		goto L56
	}
L53:
	;
	goto L54
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v45)+4)) = v154
	v450 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v451 = F_strlen(m, l3)
	mBase = m.M
	v454 = F_MemoryContextAlloc(m, v450, v451+int32(3))
	mBase = m.M
	v455 = m.ExcPending
	if v455 != 0 {
		goto L9
	} else {
		goto L140
	}
L55:
	;
	v164 = l3
	goto L57
L56:
	;
	v164 = int32(738681)
	goto L57
L57:
	;
	v165 = m.G0
	v167 = v165 - int32(80)
	m.G0 = v167
	v169 = F_strlen(m, v164)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v159))) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v159)+4)) = base.B2i32(l6 != int32(0))
	v173 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v164))))
	if v173 != 0 {
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
	v439 = m.ExcPending
	if v439 != 0 {
		goto L9
	} else {
		goto L137
	}
L60:
	;
	v175 = v169 + int32(9)
	v181 = int32(0)
	v185 = v173
	v186 = v164
	v190 = int32(4)
	goto L63
L61:
	;
	goto L62
L62:
	;
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v159)))
	if v388 != 0 {
		goto L131
	} else {
		goto L132
	}
L63:
	;
	switch v190 - int32(1) {
	case 0:
		goto L67
	default:
		goto L66
	case 3:
		goto L68
	}
L64:
	;
	if v363 != int32(4) {
		goto L59
	} else {
		goto L130
	}
L65:
	;
	v365 = F_pg_mblen_cstr(m, v186)
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L9
	} else {
		goto L128
	}
L66:
	;
	if v190&int32(-2) == int32(2) {
		goto L109
	} else {
		goto L110
	}
L67:
	;
	if v185&int32(255) == int32(94) {
		goto L94
	} else {
		goto L95
	}
L68:
	;
	v197 = F_t_isalpha_cstr(m, v186)
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L9
	} else {
		goto L69
	}
L69:
	;
	if v197 != 0 {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v199 = F_palloc0(m, v175)
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
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
		goto L83
	} else {
		goto L84
	}
L73:
	;
	if v181 != 0 {
		goto L75
	} else {
		goto L76
	}
L74:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v199)))
	*(*int32)(unsafe.Add(mBase, uint32(v199))) = v203&int32(-4) | int32(1)
	v211 = F_pg_mblen_cstr(m, v186)
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L9
	} else {
		goto L78
	}
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v181)+4)) = v199
	goto L74
L76:
	;
	goto L77
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v159))) = v199
	goto L74
L78:
	;
	if v211 != 0 {
		goto L80
	} else {
		goto L81
	}
L79:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v199)))
	*(*int32)(unsafe.Add(mBase, uint32(v199))) = v215&int32(-262141) | v211<<(uint(int32(2))%32)&int32(262140)
	v361 = v199
	v363 = int32(4)
	goto L65
L80:
	;
	v213 = F__emscripten_memcpy_bulkmem(m, v199+int32(8), v186, v211)
	mBase = m.M
	goto L82
L81:
	;
	goto L82
L82:
	;
	goto L79
L83:
	;
	v228 = F_palloc0(m, v175)
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L9
	} else {
		goto L86
	}
L84:
	;
	goto L85
L85:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L9
	} else {
		goto L91
	}
L86:
	;
	if v181 != 0 {
		goto L88
	} else {
		goto L89
	}
L87:
	;
	v232 = int32(1)
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v228)))
	*(*int32)(unsafe.Add(mBase, uint32(v228))) = v233&int32(-4) | v232
	v361 = v228
	v363 = v232
	goto L65
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v181)+4)) = v228
	goto L87
L89:
	;
	goto L90
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v159))) = v228
	goto L87
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v167)+48)) = v164
	F_errmsg_internal(m, int32(707450), v167+int32(48))
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L9
	} else {
		goto L92
	}
L92:
	;
	F_errfinish(m, int32(490897), int32(118), int32(383774))
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L9
	} else {
		goto L93
	}
L93:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L94:
	;
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v181)))
	*(*int32)(unsafe.Add(mBase, uint32(v181))) = v258&int32(-4) | int32(2)
	v361 = v181
	v363 = int32(3)
	goto L65
L95:
	;
	goto L96
L96:
	;
	v265 = F_t_isalpha_cstr(m, v186)
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L9
	} else {
		goto L97
	}
L97:
	;
	if v265 != 0 {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v269 = F_pg_mblen_cstr(m, v186)
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L9
	} else {
		goto L101
	}
L99:
	;
	goto L100
L100:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L9
	} else {
		goto L106
	}
L101:
	;
	if v269 != 0 {
		goto L103
	} else {
		goto L104
	}
L102:
	;
	v273 = int32(2)
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v181)))
	*(*int32)(unsafe.Add(mBase, uint32(v181))) = v274&int32(-262141) | v269<<(uint(v273)%32)&int32(262140)
	v361 = v181
	v363 = v273
	goto L65
L103:
	;
	v271 = F__emscripten_memcpy_bulkmem(m, v181+int32(8), v186, v269)
	mBase = m.M
	goto L105
L104:
	;
	goto L105
L105:
	;
	goto L102
L106:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v167)+64)) = v164
	F_errmsg_internal(m, int32(707450), v167-int32(-64))
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L9
	} else {
		goto L107
	}
L107:
	;
	F_errfinish(m, int32(490897), int32(133), int32(383774))
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L9
	} else {
		goto L108
	}
L108:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L109:
	;
	v302 = F_t_isalpha_cstr(m, v186)
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L9
	} else {
		goto L112
	}
L110:
	;
	goto L111
L111:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L9
	} else {
		goto L125
	}
L112:
	;
	if v302 != 0 {
		goto L113
	} else {
		goto L114
	}
L113:
	;
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v181)))
	v312 = F_pg_mblen_cstr(m, v186)
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L9
	} else {
		goto L116
	}
L114:
	;
	goto L115
L115:
	;
	v327 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v186))))
	if v327 == int32(93) {
		v361 = v181
		v363 = int32(4)
		goto L65
	} else {
		goto L121
	}
L116:
	;
	if v312 != 0 {
		goto L118
	} else {
		goto L119
	}
L117:
	;
	v316 = *(*int32)(unsafe.Add(mBase, uint32(v181)))
	*(*int32)(unsafe.Add(mBase, uint32(v181))) = (v316+v312<<(uint(int32(2))%32))&int32(262140) | v316&int32(-262141)
	v361 = v181
	v363 = v190
	goto L65
L118:
	;
	v314 = F__emscripten_memcpy_bulkmem(m, v181+int32(base.Ui32(v304)>>(uint(int32(2))%32))&int32(65535)+int32(8), v186, v312)
	mBase = m.M
	goto L120
L119:
	;
	goto L120
L120:
	;
	goto L117
L121:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L9
	} else {
		goto L122
	}
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v167)+16)) = v164
	F_errmsg_internal(m, int32(707450), v167+int32(16))
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L9
	} else {
		goto L123
	}
L123:
	;
	F_errfinish(m, int32(490897), int32(142), int32(383774))
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v167)+32)) = int32(3)
	F_errmsg_internal(m, int32(472225), v167+int32(32))
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L9
	} else {
		goto L126
	}
L126:
	;
	F_errfinish(m, int32(490897), int32(145), int32(383774))
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L9
	} else {
		goto L127
	}
L127:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L128:
	;
	v367 = v365 + v186
	v368 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v367))))
	if v368 != 0 {
		v181 = v361
		v185 = v368
		v186 = v367
		v190 = v363
		goto L63
	} else {
		goto L129
	}
L129:
	;
	goto L64
L130:
	;
	goto L62
L131:
	;
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v159)+4))
	v393 = v388
	v398 = v389
	goto L134
L132:
	;
	goto L133
L133:
	;
	m.G0 = v167 + int32(80)
	goto L58
L134:
	;
	v413 = (v398+int32(2))&int32(131070) | v398&int32(-131071)
	*(*int32)(unsafe.Add(mBase, uint32(v159)+4)) = v413
	v415 = *(*int32)(unsafe.Add(mBase, uint32(v393)+4))
	if v415 != 0 {
		v393 = v415
		v398 = v413
		goto L134
	} else {
		goto L136
	}
L135:
	;
	goto L133
L136:
	;
	goto L135
L137:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v167))) = v164
	F_errmsg_internal(m, int32(707450), v167)
	mBase = m.M
	v443 = m.ExcPending
	if v443 != 0 {
		goto L9
	} else {
		goto L138
	}
L138:
	;
	F_errfinish(m, int32(490897), int32(150), int32(383774))
	mBase = m.M
	v448 = m.ExcPending
	if v448 != 0 {
		goto L9
	} else {
		goto L139
	}
L139:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L140:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = l3
	if l6 != 0 {
		goto L141
	} else {
		goto L142
	}
L141:
	;
	v459 = int32(671426)
	goto L143
L142:
	;
	v459 = int32(175111)
	goto L143
L143:
	;
	v462 = F_pg_sprintf(m, v454, v459, v20+int32(16))
	mBase = m.M
	v463 = m.ExcPending
	if v463 != 0 {
		goto L9
	} else {
		goto L144
	}
L144:
	;
	v464 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v465 = F_strlen(m, v454)
	mBase = m.M
	v470 = F_MemoryContextAlloc(m, v464, v465<<(uint(int32(2))%32)+int32(4))
	mBase = m.M
	v471 = m.ExcPending
	if v471 != 0 {
		goto L9
	} else {
		goto L145
	}
L145:
	;
	v472 = F_pg_mb2wchar_with_len(m, v454, v470, v465)
	mBase = m.M
	v473 = m.ExcPending
	if v473 != 0 {
		goto L9
	} else {
		goto L146
	}
L146:
	;
	v475 = F_palloc(m, int32(32))
	mBase = m.M
	v476 = m.ExcPending
	if v476 != 0 {
		goto L9
	} else {
		goto L147
	}
L147:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v45)+16)) = v475
	v480 = F_pg_regcomp(m, v475, v470, v472, int32(19), int32(100))
	mBase = m.M
	v481 = m.ExcPending
	if v481 != 0 {
		goto L9
	} else {
		goto L148
	}
L148:
	;
	if v480 != 0 {
		goto L12
	} else {
		goto L149
	}
L149:
	;
	goto L13
L150:
	;
	v511 = v506
	goto L152
L151:
	;
	v511 = v506 | v507
	goto L152
L152:
	;
	if v503&int32(34) != 0 {
		goto L153
	} else {
		goto L154
	}
L153:
	;
	v514 = v511
	goto L155
L154:
	;
	v514 = v506
	goto L155
L155:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v45)+4)) = v514
	v516 = F_strlen(m, l1)
	mBase = m.M
	v518 = v516 + int32(1)
	if base.Ui32(int32(1025)) <= base.Ui32(v518) {
		goto L157
	} else {
		goto L158
	}
L156:
	;
	if (l1^v541)&int32(3) != 0 {
		goto L169
	} else {
		goto L170
	}
L157:
	;
	v521 = F_palloc0(m, v518)
	mBase = m.M
	v522 = m.ExcPending
	if v522 != 0 {
		goto L9
	} else {
		goto L160
	}
L158:
	;
	goto L159
L159:
	;
	v526 = (v516 + int32(8)) & int32(4088)
	v527 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	if base.Ui32(v526) <= base.Ui32(v527) {
		goto L162
	} else {
		goto L163
	}
L160:
	;
	v541 = v521
	goto L156
L161:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = v534 - v526
	*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = v526 + v535
	v541 = v535
	goto L156
L162:
	;
	v529 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	v534 = v527
	v535 = v529
	goto L161
L163:
	;
	goto L164
L164:
	;
	v530 = int32(8192)
	v532 = F_palloc0(m, v530)
	mBase = m.M
	v533 = m.ExcPending
	if v533 != 0 {
		goto L9
	} else {
		goto L165
	}
L165:
	;
	v534 = v530
	v535 = v532
	goto L161
L166:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v45))) = v541
	v618 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
	v621 = v618&int32(-2) | l6
	*(*int32)(unsafe.Add(mBase, uint32(v45)+4)) = v621
	if l4 == int32(0) {
		goto L188
	} else {
		goto L189
	}
L167:
	;
	goto L166
L168:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v597))) = uint8(v596)
	if v596&int32(255) == int32(0) {
		goto L167
	} else {
		goto L183
	}
L169:
	;
	v548 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	v595 = l1
	v596 = v548
	v597 = v541
	goto L168
L170:
	;
	goto L171
L171:
	;
	if l1&int32(3) != 0 {
		goto L172
	} else {
		goto L173
	}
L172:
	;
	v552 = l1
	v554 = v541
	goto L175
L173:
	;
	v566 = l1
	v568 = v541
	goto L174
L174:
	;
	v570 = *(*int32)(unsafe.Add(mBase, uint32(v566)))
	v573 = int32(-2139062144)
	if (int32(16843008)-v570|v570)&v573 != v573 {
		v595 = v566
		v596 = v570
		v597 = v568
		goto L168
	} else {
		goto L179
	}
L175:
	;
	v555 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v552))))
	*(*uint8)(unsafe.Add(mBase, uint32(v554))) = uint8(v555)
	if v555 == int32(0) {
		goto L167
	} else {
		goto L177
	}
L176:
	;
	v566 = v562
	v568 = v560
	goto L174
L177:
	;
	v559 = int32(1)
	v560 = v554 + v559
	v562 = v552 + v559
	if v562&int32(3) != 0 {
		v552 = v562
		v554 = v560
		goto L175
	} else {
		goto L178
	}
L178:
	;
	goto L176
L179:
	;
	v578 = v566
	v579 = v570
	v580 = v568
	goto L180
L180:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v580))) = v579
	v582 = int32(4)
	v583 = v580 + v582
	v584 = *(*int32)(unsafe.Add(mBase, uint32(v578)+4))
	v586 = v578 + v582
	v590 = int32(-2139062144)
	if (v584|(int32(16843008)-v584))&v590 == v590 {
		v578 = v586
		v579 = v584
		v580 = v583
		goto L180
	} else {
		goto L182
	}
L181:
	;
	v595 = v586
	v596 = v584
	v597 = v583
	goto L168
L182:
	;
	goto L181
L183:
	;
	v604 = v595
	v606 = v597
	goto L184
L184:
	;
	v607 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v604)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v606)+1)) = uint8(v607)
	v609 = int32(1)
	if v607 != 0 {
		v604 = v604 + v609
		v606 = v606 + v609
		goto L184
	} else {
		goto L186
	}
L185:
	;
	goto L167
L186:
	;
	goto L185
L187:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v45)+8)) = v733
	v738 = F_strlen(m, l5)
	mBase = m.M
	v740 = v738 & int32(16383)
	*(*int32)(unsafe.Add(mBase, uint32(v45)+4)) = v732&int32(-16776193) | v740<<(uint(int32(10))%32)
	if v740 != 0 {
		goto L223
	} else {
		goto L224
	}
L188:
	;
	v732 = v621
	v733 = int32(738681)
	goto L187
L189:
	;
	goto L190
L190:
	;
	v627 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4))))
	if v627 == int32(0) {
		v732 = v621
		v733 = int32(738681)
		goto L187
	} else {
		goto L191
	}
L191:
	;
	v630 = F_strlen(m, l4)
	mBase = m.M
	v632 = v630 + int32(1)
	if base.Ui32(int32(1025)) <= base.Ui32(v632) {
		goto L193
	} else {
		goto L194
	}
L192:
	;
	if (l4^v654)&int32(3) != 0 {
		goto L205
	} else {
		goto L206
	}
L193:
	;
	v635 = F_palloc0(m, v632)
	mBase = m.M
	v636 = m.ExcPending
	if v636 != 0 {
		goto L9
	} else {
		goto L196
	}
L194:
	;
	goto L195
L195:
	;
	v640 = (v630 + int32(8)) & int32(4088)
	v641 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	if base.Ui32(v640) <= base.Ui32(v641) {
		goto L198
	} else {
		goto L199
	}
L196:
	;
	v654 = v635
	goto L192
L197:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = v648 - v640
	*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = v649 + v640
	v654 = v649
	goto L192
L198:
	;
	v643 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	v648 = v641
	v649 = v643
	goto L197
L199:
	;
	goto L200
L200:
	;
	v644 = int32(8192)
	v646 = F_palloc0(m, v644)
	mBase = m.M
	v647 = m.ExcPending
	if v647 != 0 {
		goto L9
	} else {
		goto L201
	}
L201:
	;
	v648 = v644
	v649 = v646
	goto L197
L202:
	;
	v731 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
	v732 = v731
	v733 = v654
	goto L187
L203:
	;
	goto L202
L204:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v711))) = uint8(v710)
	if v710&int32(255) == int32(0) {
		goto L203
	} else {
		goto L219
	}
L205:
	;
	v662 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4))))
	v709 = l4
	v710 = v662
	v711 = v654
	goto L204
L206:
	;
	goto L207
L207:
	;
	if l4&int32(3) != 0 {
		goto L208
	} else {
		goto L209
	}
L208:
	;
	v666 = l4
	v668 = v654
	goto L211
L209:
	;
	v680 = l4
	v682 = v654
	goto L210
L210:
	;
	v684 = *(*int32)(unsafe.Add(mBase, uint32(v680)))
	v687 = int32(-2139062144)
	if (int32(16843008)-v684|v684)&v687 != v687 {
		v709 = v680
		v710 = v684
		v711 = v682
		goto L204
	} else {
		goto L215
	}
L211:
	;
	v669 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v666))))
	*(*uint8)(unsafe.Add(mBase, uint32(v668))) = uint8(v669)
	if v669 == int32(0) {
		goto L203
	} else {
		goto L213
	}
L212:
	;
	v680 = v676
	v682 = v674
	goto L210
L213:
	;
	v673 = int32(1)
	v674 = v668 + v673
	v676 = v666 + v673
	if v676&int32(3) != 0 {
		v666 = v676
		v668 = v674
		goto L211
	} else {
		goto L214
	}
L214:
	;
	goto L212
L215:
	;
	v692 = v680
	v693 = v684
	v694 = v682
	goto L216
L216:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v694))) = v693
	v696 = int32(4)
	v697 = v694 + v696
	v698 = *(*int32)(unsafe.Add(mBase, uint32(v692)+4))
	v700 = v692 + v696
	v704 = int32(-2139062144)
	if (v698|(int32(16843008)-v698))&v704 == v704 {
		v692 = v700
		v693 = v698
		v694 = v697
		goto L216
	} else {
		goto L218
	}
L217:
	;
	v709 = v700
	v710 = v698
	v711 = v697
	goto L204
L218:
	;
	goto L217
L219:
	;
	v718 = v709
	v720 = v711
	goto L220
L220:
	;
	v721 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v718)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v720)+1)) = uint8(v721)
	v723 = int32(1)
	if v721 != 0 {
		v718 = v718 + v723
		v720 = v720 + v723
		goto L220
	} else {
		goto L222
	}
L221:
	;
	goto L203
L222:
	;
	goto L221
L223:
	;
	v745 = F_strlen(m, l5)
	mBase = m.M
	v747 = v745 + int32(1)
	if base.Ui32(int32(1025)) <= base.Ui32(v747) {
		goto L227
	} else {
		goto L228
	}
L224:
	;
	v850 = int32(738681)
	goto L225
L225:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v45)+12)) = v850
	v852 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v852 + int32(1)
	m.G0 = v20 + int32(144)
	return
L226:
	;
	if (l5^v770)&int32(3) != 0 {
		goto L239
	} else {
		goto L240
	}
L227:
	;
	v750 = F_palloc0(m, v747)
	mBase = m.M
	v751 = m.ExcPending
	if v751 != 0 {
		goto L9
	} else {
		goto L230
	}
L228:
	;
	goto L229
L229:
	;
	v755 = (v745 + int32(8)) & int32(4088)
	v756 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	if base.Ui32(v755) <= base.Ui32(v756) {
		goto L232
	} else {
		goto L233
	}
L230:
	;
	v770 = v750
	goto L226
L231:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = v763 - v755
	*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = v764 + v755
	v770 = v764
	goto L226
L232:
	;
	v758 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	v763 = v756
	v764 = v758
	goto L231
L233:
	;
	goto L234
L234:
	;
	v759 = int32(8192)
	v761 = F_palloc0(m, v759)
	mBase = m.M
	v762 = m.ExcPending
	if v762 != 0 {
		goto L9
	} else {
		goto L235
	}
L235:
	;
	v763 = v759
	v764 = v761
	goto L231
L236:
	;
	v850 = v770
	goto L225
L237:
	;
	goto L236
L238:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v826))) = uint8(v825)
	if v825&int32(255) == int32(0) {
		goto L237
	} else {
		goto L253
	}
L239:
	;
	v777 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l5))))
	v824 = l5
	v825 = v777
	v826 = v770
	goto L238
L240:
	;
	goto L241
L241:
	;
	if l5&int32(3) != 0 {
		goto L242
	} else {
		goto L243
	}
L242:
	;
	v781 = l5
	v783 = v770
	goto L245
L243:
	;
	v795 = l5
	v797 = v770
	goto L244
L244:
	;
	v799 = *(*int32)(unsafe.Add(mBase, uint32(v795)))
	v802 = int32(-2139062144)
	if (int32(16843008)-v799|v799)&v802 != v802 {
		v824 = v795
		v825 = v799
		v826 = v797
		goto L238
	} else {
		goto L249
	}
L245:
	;
	v784 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v781))))
	*(*uint8)(unsafe.Add(mBase, uint32(v783))) = uint8(v784)
	if v784 == int32(0) {
		goto L237
	} else {
		goto L247
	}
L246:
	;
	v795 = v791
	v797 = v789
	goto L244
L247:
	;
	v788 = int32(1)
	v789 = v783 + v788
	v791 = v781 + v788
	if v791&int32(3) != 0 {
		v781 = v791
		v783 = v789
		goto L245
	} else {
		goto L248
	}
L248:
	;
	goto L246
L249:
	;
	v807 = v795
	v808 = v799
	v809 = v797
	goto L250
L250:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v809))) = v808
	v811 = int32(4)
	v812 = v809 + v811
	v813 = *(*int32)(unsafe.Add(mBase, uint32(v807)+4))
	v815 = v807 + v811
	v819 = int32(-2139062144)
	if (v813|(int32(16843008)-v813))&v819 == v819 {
		v807 = v815
		v808 = v813
		v809 = v812
		goto L250
	} else {
		goto L252
	}
L251:
	;
	v824 = v815
	v825 = v813
	v826 = v812
	goto L238
L252:
	;
	goto L251
L253:
	;
	v833 = v824
	v835 = v826
	goto L254
L254:
	;
	v836 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v833)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v835)+1)) = uint8(v836)
	v838 = int32(1)
	if v836 != 0 {
		v833 = v833 + v838
		v835 = v835 + v838
		goto L254
	} else {
		goto L256
	}
L255:
	;
	goto L237
L256:
	;
	goto L255
L257:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v867 = m.ExcPending
	if v867 != 0 {
		goto L9
	} else {
		goto L258
	}
L258:
	;
	F_errcode(m, int32(302252162))
	mBase = m.M
	v870 = m.ExcPending
	if v870 != 0 {
		goto L9
	} else {
		goto L259
	}
L259:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v20 + int32(32)
	F_errmsg(m, int32(201089), v20)
	mBase = m.M
	v876 = m.ExcPending
	if v876 != 0 {
		goto L9
	} else {
		goto L260
	}
L260:
	;
	F_errfinish(m, int32(494186), int32(752), int32(27206))
	mBase = m.M
	v881 = m.ExcPending
	if v881 != 0 {
		goto L9
	} else {
		goto L261
	}
L261:
	;
	base.Wasm_trap_unreachable()
	for {
	}
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
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	if l0 == int32(14) {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
		if v9 != 0 {
			v10 = int32(546508)
		} else {
			v10 = int32(531221)
		}
		return v10
	} else {
		v13 = l0 >> (uint(int32(16)) % 32)
		v14 = int32(65535)
		v15 = l0 & v14
		if v15 != v14 {
			v29 = int32(738681)
			switch v13 - int32(1) {
			case 0:
				if base.Ui32(int32(1)) < base.Ui32(v15) {
					v54 = v29
					return v54
				} else {
					v41 = int32(4072520)
					if v15 == int32(0) {
						return v41
					} else {
						v45 = v41
						v47 = v15
						for {
							v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45))))
							v51 = v45 + int32(1)
							if v49 != 0 {
								v45 = v51
								continue
							} else {
							}
							v53 = v47 - int32(1)
							if v53 != 0 {
								v45 = v51
								v47 = v53
								continue
							} else {
								break
							}
							break
						}
						v54 = v51
						return v54
					}
				}
			case 1:
				if base.Ui32(int32(49)) < base.Ui32(v15) {
					v54 = v29
					return v54
				} else {
					v41 = int32(4072528)
					if v15 == int32(0) {
						return v41
					} else {
						v45 = v41
						v47 = v15
						for {
							v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45))))
							v51 = v45 + int32(1)
							if v49 != 0 {
								v45 = v51
								continue
							} else {
							}
							v53 = v47 - int32(1)
							if v53 != 0 {
								v45 = v51
								v47 = v53
								continue
							} else {
								break
							}
							break
						}
						v54 = v51
						return v54
					}
				}
			default:
				v54 = v29
				return v54
			case 4:
				if base.Ui32(int32(3)) < base.Ui32(v15) {
					v54 = v29
					return v54
				} else {
					v41 = int32(4072848)
					if v15 == int32(0) {
						return v41
					} else {
						v45 = v41
						v47 = v15
						for {
							v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45))))
							v51 = v45 + int32(1)
							if v49 != 0 {
								v45 = v51
								continue
							} else {
							}
							v53 = v47 - int32(1)
							if v53 != 0 {
								v45 = v51
								v47 = v53
								continue
							} else {
								break
							}
							break
						}
						v54 = v51
						return v54
					}
				}
			}
		} else {
			if int32(5) < v13 {
				v29 = int32(738681)
				switch v13 - int32(1) {
				case 0:
					if base.Ui32(int32(1)) < base.Ui32(v15) {
						v54 = v29
						return v54
					} else {
						v41 = int32(4072520)
						if v15 == int32(0) {
							return v41
						} else {
							v45 = v41
							v47 = v15
							for {
								v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45))))
								v51 = v45 + int32(1)
								if v49 != 0 {
									v45 = v51
									continue
								} else {
								}
								v53 = v47 - int32(1)
								if v53 != 0 {
									v45 = v51
									v47 = v53
									continue
								} else {
									break
								}
								break
							}
							v54 = v51
							return v54
						}
					}
				case 1:
					if base.Ui32(int32(49)) < base.Ui32(v15) {
						v54 = v29
						return v54
					} else {
						v41 = int32(4072528)
						if v15 == int32(0) {
							return v41
						} else {
							v45 = v41
							v47 = v15
							for {
								v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45))))
								v51 = v45 + int32(1)
								if v49 != 0 {
									v45 = v51
									continue
								} else {
								}
								v53 = v47 - int32(1)
								if v53 != 0 {
									v45 = v51
									v47 = v53
									continue
								} else {
									break
								}
								break
							}
							v54 = v51
							return v54
						}
					}
				default:
					v54 = v29
					return v54
				case 4:
					if base.Ui32(int32(3)) < base.Ui32(v15) {
						v54 = v29
						return v54
					} else {
						v41 = int32(4072848)
						if v15 == int32(0) {
							return v41
						} else {
							v45 = v41
							v47 = v15
							for {
								v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45))))
								v51 = v45 + int32(1)
								if v49 != 0 {
									v45 = v51
									continue
								} else {
								}
								v53 = v47 - int32(1)
								if v53 != 0 {
									v45 = v51
									v47 = v53
									continue
								} else {
									break
								}
								break
							}
							v54 = v51
							return v54
						}
					}
				}
			} else {
				v23 = *(*int32)(unsafe.Add(mBase, uint32(l1+v13<<(uint(int32(2))%32))))
				if v23 != 0 {
					v27 = v23 + int32(8)
				} else {
					v27 = int32(541278)
				}
				return v27
			}
		}
	}
}
func F_namefastcmp_locale(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v4 = F_strlen(m, l0)
	v5 = F_strlen(m, l1)
	v6 = F_varstrfastcmp_locale(m, l0, v4, l1, v5, l2)
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return v6
	}
}
func F_namelike(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v8 = F_pg_detoast_datum_packed(m, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		v12 = int32(1)
		v13 = v8 + v12
		v14 = F_strlen(m, v6)
		mBase = m.M
		v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8))))
		v19 = v17 & v12
		if v19 != 0 {
			v20 = v13
		} else {
			v20 = v8 + int32(4)
		}
		if v17 == int32(1) {
			v23 = int32(4)
			v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
			if v25&int32(254) == int32(2) {
				v34 = v23
			} else {
				v34 = base.B2i32(v25 == int32(18)) << (uint(v23) % 32)
			}
			if v25 == int32(1) {
				v37 = v23
			} else {
				v37 = v34
			}
			v48 = v37
		} else {
			v38 = int32(1)
			if v19 != 0 {
				v48 = int32(base.Ui32(v17)>>(uint(v38)%32)) - v38
			} else {
				v42 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
				v48 = int32(base.Ui32(v42)>>(uint(int32(2))%32)) - int32(4)
			}
		}
		v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v50 = F_GenericMatchText(m, v6, v14, v20, v48, v49)
		mBase = m.M
		v51 = m.ExcPending
		if v51 != 0 {
			return int32(0)
		} else {
			return base.B2i32(v50 == int32(1))
		}
	}
}
func F_nameregexeq(m *base.Module, l0 int32) int32 {
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
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
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
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v7 = F_pg_detoast_datum_packed(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		v11 = F_strlen(m, v5)
		mBase = m.M
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v14 = F_RE_compile_and_cache(m, v7, int32(19), v13)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			v20 = F_palloc(m, v11<<(uint(int32(2))%32)+int32(4))
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return int32(0)
			} else {
				v22 = F_pg_mb2wchar_with_len(m, v5, v20, v11)
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return int32(0)
				} else {
					v24 = int32(0)
					v27 = F_RE_wchar_execute(m, v20, v22, v24, v24, v24)
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return int32(0)
					} else {
						F_pfree(m, v20)
						mBase = m.M
						v30 = m.ExcPending
						if v30 != 0 {
							return int32(0)
						} else {
							return v27
						}
					}
				}
			}
		}
	}
}
func F_namestrcpy(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
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
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	v3 = int32(64)
	if (l1^l0)&int32(3) != 0 {
		v73 = l1
		v74 = v3
		v75 = l0
		goto L5
	} else {
		goto L6
	}
L1:
	;
	v112 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+63)) = uint8(v112)
	return
L2:
	;
	v111 = F___memset(m, v108, int32(0), v107)
	mBase = m.M
	goto L1
L3:
	;
	v107 = int32(0)
	v108 = v102
	goto L2
L4:
	;
	v85 = v80
	v86 = v81
	v87 = v82
	goto L23
L5:
	;
	if v74 == int32(0) {
		v102 = v75
		goto L3
	} else {
		goto L22
	}
L6:
	;
	if l1&int32(3) == int32(0) {
		v39 = l1
		v40 = v3
		v41 = l0
		v42 = int32(1)
		goto L7
	} else {
		goto L8
	}
L7:
	;
	if v42 == int32(0) {
		v102 = v41
		goto L3
	} else {
		goto L15
	}
L8:
	;
	goto L9
L9:
	;
	v18 = l1
	v19 = v3
	v20 = l0
	goto L10
L10:
	;
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
	*(*uint8)(unsafe.Add(mBase, uint32(v20))) = uint8(v22)
	if v22 == int32(0) {
		v107 = v19
		v108 = v20
		goto L2
	} else {
		goto L12
	}
L11:
	;
	v39 = v33
	v40 = v29
	v41 = v27
	v42 = v31
	goto L7
L12:
	;
	v26 = int32(1)
	v27 = v20 + v26
	v29 = v19 - v26
	v30 = int32(0)
	v31 = base.B2i32(v29 != v30)
	v33 = v18 + v26
	if v33&int32(3) == v30 {
		v39 = v33
		v40 = v29
		v41 = v27
		v42 = v31
		goto L7
	} else {
		goto L13
	}
L13:
	;
	if v29 != 0 {
		v18 = v33
		v19 = v29
		v20 = v27
		goto L10
	} else {
		goto L14
	}
L14:
	;
	goto L11
L15:
	;
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39))))
	if v45 == int32(0) {
		v107 = v40
		v108 = v41
		goto L2
	} else {
		goto L16
	}
L16:
	;
	if base.Ui32(v40) < base.Ui32(int32(4)) {
		v73 = v39
		v74 = v40
		v75 = v41
		goto L5
	} else {
		goto L17
	}
L17:
	;
	v51 = v39
	v52 = v40
	v53 = v41
	goto L18
L18:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
	v59 = int32(-2139062144)
	if (int32(16843008)-v56|v56)&v59 != v59 {
		v80 = v51
		v81 = v52
		v82 = v53
		goto L4
	} else {
		goto L20
	}
L19:
	;
	v73 = v67
	v74 = v69
	v75 = v65
	goto L5
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v53))) = v56
	v64 = int32(4)
	v65 = v53 + v64
	v67 = v51 + v64
	v69 = v52 - v64
	if base.Ui32(int32(3)) < base.Ui32(v69) {
		v51 = v67
		v52 = v69
		v53 = v65
		goto L18
	} else {
		goto L21
	}
L21:
	;
	goto L19
L22:
	;
	v80 = v73
	v81 = v74
	v82 = v75
	goto L4
L23:
	;
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85))))
	*(*uint8)(unsafe.Add(mBase, uint32(v87))) = uint8(v89)
	if v89 == int32(0) {
		v107 = v86
		v108 = v87
		goto L2
	} else {
		goto L25
	}
L24:
	;
	v102 = v94
	goto L3
L25:
	;
	v93 = int32(1)
	v94 = v87 + v93
	v98 = v86 - v93
	if v98 != 0 {
		v85 = v85 + v93
		v86 = v98
		v87 = v94
		goto L23
	} else {
		goto L26
	}
L26:
	;
	goto L24
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
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
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
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v231 int32
	_ = v231
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	v2 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	if l0 != 0 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	m.G0 = v7 + int32(32)
	return v244
L2:
	;
	v241 = F_make_orclause(m, v238)
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L17
	} else {
		goto L68
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L17
	} else {
		goto L65
	}
L4:
	;
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch v9 - int32(7) {
	case 0:
		goto L13
	default:
		goto L7
	case 10:
		goto L12
	case 13:
		goto L11
	case 14:
		goto L10
	case 45:
		goto L9
	case 46:
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L17
	} else {
		goto L62
	}
L7:
	;
	v186 = m.G0
	v187 = int32(16)
	v188 = v186 - v187
	m.G0 = v188
	v191 = F_palloc0(m, v187)
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L17
	} else {
		goto L60
	}
L8:
	;
	v168 = F_palloc0(m, int32(16))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L17
	} else {
		goto L58
	}
L9:
	;
	v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)))
	if v151 != 0 {
		goto L7
	} else {
		goto L56
	}
L10:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	switch v77 {
	case 0:
		goto L29
	case 1:
		goto L28
	case 2:
		goto L27
	default:
		goto L26
	}
L11:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v53 = F_get_negator(m, v52)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L17
	} else {
		goto L23
	}
L12:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v28 = F_get_negator(m, v27)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L17
	} else {
		goto L20
	}
L13:
	;
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v12 == int32(1) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v17 = F_makeBoolConst(m, int32(0), int32(1))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	goto L16
L16:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v22 = int32(0)
	v25 = F_makeBoolConst(m, base.B2i32(v21 == v22), v22)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L17
	} else {
		goto L19
	}
L17:
	;
	return int32(0)
L18:
	;
	v244 = v17
	goto L1
L19:
	;
	v244 = v25
	goto L1
L20:
	;
	if v28 == int32(0) {
		goto L7
	} else {
		goto L21
	}
L21:
	;
	v33 = F_palloc0(m, int32(36))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L17
	} else {
		goto L22
	}
L22:
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
	v244 = v33
	goto L1
L23:
	;
	if v53 == int32(0) {
		goto L7
	} else {
		goto L24
	}
L24:
	;
	v58 = F_palloc0(m, int32(36))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L17
	} else {
		goto L25
	}
L25:
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
	v244 = v58
	goto L1
L26:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L17
	} else {
		goto L53
	}
L27:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v134)+12))
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v135)))
	v244 = v136
	goto L1
L28:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v101 == int32(0) {
		goto L40
	} else {
		goto L41
	}
L29:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v78 != 0 {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	v84 = v79
	v85 = v2
	goto L35
L31:
	;
	v79 = int32(0)
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v78)+4))
	if v79 < v80 {
		goto L30
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	v238 = v2
	goto L2
L34:
	;
	goto L33
L35:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v78)+12))
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v88+v84<<(uint(int32(2))%32))))
	v93 = F_negate_clause(m, v92)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L17
	} else {
		goto L37
	}
L36:
	;
	v238 = v95
	goto L2
L37:
	;
	v95 = F_lappend(m, v85, v93)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L17
	} else {
		goto L38
	}
L38:
	;
	v98 = v84 + int32(1)
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v78)+4))
	if v98 < v99 {
		v84 = v98
		v85 = v95
		goto L35
	} else {
		goto L39
	}
L39:
	;
	goto L36
L40:
	;
	v105 = F_make_andclause(m, int32(0))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L17
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v101)+4))
	if int32(0) < v107 {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	v244 = v105
	goto L1
L44:
	;
	v111 = int32(0)
	v113 = v2
	goto L47
L45:
	;
	v130 = v2
	goto L46
L46:
	;
	v132 = F_make_andclause(m, v130)
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L17
	} else {
		goto L52
	}
L47:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v101)+12))
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v115+v111<<(uint(int32(2))%32))))
	v120 = F_negate_clause(m, v119)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L17
	} else {
		goto L49
	}
L48:
	;
	v130 = v122
	goto L46
L49:
	;
	v122 = F_lappend(m, v113, v120)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L17
	} else {
		goto L50
	}
L50:
	;
	v125 = v111 + int32(1)
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v101)+4))
	if v125 < v126 {
		v111 = v125
		v113 = v122
		goto L47
	} else {
		goto L51
	}
L51:
	;
	goto L48
L52:
	;
	v244 = v132
	goto L1
L53:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v141
	F_errmsg_internal(m, int32(479617), v7)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L17
	} else {
		goto L54
	}
L54:
	;
	F_errfinish(m, int32(494352), int32(195), int32(355811))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L17
	} else {
		goto L55
	}
L55:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L56:
	;
	v153 = F_palloc0(m, int32(20))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L17
	} else {
		goto L57
	}
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v153))) = int32(52)
	v157 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v153)+4)) = v157
	v159 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v153)+8)) = base.B2i32(v159 == int32(0))
	v163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)))
	*(*uint8)(unsafe.Add(mBase, uint32(v153)+12)) = uint8(v163)
	v165 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v153)+16)) = v165
	v244 = v153
	goto L1
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v168))) = int32(53)
	v172 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v168)+4)) = v172
	v174 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if base.Ui32(int32(6)) <= base.Ui32(v174) {
		goto L3
	} else {
		goto L59
	}
L59:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v174<<(uint(int32(2))%32))+uint32(_consts[508])))
	*(*int32)(unsafe.Add(mBase, uint32(v168)+8)) = v181
	v183 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v168)+12)) = v183
	v244 = v168
	goto L1
L60:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v191))) = int64(8589934613)
	*(*int32)(unsafe.Add(mBase, uint32(v188)+8)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v188)+12)) = l0
	v200 = F_list_make1_impl(m, int32(1), v188+int32(8))
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L17
	} else {
		goto L61
	}
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v191)+12)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v191)+8)) = v200
	m.G0 = v188 + int32(16)
	v244 = v191
	goto L1
L62:
	;
	F_errmsg_internal(m, int32(267468), int32(0))
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L17
	} else {
		goto L63
	}
L63:
	;
	F_errfinish(m, int32(494352), int32(76), int32(355811))
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L17
	} else {
		goto L64
	}
L64:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L65:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v225
	F_errmsg_internal(m, int32(480844), v7+int32(16))
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L17
	} else {
		goto L66
	}
L66:
	;
	F_errfinish(m, int32(494352), int32(250), int32(355811))
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L17
	} else {
		goto L67
	}
L67:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L68:
	;
	v244 = v241
	goto L1
}
func F_neqjoinsel(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v30 int32
	_ = v30
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
	var v38 float32
	_ = v38
	var v42 float64
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 float64
	_ = v64
	var v69 float64
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	v9 = m.G0
	v11 = v9 - int32(80)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v16&int32(65534) == int32(4) {
		F_get_join_variables(m, v15, v14, v13, v11+int32(48), v11+int32(16), v11+int32(15))
		mBase = m.M
		v30 = m.ExcPending
		if v30 != 0 {
			return int32(0)
		} else {
			v31 = *(*int32)(unsafe.Add(mBase, uint32(v11)+24))
			v32 = *(*int32)(unsafe.Add(mBase, uint32(v11)+56))
			v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+15)))
			if v33 != 0 {
				v34 = v31
			} else {
				v34 = v32
			}
			if v34 != 0 {
				v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)+16))
				v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+22)))
				v38 = *(*float32)(unsafe.Add(mBase, uint32(v35+v36)+8))
				v42 = base.F64_promote_f32(v38)
			} else {
				v42 = float64(0)
			}
			if v32 != 0 {
				v43 = *(*int32)(unsafe.Add(mBase, uint32(v11)+60))
				m.T0[v43].(func(*base.Module, int32))(m, v32)
				mBase = m.M
				v45 = m.ExcPending
				if v45 != 0 {
					return int32(0)
				} else {
					v46 = *(*int32)(unsafe.Add(mBase, uint32(v11)+24))
					v47 = v46
					if v47 == int32(0) {
						v69 = v42
						v72 = F_Float8GetDatum(m, base.F64_sub(float64(1), v69))
						mBase = m.M
						v73 = m.ExcPending
						if v73 != 0 {
							return int32(0)
						} else {
							m.G0 = v11 + int32(80)
							return v72
						}
					} else {
						v50 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
						m.T0[v50].(func(*base.Module, int32))(m, v47)
						mBase = m.M
						v52 = m.ExcPending
						if v52 != 0 {
							return int32(0)
						} else {
							v69 = v42
							v72 = F_Float8GetDatum(m, base.F64_sub(float64(1), v69))
							mBase = m.M
							v73 = m.ExcPending
							if v73 != 0 {
								return int32(0)
							} else {
								m.G0 = v11 + int32(80)
								return v72
							}
						}
					}
				}
			} else {
				v47 = v31
				if v47 == int32(0) {
					v69 = v42
					v72 = F_Float8GetDatum(m, base.F64_sub(float64(1), v69))
					mBase = m.M
					v73 = m.ExcPending
					if v73 != 0 {
						return int32(0)
					} else {
						m.G0 = v11 + int32(80)
						return v72
					}
				} else {
					v50 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
					m.T0[v50].(func(*base.Module, int32))(m, v47)
					mBase = m.M
					v52 = m.ExcPending
					if v52 != 0 {
						return int32(0)
					} else {
						v69 = v42
						v72 = F_Float8GetDatum(m, base.F64_sub(float64(1), v69))
						mBase = m.M
						v73 = m.ExcPending
						if v73 != 0 {
							return int32(0)
						} else {
							m.G0 = v11 + int32(80)
							return v72
						}
					}
				}
			}
		}
	} else {
		v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v55 = F_get_negator(m, v54)
		mBase = m.M
		v56 = m.ExcPending
		if v56 != 0 {
			return int32(0)
		} else {
			if v55 == int32(0) {
				v69 = float64(0.005)
				v72 = F_Float8GetDatum(m, base.F64_sub(float64(1), v69))
				mBase = m.M
				v73 = m.ExcPending
				if v73 != 0 {
					return int32(0)
				} else {
					m.G0 = v11 + int32(80)
					return v72
				}
			} else {
				v62 = F_DirectFunctionCall5Coll(m, int32(1509), v53, v15, v55, v14, base.I32_extend16_s(v16), v13)
				mBase = m.M
				v63 = m.ExcPending
				if v63 != 0 {
					return int32(0)
				} else {
					v64 = *(*float64)(unsafe.Add(mBase, uint32(v62)))
					v69 = v64
					v72 = F_Float8GetDatum(m, base.F64_sub(float64(1), v69))
					mBase = m.M
					v73 = m.ExcPending
					if v73 != 0 {
						return int32(0)
					} else {
						m.G0 = v11 + int32(80)
						return v72
					}
				}
			}
		}
	}
}
func F_neqsel(m *base.Module, l0 int32) int32 {
	var v3 float64
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	v3 = F_eqsel_internal(m, l0, int32(1))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v7 = F_Float8GetDatum(m, v3)
		v8 = m.ExcPending
		if v8 != 0 {
			return int32(0)
		} else {
			return v7
		}
	}
}
func F_networksel(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 float64
	_ = v7
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v58 float64
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v71 float64
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v88 float64
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 float32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v106 int32
	_ = v106
	var v109 float64
	_ = v109
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 float64
	_ = v125
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v135 float64
	_ = v135
	var v136 float64
	_ = v136
	var v137 float64
	_ = v137
	var v140 float64
	_ = v140
	var v143 float64
	_ = v143
	var v151 float64
	_ = v151
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v161 float64
	_ = v161
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	v7 = float64(0)
	v11 = m.G0
	v13 = v11 - int32(128)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	switch v19 - int32(931) {
	case 0:
		v43 = int32(2)
		v50 = F_get_restriction_variable(m, v17, v16, v15, v13+int32(96), v13+int32(92), v13+int32(91))
		mBase = m.M
		v51 = m.ExcPending
		if v51 != 0 {
			return int32(0)
		} else {
			if v50 == int32(0) {
				if v19 == int32(3552) {
					v58 = float64(0.01)
				} else {
					v58 = float64(0.005)
				}
				v161 = v58
				v165 = F_Float8GetDatum(m, v161)
				mBase = m.M
				v166 = m.ExcPending
				if v166 != 0 {
					return int32(0)
				} else {
					m.G0 = v13 + int32(128)
					return v165
				}
			} else {
				v59 = *(*int32)(unsafe.Add(mBase, uint32(v13)+92))
				v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
				if v60 != int32(7) {
					v63 = *(*int32)(unsafe.Add(mBase, uint32(v13)+104))
					if v63 != 0 {
						v64 = *(*int32)(unsafe.Add(mBase, uint32(v13)+108))
						m.T0[v64].(func(*base.Module, int32))(m, v63)
						mBase = m.M
						v66 = m.ExcPending
						if v66 != 0 {
							return int32(0)
						} else {
							if v19 == int32(3552) {
								v71 = float64(0.01)
							} else {
								v71 = float64(0.005)
							}
							v161 = v71
							v165 = F_Float8GetDatum(m, v161)
							mBase = m.M
							v166 = m.ExcPending
							if v166 != 0 {
								return int32(0)
							} else {
								m.G0 = v13 + int32(128)
								return v165
							}
						}
					} else {
						if v19 == int32(3552) {
							v71 = float64(0.01)
						} else {
							v71 = float64(0.005)
						}
						v161 = v71
						v165 = F_Float8GetDatum(m, v161)
						mBase = m.M
						v166 = m.ExcPending
						if v166 != 0 {
							return int32(0)
						} else {
							m.G0 = v13 + int32(128)
							return v165
						}
					}
				} else {
					v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+24)))
					if v72 == int32(1) {
						v75 = *(*int32)(unsafe.Add(mBase, uint32(v13)+104))
						if v75 == int32(0) {
							v161 = v7
							v165 = F_Float8GetDatum(m, v161)
							mBase = m.M
							v166 = m.ExcPending
							if v166 != 0 {
								return int32(0)
							} else {
								m.G0 = v13 + int32(128)
								return v165
							}
						} else {
							v78 = *(*int32)(unsafe.Add(mBase, uint32(v13)+108))
							m.T0[v78].(func(*base.Module, int32))(m, v75)
							mBase = m.M
							v80 = m.ExcPending
							if v80 != 0 {
								return int32(0)
							} else {
								v161 = v7
								v165 = F_Float8GetDatum(m, v161)
								mBase = m.M
								v166 = m.ExcPending
								if v166 != 0 {
									return int32(0)
								} else {
									m.G0 = v13 + int32(128)
									return v165
								}
							}
						}
					} else {
						v81 = *(*int32)(unsafe.Add(mBase, uint32(v13)+104))
						if v81 == int32(0) {
							if v19 == int32(3552) {
								v88 = float64(0.01)
							} else {
								v88 = float64(0.005)
							}
							v161 = v88
							v165 = F_Float8GetDatum(m, v161)
							mBase = m.M
							v166 = m.ExcPending
							if v166 != 0 {
								return int32(0)
							} else {
								m.G0 = v13 + int32(128)
								return v165
							}
						} else {
							v89 = *(*int32)(unsafe.Add(mBase, uint32(v59)+20))
							v90 = *(*int32)(unsafe.Add(mBase, uint32(v81)+16))
							v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90)+22)))
							v93 = *(*float32)(unsafe.Add(mBase, uint32(v90+v91)+8))
							v94 = F_get_opcode(m, v19)
							mBase = m.M
							v95 = m.ExcPending
							if v95 != 0 {
								return int32(0)
							} else {
								F_fmgr_info(m, v94, v13+int32(12))
								mBase = m.M
								v99 = m.ExcPending
								if v99 != 0 {
									return int32(0)
								} else {
									v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+91)))
									v109 = F_mcv_selectivity(m, v13+int32(96), v13+int32(12), int32(0), v89, v106, v13+int32(40))
									mBase = m.M
									v110 = m.ExcPending
									if v110 != 0 {
										return int32(0)
									} else {
										v113 = *(*int32)(unsafe.Add(mBase, uint32(v13)+104))
										v117 = F_get_attstatsslot(m, v13+int32(52), v113, int32(2), int32(0), int32(1))
										mBase = m.M
										v118 = m.ExcPending
										if v118 != 0 {
											return int32(0)
										} else {
											if v117 != 0 {
												v119 = *(*int32)(unsafe.Add(mBase, uint32(v13)+64))
												v120 = *(*int32)(unsafe.Add(mBase, uint32(v13)+68))
												v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+91)))
												if v123 != 0 {
													v124 = v43
												} else {
													v124 = int32(0) - v43
												}
												v125 = F_inet_hist_value_sel(m, v119, v120, v89, v124)
												mBase = m.M
												v126 = m.ExcPending
												if v126 != 0 {
													return int32(0)
												} else {
													F_free_attstatsslot(m, v13+int32(52))
													mBase = m.M
													v130 = m.ExcPending
													if v130 != 0 {
														return int32(0)
													} else {
														v136 = v125
														v137 = float64(0)
														v140 = *(*float64)(unsafe.Add(mBase, uint32(v13)+40))
														v143 = base.F64_add(base.F64_mul(base.F64_sub(base.F64_sub(float64(1), base.F64_promote_f32(v93)), v140), v136), v109)
														if base.F64_lt(v143, v137) != 0 {
															v151 = v137
														} else {
															if base.F64_gt(v143, float64(1)) == int32(0) {
																v151 = v143
															} else {
																v151 = float64(1)
															}
														}
														v152 = *(*int32)(unsafe.Add(mBase, uint32(v13)+104))
														if v152 == int32(0) {
															v161 = v151
															v165 = F_Float8GetDatum(m, v161)
															mBase = m.M
															v166 = m.ExcPending
															if v166 != 0 {
																return int32(0)
															} else {
																m.G0 = v13 + int32(128)
																return v165
															}
														} else {
															v155 = *(*int32)(unsafe.Add(mBase, uint32(v13)+108))
															m.T0[v155].(func(*base.Module, int32))(m, v152)
															mBase = m.M
															v157 = m.ExcPending
															if v157 != 0 {
																return int32(0)
															} else {
																v161 = v151
																v165 = F_Float8GetDatum(m, v161)
																mBase = m.M
																v166 = m.ExcPending
																if v166 != 0 {
																	return int32(0)
																} else {
																	m.G0 = v13 + int32(128)
																	return v165
																}
															}
														}
													}
												}
											} else {
												if v19 == int32(3552) {
													v135 = float64(0.01)
												} else {
													v135 = float64(0.005)
												}
												v136 = v135
												v137 = float64(0)
												v140 = *(*float64)(unsafe.Add(mBase, uint32(v13)+40))
												v143 = base.F64_add(base.F64_mul(base.F64_sub(base.F64_sub(float64(1), base.F64_promote_f32(v93)), v140), v136), v109)
												if base.F64_lt(v143, v137) != 0 {
													v151 = v137
												} else {
													if base.F64_gt(v143, float64(1)) == int32(0) {
														v151 = v143
													} else {
														v151 = float64(1)
													}
												}
												v152 = *(*int32)(unsafe.Add(mBase, uint32(v13)+104))
												if v152 == int32(0) {
													v161 = v151
													v165 = F_Float8GetDatum(m, v161)
													mBase = m.M
													v166 = m.ExcPending
													if v166 != 0 {
														return int32(0)
													} else {
														m.G0 = v13 + int32(128)
														return v165
													}
												} else {
													v155 = *(*int32)(unsafe.Add(mBase, uint32(v13)+108))
													m.T0[v155].(func(*base.Module, int32))(m, v152)
													mBase = m.M
													v157 = m.ExcPending
													if v157 != 0 {
														return int32(0)
													} else {
														v161 = v151
														v165 = F_Float8GetDatum(m, v161)
														mBase = m.M
														v166 = m.ExcPending
														if v166 != 0 {
															return int32(0)
														} else {
															m.G0 = v13 + int32(128)
															return v165
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
	case 1:
		v43 = int32(1)
		v50 = F_get_restriction_variable(m, v17, v16, v15, v13+int32(96), v13+int32(92), v13+int32(91))
		mBase = m.M
		v51 = m.ExcPending
		if v51 != 0 {
			return int32(0)
		} else {
			if v50 == int32(0) {
				if v19 == int32(3552) {
					v58 = float64(0.01)
				} else {
					v58 = float64(0.005)
				}
				v161 = v58
				v165 = F_Float8GetDatum(m, v161)
				mBase = m.M
				v166 = m.ExcPending
				if v166 != 0 {
					return int32(0)
				} else {
					m.G0 = v13 + int32(128)
					return v165
				}
			} else {
				v59 = *(*int32)(unsafe.Add(mBase, uint32(v13)+92))
				v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
				if v60 != int32(7) {
					v63 = *(*int32)(unsafe.Add(mBase, uint32(v13)+104))
					if v63 != 0 {
						v64 = *(*int32)(unsafe.Add(mBase, uint32(v13)+108))
						m.T0[v64].(func(*base.Module, int32))(m, v63)
						mBase = m.M
						v66 = m.ExcPending
						if v66 != 0 {
							return int32(0)
						} else {
							if v19 == int32(3552) {
								v71 = float64(0.01)
							} else {
								v71 = float64(0.005)
							}
							v161 = v71
							v165 = F_Float8GetDatum(m, v161)
							mBase = m.M
							v166 = m.ExcPending
							if v166 != 0 {
								return int32(0)
							} else {
								m.G0 = v13 + int32(128)
								return v165
							}
						}
					} else {
						if v19 == int32(3552) {
							v71 = float64(0.01)
						} else {
							v71 = float64(0.005)
						}
						v161 = v71
						v165 = F_Float8GetDatum(m, v161)
						mBase = m.M
						v166 = m.ExcPending
						if v166 != 0 {
							return int32(0)
						} else {
							m.G0 = v13 + int32(128)
							return v165
						}
					}
				} else {
					v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+24)))
					if v72 == int32(1) {
						v75 = *(*int32)(unsafe.Add(mBase, uint32(v13)+104))
						if v75 == int32(0) {
							v161 = v7
							v165 = F_Float8GetDatum(m, v161)
							mBase = m.M
							v166 = m.ExcPending
							if v166 != 0 {
								return int32(0)
							} else {
								m.G0 = v13 + int32(128)
								return v165
							}
						} else {
							v78 = *(*int32)(unsafe.Add(mBase, uint32(v13)+108))
							m.T0[v78].(func(*base.Module, int32))(m, v75)
							mBase = m.M
							v80 = m.ExcPending
							if v80 != 0 {
								return int32(0)
							} else {
								v161 = v7
								v165 = F_Float8GetDatum(m, v161)
								mBase = m.M
								v166 = m.ExcPending
								if v166 != 0 {
									return int32(0)
								} else {
									m.G0 = v13 + int32(128)
									return v165
								}
							}
						}
					} else {
						v81 = *(*int32)(unsafe.Add(mBase, uint32(v13)+104))
						if v81 == int32(0) {
							if v19 == int32(3552) {
								v88 = float64(0.01)
							} else {
								v88 = float64(0.005)
							}
							v161 = v88
							v165 = F_Float8GetDatum(m, v161)
							mBase = m.M
							v166 = m.ExcPending
							if v166 != 0 {
								return int32(0)
							} else {
								m.G0 = v13 + int32(128)
								return v165
							}
						} else {
							v89 = *(*int32)(unsafe.Add(mBase, uint32(v59)+20))
							v90 = *(*int32)(unsafe.Add(mBase, uint32(v81)+16))
							v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90)+22)))
							v93 = *(*float32)(unsafe.Add(mBase, uint32(v90+v91)+8))
							v94 = F_get_opcode(m, v19)
							mBase = m.M
							v95 = m.ExcPending
							if v95 != 0 {
								return int32(0)
							} else {
								F_fmgr_info(m, v94, v13+int32(12))
								mBase = m.M
								v99 = m.ExcPending
								if v99 != 0 {
									return int32(0)
								} else {
									v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+91)))
									v109 = F_mcv_selectivity(m, v13+int32(96), v13+int32(12), int32(0), v89, v106, v13+int32(40))
									mBase = m.M
									v110 = m.ExcPending
									if v110 != 0 {
										return int32(0)
									} else {
										v113 = *(*int32)(unsafe.Add(mBase, uint32(v13)+104))
										v117 = F_get_attstatsslot(m, v13+int32(52), v113, int32(2), int32(0), int32(1))
										mBase = m.M
										v118 = m.ExcPending
										if v118 != 0 {
											return int32(0)
										} else {
											if v117 != 0 {
												v119 = *(*int32)(unsafe.Add(mBase, uint32(v13)+64))
												v120 = *(*int32)(unsafe.Add(mBase, uint32(v13)+68))
												v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+91)))
												if v123 != 0 {
													v124 = v43
												} else {
													v124 = int32(0) - v43
												}
												v125 = F_inet_hist_value_sel(m, v119, v120, v89, v124)
												mBase = m.M
												v126 = m.ExcPending
												if v126 != 0 {
													return int32(0)
												} else {
													F_free_attstatsslot(m, v13+int32(52))
													mBase = m.M
													v130 = m.ExcPending
													if v130 != 0 {
														return int32(0)
													} else {
														v136 = v125
														v137 = float64(0)
														v140 = *(*float64)(unsafe.Add(mBase, uint32(v13)+40))
														v143 = base.F64_add(base.F64_mul(base.F64_sub(base.F64_sub(float64(1), base.F64_promote_f32(v93)), v140), v136), v109)
														if base.F64_lt(v143, v137) != 0 {
															v151 = v137
														} else {
															if base.F64_gt(v143, float64(1)) == int32(0) {
																v151 = v143
															} else {
																v151 = float64(1)
															}
														}
														v152 = *(*int32)(unsafe.Add(mBase, uint32(v13)+104))
														if v152 == int32(0) {
															v161 = v151
															v165 = F_Float8GetDatum(m, v161)
															mBase = m.M
															v166 = m.ExcPending
															if v166 != 0 {
																return int32(0)
															} else {
																m.G0 = v13 + int32(128)
																return v165
															}
														} else {
															v155 = *(*int32)(unsafe.Add(mBase, uint32(v13)+108))
															m.T0[v155].(func(*base.Module, int32))(m, v152)
															mBase = m.M
															v157 = m.ExcPending
															if v157 != 0 {
																return int32(0)
															} else {
																v161 = v151
																v165 = F_Float8GetDatum(m, v161)
																mBase = m.M
																v166 = m.ExcPending
																if v166 != 0 {
																	return int32(0)
																} else {
																	m.G0 = v13 + int32(128)
																	return v165
																}
															}
														}
													}
												}
											} else {
												if v19 == int32(3552) {
													v135 = float64(0.01)
												} else {
													v135 = float64(0.005)
												}
												v136 = v135
												v137 = float64(0)
												v140 = *(*float64)(unsafe.Add(mBase, uint32(v13)+40))
												v143 = base.F64_add(base.F64_mul(base.F64_sub(base.F64_sub(float64(1), base.F64_promote_f32(v93)), v140), v136), v109)
												if base.F64_lt(v143, v137) != 0 {
													v151 = v137
												} else {
													if base.F64_gt(v143, float64(1)) == int32(0) {
														v151 = v143
													} else {
														v151 = float64(1)
													}
												}
												v152 = *(*int32)(unsafe.Add(mBase, uint32(v13)+104))
												if v152 == int32(0) {
													v161 = v151
													v165 = F_Float8GetDatum(m, v161)
													mBase = m.M
													v166 = m.ExcPending
													if v166 != 0 {
														return int32(0)
													} else {
														m.G0 = v13 + int32(128)
														return v165
													}
												} else {
													v155 = *(*int32)(unsafe.Add(mBase, uint32(v13)+108))
													m.T0[v155].(func(*base.Module, int32))(m, v152)
													mBase = m.M
													v157 = m.ExcPending
													if v157 != 0 {
														return int32(0)
													} else {
														v161 = v151
														v165 = F_Float8GetDatum(m, v161)
														mBase = m.M
														v166 = m.ExcPending
														if v166 != 0 {
															return int32(0)
														} else {
															m.G0 = v13 + int32(128)
															return v165
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
	case 2:
		v43 = int32(-2)
		v50 = F_get_restriction_variable(m, v17, v16, v15, v13+int32(96), v13+int32(92), v13+int32(91))
		mBase = m.M
		v51 = m.ExcPending
		if v51 != 0 {
			return int32(0)
		} else {
			if v50 == int32(0) {
				if v19 == int32(3552) {
					v58 = float64(0.01)
				} else {
					v58 = float64(0.005)
				}
				v161 = v58
				v165 = F_Float8GetDatum(m, v161)
				mBase = m.M
				v166 = m.ExcPending
				if v166 != 0 {
					return int32(0)
				} else {
					m.G0 = v13 + int32(128)
					return v165
				}
			} else {
				v59 = *(*int32)(unsafe.Add(mBase, uint32(v13)+92))
				v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
				if v60 != int32(7) {
					v63 = *(*int32)(unsafe.Add(mBase, uint32(v13)+104))
					if v63 != 0 {
						v64 = *(*int32)(unsafe.Add(mBase, uint32(v13)+108))
						m.T0[v64].(func(*base.Module, int32))(m, v63)
						mBase = m.M
						v66 = m.ExcPending
						if v66 != 0 {
							return int32(0)
						} else {
							if v19 == int32(3552) {
								v71 = float64(0.01)
							} else {
								v71 = float64(0.005)
							}
							v161 = v71
							v165 = F_Float8GetDatum(m, v161)
							mBase = m.M
							v166 = m.ExcPending
							if v166 != 0 {
								return int32(0)
							} else {
								m.G0 = v13 + int32(128)
								return v165
							}
						}
					} else {
						if v19 == int32(3552) {
							v71 = float64(0.01)
						} else {
							v71 = float64(0.005)
						}
						v161 = v71
						v165 = F_Float8GetDatum(m, v161)
						mBase = m.M
						v166 = m.ExcPending
						if v166 != 0 {
							return int32(0)
						} else {
							m.G0 = v13 + int32(128)
							return v165
						}
					}
				} else {
					v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+24)))
					if v72 == int32(1) {
						v75 = *(*int32)(unsafe.Add(mBase, uint32(v13)+104))
						if v75 == int32(0) {
							v161 = v7
							v165 = F_Float8GetDatum(m, v161)
							mBase = m.M
							v166 = m.ExcPending
							if v166 != 0 {
								return int32(0)
							} else {
								m.G0 = v13 + int32(128)
								return v165
							}
						} else {
							v78 = *(*int32)(unsafe.Add(mBase, uint32(v13)+108))
							m.T0[v78].(func(*base.Module, int32))(m, v75)
							mBase = m.M
							v80 = m.ExcPending
							if v80 != 0 {
								return int32(0)
							} else {
								v161 = v7
								v165 = F_Float8GetDatum(m, v161)
								mBase = m.M
								v166 = m.ExcPending
								if v166 != 0 {
									return int32(0)
								} else {
									m.G0 = v13 + int32(128)
									return v165
								}
							}
						}
					} else {
						v81 = *(*int32)(unsafe.Add(mBase, uint32(v13)+104))
						if v81 == int32(0) {
							if v19 == int32(3552) {
								v88 = float64(0.01)
							} else {
								v88 = float64(0.005)
							}
							v161 = v88
							v165 = F_Float8GetDatum(m, v161)
							mBase = m.M
							v166 = m.ExcPending
							if v166 != 0 {
								return int32(0)
							} else {
								m.G0 = v13 + int32(128)
								return v165
							}
						} else {
							v89 = *(*int32)(unsafe.Add(mBase, uint32(v59)+20))
							v90 = *(*int32)(unsafe.Add(mBase, uint32(v81)+16))
							v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90)+22)))
							v93 = *(*float32)(unsafe.Add(mBase, uint32(v90+v91)+8))
							v94 = F_get_opcode(m, v19)
							mBase = m.M
							v95 = m.ExcPending
							if v95 != 0 {
								return int32(0)
							} else {
								F_fmgr_info(m, v94, v13+int32(12))
								mBase = m.M
								v99 = m.ExcPending
								if v99 != 0 {
									return int32(0)
								} else {
									v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+91)))
									v109 = F_mcv_selectivity(m, v13+int32(96), v13+int32(12), int32(0), v89, v106, v13+int32(40))
									mBase = m.M
									v110 = m.ExcPending
									if v110 != 0 {
										return int32(0)
									} else {
										v113 = *(*int32)(unsafe.Add(mBase, uint32(v13)+104))
										v117 = F_get_attstatsslot(m, v13+int32(52), v113, int32(2), int32(0), int32(1))
										mBase = m.M
										v118 = m.ExcPending
										if v118 != 0 {
											return int32(0)
										} else {
											if v117 != 0 {
												v119 = *(*int32)(unsafe.Add(mBase, uint32(v13)+64))
												v120 = *(*int32)(unsafe.Add(mBase, uint32(v13)+68))
												v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+91)))
												if v123 != 0 {
													v124 = v43
												} else {
													v124 = int32(0) - v43
												}
												v125 = F_inet_hist_value_sel(m, v119, v120, v89, v124)
												mBase = m.M
												v126 = m.ExcPending
												if v126 != 0 {
													return int32(0)
												} else {
													F_free_attstatsslot(m, v13+int32(52))
													mBase = m.M
													v130 = m.ExcPending
													if v130 != 0 {
														return int32(0)
													} else {
														v136 = v125
														v137 = float64(0)
														v140 = *(*float64)(unsafe.Add(mBase, uint32(v13)+40))
														v143 = base.F64_add(base.F64_mul(base.F64_sub(base.F64_sub(float64(1), base.F64_promote_f32(v93)), v140), v136), v109)
														if base.F64_lt(v143, v137) != 0 {
															v151 = v137
														} else {
															if base.F64_gt(v143, float64(1)) == int32(0) {
																v151 = v143
															} else {
																v151 = float64(1)
															}
														}
														v152 = *(*int32)(unsafe.Add(mBase, uint32(v13)+104))
														if v152 == int32(0) {
															v161 = v151
															v165 = F_Float8GetDatum(m, v161)
															mBase = m.M
															v166 = m.ExcPending
															if v166 != 0 {
																return int32(0)
															} else {
																m.G0 = v13 + int32(128)
																return v165
															}
														} else {
															v155 = *(*int32)(unsafe.Add(mBase, uint32(v13)+108))
															m.T0[v155].(func(*base.Module, int32))(m, v152)
															mBase = m.M
															v157 = m.ExcPending
															if v157 != 0 {
																return int32(0)
															} else {
																v161 = v151
																v165 = F_Float8GetDatum(m, v161)
																mBase = m.M
																v166 = m.ExcPending
																if v166 != 0 {
																	return int32(0)
																} else {
																	m.G0 = v13 + int32(128)
																	return v165
																}
															}
														}
													}
												}
											} else {
												if v19 == int32(3552) {
													v135 = float64(0.01)
												} else {
													v135 = float64(0.005)
												}
												v136 = v135
												v137 = float64(0)
												v140 = *(*float64)(unsafe.Add(mBase, uint32(v13)+40))
												v143 = base.F64_add(base.F64_mul(base.F64_sub(base.F64_sub(float64(1), base.F64_promote_f32(v93)), v140), v136), v109)
												if base.F64_lt(v143, v137) != 0 {
													v151 = v137
												} else {
													if base.F64_gt(v143, float64(1)) == int32(0) {
														v151 = v143
													} else {
														v151 = float64(1)
													}
												}
												v152 = *(*int32)(unsafe.Add(mBase, uint32(v13)+104))
												if v152 == int32(0) {
													v161 = v151
													v165 = F_Float8GetDatum(m, v161)
													mBase = m.M
													v166 = m.ExcPending
													if v166 != 0 {
														return int32(0)
													} else {
														m.G0 = v13 + int32(128)
														return v165
													}
												} else {
													v155 = *(*int32)(unsafe.Add(mBase, uint32(v13)+108))
													m.T0[v155].(func(*base.Module, int32))(m, v152)
													mBase = m.M
													v157 = m.ExcPending
													if v157 != 0 {
														return int32(0)
													} else {
														v161 = v151
														v165 = F_Float8GetDatum(m, v161)
														mBase = m.M
														v166 = m.ExcPending
														if v166 != 0 {
															return int32(0)
														} else {
															m.G0 = v13 + int32(128)
															return v165
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
	case 3:
		v43 = int32(-1)
		v50 = F_get_restriction_variable(m, v17, v16, v15, v13+int32(96), v13+int32(92), v13+int32(91))
		mBase = m.M
		v51 = m.ExcPending
		if v51 != 0 {
			return int32(0)
		} else {
			if v50 == int32(0) {
				if v19 == int32(3552) {
					v58 = float64(0.01)
				} else {
					v58 = float64(0.005)
				}
				v161 = v58
				v165 = F_Float8GetDatum(m, v161)
				mBase = m.M
				v166 = m.ExcPending
				if v166 != 0 {
					return int32(0)
				} else {
					m.G0 = v13 + int32(128)
					return v165
				}
			} else {
				v59 = *(*int32)(unsafe.Add(mBase, uint32(v13)+92))
				v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
				if v60 != int32(7) {
					v63 = *(*int32)(unsafe.Add(mBase, uint32(v13)+104))
					if v63 != 0 {
						v64 = *(*int32)(unsafe.Add(mBase, uint32(v13)+108))
						m.T0[v64].(func(*base.Module, int32))(m, v63)
						mBase = m.M
						v66 = m.ExcPending
						if v66 != 0 {
							return int32(0)
						} else {
							if v19 == int32(3552) {
								v71 = float64(0.01)
							} else {
								v71 = float64(0.005)
							}
							v161 = v71
							v165 = F_Float8GetDatum(m, v161)
							mBase = m.M
							v166 = m.ExcPending
							if v166 != 0 {
								return int32(0)
							} else {
								m.G0 = v13 + int32(128)
								return v165
							}
						}
					} else {
						if v19 == int32(3552) {
							v71 = float64(0.01)
						} else {
							v71 = float64(0.005)
						}
						v161 = v71
						v165 = F_Float8GetDatum(m, v161)
						mBase = m.M
						v166 = m.ExcPending
						if v166 != 0 {
							return int32(0)
						} else {
							m.G0 = v13 + int32(128)
							return v165
						}
					}
				} else {
					v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+24)))
					if v72 == int32(1) {
						v75 = *(*int32)(unsafe.Add(mBase, uint32(v13)+104))
						if v75 == int32(0) {
							v161 = v7
							v165 = F_Float8GetDatum(m, v161)
							mBase = m.M
							v166 = m.ExcPending
							if v166 != 0 {
								return int32(0)
							} else {
								m.G0 = v13 + int32(128)
								return v165
							}
						} else {
							v78 = *(*int32)(unsafe.Add(mBase, uint32(v13)+108))
							m.T0[v78].(func(*base.Module, int32))(m, v75)
							mBase = m.M
							v80 = m.ExcPending
							if v80 != 0 {
								return int32(0)
							} else {
								v161 = v7
								v165 = F_Float8GetDatum(m, v161)
								mBase = m.M
								v166 = m.ExcPending
								if v166 != 0 {
									return int32(0)
								} else {
									m.G0 = v13 + int32(128)
									return v165
								}
							}
						}
					} else {
						v81 = *(*int32)(unsafe.Add(mBase, uint32(v13)+104))
						if v81 == int32(0) {
							if v19 == int32(3552) {
								v88 = float64(0.01)
							} else {
								v88 = float64(0.005)
							}
							v161 = v88
							v165 = F_Float8GetDatum(m, v161)
							mBase = m.M
							v166 = m.ExcPending
							if v166 != 0 {
								return int32(0)
							} else {
								m.G0 = v13 + int32(128)
								return v165
							}
						} else {
							v89 = *(*int32)(unsafe.Add(mBase, uint32(v59)+20))
							v90 = *(*int32)(unsafe.Add(mBase, uint32(v81)+16))
							v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90)+22)))
							v93 = *(*float32)(unsafe.Add(mBase, uint32(v90+v91)+8))
							v94 = F_get_opcode(m, v19)
							mBase = m.M
							v95 = m.ExcPending
							if v95 != 0 {
								return int32(0)
							} else {
								F_fmgr_info(m, v94, v13+int32(12))
								mBase = m.M
								v99 = m.ExcPending
								if v99 != 0 {
									return int32(0)
								} else {
									v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+91)))
									v109 = F_mcv_selectivity(m, v13+int32(96), v13+int32(12), int32(0), v89, v106, v13+int32(40))
									mBase = m.M
									v110 = m.ExcPending
									if v110 != 0 {
										return int32(0)
									} else {
										v113 = *(*int32)(unsafe.Add(mBase, uint32(v13)+104))
										v117 = F_get_attstatsslot(m, v13+int32(52), v113, int32(2), int32(0), int32(1))
										mBase = m.M
										v118 = m.ExcPending
										if v118 != 0 {
											return int32(0)
										} else {
											if v117 != 0 {
												v119 = *(*int32)(unsafe.Add(mBase, uint32(v13)+64))
												v120 = *(*int32)(unsafe.Add(mBase, uint32(v13)+68))
												v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+91)))
												if v123 != 0 {
													v124 = v43
												} else {
													v124 = int32(0) - v43
												}
												v125 = F_inet_hist_value_sel(m, v119, v120, v89, v124)
												mBase = m.M
												v126 = m.ExcPending
												if v126 != 0 {
													return int32(0)
												} else {
													F_free_attstatsslot(m, v13+int32(52))
													mBase = m.M
													v130 = m.ExcPending
													if v130 != 0 {
														return int32(0)
													} else {
														v136 = v125
														v137 = float64(0)
														v140 = *(*float64)(unsafe.Add(mBase, uint32(v13)+40))
														v143 = base.F64_add(base.F64_mul(base.F64_sub(base.F64_sub(float64(1), base.F64_promote_f32(v93)), v140), v136), v109)
														if base.F64_lt(v143, v137) != 0 {
															v151 = v137
														} else {
															if base.F64_gt(v143, float64(1)) == int32(0) {
																v151 = v143
															} else {
																v151 = float64(1)
															}
														}
														v152 = *(*int32)(unsafe.Add(mBase, uint32(v13)+104))
														if v152 == int32(0) {
															v161 = v151
															v165 = F_Float8GetDatum(m, v161)
															mBase = m.M
															v166 = m.ExcPending
															if v166 != 0 {
																return int32(0)
															} else {
																m.G0 = v13 + int32(128)
																return v165
															}
														} else {
															v155 = *(*int32)(unsafe.Add(mBase, uint32(v13)+108))
															m.T0[v155].(func(*base.Module, int32))(m, v152)
															mBase = m.M
															v157 = m.ExcPending
															if v157 != 0 {
																return int32(0)
															} else {
																v161 = v151
																v165 = F_Float8GetDatum(m, v161)
																mBase = m.M
																v166 = m.ExcPending
																if v166 != 0 {
																	return int32(0)
																} else {
																	m.G0 = v13 + int32(128)
																	return v165
																}
															}
														}
													}
												}
											} else {
												if v19 == int32(3552) {
													v135 = float64(0.01)
												} else {
													v135 = float64(0.005)
												}
												v136 = v135
												v137 = float64(0)
												v140 = *(*float64)(unsafe.Add(mBase, uint32(v13)+40))
												v143 = base.F64_add(base.F64_mul(base.F64_sub(base.F64_sub(float64(1), base.F64_promote_f32(v93)), v140), v136), v109)
												if base.F64_lt(v143, v137) != 0 {
													v151 = v137
												} else {
													if base.F64_gt(v143, float64(1)) == int32(0) {
														v151 = v143
													} else {
														v151 = float64(1)
													}
												}
												v152 = *(*int32)(unsafe.Add(mBase, uint32(v13)+104))
												if v152 == int32(0) {
													v161 = v151
													v165 = F_Float8GetDatum(m, v161)
													mBase = m.M
													v166 = m.ExcPending
													if v166 != 0 {
														return int32(0)
													} else {
														m.G0 = v13 + int32(128)
														return v165
													}
												} else {
													v155 = *(*int32)(unsafe.Add(mBase, uint32(v13)+108))
													m.T0[v155].(func(*base.Module, int32))(m, v152)
													mBase = m.M
													v157 = m.ExcPending
													if v157 != 0 {
														return int32(0)
													} else {
														v161 = v151
														v165 = F_Float8GetDatum(m, v161)
														mBase = m.M
														v166 = m.ExcPending
														if v166 != 0 {
															return int32(0)
														} else {
															m.G0 = v13 + int32(128)
															return v165
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
	default:
		if v19 == int32(3552) {
			v43 = int32(0)
			v50 = F_get_restriction_variable(m, v17, v16, v15, v13+int32(96), v13+int32(92), v13+int32(91))
			mBase = m.M
			v51 = m.ExcPending
			if v51 != 0 {
				return int32(0)
			} else {
				if v50 == int32(0) {
					if v19 == int32(3552) {
						v58 = float64(0.01)
					} else {
						v58 = float64(0.005)
					}
					v161 = v58
					v165 = F_Float8GetDatum(m, v161)
					mBase = m.M
					v166 = m.ExcPending
					if v166 != 0 {
						return int32(0)
					} else {
						m.G0 = v13 + int32(128)
						return v165
					}
				} else {
					v59 = *(*int32)(unsafe.Add(mBase, uint32(v13)+92))
					v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
					if v60 != int32(7) {
						v63 = *(*int32)(unsafe.Add(mBase, uint32(v13)+104))
						if v63 != 0 {
							v64 = *(*int32)(unsafe.Add(mBase, uint32(v13)+108))
							m.T0[v64].(func(*base.Module, int32))(m, v63)
							mBase = m.M
							v66 = m.ExcPending
							if v66 != 0 {
								return int32(0)
							} else {
								if v19 == int32(3552) {
									v71 = float64(0.01)
								} else {
									v71 = float64(0.005)
								}
								v161 = v71
								v165 = F_Float8GetDatum(m, v161)
								mBase = m.M
								v166 = m.ExcPending
								if v166 != 0 {
									return int32(0)
								} else {
									m.G0 = v13 + int32(128)
									return v165
								}
							}
						} else {
							if v19 == int32(3552) {
								v71 = float64(0.01)
							} else {
								v71 = float64(0.005)
							}
							v161 = v71
							v165 = F_Float8GetDatum(m, v161)
							mBase = m.M
							v166 = m.ExcPending
							if v166 != 0 {
								return int32(0)
							} else {
								m.G0 = v13 + int32(128)
								return v165
							}
						}
					} else {
						v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+24)))
						if v72 == int32(1) {
							v75 = *(*int32)(unsafe.Add(mBase, uint32(v13)+104))
							if v75 == int32(0) {
								v161 = v7
								v165 = F_Float8GetDatum(m, v161)
								mBase = m.M
								v166 = m.ExcPending
								if v166 != 0 {
									return int32(0)
								} else {
									m.G0 = v13 + int32(128)
									return v165
								}
							} else {
								v78 = *(*int32)(unsafe.Add(mBase, uint32(v13)+108))
								m.T0[v78].(func(*base.Module, int32))(m, v75)
								mBase = m.M
								v80 = m.ExcPending
								if v80 != 0 {
									return int32(0)
								} else {
									v161 = v7
									v165 = F_Float8GetDatum(m, v161)
									mBase = m.M
									v166 = m.ExcPending
									if v166 != 0 {
										return int32(0)
									} else {
										m.G0 = v13 + int32(128)
										return v165
									}
								}
							}
						} else {
							v81 = *(*int32)(unsafe.Add(mBase, uint32(v13)+104))
							if v81 == int32(0) {
								if v19 == int32(3552) {
									v88 = float64(0.01)
								} else {
									v88 = float64(0.005)
								}
								v161 = v88
								v165 = F_Float8GetDatum(m, v161)
								mBase = m.M
								v166 = m.ExcPending
								if v166 != 0 {
									return int32(0)
								} else {
									m.G0 = v13 + int32(128)
									return v165
								}
							} else {
								v89 = *(*int32)(unsafe.Add(mBase, uint32(v59)+20))
								v90 = *(*int32)(unsafe.Add(mBase, uint32(v81)+16))
								v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90)+22)))
								v93 = *(*float32)(unsafe.Add(mBase, uint32(v90+v91)+8))
								v94 = F_get_opcode(m, v19)
								mBase = m.M
								v95 = m.ExcPending
								if v95 != 0 {
									return int32(0)
								} else {
									F_fmgr_info(m, v94, v13+int32(12))
									mBase = m.M
									v99 = m.ExcPending
									if v99 != 0 {
										return int32(0)
									} else {
										v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+91)))
										v109 = F_mcv_selectivity(m, v13+int32(96), v13+int32(12), int32(0), v89, v106, v13+int32(40))
										mBase = m.M
										v110 = m.ExcPending
										if v110 != 0 {
											return int32(0)
										} else {
											v113 = *(*int32)(unsafe.Add(mBase, uint32(v13)+104))
											v117 = F_get_attstatsslot(m, v13+int32(52), v113, int32(2), int32(0), int32(1))
											mBase = m.M
											v118 = m.ExcPending
											if v118 != 0 {
												return int32(0)
											} else {
												if v117 != 0 {
													v119 = *(*int32)(unsafe.Add(mBase, uint32(v13)+64))
													v120 = *(*int32)(unsafe.Add(mBase, uint32(v13)+68))
													v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+91)))
													if v123 != 0 {
														v124 = v43
													} else {
														v124 = int32(0) - v43
													}
													v125 = F_inet_hist_value_sel(m, v119, v120, v89, v124)
													mBase = m.M
													v126 = m.ExcPending
													if v126 != 0 {
														return int32(0)
													} else {
														F_free_attstatsslot(m, v13+int32(52))
														mBase = m.M
														v130 = m.ExcPending
														if v130 != 0 {
															return int32(0)
														} else {
															v136 = v125
															v137 = float64(0)
															v140 = *(*float64)(unsafe.Add(mBase, uint32(v13)+40))
															v143 = base.F64_add(base.F64_mul(base.F64_sub(base.F64_sub(float64(1), base.F64_promote_f32(v93)), v140), v136), v109)
															if base.F64_lt(v143, v137) != 0 {
																v151 = v137
															} else {
																if base.F64_gt(v143, float64(1)) == int32(0) {
																	v151 = v143
																} else {
																	v151 = float64(1)
																}
															}
															v152 = *(*int32)(unsafe.Add(mBase, uint32(v13)+104))
															if v152 == int32(0) {
																v161 = v151
																v165 = F_Float8GetDatum(m, v161)
																mBase = m.M
																v166 = m.ExcPending
																if v166 != 0 {
																	return int32(0)
																} else {
																	m.G0 = v13 + int32(128)
																	return v165
																}
															} else {
																v155 = *(*int32)(unsafe.Add(mBase, uint32(v13)+108))
																m.T0[v155].(func(*base.Module, int32))(m, v152)
																mBase = m.M
																v157 = m.ExcPending
																if v157 != 0 {
																	return int32(0)
																} else {
																	v161 = v151
																	v165 = F_Float8GetDatum(m, v161)
																	mBase = m.M
																	v166 = m.ExcPending
																	if v166 != 0 {
																		return int32(0)
																	} else {
																		m.G0 = v13 + int32(128)
																		return v165
																	}
																}
															}
														}
													}
												} else {
													if v19 == int32(3552) {
														v135 = float64(0.01)
													} else {
														v135 = float64(0.005)
													}
													v136 = v135
													v137 = float64(0)
													v140 = *(*float64)(unsafe.Add(mBase, uint32(v13)+40))
													v143 = base.F64_add(base.F64_mul(base.F64_sub(base.F64_sub(float64(1), base.F64_promote_f32(v93)), v140), v136), v109)
													if base.F64_lt(v143, v137) != 0 {
														v151 = v137
													} else {
														if base.F64_gt(v143, float64(1)) == int32(0) {
															v151 = v143
														} else {
															v151 = float64(1)
														}
													}
													v152 = *(*int32)(unsafe.Add(mBase, uint32(v13)+104))
													if v152 == int32(0) {
														v161 = v151
														v165 = F_Float8GetDatum(m, v161)
														mBase = m.M
														v166 = m.ExcPending
														if v166 != 0 {
															return int32(0)
														} else {
															m.G0 = v13 + int32(128)
															return v165
														}
													} else {
														v155 = *(*int32)(unsafe.Add(mBase, uint32(v13)+108))
														m.T0[v155].(func(*base.Module, int32))(m, v152)
														mBase = m.M
														v157 = m.ExcPending
														if v157 != 0 {
															return int32(0)
														} else {
															v161 = v151
															v165 = F_Float8GetDatum(m, v161)
															mBase = m.M
															v166 = m.ExcPending
															if v166 != 0 {
																return int32(0)
															} else {
																m.G0 = v13 + int32(128)
																return v165
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
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v31 = m.ExcPending
			if v31 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v13))) = v19
				F_errmsg_internal(m, int32(9911), v13)
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(491625), int32(870), int32(285532))
					mBase = m.M
					v40 = m.ExcPending
					if v40 != 0 {
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
	v6 = *(*int32)(unsafe.Add(mBase, _consts[1]))
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
func F_nocache_index_getattr(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v118 int32
	_ = v118
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	var v138 int32
	_ = v138
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v232 int32
	_ = v232
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v256 int32
	_ = v256
	var v272 int32
	_ = v272
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v313 int32
	_ = v313
	var v324 int32
	_ = v324
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v331 int32
	_ = v331
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v342 int32
	_ = v342
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v374 int32
	_ = v374
	var v377 int32
	_ = v377
	var v380 int32
	_ = v380
	var v382 int32
	_ = v382
	var v384 int32
	_ = v384
	var v387 int32
	_ = v387
	var v392 int32
	_ = v392
	var v395 int32
	_ = v395
	var v400 int32
	_ = v400
	var v403 int32
	_ = v403
	var v406 int32
	_ = v406
	var v412 int32
	_ = v412
	var v419 int32
	_ = v419
	var v426 int32
	_ = v426
	var v429 int32
	_ = v429
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v436 int32
	_ = v436
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v446 int32
	_ = v446
	var v450 int32
	_ = v450
	var v455 int32
	_ = v455
	var v457 int32
	_ = v457
	var v470 int32
	_ = v470
	var v472 int32
	_ = v472
	v4 = int32(0)
	v15 = m.G0
	v17 = v15 - int32(32)
	m.G0 = v17
	v21 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+6)))
	v23 = base.B2i32(v4 <= v21)
	if v4 <= v21 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v24 = int32(8)
	goto L3
L2:
	;
	v24 = int32(16)
	goto L3
L3:
	;
	v26 = l1 - int32(1)
	if v4 <= v21 {
		v89 = v4
		goto L8
	} else {
		goto L9
	}
L4:
	;
	m.G0 = v17 + int32(32)
	return v472
L5:
	;
	v470 = *(*int32)(unsafe.Add(mBase, uint32(v457)))
	v472 = v470
	goto L4
L6:
	;
	v429 = v419 + v426
	v432 = l2 + v26<<(uint(int32(4))%32)
	v433 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v432)+26)))
	if v433 != int32(1) {
		v472 = v429
		goto L4
	} else {
		goto L86
	}
L7:
	;
	v287 = l0 + v24
	v290 = int32(0)
	v294 = v290
	v296 = int32(1)
	v297 = v290
	v298 = v21
	goto L47
L8:
	;
	v93 = l0 + v24
	v95 = l2 + int32(20)
	v98 = v95 + v26<<(uint(int32(4))%32)
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v98)))
	if int32(0) <= v99 {
		goto L17
	} else {
		goto L18
	}
L9:
	;
	v28 = l0 + int32(8)
	v30 = v26 >> (uint(int32(3)) % 32)
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28+v30))))
	v33 = int32(-1)
	if v32|v33<<(uint(v26&int32(7))%32) != v33 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v283 = v28
	v284 = l0 + v24
	goto L7
L11:
	;
	v40 = int32(0)
	if v30 <= v40 {
		v89 = v28
		goto L8
	} else {
		goto L12
	}
L12:
	;
	v44 = v40
	goto L13
L13:
	;
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44+v28))))
	if v58 != int32(255) {
		goto L10
	} else {
		goto L15
	}
L14:
	;
	v89 = v28
	goto L8
L15:
	;
	v62 = v44 + int32(1)
	if v30 != v62 {
		v44 = v62
		goto L13
	} else {
		goto L16
	}
L16:
	;
	goto L14
L17:
	;
	v102 = v99 + v93
	v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98)+6)))
	if v103 != int32(1) {
		v472 = v102
		goto L4
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	if v21&int32(16384) == int32(0) {
		goto L28
	} else {
		goto L29
	}
L20:
	;
	v106 = int32(*(*int16)(unsafe.Add(mBase, uint32(v98)+4)))
	switch v106&int32(65535) - int32(1) {
	case 0:
		goto L23
	case 1:
		goto L22
	default:
		goto L21
	case 3:
		v457 = v102
		goto L5
	}
L21:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L24
	} else {
		goto L25
	}
L22:
	;
	v112 = int32(*(*int16)(unsafe.Add(mBase, uint32(v102))))
	v472 = v112
	goto L4
L23:
	;
	v111 = int32(*(*int8)(unsafe.Add(mBase, uint32(v102))))
	v472 = v111
	goto L4
L24:
	;
	return int32(0)
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v106
	F_errmsg_internal(m, int32(480197), v17+int32(16))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L24
	} else {
		goto L26
	}
L26:
	;
	F_errfinish(m, int32(324700), int32(70), int32(67479))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L24
	} else {
		goto L27
	}
L27:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = int32(0)
	v176 = int32(1)
	v177 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v177 < int32(2) {
		v204 = v176
		goto L36
	} else {
		goto L37
	}
L29:
	;
	if v26 < int32(0) {
		goto L28
	} else {
		goto L30
	}
L30:
	;
	v138 = int32(0)
	goto L31
L31:
	;
	v154 = int32(*(*int16)(unsafe.Add(mBase, uint32(v95+v138<<(uint(int32(4))%32))+4)))
	if v154 <= int32(0) {
		v283 = v89
		v284 = v93
		goto L7
	} else {
		goto L33
	}
L32:
	;
	goto L28
L33:
	;
	v158 = v138 + int32(1)
	if v158 <= v26 {
		v138 = v158
		goto L31
	} else {
		goto L34
	}
L34:
	;
	goto L32
L35:
	;
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v98)))
	v419 = v272
	v426 = v93
	goto L6
L36:
	;
	if v177 <= v204 {
		goto L35
	} else {
		goto L42
	}
L37:
	;
	v181 = v176
	goto L38
L38:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v95+v181<<(uint(int32(4))%32))))
	if v197 <= int32(0) {
		v204 = v181
		goto L36
	} else {
		goto L40
	}
L39:
	;
	goto L35
L40:
	;
	v201 = v181 + int32(1)
	if v201 != v177 {
		v181 = v201
		goto L38
	} else {
		goto L41
	}
L41:
	;
	goto L39
L42:
	;
	v222 = v204<<(uint(int32(4))%32) + v95 - int32(16)
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v222)))
	v224 = int32(*(*int16)(unsafe.Add(mBase, uint32(v222)+4)))
	v227 = v204
	v232 = v223 + v224
	goto L43
L43:
	;
	v242 = v95 + v227<<(uint(int32(4))%32)
	v243 = int32(*(*int16)(unsafe.Add(mBase, uint32(v242)+4)))
	if v243 <= int32(0) {
		goto L35
	} else {
		goto L45
	}
L44:
	;
	goto L35
L45:
	;
	v246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v242)+12)))
	v248 = int32(1)
	v252 = (v232 + v246 - v248) & (int32(0) - v246)
	*(*int32)(unsafe.Add(mBase, uint32(v242))) = v252
	v256 = v227 + v248
	if v256 != v177 {
		v227 = v256
		v232 = v243 + v252
		goto L43
	} else {
		goto L46
	}
L46:
	;
	goto L44
L47:
	;
	if int32(0) <= base.I32_extend16_s(v298) {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v324 = l2 + int32(20) + v294<<(uint(int32(4))%32)
	if v296&int32(1) != 0 {
		goto L55
	} else {
		goto L56
	}
L50:
	;
	v313 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v283+v294>>(uint(int32(3))%32)))))
	if int32(base.Ui32(v313)>>(uint(v294&int32(7))%32))&int32(1) != 0 {
		goto L49
	} else {
		goto L51
	}
L51:
	;
	v294 = v294 + int32(1)
	v296 = int32(0)
	goto L47
L52:
	;
	if v294 == v26 {
		v419 = v368
		v426 = v284
		goto L6
	} else {
		goto L65
	}
L53:
	;
	v357 = int32(0)
	v359 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v297+v287))))
	if v359 != 0 {
		v368 = v297
		v369 = v357
		goto L52
	} else {
		goto L64
	}
L54:
	;
	v342 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v324)+12)))
	v348 = (v297 + v342 - int32(1)) & (int32(0) - v342)
	v349 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v324)+4)))
	if v349 != int32(65535) {
		goto L60
	} else {
		goto L61
	}
L55:
	;
	v327 = int32(1)
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v324)))
	if v328 < int32(0) {
		goto L54
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	v331 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v324)+4)))
	if v331 == int32(65535) {
		goto L53
	} else {
		goto L59
	}
L58:
	;
	v368 = v328
	v369 = v327
	goto L52
L59:
	;
	v334 = int32(0)
	v335 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v324)+12)))
	v368 = (v297 + v335 - int32(1)) & (v334 - v335)
	v369 = v334
	goto L52
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v324))) = v348
	v368 = v348
	v369 = v327
	goto L52
L61:
	;
	goto L62
L62:
	;
	if v348 != v297 {
		goto L53
	} else {
		goto L63
	}
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v324))) = v297
	v368 = v297
	v369 = v327
	goto L52
L64:
	;
	v360 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v324)+12)))
	v368 = (v297 + v360 - int32(1)) & (int32(0) - v360)
	v369 = v357
	goto L52
L65:
	;
	v371 = int32(*(*int16)(unsafe.Add(mBase, uint32(v324)+4)))
	if int32(0) < v371 {
		v406 = v371
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v412 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+6)))
	v294 = v294 + int32(1)
	v296 = v369 & base.B2i32(int32(0) < v371)
	v297 = v406 + v368
	v298 = v412
	goto L47
L67:
	;
	v374 = v368 + v287
	if v371 == int32(-1) {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v377 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v374))))
	if v377 == int32(1) {
		goto L71
	} else {
		goto L72
	}
L69:
	;
	goto L70
L70:
	;
	v403 = F_strlen(m, v374)
	mBase = m.M
	v406 = v403 + int32(1)
	goto L66
L71:
	;
	v380 = int32(6)
	v382 = int32(18)
	v384 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v374)+1)))
	if v384 == v382 {
		goto L74
	} else {
		goto L75
	}
L72:
	;
	goto L73
L73:
	;
	if v377&int32(1) != 0 {
		goto L83
	} else {
		goto L84
	}
L74:
	;
	v387 = v382
	goto L76
L75:
	;
	v387 = int32(2)
	goto L76
L76:
	;
	if v384&int32(254) == int32(2) {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v392 = v380
	goto L79
L78:
	;
	v392 = v387
	goto L79
L79:
	;
	if v384 == int32(1) {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v395 = v380
	goto L82
L81:
	;
	v395 = v392
	goto L82
L82:
	;
	v406 = v395
	goto L66
L83:
	;
	v406 = int32(base.Ui32(v377) >> (uint(int32(1)) % 32))
	goto L66
L84:
	;
	goto L85
L85:
	;
	v400 = *(*int32)(unsafe.Add(mBase, uint32(v374)))
	v406 = int32(base.Ui32(v400) >> (uint(int32(2)) % 32))
	goto L66
L86:
	;
	v436 = int32(*(*int16)(unsafe.Add(mBase, uint32(v432)+24)))
	switch v436&int32(65535) - int32(1) {
	case 0:
		goto L89
	case 1:
		goto L88
	default:
		goto L87
	case 3:
		v457 = v429
		goto L5
	}
L87:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v446 = m.ExcPending
	if v446 != 0 {
		goto L24
	} else {
		goto L90
	}
L88:
	;
	v442 = int32(*(*int16)(unsafe.Add(mBase, uint32(v429))))
	v472 = v442
	goto L4
L89:
	;
	v441 = int32(*(*int8)(unsafe.Add(mBase, uint32(v429))))
	v472 = v441
	goto L4
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v436
	F_errmsg_internal(m, int32(480197), v17)
	mBase = m.M
	v450 = m.ExcPending
	if v450 != 0 {
		goto L24
	} else {
		goto L91
	}
L91:
	;
	F_errfinish(m, int32(324700), int32(70), int32(67479))
	mBase = m.M
	v455 = m.ExcPending
	if v455 != 0 {
		goto L24
	} else {
		goto L92
	}
L92:
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
	v8 = int32(4394552)
	v9 = int32(*(*uint8)(unsafe.Add(mBase, _consts[478])))
	*(*uint8)(unsafe.Add(mBase, _consts[478])) = uint8(v2)
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
			*(*uint8)(unsafe.Add(mBase, _consts[478])) = uint8(v21)
			v23 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
			m.G0 = v6 + int32(16)
			return v23
		}
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
