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
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
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
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
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
	v127 = F_palloc(m, int32(16))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
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
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+4)))
	if v82&int32(6) != int32(2) {
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
	v52 = int32(4)
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43))))
	if v54&int32(254) == int32(2) {
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
	v63 = v52
	goto L25
L24:
	;
	v63 = base.B2i32(v54 == int32(18)) << (uint(v52) % 32)
	goto L25
L25:
	;
	if v54 == int32(1) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v66 = v52
	goto L28
L27:
	;
	v66 = v63
	goto L28
L28:
	;
	v67 = F_generate_trgm(m, v49, v66)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L10
	} else {
		goto L29
	}
L29:
	;
	v121 = v67
	goto L12
L30:
	;
	v69 = int32(1)
	v73 = F_generate_trgm(m, v49, int32(base.Ui32(v46)>>(uint(v69)%32))-v69)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L10
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
	v80 = F_generate_trgm(m, v49, int32(base.Ui32(v75)>>(uint(int32(2))%32))-int32(4))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L10
	} else {
		goto L34
	}
L33:
	;
	v121 = v73
	goto L12
L34:
	;
	v121 = v80
	goto L12
L35:
	;
	return v7
L36:
	;
	goto L37
L37:
	;
	v88 = int32(0)
	if v35 <= v88 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v114 = F_palloc(m, int32(5))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L10
	} else {
		goto L46
	}
L39:
	;
	v93 = v88
	goto L40
