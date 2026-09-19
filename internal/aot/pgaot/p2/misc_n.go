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
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v241 int32
	_ = v241
	var v247 int32
	_ = v247
	var v252 int32
	_ = v252
	var v257 int32
	_ = v257
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v284 int32
	_ = v284
	var v290 int32
	_ = v290
	var v295 int32
	_ = v295
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v313 int32
	_ = v313
	var v324 int32
	_ = v324
	var v330 int32
	_ = v330
	var v336 int32
	_ = v336
	var v341 int32
	_ = v341
	var v345 int32
	_ = v345
	var v352 int32
	_ = v352
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v390 int32
	_ = v390
	var v394 int32
	_ = v394
	var v410 int32
	_ = v410
	var v412 int32
	_ = v412
	var v436 int32
	_ = v436
	var v440 int32
	_ = v440
	var v445 int32
	_ = v445
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v456 int32
	_ = v456
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v496 int32
	_ = v496
	var v500 int32
	_ = v500
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v508 int32
	_ = v508
	var v511 int32
	_ = v511
	var v513 int32
	_ = v513
	var v515 int32
	_ = v515
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v538 int32
	_ = v538
	var v545 int32
	_ = v545
	var v549 int32
	_ = v549
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v559 int32
	_ = v559
	var v563 int32
	_ = v563
	var v565 int32
	_ = v565
	var v567 int32
	_ = v567
	var v570 int32
	_ = v570
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v582 int32
	_ = v582
	var v584 int32
	_ = v584
	var v587 int32
	_ = v587
	var v592 int32
	_ = v592
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v601 int32
	_ = v601
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v606 int32
	_ = v606
	var v615 int32
	_ = v615
	var v618 int32
	_ = v618
	var v624 int32
	_ = v624
	var v627 int32
	_ = v627
	var v629 int32
	_ = v629
	var v632 int32
	_ = v632
	var v633 int32
	_ = v633
	var v637 int32
	_ = v637
	var v638 int32
	_ = v638
	var v640 int32
	_ = v640
	var v641 int32
	_ = v641
	var v643 int32
	_ = v643
	var v644 int32
	_ = v644
	var v645 int32
	_ = v645
	var v646 int32
	_ = v646
	var v651 int32
	_ = v651
	var v659 int32
	_ = v659
	var v663 int32
	_ = v663
	var v665 int32
	_ = v665
	var v666 int32
	_ = v666
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	var v673 int32
	_ = v673
	var v677 int32
	_ = v677
	var v679 int32
	_ = v679
	var v681 int32
	_ = v681
	var v684 int32
	_ = v684
	var v689 int32
	_ = v689
	var v690 int32
	_ = v690
	var v691 int32
	_ = v691
	var v693 int32
	_ = v693
	var v694 int32
	_ = v694
	var v696 int32
	_ = v696
	var v698 int32
	_ = v698
	var v701 int32
	_ = v701
	var v706 int32
	_ = v706
	var v707 int32
	_ = v707
	var v708 int32
	_ = v708
	var v715 int32
	_ = v715
	var v717 int32
	_ = v717
	var v718 int32
	_ = v718
	var v720 int32
	_ = v720
	var v728 int32
	_ = v728
	var v729 int32
	_ = v729
	var v730 int32
	_ = v730
	var v735 int32
	_ = v735
	var v737 int32
	_ = v737
	var v742 int32
	_ = v742
	var v744 int32
	_ = v744
	var v747 int32
	_ = v747
	var v748 int32
	_ = v748
	var v752 int32
	_ = v752
	var v753 int32
	_ = v753
	var v755 int32
	_ = v755
	var v756 int32
	_ = v756
	var v758 int32
	_ = v758
	var v759 int32
	_ = v759
	var v760 int32
	_ = v760
	var v761 int32
	_ = v761
	var v768 int32
	_ = v768
	var v774 int32
	_ = v774
	var v778 int32
	_ = v778
	var v780 int32
	_ = v780
	var v781 int32
	_ = v781
	var v785 int32
	_ = v785
	var v786 int32
	_ = v786
	var v788 int32
	_ = v788
	var v792 int32
	_ = v792
	var v794 int32
	_ = v794
	var v796 int32
	_ = v796
	var v799 int32
	_ = v799
	var v804 int32
	_ = v804
	var v805 int32
	_ = v805
	var v806 int32
	_ = v806
	var v808 int32
	_ = v808
	var v809 int32
	_ = v809
	var v811 int32
	_ = v811
	var v813 int32
	_ = v813
	var v816 int32
	_ = v816
	var v821 int32
	_ = v821
	var v822 int32
	_ = v822
	var v823 int32
	_ = v823
	var v830 int32
	_ = v830
	var v832 int32
	_ = v832
	var v833 int32
	_ = v833
	var v835 int32
	_ = v835
	var v847 int32
	_ = v847
	var v849 int32
	_ = v849
	var v858 int32
	_ = v858
	var v860 int32
	_ = v860
	var v864 int32
	_ = v864
	var v867 int32
	_ = v867
	var v871 int32
	_ = v871
	var v876 int32
	_ = v876
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
	v858 = v20 + int32(32)
	F_pg_regerror(m, v477, v858)
	mBase = m.M
	v860 = m.ExcPending
	if v860 != 0 {
		goto L9
	} else {
		goto L254
	}
L13:
	;
	v496 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
	v500 = l2 << (uint(int32(1)) % 32)
	v503 = v496&int32(-255) | v500&int32(254)
	v504 = int32(28)
	if v500&v504 != 0 {
		goto L147
	} else {
		goto L148
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
	v72 = int32(4)
	v73 = l3
	v75 = v61
	goto L22
L22:
	;
	switch v72 - int32(1) {
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
	v125 = F_pg_mblen_cstr(m, v73)
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
	if v72&int32(-2) == int32(2) {
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
	v84 = F_t_isalpha_cstr(m, v73)
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
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73))))
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
	v94 = F_t_isalpha_cstr(m, v73)
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
	v102 = F_t_isalpha_cstr(m, v73)
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
		v124 = v72
		goto L24
	} else {
		goto L45
	}
L45:
	;
	v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73))))
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
	F_errmsg_internal(m, int32(_a_F_NIAddAffix_0), v58)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L9
	} else {
		goto L48
	}
L48:
	;
	F_errfinish(m, int32(_a_F_NIAddAffix_1), int32(66), int32(_a_F_NIAddAffix_2))
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
	v127 = v125 + v73
	v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127))))
	if v128 != 0 {
		v72 = v124
		v73 = v127
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
	v447 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v448 = F_strlen(m, l3)
	mBase = m.M
	v451 = F_MemoryContextAlloc(m, v447, v448+int32(3))
	mBase = m.M
	v452 = m.ExcPending
	if v452 != 0 {
		goto L9
	} else {
		goto L137
	}
L55:
	;
	v164 = l3
	goto L57
L56:
	;
	v164 = int32(_a_F_NIAddAffix_3)
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
	v436 = m.ExcPending
	if v436 != 0 {
		goto L9
	} else {
		goto L134
	}
L60:
	;
	v175 = v169 + int32(9)
	v181 = int32(0)
	v185 = v164
	v186 = v173
	v190 = int32(4)
	goto L63
L61:
	;
	goto L62
L62:
	;
	v385 = *(*int32)(unsafe.Add(mBase, uint32(v159)))
	if v385 != 0 {
		goto L128
	} else {
		goto L129
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
	if v360 != int32(4) {
		goto L59
	} else {
		goto L127
	}
L65:
	;
	v362 = F_pg_mblen_cstr(m, v185)
	mBase = m.M
	v363 = m.ExcPending
	if v363 != 0 {
		goto L9
	} else {
		goto L125
	}
L66:
	;
	if v190&int32(-2) == int32(2) {
		goto L107
	} else {
		goto L108
	}
L67:
	;
	if v186&int32(255) == int32(94) {
		goto L93
	} else {
		goto L94
	}
L68:
	;
	v197 = F_t_isalpha_cstr(m, v185)
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
	v224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v185))))
	if v224 == int32(91) {
		goto L82
	} else {
		goto L83
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
	v209 = F_pg_mblen_cstr(m, v185)
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
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
	if v209 != 0 {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	base.MemoryCopy(m, v199+int32(8), v185, v209)
	goto L81
L80:
	;
	goto L81
L81:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v199)))
	*(*int32)(unsafe.Add(mBase, uint32(v199))) = v214&int32(-262141) | v209<<(uint(int32(2))%32)&int32(_a_F_NIAddAffix_4)
	v358 = v199
	v360 = int32(4)
	goto L65
L82:
	;
	v227 = F_palloc0(m, v175)
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
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
	v241 = m.ExcPending
	if v241 != 0 {
		goto L9
	} else {
		goto L90
	}
L85:
	;
	if v181 != 0 {
		goto L87
	} else {
		goto L88
	}
L86:
	;
	v231 = int32(1)
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v227)))
	*(*int32)(unsafe.Add(mBase, uint32(v227))) = v232&int32(-4) | v231
	v358 = v227
	v360 = v231
	goto L65
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v181)+4)) = v227
	goto L86
L88:
	;
	goto L89
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v159))) = v227
	goto L86
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v167)+48)) = v164
	F_errmsg_internal(m, int32(_a_F_NIAddAffix_5), v167+int32(48))
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L9
	} else {
		goto L91
	}
L91:
	;
	F_errfinish(m, int32(_a_F_NIAddAffix_1), int32(118), int32(_a_F_NIAddAffix_6))
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
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
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v181)))
	*(*int32)(unsafe.Add(mBase, uint32(v181))) = v257&int32(-4) | int32(2)
	v358 = v181
	v360 = int32(3)
	goto L65
