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
	var v60 int32
	_ = v60
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
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v123 int32
	_ = v123
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
	return v123
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
	v123 = v24
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
	F_errmsg(m, int32(69248), v15)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L2
	} else {
		goto L12
	}
L12:
	;
	F_errfinish(m, int32(497925), int32(3949), int32(63580))
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
		v123 = v50
		goto L1
	} else {
		goto L15
	}
L15:
	;
	v59 = v50
	v60 = int32(1)
	goto L16
L16:
	;
	v68 = F_SearchSysCache2(m, int32(7), l0, base.I32_extend16_s(v60))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L2
	} else {
		goto L19
	}
L17:
	;
	v123 = v111
	goto L1
L18:
	;
	v116 = base.I32_extend16_s(v60 + int32(1))
	if v116 <= v47 {
		v59 = v111
		v60 = v116
		goto L16
	} else {
		goto L41
	}
L19:
	;
	if v68 == int32(0) {
		v111 = v59
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
	v111 = v59
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
	v108 = int32(1)
	if l3 == int32(0) {
		v123 = v108
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
	if v91 != v84 {
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
		v123 = v103
		goto L1
	} else {
		goto L39
	}
L39:
	;
	v111 = v103
	goto L18
L40:
	;
	v111 = v108
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
	var v70 int32
	_ = v70
	var v71 int64
	_ = v71
	var v92 int32
	_ = v92
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
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v186 int32
	_ = v186
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v205 int32
	_ = v205
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
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
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v309 int32
	_ = v309
	var v312 int32
	_ = v312
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v346 int32
	_ = v346
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v360 int32
	_ = v360
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v374 int32
	_ = v374
	var v379 int32
	_ = v379
	var v382 int32
	_ = v382
	var v386 int32
	_ = v386
	var v390 int32
	_ = v390
	var v395 int32
	_ = v395
	var v399 int32
	_ = v399
	var v402 int32
	_ = v402
	var v411 int32
	_ = v411
	var v416 int32
	_ = v416
	var v420 int32
	_ = v420
	var v423 int32
	_ = v423
	var v429 int32
	_ = v429
	var v434 int32
	_ = v434
	var v438 int32
	_ = v438
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v451 int32
	_ = v451
	var v456 int32
	_ = v456
	var v460 int32
	_ = v460
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v473 int32
	_ = v473
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v482 int32
	_ = v482
	var v517 int32
	_ = v517
	var v519 int32
	_ = v519
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v527 int32
	_ = v527
	var v531 int32
	_ = v531
	var v566 int32
	_ = v566
	var v570 int32
	_ = v570
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
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
	var v626 int32
	_ = v626
	var v631 int32
	_ = v631
	var v660 int32
	_ = v660
	var v662 int32
	_ = v662
	var v665 int32
	_ = v665
	var v667 int32
	_ = v667
	var v668 int32
	_ = v668
	var v669 int32
	_ = v669
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	var v672 int32
	_ = v672
	var v675 int32
	_ = v675
	var v677 int32
	_ = v677
	var v678 int32
	_ = v678
	var v679 int32
	_ = v679
	var v680 int32
	_ = v680
	var v681 int32
	_ = v681
	var v682 int32
	_ = v682
	var v685 int32
	_ = v685
	var v686 int32
	_ = v686
	var v691 int32
	_ = v691
	var v700 int32
	_ = v700
	var v701 int32
	_ = v701
	var v704 int32
	_ = v704
	var v705 int32
	_ = v705
	var v708 int32
	_ = v708
	var v709 int32
	_ = v709
	var v710 int32
	_ = v710
	var v711 int32
	_ = v711
	var v714 int32
	_ = v714
	var v717 int32
	_ = v717
	var v718 int32
	_ = v718
	var v724 int32
	_ = v724
	var v733 int32
	_ = v733
	var v738 int32
	_ = v738
	var v739 int32
	_ = v739
	var v741 int32
	_ = v741
	var v742 int32
	_ = v742
	var v743 int32
	_ = v743
	var v744 int32
	_ = v744
	var v745 int32
	_ = v745
	var v747 int32
	_ = v747
	var v748 int32
	_ = v748
	var v749 int32
	_ = v749
	var v750 int32
	_ = v750
	var v751 int32
	_ = v751
	var v752 int32
	_ = v752
	var v753 int32
	_ = v753
	var v754 int32
	_ = v754
	var v755 int32
	_ = v755
	var v759 int32
	_ = v759
	var v762 int32
	_ = v762
	var v763 int32
	_ = v763
	var v766 int32
	_ = v766
	var v767 int32
	_ = v767
	var v770 int32
	_ = v770
	var v776 int32
	_ = v776
	var v785 int32
	_ = v785
	var v790 int32
	_ = v790
	var v791 int32
	_ = v791
	var v793 int32
	_ = v793
	var v794 int32
	_ = v794
	var v795 int32
	_ = v795
	var v797 int32
	_ = v797
	var v800 int32
	_ = v800
	var v812 int32
	_ = v812
	var v815 int32
	_ = v815
	var v816 int32
	_ = v816
	var v819 int32
	_ = v819
	var v825 int32
	_ = v825
	var v834 int32
	_ = v834
	var v839 int32
	_ = v839
	var v840 int32
	_ = v840
	var v842 int32
	_ = v842
	var v843 int32
	_ = v843
	var v844 int32
	_ = v844
	var v849 int32
	_ = v849
	var v852 int32
	_ = v852
	var v853 int32
	_ = v853
	var v856 int32
	_ = v856
	var v857 int32
	_ = v857
	var v858 int32
	_ = v858
	var v859 int32
	_ = v859
	var v865 int32
	_ = v865
	var v866 int32
	_ = v866
	var v868 int64
	_ = v868
	var v879 int32
	_ = v879
	var v887 int64
	_ = v887
	var v925 int32
	_ = v925
	var v928 int32
	_ = v928
	var v930 int32
	_ = v930
	var v932 int32
	_ = v932
	var v935 int32
	_ = v935
	var v937 int32
	_ = v937
	var v940 int32
	_ = v940
	var v943 int32
	_ = v943
	var v945 int32
	_ = v945
	var v953 int32
	_ = v953
	var v957 int32
	_ = v957
	var v960 int32
	_ = v960
	var v961 int32
	_ = v961
	var v963 int32
	_ = v963
	var v973 int32
	_ = v973
	var v976 int32
	_ = v976
	var v981 int32
	_ = v981
	var v982 int32
	_ = v982
	var v988 int32
	_ = v988
	var v991 int32
	_ = v991
	var v992 int32
	_ = v992
	var v993 int32
	_ = v993
	var v1003 int32
	_ = v1003
	var v1007 int32
	_ = v1007
	var v1010 int32
	_ = v1010
	var v1014 int32
	_ = v1014
	var v1027 int32
	_ = v1027
	var v1028 int32
	_ = v1028
	var v1029 int32
	_ = v1029
	var v1033 int32
	_ = v1033
	var v1034 int32
	_ = v1034
	var v1035 int32
	_ = v1035
	var v1041 int32
	_ = v1041
	var v1044 int32
	_ = v1044
	var v1045 int32
	_ = v1045
	var v1046 int32
	_ = v1046
	var v1056 int32
	_ = v1056
	var v1059 int32
	_ = v1059
	var v1063 int32
	_ = v1063
	var v1073 int32
	_ = v1073
	var v1074 int32
	_ = v1074
	var v1078 int32
	_ = v1078
	var v1081 int32
	_ = v1081
	var v1087 int32
	_ = v1087
	var v1090 int32
	_ = v1090
	var v1091 int32
	_ = v1091
	var v1092 int32
	_ = v1092
	var v1102 int32
	_ = v1102
	var v1108 int32
	_ = v1108
	var v1112 int32
	_ = v1112
	var v1113 int32
	_ = v1113
	var v1119 int32
	_ = v1119
	var v1120 int32
	_ = v1120
	var v1121 int32
	_ = v1121
	var v1127 int32
	_ = v1127
	var v1132 int32
	_ = v1132
	var v1133 int32
	_ = v1133
	var v1134 int32
	_ = v1134
	var v1145 int32
	_ = v1145
	var v1149 int32
	_ = v1149
	var v1154 int32
	_ = v1154
	var v1156 int32
	_ = v1156
	var v1163 int32
	_ = v1163
	var v1164 int32
	_ = v1164
	var v1168 int32
	_ = v1168
	var v1173 int32
	_ = v1173
	var v1174 int32
	_ = v1174
	var v1176 int32
	_ = v1176
	var v1177 int32
	_ = v1177
	var v1179 int32
	_ = v1179
	var v1181 int32
	_ = v1181
	var v1183 int32
	_ = v1183
	var v1186 int32
	_ = v1186
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
	v49 = F_stats_fill_fcinfo_from_arg_pairs(m, l0, v36+int32(208), int32(4121808))
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
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+312)))
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+328)))
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+336)))
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+352)))
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+360)))
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+320)))
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+344)))
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+368)))
	v70 = F__emscripten_memset_bulkmem(m, v36+int32(448), base.I32_extend8_s(v53), int32(124))
	mBase = m.M
	goto L3
