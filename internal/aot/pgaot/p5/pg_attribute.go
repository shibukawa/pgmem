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
	var v18 int64
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
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
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v86 int64
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int64
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v126 int32
	_ = v126
	v13 = m.G0
	v15 = v13 - int32(16)
	m.G0 = v15
	v18 = base.I64_extend_i32_u(l0)
	v19 = F_SearchSysCache1(m, int32(57), v18)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v15 + int32(16)
	return v126
L2:
	;
	return int32(0)
L3:
	;
	if v19 == int32(0) {
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
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v19)+16))
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+22)))
	v46 = v44 + v45
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)+80))
	v48 = int32(*(*int16)(unsafe.Add(mBase, uint32(v46)+120)))
	F_ReleaseCatCache(m, v19)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L2
	} else {
		goto L14
	}
L7:
	;
	v25 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v25)
	v126 = v25
	goto L1
L8:
	;
	goto L9
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L2
	} else {
		goto L10
	}
L10:
	;
	F_errcode(m, int32(16908420))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L2
	} else {
		goto L11
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = l0
	F_errmsg(m, int32(_a_F_pg_attribute_aclcheck_all_ext_0), v15)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L2
	} else {
		goto L12
	}
L12:
	;
	F_errfinish(m, int32(_a_F_pg_attribute_aclcheck_all_ext_1), int32(3958), int32(_a_F_pg_attribute_aclcheck_all_ext_2))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
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
	v51 = int32(1)
	if v48 <= int32(0) {
		v126 = v51
		goto L1
	} else {
		goto L15
	}
L15:
	;
	v59 = int32(1)
	v60 = v51
	goto L16
L16:
	;
	v70 = F_SearchSysCache2(m, int32(7), v18, base.I64_extend16_s(base.I64_extend_i32_u(v59)))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L2
	} else {
		goto L19
	}
L17:
	;
	v126 = v114
	goto L1
L18:
	;
	v119 = base.I32_extend16_s(v59 + int32(1))
	if v119 <= v48 {
		v59 = v119
		v60 = v114
		goto L16
	} else {
		goto L41
	}
L19:
	;
	if v70 == int32(0) {
		v114 = v60
		goto L18
	} else {
		goto L20
	}
L20:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v70)+16))
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74)+22)))
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74+v75)+91)))
	if v77 == int32(1) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	F_ReleaseCatCache(m, v70)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L2
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v86 = F_SysCacheGetAttr(m, int32(7), v70, int32(22), v15+int32(15))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L2
	} else {
		goto L25
	}
L24:
	;
	v114 = v60
	goto L18
L25:
	;
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+15)))
	if v88 == int32(1) {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v111 = int32(1)
	if l3 == int32(0) {
		v126 = v111
		goto L1
	} else {
		goto L40
	}
L27:
	;
	F_ReleaseCatCache(m, v70)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L2
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v93 = base.I32_wrap_i64(v86)
	v94 = F_pg_detoast_datum(m, v93)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L2
	} else {
		goto L31
	}
L30:
	;
	goto L26
L31:
	;
	v97 = F_aclmask(m, v94, l1, v47, l2, int32(1))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L2
	} else {
		goto L32
	}
L32:
	;
	if v94 != v93 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	F_pfree(m, v94)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L2
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	F_ReleaseCatCache(m, v70)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L2
	} else {
		goto L37
	}
L36:
	;
	goto L35
L37:
	;
	if v97 == int64(0) {
		goto L26
	} else {
		goto L38
	}
L38:
	;
	v106 = int32(0)
	if l3 == int32(1) {
		v126 = v106
		goto L1
	} else {
		goto L39
	}
L39:
	;
	v114 = v106
	goto L18
L40:
	;
	v114 = v111
	goto L18
