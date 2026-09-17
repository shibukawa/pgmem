package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_AlterConstrUpdateConstraintEntry(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	v6 = F_heap_copytuple(m, l2)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(v6)+16))
		v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+22)))
		v10 = v8 + v9
		v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
		if v11 == int32(1) {
			v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+9)))
			*(*uint8)(unsafe.Add(mBase, uint32(v10)+76)) = uint8(v14)
			*(*uint8)(unsafe.Add(mBase, uint32(v10)+75)) = uint8(v14)
		} else {
		}
		v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+10)))
		if v18 == int32(1) {
			v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+11)))
			*(*uint8)(unsafe.Add(mBase, uint32(v10)+73)) = uint8(v21)
			v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)))
			*(*uint8)(unsafe.Add(mBase, uint32(v10)+74)) = uint8(v23)
		} else {
		}
		v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+13)))
		if v25 == int32(1) {
			v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+14)))
			*(*uint8)(unsafe.Add(mBase, uint32(v10)+106)) = uint8(v28)
		} else {
		}
		F_CatalogTupleUpdate(m, l1, v6+int32(4), v6)
		mBase = m.M
		v33 = m.ExcPending
		if v33 != 0 {
			return
		} else {
			v35 = *(*int32)(unsafe.Add(mBase, _c_F_AlterConstrUpdateConstraintEntry[0]))
			if v35 != 0 {
				v37 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
				v38 = int32(0)
				F_RunObjectPostAlterHook(m, int32(2606), v37, v38, v38, v38)
				mBase = m.M
				v42 = m.ExcPending
				if v42 != 0 {
					return
				} else {
					v43 = *(*int32)(unsafe.Add(mBase, uint32(v10)+80))
					F_CacheInvalidateRelcacheByRelid(m, v43)
					mBase = m.M
					v45 = m.ExcPending
					if v45 != 0 {
						return
					} else {
						F_pfree(m, v6)
						mBase = m.M
						v47 = m.ExcPending
						if v47 != 0 {
							return
						} else {
							return
						}
					}
				}
			} else {
				v43 = *(*int32)(unsafe.Add(mBase, uint32(v10)+80))
				F_CacheInvalidateRelcacheByRelid(m, v43)
				mBase = m.M
				v45 = m.ExcPending
				if v45 != 0 {
					return
				} else {
					F_pfree(m, v6)
					mBase = m.M
					v47 = m.ExcPending
					if v47 != 0 {
						return
					} else {
						return
					}
				}
			}
		}
	}
}
func F_AlterSystemSetConfigFile(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v67 int32
	_ = v67
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v97 int32
	_ = v97
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v137 int32
	_ = v137
	var v143 int32
	_ = v143
	var v152 int32
	_ = v152
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v172 int32
	_ = v172
	var v178 int32
	_ = v178
	var v185 int32
	_ = v185
	var v193 int32
	_ = v193
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v215 int32
	_ = v215
	var v221 int32
	_ = v221
	var v230 int32
	_ = v230
	var v238 int32
	_ = v238
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v262 int32
	_ = v262
	var v268 int32
	_ = v268
	var v278 int32
	_ = v278
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v323 int32
	_ = v323
	var v329 int32
	_ = v329
	var v331 int32
	_ = v331
	var v334 int32
	_ = v334
	var v343 int32
	_ = v343
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v363 int32
	_ = v363
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v370 int32
	_ = v370
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v393 int32
	_ = v393
	var v395 int32
	_ = v395
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v403 int32
	_ = v403
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v410 int32
	_ = v410
	var v413 int32
	_ = v413
	var v415 int32
	_ = v415
	var v420 int32
	_ = v420
	var v422 int32
	_ = v422
	var v427 int32
	_ = v427
	var v429 int32
	_ = v429
	var v432 int32
	_ = v432
	var v434 int32
	_ = v434
	var v437 int32
	_ = v437
	var v449 int32
	_ = v449
	var v466 int32
	_ = v466
	var v472 int32
	_ = v472
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v537 int32
	_ = v537
	var v541 int32
	_ = v541
	var v550 int32
	_ = v550
	var v556 int32
	_ = v556
	var v563 int32
	_ = v563
	var v571 int32
	_ = v571
	var v594 int32
	_ = v594
	var v599 int32
	_ = v599
	var v600 int32
	_ = v600
	var v613 int32
	_ = v613
	var v614 int32
	_ = v614
	var v619 int32
	_ = v619
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v632 int32
	_ = v632
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v649 int32
	_ = v649
	var v654 int32
	_ = v654
	var v663 int32
	_ = v663
	var v671 int32
	_ = v671
	var v676 int32
	_ = v676
	var v683 int32
	_ = v683
	var v684 int32
	_ = v684
	var v693 int32
	_ = v693
	var v699 int32
	_ = v699
	var v708 int32
	_ = v708
	var v716 int32
	_ = v716
	var v720 int32
	_ = v720
	var v721 int32
	_ = v721
	var v725 int32
	_ = v725
	var v732 int32
	_ = v732
	var v734 int32
	_ = v734
	var v742 int32
	_ = v742
	var v743 int32
	_ = v743
	var v751 int32
	_ = v751
	var v755 int32
	_ = v755
	var v760 int32
	_ = v760
	var v761 int32
	_ = v761
	var v766 int32
	_ = v766
	var v778 int32
	_ = v778
	var v779 int32
	_ = v779
	var v789 int32
	_ = v789
	var v796 int32
	_ = v796
	var v801 int32
	_ = v801
	var v802 int32
	_ = v802
	var v807 int32
	_ = v807
	var v808 int32
	_ = v808
	var v813 int32
	_ = v813
	var v818 int32
	_ = v818
	var v819 int32
	_ = v819
	var v844 int32
	_ = v844
	var v845 int32
	_ = v845
	var v849 int32
	_ = v849
	var v850 int32
	_ = v850
	var v855 int32
	_ = v855
	var v856 int32
	_ = v856
	var v864 int32
	_ = v864
	var v865 int32
	_ = v865
	var v866 int32
	_ = v866
	var v873 int32
	_ = v873
	var v878 int32
	_ = v878
	var v887 int32
	_ = v887
	var v892 int32
	_ = v892
	var v893 int32
	_ = v893
	var v906 int32
	_ = v906
	var v907 int32
	_ = v907
	var v912 int32
	_ = v912
	var v916 int32
	_ = v916
	var v917 int32
	_ = v917
	var v922 int32
	_ = v922
	var v935 int32
	_ = v935
	var v936 int32
	_ = v936
	var v937 int32
	_ = v937
	var v943 int32
	_ = v943
	var v949 int32
	_ = v949
	var v951 int32
	_ = v951
	var v952 int32
	_ = v952
	var v963 int32
	_ = v963
	var v968 int32
	_ = v968
	var v977 int32
	_ = v977
	var v985 int32
	_ = v985
	var v988 int32
	_ = v988
	var v990 int32
	_ = v990
	var v992 int32
	_ = v992
	var v1000 int32
	_ = v1000
	var v1001 int32
	_ = v1001
	var v1002 int32
	_ = v1002
	var v1007 int32
	_ = v1007
	var v1023 int32
	_ = v1023
	var v1024 int32
	_ = v1024
	var v1026 int32
	_ = v1026
	var v1028 int32
	_ = v1028
	var v1034 int32
	_ = v1034
	var v1040 int32
	_ = v1040
	var v1050 int32
	_ = v1050
	var v1051 int32
	_ = v1051
	var v1052 int32
	_ = v1052
	var v1053 int32
	_ = v1053
	var v1064 int32
	_ = v1064
	var v1075 int32
	_ = v1075
	var v1076 int32
	_ = v1076
	var v1077 int32
	_ = v1077
	var v1083 int32
	_ = v1083
	var v1088 int32
	_ = v1088
	var v1094 int32
	_ = v1094
	var v1095 int32
	_ = v1095
	var v1099 int32
	_ = v1099
	var v1102 int32
	_ = v1102
	var v1103 int32
	_ = v1103
	var v1107 int32
	_ = v1107
	var v1122 int32
	_ = v1122
	var v1124 int32
	_ = v1124
	var v1126 int32
	_ = v1126
	var v1132 int32
	_ = v1132
	var v1133 int32
	_ = v1133
	var v1145 int32
	_ = v1145
	var v1149 int32
	_ = v1149
	var v1150 int32
	_ = v1150
	var v1162 int32
	_ = v1162
	var v1165 int32
	_ = v1165
	var v1166 int32
	_ = v1166
	var v1168 int32
	_ = v1168
	var v1170 int32
	_ = v1170
	var v1180 int32
	_ = v1180
	var v1182 int32
	_ = v1182
	var v1191 int32
	_ = v1191
	var v1203 int32
	_ = v1203
	var v1214 int32
	_ = v1214
	var v1225 int32
	_ = v1225
	var v1251 int32
	_ = v1251
	var v1257 int32
	_ = v1257
	var v1264 int32
	_ = v1264
	var v1272 int32
	_ = v1272
	var v1277 int32
	_ = v1277
	var v1279 int32
	_ = v1279
	var v1286 int32
	_ = v1286
	var v1293 int32
	_ = v1293
	var v1294 int32
	_ = v1294
	var v1295 int32
	_ = v1295
	var v1296 int32
	_ = v1296
	var v1299 int32
	_ = v1299
	var v1311 int32
	_ = v1311
	var v1316 int32
	_ = v1316
	var v1327 int32
	_ = v1327
	var v1335 int32
	_ = v1335
	var v1336 int32
	_ = v1336
	var v1358 int32
	_ = v1358
	var v1363 int32
	_ = v1363
	var v1367 int32
	_ = v1367
	var v1372 int32
	_ = v1372
	var v1379 int32
	_ = v1379
	var v1384 int32
	_ = v1384
	var v1393 int32
	_ = v1393
	var v1401 int32
	_ = v1401
	var v1405 int32
	_ = v1405
	var v1407 int32
	_ = v1407
	var v1411 int32
	_ = v1411
	var v1412 int32
	_ = v1412
	var v1423 int32
	_ = v1423
	var v1424 int32
	_ = v1424
	var v1432 int32
	_ = v1432
	var v1434 int32
	_ = v1434
	var v1439 int32
	_ = v1439
	var v1443 int32
	_ = v1443
	var v1451 int32
	_ = v1451
	var v1457 int32
	_ = v1457
	var v1458 int32
	_ = v1458
	var v1464 int32
	_ = v1464
	var v1469 int32
	_ = v1469
	var v1471 int32
	_ = v1471
	var v1483 int32
	_ = v1483
	var v1488 int32
	_ = v1488
	var v1499 int32
	_ = v1499
	var v1507 int32
	_ = v1507
	var v1524 int32
	_ = v1524
	var v1525 int64
	_ = v1525
	var v1529 int32
	_ = v1529
	var v1531 int32
	_ = v1531
	var v1532 int32
	_ = v1532
	var v1535 int32
	_ = v1535
	var v1537 int32
	_ = v1537
	var v1539 int32
	_ = v1539
	var v1540 int32
	_ = v1540
	var v1541 int32
	_ = v1541
	var v1542 int32
	_ = v1542
	var v1544 int32
	_ = v1544
	v2 = int32(0)
	v17 = m.G0
	v19 = v17 - int32(2592)
	m.G0 = v19
	v22 = l0
	v23 = v19
	v24 = v2
	v25 = v2
	v26 = v2
	v30 = int32(-1)
	v31 = v2
	goto L1
L1:
	;
	goto L4
L2:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L3:
	;
	goto L2
L4:
	;
	if v30 != int32(1) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L3
L6:
	;
	v1524 = int32(m.ExcTag)
	v1525 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v1524 == int32(0) {
		goto L295
	} else {
		goto L296
	}
L7:
	;
	v40 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2560)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2556)) = v40
	v45 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_AlterSystemSetConfigFile[0])))
	if v45 == v40 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v1000 = v24
	v1001 = v25
	v1002 = v26
	v1007 = v31
	goto L9
L9:
	;
	if v1007 == int32(0) {
		goto L210
	} else {
		goto L211
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v26
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L6
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)+8))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v76)+4))
	switch v79 {
	case 0:
		goto L22
	case 1, 4:
		v111 = v26
		v112 = int32(0)
		goto L21
	default:
		goto L23
	case 5:
		goto L20
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v26
	F_errcode(m, int32(1088))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L6
	} else {
		goto L14
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v26
	F_errmsg(m, int32(_a_F_AlterSystemSetConfigFile_0), int32(0))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L6
	} else {
		goto L15
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v26
	F_errfinish(m, int32(_a_F_AlterSystemSetConfigFile_1), int32(_a_F_AlterSystemSetConfigFile_2), int32(_a_F_AlterSystemSetConfigFile_3))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L6
	} else {
		goto L16
	}
L16:
	;
	goto L3
L17:
	;
	v935 = *(*int32)(unsafe.Add(mBase, _c_F_AlterSystemSetConfigFile[1]))
	if v935 != 0 {
		goto L193
	} else {
		goto L194
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v26
	*(*int32)(unsafe.Add(mBase, uint32(v23)+208)) = int32(_a_F_AlterSystemSetConfigFile_4)
	v887 = v23 + int32(1520)
	v892 = F_pg_snprintf(m, v887, int32(1024), int32(_a_F_AlterSystemSetConfigFile_5), v23+int32(208))
	mBase = m.M
	v893 = m.ExcPending
	if v893 != 0 {
		goto L6
	} else {
		goto L190
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v111
	v201 = F_find_option(m, v77, int32(0), int32(1), int32(10))
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L6
	} else {
		goto L44
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v26
	v164 = F_superuser(m)
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L6
	} else {
		goto L36
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v111
	v116 = F_superuser(m)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L6
	} else {
		goto L28
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v26
	v109 = F_ExtractSetVariableArgs(m, v76)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L6
	} else {
		goto L27
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v26
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L6
	} else {
		goto L24
	}
L24:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v87)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v26
	*(*int32)(unsafe.Add(mBase, uint32(v23)+48)) = v88
	F_errmsg_internal(m, int32(_a_F_AlterSystemSetConfigFile_6), v23+int32(48))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L6
	} else {
		goto L25
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v26
	F_errfinish(m, int32(_a_F_AlterSystemSetConfigFile_1), int32(_a_F_AlterSystemSetConfigFile_7), int32(_a_F_AlterSystemSetConfigFile_3))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L6
	} else {
		goto L26
	}
L26:
	;
	goto L3
L27:
	;
	v111 = v109
	v112 = v109
	goto L21
L28:
	;
	if v116 != 0 {
		goto L19
	} else {
		goto L29
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v111
	v122 = *(*int32)(unsafe.Add(mBase, _c_F_AlterSystemSetConfigFile[2]))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v111
	v127 = F_pg_parameter_aclcheck(m, v77, v122, int64(8192))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L6
	} else {
		goto L30
	}
L30:
	;
	if v127 == int32(0) {
		goto L19
	} else {
		goto L31
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v111
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L6
	} else {
		goto L32
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v111
	F_errcode(m, int32(16797828))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L6
	} else {
		goto L33
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v111
	*(*int32)(unsafe.Add(mBase, uint32(v23)+176)) = v77
	F_errmsg(m, int32(_a_F_AlterSystemSetConfigFile_8), v23+int32(176))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L6
	} else {
		goto L34
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v111
	F_errfinish(m, int32(_a_F_AlterSystemSetConfigFile_1), int32(_a_F_AlterSystemSetConfigFile_9), int32(_a_F_AlterSystemSetConfigFile_3))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L6
	} else {
		goto L35
	}
L35:
	;
	goto L3
L36:
	;
	if v164 != 0 {
		goto L18
	} else {
		goto L37
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v26
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L6
	} else {
		goto L38
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v26
	F_errcode(m, int32(16797828))
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L6
	} else {
		goto L39
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v26
	F_errmsg(m, int32(_a_F_AlterSystemSetConfigFile_10), int32(0))
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L6
	} else {
		goto L40
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v26
	F_errfinish(m, int32(_a_F_AlterSystemSetConfigFile_1), int32(_a_F_AlterSystemSetConfigFile_11), int32(_a_F_AlterSystemSetConfigFile_3))
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L6
	} else {
		goto L41
	}
L41:
	;
	goto L3
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v111
	*(*int32)(unsafe.Add(mBase, uint32(v23)+128)) = int32(_a_F_AlterSystemSetConfigFile_4)
	v594 = v23 + int32(1520)
	v599 = F_pg_snprintf(m, v594, int32(1024), int32(_a_F_AlterSystemSetConfigFile_5), v23+int32(128))
	mBase = m.M
	v600 = m.ExcPending
	if v600 != 0 {
		goto L6
	} else {
		goto L123
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v111
	v534 = int32(10)
	v535 = F___strchrnul(m, v112, v534)
	mBase = m.M
	v537 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v535))))
	if v537 == v534 {
		goto L115
	} else {
		goto L116
	}
L44:
	;
	if v201 != 0 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v201)+4))
	if v203 != 0 {
		goto L49
	} else {
		goto L50
	}
L46:
	;
	goto L47
L47:
	;
	if v112 != 0 {
		goto L73
	} else {
		goto L74
	}
L48:
	;
	if v112 == int32(0) {
		goto L42
	} else {
		goto L57
	}
L49:
	;
	v204 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v201)+21)))
	if v204&int32(33) == int32(0) {
		goto L48
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v111
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L6
	} else {
		goto L53
	}
