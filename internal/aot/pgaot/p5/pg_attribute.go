package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_pg_attribute_aclcheck(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64) int32 {
	var v6 int64
	_ = v6
	var v9 int32
	_ = v9
	v6 = F_pg_attribute_aclmask_ext(m, l0, l1, l2, l3, int32(0))
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v6 == int64(0))
	}
}
func F_pg_attribute_aclcheck_all_ext(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int64
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v121 int32
	_ = v121
	v13 = m.G0
	v15 = v13 - int32(16)
	m.G0 = v15
	v18 = F_SearchSysCache1(m, int32(57), l0)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v15 + int32(16)
	return v121
L2:
	;
	return int32(0)
L3:
	;
	if v18 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	if l4 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v18)+16))
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+22)))
	v45 = v43 + v44
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)+80))
	v47 = int32(*(*int16)(unsafe.Add(mBase, uint32(v45)+120)))
	F_ReleaseCatCache(m, v18)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L2
	} else {
		goto L14
	}
L7:
	;
	v24 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v24)
	v121 = v24
	goto L1
L8:
	;
	goto L9
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L2
	} else {
		goto L10
	}
L10:
	;
	F_errcode(m, int32(16908420))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L2
	} else {
		goto L11
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = l0
	F_errmsg(m, int32(_a_F_pg_attribute_aclcheck_all_ext_0), v15)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L2
	} else {
		goto L12
	}
L12:
	;
	F_errfinish(m, int32(_a_F_pg_attribute_aclcheck_all_ext_1), int32(3949), int32(_a_F_pg_attribute_aclcheck_all_ext_2))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L2
	} else {
		goto L13
	}
L13:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L14:
	;
	v50 = int32(1)
	if v47 <= int32(0) {
		v121 = v50
		goto L1
	} else {
		goto L15
	}
L15:
	;
	v59 = v50
	v62 = int32(1)
	goto L16
L16:
	;
	v68 = F_SearchSysCache2(m, int32(7), l0, base.I32_extend16_s(v62))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L2
	} else {
		goto L19
	}
L17:
	;
	v121 = v110
	goto L1
L18:
	;
	v114 = base.I32_extend16_s(v62 + int32(1))
	if v114 <= v47 {
		v59 = v110
		v62 = v114
		goto L16
	} else {
		goto L41
	}
L19:
	;
	if v68 == int32(0) {
		v110 = v59
		goto L18
	} else {
		goto L20
	}
L20:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v68)+16))
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72)+22)))
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72+v73)+91)))
	if v75 == int32(1) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	F_ReleaseCatCache(m, v68)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L2
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v84 = F_SysCacheGetAttr(m, int32(7), v68, int32(22), v15+int32(15))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L2
	} else {
		goto L25
	}
L24:
	;
	v110 = v59
	goto L18
L25:
	;
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+15)))
	if v86 == int32(1) {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v107 = int32(1)
	if l3 == int32(0) {
		v121 = v107
		goto L1
	} else {
		goto L40
	}
L27:
	;
	F_ReleaseCatCache(m, v68)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L2
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v91 = F_pg_detoast_datum(m, v84)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L2
	} else {
		goto L31
	}
L30:
	;
	goto L26
L31:
	;
	v94 = F_aclmask(m, v91, l1, v46, l2, int32(1))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L2
	} else {
		goto L32
	}
L32:
	;
	if v84 != v91 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	F_pfree(m, v91)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L2
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	F_ReleaseCatCache(m, v68)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L2
	} else {
		goto L37
	}
L36:
	;
	goto L35
L37:
	;
	if v94 == int64(0) {
		goto L26
	} else {
		goto L38
	}
L38:
	;
	v103 = int32(0)
	if l3 == int32(1) {
		v121 = v103
		goto L1
	} else {
		goto L39
	}
L39:
	;
	v110 = v103
	goto L18
L40:
	;
	v110 = v107
	goto L18
