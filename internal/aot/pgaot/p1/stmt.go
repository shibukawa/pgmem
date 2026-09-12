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
	var v166 int32
	_ = v166
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
	var v233 int32
	_ = v233
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
	var v274 int32
	_ = v274
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
	var v341 int32
	_ = v341
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
	var v375 int32
	_ = v375
	var v379 int32
	_ = v379
	var v381 int32
	_ = v381
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v386 int32
	_ = v386
	var v389 int32
	_ = v389
	var v393 int32
	_ = v393
	var v397 int32
	_ = v397
	var v400 int32
	_ = v400
	var v405 int32
	_ = v405
	var v411 int32
	_ = v411
	var v415 int32
	_ = v415
	var v417 int32
	_ = v417
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v422 int32
	_ = v422
	var v425 int32
	_ = v425
	var v429 int32
	_ = v429
	var v433 int32
	_ = v433
	var v436 int32
	_ = v436
	var v440 int32
	_ = v440
	var v444 int32
	_ = v444
	var v447 int32
	_ = v447
	var v451 int32
	_ = v451
	var v455 int32
	_ = v455
	var v458 int32
	_ = v458
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
	var v480 int32
	_ = v480
	var v483 int32
	_ = v483
	var v487 int32
	_ = v487
	var v493 int32
	_ = v493
	var v497 int32
	_ = v497
	var v499 int32
	_ = v499
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v511 int32
	_ = v511
	var v514 int32
	_ = v514
	var v518 int32
	_ = v518
	var v521 int32
	_ = v521
	var v524 int32
	_ = v524
	var v528 int32
	_ = v528
	var v534 int32
	_ = v534
	var v538 int32
	_ = v538
	var v540 int32
	_ = v540
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v552 int32
	_ = v552
	var v555 int32
	_ = v555
	var v559 int32
	_ = v559
	var v562 int32
	_ = v562
	var v565 int32
	_ = v565
	var v569 int32
	_ = v569
	var v573 int32
	_ = v573
	var v576 int32
	_ = v576
	var v581 int32
	_ = v581
	var v587 int32
	_ = v587
	var v591 int32
	_ = v591
	var v593 int32
	_ = v593
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v598 int32
	_ = v598
	var v601 int32
	_ = v601
	var v605 int32
	_ = v605
	var v608 int32
	_ = v608
	var v611 int32
	_ = v611
	var v615 int32
	_ = v615
	var v618 int32
	_ = v618
	var v621 int32
	_ = v621
	var v625 int32
	_ = v625
	var v628 int32
	_ = v628
	var v631 int32
	_ = v631
	var v635 int32
	_ = v635
	var v639 int32
	_ = v639
	var v642 int32
	_ = v642
	var v646 int32
	_ = v646
	var v650 int32
	_ = v650
	var v653 int32
	_ = v653
	var v658 int32
	_ = v658
	var v664 int32
	_ = v664
	var v668 int32
	_ = v668
	var v671 int32
	_ = v671
	var v675 int32
	_ = v675
	var v680 int32
	_ = v680
	var v681 int32
	_ = v681
	var v683 int32
	_ = v683
	var v686 int32
	_ = v686
	var v690 int32
	_ = v690
	var v696 int32
	_ = v696
	var v700 int32
	_ = v700
	var v703 int32
	_ = v703
	var v707 int32
	_ = v707
	var v712 int32
	_ = v712
	var v713 int32
	_ = v713
	var v722 int32
	_ = v722
	var v725 int32
	_ = v725
	var v730 int32
	_ = v730
	var v736 int32
	_ = v736
	var v740 int32
	_ = v740
	var v741 int32
	_ = v741
	var v744 int32
	_ = v744
	var v748 int32
	_ = v748
	var v753 int32
	_ = v753
	var v754 int32
	_ = v754
	var v756 int32
	_ = v756
	var v759 int32
	_ = v759
	var v763 int32
	_ = v763
	var v767 int32
	_ = v767
	var v770 int32
	_ = v770
	var v774 int32
	_ = v774
	var v777 int32
	_ = v777
	var v780 int32
	_ = v780
	var v784 int32
	_ = v784
	var v787 int32
	_ = v787
	var v790 int32
	_ = v790
	var v794 int32
	_ = v794
	var v798 int32
	_ = v798
	var v801 int32
	_ = v801
	var v806 int32
	_ = v806
	var v812 int32
	_ = v812
	var v816 int32
	_ = v816
	var v819 int32
	_ = v819
	var v823 int32
	_ = v823
	var v828 int32
	_ = v828
	var v829 int32
	_ = v829
	var v831 int32
	_ = v831
	var v834 int32
	_ = v834
	var v838 int32
	_ = v838
	var v844 int32
	_ = v844
	var v848 int32
	_ = v848
	var v850 int32
	_ = v850
	var v852 int32
	_ = v852
	var v853 int32
	_ = v853
	var v862 int32
	_ = v862
	var v865 int32
	_ = v865
	var v869 int32
	_ = v869
	var v873 int32
	_ = v873
	var v876 int32
	_ = v876
	var v881 int32
	_ = v881
	var v887 int32
	_ = v887
	var v891 int32
	_ = v891
	var v894 int32
	_ = v894
	var v898 int32
	_ = v898
	var v903 int32
	_ = v903
	var v904 int32
	_ = v904
	var v906 int32
	_ = v906
	var v909 int32
	_ = v909
	var v913 int32
	_ = v913
	var v917 int32
	_ = v917
	var v920 int32
	_ = v920
	var v924 int32
	_ = v924
	var v928 int32
	_ = v928
	var v931 int32
	_ = v931
	var v935 int32
	_ = v935
	var v939 int32
	_ = v939
	var v942 int32
	_ = v942
	var v947 int32
	_ = v947
	var v953 int32
	_ = v953
	var v957 int32
	_ = v957
	var v960 int32
	_ = v960
	var v964 int32
	_ = v964
	var v969 int32
	_ = v969
	var v970 int32
	_ = v970
	var v972 int32
	_ = v972
	var v975 int32
	_ = v975
	var v979 int32
	_ = v979
	var v982 int32
	_ = v982
	var v985 int32
	_ = v985
	var v989 int32
	_ = v989
	var v992 int32
	_ = v992
	var v995 int32
	_ = v995
	var v999 int32
	_ = v999
	var v1017 int32
	_ = v1017
	var v1018 int32
	_ = v1018
	var v1023 int32
	_ = v1023
	var v1030 int32
	_ = v1030
	v2 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	if l0 == v2 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(556174))
	mBase = m.M
	v1017 = m.ExcPending
	if v1017 != 0 {
		goto L32
	} else {
		goto L294
	}
