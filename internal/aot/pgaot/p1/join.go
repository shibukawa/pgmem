package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_build_join_pathkeys(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	if base.Ui32(l2) <= base.Ui32(int32(7)) {
		if int32(1)<<(uint(l2)%32)&int32(140) != 0 {
			v16 = int32(0)
			return v16
		} else {
			v12 = F_truncate_useless_pathkeys(m, l0, l1, l3)
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				v16 = v12
				return v16
			}
		}
	} else {
		v12 = F_truncate_useless_pathkeys(m, l0, l1, l3)
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			v16 = v12
			return v16
		}
	}
}
func F_flatten_join_alias_vars(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = l1
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+39)))
	*(*uint8)(unsafe.Add(mBase, uint32(v7)+13)) = uint8(v13)
	*(*uint8)(unsafe.Add(mBase, uint32(v7)+12)) = uint8(v13)
	v16 = F_flatten_join_alias_vars_mutator(m, l2, v7)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return int32(0)
	} else {
		m.G0 = v7 + int32(16)
		return v16
	}
}
func F_make_join_rel(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
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
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int64
	_ = v53
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v98 float64
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int64
	_ = v106
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v137 int32
	_ = v137
	var v138 int64
	_ = v138
	var v148 int32
	_ = v148
	var v158 int64
	_ = v158
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v229 int64
	_ = v229
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v241 int32
	_ = v241
	var v253 int64
	_ = v253
	var v255 int32
	_ = v255
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v282 int32
	_ = v282
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	var v305 int32
	_ = v305
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v323 int32
	_ = v323
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v338 int32
	_ = v338
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v352 int32
	_ = v352
	var v354 int32
	_ = v354
	var v356 int32
	_ = v356
	var v361 int32
	_ = v361
	var v370 int32
	_ = v370
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v387 int32
	_ = v387
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v396 int32
	_ = v396
	var v403 int32
	_ = v403
	var v405 int32
	_ = v405
	var v407 int32
	_ = v407
	var v410 int32
	_ = v410
	var v412 int32
	_ = v412
	var v414 int32
	_ = v414
	var v419 int32
	_ = v419
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v443 int32
	_ = v443
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v452 int32
	_ = v452
	var v459 int32
	_ = v459
	var v461 int32
	_ = v461
	var v463 int32
	_ = v463
	var v466 int32
	_ = v466
	var v468 int32
	_ = v468
	var v470 int32
	_ = v470
	var v475 int32
	_ = v475
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v494 int32
	_ = v494
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v498 float64
	_ = v498
	var v499 float64
	_ = v499
	var v502 int32
	_ = v502
	var v503 float64
	_ = v503
	var v504 float64
	_ = v504
	var v507 int64
	_ = v507
	var v511 int64
	_ = v511
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v519 int64
	_ = v519
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v540 int64
	_ = v540
	var v542 int64
	_ = v542
	var v545 int64
	_ = v545
	var v547 int32
	_ = v547
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v576 int32
	_ = v576
	var v579 int32
	_ = v579
	var v590 int32
	_ = v590
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v609 int32
	_ = v609
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v618 int32
	_ = v618
	var v625 int32
	_ = v625
	var v627 int32
	_ = v627
	var v629 int32
	_ = v629
	var v632 int32
	_ = v632
	var v634 int32
	_ = v634
	var v636 int32
	_ = v636
	var v641 int32
	_ = v641
	var v650 int32
	_ = v650
	var v653 int32
	_ = v653
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v664 int32
	_ = v664
	var v678 int32
	_ = v678
	var v681 int32
	_ = v681
	var v689 int32
	_ = v689
	var v692 int32
	_ = v692
	var v703 int32
	_ = v703
	var v707 int32
	_ = v707
	var v708 int32
	_ = v708
	var v709 int32
	_ = v709
	var v710 int32
	_ = v710
	var v719 int32
	_ = v719
	var v720 int32
	_ = v720
	var v722 int32
	_ = v722
	var v725 int32
	_ = v725
	var v726 int32
	_ = v726
	var v731 int32
	_ = v731
	var v738 int32
	_ = v738
	var v740 int32
	_ = v740
	var v742 int32
	_ = v742
	var v745 int32
	_ = v745
	var v747 int32
	_ = v747
	var v749 int32
	_ = v749
	var v754 int32
	_ = v754
	var v763 int32
	_ = v763
	var v766 int32
	_ = v766
	var v767 int32
	_ = v767
	var v768 int32
	_ = v768
	var v770 int32
	_ = v770
	var v771 int32
	_ = v771
	var v777 int32
	_ = v777
	var v792 int32
	_ = v792
	var v793 int32
	_ = v793
	var v803 int32
	_ = v803
	var v804 int32
	_ = v804
	var v807 int32
	_ = v807
	var v811 int32
	_ = v811
	var v814 int32
	_ = v814
	var v816 int32
	_ = v816
	var v819 int32
	_ = v819
	var v826 int32
	_ = v826
	var v828 int32
	_ = v828
	var v836 int32
	_ = v836
	var v837 int32
	_ = v837
	var v850 int32
	_ = v850
	var v857 int32
	_ = v857
	var v858 int32
	_ = v858
	var v871 int32
	_ = v871
	var v873 int32
	_ = v873
	var v877 int32
	_ = v877
	var v880 int32
	_ = v880
	var v881 int32
	_ = v881
	var v882 int32
	_ = v882
	var v883 int32
	_ = v883
	var v891 int32
	_ = v891
	var v893 int32
	_ = v893
	var v894 int32
	_ = v894
	var v897 int32
	_ = v897
	var v901 int32
	_ = v901
	var v904 int32
	_ = v904
	var v906 int32
	_ = v906
	var v909 int32
	_ = v909
	var v916 int32
	_ = v916
	var v918 int32
	_ = v918
	var v926 int32
	_ = v926
	var v927 int32
	_ = v927
	var v940 int32
	_ = v940
	var v947 int32
	_ = v947
	var v970 int32
	_ = v970
	var v971 int32
	_ = v971
	var v974 int32
	_ = v974
	var v978 int32
	_ = v978
	var v981 int32
	_ = v981
	var v983 int32
	_ = v983
	var v986 int32
	_ = v986
	var v993 int32
	_ = v993
	var v995 int32
	_ = v995
	var v1003 int32
	_ = v1003
	var v1004 int32
	_ = v1004
	var v1017 int32
	_ = v1017
	var v1025 int32
	_ = v1025
	var v1038 int32
	_ = v1038
	var v1039 int32
	_ = v1039
	var v1043 int32
	_ = v1043
	var v1044 int32
	_ = v1044
	var v1047 int32
	_ = v1047
	var v1050 int32
	_ = v1050
	var v1051 int32
	_ = v1051
	var v1052 int32
	_ = v1052
	var v1061 int32
	_ = v1061
	var v1062 int32
	_ = v1062
	var v1064 int32
	_ = v1064
	var v1067 int32
	_ = v1067
	var v1068 int32
	_ = v1068
	var v1073 int32
	_ = v1073
	var v1080 int32
	_ = v1080
	var v1082 int32
	_ = v1082
	var v1084 int32
	_ = v1084
	var v1087 int32
	_ = v1087
	var v1089 int32
	_ = v1089
	var v1091 int32
	_ = v1091
	var v1096 int32
	_ = v1096
	var v1105 int32
	_ = v1105
	var v1113 int32
	_ = v1113
	var v1115 int32
	_ = v1115
	var v1116 int32
	_ = v1116
	var v1119 int32
	_ = v1119
	var v1123 int32
	_ = v1123
	var v1126 int32
	_ = v1126
	var v1128 int32
	_ = v1128
	var v1131 int32
	_ = v1131
	var v1138 int32
	_ = v1138
	var v1140 int32
	_ = v1140
	var v1148 int32
	_ = v1148
	var v1149 int32
	_ = v1149
	var v1162 int32
	_ = v1162
	var v1202 int32
	_ = v1202
	var v1205 int32
	_ = v1205
	var v1207 int32
	_ = v1207
	var v1208 int32
	_ = v1208
	var v1211 int32
	_ = v1211
	var v1214 int32
	_ = v1214
	var v1215 int32
	_ = v1215
	var v1218 int32
	_ = v1218
	var v1219 int32
	_ = v1219
	var v1220 int32
	_ = v1220
	var v1221 int32
	_ = v1221
	var v1224 int32
	_ = v1224
	var v1226 int32
	_ = v1226
	var v1227 int32
	_ = v1227
	var v1228 int32
	_ = v1228
	var v1230 int32
	_ = v1230
	var v1234 int32
	_ = v1234
	var v1235 int32
	_ = v1235
	var v1237 int32
	_ = v1237
	var v1240 int32
	_ = v1240
	var v1244 int32
	_ = v1244
	var v1245 int32
	_ = v1245
	var v1246 int32
	_ = v1246
	var v1247 int32
	_ = v1247
	var v1248 int32
	_ = v1248
	var v1256 int32
	_ = v1256
	var v1274 int32
	_ = v1274
	var v1277 int32
	_ = v1277
	var v1280 int32
	_ = v1280
	var v1296 int32
	_ = v1296
	var v1297 int32
	_ = v1297
	var v1304 int32
	_ = v1304
	var v1327 int32
	_ = v1327
	var v1329 int32
	_ = v1329
	var v1349 int32
	_ = v1349
	var v1355 int32
	_ = v1355
	v4 = int32(0)
	v19 = m.G0
	v21 = v19 - int32(80)
	m.G0 = v21
	*(*int32)(unsafe.Add(mBase, uint32(v21)+68)) = v4
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v27 = F_bms_union(m, v25, v26)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v21 + int32(80)
	return v1355
L2:
	;
	return int32(0)
L3:
	;
	v35 = F_join_is_legal(m, l0, l1, l2, v27, v21+int32(76), v21+int32(75))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	if v35 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	F_bms_free(m, v27)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L2
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v21)+76))
	v44 = F_add_outer_joins_to_relids(m, l0, v27, v41, v21+int32(68))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L2
	} else {
		goto L9
	}
