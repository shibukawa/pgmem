package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_gtrgm_compress(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int64
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
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
	var v87 int32
	_ = v87
	var v94 int32
	_ = v94
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v146 int64
	_ = v146
	v2 = int32(0)
	v6 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v7 = base.I32_wrap_i64(v6)
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v9 == v2 {
		v26 = v2
		goto L2
	} else {
		goto L3
	}
L1:
	;
	if v26&int32(1) != 0 {
		goto L7
	} else {
		goto L8
	}
L2:
	;
	goto L1
L3:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v9)+24))
	if v13 == int32(0) {
		v26 = v2
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	if v16 != int32(7) {
		v26 = v2
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if v19 != int32(17) {
		v26 = v2
		goto L2
	} else {
		goto L6
	}
L6:
	;
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+32)))
	v26 = v22 ^ int32(1)
	goto L2
L7:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v30 = F_get_fn_opclass_options(m, v29)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v35 = int32(12)
	goto L9
L9:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+18)))
	if v37 == int32(1) {
		goto L15
	} else {
		goto L16
	}
L10:
	;
	return int64(0)
L11:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	v35 = v34
	goto L9
L12:
	;
	return v146
L13:
	;
	v146 = base.I64_extend_i32_u(v136)
	goto L12
L14:
	;
	v124 = F_palloc(m, int32(24))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L10
	} else {
		goto L48
	}
L15:
	;
	v40 = F_pg_detoast_datum_packed(m, v36)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L10
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+4)))
	if v82&int32(6) != int32(2) {
		goto L37
	} else {
		goto L38
	}
L18:
	;
	v42 = int32(1)
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40))))
	if v44&v42 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v47 = v42
	goto L21
L20:
	;
	v47 = int32(4)
	goto L21
L21:
	;
	v48 = v40 + v47
	if v44 == int32(1) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+1)))
	if v54 == int32(18) {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	goto L24
L24:
	;
	if v44&int32(1) != 0 {
		goto L32
	} else {
		goto L33
	}
L25:
	;
	v57 = int32(16)
	goto L27
L26:
	;
	v57 = int32(0)
	goto L27
L27:
	;
	if base.Ui32((v54-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v64 = int32(4)
	goto L30
L29:
	;
	v64 = v57
	goto L30
L30:
	;
	v65 = F_generate_trgm(m, v48, v64)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L10
	} else {
		goto L31
	}
L31:
	;
	v119 = v65
	goto L14
L32:
	;
	v69 = int32(1)
	v73 = F_generate_trgm(m, v48, int32(base.Ui32(v44)>>(uint(v69)%32))-v69)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L10
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
	v80 = F_generate_trgm(m, v48, int32(base.Ui32(v75)>>(uint(int32(2))%32))-int32(4))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L10
	} else {
		goto L36
	}
L35:
	;
	v119 = v73
	goto L14
L36:
	;
	v119 = v80
	goto L14
L37:
	;
	v136 = v7
	goto L13
L38:
	;
	goto L39
L39:
	;
	v87 = int32(0)
	if v87 < v35 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v94 = v87
	goto L43
L41:
	;
	goto L42
L42:
	;
	v112 = F_palloc(m, int32(5))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L10
	} else {
		goto L47
	}
L43:
	;
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94+(v36+int32(5))))))
	if v100 != int32(255) {
		v146 = v6 & int64(4294967295)
		goto L12
	} else {
		goto L45
	}
L44:
	;
	goto L42
L45:
	;
	v104 = v94 + int32(1)
	if v104 != v35 {
		v94 = v104
		goto L43
	} else {
		goto L46
	}
L46:
	;
	goto L44
L47:
	;
	v114 = int32(6)
	*(*uint8)(unsafe.Add(mBase, uint32(v112)+4)) = uint8(v114)
	*(*int32)(unsafe.Add(mBase, uint32(v112))) = int32(20)
	v119 = v112
	goto L14
