package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_free_stmt(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v165 int32
	_ = v165
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v237 int32
	_ = v237
	var v243 int32
	_ = v243
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v261 int32
	_ = v261
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v273 int32
	_ = v273
	var v278 int32
	_ = v278
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v290 int32
	_ = v290
	var v294 int32
	_ = v294
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v302 int32
	_ = v302
	var v308 int32
	_ = v308
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v337 int32
	_ = v337
	var v340 int32
	_ = v340
	var v345 int32
	_ = v345
	var v351 int32
	_ = v351
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v362 int32
	_ = v362
	var v365 int32
	_ = v365
	var v369 int32
	_ = v369
	var v373 int32
	_ = v373
	var v376 int32
	_ = v376
	var v381 int32
	_ = v381
	var v387 int32
	_ = v387
	var v391 int32
	_ = v391
	var v393 int32
	_ = v393
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v398 int32
	_ = v398
	var v401 int32
	_ = v401
	var v405 int32
	_ = v405
	var v409 int32
	_ = v409
	var v412 int32
	_ = v412
	var v416 int32
	_ = v416
	var v420 int32
	_ = v420
	var v423 int32
	_ = v423
	var v427 int32
	_ = v427
	var v431 int32
	_ = v431
	var v434 int32
	_ = v434
	var v439 int32
	_ = v439
	var v445 int32
	_ = v445
	var v449 int32
	_ = v449
	var v451 int32
	_ = v451
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v456 int32
	_ = v456
	var v459 int32
	_ = v459
	var v463 int32
	_ = v463
	var v469 int32
	_ = v469
	var v473 int32
	_ = v473
	var v475 int32
	_ = v475
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v491 int32
	_ = v491
	var v495 int32
	_ = v495
	var v501 int32
	_ = v501
	var v505 int32
	_ = v505
	var v507 int32
	_ = v507
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v523 int32
	_ = v523
	var v527 int32
	_ = v527
	var v531 int32
	_ = v531
	var v534 int32
	_ = v534
	var v539 int32
	_ = v539
	var v545 int32
	_ = v545
	var v549 int32
	_ = v549
	var v551 int32
	_ = v551
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v556 int32
	_ = v556
	var v559 int32
	_ = v559
	var v563 int32
	_ = v563
	var v566 int32
	_ = v566
	var v569 int32
	_ = v569
	var v573 int32
	_ = v573
	var v576 int32
	_ = v576
	var v579 int32
	_ = v579
	var v583 int32
	_ = v583
	var v586 int32
	_ = v586
	var v589 int32
	_ = v589
	var v593 int32
	_ = v593
	var v597 int32
	_ = v597
	var v600 int32
	_ = v600
	var v604 int32
	_ = v604
	var v608 int32
	_ = v608
	var v611 int32
	_ = v611
	var v616 int32
	_ = v616
	var v622 int32
	_ = v622
	var v626 int32
	_ = v626
	var v629 int32
	_ = v629
	var v633 int32
	_ = v633
	var v638 int32
	_ = v638
	var v639 int32
	_ = v639
	var v641 int32
	_ = v641
	var v644 int32
	_ = v644
	var v648 int32
	_ = v648
	var v654 int32
	_ = v654
	var v658 int32
	_ = v658
	var v661 int32
	_ = v661
	var v665 int32
	_ = v665
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	var v680 int32
	_ = v680
	var v683 int32
	_ = v683
	var v688 int32
	_ = v688
	var v694 int32
	_ = v694
	var v698 int32
	_ = v698
	var v699 int32
	_ = v699
	var v702 int32
	_ = v702
	var v706 int32
	_ = v706
	var v711 int32
	_ = v711
	var v712 int32
	_ = v712
	var v714 int32
	_ = v714
	var v717 int32
	_ = v717
	var v721 int32
	_ = v721
	var v725 int32
	_ = v725
	var v728 int32
	_ = v728
	var v732 int32
	_ = v732
	var v735 int32
	_ = v735
	var v738 int32
	_ = v738
	var v742 int32
	_ = v742
	var v745 int32
	_ = v745
	var v748 int32
	_ = v748
	var v752 int32
	_ = v752
	var v756 int32
	_ = v756
	var v759 int32
	_ = v759
	var v764 int32
	_ = v764
	var v770 int32
	_ = v770
	var v774 int32
	_ = v774
	var v777 int32
	_ = v777
	var v781 int32
	_ = v781
	var v786 int32
	_ = v786
	var v787 int32
	_ = v787
	var v789 int32
	_ = v789
	var v792 int32
	_ = v792
	var v796 int32
	_ = v796
	var v802 int32
	_ = v802
	var v806 int32
	_ = v806
	var v808 int32
	_ = v808
	var v810 int32
	_ = v810
	var v811 int32
	_ = v811
	var v820 int32
	_ = v820
	var v821 int32
	_ = v821
	var v824 int32
	_ = v824
	var v828 int32
	_ = v828
	var v832 int32
	_ = v832
	var v835 int32
	_ = v835
	var v839 int32
	_ = v839
	var v843 int32
	_ = v843
	var v846 int32
	_ = v846
	var v850 int32
	_ = v850
	var v854 int32
	_ = v854
	var v857 int32
	_ = v857
	var v862 int32
	_ = v862
	var v868 int32
	_ = v868
	var v872 int32
	_ = v872
	var v875 int32
	_ = v875
	var v879 int32
	_ = v879
	var v884 int32
	_ = v884
	var v885 int32
	_ = v885
	var v887 int32
	_ = v887
	var v890 int32
	_ = v890
	var v894 int32
	_ = v894
	var v897 int32
	_ = v897
	var v900 int32
	_ = v900
	var v904 int32
	_ = v904
	var v907 int32
	_ = v907
	var v910 int32
	_ = v910
	var v914 int32
	_ = v914
	var v920 int32
	_ = v920
	var v921 int32
	_ = v921
	var v925 int32
	_ = v925
	var v930 int32
	_ = v930
	var v931 int32
	_ = v931
	var v932 int32
	_ = v932
	var v936 int32
	_ = v936
	var v940 int32
	_ = v940
	var v943 int32
	_ = v943
	var v948 int32
	_ = v948
	var v954 int32
	_ = v954
	var v958 int32
	_ = v958
	var v961 int32
	_ = v961
	var v965 int32
	_ = v965
	var v970 int32
	_ = v970
	var v971 int32
	_ = v971
	var v973 int32
	_ = v973
	var v977 int32
	_ = v977
	var v980 int32
	_ = v980
	var v984 int32
	_ = v984
	var v987 int32
	_ = v987
	var v991 int32
	_ = v991
	var v997 int32
	_ = v997
	var v1001 int32
	_ = v1001
	var v1003 int32
	_ = v1003
	var v1005 int32
	_ = v1005
	var v1006 int32
	_ = v1006
	v2 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	if l0 == v2 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v10 + int32(16)
	return
L2:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch v14 {
	case 0:
		goto L31
	case 1:
		goto L30
	case 2:
		goto L29
	case 3:
		goto L28
	case 4:
		goto L8
	case 5:
		goto L27
	case 6:
		goto L26
	case 7:
		goto L25
	case 8:
		goto L24
	case 9:
		goto L23
	case 10:
		goto L22
	case 11:
		goto L21
	case 12:
		goto L20
	case 13:
		goto L19
	case 14:
		goto L18
	case 15:
		goto L17
	case 16:
		goto L16
	case 17:
		goto L15
	case 18:
		goto L14
	case 19, 22, 25, 26:
		goto L1
	case 20:
		goto L13
	case 21:
		goto L12
	case 23:
		goto L11
	case 24:
		goto L10
	default:
		goto L9
	}
L3:
	;
	v987 = *(*int32)(unsafe.Add(mBase, uint32(v931)+4))
	if v987 <= int32(0) {
		goto L1
	} else {
		goto L296
	}
L4:
	;
	v980 = *(*int32)(unsafe.Add(mBase, uint32(v487)+24))
	if v980 == int32(0) {
		goto L1
	} else {
		goto L294
	}
L5:
	;
	v973 = *(*int32)(unsafe.Add(mBase, uint32(v519)+24))
	if v973 == int32(0) {
		goto L1
	} else {
		goto L292
	}
L6:
	;
	v940 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v940 == int32(0) {
		goto L1
	} else {
		goto L283
	}
L7:
	;
	v932 = *(*int32)(unsafe.Add(mBase, uint32(v820)+24))
	if v932 == int32(0) {
		goto L6
	} else {
		goto L281
	}
L8:
	;
	v931 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v931 != 0 {
		goto L3
	} else {
		goto L280
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_free_stmt_0))
	mBase = m.M
	v920 = m.ExcPending
	if v920 != 0 {
		goto L37
	} else {
		goto L277
	}
