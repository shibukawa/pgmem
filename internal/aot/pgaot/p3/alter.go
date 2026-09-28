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
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
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
	var v724 int32
	_ = v724
	var v725 int32
	_ = v725
	var v734 int32
	_ = v734
	var __phi734 int32
	_ = __phi734
	var v736 int32
	_ = v736
	var __phi736 int32
	_ = __phi736
	var v744 int32
	_ = v744
	var v745 int32
	_ = v745
	var v753 int32
	_ = v753
	var v757 int32
	_ = v757
	var v762 int32
	_ = v762
	var v763 int32
	_ = v763
	var v770 int32
	_ = v770
	var v782 int32
	_ = v782
	var v783 int32
	_ = v783
	var v793 int32
	_ = v793
	var v804 int32
	_ = v804
	var v809 int32
	_ = v809
	var v810 int32
	_ = v810
	var v815 int32
	_ = v815
	var v816 int32
	_ = v816
	var v821 int32
	_ = v821
	var v826 int32
	_ = v826
	var v833 int32
	_ = v833
	var v838 int32
	_ = v838
	var v839 int32
	_ = v839
	var v852 int32
	_ = v852
	var v853 int32
	_ = v853
	var v858 int32
	_ = v858
	var v862 int32
	_ = v862
	var v863 int32
	_ = v863
	var v886 int32
	_ = v886
	var v887 int32
	_ = v887
	var v891 int32
	_ = v891
	var v892 int32
	_ = v892
	var v897 int32
	_ = v897
	var v898 int32
	_ = v898
	var v906 int32
	_ = v906
	var v907 int32
	_ = v907
	var v908 int32
	_ = v908
	var v915 int32
	_ = v915
	var v920 int32
	_ = v920
	var v927 int32
	_ = v927
	var v940 int32
	_ = v940
	var v941 int32
	_ = v941
	var v942 int32
	_ = v942
	var v948 int32
	_ = v948
	var v954 int32
	_ = v954
	var v956 int32
	_ = v956
	var v957 int32
	_ = v957
	var v968 int32
	_ = v968
	var v973 int32
	_ = v973
	var v982 int32
	_ = v982
	var v990 int32
	_ = v990
	var v993 int32
	_ = v993
	var v995 int32
	_ = v995
	var v997 int32
	_ = v997
	var v1005 int32
	_ = v1005
	var v1006 int32
	_ = v1006
	var v1007 int32
	_ = v1007
	var v1012 int32
	_ = v1012
	var v1028 int32
	_ = v1028
	var v1029 int32
	_ = v1029
	var v1031 int32
	_ = v1031
	var v1033 int32
	_ = v1033
	var v1039 int32
	_ = v1039
	var v1045 int32
	_ = v1045
	var v1055 int32
	_ = v1055
	var v1056 int32
	_ = v1056
	var v1057 int32
	_ = v1057
	var v1058 int32
	_ = v1058
	var v1061 int32
	_ = v1061
	var v1073 int32
	_ = v1073
	var v1078 int32
	_ = v1078
	var v1089 int32
	_ = v1089
	var v1097 int32
	_ = v1097
	var v1107 int32
	_ = v1107
	var v1118 int32
	_ = v1118
	var v1119 int32
	_ = v1119
	var v1120 int32
	_ = v1120
	var v1126 int32
	_ = v1126
	var v1131 int32
	_ = v1131
	var v1137 int32
	_ = v1137
	var v1138 int32
	_ = v1138
	var v1142 int32
	_ = v1142
	var v1145 int32
	_ = v1145
	var v1146 int32
	_ = v1146
	var v1150 int32
	_ = v1150
	var v1165 int32
	_ = v1165
	var v1169 int32
	_ = v1169
	var v1170 int32
	_ = v1170
	var v1175 int32
	_ = v1175
	var v1176 int32
	_ = v1176
	var v1188 int32
	_ = v1188
	var v1192 int32
	_ = v1192
	var v1193 int32
	_ = v1193
	var v1205 int32
	_ = v1205
	var v1208 int32
	_ = v1208
	var v1209 int32
	_ = v1209
	var v1211 int32
	_ = v1211
	var v1213 int32
	_ = v1213
	var v1223 int32
	_ = v1223
	var v1228 int32
	_ = v1228
	var v1234 int32
	_ = v1234
	var v1246 int32
	_ = v1246
	var v1257 int32
	_ = v1257
	var v1268 int32
	_ = v1268
	var v1294 int32
	_ = v1294
	var v1300 int32
	_ = v1300
	var v1307 int32
	_ = v1307
	var v1315 int32
	_ = v1315
	var v1320 int32
	_ = v1320
	var v1322 int32
	_ = v1322
	var v1329 int32
	_ = v1329
	var v1336 int32
	_ = v1336
	var v1337 int32
	_ = v1337
	var v1338 int32
	_ = v1338
	var v1339 int32
	_ = v1339
	var v1341 int32
	_ = v1341
	var v1345 int32
	_ = v1345
	var v1357 int32
	_ = v1357
	var v1362 int32
	_ = v1362
	var v1373 int32
	_ = v1373
	var v1381 int32
	_ = v1381
	var v1386 int32
	_ = v1386
	var v1392 int32
	_ = v1392
	var v1393 int32
	_ = v1393
	var v1399 int32
	_ = v1399
	var v1404 int32
	_ = v1404
	var v1426 int32
	_ = v1426
	var v1431 int32
	_ = v1431
	var v1435 int32
	_ = v1435
	var v1440 int32
	_ = v1440
	var v1447 int32
	_ = v1447
	var v1452 int32
	_ = v1452
	var v1461 int32
	_ = v1461
	var v1469 int32
	_ = v1469
	var v1473 int32
	_ = v1473
	var v1475 int32
	_ = v1475
	var v1479 int32
	_ = v1479
	var v1480 int32
	_ = v1480
	var v1491 int32
	_ = v1491
	var v1492 int32
	_ = v1492
	var v1500 int32
	_ = v1500
	var v1502 int32
	_ = v1502
	var v1507 int32
	_ = v1507
	var v1511 int32
	_ = v1511
	var v1528 int32
	_ = v1528
	var v1529 int64
	_ = v1529
	var v1533 int32
	_ = v1533
	var v1535 int32
	_ = v1535
	var v1536 int32
	_ = v1536
	var v1539 int32
	_ = v1539
	var v1541 int32
	_ = v1541
	var v1543 int32
	_ = v1543
	var v1544 int32
	_ = v1544
	var v1545 int32
	_ = v1545
	var v1546 int32
	_ = v1546
	var v1548 int32
	_ = v1548
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
	goto L2
L1:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2:
	;
	goto L4
L3:
	;
	m.G0 = v23 + int32(2592)
	return
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
	v1528 = int32(m.ExcTag)
	v1529 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v1528 == int32(0) {
		goto L297
	} else {
		goto L298
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
	v1005 = v24
	v1006 = v25
	v1007 = v26
	v1012 = v31
	goto L9
L9:
	;
	if v1012 == int32(0) {
		goto L211
	} else {
		goto L212
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
		goto L23
	case 1, 4:
		v111 = v26
		v112 = int32(0)
		goto L22
	default:
		goto L24
	case 5:
		goto L21
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
	goto L1
L17:
	;
	v940 = *(*int32)(unsafe.Add(mBase, _c_F_AlterSystemSetConfigFile[1]))
	if v940 != 0 {
		goto L194
	} else {
		goto L195
	}
L18:
	;
	if v112 == int32(0) {
		v927 = v111
		goto L17
	} else {
		goto L186
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v26
	*(*int32)(unsafe.Add(mBase, uint32(v23)+208)) = int32(_a_F_AlterSystemSetConfigFile_4)
	v833 = v23 + int32(1520)
	v838 = F_pg_snprintf(m, v833, int32(1024), int32(_a_F_AlterSystemSetConfigFile_5), v23+int32(208))
	mBase = m.M
	v839 = m.ExcPending
	if v839 != 0 {
		goto L6
	} else {
		goto L183
	}
L20:
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
		goto L45
	}
L21:
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
		goto L37
	}
L22:
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
		goto L29
	}
L23:
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
		goto L28
	}
L24:
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
		goto L25
	}
L25:
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
		goto L26
	}
L26:
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
		goto L27
	}
L27:
	;
	goto L1
L28:
	;
	v111 = v109
	v112 = v109
	goto L22
L29:
	;
	if v116 != 0 {
		goto L20
	} else {
		goto L30
	}
L30:
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
		goto L31
	}
L31:
	;
	if v127 == int32(0) {
		goto L20
	} else {
		goto L32
	}
L32:
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
		goto L33
	}
L33:
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
		goto L34
	}
L34:
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
		goto L35
	}
L35:
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
		goto L36
	}
L36:
	;
	goto L1
L37:
	;
	if v164 != 0 {
		goto L19
	} else {
		goto L38
	}
L38:
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
		goto L39
	}
L39:
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
		goto L40
	}
L40:
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
		goto L41
	}
L41:
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
		goto L42
	}
L42:
	;
	goto L1
L43:
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
		goto L124
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v111
	v534 = int32(10)
	v535 = F___strchrnul(m, v112, v534)
	mBase = m.M
	v537 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v535))))
	if v537 == v534 {
		goto L116
	} else {
		goto L117
	}
L45:
	;
	if v201 != 0 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v201)+4))
	if v203 != 0 {
		goto L50
	} else {
		goto L51
	}
L47:
	;
	goto L48
L48:
	;
	if v112 != 0 {
		goto L74
	} else {
		goto L75
	}
L49:
	;
	if v112 == int32(0) {
		goto L43
	} else {
		goto L58
	}
L50:
	;
	v204 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v201)+21)))
	if v204&int32(33) == int32(0) {
		goto L49
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
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
		goto L54
	}
L53:
	;
	goto L52
L54:
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
		goto L55
	}
L55:
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
		goto L56
	}
L56:
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
		goto L57
	}
L57:
	;
	goto L1
L58:
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
		goto L59
	}
L59:
	;
	if v252 == int32(0) {
		goto L60
	} else {
		goto L61
	}
L60:
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
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v201)+24))
	if v287 != int32(3) {
		goto L67
	} else {
		goto L68
	}
L63:
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
		goto L64
	}
L64:
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
		goto L65
	}
L65:
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
		goto L66
	}
L66:
	;
	goto L1
L67:
	;
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v23)+484))
	if v299 == int32(0) {
		goto L44
	} else {
		goto L71
	}
L68:
	;
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v23)+488))
	if v290 == int32(0) {
		goto L67
	} else {
		goto L69
	}
L69:
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
		goto L70
	}
L70:
	;
	goto L67
L71:
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
		goto L72
	}
L72:
	;
	goto L44
L73:
	;
	if v112 == int32(0) {
		goto L43
	} else {
		goto L114
	}
L74:
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
		goto L113
	}
L75:
	;
	v307 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77))))
	if v307 == int32(0) {
		goto L74
	} else {
		goto L76
	}
L76:
	;
	v318 = int32(0)
	v319 = v77
	v320 = v307
	v323 = int32(1)
	goto L77
L77:
	;
	v329 = v320 & int32(255)
	v331 = base.B2i32(v329 != int32(46))
	if v331 == int32(0) {
		goto L80
	} else {
		goto L81
	}
L78:
	;
	if v466&v331 != 0 {
		goto L73
	} else {
		goto L112
	}
L79:
	;
	v472 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v319)+1)))
	if v472 != 0 {
		v318 = v466
		v319 = v319 + int32(1)
		v320 = v472
		v323 = base.B2i32(v329 == int32(46))
		goto L77
	} else {
		goto L111
	}
L80:
	;
	v334 = int32(1)
	if v323&v334 == int32(0) {
		v466 = v334
		goto L79
	} else {
		goto L83
	}
L81:
	;
	goto L82
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v111
	v342 = int32(_a_F_AlterSystemSetConfigFile_16)
	v343 = base.I32_extend8_s(v320)
	v344 = int32(54)
	goto L87
L83:
	;
	goto L74
L84:
	;
	if v449|base.B2i32(v343 < int32(0)) != 0 {
		v466 = v318
		goto L79
	} else {
		goto L109
	}
L85:
	;
	v449 = int32(0)
	goto L84
L86:
	;
	v427 = v420
	v429 = v422
	goto L103
L87:
	;
	goto L94
L94:
	;
	v383 = v343 & int32(255)
	v384 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_AlterSystemSetConfigFile[3])))
	if base.B2i32(v383 == v384)|int32(0) == int32(0) {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	v393 = v342
	v395 = v344
	goto L98
L96:
	;
	v413 = v342
	v415 = v344
	goto L97
L97:
	;
	if v415 == int32(0) {
		goto L85
	} else {
		goto L102
	}
L98:
	;
	v399 = *(*int32)(unsafe.Add(mBase, uint32(v393)))
	v400 = v399 ^ v383*int32(16843009)
	v403 = int32(-2139062144)
	if (int32(16843008)-v400|v400)&v403 != v403 {
		v420 = v393
		v422 = v395
		goto L86
	} else {
		goto L100
	}
L99:
	;
	v413 = v408
	v415 = v410
	goto L97
L100:
	;
	v407 = int32(4)
	v408 = v393 + v407
	v410 = v395 - v407
	if base.Ui32(int32(3)) < base.Ui32(v410) {
		v393 = v408
		v395 = v410
		goto L98
	} else {
		goto L101
	}
L101:
	;
	goto L99
L102:
	;
	v420 = v413
	v422 = v415
	goto L86
L103:
	;
	v432 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v427))))
	if v343&int32(255) == v432 {
		goto L105
	} else {
		goto L106
	}
L104:
	;
	goto L85
L105:
	;
	v449 = v427
	goto L84
L106:
	;
	goto L107
L107:
	;
	v434 = int32(1)
	v437 = v429 - v434
	if v437 != 0 {
		v427 = v427 + v434
		v429 = v437
		goto L103
	} else {
		goto L108
	}
L108:
	;
	goto L104
L109:
	;
	if base.B2i32(int64(1)<<(uint(base.I64_extend_i32_u(v343))%64)&int64(287948969894477825) == int64(0))|(v323&int32(1)|base.B2i32(base.Ui32(int32(63)) < base.Ui32(v329))) != 0 {
		goto L74
	} else {
		goto L110
	}
L110:
	;
	v466 = v318
	goto L79
L111:
	;
	goto L78
L112:
	;
	goto L74
L113:
	;
	goto L73
L114:
	;
	goto L44
L115:
	;
	if v541 == int32(0) {
		goto L43
	} else {
		goto L119
	}
L116:
	;
	v541 = v535
	goto L118
L117:
	;
	v541 = int32(0)
	goto L118
L118:
	;
	goto L115
L119:
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
		goto L120
	}
L120:
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
		goto L121
	}
L121:
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
		goto L122
	}
L122:
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
		goto L123
	}
L123:
	;
	goto L1
L124:
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
		goto L125
	}
L125:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v111
	v619 = *(*int32)(unsafe.Add(mBase, _c_F_AlterSystemSetConfigFile[4]))
	v623 = F_LWLockAcquire(m, v619+int32(_a_F_AlterSystemSetConfigFile_21), int32(0))
	mBase = m.M
	v624 = m.ExcPending
	if v624 != 0 {
		goto L6
	} else {
		goto L126
	}
L126:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v111
	v632 = F___fstatat(m, int32(-100), v594, v23+int32(384), int32(0))
	mBase = m.M
	goto L127
