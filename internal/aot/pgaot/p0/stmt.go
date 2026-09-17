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
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
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
	var v88 int32
	_ = v88
	var v109 int32
	_ = v109
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v249 int32
	_ = v249
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v286 int32
	_ = v286
	var v304 int32
	_ = v304
	var v319 int32
	_ = v319
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v364 int32
	_ = v364
	var v366 int32
	_ = v366
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v411 int32
	_ = v411
	var v415 int32
	_ = v415
	var v417 int32
	_ = v417
	var v419 int32
	_ = v419
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
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
	var v479 int32
	_ = v479
	var v482 int32
	_ = v482
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v501 int32
	_ = v501
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v559 int32
	_ = v559
	var v561 int32
	_ = v561
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v577 int32
	_ = v577
	var v579 int32
	_ = v579
	var v583 int32
	_ = v583
	var v588 int32
	_ = v588
	var v591 int32
	_ = v591
	var v594 int32
	_ = v594
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v640 int32
	_ = v640
	var v655 int32
	_ = v655
	var v677 int32
	_ = v677
	var v681 int32
	_ = v681
	var v696 int32
	_ = v696
	var v697 int32
	_ = v697
	var v698 int32
	_ = v698
	var v701 int32
	_ = v701
	var v704 int32
	_ = v704
	var v707 int32
	_ = v707
	var v708 int32
	_ = v708
	var v731 int32
	_ = v731
	var v746 int32
	_ = v746
	var v747 int32
	_ = v747
	var v748 int32
	_ = v748
	var v759 int32
	_ = v759
	var v788 int32
	_ = v788
	var v795 int32
	_ = v795
	var v830 int32
	_ = v830
	var v880 int32
	_ = v880
	var v883 int32
	_ = v883
	var v897 int32
	_ = v897
	var v898 int32
	_ = v898
	var v900 int32
	_ = v900
	var v901 int32
	_ = v901
	var v902 int32
	_ = v902
	var v905 int32
	_ = v905
	var v906 int32
	_ = v906
	var v910 int32
	_ = v910
	var v924 int32
	_ = v924
	var v925 int32
	_ = v925
	var v927 int32
	_ = v927
	var v928 int32
	_ = v928
	var v936 int32
	_ = v936
	var v944 int32
	_ = v944
	var v952 int32
	_ = v952
	var v960 int32
	_ = v960
	var v963 int32
	_ = v963
	var v979 int32
	_ = v979
	var v980 int32
	_ = v980
	var v997 int32
	_ = v997
	var v998 int32
	_ = v998
	var v1012 int32
	_ = v1012
	var v1013 int32
	_ = v1013
	var v1030 int32
	_ = v1030
	var v1034 int32
	_ = v1034
	var v1048 int32
	_ = v1048
	var v1049 int32
	_ = v1049
	var v1052 int32
	_ = v1052
	var v1067 int32
	_ = v1067
	var v1083 int32
	_ = v1083
	var v1124 int32
	_ = v1124
	var v1126 int32
	_ = v1126
	var v1127 int32
	_ = v1127
	var v1128 int32
	_ = v1128
	var v1129 int32
	_ = v1129
	var v1130 int32
	_ = v1130
	var v1131 int32
	_ = v1131
	var v1132 int32
	_ = v1132
	var v1133 int32
	_ = v1133
	var v1134 int32
	_ = v1134
	var v1135 int32
	_ = v1135
	var v1136 int32
	_ = v1136
	var v1137 int32
	_ = v1137
	var v1156 int32
	_ = v1156
	var v1164 int32
	_ = v1164
	var v1165 int32
	_ = v1165
	var v1166 int32
	_ = v1166
	var v1169 int32
	_ = v1169
	var v1187 int32
	_ = v1187
	var v1190 int32
	_ = v1190
	var v1193 int32
	_ = v1193
	var v1194 int32
	_ = v1194
	var v1197 int32
	_ = v1197
	var v1198 int32
	_ = v1198
	var v1201 int32
	_ = v1201
	var v1208 int32
	_ = v1208
	var v1209 int32
	_ = v1209
	var v1211 int32
	_ = v1211
	var v1214 int32
	_ = v1214
	var v1237 int32
	_ = v1237
	var v1251 int32
	_ = v1251
	var v1255 int32
	_ = v1255
	var v1273 int32
	_ = v1273
	var v1307 int32
	_ = v1307
	var v1308 int64
	_ = v1308
	var v1312 int32
	_ = v1312
	var v1314 int32
	_ = v1314
	var v1315 int32
	_ = v1315
	var v1318 int32
	_ = v1318
	var v1320 int32
	_ = v1320
	var v1322 int32
	_ = v1322
	var v1323 int32
	_ = v1323
	var v1324 int32
	_ = v1324
	var v1325 int32
	_ = v1325
	var v1326 int32
	_ = v1326
	var v1327 int32
	_ = v1327
	var v1328 int32
	_ = v1328
	var v1329 int32
	_ = v1329
	var v1330 int32
	_ = v1330
	var v1331 int32
	_ = v1331
	var v1332 int32
	_ = v1332
	var v1333 int32
	_ = v1333
	var v1334 int32
	_ = v1334
	var v1335 int32
	_ = v1335
	var v1339 int32
	_ = v1339
	v3 = int32(0)
	v34 = m.G0
	v36 = v34 - int32(256)
	m.G0 = v36
	v39 = l0 + int32(132)
	v41 = l1 + int32(28)
	v43 = l0 + int32(96)
	v45 = l0 + int32(36)
	v47 = l0 + int32(120)
	v52 = v3
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
	v66 = int32(-1)
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
	if v66 != int32(1) {
		goto L11
	} else {
		goto L12
	}
L5:
	;
	goto L3
L6:
	;
	v1307 = int32(m.ExcTag)
	v1308 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v1307 == int32(0) {
		goto L136
	} else {
		goto L137
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1124))) = int32(0)
	v1156 = *(*int32)(unsafe.Add(mBase, uint32(v36)+200))
	if v1156 != int32(1) {
		goto L119
	} else {
		goto L120
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_exec_stmt_block[0])) = v436
	*(*int32)(unsafe.Add(mBase, _c_F_exec_stmt_block[1])) = v435
	v1124 = v429
	v1126 = v431
	v1127 = v432
	v1128 = v433
	v1129 = v434
	v1130 = v435
	v1131 = v436
	v1132 = v437
	v1133 = v438
	v1134 = v439
	v1135 = v440
	v1136 = v441
	v1137 = v442
	goto L7
