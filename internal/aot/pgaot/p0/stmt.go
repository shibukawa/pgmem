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
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
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
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v90 int32
	_ = v90
	var v111 int32
	_ = v111
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v148 int64
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v188 int32
	_ = v188
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v255 int32
	_ = v255
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v292 int32
	_ = v292
	var v310 int32
	_ = v310
	var v325 int32
	_ = v325
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v366 int32
	_ = v366
	var v368 int32
	_ = v368
	var v372 int32
	_ = v372
	var v374 int32
	_ = v374
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v419 int32
	_ = v419
	var v423 int32
	_ = v423
	var v425 int32
	_ = v425
	var v427 int32
	_ = v427
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v488 int32
	_ = v488
	var v491 int32
	_ = v491
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v510 int32
	_ = v510
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v534 int32
	_ = v534
	var v535 int64
	_ = v535
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v551 int64
	_ = v551
	var v552 int32
	_ = v552
	var v568 int32
	_ = v568
	var v570 int32
	_ = v570
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v586 int32
	_ = v586
	var v588 int32
	_ = v588
	var v592 int32
	_ = v592
	var v597 int32
	_ = v597
	var v600 int32
	_ = v600
	var v603 int32
	_ = v603
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v649 int32
	_ = v649
	var v664 int32
	_ = v664
	var v686 int32
	_ = v686
	var v690 int32
	_ = v690
	var v705 int32
	_ = v705
	var v706 int32
	_ = v706
	var v707 int32
	_ = v707
	var v710 int32
	_ = v710
	var v713 int32
	_ = v713
	var v716 int32
	_ = v716
	var v717 int32
	_ = v717
	var v740 int32
	_ = v740
	var v756 int32
	_ = v756
	var v757 int32
	_ = v757
	var v758 int32
	_ = v758
	var v769 int32
	_ = v769
	var v799 int32
	_ = v799
	var v806 int32
	_ = v806
	var v842 int32
	_ = v842
	var v893 int32
	_ = v893
	var v896 int32
	_ = v896
	var v910 int32
	_ = v910
	var v911 int32
	_ = v911
	var v913 int32
	_ = v913
	var v914 int32
	_ = v914
	var v915 int32
	_ = v915
	var v918 int32
	_ = v918
	var v919 int32
	_ = v919
	var v923 int32
	_ = v923
	var v937 int32
	_ = v937
	var v938 int32
	_ = v938
	var v940 int32
	_ = v940
	var v941 int32
	_ = v941
	var v949 int32
	_ = v949
	var v957 int32
	_ = v957
	var v965 int32
	_ = v965
	var v973 int32
	_ = v973
	var v976 int32
	_ = v976
	var v992 int32
	_ = v992
	var v993 int32
	_ = v993
	var v1011 int32
	_ = v1011
	var v1012 int32
	_ = v1012
	var v1026 int32
	_ = v1026
	var v1027 int32
	_ = v1027
	var v1045 int32
	_ = v1045
	var v1049 int32
	_ = v1049
	var v1063 int32
	_ = v1063
	var v1064 int32
	_ = v1064
	var v1067 int32
	_ = v1067
	var v1082 int32
	_ = v1082
	var v1098 int32
	_ = v1098
	var v1140 int32
	_ = v1140
	var v1142 int32
	_ = v1142
	var v1143 int32
	_ = v1143
	var v1144 int32
	_ = v1144
	var v1145 int32
	_ = v1145
	var v1146 int32
	_ = v1146
	var v1147 int32
	_ = v1147
	var v1148 int32
	_ = v1148
	var v1149 int32
	_ = v1149
	var v1150 int32
	_ = v1150
	var v1151 int32
	_ = v1151
	var v1152 int32
	_ = v1152
	var v1153 int32
	_ = v1153
	var v1173 int32
	_ = v1173
	var v1181 int32
	_ = v1181
	var v1182 int32
	_ = v1182
	var v1183 int32
	_ = v1183
	var v1186 int32
	_ = v1186
	var v1204 int32
	_ = v1204
	var v1207 int32
	_ = v1207
	var v1210 int32
	_ = v1210
	var v1211 int32
	_ = v1211
	var v1214 int32
	_ = v1214
	var v1215 int32
	_ = v1215
	var v1218 int32
	_ = v1218
	var v1225 int32
	_ = v1225
	var v1226 int32
	_ = v1226
	var v1228 int32
	_ = v1228
	var v1231 int32
	_ = v1231
	var v1254 int32
	_ = v1254
	var v1268 int32
	_ = v1268
	var v1272 int32
	_ = v1272
	var v1290 int32
	_ = v1290
	var v1325 int32
	_ = v1325
	var v1326 int64
	_ = v1326
	var v1330 int32
	_ = v1330
	var v1332 int32
	_ = v1332
	var v1333 int32
	_ = v1333
	var v1336 int32
	_ = v1336
	var v1338 int32
	_ = v1338
	var v1340 int32
	_ = v1340
	var v1341 int32
	_ = v1341
	var v1342 int32
	_ = v1342
	var v1343 int32
	_ = v1343
	var v1344 int32
	_ = v1344
	var v1345 int32
	_ = v1345
	var v1346 int32
	_ = v1346
	var v1347 int32
	_ = v1347
	var v1348 int32
	_ = v1348
	var v1349 int32
	_ = v1349
	var v1350 int32
	_ = v1350
	var v1351 int32
	_ = v1351
	var v1352 int32
	_ = v1352
	var v1353 int32
	_ = v1353
	var v1357 int32
	_ = v1357
	v3 = int32(0)
	v35 = m.G0
	v37 = v35 - int32(256)
	m.G0 = v37
	v40 = l0 + int32(140)
	v42 = l1 + int32(28)
	v44 = l0 + int32(104)
	v46 = l0 + int32(44)
	v48 = l0 + int32(128)
	v53 = v3
	v54 = v3
	v55 = v3
	v56 = v3
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
	v67 = int32(-1)
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
	if v67 != int32(1) {
		goto L11
	} else {
		goto L12
	}
L5:
	;
	goto L3
L6:
	;
	v1325 = int32(m.ExcTag)
	v1326 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v1325 == int32(0) {
		goto L136
	} else {
		goto L137
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1140))) = int32(0)
	v1173 = *(*int32)(unsafe.Add(mBase, uint32(v37)+200))
	if v1173 != int32(1) {
		goto L119
	} else {
		goto L120
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_exec_stmt_block[0])) = v444
	*(*int32)(unsafe.Add(mBase, _c_F_exec_stmt_block[1])) = v443
	v1140 = v437
	v1142 = v439
	v1143 = v440
	v1144 = v441
	v1145 = v442
	v1146 = v443
	v1147 = v444
	v1148 = v445
	v1149 = v446
	v1150 = v447
	v1151 = v448
	v1152 = v449
	v1153 = v450
	goto L7