L41:
	;
	goto L17
}
func F_pg_restore_attribute_stats(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v40 int64
	_ = v40
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
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
	var v66 int32
	_ = v66
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
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v176 int32
	_ = v176
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
	var v192 int64
	_ = v192
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
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
	var v243 int64
	_ = v243
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v287 int32
	_ = v287
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v343 int32
	_ = v343
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v391 int32
	_ = v391
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v405 int32
	_ = v405
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v418 int32
	_ = v418
	var v422 int32
	_ = v422
	var v434 int32
	_ = v434
	var v466 int32
	_ = v466
	var v468 int32
	_ = v468
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v520 int32
	_ = v520
	var v562 int32
	_ = v562
	var v564 int32
	_ = v564
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v570 int32
	_ = v570
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
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v622 int32
	_ = v622
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v626 int32
	_ = v626
	var v628 int32
	_ = v628
	var v630 int32
	_ = v630
	var v637 int32
	_ = v637
	var v644 int32
	_ = v644
	var v647 int32
	_ = v647
	var v648 int32
	_ = v648
	var v655 int32
	_ = v655
	var v660 int32
	_ = v660
	var v664 int32
	_ = v664
	var v667 int32
	_ = v667
	var v668 int32
	_ = v668
	var v677 int32
	_ = v677
	var v682 int32
	_ = v682
	var v686 int32
	_ = v686
	var v690 int32
	_ = v690
	var v695 int32
	_ = v695
	var v700 int32
	_ = v700
	var v702 int32
	_ = v702
	var v703 int32
	_ = v703
	var v705 int32
	_ = v705
	var v710 int32
	_ = v710
	var v715 int32
	_ = v715
	var v716 int32
	_ = v716
	var v717 int32
	_ = v717
	var v720 int32
	_ = v720
	var v721 int32
	_ = v721
	var v727 int32
	_ = v727
	var v735 int32
	_ = v735
	var v736 int32
	_ = v736
	var v741 int32
	_ = v741
	var v742 int32
	_ = v742
	var v746 int32
	_ = v746
	var v747 int32
	_ = v747
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
	var v755 int32
	_ = v755
	var v756 int32
	_ = v756
	var v758 int32
	_ = v758
	var v759 int32
	_ = v759
	var v762 int32
	_ = v762
	var v763 int32
	_ = v763
	var v766 int32
	_ = v766
	var v772 int32
	_ = v772
	var v780 int32
	_ = v780
	var v781 int32
	_ = v781
	var v786 int32
	_ = v786
	var v787 int32
	_ = v787
	var v789 int32
	_ = v789
	var v790 int32
	_ = v790
	var v791 int32
	_ = v791
	var v792 int32
	_ = v792
	var v793 int32
	_ = v793
	var v797 int32
	_ = v797
	var v800 int32
	_ = v800
	var v803 int32
	_ = v803
	var v804 int32
	_ = v804
	var v807 int32
	_ = v807
	var v813 int32
	_ = v813
	var v821 int32
	_ = v821
	var v822 int32
	_ = v822
	var v827 int32
	_ = v827
	var v828 int32
	_ = v828
	var v830 int32
	_ = v830
	var v831 int32
	_ = v831
	var v832 int32
	_ = v832
	var v837 int32
	_ = v837
	var v840 int32
	_ = v840
	var v841 int32
	_ = v841
	var v846 int32
	_ = v846
	var v847 int32
	_ = v847
	var v848 int32
	_ = v848
	var v854 int32
	_ = v854
	var v856 int32
	_ = v856
	var v858 int32
	_ = v858
	var v860 int32
	_ = v860
	var v861 int32
	_ = v861
	var v863 int64
	_ = v863
	var v875 int32
	_ = v875
	var v885 int64
	_ = v885
	var v957 int32
	_ = v957
	var v960 int64
	_ = v960
	var v961 int32
	_ = v961
	var v965 int32
	_ = v965
	var v968 int64
	_ = v968
	var v969 int32
	_ = v969
	var v973 int32
	_ = v973
	var v976 int64
	_ = v976
	var v977 int32
	_ = v977
	var v985 int64
	_ = v985
	var v989 int64
	_ = v989
	var v990 int32
	_ = v990
	var v991 int32
	_ = v991
	var v994 int64
	_ = v994
	var v995 int32
	_ = v995
	var v997 int32
	_ = v997
	var v1001 int32
	_ = v1001
	var v1002 int32
	_ = v1002
	var v1004 int32
	_ = v1004
	var v1005 int32
	_ = v1005
	var v1006 int32
	_ = v1006
	var v1007 int32
	_ = v1007
	var v1009 int32
	_ = v1009
	var v1012 int32
	_ = v1012
	var v1013 int32
	_ = v1013
	var v1018 int32
	_ = v1018
	var v1025 int32
	_ = v1025
	var v1030 int32
	_ = v1030
	var v1038 int32
	_ = v1038
	var v1039 int32
	_ = v1039
	var v1040 int32
	_ = v1040
	var v1043 int32
	_ = v1043
	var v1048 int32
	_ = v1048
	var v1054 int32
	_ = v1054
	var v1060 int64
	_ = v1060
	var v1061 int32
	_ = v1061
	var v1062 int32
	_ = v1062
	var v1065 int64
	_ = v1065
	var v1066 int32
	_ = v1066
	var v1067 int32
	_ = v1067
	var v1077 int32
	_ = v1077
	var v1078 int32
	_ = v1078
	var v1083 int32
	_ = v1083
	var v1086 int32
	_ = v1086
	var v1090 int64
	_ = v1090
	var v1096 int32
	_ = v1096
	var v1097 int32
	_ = v1097
	var v1105 int32
	_ = v1105
	var v1106 int32
	_ = v1106
	var v1112 int32
	_ = v1112
	var v1116 int64
	_ = v1116
	var v1117 int32
	_ = v1117
	var v1123 int64
	_ = v1123
	var v1124 int32
	_ = v1124
	var v1125 int32
	_ = v1125
	var v1128 int64
	_ = v1128
	var v1129 int32
	_ = v1129
	var v1130 int32
	_ = v1130
	var v1140 int32
	_ = v1140
	var v1141 int32
	_ = v1141
	var v1142 int32
	_ = v1142
	var v1145 int32
	_ = v1145
	var v1149 int32
	_ = v1149
	var v1161 int32
	_ = v1161
	var v1162 int32
	_ = v1162
	var v1163 int64
	_ = v1163
	var v1168 int32
	_ = v1168
	var v1171 int32
	_ = v1171
	var v1174 int32
	_ = v1174
	var v1175 int32
	_ = v1175
	var v1176 int32
	_ = v1176
	var v1177 int32
	_ = v1177
	var v1178 int32
	_ = v1178
	var v1179 int32
	_ = v1179
	var v1180 int32
	_ = v1180
	var v1184 int64
	_ = v1184
	var v1185 int32
	_ = v1185
	var v1188 int64
	_ = v1188
	var v1189 int32
	_ = v1189
	var v1190 int32
	_ = v1190
	var v1200 int32
	_ = v1200
	var v1206 int32
	_ = v1206
	var v1210 int32
	_ = v1210
	var v1213 int64
	_ = v1213
	var v1219 int32
	_ = v1219
	var v1220 int32
	_ = v1220
	var v1221 int32
	_ = v1221
	var v1227 int64
	_ = v1227
	var v1232 int64
	_ = v1232
	var v1233 int32
	_ = v1233
	var v1234 int32
	_ = v1234
	var v1245 int32
	_ = v1245
	var v1250 int32
	_ = v1250
	var v1255 int32
	_ = v1255
	var v1257 int32
	_ = v1257
	var v1264 int32
	_ = v1264
	var v1265 int32
	_ = v1265
	var v1269 int32
	_ = v1269
	var v1274 int32
	_ = v1274
	var v1275 int32
	_ = v1275
	var v1277 int32
	_ = v1277
	var v1278 int32
	_ = v1278
	var v1280 int32
	_ = v1280
	var v1282 int32
	_ = v1282
	var v1284 int32
	_ = v1284
	var v1287 int32
	_ = v1287
	var v1300 int32
	_ = v1300
	var v1303 int32
	_ = v1303
	var v1307 int32
	_ = v1307
	var v1311 int32
	_ = v1311
	var v1316 int32
	_ = v1316
	var v1320 int32
	_ = v1320
	var v1323 int32
	_ = v1323
	var v1332 int32
	_ = v1332
	var v1337 int32
	_ = v1337
	var v1341 int32
	_ = v1341
	var v1344 int32
	_ = v1344
	var v1350 int32
	_ = v1350
	var v1355 int32
	_ = v1355
	v2 = int32(0)
	v40 = int64(0)
	v42 = m.G0
	v44 = v42 - int32(400)
	m.G0 = v44
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+96)) = uint8(v2)
	*(*int64)(unsafe.Add(mBase, uint32(v44)+88)) = v40
	*(*int64)(unsafe.Add(mBase, uint32(v44)+80)) = v40
	v52 = int32(18)
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+98)) = uint16(v52)
	v55 = v44 + int32(80)
	v57 = F_stats_fill_fcinfo_from_arg_pairs(m, l0, v55, int32(_a_F_pg_restore_attribute_stats_0))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v61 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+396)) = v61
	F_stats_check_required_arg(m, v55, int32(_a_F_pg_restore_attribute_stats_0), v61)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	F_stats_check_required_arg(m, v55, int32(_a_F_pg_restore_attribute_stats_0), int32(1))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v44)+104))
	v72 = F_text_to_cstring(m, v71)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v44)+120))
	v75 = F_text_to_cstring(m, v74)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v79 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_restore_attribute_stats[0])))
	if v79 == int32(1) {
		goto L10
	} else {
		goto L11
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1341 = m.ExcPending
	if v1341 != 0 {
		goto L1
	} else {
		goto L268
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1320 = m.ExcPending
	if v1320 != 0 {
		goto L1
	} else {
		goto L264
	}
L9:
	;
	if v89 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L10:
	;
	v84 = *(*int32)(unsafe.Add(mBase, _c_F_pg_restore_attribute_stats[1]))
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v84)+308))
	v87 = base.B2i32(v85 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _c_F_pg_restore_attribute_stats[0])) = uint8(v87)
	v89 = v87
	goto L12
L11:
	;
	v89 = int32(0)
	goto L12
L12:
	;
	goto L9
L13:
	;
	v93 = F_makeRangeVar(m, v72, v75, int32(-1))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L1
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1300 = m.ExcPending
	if v1300 != 0 {
		goto L1
	} else {
		goto L259
	}
L16:
	;
	v100 = F_RangeVarGetRelidExtended(m, v93, int32(4), int32(0), int32(1138), v44+int32(396))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+160)))
	v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+144)))
	if v103 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	if v182 < int32(0) {
		goto L7
	} else {
		goto L47
	}
L19:
	;
	if v102&int32(1) == int32(0) {
		goto L8
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	if v102&int32(1) == int32(0) {
		goto L30
	} else {
		goto L31
	}
L22:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v44)+136))
	v111 = F_text_to_cstring(m, v110)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	v113 = F_get_attnum(m, v100, v111)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	if v113 != 0 {
		v182 = v113
		v183 = v111
		goto L18
	} else {
		goto L25
	}
L25:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	F_errcode(m, int32(50360452))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+52)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v44)+48)) = v111
	F_errmsg(m, int32(_a_F_pg_restore_attribute_stats_15), v44+int32(48))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	F_errfinish(m, int32(_a_F_pg_restore_attribute_stats_4), int32(175), int32(_a_F_pg_restore_attribute_stats_5))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L30:
	;
	v138 = int32(*(*int16)(unsafe.Add(mBase, uint32(v44)+152)))
	v140 = F_get_attname(m, v100, v138, int32(1))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L1
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L1
	} else {
		goto L43
	}