L94:
	;
	goto L95
L95:
	;
	v264 = F_t_isalpha_cstr(m, v185)
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L9
	} else {
		goto L96
	}
L96:
	;
	if v264 != 0 {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	v266 = F_pg_mblen_cstr(m, v185)
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
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
	v284 = m.ExcPending
	if v284 != 0 {
		goto L9
	} else {
		goto L104
	}
L100:
	;
	if v266 != 0 {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	base.MemoryCopy(m, v181+int32(8), v185, v266)
	goto L103
L102:
	;
	goto L103
L103:
	;
	v271 = int32(2)
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v181)))
	*(*int32)(unsafe.Add(mBase, uint32(v181))) = v272&int32(-262141) | v266<<(uint(v271)%32)&int32(_a_F_NIAddAffix_4)
	v358 = v181
	v360 = v271
	goto L65
L104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v167)+64)) = v164
	F_errmsg_internal(m, int32(_a_F_NIAddAffix_5), v167-int32(-64))
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L9
	} else {
		goto L105
	}
L105:
	;
	F_errfinish(m, int32(_a_F_NIAddAffix_1), int32(133), int32(_a_F_NIAddAffix_6))
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
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
	v300 = F_t_isalpha_cstr(m, v185)
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
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
	v345 = m.ExcPending
	if v345 != 0 {
		goto L9
	} else {
		goto L122
	}
L110:
	;
	if v300 != 0 {
		goto L111
	} else {
		goto L112
	}
L111:
	;
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v181)))
	v303 = F_pg_mblen_cstr(m, v185)
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L9
	} else {
		goto L114
	}
L112:
	;
	goto L113
L113:
	;
	v324 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v185))))
	if v324 == int32(93) {
		v358 = v181
		v360 = int32(4)
		goto L65
	} else {
		goto L118
	}
L114:
	;
	if v303 != 0 {
		goto L115
	} else {
		goto L116
	}
L115:
	;
	base.MemoryCopy(m, v181+int32(base.Ui32(v302)>>(uint(int32(2))%32))&int32(_a_F_NIAddAffix_7)+int32(8), v185, v303)
	goto L117
L116:
	;
	goto L117
L117:
	;
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v181)))
	*(*int32)(unsafe.Add(mBase, uint32(v181))) = (v313+v303<<(uint(int32(2))%32))&int32(_a_F_NIAddAffix_4) | v313&int32(-262141)
	v358 = v181
	v360 = v190
	goto L65
L118:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L9
	} else {
		goto L119
	}
L119:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v167)+16)) = v164
	F_errmsg_internal(m, int32(_a_F_NIAddAffix_5), v167+int32(16))
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L9
	} else {
		goto L120
	}
L120:
	;
	F_errfinish(m, int32(_a_F_NIAddAffix_1), int32(142), int32(_a_F_NIAddAffix_6))
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v167)+32)) = int32(3)
	F_errmsg_internal(m, int32(_a_F_NIAddAffix_8), v167+int32(32))
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L9
	} else {
		goto L123
	}
L123:
	;
	F_errfinish(m, int32(_a_F_NIAddAffix_1), int32(145), int32(_a_F_NIAddAffix_6))
	mBase = m.M
	v357 = m.ExcPending
	if v357 != 0 {
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
	v364 = v362 + v185
	v365 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v364))))
	if v365 != 0 {
		v181 = v358
		v185 = v364
		v186 = v365
		v190 = v360
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
	v386 = *(*int32)(unsafe.Add(mBase, uint32(v159)+4))
	v390 = v385
	v394 = v386
	goto L131
L129:
	;
	goto L130
L130:
	;
	m.G0 = v167 + int32(80)
	goto L58
L131:
	;
	v410 = (v394+int32(2))&int32(_a_F_NIAddAffix_9) | v394&int32(-131071)
	*(*int32)(unsafe.Add(mBase, uint32(v159)+4)) = v410
	v412 = *(*int32)(unsafe.Add(mBase, uint32(v390)+4))
	if v412 != 0 {
		v390 = v412
		v394 = v410
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
	*(*int32)(unsafe.Add(mBase, uint32(v167))) = v164
	F_errmsg_internal(m, int32(_a_F_NIAddAffix_5), v167)
	mBase = m.M
	v440 = m.ExcPending
	if v440 != 0 {
		goto L9
	} else {
		goto L135
	}
L135:
	;
	F_errfinish(m, int32(_a_F_NIAddAffix_1), int32(150), int32(_a_F_NIAddAffix_6))
	mBase = m.M
	v445 = m.ExcPending
	if v445 != 0 {
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
	v456 = int32(_a_F_NIAddAffix_10)
	goto L140
L139:
	;
	v456 = int32(_a_F_NIAddAffix_11)
	goto L140
L140:
	;
	v459 = F_pg_sprintf(m, v451, v456, v20+int32(16))
	mBase = m.M
	v460 = m.ExcPending
	if v460 != 0 {
		goto L9
	} else {
		goto L141
	}
L141:
	;
	v461 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v462 = F_strlen(m, v451)
	mBase = m.M
	v467 = F_MemoryContextAlloc(m, v461, v462<<(uint(int32(2))%32)+int32(4))
	mBase = m.M
	v468 = m.ExcPending
	if v468 != 0 {
		goto L9
	} else {
		goto L142
	}
L142:
	;
	v469 = F_pg_mb2wchar_with_len(m, v451, v467, v462)
	mBase = m.M
	v470 = m.ExcPending
	if v470 != 0 {
		goto L9
	} else {
		goto L143
	}
L143:
	;
	v472 = F_palloc(m, int32(32))
	mBase = m.M
	v473 = m.ExcPending
	if v473 != 0 {
		goto L9
	} else {
		goto L144
	}
L144:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v45)+16)) = v472
	v477 = F_pg_regcomp(m, v472, v467, v469, int32(19), int32(100))
	mBase = m.M
	v478 = m.ExcPending
	if v478 != 0 {
		goto L9
	} else {
		goto L145
	}
L145:
	;
	if v477 != 0 {
		goto L12
	} else {
		goto L146
	}
L146:
	;
	goto L13
L147:
	;
	v508 = v503
	goto L149
L148:
	;
	v508 = v503 | v504
	goto L149
L149:
	;
	if v500&int32(34) != 0 {
		goto L150
	} else {
		goto L151
	}
L150:
	;
	v511 = v508
	goto L152
L151:
	;
	v511 = v503
	goto L152
L152:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v45)+4)) = v511
	v513 = F_strlen(m, l1)
	mBase = m.M
	v515 = v513 + int32(1)
	if base.Ui32(int32(1025)) <= base.Ui32(v515) {
		goto L154
	} else {
		goto L155
	}
L153:
	;
	if (l1^v538)&int32(3) != 0 {
		goto L166
	} else {
		goto L167
	}
L154:
	;
	v518 = F_palloc0(m, v515)
	mBase = m.M
	v519 = m.ExcPending
	if v519 != 0 {
		goto L9
	} else {
		goto L157
	}
L155:
	;
	goto L156
L156:
	;
	v523 = (v513 + int32(8)) & int32(4088)
	v524 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	if base.Ui32(v523) <= base.Ui32(v524) {
		goto L159
	} else {
		goto L160
	}
L157:
	;
	v538 = v518
	goto L153
L158:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = v531 - v523
	*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = v523 + v532
	v538 = v532
	goto L153
L159:
	;
	v526 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	v531 = v524
	v532 = v526
	goto L158
L160:
	;
	goto L161
L161:
	;
	v527 = int32(_a_F_NIAddAffix_12)
	v529 = F_palloc0(m, v527)
	mBase = m.M
	v530 = m.ExcPending
	if v530 != 0 {
		goto L9
	} else {
		goto L162
	}
L162:
	;
	v531 = v527
	v532 = v529
	goto L158
L163:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v45))) = v538
	v615 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
	v618 = v615&int32(-2) | l6
	*(*int32)(unsafe.Add(mBase, uint32(v45)+4)) = v618
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
	*(*uint8)(unsafe.Add(mBase, uint32(v594))) = uint8(v593)
	if v593&int32(255) == int32(0) {
		goto L164
	} else {
		goto L180
	}
L166:
	;
	v545 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	v592 = l1
	v593 = v545
	v594 = v538
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
	v549 = l1
	v551 = v538
	goto L172
L170:
	;
	v563 = l1
	v565 = v538
	goto L171
L171:
	;
	v567 = *(*int32)(unsafe.Add(mBase, uint32(v563)))
	v570 = int32(-2139062144)
	if (int32(16843008)-v567|v567)&v570 != v570 {
		v592 = v563
		v593 = v567
		v594 = v565
		goto L165
	} else {
		goto L176
	}
L172:
	;
	v552 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v549))))
	*(*uint8)(unsafe.Add(mBase, uint32(v551))) = uint8(v552)
	if v552 == int32(0) {
		goto L164
	} else {
		goto L174
	}
L173:
	;
	v563 = v559
	v565 = v557
	goto L171
L174:
	;
	v556 = int32(1)
	v557 = v551 + v556
	v559 = v549 + v556
	if v559&int32(3) != 0 {
		v549 = v559
		v551 = v557
		goto L172
	} else {
		goto L175
	}
L175:
	;
	goto L173
L176:
	;
	v575 = v563
	v576 = v567
	v577 = v565
	goto L177
