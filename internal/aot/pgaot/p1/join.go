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
	var v105 int64
	_ = v105
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int64
	_ = v109
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
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
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v138 int64
	_ = v138
	var v160 int64
	_ = v160
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v233 int64
	_ = v233
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v257 int64
	_ = v257
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v286 int32
	_ = v286
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v311 int32
	_ = v311
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v327 int32
	_ = v327
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v342 int32
	_ = v342
	var v349 int32
	_ = v349
	var v351 int32
	_ = v351
	var v353 int32
	_ = v353
	var v356 int32
	_ = v356
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v364 int32
	_ = v364
	var v374 int32
	_ = v374
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v391 int32
	_ = v391
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v400 int32
	_ = v400
	var v407 int32
	_ = v407
	var v409 int32
	_ = v409
	var v411 int32
	_ = v411
	var v414 int32
	_ = v414
	var v416 int32
	_ = v416
	var v418 int32
	_ = v418
	var v425 int32
	_ = v425
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v447 int32
	_ = v447
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v456 int32
	_ = v456
	var v463 int32
	_ = v463
	var v465 int32
	_ = v465
	var v467 int32
	_ = v467
	var v470 int32
	_ = v470
	var v472 int32
	_ = v472
	var v474 int32
	_ = v474
	var v481 int32
	_ = v481
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
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v498 int32
	_ = v498
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v502 float64
	_ = v502
	var v503 float64
	_ = v503
	var v506 int32
	_ = v506
	var v507 float64
	_ = v507
	var v508 float64
	_ = v508
	var v511 int64
	_ = v511
	var v515 int64
	_ = v515
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v523 int64
	_ = v523
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v544 int64
	_ = v544
	var v546 int64
	_ = v546
	var v549 int64
	_ = v549
	var v551 int32
	_ = v551
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v593 int32
	_ = v593
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v600 int32
	_ = v600
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
	var v612 int32
	_ = v612
	var v615 int32
	_ = v615
	var v616 int32
	_ = v616
	var v621 int32
	_ = v621
	var v628 int32
	_ = v628
	var v630 int32
	_ = v630
	var v632 int32
	_ = v632
	var v635 int32
	_ = v635
	var v637 int32
	_ = v637
	var v639 int32
	_ = v639
	var v646 int32
	_ = v646
	var v653 int32
	_ = v653
	var v656 int32
	_ = v656
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v660 int32
	_ = v660
	var v661 int32
	_ = v661
	var v665 int32
	_ = v665
	var v681 int32
	_ = v681
	var v684 int32
	_ = v684
	var v690 int32
	_ = v690
	var v691 int32
	_ = v691
	var v706 int32
	_ = v706
	var v710 int32
	_ = v710
	var v711 int32
	_ = v711
	var v712 int32
	_ = v712
	var v713 int32
	_ = v713
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v725 int32
	_ = v725
	var v728 int32
	_ = v728
	var v729 int32
	_ = v729
	var v734 int32
	_ = v734
	var v741 int32
	_ = v741
	var v743 int32
	_ = v743
	var v745 int32
	_ = v745
	var v748 int32
	_ = v748
	var v750 int32
	_ = v750
	var v752 int32
	_ = v752
	var v759 int32
	_ = v759
	var v766 int32
	_ = v766
	var v769 int32
	_ = v769
	var v770 int32
	_ = v770
	var v771 int32
	_ = v771
	var v773 int32
	_ = v773
	var v774 int32
	_ = v774
	var v778 int32
	_ = v778
	var v795 int32
	_ = v795
	var v796 int32
	_ = v796
	var v804 int32
	_ = v804
	var v807 int32
	_ = v807
	var v810 int32
	_ = v810
	var v814 int32
	_ = v814
	var v817 int32
	_ = v817
	var v818 int32
	_ = v818
	var v822 int32
	_ = v822
	var v829 int32
	_ = v829
	var v831 int32
	_ = v831
	var v839 int32
	_ = v839
	var v840 int32
	_ = v840
	var v853 int32
	_ = v853
	var v858 int32
	_ = v858
	var v859 int32
	_ = v859
	var v874 int32
	_ = v874
	var v876 int32
	_ = v876
	var v880 int32
	_ = v880
	var v883 int32
	_ = v883
	var v884 int32
	_ = v884
	var v885 int32
	_ = v885
	var v886 int32
	_ = v886
	var v894 int32
	_ = v894
	var v896 int32
	_ = v896
	var v897 int32
	_ = v897
	var v900 int32
	_ = v900
	var v904 int32
	_ = v904
	var v907 int32
	_ = v907
	var v909 int32
	_ = v909
	var v912 int32
	_ = v912
	var v919 int32
	_ = v919
	var v921 int32
	_ = v921
	var v929 int32
	_ = v929
	var v930 int32
	_ = v930
	var v943 int32
	_ = v943
	var v948 int32
	_ = v948
	var v971 int32
	_ = v971
	var v974 int32
	_ = v974
	var v977 int32
	_ = v977
	var v981 int32
	_ = v981
	var v984 int32
	_ = v984
	var v985 int32
	_ = v985
	var v989 int32
	_ = v989
	var v996 int32
	_ = v996
	var v998 int32
	_ = v998
	var v1006 int32
	_ = v1006
	var v1007 int32
	_ = v1007
	var v1020 int32
	_ = v1020
	var v1026 int32
	_ = v1026
	var v1041 int32
	_ = v1041
	var v1042 int32
	_ = v1042
	var v1046 int32
	_ = v1046
	var v1047 int32
	_ = v1047
	var v1050 int32
	_ = v1050
	var v1053 int32
	_ = v1053
	var v1054 int32
	_ = v1054
	var v1055 int32
	_ = v1055
	var v1064 int32
	_ = v1064
	var v1065 int32
	_ = v1065
	var v1067 int32
	_ = v1067
	var v1070 int32
	_ = v1070
	var v1071 int32
	_ = v1071
	var v1076 int32
	_ = v1076
	var v1083 int32
	_ = v1083
	var v1085 int32
	_ = v1085
	var v1087 int32
	_ = v1087
	var v1090 int32
	_ = v1090
	var v1092 int32
	_ = v1092
	var v1094 int32
	_ = v1094
	var v1101 int32
	_ = v1101
	var v1108 int32
	_ = v1108
	var v1116 int32
	_ = v1116
	var v1118 int32
	_ = v1118
	var v1119 int32
	_ = v1119
	var v1122 int32
	_ = v1122
	var v1126 int32
	_ = v1126
	var v1129 int32
	_ = v1129
	var v1131 int32
	_ = v1131
	var v1134 int32
	_ = v1134
	var v1141 int32
	_ = v1141
	var v1143 int32
	_ = v1143
	var v1151 int32
	_ = v1151
	var v1152 int32
	_ = v1152
	var v1165 int32
	_ = v1165
	var v1205 int32
	_ = v1205
	var v1208 int32
	_ = v1208
	var v1209 int32
	_ = v1209
	var v1212 int32
	_ = v1212
	var v1215 int32
	_ = v1215
	var v1216 int32
	_ = v1216
	var v1219 int32
	_ = v1219
	var v1220 int32
	_ = v1220
	var v1221 int32
	_ = v1221
	var v1222 int32
	_ = v1222
	var v1225 int32
	_ = v1225
	var v1228 int32
	_ = v1228
	var v1230 int32
	_ = v1230
	var v1232 int32
	_ = v1232
	var v1233 int32
	_ = v1233
	var v1234 int32
	_ = v1234
	var v1235 int32
	_ = v1235
	var v1237 int32
	_ = v1237
	var v1241 int32
	_ = v1241
	var v1242 int32
	_ = v1242
	var v1244 int32
	_ = v1244
	var v1247 int32
	_ = v1247
	var v1251 int32
	_ = v1251
	var v1252 int32
	_ = v1252
	var v1253 int32
	_ = v1253
	var v1254 int32
	_ = v1254
	var v1255 int32
	_ = v1255
	var v1261 int32
	_ = v1261
	var v1281 int32
	_ = v1281
	var v1284 int32
	_ = v1284
	var v1287 int32
	_ = v1287
	var v1303 int32
	_ = v1303
	var v1304 int32
	_ = v1304
	var v1311 int32
	_ = v1311
	var v1334 int32
	_ = v1334
	var v1336 int32
	_ = v1336
	var v1337 int32
	_ = v1337
	var v1339 int32
	_ = v1339
	var v1359 int32
	_ = v1359
	var v1363 int32
	_ = v1363
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
	return v1363
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
	v1363 = v4
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
	*(*int64)(unsafe.Add(mBase, uint32(v21)+60)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v21)+28)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v21)+24)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v21)+20)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v21)+16)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v21)+12)) = int32(322)
	*(*int64)(unsafe.Add(mBase, uint32(v21)+32)) = v53
	*(*int64)(unsafe.Add(mBase, uint32(v21)+40)) = v53
	*(*int64)(unsafe.Add(mBase, uint32(v21)+48)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v21)+55)) = int32(0)
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
	v1281 = *(*int32)(unsafe.Add(mBase, uint32(v1261)+44))
	if v1281 == int32(0) {
		goto L285
	} else {
		goto L286
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
		v1261 = v81
		goto L19
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v89 = F_palloc0(m, int32(304))
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
	v1261 = v81
	goto L19
L26:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v89))) = int64(4294967566)
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
	v98 = *(*float64)(unsafe.Add(mBase, uint32(l0)+312))
	v99 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v89)+25)) = uint16(v99)
	v102 = base.F64_gt(v98, float64(0))
	*(*uint8)(unsafe.Add(mBase, uint32(v89)+24)) = uint8(v102)
	v104 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v105 = *(*int64)(unsafe.Add(mBase, uint32(v104)+104))
	*(*int64)(unsafe.Add(mBase, uint32(v89)+32)) = v105
	v107 = F_create_empty_pathtarget(m)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L2
	} else {
		goto L28
	}
