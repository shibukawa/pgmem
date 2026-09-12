package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F__equalUpdateStmt(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
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
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
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
	var v33 int32
	_ = v33
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
	v3 = int32(0)
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v6 = F_equal(m, v4, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		if v6 == int32(0) {
			v40 = v3
			return v40
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
			v14 = F_equal(m, v12, v13)
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				if v14 == int32(0) {
					v40 = v3
					return v40
				} else {
					v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
					v20 = F_equal(m, v18, v19)
					mBase = m.M
					v21 = m.ExcPending
					if v21 != 0 {
						return int32(0)
					} else {
						if v20 == int32(0) {
							v40 = v3
							return v40
						} else {
							v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
							v25 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
							v26 = F_equal(m, v24, v25)
							mBase = m.M
							v27 = m.ExcPending
							if v27 != 0 {
								return int32(0)
							} else {
								if v26 == int32(0) {
									v40 = v3
									return v40
								} else {
									v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
									v31 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
									v32 = F_equal(m, v30, v31)
									mBase = m.M
									v33 = m.ExcPending
									if v33 != 0 {
										return int32(0)
									} else {
										if v32 == int32(0) {
											v40 = v3
											return v40
										} else {
											v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
											v37 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
											v38 = F_equal(m, v36, v37)
											mBase = m.M
											v39 = m.ExcPending
											if v39 != 0 {
												return int32(0)
											} else {
												v40 = v38
												return v40
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
func F_exec_stmt_block(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
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
	var v79 int32
	_ = v79
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v110 int32
	_ = v110
	var v135 int32
	_ = v135
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v296 int32
	_ = v296
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v341 int32
	_ = v341
	var v363 int32
	_ = v363
	var v382 int32
	_ = v382
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v426 int32
	_ = v426
	var v428 int32
	_ = v428
	var v432 int32
	_ = v432
	var v434 int32
	_ = v434
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v487 int32
	_ = v487
	var v491 int32
	_ = v491
	var v493 int32
	_ = v493
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
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
	var v525 int32
	_ = v525
	var v560 int32
	_ = v560
	var v563 int32
	_ = v563
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v586 int32
	_ = v586
	var v589 int32
	_ = v589
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v631 int32
	_ = v631
	var v632 int32
	_ = v632
	var v654 int32
	_ = v654
	var v656 int32
	_ = v656
	var v659 int32
	_ = v659
	var v660 int32
	_ = v660
	var v666 int32
	_ = v666
	var v667 int32
	_ = v667
	var v676 int32
	_ = v676
	var v677 int32
	_ = v677
	var v681 int32
	_ = v681
	var v686 int32
	_ = v686
	var v689 int32
	_ = v689
	var v692 int32
	_ = v692
	var v726 int32
	_ = v726
	var v727 int32
	_ = v727
	var v746 int32
	_ = v746
	var v765 int32
	_ = v765
	var v791 int32
	_ = v791
	var v795 int32
	_ = v795
	var v814 int32
	_ = v814
	var v816 int32
	_ = v816
	var v817 int32
	_ = v817
	var v820 int32
	_ = v820
	var v823 int32
	_ = v823
	var v826 int32
	_ = v826
	var v827 int32
	_ = v827
	var v859 int32
	_ = v859
	var v870 int32
	_ = v870
	var v871 int32
	_ = v871
	var v872 int32
	_ = v872
	var v882 int32
	_ = v882
	var v917 int32
	_ = v917
	var v922 int32
	_ = v922
	var v962 int32
	_ = v962
	var v1021 int32
	_ = v1021
	var v1024 int32
	_ = v1024
	var v1042 int32
	_ = v1042
	var v1043 int32
	_ = v1043
	var v1045 int32
	_ = v1045
	var v1046 int32
	_ = v1046
	var v1047 int32
	_ = v1047
	var v1050 int32
	_ = v1050
	var v1051 int32
	_ = v1051
	var v1055 int32
	_ = v1055
	var v1073 int32
	_ = v1073
	var v1074 int32
	_ = v1074
	var v1076 int32
	_ = v1076
	var v1077 int32
	_ = v1077
	var v1085 int32
	_ = v1085
	var v1093 int32
	_ = v1093
	var v1101 int32
	_ = v1101
	var v1109 int32
	_ = v1109
	var v1112 int32
	_ = v1112
	var v1132 int32
	_ = v1132
	var v1133 int32
	_ = v1133
	var v1154 int32
	_ = v1154
	var v1155 int32
	_ = v1155
	var v1173 int32
	_ = v1173
	var v1174 int32
	_ = v1174
	var v1195 int32
	_ = v1195
	var v1199 int32
	_ = v1199
	var v1217 int32
	_ = v1217
	var v1218 int32
	_ = v1218
	var v1221 int32
	_ = v1221
	var v1240 int32
	_ = v1240
	var v1260 int32
	_ = v1260
	var v1307 int32
	_ = v1307
	var v1308 int32
	_ = v1308
	var v1309 int32
	_ = v1309
	var v1310 int32
	_ = v1310
	var v1311 int32
	_ = v1311
	var v1312 int32
	_ = v1312
	var v1313 int32
	_ = v1313
	var v1314 int32
	_ = v1314
	var v1315 int32
	_ = v1315
	var v1316 int32
	_ = v1316
	var v1317 int32
	_ = v1317
	var v1318 int32
	_ = v1318
	var v1319 int32
	_ = v1319
	var v1320 int32
	_ = v1320
	var v1321 int32
	_ = v1321
	var v1322 int32
	_ = v1322
	var v1323 int32
	_ = v1323
	var v1328 int32
	_ = v1328
	var v1343 int32
	_ = v1343
	var v1350 int32
	_ = v1350
	var v1351 int32
	_ = v1351
	var v1352 int32
	_ = v1352
	var v1355 int32
	_ = v1355
	var v1377 int32
	_ = v1377
	var v1378 int32
	_ = v1378
	var v1382 int32
	_ = v1382
	var v1383 int32
	_ = v1383
	var v1386 int32
	_ = v1386
	var v1387 int32
	_ = v1387
	var v1390 int32
	_ = v1390
	var v1397 int32
	_ = v1397
	var v1398 int32
	_ = v1398
	var v1400 int32
	_ = v1400
	var v1403 int32
	_ = v1403
	var v1430 int32
	_ = v1430
	var v1431 int32
	_ = v1431
	var v1452 int32
	_ = v1452
	var v1474 int32
	_ = v1474
	var v1500 int32
	_ = v1500
	var v1513 int32
	_ = v1513
	var v1514 int64
	_ = v1514
	var v1518 int32
	_ = v1518
	var v1520 int32
	_ = v1520
	var v1521 int32
	_ = v1521
	var v1525 int32
	_ = v1525
	var v1527 int32
	_ = v1527
	var v1528 int32
	_ = v1528
	var v1529 int32
	_ = v1529
	var v1530 int32
	_ = v1530
	var v1531 int32
	_ = v1531
	var v1532 int32
	_ = v1532
	var v1533 int32
	_ = v1533
	var v1534 int32
	_ = v1534
	var v1535 int32
	_ = v1535
	var v1536 int32
	_ = v1536
	var v1537 int32
	_ = v1537
	var v1538 int32
	_ = v1538
	var v1539 int32
	_ = v1539
	var v1540 int32
	_ = v1540
	var v1541 int32
	_ = v1541
	var v1542 int32
	_ = v1542
	var v1543 int32
	_ = v1543
	var v1544 int32
	_ = v1544
	var v1545 int32
	_ = v1545
	var v1549 int32
	_ = v1549
	v3 = int32(0)
	v39 = m.G0
	v40 = int32(96)
	v41 = v39 - v40
	m.G0 = v41
	v44 = l0 + int32(132)
	v46 = l1 + int32(28)
	v48 = l0 + v40
	v50 = l0 + int32(36)
	v52 = l0 + int32(120)
	v57 = v3
	v58 = v3
	v59 = v3
	v60 = v3
	v61 = v3
	v62 = v3
	v63 = v3
	v64 = v3
	v65 = v3
	v66 = v3
	v67 = v3
	v68 = v3
	v69 = v3
	v70 = v3
	v71 = v3
	v72 = v3
	v73 = v3
	v74 = v3
	v75 = int32(-1)
	v79 = v41
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
	if v75 != int32(1) {
		goto L11
	} else {
		goto L12
	}
L5:
	;
	goto L3
L6:
	;
	v1513 = int32(m.ExcTag)
	v1514 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v1513 == int32(0) {
		goto L140
	} else {
		goto L141
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1323))) = int32(0)
	v1343 = *(*int32)(unsafe.Add(mBase, uint32(v1307)))
	if v1343 != int32(1) {
		goto L121
	} else {
		goto L122
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, _consts[74])) = v512
	*(*int32)(unsafe.Add(mBase, _consts[292])) = v511
	v1307 = v504
	v1308 = v505
	v1309 = v506
	v1310 = v507
	v1311 = v508
	v1312 = v509
	v1313 = v510
	v1314 = v511
	v1315 = v512
	v1316 = v513
	v1317 = v514
	v1318 = v515
	v1319 = v516
	v1320 = v517
	v1321 = v518
	v1322 = v519
	v1323 = v520
	v1328 = v525
	goto L7
L9:
	;
	v1045 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v1046 = *(*int32)(unsafe.Add(mBase, uint32(v816)+4))
	v1047 = int32(2)
	v1050 = *(*int32)(unsafe.Add(mBase, uint32(v1045+v1046<<(uint(v1047)%32))))
	v1051 = *(*int32)(unsafe.Add(mBase, uint32(v816)))
	v1055 = *(*int32)(unsafe.Add(mBase, uint32(v1045+v1051<<(uint(v1047)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v41)+32)) = v511
	*(*int32)(unsafe.Add(mBase, uint32(v41)+28)) = v512
	*(*int32)(unsafe.Add(mBase, uint32(v41)+36)) = v505
	*(*int32)(unsafe.Add(mBase, uint32(v41)+40)) = v515
	*(*int32)(unsafe.Add(mBase, uint32(v41)+44)) = v517
	*(*int32)(unsafe.Add(mBase, uint32(v41)+48)) = v513
	*(*int32)(unsafe.Add(mBase, uint32(v41)+52)) = v516
	*(*int32)(unsafe.Add(mBase, uint32(v41)+56)) = v506
	*(*int32)(unsafe.Add(mBase, uint32(v41)+60)) = v518
	*(*int32)(unsafe.Add(mBase, uint32(v41)+64)) = v514
	*(*int32)(unsafe.Add(mBase, uint32(v41)+68)) = v507
	*(*int32)(unsafe.Add(mBase, uint32(v41)+72)) = v519
	*(*int32)(unsafe.Add(mBase, uint32(v41)+76)) = v520
	*(*int32)(unsafe.Add(mBase, uint32(v41)+80)) = v510
	*(*int32)(unsafe.Add(mBase, uint32(v41)+84)) = v509
	*(*int32)(unsafe.Add(mBase, uint32(v41)+88)) = v508
	*(*int32)(unsafe.Add(mBase, uint32(v41)+92)) = v504
	v1073 = int32(4542760)
	v1074 = int32(63)
	v1076 = int32(48)
	v1077 = v872&v1074 + v1076
	*(*uint8)(unsafe.Add(mBase, _consts[1048])) = uint8(v1077)
	v1085 = int32(base.Ui32(v872)>>(uint(int32(24))%32))&v1074 + v1076
	*(*uint8)(unsafe.Add(mBase, _consts[1049])) = uint8(v1085)
	v1093 = int32(base.Ui32(v872)>>(uint(int32(18))%32))&v1074 + v1076
	*(*uint8)(unsafe.Add(mBase, _consts[1050])) = uint8(v1093)
	v1101 = int32(base.Ui32(v872)>>(uint(int32(12))%32))&v1074 + v1076
	*(*uint8)(unsafe.Add(mBase, _consts[1051])) = uint8(v1101)
	v1109 = int32(base.Ui32(v872)>>(uint(int32(6))%32))&v1074 + v1076
	*(*uint8)(unsafe.Add(mBase, _consts[1052])) = uint8(v1109)
	v1112 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _consts[1053])) = uint8(v1112)
	goto L112
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+132)) = int32(0)
	v1024 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v41)+32)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v41)+28)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v41)+36)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v41)+40)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v41)+44)) = v71
	*(*int32)(unsafe.Add(mBase, uint32(v41)+48)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v41)+52)) = v70
	*(*int32)(unsafe.Add(mBase, uint32(v41)+56)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v41)+60)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v41)+64)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v41)+68)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v41)+72)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v41)+76)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v41)+80)) = v104
	*(*int32)(unsafe.Add(mBase, uint32(v41)+84)) = v101
	*(*int32)(unsafe.Add(mBase, uint32(v41)+88)) = v98
	*(*int32)(unsafe.Add(mBase, uint32(v41)+92)) = v95
	v1042 = F_exec_stmts(m, l0, v1024)
	mBase = m.M
	v1043 = m.ExcPending
	if v1043 != 0 {
		v1500 = v104
		goto L6
	} else {
		goto L111
	}