L9:
	;
	v900 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v901 = *(*int32)(unsafe.Add(mBase, uint32(v697)+4))
	v902 = int32(2)
	v905 = *(*int32)(unsafe.Add(mBase, uint32(v900+v901<<(uint(v902)%32))))
	v906 = *(*int32)(unsafe.Add(mBase, uint32(v697)))
	v910 = *(*int32)(unsafe.Add(mBase, uint32(v900+v906<<(uint(v902)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v36)+208)) = v435
	*(*int32)(unsafe.Add(mBase, uint32(v36)+204)) = v436
	*(*int32)(unsafe.Add(mBase, uint32(v36)+212)) = v432
	*(*int32)(unsafe.Add(mBase, uint32(v36)+216)) = v439
	*(*int32)(unsafe.Add(mBase, uint32(v36)+220)) = v441
	*(*int32)(unsafe.Add(mBase, uint32(v36)+224)) = v437
	*(*int32)(unsafe.Add(mBase, uint32(v36)+228)) = v440
	*(*int32)(unsafe.Add(mBase, uint32(v36)+232)) = v433
	*(*int32)(unsafe.Add(mBase, uint32(v36)+236)) = v442
	*(*int32)(unsafe.Add(mBase, uint32(v36)+240)) = v438
	*(*int32)(unsafe.Add(mBase, uint32(v36)+244)) = v434
	*(*int32)(unsafe.Add(mBase, uint32(v36)+248)) = v431
	*(*int32)(unsafe.Add(mBase, uint32(v36)+252)) = v429
	v924 = int32(_a_F_exec_stmt_block_0)
	v925 = int32(63)
	v927 = int32(48)
	v928 = v748&v925 + v927
	*(*uint8)(unsafe.Add(mBase, _c_F_exec_stmt_block[2])) = uint8(v928)
	v936 = int32(base.Ui32(v748)>>(uint(int32(24))%32))&v925 + v927
	*(*uint8)(unsafe.Add(mBase, _c_F_exec_stmt_block[3])) = uint8(v936)
	v944 = int32(base.Ui32(v748)>>(uint(int32(18))%32))&v925 + v927
	*(*uint8)(unsafe.Add(mBase, _c_F_exec_stmt_block[4])) = uint8(v944)
	v952 = int32(base.Ui32(v748)>>(uint(int32(12))%32))&v925 + v927
	*(*uint8)(unsafe.Add(mBase, _c_F_exec_stmt_block[5])) = uint8(v952)
	v960 = int32(base.Ui32(v748)>>(uint(int32(6))%32))&v925 + v927
	*(*uint8)(unsafe.Add(mBase, _c_F_exec_stmt_block[6])) = uint8(v960)
	v963 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_exec_stmt_block[7])) = uint8(v963)
	goto L110
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+132)) = int32(0)
	v883 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v36)+208)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v36)+204)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v36)+212)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v36)+216)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v36)+220)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v36)+224)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v36)+228)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v36)+232)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v36)+236)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v36)+240)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v36)+244)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v36)+248)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v36)+252)) = v39
	v897 = F_exec_stmts(m, l0, v883)
	mBase = m.M
	v898 = m.ExcPending
	if v898 != 0 {
		goto L6
	} else {
		goto L109
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+200)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+132)) = int32(_a_F_exec_stmt_block_1)
	v88 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if int32(0) < v88 {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	v429 = v52
	v430 = v53
	v431 = v54
	v432 = v55
	v433 = v56
	v434 = v57
	v435 = v58
	v436 = v59
	v437 = v60
	v438 = v61
	v439 = v62
	v440 = v63
	v441 = v64
	v442 = v65
	goto L13
L13:
	;
	if v430 == int32(0) {
		goto L57
	} else {
		goto L58
	}
L14:
	;
	v109 = int32(0)
	goto L17
L15:
	;
	goto L16
L16:
	;
	v358 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+128)) = v358
	v360 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	if v360 == v358 {
		goto L10
	} else {
		goto L46
	}
L17:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v126 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v127 = int32(2)
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v126+v109<<(uint(v127)%32))))
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v125+v130<<(uint(v127)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+128)) = v134
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v134)))
	switch v136 {
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
	v322 = v109 + int32(1)
	v323 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v322 < v323 {
		v109 = v322
		goto L17
	} else {
		goto L45
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+208)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v36)+204)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v36)+212)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v36)+216)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v36)+220)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v36)+224)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v36)+228)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v36)+232)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v36)+236)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v36)+240)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v36)+244)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v36)+248)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v36)+252)) = v39
	F_exec_assign_expr(m, l0, v134, v191)
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L6
	} else {
		goto L44
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+208)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v36)+204)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v36)+212)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v36)+216)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v36)+220)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v36)+224)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v36)+228)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v36)+232)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v36)+236)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v36)+240)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v36)+244)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v36)+248)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v36)+252)) = v39
	F_errstart_cold(m, int32(21), int32(_a_F_exec_stmt_block_2))
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L6
	} else {
		goto L41
	}
L22:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v134)+20))
	if v215 == int32(0) {
		goto L36
	} else {
		goto L37
	}
L23:
	;
	v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v134)+45)))
	if v137 != int32(1) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v185 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v134)+48)) = v185
	v187 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v134)+44)) = uint16(v187)
	*(*int32)(unsafe.Add(mBase, uint32(v134)+40)) = v185
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v134)+20))
	if v191 != 0 {
		goto L20
	} else {
		goto L33
	}
L25:
	;
	v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v134)+44)))
	if v140 != 0 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v134)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v36)+208)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v36)+204)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v36)+212)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v36)+216)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v36)+220)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v36)+224)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v36)+228)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v36)+232)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v36)+236)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v36)+240)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v36)+244)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v36)+248)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v36)+252)) = v39
	F_pfree(m, v168)
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L6
	} else {
		goto L32
	}
L27:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v134)+24))
	v142 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v141)+12)))
	if v142 != int32(_a_F_exec_stmt_block_3) {
		goto L26
	} else {
		goto L28
	}
L28:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v134)+40))
	v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v145))))
	if v146 != int32(1) {
		goto L26
	} else {
		goto L29
	}
L29:
	;
	v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v145)+1)))
	if v149 != int32(3) {
		goto L26
	} else {
		goto L30
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+208)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v36)+204)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v36)+212)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v36)+216)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v36)+220)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v36)+224)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v36)+228)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v36)+232)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v36)+236)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v36)+240)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v36)+244)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v36)+248)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v36)+252)) = v39
	F_DeleteExpandedObject(m, v145)
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
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
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v134)+24))
	v193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192)+15)))
	if v193 != int32(100) {
		goto L19
	} else {
		goto L34
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+208)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v36)+204)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v36)+212)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v36)+216)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v36)+220)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v36)+224)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v36)+228)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v36)+232)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v36)+236)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v36)+240)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v36)+244)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v36)+248)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v36)+252)) = v39
	F_exec_assign_value(m, l0, v134, int32(0), int32(1), int32(705), int32(-1))
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L6
	} else {
		goto L35
	}