L28:
	;
	v109 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v89)+44)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v89)+40)) = v107
	*(*int64)(unsafe.Add(mBase, uint32(v89)+52)) = v109
	*(*int64)(unsafe.Add(mBase, uint32(v89)+60)) = v109
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v47)+68))
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v48)+68))
	v118 = F_bms_union(m, v116, v117)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L2
	} else {
		goto L29
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v89)+68)) = v118
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v89)+8))
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v47)+72))
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v48)+72))
	v124 = F_bms_union(m, v122, v123)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L2
	} else {
		goto L30
	}
L30:
	;
	v126 = F_bms_del_members(m, v124, v121)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L2
	} else {
		goto L31
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v89)+84)) = int32(2)
	v130 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v89)+76)) = v130
	*(*int32)(unsafe.Add(mBase, uint32(v89)+72)) = v126
	base.MemoryFill(m, v89+int32(88), v130, int32(68))
	v138 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v89)+160)) = v138
	*(*int32)(unsafe.Add(mBase, uint32(v89)+156)) = int32(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v89)+165)) = v138
	*(*int64)(unsafe.Add(mBase, uint32(v89)+176)) = v138
	*(*int64)(unsafe.Add(mBase, uint32(v89)+184)) = v138
	*(*int64)(unsafe.Add(mBase, uint32(v89)+192)) = v138
	*(*int64)(unsafe.Add(mBase, uint32(v89)+200)) = v138
	*(*int64)(unsafe.Add(mBase, uint32(v89)+208)) = v138
	*(*int64)(unsafe.Add(mBase, uint32(v89)+216)) = v138
	*(*int64)(unsafe.Add(mBase, uint32(v89)+236)) = v138
	*(*uint16)(unsafe.Add(mBase, uint32(v89)+232)) = uint16(v130)
	v160 = int64(4294967295)
	*(*int64)(unsafe.Add(mBase, uint32(v89)+224)) = v160
	*(*int64)(unsafe.Add(mBase, uint32(v89)+244)) = v138
	*(*int64)(unsafe.Add(mBase, uint32(v89)+252)) = v138
	*(*uint8)(unsafe.Add(mBase, uint32(v89)+268)) = uint8(v130)
	*(*int64)(unsafe.Add(mBase, uint32(v89)+260)) = v160
	*(*int64)(unsafe.Add(mBase, uint32(v89)+288)) = v138
	*(*int64)(unsafe.Add(mBase, uint32(v89)+280)) = v138
	*(*int64)(unsafe.Add(mBase, uint32(v89)+272)) = v138
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v47)+164))
	if v176 == v130 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v73)+20))
	F_build_joinrel_tlist(m, l0, v89, v47, v73, v74, base.B2i32(v217 == int32(2)))
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L2
	} else {
		goto L51
	}