L3:
	;
	v71 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v36)+439)) = v71
	*(*int64)(unsafe.Add(mBase, uint32(v36)+432)) = v71
	*(*int64)(unsafe.Add(mBase, uint32(v36)+424)) = v71
	*(*int64)(unsafe.Add(mBase, uint32(v36)+416)) = v71
	*(*int64)(unsafe.Add(mBase, uint32(v36)+407)) = v71
	*(*int64)(unsafe.Add(mBase, uint32(v36)+400)) = v71
	*(*int64)(unsafe.Add(mBase, uint32(v36)+392)) = v71
	*(*int64)(unsafe.Add(mBase, uint32(v36)+384)) = v71
	F_stats_check_required_arg(m, v36+int32(208), int32(4121808), int32(0))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	F_stats_check_required_arg(m, v36+int32(208), int32(4121808), int32(1))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v36)+228))
	v100 = F_text_to_cstring(m, v99)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v36)+236))
	v103 = F_text_to_cstring(m, v102)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v107 = int32(*(*uint8)(unsafe.Add(mBase, _consts[112])))
	if v107 == int32(1) {
		goto L18
	} else {
		goto L19
	}
L8:
	;
	v660 = int32(1)
	v662 = v232 & v249 & (v58 | v59 ^ v660)
	v665 = v237 & (v63 ^ v660)
	F_ReleaseCatCache(m, v261)
	mBase = m.M
	v667 = m.ExcPending
	if v667 != 0 {
		goto L1
	} else {
		goto L130
	}
L9:
	;
	v613 = F_exprType(m, v576)
	mBase = m.M
	v614 = m.ExcPending
	if v614 != 0 {
		goto L1
	} else {
		goto L126
	}
L10:
	;
	v610 = *(*int32)(unsafe.Add(mBase, uint32(v267)+96))
	v611 = *(*int32)(unsafe.Add(mBase, uint32(v267)+76))
	v612 = *(*int32)(unsafe.Add(mBase, uint32(v267)+68))
	v620 = v612
	v626 = v610
	v631 = v611
	goto L8
L11:
	;
	if v531 == int32(0) {
		goto L119
	} else {
		goto L120
	}
L12:
	;
	if v286&int32(1) == int32(0) {
		v531 = v482
		goto L11
	} else {
		goto L114
	}
L13:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v460 = m.ExcPending
	if v460 != 0 {
		goto L1
	} else {
		goto L110
	}
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v438 = m.ExcPending
	if v438 != 0 {
		goto L1
	} else {
		goto L106
	}
L15:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v420 = m.ExcPending
	if v420 != 0 {
		goto L1
	} else {
		goto L102
	}
L16:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v399 = m.ExcPending
	if v399 != 0 {
		goto L1
	} else {
		goto L98
	}
L17:
	;
	if v117 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L18:
	;
	v112 = *(*int32)(unsafe.Add(mBase, _consts[113]))
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v112)+316))
	v115 = base.B2i32(v113 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _consts[112])) = uint8(v115)
	v117 = v115
	goto L20
L19:
	;
	v117 = int32(0)
	goto L20