L8:
	;
	v1355 = v4
	goto L1
L9:
	;
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+75)))
	if v46 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v47 = l2
	goto L12
L11:
	;
	v47 = l1
	goto L12
L12:
	;
	if v46 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v48 = l1
	goto L15
L14:
	;
	v48 = l2
	goto L15
L15:
	;
	if v41 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v47)+8))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v48)+8))
	v53 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v21)+40)) = v53
	*(*int64)(unsafe.Add(mBase, uint32(v21)+48)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v21)+55)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v21)+60)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v21)+28)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v21)+24)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v21)+20)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v21)+16)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v21)+12)) = int32(320)
	*(*int64)(unsafe.Add(mBase, uint32(v21)+32)) = v53
	v73 = v21 + int32(12)
	goto L18
L17:
	;
	v73 = v41
	goto L18
L18:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v21)+68))
	v76 = v21 + int32(8)
	v77 = m.G0
	v79 = v77 - int32(16)
	m.G0 = v79
	v81 = F_find_join_rel(m, l0, v44)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L2
	} else {
		goto L20
	}
L19:
	;
	m.G0 = v79 + int32(16)
	v1274 = *(*int32)(unsafe.Add(mBase, uint32(v1256)+32))
	if v1274 == int32(0) {
		goto L284
	} else {
		goto L285
	}
L20:
	;
	if v81 != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	if v76 == int32(0) {
		v1256 = v81
		goto L19
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v89 = F_palloc0(m, int32(272))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L2
	} else {
		goto L26
	}
L24:
	;
	v85 = F_build_joinrel_restrictlist(m, l0, v81, v47, v48, v73)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L2
	} else {
		goto L25
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v76))) = v85
	v1256 = v81
	goto L19
L26:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v89))) = int64(4294967564)
	v93 = F_bms_copy(m, v44)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L2
	} else {
		goto L27
	}
L27:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v89)+16)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v89)+8)) = v93
	v98 = *(*float64)(unsafe.Add(mBase, uint32(l0)+296))
	v99 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v89)+25)) = uint16(v99)
	v102 = base.F64_gt(v98, float64(0))
	*(*uint8)(unsafe.Add(mBase, uint32(v89)+24)) = uint8(v102)
	v104 = F_create_empty_pathtarget(m)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L2
	} else {
		goto L28
	}
L28:
	;
	v106 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v89)+32)) = v106
	*(*int32)(unsafe.Add(mBase, uint32(v89)+28)) = v104
	*(*int64)(unsafe.Add(mBase, uint32(v89)+40)) = v106
	*(*int64)(unsafe.Add(mBase, uint32(v89)+48)) = v106
	*(*int32)(unsafe.Add(mBase, uint32(v89)+56)) = int32(0)
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v47)+60))
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v48)+60))
	v117 = F_bms_union(m, v115, v116)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L2
	} else {
		goto L29
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v89)+60)) = v117
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v89)+8))
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v47)+64))
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v48)+64))
	v123 = F_bms_union(m, v121, v122)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L2
	} else {
		goto L30
	}
L30:
	;
	v125 = F_bms_del_members(m, v123, v120)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L2
	} else {
		goto L31
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v89)+76)) = int32(2)
	v129 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v89)+68)) = v129
	*(*int32)(unsafe.Add(mBase, uint32(v89)+64)) = v125
	v137 = F__emscripten_memset_bulkmem(m, v89+int32(80), base.I32_extend8_s(v129), int32(68))
	mBase = m.M
	goto L32