L11:
	;
	v94 = int32(16)
	v95 = v79 - v94
	m.G0 = v95
	v98 = v95 - int32(160)
	m.G0 = v98
	v101 = v98 - v94
	m.G0 = v101
	v104 = v101 - v94
	m.G0 = v104
	*(*int32)(unsafe.Add(mBase, uint32(v95))) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+132)) = int32(269401)
	v110 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if int32(0) < v110 {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	v503 = v57
	v504 = v58
	v505 = v59
	v506 = v60
	v507 = v61
	v508 = v62
	v509 = v63
	v510 = v64
	v511 = v65
	v512 = v66
	v513 = v67
	v514 = v68
	v515 = v69
	v516 = v70
	v517 = v71
	v518 = v72
	v519 = v73
	v520 = v74
	v525 = v79
	goto L13
L13:
	;
	if v503 == int32(0) {
		goto L57
	} else {
		goto L58
	}
L14:
	;
	v135 = int32(0)
	goto L17
L15:
	;
	goto L16
L16:
	;
	v426 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+128)) = v426
	v428 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	if v428 == v426 {
		goto L10
	} else {
		goto L46
	}
L17:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v153 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v154 = int32(2)
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v153+v135<<(uint(v154)%32))))
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v152+v157<<(uint(v154)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+128)) = v161
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v161)))
	switch v163 {
	case 0:
		goto L23
	default:
		goto L21
	case 2:
		goto L22
	}
L18:
	;
	goto L16
L19:
	;
	v385 = v135 + int32(1)
	v386 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v385 < v386 {
		v135 = v385
		goto L17
	} else {
		goto L45
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+32)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v41)+28)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v41)+36)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v41)+40)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v41)+44)) = v71
	*(*int32)(unsafe.Add(mBase, uint32(v41)+48)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v41)+52)) = v70
	*(*int32)(unsafe.Add(mBase, uint32(v41)+56)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v41)+60)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v41)+64)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v41)+68)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v41)+72)) = v73
	*(*int32)(unsafe.Add(mBase, uint32(v41)+76)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v41)+80)) = v104
	*(*int32)(unsafe.Add(mBase, uint32(v41)+84)) = v101
	*(*int32)(unsafe.Add(mBase, uint32(v41)+88)) = v98
	*(*int32)(unsafe.Add(mBase, uint32(v41)+92)) = v95
	F_exec_assign_expr(m, l0, v161, v226)
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		v1500 = v104
		goto L6
	} else {
		goto L44
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+32)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v41)+28)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v41)+36)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v41)+40)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v41)+44)) = v71
	*(*int32)(unsafe.Add(mBase, uint32(v41)+48)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v41)+52)) = v70
	*(*int32)(unsafe.Add(mBase, uint32(v41)+56)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v41)+60)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v41)+64)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v41)+68)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v41)+72)) = v73
	*(*int32)(unsafe.Add(mBase, uint32(v41)+76)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v41)+80)) = v104
	*(*int32)(unsafe.Add(mBase, uint32(v41)+84)) = v101
	*(*int32)(unsafe.Add(mBase, uint32(v41)+88)) = v98
	*(*int32)(unsafe.Add(mBase, uint32(v41)+92)) = v95
	F_errstart_cold(m, int32(21), int32(578012))
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		v1500 = v104
		goto L6
	} else {
		goto L41
	}
L22:
	;
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v161)+20))
	if v254 == int32(0) {
		goto L36
	} else {
		goto L37
	}
L23:
	;
	v164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161)+45)))
	if v164 != int32(1) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v220 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v161)+48)) = v220
	v222 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v161)+44)) = uint16(v222)
	*(*int32)(unsafe.Add(mBase, uint32(v161)+40)) = v220
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v161)+20))
	if v226 != 0 {
		goto L20
	} else {
		goto L33
	}
L25:
	;
	v167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161)+44)))
	if v167 != 0 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v161)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v41)+32)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v41)+28)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v41)+36)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v41)+40)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v41)+44)) = v71
	*(*int32)(unsafe.Add(mBase, uint32(v41)+48)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v41)+52)) = v70
	*(*int32)(unsafe.Add(mBase, uint32(v41)+56)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v41)+60)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v41)+64)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v41)+68)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v41)+72)) = v73
	*(*int32)(unsafe.Add(mBase, uint32(v41)+76)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v41)+80)) = v104
	*(*int32)(unsafe.Add(mBase, uint32(v41)+84)) = v101
	*(*int32)(unsafe.Add(mBase, uint32(v41)+88)) = v98
	*(*int32)(unsafe.Add(mBase, uint32(v41)+92)) = v95
	F_pfree(m, v199)
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		v1500 = v104
		goto L6
	} else {
		goto L32
	}
L27:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v161)+24))
	v169 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v168)+12)))
	if v169 != int32(65535) {
		goto L26
	} else {
		goto L28
	}
L28:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v161)+40))
	v173 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v172))))
	if v173 != int32(1) {
		goto L26
	} else {
		goto L29
	}
L29:
	;
	v176 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v172)+1)))
	if v176 != int32(3) {
		goto L26
	} else {
		goto L30
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+32)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v41)+28)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v41)+36)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v41)+40)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v41)+44)) = v71
	*(*int32)(unsafe.Add(mBase, uint32(v41)+48)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v41)+52)) = v70
	*(*int32)(unsafe.Add(mBase, uint32(v41)+56)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v41)+60)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v41)+64)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v41)+68)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v41)+72)) = v73
	*(*int32)(unsafe.Add(mBase, uint32(v41)+76)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v41)+80)) = v104
	*(*int32)(unsafe.Add(mBase, uint32(v41)+84)) = v101
	*(*int32)(unsafe.Add(mBase, uint32(v41)+88)) = v98
	*(*int32)(unsafe.Add(mBase, uint32(v41)+92)) = v95
	F_DeleteExpandedObject(m, v172)
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		v1500 = v104
		goto L6
	} else {
		goto L31
	}