L20:
	;
	goto L17
L21:
	;
	v121 = F_makeRangeVar(m, v100, v103, int32(-1))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L1
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L1
	} else {
		goto L93
	}
L24:
	;
	v128 = F_RangeVarGetRelidExtended(m, v121, int32(4), int32(0), int32(1060), v36+int32(604))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+256)))
	v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+248)))
	if v131 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v214 = base.I32_extend16_s(v211)
	if v214 < int32(0) {
		goto L15
	} else {
		goto L55
	}
L27:
	;
	if v130&int32(1) == int32(0) {
		goto L16
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	if v130&int32(1) == int32(0) {
		goto L38
	} else {
		goto L39
	}
L30:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v36)+244))
	v139 = F_text_to_cstring(m, v138)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	v141 = F_get_attnum(m, v128, v139)
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	if v141 != 0 {
		v211 = v141
		v213 = v139
		goto L26
	} else {
		goto L33
	}
L33:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	F_errcode(m, int32(50360452))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+180)) = v103
	*(*int32)(unsafe.Add(mBase, uint32(v36)+176)) = v139
	F_errmsg(m, int32(71628), v36+int32(176))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	F_errfinish(m, int32(493938), int32(213), int32(356171))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L38:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v36)+252))
	v167 = base.I32_extend16_s(v166)
	v169 = F_get_attname(m, v128, v167, int32(1))
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L1
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L1
	} else {
		goto L51
	}
L41:
	;
	if v169 != 0 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v171 = F_SearchSysCacheExistsAttName(m, v128, v169)
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L1
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L1
	} else {
		goto L47
	}
L45:
	;
	if v171 != 0 {
		v211 = v166
		v213 = v169
		goto L26
	} else {
		goto L46
	}
L46:
	;
	goto L44
L47:
	;
	F_errcode(m, int32(50360452))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+20)) = v103
	*(*int32)(unsafe.Add(mBase, uint32(v36)+16)) = v167
	F_errmsg(m, int32(71475), v36+int32(16))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	F_errfinish(m, int32(493938), int32(225), int32(356171))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L51:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+4)) = int32(287342)
	*(*int32)(unsafe.Add(mBase, uint32(v36))) = int32(377907)
	F_errmsg(m, int32(700726), v36)
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	F_errfinish(m, int32(493938), int32(231), int32(356171))
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
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
	F_stats_check_required_arg(m, v36+int32(208), int32(4121808), int32(4))
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v36)+260))
	v227 = F_stats_check_arg_array(m, v36+int32(208), int32(9))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	v232 = F_stats_check_arg_array(m, v36+int32(208), int32(13))
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	v237 = F_stats_check_arg_array(m, v36+int32(208), int32(14))
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	v243 = F_stats_check_arg_pair(m, v36+int32(208), int32(8), int32(9))
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	v249 = F_stats_check_arg_pair(m, v36+int32(208), int32(12), int32(13))
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	v255 = F_stats_check_arg_pair(m, v36+int32(208), int32(15), int32(16))
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	v258 = F_relation_open(m, v128, int32(1))
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	v261 = F_SearchSysCache2(m, int32(7), v128, v214)
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	if v261 == int32(0) {
		goto L14
	} else {
		goto L65
	}
L65:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v261)+16))
	v266 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v265)+22)))
	v267 = v265 + v266
	v268 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v267)+91)))
	if v268 == int32(1) {
		goto L13
	} else {
		goto L66
	}
L66:
	;
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v258)+48))
	v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v271)+119)))
	if v272|int32(32) != int32(105) {
		goto L10
	} else {
		goto L67
	}
L67:
	;
	v277 = int32(*(*int16)(unsafe.Add(mBase, uint32(v267)+74)))
	v278 = F_RelationGetIndexExpressions(m, v258)
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L1
	} else {
		goto L68
	}
L68:
	;
	if v278 == int32(0) {
		goto L10
	} else {
		goto L69
	}
L69:
	;
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v258)+192))
	v284 = v282 + int32(48)
	v285 = int32(1)
	v286 = v277 - v285
	v290 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v284+v286<<(uint(v285)%32)))))
	if v290 != 0 {
		goto L10
	} else {
		goto L70
	}
L70:
	;
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v258)+228))
	if v291 != 0 {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v291)+12))
	v294 = v292
	goto L73
L72:
	;
	v294 = int32(0)
	goto L73
L73:
	;
	if v277 < int32(2) {
		v531 = v294
		goto L11
	} else {
		goto L74
	}
L74:
	;
	if v277 == int32(2) {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v479 = int32(0)
	v482 = v294
	goto L12
L76:
	;
	goto L77
L77:
	;
	v304 = int32(0)
	v306 = v304
	v309 = v294
	v312 = v304
	goto L78
L78:
	;
	v341 = v284 + v306<<(uint(int32(1))%32)
	v342 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v341))))
	if v342 == int32(0) {
		goto L80
	} else {
		goto L81
	}
L79:
	;
	v479 = v372
	v482 = v369
	goto L12
L80:
	;
	v346 = v309 + int32(4)
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v291)+12))
	v349 = *(*int32)(unsafe.Add(mBase, uint32(v291)+4))
	if base.Ui32(v346) < base.Ui32(v348+v349<<(uint(int32(2))%32)) {
		goto L83
	} else {
		goto L84
	}
L81:
	;
	v355 = v309
	goto L82
L82:
	;
	v356 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v341)+2)))
	if v356 == int32(0) {
		goto L86
	} else {
		goto L87
	}
L83:
	;
	v354 = v346
	goto L85
L84:
	;
	v354 = int32(0)
	goto L85
L85:
	;
	v355 = v354
	goto L82
L86:
	;
	v360 = v355 + int32(4)
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v291)+12))
	v363 = *(*int32)(unsafe.Add(mBase, uint32(v291)+4))
	if base.Ui32(v360) < base.Ui32(v362+v363<<(uint(int32(2))%32)) {
		goto L89
	} else {
		goto L90
	}
L87:
	;
	v369 = v355
	goto L88