L41:
	;
	goto L17
}
func F_pg_restore_attribute_stats(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v40 int64
	_ = v40
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
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v70 int64
	_ = v70
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v153 int32
	_ = v153
	var v158 int32
	_ = v158
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v183 int32
	_ = v183
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v202 int32
	_ = v202
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
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
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v361 int32
	_ = v361
	var v363 int32
	_ = v363
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v407 int32
	_ = v407
	var v410 int32
	_ = v410
	var v443 int32
	_ = v443
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v490 int32
	_ = v490
	var v495 int32
	_ = v495
	var v527 int32
	_ = v527
	var v529 int32
	_ = v529
	var v532 int32
	_ = v532
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v542 int32
	_ = v542
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v558 int32
	_ = v558
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v579 int32
	_ = v579
	var v581 int32
	_ = v581
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v591 int32
	_ = v591
	var v600 int32
	_ = v600
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v618 int32
	_ = v618
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v624 int32
	_ = v624
	var v629 int32
	_ = v629
	var v632 int32
	_ = v632
	var v633 int32
	_ = v633
	var v636 int32
	_ = v636
	var v642 int32
	_ = v642
	var v651 int32
	_ = v651
	var v656 int32
	_ = v656
	var v657 int32
	_ = v657
	var v659 int32
	_ = v659
	var v660 int32
	_ = v660
	var v661 int32
	_ = v661
	var v662 int32
	_ = v662
	var v663 int32
	_ = v663
	var v669 int32
	_ = v669
	var v672 int32
	_ = v672
	var v673 int32
	_ = v673
	var v676 int32
	_ = v676
	var v682 int32
	_ = v682
	var v691 int32
	_ = v691
	var v696 int32
	_ = v696
	var v697 int32
	_ = v697
	var v699 int32
	_ = v699
	var v700 int32
	_ = v700
	var v701 int32
	_ = v701
	var v706 int32
	_ = v706
	var v709 int32
	_ = v709
	var v710 int32
	_ = v710
	var v713 int32
	_ = v713
	var v714 int32
	_ = v714
	var v715 int32
	_ = v715
	var v716 int32
	_ = v716
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v725 int64
	_ = v725
	var v736 int32
	_ = v736
	var v744 int64
	_ = v744
	var v782 int32
	_ = v782
	var v787 int32
	_ = v787
	var v789 int32
	_ = v789
	var v791 int32
	_ = v791
	var v796 int32
	_ = v796
	var v798 int32
	_ = v798
	var v800 int32
	_ = v800
	var v805 int32
	_ = v805
	var v807 int32
	_ = v807
	var v809 int32
	_ = v809
	var v814 int32
	_ = v814
	var v822 int32
	_ = v822
	var v826 int32
	_ = v826
	var v829 int32
	_ = v829
	var v830 int32
	_ = v830
	var v832 int32
	_ = v832
	var v842 int32
	_ = v842
	var v845 int32
	_ = v845
	var v850 int32
	_ = v850
	var v856 int32
	_ = v856
	var v859 int32
	_ = v859
	var v860 int32
	_ = v860
	var v862 int32
	_ = v862
	var v872 int32
	_ = v872
	var v876 int32
	_ = v876
	var v879 int32
	_ = v879
	var v882 int32
	_ = v882
	var v895 int32
	_ = v895
	var v896 int32
	_ = v896
	var v897 int32
	_ = v897
	var v901 int32
	_ = v901
	var v904 int32
	_ = v904
	var v908 int32
	_ = v908
	var v911 int32
	_ = v911
	var v912 int32
	_ = v912
	var v914 int32
	_ = v914
	var v924 int32
	_ = v924
	var v927 int32
	_ = v927
	var v932 int32
	_ = v932
	var v942 int32
	_ = v942
	var v943 int32
	_ = v943
	var v947 int32
	_ = v947
	var v953 int32
	_ = v953
	var v956 int32
	_ = v956
	var v957 int32
	_ = v957
	var v959 int32
	_ = v959
	var v969 int32
	_ = v969
	var v975 int32
	_ = v975
	var v978 int32
	_ = v978
	var v979 int32
	_ = v979
	var v985 int32
	_ = v985
	var v986 int32
	_ = v986
	var v987 int32
	_ = v987
	var v991 int32
	_ = v991
	var v996 int32
	_ = v996
	var v997 int32
	_ = v997
	var v998 int32
	_ = v998
	var v1009 int32
	_ = v1009
	var v1013 int32
	_ = v1013
	var v1018 int32
	_ = v1018
	var v1020 int32
	_ = v1020
	var v1027 int32
	_ = v1027
	var v1028 int32
	_ = v1028
	var v1032 int32
	_ = v1032
	var v1037 int32
	_ = v1037
	var v1038 int32
	_ = v1038
	var v1040 int32
	_ = v1040
	var v1041 int32
	_ = v1041
	var v1043 int32
	_ = v1043
	var v1045 int32
	_ = v1045
	var v1047 int32
	_ = v1047
	var v1050 int32
	_ = v1050
	var v1059 int32
	_ = v1059
	var v1062 int32
	_ = v1062
	var v1066 int32
	_ = v1066
	var v1070 int32
	_ = v1070
	var v1075 int32
	_ = v1075
	var v1079 int32
	_ = v1079
	var v1082 int32
	_ = v1082
	var v1091 int32
	_ = v1091
	var v1096 int32
	_ = v1096
	var v1100 int32
	_ = v1100
	var v1103 int32
	_ = v1103
	var v1109 int32
	_ = v1109
	var v1114 int32
	_ = v1114
	var v1118 int32
	_ = v1118
	var v1121 int32
	_ = v1121
	var v1122 int32
	_ = v1122
	var v1131 int32
	_ = v1131
	var v1136 int32
	_ = v1136
	var v1140 int32
	_ = v1140
	var v1143 int32
	_ = v1143
	var v1144 int32
	_ = v1144
	var v1153 int32
	_ = v1153
	var v1158 int32
	_ = v1158
	var v1162 int32
	_ = v1162
	var v1166 int32
	_ = v1166
	var v1171 int32
	_ = v1171
	v2 = int32(0)
	v34 = m.G0
	v36 = v34 - int32(608)
	m.G0 = v36
	*(*uint8)(unsafe.Add(mBase, uint32(v36)+224)) = uint8(v2)
	v40 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v36)+216)) = v40
	*(*int64)(unsafe.Add(mBase, uint32(v36)+208)) = v40
	v44 = int32(18)
	*(*uint16)(unsafe.Add(mBase, uint32(v36)+226)) = uint16(v44)
	v47 = v36 + int32(208)
	v49 = F_stats_fill_fcinfo_from_arg_pairs(m, l0, v47, int32(_a_F_pg_restore_attribute_stats_0))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v53 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v36)+604)) = v53
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+304)))
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+296)))
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+328)))
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+336)))
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+352)))
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+360)))
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+312)))
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+320)))
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+344)))
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+368)))
	base.MemoryFill(m, v36+int32(448), v53, int32(124))
	v70 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v36)+439)) = v70
	*(*int64)(unsafe.Add(mBase, uint32(v36)+432)) = v70
	*(*int64)(unsafe.Add(mBase, uint32(v36)+424)) = v70
	*(*int64)(unsafe.Add(mBase, uint32(v36)+416)) = v70
	*(*int64)(unsafe.Add(mBase, uint32(v36)+407)) = v70
	*(*int64)(unsafe.Add(mBase, uint32(v36)+400)) = v70
	*(*int64)(unsafe.Add(mBase, uint32(v36)+392)) = v70
	*(*int64)(unsafe.Add(mBase, uint32(v36)+384)) = v70
	F_stats_check_required_arg(m, v47, int32(_a_F_pg_restore_attribute_stats_0), v53)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	F_stats_check_required_arg(m, v47, int32(_a_F_pg_restore_attribute_stats_0), int32(1))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v36)+228))
	v95 = F_text_to_cstring(m, v94)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v36)+236))
	v98 = F_text_to_cstring(m, v97)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v102 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_restore_attribute_stats[0])))
	if v102 == int32(1) {
		goto L13
	} else {
		goto L14
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1162 = m.ExcPending
	if v1162 != 0 {
		goto L1
	} else {
		goto L253
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1140 = m.ExcPending
	if v1140 != 0 {
		goto L1
	} else {
		goto L249
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1118 = m.ExcPending
	if v1118 != 0 {
		goto L1
	} else {
		goto L245
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1100 = m.ExcPending
	if v1100 != 0 {
		goto L1
	} else {
		goto L241
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1079 = m.ExcPending
	if v1079 != 0 {
		goto L1
	} else {
		goto L237
	}
L12:
	;
	if v112 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L13:
	;
	v107 = *(*int32)(unsafe.Add(mBase, _c_F_pg_restore_attribute_stats[1]))
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v107)+316))
	v110 = base.B2i32(v108 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _c_F_pg_restore_attribute_stats[0])) = uint8(v110)
	v112 = v110
	goto L15
L14:
	;
	v112 = int32(0)
	goto L15
L15:
	;
	goto L12
L16:
	;
	v116 = F_makeRangeVar(m, v95, v98, int32(-1))
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L1
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1059 = m.ExcPending
	if v1059 != 0 {
		goto L1
	} else {
		goto L232
	}
L19:
	;
	v123 = F_RangeVarGetRelidExtended(m, v116, int32(4), int32(0), int32(1060), v36+int32(604))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+256)))
	v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+248)))
	if v126&int32(1) == int32(0) {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v211 = base.I32_extend16_s(v208)
	if v211 < int32(0) {
		goto L10
	} else {
		goto L50
	}
L22:
	;
	if v125&int32(1) == int32(0) {
		goto L11
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	if v125&int32(1) == int32(0) {
		goto L33
	} else {
		goto L34
	}
L25:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v36)+244))
	v136 = F_text_to_cstring(m, v135)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	v138 = F_get_attnum(m, v123, v136)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	if v138 != 0 {
		v208 = v138
		v210 = v136
		goto L21
	} else {
		goto L28
	}