L10:
	;
	v907 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v907 == int32(0) {
		goto L1
	} else {
		goto L274
	}
L11:
	;
	v897 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v897 == int32(0) {
		goto L1
	} else {
		goto L271
	}
L12:
	;
	v887 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v887 == int32(0) {
		goto L1
	} else {
		goto L268
	}
L13:
	;
	v821 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v821 == int32(0) {
		goto L247
	} else {
		goto L248
	}
L14:
	;
	v789 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v789 == int32(0) {
		goto L239
	} else {
		goto L240
	}
L15:
	;
	v745 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v745 == int32(0) {
		goto L226
	} else {
		goto L227
	}
L16:
	;
	v735 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v735 == int32(0) {
		goto L1
	} else {
		goto L223
	}
L17:
	;
	v714 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v714 == int32(0) {
		goto L216
	} else {
		goto L217
	}
L18:
	;
	v641 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v641 == int32(0) {
		goto L197
	} else {
		goto L198
	}
L19:
	;
	v586 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v586 == int32(0) {
		goto L180
	} else {
		goto L181
	}
L20:
	;
	v576 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v576 == int32(0) {
		goto L1
	} else {
		goto L177
	}
L21:
	;
	v566 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v566 == int32(0) {
		goto L1
	} else {
		goto L174
	}
L22:
	;
	v556 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v556 == int32(0) {
		goto L1
	} else {
		goto L171
	}
L23:
	;
	v520 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v520 == int32(0) {
		goto L161
	} else {
		goto L162
	}
L24:
	;
	v488 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v488 == int32(0) {
		goto L153
	} else {
		goto L154
	}
L25:
	;
	v456 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v456 == int32(0) {
		goto L145
	} else {
		goto L146
	}
L26:
	;
	v398 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v398 == int32(0) {
		goto L127
	} else {
		goto L128
	}
L27:
	;
	v362 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v362 == int32(0) {
		goto L117
	} else {
		goto L118
	}
L28:
	;
	v254 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v254 == int32(0) {
		goto L90
	} else {
		goto L91
	}
L29:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v114 == int32(0) {
		goto L56
	} else {
		goto L57
	}
L30:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v104 == int32(0) {
		goto L1
	} else {
		goto L53
	}
L31:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v15 == int32(0) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v46 == int32(0) {
		goto L1
	} else {
		goto L40
	}
L33:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	if v18 <= int32(0) {
		goto L32
	} else {
		goto L34
	}
L34:
	;
	v22 = v2
	goto L35