L52:
	;
	goto L51
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v111
	F_errcode(m, int32(33685829))
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L6
	} else {
		goto L54
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v111
	*(*int32)(unsafe.Add(mBase, uint32(v23)+144)) = v77
	F_errmsg(m, int32(_a_F_AlterSystemSetConfigFile_12), v23+int32(144))
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L6
	} else {
		goto L55
	}
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v111
	F_errfinish(m, int32(_a_F_AlterSystemSetConfigFile_1), int32(_a_F_AlterSystemSetConfigFile_13), int32(_a_F_AlterSystemSetConfigFile_3))
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L6
	} else {
		goto L56
	}
L56:
	;
	goto L3
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v111
	*(*int32)(unsafe.Add(mBase, uint32(v23)+484)) = int32(0)
	v252 = F_parse_and_validate_value(m, v201, v112, int32(3), int32(21), v23+int32(488), v23+int32(484))
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L6
	} else {
		goto L58
	}
L58:
	;
	if v252 == int32(0) {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v111
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L6
	} else {
		goto L62
	}
L60:
	;
	goto L61
L61:
	;
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v201)+24))
	if v287 != int32(3) {
		goto L66
	} else {
		goto L67
	}
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v111
	F_errcode(m, int32(50856066))
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L6
	} else {
		goto L63
	}
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v111
	*(*int32)(unsafe.Add(mBase, uint32(v23)+164)) = v112
	*(*int32)(unsafe.Add(mBase, uint32(v23)+160)) = v77
	F_errmsg(m, int32(_a_F_AlterSystemSetConfigFile_14), v23+int32(160))
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L6
	} else {
		goto L64
	}
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v111
	F_errfinish(m, int32(_a_F_AlterSystemSetConfigFile_1), int32(_a_F_AlterSystemSetConfigFile_15), int32(_a_F_AlterSystemSetConfigFile_3))
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L6
	} else {
		goto L65
	}
L65:
	;
	goto L3
L66:
	;
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v23)+484))
	if v299 == int32(0) {
		goto L43
	} else {
		goto L70
	}
L67:
	;
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v23)+488))
	if v290 == int32(0) {
		goto L66
	} else {
		goto L68
	}
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v111
	F_pfree(m, v290)
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L6
	} else {
		goto L69
	}
L69:
	;
	goto L66
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v111
	F_pfree(m, v299)
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L6
	} else {
		goto L71
	}
L71:
	;
	goto L43
L72:
	;
	if v112 == int32(0) {
		goto L42
	} else {
		goto L113
	}
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v111
	v495 = F_assignable_custom_variable_name(m, v77, int32(0), int32(21))
	mBase = m.M
	v496 = m.ExcPending
	if v496 != 0 {
		goto L6
	} else {
		goto L112
	}
L74:
	;
	v307 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77))))
	if v307 == int32(0) {
		goto L73
	} else {
		goto L75
	}
L75:
	;
	v318 = int32(0)
	v319 = v77
	v320 = v307
	v323 = int32(1)
	goto L76
L76:
	;
	v329 = v320 & int32(255)
	v331 = base.B2i32(v329 != int32(46))
	if v331 == int32(0) {
		goto L79
	} else {
		goto L80
	}
L77:
	;
	if v466&v331 != 0 {
		goto L72
	} else {
		goto L111
	}
L78:
	;
	v472 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v319)+1)))
	if v472 != 0 {
		v318 = v466
		v319 = v319 + int32(1)
		v320 = v472
		v323 = base.B2i32(v329 == int32(46))
		goto L76
	} else {
		goto L110
	}
L79:
	;
	v334 = int32(1)
	if v323&v334 == int32(0) {
		v466 = v334
		goto L78
	} else {
		goto L82
	}
L80:
	;
	goto L81
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v111
	v343 = base.I32_extend8_s(v320)
	goto L87
L82:
	;
	goto L73
L83:
	;
	if v449|base.B2i32(v343 < int32(0)) != 0 {
		v466 = v318
		goto L78
	} else {
		goto L108
	}
L84:
	;
	v449 = int32(0)
	goto L83
L85:
	;
	v427 = v420
	v429 = v422
	goto L102
L86:
	;
	if base.B2i32(v366 != v367) == int32(0) {
		goto L84
	} else {
		goto L93
	}
L87:
	;
	v358 = int32(_a_F_AlterSystemSetConfigFile_16)
	v360 = int32(54)
	goto L88
L88:
	;
	v363 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v358))))
	if v363 == v343&int32(255) {
		v420 = v358
		v422 = v360
		goto L85
	} else {
		goto L90
	}
L89:
	;
	goto L86
L90:
	;
	v365 = int32(1)
	v366 = v360 - v365
	v367 = int32(0)
	v370 = v358 + v365
	if v370&int32(3) == v367 {
		goto L86
	} else {
		goto L91
	}
L91:
	;
	if v366 != 0 {
		v358 = v370
		v360 = v366
		goto L88
	} else {
		goto L92
	}
L92:
	;
	goto L89
L93:
	;
	v383 = v343 & int32(255)
	v384 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v370))))
	if base.B2i32(v383 == v384)|base.B2i32(base.Ui32(v366) < base.Ui32(int32(4))) == int32(0) {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	v393 = v370
	v395 = v366
	goto L97
L95:
	;
	v413 = v370
	v415 = v366
	goto L96
L96:
	;
	if v415 == int32(0) {
		goto L84
	} else {
		goto L101
	}
L97:
	;
	v399 = *(*int32)(unsafe.Add(mBase, uint32(v393)))
	v400 = v399 ^ v383*int32(16843009)
	v403 = int32(-2139062144)
	if (int32(16843008)-v400|v400)&v403 != v403 {
		v420 = v393
		v422 = v395
		goto L85
	} else {
		goto L99
	}
L98:
	;
	v413 = v408
	v415 = v410
	goto L96
L99:
	;
	v407 = int32(4)
	v408 = v393 + v407
	v410 = v395 - v407
	if base.Ui32(int32(3)) < base.Ui32(v410) {
		v393 = v408
		v395 = v410
		goto L97
	} else {
		goto L100
	}
L100:
	;
	goto L98
L101:
	;
	v420 = v413
	v422 = v415
	goto L85
L102:
	;
	v432 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v427))))
	if v343&int32(255) == v432 {
		goto L104
	} else {
		goto L105
	}
L103:
	;
	goto L84
L104:
	;
	v449 = v427
	goto L83
L105:
	;
	goto L106
L106:
	;
	v434 = int32(1)
	v437 = v429 - v434
	if v437 != 0 {
		v427 = v427 + v434
		v429 = v437
		goto L102
	} else {
		goto L107
	}
L107:
	;
	goto L103
L108:
	;
	if base.B2i32(int64(1)<<(uint(base.I64_extend_i32_u(v343))%64)&int64(287948969894477825) == int64(0))|(v323&int32(1)|base.B2i32(base.Ui32(int32(63)) < base.Ui32(v329))) != 0 {
		goto L73
	} else {
		goto L109
	}
L109:
	;
	v466 = v318
	goto L78
L110:
	;
	goto L77
L111:
	;
	goto L73
L112:
	;
	goto L72
L113:
	;
	goto L43
L114:
	;
	if v541 == int32(0) {
		goto L42
	} else {
		goto L118
	}
L115:
	;
	v541 = v535
	goto L117
L116:
	;
	v541 = int32(0)
	goto L117
L117:
	;
	goto L114
L118:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v111
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v550 = m.ExcPending
	if v550 != 0 {
		goto L6
	} else {
		goto L119
	}
L119:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v111
	F_errcode(m, int32(50856066))
	mBase = m.M
	v556 = m.ExcPending
	if v556 != 0 {
		goto L6
	} else {
		goto L120
	}
L120:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v111
	F_errmsg(m, int32(_a_F_AlterSystemSetConfigFile_17), int32(0))
	mBase = m.M
	v563 = m.ExcPending
	if v563 != 0 {
		goto L6
	} else {
		goto L121
	}
L121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v111
	F_errfinish(m, int32(_a_F_AlterSystemSetConfigFile_1), int32(_a_F_AlterSystemSetConfigFile_18), int32(_a_F_AlterSystemSetConfigFile_3))
	mBase = m.M
	v571 = m.ExcPending
	if v571 != 0 {
		goto L6
	} else {
		goto L122
	}
L122:
	;
	goto L3
L123:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v111
	*(*int32)(unsafe.Add(mBase, uint32(v23)+116)) = int32(_a_F_AlterSystemSetConfigFile_19)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+112)) = v594
	v613 = F_pg_snprintf(m, v23+int32(496), int32(1024), int32(_a_F_AlterSystemSetConfigFile_20), v23+int32(112))
	mBase = m.M
	v614 = m.ExcPending
	if v614 != 0 {
		goto L6
	} else {
		goto L124
	}
L124:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v111
	v619 = *(*int32)(unsafe.Add(mBase, _c_F_AlterSystemSetConfigFile[3]))
	v623 = F_LWLockAcquire(m, v619+int32(_a_F_AlterSystemSetConfigFile_21), int32(0))
	mBase = m.M
	v624 = m.ExcPending
	if v624 != 0 {
		goto L6
	} else {
		goto L125
	}
L125:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v111
	v632 = F___fstatat(m, int32(-100), v594, v23+int32(384), int32(0))
	mBase = m.M
	goto L126
L126:
	;
	if v632 == int32(0) {
		goto L127
	} else {
		goto L128
	}
L127:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v111
	v639 = F_AllocateFile(m, v594, int32(_a_F_AlterSystemSetConfigFile_22))
	mBase = m.M
	v640 = m.ExcPending
	if v640 != 0 {
		goto L6
	} else {
		goto L130
	}
L128:
	;
	goto L129
L129:
	;
	v725 = *(*int32)(unsafe.Add(mBase, uint32(v23)+2560))
	if v725 != 0 {
		goto L147
	} else {
		goto L148
	}
L130:
	;
	if v639 == int32(0) {
		goto L131
	} else {
		goto L132
	}
L131:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v111
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v649 = m.ExcPending
	if v649 != 0 {
		goto L6
	} else {
		goto L134
	}
L132:
	;
	goto L133
L133:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v111
	v676 = v23 + int32(1520)
	v683 = F_ParseConfigFp(m, v639, v676, int32(0), int32(15), v23+int32(2560), v23+int32(2556))
	mBase = m.M
	v684 = m.ExcPending
	if v684 != 0 {
		goto L6
	} else {
		goto L138
	}
L134:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v111
	F_errcode_for_file_access(m)
	mBase = m.M
	v654 = m.ExcPending
	if v654 != 0 {
		goto L6
	} else {
		goto L135
	}
L135:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v111
	*(*int32)(unsafe.Add(mBase, uint32(v23)+64)) = v594
	F_errmsg(m, int32(_a_F_AlterSystemSetConfigFile_23), v23-int32(-64))
	mBase = m.M
	v663 = m.ExcPending
	if v663 != 0 {
		goto L6
	} else {
		goto L136
	}
L136:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v111
	F_errfinish(m, int32(_a_F_AlterSystemSetConfigFile_1), int32(_a_F_AlterSystemSetConfigFile_24), int32(_a_F_AlterSystemSetConfigFile_3))
	mBase = m.M
	v671 = m.ExcPending
	if v671 != 0 {
		goto L6
	} else {
		goto L137
	}
L137:
	;
	goto L3
L138:
	;
	if v683 == int32(0) {
		goto L139
	} else {
		goto L140
	}
L139:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v111
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v693 = m.ExcPending
	if v693 != 0 {
		goto L6
	} else {
		goto L142
	}
L140:
	;
	goto L141
L141:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v111
	v720 = F_FreeFile(m, v639)
	mBase = m.M
	v721 = m.ExcPending
	if v721 != 0 {
		goto L6
	} else {
		goto L146
	}
L142:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v111
	F_errcode(m, int32(22))
	mBase = m.M
	v699 = m.ExcPending
	if v699 != 0 {
		goto L6
	} else {
		goto L143
	}
L143:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v111
	*(*int32)(unsafe.Add(mBase, uint32(v23)+96)) = v676
	F_errmsg(m, int32(_a_F_AlterSystemSetConfigFile_25), v23+int32(96))
	mBase = m.M
	v708 = m.ExcPending
	if v708 != 0 {
		goto L6
	} else {
		goto L144
	}
L144:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v111
	F_errfinish(m, int32(_a_F_AlterSystemSetConfigFile_1), int32(_a_F_AlterSystemSetConfigFile_26), int32(_a_F_AlterSystemSetConfigFile_3))
	mBase = m.M
	v716 = m.ExcPending
	if v716 != 0 {
		goto L6
	} else {
		goto L145
	}
L145:
	;
	goto L3
L146:
	;
	goto L129
L147:
	;
	v732 = int32(0)
	v734 = v725
	goto L150
L148:
	;
	goto L149
L149:
	;
	if v112 == int32(0) {
		v922 = v111
		goto L17
	} else {
		goto L182
	}
L150:
	;
	v742 = *(*int32)(unsafe.Add(mBase, uint32(v734)))
	v743 = *(*int32)(unsafe.Add(mBase, uint32(v734)+24))
	v751 = v742
	v755 = v77
	goto L154
L151:
	;
	goto L149
L152:
	;
	if v743 != 0 {
		v732 = v819
		v734 = v743
		goto L150
	} else {
		goto L181
	}
L153:
	;
	if v760 != 0 {
		goto L167
	} else {
		goto L168
	}
L154:
	;
	v760 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v755))))
	v761 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v751))))
	if v761 == int32(0) {
		goto L153
	} else {
		goto L156
	}
L155:
	;
	v819 = v734
	goto L152
L156:
	;
	if v760 == int32(0) {
		goto L157
	} else {
		goto L158
	}
L157:
	;
	v819 = v734
	goto L152
L158:
	;
	goto L159
L159:
	;
	v766 = int32(1)
	if base.Ui32((v760-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L160
	} else {
		goto L161
	}
L160:
	;
	v778 = v760 | int32(32)
	goto L162
L161:
	;
	v778 = v760
	goto L162
L162:
	;
	v779 = int32(255)
	if base.Ui32((v761-int32(65))&v779) < base.Ui32(int32(26)) {
		goto L163
	} else {
		goto L164
	}
L163:
	;
	v789 = v761 | int32(32)
	goto L165
L164:
	;
	v789 = v761
	goto L165
L165:
	;
	if v778&v779 == v789 {
		v751 = v751 + v766
		v755 = v755 + v766
		goto L154
	} else {
		goto L166
	}
L166:
	;
	goto L155
L167:
	;
	v819 = v734
	goto L152
L168:
	;
	goto L169
L169:
	;
	if v732 != 0 {
		goto L171
	} else {
		goto L172
	}
L170:
	;
	if v743 == int32(0) {
		goto L174
	} else {
		goto L175
	}
L171:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v732)+24)) = v743
	goto L170
L172:
	;
	goto L173
L173:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2560)) = v743
	goto L170
L174:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2556)) = v732
	goto L176
L175:
	;
	goto L176
L176:
	;
	v796 = *(*int32)(unsafe.Add(mBase, uint32(v734)))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v111
	F_pfree(m, v796)
	mBase = m.M
	v801 = m.ExcPending
	if v801 != 0 {
		goto L6
	} else {
		goto L177
	}
L177:
	;
	v802 = *(*int32)(unsafe.Add(mBase, uint32(v734)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v111
	F_pfree(m, v802)
	mBase = m.M
	v807 = m.ExcPending
	if v807 != 0 {
		goto L6
	} else {
		goto L178
	}
L178:
	;
	v808 = *(*int32)(unsafe.Add(mBase, uint32(v734)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v111
	F_pfree(m, v808)
	mBase = m.M
	v813 = m.ExcPending
	if v813 != 0 {
		goto L6
	} else {
		goto L179
	}
L179:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v111
	F_pfree(m, v734)
	mBase = m.M
	v818 = m.ExcPending
	if v818 != 0 {
		goto L6
	} else {
		goto L180
	}
L180:
	;
	v819 = v732
	goto L152
L181:
	;
	goto L151
L182:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v111
	v844 = F_palloc(m, int32(28))
	mBase = m.M
	v845 = m.ExcPending
	if v845 != 0 {
		goto L6
	} else {
		goto L183
	}
L183:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v111
	v849 = F_pstrdup(m, v77)
	mBase = m.M
	v850 = m.ExcPending
	if v850 != 0 {
		goto L6
	} else {
		goto L184
	}
L184:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v844))) = v849
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v111
	v855 = F_pstrdup(m, v112)
	mBase = m.M
	v856 = m.ExcPending
	if v856 != 0 {
		goto L6
	} else {
		goto L185
	}
L185:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v844)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v844)+4)) = v855
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v111
	v864 = F_pstrdup(m, int32(_a_F_AlterSystemSetConfigFile_27))
	mBase = m.M
	v865 = m.ExcPending
	if v865 != 0 {
		goto L6
	} else {
		goto L186
	}
L186:
	;
	v866 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v844)+24)) = v866
	*(*uint16)(unsafe.Add(mBase, uint32(v844)+20)) = uint16(v866)
	*(*int32)(unsafe.Add(mBase, uint32(v844)+16)) = v866
	*(*int32)(unsafe.Add(mBase, uint32(v844)+12)) = v864
	v873 = *(*int32)(unsafe.Add(mBase, uint32(v23)+2560))
	if v873 == v866 {
		goto L187
	} else {
		goto L188
	}
L187:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2560)) = v844
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2556)) = v844
	v922 = v111
	goto L17
L188:
	;
	goto L189
L189:
	;
	v878 = *(*int32)(unsafe.Add(mBase, uint32(v23)+2556))
	*(*int32)(unsafe.Add(mBase, uint32(v878)+24)) = v844
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2556)) = v844
	v922 = v111
	goto L17