L28:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	F_errcode(m, int32(50360452))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+180)) = v98
	*(*int32)(unsafe.Add(mBase, uint32(v36)+176)) = v136
	F_errmsg(m, int32(_a_F_pg_restore_attribute_stats_18), v36+int32(176))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	F_errfinish(m, int32(_a_F_pg_restore_attribute_stats_4), int32(213), int32(_a_F_pg_restore_attribute_stats_5))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L33:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v36)+252))
	v164 = base.I32_extend16_s(v163)
	v166 = F_get_attname(m, v123, v164, int32(1))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L1
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L1
	} else {
		goto L46
	}
L36:
	;
	if v166 != 0 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v168 = F_SearchSysCacheExistsAttName(m, v123, v166)
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L1
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L1
	} else {
		goto L42
	}
L40:
	;
	if v168 != 0 {
		v208 = v163
		v210 = v166
		goto L21
	} else {
		goto L41
	}
L41:
	;
	goto L39
L42:
	;
	F_errcode(m, int32(50360452))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+20)) = v98
	*(*int32)(unsafe.Add(mBase, uint32(v36)+16)) = v164
	F_errmsg(m, int32(_a_F_pg_restore_attribute_stats_7), v36+int32(16))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	F_errfinish(m, int32(_a_F_pg_restore_attribute_stats_4), int32(225), int32(_a_F_pg_restore_attribute_stats_5))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L46:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+4)) = int32(_a_F_pg_restore_attribute_stats_1)
	*(*int32)(unsafe.Add(mBase, uint32(v36))) = int32(_a_F_pg_restore_attribute_stats_2)
	F_errmsg(m, int32(_a_F_pg_restore_attribute_stats_19), v36)
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	F_errfinish(m, int32(_a_F_pg_restore_attribute_stats_4), int32(231), int32(_a_F_pg_restore_attribute_stats_5))
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L1
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
	v215 = v36 + int32(208)
	F_stats_check_required_arg(m, v215, int32(_a_F_pg_restore_attribute_stats_0), int32(4))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v36)+260))
	v222 = F_stats_check_arg_array(m, v215, int32(9))
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	v225 = F_stats_check_arg_array(m, v215, int32(13))
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	v228 = F_stats_check_arg_array(m, v215, int32(14))
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	v232 = F_stats_check_arg_pair(m, v215, int32(8), int32(9))
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	v236 = F_stats_check_arg_pair(m, v215, int32(12), int32(13))
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	v240 = F_stats_check_arg_pair(m, v215, int32(15), int32(16))
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	v243 = F_relation_open(m, v123, int32(1))
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	v246 = F_SearchSysCache2(m, int32(7), v123, v211)
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	if v246 == int32(0) {
		goto L9
	} else {
		goto L60
	}
L60:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v246)+16))
	v251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v250)+22)))
	v252 = v250 + v251
	v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v252)+91)))
	if v253 == int32(1) {
		goto L8
	} else {
		goto L61
	}
L61:
	;
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v243)+48))
	v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v256)+119)))
	if v257|int32(32) != int32(105) {
		goto L64
	} else {
		goto L65
	}
L62:
	;
	v527 = int32(1)
	v529 = v225 & v236 & (v57 | v58 ^ v527)
	v532 = v228 & (v63 ^ v527)
	F_ReleaseCatCache(m, v246)
	mBase = m.M
	v534 = m.ExcPending
	if v534 != 0 {
		goto L1
	} else {
		goto L103
	}
L63:
	;
	v480 = F_exprType(m, v443)
	mBase = m.M
	v481 = m.ExcPending
	if v481 != 0 {
		goto L1
	} else {
		goto L99
	}
L64:
	;
	v477 = *(*int32)(unsafe.Add(mBase, uint32(v252)+96))
	v478 = *(*int32)(unsafe.Add(mBase, uint32(v252)+76))
	v479 = *(*int32)(unsafe.Add(mBase, uint32(v252)+68))
	v487 = v479
	v490 = v477
	v495 = v478
	goto L62
L65:
	;
	v262 = int32(*(*int16)(unsafe.Add(mBase, uint32(v252)+74)))
	v263 = F_RelationGetIndexExpressions(m, v243)
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	if v263 == int32(0) {
		goto L64
	} else {
		goto L67
	}
L67:
	;
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v243)+192))
	v269 = v267 + int32(48)
	v270 = int32(1)
	v271 = v262 - v270
	v275 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v269+v271<<(uint(v270)%32)))))
	if v275 != 0 {
		goto L64
	} else {
		goto L68
	}
L68:
	;
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v243)+228))
	if v276 != 0 {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v276)+12))
	v279 = v277
	goto L71
L70:
	;
	v279 = int32(0)
	goto L71
L71:
	;
	if v262 < int32(2) {
		v410 = v279
		goto L72
	} else {
		goto L73
	}
L72:
	;
	if v410 == int32(0) {
		goto L7
	} else {
		goto L97
	}
L73:
	;
	v282 = int32(0)
	if v262 != int32(2) {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v290 = v282
	v292 = v279
	v293 = int32(0)
	goto L77
L75:
	;
	v361 = v282
	v363 = v279
	goto L76
L76:
	;
	v397 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v269+v361<<(uint(int32(1))%32)))))
	if v397 != 0 {
		v410 = v363
		goto L72
	} else {
		goto L93
	}
L77:
	;
	v325 = v269 + v290<<(uint(int32(1))%32)
	v326 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v325))))
	if v326 == int32(0) {
		goto L79
	} else {
		goto L80
	}