L33:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v48)+164))
	if v179 != v176 {
		goto L32
	} else {
		goto L34
	}
L34:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v48)+168))
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v47)+168))
	if v181 == v182 {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v89)+172)) = uint8(v211)
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v47)+176))
	*(*int32)(unsafe.Add(mBase, uint32(v89)+176)) = v213
	goto L32
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v89)+164)) = v176
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v47)+168))
	*(*int32)(unsafe.Add(mBase, uint32(v89)+168)) = v185
	v187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+172)))
	if v187 != 0 {
		goto L39
	} else {
		goto L40
	}
L37:
	;
	goto L38
L38:
	;
	if v181 != 0 {
		goto L43
	} else {
		goto L44
	}
L39:
	;
	v190 = int32(1)
	goto L41
L40:
	;
	v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+172)))
	v190 = v189
	goto L41
L41:
	;
	v211 = v190 & int32(1)
	goto L35
L42:
	;
	v211 = int32(1)
	goto L35
L43:
	;
	v201 = v182
	goto L45
L44:
	;
	v194 = *(*int32)(unsafe.Add(mBase, _c_F_make_join_rel[0]))
	if v182 == v194 {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	if v201 != 0 {
		goto L32
	} else {
		goto L49
	}
L46:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v47)+164))
	*(*int32)(unsafe.Add(mBase, uint32(v89)+164)) = v196
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v47)+168))
	*(*int32)(unsafe.Add(mBase, uint32(v89)+168)) = v198
	goto L42
L47:
	;
	goto L48
L48:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v47)+168))
	v201 = v200
	goto L45
L49:
	;
	v203 = *(*int32)(unsafe.Add(mBase, _c_F_make_join_rel[0]))
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v48)+168))
	if v203 != v204 {
		goto L32
	} else {
		goto L50
	}
L50:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v47)+164))
	*(*int32)(unsafe.Add(mBase, uint32(v89)+164)) = v206
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v48)+168))
	*(*int32)(unsafe.Add(mBase, uint32(v89)+168)) = v208
	goto L42
L51:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v73)+20))
	F_build_joinrel_tlist(m, l0, v89, v48, v73, v74, base.B2i32(v222 != int32(0)))
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L2
	} else {
		goto L52
	}
L52:
	;
	v227 = int32(0)
	v228 = m.G0
	v230 = v228 - int32(16)
	m.G0 = v230
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v89)+40))
	v233 = int64(*(*int32)(unsafe.Add(mBase, uint32(v232)+32)))
	v234 = *(*int32)(unsafe.Add(mBase, uint32(l0)+148))
	if v234 == v227 {
		v544 = v233
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v546 = int64(1073741823)
	if v546 <= v544 {
		goto L127
	} else {
		goto L128
	}
L54:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v234)+4))
	if v237 <= int32(0) {
		v544 = v233
		goto L53
	} else {
		goto L55
	}
L55:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v89)+8))
	v243 = v227
	v257 = v233
	goto L56
L56:
	;
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v234)+12))
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v259+v243<<(uint(int32(2))%32))))
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v263)+12))
	v265 = int32(0)
	if v264 == v265 {
		goto L59
	} else {
		goto L60
	}
L57:
	;
	v544 = v523
	goto L53
L58:
	;
	if v318 != 0 {
		goto L72
	} else {
		goto L73
	}
L59:
	;
	v318 = int32(1)
	goto L58
L60:
	;
	goto L61
L61:
	;
	if v240 == int32(0) {
		v311 = v265
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v318 = v311
	goto L58
L63:
	;
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v264)+4))
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v240)+4))
	if v275 < v274 {
		v311 = v265
		goto L62
	} else {
		goto L64
	}
L64:
	;
	v277 = int32(1)
	if v274 <= v277 {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v280 = v277
	goto L67
L66:
	;
	v280 = v274
	goto L67
L67:
	;
	v281 = int32(8)
	v286 = int32(0)
	goto L68
L68:
	;
	v293 = v286 << (uint(int32(2)) % 32)
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v264+v281+v293)))
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v240+v281+v293)))
	v300 = v295 & (v297 ^ int32(-1))
	v302 = base.B2i32(v300 == int32(0))
	if v300 != 0 {
		v311 = v302
		goto L62
	} else {
		goto L70
	}
L69:
	;
	v311 = v302
	goto L62
L70:
	;
	v304 = v286 + int32(1)
	if v304 != v280 {
		v286 = v304
		goto L68
	} else {
		goto L71
	}
L71:
	;
	goto L69
L72:
	;
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v263)+20))
	if v319 == int32(0) {
		goto L77
	} else {
		goto L78
	}
L73:
	;
	v523 = v257
	goto L74
L74:
	;
	v525 = v243 + int32(1)
	v526 = *(*int32)(unsafe.Add(mBase, uint32(v234)+4))
	if v525 < v526 {
		v243 = v525
		v257 = v523
		goto L56
	} else {
		goto L125
	}
L75:
	;
	v516 = *(*int32)(unsafe.Add(mBase, uint32(v89)+68))
	v517 = *(*int32)(unsafe.Add(mBase, uint32(v263)+16))
	v518 = F_bms_add_members(m, v516, v517)
	mBase = m.M
	v519 = m.ExcPending
	if v519 != 0 {
		goto L2
	} else {
		goto L124
	}
L76:
	;
	if v374 == int32(0) {
		v515 = v257
		goto L75
	} else {
		goto L90
	}
L77:
	;
	v374 = int32(0)
	goto L76
L78:
	;
	goto L79
L79:
	;
	v327 = int32(1)
	if v240 == int32(0) {
		v364 = v327
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v374 = v364
	goto L76
L81:
	;
	v330 = *(*int32)(unsafe.Add(mBase, uint32(v319)+4))
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v240)+4))
	if v331 < v330 {
		v364 = v327
		goto L80
	} else {
		goto L82
	}