L88:
	;
	v371 = int32(2)
	v372 = v306 + v371
	v374 = v312 + v371
	if v286&int32(-2) != v374 {
		v306 = v372
		v309 = v369
		v312 = v374
		goto L78
	} else {
		goto L92
	}
L89:
	;
	v368 = v360
	goto L91
L90:
	;
	v368 = int32(0)
	goto L91
L91:
	;
	v369 = v368
	goto L88
L92:
	;
	goto L79
L93:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L1
	} else {
		goto L94
	}
L94:
	;
	F_errmsg(m, int32(128133), int32(0))
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L1
	} else {
		goto L95
	}
L95:
	;
	F_errhint(m, int32(573425), int32(0))
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L1
	} else {
		goto L96
	}
L96:
	;
	F_errfinish(m, int32(493938), int32(192), int32(356171))
	mBase = m.M
	v395 = m.ExcPending
	if v395 != 0 {
		goto L1
	} else {
		goto L97
	}
L97:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L98:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		goto L1
	} else {
		goto L99
	}
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+196)) = int32(287342)
	*(*int32)(unsafe.Add(mBase, uint32(v36)+192)) = int32(377907)
	F_errmsg(m, int32(723070), v36+int32(192))
	mBase = m.M
	v411 = m.ExcPending
	if v411 != 0 {
		goto L1
	} else {
		goto L100
	}
L100:
	;
	F_errfinish(m, int32(493938), int32(205), int32(356171))
	mBase = m.M
	v416 = m.ExcPending
	if v416 != 0 {
		goto L1
	} else {
		goto L101
	}
L101:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L102:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v423 = m.ExcPending
	if v423 != 0 {
		goto L1
	} else {
		goto L103
	}
L103:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+32)) = v213
	F_errmsg(m, int32(711510), v36+int32(32))
	mBase = m.M
	v429 = m.ExcPending
	if v429 != 0 {
		goto L1
	} else {
		goto L104
	}
L104:
	;
	F_errfinish(m, int32(493938), int32(240), int32(356171))
	mBase = m.M
	v434 = m.ExcPending
	if v434 != 0 {
		goto L1
	} else {
		goto L105
	}
L105:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L106:
	;
	F_errcode(m, int32(50360452))
	mBase = m.M
	v441 = m.ExcPending
	if v441 != 0 {
		goto L1
	} else {
		goto L107
	}
L107:
	;
	v442 = *(*int32)(unsafe.Add(mBase, uint32(v258)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v36)+48)) = v214
	*(*int32)(unsafe.Add(mBase, uint32(v36)+52)) = v442 + int32(4)
	F_errmsg(m, int32(71475), v36+int32(48))
	mBase = m.M
	v451 = m.ExcPending
	if v451 != 0 {
		goto L1
	} else {
		goto L108
	}
L108:
	;
	F_errfinish(m, int32(493938), int32(597), int32(366684))
	mBase = m.M
	v456 = m.ExcPending
	if v456 != 0 {
		goto L1
	} else {
		goto L109
	}
L109:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L110:
	;
	F_errcode(m, int32(50360452))
	mBase = m.M
	v463 = m.ExcPending
	if v463 != 0 {
		goto L1
	} else {
		goto L111
	}
L111:
	;
	v464 = *(*int32)(unsafe.Add(mBase, uint32(v258)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v36)+64)) = v214
	*(*int32)(unsafe.Add(mBase, uint32(v36)+68)) = v464 + int32(4)
	F_errmsg(m, int32(71475), v36-int32(-64))
	mBase = m.M
	v473 = m.ExcPending
	if v473 != 0 {
		goto L1
	} else {
		goto L112
	}
L112:
	;
	F_errfinish(m, int32(493938), int32(605), int32(366684))
	mBase = m.M
	v478 = m.ExcPending
	if v478 != 0 {
		goto L1
	} else {
		goto L113
	}
L113:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L114:
	;
	v517 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v284+v479<<(uint(int32(1))%32)))))
	if v517 != 0 {
		v531 = v482
		goto L11
	} else {
		goto L115
	}
L115:
	;
	v519 = v482 + int32(4)
	v521 = *(*int32)(unsafe.Add(mBase, uint32(v291)+12))
	v522 = *(*int32)(unsafe.Add(mBase, uint32(v291)+4))
	if base.Ui32(v519) < base.Ui32(v521+v522<<(uint(int32(2))%32)) {
		goto L116
	} else {
		goto L117
	}
L116:
	;
	v527 = v519
	goto L118
L117:
	;
	v527 = int32(0)
	goto L118
L118:
	;
	v531 = v527
	goto L11
L119:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v566 = m.ExcPending
	if v566 != 0 {
		goto L1
	} else {
		goto L122
	}
L120:
	;
	goto L121
L121:
	;
	v576 = *(*int32)(unsafe.Add(mBase, uint32(v531)))
	if v576 != 0 {
		goto L9
	} else {
		goto L125
	}
L122:
	;
	F_errmsg_internal(m, int32(74522), int32(0))
	mBase = m.M
	v570 = m.ExcPending
	if v570 != 0 {
		goto L1
	} else {
		goto L123
	}
L123:
	;
	F_errfinish(m, int32(493938), int32(569), int32(207032))
	mBase = m.M
	v575 = m.ExcPending
	if v575 != 0 {
		goto L1
	} else {
		goto L124
	}
L124:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L125:
	;
	goto L10
L126:
	;
	v615 = F_exprTypmod(m, v576)
	mBase = m.M
	v616 = m.ExcPending
	if v616 != 0 {
		goto L1
	} else {
		goto L127
	}
L127:
	;
	v617 = *(*int32)(unsafe.Add(mBase, uint32(v267)+96))
	if v617 != 0 {
		v620 = v613
		v626 = v617
		v631 = v615
		goto L8
	} else {
		goto L128
	}
L128:
	;
	v618 = F_exprCollation(m, v576)
	mBase = m.M
	v619 = m.ExcPending
	if v619 != 0 {
		goto L1
	} else {
		goto L129
	}
L129:
	;
	v620 = v613
	v626 = v618
	v631 = v615
	goto L8