L9:
	;
	v913 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v914 = *(*int32)(unsafe.Add(mBase, uint32(v706)+4))
	v915 = int32(2)
	v918 = *(*int32)(unsafe.Add(mBase, uint32(v913+v914<<(uint(v915)%32))))
	v919 = *(*int32)(unsafe.Add(mBase, uint32(v706)))
	v923 = *(*int32)(unsafe.Add(mBase, uint32(v913+v919<<(uint(v915)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v37)+208)) = v443
	*(*int32)(unsafe.Add(mBase, uint32(v37)+204)) = v444
	*(*int32)(unsafe.Add(mBase, uint32(v37)+212)) = v440
	*(*int32)(unsafe.Add(mBase, uint32(v37)+216)) = v447
	*(*int32)(unsafe.Add(mBase, uint32(v37)+220)) = v449
	*(*int32)(unsafe.Add(mBase, uint32(v37)+224)) = v445
	*(*int32)(unsafe.Add(mBase, uint32(v37)+228)) = v448
	*(*int32)(unsafe.Add(mBase, uint32(v37)+232)) = v441
	*(*int32)(unsafe.Add(mBase, uint32(v37)+236)) = v450
	*(*int32)(unsafe.Add(mBase, uint32(v37)+240)) = v446
	*(*int32)(unsafe.Add(mBase, uint32(v37)+244)) = v442
	*(*int32)(unsafe.Add(mBase, uint32(v37)+248)) = v439
	*(*int32)(unsafe.Add(mBase, uint32(v37)+252)) = v437
	v937 = int32(_a_F_exec_stmt_block_0)
	v938 = int32(63)
	v940 = int32(48)
	v941 = v758&v938 + v940
	*(*uint8)(unsafe.Add(mBase, _c_F_exec_stmt_block[2])) = uint8(v941)
	v949 = int32(base.Ui32(v758)>>(uint(int32(24))%32))&v938 + v940
	*(*uint8)(unsafe.Add(mBase, _c_F_exec_stmt_block[3])) = uint8(v949)
	v957 = int32(base.Ui32(v758)>>(uint(int32(18))%32))&v938 + v940
	*(*uint8)(unsafe.Add(mBase, _c_F_exec_stmt_block[4])) = uint8(v957)
	v965 = int32(base.Ui32(v758)>>(uint(int32(12))%32))&v938 + v940
	*(*uint8)(unsafe.Add(mBase, _c_F_exec_stmt_block[5])) = uint8(v965)
	v973 = int32(base.Ui32(v758)>>(uint(int32(6))%32))&v938 + v940
	*(*uint8)(unsafe.Add(mBase, _c_F_exec_stmt_block[6])) = uint8(v973)
	v976 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_exec_stmt_block[7])) = uint8(v976)
	goto L110
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = int32(0)
	v896 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v37)+208)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v37)+204)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v37)+212)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v37)+216)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v37)+220)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v37)+224)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v37)+228)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v37)+232)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v37)+236)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v37)+240)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v37)+244)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v37)+248)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v37)+252)) = v40
	v910 = F_exec_stmts(m, l0, v896)
	mBase = m.M
	v911 = m.ExcPending
	if v911 != 0 {
		goto L6
	} else {
		goto L109
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+200)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = int32(_a_F_exec_stmt_block_1)
	v90 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if int32(0) < v90 {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	v437 = v53
	v438 = v54
	v439 = v55
	v440 = v56
	v441 = v57
	v442 = v58
	v443 = v59
	v444 = v60
	v445 = v61
	v446 = v62
	v447 = v63
	v448 = v64
	v449 = v65
	v450 = v66
	goto L13
L13:
	;
	if v438 == int32(0) {
		goto L57
	} else {
		goto L58
	}
L14:
	;
	v111 = int32(0)
	goto L17
L15:
	;
	goto L16
L16:
	;
	v366 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+136)) = v366
	v368 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	if v368 == v366 {
		goto L10
	} else {
		goto L46
	}
L17:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v129 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v130 = int32(2)
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v129+v111<<(uint(v130)%32))))
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v128+v133<<(uint(v130)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+136)) = v137
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v137)))
	switch v139 {
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
	v329 = v111 + int32(1)
	v330 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v329 < v330 {
		v111 = v329
		goto L17
	} else {
		goto L45
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+208)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v37)+204)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v37)+212)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v37)+216)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v37)+220)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v37)+224)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v37)+228)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v37)+232)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v37)+236)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v37)+240)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v37)+244)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v37)+248)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v37)+252)) = v40
	F_exec_assign_expr(m, l0, v137, v197)
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L6
	} else {
		goto L44
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+208)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v37)+204)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v37)+212)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v37)+216)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v37)+220)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v37)+224)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v37)+228)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v37)+232)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v37)+236)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v37)+240)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v37)+244)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v37)+248)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v37)+252)) = v40
	F_errstart_cold(m, int32(21), int32(_a_F_exec_stmt_block_2))
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L6
	} else {
		goto L41
	}
L22:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v137)+20))
	if v221 == int32(0) {
		goto L36
	} else {
		goto L37
	}
L23:
	;
	v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v137)+49)))
	if v140 != int32(1) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v137)+52)) = int32(0)
	v193 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v137)+48)) = uint16(v193)
	*(*int64)(unsafe.Add(mBase, uint32(v137)+40)) = int64(0)
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v137)+20))
	if v197 != 0 {
		goto L20
	} else {
		goto L33
	}
L25:
	;
	v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v137)+48)))
	if v143 != 0 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v137)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v37)+208)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v37)+204)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v37)+212)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v37)+216)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v37)+220)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v37)+224)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v37)+228)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v37)+232)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v37)+236)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v37)+240)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v37)+244)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v37)+248)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v37)+252)) = v40
	F_pfree(m, v173)
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L6
	} else {
		goto L32
	}
L27:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v137)+24))
	v145 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v144)+12)))
	if v145 != int32(_a_F_exec_stmt_block_3) {
		goto L26
	} else {
		goto L28
	}
L28:
	;
	v148 = *(*int64)(unsafe.Add(mBase, uint32(v137)+40))
	v149 = base.I32_wrap_i64(v148)
	v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149))))
	if v150 != int32(1) {
		goto L26
	} else {
		goto L29
	}
L29:
	;
	v153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149)+1)))
	if v153 != int32(3) {
		goto L26
	} else {
		goto L30
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+208)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v37)+204)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v37)+212)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v37)+216)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v37)+220)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v37)+224)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v37)+228)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v37)+232)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v37)+236)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v37)+240)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v37)+244)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v37)+248)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v37)+252)) = v40
	F_DeleteExpandedObject(m, v148)
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
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
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v137)+24))
	v199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v198)+15)))
	if v199 != int32(100) {
		goto L19
	} else {
		goto L34
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+208)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v37)+204)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v37)+212)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v37)+216)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v37)+220)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v37)+224)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v37)+228)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v37)+232)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v37)+236)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v37)+240)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v37)+244)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v37)+248)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v37)+252)) = v40
	F_exec_assign_value(m, l0, v137, int64(0), int32(1), int32(705), int32(-1))
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L6
	} else {
		goto L35
	}