L2:
	;
	m.G0 = v10 + int32(16)
	return
L3:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch v14 {
	case 0:
		goto L26
	case 1:
		goto L25
	case 2:
		goto L24
	case 3:
		goto L23
	case 4:
		goto L22
	case 5:
		goto L21
	case 6:
		goto L20
	case 7:
		goto L19
	case 8:
		goto L18
	case 9:
		goto L17
	case 10:
		goto L16
	case 11:
		goto L15
	case 12:
		goto L14
	case 13:
		goto L13
	case 14:
		goto L12
	case 15:
		goto L11
	case 16:
		goto L10
	case 17:
		goto L9
	case 18:
		goto L8
	case 19, 22, 25, 26:
		goto L2
	case 20:
		goto L7
	case 21:
		goto L6
	case 23:
		goto L5
	case 24:
		goto L4
	default:
		goto L1
	}
L4:
	;
	v992 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v992 == int32(0) {
		goto L2
	} else {
		goto L291
	}
L5:
	;
	v982 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v982 == int32(0) {
		goto L2
	} else {
		goto L288
	}
L6:
	;
	v972 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v972 == int32(0) {
		goto L2
	} else {
		goto L285
	}
L7:
	;
	v906 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v906 == int32(0) {
		goto L264
	} else {
		goto L265
	}
L8:
	;
	v831 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v831 == int32(0) {
		goto L244
	} else {
		goto L245
	}
L9:
	;
	v787 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v787 == int32(0) {
		goto L231
	} else {
		goto L232
	}
L10:
	;
	v777 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v777 == int32(0) {
		goto L2
	} else {
		goto L228
	}
L11:
	;
	v756 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v756 == int32(0) {
		goto L221
	} else {
		goto L222
	}
L12:
	;
	v683 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v683 == int32(0) {
		goto L202
	} else {
		goto L203
	}
L13:
	;
	v628 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v628 == int32(0) {
		goto L185
	} else {
		goto L186
	}
L14:
	;
	v618 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v618 == int32(0) {
		goto L2
	} else {
		goto L182
	}
L15:
	;
	v608 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v608 == int32(0) {
		goto L2
	} else {
		goto L179
	}
L16:
	;
	v598 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v598 == int32(0) {
		goto L2
	} else {
		goto L176
	}
L17:
	;
	v562 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v562 == int32(0) {
		goto L166
	} else {
		goto L167
	}
L18:
	;
	v521 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v521 == int32(0) {
		goto L156
	} else {
		goto L157
	}
L19:
	;
	v480 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v480 == int32(0) {
		goto L146
	} else {
		goto L147
	}
L20:
	;
	v422 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v422 == int32(0) {
		goto L128
	} else {
		goto L129
	}
L21:
	;
	v386 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v386 == int32(0) {
		goto L118
	} else {
		goto L119
	}
L22:
	;
	v362 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v362 == int32(0) {
		goto L2
	} else {
		goto L112
	}
L23:
	;
	v254 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v254 == int32(0) {
		goto L85
	} else {
		goto L86
	}
L24:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v114 == int32(0) {
		goto L51
	} else {
		goto L52
	}
L25:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v104 == int32(0) {
		goto L2
	} else {
		goto L48
	}
L26:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v15 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v46 == int32(0) {
		goto L2
	} else {
		goto L35
	}
L28:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	if v18 <= int32(0) {
		goto L27
	} else {
		goto L29
	}
L29:
	;
	v22 = v2
	goto L30
L30:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v28+v22<<(uint(int32(2))%32))))
	F_free_stmt(m, v32)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	goto L27
L32:
	;
	return
L33:
	;
	v36 = v22 + int32(1)
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	if v36 < v37 {
		v22 = v36
		goto L30
	} else {
		goto L34
	}
L34:
	;
	goto L31
L35:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v46)+8))
	if v49 == int32(0) {
		goto L2
	} else {
		goto L36
	}
L36:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v49)+4))
	if v52 <= int32(0) {
		goto L2
	} else {
		goto L37
	}
L37:
	;
	v56 = int32(0)
	goto L38
L38:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v49)+12))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v63+v56<<(uint(int32(2))%32))))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)+8))
	if v68 == int32(0) {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	goto L2
L40:
	;
	v101 = v56 + int32(1)
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v49)+4))
	if v101 < v102 {
		v56 = v101
		goto L38
	} else {
		goto L47
	}
L41:
	;
	v71 = int32(0)
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v68)+4))
	if v72 <= v71 {
		goto L40
	} else {
		goto L42
	}
