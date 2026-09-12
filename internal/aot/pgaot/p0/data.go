package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ThrowErrorData(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
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
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
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
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
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
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = F_errstart(m, v5, v6)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	if v7 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v9 = int32(4546812)
	v11 = *(*int32)(unsafe.Add(mBase, _consts[902]))
	*(*int32)(unsafe.Add(mBase, _consts[902])) = v11 + int32(1)
	v15 = int32(4553888)
	v16 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v19 = *(*int32)(unsafe.Add(mBase, _consts[899]))
	v21 = v19 * int32(100)
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_consts[903])))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v24
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v28 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	return
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_consts[901]))) = v28
	goto L8
L7:
	;
	goto L8
L8:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v30 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v31 = F_pstrdup(m, v30)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L1
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v34 != 0 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_consts[905]))) = v31
	goto L11
L13:
	;
	v35 = F_pstrdup(m, v34)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v38 != 0 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_consts[906]))) = v35
	goto L15
L17:
	;
	v39 = F_pstrdup(m, v38)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L1
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v42 != 0 {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_consts[909]))) = v39
	goto L19
L21:
	;
	v43 = F_pstrdup(m, v42)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L1
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	if v46 != 0 {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_consts[907]))) = v43
	goto L23
L25:
	;
	v47 = F_pstrdup(m, v46)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L1
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	if v50 != 0 {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_consts[910]))) = v47
	goto L27
L29:
	;
	v51 = F_pstrdup(m, v50)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L1
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	if v54 != 0 {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_consts[911]))) = v51
	goto L31
L33:
	;
	v55 = F_pstrdup(m, v54)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L1
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	if v58 != 0 {
		goto L37
	} else {
		goto L38
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_consts[912]))) = v55
	goto L35
L37:
	;
	v59 = F_pstrdup(m, v58)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L1
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	if v62 != 0 {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_consts[913]))) = v59
	goto L39
L41:
	;
	v63 = F_pstrdup(m, v62)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L1
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	if v66 != 0 {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_consts[914]))) = v63
	goto L43
L45:
	;
	v67 = F_pstrdup(m, v66)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L1
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	if v70 != 0 {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_consts[915]))) = v67
	goto L47
L49:
	;
	v71 = F_pstrdup(m, v70)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L1
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_consts[917]))) = v74
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_consts[918]))) = v76
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	if v78 != 0 {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_consts[916]))) = v71
	goto L51
L53:
	;
	v79 = F_pstrdup(m, v78)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L1
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v16
	v84 = int32(4546812)
	v86 = *(*int32)(unsafe.Add(mBase, _consts[902]))
	*(*int32)(unsafe.Add(mBase, _consts[902])) = v86 - int32(1)
	v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	F_errfinish(m, v90, v91, v92)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L1
	} else {
		goto L57
	}
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_consts[919]))) = v79
	goto L55