L35:
	;
	goto L19
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+208)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v37)+204)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v37)+212)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v37)+216)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v37)+220)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v37)+224)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v37)+228)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v37)+232)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v37)+236)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v37)+240)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v37)+244)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v37)+248)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v37)+252)) = v40
	v237 = int32(0)
	F_exec_move_row(m, l0, v137, v237, v237)
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L6
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+208)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v37)+204)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v37)+212)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v37)+216)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v37)+220)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v37)+224)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v37)+228)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v37)+232)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v37)+236)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v37)+240)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v37)+244)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v37)+248)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v37)+252)) = v40
	F_exec_assign_expr(m, l0, v137, v221)
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
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
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v137)))
	*(*int32)(unsafe.Add(mBase, uint32(v37)+208)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v37)+204)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v37)+212)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v37)+216)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v37)+220)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v37)+224)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v37)+228)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v37)+232)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v37)+236)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v37)+240)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v37)+244)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v37)+248)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v37)+252)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v37)+16)) = v273
	F_errmsg_internal(m, int32(_a_F_exec_stmt_block_4), v37+int32(16))
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L6
	} else {
		goto L42
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+208)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v37)+204)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v37)+212)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v37)+216)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v37)+220)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v37)+224)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v37)+228)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v37)+232)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v37)+236)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v37)+240)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v37)+244)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v37)+248)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v37)+252)) = v40
	F_errfinish(m, int32(_a_F_exec_stmt_block_5), int32(1790), int32(_a_F_exec_stmt_block_6))
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
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
	v372 = *(*int32)(unsafe.Add(mBase, _c_F_exec_stmt_block[8]))
	v374 = *(*int32)(unsafe.Add(mBase, _c_F_exec_stmt_block[9]))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = int32(_a_F_exec_stmt_block_7)
	v377 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v378 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	v379 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v379 != 0 {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+208)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v37)+204)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v37)+212)) = v403
	*(*int32)(unsafe.Add(mBase, uint32(v37)+216)) = v402
	*(*int32)(unsafe.Add(mBase, uint32(v37)+220)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v37)+224)) = v377
	*(*int32)(unsafe.Add(mBase, uint32(v37)+228)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v37)+232)) = v378
	*(*int32)(unsafe.Add(mBase, uint32(v37)+236)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v37)+240)) = v372
	*(*int32)(unsafe.Add(mBase, uint32(v37)+244)) = v374
	*(*int32)(unsafe.Add(mBase, uint32(v37)+248)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v37)+252)) = v40
	F_BeginInternalSubTransaction(m, int32(0))
	mBase = m.M
	v419 = m.ExcPending
	if v419 != 0 {
		goto L6
	} else {
		goto L52
	}
L48:
	;
	v402 = v63
	v403 = v379
	goto L47
L49:
	;
	goto L50
L50:
	;
	v380 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
	*(*int32)(unsafe.Add(mBase, uint32(v37)+208)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v37)+204)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v37)+212)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v37)+216)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v37)+220)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v37)+224)) = v377
	*(*int32)(unsafe.Add(mBase, uint32(v37)+228)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v37)+232)) = v378
	*(*int32)(unsafe.Add(mBase, uint32(v37)+236)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v37)+240)) = v372
	*(*int32)(unsafe.Add(mBase, uint32(v37)+244)) = v374
	*(*int32)(unsafe.Add(mBase, uint32(v37)+248)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v37)+252)) = v40
	v398 = F_AllocSetContextCreateInternal(m, v380, int32(_a_F_exec_stmt_block_8), int32(0), int32(_a_F_exec_stmt_block_9), int32(_a_F_exec_stmt_block_10))
	mBase = m.M
	v399 = m.ExcPending
	if v399 != 0 {
		goto L6
	} else {
		goto L51
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44))) = v398
	v402 = v398
	v403 = v398
	goto L47
L52:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_exec_stmt_block[9])) = v374
	v423 = *(*int32)(unsafe.Add(mBase, _c_F_exec_stmt_block[1]))
	v425 = *(*int32)(unsafe.Add(mBase, _c_F_exec_stmt_block[0]))
	goto L53
L53:
	;
	v427 = v37 + int32(32)
	*(*int32)(unsafe.Add(mBase, uint32(v427)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v427))) = v37 + int32(24)
	goto L56
L54:
	;
	v437 = v40
	v438 = int32(0)
	v439 = v42
	v440 = v403
	v441 = v378
	v442 = v374
	v443 = v423
	v444 = v425
	v445 = v377
	v446 = v372
	v447 = v402
	v448 = v46
	v449 = v44
	v450 = v48
	goto L13
L56:
	;
	goto L54
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+204)) = v444
	*(*int32)(unsafe.Add(mBase, _c_F_exec_stmt_block[1])) = v37 + int32(32)
	*(*int32)(unsafe.Add(mBase, uint32(v37)+208)) = v443
	*(*int32)(unsafe.Add(mBase, uint32(v37)+212)) = v440
	*(*int32)(unsafe.Add(mBase, uint32(v37)+216)) = v447
	*(*int32)(unsafe.Add(mBase, uint32(v37)+220)) = v449
	*(*int32)(unsafe.Add(mBase, uint32(v37)+224)) = v445
	*(*int32)(unsafe.Add(mBase, uint32(v37)+228)) = v448
	*(*int32)(unsafe.Add(mBase, uint32(v37)+232)) = v441
	*(*int32)(unsafe.Add(mBase, uint32(v37)+236)) = v450
	*(*int32)(unsafe.Add(mBase, uint32(v37)+240)) = v446
	*(*int32)(unsafe.Add(mBase, uint32(v37)+244)) = v442
	*(*int32)(unsafe.Add(mBase, uint32(v37)+248)) = v439
	*(*int32)(unsafe.Add(mBase, uint32(v37)+252)) = v437
	F_plpgsql_create_econtext(m, l0)
	mBase = m.M
	v488 = m.ExcPending
	if v488 != 0 {
		goto L6
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_exec_stmt_block[0])) = v444
	*(*int32)(unsafe.Add(mBase, _c_F_exec_stmt_block[1])) = v443
	*(*int32)(unsafe.Add(mBase, uint32(v437))) = int32(_a_F_exec_stmt_block_11)
	*(*int32)(unsafe.Add(mBase, _c_F_exec_stmt_block[9])) = v440
	*(*int32)(unsafe.Add(mBase, uint32(v37)+204)) = v444
	*(*int32)(unsafe.Add(mBase, uint32(v37)+208)) = v443
	*(*int32)(unsafe.Add(mBase, uint32(v37)+212)) = v440
	*(*int32)(unsafe.Add(mBase, uint32(v37)+216)) = v447
	*(*int32)(unsafe.Add(mBase, uint32(v37)+220)) = v449
	*(*int32)(unsafe.Add(mBase, uint32(v37)+224)) = v445
	*(*int32)(unsafe.Add(mBase, uint32(v37)+228)) = v448
	*(*int32)(unsafe.Add(mBase, uint32(v37)+232)) = v441
	*(*int32)(unsafe.Add(mBase, uint32(v37)+236)) = v450
	*(*int32)(unsafe.Add(mBase, uint32(v37)+240)) = v446
	*(*int32)(unsafe.Add(mBase, uint32(v37)+244)) = v442
	*(*int32)(unsafe.Add(mBase, uint32(v37)+248)) = v439
	*(*int32)(unsafe.Add(mBase, uint32(v37)+252)) = v437
	v633 = F_CopyErrorData(m)
	mBase = m.M
	v634 = m.ExcPending
	if v634 != 0 {
		goto L6
	} else {
		goto L79
	}
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v437))) = int32(0)
	v491 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v37)+208)) = v443
	*(*int32)(unsafe.Add(mBase, uint32(v37)+204)) = v444
	*(*int32)(unsafe.Add(mBase, uint32(v37)+212)) = v440
	*(*int32)(unsafe.Add(mBase, uint32(v37)+216)) = v447
	*(*int32)(unsafe.Add(mBase, uint32(v37)+220)) = v449
	*(*int32)(unsafe.Add(mBase, uint32(v37)+224)) = v445
	*(*int32)(unsafe.Add(mBase, uint32(v37)+228)) = v448
	*(*int32)(unsafe.Add(mBase, uint32(v37)+232)) = v441
	*(*int32)(unsafe.Add(mBase, uint32(v37)+236)) = v450
	*(*int32)(unsafe.Add(mBase, uint32(v37)+240)) = v446
	*(*int32)(unsafe.Add(mBase, uint32(v37)+244)) = v442
	*(*int32)(unsafe.Add(mBase, uint32(v37)+248)) = v439
	*(*int32)(unsafe.Add(mBase, uint32(v37)+252)) = v437
	v505 = F_exec_stmts(m, l0, v491)
	mBase = m.M
	v506 = m.ExcPending
	if v506 != 0 {
		goto L6
	} else {
		goto L61
	}
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+200)) = v505
	*(*int32)(unsafe.Add(mBase, uint32(v437))) = int32(_a_F_exec_stmt_block_12)
	v510 = *(*int32)(unsafe.Add(mBase, uint32(v37)+200))
	if v510 != int32(2) {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+208)) = v443
	*(*int32)(unsafe.Add(mBase, uint32(v37)+204)) = v444
	*(*int32)(unsafe.Add(mBase, uint32(v37)+212)) = v440
	*(*int32)(unsafe.Add(mBase, uint32(v37)+216)) = v447
	*(*int32)(unsafe.Add(mBase, uint32(v37)+220)) = v449
	*(*int32)(unsafe.Add(mBase, uint32(v37)+224)) = v445
	*(*int32)(unsafe.Add(mBase, uint32(v37)+228)) = v448
	*(*int32)(unsafe.Add(mBase, uint32(v37)+232)) = v441
	*(*int32)(unsafe.Add(mBase, uint32(v37)+236)) = v450
	*(*int32)(unsafe.Add(mBase, uint32(v37)+240)) = v446
	*(*int32)(unsafe.Add(mBase, uint32(v37)+244)) = v442
	*(*int32)(unsafe.Add(mBase, uint32(v37)+248)) = v439
	*(*int32)(unsafe.Add(mBase, uint32(v37)+252)) = v437
	v568 = m.G0
	v570 = v568 - int32(16)
	m.G0 = v570
	v573 = *(*int32)(unsafe.Add(mBase, _c_F_exec_stmt_block[10]))
	v574 = *(*int32)(unsafe.Add(mBase, uint32(v573)+24))
	if v574 != int32(12) {
		goto L68
	} else {
		goto L69
	}
