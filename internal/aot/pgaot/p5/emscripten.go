package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F___emscripten_stdout_seek(m *base.Module, l0 int32, l1 int64, l2 int32) int64 {
	return int64(0)
}
func F_emscripten_builtin_malloc(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v139 int32
	_ = v139
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var __phi197 int32
	_ = __phi197
	var v199 int32
	_ = v199
	var __phi199 int32
	_ = __phi199
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v249 int32
	_ = v249
	var v253 int32
	_ = v253
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v300 int32
	_ = v300
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v316 int32
	_ = v316
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v334 int32
	_ = v334
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v346 int32
	_ = v346
	var v353 int32
	_ = v353
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v368 int32
	_ = v368
	var v378 int32
	_ = v378
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v384 int32
	_ = v384
	var v387 int32
	_ = v387
	var v390 int32
	_ = v390
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v399 int32
	_ = v399
	var __phi399 int32
	_ = __phi399
	var v400 int32
	_ = v400
	var __phi400 int32
	_ = __phi400
	var v411 int32
	_ = v411
	var v414 int32
	_ = v414
	var v421 int32
	_ = v421
	var v430 int32
	_ = v430
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v437 int32
	_ = v437
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v454 int32
	_ = v454
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v466 int32
	_ = v466
	var v469 int32
	_ = v469
	var v471 int32
	_ = v471
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v484 int32
	_ = v484
	var v486 int32
	_ = v486
	var v488 int32
	_ = v488
	var v490 int32
	_ = v490
	var v506 int32
	_ = v506
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v519 int32
	_ = v519
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v528 int32
	_ = v528
	var v534 int32
	_ = v534
	var v536 int32
	_ = v536
	var v548 int32
	_ = v548
	var v550 int32
	_ = v550
	var v553 int32
	_ = v553
	var v567 int32
	_ = v567
	var v571 int32
	_ = v571
	var v573 int32
	_ = v573
	var v581 int32
	_ = v581
	var v584 int32
	_ = v584
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v592 int32
	_ = v592
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v601 int32
	_ = v601
	var v607 int32
	_ = v607
	var v619 int32
	_ = v619
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v643 int32
	_ = v643
	var v645 int32
	_ = v645
	var v661 int32
	_ = v661
	var v663 int32
	_ = v663
	var v669 int32
	_ = v669
	var v675 int32
	_ = v675
	var v679 int32
	_ = v679
	var v685 int32
	_ = v685
	var v687 int32
	_ = v687
	var v688 int32
	_ = v688
	var v691 int32
	_ = v691
	var v696 int32
	_ = v696
	var v698 int32
	_ = v698
	var v710 int32
	_ = v710
	var v711 int32
	_ = v711
	var v714 int32
	_ = v714
	var v716 int32
	_ = v716
	var v719 int32
	_ = v719
	var v724 int32
	_ = v724
	var v734 int32
	_ = v734
	var v739 int32
	_ = v739
	var v752 int32
	_ = v752
	var v756 int32
	_ = v756
	var v762 int32
	_ = v762
	var v766 int32
	_ = v766
	var v767 int32
	_ = v767
	var v771 int32
	_ = v771
	var v772 int32
	_ = v772
	var v775 int32
	_ = v775
	var v785 int32
	_ = v785
	var v789 int32
	_ = v789
	var v798 int32
	_ = v798
	var v799 int32
	_ = v799
	var v801 int32
	_ = v801
	var v803 int32
	_ = v803
	var v804 int32
	_ = v804
	var v805 int32
	_ = v805
	var v815 int32
	_ = v815
	var v821 int32
	_ = v821
	var v827 int32
	_ = v827
	var v839 int32
	_ = v839
	var v841 int32
	_ = v841
	var v842 int32
	_ = v842
	var v849 int32
	_ = v849
	var v861 int32
	_ = v861
	var v863 int32
	_ = v863
	var v864 int32
	_ = v864
	var v867 int32
	_ = v867
	var v869 int32
	_ = v869
	var v870 int32
	_ = v870
	var v873 int32
	_ = v873
	var v874 int32
	_ = v874
	var v875 int32
	_ = v875
	var v878 int32
	_ = v878
	var v888 int32
	_ = v888
	var v896 int32
	_ = v896
	var v900 int32
	_ = v900
	var v903 int32
	_ = v903
	var v904 int64
	_ = v904
	var v906 int32
	_ = v906
	var v907 int64
	_ = v907
	var v922 int32
	_ = v922
	var v942 int32
	_ = v942
	var v946 int32
	_ = v946
	var v954 int32
	_ = v954
	var v956 int32
	_ = v956
	var v958 int32
	_ = v958
	var v962 int32
	_ = v962
	var v969 int32
	_ = v969
	var v970 int32
	_ = v970
	var v981 int32
	_ = v981
	var v984 int32
	_ = v984
	var v991 int32
	_ = v991
	var v996 int32
	_ = v996
	var v1000 int32
	_ = v1000
	var v1002 int32
	_ = v1002
	var v1018 int32
	_ = v1018
	var v1020 int32
	_ = v1020
	var v1021 int32
	_ = v1021
	var v1026 int32
	_ = v1026
	var v1033 int32
	_ = v1033
	var v1043 int32
	_ = v1043
	var v1044 int32
	_ = v1044
	var v1061 int32
	_ = v1061
	var v1068 int32
	_ = v1068
	var v1070 int32
	_ = v1070
	var v1071 int32
	_ = v1071
	var v1080 int32
	_ = v1080
	var v1098 int32
	_ = v1098
	var v1101 int32
	_ = v1101
	var v1103 int32
	_ = v1103
	var v1105 int32
	_ = v1105
	var v1106 int32
	_ = v1106
	var v1133 int32
	_ = v1133
	var v1136 int32
	_ = v1136
	var v1138 int32
	_ = v1138
	var v1140 int32
	_ = v1140
	var v1148 int32
	_ = v1148
	var v1149 int32
	_ = v1149
	var v1150 int32
	_ = v1150
	var v1152 int32
	_ = v1152
	var v1156 int32
	_ = v1156
	var v1158 int32
	_ = v1158
	var v1159 int32
	_ = v1159
	var v1165 int32
	_ = v1165
	var v1169 int32
	_ = v1169
	var v1171 int32
	_ = v1171
	var v1172 int32
	_ = v1172
	var v1179 int32
	_ = v1179
	var v1185 int32
	_ = v1185
	var v1186 int32
	_ = v1186
	var v1189 int32
	_ = v1189
	var v1191 int32
	_ = v1191
	var v1193 int32
	_ = v1193
	var v1202 int32
	_ = v1202
	var v1204 int32
	_ = v1204
	var v1207 int32
	_ = v1207
	var v1210 int32
	_ = v1210
	var v1215 int32
	_ = v1215
	var v1216 int32
	_ = v1216
	var v1217 int32
	_ = v1217
	var __phi1217 int32
	_ = __phi1217
	var v1222 int32
	_ = v1222
	var __phi1222 int32
	_ = __phi1222
	var v1231 int32
	_ = v1231
	var v1234 int32
	_ = v1234
	var v1242 int32
	_ = v1242
	var v1252 int32
	_ = v1252
	var v1254 int32
	_ = v1254
	var v1257 int32
	_ = v1257
	var v1260 int32
	_ = v1260
	var v1262 int32
	_ = v1262
	var v1267 int32
	_ = v1267
	var v1274 int32
	_ = v1274
	var v1277 int32
	_ = v1277
	var v1294 int32
	_ = v1294
	var v1295 int32
	_ = v1295
	var v1299 int32
	_ = v1299
	var v1302 int32
	_ = v1302
	var v1303 int32
	_ = v1303
	var v1320 int32
	_ = v1320
	var v1322 int32
	_ = v1322
	var v1324 int32
	_ = v1324
	var v1328 int32
	_ = v1328
	var v1335 int32
	_ = v1335
	var v1336 int32
	_ = v1336
	var v1347 int32
	_ = v1347
	var v1350 int32
	_ = v1350
	var v1357 int32
	_ = v1357
	var v1362 int32
	_ = v1362
	var v1366 int32
	_ = v1366
	var v1368 int32
	_ = v1368
	var v1384 int32
	_ = v1384
	var v1386 int32
	_ = v1386
	var v1387 int32
	_ = v1387
	var v1391 int32
	_ = v1391
	var v1399 int32
	_ = v1399
	var v1409 int32
	_ = v1409
	var v1410 int32
	_ = v1410
	var v1427 int32
	_ = v1427
	var v1448 int32
	_ = v1448
	var v1462 int32
	_ = v1462
	var v1464 int32
	_ = v1464
	var v1467 int32
	_ = v1467
	var v1473 int32
	_ = v1473
	var v1475 int32
	_ = v1475
	var v1482 int32
	_ = v1482
	var v1485 int32
	_ = v1485
	var v1492 int32
	_ = v1492
	var v1495 int32
	_ = v1495
	var v1499 int32
	_ = v1499
	var v1500 int32
	_ = v1500
	var v1507 int32
	_ = v1507
	var v1516 int32
	_ = v1516
	var v1518 int32
	_ = v1518
	var v1520 int32
	_ = v1520
	var v1524 int32
	_ = v1524
	var v1531 int32
	_ = v1531
	var v1532 int32
	_ = v1532
	var v1543 int32
	_ = v1543
	var v1546 int32
	_ = v1546
	var v1553 int32
	_ = v1553
	var v1558 int32
	_ = v1558
	var v1562 int32
	_ = v1562
	var v1578 int32
	_ = v1578
	var v1580 int32
	_ = v1580
	var v1581 int32
	_ = v1581
	var v1584 int32
	_ = v1584
	var v1593 int32
	_ = v1593
	var v1603 int32
	_ = v1603
	var v1604 int32
	_ = v1604
	var v1621 int32
	_ = v1621
	var v1642 int32
	_ = v1642
	var v1656 int32
	_ = v1656
	var v1658 int32
	_ = v1658
	var v1661 int32
	_ = v1661
	var v1669 int32
	_ = v1669
	var v1676 int32
	_ = v1676
	var v1679 int32
	_ = v1679
	var v1688 int32
	_ = v1688
	var v1692 int32
	_ = v1692
	var v1693 int32
	_ = v1693
	var v1700 int32
	_ = v1700
	var v1707 int32
	_ = v1707
	var v1709 int32
	_ = v1709
	var v1711 int32
	_ = v1711
	var v1715 int32
	_ = v1715
	var v1722 int32
	_ = v1722
	var v1723 int32
	_ = v1723
	var v1741 int32
	_ = v1741
	v2 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(16)
	m.G0 = v15
	if base.Ui32(l0) <= base.Ui32(int32(244)) {
		goto L11
	} else {
		goto L12
	}
L1:
	;
	m.G0 = v15 + int32(16)
	return v1741
L2:
	;
	if v178 == int32(0) {
		goto L361
	} else {
		goto L362
	}
L3:
	;
	if v381 == int32(0) {
		v1492 = v222
		goto L319
	} else {
		goto L320
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v827))) = v675
	v1133 = *(*int32)(unsafe.Add(mBase, uint32(v827)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v827)+4)) = v1133 + v679
	v1136 = int32(-8)
	v1138 = int32(7)
	v1140 = v675 + (v1136-v675)&v1138
	*(*int32)(unsafe.Add(mBase, uint32(v1140)+4)) = v421 | int32(3)
	v1148 = v839 + (v1136-v839)&v1138
	v1149 = v421 + v1140
	v1150 = v1148 - v1149
	v1152 = *(*int32)(unsafe.Add(mBase, _consts[1499]))
	if v1152 == v1148 {
		goto L253
	} else {
		goto L254
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, _consts[87])) = int32(48)
	v1741 = int32(0)
	goto L1