L190:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v26
	*(*int32)(unsafe.Add(mBase, uint32(v23)+196)) = int32(_a_F_AlterSystemSetConfigFile_19)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+192)) = v887
	v906 = F_pg_snprintf(m, v23+int32(496), int32(1024), int32(_a_F_AlterSystemSetConfigFile_20), v23+int32(192))
	mBase = m.M
	v907 = m.ExcPending
	if v907 != 0 {
		goto L6
	} else {
		goto L191
	}
L191:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v26
	v912 = *(*int32)(unsafe.Add(mBase, _c_F_AlterSystemSetConfigFile[3]))
	v916 = F_LWLockAcquire(m, v912+int32(_a_F_AlterSystemSetConfigFile_21), int32(0))
	mBase = m.M
	v917 = m.ExcPending
	if v917 != 0 {
		goto L6
	} else {
		goto L192
	}
L192:
	;
	v922 = v26
	goto L17
L193:
	;
	v936 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	v937 = *(*int32)(unsafe.Add(mBase, uint32(v936)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v922
	F_RunObjectPostAlterHookStr(m, v77, int32(_a_F_AlterSystemSetConfigFile_28), v937)
	mBase = m.M
	v943 = m.ExcPending
	if v943 != 0 {
		goto L6
	} else {
		goto L196
	}
L194:
	;
	goto L195
L195:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v922
	v949 = v23 + int32(496)
	v951 = F_BasicOpenFile(m, v949, int32(578))
	mBase = m.M
	v952 = m.ExcPending
	if v952 != 0 {
		goto L6
	} else {
		goto L197
	}
L196:
	;
	goto L195
L197:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2552)) = v951
	if v951 < int32(0) {
		goto L198
	} else {
		goto L199
	}
L198:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v922
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v963 = m.ExcPending
	if v963 != 0 {
		goto L6
	} else {
		goto L201
	}
L199:
	;
	goto L200
L200:
	;
	v988 = *(*int32)(unsafe.Add(mBase, _c_F_AlterSystemSetConfigFile[4]))
	v990 = *(*int32)(unsafe.Add(mBase, _c_F_AlterSystemSetConfigFile[5]))
	goto L205
L201:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v922
	F_errcode_for_file_access(m)
	mBase = m.M
	v968 = m.ExcPending
	if v968 != 0 {
		goto L6
	} else {
		goto L202
	}
L202:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v922
	*(*int32)(unsafe.Add(mBase, uint32(v23)+80)) = v949
	F_errmsg(m, int32(_a_F_AlterSystemSetConfigFile_23), v23+int32(80))
	mBase = m.M
	v977 = m.ExcPending
	if v977 != 0 {
		goto L6
	} else {
		goto L203
	}
L203:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v922
	F_errfinish(m, int32(_a_F_AlterSystemSetConfigFile_1), int32(_a_F_AlterSystemSetConfigFile_29), int32(_a_F_AlterSystemSetConfigFile_3))
	mBase = m.M
	v985 = m.ExcPending
	if v985 != 0 {
		goto L6
	} else {
		goto L204
	}
L204:
	;
	goto L3
L205:
	;
	v992 = v23 + int32(224)
	*(*int32)(unsafe.Add(mBase, uint32(v992)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v992))) = v23 + int32(220)
	goto L208
L206:
	;
	v1000 = v990
	v1001 = v988
	v1002 = v922
	v1007 = int32(0)
	goto L9
L208:
	;
	goto L206
L209:
	;
	v1471 = *(*int32)(unsafe.Add(mBase, _c_F_AlterSystemSetConfigFile[6]))
	if v1471 == int32(0) {
		goto L288
	} else {
		goto L289
	}
L210:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_AlterSystemSetConfigFile[5])) = v23 + int32(224)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v1001
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v1000
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v1002
	v1023 = *(*int32)(unsafe.Add(mBase, uint32(v23)+2552))
	v1024 = *(*int32)(unsafe.Add(mBase, uint32(v23)+2560))
	v1026 = v23 + int32(2564)
	F_initStringInfo(m, v1026)
	mBase = m.M
	v1028 = m.ExcPending
	if v1028 != 0 {
		goto L6
	} else {
		goto L213
	}
L211:
	;
	goto L212
L212:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_AlterSystemSetConfigFile[4])) = v1001
	*(*int32)(unsafe.Add(mBase, _c_F_AlterSystemSetConfigFile[5])) = v1000
	v1451 = *(*int32)(unsafe.Add(mBase, uint32(v23)+2552))
	if int32(0) <= v1451 {
		goto L284
	} else {
		goto L285
	}
L213:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v1000
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v1001
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v1002
	F_appendStringInfoString(m, v1026, int32(_a_F_AlterSystemSetConfigFile_30))
	mBase = m.M
	v1034 = m.ExcPending
	if v1034 != 0 {
		goto L6
	} else {
		goto L214
	}
L214:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v1000
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v1001
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v1002
	F_appendStringInfoString(m, v1026, int32(_a_F_AlterSystemSetConfigFile_31))
	mBase = m.M
	v1040 = m.ExcPending
	if v1040 != 0 {
		goto L6
	} else {
		goto L215
	}
L215:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v1000
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v1001
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v1002
	*(*int32)(unsafe.Add(mBase, _c_F_AlterSystemSetConfigFile[6])) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v1001
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v1000
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v1002
	v1050 = *(*int32)(unsafe.Add(mBase, uint32(v23)+2564))
	v1051 = *(*int32)(unsafe.Add(mBase, uint32(v23)+2568))
	v1052 = F_write(m, v1023, v1050, v1051)
	mBase = m.M
	v1053 = *(*int32)(unsafe.Add(mBase, uint32(v23)+2568))
	if v1052 != v1053 {
		goto L209
	} else {
		goto L216
	}
L216:
	;
	if v1024 != 0 {
		goto L217
	} else {
		goto L218
	}
L217:
	;
	v1064 = v1024
	goto L220
L218:
	;
	goto L219
L219:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v1000
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v1001
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v1002
	v1358 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_AlterSystemSetConfigFile[7])))
	if v1358 != int32(1) {
		v1372 = int32(0)
		goto L267
	} else {
		goto L268
	}
L220:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v1000
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v1001
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v1002
	v1075 = v23 + int32(2564)
	v1076 = *(*int32)(unsafe.Add(mBase, uint32(v1075)))
	v1077 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1076))) = uint8(v1077)
	*(*int32)(unsafe.Add(mBase, uint32(v1075)+12)) = v1077
	*(*int32)(unsafe.Add(mBase, uint32(v1075)+4)) = v1077
	goto L222
L221:
	;
	goto L219
L222:
	;
	v1083 = *(*int32)(unsafe.Add(mBase, uint32(v1064)))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v1000
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v1001
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v1002
	F_appendStringInfoString(m, v1075, v1083)
	mBase = m.M
	v1088 = m.ExcPending
	if v1088 != 0 {
		goto L6
	} else {
		goto L223
	}
L223:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v1000
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v1001
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v1002
	F_appendStringInfoString(m, v1075, int32(_a_F_AlterSystemSetConfigFile_32))
	mBase = m.M
	v1094 = m.ExcPending
	if v1094 != 0 {
		goto L6
	} else {
		goto L224
	}
L224:
	;
	v1095 = *(*int32)(unsafe.Add(mBase, uint32(v1064)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v1000
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v1001
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v1002
	v1099 = int32(0)
	v1102 = F_strlen(m, v1095)
	mBase = m.M
	v1103 = int32(1)
	v1107 = F_emscripten_builtin_malloc(m, v1102<<(uint(v1103)%32)|v1103)
	mBase = m.M
	if v1107 != 0 {
		goto L225
	} else {
		goto L226
	}
L225:
	;
	if v1102 <= int32(0) {
		v1214 = v1099
		goto L228
	} else {
		goto L229
	}
L226:
	;
	goto L227
L227:
	;
	if v1107 == int32(0) {
		goto L246
	} else {
		goto L247
	}
L228:
	;
	v1225 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1107+v1214))) = uint8(v1225)
	goto L227
L229:
	;
	if v1102 != int32(1) {
		goto L230
	} else {
		goto L231
	}
L230:
	;
	v1122 = v1099
	v1124 = v1099
	v1126 = v1099
	goto L233
L231:
	;
	v1180 = v1099
	v1182 = v1099
	goto L232
L232:
	;
	v1191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1182+v1095))))
	if base.B2i32(v1191 != int32(92))&base.B2i32(v1191 != int32(39)) == int32(0) {
		goto L243
	} else {
		goto L244
	}
L233:
	;
	v1132 = v1124 + v1095
	v1133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1132))))
	if base.B2i32(v1133 != int32(92))&base.B2i32(v1133 != int32(39)) == int32(0) {
		goto L235
	} else {
		goto L236
	}
L234:
	;
	if v1102&int32(1) == int32(0) {
		v1214 = v1168
		goto L228
	} else {
		goto L242
	}
L235:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1107+v1122))) = uint8(v1133)
	v1145 = v1122 + int32(1)
	goto L237
L236:
	;
	v1145 = v1122
	goto L237
L237:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1107+v1145))) = uint8(v1133)
	v1149 = v1145 + int32(1)
	v1150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1132)+1)))
	if base.B2i32(v1150 != int32(92))&base.B2i32(v1150 != int32(39)) == int32(0) {
		goto L238
	} else {
		goto L239
	}
L238:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1107+v1149))) = uint8(v1150)
	v1162 = v1145 + int32(2)
	goto L240
L239:
	;
	v1162 = v1149
	goto L240
L240:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1107+v1162))) = uint8(v1150)
	v1165 = int32(2)
	v1166 = v1124 + v1165
	v1168 = v1162 + int32(1)
	v1170 = v1126 + v1165
	if v1170 != v1102&int32(2147483646) {
		v1122 = v1168
		v1124 = v1166
		v1126 = v1170
		goto L233
	} else {
		goto L241
	}
L241:
	;
	goto L234
L242:
	;
	v1180 = v1168
	v1182 = v1166
	goto L232
L243:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1107+v1180))) = uint8(v1191)
	v1203 = v1180 + int32(1)
	goto L245
L244:
	;
	v1203 = v1180
	goto L245
L245:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1107+v1203))) = uint8(v1191)
	v1214 = v1203 + int32(1)
	goto L228
L246:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v1000
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v1001
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v1002
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1251 = m.ExcPending
	if v1251 != 0 {
		goto L6
	} else {
		goto L249
	}
L247:
	;
	goto L248
L248:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v1000
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v1001
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v1002
	v1277 = v23 + int32(2564)
	F_appendStringInfoString(m, v1277, v1107)
	mBase = m.M
	v1279 = m.ExcPending
	if v1279 != 0 {
		goto L6
	} else {
		goto L253
	}
L249:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v1000
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v1001
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v1002
	F_errcode(m, int32(_a_F_AlterSystemSetConfigFile_33))
	mBase = m.M
	v1257 = m.ExcPending
	if v1257 != 0 {
		goto L6
	} else {
		goto L250
	}
L250:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v1000
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v1001
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v1002
	F_errmsg(m, int32(_a_F_AlterSystemSetConfigFile_34), int32(0))
	mBase = m.M
	v1264 = m.ExcPending
	if v1264 != 0 {
		goto L6
	} else {
		goto L251
	}
L251:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v1000
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v1001
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v1002
	F_errfinish(m, int32(_a_F_AlterSystemSetConfigFile_1), int32(_a_F_AlterSystemSetConfigFile_35), int32(_a_F_AlterSystemSetConfigFile_36))
	mBase = m.M
	v1272 = m.ExcPending
	if v1272 != 0 {
		goto L6
	} else {
		goto L252
	}
L252:
	;
	goto L3
L253:
	;
	F_emscripten_builtin_free(m, v1107)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v1000
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v1001
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v1002
	F_appendStringInfoString(m, v1277, int32(_a_F_AlterSystemSetConfigFile_37))
	mBase = m.M
	v1286 = m.ExcPending
	if v1286 != 0 {
		goto L6
	} else {
		goto L254
	}
L254:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_AlterSystemSetConfigFile[6])) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v1001
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v1000
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v1002
	v1293 = *(*int32)(unsafe.Add(mBase, uint32(v23)+2564))
	v1294 = *(*int32)(unsafe.Add(mBase, uint32(v23)+2568))
	v1295 = F_write(m, v1023, v1293, v1294)
	mBase = m.M
	v1296 = *(*int32)(unsafe.Add(mBase, uint32(v23)+2568))
	if v1295 != v1296 {
		goto L255
	} else {
		goto L256
	}
L255:
	;
	v1299 = *(*int32)(unsafe.Add(mBase, _c_F_AlterSystemSetConfigFile[6]))
	if v1299 == int32(0) {
		goto L258
	} else {
		goto L259
	}
L256:
	;
	goto L257
L257:
	;
	v1336 = *(*int32)(unsafe.Add(mBase, uint32(v1064)+24))
	if v1336 != 0 {
		v1064 = v1336
		goto L220
	} else {
		goto L265
	}
L258:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_AlterSystemSetConfigFile[6])) = int32(51)
	goto L260
L259:
	;
	goto L260
L260:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v1000
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v1001
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v1002
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1311 = m.ExcPending
	if v1311 != 0 {
		goto L6
	} else {
		goto L261
	}
L261:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v1000
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v1001
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v1002
	F_errcode_for_file_access(m)
	mBase = m.M
	v1316 = m.ExcPending
	if v1316 != 0 {
		goto L6
	} else {
		goto L262
	}
L262:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v1000
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v1001
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v1002
	*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = v23 + int32(496)
	F_errmsg(m, int32(_a_F_AlterSystemSetConfigFile_38), v23+int32(16))
	mBase = m.M
	v1327 = m.ExcPending
	if v1327 != 0 {
		goto L6
	} else {
		goto L263
	}
L263:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v1000
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v1001
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v1002
	F_errfinish(m, int32(_a_F_AlterSystemSetConfigFile_1), int32(_a_F_AlterSystemSetConfigFile_39), int32(_a_F_AlterSystemSetConfigFile_36))
	mBase = m.M
	v1335 = m.ExcPending
	if v1335 != 0 {
		goto L6
	} else {
		goto L264
	}
L264:
	;
	goto L3
L265:
	;
	goto L221
L266:
	;
	if v1372 != 0 {
		goto L273
	} else {
		goto L274
	}
L267:
	;
	goto L266
L268:
	;
	goto L269
L269:
	;
	v1363 = F_fsync(m, v1023)
	mBase = m.M
	if v1363 != int32(-1) {
		v1372 = v1363
		goto L267
	} else {
		goto L271
	}
L270:
	;
	v1372 = int32(-1)
	goto L267
L271:
	;
	v1367 = *(*int32)(unsafe.Add(mBase, _c_F_AlterSystemSetConfigFile[6]))
	if v1367 == int32(27) {
		goto L269
	} else {
		goto L272
	}
L272:
	;
	goto L270
L273:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v1000
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v1001
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v1002
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1379 = m.ExcPending
	if v1379 != 0 {
		goto L6
	} else {
		goto L276
	}
L274:
	;
	goto L275
L275:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v1001
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v1000
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v1002
	v1405 = *(*int32)(unsafe.Add(mBase, uint32(v23)+2564))
	F_pfree(m, v1405)
	mBase = m.M
	v1407 = m.ExcPending
	if v1407 != 0 {
		goto L6
	} else {
		goto L280
	}
L276:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v1000
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v1001
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v1002
	F_errcode_for_file_access(m)
	mBase = m.M
	v1384 = m.ExcPending
	if v1384 != 0 {
		goto L6
	} else {
		goto L277
	}
L277:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v1000
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v1001
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v1002
	*(*int32)(unsafe.Add(mBase, uint32(v23))) = v23 + int32(496)
	F_errmsg(m, int32(_a_F_AlterSystemSetConfigFile_40), v23)
	mBase = m.M
	v1393 = m.ExcPending
	if v1393 != 0 {
		goto L6
	} else {
		goto L278
	}
L278:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v1000
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v1001
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v1002
	F_errfinish(m, int32(_a_F_AlterSystemSetConfigFile_1), int32(_a_F_AlterSystemSetConfigFile_41), int32(_a_F_AlterSystemSetConfigFile_36))
	mBase = m.M
	v1401 = m.ExcPending
	if v1401 != 0 {
		goto L6
	} else {
		goto L279
	}
L279:
	;
	goto L3
L280:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v1001
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v1000
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v1002
	v1411 = *(*int32)(unsafe.Add(mBase, uint32(v23)+2552))
	v1412 = F_close(m, v1411)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v1001
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2552)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v1000
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v1002
	v1423 = F_durable_rename(m, v23+int32(496), v23+int32(1520), int32(21))
	mBase = m.M
	v1424 = m.ExcPending
	if v1424 != 0 {
		goto L6
	} else {
		goto L281
	}
L281:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_AlterSystemSetConfigFile[4])) = v1001
	*(*int32)(unsafe.Add(mBase, _c_F_AlterSystemSetConfigFile[5])) = v1000
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v1001
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v1000
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v1002
	v1432 = *(*int32)(unsafe.Add(mBase, uint32(v23)+2560))
	F_FreeConfigVariables(m, v1432)
	mBase = m.M
	v1434 = m.ExcPending
	if v1434 != 0 {
		goto L6
	} else {
		goto L282
	}
L282:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v1001
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v1000
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v1002
	v1439 = *(*int32)(unsafe.Add(mBase, _c_F_AlterSystemSetConfigFile[3]))
	F_LWLockRelease(m, v1439+int32(_a_F_AlterSystemSetConfigFile_21))
	mBase = m.M
	v1443 = m.ExcPending
	if v1443 != 0 {
		goto L6
	} else {
		goto L283
	}