L177:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v577))) = v576
	v579 = int32(4)
	v580 = v577 + v579
	v582 = v575 + v579
	v584 = *(*int32)(unsafe.Add(mBase, uint32(v575)+4))
	v587 = int32(-2139062144)
	if (int32(16843008)-v584|v584)&v587 == v587 {
		v575 = v582
		v576 = v584
		v577 = v580
		goto L177
	} else {
		goto L179
	}
L178:
	;
	v592 = v582
	v593 = v584
	v594 = v580
	goto L165
L179:
	;
	goto L178
L180:
	;
	v601 = v592
	v603 = v594
	goto L181
L181:
	;
	v604 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v601)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v603)+1)) = uint8(v604)
	v606 = int32(1)
	if v604 != 0 {
		v601 = v601 + v606
		v603 = v603 + v606
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
	*(*int32)(unsafe.Add(mBase, uint32(v45)+8)) = v730
	v735 = F_strlen(m, l5)
	mBase = m.M
	v737 = v735 & int32(_a_F_NIAddAffix_13)
	*(*int32)(unsafe.Add(mBase, uint32(v45)+4)) = v729&int32(-16776193) | v737<<(uint(int32(10))%32)
	if v737 != 0 {
		goto L220
	} else {
		goto L221
	}
L185:
	;
	v729 = v618
	v730 = int32(_a_F_NIAddAffix_3)
	goto L184
L186:
	;
	goto L187
L187:
	;
	v624 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4))))
	if v624 == int32(0) {
		v729 = v618
		v730 = int32(_a_F_NIAddAffix_3)
		goto L184
	} else {
		goto L188
	}
L188:
	;
	v627 = F_strlen(m, l4)
	mBase = m.M
	v629 = v627 + int32(1)
	if base.Ui32(int32(1025)) <= base.Ui32(v629) {
		goto L190
	} else {
		goto L191
	}
L189:
	;
	if (l4^v651)&int32(3) != 0 {
		goto L202
	} else {
		goto L203
	}
L190:
	;
	v632 = F_palloc0(m, v629)
	mBase = m.M
	v633 = m.ExcPending
	if v633 != 0 {
		goto L9
	} else {
		goto L193
	}
L191:
	;
	goto L192
L192:
	;
	v637 = (v627 + int32(8)) & int32(4088)
	v638 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	if base.Ui32(v637) <= base.Ui32(v638) {
		goto L195
	} else {
		goto L196
	}
L193:
	;
	v651 = v632
	goto L189
L194:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = v645 - v637
	*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = v646 + v637
	v651 = v646
	goto L189
L195:
	;
	v640 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	v645 = v638
	v646 = v640
	goto L194
L196:
	;
	goto L197
L197:
	;
	v641 = int32(_a_F_NIAddAffix_12)
	v643 = F_palloc0(m, v641)
	mBase = m.M
	v644 = m.ExcPending
	if v644 != 0 {
		goto L9
	} else {
		goto L198
	}
L198:
	;
	v645 = v641
	v646 = v643
	goto L194
L199:
	;
	v728 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
	v729 = v728
	v730 = v651
	goto L184
L200:
	;
	goto L199
L201:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v708))) = uint8(v707)
	if v707&int32(255) == int32(0) {
		goto L200
	} else {
		goto L216
	}
L202:
	;
	v659 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4))))
	v706 = l4
	v707 = v659
	v708 = v651
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
	v663 = l4
	v665 = v651
	goto L208
L206:
	;
	v677 = l4
	v679 = v651
	goto L207
L207:
	;
	v681 = *(*int32)(unsafe.Add(mBase, uint32(v677)))
	v684 = int32(-2139062144)
	if (int32(16843008)-v681|v681)&v684 != v684 {
		v706 = v677
		v707 = v681
		v708 = v679
		goto L201
	} else {
		goto L212
	}
L208:
	;
	v666 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v663))))
	*(*uint8)(unsafe.Add(mBase, uint32(v665))) = uint8(v666)
	if v666 == int32(0) {
		goto L200
	} else {
		goto L210
	}
L209:
	;
	v677 = v673
	v679 = v671
	goto L207
L210:
	;
	v670 = int32(1)
	v671 = v665 + v670
	v673 = v663 + v670
	if v673&int32(3) != 0 {
		v663 = v673
		v665 = v671
		goto L208
	} else {
		goto L211
	}
L211:
	;
	goto L209
L212:
	;
	v689 = v677
	v690 = v681
	v691 = v679
	goto L213
L213:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v691))) = v690
	v693 = int32(4)
	v694 = v691 + v693
	v696 = v689 + v693
	v698 = *(*int32)(unsafe.Add(mBase, uint32(v689)+4))
	v701 = int32(-2139062144)
	if (int32(16843008)-v698|v698)&v701 == v701 {
		v689 = v696
		v690 = v698
		v691 = v694
		goto L213
	} else {
		goto L215
	}
L214:
	;
	v706 = v696
	v707 = v698
	v708 = v694
	goto L201
L215:
	;
	goto L214
L216:
	;
	v715 = v706
	v717 = v708
	goto L217
L217:
	;
	v718 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v715)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v717)+1)) = uint8(v718)
	v720 = int32(1)
	if v718 != 0 {
		v715 = v715 + v720
		v717 = v717 + v720
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
	v742 = F_strlen(m, l5)
	mBase = m.M
	v744 = v742 + int32(1)
	if base.Ui32(int32(1025)) <= base.Ui32(v744) {
		goto L224
	} else {
		goto L225
	}
L221:
	;
	v847 = int32(_a_F_NIAddAffix_3)
	goto L222
L222:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v45)+12)) = v847
	v849 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v849 + int32(1)
	m.G0 = v20 + int32(144)
	return
L223:
	;
	if (l5^v768)&int32(3) != 0 {
		goto L236
	} else {
		goto L237
	}
L224:
	;
	v747 = F_palloc0(m, v744)
	mBase = m.M
	v748 = m.ExcPending
	if v748 != 0 {
		goto L9
	} else {
		goto L227
	}
L225:
	;
	goto L226
L226:
	;
	v752 = (v742 + int32(8)) & int32(4088)
	v753 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	if base.Ui32(v752) <= base.Ui32(v753) {
		goto L229
	} else {
		goto L230
	}
L227:
	;
	v768 = v747
	goto L223
L228:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = v760 - v752
	*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = v752 + v761
	v768 = v761
	goto L223
L229:
	;
	v755 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	v760 = v753
	v761 = v755
	goto L228
L230:
	;
	goto L231
L231:
	;
	v756 = int32(_a_F_NIAddAffix_12)
	v758 = F_palloc0(m, v756)
	mBase = m.M
	v759 = m.ExcPending
	if v759 != 0 {
		goto L9
	} else {
		goto L232
	}
L232:
	;
	v760 = v756
	v761 = v758
	goto L228
L233:
	;
	v847 = v768
	goto L222
L234:
	;
	goto L233
L235:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v823))) = uint8(v822)
	if v822&int32(255) == int32(0) {
		goto L234
	} else {
		goto L250
	}
L236:
	;
	v774 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l5))))
	v821 = l5
	v822 = v774
	v823 = v768
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
	v778 = l5
	v780 = v768
	goto L242
L240:
	;
	v792 = l5
	v794 = v768
	goto L241
L241:
	;
	v796 = *(*int32)(unsafe.Add(mBase, uint32(v792)))
	v799 = int32(-2139062144)
	if (int32(16843008)-v796|v796)&v799 != v799 {
		v821 = v792
		v822 = v796
		v823 = v794
		goto L235
	} else {
		goto L246
	}
L242:
	;
	v781 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v778))))
	*(*uint8)(unsafe.Add(mBase, uint32(v780))) = uint8(v781)
	if v781 == int32(0) {
		goto L234
	} else {
		goto L244
	}
L243:
	;
	v792 = v788
	v794 = v786
	goto L241
L244:
	;
	v785 = int32(1)
	v786 = v780 + v785
	v788 = v778 + v785
	if v788&int32(3) != 0 {
		v778 = v788
		v780 = v786
		goto L242
	} else {
		goto L245
	}
L245:
	;
	goto L243
L246:
	;
	v804 = v792
	v805 = v796
	v806 = v794
	goto L247
L247:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v806))) = v805
	v808 = int32(4)
	v809 = v806 + v808
	v811 = v804 + v808
	v813 = *(*int32)(unsafe.Add(mBase, uint32(v804)+4))
	v816 = int32(-2139062144)
	if (int32(16843008)-v813|v813)&v816 == v816 {
		v804 = v811
		v805 = v813
		v806 = v809
		goto L247
	} else {
		goto L249
	}
L248:
	;
	v821 = v811
	v822 = v813
	v823 = v809
	goto L235
L249:
	;
	goto L248
L250:
	;
	v830 = v821
	v832 = v823
	goto L251