L35:
	;
	goto L19
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+208)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v36)+204)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v36)+212)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v36)+216)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v36)+220)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v36)+224)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v36)+228)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v36)+232)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v36)+236)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v36)+240)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v36)+244)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v36)+248)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v36)+252)) = v39
	v231 = int32(0)
	F_exec_move_row(m, l0, v134, v231, v231)
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L6
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+208)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v36)+204)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v36)+212)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v36)+216)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v36)+220)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v36)+224)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v36)+228)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v36)+232)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v36)+236)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v36)+240)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v36)+244)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v36)+248)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v36)+252)) = v39
	F_exec_assign_expr(m, l0, v134, v215)
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
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
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v134)))
	*(*int32)(unsafe.Add(mBase, uint32(v36)+208)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v36)+204)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v36)+212)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v36)+216)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v36)+220)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v36)+224)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v36)+228)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v36)+232)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v36)+236)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v36)+240)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v36)+244)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v36)+248)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v36)+252)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v36)+16)) = v267
	F_errmsg_internal(m, int32(_a_F_exec_stmt_block_4), v36+int32(16))
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L6
	} else {
		goto L42
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+208)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v36)+204)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v36)+212)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v36)+216)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v36)+220)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v36)+224)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v36)+228)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v36)+232)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v36)+236)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v36)+240)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v36)+244)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v36)+248)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v36)+252)) = v39
	F_errfinish(m, int32(_a_F_exec_stmt_block_5), int32(1759), int32(_a_F_exec_stmt_block_6))
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
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
	v364 = *(*int32)(unsafe.Add(mBase, _c_F_exec_stmt_block[8]))
	v366 = *(*int32)(unsafe.Add(mBase, _c_F_exec_stmt_block[9]))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+132)) = int32(_a_F_exec_stmt_block_7)
	v369 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v370 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v371 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	if v371 != 0 {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+208)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v36)+204)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v36)+212)) = v395
	*(*int32)(unsafe.Add(mBase, uint32(v36)+216)) = v394
	*(*int32)(unsafe.Add(mBase, uint32(v36)+220)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v36)+224)) = v369
	*(*int32)(unsafe.Add(mBase, uint32(v36)+228)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v36)+232)) = v370
	*(*int32)(unsafe.Add(mBase, uint32(v36)+236)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v36)+240)) = v364
	*(*int32)(unsafe.Add(mBase, uint32(v36)+244)) = v366
	*(*int32)(unsafe.Add(mBase, uint32(v36)+248)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v36)+252)) = v39
	F_BeginInternalSubTransaction(m, int32(0))
	mBase = m.M
	v411 = m.ExcPending
	if v411 != 0 {
		goto L6
	} else {
		goto L52
	}
L48:
	;
	v394 = v62
	v395 = v371
	goto L47
L49:
	;
	goto L50
L50:
	;
	v372 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	*(*int32)(unsafe.Add(mBase, uint32(v36)+208)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v36)+204)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v36)+212)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v36)+216)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v36)+220)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v36)+224)) = v369
	*(*int32)(unsafe.Add(mBase, uint32(v36)+228)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v36)+232)) = v370
	*(*int32)(unsafe.Add(mBase, uint32(v36)+236)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v36)+240)) = v364
	*(*int32)(unsafe.Add(mBase, uint32(v36)+244)) = v366
	*(*int32)(unsafe.Add(mBase, uint32(v36)+248)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v36)+252)) = v39
	v390 = F_AllocSetContextCreateInternal(m, v372, int32(_a_F_exec_stmt_block_8), int32(0), int32(_a_F_exec_stmt_block_9), int32(_a_F_exec_stmt_block_10))
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L6
	} else {
		goto L51
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43))) = v390
	v394 = v390
	v395 = v390
	goto L47
L52:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_exec_stmt_block[9])) = v366
	v415 = *(*int32)(unsafe.Add(mBase, _c_F_exec_stmt_block[1]))
	v417 = *(*int32)(unsafe.Add(mBase, _c_F_exec_stmt_block[0]))
	goto L53
L53:
	;
	v419 = v36 + int32(32)
	*(*int32)(unsafe.Add(mBase, uint32(v419)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v419))) = v36 + int32(24)
	goto L56
L54:
	;
	v429 = v39
	v430 = int32(0)
	v431 = v41
	v432 = v395
	v433 = v370
	v434 = v366
	v435 = v415
	v436 = v417
	v437 = v369
	v438 = v364
	v439 = v394
	v440 = v45
	v441 = v43
	v442 = v47
	goto L13
L56:
	;
	goto L54
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+204)) = v436
	*(*int32)(unsafe.Add(mBase, _c_F_exec_stmt_block[1])) = v36 + int32(32)
	*(*int32)(unsafe.Add(mBase, uint32(v36)+208)) = v435
	*(*int32)(unsafe.Add(mBase, uint32(v36)+212)) = v432
	*(*int32)(unsafe.Add(mBase, uint32(v36)+216)) = v439
	*(*int32)(unsafe.Add(mBase, uint32(v36)+220)) = v441
	*(*int32)(unsafe.Add(mBase, uint32(v36)+224)) = v437
	*(*int32)(unsafe.Add(mBase, uint32(v36)+228)) = v440
	*(*int32)(unsafe.Add(mBase, uint32(v36)+232)) = v433
	*(*int32)(unsafe.Add(mBase, uint32(v36)+236)) = v442
	*(*int32)(unsafe.Add(mBase, uint32(v36)+240)) = v438
	*(*int32)(unsafe.Add(mBase, uint32(v36)+244)) = v434
	*(*int32)(unsafe.Add(mBase, uint32(v36)+248)) = v431
	*(*int32)(unsafe.Add(mBase, uint32(v36)+252)) = v429
	F_plpgsql_create_econtext(m, l0)
	mBase = m.M
	v479 = m.ExcPending
	if v479 != 0 {
		goto L6
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_exec_stmt_block[0])) = v436
	*(*int32)(unsafe.Add(mBase, _c_F_exec_stmt_block[1])) = v435
	*(*int32)(unsafe.Add(mBase, uint32(v429))) = int32(_a_F_exec_stmt_block_11)
	*(*int32)(unsafe.Add(mBase, _c_F_exec_stmt_block[9])) = v432
	*(*int32)(unsafe.Add(mBase, uint32(v36)+204)) = v436
	*(*int32)(unsafe.Add(mBase, uint32(v36)+208)) = v435
	*(*int32)(unsafe.Add(mBase, uint32(v36)+212)) = v432
	*(*int32)(unsafe.Add(mBase, uint32(v36)+216)) = v439
	*(*int32)(unsafe.Add(mBase, uint32(v36)+220)) = v441
	*(*int32)(unsafe.Add(mBase, uint32(v36)+224)) = v437
	*(*int32)(unsafe.Add(mBase, uint32(v36)+228)) = v440
	*(*int32)(unsafe.Add(mBase, uint32(v36)+232)) = v433
	*(*int32)(unsafe.Add(mBase, uint32(v36)+236)) = v442
	*(*int32)(unsafe.Add(mBase, uint32(v36)+240)) = v438
	*(*int32)(unsafe.Add(mBase, uint32(v36)+244)) = v434
	*(*int32)(unsafe.Add(mBase, uint32(v36)+248)) = v431
	*(*int32)(unsafe.Add(mBase, uint32(v36)+252)) = v429
	v624 = F_CopyErrorData(m)
	mBase = m.M
	v625 = m.ExcPending
	if v625 != 0 {
		goto L6
	} else {
		goto L79
	}
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v429))) = int32(0)
	v482 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v36)+208)) = v435
	*(*int32)(unsafe.Add(mBase, uint32(v36)+204)) = v436
	*(*int32)(unsafe.Add(mBase, uint32(v36)+212)) = v432
	*(*int32)(unsafe.Add(mBase, uint32(v36)+216)) = v439
	*(*int32)(unsafe.Add(mBase, uint32(v36)+220)) = v441
	*(*int32)(unsafe.Add(mBase, uint32(v36)+224)) = v437
	*(*int32)(unsafe.Add(mBase, uint32(v36)+228)) = v440
	*(*int32)(unsafe.Add(mBase, uint32(v36)+232)) = v433
	*(*int32)(unsafe.Add(mBase, uint32(v36)+236)) = v442
	*(*int32)(unsafe.Add(mBase, uint32(v36)+240)) = v438
	*(*int32)(unsafe.Add(mBase, uint32(v36)+244)) = v434
	*(*int32)(unsafe.Add(mBase, uint32(v36)+248)) = v431
	*(*int32)(unsafe.Add(mBase, uint32(v36)+252)) = v429
	v496 = F_exec_stmts(m, l0, v482)
	mBase = m.M
	v497 = m.ExcPending
	if v497 != 0 {
		goto L6
	} else {
		goto L61
	}
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+200)) = v496
	*(*int32)(unsafe.Add(mBase, uint32(v429))) = int32(_a_F_exec_stmt_block_12)
	v501 = *(*int32)(unsafe.Add(mBase, uint32(v36)+200))
	if v501 != int32(2) {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+208)) = v435
	*(*int32)(unsafe.Add(mBase, uint32(v36)+204)) = v436
	*(*int32)(unsafe.Add(mBase, uint32(v36)+212)) = v432
	*(*int32)(unsafe.Add(mBase, uint32(v36)+216)) = v439
	*(*int32)(unsafe.Add(mBase, uint32(v36)+220)) = v441
	*(*int32)(unsafe.Add(mBase, uint32(v36)+224)) = v437
	*(*int32)(unsafe.Add(mBase, uint32(v36)+228)) = v440
	*(*int32)(unsafe.Add(mBase, uint32(v36)+232)) = v433
	*(*int32)(unsafe.Add(mBase, uint32(v36)+236)) = v442
	*(*int32)(unsafe.Add(mBase, uint32(v36)+240)) = v438
	*(*int32)(unsafe.Add(mBase, uint32(v36)+244)) = v434
	*(*int32)(unsafe.Add(mBase, uint32(v36)+248)) = v431
	*(*int32)(unsafe.Add(mBase, uint32(v36)+252)) = v429
	v559 = m.G0
	v561 = v559 - int32(16)
	m.G0 = v561
	v564 = *(*int32)(unsafe.Add(mBase, _c_F_exec_stmt_block[10]))
	v565 = *(*int32)(unsafe.Add(mBase, uint32(v564)+24))
	if v565 != int32(12) {
		goto L68
	} else {
		goto L69
	}
