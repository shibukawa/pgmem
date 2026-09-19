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
	v9 = int32(_a_F_ThrowErrorData_0)
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_ThrowErrorData[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_ThrowErrorData[0])) = v11 + int32(1)
	v15 = int32(_a_F_ThrowErrorData_1)
	v16 = *(*int32)(unsafe.Add(mBase, _c_F_ThrowErrorData[1]))
	v19 = *(*int32)(unsafe.Add(mBase, _c_F_ThrowErrorData[2]))
	v21 = v19 * int32(100)
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_c_F_ThrowErrorData[3])))
	*(*int32)(unsafe.Add(mBase, _c_F_ThrowErrorData[1])) = v24
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
	*(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_c_F_ThrowErrorData[4]))) = v28
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
	*(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_c_F_ThrowErrorData[5]))) = v31
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
	*(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_c_F_ThrowErrorData[6]))) = v35
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
	*(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_c_F_ThrowErrorData[7]))) = v39
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
	*(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_c_F_ThrowErrorData[8]))) = v43
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
	*(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_c_F_ThrowErrorData[9]))) = v47
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
	*(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_c_F_ThrowErrorData[10]))) = v51
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
	*(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_c_F_ThrowErrorData[11]))) = v55
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
	*(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_c_F_ThrowErrorData[12]))) = v59
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
	*(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_c_F_ThrowErrorData[13]))) = v63
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
	*(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_c_F_ThrowErrorData[14]))) = v67
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
	*(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_c_F_ThrowErrorData[16]))) = v74
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_c_F_ThrowErrorData[17]))) = v76
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	if v78 != 0 {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_c_F_ThrowErrorData[15]))) = v71
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
	*(*int32)(unsafe.Add(mBase, _c_F_ThrowErrorData[1])) = v16
	v84 = int32(_a_F_ThrowErrorData_0)
	v86 = *(*int32)(unsafe.Add(mBase, _c_F_ThrowErrorData[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_ThrowErrorData[0])) = v86 - int32(1)
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
	*(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_c_F_ThrowErrorData[18]))) = v79
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
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
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
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v89 int64
	_ = v89
	var v107 int32
	_ = v107
	var v130 int32
	_ = v130
	var v131 int64
	_ = v131
	var v134 int64
	_ = v134
	var v138 int64
	_ = v138
	var v142 int32
	_ = v142
	var v148 int32
	_ = v148
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v200 int64
	_ = v200
	var v201 int64
	_ = v201
	var v202 int64
	_ = v202
	var v204 int64
	_ = v204
	var v205 int64
	_ = v205
	var v209 int64
	_ = v209
	var v210 int64
	_ = v210
	var v219 int32
	_ = v219
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v270 int32
	_ = v270
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v294 int64
	_ = v294
	var v295 int64
	_ = v295
	var v299 int64
	_ = v299
	var v307 int32
	_ = v307
	var v330 int32
	_ = v330
	var v331 int64
	_ = v331
	var v334 int64
	_ = v334
	var v338 int64
	_ = v338
	var v342 int32
	_ = v342
	var v347 int32
	_ = v347
	var v371 int32
	_ = v371
	var v392 int32
	_ = v392
	var v395 int32
	_ = v395
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v406 int32
	_ = v406
	var v408 int32
	_ = v408
	var v411 int64
	_ = v411
	var v412 int64
	_ = v412
	var v416 int64
	_ = v416
	var v417 int64
	_ = v417
	var v422 int64
	_ = v422
	var v424 int64
	_ = v424
	var v425 int64
	_ = v425
	var v428 int64
	_ = v428
	var v434 int32
	_ = v434
	var v437 int32
	_ = v437
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v445 int32
	_ = v445
	var v447 int32
	_ = v447
	var v449 int32
	_ = v449
	var v455 int32
	_ = v455
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v469 int32
	_ = v469
	var v472 int32
	_ = v472
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v479 int32
	_ = v479
	var v482 int32
	_ = v482
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v490 int32
	_ = v490
	var v495 int32
	_ = v495
	var v500 int32
	_ = v500
	var v513 int32
	_ = v513
	var v515 int32
	_ = v515
	var v521 int32
	_ = v521
	var v531 int32
	_ = v531
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v547 int32
	_ = v547
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v565 int32
	_ = v565
	var v567 int32
	_ = v567
	var v570 int32
	_ = v570
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
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v589 int64
	_ = v589
	var v591 int32
	_ = v591
	var v593 int32
	_ = v593
	var v596 int32
	_ = v596
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v607 int32
	_ = v607
	var v609 int32
	_ = v609
	var v612 int32
	_ = v612
	var v614 int32
	_ = v614
	var v620 int32
	_ = v620
	var v623 int32
	_ = v623
	var v625 int32
	_ = v625
	var v627 int32
	_ = v627
	var v633 int32
	_ = v633
	var v636 int32
	_ = v636
	var v638 int32
	_ = v638
	var v645 int32
	_ = v645
	var v648 int32
	_ = v648
	var v649 int32
	_ = v649
	var v652 int32
	_ = v652
	var v656 int32
	_ = v656
	var v657 int32
	_ = v657
	var v660 int32
	_ = v660
	var v662 int32
	_ = v662
	var v664 int32
	_ = v664
	var v694 int32
	_ = v694
	var v695 int32
	_ = v695
	var v698 int32
	_ = v698
	var v700 int32
	_ = v700
	var v702 int32
	_ = v702
	var v708 int32
	_ = v708
	var v736 int32
	_ = v736
	var v741 int32
	_ = v741
	var v746 int32
	_ = v746
	var v747 int32
	_ = v747
	var v755 int32
	_ = v755
	var v767 int32
	_ = v767
	var v770 int32
	_ = v770
	var v771 int32
	_ = v771
	var v774 int32
	_ = v774
	var v775 int32
	_ = v775
	var v780 int32
	_ = v780
	var v784 int32
	_ = v784
	var v785 int32
	_ = v785
	var v787 int32
	_ = v787
	var v788 int32
	_ = v788
	var v793 int32
	_ = v793
	var v794 int32
	_ = v794
	var v795 int32
	_ = v795
	var v797 int32
	_ = v797
	var v801 int32
	_ = v801
	var v802 int32
	_ = v802
	var v804 int32
	_ = v804
	var v807 int32
	_ = v807
	var v810 int32
	_ = v810
	var v811 int32
	_ = v811
	var v814 int32
	_ = v814
	var v816 int32
	_ = v816
	var v817 int32
	_ = v817
	var v822 int32
	_ = v822
	var v826 int32
	_ = v826
	var v829 int32
	_ = v829
	var v831 int32
	_ = v831
	var v833 int32
	_ = v833
	var v834 int32
	_ = v834
	var v835 int32
	_ = v835
	var v837 int32
	_ = v837
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
	var v921 int32
	_ = v921
	var v922 int32
	_ = v922
	var v924 int32
	_ = v924
	var v925 int32
	_ = v925
	var v932 int32
	_ = v932
	var v953 int32
	_ = v953
	var v955 int32
	_ = v955
	var v957 int32
	_ = v957
	var v959 int32
	_ = v959
	var v969 int32
	_ = v969
	var v985 int32
	_ = v985
	var v986 int32
	_ = v986
	var v988 int32
	_ = v988
	var v990 int32
	_ = v990
	var v1000 int32
	_ = v1000
	var v1016 int32
	_ = v1016
	var v1018 int32
	_ = v1018
	var v1025 int32
	_ = v1025
	var v1045 int32
	_ = v1045
	var v1047 int32
	_ = v1047
	var v1048 int32
	_ = v1048
	var v1058 int32
	_ = v1058
	var v1063 int32
	_ = v1063
	var v1079 int32
	_ = v1079
	var v1086 int32
	_ = v1086
	var v1087 int32
	_ = v1087
	var v1088 int32
	_ = v1088
	var v1094 int32
	_ = v1094
	var v1095 int32
	_ = v1095
	var v1096 int32
	_ = v1096
	var v1097 int32
	_ = v1097
	var v1102 int32
	_ = v1102
	var v1103 int32
	_ = v1103
	var v1129 int32
	_ = v1129
	var v1137 int32
	_ = v1137
	var v1141 int32
	_ = v1141
	var v1142 int32
	_ = v1142
	var v1143 int32
	_ = v1143
	var v1145 int32
	_ = v1145
	var v1146 int32
	_ = v1146
	var v1147 int32
	_ = v1147
	var v1150 int32
	_ = v1150
	var v1155 int32
	_ = v1155
	var v1156 int32
	_ = v1156
	var v1161 int32
	_ = v1161
	var v1166 int64
	_ = v1166
	var v1167 int64
	_ = v1167
	var v1170 int64
	_ = v1170
	var v1185 int32
	_ = v1185
	var v1202 int32
	_ = v1202
	var v1203 int64
	_ = v1203
	var v1206 int64
	_ = v1206
	var v1210 int64
	_ = v1210
	var v1214 int32
	_ = v1214
	var v1227 int32
	_ = v1227
	var v1244 int32
	_ = v1244
	var v1245 int32
	_ = v1245
	var v1246 int32
	_ = v1246
	var v1250 int32
	_ = v1250
	var v1253 int32
	_ = v1253
	var v1254 int32
	_ = v1254
	var v1255 int32
	_ = v1255
	var v1257 int32
	_ = v1257
	var v1261 int32
	_ = v1261
	var v1262 int32
	_ = v1262
	var v1269 int32
	_ = v1269
	var v1275 int32
	_ = v1275
	var v1277 int32
	_ = v1277
	var v1283 int32
	_ = v1283
	var v1284 int32
	_ = v1284
	var v1285 int32
	_ = v1285
	var v1286 int32
	_ = v1286
	var v1287 int32
	_ = v1287
	var v1298 int32
	_ = v1298
	var v1307 int32
	_ = v1307
	var v1313 int32
	_ = v1313
	var v1315 int32
	_ = v1315
	var v1321 int32
	_ = v1321
	var v1322 int32
	_ = v1322
	var v1323 int32
	_ = v1323
	var v1324 int32
	_ = v1324
	var v1325 int32
	_ = v1325
	var v1336 int32
	_ = v1336
	var v1339 int32
	_ = v1339
	var v1340 int32
	_ = v1340
	var v1341 int32
	_ = v1341
	var v1346 int32
	_ = v1346
	var v1351 int32
	_ = v1351
	var v1367 int32
	_ = v1367
	var v1370 int32
	_ = v1370
	var v1371 int32
	_ = v1371
	var v1372 int32
	_ = v1372
	var v1378 int32
	_ = v1378
	var v1379 int32
	_ = v1379
	var v1380 int32
	_ = v1380
	var v1381 int32
	_ = v1381
	var v1391 int32
	_ = v1391
	var v1394 int32
	_ = v1394
	var v1396 int32
	_ = v1396
	var v1401 int32
	_ = v1401
	var v1422 int32
	_ = v1422
	var v1425 int32
	_ = v1425
	var v1428 int32
	_ = v1428
	var v1429 int32
	_ = v1429
	var v1431 int32
	_ = v1431
	var v1432 int32
	_ = v1432
	var v1433 int32
	_ = v1433
	var v1437 int32
	_ = v1437
	var v1438 int32
	_ = v1438
	var v1440 int32
	_ = v1440
	var v1443 int32
	_ = v1443
	var v1444 int32
	_ = v1444
	var v1447 int32
	_ = v1447
	var v1448 int32
	_ = v1448
	var v1450 int32
	_ = v1450
	var v1451 int32
	_ = v1451
	var v1455 int32
	_ = v1455
	var v1456 int32
	_ = v1456
	var v1460 int32
	_ = v1460
	var v1464 int32
	_ = v1464
	var v1465 int32
	_ = v1465
	var v1470 int32
	_ = v1470
	var v1471 int32
	_ = v1471
	var v1472 int32
	_ = v1472
	var v1473 int32
	_ = v1473
	var v1480 int32
	_ = v1480
	var v1486 int32
	_ = v1486
	var v1490 int32
	_ = v1490
	var v1502 int32
	_ = v1502
	var v1505 int32
	_ = v1505
	var v1506 int32
	_ = v1506
	var v1512 int32
	_ = v1512
	var v1516 int32
	_ = v1516
	var v1519 int32
	_ = v1519
	var v1520 int32
	_ = v1520
	var v1549 int32
	_ = v1549
	var v1551 int32
	_ = v1551
	var v1553 int32
	_ = v1553
	var v1560 int32
	_ = v1560
	var v1561 int32
	_ = v1561
	var v1562 int32
	_ = v1562
	var v1582 int32
	_ = v1582
	var v1585 int32
	_ = v1585
	var v1586 int32
	_ = v1586
	var v1592 int32
	_ = v1592
	var v1596 int32
	_ = v1596
	var v1597 int32
	_ = v1597
	var v1600 int32
	_ = v1600
	var v1606 int32
	_ = v1606
	var v1610 int32
	_ = v1610
	var v1611 int32
	_ = v1611
	var v1618 int32
	_ = v1618
	var v1624 int32
	_ = v1624
	var v1626 int32
	_ = v1626
	var v1632 int32
	_ = v1632
	var v1633 int32
	_ = v1633
	var v1634 int64
	_ = v1634
	var v1635 int32
	_ = v1635
	var v1636 int32
	_ = v1636
	var v1645 int32
	_ = v1645
	var v1654 int32
	_ = v1654
	var v1660 int32
	_ = v1660
	var v1662 int32
	_ = v1662
	var v1668 int32
	_ = v1668
	var v1669 int32
	_ = v1669
	var v1670 int64
	_ = v1670
	var v1671 int32
	_ = v1671
	var v1672 int32
	_ = v1672
	var v1683 int32
	_ = v1683
	var v1694 int32
	_ = v1694
	var v1710 int32
	_ = v1710
	var v1714 int32
	_ = v1714
	var v1718 int32
	_ = v1718
	var v1724 int32
	_ = v1724
	var v1739 int32
	_ = v1739
	var v1745 int32
	_ = v1745
	var v1773 int32
	_ = v1773
	var v1777 int32
	_ = v1777
	var v1782 int32
	_ = v1782
	var v1810 int32
	_ = v1810
	var v1814 int32
	_ = v1814
	var v1819 int32
	_ = v1819
	v9 = int32(0)
	v25 = m.G0
	v27 = v25 - int32(_a_F_dataBeginPlaceToPage_0)
	m.G0 = v27
	if l1 < v9 {
		goto L12
	} else {
		goto L13
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1810 = m.ExcPending
	if v1810 != 0 {
		goto L29
	} else {
		goto L310
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1773 = m.ExcPending
	if v1773 != 0 {
		goto L29
	} else {
		goto L307
	}
L3:
	;
	m.G0 = v27 + int32(_a_F_dataBeginPlaceToPage_0)
	return v1745
L4:
	;
	v1739 = *(*int32)(unsafe.Add(mBase, uint32(v70)))
	*(*int32)(unsafe.Add(mBase, uint32(v70))) = v1739 + v1724
	v1745 = v1718
	goto L3
L5:
	;
	v736 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v27)+104)) = uint16(v736)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+100)) = int32(-1)
	v741 = *(*int32)(unsafe.Add(mBase, uint32(v168)+4))
	if v741 != 0 {
		goto L128
	} else {
		goto L129
	}