L57:
	;
	goto L5
}
func F_dataBeginPlaceToPage(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
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
	var v81 int32
	_ = v81
	var v88 int64
	_ = v88
	var v111 int32
	_ = v111
	var v128 int32
	_ = v128
	var v129 int64
	_ = v129
	var v132 int64
	_ = v132
	var v136 int64
	_ = v136
	var v140 int32
	_ = v140
	var v153 int32
	_ = v153
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v196 int64
	_ = v196
	var v197 int64
	_ = v197
	var v198 int64
	_ = v198
	var v200 int64
	_ = v200
	var v201 int64
	_ = v201
	var v205 int64
	_ = v205
	var v206 int64
	_ = v206
	var v215 int32
	_ = v215
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
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
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v273 int32
	_ = v273
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v288 int64
	_ = v288
	var v289 int64
	_ = v289
	var v293 int64
	_ = v293
	var v306 int32
	_ = v306
	var v323 int32
	_ = v323
	var v324 int64
	_ = v324
	var v327 int64
	_ = v327
	var v331 int64
	_ = v331
	var v335 int32
	_ = v335
	var v345 int32
	_ = v345
	var v391 int32
	_ = v391
	var v406 int32
	_ = v406
	var v409 int32
	_ = v409
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v420 int32
	_ = v420
	var v422 int32
	_ = v422
	var v425 int64
	_ = v425
	var v426 int64
	_ = v426
	var v430 int64
	_ = v430
	var v431 int64
	_ = v431
	var v436 int64
	_ = v436
	var v438 int64
	_ = v438
	var v439 int64
	_ = v439
	var v442 int64
	_ = v442
	var v448 int32
	_ = v448
	var v451 int32
	_ = v451
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v459 int32
	_ = v459
	var v461 int32
	_ = v461
	var v463 int32
	_ = v463
	var v469 int32
	_ = v469
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v483 int32
	_ = v483
	var v486 int32
	_ = v486
	var v488 int32
	_ = v488
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v502 int32
	_ = v502
	var v512 int32
	_ = v512
	var v514 int32
	_ = v514
	var v520 int32
	_ = v520
	var v526 int32
	_ = v526
	var v545 int32
	_ = v545
	var v554 int32
	_ = v554
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v569 int32
	_ = v569
	var v571 int32
	_ = v571
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
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
	var v581 int32
	_ = v581
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v589 int32
	_ = v589
	var v591 int32
	_ = v591
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v614 int64
	_ = v614
	var v618 int32
	_ = v618
	var v621 int32
	_ = v621
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v630 int32
	_ = v630
	var v631 int32
	_ = v631
	var v636 int32
	_ = v636
	var v639 int32
	_ = v639
	var v641 int32
	_ = v641
	var v644 int32
	_ = v644
	var v646 int32
	_ = v646
	var v652 int32
	_ = v652
	var v655 int32
	_ = v655
	var v657 int32
	_ = v657
	var v659 int32
	_ = v659
	var v660 int32
	_ = v660
	var v662 int32
	_ = v662
	var v670 int32
	_ = v670
	var v672 int32
	_ = v672
	var v673 int32
	_ = v673
	var v675 int32
	_ = v675
	var v678 int32
	_ = v678
	var v679 int32
	_ = v679
	var v682 int32
	_ = v682
	var v684 int32
	_ = v684
	var v685 int32
	_ = v685
	var v687 int32
	_ = v687
	var v691 int32
	_ = v691
	var v693 int32
	_ = v693
	var v699 int32
	_ = v699
	var v700 int32
	_ = v700
	var v703 int32
	_ = v703
	var v705 int32
	_ = v705
	var v707 int32
	_ = v707
	var v713 int32
	_ = v713
	var v740 int32
	_ = v740
	var v745 int32
	_ = v745
	var v755 int32
	_ = v755
	var v758 int32
	_ = v758
	var v760 int32
	_ = v760
	var v770 int32
	_ = v770
	var v773 int32
	_ = v773
	var v774 int32
	_ = v774
	var v777 int32
	_ = v777
	var v778 int32
	_ = v778
	var v781 int32
	_ = v781
	var v785 int32
	_ = v785
	var v786 int32
	_ = v786
	var v788 int32
	_ = v788
	var v789 int32
	_ = v789
	var v794 int32
	_ = v794
	var v795 int32
	_ = v795
	var v799 int32
	_ = v799
	var v800 int32
	_ = v800
	var v804 int32
	_ = v804
	var v805 int32
	_ = v805
	var v807 int32
	_ = v807
	var v810 int32
	_ = v810
	var v813 int32
	_ = v813
	var v814 int32
	_ = v814
	var v817 int32
	_ = v817
	var v819 int32
	_ = v819
	var v820 int32
	_ = v820
	var v825 int32
	_ = v825
	var v829 int32
	_ = v829
	var v832 int32
	_ = v832
	var v834 int32
	_ = v834
	var v835 int32
	_ = v835
	var v836 int32
	_ = v836
	var v838 int32
	_ = v838
	var v843 int32
	_ = v843
	var v848 int32
	_ = v848
	var v849 int32
	_ = v849
	var v851 int32
	_ = v851
	var v852 int32
	_ = v852
	var v853 int32
	_ = v853
	var v856 int32
	_ = v856
	var v857 int32
	_ = v857
	var v859 int32
	_ = v859
	var v860 int32
	_ = v860
	var v861 int32
	_ = v861
	var v862 int32
	_ = v862
	var v863 int32
	_ = v863
	var v866 int32
	_ = v866
	var v867 int32
	_ = v867
	var v869 int32
	_ = v869
	var v870 int32
	_ = v870
	var v877 int32
	_ = v877
	var v879 int32
	_ = v879
	var v882 int32
	_ = v882
	var v883 int32
	_ = v883
	var v885 int32
	_ = v885
	var v889 int32
	_ = v889
	var v895 int32
	_ = v895
	var v898 int32
	_ = v898
	var v904 int32
	_ = v904
	var v910 int32
	_ = v910
	var v916 int32
	_ = v916
	var v917 int32
	_ = v917
	var v919 int32
	_ = v919
	var v922 int32
	_ = v922
	var v924 int32
	_ = v924
	var v925 int32
	_ = v925
	var v939 int32
	_ = v939
	var v952 int32
	_ = v952
	var v954 int32
	_ = v954
	var v956 int32
	_ = v956
	var v958 int32
	_ = v958
	var v962 int32
	_ = v962
	var v983 int32
	_ = v983
	var v984 int32
	_ = v984
	var v986 int32
	_ = v986
	var v988 int32
	_ = v988
	var v992 int32
	_ = v992
	var v1013 int32
	_ = v1013
	var v1015 int32
	_ = v1015
	var v1029 int32
	_ = v1029
	var v1041 int32
	_ = v1041
	var v1043 int32
	_ = v1043
	var v1050 int32
	_ = v1050
	var v1056 int32
	_ = v1056
	var v1071 int32
	_ = v1071
	var v1080 int32
	_ = v1080
	var v1083 int32
	_ = v1083
	var v1084 int32
	_ = v1084
	var v1090 int32
	_ = v1090
	var v1091 int32
	_ = v1091
	var v1092 int32
	_ = v1092
	var v1093 int32
	_ = v1093
	var v1094 int32
	_ = v1094
	var v1095 int32
	_ = v1095
	var v1097 int32
	_ = v1097
	var v1099 int32
	_ = v1099
	var v1124 int32
	_ = v1124
	var v1132 int32
	_ = v1132
	var v1136 int32
	_ = v1136
	var v1137 int32
	_ = v1137
	var v1138 int32
	_ = v1138
	var v1140 int32
	_ = v1140
	var v1141 int32
	_ = v1141
	var v1142 int32
	_ = v1142
	var v1145 int32
	_ = v1145
	var v1150 int32
	_ = v1150
	var v1151 int32
	_ = v1151
	var v1156 int32
	_ = v1156
	var v1161 int64
	_ = v1161
	var v1162 int64
	_ = v1162
	var v1165 int64
	_ = v1165
	var v1180 int32
	_ = v1180
	var v1196 int32
	_ = v1196
	var v1197 int64
	_ = v1197
	var v1200 int64
	_ = v1200
	var v1204 int64
	_ = v1204
	var v1208 int32
	_ = v1208
	var v1221 int32
	_ = v1221
	var v1238 int32
	_ = v1238
	var v1239 int32
	_ = v1239
	var v1240 int32
	_ = v1240
	var v1244 int32
	_ = v1244
	var v1247 int32
	_ = v1247
	var v1248 int32
	_ = v1248
	var v1249 int32
	_ = v1249
	var v1251 int32
	_ = v1251
	var v1255 int32
	_ = v1255
	var v1256 int32
	_ = v1256
	var v1263 int32
	_ = v1263
	var v1269 int32
	_ = v1269
	var v1271 int32
	_ = v1271
	var v1277 int32
	_ = v1277
	var v1278 int32
	_ = v1278
	var v1279 int32
	_ = v1279
	var v1280 int32
	_ = v1280
	var v1281 int32
	_ = v1281
	var v1292 int32
	_ = v1292
	var v1301 int32
	_ = v1301
	var v1307 int32
	_ = v1307
	var v1309 int32
	_ = v1309
	var v1315 int32
	_ = v1315
	var v1316 int32
	_ = v1316
	var v1317 int32
	_ = v1317
	var v1318 int32
	_ = v1318
	var v1319 int32
	_ = v1319
	var v1330 int32
	_ = v1330
	var v1333 int32
	_ = v1333
	var v1334 int32
	_ = v1334
	var v1335 int32
	_ = v1335
	var v1339 int32
	_ = v1339
	var v1345 int32
	_ = v1345
	var v1360 int32
	_ = v1360
	var v1363 int32
	_ = v1363
	var v1364 int32
	_ = v1364
	var v1365 int32
	_ = v1365
	var v1371 int32
	_ = v1371
	var v1372 int32
	_ = v1372
	var v1373 int32
	_ = v1373
	var v1374 int32
	_ = v1374
	var v1383 int32
	_ = v1383
	var v1386 int32
	_ = v1386
	var v1388 int32
	_ = v1388
	var v1398 int32
	_ = v1398
	var v1413 int32
	_ = v1413
	var v1416 int32
	_ = v1416
	var v1419 int32
	_ = v1419
	var v1420 int32
	_ = v1420
	var v1422 int32
	_ = v1422
	var v1423 int32
	_ = v1423
	var v1424 int32
	_ = v1424
	var v1428 int32
	_ = v1428
	var v1429 int32
	_ = v1429
	var v1431 int32
	_ = v1431
	var v1434 int32
	_ = v1434
	var v1435 int32
	_ = v1435
	var v1438 int32
	_ = v1438
	var v1439 int32
	_ = v1439
	var v1441 int32
	_ = v1441
	var v1442 int32
	_ = v1442
	var v1446 int32
	_ = v1446
	var v1447 int32
	_ = v1447
	var v1451 int32
	_ = v1451
	var v1455 int32
	_ = v1455
	var v1456 int32
	_ = v1456
	var v1461 int32
	_ = v1461
	var v1463 int32
	_ = v1463
	var v1464 int32
	_ = v1464
	var v1465 int32
	_ = v1465
	var v1472 int32
	_ = v1472
	var v1482 int32
	_ = v1482
	var v1483 int32
	_ = v1483
	var v1493 int32
	_ = v1493
	var v1496 int32
	_ = v1496
	var v1497 int32
	_ = v1497
	var v1503 int32
	_ = v1503
	var v1505 int32
	_ = v1505
	var v1506 int32
	_ = v1506
	var v1509 int32
	_ = v1509
	var v1510 int32
	_ = v1510
	var v1512 int32
	_ = v1512
	var v1539 int32
	_ = v1539
	var v1541 int32
	_ = v1541
	var v1543 int32
	_ = v1543
	var v1556 int32
	_ = v1556
	var v1559 int32
	_ = v1559
	var v1560 int32
	_ = v1560
	var v1571 int32
	_ = v1571
	var v1574 int32
	_ = v1574
	var v1575 int32
	_ = v1575
	var v1581 int32
	_ = v1581
	var v1583 int32
	_ = v1583
	var v1584 int32
	_ = v1584
	var v1587 int32
	_ = v1587
	var v1588 int32
	_ = v1588
	var v1590 int32
	_ = v1590
	var v1596 int32
	_ = v1596
	var v1600 int32
	_ = v1600
	var v1601 int32
	_ = v1601
	var v1608 int32
	_ = v1608
	var v1614 int32
	_ = v1614
	var v1616 int32
	_ = v1616
	var v1622 int32
	_ = v1622
	var v1623 int32
	_ = v1623
	var v1624 int64
	_ = v1624
	var v1625 int32
	_ = v1625
	var v1626 int32
	_ = v1626
	var v1635 int32
	_ = v1635
	var v1644 int32
	_ = v1644
	var v1650 int32
	_ = v1650
	var v1652 int32
	_ = v1652
	var v1658 int32
	_ = v1658
	var v1659 int32
	_ = v1659
	var v1660 int64
	_ = v1660
	var v1661 int32
	_ = v1661
	var v1662 int32
	_ = v1662
	var v1673 int32
	_ = v1673
	var v1684 int32
	_ = v1684
	var v1699 int32
	_ = v1699
	var v1702 int32
	_ = v1702
	var v1711 int32
	_ = v1711
	var v1712 int32
	_ = v1712
	var v1726 int32
	_ = v1726
	var v1737 int32
	_ = v1737
	var v1759 int32
	_ = v1759
	var v1763 int32
	_ = v1763
	var v1768 int32
	_ = v1768
	var v1795 int32
	_ = v1795
	var v1799 int32
	_ = v1799
	var v1804 int32
	_ = v1804
	v9 = int32(0)
	v24 = m.G0
	v26 = v24 - int32(8320)
	m.G0 = v26
	if l1 < v9 {
		goto L10
	} else {
		goto L11
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1795 = m.ExcPending
	if v1795 != 0 {
		goto L29
	} else {
		goto L314
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1759 = m.ExcPending
	if v1759 != 0 {
		goto L29
	} else {
		goto L311
	}
L3:
	;
	m.G0 = v26 + int32(8320)
	return v1737
L4:
	;
	v1726 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l3)+8)) = v1726 + v1712
	v1737 = v1711
	goto L3
L5:
	;
	v740 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v26)+104)) = uint16(v740)
	*(*int32)(unsafe.Add(mBase, uint32(v26)+100)) = int32(-1)
	v745 = *(*int32)(unsafe.Add(mBase, uint32(v165)+4))
	if v745 != 0 {
		goto L128
	} else {
		goto L129
	}
L6:
	;
	v699 = F_palloc(m, int32(32))
	mBase = m.M
	v700 = m.ExcPending
	if v700 != 0 {
		goto L29
	} else {
		goto L122
	}
