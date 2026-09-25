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
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v153 int32
	_ = v153
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v170 int32
	_ = v170
	var v175 int32
	_ = v175
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v201 float64
	_ = v201
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v228 int32
	_ = v228
	var v232 int32
	_ = v232
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v265 int32
	_ = v265
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v279 int32
	_ = v279
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v293 int32
	_ = v293
	var v296 int32
	_ = v296
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v342 int32
	_ = v342
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v356 int32
	_ = v356
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v382 int32
	_ = v382
	var v386 int32
	_ = v386
	var v390 int32
	_ = v390
	var v395 int32
	_ = v395
	var v400 int32
	_ = v400
	var v403 int32
	_ = v403
	var v406 int32
	_ = v406
	var v408 int32
	_ = v408
	var v410 int32
	_ = v410
	var v415 int32
	_ = v415
	var v417 int32
	_ = v417
	var v421 int32
	_ = v421
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v429 int32
	_ = v429
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v437 int32
	_ = v437
	var v447 int32
	_ = v447
	var v452 int32
	_ = v452
	var v454 int32
	_ = v454
	var v456 int32
	_ = v456
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v475 int32
	_ = v475
	var v483 int32
	_ = v483
	var v491 int32
	_ = v491
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v498 int32
	_ = v498
	var v506 int32
	_ = v506
	var v511 int32
	_ = v511
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v534 int32
	_ = v534
	var v538 int32
	_ = v538
	var v542 int32
	_ = v542
	var v547 int32
	_ = v547
	var v552 int32
	_ = v552
	var v555 int32
	_ = v555
	var v558 int32
	_ = v558
	var v560 int32
	_ = v560
	var v562 int32
	_ = v562
	var v567 int32
	_ = v567
	var v569 int32
	_ = v569
	var v573 int32
	_ = v573
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v581 int32
	_ = v581
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v589 int32
	_ = v589
	var v599 int32
	_ = v599
	var v604 int32
	_ = v604
	var v606 int32
	_ = v606
	var v608 int32
	_ = v608
	var v613 int32
	_ = v613
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v627 int32
	_ = v627
	var v635 int32
	_ = v635
	var v643 int32
	_ = v643
	var v648 int32
	_ = v648
	var v649 int32
	_ = v649
	var v652 int32
	_ = v652
	var v660 int32
	_ = v660
	var v665 int32
	_ = v665
	var v668 int32
	_ = v668
	var v669 int32
	_ = v669
	var v674 int32
	_ = v674
	var v683 int32
	_ = v683
	var v684 int32
	_ = v684
	var v685 int32
	_ = v685
	var v690 int32
	_ = v690
	var v700 int32
	_ = v700
	var v704 int32
	_ = v704
	var v707 int32
	_ = v707
	var v711 int32
	_ = v711
	var v715 int32
	_ = v715
	var v720 int32
	_ = v720
	var v727 float64
	_ = v727
	var v735 int32
	_ = v735
	var v737 float64
	_ = v737
	var v738 int32
	_ = v738
	var v741 int32
	_ = v741
	var v742 int32
	_ = v742
	var v744 int32
	_ = v744
	var v747 int32
	_ = v747
	var v750 int32
	_ = v750
	var v752 int32
	_ = v752
	var v756 int32
	_ = v756
	var v767 int32
	_ = v767
	var v789 float64
	_ = v789
	var v790 int32
	_ = v790
	var v792 int32
	_ = v792
	var v793 int32
	_ = v793
	var v795 int32
	_ = v795
	var v796 int32
	_ = v796
	var v797 int32
	_ = v797
	var v799 int32
	_ = v799
	var v803 int32
	_ = v803
	var v804 int32
	_ = v804
	var v805 int32
	_ = v805
	var v816 int32
	_ = v816
	var v817 int32
	_ = v817
	var v818 int32
	_ = v818
	var v821 int32
	_ = v821
	var v826 int32
	_ = v826
	var v827 int32
	_ = v827
	var v830 int32
	_ = v830
	var v834 int32
	_ = v834
	var v840 int32
	_ = v840
	var v842 int32
	_ = v842
	var v848 int32
	_ = v848
	var v851 int32
	_ = v851
	var v855 int32
	_ = v855
	var v856 int32
	_ = v856
	var v859 int32
	_ = v859
	var v860 int32
	_ = v860
	var v861 int32
	_ = v861
	var v868 int32
	_ = v868
	var v878 int32
	_ = v878
	var v886 int32
	_ = v886
	var v916 float64
	_ = v916
	var v917 int32
	_ = v917
	var v918 int32
	_ = v918
	var v948 float64
	_ = v948
	var v950 int32
	_ = v950
	var v955 int32
	_ = v955
	var v956 int64
	_ = v956
	var v959 int32
	_ = v959
	var v963 int32
	_ = v963
	var v968 int32
	_ = v968
	var v970 int32
	_ = v970
	var v971 int32
	_ = v971
	var v974 int32
	_ = v974
	var v978 int32
	_ = v978
	var v980 int32
	_ = v980
	var v981 int32
	_ = v981
	var v989 int32
	_ = v989
	var v990 int32
	_ = v990
	var v996 int32
	_ = v996
	var v1001 int32
	_ = v1001
	var v1002 int32
	_ = v1002
	var v1003 int32
	_ = v1003
	var v1005 int32
	_ = v1005
	var v1007 int32
	_ = v1007
	var v1009 int32
	_ = v1009
	var v1011 int32
	_ = v1011
	var v1012 int32
	_ = v1012
	var v1023 int32
	_ = v1023
	var v1026 int32
	_ = v1026
	var v1027 int32
	_ = v1027
	var v1028 int32
	_ = v1028
	var v1029 int32
	_ = v1029
	var v1040 int32
	_ = v1040
	var v1045 int32
	_ = v1045
	v12 = int32(0)
	v30 = float64(0)
	v31 = m.G0
	v33 = v31 - int32(800)
	m.G0 = v33
	v36 = int32(1)
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	if base.Ui32(v37) < base.Ui32(int32(_a_F_heapam_index_build_range_scan_0)) {
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
	v73 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_index_build_range_scan[0]))
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
	v89 = int32(_a_F_heapam_index_build_range_scan_1)
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
	v116 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_index_build_range_scan[1]))
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
	v120 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_heapam_index_build_range_scan[2])))
	if v120&int32(1) == int32(0) {
		goto L40
	} else {
		goto L42
	}
L42:
	;
	v125 = int32(_a_F_heapam_index_build_range_scan_2)
	v127 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_index_build_range_scan[3]))
	v128 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_heapam_index_build_range_scan[3])) = v127 + v128
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v116)))
	*(*int32)(unsafe.Add(mBase, uint32(v116))) = v131 + v128
	v135 = int32(0)
	v137 = int32(_a_F_heapam_index_build_range_scan_3)
	v138 = base.AtomicRmwOr32(m, v135, v137, v135)
	*(*int64)(unsafe.Add(mBase, uint32(v116+int32(120))+232)) = v113
	v146 = base.AtomicRmwOr32(m, v135, v137, v135)
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v116)))
	*(*int32)(unsafe.Add(mBase, uint32(v116))) = v147 + v128
	v153 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_index_build_range_scan[3]))
	*(*int32)(unsafe.Add(mBase, _c_F_heapam_index_build_range_scan[3])) = v153 - v128
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
	v162 = F_heap_getnext(m, v103)
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L8
	} else {
		goto L47
	}
L46:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1023 = m.ExcPending
	if v1023 != 0 {
		goto L8
	} else {
		goto L273
	}
L47:
	;
	if v162 != 0 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v170 = int32(-1)
	v175 = v162
	v189 = v170
	v193 = v170
	v201 = v30
	goto L51
L49:
	;
	v948 = v30
	goto L50
L50:
	;
	if l5 != 0 {
		goto L256
	} else {
		goto L257
	}
L51:
	;
	v203 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_index_build_range_scan[4]))
	if v203 != 0 {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	v948 = v916
	goto L50
L53:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L8
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v103)+52))
	if l5 == int32(0) {
		v271 = v206
		v273 = v189
		goto L57
	} else {
		goto L58
	}
L56:
	;
	goto L55
L57:
	;
	if v271 != v193 {
		goto L74
	} else {
		goto L75
	}
L58:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v103)+32))
	if v209 != 0 {
		goto L60
	} else {
		goto L61
	}
L59:
	;
	if v189 == v222 {
		v271 = v206
		v273 = v189
		goto L57
	} else {
		goto L69
	}
L60:
	;
	v212 = v209 + int32(28)
	goto L62
L61:
	;
	v212 = v103 + int32(40)
	goto L62
L62:
	;
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v212)))
	if base.Ui32(v213) < base.Ui32(v206) {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v222 = v206 - v213
	goto L59
L64:
	;
	goto L65
L65:
	;
	if v209 != 0 {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v218 = v209 + int32(20)
	goto L68
L67:
	;
	v218 = v103 + int32(36)
	goto L68
L68:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v218)))
	v222 = v219 + (v206 - v213)
	goto L59
L69:
	;
	v228 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_index_build_range_scan[1]))
	if v228 == int32(0) {
		goto L71
	} else {
		goto L72
	}
L70:
	;
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v103)+52))
	v271 = v269
	v273 = v222
	goto L57
L71:
	;
	goto L70
L72:
	;
	v232 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_heapam_index_build_range_scan[2])))
	if v232&int32(1) == int32(0) {
		goto L71
	} else {
		goto L73
	}