L33:
	;
	if v140 != 0 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v142 = F_SearchSysCacheExistsAttName(m, v100, v140)
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L1
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L1
	} else {
		goto L39
	}
L37:
	;
	if v142 != 0 {
		v182 = v138
		v183 = v140
		goto L18
	} else {
		goto L38
	}
L38:
	;
	goto L36
L39:
	;
	F_errcode(m, int32(50360452))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+20)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v44)+16)) = v138
	F_errmsg(m, int32(_a_F_pg_restore_attribute_stats_7), v44+int32(16))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	F_errfinish(m, int32(_a_F_pg_restore_attribute_stats_4), int32(187), int32(_a_F_pg_restore_attribute_stats_5))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L43:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+4)) = int32(_a_F_pg_restore_attribute_stats_1)
	*(*int32)(unsafe.Add(mBase, uint32(v44))) = int32(_a_F_pg_restore_attribute_stats_2)
	F_errmsg(m, int32(_a_F_pg_restore_attribute_stats_16), v44)
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	F_errfinish(m, int32(_a_F_pg_restore_attribute_stats_4), int32(193), int32(_a_F_pg_restore_attribute_stats_5))
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L47:
	;
	v187 = v44 + int32(80)
	F_stats_check_required_arg(m, v187, int32(_a_F_pg_restore_attribute_stats_0), int32(4))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	v192 = *(*int64)(unsafe.Add(mBase, uint32(v44)+168))
	v194 = base.B2i32(v192 != int64(0))
	v195 = m.G0
	v197 = v195 - int32(496)
	m.G0 = v197
	v199 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v197)+492)) = v199
	*(*int32)(unsafe.Add(mBase, uint32(v197)+480)) = v199
	*(*int32)(unsafe.Add(mBase, uint32(v197)+476)) = v199
	*(*int32)(unsafe.Add(mBase, uint32(v197)+472)) = v199
	*(*int32)(unsafe.Add(mBase, uint32(v197)+468)) = v199
	*(*int32)(unsafe.Add(mBase, uint32(v197)+464)) = v199
	v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+176)))
	if v211 == v199 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+160)))
	v217 = v214 ^ int32(1)
	goto L51
L50:
	;
	v217 = v2
	goto L51
L51:
	;
	v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+224)))
	if v218 == int32(0) {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+240)))
	v224 = v221 ^ int32(1)
	goto L54
L53:
	;
	v224 = v2
	goto L54
L54:
	;
	v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+256)))
	v228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+272)))
	if v228 == int32(0) {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+288)))
	v234 = v231 ^ int32(1)
	goto L57
L56:
	;
	v234 = v2
	goto L57
L57:
	;
	v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+208)))
	v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+192)))
	v237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+304)))
	base.MemoryFill(m, v197+int32(176), int32(0), int32(248))
	v243 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v197)+167)) = v243
	*(*int64)(unsafe.Add(mBase, uint32(v197)+160)) = v243
	*(*int64)(unsafe.Add(mBase, uint32(v197)+152)) = v243
	*(*int64)(unsafe.Add(mBase, uint32(v197)+144)) = v243
	*(*int64)(unsafe.Add(mBase, uint32(v197)+135)) = v243
	*(*int64)(unsafe.Add(mBase, uint32(v197)+128)) = v243
	*(*int64)(unsafe.Add(mBase, uint32(v197)+120)) = v243
	*(*int64)(unsafe.Add(mBase, uint32(v197)+112)) = v243
	v260 = F_stats_check_arg_array(m, v187, int32(9))
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	v263 = F_stats_check_arg_array(m, v187, int32(13))
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	v266 = F_stats_check_arg_array(m, v187, int32(14))
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	v270 = F_stats_check_arg_pair(m, v187, int32(8), int32(9))
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	v274 = F_stats_check_arg_pair(m, v187, int32(12), int32(13))
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	v278 = F_stats_check_arg_pair(m, v187, int32(15), int32(16))
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	v281 = v197 + int32(492)
	v283 = v197 + int32(488)
	v287 = v197 + int32(480)
	v292 = m.G0
	v294 = v292 - int32(32)
	m.G0 = v294
	v297 = F_relation_open(m, v100, int32(1))
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	v302 = F_SearchSysCache2(m, int32(7), base.I64_extend_i32_u(v100), base.I64_extend_i32_s(v182))
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L1
	} else {
		goto L68
	}
L65:
	;
	v700 = v260 & (v263 & (v266 & (v270 & (v274 & v278))))
	v702 = v263 & v274 & v224
	v703 = int32(1)
	v705 = v266 & (v225 ^ int32(1))
	if v702&v703|v705&v703 != 0 {
		goto L132
	} else {
		goto L133
	}
L66:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v686 = m.ExcPending
	if v686 != 0 {
		goto L1
	} else {
		goto L128
	}
L67:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v664 = m.ExcPending
	if v664 != 0 {
		goto L1
	} else {
		goto L124
	}
L68:
	;
	if v302 != 0 {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v302)+16))
	v305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v304)+22)))
	v306 = v304 + v305
	v307 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v306)+91)))
	if v307 == int32(1) {
		goto L67
	} else {
		goto L72
	}
L70:
	;
	goto L71
L71:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v644 = m.ExcPending
	if v644 != 0 {
		goto L1
	} else {
		goto L120
	}
L72:
	;
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v297)+48))
	v311 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+119)))
	if v311|int32(32) != int32(105) {
		goto L75
	} else {
		goto L76
	}
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v287))) = v576
	F_ReleaseCatCache(m, v302)
	mBase = m.M
	v619 = m.ExcPending
	if v619 != 0 {
		goto L1
	} else {
		goto L114
	}
L74:
	;
	v567 = F_exprType(m, v520)
	mBase = m.M
	v568 = m.ExcPending
	if v568 != 0 {
		goto L1
	} else {
		goto L110
	}
L75:
	;
	v562 = *(*int32)(unsafe.Add(mBase, uint32(v306)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v281))) = v562
	v564 = *(*int32)(unsafe.Add(mBase, uint32(v306)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v283))) = v564
	v566 = *(*int32)(unsafe.Add(mBase, uint32(v306)+96))
	v576 = v566
	goto L73
L76:
	;
	v316 = int32(*(*int16)(unsafe.Add(mBase, uint32(v306)+74)))
	v317 = F_RelationGetIndexExpressions(m, v297)
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L1
	} else {
		goto L77
	}
L77:
	;
	if v317 == int32(0) {
		goto L75
	} else {
		goto L78
	}
L78:
	;
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v297)+192))
	v323 = v321 + int32(48)
	v324 = int32(1)
	v325 = v316 - v324
	v329 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v323+v325<<(uint(v324)%32)))))
	if v329 != 0 {
		goto L75
	} else {
		goto L79
	}
L79:
	;
	v330 = *(*int32)(unsafe.Add(mBase, uint32(v297)+228))
	if v330 != 0 {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v330)+12))
	v333 = v331
	goto L82
L81:
	;
	v333 = int32(0)
	goto L82
L82:
	;
	if v316 < int32(2) {
		v477 = v333
		goto L83
	} else {
		goto L84
	}
L83:
	;
	if v477 == int32(0) {
		goto L66
	} else {
		goto L108
	}
L84:
	;
	if v316 != int32(2) {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v343 = v333
	v355 = v2
	v356 = int32(0)
	goto L88
L86:
	;
	v422 = v333
	v434 = v2
	goto L87
L87:
	;
	v466 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v323+v434<<(uint(int32(1))%32)))))
	if v466 != 0 {
		v477 = v422
		goto L83
	} else {
		goto L104
	}