L251:
	;
	v833 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v830)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v832)+1)) = uint8(v833)
	v835 = int32(1)
	if v833 != 0 {
		v830 = v830 + v835
		v832 = v832 + v835
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
	v864 = m.ExcPending
	if v864 != 0 {
		goto L9
	} else {
		goto L255
	}
L255:
	;
	F_errcode(m, int32(302252162))
	mBase = m.M
	v867 = m.ExcPending
	if v867 != 0 {
		goto L9
	} else {
		goto L256
	}
L256:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v858
	F_errmsg(m, int32(_a_F_NIAddAffix_14), v20)
	mBase = m.M
	v871 = m.ExcPending
	if v871 != 0 {
		goto L9
	} else {
		goto L257
	}
L257:
	;
	F_errfinish(m, int32(_a_F_NIAddAffix_15), int32(752), int32(_a_F_NIAddAffix_16))
	mBase = m.M
	v876 = m.ExcPending
	if v876 != 0 {
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
	F_LockBuffer(m, v18, int32(1))
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
func F_name_bpchar(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = F_cstring_to_text(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
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
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
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
			v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
			if v26 == int32(18) {
				v29 = int32(16)
			} else {
				v29 = int32(0)
			}
			if base.Ui32((v26-int32(1))&int32(255)) < base.Ui32(int32(3)) {
				v36 = int32(4)
			} else {
				v36 = v29
			}
			v47 = v36
		} else {
			v37 = int32(1)
			if v19 != 0 {
				v47 = int32(base.Ui32(v17)>>(uint(v37)%32)) - v37
			} else {
				v41 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
				v47 = int32(base.Ui32(v41)>>(uint(int32(2))%32)) - int32(4)
			}
		}
		v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v49 = F_GenericMatchText(m, v6, v14, v20, v47, v48)
		mBase = m.M
		v50 = m.ExcPending
		if v50 != 0 {
			return int32(0)
		} else {
			return base.B2i32(v49 == int32(1))
		}
	}
}
func F_nameregexeq(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn13969(m, l0, int32(19))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
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
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v216 int32
	_ = v216
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v246 int32
	_ = v246
	var v251 int32
	_ = v251
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
	v239 = m.ExcPending
	if v239 != 0 {
		goto L16
	} else {
		goto L68
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
	v226 = m.ExcPending
	if v226 != 0 {
		goto L16
	} else {
		goto L65
	}
L5:
	;
	m.G0 = v7 + int32(32)
	return v216
L6:
	;
	v193 = m.G0
	v194 = int32(16)
	v195 = v193 - v194
	m.G0 = v195
	v198 = F_palloc0(m, v194)
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
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
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
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
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v22 = int32(0)
	v25 = F_makeBoolConst(m, base.B2i32(v21 == v22), v22)
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
	v216 = v17
	goto L5
L18:
	;
	v216 = v25
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
	v216 = v33
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
	v216 = v58
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
	v216 = v145
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
	v216 = v82
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
	v216 = v89
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
	v216 = v108
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
	v216 = v114
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
	v216 = v141
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
	v216 = v162
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
	v216 = v177
	goto L5
L63:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v198))) = int64(8589934613)
	*(*int32)(unsafe.Add(mBase, uint32(v195)+8)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v195)+12)) = l0
	v207 = F_list_make1_impl(m, int32(1), v195+int32(8))
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L16
	} else {
		goto L64
	}
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v198)+12)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v198)+8)) = v207
	m.G0 = v195 + int32(16)
	v216 = v198
	goto L5
L65:
	;
	F_errmsg_internal(m, int32(_a_F_negate_clause_3), int32(0))
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L16
	} else {
		goto L66
	}
L66:
	;
	F_errfinish(m, int32(_a_F_negate_clause_1), int32(76), int32(_a_F_negate_clause_2))
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L16
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
	v240 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v240
	F_errmsg_internal(m, int32(_a_F_negate_clause_4), v7+int32(16))
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L16
	} else {
		goto L69
	}
L69:
	;
	F_errfinish(m, int32(_a_F_negate_clause_1), int32(250), int32(_a_F_negate_clause_2))
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L16
	} else {
		goto L70
	}