L73:
	;
	v237 = int32(_a_F_heapam_index_build_range_scan_2)
	v239 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_index_build_range_scan[3]))
	v240 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_heapam_index_build_range_scan[3])) = v239 + v240
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v228)))
	*(*int32)(unsafe.Add(mBase, uint32(v228))) = v243 + v240
	v247 = int32(0)
	v249 = int32(_a_F_heapam_index_build_range_scan_3)
	v250 = base.AtomicRmwOr32(m, v247, v249, v247)
	*(*int64)(unsafe.Add(mBase, uint32(v228+int32(128))+232)) = base.I64_extend_i32_u(v222)
	v258 = base.AtomicRmwOr32(m, v247, v249, v247)
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v228)))
	*(*int32)(unsafe.Add(mBase, uint32(v228))) = v259 + v240
	v265 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_index_build_range_scan[3]))
	*(*int32)(unsafe.Add(mBase, _c_F_heapam_index_build_range_scan[3])) = v265 - v240
	goto L71
L74:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v103)+56))
	if v275 < int32(0) {
		goto L78
	} else {
		goto L79
	}
L75:
	;
	v307 = v193
	goto L76
L76:
	;
	if base.B2i32(v104 != int32(_a_F_heapam_index_build_range_scan_1)) == int32(0) {
		goto L86
	} else {
		goto L87
	}
L77:
	;
	F_LockBuffer(m, v275, int32(1))
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L8
	} else {
		goto L81
	}
L78:
	;
	v279 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_index_build_range_scan[5]))
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v279+(v275^int32(-1))<<(uint(int32(2))%32))))
	v293 = v285
	goto L77
L79:
	;
	goto L80
L80:
	;
	v287 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_index_build_range_scan[6]))
	v293 = v287 + v275<<(uint(int32(13))%32) + int32(-8192)
	goto L77
L81:
	;
	F_heap_get_root_tuples(m, v293, v33+int32(48))
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L8
	} else {
		goto L82
	}
L82:
	;
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v103)+56))
	F_LockBuffer(m, v301, int32(0))
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L8
	} else {
		goto L83
	}
L83:
	;
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v103)+52))
	v307 = v305
	goto L76
L84:
	;
	v917 = F_heap_getnext(m, v103)
	mBase = m.M
	v918 = m.ExcPending
	if v918 != 0 {
		goto L8
	} else {
		goto L254
	}
L85:
	;
	v790 = *(*int32)(unsafe.Add(mBase, uint32(v64)+20))
	F_MemoryContextReset(m, v790)
	mBase = m.M
	v792 = m.ExcPending
	if v792 != 0 {
		goto L8
	} else {
		goto L230
	}
L86:
	;
	goto L89
L87:
	;
	goto L88
L88:
	;
	v767 = int32(1)
	v789 = base.F64_add(v201, float64(1))
	goto L85
L89:
	;
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v103)+56))
	F_LockBuffer(m, v342, int32(1))
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L8
	} else {
		goto L91
	}
L91:
	;
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v103)+56))
	v347 = F_HeapTupleSatisfiesVacuum(m, v175, v79, v346)
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L8
	} else {
		goto L103
	}
L92:
	;
	v744 = *(*int32)(unsafe.Add(mBase, uint32(v103)+56))
	F_LockBuffer(m, v744, int32(0))
	mBase = m.M
	v747 = m.ExcPending
	if v747 != 0 {
		goto L8
	} else {
		goto L226
	}
L93:
	;
	v738 = *(*int32)(unsafe.Add(mBase, uint32(v103)+56))
	F_LockBuffer(m, v738, int32(0))
	mBase = m.M
	v741 = m.ExcPending
	if v741 != 0 {
		goto L8
	} else {
		goto L225
	}
L94:
	;
	v735 = int32(0)
	v737 = base.F64_add(v201, float64(1))
	goto L93
L95:
	;
	v735 = int32(1)
	v737 = v727
	goto L93
L96:
	;
	v727 = base.F64_add(v201, float64(1))
	goto L95
L97:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v711 = m.ExcPending
	if v711 != 0 {
		goto L8
	} else {
		goto L222
	}
L98:
	;
	v704 = *(*int32)(unsafe.Add(mBase, uint32(v103)+56))
	F_LockBuffer(m, v704, int32(0))
	mBase = m.M
	v707 = m.ExcPending
	if v707 != 0 {
		goto L8
	} else {
		goto L221
	}
L99:
	;
	v700 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l2)+122)) = uint8(v700)
	goto L98
L100:
	;
	if l4 != 0 {
		goto L94
	} else {
		goto L159
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
	v349 = int32(0)
	v350 = *(*int32)(unsafe.Add(mBase, uint32(v175)+16))
	v351 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v350)+19)))
	if v351&int32(64) == v349 {
		v735 = v349
		v737 = v201
		goto L93
	} else {
		goto L104
	}
L103:
	;
	switch v347 {
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
	v356 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v350)+20)))
	if v356&int32(2048) != 0 {
		v735 = v349
		v737 = v201
		goto L93
	} else {
		goto L105
	}
L105:
	;
	if v356&int32(768) != int32(512) {
		goto L99
	} else {
		goto L106
	}
L106:
	;
	v735 = v349
	v737 = v201
	goto L93
L107:
	;
	v364 = *(*int32)(unsafe.Add(mBase, uint32(v175)+16))
	v365 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v364)+20)))
	v366 = int32(768)
	if v365&v366 != v366 {
		goto L108
	} else {
		goto L109
	}
L108:
	;
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v364)))
	v371 = v370
	goto L110
L109:
	;
	v371 = int32(2)
	goto L110
L110:
	;
	if base.Ui32(v371) < base.Ui32(int32(3)) {
		goto L112
	} else {
		goto L113
	}
L111:
	;
	if v491 != 0 {
		goto L96
	} else {
		goto L151
	}
L112:
	;
	v491 = int32(0)
	goto L111
L113:
	;
	goto L114
L114:
	;
	v382 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_index_build_range_scan[7]))
	if v382 == v371 {
		goto L115
	} else {
		goto L116
	}
L115:
	;
	v491 = int32(1)
	goto L111
L116:
	;
	goto L117
L117:
	;
	v386 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_index_build_range_scan[8]))
	if v386 <= int32(0) {
		goto L119
	} else {
		goto L120
	}
L118:
	;
	v491 = v483
	goto L111
L119:
	;
	v390 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_index_build_range_scan[9]))
	if v390 == int32(0) {
		v483 = int32(0)
		goto L118
	} else {
		goto L122
	}
L120:
	;
	goto L121
L121:
	;
	v452 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_index_build_range_scan[10]))
	v454 = int32(0)
	v456 = v386 - int32(1)
	goto L141
L122:
	;
	v395 = v390
	goto L123
L123:
	;
	v400 = *(*int32)(unsafe.Add(mBase, uint32(v395)+20))
	if v400 == int32(4) {
		goto L125
	} else {
		goto L126
	}
L124:
	;
	v483 = int32(0)
	goto L118
L125:
	;
	v447 = *(*int32)(unsafe.Add(mBase, uint32(v395)+80))
	if v447 != 0 {
		v395 = v447
		goto L123
	} else {
		goto L140
	}
L126:
	;
	v403 = *(*int32)(unsafe.Add(mBase, uint32(v395)))
	if v403 == int32(0) {
		goto L125
	} else {
		goto L127
	}
L127:
	;
	v406 = int32(1)
	if v371 == v403 {
		v483 = v406
		goto L118
	} else {
		goto L128
	}
L128:
	;
	v408 = *(*int32)(unsafe.Add(mBase, uint32(v395)+52))
	v410 = v408 - int32(1)
	if v410 < int32(0) {
		goto L125
	} else {
		goto L129
	}
L129:
	;
	v415 = int32(0)
	v417 = v410
	goto L130
L130:
	;
	v421 = *(*int32)(unsafe.Add(mBase, uint32(v395)+48))
	v423 = int32(2)
	v424 = base.I32_div_s(v417-v415, v423)
	v425 = v424 + v415
	v429 = *(*int32)(unsafe.Add(mBase, uint32(v421+v425<<(uint(v423)%32))))
	if v429 == v371 {
		v483 = v406
		goto L118
	} else {
		goto L132
	}
L131:
	;
	goto L125
L132:
	;
	v433 = F_TransactionIdPrecedes(m, v429, v371)
	mBase = m.M
	if v433 != 0 {
		goto L133
	} else {
		goto L134
	}
L133:
	;
	v434 = v425 + int32(1)
	goto L135
L134:
	;
	v434 = v415
	goto L135
L135:
	;
	if v433 != 0 {
		goto L136
	} else {
		goto L137
	}
L136:
	;
	v437 = v417
	goto L138
L137:
	;
	v437 = v425 - int32(1)
	goto L138
L138:
	;
	if v434 <= v437 {
		v415 = v434
		v417 = v437
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
	v461 = int32(2)
	v462 = base.I32_div_s(v456-v454, v461)
	v463 = v462 + v454
	v467 = *(*int32)(unsafe.Add(mBase, uint32(v452+v463<<(uint(v461)%32))))
	v468 = base.B2i32(v467 == v371)
	if v467 == v371 {
		v483 = v468
		goto L118
	} else {
		goto L143
	}
L142:
	;
	v483 = v468
	goto L118
L143:
	;
	v471 = base.B2i32(base.Ui32(v467) < base.Ui32(v371))
	if base.Ui32(v467) < base.Ui32(v371) {
		goto L144
	} else {
		goto L145
	}
L144:
	;
	v472 = v463 + int32(1)
	goto L146
L145:
	;
	v472 = v454
	goto L146
L146:
	;
	if base.Ui32(v467) < base.Ui32(v371) {
		goto L147
	} else {
		goto L148
	}
L147:
	;
	v475 = v456
	goto L149
L148:
	;
	v475 = v463 - int32(1)
	goto L149
L149:
	;
	if v472 <= v475 {
		v454 = v472
		v456 = v475
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
		v727 = v201
		goto L95
	} else {
		goto L158
	}
L153:
	;
	v494 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v495 = m.ExcPending
	if v495 != 0 {
		goto L8
	} else {
		goto L154
	}
L154:
	;
	if v494 == int32(0) {
		goto L152
	} else {
		goto L155
	}
L155:
	;
	v498 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v33)+16)) = v498 + int32(4)
	F_errmsg_internal(m, int32(_a_F_heapam_index_build_range_scan_4), v33+int32(16))
	mBase = m.M
	v506 = m.ExcPending
	if v506 != 0 {
		goto L8
	} else {
		goto L156
	}
