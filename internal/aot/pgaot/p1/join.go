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
	var v137 int64
	_ = v137
	var v157 int64
	_ = v157
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v228 int64
	_ = v228
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v252 int64
	_ = v252
	var v254 int32
	_ = v254
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v281 int32
	_ = v281
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v306 int32
	_ = v306
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v322 int32
	_ = v322
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v328 int32
	_ = v328
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v337 int32
	_ = v337
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v348 int32
	_ = v348
	var v351 int32
	_ = v351
	var v353 int32
	_ = v353
	var v355 int32
	_ = v355
	var v359 int32
	_ = v359
	var v369 int32
	_ = v369
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v386 int32
	_ = v386
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v395 int32
	_ = v395
	var v402 int32
	_ = v402
	var v404 int32
	_ = v404
	var v406 int32
	_ = v406
	var v409 int32
	_ = v409
	var v411 int32
	_ = v411
	var v413 int32
	_ = v413
	var v420 int32
	_ = v420
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v442 int32
	_ = v442
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v451 int32
	_ = v451
	var v458 int32
	_ = v458
	var v460 int32
	_ = v460
	var v462 int32
	_ = v462
	var v465 int32
	_ = v465
	var v467 int32
	_ = v467
	var v469 int32
	_ = v469
	var v476 int32
	_ = v476
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
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v493 int32
	_ = v493
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v497 float64
	_ = v497
	var v498 float64
	_ = v498
	var v501 int32
	_ = v501
	var v502 float64
	_ = v502
	var v503 float64
	_ = v503
	var v506 int64
	_ = v506
	var v510 int64
	_ = v510
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v518 int64
	_ = v518
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v539 int64
	_ = v539
	var v541 int64
	_ = v541
	var v544 int64
	_ = v544
	var v546 int32
	_ = v546
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
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
	var v566 int32
	_ = v566
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v588 int32
	_ = v588
	var v592 int32
	_ = v592
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v604 int32
	_ = v604
	var v605 int32
	_ = v605
	var v607 int32
	_ = v607
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v616 int32
	_ = v616
	var v623 int32
	_ = v623
	var v625 int32
	_ = v625
	var v627 int32
	_ = v627
	var v630 int32
	_ = v630
	var v632 int32
	_ = v632
	var v634 int32
	_ = v634
	var v641 int32
	_ = v641
	var v648 int32
	_ = v648
	var v651 int32
	_ = v651
	var v652 int32
	_ = v652
	var v653 int32
	_ = v653
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v660 int32
	_ = v660
	var v676 int32
	_ = v676
	var v679 int32
	_ = v679
	var v685 int32
	_ = v685
	var v686 int32
	_ = v686
	var v701 int32
	_ = v701
	var v705 int32
	_ = v705
	var v706 int32
	_ = v706
	var v707 int32
	_ = v707
	var v708 int32
	_ = v708
	var v717 int32
	_ = v717
	var v718 int32
	_ = v718
	var v720 int32
	_ = v720
	var v723 int32
	_ = v723
	var v724 int32
	_ = v724
	var v729 int32
	_ = v729
	var v736 int32
	_ = v736
	var v738 int32
	_ = v738
	var v740 int32
	_ = v740
	var v743 int32
	_ = v743
	var v745 int32
	_ = v745
	var v747 int32
	_ = v747
	var v754 int32
	_ = v754
	var v761 int32
	_ = v761
	var v764 int32
	_ = v764
	var v765 int32
	_ = v765
	var v766 int32
	_ = v766
	var v768 int32
	_ = v768
	var v769 int32
	_ = v769
	var v773 int32
	_ = v773
	var v790 int32
	_ = v790
	var v791 int32
	_ = v791
	var v801 int32
	_ = v801
	var v802 int32
	_ = v802
	var v805 int32
	_ = v805
	var v809 int32
	_ = v809
	var v812 int32
	_ = v812
	var v814 int32
	_ = v814
	var v817 int32
	_ = v817
	var v824 int32
	_ = v824
	var v826 int32
	_ = v826
	var v834 int32
	_ = v834
	var v835 int32
	_ = v835
	var v848 int32
	_ = v848
	var v853 int32
	_ = v853
	var v854 int32
	_ = v854
	var v869 int32
	_ = v869
	var v871 int32
	_ = v871
	var v875 int32
	_ = v875
	var v878 int32
	_ = v878
	var v879 int32
	_ = v879
	var v880 int32
	_ = v880
	var v881 int32
	_ = v881
	var v889 int32
	_ = v889
	var v891 int32
	_ = v891
	var v892 int32
	_ = v892
	var v895 int32
	_ = v895
	var v899 int32
	_ = v899
	var v902 int32
	_ = v902
	var v904 int32
	_ = v904
	var v907 int32
	_ = v907
	var v914 int32
	_ = v914
	var v916 int32
	_ = v916
	var v924 int32
	_ = v924
	var v925 int32
	_ = v925
	var v938 int32
	_ = v938
	var v943 int32
	_ = v943
	var v968 int32
	_ = v968
	var v969 int32
	_ = v969
	var v972 int32
	_ = v972
	var v976 int32
	_ = v976
	var v979 int32
	_ = v979
	var v981 int32
	_ = v981
	var v984 int32
	_ = v984
	var v991 int32
	_ = v991
	var v993 int32
	_ = v993
	var v1001 int32
	_ = v1001
	var v1002 int32
	_ = v1002
	var v1015 int32
	_ = v1015
	var v1021 int32
	_ = v1021
	var v1036 int32
	_ = v1036
	var v1037 int32
	_ = v1037
	var v1041 int32
	_ = v1041
	var v1042 int32
	_ = v1042
	var v1045 int32
	_ = v1045
	var v1048 int32
	_ = v1048
	var v1049 int32
	_ = v1049
	var v1050 int32
	_ = v1050
	var v1059 int32
	_ = v1059
	var v1060 int32
	_ = v1060
	var v1062 int32
	_ = v1062
	var v1065 int32
	_ = v1065
	var v1066 int32
	_ = v1066
	var v1071 int32
	_ = v1071
	var v1078 int32
	_ = v1078
	var v1080 int32
	_ = v1080
	var v1082 int32
	_ = v1082
	var v1085 int32
	_ = v1085
	var v1087 int32
	_ = v1087
	var v1089 int32
	_ = v1089
	var v1096 int32
	_ = v1096
	var v1103 int32
	_ = v1103
	var v1111 int32
	_ = v1111
	var v1113 int32
	_ = v1113
	var v1114 int32
	_ = v1114
	var v1117 int32
	_ = v1117
	var v1121 int32
	_ = v1121
	var v1124 int32
	_ = v1124
	var v1126 int32
	_ = v1126
	var v1129 int32
	_ = v1129
	var v1136 int32
	_ = v1136
	var v1138 int32
	_ = v1138
	var v1146 int32
	_ = v1146
	var v1147 int32
	_ = v1147
	var v1160 int32
	_ = v1160
	var v1200 int32
	_ = v1200
	var v1203 int32
	_ = v1203
	var v1205 int32
	_ = v1205
	var v1206 int32
	_ = v1206
	var v1209 int32
	_ = v1209
	var v1212 int32
	_ = v1212
	var v1213 int32
	_ = v1213
	var v1216 int32
	_ = v1216
	var v1217 int32
	_ = v1217
	var v1218 int32
	_ = v1218
	var v1219 int32
	_ = v1219
	var v1222 int32
	_ = v1222
	var v1224 int32
	_ = v1224
	var v1225 int32
	_ = v1225
	var v1226 int32
	_ = v1226
	var v1228 int32
	_ = v1228
	var v1232 int32
	_ = v1232
	var v1233 int32
	_ = v1233
	var v1235 int32
	_ = v1235
	var v1238 int32
	_ = v1238
	var v1242 int32
	_ = v1242
	var v1243 int32
	_ = v1243
	var v1244 int32
	_ = v1244
	var v1245 int32
	_ = v1245
	var v1246 int32
	_ = v1246
	var v1252 int32
	_ = v1252
	var v1272 int32
	_ = v1272
	var v1275 int32
	_ = v1275
	var v1278 int32
	_ = v1278
	var v1294 int32
	_ = v1294
	var v1295 int32
	_ = v1295
	var v1302 int32
	_ = v1302
	var v1325 int32
	_ = v1325
	var v1327 int32
	_ = v1327
	var v1347 int32
	_ = v1347
	var v1351 int32
	_ = v1351
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
	return v1351
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
	v1351 = v4
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
	*(*int32)(unsafe.Add(mBase, uint32(v21)+12)) = int32(320)
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
	v1272 = *(*int32)(unsafe.Add(mBase, uint32(v1252)+32))
	if v1272 == int32(0) {
		goto L281
	} else {
		goto L282
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
		v1252 = v81
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
	v1252 = v81
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
	base.MemoryFill(m, v89+int32(80), v129, int32(68))
	v137 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v89)+152)) = v137
	*(*int32)(unsafe.Add(mBase, uint32(v89)+148)) = int32(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v89)+157)) = v137
	*(*int64)(unsafe.Add(mBase, uint32(v89)+168)) = v137
	*(*int64)(unsafe.Add(mBase, uint32(v89)+176)) = v137
	*(*int32)(unsafe.Add(mBase, uint32(v89)+184)) = v129
	*(*int64)(unsafe.Add(mBase, uint32(v89)+192)) = v137
	*(*int64)(unsafe.Add(mBase, uint32(v89)+200)) = v137
	*(*int64)(unsafe.Add(mBase, uint32(v89)+220)) = v137
	*(*uint16)(unsafe.Add(mBase, uint32(v89)+216)) = uint16(v129)
	v157 = int64(4294967295)
	*(*int64)(unsafe.Add(mBase, uint32(v89)+208)) = v157
	*(*int64)(unsafe.Add(mBase, uint32(v89)+228)) = v137
	*(*uint8)(unsafe.Add(mBase, uint32(v89)+244)) = uint8(v129)
	*(*int64)(unsafe.Add(mBase, uint32(v89)+236)) = v157
	*(*int64)(unsafe.Add(mBase, uint32(v89)+248)) = v137
	*(*int64)(unsafe.Add(mBase, uint32(v89)+256)) = v137
	*(*int64)(unsafe.Add(mBase, uint32(v89)+264)) = v137
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v47)+156))
	if v171 == v129 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v73)+20))
	F_build_joinrel_tlist(m, l0, v89, v47, v73, v74, base.B2i32(v212 == int32(2)))
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L2
	} else {
		goto L51
	}