L48:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v124))) = base.I64_extend_i32_u(v119)
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v124)+8)) = v128
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v124)+12)) = v130
	v132 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v7)+16)))
	v133 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v124)+18)) = uint8(v133)
	*(*uint16)(unsafe.Add(mBase, uint32(v124)+16)) = uint16(v132)
	v136 = v124
	goto L13
}
func F_gtrgm_consistent(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int64
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v175 int32
	_ = v175
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v239 int32
	_ = v239
	var v243 int32
	_ = v243
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v270 int32
	_ = v270
	var v277 int32
	_ = v277
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
	var v313 int32
	_ = v313
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v335 int32
	_ = v335
	var v346 int32
	_ = v346
	var v350 int32
	_ = v350
	var v352 int32
	_ = v352
	var v354 int32
	_ = v354
	var v369 int32
	_ = v369
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v386 int32
	_ = v386
	var v389 int32
	_ = v389
	var v396 int32
	_ = v396
	var v401 int32
	_ = v401
	var v406 int32
	_ = v406
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v418 int32
	_ = v418
	var v421 int32
	_ = v421
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v433 int32
	_ = v433
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v451 int32
	_ = v451
	var v454 int32
	_ = v454
	var v461 int32
	_ = v461
	var v469 int32
	_ = v469
	var v484 int32
	_ = v484
	var v486 int32
	_ = v486
	var v488 int32
	_ = v488
	var v492 float64
	_ = v492
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v497 int32
	_ = v497
	var v500 int32
	_ = v500
	var v501 float32
	_ = v501
	var v502 int32
	_ = v502
	var v505 int32
	_ = v505
	var v509 int32
	_ = v509
	var v513 int32
	_ = v513
	var v517 int32
	_ = v517
	var v521 int32
	_ = v521
	var v523 int32
	_ = v523
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v529 int32
	_ = v529
	var v532 int32
	_ = v532
	var v547 int32
	_ = v547
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v555 int32
	_ = v555
	var v559 int32
	_ = v559
	var v563 int32
	_ = v563
	var v565 int32
	_ = v565
	var v567 int32
	_ = v567
	var v573 int32
	_ = v573
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v578 int32
	_ = v578
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v586 int32
	_ = v586
	var v590 int32
	_ = v590
	var v593 int32
	_ = v593
	var v597 int32
	_ = v597
	var v599 int32
	_ = v599
	var v602 int32
	_ = v602
	var v604 int32
	_ = v604
	var v623 int32
	_ = v623
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v631 int32
	_ = v631
	var v638 int32
	_ = v638
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v641 int32
	_ = v641
	var v645 int32
	_ = v645
	var v647 int32
	_ = v647
	var v650 int32
	_ = v650
	var v654 int32
	_ = v654
	var v656 int32
	_ = v656
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v661 int32
	_ = v661
	var v665 int32
	_ = v665
	var v668 int32
	_ = v668
	var v670 int32
	_ = v670
	var v689 int32
	_ = v689
	var v691 int32
	_ = v691
	var v692 int32
	_ = v692
	var v693 int32
	_ = v693
	var v697 int32
	_ = v697
	var v701 int32
	_ = v701
	var v705 int32
	_ = v705
	var v706 int32
	_ = v706
	var v709 int32
	_ = v709
	var v729 int32
	_ = v729
	var v730 int32
	_ = v730
	var v740 int32
	_ = v740
	var v741 int32
	_ = v741
	var v744 int32
	_ = v744
	var v745 int32
	_ = v745
	var v748 int32
	_ = v748
	var v751 int32
	_ = v751
	var v755 int32
	_ = v755
	var v756 int32
	_ = v756
	var v767 int32
	_ = v767
	var v768 int32
	_ = v768
	var v773 int32
	_ = v773
	var v784 int32
	_ = v784
	var v788 int32
	_ = v788
	var v790 int32
	_ = v790
	var v792 int32
	_ = v792
	var v807 int32
	_ = v807
	var v821 int32
	_ = v821
	var v822 int32
	_ = v822
	var v824 int32
	_ = v824
	var v827 int32
	_ = v827
	var v834 int32
	_ = v834
	var v839 int32
	_ = v839
	var v844 int32
	_ = v844
	var v847 int32
	_ = v847
	var v848 int32
	_ = v848
	var v851 int32
	_ = v851
	var v852 int32
	_ = v852
	var v856 int32
	_ = v856
	var v859 int32
	_ = v859
	var v868 int32
	_ = v868
	var v869 int32
	_ = v869
	var v871 int32
	_ = v871
	var v874 int32
	_ = v874
	var v875 int32
	_ = v875
	var v878 int32
	_ = v878
	var v879 int32
	_ = v879
	var v880 int32
	_ = v880
	var v889 int32
	_ = v889
	var v892 int32
	_ = v892
	var v899 int32
	_ = v899
	var v907 int32
	_ = v907
	var v922 int32
	_ = v922
	var v924 int32
	_ = v924
	var v929 int32
	_ = v929
	var v954 int32
	_ = v954
	var v960 int32
	_ = v960
	var v965 int32
	_ = v965
	v19 = m.G0
	v21 = v19 - int32(32)
	m.G0 = v21
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v25 = F_pg_detoast_datum(m, v24)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	v30 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v33 = int32(0)
	if v32 == v33 {
		v49 = v33
		goto L4
	} else {
		goto L5
	}
L3:
	;
	if v49&int32(1) != 0 {
		goto L9
	} else {
		goto L10
	}
L4:
	;
	goto L3
L5:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v32)+24))
	if v36 == int32(0) {
		v49 = v33
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
	if v39 != int32(7) {
		v49 = v33
		goto L4
	} else {
		goto L7
	}
L7:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v36)+4))
	if v42 != int32(17) {
		v49 = v33
		goto L4
	} else {
		goto L8
	}