L31:
	;
	goto L24
L32:
	;
	goto L24
L33:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v161)+24))
	v228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v227)+15)))
	if v228 != int32(100) {
		goto L19
	} else {
		goto L34
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+32)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v41)+28)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v41)+36)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v41)+40)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v41)+44)) = v71
	*(*int32)(unsafe.Add(mBase, uint32(v41)+48)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v41)+52)) = v70
	*(*int32)(unsafe.Add(mBase, uint32(v41)+56)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v41)+60)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v41)+64)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v41)+68)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v41)+72)) = v73
	*(*int32)(unsafe.Add(mBase, uint32(v41)+76)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v41)+80)) = v104
	*(*int32)(unsafe.Add(mBase, uint32(v41)+84)) = v101
	*(*int32)(unsafe.Add(mBase, uint32(v41)+88)) = v98
	*(*int32)(unsafe.Add(mBase, uint32(v41)+92)) = v95
	F_exec_assign_value(m, l0, v161, int32(0), int32(1), int32(705), int32(-1))
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		v1500 = v104
		goto L6
	} else {
		goto L35
	}
L35:
	;
	goto L19
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+32)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v41)+28)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v41)+36)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v41)+40)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v41)+44)) = v71
	*(*int32)(unsafe.Add(mBase, uint32(v41)+48)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v41)+52)) = v70
	*(*int32)(unsafe.Add(mBase, uint32(v41)+56)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v41)+60)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v41)+64)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v41)+68)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v41)+72)) = v73
	*(*int32)(unsafe.Add(mBase, uint32(v41)+76)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v41)+80)) = v104
	*(*int32)(unsafe.Add(mBase, uint32(v41)+84)) = v101
	*(*int32)(unsafe.Add(mBase, uint32(v41)+88)) = v98
	*(*int32)(unsafe.Add(mBase, uint32(v41)+92)) = v95
	v274 = int32(0)
	F_exec_move_row(m, l0, v161, v274, v274)
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		v1500 = v104
		goto L6
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+32)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v41)+28)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v41)+36)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v41)+40)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v41)+44)) = v71
	*(*int32)(unsafe.Add(mBase, uint32(v41)+48)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v41)+52)) = v70
	*(*int32)(unsafe.Add(mBase, uint32(v41)+56)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v41)+60)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v41)+64)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v41)+68)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v41)+72)) = v73
	*(*int32)(unsafe.Add(mBase, uint32(v41)+76)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v41)+80)) = v104
	*(*int32)(unsafe.Add(mBase, uint32(v41)+84)) = v101
	*(*int32)(unsafe.Add(mBase, uint32(v41)+88)) = v98
	*(*int32)(unsafe.Add(mBase, uint32(v41)+92)) = v95
	F_exec_assign_expr(m, l0, v161, v254)
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		v1500 = v104
		goto L6
	} else {
		goto L40
	}
L39:
	;
	goto L19
L40:
	;
	goto L19
L41:
	;
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v161)))
	*(*int32)(unsafe.Add(mBase, uint32(v41)+32)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v41)+28)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v41)+36)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v41)+40)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v41)+44)) = v71
	*(*int32)(unsafe.Add(mBase, uint32(v41)+48)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v41)+52)) = v70
	*(*int32)(unsafe.Add(mBase, uint32(v41)+56)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v41)+60)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v41)+64)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v41)+68)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v41)+72)) = v73
	*(*int32)(unsafe.Add(mBase, uint32(v41)+76)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v41)+80)) = v104
	*(*int32)(unsafe.Add(mBase, uint32(v41)+84)) = v101
	*(*int32)(unsafe.Add(mBase, uint32(v41)+88)) = v98
	*(*int32)(unsafe.Add(mBase, uint32(v41)+92)) = v95
	*(*int32)(unsafe.Add(mBase, uint32(v41)+16)) = v318
	F_errmsg_internal(m, int32(502921), v41+int32(16))
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		v1500 = v104
		goto L6
	} else {
		goto L42
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+32)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v41)+28)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v41)+36)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v41)+40)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v41)+44)) = v71
	*(*int32)(unsafe.Add(mBase, uint32(v41)+48)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v41)+52)) = v70
	*(*int32)(unsafe.Add(mBase, uint32(v41)+56)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v41)+60)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v41)+64)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v41)+68)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v41)+72)) = v73
	*(*int32)(unsafe.Add(mBase, uint32(v41)+76)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v41)+80)) = v104
	*(*int32)(unsafe.Add(mBase, uint32(v41)+84)) = v101
	*(*int32)(unsafe.Add(mBase, uint32(v41)+88)) = v98
	*(*int32)(unsafe.Add(mBase, uint32(v41)+92)) = v95
	F_errfinish(m, int32(520770), int32(1759), int32(329261))
	mBase = m.M
	v363 = m.ExcPending
	if v363 != 0 {
		v1500 = v104
		goto L6
	} else {
		goto L43
	}
L43:
	;
	goto L3
L44:
	;
	goto L19
L45:
	;
	goto L18
L46:
	;
	v432 = *(*int32)(unsafe.Add(mBase, _consts[579]))
	v434 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+132)) = int32(12564)
	v437 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v438 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v439 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	if v439 != 0 {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+32)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v41)+28)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v41)+36)) = v467
	*(*int32)(unsafe.Add(mBase, uint32(v41)+40)) = v466
	*(*int32)(unsafe.Add(mBase, uint32(v41)+44)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v41)+48)) = v437
	*(*int32)(unsafe.Add(mBase, uint32(v41)+52)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v41)+56)) = v438
	*(*int32)(unsafe.Add(mBase, uint32(v41)+60)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v41)+64)) = v432
	*(*int32)(unsafe.Add(mBase, uint32(v41)+68)) = v434
	*(*int32)(unsafe.Add(mBase, uint32(v41)+72)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v41)+76)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v41)+80)) = v104
	*(*int32)(unsafe.Add(mBase, uint32(v41)+84)) = v101
	*(*int32)(unsafe.Add(mBase, uint32(v41)+88)) = v98
	*(*int32)(unsafe.Add(mBase, uint32(v41)+92)) = v95
	F_BeginInternalSubTransaction(m, int32(0))
	mBase = m.M
	v487 = m.ExcPending
	if v487 != 0 {
		v1500 = v104
		goto L6
	} else {
		goto L52
	}
L48:
	;
	v466 = v69
	v467 = v439
	goto L47
L49:
	;
	goto L50
L50:
	;
	v440 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	*(*int32)(unsafe.Add(mBase, uint32(v41)+32)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v41)+28)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v41)+36)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v41)+40)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v41)+44)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v41)+48)) = v437
	*(*int32)(unsafe.Add(mBase, uint32(v41)+52)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v41)+56)) = v438
	*(*int32)(unsafe.Add(mBase, uint32(v41)+60)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v41)+64)) = v432
	*(*int32)(unsafe.Add(mBase, uint32(v41)+68)) = v434
	*(*int32)(unsafe.Add(mBase, uint32(v41)+72)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v41)+76)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v41)+80)) = v104
	*(*int32)(unsafe.Add(mBase, uint32(v41)+84)) = v101
	*(*int32)(unsafe.Add(mBase, uint32(v41)+88)) = v98
	*(*int32)(unsafe.Add(mBase, uint32(v41)+92)) = v95
	v462 = F_AllocSetContextCreateInternal(m, v440, int32(525561), int32(0), int32(8192), int32(8388608))
	mBase = m.M
	v463 = m.ExcPending
	if v463 != 0 {
		v1500 = v104
		goto L6
	} else {
		goto L51
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48))) = v462
	v466 = v462
	v467 = v462
	goto L47
L52:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v434
	v491 = *(*int32)(unsafe.Add(mBase, _consts[292]))
	v493 = *(*int32)(unsafe.Add(mBase, _consts[74]))
	goto L53
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v98)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v98))) = v41 + int32(24)
	goto L56
L54:
	;
	v503 = int32(0)
	v504 = v95
	v505 = v467
	v506 = v438
	v507 = v434
	v508 = v98
	v509 = v101
	v510 = v104
	v511 = v491
	v512 = v493
	v513 = v437
	v514 = v432
	v515 = v466
	v516 = v50
	v517 = v48
	v518 = v52
	v519 = v46
	v520 = v44
	v525 = v104
	goto L13
L56:
	;
	goto L54