L63:
	;
	v504 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+29)))
	if v504 != 0 {
		goto L62
	} else {
		goto L64
	}
L64:
	;
	v505 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)))
	if v505 != 0 {
		goto L62
	} else {
		goto L65
	}
L65:
	;
	v506 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v36)+208)) = v435
	*(*int32)(unsafe.Add(mBase, uint32(v36)+204)) = v436
	*(*int32)(unsafe.Add(mBase, uint32(v36)+212)) = v432
	*(*int32)(unsafe.Add(mBase, uint32(v36)+216)) = v439
	*(*int32)(unsafe.Add(mBase, uint32(v36)+220)) = v441
	*(*int32)(unsafe.Add(mBase, uint32(v36)+224)) = v437
	*(*int32)(unsafe.Add(mBase, uint32(v36)+228)) = v440
	*(*int32)(unsafe.Add(mBase, uint32(v36)+232)) = v433
	*(*int32)(unsafe.Add(mBase, uint32(v36)+236)) = v442
	*(*int32)(unsafe.Add(mBase, uint32(v36)+240)) = v438
	*(*int32)(unsafe.Add(mBase, uint32(v36)+244)) = v434
	*(*int32)(unsafe.Add(mBase, uint32(v36)+248)) = v431
	*(*int32)(unsafe.Add(mBase, uint32(v36)+252)) = v429
	F_get_typlenbyval(m, v506, v36+int32(30), v36+int32(29))
	mBase = m.M
	v525 = m.ExcPending
	if v525 != 0 {
		goto L6
	} else {
		goto L66
	}
L66:
	;
	v526 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v36)+204)) = v436
	*(*int32)(unsafe.Add(mBase, uint32(v36)+208)) = v435
	*(*int32)(unsafe.Add(mBase, uint32(v36)+212)) = v432
	*(*int32)(unsafe.Add(mBase, uint32(v36)+216)) = v439
	*(*int32)(unsafe.Add(mBase, uint32(v36)+220)) = v441
	*(*int32)(unsafe.Add(mBase, uint32(v36)+224)) = v437
	*(*int32)(unsafe.Add(mBase, uint32(v36)+228)) = v440
	*(*int32)(unsafe.Add(mBase, uint32(v36)+232)) = v433
	*(*int32)(unsafe.Add(mBase, uint32(v36)+236)) = v442
	*(*int32)(unsafe.Add(mBase, uint32(v36)+240)) = v438
	*(*int32)(unsafe.Add(mBase, uint32(v36)+244)) = v434
	*(*int32)(unsafe.Add(mBase, uint32(v36)+248)) = v431
	*(*int32)(unsafe.Add(mBase, uint32(v36)+252)) = v429
	v540 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+29)))
	v541 = int32(*(*int16)(unsafe.Add(mBase, uint32(v36)+30)))
	v542 = F_datumTransfer(m, v526, v540, v541)
	mBase = m.M
	v543 = m.ExcPending
	if v543 != 0 {
		goto L6
	} else {
		goto L67
	}
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v542
	goto L62
L68:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v571 = m.ExcPending
	if v571 != 0 {
		goto L6
	} else {
		goto L71
	}
L69:
	;
	goto L70
L70:
	;
	v591 = *(*int32)(unsafe.Add(mBase, _c_F_exec_stmt_block[11]))
	*(*int32)(unsafe.Add(mBase, _c_F_exec_stmt_block[9])) = v591
	F_CommitSubTransaction(m)
	mBase = m.M
	v594 = m.ExcPending
	if v594 != 0 {
		goto L6
	} else {
		goto L78
	}
L71:
	;
	v572 = *(*int32)(unsafe.Add(mBase, uint32(v564)+24))
	if base.Ui32(v572) <= base.Ui32(int32(19)) {
		goto L73
	} else {
		goto L74
	}
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v561))) = v579
	F_errmsg_internal(m, int32(_a_F_exec_stmt_block_13), v561)
	mBase = m.M
	v583 = m.ExcPending
	if v583 != 0 {
		goto L6
	} else {
		goto L76
	}
L73:
	;
	v577 = *(*int32)(unsafe.Add(mBase, uint32(v572<<(uint(int32(2))%32))+uint32(_c_F_exec_stmt_block[12])))
	v579 = v577
	goto L75
L74:
	;
	v579 = int32(_a_F_exec_stmt_block_14)
	goto L75
L75:
	;
	goto L72
