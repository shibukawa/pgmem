package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_PGLC_localeconv(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
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
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
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
	var v52 int32
	_ = v52
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v198 int32
	_ = v198
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v225 int32
	_ = v225
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v238 int32
	_ = v238
	var v242 int32
	_ = v242
	var v247 int32
	_ = v247
	var v251 int64
	_ = v251
	var v253 int64
	_ = v253
	var v255 int64
	_ = v255
	var v266 int32
	_ = v266
	var v276 int32
	_ = v276
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v287 int32
	_ = v287
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v312 int32
	_ = v312
	var v315 int32
	_ = v315
	var v318 int32
	_ = v318
	var v323 int32
	_ = v323
	var v326 int32
	_ = v326
	var v331 int32
	_ = v331
	var v334 int32
	_ = v334
	var v336 int32
	_ = v336
	var v342 int32
	_ = v342
	var v350 int32
	_ = v350
	var v353 int32
	_ = v353
	var v357 int32
	_ = v357
	var v361 int32
	_ = v361
	var v369 int32
	_ = v369
	var v374 int32
	_ = v374
	var v377 int32
	_ = v377
	var v382 int32
	_ = v382
	var v386 int32
	_ = v386
	var v391 int32
	_ = v391
	var v395 int64
	_ = v395
	var v397 int64
	_ = v397
	var v399 int64
	_ = v399
	var v410 int32
	_ = v410
	var v419 int32
	_ = v419
	var v422 int64
	_ = v422
	var v439 int32
	_ = v439
	var v444 int32
	_ = v444
	var v449 int32
	_ = v449
	var v453 int32
	_ = v453
	var v477 int32
	_ = v477
	var v480 int32
	_ = v480
	var v485 int32
	_ = v485
	var v490 int32
	_ = v490
	var v493 int32
	_ = v493
	var v496 int32
	_ = v496
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v511 int32
	_ = v511
	var v535 int32
	_ = v535
	var v538 int32
	_ = v538
	var v543 int32
	_ = v543
	var v545 int32
	_ = v545
	var v548 int32
	_ = v548
	var v554 int32
	_ = v554
	var v560 int32
	_ = v560
	var v570 int32
	_ = v570
	var v580 int32
	_ = v580
	var v604 int32
	_ = v604
	var v607 int32
	_ = v607
	var v612 int32
	_ = v612
	var v617 int32
	_ = v617
	var v620 int32
	_ = v620
	var v623 int32
	_ = v623
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v630 int32
	_ = v630
	var v631 int32
	_ = v631
	var v638 int32
	_ = v638
	var v662 int32
	_ = v662
	var v665 int32
	_ = v665
	var v670 int32
	_ = v670
	var v672 int32
	_ = v672
	var v675 int32
	_ = v675
	var v681 int32
	_ = v681
	var v687 int32
	_ = v687
	var v747 int32
	_ = v747
	var v755 int32
	_ = v755
	var v761 int32
	_ = v761
	var v763 int32
	_ = v763
	var v791 int32
	_ = v791
	var v807 int32
	_ = v807
	var v821 int32
	_ = v821
	var v823 int32
	_ = v823
	var v828 int32
	_ = v828
	var v845 int32
	_ = v845
	var v846 int32
	_ = v846
	var v861 int32
	_ = v861
	var v863 int32
	_ = v863
	var v864 int32
	_ = v864
	var v868 int32
	_ = v868
	var v869 int32
	_ = v869
	var v871 int32
	_ = v871
	var v886 int32
	_ = v886
	var v888 int32
	_ = v888
	var v889 int32
	_ = v889
	var v893 int32
	_ = v893
	var v894 int32
	_ = v894
	var v896 int32
	_ = v896
	var v906 int32
	_ = v906
	var v913 int32
	_ = v913
	var v915 int32
	_ = v915
	var v916 int32
	_ = v916
	var v920 int32
	_ = v920
	var v921 int32
	_ = v921
	var v923 int32
	_ = v923
	var v938 int32
	_ = v938
	var v940 int32
	_ = v940
	var v941 int32
	_ = v941
	var v945 int32
	_ = v945
	var v946 int32
	_ = v946
	var v948 int32
	_ = v948
	var v957 int32
	_ = v957
	var v965 int32
	_ = v965
	var v967 int32
	_ = v967
	var v968 int32
	_ = v968
	var v972 int32
	_ = v972
	var v973 int32
	_ = v973
	var v975 int32
	_ = v975
	var v983 int32
	_ = v983
	var v992 int32
	_ = v992
	var v994 int32
	_ = v994
	var v995 int32
	_ = v995
	var v999 int32
	_ = v999
	var v1000 int32
	_ = v1000
	var v1002 int32
	_ = v1002
	var v1009 int32
	_ = v1009
	var v1019 int32
	_ = v1019
	var v1021 int32
	_ = v1021
	var v1022 int32
	_ = v1022
	var v1026 int32
	_ = v1026
	var v1027 int32
	_ = v1027
	var v1029 int32
	_ = v1029
	var v1035 int32
	_ = v1035
	var v1046 int32
	_ = v1046
	var v1048 int32
	_ = v1048
	var v1049 int32
	_ = v1049
	var v1053 int32
	_ = v1053
	var v1054 int32
	_ = v1054
	var v1056 int32
	_ = v1056
	var v1071 int32
	_ = v1071
	var v1073 int32
	_ = v1073
	var v1074 int32
	_ = v1074
	var v1078 int32
	_ = v1078
	var v1079 int32
	_ = v1079
	var v1081 int32
	_ = v1081
	var v1086 int32
	_ = v1086
	var v1098 int32
	_ = v1098
	var v1100 int32
	_ = v1100
	var v1101 int32
	_ = v1101
	var v1105 int32
	_ = v1105
	var v1106 int32
	_ = v1106
	var v1108 int32
	_ = v1108
	var v1110 int32
	_ = v1110
	var v1112 int32
	_ = v1112
	var v1114 int32
	_ = v1114
	var v1116 int32
	_ = v1116
	var v1118 int32
	_ = v1118
	var v1120 int32
	_ = v1120
	var v1122 int32
	_ = v1122
	var v1127 int32
	_ = v1127
	var v1163 int32
	_ = v1163
	var v1166 int32
	_ = v1166
	var v1169 int32
	_ = v1169
	var v1174 int32
	_ = v1174
	var v1176 int32
	_ = v1176
	var v1179 int32
	_ = v1179
	var v1215 int32
	_ = v1215
	var v1230 int32
	_ = v1230
	var v1246 int32
	_ = v1246
	var v1263 int32
	_ = v1263
	var v1266 int32
	_ = v1266
	var v1268 int32
	_ = v1268
	var v1275 int32
	_ = v1275
	var v1276 int32
	_ = v1276
	var v1277 int32
	_ = v1277
	var v1278 int32
	_ = v1278
	var v1279 int32
	_ = v1279
	var v1280 int32
	_ = v1280
	var v1281 int32
	_ = v1281
	var v1282 int32
	_ = v1282
	var v1283 int32
	_ = v1283
	var v1284 int32
	_ = v1284
	var v1285 int32
	_ = v1285
	var v1286 int32
	_ = v1286
	var v1287 int32
	_ = v1287
	var v1294 int32
	_ = v1294
	var v1315 int32
	_ = v1315
	var v1317 int32
	_ = v1317
	var v1318 int32
	_ = v1318
	var v1331 int32
	_ = v1331
	var v1334 int32
	_ = v1334
	var v1336 int32
	_ = v1336
	var v1350 int32
	_ = v1350
	var v1364 int32
	_ = v1364
	var v1366 int32
	_ = v1366
	var v1367 int32
	_ = v1367
	var v1380 int32
	_ = v1380
	var v1383 int32
	_ = v1383
	var v1385 int32
	_ = v1385
	var v1399 int32
	_ = v1399
	var v1413 int32
	_ = v1413
	var v1427 int32
	_ = v1427
	var v1441 int32
	_ = v1441
	var v1455 int32
	_ = v1455
	var v1461 int64
	_ = v1461
	var v1464 int64
	_ = v1464
	var v1467 int64
	_ = v1467
	var v1470 int64
	_ = v1470
	var v1473 int64
	_ = v1473
	var v1476 int64
	_ = v1476
	var v1479 int64
	_ = v1479
	var v1482 int32
	_ = v1482
	var v1534 int32
	_ = v1534
	var v1536 int32
	_ = v1536
	var v1538 int32
	_ = v1538
	var v1540 int32
	_ = v1540
	var v1542 int32
	_ = v1542
	var v1544 int32
	_ = v1544
	var v1546 int32
	_ = v1546
	var v1548 int32
	_ = v1548
	var v1550 int32
	_ = v1550
	var v1552 int32
	_ = v1552
	var v1567 int32
	_ = v1567
	var v1588 int32
	_ = v1588
	var v1594 int32
	_ = v1594
	var v1595 int64
	_ = v1595
	var v1599 int32
	_ = v1599
	var v1601 int32
	_ = v1601
	var v1602 int32
	_ = v1602
	var v1606 int32
	_ = v1606
	var v1608 int32
	_ = v1608
	var v1609 int32
	_ = v1609
	var v1610 int32
	_ = v1610
	var v1611 int32
	_ = v1611
	var v1612 int32
	_ = v1612
	var v1613 int32
	_ = v1613
	var v1614 int32
	_ = v1614
	var v1615 int32
	_ = v1615
	var v1616 int32
	_ = v1616
	var v1617 int32
	_ = v1617
	var v1618 int32
	_ = v1618
	var v1619 int32
	_ = v1619
	var v1620 int32
	_ = v1620
	var v1621 int32
	_ = v1621
	var v1625 int32
	_ = v1625
	v1 = int32(0)
	v27 = m.G0
	v29 = v27 + int32(-64)
	m.G0 = v29
	v33 = v1
	v34 = v1
	v35 = v1
	v36 = v1
	v37 = v1
	v38 = v1
	v39 = v1
	v40 = v1
	v41 = v1
	v42 = v1
	v43 = v1
	v44 = v1
	v45 = v1
	v46 = int32(-1)
	v52 = v29
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
	if v46 != int32(1) {
		goto L9
	} else {
		goto L10
	}
L5:
	;
	goto L3
L6:
	;
	v1594 = int32(m.ExcTag)
	v1595 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v1594 == int32(0) {
		goto L255
	} else {
		goto L256
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, _consts[52])) = v1280
	*(*int32)(unsafe.Add(mBase, _consts[53])) = v1279
	*(*int32)(unsafe.Add(mBase, uint32(v29)+16)) = v1280
	*(*int32)(unsafe.Add(mBase, uint32(v29)+20)) = v1279
	*(*int32)(unsafe.Add(mBase, uint32(v29)+24)) = v1281
	*(*int32)(unsafe.Add(mBase, uint32(v29)+28)) = v1282
	*(*int32)(unsafe.Add(mBase, uint32(v29)+32)) = v1283
	*(*int32)(unsafe.Add(mBase, uint32(v29)+36)) = v1284
	*(*int32)(unsafe.Add(mBase, uint32(v29)+40)) = v1285
	*(*int32)(unsafe.Add(mBase, uint32(v29)+44)) = v1286
	*(*int32)(unsafe.Add(mBase, uint32(v29)+48)) = v1287
	*(*int32)(unsafe.Add(mBase, uint32(v29)+52)) = v1278
	*(*int32)(unsafe.Add(mBase, uint32(v29)+56)) = v1275
	*(*int32)(unsafe.Add(mBase, uint32(v29)+60)) = v1276
	v1534 = *(*int32)(unsafe.Add(mBase, uint32(v1275)))
	F_emscripten_builtin_free(m, v1534)
	mBase = m.M
	v1536 = *(*int32)(unsafe.Add(mBase, uint32(v1275)+4))
	F_emscripten_builtin_free(m, v1536)
	mBase = m.M
	v1538 = *(*int32)(unsafe.Add(mBase, uint32(v1275)+8))
	F_emscripten_builtin_free(m, v1538)
	mBase = m.M
	v1540 = *(*int32)(unsafe.Add(mBase, uint32(v1275)+12))
	F_emscripten_builtin_free(m, v1540)
	mBase = m.M
	v1542 = *(*int32)(unsafe.Add(mBase, uint32(v1275)+16))
	F_emscripten_builtin_free(m, v1542)
	mBase = m.M
	v1544 = *(*int32)(unsafe.Add(mBase, uint32(v1275)+20))
	F_emscripten_builtin_free(m, v1544)
	mBase = m.M
	v1546 = *(*int32)(unsafe.Add(mBase, uint32(v1275)+24))
	F_emscripten_builtin_free(m, v1546)
	mBase = m.M
	v1548 = *(*int32)(unsafe.Add(mBase, uint32(v1275)+28))
	F_emscripten_builtin_free(m, v1548)
	mBase = m.M
	v1550 = *(*int32)(unsafe.Add(mBase, uint32(v1275)+32))
	F_emscripten_builtin_free(m, v1550)
	mBase = m.M
	v1552 = *(*int32)(unsafe.Add(mBase, uint32(v1275)+36))
	F_emscripten_builtin_free(m, v1552)
	mBase = m.M
	goto L253
L8:
	;
	m.G0 = v29 - int32(-64)
	return int32(4538692)
L9:
	;
	v61 = v52 + int32(-64)
	m.G0 = v61
	v64 = v52 + int32(-128)
	m.G0 = v64
	v67 = v64 - int32(160)
	m.G0 = v67
	*(*int64)(unsafe.Add(mBase, uint32(v64)+48)) = int64(0)
	v72 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1005])))
	if v72 != 0 {
		goto L8
	} else {
		goto L12
	}
L10:
	;
	v1275 = v33
	v1276 = v34
	v1277 = v35
	v1278 = v36
	v1279 = v37
	v1280 = v38
	v1281 = v39
	v1282 = v40
	v1283 = v41
	v1284 = v42
	v1285 = v43
	v1286 = v44
	v1287 = v45
	v1294 = v52
	goto L11
L11:
	;
	if v1277 != 0 {
		goto L7
	} else {
		goto L236
	}
L12:
	;
	v74 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1006])))
	if v74 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+20)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v29)+16)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v29)+24)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v29)+28)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v29)+32)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v29)+36)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v29)+40)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v29)+44)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v29)+48)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v29)+52)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v29)+56)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v29)+60)) = v61
	v87 = int32(4538692)
	v88 = *(*int32)(unsafe.Add(mBase, _consts[1007]))
	F_emscripten_builtin_free(m, v88)
	mBase = m.M
	v90 = *(*int32)(unsafe.Add(mBase, _consts[1008]))
	F_emscripten_builtin_free(m, v90)
	mBase = m.M
	v92 = *(*int32)(unsafe.Add(mBase, _consts[1009]))
	F_emscripten_builtin_free(m, v92)
	mBase = m.M
	v94 = *(*int32)(unsafe.Add(mBase, _consts[1010]))
	F_emscripten_builtin_free(m, v94)
	mBase = m.M
	v96 = *(*int32)(unsafe.Add(mBase, _consts[1011]))
	F_emscripten_builtin_free(m, v96)
	mBase = m.M
	v98 = *(*int32)(unsafe.Add(mBase, _consts[1012]))
	F_emscripten_builtin_free(m, v98)
	mBase = m.M
	v100 = *(*int32)(unsafe.Add(mBase, _consts[1013]))
	F_emscripten_builtin_free(m, v100)
	mBase = m.M
	v102 = *(*int32)(unsafe.Add(mBase, _consts[1014]))
	F_emscripten_builtin_free(m, v102)
	mBase = m.M
	v104 = *(*int32)(unsafe.Add(mBase, _consts[1015]))
	F_emscripten_builtin_free(m, v104)
	mBase = m.M
	v106 = *(*int32)(unsafe.Add(mBase, _consts[1016]))
	F_emscripten_builtin_free(m, v106)
	mBase = m.M
	goto L16
L14:
	;
	goto L15
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+16)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v29)+20)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v29)+24)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v29)+28)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v29)+32)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v29)+36)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v29)+40)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v29)+44)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v29)+48)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v29)+52)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v29)+56)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v29)+60)) = v61
	v124 = *(*int32)(unsafe.Add(mBase, _consts[1017]))
	v126 = *(*int32)(unsafe.Add(mBase, _consts[1018]))
	*(*int32)(unsafe.Add(mBase, _consts[43])) = int32(44)
	v132 = int32(0)
	v136 = m.G0
	v138 = v136 - int32(32)
	m.G0 = v138
	v143 = v132
	goto L21
L16:
	;
	v109 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _consts[1006])) = uint8(v109)
	goto L15
L17:
	;
	if v791 != 0 {
		goto L164
	} else {
		goto L165
	}
L18:
	;
	if v266 == int32(0) {
		v791 = int32(-1)
		goto L17
	} else {
		goto L46
	}
L19:
	;
	m.G0 = v138 + int32(32)
	goto L18
L20:
	;
	v266 = int32(0)
	goto L19
L21:
	;
	goto L24
L22:
	;
	v171 = F___loc_is_allocated(m, v132)
	mBase = m.M
	if v171 == int32(0) {
		goto L32
	} else {
		goto L33
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v138+int32(8)+v143<<(uint(int32(2))%32)))) = v162
	if v162 == int32(-1) {
		goto L20
	} else {
		goto L30
	}
L24:
	;
	if int32(1)<<(uint(v143)%32)&int32(2147483647) != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v161 = v124
	goto L29
L28:
	;
	v161 = int32(790160)
	goto L29
L29:
	;
	v162 = F___get_locale(m, v143, v161)
	mBase = m.M
	goto L23
L30:
	;
	v168 = v143 + int32(1)
	if v168 != int32(6) {
		v143 = v168
		goto L21
	} else {
		goto L31
	}
L31:
	;
	goto L22
L32:
	;
	v174 = int32(4135160)
	v179 = F_memcmp(m, v138+int32(8), v174, int32(24))
	mBase = m.M
	if v179 == int32(0) {
		v266 = v174
		goto L19
	} else {
		goto L35
	}
L33:
	;
	v247 = v132
	goto L34
L34:
	;
	v251 = *(*int64)(unsafe.Add(mBase, uint32(v138)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v247))) = v251
	v253 = *(*int64)(unsafe.Add(mBase, uint32(v138)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v247)+16)) = v253
	v255 = *(*int64)(unsafe.Add(mBase, uint32(v138)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v247)+8)) = v255
	v266 = v247
	goto L19
L35:
	;
	v182 = int32(4135184)
	v187 = F_memcmp(m, v138+int32(8), v182, int32(24))
	mBase = m.M
	if v187 == int32(0) {
		v266 = v182
		goto L19
	} else {
		goto L36
	}
L36:
	;
	v190 = int32(0)
	v192 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1019])))
	if v192 == v190 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v198 = v190
	goto L40
L38:
	;
	goto L39
L39:
	;
	v225 = int32(4719012)
	v230 = F_memcmp(m, v138+int32(8), v225, int32(24))
	mBase = m.M
	if v230 == int32(0) {
		v266 = v225
		goto L19
	} else {
		goto L43
	}
L40:
	;
	v206 = F___get_locale(m, v198, int32(790160))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v198<<(uint(int32(2))%32))+uint32(_consts[1020]))) = v206
	v209 = v198 + int32(1)
	if v209 != int32(6) {
		v198 = v209
		goto L40
	} else {
		goto L42
	}
L41:
	;
	v213 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[1019])) = uint8(v213)
	v217 = *(*int32)(unsafe.Add(mBase, _consts[1020]))
	*(*int32)(unsafe.Add(mBase, _consts[1021])) = v217
	goto L39
L42:
	;
	goto L41
L43:
	;
	v233 = int32(4719036)
	v238 = F_memcmp(m, v138+int32(8), v233, int32(24))
	mBase = m.M
	if v238 == int32(0) {
		v266 = v233
		goto L19
	} else {
		goto L44
	}
L44:
	;
	v242 = F_emscripten_builtin_malloc(m, int32(24))
	mBase = m.M
	if v242 == int32(0) {
		goto L20
	} else {
		goto L45
	}
L45:
	;
	v247 = v242
	goto L34
L46:
	;
	v276 = int32(0)
	v280 = m.G0
	v282 = v280 - int32(32)
	m.G0 = v282
	v287 = v276
	goto L50
L47:
	;
	if v410 == int32(0) {
		goto L75
	} else {
		goto L76
	}
L48:
	;
	m.G0 = v282 + int32(32)
	goto L47
L49:
	;
	v410 = int32(0)
	goto L48
L50:
	;
	goto L53
L51:
	;
	v315 = F___loc_is_allocated(m, v276)
	mBase = m.M
	if v315 == int32(0) {
		goto L61
	} else {
		goto L62
	}
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v282+int32(8)+v287<<(uint(int32(2))%32)))) = v306
	if v306 == int32(-1) {
		goto L49
	} else {
		goto L59
	}
L53:
	;
	if int32(1)<<(uint(v287)%32)&int32(2147483647) != 0 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v305 = v126
	goto L58
L57:
	;
	v305 = int32(790160)
	goto L58
L58:
	;
	v306 = F___get_locale(m, v287, v305)
	mBase = m.M
	goto L52
L59:
	;
	v312 = v287 + int32(1)
	if v312 != int32(6) {
		v287 = v312
		goto L50
	} else {
		goto L60
	}
L60:
	;
	goto L51
L61:
	;
	v318 = int32(4135160)
	v323 = F_memcmp(m, v282+int32(8), v318, int32(24))
	mBase = m.M
	if v323 == int32(0) {
		v410 = v318
		goto L48
	} else {
		goto L64
	}
L62:
	;
	v391 = v276
	goto L63
L63:
	;
	v395 = *(*int64)(unsafe.Add(mBase, uint32(v282)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v391))) = v395
	v397 = *(*int64)(unsafe.Add(mBase, uint32(v282)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v391)+16)) = v397
	v399 = *(*int64)(unsafe.Add(mBase, uint32(v282)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v391)+8)) = v399
	v410 = v391
	goto L48
L64:
	;
	v326 = int32(4135184)
	v331 = F_memcmp(m, v282+int32(8), v326, int32(24))
	mBase = m.M
	if v331 == int32(0) {
		v410 = v326
		goto L48
	} else {
		goto L65
	}
L65:
	;
	v334 = int32(0)
	v336 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1019])))
	if v336 == v334 {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v342 = v334
	goto L69
L67:
	;
	goto L68
L68:
	;
	v369 = int32(4719012)
	v374 = F_memcmp(m, v282+int32(8), v369, int32(24))
	mBase = m.M
	if v374 == int32(0) {
		v410 = v369
		goto L48
	} else {
		goto L72
	}
L69:
	;
	v350 = F___get_locale(m, v342, int32(790160))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v342<<(uint(int32(2))%32))+uint32(_consts[1020]))) = v350
	v353 = v342 + int32(1)
	if v353 != int32(6) {
		v342 = v353
		goto L69
	} else {
		goto L71
	}
L70:
	;
	v357 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[1019])) = uint8(v357)
	v361 = *(*int32)(unsafe.Add(mBase, _consts[1020]))
	*(*int32)(unsafe.Add(mBase, _consts[1021])) = v361
	goto L68
L71:
	;
	goto L70
L72:
	;
	v377 = int32(4719036)
	v382 = F_memcmp(m, v282+int32(8), v377, int32(24))
	mBase = m.M
	if v382 == int32(0) {
		v410 = v377
		goto L48
	} else {
		goto L73
	}
L73:
	;
	v386 = F_emscripten_builtin_malloc(m, int32(24))
	mBase = m.M
	if v386 == int32(0) {
		goto L49
	} else {
		goto L74
	}
L74:
	;
	v391 = v386
	goto L63
L75:
	;
	v419 = F___loc_is_allocated(m, v266)
	mBase = m.M
	if v419 != 0 {
		goto L79
	} else {
		goto L80
	}
L76:
	;
	goto L77
L77:
	;
	v422 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v61))) = v422
	*(*int64)(unsafe.Add(mBase, uint32(v61)+48)) = v422
	*(*int64)(unsafe.Add(mBase, uint32(v61)+40)) = v422
	*(*int64)(unsafe.Add(mBase, uint32(v61)+32)) = v422
	*(*int64)(unsafe.Add(mBase, uint32(v61)+24)) = v422
	*(*int64)(unsafe.Add(mBase, uint32(v61)+16)) = v422
	*(*int64)(unsafe.Add(mBase, uint32(v61)+8)) = v422
	v439 = *(*int32)(unsafe.Add(mBase, _consts[1022]))
	if v266 != 0 {
		goto L83
	} else {
		goto L84
	}
L78:
	;
	v791 = int32(-1)
	goto L17
L79:
	;
	F_emscripten_builtin_free(m, v266)
	mBase = m.M
	goto L81
L80:
	;
	goto L81
L81:
	;
	goto L78
L82:
	;
	v453 = int32(0)
	goto L94
L83:
	;
	if v266 == int32(-1) {
		goto L86
	} else {
		goto L87
	}
L84:
	;
	goto L85
L85:
	;
	if v439 == int32(4718872) {
		goto L89
	} else {
		goto L90
	}
L86:
	;
	v444 = int32(4718872)
	goto L88
L87:
	;
	v444 = v266
	goto L88
L88:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1022])) = v444
	goto L85
L89:
	;
	v449 = int32(-1)
	goto L91
L90:
	;
	v449 = v439
	goto L91
L91:
	;
	goto L82
L92:
	;
	if v449 != 0 {
		goto L147
	} else {
		goto L148
	}
L93:
	;
	*(*int32)(unsafe.Add(mBase, _consts[43])) = int32(48)
	v747 = int32(-1)
	goto L92
L94:
	;
	v477 = v453 * int32(12)
	v480 = *(*int32)(unsafe.Add(mBase, uint32(v477)+uint32(_consts[1023])))
	if v480 != int32(4) {
		goto L96
	} else {
		goto L97
	}
L95:
	;
	if v410 != 0 {
		goto L116
	} else {
		goto L117
	}
L96:
	;
	v560 = v453 + int32(1)
	if v560 != int32(18) {
		v453 = v560
		goto L94
	} else {
		goto L114
	}
L97:
	;
	v485 = *(*int32)(unsafe.Add(mBase, uint32(v477)+uint32(_consts[1024])))
	v490 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v477)+uint32(_consts[1025]))))
	if v490 == int32(1) {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v493 = *(*int32)(unsafe.Add(mBase, uint32(v485)+uint32(_consts[1026])))
	v496 = F_strlen(m, v493)
	mBase = m.M
	v498 = v496 + int32(1)
	v499 = F_emscripten_builtin_malloc(m, v498)
	mBase = m.M
	if v499 == int32(0) {
		goto L102
	} else {
		goto L103
	}
L99:
	;
	goto L100
L100:
	;
	v554 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v485)+uint32(_consts[1026]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v61+v485))) = uint8(v554)
	goto L96
L101:
	;
	if v504 == int32(0) {
		goto L105
	} else {
		goto L106
	}
L102:
	;
	v504 = int32(0)
	goto L101
L103:
	;
	goto L104
L104:
	;
	v503 = F___memcpy(m, v499, v493, v498)
	mBase = m.M
	v504 = v503
	goto L101
L105:
	;
	v511 = int32(0)
	goto L108
L106:
	;
	goto L107
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v61+v485))) = v504
	goto L96
L108:
	;
	v535 = v511 * int32(12)
	v538 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v535)+uint32(_consts[1025]))))
	if v538 == int32(1) {
		goto L110
	} else {
		goto L111
	}
L109:
	;
	goto L93
L110:
	;
	v543 = *(*int32)(unsafe.Add(mBase, uint32(v535)+uint32(_consts[1024])))
	v545 = *(*int32)(unsafe.Add(mBase, uint32(v61+v543)))
	F_emscripten_builtin_free(m, v545)
	mBase = m.M
	goto L112
L111:
	;
	goto L112
L112:
	;
	v548 = v511 + int32(1)
	if v548 != int32(18) {
		v511 = v548
		goto L108
	} else {
		goto L113
	}
L113:
	;
	goto L109
L114:
	;
	goto L95
L115:
	;
	v580 = int32(0)
	goto L125
L116:
	;
	if v410 == int32(-1) {
		goto L119
	} else {
		goto L120
	}
L117:
	;
	goto L118
L118:
	;
	goto L122
L119:
	;
	v570 = int32(4718872)
	goto L121
L120:
	;
	v570 = v410
	goto L121
L121:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1022])) = v570
	goto L118
L122:
	;
	goto L124
L124:
	;
	goto L115
L125:
	;
	v604 = v580 * int32(12)
	v607 = *(*int32)(unsafe.Add(mBase, uint32(v604)+uint32(_consts[1023])))
	if v607 != int32(1) {
		goto L127
	} else {
		goto L128
	}
L126:
	;
	v747 = int32(0)
	goto L92
L127:
	;
	v687 = v580 + int32(1)
	if v687 != int32(18) {
		v580 = v687
		goto L125
	} else {
		goto L145
	}
L128:
	;
	v612 = *(*int32)(unsafe.Add(mBase, uint32(v604)+uint32(_consts[1024])))
	v617 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v604)+uint32(_consts[1025]))))
	if v617 == int32(1) {
		goto L129
	} else {
		goto L130
	}
L129:
	;
	v620 = *(*int32)(unsafe.Add(mBase, uint32(v612)+uint32(_consts[1026])))
	v623 = F_strlen(m, v620)
	mBase = m.M
	v625 = v623 + int32(1)
	v626 = F_emscripten_builtin_malloc(m, v625)
	mBase = m.M
	if v626 == int32(0) {
		goto L133
	} else {
		goto L134
	}
L130:
	;
	goto L131
L131:
	;
	v681 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v612)+uint32(_consts[1026]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v61+v612))) = uint8(v681)
	goto L127
L132:
	;
	if v631 == int32(0) {
		goto L136
	} else {
		goto L137
	}
L133:
	;
	v631 = int32(0)
	goto L132
L134:
	;
	goto L135
L135:
	;
	v630 = F___memcpy(m, v626, v620, v625)
	mBase = m.M
	v631 = v630
	goto L132
L136:
	;
	v638 = int32(0)
	goto L139
L137:
	;
	goto L138
L138:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v61+v612))) = v631
	goto L127
L139:
	;
	v662 = v638 * int32(12)
	v665 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v662)+uint32(_consts[1025]))))
	if v665 == int32(1) {
		goto L141
	} else {
		goto L142
	}
L140:
	;
	goto L93
L141:
	;
	v670 = *(*int32)(unsafe.Add(mBase, uint32(v662)+uint32(_consts[1024])))
	v672 = *(*int32)(unsafe.Add(mBase, uint32(v61+v670)))
	F_emscripten_builtin_free(m, v672)
	mBase = m.M
	goto L143
L142:
	;
	goto L143
L143:
	;
	v675 = v638 + int32(1)
	if v675 != int32(18) {
		v638 = v675
		goto L139
	} else {
		goto L144
	}
L144:
	;
	goto L140
L145:
	;
	goto L126
L146:
	;
	v761 = F___loc_is_allocated(m, v266)
	mBase = m.M
	if v761 != 0 {
		goto L157
	} else {
		goto L158
	}
L147:
	;
	if v449 == int32(-1) {
		goto L150
	} else {
		goto L151
	}
L148:
	;
	goto L149
L149:
	;
	goto L153
L150:
	;
	v755 = int32(4718872)
	goto L152
L151:
	;
	v755 = v449
	goto L152
L152:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1022])) = v755
	goto L149
L153:
	;
	goto L155
L155:
	;
	goto L146
L156:
	;
	v763 = F___loc_is_allocated(m, v410)
	mBase = m.M
	if v763 != 0 {
		goto L161
	} else {
		goto L162
	}
L157:
	;
	F_emscripten_builtin_free(m, v266)
	mBase = m.M
	goto L159
L158:
	;
	goto L159
L159:
	;
	goto L156
L160:
	;
	v791 = v747
	goto L17
L161:
	;
	F_emscripten_builtin_free(m, v410)
	mBase = m.M
	goto L163
L162:
	;
	goto L163
L163:
	;
	goto L160
L164:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+20)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v29)+16)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v29)+24)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v29)+28)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v29)+32)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v29)+36)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v29)+40)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v29)+44)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v29)+48)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v29)+52)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v29)+56)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v29)+60)) = v61
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v807 = m.ExcPending
	if v807 != 0 {
		v1588 = v67
		goto L6
	} else {
		goto L167
	}
L165:
	;
	goto L166
L166:
	;
	v846 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
	*(*int32)(unsafe.Add(mBase, uint32(v29)+20)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v29)+16)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v29)+24)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v29)+28)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v29)+32)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v29)+36)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v29)+40)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v29)+44)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v29)+48)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v29)+52)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v29)+56)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v29)+60)) = v61
	v861 = F_strlen(m, v846)
	mBase = m.M
	v863 = v861 + int32(1)
	v864 = F_emscripten_builtin_malloc(m, v863)
	mBase = m.M
	if v864 == int32(0) {
		goto L171
	} else {
		goto L172
	}
L167:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+16)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v29)+20)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v29)+24)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v29)+28)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v29)+32)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v29)+36)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v29)+40)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v29)+44)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v29)+48)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v29)+52)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v29)+56)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v29)+60)) = v61
	v821 = *(*int32)(unsafe.Add(mBase, _consts[1017]))
	v823 = *(*int32)(unsafe.Add(mBase, _consts[1018]))
	*(*int32)(unsafe.Add(mBase, uint32(v29)+4)) = v823
	*(*int32)(unsafe.Add(mBase, uint32(v29))) = v821
	F_errmsg_internal(m, int32(313443), v29)
	mBase = m.M
	v828 = m.ExcPending
	if v828 != 0 {
		v1588 = v67
		goto L6
	} else {
		goto L168
	}
L168:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+20)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v29)+16)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v29)+24)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v29)+28)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v29)+32)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v29)+36)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v29)+40)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v29)+44)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v29)+48)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v29)+52)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v29)+56)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v29)+60)) = v61
	F_errfinish(m, int32(522638), int32(560), int32(34006))
	mBase = m.M
	v845 = m.ExcPending
	if v845 != 0 {
		v1588 = v67
		goto L6
	} else {
		goto L169
	}
L169:
	;
	goto L3
L170:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v64))) = v869
	v871 = *(*int32)(unsafe.Add(mBase, uint32(v61)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v29)+20)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v29)+16)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v29)+24)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v29)+28)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v29)+32)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v29)+36)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v29)+40)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v29)+44)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v29)+48)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v29)+52)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v29)+56)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v29)+60)) = v61
	v886 = F_strlen(m, v871)
	mBase = m.M
	v888 = v886 + int32(1)
	v889 = F_emscripten_builtin_malloc(m, v888)
	mBase = m.M
	if v889 == int32(0) {
		goto L175
	} else {
		goto L176
	}
L171:
	;
	v869 = int32(0)
	goto L170
L172:
	;
	goto L173
L173:
	;
	v868 = F___memcpy(m, v864, v846, v863)
	mBase = m.M
	v869 = v868
	goto L170
L174:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v64)+4)) = v894
	v896 = *(*int32)(unsafe.Add(mBase, uint32(v61)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v29)+20)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v29)+16)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v29)+24)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v29)+28)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v29)+32)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v29)+36)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v29)+40)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v29)+44)) = v44
	v906 = v52 + int32(-124)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+48)) = v906
	*(*int32)(unsafe.Add(mBase, uint32(v29)+52)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v29)+56)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v29)+60)) = v61
	v913 = F_strlen(m, v896)
	mBase = m.M
	v915 = v913 + int32(1)
	v916 = F_emscripten_builtin_malloc(m, v915)
	mBase = m.M
	if v916 == int32(0) {
		goto L179
	} else {
		goto L180
	}
L175:
	;
	v894 = int32(0)
	goto L174
L176:
	;
	goto L177
L177:
	;
	v893 = F___memcpy(m, v889, v871, v888)
	mBase = m.M
	v894 = v893
	goto L174
L178:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v64)+8)) = v921
	v923 = *(*int32)(unsafe.Add(mBase, uint32(v61)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v29)+20)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v29)+16)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v29)+24)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v29)+28)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v29)+32)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v29)+36)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v29)+40)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v29)+44)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v29)+48)) = v906
	*(*int32)(unsafe.Add(mBase, uint32(v29)+52)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v29)+56)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v29)+60)) = v61
	v938 = F_strlen(m, v923)
	mBase = m.M
	v940 = v938 + int32(1)
	v941 = F_emscripten_builtin_malloc(m, v940)
	mBase = m.M
	if v941 == int32(0) {
		goto L183
	} else {
		goto L184
	}
L179:
	;
	v921 = int32(0)
	goto L178
L180:
	;
	goto L181
L181:
	;
	v920 = F___memcpy(m, v916, v896, v915)
	mBase = m.M
	v921 = v920
	goto L178
L182:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v64)+12)) = v946
	v948 = *(*int32)(unsafe.Add(mBase, uint32(v61)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v29)+20)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v29)+16)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v29)+24)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v29)+28)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v29)+32)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v29)+36)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v29)+40)) = v43
	v957 = v52 + int32(-116)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+44)) = v957
	*(*int32)(unsafe.Add(mBase, uint32(v29)+48)) = v906
	*(*int32)(unsafe.Add(mBase, uint32(v29)+52)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v29)+56)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v29)+60)) = v61
	v965 = F_strlen(m, v948)
	mBase = m.M
	v967 = v965 + int32(1)
	v968 = F_emscripten_builtin_malloc(m, v967)
	mBase = m.M
	if v968 == int32(0) {
		goto L187
	} else {
		goto L188
	}
L183:
	;
	v946 = int32(0)
	goto L182
L184:
	;
	goto L185
L185:
	;
	v945 = F___memcpy(m, v941, v923, v940)
	mBase = m.M
	v946 = v945
	goto L182
L186:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v64)+16)) = v973
	v975 = *(*int32)(unsafe.Add(mBase, uint32(v61)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v29)+20)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v29)+16)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v29)+24)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v29)+28)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v29)+32)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v29)+36)) = v42
	v983 = v52 + int32(-112)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+40)) = v983
	*(*int32)(unsafe.Add(mBase, uint32(v29)+44)) = v957
	*(*int32)(unsafe.Add(mBase, uint32(v29)+48)) = v906
	*(*int32)(unsafe.Add(mBase, uint32(v29)+52)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v29)+56)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v29)+60)) = v61
	v992 = F_strlen(m, v975)
	mBase = m.M
	v994 = v992 + int32(1)
	v995 = F_emscripten_builtin_malloc(m, v994)
	mBase = m.M
	if v995 == int32(0) {
		goto L191
	} else {
		goto L192
	}
L187:
	;
	v973 = int32(0)
	goto L186
L188:
	;
	goto L189
L189:
	;
	v972 = F___memcpy(m, v968, v948, v967)
	mBase = m.M
	v973 = v972
	goto L186
L190:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v64)+20)) = v1000
	v1002 = *(*int32)(unsafe.Add(mBase, uint32(v61)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v29)+20)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v29)+16)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v29)+24)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v29)+28)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v29)+32)) = v41
	v1009 = v52 + int32(-108)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+36)) = v1009
	*(*int32)(unsafe.Add(mBase, uint32(v29)+40)) = v983
	*(*int32)(unsafe.Add(mBase, uint32(v29)+44)) = v957
	*(*int32)(unsafe.Add(mBase, uint32(v29)+48)) = v906
	*(*int32)(unsafe.Add(mBase, uint32(v29)+52)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v29)+56)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v29)+60)) = v61
	v1019 = F_strlen(m, v1002)
	mBase = m.M
	v1021 = v1019 + int32(1)
	v1022 = F_emscripten_builtin_malloc(m, v1021)
	mBase = m.M
	if v1022 == int32(0) {
		goto L195
	} else {
		goto L196
	}
L191:
	;
	v1000 = int32(0)
	goto L190
L192:
	;
	goto L193
L193:
	;
	v999 = F___memcpy(m, v995, v975, v994)
	mBase = m.M
	v1000 = v999
	goto L190
L194:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v64)+24)) = v1027
	v1029 = *(*int32)(unsafe.Add(mBase, uint32(v61)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v29)+20)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v29)+16)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v29)+24)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v29)+28)) = v40
	v1035 = v52 + int32(-104)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+32)) = v1035
	*(*int32)(unsafe.Add(mBase, uint32(v29)+36)) = v1009
	*(*int32)(unsafe.Add(mBase, uint32(v29)+40)) = v983
	*(*int32)(unsafe.Add(mBase, uint32(v29)+44)) = v957
	*(*int32)(unsafe.Add(mBase, uint32(v29)+48)) = v906
	*(*int32)(unsafe.Add(mBase, uint32(v29)+52)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v29)+56)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v29)+60)) = v61
	v1046 = F_strlen(m, v1029)
	mBase = m.M
	v1048 = v1046 + int32(1)
	v1049 = F_emscripten_builtin_malloc(m, v1048)
	mBase = m.M
	if v1049 == int32(0) {
		goto L199
	} else {
		goto L200
	}
L195:
	;
	v1027 = int32(0)
	goto L194
L196:
	;
	goto L197
L197:
	;
	v1026 = F___memcpy(m, v1022, v1002, v1021)
	mBase = m.M
	v1027 = v1026
	goto L194
L198:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v64)+28)) = v1054
	v1056 = *(*int32)(unsafe.Add(mBase, uint32(v61)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v29)+20)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v29)+16)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v29)+24)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v29)+28)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v29)+32)) = v1035
	*(*int32)(unsafe.Add(mBase, uint32(v29)+36)) = v1009
	*(*int32)(unsafe.Add(mBase, uint32(v29)+40)) = v983
	*(*int32)(unsafe.Add(mBase, uint32(v29)+44)) = v957
	*(*int32)(unsafe.Add(mBase, uint32(v29)+48)) = v906
	*(*int32)(unsafe.Add(mBase, uint32(v29)+52)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v29)+56)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v29)+60)) = v61
	v1071 = F_strlen(m, v1056)
	mBase = m.M
	v1073 = v1071 + int32(1)
	v1074 = F_emscripten_builtin_malloc(m, v1073)
	mBase = m.M
	if v1074 == int32(0) {
		goto L203
	} else {
		goto L204
	}
L199:
	;
	v1054 = int32(0)
	goto L198
L200:
	;
	goto L201
L201:
	;
	v1053 = F___memcpy(m, v1049, v1029, v1048)
	mBase = m.M
	v1054 = v1053
	goto L198
L202:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v64)+32)) = v1079
	v1081 = *(*int32)(unsafe.Add(mBase, uint32(v61)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v29)+20)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v29)+16)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v29)+24)) = v39
	v1086 = v52 + int32(-96)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+28)) = v1086
	*(*int32)(unsafe.Add(mBase, uint32(v29)+32)) = v1035
	*(*int32)(unsafe.Add(mBase, uint32(v29)+36)) = v1009
	*(*int32)(unsafe.Add(mBase, uint32(v29)+40)) = v983
	*(*int32)(unsafe.Add(mBase, uint32(v29)+44)) = v957
	*(*int32)(unsafe.Add(mBase, uint32(v29)+48)) = v906
	*(*int32)(unsafe.Add(mBase, uint32(v29)+52)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v29)+56)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v29)+60)) = v61
	v1098 = F_strlen(m, v1081)
	mBase = m.M
	v1100 = v1098 + int32(1)
	v1101 = F_emscripten_builtin_malloc(m, v1100)
	mBase = m.M
	if v1101 == int32(0) {
		goto L207
	} else {
		goto L208
	}
L203:
	;
	v1079 = int32(0)
	goto L202
L204:
	;
	goto L205
L205:
	;
	v1078 = F___memcpy(m, v1074, v1056, v1073)
	mBase = m.M
	v1079 = v1078
	goto L202
L206:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v64)+36)) = v1106
	v1108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61)+40)))
	*(*uint8)(unsafe.Add(mBase, uint32(v64)+40)) = uint8(v1108)
	v1110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61)+41)))
	*(*uint8)(unsafe.Add(mBase, uint32(v64)+41)) = uint8(v1110)
	v1112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61)+42)))
	*(*uint8)(unsafe.Add(mBase, uint32(v64)+42)) = uint8(v1112)
	v1114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61)+43)))
	*(*uint8)(unsafe.Add(mBase, uint32(v64)+43)) = uint8(v1114)
	v1116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61)+44)))
	*(*uint8)(unsafe.Add(mBase, uint32(v64)+44)) = uint8(v1116)
	v1118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61)+45)))
	*(*uint8)(unsafe.Add(mBase, uint32(v64)+45)) = uint8(v1118)
	v1120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61)+46)))
	*(*uint8)(unsafe.Add(mBase, uint32(v64)+46)) = uint8(v1120)
	v1122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61)+47)))
	*(*uint8)(unsafe.Add(mBase, uint32(v64)+47)) = uint8(v1122)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+20)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v29)+16)) = v38
	v1127 = v52 + int32(-92)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+24)) = v1127
	*(*int32)(unsafe.Add(mBase, uint32(v29)+28)) = v1086
	*(*int32)(unsafe.Add(mBase, uint32(v29)+32)) = v1035
	*(*int32)(unsafe.Add(mBase, uint32(v29)+36)) = v1009
	*(*int32)(unsafe.Add(mBase, uint32(v29)+40)) = v983
	*(*int32)(unsafe.Add(mBase, uint32(v29)+44)) = v957
	*(*int32)(unsafe.Add(mBase, uint32(v29)+48)) = v906
	*(*int32)(unsafe.Add(mBase, uint32(v29)+52)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v29)+56)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v29)+60)) = v61
	v1163 = int32(0)
	goto L210
L207:
	;
	v1106 = int32(0)
	goto L206
L208:
	;
	goto L209
L209:
	;
	v1105 = F___memcpy(m, v1101, v1081, v1100)
	mBase = m.M
	v1106 = v1105
	goto L206
L210:
	;
	v1166 = v1163 * int32(12)
	v1169 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1166)+uint32(_consts[1025]))))
	if v1169 == int32(1) {
		goto L212
	} else {
		goto L213
	}
L211:
	;
	if v869 == int32(0) {
		goto L217
	} else {
		goto L218
	}
L212:
	;
	v1174 = *(*int32)(unsafe.Add(mBase, uint32(v1166)+uint32(_consts[1024])))
	v1176 = *(*int32)(unsafe.Add(mBase, uint32(v61+v1174)))
	F_emscripten_builtin_free(m, v1176)
	mBase = m.M
	goto L214
L213:
	;
	goto L214
L214:
	;
	v1179 = v1163 + int32(1)
	if v1179 != int32(18) {
		v1163 = v1179
		goto L210
	} else {
		goto L215
	}
L215:
	;
	goto L211
L216:
	;
	v1266 = *(*int32)(unsafe.Add(mBase, _consts[52]))
	v1268 = *(*int32)(unsafe.Add(mBase, _consts[53]))
	goto L232
L217:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+20)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v29)+16)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v29)+24)) = v1127
	*(*int32)(unsafe.Add(mBase, uint32(v29)+28)) = v1086
	*(*int32)(unsafe.Add(mBase, uint32(v29)+32)) = v1035
	*(*int32)(unsafe.Add(mBase, uint32(v29)+36)) = v1009
	*(*int32)(unsafe.Add(mBase, uint32(v29)+40)) = v983
	*(*int32)(unsafe.Add(mBase, uint32(v29)+44)) = v957
	*(*int32)(unsafe.Add(mBase, uint32(v29)+48)) = v906
	*(*int32)(unsafe.Add(mBase, uint32(v29)+52)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v29)+56)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v29)+60)) = v61
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1215 = m.ExcPending
	if v1215 != 0 {
		v1588 = v67
		goto L6
	} else {
		goto L228
	}
L218:
	;
	if v894 == int32(0) {
		goto L217
	} else {
		goto L219
	}
L219:
	;
	if v921 == int32(0) {
		goto L217
	} else {
		goto L220
	}
L220:
	;
	if v946 == int32(0) {
		goto L217
	} else {
		goto L221
	}
L221:
	;
	if v973 == int32(0) {
		goto L217
	} else {
		goto L222
	}
L222:
	;
	if v1000 == int32(0) {
		goto L217
	} else {
		goto L223
	}
L223:
	;
	if v1027 == int32(0) {
		goto L217
	} else {
		goto L224
	}
L224:
	;
	if v1054 == int32(0) {
		goto L217
	} else {
		goto L225
	}
L225:
	;
	if v1079 == int32(0) {
		goto L217
	} else {
		goto L226
	}
L226:
	;
	if v1106 != 0 {
		goto L216
	} else {
		goto L227
	}
L227:
	;
	goto L217
L228:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+20)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v29)+16)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v29)+24)) = v1127
	*(*int32)(unsafe.Add(mBase, uint32(v29)+28)) = v1086
	*(*int32)(unsafe.Add(mBase, uint32(v29)+32)) = v1035
	*(*int32)(unsafe.Add(mBase, uint32(v29)+36)) = v1009
	*(*int32)(unsafe.Add(mBase, uint32(v29)+40)) = v983
	*(*int32)(unsafe.Add(mBase, uint32(v29)+44)) = v957
	*(*int32)(unsafe.Add(mBase, uint32(v29)+48)) = v906
	*(*int32)(unsafe.Add(mBase, uint32(v29)+52)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v29)+56)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v29)+60)) = v61
	F_errcode(m, int32(8389))
	mBase = m.M
	v1230 = m.ExcPending
	if v1230 != 0 {
		v1588 = v67
		goto L6
	} else {
		goto L229
	}
L229:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+20)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v29)+16)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v29)+24)) = v1127
	*(*int32)(unsafe.Add(mBase, uint32(v29)+28)) = v1086
	*(*int32)(unsafe.Add(mBase, uint32(v29)+32)) = v1035
	*(*int32)(unsafe.Add(mBase, uint32(v29)+36)) = v1009
	*(*int32)(unsafe.Add(mBase, uint32(v29)+40)) = v983
	*(*int32)(unsafe.Add(mBase, uint32(v29)+44)) = v957
	*(*int32)(unsafe.Add(mBase, uint32(v29)+48)) = v906
	*(*int32)(unsafe.Add(mBase, uint32(v29)+52)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v29)+56)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v29)+60)) = v61
	F_errmsg(m, int32(14020), int32(0))
	mBase = m.M
	v1246 = m.ExcPending
	if v1246 != 0 {
		v1588 = v67
		goto L6
	} else {
		goto L230
	}
L230:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+20)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v29)+16)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v29)+24)) = v1127
	*(*int32)(unsafe.Add(mBase, uint32(v29)+28)) = v1086
	*(*int32)(unsafe.Add(mBase, uint32(v29)+32)) = v1035
	*(*int32)(unsafe.Add(mBase, uint32(v29)+36)) = v1009
	*(*int32)(unsafe.Add(mBase, uint32(v29)+40)) = v983
	*(*int32)(unsafe.Add(mBase, uint32(v29)+44)) = v957
	*(*int32)(unsafe.Add(mBase, uint32(v29)+48)) = v906
	*(*int32)(unsafe.Add(mBase, uint32(v29)+52)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v29)+56)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v29)+60)) = v61
	F_errfinish(m, int32(522638), int32(591), int32(34006))
	mBase = m.M
	v1263 = m.ExcPending
	if v1263 != 0 {
		v1588 = v67
		goto L6
	} else {
		goto L231
	}
L231:
	;
	goto L3
L232:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v67)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v67))) = v27 + int32(-52)
	goto L235
L233:
	;
	v1275 = v64
	v1276 = v61
	v1277 = int32(0)
	v1278 = v67
	v1279 = v1268
	v1280 = v1266
	v1281 = v1127
	v1282 = v1086
	v1283 = v1035
	v1284 = v1009
	v1285 = v983
	v1286 = v957
	v1287 = v906
	v1294 = v67
	goto L11
L235:
	;
	goto L233
L236:
	;
	*(*int32)(unsafe.Add(mBase, _consts[53])) = v1278
	*(*int32)(unsafe.Add(mBase, uint32(v29)+16)) = v1280
	*(*int32)(unsafe.Add(mBase, uint32(v29)+20)) = v1279
	*(*int32)(unsafe.Add(mBase, uint32(v29)+24)) = v1281
	*(*int32)(unsafe.Add(mBase, uint32(v29)+28)) = v1282
	*(*int32)(unsafe.Add(mBase, uint32(v29)+32)) = v1283
	*(*int32)(unsafe.Add(mBase, uint32(v29)+36)) = v1284
	*(*int32)(unsafe.Add(mBase, uint32(v29)+40)) = v1285
	*(*int32)(unsafe.Add(mBase, uint32(v29)+44)) = v1286
	*(*int32)(unsafe.Add(mBase, uint32(v29)+48)) = v1287
	*(*int32)(unsafe.Add(mBase, uint32(v29)+52)) = v1278
	*(*int32)(unsafe.Add(mBase, uint32(v29)+56)) = v1275
	*(*int32)(unsafe.Add(mBase, uint32(v29)+60)) = v1276
	v1315 = *(*int32)(unsafe.Add(mBase, _consts[1018]))
	v1317 = F_pg_get_encoding_from_locale(m, v1315, int32(1))
	mBase = m.M
	v1318 = m.ExcPending
	if v1318 != 0 {
		v1588 = v1294
		goto L6
	} else {
		goto L237
	}
L237:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+20)) = v1279
	*(*int32)(unsafe.Add(mBase, uint32(v29)+16)) = v1280
	*(*int32)(unsafe.Add(mBase, uint32(v29)+24)) = v1281
	*(*int32)(unsafe.Add(mBase, uint32(v29)+28)) = v1282
	*(*int32)(unsafe.Add(mBase, uint32(v29)+32)) = v1283
	*(*int32)(unsafe.Add(mBase, uint32(v29)+36)) = v1284
	*(*int32)(unsafe.Add(mBase, uint32(v29)+40)) = v1285
	*(*int32)(unsafe.Add(mBase, uint32(v29)+44)) = v1286
	*(*int32)(unsafe.Add(mBase, uint32(v29)+48)) = v1287
	*(*int32)(unsafe.Add(mBase, uint32(v29)+52)) = v1278
	*(*int32)(unsafe.Add(mBase, uint32(v29)+56)) = v1275
	*(*int32)(unsafe.Add(mBase, uint32(v29)+60)) = v1276
	v1331 = int32(0)
	if v1331 < v1317 {
		goto L238
	} else {
		goto L239
	}
L238:
	;
	v1334 = v1317
	goto L240
L239:
	;
	v1334 = v1331
	goto L240
L240:
	;
	F_db_encoding_convert(m, v1334, v1275)
	mBase = m.M
	v1336 = m.ExcPending
	if v1336 != 0 {
		v1588 = v1294
		goto L6
	} else {
		goto L241
	}
L241:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+20)) = v1279
	*(*int32)(unsafe.Add(mBase, uint32(v29)+16)) = v1280
	*(*int32)(unsafe.Add(mBase, uint32(v29)+24)) = v1281
	*(*int32)(unsafe.Add(mBase, uint32(v29)+28)) = v1282
	*(*int32)(unsafe.Add(mBase, uint32(v29)+32)) = v1283
	*(*int32)(unsafe.Add(mBase, uint32(v29)+36)) = v1284
	*(*int32)(unsafe.Add(mBase, uint32(v29)+40)) = v1285
	*(*int32)(unsafe.Add(mBase, uint32(v29)+44)) = v1286
	*(*int32)(unsafe.Add(mBase, uint32(v29)+48)) = v1287
	*(*int32)(unsafe.Add(mBase, uint32(v29)+52)) = v1278
	*(*int32)(unsafe.Add(mBase, uint32(v29)+56)) = v1275
	*(*int32)(unsafe.Add(mBase, uint32(v29)+60)) = v1276
	F_db_encoding_convert(m, v1334, v1287)
	mBase = m.M
	v1350 = m.ExcPending
	if v1350 != 0 {
		v1588 = v1294
		goto L6
	} else {
		goto L242
	}
L242:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+16)) = v1280
	*(*int32)(unsafe.Add(mBase, uint32(v29)+20)) = v1279
	*(*int32)(unsafe.Add(mBase, uint32(v29)+24)) = v1281
	*(*int32)(unsafe.Add(mBase, uint32(v29)+28)) = v1282
	*(*int32)(unsafe.Add(mBase, uint32(v29)+32)) = v1283
	*(*int32)(unsafe.Add(mBase, uint32(v29)+36)) = v1284
	*(*int32)(unsafe.Add(mBase, uint32(v29)+40)) = v1285
	*(*int32)(unsafe.Add(mBase, uint32(v29)+44)) = v1286
	*(*int32)(unsafe.Add(mBase, uint32(v29)+48)) = v1287
	*(*int32)(unsafe.Add(mBase, uint32(v29)+52)) = v1278
	*(*int32)(unsafe.Add(mBase, uint32(v29)+56)) = v1275
	*(*int32)(unsafe.Add(mBase, uint32(v29)+60)) = v1276
	v1364 = *(*int32)(unsafe.Add(mBase, _consts[1017]))
	v1366 = F_pg_get_encoding_from_locale(m, v1364, int32(1))
	mBase = m.M
	v1367 = m.ExcPending
	if v1367 != 0 {
		v1588 = v1294
		goto L6
	} else {
		goto L243
	}
L243:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+20)) = v1279
	*(*int32)(unsafe.Add(mBase, uint32(v29)+16)) = v1280
	*(*int32)(unsafe.Add(mBase, uint32(v29)+24)) = v1281
	*(*int32)(unsafe.Add(mBase, uint32(v29)+28)) = v1282
	*(*int32)(unsafe.Add(mBase, uint32(v29)+32)) = v1283
	*(*int32)(unsafe.Add(mBase, uint32(v29)+36)) = v1284
	*(*int32)(unsafe.Add(mBase, uint32(v29)+40)) = v1285
	*(*int32)(unsafe.Add(mBase, uint32(v29)+44)) = v1286
	*(*int32)(unsafe.Add(mBase, uint32(v29)+48)) = v1287
	*(*int32)(unsafe.Add(mBase, uint32(v29)+52)) = v1278
	*(*int32)(unsafe.Add(mBase, uint32(v29)+56)) = v1275
	*(*int32)(unsafe.Add(mBase, uint32(v29)+60)) = v1276
	v1380 = int32(0)
	if v1380 < v1366 {
		goto L244
	} else {
		goto L245
	}
L244:
	;
	v1383 = v1366
	goto L246
L245:
	;
	v1383 = v1380
	goto L246
L246:
	;
	F_db_encoding_convert(m, v1383, v1286)
	mBase = m.M
	v1385 = m.ExcPending
	if v1385 != 0 {
		v1588 = v1294
		goto L6
	} else {
		goto L247
	}
L247:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+20)) = v1279
	*(*int32)(unsafe.Add(mBase, uint32(v29)+16)) = v1280
	*(*int32)(unsafe.Add(mBase, uint32(v29)+24)) = v1281
	*(*int32)(unsafe.Add(mBase, uint32(v29)+28)) = v1282
	*(*int32)(unsafe.Add(mBase, uint32(v29)+32)) = v1283
	*(*int32)(unsafe.Add(mBase, uint32(v29)+36)) = v1284
	*(*int32)(unsafe.Add(mBase, uint32(v29)+40)) = v1285
	*(*int32)(unsafe.Add(mBase, uint32(v29)+44)) = v1286
	*(*int32)(unsafe.Add(mBase, uint32(v29)+48)) = v1287
	*(*int32)(unsafe.Add(mBase, uint32(v29)+52)) = v1278
	*(*int32)(unsafe.Add(mBase, uint32(v29)+56)) = v1275
	*(*int32)(unsafe.Add(mBase, uint32(v29)+60)) = v1276
	F_db_encoding_convert(m, v1383, v1285)
	mBase = m.M
	v1399 = m.ExcPending
	if v1399 != 0 {
		v1588 = v1294
		goto L6
	} else {
		goto L248
	}
L248:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+20)) = v1279
	*(*int32)(unsafe.Add(mBase, uint32(v29)+16)) = v1280
	*(*int32)(unsafe.Add(mBase, uint32(v29)+24)) = v1281
	*(*int32)(unsafe.Add(mBase, uint32(v29)+28)) = v1282
	*(*int32)(unsafe.Add(mBase, uint32(v29)+32)) = v1283
	*(*int32)(unsafe.Add(mBase, uint32(v29)+36)) = v1284
	*(*int32)(unsafe.Add(mBase, uint32(v29)+40)) = v1285
	*(*int32)(unsafe.Add(mBase, uint32(v29)+44)) = v1286
	*(*int32)(unsafe.Add(mBase, uint32(v29)+48)) = v1287
	*(*int32)(unsafe.Add(mBase, uint32(v29)+52)) = v1278
	*(*int32)(unsafe.Add(mBase, uint32(v29)+56)) = v1275
	*(*int32)(unsafe.Add(mBase, uint32(v29)+60)) = v1276
	F_db_encoding_convert(m, v1383, v1284)
	mBase = m.M
	v1413 = m.ExcPending
	if v1413 != 0 {
		v1588 = v1294
		goto L6
	} else {
		goto L249
	}
L249:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+20)) = v1279
	*(*int32)(unsafe.Add(mBase, uint32(v29)+16)) = v1280
	*(*int32)(unsafe.Add(mBase, uint32(v29)+24)) = v1281
	*(*int32)(unsafe.Add(mBase, uint32(v29)+28)) = v1282
	*(*int32)(unsafe.Add(mBase, uint32(v29)+32)) = v1283
	*(*int32)(unsafe.Add(mBase, uint32(v29)+36)) = v1284
	*(*int32)(unsafe.Add(mBase, uint32(v29)+40)) = v1285
	*(*int32)(unsafe.Add(mBase, uint32(v29)+44)) = v1286
	*(*int32)(unsafe.Add(mBase, uint32(v29)+48)) = v1287
	*(*int32)(unsafe.Add(mBase, uint32(v29)+52)) = v1278
	*(*int32)(unsafe.Add(mBase, uint32(v29)+56)) = v1275
	*(*int32)(unsafe.Add(mBase, uint32(v29)+60)) = v1276
	F_db_encoding_convert(m, v1383, v1283)
	mBase = m.M
	v1427 = m.ExcPending
	if v1427 != 0 {
		v1588 = v1294
		goto L6
	} else {
		goto L250
	}
L250:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+20)) = v1279
	*(*int32)(unsafe.Add(mBase, uint32(v29)+16)) = v1280
	*(*int32)(unsafe.Add(mBase, uint32(v29)+24)) = v1281
	*(*int32)(unsafe.Add(mBase, uint32(v29)+28)) = v1282
	*(*int32)(unsafe.Add(mBase, uint32(v29)+32)) = v1283
	*(*int32)(unsafe.Add(mBase, uint32(v29)+36)) = v1284
	*(*int32)(unsafe.Add(mBase, uint32(v29)+40)) = v1285
	*(*int32)(unsafe.Add(mBase, uint32(v29)+44)) = v1286
	*(*int32)(unsafe.Add(mBase, uint32(v29)+48)) = v1287
	*(*int32)(unsafe.Add(mBase, uint32(v29)+52)) = v1278
	*(*int32)(unsafe.Add(mBase, uint32(v29)+56)) = v1275
	*(*int32)(unsafe.Add(mBase, uint32(v29)+60)) = v1276
	F_db_encoding_convert(m, v1383, v1282)
	mBase = m.M
	v1441 = m.ExcPending
	if v1441 != 0 {
		v1588 = v1294
		goto L6
	} else {
		goto L251
	}
L251:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+20)) = v1279
	*(*int32)(unsafe.Add(mBase, uint32(v29)+16)) = v1280
	*(*int32)(unsafe.Add(mBase, uint32(v29)+24)) = v1281
	*(*int32)(unsafe.Add(mBase, uint32(v29)+28)) = v1282
	*(*int32)(unsafe.Add(mBase, uint32(v29)+32)) = v1283
	*(*int32)(unsafe.Add(mBase, uint32(v29)+36)) = v1284
	*(*int32)(unsafe.Add(mBase, uint32(v29)+40)) = v1285
	*(*int32)(unsafe.Add(mBase, uint32(v29)+44)) = v1286
	*(*int32)(unsafe.Add(mBase, uint32(v29)+48)) = v1287
	*(*int32)(unsafe.Add(mBase, uint32(v29)+52)) = v1278
	*(*int32)(unsafe.Add(mBase, uint32(v29)+56)) = v1275
	*(*int32)(unsafe.Add(mBase, uint32(v29)+60)) = v1276
	F_db_encoding_convert(m, v1383, v1281)
	mBase = m.M
	v1455 = m.ExcPending
	if v1455 != 0 {
		v1588 = v1294
		goto L6
	} else {
		goto L252
	}
L252:
	;
	*(*int32)(unsafe.Add(mBase, _consts[52])) = v1280
	*(*int32)(unsafe.Add(mBase, _consts[53])) = v1279
	v1461 = *(*int64)(unsafe.Add(mBase, uint32(v1275)))
	*(*int64)(unsafe.Add(mBase, _consts[1007])) = v1461
	v1464 = *(*int64)(unsafe.Add(mBase, uint32(v1275)+8))
	*(*int64)(unsafe.Add(mBase, _consts[1009])) = v1464
	v1467 = *(*int64)(unsafe.Add(mBase, uint32(v1275)+16))
	*(*int64)(unsafe.Add(mBase, _consts[1011])) = v1467
	v1470 = *(*int64)(unsafe.Add(mBase, uint32(v1275)+24))
	*(*int64)(unsafe.Add(mBase, _consts[1013])) = v1470
	v1473 = *(*int64)(unsafe.Add(mBase, uint32(v1275)+32))
	*(*int64)(unsafe.Add(mBase, _consts[1015])) = v1473
	v1476 = *(*int64)(unsafe.Add(mBase, uint32(v1275)+40))
	*(*int64)(unsafe.Add(mBase, _consts[1027])) = v1476
	v1479 = *(*int64)(unsafe.Add(mBase, uint32(v1275)+48))
	*(*int64)(unsafe.Add(mBase, _consts[1028])) = v1479
	v1482 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[1006])) = uint8(v1482)
	*(*uint8)(unsafe.Add(mBase, _consts[1005])) = uint8(v1482)
	goto L8
L253:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+20)) = v1279
	*(*int32)(unsafe.Add(mBase, uint32(v29)+16)) = v1280
	*(*int32)(unsafe.Add(mBase, uint32(v29)+24)) = v1281
	*(*int32)(unsafe.Add(mBase, uint32(v29)+28)) = v1282
	*(*int32)(unsafe.Add(mBase, uint32(v29)+32)) = v1283
	*(*int32)(unsafe.Add(mBase, uint32(v29)+36)) = v1284
	*(*int32)(unsafe.Add(mBase, uint32(v29)+40)) = v1285
	*(*int32)(unsafe.Add(mBase, uint32(v29)+44)) = v1286
	*(*int32)(unsafe.Add(mBase, uint32(v29)+48)) = v1287
	*(*int32)(unsafe.Add(mBase, uint32(v29)+52)) = v1278
	*(*int32)(unsafe.Add(mBase, uint32(v29)+56)) = v1275
	*(*int32)(unsafe.Add(mBase, uint32(v29)+60)) = v1276
	F_pg_re_throw(m)
	mBase = m.M
	v1567 = m.ExcPending
	if v1567 != 0 {
		v1588 = v1294
		goto L6
	} else {
		goto L254
	}
L254:
	;
	goto L5
L255:
	;
	v1599 = int32(v1595)
	m.G0 = v1588
	v1601 = *(*int32)(unsafe.Add(mBase, uint32(v1599)+4))
	v1602 = *(*int32)(unsafe.Add(mBase, uint32(v1599)))
	v1606 = *(*int32)(unsafe.Add(mBase, uint32(v1602)))
	if v27+int32(-52) == v1606 {
		goto L258
	} else {
		goto L259
	}
L256:
	;
	m.ExcPending = 1
	goto L264
L257:
	;
	if v1609 != 0 {
		goto L261
	} else {
		goto L262
	}
L258:
	;
	v1608 = *(*int32)(unsafe.Add(mBase, uint32(v1602)+4))
	v1609 = v1608
	goto L260
L259:
	;
	v1609 = int32(0)
	goto L260
L260:
	;
	goto L257
L261:
	;
	v1610 = *(*int32)(unsafe.Add(mBase, uint32(v29)+60))
	v1611 = *(*int32)(unsafe.Add(mBase, uint32(v29)+56))
	v1612 = *(*int32)(unsafe.Add(mBase, uint32(v29)+52))
	v1613 = *(*int32)(unsafe.Add(mBase, uint32(v29)+48))
	v1614 = *(*int32)(unsafe.Add(mBase, uint32(v29)+44))
	v1615 = *(*int32)(unsafe.Add(mBase, uint32(v29)+40))
	v1616 = *(*int32)(unsafe.Add(mBase, uint32(v29)+36))
	v1617 = *(*int32)(unsafe.Add(mBase, uint32(v29)+32))
	v1618 = *(*int32)(unsafe.Add(mBase, uint32(v29)+28))
	v1619 = *(*int32)(unsafe.Add(mBase, uint32(v29)+24))
	v1620 = *(*int32)(unsafe.Add(mBase, uint32(v29)+20))
	v1621 = *(*int32)(unsafe.Add(mBase, uint32(v29)+16))
	v33 = v1611
	v34 = v1610
	v35 = v1601
	v36 = v1612
	v37 = v1620
	v38 = v1621
	v39 = v1619
	v40 = v1618
	v41 = v1617
	v42 = v1616
	v43 = v1615
	v44 = v1614
	v45 = v1613
	v46 = v1609
	v52 = v1588
	goto L1
L262:
	;
	goto L263
L263:
	;
	F___wasm_longjmp(m, v1602, v1601)
	mBase = m.M
	v1625 = m.ExcPending
	if v1625 != 0 {
		goto L264
	} else {
		goto L265
	}
L264:
	;
	return int32(0)
L265:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_ParseCommitRecord(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int64
	_ = v11
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v96 int32
	_ = v96
	var v103 int32
	_ = v103
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v223 int64
	_ = v223
	var v224 int64
	_ = v224
	v10 = F__emscripten_memset_bulkmem(m, l2, base.I32_extend8_s(int32(0)), int32(288))
	mBase = m.M
	goto L1
L1:
	;
	v11 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	*(*int64)(unsafe.Add(mBase, uint32(v10))) = v11
	if int32(0) <= base.I32_extend8_s(l0) {
		goto L2
	} else {
		goto L3
	}
L2:
	;
	return
L3:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v16
	if v16&int32(1) != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = v20
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v22
	v28 = l1 + int32(20)
	goto L6
L5:
	;
	v28 = l1 + int32(12)
	goto L6
L6:
	;
	if v16&int32(2) != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	v33 = v28 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v31
	v39 = v33 + v31<<(uint(int32(2))%32)
	goto L9
L8:
	;
	v39 = v28
	goto L9
L9:
	;
	if v16&int32(4) != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
	v45 = v39 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v10)+28)) = v43
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
	v52 = v45 + v48*int32(12)
	goto L12
L11:
	;
	v52 = v39
	goto L12
L12:
	;
	if v16&int32(256) != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
	v58 = int32(4)
	v59 = v52 + v58
	*(*int32)(unsafe.Add(mBase, uint32(v10)+40)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v10)+36)) = v57
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
	v66 = v59 + v62<<(uint(v58)%32)
	goto L15
L14:
	;
	v66 = v52
	goto L15
L15:
	;
	if v16&int32(8) != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v66)))
	v72 = int32(4)
	v73 = v66 + v72
	*(*int32)(unsafe.Add(mBase, uint32(v10)+48)) = v73
	*(*int32)(unsafe.Add(mBase, uint32(v10)+44)) = v71
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v66)))
	v80 = v73 + v76<<(uint(v72)%32)
	goto L18
L17:
	;
	v80 = v66
	goto L18
L18:
	;
	if v16&int32(16) == int32(0) {
		v217 = v16
		v218 = v80
		goto L19
	} else {
		goto L20
	}
L19:
	;
	if v217&int32(32) == int32(0) {
		goto L2
	} else {
		goto L54
	}
L20:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+52)) = v87
	v90 = v80 + int32(4)
	if v16&int32(128) == int32(0) {
		v217 = v16
		v218 = v90
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v96 = v10 + int32(56)
	goto L25
L22:
	;
	v212 = F_strlen(m, v90)
	mBase = m.M
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
	v217 = v216
	v218 = v212 + v90 + int32(1)
	goto L19
L23:
	;
	v209 = F_strlen(m, v198)
	mBase = m.M
	goto L22
L25:
	;
	goto L26
L26:
	;
	v103 = int32(199)
	if (v96^v90)&int32(3) != 0 {
		goto L30
	} else {
		goto L31
	}
L27:
	;
	v202 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v199))) = uint8(v202)
	goto L23
L28:
	;
	v183 = v178
	v184 = v179
	v185 = v180
	goto L50
L29:
	;
	if v173 == int32(0) {
		v198 = v171
		v199 = v172
		goto L27
	} else {
		goto L49
	}
L30:
	;
	v171 = v90
	v172 = v96
	v173 = v103
	goto L29
L31:
	;
	goto L32
L32:
	;
	if v90&int32(3) == int32(0) {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	if v140 == int32(0) {
		v198 = v137
		v199 = v138
		goto L27
	} else {
		goto L42
	}
L34:
	;
	v137 = v90
	v138 = v96
	v139 = v103
	v140 = int32(1)
	goto L33
L35:
	;
	goto L36
L36:
	;
	v116 = v90
	v117 = v96
	v118 = v103
	goto L37
L37:
	;
	v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116))))
	*(*uint8)(unsafe.Add(mBase, uint32(v117))) = uint8(v120)
	if v120 == int32(0) {
		v178 = v116
		v179 = v117
		v180 = v118
		goto L28
	} else {
		goto L39
	}
L38:
	;
	v137 = v131
	v138 = v125
	v139 = v127
	v140 = v129
	goto L33
L39:
	;
	v124 = int32(1)
	v125 = v117 + v124
	v127 = v118 - v124
	v128 = int32(0)
	v129 = base.B2i32(v127 != v128)
	v131 = v116 + v124
	if v131&int32(3) == v128 {
		v137 = v131
		v138 = v125
		v139 = v127
		v140 = v129
		goto L33
	} else {
		goto L40
	}
L40:
	;
	if v127 != 0 {
		v116 = v131
		v117 = v125
		v118 = v127
		goto L37
	} else {
		goto L41
	}
L41:
	;
	goto L38
L42:
	;
	v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v137))))
	if v143 == int32(0) {
		v171 = v137
		v172 = v138
		v173 = v139
		goto L29
	} else {
		goto L43
	}
L43:
	;
	if base.Ui32(v139) < base.Ui32(int32(4)) {
		v171 = v137
		v172 = v138
		v173 = v139
		goto L29
	} else {
		goto L44
	}
L44:
	;
	v149 = v137
	v150 = v138
	v151 = v139
	goto L45
L45:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v149)))
	v157 = int32(-2139062144)
	if (int32(16843008)-v154|v154)&v157 != v157 {
		v178 = v149
		v179 = v150
		v180 = v151
		goto L28
	} else {
		goto L47
	}
L46:
	;
	v171 = v165
	v172 = v163
	v173 = v167
	goto L29
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v150))) = v154
	v162 = int32(4)
	v163 = v150 + v162
	v165 = v149 + v162
	v167 = v151 - v162
	if base.Ui32(int32(3)) < base.Ui32(v167) {
		v149 = v165
		v150 = v163
		v151 = v167
		goto L45
	} else {
		goto L48
	}
L48:
	;
	goto L46
L49:
	;
	v178 = v171
	v179 = v172
	v180 = v173
	goto L28
L50:
	;
	v187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v183))))
	*(*uint8)(unsafe.Add(mBase, uint32(v184))) = uint8(v187)
	if v187 == int32(0) {
		v198 = v183
		v199 = v184
		goto L27
	} else {
		goto L52
	}
L51:
	;
	v198 = v194
	v199 = v192
	goto L27
L52:
	;
	v191 = int32(1)
	v192 = v184 + v191
	v194 = v183 + v191
	v196 = v185 - v191
	if v196 != 0 {
		v183 = v194
		v184 = v192
		v185 = v196
		goto L50
	} else {
		goto L53
	}
L53:
	;
	goto L51
L54:
	;
	v223 = *(*int64)(unsafe.Add(mBase, uint32(v218)))
	v224 = *(*int64)(unsafe.Add(mBase, uint32(v218)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v10)+280)) = v224
	*(*int64)(unsafe.Add(mBase, uint32(v10)+272)) = v223
	goto L2
}
func F_ParseTzFile(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v102 int32
	_ = v102
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v115 int32
	_ = v115
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v148 int32
	_ = v148
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v218 int32
	_ = v218
	var v223 int32
	_ = v223
	var v232 int32
	_ = v232
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v255 int32
	_ = v255
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v272 int32
	_ = v272
	var v276 int32
	_ = v276
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v294 int32
	_ = v294
	var v309 int32
	_ = v309
	var v318 int32
	_ = v318
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v348 int32
	_ = v348
	var v357 int32
	_ = v357
	var v360 int32
	_ = v360
	var v362 int32
	_ = v362
	var v379 int32
	_ = v379
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v387 int32
	_ = v387
	var v389 int32
	_ = v389
	var v391 int32
	_ = v391
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v398 int32
	_ = v398
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v408 int32
	_ = v408
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v417 int32
	_ = v417
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v453 int32
	_ = v453
	var v462 int32
	_ = v462
	var v465 int32
	_ = v465
	var v467 int32
	_ = v467
	var v484 int32
	_ = v484
	var v487 int32
	_ = v487
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v499 int32
	_ = v499
	var v501 int32
	_ = v501
	var v503 int32
	_ = v503
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v510 int32
	_ = v510
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v520 int32
	_ = v520
	var v526 int32
	_ = v526
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v534 int32
	_ = v534
	var v536 int32
	_ = v536
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v545 int32
	_ = v545
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v555 int32
	_ = v555
	var v561 int32
	_ = v561
	var v565 int32
	_ = v565
	var v580 int64
	_ = v580
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v589 int32
	_ = v589
	var v591 int32
	_ = v591
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v600 int32
	_ = v600
	var v604 int32
	_ = v604
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v610 int32
	_ = v610
	var v616 int32
	_ = v616
	var v620 int32
	_ = v620
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v628 int32
	_ = v628
	var v629 int32
	_ = v629
	var v639 int32
	_ = v639
	var v648 int32
	_ = v648
	var v651 int32
	_ = v651
	var v653 int32
	_ = v653
	var v662 int32
	_ = v662
	var v665 int32
	_ = v665
	var v666 int32
	_ = v666
	var v671 int32
	_ = v671
	var v672 int32
	_ = v672
	var v673 int32
	_ = v673
	var v676 int32
	_ = v676
	var v678 int32
	_ = v678
	var v682 int32
	_ = v682
	var v683 int32
	_ = v683
	var v684 int32
	_ = v684
	var v687 int32
	_ = v687
	var v691 int32
	_ = v691
	var v692 int32
	_ = v692
	var v693 int32
	_ = v693
	var v697 int32
	_ = v697
	var v703 int32
	_ = v703
	var v707 int32
	_ = v707
	var v709 int32
	_ = v709
	var v710 int32
	_ = v710
	var v711 int32
	_ = v711
	var v712 int32
	_ = v712
	var v716 int32
	_ = v716
	var v724 int32
	_ = v724
	var v732 int32
	_ = v732
	var v733 int32
	_ = v733
	var v738 int32
	_ = v738
	var v739 int32
	_ = v739
	var v740 int32
	_ = v740
	var v742 int32
	_ = v742
	var v746 int32
	_ = v746
	var v757 int32
	_ = v757
	var v758 int32
	_ = v758
	var v763 int32
	_ = v763
	var v770 int32
	_ = v770
	var v771 int32
	_ = v771
	var v785 int32
	_ = v785
	var v786 int32
	_ = v786
	var v795 int32
	_ = v795
	var v798 int32
	_ = v798
	var v799 int32
	_ = v799
	var v821 int32
	_ = v821
	var v822 int32
	_ = v822
	var v824 int32
	_ = v824
	var v828 int32
	_ = v828
	var v837 int32
	_ = v837
	var v838 int32
	_ = v838
	var v839 int32
	_ = v839
	var v847 int32
	_ = v847
	var v848 int32
	_ = v848
	var v864 int32
	_ = v864
	var v867 int32
	_ = v867
	var v868 int32
	_ = v868
	var v871 int32
	_ = v871
	var v872 int32
	_ = v872
	var v876 int32
	_ = v876
	var v877 int32
	_ = v877
	var v880 int32
	_ = v880
	var v881 int32
	_ = v881
	var v884 int32
	_ = v884
	var v891 int32
	_ = v891
	var v892 int32
	_ = v892
	var v893 int32
	_ = v893
	var v902 int32
	_ = v902
	var v903 int32
	_ = v903
	var v905 int32
	_ = v905
	var v908 int32
	_ = v908
	var v909 int32
	_ = v909
	var v911 int32
	_ = v911
	var v918 int32
	_ = v918
	var v919 int32
	_ = v919
	var v923 int32
	_ = v923
	var v924 int32
	_ = v924
	var v927 int32
	_ = v927
	var v928 int32
	_ = v928
	var v931 int32
	_ = v931
	var v938 int32
	_ = v938
	var v939 int32
	_ = v939
	var v943 int32
	_ = v943
	var v948 int32
	_ = v948
	var v956 int32
	_ = v956
	var v957 int32
	_ = v957
	var v960 int32
	_ = v960
	var v963 int64
	_ = v963
	var v973 int32
	_ = v973
	var v974 int32
	_ = v974
	var v1004 int32
	_ = v1004
	var v1019 int32
	_ = v1019
	var v1024 int32
	_ = v1024
	var v1027 int32
	_ = v1027
	var v1028 int32
	_ = v1028
	var v1030 int32
	_ = v1030
	var v1031 int32
	_ = v1031
	var v1033 int32
	_ = v1033
	var v1035 int32
	_ = v1035
	var v1038 int32
	_ = v1038
	var v1042 int32
	_ = v1042
	var v1049 int32
	_ = v1049
	var v1052 int32
	_ = v1052
	var v1059 int32
	_ = v1059
	var v1060 int32
	_ = v1060
	var v1061 int32
	_ = v1061
	var v1065 int32
	_ = v1065
	var v1067 int32
	_ = v1067
	var v1068 int32
	_ = v1068
	var v1070 int32
	_ = v1070
	var v1072 int32
	_ = v1072
	var v1079 int32
	_ = v1079
	var v1085 int32
	_ = v1085
	var v1086 int32
	_ = v1086
	var v1088 int32
	_ = v1088
	var v1094 int32
	_ = v1094
	var v1101 int32
	_ = v1101
	var v1105 int32
	_ = v1105
	var v1108 int32
	_ = v1108
	var v1114 int32
	_ = v1114
	var v1121 int32
	_ = v1121
	var v1125 int32
	_ = v1125
	var v1128 int32
	_ = v1128
	var v1131 int32
	_ = v1131
	var v1132 int32
	_ = v1132
	var v1133 int32
	_ = v1133
	var v1138 int32
	_ = v1138
	var v1139 int32
	_ = v1139
	var v1140 int32
	_ = v1140
	var v1142 int32
	_ = v1142
	var v1144 int32
	_ = v1144
	var v1145 int32
	_ = v1145
	var v1147 int32
	_ = v1147
	var v1149 int32
	_ = v1149
	var v1153 int32
	_ = v1153
	var v1154 int32
	_ = v1154
	var v1155 int32
	_ = v1155
	var v1160 int32
	_ = v1160
	var v1161 int32
	_ = v1161
	var v1162 int32
	_ = v1162
	var v1164 int32
	_ = v1164
	var v1166 int32
	_ = v1166
	var v1171 int32
	_ = v1171
	var v1187 int32
	_ = v1187
	var v1189 int32
	_ = v1189
	var v1199 int32
	_ = v1199
	var v1222 int32
	_ = v1222
	var v1239 int32
	_ = v1239
	var v1242 int32
	_ = v1242
	var v1243 int32
	_ = v1243
	var v1244 int32
	_ = v1244
	var v1255 int32
	_ = v1255
	var v1272 int32
	_ = v1272
	var v1273 int32
	_ = v1273
	var v1278 int32
	_ = v1278
	v6 = int32(0)
	v22 = m.G0
	v24 = v22 - int32(3328)
	m.G0 = v24
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v26 == v6 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v24 + int32(3328)
	return v1278
L2:
	;
	if base.Ui32(int32(4)) <= base.Ui32(l1) {
		goto L14
	} else {
		goto L15
	}
L3:
	;
	v35 = v26
	v36 = l0
	goto L4
L4:
	;
	if base.Ui32((v35|int32(32)-int32(97))&int32(255)) < base.Ui32(int32(26)) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v61 = int32(-1)
	if l1 == int32(0) {
		v1278 = v61
		goto L1
	} else {
		goto L10
	}
L6:
	;
	v59 = v36 + int32(1)
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59))))
	if v60 != 0 {
		v35 = v60
		v36 = v59
		goto L4
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	goto L5
L9:
	;
	goto L2
L10:
	;
	v65 = *(*int32)(unsafe.Add(mBase, _consts[43]))
	*(*int32)(unsafe.Add(mBase, _consts[508])) = v65
	goto L11
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+224)) = l0
	v73 = F_format_elog_string(m, int32(747451), v24+int32(224))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	return int32(0)
L13:
	;
	*(*int32)(unsafe.Add(mBase, _consts[527])) = v73
	v1278 = v61
	goto L1
L14:
	;
	v102 = *(*int32)(unsafe.Add(mBase, _consts[43]))
	*(*int32)(unsafe.Add(mBase, _consts[508])) = v102
	goto L17
L15:
	;
	goto L16
L16:
	;
	F_get_share_path(m, v24+int32(2288))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L12
	} else {
		goto L19
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24))) = l0
	v108 = F_format_elog_string(m, int32(748974), v24)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L12
	} else {
		goto L18
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, _consts[527])) = v108
	v1278 = int32(-1)
	goto L1
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+212)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v24)+208)) = v24 + int32(2288)
	v126 = F_pg_snprintf(m, v24+int32(1264), int32(1024), int32(187219), v24+int32(208))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L12
	} else {
		goto L20
	}
L20:
	;
	v131 = F_AllocateFile(m, v24+int32(1264), int32(242112))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L12
	} else {
		goto L23
	}
L21:
	;
	v1272 = F_FreeFile(m, v131)
	mBase = m.M
	v1273 = m.ExcPending
	if v1273 != 0 {
		goto L12
	} else {
		goto L349
	}
L22:
	;
	v218 = l4
	v223 = v6
	v232 = v6
	goto L46
L23:
	;
	if v131 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v131)+76))
	if v133 < int32(0) {
		goto L29
	} else {
		goto L30
	}
L25:
	;
	goto L26
L26:
	;
	v148 = *(*int32)(unsafe.Add(mBase, _consts[43]))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+64)) = v24 + int32(2288)
	v158 = F_pg_snprintf(m, v24+int32(1264), int32(1024), int32(132052), v24-int32(-64))
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L12
	} else {
		goto L33
	}
L27:
	;
	if int32(base.Ui32(v138)>>(uint(int32(4))%32))&int32(1) != 0 {
		v1255 = l4
		goto L21
	} else {
		goto L32
	}
L28:
	;
	goto L27
L29:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v131)))
	v138 = v136
	goto L28
L30:
	;
	goto L31
L31:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v131)))
	v138 = v137
	goto L28
L32:
	;
	goto L22
L33:
	;
	v162 = F_AllocateDir(m, v24+int32(1264))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L12
	} else {
		goto L34
	}
L34:
	;
	if v162 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v167 = *(*int32)(unsafe.Add(mBase, _consts[43]))
	*(*int32)(unsafe.Add(mBase, _consts[508])) = v167
	goto L38
L36:
	;
	goto L37
L37:
	;
	F_FreeDir(m, v162)
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L12
	} else {
		goto L42
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+32)) = v24 + int32(1264)
	v177 = F_format_elog_string(m, int32(310066), v24+int32(32))
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L12
	} else {
		goto L39
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, _consts[527])) = v177
	v181 = *(*int32)(unsafe.Add(mBase, _consts[43]))
	*(*int32)(unsafe.Add(mBase, _consts[508])) = v181
	goto L40
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+16)) = int32(4549984)
	v190 = F_format_elog_string(m, int32(644027), v24+int32(16))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L12
	} else {
		goto L41
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, _consts[712])) = v190
	v1278 = int32(-1)
	goto L1
L42:
	;
	*(*int32)(unsafe.Add(mBase, _consts[43])) = v148
	v198 = int32(-1)
	if base.B2i32(l1 == int32(0))&base.B2i32(v148 == int32(44)) != 0 {
		v1278 = v198
		goto L1
	} else {
		goto L43
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, _consts[508])) = v148
	goto L44
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+48)) = l0
	v211 = F_format_elog_string(m, int32(313009), v24+int32(48))
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L12
	} else {
		goto L45
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, _consts[527])) = v211
	v1278 = v198
	goto L1
L46:
	;
	v238 = F_fgets(m, v24+int32(240), int32(1024), v131)
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L12
	} else {
		goto L52
	}
L47:
	;
	v1255 = v1222
	goto L21
L48:
	;
	v1239 = *(*int32)(unsafe.Add(mBase, uint32(v131)+76))
	if v1239 < int32(0) {
		goto L345
	} else {
		goto L346
	}
L49:
	;
	if v1199 < int32(0) {
		v1255 = v1199
		goto L21
	} else {
		goto L342
	}
L50:
	;
	v1019 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v1019 <= v218 {
		goto L292
	} else {
		goto L293
	}
L51:
	;
	v1255 = int32(-1)
	goto L21
L52:
	;
	if v238 == int32(0) {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v131)+76))
	if v242 < int32(0) {
		goto L58
	} else {
		goto L59
	}
L54:
	;
	goto L55
L55:
	;
	v267 = v223 + int32(1)
	v269 = v24 + int32(240)
	v272 = F_strlen(m, v269)
	mBase = m.M
	if v272 == int32(1023) {
		goto L64
	} else {
		goto L65
	}
L56:
	;
	if int32(base.Ui32(v247)>>(uint(int32(5))%32))&int32(1) == int32(0) {
		v1255 = v218
		goto L21
	} else {
		goto L61
	}
L57:
	;
	goto L56
L58:
	;
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v131)))
	v247 = v245
	goto L57
L59:
	;
	goto L60
L60:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v131)))
	v247 = v246
	goto L57
L61:
	;
	v255 = *(*int32)(unsafe.Add(mBase, _consts[43]))
	*(*int32)(unsafe.Add(mBase, _consts[508])) = v255
	goto L62
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+80)) = l0
	v263 = F_format_elog_string(m, int32(313009), v24+int32(80))
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L12
	} else {
		goto L63
	}
L63:
	;
	*(*int32)(unsafe.Add(mBase, _consts[527])) = v263
	goto L51
L64:
	;
	v276 = *(*int32)(unsafe.Add(mBase, _consts[43]))
	*(*int32)(unsafe.Add(mBase, _consts[508])) = v276
	goto L67
L65:
	;
	goto L66
L66:
	;
	v294 = v269
	goto L71
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+100)) = v267
	*(*int32)(unsafe.Add(mBase, uint32(v24)+96)) = l0
	v285 = F_format_elog_string(m, int32(498551), v24+int32(96))
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L12
	} else {
		goto L68
	}
L68:
	;
	*(*int32)(unsafe.Add(mBase, _consts[527])) = v285
	goto L51
L69:
	;
	v332 = v294
	v333 = int32(566813)
	v334 = int32(8)
	goto L84
L70:
	;
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v131)+76))
	if v318 < int32(0) {
		goto L79
	} else {
		goto L80
	}
L71:
	;
	v309 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v294))))
	if base.Ui32(v309-int32(9)) < base.Ui32(int32(5)) {
		goto L74
	} else {
		goto L75
	}
L72:
	;
	if v309 != 0 {
		goto L69
	} else {
		goto L76
	}
L73:
	;
	goto L72
L74:
	;
	v294 = v294 + int32(1)
	goto L71
L75:
	;
	switch v309 - int32(32) {
	case 0:
		goto L74
	case 1, 2:
		goto L69
	case 3:
		goto L70
	default:
		goto L73
	}
L76:
	;
	goto L70
L77:
	;
	if int32(base.Ui32(v323)>>(uint(int32(4))%32))&int32(1) != 0 {
		v1255 = v218
		goto L21
	} else {
		goto L82
	}
L78:
	;
	goto L77
L79:
	;
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v131)))
	v323 = v321
	goto L78
L80:
	;
	goto L81
L81:
	;
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v131)))
	v323 = v322
	goto L78
L82:
	;
	v223 = v267
	goto L46
L83:
	;
	if v379 == int32(0) {
		goto L99
	} else {
		goto L100
	}
L84:
	;
	if v334 != 0 {
		goto L86
	} else {
		goto L87
	}
L85:
	;
	v379 = int32(0)
	goto L83
L86:
	;
	v337 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v332))))
	v338 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v333))))
	if v337 == v338 {
		v360 = v337
		goto L89
	} else {
		goto L90
	}
L87:
	;
	goto L88
L88:
	;
	goto L85
L89:
	;
	v362 = int32(1)
	if v360 != 0 {
		v332 = v332 + v362
		v333 = v333 + v362
		v334 = v334 - v362
		goto L84
	} else {
		goto L98
	}
L90:
	;
	if base.Ui32((v337-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	v348 = v337 | int32(32)
	goto L93
L92:
	;
	v348 = v337
	goto L93
L93:
	;
	if base.Ui32((v338-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	v357 = v338 | int32(32)
	goto L96
L95:
	;
	v357 = v338
	goto L96
L96:
	;
	if v348 == v357 {
		v360 = v348
		goto L89
	} else {
		goto L97
	}
L97:
	;
	v379 = v348 - v357
	goto L83
L98:
	;
	goto L88
L99:
	;
	v384 = F_pstrdup(m, v294+int32(8))
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L12
	} else {
		goto L103
	}
L100:
	;
	goto L101
L101:
	;
	v437 = v294
	v438 = int32(566862)
	v439 = int32(9)
	goto L123
L102:
	;
	v429 = F_ParseTzFile(m, v414, l1+int32(1), l2, l3, v218)
	mBase = m.M
	v430 = m.ExcPending
	if v430 != 0 {
		goto L12
	} else {
		goto L120
	}
L103:
	;
	v387 = v24 + int32(3324)
	if v384 != 0 {
		v391 = v384
		goto L105
	} else {
		goto L106
	}
L104:
	;
	if v414 != 0 {
		goto L114
	} else {
		goto L115
	}
L105:
	;
	v393 = F_strspn(m, v391, int32(779075))
	mBase = m.M
	v394 = v393 + v391
	v395 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v394))))
	if v395 == int32(0) {
		goto L108
	} else {
		goto L109
	}
L106:
	;
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v387)))
	if v389 != 0 {
		v391 = v389
		goto L105
	} else {
		goto L107
	}
L107:
	;
	v414 = int32(0)
	goto L104
L108:
	;
	v398 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v387))) = v398
	v414 = v398
	goto L104
L109:
	;
	goto L110
L110:
	;
	v402 = F_strcspn(m, v394, int32(779075))
	mBase = m.M
	v403 = v402 + v394
	v404 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v403))))
	if v404 != 0 {
		goto L111
	} else {
		goto L112
	}
L111:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v387))) = v403 + int32(1)
	v408 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v403))) = uint8(v408)
	v414 = v394
	goto L104
L112:
	;
	goto L113
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v387))) = int32(0)
	v414 = v394
	goto L104
L114:
	;
	v415 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v414))))
	if v415 != 0 {
		goto L102
	} else {
		goto L117
	}
L115:
	;
	goto L116
L116:
	;
	v417 = *(*int32)(unsafe.Add(mBase, _consts[43]))
	*(*int32)(unsafe.Add(mBase, _consts[508])) = v417
	goto L118
L117:
	;
	goto L116
L118:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+116)) = v267
	*(*int32)(unsafe.Add(mBase, uint32(v24)+112)) = l0
	v426 = F_format_elog_string(m, int32(498600), v24+int32(112))
	mBase = m.M
	v427 = m.ExcPending
	if v427 != 0 {
		goto L12
	} else {
		goto L119
	}
L119:
	;
	*(*int32)(unsafe.Add(mBase, _consts[527])) = v426
	goto L51
L120:
	;
	if v429 < int32(0) {
		v1255 = v429
		goto L21
	} else {
		goto L121
	}
L121:
	;
	v1222 = v429
	goto L48
L122:
	;
	if v484 == int32(0) {
		goto L138
	} else {
		goto L139
	}
L123:
	;
	if v439 != 0 {
		goto L125
	} else {
		goto L126
	}
L124:
	;
	v484 = int32(0)
	goto L122
L125:
	;
	v442 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v437))))
	v443 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438))))
	if v442 == v443 {
		v465 = v442
		goto L128
	} else {
		goto L129
	}
L126:
	;
	goto L127
L127:
	;
	goto L124
L128:
	;
	v467 = int32(1)
	if v465 != 0 {
		v437 = v437 + v467
		v438 = v438 + v467
		v439 = v439 - v467
		goto L123
	} else {
		goto L137
	}
L129:
	;
	if base.Ui32((v442-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L130
	} else {
		goto L131
	}
L130:
	;
	v453 = v442 | int32(32)
	goto L132
L131:
	;
	v453 = v442
	goto L132
L132:
	;
	if base.Ui32((v443-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L133
	} else {
		goto L134
	}
L133:
	;
	v462 = v443 | int32(32)
	goto L135
L134:
	;
	v462 = v443
	goto L135
L135:
	;
	if v453 == v462 {
		v465 = v453
		goto L128
	} else {
		goto L136
	}
L136:
	;
	v484 = v453 - v462
	goto L122
L137:
	;
	goto L127
L138:
	;
	v487 = *(*int32)(unsafe.Add(mBase, uint32(v131)+76))
	if v487 < int32(0) {
		goto L143
	} else {
		goto L144
	}
L139:
	;
	goto L140
L140:
	;
	v499 = v24 + int32(3324)
	if v294 != 0 {
		v503 = v294
		goto L150
	} else {
		goto L151
	}
L141:
	;
	if int32(base.Ui32(v492)>>(uint(int32(4))%32))&int32(1) != 0 {
		v1255 = v218
		goto L21
	} else {
		goto L146
	}
L142:
	;
	goto L141
L143:
	;
	v490 = *(*int32)(unsafe.Add(mBase, uint32(v131)))
	v492 = v490
	goto L142
L144:
	;
	goto L145
L145:
	;
	v491 = *(*int32)(unsafe.Add(mBase, uint32(v131)))
	v492 = v491
	goto L142
L146:
	;
	v223 = v267
	v232 = int32(1)
	goto L46
L147:
	;
	v742 = F_strlen(m, v530)
	mBase = m.M
	if base.Ui32(int32(11)) <= base.Ui32(v742) {
		goto L228
	} else {
		goto L229
	}
L148:
	;
	v724 = *(*int32)(unsafe.Add(mBase, _consts[43]))
	*(*int32)(unsafe.Add(mBase, _consts[508])) = v724
	goto L224
L149:
	;
	if v526 == int32(0) {
		goto L159
	} else {
		goto L160
	}
L150:
	;
	v505 = F_strspn(m, v503, int32(779075))
	mBase = m.M
	v506 = v505 + v503
	v507 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v506))))
	if v507 == int32(0) {
		goto L153
	} else {
		goto L154
	}
L151:
	;
	v501 = *(*int32)(unsafe.Add(mBase, uint32(v499)))
	if v501 != 0 {
		v503 = v501
		goto L150
	} else {
		goto L152
	}
L152:
	;
	v526 = int32(0)
	goto L149
L153:
	;
	v510 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v499))) = v510
	v526 = v510
	goto L149
L154:
	;
	goto L155
L155:
	;
	v514 = F_strcspn(m, v506, int32(779075))
	mBase = m.M
	v515 = v514 + v506
	v516 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v515))))
	if v516 != 0 {
		goto L156
	} else {
		goto L157
	}
L156:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v499))) = v515 + int32(1)
	v520 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v515))) = uint8(v520)
	v526 = v506
	goto L149
L157:
	;
	goto L158
L158:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v499))) = int32(0)
	v526 = v506
	goto L149
L159:
	;
	v716 = int32(498488)
	goto L148
L160:
	;
	goto L161
L161:
	;
	v530 = F_pstrdup(m, v526)
	mBase = m.M
	v531 = m.ExcPending
	if v531 != 0 {
		goto L12
	} else {
		goto L162
	}
L162:
	;
	v534 = v24 + int32(3324)
	goto L165
L163:
	;
	if v561 == int32(0) {
		goto L173
	} else {
		goto L174
	}
L164:
	;
	v540 = F_strspn(m, v536, int32(779075))
	mBase = m.M
	v541 = v540 + v536
	v542 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v541))))
	if v542 == int32(0) {
		goto L167
	} else {
		goto L168
	}
L165:
	;
	v536 = *(*int32)(unsafe.Add(mBase, uint32(v534)))
	if v536 != 0 {
		goto L164
	} else {
		goto L166
	}
L166:
	;
	v561 = int32(0)
	goto L163
L167:
	;
	v545 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v534))) = v545
	v561 = v545
	goto L163
L168:
	;
	goto L169
L169:
	;
	v549 = F_strcspn(m, v541, int32(779075))
	mBase = m.M
	v550 = v549 + v541
	v551 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v550))))
	if v551 != 0 {
		goto L170
	} else {
		goto L171
	}
L170:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v534))) = v550 + int32(1)
	v555 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v550))) = uint8(v555)
	v561 = v541
	goto L163
L171:
	;
	goto L172
L172:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v534))) = int32(0)
	v561 = v541
	goto L163
L173:
	;
	v716 = int32(498431)
	goto L148
L174:
	;
	goto L175
L175:
	;
	v565 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v561))))
	if base.Ui32((v565-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		goto L179
	} else {
		goto L180
	}
L176:
	;
	v712 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v707))))
	if v712 == int32(35) {
		v738 = v709
		v739 = v710
		v740 = v711
		goto L147
	} else {
		goto L223
	}
L177:
	;
	v676 = v24 + int32(3324)
	goto L214
L178:
	;
	v665 = F_pstrdup(m, v561)
	mBase = m.M
	v666 = m.ExcPending
	if v666 != 0 {
		goto L12
	} else {
		goto L211
	}
L179:
	;
	v580 = F_strtox_2(m, v561, v24+int32(3320), int32(10), int64(2147483648))
	mBase = m.M
	v581 = base.I32_wrap_i64(v580)
	goto L181
L180:
	;
	switch v565&int32(255) - int32(43) {
	case 0, 2:
		goto L179
	default:
		goto L178
	}
L181:
	;
	v582 = int32(498363)
	v583 = *(*int32)(unsafe.Add(mBase, uint32(v24)+3320))
	if v583 == v561 {
		v716 = v582
		goto L148
	} else {
		goto L182
	}
L182:
	;
	v585 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v583))))
	if v585 != 0 {
		v716 = v582
		goto L148
	} else {
		goto L183
	}
L183:
	;
	v586 = int32(0)
	v589 = v24 + int32(3324)
	goto L186
L184:
	;
	if v616 == int32(0) {
		goto L194
	} else {
		goto L195
	}
L185:
	;
	v595 = F_strspn(m, v591, int32(779075))
	mBase = m.M
	v596 = v595 + v591
	v597 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v596))))
	if v597 == int32(0) {
		goto L188
	} else {
		goto L189
	}
L186:
	;
	v591 = *(*int32)(unsafe.Add(mBase, uint32(v589)))
	if v591 != 0 {
		goto L185
	} else {
		goto L187
	}
L187:
	;
	v616 = int32(0)
	goto L184
L188:
	;
	v600 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v589))) = v600
	v616 = v600
	goto L184
L189:
	;
	goto L190
L190:
	;
	v604 = F_strcspn(m, v596, int32(779075))
	mBase = m.M
	v605 = v604 + v596
	v606 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v605))))
	if v606 != 0 {
		goto L191
	} else {
		goto L192
	}
L191:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v589))) = v605 + int32(1)
	v610 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v605))) = uint8(v610)
	v616 = v596
	goto L184
L192:
	;
	goto L193
L193:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v589))) = int32(0)
	v616 = v596
	goto L184
L194:
	;
	v738 = int32(0)
	v739 = v581
	v740 = v586
	goto L147
L195:
	;
	goto L196
L196:
	;
	v620 = int32(0)
	v624 = v616
	v625 = int32(569469)
	goto L198
L197:
	;
	if v662 != 0 {
		v707 = v616
		v709 = v620
		v710 = v581
		v711 = v586
		goto L176
	} else {
		goto L210
	}
L198:
	;
	v628 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v624))))
	v629 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v625))))
	if v628 == v629 {
		v651 = v628
		goto L200
	} else {
		goto L201
	}
L199:
	;
	v662 = int32(0)
	goto L197
L200:
	;
	v653 = int32(1)
	if v651 != 0 {
		v624 = v624 + v653
		v625 = v625 + v653
		goto L198
	} else {
		goto L209
	}
L201:
	;
	if base.Ui32((v628-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L202
	} else {
		goto L203
	}
L202:
	;
	v639 = v628 | int32(32)
	goto L204
L203:
	;
	v639 = v628
	goto L204
L204:
	;
	if base.Ui32((v629-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L205
	} else {
		goto L206
	}
L205:
	;
	v648 = v629 | int32(32)
	goto L207
L206:
	;
	v648 = v629
	goto L207
L207:
	;
	if v639 == v648 {
		v651 = v639
		goto L200
	} else {
		goto L208
	}
L208:
	;
	v662 = v639 - v648
	goto L197
L209:
	;
	goto L199
L210:
	;
	v671 = v620
	v672 = v581
	v673 = int32(1)
	goto L177
L211:
	;
	v671 = v665
	v672 = int32(0)
	v673 = int32(0)
	goto L177
L212:
	;
	if v703 == int32(0) {
		v738 = v671
		v739 = v672
		v740 = v673
		goto L147
	} else {
		goto L222
	}
L213:
	;
	v682 = F_strspn(m, v678, int32(779075))
	mBase = m.M
	v683 = v682 + v678
	v684 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v683))))
	if v684 == int32(0) {
		goto L216
	} else {
		goto L217
	}
L214:
	;
	v678 = *(*int32)(unsafe.Add(mBase, uint32(v676)))
	if v678 != 0 {
		goto L213
	} else {
		goto L215
	}
L215:
	;
	v703 = int32(0)
	goto L212
L216:
	;
	v687 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v676))) = v687
	v703 = v687
	goto L212
L217:
	;
	goto L218
L218:
	;
	v691 = F_strcspn(m, v683, int32(779075))
	mBase = m.M
	v692 = v691 + v683
	v693 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v692))))
	if v693 != 0 {
		goto L219
	} else {
		goto L220
	}
L219:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v676))) = v692 + int32(1)
	v697 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v692))) = uint8(v697)
	v703 = v683
	goto L212
L220:
	;
	goto L221
L221:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v676))) = int32(0)
	v703 = v683
	goto L212
L222:
	;
	v707 = v703
	v709 = v671
	v710 = v672
	v711 = v673
	goto L176
L223:
	;
	v716 = int32(498316)
	goto L148
L224:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+132)) = v267
	*(*int32)(unsafe.Add(mBase, uint32(v24)+128)) = l0
	v732 = F_format_elog_string(m, v716, v24+int32(128))
	mBase = m.M
	v733 = m.ExcPending
	if v733 != 0 {
		goto L12
	} else {
		goto L225
	}
L225:
	;
	*(*int32)(unsafe.Add(mBase, _consts[527])) = v732
	goto L51
L226:
	;
	v847 = v821
	v848 = v824
	goto L249
L227:
	;
	*(*int32)(unsafe.Add(mBase, _consts[527])) = v839
	goto L51
L228:
	;
	v746 = *(*int32)(unsafe.Add(mBase, _consts[43]))
	*(*int32)(unsafe.Add(mBase, _consts[508])) = v746
	goto L231
L229:
	;
	goto L230
L230:
	;
	if base.Ui32(int32(-100801)) <= base.Ui32(v739-int32(50401)) {
		goto L233
	} else {
		goto L234
	}
L231:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+156)) = v267
	*(*int32)(unsafe.Add(mBase, uint32(v24)+152)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v24)+148)) = int32(10)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+144)) = v530
	v757 = F_format_elog_string(m, int32(498727), v24+int32(144))
	mBase = m.M
	v758 = m.ExcPending
	if v758 != 0 {
		goto L12
	} else {
		goto L232
	}
L232:
	;
	v839 = v757
	goto L227
L233:
	;
	v763 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v530))))
	if v763 != 0 {
		goto L236
	} else {
		goto L237
	}
L234:
	;
	goto L235
L235:
	;
	v828 = *(*int32)(unsafe.Add(mBase, _consts[43]))
	*(*int32)(unsafe.Add(mBase, _consts[508])) = v828
	goto L247
L236:
	;
	v770 = v530
	v771 = v763
	goto L239
L237:
	;
	goto L238
L238:
	;
	v821 = int32(0)
	v822 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v824 = v218 - int32(1)
	if v821 <= v824 {
		goto L226
	} else {
		goto L246
	}
L239:
	;
	v785 = int32(255)
	v786 = v771 & v785
	if base.Ui32((v786-int32(65))&v785) < base.Ui32(int32(26)) {
		goto L242
	} else {
		goto L243
	}
L240:
	;
	goto L238
L241:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v770))) = uint8(v795)
	v798 = v770 + int32(1)
	v799 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v798))))
	if v799 != 0 {
		v770 = v798
		v771 = v799
		goto L239
	} else {
		goto L245
	}
L242:
	;
	v795 = v786 | int32(32)
	goto L244
L243:
	;
	v795 = v786
	goto L244
L244:
	;
	goto L241
L245:
	;
	goto L240
L246:
	;
	v1004 = v821
	goto L50
L247:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+168)) = v267
	*(*int32)(unsafe.Add(mBase, uint32(v24)+164)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v24)+160)) = v739
	v837 = F_format_elog_string(m, int32(498659), v24+int32(160))
	mBase = m.M
	v838 = m.ExcPending
	if v838 != 0 {
		goto L12
	} else {
		goto L248
	}
L248:
	;
	v839 = v837
	goto L227
L249:
	;
	v864 = (v847 + v848) >> (uint(int32(1)) % 32)
	v867 = v822 + v864*int32(24)
	v868 = *(*int32)(unsafe.Add(mBase, uint32(v867)))
	v871 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v868))))
	v872 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v530))))
	if v872 == int32(0) {
		v891 = v871
		v892 = v872
		goto L254
	} else {
		goto L255
	}
L250:
	;
	v905 = *(*int32)(unsafe.Add(mBase, uint32(v867)+4))
	if v905 == int32(0) {
		goto L267
	} else {
		goto L268
	}
L251:
	;
	goto L250
L252:
	;
	if v902 <= v903 {
		v847 = v902
		v848 = v903
		goto L249
	} else {
		goto L265
	}
L253:
	;
	if v893 < int32(0) {
		goto L261
	} else {
		goto L262
	}
L254:
	;
	v893 = v892 - v891
	goto L253
L255:
	;
	if v871 != v872 {
		v891 = v871
		v892 = v872
		goto L254
	} else {
		goto L256
	}
L256:
	;
	v876 = v530
	v877 = v868
	goto L257
L257:
	;
	v880 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v877)+1)))
	v881 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v876)+1)))
	if v881 == int32(0) {
		v891 = v880
		v892 = v881
		goto L254
	} else {
		goto L259
	}
L258:
	;
	v891 = v880
	v892 = v881
	goto L254
L259:
	;
	v884 = int32(1)
	if v880 == v881 {
		v876 = v876 + v884
		v877 = v877 + v884
		goto L257
	} else {
		goto L260
	}
L260:
	;
	goto L258
L261:
	;
	v902 = v847
	v903 = v864 - int32(1)
	goto L252
L262:
	;
	goto L263
L263:
	;
	if v893 == int32(0) {
		goto L251
	} else {
		goto L264
	}
L264:
	;
	v902 = v864 + int32(1)
	v903 = v848
	goto L252
L265:
	;
	v1004 = v902
	goto L50
L266:
	;
	if v232 != 0 {
		goto L285
	} else {
		goto L286
	}
L267:
	;
	if v738 != 0 {
		v943 = v738
		goto L266
	} else {
		goto L270
	}
L268:
	;
	goto L269
L269:
	;
	if v738 == int32(0) {
		goto L273
	} else {
		goto L274
	}
L270:
	;
	v908 = int32(0)
	v909 = *(*int32)(unsafe.Add(mBase, uint32(v867)+8))
	if v909 != v739 {
		v943 = v908
		goto L266
	} else {
		goto L271
	}
L271:
	;
	v911 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v867)+12)))
	if v911 != v740 {
		v943 = v908
		goto L266
	} else {
		goto L272
	}
L272:
	;
	v1199 = v218
	goto L49
L273:
	;
	v943 = int32(0)
	goto L266
L274:
	;
	goto L275
L275:
	;
	v918 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v738))))
	v919 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v905))))
	if v919 == int32(0) {
		v938 = v918
		v939 = v919
		goto L277
	} else {
		goto L278
	}
L276:
	;
	if v939-v938 == int32(0) {
		v1199 = v218
		goto L49
	} else {
		goto L284
	}
L277:
	;
	goto L276
L278:
	;
	if v918 != v919 {
		v938 = v918
		v939 = v919
		goto L277
	} else {
		goto L279
	}
L279:
	;
	v923 = v905
	v924 = v738
	goto L280
L280:
	;
	v927 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v924)+1)))
	v928 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v923)+1)))
	if v928 == int32(0) {
		v938 = v927
		v939 = v928
		goto L277
	} else {
		goto L282
	}
L281:
	;
	v938 = v927
	v939 = v928
	goto L277
L282:
	;
	v931 = int32(1)
	if v927 == v928 {
		v923 = v923 + v931
		v924 = v924 + v931
		goto L280
	} else {
		goto L283
	}
L283:
	;
	goto L281
L284:
	;
	v943 = v738
	goto L266
L285:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v867)+12)) = uint8(v740)
	*(*int32)(unsafe.Add(mBase, uint32(v867)+8)) = v739
	*(*int32)(unsafe.Add(mBase, uint32(v867)+4)) = v943
	v1199 = v218
	goto L49
L286:
	;
	goto L287
L287:
	;
	v948 = *(*int32)(unsafe.Add(mBase, _consts[43]))
	*(*int32)(unsafe.Add(mBase, _consts[508])) = v948
	goto L288
L288:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+192)) = v530
	v956 = F_format_elog_string(m, int32(473238), v24+int32(192))
	mBase = m.M
	v957 = m.ExcPending
	if v957 != 0 {
		goto L12
	} else {
		goto L289
	}
L289:
	;
	*(*int32)(unsafe.Add(mBase, _consts[527])) = v956
	v960 = *(*int32)(unsafe.Add(mBase, _consts[43]))
	*(*int32)(unsafe.Add(mBase, _consts[508])) = v960
	goto L290
L290:
	;
	v963 = *(*int64)(unsafe.Add(mBase, uint32(v867)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+188)) = v267
	*(*int32)(unsafe.Add(mBase, uint32(v24)+184)) = l0
	*(*int64)(unsafe.Add(mBase, uint32(v24)+176)) = base.I64_rotl(v963, int64(32))
	v973 = F_format_elog_string(m, int32(681977), v24+int32(176))
	mBase = m.M
	v974 = m.ExcPending
	if v974 != 0 {
		goto L12
	} else {
		goto L291
	}
L291:
	;
	*(*int32)(unsafe.Add(mBase, _consts[509])) = v973
	goto L51
L292:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v1019 << (uint(int32(1)) % 32)
	v1024 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v1027 = F_repalloc(m, v1024, v1019*int32(48))
	mBase = m.M
	v1028 = m.ExcPending
	if v1028 != 0 {
		goto L12
	} else {
		goto L295
	}
L293:
	;
	v1030 = v822
	goto L294
L294:
	;
	v1031 = int32(24)
	v1033 = v1030 + v1004*v1031
	v1035 = v1033 + v1031
	v1038 = (v218 - v1004) * v1031
	if v1035 == v1033 {
		goto L297
	} else {
		goto L298
	}
L295:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v1027
	v1030 = v1027
	goto L294
L296:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1033)+12)) = uint8(v740)
	*(*int32)(unsafe.Add(mBase, uint32(v1033)+8)) = v739
	*(*int32)(unsafe.Add(mBase, uint32(v1033)+4)) = v738
	*(*int32)(unsafe.Add(mBase, uint32(v1033))) = v530
	v1187 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+237)))
	*(*uint16)(unsafe.Add(mBase, uint32(v1033)+13)) = uint16(v1187)
	v1189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24+int32(239)))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1033)+15)) = uint8(v1189)
	*(*int32)(unsafe.Add(mBase, uint32(v1033)+20)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v1033)+16)) = v267
	v1199 = v218 + int32(1)
	goto L49
L297:
	;
	goto L296
L298:
	;
	v1042 = v1035 + v1038
	if base.Ui32(v1033-v1042) <= base.Ui32(int32(0)-v1038<<(uint(int32(1))%32)) {
		goto L299
	} else {
		goto L300
	}
L299:
	;
	v1049 = F___memcpy(m, v1035, v1033, v1038)
	mBase = m.M
	goto L296
L300:
	;
	goto L301
L301:
	;
	v1052 = (v1035 ^ v1033) & int32(3)
	if base.Ui32(v1035) < base.Ui32(v1033) {
		goto L304
	} else {
		goto L305
	}
L302:
	;
	if v1154 == int32(0) {
		goto L297
	} else {
		goto L338
	}
L303:
	;
	if base.Ui32(v1132) <= base.Ui32(int32(3)) {
		v1153 = v1131
		v1154 = v1132
		v1155 = v1133
		goto L302
	} else {
		goto L334
	}
L304:
	;
	if v1052 != 0 {
		goto L307
	} else {
		goto L308
	}
L305:
	;
	goto L306
L306:
	;
	if v1052 != 0 {
		v1114 = v1038
		goto L317
	} else {
		goto L318
	}
L307:
	;
	v1153 = v1033
	v1154 = v1038
	v1155 = v1035
	goto L302
L308:
	;
	goto L309
L309:
	;
	if v1035&int32(3) == int32(0) {
		goto L310
	} else {
		goto L311
	}
L310:
	;
	v1131 = v1033
	v1132 = v1038
	v1133 = v1035
	goto L303
L311:
	;
	goto L312
L312:
	;
	v1059 = v1033
	v1060 = v1038
	v1061 = v1035
	goto L313
L313:
	;
	if v1060 == int32(0) {
		goto L297
	} else {
		goto L315
	}
L314:
	;
	v1131 = v1068
	v1132 = v1070
	v1133 = v1072
	goto L303
L315:
	;
	v1065 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1059))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1061))) = uint8(v1065)
	v1067 = int32(1)
	v1068 = v1059 + v1067
	v1070 = v1060 - v1067
	v1072 = v1061 + v1067
	if v1072&int32(3) != 0 {
		v1059 = v1068
		v1060 = v1070
		v1061 = v1072
		goto L313
	} else {
		goto L316
	}
L316:
	;
	goto L314
L317:
	;
	if v1114 == int32(0) {
		goto L297
	} else {
		goto L330
	}
L318:
	;
	if v1042&int32(3) != 0 {
		goto L319
	} else {
		goto L320
	}
L319:
	;
	v1079 = v1038
	goto L322
L320:
	;
	v1094 = v1038
	goto L321
L321:
	;
	if base.Ui32(v1094) <= base.Ui32(int32(3)) {
		v1114 = v1094
		goto L317
	} else {
		goto L326
	}
L322:
	;
	if v1079 == int32(0) {
		goto L297
	} else {
		goto L324
	}
L323:
	;
	v1094 = v1085
	goto L321
L324:
	;
	v1085 = v1079 - int32(1)
	v1086 = v1035 + v1085
	v1088 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1033+v1085))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1086))) = uint8(v1088)
	if v1086&int32(3) != 0 {
		v1079 = v1085
		goto L322
	} else {
		goto L325
	}
L325:
	;
	goto L323
L326:
	;
	v1101 = v1094
	goto L327
L327:
	;
	v1105 = v1101 - int32(4)
	v1108 = *(*int32)(unsafe.Add(mBase, uint32(v1033+v1105)))
	*(*int32)(unsafe.Add(mBase, uint32(v1035+v1105))) = v1108
	if base.Ui32(int32(3)) < base.Ui32(v1105) {
		v1101 = v1105
		goto L327
	} else {
		goto L329
	}
L328:
	;
	v1114 = v1105
	goto L317
L329:
	;
	goto L328
L330:
	;
	v1121 = v1114
	goto L331
L331:
	;
	v1125 = v1121 - int32(1)
	v1128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1033+v1125))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1035+v1125))) = uint8(v1128)
	if v1125 != 0 {
		v1121 = v1125
		goto L331
	} else {
		goto L333
	}
L332:
	;
	goto L297
L333:
	;
	goto L332
L334:
	;
	v1138 = v1131
	v1139 = v1132
	v1140 = v1133
	goto L335
L335:
	;
	v1142 = *(*int32)(unsafe.Add(mBase, uint32(v1138)))
	*(*int32)(unsafe.Add(mBase, uint32(v1140))) = v1142
	v1144 = int32(4)
	v1145 = v1138 + v1144
	v1147 = v1140 + v1144
	v1149 = v1139 - v1144
	if base.Ui32(int32(3)) < base.Ui32(v1149) {
		v1138 = v1145
		v1139 = v1149
		v1140 = v1147
		goto L335
	} else {
		goto L337
	}
L336:
	;
	v1153 = v1145
	v1154 = v1149
	v1155 = v1147
	goto L302
L337:
	;
	goto L336
L338:
	;
	v1160 = v1153
	v1161 = v1154
	v1162 = v1155
	goto L339
L339:
	;
	v1164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1160))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1162))) = uint8(v1164)
	v1166 = int32(1)
	v1171 = v1161 - v1166
	if v1171 != 0 {
		v1160 = v1160 + v1166
		v1161 = v1171
		v1162 = v1162 + v1166
		goto L339
	} else {
		goto L341
	}
L340:
	;
	goto L297
L341:
	;
	goto L340
L342:
	;
	v1222 = v1199
	goto L48
L343:
	;
	if int32(base.Ui32(v1244)>>(uint(int32(4))%32))&int32(1) == int32(0) {
		v218 = v1222
		v223 = v267
		goto L46
	} else {
		goto L348
	}
L344:
	;
	goto L343
L345:
	;
	v1242 = *(*int32)(unsafe.Add(mBase, uint32(v131)))
	v1244 = v1242
	goto L344
L346:
	;
	goto L347
L347:
	;
	v1243 = *(*int32)(unsafe.Add(mBase, uint32(v131)))
	v1244 = v1243
	goto L344
L348:
	;
	goto L47
L349:
	;
	v1278 = v1255
	goto L1
}
func F_PersistHoldablePortal(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
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
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
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
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v240 int32
	_ = v240
	var v252 int32
	_ = v252
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v291 int32
	_ = v291
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v324 int64
	_ = v324
	var v325 int32
	_ = v325
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v352 int32
	_ = v352
	var v366 int32
	_ = v366
	var v381 int32
	_ = v381
	var v397 int32
	_ = v397
	var v415 int32
	_ = v415
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v479 int32
	_ = v479
	var v494 int32
	_ = v494
	var v500 int32
	_ = v500
	var v501 int64
	_ = v501
	var v505 int32
	_ = v505
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v512 int32
	_ = v512
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v527 int32
	_ = v527
	v2 = int32(0)
	v21 = m.G0
	v23 = v21 - int32(48)
	m.G0 = v23
	v27 = l0 + int32(108)
	v29 = l0 + int32(88)
	v32 = v2
	v33 = v2
	v34 = v2
	v35 = v2
	v36 = v2
	v37 = v2
	v38 = v2
	v39 = v2
	v40 = v2
	v41 = v2
	v42 = v2
	v43 = int32(-1)
	v44 = v23
	goto L1
L1:
	;
	goto L3
L2:
	;
	m.G0 = v23 + int32(48)
	return
L3:
	;
	if v43 != int32(1) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L2
L5:
	;
	v500 = int32(m.ExcTag)
	v501 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v500 == int32(0) {
		goto L54
	} else {
		goto L55
	}
L6:
	;
	v53 = v44 - int32(160)
	m.G0 = v53
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	v56 = int32(4554128)
	v57 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v59
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+12)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v23)+8)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v23)+20)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v23)+24)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v23)+28)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v23)+32)) = v27
	*(*int32)(unsafe.Add(mBase, uint32(v23)+36)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v23)+40)) = v29
	*(*int32)(unsafe.Add(mBase, uint32(v23)+44)) = v53
	v72 = F_CreateTupleDescCopy(m, v61)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		v494 = v53
		goto L5
	} else {
		goto L9
	}
L7:
	;
	v105 = v32
	v106 = v33
	v107 = v34
	v108 = v35
	v109 = v36
	v110 = v37
	v111 = v38
	v112 = v39
	v113 = v40
	v114 = v41
	v115 = v42
	v116 = v44
	goto L8
L8:
	;
	if v113 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+92)) = v72
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v23)+8)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v23)+12)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v23)+20)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v23)+24)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v23)+28)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v23)+32)) = v27
	*(*int32)(unsafe.Add(mBase, uint32(v23)+36)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v23)+40)) = v29
	*(*int32)(unsafe.Add(mBase, uint32(v23)+44)) = v53
	F_MarkPortalActive(m, l0)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		v494 = v53
		goto L5
	} else {
		goto L10
	}
L10:
	;
	v91 = *(*int32)(unsafe.Add(mBase, _consts[52]))
	v93 = *(*int32)(unsafe.Add(mBase, _consts[53]))
	v95 = *(*int32)(unsafe.Add(mBase, _consts[485]))
	v97 = *(*int32)(unsafe.Add(mBase, _consts[181]))
	v99 = *(*int32)(unsafe.Add(mBase, _consts[486]))
	goto L11
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v53)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v53))) = v23 + int32(4)
	goto L14
L12:
	;
	v105 = v55
	v106 = v53
	v107 = v57
	v108 = v93
	v109 = v91
	v110 = v95
	v111 = v97
	v112 = v99
	v113 = int32(0)
	v114 = v27
	v115 = v29
	v116 = v53
	goto L8
L14:
	;
	goto L12
L15:
	;
	*(*int32)(unsafe.Add(mBase, _consts[52])) = v109
	*(*int32)(unsafe.Add(mBase, _consts[53])) = v108
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v107
	*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = int32(2)
	*(*int32)(unsafe.Add(mBase, _consts[181])) = v111
	*(*int32)(unsafe.Add(mBase, _consts[486])) = v112
	*(*int32)(unsafe.Add(mBase, _consts[485])) = v110
	*(*int32)(unsafe.Add(mBase, uint32(v23)+8)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v23)+12)) = v108
	*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = v110
	*(*int32)(unsafe.Add(mBase, uint32(v23)+20)) = v111
	*(*int32)(unsafe.Add(mBase, uint32(v23)+24)) = v112
	*(*int32)(unsafe.Add(mBase, uint32(v23)+28)) = v107
	*(*int32)(unsafe.Add(mBase, uint32(v23)+32)) = v114
	*(*int32)(unsafe.Add(mBase, uint32(v23)+36)) = v105
	*(*int32)(unsafe.Add(mBase, uint32(v23)+40)) = v115
	*(*int32)(unsafe.Add(mBase, uint32(v23)+44)) = v106
	F_PopActiveSnapshot(m)
	mBase = m.M
	v466 = m.ExcPending
	if v466 != 0 {
		v494 = v116
		goto L5
	} else {
		goto L52
	}
L16:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, _consts[486])) = l0
	*(*int32)(unsafe.Add(mBase, _consts[53])) = v106
	v123 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v123 != 0 {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	goto L19
L19:
	;
	*(*int32)(unsafe.Add(mBase, _consts[52])) = v109
	*(*int32)(unsafe.Add(mBase, _consts[53])) = v108
	*(*int32)(unsafe.Add(mBase, uint32(v23)+8)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v23)+12)) = v108
	*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = v110
	*(*int32)(unsafe.Add(mBase, uint32(v23)+20)) = v111
	*(*int32)(unsafe.Add(mBase, uint32(v23)+24)) = v112
	*(*int32)(unsafe.Add(mBase, uint32(v23)+28)) = v107
	*(*int32)(unsafe.Add(mBase, uint32(v23)+32)) = v114
	*(*int32)(unsafe.Add(mBase, uint32(v23)+36)) = v105
	*(*int32)(unsafe.Add(mBase, uint32(v23)+40)) = v115
	*(*int32)(unsafe.Add(mBase, uint32(v23)+44)) = v106
	F_MarkPortalFailed(m, l0)
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
		v494 = v116
		goto L5
	} else {
		goto L50
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, _consts[181])) = v123
	goto L22
L21:
	;
	goto L22
L22:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v127
	*(*int32)(unsafe.Add(mBase, _consts[485])) = v127
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v105)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+12)) = v108
	*(*int32)(unsafe.Add(mBase, uint32(v23)+8)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = v110
	*(*int32)(unsafe.Add(mBase, uint32(v23)+20)) = v111
	*(*int32)(unsafe.Add(mBase, uint32(v23)+24)) = v112
	*(*int32)(unsafe.Add(mBase, uint32(v23)+28)) = v107
	*(*int32)(unsafe.Add(mBase, uint32(v23)+32)) = v114
	*(*int32)(unsafe.Add(mBase, uint32(v23)+36)) = v105
	*(*int32)(unsafe.Add(mBase, uint32(v23)+40)) = v115
	*(*int32)(unsafe.Add(mBase, uint32(v23)+44)) = v106
	F_PushActiveSnapshot(m, v131)
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		v494 = v116
		goto L5
	} else {
		goto L23
	}
L23:
	;
	v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+76)))
	if v144&int32(2) != 0 {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+12)) = v108
	*(*int32)(unsafe.Add(mBase, uint32(v23)+8)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = v110
	*(*int32)(unsafe.Add(mBase, uint32(v23)+20)) = v111
	*(*int32)(unsafe.Add(mBase, uint32(v23)+24)) = v112
	*(*int32)(unsafe.Add(mBase, uint32(v23)+28)) = v107
	*(*int32)(unsafe.Add(mBase, uint32(v23)+32)) = v114
	*(*int32)(unsafe.Add(mBase, uint32(v23)+36)) = v105
	*(*int32)(unsafe.Add(mBase, uint32(v23)+40)) = v115
	*(*int32)(unsafe.Add(mBase, uint32(v23)+44)) = v106
	v175 = F_CreateDestReceiver(m, int32(6))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		v494 = v116
		goto L5
	} else {
		goto L29
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+12)) = v108
	*(*int32)(unsafe.Add(mBase, uint32(v23)+8)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = v110
	*(*int32)(unsafe.Add(mBase, uint32(v23)+20)) = v111
	*(*int32)(unsafe.Add(mBase, uint32(v23)+24)) = v112
	*(*int32)(unsafe.Add(mBase, uint32(v23)+28)) = v107
	*(*int32)(unsafe.Add(mBase, uint32(v23)+32)) = v114
	*(*int32)(unsafe.Add(mBase, uint32(v23)+36)) = v105
	*(*int32)(unsafe.Add(mBase, uint32(v23)+40)) = v115
	*(*int32)(unsafe.Add(mBase, uint32(v23)+44)) = v106
	F_ExecutorRewind(m, v105)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		v494 = v116
		goto L5
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+117)))
	v163 = v160 ^ int32(1)
	goto L24
L28:
	;
	v163 = int32(1)
	goto L24
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v105)+20)) = v175
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v114)))
	v179 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+12)) = v108
	*(*int32)(unsafe.Add(mBase, uint32(v23)+8)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = v110
	*(*int32)(unsafe.Add(mBase, uint32(v23)+20)) = v111
	*(*int32)(unsafe.Add(mBase, uint32(v23)+24)) = v112
	*(*int32)(unsafe.Add(mBase, uint32(v23)+28)) = v107
	*(*int32)(unsafe.Add(mBase, uint32(v23)+32)) = v114
	*(*int32)(unsafe.Add(mBase, uint32(v23)+36)) = v105
	*(*int32)(unsafe.Add(mBase, uint32(v23)+40)) = v115
	*(*int32)(unsafe.Add(mBase, uint32(v23)+44)) = v106
	v190 = int32(1)
	v191 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v175)+36)) = v191
	*(*int32)(unsafe.Add(mBase, uint32(v175)+32)) = v191
	*(*uint8)(unsafe.Add(mBase, uint32(v175)+28)) = uint8(v190)
	*(*int32)(unsafe.Add(mBase, uint32(v175)+24)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v175)+20)) = v179
	goto L30
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+12)) = v108
	*(*int32)(unsafe.Add(mBase, uint32(v23)+8)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = v110
	*(*int32)(unsafe.Add(mBase, uint32(v23)+20)) = v111
	*(*int32)(unsafe.Add(mBase, uint32(v23)+24)) = v112
	*(*int32)(unsafe.Add(mBase, uint32(v23)+28)) = v107
	*(*int32)(unsafe.Add(mBase, uint32(v23)+32)) = v114
	*(*int32)(unsafe.Add(mBase, uint32(v23)+36)) = v105
	*(*int32)(unsafe.Add(mBase, uint32(v23)+40)) = v115
	*(*int32)(unsafe.Add(mBase, uint32(v23)+44)) = v106
	F_ExecutorRun(m, v105, v163, int64(0))
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		v494 = v116
		goto L5
	} else {
		goto L31
	}
L31:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v105)+20))
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v211)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+12)) = v108
	*(*int32)(unsafe.Add(mBase, uint32(v23)+8)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = v110
	*(*int32)(unsafe.Add(mBase, uint32(v23)+20)) = v111
	*(*int32)(unsafe.Add(mBase, uint32(v23)+24)) = v112
	*(*int32)(unsafe.Add(mBase, uint32(v23)+28)) = v107
	*(*int32)(unsafe.Add(mBase, uint32(v23)+32)) = v114
	*(*int32)(unsafe.Add(mBase, uint32(v23)+36)) = v105
	*(*int32)(unsafe.Add(mBase, uint32(v23)+40)) = v115
	*(*int32)(unsafe.Add(mBase, uint32(v23)+44)) = v106
	m.T0[v212].(func(*base.Module, int32))(m, v211)
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		v494 = v116
		goto L5
	} else {
		goto L32
	}
L32:
	;
	v225 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v105)+20)) = v225
	*(*int32)(unsafe.Add(mBase, uint32(v115))) = v225
	*(*int32)(unsafe.Add(mBase, uint32(v23)+12)) = v108
	*(*int32)(unsafe.Add(mBase, uint32(v23)+8)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = v110
	*(*int32)(unsafe.Add(mBase, uint32(v23)+20)) = v111
	*(*int32)(unsafe.Add(mBase, uint32(v23)+24)) = v112
	*(*int32)(unsafe.Add(mBase, uint32(v23)+28)) = v107
	*(*int32)(unsafe.Add(mBase, uint32(v23)+32)) = v114
	*(*int32)(unsafe.Add(mBase, uint32(v23)+36)) = v105
	*(*int32)(unsafe.Add(mBase, uint32(v23)+40)) = v115
	*(*int32)(unsafe.Add(mBase, uint32(v23)+44)) = v106
	F_ExecutorFinish(m, v105)
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		v494 = v116
		goto L5
	} else {
		goto L33
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+12)) = v108
	*(*int32)(unsafe.Add(mBase, uint32(v23)+8)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = v110
	*(*int32)(unsafe.Add(mBase, uint32(v23)+20)) = v111
	*(*int32)(unsafe.Add(mBase, uint32(v23)+24)) = v112
	*(*int32)(unsafe.Add(mBase, uint32(v23)+28)) = v107
	*(*int32)(unsafe.Add(mBase, uint32(v23)+32)) = v114
	*(*int32)(unsafe.Add(mBase, uint32(v23)+36)) = v105
	*(*int32)(unsafe.Add(mBase, uint32(v23)+40)) = v115
	*(*int32)(unsafe.Add(mBase, uint32(v23)+44)) = v106
	F_ExecutorEnd(m, v105)
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		v494 = v116
		goto L5
	} else {
		goto L34
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+12)) = v108
	*(*int32)(unsafe.Add(mBase, uint32(v23)+8)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = v110
	*(*int32)(unsafe.Add(mBase, uint32(v23)+20)) = v111
	*(*int32)(unsafe.Add(mBase, uint32(v23)+24)) = v112
	*(*int32)(unsafe.Add(mBase, uint32(v23)+28)) = v107
	*(*int32)(unsafe.Add(mBase, uint32(v23)+32)) = v114
	*(*int32)(unsafe.Add(mBase, uint32(v23)+36)) = v105
	*(*int32)(unsafe.Add(mBase, uint32(v23)+40)) = v115
	*(*int32)(unsafe.Add(mBase, uint32(v23)+44)) = v106
	F_FreeQueryDesc(m, v105)
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		v494 = v116
		goto L5
	} else {
		goto L35
	}
L35:
	;
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v114)))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v266
	v268 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+117)))
	if v268 == int32(1) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	goto L39
L37:
	;
	goto L38
L38:
	;
	v306 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+12)) = v108
	*(*int32)(unsafe.Add(mBase, uint32(v23)+8)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = v110
	*(*int32)(unsafe.Add(mBase, uint32(v23)+20)) = v111
	*(*int32)(unsafe.Add(mBase, uint32(v23)+24)) = v112
	*(*int32)(unsafe.Add(mBase, uint32(v23)+28)) = v107
	*(*int32)(unsafe.Add(mBase, uint32(v23)+32)) = v114
	*(*int32)(unsafe.Add(mBase, uint32(v23)+36)) = v105
	*(*int32)(unsafe.Add(mBase, uint32(v23)+40)) = v115
	*(*int32)(unsafe.Add(mBase, uint32(v23)+44)) = v106
	F_tuplestore_rescan(m, v306)
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		v494 = v116
		goto L5
	} else {
		goto L43
	}
L39:
	;
	v291 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+12)) = v108
	*(*int32)(unsafe.Add(mBase, uint32(v23)+8)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = v110
	*(*int32)(unsafe.Add(mBase, uint32(v23)+20)) = v111
	*(*int32)(unsafe.Add(mBase, uint32(v23)+24)) = v112
	*(*int32)(unsafe.Add(mBase, uint32(v23)+28)) = v107
	*(*int32)(unsafe.Add(mBase, uint32(v23)+32)) = v114
	*(*int32)(unsafe.Add(mBase, uint32(v23)+36)) = v105
	*(*int32)(unsafe.Add(mBase, uint32(v23)+40)) = v115
	*(*int32)(unsafe.Add(mBase, uint32(v23)+44)) = v106
	v304 = F_tuplestore_skiptuples(m, v291, int64(1000000), int32(1))
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		v494 = v116
		goto L5
	} else {
		goto L41
	}
L41:
	;
	if v304 != 0 {
		goto L39
	} else {
		goto L42
	}
L42:
	;
	goto L15
L43:
	;
	v319 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+76)))
	if v319&int32(2) == int32(0) {
		goto L15
	} else {
		goto L44
	}
L44:
	;
	v324 = *(*int64)(unsafe.Add(mBase, uint32(l0)+120))
	v325 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+12)) = v108
	*(*int32)(unsafe.Add(mBase, uint32(v23)+8)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = v110
	*(*int32)(unsafe.Add(mBase, uint32(v23)+20)) = v111
	*(*int32)(unsafe.Add(mBase, uint32(v23)+24)) = v112
	*(*int32)(unsafe.Add(mBase, uint32(v23)+28)) = v107
	*(*int32)(unsafe.Add(mBase, uint32(v23)+32)) = v114
	*(*int32)(unsafe.Add(mBase, uint32(v23)+36)) = v105
	*(*int32)(unsafe.Add(mBase, uint32(v23)+40)) = v115
	*(*int32)(unsafe.Add(mBase, uint32(v23)+44)) = v106
	v337 = F_tuplestore_skiptuples(m, v325, v324, int32(1))
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		v494 = v116
		goto L5
	} else {
		goto L45
	}
L45:
	;
	if v337 != 0 {
		goto L15
	} else {
		goto L46
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+12)) = v108
	*(*int32)(unsafe.Add(mBase, uint32(v23)+8)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = v110
	*(*int32)(unsafe.Add(mBase, uint32(v23)+20)) = v111
	*(*int32)(unsafe.Add(mBase, uint32(v23)+24)) = v112
	*(*int32)(unsafe.Add(mBase, uint32(v23)+28)) = v107
	*(*int32)(unsafe.Add(mBase, uint32(v23)+32)) = v114
	*(*int32)(unsafe.Add(mBase, uint32(v23)+36)) = v105
	*(*int32)(unsafe.Add(mBase, uint32(v23)+40)) = v115
	*(*int32)(unsafe.Add(mBase, uint32(v23)+44)) = v106
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		v494 = v116
		goto L5
	} else {
		goto L47
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+12)) = v108
	*(*int32)(unsafe.Add(mBase, uint32(v23)+8)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = v110
	*(*int32)(unsafe.Add(mBase, uint32(v23)+20)) = v111
	*(*int32)(unsafe.Add(mBase, uint32(v23)+24)) = v112
	*(*int32)(unsafe.Add(mBase, uint32(v23)+28)) = v107
	*(*int32)(unsafe.Add(mBase, uint32(v23)+32)) = v114
	*(*int32)(unsafe.Add(mBase, uint32(v23)+36)) = v105
	*(*int32)(unsafe.Add(mBase, uint32(v23)+40)) = v115
	*(*int32)(unsafe.Add(mBase, uint32(v23)+44)) = v106
	F_errmsg_internal(m, int32(305541), int32(0))
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		v494 = v116
		goto L5
	} else {
		goto L48
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+12)) = v108
	*(*int32)(unsafe.Add(mBase, uint32(v23)+8)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = v110
	*(*int32)(unsafe.Add(mBase, uint32(v23)+20)) = v111
	*(*int32)(unsafe.Add(mBase, uint32(v23)+24)) = v112
	*(*int32)(unsafe.Add(mBase, uint32(v23)+28)) = v107
	*(*int32)(unsafe.Add(mBase, uint32(v23)+32)) = v114
	*(*int32)(unsafe.Add(mBase, uint32(v23)+36)) = v105
	*(*int32)(unsafe.Add(mBase, uint32(v23)+40)) = v115
	*(*int32)(unsafe.Add(mBase, uint32(v23)+44)) = v106
	F_errfinish(m, int32(517246), int32(470), int32(324293))
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		v494 = v116
		goto L5
	} else {
		goto L49
	}
L49:
	;
	goto L16
L50:
	;
	*(*int32)(unsafe.Add(mBase, _consts[181])) = v111
	*(*int32)(unsafe.Add(mBase, _consts[486])) = v112
	*(*int32)(unsafe.Add(mBase, _consts[485])) = v110
	*(*int32)(unsafe.Add(mBase, uint32(v23)+8)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v23)+12)) = v108
	*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = v110
	*(*int32)(unsafe.Add(mBase, uint32(v23)+20)) = v111
	*(*int32)(unsafe.Add(mBase, uint32(v23)+24)) = v112
	*(*int32)(unsafe.Add(mBase, uint32(v23)+28)) = v107
	*(*int32)(unsafe.Add(mBase, uint32(v23)+32)) = v114
	*(*int32)(unsafe.Add(mBase, uint32(v23)+36)) = v105
	*(*int32)(unsafe.Add(mBase, uint32(v23)+40)) = v115
	*(*int32)(unsafe.Add(mBase, uint32(v23)+44)) = v106
	F_pg_re_throw(m)
	mBase = m.M
	v415 = m.ExcPending
	if v415 != 0 {
		v494 = v116
		goto L5
	} else {
		goto L51
	}
L51:
	;
	goto L16
L52:
	;
	v467 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+12)) = v108
	*(*int32)(unsafe.Add(mBase, uint32(v23)+8)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = v110
	*(*int32)(unsafe.Add(mBase, uint32(v23)+20)) = v111
	*(*int32)(unsafe.Add(mBase, uint32(v23)+24)) = v112
	*(*int32)(unsafe.Add(mBase, uint32(v23)+28)) = v107
	*(*int32)(unsafe.Add(mBase, uint32(v23)+32)) = v114
	*(*int32)(unsafe.Add(mBase, uint32(v23)+36)) = v105
	*(*int32)(unsafe.Add(mBase, uint32(v23)+40)) = v115
	*(*int32)(unsafe.Add(mBase, uint32(v23)+44)) = v106
	F_MemoryContextDeleteChildren(m, v467)
	mBase = m.M
	v479 = m.ExcPending
	if v479 != 0 {
		v494 = v116
		goto L5
	} else {
		goto L53
	}
L53:
	;
	goto L4
L54:
	;
	v505 = int32(v501)
	m.G0 = v494
	v507 = *(*int32)(unsafe.Add(mBase, uint32(v505)+4))
	v508 = *(*int32)(unsafe.Add(mBase, uint32(v505)))
	v512 = *(*int32)(unsafe.Add(mBase, uint32(v508)))
	if v23+int32(4) == v512 {
		goto L57
	} else {
		goto L58
	}
L55:
	;
	m.ExcPending = 1
	goto L63
L56:
	;
	if v515 != 0 {
		goto L60
	} else {
		goto L61
	}
L57:
	;
	v514 = *(*int32)(unsafe.Add(mBase, uint32(v508)+4))
	v515 = v514
	goto L59
L58:
	;
	v515 = int32(0)
	goto L59
L59:
	;
	goto L56
L60:
	;
	v516 = *(*int32)(unsafe.Add(mBase, uint32(v23)+44))
	v517 = *(*int32)(unsafe.Add(mBase, uint32(v23)+40))
	v518 = *(*int32)(unsafe.Add(mBase, uint32(v23)+36))
	v519 = *(*int32)(unsafe.Add(mBase, uint32(v23)+32))
	v520 = *(*int32)(unsafe.Add(mBase, uint32(v23)+28))
	v521 = *(*int32)(unsafe.Add(mBase, uint32(v23)+24))
	v522 = *(*int32)(unsafe.Add(mBase, uint32(v23)+20))
	v523 = *(*int32)(unsafe.Add(mBase, uint32(v23)+16))
	v524 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
	v525 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	v32 = v518
	v33 = v516
	v34 = v520
	v35 = v524
	v36 = v525
	v37 = v523
	v38 = v522
	v39 = v521
	v40 = v507
	v41 = v519
	v42 = v517
	v43 = v515
	v44 = v494
	goto L1
L61:
	;
	goto L62
L62:
	;
	F___wasm_longjmp(m, v508, v507)
	mBase = m.M
	v527 = m.ExcPending
	if v527 != 0 {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	return
L64:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_PrefetchBuffer(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int64
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
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
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int64
	_ = v95
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+118)))
	if v12 == int32(116) {
		v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+24)))
		if v15 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v131 = m.ExcPending
			if v131 != 0 {
				return
			} else {
				F_errcode(m, int32(1088))
				mBase = m.M
				v134 = m.ExcPending
				if v134 != 0 {
					return
				} else {
					F_errmsg(m, int32(152323), int32(0))
					mBase = m.M
					v138 = m.ExcPending
					if v138 != 0 {
						return
					} else {
						F_errfinish(m, int32(518165), int32(662), int32(237608))
						mBase = m.M
						v143 = m.ExcPending
						if v143 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		} else {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
			if v18 == int32(0) {
				v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
				v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v22
				v24 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
				*(*int64)(unsafe.Add(mBase, uint32(v9))) = v24
				v26 = F_smgropen(m, v9, v21)
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v26
					v30 = *(*int32)(unsafe.Add(mBase, uint32(v26)+72))
					if v30 != 0 {
						v38 = v30
					} else {
						v31 = *(*int32)(unsafe.Add(mBase, uint32(v26)+76))
						v32 = *(*int32)(unsafe.Add(mBase, uint32(v26)+80))
						*(*int32)(unsafe.Add(mBase, uint32(v31)+4)) = v32
						v34 = *(*int32)(unsafe.Add(mBase, uint32(v26)+76))
						*(*int32)(unsafe.Add(mBase, uint32(v32))) = v34
						v36 = *(*int32)(unsafe.Add(mBase, uint32(v26)+72))
						v38 = v36
					}
					*(*int32)(unsafe.Add(mBase, uint32(v26)+72)) = v38 + int32(1)
					v42 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
					v43 = v42
					v44 = m.G0
					v46 = v44 - int32(32)
					m.G0 = v46
					*(*int64)(unsafe.Add(mBase, uint32(l0))) = int64(0)
					v50 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
					*(*int32)(unsafe.Add(mBase, uint32(v46)+12)) = v50
					v52 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
					*(*int32)(unsafe.Add(mBase, uint32(v46)+16)) = v52
					v54 = *(*int32)(unsafe.Add(mBase, uint32(v43)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v46)+28)) = l2
					*(*int32)(unsafe.Add(mBase, uint32(v46)+24)) = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v46)+20)) = v54
					v60 = *(*int32)(unsafe.Add(mBase, _consts[788]))
					if v60 != 0 {
						v65 = v60
						v68 = int32(0)
						v70 = F_hash_search(m, v65, v46+int32(12), v68, v68)
						mBase = m.M
						v71 = m.ExcPending
						if v71 != 0 {
							return
						} else {
							if v70 != 0 {
								v72 = *(*int32)(unsafe.Add(mBase, uint32(v70)+20))
								*(*int32)(unsafe.Add(mBase, uint32(l0))) = v72 ^ int32(-1)
								m.G0 = v46 + int32(32)
								m.G0 = v9 + int32(32)
								return
							} else {
								v77 = int32(*(*uint8)(unsafe.Add(mBase, _consts[207])))
								if v77&int32(1) != 0 {
									m.G0 = v46 + int32(32)
									m.G0 = v9 + int32(32)
									return
								} else {
									v82 = F_smgrprefetch(m, v43, int32(0), l2, int32(1))
									mBase = m.M
									v83 = m.ExcPending
									if v83 != 0 {
										return
									} else {
										if v82 == int32(0) {
										} else {
											v86 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)) = uint8(v86)
										}
										m.G0 = v46 + int32(32)
										m.G0 = v9 + int32(32)
										return
									}
								}
							}
						}
					} else {
						F_InitLocalBuffers(m)
						mBase = m.M
						v62 = m.ExcPending
						if v62 != 0 {
							return
						} else {
							v64 = *(*int32)(unsafe.Add(mBase, _consts[788]))
							v65 = v64
							v68 = int32(0)
							v70 = F_hash_search(m, v65, v46+int32(12), v68, v68)
							mBase = m.M
							v71 = m.ExcPending
							if v71 != 0 {
								return
							} else {
								if v70 != 0 {
									v72 = *(*int32)(unsafe.Add(mBase, uint32(v70)+20))
									*(*int32)(unsafe.Add(mBase, uint32(l0))) = v72 ^ int32(-1)
									m.G0 = v46 + int32(32)
									m.G0 = v9 + int32(32)
									return
								} else {
									v77 = int32(*(*uint8)(unsafe.Add(mBase, _consts[207])))
									if v77&int32(1) != 0 {
										m.G0 = v46 + int32(32)
										m.G0 = v9 + int32(32)
										return
									} else {
										v82 = F_smgrprefetch(m, v43, int32(0), l2, int32(1))
										mBase = m.M
										v83 = m.ExcPending
										if v83 != 0 {
											return
										} else {
											if v82 == int32(0) {
											} else {
												v86 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)) = uint8(v86)
											}
											m.G0 = v46 + int32(32)
											m.G0 = v9 + int32(32)
											return
										}
									}
								}
							}
						}
					}
				}
			} else {
				v43 = v18
				v44 = m.G0
				v46 = v44 - int32(32)
				m.G0 = v46
				*(*int64)(unsafe.Add(mBase, uint32(l0))) = int64(0)
				v50 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
				*(*int32)(unsafe.Add(mBase, uint32(v46)+12)) = v50
				v52 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v46)+16)) = v52
				v54 = *(*int32)(unsafe.Add(mBase, uint32(v43)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v46)+28)) = l2
				*(*int32)(unsafe.Add(mBase, uint32(v46)+24)) = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v46)+20)) = v54
				v60 = *(*int32)(unsafe.Add(mBase, _consts[788]))
				if v60 != 0 {
					v65 = v60
					v68 = int32(0)
					v70 = F_hash_search(m, v65, v46+int32(12), v68, v68)
					mBase = m.M
					v71 = m.ExcPending
					if v71 != 0 {
						return
					} else {
						if v70 != 0 {
							v72 = *(*int32)(unsafe.Add(mBase, uint32(v70)+20))
							*(*int32)(unsafe.Add(mBase, uint32(l0))) = v72 ^ int32(-1)
							m.G0 = v46 + int32(32)
							m.G0 = v9 + int32(32)
							return
						} else {
							v77 = int32(*(*uint8)(unsafe.Add(mBase, _consts[207])))
							if v77&int32(1) != 0 {
								m.G0 = v46 + int32(32)
								m.G0 = v9 + int32(32)
								return
							} else {
								v82 = F_smgrprefetch(m, v43, int32(0), l2, int32(1))
								mBase = m.M
								v83 = m.ExcPending
								if v83 != 0 {
									return
								} else {
									if v82 == int32(0) {
									} else {
										v86 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)) = uint8(v86)
									}
									m.G0 = v46 + int32(32)
									m.G0 = v9 + int32(32)
									return
								}
							}
						}
					}
				} else {
					F_InitLocalBuffers(m)
					mBase = m.M
					v62 = m.ExcPending
					if v62 != 0 {
						return
					} else {
						v64 = *(*int32)(unsafe.Add(mBase, _consts[788]))
						v65 = v64
						v68 = int32(0)
						v70 = F_hash_search(m, v65, v46+int32(12), v68, v68)
						mBase = m.M
						v71 = m.ExcPending
						if v71 != 0 {
							return
						} else {
							if v70 != 0 {
								v72 = *(*int32)(unsafe.Add(mBase, uint32(v70)+20))
								*(*int32)(unsafe.Add(mBase, uint32(l0))) = v72 ^ int32(-1)
								m.G0 = v46 + int32(32)
								m.G0 = v9 + int32(32)
								return
							} else {
								v77 = int32(*(*uint8)(unsafe.Add(mBase, _consts[207])))
								if v77&int32(1) != 0 {
									m.G0 = v46 + int32(32)
									m.G0 = v9 + int32(32)
									return
								} else {
									v82 = F_smgrprefetch(m, v43, int32(0), l2, int32(1))
									mBase = m.M
									v83 = m.ExcPending
									if v83 != 0 {
										return
									} else {
										if v82 == int32(0) {
										} else {
											v86 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)) = uint8(v86)
										}
										m.G0 = v46 + int32(32)
										m.G0 = v9 + int32(32)
										return
									}
								}
							}
						}
					}
				}
			}
		}
	} else {
		v91 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
		if v91 != 0 {
			v117 = v91
			F_PrefetchSharedBuffer(m, l0, v117, int32(0), l2)
			mBase = m.M
			v120 = m.ExcPending
			if v120 != 0 {
				return
			} else {
				m.G0 = v9 + int32(32)
				return
			}
		} else {
			v92 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
			v93 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = v93
			v95 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
			*(*int64)(unsafe.Add(mBase, uint32(v9)+16)) = v95
			v99 = F_smgropen(m, v9+int32(16), v92)
			mBase = m.M
			v100 = m.ExcPending
			if v100 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v99
				v103 = *(*int32)(unsafe.Add(mBase, uint32(v99)+72))
				if v103 != 0 {
					v111 = v103
				} else {
					v104 = *(*int32)(unsafe.Add(mBase, uint32(v99)+76))
					v105 = *(*int32)(unsafe.Add(mBase, uint32(v99)+80))
					*(*int32)(unsafe.Add(mBase, uint32(v104)+4)) = v105
					v107 = *(*int32)(unsafe.Add(mBase, uint32(v99)+76))
					*(*int32)(unsafe.Add(mBase, uint32(v105))) = v107
					v109 = *(*int32)(unsafe.Add(mBase, uint32(v99)+72))
					v111 = v109
				}
				*(*int32)(unsafe.Add(mBase, uint32(v99)+72)) = v111 + int32(1)
				v115 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
				v117 = v115
				F_PrefetchSharedBuffer(m, l0, v117, int32(0), l2)
				mBase = m.M
				v120 = m.ExcPending
				if v120 != 0 {
					return
				} else {
					m.G0 = v9 + int32(32)
					return
				}
			}
		}
	}
}
func F_PrepareRedoRemove(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v21 int32
	_ = v21
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	v3 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	v13 = *(*int32)(unsafe.Add(mBase, _consts[166]))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if v14 <= v3 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L10
	} else {
		goto L28
	}
L2:
	;
	m.G0 = v10 + int32(32)
	return
L3:
	;
	v21 = v3
	goto L4
L4:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v13+int32(8)+v21<<(uint(int32(2))%32))))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+32))
	if l0 != v30 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v37 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L10
	} else {
		goto L11
	}
L6:
	;
	v33 = v21 + int32(1)
	if v14 != v33 {
		v21 = v33
		goto L4
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	goto L5
L9:
	;
	goto L2
L10:
	;
	return
L11:
	;
	if v37 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = l0
	F_errmsg_internal(m, int32(48935), v10+int32(16))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L10
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+45)))
	if v50 == int32(1) {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	F_errfinish(m, int32(522203), int32(2601), int32(359163))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L10
	} else {
		goto L16
	}
L16:
	;
	goto L14
L17:
	;
	F_RemoveTwoPhaseFile(m, l0, l1)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L10
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v56 = *(*int32)(unsafe.Add(mBase, _consts[166]))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)+4))
	if v57 <= int32(0) {
		goto L1
	} else {
		goto L21
	}
L20:
	;
	goto L19
L21:
	;
	v61 = v56 + int32(8)
	v65 = int32(0)
	goto L22
L22:
	;
	v72 = v61 + v65<<(uint(int32(2))%32)
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
	if v73 != v29 {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	v79 = v57 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v56)+4)) = v79
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v61+v79<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v72))) = v84
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
	*(*int32)(unsafe.Add(mBase, uint32(v29))) = v86
	*(*int32)(unsafe.Add(mBase, uint32(v56))) = v29
	goto L2
L24:
	;
	v76 = v65 + int32(1)
	if v57 != v76 {
		v65 = v76
		goto L22
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	goto L23
L27:
	;
	goto L1
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v29
	F_errmsg_internal(m, int32(25639), v10)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L10
	} else {
		goto L29
	}
L29:
	;
	F_errfinish(m, int32(522203), int32(650), int32(119268))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L10
	} else {
		goto L30
	}
L30:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_ProcedureCreate(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32, l12 int32, l13 int32, l14 int32, l15 int32, l16 int32, l17 int32, l18 int32, l19 int32, l20 int32, l21 int32, l22 int32, l23 int32, l24 int32, l25 int32, l26 int32, l27 float32, l28 float32) {
	mBase := m.M
	_ = mBase
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v120 int32
	_ = v120
	var v133 int32
	_ = v133
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
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v190 int32
	_ = v190
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v242 int32
	_ = v242
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v257 int32
	_ = v257
	var v263 int32
	_ = v263
	var v267 int32
	_ = v267
	var v272 int32
	_ = v272
	var v276 int32
	_ = v276
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v286 int32
	_ = v286
	var v290 int32
	_ = v290
	var v295 int32
	_ = v295
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v306 int32
	_ = v306
	var v311 int32
	_ = v311
	var v316 int32
	_ = v316
	var v320 int32
	_ = v320
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v362 int32
	_ = v362
	var v371 int64
	_ = v371
	var v384 int32
	_ = v384
	var v385 int64
	_ = v385
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v423 int32
	_ = v423
	var v426 int32
	_ = v426
	var v429 int32
	_ = v429
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v436 int32
	_ = v436
	var v439 int32
	_ = v439
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v447 int32
	_ = v447
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v454 int32
	_ = v454
	var v457 int32
	_ = v457
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v489 int32
	_ = v489
	var v492 int32
	_ = v492
	var v496 int32
	_ = v496
	var v498 int32
	_ = v498
	var v504 int32
	_ = v504
	var v509 int32
	_ = v509
	var v515 int32
	_ = v515
	var v521 int32
	_ = v521
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v527 int32
	_ = v527
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v542 int32
	_ = v542
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v556 int32
	_ = v556
	var v558 int32
	_ = v558
	var v564 int32
	_ = v564
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v575 int32
	_ = v575
	var v578 int32
	_ = v578
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	var v591 int32
	_ = v591
	var v593 int32
	_ = v593
	var v598 int32
	_ = v598
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v665 int32
	_ = v665
	var v676 int32
	_ = v676
	var v678 int32
	_ = v678
	var v681 int32
	_ = v681
	var v686 int32
	_ = v686
	var v687 int32
	_ = v687
	var v691 int32
	_ = v691
	var v692 int32
	_ = v692
	var v695 int32
	_ = v695
	var v696 int32
	_ = v696
	var v699 int32
	_ = v699
	var v706 int32
	_ = v706
	var v707 int32
	_ = v707
	var v711 int32
	_ = v711
	var v753 int32
	_ = v753
	var v756 int32
	_ = v756
	var v758 int32
	_ = v758
	var v762 int32
	_ = v762
	var v763 int32
	_ = v763
	var v764 int32
	_ = v764
	var v765 int32
	_ = v765
	var v766 int32
	_ = v766
	var v767 int32
	_ = v767
	var v768 int32
	_ = v768
	var v770 int32
	_ = v770
	var v773 int32
	_ = v773
	var v776 int32
	_ = v776
	var v777 int32
	_ = v777
	var v813 int32
	_ = v813
	var v816 int32
	_ = v816
	var v823 int32
	_ = v823
	var v824 int32
	_ = v824
	var v828 int32
	_ = v828
	var v829 int32
	_ = v829
	var v830 int32
	_ = v830
	var v831 int32
	_ = v831
	var v832 int32
	_ = v832
	var v835 int32
	_ = v835
	var v837 int32
	_ = v837
	var v838 int32
	_ = v838
	var v843 int32
	_ = v843
	var v845 int32
	_ = v845
	var v846 int32
	_ = v846
	var v888 int32
	_ = v888
	var v901 int32
	_ = v901
	var v902 int32
	_ = v902
	var v906 int32
	_ = v906
	var v908 int32
	_ = v908
	var v910 int32
	_ = v910
	var v911 int32
	_ = v911
	var v913 int32
	_ = v913
	var v915 int32
	_ = v915
	var v916 int32
	_ = v916
	var v918 int32
	_ = v918
	var v919 int32
	_ = v919
	var v921 int32
	_ = v921
	var v925 int32
	_ = v925
	var v926 int32
	_ = v926
	var v932 int32
	_ = v932
	var v933 int32
	_ = v933
	var v935 int32
	_ = v935
	var v936 int32
	_ = v936
	var v937 int32
	_ = v937
	var v939 int32
	_ = v939
	var v959 int32
	_ = v959
	var v972 int32
	_ = v972
	var v974 int32
	_ = v974
	var v981 int32
	_ = v981
	var v982 int32
	_ = v982
	var v983 int32
	_ = v983
	var v996 int32
	_ = v996
	var v1005 int32
	_ = v1005
	var v1014 int32
	_ = v1014
	var v1045 int32
	_ = v1045
	var v1060 int32
	_ = v1060
	var v1067 int32
	_ = v1067
	var v1069 int32
	_ = v1069
	var v1113 int32
	_ = v1113
	var v1114 int32
	_ = v1114
	var v1147 int32
	_ = v1147
	var v1157 int32
	_ = v1157
	var v1161 int32
	_ = v1161
	var v1170 int32
	_ = v1170
	var v1172 int32
	_ = v1172
	var v1173 int32
	_ = v1173
	var v1223 int32
	_ = v1223
	var v1226 int32
	_ = v1226
	var v1228 int32
	_ = v1228
	var v1235 int32
	_ = v1235
	var v1238 int32
	_ = v1238
	var v1243 int32
	_ = v1243
	var v1246 int32
	_ = v1246
	var v1250 int32
	_ = v1250
	var v1252 int32
	_ = v1252
	var v1254 int32
	_ = v1254
	var v1256 int32
	_ = v1256
	var v1259 int32
	_ = v1259
	var v1262 int32
	_ = v1262
	var v1266 int32
	_ = v1266
	var v1268 int32
	_ = v1268
	var v1272 int32
	_ = v1272
	var v1273 int32
	_ = v1273
	var v1275 int32
	_ = v1275
	var v1277 int32
	_ = v1277
	var v1279 int32
	_ = v1279
	var v1283 int32
	_ = v1283
	var v1284 int32
	_ = v1284
	var v1285 int32
	_ = v1285
	var v1289 int32
	_ = v1289
	var v1291 int32
	_ = v1291
	var v1292 int32
	_ = v1292
	var v1295 int32
	_ = v1295
	var v1301 int32
	_ = v1301
	var v1304 int32
	_ = v1304
	var v1311 int32
	_ = v1311
	var v1314 int32
	_ = v1314
	var v1320 int32
	_ = v1320
	var v1325 int32
	_ = v1325
	var v1329 int32
	_ = v1329
	var v1332 int32
	_ = v1332
	var v1335 int32
	_ = v1335
	var v1338 int32
	_ = v1338
	var v1339 int32
	_ = v1339
	var v1340 int32
	_ = v1340
	var v1341 int32
	_ = v1341
	var v1348 int32
	_ = v1348
	var v1353 int32
	_ = v1353
	var v1357 int32
	_ = v1357
	var v1360 int32
	_ = v1360
	var v1364 int32
	_ = v1364
	var v1368 int32
	_ = v1368
	var v1369 int32
	_ = v1369
	var v1370 int32
	_ = v1370
	var v1371 int32
	_ = v1371
	var v1378 int32
	_ = v1378
	var v1383 int32
	_ = v1383
	var v1388 int32
	_ = v1388
	var v1391 int32
	_ = v1391
	var v1392 int32
	_ = v1392
	var v1396 int32
	_ = v1396
	var v1402 int32
	_ = v1402
	var v1403 int32
	_ = v1403
	var v1404 int32
	_ = v1404
	var v1405 int32
	_ = v1405
	var v1412 int32
	_ = v1412
	var v1417 int32
	_ = v1417
	var v1421 int32
	_ = v1421
	var v1424 int32
	_ = v1424
	var v1428 int32
	_ = v1428
	var v1429 int32
	_ = v1429
	var v1430 int32
	_ = v1430
	var v1431 int32
	_ = v1431
	var v1438 int32
	_ = v1438
	var v1443 int32
	_ = v1443
	var v1447 int32
	_ = v1447
	var v1450 int32
	_ = v1450
	var v1454 int32
	_ = v1454
	var v1455 int32
	_ = v1455
	var v1456 int32
	_ = v1456
	var v1457 int32
	_ = v1457
	var v1464 int32
	_ = v1464
	var v1469 int32
	_ = v1469
	var v1473 int32
	_ = v1473
	var v1476 int32
	_ = v1476
	var v1480 int32
	_ = v1480
	var v1486 int32
	_ = v1486
	var v1491 int32
	_ = v1491
	var v1495 int32
	_ = v1495
	var v1498 int32
	_ = v1498
	var v1502 int32
	_ = v1502
	var v1508 int32
	_ = v1508
	var v1513 int32
	_ = v1513
	var v1517 int32
	_ = v1517
	var v1520 int32
	_ = v1520
	var v1524 int32
	_ = v1524
	var v1530 int32
	_ = v1530
	var v1535 int32
	_ = v1535
	var v1539 int32
	_ = v1539
	var v1542 int32
	_ = v1542
	var v1546 int32
	_ = v1546
	var v1552 int32
	_ = v1552
	var v1557 int32
	_ = v1557
	var v1561 int32
	_ = v1561
	var v1565 int32
	_ = v1565
	var v1570 int32
	_ = v1570
	var v1576 int32
	_ = v1576
	var v1580 int32
	_ = v1580
	var v1585 int32
	_ = v1585
	var v1589 int32
	_ = v1589
	var v1592 int32
	_ = v1592
	var v1593 int32
	_ = v1593
	var v1599 int32
	_ = v1599
	var v1604 int32
	_ = v1604
	v41 = m.G0
	v43 = v41 - int32(496)
	m.G0 = v43
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l18)+16))
	if base.Ui32(v45) < base.Ui32(int32(101)) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	if l19 != 0 {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	goto L3
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1589 = m.ExcPending
	if v1589 != 0 {
		goto L20
	} else {
		goto L376
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1576 = m.ExcPending
	if v1576 != 0 {
		goto L20
	} else {
		goto L373
	}
L5:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l19)+4))
	if v48 != int32(1) {
		goto L4
	} else {
		goto L8
	}
L6:
	;
	v58 = l18
	v59 = v45
	goto L7
L7:
	;
	if l20 != 0 {
		goto L13
	} else {
		goto L14
	}
L8:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l19)+16))
	if v51 <= int32(0) {
		goto L4
	} else {
		goto L9
	}
L9:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l19)+8))
	if v54 != 0 {
		goto L4
	} else {
		goto L10
	}
L10:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l19)+12))
	if v55 != int32(26) {
		goto L4
	} else {
		goto L11
	}
L11:
	;
	v58 = l19
	v59 = v51
	goto L7
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1561 = m.ExcPending
	if v1561 != 0 {
		goto L20
	} else {
		goto L370
	}
L13:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l20)+4))
	if v60 != int32(1) {
		goto L12
	} else {
		goto L16
	}
L14:
	;
	v72 = int32(0)
	goto L15
L15:
	;
	v74 = l18 + int32(24)
	v75 = F_check_valid_polymorphic_signature(m, l5, v74, v45)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L20
	} else {
		goto L21
	}
L16:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l20)+16))
	if v63 != v59 {
		goto L12
	} else {
		goto L17
	}
L17:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l20)+8))
	if v65 != 0 {
		goto L12
	} else {
		goto L18
	}
L18:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l20)+12))
	if v66 != int32(18) {
		goto L12
	} else {
		goto L19
	}
L19:
	;
	v72 = l20 + int32(24)
	goto L15
L20:
	;
	return
L21:
	;
	if v75 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v79 = F_check_valid_internal_signature(m, l5, v74, v45)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L20
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1539 = m.ExcPending
	if v1539 != 0 {
		goto L20
	} else {
		goto L365
	}
L25:
	;
	if v79 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v84 = v58 + int32(24)
	if l19 == int32(0) {
		goto L31
	} else {
		goto L32
	}
L27:
	;
	goto L28
L28:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1517 = m.ExcPending
	if v1517 != 0 {
		goto L20
	} else {
		goto L360
	}
L29:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1495 = m.ExcPending
	if v1495 != 0 {
		goto L20
	} else {
		goto L355
	}
L30:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1473 = m.ExcPending
	if v1473 != 0 {
		goto L20
	} else {
		goto L350
	}
L31:
	;
	v190 = int32(0)
	if v72 == v190 {
		v362 = v190
		goto L44
	} else {
		goto L45
	}
L32:
	;
	if v59 == int32(0) {
		goto L31
	} else {
		goto L33
	}
L33:
	;
	v120 = int32(0)
	goto L34
L34:
	;
	if v72 == int32(0) {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	goto L31
L36:
	;
	v148 = v120 + int32(1)
	if v148 != v59 {
		v120 = v148
		goto L34
	} else {
		goto L43
	}
L37:
	;
	v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v120+v72))))
	switch v133 - int32(105) {
	case 0, 13:
		goto L36
	default:
		goto L38
	}
L38:
	;
	v138 = v84 + v120<<(uint(int32(2))%32)
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v138)))
	v140 = F_check_valid_polymorphic_signature(m, v139, v74, v45)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L20
	} else {
		goto L39
	}
L39:
	;
	if v140 != 0 {
		goto L29
	} else {
		goto L40
	}
L40:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v138)))
	v143 = F_check_valid_internal_signature(m, v142, v74, v45)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L20
	} else {
		goto L41
	}
L41:
	;
	if v143 != 0 {
		goto L30
	} else {
		goto L42
	}
L42:
	;
	goto L36
L43:
	;
	goto L35
L44:
	;
	v371 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v43)+486)) = v371
	*(*int64)(unsafe.Add(mBase, uint32(v43)+480)) = v371
	*(*int64)(unsafe.Add(mBase, uint32(v43)+472)) = v371
	*(*int64)(unsafe.Add(mBase, uint32(v43)+464)) = v371
	v384 = F__emscripten_memset_bulkmem(m, v43+int32(336), base.I32_extend8_s(int32(0)), int32(120))
	mBase = m.M
	goto L81
L45:
	;
	if v59 == int32(0) {
		v362 = v190
		goto L44
	} else {
		goto L46
	}
L46:
	;
	v228 = int32(0)
	v229 = v190
	goto L47
L47:
	;
	v238 = v228 + v72
	v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v238))))
	switch v239 - int32(98) {
	case 0, 7:
		goto L56
	default:
		goto L53
	case 13:
		goto L55
	case 18:
		v327 = v229
		goto L49
	case 20:
		goto L54
	}
L48:
	;
	v362 = v327
	goto L44
L49:
	;
	v329 = v228 + int32(1)
	if v329 != v59 {
		v228 = v329
		v229 = v327
		goto L47
	} else {
		goto L80
	}
L50:
	;
	v327 = int32(5077)
	goto L49
L51:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L20
	} else {
		goto L77
	}
L52:
	;
	v327 = int32(2283)
	goto L49
L53:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L20
	} else {
		goto L74
	}
L54:
	;
	if v229 != 0 {
		goto L51
	} else {
		goto L66
	}
L55:
	;
	if l12 != int32(112) {
		v327 = v229
		goto L49
	} else {
		goto L61
	}
L56:
	;
	v242 = int32(0)
	if v229 == v242 {
		v327 = v242
		goto L49
	} else {
		goto L57
	}
L57:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L20
	} else {
		goto L58
	}
L58:
	;
	F_errmsg_internal(m, int32(83962), int32(0))
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L20
	} else {
		goto L59
	}
L59:
	;
	F_errfinish(m, int32(523705), int32(279), int32(372796))
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L20
	} else {
		goto L60
	}
L60:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L61:
	;
	if v229 == int32(0) {
		v327 = v229
		goto L49
	} else {
		goto L62
	}
L62:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L20
	} else {
		goto L63
	}
L63:
	;
	F_errmsg_internal(m, int32(83962), int32(0))
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L20
	} else {
		goto L64
	}
L64:
	;
	F_errfinish(m, int32(523705), int32(283), int32(372796))
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L20
	} else {
		goto L65
	}
L65:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L66:
	;
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v84+v228<<(uint(int32(2))%32))))
	switch v276 - int32(2276) {
	case 0:
		v327 = v276
		goto L49
	case 1:
		goto L52
	default:
		goto L67
	}
L67:
	;
	if v276 == int32(5078) {
		goto L50
	} else {
		goto L68
	}
L68:
	;
	v281 = F_get_element_type(m, v276)
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L20
	} else {
		goto L69
	}
L69:
	;
	if v281 != 0 {
		v327 = v281
		goto L49
	} else {
		goto L70
	}
L70:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L20
	} else {
		goto L71
	}
L71:
	;
	F_errmsg_internal(m, int32(25684), int32(0))
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L20
	} else {
		goto L72
	}
L72:
	;
	F_errfinish(m, int32(523705), int32(305), int32(372796))
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L20
	} else {
		goto L73
	}
L73:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L74:
	;
	v300 = int32(*(*int8)(unsafe.Add(mBase, uint32(v238))))
	*(*int32)(unsafe.Add(mBase, uint32(v43)+16)) = v300
	F_errmsg_internal(m, int32(719696), v43+int32(16))
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L20
	} else {
		goto L75
	}
L75:
	;
	F_errfinish(m, int32(523705), int32(310), int32(372796))
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L20
	} else {
		goto L76
	}
L76:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L77:
	;
	F_errmsg_internal(m, int32(83962), int32(0))
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L20
	} else {
		goto L78
	}
L78:
	;
	F_errfinish(m, int32(523705), int32(290), int32(372796))
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L20
	} else {
		goto L79
	}
L79:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L80:
	;
	goto L48
L81:
	;
	v385 = int64(72340172838076673)
	*(*int64)(unsafe.Add(mBase, uint32(v43)+326)) = v385
	*(*int64)(unsafe.Add(mBase, uint32(v43)+320)) = v385
	*(*int64)(unsafe.Add(mBase, uint32(v43)+312)) = v385
	*(*int64)(unsafe.Add(mBase, uint32(v43)+304)) = v385
	v396 = F_strncpy(m, v43+int32(240), l1, int32(64))
	mBase = m.M
	v397 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v396)+63)) = uint8(v397)
	goto L82
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43)+400)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v43)+396)) = l17
	*(*int32)(unsafe.Add(mBase, uint32(v43)+392)) = l16
	*(*int32)(unsafe.Add(mBase, uint32(v43)+388)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v43)+384)) = l15
	*(*int32)(unsafe.Add(mBase, uint32(v43)+380)) = l14
	*(*int32)(unsafe.Add(mBase, uint32(v43)+376)) = l13
	*(*int32)(unsafe.Add(mBase, uint32(v43)+372)) = l12
	*(*int32)(unsafe.Add(mBase, uint32(v43)+368)) = l26
	*(*int32)(unsafe.Add(mBase, uint32(v43)+364)) = v362
	*(*float32)(unsafe.Add(mBase, uint32(v43)+360)) = l28
	*(*float32)(unsafe.Add(mBase, uint32(v43)+356)) = l27
	*(*int32)(unsafe.Add(mBase, uint32(v43)+352)) = l7
	*(*int32)(unsafe.Add(mBase, uint32(v43)+348)) = l6
	*(*int32)(unsafe.Add(mBase, uint32(v43)+344)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v43)+340)) = v43 + int32(240)
	if l22 != 0 {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v417 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l22)+4)))
	v418 = v417
	goto L85
L84:
	;
	v418 = int32(0)
	goto L85
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43)+412)) = l18
	*(*int32)(unsafe.Add(mBase, uint32(v43)+408)) = l5
	*(*int32)(unsafe.Add(mBase, uint32(v43)+404)) = v418
	if l19 != 0 {
		goto L87
	} else {
		goto L88
	}
L86:
	;
	if l20 != 0 {
		goto L91
	} else {
		goto L92
	}
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43)+416)) = l19
	goto L86
L88:
	;
	goto L89
L89:
	;
	v423 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v43)+484)) = uint8(v423)
	goto L86
L90:
	;
	if l21 != 0 {
		goto L95
	} else {
		goto L96
	}
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43)+420)) = l20
	goto L90
L92:
	;
	goto L93
L93:
	;
	v426 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v43)+485)) = uint8(v426)
	goto L90
L94:
	;
	if l22 != 0 {
		goto L99
	} else {
		goto L100
	}
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43)+424)) = l21
	goto L94
L96:
	;
	goto L97
L97:
	;
	v429 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v43)+486)) = uint8(v429)
	goto L94
L98:
	;
	if l23 != 0 {
		goto L105
	} else {
		goto L106
	}
L99:
	;
	v431 = F_nodeToString(m, l22)
	mBase = m.M
	v432 = m.ExcPending
	if v432 != 0 {
		goto L20
	} else {
		goto L102
	}
L100:
	;
	goto L101
L101:
	;
	v436 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v43)+487)) = uint8(v436)
	goto L98
L102:
	;
	v433 = F_cstring_to_text(m, v431)
	mBase = m.M
	v434 = m.ExcPending
	if v434 != 0 {
		goto L20
	} else {
		goto L103
	}
L103:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43)+428)) = v433
	goto L98
L104:
	;
	v441 = F_cstring_to_text(m, l9)
	mBase = m.M
	v442 = m.ExcPending
	if v442 != 0 {
		goto L20
	} else {
		goto L108
	}
L105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43)+432)) = l23
	goto L104
L106:
	;
	goto L107
L107:
	;
	v439 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v43)+488)) = uint8(v439)
	goto L104
L108:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43)+436)) = v441
	if l10 != 0 {
		goto L110
	} else {
		goto L111
	}
L109:
	;
	if l11 != 0 {
		goto L115
	} else {
		goto L116
	}
L110:
	;
	v444 = F_cstring_to_text(m, l10)
	mBase = m.M
	v445 = m.ExcPending
	if v445 != 0 {
		goto L20
	} else {
		goto L113
	}
L111:
	;
	goto L112
L112:
	;
	v447 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v43)+490)) = uint8(v447)
	goto L109
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43)+440)) = v444
	goto L109
L114:
	;
	if l25 != 0 {
		goto L121
	} else {
		goto L122
	}
L115:
	;
	v449 = F_nodeToString(m, l11)
	mBase = m.M
	v450 = m.ExcPending
	if v450 != 0 {
		goto L20
	} else {
		goto L118
	}
L116:
	;
	goto L117
L117:
	;
	v454 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v43)+491)) = uint8(v454)
	goto L114
L118:
	;
	v451 = F_cstring_to_text(m, v449)
	mBase = m.M
	v452 = m.ExcPending
	if v452 != 0 {
		goto L20
	} else {
		goto L119
	}
L119:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43)+444)) = v451
	goto L114
L120:
	;
	v461 = F_table_open(m, int32(1255), int32(3))
	mBase = m.M
	v462 = m.ExcPending
	if v462 != 0 {
		goto L20
	} else {
		goto L124
	}
L121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43)+448)) = l25
	goto L120
L122:
	;
	goto L123
L123:
	;
	v457 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v43)+492)) = uint8(v457)
	goto L120
L124:
	;
	v463 = *(*int32)(unsafe.Add(mBase, uint32(v461)+52))
	v465 = F_SearchSysCache3(m, int32(46), l1, l18, l2)
	mBase = m.M
	v466 = m.ExcPending
	if v466 != 0 {
		goto L20
	} else {
		goto L132
	}
L125:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1447 = m.ExcPending
	if v1447 != 0 {
		goto L20
	} else {
		goto L344
	}
L126:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1421 = m.ExcPending
	if v1421 != 0 {
		goto L20
	} else {
		goto L338
	}
L127:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1388 = m.ExcPending
	if v1388 != 0 {
		goto L20
	} else {
		goto L332
	}
L128:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1357 = m.ExcPending
	if v1357 != 0 {
		goto L20
	} else {
		goto L325
	}
L129:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1329 = m.ExcPending
	if v1329 != 0 {
		goto L20
	} else {
		goto L316
	}
L130:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1311 = m.ExcPending
	if v1311 != 0 {
		goto L20
	} else {
		goto L312
	}
L131:
	;
	v981 = F_new_object_addresses(m)
	mBase = m.M
	v982 = m.ExcPending
	if v982 != 0 {
		goto L20
	} else {
		goto L247
	}
L132:
	;
	if v465 != 0 {
		goto L133
	} else {
		goto L134
	}
L133:
	;
	if l3 == int32(0) {
		goto L130
	} else {
		goto L136
	}
L134:
	;
	goto L135
L135:
	;
	v918 = F_get_user_default_acl(m, int32(19), l6, l2)
	mBase = m.M
	v919 = m.ExcPending
	if v919 != 0 {
		goto L20
	} else {
		goto L240
	}
L136:
	;
	v470 = *(*int32)(unsafe.Add(mBase, uint32(v465)+16))
	v471 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v470)+22)))
	v472 = v470 + v471
	v473 = *(*int32)(unsafe.Add(mBase, uint32(v472)))
	v474 = F_object_ownercheck(m, int32(1255), v473, l6)
	mBase = m.M
	v475 = m.ExcPending
	if v475 != 0 {
		goto L20
	} else {
		goto L137
	}
L137:
	;
	if v474 == int32(0) {
		goto L138
	} else {
		goto L139
	}
L138:
	;
	F_aclcheck_error(m, int32(2), int32(19), l1)
	mBase = m.M
	v481 = m.ExcPending
	if v481 != 0 {
		goto L20
	} else {
		goto L141
	}
L139:
	;
	goto L140
L140:
	;
	v482 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v472)+96)))
	if v482 != l12&int32(255) {
		goto L142
	} else {
		goto L143
	}
L141:
	;
	goto L140
L142:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v489 = m.ExcPending
	if v489 != 0 {
		goto L20
	} else {
		goto L145
	}
L143:
	;
	goto L144
L144:
	;
	if l12 == int32(97) {
		goto L155
	} else {
		goto L156
	}
L145:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v492 = m.ExcPending
	if v492 != 0 {
		goto L20
	} else {
		goto L146
	}
L146:
	;
	F_errmsg(m, int32(445251), int32(0))
	mBase = m.M
	v496 = m.ExcPending
	if v496 != 0 {
		goto L20
	} else {
		goto L147
	}
L147:
	;
	v498 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v472)+96)))
	switch v498 - int32(97) {
	case 0:
		v504 = int32(640676)
		goto L149
	default:
		goto L148
	case 5:
		goto L152
	case 15:
		goto L151
	case 22:
		goto L150
	}
L148:
	;
	F_errfinish(m, int32(523705), int32(421), int32(372796))
	mBase = m.M
	v515 = m.ExcPending
	if v515 != 0 {
		goto L20
	} else {
		goto L154
	}
L149:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43)+128)) = l1
	F_errdetail(m, v504, v43+int32(128))
	mBase = m.M
	v509 = m.ExcPending
	if v509 != 0 {
		goto L20
	} else {
		goto L153
	}
L150:
	;
	v504 = int32(640609)
	goto L149
L151:
	;
	v504 = int32(660669)
	goto L149
L152:
	;
	v504 = int32(640707)
	goto L149
L153:
	;
	goto L148
L154:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L155:
	;
	v521 = int32(563300)
	goto L157
L156:
	;
	v521 = int32(553872)
	goto L157
L157:
	;
	v523 = base.B2i32(l12 == int32(112))
	if l12 == int32(112) {
		goto L158
	} else {
		goto L159
	}
L158:
	;
	v524 = int32(564248)
	goto L160
L159:
	;
	v524 = v521
	goto L160
L160:
	;
	v525 = *(*int32)(unsafe.Add(mBase, uint32(v472)+108))
	if l5 != v525 {
		goto L129
	} else {
		goto L161
	}
L161:
	;
	v527 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v472)+100)))
	if v527 != l4 {
		goto L129
	} else {
		goto L162
	}
L162:
	;
	if l5 != int32(2249) {
		goto L163
	} else {
		goto L164
	}
L163:
	;
	v610 = F_SysCacheGetAttr(m, int32(46), v465, int32(23), v43+int32(227))
	mBase = m.M
	v611 = m.ExcPending
	if v611 != 0 {
		goto L20
	} else {
		goto L185
	}
L164:
	;
	v531 = F_build_function_result_tupdesc_t(m, v465)
	mBase = m.M
	v532 = m.ExcPending
	if v532 != 0 {
		goto L20
	} else {
		goto L165
	}
L165:
	;
	v533 = F_build_function_result_tupdesc_d(m, l12, l19, l20, l21)
	mBase = m.M
	v534 = m.ExcPending
	if v534 != 0 {
		goto L20
	} else {
		goto L166
	}
L166:
	;
	if v531|v533 == int32(0) {
		goto L163
	} else {
		goto L167
	}
L167:
	;
	if v531 == int32(0) {
		goto L128
	} else {
		goto L168
	}
L168:
	;
	if v533 == int32(0) {
		goto L128
	} else {
		goto L169
	}
L169:
	;
	v542 = int32(0)
	v546 = *(*int32)(unsafe.Add(mBase, uint32(v531)))
	v547 = *(*int32)(unsafe.Add(mBase, uint32(v533)))
	if v546 != v547 {
		v598 = v542
		goto L171
	} else {
		goto L172
	}
L170:
	;
	if v598 == int32(0) {
		goto L128
	} else {
		goto L184
	}
L171:
	;
	goto L170
L172:
	;
	v549 = *(*int32)(unsafe.Add(mBase, uint32(v531)+4))
	v550 = *(*int32)(unsafe.Add(mBase, uint32(v533)+4))
	if v549 != v550 {
		v598 = v542
		goto L171
	} else {
		goto L173
	}
L173:
	;
	if v546 <= int32(0) {
		v598 = int32(1)
		goto L171
	} else {
		goto L174
	}
L174:
	;
	v556 = v546 << (uint(int32(4)) % 32)
	v558 = int32(20)
	v564 = int32(0)
	goto L175
L175:
	;
	v571 = v564 * int32(100)
	v572 = v531 + v556 + v558 + v571
	v573 = int32(4)
	v575 = v571 + (v533 + v556 + v558)
	v578 = F_strcmp(m, v572+v573, v575+v573)
	mBase = m.M
	if v578 != 0 {
		goto L177
	} else {
		goto L178
	}
L176:
	;
	v598 = int32(0)
	goto L171
L177:
	;
	goto L176
L178:
	;
	v579 = *(*int32)(unsafe.Add(mBase, uint32(v572)+68))
	v580 = *(*int32)(unsafe.Add(mBase, uint32(v575)+68))
	if v579 != v580 {
		goto L177
	} else {
		goto L179
	}
L179:
	;
	v582 = *(*int32)(unsafe.Add(mBase, uint32(v572)+76))
	v583 = *(*int32)(unsafe.Add(mBase, uint32(v575)+76))
	if v582 != v583 {
		goto L177
	} else {
		goto L180
	}
L180:
	;
	v585 = *(*int32)(unsafe.Add(mBase, uint32(v572)+96))
	v586 = *(*int32)(unsafe.Add(mBase, uint32(v575)+96))
	if v585 != v586 {
		goto L177
	} else {
		goto L181
	}
L181:
	;
	v588 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v572)+91)))
	v589 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v575)+91)))
	if v588 != v589 {
		goto L177
	} else {
		goto L182
	}
L182:
	;
	v591 = int32(1)
	v593 = v564 + v591
	if v546 != v593 {
		v564 = v593
		goto L175
	} else {
		goto L183
	}
L183:
	;
	v598 = v591
	goto L171
L184:
	;
	goto L163
L185:
	;
	v612 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+227)))
	if v612 != 0 {
		goto L186
	} else {
		goto L187
	}
L186:
	;
	v753 = int32(*(*int16)(unsafe.Add(mBase, uint32(v472)+106)))
	if v753 == int32(0) {
		goto L212
	} else {
		goto L213
	}
L187:
	;
	v613 = int32(0)
	v619 = F_SysCacheGetAttr(m, int32(46), v465, int32(22), v43+int32(227))
	mBase = m.M
	v620 = m.ExcPending
	if v620 != 0 {
		goto L20
	} else {
		goto L188
	}
L188:
	;
	v621 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+227)))
	if v621 != 0 {
		goto L189
	} else {
		goto L190
	}
L189:
	;
	v622 = v613
	goto L191
L190:
	;
	v622 = v619
	goto L191
L191:
	;
	v625 = F_get_func_input_arg_names(m, v610, v622, v43+int32(228))
	mBase = m.M
	v626 = m.ExcPending
	if v626 != 0 {
		goto L20
	} else {
		goto L192
	}
L192:
	;
	v629 = F_get_func_input_arg_names(m, l21, l20, v43+int32(220))
	mBase = m.M
	v630 = m.ExcPending
	if v630 != 0 {
		goto L20
	} else {
		goto L193
	}
L193:
	;
	if v625 <= int32(0) {
		goto L186
	} else {
		goto L194
	}
L194:
	;
	v633 = *(*int32)(unsafe.Add(mBase, uint32(v43)+220))
	v634 = *(*int32)(unsafe.Add(mBase, uint32(v43)+228))
	v665 = v613
	goto L195
L195:
	;
	v676 = v665 << (uint(int32(2)) % 32)
	v678 = *(*int32)(unsafe.Add(mBase, uint32(v634+v676)))
	if v678 != 0 {
		goto L197
	} else {
		goto L198
	}
L196:
	;
	goto L186
L197:
	;
	if v629 <= v665 {
		goto L127
	} else {
		goto L200
	}
L198:
	;
	goto L199
L199:
	;
	v711 = v665 + int32(1)
	if v711 != v625 {
		v665 = v711
		goto L195
	} else {
		goto L211
	}
L200:
	;
	v681 = *(*int32)(unsafe.Add(mBase, uint32(v633+v676)))
	if v681 == int32(0) {
		goto L127
	} else {
		goto L201
	}
L201:
	;
	v686 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v681))))
	v687 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v678))))
	if v687 == int32(0) {
		v706 = v686
		v707 = v687
		goto L203
	} else {
		goto L204
	}
L202:
	;
	if v707-v706 != 0 {
		goto L127
	} else {
		goto L210
	}
L203:
	;
	goto L202
L204:
	;
	if v686 != v687 {
		v706 = v686
		v707 = v687
		goto L203
	} else {
		goto L205
	}
L205:
	;
	v691 = v678
	v692 = v681
	goto L206
L206:
	;
	v695 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v692)+1)))
	v696 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v691)+1)))
	if v696 == int32(0) {
		v706 = v695
		v707 = v696
		goto L203
	} else {
		goto L208
	}
L207:
	;
	v706 = v695
	v707 = v696
	goto L203
L208:
	;
	v699 = int32(1)
	if v695 == v696 {
		v691 = v691 + v699
		v692 = v692 + v699
		goto L206
	} else {
		goto L209
	}
L209:
	;
	goto L207
L210:
	;
	goto L199
L211:
	;
	goto L196
L212:
	;
	v888 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v43)+333)) = uint8(v888)
	*(*uint8)(unsafe.Add(mBase, uint32(v43)+307)) = uint8(v888)
	*(*uint8)(unsafe.Add(mBase, uint32(v43)+304)) = uint8(v888)
	v901 = F_heap_modify_tuple(m, v465, v463, v43+int32(336), v43+int32(464), v43+int32(304))
	mBase = m.M
	v902 = m.ExcPending
	if v902 != 0 {
		goto L20
	} else {
		goto L235
	}
L213:
	;
	if l22 != 0 {
		goto L214
	} else {
		goto L215
	}
L214:
	;
	v756 = *(*int32)(unsafe.Add(mBase, uint32(l22)+4))
	v758 = v756
	goto L216
L215:
	;
	v758 = int32(0)
	goto L216
L216:
	;
	if v758 < v753 {
		goto L126
	} else {
		goto L217
	}
L217:
	;
	v762 = F_SysCacheGetAttrNotNull(m, int32(46), v465, int32(24))
	mBase = m.M
	v763 = m.ExcPending
	if v763 != 0 {
		goto L20
	} else {
		goto L218
	}
L218:
	;
	v764 = F_text_to_cstring(m, v762)
	mBase = m.M
	v765 = m.ExcPending
	if v765 != 0 {
		goto L20
	} else {
		goto L219
	}
L219:
	;
	v766 = F_stringToNode(m, v764)
	mBase = m.M
	v767 = m.ExcPending
	if v767 != 0 {
		goto L20
	} else {
		goto L220
	}
L220:
	;
	if l22 != 0 {
		goto L221
	} else {
		goto L222
	}
L221:
	;
	v768 = *(*int32)(unsafe.Add(mBase, uint32(l22)+4))
	v770 = v768
	goto L223
L222:
	;
	v770 = int32(0)
	goto L223
L223:
	;
	if v766 == int32(0) {
		goto L212
	} else {
		goto L224
	}
L224:
	;
	v773 = *(*int32)(unsafe.Add(mBase, uint32(v766)+4))
	if v773 <= int32(0) {
		goto L212
	} else {
		goto L225
	}
L225:
	;
	v776 = *(*int32)(unsafe.Add(mBase, uint32(l22)+12))
	v777 = int32(*(*int16)(unsafe.Add(mBase, uint32(v472)+106)))
	v813 = v776 + (v770-v777)<<(uint(int32(2))%32)
	v816 = int32(0)
	goto L226
L226:
	;
	v823 = *(*int32)(unsafe.Add(mBase, uint32(v813)))
	v824 = *(*int32)(unsafe.Add(mBase, uint32(v766)+12))
	v828 = *(*int32)(unsafe.Add(mBase, uint32(v824+v816<<(uint(int32(2))%32))))
	v829 = F_exprType(m, v828)
	mBase = m.M
	v830 = m.ExcPending
	if v830 != 0 {
		goto L20
	} else {
		goto L228
	}
L227:
	;
	goto L212
L228:
	;
	v831 = F_exprType(m, v823)
	mBase = m.M
	v832 = m.ExcPending
	if v832 != 0 {
		goto L20
	} else {
		goto L229
	}
L229:
	;
	if v829 != v831 {
		goto L125
	} else {
		goto L230
	}
L230:
	;
	v835 = v813 + int32(4)
	v837 = *(*int32)(unsafe.Add(mBase, uint32(l22)+12))
	v838 = *(*int32)(unsafe.Add(mBase, uint32(l22)+4))
	if base.Ui32(v835) < base.Ui32(v837+v838<<(uint(int32(2))%32)) {
		goto L231
	} else {
		goto L232
	}
L231:
	;
	v843 = v835
	goto L233
L232:
	;
	v843 = int32(0)
	goto L233
L233:
	;
	v845 = v816 + int32(1)
	v846 = *(*int32)(unsafe.Add(mBase, uint32(v766)+4))
	if v845 < v846 {
		v813 = v843
		v816 = v845
		goto L226
	} else {
		goto L234
	}
L234:
	;
	goto L227
L235:
	;
	F_CatalogTupleUpdate(m, v461, v901+int32(4), v901)
	mBase = m.M
	v906 = m.ExcPending
	if v906 != 0 {
		goto L20
	} else {
		goto L236
	}
L236:
	;
	F_ReleaseCatCache(m, v465)
	mBase = m.M
	v908 = m.ExcPending
	if v908 != 0 {
		goto L20
	} else {
		goto L237
	}
L237:
	;
	v910 = *(*int32)(unsafe.Add(mBase, uint32(v901)+16))
	v911 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v910)+22)))
	v913 = *(*int32)(unsafe.Add(mBase, uint32(v910+v911)))
	v915 = F_deleteDependencyRecordsFor(m, int32(1255), v913, int32(1))
	mBase = m.M
	v916 = m.ExcPending
	if v916 != 0 {
		goto L20
	} else {
		goto L238
	}
L238:
	;
	v959 = v888
	v972 = v913
	v974 = v901
	goto L131
L239:
	;
	v925 = F_GetNewOidWithIndex(m, v461, int32(2690), int32(1))
	mBase = m.M
	v926 = m.ExcPending
	if v926 != 0 {
		goto L20
	} else {
		goto L244
	}
L240:
	;
	if v918 != 0 {
		goto L241
	} else {
		goto L242
	}
L241:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43)+452)) = v918
	goto L239
L242:
	;
	goto L243
L243:
	;
	v921 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v43)+493)) = uint8(v921)
	goto L239
L244:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43)+336)) = v925
	v932 = F_heap_form_tuple(m, v463, v43+int32(336), v43+int32(464))
	mBase = m.M
	v933 = m.ExcPending
	if v933 != 0 {
		goto L20
	} else {
		goto L245
	}
L245:
	;
	F_CatalogTupleInsert(m, v461, v932)
	mBase = m.M
	v935 = m.ExcPending
	if v935 != 0 {
		goto L20
	} else {
		goto L246
	}
L246:
	;
	v936 = *(*int32)(unsafe.Add(mBase, uint32(v932)+16))
	v937 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v936)+22)))
	v939 = *(*int32)(unsafe.Add(mBase, uint32(v936+v937)))
	v959 = v918
	v972 = v939
	v974 = v932
	goto L131
L247:
	;
	v983 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v983
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v972
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1255)
	*(*int32)(unsafe.Add(mBase, uint32(v43)+236)) = v983
	*(*int32)(unsafe.Add(mBase, uint32(v43)+232)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v43)+228)) = int32(2615)
	F_add_exact_object_address(m, v43+int32(228), v981)
	mBase = m.M
	v996 = m.ExcPending
	if v996 != 0 {
		goto L20
	} else {
		goto L248
	}
L248:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43)+236)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v43)+232)) = l7
	*(*int32)(unsafe.Add(mBase, uint32(v43)+228)) = int32(2612)
	F_add_exact_object_address(m, v43+int32(228), v981)
	mBase = m.M
	v1005 = m.ExcPending
	if v1005 != 0 {
		goto L20
	} else {
		goto L249
	}
L249:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43)+236)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v43)+232)) = l5
	*(*int32)(unsafe.Add(mBase, uint32(v43)+228)) = int32(1247)
	F_add_exact_object_address(m, v43+int32(228), v981)
	mBase = m.M
	v1014 = m.ExcPending
	if v1014 != 0 {
		goto L20
	} else {
		goto L250
	}
L250:
	;
	if v59 != 0 {
		goto L251
	} else {
		goto L252
	}
L251:
	;
	v1045 = int32(0)
	goto L254
L252:
	;
	goto L253
L253:
	;
	if l24 == int32(0) {
		goto L258
	} else {
		goto L259
	}
L254:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43)+228)) = int32(1247)
	v1060 = *(*int32)(unsafe.Add(mBase, uint32(v84+v1045<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v43)+236)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v43)+232)) = v1060
	F_add_exact_object_address(m, v43+int32(228), v981)
	mBase = m.M
	v1067 = m.ExcPending
	if v1067 != 0 {
		goto L20
	} else {
		goto L256
	}
L255:
	;
	goto L253
L256:
	;
	v1069 = v1045 + int32(1)
	if v1069 != v59 {
		v1045 = v1069
		goto L254
	} else {
		goto L257
	}
L257:
	;
	goto L255
L258:
	;
	if l26 != 0 {
		goto L265
	} else {
		goto L266
	}
L259:
	;
	v1113 = int32(0)
	v1114 = *(*int32)(unsafe.Add(mBase, uint32(l24)+4))
	if v1114 <= v1113 {
		goto L258
	} else {
		goto L260
	}
L260:
	;
	v1147 = v1113
	goto L261
L261:
	;
	v1157 = *(*int32)(unsafe.Add(mBase, uint32(l24)+12))
	v1161 = *(*int32)(unsafe.Add(mBase, uint32(v1157+v1147<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v43)+236)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v43)+232)) = v1161
	*(*int32)(unsafe.Add(mBase, uint32(v43)+228)) = int32(3576)
	F_add_exact_object_address(m, v43+int32(228), v981)
	mBase = m.M
	v1170 = m.ExcPending
	if v1170 != 0 {
		goto L20
	} else {
		goto L263
	}
L262:
	;
	goto L258
L263:
	;
	v1172 = v1147 + int32(1)
	v1173 = *(*int32)(unsafe.Add(mBase, uint32(l24)+4))
	if v1172 < v1173 {
		v1147 = v1172
		goto L261
	} else {
		goto L264
	}
L264:
	;
	goto L262
L265:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43)+236)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v43)+232)) = l26
	*(*int32)(unsafe.Add(mBase, uint32(v43)+228)) = int32(1255)
	F_add_exact_object_address(m, v43+int32(228), v981)
	mBase = m.M
	v1223 = m.ExcPending
	if v1223 != 0 {
		goto L20
	} else {
		goto L268
	}
L266:
	;
	goto L267
L267:
	;
	F_record_object_address_dependencies(m, l0, v981, int32(110))
	mBase = m.M
	v1226 = m.ExcPending
	if v1226 != 0 {
		goto L20
	} else {
		goto L269
	}
L268:
	;
	goto L267
L269:
	;
	F_free_object_addresses(m, v981)
	mBase = m.M
	v1228 = m.ExcPending
	if v1228 != 0 {
		goto L20
	} else {
		goto L270
	}
L270:
	;
	if l7 != int32(14) {
		goto L271
	} else {
		goto L272
	}
L271:
	;
	if l22 != 0 {
		goto L275
	} else {
		goto L276
	}
L272:
	;
	if l11 == int32(0) {
		goto L271
	} else {
		goto L273
	}
L273:
	;
	F_recordDependencyOnExpr(m, l0, l11, int32(0))
	mBase = m.M
	v1235 = m.ExcPending
	if v1235 != 0 {
		goto L20
	} else {
		goto L274
	}
L274:
	;
	goto L271
L275:
	;
	F_recordDependencyOnExpr(m, l0, l22, int32(0))
	mBase = m.M
	v1238 = m.ExcPending
	if v1238 != 0 {
		goto L20
	} else {
		goto L278
	}
L276:
	;
	goto L277
L277:
	;
	if v465 == int32(0) {
		goto L279
	} else {
		goto L280
	}
L278:
	;
	goto L277
L279:
	;
	F_recordDependencyOnOwner(m, int32(1255), v972, l6)
	mBase = m.M
	v1243 = m.ExcPending
	if v1243 != 0 {
		goto L20
	} else {
		goto L282
	}
L280:
	;
	goto L281
L281:
	;
	F_recordDependencyOnCurrentExtension(m, l0, base.B2i32(v465 != int32(0)))
	mBase = m.M
	v1250 = m.ExcPending
	if v1250 != 0 {
		goto L20
	} else {
		goto L284
	}
L282:
	;
	F_recordDependencyOnNewAcl(m, int32(1255), v972, l6, v959)
	mBase = m.M
	v1246 = m.ExcPending
	if v1246 != 0 {
		goto L20
	} else {
		goto L283
	}
L283:
	;
	goto L281
L284:
	;
	F_pfree(m, v974)
	mBase = m.M
	v1252 = m.ExcPending
	if v1252 != 0 {
		goto L20
	} else {
		goto L285
	}
L285:
	;
	v1254 = *(*int32)(unsafe.Add(mBase, _consts[444]))
	if v1254 != 0 {
		goto L286
	} else {
		goto L287
	}
L286:
	;
	v1256 = int32(0)
	F_RunObjectPostCreateHook(m, int32(1255), v972, v1256, v1256)
	mBase = m.M
	v1259 = m.ExcPending
	if v1259 != 0 {
		goto L20
	} else {
		goto L289
	}
L287:
	;
	goto L288
L288:
	;
	F_sequence_close(m, v461, int32(3))
	mBase = m.M
	v1262 = m.ExcPending
	if v1262 != 0 {
		goto L20
	} else {
		goto L290
	}
L289:
	;
	goto L288
L290:
	;
	if l8 == int32(0) {
		goto L291
	} else {
		goto L292
	}
L291:
	;
	if v465 == int32(0) {
		goto L308
	} else {
		goto L309
	}
L292:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v1266 = m.ExcPending
	if v1266 != 0 {
		goto L20
	} else {
		goto L293
	}
L293:
	;
	if l25 != 0 {
		goto L295
	} else {
		goto L296
	}
L294:
	;
	v1275 = int32(4552168)
	v1277 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v1279 = v1277 + int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[5])) = v1279
	goto L300
L295:
	;
	v1268 = int32(*(*uint8)(unsafe.Add(mBase, _consts[449])))
	if v1268&int32(1) != 0 {
		goto L294
	} else {
		goto L298
	}
L296:
	;
	goto L297
L297:
	;
	v1272 = F_OidFunctionCall1Coll(m, l8, int32(0), v972)
	mBase = m.M
	v1273 = m.ExcPending
	if v1273 != 0 {
		goto L20
	} else {
		goto L299
	}
L298:
	;
	goto L297
L299:
	;
	goto L291
L300:
	;
	v1283 = F_superuser(m)
	mBase = m.M
	v1284 = m.ExcPending
	if v1284 != 0 {
		goto L20
	} else {
		goto L301
	}
L301:
	;
	if v1283 != 0 {
		goto L302
	} else {
		goto L303
	}
L302:
	;
	v1285 = int32(5)
	goto L304
L303:
	;
	v1285 = int32(6)
	goto L304
L304:
	;
	F_ProcessGUCArray(m, l25, v1285, int32(13), int32(2))
	mBase = m.M
	v1289 = m.ExcPending
	if v1289 != 0 {
		goto L20
	} else {
		goto L305
	}
L305:
	;
	v1291 = F_OidFunctionCall1Coll(m, l8, int32(0), v972)
	mBase = m.M
	v1292 = m.ExcPending
	if v1292 != 0 {
		goto L20
	} else {
		goto L306
	}
L306:
	;
	F_AtEOXact_GUC(m, int32(1), v1279)
	mBase = m.M
	v1295 = m.ExcPending
	if v1295 != 0 {
		goto L20
	} else {
		goto L307
	}
L307:
	;
	goto L291
L308:
	;
	v1301 = *(*int32)(unsafe.Add(mBase, _consts[130]))
	F_pgstat_create_transactional(m, int32(3), v1301, base.I64_extend_i32_u(v972))
	mBase = m.M
	v1304 = m.ExcPending
	if v1304 != 0 {
		goto L20
	} else {
		goto L311
	}
L309:
	;
	goto L310
L310:
	;
	m.G0 = v43 + int32(496)
	return
L311:
	;
	goto L310
L312:
	;
	F_errcode(m, int32(50884740))
	mBase = m.M
	v1314 = m.ExcPending
	if v1314 != 0 {
		goto L20
	} else {
		goto L313
	}
L313:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43)+144)) = l1
	F_errmsg(m, int32(171837), v43+int32(144))
	mBase = m.M
	v1320 = m.ExcPending
	if v1320 != 0 {
		goto L20
	} else {
		goto L314
	}
L314:
	;
	F_errfinish(m, int32(523705), int32(403), int32(372796))
	mBase = m.M
	v1325 = m.ExcPending
	if v1325 != 0 {
		goto L20
	} else {
		goto L315
	}
L315:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L316:
	;
	F_errcode(m, int32(50724996))
	mBase = m.M
	v1332 = m.ExcPending
	if v1332 != 0 {
		goto L20
	} else {
		goto L317
	}
L317:
	;
	if l12 == int32(112) {
		goto L318
	} else {
		goto L319
	}
L318:
	;
	v1335 = int32(140680)
	goto L320
L319:
	;
	v1335 = int32(265062)
	goto L320
L320:
	;
	F_errmsg(m, v1335, int32(0))
	mBase = m.M
	v1338 = m.ExcPending
	if v1338 != 0 {
		goto L20
	} else {
		goto L321
	}
L321:
	;
	v1339 = *(*int32)(unsafe.Add(mBase, uint32(v472)))
	v1340 = F_format_procedure(m, v1339)
	mBase = m.M
	v1341 = m.ExcPending
	if v1341 != 0 {
		goto L20
	} else {
		goto L322
	}
L322:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43)+116)) = v1340
	*(*int32)(unsafe.Add(mBase, uint32(v43)+112)) = v524
	F_errhint(m, int32(604974), v43+int32(112))
	mBase = m.M
	v1348 = m.ExcPending
	if v1348 != 0 {
		goto L20
	} else {
		goto L323
	}
L323:
	;
	F_errfinish(m, int32(523705), int32(449), int32(372796))
	mBase = m.M
	v1353 = m.ExcPending
	if v1353 != 0 {
		goto L20
	} else {
		goto L324
	}
L324:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L325:
	;
	F_errcode(m, int32(50724996))
	mBase = m.M
	v1360 = m.ExcPending
	if v1360 != 0 {
		goto L20
	} else {
		goto L326
	}
L326:
	;
	F_errmsg(m, int32(265062), int32(0))
	mBase = m.M
	v1364 = m.ExcPending
	if v1364 != 0 {
		goto L20
	} else {
		goto L327
	}
L327:
	;
	F_errdetail(m, int32(606883), int32(0))
	mBase = m.M
	v1368 = m.ExcPending
	if v1368 != 0 {
		goto L20
	} else {
		goto L328
	}
L328:
	;
	v1369 = *(*int32)(unsafe.Add(mBase, uint32(v472)))
	v1370 = F_format_procedure(m, v1369)
	mBase = m.M
	v1371 = m.ExcPending
	if v1371 != 0 {
		goto L20
	} else {
		goto L329
	}
L329:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43)+100)) = v1370
	*(*int32)(unsafe.Add(mBase, uint32(v43)+96)) = v524
	F_errhint(m, int32(604974), v43+int32(96))
	mBase = m.M
	v1378 = m.ExcPending
	if v1378 != 0 {
		goto L20
	} else {
		goto L330
	}
L330:
	;
	F_errfinish(m, int32(523705), int32(476), int32(372796))
	mBase = m.M
	v1383 = m.ExcPending
	if v1383 != 0 {
		goto L20
	} else {
		goto L331
	}
L331:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L332:
	;
	F_errcode(m, int32(50724996))
	mBase = m.M
	v1391 = m.ExcPending
	if v1391 != 0 {
		goto L20
	} else {
		goto L333
	}
L333:
	;
	v1392 = *(*int32)(unsafe.Add(mBase, uint32(v43)+228))
	v1396 = *(*int32)(unsafe.Add(mBase, uint32(v1392+v665<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v43)+80)) = v1396
	F_errmsg(m, int32(732880), v43+int32(80))
	mBase = m.M
	v1402 = m.ExcPending
	if v1402 != 0 {
		goto L20
	} else {
		goto L334
	}
L334:
	;
	v1403 = *(*int32)(unsafe.Add(mBase, uint32(v472)))
	v1404 = F_format_procedure(m, v1403)
	mBase = m.M
	v1405 = m.ExcPending
	if v1405 != 0 {
		goto L20
	} else {
		goto L335
	}
L335:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43)+68)) = v1404
	*(*int32)(unsafe.Add(mBase, uint32(v43)+64)) = v524
	F_errhint(m, int32(604974), v43-int32(-64))
	mBase = m.M
	v1412 = m.ExcPending
	if v1412 != 0 {
		goto L20
	} else {
		goto L336
	}
L336:
	;
	F_errfinish(m, int32(523705), int32(521), int32(372796))
	mBase = m.M
	v1417 = m.ExcPending
	if v1417 != 0 {
		goto L20
	} else {
		goto L337
	}
L337:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L338:
	;
	F_errcode(m, int32(50724996))
	mBase = m.M
	v1424 = m.ExcPending
	if v1424 != 0 {
		goto L20
	} else {
		goto L339
	}
L339:
	;
	F_errmsg(m, int32(265006), int32(0))
	mBase = m.M
	v1428 = m.ExcPending
	if v1428 != 0 {
		goto L20
	} else {
		goto L340
	}
L340:
	;
	v1429 = *(*int32)(unsafe.Add(mBase, uint32(v472)))
	v1430 = F_format_procedure(m, v1429)
	mBase = m.M
	v1431 = m.ExcPending
	if v1431 != 0 {
		goto L20
	} else {
		goto L341
	}
L341:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43)+36)) = v1430
	*(*int32)(unsafe.Add(mBase, uint32(v43)+32)) = v524
	F_errhint(m, int32(604974), v43+int32(32))
	mBase = m.M
	v1438 = m.ExcPending
	if v1438 != 0 {
		goto L20
	} else {
		goto L342
	}
L342:
	;
	F_errfinish(m, int32(523705), int32(547), int32(372796))
	mBase = m.M
	v1443 = m.ExcPending
	if v1443 != 0 {
		goto L20
	} else {
		goto L343
	}
L343:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L344:
	;
	F_errcode(m, int32(50724996))
	mBase = m.M
	v1450 = m.ExcPending
	if v1450 != 0 {
		goto L20
	} else {
		goto L345
	}
L345:
	;
	F_errmsg(m, int32(362247), int32(0))
	mBase = m.M
	v1454 = m.ExcPending
	if v1454 != 0 {
		goto L20
	} else {
		goto L346
	}
L346:
	;
	v1455 = *(*int32)(unsafe.Add(mBase, uint32(v472)))
	v1456 = F_format_procedure(m, v1455)
	mBase = m.M
	v1457 = m.ExcPending
	if v1457 != 0 {
		goto L20
	} else {
		goto L347
	}
L347:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43)+52)) = v1456
	*(*int32)(unsafe.Add(mBase, uint32(v43)+48)) = v524
	F_errhint(m, int32(604974), v43+int32(48))
	mBase = m.M
	v1464 = m.ExcPending
	if v1464 != 0 {
		goto L20
	} else {
		goto L348
	}
L348:
	;
	F_errfinish(m, int32(523705), int32(571), int32(372796))
	mBase = m.M
	v1469 = m.ExcPending
	if v1469 != 0 {
		goto L20
	} else {
		goto L349
	}
L349:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L350:
	;
	F_errcode(m, int32(50724996))
	mBase = m.M
	v1476 = m.ExcPending
	if v1476 != 0 {
		goto L20
	} else {
		goto L351
	}
L351:
	;
	F_errmsg(m, int32(761111), int32(0))
	mBase = m.M
	v1480 = m.ExcPending
	if v1480 != 0 {
		goto L20
	} else {
		goto L352
	}
L352:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43)+160)) = v143
	F_errdetail_internal(m, int32(216470), v43+int32(160))
	mBase = m.M
	v1486 = m.ExcPending
	if v1486 != 0 {
		goto L20
	} else {
		goto L353
	}
L353:
	;
	F_errfinish(m, int32(523705), int32(260), int32(372796))
	mBase = m.M
	v1491 = m.ExcPending
	if v1491 != 0 {
		goto L20
	} else {
		goto L354
	}
L354:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L355:
	;
	F_errcode(m, int32(50724996))
	mBase = m.M
	v1498 = m.ExcPending
	if v1498 != 0 {
		goto L20
	} else {
		goto L356
	}
L356:
	;
	F_errmsg(m, int32(388225), int32(0))
	mBase = m.M
	v1502 = m.ExcPending
	if v1502 != 0 {
		goto L20
	} else {
		goto L357
	}
L357:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43)+176)) = v140
	F_errdetail_internal(m, int32(216470), v43+int32(176))
	mBase = m.M
	v1508 = m.ExcPending
	if v1508 != 0 {
		goto L20
	} else {
		goto L358
	}
L358:
	;
	F_errfinish(m, int32(523705), int32(252), int32(372796))
	mBase = m.M
	v1513 = m.ExcPending
	if v1513 != 0 {
		goto L20
	} else {
		goto L359
	}
L359:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L360:
	;
	F_errcode(m, int32(50724996))
	mBase = m.M
	v1520 = m.ExcPending
	if v1520 != 0 {
		goto L20
	} else {
		goto L361
	}
L361:
	;
	F_errmsg(m, int32(761111), int32(0))
	mBase = m.M
	v1524 = m.ExcPending
	if v1524 != 0 {
		goto L20
	} else {
		goto L362
	}
L362:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43)+192)) = v79
	F_errdetail_internal(m, int32(216470), v43+int32(192))
	mBase = m.M
	v1530 = m.ExcPending
	if v1530 != 0 {
		goto L20
	} else {
		goto L363
	}
L363:
	;
	F_errfinish(m, int32(523705), int32(231), int32(372796))
	mBase = m.M
	v1535 = m.ExcPending
	if v1535 != 0 {
		goto L20
	} else {
		goto L364
	}
L364:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L365:
	;
	F_errcode(m, int32(50724996))
	mBase = m.M
	v1542 = m.ExcPending
	if v1542 != 0 {
		goto L20
	} else {
		goto L366
	}
L366:
	;
	F_errmsg(m, int32(388225), int32(0))
	mBase = m.M
	v1546 = m.ExcPending
	if v1546 != 0 {
		goto L20
	} else {
		goto L367
	}
L367:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43)+208)) = v75
	F_errdetail_internal(m, int32(216470), v43+int32(208))
	mBase = m.M
	v1552 = m.ExcPending
	if v1552 != 0 {
		goto L20
	} else {
		goto L368
	}
L368:
	;
	F_errfinish(m, int32(523705), int32(218), int32(372796))
	mBase = m.M
	v1557 = m.ExcPending
	if v1557 != 0 {
		goto L20
	} else {
		goto L369
	}
L369:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L370:
	;
	F_errmsg_internal(m, int32(25519), int32(0))
	mBase = m.M
	v1565 = m.ExcPending
	if v1565 != 0 {
		goto L20
	} else {
		goto L371
	}
L371:
	;
	F_errfinish(m, int32(523705), int32(203), int32(372796))
	mBase = m.M
	v1570 = m.ExcPending
	if v1570 != 0 {
		goto L20
	} else {
		goto L372
	}
L372:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L373:
	;
	F_errmsg_internal(m, int32(26351), int32(0))
	mBase = m.M
	v1580 = m.ExcPending
	if v1580 != 0 {
		goto L20
	} else {
		goto L374
	}
L374:
	;
	F_errfinish(m, int32(523705), int32(179), int32(372796))
	mBase = m.M
	v1585 = m.ExcPending
	if v1585 != 0 {
		goto L20
	} else {
		goto L375
	}
L375:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L376:
	;
	F_errcode(m, int32(50856197))
	mBase = m.M
	v1592 = m.ExcPending
	if v1592 != 0 {
		goto L20
	} else {
		goto L377
	}
L377:
	;
	v1593 = int32(100)
	*(*int32)(unsafe.Add(mBase, uint32(v43))) = v1593
	F_errmsg_plural(m, int32(100768), int32(130051), v1593, v43)
	mBase = m.M
	v1599 = m.ExcPending
	if v1599 != 0 {
		goto L20
	} else {
		goto L378
	}
L378:
	;
	F_errfinish(m, int32(523705), int32(161), int32(372796))
	mBase = m.M
	v1604 = m.ExcPending
	if v1604 != 0 {
		goto L20
	} else {
		goto L379
	}
L379:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_ProcessUtility(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	v11 = *(*int32)(unsafe.Add(mBase, _consts[911]))
	if v11 != 0 {
		m.T0[v11].(func(*base.Module, int32, int32, int32, int32, int32, int32, int32, int32))(m, l0, l1, l2, l3, l4, l5, l6, l7)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return
		} else {
			return
		}
	} else {
		F_standard_ProcessUtility(m, l0, l1, l2, l3, l4, l5, l6, l7)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return
		} else {
			return
		}
	}
}
func F_p_isEOF(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v6 != v7 {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
		v12 = base.B2i32(v9 == int32(0))
	} else {
		v12 = int32(1)
	}
	return v12
}
func F_p_isalpha(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	v3 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)))
	if v3 == int32(1) {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		if v6 != 0 {
			v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
			v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
			v12 = *(*int32)(unsafe.Add(mBase, uint32(v6+v8<<(uint(int32(2))%32))))
			return base.B2i32(base.Ui32(int32(127)) < base.Ui32(v12)) | base.B2i32(base.Ui32(v12|int32(32)-int32(97)) < base.Ui32(int32(26)))
		} else {
			v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
			v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
			v29 = *(*int32)(unsafe.Add(mBase, uint32(v23+v25<<(uint(int32(2))%32))))
			if base.Ui32(v29) <= base.Ui32(int32(131071)) {
				v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v29)>>(uint(int32(8))%32)))+uint32(_consts[914]))))
				v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v29)>>(uint(int32(3))%32))&int32(31)|v40<<(uint(int32(5))%32))+uint32(_consts[914]))))
				v54 = int32(base.Ui32(v46)>>(uint(v29&int32(7))%32)) & int32(1)
			} else {
				v54 = base.B2i32(base.Ui32(v29) < base.Ui32(int32(196606)))
			}
			return v54
		}
	} else {
		v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
		v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
		v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56+v58))))
		return base.B2i32(base.Ui32((v60|int32(32)-int32(97))&int32(255)) < base.Ui32(int32(26)))
	}
}
func F_p_ishost(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int64
	_ = v41
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v104 int32
	_ = v104
	v2 = int32(0)
	v6 = F_palloc0(m, int32(48))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v6)+20)) = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	*(*int32)(unsafe.Add(mBase, uint32(v6))) = v12 + v14
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v17 - v19
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)))
	*(*uint8)(unsafe.Add(mBase, uint32(v6)+16)) = uint8(v22)
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v24 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = v24 + v26<<(uint(int32(2))%32)
	goto L5
L4:
	;
	goto L5
L5:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v31 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v6)+8)) = v31 + v33<<(uint(int32(2))%32)
	goto L8
L7:
	;
	goto L8
L8:
	;
	v39 = F_palloc(m, int32(32))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	v41 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v39))) = v41
	*(*int64)(unsafe.Add(mBase, uint32(v39)+24)) = v41
	*(*int64)(unsafe.Add(mBase, uint32(v39)+16)) = v41
	*(*int64)(unsafe.Add(mBase, uint32(v39)+8)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v6)+24)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v39)+20)) = int32(0)
	v52 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v6)+29)) = uint8(v52)
	F_check_stack_depth(m)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	v56 = F_TParserGet(m, v6)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L1
	} else {
		goto L12
	}
L11:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v6)+24))
	if v90 != 0 {
		goto L15
	} else {
		goto L16
	}
L12:
	;
	if v56 == int32(0) {
		v89 = v2
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v6)+44))
	if v60 != int32(6) {
		v89 = v2
		goto L11
	} else {
		goto L14
	}
L14:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v6)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v63))) = v64 + v65
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)+4))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v6)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v68)+4)) = v69 + v70
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v73)+12))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v6)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v73)+12)) = v74 + v75
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v78)+16))
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v6)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v78)+16)) = v79 + v80
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v6)+24))
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v84)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v83)+8)) = v85
	v89 = int32(1)
	goto L11
L15:
	;
	v92 = v90
	goto L18
L16:
	;
	goto L17
L17:
	;
	F_pfree(m, v6)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L1
	} else {
		goto L22
	}
L18:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v92)+24))
	F_pfree(m, v92)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L1
	} else {
		goto L20
	}
L19:
	;
	goto L17
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6)+24)) = v95
	if v95 != 0 {
		v92 = v95
		goto L18
	} else {
		goto L21
	}
L21:
	;
	goto L19
L22:
	;
	return v89
}
func F_p_isignore(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
	return v2
}
func F_p_isspace(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v67 int32
	_ = v67
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v92 int32
	_ = v92
	v3 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)))
	if v3 == int32(1) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v92
L2:
	;
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v6 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)))
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76+v78))))
	v92 = base.B2i32(v80 == int32(32)) | base.B2i32(base.Ui32((v80-int32(9))&int32(255)) < base.Ui32(int32(5)))
	goto L1
L5:
	;
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v6+v9<<(uint(int32(2))%32))))
	if base.Ui32(int32(127)) < base.Ui32(v13) {
		v92 = int32(0)
		goto L1
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v24+v26<<(uint(int32(2))%32))))
	if v30 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	return base.B2i32(v13 == int32(32)) | base.B2i32(base.Ui32(v13-int32(9)) < base.Ui32(int32(5)))
L9:
	;
	return v74
L10:
	;
	v74 = int32(0)
	goto L9
L11:
	;
	goto L12
L12:
	;
	if v30 != 0 {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	v74 = base.B2i32(v67 != int32(0))
	goto L9
L14:
	;
	v40 = int32(4130160)
	goto L17
L15:
	;
	goto L16
L16:
	;
	v50 = int32(4130160)
	goto L26
L17:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
	if v42 != 0 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	if v42 != 0 {
		goto L23
	} else {
		goto L24
	}
L19:
	;
	if v30 != v42 {
		v40 = v40 + int32(4)
		goto L17
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	goto L18
L22:
	;
	goto L21
L23:
	;
	v48 = v40
	goto L25
L24:
	;
	v48 = int32(0)
	goto L25
L25:
	;
	v67 = v48
	goto L13
L26:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
	if v56 != 0 {
		v50 = v50 + int32(4)
		goto L26
	} else {
		goto L28
	}
L27:
	;
	v57 = int32(4130160)
	v67 = (v50-v57)&int32(-4) + v57
	goto L13
L28:
	;
	goto L27
}
func F_p_isstophost(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v3 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+29)))
	if v3 == int32(1) {
		v6 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+29)) = uint8(v6)
		v9 = int32(1)
	} else {
		v9 = int32(0)
	}
	return v9
}
func F_palloc(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	v2 = int32(0)
	v4 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	*(*uint8)(unsafe.Add(mBase, uint32(v4)+4)) = uint8(v2)
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v4)+12))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	v10 = m.T0[v9].(func(*base.Module, int32, int32, int32) int32)(m, v4, l0, v2)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		return v10
	}
}
func F_palloc0(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	v2 = int32(0)
	v4 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	*(*uint8)(unsafe.Add(mBase, uint32(v4)+4)) = uint8(v2)
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v4)+12))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	v10 = m.T0[v9].(func(*base.Module, int32, int32, int32) int32)(m, v4, l0, v2)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		if base.Ui32(int32(1024)) < base.Ui32(l0) {
			v37 = F__emscripten_memset_bulkmem(m, v10, base.I32_extend8_s(int32(0)), l0)
			mBase = m.M
			return v10
		} else {
			if l0&int32(3) != 0 {
				v37 = F__emscripten_memset_bulkmem(m, v10, base.I32_extend8_s(int32(0)), l0)
				mBase = m.M
				return v10
			} else {
				v18 = l0 + v10
				if base.Ui32(v18) <= base.Ui32(v10) {
					return v10
				} else {
					v24 = v10 + int32(4)
					if base.Ui32(v24) < base.Ui32(v18) {
						v26 = v18
					} else {
						v26 = v24
					}
					v33 = F__emscripten_memset_bulkmem(m, v10, base.I32_extend8_s(int32(0)), (v10^int32(-1)+v26)&int32(-4)+int32(4))
					mBase = m.M
					return v33
				}
			}
		}
	}
}
func F_paramlist_param_ref(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	v3 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v11 <= v3 {
		v50 = v3
		m.G0 = v9 + int32(16)
		return v50
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+28))
		if v15 < v11 {
			v50 = v3
			m.G0 = v9 + int32(16)
			return v50
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
			if v17 != 0 {
				v21 = m.T0[v17].(func(*base.Module, int32, int32, int32, int32) int32)(m, v14, v11, int32(0), v9+int32(4))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					v30 = v21
					v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+8))
					if v31 == int32(0) {
						v50 = v3
						m.G0 = v9 + int32(16)
						return v50
					} else {
						v35 = F_palloc0(m, int32(28))
						mBase = m.M
						v36 = m.ExcPending
						if v36 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v35)+8)) = v11
							*(*int64)(unsafe.Add(mBase, uint32(v35))) = int64(8)
							v40 = *(*int32)(unsafe.Add(mBase, uint32(v30)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v35)+16)) = int32(-1)
							*(*int32)(unsafe.Add(mBase, uint32(v35)+12)) = v40
							v44 = F_get_typcollation(m, v40)
							mBase = m.M
							v45 = m.ExcPending
							if v45 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v35)+20)) = v44
								v47 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
								*(*int32)(unsafe.Add(mBase, uint32(v35)+24)) = v47
								v50 = v35
								m.G0 = v9 + int32(16)
								return v50
							}
						}
					}
				}
			} else {
				v30 = v11*int32(12) + v14 + int32(20)
				v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+8))
				if v31 == int32(0) {
					v50 = v3
					m.G0 = v9 + int32(16)
					return v50
				} else {
					v35 = F_palloc0(m, int32(28))
					mBase = m.M
					v36 = m.ExcPending
					if v36 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v35)+8)) = v11
						*(*int64)(unsafe.Add(mBase, uint32(v35))) = int64(8)
						v40 = *(*int32)(unsafe.Add(mBase, uint32(v30)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v35)+16)) = int32(-1)
						*(*int32)(unsafe.Add(mBase, uint32(v35)+12)) = v40
						v44 = F_get_typcollation(m, v40)
						mBase = m.M
						v45 = m.ExcPending
						if v45 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v35)+20)) = v44
							v47 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v35)+24)) = v47
							v50 = v35
							m.G0 = v9 + int32(16)
							return v50
						}
					}
				}
			}
		}
	}
}
func F_parserOpenTable(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v103 int32
	_ = v103
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v155 int64
	_ = v155
	var v161 int32
	_ = v161
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v180 int32
	_ = v180
	v8 = m.G0
	v10 = v8 + int32(-64)
	m.G0 = v10
	v13 = v8 + int32(-20)
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = int32(489)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v14
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v13
	v20 = int32(4547032)
	v21 = *(*int32)(unsafe.Add(mBase, _consts[52]))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = v21
	*(*int32)(unsafe.Add(mBase, _consts[52])) = v8 + int32(-12)
	goto L1
L1:
	;
	v28 = F_table_openrv_extended(m, l1, l2, int32(1))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v167
	F_errdetail(m, int32(600224), v10)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L4
	} else {
		goto L47
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L4
	} else {
		goto L43
	}
L4:
	;
	return int32(0)
L5:
	;
	if v28 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v34 != 0 {
		goto L3
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v8+int32(-20))+8))
	*(*int32)(unsafe.Add(mBase, _consts[52])) = v142
	goto L42
L9:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if l0 != 0 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L4
	} else {
		goto L37
	}
L11:
	;
	v36 = l0
	goto L14
L12:
	;
	goto L13
L13:
	;
	v119 = int32(0)
	goto L10
L14:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v36)+40))
	if v43 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	goto L13
L16:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
	if v103 != 0 {
		v36 = v103
		goto L14
	} else {
		goto L36
	}
L17:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
	if v46 <= int32(0) {
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v49 = int32(0)
	if v49 < v46 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v52 = v46
	goto L21
L20:
	;
	v52 = v49
	goto L21
L21:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v43)+12))
	v57 = int32(0)
	goto L22
L22:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v53+v57<<(uint(int32(2))%32))))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)+4))
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35))))
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66))))
	if v70 == int32(0) {
		v89 = v69
		v90 = v70
		goto L25
	} else {
		goto L26
	}
L23:
	;
	v119 = int32(1)
	goto L10
L24:
	;
	if v90-v89 != 0 {
		goto L32
	} else {
		goto L33
	}
L25:
	;
	goto L24
L26:
	;
	if v69 != v70 {
		v89 = v69
		v90 = v70
		goto L25
	} else {
		goto L27
	}
L27:
	;
	v74 = v66
	v75 = v35
	goto L28
L28:
	;
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75)+1)))
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74)+1)))
	if v79 == int32(0) {
		v89 = v78
		v90 = v79
		goto L25
	} else {
		goto L30
	}
L29:
	;
	v89 = v78
	v90 = v79
	goto L25
L30:
	;
	v82 = int32(1)
	if v78 == v79 {
		v74 = v74 + v82
		v75 = v75 + v82
		goto L28
	} else {
		goto L31
	}
L31:
	;
	goto L29
L32:
	;
	v93 = v57 + int32(1)
	if v52 != v93 {
		v57 = v93
		goto L22
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	goto L23
L35:
	;
	goto L16
L36:
	;
	goto L15
L37:
	;
	F_errcode(m, int32(16908420))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L4
	} else {
		goto L38
	}
L38:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v127
	F_errmsg(m, int32(77524), v8+int32(-48))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L4
	} else {
		goto L39
	}
L39:
	;
	if v119 != 0 {
		goto L2
	} else {
		goto L40
	}
L40:
	;
	F_errfinish(m, int32(519399), int32(1469), int32(415280))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L4
	} else {
		goto L41
	}
L41:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L42:
	;
	m.G0 = v10 - int32(-64)
	return v28
L43:
	;
	F_errcode(m, int32(16908420))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L4
	} else {
		goto L44
	}
L44:
	;
	v155 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v10)+32)) = v155
	F_errmsg(m, int32(76293), v8+int32(-32))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L4
	} else {
		goto L45
	}
L45:
	;
	F_errfinish(m, int32(519399), int32(1448), int32(415280))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L4
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
	F_errhint(m, int32(628767), int32(0))
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L4
	} else {
		goto L48
	}
L48:
	;
	F_errfinish(m, int32(519399), int32(1464), int32(415280))
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L4
	} else {
		goto L49
	}
L49:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_per_MultiFuncCall(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(v2)+16))
	return v3
}
func F_performDeletion(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v13 = F_table_open(m, int32(2608), int32(3))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v13
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		switch v17 - int32(1259) {
		case 0:
			F_LockRelationOid(m, v16, int32(8))
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return
			} else {
				v31 = F_palloc(m, int32(16))
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return
				} else {
					*(*int64)(unsafe.Add(mBase, uint32(v31)+8)) = int64(137438953472)
					v36 = F_palloc(m, int32(384))
					mBase = m.M
					v37 = m.ExcPending
					if v37 != 0 {
						return
					} else {
						v38 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v31)+4)) = v38
						*(*int32)(unsafe.Add(mBase, uint32(v31))) = v36
						F_findDependentObjects(m, l0, int32(1), l2, v38, v31, v38, v9+int32(12))
						mBase = m.M
						v47 = m.ExcPending
						if v47 != 0 {
							return
						} else {
							F_reportDependentObjects(m, v31, l1, l2, l0)
							mBase = m.M
							v49 = m.ExcPending
							if v49 != 0 {
								return
							} else {
								F_deleteObjectsInList(m, v31, v9+int32(12), l2)
								mBase = m.M
								v53 = m.ExcPending
								if v53 != 0 {
									return
								} else {
									v54 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
									F_pfree(m, v54)
									mBase = m.M
									v56 = m.ExcPending
									if v56 != 0 {
										return
									} else {
										v57 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
										if v57 != 0 {
											F_pfree(m, v57)
											mBase = m.M
											v59 = m.ExcPending
											if v59 != 0 {
												return
											} else {
												F_pfree(m, v31)
												mBase = m.M
												v61 = m.ExcPending
												if v61 != 0 {
													return
												} else {
													v62 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
													F_sequence_close(m, v62, int32(3))
													mBase = m.M
													v65 = m.ExcPending
													if v65 != 0 {
														return
													} else {
														m.G0 = v9 + int32(16)
														return
													}
												}
											}
										} else {
											F_pfree(m, v31)
											mBase = m.M
											v61 = m.ExcPending
											if v61 != 0 {
												return
											} else {
												v62 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
												F_sequence_close(m, v62, int32(3))
												mBase = m.M
												v65 = m.ExcPending
												if v65 != 0 {
													return
												} else {
													m.G0 = v9 + int32(16)
													return
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		default:
			F_LockDatabaseObject(m, v17, v16, int32(8))
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return
			} else {
				v31 = F_palloc(m, int32(16))
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return
				} else {
					*(*int64)(unsafe.Add(mBase, uint32(v31)+8)) = int64(137438953472)
					v36 = F_palloc(m, int32(384))
					mBase = m.M
					v37 = m.ExcPending
					if v37 != 0 {
						return
					} else {
						v38 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v31)+4)) = v38
						*(*int32)(unsafe.Add(mBase, uint32(v31))) = v36
						F_findDependentObjects(m, l0, int32(1), l2, v38, v31, v38, v9+int32(12))
						mBase = m.M
						v47 = m.ExcPending
						if v47 != 0 {
							return
						} else {
							F_reportDependentObjects(m, v31, l1, l2, l0)
							mBase = m.M
							v49 = m.ExcPending
							if v49 != 0 {
								return
							} else {
								F_deleteObjectsInList(m, v31, v9+int32(12), l2)
								mBase = m.M
								v53 = m.ExcPending
								if v53 != 0 {
									return
								} else {
									v54 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
									F_pfree(m, v54)
									mBase = m.M
									v56 = m.ExcPending
									if v56 != 0 {
										return
									} else {
										v57 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
										if v57 != 0 {
											F_pfree(m, v57)
											mBase = m.M
											v59 = m.ExcPending
											if v59 != 0 {
												return
											} else {
												F_pfree(m, v31)
												mBase = m.M
												v61 = m.ExcPending
												if v61 != 0 {
													return
												} else {
													v62 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
													F_sequence_close(m, v62, int32(3))
													mBase = m.M
													v65 = m.ExcPending
													if v65 != 0 {
														return
													} else {
														m.G0 = v9 + int32(16)
														return
													}
												}
											}
										} else {
											F_pfree(m, v31)
											mBase = m.M
											v61 = m.ExcPending
											if v61 != 0 {
												return
											} else {
												v62 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
												F_sequence_close(m, v62, int32(3))
												mBase = m.M
												v65 = m.ExcPending
												if v65 != 0 {
													return
												} else {
													m.G0 = v9 + int32(16)
													return
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		case 2:
			F_LockSharedObject(m, int32(1261), v16, int32(8))
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return
			} else {
				v31 = F_palloc(m, int32(16))
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return
				} else {
					*(*int64)(unsafe.Add(mBase, uint32(v31)+8)) = int64(137438953472)
					v36 = F_palloc(m, int32(384))
					mBase = m.M
					v37 = m.ExcPending
					if v37 != 0 {
						return
					} else {
						v38 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v31)+4)) = v38
						*(*int32)(unsafe.Add(mBase, uint32(v31))) = v36
						F_findDependentObjects(m, l0, int32(1), l2, v38, v31, v38, v9+int32(12))
						mBase = m.M
						v47 = m.ExcPending
						if v47 != 0 {
							return
						} else {
							F_reportDependentObjects(m, v31, l1, l2, l0)
							mBase = m.M
							v49 = m.ExcPending
							if v49 != 0 {
								return
							} else {
								F_deleteObjectsInList(m, v31, v9+int32(12), l2)
								mBase = m.M
								v53 = m.ExcPending
								if v53 != 0 {
									return
								} else {
									v54 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
									F_pfree(m, v54)
									mBase = m.M
									v56 = m.ExcPending
									if v56 != 0 {
										return
									} else {
										v57 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
										if v57 != 0 {
											F_pfree(m, v57)
											mBase = m.M
											v59 = m.ExcPending
											if v59 != 0 {
												return
											} else {
												F_pfree(m, v31)
												mBase = m.M
												v61 = m.ExcPending
												if v61 != 0 {
													return
												} else {
													v62 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
													F_sequence_close(m, v62, int32(3))
													mBase = m.M
													v65 = m.ExcPending
													if v65 != 0 {
														return
													} else {
														m.G0 = v9 + int32(16)
														return
													}
												}
											}
										} else {
											F_pfree(m, v31)
											mBase = m.M
											v61 = m.ExcPending
											if v61 != 0 {
												return
											} else {
												v62 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
												F_sequence_close(m, v62, int32(3))
												mBase = m.M
												v65 = m.ExcPending
												if v65 != 0 {
													return
												} else {
													m.G0 = v9 + int32(16)
													return
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_pgstatginindex_v1_5(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = F_pgstatginindex_internal(m, v2, l0)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_pgstatindexbyid(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = F_superuser(m)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		if v4 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(16797828))
				mBase = m.M
				v16 = m.ExcPending
				if v16 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(149452), int32(0))
					mBase = m.M
					v22 = m.ExcPending
					if v22 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(514575), int32(193), int32(452045))
						mBase = m.M
						v29 = m.ExcPending
						if v29 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		} else {
			v31 = F_relation_open(m, v3, int32(1))
			mBase = m.M
			v32 = m.ExcPending
			if v32 != 0 {
				return int32(0)
			} else {
				v33 = F_pgstatindex_impl(m, v31, l0)
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return int32(0)
				} else {
					return v33
				}
			}
		}
	}
}
func F_pgstattuple(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = F_pg_detoast_datum_packed(m, v3)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		v8 = F_superuser(m)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			if v8 == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v15 = m.ExcPending
				if v15 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(16797828))
					mBase = m.M
					v18 = m.ExcPending
					if v18 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(149452), int32(0))
						mBase = m.M
						v24 = m.ExcPending
						if v24 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(522428), int32(178), int32(401022))
							mBase = m.M
							v31 = m.ExcPending
							if v31 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			} else {
				v32 = F_textToQualifiedNameList(m, v4)
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return int32(0)
				} else {
					v34 = F_makeRangeVarFromNameList(m, v32)
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return int32(0)
					} else {
						v37 = F_relation_openrv(m, v34, int32(1))
						mBase = m.M
						v38 = m.ExcPending
						if v38 != 0 {
							return int32(0)
						} else {
							v39 = F_pgstat_relation(m, v37, l0)
							mBase = m.M
							v40 = m.ExcPending
							if v40 != 0 {
								return int32(0)
							} else {
								return v39
							}
						}
					}
				}
			}
		}
	}
}
func F_pgstattuplebyid_v1_5(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = F_relation_open(m, v2, int32(1))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		v8 = F_pgstat_relation(m, v4, l0)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			return v8
		}
	}
}
func F_pkt_stream_flush(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v9 != 0 {
		v24 = int32(0)
		m.G0 = v7 + int32(16)
		return v24
	} else {
		v10 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v7)+8)) = uint8(v10)
		v15 = F_pushf_write(m, l0, v7+int32(8), int32(1))
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			if v15 < int32(0) {
				v24 = v15
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(1)
				v24 = int32(0)
			}
			m.G0 = v7 + int32(16)
			return v24
		}
	}
}
func F_plain_crypt_verify(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int64
	_ = v37
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v93 int32
	_ = v93
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v113 int32
	_ = v113
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v170 int32
	_ = v170
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v220 int32
	_ = v220
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v238 int32
	_ = v238
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v258 int32
	_ = v258
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v289 int32
	_ = v289
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v302 int32
	_ = v302
	var v310 int32
	_ = v310
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v335 int32
	_ = v335
	var v343 int32
	_ = v343
	var v359 int32
	_ = v359
	var v372 int32
	_ = v372
	var v375 int32
	_ = v375
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v385 int32
	_ = v385
	var v390 int32
	_ = v390
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v418 int32
	_ = v418
	var v420 int32
	_ = v420
	var v422 int32
	_ = v422
	var v424 int32
	_ = v424
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v443 int32
	_ = v443
	var v445 int32
	_ = v445
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v468 int32
	_ = v468
	var v473 int32
	_ = v473
	var v486 int32
	_ = v486
	var v491 int32
	_ = v491
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v506 int32
	_ = v506
	var v511 int32
	_ = v511
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v530 int32
	_ = v530
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v546 int32
	_ = v546
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v573 int32
	_ = v573
	var v581 int32
	_ = v581
	v5 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(144)
	m.G0 = v10
	*(*int32)(unsafe.Add(mBase, uint32(v10)+44)) = v5
	*(*int32)(unsafe.Add(mBase, uint32(v10)+132)) = v5
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v16 != int32(109) {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	m.G0 = v10 + int32(144)
	return v581
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v573
	v581 = int32(-1)
	goto L1
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = l0
	v566 = F_psprintf(m, int32(609482), v10+int32(32))
	mBase = m.M
	v567 = m.ExcPending
	if v567 != 0 {
		goto L32
	} else {
		goto L148
	}
L4:
	;
	v518 = F_strlen(m, l0)
	mBase = m.M
	v523 = F_pg_md5_encrypt(m, l2, l0, v518, v10+int32(48), v10+int32(44))
	mBase = m.M
	v524 = m.ExcPending
	if v524 != 0 {
		goto L32
	} else {
		goto L134
	}
L5:
	;
	v128 = F_parse_scram_secret(m, l1, v10+int32(136), v10+int32(128), v10+int32(132), v10+int32(140), v10+int32(48), v10+int32(96))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L32
	} else {
		goto L33
	}
L6:
	;
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
	if v19 != int32(100) {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+2)))
	if v22 != int32(53) {
		goto L5
	} else {
		goto L8
	}
L8:
	;
	v25 = F_strlen(m, l1)
	mBase = m.M
	if v25 != int32(35) {
		goto L5
	} else {
		goto L9
	}
L9:
	;
	v29 = l1 + int32(3)
	v30 = int32(355825)
	v34 = m.G0
	v36 = v34 - int32(32)
	v37 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v36)+24)) = v37
	*(*int64)(unsafe.Add(mBase, uint32(v36)+16)) = v37
	*(*int64)(unsafe.Add(mBase, uint32(v36)+8)) = v37
	*(*int64)(unsafe.Add(mBase, uint32(v36))) = v37
	v45 = int32(*(*uint8)(unsafe.Add(mBase, _consts[571])))
	if v45 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	if v113 == int32(32) {
		goto L4
	} else {
		goto L31
	}
L11:
	;
	v113 = int32(0)
	goto L10
L12:
	;
	goto L13
L13:
	;
	v49 = int32(*(*uint8)(unsafe.Add(mBase, _consts[572])))
	if v49 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v53 = v29
	goto L17
L15:
	;
	goto L16
L16:
	;
	v63 = v30
	v64 = v45
	goto L20
L17:
	;
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53))))
	if v59 == v45 {
		v53 = v53 + int32(1)
		goto L17
	} else {
		goto L19
	}
L18:
	;
	v113 = v53 - v29
	goto L10
L19:
	;
	goto L18
L20:
	;
	v71 = v36 + int32(base.Ui32(v64)>>(uint(int32(3))%32))&int32(28)
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v71)))
	v73 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v71))) = v72 | v73<<(uint(v64)%32)
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63)+1)))
	if v77 != 0 {
		v63 = v63 + v73
		v64 = v77
		goto L20
	} else {
		goto L22
	}
L21:
	;
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29))))
	if v80 == int32(0) {
		v105 = v29
		goto L23
	} else {
		goto L24
	}
L22:
	;
	goto L21
L23:
	;
	v113 = v105 - v29
	goto L10
L24:
	;
	v84 = v29
	v85 = v80
	goto L25
L25:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v36+int32(base.Ui32(v85)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v93)>>(uint(v85)%32))&int32(1) == int32(0) {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v105 = v101
	goto L23
L27:
	;
	v105 = v84
	goto L23
L28:
	;
	goto L29
L29:
	;
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84)+1)))
	v101 = v84 + int32(1)
	if v99 != 0 {
		v84 = v101
		v85 = v99
		goto L25
	} else {
		goto L30
	}
L30:
	;
	goto L26
L31:
	;
	goto L5
L32:
	;
	return int32(0)
L33:
	;
	if v128 == int32(0) {
		goto L3
	} else {
		goto L34
	}
L34:
	;
	v134 = int32(0)
	v135 = m.G0
	v137 = v135 - int32(192)
	m.G0 = v137
	*(*int32)(unsafe.Add(mBase, uint32(v137)+180)) = v134
	*(*int32)(unsafe.Add(mBase, uint32(v137)+40)) = v134
	v155 = F_parse_scram_secret(m, l1, v137+int32(184), v137+int32(176), v137+int32(180), v137+int32(188), v137+int32(112), v137+int32(80))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L32
	} else {
		goto L38
	}
L35:
	;
	if v491 != 0 {
		v581 = v134
		goto L1
	} else {
		goto L132
	}
L36:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v499 = m.ExcPending
	if v499 != 0 {
		goto L32
	} else {
		goto L129
	}
L37:
	;
	m.G0 = v137 + int32(192)
	goto L35
L38:
	;
	if v155 == int32(0) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v161 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L32
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v137)+188))
	v177 = F_strlen(m, v176)
	mBase = m.M
	v181 = v177 * int32(3) >> (uint(int32(2)) % 32)
	goto L46
L42:
	;
	if v161 == int32(0) {
		v491 = v5
		goto L37
	} else {
		goto L43
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v137)+32)) = l0
	F_errmsg(m, int32(733335), v137+int32(32))
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L32
	} else {
		goto L44
	}
L44:
	;
	F_errfinish(m, int32(520686), int32(547), int32(439650))
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L32
	} else {
		goto L45
	}
L45:
	;
	v491 = v5
	goto L37
L46:
	;
	v182 = F_palloc(m, v181)
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L32
	} else {
		goto L47
	}
L47:
	;
	v184 = F_strlen(m, v176)
	mBase = m.M
	v185 = int32(0)
	v192 = v176 + v184
	if base.Ui32(v176) < base.Ui32(v192) {
		goto L50
	} else {
		goto L51
	}
L48:
	;
	if v372 < int32(0) {
		goto L92
	} else {
		goto L93
	}
L49:
	;
	v359 = F___memset(m, v182, int32(0), v181)
	mBase = m.M
	v372 = int32(-1)
	goto L48
L50:
	;
	v194 = v176
	v198 = v185
	v199 = v182
	v200 = v185
	v202 = v185
	goto L53
L51:
	;
	v343 = v182
	goto L52
L52:
	;
	v372 = v343 - v182
	goto L48
L53:
	;
	v206 = v194 + int32(1)
	v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v194))))
	if v207 != int32(61) {
		goto L61
	} else {
		goto L62
	}
L54:
	;
	if v331 != 0 {
		goto L49
	} else {
		goto L91
	}
L55:
	;
	if v329 != v192 {
		v194 = v329
		v198 = v331
		v199 = v332
		v200 = v333
		v202 = v335
		goto L53
	} else {
		goto L90
	}
L56:
	;
	if v181 < v199-v182+int32(1) {
		goto L49
	} else {
		goto L77
	}
L57:
	;
	v285 = int32(2)
	v286 = v206
	v289 = v202 << (uint(int32(6)) % 32)
	goto L56
L58:
	;
	v274 = v271 + v270<<(uint(int32(6))%32)
	v276 = v267 + int32(1)
	if v276 == int32(4) {
		v285 = v268
		v286 = v269
		v289 = v274
		goto L56
	} else {
		goto L76
	}
L59:
	;
	v267 = int32(3)
	v268 = int32(1)
	v269 = v194 + int32(2)
	v270 = v225
	v271 = v220
	goto L58
L60:
	;
	if base.Ui32(int32(125)) < base.Ui32((v244-int32(1))&int32(255)) {
		goto L49
	} else {
		goto L74
	}
L61:
	;
	v211 = v207 - int32(9)
	if base.Ui32(int32(23)) < base.Ui32(v211) {
		v244 = v207
		v245 = v198
		v246 = v200
		v247 = v206
		v248 = v202
		goto L60
	} else {
		goto L64
	}
L62:
	;
	goto L63
L63:
	;
	v220 = int32(0)
	if v200 != 0 {
		v267 = v198
		v268 = v200
		v269 = v206
		v270 = v202
		v271 = v220
		goto L58
	} else {
		goto L66
	}
L64:
	;
	if int32(1)<<(uint(v211)%32)&int32(8388627) == int32(0) {
		v244 = v207
		v245 = v198
		v246 = v200
		v247 = v206
		v248 = v202
		goto L60
	} else {
		goto L65
	}
L65:
	;
	goto L49
L66:
	;
	switch v198 - int32(2) {
	case 0:
		goto L67
	case 1:
		goto L57
	default:
		goto L49
	}
L67:
	;
	if v206 == v192 {
		goto L49
	} else {
		goto L68
	}
L68:
	;
	v225 = v202 << (uint(int32(6)) % 32)
	v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206))))
	if v226 == int32(61) {
		goto L59
	} else {
		goto L69
	}
L69:
	;
	v230 = v226 - int32(9)
	if int32(1)<<(uint(v230)%32)&int32(8388627) != 0 {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v238 = base.B2i32(base.Ui32(v230) <= base.Ui32(int32(23)))
	goto L72
L71:
	;
	v238 = int32(0)
	goto L72
L72:
	;
	if v238 != 0 {
		goto L49
	} else {
		goto L73
	}
L73:
	;
	v244 = v226
	v245 = int32(3)
	v246 = int32(1)
	v247 = v194 + int32(2)
	v248 = v225
	goto L60
L74:
	;
	v258 = int32(*(*int8)(unsafe.Add(mBase, uint32(v244)+uint32(_consts[573]))))
	if v258 < int32(0) {
		goto L49
	} else {
		goto L75
	}
L75:
	;
	v267 = v245
	v268 = v246
	v269 = v247
	v270 = v248
	v271 = v258
	goto L58
L76:
	;
	v329 = v269
	v331 = v276
	v332 = v199
	v333 = v268
	v335 = v274
	goto L55
L77:
	;
	v295 = int32(base.Ui32(v289) >> (uint(int32(16)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v199))) = uint8(v295)
	v298 = v199 + int32(1)
	if base.Ui32(v285) < base.Ui32(int32(2)) {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v302 = v285
	goto L80
L79:
	;
	v302 = int32(0)
	goto L80
L80:
	;
	if v302 == int32(0) {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	if v181 < v298-v182+int32(1) {
		goto L49
	} else {
		goto L84
	}
L82:
	;
	v314 = v298
	goto L83
L83:
	;
	v315 = int32(0)
	if v285 == v315 {
		goto L86
	} else {
		goto L87
	}
L84:
	;
	v310 = int32(base.Ui32(v289) >> (uint(int32(8)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v199)+1)) = uint8(v310)
	v314 = v199 + int32(2)
	goto L83
L85:
	;
	v329 = v286
	v331 = int32(0)
	v332 = v327
	v333 = v285
	v335 = v315
	goto L55
L86:
	;
	if v181 < v314-v182+int32(1) {
		goto L49
	} else {
		goto L89
	}
L87:
	;
	if base.Ui32(int32(3)) <= base.Ui32(v285) {
		goto L86
	} else {
		goto L88
	}
L88:
	;
	v327 = v314
	goto L85
L89:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v314))) = uint8(v289)
	v327 = v314 + int32(1)
	goto L85
L90:
	;
	goto L54
L91:
	;
	v343 = v332
	goto L52
L92:
	;
	v375 = int32(0)
	v378 = F_errstart(m, int32(15), v375)
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L32
	} else {
		goto L95
	}
L93:
	;
	goto L94
L94:
	;
	v393 = F_pg_saslprep(m, l2, v137+int32(44))
	mBase = m.M
	v394 = m.ExcPending
	if v394 != 0 {
		goto L32
	} else {
		goto L99
	}
L95:
	;
	if v378 == int32(0) {
		v491 = v375
		goto L37
	} else {
		goto L96
	}
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v137))) = l0
	F_errmsg(m, int32(733335), v137)
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L32
	} else {
		goto L97
	}
L97:
	;
	F_errfinish(m, int32(520686), int32(558), int32(439650))
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L32
	} else {
		goto L98
	}
L98:
	;
	v491 = v375
	goto L37
L99:
	;
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v137)+44))
	if v393 != 0 {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	v396 = l2
	goto L102
L101:
	;
	v396 = v395
	goto L102
L102:
	;
	v397 = *(*int32)(unsafe.Add(mBase, uint32(v137)+176))
	v398 = *(*int32)(unsafe.Add(mBase, uint32(v137)+180))
	v399 = *(*int32)(unsafe.Add(mBase, uint32(v137)+184))
	v404 = F_scram_SaltedPassword(m, v396, v397, v398, v182, v372, v399, v137+int32(144), v137+int32(40))
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
		goto L32
	} else {
		goto L103
	}
L103:
	;
	if v404 < int32(0) {
		goto L36
	} else {
		goto L104
	}
L104:
	;
	v414 = F_scram_ServerKey(m, v137+int32(144), v397, v398, v137+int32(48), v137+int32(40))
	mBase = m.M
	v415 = m.ExcPending
	if v415 != 0 {
		goto L32
	} else {
		goto L105
	}
L105:
	;
	if v414 < int32(0) {
		goto L36
	} else {
		goto L106
	}
L106:
	;
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v137)+44))
	if v418 != 0 {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	F_pfree(m, v418)
	mBase = m.M
	v420 = m.ExcPending
	if v420 != 0 {
		goto L32
	} else {
		goto L110
	}
L108:
	;
	goto L109
L109:
	;
	v422 = v137 + int32(48)
	v424 = v137 + int32(80)
	if base.Ui32(int32(4)) <= base.Ui32(v398) {
		goto L114
	} else {
		goto L115
	}
L110:
	;
	goto L109
L111:
	;
	v491 = base.B2i32(v486 == int32(0))
	goto L37
L112:
	;
	v486 = int32(0)
	goto L111
L113:
	;
	v460 = v455
	v461 = v456
	v462 = v457
	goto L123
L114:
	;
	if (v422|v424)&int32(3) != 0 {
		v455 = v422
		v456 = v424
		v457 = v398
		goto L113
	} else {
		goto L117
	}
L115:
	;
	v448 = v422
	v449 = v424
	v450 = v398
	goto L116
L116:
	;
	if v450 == int32(0) {
		goto L112
	} else {
		goto L122
	}
L117:
	;
	v432 = v422
	v433 = v424
	v434 = v398
	goto L118
L118:
	;
	v437 = *(*int32)(unsafe.Add(mBase, uint32(v432)))
	v438 = *(*int32)(unsafe.Add(mBase, uint32(v433)))
	if v437 != v438 {
		v455 = v432
		v456 = v433
		v457 = v434
		goto L113
	} else {
		goto L120
	}
L119:
	;
	v448 = v443
	v449 = v441
	v450 = v445
	goto L116
L120:
	;
	v440 = int32(4)
	v441 = v433 + v440
	v443 = v432 + v440
	v445 = v434 - v440
	if base.Ui32(int32(3)) < base.Ui32(v445) {
		v432 = v443
		v433 = v441
		v434 = v445
		goto L118
	} else {
		goto L121
	}
L121:
	;
	goto L119
L122:
	;
	v455 = v448
	v456 = v449
	v457 = v450
	goto L113
L123:
	;
	v465 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v460))))
	v466 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v461))))
	if v465 == v466 {
		goto L125
	} else {
		goto L126
	}
L124:
	;
	v486 = v465 - v466
	goto L111
L125:
	;
	v468 = int32(1)
	v473 = v462 - v468
	if v473 != 0 {
		v460 = v460 + v468
		v461 = v461 + v468
		v462 = v473
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
	goto L112
L129:
	;
	v500 = *(*int32)(unsafe.Add(mBase, uint32(v137)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v137)+16)) = v500
	F_errmsg_internal(m, int32(209297), v137+int32(16))
	mBase = m.M
	v506 = m.ExcPending
	if v506 != 0 {
		goto L32
	} else {
		goto L130
	}
L130:
	;
	F_errfinish(m, int32(520686), int32(574), int32(439650))
	mBase = m.M
	v511 = m.ExcPending
	if v511 != 0 {
		goto L32
	} else {
		goto L131
	}
L131:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L132:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = l0
	v516 = F_psprintf(m, int32(694668), v10+int32(16))
	mBase = m.M
	v517 = m.ExcPending
	if v517 != 0 {
		goto L32
	} else {
		goto L133
	}
L133:
	;
	v573 = v516
	goto L2
L134:
	;
	if v523 == int32(0) {
		goto L135
	} else {
		goto L136
	}
L135:
	;
	v527 = *(*int32)(unsafe.Add(mBase, uint32(v10)+44))
	v573 = v527
	goto L2
L136:
	;
	goto L137
L137:
	;
	v528 = int32(0)
	v530 = v10 + int32(48)
	v533 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	v534 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v530))))
	if v534 == v528 {
		v553 = v533
		v554 = v534
		goto L139
	} else {
		goto L140
	}
L138:
	;
	if v554-v553 == int32(0) {
		v581 = v528
		goto L1
	} else {
		goto L146
	}
L139:
	;
	goto L138
L140:
	;
	if v533 != v534 {
		v553 = v533
		v554 = v534
		goto L139
	} else {
		goto L141
	}
L141:
	;
	v538 = v530
	v539 = l1
	goto L142
L142:
	;
	v542 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v539)+1)))
	v543 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v538)+1)))
	if v543 == int32(0) {
		v553 = v542
		v554 = v543
		goto L139
	} else {
		goto L144
	}
L143:
	;
	v553 = v542
	v554 = v543
	goto L139
L144:
	;
	v546 = int32(1)
	if v542 == v543 {
		v538 = v538 + v546
		v539 = v539 + v546
		goto L142
	} else {
		goto L145
	}
L145:
	;
	goto L143
L146:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = l0
	v560 = F_psprintf(m, int32(694668), v10)
	mBase = m.M
	v561 = m.ExcPending
	if v561 != 0 {
		goto L32
	} else {
		goto L147
	}
L147:
	;
	v573 = v560
	goto L2
L148:
	;
	v573 = v566
	goto L2
}
func F_plainto_tsquery(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = F_pg_detoast_datum_packed(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v9 = F_getTSCurrentConfig(m)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			v11 = F_DirectFunctionCall2Coll(m, int32(1176), int32(0), v9, v3)
			mBase = m.M
			v12 = m.ExcPending
			if v12 != 0 {
				return int32(0)
			} else {
				return v11
			}
		}
	}
}
func F_plan_recursive_revoke(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	v12 = l2 << (uint(int32(2)) % 32)
	v13 = l1 + v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	switch v14 - int32(1) {
	case 0:
		goto L7
	default:
		goto L6
	case 3:
		goto L2
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L26
	} else {
		goto L29
	}
L2:
	;
	return
L3:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v47 <= int32(0) {
		goto L2
	} else {
		goto L12
	}
L4:
	;
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+16)))
	if v39 != int32(1) {
		goto L2
	} else {
		goto L11
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = int32(4)
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+16)))
	if v38 != 0 {
		v45 = v34
		v46 = v35
		goto L3
	} else {
		goto L10
	}
L6:
	;
	v27 = l0 + int32(48)
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v27+v12)))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+56))
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+22)))
	v32 = v30 + v31
	if l3 != 0 {
		goto L4
	} else {
		goto L9
	}
L7:
	;
	if l3 != 0 {
		goto L2
	} else {
		goto L8
	}
L8:
	;
	v18 = l0 + int32(48)
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v18+l2<<(uint(int32(2))%32))))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+56))
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+22)))
	v34 = v23 + v24
	v35 = v18
	goto L5
L9:
	;
	v34 = v32
	v35 = v27
	goto L5
L10:
	;
	goto L2
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = int32(1)
	v45 = v32
	v46 = v27
	goto L3
L12:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v45)+8))
	v54 = int32(0)
	goto L13
L13:
	;
	v63 = v54 << (uint(int32(2)) % 32)
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v46+v63)))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)+56))
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+22)))
	v68 = v66 + v67
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)+8))
	if v69 != v50 {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v84 = int32(0)
	v85 = v47
	goto L20
L15:
	;
	v79 = v54 + int32(1)
	if v79 != v47 {
		v54 = v79
		goto L13
	} else {
		goto L19
	}
L16:
	;
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68)+16)))
	if v71 != int32(1) {
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l1+v63)))
	if v75 == int32(0) {
		goto L2
	} else {
		goto L18
	}
L18:
	;
	goto L15
L19:
	;
	goto L14
L20:
	;
	v93 = v84 << (uint(int32(2)) % 32)
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v46+v93)))
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v95)+56))
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96)+22)))
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v96+v97)+12))
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v45)+8))
	if v99 != v100 {
		v112 = v85
		goto L22
	} else {
		goto L23
	}
L21:
	;
	goto L2
L22:
	;
	v114 = v84 + int32(1)
	if v114 < v112 {
		v84 = v114
		v85 = v112
		goto L20
	} else {
		goto L28
	}
L23:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(l1+v93)))
	if v103 == int32(4) {
		v112 = v85
		goto L22
	} else {
		goto L24
	}
L24:
	;
	if l4 == int32(0) {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	F_plan_recursive_revoke(m, l0, l1, v84, int32(0), l4)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	return
L27:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v112 = v111
	goto L22
L28:
	;
	goto L21
L29:
	;
	F_errcode(m, int32(16909442))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L26
	} else {
		goto L30
	}
L30:
	;
	F_errmsg(m, int32(78928), int32(0))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L26
	} else {
		goto L31
	}
L31:
	;
	F_errhint(m, int32(639537), int32(0))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L26
	} else {
		goto L32
	}
L32:
	;
	F_errfinish(m, int32(518278), int32(2494), int32(416640))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L26
	} else {
		goto L33
	}
L33:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pop_arg_long_double(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v11 int64
	_ = v11
	var v12 int64
	_ = v12
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v25 int64
	_ = v25
	var v29 int64
	_ = v29
	var v30 int32
	_ = v30
	var v39 int64
	_ = v39
	var v44 int64
	_ = v44
	var v54 int64
	_ = v54
	var v57 int32
	_ = v57
	var v58 int64
	_ = v58
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int64
	_ = v86
	var v92 int64
	_ = v92
	var v97 int64
	_ = v97
	var v100 int64
	_ = v100
	var v103 int64
	_ = v103
	var v105 int64
	_ = v105
	var v106 int64
	_ = v106
	var v114 int64
	_ = v114
	var v124 int64
	_ = v124
	var v128 int32
	_ = v128
	var v129 int64
	_ = v129
	var v131 int64
	_ = v131
	var v138 int64
	_ = v138
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v7 = (v3 + int32(7)) & int32(-8)
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v7 + int32(16)
	v11 = *(*int64)(unsafe.Add(mBase, uint32(v7)))
	v12 = *(*int64)(unsafe.Add(mBase, uint32(v7)+8))
	v20 = m.G0
	v22 = v20 - int32(32)
	m.G0 = v22
	v25 = v12 & int64(281474976710655)
	v29 = int64(base.Ui64(v12)>>(uint(int64(48))%64)) & int64(32767)
	v30 = base.I32_wrap_i64(v29)
	if base.Ui32(v30-int32(15361)) <= base.Ui32(int32(2045)) {
		v39 = v25<<(uint(int64(4))%64) | int64(base.Ui64(v11)>>(uint(int64(60))%64))
		v44 = v11 & int64(1152921504606846975)
		if base.Ui64(int64(576460752303423489)) <= base.Ui64(v44) {
			v54 = v39 + int64(1)
		} else {
			if v44 != int64(576460752303423488) {
				v54 = v39
			} else {
				v54 = v39&int64(1) + v39
			}
		}
		v57 = base.B2i32(base.Ui64(int64(4503599627370495)) < base.Ui64(v54))
		if base.Ui64(int64(4503599627370495)) < base.Ui64(v54) {
			v58 = int64(0)
		} else {
			v58 = v54
		}
		v131 = v58
		v138 = base.I64_extend_i32_u(v57) + base.I64_extend_i32_u(v30-int32(15360))
	} else {
		if v11|v25 == int64(0) {
			if base.Ui32(int32(17406)) < base.Ui32(v30) {
				v131 = int64(0)
				v138 = int64(2047)
			} else {
				v81 = base.B2i32(v29 == int64(0))
				if v29 == int64(0) {
					v82 = int32(15360)
				} else {
					v82 = int32(15361)
				}
				v83 = v82 - v30
				if int32(112) < v83 {
					v86 = int64(0)
					v131 = v86
					v138 = v86
				} else {
					if v29 == int64(0) {
						v92 = v25
					} else {
						v92 = v25 | int64(281474976710656)
					}
					F___ashlti3(m, v22+int32(16), v11, v92, int32(128)-v83)
					mBase = m.M
					F___lshrti3(m, v22, v11, v92, v83)
					mBase = m.M
					v97 = *(*int64)(unsafe.Add(mBase, uint32(v22)+8))
					v100 = *(*int64)(unsafe.Add(mBase, uint32(v22)))
					v103 = v97<<(uint(int64(4))%64) | int64(base.Ui64(v100)>>(uint(int64(60))%64))
					v105 = *(*int64)(unsafe.Add(mBase, uint32(v22)+16))
					v106 = *(*int64)(unsafe.Add(mBase, uint32(v22)+24))
					v114 = base.I64_extend_i32_u(base.B2i32(v30 != v82)&base.B2i32(v105|v106 != int64(0))) | v100&int64(1152921504606846975)
					if base.Ui64(int64(576460752303423489)) <= base.Ui64(v114) {
						v124 = v103 + int64(1)
					} else {
						if v114 != int64(576460752303423488) {
							v124 = v103
						} else {
							v124 = v103&int64(1) + v103
						}
					}
					v128 = base.B2i32(base.Ui64(int64(4503599627370495)) < base.Ui64(v124))
					if base.Ui64(int64(4503599627370495)) < base.Ui64(v124) {
						v129 = v124 ^ int64(4503599627370496)
					} else {
						v129 = v124
					}
					v131 = v129
					v138 = base.I64_extend_i32_u(v128)
				}
			}
		} else {
			if v29 != int64(32767) {
				if base.Ui32(int32(17406)) < base.Ui32(v30) {
					v131 = int64(0)
					v138 = int64(2047)
				} else {
					v81 = base.B2i32(v29 == int64(0))
					if v29 == int64(0) {
						v82 = int32(15360)
					} else {
						v82 = int32(15361)
					}
					v83 = v82 - v30
					if int32(112) < v83 {
						v86 = int64(0)
						v131 = v86
						v138 = v86
					} else {
						if v29 == int64(0) {
							v92 = v25
						} else {
							v92 = v25 | int64(281474976710656)
						}
						F___ashlti3(m, v22+int32(16), v11, v92, int32(128)-v83)
						mBase = m.M
						F___lshrti3(m, v22, v11, v92, v83)
						mBase = m.M
						v97 = *(*int64)(unsafe.Add(mBase, uint32(v22)+8))
						v100 = *(*int64)(unsafe.Add(mBase, uint32(v22)))
						v103 = v97<<(uint(int64(4))%64) | int64(base.Ui64(v100)>>(uint(int64(60))%64))
						v105 = *(*int64)(unsafe.Add(mBase, uint32(v22)+16))
						v106 = *(*int64)(unsafe.Add(mBase, uint32(v22)+24))
						v114 = base.I64_extend_i32_u(base.B2i32(v30 != v82)&base.B2i32(v105|v106 != int64(0))) | v100&int64(1152921504606846975)
						if base.Ui64(int64(576460752303423489)) <= base.Ui64(v114) {
							v124 = v103 + int64(1)
						} else {
							if v114 != int64(576460752303423488) {
								v124 = v103
							} else {
								v124 = v103&int64(1) + v103
							}
						}
						v128 = base.B2i32(base.Ui64(int64(4503599627370495)) < base.Ui64(v124))
						if base.Ui64(int64(4503599627370495)) < base.Ui64(v124) {
							v129 = v124 ^ int64(4503599627370496)
						} else {
							v129 = v124
						}
						v131 = v129
						v138 = base.I64_extend_i32_u(v128)
					}
				}
			} else {
				v131 = v25<<(uint(int64(4))%64) | int64(base.Ui64(v11)>>(uint(int64(60))%64)) | int64(2251799813685248)
				v138 = int64(2047)
			}
		}
	}
	m.G0 = v22 + int32(32)
	*(*float64)(unsafe.Add(mBase, uint32(l0))) = base.F64_reinterpret_i64(v12&int64(-9223372036854775807-1) | v138<<(uint(int64(52))%64) | v131)
	return
}
func F_positionsel(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_Float8GetDatum(m, float64(0.1))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_postgresql_fdw_validator(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
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
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v322 int32
	_ = v322
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v338 int32
	_ = v338
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v350 int32
	_ = v350
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v366 int32
	_ = v366
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v378 int32
	_ = v378
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v394 int32
	_ = v394
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v407 int32
	_ = v407
	var v419 int32
	_ = v419
	var v424 int32
	_ = v424
	var v429 int32
	_ = v429
	var v434 int32
	_ = v434
	var v439 int32
	_ = v439
	var v444 int32
	_ = v444
	var v449 int32
	_ = v449
	var v454 int32
	_ = v454
	var v459 int32
	_ = v459
	var v464 int32
	_ = v464
	var v469 int32
	_ = v469
	var v474 int32
	_ = v474
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v496 int32
	_ = v496
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v508 int32
	_ = v508
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v524 int32
	_ = v524
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v537 int32
	_ = v537
	var v549 int32
	_ = v549
	var v554 int32
	_ = v554
	var v557 int32
	_ = v557
	var v572 int32
	_ = v572
	var v581 int32
	_ = v581
	var v584 int32
	_ = v584
	var v588 int32
	_ = v588
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v598 int32
	_ = v598
	var v604 int32
	_ = v604
	var v608 int32
	_ = v608
	var v613 int32
	_ = v613
	v8 = m.G0
	v10 = v8 - int32(48)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = F_untransformRelOptions(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	v584 = *(*int32)(unsafe.Add(mBase, uint32(v10+int32(32))+12))
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v588 = m.ExcPending
	if v588 != 0 {
		goto L4
	} else {
		goto L171
	}
L2:
	;
	v572 = v10 + int32(32)
	*(*int32)(unsafe.Add(mBase, uint32(v572)+12)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v572)+8)) = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v572)+4)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v572))) = v41
	goto L170
L3:
	;
	m.G0 = v10 + int32(48)
	return int32(1)
L4:
	;
	return int32(0)
L5:
	;
	if v13 == int32(0) {
		goto L3
	} else {
		goto L6
	}
L6:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if v19 <= int32(0) {
		goto L3
	} else {
		goto L7
	}
L7:
	;
	v22 = int32(0)
	if v22 < v19 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v25 = v19
	goto L10
L9:
	;
	v25 = v22
	goto L10
L10:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v32 = int32(0)
	goto L11
L11:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v26+v32<<(uint(int32(2))%32))))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+8))
	switch v27 - int32(1417) {
	case 0:
		goto L16
	case 1:
		goto L15
	default:
		goto L2
	}
L12:
	;
	goto L3
L13:
	;
	v557 = v32 + int32(1)
	if v557 != v25 {
		v32 = v557
		goto L11
	} else {
		goto L169
	}
L14:
	;
	v581 = int32(1)
	goto L1
L15:
	;
	v480 = int32(228489)
	v483 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41))))
	v484 = int32(*(*uint8)(unsafe.Add(mBase, _consts[542])))
	if v484 == int32(0) {
		v503 = v483
		v504 = v484
		goto L149
	} else {
		goto L150
	}
L16:
	;
	v42 = int32(383377)
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41))))
	v46 = int32(*(*uint8)(unsafe.Add(mBase, _consts[543])))
	if v46 == int32(0) {
		v65 = v45
		v66 = v46
		goto L18
	} else {
		goto L19
	}
L17:
	;
	if v66-v65 == int32(0) {
		goto L13
	} else {
		goto L25
	}
L18:
	;
	goto L17
L19:
	;
	if v45 != v46 {
		v65 = v45
		v66 = v46
		goto L18
	} else {
		goto L20
	}
L20:
	;
	v50 = v42
	v51 = v41
	goto L21
L21:
	;
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51)+1)))
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50)+1)))
	if v55 == int32(0) {
		v65 = v54
		v66 = v55
		goto L18
	} else {
		goto L23
	}
L22:
	;
	v65 = v54
	v66 = v55
	goto L18
L23:
	;
	v58 = int32(1)
	if v54 == v55 {
		v50 = v50 + v58
		v51 = v51 + v58
		goto L21
	} else {
		goto L24
	}
L24:
	;
	goto L22
L25:
	;
	v70 = int32(437650)
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41))))
	v74 = int32(*(*uint8)(unsafe.Add(mBase, _consts[544])))
	if v74 == int32(0) {
		v93 = v73
		v94 = v74
		goto L27
	} else {
		goto L28
	}
L26:
	;
	if v94-v93 == int32(0) {
		goto L13
	} else {
		goto L34
	}
L27:
	;
	goto L26
L28:
	;
	if v73 != v74 {
		v93 = v73
		v94 = v74
		goto L27
	} else {
		goto L29
	}
L29:
	;
	v78 = v70
	v79 = v41
	goto L30
L30:
	;
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79)+1)))
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78)+1)))
	if v83 == int32(0) {
		v93 = v82
		v94 = v83
		goto L27
	} else {
		goto L32
	}
L31:
	;
	v93 = v82
	v94 = v83
	goto L27
L32:
	;
	v86 = int32(1)
	if v82 == v83 {
		v78 = v78 + v86
		v79 = v79 + v86
		goto L30
	} else {
		goto L33
	}
L33:
	;
	goto L31
L34:
	;
	v98 = int32(71320)
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41))))
	v102 = int32(*(*uint8)(unsafe.Add(mBase, _consts[545])))
	if v102 == int32(0) {
		v121 = v101
		v122 = v102
		goto L36
	} else {
		goto L37
	}
L35:
	;
	if v122-v121 == int32(0) {
		goto L13
	} else {
		goto L43
	}
L36:
	;
	goto L35
L37:
	;
	if v101 != v102 {
		v121 = v101
		v122 = v102
		goto L36
	} else {
		goto L38
	}
L38:
	;
	v106 = v98
	v107 = v41
	goto L39
L39:
	;
	v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+1)))
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106)+1)))
	if v111 == int32(0) {
		v121 = v110
		v122 = v111
		goto L36
	} else {
		goto L41
	}
L40:
	;
	v121 = v110
	v122 = v111
	goto L36
L41:
	;
	v114 = int32(1)
	if v110 == v111 {
		v106 = v106 + v114
		v107 = v107 + v114
		goto L39
	} else {
		goto L42
	}
L42:
	;
	goto L40
L43:
	;
	v126 = int32(395777)
	v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41))))
	v130 = int32(*(*uint8)(unsafe.Add(mBase, _consts[546])))
	if v130 == int32(0) {
		v149 = v129
		v150 = v130
		goto L45
	} else {
		goto L46
	}
L44:
	;
	if v150-v149 == int32(0) {
		goto L13
	} else {
		goto L52
	}
L45:
	;
	goto L44
L46:
	;
	if v129 != v130 {
		v149 = v129
		v150 = v130
		goto L45
	} else {
		goto L47
	}
L47:
	;
	v134 = v126
	v135 = v41
	goto L48
L48:
	;
	v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135)+1)))
	v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v134)+1)))
	if v139 == int32(0) {
		v149 = v138
		v150 = v139
		goto L45
	} else {
		goto L50
	}
L49:
	;
	v149 = v138
	v150 = v139
	goto L45
L50:
	;
	v142 = int32(1)
	if v138 == v139 {
		v134 = v134 + v142
		v135 = v135 + v142
		goto L48
	} else {
		goto L51
	}
L51:
	;
	goto L49
L52:
	;
	v154 = int32(74060)
	v157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41))))
	v158 = int32(*(*uint8)(unsafe.Add(mBase, _consts[547])))
	if v158 == int32(0) {
		v177 = v157
		v178 = v158
		goto L54
	} else {
		goto L55
	}
L53:
	;
	if v178-v177 == int32(0) {
		goto L13
	} else {
		goto L61
	}
L54:
	;
	goto L53
L55:
	;
	if v157 != v158 {
		v177 = v157
		v178 = v158
		goto L54
	} else {
		goto L56
	}
L56:
	;
	v162 = v154
	v163 = v41
	goto L57
L57:
	;
	v166 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v163)+1)))
	v167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162)+1)))
	if v167 == int32(0) {
		v177 = v166
		v178 = v167
		goto L54
	} else {
		goto L59
	}
L58:
	;
	v177 = v166
	v178 = v167
	goto L54
L59:
	;
	v170 = int32(1)
	if v166 == v167 {
		v162 = v162 + v170
		v163 = v163 + v170
		goto L57
	} else {
		goto L60
	}
L60:
	;
	goto L58
L61:
	;
	v182 = int32(240240)
	v185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41))))
	v186 = int32(*(*uint8)(unsafe.Add(mBase, _consts[548])))
	if v186 == int32(0) {
		v205 = v185
		v206 = v186
		goto L63
	} else {
		goto L64
	}
L62:
	;
	if v206-v205 == int32(0) {
		goto L13
	} else {
		goto L70
	}
L63:
	;
	goto L62
L64:
	;
	if v185 != v186 {
		v205 = v185
		v206 = v186
		goto L63
	} else {
		goto L65
	}
L65:
	;
	v190 = v182
	v191 = v41
	goto L66
L66:
	;
	v194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v191)+1)))
	v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v190)+1)))
	if v195 == int32(0) {
		v205 = v194
		v206 = v195
		goto L63
	} else {
		goto L68
	}
L67:
	;
	v205 = v194
	v206 = v195
	goto L63
L68:
	;
	v198 = int32(1)
	if v194 == v195 {
		v190 = v190 + v198
		v191 = v191 + v198
		goto L66
	} else {
		goto L69
	}
L69:
	;
	goto L67
L70:
	;
	v210 = int32(87096)
	v213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41))))
	v214 = int32(*(*uint8)(unsafe.Add(mBase, _consts[549])))
	if v214 == int32(0) {
		v233 = v213
		v234 = v214
		goto L72
	} else {
		goto L73
	}
L71:
	;
	if v234-v233 == int32(0) {
		goto L13
	} else {
		goto L79
	}
L72:
	;
	goto L71
L73:
	;
	if v213 != v214 {
		v233 = v213
		v234 = v214
		goto L72
	} else {
		goto L74
	}
L74:
	;
	v218 = v210
	v219 = v41
	goto L75
L75:
	;
	v222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v219)+1)))
	v223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v218)+1)))
	if v223 == int32(0) {
		v233 = v222
		v234 = v223
		goto L72
	} else {
		goto L77
	}
L76:
	;
	v233 = v222
	v234 = v223
	goto L72
L77:
	;
	v226 = int32(1)
	if v222 == v223 {
		v218 = v218 + v226
		v219 = v219 + v226
		goto L75
	} else {
		goto L78
	}
L78:
	;
	goto L76
L79:
	;
	v238 = int32(8336)
	v241 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41))))
	v242 = int32(*(*uint8)(unsafe.Add(mBase, _consts[35])))
	if v242 == int32(0) {
		v261 = v241
		v262 = v242
		goto L81
	} else {
		goto L82
	}
L80:
	;
	if v262-v261 == int32(0) {
		goto L13
	} else {
		goto L88
	}
L81:
	;
	goto L80
L82:
	;
	if v241 != v242 {
		v261 = v241
		v262 = v242
		goto L81
	} else {
		goto L83
	}
L83:
	;
	v246 = v238
	v247 = v41
	goto L84
L84:
	;
	v250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v247)+1)))
	v251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v246)+1)))
	if v251 == int32(0) {
		v261 = v250
		v262 = v251
		goto L81
	} else {
		goto L86
	}
L85:
	;
	v261 = v250
	v262 = v251
	goto L81
L86:
	;
	v254 = int32(1)
	if v250 == v251 {
		v246 = v246 + v254
		v247 = v247 + v254
		goto L84
	} else {
		goto L87
	}
L87:
	;
	goto L85
L88:
	;
	v266 = int32(146572)
	v269 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41))))
	v270 = int32(*(*uint8)(unsafe.Add(mBase, _consts[550])))
	if v270 == int32(0) {
		v289 = v269
		v290 = v270
		goto L90
	} else {
		goto L91
	}
L89:
	;
	if v290-v289 == int32(0) {
		goto L13
	} else {
		goto L97
	}
L90:
	;
	goto L89
L91:
	;
	if v269 != v270 {
		v289 = v269
		v290 = v270
		goto L90
	} else {
		goto L92
	}
L92:
	;
	v274 = v266
	v275 = v41
	goto L93
L93:
	;
	v278 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v275)+1)))
	v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v274)+1)))
	if v279 == int32(0) {
		v289 = v278
		v290 = v279
		goto L90
	} else {
		goto L95
	}
L94:
	;
	v289 = v278
	v290 = v279
	goto L90
L95:
	;
	v282 = int32(1)
	if v278 == v279 {
		v274 = v274 + v282
		v275 = v275 + v282
		goto L93
	} else {
		goto L96
	}
L96:
	;
	goto L94
L97:
	;
	v294 = int32(313944)
	v297 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41))))
	v298 = int32(*(*uint8)(unsafe.Add(mBase, _consts[551])))
	if v298 == int32(0) {
		v317 = v297
		v318 = v298
		goto L99
	} else {
		goto L100
	}
L98:
	;
	if v318-v317 == int32(0) {
		goto L13
	} else {
		goto L106
	}
L99:
	;
	goto L98
L100:
	;
	if v297 != v298 {
		v317 = v297
		v318 = v298
		goto L99
	} else {
		goto L101
	}
L101:
	;
	v302 = v294
	v303 = v41
	goto L102
L102:
	;
	v306 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v303)+1)))
	v307 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v302)+1)))
	if v307 == int32(0) {
		v317 = v306
		v318 = v307
		goto L99
	} else {
		goto L104
	}
L103:
	;
	v317 = v306
	v318 = v307
	goto L99
L104:
	;
	v310 = int32(1)
	if v306 == v307 {
		v302 = v302 + v310
		v303 = v303 + v310
		goto L102
	} else {
		goto L105
	}
L105:
	;
	goto L103
L106:
	;
	v322 = int32(430542)
	v325 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41))))
	v326 = int32(*(*uint8)(unsafe.Add(mBase, _consts[552])))
	if v326 == int32(0) {
		v345 = v325
		v346 = v326
		goto L108
	} else {
		goto L109
	}
L107:
	;
	if v346-v345 == int32(0) {
		goto L13
	} else {
		goto L115
	}
L108:
	;
	goto L107
L109:
	;
	if v325 != v326 {
		v345 = v325
		v346 = v326
		goto L108
	} else {
		goto L110
	}
L110:
	;
	v330 = v322
	v331 = v41
	goto L111
L111:
	;
	v334 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v331)+1)))
	v335 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v330)+1)))
	if v335 == int32(0) {
		v345 = v334
		v346 = v335
		goto L108
	} else {
		goto L113
	}
L112:
	;
	v345 = v334
	v346 = v335
	goto L108
L113:
	;
	v338 = int32(1)
	if v334 == v335 {
		v330 = v330 + v338
		v331 = v331 + v338
		goto L111
	} else {
		goto L114
	}
L114:
	;
	goto L112
L115:
	;
	v350 = int32(527189)
	v353 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41))))
	v354 = int32(*(*uint8)(unsafe.Add(mBase, _consts[553])))
	if v354 == int32(0) {
		v373 = v353
		v374 = v354
		goto L117
	} else {
		goto L118
	}
L116:
	;
	if v374-v373 == int32(0) {
		goto L13
	} else {
		goto L124
	}
L117:
	;
	goto L116
L118:
	;
	if v353 != v354 {
		v373 = v353
		v374 = v354
		goto L117
	} else {
		goto L119
	}
L119:
	;
	v358 = v350
	v359 = v41
	goto L120
L120:
	;
	v362 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v359)+1)))
	v363 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v358)+1)))
	if v363 == int32(0) {
		v373 = v362
		v374 = v363
		goto L117
	} else {
		goto L122
	}
L121:
	;
	v373 = v362
	v374 = v363
	goto L117
L122:
	;
	v366 = int32(1)
	if v362 == v363 {
		v358 = v358 + v366
		v359 = v359 + v366
		goto L120
	} else {
		goto L123
	}
L123:
	;
	goto L121
L124:
	;
	v378 = int32(277001)
	v381 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41))))
	v382 = int32(*(*uint8)(unsafe.Add(mBase, _consts[554])))
	if v382 == int32(0) {
		v401 = v381
		v402 = v382
		goto L126
	} else {
		goto L127
	}
L125:
	;
	if v402-v401 == int32(0) {
		goto L13
	} else {
		goto L133
	}
L126:
	;
	goto L125
L127:
	;
	if v381 != v382 {
		v401 = v381
		v402 = v382
		goto L126
	} else {
		goto L128
	}
L128:
	;
	v386 = v378
	v387 = v41
	goto L129
L129:
	;
	v390 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v387)+1)))
	v391 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v386)+1)))
	if v391 == int32(0) {
		v401 = v390
		v402 = v391
		goto L126
	} else {
		goto L131
	}
L130:
	;
	v401 = v390
	v402 = v391
	goto L126
L131:
	;
	v394 = int32(1)
	if v390 == v391 {
		v386 = v386 + v394
		v387 = v387 + v394
		goto L129
	} else {
		goto L132
	}
L132:
	;
	goto L130
L133:
	;
	v407 = v10 + int32(32)
	*(*int32)(unsafe.Add(mBase, uint32(v407)+12)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v407)+8)) = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v407)+4)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v407))) = v41
	goto L134
L134:
	;
	F_updateClosestMatch(m, v10+int32(32), int32(383377))
	mBase = m.M
	v419 = m.ExcPending
	if v419 != 0 {
		goto L4
	} else {
		goto L135
	}
L135:
	;
	F_updateClosestMatch(m, v10+int32(32), int32(437650))
	mBase = m.M
	v424 = m.ExcPending
	if v424 != 0 {
		goto L4
	} else {
		goto L136
	}
L136:
	;
	F_updateClosestMatch(m, v10+int32(32), int32(71320))
	mBase = m.M
	v429 = m.ExcPending
	if v429 != 0 {
		goto L4
	} else {
		goto L137
	}
L137:
	;
	F_updateClosestMatch(m, v10+int32(32), int32(395777))
	mBase = m.M
	v434 = m.ExcPending
	if v434 != 0 {
		goto L4
	} else {
		goto L138
	}
L138:
	;
	F_updateClosestMatch(m, v10+int32(32), int32(74060))
	mBase = m.M
	v439 = m.ExcPending
	if v439 != 0 {
		goto L4
	} else {
		goto L139
	}
L139:
	;
	F_updateClosestMatch(m, v10+int32(32), int32(240240))
	mBase = m.M
	v444 = m.ExcPending
	if v444 != 0 {
		goto L4
	} else {
		goto L140
	}
L140:
	;
	F_updateClosestMatch(m, v10+int32(32), int32(87096))
	mBase = m.M
	v449 = m.ExcPending
	if v449 != 0 {
		goto L4
	} else {
		goto L141
	}
L141:
	;
	F_updateClosestMatch(m, v10+int32(32), int32(8336))
	mBase = m.M
	v454 = m.ExcPending
	if v454 != 0 {
		goto L4
	} else {
		goto L142
	}
L142:
	;
	F_updateClosestMatch(m, v10+int32(32), int32(146572))
	mBase = m.M
	v459 = m.ExcPending
	if v459 != 0 {
		goto L4
	} else {
		goto L143
	}
L143:
	;
	F_updateClosestMatch(m, v10+int32(32), int32(313944))
	mBase = m.M
	v464 = m.ExcPending
	if v464 != 0 {
		goto L4
	} else {
		goto L144
	}
L144:
	;
	F_updateClosestMatch(m, v10+int32(32), int32(430542))
	mBase = m.M
	v469 = m.ExcPending
	if v469 != 0 {
		goto L4
	} else {
		goto L145
	}
L145:
	;
	F_updateClosestMatch(m, v10+int32(32), int32(527189))
	mBase = m.M
	v474 = m.ExcPending
	if v474 != 0 {
		goto L4
	} else {
		goto L146
	}
L146:
	;
	F_updateClosestMatch(m, v10+int32(32), int32(277001))
	mBase = m.M
	v479 = m.ExcPending
	if v479 != 0 {
		goto L4
	} else {
		goto L147
	}
L147:
	;
	goto L14
L148:
	;
	if v504-v503 == int32(0) {
		goto L13
	} else {
		goto L156
	}
L149:
	;
	goto L148
L150:
	;
	if v483 != v484 {
		v503 = v483
		v504 = v484
		goto L149
	} else {
		goto L151
	}
L151:
	;
	v488 = v480
	v489 = v41
	goto L152
L152:
	;
	v492 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v489)+1)))
	v493 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v488)+1)))
	if v493 == int32(0) {
		v503 = v492
		v504 = v493
		goto L149
	} else {
		goto L154
	}
L153:
	;
	v503 = v492
	v504 = v493
	goto L149
L154:
	;
	v496 = int32(1)
	if v492 == v493 {
		v488 = v488 + v496
		v489 = v489 + v496
		goto L152
	} else {
		goto L155
	}
L155:
	;
	goto L153
L156:
	;
	v508 = int32(439861)
	v511 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41))))
	v512 = int32(*(*uint8)(unsafe.Add(mBase, _consts[555])))
	if v512 == int32(0) {
		v531 = v511
		v532 = v512
		goto L158
	} else {
		goto L159
	}
L157:
	;
	if v532-v531 == int32(0) {
		goto L13
	} else {
		goto L165
	}
L158:
	;
	goto L157
L159:
	;
	if v511 != v512 {
		v531 = v511
		v532 = v512
		goto L158
	} else {
		goto L160
	}
L160:
	;
	v516 = v508
	v517 = v41
	goto L161
L161:
	;
	v520 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v517)+1)))
	v521 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516)+1)))
	if v521 == int32(0) {
		v531 = v520
		v532 = v521
		goto L158
	} else {
		goto L163
	}
L162:
	;
	v531 = v520
	v532 = v521
	goto L158
L163:
	;
	v524 = int32(1)
	if v520 == v521 {
		v516 = v516 + v524
		v517 = v517 + v524
		goto L161
	} else {
		goto L164
	}
L164:
	;
	goto L162
L165:
	;
	v537 = v10 + int32(32)
	*(*int32)(unsafe.Add(mBase, uint32(v537)+12)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v537)+8)) = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v537)+4)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v537))) = v41
	goto L166
L166:
	;
	F_updateClosestMatch(m, v10+int32(32), int32(228489))
	mBase = m.M
	v549 = m.ExcPending
	if v549 != 0 {
		goto L4
	} else {
		goto L167
	}
L167:
	;
	F_updateClosestMatch(m, v10+int32(32), int32(439861))
	mBase = m.M
	v554 = m.ExcPending
	if v554 != 0 {
		goto L4
	} else {
		goto L168
	}
L168:
	;
	goto L14
L169:
	;
	goto L12
L170:
	;
	v581 = int32(0)
	goto L1
L171:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v591 = m.ExcPending
	if v591 != 0 {
		goto L4
	} else {
		goto L172
	}
L172:
	;
	v592 = *(*int32)(unsafe.Add(mBase, uint32(v40)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v592
	F_errmsg(m, int32(735030), v10+int32(16))
	mBase = m.M
	v598 = m.ExcPending
	if v598 != 0 {
		goto L4
	} else {
		goto L173
	}
L173:
	;
	if v581 != 0 {
		goto L175
	} else {
		goto L176
	}
L174:
	;
	F_errfinish(m, int32(519766), int32(665), int32(220252))
	mBase = m.M
	v613 = m.ExcPending
	if v613 != 0 {
		goto L4
	} else {
		goto L181
	}
L175:
	;
	if v584 == int32(0) {
		goto L174
	} else {
		goto L178
	}
L176:
	;
	goto L177
L177:
	;
	F_errhint(m, int32(604653), int32(0))
	mBase = m.M
	v608 = m.ExcPending
	if v608 != 0 {
		goto L4
	} else {
		goto L180
	}
L178:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v584
	F_errhint(m, int32(694839), v10)
	mBase = m.M
	v604 = m.ExcPending
	if v604 != 0 {
		goto L4
	} else {
		goto L179
	}
L179:
	;
	goto L174
L180:
	;
	goto L174
L181:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pre_sync_fname(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int64
	_ = v22
	var v24 int64
	_ = v24
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	v5 = m.G0
	v7 = v5 - int32(48)
	m.G0 = v7
	if l1 != 0 {
		m.G0 = v7 + int32(48)
		return
	} else {
		v16 = m.G0
		v18 = v16 - int32(16)
		m.G0 = v18
		v21 = *(*int32)(unsafe.Add(mBase, _consts[304]))
		if v21 != 0 {
			v22 = F_GetCurrentTimestamp(m)
			mBase = m.M
			v24 = *(*int64)(unsafe.Add(mBase, _consts[305]))
			F_TimestampDifference(m, v24, v22, v18+int32(12), v18+int32(8))
			mBase = m.M
			v30 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
			*(*int32)(unsafe.Add(mBase, uint32(v7+int32(44)))) = v30
			v32 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v7+int32(40)))) = v32
			*(*int32)(unsafe.Add(mBase, _consts[304])) = int32(0)
		} else {
		}
		m.G0 = v18 + int32(16)
		if base.B2i32(v21 != int32(0)) == int32(0) {
			v70 = *(*int32)(unsafe.Add(mBase, _consts[423]))
			v71 = F_OpenTransientFilePerm(m, l0, int32(0), v70)
			mBase = m.M
			v72 = m.ExcPending
			if v72 != 0 {
				return
			} else {
				if v71 < int32(0) {
					v76 = *(*int32)(unsafe.Add(mBase, _consts[43]))
					if v76 == int32(2) {
						m.G0 = v7 + int32(48)
						return
					} else {
						v80 = F_errstart(m, l2, int32(0))
						mBase = m.M
						v81 = m.ExcPending
						if v81 != 0 {
							return
						} else {
							if v80 == int32(0) {
								m.G0 = v7 + int32(48)
								return
							} else {
								v98 = int32(312236)
								v99 = int32(3810)
								F_errcode_for_file_access(m)
								mBase = m.M
								v101 = m.ExcPending
								if v101 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
									F_errmsg(m, v98, v7)
									mBase = m.M
									v104 = m.ExcPending
									if v104 != 0 {
										return
									} else {
										F_errfinish(m, int32(523542), v99, int32(395499))
										mBase = m.M
										v108 = m.ExcPending
										if v108 != 0 {
											return
										} else {
											m.G0 = v7 + int32(48)
											return
										}
									}
								}
							}
						}
					}
				} else {
					v86 = F_fsync(m, v71)
					mBase = m.M
					v87 = F_CloseTransientFile(m, v71)
					mBase = m.M
					v88 = m.ExcPending
					if v88 != 0 {
						return
					} else {
						if v87 == int32(0) {
							m.G0 = v7 + int32(48)
							return
						} else {
							v92 = F_errstart(m, l2, int32(0))
							mBase = m.M
							v93 = m.ExcPending
							if v93 != 0 {
								return
							} else {
								if v92 == int32(0) {
									m.G0 = v7 + int32(48)
									return
								} else {
									v98 = int32(312979)
									v99 = int32(3823)
									F_errcode_for_file_access(m)
									mBase = m.M
									v101 = m.ExcPending
									if v101 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
										F_errmsg(m, v98, v7)
										mBase = m.M
										v104 = m.ExcPending
										if v104 != 0 {
											return
										} else {
											F_errfinish(m, int32(523542), v99, int32(395499))
											mBase = m.M
											v108 = m.ExcPending
											if v108 != 0 {
												return
											} else {
												m.G0 = v7 + int32(48)
												return
											}
										}
									}
								}
							}
						}
					}
				}
			}
		} else {
			v47 = F_errstart(m, int32(15), int32(0))
			mBase = m.M
			v48 = m.ExcPending
			if v48 != 0 {
				return
			} else {
				if v47 == int32(0) {
					v70 = *(*int32)(unsafe.Add(mBase, _consts[423]))
					v71 = F_OpenTransientFilePerm(m, l0, int32(0), v70)
					mBase = m.M
					v72 = m.ExcPending
					if v72 != 0 {
						return
					} else {
						if v71 < int32(0) {
							v76 = *(*int32)(unsafe.Add(mBase, _consts[43]))
							if v76 == int32(2) {
								m.G0 = v7 + int32(48)
								return
							} else {
								v80 = F_errstart(m, l2, int32(0))
								mBase = m.M
								v81 = m.ExcPending
								if v81 != 0 {
									return
								} else {
									if v80 == int32(0) {
										m.G0 = v7 + int32(48)
										return
									} else {
										v98 = int32(312236)
										v99 = int32(3810)
										F_errcode_for_file_access(m)
										mBase = m.M
										v101 = m.ExcPending
										if v101 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
											F_errmsg(m, v98, v7)
											mBase = m.M
											v104 = m.ExcPending
											if v104 != 0 {
												return
											} else {
												F_errfinish(m, int32(523542), v99, int32(395499))
												mBase = m.M
												v108 = m.ExcPending
												if v108 != 0 {
													return
												} else {
													m.G0 = v7 + int32(48)
													return
												}
											}
										}
									}
								}
							}
						} else {
							v86 = F_fsync(m, v71)
							mBase = m.M
							v87 = F_CloseTransientFile(m, v71)
							mBase = m.M
							v88 = m.ExcPending
							if v88 != 0 {
								return
							} else {
								if v87 == int32(0) {
									m.G0 = v7 + int32(48)
									return
								} else {
									v92 = F_errstart(m, l2, int32(0))
									mBase = m.M
									v93 = m.ExcPending
									if v93 != 0 {
										return
									} else {
										if v92 == int32(0) {
											m.G0 = v7 + int32(48)
											return
										} else {
											v98 = int32(312979)
											v99 = int32(3823)
											F_errcode_for_file_access(m)
											mBase = m.M
											v101 = m.ExcPending
											if v101 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
												F_errmsg(m, v98, v7)
												mBase = m.M
												v104 = m.ExcPending
												if v104 != 0 {
													return
												} else {
													F_errfinish(m, int32(523542), v99, int32(395499))
													mBase = m.M
													v108 = m.ExcPending
													if v108 != 0 {
														return
													} else {
														m.G0 = v7 + int32(48)
														return
													}
												}
											}
										}
									}
								}
							}
						}
					}
				} else {
					v51 = *(*int32)(unsafe.Add(mBase, uint32(v7)+44))
					*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v51
					*(*int32)(unsafe.Add(mBase, uint32(v7)+24)) = l0
					v54 = *(*int32)(unsafe.Add(mBase, uint32(v7)+40))
					v56 = base.I32_div_s(v54, int32(10000))
					*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = v56
					F_errmsg(m, int32(212838), v7+int32(16))
					mBase = m.M
					v62 = m.ExcPending
					if v62 != 0 {
						return
					} else {
						F_errfinish(m, int32(523542), int32(3800), int32(395499))
						mBase = m.M
						v67 = m.ExcPending
						if v67 != 0 {
							return
						} else {
							v70 = *(*int32)(unsafe.Add(mBase, _consts[423]))
							v71 = F_OpenTransientFilePerm(m, l0, int32(0), v70)
							mBase = m.M
							v72 = m.ExcPending
							if v72 != 0 {
								return
							} else {
								if v71 < int32(0) {
									v76 = *(*int32)(unsafe.Add(mBase, _consts[43]))
									if v76 == int32(2) {
										m.G0 = v7 + int32(48)
										return
									} else {
										v80 = F_errstart(m, l2, int32(0))
										mBase = m.M
										v81 = m.ExcPending
										if v81 != 0 {
											return
										} else {
											if v80 == int32(0) {
												m.G0 = v7 + int32(48)
												return
											} else {
												v98 = int32(312236)
												v99 = int32(3810)
												F_errcode_for_file_access(m)
												mBase = m.M
												v101 = m.ExcPending
												if v101 != 0 {
													return
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
													F_errmsg(m, v98, v7)
													mBase = m.M
													v104 = m.ExcPending
													if v104 != 0 {
														return
													} else {
														F_errfinish(m, int32(523542), v99, int32(395499))
														mBase = m.M
														v108 = m.ExcPending
														if v108 != 0 {
															return
														} else {
															m.G0 = v7 + int32(48)
															return
														}
													}
												}
											}
										}
									}
								} else {
									v86 = F_fsync(m, v71)
									mBase = m.M
									v87 = F_CloseTransientFile(m, v71)
									mBase = m.M
									v88 = m.ExcPending
									if v88 != 0 {
										return
									} else {
										if v87 == int32(0) {
											m.G0 = v7 + int32(48)
											return
										} else {
											v92 = F_errstart(m, l2, int32(0))
											mBase = m.M
											v93 = m.ExcPending
											if v93 != 0 {
												return
											} else {
												if v92 == int32(0) {
													m.G0 = v7 + int32(48)
													return
												} else {
													v98 = int32(312979)
													v99 = int32(3823)
													F_errcode_for_file_access(m)
													mBase = m.M
													v101 = m.ExcPending
													if v101 != 0 {
														return
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
														F_errmsg(m, v98, v7)
														mBase = m.M
														v104 = m.ExcPending
														if v104 != 0 {
															return
														} else {
															F_errfinish(m, int32(523542), v99, int32(395499))
															mBase = m.M
															v108 = m.ExcPending
															if v108 != 0 {
																return
															} else {
																m.G0 = v7 + int32(48)
																return
															}
														}
													}
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_preprocess_aggrefs_walker(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
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
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
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
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
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
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
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
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v276 int32
	_ = v276
	var v281 int32
	_ = v281
	var v288 int32
	_ = v288
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v328 int32
	_ = v328
	var v331 int32
	_ = v331
	var v334 int32
	_ = v334
	var v336 int32
	_ = v336
	var v345 int32
	_ = v345
	var v350 int32
	_ = v350
	var v353 int32
	_ = v353
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v358 int32
	_ = v358
	var v363 int32
	_ = v363
	var v370 int32
	_ = v370
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v389 int32
	_ = v389
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v396 int32
	_ = v396
	var v398 int32
	_ = v398
	var v400 int32
	_ = v400
	var v402 int32
	_ = v402
	var v404 int32
	_ = v404
	var v411 int32
	_ = v411
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
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v449 int32
	_ = v449
	var v451 int32
	_ = v451
	var v459 int32
	_ = v459
	var v461 int32
	_ = v461
	var v465 int32
	_ = v465
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v478 int32
	_ = v478
	var v480 int32
	_ = v480
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v491 int32
	_ = v491
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v497 int32
	_ = v497
	var v499 int32
	_ = v499
	var v504 int32
	_ = v504
	var v515 int32
	_ = v515
	var v529 int32
	_ = v529
	var v533 int32
	_ = v533
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
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v553 int32
	_ = v553
	var v556 int32
	_ = v556
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v564 int32
	_ = v564
	var v589 int32
	_ = v589
	var v593 int32
	_ = v593
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v606 int32
	_ = v606
	var v629 int32
	_ = v629
	var v631 int32
	_ = v631
	var v671 int32
	_ = v671
	v3 = int32(0)
	v23 = m.G0
	v25 = v23 - int32(432)
	m.G0 = v25
	if l0 == v3 {
		v671 = v3
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v25 + int32(432)
	return v671
L2:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v30 == int32(9) {
		goto L6
	} else {
		goto L7
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = v629
	*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v631
	v671 = int32(0)
	goto L1
L4:
	;
	v305 = F_palloc0(m, int32(20))
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L9
	} else {
		goto L75
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L9
	} else {
		goto L72
	}
L6:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v35 = F_SearchSysCache1(m, int32(0), v34)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	v266 = F_expression_tree_walker_impl(m, l0, int32(850), l1)
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L9
	} else {
		goto L71
	}
L9:
	;
	return int32(0)
L10:
	;
	if v35 == int32(0) {
		goto L5
	} else {
		goto L11
	}
L11:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v35)+16))
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+22)))
	v43 = v41 + v42
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+52))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v43)+24))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v43)+20))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v43)+16))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v43)+12))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v43)+8))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v43)+48))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v55 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v90 = F_resolve_aggregate_transtype(m, v87, v50, v25+int32(16))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L9
	} else {
		goto L22
	}
L13:
	;
	goto L12
L14:
	;
	goto L15
L15:
	;
	v59 = int32(0)
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v55)+4))
	if v59 < v60 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v63 = v59
	goto L19
L17:
	;
	goto L18
L18:
	;
	goto L12
L19:
	;
	v68 = v63 << (uint(int32(2)) % 32)
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v55)+12))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v70+v68)))
	*(*int32)(unsafe.Add(mBase, uint32(v25+int32(16)+v68))) = v72
	v75 = v63 + int32(1)
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v55)+4))
	if v75 < v76 {
		v63 = v75
		goto L19
	} else {
		goto L21
	}
L20:
	;
	goto L18
L21:
	;
	goto L20
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v90
	v93 = int32(-1)
	v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v94 == int32(0) {
		v107 = v93
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+42)))
	v109 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_get_typlenbyval(m, v109, v25+int32(422), v25+int32(421))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L9
	} else {
		goto L28
	}
L24:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v94)+12))
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v97)))
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v98)+4))
	v100 = F_exprType(m, v99)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L9
	} else {
		goto L25
	}
L25:
	;
	if v100 != v90 {
		v107 = v93
		goto L23
	} else {
		goto L26
	}
L26:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v98)+4))
	v104 = F_exprTypmod(m, v103)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L9
	} else {
		goto L27
	}
L27:
	;
	v107 = v104
	goto L23
L28:
	;
	v120 = F_SysCacheGetAttr(m, int32(0), v35, int32(21), v25+int32(420))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L9
	} else {
		goto L29
	}
L29:
	;
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+420)))
	if v122 == int32(0) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	F_getTypeInputInfo(m, v90, v25+int32(428), v25+int32(424))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L9
	} else {
		goto L33
	}
L31:
	;
	v141 = v3
	goto L32
L32:
	;
	F_ReleaseCatCache(m, v35)
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L9
	} else {
		goto L37
	}
L33:
	;
	v131 = F_text_to_cstring(m, v120)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L9
	} else {
		goto L34
	}
L34:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v25)+428))
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v25)+424))
	v136 = F_OidInputFunctionCall(m, v133, v131, v134, int32(-1))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L9
	} else {
		goto L35
	}
L35:
	;
	F_pfree(m, v131)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L9
	} else {
		goto L36
	}
L36:
	;
	v141 = v136
	goto L32
L37:
	;
	v144 = F_contain_volatile_functions(m, l0)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L9
	} else {
		goto L38
	}
L38:
	;
	if v144 != 0 {
		v288 = v3
		goto L4
	} else {
		goto L39
	}
L39:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(l1)+328))
	if v146 == int32(0) {
		v288 = v3
		goto L4
	} else {
		goto L40
	}
L40:
	;
	v149 = int32(0)
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v146)+4))
	if v150 <= v149 {
		v288 = v3
		goto L4
	} else {
		goto L41
	}
L41:
	;
	v156 = v149
	v160 = v3
	v161 = int32(-1)
	goto L42
L42:
	;
	v177 = v161 + int32(1)
	v178 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v146)+12))
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v179+v156<<(uint(int32(2))%32))))
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v183)+4))
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v184)+12))
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v185)))
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v186)+16))
	if v178 != v187 {
		v244 = v160
		goto L45
	} else {
		goto L46
	}
L43:
	;
	F_list_free(m, v160)
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L9
	} else {
		goto L68
	}
L44:
	;
	goto L43
L45:
	;
	v246 = v156 + int32(1)
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v146)+4))
	if v246 < v247 {
		v156 = v246
		v160 = v244
		v161 = v177
		goto L42
	} else {
		goto L67
	}
L46:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v186)+20))
	if v189 != v190 {
		v244 = v160
		goto L45
	} else {
		goto L47
	}
L47:
	;
	v192 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)))
	v193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v186)+48)))
	if v192 != v193 {
		v244 = v160
		goto L45
	} else {
		goto L48
	}
L48:
	;
	v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+49)))
	v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v186)+49)))
	if v195 != v196 {
		v244 = v160
		goto L45
	} else {
		goto L49
	}
L49:
	;
	v198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+50)))
	v199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v186)+50)))
	if v198 != v199 {
		v244 = v160
		goto L45
	} else {
		goto L50
	}
L50:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v186)+32))
	v203 = F_equal(m, v201, v202)
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L9
	} else {
		goto L51
	}
L51:
	;
	if v203 == int32(0) {
		v244 = v160
		goto L45
	} else {
		goto L52
	}
L52:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v186)+36))
	v209 = F_equal(m, v207, v208)
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L9
	} else {
		goto L53
	}
L53:
	;
	if v209 == int32(0) {
		v244 = v160
		goto L45
	} else {
		goto L54
	}
L54:
	;
	v213 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v186)+40))
	v215 = F_equal(m, v213, v214)
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L9
	} else {
		goto L55
	}
L55:
	;
	if v215 == int32(0) {
		v244 = v160
		goto L45
	} else {
		goto L56
	}
L56:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v186)+44))
	v221 = F_equal(m, v219, v220)
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L9
	} else {
		goto L57
	}
L57:
	;
	if v221 == int32(0) {
		v244 = v160
		goto L45
	} else {
		goto L58
	}
L58:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v186)+4))
	if v225 != v226 {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v238 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v183)+12)))
	if v238 != int32(1) {
		v244 = v160
		goto L45
	} else {
		goto L65
	}
L60:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v186)+8))
	if v228 != v229 {
		goto L59
	} else {
		goto L61
	}
L61:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v186)+12))
	if v231 != v232 {
		goto L59
	} else {
		goto L62
	}
L62:
	;
	v234 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v186)+28))
	v236 = F_equal(m, v234, v235)
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L9
	} else {
		goto L63
	}
L63:
	;
	if v236 != 0 {
		goto L44
	} else {
		goto L64
	}
L64:
	;
	goto L59
L65:
	;
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v183)+8))
	v242 = F_lappend_int(m, v160, v241)
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L9
	} else {
		goto L66
	}
L66:
	;
	v244 = v242
	goto L45
L67:
	;
	v288 = v244
	goto L4
L68:
	;
	if v177 == int32(-1) {
		v288 = int32(0)
		goto L4
	} else {
		goto L69
	}
L69:
	;
	v254 = *(*int32)(unsafe.Add(mBase, uint32(l1)+328))
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v254)+12))
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v255+v177<<(uint(int32(2))%32))))
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v259)+4))
	v261 = F_lappend(m, v260, l0)
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L9
	} else {
		goto L70
	}
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v259)+4)) = v261
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v259)+8))
	v629 = v264
	v631 = v177
	goto L3
L71:
	;
	v671 = v266
	goto L1
L72:
	;
	v272 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v25))) = v272
	F_errmsg_internal(m, int32(54012), v25)
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L9
	} else {
		goto L73
	}
L73:
	;
	F_errfinish(m, int32(521708), int32(153), int32(355431))
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L9
	} else {
		goto L74
	}
L74:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v305)+16)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v305))) = int32(327)
	*(*int32)(unsafe.Add(mBase, uint32(v25)+12)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v25)+428)) = l0
	v315 = F_list_make1_impl(m, int32(1), v25+int32(12))
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L9
	} else {
		goto L76
	}
L76:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v305)+12)) = uint8(base.B2i32(v108 != int32(119)))
	*(*int32)(unsafe.Add(mBase, uint32(v305)+4)) = v315
	v321 = *(*int32)(unsafe.Add(mBase, uint32(l1)+328))
	if v321 != 0 {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v321)+4))
	v324 = v322
	goto L79
L78:
	;
	v324 = int32(0)
	goto L79
L79:
	;
	v325 = F_lappend(m, v321, v305)
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L9
	} else {
		goto L80
	}
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+328)) = v325
	v328 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v328 == int32(0) {
		goto L82
	} else {
		goto L83
	}
L81:
	;
	F_get_typlenbyval(m, v90, v25+int32(424), v25+int32(419))
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L9
	} else {
		goto L86
	}
L82:
	;
	v331 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v331 == int32(0) {
		goto L81
	} else {
		goto L85
	}
L83:
	;
	goto L84
L84:
	;
	v334 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+340)) = uint8(v334)
	v336 = *(*int32)(unsafe.Add(mBase, uint32(l1)+336))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+336)) = v336 + v334
	goto L81
L85:
	;
	goto L84
L86:
	;
	if v108 == int32(119) {
		goto L88
	} else {
		goto L89
	}
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v305)+8)) = v606
	v629 = v606
	v631 = v324
	goto L3
L88:
	;
	v445 = F_palloc0(m, int32(56))
	mBase = m.M
	v446 = m.ExcPending
	if v446 != 0 {
		goto L9
	} else {
		goto L110
	}
L89:
	;
	if v288 == int32(0) {
		goto L88
	} else {
		goto L90
	}
L90:
	;
	v350 = *(*int32)(unsafe.Add(mBase, uint32(v288)+4))
	if v350 <= int32(0) {
		goto L88
	} else {
		goto L91
	}
L91:
	;
	v353 = int32(*(*int16)(unsafe.Add(mBase, uint32(v25)+424)))
	v355 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+420)))
	v356 = int32(1)
	v358 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+419)))
	v363 = int32(0)
	v370 = v350
	goto L92
L92:
	;
	v383 = *(*int32)(unsafe.Add(mBase, uint32(l1)+332))
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v383)+12))
	v385 = *(*int32)(unsafe.Add(mBase, uint32(v288)+12))
	v386 = int32(2)
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v385+v363<<(uint(v386)%32))))
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v384+v389<<(uint(v386)%32))))
	v394 = *(*int32)(unsafe.Add(mBase, uint32(v393)+12))
	if v49 != v394 {
		v415 = v370
		goto L95
	} else {
		goto L96
	}
L93:
	;
	if v389 != int32(-1) {
		v606 = v389
		goto L87
	} else {
		goto L109
	}
L94:
	;
	goto L93
L95:
	;
	v418 = v363 + int32(1)
	if v418 < v415 {
		v363 = v418
		v370 = v415
		goto L92
	} else {
		goto L108
	}
L96:
	;
	v396 = *(*int32)(unsafe.Add(mBase, uint32(v393)+28))
	if v90 != v396 {
		v415 = v370
		goto L95
	} else {
		goto L97
	}
L97:
	;
	v398 = *(*int32)(unsafe.Add(mBase, uint32(v393)+16))
	if v46 != v398 {
		v415 = v370
		goto L95
	} else {
		goto L98
	}
L98:
	;
	v400 = *(*int32)(unsafe.Add(mBase, uint32(v393)+20))
	if v45 != v400 {
		v415 = v370
		goto L95
	} else {
		goto L99
	}
L99:
	;
	v402 = *(*int32)(unsafe.Add(mBase, uint32(v393)+24))
	if v47 != v402 {
		v415 = v370
		goto L95
	} else {
		goto L100
	}
L100:
	;
	v404 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v393)+52)))
	if v355&v356 != 0 {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	if v404&int32(1) == int32(0) {
		v415 = v370
		goto L95
	} else {
		goto L104
	}
L102:
	;
	goto L103
L103:
	;
	if v404&int32(1) != 0 {
		v415 = v370
		goto L95
	} else {
		goto L105
	}
L104:
	;
	goto L94
L105:
	;
	v411 = *(*int32)(unsafe.Add(mBase, uint32(v393)+48))
	v412 = F_datumIsEqual(m, v141, v411, v358&v356, v353)
	mBase = m.M
	v413 = m.ExcPending
	if v413 != 0 {
		goto L9
	} else {
		goto L106
	}
L106:
	;
	if v412 != 0 {
		goto L94
	} else {
		goto L107
	}
L107:
	;
	v414 = *(*int32)(unsafe.Add(mBase, uint32(v288)+4))
	v415 = v414
	goto L95
L108:
	;
	goto L88
L109:
	;
	goto L88
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v445))) = int32(328)
	v449 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v445)+4)) = v449
	v451 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v445)+24)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v445)+12)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v445)+8)) = v451
	*(*int32)(unsafe.Add(mBase, uint32(v445)+32)) = v107
	*(*int32)(unsafe.Add(mBase, uint32(v445)+28)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v445)+20)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v445)+16)) = v46
	v459 = int32(*(*int16)(unsafe.Add(mBase, uint32(v25)+424)))
	*(*int32)(unsafe.Add(mBase, uint32(v445)+36)) = v459
	v461 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+419)))
	*(*int32)(unsafe.Add(mBase, uint32(v445)+48)) = v141
	*(*int32)(unsafe.Add(mBase, uint32(v445)+44)) = v44
	*(*uint8)(unsafe.Add(mBase, uint32(v445)+40)) = uint8(v461)
	v465 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+420)))
	*(*uint8)(unsafe.Add(mBase, uint32(v445)+52)) = uint8(v465)
	v467 = *(*int32)(unsafe.Add(mBase, uint32(l1)+332))
	if v467 != 0 {
		goto L111
	} else {
		goto L112
	}
L111:
	;
	v468 = *(*int32)(unsafe.Add(mBase, uint32(v467)+4))
	v470 = v468
	goto L113
L112:
	;
	v470 = int32(0)
	goto L113
L113:
	;
	v471 = F_lappend(m, v467, v445)
	mBase = m.M
	v472 = m.ExcPending
	if v472 != 0 {
		goto L9
	} else {
		goto L114
	}
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+332)) = v471
	v474 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+340)))
	if v474 != 0 {
		v606 = v470
		goto L87
	} else {
		goto L115
	}
L115:
	;
	v475 = *(*int32)(unsafe.Add(mBase, uint32(v445)+24))
	if v475 == int32(0) {
		goto L116
	} else {
		goto L117
	}
L116:
	;
	v478 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+340)) = uint8(v478)
	v606 = v470
	goto L87
L117:
	;
	goto L118
L118:
	;
	v480 = *(*int32)(unsafe.Add(mBase, uint32(v445)+28))
	if v480 != int32(2281) {
		v606 = v470
		goto L87
	} else {
		goto L119
	}
L119:
	;
	v483 = *(*int32)(unsafe.Add(mBase, uint32(v445)+16))
	if v483 != 0 {
		goto L121
	} else {
		goto L122
	}
L120:
	;
	if v488 != int32(6294) {
		goto L125
	} else {
		goto L126
	}
L121:
	;
	v484 = *(*int32)(unsafe.Add(mBase, uint32(v445)+20))
	if v484 != 0 {
		v488 = v483
		goto L120
	} else {
		goto L124
	}
L122:
	;
	goto L123
L123:
	;
	v485 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+341)) = uint8(v485)
	v487 = *(*int32)(unsafe.Add(mBase, uint32(v445)+16))
	v488 = v487
	goto L120
L124:
	;
	goto L123
L125:
	;
	v491 = *(*int32)(unsafe.Add(mBase, uint32(v445)+20))
	if v491 != int32(6295) {
		v606 = v470
		goto L87
	} else {
		goto L128
	}
L126:
	;
	goto L127
L127:
	;
	v494 = int32(0)
	v495 = m.G0
	v497 = v495 - int32(16)
	m.G0 = v497
	v499 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v499 == v494 {
		goto L132
	} else {
		goto L133
	}
L128:
	;
	goto L127
L129:
	;
	if v564 != 0 {
		v606 = v470
		goto L87
	} else {
		goto L154
	}
L130:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v589 = m.ExcPending
	if v589 != 0 {
		goto L9
	} else {
		goto L151
	}
L131:
	;
	m.G0 = v497 + int32(16)
	goto L129
L132:
	;
	v564 = int32(1)
	goto L131
L133:
	;
	goto L134
L134:
	;
	v504 = *(*int32)(unsafe.Add(mBase, uint32(v499)+4))
	if v504 <= int32(0) {
		v564 = int32(1)
		goto L131
	} else {
		goto L135
	}
L135:
	;
	v515 = v494
	goto L136
L136:
	;
	v529 = *(*int32)(unsafe.Add(mBase, uint32(v499)+12))
	v533 = *(*int32)(unsafe.Add(mBase, uint32(v529+v515<<(uint(int32(2))%32))))
	v534 = *(*int32)(unsafe.Add(mBase, uint32(v533)+4))
	v535 = F_exprType(m, v534)
	mBase = m.M
	v536 = m.ExcPending
	if v536 != 0 {
		goto L9
	} else {
		goto L138
	}
L137:
	;
	v564 = v538
	goto L131
L138:
	;
	v537 = int32(2249)
	v538 = base.B2i32(v535 != v537)
	if v535 == v537 {
		v564 = v538
		goto L131
	} else {
		goto L139
	}
L139:
	;
	v542 = F_SearchSysCache1(m, int32(82), v535)
	mBase = m.M
	v543 = m.ExcPending
	if v543 != 0 {
		goto L9
	} else {
		goto L140
	}
L140:
	;
	if v542 == int32(0) {
		goto L130
	} else {
		goto L141
	}
L141:
	;
	v546 = *(*int32)(unsafe.Add(mBase, uint32(v542)+16))
	v547 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v546)+22)))
	v548 = v546 + v547
	v549 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v548)+78)))
	if v549 != 0 {
		goto L142
	} else {
		goto L143
	}
L142:
	;
	F_ReleaseCatCache(m, v542)
	mBase = m.M
	v556 = m.ExcPending
	if v556 != 0 {
		goto L9
	} else {
		goto L149
	}
L143:
	;
	v550 = *(*int32)(unsafe.Add(mBase, uint32(v548)+112))
	if v550 != 0 {
		goto L144
	} else {
		goto L145
	}
L144:
	;
	v551 = *(*int32)(unsafe.Add(mBase, uint32(v548)+108))
	if v551 != 0 {
		goto L142
	} else {
		goto L147
	}
L145:
	;
	goto L146
L146:
	;
	F_ReleaseCatCache(m, v542)
	mBase = m.M
	v553 = m.ExcPending
	if v553 != 0 {
		goto L9
	} else {
		goto L148
	}
L147:
	;
	goto L146
L148:
	;
	v564 = int32(0)
	goto L131
L149:
	;
	v558 = v515 + int32(1)
	v559 = *(*int32)(unsafe.Add(mBase, uint32(v499)+4))
	if v558 < v559 {
		v515 = v558
		goto L136
	} else {
		goto L150
	}
L150:
	;
	goto L137
L151:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v497))) = v535
	F_errmsg_internal(m, int32(55124), v497)
	mBase = m.M
	v593 = m.ExcPending
	if v593 != 0 {
		goto L9
	} else {
		goto L152
	}
L152:
	;
	F_errfinish(m, int32(521728), int32(2134), int32(360743))
	mBase = m.M
	v598 = m.ExcPending
	if v598 != 0 {
		goto L9
	} else {
		goto L153
	}
L153:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L154:
	;
	v599 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+341)) = uint8(v599)
	v606 = v470
	goto L87
}
func F_printsimple_startup(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v166 int32
	_ = v166
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	F_pq_beginmessage(m, v10, int32(84))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v15 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2))))
	F_enlargeStringInfo(m, v10, int32(2))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	v22 = int32(8)
	v26 = v15<<(uint(v22)%32) | int32(base.Ui32(v15)>>(uint(v22)%32))
	*(*uint16)(unsafe.Add(mBase, uint32(v19+v20))) = uint16(v26)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v19 + int32(2)
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if int32(0) < v31 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v37 = v31
	v40 = int32(0)
	goto L7
L5:
	;
	goto L6
L6:
	;
	F_pq_endmessage(m, v10)
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L1
	} else {
		goto L17
	}
L7:
	;
	v43 = int32(4)
	v48 = l2 + int32(20) + v37<<(uint(v43)%32) + v40*int32(100)
	F_pq_sendstring(m, v10, v48+v43)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L1
	} else {
		goto L9
	}
L8:
	;
	goto L6
L9:
	;
	F_enlargeStringInfo(m, v10, int32(4))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	*(*int32)(unsafe.Add(mBase, uint32(v56+v57))) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v56 + int32(4)
	F_enlargeStringInfo(m, v10, int32(2))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	v70 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v67+v68))) = uint16(v70)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v67 + int32(2)
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v48)+68))
	F_enlargeStringInfo(m, v10, int32(4))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	v82 = int32(24)
	v84 = int32(65280)
	v86 = int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v79+v80))) = v75<<(uint(v82)%32) | v75&v84<<(uint(v86)%32) | (int32(base.Ui32(v75)>>(uint(v86)%32))&v84 | int32(base.Ui32(v75)>>(uint(v82)%32)))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v79 + int32(4)
	v101 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v48)+72)))
	F_enlargeStringInfo(m, v10, int32(2))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	v108 = int32(8)
	v112 = v101<<(uint(v108)%32) | int32(base.Ui32(v101)>>(uint(v108)%32))
	*(*uint16)(unsafe.Add(mBase, uint32(v105+v106))) = uint16(v112)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v105 + int32(2)
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v48)+76))
	F_enlargeStringInfo(m, v10, int32(4))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	v124 = int32(24)
	v126 = int32(65280)
	v128 = int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v121+v122))) = v117<<(uint(v124)%32) | v117&v126<<(uint(v128)%32) | (int32(base.Ui32(v117)>>(uint(v128)%32))&v126 | int32(base.Ui32(v117)>>(uint(v124)%32)))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v121 + int32(4)
	F_enlargeStringInfo(m, v10, int32(2))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	v149 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v146+v147))) = uint16(v149)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v146 + int32(2)
	v155 = v40 + int32(1)
	v156 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v155 < v156 {
		v37 = v156
		v40 = v155
		goto L7
	} else {
		goto L16
	}
L16:
	;
	goto L8
L17:
	;
	m.G0 = v10 + int32(16)
	return
}
func F_processIndirection(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
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
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	v3 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	if l0 == v3 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L17
	} else {
		goto L30
	}
L2:
	;
	m.G0 = v9 + int32(32)
	return v76
L3:
	;
	v76 = int32(0)
	goto L2
L4:
	;
	goto L5
L5:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v15 = l0
	v19 = v3
	goto L8
L6:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v72)+4))
	if v70 == v73 {
		goto L27
	} else {
		goto L28
	}
L7:
	;
	if v67 == int32(0) {
		v76 = v65
		goto L2
	} else {
		goto L26
	}
L8:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	switch v21 - int32(14) {
	case 0:
		goto L12
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11:
		v65 = v15
		v67 = v19
		goto L7
	case 12:
		goto L13
	default:
		goto L14
	}
L9:
	;
	v65 = int32(0)
	v67 = v62
	goto L7
L10:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
	if v63 != 0 {
		v15 = v63
		v19 = v62
		goto L8
	} else {
		goto L25
	}
L11:
	;
	v61 = v15 + int32(4)
	v62 = v15
	goto L10
L12:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v15)+36))
	if v52 == int32(0) {
		v65 = v15
		v67 = v19
		goto L7
	} else {
		goto L23
	}
L13:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v15)+16))
	v30 = F_get_typ_typrelid(m, v29)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L17
	} else {
		goto L18
	}
L14:
	;
	if v21 != int32(55) {
		v65 = v15
		v67 = v19
		goto L7
	} else {
		goto L15
	}
L15:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v15)+20))
	if v26 == int32(2) {
		goto L11
	} else {
		goto L16
	}
L16:
	;
	v70 = v15
	v72 = v15
	goto L6
L17:
	;
	return int32(0)
L18:
	;
	if v30 == int32(0) {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)+12))
	v38 = int32(*(*int16)(unsafe.Add(mBase, uint32(v37))))
	v40 = F_get_attname(m, v30, v38, int32(0))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L17
	} else {
		goto L20
	}
L20:
	;
	v42 = F_quote_identifier(m, v40)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L17
	} else {
		goto L21
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v42
	F_appendStringInfo(m, v14, int32(187524), v9+int32(16))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L17
	} else {
		goto L22
	}
L22:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)+12))
	v61 = v51
	v62 = v19
	goto L10
L23:
	;
	F_printSubscripts(m, v15, l1)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L17
	} else {
		goto L24
	}
L24:
	;
	v61 = v15 + int32(36)
	v62 = v19
	goto L10
L25:
	;
	goto L9
L26:
	;
	v70 = v65
	v72 = v67
	goto L6
L27:
	;
	v75 = v72
	goto L29
L28:
	;
	v75 = v70
	goto L29
L29:
	;
	v76 = v75
	goto L2
L30:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v15)+16))
	v91 = F_format_type_be(m, v90)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L17
	} else {
		goto L31
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v91
	F_errmsg_internal(m, int32(387505), v9)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L17
	} else {
		goto L32
	}
L32:
	;
	F_errfinish(m, int32(516743), int32(12940), int32(267093))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L17
	} else {
		goto L33
	}
L33:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_processTypesSpec(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
	v7 = F_typenameTypeId(m, int32(0), v6)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l1))) = v7
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		if int32(2) <= v10 {
			v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
			v16 = F_typenameTypeId(m, int32(0), v15)
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return
			} else {
				v18 = v16
				*(*int32)(unsafe.Add(mBase, uint32(l2))) = v18
				v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				if int32(3) <= v20 {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return
					} else {
						F_errcode(m, int32(16801924))
						mBase = m.M
						v29 = m.ExcPending
						if v29 != 0 {
							return
						} else {
							F_errmsg(m, int32(478383), int32(0))
							mBase = m.M
							v33 = m.ExcPending
							if v33 != 0 {
								return
							} else {
								F_errfinish(m, int32(517098), int32(1128), int32(513785))
								mBase = m.M
								v38 = m.ExcPending
								if v38 != 0 {
									return
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					}
				} else {
					return
				}
			}
		} else {
			v18 = v7
			*(*int32)(unsafe.Add(mBase, uint32(l2))) = v18
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			if int32(3) <= v20 {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return
				} else {
					F_errcode(m, int32(16801924))
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
						return
					} else {
						F_errmsg(m, int32(478383), int32(0))
						mBase = m.M
						v33 = m.ExcPending
						if v33 != 0 {
							return
						} else {
							F_errfinish(m, int32(517098), int32(1128), int32(513785))
							mBase = m.M
							v38 = m.ExcPending
							if v38 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			} else {
				return
			}
		}
	}
}
func F_process_implied_equality(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
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
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v151 int32
	_ = v151
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	v10 = F_copyObjectImpl(m, l3)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v14 = F_copyObjectImpl(m, l4)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v16 = F_make_opclause(m, l1, v10, v14, l2)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	if l7 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	return v249
L6:
	;
	v33 = F_pull_varnos(m, l0, v32)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L19
	}
L7:
	;
	v32 = v16
	goto L6
L8:
	;
	goto L9
L9:
	;
	v20 = F_eval_const_expressions(m, l0, v16)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	if v20 == int32(0) {
		v32 = int32(0)
		goto L6
	} else {
		goto L11
	}
L11:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	if v24 != int32(7) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v32 = v20
	goto L6
L13:
	;
	goto L14
L14:
	;
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+24)))
	if v27 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v32 = v20
	goto L6
L16:
	;
	goto L17
L17:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v20)+20))
	if v29 != 0 {
		v249 = int32(0)
		goto L5
	} else {
		goto L18
	}
L18:
	;
	v32 = v20
	goto L6
L19:
	;
	if v33 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v37 = F_bms_copy(m, l5)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L1
	} else {
		goto L24
	}
L21:
	;
	v142 = v33
	goto L22
L22:
	;
	v151 = int32(0)
	v157 = F_make_restrictinfo(m, l0, v32, int32(1), v151, v151, base.B2i32(v33 == v151), l6, v142, v151, v151)
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L1
	} else {
		goto L49
	}
L23:
	;
	v139 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+319)) = uint8(v139)
	v142 = v131
	goto L22
L24:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v40 = int32(0)
	v47 = base.B2i32(v37|v39 == v40)
	if v37 == v40 {
		v86 = v47
		goto L26
	} else {
		goto L27
	}
L25:
	;
	if v86 != 0 {
		v131 = v37
		goto L23
	} else {
		goto L37
	}
L26:
	;
	goto L25
L27:
	;
	if v39 == int32(0) {
		v86 = v47
		goto L26
	} else {
		goto L28
	}
L28:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
	if v53 != v54 {
		v86 = int32(0)
		goto L26
	} else {
		goto L29
	}
L29:
	;
	v56 = int32(1)
	if v53 <= v56 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v59 = v56
	goto L32
L31:
	;
	v59 = v53
	goto L32
L32:
	;
	v60 = int32(8)
	v65 = int32(0)
	goto L33
L33:
	;
	v73 = v65 << (uint(int32(2)) % 32)
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v37+v60+v73)))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v73+(v39+v60))))
	v78 = base.B2i32(v75 == v77)
	if v77 != v75 {
		v86 = v78
		goto L26
	} else {
		goto L35
	}
L34:
	;
	v86 = v78
	goto L26
L35:
	;
	v81 = v65 + int32(1)
	if v81 != v59 {
		v65 = v81
		goto L33
	} else {
		goto L36
	}
L36:
	;
	goto L34
L37:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	if v90 == int32(0) {
		v131 = v37
		goto L23
	} else {
		goto L38
	}
L38:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v90)+4))
	if v93 <= int32(0) {
		v131 = v37
		goto L23
	} else {
		goto L39
	}
L39:
	;
	v98 = v37
	v99 = int32(0)
	goto L40
L40:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v90)+12))
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v106+v99<<(uint(int32(2))%32))))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v110)+20))
	if v111 != int32(1) {
		v125 = v98
		goto L42
	} else {
		goto L43
	}
L41:
	;
	v131 = v125
	goto L23
L42:
	;
	v127 = v99 + int32(1)
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v90)+4))
	if v127 < v128 {
		v98 = v125
		v99 = v127
		goto L40
	} else {
		goto L48
	}
L43:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v110)+24))
	v115 = F_bms_is_member(m, v114, v98)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	if v115 == int32(0) {
		v125 = v98
		goto L42
	} else {
		goto L45
	}
L45:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v110)+24))
	v120 = F_bms_del_member(m, v98, v119)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v110)+16))
	v123 = F_bms_del_members(m, v120, v122)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	v125 = v123
	goto L42
L48:
	;
	goto L41
L49:
	;
	if v142 == int32(0) {
		goto L51
	} else {
		goto L52
	}
L50:
	;
	if v205 == int32(2) {
		goto L66
	} else {
		goto L67
	}
L51:
	;
	v205 = int32(0)
	goto L50
L52:
	;
	goto L53
L53:
	;
	v167 = int32(1)
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v142)+4))
	if v168 <= v167 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v171 = v167
	goto L56
L55:
	;
	v171 = v168
	goto L56
L56:
	;
	v174 = int32(0)
	v176 = v174
	v177 = v174
	goto L57
L57:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v142+int32(8)+v176<<(uint(int32(2))%32))))
	if v185 != 0 {
		goto L60
	} else {
		goto L61
	}
L58:
	;
	v205 = v198
	goto L50
L59:
	;
	goto L58
L60:
	;
	v186 = int32(2)
	if v177 != 0 {
		v198 = v186
		goto L59
	} else {
		goto L63
	}
L61:
	;
	v191 = v177
	goto L62
L62:
	;
	v194 = v176 + int32(1)
	if v194 != v171 {
		v176 = v194
		v177 = v191
		goto L57
	} else {
		goto L65
	}
L63:
	;
	v187 = int32(1)
	if base.Ui32(v187) < base.Ui32(base.I32_popcnt(v185)) {
		v198 = v186
		goto L59
	} else {
		goto L64
	}
L64:
	;
	v191 = v187
	goto L62
L65:
	;
	v198 = v191
	goto L59
L66:
	;
	v209 = F_pull_var_clause(m, v32, int32(26))
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L1
	} else {
		goto L69
	}
L67:
	;
	goto L68
L68:
	;
	v216 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157)+10)))
	if v216 != 0 {
		goto L72
	} else {
		goto L73
	}
L69:
	;
	F_add_vars_to_targetlist(m, l0, v209, v142)
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L1
	} else {
		goto L70
	}
L70:
	;
	F_list_free(m, v209)
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L1
	} else {
		goto L71
	}
L71:
	;
	goto L68
L72:
	;
	F_distribute_restrictinfo_to_rels(m, l0, v157)
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L1
	} else {
		goto L84
	}
L73:
	;
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v157)+4))
	if v217 == int32(0) {
		goto L72
	} else {
		goto L74
	}
L74:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v217)))
	if v220 != int32(17) {
		goto L72
	} else {
		goto L75
	}
L75:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v217)+28))
	if v223 == int32(0) {
		goto L72
	} else {
		goto L76
	}
L76:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v223)+4))
	if v226 != int32(2) {
		goto L72
	} else {
		goto L77
	}
L77:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v217)+4))
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v223)+12))
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v230)))
	v232 = F_exprType(m, v231)
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L1
	} else {
		goto L78
	}
L78:
	;
	v234 = F_op_mergejoinable(m, v229, v232)
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L1
	} else {
		goto L79
	}
L79:
	;
	if v234 == int32(0) {
		goto L72
	} else {
		goto L80
	}
L80:
	;
	v238 = F_contain_volatile_functions(m, v157)
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	if v238 != 0 {
		goto L72
	} else {
		goto L82
	}
L82:
	;
	v240 = F_get_mergejoin_opfamilies(m, v229)
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L1
	} else {
		goto L83
	}
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v157)+96)) = v240
	goto L72
L84:
	;
	v249 = v157
	goto L5
}
func F_prsd_end(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)+24))
	if v5 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v8 = v5
	goto L4
L2:
	;
	goto L3
L3:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v4)+8))
	if v18 != 0 {
		goto L9
	} else {
		goto L10
	}
L4:
	;
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+24))
	F_pfree(m, v8)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L3
L6:
	;
	return int32(0)
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4)+24)) = v9
	if v9 != 0 {
		v8 = v9
		goto L4
	} else {
		goto L8
	}
L8:
	;
	goto L5
L9:
	;
	F_pfree(m, v18)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L6
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v4)+12))
	if v21 != 0 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	goto L11
L13:
	;
	F_pfree(m, v21)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L6
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	F_pfree(m, v4)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L6
	} else {
		goto L17
	}
L16:
	;
	goto L15
L17:
	;
	return int32(0)
}
func F_pub_collist_validate(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	v3 = int32(0)
	v10 = m.G0
	v12 = v10 + int32(-64)
	m.G0 = v12
	if l1 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v12 - int32(-64)
	return v138
L2:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v24 = v3
	v26 = v3
	goto L10
L3:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if int32(0) < v14 {
		goto L2
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v138 = v3
	goto L1
L6:
	;
	goto L5
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L12
	} else {
		goto L35
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L12
	} else {
		goto L31
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L12
	} else {
		goto L27
	}
L10:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v30+v26<<(uint(int32(2))%32))))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)+4))
	v36 = F_get_attnum(m, v29, v35)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L12
	} else {
		goto L23
	}
L12:
	;
	return int32(0)
L13:
	;
	if v36 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	if v36 <= int32(0) {
		goto L9
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	goto L11
L17:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17+int32(10)+v42<<(uint(int32(4))%32)+v36*int32(100)))))
	if v49 == int32(118) {
		goto L8
	} else {
		goto L18
	}
L18:
	;
	v52 = F_bms_is_member(m, v36, v24)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L12
	} else {
		goto L19
	}
L19:
	;
	if v52 != 0 {
		goto L7
	} else {
		goto L20
	}
L20:
	;
	v54 = F_bms_add_member(m, v24, v36)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L12
	} else {
		goto L21
	}
L21:
	;
	v57 = v26 + int32(1)
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v57 < v58 {
		v24 = v54
		v26 = v57
		goto L10
	} else {
		goto L22
	}
L22:
	;
	v138 = v54
	goto L1
L23:
	;
	F_errcode(m, int32(50360452))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L12
	} else {
		goto L24
	}
L24:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v67 + int32(4)
	F_errmsg(m, int32(77509), v12)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L12
	} else {
		goto L25
	}
L25:
	;
	F_errfinish(m, int32(519434), int32(571), int32(373836))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L12
	} else {
		goto L26
	}
L26:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L27:
	;
	F_errcode(m, int32(393348))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L12
	} else {
		goto L28
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = v35
	F_errmsg(m, int32(81081), v10+int32(-16))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L12
	} else {
		goto L29
	}
L29:
	;
	F_errfinish(m, int32(519434), int32(577), int32(373836))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L12
	} else {
		goto L30
	}
L30:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L31:
	;
	F_errcode(m, int32(393348))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L12
	} else {
		goto L32
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v35
	F_errmsg(m, int32(81187), v10+int32(-48))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L12
	} else {
		goto L33
	}
L33:
	;
	F_errfinish(m, int32(519434), int32(583), int32(373836))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L12
	} else {
		goto L34
	}
L34:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L35:
	;
	F_errcode(m, int32(290948))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L12
	} else {
		goto L36
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v35
	F_errmsg(m, int32(81138), v10+int32(-32))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L12
	} else {
		goto L37
	}
L37:
	;
	F_errfinish(m, int32(519434), int32(589), int32(373836))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L12
	} else {
		goto L38
	}
L38:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pull_up_sublinks_qual_recurse(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v142 int32
	_ = v142
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v167 int32
	_ = v167
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v292 int32
	_ = v292
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v316 int32
	_ = v316
	var v321 int32
	_ = v321
	var v328 int32
	_ = v328
	var v337 int32
	_ = v337
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v359 int32
	_ = v359
	var v371 int32
	_ = v371
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v384 int32
	_ = v384
	v7 = int32(0)
	v18 = m.G0
	v20 = v18 - int32(16)
	m.G0 = v20
	if l1 == v7 {
		v384 = v7
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v20 + int32(16)
	return v384
L2:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	switch v24 - int32(21) {
	case 0:
		goto L3
	case 1:
		goto L4
	default:
		v384 = l1
		goto L1
	}
L3:
	;
	v269 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	switch v269 {
	case 0:
		goto L68
	default:
		v384 = l1
		goto L1
	case 2:
		goto L69
	}
L4:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	switch v27 {
	case 0:
		goto L5
	default:
		v384 = l1
		goto L1
	case 2:
		goto L6
	}
L5:
	;
	v221 = int32(0)
	v223 = F_convert_EXISTS_sublink_to_join(m, l0, l1, v221, l3)
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L23
	} else {
		goto L55
	}
L6:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v30 = m.G0
	v32 = v30 - int32(16)
	m.G0 = v32
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	if v34 != int32(17) {
		v167 = v7
		goto L7
	} else {
		goto L8
	}
L7:
	;
	m.G0 = v32 + int32(16)
	if v167 != 0 {
		v384 = v167
		goto L1
	} else {
		goto L41
	}
L8:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v28)+28))
	if v37 == int32(0) {
		v167 = v7
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	if v40 != int32(2) {
		v167 = v7
		goto L7
	} else {
		goto L10
	}
L10:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v29)+76))
	if v43 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
	if int32(1) < v44 {
		v167 = v7
		goto L7
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v29)+132))
	if v47 != 0 {
		v167 = v7
		goto L7
	} else {
		goto L15
	}
L14:
	;
	goto L13
L15:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v29)+128))
	if v48 != 0 {
		v167 = v7
		goto L7
	} else {
		goto L16
	}
L16:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v29)+124))
	if v49 != 0 {
		v167 = v7
		goto L7
	} else {
		goto L17
	}
L17:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v29)+52))
	if v50 == int32(0) {
		v167 = v7
		goto L7
	} else {
		goto L18
	}
L18:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
	if v53 != int32(1) {
		v167 = v7
		goto L7
	} else {
		goto L19
	}
L19:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v50)+12))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)+12))
	if v58 != int32(5) {
		v167 = v7
		goto L7
	} else {
		goto L20
	}
L20:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v57)+80))
	if v61 == int32(0) {
		v167 = v7
		goto L7
	} else {
		goto L21
	}
L21:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v61)+4))
	if v64 < int32(2) {
		v167 = v7
		goto L7
	} else {
		goto L22
	}
L22:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v37)+12))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)+4))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v28)+24))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
	v72 = F_contain_volatile_functions(m, v61)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	return int32(0)
L24:
	;
	if v72 != 0 {
		v167 = v7
		goto L7
	} else {
		goto L25
	}
L25:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v57)+80))
	if v76 == int32(0) {
		v142 = v7
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v147 = F_exprType(m, v68)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L23
	} else {
		goto L39
	}
L27:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v76)+4))
	if v79 <= int32(0) {
		v142 = v7
		goto L26
	} else {
		goto L28
	}
L28:
	;
	v93 = int32(0)
	v95 = v7
	goto L29
L29:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v76)+12))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v100+v93<<(uint(int32(2))%32))))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v104)+12))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v105)))
	*(*int32)(unsafe.Add(mBase, uint32(v32))) = v106
	*(*int32)(unsafe.Add(mBase, uint32(v32)+4)) = v106
	v110 = F_list_make1_impl(m, int32(1), v32)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L23
	} else {
		goto L31
	}
L30:
	;
	v142 = v124
	goto L26
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+12)) = v110
	*(*int32)(unsafe.Add(mBase, uint32(v32)+8)) = l0
	v116 = F_convert_testexpr_mutator(m, v68, v32+int32(8))
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L23
	} else {
		goto L32
	}
L32:
	;
	v118 = F_eval_const_expressions(m, l0, v116)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L23
	} else {
		goto L33
	}
L33:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v118)))
	if v120 != int32(7) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v167 = int32(0)
	goto L7
L35:
	;
	goto L36
L36:
	;
	v124 = F_lappend(m, v95, v118)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L23
	} else {
		goto L37
	}
L37:
	;
	v127 = v93 + int32(1)
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v76)+4))
	if v127 < v128 {
		v93 = v127
		v95 = v124
		goto L29
	} else {
		goto L38
	}
L38:
	;
	goto L30
L39:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v57)+104))
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v149)+12))
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v150)))
	v153 = F_make_SAOP_expr(m, v70, v71, v147, v151, v69, v142, int32(0))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L23
	} else {
		goto L40
	}
L40:
	;
	v167 = v153
	goto L7
L41:
	;
	v175 = F_convert_ANY_sublink_to_join(m, l0, l1, l3)
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L23
	} else {
		goto L42
	}
L42:
	;
	if v175 != 0 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*int32)(unsafe.Add(mBase, uint32(v175)+12)) = v177
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v175
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v175)+16))
	v183 = F_pull_up_sublinks_jointree_recurse(m, l0, v180, v20+int32(12))
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L23
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	if l5 == int32(0) {
		goto L48
	} else {
		goto L49
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v175)+16)) = v183
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v175)+28))
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	v192 = F_pull_up_sublinks_qual_recurse(m, l0, v186, v175+int32(12), l3, v175+int32(16), v191)
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L23
	} else {
		goto L47
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v175)+28)) = v192
	v384 = int32(0)
	goto L1
L48:
	;
	v384 = l1
	goto L1
L49:
	;
	goto L50
L50:
	;
	v198 = F_convert_ANY_sublink_to_join(m, l0, l1, l5)
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L23
	} else {
		goto L51
	}
L51:
	;
	if v198 == int32(0) {
		v384 = l1
		goto L1
	} else {
		goto L52
	}
L52:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	*(*int32)(unsafe.Add(mBase, uint32(v198)+12)) = v202
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v198
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v198)+16))
	v208 = F_pull_up_sublinks_jointree_recurse(m, l0, v205, v20+int32(12))
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L23
	} else {
		goto L53
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v198)+16)) = v208
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v198)+28))
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	v217 = F_pull_up_sublinks_qual_recurse(m, l0, v211, v198+int32(12), l5, v198+int32(16), v216)
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L23
	} else {
		goto L54
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v198)+28)) = v217
	v384 = int32(0)
	goto L1
L55:
	;
	if v223 != 0 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*int32)(unsafe.Add(mBase, uint32(v223)+12)) = v225
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v223
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v223)+16))
	v231 = F_pull_up_sublinks_jointree_recurse(m, l0, v228, v20+int32(12))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L23
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	if l5 == int32(0) {
		goto L61
	} else {
		goto L62
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v223)+16)) = v231
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v223)+28))
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	v240 = F_pull_up_sublinks_qual_recurse(m, l0, v234, v223+int32(12), l3, v223+int32(16), v239)
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L23
	} else {
		goto L60
	}
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v223)+28)) = v240
	v384 = v221
	goto L1
L61:
	;
	v384 = l1
	goto L1
L62:
	;
	goto L63
L63:
	;
	v246 = F_convert_EXISTS_sublink_to_join(m, l0, l1, int32(0), l5)
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L23
	} else {
		goto L64
	}
L64:
	;
	if v246 == int32(0) {
		v384 = l1
		goto L1
	} else {
		goto L65
	}
L65:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	*(*int32)(unsafe.Add(mBase, uint32(v246)+12)) = v250
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v246
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v246)+16))
	v256 = F_pull_up_sublinks_jointree_recurse(m, l0, v253, v20+int32(12))
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L23
	} else {
		goto L66
	}
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v246)+16)) = v256
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v246)+28))
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	v265 = F_pull_up_sublinks_qual_recurse(m, l0, v259, v246+int32(12), l5, v246+int32(16), v264)
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L23
	} else {
		goto L67
	}
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v246)+28)) = v265
	v384 = int32(0)
	goto L1
L68:
	;
	v312 = int32(0)
	v313 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v313 == v312 {
		v384 = v312
		goto L1
	} else {
		goto L83
	}
L69:
	;
	v270 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v270)+12))
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v271)))
	if v272 == int32(0) {
		v384 = l1
		goto L1
	} else {
		goto L70
	}
L70:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v272)))
	if v275 != int32(22) {
		v384 = l1
		goto L1
	} else {
		goto L71
	}
L71:
	;
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v272)+4))
	if v278 != 0 {
		v384 = l1
		goto L1
	} else {
		goto L72
	}
L72:
	;
	v280 = F_convert_EXISTS_sublink_to_join(m, l0, v272, int32(1), l3)
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L23
	} else {
		goto L74
	}
L73:
	;
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v295)+16))
	v299 = F_pull_up_sublinks_jointree_recurse(m, l0, v296, v20+int32(8))
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L23
	} else {
		goto L81
	}
L74:
	;
	if v280 != 0 {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v282 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*int32)(unsafe.Add(mBase, uint32(v280)+12)) = v282
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v280
	v295 = v280
	goto L73
L76:
	;
	goto L77
L77:
	;
	if l5 == int32(0) {
		v384 = l1
		goto L1
	} else {
		goto L78
	}
L78:
	;
	v288 = F_convert_EXISTS_sublink_to_join(m, l0, v272, int32(1), l5)
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L23
	} else {
		goto L79
	}
L79:
	;
	if v288 == int32(0) {
		v384 = l1
		goto L1
	} else {
		goto L80
	}
L80:
	;
	v292 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	*(*int32)(unsafe.Add(mBase, uint32(v288)+12)) = v292
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v288
	v295 = v288
	goto L73
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v295)+16)) = v299
	v302 = int32(0)
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v295)+28))
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
	v309 = F_pull_up_sublinks_qual_recurse(m, l0, v303, v295+int32(16), v306, v302, v302)
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L23
	} else {
		goto L82
	}
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v295)+28)) = v309
	v384 = v302
	goto L1
L83:
	;
	v316 = *(*int32)(unsafe.Add(mBase, uint32(v313)+4))
	if int32(0) < v316 {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	v321 = int32(0)
	v328 = v7
	goto L87
L85:
	;
	v359 = v7
	goto L86
L86:
	;
	if v359 == int32(0) {
		goto L95
	} else {
		goto L96
	}
L87:
	;
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v313)+12))
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v337+v321<<(uint(int32(2))%32))))
	v342 = F_pull_up_sublinks_qual_recurse(m, l0, v341, l2, l3, l4, l5)
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L23
	} else {
		goto L89
	}
L88:
	;
	v359 = v346
	goto L86
L89:
	;
	if v342 != 0 {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v344 = F_lappend(m, v328, v342)
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L23
	} else {
		goto L93
	}
L91:
	;
	v346 = v328
	goto L92
L92:
	;
	v348 = v321 + int32(1)
	v349 = *(*int32)(unsafe.Add(mBase, uint32(v313)+4))
	if v348 < v349 {
		v321 = v348
		v328 = v346
		goto L87
	} else {
		goto L94
	}
L93:
	;
	v346 = v344
	goto L92
L94:
	;
	goto L88
L95:
	;
	v384 = int32(0)
	goto L1
L96:
	;
	goto L97
L97:
	;
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v359)+4))
	if v371 == int32(1) {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v359)+12))
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v374)))
	v384 = v375
	goto L1
L99:
	;
	goto L100
L100:
	;
	v376 = F_make_andclause(m, v359)
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L23
	} else {
		goto L101
	}
L101:
	;
	v384 = v376
	goto L1
}
func F_pushf_free_all(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	if l0 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v4 = l0
	goto L4
L2:
	;
	goto L3
L3:
	;
	return
L4:
	;
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v4)+4))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	if v9 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L3
L6:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v4)+20))
	m.T0[v9].(func(*base.Module, int32))(m, v10)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v4)+12))
	if v13 != 0 {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	return
L10:
	;
	goto L8
L11:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v4)+8))
	v16 = F___memset(m, v13, int32(0), v15)
	mBase = m.M
	goto L14
L12:
	;
	goto L13
L13:
	;
	v22 = F___memset(m, v4, int32(0), int32(24))
	mBase = m.M
	goto L16
L14:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v4)+12))
	F_pfree(m, v17)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L9
	} else {
		goto L15
	}
L15:
	;
	goto L13
L16:
	;
	F_pfree(m, v4)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L9
	} else {
		goto L17
	}
L17:
	;
	if v7 != 0 {
		v4 = v7
		goto L4
	} else {
		goto L18
	}
L18:
	;
	goto L5
}
func F_pwrite(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v17 int32
	_ = v17
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = l1
	v17 = m.Wasi_snapshot_preview1.Fd_pwrite(m, l0, v8+int32(8), int32(1), l3, v8+int32(4))
	mBase = m.M
	if v17 == int32(0) {
		v24 = int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, _consts[43])) = v17
		v24 = int32(-1)
	}
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	m.G0 = v8 + int32(16)
	if v24 != 0 {
		v30 = int32(-1)
	} else {
		v30 = v25
	}
	return v30
}
func F_pwritev(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v12 = m.Wasi_snapshot_preview1.Fd_pwrite(m, l0, l1, l2, l3, v8+int32(12))
	mBase = m.M
	if v12 == int32(0) {
		v19 = int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, _consts[43])) = v12
		v19 = int32(-1)
	}
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	m.G0 = v8 + int32(16)
	if v19 != 0 {
		v25 = int32(-1)
	} else {
		v25 = v20
	}
	return v25
}