L127:
	;
	if v632 == int32(0) {
		goto L128
	} else {
		goto L129
	}
L128:
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
		goto L131
	}
L129:
	;
	goto L130
L130:
	;
	v724 = int32(0)
	v725 = *(*int32)(unsafe.Add(mBase, uint32(v23)+2560))
	if v725 == v724 {
		goto L18
	} else {
		goto L148
	}
L131:
	;
	if v639 == int32(0) {
		goto L132
	} else {
		goto L133
	}
L132:
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
		goto L135
	}
L133:
	;
	goto L134
L134:
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
		goto L139
	}
L135:
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
		goto L136
	}
L136:
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
		goto L137
	}
L137:
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
		goto L138
	}
L138:
	;
	goto L1
L139:
	;
	if v683 == int32(0) {
		goto L140
	} else {
		goto L141
	}
L140:
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
		goto L143
	}
L141:
	;
	goto L142
L142:
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
		goto L147
	}
L143:
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
		goto L144
	}
L144:
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
		goto L145
	}
L145:
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
		goto L146
	}
L146:
	;
	goto L1
L147:
	;
	goto L130
L148:
	;
	__phi734 = v724
	__phi736 = v725
	v734 = __phi734
	v736 = __phi736
	goto L149
L149:
	;
	v744 = *(*int32)(unsafe.Add(mBase, uint32(v736)))
	v745 = *(*int32)(unsafe.Add(mBase, uint32(v736)+24))
	v753 = v744
	v757 = v77
	goto L152
L150:
	;
	goto L18
L151:
	;
	if v762 != 0 {
		goto L167
	} else {
		goto L168
	}
L152:
	;
	v762 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v757))))
	v763 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v753))))
	if v763 == int32(0) {
		goto L151
	} else {
		goto L154
	}
L153:
	;
	if v745 == int32(0) {
		goto L18
	} else {
		goto L166
	}
L154:
	;
	if v762 == int32(0) {
		goto L155
	} else {
		goto L156
	}
L155:
	;
	if v745 == int32(0) {
		goto L18
	} else {
		goto L158
	}
L156:
	;
	goto L157
L157:
	;
	v770 = int32(1)
	if base.Ui32((v762-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L159
	} else {
		goto L160
	}
L158:
	;
	__phi734 = v736
	__phi736 = v745
	v734 = __phi734
	v736 = __phi736
	goto L149
L159:
	;
	v782 = v762 | int32(32)
	goto L161
L160:
	;
	v782 = v762
	goto L161
L161:
	;
	v783 = int32(255)
	if base.Ui32((v763-int32(65))&v783) < base.Ui32(int32(26)) {
		goto L162
	} else {
		goto L163
	}
L162:
	;
	v793 = v763 | int32(32)
	goto L164
L163:
	;
	v793 = v763
	goto L164
L164:
	;
	if v782&v783 == v793 {
		v753 = v753 + v770
		v757 = v757 + v770
		goto L152
	} else {
		goto L165
	}
L165:
	;
	goto L153
L166:
	;
	__phi734 = v736
	__phi736 = v745
	v734 = __phi734
	v736 = __phi736
	goto L149
L167:
	;
	if v745 == int32(0) {
		goto L18
	} else {
		goto L170
	}
L168:
	;
	goto L169
L169:
	;
	if v734 != 0 {
		goto L172
	} else {
		goto L173
	}
L170:
	;
	__phi734 = v736
	__phi736 = v745
	v734 = __phi734
	v736 = __phi736
	goto L149
L171:
	;
	if v745 == int32(0) {
		goto L175
	} else {
		goto L176
	}
L172:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v734)+24)) = v745
	goto L171
L173:
	;
	goto L174
L174:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2560)) = v745
	goto L171
L175:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2556)) = v734
	goto L177
L176:
	;
	goto L177
L177:
	;
	v804 = *(*int32)(unsafe.Add(mBase, uint32(v736)))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v111
	F_pfree(m, v804)
	mBase = m.M
	v809 = m.ExcPending
	if v809 != 0 {
		goto L6
	} else {
		goto L178
	}
L178:
	;
	v810 = *(*int32)(unsafe.Add(mBase, uint32(v736)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v111
	F_pfree(m, v810)
	mBase = m.M
	v815 = m.ExcPending
	if v815 != 0 {
		goto L6
	} else {
		goto L179
	}
L179:
	;
	v816 = *(*int32)(unsafe.Add(mBase, uint32(v736)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v111
	F_pfree(m, v816)
	mBase = m.M
	v821 = m.ExcPending
	if v821 != 0 {
		goto L6
	} else {
		goto L180
	}
L180:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v111
	F_pfree(m, v736)
	mBase = m.M
	v826 = m.ExcPending
	if v826 != 0 {
		goto L6
	} else {
		goto L181
	}
L181:
	;
	if v745 != 0 {
		__phi736 = v745
		v736 = __phi736
		goto L149
	} else {
		goto L182
	}
L182:
	;
	goto L150
L183:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v26
	*(*int32)(unsafe.Add(mBase, uint32(v23)+196)) = int32(_a_F_AlterSystemSetConfigFile_19)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+192)) = v833
	v852 = F_pg_snprintf(m, v23+int32(496), int32(1024), int32(_a_F_AlterSystemSetConfigFile_20), v23+int32(192))
	mBase = m.M
	v853 = m.ExcPending
	if v853 != 0 {
		goto L6
	} else {
		goto L184
	}
L184:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v26
	v858 = *(*int32)(unsafe.Add(mBase, _c_F_AlterSystemSetConfigFile[4]))
	v862 = F_LWLockAcquire(m, v858+int32(_a_F_AlterSystemSetConfigFile_21), int32(0))
	mBase = m.M
	v863 = m.ExcPending
	if v863 != 0 {
		goto L6
	} else {
		goto L185
	}
L185:
	;
	v927 = v26
	goto L17
L186:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v111
	v886 = F_palloc(m, int32(28))
	mBase = m.M
	v887 = m.ExcPending
	if v887 != 0 {
		goto L6
	} else {
		goto L187
	}
L187:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v111
	v891 = F_pstrdup(m, v77)
	mBase = m.M
	v892 = m.ExcPending
	if v892 != 0 {
		goto L6
	} else {
		goto L188
	}
L188:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v886))) = v891
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v111
	v897 = F_pstrdup(m, v112)
	mBase = m.M
	v898 = m.ExcPending
	if v898 != 0 {
		goto L6
	} else {
		goto L189
	}
L189:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v886)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v886)+4)) = v897
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v111
	v906 = F_pstrdup(m, int32(_a_F_AlterSystemSetConfigFile_27))
	mBase = m.M
	v907 = m.ExcPending
	if v907 != 0 {
		goto L6
	} else {
		goto L190
	}
L190:
	;
	v908 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v886)+24)) = v908
	*(*uint16)(unsafe.Add(mBase, uint32(v886)+20)) = uint16(v908)
	*(*int32)(unsafe.Add(mBase, uint32(v886)+16)) = v908
	*(*int32)(unsafe.Add(mBase, uint32(v886)+12)) = v906
	v915 = *(*int32)(unsafe.Add(mBase, uint32(v23)+2560))
	if v915 == v908 {
		goto L191
	} else {
		goto L192
	}
L191:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2560)) = v886
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2556)) = v886
	v927 = v111
	goto L17
L192:
	;
	goto L193
L193:
	;
	v920 = *(*int32)(unsafe.Add(mBase, uint32(v23)+2556))
	*(*int32)(unsafe.Add(mBase, uint32(v920)+24)) = v886
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2556)) = v886
	v927 = v111
	goto L17
L194:
	;
	v941 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	v942 = *(*int32)(unsafe.Add(mBase, uint32(v941)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v927
	F_RunObjectPostAlterHookStr(m, v77, int32(_a_F_AlterSystemSetConfigFile_28), v942)
	mBase = m.M
	v948 = m.ExcPending
	if v948 != 0 {
		goto L6
	} else {
		goto L197
	}
L195:
	;
	goto L196
L196:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v927
	v954 = v23 + int32(496)
	v956 = F_BasicOpenFile(m, v954, int32(578))
	mBase = m.M
	v957 = m.ExcPending
	if v957 != 0 {
		goto L6
	} else {
		goto L198
	}
L197:
	;
	goto L196
L198:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2552)) = v956
	if v956 < int32(0) {
		goto L199
	} else {
		goto L200
	}
L199:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v927
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v968 = m.ExcPending
	if v968 != 0 {
		goto L6
	} else {
		goto L202
	}
L200:
	;
	goto L201
L201:
	;
	v993 = *(*int32)(unsafe.Add(mBase, _c_F_AlterSystemSetConfigFile[5]))
	v995 = *(*int32)(unsafe.Add(mBase, _c_F_AlterSystemSetConfigFile[6]))
	goto L206
L202:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v927
	F_errcode_for_file_access(m)
	mBase = m.M
	v973 = m.ExcPending
	if v973 != 0 {
		goto L6
	} else {
		goto L203
	}
L203:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v927
	*(*int32)(unsafe.Add(mBase, uint32(v23)+80)) = v954
	F_errmsg(m, int32(_a_F_AlterSystemSetConfigFile_23), v23+int32(80))
	mBase = m.M
	v982 = m.ExcPending
	if v982 != 0 {
		goto L6
	} else {
		goto L204
	}
L204:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v927
	F_errfinish(m, int32(_a_F_AlterSystemSetConfigFile_1), int32(_a_F_AlterSystemSetConfigFile_29), int32(_a_F_AlterSystemSetConfigFile_3))
	mBase = m.M
	v990 = m.ExcPending
	if v990 != 0 {
		goto L6
	} else {
		goto L205
	}
L205:
	;
	goto L1
L206:
	;
	v997 = v23 + int32(224)
	*(*int32)(unsafe.Add(mBase, uint32(v997)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v997))) = v23 + int32(220)
	goto L209
L207:
	;
	v1005 = v995
	v1006 = v993
	v1007 = v927
	v1012 = int32(0)
	goto L9
L209:
	;
	goto L207
L210:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v1005
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v1006
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v1007
	v1426 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_AlterSystemSetConfigFile[7])))
	if v1426 != int32(1) {
		v1440 = int32(0)
		goto L280
	} else {
		goto L281
	}
L211:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_AlterSystemSetConfigFile[6])) = v23 + int32(224)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v1006
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v1005
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v1007
	v1028 = *(*int32)(unsafe.Add(mBase, uint32(v23)+2552))
	v1029 = *(*int32)(unsafe.Add(mBase, uint32(v23)+2560))
	v1031 = v23 + int32(2564)
	F_initStringInfo(m, v1031)
	mBase = m.M
	v1033 = m.ExcPending
	if v1033 != 0 {
		goto L6
	} else {
		goto L214
	}
L212:
	;
	goto L213
L213:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_AlterSystemSetConfigFile[5])) = v1006
	*(*int32)(unsafe.Add(mBase, _c_F_AlterSystemSetConfigFile[6])) = v1005
	v1386 = *(*int32)(unsafe.Add(mBase, uint32(v23)+2552))
	if int32(0) <= v1386 {
		goto L275
	} else {
		goto L276
	}
L214:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v1005
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v1006
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v1007
	F_appendStringInfoString(m, v1031, int32(_a_F_AlterSystemSetConfigFile_30))
	mBase = m.M
	v1039 = m.ExcPending
	if v1039 != 0 {
		goto L6
	} else {
		goto L215
	}
L215:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v1005
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v1006
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v1007
	F_appendStringInfoString(m, v1031, int32(_a_F_AlterSystemSetConfigFile_31))
	mBase = m.M
	v1045 = m.ExcPending
	if v1045 != 0 {
		goto L6
	} else {
		goto L216
	}
L216:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v1005
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v1006
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v1007
	*(*int32)(unsafe.Add(mBase, _c_F_AlterSystemSetConfigFile[8])) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v1006
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v1005
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v1007
	v1055 = *(*int32)(unsafe.Add(mBase, uint32(v23)+2564))
	v1056 = *(*int32)(unsafe.Add(mBase, uint32(v23)+2568))
	v1057 = F_write(m, v1028, v1055, v1056)
	mBase = m.M
	v1058 = *(*int32)(unsafe.Add(mBase, uint32(v23)+2568))
	if v1057 == v1058 {
		goto L218
	} else {
		goto L219
	}
L217:
	;
	v1107 = v1029
	goto L229
L218:
	;
	if v1029 != 0 {
		goto L217
	} else {
		goto L221
	}
L219:
	;
	goto L220
L220:
	;
	v1061 = *(*int32)(unsafe.Add(mBase, _c_F_AlterSystemSetConfigFile[8]))
	if v1061 == int32(0) {
		goto L222
	} else {
		goto L223
	}
L221:
	;
	goto L210
L222:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_AlterSystemSetConfigFile[8])) = int32(51)
	goto L224
L223:
	;
	goto L224
L224:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v1005
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v1006
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v1007
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1073 = m.ExcPending
	if v1073 != 0 {
		goto L6
	} else {
		goto L225
	}
L225:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v1005
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v1006
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v1007
	F_errcode_for_file_access(m)
	mBase = m.M
	v1078 = m.ExcPending
	if v1078 != 0 {
		goto L6
	} else {
		goto L226
	}
L226:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v1005
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v1006
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v1007
	*(*int32)(unsafe.Add(mBase, uint32(v23)+32)) = v23 + int32(496)
	F_errmsg(m, int32(_a_F_AlterSystemSetConfigFile_32), v23+int32(32))
	mBase = m.M
	v1089 = m.ExcPending
	if v1089 != 0 {
		goto L6
	} else {
		goto L227
	}
L227:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v1005
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v1006
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v1007
	F_errfinish(m, int32(_a_F_AlterSystemSetConfigFile_1), int32(_a_F_AlterSystemSetConfigFile_33), int32(_a_F_AlterSystemSetConfigFile_34))
	mBase = m.M
	v1097 = m.ExcPending
	if v1097 != 0 {
		goto L6
	} else {
		goto L228
	}
L228:
	;
	goto L1
L229:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v1005
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v1006
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v1007
	v1118 = v23 + int32(2564)
	v1119 = *(*int32)(unsafe.Add(mBase, uint32(v1118)))
	v1120 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1119))) = uint8(v1120)
	*(*int32)(unsafe.Add(mBase, uint32(v1118)+12)) = v1120
	*(*int32)(unsafe.Add(mBase, uint32(v1118)+4)) = v1120
	goto L231
L230:
	;
	v1345 = *(*int32)(unsafe.Add(mBase, _c_F_AlterSystemSetConfigFile[8]))
	if v1345 == int32(0) {
		goto L268
	} else {
		goto L269
	}
L231:
	;
	v1126 = *(*int32)(unsafe.Add(mBase, uint32(v1107)))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v1005
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v1006
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v1007
	F_appendStringInfoString(m, v1118, v1126)
	mBase = m.M
	v1131 = m.ExcPending
	if v1131 != 0 {
		goto L6
	} else {
		goto L232
	}
L232:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v1005
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v1006
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v1007
	F_appendStringInfoString(m, v1118, int32(_a_F_AlterSystemSetConfigFile_35))
	mBase = m.M
	v1137 = m.ExcPending
	if v1137 != 0 {
		goto L6
	} else {
		goto L233
	}