L70:
	;
	base.Wasm_trap_unreachable()
	for {
	}
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
	if v16&int32(_a_F_neqjoinsel_0) == int32(4) {
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
				v62 = F_DirectFunctionCall5Coll(m, int32(1493), v53, v15, v55, v14, base.I32_extend16_s(v16), v13)
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
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v107 float64
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 float64
	_ = v123
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v131 float64
	_ = v131
	var v132 float64
	_ = v132
	var v133 float64
	_ = v133
	var v136 float64
	_ = v136
	var v139 float64
	_ = v139
	var v147 float64
	_ = v147
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v158 float64
	_ = v158
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
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
				v158 = v58
				v161 = F_Float8GetDatum(m, v158)
				mBase = m.M
				v162 = m.ExcPending
				if v162 != 0 {
					return int32(0)
				} else {
					m.G0 = v13 + int32(128)
					return v161
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
							v158 = v71
							v161 = F_Float8GetDatum(m, v158)
							mBase = m.M
							v162 = m.ExcPending
							if v162 != 0 {
								return int32(0)
							} else {
								m.G0 = v13 + int32(128)
								return v161
							}
						}
					} else {
						if v19 == int32(3552) {
							v71 = float64(0.01)
						} else {
							v71 = float64(0.005)
						}
						v158 = v71
						v161 = F_Float8GetDatum(m, v158)
						mBase = m.M
						v162 = m.ExcPending
						if v162 != 0 {
							return int32(0)
						} else {
							m.G0 = v13 + int32(128)
							return v161
						}
					}
				} else {
					v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+24)))
					if v72 == int32(1) {
						v75 = *(*int32)(unsafe.Add(mBase, uint32(v13)+104))
						if v75 == int32(0) {
							v158 = v7
							v161 = F_Float8GetDatum(m, v158)
							mBase = m.M
							v162 = m.ExcPending
							if v162 != 0 {
								return int32(0)
							} else {
								m.G0 = v13 + int32(128)
								return v161
							}
						} else {
							v78 = *(*int32)(unsafe.Add(mBase, uint32(v13)+108))
							m.T0[v78].(func(*base.Module, int32))(m, v75)
							mBase = m.M
							v80 = m.ExcPending
							if v80 != 0 {
								return int32(0)
							} else {
								v158 = v7
								v161 = F_Float8GetDatum(m, v158)
								mBase = m.M
								v162 = m.ExcPending
								if v162 != 0 {
									return int32(0)
								} else {
									m.G0 = v13 + int32(128)
									return v161
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
							v158 = v88
							v161 = F_Float8GetDatum(m, v158)
							mBase = m.M
							v162 = m.ExcPending
							if v162 != 0 {
								return int32(0)
							} else {
								m.G0 = v13 + int32(128)
								return v161
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
								v97 = v13 + int32(12)
								F_fmgr_info(m, v94, v97)
								mBase = m.M
								v99 = m.ExcPending
								if v99 != 0 {
									return int32(0)
								} else {
									v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+91)))
									v107 = F_mcv_selectivity(m, v13+int32(96), v97, int32(0), v89, v104, v13+int32(40))
									mBase = m.M
									v108 = m.ExcPending
									if v108 != 0 {
										return int32(0)
									} else {
										v110 = v13 + int32(52)
										v111 = *(*int32)(unsafe.Add(mBase, uint32(v13)+104))
										v115 = F_get_attstatsslot(m, v110, v111, int32(2), int32(0), int32(1))
										mBase = m.M
										v116 = m.ExcPending
										if v116 != 0 {
											return int32(0)
										} else {
											if v115 != 0 {
												v117 = *(*int32)(unsafe.Add(mBase, uint32(v13)+64))
												v118 = *(*int32)(unsafe.Add(mBase, uint32(v13)+68))
												v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+91)))
												if v121 != 0 {
													v122 = v43
												} else {
													v122 = int32(0) - v43
												}
												v123 = F_inet_hist_value_sel(m, v117, v118, v89, v122)
												mBase = m.M
												v124 = m.ExcPending
												if v124 != 0 {
													return int32(0)
												} else {
													F_free_attstatsslot(m, v110)
													mBase = m.M
													v126 = m.ExcPending
													if v126 != 0 {
														return int32(0)
													} else {
														v132 = v123
														v133 = float64(0)
														v136 = *(*float64)(unsafe.Add(mBase, uint32(v13)+40))
														v139 = base.F64_add(base.F64_mul(base.F64_sub(base.F64_sub(float64(1), base.F64_promote_f32(v93)), v136), v132), v107)
														if base.F64_lt(v139, v133) != 0 {
															v147 = v133
														} else {
															if base.F64_gt(v139, float64(1)) == int32(0) {
																v147 = v139
															} else {
																v147 = float64(1)
															}
														}
														v148 = *(*int32)(unsafe.Add(mBase, uint32(v13)+104))
														if v148 == int32(0) {
															v158 = v147
															v161 = F_Float8GetDatum(m, v158)
															mBase = m.M
															v162 = m.ExcPending
															if v162 != 0 {
																return int32(0)
															} else {
																m.G0 = v13 + int32(128)
																return v161
															}
														} else {
															v151 = *(*int32)(unsafe.Add(mBase, uint32(v13)+108))
															m.T0[v151].(func(*base.Module, int32))(m, v148)
															mBase = m.M
															v153 = m.ExcPending
															if v153 != 0 {
																return int32(0)
															} else {
																v158 = v147
																v161 = F_Float8GetDatum(m, v158)
																mBase = m.M
																v162 = m.ExcPending
																if v162 != 0 {
																	return int32(0)
																} else {
																	m.G0 = v13 + int32(128)
																	return v161
																}
															}
														}
													}
												}
											} else {
												if v19 == int32(3552) {
													v131 = float64(0.01)
												} else {
													v131 = float64(0.005)
												}
												v132 = v131
												v133 = float64(0)
												v136 = *(*float64)(unsafe.Add(mBase, uint32(v13)+40))
												v139 = base.F64_add(base.F64_mul(base.F64_sub(base.F64_sub(float64(1), base.F64_promote_f32(v93)), v136), v132), v107)
												if base.F64_lt(v139, v133) != 0 {
													v147 = v133
												} else {
													if base.F64_gt(v139, float64(1)) == int32(0) {
														v147 = v139
													} else {
														v147 = float64(1)
													}
												}
												v148 = *(*int32)(unsafe.Add(mBase, uint32(v13)+104))
												if v148 == int32(0) {
													v158 = v147
													v161 = F_Float8GetDatum(m, v158)
													mBase = m.M
													v162 = m.ExcPending
													if v162 != 0 {
														return int32(0)
													} else {
														m.G0 = v13 + int32(128)
														return v161
													}
												} else {
													v151 = *(*int32)(unsafe.Add(mBase, uint32(v13)+108))
													m.T0[v151].(func(*base.Module, int32))(m, v148)
													mBase = m.M
													v153 = m.ExcPending
													if v153 != 0 {
														return int32(0)
													} else {
														v158 = v147
														v161 = F_Float8GetDatum(m, v158)
														mBase = m.M
														v162 = m.ExcPending
														if v162 != 0 {
															return int32(0)
														} else {
															m.G0 = v13 + int32(128)
															return v161
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
				v158 = v58
				v161 = F_Float8GetDatum(m, v158)
				mBase = m.M
				v162 = m.ExcPending
				if v162 != 0 {
					return int32(0)
				} else {
					m.G0 = v13 + int32(128)
					return v161
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
							v158 = v71
							v161 = F_Float8GetDatum(m, v158)
							mBase = m.M
							v162 = m.ExcPending
							if v162 != 0 {
								return int32(0)
							} else {
								m.G0 = v13 + int32(128)
								return v161
							}
						}
					} else {
						if v19 == int32(3552) {
							v71 = float64(0.01)
						} else {
							v71 = float64(0.005)
						}
						v158 = v71
						v161 = F_Float8GetDatum(m, v158)
						mBase = m.M
						v162 = m.ExcPending
						if v162 != 0 {
							return int32(0)
						} else {
							m.G0 = v13 + int32(128)
							return v161
						}
					}
				} else {
					v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+24)))
					if v72 == int32(1) {
						v75 = *(*int32)(unsafe.Add(mBase, uint32(v13)+104))
						if v75 == int32(0) {
							v158 = v7
							v161 = F_Float8GetDatum(m, v158)
							mBase = m.M
							v162 = m.ExcPending
							if v162 != 0 {
								return int32(0)
							} else {
								m.G0 = v13 + int32(128)
								return v161
							}
						} else {
							v78 = *(*int32)(unsafe.Add(mBase, uint32(v13)+108))
							m.T0[v78].(func(*base.Module, int32))(m, v75)
							mBase = m.M
							v80 = m.ExcPending
							if v80 != 0 {
								return int32(0)
							} else {
								v158 = v7
								v161 = F_Float8GetDatum(m, v158)
								mBase = m.M
								v162 = m.ExcPending
								if v162 != 0 {
									return int32(0)
								} else {
									m.G0 = v13 + int32(128)
									return v161
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
							v158 = v88
							v161 = F_Float8GetDatum(m, v158)
							mBase = m.M
							v162 = m.ExcPending
							if v162 != 0 {
								return int32(0)
							} else {
								m.G0 = v13 + int32(128)
								return v161
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
								v97 = v13 + int32(12)
								F_fmgr_info(m, v94, v97)
								mBase = m.M
								v99 = m.ExcPending
								if v99 != 0 {
									return int32(0)
								} else {
									v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+91)))
									v107 = F_mcv_selectivity(m, v13+int32(96), v97, int32(0), v89, v104, v13+int32(40))
									mBase = m.M
									v108 = m.ExcPending
									if v108 != 0 {
										return int32(0)
									} else {
										v110 = v13 + int32(52)
										v111 = *(*int32)(unsafe.Add(mBase, uint32(v13)+104))
										v115 = F_get_attstatsslot(m, v110, v111, int32(2), int32(0), int32(1))
										mBase = m.M
										v116 = m.ExcPending
										if v116 != 0 {
											return int32(0)
										} else {
											if v115 != 0 {
												v117 = *(*int32)(unsafe.Add(mBase, uint32(v13)+64))
												v118 = *(*int32)(unsafe.Add(mBase, uint32(v13)+68))
												v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+91)))
												if v121 != 0 {
													v122 = v43
												} else {
													v122 = int32(0) - v43
												}
												v123 = F_inet_hist_value_sel(m, v117, v118, v89, v122)
												mBase = m.M
												v124 = m.ExcPending
												if v124 != 0 {
													return int32(0)
												} else {
													F_free_attstatsslot(m, v110)
													mBase = m.M
													v126 = m.ExcPending
													if v126 != 0 {
														return int32(0)
													} else {
														v132 = v123
														v133 = float64(0)
														v136 = *(*float64)(unsafe.Add(mBase, uint32(v13)+40))
														v139 = base.F64_add(base.F64_mul(base.F64_sub(base.F64_sub(float64(1), base.F64_promote_f32(v93)), v136), v132), v107)
														if base.F64_lt(v139, v133) != 0 {
															v147 = v133
														} else {
															if base.F64_gt(v139, float64(1)) == int32(0) {
																v147 = v139
															} else {
																v147 = float64(1)
															}
														}
														v148 = *(*int32)(unsafe.Add(mBase, uint32(v13)+104))
														if v148 == int32(0) {
															v158 = v147
															v161 = F_Float8GetDatum(m, v158)
															mBase = m.M
															v162 = m.ExcPending
															if v162 != 0 {
																return int32(0)
															} else {
																m.G0 = v13 + int32(128)
																return v161
															}
														} else {
															v151 = *(*int32)(unsafe.Add(mBase, uint32(v13)+108))
															m.T0[v151].(func(*base.Module, int32))(m, v148)
															mBase = m.M
															v153 = m.ExcPending
															if v153 != 0 {
																return int32(0)
															} else {
																v158 = v147
																v161 = F_Float8GetDatum(m, v158)
																mBase = m.M
																v162 = m.ExcPending
																if v162 != 0 {
																	return int32(0)
																} else {
																	m.G0 = v13 + int32(128)
																	return v161
																}
															}
														}
													}
												}
											} else {
												if v19 == int32(3552) {
													v131 = float64(0.01)
												} else {
													v131 = float64(0.005)
												}
												v132 = v131
												v133 = float64(0)
												v136 = *(*float64)(unsafe.Add(mBase, uint32(v13)+40))
												v139 = base.F64_add(base.F64_mul(base.F64_sub(base.F64_sub(float64(1), base.F64_promote_f32(v93)), v136), v132), v107)
												if base.F64_lt(v139, v133) != 0 {
													v147 = v133
												} else {
													if base.F64_gt(v139, float64(1)) == int32(0) {
														v147 = v139
													} else {
														v147 = float64(1)
													}
												}
												v148 = *(*int32)(unsafe.Add(mBase, uint32(v13)+104))
												if v148 == int32(0) {
													v158 = v147
													v161 = F_Float8GetDatum(m, v158)
													mBase = m.M
													v162 = m.ExcPending
													if v162 != 0 {
														return int32(0)
													} else {
														m.G0 = v13 + int32(128)
														return v161
													}
												} else {
													v151 = *(*int32)(unsafe.Add(mBase, uint32(v13)+108))
													m.T0[v151].(func(*base.Module, int32))(m, v148)
													mBase = m.M
													v153 = m.ExcPending
													if v153 != 0 {
														return int32(0)
													} else {
														v158 = v147
														v161 = F_Float8GetDatum(m, v158)
														mBase = m.M
														v162 = m.ExcPending
														if v162 != 0 {
															return int32(0)
														} else {
															m.G0 = v13 + int32(128)
															return v161
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
				v158 = v58
				v161 = F_Float8GetDatum(m, v158)
				mBase = m.M
				v162 = m.ExcPending
				if v162 != 0 {
					return int32(0)
				} else {
					m.G0 = v13 + int32(128)
					return v161
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
							v158 = v71
							v161 = F_Float8GetDatum(m, v158)
							mBase = m.M
							v162 = m.ExcPending
							if v162 != 0 {
								return int32(0)
							} else {
								m.G0 = v13 + int32(128)
								return v161
							}
						}
					} else {
						if v19 == int32(3552) {
							v71 = float64(0.01)
						} else {
							v71 = float64(0.005)
						}
						v158 = v71
						v161 = F_Float8GetDatum(m, v158)
						mBase = m.M
						v162 = m.ExcPending
						if v162 != 0 {
							return int32(0)
						} else {
							m.G0 = v13 + int32(128)
							return v161
						}
					}
				} else {
					v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+24)))
					if v72 == int32(1) {
						v75 = *(*int32)(unsafe.Add(mBase, uint32(v13)+104))
						if v75 == int32(0) {
							v158 = v7
							v161 = F_Float8GetDatum(m, v158)
							mBase = m.M
							v162 = m.ExcPending
							if v162 != 0 {
								return int32(0)
							} else {
								m.G0 = v13 + int32(128)
								return v161
							}
						} else {
							v78 = *(*int32)(unsafe.Add(mBase, uint32(v13)+108))
							m.T0[v78].(func(*base.Module, int32))(m, v75)
							mBase = m.M
							v80 = m.ExcPending
							if v80 != 0 {
								return int32(0)
							} else {
								v158 = v7
								v161 = F_Float8GetDatum(m, v158)
								mBase = m.M
								v162 = m.ExcPending
								if v162 != 0 {
									return int32(0)
								} else {
									m.G0 = v13 + int32(128)
									return v161
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
							v158 = v88
							v161 = F_Float8GetDatum(m, v158)
							mBase = m.M
							v162 = m.ExcPending
							if v162 != 0 {
								return int32(0)
							} else {
								m.G0 = v13 + int32(128)
								return v161
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
								v97 = v13 + int32(12)
								F_fmgr_info(m, v94, v97)
								mBase = m.M
								v99 = m.ExcPending
								if v99 != 0 {
									return int32(0)
								} else {
									v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+91)))
									v107 = F_mcv_selectivity(m, v13+int32(96), v97, int32(0), v89, v104, v13+int32(40))
									mBase = m.M
									v108 = m.ExcPending
									if v108 != 0 {
										return int32(0)
									} else {
										v110 = v13 + int32(52)
										v111 = *(*int32)(unsafe.Add(mBase, uint32(v13)+104))
										v115 = F_get_attstatsslot(m, v110, v111, int32(2), int32(0), int32(1))
										mBase = m.M
										v116 = m.ExcPending
										if v116 != 0 {
											return int32(0)
										} else {
											if v115 != 0 {
												v117 = *(*int32)(unsafe.Add(mBase, uint32(v13)+64))
												v118 = *(*int32)(unsafe.Add(mBase, uint32(v13)+68))
												v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+91)))
												if v121 != 0 {
													v122 = v43
												} else {
													v122 = int32(0) - v43
												}
												v123 = F_inet_hist_value_sel(m, v117, v118, v89, v122)
												mBase = m.M
												v124 = m.ExcPending
												if v124 != 0 {
													return int32(0)
												} else {
													F_free_attstatsslot(m, v110)
													mBase = m.M
													v126 = m.ExcPending
													if v126 != 0 {
														return int32(0)
													} else {
														v132 = v123
														v133 = float64(0)
														v136 = *(*float64)(unsafe.Add(mBase, uint32(v13)+40))
														v139 = base.F64_add(base.F64_mul(base.F64_sub(base.F64_sub(float64(1), base.F64_promote_f32(v93)), v136), v132), v107)
														if base.F64_lt(v139, v133) != 0 {
															v147 = v133
														} else {
															if base.F64_gt(v139, float64(1)) == int32(0) {
																v147 = v139
															} else {
																v147 = float64(1)
															}
														}
														v148 = *(*int32)(unsafe.Add(mBase, uint32(v13)+104))
														if v148 == int32(0) {
															v158 = v147
															v161 = F_Float8GetDatum(m, v158)
															mBase = m.M
															v162 = m.ExcPending
															if v162 != 0 {
																return int32(0)
															} else {
																m.G0 = v13 + int32(128)
																return v161
															}
														} else {
															v151 = *(*int32)(unsafe.Add(mBase, uint32(v13)+108))
															m.T0[v151].(func(*base.Module, int32))(m, v148)
															mBase = m.M
															v153 = m.ExcPending
															if v153 != 0 {
																return int32(0)
															} else {
																v158 = v147
																v161 = F_Float8GetDatum(m, v158)
																mBase = m.M
																v162 = m.ExcPending
																if v162 != 0 {
																	return int32(0)
																} else {
																	m.G0 = v13 + int32(128)
																	return v161
																}
															}
														}
													}
												}
											} else {
												if v19 == int32(3552) {
													v131 = float64(0.01)
												} else {
													v131 = float64(0.005)
												}
												v132 = v131
												v133 = float64(0)
												v136 = *(*float64)(unsafe.Add(mBase, uint32(v13)+40))
												v139 = base.F64_add(base.F64_mul(base.F64_sub(base.F64_sub(float64(1), base.F64_promote_f32(v93)), v136), v132), v107)
												if base.F64_lt(v139, v133) != 0 {
													v147 = v133
												} else {
													if base.F64_gt(v139, float64(1)) == int32(0) {
														v147 = v139
													} else {
														v147 = float64(1)
													}
												}
												v148 = *(*int32)(unsafe.Add(mBase, uint32(v13)+104))
												if v148 == int32(0) {
													v158 = v147
													v161 = F_Float8GetDatum(m, v158)
													mBase = m.M
													v162 = m.ExcPending
													if v162 != 0 {
														return int32(0)
													} else {
														m.G0 = v13 + int32(128)
														return v161
													}
												} else {
													v151 = *(*int32)(unsafe.Add(mBase, uint32(v13)+108))
													m.T0[v151].(func(*base.Module, int32))(m, v148)
													mBase = m.M
													v153 = m.ExcPending
													if v153 != 0 {
														return int32(0)
													} else {
														v158 = v147
														v161 = F_Float8GetDatum(m, v158)
														mBase = m.M
														v162 = m.ExcPending
														if v162 != 0 {
															return int32(0)
														} else {
															m.G0 = v13 + int32(128)
															return v161
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
				v158 = v58
				v161 = F_Float8GetDatum(m, v158)
				mBase = m.M
				v162 = m.ExcPending
				if v162 != 0 {
					return int32(0)
				} else {
					m.G0 = v13 + int32(128)
					return v161
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
							v158 = v71
							v161 = F_Float8GetDatum(m, v158)
							mBase = m.M
							v162 = m.ExcPending
							if v162 != 0 {
								return int32(0)
							} else {
								m.G0 = v13 + int32(128)
								return v161
							}
						}
					} else {
						if v19 == int32(3552) {
							v71 = float64(0.01)
						} else {
							v71 = float64(0.005)
						}
						v158 = v71
						v161 = F_Float8GetDatum(m, v158)
						mBase = m.M
						v162 = m.ExcPending
						if v162 != 0 {
							return int32(0)
						} else {
							m.G0 = v13 + int32(128)
							return v161
						}
					}
				} else {
					v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+24)))
					if v72 == int32(1) {
						v75 = *(*int32)(unsafe.Add(mBase, uint32(v13)+104))
						if v75 == int32(0) {
							v158 = v7
							v161 = F_Float8GetDatum(m, v158)
							mBase = m.M
							v162 = m.ExcPending
							if v162 != 0 {
								return int32(0)
							} else {
								m.G0 = v13 + int32(128)
								return v161
							}
						} else {
							v78 = *(*int32)(unsafe.Add(mBase, uint32(v13)+108))
							m.T0[v78].(func(*base.Module, int32))(m, v75)
							mBase = m.M
							v80 = m.ExcPending
							if v80 != 0 {
								return int32(0)
							} else {
								v158 = v7
								v161 = F_Float8GetDatum(m, v158)
								mBase = m.M
								v162 = m.ExcPending
								if v162 != 0 {
									return int32(0)
								} else {
									m.G0 = v13 + int32(128)
									return v161
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
							v158 = v88
							v161 = F_Float8GetDatum(m, v158)
							mBase = m.M
							v162 = m.ExcPending
							if v162 != 0 {
								return int32(0)
							} else {
								m.G0 = v13 + int32(128)
								return v161
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
								v97 = v13 + int32(12)
								F_fmgr_info(m, v94, v97)
								mBase = m.M
								v99 = m.ExcPending
								if v99 != 0 {
									return int32(0)
								} else {
									v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+91)))
									v107 = F_mcv_selectivity(m, v13+int32(96), v97, int32(0), v89, v104, v13+int32(40))
									mBase = m.M
									v108 = m.ExcPending
									if v108 != 0 {
										return int32(0)
									} else {
										v110 = v13 + int32(52)
										v111 = *(*int32)(unsafe.Add(mBase, uint32(v13)+104))
										v115 = F_get_attstatsslot(m, v110, v111, int32(2), int32(0), int32(1))
										mBase = m.M
										v116 = m.ExcPending
										if v116 != 0 {
											return int32(0)
										} else {
											if v115 != 0 {
												v117 = *(*int32)(unsafe.Add(mBase, uint32(v13)+64))
												v118 = *(*int32)(unsafe.Add(mBase, uint32(v13)+68))
												v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+91)))
												if v121 != 0 {
													v122 = v43
												} else {
													v122 = int32(0) - v43
												}
												v123 = F_inet_hist_value_sel(m, v117, v118, v89, v122)
												mBase = m.M
												v124 = m.ExcPending
												if v124 != 0 {
													return int32(0)
												} else {
													F_free_attstatsslot(m, v110)
													mBase = m.M
													v126 = m.ExcPending
													if v126 != 0 {
														return int32(0)
													} else {
														v132 = v123
														v133 = float64(0)
														v136 = *(*float64)(unsafe.Add(mBase, uint32(v13)+40))
														v139 = base.F64_add(base.F64_mul(base.F64_sub(base.F64_sub(float64(1), base.F64_promote_f32(v93)), v136), v132), v107)
														if base.F64_lt(v139, v133) != 0 {
															v147 = v133
														} else {
															if base.F64_gt(v139, float64(1)) == int32(0) {
																v147 = v139
															} else {
																v147 = float64(1)
															}
														}
														v148 = *(*int32)(unsafe.Add(mBase, uint32(v13)+104))
														if v148 == int32(0) {
															v158 = v147
															v161 = F_Float8GetDatum(m, v158)
															mBase = m.M
															v162 = m.ExcPending
															if v162 != 0 {
																return int32(0)
															} else {
																m.G0 = v13 + int32(128)
																return v161
															}
														} else {
															v151 = *(*int32)(unsafe.Add(mBase, uint32(v13)+108))
															m.T0[v151].(func(*base.Module, int32))(m, v148)
															mBase = m.M
															v153 = m.ExcPending
															if v153 != 0 {
																return int32(0)
															} else {
																v158 = v147
																v161 = F_Float8GetDatum(m, v158)
																mBase = m.M
																v162 = m.ExcPending
																if v162 != 0 {
																	return int32(0)
																} else {
																	m.G0 = v13 + int32(128)
																	return v161
																}
															}
														}
													}
												}
											} else {
												if v19 == int32(3552) {
													v131 = float64(0.01)
												} else {
													v131 = float64(0.005)
												}
												v132 = v131
												v133 = float64(0)
												v136 = *(*float64)(unsafe.Add(mBase, uint32(v13)+40))
												v139 = base.F64_add(base.F64_mul(base.F64_sub(base.F64_sub(float64(1), base.F64_promote_f32(v93)), v136), v132), v107)
												if base.F64_lt(v139, v133) != 0 {
													v147 = v133
												} else {
													if base.F64_gt(v139, float64(1)) == int32(0) {
														v147 = v139
													} else {
														v147 = float64(1)
													}
												}
												v148 = *(*int32)(unsafe.Add(mBase, uint32(v13)+104))
												if v148 == int32(0) {
													v158 = v147
													v161 = F_Float8GetDatum(m, v158)
													mBase = m.M
													v162 = m.ExcPending
													if v162 != 0 {
														return int32(0)
													} else {
														m.G0 = v13 + int32(128)
														return v161
													}
												} else {
													v151 = *(*int32)(unsafe.Add(mBase, uint32(v13)+108))
													m.T0[v151].(func(*base.Module, int32))(m, v148)
													mBase = m.M
													v153 = m.ExcPending
													if v153 != 0 {
														return int32(0)
													} else {
														v158 = v147
														v161 = F_Float8GetDatum(m, v158)
														mBase = m.M
														v162 = m.ExcPending
														if v162 != 0 {
															return int32(0)
														} else {
															m.G0 = v13 + int32(128)
															return v161
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
					v158 = v58
					v161 = F_Float8GetDatum(m, v158)
					mBase = m.M
					v162 = m.ExcPending
					if v162 != 0 {
						return int32(0)
					} else {
						m.G0 = v13 + int32(128)
						return v161
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
								v158 = v71
								v161 = F_Float8GetDatum(m, v158)
								mBase = m.M
								v162 = m.ExcPending
								if v162 != 0 {
									return int32(0)
								} else {
									m.G0 = v13 + int32(128)
									return v161
								}
							}
						} else {
							if v19 == int32(3552) {
								v71 = float64(0.01)
							} else {
								v71 = float64(0.005)
							}
							v158 = v71
							v161 = F_Float8GetDatum(m, v158)
							mBase = m.M
							v162 = m.ExcPending
							if v162 != 0 {
								return int32(0)
							} else {
								m.G0 = v13 + int32(128)
								return v161
							}
						}
					} else {
						v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+24)))
						if v72 == int32(1) {
							v75 = *(*int32)(unsafe.Add(mBase, uint32(v13)+104))
							if v75 == int32(0) {
								v158 = v7
								v161 = F_Float8GetDatum(m, v158)
								mBase = m.M
								v162 = m.ExcPending
								if v162 != 0 {
									return int32(0)
								} else {
									m.G0 = v13 + int32(128)
									return v161
								}
							} else {
								v78 = *(*int32)(unsafe.Add(mBase, uint32(v13)+108))
								m.T0[v78].(func(*base.Module, int32))(m, v75)
								mBase = m.M
								v80 = m.ExcPending
								if v80 != 0 {
									return int32(0)
								} else {
									v158 = v7
									v161 = F_Float8GetDatum(m, v158)
									mBase = m.M
									v162 = m.ExcPending
									if v162 != 0 {
										return int32(0)
									} else {
										m.G0 = v13 + int32(128)
										return v161
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
								v158 = v88
								v161 = F_Float8GetDatum(m, v158)
								mBase = m.M
								v162 = m.ExcPending
								if v162 != 0 {
									return int32(0)
								} else {
									m.G0 = v13 + int32(128)
									return v161
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
									v97 = v13 + int32(12)
									F_fmgr_info(m, v94, v97)
									mBase = m.M
									v99 = m.ExcPending
									if v99 != 0 {
										return int32(0)
									} else {
										v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+91)))
										v107 = F_mcv_selectivity(m, v13+int32(96), v97, int32(0), v89, v104, v13+int32(40))
										mBase = m.M
										v108 = m.ExcPending
										if v108 != 0 {
											return int32(0)
										} else {
											v110 = v13 + int32(52)
											v111 = *(*int32)(unsafe.Add(mBase, uint32(v13)+104))
											v115 = F_get_attstatsslot(m, v110, v111, int32(2), int32(0), int32(1))
											mBase = m.M
											v116 = m.ExcPending
											if v116 != 0 {
												return int32(0)
											} else {
												if v115 != 0 {
													v117 = *(*int32)(unsafe.Add(mBase, uint32(v13)+64))
													v118 = *(*int32)(unsafe.Add(mBase, uint32(v13)+68))
													v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+91)))
													if v121 != 0 {
														v122 = v43
													} else {
														v122 = int32(0) - v43
													}
													v123 = F_inet_hist_value_sel(m, v117, v118, v89, v122)
													mBase = m.M
													v124 = m.ExcPending
													if v124 != 0 {
														return int32(0)
													} else {
														F_free_attstatsslot(m, v110)
														mBase = m.M
														v126 = m.ExcPending
														if v126 != 0 {
															return int32(0)
														} else {
															v132 = v123
															v133 = float64(0)
															v136 = *(*float64)(unsafe.Add(mBase, uint32(v13)+40))
															v139 = base.F64_add(base.F64_mul(base.F64_sub(base.F64_sub(float64(1), base.F64_promote_f32(v93)), v136), v132), v107)
															if base.F64_lt(v139, v133) != 0 {
																v147 = v133
															} else {
																if base.F64_gt(v139, float64(1)) == int32(0) {
																	v147 = v139
																} else {
																	v147 = float64(1)
																}
															}
															v148 = *(*int32)(unsafe.Add(mBase, uint32(v13)+104))
															if v148 == int32(0) {
																v158 = v147
																v161 = F_Float8GetDatum(m, v158)
																mBase = m.M
																v162 = m.ExcPending
																if v162 != 0 {
																	return int32(0)
																} else {
																	m.G0 = v13 + int32(128)
																	return v161
																}
															} else {
																v151 = *(*int32)(unsafe.Add(mBase, uint32(v13)+108))
																m.T0[v151].(func(*base.Module, int32))(m, v148)
																mBase = m.M
																v153 = m.ExcPending
																if v153 != 0 {
																	return int32(0)
																} else {
																	v158 = v147
																	v161 = F_Float8GetDatum(m, v158)
																	mBase = m.M
																	v162 = m.ExcPending
																	if v162 != 0 {
																		return int32(0)
																	} else {
																		m.G0 = v13 + int32(128)
																		return v161
																	}
																}
															}
														}
													}
												} else {
													if v19 == int32(3552) {
														v131 = float64(0.01)
													} else {
														v131 = float64(0.005)
													}
													v132 = v131
													v133 = float64(0)
													v136 = *(*float64)(unsafe.Add(mBase, uint32(v13)+40))
													v139 = base.F64_add(base.F64_mul(base.F64_sub(base.F64_sub(float64(1), base.F64_promote_f32(v93)), v136), v132), v107)
													if base.F64_lt(v139, v133) != 0 {
														v147 = v133
													} else {
														if base.F64_gt(v139, float64(1)) == int32(0) {
															v147 = v139
														} else {
															v147 = float64(1)
														}
													}
													v148 = *(*int32)(unsafe.Add(mBase, uint32(v13)+104))
													if v148 == int32(0) {
														v158 = v147
														v161 = F_Float8GetDatum(m, v158)
														mBase = m.M
														v162 = m.ExcPending
														if v162 != 0 {
															return int32(0)
														} else {
															m.G0 = v13 + int32(128)
															return v161
														}
													} else {
														v151 = *(*int32)(unsafe.Add(mBase, uint32(v13)+108))
														m.T0[v151].(func(*base.Module, int32))(m, v148)
														mBase = m.M
														v153 = m.ExcPending
														if v153 != 0 {
															return int32(0)
														} else {
															v158 = v147
															v161 = F_Float8GetDatum(m, v158)
															mBase = m.M
															v162 = m.ExcPending
															if v162 != 0 {
																return int32(0)
															} else {
																m.G0 = v13 + int32(128)
																return v161
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
				F_errmsg_internal(m, int32(_a_F_networksel_0), v13)
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_networksel_1), int32(870), int32(_a_F_networksel_2))
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
	var v90 int32
	_ = v90
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
	var v132 int32
	_ = v132
	var v141 int32
	_ = v141
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v239 int32
	_ = v239
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v257 int32
	_ = v257
	var v261 int32
	_ = v261
	var v277 int32
	_ = v277
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v292 int32
	_ = v292
	var v295 int32
	_ = v295
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v318 int32
	_ = v318
	var v329 int32
	_ = v329
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v336 int32
	_ = v336
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v347 int32
	_ = v347
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v376 int32
	_ = v376
	var v379 int32
	_ = v379
	var v382 int32
	_ = v382
	var v386 int32
	_ = v386
	var v388 int32
	_ = v388
	var v391 int32
	_ = v391
	var v398 int32
	_ = v398
	var v403 int32
	_ = v403
	var v406 int32
	_ = v406
	var v409 int32
	_ = v409
	var v415 int32
	_ = v415
	var v422 int32
	_ = v422
	var v428 int32
	_ = v428
	var v432 int32
	_ = v432
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v439 int32
	_ = v439
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v449 int32
	_ = v449
	var v453 int32
	_ = v453
	var v458 int32
	_ = v458
	var v460 int32
	_ = v460
	var v473 int32
	_ = v473
	var v488 int32
	_ = v488
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
		v90 = v4
		goto L8
	} else {
		goto L9
	}
L4:
	;
	m.G0 = v17 + int32(32)
	return v488
L5:
	;
	v473 = *(*int32)(unsafe.Add(mBase, uint32(v460)))
	v488 = v473
	goto L4
L6:
	;
	v432 = v422 + v428
	v435 = l2 + v26<<(uint(int32(4))%32)
	v436 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v435)+26)))
	if v436 != int32(1) {
		v488 = v432
		goto L4
	} else {
		goto L83
	}
L7:
	;
	v292 = l0 + v24
	v295 = int32(0)
	v299 = v295
	v301 = int32(1)
	v302 = v295
	v303 = v21
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
	v288 = l0 + v24
	v289 = v28
	goto L7
L11:
	;
	v40 = int32(0)
	if v30 <= v40 {
		v90 = v28
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
	v90 = v28
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
		v488 = v102
		goto L4
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v132 = int32(0)
	if base.B2i32(v21&int32(_a_F_nocache_index_getattr_0) == v132)|base.B2i32(v26 < v132) == v132 {
		goto L28
	} else {
		goto L29
	}
L20:
	;
	v106 = int32(*(*int16)(unsafe.Add(mBase, uint32(v98)+4)))
	switch v106&int32(_a_F_nocache_index_getattr_1) - int32(1) {
	case 0:
		goto L23
	case 1:
		goto L22
	default:
		goto L21
	case 3:
		v460 = v102
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
	v488 = v112
	goto L4
L23:
	;
	v111 = int32(*(*int8)(unsafe.Add(mBase, uint32(v102))))
	v488 = v111
	goto L4
L24:
	;
	return int32(0)
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v106
	F_errmsg_internal(m, int32(_a_F_nocache_index_getattr_2), v17+int32(16))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L24
	} else {
		goto L26
	}
L26:
	;
	F_errfinish(m, int32(_a_F_nocache_index_getattr_3), int32(70), int32(_a_F_nocache_index_getattr_4))
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
	v141 = int32(0)
	goto L31
L29:
	;
	goto L30
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = int32(0)
	v179 = int32(1)
	v180 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v180 < int32(2) {
		v207 = v179
		goto L36
	} else {
		goto L37
	}
L31:
	;
	v157 = int32(*(*int16)(unsafe.Add(mBase, uint32(v95+v141<<(uint(int32(4))%32))+4)))
	if v157 <= int32(0) {
		v288 = v93
		v289 = v90
		goto L7
	} else {
		goto L33
	}
L32:
	;
	goto L30
L33:
	;
	v161 = v141 + int32(1)
	if v161 <= v26 {
		v141 = v161
		goto L31
	} else {
		goto L34
	}
L34:
	;
	goto L32
L35:
	;
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v98)))
	v422 = v277
	v428 = v93
	goto L6
L36:
	;
	if v180 <= v207 {
		goto L35
	} else {
		goto L42
	}
L37:
	;
	v184 = v179
	goto L38
L38:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v95+v184<<(uint(int32(4))%32))))
	if v200 <= int32(0) {
		v207 = v184
		goto L36
	} else {
		goto L40
	}
L39:
	;
	goto L35
L40:
	;
	v204 = v184 + int32(1)
	if v204 != v180 {
		v184 = v204
		goto L38
	} else {
		goto L41
	}
L41:
	;
	goto L39
L42:
	;
	v223 = v95 + v207<<(uint(int32(4))%32)
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v223-int32(16))))
	v229 = int32(*(*int16)(unsafe.Add(mBase, uint32(v223-int32(12)))))
	v232 = v207
	v239 = v226 + v229
	goto L43