L156:
	;
	F_errfinish(m, int32(_a_F_heapam_index_build_range_scan_5), int32(1484), int32(_a_F_heapam_index_build_range_scan_6))
	mBase = m.M
	v511 = m.ExcPending
	if v511 != 0 {
		goto L8
	} else {
		goto L157
	}
L157:
	;
	goto L152
L158:
	;
	v742 = v371
	goto L92
L159:
	;
	v514 = *(*int32)(unsafe.Add(mBase, uint32(v175)+16))
	v515 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v514)+20)))
	if v515&int32(_a_F_heapam_index_build_range_scan_7) == int32(_a_F_heapam_index_build_range_scan_8) {
		goto L161
	} else {
		goto L162
	}
L160:
	;
	if base.Ui32(v523) < base.Ui32(int32(3)) {
		goto L166
	} else {
		goto L167
	}
L161:
	;
	v520 = F_HeapTupleGetUpdateXid(m, v514)
	mBase = m.M
	v521 = m.ExcPending
	if v521 != 0 {
		goto L8
	} else {
		goto L164
	}
L162:
	;
	goto L163
L163:
	;
	v522 = *(*int32)(unsafe.Add(mBase, uint32(v514)+4))
	v523 = v522
	goto L160
L164:
	;
	v523 = v520
	goto L160
L165:
	;
	if v643 == int32(0) {
		goto L205
	} else {
		goto L206
	}
L166:
	;
	v643 = int32(0)
	goto L165
L167:
	;
	goto L168
L168:
	;
	v534 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_index_build_range_scan[7]))
	if v534 == v523 {
		goto L169
	} else {
		goto L170
	}
L169:
	;
	v643 = int32(1)
	goto L165
L170:
	;
	goto L171
L171:
	;
	v538 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_index_build_range_scan[8]))
	if v538 <= int32(0) {
		goto L173
	} else {
		goto L174
	}
L172:
	;
	v643 = v635
	goto L165
L173:
	;
	v542 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_index_build_range_scan[9]))
	if v542 == int32(0) {
		v635 = int32(0)
		goto L172
	} else {
		goto L176
	}
L174:
	;
	goto L175
L175:
	;
	v604 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_index_build_range_scan[10]))
	v606 = int32(0)
	v608 = v538 - int32(1)
	goto L195
L176:
	;
	v547 = v542
	goto L177
L177:
	;
	v552 = *(*int32)(unsafe.Add(mBase, uint32(v547)+20))
	if v552 == int32(4) {
		goto L179
	} else {
		goto L180
	}
L178:
	;
	v635 = int32(0)
	goto L172
L179:
	;
	v599 = *(*int32)(unsafe.Add(mBase, uint32(v547)+80))
	if v599 != 0 {
		v547 = v599
		goto L177
	} else {
		goto L194
	}
L180:
	;
	v555 = *(*int32)(unsafe.Add(mBase, uint32(v547)))
	if v555 == int32(0) {
		goto L179
	} else {
		goto L181
	}
L181:
	;
	v558 = int32(1)
	if v523 == v555 {
		v635 = v558
		goto L172
	} else {
		goto L182
	}
L182:
	;
	v560 = *(*int32)(unsafe.Add(mBase, uint32(v547)+52))
	v562 = v560 - int32(1)
	if v562 < int32(0) {
		goto L179
	} else {
		goto L183
	}
L183:
	;
	v567 = int32(0)
	v569 = v562
	goto L184
L184:
	;
	v573 = *(*int32)(unsafe.Add(mBase, uint32(v547)+48))
	v575 = int32(2)
	v576 = base.I32_div_s(v569-v567, v575)
	v577 = v576 + v567
	v581 = *(*int32)(unsafe.Add(mBase, uint32(v573+v577<<(uint(v575)%32))))
	if v581 == v523 {
		v635 = v558
		goto L172
	} else {
		goto L186
	}
L185:
	;
	goto L179
L186:
	;
	v585 = F_TransactionIdPrecedes(m, v581, v523)
	mBase = m.M
	if v585 != 0 {
		goto L187
	} else {
		goto L188
	}
L187:
	;
	v586 = v577 + int32(1)
	goto L189
L188:
	;
	v586 = v567
	goto L189
L189:
	;
	if v585 != 0 {
		goto L190
	} else {
		goto L191
	}
L190:
	;
	v589 = v569
	goto L192
L191:
	;
	v589 = v577 - int32(1)
	goto L192
L192:
	;
	if v586 <= v589 {
		v567 = v586
		v569 = v589
		goto L184
	} else {
		goto L193
	}
L193:
	;
	goto L185
L194:
	;
	goto L178
L195:
	;
	v613 = int32(2)
	v614 = base.I32_div_s(v608-v606, v613)
	v615 = v614 + v606
	v619 = *(*int32)(unsafe.Add(mBase, uint32(v604+v615<<(uint(v613)%32))))
	v620 = base.B2i32(v619 == v523)
	if v619 == v523 {
		v635 = v620
		goto L172
	} else {
		goto L197
	}
L196:
	;
	v635 = v620
	goto L172
L197:
	;
	v623 = base.B2i32(base.Ui32(v619) < base.Ui32(v523))
	if base.Ui32(v619) < base.Ui32(v523) {
		goto L198
	} else {
		goto L199
	}
L198:
	;
	v624 = v615 + int32(1)
	goto L200
L199:
	;
	v624 = v606
	goto L200
L200:
	;
	if base.Ui32(v619) < base.Ui32(v523) {
		goto L201
	} else {
		goto L202
	}
L201:
	;
	v627 = v608
	goto L203
L202:
	;
	v627 = v615 - int32(1)
	goto L203
L203:
	;
	if v624 <= v627 {
		v606 = v624
		v608 = v627
		goto L195
	} else {
		goto L204
	}
L204:
	;
	goto L196
L205:
	;
	if v46 != 0 {
		goto L208
	} else {
		goto L209
	}
L206:
	;
	goto L207
L207:
	;
	v683 = int32(0)
	v684 = *(*int32)(unsafe.Add(mBase, uint32(v175)+16))
	v685 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v684)+19)))
	if v685&int32(64) == v683 {
		v735 = v683
		v737 = v201
		goto L93
	} else {
		goto L219
	}
L208:
	;
	if v54 == int32(0) {
		goto L214
	} else {
		goto L215
	}
L209:
	;
	v648 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v649 = m.ExcPending
	if v649 != 0 {
		goto L8
	} else {
		goto L210
	}
L210:
	;
	if v648 == int32(0) {
		goto L208
	} else {
		goto L211
	}
L211:
	;
	v652 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v33)+32)) = v652 + int32(4)
	F_errmsg_internal(m, int32(_a_F_heapam_index_build_range_scan_9), v33+int32(32))
	mBase = m.M
	v660 = m.ExcPending
	if v660 != 0 {
		goto L8
	} else {
		goto L212
	}
L212:
	;
	F_errfinish(m, int32(_a_F_heapam_index_build_range_scan_5), int32(1543), int32(_a_F_heapam_index_build_range_scan_6))
	mBase = m.M
	v665 = m.ExcPending
	if v665 != 0 {
		goto L8
	} else {
		goto L213
	}
L213:
	;
	goto L208
L214:
	;
	v668 = *(*int32)(unsafe.Add(mBase, uint32(v175)+16))
	v669 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v668)+19)))
	if v669&int32(64) == int32(0) {
		goto L94
	} else {
		goto L217
	}
L215:
	;
	goto L216
L216:
	;
	v742 = v523
	goto L92
L217:
	;
	v674 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v668)+20)))
	if v674&int32(2048)|base.B2i32(v674&int32(768) == int32(512)) != 0 {
		goto L94
	} else {
		goto L218
	}
L218:
	;
	goto L216
L219:
	;
	v690 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v684)+20)))
	if v690&int32(2048)|base.B2i32(v690&int32(768) == int32(512)) != 0 {
		v735 = v683
		v737 = v201
		goto L93
	} else {
		goto L220
	}
L220:
	;
	goto L99
L221:
	;
	v916 = v201
	goto L84
L222:
	;
	F_errmsg_internal(m, int32(_a_F_heapam_index_build_range_scan_10), int32(0))
	mBase = m.M
	v715 = m.ExcPending
	if v715 != 0 {
		goto L8
	} else {
		goto L223
	}
L223:
	;
	F_errfinish(m, int32(_a_F_heapam_index_build_range_scan_5), int32(1613), int32(_a_F_heapam_index_build_range_scan_6))
	mBase = m.M
	v720 = m.ExcPending
	if v720 != 0 {
		goto L8
	} else {
		goto L224
	}
L224:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L225:
	;
	v767 = v735
	v789 = v737
	goto L85
L226:
	;
	F_XactLockTableWait(m, v742, l0, v175+int32(4), int32(6))
	mBase = m.M
	v750 = m.ExcPending
	if v750 != 0 {
		goto L8
	} else {
		goto L227
	}
L227:
	;
	v752 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_index_build_range_scan[4]))
	if v752 == int32(0) {
		goto L89
	} else {
		goto L228
	}
L228:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v756 = m.ExcPending
	if v756 != 0 {
		goto L8
	} else {
		goto L229
	}
L229:
	;
	goto L89
L230:
	;
	v793 = *(*int32)(unsafe.Add(mBase, uint32(v103)+56))
	F_ExecStoreBufferHeapTuple(m, v175, v66, v793)
	mBase = m.M
	v795 = m.ExcPending
	if v795 != 0 {
		goto L8
	} else {
		goto L231
	}