L40:
	;
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v93+(v36+int32(5))))))
	if v100 == int32(255) {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	return v7
L42:
	;
	v104 = v93 + int32(1)
	if v35 != v104 {
		v93 = v104
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
	v116 = int32(6)
	*(*uint8)(unsafe.Add(mBase, uint32(v114)+4)) = uint8(v116)
	*(*int32)(unsafe.Add(mBase, uint32(v114))) = int32(20)
	v121 = v114
	goto L12
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v127))) = v121
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v127)+4)) = v130
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v127)+8)) = v132
	v134 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v7)+12)))
	v135 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v127)+14)) = uint8(v135)
	*(*uint16)(unsafe.Add(mBase, uint32(v127)+12)) = uint16(v134)
	return v127
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
	var v148 int32
	_ = v148
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v171 int32
	_ = v171
	var v177 int32
	_ = v177
	var v183 int32
	_ = v183
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
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
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
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v271 int32
	_ = v271
	var v280 int32
	_ = v280
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v316 int32
	_ = v316
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v338 int32
	_ = v338
	var v349 int32
	_ = v349
	var v353 int32
	_ = v353
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v372 int32
	_ = v372
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v389 int32
	_ = v389
	var v392 int32
	_ = v392
	var v398 int32
	_ = v398
	var v400 int32
	_ = v400
	var v409 int32
	_ = v409
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v421 int32
	_ = v421
	var v423 int32
	_ = v423
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v436 int32
	_ = v436
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v454 int32
	_ = v454
	var v457 int32
	_ = v457
	var v463 int32
	_ = v463
	var v472 int32
	_ = v472
	var v487 int32
	_ = v487
	var v489 int32
	_ = v489
	var v493 float64
	_ = v493
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v498 int32
	_ = v498
	var v501 int32
	_ = v501
	var v502 float32
	_ = v502
	var v503 int32
	_ = v503
	var v506 int32
	_ = v506
	var v510 int32
	_ = v510
	var v514 int32
	_ = v514
	var v518 int32
	_ = v518
	var v522 int32
	_ = v522
	var v524 int32
	_ = v524
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
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
	var v622 int32
	_ = v622
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v630 int32
	_ = v630
	var v637 int32
	_ = v637
	var v638 int32
	_ = v638
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v644 int32
	_ = v644
	var v646 int32
	_ = v646
	var v649 int32
	_ = v649
	var v653 int32
	_ = v653
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v657 int32
	_ = v657
	var v660 int32
	_ = v660
	var v664 int32
	_ = v664
	var v667 int32
	_ = v667
	var v669 int32
	_ = v669
	var v687 int32
	_ = v687
	var v689 int32
	_ = v689
	var v690 int32
	_ = v690
	var v691 int32
	_ = v691
	var v695 int32
	_ = v695
	var v699 int32
	_ = v699
	var v703 int32
	_ = v703
	var v704 int32
	_ = v704
	var v707 int32
	_ = v707
	var v726 int32
	_ = v726
	var v727 int32
	_ = v727
	var v737 int32
	_ = v737
	var v739 int32
	_ = v739
	var v740 int32
	_ = v740
	var v741 int32
	_ = v741
	var v743 int32
	_ = v743
	var v744 int32
	_ = v744
	var v745 int32
	_ = v745
	var v748 int32
	_ = v748
	var v752 int32
	_ = v752
	var v754 int32
	_ = v754
	var v764 int32
	_ = v764
	var v765 int32
	_ = v765
	var v770 int32
	_ = v770
	var v781 int32
	_ = v781
	var v785 int32
	_ = v785
	var v787 int32
	_ = v787
	var v789 int32
	_ = v789
	var v804 int32
	_ = v804
	var v818 int32
	_ = v818
	var v819 int32
	_ = v819
	var v821 int32
	_ = v821
	var v824 int32
	_ = v824
	var v830 int32
	_ = v830
	var v832 int32
	_ = v832
	var v841 int32
	_ = v841
	var v844 int32
	_ = v844
	var v845 int32
	_ = v845
	var v848 int32
	_ = v848
	var v849 int32
	_ = v849
	var v853 int32
	_ = v853
	var v855 int32
	_ = v855
	var v865 int32
	_ = v865
	var v866 int32
	_ = v866
	var v868 int32
	_ = v868
	var v871 int32
	_ = v871
	var v872 int32
	_ = v872
	var v875 int32
	_ = v875
	var v876 int32
	_ = v876
	var v877 int32
	_ = v877
	var v886 int32
	_ = v886
	var v889 int32
	_ = v889
	var v895 int32
	_ = v895
	var v904 int32
	_ = v904
	var v919 int32
	_ = v919
	var v921 int32
	_ = v921
	var v924 int32
	_ = v924
	var v949 int32
	_ = v949
	var v957 int32
	_ = v957
	var v964 int32
	_ = v964
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
		goto L78
	} else {
		goto L79
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
	v256 = v65
	goto L13
L37:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v221)+20))
	v226 = (v62 + int32(7)) & int32(2147483640)
	v230 = F_MemoryContextAlloc(m, v222, v226+v219+int32(16))
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
	v218 = v216
	v219 = v216
	v220 = v148
	goto L37
L39:
	;
	v218 = v211
	v219 = v212
	v220 = int32(0)
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
	v201 = F_generate_wildcard_trgm(m, v24+v197, v62-v197)
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
	v195 = F_generate_trgm(m, v24+v191, v62-v191)
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
	v177 = m.ExcPending
	if v177 != 0 {
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
	v148 = int32(1)
	if v148<<(uint(v29)%32)&int32(96) == int32(0) {
		goto L43
	} else {
		goto L47
	}
L47:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v64)+20))
	v159 = F_createTrgmNFA(m, v24, v155, v20+int32(28), v158)
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	if v159 == int32(0) {
		goto L38
	} else {
		goto L49
	}
L49:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v159)))
	v164 = int32(2)
	v165 = int32(base.Ui32(v163) >> (uint(v164) % 32))
	if base.Ui32(v164) < base.Ui32(v165-int32(5)) {
		v211 = v159
		v212 = v165
		goto L39
	} else {
		goto L50
	}
L50:
	;
	F_pfree(m, v159)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
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
	F_errmsg_internal(m, int32(473734), v20)
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	F_errfinish(m, int32(483861), int32(271), int32(91005))
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
	v218 = int32(0)
	v219 = int32(0)
	v220 = int32(1)
	goto L37
L58:
	;
	goto L59