L88:
	;
	v386 = v323 + v355<<(uint(int32(1))%32)
	v387 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v386))))
	if v387 == int32(0) {
		goto L90
	} else {
		goto L91
	}
L89:
	;
	if v325&int32(1) == int32(0) {
		v477 = v414
		goto L83
	} else {
		goto L103
	}
L90:
	;
	v391 = v343 + int32(4)
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v330)+12))
	v394 = *(*int32)(unsafe.Add(mBase, uint32(v330)+4))
	if base.Ui32(v391) < base.Ui32(v393+v394<<(uint(int32(2))%32)) {
		goto L93
	} else {
		goto L94
	}
L91:
	;
	v400 = v343
	goto L92
L92:
	;
	v401 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v386)+2)))
	if v401 == int32(0) {
		goto L96
	} else {
		goto L97
	}
L93:
	;
	v399 = v391
	goto L95
L94:
	;
	v399 = int32(0)
	goto L95
L95:
	;
	v400 = v399
	goto L92
L96:
	;
	v405 = v400 + int32(4)
	v407 = *(*int32)(unsafe.Add(mBase, uint32(v330)+12))
	v408 = *(*int32)(unsafe.Add(mBase, uint32(v330)+4))
	if base.Ui32(v405) < base.Ui32(v407+v408<<(uint(int32(2))%32)) {
		goto L99
	} else {
		goto L100
	}
L97:
	;
	v414 = v400
	goto L98
L98:
	;
	v415 = int32(2)
	v416 = v355 + v415
	v418 = v356 + v415
	if v418 != v325&int32(-2) {
		v343 = v414
		v355 = v416
		v356 = v418
		goto L88
	} else {
		goto L102
	}
L99:
	;
	v413 = v405
	goto L101
L100:
	;
	v413 = int32(0)
	goto L101
L101:
	;
	v414 = v413
	goto L98
L102:
	;
	goto L89
L103:
	;
	v422 = v414
	v434 = v416
	goto L87
L104:
	;
	v468 = v422 + int32(4)
	v470 = *(*int32)(unsafe.Add(mBase, uint32(v330)+12))
	v471 = *(*int32)(unsafe.Add(mBase, uint32(v330)+4))
	if base.Ui32(v468) < base.Ui32(v470+v471<<(uint(int32(2))%32)) {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	v476 = v468
	goto L107
L106:
	;
	v476 = int32(0)
	goto L107
L107:
	;
	v477 = v476
	goto L83
L108:
	;
	v520 = *(*int32)(unsafe.Add(mBase, uint32(v477)))
	if v520 != 0 {
		goto L74
	} else {
		goto L109
	}
L109:
	;
	goto L75
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v281))) = v567
	v570 = F_exprTypmod(m, v520)
	mBase = m.M
	v571 = m.ExcPending
	if v571 != 0 {
		goto L1
	} else {
		goto L111
	}
L111:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v283))) = v570
	v573 = *(*int32)(unsafe.Add(mBase, uint32(v306)+96))
	if v573 != 0 {
		v576 = v573
		goto L73
	} else {
		goto L112
	}
L112:
	;
	v574 = F_exprCollation(m, v520)
	mBase = m.M
	v575 = m.ExcPending
	if v575 != 0 {
		goto L1
	} else {
		goto L113
	}
L113:
	;
	v576 = v574
	goto L73
L114:
	;
	v620 = *(*int32)(unsafe.Add(mBase, uint32(v281)))
	v622 = F_lookup_type_cache(m, v620, int32(3))
	mBase = m.M
	v623 = m.ExcPending
	if v623 != 0 {
		goto L1
	} else {
		goto L115
	}
L115:
	;
	v624 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v622)+13)))
	*(*uint8)(unsafe.Add(mBase, uint32(v197+int32(487)))) = uint8(v624)
	v626 = *(*int32)(unsafe.Add(mBase, uint32(v622)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v197+int32(476)))) = v626
	v628 = *(*int32)(unsafe.Add(mBase, uint32(v622)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v197+int32(472)))) = v628
	v630 = *(*int32)(unsafe.Add(mBase, uint32(v281)))
	if v630 == int32(3614) {
		goto L116
	} else {
		goto L117
	}
L116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v287))) = int32(100)
	goto L118
L117:
	;
	goto L118
L118:
	;
	F_relation_close(m, v297, int32(0))
	mBase = m.M
	v637 = m.ExcPending
	if v637 != 0 {
		goto L1
	} else {
		goto L119
	}
L119:
	;
	m.G0 = v294 + int32(32)
	goto L65
L120:
	;
	F_errcode(m, int32(50360452))
	mBase = m.M
	v647 = m.ExcPending
	if v647 != 0 {
		goto L1
	} else {
		goto L121
	}
L121:
	;
	v648 = *(*int32)(unsafe.Add(mBase, uint32(v297)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v294))) = v182
	*(*int32)(unsafe.Add(mBase, uint32(v294)+4)) = v648 + int32(4)
	F_errmsg(m, int32(_a_F_pg_restore_attribute_stats_7), v294)
	mBase = m.M
	v655 = m.ExcPending
	if v655 != 0 {
		goto L1
	} else {
		goto L122
	}
L122:
	;
	F_errfinish(m, int32(_a_F_pg_restore_attribute_stats_8), int32(461), int32(_a_F_pg_restore_attribute_stats_9))
	mBase = m.M
	v660 = m.ExcPending
	if v660 != 0 {
		goto L1
	} else {
		goto L123
	}
L123:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L124:
	;
	F_errcode(m, int32(50360452))
	mBase = m.M
	v667 = m.ExcPending
	if v667 != 0 {
		goto L1
	} else {
		goto L125
	}
L125:
	;
	v668 = *(*int32)(unsafe.Add(mBase, uint32(v297)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v294)+16)) = v182
	*(*int32)(unsafe.Add(mBase, uint32(v294)+20)) = v668 + int32(4)
	F_errmsg(m, int32(_a_F_pg_restore_attribute_stats_7), v294+int32(16))
	mBase = m.M
	v677 = m.ExcPending
	if v677 != 0 {
		goto L1
	} else {
		goto L126
	}
L126:
	;
	F_errfinish(m, int32(_a_F_pg_restore_attribute_stats_8), int32(469), int32(_a_F_pg_restore_attribute_stats_9))
	mBase = m.M
	v682 = m.ExcPending
	if v682 != 0 {
		goto L1
	} else {
		goto L127
	}
L127:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L128:
	;
	F_errmsg_internal(m, int32(_a_F_pg_restore_attribute_stats_17), int32(0))
	mBase = m.M
	v690 = m.ExcPending
	if v690 != 0 {
		goto L1
	} else {
		goto L129
	}
L129:
	;
	F_errfinish(m, int32(_a_F_pg_restore_attribute_stats_8), int32(335), int32(_a_F_pg_restore_attribute_stats_18))
	mBase = m.M
	v695 = m.ExcPending
	if v695 != 0 {
		goto L1
	} else {
		goto L130
	}
L130:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L131:
	;
	v752 = int32(1)
	v753 = v235 ^ v752
	v755 = v236 ^ v752
	v756 = v234 & v278
	if v235&v236 != 0 {
		v789 = v749
		v790 = v755
		v791 = v753
		goto L144
	} else {
		goto L145
	}
L132:
	;
	v710 = *(*int32)(unsafe.Add(mBase, uint32(v197)+492))
	v715 = F_statatt_get_elem_type(m, v710, v197+int32(468), v197+int32(464))
	mBase = m.M
	v716 = m.ExcPending
	if v716 != 0 {
		goto L1
	} else {
		goto L135
	}