L35:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v28+v22<<(uint(int32(2))%32))))
	F_free_stmt(m, v32)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L37
	} else {
		goto L38
	}
L36:
	;
	goto L32
L37:
	;
	return
L38:
	;
	v36 = v22 + int32(1)
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	if v36 < v37 {
		v22 = v36
		goto L35
	} else {
		goto L39
	}
L39:
	;
	goto L36
L40:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v46)+8))
	if v49 == int32(0) {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v49)+4))
	if v52 <= int32(0) {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	v56 = int32(0)
	goto L43
L43:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v49)+12))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v63+v56<<(uint(int32(2))%32))))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)+8))
	if v68 == int32(0) {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	goto L1
L45:
	;
	v101 = v56 + int32(1)
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v49)+4))
	if v101 < v102 {
		v56 = v101
		goto L43
	} else {
		goto L52
	}
L46:
	;
	v71 = int32(0)
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v68)+4))
	if v72 <= v71 {
		goto L45
	} else {
		goto L47
	}
L47:
	;
	v76 = v71
	goto L48
L48:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v68)+12))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v82+v76<<(uint(int32(2))%32))))
	F_free_stmt(m, v86)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L37
	} else {
		goto L50
	}
L49:
	;
	goto L45
L50:
	;
	v90 = v76 + int32(1)
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v68)+4))
	if v90 < v91 {
		v76 = v90
		goto L48
	} else {
		goto L51
	}
L51:
	;
	goto L49
L52:
	;
	goto L44
L53:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v104)+24))
	if v107 == int32(0) {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	F_SPI_freeplan(m, v107)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L37
	} else {
		goto L55
	}
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v104)+24)) = int32(0)
	goto L1
L56:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v125 == int32(0) {
		goto L60
	} else {
		goto L61
	}
L57:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v114)+24))
	if v117 == int32(0) {
		goto L56
	} else {
		goto L58
	}
L58:
	;
	F_SPI_freeplan(m, v117)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L37
	} else {
		goto L59
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v114)+24)) = int32(0)
	goto L56
L60:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v157 == int32(0) {
		goto L67
	} else {
		goto L68
	}
L61:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v125)+4))
	if v128 <= int32(0) {
		goto L60
	} else {
		goto L62
	}
L62:
	;
	v133 = int32(0)
	goto L63
L63:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v125)+12))
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v139+v133<<(uint(int32(2))%32))))
	F_free_stmt(m, v143)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L37
	} else {
		goto L65
	}
L64:
	;
	goto L60
L65:
	;
	v147 = v133 + int32(1)
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v125)+4))
	if v147 < v148 {
		v133 = v147
		goto L63
	} else {
		goto L66
	}
L66:
	;
	goto L64
L67:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v229 == int32(0) {
		goto L1
	} else {
		goto L84
	}
L68:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v157)+4))
	if v160 <= int32(0) {
		goto L67
	} else {
		goto L69
	}
L69:
	;
	v165 = v2
	goto L70
L70:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v157)+12))
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v170+v165<<(uint(int32(2))%32))))
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v174)+4))
	if v175 == int32(0) {
		goto L72
	} else {
		goto L73
	}
L71:
	;
	goto L67
L72:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v174)+8))
	if v186 == int32(0) {
		goto L76
	} else {
		goto L77
	}
L73:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v175)+24))
	if v178 == int32(0) {
		goto L72
	} else {
		goto L74
	}
L74:
	;
	F_SPI_freeplan(m, v178)
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L37
	} else {
		goto L75
	}
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v175)+24)) = int32(0)
	goto L72
L76:
	;
	v219 = v165 + int32(1)
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v157)+4))
	if v219 < v220 {
		v165 = v219
		goto L70
	} else {
		goto L83
	}
L77:
	;
	v189 = int32(0)
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v186)+4))
	if v190 <= v189 {
		goto L76
	} else {
		goto L78
	}
L78:
	;
	v194 = v189
	goto L79
L79:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v186)+12))
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v200+v194<<(uint(int32(2))%32))))
	F_free_stmt(m, v204)
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L37
	} else {
		goto L81
	}
L80:
	;
	goto L76
L81:
	;
	v208 = v194 + int32(1)
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v186)+4))
	if v208 < v209 {
		v194 = v208
		goto L79
	} else {
		goto L82
	}
L82:
	;
	goto L80
L83:
	;
	goto L71
L84:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v229)+4))
	if v232 <= int32(0) {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	v237 = int32(0)
	goto L86
L86:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v229)+12))
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v243+v237<<(uint(int32(2))%32))))
	F_free_stmt(m, v247)
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L37
	} else {
		goto L88
	}
L87:
	;
	goto L1
L88:
	;
	v251 = v237 + int32(1)
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v229)+4))
	if v251 < v252 {
		v237 = v251
		goto L86
	} else {
		goto L89
	}
L89:
	;
	goto L87
L90:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v265 == int32(0) {
		goto L94
	} else {
		goto L95
	}
L91:
	;
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v254)+24))
	if v257 == int32(0) {
		goto L90
	} else {
		goto L92
	}
L92:
	;
	F_SPI_freeplan(m, v257)
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L37
	} else {
		goto L93
	}
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v254)+24)) = int32(0)
	goto L90
L94:
	;
	v337 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v337 == int32(0) {
		goto L1
	} else {
		goto L111
	}
L95:
	;
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v265)+4))
	if v268 <= int32(0) {
		goto L94
	} else {
		goto L96
	}
L96:
	;
	v273 = v2
	goto L97
L97:
	;
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v265)+12))
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v278+v273<<(uint(int32(2))%32))))
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v282)+4))
	if v283 == int32(0) {
		goto L99
	} else {
		goto L100
	}