L33:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v48)+156))
	if v174 != v171 {
		goto L32
	} else {
		goto L34
	}
L34:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v48)+160))
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v47)+160))
	if v176 == v177 {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v89)+164)) = uint8(v206)
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v47)+168))
	*(*int32)(unsafe.Add(mBase, uint32(v89)+168)) = v208
	goto L32
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v89)+156)) = v171
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v47)+160))
	*(*int32)(unsafe.Add(mBase, uint32(v89)+160)) = v180
	v182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+164)))
	if v182 != 0 {
		goto L39
	} else {
		goto L40
	}
L37:
	;
	goto L38
L38:
	;
	if v176 != 0 {
		goto L43
	} else {
		goto L44
	}
L39:
	;
	v185 = int32(1)
	goto L41
L40:
	;
	v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+164)))
	v185 = v184
	goto L41
L41:
	;
	v206 = v185 & int32(1)
	goto L35
L42:
	;
	v206 = int32(1)
	goto L35
L43:
	;
	v196 = v177
	goto L45
L44:
	;
	v189 = *(*int32)(unsafe.Add(mBase, _c_F_make_join_rel[0]))
	if v177 == v189 {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	if v196 != 0 {
		goto L32
	} else {
		goto L49
	}
L46:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v47)+156))
	*(*int32)(unsafe.Add(mBase, uint32(v89)+156)) = v191
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v47)+160))
	*(*int32)(unsafe.Add(mBase, uint32(v89)+160)) = v193
	goto L42