L133:
	;
	v746 = v700
	goto L134
L134:
	;
	v747 = int32(0)
	v749 = v746
	v750 = v747
	v751 = v747
	goto L131
L135:
	;
	if v715 != 0 {
		v749 = v700
		v750 = v705
		v751 = v702
		goto L131
	} else {
		goto L136
	}
L136:
	;
	v717 = int32(0)
	v720 = F_errstart(m, int32(19), v717)
	mBase = m.M
	v721 = m.ExcPending
	if v721 != 0 {
		goto L1
	} else {
		goto L137
	}
L137:
	;
	if v720 != 0 {
		goto L138
	} else {
		goto L139
	}
L138:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v197)+96)) = v183
	F_errmsg(m, int32(_a_F_pg_restore_attribute_stats_19), v197+int32(96))
	mBase = m.M
	v727 = m.ExcPending
	if v727 != 0 {
		goto L1
	} else {
		goto L141
	}
L139:
	;
	goto L140
L140:
	;
	v742 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v197)+464)) = v742
	*(*int32)(unsafe.Add(mBase, uint32(v197)+468)) = v742
	v746 = v717
	goto L134
L141:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v197)+84)) = int32(_a_F_pg_restore_attribute_stats_20)
	*(*int32)(unsafe.Add(mBase, uint32(v197)+80)) = int32(_a_F_pg_restore_attribute_stats_21)
	v735 = F_errdetail(m, int32(_a_F_pg_restore_attribute_stats_22), v197+int32(80))
	mBase = m.M
	v736 = m.ExcPending
	if v736 != 0 {
		goto L1
	} else {
		goto L142
	}
L142:
	;
	F_errfinish(m, int32(_a_F_pg_restore_attribute_stats_4), int32(311), int32(_a_F_pg_restore_attribute_stats_23))
	mBase = m.M
	v741 = m.ExcPending
	if v741 != 0 {
		goto L1
	} else {
		goto L143
	}
L143:
	;
	goto L140
L144:
	;
	v792 = int32(1)
	v793 = v237 ^ v792
	if (v756^v792)&v237 != 0 {
		v830 = v789
		v831 = v756
		v832 = v793
		goto L155
	} else {
		goto L156
	}
L145:
	;
	v758 = *(*int32)(unsafe.Add(mBase, uint32(v197)+472))
	if v758 != 0 {
		v789 = v749
		v790 = v755
		v791 = v753
		goto L144
	} else {
		goto L146
	}
L146:
	;
	v759 = int32(0)
	v762 = F_errstart(m, int32(19), v759)
	mBase = m.M
	v763 = m.ExcPending
	if v763 != 0 {
		goto L1
	} else {
		goto L147
	}
L147:
	;
	if v762 != 0 {
		goto L148
	} else {
		goto L149
	}
L148:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v766 = m.ExcPending
	if v766 != 0 {
		goto L1
	} else {
		goto L151
	}
L149:
	;
	goto L150
L150:
	;
	v787 = int32(0)
	v789 = v759
	v790 = v787
	v791 = v787
	goto L144
L151:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v197)+64)) = v183
	F_errmsg(m, int32(_a_F_pg_restore_attribute_stats_24), v197-int32(-64))
	mBase = m.M
	v772 = m.ExcPending
	if v772 != 0 {
		goto L1
	} else {
		goto L152
	}
L152:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v197)+52)) = int32(_a_F_pg_restore_attribute_stats_25)
	*(*int32)(unsafe.Add(mBase, uint32(v197)+48)) = int32(_a_F_pg_restore_attribute_stats_26)
	v780 = F_errdetail(m, int32(_a_F_pg_restore_attribute_stats_22), v197+int32(48))
	mBase = m.M
	v781 = m.ExcPending
	if v781 != 0 {
		goto L1
	} else {
		goto L153
	}
L153:
	;
	F_errfinish(m, int32(_a_F_pg_restore_attribute_stats_4), int32(328), int32(_a_F_pg_restore_attribute_stats_23))
	mBase = m.M
	v786 = m.ExcPending
	if v786 != 0 {
		goto L1
	} else {
		goto L154
	}
L154:
	;
	goto L150
L155:
	;
	F_fmgr_info(m, int32(750), v197+int32(436))
	mBase = m.M
	v837 = m.ExcPending
	if v837 != 0 {
		goto L1
	} else {
		goto L166
	}
L156:
	;
	v797 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v197)+487)))
	switch v797 - int32(109) {
	case 0, 5:
		v830 = v789
		v831 = v756
		v832 = v793
		goto L155
	default:
		goto L157
	}
L157:
	;
	v800 = int32(0)
	v803 = F_errstart(m, int32(19), v800)
	mBase = m.M
	v804 = m.ExcPending
	if v804 != 0 {
		goto L1
	} else {
		goto L158
	}
L158:
	;
	if v803 != 0 {
		goto L159
	} else {
		goto L160
	}
L159:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v807 = m.ExcPending
	if v807 != 0 {
		goto L1
	} else {
		goto L162
	}
L160:
	;
	goto L161
L161:
	;
	v828 = int32(0)
	v830 = v800
	v831 = v828
	v832 = v828
	goto L155
L162:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v197)+32)) = v183
	F_errmsg(m, int32(_a_F_pg_restore_attribute_stats_27), v197+int32(32))
	mBase = m.M
	v813 = m.ExcPending
	if v813 != 0 {
		goto L1
	} else {
		goto L163
	}
L163:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v197)+20)) = int32(_a_F_pg_restore_attribute_stats_28)
	*(*int32)(unsafe.Add(mBase, uint32(v197)+16)) = int32(_a_F_pg_restore_attribute_stats_29)
	v821 = F_errdetail(m, int32(_a_F_pg_restore_attribute_stats_22), v197+int32(16))
	mBase = m.M
	v822 = m.ExcPending
	if v822 != 0 {
		goto L1
	} else {
		goto L164
	}
L164:
	;
	F_errfinish(m, int32(_a_F_pg_restore_attribute_stats_4), int32(343), int32(_a_F_pg_restore_attribute_stats_23))
	mBase = m.M
	v827 = m.ExcPending
	if v827 != 0 {
		goto L1
	} else {
		goto L165
	}
L165:
	;
	goto L161
L166:
	;
	v840 = F_table_open(m, int32(2619), int32(3))
	mBase = m.M
	v841 = m.ExcPending
	if v841 != 0 {
		goto L1
	} else {
		goto L167
	}
L167:
	;
	v846 = F_SearchSysCache3(m, int32(65), base.I64_extend_i32_u(v100), base.I64_extend_i32_s(v182), base.I64_extend_i32_u(v194))
	mBase = m.M
	v847 = m.ExcPending
	if v847 != 0 {
		goto L1
	} else {
		goto L169
	}
L168:
	;
	v957 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+112)))
	if v957 == int32(0) {
		goto L175
	} else {
		goto L176
	}
L169:
	;
	if v846 != 0 {
		goto L170
	} else {
		goto L171
	}
L170:
	;
	v848 = *(*int32)(unsafe.Add(mBase, uint32(v840)+52))
	F_heap_deform_tuple(m, v846, v848, v197+int32(176), v197+int32(144))
	mBase = m.M
	v854 = m.ExcPending
	if v854 != 0 {
		goto L1
	} else {
		goto L173
	}
L171:
	;
	goto L172