L7:
	;
	v565 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+8)))
	v566 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v563)+19)))
	v569 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v563)+28)))
	*(*uint16)(unsafe.Add(mBase, uint32(v26)+uint32(_consts[11]))) = uint16(v569)
	v571 = *(*int32)(unsafe.Add(mBase, uint32(v563)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+uint32(_consts[12]))) = v571
	v573 = F_PageGetTempPage(m, v563)
	mBase = m.M
	v574 = m.ExcPending
	if v574 != 0 {
		goto L29
	} else {
		goto L98
	}
L8:
	;
	v554 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v61)+4)))
	if base.Ui32(v554*int32(-10)+int32(8152)) <= base.Ui32(int32(9)) {
		v563 = v57
		v564 = v554
		goto L7
	} else {
		goto L97
	}
L9:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v72 = v70 - v71
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v76 = v73 + v71*int32(6)
	v77 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v68)+28)))
	v78 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v68)+26)))
	v79 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v68)+24)))
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v67+v68)))
	if v81 == int32(-1) {
		v153 = v72
		goto L16
	} else {
		goto L17
	}
L10:
	;
	v31 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v34 = int32(2)
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v31+(l1^int32(-1))<<(uint(v34)%32))))
	v38 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v37)+16)))
	v39 = v37 + v38
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+6)))
	if v40&v34 != 0 {
		v67 = v38
		v68 = v37
		goto L9
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	v52 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v55 = v52 + l1<<(uint(int32(13))%32)
	v57 = v55 + int32(-8192)
	v60 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v55-int32(8176)))))
	v61 = v57 + v60
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61)+6)))
	if v62&int32(2) == int32(0) {
		goto L8
	} else {
		goto L15
	}
L13:
	;
	v43 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v39)+4)))
	if base.Ui32(v43*int32(-10)+int32(8152)) < base.Ui32(int32(10)) {
		v563 = v37
		v564 = v43
		goto L7
	} else {
		goto L14
	}
L14:
	;
	v1737 = int32(1)
	goto L3
L15:
	;
	v67 = v60
	v68 = v57
	goto L9
L16:
	;
	v165 = F_disassembleLeaf(m, v68)
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L29
	} else {
		goto L30
	}
L17:
	;
	if v72 <= int32(0) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v153 = int32(0)
	goto L16
L19:
	;
	goto L20
L20:
	;
	v88 = int64(65535)
	v111 = int32(0)
	goto L21
L21:
	;
	v128 = v76 + v111*int32(6)
	v129 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v128)+2)))
	v132 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v128))))
	v136 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v128)+4)))
	if base.Ui64(base.I64_extend_i32_u(v77)&v88|(base.I64_extend_i32_u(v78)&v88<<(uint(int64(32))%64)|base.I64_extend_i32_u(v79)&v88<<(uint(int64(48))%64))) < base.Ui64(v129<<(uint(int64(32))%64)|v132<<(uint(int64(48))%64)|v136) {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v153 = v72
	goto L16
L23:
	;
	v153 = v111
	goto L16
L24:
	;
	goto L25
L25:
	;
	v140 = v111 + int32(1)
	if v140 != v72 {
		v111 = v140
		goto L21
	} else {
		goto L26
	}
L26:
	;
	goto L22
L27:
	;
	v225 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v68)+16)))
	v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68+v225)+6)))
	if v227&int32(128) != 0 {
		goto L37
	} else {
		goto L38
	}
L28:
	;
	v215 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v26)+uint32(_consts[11]))) = uint16(v215)
	*(*int32)(unsafe.Add(mBase, uint32(v26)+uint32(_consts[12]))) = v215
	v223 = int32(1)
	goto L27
L29:
	;
	return int32(0)
L30:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v165)+4))
	if v169 == int32(0) {
		goto L28
	} else {
		goto L31
	}
L31:
	;
	if v165 == v169 {
		goto L28
	} else {
		goto L32
	}
L32:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v165)))
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v173)+24))
	if v174 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v173)+20))
	v180 = F_ginPostingListDecode(m, v177, v173+int32(28))
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L29
	} else {
		goto L36
	}
L34:
	;
	v183 = v174
	goto L35
L35:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v173)+28))
	v187 = int32(6)
	v191 = v183 + v186*v187 - v187
	v192 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v191)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v26)+uint32(_consts[11]))) = uint16(v192)
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v191)))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+uint32(_consts[12]))) = v194
	v196 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v76)+4)))
	v197 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v76)+2)))
	v198 = int64(32)
	v200 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v76))))
	v201 = int64(48)
	v205 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v26)+uint32(_consts[11]))))
	v206 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v26)+uint32(_consts[13]))))
	v223 = base.B2i32(base.Ui64(v205|(v206<<(uint(v198)%64)|base.I64_extend_i32_u(v194)<<(uint(v201)%64))) <= base.Ui64(v196|(v197<<(uint(v198)%64)|v200<<(uint(v201)%64))))
	goto L27
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v173)+24)) = v180
	v183 = v180
	goto L35
L37:
	;
	v230 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v68)+14)))
	v231 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v68)+12)))
	v232 = v230 - v231
	v233 = int32(0)
	if v233 < v232 {
		goto L41
	} else {
		goto L42
	}
L38:
	;
	v237 = int32(0)
	goto L39
L39:
	;
	if v223 != 0 {
		goto L45
	} else {
		goto L46
	}
L40:
	;
	v237 = v236
	goto L39
L41:
	;
	v236 = v232
	goto L43
L42:
	;
	v236 = v233
	goto L43
L43:
	;
	goto L40
L44:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v165)+4))
	if v252 == int32(0) {
		goto L6
	} else {
		goto L54
	}
L45:
	;
	v239 = v237 + int32(8152)
	if base.Ui32(v153) < base.Ui32(v239) {
		goto L48
	} else {
		goto L49
	}
L46:
	;
	goto L47
L47:
	;
	v243 = base.I32_div_u_s(v237, int32(384))
	v247 = v243*int32(63) + int32(1323)
	if v153 < v247 {
		goto L51
	} else {
		goto L52
	}
L48:
	;
	v241 = v153
	goto L50
L49:
	;
	v241 = v239
	goto L50
L50:
	;
	v251 = v241
	goto L44
L51:
	;
	v249 = v153
	goto L53
L52:
	;
	v249 = v247
	goto L53
L53:
	;
	v251 = v249
	goto L44
L54:
	;
	if v165 == v252 {
		goto L6
	} else {
		goto L55
	}
L55:
	;
	v265 = v251
	v267 = v76
	v268 = v252
	v273 = v9
	goto L56
L56:
	;
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v268)+4))
	if v279 == v165 {
		v345 = v265
		goto L62
	} else {
		goto L63
	}
L57:
	;
	if v545&int32(1) != 0 {
		goto L5
	} else {
		goto L96
	}
L58:
	;
	goto L57
L59:
	;
	if v526 != v165 {
		v265 = v512
		v267 = v514
		v268 = v526
		v273 = v520
		goto L56
	} else {
		goto L95
	}
L60:
	;
	v406 = *(*int32)(unsafe.Add(mBase, uint32(v268)+24))
	if v406 == int32(0) {
		goto L73
	} else {
		goto L74
	}
L61:
	;
	v512 = v265
	v514 = v267
	v520 = v273
	v526 = v279
	goto L59
L62:
	;
	if v345 != 0 {
		v391 = v345
		goto L60
	} else {
		goto L72
	}
L63:
	;
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v279)+24))
	if v281 == int32(0) {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v279)+20))
	v285 = v284
	goto L66
L65:
	;
	v285 = v281
	goto L66
L66:
	;
	if v265 <= int32(0) {
		goto L61
	} else {
		goto L67
	}
L67:
	;
	v288 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v285)+4)))
	v289 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v285)+2)))
	v293 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v285))))
	v306 = int32(0)
	goto L68
L68:
	;
	v323 = v267 + v306*int32(6)
	v324 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v323)+2)))
	v327 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v323))))
	v331 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v323)+4)))
	if base.Ui64(v288|v289<<(uint(int64(32))%64)|v293<<(uint(int64(48))%64)) <= base.Ui64(v324<<(uint(int64(32))%64)|v327<<(uint(int64(48))%64)|v331) {
		v345 = v306
		goto L62
	} else {
		goto L70
	}
L69:
	;
	v391 = v265
	goto L60
L70:
	;
	v335 = v306 + int32(1)
	if v335 != v265 {
		v306 = v335
		goto L68
	} else {
		goto L71
	}
L71:
	;
	goto L69
L72:
	;
	goto L61
L73:
	;
	v409 = *(*int32)(unsafe.Add(mBase, uint32(v268)+20))
	v412 = F_ginPostingListDecode(m, v409, v268+int32(28))
	mBase = m.M
	v413 = m.ExcPending
	if v413 != 0 {
		goto L29
	} else {
		goto L76
	}