L78:
	;
	if v271&int32(1) == int32(0) {
		v410 = v353
		goto L72
	} else {
		goto L92
	}
L79:
	;
	v330 = v292 + int32(4)
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v276)+12))
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v276)+4))
	if base.Ui32(v330) < base.Ui32(v332+v333<<(uint(int32(2))%32)) {
		goto L82
	} else {
		goto L83
	}
L80:
	;
	v339 = v292
	goto L81
L81:
	;
	v340 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v325)+2)))
	if v340 == int32(0) {
		goto L85
	} else {
		goto L86
	}
L82:
	;
	v338 = v330
	goto L84
L83:
	;
	v338 = int32(0)
	goto L84
L84:
	;
	v339 = v338
	goto L81
L85:
	;
	v344 = v339 + int32(4)
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v276)+12))
	v347 = *(*int32)(unsafe.Add(mBase, uint32(v276)+4))
	if base.Ui32(v344) < base.Ui32(v346+v347<<(uint(int32(2))%32)) {
		goto L88
	} else {
		goto L89
	}
L86:
	;
	v353 = v339
	goto L87
L87:
	;
	v354 = int32(2)
	v355 = v290 + v354
	v357 = v293 + v354
	if v357 != v271&int32(-2) {
		v290 = v355
		v292 = v353
		v293 = v357
		goto L77
	} else {
		goto L91
	}
L88:
	;
	v352 = v344
	goto L90
L89:
	;
	v352 = int32(0)
	goto L90
L90:
	;
	v353 = v352
	goto L87
L91:
	;
	goto L78
L92:
	;
	v361 = v355
	v363 = v353
	goto L76
L93:
	;
	v399 = v363 + int32(4)
	v401 = *(*int32)(unsafe.Add(mBase, uint32(v276)+12))
	v402 = *(*int32)(unsafe.Add(mBase, uint32(v276)+4))
	if base.Ui32(v399) < base.Ui32(v401+v402<<(uint(int32(2))%32)) {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	v407 = v399
	goto L96
L95:
	;
	v407 = int32(0)
	goto L96
L96:
	;
	v410 = v407
	goto L72
L97:
	;
	v443 = *(*int32)(unsafe.Add(mBase, uint32(v410)))
	if v443 != 0 {
		goto L63
	} else {
		goto L98
	}
L98:
	;
	goto L64
L99:
	;
	v482 = F_exprTypmod(m, v443)
	mBase = m.M
	v483 = m.ExcPending
	if v483 != 0 {
		goto L1
	} else {
		goto L100
	}
L100:
	;
	v484 = *(*int32)(unsafe.Add(mBase, uint32(v252)+96))
	if v484 != 0 {
		v487 = v480
		v490 = v484
		v495 = v482
		goto L62
	} else {
		goto L101
	}
L101:
	;
	v485 = F_exprCollation(m, v443)
	mBase = m.M
	v486 = m.ExcPending
	if v486 != 0 {
		goto L1
	} else {
		goto L102
	}
L102:
	;
	v487 = v480
	v490 = v485
	v495 = v482
	goto L62
L103:
	;
	v535 = F_type_is_multirange(m, v487)
	mBase = m.M
	v536 = m.ExcPending
	if v536 != 0 {
		goto L1
	} else {
		goto L104
	}
L104:
	;
	if v535 != 0 {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	v537 = F_get_multirange_range(m, v487)
	mBase = m.M
	v538 = m.ExcPending
	if v538 != 0 {
		goto L1
	} else {
		goto L108
	}
L106:
	;
	v539 = v487
	goto L107
L107:
	;
	v542 = v236 & v240 & v232 & v228 & v225 & v222
	v544 = F_lookup_type_cache(m, v539, int32(3))
	mBase = m.M
	v545 = m.ExcPending
	if v545 != 0 {
		goto L1
	} else {
		goto L109
	}
L108:
	;
	v539 = v537
	goto L107
L109:
	;
	v546 = *(*int32)(unsafe.Add(mBase, uint32(v544)+56))
	v547 = *(*int32)(unsafe.Add(mBase, uint32(v544)+52))
	v548 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v544)+13)))
	v549 = int32(0)
	F_relation_close(m, v243, v549)
	mBase = m.M
	v552 = m.ExcPending
	if v552 != 0 {
		goto L1
	} else {
		goto L110
	}
L110:
	;
	v553 = v532 | v529
	if v539 == int32(3614) {
		goto L115
	} else {
		goto L116
	}
L111:
	;
	v620 = int32(1)
	v624 = v240 & (v59 | v60 ^ int32(1))
	if v61&v62|v546 == int32(0) {
		goto L133
	} else {
		goto L134
	}
L112:
	;
	v612 = int32(0)
	v614 = v608
	v615 = v609
	v616 = v610
	v617 = v611
	v618 = v612
	v619 = v612
	goto L111
L113:
	;
	v581 = int32(0)
	v584 = F_errstart(m, int32(19), v581)
	mBase = m.M
	v585 = m.ExcPending
	if v585 != 0 {
		goto L1
	} else {
		goto L126
	}
L114:
	;
	v575 = F_lookup_type_cache(m, v571, int32(1))
	mBase = m.M
	v576 = m.ExcPending
	if v576 != 0 {
		goto L1
	} else {
		goto L124
	}
L115:
	;
	v558 = int32(100)
	if v553&int32(1) != 0 {
		v571 = int32(25)
		v572 = v558
		goto L114
	} else {
		goto L118
	}
L116:
	;
	goto L117
L117:
	;
	if v553&int32(1) == int32(0) {
		goto L119
	} else {
		goto L120
	}
L118:
	;
	v608 = int32(0)
	v609 = v558
	v610 = int32(0)
	v611 = v542
	goto L112
L119:
	;
	v608 = v549
	v609 = v490
	v610 = int32(0)
	v611 = v542
	goto L112
L120:
	;
	goto L121
L121:
	;
	v567 = F_get_base_element_type(m, v539)
	mBase = m.M
	v568 = m.ExcPending
	if v568 != 0 {
		goto L1
	} else {
		goto L122
	}
L122:
	;
	if v567 == int32(0) {
		v579 = v490
		goto L113
	} else {
		goto L123
	}
L123:
	;
	v571 = v567
	v572 = v490
	goto L114
L124:
	;
	v577 = *(*int32)(unsafe.Add(mBase, uint32(v575)+52))
	if v577 != 0 {
		v614 = v571
		v615 = v572
		v616 = v577
		v617 = v542
		v618 = v532
		v619 = v529
		goto L111
	} else {
		goto L125
	}
