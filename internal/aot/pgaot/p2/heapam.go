package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_heapam_index_build_range_scan(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32) float64 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v30 float64
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
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
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v113 int64
	_ = v113
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v137 int32
	_ = v137
	var v143 int32
	_ = v143
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v160 int32
	_ = v160
	var v165 int32
	_ = v165
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v191 float64
	_ = v191
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v239 int32
	_ = v239
	var v245 int32
	_ = v245
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v259 int32
	_ = v259
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v273 int32
	_ = v273
	var v276 int32
	_ = v276
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v292 int32
	_ = v292
	var v323 int32
	_ = v323
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v337 int32
	_ = v337
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v363 int32
	_ = v363
	var v367 int32
	_ = v367
	var v371 int32
	_ = v371
	var v376 int32
	_ = v376
	var v381 int32
	_ = v381
	var v384 int32
	_ = v384
	var v387 int32
	_ = v387
	var v389 int32
	_ = v389
	var v391 int32
	_ = v391
	var v396 int32
	_ = v396
	var v398 int32
	_ = v398
	var v402 int32
	_ = v402
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v410 int32
	_ = v410
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v418 int32
	_ = v418
	var v428 int32
	_ = v428
	var v433 int32
	_ = v433
	var v435 int32
	_ = v435
	var v437 int32
	_ = v437
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v456 int32
	_ = v456
	var v464 int32
	_ = v464
	var v472 int32
	_ = v472
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v479 int32
	_ = v479
	var v487 int32
	_ = v487
	var v492 int32
	_ = v492
	var v495 int32
	_ = v495
	var v498 int32
	_ = v498
	var v501 int32
	_ = v501
	var v503 int32
	_ = v503
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v526 int32
	_ = v526
	var v530 int32
	_ = v530
	var v534 int32
	_ = v534
	var v539 int32
	_ = v539
	var v544 int32
	_ = v544
	var v547 int32
	_ = v547
	var v550 int32
	_ = v550
	var v552 int32
	_ = v552
	var v554 int32
	_ = v554
	var v559 int32
	_ = v559
	var v561 int32
	_ = v561
	var v565 int32
	_ = v565
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v573 int32
	_ = v573
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v581 int32
	_ = v581
	var v591 int32
	_ = v591
	var v596 int32
	_ = v596
	var v598 int32
	_ = v598
	var v600 int32
	_ = v600
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v615 int32
	_ = v615
	var v616 int32
	_ = v616
	var v619 int32
	_ = v619
	var v627 int32
	_ = v627
	var v635 int32
	_ = v635
	var v640 int32
	_ = v640
	var v641 int32
	_ = v641
	var v644 int32
	_ = v644
	var v652 int32
	_ = v652
	var v657 int32
	_ = v657
	var v660 int32
	_ = v660
	var v661 int32
	_ = v661
	var v666 int32
	_ = v666
	var v674 int32
	_ = v674
	var v677 int32
	_ = v677
	var v680 int32
	_ = v680
	var v682 int32
	_ = v682
	var v685 int32
	_ = v685
	var v686 int32
	_ = v686
	var v687 int32
	_ = v687
	var v692 int32
	_ = v692
	var v701 int32
	_ = v701
	var v705 int32
	_ = v705
	var v708 int32
	_ = v708
	var v712 int32
	_ = v712
	var v716 int32
	_ = v716
	var v721 int32
	_ = v721
	var v728 float64
	_ = v728
	var v736 int32
	_ = v736
	var v738 float64
	_ = v738
	var v739 int32
	_ = v739
	var v742 int32
	_ = v742
	var v746 int32
	_ = v746
	var v757 int32
	_ = v757
	var v779 float64
	_ = v779
	var v780 int32
	_ = v780
	var v782 int32
	_ = v782
	var v783 int32
	_ = v783
	var v785 int32
	_ = v785
	var v786 int32
	_ = v786
	var v787 int32
	_ = v787
	var v789 int32
	_ = v789
	var v793 int32
	_ = v793
	var v794 int32
	_ = v794
	var v795 int32
	_ = v795
	var v807 int32
	_ = v807
	var v808 int32
	_ = v808
	var v809 int32
	_ = v809
	var v812 int32
	_ = v812
	var v817 int32
	_ = v817
	var v818 int32
	_ = v818
	var v821 int32
	_ = v821
	var v825 int32
	_ = v825
	var v831 int32
	_ = v831
	var v833 int32
	_ = v833
	var v839 int32
	_ = v839
	var v842 int32
	_ = v842
	var v846 int32
	_ = v846
	var v847 int32
	_ = v847
	var v850 int32
	_ = v850
	var v851 int32
	_ = v851
	var v852 int32
	_ = v852
	var v860 int32
	_ = v860
	var v870 int32
	_ = v870
	var v878 int32
	_ = v878
	var v908 float64
	_ = v908
	var v909 int32
	_ = v909
	var v910 int32
	_ = v910
	var v940 float64
	_ = v940
	var v942 int32
	_ = v942
	var v947 int32
	_ = v947
	var v948 int64
	_ = v948
	var v951 int32
	_ = v951
	var v955 int32
	_ = v955
	var v958 int32
	_ = v958
	var v960 int32
	_ = v960
	var v961 int32
	_ = v961
	var v964 int32
	_ = v964
	var v972 int32
	_ = v972
	var v978 int32
	_ = v978
	var v983 int32
	_ = v983
	var v984 int32
	_ = v984
	var v985 int32
	_ = v985
	var v987 int32
	_ = v987
	var v989 int32
	_ = v989
	var v991 int32
	_ = v991
	var v993 int32
	_ = v993
	var v994 int32
	_ = v994
	var v1005 int32
	_ = v1005
	var v1008 int32
	_ = v1008
	var v1009 int32
	_ = v1009
	var v1010 int32
	_ = v1010
	var v1011 int32
	_ = v1011
	var v1022 int32
	_ = v1022
	var v1027 int32
	_ = v1027
	v12 = int32(0)
	v30 = float64(0)
	v31 = m.G0
	v33 = v31 - int32(800)
	m.G0 = v33
	v36 = int32(1)
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	if base.Ui32(v37) < base.Ui32(int32(12000)) {
		v46 = v36
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+116)))
	if v48 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	goto L1
L3:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+68))
	if v41 == int32(99) {
		v46 = v36
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v44 = F_isTempToastNamespace(m, v41)
	mBase = m.M
	v46 = v44
	goto L2
L5:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l2)+92))
	v54 = base.B2i32(v51 != int32(0))
	goto L7
L6:
	;
	v54 = int32(1)
	goto L7
L7:
	;
	v55 = F_CreateExecutorState(m)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	return float64(0)
L9:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v55)+152))
	if v59 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v62 = F_MakePerTupleExprContext(m, v55)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L8
	} else {
		goto L13
	}
L11:
	;
	v64 = v59
	goto L12
L12:
	;
	v66 = F_table_slot_create(m, l0, int32(0))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L8
	} else {
		goto L14
	}
L13:
	;
	v64 = v62
	goto L12
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v64)+4)) = v66
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l2)+84))
	v70 = F_ExecPrepareQual(m, v69, v55)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L8
	} else {
		goto L15
	}
L15:
	;
	v73 = *(*int32)(unsafe.Add(mBase, _consts[75]))
	if v73 == int32(0) {
		v79 = v12
		goto L16
	} else {
		goto L17
	}
L16:
	;
	if l10 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L17:
	;
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+121)))
	if v76 != 0 {
		v79 = v12
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v77 = F_GetOldestNonRemovableTransactionId(m, l0)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L8
	} else {
		goto L19
	}
L19:
	;
	v79 = v77
	goto L16
L20:
	;
	if l5 != 0 {
		goto L33
	} else {
		goto L34
	}
L21:
	;
	if v79 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L22:
	;
	goto L23
L23:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l10)+4))
	v103 = l10
	v104 = v102
	v105 = v12
	goto L20
L24:
	;
	v85 = F_GetTransactionSnapshot(m)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L8
	} else {
		goto L27
	}
L25:
	;
	v89 = int32(4173712)
	goto L26
L26:
	;
	v90 = int32(0)
	if l3 != 0 {
		goto L29
	} else {
		goto L30
	}
L27:
	;
	v87 = F_RegisterSnapshot(m, v85)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L8
	} else {
		goto L28
	}
L28:
	;
	v89 = v87
	goto L26
L29:
	;
	v97 = int32(449)
	goto L31
L30:
	;
	v97 = int32(321)
	goto L31
L31:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v98)+8))
	v100 = m.T0[v99].(func(*base.Module, int32, int32, int32, int32, int32, int32) int32)(m, l0, v89, v90, v90, v90, v97)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L8
	} else {
		goto L32
	}
L32:
	;
	v103 = v100
	v104 = v89
	v105 = base.B2i32(v79 == v90)
	goto L20
L33:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v103)+32))
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
	if l3 == int32(0) {
		goto L43
	} else {
		goto L44
	}
L36:
	;
	v112 = v107 + int32(20)
	goto L38
L37:
	;
	v112 = v103 + int32(36)
	goto L38
L38:
	;
	v113 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v112))))
	v116 = *(*int32)(unsafe.Add(mBase, _consts[26]))
	if v116 == int32(0) {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	goto L35
L40:
	;
	goto L39
L41:
	;
	v120 = int32(*(*uint8)(unsafe.Add(mBase, _consts[28])))
	if v120 != int32(1) {
		goto L40
	} else {
		goto L42
	}
L42:
	;
	v123 = int32(4509780)
	v125 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	v126 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[7])) = v125 + v126
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v116)))
	*(*int32)(unsafe.Add(mBase, uint32(v116))) = v129 + v126
	*(*int64)(unsafe.Add(mBase, uint32(v116+int32(120))+232)) = v113
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v116)))
	*(*int32)(unsafe.Add(mBase, uint32(v116))) = v137 + v126
	v143 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	*(*int32)(unsafe.Add(mBase, _consts[7])) = v143 - v126
	goto L40
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v103)+44)) = l7
	*(*int32)(unsafe.Add(mBase, uint32(v103)+40)) = l6
	goto L45
L44:
	;
	goto L45
L45:
	;
	v152 = F_heap_getnext(m, v103)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L8
	} else {
		goto L47
	}