L47:
	;
	goto L48
L48:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v47)+160))
	v196 = v195
	goto L45
L49:
	;
	v198 = *(*int32)(unsafe.Add(mBase, _c_F_make_join_rel[0]))
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v48)+160))
	if v198 != v199 {
		goto L32
	} else {
		goto L50
	}
L50:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v47)+156))
	*(*int32)(unsafe.Add(mBase, uint32(v89)+156)) = v201
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v48)+160))
	*(*int32)(unsafe.Add(mBase, uint32(v89)+160)) = v203
	goto L42
L51:
	;
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v73)+20))
	F_build_joinrel_tlist(m, l0, v89, v48, v73, v74, base.B2i32(v217 != int32(0)))
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L2
	} else {
		goto L52
	}
L52:
	;
	v222 = int32(0)
	v223 = m.G0
	v225 = v223 - int32(16)
	m.G0 = v225
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v89)+28))
	v228 = int64(*(*int32)(unsafe.Add(mBase, uint32(v227)+32)))
	v229 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	if v229 == v222 {
		v539 = v228
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v541 = int64(1073741823)
	if v541 <= v539 {
		goto L127
	} else {
		goto L128
	}
L54:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v229)+4))
	if v232 <= int32(0) {
		v539 = v228
		goto L53
	} else {
		goto L55
	}
L55:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v89)+8))
	v238 = v222
	v252 = v228
	goto L56
L56:
	;
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v229)+12))
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v254+v238<<(uint(int32(2))%32))))
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v258)+12))
	v260 = int32(0)
	if v259 == v260 {
		goto L59
	} else {
		goto L60
	}
L57:
	;
	v539 = v518
	goto L53
L58:
	;
	if v313 != 0 {
		goto L72
	} else {
		goto L73
	}
L59:
	;
	v313 = int32(1)
	goto L58
L60:
	;
	goto L61
L61:
	;
	if v235 == int32(0) {
		v306 = v260
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v313 = v306
	goto L58
L63:
	;
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v259)+4))
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v235)+4))
	if v270 < v269 {
		v306 = v260
		goto L62
	} else {
		goto L64
	}
L64:
	;
	v272 = int32(1)
	if v269 <= v272 {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v275 = v272
	goto L67
L66:
	;
	v275 = v269
	goto L67
L67:
	;
	v276 = int32(8)
	v281 = int32(0)
	goto L68
L68:
	;
	v288 = v281 << (uint(int32(2)) % 32)
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v259+v276+v288)))
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v235+v276+v288)))
	v295 = v290 & (v292 ^ int32(-1))
	v297 = base.B2i32(v295 == int32(0))
	if v295 != 0 {
		v306 = v297
		goto L62
	} else {
		goto L70
	}
L69:
	;
	v306 = v297
	goto L62
L70:
	;
	v299 = v281 + int32(1)
	if v299 != v275 {
		v281 = v299
		goto L68
	} else {
		goto L71
	}
L71:
	;
	goto L69
L72:
	;
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v258)+20))
	if v314 == int32(0) {
		goto L77
	} else {
		goto L78
	}
L73:
	;
	v518 = v252
	goto L74
L74:
	;
	v520 = v238 + int32(1)
	v521 = *(*int32)(unsafe.Add(mBase, uint32(v229)+4))
	if v520 < v521 {
		v238 = v520
		v252 = v518
		goto L56
	} else {
		goto L125
	}
L75:
	;
	v511 = *(*int32)(unsafe.Add(mBase, uint32(v89)+60))
	v512 = *(*int32)(unsafe.Add(mBase, uint32(v258)+16))
	v513 = F_bms_add_members(m, v511, v512)
	mBase = m.M
	v514 = m.ExcPending
	if v514 != 0 {
		goto L2
	} else {
		goto L124
	}
L76:
	;
	if v369 == int32(0) {
		v510 = v252
		goto L75
	} else {
		goto L90
	}
L77:
	;
	v369 = int32(0)
	goto L76
L78:
	;
	goto L79
L79:
	;
	v322 = int32(1)
	if v235 == int32(0) {
		v359 = v322
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v369 = v359
	goto L76
L81:
	;
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v314)+4))
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v235)+4))
	if v326 < v325 {
		v359 = v322
		goto L80
	} else {
		goto L82
	}
L82:
	;
	v328 = int32(1)
	if v325 <= v328 {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v331 = v328
	goto L85
L84:
	;
	v331 = v325
	goto L85
L85:
	;
	v332 = int32(8)
	v337 = int32(0)
	goto L86
L86:
	;
	v344 = v337 << (uint(int32(2)) % 32)
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v314+v332+v344)))
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v235+v332+v344)))
	v351 = v346 & (v348 ^ int32(-1))
	v353 = base.B2i32(v351 != int32(0))
	if v351 != 0 {
		v359 = v353
		goto L80
	} else {
		goto L88
	}
L87:
	;
	v359 = v353
	goto L80
L88:
	;
	v355 = v337 + int32(1)
	if v355 != v331 {
		v337 = v355
		goto L86
	} else {
		goto L89
	}
L89:
	;
	goto L87