L57:
	;
	*(*int32)(unsafe.Add(mBase, _consts[292])) = v508
	*(*int32)(unsafe.Add(mBase, uint32(v41)+28)) = v512
	*(*int32)(unsafe.Add(mBase, uint32(v41)+32)) = v511
	*(*int32)(unsafe.Add(mBase, uint32(v41)+36)) = v505
	*(*int32)(unsafe.Add(mBase, uint32(v41)+40)) = v515
	*(*int32)(unsafe.Add(mBase, uint32(v41)+44)) = v517
	*(*int32)(unsafe.Add(mBase, uint32(v41)+48)) = v513
	*(*int32)(unsafe.Add(mBase, uint32(v41)+52)) = v516
	*(*int32)(unsafe.Add(mBase, uint32(v41)+56)) = v506
	*(*int32)(unsafe.Add(mBase, uint32(v41)+60)) = v518
	*(*int32)(unsafe.Add(mBase, uint32(v41)+64)) = v514
	*(*int32)(unsafe.Add(mBase, uint32(v41)+68)) = v507
	*(*int32)(unsafe.Add(mBase, uint32(v41)+72)) = v519
	*(*int32)(unsafe.Add(mBase, uint32(v41)+76)) = v520
	*(*int32)(unsafe.Add(mBase, uint32(v41)+80)) = v510
	*(*int32)(unsafe.Add(mBase, uint32(v41)+84)) = v509
	*(*int32)(unsafe.Add(mBase, uint32(v41)+88)) = v508
	*(*int32)(unsafe.Add(mBase, uint32(v41)+92)) = v504
	F_plpgsql_create_econtext(m, l0)
	mBase = m.M
	v560 = m.ExcPending
	if v560 != 0 {
		v1500 = v525
		goto L6
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	*(*int32)(unsafe.Add(mBase, _consts[74])) = v512
	*(*int32)(unsafe.Add(mBase, _consts[292])) = v511
	*(*int32)(unsafe.Add(mBase, uint32(v520))) = int32(242745)
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v505
	*(*int32)(unsafe.Add(mBase, uint32(v41)+28)) = v512
	*(*int32)(unsafe.Add(mBase, uint32(v41)+32)) = v511
	*(*int32)(unsafe.Add(mBase, uint32(v41)+36)) = v505
	*(*int32)(unsafe.Add(mBase, uint32(v41)+40)) = v515
	*(*int32)(unsafe.Add(mBase, uint32(v41)+44)) = v517
	*(*int32)(unsafe.Add(mBase, uint32(v41)+48)) = v513
	*(*int32)(unsafe.Add(mBase, uint32(v41)+52)) = v516
	*(*int32)(unsafe.Add(mBase, uint32(v41)+56)) = v506
	*(*int32)(unsafe.Add(mBase, uint32(v41)+60)) = v518
	*(*int32)(unsafe.Add(mBase, uint32(v41)+64)) = v514
	*(*int32)(unsafe.Add(mBase, uint32(v41)+68)) = v507
	*(*int32)(unsafe.Add(mBase, uint32(v41)+72)) = v519
	*(*int32)(unsafe.Add(mBase, uint32(v41)+76)) = v520
	*(*int32)(unsafe.Add(mBase, uint32(v41)+80)) = v510
	*(*int32)(unsafe.Add(mBase, uint32(v41)+84)) = v509
	*(*int32)(unsafe.Add(mBase, uint32(v41)+88)) = v508
	*(*int32)(unsafe.Add(mBase, uint32(v41)+92)) = v504
	v726 = F_CopyErrorData(m)
	mBase = m.M
	v727 = m.ExcPending
	if v727 != 0 {
		v1500 = v525
		goto L6
	} else {
		goto L79
	}
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v520))) = int32(0)
	v563 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v41)+32)) = v511
	*(*int32)(unsafe.Add(mBase, uint32(v41)+28)) = v512
	*(*int32)(unsafe.Add(mBase, uint32(v41)+36)) = v505
	*(*int32)(unsafe.Add(mBase, uint32(v41)+40)) = v515
	*(*int32)(unsafe.Add(mBase, uint32(v41)+44)) = v517
	*(*int32)(unsafe.Add(mBase, uint32(v41)+48)) = v513
	*(*int32)(unsafe.Add(mBase, uint32(v41)+52)) = v516
	*(*int32)(unsafe.Add(mBase, uint32(v41)+56)) = v506
	*(*int32)(unsafe.Add(mBase, uint32(v41)+60)) = v518
	*(*int32)(unsafe.Add(mBase, uint32(v41)+64)) = v514
	*(*int32)(unsafe.Add(mBase, uint32(v41)+68)) = v507
	*(*int32)(unsafe.Add(mBase, uint32(v41)+72)) = v519
	*(*int32)(unsafe.Add(mBase, uint32(v41)+76)) = v520
	*(*int32)(unsafe.Add(mBase, uint32(v41)+80)) = v510
	*(*int32)(unsafe.Add(mBase, uint32(v41)+84)) = v509
	*(*int32)(unsafe.Add(mBase, uint32(v41)+88)) = v508
	*(*int32)(unsafe.Add(mBase, uint32(v41)+92)) = v504
	v581 = F_exec_stmts(m, l0, v563)
	mBase = m.M
	v582 = m.ExcPending
	if v582 != 0 {
		v1500 = v525
		goto L6
	} else {
		goto L61
	}
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v504))) = v581
	*(*int32)(unsafe.Add(mBase, uint32(v520))) = int32(105639)
	v586 = *(*int32)(unsafe.Add(mBase, uint32(v504)))
	if v586 != int32(2) {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+32)) = v511
	*(*int32)(unsafe.Add(mBase, uint32(v41)+28)) = v512
	*(*int32)(unsafe.Add(mBase, uint32(v41)+36)) = v505
	*(*int32)(unsafe.Add(mBase, uint32(v41)+40)) = v515
	*(*int32)(unsafe.Add(mBase, uint32(v41)+44)) = v517
	*(*int32)(unsafe.Add(mBase, uint32(v41)+48)) = v513
	*(*int32)(unsafe.Add(mBase, uint32(v41)+52)) = v516
	*(*int32)(unsafe.Add(mBase, uint32(v41)+56)) = v506
	*(*int32)(unsafe.Add(mBase, uint32(v41)+60)) = v518
	*(*int32)(unsafe.Add(mBase, uint32(v41)+64)) = v514
	*(*int32)(unsafe.Add(mBase, uint32(v41)+68)) = v507
	*(*int32)(unsafe.Add(mBase, uint32(v41)+72)) = v519
	*(*int32)(unsafe.Add(mBase, uint32(v41)+76)) = v520
	*(*int32)(unsafe.Add(mBase, uint32(v41)+80)) = v510
	*(*int32)(unsafe.Add(mBase, uint32(v41)+84)) = v509
	*(*int32)(unsafe.Add(mBase, uint32(v41)+88)) = v508
	*(*int32)(unsafe.Add(mBase, uint32(v41)+92)) = v504
	v654 = m.G0
	v656 = v654 - int32(16)
	m.G0 = v656
	v659 = *(*int32)(unsafe.Add(mBase, _consts[37]))
	v660 = *(*int32)(unsafe.Add(mBase, uint32(v659)+24))
	if v660 != int32(12) {
		goto L68
	} else {
		goto L69
	}
L63:
	;
	v589 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+29)))
	if v589 != 0 {
		goto L62
	} else {
		goto L64
	}
L64:
	;
	v590 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)))
	if v590 != 0 {
		goto L62
	} else {
		goto L65
	}
L65:
	;
	v591 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v41)+32)) = v511
	*(*int32)(unsafe.Add(mBase, uint32(v41)+28)) = v512
	*(*int32)(unsafe.Add(mBase, uint32(v41)+36)) = v505
	*(*int32)(unsafe.Add(mBase, uint32(v41)+40)) = v515
	*(*int32)(unsafe.Add(mBase, uint32(v41)+44)) = v517
	*(*int32)(unsafe.Add(mBase, uint32(v41)+48)) = v513
	*(*int32)(unsafe.Add(mBase, uint32(v41)+52)) = v516
	*(*int32)(unsafe.Add(mBase, uint32(v41)+56)) = v506
	*(*int32)(unsafe.Add(mBase, uint32(v41)+60)) = v518
	*(*int32)(unsafe.Add(mBase, uint32(v41)+64)) = v514
	*(*int32)(unsafe.Add(mBase, uint32(v41)+68)) = v507
	*(*int32)(unsafe.Add(mBase, uint32(v41)+72)) = v519
	*(*int32)(unsafe.Add(mBase, uint32(v41)+76)) = v520
	*(*int32)(unsafe.Add(mBase, uint32(v41)+80)) = v510
	*(*int32)(unsafe.Add(mBase, uint32(v41)+84)) = v509
	*(*int32)(unsafe.Add(mBase, uint32(v41)+88)) = v508
	*(*int32)(unsafe.Add(mBase, uint32(v41)+92)) = v504
	F_get_typlenbyval(m, v591, v509, v510)
	mBase = m.M
	v610 = m.ExcPending
	if v610 != 0 {
		v1500 = v525
		goto L6
	} else {
		goto L66
	}
L66:
	;
	v611 = int32(*(*int16)(unsafe.Add(mBase, uint32(v509))))
	v612 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v510))))
	v613 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v41)+32)) = v511
	*(*int32)(unsafe.Add(mBase, uint32(v41)+28)) = v512
	*(*int32)(unsafe.Add(mBase, uint32(v41)+36)) = v505
	*(*int32)(unsafe.Add(mBase, uint32(v41)+40)) = v515
	*(*int32)(unsafe.Add(mBase, uint32(v41)+44)) = v517
	*(*int32)(unsafe.Add(mBase, uint32(v41)+48)) = v513
	*(*int32)(unsafe.Add(mBase, uint32(v41)+52)) = v516
	*(*int32)(unsafe.Add(mBase, uint32(v41)+56)) = v506
	*(*int32)(unsafe.Add(mBase, uint32(v41)+60)) = v518
	*(*int32)(unsafe.Add(mBase, uint32(v41)+64)) = v514
	*(*int32)(unsafe.Add(mBase, uint32(v41)+68)) = v507
	*(*int32)(unsafe.Add(mBase, uint32(v41)+72)) = v519
	*(*int32)(unsafe.Add(mBase, uint32(v41)+76)) = v520
	*(*int32)(unsafe.Add(mBase, uint32(v41)+80)) = v510
	*(*int32)(unsafe.Add(mBase, uint32(v41)+84)) = v509
	*(*int32)(unsafe.Add(mBase, uint32(v41)+88)) = v508
	*(*int32)(unsafe.Add(mBase, uint32(v41)+92)) = v504
	v631 = F_datumTransfer(m, v613, v612, v611)
	mBase = m.M
	v632 = m.ExcPending
	if v632 != 0 {
		v1500 = v525
		goto L6
	} else {
		goto L67
	}
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v631
	goto L62