L46:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1005 = m.ExcPending
	if v1005 != 0 {
		goto L8
	} else {
		goto L278
	}
L47:
	;
	if v152 != 0 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v160 = int32(-1)
	v165 = v152
	v182 = v160
	v184 = v160
	v191 = v30
	goto L51
L49:
	;
	v940 = v30
	goto L50
L50:
	;
	if l5 != 0 {
		goto L261
	} else {
		goto L262
	}
L51:
	;
	v193 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	if v193 != 0 {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	v940 = v908
	goto L50
L53:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L8
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v103)+52))
	if l5 == int32(0) {
		v251 = v196
		v253 = v184
		goto L57
	} else {
		goto L58
	}
L56:
	;
	goto L55
L57:
	;
	if v251 != v182 {
		goto L74
	} else {
		goto L75
	}
L58:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v103)+32))
	if v199 != 0 {
		goto L60
	} else {
		goto L61
	}
L59:
	;
	if v184 == v212 {
		v251 = v196
		v253 = v184
		goto L57
	} else {
		goto L69
	}
L60:
	;
	v202 = v199 + int32(28)
	goto L62
L61:
	;
	v202 = v103 + int32(40)
	goto L62
L62:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v202)))
	if base.Ui32(v203) < base.Ui32(v196) {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v212 = v196 - v203
	goto L59
L64:
	;
	goto L65
L65:
	;
	if v199 != 0 {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v208 = v199 + int32(20)
	goto L68
L67:
	;
	v208 = v103 + int32(36)
	goto L68
L68:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v208)))
	v212 = v209 + (v196 - v203)
	goto L59
L69:
	;
	v218 = *(*int32)(unsafe.Add(mBase, _consts[26]))
	if v218 == int32(0) {
		goto L71
	} else {
		goto L72
	}
L70:
	;
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v103)+52))
	v251 = v249
	v253 = v212
	goto L57
L71:
	;
	goto L70
L72:
	;
	v222 = int32(*(*uint8)(unsafe.Add(mBase, _consts[28])))
	if v222 != int32(1) {
		goto L71
	} else {
		goto L73
	}
L73:
	;
	v225 = int32(4509780)
	v227 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	v228 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[7])) = v227 + v228
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v218)))
	*(*int32)(unsafe.Add(mBase, uint32(v218))) = v231 + v228
	*(*int64)(unsafe.Add(mBase, uint32(v218+int32(128))+232)) = base.I64_extend_i32_u(v212)
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v218)))
	*(*int32)(unsafe.Add(mBase, uint32(v218))) = v239 + v228
	v245 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	*(*int32)(unsafe.Add(mBase, _consts[7])) = v245 - v228
	goto L71
L74:
	;
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v103)+56))
	if v255 < int32(0) {
		goto L78
	} else {
		goto L79
	}
L75:
	;
	v288 = v182
	goto L76
L76:
	;
	if base.B2i32(v104 != int32(4173712)) == int32(0) {
		goto L86
	} else {
		goto L87
	}
L77:
	;
	F_LockBuffer(m, v255, int32(1))
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L8
	} else {
		goto L81
	}
L78:
	;
	v259 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v259+(v255^int32(-1))<<(uint(int32(2))%32))))
	v273 = v265
	goto L77
L79:
	;
	goto L80
L80:
	;
	v267 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v273 = v267 + v255<<(uint(int32(13))%32) + int32(-8192)
	goto L77
L81:
	;
	F_heap_get_root_tuples(m, v273, v33+int32(48))
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L8
	} else {
		goto L82
	}
L82:
	;
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v103)+56))
	F_LockBuffer(m, v281, int32(0))
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L8
	} else {
		goto L83
	}
L83:
	;
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v103)+52))
	v288 = v285
	goto L76
L84:
	;
	v909 = F_heap_getnext(m, v103)
	mBase = m.M
	v910 = m.ExcPending
	if v910 != 0 {
		goto L8
	} else {
		goto L259
	}
L85:
	;
	v780 = *(*int32)(unsafe.Add(mBase, uint32(v64)+20))
	F_MemoryContextReset(m, v780)
	mBase = m.M
	v782 = m.ExcPending
	if v782 != 0 {
		goto L8
	} else {
		goto L235
	}
L86:
	;
	v292 = v165 + int32(4)
	goto L89
L87:
	;
	goto L88
L88:
	;
	v757 = int32(1)
	v779 = base.F64_add(v191, float64(1))
	goto L85
L89:
	;
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v103)+56))
	F_LockBuffer(m, v323, int32(1))
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L8
	} else {
		goto L91
	}
L91:
	;
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v103)+56))
	v328 = F_HeapTupleSatisfiesVacuum(m, v165, v79, v327)
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L8
	} else {
		goto L103
	}
L92:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v746 = m.ExcPending
	if v746 != 0 {
		goto L8
	} else {
		goto L234
	}
L93:
	;
	v739 = *(*int32)(unsafe.Add(mBase, uint32(v103)+56))
	F_LockBuffer(m, v739, int32(0))
	mBase = m.M
	v742 = m.ExcPending
	if v742 != 0 {
		goto L8
	} else {
		goto L233
	}
L94:
	;
	v736 = int32(0)
	v738 = base.F64_add(v191, float64(1))
	goto L93
L95:
	;
	v736 = int32(1)
	v738 = v728
	goto L93
L96:
	;
	v728 = base.F64_add(v191, float64(1))
	goto L95
L97:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v712 = m.ExcPending
	if v712 != 0 {
		goto L8
	} else {
		goto L230
	}
L98:
	;
	v705 = *(*int32)(unsafe.Add(mBase, uint32(v103)+56))
	F_LockBuffer(m, v705, int32(0))
	mBase = m.M
	v708 = m.ExcPending
	if v708 != 0 {
		goto L8
	} else {
		goto L229
	}
L99:
	;
	v701 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l2)+122)) = uint8(v701)
	goto L98
L100:
	;
	if l4 != 0 {
		goto L94
	} else {
		goto L162
	}
L101:
	;
	if l4 != 0 {
		goto L96
	} else {
		goto L107
	}
L102:
	;
	v330 = int32(0)
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v165)+16))
	v332 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v331)+19)))
	if v332&int32(64) == v330 {
		v736 = v330
		v738 = v191
		goto L93
	} else {
		goto L104
	}
L103:
	;
	switch v328 {
	case 0:
		goto L98
	case 1:
		goto L96
	case 2:
		goto L102
	case 3:
		goto L101
	case 4:
		goto L100
	default:
		goto L97
	}
L104:
	;
	v337 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v331)+20)))
	if v337&int32(2048) != 0 {
		v736 = v330
		v738 = v191
		goto L93
	} else {
		goto L105
	}
L105:
	;
	if v337&int32(768) != int32(512) {
		goto L99
	} else {
		goto L106
	}
L106:
	;
	v736 = v330
	v738 = v191
	goto L93
L107:
	;
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v165)+16))
	v346 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v345)+20)))
	v347 = int32(768)
	if v346&v347 != v347 {
		goto L108
	} else {
		goto L109
	}
L108:
	;
	v351 = *(*int32)(unsafe.Add(mBase, uint32(v345)))
	v352 = v351
	goto L110
L109:
	;
	v352 = int32(2)
	goto L110
L110:
	;
	if base.Ui32(v352) < base.Ui32(int32(3)) {
		goto L112
	} else {
		goto L113
	}
L111:
	;
	if v472 != 0 {
		goto L96
	} else {
		goto L151
	}
L112:
	;
	v472 = int32(0)
	goto L111
L113:
	;
	goto L114
L114:
	;
	v363 = *(*int32)(unsafe.Add(mBase, _consts[67]))
	if v363 == v352 {
		goto L115
	} else {
		goto L116
	}
L115:
	;
	v472 = int32(1)
	goto L111
L116:
	;
	goto L117
L117:
	;
	v367 = *(*int32)(unsafe.Add(mBase, _consts[68]))
	if v367 <= int32(0) {
		goto L119
	} else {
		goto L120
	}
L118:
	;
	v472 = v464
	goto L111
L119:
	;
	v371 = *(*int32)(unsafe.Add(mBase, _consts[65]))
	if v371 == int32(0) {
		v464 = int32(0)
		goto L118
	} else {
		goto L122
	}
L120:
	;
	goto L121
L121:
	;
	v433 = *(*int32)(unsafe.Add(mBase, _consts[69]))
	v435 = int32(0)
	v437 = v367 - int32(1)
	goto L141
L122:
	;
	v376 = v371
	goto L123
L123:
	;
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v376)+20))
	if v381 == int32(4) {
		goto L125
	} else {
		goto L126
	}
L124:
	;
	v464 = int32(0)
	goto L118
L125:
	;
	v428 = *(*int32)(unsafe.Add(mBase, uint32(v376)+80))
	if v428 != 0 {
		v376 = v428
		goto L123
	} else {
		goto L140
	}
L126:
	;
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v376)))
	if v384 == int32(0) {
		goto L125
	} else {
		goto L127
	}
L127:
	;
	v387 = int32(1)
	if v352 == v384 {
		v464 = v387
		goto L118
	} else {
		goto L128
	}
L128:
	;
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v376)+52))
	v391 = v389 - int32(1)
	if v391 < int32(0) {
		goto L125
	} else {
		goto L129
	}
L129:
	;
	v396 = int32(0)
	v398 = v391
	goto L130
L130:
	;
	v402 = *(*int32)(unsafe.Add(mBase, uint32(v376)+48))
	v404 = int32(2)
	v405 = base.I32_div_s(v398-v396, v404)
	v406 = v405 + v396
	v410 = *(*int32)(unsafe.Add(mBase, uint32(v402+v406<<(uint(v404)%32))))
	if v410 == v352 {
		v464 = v387
		goto L118
	} else {
		goto L132
	}
L131:
	;
	goto L125
L132:
	;
	v414 = F_TransactionIdPrecedes(m, v410, v352)
	mBase = m.M
	if v414 != 0 {
		goto L133
	} else {
		goto L134
	}
L133:
	;
	v415 = v406 + int32(1)
	goto L135
L134:
	;
	v415 = v396
	goto L135
L135:
	;
	if v414 != 0 {
		goto L136
	} else {
		goto L137
	}
L136:
	;
	v418 = v398
	goto L138