L63:
	;
	v513 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+37)))
	if v513 != 0 {
		goto L62
	} else {
		goto L64
	}
L64:
	;
	v514 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v514 != 0 {
		goto L62
	} else {
		goto L65
	}
L65:
	;
	v515 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v37)+208)) = v443
	*(*int32)(unsafe.Add(mBase, uint32(v37)+204)) = v444
	*(*int32)(unsafe.Add(mBase, uint32(v37)+212)) = v440
	*(*int32)(unsafe.Add(mBase, uint32(v37)+216)) = v447
	*(*int32)(unsafe.Add(mBase, uint32(v37)+220)) = v449
	*(*int32)(unsafe.Add(mBase, uint32(v37)+224)) = v445
	*(*int32)(unsafe.Add(mBase, uint32(v37)+228)) = v448
	*(*int32)(unsafe.Add(mBase, uint32(v37)+232)) = v441
	*(*int32)(unsafe.Add(mBase, uint32(v37)+236)) = v450
	*(*int32)(unsafe.Add(mBase, uint32(v37)+240)) = v446
	*(*int32)(unsafe.Add(mBase, uint32(v37)+244)) = v442
	*(*int32)(unsafe.Add(mBase, uint32(v37)+248)) = v439
	*(*int32)(unsafe.Add(mBase, uint32(v37)+252)) = v437
	F_get_typlenbyval(m, v515, v37+int32(30), v37+int32(29))
	mBase = m.M
	v534 = m.ExcPending
	if v534 != 0 {
		goto L6
	} else {
		goto L66
	}
L66:
	;
	v535 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v37)+204)) = v444
	*(*int32)(unsafe.Add(mBase, uint32(v37)+208)) = v443
	*(*int32)(unsafe.Add(mBase, uint32(v37)+212)) = v440
	*(*int32)(unsafe.Add(mBase, uint32(v37)+216)) = v447
	*(*int32)(unsafe.Add(mBase, uint32(v37)+220)) = v449
	*(*int32)(unsafe.Add(mBase, uint32(v37)+224)) = v445
	*(*int32)(unsafe.Add(mBase, uint32(v37)+228)) = v448
	*(*int32)(unsafe.Add(mBase, uint32(v37)+232)) = v441
	*(*int32)(unsafe.Add(mBase, uint32(v37)+236)) = v450
	*(*int32)(unsafe.Add(mBase, uint32(v37)+240)) = v446
	*(*int32)(unsafe.Add(mBase, uint32(v37)+244)) = v442
	*(*int32)(unsafe.Add(mBase, uint32(v37)+248)) = v439
	*(*int32)(unsafe.Add(mBase, uint32(v37)+252)) = v437
	v549 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+29)))
	v550 = int32(*(*int16)(unsafe.Add(mBase, uint32(v37)+30)))
	v551 = F_datumTransfer(m, v535, v549, v550)
	mBase = m.M
	v552 = m.ExcPending
	if v552 != 0 {
		goto L6
	} else {
		goto L67
	}
L67:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v551
	goto L62
L68:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v580 = m.ExcPending
	if v580 != 0 {
		goto L6
	} else {
		goto L71
	}
L69:
	;
	goto L70
L70:
	;
	v600 = *(*int32)(unsafe.Add(mBase, _c_F_exec_stmt_block[11]))
	*(*int32)(unsafe.Add(mBase, _c_F_exec_stmt_block[9])) = v600
	F_CommitSubTransaction(m)
	mBase = m.M
	v603 = m.ExcPending
	if v603 != 0 {
		goto L6
	} else {
		goto L78
	}
L71:
	;
	v581 = *(*int32)(unsafe.Add(mBase, uint32(v573)+24))
	if base.Ui32(v581) <= base.Ui32(int32(19)) {
		goto L73
	} else {
		goto L74
	}
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v570))) = v588
	F_errmsg_internal(m, int32(_a_F_exec_stmt_block_13), v570)
	mBase = m.M
	v592 = m.ExcPending
	if v592 != 0 {
		goto L6
	} else {
		goto L76
	}
L73:
	;
	v586 = *(*int32)(unsafe.Add(mBase, uint32(v581<<(uint(int32(2))%32))+uint32(_c_F_exec_stmt_block[12])))
	v588 = v586
	goto L75
L74:
	;
	v588 = int32(_a_F_exec_stmt_block_14)
	goto L75
L75:
	;
	goto L72