L82:
	;
	v333 = int32(1)
	if v330 <= v333 {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v336 = v333
	goto L85
L84:
	;
	v336 = v330
	goto L85
L85:
	;
	v337 = int32(8)
	v342 = int32(0)
	goto L86
L86:
	;
	v349 = v342 << (uint(int32(2)) % 32)
	v351 = *(*int32)(unsafe.Add(mBase, uint32(v319+v337+v349)))
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v240+v337+v349)))
	v356 = v351 & (v353 ^ int32(-1))
	v358 = base.B2i32(v356 != int32(0))
	if v356 != 0 {
		v364 = v358
		goto L80
	} else {
		goto L88
	}
L87:
	;
	v364 = v358
	goto L80
L88:
	;
	v360 = v342 + int32(1)
	if v360 != v336 {
		v342 = v360
		goto L86
	} else {
		goto L89
	}
L89:
	;
	goto L87
L90:
	;
	v377 = *(*int32)(unsafe.Add(mBase, uint32(v263)+12))
	v378 = *(*int32)(unsafe.Add(mBase, uint32(v47)+8))
	v379 = int32(0)
	if v377 == v379 {
		goto L92
	} else {
		goto L93
	}
L91:
	;
	if v432 != 0 {
		v515 = v257
		goto L75
	} else {
		goto L105
	}
L92:
	;
	v432 = int32(1)
	goto L91
L93:
	;
	goto L94
L94:
	;
	if v378 == int32(0) {
		v425 = v379
		goto L95
	} else {
		goto L96
	}
L95:
	;
	v432 = v425
	goto L91
L96:
	;
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v377)+4))
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v378)+4))
	if v389 < v388 {
		v425 = v379
		goto L95
	} else {
		goto L97
	}
L97:
	;
	v391 = int32(1)
	if v388 <= v391 {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v394 = v391
	goto L100
L99:
	;
	v394 = v388
	goto L100
L100:
	;
	v395 = int32(8)
	v400 = int32(0)
	goto L101
L101:
	;
	v407 = v400 << (uint(int32(2)) % 32)
	v409 = *(*int32)(unsafe.Add(mBase, uint32(v377+v395+v407)))
	v411 = *(*int32)(unsafe.Add(mBase, uint32(v378+v395+v407)))
	v414 = v409 & (v411 ^ int32(-1))
	v416 = base.B2i32(v414 == int32(0))
	if v414 != 0 {
		v425 = v416
		goto L95
	} else {
		goto L103
	}
L102:
	;
	v425 = v416
	goto L95
L103:
	;
	v418 = v400 + int32(1)
	if v418 != v394 {
		v400 = v418
		goto L101
	} else {
		goto L104
	}
L104:
	;
	goto L102
L105:
	;
	v433 = *(*int32)(unsafe.Add(mBase, uint32(v263)+12))
	v434 = *(*int32)(unsafe.Add(mBase, uint32(v48)+8))
	v435 = int32(0)
	if v433 == v435 {
		goto L107
	} else {
		goto L108
	}
L106:
	;
	if v488 != 0 {
		v515 = v257
		goto L75
	} else {
		goto L120
	}
L107:
	;
	v488 = int32(1)
	goto L106
L108:
	;
	goto L109
L109:
	;
	if v434 == int32(0) {
		v481 = v435
		goto L110
	} else {
		goto L111
	}
L110:
	;
	v488 = v481
	goto L106
L111:
	;
	v444 = *(*int32)(unsafe.Add(mBase, uint32(v433)+4))
	v445 = *(*int32)(unsafe.Add(mBase, uint32(v434)+4))
	if v445 < v444 {
		v481 = v435
		goto L110
	} else {
		goto L112
	}
L112:
	;
	v447 = int32(1)
	if v444 <= v447 {
		goto L113
	} else {
		goto L114
	}
L113:
	;
	v450 = v447
	goto L115
L114:
	;
	v450 = v444
	goto L115
L115:
	;
	v451 = int32(8)
	v456 = int32(0)
	goto L116
L116:
	;
	v463 = v456 << (uint(int32(2)) % 32)
	v465 = *(*int32)(unsafe.Add(mBase, uint32(v433+v451+v463)))
	v467 = *(*int32)(unsafe.Add(mBase, uint32(v434+v451+v463)))
	v470 = v465 & (v467 ^ int32(-1))
	v472 = base.B2i32(v470 == int32(0))
	if v470 != 0 {
		v481 = v472
		goto L110
	} else {
		goto L118
	}
L117:
	;
	v481 = v472
	goto L110
L118:
	;
	v474 = v456 + int32(1)
	if v474 != v450 {
		v456 = v474
		goto L116
	} else {
		goto L119
	}
L119:
	;
	goto L117
L120:
	;
	v489 = *(*int32)(unsafe.Add(mBase, uint32(v263)+8))
	v490 = F_copyObjectImpl(m, v489)
	mBase = m.M
	v491 = m.ExcPending
	if v491 != 0 {
		goto L2
	} else {
		goto L121
	}
L121:
	;
	v492 = *(*int32)(unsafe.Add(mBase, uint32(v89)+40))
	v493 = *(*int32)(unsafe.Add(mBase, uint32(v492)+4))
	v494 = F_lappend(m, v493, v490)
	mBase = m.M
	v495 = m.ExcPending
	if v495 != 0 {
		goto L2
	} else {
		goto L122
	}
L122:
	;
	v496 = *(*int32)(unsafe.Add(mBase, uint32(v89)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v496)+4)) = v494
	v498 = *(*int32)(unsafe.Add(mBase, uint32(v490)+4))
	F_cost_qual_eval_node(m, v230, v498, l0)
	mBase = m.M
	v500 = m.ExcPending
	if v500 != 0 {
		goto L2
	} else {
		goto L123
	}