L6:
	;
	v694 = F_palloc(m, int32(32))
	mBase = m.M
	v695 = m.ExcPending
	if v695 != 0 {
		goto L29
	} else {
		goto L122
	}
L7:
	;
	v1718 = int32(0)
	v1724 = v255
	goto L4
L8:
	;
	v543 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+8)))
	v544 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v540)+19)))
	v545 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v540)+28)))
	*(*uint16)(unsafe.Add(mBase, uint32(v27)+uint32(_c_F_dataBeginPlaceToPage[0]))) = uint16(v545)
	v547 = *(*int32)(unsafe.Add(mBase, uint32(v540)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_c_F_dataBeginPlaceToPage[1]))) = v547
	v549 = F_PageGetTempPage(m, v540)
	mBase = m.M
	v550 = m.ExcPending
	if v550 != 0 {
		goto L29
	} else {
		goto L102
	}
L9:
	;
	v531 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v40)+4)))
	if base.Ui32(v531*int32(-10)+int32(_a_F_dataBeginPlaceToPage_1)) < base.Ui32(int32(10)) {
		v540 = v38
		v541 = v531
		goto L8
	} else {
		goto L101
	}
L10:
	;
	v521 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v54-int32(_a_F_dataBeginPlaceToPage_2)))))
	if base.Ui32(int32(9)) < base.Ui32(v521*int32(-10)+int32(_a_F_dataBeginPlaceToPage_1)) {
		goto L98
	} else {
		goto L99
	}
L11:
	;
	v68 = l3 + int32(4)
	v70 = l3 + int32(8)
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v73 = v71 - v72
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v77 = v74 + v72*int32(6)
	v78 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v65)+28)))
	v79 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v65)+26)))
	v80 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v65)+24)))
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v65+v66)))
	if v82 == int32(-1) {
		v148 = v73
		goto L17
	} else {
		goto L18
	}
L12:
	;
	v32 = *(*int32)(unsafe.Add(mBase, _c_F_dataBeginPlaceToPage[2]))
	v35 = int32(2)
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v32+(l1^int32(-1))<<(uint(v35)%32))))
	v39 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v38)+16)))
	v40 = v38 + v39
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+6)))
	if v41&v35 == int32(0) {
		goto L9
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	v47 = *(*int32)(unsafe.Add(mBase, _c_F_dataBeginPlaceToPage[3]))
	v50 = v47 + l1<<(uint(int32(13))%32)
	v53 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v50-int32(_a_F_dataBeginPlaceToPage_3)))))
	v54 = v50 + v53
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54-int32(_a_F_dataBeginPlaceToPage_4)))))
	if v57&int32(2) == int32(0) {
		goto L10
	} else {
		goto L16
	}