L32:
	;
	v138 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v89)+152)) = v138
	*(*int32)(unsafe.Add(mBase, uint32(v89)+148)) = int32(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v89)+157)) = v138
	*(*int64)(unsafe.Add(mBase, uint32(v89)+168)) = v138
	*(*int64)(unsafe.Add(mBase, uint32(v89)+176)) = v138
	v148 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v89)+184)) = v148
	*(*int64)(unsafe.Add(mBase, uint32(v89)+192)) = v138
	*(*int64)(unsafe.Add(mBase, uint32(v89)+200)) = v138
	*(*int64)(unsafe.Add(mBase, uint32(v89)+220)) = v138
	*(*uint16)(unsafe.Add(mBase, uint32(v89)+216)) = uint16(v148)
	v158 = int64(4294967295)
	*(*int64)(unsafe.Add(mBase, uint32(v89)+208)) = v158
	*(*int64)(unsafe.Add(mBase, uint32(v89)+228)) = v138
	*(*uint8)(unsafe.Add(mBase, uint32(v89)+244)) = uint8(v148)
	*(*int64)(unsafe.Add(mBase, uint32(v89)+236)) = v158
	*(*int64)(unsafe.Add(mBase, uint32(v89)+248)) = v138
	*(*int64)(unsafe.Add(mBase, uint32(v89)+256)) = v138
	*(*int64)(unsafe.Add(mBase, uint32(v89)+264)) = v138
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v47)+156))
	if v172 == v148 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v73)+20))
	F_build_joinrel_tlist(m, l0, v89, v47, v73, v74, base.B2i32(v213 == int32(2)))
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L2
	} else {
		goto L52
	}
L34:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v48)+156))
	if v175 != v172 {
		goto L33
	} else {
		goto L35
	}
L35:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v48)+160))
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v47)+160))
	if v177 == v178 {
		goto L37
	} else {
		goto L38
	}
L36:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v89)+164)) = uint8(v207)
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v47)+168))
	*(*int32)(unsafe.Add(mBase, uint32(v89)+168)) = v209
	goto L33
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v89)+156)) = v172
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v47)+160))
	*(*int32)(unsafe.Add(mBase, uint32(v89)+160)) = v181
	v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+164)))
	if v183 != 0 {
		goto L40
	} else {
		goto L41
	}
L38:
	;
	goto L39
L39:
	;
	if v177 != 0 {
		goto L44
	} else {
		goto L45
	}
L40:
	;
	v186 = int32(1)
	goto L42
L41:
	;
	v185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+164)))
	v186 = v185
	goto L42
L42:
	;
	v207 = v186 & int32(1)
	goto L36
L43:
	;
	v207 = int32(1)
	goto L36
L44:
	;
	v197 = v178
	goto L46
L45:
	;
	v190 = *(*int32)(unsafe.Add(mBase, _consts[31]))
	if v178 == v190 {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	if v197 != 0 {
		goto L33
	} else {
		goto L50
	}
L47:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v47)+156))
	*(*int32)(unsafe.Add(mBase, uint32(v89)+156)) = v192
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v47)+160))
	*(*int32)(unsafe.Add(mBase, uint32(v89)+160)) = v194
	goto L43
L48:
	;
	goto L49
L49:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v47)+160))
	v197 = v196
	goto L46
L50:
	;
	v199 = *(*int32)(unsafe.Add(mBase, _consts[31]))
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v48)+160))
	if v199 != v200 {
		goto L33
	} else {
		goto L51
	}
L51:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v47)+156))
	*(*int32)(unsafe.Add(mBase, uint32(v89)+156)) = v202
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v48)+160))
	*(*int32)(unsafe.Add(mBase, uint32(v89)+160)) = v204
	goto L43
L52:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v73)+20))
	F_build_joinrel_tlist(m, l0, v89, v48, v73, v74, base.B2i32(v218 != int32(0)))
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L2
	} else {
		goto L53
	}
L53:
	;
	v223 = int32(0)
	v224 = m.G0
	v226 = v224 - int32(16)
	m.G0 = v226
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v89)+28))
	v229 = int64(*(*int32)(unsafe.Add(mBase, uint32(v228)+32)))
	v230 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	if v230 == v223 {
		v540 = v229
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v542 = int64(1073741823)
	if v542 <= v540 {
		goto L128
	} else {
		goto L129
	}
L55:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v230)+4))
	if v233 <= int32(0) {
		v540 = v229
		goto L54
	} else {
		goto L56
	}
L56:
	;
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v89)+8))
	v241 = v223
	v253 = v229
	goto L57
L57:
	;
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v230)+12))
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v255+v241<<(uint(int32(2))%32))))
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v259)+12))
	v261 = int32(0)
	if v260 == v261 {
		goto L60
	} else {
		goto L61
	}
L58:
	;
	v540 = v519
	goto L54
L59:
	;
	if v314 != 0 {
		goto L73
	} else {
		goto L74
	}
L60:
	;
	v314 = int32(1)
	goto L59
L61:
	;
	goto L62
L62:
	;
	if v236 == int32(0) {
		v305 = v261
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v314 = v305
	goto L59
L64:
	;
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v260)+4))
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v236)+4))
	if v271 < v270 {
		v305 = v261
		goto L63
	} else {
		goto L65
	}
L65:
	;
	v273 = int32(1)
	if v270 <= v273 {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v276 = v273
	goto L68
L67:
	;
	v276 = v270
	goto L68
L68:
	;
	v277 = int32(8)
	v282 = int32(0)
	goto L69
L69:
	;
	v289 = v282 << (uint(int32(2)) % 32)
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v260+v277+v289)))
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v289+(v236+v277))))
	v296 = v291 & (v293 ^ int32(-1))
	v298 = base.B2i32(v296 == int32(0))
	if v296 != 0 {
		v305 = v298
		goto L63
	} else {
		goto L71
	}
L70:
	;
	v305 = v298
	goto L63
L71:
	;
	v300 = v282 + int32(1)
	if v300 != v276 {
		v282 = v300
		goto L69
	} else {
		goto L72
	}
L72:
	;
	goto L70
L73:
	;
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v259)+20))
	if v315 == int32(0) {
		goto L78
	} else {
		goto L79
	}
L74:
	;
	v519 = v253
	goto L75
L75:
	;
	v521 = v241 + int32(1)
	v522 = *(*int32)(unsafe.Add(mBase, uint32(v230)+4))
	if v521 < v522 {
		v241 = v521
		v253 = v519
		goto L57
	} else {
		goto L126
	}
L76:
	;
	v512 = *(*int32)(unsafe.Add(mBase, uint32(v89)+60))
	v513 = *(*int32)(unsafe.Add(mBase, uint32(v259)+16))
	v514 = F_bms_add_members(m, v512, v513)
	mBase = m.M
	v515 = m.ExcPending
	if v515 != 0 {
		goto L2
	} else {
		goto L125
	}