L68:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v666 = m.ExcPending
	if v666 != 0 {
		v1500 = v525
		goto L6
	} else {
		goto L71
	}
L69:
	;
	goto L70
L70:
	;
	v689 = *(*int32)(unsafe.Add(mBase, _consts[298]))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v689
	F_CommitSubTransaction(m)
	mBase = m.M
	v692 = m.ExcPending
	if v692 != 0 {
		v1500 = v525
		goto L6
	} else {
		goto L78
	}
L71:
	;
	v667 = *(*int32)(unsafe.Add(mBase, uint32(v659)+24))
	if base.Ui32(v667) <= base.Ui32(int32(19)) {
		goto L73
	} else {
		goto L74
	}
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v656))) = v677
	F_errmsg_internal(m, int32(196133), v656)
	mBase = m.M
	v681 = m.ExcPending
	if v681 != 0 {
		v1500 = v525
		goto L6
	} else {
		goto L76
	}
L73:
	;
	v676 = *(*int32)(unsafe.Add(mBase, uint32(v667<<(uint(int32(2))%32))+uint32(_consts[154])))
	v677 = v676
	goto L75
L74:
	;
	v677 = int32(565099)
	goto L75
L75:
	;
	goto L72
L76:
	;
	F_errfinish(m, int32(512898), int32(4780), int32(268816))
	mBase = m.M
	v686 = m.ExcPending
	if v686 != 0 {
		v1500 = v525
		goto L6
	} else {
		goto L77
	}
L77:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L78:
	;
	m.G0 = v656 + int32(16)
	*(*int32)(unsafe.Add(mBase, _consts[579])) = v514
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v507
	*(*int32)(unsafe.Add(mBase, uint32(v518))) = v506
	goto L8
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+32)) = v511
	*(*int32)(unsafe.Add(mBase, uint32(v41)+28)) = v512
	*(*int32)(unsafe.Add(mBase, uint32(v41)+36)) = v505
	*(*int32)(unsafe.Add(mBase, uint32(v41)+40)) = v515
	*(*int32)(unsafe.Add(mBase, uint32(v41)+44)) = v517
	*(*int32)(unsafe.Add(mBase, uint32(v41)+48)) = v513
	*(*int32)(unsafe.Add(mBase, uint32(v41)+52)) = v516
	*(*int32)(unsafe.Add(mBase, uint32(v41)+56)) = v506
	*(*int32)(unsafe.Add(mBase, uint32(v41)+60)) = v518
	*(*int32)(unsafe.Add(mBase, uint32(v41)+64)) = v514
	*(*int32)(unsafe.Add(mBase, uint32(v41)+68)) = v507
	*(*int32)(unsafe.Add(mBase, uint32(v41)+72)) = v519
	*(*int32)(unsafe.Add(mBase, uint32(v41)+76)) = v520
	*(*int32)(unsafe.Add(mBase, uint32(v41)+80)) = v510
	*(*int32)(unsafe.Add(mBase, uint32(v41)+84)) = v509
	*(*int32)(unsafe.Add(mBase, uint32(v41)+88)) = v508
	*(*int32)(unsafe.Add(mBase, uint32(v41)+92)) = v504
	F_FlushErrorState(m)
	mBase = m.M
	v746 = m.ExcPending
	if v746 != 0 {
		v1500 = v525
		goto L6
	} else {
		goto L80
	}
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+32)) = v511
	*(*int32)(unsafe.Add(mBase, uint32(v41)+28)) = v512
	*(*int32)(unsafe.Add(mBase, uint32(v41)+36)) = v505
	*(*int32)(unsafe.Add(mBase, uint32(v41)+40)) = v515
	*(*int32)(unsafe.Add(mBase, uint32(v41)+44)) = v517
	*(*int32)(unsafe.Add(mBase, uint32(v41)+48)) = v513
	*(*int32)(unsafe.Add(mBase, uint32(v41)+52)) = v516
	*(*int32)(unsafe.Add(mBase, uint32(v41)+56)) = v506
	*(*int32)(unsafe.Add(mBase, uint32(v41)+60)) = v518
	*(*int32)(unsafe.Add(mBase, uint32(v41)+64)) = v514
	*(*int32)(unsafe.Add(mBase, uint32(v41)+68)) = v507
	*(*int32)(unsafe.Add(mBase, uint32(v41)+72)) = v519
	*(*int32)(unsafe.Add(mBase, uint32(v41)+76)) = v520
	*(*int32)(unsafe.Add(mBase, uint32(v41)+80)) = v510
	*(*int32)(unsafe.Add(mBase, uint32(v41)+84)) = v509
	*(*int32)(unsafe.Add(mBase, uint32(v41)+88)) = v508
	*(*int32)(unsafe.Add(mBase, uint32(v41)+92)) = v504
	F_RollbackAndReleaseCurrentSubTransaction(m)
	mBase = m.M
	v765 = m.ExcPending
	if v765 != 0 {
		v1500 = v525
		goto L6
	} else {
		goto L81
	}
L81:
	;
	*(*int32)(unsafe.Add(mBase, _consts[579])) = v514
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v507
	*(*int32)(unsafe.Add(mBase, uint32(l0)+100)) = v505
	*(*int32)(unsafe.Add(mBase, uint32(v517))) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v41)+32)) = v511
	*(*int32)(unsafe.Add(mBase, uint32(v41)+28)) = v512
	*(*int32)(unsafe.Add(mBase, uint32(v41)+36)) = v505
	*(*int32)(unsafe.Add(mBase, uint32(v41)+40)) = v515
	*(*int32)(unsafe.Add(mBase, uint32(v41)+44)) = v517
	*(*int32)(unsafe.Add(mBase, uint32(v41)+48)) = v513
	*(*int32)(unsafe.Add(mBase, uint32(v41)+52)) = v516
	*(*int32)(unsafe.Add(mBase, uint32(v41)+56)) = v506
	*(*int32)(unsafe.Add(mBase, uint32(v41)+60)) = v518
	*(*int32)(unsafe.Add(mBase, uint32(v41)+64)) = v514
	*(*int32)(unsafe.Add(mBase, uint32(v41)+68)) = v507
	*(*int32)(unsafe.Add(mBase, uint32(v41)+72)) = v519
	*(*int32)(unsafe.Add(mBase, uint32(v41)+76)) = v520
	*(*int32)(unsafe.Add(mBase, uint32(v41)+80)) = v510
	*(*int32)(unsafe.Add(mBase, uint32(v41)+84)) = v509
	*(*int32)(unsafe.Add(mBase, uint32(v41)+88)) = v508
	*(*int32)(unsafe.Add(mBase, uint32(v41)+92)) = v504
	F_MemoryContextDeleteChildren(m, v505)
	mBase = m.M
	v791 = m.ExcPending
	if v791 != 0 {
		v1500 = v525
		goto L6
	} else {
		goto L82
	}
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v518))) = v506
	*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = int32(0)
	if v506 != 0 {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v795 = *(*int32)(unsafe.Add(mBase, uint32(v506)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v41)+32)) = v511
	*(*int32)(unsafe.Add(mBase, uint32(v41)+28)) = v512
	*(*int32)(unsafe.Add(mBase, uint32(v41)+36)) = v505
	*(*int32)(unsafe.Add(mBase, uint32(v41)+40)) = v515
	*(*int32)(unsafe.Add(mBase, uint32(v41)+44)) = v517
	*(*int32)(unsafe.Add(mBase, uint32(v41)+48)) = v513
	*(*int32)(unsafe.Add(mBase, uint32(v41)+52)) = v516
	*(*int32)(unsafe.Add(mBase, uint32(v41)+56)) = v506
	*(*int32)(unsafe.Add(mBase, uint32(v41)+60)) = v518
	*(*int32)(unsafe.Add(mBase, uint32(v41)+64)) = v514
	*(*int32)(unsafe.Add(mBase, uint32(v41)+68)) = v507
	*(*int32)(unsafe.Add(mBase, uint32(v41)+72)) = v519
	*(*int32)(unsafe.Add(mBase, uint32(v41)+76)) = v520
	*(*int32)(unsafe.Add(mBase, uint32(v41)+80)) = v510
	*(*int32)(unsafe.Add(mBase, uint32(v41)+84)) = v509
	*(*int32)(unsafe.Add(mBase, uint32(v41)+88)) = v508
	*(*int32)(unsafe.Add(mBase, uint32(v41)+92)) = v504
	F_MemoryContextReset(m, v795)
	mBase = m.M
	v814 = m.ExcPending
	if v814 != 0 {
		v1500 = v525
		goto L6
	} else {
		goto L86
	}
L84:
	;
	goto L85
L85:
	;
	v816 = *(*int32)(unsafe.Add(mBase, uint32(v519)))
	v817 = *(*int32)(unsafe.Add(mBase, uint32(v816)+8))
	if v817 == int32(0) {
		goto L87
	} else {
		goto L88
	}