L233:
	;
	v1138 = *(*int32)(unsafe.Add(mBase, uint32(v1107)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v1005
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v1006
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v1007
	v1142 = int32(0)
	v1145 = F_strlen(m, v1138)
	mBase = m.M
	v1146 = int32(1)
	v1150 = F_emscripten_builtin_malloc(m, v1145<<(uint(v1146)%32)|v1146)
	mBase = m.M
	if v1150 != 0 {
		goto L234
	} else {
		goto L235
	}
L234:
	;
	if v1145 <= int32(0) {
		v1257 = v1142
		goto L237
	} else {
		goto L238
	}
L235:
	;
	goto L236
L236:
	;
	if v1150 == int32(0) {
		goto L255
	} else {
		goto L256
	}
L237:
	;
	v1268 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1150+v1257))) = uint8(v1268)
	goto L236
L238:
	;
	if v1145 != int32(1) {
		goto L239
	} else {
		goto L240
	}
L239:
	;
	v1165 = v1142
	v1169 = v1142
	v1170 = v1142
	goto L242
L240:
	;
	v1223 = v1142
	v1228 = v1142
	goto L241
L241:
	;
	v1234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1228+v1138))))
	if base.B2i32(v1234 != int32(92))&base.B2i32(v1234 != int32(39)) == int32(0) {
		goto L252
	} else {
		goto L253
	}
L242:
	;
	v1175 = v1170 + v1138
	v1176 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1175))))
	if base.B2i32(v1176 != int32(92))&base.B2i32(v1176 != int32(39)) == int32(0) {
		goto L244
	} else {
		goto L245
	}
L243:
	;
	if v1145&int32(1) == int32(0) {
		v1257 = v1211
		goto L237
	} else {
		goto L251
	}
L244:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1150+v1165))) = uint8(v1176)
	v1188 = v1165 + int32(1)
	goto L246
L245:
	;
	v1188 = v1165
	goto L246
L246:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1150+v1188))) = uint8(v1176)
	v1192 = v1188 + int32(1)
	v1193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1175)+1)))
	if base.B2i32(v1193 != int32(92))&base.B2i32(v1193 != int32(39)) == int32(0) {
		goto L247
	} else {
		goto L248
	}
L247:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1150+v1192))) = uint8(v1193)
	v1205 = v1188 + int32(2)
	goto L249
L248:
	;
	v1205 = v1192
	goto L249
L249:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1150+v1205))) = uint8(v1193)
	v1208 = int32(2)
	v1209 = v1170 + v1208
	v1211 = v1205 + int32(1)
	v1213 = v1169 + v1208
	if v1213 != v1145&int32(2147483646) {
		v1165 = v1211
		v1169 = v1213
		v1170 = v1209
		goto L242
	} else {
		goto L250
	}
L250:
	;
	goto L243
L251:
	;
	v1223 = v1211
	v1228 = v1209
	goto L241
L252:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1150+v1223))) = uint8(v1234)
	v1246 = v1223 + int32(1)
	goto L254
L253:
	;
	v1246 = v1223
	goto L254
L254:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1150+v1246))) = uint8(v1234)
	v1257 = v1246 + int32(1)
	goto L237
L255:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v1005
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v1006
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v1007
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1294 = m.ExcPending
	if v1294 != 0 {
		goto L6
	} else {
		goto L258
	}
L256:
	;
	goto L257
L257:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v1005
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v1006
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v1007
	v1320 = v23 + int32(2564)
	F_appendStringInfoString(m, v1320, v1150)
	mBase = m.M
	v1322 = m.ExcPending
	if v1322 != 0 {
		goto L6
	} else {
		goto L262
	}
L258:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v1005
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v1006
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v1007
	F_errcode(m, int32(_a_F_AlterSystemSetConfigFile_36))
	mBase = m.M
	v1300 = m.ExcPending
	if v1300 != 0 {
		goto L6
	} else {
		goto L259
	}
L259:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v1005
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v1006
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v1007
	F_errmsg(m, int32(_a_F_AlterSystemSetConfigFile_37), int32(0))
	mBase = m.M
	v1307 = m.ExcPending
	if v1307 != 0 {
		goto L6
	} else {
		goto L260
	}
L260:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v1005
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v1006
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v1007
	F_errfinish(m, int32(_a_F_AlterSystemSetConfigFile_1), int32(_a_F_AlterSystemSetConfigFile_38), int32(_a_F_AlterSystemSetConfigFile_34))
	mBase = m.M
	v1315 = m.ExcPending
	if v1315 != 0 {
		goto L6
	} else {
		goto L261
	}
L261:
	;
	goto L1
L262:
	;
	F_emscripten_builtin_free(m, v1150)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v1005
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v1006
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v1007
	F_appendStringInfoString(m, v1320, int32(_a_F_AlterSystemSetConfigFile_39))
	mBase = m.M
	v1329 = m.ExcPending
	if v1329 != 0 {
		goto L6
	} else {
		goto L263
	}
L263:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_AlterSystemSetConfigFile[8])) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v1006
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v1005
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v1007
	v1336 = *(*int32)(unsafe.Add(mBase, uint32(v23)+2564))
	v1337 = *(*int32)(unsafe.Add(mBase, uint32(v23)+2568))
	v1338 = F_write(m, v1028, v1336, v1337)
	mBase = m.M
	v1339 = *(*int32)(unsafe.Add(mBase, uint32(v23)+2568))
	if v1338 == v1339 {
		goto L264
	} else {
		goto L265
	}
L264:
	;
	v1341 = *(*int32)(unsafe.Add(mBase, uint32(v1107)+24))
	if v1341 == int32(0) {
		goto L210
	} else {
		goto L267
	}
L265:
	;
	goto L266
L266:
	;
	goto L230
L267:
	;
	v1107 = v1341
	goto L229
L268:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_AlterSystemSetConfigFile[8])) = int32(51)
	goto L270
L269:
	;
	goto L270
L270:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v1005
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v1006
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v1007
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1357 = m.ExcPending
	if v1357 != 0 {
		goto L6
	} else {
		goto L271
	}
L271:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v1005
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v1006
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v1007
	F_errcode_for_file_access(m)
	mBase = m.M
	v1362 = m.ExcPending
	if v1362 != 0 {
		goto L6
	} else {
		goto L272
	}
L272:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v1005
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v1006
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v1007
	*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = v23 + int32(496)
	F_errmsg(m, int32(_a_F_AlterSystemSetConfigFile_32), v23+int32(16))
	mBase = m.M
	v1373 = m.ExcPending
	if v1373 != 0 {
		goto L6
	} else {
		goto L273
	}
L273:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v1005
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v1006
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v1007
	F_errfinish(m, int32(_a_F_AlterSystemSetConfigFile_1), int32(_a_F_AlterSystemSetConfigFile_40), int32(_a_F_AlterSystemSetConfigFile_34))
	mBase = m.M
	v1381 = m.ExcPending
	if v1381 != 0 {
		goto L6
	} else {
		goto L274
	}
L274:
	;
	goto L1
L275:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v1006
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v1005
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v1007
	v1392 = *(*int32)(unsafe.Add(mBase, uint32(v23)+2552))
	v1393 = F_close(m, v1392)
	mBase = m.M
	goto L277
L276:
	;
	goto L277
L277:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v1005
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v1006
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v1007
	v1399 = F_unlink(m, v23+int32(496))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v1005
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v1006
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v1007
	F_pg_re_throw(m)
	mBase = m.M
	v1404 = m.ExcPending
	if v1404 != 0 {
		goto L6
	} else {
		goto L278
	}
L278:
	;
	goto L1
L279:
	;
	if v1440 != 0 {
		goto L286
	} else {
		goto L287
	}
L280:
	;
	goto L279
L281:
	;
	goto L282
L282:
	;
	v1431 = F_fsync(m, v1028)
	mBase = m.M
	if v1431 != int32(-1) {
		v1440 = v1431
		goto L280
	} else {
		goto L284
	}
L283:
	;
	v1440 = int32(-1)
	goto L280
L284:
	;
	v1435 = *(*int32)(unsafe.Add(mBase, _c_F_AlterSystemSetConfigFile[8]))
	if v1435 == int32(27) {
		goto L282
	} else {
		goto L285
	}
L285:
	;
	goto L283
L286:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v1005
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v1006
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v1007
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1447 = m.ExcPending
	if v1447 != 0 {
		goto L6
	} else {
		goto L289
	}
L287:
	;
	goto L288
L288:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v1006
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v1005
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v1007
	v1473 = *(*int32)(unsafe.Add(mBase, uint32(v23)+2564))
	F_pfree(m, v1473)
	mBase = m.M
	v1475 = m.ExcPending
	if v1475 != 0 {
		goto L6
	} else {
		goto L293
	}
L289:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v1005
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v1006
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v1007
	F_errcode_for_file_access(m)
	mBase = m.M
	v1452 = m.ExcPending
	if v1452 != 0 {
		goto L6
	} else {
		goto L290
	}
L290:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v1005
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v1006
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v1007
	*(*int32)(unsafe.Add(mBase, uint32(v23))) = v23 + int32(496)
	F_errmsg(m, int32(_a_F_AlterSystemSetConfigFile_41), v23)
	mBase = m.M
	v1461 = m.ExcPending
	if v1461 != 0 {
		goto L6
	} else {
		goto L291
	}
L291:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v1005
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v1006
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v1007
	F_errfinish(m, int32(_a_F_AlterSystemSetConfigFile_1), int32(_a_F_AlterSystemSetConfigFile_42), int32(_a_F_AlterSystemSetConfigFile_34))
	mBase = m.M
	v1469 = m.ExcPending
	if v1469 != 0 {
		goto L6
	} else {
		goto L292
	}
L292:
	;
	goto L1
L293:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v1006
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v1005
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v1007
	v1479 = *(*int32)(unsafe.Add(mBase, uint32(v23)+2552))
	v1480 = F_close(m, v1479)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v1006
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2552)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v1005
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v1007
	v1491 = F_durable_rename(m, v23+int32(496), v23+int32(1520), int32(21))
	mBase = m.M
	v1492 = m.ExcPending
	if v1492 != 0 {
		goto L6
	} else {
		goto L294
	}
L294:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_AlterSystemSetConfigFile[5])) = v1006
	*(*int32)(unsafe.Add(mBase, _c_F_AlterSystemSetConfigFile[6])) = v1005
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v1006
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v1005
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v1007
	v1500 = *(*int32)(unsafe.Add(mBase, uint32(v23)+2560))
	F_FreeConfigVariables(m, v1500)
	mBase = m.M
	v1502 = m.ExcPending
	if v1502 != 0 {
		goto L6
	} else {
		goto L295
	}
L295:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2580)) = v1006
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2584)) = v1005
	*(*int32)(unsafe.Add(mBase, uint32(v23)+2588)) = v1007
	v1507 = *(*int32)(unsafe.Add(mBase, _c_F_AlterSystemSetConfigFile[4]))
	F_LWLockRelease(m, v1507+int32(_a_F_AlterSystemSetConfigFile_21))
	mBase = m.M
	v1511 = m.ExcPending
	if v1511 != 0 {
		goto L6
	} else {
		goto L296
	}
L296:
	;
	goto L5
L297:
	;
	v1533 = int32(v1529)
	m.G0 = v23
	v1535 = *(*int32)(unsafe.Add(mBase, uint32(v1533)+4))
	v1536 = *(*int32)(unsafe.Add(mBase, uint32(v1533)))
	v1539 = *(*int32)(unsafe.Add(mBase, uint32(v1536)))
	if v23+int32(220) == v1539 {
		goto L300
	} else {
		goto L301
	}
L298:
	;
	m.ExcPending = 1
	goto L306
L299:
	;
	if v1543 != 0 {
		goto L303
	} else {
		goto L304
	}
L300:
	;
	v1541 = *(*int32)(unsafe.Add(mBase, uint32(v1536)+4))
	v1543 = v1541
	goto L302
L301:
	;
	v1543 = int32(0)
	goto L302
L302:
	;
	goto L299
L303:
	;
	v1544 = *(*int32)(unsafe.Add(mBase, uint32(v23)+2588))
	v1545 = *(*int32)(unsafe.Add(mBase, uint32(v23)+2584))
	v1546 = *(*int32)(unsafe.Add(mBase, uint32(v23)+2580))
	v24 = v1545
	v25 = v1546
	v26 = v1544
	v30 = v1543
	v31 = v1535
	goto L2
L304:
	;
	goto L305
L305:
	;
	F___wasm_longjmp(m, v1536, v1535)
	mBase = m.M
	v1548 = m.ExcPending
	if v1548 != 0 {
		goto L306
	} else {
		goto L307
	}
L306:
	;
	return