L8:
	;
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+32)))
	v49 = v45 ^ int32(1)
	goto L4
L9:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v53 = F_get_fn_opclass_options(m, v52)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L1
	} else {
		goto L12
	}
L10:
	;
	v60 = int32(95)
	goto L11
L11:
	;
	v61 = base.I32_wrap_i64(v30)
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	v64 = int32(base.Ui32(v62) >> (uint(int32(2)) % 32))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v66)+16))
	if v67 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
	v60 = v55<<(uint(int32(3))%32) - int32(1)
	goto L11
L13:
	;
	v260 = v61 & int32(_a_F_gtrgm_consistent_0)
	if base.Ui32(int32(11)) < base.Ui32(v260) {
		goto L76
	} else {
		goto L77
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+28)) = int32(0)
	v145 = v61 & int32(_a_F_gtrgm_consistent_0)
	if base.Ui32(int32(11)) < base.Ui32(v145) {
		goto L43
	} else {
		goto L44
	}
L15:
	;
	v70 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v67))))
	if v70 != v61&int32(_a_F_gtrgm_consistent_0) {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v67)+4))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)))
	if int32(base.Ui32(v75)>>(uint(int32(2))%32)) != v64 {
		goto L14
	} else {
		goto L17
	}
L17:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v64) {
		goto L21
	} else {
		goto L22
	}
L18:
	;
	if v140 != 0 {
		goto L14
	} else {
		goto L36
	}
L19:
	;
	v140 = int32(0)
	goto L18
L20:
	;
	v114 = v109
	v115 = v110
	v116 = v111
	goto L30
L21:
	;
	if (v74|v25)&int32(3) != 0 {
		v109 = v74
		v110 = v25
		v111 = v64
		goto L20
	} else {
		goto L24
	}
L22:
	;
	v102 = v74
	v103 = v25
	v104 = v64
	goto L23
L23:
	;
	if v104 == int32(0) {
		goto L19
	} else {
		goto L29
	}
L24:
	;
	v86 = v74
	v87 = v25
	v88 = v64
	goto L25
L25:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v87)))
	if v91 != v92 {
		v109 = v86
		v110 = v87
		v111 = v88
		goto L20
	} else {
		goto L27
	}
L26:
	;
	v102 = v97
	v103 = v95
	v104 = v99
	goto L23
L27:
	;
	v94 = int32(4)
	v95 = v87 + v94
	v97 = v86 + v94
	v99 = v88 - v94
	if base.Ui32(int32(3)) < base.Ui32(v99) {
		v86 = v97
		v87 = v95
		v88 = v99
		goto L25
	} else {
		goto L28
	}
L28:
	;
	goto L26
L29:
	;
	v109 = v102
	v110 = v103
	v111 = v104
	goto L20
L30:
	;
	v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114))))
	v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115))))
	if v119 == v120 {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	v140 = v119 - v120
	goto L18
L32:
	;
	v122 = int32(1)
	v127 = v116 - v122
	if v127 != 0 {
		v114 = v114 + v122
		v115 = v115 + v122
		v116 = v127
		goto L30
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	goto L31
L35:
	;
	goto L19
L36:
	;
	v254 = v67
	goto L13
L37:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v221)+20))
	v226 = (v64 + int32(7)) & int32(2147483640)
	v230 = F_MemoryContextAlloc(m, v222, v226+v220+int32(16))
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L1
	} else {
		goto L60
	}
L38:
	;
	v216 = int32(0)
	v218 = int32(1)
	v219 = v216
	v220 = v216
	goto L37
L39:
	;
	v218 = int32(0)
	v219 = v212
	v220 = v213
	goto L37