L123:
	;
	v501 = *(*int32)(unsafe.Add(mBase, uint32(v89)+40))
	v502 = *(*float64)(unsafe.Add(mBase, uint32(v230)))
	v503 = *(*float64)(unsafe.Add(mBase, uint32(v501)+16))
	*(*float64)(unsafe.Add(mBase, uint32(v501)+16)) = base.F64_add(v502, v503)
	v506 = *(*int32)(unsafe.Add(mBase, uint32(v89)+40))
	v507 = *(*float64)(unsafe.Add(mBase, uint32(v230)+8))
	v508 = *(*float64)(unsafe.Add(mBase, uint32(v506)+24))
	*(*float64)(unsafe.Add(mBase, uint32(v506)+24)) = base.F64_add(v507, v508)
	v511 = int64(*(*int32)(unsafe.Add(mBase, uint32(v263)+24)))
	v515 = v257 + v511
	goto L75
L124:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v89)+68)) = v518
	v523 = v515
	goto L74
L125:
	;
	goto L57
L126:
	;
	v551 = *(*int32)(unsafe.Add(mBase, uint32(v89)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v551)+32)) = base.I32_wrap_i64(v549)
	m.G0 = v230 + int32(16)
	v556 = *(*int32)(unsafe.Add(mBase, uint32(v89)+68))
	v557 = *(*int32)(unsafe.Add(mBase, uint32(v89)+8))
	v558 = F_bms_del_members(m, v556, v557)
	mBase = m.M
	v559 = m.ExcPending
	if v559 != 0 {
		goto L2
	} else {
		goto L130
	}
L127:
	;
	v549 = v546
	goto L129
L128:
	;
	v549 = v544
	goto L129
L129:
	;
	goto L126
L130:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v89)+68)) = v558
	v561 = F_build_joinrel_restrictlist(m, l0, v89, v47, v48, v73)
	mBase = m.M
	v562 = m.ExcPending
	if v562 != 0 {
		goto L2
	} else {
		goto L131
	}
L131:
	;
	if v76 != 0 {
		goto L132
	} else {
		goto L133
	}
L132:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v76))) = v561
	goto L134
L133:
	;
	goto L134
L134:
	;
	v565 = v89 + int32(8)
	v566 = *(*int32)(unsafe.Add(mBase, uint32(v47)+228))
	if v566 == int32(0) {
		goto L136
	} else {
		goto L137
	}
L135:
	;
	v681 = *(*int32)(unsafe.Add(mBase, uint32(v48)+228))
	if v681 == int32(0) {
		v778 = v665
		goto L161
	} else {
		goto L162
	}
L136:
	;
	v665 = int32(0)
	goto L135
L137:
	;
	goto L138
L138:
	;
	v570 = int32(0)
	v571 = *(*int32)(unsafe.Add(mBase, uint32(v566)+4))
	if v571 <= v570 {
		v665 = v570
		goto L135
	} else {
		goto L139
	}
L139:
	;
	v577 = v570
	v578 = int32(0)
	goto L140
L140:
	;
	v593 = *(*int32)(unsafe.Add(mBase, uint32(v566)+12))
	v597 = *(*int32)(unsafe.Add(mBase, uint32(v593+v578<<(uint(int32(2))%32))))
	v598 = *(*int32)(unsafe.Add(mBase, uint32(v597)+32))
	v599 = *(*int32)(unsafe.Add(mBase, uint32(v565)))
	v600 = int32(0)
	if v598 == v600 {
		goto L143
	} else {
		goto L144
	}
L141:
	;
	v665 = v658
	goto L135
L142:
	;
	if v653 == int32(0) {
		goto L156
	} else {
		goto L157
	}
L143:
	;
	v653 = int32(1)
	goto L142
L144:
	;
	goto L145
L145:
	;
	if v599 == int32(0) {
		v646 = v600
		goto L146
	} else {
		goto L147
	}
L146:
	;
	v653 = v646
	goto L142
L147:
	;
	v609 = *(*int32)(unsafe.Add(mBase, uint32(v598)+4))
	v610 = *(*int32)(unsafe.Add(mBase, uint32(v599)+4))
	if v610 < v609 {
		v646 = v600
		goto L146
	} else {
		goto L148
	}
L148:
	;
	v612 = int32(1)
	if v609 <= v612 {
		goto L149
	} else {
		goto L150
	}
L149:
	;
	v615 = v612
	goto L151
L150:
	;
	v615 = v609
	goto L151
L151:
	;
	v616 = int32(8)
	v621 = int32(0)
	goto L152
L152:
	;
	v628 = v621 << (uint(int32(2)) % 32)
	v630 = *(*int32)(unsafe.Add(mBase, uint32(v598+v616+v628)))
	v632 = *(*int32)(unsafe.Add(mBase, uint32(v599+v616+v628)))
	v635 = v630 & (v632 ^ int32(-1))
	v637 = base.B2i32(v635 == int32(0))
	if v635 != 0 {
		v646 = v637
		goto L146
	} else {
		goto L154
	}
L153:
	;
	v646 = v637
	goto L146
L154:
	;
	v639 = v621 + int32(1)
	if v639 != v615 {
		v621 = v639
		goto L152
	} else {
		goto L155
	}
L155:
	;
	goto L153
L156:
	;
	v656 = F_list_append_unique_ptr(m, v577, v597)
	mBase = m.M
	v657 = m.ExcPending
	if v657 != 0 {
		goto L2
	} else {
		goto L159
	}
L157:
	;
	v658 = v577
	goto L158
L158:
	;
	v660 = v578 + int32(1)
	v661 = *(*int32)(unsafe.Add(mBase, uint32(v566)+4))
	if v660 < v661 {
		v577 = v658
		v578 = v660
		goto L140
	} else {
		goto L160
	}
L159:
	;
	v658 = v656
	goto L158
L160:
	;
	goto L141
L161:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v89)+228)) = v778
	v795 = int32(0)
	v796 = *(*int32)(unsafe.Add(mBase, uint32(v89)+8))
	if v796 == v795 {
		goto L188
	} else {
		goto L189
	}
L162:
	;
	v684 = *(*int32)(unsafe.Add(mBase, uint32(v681)+4))
	if v684 <= int32(0) {
		v778 = v665
		goto L161
	} else {
		goto L163
	}
L163:
	;
	v690 = v665
	v691 = int32(0)
	goto L164
L164:
	;
	v706 = *(*int32)(unsafe.Add(mBase, uint32(v681)+12))
	v710 = *(*int32)(unsafe.Add(mBase, uint32(v706+v691<<(uint(int32(2))%32))))
	v711 = *(*int32)(unsafe.Add(mBase, uint32(v710)+32))
	v712 = *(*int32)(unsafe.Add(mBase, uint32(v565)))
	v713 = int32(0)
	if v711 == v713 {
		goto L167
	} else {
		goto L168
	}
