package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_gtrgm_compress(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
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
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	v2 = int32(0)
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
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
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+24)))
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
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+14)))
	if v37 == int32(1) {
		goto L13
	} else {
		goto L14
	}
L10:
	;
	return int32(0)
L11:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	v35 = v34
	goto L9
L12:
	;
	v126 = F_palloc(m, int32(16))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L10
	} else {
		goto L47
	}
L13:
	;
	v40 = F_pg_detoast_datum_packed(m, v36)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L10
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+4)))
	if v81&int32(6) != int32(2) {
		goto L35
	} else {
		goto L36
	}
L16:
	;
	v42 = int32(1)
	v43 = v40 + v42
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40))))
	v48 = v46 & v42
	if v48 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v49 = v43
	goto L19
L18:
	;
	v49 = v40 + int32(4)
	goto L19
L19:
	;
	if v46 == int32(1) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43))))
	if v55 == int32(18) {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	goto L22
L22:
	;
	if v48 != 0 {
		goto L30
	} else {
		goto L31
	}
L23:
	;
	v58 = int32(16)
	goto L25
L24:
	;
	v58 = int32(0)
	goto L25
L25:
	;
	if base.Ui32((v55-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v65 = int32(4)
	goto L28
L27:
	;
	v65 = v58
	goto L28
L28:
	;
	v66 = F_generate_trgm(m, v49, v65)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L10
	} else {
		goto L29
	}
L29:
	;
	v120 = v66
	goto L12
L30:
	;
	v68 = int32(1)
	v72 = F_generate_trgm(m, v49, int32(base.Ui32(v46)>>(uint(v68)%32))-v68)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L10
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
	v79 = F_generate_trgm(m, v49, int32(base.Ui32(v74)>>(uint(int32(2))%32))-int32(4))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L10
	} else {
		goto L34
	}
L33:
	;
	v120 = v72
	goto L12
L34:
	;
	v120 = v79
	goto L12
L35:
	;
	return v7
L36:
	;
	goto L37
L37:
	;
	v87 = int32(0)
	if v35 <= v87 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v113 = F_palloc(m, int32(5))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L10
	} else {
		goto L46
	}
L39:
	;
	v92 = v87
	goto L40