L172:
	;
	v856 = v197 + int32(176)
	v858 = v197 + int32(144)
	v860 = v197 + int32(112)
	v861 = int32(257)
	*(*uint16)(unsafe.Add(mBase, uint32(v858)+29)) = uint16(v861)
	v863 = int64(72340172838076673)
	*(*int64)(unsafe.Add(mBase, uint32(v858)+21)) = v863
	*(*int64)(unsafe.Add(mBase, uint32(v860)+23)) = v863
	*(*int64)(unsafe.Add(mBase, uint32(v860)+16)) = v863
	*(*int64)(unsafe.Add(mBase, uint32(v860)+8)) = v863
	*(*int64)(unsafe.Add(mBase, uint32(v860))) = v863
	*(*int64)(unsafe.Add(mBase, uint32(v856))) = base.I64_extend_i32_u(v100)
	v875 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v858))) = uint8(v875)
	*(*int64)(unsafe.Add(mBase, uint32(v856)+8)) = base.I64_extend_i32_s(v182)
	*(*uint8)(unsafe.Add(mBase, uint32(v858)+1)) = uint8(v875)
	*(*int64)(unsafe.Add(mBase, uint32(v856)+16)) = base.I64_extend_i32_u(v194)
	*(*uint8)(unsafe.Add(mBase, uint32(v858)+2)) = uint8(v875)
	v885 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v856)+24)) = v885
	*(*uint8)(unsafe.Add(mBase, uint32(v858)+3)) = uint8(v875)
	*(*int64)(unsafe.Add(mBase, uint32(v856)+32)) = v885
	*(*uint8)(unsafe.Add(mBase, uint32(v858)+4)) = uint8(v875)
	*(*int64)(unsafe.Add(mBase, uint32(v856)+40)) = v885
	*(*uint8)(unsafe.Add(mBase, uint32(v858)+5)) = uint8(v875)
	*(*int64)(unsafe.Add(mBase, uint32(v856)+48)) = v885
	*(*uint8)(unsafe.Add(mBase, uint32(v858)+6)) = uint8(v875)
	*(*int64)(unsafe.Add(mBase, uint32(v856)+88)) = v885
	*(*uint8)(unsafe.Add(mBase, uint32(v858)+11)) = uint8(v875)
	*(*int64)(unsafe.Add(mBase, uint32(v856)+128)) = v885
	*(*uint8)(unsafe.Add(mBase, uint32(v858)+16)) = uint8(v875)
	*(*int64)(unsafe.Add(mBase, uint32(v856)+56)) = v885
	*(*uint8)(unsafe.Add(mBase, uint32(v858)+7)) = uint8(v875)
	*(*int64)(unsafe.Add(mBase, uint32(v856)+96)) = v885
	*(*uint8)(unsafe.Add(mBase, uint32(v858)+12)) = uint8(v875)
	*(*int64)(unsafe.Add(mBase, uint32(v856)+136)) = v885
	*(*uint8)(unsafe.Add(mBase, uint32(v858)+17)) = uint8(v875)
	*(*int64)(unsafe.Add(mBase, uint32(v856)+64)) = v885
	*(*uint8)(unsafe.Add(mBase, uint32(v858)+8)) = uint8(v875)
	*(*int64)(unsafe.Add(mBase, uint32(v856)+104)) = v885
	*(*uint8)(unsafe.Add(mBase, uint32(v858)+13)) = uint8(v875)
	*(*int64)(unsafe.Add(mBase, uint32(v856)+144)) = v885
	*(*uint8)(unsafe.Add(mBase, uint32(v858)+18)) = uint8(v875)
	*(*int64)(unsafe.Add(mBase, uint32(v856)+72)) = v885
	*(*uint8)(unsafe.Add(mBase, uint32(v858)+9)) = uint8(v875)
	*(*int64)(unsafe.Add(mBase, uint32(v856)+112)) = v885
	*(*uint8)(unsafe.Add(mBase, uint32(v858)+14)) = uint8(v875)
	*(*int64)(unsafe.Add(mBase, uint32(v856)+152)) = v885
	*(*uint8)(unsafe.Add(mBase, uint32(v858)+19)) = uint8(v875)
	*(*int64)(unsafe.Add(mBase, uint32(v856)+80)) = v885
	*(*uint8)(unsafe.Add(mBase, uint32(v858)+10)) = uint8(v875)
	*(*int64)(unsafe.Add(mBase, uint32(v856)+120)) = v885
	*(*uint8)(unsafe.Add(mBase, uint32(v858)+15)) = uint8(v875)
	*(*int64)(unsafe.Add(mBase, uint32(v856)+160)) = v885
	*(*uint8)(unsafe.Add(mBase, uint32(v858)+20)) = uint8(v875)
	goto L174
L173:
	;
	goto L168
L174:
	;
	goto L168
L175:
	;
	v960 = *(*int64)(unsafe.Add(mBase, uint32(v187)+104))
	v961 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v197)+115)) = uint8(v961)
	*(*int64)(unsafe.Add(mBase, uint32(v197)+200)) = v960
	goto L177
L176:
	;
	goto L177
L177:
	;
	v965 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+128)))
	if v965 == int32(0) {
		goto L178
	} else {
		goto L179
	}
L178:
	;
	v968 = *(*int64)(unsafe.Add(mBase, uint32(v187)+120))
	v969 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v197)+116)) = uint8(v969)
	*(*int64)(unsafe.Add(mBase, uint32(v197)+208)) = v968
	goto L180
L179:
	;
	goto L180
L180:
	;
	v973 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+144)))
	if v973 == int32(0) {
		goto L181
	} else {
		goto L182
	}
L181:
	;
	v976 = *(*int64)(unsafe.Add(mBase, uint32(v187)+136))
	v977 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v197)+117)) = uint8(v977)
	*(*int64)(unsafe.Add(mBase, uint32(v197)+216)) = v976
	goto L183
L182:
	;
	goto L183
L183:
	;
	if v260&v270&v217&int32(1) != 0 {
		goto L185
	} else {
		goto L186
	}
L184:
	;
	if v790&int32(1) != 0 {
		goto L202
	} else {
		goto L203
	}
L185:
	;
	v985 = *(*int64)(unsafe.Add(mBase, uint32(v187)+168))
	v989 = *(*int64)(unsafe.Add(mBase, uint32(v187)+152))
	v990 = *(*int32)(unsafe.Add(mBase, uint32(v197)+492))
	v991 = *(*int32)(unsafe.Add(mBase, uint32(v197)+488))
	v994 = F_statatt_build_stavalues(m, int32(_a_F_pg_restore_attribute_stats_10), v197+int32(436), v989, v990, v991, v197+int32(104))
	mBase = m.M
	v995 = m.ExcPending
	if v995 != 0 {
		goto L1
	} else {
		goto L188
	}
L186:
	;
	goto L187
L187:
	;
	v1048 = v830
	goto L184
L188:
	;
	v997 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v197)+104)))
	if v997 != int32(1) {
		v1048 = int32(0)
		goto L184
	} else {
		goto L189
	}
L189:
	;
	v1001 = F_pg_detoast_datum(m, base.I32_wrap_i64(v994))
	mBase = m.M
	v1002 = m.ExcPending
	if v1002 != 0 {
		goto L1
	} else {
		goto L190
	}
L190:
	;
	v1004 = F_pg_detoast_datum(m, base.I32_wrap_i64(v985))
	mBase = m.M
	v1005 = m.ExcPending
	if v1005 != 0 {
		goto L1
	} else {
		goto L191
	}
L191:
	;
	v1006 = *(*int32)(unsafe.Add(mBase, uint32(v1001)+16))
	v1007 = *(*int32)(unsafe.Add(mBase, uint32(v1004)+16))
	if v1006 != v1007 {
		goto L192
	} else {
		goto L193
	}
L192:
	;
	v1009 = int32(0)
	v1012 = F_errstart(m, int32(19), v1009)
	mBase = m.M
	v1013 = m.ExcPending
	if v1013 != 0 {
		goto L1
	} else {
		goto L195
	}