L76:
	;
	F_errfinish(m, int32(_a_F_exec_stmt_block_15), int32(_a_F_exec_stmt_block_16), int32(_a_F_exec_stmt_block_17))
	mBase = m.M
	v597 = m.ExcPending
	if v597 != 0 {
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
	m.G0 = v570 + int32(16)
	*(*int32)(unsafe.Add(mBase, _c_F_exec_stmt_block[8])) = v446
	*(*int32)(unsafe.Add(mBase, _c_F_exec_stmt_block[9])) = v442
	*(*int32)(unsafe.Add(mBase, uint32(v450))) = v441
	goto L8
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+208)) = v443
	*(*int32)(unsafe.Add(mBase, uint32(v37)+204)) = v444
	*(*int32)(unsafe.Add(mBase, uint32(v37)+212)) = v440
	*(*int32)(unsafe.Add(mBase, uint32(v37)+216)) = v447
	*(*int32)(unsafe.Add(mBase, uint32(v37)+220)) = v449
	*(*int32)(unsafe.Add(mBase, uint32(v37)+224)) = v445
	*(*int32)(unsafe.Add(mBase, uint32(v37)+228)) = v448
	*(*int32)(unsafe.Add(mBase, uint32(v37)+232)) = v441
	*(*int32)(unsafe.Add(mBase, uint32(v37)+236)) = v450
	*(*int32)(unsafe.Add(mBase, uint32(v37)+240)) = v446
	*(*int32)(unsafe.Add(mBase, uint32(v37)+244)) = v442
	*(*int32)(unsafe.Add(mBase, uint32(v37)+248)) = v439
	*(*int32)(unsafe.Add(mBase, uint32(v37)+252)) = v437
	F_FlushErrorState(m)
	mBase = m.M
	v649 = m.ExcPending
	if v649 != 0 {
		goto L6
	} else {
		goto L80
	}
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+208)) = v443
	*(*int32)(unsafe.Add(mBase, uint32(v37)+204)) = v444
	*(*int32)(unsafe.Add(mBase, uint32(v37)+212)) = v440
	*(*int32)(unsafe.Add(mBase, uint32(v37)+216)) = v447
	*(*int32)(unsafe.Add(mBase, uint32(v37)+220)) = v449
	*(*int32)(unsafe.Add(mBase, uint32(v37)+224)) = v445
	*(*int32)(unsafe.Add(mBase, uint32(v37)+228)) = v448
	*(*int32)(unsafe.Add(mBase, uint32(v37)+232)) = v441
	*(*int32)(unsafe.Add(mBase, uint32(v37)+236)) = v450
	*(*int32)(unsafe.Add(mBase, uint32(v37)+240)) = v446
	*(*int32)(unsafe.Add(mBase, uint32(v37)+244)) = v442
	*(*int32)(unsafe.Add(mBase, uint32(v37)+248)) = v439
	*(*int32)(unsafe.Add(mBase, uint32(v37)+252)) = v437
	F_RollbackAndReleaseCurrentSubTransaction(m)
	mBase = m.M
	v664 = m.ExcPending
	if v664 != 0 {
		goto L6
	} else {
		goto L81
	}
L81:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_exec_stmt_block[8])) = v446
	*(*int32)(unsafe.Add(mBase, _c_F_exec_stmt_block[9])) = v442
	*(*int32)(unsafe.Add(mBase, uint32(l0)+108)) = v440
	*(*int32)(unsafe.Add(mBase, uint32(v449))) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v37)+208)) = v443
	*(*int32)(unsafe.Add(mBase, uint32(v37)+204)) = v444
	*(*int32)(unsafe.Add(mBase, uint32(v37)+212)) = v440
	*(*int32)(unsafe.Add(mBase, uint32(v37)+216)) = v447
	*(*int32)(unsafe.Add(mBase, uint32(v37)+220)) = v449
	*(*int32)(unsafe.Add(mBase, uint32(v37)+224)) = v445
	*(*int32)(unsafe.Add(mBase, uint32(v37)+228)) = v448
	*(*int32)(unsafe.Add(mBase, uint32(v37)+232)) = v441
	*(*int32)(unsafe.Add(mBase, uint32(v37)+236)) = v450
	*(*int32)(unsafe.Add(mBase, uint32(v37)+240)) = v446
	*(*int32)(unsafe.Add(mBase, uint32(v37)+244)) = v442
	*(*int32)(unsafe.Add(mBase, uint32(v37)+248)) = v439
	*(*int32)(unsafe.Add(mBase, uint32(v37)+252)) = v437
	F_MemoryContextDeleteChildren(m, v440)
	mBase = m.M
	v686 = m.ExcPending
	if v686 != 0 {
		goto L6
	} else {
		goto L82
	}
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v450))) = v441
	*(*int32)(unsafe.Add(mBase, uint32(l0)+112)) = int32(0)
	if v441 != 0 {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v690 = *(*int32)(unsafe.Add(mBase, uint32(v441)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v37)+208)) = v443
	*(*int32)(unsafe.Add(mBase, uint32(v37)+204)) = v444
	*(*int32)(unsafe.Add(mBase, uint32(v37)+212)) = v440
	*(*int32)(unsafe.Add(mBase, uint32(v37)+216)) = v447
	*(*int32)(unsafe.Add(mBase, uint32(v37)+220)) = v449
	*(*int32)(unsafe.Add(mBase, uint32(v37)+224)) = v445
	*(*int32)(unsafe.Add(mBase, uint32(v37)+228)) = v448
	*(*int32)(unsafe.Add(mBase, uint32(v37)+232)) = v441
	*(*int32)(unsafe.Add(mBase, uint32(v37)+236)) = v450
	*(*int32)(unsafe.Add(mBase, uint32(v37)+240)) = v446
	*(*int32)(unsafe.Add(mBase, uint32(v37)+244)) = v442
	*(*int32)(unsafe.Add(mBase, uint32(v37)+248)) = v439
	*(*int32)(unsafe.Add(mBase, uint32(v37)+252)) = v437
	F_MemoryContextReset(m, v690)
	mBase = m.M
	v705 = m.ExcPending
	if v705 != 0 {
		goto L6
	} else {
		goto L86
	}
L84:
	;
	goto L85
L85:
	;
	v706 = *(*int32)(unsafe.Add(mBase, uint32(v439)))
	v707 = *(*int32)(unsafe.Add(mBase, uint32(v706)+8))
	if v707 == int32(0) {
		goto L87
	} else {
		goto L88
	}
L86:
	;
	goto L85
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v448))) = v445
	*(*int32)(unsafe.Add(mBase, uint32(v37)+208)) = v443
	*(*int32)(unsafe.Add(mBase, uint32(v37)+204)) = v444
	*(*int32)(unsafe.Add(mBase, uint32(v37)+212)) = v440
	*(*int32)(unsafe.Add(mBase, uint32(v37)+216)) = v447
	*(*int32)(unsafe.Add(mBase, uint32(v37)+220)) = v449
	*(*int32)(unsafe.Add(mBase, uint32(v37)+224)) = v445
	*(*int32)(unsafe.Add(mBase, uint32(v37)+228)) = v448
	*(*int32)(unsafe.Add(mBase, uint32(v37)+232)) = v441
	*(*int32)(unsafe.Add(mBase, uint32(v37)+236)) = v450
	*(*int32)(unsafe.Add(mBase, uint32(v37)+240)) = v446
	*(*int32)(unsafe.Add(mBase, uint32(v37)+244)) = v442
	*(*int32)(unsafe.Add(mBase, uint32(v37)+248)) = v439
	*(*int32)(unsafe.Add(mBase, uint32(v37)+252)) = v437
	F_ReThrowError(m, v633)
	mBase = m.M
	v893 = m.ExcPending
	if v893 != 0 {
		goto L6
	} else {
		goto L108
	}