L42:
	;
	v76 = v71
	goto L43
L43:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v68)+12))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v82+v76<<(uint(int32(2))%32))))
	F_free_stmt(m, v86)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L32
	} else {
		goto L45
	}
L44:
	;
	goto L40
L45:
	;
	v90 = v76 + int32(1)
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v68)+4))
	if v90 < v91 {
		v76 = v90
		goto L43
	} else {
		goto L46
	}
L46:
	;
	goto L44
L47:
	;
	goto L39
L48:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v104)+24))
	if v107 == int32(0) {
		goto L2
	} else {
		goto L49
	}
L49:
	;
	F_SPI_freeplan(m, v107)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L32
	} else {
		goto L50
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v104)+24)) = int32(0)
	goto L2
L51:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v125 == int32(0) {
		goto L55
	} else {
		goto L56
	}
L52:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v114)+24))
	if v117 == int32(0) {
		goto L51
	} else {
		goto L53
	}
L53:
	;
	F_SPI_freeplan(m, v117)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L32
	} else {
		goto L54
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v114)+24)) = int32(0)
	goto L51
L55:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v157 == int32(0) {
		goto L62
	} else {
		goto L63
	}
L56:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v125)+4))
	if v128 <= int32(0) {
		goto L55
	} else {
		goto L57
	}
L57:
	;
	v133 = int32(0)
	goto L58
L58:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v125)+12))
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v139+v133<<(uint(int32(2))%32))))
	F_free_stmt(m, v143)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L32
	} else {
		goto L60
	}
L59:
	;
	goto L55
L60:
	;
	v147 = v133 + int32(1)
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v125)+4))
	if v147 < v148 {
		v133 = v147
		goto L58
	} else {
		goto L61
	}
L61:
	;
	goto L59
L62:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v229 == int32(0) {
		goto L2
	} else {
		goto L79
	}
L63:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v157)+4))
	if v160 <= int32(0) {
		goto L62
	} else {
		goto L64
	}
L64:
	;
	v166 = v2
	goto L65
L65:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v157)+12))
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v170+v166<<(uint(int32(2))%32))))
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v174)+4))
	if v175 == int32(0) {
		goto L67
	} else {
		goto L68
	}
L66:
	;
	goto L62
L67:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v174)+8))
	if v186 == int32(0) {
		goto L71
	} else {
		goto L72
	}
L68:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v175)+24))
	if v178 == int32(0) {
		goto L67
	} else {
		goto L69
	}
L69:
	;
	F_SPI_freeplan(m, v178)
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L32
	} else {
		goto L70
	}
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v175)+24)) = int32(0)
	goto L67
L71:
	;
	v219 = v166 + int32(1)
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v157)+4))
	if v219 < v220 {
		v166 = v219
		goto L65
	} else {
		goto L78
	}
L72:
	;
	v189 = int32(0)
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v186)+4))
	if v190 <= v189 {
		goto L71
	} else {
		goto L73
	}
L73:
	;
	v194 = v189
	goto L74
L74:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v186)+12))
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v200+v194<<(uint(int32(2))%32))))
	F_free_stmt(m, v204)
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L32
	} else {
		goto L76
	}
L75:
	;
	goto L71
L76:
	;
	v208 = v194 + int32(1)
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v186)+4))
	if v208 < v209 {
		v194 = v208
		goto L74
	} else {
		goto L77
	}
L77:
	;
	goto L75
L78:
	;
	goto L66
L79:
	;
	v232 = int32(0)
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v229)+4))
	if v233 <= v232 {
		goto L2
	} else {
		goto L80
	}
L80:
	;
	v237 = v232
	goto L81
L81:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v229)+12))
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v243+v237<<(uint(int32(2))%32))))
	F_free_stmt(m, v247)
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L32
	} else {
		goto L83
	}
L82:
	;
	goto L2
L83:
	;
	v251 = v237 + int32(1)
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v229)+4))
	if v251 < v252 {
		v237 = v251
		goto L81
	} else {
		goto L84
	}
L84:
	;
	goto L82
L85:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v265 == int32(0) {
		goto L89
	} else {
		goto L90
	}
L86:
	;
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v254)+24))
	if v257 == int32(0) {
		goto L85
	} else {
		goto L87
	}
L87:
	;
	F_SPI_freeplan(m, v257)
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L32
	} else {
		goto L88
	}
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v254)+24)) = int32(0)
	goto L85
L89:
	;
	v337 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v337 == int32(0) {
		goto L2
	} else {
		goto L106
	}
L90:
	;
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v265)+4))
	if v268 <= int32(0) {
		goto L89
	} else {
		goto L91
	}
L91:
	;
	v274 = v2
	goto L92
L92:
	;
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v265)+12))
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v278+v274<<(uint(int32(2))%32))))
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v282)+4))
	if v283 == int32(0) {
		goto L94
	} else {
		goto L95
	}
L93:
	;
	goto L89
L94:
	;
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v282)+8))
	if v294 == int32(0) {
		goto L98
	} else {
		goto L99
	}
L95:
	;
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v283)+24))
	if v286 == int32(0) {
		goto L94
	} else {
		goto L96
	}
L96:
	;
	F_SPI_freeplan(m, v286)
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L32
	} else {
		goto L97
	}
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v283)+24)) = int32(0)
	goto L94
L98:
	;
	v327 = v274 + int32(1)
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v265)+4))
	if v327 < v328 {
		v274 = v327
		goto L92
	} else {
		goto L105
	}
L99:
	;
	v297 = int32(0)
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v294)+4))
	if v298 <= v297 {
		goto L98
	} else {
		goto L100
	}