L283:
	;
	m.G0 = v23 + int32(2592)
	return
L284:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v1001
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v1000
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v1002
	v1457 = *(*int32)(unsafe.Add(mBase, uint32(v23)+2552))
	v1458 = F_close(m, v1457)
	mBase = m.M
	goto L286
L285:
	;
	goto L286
L286:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v1000
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v1001
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v1002
	v1464 = F_unlink(m, v23+int32(496))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v1000
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v1001
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v1002
	F_pg_re_throw(m)
	mBase = m.M
	v1469 = m.ExcPending
	if v1469 != 0 {
		goto L6
	} else {
		goto L287
	}
L287:
	;
	goto L3
L288:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_AlterSystemSetConfigFile[6])) = int32(51)
	goto L290
L289:
	;
	goto L290
L290:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v1000
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v1001
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v1002
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1483 = m.ExcPending
	if v1483 != 0 {
		goto L6
	} else {
		goto L291
	}
L291:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v1000
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v1001
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v1002
	F_errcode_for_file_access(m)
	mBase = m.M
	v1488 = m.ExcPending
	if v1488 != 0 {
		goto L6
	} else {
		goto L292
	}
L292:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v1000
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v1001
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v1002
	*(*int32)(unsafe.Add(mBase, uint32(v23)+32)) = v23 + int32(496)
	F_errmsg(m, int32(_a_F_AlterSystemSetConfigFile_38), v23+int32(32))
	mBase = m.M
	v1499 = m.ExcPending
	if v1499 != 0 {
		goto L6
	} else {
		goto L293
	}
L293:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v1000
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v1001
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v1002
	F_errfinish(m, int32(_a_F_AlterSystemSetConfigFile_1), int32(_a_F_AlterSystemSetConfigFile_42), int32(_a_F_AlterSystemSetConfigFile_36))
	mBase = m.M
	v1507 = m.ExcPending
	if v1507 != 0 {
		goto L6
	} else {
		goto L294
	}
L294:
	;
	goto L5
L295:
	;
	v1529 = int32(v1525)
	m.G0 = v23
	v1531 = *(*int32)(unsafe.Add(mBase, uint32(v1529)+4))
	v1532 = *(*int32)(unsafe.Add(mBase, uint32(v1529)))
	v1535 = *(*int32)(unsafe.Add(mBase, uint32(v1532)))
	if v23+int32(220) == v1535 {
		goto L298
	} else {
		goto L299
	}
L296:
	;
	m.ExcPending = 1
	goto L304
L297:
	;
	if v1539 != 0 {
		goto L301
	} else {
		goto L302
	}
L298:
	;
	v1537 = *(*int32)(unsafe.Add(mBase, uint32(v1532)+4))
	v1539 = v1537
	goto L300
L299:
	;
	v1539 = int32(0)
	goto L300
L300:
	;
	goto L297
L301:
	;
	v1540 = *(*int32)(unsafe.Add(mBase, uint32(v23)+2588))
	v1541 = *(*int32)(unsafe.Add(mBase, uint32(v23)+2584))
	v1542 = *(*int32)(unsafe.Add(mBase, uint32(v23)+2580))
	v24 = v1541
	v25 = v1542
	v26 = v1540
	v30 = v1539
	v31 = v1531
	goto L1
L302:
	;
	goto L303
L303:
	;
	F___wasm_longjmp(m, v1532, v1531)
	mBase = m.M
	v1544 = m.ExcPending
	if v1544 != 0 {
		goto L304
	} else {
		goto L305
	}
L304:
	;
	return
L305:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_ExecAlterExtensionContentsRecurse(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v79 int32
	_ = v79
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v231 int32
	_ = v231
	var v237 int32
	_ = v237
	var v242 int32
	_ = v242
	var v246 int32
	_ = v246
	var v252 int32
	_ = v252
	var v257 int32
	_ = v257
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v268 int32
	_ = v268
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v308 int32
	_ = v308
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v319 int32
	_ = v319
	var v322 int32
	_ = v322
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v338 int32
	_ = v338
	var v351 int32
	_ = v351
	var v354 int32
	_ = v354
	var v356 int64
	_ = v356
	var v370 int32
	_ = v370
	var v372 int32
	_ = v372
	var v380 int32
	_ = v380
	var v382 int32
	_ = v382
	var v386 int32
	_ = v386
	var v392 int32
	_ = v392
	var v395 int32
	_ = v395
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v410 int32
	_ = v410
	var v413 int32
	_ = v413
	var v420 int32
	_ = v420
	var v439 int32
	_ = v439
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v455 int32
	_ = v455
	var v457 int32
	_ = v457
	var v460 int32
	_ = v460
	var v462 int32
	_ = v462
	var v464 int32
	_ = v464
	var v467 int32
	_ = v467
	var v469 int32
	_ = v469
	var v471 int32
	_ = v471
	var v474 int32
	_ = v474
	var v478 int32
	_ = v478
	var v495 int32
	_ = v495
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v515 int32
	_ = v515
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v528 int32
	_ = v528
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v535 int32
	_ = v535
	var v539 int32
	_ = v539
	var v541 int32
	_ = v541
	var v549 int32
	_ = v549
	var v551 int32
	_ = v551
	var v558 int32
	_ = v558
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v582 int32
	_ = v582
	var v585 int32
	_ = v585
	var v591 int32
	_ = v591
	var v607 int32
	_ = v607
	var v617 int32
	_ = v617
	var v618 int32
	_ = v618
	var v619 int32
	_ = v619
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v624 int32
	_ = v624
	var v626 int32
	_ = v626
	var v629 int32
	_ = v629
	var v631 int32
	_ = v631
	var v633 int32
	_ = v633
	var v636 int32
	_ = v636
	var v638 int32
	_ = v638
	var v640 int32
	_ = v640
	var v643 int32
	_ = v643
	var v647 int32
	_ = v647
	var v664 int32
	_ = v664
	var v666 int32
	_ = v666
	var v667 int32
	_ = v667
	var v683 int32
	_ = v683
	var v690 int32
	_ = v690
	var v691 int32
	_ = v691
	var v695 int32
	_ = v695
	var v711 int32
	_ = v711
	var v714 int32
	_ = v714
	var v715 int32
	_ = v715
	var v721 int32
	_ = v721
	var v730 int32
	_ = v730
	var v731 int32
	_ = v731
	var v733 int32
	_ = v733
	var v738 int32
	_ = v738
	var v739 int32
	_ = v739
	var v742 int32
	_ = v742
	var v743 int32
	_ = v743
	var v744 int32
	_ = v744
	var v745 int32
	_ = v745
	var v749 int32
	_ = v749
	var v756 int32
	_ = v756
	var v767 int32
	_ = v767
	var v768 int32
	_ = v768
	var v772 int32
	_ = v772
	var v774 int32
	_ = v774
	var v777 int32
	_ = v777
	var v794 int32
	_ = v794
	var v809 int32
	_ = v809
	var v812 int32
	_ = v812
	var v816 int32
	_ = v816
	var v820 int32
	_ = v820
	var v825 int32
	_ = v825
	var v827 int32
	_ = v827
	var v859 int32
	_ = v859
	var v866 int32
	_ = v866
	var v867 int32
	_ = v867
	var v868 int32
	_ = v868
	var v870 int32
	_ = v870
	var v872 int64
	_ = v872
	var v874 int64
	_ = v874
	var v876 int32
	_ = v876
	var v883 int32
	_ = v883
	var v884 int32
	_ = v884
	var v885 int32
	_ = v885
	var v886 int32
	_ = v886
	var v887 int32
	_ = v887
	var v888 int32
	_ = v888
	var v889 int32
	_ = v889
	var v893 int32
	_ = v893
	var v895 int64
	_ = v895
	var v897 int64
	_ = v897
	var v899 int32
	_ = v899
	var v906 int32
	_ = v906
	var v908 int32
	_ = v908
	var v910 int32
	_ = v910
	var v917 int32
	_ = v917
	var v918 int32
	_ = v918
	var v919 int32
	_ = v919
	var v923 int32
	_ = v923
	var v925 int64
	_ = v925
	var v927 int64
	_ = v927
	var v929 int32
	_ = v929
	var v936 int32
	_ = v936
	var v944 int32
	_ = v944
	var v947 int32
	_ = v947
	var v949 int32
	_ = v949
	var v950 int32
	_ = v950
	var v951 int32
	_ = v951
	var v952 int32
	_ = v952
	var v959 int32
	_ = v959
	var v964 int32
	_ = v964
	var v968 int32
	_ = v968
	var v971 int32
	_ = v971
	var v972 int32
	_ = v972
	var v973 int32
	_ = v973
	var v974 int32
	_ = v974
	var v975 int32
	_ = v975
	var v980 int32
	_ = v980
	var v985 int32
	_ = v985
	var v989 int32
	_ = v989
	var v992 int32
	_ = v992
	var v994 int32
	_ = v994
	var v995 int32
	_ = v995
	var v996 int32
	_ = v996
	var v1003 int32
	_ = v1003
	var v1008 int32
	_ = v1008
	var v1012 int32
	_ = v1012
	var v1016 int32
	_ = v1016
	var v1021 int32
	_ = v1021
	var v1025 int32
	_ = v1025
	var v1031 int32
	_ = v1031
	var v1036 int32
	_ = v1036
	var v1041 int32
	_ = v1041
	var v1045 int32
	_ = v1045
	var v1050 int32
	_ = v1050
	var v1054 int32
	_ = v1054
	var v1058 int32
	_ = v1058
	var v1063 int32
	_ = v1063
	var v1067 int32
	_ = v1067
	var v1071 int32
	_ = v1071
	var v1076 int32
	_ = v1076
	var v1080 int32
	_ = v1080
	var v1084 int32
	_ = v1084
	var v1089 int32
	_ = v1089
	var v1093 int32
	_ = v1093
	var v1096 int32
	_ = v1096
	var v1097 int32
	_ = v1097
	var v1098 int32
	_ = v1098
	var v1099 int32
	_ = v1099
	var v1105 int32
	_ = v1105
	var v1110 int32
	_ = v1110
	v15 = m.G0
	v17 = v15 - int32(288)
	m.G0 = v17
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v21 = F_getExtensionOfObject(m, v19, v20)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if int32(0) < v23 {
		goto L14
	} else {
		goto L15
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1093 = m.ExcPending
	if v1093 != 0 {
		goto L1
	} else {
		goto L257
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1080 = m.ExcPending
	if v1080 != 0 {
		goto L1
	} else {
		goto L254
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1067 = m.ExcPending
	if v1067 != 0 {
		goto L1
	} else {
		goto L251
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1054 = m.ExcPending
	if v1054 != 0 {
		goto L1
	} else {
		goto L248
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1041 = m.ExcPending
	if v1041 != 0 {
		goto L1
	} else {
		goto L245
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1025 = m.ExcPending
	if v1025 != 0 {
		goto L1
	} else {
		goto L242
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1012 = m.ExcPending
	if v1012 != 0 {
		goto L1
	} else {
		goto L239
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v989 = m.ExcPending
	if v989 != 0 {
		goto L1
	} else {
		goto L234
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v968 = m.ExcPending
	if v968 != 0 {
		goto L1
	} else {
		goto L229
	}
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v944 = m.ExcPending
	if v944 != 0 {
		goto L1
	} else {
		goto L223
	}
L13:
	;
	v859 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v859 == int32(1247) {
		goto L204
	} else {
		goto L205
	}
L14:
	;
	if v21 != 0 {
		goto L12
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v274 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v21 != v274 {
		goto L10
	} else {
		goto L102
	}
L17:
	;
	if v19 == int32(2615) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v30 = F_SearchSysCache1(m, int32(28), v29)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L22
	}
L19:
	;
	goto L20
L20:
	;
	F_recordDependencyOn(m, l2, l1, int32(101))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L1
	} else {
		goto L28
	}
L21:
	;
	if v41 == v20 {
		goto L11
	} else {
		goto L27
	}
L22:
	;
	if v30 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v41 = int32(0)
	goto L21
L24:
	;
	goto L25
L25:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v30)+16))
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+22)))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v35+v36)+72))
	F_ReleaseCatCache(m, v30)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	v41 = v38
	goto L21
L27:
	;
	goto L20
L28:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v50 = m.G0
	v52 = v50 - int32(112)
	m.G0 = v52
	if v49 != int32(2613) {
		goto L35
	} else {
		goto L36
	}
L29:
	;
	goto L13
L30:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L1
	} else {
		goto L98
	}
L31:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L1
	} else {
		goto L95
	}
L32:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L1
	} else {
		goto L92
	}
L33:
	;
	m.G0 = v52 + int32(112)
	goto L29
L34:
	;
	v185 = F_get_object_attnum_acl(m, v49)
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L1
	} else {
		goto L79
	}
L35:
	;
	if v49 != int32(1259) {
		goto L34
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	v149 = F_table_open(m, int32(2995), int32(3))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L1
	} else {
		goto L67
	}
L38:
	;
	v59 = F_SearchSysCache1(m, int32(57), v48)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	if v59 == int32(0) {
		goto L32
	} else {
		goto L40
	}
L40:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v59)+16))
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63)+22)))
	v65 = v63 + v64
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65)+119)))
	switch v66 - int32(73) {
	case 0, 26, 32:
		goto L43
	default:
		goto L42
	case 10:
		goto L41
	}
L41:
	;
	v134 = F_SysCacheGetAttr(m, int32(57), v59, int32(32), v52+int32(48))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L1
	} else {
		goto L60
	}
L42:
	;
	v72 = int32(*(*int16)(unsafe.Add(mBase, uint32(v65)+120)))
	if v72 <= int32(0) {
		goto L41
	} else {
		goto L45
	}
L43:
	;
	F_ReleaseCatCache(m, v59)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	goto L33
L45:
	;
	v79 = int32(1)
	goto L46
L46:
	;
	v90 = F_SearchSysCache2(m, int32(7), v48, v79)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L1
	} else {
		goto L48
	}
L47:
	;
	goto L41
L48:
	;
	if v90 != 0 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v90)+16))
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92)+22)))
	v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92+v93)+91)))
	if v95 != 0 {
		goto L52
	} else {
		goto L53
	}
L50:
	;
	goto L51
L51:
	;
	v114 = base.I32_extend16_s(v79 + int32(1))
	if v114 <= v72 {
		v79 = v114
		goto L46
	} else {
		goto L59
	}
L52:
	;
	F_ReleaseCatCache(m, v90)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L1
	} else {
		goto L58
	}
L53:
	;
	v100 = F_SysCacheGetAttr(m, int32(7), v90, int32(22), v52+int32(48))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+48)))
	if v102 != 0 {
		goto L52
	} else {
		goto L55
	}
L55:
	;
	v104 = F_pg_detoast_datum(m, v100)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	F_recordExtensionInitPrivWorker(m, v48, int32(1259), v79, v104)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	goto L52
L58:
	;
	goto L51
L59:
	;
	goto L47
L60:
	;
	v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+48)))
	if v136 == int32(0) {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v141 = F_pg_detoast_datum(m, v134)
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L1
	} else {
		goto L64
	}
L62:
	;
	goto L63
L63:
	;
	F_ReleaseCatCache(m, v59)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L1
	} else {
		goto L66
	}
L64:
	;
	F_recordExtensionInitPrivWorker(m, v48, int32(1259), int32(0), v141)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	goto L63
L66:
	;
	goto L33
L67:
	;
	v152 = v52 + int32(48)
	F_ScanKeyInit(m, v152, int32(1), int32(3), int32(184), v48)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L1
	} else {
		goto L68
	}
L68:
	;
	v159 = int32(1)
	v162 = F_systable_beginscan(m, v149, int32(2996), v159, int32(0), v159, v152)
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L1
	} else {
		goto L69
	}
L69:
	;
	v164 = F_systable_getnext(m, v162)
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L1
	} else {
		goto L70
	}
L70:
	;
	if v164 == int32(0) {
		goto L31
	} else {
		goto L71
	}
L71:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v149)+52))
	v172 = F_heap_getattr_2(m, v164, int32(3), v169, v52+int32(111))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+111)))
	if v174 == int32(0) {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v179 = F_pg_detoast_datum(m, v172)
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L1
	} else {
		goto L76
	}
L74:
	;
	goto L75
L75:
	;
	F_systable_endscan(m, v162)
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L1
	} else {
		goto L78
	}
L76:
	;
	F_recordExtensionInitPrivWorker(m, v48, int32(2613), int32(0), v179)
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L1
	} else {
		goto L77
	}
L77:
	;
	goto L75
L78:
	;
	goto L33
L79:
	;
	if v185 == int32(0) {
		goto L33
	} else {
		goto L80
	}
L80:
	;
	v189 = F_get_object_catcache_oid(m, v49)
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	v191 = F_SearchSysCache1(m, v189, v48)
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L1
	} else {
		goto L82
	}
L82:
	;
	if v191 == int32(0) {
		goto L30
	} else {
		goto L83
	}
L83:
	;
	v195 = F_get_object_attnum_acl(m, v49)
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L1
	} else {
		goto L84
	}
L84:
	;
	v199 = F_SysCacheGetAttr(m, v189, v191, v195, v52+int32(48))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+48)))
	if v201 == int32(0) {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v205 = F_pg_detoast_datum(m, v199)
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L1
	} else {
		goto L89
	}
L87:
	;
	goto L88
L88:
	;
	F_ReleaseCatCache(m, v191)
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L1
	} else {
		goto L91
	}
L89:
	;
	F_recordExtensionInitPrivWorker(m, v48, v49, int32(0), v205)
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L1
	} else {
		goto L90
	}