L6:
	;
	v1098 = *(*int32)(unsafe.Add(mBase, _consts[1500]))
	if base.Ui32(v1098) <= base.Ui32(v421) {
		goto L5
	} else {
		goto L251
	}
L7:
	;
	v821 = *(*int32)(unsafe.Add(mBase, _consts[1501]))
	if base.Ui32(v675) < base.Ui32(v821) {
		goto L203
	} else {
		goto L204
	}
L8:
	;
	v1448 = int32(0)
	goto L3
L9:
	;
	v1642 = int32(0)
	goto L2
L10:
	;
	v430 = *(*int32)(unsafe.Add(mBase, _consts[1502]))
	if base.Ui32(v421) <= base.Ui32(v430) {
		goto L120
	} else {
		goto L121
	}
L11:
	;
	v20 = *(*int32)(unsafe.Add(mBase, _consts[1503]))
	v22 = int32(11)
	if base.Ui32(l0) < base.Ui32(v22) {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	goto L13
L13:
	;
	if base.Ui32(int32(-65)) < base.Ui32(l0) {
		v421 = int32(-1)
		goto L10
	} else {
		goto L64
	}
L14:
	;
	v28 = int32(16)
	goto L16
L15:
	;
	v28 = (l0 + v22) & int32(504)
	goto L16
L16:
	;
	v29 = int32(3)
	v30 = int32(base.Ui32(v28) >> (uint(v29) % 32))
	v31 = int32(base.Ui32(v20) >> (uint(v30) % 32))
	if v31&v29 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v38 = (v31^int32(-1))&int32(1) + v30
	v40 = v38 << (uint(int32(3)) % 32)
	v42 = v40 + int32(4730976)
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v40)+uint32(_consts[1504])))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)+8))
	if v42 == v46 {
		goto L21
	} else {
		goto L22
	}
L18:
	;
	goto L19
L19:
	;
	v68 = *(*int32)(unsafe.Add(mBase, _consts[1502]))
	if base.Ui32(v28) <= base.Ui32(v68) {
		v421 = v28
		goto L10
	} else {
		goto L24
	}
L20:
	;
	v57 = int32(3)
	v58 = v38 << (uint(v57) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v45)+4)) = v58 | v57
	v62 = v45 + v58
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v62)+4)) = v63 | int32(1)
	v1741 = v45 + int32(8)
	goto L1
L21:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1503])) = v20 & base.I32_rotl(int32(-2), v38)
	goto L20
L22:
	;
	goto L23
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+12)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v40)+uint32(_consts[1504]))) = v46
	goto L20
L24:
	;
	if v31 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v72 = int32(2) << (uint(v30) % 32)
	v77 = base.I32_ctz(v31 << (uint(v30) % 32) & (v72 | (int32(0) - v72)))
	v79 = v77 << (uint(int32(3)) % 32)
	v81 = v79 + int32(4730976)
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v79)+uint32(_consts[1504])))
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v84)+8))
	if v81 == v85 {
		goto L29
	} else {
		goto L30
	}
L26:
	;
	goto L27
L27:
	;
	v139 = *(*int32)(unsafe.Add(mBase, _consts[1505]))
	if v139 == int32(0) {
		v421 = v28
		goto L10
	} else {
		goto L39
	}
L28:
	;
	v95 = int32(3)
	*(*int32)(unsafe.Add(mBase, uint32(v84)+4)) = v28 | v95
	v98 = v84 + v28
	v100 = v77 << (uint(v95) % 32)
	v101 = v100 - v28
	*(*int32)(unsafe.Add(mBase, uint32(v98)+4)) = v101 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v84+v100))) = v101
	if v68 != 0 {
		goto L32
	} else {
		goto L33
	}
L29:
	;
	v90 = v20 & base.I32_rotl(int32(-2), v77)
	*(*int32)(unsafe.Add(mBase, _consts[1503])) = v90
	v94 = v90
	goto L28
L30:
	;
	goto L31
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v85)+12)) = v81
	*(*int32)(unsafe.Add(mBase, uint32(v79)+uint32(_consts[1504]))) = v85
	v94 = v20
	goto L28
L32:
	;
	v108 = v68 & int32(-8)
	v110 = v108 + int32(4730976)
	v112 = *(*int32)(unsafe.Add(mBase, _consts[1506]))
	v116 = int32(1) << (uint(int32(base.Ui32(v68)>>(uint(int32(3))%32))) % 32)
	if v94&v116 == int32(0) {
		goto L36
	} else {
		goto L37
	}
L33:
	;
	goto L34
L34:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1506])) = v98
	*(*int32)(unsafe.Add(mBase, _consts[1502])) = v101
	v1741 = v84 + int32(8)
	goto L1
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v108)+uint32(_consts[1504]))) = v112
	*(*int32)(unsafe.Add(mBase, uint32(v124)+12)) = v112
	*(*int32)(unsafe.Add(mBase, uint32(v112)+12)) = v110
	*(*int32)(unsafe.Add(mBase, uint32(v112)+8)) = v124
	goto L34
L36:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1503])) = v116 | v94
	v124 = v110
	goto L35
L37:
	;
	goto L38
L38:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v108)+uint32(_consts[1504])))
	v124 = v123
	goto L35
L39:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(base.I32_ctz(v139)<<(uint(int32(2))%32))+uint32(_consts[1507])))
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v147)+4))
	v153 = v148&int32(-8) - v28
	v154 = v147
	v155 = v147
	goto L40
L40:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v155)+16))
	if v164 == int32(0) {
		goto L43
	} else {
		goto L44
	}
L41:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v154)+24))
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v154)+12))
	if v154 != v179 {
		goto L53
	} else {
		goto L54
	}
L42:
	;
	goto L41
L43:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v155)+20))
	if v167 == int32(0) {
		goto L42
	} else {
		goto L46
	}
L44:
	;
	v170 = v164
	goto L45