L137:
	;
	v418 = v406 - int32(1)
	goto L138
L138:
	;
	if v415 <= v418 {
		v396 = v415
		v398 = v418
		goto L130
	} else {
		goto L139
	}
L139:
	;
	goto L131
L140:
	;
	goto L124
L141:
	;
	v442 = int32(2)
	v443 = base.I32_div_s(v437-v435, v442)
	v444 = v443 + v435
	v448 = *(*int32)(unsafe.Add(mBase, uint32(v433+v444<<(uint(v442)%32))))
	v449 = base.B2i32(v448 == v352)
	if v448 == v352 {
		v464 = v449
		goto L118
	} else {
		goto L143
	}
L142:
	;
	v464 = v449
	goto L118
L143:
	;
	v452 = base.B2i32(base.Ui32(v448) < base.Ui32(v352))
	if base.Ui32(v448) < base.Ui32(v352) {
		goto L144
	} else {
		goto L145
	}
L144:
	;
	v453 = v444 + int32(1)
	goto L146
L145:
	;
	v453 = v435
	goto L146
L146:
	;
	if base.Ui32(v448) < base.Ui32(v352) {
		goto L147
	} else {
		goto L148
	}
L147:
	;
	v456 = v437
	goto L149
L148:
	;
	v456 = v444 - int32(1)
	goto L149
L149:
	;
	if v453 <= v456 {
		v435 = v453
		v437 = v456
		goto L141
	} else {
		goto L150
	}
L150:
	;
	goto L142
L151:
	;
	if v46 != 0 {
		goto L152
	} else {
		goto L153
	}
L152:
	;
	if v54 == int32(0) {
		v728 = v191
		goto L95
	} else {
		goto L158
	}
L153:
	;
	v475 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v476 = m.ExcPending
	if v476 != 0 {
		goto L8
	} else {
		goto L154
	}
L154:
	;
	if v475 == int32(0) {
		goto L152
	} else {
		goto L155
	}
L155:
	;
	v479 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v33)+16)) = v479 + int32(4)
	F_errmsg_internal(m, int32(718605), v33+int32(16))
	mBase = m.M
	v487 = m.ExcPending
	if v487 != 0 {
		goto L8
	} else {
		goto L156
	}
L156:
	;
	F_errfinish(m, int32(495072), int32(1484), int32(284533))
	mBase = m.M
	v492 = m.ExcPending
	if v492 != 0 {
		goto L8
	} else {
		goto L157
	}
L157:
	;
	goto L152
L158:
	;
	v495 = *(*int32)(unsafe.Add(mBase, uint32(v103)+56))
	F_LockBuffer(m, v495, int32(0))
	mBase = m.M
	v498 = m.ExcPending
	if v498 != 0 {
		goto L8
	} else {
		goto L159
	}
L159:
	;
	F_XactLockTableWait(m, v352, l0, v292, int32(6))
	mBase = m.M
	v501 = m.ExcPending
	if v501 != 0 {
		goto L8
	} else {
		goto L160
	}
L160:
	;
	v503 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	if v503 == int32(0) {
		goto L89
	} else {
		goto L161
	}
L161:
	;
	goto L92
L162:
	;
	v506 = *(*int32)(unsafe.Add(mBase, uint32(v165)+16))
	v507 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v506)+20)))
	if v507&int32(6272) == int32(4096) {
		goto L164
	} else {
		goto L165
	}
L163:
	;
	if base.Ui32(v515) < base.Ui32(int32(3)) {
		goto L169
	} else {
		goto L170
	}
L164:
	;
	v512 = F_HeapTupleGetUpdateXid(m, v506)
	mBase = m.M
	v513 = m.ExcPending
	if v513 != 0 {
		goto L8
	} else {
		goto L167
	}
L165:
	;
	goto L166
L166:
	;
	v514 = *(*int32)(unsafe.Add(mBase, uint32(v506)+4))
	v515 = v514
	goto L163
L167:
	;
	v515 = v512
	goto L163
L168:
	;
	if v635 == int32(0) {
		goto L208
	} else {
		goto L209
	}
L169:
	;
	v635 = int32(0)
	goto L168
L170:
	;
	goto L171
L171:
	;
	v526 = *(*int32)(unsafe.Add(mBase, _consts[67]))
	if v526 == v515 {
		goto L172
	} else {
		goto L173
	}
L172:
	;
	v635 = int32(1)
	goto L168
L173:
	;
	goto L174
L174:
	;
	v530 = *(*int32)(unsafe.Add(mBase, _consts[68]))
	if v530 <= int32(0) {
		goto L176
	} else {
		goto L177
	}
L175:
	;
	v635 = v627
	goto L168
L176:
	;
	v534 = *(*int32)(unsafe.Add(mBase, _consts[65]))
	if v534 == int32(0) {
		v627 = int32(0)
		goto L175
	} else {
		goto L179
	}
L177:
	;
	goto L178
L178:
	;
	v596 = *(*int32)(unsafe.Add(mBase, _consts[69]))
	v598 = int32(0)
	v600 = v530 - int32(1)
	goto L198
L179:
	;
	v539 = v534
	goto L180
L180:
	;
	v544 = *(*int32)(unsafe.Add(mBase, uint32(v539)+20))
	if v544 == int32(4) {
		goto L182
	} else {
		goto L183
	}
L181:
	;
	v627 = int32(0)
	goto L175
L182:
	;
	v591 = *(*int32)(unsafe.Add(mBase, uint32(v539)+80))
	if v591 != 0 {
		v539 = v591
		goto L180
	} else {
		goto L197
	}
L183:
	;
	v547 = *(*int32)(unsafe.Add(mBase, uint32(v539)))
	if v547 == int32(0) {
		goto L182
	} else {
		goto L184
	}
L184:
	;
	v550 = int32(1)
	if v515 == v547 {
		v627 = v550
		goto L175
	} else {
		goto L185
	}
L185:
	;
	v552 = *(*int32)(unsafe.Add(mBase, uint32(v539)+52))
	v554 = v552 - int32(1)
	if v554 < int32(0) {
		goto L182
	} else {
		goto L186
	}
L186:
	;
	v559 = int32(0)
	v561 = v554
	goto L187
L187:
	;
	v565 = *(*int32)(unsafe.Add(mBase, uint32(v539)+48))
	v567 = int32(2)
	v568 = base.I32_div_s(v561-v559, v567)
	v569 = v568 + v559
	v573 = *(*int32)(unsafe.Add(mBase, uint32(v565+v569<<(uint(v567)%32))))
	if v573 == v515 {
		v627 = v550
		goto L175
	} else {
		goto L189
	}
L188:
	;
	goto L182
L189:
	;
	v577 = F_TransactionIdPrecedes(m, v573, v515)
	mBase = m.M
	if v577 != 0 {
		goto L190
	} else {
		goto L191
	}
L190:
	;
	v578 = v569 + int32(1)
	goto L192
L191:
	;
	v578 = v559
	goto L192
L192:
	;
	if v577 != 0 {
		goto L193
	} else {
		goto L194
	}
L193:
	;
	v581 = v561
	goto L195
L194:
	;
	v581 = v569 - int32(1)
	goto L195
L195:
	;
	if v578 <= v581 {
		v559 = v578
		v561 = v581
		goto L187
	} else {
		goto L196
	}
L196:
	;
	goto L188
L197:
	;
	goto L181
L198:
	;
	v605 = int32(2)
	v606 = base.I32_div_s(v600-v598, v605)
	v607 = v606 + v598
	v611 = *(*int32)(unsafe.Add(mBase, uint32(v596+v607<<(uint(v605)%32))))
	v612 = base.B2i32(v611 == v515)
	if v611 == v515 {
		v627 = v612
		goto L175
	} else {
		goto L200
	}
L199:
	;
	v627 = v612
	goto L175
L200:
	;
	v615 = base.B2i32(base.Ui32(v611) < base.Ui32(v515))
	if base.Ui32(v611) < base.Ui32(v515) {
		goto L201
	} else {
		goto L202
	}
L201:
	;
	v616 = v607 + int32(1)
	goto L203
L202:
	;
	v616 = v598
	goto L203
L203:
	;
	if base.Ui32(v611) < base.Ui32(v515) {
		goto L204
	} else {
		goto L205
	}
L204:
	;
	v619 = v600
	goto L206
L205:
	;
	v619 = v607 - int32(1)
	goto L206
L206:
	;
	if v616 <= v619 {
		v598 = v616
		v600 = v619
		goto L198
	} else {
		goto L207
	}
L207:
	;
	goto L199
L208:
	;
	if v46 != 0 {
		goto L211
	} else {
		goto L212
	}
L209:
	;
	goto L210
L210:
	;
	v685 = int32(0)
	v686 = *(*int32)(unsafe.Add(mBase, uint32(v165)+16))
	v687 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v686)+19)))
	if v687&int32(64) == v685 {
		v736 = v685
		v738 = v191
		goto L93
	} else {
		goto L226
	}
L211:
	;
	if v54 == int32(0) {
		goto L217
	} else {
		goto L218
	}
L212:
	;
	v640 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v641 = m.ExcPending
	if v641 != 0 {
		goto L8
	} else {
		goto L213
	}
L213:
	;
	if v640 == int32(0) {
		goto L211
	} else {
		goto L214
	}
L214:
	;
	v644 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v33)+32)) = v644 + int32(4)
	F_errmsg_internal(m, int32(718653), v33+int32(32))
	mBase = m.M
	v652 = m.ExcPending
	if v652 != 0 {
		goto L8
	} else {
		goto L215
	}
L215:
	;
	F_errfinish(m, int32(495072), int32(1543), int32(284533))
	mBase = m.M
	v657 = m.ExcPending
	if v657 != 0 {
		goto L8
	} else {
		goto L216
	}
L216:
	;
	goto L211
L217:
	;
	v660 = *(*int32)(unsafe.Add(mBase, uint32(v165)+16))
	v661 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v660)+19)))
	if v661&int32(64) == int32(0) {
		goto L94
	} else {
		goto L220
	}
L218:
	;
	goto L219
L219:
	;
	v674 = *(*int32)(unsafe.Add(mBase, uint32(v103)+56))
	F_LockBuffer(m, v674, int32(0))
	mBase = m.M
	v677 = m.ExcPending
	if v677 != 0 {
		goto L8
	} else {
		goto L223
	}