L74:
	;
	v416 = v406
	v417 = v279
	goto L75
L75:
	;
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v268)+28))
	if v165 != v417 {
		goto L77
	} else {
		goto L78
	}
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v268)+24)) = v412
	v415 = *(*int32)(unsafe.Add(mBase, uint32(v268)+4))
	v416 = v412
	v417 = v415
	goto L75
L77:
	;
	v476 = F_ginMergeItemPointers(m, v416, v418, v267, v391, v26+int32(112))
	mBase = m.M
	v477 = m.ExcPending
	if v477 != 0 {
		goto L29
	} else {
		goto L86
	}
L78:
	;
	v420 = int32(6)
	v422 = v416 + v418*v420
	v425 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v422-int32(4)))))
	v426 = int64(32)
	v430 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v422-v420))))
	v431 = int64(48)
	v436 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v422-int32(2)))))
	v438 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v267)+4)))
	v439 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v267)+2)))
	v442 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v267))))
	if base.Ui64(v438|(v439<<(uint(v426)%64)|v442<<(uint(v431)%64))) <= base.Ui64(v425<<(uint(v426)%64)|v430<<(uint(v431)%64)|v436) {
		goto L77
	} else {
		goto L79
	}
L79:
	;
	v448 = *(*int32)(unsafe.Add(mBase, uint32(v268)+20))
	if v448 == int32(0) {
		goto L77
	} else {
		goto L80
	}
L80:
	;
	v451 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v448)+6)))
	if base.Ui32(v451) < base.Ui32(int32(247)) {
		goto L77
	} else {
		goto L81
	}
L81:
	;
	v455 = F_palloc(m, int32(32))
	mBase = m.M
	v456 = m.ExcPending
	if v456 != 0 {
		goto L29
	} else {
		goto L82
	}
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v455)+28)) = v391
	*(*int32)(unsafe.Add(mBase, uint32(v455)+24)) = v267
	v459 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v455)+20)) = v459
	v461 = int32(2)
	*(*uint8)(unsafe.Add(mBase, uint32(v455)+8)) = uint8(v461)
	v463 = *(*int32)(unsafe.Add(mBase, uint32(v165)+4))
	if v463 == v459 {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v165)+4)) = v165
	*(*int32)(unsafe.Add(mBase, uint32(v165))) = v165
	goto L85
L84:
	;
	goto L85
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v455)+4)) = v165
	v469 = *(*int32)(unsafe.Add(mBase, uint32(v165)))
	*(*int32)(unsafe.Add(mBase, uint32(v455))) = v469
	*(*int32)(unsafe.Add(mBase, uint32(v469)+4)) = v455
	*(*int32)(unsafe.Add(mBase, uint32(v165))) = v455
	goto L5
L86:
	;
	v478 = *(*int32)(unsafe.Add(mBase, uint32(v26)+112))
	v479 = *(*int32)(unsafe.Add(mBase, uint32(v268)+28))
	if v478 != v479 {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	if v478 != v391+v479 {
		goto L91
	} else {
		goto L92
	}
L88:
	;
	v495 = v273
	goto L89
L89:
	;
	v496 = v265 - v391
	if v496 == int32(0) {
		v545 = v495
		goto L58
	} else {
		goto L94
	}
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v268)+28)) = v478
	*(*int32)(unsafe.Add(mBase, uint32(v268)+24)) = v476
	*(*int32)(unsafe.Add(mBase, uint32(v268)+20)) = int32(0)
	v495 = int32(1)
	goto L89
L91:
	;
	v488 = int32(3)
	*(*uint8)(unsafe.Add(mBase, uint32(v268)+8)) = uint8(v488)
	goto L90
L92:
	;
	v483 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v268)+8)))
	if v483 != 0 {
		goto L91
	} else {
		goto L93
	}
L93:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v268)+16)) = uint16(v391)
	*(*int32)(unsafe.Add(mBase, uint32(v268)+12)) = v267
	v486 = int32(4)
	*(*uint8)(unsafe.Add(mBase, uint32(v268)+8)) = uint8(v486)
	goto L90
L94:
	;
	v502 = *(*int32)(unsafe.Add(mBase, uint32(v268)+4))
	v512 = v496
	v514 = v267 + v391*int32(6)
	v520 = v495
	v526 = v502
	goto L59
L95:
	;
	v545 = v520
	goto L58
L96:
	;
	v1711 = int32(0)
	v1712 = v251
	goto L4
L97:
	;
	v1737 = int32(1)
	goto L3
L98:
	;
	v575 = F_PageGetTempPage(m, v563)
	mBase = m.M
	v576 = m.ExcPending
	if v576 != 0 {
		goto L29
	} else {
		goto L99
	}
L99:
	;
	v577 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v563)+16)))
	v579 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v563+v577)+6)))
	v580 = int32(8)
	v581 = v566 << (uint(v580) % 32)
	F_PageInit(m, v573, v581, v580)
	mBase = m.M
	v584 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v573)+16)))
	v585 = v573 + v584
	*(*int32)(unsafe.Add(mBase, uint32(v585))) = int32(-1)
	*(*uint16)(unsafe.Add(mBase, uint32(v585)+6)) = uint16(v579)
	goto L100
L100:
	;
	v589 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v563)+16)))
	v591 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v563+v589)+6)))
	F_PageInit(m, v575, v581, int32(8))
	mBase = m.M
	v594 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v575)+16)))
	v595 = v575 + v594
	*(*int32)(unsafe.Add(mBase, uint32(v595))) = int32(-1)
	*(*uint16)(unsafe.Add(mBase, uint32(v595)+6)) = uint16(v591)
	goto L101
L101:
	;
	v602 = v563 + int32(32)
	v603 = int32(10)
	v604 = v565 * v603
	v606 = v604 - v603
	if v606 != 0 {
		goto L103
	} else {
		goto L104
	}
L102:
	;
	v610 = v26 + int32(112)
	v611 = v610 + v606
	v612 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l3)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(v611)+8)) = uint16(v612)
	v614 = *(*int64)(unsafe.Add(mBase, uint32(l3)))
	*(*int64)(unsafe.Add(mBase, uint32(v611))) = v614
	v618 = v610 + v604
	v621 = int32(10)
	v624 = (v564-v565)*v621 + v621
	if v624 != 0 {
		goto L107
	} else {
		goto L108
	}
L103:
	;
	v607 = F__emscripten_memcpy_bulkmem(m, v26+int32(112), v602, v606)
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
	*(*int32)(unsafe.Add(mBase, uint32(v626))) = base.I32_rotr(l4, int32(16))
	v630 = int32(1)
	v631 = v564 + v630
	v636 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+53)))
	if v636 != v630 {
		goto L111
	} else {
		goto L112
	}
L107:
	;
	v625 = F__emscripten_memcpy_bulkmem(m, v618, v606+v602, v624)
	mBase = m.M
	v626 = v625
	goto L109
L108:
	;
	v626 = v618
	goto L109
L109:
	;
	goto L106
L110:
	;
	v657 = v655 & int32(65535)
	v659 = v657 * int32(10)
	if v659 != 0 {
		goto L115
	} else {
		goto L116
	}
L111:
	;
	v655 = int32(base.Ui32(v631) >> (uint(int32(1)) % 32))
	goto L110
L112:
	;
	v639 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v563)+16)))
	v641 = *(*int32)(unsafe.Add(mBase, uint32(v563+v639)))
	if v641 != int32(-1) {
		goto L111
	} else {
		goto L113
	}
L113:
	;
	v644 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v575)+16)))
	v646 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v575+v644)+4)))
	v652 = base.I32_div_u_s(v646*int32(-10)+int32(8152), int32(10))
	v655 = v652
	goto L110
L114:
	;
	v662 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v573)+16)))
	*(*uint16)(unsafe.Add(mBase, uint32(v573+v662)+4)) = uint16(v655)
	v670 = v631 - v657
	v672 = v670 * int32(10)
	if v672 != 0 {
		goto L119
	} else {
		goto L120
	}
L115:
	;
	v660 = F__emscripten_memcpy_bulkmem(m, v573+int32(32), v26+int32(112), v659)
	mBase = m.M
	goto L117
L116:
	;
	goto L117
L117:
	;
	goto L114
L118:
	;
	v675 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v575)+16)))
	*(*uint16)(unsafe.Add(mBase, uint32(v575+v675)+4)) = uint16(v670)
	v678 = int32(32)
	v679 = v659 + v678
	*(*uint16)(unsafe.Add(mBase, uint32(v573)+12)) = uint16(v679)
	v682 = v672 + v678
	*(*uint16)(unsafe.Add(mBase, uint32(v575)+12)) = uint16(v682)
	v684 = v573 + v659
	v685 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v684)+30)))
	*(*uint16)(unsafe.Add(mBase, uint32(v573)+28)) = uint16(v685)
	v687 = *(*int32)(unsafe.Add(mBase, uint32(v684)+26))
	*(*int32)(unsafe.Add(mBase, uint32(v573)+24)) = v687
	v691 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v26)+uint32(_consts[11]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v575)+28)) = uint16(v691)
	v693 = *(*int32)(unsafe.Add(mBase, uint32(v26)+uint32(_consts[12])))
	*(*int32)(unsafe.Add(mBase, uint32(v575)+24)) = v693
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v573
	*(*int32)(unsafe.Add(mBase, uint32(l7))) = v575
	v1737 = int32(2)
	goto L3