L40:
	;
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92+(v36+int32(5))))))
	if v99 == int32(255) {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	return v7
L42:
	;
	v103 = v92 + int32(1)
	if v35 != v103 {
		v92 = v103
		goto L40
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	goto L41
L45:
	;
	goto L38
L46:
	;
	v115 = int32(6)
	*(*uint8)(unsafe.Add(mBase, uint32(v113)+4)) = uint8(v115)
	*(*int32)(unsafe.Add(mBase, uint32(v113))) = int32(20)
	v120 = v113
	goto L12
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v126))) = v120
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v126)+4)) = v129
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v126)+8)) = v131
	v133 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v7)+12)))
	v134 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v126)+14)) = uint8(v134)
	*(*uint16)(unsafe.Add(mBase, uint32(v126)+12)) = uint16(v133)
	return v126
}
func F_gtrgm_consistent(m *base.Module, l0 int32) int32 {
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
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
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
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v136 int32
	_ = v136
	var v143 int32
	_ = v143
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v169 int32
	_ = v169
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
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
	var v197 int32
	_ = v197
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v263 int32
	_ = v263
	var v270 int32
	_ = v270
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v328 int32
	_ = v328
	var v339 int32
	_ = v339
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v362 int32
	_ = v362
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v379 int32
	_ = v379
	var v382 int32
	_ = v382
	var v389 int32
	_ = v389
	var v394 int32
	_ = v394
	var v399 int32
	_ = v399
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v411 int32
	_ = v411
	var v414 int32
	_ = v414
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v426 int32
	_ = v426
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v444 int32
	_ = v444
	var v447 int32
	_ = v447
	var v454 int32
	_ = v454
	var v462 int32
	_ = v462
	var v477 int32
	_ = v477
	var v479 int32
	_ = v479
	var v483 float64
	_ = v483
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v488 int32
	_ = v488
	var v491 int32
	_ = v491
	var v492 float32
	_ = v492
	var v493 int32
	_ = v493
	var v496 int32
	_ = v496
	var v500 int32
	_ = v500
	var v504 int32
	_ = v504
	var v508 int32
	_ = v508
	var v512 int32
	_ = v512
	var v514 int32
	_ = v514
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v537 int32
	_ = v537
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v545 int32
	_ = v545
	var v549 int32
	_ = v549
	var v553 int32
	_ = v553
	var v555 int32
	_ = v555
	var v557 int32
	_ = v557
	var v563 int32
	_ = v563
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v568 int32
	_ = v568
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v576 int32
	_ = v576
	var v580 int32
	_ = v580
	var v583 int32
	_ = v583
	var v587 int32
	_ = v587
	var v589 int32
	_ = v589
	var v592 int32
	_ = v592
	var v594 int32
	_ = v594
	var v612 int32
	_ = v612
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v616 int32
	_ = v616
	var v620 int32
	_ = v620
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v634 int32
	_ = v634
	var v636 int32
	_ = v636
	var v639 int32
	_ = v639
	var v643 int32
	_ = v643
	var v645 int32
	_ = v645
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v650 int32
	_ = v650
	var v654 int32
	_ = v654
	var v657 int32
	_ = v657
	var v659 int32
	_ = v659
	var v677 int32
	_ = v677
	var v679 int32
	_ = v679
	var v680 int32
	_ = v680
	var v681 int32
	_ = v681
	var v685 int32
	_ = v685
	var v689 int32
	_ = v689
	var v693 int32
	_ = v693
	var v694 int32
	_ = v694
	var v697 int32
	_ = v697
	var v716 int32
	_ = v716
	var v717 int32
	_ = v717
	var v727 int32
	_ = v727
	var v728 int32
	_ = v728
	var v731 int32
	_ = v731
	var v732 int32
	_ = v732
	var v735 int32
	_ = v735
	var v738 int32
	_ = v738
	var v742 int32
	_ = v742
	var v743 int32
	_ = v743
	var v754 int32
	_ = v754
	var v755 int32
	_ = v755
	var v760 int32
	_ = v760
	var v771 int32
	_ = v771
	var v775 int32
	_ = v775
	var v777 int32
	_ = v777
	var v779 int32
	_ = v779
	var v794 int32
	_ = v794
	var v808 int32
	_ = v808
	var v809 int32
	_ = v809
	var v811 int32
	_ = v811
	var v814 int32
	_ = v814
	var v821 int32
	_ = v821
	var v826 int32
	_ = v826
	var v831 int32
	_ = v831
	var v834 int32
	_ = v834
	var v835 int32
	_ = v835
	var v838 int32
	_ = v838
	var v839 int32
	_ = v839
	var v843 int32
	_ = v843
	var v846 int32
	_ = v846
	var v855 int32
	_ = v855
	var v856 int32
	_ = v856
	var v858 int32
	_ = v858
	var v861 int32
	_ = v861
	var v862 int32
	_ = v862
	var v865 int32
	_ = v865
	var v866 int32
	_ = v866
	var v867 int32
	_ = v867
	var v876 int32
	_ = v876
	var v879 int32
	_ = v879
	var v886 int32
	_ = v886
	var v894 int32
	_ = v894
	var v909 int32
	_ = v909
	var v911 int32
	_ = v911
	var v914 int32
	_ = v914
	var v939 int32
	_ = v939
	var v945 int32
	_ = v945
	var v950 int32
	_ = v950
	v18 = m.G0
	v20 = v18 - int32(32)
	m.G0 = v20
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v24 = F_pg_detoast_datum(m, v23)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v29 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+36)))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v32 = int32(0)
	if v31 == v32 {
		v48 = v32
		goto L4
	} else {
		goto L5
	}
L3:
	;
	if v48&int32(1) != 0 {
		goto L9
	} else {
		goto L10
	}
L4:
	;
	goto L3
L5:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v31)+24))
	if v35 == int32(0) {
		v48 = v32
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	if v38 != int32(7) {
		v48 = v32
		goto L4
	} else {
		goto L7
	}
L7:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
	if v41 != int32(17) {
		v48 = v32
		goto L4
	} else {
		goto L8
	}
L8:
	;
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+24)))
	v48 = v44 ^ int32(1)
	goto L4