L77:
	;
	if v370 == int32(0) {
		v511 = v253
		goto L76
	} else {
		goto L91
	}
L78:
	;
	v370 = int32(0)
	goto L77
L79:
	;
	goto L80
L80:
	;
	v323 = int32(1)
	if v236 == int32(0) {
		v361 = v323
		goto L81
	} else {
		goto L82
	}
L81:
	;
	v370 = v361
	goto L77
L82:
	;
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v315)+4))
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v236)+4))
	if v327 < v326 {
		v361 = v323
		goto L81
	} else {
		goto L83
	}
L83:
	;
	v329 = int32(1)
	if v326 <= v329 {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	v332 = v329
	goto L86
L85:
	;
	v332 = v326
	goto L86
L86:
	;
	v333 = int32(8)
	v338 = int32(0)
	goto L87
L87:
	;
	v345 = v338 << (uint(int32(2)) % 32)
	v347 = *(*int32)(unsafe.Add(mBase, uint32(v315+v333+v345)))
	v349 = *(*int32)(unsafe.Add(mBase, uint32(v345+(v236+v333))))
	v352 = v347 & (v349 ^ int32(-1))
	v354 = base.B2i32(v352 != int32(0))
	if v352 != 0 {
		v361 = v354
		goto L81
	} else {
		goto L89
	}
L88:
	;
	v361 = v354
	goto L81
L89:
	;
	v356 = v338 + int32(1)
	if v356 != v332 {
		v338 = v356
		goto L87
	} else {
		goto L90
	}
L90:
	;
	goto L88
L91:
	;
	v373 = *(*int32)(unsafe.Add(mBase, uint32(v259)+12))
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v47)+8))
	v375 = int32(0)
	if v373 == v375 {
		goto L93
	} else {
		goto L94
	}
L92:
	;
	if v428 != 0 {
		v511 = v253
		goto L76
	} else {
		goto L106
	}
L93:
	;
	v428 = int32(1)
	goto L92
L94:
	;
	goto L95
L95:
	;
	if v374 == int32(0) {
		v419 = v375
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v428 = v419
	goto L92
L97:
	;
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v373)+4))
	v385 = *(*int32)(unsafe.Add(mBase, uint32(v374)+4))
	if v385 < v384 {
		v419 = v375
		goto L96
	} else {
		goto L98
	}
L98:
	;
	v387 = int32(1)
	if v384 <= v387 {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	v390 = v387
	goto L101
L100:
	;
	v390 = v384
	goto L101
L101:
	;
	v391 = int32(8)
	v396 = int32(0)
	goto L102
L102:
	;
	v403 = v396 << (uint(int32(2)) % 32)
	v405 = *(*int32)(unsafe.Add(mBase, uint32(v373+v391+v403)))
	v407 = *(*int32)(unsafe.Add(mBase, uint32(v403+(v374+v391))))
	v410 = v405 & (v407 ^ int32(-1))
	v412 = base.B2i32(v410 == int32(0))
	if v410 != 0 {
		v419 = v412
		goto L96
	} else {
		goto L104
	}
L103:
	;
	v419 = v412
	goto L96
L104:
	;
	v414 = v396 + int32(1)
	if v414 != v390 {
		v396 = v414
		goto L102
	} else {
		goto L105
	}
L105:
	;
	goto L103
L106:
	;
	v429 = *(*int32)(unsafe.Add(mBase, uint32(v259)+12))
	v430 = *(*int32)(unsafe.Add(mBase, uint32(v48)+8))
	v431 = int32(0)
	if v429 == v431 {
		goto L108
	} else {
		goto L109
	}
L107:
	;
	if v484 != 0 {
		v511 = v253
		goto L76
	} else {
		goto L121
	}
L108:
	;
	v484 = int32(1)
	goto L107
L109:
	;
	goto L110
L110:
	;
	if v430 == int32(0) {
		v475 = v431
		goto L111
	} else {
		goto L112
	}
L111:
	;
	v484 = v475
	goto L107
L112:
	;
	v440 = *(*int32)(unsafe.Add(mBase, uint32(v429)+4))
	v441 = *(*int32)(unsafe.Add(mBase, uint32(v430)+4))
	if v441 < v440 {
		v475 = v431
		goto L111
	} else {
		goto L113
	}
L113:
	;
	v443 = int32(1)
	if v440 <= v443 {
		goto L114
	} else {
		goto L115
	}
L114:
	;
	v446 = v443
	goto L116
L115:
	;
	v446 = v440
	goto L116
L116:
	;
	v447 = int32(8)
	v452 = int32(0)
	goto L117
L117:
	;
	v459 = v452 << (uint(int32(2)) % 32)
	v461 = *(*int32)(unsafe.Add(mBase, uint32(v429+v447+v459)))
	v463 = *(*int32)(unsafe.Add(mBase, uint32(v459+(v430+v447))))
	v466 = v461 & (v463 ^ int32(-1))
	v468 = base.B2i32(v466 == int32(0))
	if v466 != 0 {
		v475 = v468
		goto L111
	} else {
		goto L119
	}
L118:
	;
	v475 = v468
	goto L111
L119:
	;
	v470 = v452 + int32(1)
	if v470 != v446 {
		v452 = v470
		goto L117
	} else {
		goto L120
	}
L120:
	;
	goto L118
L121:
	;
	v485 = *(*int32)(unsafe.Add(mBase, uint32(v259)+8))
	v486 = F_copyObjectImpl(m, v485)
	mBase = m.M
	v487 = m.ExcPending
	if v487 != 0 {
		goto L2
	} else {
		goto L122
	}
L122:
	;
	v488 = *(*int32)(unsafe.Add(mBase, uint32(v89)+28))
	v489 = *(*int32)(unsafe.Add(mBase, uint32(v488)+4))
	v490 = F_lappend(m, v489, v486)
	mBase = m.M
	v491 = m.ExcPending
	if v491 != 0 {
		goto L2
	} else {
		goto L123
	}
L123:
	;
	v492 = *(*int32)(unsafe.Add(mBase, uint32(v89)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v492)+4)) = v490
	v494 = *(*int32)(unsafe.Add(mBase, uint32(v486)+4))
	F_cost_qual_eval_node(m, v226, v494, l0)
	mBase = m.M
	v496 = m.ExcPending
	if v496 != 0 {
		goto L2
	} else {
		goto L124
	}