L119:
	;
	v673 = F__emscripten_memcpy_bulkmem(m, v575+int32(32), v26+int32(112)+v659, v672)
	mBase = m.M
	goto L121
L120:
	;
	goto L121
L121:
	;
	goto L118
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v699)+28)) = v251
	*(*int32)(unsafe.Add(mBase, uint32(v699)+24)) = v76
	v703 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v699)+20)) = v703
	v705 = int32(2)
	*(*uint8)(unsafe.Add(mBase, uint32(v699)+8)) = uint8(v705)
	v707 = *(*int32)(unsafe.Add(mBase, uint32(v165)+4))
	if v707 == v703 {
		goto L123
	} else {
		goto L124
	}
L123:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v165)+4)) = v165
	*(*int32)(unsafe.Add(mBase, uint32(v165))) = v165
	goto L125
L124:
	;
	goto L125
L125:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v699)+4)) = v165
	v713 = *(*int32)(unsafe.Add(mBase, uint32(v165)))
	*(*int32)(unsafe.Add(mBase, uint32(v699))) = v713
	*(*int32)(unsafe.Add(mBase, uint32(v713)+4)) = v699
	*(*int32)(unsafe.Add(mBase, uint32(v165))) = v699
	goto L5
L126:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v165)+16)) = v1029
	v1043 = *(*int32)(unsafe.Add(mBase, uint32(v165)+4))
	if v1043 == int32(0) {
		goto L181
	} else {
		goto L182
	}
L127:
	;
	v954 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v836)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v26)+104)) = uint16(v954)
	v956 = *(*int32)(unsafe.Add(mBase, uint32(v836)))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+100)) = v956
	v958 = *(*int32)(unsafe.Add(mBase, uint32(v755)+4))
	if v165 != v958 {
		goto L175
	} else {
		goto L176
	}
L128:
	;
	v755 = v745
	v758 = v740
	v760 = int32(0)
	goto L131
L129:
	;
	v939 = v740
	goto L130
L130:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v165)+12)) = v939
	v952 = int32(0)
	v1029 = v952
	v1041 = v952
	goto L126
L131:
	;
	v770 = *(*int32)(unsafe.Add(mBase, uint32(v755)+4))
	if v770 != v165 {
		goto L133
	} else {
		goto L134
	}
L132:
	;
	v925 = int32(1)
	if v924&v925 != 0 {
		v1029 = v922
		v1041 = v925
		goto L126
	} else {
		goto L174
	}
L133:
	;
	v773 = v770
	goto L135
L134:
	;
	v773 = int32(0)
	goto L135
L135:
	;
	v774 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v755)+8)))
	if v774 == int32(1) {
		v919 = v773
		v922 = v758
		v924 = v760
		goto L136
	} else {
		goto L137
	}
L136:
	;
	if v919 != 0 {
		v755 = v919
		v758 = v922
		v760 = v924
		goto L131
	} else {
		goto L173
	}
L137:
	;
	v777 = *(*int32)(unsafe.Add(mBase, uint32(v755)+20))
	if v777 != 0 {
		v835 = v773
		v836 = v777
		goto L138
	} else {
		goto L139
	}
L138:
	;
	v838 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v836)+6)))
	if base.Ui32(int32(118)) < base.Ui32(v838) {
		goto L153
	} else {
		goto L154
	}
L139:
	;
	v778 = *(*int32)(unsafe.Add(mBase, uint32(v755)+28))
	if v778 <= int32(384) {
		goto L141
	} else {
		goto L142
	}
L140:
	;
	v800 = *(*int32)(unsafe.Add(mBase, uint32(v755)+24))
	v804 = F_ginCompressPostingList(m, v800, v799, int32(256), v26+int32(112))
	mBase = m.M
	v805 = m.ExcPending
	if v805 != 0 {
		goto L29
	} else {
		goto L148
	}
L141:
	;
	v781 = *(*int32)(unsafe.Add(mBase, uint32(v755)+24))
	v785 = F_ginCompressPostingList(m, v781, v778, int32(384), v26+int32(112))
	mBase = m.M
	v786 = m.ExcPending
	if v786 != 0 {
		goto L29
	} else {
		goto L144
	}
L142:
	;
	goto L143
L143:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+112)) = int32(0)
	v799 = v778
	goto L140
L144:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v755)+20)) = v785
	v788 = *(*int32)(unsafe.Add(mBase, uint32(v755)+28))
	v789 = *(*int32)(unsafe.Add(mBase, uint32(v26)+112))
	if v788 == v789 {
		v835 = v773
		v836 = v785
		goto L138
	} else {
		goto L145
	}
L145:
	;
	if v785 == int32(0) {
		v799 = v788
		goto L140
	} else {
		goto L146
	}
L146:
	;
	F_pfree(m, v785)
	mBase = m.M
	v794 = m.ExcPending
	if v794 != 0 {
		goto L29
	} else {
		goto L147
	}
L147:
	;
	v795 = *(*int32)(unsafe.Add(mBase, uint32(v755)+28))
	v799 = v795
	goto L140
L148:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v755)+20)) = v804
	v807 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v755)+8)))
	if v807 != int32(2) {
		goto L149
	} else {
		goto L150
	}
L149:
	;
	v810 = int32(3)
	*(*uint8)(unsafe.Add(mBase, uint32(v755)+8)) = uint8(v810)
	goto L151
L150:
	;
	goto L151
L151:
	;
	v813 = F_palloc(m, int32(32))
	mBase = m.M
	v814 = m.ExcPending
	if v814 != 0 {
		goto L29
	} else {
		goto L152
	}
L152:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v813)+20)) = int32(0)
	v817 = int32(2)
	*(*uint8)(unsafe.Add(mBase, uint32(v813)+8)) = uint8(v817)
	v819 = *(*int32)(unsafe.Add(mBase, uint32(v755)+24))
	v820 = *(*int32)(unsafe.Add(mBase, uint32(v26)+112))
	*(*int32)(unsafe.Add(mBase, uint32(v813)+24)) = v819 + v820*int32(6)
	v825 = *(*int32)(unsafe.Add(mBase, uint32(v755)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v813))) = v755
	*(*int32)(unsafe.Add(mBase, uint32(v813)+28)) = v825 - v820
	v829 = *(*int32)(unsafe.Add(mBase, uint32(v755)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v813)+4)) = v829
	*(*int32)(unsafe.Add(mBase, uint32(v755)+4)) = v813
	v832 = *(*int32)(unsafe.Add(mBase, uint32(v813)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v832))) = v813
	v834 = *(*int32)(unsafe.Add(mBase, uint32(v755)+20))
	v835 = v813
	v836 = v834
	goto L138
L153:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v755)+24)) = int64(0)
	v895 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v755)+8)))
	if v895 == int32(1) {
		v919 = v835
		v922 = v758
		v924 = v760
		goto L136
	} else {
		goto L168
	}
L154:
	;
	if v835 == int32(0) {
		goto L153
	} else {
		goto L155
	}
L155:
	;
	v843 = *(*int32)(unsafe.Add(mBase, uint32(v755)+24))
	if v843 == int32(0) {
		goto L156
	} else {
		goto L157
	}
L156:
	;
	v848 = F_ginPostingListDecode(m, v836, v755+int32(28))
	mBase = m.M
	v849 = m.ExcPending
	if v849 != 0 {
		goto L29
	} else {
		goto L159
	}
L157:
	;
	v851 = v843
	goto L158
L158:
	;
	v852 = *(*int32)(unsafe.Add(mBase, uint32(v835)+24))
	if v852 != 0 {
		goto L160
	} else {
		goto L161
	}
L159:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v755)+24)) = v848
	v851 = v848
	goto L158
L160:
	;
	v860 = v852
	v861 = v851
	goto L162
L161:
	;
	v853 = *(*int32)(unsafe.Add(mBase, uint32(v835)+20))
	v856 = F_ginPostingListDecode(m, v853, v835+int32(28))
	mBase = m.M
	v857 = m.ExcPending
	if v857 != 0 {
		goto L29
	} else {
		goto L163
	}
L162:
	;
	v862 = *(*int32)(unsafe.Add(mBase, uint32(v755)+28))
	v863 = *(*int32)(unsafe.Add(mBase, uint32(v835)+28))
	v866 = F_ginMergeItemPointers(m, v861, v862, v860, v863, v26+int32(108))
	mBase = m.M
	v867 = m.ExcPending
	if v867 != 0 {
		goto L29
	} else {
		goto L164
	}