L9:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v52 = F_get_fn_opclass_options(m, v51)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L1
	} else {
		goto L12
	}
L10:
	;
	v59 = int32(95)
	goto L11
L11:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	v62 = int32(base.Ui32(v60) >> (uint(int32(2)) % 32))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+16))
	if v65 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v52)+4))
	v59 = v54<<(uint(int32(3))%32) - int32(1)
	goto L11
L13:
	;
	if base.Ui32(int32(11)) < base.Ui32(v29) {
		goto L76
	} else {
		goto L77
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+28)) = int32(0)
	if base.Ui32(int32(11)) < base.Ui32(v29) {
		goto L43
	} else {
		goto L44
	}
L15:
	;
	v68 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v65))))
	if v68 != v29 {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v65)+4))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)))
	if int32(base.Ui32(v71)>>(uint(int32(2))%32)) != v62 {
		goto L14
	} else {
		goto L17
	}
L17:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v62) {
		goto L21
	} else {
		goto L22
	}
L18:
	;
	if v136 != 0 {
		goto L14
	} else {
		goto L36
	}
L19:
	;
	v136 = int32(0)
	goto L18
L20:
	;
	v110 = v105
	v111 = v106
	v112 = v107
	goto L30
L21:
	;
	if (v70|v24)&int32(3) != 0 {
		v105 = v70
		v106 = v24
		v107 = v62
		goto L20
	} else {
		goto L24
	}
L22:
	;
	v98 = v70
	v99 = v24
	v100 = v62
	goto L23
L23:
	;
	if v100 == int32(0) {
		goto L19
	} else {
		goto L29
	}
L24:
	;
	v82 = v70
	v83 = v24
	v84 = v62
	goto L25
L25:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
	if v87 != v88 {
		v105 = v82
		v106 = v83
		v107 = v84
		goto L20
	} else {
		goto L27
	}
L26:
	;
	v98 = v93
	v99 = v91
	v100 = v95
	goto L23
L27:
	;
	v90 = int32(4)
	v91 = v83 + v90
	v93 = v82 + v90
	v95 = v84 - v90
	if base.Ui32(int32(3)) < base.Ui32(v95) {
		v82 = v93
		v83 = v91
		v84 = v95
		goto L25
	} else {
		goto L28
	}
L28:
	;
	goto L26
L29:
	;
	v105 = v98
	v106 = v99
	v107 = v100
	goto L20
L30:
	;
	v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v110))))
	v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111))))
	if v115 == v116 {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	v136 = v115 - v116
	goto L18
L32:
	;
	v118 = int32(1)
	v123 = v112 - v118
	if v123 != 0 {
		v110 = v110 + v118
		v111 = v111 + v118
		v112 = v123
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
	v247 = v65
	goto L13
L37:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v215)+20))
	v220 = (v62 + int32(7)) & int32(2147483640)
	v224 = F_MemoryContextAlloc(m, v216, v220+v213+int32(16))
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L1
	} else {
		goto L60
	}
L38:
	;
	v210 = int32(0)
	v212 = v210
	v213 = v210
	v214 = int32(1)
	goto L37
L39:
	;
	v212 = v205
	v213 = v206
	v214 = int32(0)
	goto L37
L40:
	;
	if v197 == int32(0) {
		goto L57
	} else {
		goto L58
	}
L41:
	;
	v191 = int32(4)
	v195 = F_generate_wildcard_trgm(m, v24+v191, v62-v191)
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L1
	} else {
		goto L56
	}
L42:
	;
	v185 = int32(4)
	v189 = F_generate_trgm(m, v24+v185, v62-v185)
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L1
	} else {
		goto L55
	}
L43:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L1
	} else {
		goto L52
	}
L44:
	;
	v143 = int32(1) << (uint(v29) % 32)
	if v143&int32(2690) != 0 {
		goto L42
	} else {
		goto L45
	}
L45:
	;
	if v143&int32(24) != 0 {
		goto L41
	} else {
		goto L46
	}