L130:
	;
	v668 = F_type_is_multirange(m, v620)
	mBase = m.M
	v669 = m.ExcPending
	if v669 != 0 {
		goto L1
	} else {
		goto L131
	}
L131:
	;
	if v668 != 0 {
		goto L132
	} else {
		goto L133
	}
L132:
	;
	v670 = F_get_multirange_range(m, v620)
	mBase = m.M
	v671 = m.ExcPending
	if v671 != 0 {
		goto L1
	} else {
		goto L135
	}
L133:
	;
	v672 = v620
	goto L134
L134:
	;
	v675 = v249 & v255 & v243 & v237 & v232 & v227
	v677 = F_lookup_type_cache(m, v672, int32(3))
	mBase = m.M
	v678 = m.ExcPending
	if v678 != 0 {
		goto L1
	} else {
		goto L136
	}
L135:
	;
	v672 = v670
	goto L134
L136:
	;
	v679 = *(*int32)(unsafe.Add(mBase, uint32(v677)+56))
	v680 = *(*int32)(unsafe.Add(mBase, uint32(v677)+52))
	v681 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v677)+13)))
	v682 = int32(0)
	F_relation_close(m, v258, v682)
	mBase = m.M
	v685 = m.ExcPending
	if v685 != 0 {
		goto L1
	} else {
		goto L137
	}
L137:
	;
	v686 = v662 | v665
	if v672 == int32(3614) {
		goto L142
	} else {
		goto L143
	}
L138:
	;
	v753 = int32(1)
	v754 = v62 ^ v753
	v755 = v255 & (v60 | v61 ^ int32(1))
	if v57&v753 != 0 {
		goto L161
	} else {
		goto L162
	}
L139:
	;
	v745 = int32(0)
	v747 = v741
	v748 = v745
	v749 = v742
	v750 = v745
	v751 = v743
	v752 = v744
	goto L138
L140:
	;
	v714 = int32(0)
	v717 = F_errstart(m, int32(19), v714)
	mBase = m.M
	v718 = m.ExcPending
	if v718 != 0 {
		goto L1
	} else {
		goto L153
	}
L141:
	;
	v708 = F_lookup_type_cache(m, v705, int32(1))
	mBase = m.M
	v709 = m.ExcPending
	if v709 != 0 {
		goto L1
	} else {
		goto L151
	}
L142:
	;
	v691 = int32(100)
	if v686&int32(1) != 0 {
		v704 = v691
		v705 = int32(25)
		goto L141
	} else {
		goto L145
	}
L143:
	;
	goto L144
L144:
	;
	if v686&int32(1) == int32(0) {
		goto L146
	} else {
		goto L147
	}
L145:
	;
	v741 = v691
	v742 = int32(0)
	v743 = int32(0)
	v744 = v675
	goto L139
L146:
	;
	v741 = v626
	v742 = v682
	v743 = int32(0)
	v744 = v675
	goto L139
L147:
	;
	goto L148
L148:
	;
	v700 = F_get_base_element_type(m, v672)
	mBase = m.M
	v701 = m.ExcPending
	if v701 != 0 {
		goto L1
	} else {
		goto L149
	}
L149:
	;
	if v700 == int32(0) {
		v711 = v626
		goto L140
	} else {
		goto L150
	}
L150:
	;
	v704 = v626
	v705 = v700
	goto L141
L151:
	;
	v710 = *(*int32)(unsafe.Add(mBase, uint32(v708)+52))
	if v710 != 0 {
		v747 = v704
		v748 = v662
		v749 = v705
		v750 = v665
		v751 = v710
		v752 = v675
		goto L138
	} else {
		goto L152
	}
L152:
	;
	v711 = v704
	goto L140
L153:
	;
	if v717 != 0 {
		goto L154
	} else {
		goto L155
	}
L154:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+160)) = v213
	F_errmsg(m, int32(711782), v36+int32(160))
	mBase = m.M
	v724 = m.ExcPending
	if v724 != 0 {
		goto L1
	} else {
		goto L157
	}
L155:
	;
	goto L156
L156:
	;
	v739 = int32(0)
	v741 = v711
	v742 = v714
	v743 = v739
	v744 = v739
	goto L139
L157:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+148)) = int32(517371)
	*(*int32)(unsafe.Add(mBase, uint32(v36)+144)) = int32(532121)
	F_errdetail(m, int32(604777), v36+int32(144))
	mBase = m.M
	v733 = m.ExcPending
	if v733 != 0 {
		goto L1
	} else {
		goto L158
	}
L158:
	;
	F_errfinish(m, int32(493938), int32(305), int32(356171))
	mBase = m.M
	v738 = m.ExcPending
	if v738 != 0 {
		goto L1
	} else {
		goto L159
	}
L159:
	;
	goto L156
L160:
	;
	v797 = v64 ^ int32(1)
	if v755 == int32(0) {
		goto L177
	} else {
		goto L178
	}
L161:
	;
	v759 = int32(0)
	if v62&int32(1) != 0 {
		v793 = v754
		v794 = v752
		v795 = v759
		goto L160
	} else {
		goto L164
	}
L162:
	;
	v762 = v753
	goto L163
L163:
	;
	if v679 != 0 {
		v793 = v754
		v794 = v752
		v795 = v762
		goto L160
	} else {
		goto L165
	}
L164:
	;
	v762 = v759
	goto L163
L165:
	;
	v763 = int32(0)
	v766 = F_errstart(m, int32(19), v763)
	mBase = m.M
	v767 = m.ExcPending
	if v767 != 0 {
		goto L1
	} else {
		goto L166
	}
L166:
	;
	if v766 != 0 {
		goto L167
	} else {
		goto L168
	}
L167:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v770 = m.ExcPending
	if v770 != 0 {
		goto L1
	} else {
		goto L170
	}
L168:
	;
	goto L169
L169:
	;
	v791 = int32(0)
	v793 = v791
	v794 = v763
	v795 = v791
	goto L160
L170:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+128)) = v213
	F_errmsg(m, int32(710432), v36+int32(128))
	mBase = m.M
	v776 = m.ExcPending
	if v776 != 0 {
		goto L1
	} else {
		goto L171
	}