L125:
	;
	v579 = v572
	goto L113
L126:
	;
	if v584 != 0 {
		goto L127
	} else {
		goto L128
	}
L127:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+160)) = v210
	F_errmsg(m, int32(_a_F_pg_restore_attribute_stats_20), v36+int32(160))
	mBase = m.M
	v591 = m.ExcPending
	if v591 != 0 {
		goto L1
	} else {
		goto L130
	}
L128:
	;
	goto L129
L129:
	;
	v606 = int32(0)
	v608 = v581
	v609 = v579
	v610 = v606
	v611 = v606
	goto L112
L130:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+148)) = int32(_a_F_pg_restore_attribute_stats_21)
	*(*int32)(unsafe.Add(mBase, uint32(v36)+144)) = int32(_a_F_pg_restore_attribute_stats_22)
	F_errdetail(m, int32(_a_F_pg_restore_attribute_stats_12), v36+int32(144))
	mBase = m.M
	v600 = m.ExcPending
	if v600 != 0 {
		goto L1
	} else {
		goto L131
	}
L131:
	;
	F_errfinish(m, int32(_a_F_pg_restore_attribute_stats_4), int32(305), int32(_a_F_pg_restore_attribute_stats_5))
	mBase = m.M
	v605 = m.ExcPending
	if v605 != 0 {
		goto L1
	} else {
		goto L132
	}
L132:
	;
	goto L129
L133:
	;
	v629 = int32(0)
	v632 = F_errstart(m, int32(19), v629)
	mBase = m.M
	v633 = m.ExcPending
	if v633 != 0 {
		goto L1
	} else {
		goto L136
	}
L134:
	;
	v659 = v617
	v660 = v61 ^ v620
	v661 = v62 ^ v620
	goto L135
L135:
	;
	v662 = int32(1)
	v663 = v64 ^ v662
	if (v624^v662)&v64 != 0 {
		v699 = v659
		v700 = v624
		v701 = v663
		goto L144
	} else {
		goto L145
	}
L136:
	;
	if v632 != 0 {
		goto L137
	} else {
		goto L138
	}
L137:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v636 = m.ExcPending
	if v636 != 0 {
		goto L1
	} else {
		goto L140
	}
L138:
	;
	goto L139
L139:
	;
	v657 = int32(0)
	v659 = v629
	v660 = v657
	v661 = v657
	goto L135
L140:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+128)) = v210
	F_errmsg(m, int32(_a_F_pg_restore_attribute_stats_9), v36+int32(128))
	mBase = m.M
	v642 = m.ExcPending
	if v642 != 0 {
		goto L1
	} else {
		goto L141
	}
L141:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+116)) = int32(_a_F_pg_restore_attribute_stats_10)
	*(*int32)(unsafe.Add(mBase, uint32(v36)+112)) = int32(_a_F_pg_restore_attribute_stats_11)
	F_errdetail(m, int32(_a_F_pg_restore_attribute_stats_12), v36+int32(112))
	mBase = m.M
	v651 = m.ExcPending
	if v651 != 0 {
		goto L1
	} else {
		goto L142
	}
L142:
	;
	F_errfinish(m, int32(_a_F_pg_restore_attribute_stats_4), int32(322), int32(_a_F_pg_restore_attribute_stats_5))
	mBase = m.M
	v656 = m.ExcPending
	if v656 != 0 {
		goto L1
	} else {
		goto L143
	}
L143:
	;
	goto L139
L144:
	;
	F_fmgr_info(m, int32(750), v36+int32(576))
	mBase = m.M
	v706 = m.ExcPending
	if v706 != 0 {
		goto L1
	} else {
		goto L155
	}
L145:
	;
	switch v548 - int32(109) {
	case 0, 5:
		v699 = v659
		v700 = v624
		v701 = v663
		goto L144
	default:
		goto L146
	}
L146:
	;
	v669 = int32(0)
	v672 = F_errstart(m, int32(19), v669)
	mBase = m.M
	v673 = m.ExcPending
	if v673 != 0 {
		goto L1
	} else {
		goto L147
	}
L147:
	;
	if v672 != 0 {
		goto L148
	} else {
		goto L149
	}
L148:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v676 = m.ExcPending
	if v676 != 0 {
		goto L1
	} else {
		goto L151
	}
L149:
	;
	goto L150
L150:
	;
	v697 = int32(0)
	v699 = v669
	v700 = v697
	v701 = v697
	goto L144
L151:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+96)) = v210
	F_errmsg(m, int32(_a_F_pg_restore_attribute_stats_23), v36+int32(96))
	mBase = m.M
	v682 = m.ExcPending
	if v682 != 0 {
		goto L1
	} else {
		goto L152
	}
L152:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+84)) = int32(_a_F_pg_restore_attribute_stats_24)
	*(*int32)(unsafe.Add(mBase, uint32(v36)+80)) = int32(_a_F_pg_restore_attribute_stats_25)
	F_errdetail(m, int32(_a_F_pg_restore_attribute_stats_12), v36+int32(80))
	mBase = m.M
	v691 = m.ExcPending
	if v691 != 0 {
		goto L1
	} else {
		goto L153
	}
L153:
	;
	F_errfinish(m, int32(_a_F_pg_restore_attribute_stats_4), int32(337), int32(_a_F_pg_restore_attribute_stats_5))
	mBase = m.M
	v696 = m.ExcPending
	if v696 != 0 {
		goto L1
	} else {
		goto L154
	}
L154:
	;
	goto L150
L155:
	;
	v709 = F_table_open(m, int32(2619), int32(3))
	mBase = m.M
	v710 = m.ExcPending
	if v710 != 0 {
		goto L1
	} else {
		goto L156
	}
L156:
	;
	v713 = base.B2i32(v220 != int32(0))
	v714 = F_SearchSysCache3(m, int32(65), v123, v211, v713)
	mBase = m.M
	v715 = m.ExcPending
	if v715 != 0 {
		goto L1
	} else {
		goto L158
	}
L157:
	;
	v782 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+272)))
	if v782&int32(1) == int32(0) {
		goto L163
	} else {
		goto L164
	}
L158:
	;
	if v714 != 0 {
		goto L159
	} else {
		goto L160
	}
L159:
	;
	v716 = *(*int32)(unsafe.Add(mBase, uint32(v709)+52))
	F_heap_deform_tuple(m, v714, v716, v36+int32(448), v36+int32(416))
	mBase = m.M
	v722 = m.ExcPending
	if v722 != 0 {
		goto L1
	} else {
		goto L162
	}