L40:
	;
	if v203 == int32(0) {
		goto L57
	} else {
		goto L58
	}
L41:
	;
	v197 = int32(4)
	v201 = F_generate_wildcard_trgm(m, v25+v197, v64-v197)
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L1
	} else {
		goto L56
	}
L42:
	;
	v191 = int32(4)
	v195 = F_generate_trgm(m, v25+v191, v64-v191)
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L1
	} else {
		goto L55
	}
L43:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L1
	} else {
		goto L52
	}
L44:
	;
	v149 = int32(1) << (uint(v145) % 32)
	if v149&int32(2690) != 0 {
		goto L42
	} else {
		goto L45
	}
L45:
	;
	if v149&int32(24) != 0 {
		goto L41
	} else {
		goto L46
	}
L46:
	;
	if v149&int32(96) == int32(0) {
		goto L43
	} else {
		goto L47
	}
L47:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v66)+20))
	v163 = F_createTrgmNFA(m, v25, v159, v21+int32(28), v162)
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	if v163 == int32(0) {
		goto L38
	} else {
		goto L49
	}
L49:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v163)))
	v168 = int32(2)
	v169 = int32(base.Ui32(v167) >> (uint(v168) % 32))
	if base.Ui32(v168) < base.Ui32(v169-int32(5)) {
		v212 = v163
		v213 = v169
		goto L39
	} else {
		goto L50
	}
L50:
	;
	F_pfree(m, v163)
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	goto L38
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = v145
	F_errmsg_internal(m, int32(_a_F_gtrgm_consistent_1), v21)
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	F_errfinish(m, int32(_a_F_gtrgm_consistent_2), int32(272), int32(_a_F_gtrgm_consistent_3))
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L55:
	;
	v203 = v195
	goto L40
L56:
	;
	v203 = v201
	goto L40
L57:
	;
	v218 = int32(1)
	v219 = int32(0)
	v220 = int32(0)
	goto L37
L58:
	;
	goto L59
L59:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v203)))
	v212 = v203
	v213 = int32(base.Ui32(v208) >> (uint(int32(2)) % 32))
	goto L39
L60:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v230))) = uint16(v61)
	v234 = v230 + int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v230)+4)) = v234
	if v64 != 0 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	base.MemoryCopy(m, v234, v25, v64)
	goto L63
L62:
	;
	goto L63
L63:
	;
	if v218 == int32(0) {
		goto L65
	} else {
		goto L66
	}
L64:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v21)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v230)+12)) = v247
	if v67 != 0 {
		goto L72
	} else {
		goto L73
	}
L65:
	;
	v239 = v226 + v234
	*(*int32)(unsafe.Add(mBase, uint32(v230)+8)) = v239
	if v220 != 0 {
		goto L68
	} else {
		goto L69
	}
L66:
	;
	goto L67
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v230)+8)) = int32(0)
	goto L64
L68:
	;
	base.MemoryCopy(m, v239, v219, v220)
	goto L70
L69:
	;
	goto L70
L70:
	;
	F_pfree(m, v219)
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L1
	} else {
		goto L71
	}
L71:
	;
	goto L64
L72:
	;
	F_pfree(m, v67)
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L1
	} else {
		goto L75
	}
L73:
	;
	goto L74
L74:
	;
	v251 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v251)+16)) = v230
	v254 = v230
	goto L13
L75:
	;
	goto L74
L76:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v954 = m.ExcPending
	if v954 != 0 {
		goto L1
	} else {
		goto L202
	}
L77:
	;
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v254)+8))
	v265 = int32(1) << (uint(v260) % 32)
	if v265&int32(642) == int32(0) {
		goto L81
	} else {
		goto L82
	}
L78:
	;
	m.G0 = v21 + int32(32)
	return base.I64_extend_i32_u(v929)
L79:
	;
	v647 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65)+4)))
	if v647&int32(4) != 0 {
		v929 = v270
		goto L78
	} else {
		goto L155
	}
L80:
	;
	v573 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v29))) = uint8(v573)
	v575 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
	v576 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v575)+16)))
	v578 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v575+v576)+12)))
	if v578&v573 != 0 {
		goto L142
	} else {
		goto L143
	}
L81:
	;
	v270 = int32(1)
	if v265&int32(2072) != 0 {
		goto L80
	} else {
		goto L84
	}
L82:
	;
	goto L83