L220:
	;
	v666 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v660)+20)))
	if v666&int32(2048) != 0 {
		goto L94
	} else {
		goto L221
	}
L221:
	;
	if v666&int32(768) == int32(512) {
		goto L94
	} else {
		goto L222
	}
L222:
	;
	goto L219
L223:
	;
	F_XactLockTableWait(m, v515, l0, v292, int32(6))
	mBase = m.M
	v680 = m.ExcPending
	if v680 != 0 {
		goto L8
	} else {
		goto L224
	}
L224:
	;
	v682 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	if v682 == int32(0) {
		goto L89
	} else {
		goto L225
	}
L225:
	;
	goto L92
L226:
	;
	v692 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v686)+20)))
	if v692&int32(2048) != 0 {
		v736 = v685
		v738 = v191
		goto L93
	} else {
		goto L227
	}
L227:
	;
	if v692&int32(768) == int32(512) {
		v736 = v685
		v738 = v191
		goto L93
	} else {
		goto L228
	}
L228:
	;
	goto L99
L229:
	;
	v908 = v191
	goto L84
L230:
	;
	F_errmsg_internal(m, int32(97885), int32(0))
	mBase = m.M
	v716 = m.ExcPending
	if v716 != 0 {
		goto L8
	} else {
		goto L231
	}
L231:
	;
	F_errfinish(m, int32(495072), int32(1613), int32(284533))
	mBase = m.M
	v721 = m.ExcPending
	if v721 != 0 {
		goto L8
	} else {
		goto L232
	}
L232:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L233:
	;
	v757 = v736
	v779 = v738
	goto L85
L234:
	;
	goto L89
L235:
	;
	v783 = *(*int32)(unsafe.Add(mBase, uint32(v103)+56))
	F_ExecStoreBufferHeapTuple(m, v165, v66, v783)
	mBase = m.M
	v785 = m.ExcPending
	if v785 != 0 {
		goto L8
	} else {
		goto L236
	}
L236:
	;
	if v70 != 0 {
		goto L237
	} else {
		goto L238
	}
L237:
	;
	v786 = int32(4515120)
	v787 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v789 = *(*int32)(unsafe.Add(mBase, uint32(v64)+20))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v789
	v793 = *(*int32)(unsafe.Add(mBase, uint32(v70)+20))
	v794 = m.T0[v793].(func(*base.Module, int32, int32, int32) int32)(m, v70, v64, v33+int32(42))
	mBase = m.M
	v795 = m.ExcPending
	if v795 != 0 {
		goto L8
	} else {
		goto L240
	}
L238:
	;
	goto L239
L239:
	;
	F_FormIndexDatum(m, l2, v66, v55, v33+int32(672), v33+int32(640))
	mBase = m.M
	v807 = m.ExcPending
	if v807 != 0 {
		goto L8
	} else {
		goto L242
	}
L240:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v787
	if v794 == int32(0) {
		v908 = v779
		goto L84
	} else {
		goto L241
	}
L241:
	;
	goto L239
L242:
	;
	v808 = *(*int32)(unsafe.Add(mBase, uint32(v165)+16))
	v809 = int32(*(*int16)(unsafe.Add(mBase, uint32(v808)+18)))
	if v809 < int32(0) {
		goto L243
	} else {
		goto L244
	}
L243:
	;
	v812 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v165)+8)))
	v817 = v812<<(uint(int32(1))%32) + v33 + int32(46)
	v818 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v817))))
	if v818 == int32(0) {
		goto L246
	} else {
		goto L247
	}
L244:
	;
	goto L245
L245:
	;
	m.T0[l8].(func(*base.Module, int32, int32, int32, int32, int32, int32))(m, l1, v165+int32(4), v33+int32(672), v33+int32(640), v757, l9)
	mBase = m.M
	v878 = m.ExcPending
	if v878 != 0 {
		goto L8
	} else {
		goto L258
	}
L246:
	;
	v821 = *(*int32)(unsafe.Add(mBase, uint32(v103)+56))
	if v821 < int32(0) {
		goto L250
	} else {
		goto L251
	}
L247:
	;
	v852 = v818
	goto L248
L248:
	;
	if base.Ui32(int32(2048)) <= base.Ui32((v852-int32(1))&int32(65535)) {
		goto L46
	} else {
		goto L256
	}
L249:
	;
	F_LockBuffer(m, v821, int32(1))
	mBase = m.M
	v842 = m.ExcPending
	if v842 != 0 {
		goto L8
	} else {
		goto L253
	}
L250:
	;
	v825 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v831 = *(*int32)(unsafe.Add(mBase, uint32(v825+(v821^int32(-1))<<(uint(int32(2))%32))))
	v839 = v831
	goto L249
L251:
	;
	goto L252
L252:
	;
	v833 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v839 = v833 + v821<<(uint(int32(13))%32) + int32(-8192)
	goto L249
L253:
	;
	F_heap_get_root_tuples(m, v839, v33+int32(48))
	mBase = m.M
	v846 = m.ExcPending
	if v846 != 0 {
		goto L8
	} else {
		goto L254
	}
L254:
	;
	v847 = *(*int32)(unsafe.Add(mBase, uint32(v103)+56))
	F_LockBuffer(m, v847, int32(0))
	mBase = m.M
	v850 = m.ExcPending
	if v850 != 0 {
		goto L8
	} else {
		goto L255
	}
L255:
	;
	v851 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v817))))
	v852 = v851
	goto L248
L256:
	;
	v860 = *(*int32)(unsafe.Add(mBase, uint32(v165)+4))
	*(*uint16)(unsafe.Add(mBase, uint32(v33)+46)) = uint16(v852)
	*(*int32)(unsafe.Add(mBase, uint32(v33)+42)) = v860
	m.T0[l8].(func(*base.Module, int32, int32, int32, int32, int32, int32))(m, l1, v33+int32(42), v33+int32(672), v33+int32(640), v757, l9)
	mBase = m.M
	v870 = m.ExcPending
	if v870 != 0 {
		goto L8
	} else {
		goto L257
	}
L257:
	;
	v908 = v779
	goto L84
L258:
	;
	v908 = v779
	goto L84
L259:
	;
	if v909 != 0 {
		v165 = v909
		v182 = v288
		v184 = v253
		v191 = v908
		goto L51
	} else {
		goto L260
	}
L260:
	;
	goto L52
L261:
	;
	v942 = *(*int32)(unsafe.Add(mBase, uint32(v103)+32))
	if v942 != 0 {
		goto L264
	} else {
		goto L265
	}
L262:
	;
	goto L263
L263:
	;
	v983 = *(*int32)(unsafe.Add(mBase, uint32(v103)))
	v984 = *(*int32)(unsafe.Add(mBase, uint32(v983)+188))
	v985 = *(*int32)(unsafe.Add(mBase, uint32(v984)+12))
	m.T0[v985].(func(*base.Module, int32))(m, v103)
	mBase = m.M
	v987 = m.ExcPending
	if v987 != 0 {
		goto L8
	} else {
		goto L271
	}
L264:
	;
	v947 = v942 + int32(20)
	goto L266
L265:
	;
	v947 = v103 + int32(36)
	goto L266
L266:
	;
	v948 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v947))))
	v951 = *(*int32)(unsafe.Add(mBase, _consts[26]))
	if v951 == int32(0) {
		goto L268
	} else {
		goto L269
	}
L267:
	;
	goto L263
L268:
	;
	goto L267
L269:
	;
	v955 = int32(*(*uint8)(unsafe.Add(mBase, _consts[28])))
	if v955 != int32(1) {
		goto L268
	} else {
		goto L270
	}
L270:
	;
	v958 = int32(4509780)
	v960 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	v961 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[7])) = v960 + v961
	v964 = *(*int32)(unsafe.Add(mBase, uint32(v951)))
	*(*int32)(unsafe.Add(mBase, uint32(v951))) = v964 + v961
	*(*int64)(unsafe.Add(mBase, uint32(v951+int32(128))+232)) = v948
	v972 = *(*int32)(unsafe.Add(mBase, uint32(v951)))
	*(*int32)(unsafe.Add(mBase, uint32(v951))) = v972 + v961
	v978 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	*(*int32)(unsafe.Add(mBase, _consts[7])) = v978 - v961
	goto L268
L271:
	;
	if v105 != 0 {
		goto L272
	} else {
		goto L273
	}
L272:
	;
	F_UnregisterSnapshot(m, v104)
	mBase = m.M
	v989 = m.ExcPending
	if v989 != 0 {
		goto L8
	} else {
		goto L275
	}
L273:
	;
	goto L274
L274:
	;
	F_ExecDropSingleTupleTableSlot(m, v66)
	mBase = m.M
	v991 = m.ExcPending
	if v991 != 0 {
		goto L8
	} else {
		goto L276
	}
L275:
	;
	goto L274
L276:
	;
	F_FreeExecutorState(m, v55)
	mBase = m.M
	v993 = m.ExcPending
	if v993 != 0 {
		goto L8
	} else {
		goto L277
	}
L277:
	;
	v994 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l2)+88)) = v994
	*(*int32)(unsafe.Add(mBase, uint32(l2)+80)) = v994
	m.G0 = v33 + int32(800)
	return v940
L278:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v1008 = m.ExcPending
	if v1008 != 0 {
		goto L8
	} else {
		goto L279
	}
L279:
	;
	v1009 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v165)+6)))
	v1010 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v165)+4)))
	v1011 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v33)+4)) = v812
	*(*int32)(unsafe.Add(mBase, uint32(v33)+8)) = v1011 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v33))) = v1009 | v1010<<(uint(int32(16))%32)
	F_errmsg_internal(m, int32(718771), v33)
	mBase = m.M
	v1022 = m.ExcPending
	if v1022 != 0 {
		goto L8
	} else {
		goto L280
	}
L280:
	;
	F_errfinish(m, int32(495072), int32(1693), int32(284533))
	mBase = m.M
	v1027 = m.ExcPending
	if v1027 != 0 {
		goto L8
	} else {
		goto L281
	}