L231:
	;
	if v70 != 0 {
		goto L232
	} else {
		goto L233
	}
L232:
	;
	v796 = int32(_a_F_heapam_index_build_range_scan_11)
	v797 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_index_build_range_scan[11]))
	v799 = *(*int32)(unsafe.Add(mBase, uint32(v64)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_heapam_index_build_range_scan[11])) = v799
	v803 = *(*int32)(unsafe.Add(mBase, uint32(v70)+20))
	v804 = m.T0[v803].(func(*base.Module, int32, int32, int32) int32)(m, v70, v64, v33+int32(42))
	mBase = m.M
	v805 = m.ExcPending
	if v805 != 0 {
		goto L8
	} else {
		goto L235
	}
L233:
	;
	goto L234
L234:
	;
	F_FormIndexDatum(m, l2, v66, v55, v33+int32(672), v33+int32(640))
	mBase = m.M
	v816 = m.ExcPending
	if v816 != 0 {
		goto L8
	} else {
		goto L237
	}
L235:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_heapam_index_build_range_scan[11])) = v797
	if v804 == int32(0) {
		v916 = v789
		goto L84
	} else {
		goto L236
	}
L236:
	;
	goto L234
L237:
	;
	v817 = *(*int32)(unsafe.Add(mBase, uint32(v175)+16))
	v818 = int32(*(*int16)(unsafe.Add(mBase, uint32(v817)+18)))
	if v818 < int32(0) {
		goto L238
	} else {
		goto L239
	}
L238:
	;
	v821 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v175)+8)))
	v826 = v821<<(uint(int32(1))%32) + v33 + int32(46)
	v827 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v826))))
	if v827 == int32(0) {
		goto L241
	} else {
		goto L242
	}
L239:
	;
	goto L240
L240:
	;
	m.T0[l8].(func(*base.Module, int32, int32, int32, int32, int32, int32))(m, l1, v175+int32(4), v33+int32(672), v33+int32(640), v767, l9)
	mBase = m.M
	v886 = m.ExcPending
	if v886 != 0 {
		goto L8
	} else {
		goto L253
	}
L241:
	;
	v830 = *(*int32)(unsafe.Add(mBase, uint32(v103)+56))
	if v830 < int32(0) {
		goto L245
	} else {
		goto L246
	}
L242:
	;
	v861 = v827
	goto L243
L243:
	;
	if base.Ui32(int32(2048)) <= base.Ui32((v861-int32(1))&int32(_a_F_heapam_index_build_range_scan_12)) {
		goto L46
	} else {
		goto L251
	}
L244:
	;
	F_LockBuffer(m, v830, int32(1))
	mBase = m.M
	v851 = m.ExcPending
	if v851 != 0 {
		goto L8
	} else {
		goto L248
	}
L245:
	;
	v834 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_index_build_range_scan[5]))
	v840 = *(*int32)(unsafe.Add(mBase, uint32(v834+(v830^int32(-1))<<(uint(int32(2))%32))))
	v848 = v840
	goto L244
L246:
	;
	goto L247
L247:
	;
	v842 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_index_build_range_scan[6]))
	v848 = v842 + v830<<(uint(int32(13))%32) + int32(-8192)
	goto L244
L248:
	;
	F_heap_get_root_tuples(m, v848, v33+int32(48))
	mBase = m.M
	v855 = m.ExcPending
	if v855 != 0 {
		goto L8
	} else {
		goto L249
	}
L249:
	;
	v856 = *(*int32)(unsafe.Add(mBase, uint32(v103)+56))
	F_LockBuffer(m, v856, int32(0))
	mBase = m.M
	v859 = m.ExcPending
	if v859 != 0 {
		goto L8
	} else {
		goto L250
	}
L250:
	;
	v860 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v826))))
	v861 = v860
	goto L243
L251:
	;
	v868 = *(*int32)(unsafe.Add(mBase, uint32(v175)+4))
	*(*uint16)(unsafe.Add(mBase, uint32(v33)+46)) = uint16(v861)
	*(*int32)(unsafe.Add(mBase, uint32(v33)+42)) = v868
	m.T0[l8].(func(*base.Module, int32, int32, int32, int32, int32, int32))(m, l1, v33+int32(42), v33+int32(672), v33+int32(640), v767, l9)
	mBase = m.M
	v878 = m.ExcPending
	if v878 != 0 {
		goto L8
	} else {
		goto L252
	}
L252:
	;
	v916 = v789
	goto L84
L253:
	;
	v916 = v789
	goto L84
L254:
	;
	if v917 != 0 {
		v175 = v917
		v189 = v273
		v193 = v307
		v201 = v916
		goto L51
	} else {
		goto L255
	}
L255:
	;
	goto L52
L256:
	;
	v950 = *(*int32)(unsafe.Add(mBase, uint32(v103)+32))
	if v950 != 0 {
		goto L259
	} else {
		goto L260
	}
L257:
	;
	goto L258
L258:
	;
	v1001 = *(*int32)(unsafe.Add(mBase, uint32(v103)))
	v1002 = *(*int32)(unsafe.Add(mBase, uint32(v1001)+188))
	v1003 = *(*int32)(unsafe.Add(mBase, uint32(v1002)+12))
	m.T0[v1003].(func(*base.Module, int32))(m, v103)
	mBase = m.M
	v1005 = m.ExcPending
	if v1005 != 0 {
		goto L8
	} else {
		goto L266
	}
L259:
	;
	v955 = v950 + int32(20)
	goto L261
L260:
	;
	v955 = v103 + int32(36)
	goto L261
L261:
	;
	v956 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v955))))
	v959 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_index_build_range_scan[1]))
	if v959 == int32(0) {
		goto L263
	} else {
		goto L264
	}
L262:
	;
	goto L258
L263:
	;
	goto L262
L264:
	;
	v963 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_heapam_index_build_range_scan[2])))
	if v963&int32(1) == int32(0) {
		goto L263
	} else {
		goto L265
	}
L265:
	;
	v968 = int32(_a_F_heapam_index_build_range_scan_2)
	v970 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_index_build_range_scan[3]))
	v971 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_heapam_index_build_range_scan[3])) = v970 + v971
	v974 = *(*int32)(unsafe.Add(mBase, uint32(v959)))
	*(*int32)(unsafe.Add(mBase, uint32(v959))) = v974 + v971
	v978 = int32(0)
	v980 = int32(_a_F_heapam_index_build_range_scan_3)
	v981 = base.AtomicRmwOr32(m, v978, v980, v978)
	*(*int64)(unsafe.Add(mBase, uint32(v959+int32(128))+232)) = v956
	v989 = base.AtomicRmwOr32(m, v978, v980, v978)
	v990 = *(*int32)(unsafe.Add(mBase, uint32(v959)))
	*(*int32)(unsafe.Add(mBase, uint32(v959))) = v990 + v971
	v996 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_index_build_range_scan[3]))
	*(*int32)(unsafe.Add(mBase, _c_F_heapam_index_build_range_scan[3])) = v996 - v971
	goto L263
L266:
	;
	if v105 != 0 {
		goto L267
	} else {
		goto L268
	}
L267:
	;
	F_UnregisterSnapshot(m, v104)
	mBase = m.M
	v1007 = m.ExcPending
	if v1007 != 0 {
		goto L8
	} else {
		goto L270
	}
L268:
	;
	goto L269
L269:
	;
	F_ExecDropSingleTupleTableSlot(m, v66)
	mBase = m.M
	v1009 = m.ExcPending
	if v1009 != 0 {
		goto L8
	} else {
		goto L271
	}
L270:
	;
	goto L269
L271:
	;
	F_FreeExecutorState(m, v55)
	mBase = m.M
	v1011 = m.ExcPending
	if v1011 != 0 {
		goto L8
	} else {
		goto L272
	}
L272:
	;
	v1012 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l2)+88)) = v1012
	*(*int32)(unsafe.Add(mBase, uint32(l2)+80)) = v1012
	m.G0 = v33 + int32(800)
	return v948
L273:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v1026 = m.ExcPending
	if v1026 != 0 {
		goto L8
	} else {
		goto L274
	}
L274:
	;
	v1027 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v175)+6)))
	v1028 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v175)+4)))
	v1029 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v33)+4)) = v821
	*(*int32)(unsafe.Add(mBase, uint32(v33)+8)) = v1029 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v33))) = v1027 | v1028<<(uint(int32(16))%32)
	F_errmsg_internal(m, int32(_a_F_heapam_index_build_range_scan_13), v33)
	mBase = m.M
	v1040 = m.ExcPending
	if v1040 != 0 {
		goto L8
	} else {
		goto L275
	}
L275:
	;
	F_errfinish(m, int32(_a_F_heapam_index_build_range_scan_5), int32(1693), int32(_a_F_heapam_index_build_range_scan_6))
	mBase = m.M
	v1045 = m.ExcPending
	if v1045 != 0 {
		goto L8
	} else {
		goto L276
	}
L276:
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
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v116 float64
	_ = v116
	var v120 int32
	_ = v120
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
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
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v234 int32
	_ = v234
	var v240 int32
	_ = v240
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v289 int32
	_ = v289
	var v290 int64
	_ = v290
	var v293 int64
	_ = v293
	var v296 int64
	_ = v296
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v331 int32
	_ = v331
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v342 int32
	_ = v342
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v351 int32
	_ = v351
	var v353 int32
	_ = v353
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v367 int32
	_ = v367
	var v368 int64
	_ = v368
	var v371 int64
	_ = v371
	var v374 int64
	_ = v374
	var v403 int32
	_ = v403
	var v415 int32
	_ = v415
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v433 int32
	_ = v433
	var v438 int32
	_ = v438
	var v448 int32
	_ = v448
	var v456 int32
	_ = v456
	var v458 int32
	_ = v458
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v471 int32
	_ = v471
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v482 int32
	_ = v482
	var v494 int32
	_ = v494
	var v500 int32
	_ = v500
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v506 int32
	_ = v506
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v513 int32
	_ = v513
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v526 int32
	_ = v526
	var v528 int32
	_ = v528
	var v530 int32
	_ = v530
	var v533 int32
	_ = v533
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v537 float64
	_ = v537
	var v550 int32
	_ = v550
	var v556 int32
	_ = v556
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v579 int32
	_ = v579
	var v581 int32
	_ = v581
	var v583 int32
	_ = v583
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
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
	v32 = F_MakeTupleTableSlot(m, v30, int32(_a_F_heapam_index_validate_scan_0))
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
	v50 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_index_validate_scan[0]))
	if v50 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v91 = F_heap_getnext(m, v45)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L1
	} else {
		goto L14
	}