L76:
	;
	F_errfinish(m, int32(_a_F_exec_stmt_block_15), int32(_a_F_exec_stmt_block_16), int32(_a_F_exec_stmt_block_17))
	mBase = m.M
	v588 = m.ExcPending
	if v588 != 0 {
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
	m.G0 = v561 + int32(16)
	*(*int32)(unsafe.Add(mBase, _c_F_exec_stmt_block[8])) = v438
	*(*int32)(unsafe.Add(mBase, _c_F_exec_stmt_block[9])) = v434
	*(*int32)(unsafe.Add(mBase, uint32(v442))) = v433
	goto L8
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+208)) = v435
	*(*int32)(unsafe.Add(mBase, uint32(v36)+204)) = v436
	*(*int32)(unsafe.Add(mBase, uint32(v36)+212)) = v432
	*(*int32)(unsafe.Add(mBase, uint32(v36)+216)) = v439
	*(*int32)(unsafe.Add(mBase, uint32(v36)+220)) = v441
	*(*int32)(unsafe.Add(mBase, uint32(v36)+224)) = v437
	*(*int32)(unsafe.Add(mBase, uint32(v36)+228)) = v440
	*(*int32)(unsafe.Add(mBase, uint32(v36)+232)) = v433
	*(*int32)(unsafe.Add(mBase, uint32(v36)+236)) = v442
	*(*int32)(unsafe.Add(mBase, uint32(v36)+240)) = v438
	*(*int32)(unsafe.Add(mBase, uint32(v36)+244)) = v434
	*(*int32)(unsafe.Add(mBase, uint32(v36)+248)) = v431
	*(*int32)(unsafe.Add(mBase, uint32(v36)+252)) = v429
	F_FlushErrorState(m)
	mBase = m.M
	v640 = m.ExcPending
	if v640 != 0 {
		goto L6
	} else {
		goto L80
	}
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+208)) = v435
	*(*int32)(unsafe.Add(mBase, uint32(v36)+204)) = v436
	*(*int32)(unsafe.Add(mBase, uint32(v36)+212)) = v432
	*(*int32)(unsafe.Add(mBase, uint32(v36)+216)) = v439
	*(*int32)(unsafe.Add(mBase, uint32(v36)+220)) = v441
	*(*int32)(unsafe.Add(mBase, uint32(v36)+224)) = v437
	*(*int32)(unsafe.Add(mBase, uint32(v36)+228)) = v440
	*(*int32)(unsafe.Add(mBase, uint32(v36)+232)) = v433
	*(*int32)(unsafe.Add(mBase, uint32(v36)+236)) = v442
	*(*int32)(unsafe.Add(mBase, uint32(v36)+240)) = v438
	*(*int32)(unsafe.Add(mBase, uint32(v36)+244)) = v434
	*(*int32)(unsafe.Add(mBase, uint32(v36)+248)) = v431
	*(*int32)(unsafe.Add(mBase, uint32(v36)+252)) = v429
	F_RollbackAndReleaseCurrentSubTransaction(m)
	mBase = m.M
	v655 = m.ExcPending
	if v655 != 0 {
		goto L6
	} else {
		goto L81
	}
L81:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_exec_stmt_block[8])) = v438
	*(*int32)(unsafe.Add(mBase, _c_F_exec_stmt_block[9])) = v434
	*(*int32)(unsafe.Add(mBase, uint32(l0)+100)) = v432
	*(*int32)(unsafe.Add(mBase, uint32(v441))) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v36)+208)) = v435
	*(*int32)(unsafe.Add(mBase, uint32(v36)+204)) = v436
	*(*int32)(unsafe.Add(mBase, uint32(v36)+212)) = v432
	*(*int32)(unsafe.Add(mBase, uint32(v36)+216)) = v439
	*(*int32)(unsafe.Add(mBase, uint32(v36)+220)) = v441
	*(*int32)(unsafe.Add(mBase, uint32(v36)+224)) = v437
	*(*int32)(unsafe.Add(mBase, uint32(v36)+228)) = v440
	*(*int32)(unsafe.Add(mBase, uint32(v36)+232)) = v433
	*(*int32)(unsafe.Add(mBase, uint32(v36)+236)) = v442
	*(*int32)(unsafe.Add(mBase, uint32(v36)+240)) = v438
	*(*int32)(unsafe.Add(mBase, uint32(v36)+244)) = v434
	*(*int32)(unsafe.Add(mBase, uint32(v36)+248)) = v431
	*(*int32)(unsafe.Add(mBase, uint32(v36)+252)) = v429
	F_MemoryContextDeleteChildren(m, v432)
	mBase = m.M
	v677 = m.ExcPending
	if v677 != 0 {
		goto L6
	} else {
		goto L82
	}
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v442))) = v433
	*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = int32(0)
	if v433 != 0 {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v681 = *(*int32)(unsafe.Add(mBase, uint32(v433)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v36)+208)) = v435
	*(*int32)(unsafe.Add(mBase, uint32(v36)+204)) = v436
	*(*int32)(unsafe.Add(mBase, uint32(v36)+212)) = v432
	*(*int32)(unsafe.Add(mBase, uint32(v36)+216)) = v439
	*(*int32)(unsafe.Add(mBase, uint32(v36)+220)) = v441
	*(*int32)(unsafe.Add(mBase, uint32(v36)+224)) = v437
	*(*int32)(unsafe.Add(mBase, uint32(v36)+228)) = v440
	*(*int32)(unsafe.Add(mBase, uint32(v36)+232)) = v433
	*(*int32)(unsafe.Add(mBase, uint32(v36)+236)) = v442
	*(*int32)(unsafe.Add(mBase, uint32(v36)+240)) = v438
	*(*int32)(unsafe.Add(mBase, uint32(v36)+244)) = v434
	*(*int32)(unsafe.Add(mBase, uint32(v36)+248)) = v431
	*(*int32)(unsafe.Add(mBase, uint32(v36)+252)) = v429
	F_MemoryContextReset(m, v681)
	mBase = m.M
	v696 = m.ExcPending
	if v696 != 0 {
		goto L6
	} else {
		goto L86
	}
L84:
	;
	goto L85
L85:
	;
	v697 = *(*int32)(unsafe.Add(mBase, uint32(v431)))
	v698 = *(*int32)(unsafe.Add(mBase, uint32(v697)+8))
	if v698 == int32(0) {
		goto L87
	} else {
		goto L88
	}
L86:
	;
	goto L85
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v440))) = v437
	*(*int32)(unsafe.Add(mBase, uint32(v36)+208)) = v435
	*(*int32)(unsafe.Add(mBase, uint32(v36)+204)) = v436
	*(*int32)(unsafe.Add(mBase, uint32(v36)+212)) = v432
	*(*int32)(unsafe.Add(mBase, uint32(v36)+216)) = v439
	*(*int32)(unsafe.Add(mBase, uint32(v36)+220)) = v441
	*(*int32)(unsafe.Add(mBase, uint32(v36)+224)) = v437
	*(*int32)(unsafe.Add(mBase, uint32(v36)+228)) = v440
	*(*int32)(unsafe.Add(mBase, uint32(v36)+232)) = v433
	*(*int32)(unsafe.Add(mBase, uint32(v36)+236)) = v442
	*(*int32)(unsafe.Add(mBase, uint32(v36)+240)) = v438
	*(*int32)(unsafe.Add(mBase, uint32(v36)+244)) = v434
	*(*int32)(unsafe.Add(mBase, uint32(v36)+248)) = v431
	*(*int32)(unsafe.Add(mBase, uint32(v36)+252)) = v429
	F_ReThrowError(m, v624)
	mBase = m.M
	v880 = m.ExcPending
	if v880 != 0 {
		goto L6
	} else {
		goto L108
	}