L160:
	;
	goto L161
L161:
	;
	v723 = int32(257)
	*(*uint16)(unsafe.Add(mBase, uint32(v36)+445)) = uint16(v723)
	v725 = int64(72340172838076673)
	*(*int64)(unsafe.Add(mBase, uint32(v36)+437)) = v725
	*(*int64)(unsafe.Add(mBase, uint32(v36)+384)) = v725
	*(*int64)(unsafe.Add(mBase, uint32(v36)+392)) = v725
	*(*int64)(unsafe.Add(mBase, uint32(v36)+400)) = v725
	*(*int64)(unsafe.Add(mBase, uint32(v36)+407)) = v725
	*(*int32)(unsafe.Add(mBase, uint32(v36)+448)) = v123
	v736 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v36)+416)) = uint8(v736)
	*(*int32)(unsafe.Add(mBase, uint32(v36)+452)) = v211
	*(*int32)(unsafe.Add(mBase, uint32(v36)+456)) = v713
	*(*int32)(unsafe.Add(mBase, uint32(v36)+460)) = v736
	*(*int32)(unsafe.Add(mBase, uint32(v36)+417)) = v736
	v744 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v36)+464)) = v744
	*(*int32)(unsafe.Add(mBase, uint32(v36)+472)) = v736
	*(*uint16)(unsafe.Add(mBase, uint32(v36)+421)) = uint16(v736)
	*(*int32)(unsafe.Add(mBase, uint32(v36)+492)) = v736
	*(*uint8)(unsafe.Add(mBase, uint32(v36)+427)) = uint8(v736)
	*(*int64)(unsafe.Add(mBase, uint32(v36)+429)) = v744
	*(*int64)(unsafe.Add(mBase, uint32(v36)+512)) = v744
	*(*int32)(unsafe.Add(mBase, uint32(v36)+476)) = v736
	*(*uint8)(unsafe.Add(mBase, uint32(v36)+423)) = uint8(v736)
	*(*int32)(unsafe.Add(mBase, uint32(v36)+496)) = v736
	*(*uint8)(unsafe.Add(mBase, uint32(v36)+428)) = uint8(v736)
	*(*uint16)(unsafe.Add(mBase, uint32(v36)+424)) = uint16(v736)
	*(*int64)(unsafe.Add(mBase, uint32(v36)+520)) = v744
	*(*int64)(unsafe.Add(mBase, uint32(v36)+500)) = v744
	*(*int64)(unsafe.Add(mBase, uint32(v36)+480)) = v744
	*(*int32)(unsafe.Add(mBase, uint32(v36)+488)) = v736
	*(*uint8)(unsafe.Add(mBase, uint32(v36)+426)) = uint8(v736)
	*(*int32)(unsafe.Add(mBase, uint32(v36)+508)) = v736
	*(*int32)(unsafe.Add(mBase, uint32(v36)+528)) = v736
	goto L157
L162:
	;
	goto L157
L163:
	;
	v787 = *(*int32)(unsafe.Add(mBase, uint32(v36)+268))
	*(*int32)(unsafe.Add(mBase, uint32(v36)+460)) = v787
	v789 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v36)+387)) = uint8(v789)
	goto L165
L164:
	;
	goto L165
L165:
	;
	v791 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+280)))
	if v791&int32(1) == int32(0) {
		goto L166
	} else {
		goto L167
	}
L166:
	;
	v796 = *(*int32)(unsafe.Add(mBase, uint32(v36)+276))
	*(*int32)(unsafe.Add(mBase, uint32(v36)+464)) = v796
	v798 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v36)+388)) = uint8(v798)
	goto L168
L167:
	;
	goto L168
L168:
	;
	v800 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+288)))
	if v800&int32(1) == int32(0) {
		goto L169
	} else {
		goto L170
	}
L169:
	;
	v805 = *(*int32)(unsafe.Add(mBase, uint32(v36)+284))
	*(*int32)(unsafe.Add(mBase, uint32(v36)+468)) = v805
	v807 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v36)+389)) = uint8(v807)
	goto L171
L170:
	;
	goto L171
L171:
	;
	v809 = int32(1)
	v814 = int32(0)
	if v56&v809|(base.B2i32(v222&v232&v809 == v814)|v55&v809) == v814 {
		goto L173
	} else {
		goto L174
	}
L172:
	;
	if v660&int32(1) != 0 {
		goto L180
	} else {
		goto L181
	}
L173:
	;
	v822 = *(*int32)(unsafe.Add(mBase, uint32(v36)+300))
	v826 = *(*int32)(unsafe.Add(mBase, uint32(v36)+292))
	v829 = F_text_to_stavalues(m, int32(_a_F_pg_restore_attribute_stats_13), v36+int32(576), v826, v539, v495, v36+int32(380))
	mBase = m.M
	v830 = m.ExcPending
	if v830 != 0 {
		goto L1
	} else {
		goto L176
	}
L174:
	;
	goto L175
L175:
	;
	v850 = v699
	goto L172
L176:
	;
	v832 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+380)))
	if v832 != int32(1) {
		v850 = int32(0)
		goto L172
	} else {
		goto L177
	}
L177:
	;
	v842 = int32(0)
	F_set_stats_slot(m, v36+int32(448), v36+int32(416), v36+int32(384), int32(1), v547, v615, v822, v842, v829, v842)
	mBase = m.M
	v845 = m.ExcPending
	if v845 != 0 {
		goto L1
	} else {
		goto L178
	}
L178:
	;
	goto L175
L179:
	;
	if v661&int32(1) != 0 {
		goto L186
	} else {
		goto L187
	}
L180:
	;
	v856 = *(*int32)(unsafe.Add(mBase, uint32(v36)+308))
	v859 = F_text_to_stavalues(m, int32(_a_F_pg_restore_attribute_stats_14), v36+int32(576), v856, v539, v495, v36+int32(380))
	mBase = m.M
	v860 = m.ExcPending
	if v860 != 0 {
		goto L1
	} else {
		goto L183
	}
L181:
	;
	goto L182
L182:
	;
	v879 = v850
	goto L179
L183:
	;
	v862 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+380)))
	if v862 != int32(1) {
		v879 = int32(0)
		goto L179
	} else {
		goto L184
	}