L59:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v203)))
	v211 = v203
	v212 = int32(base.Ui32(v208) >> (uint(int32(2)) % 32))
	goto L39
L60:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v230))) = uint16(v29)
	v234 = v230 + int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v230)+4)) = v234
	if v62 != 0 {
		goto L62
	} else {
		goto L63
	}
L61:
	;
	if v220 == int32(0) {
		goto L66
	} else {
		goto L67
	}
L62:
	;
	v236 = F__emscripten_memcpy_bulkmem(m, v234, v24, v62)
	mBase = m.M
	v237 = v236
	goto L64
L63:
	;
	v237 = v234
	goto L64
L64:
	;
	goto L61
L65:
	;
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v20)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v230)+12)) = v249
	if v65 != 0 {
		goto L74
	} else {
		goto L75
	}
L66:
	;
	v240 = v237 + v226
	*(*int32)(unsafe.Add(mBase, uint32(v230)+8)) = v240
	if v219 != 0 {
		goto L70
	} else {
		goto L71
	}
L67:
	;
	goto L68
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v230)+8)) = int32(0)
	goto L65
L69:
	;
	F_pfree(m, v218)
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L1
	} else {
		goto L73
	}
L70:
	;
	v242 = F__emscripten_memcpy_bulkmem(m, v240, v218, v219)
	mBase = m.M
	goto L72
L71:
	;
	goto L72
L72:
	;
	goto L69
L73:
	;
	goto L65
L74:
	;
	F_pfree(m, v65)
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L1
	} else {
		goto L77
	}
L75:
	;
	goto L76
L76:
	;
	v253 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v253)+16)) = v230
	v256 = v230
	goto L13
L77:
	;
	goto L76
L78:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v949 = m.ExcPending
	if v949 != 0 {
		goto L1
	} else {
		goto L192
	}
L79:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v256)+8))
	v266 = int32(1) << (uint(v29) % 32)
	if v266&int32(642) == int32(0) {
		goto L83
	} else {
		goto L84
	}
L80:
	;
	m.G0 = v20 + int32(32)
	return v924
L81:
	;
	v646 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63)+4)))
	if v646&int32(4) != 0 {
		v924 = v271
		goto L80
	} else {
		goto L151
	}
L82:
	;
	v573 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v28))) = uint8(v573)
	v575 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
	v576 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v575)+16)))
	v578 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v575+v576)+12)))
	if v578&v573 != 0 {
		goto L138
	} else {
		goto L139
	}
L83:
	;
	v271 = int32(1)
	if v266&int32(2072) != 0 {
		goto L82
	} else {
		goto L86
	}
L84:
	;
	goto L85
L85:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v28))) = uint8(base.B2i32(v29 != int32(1)))
	v493 = F_index_strategy_get_limit(m, v29)
	mBase = m.M
	v494 = m.ExcPending
	if v494 != 0 {
		goto L1
	} else {
		goto L121
	}
L86:
	;
	if int32(1)<<(uint(v29)%32)&int32(96) == int32(0) {
		goto L78
	} else {
		goto L87
	}
L87:
	;
	v280 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v28))) = uint8(v280)
	if v264 == int32(0) {
		v924 = v271
		goto L80
	} else {
		goto L88
	}
L88:
	;
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
	v285 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v284)+16)))
	v287 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v284+v285)+12)))
	if v287&int32(1) == int32(0) {
		goto L81
	} else {
		goto L89
	}
L89:
	;
	v292 = F_trgm_presence_map(m, v264, v63)
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L1
	} else {
		goto L90
	}
L90:
	;
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v256)+12))
	v295 = int32(0)
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v294)+16))
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v294)))
	v308 = F___memset(m, v305, v295, v307)
	mBase = m.M
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v294)+20))
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v294)+8))
	v312 = F___memset(m, v309, v295, v311)
	mBase = m.M
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v294)))
	if v295 < v313 {
		goto L92
	} else {
		goto L93
	}