L83:
	;
	v488 = v61 & int32(_a_F_gtrgm_consistent_0)
	*(*uint8)(unsafe.Add(mBase, uint32(v29))) = uint8(base.B2i32(v488 != int32(1)))
	v492 = F_index_strategy_get_limit(m, v488)
	mBase = m.M
	v493 = m.ExcPending
	if v493 != 0 {
		goto L1
	} else {
		goto L125
	}
L84:
	;
	if v265&int32(96) == int32(0) {
		goto L76
	} else {
		goto L85
	}
L85:
	;
	v277 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v29))) = uint8(v277)
	if v263 == int32(0) {
		v929 = v270
		goto L78
	} else {
		goto L86
	}
L86:
	;
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
	v282 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v281)+16)))
	v284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v281+v282)+12)))
	if v284&int32(1) == int32(0) {
		goto L79
	} else {
		goto L87
	}
L87:
	;
	v289 = F_trgm_presence_map(m, v263, v65)
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L1
	} else {
		goto L88
	}
L88:
	;
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v254)+12))
	v292 = int32(0)
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v291)))
	if v302 != 0 {
		goto L90
	} else {
		goto L91
	}
L89:
	;
	F_pfree(m, v289)
	mBase = m.M
	v486 = m.ExcPending
	if v486 != 0 {
		goto L1
	} else {
		goto L124
	}
L90:
	;
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v291)+16))
	base.MemoryFill(m, v303, int32(0), v302)
	goto L92
L91:
	;
	goto L92
L92:
	;
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v291)+8))
	if v306 != 0 {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v291)+20))
	base.MemoryFill(m, v307, int32(0), v306)
	goto L95
L94:
	;
	goto L95
L95:
	;
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v291)))
	if int32(0) < v310 {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v291)+4))
	v317 = v292
	v318 = v292
	goto L99
L97:
	;
	goto L98
L98:
	;
	v383 = *(*int32)(unsafe.Add(mBase, uint32(v291)+20))
	v384 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v383))) = uint8(v384)
	v386 = *(*int32)(unsafe.Add(mBase, uint32(v291)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v386))) = int32(0)
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v291)+12))
	v396 = v384
	v401 = v292
	goto L111
L99:
	;
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v313+v317<<(uint(int32(2))%32))))
	v330 = v329 + v318
	if v329 <= int32(0) {
		goto L101
	} else {
		goto L102
	}
L100:
	;
	goto L98
L101:
	;
	v369 = v317 + int32(1)
	if v369 != v310 {
		v317 = v369
		v318 = v330
		goto L99
	} else {
		goto L109
	}
L102:
	;
	v335 = v318
	goto L103
L103:
	;
	v346 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v289+v335))))
	if v346 != int32(1) {
		goto L105
	} else {
		goto L106
	}
L104:
	;
	v352 = *(*int32)(unsafe.Add(mBase, uint32(v291)+16))
	v354 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v352+v317))) = uint8(v354)
	goto L101
L105:
	;
	v350 = v335 + int32(1)
	if v350 < v330 {
		v335 = v350
		goto L103
	} else {
		goto L108
	}
L106:
	;
	goto L107
L107:
	;
	goto L104
L108:
	;
	goto L101
L109:
	;
	goto L100
L110:
	;
	goto L89
L111:
	;
	v406 = *(*int32)(unsafe.Add(mBase, uint32(v386+v401<<(uint(int32(2))%32))))
	v409 = v389 + v406<<(uint(int32(3))%32)
	v410 = *(*int32)(unsafe.Add(mBase, uint32(v409)))
	if int32(0) < v410 {
		goto L113
	} else {
		goto L114
	}
L112:
	;
	v484 = int32(0)
	goto L110
L113:
	;
	v413 = *(*int32)(unsafe.Add(mBase, uint32(v291)+16))
	v414 = *(*int32)(unsafe.Add(mBase, uint32(v409)+4))
	v418 = int32(0)
	v421 = v396
	goto L116
L114:
	;
	v461 = v396
	goto L115
L115:
	;
	v469 = v401 + int32(1)
	if v469 < v461 {
		v396 = v461
		v401 = v469
		goto L111
	} else {
		goto L123
	}
L116:
	;
	v430 = v414 + v418<<(uint(int32(3))%32)
	v431 = *(*int32)(unsafe.Add(mBase, uint32(v430)+4))
	v433 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v413+v431))))
	if v433 != int32(1) {
		v451 = v421
		goto L118
	} else {
		goto L119
	}
L117:
	;
	v461 = v451
	goto L115
L118:
	;
	v454 = v418 + int32(1)
	if v454 != v410 {
		v418 = v454
		v421 = v451
		goto L116
	} else {
		goto L122
	}