L15:
	;
	v65 = v38
	v66 = v39
	goto L11
L16:
	;
	v65 = v50 + int32(-8192)
	v66 = v53
	goto L11
L17:
	;
	v168 = F_disassembleLeaf(m, v65)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L29
	} else {
		goto L30
	}
L18:
	;
	if v73 <= int32(0) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v148 = int32(0)
	goto L17
L20:
	;
	goto L21
L21:
	;
	v89 = int64(65535)
	v107 = int32(0)
	goto L22
L22:
	;
	v130 = v77 + v107*int32(6)
	v131 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v130)+2)))
	v134 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v130))))
	v138 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v130)+4)))
	if base.Ui64(base.I64_extend_i32_u(v78)&v89|(base.I64_extend_i32_u(v79)&v89<<(uint(int64(32))%64)|base.I64_extend_i32_u(v80)&v89<<(uint(int64(48))%64))) < base.Ui64(v131<<(uint(int64(32))%64)|v134<<(uint(int64(48))%64)|v138) {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	v148 = v73
	goto L17
L24:
	;
	v148 = v107
	goto L17
L25:
	;
	goto L26
L26:
	;
	v142 = v107 + int32(1)
	if v142 != v73 {
		v107 = v142
		goto L22
	} else {
		goto L27
	}
L27:
	;
	goto L23
L28:
	;
	v229 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v65)+16)))
	v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65+v229)+6)))
	if v231&int32(128) != 0 {
		goto L38
	} else {
		goto L39
	}
L29:
	;
	return int32(0)
L30:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v168)+4))
	v173 = int32(0)
	if base.B2i32(v172 == v173)|base.B2i32(v172 == v168) == v173 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v168)))
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v179)+24))
	if v180 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L32:
	;
	goto L33
L33:
	;
	v219 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v27)+uint32(_c_F_dataBeginPlaceToPage[0]))) = uint16(v219)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_c_F_dataBeginPlaceToPage[1]))) = v219
	v227 = int32(1)
	goto L28
L34:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v179)+20))
	v186 = F_ginPostingListDecode(m, v183, v179+int32(28))
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L29
	} else {
		goto L37
	}
L35:
	;
	v189 = v180
	goto L36
L36:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v179)+28))
	v191 = int32(6)
	v195 = v189 + v190*v191 - v191
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v195)))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_c_F_dataBeginPlaceToPage[1]))) = v196
	v198 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v195)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v27)+uint32(_c_F_dataBeginPlaceToPage[0]))) = uint16(v198)
	v200 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v77)+4)))
	v201 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v77)+2)))
	v202 = int64(32)
	v204 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v77))))
	v205 = int64(48)
	v209 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v27)+uint32(_c_F_dataBeginPlaceToPage[0]))))
	v210 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v27)+uint32(_c_F_dataBeginPlaceToPage[4]))))
	v227 = base.B2i32(base.Ui64(v209|(v210<<(uint(v202)%64)|base.I64_extend_i32_u(v196)<<(uint(v205)%64))) <= base.Ui64(v200|(v201<<(uint(v202)%64)|v204<<(uint(v205)%64))))
	goto L28
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v179)+24)) = v186
	v189 = v186
	goto L36
L38:
	;
	v234 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v65)+14)))
	v235 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v65)+12)))
	v236 = v234 - v235
	v237 = int32(0)
	if v237 < v236 {
		goto L42
	} else {
		goto L43
	}
L39:
	;
	v241 = int32(0)
	goto L40
L40:
	;
	if v227 != 0 {
		goto L46
	} else {
		goto L47
	}
L41:
	;
	v241 = v240
	goto L40
L42:
	;
	v240 = v236
	goto L44
L43:
	;
	v240 = v237
	goto L44
L44:
	;
	goto L41
L45:
	;
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v168)+4))
	if base.B2i32(v256 == int32(0))|base.B2i32(v256 == v168) != 0 {
		goto L6
	} else {
		goto L55
	}
L46:
	;
	v243 = v241 + int32(_a_F_dataBeginPlaceToPage_1)
	if base.Ui32(v148) < base.Ui32(v243) {
		goto L49
	} else {
		goto L50
	}
L47:
	;
	goto L48
L48:
	;
	v247 = base.I32_div_u_s(v241, int32(384))
	v251 = v247*int32(63) + int32(1323)
	if v148 < v251 {
		goto L52
	} else {
		goto L53
	}
L49:
	;
	v245 = v148
	goto L51
L50:
	;
	v245 = v243
	goto L51
L51:
	;
	v255 = v245
	goto L45
L52:
	;
	v253 = v148
	goto L54
L53:
	;
	v253 = v251
	goto L54
L54:
	;
	v255 = v253
	goto L45
L55:
	;
	v263 = v256
	v265 = v77
	v270 = v255
	v283 = v9
	goto L56
L56:
	;
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v263)+4))
	if v285 == v168 {
		v347 = v270
		goto L60
	} else {
		goto L61
	}
L57:
	;
	if v513 == int32(0) {
		goto L7
	} else {
		goto L97
	}
L58:
	;
	if v168 != v515 {
		v263 = v515
		v265 = v495
		v270 = v500
		v283 = v513
		goto L56
	} else {
		goto L96
	}
L59:
	;
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v263)+24))
	if v392 == int32(0) {
		goto L71
	} else {
		goto L72
	}
L60:
	;
	if v347 != 0 {
		v371 = v347
		goto L59
	} else {
		goto L70
	}
L61:
	;
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v285)+24))
	if v287 == int32(0) {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v285)+20))
	v291 = v290
	goto L64
L63:
	;
	v291 = v287
	goto L64
L64:
	;
	if v270 <= int32(0) {
		v495 = v265
		v500 = v270
		v513 = v283
		v515 = v285
		goto L58
	} else {
		goto L65
	}
L65:
	;
	v294 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v291)+4)))
	v295 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v291)+2)))
	v299 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v291))))
	v307 = int32(0)
	goto L66
L66:
	;
	v330 = v265 + v307*int32(6)
	v331 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v330)+2)))
	v334 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v330))))
	v338 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v330)+4)))
	if base.Ui64(v294|v295<<(uint(int64(32))%64)|v299<<(uint(int64(48))%64)) <= base.Ui64(v331<<(uint(int64(32))%64)|v334<<(uint(int64(48))%64)|v338) {
		v347 = v307
		goto L60
	} else {
		goto L68
	}
L67:
	;
	v371 = v270
	goto L59
L68:
	;
	v342 = v307 + int32(1)
	if v342 != v270 {
		v307 = v342
		goto L66
	} else {
		goto L69
	}
L69:
	;
	goto L67
L70:
	;
	v495 = v265
	v500 = v270
	v513 = v283
	v515 = v285
	goto L58
L71:
	;
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v263)+20))
	v398 = F_ginPostingListDecode(m, v395, v263+int32(28))
	mBase = m.M
	v399 = m.ExcPending
	if v399 != 0 {
		goto L29
	} else {
		goto L74
	}
L72:
	;
	v402 = v392
	v403 = v285
	goto L73
L73:
	;
	v404 = *(*int32)(unsafe.Add(mBase, uint32(v263)+28))
	if v168 != v403 {
		goto L75
	} else {
		goto L76
	}
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v263)+24)) = v398
	v401 = *(*int32)(unsafe.Add(mBase, uint32(v263)+4))
	v402 = v398
	v403 = v401
	goto L73
L75:
	;
	v462 = F_ginMergeItemPointers(m, v402, v404, v265, v371, v27+int32(112))
	mBase = m.M
	v463 = m.ExcPending
	if v463 != 0 {
		goto L29
	} else {
		goto L84
	}
L76:
	;
	v406 = int32(6)
	v408 = v402 + v404*v406
	v411 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v408-int32(4)))))
	v412 = int64(32)
	v416 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v408-v406))))
	v417 = int64(48)
	v422 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v408-int32(2)))))
	v424 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v265)+4)))
	v425 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v265)+2)))
	v428 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v265))))
	if base.Ui64(v424|(v425<<(uint(v412)%64)|v428<<(uint(v417)%64))) <= base.Ui64(v411<<(uint(v412)%64)|v416<<(uint(v417)%64)|v422) {
		goto L75
	} else {
		goto L77
	}
L77:
	;
	v434 = *(*int32)(unsafe.Add(mBase, uint32(v263)+20))
	if v434 == int32(0) {
		goto L75
	} else {
		goto L78
	}