L90:
	;
	goto L88
L91:
	;
	goto L33
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+16)) = v48
	F_errmsg_internal(m, int32(_a_F_ExecAlterExtensionContentsRecurse_0), v52+int32(16))
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L1
	} else {
		goto L93
	}
L93:
	;
	F_errfinish(m, int32(_a_F_ExecAlterExtensionContentsRecurse_1), int32(_a_F_ExecAlterExtensionContentsRecurse_2), int32(_a_F_ExecAlterExtensionContentsRecurse_3))
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L1
	} else {
		goto L94
	}
L94:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+32)) = v48
	F_errmsg_internal(m, int32(_a_F_ExecAlterExtensionContentsRecurse_4), v52+int32(32))
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L1
	} else {
		goto L96
	}
L96:
	;
	F_errfinish(m, int32(_a_F_ExecAlterExtensionContentsRecurse_1), int32(_a_F_ExecAlterExtensionContentsRecurse_5), int32(_a_F_ExecAlterExtensionContentsRecurse_3))
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
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
	v262 = F_get_object_class_descr(m, v49)
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L1
	} else {
		goto L99
	}
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+4)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v52))) = v262
	F_errmsg_internal(m, int32(_a_F_ExecAlterExtensionContentsRecurse_6), v52)
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L1
	} else {
		goto L100
	}
L100:
	;
	F_errfinish(m, int32(_a_F_ExecAlterExtensionContentsRecurse_1), int32(_a_F_ExecAlterExtensionContentsRecurse_7), int32(_a_F_ExecAlterExtensionContentsRecurse_3))
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
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
	v278 = F_deleteDependencyRecordsForClass(m, v19, v20, int32(3079), int32(101))
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L1
	} else {
		goto L103
	}
L103:
	;
	if v278 != int32(1) {
		goto L9
	} else {
		goto L104
	}
L104:
	;
	v282 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v282 == int32(1259) {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	v285 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v288 = F_table_open(m, int32(3079), int32(3))
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L1
	} else {
		goto L108
	}
L106:
	;
	v721 = v282
	goto L107
L107:
	;
	v730 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v731 = m.G0
	v733 = v731 - int32(16)
	m.G0 = v733
	if v721 == int32(1259) {
		goto L180
	} else {
		goto L181
	}
L108:
	;
	v291 = v17 + int32(240)
	F_ScanKeyInit(m, v291, int32(1), int32(3), int32(184), v21)
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L1
	} else {
		goto L109
	}
L109:
	;
	v298 = int32(1)
	v301 = F_systable_beginscan(m, v288, int32(3080), v298, int32(0), v298, v291)
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L1
	} else {
		goto L110
	}
L110:
	;
	v303 = F_systable_getnext(m, v301)
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L1
	} else {
		goto L111
	}
L111:
	;
	if v303 == int32(0) {
		goto L8
	} else {
		goto L112
	}
L112:
	;
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v288)+52))
	v311 = F_heap_getattr_7(m, v303, int32(7), v308, v17+int32(239))
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L1
	} else {
		goto L113
	}
L113:
	;
	v313 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+239)))
	if v313 != 0 {
		goto L114
	} else {
		goto L115
	}
L114:
	;
	F_systable_endscan(m, v301)
	mBase = m.M
	v711 = m.ExcPending
	if v711 != 0 {
		goto L1
	} else {
		goto L175
	}
L115:
	;
	v314 = F_pg_detoast_datum(m, v311)
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L1
	} else {
		goto L116
	}
L116:
	;
	v316 = *(*int32)(unsafe.Add(mBase, uint32(v314)+4))
	if v316 != int32(1) {
		goto L7
	} else {
		goto L117
	}
L117:
	;
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v314)+20))
	if v319 != int32(1) {
		goto L7
	} else {
		goto L118
	}
L118:
	;
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v314)+16))
	if v322 < int32(0) {
		goto L7
	} else {
		goto L119
	}
L119:
	;
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v314)+8))
	if v325 != 0 {
		goto L7
	} else {
		goto L120
	}
L120:
	;
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v314)+12))
	if v326 != int32(26) {
		goto L7
	} else {
		goto L121
	}
L121:
	;
	if v322 == int32(0) {
		goto L114
	} else {
		goto L122
	}
L122:
	;
	v338 = int32(0)
	goto L123
L123:
	;
	v351 = *(*int32)(unsafe.Add(mBase, uint32(v314+int32(24)+v338<<(uint(int32(2))%32))))
	if v285 != v351 {
		goto L125
	} else {
		goto L126
	}
L124:
	;
	v356 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v17)+216)) = v356
	*(*int64)(unsafe.Add(mBase, uint32(v17)+208)) = v356
	*(*int64)(unsafe.Add(mBase, uint32(v17)+200)) = v356
	*(*int64)(unsafe.Add(mBase, uint32(v17)+192)) = v356
	*(*int64)(unsafe.Add(mBase, uint32(v17)+184)) = v356
	*(*int64)(unsafe.Add(mBase, uint32(v17)+176)) = int64(72339069014638592)
	if v322 == int32(1) {
		goto L130
	} else {
		goto L131
	}
L125:
	;
	v354 = v338 + int32(1)
	if v354 != v322 {
		v338 = v354
		goto L123
	} else {
		goto L128
	}
L126:
	;
	goto L127
L127:
	;
	goto L124
L128:
	;
	goto L114
L129:
	;
	v515 = *(*int32)(unsafe.Add(mBase, uint32(v288)+52))
	v518 = F_heap_getattr_7(m, v303, int32(8), v515, v17+int32(239))
	mBase = m.M
	v519 = m.ExcPending
	if v519 != 0 {
		goto L1
	} else {
		goto L147
	}
L130:
	;
	v370 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+190)) = uint8(v370)
	goto L129
L131:
	;
	goto L132
L132:
	;
	v372 = int32(0)
	F_deconstruct_array_builtin(m, v314, int32(26), v17+int32(172), v372, v17+int32(168))
	mBase = m.M
	v380 = m.ExcPending
	if v380 != 0 {
		goto L1
	} else {
		goto L133
	}
L133:
	;
	v382 = v322 - int32(1)
	if v382 <= v338 {
		goto L134
	} else {
		goto L135
	}
L134:
	;
	v495 = *(*int32)(unsafe.Add(mBase, uint32(v17)+172))
	v497 = F_construct_array_builtin(m, v495, v382, int32(26))
	mBase = m.M
	v498 = m.ExcPending
	if v498 != 0 {
		goto L1
	} else {
		goto L146
	}
L135:
	;
	v386 = (v382 - v338) & int32(3)
	if v386 != 0 {
		goto L136
	} else {
		goto L137
	}
L136:
	;
	v392 = v338
	v395 = v372
	goto L139
L137:
	;
	v420 = v338
	goto L138
L138:
	;
	if base.Ui32(v322-v338-int32(2)) < base.Ui32(int32(3)) {
		goto L134
	} else {
		goto L142
	}
L139:
	;
	v401 = *(*int32)(unsafe.Add(mBase, uint32(v17)+172))
	v402 = int32(2)
	v405 = int32(1)
	v406 = v392 + v405
	v410 = *(*int32)(unsafe.Add(mBase, uint32(v401+v406<<(uint(v402)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v401+v392<<(uint(v402)%32)))) = v410
	v413 = v395 + v405
	if v413 != v386 {
		v392 = v406
		v395 = v413
		goto L139
	} else {
		goto L141
	}
L140:
	;
	v420 = v406
	goto L138
L141:
	;
	goto L140
L142:
	;
	v439 = v420
	goto L143
L143:
	;
	v448 = *(*int32)(unsafe.Add(mBase, uint32(v17)+172))
	v449 = int32(2)
	v450 = v439 << (uint(v449) % 32)
	v452 = int32(4)
	v453 = v450 + v452
	v455 = *(*int32)(unsafe.Add(mBase, uint32(v448+v453)))
	*(*int32)(unsafe.Add(mBase, uint32(v448+v450))) = v455
	v457 = *(*int32)(unsafe.Add(mBase, uint32(v17)+172))
	v460 = v450 + int32(8)
	v462 = *(*int32)(unsafe.Add(mBase, uint32(v457+v460)))
	*(*int32)(unsafe.Add(mBase, uint32(v457+v453))) = v462
	v464 = *(*int32)(unsafe.Add(mBase, uint32(v17)+172))
	v467 = v450 + int32(12)
	v469 = *(*int32)(unsafe.Add(mBase, uint32(v464+v467)))
	*(*int32)(unsafe.Add(mBase, uint32(v464+v460))) = v469
	v471 = *(*int32)(unsafe.Add(mBase, uint32(v17)+172))
	v474 = v439 + v452
	v478 = *(*int32)(unsafe.Add(mBase, uint32(v471+v474<<(uint(v449)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v467+v471))) = v478
	if v474 != v382 {
		v439 = v474
		goto L143
	} else {
		goto L145
	}
L144:
	;
	goto L134
L145:
	;
	goto L144
L146:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+216)) = v497
	goto L129
L147:
	;
	v520 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+239)))
	if v520 == int32(1) {
		goto L6
	} else {
		goto L148
	}
L148:
	;
	v523 = F_pg_detoast_datum(m, v518)
	mBase = m.M
	v524 = m.ExcPending
	if v524 != 0 {
		goto L1
	} else {
		goto L149
	}
L149:
	;
	v525 = *(*int32)(unsafe.Add(mBase, uint32(v523)+4))
	if v525 != int32(1) {
		goto L5
	} else {
		goto L150
	}
L150:
	;
	v528 = *(*int32)(unsafe.Add(mBase, uint32(v523)+20))
	if v528 != int32(1) {
		goto L5
	} else {
		goto L151
	}
L151:
	;
	v531 = *(*int32)(unsafe.Add(mBase, uint32(v523)+8))
	if v531 != 0 {
		goto L5
	} else {
		goto L152
	}
L152:
	;
	v532 = *(*int32)(unsafe.Add(mBase, uint32(v523)+12))
	if v532 != int32(25) {
		goto L5
	} else {
		goto L153
	}
L153:
	;
	v535 = *(*int32)(unsafe.Add(mBase, uint32(v523)+16))
	if v535 != v322 {
		goto L4
	} else {
		goto L154
	}
L154:
	;
	if v322 == int32(1) {
		goto L156
	} else {
		goto L157
	}
L155:
	;
	v683 = *(*int32)(unsafe.Add(mBase, uint32(v288)+52))
	v690 = F_heap_modify_tuple(m, v303, v683, v17+int32(192), v17+int32(184), v17+int32(176))
	mBase = m.M
	v691 = m.ExcPending
	if v691 != 0 {
		goto L1
	} else {
		goto L173
	}
L156:
	;
	v539 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+191)) = uint8(v539)
	goto L155
L157:
	;
	goto L158
L158:
	;
	v541 = int32(0)
	F_deconstruct_array_builtin(m, v523, int32(25), v17+int32(172), v541, v17+int32(168))
	mBase = m.M
	v549 = m.ExcPending
	if v549 != 0 {
		goto L1
	} else {
		goto L159
	}
L159:
	;
	v551 = v322 - int32(1)
	if v551 <= v338 {
		goto L160
	} else {
		goto L161
	}
L160:
	;
	v664 = *(*int32)(unsafe.Add(mBase, uint32(v17)+172))
	v666 = F_construct_array_builtin(m, v664, v551, int32(25))
	mBase = m.M
	v667 = m.ExcPending
	if v667 != 0 {
		goto L1
	} else {
		goto L172
	}
L161:
	;
	v558 = (v551 - v338) & int32(3)
	if v558 != 0 {
		goto L162
	} else {
		goto L163
	}
L162:
	;
	v563 = v338
	v564 = v541
	goto L165
L163:
	;
	v591 = v338
	goto L164
L164:
	;
	if base.Ui32(v322-v338-int32(2)) < base.Ui32(int32(3)) {
		goto L160
	} else {
		goto L168
	}
L165:
	;
	v573 = *(*int32)(unsafe.Add(mBase, uint32(v17)+172))
	v574 = int32(2)
	v577 = int32(1)
	v578 = v563 + v577
	v582 = *(*int32)(unsafe.Add(mBase, uint32(v573+v578<<(uint(v574)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v573+v563<<(uint(v574)%32)))) = v582
	v585 = v564 + v577
	if v585 != v558 {
		v563 = v578
		v564 = v585
		goto L165
	} else {
		goto L167
	}
L166:
	;
	v591 = v578
	goto L164
L167:
	;
	goto L166
L168:
	;
	v607 = v591
	goto L169
L169:
	;
	v617 = *(*int32)(unsafe.Add(mBase, uint32(v17)+172))
	v618 = int32(2)
	v619 = v607 << (uint(v618) % 32)
	v621 = int32(4)
	v622 = v619 + v621
	v624 = *(*int32)(unsafe.Add(mBase, uint32(v617+v622)))
	*(*int32)(unsafe.Add(mBase, uint32(v617+v619))) = v624
	v626 = *(*int32)(unsafe.Add(mBase, uint32(v17)+172))
	v629 = v619 + int32(8)
	v631 = *(*int32)(unsafe.Add(mBase, uint32(v626+v629)))
	*(*int32)(unsafe.Add(mBase, uint32(v626+v622))) = v631
	v633 = *(*int32)(unsafe.Add(mBase, uint32(v17)+172))
	v636 = v619 + int32(12)
	v638 = *(*int32)(unsafe.Add(mBase, uint32(v633+v636)))
	*(*int32)(unsafe.Add(mBase, uint32(v633+v629))) = v638
	v640 = *(*int32)(unsafe.Add(mBase, uint32(v17)+172))
	v643 = v607 + v621
	v647 = *(*int32)(unsafe.Add(mBase, uint32(v640+v643<<(uint(v618)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v636+v640))) = v647
	if v643 != v551 {
		v607 = v643
		goto L169
	} else {
		goto L171
	}
L170:
	;
	goto L160
L171:
	;
	goto L170
L172:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+220)) = v666
	goto L155
L173:
	;
	F_CatalogTupleUpdate(m, v288, v690+int32(4), v690)
	mBase = m.M
	v695 = m.ExcPending
	if v695 != 0 {
		goto L1
	} else {
		goto L174
	}
L174:
	;
	goto L114
L175:
	;
	F_relation_close(m, v288, int32(3))
	mBase = m.M
	v714 = m.ExcPending
	if v714 != 0 {
		goto L1
	} else {
		goto L176
	}
L176:
	;
	v715 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v721 = v715
	goto L107
L177:
	;
	m.G0 = v733 + int32(16)
	goto L13
L178:
	;
	F_ReleaseCatCache(m, v738)
	mBase = m.M
	v827 = m.ExcPending
	if v827 != 0 {
		goto L1
	} else {
		goto L202
	}
L179:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v816 = m.ExcPending
	if v816 != 0 {
		goto L1
	} else {
		goto L199
	}
L180:
	;
	v738 = F_SearchSysCache1(m, int32(57), v730)
	mBase = m.M
	v739 = m.ExcPending
	if v739 != 0 {
		goto L1
	} else {
		goto L183
	}
L181:
	;
	goto L182
L182:
	;
	v809 = int32(0)
	F_recordExtensionInitPrivWorker(m, v730, v721, v809, v809)
	mBase = m.M
	v812 = m.ExcPending
	if v812 != 0 {
		goto L1
	} else {
		goto L198
	}
L183:
	;
	if v738 == int32(0) {
		goto L179
	} else {
		goto L184
	}
L184:
	;
	v742 = *(*int32)(unsafe.Add(mBase, uint32(v738)+16))
	v743 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v742)+22)))
	v744 = v742 + v743
	v745 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v744)+119)))
	switch v745 - int32(73) {
	case 0, 26, 32:
		goto L178
	default:
		goto L186
	case 10:
		goto L185
	}
L185:
	;
	F_ReleaseCatCache(m, v738)
	mBase = m.M
	v794 = m.ExcPending
	if v794 != 0 {
		goto L1
	} else {
		goto L197
	}
L186:
	;
	v749 = int32(*(*int16)(unsafe.Add(mBase, uint32(v744)+120)))
	if v749 <= int32(0) {
		goto L185
	} else {
		goto L187
	}
L187:
	;
	v756 = int32(1)
	goto L188
L188:
	;
	v767 = F_SearchSysCache2(m, int32(7), v730, v756)
	mBase = m.M
	v768 = m.ExcPending
	if v768 != 0 {
		goto L1
	} else {
		goto L190
	}
L189:
	;
	goto L185
L190:
	;
	if v767 != 0 {
		goto L191
	} else {
		goto L192
	}
L191:
	;
	F_recordExtensionInitPrivWorker(m, v730, int32(1259), v756, int32(0))
	mBase = m.M
	v772 = m.ExcPending
	if v772 != 0 {
		goto L1
	} else {
		goto L194
	}
L192:
	;
	goto L193
L193:
	;
	v777 = base.I32_extend16_s(v756 + int32(1))
	if v777 <= v749 {
		v756 = v777
		goto L188
	} else {
		goto L196
	}
L194:
	;
	F_ReleaseCatCache(m, v767)
	mBase = m.M
	v774 = m.ExcPending
	if v774 != 0 {
		goto L1
	} else {
		goto L195
	}
L195:
	;
	goto L193
L196:
	;
	goto L189
L197:
	;
	goto L182
L198:
	;
	goto L177
L199:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v733))) = v730
	F_errmsg_internal(m, int32(_a_F_ExecAlterExtensionContentsRecurse_0), v733)
	mBase = m.M
	v820 = m.ExcPending
	if v820 != 0 {
		goto L1
	} else {
		goto L200
	}