L86:
	;
	goto L85
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v516))) = v513
	*(*int32)(unsafe.Add(mBase, uint32(v41)+32)) = v511
	*(*int32)(unsafe.Add(mBase, uint32(v41)+28)) = v512
	*(*int32)(unsafe.Add(mBase, uint32(v41)+36)) = v505
	*(*int32)(unsafe.Add(mBase, uint32(v41)+40)) = v515
	*(*int32)(unsafe.Add(mBase, uint32(v41)+44)) = v517
	*(*int32)(unsafe.Add(mBase, uint32(v41)+48)) = v513
	*(*int32)(unsafe.Add(mBase, uint32(v41)+52)) = v516
	*(*int32)(unsafe.Add(mBase, uint32(v41)+56)) = v506
	*(*int32)(unsafe.Add(mBase, uint32(v41)+60)) = v518
	*(*int32)(unsafe.Add(mBase, uint32(v41)+64)) = v514
	*(*int32)(unsafe.Add(mBase, uint32(v41)+68)) = v507
	*(*int32)(unsafe.Add(mBase, uint32(v41)+72)) = v519
	*(*int32)(unsafe.Add(mBase, uint32(v41)+76)) = v520
	*(*int32)(unsafe.Add(mBase, uint32(v41)+80)) = v510
	*(*int32)(unsafe.Add(mBase, uint32(v41)+84)) = v509
	*(*int32)(unsafe.Add(mBase, uint32(v41)+88)) = v508
	*(*int32)(unsafe.Add(mBase, uint32(v41)+92)) = v504
	F_ReThrowError(m, v726)
	mBase = m.M
	v1021 = m.ExcPending
	if v1021 != 0 {
		v1500 = v525
		goto L6
	} else {
		goto L110
	}
L88:
	;
	v820 = *(*int32)(unsafe.Add(mBase, uint32(v817)+4))
	if v820 <= int32(0) {
		goto L87
	} else {
		goto L89
	}
L89:
	;
	v823 = int32(0)
	if v823 < v820 {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v826 = v820
	goto L92
L91:
	;
	v826 = v823
	goto L92
L92:
	;
	v827 = *(*int32)(unsafe.Add(mBase, uint32(v817)+12))
	v859 = int32(0)
	goto L93
L93:
	;
	v870 = *(*int32)(unsafe.Add(mBase, uint32(v827+v859<<(uint(int32(2))%32))))
	v871 = *(*int32)(unsafe.Add(mBase, uint32(v870)+4))
	if v871 != 0 {
		goto L95
	} else {
		goto L96
	}
L94:
	;
	goto L87
L95:
	;
	v872 = *(*int32)(unsafe.Add(mBase, uint32(v726)+28))
	v882 = v871
	goto L98
L96:
	;
	goto L97
L97:
	;
	v962 = v859 + int32(1)
	if v962 != v826 {
		v859 = v962
		goto L93
	} else {
		goto L109
	}
L98:
	;
	v917 = *(*int32)(unsafe.Add(mBase, uint32(v882)))
	if v917 == int32(-1) {
		goto L101
	} else {
		goto L102
	}
L99:
	;
	goto L97
L100:
	;
	v922 = *(*int32)(unsafe.Add(mBase, uint32(v882)+8))
	if v922 != 0 {
		v882 = v922
		goto L98
	} else {
		goto L108
	}
L101:
	;
	if v872 == int32(67108896) {
		goto L100
	} else {
		goto L104
	}
L102:
	;
	goto L103
L103:
	;
	if v917 == v872 {
		goto L9
	} else {
		goto L106
	}
L104:
	;
	if v872 == int32(67371461) {
		goto L100
	} else {
		goto L105
	}
L105:
	;
	goto L9
L106:
	;
	if v917 == v872&int32(4095) {
		goto L9
	} else {
		goto L107
	}
L107:
	;
	goto L100
L108:
	;
	goto L99
L109:
	;
	goto L94
L110:
	;
	goto L3
L111:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v95))) = v1042
	v1307 = v95
	v1308 = v59
	v1309 = v60
	v1310 = v61
	v1311 = v98
	v1312 = v101
	v1313 = v104
	v1314 = v65
	v1315 = v66
	v1316 = v67
	v1317 = v68
	v1318 = v69
	v1319 = v70
	v1320 = v71
	v1321 = v72
	v1322 = v46
	v1323 = v44
	v1328 = v104
	goto L7
L112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+32)) = v511
	*(*int32)(unsafe.Add(mBase, uint32(v41)+28)) = v512
	*(*int32)(unsafe.Add(mBase, uint32(v41)+36)) = v505
	*(*int32)(unsafe.Add(mBase, uint32(v41)+40)) = v515
	*(*int32)(unsafe.Add(mBase, uint32(v41)+44)) = v517
	*(*int32)(unsafe.Add(mBase, uint32(v41)+48)) = v513
	*(*int32)(unsafe.Add(mBase, uint32(v41)+52)) = v516
	*(*int32)(unsafe.Add(mBase, uint32(v41)+56)) = v506
	*(*int32)(unsafe.Add(mBase, uint32(v41)+60)) = v518
	*(*int32)(unsafe.Add(mBase, uint32(v41)+64)) = v514
	*(*int32)(unsafe.Add(mBase, uint32(v41)+68)) = v507
	*(*int32)(unsafe.Add(mBase, uint32(v41)+72)) = v519
	*(*int32)(unsafe.Add(mBase, uint32(v41)+76)) = v520
	*(*int32)(unsafe.Add(mBase, uint32(v41)+80)) = v510
	*(*int32)(unsafe.Add(mBase, uint32(v41)+84)) = v509
	*(*int32)(unsafe.Add(mBase, uint32(v41)+88)) = v508
	*(*int32)(unsafe.Add(mBase, uint32(v41)+92)) = v504
	v1132 = F_cstring_to_text(m, v1073)
	mBase = m.M
	v1133 = m.ExcPending
	if v1133 != 0 {
		v1500 = v525
		goto L6
	} else {
		goto L113
	}
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+32)) = v511
	*(*int32)(unsafe.Add(mBase, uint32(v41)+28)) = v512
	*(*int32)(unsafe.Add(mBase, uint32(v41)+36)) = v505
	*(*int32)(unsafe.Add(mBase, uint32(v41)+40)) = v515
	*(*int32)(unsafe.Add(mBase, uint32(v41)+44)) = v517
	*(*int32)(unsafe.Add(mBase, uint32(v41)+48)) = v513
	*(*int32)(unsafe.Add(mBase, uint32(v41)+52)) = v516
	*(*int32)(unsafe.Add(mBase, uint32(v41)+56)) = v506
	*(*int32)(unsafe.Add(mBase, uint32(v41)+60)) = v518
	*(*int32)(unsafe.Add(mBase, uint32(v41)+64)) = v514
	*(*int32)(unsafe.Add(mBase, uint32(v41)+68)) = v507
	*(*int32)(unsafe.Add(mBase, uint32(v41)+72)) = v519
	*(*int32)(unsafe.Add(mBase, uint32(v41)+76)) = v520
	*(*int32)(unsafe.Add(mBase, uint32(v41)+80)) = v510
	*(*int32)(unsafe.Add(mBase, uint32(v41)+84)) = v509
	*(*int32)(unsafe.Add(mBase, uint32(v41)+88)) = v508
	*(*int32)(unsafe.Add(mBase, uint32(v41)+92)) = v504
	F_assign_simple_var(m, l0, v1055, v1132, int32(0), int32(1))
	mBase = m.M
	v1154 = m.ExcPending
	if v1154 != 0 {
		v1500 = v525
		goto L6
	} else {
		goto L114
	}
L114:
	;
	v1155 = *(*int32)(unsafe.Add(mBase, uint32(v726)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v41)+32)) = v511
	*(*int32)(unsafe.Add(mBase, uint32(v41)+28)) = v512
	*(*int32)(unsafe.Add(mBase, uint32(v41)+36)) = v505
	*(*int32)(unsafe.Add(mBase, uint32(v41)+40)) = v515
	*(*int32)(unsafe.Add(mBase, uint32(v41)+44)) = v517
	*(*int32)(unsafe.Add(mBase, uint32(v41)+48)) = v513
	*(*int32)(unsafe.Add(mBase, uint32(v41)+52)) = v516
	*(*int32)(unsafe.Add(mBase, uint32(v41)+56)) = v506
	*(*int32)(unsafe.Add(mBase, uint32(v41)+60)) = v518
	*(*int32)(unsafe.Add(mBase, uint32(v41)+64)) = v514
	*(*int32)(unsafe.Add(mBase, uint32(v41)+68)) = v507
	*(*int32)(unsafe.Add(mBase, uint32(v41)+72)) = v519
	*(*int32)(unsafe.Add(mBase, uint32(v41)+76)) = v520
	*(*int32)(unsafe.Add(mBase, uint32(v41)+80)) = v510
	*(*int32)(unsafe.Add(mBase, uint32(v41)+84)) = v509
	*(*int32)(unsafe.Add(mBase, uint32(v41)+88)) = v508
	*(*int32)(unsafe.Add(mBase, uint32(v41)+92)) = v504
	v1173 = F_cstring_to_text(m, v1155)
	mBase = m.M
	v1174 = m.ExcPending
	if v1174 != 0 {
		v1500 = v525
		goto L6
	} else {
		goto L115
	}