L11:
	;
	goto L10
L12:
	;
	v54 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_heapam_index_validate_scan[1])))
	if v54&int32(1) == int32(0) {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v59 = int32(_a_F_heapam_index_validate_scan_1)
	v61 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_index_validate_scan[2]))
	v62 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_heapam_index_validate_scan[2])) = v61 + v62
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
	*(*int32)(unsafe.Add(mBase, uint32(v50))) = v65 + v62
	v69 = int32(0)
	v71 = int32(_a_F_heapam_index_validate_scan_2)
	v72 = base.AtomicRmwOr32(m, v69, v71, v69)
	*(*int64)(unsafe.Add(mBase, uint32(v50+int32(120))+232)) = v47
	v80 = base.AtomicRmwOr32(m, v69, v71, v69)
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
	*(*int32)(unsafe.Add(mBase, uint32(v50))) = v81 + v62
	v87 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_index_validate_scan[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_heapam_index_validate_scan[2])) = v87 - v62
	goto L11
L14:
	;
	if v91 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v93 = int32(-1)
	v98 = v91
	v103 = v93
	v104 = v6
	v107 = v93
	v110 = v6
	goto L18
L16:
	;
	goto L17
L17:
	;
	v577 = *(*int32)(unsafe.Add(mBase, uint32(v45)))
	v578 = *(*int32)(unsafe.Add(mBase, uint32(v577)+188))
	v579 = *(*int32)(unsafe.Add(mBase, uint32(v578)+12))
	m.T0[v579].(func(*base.Module, int32))(m, v45)
	mBase = m.M
	v581 = m.ExcPending
	if v581 != 0 {
		goto L1
	} else {
		goto L99
	}
L18:
	;
	v113 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_index_validate_scan[3]))
	if v113 != 0 {
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
	v115 = m.ExcPending
	if v115 != 0 {
		goto L1
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v116 = *(*float64)(unsafe.Add(mBase, uint32(l4)+8))
	*(*float64)(unsafe.Add(mBase, uint32(l4)+8)) = base.F64_add(v116, float64(1))
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v45)+52))
	if base.B2i32(v120 == v107)&base.B2i32(v107 != int32(-1)) == int32(0) {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	goto L22
L24:
	;
	v131 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_index_validate_scan[0]))
	if v131 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L25:
	;
	v173 = v107
	goto L26
L26:
	;
	if v103 != v173 {
		goto L31
	} else {
		goto L32
	}
L27:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v45)+52))
	v173 = v172
	goto L26
L28:
	;
	goto L27
L29:
	;
	v135 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_heapam_index_validate_scan[1])))
	if v135&int32(1) == int32(0) {
		goto L28
	} else {
		goto L30
	}
L30:
	;
	v140 = int32(_a_F_heapam_index_validate_scan_1)
	v142 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_index_validate_scan[2]))
	v143 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_heapam_index_validate_scan[2])) = v142 + v143
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v131)))
	*(*int32)(unsafe.Add(mBase, uint32(v131))) = v146 + v143
	v150 = int32(0)
	v152 = int32(_a_F_heapam_index_validate_scan_2)
	v153 = base.AtomicRmwOr32(m, v150, v152, v150)
	*(*int64)(unsafe.Add(mBase, uint32(v131+int32(128))+232)) = base.I64_extend_i32_u(v120)
	v161 = base.AtomicRmwOr32(m, v150, v152, v150)
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v131)))
	*(*int32)(unsafe.Add(mBase, uint32(v131))) = v162 + v143
	v168 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_index_validate_scan[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_heapam_index_validate_scan[2])) = v168 - v143
	goto L28
L31:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v45)+56))
	if v175 < int32(0) {
		goto L35
	} else {
		goto L36
	}
L32:
	;
	v211 = v103
	goto L33
L33:
	;
	v213 = v98 + int32(4)
	v214 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v213)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v20)+24)) = uint16(v214)
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v213)))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+20)) = v216
	v218 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v98)+8)))
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v98)+16))
	v220 = int32(*(*int16)(unsafe.Add(mBase, uint32(v219)+18)))
	if v220 < int32(0) {
		goto L45
	} else {
		goto L46
	}
L34:
	;
	F_LockBuffer(m, v175, int32(1))
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L1
	} else {
		goto L38
	}
L35:
	;
	v179 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_index_validate_scan[4]))
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v179+(v175^int32(-1))<<(uint(int32(2))%32))))
	v193 = v185
	goto L34
L36:
	;
	goto L37
L37:
	;
	v187 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_index_validate_scan[5]))
	v193 = v187 + v175<<(uint(int32(13))%32) + int32(-8192)
	goto L34
L38:
	;
	F_heap_get_root_tuples(m, v193, v20+int32(336))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v45)+56))
	F_LockBuffer(m, v201, int32(0))
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	base.MemoryFill(m, v20+int32(32), int32(0), int32(291))
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v45)+52))
	v211 = v210
	goto L33
L41:
	;
	v558 = F_heap_getnext(m, v45)
	mBase = m.M
	v559 = m.ExcPending
	if v559 != 0 {
		goto L1
	} else {
		goto L97
	}
L42:
	;
	v503 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20+v234)+31)))
	if v503 != 0 {
		v550 = v494
		v556 = v500
		goto L41
	} else {
		goto L87
	}
L43:
	;
	v456 = int32(0)
	v458 = v20 + int32(20)
	v462 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v448)+2)))
	v463 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v448))))
	v464 = int32(16)
	v466 = v462 | v463<<(uint(v464)%32)
	v467 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v458)+2)))
	v468 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v458))))
	v471 = v467 | v468<<(uint(v464)%32)
	if base.Ui32(v466) < base.Ui32(v471) {
		v482 = int32(-1)
		goto L82
	} else {
		goto L83
	}
L44:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v415 = m.ExcPending
	if v415 != 0 {
		goto L1
	} else {
		goto L77
	}
L45:
	;
	v223 = int32(1)
	v226 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v218<<(uint(v223)%32)+v20)+334)))
	if base.Ui32(int32(2048)) <= base.Ui32((v226-v223)&int32(_a_F_heapam_index_validate_scan_3)) {
		goto L44
	} else {
		goto L48
	}
L46:
	;
	v234 = v218
	goto L47
L47:
	;
	if v110 == int32(0) {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v20)+24)) = uint16(v226)
	v234 = v226
	goto L47
L49:
	;
	if v104 == int32(0) {
		goto L52
	} else {
		goto L53
	}
L50:
	;
	v403 = v104
	goto L51
L51:
	;
	v494 = v403
	v500 = int32(1)
	goto L42
L52:
	;
	v277 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	v279 = int32(0)
	v285 = F_tuplesort_getdatum(m, v277, int32(1), v279, v20+int32(16), v20+int32(15), v279)
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L1
	} else {
		goto L62
	}
L53:
	;
	v240 = v20 + int32(20)
	v244 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v104)+2)))
	v245 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v104))))
	v246 = int32(16)
	v248 = v244 | v245<<(uint(v246)%32)
	v249 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v240)+2)))
	v250 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v240))))
	v253 = v249 | v250<<(uint(v246)%32)
	if base.Ui32(v248) < base.Ui32(v253) {
		v264 = int32(-1)
		goto L55
	} else {
		goto L56
	}
L54:
	;
	if int32(0) <= v264 {
		v448 = v104
		goto L43
	} else {
		goto L59
	}
L55:
	;
	goto L54
L56:
	;
	if base.Ui32(v253) < base.Ui32(v248) {
		v264 = int32(1)
		goto L55
	} else {
		goto L57
	}
L57:
	;
	v258 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v104)+4)))
	v259 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v240)+4)))
	if base.Ui32(v258) < base.Ui32(v259) {
		v264 = int32(-1)
		goto L55
	} else {
		goto L58
	}
L58:
	;
	v264 = base.B2i32(base.Ui32(v259) < base.Ui32(v258))
	goto L55
L59:
	;
	v267 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v104)+2)))
	v268 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v104))))
	if v267|v268<<(uint(int32(16))%32) != v211 {
		goto L52
	} else {
		goto L60
	}
L60:
	;
	v273 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v104)+4)))
	v275 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v20+v273)+31)) = uint8(v275)
	goto L52
L61:
	;
	v403 = int32(0)
	goto L51
L62:
	;
	if v285 == int32(0) {
		goto L61
	} else {
		goto L63
	}
L63:
	;
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v20)+16))
	v290 = *(*int64)(unsafe.Add(mBase, uint32(v289)))
	*(*uint16)(unsafe.Add(mBase, uint32(v20)+30)) = uint16(v290)
	v293 = int64(base.Ui64(v290) >> (uint(int64(16)) % 64))
	*(*uint16)(unsafe.Add(mBase, uint32(v20)+28)) = uint16(v293)
	v296 = int64(base.Ui64(v290) >> (uint(int64(32)) % 64))
	*(*uint16)(unsafe.Add(mBase, uint32(v20)+26)) = uint16(v296)
	goto L64