L78:
	;
	v437 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v434)+6)))
	if base.Ui32(v437) < base.Ui32(int32(247)) {
		goto L75
	} else {
		goto L79
	}
L79:
	;
	v441 = F_palloc(m, int32(32))
	mBase = m.M
	v442 = m.ExcPending
	if v442 != 0 {
		goto L29
	} else {
		goto L80
	}
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v441)+28)) = v371
	*(*int32)(unsafe.Add(mBase, uint32(v441)+24)) = v265
	v445 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v441)+20)) = v445
	v447 = int32(2)
	*(*uint8)(unsafe.Add(mBase, uint32(v441)+8)) = uint8(v447)
	v449 = *(*int32)(unsafe.Add(mBase, uint32(v168)+4))
	if v449 == v445 {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v168)+4)) = v168
	*(*int32)(unsafe.Add(mBase, uint32(v168))) = v168
	goto L83
L82:
	;
	goto L83
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v441)+4)) = v168
	v455 = *(*int32)(unsafe.Add(mBase, uint32(v168)))
	*(*int32)(unsafe.Add(mBase, uint32(v441))) = v455
	*(*int32)(unsafe.Add(mBase, uint32(v455)+4)) = v441
	*(*int32)(unsafe.Add(mBase, uint32(v168))) = v441
	goto L5
L84:
	;
	v464 = *(*int32)(unsafe.Add(mBase, uint32(v27)+112))
	v465 = *(*int32)(unsafe.Add(mBase, uint32(v263)+28))
	if v464 != v465 {
		goto L88
	} else {
		goto L89
	}
L85:
	;
	v490 = *(*int32)(unsafe.Add(mBase, uint32(v263)+4))
	v495 = v265 + v371*int32(6)
	v500 = v485
	v513 = v486
	v515 = v490
	goto L58
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v263)+28)) = v464
	*(*int32)(unsafe.Add(mBase, uint32(v263)+24)) = v462
	v479 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v263)+20)) = v479
	v482 = v270 - v371
	if v482 == v479 {
		goto L5
	} else {
		goto L95
	}
L87:
	;
	v475 = int32(3)
	*(*uint8)(unsafe.Add(mBase, uint32(v263)+8)) = uint8(v475)
	goto L86
L88:
	;
	if v464 != v371+v465 {
		goto L87
	} else {
		goto L91
	}
L89:
	;
	goto L90
L90:
	;
	v474 = v270 - v371
	if v474 != 0 {
		v485 = v474
		v486 = v283
		goto L85
	} else {
		goto L93
	}
L91:
	;
	v469 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v263)+8)))
	if v469 != 0 {
		goto L87
	} else {
		goto L92
	}
L92:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v263)+16)) = uint16(v371)
	*(*int32)(unsafe.Add(mBase, uint32(v263)+12)) = v265
	v472 = int32(4)
	*(*uint8)(unsafe.Add(mBase, uint32(v263)+8)) = uint8(v472)
	goto L86
L93:
	;
	if v283 != 0 {
		goto L5
	} else {
		goto L94
	}
L94:
	;
	goto L7
L95:
	;
	v485 = v482
	v486 = int32(1)
	goto L85
L96:
	;
	goto L57
L97:
	;
	goto L5
L98:
	;
	v1745 = int32(1)
	goto L3
L99:
	;
	goto L100
L100:
	;
	v540 = v50 + int32(-8192)
	v541 = v521
	goto L8
L101:
	;
	v1745 = int32(1)
	goto L3
L102:
	;
	v551 = F_PageGetTempPage(m, v540)
	mBase = m.M
	v552 = m.ExcPending
	if v552 != 0 {
		goto L29
	} else {
		goto L103
	}
L103:
	;
	v553 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v540)+16)))
	v555 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v540+v553)+6)))
	v556 = int32(8)
	v557 = v544 << (uint(v556) % 32)
	F_PageInit(m, v549, v557, v556)
	mBase = m.M
	v560 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v549)+16)))
	v561 = v549 + v560
	*(*int32)(unsafe.Add(mBase, uint32(v561))) = int32(-1)
	*(*uint16)(unsafe.Add(mBase, uint32(v561)+6)) = uint16(v555)
	goto L104
L104:
	;
	v565 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v540)+16)))
	v567 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v540+v565)+6)))
	F_PageInit(m, v551, v557, int32(8))
	mBase = m.M
	v570 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v551)+16)))
	v571 = v551 + v570
	*(*int32)(unsafe.Add(mBase, uint32(v571))) = int32(-1)
	*(*uint16)(unsafe.Add(mBase, uint32(v571)+6)) = uint16(v567)
	goto L105
L105:
	;
	v576 = v540 + int32(32)
	v577 = int32(10)
	v578 = v543 * v577
	v580 = v578 - v577
	if v580 != 0 {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	base.MemoryCopy(m, v27+int32(112), v576, v580)
	goto L108
L107:
	;
	goto L108
L108:
	;
	v585 = v27 + int32(112)
	v586 = v585 + v580
	v587 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l3)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(v586)+8)) = uint16(v587)
	v589 = *(*int64)(unsafe.Add(mBase, uint32(l3)))
	*(*int64)(unsafe.Add(mBase, uint32(v586))) = v589
	v591 = v585 + v578
	v593 = int32(10)
	v596 = (v541-v543)*v593 + v593
	if v596 != 0 {
		goto L109
	} else {
		goto L110
	}
L109:
	;
	base.MemoryCopy(m, v591, v580+v576, v596)
	goto L111
L110:
	;
	goto L111
L111:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v591))) = base.I32_rotr(l4, int32(16))
	v602 = int32(1)
	v603 = v541 + v602
	v604 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+53)))
	if v604 != v602 {
		goto L113
	} else {
		goto L114
	}
L112:
	;
	v625 = v623 & int32(_a_F_dataBeginPlaceToPage_5)
	v627 = v625 * int32(10)
	if v627 != 0 {
		goto L116
	} else {
		goto L117
	}
L113:
	;
	v623 = int32(base.Ui32(v603) >> (uint(int32(1)) % 32))
	goto L112
L114:
	;
	v607 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v540)+16)))
	v609 = *(*int32)(unsafe.Add(mBase, uint32(v540+v607)))
	if v609 != int32(-1) {
		goto L113
	} else {
		goto L115
	}
L115:
	;
	v612 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v551)+16)))
	v614 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v551+v612)+4)))
	v620 = base.I32_div_u_s(v614*int32(-10)+int32(_a_F_dataBeginPlaceToPage_1), int32(10))
	v623 = v620
	goto L112
L116:
	;
	base.MemoryCopy(m, v549+int32(32), v27+int32(112), v627)
	goto L118
L117:
	;
	goto L118
L118:
	;
	v633 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v549)+16)))
	*(*uint16)(unsafe.Add(mBase, uint32(v549+v633)+4)) = uint16(v623)
	v636 = v603 - v625
	v638 = v636 * int32(10)
	if v638 != 0 {
		goto L119
	} else {
		goto L120
	}
L119:
	;
	base.MemoryCopy(m, v551+int32(32), v27+int32(112)+v627, v638)
	goto L121
L120:
	;
	goto L121
L121:
	;
	v645 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v551)+16)))
	*(*uint16)(unsafe.Add(mBase, uint32(v551+v645)+4)) = uint16(v636)
	v648 = int32(32)
	v649 = v627 + v648
	*(*uint16)(unsafe.Add(mBase, uint32(v549)+12)) = uint16(v649)
	v652 = v638 + v648
	*(*uint16)(unsafe.Add(mBase, uint32(v551)+12)) = uint16(v652)
	v656 = v549 + int32(24) + v627
	v657 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v656)+6)))
	*(*uint16)(unsafe.Add(mBase, uint32(v549)+28)) = uint16(v657)
	v660 = *(*int32)(unsafe.Add(mBase, uint32(v656)+2))
	*(*int32)(unsafe.Add(mBase, uint32(v549)+24)) = v660
	v662 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v27)+uint32(_c_F_dataBeginPlaceToPage[0]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v551)+28)) = uint16(v662)
	v664 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_c_F_dataBeginPlaceToPage[1])))
	*(*int32)(unsafe.Add(mBase, uint32(v551)+24)) = v664
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v549
	*(*int32)(unsafe.Add(mBase, uint32(l7))) = v551
	v1745 = int32(2)
	goto L3
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v694)+28)) = v255
	*(*int32)(unsafe.Add(mBase, uint32(v694)+24)) = v77
	v698 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v694)+20)) = v698
	v700 = int32(2)
	*(*uint8)(unsafe.Add(mBase, uint32(v694)+8)) = uint8(v700)
	v702 = *(*int32)(unsafe.Add(mBase, uint32(v168)+4))
	if v702 == v698 {
		goto L123
	} else {
		goto L124
	}
L123:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v168)+4)) = v168
	*(*int32)(unsafe.Add(mBase, uint32(v168))) = v168
	goto L125