L115:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+32)) = v511
	*(*int32)(unsafe.Add(mBase, uint32(v41)+28)) = v512
	*(*int32)(unsafe.Add(mBase, uint32(v41)+36)) = v505
	*(*int32)(unsafe.Add(mBase, uint32(v41)+40)) = v515
	*(*int32)(unsafe.Add(mBase, uint32(v41)+44)) = v517
	*(*int32)(unsafe.Add(mBase, uint32(v41)+48)) = v513
	*(*int32)(unsafe.Add(mBase, uint32(v41)+52)) = v516
	*(*int32)(unsafe.Add(mBase, uint32(v41)+56)) = v506
	*(*int32)(unsafe.Add(mBase, uint32(v41)+60)) = v518
	*(*int32)(unsafe.Add(mBase, uint32(v41)+64)) = v514
	*(*int32)(unsafe.Add(mBase, uint32(v41)+68)) = v507
	*(*int32)(unsafe.Add(mBase, uint32(v41)+72)) = v519
	*(*int32)(unsafe.Add(mBase, uint32(v41)+76)) = v520
	*(*int32)(unsafe.Add(mBase, uint32(v41)+80)) = v510
	*(*int32)(unsafe.Add(mBase, uint32(v41)+84)) = v509
	*(*int32)(unsafe.Add(mBase, uint32(v41)+88)) = v508
	*(*int32)(unsafe.Add(mBase, uint32(v41)+92)) = v504
	F_assign_simple_var(m, l0, v1050, v1173, int32(0), int32(1))
	mBase = m.M
	v1195 = m.ExcPending
	if v1195 != 0 {
		v1500 = v525
		goto L6
	} else {
		goto L116
	}
L116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v516))) = v726
	*(*int32)(unsafe.Add(mBase, uint32(v520))) = int32(0)
	v1199 = *(*int32)(unsafe.Add(mBase, uint32(v870)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v41)+32)) = v511
	*(*int32)(unsafe.Add(mBase, uint32(v41)+28)) = v512
	*(*int32)(unsafe.Add(mBase, uint32(v41)+36)) = v505
	*(*int32)(unsafe.Add(mBase, uint32(v41)+40)) = v515
	*(*int32)(unsafe.Add(mBase, uint32(v41)+44)) = v517
	*(*int32)(unsafe.Add(mBase, uint32(v41)+48)) = v513
	*(*int32)(unsafe.Add(mBase, uint32(v41)+52)) = v516
	*(*int32)(unsafe.Add(mBase, uint32(v41)+56)) = v506
	*(*int32)(unsafe.Add(mBase, uint32(v41)+60)) = v518
	*(*int32)(unsafe.Add(mBase, uint32(v41)+64)) = v514
	*(*int32)(unsafe.Add(mBase, uint32(v41)+68)) = v507
	*(*int32)(unsafe.Add(mBase, uint32(v41)+72)) = v519
	*(*int32)(unsafe.Add(mBase, uint32(v41)+76)) = v520
	*(*int32)(unsafe.Add(mBase, uint32(v41)+80)) = v510
	*(*int32)(unsafe.Add(mBase, uint32(v41)+84)) = v509
	*(*int32)(unsafe.Add(mBase, uint32(v41)+88)) = v508
	*(*int32)(unsafe.Add(mBase, uint32(v41)+92)) = v504
	v1217 = F_exec_stmts(m, l0, v1199)
	mBase = m.M
	v1218 = m.ExcPending
	if v1218 != 0 {
		v1500 = v525
		goto L6
	} else {
		goto L117
	}
L117:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v504))) = v1217
	*(*int32)(unsafe.Add(mBase, uint32(v516))) = v513
	v1221 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	*(*int32)(unsafe.Add(mBase, uint32(v517))) = v1221
	*(*int32)(unsafe.Add(mBase, uint32(v41)+32)) = v511
	*(*int32)(unsafe.Add(mBase, uint32(v41)+28)) = v512
	*(*int32)(unsafe.Add(mBase, uint32(v41)+36)) = v505
	*(*int32)(unsafe.Add(mBase, uint32(v41)+40)) = v515
	*(*int32)(unsafe.Add(mBase, uint32(v41)+44)) = v517
	*(*int32)(unsafe.Add(mBase, uint32(v41)+48)) = v513
	*(*int32)(unsafe.Add(mBase, uint32(v41)+52)) = v516
	*(*int32)(unsafe.Add(mBase, uint32(v41)+56)) = v506
	*(*int32)(unsafe.Add(mBase, uint32(v41)+60)) = v518
	*(*int32)(unsafe.Add(mBase, uint32(v41)+64)) = v514
	*(*int32)(unsafe.Add(mBase, uint32(v41)+68)) = v507
	*(*int32)(unsafe.Add(mBase, uint32(v41)+72)) = v519
	*(*int32)(unsafe.Add(mBase, uint32(v41)+76)) = v520
	*(*int32)(unsafe.Add(mBase, uint32(v41)+80)) = v510
	*(*int32)(unsafe.Add(mBase, uint32(v41)+84)) = v509
	*(*int32)(unsafe.Add(mBase, uint32(v41)+88)) = v508
	*(*int32)(unsafe.Add(mBase, uint32(v41)+92)) = v504
	v1240 = *(*int32)(unsafe.Add(mBase, uint32(v1221)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+100)) = v1240
	*(*int32)(unsafe.Add(mBase, uint32(v41)+32)) = v511
	*(*int32)(unsafe.Add(mBase, uint32(v41)+28)) = v512
	*(*int32)(unsafe.Add(mBase, uint32(v41)+36)) = v505
	*(*int32)(unsafe.Add(mBase, uint32(v41)+40)) = v515
	*(*int32)(unsafe.Add(mBase, uint32(v41)+44)) = v517
	*(*int32)(unsafe.Add(mBase, uint32(v41)+48)) = v513
	*(*int32)(unsafe.Add(mBase, uint32(v41)+52)) = v516
	*(*int32)(unsafe.Add(mBase, uint32(v41)+56)) = v506
	*(*int32)(unsafe.Add(mBase, uint32(v41)+60)) = v518
	*(*int32)(unsafe.Add(mBase, uint32(v41)+64)) = v514
	*(*int32)(unsafe.Add(mBase, uint32(v41)+68)) = v507
	*(*int32)(unsafe.Add(mBase, uint32(v41)+72)) = v519
	*(*int32)(unsafe.Add(mBase, uint32(v41)+76)) = v520
	*(*int32)(unsafe.Add(mBase, uint32(v41)+80)) = v510
	*(*int32)(unsafe.Add(mBase, uint32(v41)+84)) = v509
	*(*int32)(unsafe.Add(mBase, uint32(v41)+88)) = v508
	*(*int32)(unsafe.Add(mBase, uint32(v41)+92)) = v504
	F_MemoryContextReset(m, v505)
	mBase = m.M
	v1260 = m.ExcPending
	if v1260 != 0 {
		v1500 = v525
		goto L6
	} else {
		goto L118
	}
L118:
	;
	goto L8
L119:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+32)) = v1314
	*(*int32)(unsafe.Add(mBase, uint32(v41)+28)) = v1315
	*(*int32)(unsafe.Add(mBase, uint32(v41)+36)) = v1308
	*(*int32)(unsafe.Add(mBase, uint32(v41)+40)) = v1318
	*(*int32)(unsafe.Add(mBase, uint32(v41)+44)) = v1320
	*(*int32)(unsafe.Add(mBase, uint32(v41)+48)) = v1316
	*(*int32)(unsafe.Add(mBase, uint32(v41)+52)) = v1319
	*(*int32)(unsafe.Add(mBase, uint32(v41)+56)) = v1309
	*(*int32)(unsafe.Add(mBase, uint32(v41)+60)) = v1321
	*(*int32)(unsafe.Add(mBase, uint32(v41)+64)) = v1317
	*(*int32)(unsafe.Add(mBase, uint32(v41)+68)) = v1310
	*(*int32)(unsafe.Add(mBase, uint32(v41)+72)) = v1322
	*(*int32)(unsafe.Add(mBase, uint32(v41)+76)) = v1323
	*(*int32)(unsafe.Add(mBase, uint32(v41)+80)) = v1313
	*(*int32)(unsafe.Add(mBase, uint32(v41)+84)) = v1312
	*(*int32)(unsafe.Add(mBase, uint32(v41)+88)) = v1311
	*(*int32)(unsafe.Add(mBase, uint32(v41)+92)) = v1307
	F_errstart_cold(m, int32(21), int32(578012))
	mBase = m.M
	v1430 = m.ExcPending
	if v1430 != 0 {
		v1500 = v1328
		goto L6
	} else {
		goto L137
	}
L120:
	;
	m.G0 = v41 + int32(96)
	return v1403
L121:
	;
	if base.Ui32(int32(3)) < base.Ui32(v1343) {
		goto L119
	} else {
		goto L124
	}
L122:
	;
	goto L123
L123:
	;
	v1351 = int32(1)
	v1352 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v1352 == int32(0) {
		v1403 = v1351
		goto L120
	} else {
		goto L126
	}
L124:
	;
	if v1343 == int32(1) {
		goto L119
	} else {
		goto L125
	}
L125:
	;
	v1350 = *(*int32)(unsafe.Add(mBase, uint32(v1307)))
	v1403 = v1350
	goto L120
L126:
	;
	v1355 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v1355 == int32(0) {
		v1403 = v1351
		goto L120
	} else {
		goto L127
	}