L200:
	;
	F_errfinish(m, int32(_a_F_ExecAlterExtensionContentsRecurse_1), int32(_a_F_ExecAlterExtensionContentsRecurse_8), int32(_a_F_ExecAlterExtensionContentsRecurse_9))
	mBase = m.M
	v825 = m.ExcPending
	if v825 != 0 {
		goto L1
	} else {
		goto L201
	}
L201:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L202:
	;
	goto L177
L203:
	;
	m.G0 = v17 + int32(288)
	return
L204:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+248)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+240)) = int32(1247)
	v866 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v867 = F_get_array_type(m, v866)
	mBase = m.M
	v868 = m.ExcPending
	if v868 != 0 {
		goto L1
	} else {
		goto L207
	}
L205:
	;
	v910 = v859
	goto L206
L206:
	;
	if v910 != int32(1259) {
		goto L203
	} else {
		goto L219
	}
L207:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+244)) = v867
	if v867 != 0 {
		goto L208
	} else {
		goto L209
	}
L208:
	;
	v870 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+120)) = v870
	v872 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+112)) = v872
	v874 = *(*int64)(unsafe.Add(mBase, uint32(v17)+240))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+96)) = v874
	v876 = *(*int32)(unsafe.Add(mBase, uint32(v17)+248))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+104)) = v876
	F_ExecAlterExtensionContentsRecurse(m, l0, v17+int32(112), v17+int32(96))
	mBase = m.M
	v883 = m.ExcPending
	if v883 != 0 {
		goto L1
	} else {
		goto L211
	}
L209:
	;
	goto L210
L210:
	;
	v884 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v885 = F_type_is_range(m, v884)
	mBase = m.M
	v886 = m.ExcPending
	if v886 != 0 {
		goto L1
	} else {
		goto L212
	}
L211:
	;
	goto L210
L212:
	;
	if v885 != 0 {
		goto L213
	} else {
		goto L214
	}
L213:
	;
	v887 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v888 = F_get_range_multirange(m, v887)
	mBase = m.M
	v889 = m.ExcPending
	if v889 != 0 {
		goto L1
	} else {
		goto L216
	}
L214:
	;
	goto L215
L215:
	;
	v908 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v910 = v908
	goto L206
L216:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+244)) = v888
	if v888 == int32(0) {
		goto L3
	} else {
		goto L217
	}
L217:
	;
	v893 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+88)) = v893
	v895 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+80)) = v895
	v897 = *(*int64)(unsafe.Add(mBase, uint32(v17)+240))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+64)) = v897
	v899 = *(*int32)(unsafe.Add(mBase, uint32(v17)+248))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+72)) = v899
	F_ExecAlterExtensionContentsRecurse(m, l0, v17+int32(80), v17-int32(-64))
	mBase = m.M
	v906 = m.ExcPending
	if v906 != 0 {
		goto L1
	} else {
		goto L218
	}
L218:
	;
	goto L215
L219:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+248)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+240)) = int32(1247)
	v917 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v918 = F_get_rel_type_id(m, v917)
	mBase = m.M
	v919 = m.ExcPending
	if v919 != 0 {
		goto L1
	} else {
		goto L220
	}
L220:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+244)) = v918
	if v918 == int32(0) {
		goto L203
	} else {
		goto L221
	}
L221:
	;
	v923 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+56)) = v923
	v925 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+48)) = v925
	v927 = *(*int64)(unsafe.Add(mBase, uint32(v17)+240))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+32)) = v927
	v929 = *(*int32)(unsafe.Add(mBase, uint32(v17)+248))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+40)) = v929
	F_ExecAlterExtensionContentsRecurse(m, l0, v17+int32(48), v17+int32(32))
	mBase = m.M
	v936 = m.ExcPending
	if v936 != 0 {
		goto L1
	} else {
		goto L222
	}
L222:
	;
	goto L203
L223:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v947 = m.ExcPending
	if v947 != 0 {
		goto L1
	} else {
		goto L224
	}
L224:
	;
	v949 = F_getObjectDescription(m, l2, int32(0))
	mBase = m.M
	v950 = m.ExcPending
	if v950 != 0 {
		goto L1
	} else {
		goto L225
	}
L225:
	;
	v951 = F_get_extension_name(m, v21)
	mBase = m.M
	v952 = m.ExcPending
	if v952 != 0 {
		goto L1
	} else {
		goto L226
	}
L226:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+132)) = v951
	*(*int32)(unsafe.Add(mBase, uint32(v17)+128)) = v949
	F_errmsg(m, int32(_a_F_ExecAlterExtensionContentsRecurse_10), v17+int32(128))
	mBase = m.M
	v959 = m.ExcPending
	if v959 != 0 {
		goto L1
	} else {
		goto L227
	}
L227:
	;
	F_errfinish(m, int32(_a_F_ExecAlterExtensionContentsRecurse_11), int32(3820), int32(_a_F_ExecAlterExtensionContentsRecurse_12))
	mBase = m.M
	v964 = m.ExcPending
	if v964 != 0 {
		goto L1
	} else {
		goto L228
	}
L228:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L229:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v971 = m.ExcPending
	if v971 != 0 {
		goto L1
	} else {
		goto L230
	}
L230:
	;
	v972 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v973 = F_get_namespace_name(m, v972)
	mBase = m.M
	v974 = m.ExcPending
	if v974 != 0 {
		goto L1
	} else {
		goto L231
	}
L231:
	;
	v975 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+4)) = v975
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v973
	F_errmsg(m, int32(_a_F_ExecAlterExtensionContentsRecurse_13), v17)
	mBase = m.M
	v980 = m.ExcPending
	if v980 != 0 {
		goto L1
	} else {
		goto L232
	}
L232:
	;
	F_errfinish(m, int32(_a_F_ExecAlterExtensionContentsRecurse_11), int32(3833), int32(_a_F_ExecAlterExtensionContentsRecurse_12))
	mBase = m.M
	v985 = m.ExcPending
	if v985 != 0 {
		goto L1
	} else {
		goto L233
	}
L233:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L234:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v992 = m.ExcPending
	if v992 != 0 {
		goto L1
	} else {
		goto L235
	}
L235:
	;
	v994 = F_getObjectDescription(m, l2, int32(0))
	mBase = m.M
	v995 = m.ExcPending
	if v995 != 0 {
		goto L1
	} else {
		goto L236
	}
L236:
	;
	v996 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+164)) = v996
	*(*int32)(unsafe.Add(mBase, uint32(v17)+160)) = v994
	F_errmsg(m, int32(_a_F_ExecAlterExtensionContentsRecurse_14), v17+int32(160))
	mBase = m.M
	v1003 = m.ExcPending
	if v1003 != 0 {
		goto L1
	} else {
		goto L237
	}
L237:
	;
	F_errfinish(m, int32(_a_F_ExecAlterExtensionContentsRecurse_11), int32(3860), int32(_a_F_ExecAlterExtensionContentsRecurse_12))
	mBase = m.M
	v1008 = m.ExcPending
	if v1008 != 0 {
		goto L1
	} else {
		goto L238
	}
L238:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L239:
	;
	F_errmsg_internal(m, int32(_a_F_ExecAlterExtensionContentsRecurse_15), int32(0))
	mBase = m.M
	v1016 = m.ExcPending
	if v1016 != 0 {
		goto L1
	} else {
		goto L240
	}
L240:
	;
	F_errfinish(m, int32(_a_F_ExecAlterExtensionContentsRecurse_11), int32(3868), int32(_a_F_ExecAlterExtensionContentsRecurse_12))
	mBase = m.M
	v1021 = m.ExcPending
	if v1021 != 0 {
		goto L1
	} else {
		goto L241
	}
L241:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L242:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+144)) = v21
	F_errmsg_internal(m, int32(_a_F_ExecAlterExtensionContentsRecurse_16), v17+int32(144))
	mBase = m.M
	v1031 = m.ExcPending
	if v1031 != 0 {
		goto L1
	} else {
		goto L243
	}
L243:
	;
	F_errfinish(m, int32(_a_F_ExecAlterExtensionContentsRecurse_11), int32(3058), int32(_a_F_ExecAlterExtensionContentsRecurse_17))
	mBase = m.M
	v1036 = m.ExcPending
	if v1036 != 0 {
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
	F_errmsg_internal(m, int32(_a_F_ExecAlterExtensionContentsRecurse_18), int32(0))
	mBase = m.M
	v1045 = m.ExcPending
	if v1045 != 0 {
		goto L1
	} else {
		goto L246
	}
L246:
	;
	F_errfinish(m, int32(_a_F_ExecAlterExtensionContentsRecurse_11), int32(3083), int32(_a_F_ExecAlterExtensionContentsRecurse_17))
	mBase = m.M
	v1050 = m.ExcPending
	if v1050 != 0 {
		goto L1
	} else {
		goto L247
	}
L247:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L248:
	;
	F_errmsg_internal(m, int32(_a_F_ExecAlterExtensionContentsRecurse_19), int32(0))
	mBase = m.M
	v1058 = m.ExcPending
	if v1058 != 0 {
		goto L1
	} else {
		goto L249
	}
L249:
	;
	F_errfinish(m, int32(_a_F_ExecAlterExtensionContentsRecurse_11), int32(3140), int32(_a_F_ExecAlterExtensionContentsRecurse_17))
	mBase = m.M
	v1063 = m.ExcPending
	if v1063 != 0 {
		goto L1
	} else {
		goto L250
	}
L250:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L251:
	;
	F_errmsg_internal(m, int32(_a_F_ExecAlterExtensionContentsRecurse_20), int32(0))
	mBase = m.M
	v1071 = m.ExcPending
	if v1071 != 0 {
		goto L1
	} else {
		goto L252
	}
L252:
	;
	F_errfinish(m, int32(_a_F_ExecAlterExtensionContentsRecurse_11), int32(3150), int32(_a_F_ExecAlterExtensionContentsRecurse_17))
	mBase = m.M
	v1076 = m.ExcPending
	if v1076 != 0 {
		goto L1
	} else {
		goto L253
	}
L253:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L254:
	;
	F_errmsg_internal(m, int32(_a_F_ExecAlterExtensionContentsRecurse_19), int32(0))
	mBase = m.M
	v1084 = m.ExcPending
	if v1084 != 0 {
		goto L1
	} else {
		goto L255
	}
L255:
	;
	F_errfinish(m, int32(_a_F_ExecAlterExtensionContentsRecurse_11), int32(3152), int32(_a_F_ExecAlterExtensionContentsRecurse_17))
	mBase = m.M
	v1089 = m.ExcPending
	if v1089 != 0 {
		goto L1
	} else {
		goto L256
	}
L256:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L257:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v1096 = m.ExcPending
	if v1096 != 0 {
		goto L1
	} else {
		goto L258
	}
L258:
	;
	v1097 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v1098 = F_format_type_be(m, v1097)
	mBase = m.M
	v1099 = m.ExcPending
	if v1099 != 0 {
		goto L1
	} else {
		goto L259
	}
L259:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v1098
	F_errmsg(m, int32(_a_F_ExecAlterExtensionContentsRecurse_21), v17+int32(16))
	mBase = m.M
	v1105 = m.ExcPending
	if v1105 != 0 {
		goto L1
	} else {
		goto L260
	}
L260:
	;
	F_errfinish(m, int32(_a_F_ExecAlterExtensionContentsRecurse_11), int32(3913), int32(_a_F_ExecAlterExtensionContentsRecurse_12))
	mBase = m.M
	v1110 = m.ExcPending
	if v1110 != 0 {
		goto L1
	} else {
		goto L261
	}
L261:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_ExecAlterObjectSchemaStmt(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v111 int64
	_ = v111
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v130 int32
	_ = v130
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v288 int32
	_ = v288
	var v295 int32
	_ = v295
	var v300 int32
	_ = v300
	var v322 int32
	_ = v322
	var v344 int32
	_ = v344
	var v349 int32
	_ = v349
	var v351 int32
	_ = v351
	var v353 int32
	_ = v353
	var v355 int32
	_ = v355
	var v361 int32
	_ = v361
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v394 int32
	_ = v394
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v402 int32
	_ = v402
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v411 int32
	_ = v411
	var v415 int32
	_ = v415
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v465 int32
	_ = v465
	var v468 int32
	_ = v468
	var v473 int32
	_ = v473
	var v476 int32
	_ = v476
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v484 int32
	_ = v484
	var v486 int32
	_ = v486
	var v490 int32
	_ = v490
	var v523 int32
	_ = v523
	var v526 int32
	_ = v526
	var v530 int32
	_ = v530
	var v535 int32
	_ = v535
	var v539 int32
	_ = v539
	var v542 int32
	_ = v542
	var v549 int32
	_ = v549
	var v554 int32
	_ = v554
	var v558 int32
	_ = v558
	var v564 int32
	_ = v564
	var v569 int32
	_ = v569
	var v573 int32
	_ = v573
	var v576 int32
	_ = v576
	var v584 int32
	_ = v584
	var v589 int32
	_ = v589
	var v593 int32
	_ = v593
	var v597 int32
	_ = v597
	var v602 int32
	_ = v602
	var v606 int32
	_ = v606
	var v609 int32
	_ = v609
	var v615 int32
	_ = v615
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v629 int32
	_ = v629
	var v634 int32
	_ = v634
	var v638 int32
	_ = v638
	var v646 int32
	_ = v646
	var v651 int32
	_ = v651
	var v653 int32
	_ = v653
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v660 int32
	_ = v660
	var v662 int32
	_ = v662
	var v664 int32
	_ = v664
	var v666 int32
	_ = v666
	var v667 int32
	_ = v667
	var v672 int32
	_ = v672
	var v673 int32
	_ = v673
	var v674 int32
	_ = v674
	var v675 int32
	_ = v675
	var v679 int32
	_ = v679
	var v684 int32
	_ = v684
	var v686 int32
	_ = v686
	var v689 int64
	_ = v689
	var v692 int32
	_ = v692
	var v693 int32
	_ = v693
	var v694 int32
	_ = v694
	var v695 int32
	_ = v695
	var v696 int32
	_ = v696
	var v701 int32
	_ = v701
	var v703 int32
	_ = v703
	var v704 int32
	_ = v704
	var v705 int32
	_ = v705
	var v707 int32
	_ = v707
	var v708 int32
	_ = v708
	var v709 int32
	_ = v709
	var v710 int32
	_ = v710
	var v712 int32
	_ = v712
	var v716 int32
	_ = v716
	var v717 int32
	_ = v717
	var v718 int32
	_ = v718
	var v720 int32
	_ = v720
	var v721 int32
	_ = v721
	var v723 int32
	_ = v723
	var v724 int32
	_ = v724
	var v725 int32
	_ = v725
	var v727 int32
	_ = v727
	var v729 int32
	_ = v729
	var v733 int32
	_ = v733
	var v750 int32
	_ = v750
	var v753 int32
	_ = v753
	var v757 int32
	_ = v757
	var v758 int32
	_ = v758
	var v759 int32
	_ = v759
	var v760 int32
	_ = v760
	var v761 int32
	_ = v761
	var v770 int32
	_ = v770
	var v775 int32
	_ = v775
	var v777 int32
	_ = v777
	var v778 int32
	_ = v778
	var v779 int32
	_ = v779
	var v783 int32
	_ = v783
	var v784 int32
	_ = v784
	var v786 int32
	_ = v786
	var v789 int32
	_ = v789
	var v790 int32
	_ = v790
	var v791 int32
	_ = v791
	var v792 int32
	_ = v792
	var v795 int32
	_ = v795
	var v796 int32
	_ = v796
	var v799 int32
	_ = v799
	var v800 int32
	_ = v800
	var v802 int32
	_ = v802
	var v803 int32
	_ = v803
	var v804 int32
	_ = v804
	var v805 int32
	_ = v805
	var v807 int32
	_ = v807
	var v820 int32
	_ = v820
	var v823 int32
	_ = v823
	var v824 int32
	_ = v824
	var v825 int32
	_ = v825
	var v829 int32
	_ = v829
	var v834 int32
	_ = v834
	var v838 int32
	_ = v838
	var v839 int32
	_ = v839
	var v843 int32
	_ = v843
	var v848 int32
	_ = v848
	var v851 int32
	_ = v851
	var v852 int32
	_ = v852
	var v856 int32
	_ = v856
	var v857 int32
	_ = v857
	var v858 int32
	_ = v858
	var v859 int32
	_ = v859
	var v861 int32
	_ = v861
	var v862 int32
	_ = v862
	var v863 int32
	_ = v863
	var v864 int32
	_ = v864
	var v865 int32
	_ = v865
	var v866 int32
	_ = v866
	var v867 int32
	_ = v867
	var v871 int32
	_ = v871
	var v893 int32
	_ = v893
	var v894 int32
	_ = v894
	var v895 int32
	_ = v895
	var v899 int32
	_ = v899
	var v900 int32
	_ = v900
	var v901 int32
	_ = v901
	var v919 int32
	_ = v919
	v22 = m.G0
	v24 = v22 - int32(32)
	m.G0 = v24
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	switch v26 - int32(1) {
	case 0, 6, 7, 18, 23, 24, 25, 28, 33, 38, 44, 45, 46, 47:
		goto L3
	default:
		goto L4
	case 11, 48:
		goto L5
	case 14:
		goto L7
	case 17, 22, 36, 40, 50:
		goto L6
	}
L1:
	;
	if l2 != 0 {
		goto L240
	} else {
		goto L241
	}
L2:
	;
	v893 = *(*int32)(unsafe.Add(mBase, uint32(v24)+24))
	v894 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
	v895 = *(*int32)(unsafe.Add(mBase, uint32(v24)+16))
	v899 = v895
	v900 = v894
	v901 = v893
	goto L1
L3:
	;
	v851 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v852 = int32(0)
	F_get_object_address(m, v24+int32(16), v26, v851, v852, int32(8), v852)
	mBase = m.M
	v856 = m.ExcPending
	if v856 != 0 {
		goto L18
	} else {
		goto L235
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v838 = m.ExcPending
	if v838 != 0 {
		goto L18
	} else {
		goto L232
	}
L5:
	;
	v777 = v24 + int32(16)
	v778 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v779 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if l2 != 0 {
		goto L208
	} else {
		goto L209
	}
L6:
	;
	v653 = v24 + int32(16)
	if l2 != 0 {
		goto L168
	} else {
		goto L169
	}
L7:
	;
	v30 = v24 + int32(16)
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if l2 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v37 = v24 + int32(28)
	goto L10
L9:
	;
	v37 = int32(0)
	goto L10
L10:
	;
	v38 = m.G0
	v40 = v38 - int32(240)
	m.G0 = v40
	v43 = int32(0)
	v46 = F_GetSysCacheOid(m, int32(27), v32, v43, v43, v43)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L18
	} else {
		goto L19
	}
L11:
	;
	goto L2
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v638 = m.ExcPending
	if v638 != 0 {
		goto L18
	} else {
		goto L165
	}
L13:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v606 = m.ExcPending
	if v606 != 0 {
		goto L18
	} else {
		goto L158
	}
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v593 = m.ExcPending
	if v593 != 0 {
		goto L18
	} else {
		goto L155
	}
L15:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v573 = m.ExcPending
	if v573 != 0 {
		goto L18
	} else {
		goto L151
	}
L16:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v558 = m.ExcPending
	if v558 != 0 {
		goto L18
	} else {
		goto L148
	}
L17:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v539 = m.ExcPending
	if v539 != 0 {
		goto L18
	} else {
		goto L144
	}
L18:
	;
	return
L19:
	;
	if v46 != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v48 = F_LookupCreationNamespace(m, v33)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L18
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v523 = m.ExcPending
	if v523 != 0 {
		goto L18
	} else {
		goto L140
	}
L23:
	;
	v52 = *(*int32)(unsafe.Add(mBase, _c_F_ExecAlterObjectSchemaStmt[0]))
	v53 = F_object_ownercheck(m, int32(3079), v46, v52)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L18
	} else {
		goto L24
	}
L24:
	;
	if v53 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	F_aclcheck_error(m, int32(2), int32(15), v32)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L18
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	v63 = *(*int32)(unsafe.Add(mBase, _c_F_ExecAlterObjectSchemaStmt[0]))
	v65 = F_object_aclcheck(m, int32(2615), v48, v63, int64(512))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L18
	} else {
		goto L29
	}