L91:
	;
	F_pfree(m, v292)
	mBase = m.M
	v489 = m.ExcPending
	if v489 != 0 {
		goto L1
	} else {
		goto L120
	}
L92:
	;
	v316 = *(*int32)(unsafe.Add(mBase, uint32(v294)+4))
	v320 = v295
	v322 = v295
	goto L95
L93:
	;
	goto L94
L94:
	;
	v386 = *(*int32)(unsafe.Add(mBase, uint32(v294)+20))
	v387 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v386))) = uint8(v387)
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v294)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v389))) = int32(0)
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v294)+12))
	v398 = v387
	v400 = v295
	goto L107
L95:
	;
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v316+v322<<(uint(int32(2))%32))))
	v333 = v332 + v320
	if v332 <= int32(0) {
		goto L97
	} else {
		goto L98
	}
L96:
	;
	goto L94
L97:
	;
	v372 = v322 + int32(1)
	if v372 != v313 {
		v320 = v333
		v322 = v372
		goto L95
	} else {
		goto L105
	}
L98:
	;
	v338 = v320
	goto L99
L99:
	;
	v349 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v292+v338))))
	if v349 != int32(1) {
		goto L101
	} else {
		goto L102
	}
L100:
	;
	v355 = *(*int32)(unsafe.Add(mBase, uint32(v294)+16))
	v357 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v355+v322))) = uint8(v357)
	goto L97
L101:
	;
	v353 = v338 + int32(1)
	if v353 < v333 {
		v338 = v353
		goto L99
	} else {
		goto L104
	}
L102:
	;
	goto L103
L103:
	;
	goto L100
L104:
	;
	goto L97
L105:
	;
	goto L96
L106:
	;
	goto L91
L107:
	;
	v409 = *(*int32)(unsafe.Add(mBase, uint32(v389+v400<<(uint(int32(2))%32))))
	v412 = v392 + v409<<(uint(int32(3))%32)
	v413 = *(*int32)(unsafe.Add(mBase, uint32(v412)))
	if int32(0) < v413 {
		goto L109
	} else {
		goto L110
	}
L108:
	;
	v487 = int32(0)
	goto L106
L109:
	;
	v416 = *(*int32)(unsafe.Add(mBase, uint32(v294)+16))
	v417 = *(*int32)(unsafe.Add(mBase, uint32(v412)+4))
	v421 = int32(0)
	v423 = v398
	goto L112
L110:
	;
	v463 = v398
	goto L111
L111:
	;
	v472 = v400 + int32(1)
	if v472 < v463 {
		v398 = v463
		v400 = v472
		goto L107
	} else {
		goto L119
	}
L112:
	;
	v433 = v417 + v421<<(uint(int32(3))%32)
	v434 = *(*int32)(unsafe.Add(mBase, uint32(v433)+4))
	v436 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v416+v434))))
	if v436 != int32(1) {
		v454 = v423
		goto L114
	} else {
		goto L115
	}
L113:
	;
	v463 = v454
	goto L111
L114:
	;
	v457 = v421 + int32(1)
	if v457 != v413 {
		v421 = v457
		v423 = v454
		goto L112
	} else {
		goto L118
	}
L115:
	;
	v439 = int32(1)
	v440 = *(*int32)(unsafe.Add(mBase, uint32(v433)))
	if v440 == v439 {
		v487 = v439
		goto L106
	} else {
		goto L116
	}
L116:
	;
	v443 = v440 + v386
	v444 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v443))))
	if v444 != 0 {
		v454 = v423
		goto L114
	} else {
		goto L117
	}
L117:
	;
	v445 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v443))) = uint8(v445)
	*(*int32)(unsafe.Add(mBase, uint32(v389+v423<<(uint(int32(2))%32)))) = v440
	v454 = v423 + v445
	goto L114
L118:
	;
	goto L113
L119:
	;
	goto L108
L120:
	;
	v924 = v487
	goto L80
L121:
	;
	v495 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
	v496 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v495)+16)))
	v498 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v495+v496)+12)))
	if v498&int32(1) != 0 {
		goto L122
	} else {
		goto L123
	}