L98:
	;
	goto L94
L99:
	;
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v282)+8))
	if v294 == int32(0) {
		goto L103
	} else {
		goto L104
	}
L100:
	;
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v283)+24))
	if v286 == int32(0) {
		goto L99
	} else {
		goto L101
	}
L101:
	;
	F_SPI_freeplan(m, v286)
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L37
	} else {
		goto L102
	}
L102:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v283)+24)) = int32(0)
	goto L99
L103:
	;
	v327 = v273 + int32(1)
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v265)+4))
	if v327 < v328 {
		v273 = v327
		goto L97
	} else {
		goto L110
	}
L104:
	;
	v297 = int32(0)
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v294)+4))
	if v298 <= v297 {
		goto L103
	} else {
		goto L105
	}
L105:
	;
	v302 = v297
	goto L106
L106:
	;
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v294)+12))
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v308+v302<<(uint(int32(2))%32))))
	F_free_stmt(m, v312)
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L37
	} else {
		goto L108
	}
L107:
	;
	goto L103
L108:
	;
	v316 = v302 + int32(1)
	v317 = *(*int32)(unsafe.Add(mBase, uint32(v294)+4))
	if v316 < v317 {
		v302 = v316
		goto L106
	} else {
		goto L109
	}
L109:
	;
	goto L107
L110:
	;
	goto L98
L111:
	;
	v340 = *(*int32)(unsafe.Add(mBase, uint32(v337)+4))
	if v340 <= int32(0) {
		goto L1
	} else {
		goto L112
	}
L112:
	;
	v345 = int32(0)
	goto L113
L113:
	;
	v351 = *(*int32)(unsafe.Add(mBase, uint32(v337)+12))
	v355 = *(*int32)(unsafe.Add(mBase, uint32(v351+v345<<(uint(int32(2))%32))))
	F_free_stmt(m, v355)
	mBase = m.M
	v357 = m.ExcPending
	if v357 != 0 {
		goto L37
	} else {
		goto L115
	}
L114:
	;
	goto L1
L115:
	;
	v359 = v345 + int32(1)
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v337)+4))
	if v359 < v360 {
		v345 = v359
		goto L113
	} else {
		goto L116
	}
L116:
	;
	goto L114
L117:
	;
	v373 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v373 == int32(0) {
		goto L1
	} else {
		goto L121
	}
L118:
	;
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v362)+24))
	if v365 == int32(0) {
		goto L117
	} else {
		goto L119
	}
L119:
	;
	F_SPI_freeplan(m, v365)
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L37
	} else {
		goto L120
	}
L120:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v362)+24)) = int32(0)
	goto L117
L121:
	;
	v376 = *(*int32)(unsafe.Add(mBase, uint32(v373)+4))
	if v376 <= int32(0) {
		goto L1
	} else {
		goto L122
	}
L122:
	;
	v381 = int32(0)
	goto L123
L123:
	;
	v387 = *(*int32)(unsafe.Add(mBase, uint32(v373)+12))
	v391 = *(*int32)(unsafe.Add(mBase, uint32(v387+v381<<(uint(int32(2))%32))))
	F_free_stmt(m, v391)
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L37
	} else {
		goto L125
	}
L124:
	;
	goto L1
L125:
	;
	v395 = v381 + int32(1)
	v396 = *(*int32)(unsafe.Add(mBase, uint32(v373)+4))
	if v395 < v396 {
		v381 = v395
		goto L123
	} else {
		goto L126
	}
L126:
	;
	goto L124
L127:
	;
	v409 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v409 == int32(0) {
		goto L131
	} else {
		goto L132
	}
L128:
	;
	v401 = *(*int32)(unsafe.Add(mBase, uint32(v398)+24))
	if v401 == int32(0) {
		goto L127
	} else {
		goto L129
	}
L129:
	;
	F_SPI_freeplan(m, v401)
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
		goto L37
	} else {
		goto L130
	}
L130:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v398)+24)) = int32(0)
	goto L127
L131:
	;
	v420 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v420 == int32(0) {
		goto L135
	} else {
		goto L136
	}
L132:
	;
	v412 = *(*int32)(unsafe.Add(mBase, uint32(v409)+24))
	if v412 == int32(0) {
		goto L131
	} else {
		goto L133
	}
L133:
	;
	F_SPI_freeplan(m, v412)
	mBase = m.M
	v416 = m.ExcPending
	if v416 != 0 {
		goto L37
	} else {
		goto L134
	}
L134:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v409)+24)) = int32(0)
	goto L131
L135:
	;
	v431 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v431 == int32(0) {
		goto L1
	} else {
		goto L139
	}
L136:
	;
	v423 = *(*int32)(unsafe.Add(mBase, uint32(v420)+24))
	if v423 == int32(0) {
		goto L135
	} else {
		goto L137
	}
L137:
	;
	F_SPI_freeplan(m, v423)
	mBase = m.M
	v427 = m.ExcPending
	if v427 != 0 {
		goto L37
	} else {
		goto L138
	}
L138:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v420)+24)) = int32(0)
	goto L135
L139:
	;
	v434 = *(*int32)(unsafe.Add(mBase, uint32(v431)+4))
	if v434 <= int32(0) {
		goto L1
	} else {
		goto L140
	}
L140:
	;
	v439 = int32(0)
	goto L141
L141:
	;
	v445 = *(*int32)(unsafe.Add(mBase, uint32(v431)+12))
	v449 = *(*int32)(unsafe.Add(mBase, uint32(v445+v439<<(uint(int32(2))%32))))
	F_free_stmt(m, v449)
	mBase = m.M
	v451 = m.ExcPending
	if v451 != 0 {
		goto L37
	} else {
		goto L143
	}