L45:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v170)+4))
	v174 = v171&int32(-8) - v28
	v175 = base.B2i32(base.Ui32(v174) < base.Ui32(v153))
	if base.Ui32(v174) < base.Ui32(v153) {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	v170 = v167
	goto L45
L47:
	;
	v176 = v174
	goto L49
L48:
	;
	v176 = v153
	goto L49
L49:
	;
	if base.Ui32(v174) < base.Ui32(v153) {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v177 = v170
	goto L52
L51:
	;
	v177 = v154
	goto L52
L52:
	;
	v153 = v176
	v154 = v177
	v155 = v170
	goto L40
L53:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v154)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v181)+12)) = v179
	*(*int32)(unsafe.Add(mBase, uint32(v179)+8)) = v181
	v1642 = v179
	goto L2
L54:
	;
	goto L55
L55:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v154)+20))
	if v184 != 0 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v192 = v184
	v193 = v154 + int32(20)
	goto L58
L57:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v154)+16))
	if v187 == int32(0) {
		goto L9
	} else {
		goto L59
	}
L58:
	;
	__phi197 = v192
	__phi199 = v193
	v197 = __phi197
	v199 = __phi199
	goto L60
L59:
	;
	v192 = v187
	v193 = v154 + int32(16)
	goto L58
L60:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v197)+20))
	if v208 != 0 {
		__phi197 = v208
		__phi199 = v197 + int32(20)
		v197 = __phi197
		v199 = __phi199
		goto L60
	} else {
		goto L62
	}
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v199))) = int32(0)
	v1642 = v197
	goto L2
L62:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v197)+16))
	if v211 != 0 {
		__phi197 = v211
		__phi199 = v197 + int32(16)
		v197 = __phi197
		v199 = __phi199
		goto L60
	} else {
		goto L63
	}
L63:
	;
	goto L61
L64:
	;
	v218 = l0 + int32(11)
	v220 = v218 & int32(-8)
	v222 = *(*int32)(unsafe.Add(mBase, _consts[1505]))
	if v222 == int32(0) {
		v421 = v220
		goto L10
	} else {
		goto L65
	}
L65:
	;
	if base.Ui32(l0) <= base.Ui32(int32(16777204)) {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v231 = base.I32_clz(int32(base.Ui32(v218) >> (uint(int32(8)) % 32)))
	v234 = int32(1)
	v242 = int32(base.Ui32(v220)>>(uint(int32(38)-v231)%32))&v234 - v231<<(uint(v234)%32) + int32(62)
	goto L68
L67:
	;
	v242 = int32(31)
	goto L68
L68:
	;
	v243 = int32(0)
	v244 = v243 - v220
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v242<<(uint(int32(2))%32))+uint32(_consts[1507])))
	if v249 == v243 {
		goto L72
	} else {
		goto L73
	}
L69:
	;
	if v368 == int32(0) {
		v421 = v220
		goto L10
	} else {
		goto L107
	}
L70:
	;
	v341 = v329
	v342 = v330
	v346 = v334
	goto L95
L71:
	;
	if v295|v300 == int32(0) {
		goto L90
	} else {
		goto L91
	}
L72:
	;
	v295 = int32(0)
	v296 = v244
	v300 = v2
	goto L71
L73:
	;
	goto L74
L74:
	;
	v253 = int32(0)
	if v242 != int32(31) {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v261 = int32(25) - int32(base.Ui32(v242)>>(uint(int32(1))%32))
	goto L77
L76:
	;
	v261 = v253
	goto L77
L77:
	;
	v263 = v253
	v264 = v244
	v265 = v220 << (uint(v261) % 32)
	v266 = v249
	v268 = v2
	goto L78
L78:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v266)+4))
	v278 = v275&int32(-8) - v220
	if base.Ui32(v264) <= base.Ui32(v278) {
		v281 = v264
		v282 = v268
		goto L80
	} else {
		goto L81
	}
L79:
	;
	v295 = v292
	v296 = v281
	v300 = v282
	goto L71
L80:
	;
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v266)+20))
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v266+int32(base.Ui32(v265)>>(uint(int32(29))%32))&int32(4))+16))
	if v283 == v289 {
		goto L83
	} else {
		goto L84
	}
L81:
	;
	if v278 != 0 {
		v281 = v278
		v282 = v266
		goto L80
	} else {
		goto L82
	}
L82:
	;
	v329 = v266
	v330 = int32(0)
	v334 = v266
	goto L70
L83:
	;
	v291 = v263
	goto L85
L84:
	;
	v291 = v283
	goto L85
L85:
	;
	if v283 != 0 {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v292 = v291
	goto L88
L87:
	;
	v292 = v263
	goto L88
L88:
	;
	if v289 != 0 {
		v263 = v292
		v264 = v281
		v265 = v265 << (uint(int32(1)) % 32)
		v266 = v289
		v268 = v282
		goto L78
	} else {
		goto L89
	}
L89:
	;
	goto L79
L90:
	;
	v310 = int32(0)
	v312 = int32(2) << (uint(v242) % 32)
	v316 = (v312 | (v310 - v312)) & v222
	if v316 == v310 {
		v421 = v220
		goto L10
	} else {
		goto L93
	}
L91:
	;
	v325 = v295
	v326 = v300
	goto L92
L92:
	;
	if v325 == int32(0) {
		v364 = v296
		v368 = v326
		goto L69
	} else {
		goto L94
	}
L93:
	;
	v324 = *(*int32)(unsafe.Add(mBase, uint32(base.I32_ctz(v316)<<(uint(int32(2))%32))+uint32(_consts[1507])))
	v325 = v324
	v326 = v310
	goto L92
L94:
	;
	v329 = v325
	v330 = v296
	v334 = v326
	goto L70
L95:
	;
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v341)+4))
	v356 = v353&int32(-8) - v220
	v357 = base.B2i32(base.Ui32(v356) < base.Ui32(v342))
	if base.Ui32(v356) < base.Ui32(v342) {
		goto L97
	} else {
		goto L98
	}
L96:
	;
	v364 = v358
	v368 = v359
	goto L69
L97:
	;
	v358 = v356
	goto L99
L98:
	;
	v358 = v342
	goto L99
L99:
	;
	if base.Ui32(v356) < base.Ui32(v342) {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	v359 = v341
	goto L102
L101:
	;
	v359 = v346
	goto L102
L102:
	;
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v341)+16))
	if v360 != 0 {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	v362 = v360
	goto L105
L104:
	;
	v361 = *(*int32)(unsafe.Add(mBase, uint32(v341)+20))
	v362 = v361
	goto L105
L105:
	;
	if v362 != 0 {
		v341 = v362
		v342 = v358
		v346 = v359
		goto L95
	} else {
		goto L106
	}
L106:
	;
	goto L96
L107:
	;
	v378 = *(*int32)(unsafe.Add(mBase, _consts[1502]))
	if base.Ui32(v378-v220) <= base.Ui32(v364) {
		v421 = v220
		goto L10
	} else {
		goto L108
	}
L108:
	;
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v368)+24))
	v382 = *(*int32)(unsafe.Add(mBase, uint32(v368)+12))
	if v368 != v382 {
		goto L109
	} else {
		goto L110
	}
L109:
	;
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v368)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v384)+12)) = v382
	*(*int32)(unsafe.Add(mBase, uint32(v382)+8)) = v384
	v1448 = v382
	goto L3
L110:
	;
	goto L111
L111:
	;
	v387 = *(*int32)(unsafe.Add(mBase, uint32(v368)+20))
	if v387 != 0 {
		goto L112
	} else {
		goto L113
	}
L112:
	;
	v395 = v387
	v396 = v368 + int32(20)
	goto L114
L113:
	;
	v390 = *(*int32)(unsafe.Add(mBase, uint32(v368)+16))
	if v390 == int32(0) {
		goto L8
	} else {
		goto L115
	}
L114:
	;
	__phi399 = v396
	__phi400 = v395
	v399 = __phi399
	v400 = __phi400
	goto L116
L115:
	;
	v395 = v390
	v396 = v368 + int32(16)
	goto L114
L116:
	;
	v411 = *(*int32)(unsafe.Add(mBase, uint32(v400)+20))
	if v411 != 0 {
		__phi399 = v400 + int32(20)
		__phi400 = v411
		v399 = __phi399
		v400 = __phi400
		goto L116
	} else {
		goto L118
	}
L117:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v399))) = int32(0)
	v1448 = v400
	goto L3
L118:
	;
	v414 = *(*int32)(unsafe.Add(mBase, uint32(v400)+16))
	if v414 != 0 {
		__phi399 = v400 + int32(16)
		__phi400 = v414
		v399 = __phi399
		v400 = __phi400
		goto L116
	} else {
		goto L119
	}
L119:
	;
	goto L117
L120:
	;
	v433 = *(*int32)(unsafe.Add(mBase, _consts[1506]))
	v434 = v430 - v421
	if base.Ui32(int32(16)) <= base.Ui32(v434) {
		goto L124
	} else {
		goto L125
	}
L121:
	;
	goto L122