L307:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_ExecAlterExtensionContentsRecurse(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v31 int64
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
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v61 int64
	_ = v61
	var v62 int32
	_ = v62
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
	var v75 int32
	_ = v75
	var v83 int32
	_ = v83
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v107 int64
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v145 int64
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v185 int64
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v214 int64
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v249 int32
	_ = v249
	var v255 int32
	_ = v255
	var v260 int32
	_ = v260
	var v264 int32
	_ = v264
	var v270 int32
	_ = v270
	var v275 int32
	_ = v275
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v286 int32
	_ = v286
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v327 int32
	_ = v327
	var v330 int64
	_ = v330
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v339 int32
	_ = v339
	var v342 int32
	_ = v342
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v359 int32
	_ = v359
	var v373 int32
	_ = v373
	var v376 int32
	_ = v376
	var v378 int64
	_ = v378
	var v400 int32
	_ = v400
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v412 int32
	_ = v412
	var v416 int32
	_ = v416
	var v421 int32
	_ = v421
	var v429 int32
	_ = v429
	var v433 int32
	_ = v433
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v441 int64
	_ = v441
	var v444 int32
	_ = v444
	var v450 int32
	_ = v450
	var v471 int32
	_ = v471
	var v483 int32
	_ = v483
	var v485 int32
	_ = v485
	var v487 int32
	_ = v487
	var v488 int64
	_ = v488
	var v491 int32
	_ = v491
	var v492 int64
	_ = v492
	var v495 int32
	_ = v495
	var v496 int64
	_ = v496
	var v499 int32
	_ = v499
	var v503 int64
	_ = v503
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v544 int32
	_ = v544
	var v547 int64
	_ = v547
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v558 int32
	_ = v558
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v565 int32
	_ = v565
	var v569 int32
	_ = v569
	var v571 int32
	_ = v571
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v582 int32
	_ = v582
	var v589 int32
	_ = v589
	var v595 int32
	_ = v595
	var v599 int32
	_ = v599
	var v606 int32
	_ = v606
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
	var v614 int64
	_ = v614
	var v617 int32
	_ = v617
	var v624 int32
	_ = v624
	var v642 int32
	_ = v642
	var v653 int32
	_ = v653
	var v655 int32
	_ = v655
	var v657 int32
	_ = v657
	var v658 int64
	_ = v658
	var v661 int32
	_ = v661
	var v662 int64
	_ = v662
	var v665 int32
	_ = v665
	var v666 int64
	_ = v666
	var v669 int32
	_ = v669
	var v673 int64
	_ = v673
	var v693 int32
	_ = v693
	var v694 int32
	_ = v694
	var v713 int32
	_ = v713
	var v720 int32
	_ = v720
	var v721 int32
	_ = v721
	var v725 int32
	_ = v725
	var v743 int32
	_ = v743
	var v746 int32
	_ = v746
	var v747 int32
	_ = v747
	var v752 int32
	_ = v752
	var v764 int32
	_ = v764
	var v765 int32
	_ = v765
	var v767 int32
	_ = v767
	var v772 int64
	_ = v772
	var v773 int32
	_ = v773
	var v774 int32
	_ = v774
	var v777 int32
	_ = v777
	var v778 int32
	_ = v778
	var v779 int32
	_ = v779
	var v780 int32
	_ = v780
	var v784 int32
	_ = v784
	var v792 int32
	_ = v792
	var v806 int32
	_ = v806
	var v807 int32
	_ = v807
	var v811 int32
	_ = v811
	var v813 int32
	_ = v813
	var v816 int32
	_ = v816
	var v835 int32
	_ = v835
	var v852 int32
	_ = v852
	var v855 int32
	_ = v855
	var v859 int32
	_ = v859
	var v863 int32
	_ = v863
	var v868 int32
	_ = v868
	var v870 int32
	_ = v870
	var v906 int32
	_ = v906
	var v913 int32
	_ = v913
	var v914 int32
	_ = v914
	var v915 int32
	_ = v915
	var v917 int32
	_ = v917
	var v919 int64
	_ = v919
	var v921 int64
	_ = v921
	var v923 int32
	_ = v923
	var v930 int32
	_ = v930
	var v931 int32
	_ = v931
	var v932 int32
	_ = v932
	var v933 int32
	_ = v933
	var v934 int32
	_ = v934
	var v935 int32
	_ = v935
	var v936 int32
	_ = v936
	var v940 int32
	_ = v940
	var v942 int64
	_ = v942
	var v944 int64
	_ = v944
	var v946 int32
	_ = v946
	var v953 int32
	_ = v953
	var v955 int32
	_ = v955
	var v957 int32
	_ = v957
	var v964 int32
	_ = v964
	var v965 int32
	_ = v965
	var v966 int32
	_ = v966
	var v970 int32
	_ = v970
	var v972 int64
	_ = v972
	var v974 int64
	_ = v974
	var v976 int32
	_ = v976
	var v983 int32
	_ = v983
	var v991 int32
	_ = v991
	var v994 int32
	_ = v994
	var v996 int32
	_ = v996
	var v997 int32
	_ = v997
	var v998 int32
	_ = v998
	var v999 int32
	_ = v999
	var v1006 int32
	_ = v1006
	var v1011 int32
	_ = v1011
	var v1015 int32
	_ = v1015
	var v1018 int32
	_ = v1018
	var v1019 int32
	_ = v1019
	var v1020 int32
	_ = v1020
	var v1021 int32
	_ = v1021
	var v1022 int32
	_ = v1022
	var v1027 int32
	_ = v1027
	var v1032 int32
	_ = v1032
	var v1036 int32
	_ = v1036
	var v1039 int32
	_ = v1039
	var v1041 int32
	_ = v1041
	var v1042 int32
	_ = v1042
	var v1043 int32
	_ = v1043
	var v1050 int32
	_ = v1050
	var v1055 int32
	_ = v1055
	var v1059 int32
	_ = v1059
	var v1063 int32
	_ = v1063
	var v1068 int32
	_ = v1068
	var v1072 int32
	_ = v1072
	var v1078 int32
	_ = v1078
	var v1083 int32
	_ = v1083
	var v1088 int32
	_ = v1088
	var v1092 int32
	_ = v1092
	var v1097 int32
	_ = v1097
	var v1101 int32
	_ = v1101
	var v1105 int32
	_ = v1105
	var v1110 int32
	_ = v1110
	var v1114 int32
	_ = v1114
	var v1118 int32
	_ = v1118
	var v1123 int32
	_ = v1123
	var v1127 int32
	_ = v1127
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
	var v1145 int32
	_ = v1145
	var v1146 int32
	_ = v1146
	var v1152 int32
	_ = v1152
	var v1157 int32
	_ = v1157
	v17 = m.G0
	v19 = v17 - int32(336)
	m.G0 = v19
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v23 = F_getExtensionOfObject(m, v21, v22)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if int32(0) < v25 {
		goto L14
	} else {
		goto L15
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1140 = m.ExcPending
	if v1140 != 0 {
		goto L1
	} else {
		goto L257
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1127 = m.ExcPending
	if v1127 != 0 {
		goto L1
	} else {
		goto L254
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1114 = m.ExcPending
	if v1114 != 0 {
		goto L1
	} else {
		goto L251
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1101 = m.ExcPending
	if v1101 != 0 {
		goto L1
	} else {
		goto L248
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1088 = m.ExcPending
	if v1088 != 0 {
		goto L1
	} else {
		goto L245
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1072 = m.ExcPending
	if v1072 != 0 {
		goto L1
	} else {
		goto L242
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1059 = m.ExcPending
	if v1059 != 0 {
		goto L1
	} else {
		goto L239
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1036 = m.ExcPending
	if v1036 != 0 {
		goto L1
	} else {
		goto L234
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1015 = m.ExcPending
	if v1015 != 0 {
		goto L1
	} else {
		goto L229
	}
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v991 = m.ExcPending
	if v991 != 0 {
		goto L1
	} else {
		goto L223
	}
L13:
	;
	v906 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v906 == int32(1247) {
		goto L204
	} else {
		goto L205
	}
L14:
	;
	if v23 != 0 {
		goto L12
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v292 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v23 != v292 {
		goto L10
	} else {
		goto L102
	}
L17:
	;
	if v21 == int32(2615) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v31 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l1)+4)))
	v32 = F_SearchSysCache1(m, int32(28), v31)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
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
	v49 = m.ExcPending
	if v49 != 0 {
		goto L1
	} else {
		goto L28
	}
L21:
	;
	if v43 == v22 {
		goto L11
	} else {
		goto L27
	}
L22:
	;
	if v32 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v43 = int32(0)
	goto L21
L24:
	;
	goto L25
L25:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v32)+16))
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+22)))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v37+v38)+72))
	F_ReleaseCatCache(m, v32)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	v43 = v40
	goto L21
L27:
	;
	goto L20
L28:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v52 = m.G0
	v54 = v52 - int32(112)
	m.G0 = v54
	if v51 != int32(2613) {
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
	v279 = m.ExcPending
	if v279 != 0 {
		goto L1
	} else {
		goto L98
	}
L31:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L1
	} else {
		goto L95
	}
L32:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L1
	} else {
		goto L92
	}
L33:
	;
	m.G0 = v54 + int32(112)
	goto L29
L34:
	;
	v199 = F_get_object_attnum_acl(m, v51)
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L1
	} else {
		goto L79
	}
L35:
	;
	if v51 != int32(1259) {
		goto L34
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	v161 = F_table_open(m, int32(2995), int32(3))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L1
	} else {
		goto L67
	}
L38:
	;
	v61 = base.I64_extend_i32_u(v50)
	v62 = F_SearchSysCache1(m, int32(57), v61)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	if v62 == int32(0) {
		goto L32
	} else {
		goto L40
	}
L40:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v62)+16))
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+22)))
	v68 = v66 + v67
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68)+119)))
	switch v69 - int32(73) {
	case 0, 26, 32:
		goto L43
	default:
		goto L42
	case 10:
		goto L41
	}
L41:
	;
	v145 = F_SysCacheGetAttr(m, int32(57), v62, int32(32), v54+int32(48))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L1
	} else {
		goto L60
	}
L42:
	;
	v75 = int32(*(*int16)(unsafe.Add(mBase, uint32(v68)+120)))
	if v75 <= int32(0) {
		goto L41
	} else {
		goto L45
	}
L43:
	;
	F_ReleaseCatCache(m, v62)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	goto L33
L45:
	;
	v83 = int32(1)
	goto L46
L46:
	;
	v97 = F_SearchSysCache2(m, int32(7), v61, base.I64_extend16_s(base.I64_extend_i32_u(v83)))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L1
	} else {
		goto L48
	}
L47:
	;
	goto L41
L48:
	;
	if v97 != 0 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v97)+16))
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v99)+22)))
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v99+v100)+91)))
	if v102 != 0 {
		goto L52
	} else {
		goto L53
	}
L50:
	;
	goto L51
L51:
	;
	v123 = base.I32_extend16_s(v83 + int32(1))
	if v123 <= v75 {
		v83 = v123
		goto L46
	} else {
		goto L59
	}
L52:
	;
	F_ReleaseCatCache(m, v97)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L1
	} else {
		goto L58
	}
L53:
	;
	v107 = F_SysCacheGetAttr(m, int32(7), v97, int32(22), v54+int32(48))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54)+48)))
	if v109 != 0 {
		goto L52
	} else {
		goto L55
	}
L55:
	;
	v112 = F_pg_detoast_datum(m, base.I32_wrap_i64(v107))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	F_recordExtensionInitPrivWorker(m, v50, int32(1259), v83, v112)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
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
	v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54)+48)))
	if v147 == int32(0) {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v153 = F_pg_detoast_datum(m, base.I32_wrap_i64(v145))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L1
	} else {
		goto L64
	}
L62:
	;
	goto L63
L63:
	;
	F_ReleaseCatCache(m, v62)
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L1
	} else {
		goto L66
	}
L64:
	;
	F_recordExtensionInitPrivWorker(m, v50, int32(1259), int32(0), v153)
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
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
	v164 = v54 + int32(48)
	F_ScanKeyInit(m, v164, int32(1), int32(3), int32(184), base.I64_extend_i32_u(v50))
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L1
	} else {
		goto L68
	}
L68:
	;
	v172 = int32(1)
	v175 = F_systable_beginscan(m, v161, int32(2996), v172, int32(0), v172, v164)
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L1
	} else {
		goto L69
	}
L69:
	;
	v177 = F_systable_getnext(m, v175)
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L1
	} else {
		goto L70
	}
L70:
	;
	if v177 == int32(0) {
		goto L31
	} else {
		goto L71
	}
L71:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v161)+52))
	v185 = F_heap_getattr_2(m, v177, int32(3), v182, v54+int32(111))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	v187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54)+111)))
	if v187 == int32(0) {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v193 = F_pg_detoast_datum(m, base.I32_wrap_i64(v185))
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L1
	} else {
		goto L76
	}
L74:
	;
	goto L75
L75:
	;
	F_systable_endscan(m, v175)
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L1
	} else {
		goto L78
	}
L76:
	;
	F_recordExtensionInitPrivWorker(m, v50, int32(2613), int32(0), v193)
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
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
	if v199 == int32(0) {
		goto L33
	} else {
		goto L80
	}
L80:
	;
	v203 = F_get_object_catcache_oid(m, v51)
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	v206 = F_SearchSysCache1(m, v203, base.I64_extend_i32_u(v50))
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L1
	} else {
		goto L82
	}
L82:
	;
	if v206 == int32(0) {
		goto L30
	} else {
		goto L83
	}
L83:
	;
	v210 = F_get_object_attnum_acl(m, v51)
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L1
	} else {
		goto L84
	}
L84:
	;
	v214 = F_SysCacheGetAttr(m, v203, v206, v210, v54+int32(48))
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	v216 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54)+48)))
	if v216 == int32(0) {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v221 = F_pg_detoast_datum(m, base.I32_wrap_i64(v214))
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L1
	} else {
		goto L89
	}
L87:
	;
	goto L88
L88:
	;
	F_ReleaseCatCache(m, v206)
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L1
	} else {
		goto L91
	}
L89:
	;
	F_recordExtensionInitPrivWorker(m, v50, v51, int32(0), v221)
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v54)+16)) = v50
	F_errmsg_internal(m, int32(_a_F_ExecAlterExtensionContentsRecurse_0), v54+int32(16))
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L1
	} else {
		goto L93
	}
L93:
	;
	F_errfinish(m, int32(_a_F_ExecAlterExtensionContentsRecurse_1), int32(_a_F_ExecAlterExtensionContentsRecurse_2), int32(_a_F_ExecAlterExtensionContentsRecurse_3))
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v54)+32)) = v50
	F_errmsg_internal(m, int32(_a_F_ExecAlterExtensionContentsRecurse_4), v54+int32(32))
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L1
	} else {
		goto L96
	}
L96:
	;
	F_errfinish(m, int32(_a_F_ExecAlterExtensionContentsRecurse_1), int32(_a_F_ExecAlterExtensionContentsRecurse_5), int32(_a_F_ExecAlterExtensionContentsRecurse_3))
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
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
	v280 = F_get_object_class_descr(m, v51)
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L1
	} else {
		goto L99
	}
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54)+4)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v54))) = v280
	F_errmsg_internal(m, int32(_a_F_ExecAlterExtensionContentsRecurse_6), v54)
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L1
	} else {
		goto L100
	}
L100:
	;
	F_errfinish(m, int32(_a_F_ExecAlterExtensionContentsRecurse_1), int32(_a_F_ExecAlterExtensionContentsRecurse_7), int32(_a_F_ExecAlterExtensionContentsRecurse_3))
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
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
	v296 = F_deleteDependencyRecordsForClass(m, v21, v22, int32(3079), int32(101))
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L1
	} else {
		goto L103
	}
L103:
	;
	if v296 != int32(1) {
		goto L9
	} else {
		goto L104
	}
L104:
	;
	v300 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v300 == int32(1259) {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	v303 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v306 = F_table_open(m, int32(3079), int32(3))
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L1
	} else {
		goto L108
	}
L106:
	;
	v752 = v300
	goto L107
L107:
	;
	v764 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v765 = m.G0
	v767 = v765 - int32(16)
	m.G0 = v767
	if v752 == int32(1259) {
		goto L180
	} else {
		goto L181
	}
L108:
	;
	v309 = v19 + int32(272)
	F_ScanKeyInit(m, v309, int32(1), int32(3), int32(184), base.I64_extend_i32_u(v23))
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L1
	} else {
		goto L109
	}
L109:
	;
	v317 = int32(1)
	v320 = F_systable_beginscan(m, v306, int32(3080), v317, int32(0), v317, v309)
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L1
	} else {
		goto L110
	}
L110:
	;
	v322 = F_systable_getnext(m, v320)
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L1
	} else {
		goto L111
	}
L111:
	;
	if v322 == int32(0) {
		goto L8
	} else {
		goto L112
	}
L112:
	;
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v306)+52))
	v330 = F_heap_getattr_7(m, v322, int32(7), v327, v19+int32(271))
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L1
	} else {
		goto L113
	}
L113:
	;
	v332 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+271)))
	if v332 != 0 {
		goto L114
	} else {
		goto L115
	}
L114:
	;
	F_systable_endscan(m, v320)
	mBase = m.M
	v743 = m.ExcPending
	if v743 != 0 {
		goto L1
	} else {
		goto L175
	}
L115:
	;
	v334 = F_pg_detoast_datum(m, base.I32_wrap_i64(v330))
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L1
	} else {
		goto L116
	}
L116:
	;
	v336 = *(*int32)(unsafe.Add(mBase, uint32(v334)+4))
	if v336 != int32(1) {
		goto L7
	} else {
		goto L117
	}
L117:
	;
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v334)+20))
	if v339 != int32(1) {
		goto L7
	} else {
		goto L118
	}
L118:
	;
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v334)+16))
	if v342 < int32(0) {
		goto L7
	} else {
		goto L119
	}
L119:
	;
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v334)+8))
	if v345 != 0 {
		goto L7
	} else {
		goto L120
	}
L120:
	;
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v334)+12))
	if v346 != int32(26) {
		goto L7
	} else {
		goto L121
	}
L121:
	;
	if v342 == int32(0) {
		goto L114
	} else {
		goto L122
	}
L122:
	;
	v359 = int32(0)
	goto L123
L123:
	;
	v373 = *(*int32)(unsafe.Add(mBase, uint32(v334+int32(24)+v359<<(uint(int32(2))%32))))
	if v303 != v373 {
		goto L125
	} else {
		goto L126
	}