L281:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_heapam_index_fetch_begin(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = F_palloc0(m, int32(8))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v4)+4)) = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v4))) = l0
		return v4
	}
}
func F_heapam_index_validate_scan(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v18 int32
	_ = v18
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
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int64
	_ = v47
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v106 float64
	_ = v106
	var v110 int32
	_ = v110
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v142 int32
	_ = v142
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v216 int32
	_ = v216
	var v222 int32
	_ = v222
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v272 int64
	_ = v272
	var v275 int64
	_ = v275
	var v278 int64
	_ = v278
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v313 int32
	_ = v313
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v324 int32
	_ = v324
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v351 int32
	_ = v351
	var v352 int64
	_ = v352
	var v355 int64
	_ = v355
	var v358 int64
	_ = v358
	var v387 int32
	_ = v387
	var v399 int32
	_ = v399
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v417 int32
	_ = v417
	var v422 int32
	_ = v422
	var v432 int32
	_ = v432
	var v440 int32
	_ = v440
	var v442 int32
	_ = v442
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v455 int32
	_ = v455
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v466 int32
	_ = v466
	var v478 int32
	_ = v478
	var v484 int32
	_ = v484
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v492 int32
	_ = v492
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v499 int32
	_ = v499
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v517 int32
	_ = v517
	var v524 int32
	_ = v524
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v528 float64
	_ = v528
	var v541 int32
	_ = v541
	var v547 int32
	_ = v547
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v572 int32
	_ = v572
	var v574 int32
	_ = v574
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	v6 = int32(0)
	v18 = m.G0
	v20 = v18 - int32(1088)
	m.G0 = v20
	v22 = F_CreateExecutorState(m)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v22)+152))
	if v24 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v27 = F_MakePerTupleExprContext(m, v22)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L6
	}
L4:
	;
	v29 = v24
	goto L5
L5:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v32 = F_MakeSingleTupleTableSlot(m, v30, int32(1617464))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L7
	}
L6:
	;
	v29 = v27
	goto L5
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+4)) = v32
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l2)+84))
	v36 = F_ExecPrepareQual(m, v35, v22)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v39 = int32(0)
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+8))
	v45 = m.T0[v44].(func(*base.Module, int32, int32, int32, int32, int32, int32) int32)(m, l0, l3, v39, v39, v39, int32(321))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	v47 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v45)+36)))
	v50 = *(*int32)(unsafe.Add(mBase, _consts[26]))
	if v50 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v81 = F_heap_getnext(m, v45)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L1
	} else {
		goto L14
	}
L11:
	;
	goto L10
L12:
	;
	v54 = int32(*(*uint8)(unsafe.Add(mBase, _consts[28])))
	if v54 != int32(1) {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v57 = int32(4509780)
	v59 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	v60 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[7])) = v59 + v60
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
	*(*int32)(unsafe.Add(mBase, uint32(v50))) = v63 + v60
	*(*int64)(unsafe.Add(mBase, uint32(v50+int32(120))+232)) = v47
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
	*(*int32)(unsafe.Add(mBase, uint32(v50))) = v71 + v60
	v77 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	*(*int32)(unsafe.Add(mBase, _consts[7])) = v77 - v60
	goto L11
L14:
	;
	if v81 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v83 = int32(-1)
	v88 = v81
	v93 = v83
	v94 = v6
	v97 = v83
	v100 = v6
	goto L18
L16:
	;
	goto L17
L17:
	;
	v568 = *(*int32)(unsafe.Add(mBase, uint32(v45)))
	v569 = *(*int32)(unsafe.Add(mBase, uint32(v568)+188))
	v570 = *(*int32)(unsafe.Add(mBase, uint32(v569)+12))
	m.T0[v570].(func(*base.Module, int32))(m, v45)
	mBase = m.M
	v572 = m.ExcPending
	if v572 != 0 {
		goto L1
	} else {
		goto L102
	}
L18:
	;
	v103 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	if v103 != 0 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	goto L17
L20:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L1
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v106 = *(*float64)(unsafe.Add(mBase, uint32(l4)+8))
	*(*float64)(unsafe.Add(mBase, uint32(l4)+8)) = base.F64_add(v106, float64(1))
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v45)+52))
	if base.B2i32(v110 == v97)&base.B2i32(v97 != int32(-1)) == int32(0) {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	goto L22
L24:
	;
	v121 = *(*int32)(unsafe.Add(mBase, _consts[26]))
	if v121 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L25:
	;
	v153 = v97
	goto L26
L26:
	;
	if v93 != v153 {
		goto L31
	} else {
		goto L32
	}
L27:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v45)+52))
	v153 = v152
	goto L26
L28:
	;
	goto L27
L29:
	;
	v125 = int32(*(*uint8)(unsafe.Add(mBase, _consts[28])))
	if v125 != int32(1) {
		goto L28
	} else {
		goto L30
	}
L30:
	;
	v128 = int32(4509780)
	v130 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	v131 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[7])) = v130 + v131
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v121)))
	*(*int32)(unsafe.Add(mBase, uint32(v121))) = v134 + v131
	*(*int64)(unsafe.Add(mBase, uint32(v121+int32(128))+232)) = base.I64_extend_i32_u(v110)
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v121)))
	*(*int32)(unsafe.Add(mBase, uint32(v121))) = v142 + v131
	v148 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	*(*int32)(unsafe.Add(mBase, _consts[7])) = v148 - v131
	goto L28
L31:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v45)+56))
	if v155 < int32(0) {
		goto L35
	} else {
		goto L36
	}
L32:
	;
	v193 = v93
	goto L33
L33:
	;
	v195 = v88 + int32(4)
	v196 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v195)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v20)+24)) = uint16(v196)
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v195)))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+20)) = v198
	v200 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v88)+8)))
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v88)+16))
	v202 = int32(*(*int16)(unsafe.Add(mBase, uint32(v201)+18)))
	if v202 < int32(0) {
		goto L46
	} else {
		goto L47
	}
L34:
	;
	F_LockBuffer(m, v155, int32(1))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L1
	} else {
		goto L38
	}
L35:
	;
	v159 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v159+(v155^int32(-1))<<(uint(int32(2))%32))))
	v173 = v165
	goto L34
L36:
	;
	goto L37
L37:
	;
	v167 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v173 = v167 + v155<<(uint(int32(13))%32) + int32(-8192)
	goto L34
L38:
	;
	F_heap_get_root_tuples(m, v173, v20+int32(336))
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v45)+56))
	F_LockBuffer(m, v181, int32(0))
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	v190 = F__emscripten_memset_bulkmem(m, v20+int32(32), base.I32_extend8_s(int32(0)), int32(291))
	mBase = m.M
	goto L41
L41:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v45)+52))
	v193 = v191
	goto L33
L42:
	;
	v549 = F_heap_getnext(m, v45)
	mBase = m.M
	v550 = m.ExcPending
	if v550 != 0 {
		goto L1
	} else {
		goto L100
	}
L43:
	;
	v489 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v216&int32(65535)+v20)+31)))
	if v489 != 0 {
		v541 = v478
		v547 = v484
		goto L42
	} else {
		goto L90
	}
L44:
	;
	v440 = int32(0)
	v442 = v20 + int32(20)
	v446 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v432)+2)))
	v447 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v432))))
	v448 = int32(16)
	v450 = v446 | v447<<(uint(v448)%32)
	v451 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v442)+2)))
	v452 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v442))))
	v455 = v451 | v452<<(uint(v448)%32)
	if base.Ui32(v450) < base.Ui32(v455) {
		v466 = int32(-1)
		goto L85
	} else {
		goto L86
	}
L45:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v399 = m.ExcPending
	if v399 != 0 {
		goto L1
	} else {
		goto L80
	}
L46:
	;
	v205 = int32(1)
	v208 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v200<<(uint(v205)%32)+v20)+334)))
	if base.Ui32(int32(2048)) <= base.Ui32((v208-v205)&int32(65535)) {
		goto L45
	} else {
		goto L49
	}
L47:
	;
	v216 = v200
	goto L48
L48:
	;
	if v100 == int32(0) {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v20)+24)) = uint16(v208)
	v216 = v208
	goto L48
L50:
	;
	if v94 == int32(0) {
		goto L53
	} else {
		goto L54
	}
L51:
	;
	v387 = v94
	goto L52
L52:
	;
	v478 = v387
	v484 = int32(1)
	goto L43
L53:
	;
	v259 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	v261 = int32(0)
	v267 = F_tuplesort_getdatum(m, v259, int32(1), v261, v20+int32(16), v20+int32(15), v261)
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L1
	} else {
		goto L63
	}
L54:
	;
	v222 = v20 + int32(20)
	v226 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v94)+2)))
	v227 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v94))))
	v228 = int32(16)
	v230 = v226 | v227<<(uint(v228)%32)
	v231 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v222)+2)))
	v232 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v222))))
	v235 = v231 | v232<<(uint(v228)%32)
	if base.Ui32(v230) < base.Ui32(v235) {
		v246 = int32(-1)
		goto L56
	} else {
		goto L57
	}
L55:
	;
	if int32(0) <= v246 {
		v432 = v94
		goto L44
	} else {
		goto L60
	}
L56:
	;
	goto L55
L57:
	;
	if base.Ui32(v235) < base.Ui32(v230) {
		v246 = int32(1)
		goto L56
	} else {
		goto L58
	}
L58:
	;
	v240 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v94)+4)))
	v241 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v222)+4)))
	if base.Ui32(v240) < base.Ui32(v241) {
		v246 = int32(-1)
		goto L56
	} else {
		goto L59
	}
L59:
	;
	v246 = base.B2i32(base.Ui32(v241) < base.Ui32(v240))
	goto L56
L60:
	;
	v249 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v94)+2)))
	v250 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v94))))
	if v249|v250<<(uint(int32(16))%32) != v193 {
		goto L53
	} else {
		goto L61
	}
L61:
	;
	v255 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v94)+4)))
	v257 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v20+v255)+31)) = uint8(v257)
	goto L53
L62:
	;
	v387 = int32(0)
	goto L52
L63:
	;
	if v267 == int32(0) {
		goto L62
	} else {
		goto L64
	}