L46:
	;
	if v143&int32(96) == int32(0) {
		goto L43
	} else {
		goto L47
	}
L47:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v64)+20))
	v157 = F_createTrgmNFA(m, v24, v153, v20+int32(28), v156)
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	if v157 == int32(0) {
		goto L38
	} else {
		goto L49
	}
L49:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v157)))
	v162 = int32(2)
	v163 = int32(base.Ui32(v161) >> (uint(v162) % 32))
	if base.Ui32(v162) < base.Ui32(v163-int32(5)) {
		v205 = v157
		v206 = v163
		goto L39
	} else {
		goto L50
	}
L50:
	;
	F_pfree(m, v157)
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	goto L38
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v29
	F_errmsg_internal(m, int32(_a_F_gtrgm_consistent_0), v20)
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	F_errfinish(m, int32(_a_F_gtrgm_consistent_1), int32(271), int32(_a_F_gtrgm_consistent_2))
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
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
	v197 = v189
	goto L40
L56:
	;
	v197 = v195
	goto L40
L57:
	;
	v212 = int32(0)
	v213 = int32(0)
	v214 = int32(1)
	goto L37
L58:
	;
	goto L59
L59:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v197)))
	v205 = v197
	v206 = int32(base.Ui32(v202) >> (uint(int32(2)) % 32))
	goto L39
L60:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v224))) = uint16(v29)
	v228 = v224 + int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v224)+4)) = v228
	if v62 != 0 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	base.MemoryCopy(m, v228, v24, v62)
	goto L63
L62:
	;
	goto L63
L63:
	;
	if v214 == int32(0) {
		goto L65
	} else {
		goto L66
	}
L64:
	;
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v20)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v224)+12)) = v241
	if v65 != 0 {
		goto L72
	} else {
		goto L73
	}
L65:
	;
	v233 = v228 + v220
	*(*int32)(unsafe.Add(mBase, uint32(v224)+8)) = v233
	if v213 != 0 {
		goto L68
	} else {
		goto L69
	}
L66:
	;
	goto L67
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v224)+8)) = int32(0)
	goto L64
L68:
	;
	base.MemoryCopy(m, v233, v212, v213)
	goto L70
L69:
	;
	goto L70
L70:
	;
	F_pfree(m, v212)
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L1
	} else {
		goto L71
	}
L71:
	;
	goto L64
L72:
	;
	F_pfree(m, v65)
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L1
	} else {
		goto L75
	}
L73:
	;
	goto L74
L74:
	;
	v245 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v245)+16)) = v224
	v247 = v224
	goto L13
L75:
	;
	goto L74
L76:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v939 = m.ExcPending
	if v939 != 0 {
		goto L1
	} else {
		goto L202
	}
L77:
	;
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v247)+8))
	v258 = int32(1) << (uint(v29) % 32)
	if v258&int32(642) == int32(0) {
		goto L81
	} else {
		goto L82
	}
L78:
	;
	m.G0 = v20 + int32(32)
	return v914
L79:
	;
	v636 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63)+4)))
	if v636&int32(4) != 0 {
		v914 = v263
		goto L78
	} else {
		goto L155
	}
L80:
	;
	v563 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v28))) = uint8(v563)
	v565 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
	v566 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v565)+16)))
	v568 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v565+v566)+12)))
	if v568&v563 != 0 {
		goto L142
	} else {
		goto L143
	}
L81:
	;
	v263 = int32(1)
	if v258&int32(2072) != 0 {
		goto L80
	} else {
		goto L84
	}
L82:
	;
	goto L83
L83:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v28))) = uint8(base.B2i32(v29 != int32(1)))
	v483 = F_index_strategy_get_limit(m, v29)
	mBase = m.M
	v484 = m.ExcPending
	if v484 != 0 {
		goto L1
	} else {
		goto L125
	}
L84:
	;
	if v258&int32(96) == int32(0) {
		goto L76
	} else {
		goto L85
	}
L85:
	;
	v270 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v28))) = uint8(v270)
	if v256 == int32(0) {
		v914 = v263
		goto L78
	} else {
		goto L86
	}