L119:
	;
	v436 = int32(1)
	v437 = *(*int32)(unsafe.Add(mBase, uint32(v430)))
	if v437 == v436 {
		v484 = v436
		goto L110
	} else {
		goto L120
	}
L120:
	;
	v440 = v383 + v437
	v441 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v440))))
	if v441 != 0 {
		v451 = v421
		goto L118
	} else {
		goto L121
	}
L121:
	;
	v442 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v440))) = uint8(v442)
	*(*int32)(unsafe.Add(mBase, uint32(v386+v421<<(uint(int32(2))%32)))) = v437
	v451 = v421 + v442
	goto L118
L122:
	;
	goto L117
L123:
	;
	goto L112
L124:
	;
	v929 = v484
	goto L78
L125:
	;
	v494 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
	v495 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v494)+16)))
	v497 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v494+v495)+12)))
	if v497&int32(1) != 0 {
		goto L126
	} else {
		goto L127
	}
L126:
	;
	v500 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29))))
	v501 = F_cnt_sml(m, v263, v65, v500)
	mBase = m.M
	v502 = m.ExcPending
	if v502 != 0 {
		goto L1
	} else {
		goto L129
	}
L127:
	;
	goto L128
L128:
	;
	v505 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65)+4)))
	if v505&int32(4) != 0 {
		goto L130
	} else {
		goto L131
	}
L129:
	;
	v929 = base.F64_le(v492, base.F64_promote_f32(v501))
	goto L78
L130:
	;
	v929 = int32(1)
	goto L78
L131:
	;
	goto L132
L132:
	;
	v509 = *(*int32)(unsafe.Add(mBase, uint32(v263)))
	v513 = int32(base.Ui32(v509)>>(uint(int32(2))%32)) - int32(5)
	if base.Ui32(v513) < base.Ui32(int32(3)) {
		goto L133
	} else {
		goto L134
	}
L133:
	;
	v929 = int32(0)
	goto L78
L134:
	;
	goto L135
L135:
	;
	v517 = int32(5)
	v521 = int32(1)
	v523 = base.I32_div_u_s(v513, int32(3))
	if base.Ui32(v523) <= base.Ui32(v521) {
		goto L136
	} else {
		goto L137
	}
L136:
	;
	v526 = v521
	goto L138
L137:
	;
	v526 = v523
	goto L138
L138:
	;
	v527 = int32(0)
	v529 = v527
	v532 = v527
	goto L139
L139:
	;
	v547 = int32(3)
	v549 = v263 + v517 + v529*v547
	v550 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v549))))
	v551 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v549)+2)))
	v555 = base.I32_rem_u_s(v550|v551<<(uint(int32(16))%32), v60)
	v559 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65+v517+int32(base.Ui32(v555)>>(uint(v547)%32))))))
	v563 = int32(1)
	v565 = int32(base.Ui32(v559)>>(uint(v555&int32(7))%32))&v563 + v532
	v567 = v529 + v563
	if v567 != v526 {
		v529 = v567
		v532 = v565
		goto L139
	} else {
		goto L141
	}
L140:
	;
	v929 = base.F64_ge(base.F64_div(base.F64_convert_i32_u(v565), base.F64_convert_i32_u(v523)), v492)
	goto L78
L141:
	;
	goto L140
L142:
	;
	v581 = F_trgm_contained_by(m, v263, v65)
	mBase = m.M
	v582 = m.ExcPending
	if v582 != 0 {
		goto L1
	} else {
		goto L145
	}
L143:
	;
	goto L144
L144:
	;
	v583 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65)+4)))
	if v583&int32(4) != 0 {
		v929 = v270
		goto L78
	} else {
		goto L146
	}
L145:
	;
	v929 = v581
	goto L78
L146:
	;
	v586 = *(*int32)(unsafe.Add(mBase, uint32(v263)))
	v590 = int32(base.Ui32(v586)>>(uint(int32(2))%32)) - int32(5)
	if base.Ui32(v590) < base.Ui32(int32(3)) {
		v929 = v270
		goto L78
	} else {
		goto L147
	}
L147:
	;
	v593 = int32(5)
	v597 = int32(1)
	v599 = base.I32_div_u_s(v590, int32(3))
	if base.Ui32(v599) <= base.Ui32(v597) {
		goto L148
	} else {
		goto L149
	}
L148:
	;
	v602 = v597
	goto L150