L122:
	;
	v501 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28))))
	v502 = F_cnt_sml(m, v264, v63, v501)
	mBase = m.M
	v503 = m.ExcPending
	if v503 != 0 {
		goto L1
	} else {
		goto L125
	}
L123:
	;
	goto L124
L124:
	;
	v506 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63)+4)))
	if v506&int32(4) != 0 {
		goto L126
	} else {
		goto L127
	}
L125:
	;
	v924 = base.F64_le(v493, base.F64_promote_f32(v502))
	goto L80
L126:
	;
	v924 = int32(1)
	goto L80
L127:
	;
	goto L128
L128:
	;
	v510 = *(*int32)(unsafe.Add(mBase, uint32(v264)))
	v514 = int32(base.Ui32(v510)>>(uint(int32(2))%32)) - int32(5)
	if base.Ui32(v514) < base.Ui32(int32(3)) {
		goto L129
	} else {
		goto L130
	}
L129:
	;
	v924 = int32(0)
	goto L80
L130:
	;
	goto L131
L131:
	;
	v518 = int32(5)
	v522 = int32(1)
	v524 = base.I32_div_u_s(v514, int32(3))
	if base.Ui32(v524) <= base.Ui32(v522) {
		goto L132
	} else {
		goto L133
	}
L132:
	;
	v527 = v522
	goto L134
L133:
	;
	v527 = v524
	goto L134
L134:
	;
	v528 = int32(0)
	v530 = v528
	v531 = v528
	goto L135
L135:
	;
	v547 = int32(3)
	v549 = v264 + v518 + v530*v547
	v550 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v549))))
	v551 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v549)+2)))
	v555 = base.I32_rem_u_s(v550|v551<<(uint(int32(16))%32), v59)
	v559 = int32(*(*int8)(unsafe.Add(mBase, uint32(v63+v518+int32(base.Ui32(v555)>>(uint(v547)%32))))))
	v563 = int32(1)
	v565 = int32(base.Ui32(v559)>>(uint(v555&int32(7))%32))&v563 + v531
	v567 = v530 + v563
	if v567 != v527 {
		v530 = v567
		v531 = v565
		goto L135
	} else {
		goto L137
	}
L136:
	;
	v924 = base.F64_ge(base.F64_div(base.F64_convert_i32_u(v565), base.F64_convert_i32_u(v524)), v493)
	goto L80
L137:
	;
	goto L136
L138:
	;
	v581 = F_trgm_contained_by(m, v264, v63)
	mBase = m.M
	v582 = m.ExcPending
	if v582 != 0 {
		goto L1
	} else {
		goto L141
	}
L139:
	;
	goto L140
L140:
	;
	v583 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63)+4)))
	if v583&int32(4) != 0 {
		v924 = v271
		goto L80
	} else {
		goto L142
	}
L141:
	;
	v924 = v581
	goto L80
L142:
	;
	v586 = *(*int32)(unsafe.Add(mBase, uint32(v264)))
	v590 = int32(base.Ui32(v586)>>(uint(int32(2))%32)) - int32(5)
	if base.Ui32(v590) < base.Ui32(int32(3)) {
		v924 = v271
		goto L80
	} else {
		goto L143
	}
L143:
	;
	v593 = int32(5)
	v597 = int32(1)
	v599 = base.I32_div_u_s(v590, int32(3))
	if base.Ui32(v599) <= base.Ui32(v597) {
		goto L144
	} else {
		goto L145
	}
L144:
	;
	v602 = v597
	goto L146
L145:
	;
	v602 = v599
	goto L146
L146:
	;
	v604 = int32(0)
	goto L147
L147:
	;
	v622 = int32(3)
	v624 = v264 + v593 + v604*v622
	v625 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v624))))
	v626 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v624)+2)))
	v630 = base.I32_rem_u_s(v625|v626<<(uint(int32(16))%32), v59)
	v637 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63+v593+int32(base.Ui32(v630)>>(uint(v622)%32))))))
	v638 = int32(1) << (uint(v630&int32(7)) % 32) & v637
	v639 = int32(0)
	v640 = base.B2i32(v638 != v639)
	if v638 == v639 {
		v924 = v640
		goto L80
	} else {
		goto L149
	}