L100:
	;
	v302 = v297
	goto L101
L101:
	;
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v294)+12))
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v308+v302<<(uint(int32(2))%32))))
	F_free_stmt(m, v312)
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L32
	} else {
		goto L103
	}
L102:
	;
	goto L98
L103:
	;
	v316 = v302 + int32(1)
	v317 = *(*int32)(unsafe.Add(mBase, uint32(v294)+4))
	if v316 < v317 {
		v302 = v316
		goto L101
	} else {
		goto L104
	}
L104:
	;
	goto L102
L105:
	;
	goto L93
L106:
	;
	v340 = int32(0)
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v337)+4))
	if v341 <= v340 {
		goto L2
	} else {
		goto L107
	}
L107:
	;
	v345 = v340
	goto L108
L108:
	;
	v351 = *(*int32)(unsafe.Add(mBase, uint32(v337)+12))
	v355 = *(*int32)(unsafe.Add(mBase, uint32(v351+v345<<(uint(int32(2))%32))))
	F_free_stmt(m, v355)
	mBase = m.M
	v357 = m.ExcPending
	if v357 != 0 {
		goto L32
	} else {
		goto L110
	}
L109:
	;
	goto L2
L110:
	;
	v359 = v345 + int32(1)
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v337)+4))
	if v359 < v360 {
		v345 = v359
		goto L108
	} else {
		goto L111
	}
L111:
	;
	goto L109
L112:
	;
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v362)+4))
	if v365 <= int32(0) {
		goto L2
	} else {
		goto L113
	}
L113:
	;
	v369 = v2
	goto L114
L114:
	;
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v362)+12))
	v379 = *(*int32)(unsafe.Add(mBase, uint32(v375+v369<<(uint(int32(2))%32))))
	F_free_stmt(m, v379)
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L32
	} else {
		goto L116
	}
L115:
	;
	goto L2
L116:
	;
	v383 = v369 + int32(1)
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v362)+4))
	if v383 < v384 {
		v369 = v383
		goto L114
	} else {
		goto L117
	}
L117:
	;
	goto L115
L118:
	;
	v397 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v397 == int32(0) {
		goto L2
	} else {
		goto L122
	}
L119:
	;
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v386)+24))
	if v389 == int32(0) {
		goto L118
	} else {
		goto L120
	}
L120:
	;
	F_SPI_freeplan(m, v389)
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L32
	} else {
		goto L121
	}
L121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v386)+24)) = int32(0)
	goto L118
L122:
	;
	v400 = *(*int32)(unsafe.Add(mBase, uint32(v397)+4))
	if v400 <= int32(0) {
		goto L2
	} else {
		goto L123
	}
L123:
	;
	v405 = int32(0)
	goto L124
L124:
	;
	v411 = *(*int32)(unsafe.Add(mBase, uint32(v397)+12))
	v415 = *(*int32)(unsafe.Add(mBase, uint32(v411+v405<<(uint(int32(2))%32))))
	F_free_stmt(m, v415)
	mBase = m.M
	v417 = m.ExcPending
	if v417 != 0 {
		goto L32
	} else {
		goto L126
	}
L125:
	;
	goto L2
L126:
	;
	v419 = v405 + int32(1)
	v420 = *(*int32)(unsafe.Add(mBase, uint32(v397)+4))
	if v419 < v420 {
		v405 = v419
		goto L124
	} else {
		goto L127
	}
L127:
	;
	goto L125
L128:
	;
	v433 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v433 == int32(0) {
		goto L132
	} else {
		goto L133
	}
L129:
	;
	v425 = *(*int32)(unsafe.Add(mBase, uint32(v422)+24))
	if v425 == int32(0) {
		goto L128
	} else {
		goto L130
	}
L130:
	;
	F_SPI_freeplan(m, v425)
	mBase = m.M
	v429 = m.ExcPending
	if v429 != 0 {
		goto L32
	} else {
		goto L131
	}
L131:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v422)+24)) = int32(0)
	goto L128
L132:
	;
	v444 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v444 == int32(0) {
		goto L136
	} else {
		goto L137
	}
L133:
	;
	v436 = *(*int32)(unsafe.Add(mBase, uint32(v433)+24))
	if v436 == int32(0) {
		goto L132
	} else {
		goto L134
	}
L134:
	;
	F_SPI_freeplan(m, v436)
	mBase = m.M
	v440 = m.ExcPending
	if v440 != 0 {
		goto L32
	} else {
		goto L135
	}
L135:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v433)+24)) = int32(0)
	goto L132
L136:
	;
	v455 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v455 == int32(0) {
		goto L2
	} else {
		goto L140
	}
L137:
	;
	v447 = *(*int32)(unsafe.Add(mBase, uint32(v444)+24))
	if v447 == int32(0) {
		goto L136
	} else {
		goto L138
	}
L138:
	;
	F_SPI_freeplan(m, v447)
	mBase = m.M
	v451 = m.ExcPending
	if v451 != 0 {
		goto L32
	} else {
		goto L139
	}
L139:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v444)+24)) = int32(0)
	goto L136
L140:
	;
	v458 = *(*int32)(unsafe.Add(mBase, uint32(v455)+4))
	if v458 <= int32(0) {
		goto L2
	} else {
		goto L141
	}
L141:
	;
	v463 = int32(0)
	goto L142
L142:
	;
	v469 = *(*int32)(unsafe.Add(mBase, uint32(v455)+12))
	v473 = *(*int32)(unsafe.Add(mBase, uint32(v469+v463<<(uint(int32(2))%32))))
	F_free_stmt(m, v473)
	mBase = m.M
	v475 = m.ExcPending
	if v475 != 0 {
		goto L32
	} else {
		goto L144
	}