L149:
	;
	v602 = v599
	goto L150
L150:
	;
	v604 = int32(0)
	goto L151
L151:
	;
	v623 = int32(3)
	v625 = v263 + v593 + v604*v623
	v626 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v625))))
	v627 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v625)+2)))
	v631 = base.I32_rem_u_s(v626|v627<<(uint(int32(16))%32), v60)
	v638 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65+v593+int32(base.Ui32(v631)>>(uint(v623)%32))))))
	v639 = int32(1) << (uint(v631&int32(7)) % 32) & v638
	v640 = int32(0)
	v641 = base.B2i32(v639 != v640)
	if v639 == v640 {
		v929 = v641
		goto L78
	} else {
		goto L153
	}
L152:
	;
	v929 = v641
	goto L78
L153:
	;
	v645 = v604 + int32(1)
	if v645 != v602 {
		v604 = v645
		goto L151
	} else {
		goto L154
	}
L154:
	;
	goto L152
L155:
	;
	v650 = *(*int32)(unsafe.Add(mBase, uint32(v263)))
	v654 = int32(base.Ui32(v650)>>(uint(int32(2))%32)) - int32(5)
	v656 = base.I32_div_u_s(v654, int32(3))
	v657 = F_palloc(m, v656)
	mBase = m.M
	v658 = m.ExcPending
	if v658 != 0 {
		goto L1
	} else {
		goto L156
	}
L156:
	;
	if base.Ui32(int32(3)) <= base.Ui32(v654) {
		goto L157
	} else {
		goto L158
	}
L157:
	;
	v661 = int32(5)
	v665 = int32(1)
	if base.Ui32(v656) <= base.Ui32(v665) {
		goto L160
	} else {
		goto L161
	}
L158:
	;
	goto L159
L159:
	;
	v729 = *(*int32)(unsafe.Add(mBase, uint32(v254)+12))
	v730 = int32(0)
	v740 = *(*int32)(unsafe.Add(mBase, uint32(v729)))
	if v740 != 0 {
		goto L167
	} else {
		goto L168
	}
L160:
	;
	v668 = v665
	goto L162
L161:
	;
	v668 = v656
	goto L162
L162:
	;
	v670 = int32(0)
	goto L163
L163:
	;
	v689 = int32(3)
	v691 = v263 + v661 + v670*v689
	v692 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v691))))
	v693 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v691)+2)))
	v697 = base.I32_rem_u_s(v692|v693<<(uint(int32(16))%32), v60)
	v701 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65+v661+int32(base.Ui32(v697)>>(uint(v689)%32))))))
	v705 = int32(1)
	v706 = int32(base.Ui32(v701)>>(uint(v697&int32(7))%32)) & v705
	*(*uint8)(unsafe.Add(mBase, uint32(v670+v657))) = uint8(v706)
	v709 = v670 + v705
	if v709 != v668 {
		v670 = v709
		goto L163
	} else {
		goto L165
	}
L164:
	;
	goto L159
L165:
	;
	goto L164
L166:
	;
	F_pfree(m, v657)
	mBase = m.M
	v924 = m.ExcPending
	if v924 != 0 {
		goto L1
	} else {
		goto L201
	}
L167:
	;
	v741 = *(*int32)(unsafe.Add(mBase, uint32(v729)+16))
	base.MemoryFill(m, v741, int32(0), v740)
	goto L169
L168:
	;
	goto L169
L169:
	;
	v744 = *(*int32)(unsafe.Add(mBase, uint32(v729)+8))
	if v744 != 0 {
		goto L170
	} else {
		goto L171
	}
L170:
	;
	v745 = *(*int32)(unsafe.Add(mBase, uint32(v729)+20))
	base.MemoryFill(m, v745, int32(0), v744)
	goto L172
L171:
	;
	goto L172
L172:
	;
	v748 = *(*int32)(unsafe.Add(mBase, uint32(v729)))
	if int32(0) < v748 {
		goto L173
	} else {
		goto L174
	}
L173:
	;
	v751 = *(*int32)(unsafe.Add(mBase, uint32(v729)+4))
	v755 = v730
	v756 = v730
	goto L176
L174:
	;
	goto L175
L175:
	;
	v821 = *(*int32)(unsafe.Add(mBase, uint32(v729)+20))
	v822 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v821))) = uint8(v822)
	v824 = *(*int32)(unsafe.Add(mBase, uint32(v729)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v824))) = int32(0)
	v827 = *(*int32)(unsafe.Add(mBase, uint32(v729)+12))
	v834 = v822
	v839 = v730
	goto L188