L28:
	;
	goto L27
L29:
	;
	if v65 != 0 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	F_aclcheck_error(m, v65, int32(36), v33)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L18
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	v71 = F_getExtensionOfObject(m, int32(2615), v48)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L18
	} else {
		goto L34
	}
L33:
	;
	goto L32
L34:
	;
	if v71 == v46 {
		goto L17
	} else {
		goto L35
	}
L35:
	;
	v76 = F_table_open(m, int32(3079), int32(3))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L18
	} else {
		goto L36
	}
L36:
	;
	v79 = v40 + int32(144)
	F_ScanKeyInit(m, v79, int32(1), int32(3), int32(184), v46)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L18
	} else {
		goto L37
	}
L37:
	;
	v86 = int32(1)
	v89 = F_systable_beginscan(m, v76, int32(3080), v86, int32(0), v86, v79)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L18
	} else {
		goto L38
	}
L38:
	;
	v91 = F_systable_getnext(m, v89)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L18
	} else {
		goto L39
	}
L39:
	;
	if v91 == int32(0) {
		goto L16
	} else {
		goto L40
	}
L40:
	;
	v95 = F_heap_copytuple(m, v91)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L18
	} else {
		goto L41
	}
L41:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v95)+16))
	v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97)+22)))
	F_systable_endscan(m, v89)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L18
	} else {
		goto L42
	}
L42:
	;
	v101 = v97 + v98
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v101)+72))
	if v48 == v102 {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	m.G0 = v40 + int32(240)
	goto L11
L44:
	;
	F_relation_close(m, v76, int32(3))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L18
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101)+76)))
	if v113 == int32(0) {
		goto L15
	} else {
		goto L48
	}
L47:
	;
	v108 = *(*int32)(unsafe.Add(mBase, _c_F_ExecAlterObjectSchemaStmt[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v30)+8)) = v108
	v111 = *(*int64)(unsafe.Add(mBase, _c_F_ExecAlterObjectSchemaStmt[2]))
	*(*int64)(unsafe.Add(mBase, uint32(v30))) = v111
	goto L43
L48:
	;
	v116 = F_new_object_addresses(m)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L18
	} else {
		goto L49
	}
L49:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v101)+72))
	v121 = F_table_open(m, int32(2608), int32(1))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L18
	} else {
		goto L50
	}
L50:
	;
	v124 = v40 + int32(144)
	F_ScanKeyInit(m, v124, int32(4), int32(3), int32(184), int32(3079))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L18
	} else {
		goto L51
	}
L51:
	;
	F_ScanKeyInit(m, v40+int32(192), int32(5), int32(3), int32(184), v46)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L18
	} else {
		goto L52
	}
L52:
	;
	v142 = F_systable_beginscan(m, v121, int32(2674), int32(1), int32(0), int32(2), v124)
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L18
	} else {
		goto L53
	}
L53:
	;
	v144 = F_systable_getnext(m, v142)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L18
	} else {
		goto L54
	}
L54:
	;
	if v144 != 0 {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v147 = v101 + int32(4)
	v149 = v144
	goto L58
L56:
	;
	goto L57
L57:
	;
	if v37 != 0 {
		goto L127
	} else {
		goto L128
	}
L58:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v149)+16))
	v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v169)+22)))
	v171 = v169 + v170
	v172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v171)+24)))
	if v172 == int32(110) {
		goto L61
	} else {
		goto L62
	}
L59:
	;
	goto L57
L60:
	;
	v440 = F_systable_getnext(m, v142)
	mBase = m.M
	v441 = m.ExcPending
	if v441 != 0 {
		goto L18
	} else {
		goto L125
	}
L61:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v171)))
	if v175 != int32(3079) {
		goto L60
	} else {
		goto L64
	}
L62:
	;
	v344 = v172
	goto L63
L63:
	;
	if v344&int32(255) != int32(101) {
		goto L60
	} else {
		goto L99
	}
L64:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v171)+4))
	v180 = F_SearchSysCache1(m, int32(28), v179)
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L18
	} else {
		goto L66
	}
L65:
	;
	v196 = F_palloc0(m, int32(48))
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L18
	} else {
		goto L72
	}
L66:
	;
	if v180 == int32(0) {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v194 = int32(0)
	goto L65
L68:
	;
	goto L69
L69:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v180)+16))
	v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v185)+22)))
	v190 = F_pstrdup(m, v185+v186+int32(4))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L18
	} else {
		goto L70
	}
L70:
	;
	F_ReleaseCatCache(m, v180)
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L18
	} else {
		goto L71
	}
L71:
	;
	v194 = v190
	goto L65
L72:
	;
	v198 = F_pstrdup(m, v194)
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L18
	} else {
		goto L73
	}
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v196)+36)) = int32(-1)
	v202 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v196)+34)) = uint8(v202)
	v204 = int32(256)
	*(*uint16)(unsafe.Add(mBase, uint32(v196)+32)) = uint16(v204)
	*(*int32)(unsafe.Add(mBase, uint32(v196))) = v198
	F_parse_extension_control_file(m, v196, v202)
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L18
	} else {
		goto L74
	}
L74:
	;
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v196)+44))
	if v210 == int32(0) {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v322 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v171)+24)))
	v344 = v322
	goto L63
L76:
	;
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v210)+4))
	if v213 <= int32(0) {
		goto L75
	} else {
		goto L77
	}
L77:
	;
	v216 = int32(0)
	if v216 < v213 {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v219 = v213
	goto L80
L79:
	;
	v219 = v216
	goto L80
L80:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v210)+12))
	v223 = int32(0)
	goto L81
L81:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v220+v223<<(uint(int32(2))%32))))
	v249 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v246))))
	v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v147))))
	if base.B2i32(v249 == int32(0))|base.B2i32(v249 != v252) != 0 {
		v270 = v249
		v271 = v252
		goto L84
	} else {
		goto L85
	}
L82:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L18
	} else {
		goto L94
	}
L83:
	;
	if v270-v271 != 0 {
		goto L90
	} else {
		goto L91
	}
L84:
	;
	goto L83
L85:
	;
	v255 = v246
	v256 = v147
	goto L86
L86:
	;
	v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v256)+1)))
	v260 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v255)+1)))
	if v260 == int32(0) {
		v270 = v260
		v271 = v259
		goto L84
	} else {
		goto L88
	}
L87:
	;
	v270 = v260
	v271 = v259
	goto L84
L88:
	;
	v263 = int32(1)
	if v260 == v259 {
		v255 = v255 + v263
		v256 = v256 + v263
		goto L86
	} else {
		goto L89
	}
L89:
	;
	goto L87
L90:
	;
	v274 = v223 + int32(1)
	if v219 != v274 {
		v223 = v274
		goto L81
	} else {
		goto L93
	}
L91:
	;
	goto L92
L92:
	;
	goto L82
L93:
	;
	goto L75
L94:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L18
	} else {
		goto L95
	}
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+112)) = v147
	F_errmsg(m, int32(_a_F_ExecAlterObjectSchemaStmt_0), v40+int32(112))
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L18
	} else {
		goto L96
	}
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+100)) = v147
	*(*int32)(unsafe.Add(mBase, uint32(v40)+96)) = v194
	F_errdetail(m, int32(_a_F_ExecAlterObjectSchemaStmt_1), v40+int32(96))
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L18
	} else {
		goto L97
	}
L97:
	;
	F_errfinish(m, int32(_a_F_ExecAlterObjectSchemaStmt_2), int32(3336), int32(_a_F_ExecAlterObjectSchemaStmt_3))
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L18
	} else {
		goto L98
	}
L98:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L99:
	;
	v349 = *(*int32)(unsafe.Add(mBase, uint32(v171)))
	*(*int32)(unsafe.Add(mBase, uint32(v40)+132)) = v349
	v351 = *(*int32)(unsafe.Add(mBase, uint32(v171)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v40)+136)) = v351
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v171)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v40)+140)) = v353
	if v353 != 0 {
		goto L14
	} else {
		goto L100
	}
L100:
	;
	v355 = int32(0)
	if v349 <= int32(2752) {
		goto L106
	} else {
		goto L107
	}
L101:
	;
	if v415 == int32(0) {
		goto L60
	} else {
		goto L123
	}
L102:
	;
	v405 = F_table_open(m, v349, int32(3))
	mBase = m.M
	v406 = m.ExcPending
	if v406 != 0 {
		goto L18
	} else {
		goto L120
	}
L103:
	;
	v415 = v402
	goto L101
L104:
	;
	v399 = F_AlterTypeNamespace_oid(m, v351, v48, int32(1), v116)
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L18
	} else {
		goto L119
	}
L105:
	;
	v389 = F_relation_open(m, v351, int32(8))
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L18
	} else {
		goto L116
	}
L106:
	;
	switch v349 - int32(1247) {
	case 0:
		goto L104
	case 1, 2, 3, 4, 5, 6, 7, 9, 10, 11:
		v402 = v355
		goto L103
	case 8:
		goto L102
	case 12:
		goto L105
	default:
		goto L109
	}
L107:
	;
	goto L108
L108:
	;
	if v349 <= int32(3599) {
		goto L111
	} else {
		goto L112
	}
L109:
	;
	v361 = v349 - int32(2607)
	if base.B2i32(base.Ui32(int32(10)) < base.Ui32(v361))|base.B2i32(int32(1)<<(uint(v361)%32)&int32(1537) == int32(0)) != 0 {
		v402 = v355
		goto L103
	} else {
		goto L110
	}
L110:
	;
	goto L102
L111:
	;
	if base.B2i32(v349 == int32(2753))|base.B2i32(v349 == int32(3381))|base.B2i32(v349 == int32(3456)) != 0 {
		goto L102
	} else {
		goto L114
	}
L112:
	;
	goto L113
L113:
	;
	if base.B2i32(v349 == int32(3764))|base.B2i32(base.Ui32(v349-int32(3600)) < base.Ui32(int32(3))) != 0 {
		goto L102
	} else {
		goto L115
	}
L114:
	;
	v402 = v355
	goto L103
L115:
	;
	v402 = v355
	goto L103
L116:
	;
	v391 = *(*int32)(unsafe.Add(mBase, uint32(v389)+48))
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v391)+68))
	F_AlterTableNamespaceInternal(m, v389, v392, v48, v116)
	mBase = m.M
	v394 = m.ExcPending
	if v394 != 0 {
		goto L18
	} else {
		goto L117
	}
L117:
	;
	F_relation_close(m, v389, int32(0))
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
		goto L18
	} else {
		goto L118
	}
L118:
	;
	v415 = v392
	goto L101
L119:
	;
	v402 = v399
	goto L103
L120:
	;
	v407 = F_AlterObjectNamespace_internal(m, v405, v351, v48)
	mBase = m.M
	v408 = m.ExcPending
	if v408 != 0 {
		goto L18
	} else {
		goto L121
	}
L121:
	;
	F_relation_close(m, v405, int32(3))
	mBase = m.M
	v411 = m.ExcPending
	if v411 != 0 {
		goto L18
	} else {
		goto L122
	}
L122:
	;
	v415 = v407
	goto L101
L123:
	;
	if v415 != v118 {
		goto L13
	} else {
		goto L124
	}
L124:
	;
	goto L60
L125:
	;
	if v440 != 0 {
		v149 = v440
		goto L58
	} else {
		goto L126
	}
L126:
	;
	goto L59
L127:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37))) = v118
	goto L129
L128:
	;
	goto L129
L129:
	;
	F_systable_endscan(m, v142)
	mBase = m.M
	v465 = m.ExcPending
	if v465 != 0 {
		goto L18
	} else {
		goto L130
	}
L130:
	;
	F_relation_close(m, v121, int32(1))
	mBase = m.M
	v468 = m.ExcPending
	if v468 != 0 {
		goto L18
	} else {
		goto L131
	}
L131:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v101)+72)) = v48
	F_CatalogTupleUpdate(m, v76, v95+int32(4), v95)
	mBase = m.M
	v473 = m.ExcPending
	if v473 != 0 {
		goto L18
	} else {
		goto L132
	}
L132:
	;
	F_relation_close(m, v76, int32(3))
	mBase = m.M
	v476 = m.ExcPending
	if v476 != 0 {
		goto L18
	} else {
		goto L133
	}
L133:
	;
	v479 = F_changeDependencyFor(m, int32(3079), v46, int32(2615), v118, v48)
	mBase = m.M
	v480 = m.ExcPending
	if v480 != 0 {
		goto L18
	} else {
		goto L134
	}
L134:
	;
	if v479 != int32(1) {
		goto L12
	} else {
		goto L135
	}
L135:
	;
	v484 = *(*int32)(unsafe.Add(mBase, _c_F_ExecAlterObjectSchemaStmt[3]))
	if v484 != 0 {
		goto L136
	} else {
		goto L137
	}
L136:
	;
	v486 = int32(0)
	F_RunObjectPostAlterHook(m, int32(3079), v46, v486, v486, v486)
	mBase = m.M
	v490 = m.ExcPending
	if v490 != 0 {
		goto L18
	} else {
		goto L139
	}
L137:
	;
	goto L138
L138:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v30)+4)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v30))) = int32(3079)
	goto L43
L139:
	;
	goto L138
L140:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v526 = m.ExcPending
	if v526 != 0 {
		goto L18
	} else {
		goto L141
	}
L141:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40))) = v32
	F_errmsg(m, int32(_a_F_ExecAlterObjectSchemaStmt_4), v40)
	mBase = m.M
	v530 = m.ExcPending
	if v530 != 0 {
		goto L18
	} else {
		goto L142
	}
L142:
	;
	F_errfinish(m, int32(_a_F_ExecAlterObjectSchemaStmt_2), int32(199), int32(_a_F_ExecAlterObjectSchemaStmt_5))
	mBase = m.M
	v535 = m.ExcPending
	if v535 != 0 {
		goto L18
	} else {
		goto L143
	}
L143:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L144:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v542 = m.ExcPending
	if v542 != 0 {
		goto L18
	} else {
		goto L145
	}
L145:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+20)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v40)+16)) = v32
	F_errmsg(m, int32(_a_F_ExecAlterObjectSchemaStmt_6), v40+int32(16))
	mBase = m.M
	v549 = m.ExcPending
	if v549 != 0 {
		goto L18
	} else {
		goto L146
	}
L146:
	;
	F_errfinish(m, int32(_a_F_ExecAlterObjectSchemaStmt_2), int32(3236), int32(_a_F_ExecAlterObjectSchemaStmt_3))
	mBase = m.M
	v554 = m.ExcPending
	if v554 != 0 {
		goto L18
	} else {
		goto L147
	}
L147:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L148:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+32)) = v46
	F_errmsg_internal(m, int32(_a_F_ExecAlterObjectSchemaStmt_7), v40+int32(32))
	mBase = m.M
	v564 = m.ExcPending
	if v564 != 0 {
		goto L18
	} else {
		goto L149
	}