L142:
	;
	goto L1
L143:
	;
	v453 = v439 + int32(1)
	v454 = *(*int32)(unsafe.Add(mBase, uint32(v431)+4))
	if v453 < v454 {
		v439 = v453
		goto L141
	} else {
		goto L144
	}
L144:
	;
	goto L142
L145:
	;
	v487 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v487 != 0 {
		goto L4
	} else {
		goto L152
	}
L146:
	;
	v459 = *(*int32)(unsafe.Add(mBase, uint32(v456)+4))
	if v459 <= int32(0) {
		goto L145
	} else {
		goto L147
	}
L147:
	;
	v463 = v2
	goto L148
L148:
	;
	v469 = *(*int32)(unsafe.Add(mBase, uint32(v456)+12))
	v473 = *(*int32)(unsafe.Add(mBase, uint32(v469+v463<<(uint(int32(2))%32))))
	F_free_stmt(m, v473)
	mBase = m.M
	v475 = m.ExcPending
	if v475 != 0 {
		goto L37
	} else {
		goto L150
	}
L149:
	;
	goto L145
L150:
	;
	v477 = v463 + int32(1)
	v478 = *(*int32)(unsafe.Add(mBase, uint32(v456)+4))
	if v477 < v478 {
		v463 = v477
		goto L148
	} else {
		goto L151
	}
L151:
	;
	goto L149
L152:
	;
	goto L1
L153:
	;
	v519 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v519 != 0 {
		goto L5
	} else {
		goto L160
	}
L154:
	;
	v491 = *(*int32)(unsafe.Add(mBase, uint32(v488)+4))
	if v491 <= int32(0) {
		goto L153
	} else {
		goto L155
	}
L155:
	;
	v495 = v2
	goto L156
L156:
	;
	v501 = *(*int32)(unsafe.Add(mBase, uint32(v488)+12))
	v505 = *(*int32)(unsafe.Add(mBase, uint32(v501+v495<<(uint(int32(2))%32))))
	F_free_stmt(m, v505)
	mBase = m.M
	v507 = m.ExcPending
	if v507 != 0 {
		goto L37
	} else {
		goto L158
	}
L157:
	;
	goto L153
L158:
	;
	v509 = v495 + int32(1)
	v510 = *(*int32)(unsafe.Add(mBase, uint32(v488)+4))
	if v509 < v510 {
		v495 = v509
		goto L156
	} else {
		goto L159
	}
L159:
	;
	goto L157
L160:
	;
	goto L1
L161:
	;
	v531 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v531 == int32(0) {
		goto L1
	} else {
		goto L165
	}
L162:
	;
	v523 = *(*int32)(unsafe.Add(mBase, uint32(v520)+24))
	if v523 == int32(0) {
		goto L161
	} else {
		goto L163
	}
L163:
	;
	F_SPI_freeplan(m, v523)
	mBase = m.M
	v527 = m.ExcPending
	if v527 != 0 {
		goto L37
	} else {
		goto L164
	}
L164:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v520)+24)) = int32(0)
	goto L161
L165:
	;
	v534 = *(*int32)(unsafe.Add(mBase, uint32(v531)+4))
	if v534 <= int32(0) {
		goto L1
	} else {
		goto L166
	}
L166:
	;
	v539 = int32(0)
	goto L167
L167:
	;
	v545 = *(*int32)(unsafe.Add(mBase, uint32(v531)+12))
	v549 = *(*int32)(unsafe.Add(mBase, uint32(v545+v539<<(uint(int32(2))%32))))
	F_free_stmt(m, v549)
	mBase = m.M
	v551 = m.ExcPending
	if v551 != 0 {
		goto L37
	} else {
		goto L169
	}
L168:
	;
	goto L1
L169:
	;
	v553 = v539 + int32(1)
	v554 = *(*int32)(unsafe.Add(mBase, uint32(v531)+4))
	if v553 < v554 {
		v539 = v553
		goto L167
	} else {
		goto L170
	}
L170:
	;
	goto L168
L171:
	;
	v559 = *(*int32)(unsafe.Add(mBase, uint32(v556)+24))
	if v559 == int32(0) {
		goto L1
	} else {
		goto L172
	}
L172:
	;
	F_SPI_freeplan(m, v559)
	mBase = m.M
	v563 = m.ExcPending
	if v563 != 0 {
		goto L37
	} else {
		goto L173
	}
L173:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v556)+24)) = int32(0)
	goto L1
L174:
	;
	v569 = *(*int32)(unsafe.Add(mBase, uint32(v566)+24))
	if v569 == int32(0) {
		goto L1
	} else {
		goto L175
	}
L175:
	;
	F_SPI_freeplan(m, v569)
	mBase = m.M
	v573 = m.ExcPending
	if v573 != 0 {
		goto L37
	} else {
		goto L176
	}
L176:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v566)+24)) = int32(0)
	goto L1
L177:
	;
	v579 = *(*int32)(unsafe.Add(mBase, uint32(v576)+24))
	if v579 == int32(0) {
		goto L1
	} else {
		goto L178
	}
L178:
	;
	F_SPI_freeplan(m, v579)
	mBase = m.M
	v583 = m.ExcPending
	if v583 != 0 {
		goto L37
	} else {
		goto L179
	}
L179:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v576)+24)) = int32(0)
	goto L1
L180:
	;
	v597 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v597 == int32(0) {
		goto L184
	} else {
		goto L185
	}
L181:
	;
	v589 = *(*int32)(unsafe.Add(mBase, uint32(v586)+24))
	if v589 == int32(0) {
		goto L180
	} else {
		goto L182
	}