L86:
	;
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
	v275 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v274)+16)))
	v277 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v274+v275)+12)))
	if v277&int32(1) == int32(0) {
		goto L79
	} else {
		goto L87
	}
L87:
	;
	v282 = F_trgm_presence_map(m, v256, v63)
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L1
	} else {
		goto L88
	}
L88:
	;
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v247)+12))
	v285 = int32(0)
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v284)))
	if v295 != 0 {
		goto L90
	} else {
		goto L91
	}
L89:
	;
	F_pfree(m, v282)
	mBase = m.M
	v479 = m.ExcPending
	if v479 != 0 {
		goto L1
	} else {
		goto L124
	}
L90:
	;
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v284)+16))
	base.MemoryFill(m, v296, int32(0), v295)
	goto L92
L91:
	;
	goto L92
L92:
	;
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v284)+8))
	if v299 != 0 {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v284)+20))
	base.MemoryFill(m, v300, int32(0), v299)
	goto L95
L94:
	;
	goto L95
L95:
	;
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v284)))
	if int32(0) < v303 {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v284)+4))
	v310 = v285
	v311 = v285
	goto L99
L97:
	;
	goto L98
L98:
	;
	v376 = *(*int32)(unsafe.Add(mBase, uint32(v284)+20))
	v377 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v376))) = uint8(v377)
	v379 = *(*int32)(unsafe.Add(mBase, uint32(v284)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v379))) = int32(0)
	v382 = *(*int32)(unsafe.Add(mBase, uint32(v284)+12))
	v389 = v377
	v394 = v285
	goto L111
L99:
	;
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v306+v310<<(uint(int32(2))%32))))
	v323 = v322 + v311
	if v322 <= int32(0) {
		goto L101
	} else {
		goto L102
	}
L100:
	;
	goto L98
L101:
	;
	v362 = v310 + int32(1)
	if v362 != v303 {
		v310 = v362
		v311 = v323
		goto L99
	} else {
		goto L109
	}
L102:
	;
	v328 = v311
	goto L103
L103:
	;
	v339 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v282+v328))))
	if v339 != int32(1) {
		goto L105
	} else {
		goto L106
	}
L104:
	;
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v284)+16))
	v347 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v345+v310))) = uint8(v347)
	goto L101
L105:
	;
	v343 = v328 + int32(1)
	if v343 < v323 {
		v328 = v343
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
	v399 = *(*int32)(unsafe.Add(mBase, uint32(v379+v394<<(uint(int32(2))%32))))
	v402 = v382 + v399<<(uint(int32(3))%32)
	v403 = *(*int32)(unsafe.Add(mBase, uint32(v402)))
	if int32(0) < v403 {
		goto L113
	} else {
		goto L114
	}
L112:
	;
	v477 = int32(0)
	goto L110
L113:
	;
	v406 = *(*int32)(unsafe.Add(mBase, uint32(v284)+16))
	v407 = *(*int32)(unsafe.Add(mBase, uint32(v402)+4))
	v411 = int32(0)
	v414 = v389
	goto L116
L114:
	;
	v454 = v389
	goto L115
L115:
	;
	v462 = v394 + int32(1)
	if v462 < v454 {
		v389 = v454
		v394 = v462
		goto L111
	} else {
		goto L123
	}
L116:
	;
	v423 = v407 + v411<<(uint(int32(3))%32)
	v424 = *(*int32)(unsafe.Add(mBase, uint32(v423)+4))
	v426 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v406+v424))))
	if v426 != int32(1) {
		v444 = v414
		goto L118
	} else {
		goto L119
	}
L117:
	;
	v454 = v444
	goto L115
L118:
	;
	v447 = v411 + int32(1)
	if v447 != v403 {
		v411 = v447
		v414 = v444
		goto L116
	} else {
		goto L122
	}
L119:
	;
	v429 = int32(1)
	v430 = *(*int32)(unsafe.Add(mBase, uint32(v423)))
	if v430 == v429 {
		v477 = v429
		goto L110
	} else {
		goto L120
	}
L120:
	;
	v433 = v376 + v430
	v434 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v433))))
	if v434 != 0 {
		v444 = v414
		goto L118
	} else {
		goto L121
	}