L165:
	;
	v778 = v771
	goto L161
L166:
	;
	if v766 == int32(0) {
		goto L180
	} else {
		goto L181
	}
L167:
	;
	v766 = int32(1)
	goto L166
L168:
	;
	goto L169
L169:
	;
	if v712 == int32(0) {
		v759 = v713
		goto L170
	} else {
		goto L171
	}
L170:
	;
	v766 = v759
	goto L166
L171:
	;
	v722 = *(*int32)(unsafe.Add(mBase, uint32(v711)+4))
	v723 = *(*int32)(unsafe.Add(mBase, uint32(v712)+4))
	if v723 < v722 {
		v759 = v713
		goto L170
	} else {
		goto L172
	}
L172:
	;
	v725 = int32(1)
	if v722 <= v725 {
		goto L173
	} else {
		goto L174
	}
L173:
	;
	v728 = v725
	goto L175
L174:
	;
	v728 = v722
	goto L175
L175:
	;
	v729 = int32(8)
	v734 = int32(0)
	goto L176
L176:
	;
	v741 = v734 << (uint(int32(2)) % 32)
	v743 = *(*int32)(unsafe.Add(mBase, uint32(v711+v729+v741)))
	v745 = *(*int32)(unsafe.Add(mBase, uint32(v712+v729+v741)))
	v748 = v743 & (v745 ^ int32(-1))
	v750 = base.B2i32(v748 == int32(0))
	if v748 != 0 {
		v759 = v750
		goto L170
	} else {
		goto L178
	}
L177:
	;
	v759 = v750
	goto L170
L178:
	;
	v752 = v734 + int32(1)
	if v752 != v728 {
		v734 = v752
		goto L176
	} else {
		goto L179
	}
L179:
	;
	goto L177
L180:
	;
	v769 = F_list_append_unique_ptr(m, v690, v710)
	mBase = m.M
	v770 = m.ExcPending
	if v770 != 0 {
		goto L2
	} else {
		goto L183
	}
L181:
	;
	v771 = v690
	goto L182
L182:
	;
	v773 = v691 + int32(1)
	v774 = *(*int32)(unsafe.Add(mBase, uint32(v681)+4))
	if v773 < v774 {
		v690 = v771
		v691 = v773
		goto L164
	} else {
		goto L184
	}
L183:
	;
	v771 = v769
	goto L182
L184:
	;
	goto L165
L185:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v89)+232)) = uint8(v1205)
	F_set_joinrel_size_estimates(m, l0, v89, v47, v48, v73, v561)
	mBase = m.M
	v1208 = m.ExcPending
	if v1208 != 0 {
		goto L2
	} else {
		goto L264
	}
L186:
	;
	if int32(0) < v853 {
		goto L197
	} else {
		goto L198
	}
L187:
	;
	v853 = base.I32_ctz(v839) | v840<<(uint(int32(5))%32)
	goto L186
L188:
	;
	v853 = int32(-2)
	goto L186
L189:
	;
	v804 = int32(0)
	v807 = *(*int32)(unsafe.Add(mBase, uint32(v796)+4))
	if v807 <= v804 {
		goto L188
	} else {
		goto L190
	}
L190:
	;
	v810 = v796 + int32(8)
	v814 = *(*int32)(unsafe.Add(mBase, uint32(v810)))
	v817 = v814 & int32(-1)
	if v817 != 0 {
		v839 = v817
		v840 = v804
		goto L187
	} else {
		goto L191
	}
L191:
	;
	v818 = int32(1)
	if v818 == v807 {
		goto L188
	} else {
		goto L192
	}
L192:
	;
	v822 = v818
	goto L193
L193:
	;
	v829 = *(*int32)(unsafe.Add(mBase, uint32(v810+v822<<(uint(int32(2))%32))))
	if v829 != 0 {
		v839 = v829
		v840 = v822
		goto L187
	} else {
		goto L195
	}
L194:
	;
	goto L188
L195:
	;
	v831 = v822 + int32(1)
	if v831 != v807 {
		v822 = v831
		goto L193
	} else {
		goto L196
	}
L196:
	;
	goto L194
L197:
	;
	v858 = v795
	v859 = v853
	goto L200
L198:
	;
	v948 = v795
	goto L199
L199:
	;
	if v948 == int32(0) {
		goto L220
	} else {
		goto L221
	}
L200:
	;
	v874 = *(*int32)(unsafe.Add(mBase, uint32(l0)+340))
	if v859 == v874 {
		v886 = v858
		goto L202
	} else {
		goto L203
	}
L201:
	;
	v948 = v886
	goto L199
L202:
	;
	if v796 == int32(0) {
		goto L208
	} else {
		goto L209
	}
L203:
	;
	v876 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v880 = *(*int32)(unsafe.Add(mBase, uint32(v876+v859<<(uint(int32(2))%32))))
	if v880 == int32(0) {
		v886 = v858
		goto L202
	} else {
		goto L204
	}
L204:
	;
	v883 = *(*int32)(unsafe.Add(mBase, uint32(v880)+144))
	v884 = F_bms_add_members(m, v858, v883)
	mBase = m.M
	v885 = m.ExcPending
	if v885 != 0 {
		goto L2
	} else {
		goto L205
	}
L205:
	;
	v886 = v884
	goto L202
L206:
	;
	if int32(0) < v943 {
		v858 = v886
		v859 = v943
		goto L200
	} else {
		goto L217
	}
L207:
	;
	v943 = base.I32_ctz(v929) | v930<<(uint(int32(5))%32)
	goto L206
L208:
	;
	v943 = int32(-2)
	goto L206
L209:
	;
	v894 = v859 + int32(1)
	v896 = int32(base.Ui32(v894) >> (uint(int32(5)) % 32))
	v897 = *(*int32)(unsafe.Add(mBase, uint32(v796)+4))
	if v897 <= v896 {
		goto L208
	} else {
		goto L210
	}