L88:
	;
	v701 = *(*int32)(unsafe.Add(mBase, uint32(v698)+4))
	if v701 <= int32(0) {
		goto L87
	} else {
		goto L89
	}
L89:
	;
	v704 = int32(0)
	if v704 < v701 {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v707 = v701
	goto L92
L91:
	;
	v707 = v704
	goto L92
L92:
	;
	v708 = *(*int32)(unsafe.Add(mBase, uint32(v698)+12))
	v731 = int32(0)
	goto L93
L93:
	;
	v746 = *(*int32)(unsafe.Add(mBase, uint32(v708+v731<<(uint(int32(2))%32))))
	v747 = *(*int32)(unsafe.Add(mBase, uint32(v746)+4))
	if v747 != 0 {
		goto L95
	} else {
		goto L96
	}
L94:
	;
	goto L87
L95:
	;
	v748 = *(*int32)(unsafe.Add(mBase, uint32(v624)+28))
	v759 = v747
	goto L98
L96:
	;
	goto L97
L97:
	;
	v830 = v731 + int32(1)
	if v830 != v707 {
		v731 = v830
		goto L93
	} else {
		goto L107
	}
L98:
	;
	v788 = *(*int32)(unsafe.Add(mBase, uint32(v759)))
	if v788 == int32(-1) {
		goto L101
	} else {
		goto L102
	}
L99:
	;
	goto L97
L100:
	;
	v795 = *(*int32)(unsafe.Add(mBase, uint32(v759)+8))
	if v795 != 0 {
		v759 = v795
		goto L98
	} else {
		goto L106
	}
L101:
	;
	if base.B2i32(v748 == int32(67108896))|base.B2i32(v748 == int32(67371461)) != 0 {
		goto L100
	} else {
		goto L104
	}
L102:
	;
	goto L103
L103:
	;
	if base.B2i32(v748 == v788)|base.B2i32(v788 == v748&int32(4095)) != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v36)+200)) = v897
	v1124 = v39
	v1126 = v41
	v1127 = v55
	v1128 = v56
	v1129 = v57
	v1130 = v58
	v1131 = v59
	v1132 = v60
	v1133 = v61
	v1134 = v62
	v1135 = v63
	v1136 = v64
	v1137 = v65
	goto L7
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+208)) = v435
	*(*int32)(unsafe.Add(mBase, uint32(v36)+204)) = v436
	*(*int32)(unsafe.Add(mBase, uint32(v36)+212)) = v432
	*(*int32)(unsafe.Add(mBase, uint32(v36)+216)) = v439
	*(*int32)(unsafe.Add(mBase, uint32(v36)+220)) = v441
	*(*int32)(unsafe.Add(mBase, uint32(v36)+224)) = v437
	*(*int32)(unsafe.Add(mBase, uint32(v36)+228)) = v440
	*(*int32)(unsafe.Add(mBase, uint32(v36)+232)) = v433
	*(*int32)(unsafe.Add(mBase, uint32(v36)+236)) = v442
	*(*int32)(unsafe.Add(mBase, uint32(v36)+240)) = v438
	*(*int32)(unsafe.Add(mBase, uint32(v36)+244)) = v434
	*(*int32)(unsafe.Add(mBase, uint32(v36)+248)) = v431
	*(*int32)(unsafe.Add(mBase, uint32(v36)+252)) = v429
	v979 = F_cstring_to_text(m, v924)
	mBase = m.M
	v980 = m.ExcPending
	if v980 != 0 {
		goto L6
	} else {
		goto L111
	}
L111:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+208)) = v435
	*(*int32)(unsafe.Add(mBase, uint32(v36)+204)) = v436
	*(*int32)(unsafe.Add(mBase, uint32(v36)+212)) = v432
	*(*int32)(unsafe.Add(mBase, uint32(v36)+216)) = v439
	*(*int32)(unsafe.Add(mBase, uint32(v36)+220)) = v441
	*(*int32)(unsafe.Add(mBase, uint32(v36)+224)) = v437
	*(*int32)(unsafe.Add(mBase, uint32(v36)+228)) = v440
	*(*int32)(unsafe.Add(mBase, uint32(v36)+232)) = v433
	*(*int32)(unsafe.Add(mBase, uint32(v36)+236)) = v442
	*(*int32)(unsafe.Add(mBase, uint32(v36)+240)) = v438
	*(*int32)(unsafe.Add(mBase, uint32(v36)+244)) = v434
	*(*int32)(unsafe.Add(mBase, uint32(v36)+248)) = v431
	*(*int32)(unsafe.Add(mBase, uint32(v36)+252)) = v429
	F_assign_simple_var(m, l0, v910, v979, int32(0), int32(1))
	mBase = m.M
	v997 = m.ExcPending
	if v997 != 0 {
		goto L6
	} else {
		goto L112
	}
L112:
	;
	v998 = *(*int32)(unsafe.Add(mBase, uint32(v624)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v36)+208)) = v435
	*(*int32)(unsafe.Add(mBase, uint32(v36)+204)) = v436
	*(*int32)(unsafe.Add(mBase, uint32(v36)+212)) = v432
	*(*int32)(unsafe.Add(mBase, uint32(v36)+216)) = v439
	*(*int32)(unsafe.Add(mBase, uint32(v36)+220)) = v441
	*(*int32)(unsafe.Add(mBase, uint32(v36)+224)) = v437
	*(*int32)(unsafe.Add(mBase, uint32(v36)+228)) = v440
	*(*int32)(unsafe.Add(mBase, uint32(v36)+232)) = v433
	*(*int32)(unsafe.Add(mBase, uint32(v36)+236)) = v442
	*(*int32)(unsafe.Add(mBase, uint32(v36)+240)) = v438
	*(*int32)(unsafe.Add(mBase, uint32(v36)+244)) = v434
	*(*int32)(unsafe.Add(mBase, uint32(v36)+248)) = v431
	*(*int32)(unsafe.Add(mBase, uint32(v36)+252)) = v429
	v1012 = F_cstring_to_text(m, v998)
	mBase = m.M
	v1013 = m.ExcPending
	if v1013 != 0 {
		goto L6
	} else {
		goto L113
	}
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+208)) = v435
	*(*int32)(unsafe.Add(mBase, uint32(v36)+204)) = v436
	*(*int32)(unsafe.Add(mBase, uint32(v36)+212)) = v432
	*(*int32)(unsafe.Add(mBase, uint32(v36)+216)) = v439
	*(*int32)(unsafe.Add(mBase, uint32(v36)+220)) = v441
	*(*int32)(unsafe.Add(mBase, uint32(v36)+224)) = v437
	*(*int32)(unsafe.Add(mBase, uint32(v36)+228)) = v440
	*(*int32)(unsafe.Add(mBase, uint32(v36)+232)) = v433
	*(*int32)(unsafe.Add(mBase, uint32(v36)+236)) = v442
	*(*int32)(unsafe.Add(mBase, uint32(v36)+240)) = v438
	*(*int32)(unsafe.Add(mBase, uint32(v36)+244)) = v434
	*(*int32)(unsafe.Add(mBase, uint32(v36)+248)) = v431
	*(*int32)(unsafe.Add(mBase, uint32(v36)+252)) = v429
	F_assign_simple_var(m, l0, v905, v1012, int32(0), int32(1))
	mBase = m.M
	v1030 = m.ExcPending
	if v1030 != 0 {
		goto L6
	} else {
		goto L114
	}
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v440))) = v624
	*(*int32)(unsafe.Add(mBase, uint32(v429))) = int32(0)
	v1034 = *(*int32)(unsafe.Add(mBase, uint32(v746)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v36)+208)) = v435
	*(*int32)(unsafe.Add(mBase, uint32(v36)+204)) = v436
	*(*int32)(unsafe.Add(mBase, uint32(v36)+212)) = v432
	*(*int32)(unsafe.Add(mBase, uint32(v36)+216)) = v439
	*(*int32)(unsafe.Add(mBase, uint32(v36)+220)) = v441
	*(*int32)(unsafe.Add(mBase, uint32(v36)+224)) = v437
	*(*int32)(unsafe.Add(mBase, uint32(v36)+228)) = v440
	*(*int32)(unsafe.Add(mBase, uint32(v36)+232)) = v433
	*(*int32)(unsafe.Add(mBase, uint32(v36)+236)) = v442
	*(*int32)(unsafe.Add(mBase, uint32(v36)+240)) = v438
	*(*int32)(unsafe.Add(mBase, uint32(v36)+244)) = v434
	*(*int32)(unsafe.Add(mBase, uint32(v36)+248)) = v431
	*(*int32)(unsafe.Add(mBase, uint32(v36)+252)) = v429
	v1048 = F_exec_stmts(m, l0, v1034)
	mBase = m.M
	v1049 = m.ExcPending
	if v1049 != 0 {
		goto L6
	} else {
		goto L115
	}