L121:
	;
	v435 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v433))) = uint8(v435)
	*(*int32)(unsafe.Add(mBase, uint32(v379+v414<<(uint(int32(2))%32)))) = v430
	v444 = v414 + v435
	goto L118
L122:
	;
	goto L117
L123:
	;
	goto L112
L124:
	;
	v914 = v477
	goto L78
L125:
	;
	v485 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
	v486 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v485)+16)))
	v488 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v485+v486)+12)))
	if v488&int32(1) != 0 {
		goto L126
	} else {
		goto L127
	}
L126:
	;
	v491 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28))))
	v492 = F_cnt_sml(m, v256, v63, v491)
	mBase = m.M
	v493 = m.ExcPending
	if v493 != 0 {
		goto L1
	} else {
		goto L129
	}
L127:
	;
	goto L128
L128:
	;
	v496 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63)+4)))
	if v496&int32(4) != 0 {
		goto L130
	} else {
		goto L131
	}
L129:
	;
	v914 = base.F64_le(v483, base.F64_promote_f32(v492))
	goto L78
L130:
	;
	v914 = int32(1)
	goto L78
L131:
	;
	goto L132
L132:
	;
	v500 = *(*int32)(unsafe.Add(mBase, uint32(v256)))
	v504 = int32(base.Ui32(v500)>>(uint(int32(2))%32)) - int32(5)
	if base.Ui32(v504) < base.Ui32(int32(3)) {
		goto L133
	} else {
		goto L134
	}
L133:
	;
	v914 = int32(0)
	goto L78
L134:
	;
	goto L135
L135:
	;
	v508 = int32(5)
	v512 = int32(1)
	v514 = base.I32_div_u_s(v504, int32(3))
	if base.Ui32(v514) <= base.Ui32(v512) {
		goto L136
	} else {
		goto L137
	}
L136:
	;
	v517 = v512
	goto L138
L137:
	;
	v517 = v514
	goto L138
L138:
	;
	v518 = int32(0)
	v520 = v518
	v521 = v518
	goto L139
L139:
	;
	v537 = int32(3)
	v539 = v256 + v508 + v520*v537
	v540 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v539))))
	v541 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v539)+2)))
	v545 = base.I32_rem_u_s(v540|v541<<(uint(int32(16))%32), v59)
	v549 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63+v508+int32(base.Ui32(v545)>>(uint(v537)%32))))))
	v553 = int32(1)
	v555 = v521 + int32(base.Ui32(v549)>>(uint(v545&int32(7))%32))&v553
	v557 = v520 + v553
	if v557 != v517 {
		v520 = v557
		v521 = v555
		goto L139
	} else {
		goto L141
	}
L140:
	;
	v914 = base.F64_ge(base.F64_div(base.F64_convert_i32_u(v555), base.F64_convert_i32_u(v514)), v483)
	goto L78
L141:
	;
	goto L140
L142:
	;
	v571 = F_trgm_contained_by(m, v256, v63)
	mBase = m.M
	v572 = m.ExcPending
	if v572 != 0 {
		goto L1
	} else {
		goto L145
	}
L143:
	;
	goto L144
L144:
	;
	v573 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63)+4)))
	if v573&int32(4) != 0 {
		v914 = v263
		goto L78
	} else {
		goto L146
	}
L145:
	;
	v914 = v571
	goto L78
L146:
	;
	v576 = *(*int32)(unsafe.Add(mBase, uint32(v256)))
	v580 = int32(base.Ui32(v576)>>(uint(int32(2))%32)) - int32(5)
	if base.Ui32(v580) < base.Ui32(int32(3)) {
		v914 = v263
		goto L78
	} else {
		goto L147
	}
L147:
	;
	v583 = int32(5)
	v587 = int32(1)
	v589 = base.I32_div_u_s(v580, int32(3))
	if base.Ui32(v589) <= base.Ui32(v587) {
		goto L148
	} else {
		goto L149
	}
L148:
	;
	v592 = v587
	goto L150
L149:
	;
	v592 = v589
	goto L150
L150:
	;
	v594 = int32(0)
	goto L151