L210:
	;
	v900 = v796 + int32(8)
	v904 = *(*int32)(unsafe.Add(mBase, uint32(v900+v896<<(uint(int32(2))%32))))
	v907 = v904 & (int32(-1) << (uint(v894) % 32))
	if v907 != 0 {
		v929 = v907
		v930 = v896
		goto L207
	} else {
		goto L211
	}
L211:
	;
	v909 = v896 + int32(1)
	if v909 == v897 {
		goto L208
	} else {
		goto L212
	}
L212:
	;
	v912 = v909
	goto L213
L213:
	;
	v919 = *(*int32)(unsafe.Add(mBase, uint32(v900+v912<<(uint(int32(2))%32))))
	if v919 != 0 {
		v929 = v919
		v930 = v912
		goto L207
	} else {
		goto L215
	}
L214:
	;
	goto L208
L215:
	;
	v921 = v912 + int32(1)
	if v921 != v897 {
		v912 = v921
		goto L213
	} else {
		goto L216
	}
L216:
	;
	goto L214
L217:
	;
	goto L201
L218:
	;
	if int32(0) <= v1020 {
		goto L229
	} else {
		goto L230
	}
L219:
	;
	v1020 = base.I32_ctz(v1006) | v1007<<(uint(int32(5))%32)
	goto L218
L220:
	;
	v1020 = int32(-2)
	goto L218
L221:
	;
	v971 = int32(0)
	v974 = *(*int32)(unsafe.Add(mBase, uint32(v948)+4))
	if v974 <= v971 {
		goto L220
	} else {
		goto L222
	}
L222:
	;
	v977 = v948 + int32(8)
	v981 = *(*int32)(unsafe.Add(mBase, uint32(v977)))
	v984 = v981 & int32(-1)
	if v984 != 0 {
		v1006 = v984
		v1007 = v971
		goto L219
	} else {
		goto L223
	}
L223:
	;
	v985 = int32(1)
	if v985 == v974 {
		goto L220
	} else {
		goto L224
	}
L224:
	;
	v989 = v985
	goto L225
L225:
	;
	v996 = *(*int32)(unsafe.Add(mBase, uint32(v977+v989<<(uint(int32(2))%32))))
	if v996 != 0 {
		v1006 = v996
		v1007 = v989
		goto L219
	} else {
		goto L227
	}
L226:
	;
	goto L220
L227:
	;
	v998 = v989 + int32(1)
	if v998 != v974 {
		v989 = v998
		goto L225
	} else {
		goto L228
	}
L228:
	;
	goto L226
L229:
	;
	v1026 = v1020
	goto L232
L230:
	;
	goto L231
L231:
	;
	v1205 = int32(0)
	goto L185
L232:
	;
	v1041 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v1042 = *(*int32)(unsafe.Add(mBase, uint32(v1041)+12))
	v1046 = *(*int32)(unsafe.Add(mBase, uint32(v1042+v1026<<(uint(int32(2))%32))))
	v1047 = *(*int32)(unsafe.Add(mBase, uint32(v1046)+16))
	if v1047 == int32(0) {
		goto L234
	} else {
		goto L235
	}
L233:
	;
	goto L231
L234:
	;
	if v948 == int32(0) {
		goto L254
	} else {
		goto L255
	}
L235:
	;
	v1050 = *(*int32)(unsafe.Add(mBase, uint32(v1047)+4))
	if v1050 < int32(2) {
		goto L234
	} else {
		goto L236
	}
L236:
	;
	v1053 = *(*int32)(unsafe.Add(mBase, uint32(v1046)+36))
	v1054 = *(*int32)(unsafe.Add(mBase, uint32(v89)+8))
	v1055 = int32(0)
	if v1053 == v1055 {
		goto L238
	} else {
		goto L239
	}
L237:
	;
	if v1108 != 0 {
		goto L234
	} else {
		goto L251
	}
L238:
	;
	v1108 = int32(1)
	goto L237
L239:
	;
	goto L240
L240:
	;
	if v1054 == int32(0) {
		v1101 = v1055
		goto L241
	} else {
		goto L242
	}
L241:
	;
	v1108 = v1101
	goto L237
L242:
	;
	v1064 = *(*int32)(unsafe.Add(mBase, uint32(v1053)+4))
	v1065 = *(*int32)(unsafe.Add(mBase, uint32(v1054)+4))
	if v1065 < v1064 {
		v1101 = v1055
		goto L241
	} else {
		goto L243
	}
L243:
	;
	v1067 = int32(1)
	if v1064 <= v1067 {
		goto L244
	} else {
		goto L245
	}
L244:
	;
	v1070 = v1067
	goto L246
L245:
	;
	v1070 = v1064
	goto L246
L246:
	;
	v1071 = int32(8)
	v1076 = int32(0)
	goto L247
L247:
	;
	v1083 = v1076 << (uint(int32(2)) % 32)
	v1085 = *(*int32)(unsafe.Add(mBase, uint32(v1053+v1071+v1083)))
	v1087 = *(*int32)(unsafe.Add(mBase, uint32(v1054+v1071+v1083)))
	v1090 = v1085 & (v1087 ^ int32(-1))
	v1092 = base.B2i32(v1090 == int32(0))
	if v1090 != 0 {
		v1101 = v1092
		goto L241
	} else {
		goto L249
	}
L248:
	;
	v1101 = v1092
	goto L241
L249:
	;
	v1094 = v1076 + int32(1)
	if v1094 != v1070 {
		v1076 = v1094
		goto L247
	} else {
		goto L250
	}
L250:
	;
	goto L248
L251:
	;
	v1205 = int32(1)
	goto L185
L252:
	;
	if int32(0) <= v1165 {
		v1026 = v1165
		goto L232
	} else {
		goto L263
	}
L253:
	;
	v1165 = base.I32_ctz(v1151) | v1152<<(uint(int32(5))%32)
	goto L252
L254:
	;
	v1165 = int32(-2)
	goto L252
L255:
	;
	v1116 = v1026 + int32(1)
	v1118 = int32(base.Ui32(v1116) >> (uint(int32(5)) % 32))
	v1119 = *(*int32)(unsafe.Add(mBase, uint32(v948)+4))
	if v1119 <= v1118 {
		goto L254
	} else {
		goto L256
	}