L182:
	;
	F_SPI_freeplan(m, v589)
	mBase = m.M
	v593 = m.ExcPending
	if v593 != 0 {
		goto L37
	} else {
		goto L183
	}
L183:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v586)+24)) = int32(0)
	goto L180
L184:
	;
	v608 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v608 == int32(0) {
		goto L1
	} else {
		goto L188
	}
L185:
	;
	v600 = *(*int32)(unsafe.Add(mBase, uint32(v597)+24))
	if v600 == int32(0) {
		goto L184
	} else {
		goto L186
	}
L186:
	;
	F_SPI_freeplan(m, v600)
	mBase = m.M
	v604 = m.ExcPending
	if v604 != 0 {
		goto L37
	} else {
		goto L187
	}
L187:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v597)+24)) = int32(0)
	goto L184
L188:
	;
	v611 = *(*int32)(unsafe.Add(mBase, uint32(v608)+4))
	if v611 <= int32(0) {
		goto L1
	} else {
		goto L189
	}
L189:
	;
	v616 = int32(0)
	goto L190
L190:
	;
	v622 = *(*int32)(unsafe.Add(mBase, uint32(v608)+12))
	v626 = *(*int32)(unsafe.Add(mBase, uint32(v622+v616<<(uint(int32(2))%32))))
	if v626 == int32(0) {
		goto L192
	} else {
		goto L193
	}
L191:
	;
	goto L1
L192:
	;
	v638 = v616 + int32(1)
	v639 = *(*int32)(unsafe.Add(mBase, uint32(v608)+4))
	if v638 < v639 {
		v616 = v638
		goto L190
	} else {
		goto L196
	}
L193:
	;
	v629 = *(*int32)(unsafe.Add(mBase, uint32(v626)+24))
	if v629 == int32(0) {
		goto L192
	} else {
		goto L194
	}
L194:
	;
	F_SPI_freeplan(m, v629)
	mBase = m.M
	v633 = m.ExcPending
	if v633 != 0 {
		goto L37
	} else {
		goto L195
	}
L195:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v626)+24)) = int32(0)
	goto L192
L196:
	;
	goto L191
L197:
	;
	v680 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v680 == int32(0) {
		goto L1
	} else {
		goto L207
	}
L198:
	;
	v644 = *(*int32)(unsafe.Add(mBase, uint32(v641)+4))
	if v644 <= int32(0) {
		goto L197
	} else {
		goto L199
	}
L199:
	;
	v648 = v2
	goto L200
L200:
	;
	v654 = *(*int32)(unsafe.Add(mBase, uint32(v641)+12))
	v658 = *(*int32)(unsafe.Add(mBase, uint32(v654+v648<<(uint(int32(2))%32))))
	if v658 == int32(0) {
		goto L202
	} else {
		goto L203
	}
L201:
	;
	goto L197
L202:
	;
	v670 = v648 + int32(1)
	v671 = *(*int32)(unsafe.Add(mBase, uint32(v641)+4))
	if v670 < v671 {
		v648 = v670
		goto L200
	} else {
		goto L206
	}
L203:
	;
	v661 = *(*int32)(unsafe.Add(mBase, uint32(v658)+24))
	if v661 == int32(0) {
		goto L202
	} else {
		goto L204
	}
L204:
	;
	F_SPI_freeplan(m, v661)
	mBase = m.M
	v665 = m.ExcPending
	if v665 != 0 {
		goto L37
	} else {
		goto L205
	}
L205:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v658)+24)) = int32(0)
	goto L202
L206:
	;
	goto L201
L207:
	;
	v683 = *(*int32)(unsafe.Add(mBase, uint32(v680)+4))
	if v683 <= int32(0) {
		goto L1
	} else {
		goto L208
	}
L208:
	;
	v688 = int32(0)
	goto L209
L209:
	;
	v694 = *(*int32)(unsafe.Add(mBase, uint32(v680)+12))
	v698 = *(*int32)(unsafe.Add(mBase, uint32(v694+v688<<(uint(int32(2))%32))))
	v699 = *(*int32)(unsafe.Add(mBase, uint32(v698)+4))
	if v699 == int32(0) {
		goto L211
	} else {
		goto L212
	}
L210:
	;
	goto L1
L211:
	;
	v711 = v688 + int32(1)
	v712 = *(*int32)(unsafe.Add(mBase, uint32(v680)+4))
	if v711 < v712 {
		v688 = v711
		goto L209
	} else {
		goto L215
	}
L212:
	;
	v702 = *(*int32)(unsafe.Add(mBase, uint32(v699)+24))
	if v702 == int32(0) {
		goto L211
	} else {
		goto L213
	}
L213:
	;
	F_SPI_freeplan(m, v702)
	mBase = m.M
	v706 = m.ExcPending
	if v706 != 0 {
		goto L37
	} else {
		goto L214
	}
L214:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v699)+24)) = int32(0)
	goto L211
L215:
	;
	goto L210
L216:
	;
	v725 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v725 == int32(0) {
		goto L1
	} else {
		goto L220
	}
L217:
	;
	v717 = *(*int32)(unsafe.Add(mBase, uint32(v714)+24))
	if v717 == int32(0) {
		goto L216
	} else {
		goto L218
	}
L218:
	;
	F_SPI_freeplan(m, v717)
	mBase = m.M
	v721 = m.ExcPending
	if v721 != 0 {
		goto L37
	} else {
		goto L219
	}
L219:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v714)+24)) = int32(0)
	goto L216
L220:
	;
	v728 = *(*int32)(unsafe.Add(mBase, uint32(v725)+24))
	if v728 == int32(0) {
		goto L1
	} else {
		goto L221
	}