L122:
	;
	v466 = *(*int32)(unsafe.Add(mBase, _consts[1500]))
	if base.Ui32(v421) < base.Ui32(v466) {
		goto L127
	} else {
		goto L128
	}
L123:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1502])) = v458
	*(*int32)(unsafe.Add(mBase, _consts[1506])) = v457
	v1741 = v433 + int32(8)
	goto L1
L124:
	;
	v437 = v433 + v421
	*(*int32)(unsafe.Add(mBase, uint32(v437)+4)) = v434 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v430+v433))) = v434
	*(*int32)(unsafe.Add(mBase, uint32(v433)+4)) = v421 | int32(3)
	v457 = v437
	v458 = v434
	goto L123
L125:
	;
	goto L126
L126:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v433)+4)) = v430 | int32(3)
	v449 = v430 + v433
	v450 = *(*int32)(unsafe.Add(mBase, uint32(v449)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v449)+4)) = v450 | int32(1)
	v454 = int32(0)
	v457 = v454
	v458 = v454
	goto L123
L127:
	;
	v469 = v466 - v421
	*(*int32)(unsafe.Add(mBase, _consts[1500])) = v469
	v471 = int32(4730960)
	v473 = *(*int32)(unsafe.Add(mBase, _consts[1499]))
	v474 = v473 + v421
	*(*int32)(unsafe.Add(mBase, _consts[1499])) = v474
	*(*int32)(unsafe.Add(mBase, uint32(v474)+4)) = v469 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v473)+4)) = v421 | int32(3)
	v1741 = v473 + int32(8)
	goto L1
L128:
	;
	goto L129
L129:
	;
	v484 = int32(0)
	v486 = v421 + int32(47)
	v488 = *(*int32)(unsafe.Add(mBase, _consts[1508]))
	if v488 != 0 {
		goto L131
	} else {
		goto L132
	}
L130:
	;
	v513 = v486 + v512
	v515 = int32(0) - v512
	v516 = v513 & v515
	if base.Ui32(v516) <= base.Ui32(v421) {
		v1741 = v484
		goto L1
	} else {
		goto L134
	}
L131:
	;
	v490 = *(*int32)(unsafe.Add(mBase, _consts[1509]))
	v512 = v490
	goto L130
L132:
	;
	goto L133
L133:
	;
	*(*int64)(unsafe.Add(mBase, _consts[1510])) = int64(-1)
	*(*int64)(unsafe.Add(mBase, _consts[1511])) = int64(17592186048512)
	*(*int32)(unsafe.Add(mBase, _consts[1508])) = (v15+int32(12))&int32(-16) ^ int32(1431655768)
	v506 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[1512])) = v506
	*(*int32)(unsafe.Add(mBase, _consts[1513])) = v506
	v512 = int32(4096)
	goto L130
L134:
	;
	v519 = *(*int32)(unsafe.Add(mBase, _consts[1514]))
	if v519 != 0 {
		goto L135
	} else {
		goto L136
	}
L135:
	;
	v521 = *(*int32)(unsafe.Add(mBase, _consts[1515]))
	v522 = v521 + v516
	if base.Ui32(v522) <= base.Ui32(v521) {
		v1741 = v484
		goto L1
	} else {
		goto L138
	}
L136:
	;
	goto L137
L137:
	;
	v528 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1513])))
	if v528&int32(4) == int32(0) {
		goto L141
	} else {
		goto L142
	}
L138:
	;
	if base.Ui32(v519) < base.Ui32(v522) {
		v1741 = v484
		goto L1
	} else {
		goto L139
	}
L139:
	;
	goto L137
L140:
	;
	v685 = int32(4731368)
	v687 = *(*int32)(unsafe.Add(mBase, _consts[1515]))
	v688 = v687 + v679
	*(*int32)(unsafe.Add(mBase, _consts[1515])) = v688
	v691 = *(*int32)(unsafe.Add(mBase, _consts[1516]))
	if base.Ui32(v691) < base.Ui32(v688) {
		goto L180
	} else {
		goto L181
	}
L141:
	;
	v534 = *(*int32)(unsafe.Add(mBase, _consts[1499]))
	if v534 != 0 {
		goto L148
	} else {
		goto L149
	}
L142:
	;
	goto L143
L143:
	;
	v661 = F_sbrk(m, v516)
	mBase = m.M
	v663 = F_sbrk(m, int32(0))
	mBase = m.M
	if v661 == int32(-1) {
		goto L5
	} else {
		goto L176
	}
L144:
	;
	v643 = int32(4731380)
	v645 = *(*int32)(unsafe.Add(mBase, _consts[1513]))
	*(*int32)(unsafe.Add(mBase, _consts[1513])) = v645 | int32(4)
	goto L143
L145:
	;
	if v596 != int32(-1) {
		v675 = v596
		v679 = v595
		goto L140
	} else {
		goto L175
	}
L146:
	;
	if v601 == int32(-1) {
		goto L144
	} else {
		goto L170
	}
L147:
	;
	v595 = (v513 - v466) & v515
	v596 = F_sbrk(m, v595)
	mBase = m.M
	v597 = *(*int32)(unsafe.Add(mBase, uint32(v536)))
	v598 = *(*int32)(unsafe.Add(mBase, uint32(v536)+4))
	if v596 == v597+v598 {
		goto L145
	} else {
		goto L169
	}
L148:
	;
	v536 = int32(4731384)
	goto L151
L149:
	;
	goto L150
L150:
	;
	v567 = F_sbrk(m, int32(0))
	mBase = m.M
	if v567 == int32(-1) {
		goto L144
	} else {
		goto L158
	}
L151:
	;
	v548 = *(*int32)(unsafe.Add(mBase, uint32(v536)))
	if base.Ui32(v548) <= base.Ui32(v534) {
		goto L153
	} else {
		goto L154
	}
L152:
	;
	goto L150
L153:
	;
	v550 = *(*int32)(unsafe.Add(mBase, uint32(v536)+4))
	if base.Ui32(v534) < base.Ui32(v548+v550) {
		goto L147
	} else {
		goto L156
	}
L154:
	;
	goto L155
L155:
	;
	v553 = *(*int32)(unsafe.Add(mBase, uint32(v536)+8))
	if v553 != 0 {
		v536 = v553
		goto L151
	} else {
		goto L157
	}
L156:
	;
	goto L155
L157:
	;
	goto L152
L158:
	;
	v571 = *(*int32)(unsafe.Add(mBase, _consts[1511]))
	v573 = v571 - int32(1)
	if v573&v567 != 0 {
		goto L159
	} else {
		goto L160
	}
L159:
	;
	v581 = v516 - v567 + (v573+v567)&(int32(0)-v571)
	goto L161
L160:
	;
	v581 = v516
	goto L161
L161:
	;
	if base.Ui32(v581) <= base.Ui32(v421) {
		goto L144
	} else {
		goto L162
	}
L162:
	;
	v584 = *(*int32)(unsafe.Add(mBase, _consts[1514]))
	if v584 != 0 {
		goto L163
	} else {
		goto L164
	}
L163:
	;
	v586 = *(*int32)(unsafe.Add(mBase, _consts[1515]))
	v587 = v586 + v581
	if base.Ui32(v587) <= base.Ui32(v586) {
		goto L144
	} else {
		goto L166
	}
L164:
	;
	goto L165
L165:
	;
	v592 = F_sbrk(m, v581)
	mBase = m.M
	if v592 != v567 {
		v601 = v592
		v607 = v581
		goto L146
	} else {
		goto L168
	}
L166:
	;
	if base.Ui32(v584) < base.Ui32(v587) {
		goto L144
	} else {
		goto L167
	}
L167:
	;
	goto L165
L168:
	;
	v675 = v567
	v679 = v581
	goto L140
L169:
	;
	v601 = v596
	v607 = v595
	goto L146
L170:
	;
	if base.Ui32(v421+int32(48)) <= base.Ui32(v607) {
		goto L171
	} else {
		goto L172
	}
L171:
	;
	v675 = v601
	v679 = v607
	goto L140
L172:
	;
	goto L173
L173:
	;
	v619 = *(*int32)(unsafe.Add(mBase, _consts[1509]))
	v624 = (v619 + (v486 - v607)) & (int32(0) - v619)
	v625 = F_sbrk(m, v624)
	mBase = m.M
	if v625 == int32(-1) {
		goto L144
	} else {
		goto L174
	}
L174:
	;
	v675 = v601
	v679 = v624 + v607
	goto L140
L175:
	;
	goto L144
L176:
	;
	if v663 == int32(-1) {
		goto L5
	} else {
		goto L177
	}
L177:
	;
	if base.Ui32(v663) <= base.Ui32(v661) {
		goto L5
	} else {
		goto L178
	}
L178:
	;
	v669 = v663 - v661
	if base.Ui32(v669) <= base.Ui32(v421+int32(40)) {
		goto L5
	} else {
		goto L179
	}
L179:
	;
	v675 = v661
	v679 = v669
	goto L140