L256:
	;
	v1122 = v948 + int32(8)
	v1126 = *(*int32)(unsafe.Add(mBase, uint32(v1122+v1118<<(uint(int32(2))%32))))
	v1129 = v1126 & (int32(-1) << (uint(v1116) % 32))
	if v1129 != 0 {
		v1151 = v1129
		v1152 = v1118
		goto L253
	} else {
		goto L257
	}
L257:
	;
	v1131 = v1118 + int32(1)
	if v1131 == v1119 {
		goto L254
	} else {
		goto L258
	}
L258:
	;
	v1134 = v1131
	goto L259
L259:
	;
	v1141 = *(*int32)(unsafe.Add(mBase, uint32(v1122+v1134<<(uint(int32(2))%32))))
	if v1141 != 0 {
		v1151 = v1141
		v1152 = v1134
		goto L253
	} else {
		goto L261
	}
L260:
	;
	goto L254
L261:
	;
	v1143 = v1134 + int32(1)
	if v1143 != v1119 {
		v1134 = v1143
		goto L259
	} else {
		goto L262
	}
L262:
	;
	goto L260
L263:
	;
	goto L233
L264:
	;
	v1209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+26)))
	if v1209 != int32(1) {
		goto L265
	} else {
		goto L266
	}
L265:
	;
	v1228 = *(*int32)(unsafe.Add(mBase, _c_F_make_join_rel[1]))
	if v1228 != 0 {
		goto L272
	} else {
		goto L273
	}
L266:
	;
	v1212 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+26)))
	if v1212 != int32(1) {
		goto L265
	} else {
		goto L267
	}
L267:
	;
	v1215 = F_is_parallel_safe(m, l0, v561)
	mBase = m.M
	v1216 = m.ExcPending
	if v1216 != 0 {
		goto L2
	} else {
		goto L268
	}
L268:
	;
	if v1215 == int32(0) {
		goto L265
	} else {
		goto L269
	}
L269:
	;
	v1219 = *(*int32)(unsafe.Add(mBase, uint32(v89)+40))
	v1220 = *(*int32)(unsafe.Add(mBase, uint32(v1219)+4))
	v1221 = F_is_parallel_safe(m, l0, v1220)
	mBase = m.M
	v1222 = m.ExcPending
	if v1222 != 0 {
		goto L2
	} else {
		goto L270
	}
L270:
	;
	if v1221 == int32(0) {
		goto L265
	} else {
		goto L271
	}
L271:
	;
	v1225 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v89)+26)) = uint8(v1225)
	goto L265
L272:
	;
	m.T0[v1228].(func(*base.Module, int32, int32, int32, int32, int32, int32))(m, l0, v89, v47, v48, v73, v561)
	mBase = m.M
	v1230 = m.ExcPending
	if v1230 != 0 {
		goto L2
	} else {
		goto L275
	}
L273:
	;
	goto L274
L274:
	;
	F_build_joinrel_partition_info(m, l0, v89, v47, v48, v73, v561)
	mBase = m.M
	v1232 = m.ExcPending
	if v1232 != 0 {
		goto L2
	} else {
		goto L276
	}
L275:
	;
	goto L274
L276:
	;
	v1233 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v1234 = F_lappend(m, v1233, v89)
	mBase = m.M
	v1235 = m.ExcPending
	if v1235 != 0 {
		goto L2
	} else {
		goto L277
	}
L277:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = v1234
	v1237 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	if v1237 != 0 {
		goto L278
	} else {
		goto L279
	}
L278:
	;
	v1241 = F_hash_search(m, v1237, v565, int32(1), v79+int32(15))
	mBase = m.M
	v1242 = m.ExcPending
	if v1242 != 0 {
		goto L2
	} else {
		goto L281
	}
L279:
	;
	goto L280
L280:
	;
	v1244 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	if v1244 == int32(0) {
		v1261 = v89
		goto L19
	} else {
		goto L282
	}
L281:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1241)+4)) = v89
	goto L280
L282:
	;
	v1247 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v1251 = *(*int32)(unsafe.Add(mBase, uint32(v1244+v1247<<(uint(int32(2))%32))))
	v1252 = F_lappend(m, v1251, v89)
	mBase = m.M
	v1253 = m.ExcPending
	if v1253 != 0 {
		goto L2
	} else {
		goto L283
	}
L283:
	;
	v1254 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	v1255 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v1254+v1255<<(uint(int32(2))%32)))) = v1252
	v1261 = v89
	goto L19
L284:
	;
	F_bms_free(m, v44)
	mBase = m.M
	v1359 = m.ExcPending
	if v1359 != 0 {
		goto L2
	} else {
		goto L296
	}
L285:
	;
	v1334 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
	F_make_grouped_join_rel(m, l0, v47, v48, v1261, v73, v1334)
	mBase = m.M
	v1336 = m.ExcPending
	if v1336 != 0 {
		goto L2
	} else {
		goto L294
	}
L286:
	;
	v1284 = *(*int32)(unsafe.Add(mBase, uint32(v1281)+12))
	v1287 = v1284
	goto L287
L287:
	;
	v1303 = *(*int32)(unsafe.Add(mBase, uint32(v1287)))
	v1304 = *(*int32)(unsafe.Add(mBase, uint32(v1303)))
	if base.Ui32(int32(2)) <= base.Ui32(v1304-int32(303)) {
		goto L289
	} else {
		goto L290
	}
L288:
	;
	goto L285
L289:
	;
	if v1304 != int32(293) {
		goto L285
	} else {
		goto L292
	}
L290:
	;
	v1287 = v1303 + int32(72)
	goto L287
L291:
	;
	goto L288
L292:
	;
	v1311 = *(*int32)(unsafe.Add(mBase, uint32(v1303)+72))
	if v1311 == int32(0) {
		goto L284
	} else {
		goto L293
	}
L293:
	;
	goto L291
L294:
	;
	v1337 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
	F_populate_joinrel_with_paths(m, l0, v47, v48, v1261, v73, v1337)
	mBase = m.M
	v1339 = m.ExcPending
	if v1339 != 0 {
		goto L2
	} else {
		goto L295
	}
L295:
	;
	goto L284
L296:
	;
	v1363 = v1261
	goto L1
}