L221:
	;
	F_SPI_freeplan(m, v728)
	mBase = m.M
	v732 = m.ExcPending
	if v732 != 0 {
		goto L37
	} else {
		goto L222
	}
L222:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v725)+24)) = int32(0)
	goto L1
L223:
	;
	v738 = *(*int32)(unsafe.Add(mBase, uint32(v735)+24))
	if v738 == int32(0) {
		goto L1
	} else {
		goto L224
	}
L224:
	;
	F_SPI_freeplan(m, v738)
	mBase = m.M
	v742 = m.ExcPending
	if v742 != 0 {
		goto L37
	} else {
		goto L225
	}
L225:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v735)+24)) = int32(0)
	goto L1
L226:
	;
	v756 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v756 == int32(0) {
		goto L1
	} else {
		goto L230
	}
L227:
	;
	v748 = *(*int32)(unsafe.Add(mBase, uint32(v745)+24))
	if v748 == int32(0) {
		goto L226
	} else {
		goto L228
	}
L228:
	;
	F_SPI_freeplan(m, v748)
	mBase = m.M
	v752 = m.ExcPending
	if v752 != 0 {
		goto L37
	} else {
		goto L229
	}
L229:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v745)+24)) = int32(0)
	goto L226
L230:
	;
	v759 = *(*int32)(unsafe.Add(mBase, uint32(v756)+4))
	if v759 <= int32(0) {
		goto L1
	} else {
		goto L231
	}
L231:
	;
	v764 = int32(0)
	goto L232
L232:
	;
	v770 = *(*int32)(unsafe.Add(mBase, uint32(v756)+12))
	v774 = *(*int32)(unsafe.Add(mBase, uint32(v770+v764<<(uint(int32(2))%32))))
	if v774 == int32(0) {
		goto L234
	} else {
		goto L235
	}
L233:
	;
	goto L1
L234:
	;
	v786 = v764 + int32(1)
	v787 = *(*int32)(unsafe.Add(mBase, uint32(v756)+4))
	if v786 < v787 {
		v764 = v786
		goto L232
	} else {
		goto L238
	}
L235:
	;
	v777 = *(*int32)(unsafe.Add(mBase, uint32(v774)+24))
	if v777 == int32(0) {
		goto L234
	} else {
		goto L236
	}
L236:
	;
	F_SPI_freeplan(m, v777)
	mBase = m.M
	v781 = m.ExcPending
	if v781 != 0 {
		goto L37
	} else {
		goto L237
	}
L237:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v774)+24)) = int32(0)
	goto L234
L238:
	;
	goto L233
L239:
	;
	v820 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v820 != 0 {
		goto L7
	} else {
		goto L246
	}
L240:
	;
	v792 = *(*int32)(unsafe.Add(mBase, uint32(v789)+4))
	if v792 <= int32(0) {
		goto L239
	} else {
		goto L241
	}
L241:
	;
	v796 = v2
	goto L242
L242:
	;
	v802 = *(*int32)(unsafe.Add(mBase, uint32(v789)+12))
	v806 = *(*int32)(unsafe.Add(mBase, uint32(v802+v796<<(uint(int32(2))%32))))
	F_free_stmt(m, v806)
	mBase = m.M
	v808 = m.ExcPending
	if v808 != 0 {
		goto L37
	} else {
		goto L244
	}
L243:
	;
	goto L239
L244:
	;
	v810 = v796 + int32(1)
	v811 = *(*int32)(unsafe.Add(mBase, uint32(v789)+4))
	if v810 < v811 {
		v796 = v810
		goto L242
	} else {
		goto L245
	}
L245:
	;
	goto L243
L246:
	;
	goto L6
L247:
	;
	v832 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v832 == int32(0) {
		goto L251
	} else {
		goto L252
	}
L248:
	;
	v824 = *(*int32)(unsafe.Add(mBase, uint32(v821)+24))
	if v824 == int32(0) {
		goto L247
	} else {
		goto L249
	}
L249:
	;
	F_SPI_freeplan(m, v824)
	mBase = m.M
	v828 = m.ExcPending
	if v828 != 0 {
		goto L37
	} else {
		goto L250
	}
L250:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v821)+24)) = int32(0)
	goto L247
L251:
	;
	v843 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v843 == int32(0) {
		goto L255
	} else {
		goto L256
	}
L252:
	;
	v835 = *(*int32)(unsafe.Add(mBase, uint32(v832)+24))
	if v835 == int32(0) {
		goto L251
	} else {
		goto L253
	}
L253:
	;
	F_SPI_freeplan(m, v835)
	mBase = m.M
	v839 = m.ExcPending
	if v839 != 0 {
		goto L37
	} else {
		goto L254
	}
L254:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v832)+24)) = int32(0)
	goto L251
L255:
	;
	v854 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v854 == int32(0) {
		goto L1
	} else {
		goto L259
	}
L256:
	;
	v846 = *(*int32)(unsafe.Add(mBase, uint32(v843)+24))
	if v846 == int32(0) {
		goto L255
	} else {
		goto L257
	}
L257:
	;
	F_SPI_freeplan(m, v846)
	mBase = m.M
	v850 = m.ExcPending
	if v850 != 0 {
		goto L37
	} else {
		goto L258
	}
L258:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v843)+24)) = int32(0)
	goto L255
L259:
	;
	v857 = *(*int32)(unsafe.Add(mBase, uint32(v854)+4))
	if v857 <= int32(0) {
		goto L1
	} else {
		goto L260
	}
L260:
	;
	v862 = int32(0)
	goto L261