L43:
	;
	v247 = v95 + v232<<(uint(int32(4))%32)
	v248 = int32(*(*int16)(unsafe.Add(mBase, uint32(v247)+4)))
	if v248 <= int32(0) {
		goto L35
	} else {
		goto L45
	}
L44:
	;
	goto L35
L45:
	;
	v251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v247)+12)))
	v253 = int32(1)
	v257 = (v239 + v251 - v253) & (int32(0) - v251)
	*(*int32)(unsafe.Add(mBase, uint32(v247))) = v257
	v261 = v232 + v253
	if v261 != v180 {
		v232 = v261
		v239 = v257 + v248
		goto L43
	} else {
		goto L46
	}
L46:
	;
	goto L44
L47:
	;
	if int32(0) <= base.I32_extend16_s(v303) {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v329 = l2 + int32(20) + v299<<(uint(int32(4))%32)
	if v301&int32(1) != 0 {
		goto L55
	} else {
		goto L56
	}
L50:
	;
	v318 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v289+v299>>(uint(int32(3))%32)))))
	if int32(base.Ui32(v318)>>(uint(v299&int32(7))%32))&int32(1) != 0 {
		goto L49
	} else {
		goto L51
	}
L51:
	;
	v299 = v299 + int32(1)
	v301 = int32(0)
	goto L47