L90:
	;
	v372 = *(*int32)(unsafe.Add(mBase, uint32(v258)+12))
	v373 = *(*int32)(unsafe.Add(mBase, uint32(v47)+8))
	v374 = int32(0)
	if v372 == v374 {
		goto L92
	} else {
		goto L93
	}
L91:
	;
	if v427 != 0 {
		v510 = v252
		goto L75
	} else {
		goto L105
	}
L92:
	;
	v427 = int32(1)
	goto L91
L93:
	;
	goto L94
L94:
	;
	if v373 == int32(0) {
		v420 = v374
		goto L95
	} else {
		goto L96
	}
L95:
	;
	v427 = v420
	goto L91
L96:
	;
	v383 = *(*int32)(unsafe.Add(mBase, uint32(v372)+4))
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v373)+4))
	if v384 < v383 {
		v420 = v374
		goto L95
	} else {
		goto L97
	}
L97:
	;
	v386 = int32(1)
	if v383 <= v386 {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v389 = v386
	goto L100
L99:
	;
	v389 = v383
	goto L100
L100:
	;
	v390 = int32(8)
	v395 = int32(0)
	goto L101
L101:
	;
	v402 = v395 << (uint(int32(2)) % 32)
	v404 = *(*int32)(unsafe.Add(mBase, uint32(v372+v390+v402)))
	v406 = *(*int32)(unsafe.Add(mBase, uint32(v373+v390+v402)))
	v409 = v404 & (v406 ^ int32(-1))
	v411 = base.B2i32(v409 == int32(0))
	if v409 != 0 {
		v420 = v411
		goto L95
	} else {
		goto L103
	}
L102:
	;
	v420 = v411
	goto L95
L103:
	;
	v413 = v395 + int32(1)
	if v413 != v389 {
		v395 = v413
		goto L101
	} else {
		goto L104
	}
L104:
	;
	goto L102
L105:
	;
	v428 = *(*int32)(unsafe.Add(mBase, uint32(v258)+12))
	v429 = *(*int32)(unsafe.Add(mBase, uint32(v48)+8))
	v430 = int32(0)
	if v428 == v430 {
		goto L107
	} else {
		goto L108
	}
L106:
	;
	if v483 != 0 {
		v510 = v252
		goto L75
	} else {
		goto L120
	}
L107:
	;
	v483 = int32(1)
	goto L106
L108:
	;
	goto L109
L109:
	;
	if v429 == int32(0) {
		v476 = v430
		goto L110
	} else {
		goto L111
	}
L110:
	;
	v483 = v476
	goto L106
L111:
	;
	v439 = *(*int32)(unsafe.Add(mBase, uint32(v428)+4))
	v440 = *(*int32)(unsafe.Add(mBase, uint32(v429)+4))
	if v440 < v439 {
		v476 = v430
		goto L110
	} else {
		goto L112
	}
L112:
	;
	v442 = int32(1)
	if v439 <= v442 {
		goto L113
	} else {
		goto L114
	}
L113:
	;
	v445 = v442
	goto L115
L114:
	;
	v445 = v439
	goto L115
L115:
	;
	v446 = int32(8)
	v451 = int32(0)
	goto L116
L116:
	;
	v458 = v451 << (uint(int32(2)) % 32)
	v460 = *(*int32)(unsafe.Add(mBase, uint32(v428+v446+v458)))
	v462 = *(*int32)(unsafe.Add(mBase, uint32(v429+v446+v458)))
	v465 = v460 & (v462 ^ int32(-1))
	v467 = base.B2i32(v465 == int32(0))
	if v465 != 0 {
		v476 = v467
		goto L110
	} else {
		goto L118
	}
L117:
	;
	v476 = v467
	goto L110
L118:
	;
	v469 = v451 + int32(1)
	if v469 != v445 {
		v451 = v469
		goto L116
	} else {
		goto L119
	}
L119:
	;
	goto L117
L120:
	;
	v484 = *(*int32)(unsafe.Add(mBase, uint32(v258)+8))
	v485 = F_copyObjectImpl(m, v484)
	mBase = m.M
	v486 = m.ExcPending
	if v486 != 0 {
		goto L2
	} else {
		goto L121
	}
L121:
	;
	v487 = *(*int32)(unsafe.Add(mBase, uint32(v89)+28))
	v488 = *(*int32)(unsafe.Add(mBase, uint32(v487)+4))
	v489 = F_lappend(m, v488, v485)
	mBase = m.M
	v490 = m.ExcPending
	if v490 != 0 {
		goto L2
	} else {
		goto L122
	}
L122:
	;
	v491 = *(*int32)(unsafe.Add(mBase, uint32(v89)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v491)+4)) = v489
	v493 = *(*int32)(unsafe.Add(mBase, uint32(v485)+4))
	F_cost_qual_eval_node(m, v225, v493, l0)
	mBase = m.M
	v495 = m.ExcPending
	if v495 != 0 {
		goto L2
	} else {
		goto L123
	}
L123:
	;
	v496 = *(*int32)(unsafe.Add(mBase, uint32(v89)+28))
	v497 = *(*float64)(unsafe.Add(mBase, uint32(v225)))
	v498 = *(*float64)(unsafe.Add(mBase, uint32(v496)+16))
	*(*float64)(unsafe.Add(mBase, uint32(v496)+16)) = base.F64_add(v497, v498)
	v501 = *(*int32)(unsafe.Add(mBase, uint32(v89)+28))
	v502 = *(*float64)(unsafe.Add(mBase, uint32(v225)+8))
	v503 = *(*float64)(unsafe.Add(mBase, uint32(v501)+24))
	*(*float64)(unsafe.Add(mBase, uint32(v501)+24)) = base.F64_add(v502, v503)
	v506 = int64(*(*int32)(unsafe.Add(mBase, uint32(v258)+24)))
	v510 = v252 + v506
	goto L75
L124:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v89)+60)) = v513
	v518 = v510
	goto L74