L64:
	;
	v316 = v20 + int32(26)
	v318 = v20 + int32(20)
	v322 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v316)+2)))
	v323 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v316))))
	v324 = int32(16)
	v326 = v322 | v323<<(uint(v324)%32)
	v327 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v318)+2)))
	v328 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v318))))
	v331 = v327 | v328<<(uint(v324)%32)
	if base.Ui32(v326) < base.Ui32(v331) {
		v342 = int32(-1)
		goto L67
	} else {
		goto L68
	}
L66:
	;
	if int32(0) <= v342 {
		v448 = v316
		goto L43
	} else {
		goto L71
	}
L67:
	;
	goto L66
L68:
	;
	if base.Ui32(v331) < base.Ui32(v326) {
		v342 = int32(1)
		goto L67
	} else {
		goto L69
	}
L69:
	;
	v336 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v316)+4)))
	v337 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v318)+4)))
	if base.Ui32(v336) < base.Ui32(v337) {
		v342 = int32(-1)
		goto L67
	} else {
		goto L70
	}
L70:
	;
	v342 = base.B2i32(base.Ui32(v337) < base.Ui32(v336))
	goto L67
L71:
	;
	v345 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v20)+28)))
	v346 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v20)+26)))
	if v211 == v345|v346<<(uint(int32(16))%32) {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v351 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v20)+30)))
	v353 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v20+v351)+31)) = uint8(v353)
	goto L74
L73:
	;
	goto L74
L74:
	;
	v355 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	v357 = int32(0)
	v363 = F_tuplesort_getdatum(m, v355, int32(1), v357, v20+int32(16), v20+int32(15), v357)
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L1
	} else {
		goto L75
	}
L75:
	;
	if v363 == int32(0) {
		goto L61
	} else {
		goto L76
	}
L76:
	;
	v367 = *(*int32)(unsafe.Add(mBase, uint32(v20)+16))
	v368 = *(*int64)(unsafe.Add(mBase, uint32(v367)))
	*(*uint16)(unsafe.Add(mBase, uint32(v20)+30)) = uint16(v368)
	v371 = int64(base.Ui64(v368) >> (uint(int64(16)) % 64))
	*(*uint16)(unsafe.Add(mBase, uint32(v20)+28)) = uint16(v371)
	v374 = int64(base.Ui64(v368) >> (uint(int64(32)) % 64))
	*(*uint16)(unsafe.Add(mBase, uint32(v20)+26)) = uint16(v374)
	goto L64
L77:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v418 = m.ExcPending
	if v418 != 0 {
		goto L1
	} else {
		goto L78
	}
L78:
	;
	v419 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v98)+6)))
	v420 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v98)+4)))
	v421 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v422 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v98)+8)))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+4)) = v422
	*(*int32)(unsafe.Add(mBase, uint32(v20)+8)) = v421 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v419 | v420<<(uint(int32(16))%32)
	F_errmsg_internal(m, int32(_a_F_heapam_index_validate_scan_4), v20)
	mBase = m.M
	v433 = m.ExcPending
	if v433 != 0 {
		goto L1
	} else {
		goto L79
	}
L79:
	;
	F_errfinish(m, int32(_a_F_heapam_index_validate_scan_5), int32(1870), int32(_a_F_heapam_index_validate_scan_6))
	mBase = m.M
	v438 = m.ExcPending
	if v438 != 0 {
		goto L1
	} else {
		goto L80
	}
L80:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L81:
	;
	if v482 <= int32(0) {
		v550 = v448
		v556 = v456
		goto L41
	} else {
		goto L86
	}
L82:
	;
	goto L81
L83:
	;
	if base.Ui32(v471) < base.Ui32(v466) {
		v482 = int32(1)
		goto L82
	} else {
		goto L84
	}
L84:
	;
	v476 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v448)+4)))
	v477 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v458)+4)))
	if base.Ui32(v476) < base.Ui32(v477) {
		v482 = int32(-1)
		goto L82
	} else {
		goto L85
	}
L85:
	;
	v482 = base.B2i32(base.Ui32(v477) < base.Ui32(v476))
	goto L82
L86:
	;
	v494 = v448
	v500 = v456
	goto L42
L87:
	;
	v504 = *(*int32)(unsafe.Add(mBase, uint32(v29)+20))
	F_MemoryContextReset(m, v504)
	mBase = m.M
	v506 = m.ExcPending
	if v506 != 0 {
		goto L1
	} else {
		goto L88
	}
L88:
	;
	v508 = F_ExecStoreHeapTuple(m, v98, v32, int32(0))
	mBase = m.M
	v509 = m.ExcPending
	if v509 != 0 {
		goto L1
	} else {
		goto L89
	}
L89:
	;
	if v36 != 0 {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v510 = int32(_a_F_heapam_index_validate_scan_7)
	v511 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_index_validate_scan[6]))
	v513 = *(*int32)(unsafe.Add(mBase, uint32(v29)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_heapam_index_validate_scan[6])) = v513
	v517 = *(*int32)(unsafe.Add(mBase, uint32(v36)+20))
	v518 = m.T0[v517].(func(*base.Module, int32, int32, int32) int32)(m, v36, v29, v20+int32(16))
	mBase = m.M
	v519 = m.ExcPending
	if v519 != 0 {
		goto L1
	} else {
		goto L93
	}
L91:
	;
	goto L92
L92:
	;
	v526 = v20 + int32(960)
	v528 = v20 + int32(928)
	F_FormIndexDatum(m, l2, v32, v22, v526, v528)
	mBase = m.M
	v530 = m.ExcPending
	if v530 != 0 {
		goto L1
	} else {
		goto L95
	}
L93:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_heapam_index_validate_scan[6])) = v511
	if v518 == int32(0) {
		v550 = v494
		v556 = v500
		goto L41
	} else {
		goto L94
	}
L94:
	;
	goto L92
L95:
	;
	v533 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+116)))
	v535 = F_index_insert(m, l1, v526, v528, v20+int32(20), l0, v533, int32(0), l2)
	mBase = m.M
	v536 = m.ExcPending
	if v536 != 0 {
		goto L1
	} else {
		goto L96
	}
L96:
	;
	v537 = *(*float64)(unsafe.Add(mBase, uint32(l4)+24))
	*(*float64)(unsafe.Add(mBase, uint32(l4)+24)) = base.F64_add(v537, float64(1))
	v550 = v494
	v556 = v500
	goto L41
L97:
	;
	if v558 != 0 {
		v98 = v558
		v103 = v211
		v104 = v550
		v107 = v173
		v110 = v556
		goto L18
	} else {
		goto L98
	}
L98:
	;
	goto L19
L99:
	;
	F_ExecDropSingleTupleTableSlot(m, v32)
	mBase = m.M
	v583 = m.ExcPending
	if v583 != 0 {
		goto L1
	} else {
		goto L100
	}
L100:
	;
	F_FreeExecutorState(m, v22)
	mBase = m.M
	v585 = m.ExcPending
	if v585 != 0 {
		goto L1
	} else {
		goto L101
	}
L101:
	;
	v586 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l2)+88)) = v586
	*(*int32)(unsafe.Add(mBase, uint32(l2)+80)) = v586
	m.G0 = v20 + int32(1088)
	return
}
func F_heapam_relation_needs_toast_table(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v76 int32
	_ = v76
	var v90 int32
	_ = v90
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
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v146 int32
	_ = v146
	v2 = int32(0)
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	if v10 <= v2 {
		v146 = v2
	} else {
		v14 = int32(0)
		v15 = v2
		v16 = v10
		v18 = v2
		v20 = v2
		for {
			v27 = v9 + v16<<(uint(int32(4))%32) + v14*int32(100)
			v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+111)))
			if v28 != 0 {
				v113 = v15
				v114 = v16
				v116 = v18
				v117 = v20
			} else {
				v30 = v27 + int32(20)
				v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+90)))
				if v31 == int32(118) {
					v113 = v15
					v114 = v16
					v116 = v18
					v117 = v20
				} else {
					v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+83)))
					switch v34 - int32(99) {
					case 0:
						v49 = v15
					case 1:
						v49 = (v15 + int32(7)) & int32(-8)
					default:
						v49 = (v15 + int32(1)) & int32(-2)
					case 6:
						v49 = (v15 + int32(3)) & int32(-4)
					}
					v50 = int32(*(*int16)(unsafe.Add(mBase, uint32(v30)+72)))
					if int32(0) < v50 {
						v113 = v49 + v50
						v114 = v16
						v116 = v18
						v117 = v20
					} else {
						v54 = *(*int32)(unsafe.Add(mBase, uint32(v30)+68))
						v55 = *(*int32)(unsafe.Add(mBase, uint32(v30)+76))
						v57 = int32(-1)
						if v55 < int32(0) {
							v97 = v57
							v99 = v97
						} else {
							if base.Ui32(int32(2)) <= base.Ui32(v54-int32(1042)) {
								switch v54 - int32(1560) {
								case 0, 2:
									v93 = int32(8)
									v94 = base.I32_div_s(v55+int32(7), v93)
									v97 = v94 + v93
									v99 = v97
								case 1:
									v97 = v57
									v99 = v97
								default:
									if v54 != int32(1700) {
										v97 = v57
										v99 = v97
									} else {
										v76 = int32(4)
										if v55 < v76 {
											v90 = int32(-1)
										} else {
											v90 = int32(base.Ui32(int32(base.Ui32(v55-v76)>>(uint(int32(16))%32))+int32(6))>>(uint(int32(1))%32))&int32(_a_F_heapam_relation_needs_toast_table_0) + int32(8)
										}
										v99 = v90
									}
								}
							} else {
								v66 = F_GetDatabaseEncoding(m)
								mBase = m.M
								v67 = F_pg_encoding_max_length(m, v66)
								mBase = m.M
								v68 = int32(4)
								v99 = v67*(v55-v68) + v68
							}
						}
						v100 = int32(0)
						if v100 < v99 {
							v103 = v99
						} else {
							v103 = v100
						}
						v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+84)))
						v112 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
						v113 = v103 + v49
						v114 = v112
						v116 = base.B2i32(v108 != int32(112)) | v18
						v117 = base.B2i32(v99 < int32(0)) | v20
					}
				}
			}
			v120 = v14 + int32(1)
			if v120 < v114 {
				v14 = v120
				v15 = v113
				v16 = v114
				v18 = v116
				v20 = v117
				continue
			} else {
				break
			}
			break
		}
		if (v116^int32(-1)|v117)&int32(1) != 0 {
			v146 = v116
		} else {
			v127 = int32(7)
			v130 = base.I32_div_s(v114+v127, int32(8))
			v133 = int32(-8)
			v146 = base.B2i32(base.Ui32(int32(2032)) < base.Ui32((v130+int32(30))&v133+(v113+v127)&v133))
		}
	}
	return v146 & int32(1)
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
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v214 int32
	_ = v214
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v243 int32
	_ = v243
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
	var v256 int32
	_ = v256
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v290 int32
	_ = v290
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v308 int32
	_ = v308
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v336 int32
	_ = v336
	var v344 int32
	_ = v344
	var v352 int32
	_ = v352
	var v354 float64
	_ = v354
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v369 float64
	_ = v369
	var v375 int32
	_ = v375
	var v377 int32
	_ = v377
	var v380 int32
	_ = v380
	var v393 int32
	_ = v393
	var v395 int32
	_ = v395
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v407 int32
	_ = v407
	var v411 int32
	_ = v411
	var v416 int32
	_ = v416
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
	v407 = m.ExcPending
	if v407 != 0 {
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
	v18 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_scan_analyze_next_tuple[0]))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v18+(v14^int32(-1))<<(uint(int32(2))%32))))
	v32 = v24
	goto L2