L124:
	;
	v497 = *(*int32)(unsafe.Add(mBase, uint32(v89)+28))
	v498 = *(*float64)(unsafe.Add(mBase, uint32(v226)))
	v499 = *(*float64)(unsafe.Add(mBase, uint32(v497)+16))
	*(*float64)(unsafe.Add(mBase, uint32(v497)+16)) = base.F64_add(v498, v499)
	v502 = *(*int32)(unsafe.Add(mBase, uint32(v89)+28))
	v503 = *(*float64)(unsafe.Add(mBase, uint32(v226)+8))
	v504 = *(*float64)(unsafe.Add(mBase, uint32(v502)+24))
	*(*float64)(unsafe.Add(mBase, uint32(v502)+24)) = base.F64_add(v503, v504)
	v507 = int64(*(*int32)(unsafe.Add(mBase, uint32(v259)+24)))
	v511 = v253 + v507
	goto L76
L125:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v89)+60)) = v514
	v519 = v511
	goto L75
L126:
	;
	goto L58
L127:
	;
	v547 = *(*int32)(unsafe.Add(mBase, uint32(v89)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v547)+32)) = base.I32_wrap_i64(v545)
	m.G0 = v226 + int32(16)
	v552 = *(*int32)(unsafe.Add(mBase, uint32(v89)+60))
	v553 = *(*int32)(unsafe.Add(mBase, uint32(v89)+8))
	v554 = F_bms_del_members(m, v552, v553)
	mBase = m.M
	v555 = m.ExcPending
	if v555 != 0 {
		goto L2
	} else {
		goto L131
	}
L128:
	;
	v545 = v542
	goto L130
L129:
	;
	v545 = v540
	goto L130
L130:
	;
	goto L127
L131:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v89)+60)) = v554
	v557 = F_build_joinrel_restrictlist(m, l0, v89, v47, v48, v73)
	mBase = m.M
	v558 = m.ExcPending
	if v558 != 0 {
		goto L2
	} else {
		goto L132
	}
L132:
	;
	if v76 != 0 {
		goto L133
	} else {
		goto L134
	}
L133:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v76))) = v557
	goto L135
L134:
	;
	goto L135
L135:
	;
	v561 = v89 + int32(8)
	v562 = *(*int32)(unsafe.Add(mBase, uint32(v47)+212))
	if v562 == int32(0) {
		goto L137
	} else {
		goto L138
	}
L136:
	;
	v678 = *(*int32)(unsafe.Add(mBase, uint32(v48)+212))
	if v678 == int32(0) {
		v777 = v664
		goto L164
	} else {
		goto L165
	}
L137:
	;
	v664 = int32(0)
	goto L136
L138:
	;
	goto L139
L139:
	;
	v566 = int32(0)
	v567 = *(*int32)(unsafe.Add(mBase, uint32(v562)+4))
	if v567 <= v566 {
		goto L140
	} else {
		goto L141
	}
L140:
	;
	v664 = int32(0)
	goto L136
L141:
	;
	goto L142
L142:
	;
	v576 = int32(0)
	v579 = v566
	goto L143
L143:
	;
	v590 = *(*int32)(unsafe.Add(mBase, uint32(v562)+12))
	v594 = *(*int32)(unsafe.Add(mBase, uint32(v590+v579<<(uint(int32(2))%32))))
	v595 = *(*int32)(unsafe.Add(mBase, uint32(v594)+32))
	v596 = *(*int32)(unsafe.Add(mBase, uint32(v561)))
	v597 = int32(0)
	if v595 == v597 {
		goto L146
	} else {
		goto L147
	}
L144:
	;
	v664 = v655
	goto L136
L145:
	;
	if v650 == int32(0) {
		goto L159
	} else {
		goto L160
	}
L146:
	;
	v650 = int32(1)
	goto L145
L147:
	;
	goto L148
L148:
	;
	if v596 == int32(0) {
		v641 = v597
		goto L149
	} else {
		goto L150
	}
L149:
	;
	v650 = v641
	goto L145
L150:
	;
	v606 = *(*int32)(unsafe.Add(mBase, uint32(v595)+4))
	v607 = *(*int32)(unsafe.Add(mBase, uint32(v596)+4))
	if v607 < v606 {
		v641 = v597
		goto L149
	} else {
		goto L151
	}
L151:
	;
	v609 = int32(1)
	if v606 <= v609 {
		goto L152
	} else {
		goto L153
	}
L152:
	;
	v612 = v609
	goto L154
L153:
	;
	v612 = v606
	goto L154
L154:
	;
	v613 = int32(8)
	v618 = int32(0)
	goto L155
L155:
	;
	v625 = v618 << (uint(int32(2)) % 32)
	v627 = *(*int32)(unsafe.Add(mBase, uint32(v595+v613+v625)))
	v629 = *(*int32)(unsafe.Add(mBase, uint32(v625+(v596+v613))))
	v632 = v627 & (v629 ^ int32(-1))
	v634 = base.B2i32(v632 == int32(0))
	if v632 != 0 {
		v641 = v634
		goto L149
	} else {
		goto L157
	}
L156:
	;
	v641 = v634
	goto L149
L157:
	;
	v636 = v618 + int32(1)
	if v636 != v612 {
		v618 = v636
		goto L155
	} else {
		goto L158
	}
L158:
	;
	goto L156
L159:
	;
	v653 = F_list_append_unique_ptr(m, v576, v594)
	mBase = m.M
	v654 = m.ExcPending
	if v654 != 0 {
		goto L2
	} else {
		goto L162
	}
L160:
	;
	v655 = v576
	goto L161
L161:
	;
	v657 = v579 + int32(1)
	v658 = *(*int32)(unsafe.Add(mBase, uint32(v562)+4))
	if v657 < v658 {
		v576 = v655
		v579 = v657
		goto L143
	} else {
		goto L163
	}
L162:
	;
	v655 = v653
	goto L161
L163:
	;
	goto L144
L164:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v89)+212)) = v777
	v792 = int32(0)
	v793 = *(*int32)(unsafe.Add(mBase, uint32(v89)+8))
	if v793 == v792 {
		goto L191
	} else {
		goto L192
	}
L165:
	;
	v681 = *(*int32)(unsafe.Add(mBase, uint32(v678)+4))
	if v681 <= int32(0) {
		v777 = v664
		goto L164
	} else {
		goto L166
	}
L166:
	;
	v689 = v664
	v692 = int32(0)
	goto L167
L167:
	;
	v703 = *(*int32)(unsafe.Add(mBase, uint32(v678)+12))
	v707 = *(*int32)(unsafe.Add(mBase, uint32(v703+v692<<(uint(int32(2))%32))))
	v708 = *(*int32)(unsafe.Add(mBase, uint32(v707)+32))
	v709 = *(*int32)(unsafe.Add(mBase, uint32(v561)))
	v710 = int32(0)
	if v708 == v710 {
		goto L170
	} else {
		goto L171
	}
L168:
	;
	v777 = v768
	goto L164
L169:
	;
	if v763 == int32(0) {
		goto L183
	} else {
		goto L184
	}
L170:
	;
	v763 = int32(1)
	goto L169
L171:
	;
	goto L172
L172:
	;
	if v709 == int32(0) {
		v754 = v710
		goto L173
	} else {
		goto L174
	}