L180:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1516])) = v688
	goto L182
L181:
	;
	goto L182
L182:
	;
	v696 = *(*int32)(unsafe.Add(mBase, _consts[1499]))
	if v696 != 0 {
		goto L184
	} else {
		goto L185
	}
L183:
	;
	if base.Ui32(v675) <= base.Ui32(v696) {
		goto L7
	} else {
		goto L200
	}
L184:
	;
	v698 = int32(4731384)
	goto L187
L185:
	;
	goto L186
L186:
	;
	v716 = *(*int32)(unsafe.Add(mBase, _consts[1501]))
	if base.Ui32(v716) <= base.Ui32(v675) {
		goto L191
	} else {
		goto L192
	}
L187:
	;
	v710 = *(*int32)(unsafe.Add(mBase, uint32(v698)))
	v711 = *(*int32)(unsafe.Add(mBase, uint32(v698)+4))
	if v675 == v710+v711 {
		goto L183
	} else {
		goto L189
	}
L188:
	;
	goto L7
L189:
	;
	v714 = *(*int32)(unsafe.Add(mBase, uint32(v698)+8))
	if v714 != 0 {
		v698 = v714
		goto L187
	} else {
		goto L190
	}
L190:
	;
	goto L188
L191:
	;
	v719 = v716
	goto L193
L192:
	;
	v719 = int32(0)
	goto L193
L193:
	;
	if v719 == int32(0) {
		goto L194
	} else {
		goto L195
	}
L194:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1501])) = v675
	goto L196
L195:
	;
	goto L196
L196:
	;
	v724 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[1517])) = v679
	*(*int32)(unsafe.Add(mBase, _consts[1518])) = v675
	*(*int32)(unsafe.Add(mBase, _consts[1519])) = int32(-1)
	v734 = *(*int32)(unsafe.Add(mBase, _consts[1508]))
	*(*int32)(unsafe.Add(mBase, _consts[1520])) = v734
	*(*int32)(unsafe.Add(mBase, _consts[1521])) = v724
	v739 = v724
	goto L197
L197:
	;
	v752 = v739 << (uint(int32(3)) % 32)
	v756 = v752 + int32(4730976)
	*(*int32)(unsafe.Add(mBase, uint32(v752)+uint32(_consts[1504]))) = v756
	*(*int32)(unsafe.Add(mBase, uint32(v752)+uint32(_consts[1522]))) = v756
	v762 = v739 + int32(1)
	if v762 != int32(32) {
		v739 = v762
		goto L197
	} else {
		goto L199
	}
L198:
	;
	v766 = int32(40)
	v767 = v679 - v766
	v771 = (int32(-8) - v675) & int32(7)
	v772 = v767 - v771
	*(*int32)(unsafe.Add(mBase, _consts[1500])) = v772
	v775 = v771 + v675
	*(*int32)(unsafe.Add(mBase, _consts[1499])) = v775
	*(*int32)(unsafe.Add(mBase, uint32(v775)+4)) = v772 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v767+v675)+4)) = v766
	v785 = *(*int32)(unsafe.Add(mBase, _consts[1523]))
	*(*int32)(unsafe.Add(mBase, _consts[1524])) = v785
	goto L6
L199:
	;
	goto L198
L200:
	;
	if base.Ui32(v696) < base.Ui32(v710) {
		goto L7
	} else {
		goto L201
	}
L201:
	;
	v789 = *(*int32)(unsafe.Add(mBase, uint32(v698)+12))
	if v789&int32(8) != 0 {
		goto L7
	} else {
		goto L202
	}
L202:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v698)+4)) = v711 + v679
	v798 = (int32(-8) - v696) & int32(7)
	v799 = v696 + v798
	*(*int32)(unsafe.Add(mBase, _consts[1499])) = v799
	v801 = int32(4730948)
	v803 = *(*int32)(unsafe.Add(mBase, _consts[1500]))
	v804 = v803 + v679
	v805 = v804 - v798
	*(*int32)(unsafe.Add(mBase, _consts[1500])) = v805
	*(*int32)(unsafe.Add(mBase, uint32(v799)+4)) = v805 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v696+v804)+4)) = int32(40)
	v815 = *(*int32)(unsafe.Add(mBase, _consts[1523]))
	*(*int32)(unsafe.Add(mBase, _consts[1524])) = v815
	goto L6
L203:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1501])) = v675
	goto L205
L204:
	;
	goto L205
L205:
	;
	v827 = int32(4731384)
	goto L207
L206:
	;
	v849 = int32(4731384)
	goto L214
L207:
	;
	v839 = *(*int32)(unsafe.Add(mBase, uint32(v827)))
	if v675+v679 != v839 {
		goto L209
	} else {
		goto L210
	}
L208:
	;
	v842 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v827)+12)))
	if v842&int32(8) == int32(0) {
		goto L4
	} else {
		goto L213
	}
L209:
	;
	v841 = *(*int32)(unsafe.Add(mBase, uint32(v827)+8))
	if v841 != 0 {
		v827 = v841
		goto L207
	} else {
		goto L212
	}
L210:
	;
	goto L211
L211:
	;
	goto L208
L212:
	;
	goto L206
L213:
	;
	goto L206
L214:
	;
	v861 = *(*int32)(unsafe.Add(mBase, uint32(v849)))
	if base.Ui32(v861) <= base.Ui32(v696) {
		goto L217
	} else {
		goto L218
	}
L215:
	;
	v869 = int32(40)
	v870 = v679 - v869
	v873 = int32(7)
	v874 = (int32(-8) - v675) & v873
	v875 = v870 - v874
	*(*int32)(unsafe.Add(mBase, _consts[1500])) = v875
	v878 = v675 + v874
	*(*int32)(unsafe.Add(mBase, _consts[1499])) = v878
	*(*int32)(unsafe.Add(mBase, uint32(v878)+4)) = v875 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v870+v675)+4)) = v869
	v888 = *(*int32)(unsafe.Add(mBase, _consts[1523]))
	*(*int32)(unsafe.Add(mBase, _consts[1524])) = v888
	v896 = v864 + (int32(39)-v864)&v873 - int32(47)
	if base.Ui32(v896) < base.Ui32(v696+int32(16)) {
		goto L221
	} else {
		goto L222
	}
L216:
	;
	goto L215
L217:
	;
	v863 = *(*int32)(unsafe.Add(mBase, uint32(v849)+4))
	v864 = v861 + v863
	if base.Ui32(v696) < base.Ui32(v864) {
		goto L216
	} else {
		goto L220
	}
L218:
	;
	goto L219
L219:
	;
	v867 = *(*int32)(unsafe.Add(mBase, uint32(v849)+8))
	v849 = v867
	goto L214
L220:
	;
	goto L219
L221:
	;
	v900 = v696
	goto L223
L222:
	;
	v900 = v896
	goto L223
L223:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v900)+4)) = int32(27)
	v903 = int32(4731392)
	v904 = *(*int64)(unsafe.Add(mBase, _consts[1525]))
	*(*int64)(unsafe.Add(mBase, uint32(v900)+16)) = v904
	v906 = int32(4731384)
	v907 = *(*int64)(unsafe.Add(mBase, _consts[1518]))
	*(*int64)(unsafe.Add(mBase, uint32(v900)+8)) = v907
	*(*int32)(unsafe.Add(mBase, _consts[1525])) = v900 + int32(8)
	*(*int32)(unsafe.Add(mBase, _consts[1517])) = v679
	*(*int32)(unsafe.Add(mBase, _consts[1518])) = v675
	*(*int32)(unsafe.Add(mBase, _consts[1521])) = int32(0)
	v922 = v900 + int32(24)
	goto L224
L224:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v922)+4)) = int32(7)
	if base.Ui32(v922+int32(8)) < base.Ui32(v864) {
		v922 = v922 + int32(4)
		goto L224
	} else {
		goto L226
	}
L225:
	;
	if v696 == v900 {
		goto L6
	} else {
		goto L227
	}
L226:
	;
	goto L225
L227:
	;
	v942 = *(*int32)(unsafe.Add(mBase, uint32(v900)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v900)+4)) = v942 & int32(-2)
	v946 = v900 - v696
	*(*int32)(unsafe.Add(mBase, uint32(v696)+4)) = v946 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v900))) = v946
	if base.Ui32(v946) <= base.Ui32(int32(255)) {
		goto L229
	} else {
		goto L230
	}
L228:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1080+v696))) = v1071
	*(*int32)(unsafe.Add(mBase, uint32(v696+v1070))) = v1068
	goto L6
L229:
	;
	v954 = v946 & int32(-8)
	v956 = v954 + int32(4730976)
	v958 = *(*int32)(unsafe.Add(mBase, _consts[1503]))
	v962 = int32(1) << (uint(int32(base.Ui32(v946)>>(uint(int32(3))%32))) % 32)
	if v958&v962 == int32(0) {
		goto L233
	} else {
		goto L234
	}
L230:
	;
	goto L231