L4:
	;
	goto L5
L5:
	;
	v26 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_scan_analyze_next_tuple[1]))
	v32 = v26 + v14<<(uint(int32(13))%32) + int32(-8192)
	goto L2
L6:
	;
	v43 = int32(base.Ui32(v33+int32(_a_F_heapam_scan_analyze_next_tuple_0))>>(uint(int32(2))%32)) & int32(_a_F_heapam_scan_analyze_next_tuple_1)
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
	v393 = v14
	goto L11
L11:
	;
	F_UnlockReleaseBuffer(m, v393)
	mBase = m.M
	v395 = m.ExcPending
	if v395 != 0 {
		goto L20
	} else {
		goto L114
	}
L12:
	;
	v63 = v32 + int32(20) + v54<<(uint(int32(2))%32)
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
	switch int32(base.Ui32(v64)>>(uint(int32(15))%32))&int32(3) - int32(1) {
	case 0:
		goto L16
	default:
		goto L14
	case 2:
		goto L15
	}
L13:
	;
	v380 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v393 = v380
	goto L11
L14:
	;
	v375 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	v377 = v375 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+100)) = v377
	if base.Ui32(v377) <= base.Ui32(v43) {
		v54 = v377
		goto L12
	} else {
		goto L113
	}
L15:
	;
	v369 = *(*float64)(unsafe.Add(mBase, uint32(l3)))
	*(*float64)(unsafe.Add(mBase, uint32(l3))) = base.F64_add(v369, float64(1))
	goto L14
L16:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	*(*uint16)(unsafe.Add(mBase, uint32(l4)+56)) = uint16(v54)
	*(*uint16)(unsafe.Add(mBase, uint32(l4)+54)) = uint16(v71)
	v75 = int32(base.Ui32(v71) >> (uint(int32(16)) % 32))
	*(*uint16)(unsafe.Add(mBase, uint32(l4)+52)) = uint16(v75)
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)+56))
	*(*int32)(unsafe.Add(mBase, uint32(l4)+60)) = v78
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
	*(*int32)(unsafe.Add(mBase, uint32(l4)+64)) = v32 + v80&int32(_a_F_heapam_scan_analyze_next_tuple_2)
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
	*(*int32)(unsafe.Add(mBase, uint32(l4)+48)) = int32(base.Ui32(v85) >> (uint(int32(17)) % 32))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v90 = F_HeapTupleSatisfiesVacuum(m, v48, l1, v89)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L20
	} else {
		goto L21
	}
L17:
	;
	v354 = *(*float64)(unsafe.Add(mBase, uint32(l2)))
	*(*float64)(unsafe.Add(mBase, uint32(l2))) = base.F64_add(v354, float64(1))
	v358 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	F_ExecStoreBufferHeapTuple(m, v48, l4, v358)
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L20
	} else {
		goto L112
	}
L18:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(l4)+64))
	v224 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v223)+20)))
	if v224&int32(_a_F_heapam_scan_analyze_next_tuple_3) == int32(_a_F_heapam_scan_analyze_next_tuple_4) {
		goto L67
	} else {
		goto L68
	}
L19:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(l4)+64))
	v95 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v94)+20)))
	v96 = int32(768)
	if v95&v96 != v96 {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	return int32(0)
L21:
	;
	switch v90 {
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
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v94)))
	v102 = v100
	goto L24
L23:
	;
	v102 = int32(2)
	goto L24
L24:
	;
	if base.Ui32(v102) < base.Ui32(int32(3)) {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	if v222 != 0 {
		goto L17
	} else {
		goto L65
	}
L26:
	;
	v222 = int32(0)
	goto L25
L27:
	;
	goto L28
L28:
	;
	v113 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_scan_analyze_next_tuple[2]))
	if v113 == v102 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v222 = int32(1)
	goto L25
L30:
	;
	goto L31
L31:
	;
	v117 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_scan_analyze_next_tuple[3]))
	if v117 <= int32(0) {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	v222 = v214
	goto L25
L33:
	;
	v121 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_scan_analyze_next_tuple[4]))
	if v121 == int32(0) {
		v214 = int32(0)
		goto L32
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	v183 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_scan_analyze_next_tuple[5]))
	v185 = int32(0)
	v187 = v117 - int32(1)
	goto L55
L36:
	;
	v126 = v121
	goto L37
L37:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v126)+20))
	if v131 == int32(4) {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	v214 = int32(0)
	goto L32
L39:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v126)+80))
	if v178 != 0 {
		v126 = v178
		goto L37
	} else {
		goto L54
	}
L40:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v126)))
	if v134 == int32(0) {
		goto L39
	} else {
		goto L41
	}
L41:
	;
	v137 = int32(1)
	if v102 == v134 {
		v214 = v137
		goto L32
	} else {
		goto L42
	}
L42:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v126)+52))
	v141 = v139 - int32(1)
	if v141 < int32(0) {
		goto L39
	} else {
		goto L43
	}
L43:
	;
	v146 = int32(0)
	v148 = v141
	goto L44
L44:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v126)+48))
	v154 = int32(2)
	v155 = base.I32_div_s(v148-v146, v154)
	v156 = v155 + v146
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v152+v156<<(uint(v154)%32))))
	if v160 == v102 {
		v214 = v137
		goto L32
	} else {
		goto L46
	}
L45:
	;
	goto L39
L46:
	;
	v164 = F_TransactionIdPrecedes(m, v160, v102)
	mBase = m.M
	if v164 != 0 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v165 = v156 + int32(1)
	goto L49
L48:
	;
	v165 = v146
	goto L49
L49:
	;
	if v164 != 0 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v168 = v148
	goto L52
L51:
	;
	v168 = v156 - int32(1)
	goto L52
L52:
	;
	if v165 <= v168 {
		v146 = v165
		v148 = v168
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
	v192 = int32(2)
	v193 = base.I32_div_s(v187-v185, v192)
	v194 = v193 + v185
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v183+v194<<(uint(v192)%32))))
	v199 = base.B2i32(v198 == v102)
	if v198 == v102 {
		v214 = v199
		goto L32
	} else {
		goto L57
	}
L56:
	;
	v214 = v199
	goto L32
L57:
	;
	v202 = base.B2i32(base.Ui32(v198) < base.Ui32(v102))
	if base.Ui32(v198) < base.Ui32(v102) {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v203 = v194 + int32(1)
	goto L60
L59:
	;
	v203 = v185
	goto L60
L60:
	;
	if base.Ui32(v198) < base.Ui32(v102) {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v206 = v187
	goto L63
L62:
	;
	v206 = v194 - int32(1)
	goto L63
L63:
	;
	if v203 <= v206 {
		v185 = v203
		v187 = v206
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
	if base.Ui32(v232) < base.Ui32(int32(3)) {
		goto L72
	} else {
		goto L73
	}
L67:
	;
	v229 = F_HeapTupleGetUpdateXid(m, v223)
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L20
	} else {
		goto L70
	}
L68:
	;
	goto L69
L69:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v223)+4))
	v232 = v231
	goto L66
L70:
	;
	v232 = v229
	goto L66
L71:
	;
	if v352 != 0 {
		goto L15
	} else {
		goto L111
	}
L72:
	;
	v352 = int32(0)
	goto L71
L73:
	;
	goto L74
L74:
	;
	v243 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_scan_analyze_next_tuple[2]))
	if v243 == v232 {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v352 = int32(1)
	goto L71
L76:
	;
	goto L77
L77:
	;
	v247 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_scan_analyze_next_tuple[3]))
	if v247 <= int32(0) {
		goto L79
	} else {
		goto L80
	}
L78:
	;
	v352 = v344
	goto L71
L79:
	;
	v251 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_scan_analyze_next_tuple[4]))
	if v251 == int32(0) {
		v344 = int32(0)
		goto L78
	} else {
		goto L82
	}
L80:
	;
	goto L81
L81:
	;
	v313 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_scan_analyze_next_tuple[5]))
	v315 = int32(0)
	v317 = v247 - int32(1)
	goto L101
L82:
	;
	v256 = v251
	goto L83