L184:
	;
	v872 = int32(0)
	F_set_stats_slot(m, v36+int32(448), v36+int32(416), v36+int32(384), int32(2), v546, v615, v872, int32(1), v859, v872)
	mBase = m.M
	v876 = m.ExcPending
	if v876 != 0 {
		goto L1
	} else {
		goto L185
	}
L185:
	;
	goto L182
L186:
	;
	v882 = *(*int32)(unsafe.Add(mBase, uint32(v36)+316))
	*(*int32)(unsafe.Add(mBase, uint32(v36)+380)) = v882
	v895 = F_construct_array_builtin(m, v36+int32(380), int32(1), int32(700))
	mBase = m.M
	v896 = m.ExcPending
	if v896 != 0 {
		goto L1
	} else {
		goto L189
	}
L187:
	;
	goto L188
L188:
	;
	if v619&int32(1) != 0 {
		goto L192
	} else {
		goto L193
	}
L189:
	;
	v897 = int32(0)
	F_set_stats_slot(m, v36+int32(448), v36+int32(416), v36+int32(384), int32(3), v546, v615, v895, v897, v897, int32(1))
	mBase = m.M
	v901 = m.ExcPending
	if v901 != 0 {
		goto L1
	} else {
		goto L190
	}
L190:
	;
	goto L188
L191:
	;
	if v618&int32(1) != 0 {
		goto L198
	} else {
		goto L199
	}
L192:
	;
	v904 = *(*int32)(unsafe.Add(mBase, uint32(v36)+332))
	v908 = *(*int32)(unsafe.Add(mBase, uint32(v36)+324))
	v911 = F_text_to_stavalues(m, int32(_a_F_pg_restore_attribute_stats_15), v36+int32(576), v908, v614, v495, v36+int32(380))
	mBase = m.M
	v912 = m.ExcPending
	if v912 != 0 {
		goto L1
	} else {
		goto L195
	}
L193:
	;
	goto L194
L194:
	;
	v932 = v879
	goto L191
L195:
	;
	v914 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+380)))
	if v914 != int32(1) {
		v932 = int32(0)
		goto L191
	} else {
		goto L196
	}
L196:
	;
	v924 = int32(0)
	F_set_stats_slot(m, v36+int32(448), v36+int32(416), v36+int32(384), int32(4), v616, v615, v904, v924, v911, v924)
	mBase = m.M
	v927 = m.ExcPending
	if v927 != 0 {
		goto L1
	} else {
		goto L197
	}
L197:
	;
	goto L194
L198:
	;
	v942 = *(*int32)(unsafe.Add(mBase, uint32(v36)+340))
	v943 = int32(0)
	F_set_stats_slot(m, v36+int32(448), v36+int32(416), v36+int32(384), int32(5), v616, v615, v942, v943, v943, int32(1))
	mBase = m.M
	v947 = m.ExcPending
	if v947 != 0 {
		goto L1
	} else {
		goto L201
	}
L199:
	;
	goto L200
L200:
	;
	if v701&int32(1) != 0 {
		goto L203
	} else {
		goto L204
	}
L201:
	;
	goto L200
L202:
	;
	if v700 != 0 {
		goto L210
	} else {
		goto L211
	}
L203:
	;
	v953 = *(*int32)(unsafe.Add(mBase, uint32(v36)+364))
	v956 = F_text_to_stavalues(m, int32(_a_F_pg_restore_attribute_stats_16), v36+int32(576), v953, v539, v495, v36+int32(380))
	mBase = m.M
	v957 = m.ExcPending
	if v957 != 0 {
		goto L1
	} else {
		goto L206
	}
L204:
	;
	goto L205
L205:
	;
	v978 = v932
	goto L202
L206:
	;
	v959 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+380)))
	if v959 != int32(1) {
		v978 = int32(0)
		goto L202
	} else {
		goto L207
	}
L207:
	;
	v969 = int32(0)
	F_set_stats_slot(m, v36+int32(448), v36+int32(416), v36+int32(384), int32(7), v969, v969, v969, int32(1), v956, v969)
	mBase = m.M
	v975 = m.ExcPending
	if v975 != 0 {
		goto L1
	} else {
		goto L208
	}
L208:
	;
	goto L205
L209:
	;
	v1020 = *(*int32)(unsafe.Add(mBase, uint32(v709)+52))
	if v714 != 0 {
		goto L218
	} else {
		goto L219
	}
L210:
	;
	v979 = *(*int32)(unsafe.Add(mBase, uint32(v36)+356))
	*(*int32)(unsafe.Add(mBase, uint32(v36)+380)) = v979
	v985 = F_construct_array_builtin(m, v36+int32(380), int32(1), int32(700))
	mBase = m.M
	v986 = m.ExcPending
	if v986 != 0 {
		goto L1
	} else {
		goto L213
	}
L211:
	;
	goto L212
L212:
	;
	v1018 = v978
	goto L209
L213:
	;
	v987 = int32(0)
	v991 = *(*int32)(unsafe.Add(mBase, uint32(v36)+348))
	v996 = F_text_to_stavalues(m, int32(_a_F_pg_restore_attribute_stats_17), v36+int32(576), v991, int32(701), v987, v36+int32(379))
	mBase = m.M
	v997 = m.ExcPending
	if v997 != 0 {
		goto L1
	} else {
		goto L214
	}
L214:
	;
	v998 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+379)))
	if v998 != int32(1) {
		v1018 = v987
		goto L209
	} else {
		goto L215
	}
L215:
	;
	v1009 = int32(0)
	F_set_stats_slot(m, v36+int32(448), v36+int32(416), v36+int32(384), int32(6), int32(672), v1009, v985, v1009, v996, v1009)
	mBase = m.M
	v1013 = m.ExcPending
	if v1013 != 0 {
		goto L1
	} else {
		goto L216
	}
L216:
	;
	goto L212
L217:
	;
	F_pfree(m, v1041)
	mBase = m.M
	v1043 = m.ExcPending
	if v1043 != 0 {
		goto L1
	} else {
		goto L225
	}
L218:
	;
	v1027 = F_heap_modify_tuple(m, v714, v1020, v36+int32(448), v36+int32(416), v36+int32(384))
	mBase = m.M
	v1028 = m.ExcPending
	if v1028 != 0 {
		goto L1
	} else {
		goto L221
	}
L219:
	;
	goto L220
L220:
	;
	v1037 = F_heap_form_tuple(m, v1020, v36+int32(448), v36+int32(416))
	mBase = m.M
	v1038 = m.ExcPending
	if v1038 != 0 {
		goto L1
	} else {
		goto L223
	}