L125:
	;
	goto L57
L126:
	;
	v546 = *(*int32)(unsafe.Add(mBase, uint32(v89)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v546)+32)) = base.I32_wrap_i64(v544)
	m.G0 = v225 + int32(16)
	v551 = *(*int32)(unsafe.Add(mBase, uint32(v89)+60))
	v552 = *(*int32)(unsafe.Add(mBase, uint32(v89)+8))
	v553 = F_bms_del_members(m, v551, v552)
	mBase = m.M
	v554 = m.ExcPending
	if v554 != 0 {
		goto L2
	} else {
		goto L130
	}
L127:
	;
	v544 = v541
	goto L129
L128:
	;
	v544 = v539
	goto L129
L129:
	;
	goto L126
L130:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v89)+60)) = v553
	v556 = F_build_joinrel_restrictlist(m, l0, v89, v47, v48, v73)
	mBase = m.M
	v557 = m.ExcPending
	if v557 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v76))) = v556
	goto L134
L133:
	;
	goto L134
L134:
	;
	v560 = v89 + int32(8)
	v561 = *(*int32)(unsafe.Add(mBase, uint32(v47)+212))
	if v561 == int32(0) {
		goto L136
	} else {
		goto L137
	}
L135:
	;
	v676 = *(*int32)(unsafe.Add(mBase, uint32(v48)+212))
	if v676 == int32(0) {
		v773 = v660
		goto L161
	} else {
		goto L162
	}
L136:
	;
	v660 = int32(0)
	goto L135
L137:
	;
	goto L138
L138:
	;
	v565 = int32(0)
	v566 = *(*int32)(unsafe.Add(mBase, uint32(v561)+4))
	if v566 <= v565 {
		v660 = v565
		goto L135
	} else {
		goto L139
	}
L139:
	;
	v572 = v565
	v573 = int32(0)
	goto L140
L140:
	;
	v588 = *(*int32)(unsafe.Add(mBase, uint32(v561)+12))
	v592 = *(*int32)(unsafe.Add(mBase, uint32(v588+v573<<(uint(int32(2))%32))))
	v593 = *(*int32)(unsafe.Add(mBase, uint32(v592)+32))
	v594 = *(*int32)(unsafe.Add(mBase, uint32(v560)))
	v595 = int32(0)
	if v593 == v595 {
		goto L143
	} else {
		goto L144
	}
L141:
	;
	v660 = v653
	goto L135
L142:
	;
	if v648 == int32(0) {
		goto L156
	} else {
		goto L157
	}
L143:
	;
	v648 = int32(1)
	goto L142
L144:
	;
	goto L145
L145:
	;
	if v594 == int32(0) {
		v641 = v595
		goto L146
	} else {
		goto L147
	}
L146:
	;
	v648 = v641
	goto L142
L147:
	;
	v604 = *(*int32)(unsafe.Add(mBase, uint32(v593)+4))
	v605 = *(*int32)(unsafe.Add(mBase, uint32(v594)+4))
	if v605 < v604 {
		v641 = v595
		goto L146
	} else {
		goto L148
	}
L148:
	;
	v607 = int32(1)
	if v604 <= v607 {
		goto L149
	} else {
		goto L150
	}
L149:
	;
	v610 = v607
	goto L151
L150:
	;
	v610 = v604
	goto L151
L151:
	;
	v611 = int32(8)
	v616 = int32(0)
	goto L152
L152:
	;
	v623 = v616 << (uint(int32(2)) % 32)
	v625 = *(*int32)(unsafe.Add(mBase, uint32(v593+v611+v623)))
	v627 = *(*int32)(unsafe.Add(mBase, uint32(v594+v611+v623)))
	v630 = v625 & (v627 ^ int32(-1))
	v632 = base.B2i32(v630 == int32(0))
	if v630 != 0 {
		v641 = v632
		goto L146
	} else {
		goto L154
	}
L153:
	;
	v641 = v632
	goto L146
L154:
	;
	v634 = v616 + int32(1)
	if v634 != v610 {
		v616 = v634
		goto L152
	} else {
		goto L155
	}
L155:
	;
	goto L153
L156:
	;
	v651 = F_list_append_unique_ptr(m, v572, v592)
	mBase = m.M
	v652 = m.ExcPending
	if v652 != 0 {
		goto L2
	} else {
		goto L159
	}
L157:
	;
	v653 = v572
	goto L158
L158:
	;
	v655 = v573 + int32(1)
	v656 = *(*int32)(unsafe.Add(mBase, uint32(v561)+4))
	if v655 < v656 {
		v572 = v653
		v573 = v655
		goto L140
	} else {
		goto L160
	}
L159:
	;
	v653 = v651
	goto L158
L160:
	;
	goto L141
L161:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v89)+212)) = v773
	v790 = int32(0)
	v791 = *(*int32)(unsafe.Add(mBase, uint32(v89)+8))
	if v791 == v790 {
		goto L188
	} else {
		goto L189
	}
L162:
	;
	v679 = *(*int32)(unsafe.Add(mBase, uint32(v676)+4))
	if v679 <= int32(0) {
		v773 = v660
		goto L161
	} else {
		goto L163
	}
L163:
	;
	v685 = v660
	v686 = int32(0)
	goto L164
L164:
	;
	v701 = *(*int32)(unsafe.Add(mBase, uint32(v676)+12))
	v705 = *(*int32)(unsafe.Add(mBase, uint32(v701+v686<<(uint(int32(2))%32))))
	v706 = *(*int32)(unsafe.Add(mBase, uint32(v705)+32))
	v707 = *(*int32)(unsafe.Add(mBase, uint32(v560)))
	v708 = int32(0)
	if v706 == v708 {
		goto L167
	} else {
		goto L168
	}
L165:
	;
	v773 = v766
	goto L161