L193:
	;
	goto L194
L194:
	;
	v1038 = *(*int32)(unsafe.Add(mBase, uint32(v197)+476))
	v1039 = *(*int32)(unsafe.Add(mBase, uint32(v197)+480))
	v1040 = int32(0)
	F_statatt_set_slot(m, v197+int32(176), v197+int32(144), v197+int32(112), int32(1), v1038, v1039, v985, v1040, v994, v1040)
	mBase = m.M
	v1043 = m.ExcPending
	if v1043 != 0 {
		goto L1
	} else {
		goto L200
	}
L195:
	;
	if v1012 == int32(0) {
		v1048 = v1009
		goto L184
	} else {
		goto L196
	}
L196:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v1018 = m.ExcPending
	if v1018 != 0 {
		goto L1
	} else {
		goto L197
	}
L197:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v197)+4)) = int32(_a_F_pg_restore_attribute_stats_30)
	*(*int32)(unsafe.Add(mBase, uint32(v197))) = int32(_a_F_pg_restore_attribute_stats_10)
	F_errmsg(m, int32(_a_F_pg_restore_attribute_stats_31), v197)
	mBase = m.M
	v1025 = m.ExcPending
	if v1025 != 0 {
		goto L1
	} else {
		goto L198
	}
L198:
	;
	F_errfinish(m, int32(_a_F_pg_restore_attribute_stats_4), int32(404), int32(_a_F_pg_restore_attribute_stats_23))
	mBase = m.M
	v1030 = m.ExcPending
	if v1030 != 0 {
		goto L1
	} else {
		goto L199
	}
L199:
	;
	v1048 = v1009
	goto L184
L200:
	;
	goto L187
L201:
	;
	if v791&int32(1) != 0 {
		goto L208
	} else {
		goto L209
	}
L202:
	;
	v1054 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v197)+104)) = uint8(v1054)
	v1060 = *(*int64)(unsafe.Add(mBase, uint32(v187)+184))
	v1061 = *(*int32)(unsafe.Add(mBase, uint32(v197)+492))
	v1062 = *(*int32)(unsafe.Add(mBase, uint32(v197)+488))
	v1065 = F_statatt_build_stavalues(m, int32(_a_F_pg_restore_attribute_stats_11), v197+int32(436), v1060, v1061, v1062, v197+int32(104))
	mBase = m.M
	v1066 = m.ExcPending
	if v1066 != 0 {
		goto L1
	} else {
		goto L205
	}
L203:
	;
	goto L204
L204:
	;
	v1086 = v1048
	goto L201
L205:
	;
	v1067 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v197)+104)))
	if v1067 != int32(1) {
		v1086 = v1054
		goto L201
	} else {
		goto L206
	}
L206:
	;
	v1077 = *(*int32)(unsafe.Add(mBase, uint32(v197)+472))
	v1078 = *(*int32)(unsafe.Add(mBase, uint32(v197)+480))
	F_statatt_set_slot(m, v197+int32(176), v197+int32(144), v197+int32(112), int32(2), v1077, v1078, int64(0), int32(1), v1065, int32(0))
	mBase = m.M
	v1083 = m.ExcPending
	if v1083 != 0 {
		goto L1
	} else {
		goto L207
	}
L207:
	;
	goto L204
L208:
	;
	v1090 = *(*int64)(unsafe.Add(mBase, uint32(v187)+200))
	*(*int64)(unsafe.Add(mBase, uint32(v197)+104)) = v1090
	v1096 = F_construct_array_builtin(m, v197+int32(104), int32(1), int32(700))
	mBase = m.M
	v1097 = m.ExcPending
	if v1097 != 0 {
		goto L1
	} else {
		goto L211
	}
L209:
	;
	goto L210
L210:
	;
	if v751&int32(1) != 0 {
		goto L214
	} else {
		goto L215
	}
L211:
	;
	v1105 = *(*int32)(unsafe.Add(mBase, uint32(v197)+472))
	v1106 = *(*int32)(unsafe.Add(mBase, uint32(v197)+480))
	F_statatt_set_slot(m, v197+int32(176), v197+int32(144), v197+int32(112), int32(3), v1105, v1106, base.I64_extend_i32_u(v1096), int32(0), int64(0), int32(1))
	mBase = m.M
	v1112 = m.ExcPending
	if v1112 != 0 {
		goto L1
	} else {
		goto L212
	}
L212:
	;
	goto L210
L213:
	;
	if v750&int32(1) != 0 {
		goto L220
	} else {
		goto L221
	}
L214:
	;
	v1116 = *(*int64)(unsafe.Add(mBase, uint32(v187)+232))
	v1117 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v197)+104)) = uint8(v1117)
	v1123 = *(*int64)(unsafe.Add(mBase, uint32(v187)+216))
	v1124 = *(*int32)(unsafe.Add(mBase, uint32(v197)+468))
	v1125 = *(*int32)(unsafe.Add(mBase, uint32(v197)+488))
	v1128 = F_statatt_build_stavalues(m, int32(_a_F_pg_restore_attribute_stats_12), v197+int32(436), v1123, v1124, v1125, v197+int32(104))
	mBase = m.M
	v1129 = m.ExcPending
	if v1129 != 0 {
		goto L1
	} else {
		goto L217
	}
L215:
	;
	goto L216
L216:
	;
	v1149 = v1086
	goto L213
L217:
	;
	v1130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v197)+104)))
	if v1130 != int32(1) {
		v1149 = v1117
		goto L213
	} else {
		goto L218
	}
L218:
	;
	v1140 = *(*int32)(unsafe.Add(mBase, uint32(v197)+464))
	v1141 = *(*int32)(unsafe.Add(mBase, uint32(v197)+480))
	v1142 = int32(0)
	F_statatt_set_slot(m, v197+int32(176), v197+int32(144), v197+int32(112), int32(4), v1140, v1141, v1116, v1142, v1128, v1142)
	mBase = m.M
	v1145 = m.ExcPending
	if v1145 != 0 {
		goto L1
	} else {
		goto L219
	}
L219:
	;
	goto L216
L220:
	;
	v1161 = *(*int32)(unsafe.Add(mBase, uint32(v197)+464))
	v1162 = *(*int32)(unsafe.Add(mBase, uint32(v197)+480))
	v1163 = *(*int64)(unsafe.Add(mBase, uint32(v187)+248))
	F_statatt_set_slot(m, v197+int32(176), v197+int32(144), v197+int32(112), int32(5), v1161, v1162, v1163, int32(0), int64(0), int32(1))
	mBase = m.M
	v1168 = m.ExcPending
	if v1168 != 0 {
		goto L1
	} else {
		goto L223
	}
L221:
	;
	goto L222
L222:
	;
	if v832&int32(1) != 0 {
		goto L225
	} else {
		goto L226
	}
L223:
	;
	goto L222
L224:
	;
	if v831 != 0 {
		goto L237
	} else {
		goto L238
	}
L225:
	;
	v1171 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v197)+104)) = uint8(v1171)
	v1174 = *(*int32)(unsafe.Add(mBase, uint32(v197)+492))
	v1175 = F_type_is_multirange(m, v1174)
	mBase = m.M
	v1176 = m.ExcPending
	if v1176 != 0 {
		goto L1
	} else {
		goto L228
	}
L226:
	;
	goto L227
L227:
	;
	v1210 = v1149
	goto L224
L228:
	;
	if v1175 != 0 {
		goto L229
	} else {
		goto L230
	}
L229:
	;
	v1177 = *(*int32)(unsafe.Add(mBase, uint32(v197)+492))
	v1178 = F_get_multirange_range(m, v1177)
	mBase = m.M
	v1179 = m.ExcPending
	if v1179 != 0 {
		goto L1
	} else {
		goto L232
	}