L124:
	;
	goto L125
L125:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v694)+4)) = v168
	v708 = *(*int32)(unsafe.Add(mBase, uint32(v168)))
	*(*int32)(unsafe.Add(mBase, uint32(v694))) = v708
	*(*int32)(unsafe.Add(mBase, uint32(v708)+4)) = v694
	*(*int32)(unsafe.Add(mBase, uint32(v168))) = v694
	goto L5
L126:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v168)+16)) = v1025
	v1047 = *(*int32)(unsafe.Add(mBase, uint32(v168)+4))
	v1048 = int32(0)
	if base.B2i32(v1047 == v1048)|base.B2i32(v1047 == v168) == v1048 {
		goto L181
	} else {
		goto L182
	}
L127:
	;
	v955 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v834)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v27)+104)) = uint16(v955)
	v957 = *(*int32)(unsafe.Add(mBase, uint32(v834)))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+100)) = v957
	v959 = *(*int32)(unsafe.Add(mBase, uint32(v746)+4))
	if v168 != v959 {
		goto L175
	} else {
		goto L176
	}
L128:
	;
	v746 = v741
	v747 = v736
	v755 = int32(0)
	goto L131
L129:
	;
	v932 = v736
	goto L130
L130:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v168)+12)) = v932
	v953 = int32(0)
	v1025 = v953
	v1045 = v953
	goto L126
L131:
	;
	v767 = *(*int32)(unsafe.Add(mBase, uint32(v746)+4))
	if v767 != v168 {
		goto L133
	} else {
		goto L134
	}
L132:
	;
	v925 = int32(1)
	if v924&v925 != 0 {
		v1025 = v921
		v1045 = v925
		goto L126
	} else {
		goto L174
	}
L133:
	;
	v770 = v767
	goto L135
L134:
	;
	v770 = int32(0)
	goto L135
L135:
	;
	v771 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v746)+8)))
	if v771 == int32(1) {
		v921 = v747
		v922 = v770
		v924 = v755
		goto L136
	} else {
		goto L137
	}
L136:
	;
	if v922 != 0 {
		v746 = v922
		v747 = v921
		v755 = v924
		goto L131
	} else {
		goto L173
	}
L137:
	;
	v774 = *(*int32)(unsafe.Add(mBase, uint32(v746)+20))
	if v774 != 0 {
		v833 = v770
		v834 = v774
		goto L138
	} else {
		goto L139
	}
L138:
	;
	v835 = int32(0)
	v837 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v834)+6)))
	if base.B2i32(v833 == v835)|base.B2i32(base.Ui32(int32(118)) < base.Ui32(v837)) == v835 {
		goto L153
	} else {
		goto L154
	}
L139:
	;
	v775 = *(*int32)(unsafe.Add(mBase, uint32(v746)+28))
	if int32(385) <= v775 {
		goto L141
	} else {
		goto L142
	}
L140:
	;
	v797 = *(*int32)(unsafe.Add(mBase, uint32(v746)+24))
	v801 = F_ginCompressPostingList(m, v797, v795, int32(256), v27+int32(112))
	mBase = m.M
	v802 = m.ExcPending
	if v802 != 0 {
		goto L29
	} else {
		goto L148
	}
L141:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+112)) = int32(0)
	v795 = v775
	goto L140
L142:
	;
	goto L143
L143:
	;
	v780 = *(*int32)(unsafe.Add(mBase, uint32(v746)+24))
	v784 = F_ginCompressPostingList(m, v780, v775, int32(384), v27+int32(112))
	mBase = m.M
	v785 = m.ExcPending
	if v785 != 0 {
		goto L29
	} else {
		goto L144
	}
L144:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v746)+20)) = v784
	v787 = *(*int32)(unsafe.Add(mBase, uint32(v746)+28))
	v788 = *(*int32)(unsafe.Add(mBase, uint32(v27)+112))
	if v787 == v788 {
		v833 = v770
		v834 = v784
		goto L138
	} else {
		goto L145
	}
L145:
	;
	if v784 == int32(0) {
		v795 = v787
		goto L140
	} else {
		goto L146
	}
L146:
	;
	F_pfree(m, v784)
	mBase = m.M
	v793 = m.ExcPending
	if v793 != 0 {
		goto L29
	} else {
		goto L147
	}
L147:
	;
	v794 = *(*int32)(unsafe.Add(mBase, uint32(v746)+28))
	v795 = v794
	goto L140
L148:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v746)+20)) = v801
	v804 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v746)+8)))
	if v804 != int32(2) {
		goto L149
	} else {
		goto L150
	}
L149:
	;
	v807 = int32(3)
	*(*uint8)(unsafe.Add(mBase, uint32(v746)+8)) = uint8(v807)
	goto L151
L150:
	;
	goto L151
L151:
	;
	v810 = F_palloc(m, int32(32))
	mBase = m.M
	v811 = m.ExcPending
	if v811 != 0 {
		goto L29
	} else {
		goto L152
	}
L152:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v810)+20)) = int32(0)
	v814 = int32(2)
	*(*uint8)(unsafe.Add(mBase, uint32(v810)+8)) = uint8(v814)
	v816 = *(*int32)(unsafe.Add(mBase, uint32(v746)+24))
	v817 = *(*int32)(unsafe.Add(mBase, uint32(v27)+112))
	*(*int32)(unsafe.Add(mBase, uint32(v810)+24)) = v816 + v817*int32(6)
	v822 = *(*int32)(unsafe.Add(mBase, uint32(v746)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v810))) = v746
	*(*int32)(unsafe.Add(mBase, uint32(v810)+28)) = v822 - v817
	v826 = *(*int32)(unsafe.Add(mBase, uint32(v746)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v810)+4)) = v826
	*(*int32)(unsafe.Add(mBase, uint32(v746)+4)) = v810
	v829 = *(*int32)(unsafe.Add(mBase, uint32(v810)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v829))) = v810
	v831 = *(*int32)(unsafe.Add(mBase, uint32(v746)+20))
	v833 = v810
	v834 = v831
	goto L138
L153:
	;
	v843 = *(*int32)(unsafe.Add(mBase, uint32(v746)+24))
	if v843 == int32(0) {
		goto L156
	} else {
		goto L157
	}
L154:
	;
	goto L155
L155:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v746)+24)) = int64(0)
	v895 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v746)+8)))
	if v895 == int32(1) {
		v921 = v747
		v922 = v833
		v924 = v755
		goto L136
	} else {
		goto L168
	}
L156:
	;
	v848 = F_ginPostingListDecode(m, v834, v746+int32(28))
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
	v852 = *(*int32)(unsafe.Add(mBase, uint32(v833)+24))
	if v852 != 0 {
		goto L160
	} else {
		goto L161
	}
L159:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v746)+24)) = v848
	v851 = v848
	goto L158
L160:
	;
	v860 = v852
	v861 = v851
	goto L162
L161:
	;
	v853 = *(*int32)(unsafe.Add(mBase, uint32(v833)+20))
	v856 = F_ginPostingListDecode(m, v853, v833+int32(28))
	mBase = m.M
	v857 = m.ExcPending
	if v857 != 0 {
		goto L29
	} else {
		goto L163
	}
L162:
	;
	v862 = *(*int32)(unsafe.Add(mBase, uint32(v746)+28))
	v863 = *(*int32)(unsafe.Add(mBase, uint32(v833)+28))
	v866 = F_ginMergeItemPointers(m, v861, v862, v860, v863, v27+int32(108))
	mBase = m.M
	v867 = m.ExcPending
	if v867 != 0 {
		goto L29
	} else {
		goto L164
	}
L163:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v833)+24)) = v856
	v859 = *(*int32)(unsafe.Add(mBase, uint32(v746)+24))
	v860 = v856
	v861 = v859
	goto L162
L164:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v833)+24)) = v866
	v869 = *(*int32)(unsafe.Add(mBase, uint32(v27)+108))
	v870 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v833)+20)) = v870
	*(*int32)(unsafe.Add(mBase, uint32(v833)+28)) = v869
	*(*uint16)(unsafe.Add(mBase, uint32(v833)+16)) = uint16(v870)
	*(*int32)(unsafe.Add(mBase, uint32(v833)+12)) = v870
	v877 = int32(3)
	*(*uint8)(unsafe.Add(mBase, uint32(v833)+8)) = uint8(v877)
	v879 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v746)+8)))
	if v879 == int32(2) {
		goto L165
	} else {
		goto L166
	}
L165:
	;
	v882 = *(*int32)(unsafe.Add(mBase, uint32(v746)))
	v883 = *(*int32)(unsafe.Add(mBase, uint32(v746)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v882)+4)) = v883
	v885 = *(*int32)(unsafe.Add(mBase, uint32(v746)))
	*(*int32)(unsafe.Add(mBase, uint32(v883))) = v885
	v921 = v747
	v922 = v833
	v924 = v755
	goto L136