L231:
	;
	if base.Ui32(v946) <= base.Ui32(int32(16777215)) {
		goto L236
	} else {
		goto L237
	}
L232:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v954)+uint32(_consts[1504]))) = v696
	*(*int32)(unsafe.Add(mBase, uint32(v970)+12)) = v696
	v1068 = v956
	v1070 = int32(12)
	v1071 = v970
	v1080 = int32(8)
	goto L228
L233:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1503])) = v962 | v958
	v970 = v956
	goto L232
L234:
	;
	goto L235
L235:
	;
	v969 = *(*int32)(unsafe.Add(mBase, uint32(v954)+uint32(_consts[1504])))
	v970 = v969
	goto L232
L236:
	;
	v981 = base.I32_clz(int32(base.Ui32(v946) >> (uint(int32(8)) % 32)))
	v984 = int32(1)
	v991 = int32(base.Ui32(v946)>>(uint(int32(38)-v981)%32))&v984 - v981<<(uint(v984)%32) + int32(62)
	goto L238
L237:
	;
	v991 = int32(31)
	goto L238
L238:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v696)+28)) = v991
	*(*int64)(unsafe.Add(mBase, uint32(v696)+16)) = int64(0)
	v996 = v991 << (uint(int32(2)) % 32)
	v1000 = *(*int32)(unsafe.Add(mBase, _consts[1505]))
	v1002 = int32(1) << (uint(v991) % 32)
	if v1000&v1002 == int32(0) {
		goto L241
	} else {
		goto L242
	}
L239:
	;
	v1061 = *(*int32)(unsafe.Add(mBase, uint32(v1026)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1061)+12)) = v696
	*(*int32)(unsafe.Add(mBase, uint32(v1026)+8)) = v696
	*(*int32)(unsafe.Add(mBase, uint32(v696)+8)) = v1061
	v1068 = int32(0)
	v1070 = int32(24)
	v1071 = v1026
	v1080 = int32(12)
	goto L228
L240:
	;
	v1068 = v696
	v1070 = int32(8)
	v1071 = v696
	v1080 = int32(12)
	goto L228
L241:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1505])) = v1000 | v1002
	*(*int32)(unsafe.Add(mBase, uint32(v996)+uint32(_consts[1507]))) = v696
	*(*int32)(unsafe.Add(mBase, uint32(v696)+24)) = v996 + int32(4731240)
	goto L240
L242:
	;
	goto L243
L243:
	;
	if v991 != int32(31) {
		goto L244
	} else {
		goto L245
	}
L244:
	;
	v1018 = int32(25) - int32(base.Ui32(v991)>>(uint(int32(1))%32))
	goto L246
L245:
	;
	v1018 = int32(0)
	goto L246
L246:
	;
	v1020 = *(*int32)(unsafe.Add(mBase, uint32(v996)+uint32(_consts[1507])))
	v1021 = v946 << (uint(v1018) % 32)
	v1026 = v1020
	goto L247
L247:
	;
	v1033 = *(*int32)(unsafe.Add(mBase, uint32(v1026)+4))
	if v1033&int32(-8) == v946 {
		goto L239
	} else {
		goto L249
	}
L248:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1043)+16)) = v696
	*(*int32)(unsafe.Add(mBase, uint32(v696)+24)) = v1026
	goto L240
L249:
	;
	v1043 = v1026 + int32(base.Ui32(v1021)>>(uint(int32(29))%32))&int32(4)
	v1044 = *(*int32)(unsafe.Add(mBase, uint32(v1043)+16))
	if v1044 != 0 {
		v1021 = v1021 << (uint(int32(1)) % 32)
		v1026 = v1044
		goto L247
	} else {
		goto L250
	}
L250:
	;
	goto L248
L251:
	;
	v1101 = v1098 - v421
	*(*int32)(unsafe.Add(mBase, _consts[1500])) = v1101
	v1103 = int32(4730960)
	v1105 = *(*int32)(unsafe.Add(mBase, _consts[1499]))
	v1106 = v1105 + v421
	*(*int32)(unsafe.Add(mBase, _consts[1499])) = v1106
	*(*int32)(unsafe.Add(mBase, uint32(v1106)+4)) = v1101 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1105)+4)) = v421 | int32(3)
	v1741 = v1105 + int32(8)
	goto L1
L252:
	;
	v1741 = v1140 + int32(8)
	goto L1
L253:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1499])) = v1149
	v1156 = int32(4730948)
	v1158 = *(*int32)(unsafe.Add(mBase, _consts[1500]))
	v1159 = v1158 + v1150
	*(*int32)(unsafe.Add(mBase, _consts[1500])) = v1159
	*(*int32)(unsafe.Add(mBase, uint32(v1149)+4)) = v1159 | int32(1)
	goto L252
L254:
	;
	goto L255
L255:
	;
	v1165 = *(*int32)(unsafe.Add(mBase, _consts[1506]))
	if v1165 == v1148 {
		goto L256
	} else {
		goto L257
	}
L256:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1506])) = v1149
	v1169 = int32(4730944)
	v1171 = *(*int32)(unsafe.Add(mBase, _consts[1502]))
	v1172 = v1171 + v1150
	*(*int32)(unsafe.Add(mBase, _consts[1502])) = v1172
	*(*int32)(unsafe.Add(mBase, uint32(v1149)+4)) = v1172 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1149+v1172))) = v1172
	goto L252
L257:
	;
	goto L258
L258:
	;
	v1179 = *(*int32)(unsafe.Add(mBase, uint32(v1148)+4))
	if v1179&int32(3) == int32(1) {
		goto L259
	} else {
		goto L260
	}
L259:
	;
	v1185 = v1179 & int32(-8)
	v1186 = *(*int32)(unsafe.Add(mBase, uint32(v1148)+12))
	if base.Ui32(v1179) <= base.Ui32(int32(255)) {
		goto L263
	} else {
		goto L264
	}
L260:
	;
	v1299 = v1150
	v1302 = v1179
	v1303 = v1148
	goto L261
L261:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1303)+4)) = v1302 & int32(-2)
	*(*int32)(unsafe.Add(mBase, uint32(v1149)+4)) = v1299 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1149+v1299))) = v1299
	if base.Ui32(v1299) <= base.Ui32(int32(255)) {
		goto L297
	} else {
		goto L298
	}
L262:
	;
	v1294 = v1148 + v1185
	v1295 = *(*int32)(unsafe.Add(mBase, uint32(v1294)+4))
	v1299 = v1150 + v1185
	v1302 = v1295
	v1303 = v1294
	goto L261
L263:
	;
	v1189 = *(*int32)(unsafe.Add(mBase, uint32(v1148)+8))
	if v1189 == v1186 {
		goto L266
	} else {
		goto L267
	}
L264:
	;
	goto L265
L265:
	;
	v1202 = *(*int32)(unsafe.Add(mBase, uint32(v1148)+24))
	if v1186 != v1148 {
		goto L270
	} else {
		goto L271
	}
L266:
	;
	v1191 = int32(4730936)
	v1193 = *(*int32)(unsafe.Add(mBase, _consts[1503]))
	*(*int32)(unsafe.Add(mBase, _consts[1503])) = v1193 & base.I32_rotl(int32(-2), int32(base.Ui32(v1179)>>(uint(int32(3))%32)))
	goto L262
L267:
	;
	goto L268
L268:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1189)+12)) = v1186
	*(*int32)(unsafe.Add(mBase, uint32(v1186)+8)) = v1189
	goto L262
L269:
	;
	if v1202 == int32(0) {
		goto L262
	} else {
		goto L282
	}
L270:
	;
	v1204 = *(*int32)(unsafe.Add(mBase, uint32(v1148)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1204)+12)) = v1186
	*(*int32)(unsafe.Add(mBase, uint32(v1186)+8)) = v1204
	v1242 = v1186
	goto L269
L271:
	;
	goto L272
L272:
	;
	v1207 = *(*int32)(unsafe.Add(mBase, uint32(v1148)+20))
	if v1207 != 0 {
		goto L274
	} else {
		goto L275
	}
L273:
	;
	v1242 = int32(0)
	goto L269
L274:
	;
	v1215 = v1207
	v1216 = v1148 + int32(20)
	goto L276
L275:
	;
	v1210 = *(*int32)(unsafe.Add(mBase, uint32(v1148)+16))
	if v1210 == int32(0) {
		goto L273
	} else {
		goto L277
	}
L276:
	;
	__phi1217 = v1216
	__phi1222 = v1215
	v1217 = __phi1217
	v1222 = __phi1222
	goto L278
L277:
	;
	v1215 = v1210
	v1216 = v1148 + int32(16)
	goto L276
L278:
	;
	v1231 = *(*int32)(unsafe.Add(mBase, uint32(v1222)+20))
	if v1231 != 0 {
		__phi1217 = v1222 + int32(20)
		__phi1222 = v1231
		v1217 = __phi1217
		v1222 = __phi1222
		goto L278
	} else {
		goto L280
	}