L163:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v835)+24)) = v856
	v859 = *(*int32)(unsafe.Add(mBase, uint32(v755)+24))
	v860 = v856
	v861 = v859
	goto L162
L164:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v835)+24)) = v866
	v869 = *(*int32)(unsafe.Add(mBase, uint32(v26)+108))
	v870 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v835)+20)) = v870
	*(*int32)(unsafe.Add(mBase, uint32(v835)+28)) = v869
	*(*uint16)(unsafe.Add(mBase, uint32(v835)+16)) = uint16(v870)
	*(*int32)(unsafe.Add(mBase, uint32(v835)+12)) = v870
	v877 = int32(3)
	*(*uint8)(unsafe.Add(mBase, uint32(v835)+8)) = uint8(v877)
	v879 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v755)+8)))
	if v879 == int32(2) {
		goto L165
	} else {
		goto L166
	}
L165:
	;
	v882 = *(*int32)(unsafe.Add(mBase, uint32(v755)))
	v883 = *(*int32)(unsafe.Add(mBase, uint32(v755)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v882)+4)) = v883
	v885 = *(*int32)(unsafe.Add(mBase, uint32(v755)))
	*(*int32)(unsafe.Add(mBase, uint32(v883))) = v885
	v919 = v835
	v922 = v758
	v924 = v760
	goto L136
L166:
	;
	goto L167
L167:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v755)+20)) = int32(0)
	v889 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v755)+8)) = uint8(v889)
	*(*int64)(unsafe.Add(mBase, uint32(v755)+24)) = int64(0)
	v919 = v835
	v922 = v758
	v924 = v760
	goto L136
L168:
	;
	v898 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v836)+6)))
	v904 = (v898+int32(1))&int32(131070) + int32(8)
	if base.Ui32(int32(8153)) <= base.Ui32(v904+v758) {
		goto L169
	} else {
		goto L170
	}
L169:
	;
	if v760&int32(1) != 0 {
		goto L127
	} else {
		goto L172
	}
L170:
	;
	v916 = v760
	v917 = v758
	goto L171
L171:
	;
	v919 = v835
	v922 = v917 + v904
	v924 = v916
	goto L136
L172:
	;
	v910 = *(*int32)(unsafe.Add(mBase, uint32(v755)))
	*(*int32)(unsafe.Add(mBase, uint32(v165)+12)) = v758
	*(*int32)(unsafe.Add(mBase, uint32(v165)+8)) = v910
	v916 = int32(1)
	v917 = int32(0)
	goto L171
L173:
	;
	goto L132
L174:
	;
	v939 = v922
	goto L130
L175:
	;
	v962 = v958
	goto L178
L176:
	;
	v992 = v958
	goto L177
L177:
	;
	v1013 = *(*int32)(unsafe.Add(mBase, uint32(v755)))
	*(*int32)(unsafe.Add(mBase, uint32(v1013)+4)) = v992
	v1015 = *(*int32)(unsafe.Add(mBase, uint32(v755)))
	*(*int32)(unsafe.Add(mBase, uint32(v992))) = v1015
	v1029 = v758
	v1041 = int32(1)
	goto L126
L178:
	;
	v983 = *(*int32)(unsafe.Add(mBase, uint32(v962)))
	v984 = *(*int32)(unsafe.Add(mBase, uint32(v962)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v983)+4)) = v984
	v986 = *(*int32)(unsafe.Add(mBase, uint32(v962)))
	*(*int32)(unsafe.Add(mBase, uint32(v984))) = v986
	v988 = *(*int32)(unsafe.Add(mBase, uint32(v755)+4))
	if v988 != v165 {
		v962 = v988
		goto L178
	} else {
		goto L180
	}
L179:
	;
	v992 = v988
	goto L177
L180:
	;
	goto L179
L181:
	;
	v1124 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v26)+104)))
	if v1124 == int32(0) {
		goto L198
	} else {
		goto L199
	}
L182:
	;
	if v165 == v1043 {
		goto L181
	} else {
		goto L183
	}
L183:
	;
	v1050 = int32(0)
	v1056 = v1043
	goto L184
L184:
	;
	v1071 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1056)+8)))
	if v1050&int32(1) == int32(0) {
		goto L187
	} else {
		goto L188
	}
L185:
	;
	goto L181
L186:
	;
	v1099 = *(*int32)(unsafe.Add(mBase, uint32(v1056)+4))
	if v1099 != v165 {
		v1050 = v1097
		v1056 = v1099
		goto L184
	} else {
		goto L196
	}
L187:
	;
	v1097 = base.B2i32(v1071&int32(255) != int32(0))
	goto L186
L188:
	;
	goto L189
L189:
	;
	v1080 = int32(1)
	if v1071&int32(255) != 0 {
		v1097 = v1080
		goto L186
	} else {
		goto L190
	}
L190:
	;
	v1083 = *(*int32)(unsafe.Add(mBase, uint32(v1056)+20))
	v1084 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1083)+6)))
	v1090 = (v1084+int32(1))&int32(131070) + int32(8)
	v1091 = F_palloc(m, v1090)
	mBase = m.M
	v1092 = m.ExcPending
	if v1092 != 0 {
		goto L29
	} else {
		goto L191
	}
L191:
	;
	v1093 = *(*int32)(unsafe.Add(mBase, uint32(v1056)+20))
	if v1090 != 0 {
		goto L193
	} else {
		goto L194
	}
L192:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1056)+20)) = v1095
	v1097 = v1080
	goto L186
L193:
	;
	v1094 = F__emscripten_memcpy_bulkmem(m, v1091, v1093, v1090)
	mBase = m.M
	v1095 = v1094
	goto L195
L194:
	;
	v1095 = v1091
	goto L195
L195:
	;
	goto L192
L196:
	;
	goto L185
L197:
	;
	if v1041 == int32(0) {
		goto L216
	} else {
		goto L217
	}
L198:
	;
	v1221 = v251
	goto L197
L199:
	;
	goto L200
L200:
	;
	if v223 == int32(0) {
		goto L2
	} else {
		goto L201
	}
L201:
	;
	v1132 = v26 + int32(100)
	v1136 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v26)+uint32(_consts[13]))))
	v1137 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v26)+uint32(_consts[12]))))
	v1138 = int32(16)
	v1140 = v1136 | v1137<<(uint(v1138)%32)
	v1141 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1132)+2)))
	v1142 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1132))))
	v1145 = v1141 | v1142<<(uint(v1138)%32)
	if base.Ui32(v1140) < base.Ui32(v1145) {
		v1156 = int32(-1)
		goto L203
	} else {
		goto L204
	}
L202:
	;
	if int32(0) <= v1156 {
		goto L2
	} else {
		goto L207
	}
L203:
	;
	goto L202
L204:
	;
	if base.Ui32(v1145) < base.Ui32(v1140) {
		v1156 = int32(1)
		goto L203
	} else {
		goto L205
	}
L205:
	;
	v1150 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v26)+uint32(_consts[11]))))
	v1151 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1132)+4)))
	if base.Ui32(v1150) < base.Ui32(v1151) {
		v1156 = int32(-1)
		goto L203
	} else {
		goto L206
	}
L206:
	;
	v1156 = base.B2i32(base.Ui32(v1151) < base.Ui32(v1150))
	goto L203
L207:
	;
	if v251 <= int32(0) {
		goto L1
	} else {
		goto L208
	}
L208:
	;
	v1161 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v26)+104)))
	v1162 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v26)+102)))
	v1165 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v26)+100)))
	v1180 = int32(0)
	goto L210
L209:
	;
	if v1180 == int32(0) {
		goto L1
	} else {
		goto L214
	}
L210:
	;
	v1196 = v76 + v1180*int32(6)
	v1197 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v1196)+2)))
	v1200 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v1196))))
	v1204 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v1196)+4)))
	if base.Ui64(v1161|(v1162<<(uint(int64(32))%64)|v1165<<(uint(int64(48))%64))) <= base.Ui64(v1197<<(uint(int64(32))%64)|v1200<<(uint(int64(48))%64)|v1204) {
		goto L209
	} else {
		goto L212
	}
L211:
	;
	v1221 = v251
	goto L197
L212:
	;
	v1208 = v1180 + int32(1)
	if v1208 != v251 {
		v1180 = v1208
		goto L210
	} else {
		goto L213
	}
L213:
	;
	goto L211
L214:
	;
	v1221 = v1180
	goto L197
L215:
	;
	F_errfinish(m, int32(522227), v1699, int32(355522))
	mBase = m.M
	v1702 = m.ExcPending
	if v1702 != 0 {
		goto L29
	} else {
		goto L310
	}
L216:
	;
	v1238 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v1239 = *(*int32)(unsafe.Add(mBase, uint32(v1238)+48))
	v1240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1239)+118)))
	if v1240 != int32(112) {
		goto L219
	} else {
		goto L220
	}