L115:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+200)) = v1048
	*(*int32)(unsafe.Add(mBase, uint32(v440))) = v437
	v1052 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	*(*int32)(unsafe.Add(mBase, uint32(v441))) = v1052
	*(*int32)(unsafe.Add(mBase, uint32(v36)+208)) = v435
	*(*int32)(unsafe.Add(mBase, uint32(v36)+204)) = v436
	*(*int32)(unsafe.Add(mBase, uint32(v36)+212)) = v432
	*(*int32)(unsafe.Add(mBase, uint32(v36)+216)) = v439
	*(*int32)(unsafe.Add(mBase, uint32(v36)+220)) = v441
	*(*int32)(unsafe.Add(mBase, uint32(v36)+224)) = v437
	*(*int32)(unsafe.Add(mBase, uint32(v36)+228)) = v440
	*(*int32)(unsafe.Add(mBase, uint32(v36)+232)) = v433
	*(*int32)(unsafe.Add(mBase, uint32(v36)+236)) = v442
	*(*int32)(unsafe.Add(mBase, uint32(v36)+240)) = v438
	*(*int32)(unsafe.Add(mBase, uint32(v36)+244)) = v434
	*(*int32)(unsafe.Add(mBase, uint32(v36)+248)) = v431
	*(*int32)(unsafe.Add(mBase, uint32(v36)+252)) = v429
	v1067 = *(*int32)(unsafe.Add(mBase, uint32(v1052)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+100)) = v1067
	*(*int32)(unsafe.Add(mBase, uint32(v36)+208)) = v435
	*(*int32)(unsafe.Add(mBase, uint32(v36)+204)) = v436
	*(*int32)(unsafe.Add(mBase, uint32(v36)+212)) = v432
	*(*int32)(unsafe.Add(mBase, uint32(v36)+216)) = v439
	*(*int32)(unsafe.Add(mBase, uint32(v36)+220)) = v441
	*(*int32)(unsafe.Add(mBase, uint32(v36)+224)) = v437
	*(*int32)(unsafe.Add(mBase, uint32(v36)+228)) = v440
	*(*int32)(unsafe.Add(mBase, uint32(v36)+232)) = v433
	*(*int32)(unsafe.Add(mBase, uint32(v36)+236)) = v442
	*(*int32)(unsafe.Add(mBase, uint32(v36)+240)) = v438
	*(*int32)(unsafe.Add(mBase, uint32(v36)+244)) = v434
	*(*int32)(unsafe.Add(mBase, uint32(v36)+248)) = v431
	*(*int32)(unsafe.Add(mBase, uint32(v36)+252)) = v429
	F_MemoryContextReset(m, v432)
	mBase = m.M
	v1083 = m.ExcPending
	if v1083 != 0 {
		goto L6
	} else {
		goto L116
	}
L116:
	;
	goto L8
L117:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+208)) = v1130
	*(*int32)(unsafe.Add(mBase, uint32(v36)+204)) = v1131
	*(*int32)(unsafe.Add(mBase, uint32(v36)+212)) = v1127
	*(*int32)(unsafe.Add(mBase, uint32(v36)+216)) = v1134
	*(*int32)(unsafe.Add(mBase, uint32(v36)+220)) = v1136
	*(*int32)(unsafe.Add(mBase, uint32(v36)+224)) = v1132
	*(*int32)(unsafe.Add(mBase, uint32(v36)+228)) = v1135
	*(*int32)(unsafe.Add(mBase, uint32(v36)+232)) = v1128
	*(*int32)(unsafe.Add(mBase, uint32(v36)+236)) = v1137
	*(*int32)(unsafe.Add(mBase, uint32(v36)+240)) = v1133
	*(*int32)(unsafe.Add(mBase, uint32(v36)+244)) = v1129
	*(*int32)(unsafe.Add(mBase, uint32(v36)+248)) = v1126
	*(*int32)(unsafe.Add(mBase, uint32(v36)+252)) = v1124
	F_errstart_cold(m, int32(21), int32(_a_F_exec_stmt_block_2))
	mBase = m.M
	v1237 = m.ExcPending
	if v1237 != 0 {
		goto L6
	} else {
		goto L133
	}
L118:
	;
	m.G0 = v36 + int32(256)
	return v1214
L119:
	;
	if base.B2i32(v1156 == int32(1))|base.B2i32(base.Ui32(int32(3)) < base.Ui32(v1156)) != 0 {
		goto L117
	} else {
		goto L122
	}
L120:
	;
	goto L121
L121:
	;
	v1165 = int32(1)
	v1166 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v1166 == int32(0) {
		v1214 = v1165
		goto L118
	} else {
		goto L123
	}
L122:
	;
	v1164 = *(*int32)(unsafe.Add(mBase, uint32(v36)+200))
	v1214 = v1164
	goto L118
L123:
	;
	v1169 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v1169 == int32(0) {
		v1214 = v1165
		goto L118
	} else {
		goto L124
	}