L143:
	;
	goto L2
L144:
	;
	v477 = v463 + int32(1)
	v478 = *(*int32)(unsafe.Add(mBase, uint32(v455)+4))
	if v477 < v478 {
		v463 = v477
		goto L142
	} else {
		goto L145
	}
L145:
	;
	goto L143
L146:
	;
	v511 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v511 == int32(0) {
		goto L2
	} else {
		goto L153
	}
L147:
	;
	v483 = *(*int32)(unsafe.Add(mBase, uint32(v480)+4))
	if v483 <= int32(0) {
		goto L146
	} else {
		goto L148
	}
L148:
	;
	v487 = v2
	goto L149
L149:
	;
	v493 = *(*int32)(unsafe.Add(mBase, uint32(v480)+12))
	v497 = *(*int32)(unsafe.Add(mBase, uint32(v493+v487<<(uint(int32(2))%32))))
	F_free_stmt(m, v497)
	mBase = m.M
	v499 = m.ExcPending
	if v499 != 0 {
		goto L32
	} else {
		goto L151
	}
L150:
	;
	goto L146
L151:
	;
	v501 = v487 + int32(1)
	v502 = *(*int32)(unsafe.Add(mBase, uint32(v480)+4))
	if v501 < v502 {
		v487 = v501
		goto L149
	} else {
		goto L152
	}
L152:
	;
	goto L150
L153:
	;
	v514 = *(*int32)(unsafe.Add(mBase, uint32(v511)+24))
	if v514 == int32(0) {
		goto L2
	} else {
		goto L154
	}
L154:
	;
	F_SPI_freeplan(m, v514)
	mBase = m.M
	v518 = m.ExcPending
	if v518 != 0 {
		goto L32
	} else {
		goto L155
	}
L155:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v511)+24)) = int32(0)
	goto L2
L156:
	;
	v552 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v552 == int32(0) {
		goto L2
	} else {
		goto L163
	}
L157:
	;
	v524 = *(*int32)(unsafe.Add(mBase, uint32(v521)+4))
	if v524 <= int32(0) {
		goto L156
	} else {
		goto L158
	}
L158:
	;
	v528 = v2
	goto L159
L159:
	;
	v534 = *(*int32)(unsafe.Add(mBase, uint32(v521)+12))
	v538 = *(*int32)(unsafe.Add(mBase, uint32(v534+v528<<(uint(int32(2))%32))))
	F_free_stmt(m, v538)
	mBase = m.M
	v540 = m.ExcPending
	if v540 != 0 {
		goto L32
	} else {
		goto L161
	}
L160:
	;
	goto L156
L161:
	;
	v542 = v528 + int32(1)
	v543 = *(*int32)(unsafe.Add(mBase, uint32(v521)+4))
	if v542 < v543 {
		v528 = v542
		goto L159
	} else {
		goto L162
	}
L162:
	;
	goto L160
L163:
	;
	v555 = *(*int32)(unsafe.Add(mBase, uint32(v552)+24))
	if v555 == int32(0) {
		goto L2
	} else {
		goto L164
	}
L164:
	;
	F_SPI_freeplan(m, v555)
	mBase = m.M
	v559 = m.ExcPending
	if v559 != 0 {
		goto L32
	} else {
		goto L165
	}
L165:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v552)+24)) = int32(0)
	goto L2
L166:
	;
	v573 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v573 == int32(0) {
		goto L2
	} else {
		goto L170
	}
L167:
	;
	v565 = *(*int32)(unsafe.Add(mBase, uint32(v562)+24))
	if v565 == int32(0) {
		goto L166
	} else {
		goto L168
	}
L168:
	;
	F_SPI_freeplan(m, v565)
	mBase = m.M
	v569 = m.ExcPending
	if v569 != 0 {
		goto L32
	} else {
		goto L169
	}
L169:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v562)+24)) = int32(0)
	goto L166
L170:
	;
	v576 = *(*int32)(unsafe.Add(mBase, uint32(v573)+4))
	if v576 <= int32(0) {
		goto L2
	} else {
		goto L171
	}
L171:
	;
	v581 = int32(0)
	goto L172
L172:
	;
	v587 = *(*int32)(unsafe.Add(mBase, uint32(v573)+12))
	v591 = *(*int32)(unsafe.Add(mBase, uint32(v587+v581<<(uint(int32(2))%32))))
	F_free_stmt(m, v591)
	mBase = m.M
	v593 = m.ExcPending
	if v593 != 0 {
		goto L32
	} else {
		goto L174
	}
L173:
	;
	goto L2
L174:
	;
	v595 = v581 + int32(1)
	v596 = *(*int32)(unsafe.Add(mBase, uint32(v573)+4))
	if v595 < v596 {
		v581 = v595
		goto L172
	} else {
		goto L175
	}
L175:
	;
	goto L173
L176:
	;
	v601 = *(*int32)(unsafe.Add(mBase, uint32(v598)+24))
	if v601 == int32(0) {
		goto L2
	} else {
		goto L177
	}
L177:
	;
	F_SPI_freeplan(m, v601)
	mBase = m.M
	v605 = m.ExcPending
	if v605 != 0 {
		goto L32
	} else {
		goto L178
	}
L178:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v598)+24)) = int32(0)
	goto L2
L179:
	;
	v611 = *(*int32)(unsafe.Add(mBase, uint32(v608)+24))
	if v611 == int32(0) {
		goto L2
	} else {
		goto L180
	}