L124:
	;
	v378 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v19)+248)) = v378
	*(*int64)(unsafe.Add(mBase, uint32(v19)+240)) = v378
	*(*int64)(unsafe.Add(mBase, uint32(v19)+232)) = v378
	*(*int64)(unsafe.Add(mBase, uint32(v19)+224)) = v378
	*(*int64)(unsafe.Add(mBase, uint32(v19)+216)) = v378
	*(*int64)(unsafe.Add(mBase, uint32(v19)+208)) = v378
	*(*int64)(unsafe.Add(mBase, uint32(v19)+200)) = v378
	*(*int64)(unsafe.Add(mBase, uint32(v19)+192)) = v378
	*(*int64)(unsafe.Add(mBase, uint32(v19)+184)) = v378
	*(*int64)(unsafe.Add(mBase, uint32(v19)+176)) = int64(72339069014638592)
	if v342 == int32(1) {
		goto L130
	} else {
		goto L131
	}
L125:
	;
	v376 = v359 + int32(1)
	if v376 != v342 {
		v359 = v376
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
	v544 = *(*int32)(unsafe.Add(mBase, uint32(v306)+52))
	v547 = F_heap_getattr_7(m, v322, int32(8), v544, v19+int32(271))
	mBase = m.M
	v548 = m.ExcPending
	if v548 != 0 {
		goto L1
	} else {
		goto L147
	}
L130:
	;
	v400 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+190)) = uint8(v400)
	goto L129
L131:
	;
	goto L132
L132:
	;
	F_deconstruct_array_builtin(m, v334, int32(26), v19+int32(172), int32(0), v19+int32(168))
	mBase = m.M
	v409 = m.ExcPending
	if v409 != 0 {
		goto L1
	} else {
		goto L133
	}
L133:
	;
	v410 = *(*int32)(unsafe.Add(mBase, uint32(v19)+172))
	v412 = v342 - int32(1)
	if v412 <= v359 {
		goto L134
	} else {
		goto L135
	}
L134:
	;
	v523 = F_construct_array_builtin(m, v410, v412, int32(26))
	mBase = m.M
	v524 = m.ExcPending
	if v524 != 0 {
		goto L1
	} else {
		goto L146
	}
L135:
	;
	v416 = (v412 - v359) & int32(3)
	if v416 != 0 {
		goto L136
	} else {
		goto L137
	}
L136:
	;
	v421 = v359
	v429 = int32(0)
	goto L139
L137:
	;
	v450 = v359
	goto L138
L138:
	;
	if base.Ui32(v342-v359-int32(2)) < base.Ui32(int32(3)) {
		goto L134
	} else {
		goto L142
	}
L139:
	;
	v433 = int32(3)
	v436 = int32(1)
	v437 = v421 + v436
	v441 = *(*int64)(unsafe.Add(mBase, uint32(v410+v437<<(uint(v433)%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v410+v421<<(uint(v433)%32)))) = v441
	v444 = v429 + v436
	if v444 != v416 {
		v421 = v437
		v429 = v444
		goto L139
	} else {
		goto L141
	}
L140:
	;
	v450 = v437
	goto L138
L141:
	;
	goto L140
L142:
	;
	v471 = v450
	goto L143
L143:
	;
	v483 = int32(3)
	v485 = v410 + v471<<(uint(v483)%32)
	v487 = v485 + int32(8)
	v488 = *(*int64)(unsafe.Add(mBase, uint32(v487)))
	*(*int64)(unsafe.Add(mBase, uint32(v485))) = v488
	v491 = v485 + int32(16)
	v492 = *(*int64)(unsafe.Add(mBase, uint32(v491)))
	*(*int64)(unsafe.Add(mBase, uint32(v487))) = v492
	v495 = v485 + int32(24)
	v496 = *(*int64)(unsafe.Add(mBase, uint32(v495)))
	*(*int64)(unsafe.Add(mBase, uint32(v491))) = v496
	v499 = v471 + int32(4)
	v503 = *(*int64)(unsafe.Add(mBase, uint32(v410+v499<<(uint(v483)%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v495))) = v503
	if v499 != v412 {
		v471 = v499
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
	*(*int64)(unsafe.Add(mBase, uint32(v19)+240)) = base.I64_extend_i32_u(v523)
	goto L129
L147:
	;
	v549 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+271)))
	if v549 == int32(1) {
		goto L6
	} else {
		goto L148
	}
L148:
	;
	v553 = F_pg_detoast_datum(m, base.I32_wrap_i64(v547))
	mBase = m.M
	v554 = m.ExcPending
	if v554 != 0 {
		goto L1
	} else {
		goto L149
	}
L149:
	;
	v555 = *(*int32)(unsafe.Add(mBase, uint32(v553)+4))
	if v555 != int32(1) {
		goto L5
	} else {
		goto L150
	}
L150:
	;
	v558 = *(*int32)(unsafe.Add(mBase, uint32(v553)+20))
	if v558 != int32(1) {
		goto L5
	} else {
		goto L151
	}
L151:
	;
	v561 = *(*int32)(unsafe.Add(mBase, uint32(v553)+8))
	if v561 != 0 {
		goto L5
	} else {
		goto L152
	}
L152:
	;
	v562 = *(*int32)(unsafe.Add(mBase, uint32(v553)+12))
	if v562 != int32(25) {
		goto L5
	} else {
		goto L153
	}
L153:
	;
	v565 = *(*int32)(unsafe.Add(mBase, uint32(v553)+16))
	if v565 != v342 {
		goto L4
	} else {
		goto L154
	}
L154:
	;
	if v342 == int32(1) {
		goto L156
	} else {
		goto L157
	}
L155:
	;
	v713 = *(*int32)(unsafe.Add(mBase, uint32(v306)+52))
	v720 = F_heap_modify_tuple(m, v322, v713, v19+int32(192), v19+int32(184), v19+int32(176))
	mBase = m.M
	v721 = m.ExcPending
	if v721 != 0 {
		goto L1
	} else {
		goto L173
	}
L156:
	;
	v569 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+191)) = uint8(v569)
	goto L155
L157:
	;
	goto L158
L158:
	;
	v571 = int32(0)
	F_deconstruct_array_builtin(m, v553, int32(25), v19+int32(172), v571, v19+int32(168))
	mBase = m.M
	v579 = m.ExcPending
	if v579 != 0 {
		goto L1
	} else {
		goto L159
	}
L159:
	;
	v580 = *(*int32)(unsafe.Add(mBase, uint32(v19)+172))
	v582 = v342 - int32(1)
	if v582 <= v359 {
		goto L160
	} else {
		goto L161
	}
L160:
	;
	v693 = F_construct_array_builtin(m, v580, v582, int32(25))
	mBase = m.M
	v694 = m.ExcPending
	if v694 != 0 {
		goto L1
	} else {
		goto L172
	}
L161:
	;
	v589 = (v582 - v359) & int32(3)
	if v589 != 0 {
		goto L162
	} else {
		goto L163
	}
L162:
	;
	v595 = v359
	v599 = v571
	goto L165
L163:
	;
	v624 = v359
	goto L164
L164:
	;
	if base.Ui32(v342-v359-int32(2)) < base.Ui32(int32(3)) {
		goto L160
	} else {
		goto L168
	}
L165:
	;
	v606 = int32(3)
	v609 = int32(1)
	v610 = v595 + v609
	v614 = *(*int64)(unsafe.Add(mBase, uint32(v580+v610<<(uint(v606)%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v580+v595<<(uint(v606)%32)))) = v614
	v617 = v599 + v609
	if v617 != v589 {
		v595 = v610
		v599 = v617
		goto L165
	} else {
		goto L167
	}
L166:
	;
	v624 = v610
	goto L164
L167:
	;
	goto L166
L168:
	;
	v642 = v624
	goto L169
L169:
	;
	v653 = int32(3)
	v655 = v580 + v642<<(uint(v653)%32)
	v657 = v655 + int32(8)
	v658 = *(*int64)(unsafe.Add(mBase, uint32(v657)))
	*(*int64)(unsafe.Add(mBase, uint32(v655))) = v658
	v661 = v655 + int32(16)
	v662 = *(*int64)(unsafe.Add(mBase, uint32(v661)))
	*(*int64)(unsafe.Add(mBase, uint32(v657))) = v662
	v665 = v655 + int32(24)
	v666 = *(*int64)(unsafe.Add(mBase, uint32(v665)))
	*(*int64)(unsafe.Add(mBase, uint32(v661))) = v666
	v669 = v642 + int32(4)
	v673 = *(*int64)(unsafe.Add(mBase, uint32(v580+v669<<(uint(v653)%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v665))) = v673
	if v669 != v582 {
		v642 = v669
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
	*(*int64)(unsafe.Add(mBase, uint32(v19)+248)) = base.I64_extend_i32_u(v693)
	goto L155
L173:
	;
	F_CatalogTupleUpdate(m, v306, v720+int32(4), v720)
	mBase = m.M
	v725 = m.ExcPending
	if v725 != 0 {
		goto L1
	} else {
		goto L174
	}
L174:
	;
	goto L114
L175:
	;
	F_relation_close(m, v306, int32(3))
	mBase = m.M
	v746 = m.ExcPending
	if v746 != 0 {
		goto L1
	} else {
		goto L176
	}
L176:
	;
	v747 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v752 = v747
	goto L107
L177:
	;
	m.G0 = v767 + int32(16)
	goto L13
L178:
	;
	F_ReleaseCatCache(m, v773)
	mBase = m.M
	v870 = m.ExcPending
	if v870 != 0 {
		goto L1
	} else {
		goto L202
	}
L179:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v859 = m.ExcPending
	if v859 != 0 {
		goto L1
	} else {
		goto L199
	}
L180:
	;
	v772 = base.I64_extend_i32_u(v764)
	v773 = F_SearchSysCache1(m, int32(57), v772)
	mBase = m.M
	v774 = m.ExcPending
	if v774 != 0 {
		goto L1
	} else {
		goto L183
	}
L181:
	;
	goto L182
L182:
	;
	v852 = int32(0)
	F_recordExtensionInitPrivWorker(m, v764, v752, v852, v852)
	mBase = m.M
	v855 = m.ExcPending
	if v855 != 0 {
		goto L1
	} else {
		goto L198
	}
L183:
	;
	if v773 == int32(0) {
		goto L179
	} else {
		goto L184
	}
L184:
	;
	v777 = *(*int32)(unsafe.Add(mBase, uint32(v773)+16))
	v778 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v777)+22)))
	v779 = v777 + v778
	v780 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v779)+119)))
	switch v780 - int32(73) {
	case 0, 26, 32:
		goto L178
	default:
		goto L186
	case 10:
		goto L185
	}
L185:
	;
	F_ReleaseCatCache(m, v773)
	mBase = m.M
	v835 = m.ExcPending
	if v835 != 0 {
		goto L1
	} else {
		goto L197
	}
L186:
	;
	v784 = int32(*(*int16)(unsafe.Add(mBase, uint32(v779)+120)))
	if v784 <= int32(0) {
		goto L185
	} else {
		goto L187
	}
L187:
	;
	v792 = int32(1)
	goto L188
L188:
	;
	v806 = F_SearchSysCache2(m, int32(7), v772, base.I64_extend16_s(base.I64_extend_i32_u(v792)))
	mBase = m.M
	v807 = m.ExcPending
	if v807 != 0 {
		goto L1
	} else {
		goto L190
	}
L189:
	;
	goto L185
L190:
	;
	if v806 != 0 {
		goto L191
	} else {
		goto L192
	}
L191:
	;
	F_recordExtensionInitPrivWorker(m, v764, int32(1259), v792, int32(0))
	mBase = m.M
	v811 = m.ExcPending
	if v811 != 0 {
		goto L1
	} else {
		goto L194
	}
L192:
	;
	goto L193
L193:
	;
	v816 = base.I32_extend16_s(v792 + int32(1))
	if v816 <= v784 {
		v792 = v816
		goto L188
	} else {
		goto L196
	}
L194:
	;
	F_ReleaseCatCache(m, v806)
	mBase = m.M
	v813 = m.ExcPending
	if v813 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v767))) = v764
	F_errmsg_internal(m, int32(_a_F_ExecAlterExtensionContentsRecurse_0), v767)
	mBase = m.M
	v863 = m.ExcPending
	if v863 != 0 {
		goto L1
	} else {
		goto L200
	}
L200:
	;
	F_errfinish(m, int32(_a_F_ExecAlterExtensionContentsRecurse_1), int32(_a_F_ExecAlterExtensionContentsRecurse_8), int32(_a_F_ExecAlterExtensionContentsRecurse_9))
	mBase = m.M
	v868 = m.ExcPending
	if v868 != 0 {
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
	m.G0 = v19 + int32(336)
	return
L204:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+200)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+192)) = int32(1247)
	v913 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v914 = F_get_array_type(m, v913)
	mBase = m.M
	v915 = m.ExcPending
	if v915 != 0 {
		goto L1
	} else {
		goto L207
	}
L205:
	;
	v957 = v906
	goto L206
L206:
	;
	if v957 != int32(1259) {
		goto L203
	} else {
		goto L219
	}
L207:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+196)) = v914
	if v914 != 0 {
		goto L208
	} else {
		goto L209
	}
L208:
	;
	v917 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+120)) = v917
	v919 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	*(*int64)(unsafe.Add(mBase, uint32(v19)+112)) = v919
	v921 = *(*int64)(unsafe.Add(mBase, uint32(v19)+192))
	*(*int64)(unsafe.Add(mBase, uint32(v19)+96)) = v921
	v923 = *(*int32)(unsafe.Add(mBase, uint32(v19)+200))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+104)) = v923
	F_ExecAlterExtensionContentsRecurse(m, l0, v19+int32(112), v19+int32(96))
	mBase = m.M
	v930 = m.ExcPending
	if v930 != 0 {
		goto L1
	} else {
		goto L211
	}
L209:
	;
	goto L210
L210:
	;
	v931 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v932 = F_type_is_range(m, v931)
	mBase = m.M
	v933 = m.ExcPending
	if v933 != 0 {
		goto L1
	} else {
		goto L212
	}
L211:
	;
	goto L210
L212:
	;
	if v932 != 0 {
		goto L213
	} else {
		goto L214
	}
L213:
	;
	v934 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v935 = F_get_range_multirange(m, v934)
	mBase = m.M
	v936 = m.ExcPending
	if v936 != 0 {
		goto L1
	} else {
		goto L216
	}
L214:
	;
	goto L215
L215:
	;
	v955 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v957 = v955
	goto L206
L216:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+196)) = v935
	if v935 == int32(0) {
		goto L3
	} else {
		goto L217
	}
L217:
	;
	v940 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+88)) = v940
	v942 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	*(*int64)(unsafe.Add(mBase, uint32(v19)+80)) = v942
	v944 = *(*int64)(unsafe.Add(mBase, uint32(v19)+192))
	*(*int64)(unsafe.Add(mBase, uint32(v19)+64)) = v944
	v946 = *(*int32)(unsafe.Add(mBase, uint32(v19)+200))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+72)) = v946
	F_ExecAlterExtensionContentsRecurse(m, l0, v19+int32(80), v19-int32(-64))
	mBase = m.M
	v953 = m.ExcPending
	if v953 != 0 {
		goto L1
	} else {
		goto L218
	}
L218:
	;
	goto L215