L217:
	;
	goto L218
L218:
	;
	v1333 = *(*int32)(unsafe.Add(mBase, uint32(v165)+8))
	v1334 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+53)))
	if v1334 != 0 {
		v1398 = v1333
		goto L248
	} else {
		goto L249
	}
L219:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v165
	v1255 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v1256 = m.ExcPending
	if v1256 != 0 {
		goto L29
	} else {
		goto L228
	}
L220:
	;
	v1244 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	if v1244 <= int32(0) {
		goto L221
	} else {
		goto L222
	}
L221:
	;
	v1247 = *(*int32)(unsafe.Add(mBase, uint32(v1238)+32))
	if v1247 != 0 {
		goto L219
	} else {
		goto L224
	}
L222:
	;
	goto L223
L223:
	;
	v1249 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+53)))
	if v1249 != 0 {
		goto L219
	} else {
		goto L226
	}
L224:
	;
	v1248 = *(*int32)(unsafe.Add(mBase, uint32(v1238)+40))
	if v1248 != 0 {
		goto L219
	} else {
		goto L225
	}
L225:
	;
	goto L223
L226:
	;
	F_computeLeafRecompressWALData(m, v165)
	mBase = m.M
	v1251 = m.ExcPending
	if v1251 != 0 {
		goto L29
	} else {
		goto L227
	}
L227:
	;
	goto L219
L228:
	;
	if v223 != 0 {
		goto L229
	} else {
		goto L230
	}
L229:
	;
	if v1255 == int32(0) {
		goto L232
	} else {
		goto L233
	}
L230:
	;
	goto L231
L231:
	;
	if v1255 == int32(0) {
		goto L240
	} else {
		goto L241
	}
L232:
	;
	v1711 = int32(1)
	v1712 = v1221
	goto L4
L233:
	;
	goto L234
L234:
	;
	if l1 < int32(0) {
		goto L236
	} else {
		goto L237
	}
L235:
	;
	v1279 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	v1280 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v1281 = *(*int32)(unsafe.Add(mBase, uint32(v165)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+72)) = v1281
	*(*int32)(unsafe.Add(mBase, uint32(v26)+68)) = v1278
	*(*int32)(unsafe.Add(mBase, uint32(v26)+64)) = v1221
	*(*int32)(unsafe.Add(mBase, uint32(v26)+76)) = v1279 - (v1221 + v1280)
	F_errmsg_internal(m, int32(704929), v26-int32(-64))
	mBase = m.M
	v1292 = m.ExcPending
	if v1292 != 0 {
		goto L29
	} else {
		goto L239
	}
L236:
	;
	v1263 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v1269 = *(*int32)(unsafe.Add(mBase, uint32(v1263+(l1^int32(-1))<<(uint(int32(6))%32))+16))
	v1278 = v1269
	goto L235
L237:
	;
	goto L238
L238:
	;
	v1271 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	v1277 = *(*int32)(unsafe.Add(mBase, uint32(v1271+l1<<(uint(int32(6))%32)+int32(-64))+16))
	v1278 = v1277
	goto L235
L239:
	;
	v1684 = int32(1)
	v1699 = int32(611)
	goto L215
L240:
	;
	v1711 = int32(1)
	v1712 = v1221
	goto L4
L241:
	;
	goto L242
L242:
	;
	if l1 < int32(0) {
		goto L244
	} else {
		goto L245
	}
L243:
	;
	v1317 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	v1318 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v1319 = *(*int32)(unsafe.Add(mBase, uint32(v165)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+88)) = v1319
	*(*int32)(unsafe.Add(mBase, uint32(v26)+84)) = v1316
	*(*int32)(unsafe.Add(mBase, uint32(v26)+80)) = v1221
	*(*int32)(unsafe.Add(mBase, uint32(v26)+92)) = v1317 - (v1221 + v1318)
	F_errmsg_internal(m, int32(704874), v26+int32(80))
	mBase = m.M
	v1330 = m.ExcPending
	if v1330 != 0 {
		goto L29
	} else {
		goto L247
	}
L244:
	;
	v1301 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v1307 = *(*int32)(unsafe.Add(mBase, uint32(v1301+(l1^int32(-1))<<(uint(int32(6))%32))+16))
	v1316 = v1307
	goto L243
L245:
	;
	goto L246
L246:
	;
	v1309 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	v1315 = *(*int32)(unsafe.Add(mBase, uint32(v1309+l1<<(uint(int32(6))%32)+int32(-64))+16))
	v1316 = v1315
	goto L243
L247:
	;
	v1684 = int32(1)
	v1699 = int32(615)
	goto L215
L248:
	;
	v1413 = *(*int32)(unsafe.Add(mBase, uint32(v1398)+24))
	if v1413 == int32(0) {
		goto L259
	} else {
		goto L260
	}
L249:
	;
	v1335 = *(*int32)(unsafe.Add(mBase, uint32(v1333)))
	if v1335 == v165 {
		v1398 = v1333
		goto L248
	} else {
		goto L250
	}
L250:
	;
	v1339 = v1335
	v1345 = v1333
	goto L251
L251:
	;
	v1360 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1345)+8)))
	if v1360 != int32(1) {
		goto L253
	} else {
		goto L254
	}
L252:
	;
	v1398 = v1386
	goto L248
L253:
	;
	v1363 = *(*int32)(unsafe.Add(mBase, uint32(v165)+12))
	v1364 = *(*int32)(unsafe.Add(mBase, uint32(v1345)+20))
	v1365 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1364)+6)))
	v1371 = (v1365+int32(1))&int32(131070) + int32(8)
	v1372 = v1363 - v1371
	v1373 = *(*int32)(unsafe.Add(mBase, uint32(v165)+16))
	v1374 = v1371 + v1373
	if v1372-v1374 < int32(0) {
		v1398 = v1345
		goto L248
	} else {
		goto L256
	}
L254:
	;
	v1386 = v1339
	goto L255
L255:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v165)+8)) = v1386
	v1388 = *(*int32)(unsafe.Add(mBase, uint32(v1386)))
	if v1388 != v165 {
		v1339 = v1388
		v1345 = v1386
		goto L251
	} else {
		goto L258
	}
L256:
	;
	if v223&base.B2i32(v1372 < int32(6144)) != 0 {
		v1398 = v1345
		goto L248
	} else {
		goto L257
	}
L257:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v165)+16)) = v1374
	*(*int32)(unsafe.Add(mBase, uint32(v165)+12)) = v1372
	v1383 = *(*int32)(unsafe.Add(mBase, uint32(v1345)))
	v1386 = v1383
	goto L255
L258:
	;
	goto L252
L259:
	;
	v1416 = *(*int32)(unsafe.Add(mBase, uint32(v1398)+20))
	v1419 = F_ginPostingListDecode(m, v1416, v1398+int32(28))
	mBase = m.M
	v1420 = m.ExcPending
	if v1420 != 0 {
		goto L29
	} else {
		goto L262
	}
L260:
	;
	v1422 = v1413
	goto L261
L261:
	;
	v1423 = *(*int32)(unsafe.Add(mBase, uint32(v1398)+28))
	v1424 = int32(6)
	v1428 = v1422 + v1423*v1424 - v1424
	v1429 = *(*int32)(unsafe.Add(mBase, uint32(v1428)))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+112)) = v1429
	v1431 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1428)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v26)+116)) = uint16(v1431)
	v1434 = F_palloc(m, int32(8192))
	mBase = m.M
	v1435 = m.ExcPending
	if v1435 != 0 {
		goto L29
	} else {
		goto L263
	}
L262:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1398)+24)) = v1419
	v1422 = v1419
	goto L261
L263:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v1434
	v1438 = F_palloc(m, int32(8192))
	mBase = m.M
	v1439 = m.ExcPending
	if v1439 != 0 {
		goto L29
	} else {
		goto L264
	}
L264:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l7))) = v1438
	v1441 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	v1442 = int32(131)
	F_PageInit(m, v1441, int32(8192), int32(8))
	mBase = m.M
	v1446 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1441)+16)))
	v1447 = v1441 + v1446
	*(*int32)(unsafe.Add(mBase, uint32(v1447))) = int32(-1)
	*(*uint16)(unsafe.Add(mBase, uint32(v1447)+6)) = uint16(v1442)
	goto L265
L265:
	;
	v1451 = int32(131)
	F_PageInit(m, v1438, int32(8192), int32(8))
	mBase = m.M
	v1455 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1438)+16)))
	v1456 = v1438 + v1455
	*(*int32)(unsafe.Add(mBase, uint32(v1456))) = int32(-1)
	*(*uint16)(unsafe.Add(mBase, uint32(v1456)+6)) = uint16(v1451)
	goto L266