L52:
	;
	if v299 == v26 {
		v422 = v373
		v428 = v288
		goto L6
	} else {
		goto L65
	}
L53:
	;
	v362 = int32(0)
	v364 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v302+v292))))
	if v364 != 0 {
		v373 = v302
		v374 = v362
		goto L52
	} else {
		goto L64
	}
L54:
	;
	v347 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v329)+12)))
	v353 = (v302 + v347 - int32(1)) & (int32(0) - v347)
	v354 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v329)+4)))
	if v354 != int32(_a_F_nocache_index_getattr_1) {
		goto L60
	} else {
		goto L61
	}
L55:
	;
	v332 = int32(1)
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v329)))
	if v333 < int32(0) {
		goto L54
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	v336 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v329)+4)))
	if v336 == int32(_a_F_nocache_index_getattr_1) {
		goto L53
	} else {
		goto L59
	}
L58:
	;
	v373 = v333
	v374 = v332
	goto L52
L59:
	;
	v339 = int32(0)
	v340 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v329)+12)))
	v373 = (v302 + v340 - int32(1)) & (v339 - v340)
	v374 = v339
	goto L52
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v329))) = v353
	v373 = v353
	v374 = v332
	goto L52
L61:
	;
	goto L62
L62:
	;
	if v353 != v302 {
		goto L53
	} else {
		goto L63
	}
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v329))) = v302
	v373 = v302
	v374 = v332
	goto L52