L83:
	;
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v256)+20))
	if v261 == int32(4) {
		goto L85
	} else {
		goto L86
	}
L84:
	;
	v344 = int32(0)
	goto L78
L85:
	;
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v256)+80))
	if v308 != 0 {
		v256 = v308
		goto L83
	} else {
		goto L100
	}
L86:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v256)))
	if v264 == int32(0) {
		goto L85
	} else {
		goto L87
	}
L87:
	;
	v267 = int32(1)
	if v232 == v264 {
		v344 = v267
		goto L78
	} else {
		goto L88
	}
L88:
	;
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v256)+52))
	v271 = v269 - int32(1)
	if v271 < int32(0) {
		goto L85
	} else {
		goto L89
	}
L89:
	;
	v276 = int32(0)
	v278 = v271
	goto L90
L90:
	;
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v256)+48))
	v284 = int32(2)
	v285 = base.I32_div_s(v278-v276, v284)
	v286 = v285 + v276
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v282+v286<<(uint(v284)%32))))
	if v290 == v232 {
		v344 = v267
		goto L78
	} else {
		goto L92
	}
L91:
	;
	goto L85
L92:
	;
	v294 = F_TransactionIdPrecedes(m, v290, v232)
	mBase = m.M
	if v294 != 0 {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v295 = v286 + int32(1)
	goto L95
L94:
	;
	v295 = v276
	goto L95
L95:
	;
	if v294 != 0 {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v298 = v278
	goto L98
L97:
	;
	v298 = v286 - int32(1)
	goto L98
L98:
	;
	if v295 <= v298 {
		v276 = v295
		v278 = v298
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
	v322 = int32(2)
	v323 = base.I32_div_s(v317-v315, v322)
	v324 = v323 + v315
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v313+v324<<(uint(v322)%32))))
	v329 = base.B2i32(v328 == v232)
	if v328 == v232 {
		v344 = v329
		goto L78
	} else {
		goto L103
	}
L102:
	;
	v344 = v329
	goto L78
L103:
	;
	v332 = base.B2i32(base.Ui32(v328) < base.Ui32(v232))
	if base.Ui32(v328) < base.Ui32(v232) {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	v333 = v324 + int32(1)
	goto L106
L105:
	;
	v333 = v315
	goto L106
L106:
	;
	if base.Ui32(v328) < base.Ui32(v232) {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	v336 = v317
	goto L109
L108:
	;
	v336 = v324 - int32(1)
	goto L109
L109:
	;
	if v333 <= v336 {
		v315 = v333
		v317 = v336
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
	v361 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	v362 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+100)) = v361 + v362
	return v362
L113:
	;
	goto L13
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = int32(0)
	v398 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
	v399 = *(*int32)(unsafe.Add(mBase, uint32(v398)+12))
	m.T0[v399].(func(*base.Module, int32))(m, l4)
	mBase = m.M
	v401 = m.ExcPending
	if v401 != 0 {
		goto L20
	} else {
		goto L115
	}
L115:
	;
	return int32(0)
L116:
	;
	F_errmsg_internal(m, int32(_a_F_heapam_scan_analyze_next_tuple_5), int32(0))
	mBase = m.M
	v411 = m.ExcPending
	if v411 != 0 {
		goto L20
	} else {
		goto L117
	}
L117:
	;
	F_errfinish(m, int32(_a_F_heapam_scan_analyze_next_tuple_6), int32(1147), int32(_a_F_heapam_scan_analyze_next_tuple_7))
	mBase = m.M
	v416 = m.ExcPending
	if v416 != 0 {
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
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v77 int32
	_ = v77
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v125 int32
	_ = v125
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v145 int32
	_ = v145
	var v152 int32
	_ = v152
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v185 int32
	_ = v185
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int64
	_ = v245
	v4 = int32(0)
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)+124))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v25 = v23 & int32(256)
	if v25 == v4 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	F_LockBuffer(m, v28, int32(1))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	if v34 < int32(0) {
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
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+10)))
	if v53&int32(4) != 0 {
		goto L10
	} else {
		goto L11
	}
L7:
	;
	v38 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_scan_sample_next_tuple[0]))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v38+(v34^int32(-1))<<(uint(int32(2))%32))))
	v52 = v44
	goto L6
L8:
	;
	goto L9
L9:
	;
	v46 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_scan_sample_next_tuple[1]))
	v52 = v46 + v34<<(uint(int32(13))%32) + int32(-8192)
	goto L6
L10:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+29)))
	v60 = v57 ^ int32(1)
	goto L12
L11:
	;
	v60 = v4
	goto L12
L12:
	;
	v64 = int32(base.Ui32(v21) >> (uint(int32(16)) % 32))
	v68 = l0 - int32(-64)
	v69 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v52)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v69) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v77 = int32(base.Ui32(v69+int32(_a_F_heapam_scan_sample_next_tuple_0)) >> (uint(int32(2)) % 32))
	goto L15
L14:
	;
	v77 = int32(0)
	goto L15
L15:
	;
	goto L17
L16:
	;
	return base.B2i32(base.Ui32(v108&int32(_a_F_heapam_scan_sample_next_tuple_1)) < base.Ui32(int32(2048)))
L17:
	;
	v101 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_scan_sample_next_tuple[2]))
	if v101 != 0 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	F_ExecStoreBufferHeapTuple(m, v68, l2, v230)
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L4
	} else {
		goto L60
	}
L19:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L4
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v22)+28))
	v105 = m.T0[v104].(func(*base.Module, int32, int32, int32) int32)(m, l1, v21, v77&int32(_a_F_heapam_scan_sample_next_tuple_1))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
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
	if v185 == int32(0) {
		goto L17
	} else {
		goto L59
	}
L25:
	;
	if v25 != 0 {
		goto L17
	} else {
		goto L57
	}
L26:
	;
	v108 = v105 - int32(1)
	if base.Ui32(v108&int32(_a_F_heapam_scan_sample_next_tuple_1)) <= base.Ui32(int32(2047)) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v115 = v52 + int32(20) + v105<<(uint(int32(2))%32)
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v115)))
	if v116&int32(_a_F_heapam_scan_sample_next_tuple_2) != int32(_a_F_heapam_scan_sample_next_tuple_3) {
		goto L17
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	if v25 == int32(0) {
		goto L52
	} else {
		goto L53
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = v52 + v116&int32(_a_F_heapam_scan_sample_next_tuple_4)
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v115)))
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+72)) = uint16(v105)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+70)) = uint16(v21)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+68)) = uint16(v64)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = int32(base.Ui32(v125) >> (uint(int32(17)) % 32))
	v132 = int32(1)
	if v60&v132 != 0 {
		v185 = v132
		goto L31
	} else {
		goto L32
	}
L31:
	;
	if v25 != 0 {
		goto L24
	} else {
		goto L48
	}
L32:
	;
	v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+29)))
	if v135&int32(1) != 0 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v138 = int32(0)
	v139 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v139 == v138 {
		goto L25
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v178 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v179 = F_HeapTupleSatisfiesVisibility(m, v68, v177, v178)
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L4
	} else {
		goto L47
	}
L36:
	;
	v145 = v138
	v152 = v139
	goto L37
L37:
	;
	v163 = int32(1)
	v165 = int32(base.Ui32(v152-v145)>>(uint(v163)%32)) + v145
	v169 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0+int32(108)+v165<<(uint(v163)%32)))))
	v170 = base.B2i32(v105 == v169)
	if v105 == v169 {
		v185 = v170
		goto L31
	} else {
		goto L39
	}
L38:
	;
	v185 = v170
	goto L31
L39:
	;
	v173 = base.B2i32(base.Ui32(v105) < base.Ui32(v169))
	if base.Ui32(v105) < base.Ui32(v169) {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v174 = v145
	goto L42
L41:
	;
	v174 = v165 + int32(1)
	goto L42
L42:
	;
	if base.Ui32(v105) < base.Ui32(v169) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v175 = v165
	goto L45
L44:
	;
	v175 = v152
	goto L45
L45:
	;
	if base.Ui32(v174) < base.Ui32(v175) {
		v145 = v174
		v152 = v175
		goto L37
	} else {
		goto L46
	}
L46:
	;
	goto L38
L47:
	;
	v185 = v179
	goto L31
L48:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v202 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v203 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	F_HeapCheckForSerializableConflictOut(m, v185, v201, v68, v202, v203)
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L4
	} else {
		goto L49
	}
L49:
	;
	if v185 == int32(0) {
		goto L17
	} else {
		goto L50
	}
L50:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	F_LockBuffer(m, v208, int32(0))
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L4
	} else {
		goto L51
	}
L51:
	;
	goto L23
L52:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	F_LockBuffer(m, v214, int32(0))
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L4
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v218)+12))
	m.T0[v219].(func(*base.Module, int32))(m, l2)
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L4
	} else {
		goto L56
	}
L55:
	;
	goto L54
L56:
	;
	goto L16
L57:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v224 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v225 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	F_HeapCheckForSerializableConflictOut(m, int32(0), v223, v68, v224, v225)
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L4
	} else {
		goto L58
	}
L58:
	;
	goto L17
L59:
	;
	goto L23
L60:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v233)+272))
	if v234 == int32(0) {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v233)+268)))
	if v237 != int32(1) {
		goto L16
	} else {
		goto L64
	}
L62:
	;
	v244 = v234
	goto L63
L63:
	;
	v245 = *(*int64)(unsafe.Add(mBase, uint32(v244)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v244)+24)) = v245 + int64(1)
	goto L16
L64:
	;
	F_pgstat_assoc_relation(m, v233)
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L4
	} else {
		goto L65
	}
L65:
	;
	v242 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v242)+272))
	v244 = v243
	goto L63
}
func F_heapam_slot_callbacks(m *base.Module, l0 int32) int32 {
	return int32(_a_F_heapam_slot_callbacks_0)
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