L266:
	;
	v1461 = v1441 + int32(24)
	v1463 = *(*int32)(unsafe.Add(mBase, uint32(v165)+4))
	v1464 = *(*int32)(unsafe.Add(mBase, uint32(v165)+8))
	v1465 = *(*int32)(unsafe.Add(mBase, uint32(v1464)+4))
	if v1463 != v1465 {
		goto L267
	} else {
		goto L268
	}
L267:
	;
	v1472 = v1463
	v1482 = int32(0)
	v1483 = v1441 + int32(32)
	goto L270
L268:
	;
	v1539 = int32(32)
	goto L269
L269:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v1441)+12)) = uint16(v1539)
	v1541 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v26)+116)))
	*(*uint16)(unsafe.Add(mBase, uint32(v1461)+4)) = uint16(v1541)
	v1543 = *(*int32)(unsafe.Add(mBase, uint32(v26)+112))
	*(*int32)(unsafe.Add(mBase, uint32(v1461))) = v1543
	v1556 = v1465
	v1559 = int32(0)
	v1560 = v1438 + int32(32)
	goto L280
L270:
	;
	v1493 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1472)+8)))
	if v1493 != int32(1) {
		goto L272
	} else {
		goto L273
	}
L271:
	;
	v1539 = v1509 + int32(32)
	goto L269
L272:
	;
	v1496 = *(*int32)(unsafe.Add(mBase, uint32(v1472)+20))
	v1497 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1496)+6)))
	v1503 = (v1497+int32(1))&int32(131070) + int32(8)
	if v1503 != 0 {
		goto L276
	} else {
		goto L277
	}
L273:
	;
	v1509 = v1482
	v1510 = v1483
	goto L274
L274:
	;
	v1512 = *(*int32)(unsafe.Add(mBase, uint32(v1472)+4))
	if v1512 != v1465 {
		v1472 = v1512
		v1482 = v1509
		v1483 = v1510
		goto L270
	} else {
		goto L279
	}
L275:
	;
	v1509 = v1503 + v1482
	v1510 = v1506 + v1503
	goto L274
L276:
	;
	v1505 = F__emscripten_memcpy_bulkmem(m, v1483, v1496, v1503)
	mBase = m.M
	v1506 = v1505
	goto L278
L277:
	;
	v1506 = v1483
	goto L278
L278:
	;
	goto L275
L279:
	;
	goto L271
L280:
	;
	v1571 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1556)+8)))
	if v1571 != int32(1) {
		goto L282
	} else {
		goto L283
	}
L281:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v1438)+28)) = uint16(v77)
	*(*uint16)(unsafe.Add(mBase, uint32(v1438)+26)) = uint16(v78)
	*(*uint16)(unsafe.Add(mBase, uint32(v1438)+24)) = uint16(v79)
	v1596 = v1587 + int32(32)
	*(*uint16)(unsafe.Add(mBase, uint32(v1438)+12)) = uint16(v1596)
	v1600 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v1601 = m.ExcPending
	if v1601 != 0 {
		goto L29
	} else {
		goto L290
	}
L282:
	;
	v1574 = *(*int32)(unsafe.Add(mBase, uint32(v1556)+20))
	v1575 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1574)+6)))
	v1581 = (v1575+int32(1))&int32(131070) + int32(8)
	if v1581 != 0 {
		goto L286
	} else {
		goto L287
	}
L283:
	;
	v1587 = v1559
	v1588 = v1560
	goto L284
L284:
	;
	v1590 = *(*int32)(unsafe.Add(mBase, uint32(v1556)+4))
	if v1590 != v165 {
		v1556 = v1590
		v1559 = v1587
		v1560 = v1588
		goto L280
	} else {
		goto L289
	}
L285:
	;
	v1587 = v1581 + v1559
	v1588 = v1584 + v1581
	goto L284
L286:
	;
	v1583 = F__emscripten_memcpy_bulkmem(m, v1560, v1574, v1581)
	mBase = m.M
	v1584 = v1583
	goto L288
L287:
	;
	v1584 = v1560
	goto L288
L288:
	;
	goto L285
L289:
	;
	goto L281
L290:
	;
	if v223 != 0 {
		goto L291
	} else {
		goto L292
	}
L291:
	;
	if v1600 == int32(0) {
		goto L294
	} else {
		goto L295
	}
L292:
	;
	goto L293
L293:
	;
	if v1600 == int32(0) {
		goto L302
	} else {
		goto L303
	}
L294:
	;
	v1711 = int32(2)
	v1712 = v1221
	goto L4
L295:
	;
	goto L296
L296:
	;
	if l1 < int32(0) {
		goto L298
	} else {
		goto L299
	}
L297:
	;
	v1624 = *(*int64)(unsafe.Add(mBase, uint32(v165)+12))
	v1625 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	v1626 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v26))) = v1221
	*(*int32)(unsafe.Add(mBase, uint32(v26)+16)) = v1625 - (v1626 + v1221)
	*(*int32)(unsafe.Add(mBase, uint32(v26)+4)) = v1623
	*(*int64)(unsafe.Add(mBase, uint32(v26)+8)) = v1624
	F_errmsg_internal(m, int32(705038), v26)
	mBase = m.M
	v1635 = m.ExcPending
	if v1635 != 0 {
		goto L29
	} else {
		goto L301
	}
L298:
	;
	v1608 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v1614 = *(*int32)(unsafe.Add(mBase, uint32(v1608+(l1^int32(-1))<<(uint(int32(6))%32))+16))
	v1623 = v1614
	goto L297
L299:
	;
	goto L300
L300:
	;
	v1616 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	v1622 = *(*int32)(unsafe.Add(mBase, uint32(v1616+l1<<(uint(int32(6))%32)+int32(-64))+16))
	v1623 = v1622
	goto L297
L301:
	;
	v1684 = int32(2)
	v1699 = int32(697)
	goto L215
L302:
	;
	v1711 = int32(2)
	v1712 = v1221
	goto L4
L303:
	;
	goto L304
L304:
	;
	if l1 < int32(0) {
		goto L306
	} else {
		goto L307
	}
L305:
	;
	v1660 = *(*int64)(unsafe.Add(mBase, uint32(v165)+12))
	v1661 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	v1662 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+32)) = v1221
	*(*int32)(unsafe.Add(mBase, uint32(v26)+48)) = v1661 - (v1662 + v1221)
	*(*int32)(unsafe.Add(mBase, uint32(v26)+36)) = v1659
	*(*int64)(unsafe.Add(mBase, uint32(v26)+40)) = v1660
	F_errmsg_internal(m, int32(704984), v26+int32(32))
	mBase = m.M
	v1673 = m.ExcPending
	if v1673 != 0 {
		goto L29
	} else {
		goto L309
	}
L306:
	;
	v1644 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v1650 = *(*int32)(unsafe.Add(mBase, uint32(v1644+(l1^int32(-1))<<(uint(int32(6))%32))+16))
	v1659 = v1650
	goto L305
L307:
	;
	goto L308
L308:
	;
	v1652 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	v1658 = *(*int32)(unsafe.Add(mBase, uint32(v1652+l1<<(uint(int32(6))%32)+int32(-64))+16))
	v1659 = v1658
	goto L305
L309:
	;
	v1684 = int32(2)
	v1699 = int32(701)
	goto L215
L310:
	;
	v1711 = v1684
	v1712 = v1221
	goto L4
L311:
	;
	F_errmsg_internal(m, int32(109502), int32(0))
	mBase = m.M
	v1763 = m.ExcPending
	if v1763 != 0 {
		goto L29
	} else {
		goto L312
	}
L312:
	;
	F_errfinish(m, int32(522227), int32(580), int32(355522))
	mBase = m.M
	v1768 = m.ExcPending
	if v1768 != 0 {
		goto L29
	} else {
		goto L313
	}
L313:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L314:
	;
	F_errmsg_internal(m, int32(109553), int32(0))
	mBase = m.M
	v1799 = m.ExcPending
	if v1799 != 0 {
		goto L29
	} else {
		goto L315
	}
L315:
	;
	F_errfinish(m, int32(522227), int32(589), int32(355522))
	mBase = m.M
	v1804 = m.ExcPending
	if v1804 != 0 {
		goto L29
	} else {
		goto L316
	}
L316:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_data_sync_elevel(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	v4 = int32(*(*uint8)(unsafe.Add(mBase, _consts[39])))
	if v4 != 0 {
		v5 = int32(21)
	} else {
		v5 = int32(23)
	}
	return v5
}
func F_show_data_directory_mode(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	v2 = m.G0
	v4 = v2 - int32(16)
	m.G0 = v4
	v7 = *(*int32)(unsafe.Add(mBase, _consts[321]))
	*(*int32)(unsafe.Add(mBase, uint32(v4))) = v7
	v12 = F_pg_snprintf(m, int32(4451121), int32(12), int32(254297), v4)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		m.G0 = v4 + int32(16)
		return int32(4451121)
	}
}