L148:
	;
	v924 = v640
	goto L80
L149:
	;
	v644 = v604 + int32(1)
	if v644 != v602 {
		v604 = v644
		goto L147
	} else {
		goto L150
	}
L150:
	;
	goto L148
L151:
	;
	v649 = *(*int32)(unsafe.Add(mBase, uint32(v264)))
	v653 = int32(base.Ui32(v649)>>(uint(int32(2))%32)) - int32(5)
	v655 = base.I32_div_u_s(v653, int32(3))
	v656 = F_palloc(m, v655)
	mBase = m.M
	v657 = m.ExcPending
	if v657 != 0 {
		goto L1
	} else {
		goto L152
	}
L152:
	;
	if base.Ui32(int32(3)) <= base.Ui32(v653) {
		goto L153
	} else {
		goto L154
	}
L153:
	;
	v660 = int32(5)
	v664 = int32(1)
	if base.Ui32(v655) <= base.Ui32(v664) {
		goto L156
	} else {
		goto L157
	}
L154:
	;
	goto L155
L155:
	;
	v726 = *(*int32)(unsafe.Add(mBase, uint32(v256)+12))
	v727 = int32(0)
	v737 = *(*int32)(unsafe.Add(mBase, uint32(v726)+16))
	v739 = *(*int32)(unsafe.Add(mBase, uint32(v726)))
	v740 = F___memset(m, v737, v727, v739)
	mBase = m.M
	v741 = *(*int32)(unsafe.Add(mBase, uint32(v726)+20))
	v743 = *(*int32)(unsafe.Add(mBase, uint32(v726)+8))
	v744 = F___memset(m, v741, v727, v743)
	mBase = m.M
	v745 = *(*int32)(unsafe.Add(mBase, uint32(v726)))
	if v727 < v745 {
		goto L163
	} else {
		goto L164
	}
L156:
	;
	v667 = v664
	goto L158
L157:
	;
	v667 = v655
	goto L158
L158:
	;
	v669 = int32(0)
	goto L159
L159:
	;
	v687 = int32(3)
	v689 = v264 + v660 + v669*v687
	v690 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v689))))
	v691 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v689)+2)))
	v695 = base.I32_rem_u_s(v690|v691<<(uint(int32(16))%32), v59)
	v699 = int32(*(*int8)(unsafe.Add(mBase, uint32(v63+v660+int32(base.Ui32(v695)>>(uint(v687)%32))))))
	v703 = int32(1)
	v704 = int32(base.Ui32(v699)>>(uint(v695&int32(7))%32)) & v703
	*(*uint8)(unsafe.Add(mBase, uint32(v669+v656))) = uint8(v704)
	v707 = v669 + v703
	if v707 != v667 {
		v669 = v707
		goto L159
	} else {
		goto L161
	}
L160:
	;
	goto L155
L161:
	;
	goto L160
L162:
	;
	F_pfree(m, v656)
	mBase = m.M
	v921 = m.ExcPending
	if v921 != 0 {
		goto L1
	} else {
		goto L191
	}
L163:
	;
	v748 = *(*int32)(unsafe.Add(mBase, uint32(v726)+4))
	v752 = v727
	v754 = v727
	goto L166
L164:
	;
	goto L165
L165:
	;
	v818 = *(*int32)(unsafe.Add(mBase, uint32(v726)+20))
	v819 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v818))) = uint8(v819)
	v821 = *(*int32)(unsafe.Add(mBase, uint32(v726)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v821))) = int32(0)
	v824 = *(*int32)(unsafe.Add(mBase, uint32(v726)+12))
	v830 = v819
	v832 = v727
	goto L178