L166:
	;
	goto L167
L167:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v746)+20)) = int32(0)
	v889 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v746)+8)) = uint8(v889)
	*(*int64)(unsafe.Add(mBase, uint32(v746)+24)) = int64(0)
	v921 = v747
	v922 = v833
	v924 = v755
	goto L136
L168:
	;
	v898 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v834)+6)))
	v904 = (v898+int32(1))&int32(_a_F_dataBeginPlaceToPage_6) + int32(8)
	if base.Ui32(int32(_a_F_dataBeginPlaceToPage_7)) <= base.Ui32(v904+v747) {
		goto L169
	} else {
		goto L170
	}
L169:
	;
	if v755&int32(1) != 0 {
		goto L127
	} else {
		goto L172
	}
L170:
	;
	v916 = v755
	v917 = v747
	goto L171
L171:
	;
	v921 = v917 + v904
	v922 = v833
	v924 = v916
	goto L136
L172:
	;
	v910 = *(*int32)(unsafe.Add(mBase, uint32(v746)))
	*(*int32)(unsafe.Add(mBase, uint32(v168)+12)) = v747
	*(*int32)(unsafe.Add(mBase, uint32(v168)+8)) = v910
	v916 = int32(1)
	v917 = int32(0)
	goto L171
L173:
	;
	goto L132
L174:
	;
	v932 = v921
	goto L130
L175:
	;
	v969 = v959
	goto L178
L176:
	;
	v1000 = v959
	goto L177
L177:
	;
	v1016 = *(*int32)(unsafe.Add(mBase, uint32(v746)))
	*(*int32)(unsafe.Add(mBase, uint32(v1016)+4)) = v1000
	v1018 = *(*int32)(unsafe.Add(mBase, uint32(v746)))
	*(*int32)(unsafe.Add(mBase, uint32(v1000))) = v1018
	v1025 = v747
	v1045 = int32(1)
	goto L126
L178:
	;
	v985 = *(*int32)(unsafe.Add(mBase, uint32(v969)))
	v986 = *(*int32)(unsafe.Add(mBase, uint32(v969)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v985)+4)) = v986
	v988 = *(*int32)(unsafe.Add(mBase, uint32(v969)))
	*(*int32)(unsafe.Add(mBase, uint32(v986))) = v988
	v990 = *(*int32)(unsafe.Add(mBase, uint32(v746)+4))
	if v990 != v168 {
		v969 = v990
		goto L178
	} else {
		goto L180
	}
L179:
	;
	v1000 = v990
	goto L177
L180:
	;
	goto L179
L181:
	;
	v1058 = v1047
	v1063 = int32(0)
	goto L184
L182:
	;
	goto L183
L183:
	;
	v1129 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v27)+104)))
	if v1129 == int32(0) {
		goto L197
	} else {
		goto L198
	}
L184:
	;
	v1079 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1058)+8)))
	if v1063&int32(1) == int32(0) {
		goto L187
	} else {
		goto L188
	}
L185:
	;
	goto L183
L186:
	;
	v1103 = *(*int32)(unsafe.Add(mBase, uint32(v1058)+4))
	if v1103 != v168 {
		v1058 = v1103
		v1063 = v1102
		goto L184
	} else {
		goto L195
	}
L187:
	;
	v1102 = base.B2i32(v1079 != int32(0))
	goto L186
L188:
	;
	goto L189
L189:
	;
	v1086 = int32(1)
	if v1079 != 0 {
		v1102 = v1086
		goto L186
	} else {
		goto L190
	}
L190:
	;
	v1087 = *(*int32)(unsafe.Add(mBase, uint32(v1058)+20))
	v1088 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1087)+6)))
	v1094 = (v1088+int32(1))&int32(_a_F_dataBeginPlaceToPage_6) + int32(8)
	v1095 = F_palloc(m, v1094)
	mBase = m.M
	v1096 = m.ExcPending
	if v1096 != 0 {
		goto L29
	} else {
		goto L191
	}
L191:
	;
	if v1094 != 0 {
		goto L192
	} else {
		goto L193
	}
L192:
	;
	v1097 = *(*int32)(unsafe.Add(mBase, uint32(v1058)+20))
	base.MemoryCopy(m, v1095, v1097, v1094)
	goto L194
L193:
	;
	goto L194
L194:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1058)+20)) = v1095
	v1102 = v1086
	goto L186
L195:
	;
	goto L185
L196:
	;
	if v1045 == int32(0) {
		goto L215
	} else {
		goto L216
	}
L197:
	;
	v1227 = v255
	goto L196
L198:
	;
	goto L199
L199:
	;
	if v227 == int32(0) {
		goto L2
	} else {
		goto L200
	}
L200:
	;
	v1137 = v27 + int32(100)
	v1141 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v27)+uint32(_c_F_dataBeginPlaceToPage[4]))))
	v1142 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v27)+uint32(_c_F_dataBeginPlaceToPage[1]))))
	v1143 = int32(16)
	v1145 = v1141 | v1142<<(uint(v1143)%32)
	v1146 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1137)+2)))
	v1147 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1137))))
	v1150 = v1146 | v1147<<(uint(v1143)%32)
	if base.Ui32(v1145) < base.Ui32(v1150) {
		v1161 = int32(-1)
		goto L202
	} else {
		goto L203
	}
L201:
	;
	if int32(0) <= v1161 {
		goto L2
	} else {
		goto L206
	}
L202:
	;
	goto L201
L203:
	;
	if base.Ui32(v1150) < base.Ui32(v1145) {
		v1161 = int32(1)
		goto L202
	} else {
		goto L204
	}
L204:
	;
	v1155 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v27)+uint32(_c_F_dataBeginPlaceToPage[0]))))
	v1156 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1137)+4)))
	if base.Ui32(v1155) < base.Ui32(v1156) {
		v1161 = int32(-1)
		goto L202
	} else {
		goto L205
	}
L205:
	;
	v1161 = base.B2i32(base.Ui32(v1156) < base.Ui32(v1155))
	goto L202
L206:
	;
	if v255 <= int32(0) {
		goto L1
	} else {
		goto L207
	}
L207:
	;
	v1166 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v27)+104)))
	v1167 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v27)+102)))
	v1170 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v27)+100)))
	v1185 = int32(0)
	goto L209
L208:
	;
	if v1185 == int32(0) {
		goto L1
	} else {
		goto L213
	}
L209:
	;
	v1202 = v77 + v1185*int32(6)
	v1203 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v1202)+2)))
	v1206 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v1202))))
	v1210 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v1202)+4)))
	if base.Ui64(v1166|(v1167<<(uint(int64(32))%64)|v1170<<(uint(int64(48))%64))) <= base.Ui64(v1203<<(uint(int64(32))%64)|v1206<<(uint(int64(48))%64)|v1210) {
		goto L208
	} else {
		goto L211
	}
L210:
	;
	v1227 = v255
	goto L196
L211:
	;
	v1214 = v1185 + int32(1)
	if v1214 != v255 {
		v1185 = v1214
		goto L209
	} else {
		goto L212
	}
L212:
	;
	goto L210
L213:
	;
	v1227 = v1185
	goto L196
L214:
	;
	F_errfinish(m, int32(_a_F_dataBeginPlaceToPage_8), v1694, int32(_a_F_dataBeginPlaceToPage_9))
	mBase = m.M
	v1714 = m.ExcPending
	if v1714 != 0 {
		goto L29
	} else {
		goto L306
	}
L215:
	;
	v1244 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v1245 = *(*int32)(unsafe.Add(mBase, uint32(v1244)+48))
	v1246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1245)+118)))
	if v1246 != int32(112) {
		goto L218
	} else {
		goto L219
	}
L216:
	;
	goto L217
L217:
	;
	v1339 = *(*int32)(unsafe.Add(mBase, uint32(v168)+8))
	v1340 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+53)))
	if v1340 != 0 {
		v1401 = v1339
		goto L247
	} else {
		goto L248
	}
L218:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v168
	v1261 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v1262 = m.ExcPending
	if v1262 != 0 {
		goto L29
	} else {
		goto L227
	}
L219:
	;
	v1250 = *(*int32)(unsafe.Add(mBase, _c_F_dataBeginPlaceToPage[5]))
	if v1250 <= int32(0) {
		goto L220
	} else {
		goto L221
	}
L220:
	;
	v1253 = *(*int32)(unsafe.Add(mBase, uint32(v1244)+32))
	if v1253 != 0 {
		goto L218
	} else {
		goto L223
	}
L221:
	;
	goto L222
L222:
	;
	v1255 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+53)))
	if v1255 != 0 {
		goto L218
	} else {
		goto L225
	}
L223:
	;
	v1254 = *(*int32)(unsafe.Add(mBase, uint32(v1244)+40))
	if v1254 != 0 {
		goto L218
	} else {
		goto L224
	}