L261:
	;
	v868 = *(*int32)(unsafe.Add(mBase, uint32(v854)+12))
	v872 = *(*int32)(unsafe.Add(mBase, uint32(v868+v862<<(uint(int32(2))%32))))
	if v872 == int32(0) {
		goto L263
	} else {
		goto L264
	}
L262:
	;
	goto L1
L263:
	;
	v884 = v862 + int32(1)
	v885 = *(*int32)(unsafe.Add(mBase, uint32(v854)+4))
	if v884 < v885 {
		v862 = v884
		goto L261
	} else {
		goto L267
	}
L264:
	;
	v875 = *(*int32)(unsafe.Add(mBase, uint32(v872)+24))
	if v875 == int32(0) {
		goto L263
	} else {
		goto L265
	}
L265:
	;
	F_SPI_freeplan(m, v875)
	mBase = m.M
	v879 = m.ExcPending
	if v879 != 0 {
		goto L37
	} else {
		goto L266
	}
L266:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v872)+24)) = int32(0)
	goto L263
L267:
	;
	goto L262
L268:
	;
	v890 = *(*int32)(unsafe.Add(mBase, uint32(v887)+24))
	if v890 == int32(0) {
		goto L1
	} else {
		goto L269
	}
L269:
	;
	F_SPI_freeplan(m, v890)
	mBase = m.M
	v894 = m.ExcPending
	if v894 != 0 {
		goto L37
	} else {
		goto L270
	}
L270:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v887)+24)) = int32(0)
	goto L1
L271:
	;
	v900 = *(*int32)(unsafe.Add(mBase, uint32(v897)+24))
	if v900 == int32(0) {
		goto L1
	} else {
		goto L272
	}
L272:
	;
	F_SPI_freeplan(m, v900)
	mBase = m.M
	v904 = m.ExcPending
	if v904 != 0 {
		goto L37
	} else {
		goto L273
	}
L273:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v897)+24)) = int32(0)
	goto L1
L274:
	;
	v910 = *(*int32)(unsafe.Add(mBase, uint32(v907)+24))
	if v910 == int32(0) {
		goto L1
	} else {
		goto L275
	}
L275:
	;
	F_SPI_freeplan(m, v910)
	mBase = m.M
	v914 = m.ExcPending
	if v914 != 0 {
		goto L37
	} else {
		goto L276
	}
L276:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v907)+24)) = int32(0)
	goto L1
L277:
	;
	v921 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v921
	F_errmsg_internal(m, int32(_a_F_free_stmt_1), v10)
	mBase = m.M
	v925 = m.ExcPending
	if v925 != 0 {
		goto L37
	} else {
		goto L278
	}
L278:
	;
	F_errfinish(m, int32(_a_F_free_stmt_2), int32(595), int32(_a_F_free_stmt_3))
	mBase = m.M
	v930 = m.ExcPending
	if v930 != 0 {
		goto L37
	} else {
		goto L279
	}
L279:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L280:
	;
	goto L1
L281:
	;
	F_SPI_freeplan(m, v932)
	mBase = m.M
	v936 = m.ExcPending
	if v936 != 0 {
		goto L37
	} else {
		goto L282
	}
L282:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v820)+24)) = int32(0)
	goto L6
L283:
	;
	v943 = *(*int32)(unsafe.Add(mBase, uint32(v940)+4))
	if v943 <= int32(0) {
		goto L1
	} else {
		goto L284
	}
L284:
	;
	v948 = int32(0)
	goto L285
L285:
	;
	v954 = *(*int32)(unsafe.Add(mBase, uint32(v940)+12))
	v958 = *(*int32)(unsafe.Add(mBase, uint32(v954+v948<<(uint(int32(2))%32))))
	if v958 == int32(0) {
		goto L287
	} else {
		goto L288
	}
L286:
	;
	goto L1
L287:
	;
	v970 = v948 + int32(1)
	v971 = *(*int32)(unsafe.Add(mBase, uint32(v940)+4))
	if v970 < v971 {
		v948 = v970
		goto L285
	} else {
		goto L291
	}
L288:
	;
	v961 = *(*int32)(unsafe.Add(mBase, uint32(v958)+24))
	if v961 == int32(0) {
		goto L287
	} else {
		goto L289
	}
L289:
	;
	F_SPI_freeplan(m, v961)
	mBase = m.M
	v965 = m.ExcPending
	if v965 != 0 {
		goto L37
	} else {
		goto L290
	}
L290:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v958)+24)) = int32(0)
	goto L287
L291:
	;
	goto L286
L292:
	;
	F_SPI_freeplan(m, v973)
	mBase = m.M
	v977 = m.ExcPending
	if v977 != 0 {
		goto L37
	} else {
		goto L293
	}
L293:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v519)+24)) = int32(0)
	goto L1
L294:
	;
	F_SPI_freeplan(m, v980)
	mBase = m.M
	v984 = m.ExcPending
	if v984 != 0 {
		goto L37
	} else {
		goto L295
	}
L295:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v487)+24)) = int32(0)
	goto L1
L296:
	;
	v991 = v2
	goto L297
L297:
	;
	v997 = *(*int32)(unsafe.Add(mBase, uint32(v931)+12))
	v1001 = *(*int32)(unsafe.Add(mBase, uint32(v997+v991<<(uint(int32(2))%32))))
	F_free_stmt(m, v1001)
	mBase = m.M
	v1003 = m.ExcPending
	if v1003 != 0 {
		goto L37
	} else {
		goto L299
	}
L298:
	;
	goto L1
L299:
	;
	v1005 = v991 + int32(1)
	v1006 = *(*int32)(unsafe.Add(mBase, uint32(v931)+4))
	if v1005 < v1006 {
		v991 = v1005
		goto L297
	} else {
		goto L300
	}
L300:
	;
	goto L298
}