L64:
	;
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v20)+16))
	v272 = *(*int64)(unsafe.Add(mBase, uint32(v271)))
	*(*uint16)(unsafe.Add(mBase, uint32(v20)+30)) = uint16(v272)
	v275 = int64(base.Ui64(v272) >> (uint(int64(16)) % 64))
	*(*uint16)(unsafe.Add(mBase, uint32(v20)+28)) = uint16(v275)
	v278 = int64(base.Ui64(v272) >> (uint(int64(32)) % 64))
	*(*uint16)(unsafe.Add(mBase, uint32(v20)+26)) = uint16(v278)
	goto L65
L65:
	;
	v298 = v20 + int32(26)
	v300 = v20 + int32(20)
	v304 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v298)+2)))
	v305 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v298))))
	v306 = int32(16)
	v308 = v304 | v305<<(uint(v306)%32)
	v309 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v300)+2)))
	v310 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v300))))
	v313 = v309 | v310<<(uint(v306)%32)
	if base.Ui32(v308) < base.Ui32(v313) {
		v324 = int32(-1)
		goto L68
	} else {
		goto L69
	}
L67:
	;
	if int32(0) <= v324 {
		goto L72
	} else {
		goto L73
	}
L68:
	;
	goto L67
L69:
	;
	if base.Ui32(v313) < base.Ui32(v308) {
		v324 = int32(1)
		goto L68
	} else {
		goto L70
	}
L70:
	;
	v318 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v298)+4)))
	v319 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v300)+4)))
	if base.Ui32(v318) < base.Ui32(v319) {
		v324 = int32(-1)
		goto L68
	} else {
		goto L71
	}
L71:
	;
	v324 = base.B2i32(base.Ui32(v319) < base.Ui32(v318))
	goto L68
L72:
	;
	v432 = v20 + int32(26)
	goto L44
L73:
	;
	goto L74
L74:
	;
	v329 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v20)+28)))
	v330 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v20)+26)))
	if v193 == v329|v330<<(uint(int32(16))%32) {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v335 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v20)+30)))
	v337 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v20+v335)+31)) = uint8(v337)
	goto L77
L76:
	;
	goto L77
L77:
	;
	v339 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	v341 = int32(0)
	v347 = F_tuplesort_getdatum(m, v339, int32(1), v341, v20+int32(16), v20+int32(15), v341)
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L1
	} else {
		goto L78
	}
L78:
	;
	if v347 == int32(0) {
		goto L62
	} else {
		goto L79
	}
L79:
	;
	v351 = *(*int32)(unsafe.Add(mBase, uint32(v20)+16))
	v352 = *(*int64)(unsafe.Add(mBase, uint32(v351)))
	*(*uint16)(unsafe.Add(mBase, uint32(v20)+30)) = uint16(v352)
	v355 = int64(base.Ui64(v352) >> (uint(int64(16)) % 64))
	*(*uint16)(unsafe.Add(mBase, uint32(v20)+28)) = uint16(v355)
	v358 = int64(base.Ui64(v352) >> (uint(int64(32)) % 64))
	*(*uint16)(unsafe.Add(mBase, uint32(v20)+26)) = uint16(v358)
	goto L65
L80:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	v403 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v88)+6)))
	v404 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v88)+4)))
	v405 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v406 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v88)+8)))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+4)) = v406
	*(*int32)(unsafe.Add(mBase, uint32(v20)+8)) = v405 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v403 | v404<<(uint(int32(16))%32)
	F_errmsg_internal(m, int32(718771), v20)
	mBase = m.M
	v417 = m.ExcPending
	if v417 != 0 {
		goto L1
	} else {
		goto L82
	}
L82:
	;
	F_errfinish(m, int32(495072), int32(1870), int32(284506))
	mBase = m.M
	v422 = m.ExcPending
	if v422 != 0 {
		goto L1
	} else {
		goto L83
	}
L83:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L84:
	;
	if v466 <= int32(0) {
		v541 = v432
		v547 = v440
		goto L42
	} else {
		goto L89
	}
L85:
	;
	goto L84
L86:
	;
	if base.Ui32(v455) < base.Ui32(v450) {
		v466 = int32(1)
		goto L85
	} else {
		goto L87
	}
L87:
	;
	v460 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v432)+4)))
	v461 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v442)+4)))
	if base.Ui32(v460) < base.Ui32(v461) {
		v466 = int32(-1)
		goto L85
	} else {
		goto L88
	}
L88:
	;
	v466 = base.B2i32(base.Ui32(v461) < base.Ui32(v460))
	goto L85
L89:
	;
	v478 = v432
	v484 = v440
	goto L43
L90:
	;
	v490 = *(*int32)(unsafe.Add(mBase, uint32(v29)+20))
	F_MemoryContextReset(m, v490)
	mBase = m.M
	v492 = m.ExcPending
	if v492 != 0 {
		goto L1
	} else {
		goto L91
	}
L91:
	;
	v494 = F_ExecStoreHeapTuple(m, v88, v32, int32(0))
	mBase = m.M
	v495 = m.ExcPending
	if v495 != 0 {
		goto L1
	} else {
		goto L92
	}
L92:
	;
	if v36 != 0 {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v496 = int32(4515120)
	v497 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v499 = *(*int32)(unsafe.Add(mBase, uint32(v29)+20))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v499
	v503 = *(*int32)(unsafe.Add(mBase, uint32(v36)+20))
	v504 = m.T0[v503].(func(*base.Module, int32, int32, int32) int32)(m, v36, v29, v20+int32(16))
	mBase = m.M
	v505 = m.ExcPending
	if v505 != 0 {
		goto L1
	} else {
		goto L96
	}
L94:
	;
	goto L95
L95:
	;
	F_FormIndexDatum(m, l2, v32, v22, v20+int32(960), v20+int32(928))
	mBase = m.M
	v517 = m.ExcPending
	if v517 != 0 {
		goto L1
	} else {
		goto L98
	}
L96:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v497
	if v504 == int32(0) {
		v541 = v478
		v547 = v484
		goto L42
	} else {
		goto L97
	}
L97:
	;
	goto L95
L98:
	;
	v524 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+116)))
	v526 = F_index_insert(m, l1, v20+int32(960), v20+int32(928), v20+int32(20), l0, v524, int32(0), l2)
	mBase = m.M
	v527 = m.ExcPending
	if v527 != 0 {
		goto L1
	} else {
		goto L99
	}
L99:
	;
	v528 = *(*float64)(unsafe.Add(mBase, uint32(l4)+24))
	*(*float64)(unsafe.Add(mBase, uint32(l4)+24)) = base.F64_add(v528, float64(1))
	v541 = v478
	v547 = v484
	goto L42
L100:
	;
	if v549 != 0 {
		v88 = v549
		v93 = v193
		v94 = v541
		v97 = v153
		v100 = v547
		goto L18
	} else {
		goto L101
	}
L101:
	;
	goto L19
L102:
	;
	F_ExecDropSingleTupleTableSlot(m, v32)
	mBase = m.M
	v574 = m.ExcPending
	if v574 != 0 {
		goto L1
	} else {
		goto L103
	}
L103:
	;
	F_FreeExecutorState(m, v22)
	mBase = m.M
	v576 = m.ExcPending
	if v576 != 0 {
		goto L1
	} else {
		goto L104
	}
L104:
	;
	v577 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l2)+88)) = v577
	*(*int32)(unsafe.Add(mBase, uint32(l2)+80)) = v577
	m.G0 = v20 + int32(1088)
	return
}
func F_heapam_relation_needs_toast_table(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
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
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v145 int32
	_ = v145
	v2 = int32(0)
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	if v11 <= v2 {
		return int32(0)
	} else {
		v19 = int32(0)
		v20 = v2
		v21 = v11
		v23 = v2
		v24 = v2
		for {
			v33 = v10 + int32(20) + v21<<(uint(int32(4))%32) + v19*int32(100)
			v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+91)))
			if v34 != 0 {
				v117 = v20
				v118 = v21
				v119 = v23
				v120 = v24
			} else {
				v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+90)))
				if v35 == int32(118) {
					v117 = v20
					v118 = v21
					v119 = v23
					v120 = v24
				} else {
					v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+83)))
					switch v38 - int32(99) {
					case 0:
						v53 = v20
					case 1:
						v53 = (v20 + int32(7)) & int32(-8)
					default:
						v53 = (v20 + int32(1)) & int32(-2)
					case 6:
						v53 = (v20 + int32(3)) & int32(-4)
					}
					v54 = int32(*(*int16)(unsafe.Add(mBase, uint32(v33)+72)))
					if int32(0) < v54 {
						v117 = v53 + v54
						v118 = v21
						v119 = v23
						v120 = v24
					} else {
						v58 = *(*int32)(unsafe.Add(mBase, uint32(v33)+68))
						v59 = *(*int32)(unsafe.Add(mBase, uint32(v33)+76))
						v61 = int32(-1)
						if v59 < int32(0) {
							v101 = v61
							v103 = v101
						} else {
							if base.Ui32(int32(2)) <= base.Ui32(v58-int32(1042)) {
								switch v58 - int32(1560) {
								case 0, 2:
									v97 = int32(8)
									v98 = base.I32_div_s(v59+int32(7), v97)
									v101 = v98 + v97
									v103 = v101
								case 1:
									v101 = v61
									v103 = v101
								default:
									if v58 != int32(1700) {
										v101 = v61
										v103 = v101
									} else {
										if int32(4) <= v59 {
											v94 = int32(base.Ui32(int32(base.Ui32(v59-int32(4))>>(uint(int32(16))%32))+int32(6))>>(uint(int32(1))%32))&int32(65534) + int32(8)
										} else {
											v94 = int32(-1)
										}
										v103 = v94
									}
								}
							} else {
								v70 = F_GetDatabaseEncoding(m)
								mBase = m.M
								v71 = F_pg_encoding_max_length(m, v70)
								mBase = m.M
								v72 = int32(4)
								v103 = v71*(v59-v72) + v72
							}
						}
						v104 = int32(0)
						if v104 < v103 {
							v107 = v103
						} else {
							v107 = v104
						}
						v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+84)))
						v116 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
						v117 = v107 + v53
						v118 = v116
						v119 = base.B2i32(v112 != int32(112)) | v23
						v120 = base.B2i32(v103 < int32(0)) | v24
					}
				}
			}
			v123 = v19 + int32(1)
			if v123 < v118 {
				v19 = v123
				v20 = v117
				v21 = v118
				v23 = v119
				v24 = v120
				continue
			} else {
				break
			}
			break
		}
		if (v119^int32(-1)|v120)&int32(1) != 0 {
			v145 = v119
		} else {
			v130 = int32(7)
			v133 = base.I32_div_s(v118+v130, int32(8))
			v136 = int32(-8)
			v145 = base.B2i32(base.Ui32(int32(2032)) < base.Ui32((v133+int32(30))&v136+(v117+v130)&v136))
		}
		return v145 & int32(1)
	}
}
func F_heapam_scan_analyze_next_tuple(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v182 int32
	_ = v182
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v218 int32
	_ = v218
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v260 int32
	_ = v260
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v294 int32
	_ = v294
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v302 int32
	_ = v302
	var v312 int32
	_ = v312
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v340 int32
	_ = v340
	var v348 int32
	_ = v348
	var v356 int32
	_ = v356
	var v359 float64
	_ = v359
	var v363 int32
	_ = v363
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v374 float64
	_ = v374
	var v381 int32
	_ = v381
	var v383 int32
	_ = v383
	var v386 int32
	_ = v386
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v407 int32
	_ = v407
	var v413 int32
	_ = v413
	var v417 int32
	_ = v417
	var v422 int32
	_ = v422
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	if v14 < int32(0) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v413 = m.ExcPending
	if v413 != 0 {
		goto L20
	} else {
		goto L116
	}
L2:
	;
	v33 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v32)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v33) {
		goto L6
	} else {
		goto L7
	}