L88:
	;
	v710 = *(*int32)(unsafe.Add(mBase, uint32(v707)+4))
	if v710 <= int32(0) {
		goto L87
	} else {
		goto L89
	}
L89:
	;
	v713 = int32(0)
	if v713 < v710 {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v716 = v710
	goto L92
L91:
	;
	v716 = v713
	goto L92
L92:
	;
	v717 = *(*int32)(unsafe.Add(mBase, uint32(v707)+12))
	v740 = int32(0)
	goto L93
L93:
	;
	v756 = *(*int32)(unsafe.Add(mBase, uint32(v717+v740<<(uint(int32(2))%32))))
	v757 = *(*int32)(unsafe.Add(mBase, uint32(v756)+4))
	if v757 != 0 {
		goto L95
	} else {
		goto L96
	}
L94:
	;
	goto L87
L95:
	;
	v758 = *(*int32)(unsafe.Add(mBase, uint32(v633)+28))
	v769 = v757
	goto L98
L96:
	;
	goto L97
L97:
	;
	v842 = v740 + int32(1)
	if v842 != v716 {
		v740 = v842
		goto L93
	} else {
		goto L107
	}
L98:
	;
	v799 = *(*int32)(unsafe.Add(mBase, uint32(v769)))
	if v799 == int32(-1) {
		goto L101
	} else {
		goto L102
	}
L99:
	;
	goto L97
L100:
	;
	v806 = *(*int32)(unsafe.Add(mBase, uint32(v769)+8))
	if v806 != 0 {
		v769 = v806
		goto L98
	} else {
		goto L106
	}
L101:
	;
	if base.B2i32(v758 == int32(67108896))|base.B2i32(v758 == int32(67371461)) != 0 {
		goto L100
	} else {
		goto L104
	}
L102:
	;
	goto L103
L103:
	;
	if base.B2i32(v758 == v799)|base.B2i32(v799 == v758&int32(4095)) != 0 {
		goto L9
	} else {
		goto L105
	}
L104:
	;
	goto L9
L105:
	;
	goto L100
L106:
	;
	goto L99
L107:
	;
	goto L94
L108:
	;
	goto L3
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+200)) = v910
	v1140 = v40
	v1142 = v42
	v1143 = v56
	v1144 = v57
	v1145 = v58
	v1146 = v59
	v1147 = v60
	v1148 = v61
	v1149 = v62
	v1150 = v63
	v1151 = v64
	v1152 = v65
	v1153 = v66
	goto L7
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+208)) = v443
	*(*int32)(unsafe.Add(mBase, uint32(v37)+204)) = v444
	*(*int32)(unsafe.Add(mBase, uint32(v37)+212)) = v440
	*(*int32)(unsafe.Add(mBase, uint32(v37)+216)) = v447
	*(*int32)(unsafe.Add(mBase, uint32(v37)+220)) = v449
	*(*int32)(unsafe.Add(mBase, uint32(v37)+224)) = v445
	*(*int32)(unsafe.Add(mBase, uint32(v37)+228)) = v448
	*(*int32)(unsafe.Add(mBase, uint32(v37)+232)) = v441
	*(*int32)(unsafe.Add(mBase, uint32(v37)+236)) = v450
	*(*int32)(unsafe.Add(mBase, uint32(v37)+240)) = v446
	*(*int32)(unsafe.Add(mBase, uint32(v37)+244)) = v442
	*(*int32)(unsafe.Add(mBase, uint32(v37)+248)) = v439
	*(*int32)(unsafe.Add(mBase, uint32(v37)+252)) = v437
	v992 = F_cstring_to_text(m, v937)
	mBase = m.M
	v993 = m.ExcPending
	if v993 != 0 {
		goto L6
	} else {
		goto L111
	}
L111:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+208)) = v443
	*(*int32)(unsafe.Add(mBase, uint32(v37)+204)) = v444
	*(*int32)(unsafe.Add(mBase, uint32(v37)+212)) = v440
	*(*int32)(unsafe.Add(mBase, uint32(v37)+216)) = v447
	*(*int32)(unsafe.Add(mBase, uint32(v37)+220)) = v449
	*(*int32)(unsafe.Add(mBase, uint32(v37)+224)) = v445
	*(*int32)(unsafe.Add(mBase, uint32(v37)+228)) = v448
	*(*int32)(unsafe.Add(mBase, uint32(v37)+232)) = v441
	*(*int32)(unsafe.Add(mBase, uint32(v37)+236)) = v450
	*(*int32)(unsafe.Add(mBase, uint32(v37)+240)) = v446
	*(*int32)(unsafe.Add(mBase, uint32(v37)+244)) = v442
	*(*int32)(unsafe.Add(mBase, uint32(v37)+248)) = v439
	*(*int32)(unsafe.Add(mBase, uint32(v37)+252)) = v437
	F_assign_simple_var(m, l0, v923, base.I64_extend_i32_u(v992), int32(0), int32(1))
	mBase = m.M
	v1011 = m.ExcPending
	if v1011 != 0 {
		goto L6
	} else {
		goto L112
	}
L112:
	;
	v1012 = *(*int32)(unsafe.Add(mBase, uint32(v633)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v37)+208)) = v443
	*(*int32)(unsafe.Add(mBase, uint32(v37)+204)) = v444
	*(*int32)(unsafe.Add(mBase, uint32(v37)+212)) = v440
	*(*int32)(unsafe.Add(mBase, uint32(v37)+216)) = v447
	*(*int32)(unsafe.Add(mBase, uint32(v37)+220)) = v449
	*(*int32)(unsafe.Add(mBase, uint32(v37)+224)) = v445
	*(*int32)(unsafe.Add(mBase, uint32(v37)+228)) = v448
	*(*int32)(unsafe.Add(mBase, uint32(v37)+232)) = v441
	*(*int32)(unsafe.Add(mBase, uint32(v37)+236)) = v450
	*(*int32)(unsafe.Add(mBase, uint32(v37)+240)) = v446
	*(*int32)(unsafe.Add(mBase, uint32(v37)+244)) = v442
	*(*int32)(unsafe.Add(mBase, uint32(v37)+248)) = v439
	*(*int32)(unsafe.Add(mBase, uint32(v37)+252)) = v437
	v1026 = F_cstring_to_text(m, v1012)
	mBase = m.M
	v1027 = m.ExcPending
	if v1027 != 0 {
		goto L6
	} else {
		goto L113
	}
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+208)) = v443
	*(*int32)(unsafe.Add(mBase, uint32(v37)+204)) = v444
	*(*int32)(unsafe.Add(mBase, uint32(v37)+212)) = v440
	*(*int32)(unsafe.Add(mBase, uint32(v37)+216)) = v447
	*(*int32)(unsafe.Add(mBase, uint32(v37)+220)) = v449
	*(*int32)(unsafe.Add(mBase, uint32(v37)+224)) = v445
	*(*int32)(unsafe.Add(mBase, uint32(v37)+228)) = v448
	*(*int32)(unsafe.Add(mBase, uint32(v37)+232)) = v441
	*(*int32)(unsafe.Add(mBase, uint32(v37)+236)) = v450
	*(*int32)(unsafe.Add(mBase, uint32(v37)+240)) = v446
	*(*int32)(unsafe.Add(mBase, uint32(v37)+244)) = v442
	*(*int32)(unsafe.Add(mBase, uint32(v37)+248)) = v439
	*(*int32)(unsafe.Add(mBase, uint32(v37)+252)) = v437
	F_assign_simple_var(m, l0, v918, base.I64_extend_i32_u(v1026), int32(0), int32(1))
	mBase = m.M
	v1045 = m.ExcPending
	if v1045 != 0 {
		goto L6
	} else {
		goto L114
	}
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v448))) = v633
	*(*int32)(unsafe.Add(mBase, uint32(v437))) = int32(0)
	v1049 = *(*int32)(unsafe.Add(mBase, uint32(v756)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v37)+208)) = v443
	*(*int32)(unsafe.Add(mBase, uint32(v37)+204)) = v444
	*(*int32)(unsafe.Add(mBase, uint32(v37)+212)) = v440
	*(*int32)(unsafe.Add(mBase, uint32(v37)+216)) = v447
	*(*int32)(unsafe.Add(mBase, uint32(v37)+220)) = v449
	*(*int32)(unsafe.Add(mBase, uint32(v37)+224)) = v445
	*(*int32)(unsafe.Add(mBase, uint32(v37)+228)) = v448
	*(*int32)(unsafe.Add(mBase, uint32(v37)+232)) = v441
	*(*int32)(unsafe.Add(mBase, uint32(v37)+236)) = v450
	*(*int32)(unsafe.Add(mBase, uint32(v37)+240)) = v446
	*(*int32)(unsafe.Add(mBase, uint32(v37)+244)) = v442
	*(*int32)(unsafe.Add(mBase, uint32(v37)+248)) = v439
	*(*int32)(unsafe.Add(mBase, uint32(v37)+252)) = v437
	v1063 = F_exec_stmts(m, l0, v1049)
	mBase = m.M
	v1064 = m.ExcPending
	if v1064 != 0 {
		goto L6
	} else {
		goto L115
	}