L279:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1217))) = int32(0)
	v1242 = v1222
	goto L269
L280:
	;
	v1234 = *(*int32)(unsafe.Add(mBase, uint32(v1222)+16))
	if v1234 != 0 {
		__phi1217 = v1222 + int32(16)
		__phi1222 = v1234
		v1217 = __phi1217
		v1222 = __phi1222
		goto L278
	} else {
		goto L281
	}
L281:
	;
	goto L279
L282:
	;
	v1252 = *(*int32)(unsafe.Add(mBase, uint32(v1148)+28))
	v1254 = v1252 << (uint(int32(2)) % 32)
	v1257 = *(*int32)(unsafe.Add(mBase, uint32(v1254)+uint32(_consts[1507])))
	if v1257 == v1148 {
		goto L284
	} else {
		goto L285
	}
L283:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1242)+24)) = v1202
	v1274 = *(*int32)(unsafe.Add(mBase, uint32(v1148)+16))
	if v1274 != 0 {
		goto L293
	} else {
		goto L294
	}
L284:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1254)+uint32(_consts[1507]))) = v1242
	if v1242 != 0 {
		goto L283
	} else {
		goto L287
	}
L285:
	;
	goto L286
L286:
	;
	v1267 = *(*int32)(unsafe.Add(mBase, uint32(v1202)+16))
	if v1148 == v1267 {
		goto L289
	} else {
		goto L290
	}
L287:
	;
	v1260 = int32(4730940)
	v1262 = *(*int32)(unsafe.Add(mBase, _consts[1505]))
	*(*int32)(unsafe.Add(mBase, _consts[1505])) = v1262 & base.I32_rotl(int32(-2), v1252)
	goto L262
L288:
	;
	if v1242 == int32(0) {
		goto L262
	} else {
		goto L292
	}
L289:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1202)+16)) = v1242
	goto L288
L290:
	;
	goto L291
L291:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1202)+20)) = v1242
	goto L288
L292:
	;
	goto L283
L293:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1242)+16)) = v1274
	*(*int32)(unsafe.Add(mBase, uint32(v1274)+24)) = v1242
	goto L295
L294:
	;
	goto L295
L295:
	;
	v1277 = *(*int32)(unsafe.Add(mBase, uint32(v1148)+20))
	if v1277 == int32(0) {
		goto L262
	} else {
		goto L296
	}
L296:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1242)+20)) = v1277
	*(*int32)(unsafe.Add(mBase, uint32(v1277)+24)) = v1242
	goto L262
L297:
	;
	v1320 = v1299 & int32(-8)
	v1322 = v1320 + int32(4730976)
	v1324 = *(*int32)(unsafe.Add(mBase, _consts[1503]))
	v1328 = int32(1) << (uint(int32(base.Ui32(v1299)>>(uint(int32(3))%32))) % 32)
	if v1324&v1328 == int32(0) {
		goto L301
	} else {
		goto L302
	}
L298:
	;
	goto L299
L299:
	;
	if base.Ui32(v1299) <= base.Ui32(int32(16777215)) {
		goto L304
	} else {
		goto L305
	}
L300:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1320)+uint32(_consts[1504]))) = v1149
	*(*int32)(unsafe.Add(mBase, uint32(v1336)+12)) = v1149
	*(*int32)(unsafe.Add(mBase, uint32(v1149)+12)) = v1322
	*(*int32)(unsafe.Add(mBase, uint32(v1149)+8)) = v1336
	goto L252
L301:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1503])) = v1328 | v1324
	v1336 = v1322
	goto L300
L302:
	;
	goto L303
L303:
	;
	v1335 = *(*int32)(unsafe.Add(mBase, uint32(v1320)+uint32(_consts[1504])))
	v1336 = v1335
	goto L300
L304:
	;
	v1347 = base.I32_clz(int32(base.Ui32(v1299) >> (uint(int32(8)) % 32)))
	v1350 = int32(1)
	v1357 = int32(base.Ui32(v1299)>>(uint(int32(38)-v1347)%32))&v1350 - v1347<<(uint(v1350)%32) + int32(62)
	goto L306
L305:
	;
	v1357 = int32(31)
	goto L306
L306:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1149)+28)) = v1357
	*(*int64)(unsafe.Add(mBase, uint32(v1149)+16)) = int64(0)
	v1362 = v1357 << (uint(int32(2)) % 32)
	v1366 = *(*int32)(unsafe.Add(mBase, _consts[1505]))
	v1368 = int32(1) << (uint(v1357) % 32)
	if v1366&v1368 == int32(0) {
		goto L309
	} else {
		goto L310
	}
L307:
	;
	v1427 = *(*int32)(unsafe.Add(mBase, uint32(v1387)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1427)+12)) = v1149
	*(*int32)(unsafe.Add(mBase, uint32(v1387)+8)) = v1149
	*(*int32)(unsafe.Add(mBase, uint32(v1149)+24)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1149)+12)) = v1387
	*(*int32)(unsafe.Add(mBase, uint32(v1149)+8)) = v1427
	goto L252
L308:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1149)+12)) = v1149
	*(*int32)(unsafe.Add(mBase, uint32(v1149)+8)) = v1149
	goto L252
L309:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1505])) = v1366 | v1368
	*(*int32)(unsafe.Add(mBase, uint32(v1362)+uint32(_consts[1507]))) = v1149
	*(*int32)(unsafe.Add(mBase, uint32(v1149)+24)) = v1362 + int32(4731240)
	goto L308
L310:
	;
	goto L311
L311:
	;
	if v1357 != int32(31) {
		goto L312
	} else {
		goto L313
	}
L312:
	;
	v1384 = int32(25) - int32(base.Ui32(v1357)>>(uint(int32(1))%32))
	goto L314
L313:
	;
	v1384 = int32(0)
	goto L314
L314:
	;
	v1386 = *(*int32)(unsafe.Add(mBase, uint32(v1362)+uint32(_consts[1507])))
	v1387 = v1386
	v1391 = v1299 << (uint(v1384) % 32)
	goto L315
L315:
	;
	v1399 = *(*int32)(unsafe.Add(mBase, uint32(v1387)+4))
	if v1399&int32(-8) == v1299 {
		goto L307
	} else {
		goto L317
	}
L316:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1409)+16)) = v1149
	*(*int32)(unsafe.Add(mBase, uint32(v1149)+24)) = v1387
	goto L308
L317:
	;
	v1409 = v1387 + int32(base.Ui32(v1391)>>(uint(int32(29))%32))&int32(4)
	v1410 = *(*int32)(unsafe.Add(mBase, uint32(v1409)+16))
	if v1410 != 0 {
		v1387 = v1410
		v1391 = v1391 << (uint(int32(1)) % 32)
		goto L315
	} else {
		goto L318
	}
L318:
	;
	goto L316
L319:
	;
	if base.Ui32(v364) <= base.Ui32(int32(15)) {
		goto L336
	} else {
		goto L337
	}
L320:
	;
	v1462 = *(*int32)(unsafe.Add(mBase, uint32(v368)+28))
	v1464 = v1462 << (uint(int32(2)) % 32)
	v1467 = *(*int32)(unsafe.Add(mBase, uint32(v1464)+uint32(_consts[1507])))
	if v1467 == v368 {
		goto L322
	} else {
		goto L323
	}
L321:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1448)+24)) = v381
	v1482 = *(*int32)(unsafe.Add(mBase, uint32(v368)+16))
	if v1482 != 0 {
		goto L331
	} else {
		goto L332
	}
L322:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1464)+uint32(_consts[1507]))) = v1448
	if v1448 != 0 {
		goto L321
	} else {
		goto L325
	}
L323:
	;
	goto L324
L324:
	;
	v1475 = *(*int32)(unsafe.Add(mBase, uint32(v381)+16))
	if v368 == v1475 {
		goto L327
	} else {
		goto L328
	}
L325:
	;
	v1473 = v222 & base.I32_rotl(int32(-2), v1462)
	*(*int32)(unsafe.Add(mBase, _consts[1505])) = v1473
	v1492 = v1473
	goto L319
L326:
	;
	if v1448 == int32(0) {
		v1492 = v222
		goto L319
	} else {
		goto L330
	}
L327:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v381)+16)) = v1448
	goto L326
L328:
	;
	goto L329
L329:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v381)+20)) = v1448
	goto L326
L330:
	;
	goto L321
L331:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1448)+16)) = v1482
	*(*int32)(unsafe.Add(mBase, uint32(v1482)+24)) = v1448
	goto L333
L332:
	;
	goto L333
L333:
	;
	v1485 = *(*int32)(unsafe.Add(mBase, uint32(v368)+20))
	if v1485 == int32(0) {
		v1492 = v222
		goto L319
	} else {
		goto L334
	}
L334:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1448)+20)) = v1485
	*(*int32)(unsafe.Add(mBase, uint32(v1485)+24)) = v1448
	v1492 = v222
	goto L319