L224:
	;
	goto L222
L225:
	;
	F_computeLeafRecompressWALData(m, v168)
	mBase = m.M
	v1257 = m.ExcPending
	if v1257 != 0 {
		goto L29
	} else {
		goto L226
	}
L226:
	;
	goto L218
L227:
	;
	if v227 != 0 {
		goto L228
	} else {
		goto L229
	}
L228:
	;
	if v1261 == int32(0) {
		goto L231
	} else {
		goto L232
	}
L229:
	;
	goto L230
L230:
	;
	if v1261 == int32(0) {
		goto L239
	} else {
		goto L240
	}
L231:
	;
	v1718 = int32(1)
	v1724 = v1227
	goto L4
L232:
	;
	goto L233
L233:
	;
	if l1 < int32(0) {
		goto L235
	} else {
		goto L236
	}
L234:
	;
	v1285 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
	v1286 = *(*int32)(unsafe.Add(mBase, uint32(v70)))
	v1287 = *(*int32)(unsafe.Add(mBase, uint32(v168)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+72)) = v1287
	*(*int32)(unsafe.Add(mBase, uint32(v27)+68)) = v1284
	*(*int32)(unsafe.Add(mBase, uint32(v27)+64)) = v1227
	*(*int32)(unsafe.Add(mBase, uint32(v27)+76)) = v1285 - (v1286 + v1227)
	F_errmsg_internal(m, int32(_a_F_dataBeginPlaceToPage_10), v27-int32(-64))
	mBase = m.M
	v1298 = m.ExcPending
	if v1298 != 0 {
		goto L29
	} else {
		goto L238
	}
L235:
	;
	v1269 = *(*int32)(unsafe.Add(mBase, _c_F_dataBeginPlaceToPage[6]))
	v1275 = *(*int32)(unsafe.Add(mBase, uint32(v1269+(l1^int32(-1))<<(uint(int32(6))%32))+16))
	v1284 = v1275
	goto L234
L236:
	;
	goto L237
L237:
	;
	v1277 = *(*int32)(unsafe.Add(mBase, _c_F_dataBeginPlaceToPage[7]))
	v1283 = *(*int32)(unsafe.Add(mBase, uint32(v1277+l1<<(uint(int32(6))%32)+int32(-64))+16))
	v1284 = v1283
	goto L234
L238:
	;
	v1694 = int32(611)
	v1710 = int32(1)
	goto L214
L239:
	;
	v1718 = int32(1)
	v1724 = v1227
	goto L4
L240:
	;
	goto L241
L241:
	;
	if l1 < int32(0) {
		goto L243
	} else {
		goto L244
	}
L242:
	;
	v1323 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
	v1324 = *(*int32)(unsafe.Add(mBase, uint32(v70)))
	v1325 = *(*int32)(unsafe.Add(mBase, uint32(v168)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+88)) = v1325
	*(*int32)(unsafe.Add(mBase, uint32(v27)+84)) = v1322
	*(*int32)(unsafe.Add(mBase, uint32(v27)+80)) = v1227
	*(*int32)(unsafe.Add(mBase, uint32(v27)+92)) = v1323 - (v1324 + v1227)
	F_errmsg_internal(m, int32(_a_F_dataBeginPlaceToPage_11), v27+int32(80))
	mBase = m.M
	v1336 = m.ExcPending
	if v1336 != 0 {
		goto L29
	} else {
		goto L246
	}
L243:
	;
	v1307 = *(*int32)(unsafe.Add(mBase, _c_F_dataBeginPlaceToPage[6]))
	v1313 = *(*int32)(unsafe.Add(mBase, uint32(v1307+(l1^int32(-1))<<(uint(int32(6))%32))+16))
	v1322 = v1313
	goto L242
L244:
	;
	goto L245
L245:
	;
	v1315 = *(*int32)(unsafe.Add(mBase, _c_F_dataBeginPlaceToPage[7]))
	v1321 = *(*int32)(unsafe.Add(mBase, uint32(v1315+l1<<(uint(int32(6))%32)+int32(-64))+16))
	v1322 = v1321
	goto L242
L246:
	;
	v1694 = int32(615)
	v1710 = int32(1)
	goto L214
L247:
	;
	v1422 = *(*int32)(unsafe.Add(mBase, uint32(v1401)+24))
	if v1422 == int32(0) {
		goto L257
	} else {
		goto L258
	}
L248:
	;
	v1341 = *(*int32)(unsafe.Add(mBase, uint32(v1339)))
	if v1341 == v168 {
		v1401 = v1339
		goto L247
	} else {
		goto L249
	}
L249:
	;
	v1346 = v1339
	v1351 = v1341
	goto L250
L250:
	;
	v1367 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1346)+8)))
	if v1367 != int32(1) {
		goto L252
	} else {
		goto L253
	}
L251:
	;
	v1401 = v1394
	goto L247
L252:
	;
	v1370 = *(*int32)(unsafe.Add(mBase, uint32(v168)+12))
	v1371 = *(*int32)(unsafe.Add(mBase, uint32(v1346)+20))
	v1372 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1371)+6)))
	v1378 = (v1372+int32(1))&int32(_a_F_dataBeginPlaceToPage_6) + int32(8)
	v1379 = v1370 - v1378
	v1380 = *(*int32)(unsafe.Add(mBase, uint32(v168)+16))
	v1381 = v1378 + v1380
	if base.B2i32(v1379-v1381 < int32(0))|v227&base.B2i32(v1379 < int32(_a_F_dataBeginPlaceToPage_12)) != 0 {
		v1401 = v1346
		goto L247
	} else {
		goto L255
	}
L253:
	;
	v1394 = v1351
	goto L254
L254:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v168)+8)) = v1394
	v1396 = *(*int32)(unsafe.Add(mBase, uint32(v1394)))
	if v1396 != v168 {
		v1346 = v1394
		v1351 = v1396
		goto L250
	} else {
		goto L256
	}
L255:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v168)+16)) = v1381
	*(*int32)(unsafe.Add(mBase, uint32(v168)+12)) = v1379
	v1391 = *(*int32)(unsafe.Add(mBase, uint32(v1346)))
	v1394 = v1391
	goto L254
L256:
	;
	goto L251
L257:
	;
	v1425 = *(*int32)(unsafe.Add(mBase, uint32(v1401)+20))
	v1428 = F_ginPostingListDecode(m, v1425, v1401+int32(28))
	mBase = m.M
	v1429 = m.ExcPending
	if v1429 != 0 {
		goto L29
	} else {
		goto L260
	}
L258:
	;
	v1431 = v1422
	goto L259
L259:
	;
	v1432 = *(*int32)(unsafe.Add(mBase, uint32(v1401)+28))
	v1433 = int32(6)
	v1437 = v1431 + v1432*v1433 - v1433
	v1438 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1437)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v27)+116)) = uint16(v1438)
	v1440 = *(*int32)(unsafe.Add(mBase, uint32(v1437)))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+112)) = v1440
	v1443 = F_palloc(m, int32(_a_F_dataBeginPlaceToPage_13))
	mBase = m.M
	v1444 = m.ExcPending
	if v1444 != 0 {
		goto L29
	} else {
		goto L261
	}
L260:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1401)+24)) = v1428
	v1431 = v1428
	goto L259
L261:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v1443
	v1447 = F_palloc(m, int32(_a_F_dataBeginPlaceToPage_13))
	mBase = m.M
	v1448 = m.ExcPending
	if v1448 != 0 {
		goto L29
	} else {
		goto L262
	}
L262:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l7))) = v1447
	v1450 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	v1451 = int32(131)
	F_PageInit(m, v1450, int32(_a_F_dataBeginPlaceToPage_13), int32(8))
	mBase = m.M
	v1455 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1450)+16)))
	v1456 = v1450 + v1455
	*(*int32)(unsafe.Add(mBase, uint32(v1456))) = int32(-1)
	*(*uint16)(unsafe.Add(mBase, uint32(v1456)+6)) = uint16(v1451)
	goto L263
L263:
	;
	v1460 = int32(131)
	F_PageInit(m, v1447, int32(_a_F_dataBeginPlaceToPage_13), int32(8))
	mBase = m.M
	v1464 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1447)+16)))
	v1465 = v1447 + v1464
	*(*int32)(unsafe.Add(mBase, uint32(v1465))) = int32(-1)
	*(*uint16)(unsafe.Add(mBase, uint32(v1465)+6)) = uint16(v1460)
	goto L264
L264:
	;
	v1470 = v1450 + int32(24)
	v1471 = *(*int32)(unsafe.Add(mBase, uint32(v168)+4))
	v1472 = *(*int32)(unsafe.Add(mBase, uint32(v168)+8))
	v1473 = *(*int32)(unsafe.Add(mBase, uint32(v1472)+4))
	if v1471 != v1473 {
		goto L265
	} else {
		goto L266
	}