L180:
	;
	F_SPI_freeplan(m, v611)
	mBase = m.M
	v615 = m.ExcPending
	if v615 != 0 {
		goto L32
	} else {
		goto L181
	}
L181:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v608)+24)) = int32(0)
	goto L2
L182:
	;
	v621 = *(*int32)(unsafe.Add(mBase, uint32(v618)+24))
	if v621 == int32(0) {
		goto L2
	} else {
		goto L183
	}
L183:
	;
	F_SPI_freeplan(m, v621)
	mBase = m.M
	v625 = m.ExcPending
	if v625 != 0 {
		goto L32
	} else {
		goto L184
	}
L184:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v618)+24)) = int32(0)
	goto L2
L185:
	;
	v639 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v639 == int32(0) {
		goto L189
	} else {
		goto L190
	}
L186:
	;
	v631 = *(*int32)(unsafe.Add(mBase, uint32(v628)+24))
	if v631 == int32(0) {
		goto L185
	} else {
		goto L187
	}
L187:
	;
	F_SPI_freeplan(m, v631)
	mBase = m.M
	v635 = m.ExcPending
	if v635 != 0 {
		goto L32
	} else {
		goto L188
	}
L188:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v628)+24)) = int32(0)
	goto L185
L189:
	;
	v650 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v650 == int32(0) {
		goto L2
	} else {
		goto L193
	}
L190:
	;
	v642 = *(*int32)(unsafe.Add(mBase, uint32(v639)+24))
	if v642 == int32(0) {
		goto L189
	} else {
		goto L191
	}
L191:
	;
	F_SPI_freeplan(m, v642)
	mBase = m.M
	v646 = m.ExcPending
	if v646 != 0 {
		goto L32
	} else {
		goto L192
	}
L192:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v639)+24)) = int32(0)
	goto L189
L193:
	;
	v653 = *(*int32)(unsafe.Add(mBase, uint32(v650)+4))
	if v653 <= int32(0) {
		goto L2
	} else {
		goto L194
	}
L194:
	;
	v658 = int32(0)
	goto L195
L195:
	;
	v664 = *(*int32)(unsafe.Add(mBase, uint32(v650)+12))
	v668 = *(*int32)(unsafe.Add(mBase, uint32(v664+v658<<(uint(int32(2))%32))))
	if v668 == int32(0) {
		goto L197
	} else {
		goto L198
	}
L196:
	;
	goto L2
L197:
	;
	v680 = v658 + int32(1)
	v681 = *(*int32)(unsafe.Add(mBase, uint32(v650)+4))
	if v680 < v681 {
		v658 = v680
		goto L195
	} else {
		goto L201
	}
L198:
	;
	v671 = *(*int32)(unsafe.Add(mBase, uint32(v668)+24))
	if v671 == int32(0) {
		goto L197
	} else {
		goto L199
	}
L199:
	;
	F_SPI_freeplan(m, v671)
	mBase = m.M
	v675 = m.ExcPending
	if v675 != 0 {
		goto L32
	} else {
		goto L200
	}
L200:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v668)+24)) = int32(0)
	goto L197
L201:
	;
	goto L196
L202:
	;
	v722 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v722 == int32(0) {
		goto L2
	} else {
		goto L212
	}
L203:
	;
	v686 = *(*int32)(unsafe.Add(mBase, uint32(v683)+4))
	if v686 <= int32(0) {
		goto L202
	} else {
		goto L204
	}
L204:
	;
	v690 = v2
	goto L205
L205:
	;
	v696 = *(*int32)(unsafe.Add(mBase, uint32(v683)+12))
	v700 = *(*int32)(unsafe.Add(mBase, uint32(v696+v690<<(uint(int32(2))%32))))
	if v700 == int32(0) {
		goto L207
	} else {
		goto L208
	}
L206:
	;
	goto L202
L207:
	;
	v712 = v690 + int32(1)
	v713 = *(*int32)(unsafe.Add(mBase, uint32(v683)+4))
	if v712 < v713 {
		v690 = v712
		goto L205
	} else {
		goto L211
	}
L208:
	;
	v703 = *(*int32)(unsafe.Add(mBase, uint32(v700)+24))
	if v703 == int32(0) {
		goto L207
	} else {
		goto L209
	}
L209:
	;
	F_SPI_freeplan(m, v703)
	mBase = m.M
	v707 = m.ExcPending
	if v707 != 0 {
		goto L32
	} else {
		goto L210
	}
L210:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v700)+24)) = int32(0)
	goto L207
L211:
	;
	goto L206
L212:
	;
	v725 = *(*int32)(unsafe.Add(mBase, uint32(v722)+4))
	if v725 <= int32(0) {
		goto L2
	} else {
		goto L213
	}
L213:
	;
	v730 = int32(0)
	goto L214
L214:
	;
	v736 = *(*int32)(unsafe.Add(mBase, uint32(v722)+12))
	v740 = *(*int32)(unsafe.Add(mBase, uint32(v736+v730<<(uint(int32(2))%32))))
	v741 = *(*int32)(unsafe.Add(mBase, uint32(v740)+4))
	if v741 == int32(0) {
		goto L216
	} else {
		goto L217
	}
L215:
	;
	goto L2
L216:
	;
	v753 = v730 + int32(1)
	v754 = *(*int32)(unsafe.Add(mBase, uint32(v722)+4))
	if v753 < v754 {
		v730 = v753
		goto L214
	} else {
		goto L220
	}
L217:
	;
	v744 = *(*int32)(unsafe.Add(mBase, uint32(v741)+24))
	if v744 == int32(0) {
		goto L216
	} else {
		goto L218
	}