L171:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+116)) = int32(529831)
	*(*int32)(unsafe.Add(mBase, uint32(v36)+112)) = int32(532220)
	F_errdetail(m, int32(604777), v36+int32(112))
	mBase = m.M
	v785 = m.ExcPending
	if v785 != 0 {
		goto L1
	} else {
		goto L172
	}
L172:
	;
	F_errfinish(m, int32(493938), int32(322), int32(356171))
	mBase = m.M
	v790 = m.ExcPending
	if v790 != 0 {
		goto L1
	} else {
		goto L173
	}
L173:
	;
	goto L169
L174:
	;
	F_fmgr_info(m, int32(750), v36+int32(576))
	mBase = m.M
	v849 = m.ExcPending
	if v849 != 0 {
		goto L1
	} else {
		goto L191
	}
L175:
	;
	v812 = int32(0)
	v815 = F_errstart(m, int32(19), v812)
	mBase = m.M
	v816 = m.ExcPending
	if v816 != 0 {
		goto L1
	} else {
		goto L183
	}
L176:
	;
	v842 = v755
	v843 = v794
	v844 = v797
	goto L174
L177:
	;
	v800 = int32(0)
	if v64&int32(1) != 0 {
		v842 = v800
		v843 = v794
		v844 = v797
		goto L174
	} else {
		goto L180
	}
L178:
	;
	goto L179
L179:
	;
	switch v681 - int32(109) {
	case 0:
		goto L176
	default:
		goto L175
	case 5:
		v842 = int32(1)
		v843 = v794
		v844 = v797
		goto L174
	}
L180:
	;
	if v681 == int32(114) {
		v842 = v800
		v843 = v794
		v844 = v797
		goto L174
	} else {
		goto L181
	}
L181:
	;
	if v681 == int32(109) {
		goto L176
	} else {
		goto L182
	}
L182:
	;
	goto L175
L183:
	;
	if v815 != 0 {
		goto L184
	} else {
		goto L185
	}
L184:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v819 = m.ExcPending
	if v819 != 0 {
		goto L1
	} else {
		goto L187
	}
L185:
	;
	goto L186
L186:
	;
	v840 = int32(0)
	v842 = v840
	v843 = v812
	v844 = v840
	goto L174
L187:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+96)) = v213
	F_errmsg(m, int32(370702), v36+int32(96))
	mBase = m.M
	v825 = m.ExcPending
	if v825 != 0 {
		goto L1
	} else {
		goto L188
	}
L188:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+84)) = int32(532150)
	*(*int32)(unsafe.Add(mBase, uint32(v36)+80)) = int32(532182)
	F_errdetail(m, int32(604777), v36+int32(80))
	mBase = m.M
	v834 = m.ExcPending
	if v834 != 0 {
		goto L1
	} else {
		goto L189
	}
L189:
	;
	F_errfinish(m, int32(493938), int32(337), int32(356171))
	mBase = m.M
	v839 = m.ExcPending
	if v839 != 0 {
		goto L1
	} else {
		goto L190
	}
L190:
	;
	goto L186
L191:
	;
	v852 = F_table_open(m, int32(2619), int32(3))
	mBase = m.M
	v853 = m.ExcPending
	if v853 != 0 {
		goto L1
	} else {
		goto L192
	}
L192:
	;
	v856 = base.B2i32(v223 != int32(0))
	v857 = F_SearchSysCache3(m, int32(65), v128, v214, v856)
	mBase = m.M
	v858 = m.ExcPending
	if v858 != 0 {
		goto L1
	} else {
		goto L194
	}
L193:
	;
	v925 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+272)))
	if v925 == int32(0) {
		goto L199
	} else {
		goto L200
	}
L194:
	;
	if v857 != 0 {
		goto L195
	} else {
		goto L196
	}
L195:
	;
	v859 = *(*int32)(unsafe.Add(mBase, uint32(v852)+52))
	F_heap_deform_tuple(m, v857, v859, v36+int32(448), v36+int32(416))
	mBase = m.M
	v865 = m.ExcPending
	if v865 != 0 {
		goto L1
	} else {
		goto L198
	}
L196:
	;
	goto L197
L197:
	;
	v866 = int32(257)
	*(*uint16)(unsafe.Add(mBase, uint32(v36)+445)) = uint16(v866)
	v868 = int64(72340172838076673)
	*(*int64)(unsafe.Add(mBase, uint32(v36)+400)) = v868
	*(*int64)(unsafe.Add(mBase, uint32(v36)+407)) = v868
	*(*int64)(unsafe.Add(mBase, uint32(v36)+437)) = v868
	*(*int64)(unsafe.Add(mBase, uint32(v36)+384)) = v868
	*(*int64)(unsafe.Add(mBase, uint32(v36)+392)) = v868
	*(*int32)(unsafe.Add(mBase, uint32(v36)+448)) = v128
	v879 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v36)+416)) = uint8(v879)
	*(*int32)(unsafe.Add(mBase, uint32(v36)+452)) = v214
	*(*int32)(unsafe.Add(mBase, uint32(v36)+456)) = v856
	*(*int32)(unsafe.Add(mBase, uint32(v36)+460)) = v879
	*(*int32)(unsafe.Add(mBase, uint32(v36)+417)) = v879
	v887 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v36)+464)) = v887
	*(*int32)(unsafe.Add(mBase, uint32(v36)+472)) = v879
	*(*uint16)(unsafe.Add(mBase, uint32(v36)+421)) = uint16(v879)
	*(*int32)(unsafe.Add(mBase, uint32(v36)+492)) = v879
	*(*uint8)(unsafe.Add(mBase, uint32(v36)+427)) = uint8(v879)
	*(*int64)(unsafe.Add(mBase, uint32(v36)+429)) = v887
	*(*int64)(unsafe.Add(mBase, uint32(v36)+512)) = v887
	*(*int32)(unsafe.Add(mBase, uint32(v36)+476)) = v879
	*(*uint8)(unsafe.Add(mBase, uint32(v36)+423)) = uint8(v879)
	*(*int32)(unsafe.Add(mBase, uint32(v36)+496)) = v879
	*(*uint8)(unsafe.Add(mBase, uint32(v36)+428)) = uint8(v879)
	*(*uint16)(unsafe.Add(mBase, uint32(v36)+424)) = uint16(v879)
	*(*int64)(unsafe.Add(mBase, uint32(v36)+520)) = v887
	*(*int64)(unsafe.Add(mBase, uint32(v36)+500)) = v887
	*(*int64)(unsafe.Add(mBase, uint32(v36)+480)) = v887
	*(*int32)(unsafe.Add(mBase, uint32(v36)+488)) = v879
	*(*uint8)(unsafe.Add(mBase, uint32(v36)+426)) = uint8(v879)
	*(*int32)(unsafe.Add(mBase, uint32(v36)+508)) = v879
	*(*int32)(unsafe.Add(mBase, uint32(v36)+528)) = v879
	goto L193