L151:
	;
	v612 = int32(3)
	v614 = v256 + v583 + v594*v612
	v615 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v614))))
	v616 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v614)+2)))
	v620 = base.I32_rem_u_s(v615|v616<<(uint(int32(16))%32), v59)
	v627 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63+v583+int32(base.Ui32(v620)>>(uint(v612)%32))))))
	v628 = int32(1) << (uint(v620&int32(7)) % 32) & v627
	v629 = int32(0)
	v630 = base.B2i32(v628 != v629)
	if v628 == v629 {
		v914 = v630
		goto L78
	} else {
		goto L153
	}
L152:
	;
	v914 = v630
	goto L78
L153:
	;
	v634 = v594 + int32(1)
	if v634 != v592 {
		v594 = v634
		goto L151
	} else {
		goto L154
	}
L154:
	;
	goto L152
L155:
	;
	v639 = *(*int32)(unsafe.Add(mBase, uint32(v256)))
	v643 = int32(base.Ui32(v639)>>(uint(int32(2))%32)) - int32(5)
	v645 = base.I32_div_u_s(v643, int32(3))
	v646 = F_palloc(m, v645)
	mBase = m.M
	v647 = m.ExcPending
	if v647 != 0 {
		goto L1
	} else {
		goto L156
	}
L156:
	;
	if base.Ui32(int32(3)) <= base.Ui32(v643) {
		goto L157
	} else {
		goto L158
	}
L157:
	;
	v650 = int32(5)
	v654 = int32(1)
	if base.Ui32(v645) <= base.Ui32(v654) {
		goto L160
	} else {
		goto L161
	}
L158:
	;
	goto L159
L159:
	;
	v716 = *(*int32)(unsafe.Add(mBase, uint32(v247)+12))
	v717 = int32(0)
	v727 = *(*int32)(unsafe.Add(mBase, uint32(v716)))
	if v727 != 0 {
		goto L167
	} else {
		goto L168
	}
L160:
	;
	v657 = v654
	goto L162
L161:
	;
	v657 = v645
	goto L162
L162:
	;
	v659 = int32(0)
	goto L163
L163:
	;
	v677 = int32(3)
	v679 = v256 + v650 + v659*v677
	v680 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v679))))
	v681 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v679)+2)))
	v685 = base.I32_rem_u_s(v680|v681<<(uint(int32(16))%32), v59)
	v689 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63+v650+int32(base.Ui32(v685)>>(uint(v677)%32))))))
	v693 = int32(1)
	v694 = int32(base.Ui32(v689)>>(uint(v685&int32(7))%32)) & v693
	*(*uint8)(unsafe.Add(mBase, uint32(v659+v646))) = uint8(v694)
	v697 = v659 + v693
	if v697 != v657 {
		v659 = v697
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
	F_pfree(m, v646)
	mBase = m.M
	v911 = m.ExcPending
	if v911 != 0 {
		goto L1
	} else {
		goto L201
	}
L167:
	;
	v728 = *(*int32)(unsafe.Add(mBase, uint32(v716)+16))
	base.MemoryFill(m, v728, int32(0), v727)
	goto L169
L168:
	;
	goto L169
L169:
	;
	v731 = *(*int32)(unsafe.Add(mBase, uint32(v716)+8))
	if v731 != 0 {
		goto L170
	} else {
		goto L171
	}
L170:
	;
	v732 = *(*int32)(unsafe.Add(mBase, uint32(v716)+20))
	base.MemoryFill(m, v732, int32(0), v731)
	goto L172
L171:
	;
	goto L172
L172:
	;
	v735 = *(*int32)(unsafe.Add(mBase, uint32(v716)))
	if int32(0) < v735 {
		goto L173
	} else {
		goto L174
	}
L173:
	;
	v738 = *(*int32)(unsafe.Add(mBase, uint32(v716)+4))
	v742 = v717
	v743 = v717
	goto L176
L174:
	;
	goto L175
L175:
	;
	v808 = *(*int32)(unsafe.Add(mBase, uint32(v716)+20))
	v809 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v808))) = uint8(v809)
	v811 = *(*int32)(unsafe.Add(mBase, uint32(v716)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v811))) = int32(0)
	v814 = *(*int32)(unsafe.Add(mBase, uint32(v716)+12))
	v821 = v809
	v826 = v717
	goto L188