L218:
	;
	F_SPI_freeplan(m, v744)
	mBase = m.M
	v748 = m.ExcPending
	if v748 != 0 {
		goto L32
	} else {
		goto L219
	}
L219:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v741)+24)) = int32(0)
	goto L216
L220:
	;
	goto L215
L221:
	;
	v767 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v767 == int32(0) {
		goto L2
	} else {
		goto L225
	}
L222:
	;
	v759 = *(*int32)(unsafe.Add(mBase, uint32(v756)+24))
	if v759 == int32(0) {
		goto L221
	} else {
		goto L223
	}
L223:
	;
	F_SPI_freeplan(m, v759)
	mBase = m.M
	v763 = m.ExcPending
	if v763 != 0 {
		goto L32
	} else {
		goto L224
	}
L224:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v756)+24)) = int32(0)
	goto L221
L225:
	;
	v770 = *(*int32)(unsafe.Add(mBase, uint32(v767)+24))
	if v770 == int32(0) {
		goto L2
	} else {
		goto L226
	}
L226:
	;
	F_SPI_freeplan(m, v770)
	mBase = m.M
	v774 = m.ExcPending
	if v774 != 0 {
		goto L32
	} else {
		goto L227
	}
L227:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v767)+24)) = int32(0)
	goto L2
L228:
	;
	v780 = *(*int32)(unsafe.Add(mBase, uint32(v777)+24))
	if v780 == int32(0) {
		goto L2
	} else {
		goto L229
	}
L229:
	;
	F_SPI_freeplan(m, v780)
	mBase = m.M
	v784 = m.ExcPending
	if v784 != 0 {
		goto L32
	} else {
		goto L230
	}
L230:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v777)+24)) = int32(0)
	goto L2
L231:
	;
	v798 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v798 == int32(0) {
		goto L2
	} else {
		goto L235
	}
L232:
	;
	v790 = *(*int32)(unsafe.Add(mBase, uint32(v787)+24))
	if v790 == int32(0) {
		goto L231
	} else {
		goto L233
	}
L233:
	;
	F_SPI_freeplan(m, v790)
	mBase = m.M
	v794 = m.ExcPending
	if v794 != 0 {
		goto L32
	} else {
		goto L234
	}
L234:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v787)+24)) = int32(0)
	goto L231
L235:
	;
	v801 = *(*int32)(unsafe.Add(mBase, uint32(v798)+4))
	if v801 <= int32(0) {
		goto L2
	} else {
		goto L236
	}
L236:
	;
	v806 = int32(0)
	goto L237
L237:
	;
	v812 = *(*int32)(unsafe.Add(mBase, uint32(v798)+12))
	v816 = *(*int32)(unsafe.Add(mBase, uint32(v812+v806<<(uint(int32(2))%32))))
	if v816 == int32(0) {
		goto L239
	} else {
		goto L240
	}
L238:
	;
	goto L2
L239:
	;
	v828 = v806 + int32(1)
	v829 = *(*int32)(unsafe.Add(mBase, uint32(v798)+4))
	if v828 < v829 {
		v806 = v828
		goto L237
	} else {
		goto L243
	}
L240:
	;
	v819 = *(*int32)(unsafe.Add(mBase, uint32(v816)+24))
	if v819 == int32(0) {
		goto L239
	} else {
		goto L241
	}
L241:
	;
	F_SPI_freeplan(m, v819)
	mBase = m.M
	v823 = m.ExcPending
	if v823 != 0 {
		goto L32
	} else {
		goto L242
	}
L242:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v816)+24)) = int32(0)
	goto L239
L243:
	;
	goto L238
L244:
	;
	v862 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v862 == int32(0) {
		goto L251
	} else {
		goto L252
	}
L245:
	;
	v834 = *(*int32)(unsafe.Add(mBase, uint32(v831)+4))
	if v834 <= int32(0) {
		goto L244
	} else {
		goto L246
	}
L246:
	;
	v838 = v2
	goto L247
L247:
	;
	v844 = *(*int32)(unsafe.Add(mBase, uint32(v831)+12))
	v848 = *(*int32)(unsafe.Add(mBase, uint32(v844+v838<<(uint(int32(2))%32))))
	F_free_stmt(m, v848)
	mBase = m.M
	v850 = m.ExcPending
	if v850 != 0 {
		goto L32
	} else {
		goto L249
	}
L248:
	;
	goto L244
L249:
	;
	v852 = v838 + int32(1)
	v853 = *(*int32)(unsafe.Add(mBase, uint32(v831)+4))
	if v852 < v853 {
		v838 = v852
		goto L247
	} else {
		goto L250
	}
L250:
	;
	goto L248
L251:
	;
	v873 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v873 == int32(0) {
		goto L2
	} else {
		goto L255
	}
L252:
	;
	v865 = *(*int32)(unsafe.Add(mBase, uint32(v862)+24))
	if v865 == int32(0) {
		goto L251
	} else {
		goto L253
	}
L253:
	;
	F_SPI_freeplan(m, v865)
	mBase = m.M
	v869 = m.ExcPending
	if v869 != 0 {
		goto L32
	} else {
		goto L254
	}
L254:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v862)+24)) = int32(0)
	goto L251
L255:
	;
	v876 = *(*int32)(unsafe.Add(mBase, uint32(v873)+4))
	if v876 <= int32(0) {
		goto L2
	} else {
		goto L256
	}
L256:
	;
	v881 = int32(0)
	goto L257
L257:
	;
	v887 = *(*int32)(unsafe.Add(mBase, uint32(v873)+12))
	v891 = *(*int32)(unsafe.Add(mBase, uint32(v887+v881<<(uint(int32(2))%32))))
	if v891 == int32(0) {
		goto L259
	} else {
		goto L260
	}