L64:
	;
	v365 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v329)+12)))
	v373 = (v302 + v365 - int32(1)) & (int32(0) - v365)
	v374 = v362
	goto L52
L65:
	;
	v376 = int32(*(*int16)(unsafe.Add(mBase, uint32(v329)+4)))
	if int32(0) < v376 {
		v409 = v376
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v415 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+6)))
	v299 = v299 + int32(1)
	v301 = v374 & base.B2i32(int32(0) < v376)
	v302 = v409 + v373
	v303 = v415
	goto L47
L67:
	;
	v379 = v373 + v292
	if v376 == int32(-1) {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v382 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v379))))
	if v382 == int32(1) {
		goto L71
	} else {
		goto L72
	}
L69:
	;
	goto L70
L70:
	;
	v406 = F_strlen(m, v379)
	mBase = m.M
	v409 = v406 + int32(1)
	goto L66
L71:
	;
	v386 = int32(18)
	v388 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v379)+1)))
	if v388 == v386 {
		goto L74
	} else {
		goto L75
	}
L72:
	;
	goto L73
L73:
	;
	if v382&int32(1) != 0 {
		goto L80
	} else {
		goto L81
	}
L74:
	;
	v391 = v386
	goto L76
L75:
	;
	v391 = int32(2)
	goto L76
L76:
	;
	if base.Ui32((v388-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v398 = int32(6)
	goto L79
L78:
	;
	v398 = v391
	goto L79
L79:
	;
	v409 = v398
	goto L66
L80:
	;
	v409 = int32(base.Ui32(v382) >> (uint(int32(1)) % 32))
	goto L66
L81:
	;
	goto L82
L82:
	;
	v403 = *(*int32)(unsafe.Add(mBase, uint32(v379)))
	v409 = int32(base.Ui32(v403) >> (uint(int32(2)) % 32))
	goto L66
L83:
	;
	v439 = int32(*(*int16)(unsafe.Add(mBase, uint32(v435)+24)))
	switch v439&int32(_a_F_nocache_index_getattr_1) - int32(1) {
	case 0:
		goto L86
	case 1:
		goto L85
	default:
		goto L84
	case 3:
		v460 = v432
		goto L5
	}
L84:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v449 = m.ExcPending
	if v449 != 0 {
		goto L24
	} else {
		goto L87
	}
L85:
	;
	v445 = int32(*(*int16)(unsafe.Add(mBase, uint32(v432))))
	v488 = v445
	goto L4
L86:
	;
	v444 = int32(*(*int8)(unsafe.Add(mBase, uint32(v432))))
	v488 = v444
	goto L4
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v439
	F_errmsg_internal(m, int32(_a_F_nocache_index_getattr_2), v17)
	mBase = m.M
	v453 = m.ExcPending
	if v453 != 0 {
		goto L24
	} else {
		goto L88
	}
L88:
	;
	F_errfinish(m, int32(_a_F_nocache_index_getattr_3), int32(70), int32(_a_F_nocache_index_getattr_4))
	mBase = m.M
	v458 = m.ExcPending
	if v458 != 0 {
		goto L24
	} else {
		goto L89
	}
L89:
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