L176:
	;
	v754 = *(*int32)(unsafe.Add(mBase, uint32(v738+v742<<(uint(int32(2))%32))))
	v755 = v754 + v743
	if v754 <= int32(0) {
		goto L178
	} else {
		goto L179
	}
L177:
	;
	goto L175
L178:
	;
	v794 = v742 + int32(1)
	if v794 != v735 {
		v742 = v794
		v743 = v755
		goto L176
	} else {
		goto L186
	}
L179:
	;
	v760 = v743
	goto L180
L180:
	;
	v771 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v646+v760))))
	if v771 != int32(1) {
		goto L182
	} else {
		goto L183
	}
L181:
	;
	v777 = *(*int32)(unsafe.Add(mBase, uint32(v716)+16))
	v779 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v777+v742))) = uint8(v779)
	goto L178
L182:
	;
	v775 = v760 + int32(1)
	if v775 < v755 {
		v760 = v775
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
	v831 = *(*int32)(unsafe.Add(mBase, uint32(v811+v826<<(uint(int32(2))%32))))
	v834 = v814 + v831<<(uint(int32(3))%32)
	v835 = *(*int32)(unsafe.Add(mBase, uint32(v834)))
	if int32(0) < v835 {
		goto L190
	} else {
		goto L191
	}
L189:
	;
	v909 = int32(0)
	goto L187
L190:
	;
	v838 = *(*int32)(unsafe.Add(mBase, uint32(v716)+16))
	v839 = *(*int32)(unsafe.Add(mBase, uint32(v834)+4))
	v843 = int32(0)
	v846 = v821
	goto L193
L191:
	;
	v886 = v821
	goto L192
L192:
	;
	v894 = v826 + int32(1)
	if v894 < v886 {
		v821 = v886
		v826 = v894
		goto L188
	} else {
		goto L200
	}
L193:
	;
	v855 = v839 + v843<<(uint(int32(3))%32)
	v856 = *(*int32)(unsafe.Add(mBase, uint32(v855)+4))
	v858 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v838+v856))))
	if v858 != int32(1) {
		v876 = v846
		goto L195
	} else {
		goto L196
	}
L194:
	;
	v886 = v876
	goto L192
L195:
	;
	v879 = v843 + int32(1)
	if v879 != v835 {
		v843 = v879
		v846 = v876
		goto L193
	} else {
		goto L199
	}
L196:
	;
	v861 = int32(1)
	v862 = *(*int32)(unsafe.Add(mBase, uint32(v855)))
	if v862 == v861 {
		v909 = v861
		goto L187
	} else {
		goto L197
	}
L197:
	;
	v865 = v808 + v862
	v866 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v865))))
	if v866 != 0 {
		v876 = v846
		goto L195
	} else {
		goto L198
	}
L198:
	;
	v867 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v865))) = uint8(v867)
	*(*int32)(unsafe.Add(mBase, uint32(v811+v846<<(uint(int32(2))%32)))) = v862
	v876 = v846 + v867
	goto L195
L199:
	;
	goto L194
L200:
	;
	goto L189
L201:
	;
	v914 = v909
	goto L78
L202:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = v29
	F_errmsg_internal(m, int32(_a_F_gtrgm_consistent_0), v20+int32(16))
	mBase = m.M
	v945 = m.ExcPending
	if v945 != 0 {
		goto L1
	} else {
		goto L203
	}
L203:
	;
	F_errfinish(m, int32(_a_F_gtrgm_consistent_1), int32(443), int32(_a_F_gtrgm_consistent_2))
	mBase = m.M
	v950 = m.ExcPending
	if v950 != 0 {
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
func F_gtrgm_out(m *base.Module, l0 int32) int32 {
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v7 = Fn13852(m, l0, int32(_a_F_gtrgm_out_0), int32(71), int32(_a_F_gtrgm_out_1), int32(_a_F_gtrgm_out_2), int32(_a_F_gtrgm_out_3))
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
	}
}