L115:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+200)) = v1063
	*(*int32)(unsafe.Add(mBase, uint32(v448))) = v445
	v1067 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
	*(*int32)(unsafe.Add(mBase, uint32(v449))) = v1067
	*(*int32)(unsafe.Add(mBase, uint32(v37)+208)) = v443
	*(*int32)(unsafe.Add(mBase, uint32(v37)+204)) = v444
	*(*int32)(unsafe.Add(mBase, uint32(v37)+212)) = v440
	*(*int32)(unsafe.Add(mBase, uint32(v37)+216)) = v447
	*(*int32)(unsafe.Add(mBase, uint32(v37)+220)) = v449
	*(*int32)(unsafe.Add(mBase, uint32(v37)+224)) = v445
	*(*int32)(unsafe.Add(mBase, uint32(v37)+228)) = v448
	*(*int32)(unsafe.Add(mBase, uint32(v37)+232)) = v441
	*(*int32)(unsafe.Add(mBase, uint32(v37)+236)) = v450
	*(*int32)(unsafe.Add(mBase, uint32(v37)+240)) = v446
	*(*int32)(unsafe.Add(mBase, uint32(v37)+244)) = v442
	*(*int32)(unsafe.Add(mBase, uint32(v37)+248)) = v439
	*(*int32)(unsafe.Add(mBase, uint32(v37)+252)) = v437
	v1082 = *(*int32)(unsafe.Add(mBase, uint32(v1067)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+108)) = v1082
	*(*int32)(unsafe.Add(mBase, uint32(v37)+208)) = v443
	*(*int32)(unsafe.Add(mBase, uint32(v37)+204)) = v444
	*(*int32)(unsafe.Add(mBase, uint32(v37)+212)) = v440
	*(*int32)(unsafe.Add(mBase, uint32(v37)+216)) = v447
	*(*int32)(unsafe.Add(mBase, uint32(v37)+220)) = v449
	*(*int32)(unsafe.Add(mBase, uint32(v37)+224)) = v445
	*(*int32)(unsafe.Add(mBase, uint32(v37)+228)) = v448
	*(*int32)(unsafe.Add(mBase, uint32(v37)+232)) = v441
	*(*int32)(unsafe.Add(mBase, uint32(v37)+236)) = v450
	*(*int32)(unsafe.Add(mBase, uint32(v37)+240)) = v446
	*(*int32)(unsafe.Add(mBase, uint32(v37)+244)) = v442
	*(*int32)(unsafe.Add(mBase, uint32(v37)+248)) = v439
	*(*int32)(unsafe.Add(mBase, uint32(v37)+252)) = v437
	F_MemoryContextReset(m, v440)
	mBase = m.M
	v1098 = m.ExcPending
	if v1098 != 0 {
		goto L6
	} else {
		goto L116
	}
L116:
	;
	goto L8
L117:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+208)) = v1146
	*(*int32)(unsafe.Add(mBase, uint32(v37)+204)) = v1147
	*(*int32)(unsafe.Add(mBase, uint32(v37)+212)) = v1143
	*(*int32)(unsafe.Add(mBase, uint32(v37)+216)) = v1150
	*(*int32)(unsafe.Add(mBase, uint32(v37)+220)) = v1152
	*(*int32)(unsafe.Add(mBase, uint32(v37)+224)) = v1148
	*(*int32)(unsafe.Add(mBase, uint32(v37)+228)) = v1151
	*(*int32)(unsafe.Add(mBase, uint32(v37)+232)) = v1144
	*(*int32)(unsafe.Add(mBase, uint32(v37)+236)) = v1153
	*(*int32)(unsafe.Add(mBase, uint32(v37)+240)) = v1149
	*(*int32)(unsafe.Add(mBase, uint32(v37)+244)) = v1145
	*(*int32)(unsafe.Add(mBase, uint32(v37)+248)) = v1142
	*(*int32)(unsafe.Add(mBase, uint32(v37)+252)) = v1140
	F_errstart_cold(m, int32(21), int32(_a_F_exec_stmt_block_2))
	mBase = m.M
	v1254 = m.ExcPending
	if v1254 != 0 {
		goto L6
	} else {
		goto L133
	}
L118:
	;
	m.G0 = v37 + int32(256)
	return v1231
L119:
	;
	if base.B2i32(v1173 == int32(1))|base.B2i32(base.Ui32(int32(3)) < base.Ui32(v1173)) != 0 {
		goto L117
	} else {
		goto L122
	}
L120:
	;
	goto L121
L121:
	;
	v1182 = int32(1)
	v1183 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v1183 == int32(0) {
		v1231 = v1182
		goto L118
	} else {
		goto L123
	}
L122:
	;
	v1181 = *(*int32)(unsafe.Add(mBase, uint32(v37)+200))
	v1231 = v1181
	goto L118
L123:
	;
	v1186 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v1186 == int32(0) {
		v1231 = v1182
		goto L118
	} else {
		goto L124
	}