L219:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+200)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+192)) = int32(1247)
	v964 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v965 = F_get_rel_type_id(m, v964)
	mBase = m.M
	v966 = m.ExcPending
	if v966 != 0 {
		goto L1
	} else {
		goto L220
	}
L220:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+196)) = v965
	if v965 == int32(0) {
		goto L203
	} else {
		goto L221
	}
L221:
	;
	v970 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+56)) = v970
	v972 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	*(*int64)(unsafe.Add(mBase, uint32(v19)+48)) = v972
	v974 = *(*int64)(unsafe.Add(mBase, uint32(v19)+192))
	*(*int64)(unsafe.Add(mBase, uint32(v19)+32)) = v974
	v976 = *(*int32)(unsafe.Add(mBase, uint32(v19)+200))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+40)) = v976
	F_ExecAlterExtensionContentsRecurse(m, l0, v19+int32(48), v19+int32(32))
	mBase = m.M
	v983 = m.ExcPending
	if v983 != 0 {
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
	v994 = m.ExcPending
	if v994 != 0 {
		goto L1
	} else {
		goto L224
	}
L224:
	;
	v996 = F_getObjectDescription(m, l2, int32(0))
	mBase = m.M
	v997 = m.ExcPending
	if v997 != 0 {
		goto L1
	} else {
		goto L225
	}
L225:
	;
	v998 = F_get_extension_name(m, v23)
	mBase = m.M
	v999 = m.ExcPending
	if v999 != 0 {
		goto L1
	} else {
		goto L226
	}
L226:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+132)) = v998
	*(*int32)(unsafe.Add(mBase, uint32(v19)+128)) = v996
	F_errmsg(m, int32(_a_F_ExecAlterExtensionContentsRecurse_10), v19+int32(128))
	mBase = m.M
	v1006 = m.ExcPending
	if v1006 != 0 {
		goto L1
	} else {
		goto L227
	}
L227:
	;
	F_errfinish(m, int32(_a_F_ExecAlterExtensionContentsRecurse_11), int32(3885), int32(_a_F_ExecAlterExtensionContentsRecurse_12))
	mBase = m.M
	v1011 = m.ExcPending
	if v1011 != 0 {
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
	v1018 = m.ExcPending
	if v1018 != 0 {
		goto L1
	} else {
		goto L230
	}
L230:
	;
	v1019 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v1020 = F_get_namespace_name(m, v1019)
	mBase = m.M
	v1021 = m.ExcPending
	if v1021 != 0 {
		goto L1
	} else {
		goto L231
	}
L231:
	;
	v1022 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = v1022
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v1020
	F_errmsg(m, int32(_a_F_ExecAlterExtensionContentsRecurse_13), v19)
	mBase = m.M
	v1027 = m.ExcPending
	if v1027 != 0 {
		goto L1
	} else {
		goto L232
	}
L232:
	;
	F_errfinish(m, int32(_a_F_ExecAlterExtensionContentsRecurse_11), int32(3898), int32(_a_F_ExecAlterExtensionContentsRecurse_12))
	mBase = m.M
	v1032 = m.ExcPending
	if v1032 != 0 {
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
	v1039 = m.ExcPending
	if v1039 != 0 {
		goto L1
	} else {
		goto L235
	}
L235:
	;
	v1041 = F_getObjectDescription(m, l2, int32(0))
	mBase = m.M
	v1042 = m.ExcPending
	if v1042 != 0 {
		goto L1
	} else {
		goto L236
	}
L236:
	;
	v1043 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+164)) = v1043
	*(*int32)(unsafe.Add(mBase, uint32(v19)+160)) = v1041
	F_errmsg(m, int32(_a_F_ExecAlterExtensionContentsRecurse_14), v19+int32(160))
	mBase = m.M
	v1050 = m.ExcPending
	if v1050 != 0 {
		goto L1
	} else {
		goto L237
	}
L237:
	;
	F_errfinish(m, int32(_a_F_ExecAlterExtensionContentsRecurse_11), int32(3925), int32(_a_F_ExecAlterExtensionContentsRecurse_12))
	mBase = m.M
	v1055 = m.ExcPending
	if v1055 != 0 {
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
	v1063 = m.ExcPending
	if v1063 != 0 {
		goto L1
	} else {
		goto L240
	}
L240:
	;
	F_errfinish(m, int32(_a_F_ExecAlterExtensionContentsRecurse_11), int32(3933), int32(_a_F_ExecAlterExtensionContentsRecurse_12))
	mBase = m.M
	v1068 = m.ExcPending
	if v1068 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v19)+144)) = v23
	F_errmsg_internal(m, int32(_a_F_ExecAlterExtensionContentsRecurse_16), v19+int32(144))
	mBase = m.M
	v1078 = m.ExcPending
	if v1078 != 0 {
		goto L1
	} else {
		goto L243
	}
L243:
	;
	F_errfinish(m, int32(_a_F_ExecAlterExtensionContentsRecurse_11), int32(3123), int32(_a_F_ExecAlterExtensionContentsRecurse_17))
	mBase = m.M
	v1083 = m.ExcPending
	if v1083 != 0 {
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
	v1092 = m.ExcPending
	if v1092 != 0 {
		goto L1
	} else {
		goto L246
	}
L246:
	;
	F_errfinish(m, int32(_a_F_ExecAlterExtensionContentsRecurse_11), int32(3148), int32(_a_F_ExecAlterExtensionContentsRecurse_17))
	mBase = m.M
	v1097 = m.ExcPending
	if v1097 != 0 {
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
	v1105 = m.ExcPending
	if v1105 != 0 {
		goto L1
	} else {
		goto L249
	}
L249:
	;
	F_errfinish(m, int32(_a_F_ExecAlterExtensionContentsRecurse_11), int32(3205), int32(_a_F_ExecAlterExtensionContentsRecurse_17))
	mBase = m.M
	v1110 = m.ExcPending
	if v1110 != 0 {
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
	v1118 = m.ExcPending
	if v1118 != 0 {
		goto L1
	} else {
		goto L252
	}
L252:
	;
	F_errfinish(m, int32(_a_F_ExecAlterExtensionContentsRecurse_11), int32(3215), int32(_a_F_ExecAlterExtensionContentsRecurse_17))
	mBase = m.M
	v1123 = m.ExcPending
	if v1123 != 0 {
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
	v1131 = m.ExcPending
	if v1131 != 0 {
		goto L1
	} else {
		goto L255
	}
L255:
	;
	F_errfinish(m, int32(_a_F_ExecAlterExtensionContentsRecurse_11), int32(3217), int32(_a_F_ExecAlterExtensionContentsRecurse_17))
	mBase = m.M
	v1136 = m.ExcPending
	if v1136 != 0 {
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
	v1143 = m.ExcPending
	if v1143 != 0 {
		goto L1
	} else {
		goto L258
	}
L258:
	;
	v1144 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v1145 = F_format_type_be(m, v1144)
	mBase = m.M
	v1146 = m.ExcPending
	if v1146 != 0 {
		goto L1
	} else {
		goto L259
	}
L259:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = v1145
	F_errmsg(m, int32(_a_F_ExecAlterExtensionContentsRecurse_21), v19+int32(16))
	mBase = m.M
	v1152 = m.ExcPending
	if v1152 != 0 {
		goto L1
	} else {
		goto L260
	}
L260:
	;
	F_errfinish(m, int32(_a_F_ExecAlterExtensionContentsRecurse_11), int32(3978), int32(_a_F_ExecAlterExtensionContentsRecurse_12))
	mBase = m.M
	v1157 = m.ExcPending
	if v1157 != 0 {
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
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v45 int64
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v85 int64
	_ = v85
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v114 int64
	_ = v114
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v133 int32
	_ = v133
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v183 int64
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v293 int32
	_ = v293
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v305 int32
	_ = v305
	var v328 int32
	_ = v328
	var v351 int32
	_ = v351
	var v356 int32
	_ = v356
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v362 int32
	_ = v362
	var v368 int32
	_ = v368
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v404 int32
	_ = v404
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v409 int32
	_ = v409
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v418 int32
	_ = v418
	var v422 int32
	_ = v422
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v474 int32
	_ = v474
	var v477 int32
	_ = v477
	var v482 int32
	_ = v482
	var v485 int32
	_ = v485
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v493 int32
	_ = v493
	var v495 int32
	_ = v495
	var v499 int32
	_ = v499
	var v533 int32
	_ = v533
	var v536 int32
	_ = v536
	var v540 int32
	_ = v540
	var v545 int32
	_ = v545
	var v549 int32
	_ = v549
	var v552 int32
	_ = v552
	var v559 int32
	_ = v559
	var v564 int32
	_ = v564
	var v568 int32
	_ = v568
	var v574 int32
	_ = v574
	var v579 int32
	_ = v579
	var v583 int32
	_ = v583
	var v586 int32
	_ = v586
	var v594 int32
	_ = v594
	var v599 int32
	_ = v599
	var v603 int32
	_ = v603
	var v607 int32
	_ = v607
	var v612 int32
	_ = v612
	var v616 int32
	_ = v616
	var v619 int32
	_ = v619
	var v625 int32
	_ = v625
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v631 int32
	_ = v631
	var v632 int32
	_ = v632
	var v638 int32
	_ = v638
	var v639 int32
	_ = v639
	var v644 int32
	_ = v644
	var v648 int32
	_ = v648
	var v656 int32
	_ = v656
	var v661 int32
	_ = v661
	var v663 int32
	_ = v663
	var v667 int32
	_ = v667
	var v668 int32
	_ = v668
	var v670 int32
	_ = v670
	var v672 int32
	_ = v672
	var v674 int32
	_ = v674
	var v676 int32
	_ = v676
	var v677 int32
	_ = v677
	var v682 int32
	_ = v682
	var v683 int32
	_ = v683
	var v684 int32
	_ = v684
	var v685 int32
	_ = v685
	var v689 int32
	_ = v689
	var v694 int32
	_ = v694
	var v696 int32
	_ = v696
	var v699 int64
	_ = v699
	var v702 int32
	_ = v702
	var v703 int32
	_ = v703
	var v704 int32
	_ = v704
	var v705 int32
	_ = v705
	var v706 int32
	_ = v706
	var v711 int32
	_ = v711
	var v713 int32
	_ = v713
	var v714 int32
	_ = v714
	var v715 int32
	_ = v715
	var v717 int32
	_ = v717
	var v718 int32
	_ = v718
	var v719 int32
	_ = v719
	var v720 int32
	_ = v720
	var v722 int32
	_ = v722
	var v726 int32
	_ = v726
	var v727 int32
	_ = v727
	var v728 int32
	_ = v728
	var v730 int32
	_ = v730
	var v731 int32
	_ = v731
	var v733 int32
	_ = v733
	var v734 int32
	_ = v734
	var v735 int32
	_ = v735
	var v737 int32
	_ = v737
	var v739 int32
	_ = v739
	var v743 int32
	_ = v743
	var v760 int32
	_ = v760
	var v763 int32
	_ = v763
	var v767 int32
	_ = v767
	var v768 int32
	_ = v768
	var v769 int32
	_ = v769
	var v770 int32
	_ = v770
	var v771 int32
	_ = v771
	var v779 int32
	_ = v779
	var v780 int32
	_ = v780
	var v785 int32
	_ = v785
	var v787 int32
	_ = v787
	var v788 int32
	_ = v788
	var v789 int32
	_ = v789
	var v793 int32
	_ = v793
	var v794 int32
	_ = v794
	var v796 int32
	_ = v796
	var v799 int32
	_ = v799
	var v800 int32
	_ = v800
	var v801 int32
	_ = v801
	var v802 int32
	_ = v802
	var v805 int32
	_ = v805
	var v806 int32
	_ = v806
	var v809 int32
	_ = v809
	var v810 int32
	_ = v810
	var v812 int32
	_ = v812
	var v813 int32
	_ = v813
	var v814 int32
	_ = v814
	var v815 int32
	_ = v815
	var v817 int32
	_ = v817
	var v830 int32
	_ = v830
	var v833 int32
	_ = v833
	var v834 int32
	_ = v834
	var v835 int32
	_ = v835
	var v839 int32
	_ = v839
	var v844 int32
	_ = v844
	var v848 int32
	_ = v848
	var v849 int32
	_ = v849
	var v853 int32
	_ = v853
	var v858 int32
	_ = v858
	var v861 int32
	_ = v861
	var v862 int32
	_ = v862
	var v866 int32
	_ = v866
	var v867 int32
	_ = v867
	var v868 int32
	_ = v868
	var v869 int32
	_ = v869
	var v871 int32
	_ = v871
	var v872 int32
	_ = v872
	var v873 int32
	_ = v873
	var v874 int32
	_ = v874
	var v875 int32
	_ = v875
	var v876 int32
	_ = v876
	var v877 int32
	_ = v877
	var v881 int32
	_ = v881
	var v904 int32
	_ = v904
	var v905 int32
	_ = v905
	var v906 int32
	_ = v906
	var v910 int32
	_ = v910
	var v911 int32
	_ = v911
	var v912 int32
	_ = v912
	var v931 int32
	_ = v931
	v23 = m.G0
	v25 = v23 - int32(32)
	m.G0 = v25
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	switch v27 - int32(1) {
	case 0, 6, 7, 18, 23, 24, 25, 28, 34, 39, 45, 46, 47, 48:
		goto L3
	default:
		goto L4
	case 11, 49:
		goto L5
	case 14:
		goto L7
	case 17, 22, 37, 41, 51:
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
	v904 = *(*int32)(unsafe.Add(mBase, uint32(v25)+24))
	v905 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
	v906 = *(*int32)(unsafe.Add(mBase, uint32(v25)+16))
	v910 = v906
	v911 = v905
	v912 = v904
	goto L1
L3:
	;
	v861 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v862 = int32(0)
	F_get_object_address(m, v25+int32(16), v27, v861, v862, int32(8), v862)
	mBase = m.M
	v866 = m.ExcPending
	if v866 != 0 {
		goto L18
	} else {
		goto L235
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v848 = m.ExcPending
	if v848 != 0 {
		goto L18
	} else {
		goto L232
	}
L5:
	;
	v787 = v25 + int32(16)
	v788 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v789 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if l2 != 0 {
		goto L208
	} else {
		goto L209
	}
L6:
	;
	v663 = v25 + int32(16)
	if l2 != 0 {
		goto L168
	} else {
		goto L169
	}
L7:
	;
	v31 = v25 + int32(16)
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+4))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if l2 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v38 = v25 + int32(28)
	goto L10
L9:
	;
	v38 = int32(0)
	goto L10
L10:
	;
	v39 = m.G0
	v41 = v39 - int32(256)
	m.G0 = v41
	v45 = int64(0)
	v48 = F_GetSysCacheOid(m, int32(27), base.I64_extend_i32_u(v33), v45, v45, v45)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
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
	v648 = m.ExcPending
	if v648 != 0 {
		goto L18
	} else {
		goto L165
	}
L13:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v616 = m.ExcPending
	if v616 != 0 {
		goto L18
	} else {
		goto L158
	}
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v603 = m.ExcPending
	if v603 != 0 {
		goto L18
	} else {
		goto L155
	}
L15:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v583 = m.ExcPending
	if v583 != 0 {
		goto L18
	} else {
		goto L151
	}
L16:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v568 = m.ExcPending
	if v568 != 0 {
		goto L18
	} else {
		goto L148
	}
L17:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v549 = m.ExcPending
	if v549 != 0 {
		goto L18
	} else {
		goto L144
	}
L18:
	;
	return
L19:
	;
	if v48 != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v50 = F_LookupCreationNamespace(m, v34)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
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
	v533 = m.ExcPending
	if v533 != 0 {
		goto L18
	} else {
		goto L140
	}
L23:
	;
	v54 = *(*int32)(unsafe.Add(mBase, _c_F_ExecAlterObjectSchemaStmt[0]))
	v55 = F_object_ownercheck(m, int32(3079), v48, v54)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L18
	} else {
		goto L24
	}
L24:
	;
	if v55 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	F_aclcheck_error(m, int32(2), int32(15), v33)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L18
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	v65 = *(*int32)(unsafe.Add(mBase, _c_F_ExecAlterObjectSchemaStmt[0]))
	v67 = F_object_aclcheck(m, int32(2615), v50, v65, int64(512))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L18
	} else {
		goto L29
	}
L28:
	;
	goto L27
L29:
	;
	if v67 != 0 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	F_aclcheck_error(m, v67, int32(37), v34)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L18
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	v73 = F_getExtensionOfObject(m, int32(2615), v50)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L18
	} else {
		goto L34
	}
L33:
	;
	goto L32
L34:
	;
	if v73 == v48 {
		goto L17
	} else {
		goto L35
	}
L35:
	;
	v78 = F_table_open(m, int32(3079), int32(3))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L18
	} else {
		goto L36
	}
L36:
	;
	v81 = v41 + int32(144)
	v85 = base.I64_extend_i32_u(v48)
	F_ScanKeyInit(m, v81, int32(1), int32(3), int32(184), v85)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L18
	} else {
		goto L37
	}
L37:
	;
	v89 = int32(1)
	v92 = F_systable_beginscan(m, v78, int32(3080), v89, int32(0), v89, v81)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L18
	} else {
		goto L38
	}