L173:
	;
	v763 = v754
	goto L169
L174:
	;
	v719 = *(*int32)(unsafe.Add(mBase, uint32(v708)+4))
	v720 = *(*int32)(unsafe.Add(mBase, uint32(v709)+4))
	if v720 < v719 {
		v754 = v710
		goto L173
	} else {
		goto L175
	}
L175:
	;
	v722 = int32(1)
	if v719 <= v722 {
		goto L176
	} else {
		goto L177
	}
L176:
	;
	v725 = v722
	goto L178
L177:
	;
	v725 = v719
	goto L178
L178:
	;
	v726 = int32(8)
	v731 = int32(0)
	goto L179
L179:
	;
	v738 = v731 << (uint(int32(2)) % 32)
	v740 = *(*int32)(unsafe.Add(mBase, uint32(v708+v726+v738)))
	v742 = *(*int32)(unsafe.Add(mBase, uint32(v738+(v709+v726))))
	v745 = v740 & (v742 ^ int32(-1))
	v747 = base.B2i32(v745 == int32(0))
	if v745 != 0 {
		v754 = v747
		goto L173
	} else {
		goto L181
	}
L180:
	;
	v754 = v747
	goto L173
L181:
	;
	v749 = v731 + int32(1)
	if v749 != v725 {
		v731 = v749
		goto L179
	} else {
		goto L182
	}
L182:
	;
	goto L180
L183:
	;
	v766 = F_list_append_unique_ptr(m, v689, v707)
	mBase = m.M
	v767 = m.ExcPending
	if v767 != 0 {
		goto L2
	} else {
		goto L186
	}
L184:
	;
	v768 = v689
	goto L185
L185:
	;
	v770 = v692 + int32(1)
	v771 = *(*int32)(unsafe.Add(mBase, uint32(v678)+4))
	if v770 < v771 {
		v689 = v768
		v692 = v770
		goto L167
	} else {
		goto L187
	}
L186:
	;
	v768 = v766
	goto L185
L187:
	;
	goto L168
L188:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v89)+216)) = uint8(v1202)
	F_build_joinrel_partition_info(m, l0, v89, v47, v48, v73, v557)
	mBase = m.M
	v1205 = m.ExcPending
	if v1205 != 0 {
		goto L2
	} else {
		goto L267
	}
L189:
	;
	if int32(0) < v850 {
		goto L200
	} else {
		goto L201
	}
L190:
	;
	v850 = base.I32_ctz(v836) | v837<<(uint(int32(5))%32)
	goto L189
L191:
	;
	v850 = int32(-2)
	goto L189
L192:
	;
	v803 = base.I32_div_s(int32(0), int32(32))
	v804 = *(*int32)(unsafe.Add(mBase, uint32(v793)+4))
	if v804 <= v803 {
		goto L191
	} else {
		goto L193
	}
L193:
	;
	v807 = v793 + int32(8)
	v811 = *(*int32)(unsafe.Add(mBase, uint32(v807+v803<<(uint(int32(2))%32))))
	v814 = v811 & int32(-1)
	if v814 != 0 {
		v836 = v814
		v837 = v803
		goto L190
	} else {
		goto L194
	}
L194:
	;
	v816 = v803 + int32(1)
	if v816 == v804 {
		goto L191
	} else {
		goto L195
	}
L195:
	;
	v819 = v816
	goto L196
L196:
	;
	v826 = *(*int32)(unsafe.Add(mBase, uint32(v807+v819<<(uint(int32(2))%32))))
	if v826 != 0 {
		v836 = v826
		v837 = v819
		goto L190
	} else {
		goto L198
	}
L197:
	;
	goto L191
L198:
	;
	v828 = v819 + int32(1)
	if v828 != v804 {
		v819 = v828
		goto L196
	} else {
		goto L199
	}
L199:
	;
	goto L197
L200:
	;
	v857 = v792
	v858 = v850
	goto L203
L201:
	;
	v947 = v792
	goto L202
L202:
	;
	if v947 == int32(0) {
		goto L223
	} else {
		goto L224
	}
L203:
	;
	v871 = *(*int32)(unsafe.Add(mBase, uint32(l0)+324))
	if v858 == v871 {
		v883 = v857
		goto L205
	} else {
		goto L206
	}
L204:
	;
	v947 = v883
	goto L202
L205:
	;
	if v793 == int32(0) {
		goto L211
	} else {
		goto L212
	}
L206:
	;
	v873 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v877 = *(*int32)(unsafe.Add(mBase, uint32(v873+v858<<(uint(int32(2))%32))))
	if v877 == int32(0) {
		v883 = v857
		goto L205
	} else {
		goto L207
	}
L207:
	;
	v880 = *(*int32)(unsafe.Add(mBase, uint32(v877)+136))
	v881 = F_bms_add_members(m, v857, v880)
	mBase = m.M
	v882 = m.ExcPending
	if v882 != 0 {
		goto L2
	} else {
		goto L208
	}
L208:
	;
	v883 = v881
	goto L205
L209:
	;
	if int32(0) < v940 {
		v857 = v883
		v858 = v940
		goto L203
	} else {
		goto L220
	}
L210:
	;
	v940 = base.I32_ctz(v926) | v927<<(uint(int32(5))%32)
	goto L209
L211:
	;
	v940 = int32(-2)
	goto L209
L212:
	;
	v891 = v858 + int32(1)
	v893 = base.I32_div_s(v891, int32(32))
	v894 = *(*int32)(unsafe.Add(mBase, uint32(v793)+4))
	if v894 <= v893 {
		goto L211
	} else {
		goto L213
	}
L213:
	;
	v897 = v793 + int32(8)
	v901 = *(*int32)(unsafe.Add(mBase, uint32(v897+v893<<(uint(int32(2))%32))))
	v904 = v901 & (int32(-1) << (uint(v891) % 32))
	if v904 != 0 {
		v926 = v904
		v927 = v893
		goto L210
	} else {
		goto L214
	}
L214:
	;
	v906 = v893 + int32(1)
	if v906 == v894 {
		goto L211
	} else {
		goto L215
	}
L215:
	;
	v909 = v906
	goto L216
L216:
	;
	v916 = *(*int32)(unsafe.Add(mBase, uint32(v897+v909<<(uint(int32(2))%32))))
	if v916 != 0 {
		v926 = v916
		v927 = v909
		goto L210
	} else {
		goto L218
	}
L217:
	;
	goto L211
L218:
	;
	v918 = v909 + int32(1)
	if v918 != v894 {
		v909 = v918
		goto L216
	} else {
		goto L219
	}
L219:
	;
	goto L217
L220:
	;
	goto L204
L221:
	;
	if int32(0) <= v1017 {
		goto L232
	} else {
		goto L233
	}