L124:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+208)) = v1130
	*(*int32)(unsafe.Add(mBase, uint32(v36)+204)) = v1131
	*(*int32)(unsafe.Add(mBase, uint32(v36)+212)) = v1127
	*(*int32)(unsafe.Add(mBase, uint32(v36)+216)) = v1134
	*(*int32)(unsafe.Add(mBase, uint32(v36)+220)) = v1136
	*(*int32)(unsafe.Add(mBase, uint32(v36)+224)) = v1132
	*(*int32)(unsafe.Add(mBase, uint32(v36)+228)) = v1135
	*(*int32)(unsafe.Add(mBase, uint32(v36)+232)) = v1128
	*(*int32)(unsafe.Add(mBase, uint32(v36)+236)) = v1137
	*(*int32)(unsafe.Add(mBase, uint32(v36)+240)) = v1133
	*(*int32)(unsafe.Add(mBase, uint32(v36)+244)) = v1129
	*(*int32)(unsafe.Add(mBase, uint32(v36)+248)) = v1126
	*(*int32)(unsafe.Add(mBase, uint32(v36)+252)) = v1124
	v1187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1169))))
	v1190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1166))))
	if base.B2i32(v1187 == int32(0))|base.B2i32(v1187 != v1190) != 0 {
		v1208 = v1187
		v1209 = v1190
		goto L126
	} else {
		goto L127
	}
L125:
	;
	if v1208-v1209 != 0 {
		v1214 = v1165
		goto L118
	} else {
		goto L132
	}
L126:
	;
	goto L125
L127:
	;
	v1193 = v1169
	v1194 = v1166
	goto L128
L128:
	;
	v1197 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1194)+1)))
	v1198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1193)+1)))
	if v1198 == int32(0) {
		v1208 = v1198
		v1209 = v1197
		goto L126
	} else {
		goto L130
	}
L129:
	;
	v1208 = v1198
	v1209 = v1197
	goto L126
L130:
	;
	v1201 = int32(1)
	if v1198 == v1197 {
		v1193 = v1193 + v1201
		v1194 = v1194 + v1201
		goto L128
	} else {
		goto L131
	}
L131:
	;
	goto L129
L132:
	;
	v1211 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v1211
	v1214 = v1211
	goto L118
L133:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+204)) = v1131
	*(*int32)(unsafe.Add(mBase, uint32(v36)+208)) = v1130
	*(*int32)(unsafe.Add(mBase, uint32(v36)+212)) = v1127
	*(*int32)(unsafe.Add(mBase, uint32(v36)+216)) = v1134
	*(*int32)(unsafe.Add(mBase, uint32(v36)+220)) = v1136
	*(*int32)(unsafe.Add(mBase, uint32(v36)+224)) = v1132
	*(*int32)(unsafe.Add(mBase, uint32(v36)+228)) = v1135
	*(*int32)(unsafe.Add(mBase, uint32(v36)+232)) = v1128
	*(*int32)(unsafe.Add(mBase, uint32(v36)+236)) = v1137
	*(*int32)(unsafe.Add(mBase, uint32(v36)+240)) = v1133
	*(*int32)(unsafe.Add(mBase, uint32(v36)+244)) = v1129
	*(*int32)(unsafe.Add(mBase, uint32(v36)+248)) = v1126
	*(*int32)(unsafe.Add(mBase, uint32(v36)+252)) = v1124
	v1251 = *(*int32)(unsafe.Add(mBase, uint32(v36)+200))
	*(*int32)(unsafe.Add(mBase, uint32(v36))) = v1251
	F_errmsg_internal(m, int32(_a_F_exec_stmt_block_18), v36)
	mBase = m.M
	v1255 = m.ExcPending
	if v1255 != 0 {
		goto L6
	} else {
		goto L134
	}
L134:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+208)) = v1130
	*(*int32)(unsafe.Add(mBase, uint32(v36)+204)) = v1131
	*(*int32)(unsafe.Add(mBase, uint32(v36)+212)) = v1127
	*(*int32)(unsafe.Add(mBase, uint32(v36)+216)) = v1134
	*(*int32)(unsafe.Add(mBase, uint32(v36)+220)) = v1136
	*(*int32)(unsafe.Add(mBase, uint32(v36)+224)) = v1132
	*(*int32)(unsafe.Add(mBase, uint32(v36)+228)) = v1135
	*(*int32)(unsafe.Add(mBase, uint32(v36)+232)) = v1128
	*(*int32)(unsafe.Add(mBase, uint32(v36)+236)) = v1137
	*(*int32)(unsafe.Add(mBase, uint32(v36)+240)) = v1133
	*(*int32)(unsafe.Add(mBase, uint32(v36)+244)) = v1129
	*(*int32)(unsafe.Add(mBase, uint32(v36)+248)) = v1126
	*(*int32)(unsafe.Add(mBase, uint32(v36)+252)) = v1124
	F_errfinish(m, int32(_a_F_exec_stmt_block_5), int32(1983), int32(_a_F_exec_stmt_block_6))
	mBase = m.M
	v1273 = m.ExcPending
	if v1273 != 0 {
		goto L6
	} else {
		goto L135
	}
L135:
	;
	goto L5
L136:
	;
	v1312 = int32(v1308)
	m.G0 = v36
	v1314 = *(*int32)(unsafe.Add(mBase, uint32(v1312)+4))
	v1315 = *(*int32)(unsafe.Add(mBase, uint32(v1312)))
	v1318 = *(*int32)(unsafe.Add(mBase, uint32(v1315)))
	if v36+int32(24) == v1318 {
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
	if v1322 != 0 {
		goto L142
	} else {
		goto L143
	}
L139:
	;
	v1320 = *(*int32)(unsafe.Add(mBase, uint32(v1315)+4))
	v1322 = v1320
	goto L141
L140:
	;
	v1322 = int32(0)
	goto L141
L141:
	;
	goto L138
L142:
	;
	v1323 = *(*int32)(unsafe.Add(mBase, uint32(v36)+252))
	v1324 = *(*int32)(unsafe.Add(mBase, uint32(v36)+248))
	v1325 = *(*int32)(unsafe.Add(mBase, uint32(v36)+244))
	v1326 = *(*int32)(unsafe.Add(mBase, uint32(v36)+240))
	v1327 = *(*int32)(unsafe.Add(mBase, uint32(v36)+236))
	v1328 = *(*int32)(unsafe.Add(mBase, uint32(v36)+232))
	v1329 = *(*int32)(unsafe.Add(mBase, uint32(v36)+228))
	v1330 = *(*int32)(unsafe.Add(mBase, uint32(v36)+224))
	v1331 = *(*int32)(unsafe.Add(mBase, uint32(v36)+220))
	v1332 = *(*int32)(unsafe.Add(mBase, uint32(v36)+216))
	v1333 = *(*int32)(unsafe.Add(mBase, uint32(v36)+212))
	v1334 = *(*int32)(unsafe.Add(mBase, uint32(v36)+208))
	v1335 = *(*int32)(unsafe.Add(mBase, uint32(v36)+204))
	v52 = v1323
	v53 = v1314
	v54 = v1324
	v55 = v1333
	v56 = v1328
	v57 = v1325
	v58 = v1334
	v59 = v1335
	v60 = v1330
	v61 = v1326
	v62 = v1332
	v63 = v1329
	v64 = v1331
	v65 = v1327
	v66 = v1322
	goto L1
L143:
	;
	goto L144
L144:
	;
	F___wasm_longjmp(m, v1315, v1314)
	mBase = m.M
	v1339 = m.ExcPending
	if v1339 != 0 {
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