L198:
	;
	goto L193
L199:
	;
	v928 = *(*int32)(unsafe.Add(mBase, uint32(v36)+268))
	*(*int32)(unsafe.Add(mBase, uint32(v36)+460)) = v928
	v930 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v36)+387)) = uint8(v930)
	goto L201
L200:
	;
	goto L201
L201:
	;
	v932 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+280)))
	if v932 == int32(0) {
		goto L202
	} else {
		goto L203
	}
L202:
	;
	v935 = *(*int32)(unsafe.Add(mBase, uint32(v36)+276))
	*(*int32)(unsafe.Add(mBase, uint32(v36)+464)) = v935
	v937 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v36)+388)) = uint8(v937)
	goto L204
L203:
	;
	goto L204
L204:
	;
	v940 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+288)))
	if v940 == int32(0) {
		goto L205
	} else {
		goto L206
	}
L205:
	;
	v943 = *(*int32)(unsafe.Add(mBase, uint32(v36)+284))
	*(*int32)(unsafe.Add(mBase, uint32(v36)+468)) = v943
	v945 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v36)+389)) = uint8(v945)
	goto L207
L206:
	;
	goto L207
L207:
	;
	if v227&v243 == int32(0) {
		goto L209
	} else {
		goto L210
	}
L208:
	;
	if v795 != 0 {
		goto L217
	} else {
		goto L218
	}
L209:
	;
	v981 = v843
	goto L208
L210:
	;
	if v55&int32(1) != 0 {
		goto L209
	} else {
		goto L211
	}
L211:
	;
	if v56&int32(1) != 0 {
		goto L209
	} else {
		goto L212
	}
L212:
	;
	v953 = *(*int32)(unsafe.Add(mBase, uint32(v36)+300))
	v957 = *(*int32)(unsafe.Add(mBase, uint32(v36)+292))
	v960 = F_text_to_stavalues(m, int32(153479), v36+int32(576), v957, v672, v631, v36+int32(380))
	mBase = m.M
	v961 = m.ExcPending
	if v961 != 0 {
		goto L1
	} else {
		goto L213
	}
L213:
	;
	v963 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+380)))
	if v963 != int32(1) {
		v981 = int32(0)
		goto L208
	} else {
		goto L214
	}
L214:
	;
	v973 = int32(0)
	F_set_stats_slot(m, v36+int32(448), v36+int32(416), v36+int32(384), int32(1), v680, v747, v953, v973, v960, v973)
	mBase = m.M
	v976 = m.ExcPending
	if v976 != 0 {
		goto L1
	} else {
		goto L215
	}
L215:
	;
	goto L209
L216:
	;
	if v793&int32(1) != 0 {
		goto L223
	} else {
		goto L224
	}
L217:
	;
	v982 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v36)+380)) = uint8(v982)
	v988 = *(*int32)(unsafe.Add(mBase, uint32(v36)+308))
	v991 = F_text_to_stavalues(m, int32(172300), v36+int32(576), v988, v672, v631, v36+int32(380))
	mBase = m.M
	v992 = m.ExcPending
	if v992 != 0 {
		goto L1
	} else {
		goto L220
	}
L218:
	;
	goto L219
L219:
	;
	v1010 = v981
	goto L216
L220:
	;
	v993 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+380)))
	if v993 != int32(1) {
		v1010 = v982
		goto L216
	} else {
		goto L221
	}
L221:
	;
	v1003 = int32(0)
	F_set_stats_slot(m, v36+int32(448), v36+int32(416), v36+int32(384), int32(2), v679, v747, v1003, int32(1), v991, v1003)
	mBase = m.M
	v1007 = m.ExcPending
	if v1007 != 0 {
		goto L1
	} else {
		goto L222
	}
L222:
	;
	goto L219
L223:
	;
	v1014 = *(*int32)(unsafe.Add(mBase, uint32(v36)+316))
	*(*int32)(unsafe.Add(mBase, uint32(v36)+380)) = v1014
	v1027 = F_construct_array_builtin(m, v36+int32(380), int32(1), int32(700))
	mBase = m.M
	v1028 = m.ExcPending
	if v1028 != 0 {
		goto L1
	} else {
		goto L226
	}
L224:
	;
	goto L225
L225:
	;
	if v748 != 0 {
		goto L229
	} else {
		goto L230
	}
L226:
	;
	v1029 = int32(0)
	F_set_stats_slot(m, v36+int32(448), v36+int32(416), v36+int32(384), int32(3), v679, v747, v1027, v1029, v1029, int32(1))
	mBase = m.M
	v1033 = m.ExcPending
	if v1033 != 0 {
		goto L1
	} else {
		goto L227
	}
L227:
	;
	goto L225
L228:
	;
	if v750 != 0 {
		goto L235
	} else {
		goto L236
	}
L229:
	;
	v1034 = *(*int32)(unsafe.Add(mBase, uint32(v36)+332))
	v1035 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v36)+380)) = uint8(v1035)
	v1041 = *(*int32)(unsafe.Add(mBase, uint32(v36)+324))
	v1044 = F_text_to_stavalues(m, int32(150915), v36+int32(576), v1041, v749, v631, v36+int32(380))
	mBase = m.M
	v1045 = m.ExcPending
	if v1045 != 0 {
		goto L1
	} else {
		goto L232
	}
L230:
	;
	goto L231
L231:
	;
	v1063 = v1010
	goto L228