L221:
	;
	F_CatalogTupleUpdate(m, v709, v1027+int32(4), v1027)
	mBase = m.M
	v1032 = m.ExcPending
	if v1032 != 0 {
		goto L1
	} else {
		goto L222
	}
L222:
	;
	v1041 = v1027
	goto L217
L223:
	;
	F_CatalogTupleInsert(m, v709, v1037)
	mBase = m.M
	v1040 = m.ExcPending
	if v1040 != 0 {
		goto L1
	} else {
		goto L224
	}
L224:
	;
	v1041 = v1037
	goto L217
L225:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v1045 = m.ExcPending
	if v1045 != 0 {
		goto L1
	} else {
		goto L226
	}
L226:
	;
	if v714 != 0 {
		goto L227
	} else {
		goto L228
	}
L227:
	;
	F_ReleaseCatCache(m, v714)
	mBase = m.M
	v1047 = m.ExcPending
	if v1047 != 0 {
		goto L1
	} else {
		goto L230
	}
L228:
	;
	goto L229
L229:
	;
	F_relation_close(m, v709, int32(3))
	mBase = m.M
	v1050 = m.ExcPending
	if v1050 != 0 {
		goto L1
	} else {
		goto L231
	}
L230:
	;
	goto L229
L231:
	;
	m.G0 = v36 + int32(608)
	return v1018 & v49
L232:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v1062 = m.ExcPending
	if v1062 != 0 {
		goto L1
	} else {
		goto L233
	}
L233:
	;
	F_errmsg(m, int32(_a_F_pg_restore_attribute_stats_26), int32(0))
	mBase = m.M
	v1066 = m.ExcPending
	if v1066 != 0 {
		goto L1
	} else {
		goto L234
	}
L234:
	;
	F_errhint(m, int32(_a_F_pg_restore_attribute_stats_27), int32(0))
	mBase = m.M
	v1070 = m.ExcPending
	if v1070 != 0 {
		goto L1
	} else {
		goto L235
	}
L235:
	;
	F_errfinish(m, int32(_a_F_pg_restore_attribute_stats_4), int32(192), int32(_a_F_pg_restore_attribute_stats_5))
	mBase = m.M
	v1075 = m.ExcPending
	if v1075 != 0 {
		goto L1
	} else {
		goto L236
	}
L236:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L237:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v1082 = m.ExcPending
	if v1082 != 0 {
		goto L1
	} else {
		goto L238
	}
L238:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+196)) = int32(_a_F_pg_restore_attribute_stats_1)
	*(*int32)(unsafe.Add(mBase, uint32(v36)+192)) = int32(_a_F_pg_restore_attribute_stats_2)
	F_errmsg(m, int32(_a_F_pg_restore_attribute_stats_3), v36+int32(192))
	mBase = m.M
	v1091 = m.ExcPending
	if v1091 != 0 {
		goto L1
	} else {
		goto L239
	}
L239:
	;
	F_errfinish(m, int32(_a_F_pg_restore_attribute_stats_4), int32(205), int32(_a_F_pg_restore_attribute_stats_5))
	mBase = m.M
	v1096 = m.ExcPending
	if v1096 != 0 {
		goto L1
	} else {
		goto L240
	}
L240:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L241:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v1103 = m.ExcPending
	if v1103 != 0 {
		goto L1
	} else {
		goto L242
	}
L242:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+32)) = v210
	F_errmsg(m, int32(_a_F_pg_restore_attribute_stats_6), v36+int32(32))
	mBase = m.M
	v1109 = m.ExcPending
	if v1109 != 0 {
		goto L1
	} else {
		goto L243
	}
L243:
	;
	F_errfinish(m, int32(_a_F_pg_restore_attribute_stats_4), int32(240), int32(_a_F_pg_restore_attribute_stats_5))
	mBase = m.M
	v1114 = m.ExcPending
	if v1114 != 0 {
		goto L1
	} else {
		goto L244
	}
L244:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L245:
	;
	F_errcode(m, int32(50360452))
	mBase = m.M
	v1121 = m.ExcPending
	if v1121 != 0 {
		goto L1
	} else {
		goto L246
	}
L246:
	;
	v1122 = *(*int32)(unsafe.Add(mBase, uint32(v243)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v36)+48)) = v211
	*(*int32)(unsafe.Add(mBase, uint32(v36)+52)) = v1122 + int32(4)
	F_errmsg(m, int32(_a_F_pg_restore_attribute_stats_7), v36+int32(48))
	mBase = m.M
	v1131 = m.ExcPending
	if v1131 != 0 {
		goto L1
	} else {
		goto L247
	}
L247:
	;
	F_errfinish(m, int32(_a_F_pg_restore_attribute_stats_4), int32(597), int32(_a_F_pg_restore_attribute_stats_8))
	mBase = m.M
	v1136 = m.ExcPending
	if v1136 != 0 {
		goto L1
	} else {
		goto L248
	}
L248:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L249:
	;
	F_errcode(m, int32(50360452))
	mBase = m.M
	v1143 = m.ExcPending
	if v1143 != 0 {
		goto L1
	} else {
		goto L250
	}
L250:
	;
	v1144 = *(*int32)(unsafe.Add(mBase, uint32(v243)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v36)+64)) = v211
	*(*int32)(unsafe.Add(mBase, uint32(v36)+68)) = v1144 + int32(4)
	F_errmsg(m, int32(_a_F_pg_restore_attribute_stats_7), v36-int32(-64))
	mBase = m.M
	v1153 = m.ExcPending
	if v1153 != 0 {
		goto L1
	} else {
		goto L251
	}
L251:
	;
	F_errfinish(m, int32(_a_F_pg_restore_attribute_stats_4), int32(605), int32(_a_F_pg_restore_attribute_stats_8))
	mBase = m.M
	v1158 = m.ExcPending
	if v1158 != 0 {
		goto L1
	} else {
		goto L252
	}
L252:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L253:
	;
	F_errmsg_internal(m, int32(_a_F_pg_restore_attribute_stats_28), int32(0))
	mBase = m.M
	v1166 = m.ExcPending
	if v1166 != 0 {
		goto L1
	} else {
		goto L254
	}
L254:
	;
	F_errfinish(m, int32(_a_F_pg_restore_attribute_stats_4), int32(569), int32(_a_F_pg_restore_attribute_stats_29))
	mBase = m.M
	v1171 = m.ExcPending
	if v1171 != 0 {
		goto L1
	} else {
		goto L255
	}
L255:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