L3:
	;
	v18 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v18+(v14^int32(-1))<<(uint(int32(2))%32))))
	v32 = v24
	goto L2
L4:
	;
	goto L5
L5:
	;
	v26 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v32 = v26 + v14<<(uint(int32(13))%32) + int32(-8192)
	goto L2
L6:
	;
	v43 = int32(base.Ui32(v33+int32(262120))>>(uint(int32(2))%32)) & int32(65535)
	goto L8
L7:
	;
	v43 = int32(0)
	goto L8
L8:
	;
	if base.Ui32(v13) <= base.Ui32(v43) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v48 = l4 + int32(48)
	v54 = v13
	goto L12
L10:
	;
	v399 = v14
	goto L11
L11:
	;
	F_UnlockReleaseBuffer(m, v399)
	mBase = m.M
	v401 = m.ExcPending
	if v401 != 0 {
		goto L20
	} else {
		goto L114
	}
L12:
	;
	v67 = v54&int32(65535)<<(uint(int32(2))%32) + (v32 + int32(24)) - int32(4)
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
	switch int32(base.Ui32(v68)>>(uint(int32(15))%32))&int32(3) - int32(1) {
	case 0:
		goto L16
	default:
		goto L14
	case 2:
		goto L15
	}
L13:
	;
	v386 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v399 = v386
	goto L11
L14:
	;
	v381 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	v383 = v381 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+100)) = v383
	if base.Ui32(v383) <= base.Ui32(v43) {
		v54 = v383
		goto L12
	} else {
		goto L113
	}
L15:
	;
	v374 = *(*float64)(unsafe.Add(mBase, uint32(l3)))
	*(*float64)(unsafe.Add(mBase, uint32(l3))) = base.F64_add(v374, float64(1))
	goto L14
L16:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	*(*uint16)(unsafe.Add(mBase, uint32(l4)+56)) = uint16(v54)
	*(*uint16)(unsafe.Add(mBase, uint32(l4)+54)) = uint16(v75)
	v79 = int32(base.Ui32(v75) >> (uint(int32(16)) % 32))
	*(*uint16)(unsafe.Add(mBase, uint32(l4)+52)) = uint16(v79)
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v81)+56))
	*(*int32)(unsafe.Add(mBase, uint32(l4)+60)) = v82
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
	*(*int32)(unsafe.Add(mBase, uint32(l4)+64)) = v32 + v84&int32(32767)
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
	*(*int32)(unsafe.Add(mBase, uint32(l4)+48)) = int32(base.Ui32(v89) >> (uint(int32(17)) % 32))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v94 = F_HeapTupleSatisfiesVacuum(m, v48, l1, v93)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L20
	} else {
		goto L21
	}
L17:
	;
	v359 = *(*float64)(unsafe.Add(mBase, uint32(l2)))
	*(*float64)(unsafe.Add(mBase, uint32(l2))) = base.F64_add(v359, float64(1))
	v363 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	F_ExecStoreBufferHeapTuple(m, v48, l4, v363)
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L20
	} else {
		goto L112
	}
L18:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(l4)+64))
	v228 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v227)+20)))
	if v228&int32(6272) == int32(4096) {
		goto L67
	} else {
		goto L68
	}
L19:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(l4)+64))
	v100 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v99)+20)))
	v101 = int32(768)
	if v100&v101 != v101 {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	return int32(0)
L21:
	;
	switch v94 {
	case 0, 2:
		goto L15
	case 1:
		goto L17
	case 3:
		goto L19
	case 4:
		goto L18
	default:
		goto L1
	}
L22:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v99)))
	v106 = v105
	goto L24
L23:
	;
	v106 = int32(2)
	goto L24
L24:
	;
	if base.Ui32(v106) < base.Ui32(int32(3)) {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	if v226 != 0 {
		goto L17
	} else {
		goto L65
	}
L26:
	;
	v226 = int32(0)
	goto L25
L27:
	;
	goto L28
L28:
	;
	v117 = *(*int32)(unsafe.Add(mBase, _consts[67]))
	if v117 == v106 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v226 = int32(1)
	goto L25
L30:
	;
	goto L31
L31:
	;
	v121 = *(*int32)(unsafe.Add(mBase, _consts[68]))
	if v121 <= int32(0) {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	v226 = v218
	goto L25
L33:
	;
	v125 = *(*int32)(unsafe.Add(mBase, _consts[65]))
	if v125 == int32(0) {
		v218 = int32(0)
		goto L32
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	v187 = *(*int32)(unsafe.Add(mBase, _consts[69]))
	v189 = int32(0)
	v191 = v121 - int32(1)
	goto L55
L36:
	;
	v130 = v125
	goto L37
L37:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v130)+20))
	if v135 == int32(4) {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	v218 = int32(0)
	goto L32
L39:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v130)+80))
	if v182 != 0 {
		v130 = v182
		goto L37
	} else {
		goto L54
	}
L40:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v130)))
	if v138 == int32(0) {
		goto L39
	} else {
		goto L41
	}
L41:
	;
	v141 = int32(1)
	if v106 == v138 {
		v218 = v141
		goto L32
	} else {
		goto L42
	}
L42:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v130)+52))
	v145 = v143 - int32(1)
	if v145 < int32(0) {
		goto L39
	} else {
		goto L43
	}
L43:
	;
	v150 = int32(0)
	v152 = v145
	goto L44
L44:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v130)+48))
	v158 = int32(2)
	v159 = base.I32_div_s(v152-v150, v158)
	v160 = v159 + v150
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v156+v160<<(uint(v158)%32))))
	if v164 == v106 {
		v218 = v141
		goto L32
	} else {
		goto L46
	}
L45:
	;
	goto L39
L46:
	;
	v168 = F_TransactionIdPrecedes(m, v164, v106)
	mBase = m.M
	if v168 != 0 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v169 = v160 + int32(1)
	goto L49
L48:
	;
	v169 = v150
	goto L49
L49:
	;
	if v168 != 0 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v172 = v152
	goto L52
L51:
	;
	v172 = v160 - int32(1)
	goto L52
L52:
	;
	if v169 <= v172 {
		v150 = v169
		v152 = v172
		goto L44
	} else {
		goto L53
	}
L53:
	;
	goto L45
L54:
	;
	goto L38
L55:
	;
	v196 = int32(2)
	v197 = base.I32_div_s(v191-v189, v196)
	v198 = v197 + v189
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v187+v198<<(uint(v196)%32))))
	v203 = base.B2i32(v202 == v106)
	if v202 == v106 {
		v218 = v203
		goto L32
	} else {
		goto L57
	}
L56:
	;
	v218 = v203
	goto L32
L57:
	;
	v206 = base.B2i32(base.Ui32(v202) < base.Ui32(v106))
	if base.Ui32(v202) < base.Ui32(v106) {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v207 = v198 + int32(1)
	goto L60
L59:
	;
	v207 = v189
	goto L60
L60:
	;
	if base.Ui32(v202) < base.Ui32(v106) {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v210 = v191
	goto L63
L62:
	;
	v210 = v198 - int32(1)
	goto L63
L63:
	;
	if v207 <= v210 {
		v189 = v207
		v191 = v210
		goto L55
	} else {
		goto L64
	}
L64:
	;
	goto L56
L65:
	;
	goto L14
L66:
	;
	if base.Ui32(v236) < base.Ui32(int32(3)) {
		goto L72
	} else {
		goto L73
	}
L67:
	;
	v233 = F_HeapTupleGetUpdateXid(m, v227)
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L20
	} else {
		goto L70
	}
L68:
	;
	goto L69
L69:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v227)+4))
	v236 = v235
	goto L66
L70:
	;
	v236 = v233
	goto L66
L71:
	;
	if v356 != 0 {
		goto L15
	} else {
		goto L111
	}
L72:
	;
	v356 = int32(0)
	goto L71
L73:
	;
	goto L74
L74:
	;
	v247 = *(*int32)(unsafe.Add(mBase, _consts[67]))
	if v247 == v236 {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v356 = int32(1)
	goto L71
L76:
	;
	goto L77
L77:
	;
	v251 = *(*int32)(unsafe.Add(mBase, _consts[68]))
	if v251 <= int32(0) {
		goto L79
	} else {
		goto L80
	}
L78:
	;
	v356 = v348
	goto L71
L79:
	;
	v255 = *(*int32)(unsafe.Add(mBase, _consts[65]))
	if v255 == int32(0) {
		v348 = int32(0)
		goto L78
	} else {
		goto L82
	}
L80:
	;
	goto L81
L81:
	;
	v317 = *(*int32)(unsafe.Add(mBase, _consts[69]))
	v319 = int32(0)
	v321 = v251 - int32(1)
	goto L101
L82:
	;
	v260 = v255
	goto L83
L83:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v260)+20))
	if v265 == int32(4) {
		goto L85
	} else {
		goto L86
	}