L222:
	;
	v1017 = base.I32_ctz(v1003) | v1004<<(uint(int32(5))%32)
	goto L221
L223:
	;
	v1017 = int32(-2)
	goto L221
L224:
	;
	v970 = base.I32_div_s(int32(0), int32(32))
	v971 = *(*int32)(unsafe.Add(mBase, uint32(v947)+4))
	if v971 <= v970 {
		goto L223
	} else {
		goto L225
	}
L225:
	;
	v974 = v947 + int32(8)
	v978 = *(*int32)(unsafe.Add(mBase, uint32(v974+v970<<(uint(int32(2))%32))))
	v981 = v978 & int32(-1)
	if v981 != 0 {
		v1003 = v981
		v1004 = v970
		goto L222
	} else {
		goto L226
	}
L226:
	;
	v983 = v970 + int32(1)
	if v983 == v971 {
		goto L223
	} else {
		goto L227
	}
L227:
	;
	v986 = v983
	goto L228
L228:
	;
	v993 = *(*int32)(unsafe.Add(mBase, uint32(v974+v986<<(uint(int32(2))%32))))
	if v993 != 0 {
		v1003 = v993
		v1004 = v986
		goto L222
	} else {
		goto L230
	}
L229:
	;
	goto L223
L230:
	;
	v995 = v986 + int32(1)
	if v995 != v971 {
		v986 = v995
		goto L228
	} else {
		goto L231
	}
L231:
	;
	goto L229
L232:
	;
	v1025 = v1017
	goto L235
L233:
	;
	goto L234
L234:
	;
	v1202 = int32(0)
	goto L188
L235:
	;
	v1038 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	v1039 = *(*int32)(unsafe.Add(mBase, uint32(v1038)+12))
	v1043 = *(*int32)(unsafe.Add(mBase, uint32(v1039+v1025<<(uint(int32(2))%32))))
	v1044 = *(*int32)(unsafe.Add(mBase, uint32(v1043)+16))
	if v1044 == int32(0) {
		goto L237
	} else {
		goto L238
	}
L236:
	;
	goto L234
L237:
	;
	if v947 == int32(0) {
		goto L257
	} else {
		goto L258
	}
L238:
	;
	v1047 = *(*int32)(unsafe.Add(mBase, uint32(v1044)+4))
	if v1047 < int32(2) {
		goto L237
	} else {
		goto L239
	}
L239:
	;
	v1050 = *(*int32)(unsafe.Add(mBase, uint32(v1043)+36))
	v1051 = *(*int32)(unsafe.Add(mBase, uint32(v89)+8))
	v1052 = int32(0)
	if v1050 == v1052 {
		goto L241
	} else {
		goto L242
	}
L240:
	;
	if v1105 != 0 {
		goto L237
	} else {
		goto L254
	}
L241:
	;
	v1105 = int32(1)
	goto L240
L242:
	;
	goto L243
L243:
	;
	if v1051 == int32(0) {
		v1096 = v1052
		goto L244
	} else {
		goto L245
	}
L244:
	;
	v1105 = v1096
	goto L240
L245:
	;
	v1061 = *(*int32)(unsafe.Add(mBase, uint32(v1050)+4))
	v1062 = *(*int32)(unsafe.Add(mBase, uint32(v1051)+4))
	if v1062 < v1061 {
		v1096 = v1052
		goto L244
	} else {
		goto L246
	}
L246:
	;
	v1064 = int32(1)
	if v1061 <= v1064 {
		goto L247
	} else {
		goto L248
	}
L247:
	;
	v1067 = v1064
	goto L249
L248:
	;
	v1067 = v1061
	goto L249
L249:
	;
	v1068 = int32(8)
	v1073 = int32(0)
	goto L250
L250:
	;
	v1080 = v1073 << (uint(int32(2)) % 32)
	v1082 = *(*int32)(unsafe.Add(mBase, uint32(v1050+v1068+v1080)))
	v1084 = *(*int32)(unsafe.Add(mBase, uint32(v1080+(v1051+v1068))))
	v1087 = v1082 & (v1084 ^ int32(-1))
	v1089 = base.B2i32(v1087 == int32(0))
	if v1087 != 0 {
		v1096 = v1089
		goto L244
	} else {
		goto L252
	}
L251:
	;
	v1096 = v1089
	goto L244
L252:
	;
	v1091 = v1073 + int32(1)
	if v1091 != v1067 {
		v1073 = v1091
		goto L250
	} else {
		goto L253
	}
L253:
	;
	goto L251
L254:
	;
	v1202 = int32(1)
	goto L188
L255:
	;
	if int32(0) <= v1162 {
		v1025 = v1162
		goto L235
	} else {
		goto L266
	}
L256:
	;
	v1162 = base.I32_ctz(v1148) | v1149<<(uint(int32(5))%32)
	goto L255
L257:
	;
	v1162 = int32(-2)
	goto L255
L258:
	;
	v1113 = v1025 + int32(1)
	v1115 = base.I32_div_s(v1113, int32(32))
	v1116 = *(*int32)(unsafe.Add(mBase, uint32(v947)+4))
	if v1116 <= v1115 {
		goto L257
	} else {
		goto L259
	}
L259:
	;
	v1119 = v947 + int32(8)
	v1123 = *(*int32)(unsafe.Add(mBase, uint32(v1119+v1115<<(uint(int32(2))%32))))
	v1126 = v1123 & (int32(-1) << (uint(v1113) % 32))
	if v1126 != 0 {
		v1148 = v1126
		v1149 = v1115
		goto L256
	} else {
		goto L260
	}
L260:
	;
	v1128 = v1115 + int32(1)
	if v1128 == v1116 {
		goto L257
	} else {
		goto L261
	}
L261:
	;
	v1131 = v1128
	goto L262
L262:
	;
	v1138 = *(*int32)(unsafe.Add(mBase, uint32(v1119+v1131<<(uint(int32(2))%32))))
	if v1138 != 0 {
		v1148 = v1138
		v1149 = v1131
		goto L256
	} else {
		goto L264
	}
L263:
	;
	goto L257
L264:
	;
	v1140 = v1131 + int32(1)
	if v1140 != v1116 {
		v1131 = v1140
		goto L262
	} else {
		goto L265
	}
L265:
	;
	goto L263
L266:
	;
	goto L236
L267:
	;
	F_set_joinrel_size_estimates(m, l0, v89, v47, v48, v73, v557)
	mBase = m.M
	v1207 = m.ExcPending
	if v1207 != 0 {
		goto L2
	} else {
		goto L268
	}
L268:
	;
	v1208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+26)))
	if v1208 != int32(1) {
		goto L269
	} else {
		goto L270
	}
L269:
	;
	v1226 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v1227 = F_lappend(m, v1226, v89)
	mBase = m.M
	v1228 = m.ExcPending
	if v1228 != 0 {
		goto L2
	} else {
		goto L276
	}