L166:
	;
	v764 = *(*int32)(unsafe.Add(mBase, uint32(v748+v754<<(uint(int32(2))%32))))
	v765 = v764 + v752
	if v764 <= int32(0) {
		goto L168
	} else {
		goto L169
	}
L167:
	;
	goto L165
L168:
	;
	v804 = v754 + int32(1)
	if v804 != v745 {
		v752 = v765
		v754 = v804
		goto L166
	} else {
		goto L176
	}
L169:
	;
	v770 = v752
	goto L170
L170:
	;
	v781 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v656+v770))))
	if v781 != int32(1) {
		goto L172
	} else {
		goto L173
	}
L171:
	;
	v787 = *(*int32)(unsafe.Add(mBase, uint32(v726)+16))
	v789 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v787+v754))) = uint8(v789)
	goto L168
L172:
	;
	v785 = v770 + int32(1)
	if v785 < v765 {
		v770 = v785
		goto L170
	} else {
		goto L175
	}
L173:
	;
	goto L174
L174:
	;
	goto L171
L175:
	;
	goto L168
L176:
	;
	goto L167
L177:
	;
	goto L162
L178:
	;
	v841 = *(*int32)(unsafe.Add(mBase, uint32(v821+v832<<(uint(int32(2))%32))))
	v844 = v824 + v841<<(uint(int32(3))%32)
	v845 = *(*int32)(unsafe.Add(mBase, uint32(v844)))
	if int32(0) < v845 {
		goto L180
	} else {
		goto L181
	}
L179:
	;
	v919 = int32(0)
	goto L177
L180:
	;
	v848 = *(*int32)(unsafe.Add(mBase, uint32(v726)+16))
	v849 = *(*int32)(unsafe.Add(mBase, uint32(v844)+4))
	v853 = int32(0)
	v855 = v830
	goto L183
L181:
	;
	v895 = v830
	goto L182
L182:
	;
	v904 = v832 + int32(1)
	if v904 < v895 {
		v830 = v895
		v832 = v904
		goto L178
	} else {
		goto L190
	}
L183:
	;
	v865 = v849 + v853<<(uint(int32(3))%32)
	v866 = *(*int32)(unsafe.Add(mBase, uint32(v865)+4))
	v868 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v848+v866))))
	if v868 != int32(1) {
		v886 = v855
		goto L185
	} else {
		goto L186
	}
L184:
	;
	v895 = v886
	goto L182
L185:
	;
	v889 = v853 + int32(1)
	if v889 != v845 {
		v853 = v889
		v855 = v886
		goto L183
	} else {
		goto L189
	}
L186:
	;
	v871 = int32(1)
	v872 = *(*int32)(unsafe.Add(mBase, uint32(v865)))
	if v872 == v871 {
		v919 = v871
		goto L177
	} else {
		goto L187
	}
L187:
	;
	v875 = v872 + v818
	v876 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v875))))
	if v876 != 0 {
		v886 = v855
		goto L185
	} else {
		goto L188
	}
L188:
	;
	v877 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v875))) = uint8(v877)
	*(*int32)(unsafe.Add(mBase, uint32(v821+v855<<(uint(int32(2))%32)))) = v872
	v886 = v855 + v877
	goto L185
L189:
	;
	goto L184
L190:
	;
	goto L179
L191:
	;
	v924 = v919
	goto L80
L192:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = v29
	F_errmsg_internal(m, int32(473734), v20+int32(16))
	mBase = m.M
	v957 = m.ExcPending
	if v957 != 0 {
		goto L1
	} else {
		goto L193
	}
L193:
	;
	F_errfinish(m, int32(483861), int32(443), int32(91005))
	mBase = m.M
	v964 = m.ExcPending
	if v964 != 0 {
		goto L1
	} else {
		goto L194
	}
L194:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_gtrgm_out(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v23 int32
	_ = v23
	var v30 int32
	_ = v30
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		F_errcode(m, int32(1088))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v5))) = int32(284092)
			F_errmsg(m, int32(189401), v5)
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(483861), int32(71), int32(66191))
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
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