L166:
	;
	if v761 == int32(0) {
		goto L180
	} else {
		goto L181
	}
L167:
	;
	v761 = int32(1)
	goto L166
L168:
	;
	goto L169
L169:
	;
	if v707 == int32(0) {
		v754 = v708
		goto L170
	} else {
		goto L171
	}
L170:
	;
	v761 = v754
	goto L166
L171:
	;
	v717 = *(*int32)(unsafe.Add(mBase, uint32(v706)+4))
	v718 = *(*int32)(unsafe.Add(mBase, uint32(v707)+4))
	if v718 < v717 {
		v754 = v708
		goto L170
	} else {
		goto L172
	}
L172:
	;
	v720 = int32(1)
	if v717 <= v720 {
		goto L173
	} else {
		goto L174
	}
L173:
	;
	v723 = v720
	goto L175
L174:
	;
	v723 = v717
	goto L175
L175:
	;
	v724 = int32(8)
	v729 = int32(0)
	goto L176
L176:
	;
	v736 = v729 << (uint(int32(2)) % 32)
	v738 = *(*int32)(unsafe.Add(mBase, uint32(v706+v724+v736)))
	v740 = *(*int32)(unsafe.Add(mBase, uint32(v707+v724+v736)))
	v743 = v738 & (v740 ^ int32(-1))
	v745 = base.B2i32(v743 == int32(0))
	if v743 != 0 {
		v754 = v745
		goto L170
	} else {
		goto L178
	}
L177:
	;
	v754 = v745
	goto L170
L178:
	;
	v747 = v729 + int32(1)
	if v747 != v723 {
		v729 = v747
		goto L176
	} else {
		goto L179
	}
L179:
	;
	goto L177
L180:
	;
	v764 = F_list_append_unique_ptr(m, v685, v705)
	mBase = m.M
	v765 = m.ExcPending
	if v765 != 0 {
		goto L2
	} else {
		goto L183
	}
L181:
	;
	v766 = v685
	goto L182
L182:
	;
	v768 = v686 + int32(1)
	v769 = *(*int32)(unsafe.Add(mBase, uint32(v676)+4))
	if v768 < v769 {
		v685 = v766
		v686 = v768
		goto L164
	} else {
		goto L184
	}
L183:
	;
	v766 = v764
	goto L182
L184:
	;
	goto L165
L185:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v89)+216)) = uint8(v1200)
	F_build_joinrel_partition_info(m, l0, v89, v47, v48, v73, v556)
	mBase = m.M
	v1203 = m.ExcPending
	if v1203 != 0 {
		goto L2
	} else {
		goto L264
	}
L186:
	;
	if int32(0) < v848 {
		goto L197
	} else {
		goto L198
	}
L187:
	;
	v848 = base.I32_ctz(v834) | v835<<(uint(int32(5))%32)
	goto L186
L188:
	;
	v848 = int32(-2)
	goto L186
L189:
	;
	v801 = base.I32_div_s(int32(0), int32(32))
	v802 = *(*int32)(unsafe.Add(mBase, uint32(v791)+4))
	if v802 <= v801 {
		goto L188
	} else {
		goto L190
	}
L190:
	;
	v805 = v791 + int32(8)
	v809 = *(*int32)(unsafe.Add(mBase, uint32(v805+v801<<(uint(int32(2))%32))))
	v812 = v809 & int32(-1)
	if v812 != 0 {
		v834 = v812
		v835 = v801
		goto L187
	} else {
		goto L191
	}
L191:
	;
	v814 = v801 + int32(1)
	if v814 == v802 {
		goto L188
	} else {
		goto L192
	}
L192:
	;
	v817 = v814
	goto L193
L193:
	;
	v824 = *(*int32)(unsafe.Add(mBase, uint32(v805+v817<<(uint(int32(2))%32))))
	if v824 != 0 {
		v834 = v824
		v835 = v817
		goto L187
	} else {
		goto L195
	}
L194:
	;
	goto L188
L195:
	;
	v826 = v817 + int32(1)
	if v826 != v802 {
		v817 = v826
		goto L193
	} else {
		goto L196
	}
L196:
	;
	goto L194
L197:
	;
	v853 = v790
	v854 = v848
	goto L200
L198:
	;
	v943 = v790
	goto L199
L199:
	;
	if v943 == int32(0) {
		goto L220
	} else {
		goto L221
	}
L200:
	;
	v869 = *(*int32)(unsafe.Add(mBase, uint32(l0)+324))
	if v854 == v869 {
		v881 = v853
		goto L202
	} else {
		goto L203
	}
L201:
	;
	v943 = v881
	goto L199
L202:
	;
	if v791 == int32(0) {
		goto L208
	} else {
		goto L209
	}
L203:
	;
	v871 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v875 = *(*int32)(unsafe.Add(mBase, uint32(v871+v854<<(uint(int32(2))%32))))
	if v875 == int32(0) {
		v881 = v853
		goto L202
	} else {
		goto L204
	}
L204:
	;
	v878 = *(*int32)(unsafe.Add(mBase, uint32(v875)+136))
	v879 = F_bms_add_members(m, v853, v878)
	mBase = m.M
	v880 = m.ExcPending
	if v880 != 0 {
		goto L2
	} else {
		goto L205
	}
L205:
	;
	v881 = v879
	goto L202
L206:
	;
	if int32(0) < v938 {
		v853 = v881
		v854 = v938
		goto L200
	} else {
		goto L217
	}
L207:
	;
	v938 = base.I32_ctz(v924) | v925<<(uint(int32(5))%32)
	goto L206
L208:
	;
	v938 = int32(-2)
	goto L206
L209:
	;
	v889 = v854 + int32(1)
	v891 = base.I32_div_s(v889, int32(32))
	v892 = *(*int32)(unsafe.Add(mBase, uint32(v791)+4))
	if v892 <= v891 {
		goto L208
	} else {
		goto L210
	}