L127:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+32)) = v1314
	*(*int32)(unsafe.Add(mBase, uint32(v41)+28)) = v1315
	*(*int32)(unsafe.Add(mBase, uint32(v41)+36)) = v1308
	*(*int32)(unsafe.Add(mBase, uint32(v41)+40)) = v1318
	*(*int32)(unsafe.Add(mBase, uint32(v41)+44)) = v1320
	*(*int32)(unsafe.Add(mBase, uint32(v41)+48)) = v1316
	*(*int32)(unsafe.Add(mBase, uint32(v41)+52)) = v1319
	*(*int32)(unsafe.Add(mBase, uint32(v41)+56)) = v1309
	*(*int32)(unsafe.Add(mBase, uint32(v41)+60)) = v1321
	*(*int32)(unsafe.Add(mBase, uint32(v41)+64)) = v1317
	*(*int32)(unsafe.Add(mBase, uint32(v41)+68)) = v1310
	*(*int32)(unsafe.Add(mBase, uint32(v41)+72)) = v1322
	*(*int32)(unsafe.Add(mBase, uint32(v41)+76)) = v1323
	*(*int32)(unsafe.Add(mBase, uint32(v41)+80)) = v1313
	*(*int32)(unsafe.Add(mBase, uint32(v41)+84)) = v1312
	*(*int32)(unsafe.Add(mBase, uint32(v41)+88)) = v1311
	*(*int32)(unsafe.Add(mBase, uint32(v41)+92)) = v1307
	v1377 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1352))))
	v1378 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1355))))
	if v1378 == int32(0) {
		v1397 = v1377
		v1398 = v1378
		goto L129
	} else {
		goto L130
	}
L128:
	;
	if v1398-v1397 != 0 {
		v1403 = v1351
		goto L120
	} else {
		goto L136
	}
L129:
	;
	goto L128
L130:
	;
	if v1377 != v1378 {
		v1397 = v1377
		v1398 = v1378
		goto L129
	} else {
		goto L131
	}
L131:
	;
	v1382 = v1355
	v1383 = v1352
	goto L132
L132:
	;
	v1386 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1383)+1)))
	v1387 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1382)+1)))
	if v1387 == int32(0) {
		v1397 = v1386
		v1398 = v1387
		goto L129
	} else {
		goto L134
	}
L133:
	;
	v1397 = v1386
	v1398 = v1387
	goto L129
L134:
	;
	v1390 = int32(1)
	if v1386 == v1387 {
		v1382 = v1382 + v1390
		v1383 = v1383 + v1390
		goto L132
	} else {
		goto L135
	}
L135:
	;
	goto L133
L136:
	;
	v1400 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v1400
	v1403 = v1400
	goto L120
L137:
	;
	v1431 = *(*int32)(unsafe.Add(mBase, uint32(v1307)))
	*(*int32)(unsafe.Add(mBase, uint32(v41)+32)) = v1314
	*(*int32)(unsafe.Add(mBase, uint32(v41)+28)) = v1315
	*(*int32)(unsafe.Add(mBase, uint32(v41)+36)) = v1308
	*(*int32)(unsafe.Add(mBase, uint32(v41)+40)) = v1318
	*(*int32)(unsafe.Add(mBase, uint32(v41)+44)) = v1320
	*(*int32)(unsafe.Add(mBase, uint32(v41)+48)) = v1316
	*(*int32)(unsafe.Add(mBase, uint32(v41)+52)) = v1319
	*(*int32)(unsafe.Add(mBase, uint32(v41)+56)) = v1309
	*(*int32)(unsafe.Add(mBase, uint32(v41)+60)) = v1321
	*(*int32)(unsafe.Add(mBase, uint32(v41)+64)) = v1317
	*(*int32)(unsafe.Add(mBase, uint32(v41)+68)) = v1310
	*(*int32)(unsafe.Add(mBase, uint32(v41)+72)) = v1322
	*(*int32)(unsafe.Add(mBase, uint32(v41)+76)) = v1323
	*(*int32)(unsafe.Add(mBase, uint32(v41)+80)) = v1313
	*(*int32)(unsafe.Add(mBase, uint32(v41)+84)) = v1312
	*(*int32)(unsafe.Add(mBase, uint32(v41)+88)) = v1311
	*(*int32)(unsafe.Add(mBase, uint32(v41)+92)) = v1307
	*(*int32)(unsafe.Add(mBase, uint32(v41))) = v1431
	F_errmsg_internal(m, int32(506366), v41)
	mBase = m.M
	v1452 = m.ExcPending
	if v1452 != 0 {
		v1500 = v1328
		goto L6
	} else {
		goto L138
	}
L138:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+32)) = v1314
	*(*int32)(unsafe.Add(mBase, uint32(v41)+28)) = v1315
	*(*int32)(unsafe.Add(mBase, uint32(v41)+36)) = v1308
	*(*int32)(unsafe.Add(mBase, uint32(v41)+40)) = v1318
	*(*int32)(unsafe.Add(mBase, uint32(v41)+44)) = v1320
	*(*int32)(unsafe.Add(mBase, uint32(v41)+48)) = v1316
	*(*int32)(unsafe.Add(mBase, uint32(v41)+52)) = v1319
	*(*int32)(unsafe.Add(mBase, uint32(v41)+56)) = v1309
	*(*int32)(unsafe.Add(mBase, uint32(v41)+60)) = v1321
	*(*int32)(unsafe.Add(mBase, uint32(v41)+64)) = v1317
	*(*int32)(unsafe.Add(mBase, uint32(v41)+68)) = v1310
	*(*int32)(unsafe.Add(mBase, uint32(v41)+72)) = v1322
	*(*int32)(unsafe.Add(mBase, uint32(v41)+76)) = v1323
	*(*int32)(unsafe.Add(mBase, uint32(v41)+80)) = v1313
	*(*int32)(unsafe.Add(mBase, uint32(v41)+84)) = v1312
	*(*int32)(unsafe.Add(mBase, uint32(v41)+88)) = v1311
	*(*int32)(unsafe.Add(mBase, uint32(v41)+92)) = v1307
	F_errfinish(m, int32(520770), int32(1983), int32(329261))
	mBase = m.M
	v1474 = m.ExcPending
	if v1474 != 0 {
		v1500 = v1328
		goto L6
	} else {
		goto L139
	}
L139:
	;
	goto L5
L140:
	;
	v1518 = int32(v1514)
	m.G0 = v1500
	v1520 = *(*int32)(unsafe.Add(mBase, uint32(v1518)+4))
	v1521 = *(*int32)(unsafe.Add(mBase, uint32(v1518)))
	v1525 = *(*int32)(unsafe.Add(mBase, uint32(v1521)))
	if v41+int32(24) == v1525 {
		goto L143
	} else {
		goto L144
	}
L141:
	;
	m.ExcPending = 1
	goto L149
L142:
	;
	if v1528 != 0 {
		goto L146
	} else {
		goto L147
	}
L143:
	;
	v1527 = *(*int32)(unsafe.Add(mBase, uint32(v1521)+4))
	v1528 = v1527
	goto L145
L144:
	;
	v1528 = int32(0)
	goto L145
L145:
	;
	goto L142
L146:
	;
	v1529 = *(*int32)(unsafe.Add(mBase, uint32(v41)+92))
	v1530 = *(*int32)(unsafe.Add(mBase, uint32(v41)+88))
	v1531 = *(*int32)(unsafe.Add(mBase, uint32(v41)+84))
	v1532 = *(*int32)(unsafe.Add(mBase, uint32(v41)+80))
	v1533 = *(*int32)(unsafe.Add(mBase, uint32(v41)+76))
	v1534 = *(*int32)(unsafe.Add(mBase, uint32(v41)+72))
	v1535 = *(*int32)(unsafe.Add(mBase, uint32(v41)+68))
	v1536 = *(*int32)(unsafe.Add(mBase, uint32(v41)+64))
	v1537 = *(*int32)(unsafe.Add(mBase, uint32(v41)+60))
	v1538 = *(*int32)(unsafe.Add(mBase, uint32(v41)+56))
	v1539 = *(*int32)(unsafe.Add(mBase, uint32(v41)+52))
	v1540 = *(*int32)(unsafe.Add(mBase, uint32(v41)+48))
	v1541 = *(*int32)(unsafe.Add(mBase, uint32(v41)+44))
	v1542 = *(*int32)(unsafe.Add(mBase, uint32(v41)+40))
	v1543 = *(*int32)(unsafe.Add(mBase, uint32(v41)+36))
	v1544 = *(*int32)(unsafe.Add(mBase, uint32(v41)+32))
	v1545 = *(*int32)(unsafe.Add(mBase, uint32(v41)+28))
	v57 = v1520
	v58 = v1529
	v59 = v1543
	v60 = v1538
	v61 = v1535
	v62 = v1530
	v63 = v1531
	v64 = v1532
	v65 = v1544
	v66 = v1545
	v67 = v1540
	v68 = v1536
	v69 = v1542
	v70 = v1539
	v71 = v1541
	v72 = v1537
	v73 = v1534
	v74 = v1533
	v75 = v1528
	v79 = v1500
	goto L1
L147:
	;
	goto L148
L148:
	;
	F___wasm_longjmp(m, v1521, v1520)
	mBase = m.M
	v1549 = m.ExcPending
	if v1549 != 0 {
		goto L149
	} else {
		goto L150
	}
L149:
	;
	return int32(0)
L150:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