L84:
	;
	v348 = int32(0)
	goto L78
L85:
	;
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v260)+80))
	if v312 != 0 {
		v260 = v312
		goto L83
	} else {
		goto L100
	}
L86:
	;
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v260)))
	if v268 == int32(0) {
		goto L85
	} else {
		goto L87
	}
L87:
	;
	v271 = int32(1)
	if v236 == v268 {
		v348 = v271
		goto L78
	} else {
		goto L88
	}
L88:
	;
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v260)+52))
	v275 = v273 - int32(1)
	if v275 < int32(0) {
		goto L85
	} else {
		goto L89
	}
L89:
	;
	v280 = int32(0)
	v282 = v275
	goto L90
L90:
	;
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v260)+48))
	v288 = int32(2)
	v289 = base.I32_div_s(v282-v280, v288)
	v290 = v289 + v280
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v286+v290<<(uint(v288)%32))))
	if v294 == v236 {
		v348 = v271
		goto L78
	} else {
		goto L92
	}
L91:
	;
	goto L85
L92:
	;
	v298 = F_TransactionIdPrecedes(m, v294, v236)
	mBase = m.M
	if v298 != 0 {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v299 = v290 + int32(1)
	goto L95
L94:
	;
	v299 = v280
	goto L95
L95:
	;
	if v298 != 0 {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v302 = v282
	goto L98
L97:
	;
	v302 = v290 - int32(1)
	goto L98
L98:
	;
	if v299 <= v302 {
		v280 = v299
		v282 = v302
		goto L90
	} else {
		goto L99
	}
L99:
	;
	goto L91
L100:
	;
	goto L84
L101:
	;
	v326 = int32(2)
	v327 = base.I32_div_s(v321-v319, v326)
	v328 = v327 + v319
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v317+v328<<(uint(v326)%32))))
	v333 = base.B2i32(v332 == v236)
	if v332 == v236 {
		v348 = v333
		goto L78
	} else {
		goto L103
	}
L102:
	;
	v348 = v333
	goto L78
L103:
	;
	v336 = base.B2i32(base.Ui32(v332) < base.Ui32(v236))
	if base.Ui32(v332) < base.Ui32(v236) {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	v337 = v328 + int32(1)
	goto L106
L105:
	;
	v337 = v319
	goto L106
L106:
	;
	if base.Ui32(v332) < base.Ui32(v236) {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	v340 = v321
	goto L109
L108:
	;
	v340 = v328 - int32(1)
	goto L109
L109:
	;
	if v337 <= v340 {
		v319 = v337
		v321 = v340
		goto L101
	} else {
		goto L110
	}
L110:
	;
	goto L102
L111:
	;
	goto L17
L112:
	;
	v366 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	v367 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+100)) = v366 + v367
	return v367
L113:
	;
	goto L13
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = int32(0)
	v404 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
	v405 = *(*int32)(unsafe.Add(mBase, uint32(v404)+12))
	m.T0[v405].(func(*base.Module, int32))(m, l4)
	mBase = m.M
	v407 = m.ExcPending
	if v407 != 0 {
		goto L20
	} else {
		goto L115
	}
L115:
	;
	return int32(0)
L116:
	;
	F_errmsg_internal(m, int32(97885), int32(0))
	mBase = m.M
	v417 = m.ExcPending
	if v417 != 0 {
		goto L20
	} else {
		goto L117
	}
L117:
	;
	F_errfinish(m, int32(495072), int32(1147), int32(383147))
	mBase = m.M
	v422 = m.ExcPending
	if v422 != 0 {
		goto L20
	} else {
		goto L118
	}
L118:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_heapam_scan_sample_next_tuple(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v78 int32
	_ = v78
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v129 int32
	_ = v129
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v189 int32
	_ = v189
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v249 int64
	_ = v249
	v4 = int32(0)
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)+124))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v26 = v24 & int32(256)
	if v26 == v4 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	F_LockBuffer(m, v29, int32(1))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	if v35 < int32(0) {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	return int32(0)
L5:
	;
	goto L3
L6:
	;
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+10)))
	if v54&int32(4) != 0 {
		goto L10
	} else {
		goto L11
	}
L7:
	;
	v39 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v39+(v35^int32(-1))<<(uint(int32(2))%32))))
	v53 = v45
	goto L6
L8:
	;
	goto L9
L9:
	;
	v47 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v53 = v47 + v35<<(uint(int32(13))%32) + int32(-8192)
	goto L6
L10:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57)+29)))
	v61 = v58 ^ int32(1)
	goto L12
L11:
	;
	v61 = v4
	goto L12
L12:
	;
	v65 = int32(base.Ui32(v22) >> (uint(int32(16)) % 32))
	v69 = l0 - int32(-64)
	v70 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v53)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v70) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v78 = int32(base.Ui32(v70+int32(262120)) >> (uint(int32(2)) % 32))
	goto L15
L14:
	;
	v78 = int32(0)
	goto L15
L15:
	;
	goto L17
L16:
	;
	return base.B2i32(base.Ui32(v110&int32(65535)) < base.Ui32(int32(2048)))
L17:
	;
	v103 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	if v103 != 0 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v234 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	F_ExecStoreBufferHeapTuple(m, v69, l2, v234)
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L4
	} else {
		goto L59
	}
L19:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L4
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v23)+28))
	v107 = m.T0[v106].(func(*base.Module, int32, int32, int32) int32)(m, l1, v22, v78&int32(65535))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L4
	} else {
		goto L26
	}
L22:
	;
	goto L21
L23:
	;
	goto L18
L24:
	;
	if v189 == int32(0) {
		goto L17
	} else {
		goto L58
	}
L25:
	;
	if v26 != 0 {
		goto L17
	} else {
		goto L56
	}
L26:
	;
	v110 = v107 - int32(1)
	if base.Ui32(v110&int32(65535)) <= base.Ui32(int32(2047)) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v119 = v107<<(uint(int32(2))%32) + (v53 + int32(24)) - int32(4)
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v119)))
	if v120&int32(98304) != int32(32768) {
		goto L17
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	if v26 == int32(0) {
		goto L51
	} else {
		goto L52
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = v53 + v120&int32(32767)
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v119)))
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+72)) = uint16(v107)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+70)) = uint16(v22)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+68)) = uint16(v65)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = int32(base.Ui32(v129) >> (uint(int32(17)) % 32))
	v136 = int32(1)
	if v61&v136 != 0 {
		v189 = v136
		goto L31
	} else {
		goto L32
	}
L31:
	;
	if v26 != 0 {
		goto L24
	} else {
		goto L47
	}
L32:
	;
	v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+29)))
	if v139&int32(1) != 0 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	v147 = int32(0)
	v148 = v142
	goto L36
L34:
	;
	goto L35
L35:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v181 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v182 = F_HeapTupleSatisfiesVisibility(m, v69, v180, v181)
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L4
	} else {
		goto L46
	}
L36:
	;
	if base.Ui32(v148) <= base.Ui32(v147) {
		goto L25
	} else {
		goto L38
	}
L37:
	;
	v189 = v136
	goto L31
L38:
	;
	v167 = int32(1)
	v169 = int32(base.Ui32(v148-v147)>>(uint(v167)%32)) + v147
	v175 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0+int32(108)+v169<<(uint(v167)%32)))))
	v176 = base.B2i32(base.Ui32(v107) < base.Ui32(v175))
	if base.Ui32(v107) < base.Ui32(v175) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v177 = v147
	goto L41
L40:
	;
	v177 = v169 + v167
	goto L41
L41:
	;
	if base.Ui32(v107) < base.Ui32(v175) {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v178 = v169
	goto L44
L43:
	;
	v178 = v148
	goto L44
L44:
	;
	if v107 != v175 {
		v147 = v177
		v148 = v178
		goto L36
	} else {
		goto L45
	}
L45:
	;
	goto L37
L46:
	;
	v189 = v182
	goto L31
L47:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v206 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v207 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	F_HeapCheckForSerializableConflictOut(m, v189, v205, v69, v206, v207)
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L4
	} else {
		goto L48
	}
L48:
	;
	if v189 == int32(0) {
		goto L17
	} else {
		goto L49
	}
L49:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	F_LockBuffer(m, v212, int32(0))
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L4
	} else {
		goto L50
	}
L50:
	;
	goto L23
L51:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	F_LockBuffer(m, v218, int32(0))
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L4
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v222)+12))
	m.T0[v223].(func(*base.Module, int32))(m, l2)
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L4
	} else {
		goto L55
	}
L54:
	;
	goto L53
L55:
	;
	goto L16
L56:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v228 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v229 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	F_HeapCheckForSerializableConflictOut(m, int32(0), v227, v69, v228, v229)
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L4
	} else {
		goto L57
	}
L57:
	;
	goto L17
L58:
	;
	goto L23
L59:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v237)+272))
	if v238 == int32(0) {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v241 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v237)+268)))
	if v241 != int32(1) {
		goto L16
	} else {
		goto L63
	}
L61:
	;
	v248 = v238
	goto L62
L62:
	;
	v249 = *(*int64)(unsafe.Add(mBase, uint32(v248)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v248)+24)) = v249 + int64(1)
	goto L16
L63:
	;
	F_pgstat_assoc_relation(m, v237)
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L4
	} else {
		goto L64
	}
L64:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v246)+272))
	v248 = v247
	goto L62
}
func F_heapam_slot_callbacks(m *base.Module, l0 int32) int32 {
	return int32(1617568)
}
func F_heapam_tuple_satisfies_snapshot(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
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
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
	F_LockBuffer(m, v4, int32(1))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
		v12 = F_HeapTupleSatisfiesVisibility(m, v10, l2, v11)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int32(0)
		} else {
			v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
			F_LockBuffer(m, v14, int32(0))
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int32(0)
			} else {
				return v12
			}
		}
	}
}
func F_heapam_tuple_tid_valid(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	v3 = int32(0)
	if l1 == v3 {
		v16 = v3
	} else {
		v6 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
		if v6 == int32(0) {
			v16 = v3
		} else {
			v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
			v10 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+2)))
			v11 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1))))
			v16 = base.B2i32(base.Ui32(v10|v11<<(uint(int32(16))%32)) < base.Ui32(v9))
		}
	}
	return v16
}