L210:
	;
	v895 = v791 + int32(8)
	v899 = *(*int32)(unsafe.Add(mBase, uint32(v895+v891<<(uint(int32(2))%32))))
	v902 = v899 & (int32(-1) << (uint(v889) % 32))
	if v902 != 0 {
		v924 = v902
		v925 = v891
		goto L207
	} else {
		goto L211
	}
L211:
	;
	v904 = v891 + int32(1)
	if v904 == v892 {
		goto L208
	} else {
		goto L212
	}
L212:
	;
	v907 = v904
	goto L213
L213:
	;
	v914 = *(*int32)(unsafe.Add(mBase, uint32(v895+v907<<(uint(int32(2))%32))))
	if v914 != 0 {
		v924 = v914
		v925 = v907
		goto L207
	} else {
		goto L215
	}
L214:
	;
	goto L208
L215:
	;
	v916 = v907 + int32(1)
	if v916 != v892 {
		v907 = v916
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
	if int32(0) <= v1015 {
		goto L229
	} else {
		goto L230
	}
L219:
	;
	v1015 = base.I32_ctz(v1001) | v1002<<(uint(int32(5))%32)
	goto L218
L220:
	;
	v1015 = int32(-2)
	goto L218
L221:
	;
	v968 = base.I32_div_s(int32(0), int32(32))
	v969 = *(*int32)(unsafe.Add(mBase, uint32(v943)+4))
	if v969 <= v968 {
		goto L220
	} else {
		goto L222
	}
L222:
	;
	v972 = v943 + int32(8)
	v976 = *(*int32)(unsafe.Add(mBase, uint32(v972+v968<<(uint(int32(2))%32))))
	v979 = v976 & int32(-1)
	if v979 != 0 {
		v1001 = v979
		v1002 = v968
		goto L219
	} else {
		goto L223
	}
L223:
	;
	v981 = v968 + int32(1)
	if v981 == v969 {
		goto L220
	} else {
		goto L224
	}
L224:
	;
	v984 = v981
	goto L225
L225:
	;
	v991 = *(*int32)(unsafe.Add(mBase, uint32(v972+v984<<(uint(int32(2))%32))))
	if v991 != 0 {
		v1001 = v991
		v1002 = v984
		goto L219
	} else {
		goto L227
	}
L226:
	;
	goto L220
L227:
	;
	v993 = v984 + int32(1)
	if v993 != v969 {
		v984 = v993
		goto L225
	} else {
		goto L228
	}
L228:
	;
	goto L226
L229:
	;
	v1021 = v1015
	goto L232
L230:
	;
	goto L231
L231:
	;
	v1200 = int32(0)
	goto L185
L232:
	;
	v1036 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	v1037 = *(*int32)(unsafe.Add(mBase, uint32(v1036)+12))
	v1041 = *(*int32)(unsafe.Add(mBase, uint32(v1037+v1021<<(uint(int32(2))%32))))
	v1042 = *(*int32)(unsafe.Add(mBase, uint32(v1041)+16))
	if v1042 == int32(0) {
		goto L234
	} else {
		goto L235
	}
L233:
	;
	goto L231
L234:
	;
	if v943 == int32(0) {
		goto L254
	} else {
		goto L255
	}
L235:
	;
	v1045 = *(*int32)(unsafe.Add(mBase, uint32(v1042)+4))
	if v1045 < int32(2) {
		goto L234
	} else {
		goto L236
	}
L236:
	;
	v1048 = *(*int32)(unsafe.Add(mBase, uint32(v1041)+36))
	v1049 = *(*int32)(unsafe.Add(mBase, uint32(v89)+8))
	v1050 = int32(0)
	if v1048 == v1050 {
		goto L238
	} else {
		goto L239
	}
L237:
	;
	if v1103 != 0 {
		goto L234
	} else {
		goto L251
	}
L238:
	;
	v1103 = int32(1)
	goto L237
L239:
	;
	goto L240
L240:
	;
	if v1049 == int32(0) {
		v1096 = v1050
		goto L241
	} else {
		goto L242
	}
L241:
	;
	v1103 = v1096
	goto L237
L242:
	;
	v1059 = *(*int32)(unsafe.Add(mBase, uint32(v1048)+4))
	v1060 = *(*int32)(unsafe.Add(mBase, uint32(v1049)+4))
	if v1060 < v1059 {
		v1096 = v1050
		goto L241
	} else {
		goto L243
	}
L243:
	;
	v1062 = int32(1)
	if v1059 <= v1062 {
		goto L244
	} else {
		goto L245
	}
L244:
	;
	v1065 = v1062
	goto L246
L245:
	;
	v1065 = v1059
	goto L246
L246:
	;
	v1066 = int32(8)
	v1071 = int32(0)
	goto L247
L247:
	;
	v1078 = v1071 << (uint(int32(2)) % 32)
	v1080 = *(*int32)(unsafe.Add(mBase, uint32(v1048+v1066+v1078)))
	v1082 = *(*int32)(unsafe.Add(mBase, uint32(v1049+v1066+v1078)))
	v1085 = v1080 & (v1082 ^ int32(-1))
	v1087 = base.B2i32(v1085 == int32(0))
	if v1085 != 0 {
		v1096 = v1087
		goto L241
	} else {
		goto L249
	}
L248:
	;
	v1096 = v1087
	goto L241
L249:
	;
	v1089 = v1071 + int32(1)
	if v1089 != v1065 {
		v1071 = v1089
		goto L247
	} else {
		goto L250
	}
L250:
	;
	goto L248
L251:
	;
	v1200 = int32(1)
	goto L185
L252:
	;
	if int32(0) <= v1160 {
		v1021 = v1160
		goto L232
	} else {
		goto L263
	}
L253:
	;
	v1160 = base.I32_ctz(v1146) | v1147<<(uint(int32(5))%32)
	goto L252
L254:
	;
	v1160 = int32(-2)
	goto L252
L255:
	;
	v1111 = v1021 + int32(1)
	v1113 = base.I32_div_s(v1111, int32(32))
	v1114 = *(*int32)(unsafe.Add(mBase, uint32(v943)+4))
	if v1114 <= v1113 {
		goto L254
	} else {
		goto L256
	}
L256:
	;
	v1117 = v943 + int32(8)
	v1121 = *(*int32)(unsafe.Add(mBase, uint32(v1117+v1113<<(uint(int32(2))%32))))
	v1124 = v1121 & (int32(-1) << (uint(v1111) % 32))
	if v1124 != 0 {
		v1146 = v1124
		v1147 = v1113
		goto L253
	} else {
		goto L257
	}
L257:
	;
	v1126 = v1113 + int32(1)
	if v1126 == v1114 {
		goto L254
	} else {
		goto L258
	}
L258:
	;
	v1129 = v1126
	goto L259
L259:
	;
	v1136 = *(*int32)(unsafe.Add(mBase, uint32(v1117+v1129<<(uint(int32(2))%32))))
	if v1136 != 0 {
		v1146 = v1136
		v1147 = v1129
		goto L253
	} else {
		goto L261
	}
L260:
	;
	goto L254
L261:
	;
	v1138 = v1129 + int32(1)
	if v1138 != v1114 {
		v1129 = v1138
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
	F_set_joinrel_size_estimates(m, l0, v89, v47, v48, v73, v556)
	mBase = m.M
	v1205 = m.ExcPending
	if v1205 != 0 {
		goto L2
	} else {
		goto L265
	}
L265:
	;
	v1206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+26)))
	if v1206 != int32(1) {
		goto L266
	} else {
		goto L267
	}