L232:
	;
	v1046 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+380)))
	if v1046 != int32(1) {
		v1063 = v1035
		goto L228
	} else {
		goto L233
	}
L233:
	;
	v1056 = int32(0)
	F_set_stats_slot(m, v36+int32(448), v36+int32(416), v36+int32(384), int32(4), v751, v747, v1034, v1056, v1044, v1056)
	mBase = m.M
	v1059 = m.ExcPending
	if v1059 != 0 {
		goto L1
	} else {
		goto L234
	}
L234:
	;
	goto L231
L235:
	;
	v1073 = *(*int32)(unsafe.Add(mBase, uint32(v36)+340))
	v1074 = int32(0)
	F_set_stats_slot(m, v36+int32(448), v36+int32(416), v36+int32(384), int32(5), v751, v747, v1073, v1074, v1074, int32(1))
	mBase = m.M
	v1078 = m.ExcPending
	if v1078 != 0 {
		goto L1
	} else {
		goto L238
	}
L236:
	;
	goto L237
L237:
	;
	if v844&int32(1) != 0 {
		goto L240
	} else {
		goto L241
	}
L238:
	;
	goto L237
L239:
	;
	if v842 != 0 {
		goto L247
	} else {
		goto L248
	}
L240:
	;
	v1081 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v36)+380)) = uint8(v1081)
	v1087 = *(*int32)(unsafe.Add(mBase, uint32(v36)+364))
	v1090 = F_text_to_stavalues(m, int32(291964), v36+int32(576), v1087, v672, v631, v36+int32(380))
	mBase = m.M
	v1091 = m.ExcPending
	if v1091 != 0 {
		goto L1
	} else {
		goto L243
	}
L241:
	;
	goto L242
L242:
	;
	v1112 = v1063
	goto L239
L243:
	;
	v1092 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+380)))
	if v1092 != int32(1) {
		v1112 = v1081
		goto L239
	} else {
		goto L244
	}
L244:
	;
	v1102 = int32(0)
	F_set_stats_slot(m, v36+int32(448), v36+int32(416), v36+int32(384), int32(7), v1102, v1102, v1102, int32(1), v1090, v1102)
	mBase = m.M
	v1108 = m.ExcPending
	if v1108 != 0 {
		goto L1
	} else {
		goto L245
	}
L245:
	;
	goto L242
L246:
	;
	v1156 = *(*int32)(unsafe.Add(mBase, uint32(v852)+52))
	if v857 != 0 {
		goto L255
	} else {
		goto L256
	}
L247:
	;
	v1113 = *(*int32)(unsafe.Add(mBase, uint32(v36)+356))
	*(*int32)(unsafe.Add(mBase, uint32(v36)+380)) = v1113
	v1119 = F_construct_array_builtin(m, v36+int32(380), int32(1), int32(700))
	mBase = m.M
	v1120 = m.ExcPending
	if v1120 != 0 {
		goto L1
	} else {
		goto L250
	}
L248:
	;
	goto L249
L249:
	;
	v1154 = v1112
	goto L246
L250:
	;
	v1121 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v36)+379)) = uint8(v1121)
	v1127 = *(*int32)(unsafe.Add(mBase, uint32(v36)+348))
	v1132 = F_text_to_stavalues(m, int32(291987), v36+int32(576), v1127, int32(701), v1121, v36+int32(379))
	mBase = m.M
	v1133 = m.ExcPending
	if v1133 != 0 {
		goto L1
	} else {
		goto L251
	}
L251:
	;
	v1134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+379)))
	if v1134 != int32(1) {
		v1154 = v1121
		goto L246
	} else {
		goto L252
	}
L252:
	;
	v1145 = int32(0)
	F_set_stats_slot(m, v36+int32(448), v36+int32(416), v36+int32(384), int32(6), int32(672), v1145, v1119, v1145, v1132, v1145)
	mBase = m.M
	v1149 = m.ExcPending
	if v1149 != 0 {
		goto L1
	} else {
		goto L253
	}
L253:
	;
	goto L249
L254:
	;
	F_pfree(m, v1177)
	mBase = m.M
	v1179 = m.ExcPending
	if v1179 != 0 {
		goto L1
	} else {
		goto L262
	}
L255:
	;
	v1163 = F_heap_modify_tuple(m, v857, v1156, v36+int32(448), v36+int32(416), v36+int32(384))
	mBase = m.M
	v1164 = m.ExcPending
	if v1164 != 0 {
		goto L1
	} else {
		goto L258
	}
L256:
	;
	goto L257
L257:
	;
	v1173 = F_heap_form_tuple(m, v1156, v36+int32(448), v36+int32(416))
	mBase = m.M
	v1174 = m.ExcPending
	if v1174 != 0 {
		goto L1
	} else {
		goto L260
	}
L258:
	;
	F_CatalogTupleUpdate(m, v852, v1163+int32(4), v1163)
	mBase = m.M
	v1168 = m.ExcPending
	if v1168 != 0 {
		goto L1
	} else {
		goto L259
	}
L259:
	;
	v1177 = v1163
	goto L254
L260:
	;
	F_CatalogTupleInsert(m, v852, v1173)
	mBase = m.M
	v1176 = m.ExcPending
	if v1176 != 0 {
		goto L1
	} else {
		goto L261
	}
L261:
	;
	v1177 = v1173
	goto L254
L262:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v1181 = m.ExcPending
	if v1181 != 0 {
		goto L1
	} else {
		goto L263
	}
L263:
	;
	if v857 != 0 {
		goto L264
	} else {
		goto L265
	}
L264:
	;
	F_ReleaseCatCache(m, v857)
	mBase = m.M
	v1183 = m.ExcPending
	if v1183 != 0 {
		goto L1
	} else {
		goto L267
	}
L265:
	;
	goto L266
L266:
	;
	F_sequence_close(m, v852, int32(3))
	mBase = m.M
	v1186 = m.ExcPending
	if v1186 != 0 {
		goto L1
	} else {
		goto L268
	}
L267:
	;
	goto L266
L268:
	;
	m.G0 = v36 + int32(608)
	return v1154 & v49
}