L38:
	;
	v94 = F_systable_getnext(m, v92)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L18
	} else {
		goto L39
	}
L39:
	;
	if v94 == int32(0) {
		goto L16
	} else {
		goto L40
	}
L40:
	;
	v98 = F_heap_copytuple(m, v94)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L18
	} else {
		goto L41
	}
L41:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v98)+16))
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100)+22)))
	F_systable_endscan(m, v92)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L18
	} else {
		goto L42
	}
L42:
	;
	v104 = v100 + v101
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v104)+72))
	if v50 == v105 {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	m.G0 = v41 + int32(256)
	goto L11
L44:
	;
	F_relation_close(m, v78, int32(3))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L18
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104)+76)))
	if v116 == int32(0) {
		goto L15
	} else {
		goto L48
	}
L47:
	;
	v111 = *(*int32)(unsafe.Add(mBase, _c_F_ExecAlterObjectSchemaStmt[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+8)) = v111
	v114 = *(*int64)(unsafe.Add(mBase, _c_F_ExecAlterObjectSchemaStmt[2]))
	*(*int64)(unsafe.Add(mBase, uint32(v31))) = v114
	goto L43
L48:
	;
	v119 = F_new_object_addresses(m)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L18
	} else {
		goto L49
	}
L49:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v104)+72))
	v124 = F_table_open(m, int32(2608), int32(1))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L18
	} else {
		goto L50
	}
L50:
	;
	v127 = v41 + int32(144)
	F_ScanKeyInit(m, v127, int32(4), int32(3), int32(184), int64(3079))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L18
	} else {
		goto L51
	}
L51:
	;
	F_ScanKeyInit(m, v41+int32(200), int32(5), int32(3), int32(184), v85)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L18
	} else {
		goto L52
	}
L52:
	;
	v145 = F_systable_beginscan(m, v124, int32(2674), int32(1), int32(0), int32(2), v127)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L18
	} else {
		goto L53
	}
L53:
	;
	v147 = F_systable_getnext(m, v145)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L18
	} else {
		goto L54
	}
L54:
	;
	if v147 != 0 {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v150 = v104 + int32(4)
	v152 = v147
	goto L58
L56:
	;
	goto L57
L57:
	;
	if v38 != 0 {
		goto L127
	} else {
		goto L128
	}
L58:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v152)+16))
	v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+22)))
	v175 = v173 + v174
	v176 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175)+24)))
	if v176 == int32(110) {
		goto L61
	} else {
		goto L62
	}
L59:
	;
	goto L57
L60:
	;
	v448 = F_systable_getnext(m, v145)
	mBase = m.M
	v449 = m.ExcPending
	if v449 != 0 {
		goto L18
	} else {
		goto L125
	}
L61:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v175)))
	if v179 != int32(3079) {
		goto L60
	} else {
		goto L64
	}
L62:
	;
	v351 = v176
	goto L63
L63:
	;
	if v351&int32(255) != int32(101) {
		goto L60
	} else {
		goto L99
	}
L64:
	;
	v183 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v175)+4)))
	v184 = F_SearchSysCache1(m, int32(28), v183)
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L18
	} else {
		goto L66
	}
L65:
	;
	v200 = F_palloc0(m, int32(48))
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L18
	} else {
		goto L72
	}
L66:
	;
	if v184 == int32(0) {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v198 = int32(0)
	goto L65
L68:
	;
	goto L69
L69:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v184)+16))
	v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v189)+22)))
	v194 = F_pstrdup(m, v189+v190+int32(4))
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L18
	} else {
		goto L70
	}
L70:
	;
	F_ReleaseCatCache(m, v184)
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L18
	} else {
		goto L71
	}
L71:
	;
	v198 = v194
	goto L65
L72:
	;
	v202 = F_pstrdup(m, v198)
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L18
	} else {
		goto L73
	}
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v200)+36)) = int32(-1)
	v206 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v200)+34)) = uint8(v206)
	v208 = int32(256)
	*(*uint16)(unsafe.Add(mBase, uint32(v200)+32)) = uint16(v208)
	*(*int32)(unsafe.Add(mBase, uint32(v200))) = v202
	F_parse_extension_control_file(m, v200, v206)
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L18
	} else {
		goto L74
	}
L74:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v200)+44))
	if v214 == int32(0) {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v328 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175)+24)))
	v351 = v328
	goto L63
L76:
	;
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v214)+4))
	if v217 <= int32(0) {
		goto L75
	} else {
		goto L77
	}
L77:
	;
	v220 = int32(0)
	if v220 < v217 {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v223 = v217
	goto L80
L79:
	;
	v223 = v220
	goto L80
L80:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v214)+12))
	v227 = int32(0)
	goto L81
L81:
	;
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v224+v227<<(uint(int32(2))%32))))
	v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v251))))
	v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v150))))
	if base.B2i32(v254 == int32(0))|base.B2i32(v254 != v257) != 0 {
		v275 = v254
		v276 = v257
		goto L84
	} else {
		goto L85
	}
L82:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L18
	} else {
		goto L94
	}
L83:
	;
	if v275-v276 != 0 {
		goto L90
	} else {
		goto L91
	}
L84:
	;
	goto L83
L85:
	;
	v260 = v251
	v261 = v150
	goto L86
L86:
	;
	v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v261)+1)))
	v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v260)+1)))
	if v265 == int32(0) {
		v275 = v265
		v276 = v264
		goto L84
	} else {
		goto L88
	}
L87:
	;
	v275 = v265
	v276 = v264
	goto L84
L88:
	;
	v268 = int32(1)
	if v265 == v264 {
		v260 = v260 + v268
		v261 = v261 + v268
		goto L86
	} else {
		goto L89
	}
L89:
	;
	goto L87
L90:
	;
	v279 = v227 + int32(1)
	if v223 != v279 {
		v227 = v279
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
	v287 = m.ExcPending
	if v287 != 0 {
		goto L18
	} else {
		goto L95
	}
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+112)) = v150
	F_errmsg(m, int32(_a_F_ExecAlterObjectSchemaStmt_0), v41+int32(112))
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L18
	} else {
		goto L96
	}
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+100)) = v150
	*(*int32)(unsafe.Add(mBase, uint32(v41)+96)) = v198
	v299 = F_errdetail(m, int32(_a_F_ExecAlterObjectSchemaStmt_1), v41+int32(96))
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L18
	} else {
		goto L97
	}
L97:
	;
	F_errfinish(m, int32(_a_F_ExecAlterObjectSchemaStmt_2), int32(3401), int32(_a_F_ExecAlterObjectSchemaStmt_3))
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
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
	v356 = *(*int32)(unsafe.Add(mBase, uint32(v175)))
	*(*int32)(unsafe.Add(mBase, uint32(v41)+132)) = v356
	v358 = *(*int32)(unsafe.Add(mBase, uint32(v175)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v41)+136)) = v358
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v175)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v41)+140)) = v360
	if v360 != 0 {
		goto L14
	} else {
		goto L100
	}
L100:
	;
	v362 = int32(0)
	if v356 <= int32(2752) {
		goto L106
	} else {
		goto L107
	}
L101:
	;
	if v422 == int32(0) {
		goto L60
	} else {
		goto L123
	}
L102:
	;
	v412 = F_table_open(m, v356, int32(3))
	mBase = m.M
	v413 = m.ExcPending
	if v413 != 0 {
		goto L18
	} else {
		goto L120
	}
L103:
	;
	v422 = v409
	goto L101
L104:
	;
	v406 = F_AlterTypeNamespace_oid(m, v358, v50, int32(1), v119)
	mBase = m.M
	v407 = m.ExcPending
	if v407 != 0 {
		goto L18
	} else {
		goto L119
	}
L105:
	;
	v396 = F_relation_open(m, v358, int32(8))
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
		goto L18
	} else {
		goto L116
	}
L106:
	;
	switch v356 - int32(1247) {
	case 0:
		goto L104
	case 1, 2, 3, 4, 5, 6, 7, 9, 10, 11:
		v409 = v362
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
	if v356 <= int32(3599) {
		goto L111
	} else {
		goto L112
	}
L109:
	;
	v368 = v356 - int32(2607)
	if base.B2i32(base.Ui32(int32(10)) < base.Ui32(v368))|base.B2i32(int32(1)<<(uint(v368)%32)&int32(1537) == int32(0)) != 0 {
		v409 = v362
		goto L103
	} else {
		goto L110
	}
L110:
	;
	goto L102
L111:
	;
	if base.B2i32(v356 == int32(2753))|base.B2i32(v356 == int32(3381))|base.B2i32(v356 == int32(3456)) != 0 {
		goto L102
	} else {
		goto L114
	}
L112:
	;
	goto L113
L113:
	;
	if base.B2i32(v356 == int32(3764))|base.B2i32(base.Ui32(v356-int32(3600)) < base.Ui32(int32(3))) != 0 {
		goto L102
	} else {
		goto L115
	}
L114:
	;
	v409 = v362
	goto L103
L115:
	;
	v409 = v362
	goto L103
L116:
	;
	v398 = *(*int32)(unsafe.Add(mBase, uint32(v396)+48))
	v399 = *(*int32)(unsafe.Add(mBase, uint32(v398)+68))
	F_AlterTableNamespaceInternal(m, v396, v399, v50, v119)
	mBase = m.M
	v401 = m.ExcPending
	if v401 != 0 {
		goto L18
	} else {
		goto L117
	}
L117:
	;
	F_relation_close(m, v396, int32(0))
	mBase = m.M
	v404 = m.ExcPending
	if v404 != 0 {
		goto L18
	} else {
		goto L118
	}
L118:
	;
	v422 = v399
	goto L101
L119:
	;
	v409 = v406
	goto L103
L120:
	;
	v414 = F_AlterObjectNamespace_internal(m, v412, v358, v50)
	mBase = m.M
	v415 = m.ExcPending
	if v415 != 0 {
		goto L18
	} else {
		goto L121
	}
L121:
	;
	F_relation_close(m, v412, int32(3))
	mBase = m.M
	v418 = m.ExcPending
	if v418 != 0 {
		goto L18
	} else {
		goto L122
	}
L122:
	;
	v422 = v414
	goto L101
L123:
	;
	if v422 != v121 {
		goto L13
	} else {
		goto L124
	}
L124:
	;
	goto L60
L125:
	;
	if v448 != 0 {
		v152 = v448
		goto L58
	} else {
		goto L126
	}
L126:
	;
	goto L59
L127:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38))) = v121
	goto L129
L128:
	;
	goto L129
L129:
	;
	F_systable_endscan(m, v145)
	mBase = m.M
	v474 = m.ExcPending
	if v474 != 0 {
		goto L18
	} else {
		goto L130
	}
L130:
	;
	F_relation_close(m, v124, int32(1))
	mBase = m.M
	v477 = m.ExcPending
	if v477 != 0 {
		goto L18
	} else {
		goto L131
	}
L131:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v104)+72)) = v50
	F_CatalogTupleUpdate(m, v78, v98+int32(4), v98)
	mBase = m.M
	v482 = m.ExcPending
	if v482 != 0 {
		goto L18
	} else {
		goto L132
	}
L132:
	;
	F_relation_close(m, v78, int32(3))
	mBase = m.M
	v485 = m.ExcPending
	if v485 != 0 {
		goto L18
	} else {
		goto L133
	}
L133:
	;
	v488 = F_changeDependencyFor(m, int32(3079), v48, int32(2615), v121, v50)
	mBase = m.M
	v489 = m.ExcPending
	if v489 != 0 {
		goto L18
	} else {
		goto L134
	}
L134:
	;
	if v488 != int32(1) {
		goto L12
	} else {
		goto L135
	}
L135:
	;
	v493 = *(*int32)(unsafe.Add(mBase, _c_F_ExecAlterObjectSchemaStmt[3]))
	if v493 != 0 {
		goto L136
	} else {
		goto L137
	}
L136:
	;
	v495 = int32(0)
	F_RunObjectPostAlterHook(m, int32(3079), v48, v495, v495, v495)
	mBase = m.M
	v499 = m.ExcPending
	if v499 != 0 {
		goto L18
	} else {
		goto L139
	}
L137:
	;
	goto L138
L138:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+4)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v31))) = int32(3079)
	goto L43
L139:
	;
	goto L138
L140:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v536 = m.ExcPending
	if v536 != 0 {
		goto L18
	} else {
		goto L141
	}
L141:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41))) = v33
	F_errmsg(m, int32(_a_F_ExecAlterObjectSchemaStmt_4), v41)
	mBase = m.M
	v540 = m.ExcPending
	if v540 != 0 {
		goto L18
	} else {
		goto L142
	}
L142:
	;
	F_errfinish(m, int32(_a_F_ExecAlterObjectSchemaStmt_2), int32(240), int32(_a_F_ExecAlterObjectSchemaStmt_5))
	mBase = m.M
	v545 = m.ExcPending
	if v545 != 0 {
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
	v552 = m.ExcPending
	if v552 != 0 {
		goto L18
	} else {
		goto L145
	}
L145:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+20)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v41)+16)) = v33
	F_errmsg(m, int32(_a_F_ExecAlterObjectSchemaStmt_6), v41+int32(16))
	mBase = m.M
	v559 = m.ExcPending
	if v559 != 0 {
		goto L18
	} else {
		goto L146
	}
L146:
	;
	F_errfinish(m, int32(_a_F_ExecAlterObjectSchemaStmt_2), int32(3301), int32(_a_F_ExecAlterObjectSchemaStmt_3))
	mBase = m.M
	v564 = m.ExcPending
	if v564 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v41)+32)) = v48
	F_errmsg_internal(m, int32(_a_F_ExecAlterObjectSchemaStmt_7), v41+int32(32))
	mBase = m.M
	v574 = m.ExcPending
	if v574 != 0 {
		goto L18
	} else {
		goto L149
	}