L270:
	;
	v1211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+26)))
	if v1211 != int32(1) {
		goto L269
	} else {
		goto L271
	}
L271:
	;
	v1214 = F_is_parallel_safe(m, l0, v557)
	mBase = m.M
	v1215 = m.ExcPending
	if v1215 != 0 {
		goto L2
	} else {
		goto L272
	}
L272:
	;
	if v1214 == int32(0) {
		goto L269
	} else {
		goto L273
	}
L273:
	;
	v1218 = *(*int32)(unsafe.Add(mBase, uint32(v89)+28))
	v1219 = *(*int32)(unsafe.Add(mBase, uint32(v1218)+4))
	v1220 = F_is_parallel_safe(m, l0, v1219)
	mBase = m.M
	v1221 = m.ExcPending
	if v1221 != 0 {
		goto L2
	} else {
		goto L274
	}
L274:
	;
	if v1220 == int32(0) {
		goto L269
	} else {
		goto L275
	}
L275:
	;
	v1224 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v89)+26)) = uint8(v1224)
	goto L269
L276:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = v1227
	v1230 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	if v1230 != 0 {
		goto L277
	} else {
		goto L278
	}
L277:
	;
	v1234 = F_hash_search(m, v1230, v561, int32(1), v79+int32(15))
	mBase = m.M
	v1235 = m.ExcPending
	if v1235 != 0 {
		goto L2
	} else {
		goto L280
	}
L278:
	;
	goto L279
L279:
	;
	v1237 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	if v1237 == int32(0) {
		v1256 = v89
		goto L19
	} else {
		goto L281
	}
L280:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1234)+4)) = v89
	goto L279
L281:
	;
	v1240 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v1244 = *(*int32)(unsafe.Add(mBase, uint32(v1237+v1240<<(uint(int32(2))%32))))
	v1245 = F_lappend(m, v1244, v89)
	mBase = m.M
	v1246 = m.ExcPending
	if v1246 != 0 {
		goto L2
	} else {
		goto L282
	}
L282:
	;
	v1247 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v1248 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v1247+v1248<<(uint(int32(2))%32)))) = v1245
	v1256 = v89
	goto L19
L283:
	;
	F_bms_free(m, v44)
	mBase = m.M
	v1349 = m.ExcPending
	if v1349 != 0 {
		goto L2
	} else {
		goto L294
	}
L284:
	;
	v1327 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
	F_populate_joinrel_with_paths(m, l0, v47, v48, v1256, v73, v1327)
	mBase = m.M
	v1329 = m.ExcPending
	if v1329 != 0 {
		goto L2
	} else {
		goto L293
	}
L285:
	;
	v1277 = *(*int32)(unsafe.Add(mBase, uint32(v1274)+12))
	v1280 = v1277
	goto L286
L286:
	;
	v1296 = *(*int32)(unsafe.Add(mBase, uint32(v1280)))
	v1297 = *(*int32)(unsafe.Add(mBase, uint32(v1296)))
	if base.Ui32(int32(2)) <= base.Ui32(v1297-int32(301)) {
		goto L288
	} else {
		goto L289
	}
L287:
	;
	goto L284
L288:
	;
	if v1297 != int32(290) {
		goto L284
	} else {
		goto L291
	}
L289:
	;
	v1280 = v1296 + int32(72)
	goto L286
L290:
	;
	goto L287
L291:
	;
	v1304 = *(*int32)(unsafe.Add(mBase, uint32(v1296)+72))
	if v1304 == int32(0) {
		goto L283
	} else {
		goto L292
	}
L292:
	;
	goto L290
L293:
	;
	goto L283
L294:
	;
	v1355 = v1256
	goto L1
}
func F_remove_join_clause_from_rels(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v62 int32
	_ = v62
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
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v131 int32
	_ = v131
	if l2 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	if int32(0) <= v62 {
		goto L12
	} else {
		goto L13
	}
L2:
	;
	v62 = base.I32_ctz(v48) | v49<<(uint(int32(5))%32)
	goto L1
L3:
	;
	v62 = int32(-2)
	goto L1
L4:
	;
	v15 = base.I32_div_s(int32(0), int32(32))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v16 <= v15 {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v19 = l2 + int32(8)
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v19+v15<<(uint(int32(2))%32))))
	v26 = v23 & int32(-1)
	if v26 != 0 {
		v48 = v26
		v49 = v15
		goto L2
	} else {
		goto L6
	}
L6:
	;
	v28 = v15 + int32(1)
	if v28 == v16 {
		goto L3
	} else {
		goto L7
	}
L7:
	;
	v31 = v28
	goto L8
L8:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v19+v31<<(uint(int32(2))%32))))
	if v38 != 0 {
		v48 = v38
		v49 = v31
		goto L2
	} else {
		goto L10
	}
L9:
	;
	goto L3
L10:
	;
	v40 = v31 + int32(1)
	if v40 != v16 {
		v31 = v40
		goto L8
	} else {
		goto L11
	}
L11:
	;
	goto L9
L12:
	;
	v68 = v62
	goto L15
L13:
	;
	goto L14
L14:
	;
	return
L15:
	;
	v70 = F_find_base_rel_ignore_join(m, l0, v68)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	goto L14
L17:
	;
	return
L18:
	;
	if v70 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v70)+212))
	v73 = F_list_delete_ptr(m, v72, l1)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L17
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	if l2 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v70)+212)) = v73
	goto L21
L23:
	;
	if int32(0) <= v131 {
		v68 = v131
		goto L15
	} else {
		goto L34
	}
L24:
	;
	v131 = base.I32_ctz(v117) | v118<<(uint(int32(5))%32)
	goto L23
L25:
	;
	v131 = int32(-2)
	goto L23
L26:
	;
	v82 = v68 + int32(1)
	v84 = base.I32_div_s(v82, int32(32))
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v85 <= v84 {
		goto L25
	} else {
		goto L27
	}
L27:
	;
	v88 = l2 + int32(8)
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v88+v84<<(uint(int32(2))%32))))
	v95 = v92 & (int32(-1) << (uint(v82) % 32))
	if v95 != 0 {
		v117 = v95
		v118 = v84
		goto L24
	} else {
		goto L28
	}
L28:
	;
	v97 = v84 + int32(1)
	if v97 == v85 {
		goto L25
	} else {
		goto L29
	}
L29:
	;
	v100 = v97
	goto L30
L30:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v88+v100<<(uint(int32(2))%32))))
	if v107 != 0 {
		v117 = v107
		v118 = v100
		goto L24
	} else {
		goto L32
	}
L31:
	;
	goto L25
L32:
	;
	v109 = v100 + int32(1)
	if v109 != v85 {
		v100 = v109
		goto L30
	} else {
		goto L33
	}
L33:
	;
	goto L31
L34:
	;
	goto L16
}