L258:
	;
	goto L2
L259:
	;
	v903 = v881 + int32(1)
	v904 = *(*int32)(unsafe.Add(mBase, uint32(v873)+4))
	if v903 < v904 {
		v881 = v903
		goto L257
	} else {
		goto L263
	}
L260:
	;
	v894 = *(*int32)(unsafe.Add(mBase, uint32(v891)+24))
	if v894 == int32(0) {
		goto L259
	} else {
		goto L261
	}
L261:
	;
	F_SPI_freeplan(m, v894)
	mBase = m.M
	v898 = m.ExcPending
	if v898 != 0 {
		goto L32
	} else {
		goto L262
	}
L262:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v891)+24)) = int32(0)
	goto L259
L263:
	;
	goto L258
L264:
	;
	v917 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v917 == int32(0) {
		goto L268
	} else {
		goto L269
	}
L265:
	;
	v909 = *(*int32)(unsafe.Add(mBase, uint32(v906)+24))
	if v909 == int32(0) {
		goto L264
	} else {
		goto L266
	}
L266:
	;
	F_SPI_freeplan(m, v909)
	mBase = m.M
	v913 = m.ExcPending
	if v913 != 0 {
		goto L32
	} else {
		goto L267
	}
L267:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v906)+24)) = int32(0)
	goto L264
L268:
	;
	v928 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v928 == int32(0) {
		goto L272
	} else {
		goto L273
	}
L269:
	;
	v920 = *(*int32)(unsafe.Add(mBase, uint32(v917)+24))
	if v920 == int32(0) {
		goto L268
	} else {
		goto L270
	}
L270:
	;
	F_SPI_freeplan(m, v920)
	mBase = m.M
	v924 = m.ExcPending
	if v924 != 0 {
		goto L32
	} else {
		goto L271
	}
L271:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v917)+24)) = int32(0)
	goto L268
L272:
	;
	v939 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v939 == int32(0) {
		goto L2
	} else {
		goto L276
	}
L273:
	;
	v931 = *(*int32)(unsafe.Add(mBase, uint32(v928)+24))
	if v931 == int32(0) {
		goto L272
	} else {
		goto L274
	}
L274:
	;
	F_SPI_freeplan(m, v931)
	mBase = m.M
	v935 = m.ExcPending
	if v935 != 0 {
		goto L32
	} else {
		goto L275
	}
L275:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v928)+24)) = int32(0)
	goto L272
L276:
	;
	v942 = *(*int32)(unsafe.Add(mBase, uint32(v939)+4))
	if v942 <= int32(0) {
		goto L2
	} else {
		goto L277
	}
L277:
	;
	v947 = int32(0)
	goto L278
L278:
	;
	v953 = *(*int32)(unsafe.Add(mBase, uint32(v939)+12))
	v957 = *(*int32)(unsafe.Add(mBase, uint32(v953+v947<<(uint(int32(2))%32))))
	if v957 == int32(0) {
		goto L280
	} else {
		goto L281
	}
L279:
	;
	goto L2
L280:
	;
	v969 = v947 + int32(1)
	v970 = *(*int32)(unsafe.Add(mBase, uint32(v939)+4))
	if v969 < v970 {
		v947 = v969
		goto L278
	} else {
		goto L284
	}
L281:
	;
	v960 = *(*int32)(unsafe.Add(mBase, uint32(v957)+24))
	if v960 == int32(0) {
		goto L280
	} else {
		goto L282
	}
L282:
	;
	F_SPI_freeplan(m, v960)
	mBase = m.M
	v964 = m.ExcPending
	if v964 != 0 {
		goto L32
	} else {
		goto L283
	}
L283:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v957)+24)) = int32(0)
	goto L280
L284:
	;
	goto L279
L285:
	;
	v975 = *(*int32)(unsafe.Add(mBase, uint32(v972)+24))
	if v975 == int32(0) {
		goto L2
	} else {
		goto L286
	}
L286:
	;
	F_SPI_freeplan(m, v975)
	mBase = m.M
	v979 = m.ExcPending
	if v979 != 0 {
		goto L32
	} else {
		goto L287
	}
L287:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v972)+24)) = int32(0)
	goto L2
L288:
	;
	v985 = *(*int32)(unsafe.Add(mBase, uint32(v982)+24))
	if v985 == int32(0) {
		goto L2
	} else {
		goto L289
	}
L289:
	;
	F_SPI_freeplan(m, v985)
	mBase = m.M
	v989 = m.ExcPending
	if v989 != 0 {
		goto L32
	} else {
		goto L290
	}
L290:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v982)+24)) = int32(0)
	goto L2
L291:
	;
	v995 = *(*int32)(unsafe.Add(mBase, uint32(v992)+24))
	if v995 == int32(0) {
		goto L2
	} else {
		goto L292
	}
L292:
	;
	F_SPI_freeplan(m, v995)
	mBase = m.M
	v999 = m.ExcPending
	if v999 != 0 {
		goto L32
	} else {
		goto L293
	}
L293:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v992)+24)) = int32(0)
	goto L2
L294:
	;
	v1018 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v1018
	F_errmsg_internal(m, int32(484312), v10)
	mBase = m.M
	v1023 = m.ExcPending
	if v1023 != 0 {
		goto L32
	} else {
		goto L295
	}
L295:
	;
	F_errfinish(m, int32(494883), int32(595), int32(301025))
	mBase = m.M
	v1030 = m.ExcPending
	if v1030 != 0 {
		goto L32
	} else {
		goto L296
	}
L296:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