L149:
	;
	F_errfinish(m, int32(_a_F_ExecAlterObjectSchemaStmt_2), int32(3318), int32(_a_F_ExecAlterObjectSchemaStmt_3))
	mBase = m.M
	v579 = m.ExcPending
	if v579 != 0 {
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
	v586 = m.ExcPending
	if v586 != 0 {
		goto L18
	} else {
		goto L152
	}
L152:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+128)) = v104 + int32(4)
	F_errmsg(m, int32(_a_F_ExecAlterObjectSchemaStmt_8), v41+int32(128))
	mBase = m.M
	v594 = m.ExcPending
	if v594 != 0 {
		goto L18
	} else {
		goto L153
	}
L153:
	;
	F_errfinish(m, int32(_a_F_ExecAlterObjectSchemaStmt_2), int32(3341), int32(_a_F_ExecAlterObjectSchemaStmt_3))
	mBase = m.M
	v599 = m.ExcPending
	if v599 != 0 {
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
	v607 = m.ExcPending
	if v607 != 0 {
		goto L18
	} else {
		goto L156
	}
L156:
	;
	F_errfinish(m, int32(_a_F_ExecAlterObjectSchemaStmt_2), int32(3419), int32(_a_F_ExecAlterObjectSchemaStmt_3))
	mBase = m.M
	v612 = m.ExcPending
	if v612 != 0 {
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
	v619 = m.ExcPending
	if v619 != 0 {
		goto L18
	} else {
		goto L159
	}
L159:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+80)) = v150
	F_errmsg(m, int32(_a_F_ExecAlterObjectSchemaStmt_8), v41+int32(80))
	mBase = m.M
	v625 = m.ExcPending
	if v625 != 0 {
		goto L18
	} else {
		goto L160
	}
L160:
	;
	v629 = F_getObjectDescription(m, v41+int32(132), int32(0))
	mBase = m.M
	v630 = m.ExcPending
	if v630 != 0 {
		goto L18
	} else {
		goto L161
	}
L161:
	;
	v631 = F_get_namespace_name(m, v121)
	mBase = m.M
	v632 = m.ExcPending
	if v632 != 0 {
		goto L18
	} else {
		goto L162
	}
L162:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+68)) = v631
	*(*int32)(unsafe.Add(mBase, uint32(v41)+64)) = v629
	v638 = F_errdetail(m, int32(_a_F_ExecAlterObjectSchemaStmt_10), v41-int32(-64))
	mBase = m.M
	v639 = m.ExcPending
	if v639 != 0 {
		goto L18
	} else {
		goto L163
	}
L163:
	;
	F_errfinish(m, int32(_a_F_ExecAlterObjectSchemaStmt_2), int32(3438), int32(_a_F_ExecAlterObjectSchemaStmt_3))
	mBase = m.M
	v644 = m.ExcPending
	if v644 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v41)+48)) = v104 + int32(4)
	F_errmsg_internal(m, int32(_a_F_ExecAlterObjectSchemaStmt_11), v41+int32(48))
	mBase = m.M
	v656 = m.ExcPending
	if v656 != 0 {
		goto L18
	} else {
		goto L166
	}
L166:
	;
	F_errfinish(m, int32(_a_F_ExecAlterObjectSchemaStmt_2), int32(3460), int32(_a_F_ExecAlterObjectSchemaStmt_3))
	mBase = m.M
	v661 = m.ExcPending
	if v661 != 0 {
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
	v667 = v25 + int32(28)
	goto L170
L169:
	;
	v667 = int32(0)
	goto L170
L170:
	;
	v668 = m.G0
	v670 = v668 - int32(32)
	m.G0 = v670
	v672 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v674 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)))
	v676 = F_RangeVarGetRelidExtended(m, v672, int32(8), v674, int32(621), l1)
	mBase = m.M
	v677 = m.ExcPending
	if v677 != 0 {
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
	v760 = m.ExcPending
	if v760 != 0 {
		goto L18
	} else {
		goto L202
	}
L173:
	;
	m.G0 = v670 + int32(32)
	goto L171
L174:
	;
	if v676 == int32(0) {
		goto L175
	} else {
		goto L176
	}
L175:
	;
	v682 = F_errstart(m, int32(18), int32(0))
	mBase = m.M
	v683 = m.ExcPending
	if v683 != 0 {
		goto L18
	} else {
		goto L178
	}
L176:
	;
	goto L177
L177:
	;
	v702 = F_relation_open(m, v676, int32(0))
	mBase = m.M
	v703 = m.ExcPending
	if v703 != 0 {
		goto L18
	} else {
		goto L184
	}
L178:
	;
	if v682 != 0 {
		goto L179
	} else {
		goto L180
	}
L179:
	;
	v684 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v685 = *(*int32)(unsafe.Add(mBase, uint32(v684)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v670))) = v685
	F_errmsg(m, int32(_a_F_ExecAlterObjectSchemaStmt_12), v670)
	mBase = m.M
	v689 = m.ExcPending
	if v689 != 0 {
		goto L18
	} else {
		goto L182
	}
L180:
	;
	goto L181
L181:
	;
	v696 = *(*int32)(unsafe.Add(mBase, _c_F_ExecAlterObjectSchemaStmt[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v663)+8)) = v696
	v699 = *(*int64)(unsafe.Add(mBase, _c_F_ExecAlterObjectSchemaStmt[2]))
	*(*int64)(unsafe.Add(mBase, uint32(v663))) = v699
	goto L173
L182:
	;
	F_errfinish(m, int32(_a_F_ExecAlterObjectSchemaStmt_13), int32(_a_F_ExecAlterObjectSchemaStmt_14), int32(_a_F_ExecAlterObjectSchemaStmt_15))
	mBase = m.M
	v694 = m.ExcPending
	if v694 != 0 {
		goto L18
	} else {
		goto L183
	}
L183:
	;
	goto L181
L184:
	;
	v704 = *(*int32)(unsafe.Add(mBase, uint32(v702)+48))
	v705 = *(*int32)(unsafe.Add(mBase, uint32(v704)+68))
	v706 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v704)+119)))
	if v706 == int32(83) {
		goto L185
	} else {
		goto L186
	}
L185:
	;
	v711 = v670 + int32(28)
	v713 = v670 + int32(24)
	v714 = F_sequenceIsOwned(m, v676, int32(97), v711, v713)
	mBase = m.M
	v715 = m.ExcPending
	if v715 != 0 {
		goto L18
	} else {
		goto L188
	}
L186:
	;
	v720 = v704
	goto L187
L187:
	;
	v722 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v726 = F_makeRangeVar(m, v722, v720+int32(4), int32(-1))
	mBase = m.M
	v727 = m.ExcPending
	if v727 != 0 {
		goto L18
	} else {
		goto L192
	}
L188:
	;
	if v714 != 0 {
		goto L172
	} else {
		goto L189
	}
L189:
	;
	v717 = F_sequenceIsOwned(m, v676, int32(105), v711, v713)
	mBase = m.M
	v718 = m.ExcPending
	if v718 != 0 {
		goto L18
	} else {
		goto L190
	}
L190:
	;
	if v717 != 0 {
		goto L172
	} else {
		goto L191
	}
L191:
	;
	v719 = *(*int32)(unsafe.Add(mBase, uint32(v702)+48))
	v720 = v719
	goto L187
L192:
	;
	v728 = int32(0)
	v730 = F_RangeVarGetAndCheckCreationNamespace(m, v726, v728, v728)
	mBase = m.M
	v731 = m.ExcPending
	if v731 != 0 {
		goto L18
	} else {
		goto L193
	}
L193:
	;
	F_CheckSetNamespace(m, v705, v730)
	mBase = m.M
	v733 = m.ExcPending
	if v733 != 0 {
		goto L18
	} else {
		goto L194
	}
L194:
	;
	v734 = F_new_object_addresses(m)
	mBase = m.M
	v735 = m.ExcPending
	if v735 != 0 {
		goto L18
	} else {
		goto L195
	}
L195:
	;
	F_AlterTableNamespaceInternal(m, v702, v705, v730, v734)
	mBase = m.M
	v737 = m.ExcPending
	if v737 != 0 {
		goto L18
	} else {
		goto L196
	}
L196:
	;
	F_free_object_addresses(m, v734)
	mBase = m.M
	v739 = m.ExcPending
	if v739 != 0 {
		goto L18
	} else {
		goto L197
	}
L197:
	;
	if v667 != 0 {
		goto L198
	} else {
		goto L199
	}
L198:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v667))) = v705
	goto L200
L199:
	;
	goto L200
L200:
	;
	F_relation_close(m, v702, int32(0))
	mBase = m.M
	v743 = m.ExcPending
	if v743 != 0 {
		goto L18
	} else {
		goto L201
	}
L201:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v663)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v663)+4)) = v676
	*(*int32)(unsafe.Add(mBase, uint32(v663))) = int32(1259)
	goto L173
L202:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v763 = m.ExcPending
	if v763 != 0 {
		goto L18
	} else {
		goto L203
	}
L203:
	;
	F_errmsg(m, int32(_a_F_ExecAlterObjectSchemaStmt_16), int32(0))
	mBase = m.M
	v767 = m.ExcPending
	if v767 != 0 {
		goto L18
	} else {
		goto L204
	}
L204:
	;
	v768 = *(*int32)(unsafe.Add(mBase, uint32(v702)+48))
	v769 = *(*int32)(unsafe.Add(mBase, uint32(v670)+28))
	v770 = F_get_rel_name(m, v769)
	mBase = m.M
	v771 = m.ExcPending
	if v771 != 0 {
		goto L18
	} else {
		goto L205
	}
L205:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v670)+20)) = v770
	*(*int32)(unsafe.Add(mBase, uint32(v670)+16)) = v768 + int32(4)
	v779 = F_errdetail(m, int32(_a_F_ExecAlterObjectSchemaStmt_17), v670+int32(16))
	mBase = m.M
	v780 = m.ExcPending
	if v780 != 0 {
		goto L18
	} else {
		goto L206
	}
L206:
	;
	F_errfinish(m, int32(_a_F_ExecAlterObjectSchemaStmt_13), int32(_a_F_ExecAlterObjectSchemaStmt_18), int32(_a_F_ExecAlterObjectSchemaStmt_15))
	mBase = m.M
	v785 = m.ExcPending
	if v785 != 0 {
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
	v793 = v25 + int32(28)
	goto L210
L209:
	;
	v793 = int32(0)
	goto L210
L210:
	;
	v794 = m.G0
	v796 = v794 - int32(16)
	m.G0 = v796
	v799 = F_makeTypeNameFromNameList(m, v788)
	mBase = m.M
	v800 = m.ExcPending
	if v800 != 0 {
		goto L18
	} else {
		goto L211
	}
L211:
	;
	v801 = F_typenameTypeId(m, int32(0), v799)
	mBase = m.M
	v802 = m.ExcPending
	if v802 != 0 {
		goto L18
	} else {
		goto L212
	}
L212:
	;
	if v27 == int32(12) {
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
	v830 = m.ExcPending
	if v830 != 0 {
		goto L18
	} else {
		goto L227
	}
L215:
	;
	v805 = F_get_typtype(m, v801)
	mBase = m.M
	v806 = m.ExcPending
	if v806 != 0 {
		goto L18
	} else {
		goto L218
	}
L216:
	;
	goto L217
L217:
	;
	v809 = F_LookupCreationNamespace(m, v789)
	mBase = m.M
	v810 = m.ExcPending
	if v810 != 0 {
		goto L18
	} else {
		goto L220
	}
L218:
	;
	if v805 != int32(100) {
		goto L214
	} else {
		goto L219
	}
L219:
	;
	goto L217
L220:
	;
	v812 = F_new_object_addresses(m)
	mBase = m.M
	v813 = m.ExcPending
	if v813 != 0 {
		goto L18
	} else {
		goto L221
	}
L221:
	;
	v814 = F_AlterTypeNamespace_oid(m, v801, v809, int32(0), v812)
	mBase = m.M
	v815 = m.ExcPending
	if v815 != 0 {
		goto L18
	} else {
		goto L222
	}
L222:
	;
	F_free_object_addresses(m, v812)
	mBase = m.M
	v817 = m.ExcPending
	if v817 != 0 {
		goto L18
	} else {
		goto L223
	}
L223:
	;
	if v793 != 0 {
		goto L224
	} else {
		goto L225
	}
L224:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v793))) = v814
	goto L226
L225:
	;
	goto L226
L226:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v787)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v787)+4)) = v801
	*(*int32)(unsafe.Add(mBase, uint32(v787))) = int32(1247)
	m.G0 = v796 + int32(16)
	goto L213
L227:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v833 = m.ExcPending
	if v833 != 0 {
		goto L18
	} else {
		goto L228
	}
L228:
	;
	v834 = F_format_type_be(m, v801)
	mBase = m.M
	v835 = m.ExcPending
	if v835 != 0 {
		goto L18
	} else {
		goto L229
	}
L229:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v796))) = v834
	F_errmsg(m, int32(_a_F_ExecAlterObjectSchemaStmt_19), v796)
	mBase = m.M
	v839 = m.ExcPending
	if v839 != 0 {
		goto L18
	} else {
		goto L230
	}
L230:
	;
	F_errfinish(m, int32(_a_F_ExecAlterObjectSchemaStmt_20), int32(_a_F_ExecAlterObjectSchemaStmt_21), int32(_a_F_ExecAlterObjectSchemaStmt_22))
	mBase = m.M
	v844 = m.ExcPending
	if v844 != 0 {
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
	v849 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v25))) = v849
	F_errmsg_internal(m, int32(_a_F_ExecAlterObjectSchemaStmt_23), v25)
	mBase = m.M
	v853 = m.ExcPending
	if v853 != 0 {
		goto L18
	} else {
		goto L233
	}
L233:
	;
	F_errfinish(m, int32(_a_F_ExecAlterObjectSchemaStmt_24), int32(594), int32(_a_F_ExecAlterObjectSchemaStmt_25))
	mBase = m.M
	v858 = m.ExcPending
	if v858 != 0 {
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
	v867 = *(*int32)(unsafe.Add(mBase, uint32(v25)+24))
	v868 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
	v869 = *(*int32)(unsafe.Add(mBase, uint32(v25)+16))
	v871 = F_table_open(m, v869, int32(3))
	mBase = m.M
	v872 = m.ExcPending
	if v872 != 0 {
		goto L18
	} else {
		goto L236
	}
L236:
	;
	v873 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v874 = F_LookupCreationNamespace(m, v873)
	mBase = m.M
	v875 = m.ExcPending
	if v875 != 0 {
		goto L18
	} else {
		goto L237
	}
L237:
	;
	v876 = F_AlterObjectNamespace_internal(m, v871, v868, v874)
	mBase = m.M
	v877 = m.ExcPending
	if v877 != 0 {
		goto L18
	} else {
		goto L238
	}
L238:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+28)) = v876
	F_relation_close(m, v871, int32(3))
	mBase = m.M
	v881 = m.ExcPending
	if v881 != 0 {
		goto L18
	} else {
		goto L239
	}
L239:
	;
	v910 = v869
	v911 = v868
	v912 = v867
	goto L1
L240:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(2615)
	v931 = *(*int32)(unsafe.Add(mBase, uint32(v25)+28))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v931
	goto L242
L241:
	;
	goto L242
L242:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v912
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v911
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v910
	m.G0 = v25 + int32(32)
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