L176:
	;
	v767 = *(*int32)(unsafe.Add(mBase, uint32(v751+v755<<(uint(int32(2))%32))))
	v768 = v767 + v756
	if v767 <= int32(0) {
		goto L178
	} else {
		goto L179
	}
L177:
	;
	goto L175
L178:
	;
	v807 = v755 + int32(1)
	if v807 != v748 {
		v755 = v807
		v756 = v768
		goto L176
	} else {
		goto L186
	}
L179:
	;
	v773 = v756
	goto L180
L180:
	;
	v784 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v657+v773))))
	if v784 != int32(1) {
		goto L182
	} else {
		goto L183
	}
L181:
	;
	v790 = *(*int32)(unsafe.Add(mBase, uint32(v729)+16))
	v792 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v790+v755))) = uint8(v792)
	goto L178
L182:
	;
	v788 = v773 + int32(1)
	if v788 < v768 {
		v773 = v788
		goto L180
	} else {
		goto L185
	}
L183:
	;
	goto L184
L184:
	;
	goto L181
L185:
	;
	goto L178
L186:
	;
	goto L177
L187:
	;
	goto L166
L188:
	;
	v844 = *(*int32)(unsafe.Add(mBase, uint32(v824+v839<<(uint(int32(2))%32))))
	v847 = v827 + v844<<(uint(int32(3))%32)
	v848 = *(*int32)(unsafe.Add(mBase, uint32(v847)))
	if int32(0) < v848 {
		goto L190
	} else {
		goto L191
	}
L189:
	;
	v922 = int32(0)
	goto L187
L190:
	;
	v851 = *(*int32)(unsafe.Add(mBase, uint32(v729)+16))
	v852 = *(*int32)(unsafe.Add(mBase, uint32(v847)+4))
	v856 = int32(0)
	v859 = v834
	goto L193
L191:
	;
	v899 = v834
	goto L192
L192:
	;
	v907 = v839 + int32(1)
	if v907 < v899 {
		v834 = v899
		v839 = v907
		goto L188
	} else {
		goto L200
	}
L193:
	;
	v868 = v852 + v856<<(uint(int32(3))%32)
	v869 = *(*int32)(unsafe.Add(mBase, uint32(v868)+4))
	v871 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v851+v869))))
	if v871 != int32(1) {
		v889 = v859
		goto L195
	} else {
		goto L196
	}
L194:
	;
	v899 = v889
	goto L192
L195:
	;
	v892 = v856 + int32(1)
	if v892 != v848 {
		v856 = v892
		v859 = v889
		goto L193
	} else {
		goto L199
	}
L196:
	;
	v874 = int32(1)
	v875 = *(*int32)(unsafe.Add(mBase, uint32(v868)))
	if v875 == v874 {
		v922 = v874
		goto L187
	} else {
		goto L197
	}
L197:
	;
	v878 = v821 + v875
	v879 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v878))))
	if v879 != 0 {
		v889 = v859
		goto L195
	} else {
		goto L198
	}
L198:
	;
	v880 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v878))) = uint8(v880)
	*(*int32)(unsafe.Add(mBase, uint32(v824+v859<<(uint(int32(2))%32)))) = v875
	v889 = v859 + v880
	goto L195
L199:
	;
	goto L194
L200:
	;
	goto L189
L201:
	;
	v929 = v922
	goto L78
L202:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+16)) = v260
	F_errmsg_internal(m, int32(_a_F_gtrgm_consistent_1), v21+int32(16))
	mBase = m.M
	v960 = m.ExcPending
	if v960 != 0 {
		goto L1
	} else {
		goto L203
	}
L203:
	;
	F_errfinish(m, int32(_a_F_gtrgm_consistent_2), int32(444), int32(_a_F_gtrgm_consistent_3))
	mBase = m.M
	v965 = m.ExcPending
	if v965 != 0 {
		goto L1
	} else {
		goto L204
	}
L204:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_gtrgm_out(m *base.Module, l0 int32) int64 {
	var v7 int64
	_ = v7
	var v10 int32
	_ = v10
	v7 = Fn14235(m, l0, int32(_a_F_gtrgm_out_0), int32(71), int32(_a_F_gtrgm_out_1), int32(_a_F_gtrgm_out_2), int32(_a_F_gtrgm_out_3))
	v10 = m.ExcPending
	if v10 != 0 {
		return int64(0)
	} else {
		return v7
	}
}