L149:
	;
	F_errfinish(m, int32(_a_F_ExecAlterObjectSchemaStmt_2), int32(3253), int32(_a_F_ExecAlterObjectSchemaStmt_3))
	mBase = m.M
	v569 = m.ExcPending
	if v569 != 0 {
		goto L18
	} else {
		goto L150
	}
L150:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L151:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v576 = m.ExcPending
	if v576 != 0 {
		goto L18
	} else {
		goto L152
	}
L152:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+128)) = v101 + int32(4)
	F_errmsg(m, int32(_a_F_ExecAlterObjectSchemaStmt_8), v40+int32(128))
	mBase = m.M
	v584 = m.ExcPending
	if v584 != 0 {
		goto L18
	} else {
		goto L153
	}
L153:
	;
	F_errfinish(m, int32(_a_F_ExecAlterObjectSchemaStmt_2), int32(3276), int32(_a_F_ExecAlterObjectSchemaStmt_3))
	mBase = m.M
	v589 = m.ExcPending
	if v589 != 0 {
		goto L18
	} else {
		goto L154
	}
L154:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L155:
	;
	F_errmsg_internal(m, int32(_a_F_ExecAlterObjectSchemaStmt_9), int32(0))
	mBase = m.M
	v597 = m.ExcPending
	if v597 != 0 {
		goto L18
	} else {
		goto L156
	}
L156:
	;
	F_errfinish(m, int32(_a_F_ExecAlterObjectSchemaStmt_2), int32(3354), int32(_a_F_ExecAlterObjectSchemaStmt_3))
	mBase = m.M
	v602 = m.ExcPending
	if v602 != 0 {
		goto L18
	} else {
		goto L157
	}
L157:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L158:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v609 = m.ExcPending
	if v609 != 0 {
		goto L18
	} else {
		goto L159
	}
L159:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+80)) = v147
	F_errmsg(m, int32(_a_F_ExecAlterObjectSchemaStmt_8), v40+int32(80))
	mBase = m.M
	v615 = m.ExcPending
	if v615 != 0 {
		goto L18
	} else {
		goto L160
	}
L160:
	;
	v619 = F_getObjectDescription(m, v40+int32(132), int32(0))
	mBase = m.M
	v620 = m.ExcPending
	if v620 != 0 {
		goto L18
	} else {
		goto L161
	}
L161:
	;
	v621 = F_get_namespace_name(m, v118)
	mBase = m.M
	v622 = m.ExcPending
	if v622 != 0 {
		goto L18
	} else {
		goto L162
	}
L162:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+68)) = v621
	*(*int32)(unsafe.Add(mBase, uint32(v40)+64)) = v619
	F_errdetail(m, int32(_a_F_ExecAlterObjectSchemaStmt_10), v40-int32(-64))
	mBase = m.M
	v629 = m.ExcPending
	if v629 != 0 {
		goto L18
	} else {
		goto L163
	}
L163:
	;
	F_errfinish(m, int32(_a_F_ExecAlterObjectSchemaStmt_2), int32(3373), int32(_a_F_ExecAlterObjectSchemaStmt_3))
	mBase = m.M
	v634 = m.ExcPending
	if v634 != 0 {
		goto L18
	} else {
		goto L164
	}
L164:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L165:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+48)) = v101 + int32(4)
	F_errmsg_internal(m, int32(_a_F_ExecAlterObjectSchemaStmt_11), v40+int32(48))
	mBase = m.M
	v646 = m.ExcPending
	if v646 != 0 {
		goto L18
	} else {
		goto L166
	}
L166:
	;
	F_errfinish(m, int32(_a_F_ExecAlterObjectSchemaStmt_2), int32(3395), int32(_a_F_ExecAlterObjectSchemaStmt_3))
	mBase = m.M
	v651 = m.ExcPending
	if v651 != 0 {
		goto L18
	} else {
		goto L167
	}
L167:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L168:
	;
	v657 = v24 + int32(28)
	goto L170
L169:
	;
	v657 = int32(0)
	goto L170
L170:
	;
	v658 = m.G0
	v660 = v658 - int32(32)
	m.G0 = v660
	v662 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v664 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)))
	v666 = F_RangeVarGetRelidExtended(m, v662, int32(8), v664, int32(576), l1)
	mBase = m.M
	v667 = m.ExcPending
	if v667 != 0 {
		goto L18
	} else {
		goto L174
	}
L171:
	;
	goto L2
L172:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v750 = m.ExcPending
	if v750 != 0 {
		goto L18
	} else {
		goto L202
	}
L173:
	;
	m.G0 = v660 + int32(32)
	goto L171
L174:
	;
	if v666 == int32(0) {
		goto L175
	} else {
		goto L176
	}
L175:
	;
	v672 = F_errstart(m, int32(18), int32(0))
	mBase = m.M
	v673 = m.ExcPending
	if v673 != 0 {
		goto L18
	} else {
		goto L178
	}
L176:
	;
	goto L177
L177:
	;
	v692 = F_relation_open(m, v666, int32(0))
	mBase = m.M
	v693 = m.ExcPending
	if v693 != 0 {
		goto L18
	} else {
		goto L184
	}
L178:
	;
	if v672 != 0 {
		goto L179
	} else {
		goto L180
	}
L179:
	;
	v674 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v675 = *(*int32)(unsafe.Add(mBase, uint32(v674)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v660))) = v675
	F_errmsg(m, int32(_a_F_ExecAlterObjectSchemaStmt_12), v660)
	mBase = m.M
	v679 = m.ExcPending
	if v679 != 0 {
		goto L18
	} else {
		goto L182
	}
L180:
	;
	goto L181
L181:
	;
	v686 = *(*int32)(unsafe.Add(mBase, _c_F_ExecAlterObjectSchemaStmt[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v653)+8)) = v686
	v689 = *(*int64)(unsafe.Add(mBase, _c_F_ExecAlterObjectSchemaStmt[2]))
	*(*int64)(unsafe.Add(mBase, uint32(v653))) = v689
	goto L173
L182:
	;
	F_errfinish(m, int32(_a_F_ExecAlterObjectSchemaStmt_13), int32(_a_F_ExecAlterObjectSchemaStmt_14), int32(_a_F_ExecAlterObjectSchemaStmt_15))
	mBase = m.M
	v684 = m.ExcPending
	if v684 != 0 {
		goto L18
	} else {
		goto L183
	}
L183:
	;
	goto L181
L184:
	;
	v694 = *(*int32)(unsafe.Add(mBase, uint32(v692)+48))
	v695 = *(*int32)(unsafe.Add(mBase, uint32(v694)+68))
	v696 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v694)+119)))
	if v696 == int32(83) {
		goto L185
	} else {
		goto L186
	}
L185:
	;
	v701 = v660 + int32(28)
	v703 = v660 + int32(24)
	v704 = F_sequenceIsOwned(m, v666, int32(97), v701, v703)
	mBase = m.M
	v705 = m.ExcPending
	if v705 != 0 {
		goto L18
	} else {
		goto L188
	}
L186:
	;
	v710 = v694
	goto L187
L187:
	;
	v712 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v716 = F_makeRangeVar(m, v712, v710+int32(4), int32(-1))
	mBase = m.M
	v717 = m.ExcPending
	if v717 != 0 {
		goto L18
	} else {
		goto L192
	}
L188:
	;
	if v704 != 0 {
		goto L172
	} else {
		goto L189
	}
L189:
	;
	v707 = F_sequenceIsOwned(m, v666, int32(105), v701, v703)
	mBase = m.M
	v708 = m.ExcPending
	if v708 != 0 {
		goto L18
	} else {
		goto L190
	}
L190:
	;
	if v707 != 0 {
		goto L172
	} else {
		goto L191
	}
L191:
	;
	v709 = *(*int32)(unsafe.Add(mBase, uint32(v692)+48))
	v710 = v709
	goto L187
L192:
	;
	v718 = int32(0)
	v720 = F_RangeVarGetAndCheckCreationNamespace(m, v716, v718, v718)
	mBase = m.M
	v721 = m.ExcPending
	if v721 != 0 {
		goto L18
	} else {
		goto L193
	}
L193:
	;
	F_CheckSetNamespace(m, v695, v720)
	mBase = m.M
	v723 = m.ExcPending
	if v723 != 0 {
		goto L18
	} else {
		goto L194
	}
L194:
	;
	v724 = F_new_object_addresses(m)
	mBase = m.M
	v725 = m.ExcPending
	if v725 != 0 {
		goto L18
	} else {
		goto L195
	}
L195:
	;
	F_AlterTableNamespaceInternal(m, v692, v695, v720, v724)
	mBase = m.M
	v727 = m.ExcPending
	if v727 != 0 {
		goto L18
	} else {
		goto L196
	}
L196:
	;
	F_free_object_addresses(m, v724)
	mBase = m.M
	v729 = m.ExcPending
	if v729 != 0 {
		goto L18
	} else {
		goto L197
	}
L197:
	;
	if v657 != 0 {
		goto L198
	} else {
		goto L199
	}
L198:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v657))) = v695
	goto L200
L199:
	;
	goto L200
L200:
	;
	F_relation_close(m, v692, int32(0))
	mBase = m.M
	v733 = m.ExcPending
	if v733 != 0 {
		goto L18
	} else {
		goto L201
	}
L201:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v653)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v653)+4)) = v666
	*(*int32)(unsafe.Add(mBase, uint32(v653))) = int32(1259)
	goto L173
L202:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v753 = m.ExcPending
	if v753 != 0 {
		goto L18
	} else {
		goto L203
	}
L203:
	;
	F_errmsg(m, int32(_a_F_ExecAlterObjectSchemaStmt_16), int32(0))
	mBase = m.M
	v757 = m.ExcPending
	if v757 != 0 {
		goto L18
	} else {
		goto L204
	}
L204:
	;
	v758 = *(*int32)(unsafe.Add(mBase, uint32(v692)+48))
	v759 = *(*int32)(unsafe.Add(mBase, uint32(v660)+28))
	v760 = F_get_rel_name(m, v759)
	mBase = m.M
	v761 = m.ExcPending
	if v761 != 0 {
		goto L18
	} else {
		goto L205
	}
L205:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v660)+20)) = v760
	*(*int32)(unsafe.Add(mBase, uint32(v660)+16)) = v758 + int32(4)
	F_errdetail(m, int32(_a_F_ExecAlterObjectSchemaStmt_17), v660+int32(16))
	mBase = m.M
	v770 = m.ExcPending
	if v770 != 0 {
		goto L18
	} else {
		goto L206
	}
L206:
	;
	F_errfinish(m, int32(_a_F_ExecAlterObjectSchemaStmt_13), int32(_a_F_ExecAlterObjectSchemaStmt_18), int32(_a_F_ExecAlterObjectSchemaStmt_15))
	mBase = m.M
	v775 = m.ExcPending
	if v775 != 0 {
		goto L18
	} else {
		goto L207
	}
L207:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L208:
	;
	v783 = v24 + int32(28)
	goto L210
L209:
	;
	v783 = int32(0)
	goto L210
L210:
	;
	v784 = m.G0
	v786 = v784 - int32(16)
	m.G0 = v786
	v789 = F_makeTypeNameFromNameList(m, v778)
	mBase = m.M
	v790 = m.ExcPending
	if v790 != 0 {
		goto L18
	} else {
		goto L211
	}
L211:
	;
	v791 = F_typenameTypeId(m, int32(0), v789)
	mBase = m.M
	v792 = m.ExcPending
	if v792 != 0 {
		goto L18
	} else {
		goto L212
	}
L212:
	;
	if v26 == int32(12) {
		goto L215
	} else {
		goto L216
	}
L213:
	;
	goto L2
L214:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v820 = m.ExcPending
	if v820 != 0 {
		goto L18
	} else {
		goto L227
	}
L215:
	;
	v795 = F_get_typtype(m, v791)
	mBase = m.M
	v796 = m.ExcPending
	if v796 != 0 {
		goto L18
	} else {
		goto L218
	}
L216:
	;
	goto L217
L217:
	;
	v799 = F_LookupCreationNamespace(m, v779)
	mBase = m.M
	v800 = m.ExcPending
	if v800 != 0 {
		goto L18
	} else {
		goto L220
	}
L218:
	;
	if v795 != int32(100) {
		goto L214
	} else {
		goto L219
	}
L219:
	;
	goto L217
L220:
	;
	v802 = F_new_object_addresses(m)
	mBase = m.M
	v803 = m.ExcPending
	if v803 != 0 {
		goto L18
	} else {
		goto L221
	}
L221:
	;
	v804 = F_AlterTypeNamespace_oid(m, v791, v799, int32(0), v802)
	mBase = m.M
	v805 = m.ExcPending
	if v805 != 0 {
		goto L18
	} else {
		goto L222
	}
L222:
	;
	F_free_object_addresses(m, v802)
	mBase = m.M
	v807 = m.ExcPending
	if v807 != 0 {
		goto L18
	} else {
		goto L223
	}
L223:
	;
	if v783 != 0 {
		goto L224
	} else {
		goto L225
	}
L224:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v783))) = v804
	goto L226
L225:
	;
	goto L226
L226:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v777)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v777)+4)) = v791
	*(*int32)(unsafe.Add(mBase, uint32(v777))) = int32(1247)
	m.G0 = v786 + int32(16)
	goto L213
L227:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v823 = m.ExcPending
	if v823 != 0 {
		goto L18
	} else {
		goto L228
	}
L228:
	;
	v824 = F_format_type_be(m, v791)
	mBase = m.M
	v825 = m.ExcPending
	if v825 != 0 {
		goto L18
	} else {
		goto L229
	}
L229:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v786))) = v824
	F_errmsg(m, int32(_a_F_ExecAlterObjectSchemaStmt_19), v786)
	mBase = m.M
	v829 = m.ExcPending
	if v829 != 0 {
		goto L18
	} else {
		goto L230
	}
L230:
	;
	F_errfinish(m, int32(_a_F_ExecAlterObjectSchemaStmt_20), int32(4073), int32(_a_F_ExecAlterObjectSchemaStmt_21))
	mBase = m.M
	v834 = m.ExcPending
	if v834 != 0 {
		goto L18
	} else {
		goto L231
	}
L231:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L232:
	;
	v839 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v24))) = v839
	F_errmsg_internal(m, int32(_a_F_ExecAlterObjectSchemaStmt_22), v24)
	mBase = m.M
	v843 = m.ExcPending
	if v843 != 0 {
		goto L18
	} else {
		goto L233
	}
L233:
	;
	F_errfinish(m, int32(_a_F_ExecAlterObjectSchemaStmt_23), int32(600), int32(_a_F_ExecAlterObjectSchemaStmt_24))
	mBase = m.M
	v848 = m.ExcPending
	if v848 != 0 {
		goto L18
	} else {
		goto L234
	}
L234:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L235:
	;
	v857 = *(*int32)(unsafe.Add(mBase, uint32(v24)+24))
	v858 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
	v859 = *(*int32)(unsafe.Add(mBase, uint32(v24)+16))
	v861 = F_table_open(m, v859, int32(3))
	mBase = m.M
	v862 = m.ExcPending
	if v862 != 0 {
		goto L18
	} else {
		goto L236
	}
L236:
	;
	v863 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v864 = F_LookupCreationNamespace(m, v863)
	mBase = m.M
	v865 = m.ExcPending
	if v865 != 0 {
		goto L18
	} else {
		goto L237
	}
L237:
	;
	v866 = F_AlterObjectNamespace_internal(m, v861, v858, v864)
	mBase = m.M
	v867 = m.ExcPending
	if v867 != 0 {
		goto L18
	} else {
		goto L238
	}
L238:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+28)) = v866
	F_relation_close(m, v861, int32(3))
	mBase = m.M
	v871 = m.ExcPending
	if v871 != 0 {
		goto L18
	} else {
		goto L239
	}
L239:
	;
	v899 = v859
	v900 = v858
	v901 = v857
	goto L1
L240:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(2615)
	v919 = *(*int32)(unsafe.Add(mBase, uint32(v24)+28))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v919
	goto L242
L241:
	;
	goto L242
L242:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v901
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v900
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v899
	m.G0 = v24 + int32(32)
	return
}
func F__equalAlterTableSpaceOptionsStmt(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	v3 = int32(0)
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v7 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return v52
L2:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v41 = F_equal(m, v39, v40)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L16
	} else {
		goto L17
	}
L3:
	;
	if v6 == int32(0) {
		v52 = v3
		goto L1
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	if v6 == v7 {
		goto L2
	} else {
		goto L15
	}
L6:
	;
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6))))
	if base.B2i32(v12 == int32(0))|base.B2i32(v12 != v15) != 0 {
		v33 = v12
		v34 = v15
		goto L8
	} else {
		goto L9
	}
L7:
	;
	if v33-v34 != 0 {
		v52 = v3
		goto L1
	} else {
		goto L14
	}
L8:
	;
	goto L7
L9:
	;
	v18 = v7
	v19 = v6
	goto L10
L10:
	;
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+1)))
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+1)))
	if v23 == int32(0) {
		v33 = v23
		v34 = v22
		goto L8
	} else {
		goto L12
	}
L11:
	;
	v33 = v23
	v34 = v22
	goto L8
L12:
	;
	v26 = int32(1)
	if v23 == v22 {
		v18 = v18 + v26
		v19 = v19 + v26
		goto L10
	} else {
		goto L13
	}
L13:
	;
	goto L11
L14:
	;
	goto L2
L15:
	;
	return int32(0)
L16:
	;
	return int32(0)
L17:
	;
	if v41 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	return int32(0)
L19:
	;
	goto L20
L20:
	;
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)))
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+12)))
	v52 = base.B2i32(v49 == v50)
	goto L1
}