L266:
	;
	v1224 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v1225 = F_lappend(m, v1224, v89)
	mBase = m.M
	v1226 = m.ExcPending
	if v1226 != 0 {
		goto L2
	} else {
		goto L273
	}
L267:
	;
	v1209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+26)))
	if v1209 != int32(1) {
		goto L266
	} else {
		goto L268
	}
L268:
	;
	v1212 = F_is_parallel_safe(m, l0, v556)
	mBase = m.M
	v1213 = m.ExcPending
	if v1213 != 0 {
		goto L2
	} else {
		goto L269
	}
L269:
	;
	if v1212 == int32(0) {
		goto L266
	} else {
		goto L270
	}
L270:
	;
	v1216 = *(*int32)(unsafe.Add(mBase, uint32(v89)+28))
	v1217 = *(*int32)(unsafe.Add(mBase, uint32(v1216)+4))
	v1218 = F_is_parallel_safe(m, l0, v1217)
	mBase = m.M
	v1219 = m.ExcPending
	if v1219 != 0 {
		goto L2
	} else {
		goto L271
	}
L271:
	;
	if v1218 == int32(0) {
		goto L266
	} else {
		goto L272
	}
L272:
	;
	v1222 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v89)+26)) = uint8(v1222)
	goto L266
L273:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = v1225
	v1228 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	if v1228 != 0 {
		goto L274
	} else {
		goto L275
	}
L274:
	;
	v1232 = F_hash_search(m, v1228, v560, int32(1), v79+int32(15))
	mBase = m.M
	v1233 = m.ExcPending
	if v1233 != 0 {
		goto L2
	} else {
		goto L277
	}
L275:
	;
	goto L276
L276:
	;
	v1235 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	if v1235 == int32(0) {
		v1252 = v89
		goto L19
	} else {
		goto L278
	}
L277:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1232)+4)) = v89
	goto L276
L278:
	;
	v1238 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v1242 = *(*int32)(unsafe.Add(mBase, uint32(v1235+v1238<<(uint(int32(2))%32))))
	v1243 = F_lappend(m, v1242, v89)
	mBase = m.M
	v1244 = m.ExcPending
	if v1244 != 0 {
		goto L2
	} else {
		goto L279
	}
L279:
	;
	v1245 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v1246 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v1245+v1246<<(uint(int32(2))%32)))) = v1243
	v1252 = v89
	goto L19
L280:
	;
	F_bms_free(m, v44)
	mBase = m.M
	v1347 = m.ExcPending
	if v1347 != 0 {
		goto L2
	} else {
		goto L291
	}
L281:
	;
	v1325 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
	F_populate_joinrel_with_paths(m, l0, v47, v48, v1252, v73, v1325)
	mBase = m.M
	v1327 = m.ExcPending
	if v1327 != 0 {
		goto L2
	} else {
		goto L290
	}
L282:
	;
	v1275 = *(*int32)(unsafe.Add(mBase, uint32(v1272)+12))
	v1278 = v1275
	goto L283
L283:
	;
	v1294 = *(*int32)(unsafe.Add(mBase, uint32(v1278)))
	v1295 = *(*int32)(unsafe.Add(mBase, uint32(v1294)))
	if base.Ui32(int32(2)) <= base.Ui32(v1295-int32(301)) {
		goto L285
	} else {
		goto L286
	}
L284:
	;
	goto L281
L285:
	;
	if v1295 != int32(290) {
		goto L281
	} else {
		goto L288
	}
L286:
	;
	v1278 = v1294 + int32(72)
	goto L283
L287:
	;
	goto L284
L288:
	;
	v1302 = *(*int32)(unsafe.Add(mBase, uint32(v1294)+72))
	if v1302 == int32(0) {
		goto L280
	} else {
		goto L289
	}
L289:
	;
	goto L287
L290:
	;
	goto L280
L291:
	;
	v1351 = v1252
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