L124:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+208)) = v1146
	*(*int32)(unsafe.Add(mBase, uint32(v37)+204)) = v1147
	*(*int32)(unsafe.Add(mBase, uint32(v37)+212)) = v1143
	*(*int32)(unsafe.Add(mBase, uint32(v37)+216)) = v1150
	*(*int32)(unsafe.Add(mBase, uint32(v37)+220)) = v1152
	*(*int32)(unsafe.Add(mBase, uint32(v37)+224)) = v1148
	*(*int32)(unsafe.Add(mBase, uint32(v37)+228)) = v1151
	*(*int32)(unsafe.Add(mBase, uint32(v37)+232)) = v1144
	*(*int32)(unsafe.Add(mBase, uint32(v37)+236)) = v1153
	*(*int32)(unsafe.Add(mBase, uint32(v37)+240)) = v1149
	*(*int32)(unsafe.Add(mBase, uint32(v37)+244)) = v1145
	*(*int32)(unsafe.Add(mBase, uint32(v37)+248)) = v1142
	*(*int32)(unsafe.Add(mBase, uint32(v37)+252)) = v1140
	v1204 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1186))))
	v1207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1183))))
	if base.B2i32(v1204 == int32(0))|base.B2i32(v1204 != v1207) != 0 {
		v1225 = v1204
		v1226 = v1207
		goto L126
	} else {
		goto L127
	}
L125:
	;
	if v1225-v1226 != 0 {
		v1231 = v1182
		goto L118
	} else {
		goto L132
	}
L126:
	;
	goto L125
L127:
	;
	v1210 = v1186
	v1211 = v1183
	goto L128
L128:
	;
	v1214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1211)+1)))
	v1215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1210)+1)))
	if v1215 == int32(0) {
		v1225 = v1215
		v1226 = v1214
		goto L126
	} else {
		goto L130
	}
L129:
	;
	v1225 = v1215
	v1226 = v1214
	goto L126
L130:
	;
	v1218 = int32(1)
	if v1215 == v1214 {
		v1210 = v1210 + v1218
		v1211 = v1211 + v1218
		goto L128
	} else {
		goto L131
	}
L131:
	;
	goto L129
L132:
	;
	v1228 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v1228
	v1231 = v1228
	goto L118
L133:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+204)) = v1147
	*(*int32)(unsafe.Add(mBase, uint32(v37)+208)) = v1146
	*(*int32)(unsafe.Add(mBase, uint32(v37)+212)) = v1143
	*(*int32)(unsafe.Add(mBase, uint32(v37)+216)) = v1150
	*(*int32)(unsafe.Add(mBase, uint32(v37)+220)) = v1152
	*(*int32)(unsafe.Add(mBase, uint32(v37)+224)) = v1148
	*(*int32)(unsafe.Add(mBase, uint32(v37)+228)) = v1151
	*(*int32)(unsafe.Add(mBase, uint32(v37)+232)) = v1144
	*(*int32)(unsafe.Add(mBase, uint32(v37)+236)) = v1153
	*(*int32)(unsafe.Add(mBase, uint32(v37)+240)) = v1149
	*(*int32)(unsafe.Add(mBase, uint32(v37)+244)) = v1145
	*(*int32)(unsafe.Add(mBase, uint32(v37)+248)) = v1142
	*(*int32)(unsafe.Add(mBase, uint32(v37)+252)) = v1140
	v1268 = *(*int32)(unsafe.Add(mBase, uint32(v37)+200))
	*(*int32)(unsafe.Add(mBase, uint32(v37))) = v1268
	F_errmsg_internal(m, int32(_a_F_exec_stmt_block_18), v37)
	mBase = m.M
	v1272 = m.ExcPending
	if v1272 != 0 {
		goto L6
	} else {
		goto L134
	}
L134:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+208)) = v1146
	*(*int32)(unsafe.Add(mBase, uint32(v37)+204)) = v1147
	*(*int32)(unsafe.Add(mBase, uint32(v37)+212)) = v1143
	*(*int32)(unsafe.Add(mBase, uint32(v37)+216)) = v1150
	*(*int32)(unsafe.Add(mBase, uint32(v37)+220)) = v1152
	*(*int32)(unsafe.Add(mBase, uint32(v37)+224)) = v1148
	*(*int32)(unsafe.Add(mBase, uint32(v37)+228)) = v1151
	*(*int32)(unsafe.Add(mBase, uint32(v37)+232)) = v1144
	*(*int32)(unsafe.Add(mBase, uint32(v37)+236)) = v1153
	*(*int32)(unsafe.Add(mBase, uint32(v37)+240)) = v1149
	*(*int32)(unsafe.Add(mBase, uint32(v37)+244)) = v1145
	*(*int32)(unsafe.Add(mBase, uint32(v37)+248)) = v1142
	*(*int32)(unsafe.Add(mBase, uint32(v37)+252)) = v1140
	F_errfinish(m, int32(_a_F_exec_stmt_block_5), int32(2014), int32(_a_F_exec_stmt_block_6))
	mBase = m.M
	v1290 = m.ExcPending
	if v1290 != 0 {
		goto L6
	} else {
		goto L135
	}
L135:
	;
	goto L5
L136:
	;
	v1330 = int32(v1326)
	m.G0 = v37
	v1332 = *(*int32)(unsafe.Add(mBase, uint32(v1330)+4))
	v1333 = *(*int32)(unsafe.Add(mBase, uint32(v1330)))
	v1336 = *(*int32)(unsafe.Add(mBase, uint32(v1333)))
	if v37+int32(24) == v1336 {
		goto L139
	} else {
		goto L140
	}
L137:
	;
	m.ExcPending = 1
	goto L145
L138:
	;
	if v1340 != 0 {
		goto L142
	} else {
		goto L143
	}
L139:
	;
	v1338 = *(*int32)(unsafe.Add(mBase, uint32(v1333)+4))
	v1340 = v1338
	goto L141
L140:
	;
	v1340 = int32(0)
	goto L141
L141:
	;
	goto L138
L142:
	;
	v1341 = *(*int32)(unsafe.Add(mBase, uint32(v37)+252))
	v1342 = *(*int32)(unsafe.Add(mBase, uint32(v37)+248))
	v1343 = *(*int32)(unsafe.Add(mBase, uint32(v37)+244))
	v1344 = *(*int32)(unsafe.Add(mBase, uint32(v37)+240))
	v1345 = *(*int32)(unsafe.Add(mBase, uint32(v37)+236))
	v1346 = *(*int32)(unsafe.Add(mBase, uint32(v37)+232))
	v1347 = *(*int32)(unsafe.Add(mBase, uint32(v37)+228))
	v1348 = *(*int32)(unsafe.Add(mBase, uint32(v37)+224))
	v1349 = *(*int32)(unsafe.Add(mBase, uint32(v37)+220))
	v1350 = *(*int32)(unsafe.Add(mBase, uint32(v37)+216))
	v1351 = *(*int32)(unsafe.Add(mBase, uint32(v37)+212))
	v1352 = *(*int32)(unsafe.Add(mBase, uint32(v37)+208))
	v1353 = *(*int32)(unsafe.Add(mBase, uint32(v37)+204))
	v53 = v1341
	v54 = v1332
	v55 = v1342
	v56 = v1351
	v57 = v1346
	v58 = v1343
	v59 = v1352
	v60 = v1353
	v61 = v1348
	v62 = v1344
	v63 = v1350
	v64 = v1347
	v65 = v1349
	v66 = v1345
	v67 = v1340
	goto L1
L143:
	;
	goto L144
L144:
	;
	F___wasm_longjmp(m, v1333, v1332)
	mBase = m.M
	v1357 = m.ExcPending
	if v1357 != 0 {
		goto L145
	} else {
		goto L146
	}
L145:
	;
	return int32(0)
L146:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