L335:
	;
	v1741 = v368 + int32(8)
	goto L1
L336:
	;
	v1495 = v364 + v220
	*(*int32)(unsafe.Add(mBase, uint32(v368)+4)) = v1495 | int32(3)
	v1499 = v1495 + v368
	v1500 = *(*int32)(unsafe.Add(mBase, uint32(v1499)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1499)+4)) = v1500 | int32(1)
	goto L335
L337:
	;
	goto L338
L338:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v368)+4)) = v220 | int32(3)
	v1507 = v220 + v368
	*(*int32)(unsafe.Add(mBase, uint32(v1507)+4)) = v364 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v364+v1507))) = v364
	if base.Ui32(v364) <= base.Ui32(int32(255)) {
		goto L339
	} else {
		goto L340
	}
L339:
	;
	v1516 = v364 & int32(-8)
	v1518 = v1516 + int32(4730976)
	v1520 = *(*int32)(unsafe.Add(mBase, _consts[1503]))
	v1524 = int32(1) << (uint(int32(base.Ui32(v364)>>(uint(int32(3))%32))) % 32)
	if v1520&v1524 == int32(0) {
		goto L343
	} else {
		goto L344
	}
L340:
	;
	goto L341
L341:
	;
	if base.Ui32(v364) <= base.Ui32(int32(16777215)) {
		goto L346
	} else {
		goto L347
	}
L342:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1516)+uint32(_consts[1504]))) = v1507
	*(*int32)(unsafe.Add(mBase, uint32(v1532)+12)) = v1507
	*(*int32)(unsafe.Add(mBase, uint32(v1507)+12)) = v1518
	*(*int32)(unsafe.Add(mBase, uint32(v1507)+8)) = v1532
	goto L335
L343:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1503])) = v1524 | v1520
	v1532 = v1518
	goto L342
L344:
	;
	goto L345
L345:
	;
	v1531 = *(*int32)(unsafe.Add(mBase, uint32(v1516)+uint32(_consts[1504])))
	v1532 = v1531
	goto L342
L346:
	;
	v1543 = base.I32_clz(int32(base.Ui32(v364) >> (uint(int32(8)) % 32)))
	v1546 = int32(1)
	v1553 = int32(base.Ui32(v364)>>(uint(int32(38)-v1543)%32))&v1546 - v1543<<(uint(v1546)%32) + int32(62)
	goto L348
L347:
	;
	v1553 = int32(31)
	goto L348
L348:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1507)+28)) = v1553
	*(*int64)(unsafe.Add(mBase, uint32(v1507)+16)) = int64(0)
	v1558 = v1553 << (uint(int32(2)) % 32)
	v1562 = int32(1) << (uint(v1553) % 32)
	if v1492&v1562 == int32(0) {
		goto L351
	} else {
		goto L352
	}
L349:
	;
	v1621 = *(*int32)(unsafe.Add(mBase, uint32(v1584)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1621)+12)) = v1507
	*(*int32)(unsafe.Add(mBase, uint32(v1584)+8)) = v1507
	*(*int32)(unsafe.Add(mBase, uint32(v1507)+24)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1507)+12)) = v1584
	*(*int32)(unsafe.Add(mBase, uint32(v1507)+8)) = v1621
	goto L335
L350:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1507)+12)) = v1507
	*(*int32)(unsafe.Add(mBase, uint32(v1507)+8)) = v1507
	goto L335
L351:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1505])) = v1562 | v1492
	*(*int32)(unsafe.Add(mBase, uint32(v1558)+uint32(_consts[1507]))) = v1507
	*(*int32)(unsafe.Add(mBase, uint32(v1507)+24)) = v1558 + int32(4731240)
	goto L350
L352:
	;
	goto L353
L353:
	;
	if v1553 != int32(31) {
		goto L354
	} else {
		goto L355
	}
L354:
	;
	v1578 = int32(25) - int32(base.Ui32(v1553)>>(uint(int32(1))%32))
	goto L356
L355:
	;
	v1578 = int32(0)
	goto L356
L356:
	;
	v1580 = *(*int32)(unsafe.Add(mBase, uint32(v1558)+uint32(_consts[1507])))
	v1581 = v364 << (uint(v1578) % 32)
	v1584 = v1580
	goto L357
L357:
	;
	v1593 = *(*int32)(unsafe.Add(mBase, uint32(v1584)+4))
	if v1593&int32(-8) == v364 {
		goto L349
	} else {
		goto L359
	}
L358:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1603)+16)) = v1507
	*(*int32)(unsafe.Add(mBase, uint32(v1507)+24)) = v1584
	goto L350
L359:
	;
	v1603 = v1584 + int32(base.Ui32(v1581)>>(uint(int32(29))%32))&int32(4)
	v1604 = *(*int32)(unsafe.Add(mBase, uint32(v1603)+16))
	if v1604 != 0 {
		v1581 = v1581 << (uint(int32(1)) % 32)
		v1584 = v1604
		goto L357
	} else {
		goto L360
	}
L360:
	;
	goto L358
L361:
	;
	if base.Ui32(v153) <= base.Ui32(int32(15)) {
		goto L378
	} else {
		goto L379
	}
L362:
	;
	v1656 = *(*int32)(unsafe.Add(mBase, uint32(v154)+28))
	v1658 = v1656 << (uint(int32(2)) % 32)
	v1661 = *(*int32)(unsafe.Add(mBase, uint32(v1658)+uint32(_consts[1507])))
	if v1661 == v154 {
		goto L364
	} else {
		goto L365
	}
L363:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1642)+24)) = v178
	v1676 = *(*int32)(unsafe.Add(mBase, uint32(v154)+16))
	if v1676 != 0 {
		goto L373
	} else {
		goto L374
	}
L364:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1658)+uint32(_consts[1507]))) = v1642
	if v1642 != 0 {
		goto L363
	} else {
		goto L367
	}
L365:
	;
	goto L366
L366:
	;
	v1669 = *(*int32)(unsafe.Add(mBase, uint32(v178)+16))
	if v154 == v1669 {
		goto L369
	} else {
		goto L370
	}
L367:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1505])) = v139 & base.I32_rotl(int32(-2), v1656)
	goto L361
L368:
	;
	if v1642 == int32(0) {
		goto L361
	} else {
		goto L372
	}
L369:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v178)+16)) = v1642
	goto L368
L370:
	;
	goto L371
L371:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v178)+20)) = v1642
	goto L368
L372:
	;
	goto L363
L373:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1642)+16)) = v1676
	*(*int32)(unsafe.Add(mBase, uint32(v1676)+24)) = v1642
	goto L375
L374:
	;
	goto L375
L375:
	;
	v1679 = *(*int32)(unsafe.Add(mBase, uint32(v154)+20))
	if v1679 == int32(0) {
		goto L361
	} else {
		goto L376
	}
L376:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1642)+20)) = v1679
	*(*int32)(unsafe.Add(mBase, uint32(v1679)+24)) = v1642
	goto L361
L377:
	;
	v1741 = v154 + int32(8)
	goto L1
L378:
	;
	v1688 = v153 + v28
	*(*int32)(unsafe.Add(mBase, uint32(v154)+4)) = v1688 | int32(3)
	v1692 = v1688 + v154
	v1693 = *(*int32)(unsafe.Add(mBase, uint32(v1692)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1692)+4)) = v1693 | int32(1)
	goto L377
L379:
	;
	goto L380
L380:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v154)+4)) = v28 | int32(3)
	v1700 = v154 + v28
	*(*int32)(unsafe.Add(mBase, uint32(v1700)+4)) = v153 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v153+v1700))) = v153
	if v68 != 0 {
		goto L381
	} else {
		goto L382
	}
L381:
	;
	v1707 = v68 & int32(-8)
	v1709 = v1707 + int32(4730976)
	v1711 = *(*int32)(unsafe.Add(mBase, _consts[1506]))
	v1715 = int32(1) << (uint(int32(base.Ui32(v68)>>(uint(int32(3))%32))) % 32)
	if v1715&v20 == int32(0) {
		goto L385
	} else {
		goto L386
	}
L382:
	;
	goto L383
L383:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1506])) = v1700
	*(*int32)(unsafe.Add(mBase, _consts[1502])) = v153
	goto L377
L384:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1707)+uint32(_consts[1504]))) = v1711
	*(*int32)(unsafe.Add(mBase, uint32(v1723)+12)) = v1711
	*(*int32)(unsafe.Add(mBase, uint32(v1711)+12)) = v1709
	*(*int32)(unsafe.Add(mBase, uint32(v1711)+8)) = v1723
	goto L383
L385:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1503])) = v1715 | v20
	v1723 = v1709
	goto L384
L386:
	;
	goto L387
L387:
	;
	v1722 = *(*int32)(unsafe.Add(mBase, uint32(v1707)+uint32(_consts[1504])))
	v1723 = v1722
	goto L384
}