L265:
	;
	v1480 = v1450 + int32(32)
	v1486 = v1471
	v1490 = int32(0)
	goto L268
L266:
	;
	v1549 = int32(32)
	goto L267
L267:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v1450)+12)) = uint16(v1549)
	v1551 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v27)+116)))
	*(*uint16)(unsafe.Add(mBase, uint32(v1470)+4)) = uint16(v1551)
	v1553 = *(*int32)(unsafe.Add(mBase, uint32(v27)+112))
	*(*int32)(unsafe.Add(mBase, uint32(v1470))) = v1553
	v1560 = int32(0)
	v1561 = v1473
	v1562 = v1447 + int32(32)
	goto L277
L268:
	;
	v1502 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1486)+8)))
	if v1502 != int32(1) {
		goto L270
	} else {
		goto L271
	}
L269:
	;
	v1549 = v1519 + int32(32)
	goto L267
L270:
	;
	v1505 = *(*int32)(unsafe.Add(mBase, uint32(v1486)+20))
	v1506 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1505)+6)))
	v1512 = (v1506+int32(1))&int32(_a_F_dataBeginPlaceToPage_6) + int32(8)
	if v1512 != 0 {
		goto L273
	} else {
		goto L274
	}
L271:
	;
	v1516 = v1480
	v1519 = v1490
	goto L272
L272:
	;
	v1520 = *(*int32)(unsafe.Add(mBase, uint32(v1486)+4))
	if v1520 != v1473 {
		v1480 = v1516
		v1486 = v1520
		v1490 = v1519
		goto L268
	} else {
		goto L276
	}
L273:
	;
	base.MemoryCopy(m, v1480, v1505, v1512)
	goto L275
L274:
	;
	goto L275
L275:
	;
	v1516 = v1480 + v1512
	v1519 = v1512 + v1490
	goto L272
L276:
	;
	goto L269
L277:
	;
	v1582 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1561)+8)))
	if v1582 != int32(1) {
		goto L279
	} else {
		goto L280
	}
L278:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v1447)+28)) = uint16(v78)
	*(*uint16)(unsafe.Add(mBase, uint32(v1447)+26)) = uint16(v79)
	*(*uint16)(unsafe.Add(mBase, uint32(v1447)+24)) = uint16(v80)
	v1606 = v1596 + int32(32)
	*(*uint16)(unsafe.Add(mBase, uint32(v1447)+12)) = uint16(v1606)
	v1610 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v1611 = m.ExcPending
	if v1611 != 0 {
		goto L29
	} else {
		goto L286
	}
L279:
	;
	v1585 = *(*int32)(unsafe.Add(mBase, uint32(v1561)+20))
	v1586 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1585)+6)))
	v1592 = (v1586+int32(1))&int32(_a_F_dataBeginPlaceToPage_6) + int32(8)
	if v1592 != 0 {
		goto L282
	} else {
		goto L283
	}
L280:
	;
	v1596 = v1560
	v1597 = v1562
	goto L281
L281:
	;
	v1600 = *(*int32)(unsafe.Add(mBase, uint32(v1561)+4))
	if v1600 != v168 {
		v1560 = v1596
		v1561 = v1600
		v1562 = v1597
		goto L277
	} else {
		goto L285
	}
L282:
	;
	base.MemoryCopy(m, v1562, v1585, v1592)
	goto L284
L283:
	;
	goto L284
L284:
	;
	v1596 = v1560 + v1592
	v1597 = v1562 + v1592
	goto L281
L285:
	;
	goto L278
L286:
	;
	if v227 != 0 {
		goto L287
	} else {
		goto L288
	}
L287:
	;
	if v1610 == int32(0) {
		goto L290
	} else {
		goto L291
	}
L288:
	;
	goto L289
L289:
	;
	if v1610 == int32(0) {
		goto L298
	} else {
		goto L299
	}
L290:
	;
	v1718 = int32(2)
	v1724 = v1227
	goto L4
L291:
	;
	goto L292
L292:
	;
	if l1 < int32(0) {
		goto L294
	} else {
		goto L295
	}
L293:
	;
	v1634 = *(*int64)(unsafe.Add(mBase, uint32(v168)+12))
	v1635 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
	v1636 = *(*int32)(unsafe.Add(mBase, uint32(v70)))
	*(*int32)(unsafe.Add(mBase, uint32(v27))) = v1227
	*(*int32)(unsafe.Add(mBase, uint32(v27)+16)) = v1635 - (v1636 + v1227)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+4)) = v1633
	*(*int64)(unsafe.Add(mBase, uint32(v27)+8)) = v1634
	F_errmsg_internal(m, int32(_a_F_dataBeginPlaceToPage_14), v27)
	mBase = m.M
	v1645 = m.ExcPending
	if v1645 != 0 {
		goto L29
	} else {
		goto L297
	}
L294:
	;
	v1618 = *(*int32)(unsafe.Add(mBase, _c_F_dataBeginPlaceToPage[6]))
	v1624 = *(*int32)(unsafe.Add(mBase, uint32(v1618+(l1^int32(-1))<<(uint(int32(6))%32))+16))
	v1633 = v1624
	goto L293
L295:
	;
	goto L296
L296:
	;
	v1626 = *(*int32)(unsafe.Add(mBase, _c_F_dataBeginPlaceToPage[7]))
	v1632 = *(*int32)(unsafe.Add(mBase, uint32(v1626+l1<<(uint(int32(6))%32)+int32(-64))+16))
	v1633 = v1632
	goto L293
L297:
	;
	v1694 = int32(697)
	v1710 = int32(2)
	goto L214
L298:
	;
	v1718 = int32(2)
	v1724 = v1227
	goto L4
L299:
	;
	goto L300
L300:
	;
	if l1 < int32(0) {
		goto L302
	} else {
		goto L303
	}
L301:
	;
	v1670 = *(*int64)(unsafe.Add(mBase, uint32(v168)+12))
	v1671 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
	v1672 = *(*int32)(unsafe.Add(mBase, uint32(v70)))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+32)) = v1227
	*(*int32)(unsafe.Add(mBase, uint32(v27)+48)) = v1671 - (v1672 + v1227)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+36)) = v1669
	*(*int64)(unsafe.Add(mBase, uint32(v27)+40)) = v1670
	F_errmsg_internal(m, int32(_a_F_dataBeginPlaceToPage_15), v27+int32(32))
	mBase = m.M
	v1683 = m.ExcPending
	if v1683 != 0 {
		goto L29
	} else {
		goto L305
	}
L302:
	;
	v1654 = *(*int32)(unsafe.Add(mBase, _c_F_dataBeginPlaceToPage[6]))
	v1660 = *(*int32)(unsafe.Add(mBase, uint32(v1654+(l1^int32(-1))<<(uint(int32(6))%32))+16))
	v1669 = v1660
	goto L301
L303:
	;
	goto L304
L304:
	;
	v1662 = *(*int32)(unsafe.Add(mBase, _c_F_dataBeginPlaceToPage[7]))
	v1668 = *(*int32)(unsafe.Add(mBase, uint32(v1662+l1<<(uint(int32(6))%32)+int32(-64))+16))
	v1669 = v1668
	goto L301
L305:
	;
	v1694 = int32(701)
	v1710 = int32(2)
	goto L214
L306:
	;
	v1718 = v1710
	v1724 = v1227
	goto L4
L307:
	;
	F_errmsg_internal(m, int32(_a_F_dataBeginPlaceToPage_16), int32(0))
	mBase = m.M
	v1777 = m.ExcPending
	if v1777 != 0 {
		goto L29
	} else {
		goto L308
	}
L308:
	;
	F_errfinish(m, int32(_a_F_dataBeginPlaceToPage_8), int32(580), int32(_a_F_dataBeginPlaceToPage_9))
	mBase = m.M
	v1782 = m.ExcPending
	if v1782 != 0 {
		goto L29
	} else {
		goto L309
	}
L309:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L310:
	;
	F_errmsg_internal(m, int32(_a_F_dataBeginPlaceToPage_17), int32(0))
	mBase = m.M
	v1814 = m.ExcPending
	if v1814 != 0 {
		goto L29
	} else {
		goto L311
	}
L311:
	;
	F_errfinish(m, int32(_a_F_dataBeginPlaceToPage_8), int32(589), int32(_a_F_dataBeginPlaceToPage_9))
	mBase = m.M
	v1819 = m.ExcPending
	if v1819 != 0 {
		goto L29
	} else {
		goto L312
	}
L312:
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
	v4 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_data_sync_elevel[0])))
	if v4 != 0 {
		v5 = int32(21)
	} else {
		v5 = int32(23)
	}
	return v5
}
func F_show_data_directory_mode(m *base.Module) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn13994(m, int32(_a_F_show_data_directory_mode_0), int32(_a_F_show_data_directory_mode_1))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