L230:
	;
	v1180 = v1174
	goto L231
L231:
	;
	v1184 = *(*int64)(unsafe.Add(mBase, uint32(v187)+296))
	v1185 = *(*int32)(unsafe.Add(mBase, uint32(v197)+488))
	v1188 = F_statatt_build_stavalues(m, int32(_a_F_pg_restore_attribute_stats_13), v197+int32(436), v1184, v1180, v1185, v197+int32(104))
	mBase = m.M
	v1189 = m.ExcPending
	if v1189 != 0 {
		goto L1
	} else {
		goto L233
	}
L232:
	;
	v1180 = v1178
	goto L231
L233:
	;
	v1190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v197)+104)))
	if v1190 != int32(1) {
		v1210 = v1171
		goto L224
	} else {
		goto L234
	}
L234:
	;
	v1200 = int32(0)
	F_statatt_set_slot(m, v197+int32(176), v197+int32(144), v197+int32(112), int32(7), v1200, v1200, int64(0), int32(1), v1188, v1200)
	mBase = m.M
	v1206 = m.ExcPending
	if v1206 != 0 {
		goto L1
	} else {
		goto L235
	}
L235:
	;
	goto L227
L236:
	;
	v1257 = *(*int32)(unsafe.Add(mBase, uint32(v840)+52))
	if v846 != 0 {
		goto L245
	} else {
		goto L246
	}
L237:
	;
	v1213 = *(*int64)(unsafe.Add(mBase, uint32(v187)+280))
	*(*int64)(unsafe.Add(mBase, uint32(v197)+104)) = v1213
	v1219 = F_construct_array_builtin(m, v197+int32(104), int32(1), int32(700))
	mBase = m.M
	v1220 = m.ExcPending
	if v1220 != 0 {
		goto L1
	} else {
		goto L240
	}
L238:
	;
	goto L239
L239:
	;
	v1255 = v1210
	goto L236
L240:
	;
	v1221 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v197)+103)) = uint8(v1221)
	v1227 = *(*int64)(unsafe.Add(mBase, uint32(v187)+264))
	v1232 = F_statatt_build_stavalues(m, int32(_a_F_pg_restore_attribute_stats_14), v197+int32(436), v1227, int32(701), v1221, v197+int32(103))
	mBase = m.M
	v1233 = m.ExcPending
	if v1233 != 0 {
		goto L1
	} else {
		goto L241
	}
L241:
	;
	v1234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v197)+103)))
	if v1234 != int32(1) {
		v1255 = v1221
		goto L236
	} else {
		goto L242
	}
L242:
	;
	v1245 = int32(0)
	F_statatt_set_slot(m, v197+int32(176), v197+int32(144), v197+int32(112), int32(6), int32(672), v1245, base.I64_extend_i32_u(v1219), v1245, v1232, v1245)
	mBase = m.M
	v1250 = m.ExcPending
	if v1250 != 0 {
		goto L1
	} else {
		goto L243
	}
L243:
	;
	goto L239
L244:
	;
	F_pfree(m, v1278)
	mBase = m.M
	v1280 = m.ExcPending
	if v1280 != 0 {
		goto L1
	} else {
		goto L252
	}
L245:
	;
	v1264 = F_heap_modify_tuple(m, v846, v1257, v197+int32(176), v197+int32(144), v197+int32(112))
	mBase = m.M
	v1265 = m.ExcPending
	if v1265 != 0 {
		goto L1
	} else {
		goto L248
	}
L246:
	;
	goto L247
L247:
	;
	v1274 = F_heap_form_tuple(m, v1257, v197+int32(176), v197+int32(144))
	mBase = m.M
	v1275 = m.ExcPending
	if v1275 != 0 {
		goto L1
	} else {
		goto L250
	}
L248:
	;
	F_CatalogTupleUpdate(m, v840, v1264+int32(4), v1264)
	mBase = m.M
	v1269 = m.ExcPending
	if v1269 != 0 {
		goto L1
	} else {
		goto L249
	}
L249:
	;
	v1278 = v1264
	goto L244
L250:
	;
	F_CatalogTupleInsert(m, v840, v1274)
	mBase = m.M
	v1277 = m.ExcPending
	if v1277 != 0 {
		goto L1
	} else {
		goto L251
	}
L251:
	;
	v1278 = v1274
	goto L244
L252:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v1282 = m.ExcPending
	if v1282 != 0 {
		goto L1
	} else {
		goto L253
	}
L253:
	;
	if v846 != 0 {
		goto L254
	} else {
		goto L255
	}
L254:
	;
	F_ReleaseCatCache(m, v846)
	mBase = m.M
	v1284 = m.ExcPending
	if v1284 != 0 {
		goto L1
	} else {
		goto L257
	}
L255:
	;
	goto L256
L256:
	;
	F_relation_close(m, v840, int32(3))
	mBase = m.M
	v1287 = m.ExcPending
	if v1287 != 0 {
		goto L1
	} else {
		goto L258
	}
L257:
	;
	goto L256
L258:
	;
	m.G0 = v197 + int32(496)
	m.G0 = v44 + int32(400)
	return base.I64_extend_i32_u(v1255 & v57)
L259:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v1303 = m.ExcPending
	if v1303 != 0 {
		goto L1
	} else {
		goto L260
	}
L260:
	;
	F_errmsg(m, int32(_a_F_pg_restore_attribute_stats_32), int32(0))
	mBase = m.M
	v1307 = m.ExcPending
	if v1307 != 0 {
		goto L1
	} else {
		goto L261
	}
L261:
	;
	F_errhint(m, int32(_a_F_pg_restore_attribute_stats_33), int32(0))
	mBase = m.M
	v1311 = m.ExcPending
	if v1311 != 0 {
		goto L1
	} else {
		goto L262
	}
L262:
	;
	F_errfinish(m, int32(_a_F_pg_restore_attribute_stats_4), int32(154), int32(_a_F_pg_restore_attribute_stats_5))
	mBase = m.M
	v1316 = m.ExcPending
	if v1316 != 0 {
		goto L1
	} else {
		goto L263
	}
L263:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L264:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v1323 = m.ExcPending
	if v1323 != 0 {
		goto L1
	} else {
		goto L265
	}
L265:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+68)) = int32(_a_F_pg_restore_attribute_stats_1)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+64)) = int32(_a_F_pg_restore_attribute_stats_2)
	F_errmsg(m, int32(_a_F_pg_restore_attribute_stats_3), v44-int32(-64))
	mBase = m.M
	v1332 = m.ExcPending
	if v1332 != 0 {
		goto L1
	} else {
		goto L266
	}
L266:
	;
	F_errfinish(m, int32(_a_F_pg_restore_attribute_stats_4), int32(167), int32(_a_F_pg_restore_attribute_stats_5))
	mBase = m.M
	v1337 = m.ExcPending
	if v1337 != 0 {
		goto L1
	} else {
		goto L267
	}
L267:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L268:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v1344 = m.ExcPending
	if v1344 != 0 {
		goto L1
	} else {
		goto L269
	}
L269:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+32)) = v183
	F_errmsg(m, int32(_a_F_pg_restore_attribute_stats_6), v44+int32(32))
	mBase = m.M
	v1350 = m.ExcPending
	if v1350 != 0 {
		goto L1
	} else {
		goto L270
	}
L270:
	;
	F_errfinish(m, int32(_a_F_pg_restore_attribute_stats_4), int32(202), int32(_a_F_pg_restore_attribute_stats_5))
	mBase = m.M
	v1355 = m.ExcPending
	if v1355 != 0 {
		goto L1
	} else {
		goto L271
	}
L271:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
