package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_pgss_ExecutorStart(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int64
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_ExecutorStart[0]))
	if v5 != 0 {
		m.T0[v5].(func(*base.Module, int32, int32))(m, l0, l1)
		mBase = m.M
		v7 = m.ExcPending
		if v7 != 0 {
			return
		} else {
			v11 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_ExecutorStart[1]))
			if int32(0) <= v11 {
				return
			} else {
				v15 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_ExecutorStart[2]))
				if v15 != int32(2) {
					if v15 != int32(1) {
						return
					} else {
						v21 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_ExecutorStart[3]))
						if v21 != 0 {
							return
						} else {
							v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							v23 = *(*int64)(unsafe.Add(mBase, uint32(v22)+8))
							if v23 == int64(0) {
								return
							} else {
								v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
								if v26 != 0 {
									return
								} else {
									v27 = int32(_a_F_pgss_ExecutorStart_0)
									v28 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_ExecutorStart[4]))
									v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
									v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+100))
									*(*int32)(unsafe.Add(mBase, _c_F_pgss_ExecutorStart[4])) = v31
									v36 = F_InstrAlloc(m, int32(1), int32(2147483647), int32(0))
									mBase = m.M
									v37 = m.ExcPending
									if v37 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v36
										*(*int32)(unsafe.Add(mBase, _c_F_pgss_ExecutorStart[4])) = v28
										return
									}
								}
							}
						}
					}
				} else {
					v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v23 = *(*int64)(unsafe.Add(mBase, uint32(v22)+8))
					if v23 == int64(0) {
						return
					} else {
						v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
						if v26 != 0 {
							return
						} else {
							v27 = int32(_a_F_pgss_ExecutorStart_0)
							v28 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_ExecutorStart[4]))
							v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
							v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+100))
							*(*int32)(unsafe.Add(mBase, _c_F_pgss_ExecutorStart[4])) = v31
							v36 = F_InstrAlloc(m, int32(1), int32(2147483647), int32(0))
							mBase = m.M
							v37 = m.ExcPending
							if v37 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v36
								*(*int32)(unsafe.Add(mBase, _c_F_pgss_ExecutorStart[4])) = v28
								return
							}
						}
					}
				}
			}
		}
	} else {
		F_standard_ExecutorStart(m, l0, l1)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return
		} else {
			v11 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_ExecutorStart[1]))
			if int32(0) <= v11 {
				return
			} else {
				v15 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_ExecutorStart[2]))
				if v15 != int32(2) {
					if v15 != int32(1) {
						return
					} else {
						v21 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_ExecutorStart[3]))
						if v21 != 0 {
							return
						} else {
							v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							v23 = *(*int64)(unsafe.Add(mBase, uint32(v22)+8))
							if v23 == int64(0) {
								return
							} else {
								v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
								if v26 != 0 {
									return
								} else {
									v27 = int32(_a_F_pgss_ExecutorStart_0)
									v28 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_ExecutorStart[4]))
									v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
									v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+100))
									*(*int32)(unsafe.Add(mBase, _c_F_pgss_ExecutorStart[4])) = v31
									v36 = F_InstrAlloc(m, int32(1), int32(2147483647), int32(0))
									mBase = m.M
									v37 = m.ExcPending
									if v37 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v36
										*(*int32)(unsafe.Add(mBase, _c_F_pgss_ExecutorStart[4])) = v28
										return
									}
								}
							}
						}
					}
				} else {
					v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v23 = *(*int64)(unsafe.Add(mBase, uint32(v22)+8))
					if v23 == int64(0) {
						return
					} else {
						v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
						if v26 != 0 {
							return
						} else {
							v27 = int32(_a_F_pgss_ExecutorStart_0)
							v28 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_ExecutorStart[4]))
							v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
							v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+100))
							*(*int32)(unsafe.Add(mBase, _c_F_pgss_ExecutorStart[4])) = v31
							v36 = F_InstrAlloc(m, int32(1), int32(2147483647), int32(0))
							mBase = m.M
							v37 = m.ExcPending
							if v37 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v36
								*(*int32)(unsafe.Add(mBase, _c_F_pgss_ExecutorStart[4])) = v28
								return
							}
						}
					}
				}
			}
		}
	}
}
func F_pgss_store(m *base.Module, l0 int32, l1 int64, l2 int32, l3 int32, l4 int32, l5 float64, l6 int64, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32, l12 int32) {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v90 int32
	_ = v90
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v124 int64
	_ = v124
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
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
	var v200 int32
	_ = v200
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v302 int32
	_ = v302
	var v317 int32
	_ = v317
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v366 int32
	_ = v366
	var v377 int32
	_ = v377
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v383 int32
	_ = v383
	var v387 int32
	_ = v387
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v400 int32
	_ = v400
	var v403 int32
	_ = v403
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v413 int32
	_ = v413
	var v415 int32
	_ = v415
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v423 int32
	_ = v423
	var v425 int32
	_ = v425
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v434 int32
	_ = v434
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
	var v443 int32
	_ = v443
	var v462 int32
	_ = v462
	var v465 int32
	_ = v465
	var v475 int32
	_ = v475
	var v479 int32
	_ = v479
	var v481 int32
	_ = v481
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v505 int32
	_ = v505
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v534 int32
	_ = v534
	var v541 int32
	_ = v541
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v547 int64
	_ = v547
	var v549 int64
	_ = v549
	var v553 int64
	_ = v553
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v561 int32
	_ = v561
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v573 int32
	_ = v573
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v592 int32
	_ = v592
	var v593 int32
	_ = v593
	var v597 int32
	_ = v597
	var v604 int32
	_ = v604
	var v606 int32
	_ = v606
	var v609 int64
	_ = v609
	var v611 int64
	_ = v611
	var v615 int64
	_ = v615
	var v622 int32
	_ = v622
	var v623 int32
	_ = v623
	var v628 int32
	_ = v628
	var v629 int32
	_ = v629
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v639 int32
	_ = v639
	var v646 int32
	_ = v646
	var v651 int32
	_ = v651
	var v653 int32
	_ = v653
	var v655 int32
	_ = v655
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v662 int32
	_ = v662
	var v664 int32
	_ = v664
	var v665 int32
	_ = v665
	var v669 int32
	_ = v669
	var v681 int32
	_ = v681
	var v685 int32
	_ = v685
	var v697 int32
	_ = v697
	var v700 int32
	_ = v700
	var v701 int32
	_ = v701
	var v704 int32
	_ = v704
	var v712 int32
	_ = v712
	var v714 int32
	_ = v714
	var v715 int32
	_ = v715
	var v716 int32
	_ = v716
	var v720 int32
	_ = v720
	var v721 int32
	_ = v721
	var v723 int32
	_ = v723
	var v730 int32
	_ = v730
	var v735 int32
	_ = v735
	var v739 int32
	_ = v739
	var v740 int32
	_ = v740
	var v741 int32
	_ = v741
	var v746 int32
	_ = v746
	var v748 int32
	_ = v748
	var v753 int32
	_ = v753
	var v754 int32
	_ = v754
	var v769 int32
	_ = v769
	var v773 int32
	_ = v773
	var v785 int32
	_ = v785
	var v792 int32
	_ = v792
	var v794 int32
	_ = v794
	var v799 int32
	_ = v799
	var v800 int32
	_ = v800
	var v804 int32
	_ = v804
	var v811 int32
	_ = v811
	var v816 int32
	_ = v816
	var v817 int32
	_ = v817
	var v818 int32
	_ = v818
	var v821 int32
	_ = v821
	var v822 int32
	_ = v822
	var v826 int32
	_ = v826
	var v833 int32
	_ = v833
	var v838 int32
	_ = v838
	var v841 int32
	_ = v841
	var v842 int32
	_ = v842
	var v844 int32
	_ = v844
	var v845 int32
	_ = v845
	var v852 int32
	_ = v852
	var v857 int32
	_ = v857
	var v859 int32
	_ = v859
	var v864 int32
	_ = v864
	var v865 int32
	_ = v865
	var v868 int32
	_ = v868
	var v906 int32
	_ = v906
	var v908 int32
	_ = v908
	var v910 int32
	_ = v910
	var v911 int32
	_ = v911
	var v912 int32
	_ = v912
	var v915 int32
	_ = v915
	var v947 int32
	_ = v947
	var v948 int32
	_ = v948
	var v979 int32
	_ = v979
	var v980 int32
	_ = v980
	var v983 int32
	_ = v983
	var v984 int32
	_ = v984
	var v989 int32
	_ = v989
	var v990 int32
	_ = v990
	var v994 int32
	_ = v994
	var v999 int32
	_ = v999
	var v1004 int32
	_ = v1004
	var v1005 int32
	_ = v1005
	var v1006 int32
	_ = v1006
	var v1008 int32
	_ = v1008
	var v1011 int32
	_ = v1011
	var v1013 int32
	_ = v1013
	var v1049 int32
	_ = v1049
	var v1051 int32
	_ = v1051
	var v1057 int32
	_ = v1057
	var v1089 int32
	_ = v1089
	var v1092 int32
	_ = v1092
	var v1099 int32
	_ = v1099
	var v1111 int32
	_ = v1111
	var v1126 int32
	_ = v1126
	var v1135 int32
	_ = v1135
	var v1136 int64
	_ = v1136
	var v1138 int64
	_ = v1138
	var v1144 int32
	_ = v1144
	var v1147 int32
	_ = v1147
	var v1148 int64
	_ = v1148
	var v1150 int64
	_ = v1150
	var v1152 int32
	_ = v1152
	var v1154 int32
	_ = v1154
	var v1155 float64
	_ = v1155
	var v1164 int32
	_ = v1164
	var v1165 float64
	_ = v1165
	var v1166 float64
	_ = v1166
	var v1169 float64
	_ = v1169
	var v1172 int32
	_ = v1172
	var v1175 float64
	_ = v1175
	var v1179 int32
	_ = v1179
	var v1180 float64
	_ = v1180
	var v1187 int32
	_ = v1187
	var v1188 float64
	_ = v1188
	var v1200 int32
	_ = v1200
	var v1201 float64
	_ = v1201
	var v1210 int64
	_ = v1210
	var v1213 int64
	_ = v1213
	var v1214 int64
	_ = v1214
	var v1217 int64
	_ = v1217
	var v1218 int64
	_ = v1218
	var v1221 int64
	_ = v1221
	var v1222 int64
	_ = v1222
	var v1225 int64
	_ = v1225
	var v1226 int64
	_ = v1226
	var v1229 int64
	_ = v1229
	var v1230 int64
	_ = v1230
	var v1233 int64
	_ = v1233
	var v1234 int64
	_ = v1234
	var v1237 int64
	_ = v1237
	var v1238 int64
	_ = v1238
	var v1241 int64
	_ = v1241
	var v1242 int64
	_ = v1242
	var v1245 int64
	_ = v1245
	var v1246 int64
	_ = v1246
	var v1249 int64
	_ = v1249
	var v1250 int64
	_ = v1250
	var v1253 float64
	_ = v1253
	var v1254 int64
	_ = v1254
	var v1256 float64
	_ = v1256
	var v1260 float64
	_ = v1260
	var v1261 int64
	_ = v1261
	var v1267 float64
	_ = v1267
	var v1268 int64
	_ = v1268
	var v1274 float64
	_ = v1274
	var v1275 int64
	_ = v1275
	var v1281 float64
	_ = v1281
	var v1282 int64
	_ = v1282
	var v1288 float64
	_ = v1288
	var v1289 int64
	_ = v1289
	var v1295 float64
	_ = v1295
	var v1299 int64
	_ = v1299
	var v1300 int64
	_ = v1300
	var v1303 int64
	_ = v1303
	var v1304 int64
	_ = v1304
	var v1307 int64
	_ = v1307
	var v1308 int64
	_ = v1308
	var v1311 int64
	_ = v1311
	var v1312 int64
	_ = v1312
	var v1315 int64
	_ = v1315
	var v1316 int64
	_ = v1316
	var v1319 float64
	_ = v1319
	var v1320 int64
	_ = v1320
	var v1322 float64
	_ = v1322
	var v1326 int64
	_ = v1326
	var v1329 float64
	_ = v1329
	var v1332 int64
	_ = v1332
	var v1336 float64
	_ = v1336
	var v1339 int64
	_ = v1339
	var v1342 float64
	_ = v1342
	var v1345 int64
	_ = v1345
	var v1349 float64
	_ = v1349
	var v1352 int64
	_ = v1352
	var v1355 float64
	_ = v1355
	var v1358 int64
	_ = v1358
	var v1362 float64
	_ = v1362
	var v1365 int64
	_ = v1365
	var v1368 float64
	_ = v1368
	var v1371 int64
	_ = v1371
	var v1375 float64
	_ = v1375
	var v1381 int64
	_ = v1381
	var v1385 int64
	_ = v1385
	var v1404 int32
	_ = v1404
	var v1420 int32
	_ = v1420
	var v1421 int32
	_ = v1421
	var v1423 int32
	_ = v1423
	var v1427 int32
	_ = v1427
	v14 = int32(0)
	v31 = m.G0
	v33 = v31 - int32(224)
	m.G0 = v33
	*(*int32)(unsafe.Add(mBase, uint32(v33)+148)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v33)+152)) = l2
	v38 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_store[0]))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
	goto L1
L1:
	;
	if l1 == int64(0) {
		goto L2
	} else {
		goto L3
	}
L2:
	;
	m.G0 = v33 + int32(224)
	return
L3:
	;
	v43 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_store[1]))
	if v43 == int32(0) {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v47 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_store[2]))
	if v47 == int32(0) {
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v51 = v33 + int32(152)
	v53 = v33 + int32(148)
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
	if v58 < int32(0) {
		goto L10
	} else {
		goto L11
	}
L6:
	;
	v124 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v33)+136)) = v124
	*(*int64)(unsafe.Add(mBase, uint32(v33)+128)) = v124
	*(*int64)(unsafe.Add(mBase, uint32(v33)+120)) = v124
	v131 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_store[3]))
	*(*int64)(unsafe.Add(mBase, uint32(v33)+128)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v33)+120)) = v131
	v135 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_store[4]))
	*(*int32)(unsafe.Add(mBase, uint32(v33)+124)) = v135
	v138 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_store[5]))
	*(*uint8)(unsafe.Add(mBase, uint32(v33)+136)) = uint8(base.B2i32(v138 == int32(0)))
	v143 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_store[1]))
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v143)))
	v146 = F_LWLockAcquire(m, v144, int32(1))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L25
	} else {
		goto L26
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v51))) = v119
	*(*int32)(unsafe.Add(mBase, uint32(v53))) = v118
	goto L6
L8:
	;
	v76 = v72
	v79 = v73
	v80 = v74
	goto L15
L9:
	;
	v69 = F_strlen(m, v66)
	mBase = m.M
	if v69 <= int32(0) {
		v115 = v66
		v118 = v69
		v119 = v68
		goto L7
	} else {
		goto L14
	}
L10:
	;
	v66 = l0
	v68 = int32(0)
	goto L9
L11:
	;
	goto L12
L12:
	;
	v62 = l0 + v58
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
	if int32(0) < v63 {
		v72 = v62
		v73 = v63
		v74 = v58
		goto L8
	} else {
		goto L13
	}
L13:
	;
	v66 = v62
	v68 = v58
	goto L9
L14:
	;
	v72 = v66
	v73 = v69
	v74 = v68
	goto L8
L15:
	;
	v83 = int32(*(*int8)(unsafe.Add(mBase, uint32(v76))))
	v84 = F_scanner_isspace(m, v83)
	mBase = m.M
	if v84 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v115 = v109
	v118 = int32(0)
	v119 = v73 + v74
	goto L7
L17:
	;
	v90 = v79
	goto L20
L18:
	;
	goto L19
L19:
	;
	v106 = int32(1)
	v109 = v76 + v106
	if v106 < v79 {
		v76 = v109
		v79 = v79 - v106
		v80 = v80 + v106
		goto L15
	} else {
		goto L24
	}
L20:
	;
	v97 = int32(*(*int8)(unsafe.Add(mBase, uint32(v76+v90-int32(1)))))
	v98 = F_scanner_isspace(m, v97)
	mBase = m.M
	if v98 == int32(0) {
		v115 = v76
		v118 = v90
		v119 = v80
		goto L7
	} else {
		goto L22
	}
L21:
	;
	v115 = v76
	v118 = int32(0)
	v119 = v80
	goto L7
L22:
	;
	v101 = int32(1)
	if v101 < v90 {
		v90 = v90 - v101
		goto L20
	} else {
		goto L23
	}
L23:
	;
	goto L21
L24:
	;
	goto L16
L25:
	;
	return
L26:
	;
	v149 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_store[2]))
	v152 = int32(0)
	v154 = F_hash_search(m, v149, v33+int32(120), v152, v152)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L25
	} else {
		goto L29
	}
L27:
	;
	v1420 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_store[1]))
	v1421 = *(*int32)(unsafe.Add(mBase, uint32(v1420)))
	F_LWLockRelease(m, v1421)
	mBase = m.M
	v1423 = m.ExcPending
	if v1423 != 0 {
		goto L25
	} else {
		goto L242
	}
L28:
	;
	if l10 != 0 {
		v1404 = v1111
		goto L27
	} else {
		goto L208
	}
L29:
	;
	if v154 != 0 {
		v1099 = v154
		v1111 = v14
		goto L28
	} else {
		goto L30
	}
L30:
	;
	if l10 != 0 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v157 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_store[1]))
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v157)))
	F_LWLockRelease(m, v158)
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L25
	} else {
		goto L34
	}
L32:
	;
	v505 = v14
	goto L33
L33:
	;
	if v505 != 0 {
		goto L85
	} else {
		goto L86
	}
L34:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v33)+148))
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v33)+152))
	v163 = *(*int32)(unsafe.Add(mBase, uint32(l10)+16))
	if int32(2) <= v163 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(l10)+8))
	F_pg_qsort(m, v166, v163, int32(12), int32(_a_F_pgss_store_0))
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L25
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(l10)+8))
	v177 = F_scanner_init(m, v115, v33+int32(160), int32(_a_F_pgss_store_1), int32(_a_F_pgss_store_2))
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L25
	} else {
		goto L39
	}
L38:
	;
	goto L37
L39:
	;
	v179 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v33)+180)) = uint8(v179)
	v181 = *(*int32)(unsafe.Add(mBase, uint32(l10)+16))
	if v181 <= v179 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	F_scanner_finish(m, v177)
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L25
	} else {
		goto L60
	}
L41:
	;
	v184 = int32(0)
	v200 = v181
	goto L42
L42:
	;
	if v184 == int32(0) {
		goto L45
	} else {
		goto L46
	}
L43:
	;
	goto L40
L44:
	;
	v317 = v184 + int32(1)
	if v317 < v302 {
		v184 = v317
		v200 = v302
		goto L42
	} else {
		goto L59
	}
L45:
	;
	v229 = v171 + v184*int32(12)
	v230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v229)+8)))
	if v230 != 0 {
		v302 = v200
		goto L44
	} else {
		goto L48
	}
L46:
	;
	v216 = int32(12)
	v218 = v171 + v184*v216
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v218)))
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v218-v216)))
	if v219 != v222 {
		goto L45
	} else {
		goto L47
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v218)+4)) = int32(-1)
	v302 = v200
	goto L44
L48:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v229)))
	v232 = v231 - v162
	goto L49
L49:
	;
	v264 = v33 + int32(220)
	v266 = v33 + int32(156)
	v267 = F_core_yylex(m, v264, v266, v177)
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L25
	} else {
		goto L51
	}
L50:
	;
	v274 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115+v232))))
	if v274 == int32(45) {
		goto L54
	} else {
		goto L55
	}
L51:
	;
	if v267 == int32(0) {
		goto L40
	} else {
		goto L52
	}
L52:
	;
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v33)+156))
	if v271 < v232 {
		goto L49
	} else {
		goto L53
	}
L53:
	;
	goto L50
L54:
	;
	v277 = F_core_yylex(m, v264, v266, v177)
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L25
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v33)+160))
	v283 = F_strlen(m, v281+v232)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v229)+4)) = v283
	v285 = *(*int32)(unsafe.Add(mBase, uint32(l10)+16))
	v302 = v285
	goto L44
L57:
	;
	if v277 == int32(0) {
		goto L40
	} else {
		goto L58
	}
L58:
	;
	goto L56
L59:
	;
	goto L43
L60:
	;
	v351 = *(*int32)(unsafe.Add(mBase, uint32(l10)+16))
	v357 = F_palloc(m, v161+v351*int32(10)+int32(1))
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L25
	} else {
		goto L61
	}
L61:
	;
	v359 = *(*int32)(unsafe.Add(mBase, uint32(l10)+16))
	if int32(0) < v359 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v366 = int32(0)
	v377 = v14
	v379 = v359
	v380 = v14
	v381 = v14
	v383 = v14
	v387 = v14
	goto L65
L63:
	;
	v462 = v14
	v465 = v14
	goto L64
L64:
	;
	v475 = v161 - v465
	if v475 != 0 {
		goto L81
	} else {
		goto L82
	}
L65:
	;
	v394 = v366 * int32(12)
	v395 = *(*int32)(unsafe.Add(mBase, uint32(l10)+8))
	v396 = v394 + v395
	v397 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v396)+9)))
	if v397 == int32(1) {
		goto L68
	} else {
		goto L69
	}
L66:
	;
	v462 = v438
	v465 = v440
	goto L64
L67:
	;
	v443 = v366 + int32(1)
	if v443 < v437 {
		v366 = v443
		v377 = v436
		v379 = v437
		v380 = v438
		v381 = v439
		v383 = v440
		v387 = v441
		goto L65
	} else {
		goto L80
	}
L68:
	;
	v400 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l10)+24)))
	if v400 != int32(1) {
		v436 = v377
		v437 = v379
		v438 = v380
		v439 = v381
		v440 = v383
		v441 = v387
		goto L67
	} else {
		goto L71
	}
L69:
	;
	goto L70
L70:
	;
	v403 = *(*int32)(unsafe.Add(mBase, uint32(v396)+4))
	if v403 < int32(0) {
		v436 = v377
		v437 = v379
		v438 = v380
		v439 = v381
		v440 = v383
		v441 = v387
		goto L67
	} else {
		goto L72
	}
L71:
	;
	goto L70
L72:
	;
	v407 = *(*int32)(unsafe.Add(mBase, uint32(v396)))
	v408 = v407 - v162
	v409 = v408 - (v377 + v387)
	if v409 != 0 {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	base.MemoryCopy(m, v357+v380, v115+v383, v409)
	goto L75
L74:
	;
	goto L75
L75:
	;
	v413 = *(*int32)(unsafe.Add(mBase, uint32(l10)+8))
	v415 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v413+v394)+8)))
	v417 = v381 + int32(1)
	v418 = *(*int32)(unsafe.Add(mBase, uint32(l10)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v33)+96)) = v417 + v418
	if v415 != 0 {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v423 = int32(_a_F_pgss_store_3)
	goto L78
L77:
	;
	v423 = int32(_a_F_pgss_store_4)
	goto L78
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+100)) = v423
	v425 = v409 + v380
	v430 = F_pg_sprintf(m, v357+v425, int32(_a_F_pgss_store_5), v33+int32(96))
	mBase = m.M
	v431 = m.ExcPending
	if v431 != 0 {
		goto L25
	} else {
		goto L79
	}
L79:
	;
	v434 = *(*int32)(unsafe.Add(mBase, uint32(l10)+16))
	v436 = v403
	v437 = v434
	v438 = v430 + v425
	v439 = v417
	v440 = v403 + v408
	v441 = v408
	goto L67
L80:
	;
	goto L66
L81:
	;
	base.MemoryCopy(m, v357+v462, v115+v465, v475)
	goto L83
L82:
	;
	goto L83
L83:
	;
	v479 = v475 + v462
	v481 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v357+v479))) = uint8(v481)
	*(*int32)(unsafe.Add(mBase, uint32(v33)+148)) = v479
	v485 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_store[1]))
	v486 = *(*int32)(unsafe.Add(mBase, uint32(v485)))
	v488 = F_LWLockAcquire(m, v486, int32(1))
	mBase = m.M
	v489 = m.ExcPending
	if v489 != 0 {
		goto L25
	} else {
		goto L84
	}
L84:
	;
	v505 = v357
	goto L33
L85:
	;
	v520 = v505
	goto L87
L86:
	;
	v520 = v115
	goto L87
L87:
	;
	v521 = *(*int32)(unsafe.Add(mBase, uint32(v33)+148))
	v526 = F_qtext_store(m, v520, v521, v33+int32(116), v33+int32(112))
	mBase = m.M
	v527 = m.ExcPending
	if v527 != 0 {
		goto L25
	} else {
		goto L88
	}
L88:
	;
	v529 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_store[1]))
	v530 = *(*int32)(unsafe.Add(mBase, uint32(v529)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v529)+20)) = int32(1)
	if v530 != 0 {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	v534 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_store[1]))
	F_s_lock(m, v534+int32(20), int32(_a_F_pgss_store_6), int32(2428), int32(_a_F_pgss_store_7))
	mBase = m.M
	v541 = m.ExcPending
	if v541 != 0 {
		goto L25
	} else {
		goto L92
	}
L90:
	;
	goto L91
L91:
	;
	v543 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_store[1]))
	v544 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v543)+20)) = v544
	v547 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v543)+24)))
	v549 = int64(*(*int32)(unsafe.Add(mBase, _c_F_pgss_store[6])))
	if base.Ui64(v549<<(uint(int64(9))%64)) <= base.Ui64(v547) {
		goto L93
	} else {
		goto L94
	}
L92:
	;
	goto L91
L93:
	;
	v553 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v543)+16)))
	v558 = base.B2i32(base.Ui64(v549*v553<<(uint(int64(1))%64)) <= base.Ui64(v547))
	goto L95
L94:
	;
	v558 = v544
	goto L95
L95:
	;
	v559 = *(*int32)(unsafe.Add(mBase, uint32(v543)))
	F_LWLockRelease(m, v559)
	mBase = m.M
	v561 = m.ExcPending
	if v561 != 0 {
		goto L25
	} else {
		goto L96
	}
L96:
	;
	v563 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_store[1]))
	v564 = *(*int32)(unsafe.Add(mBase, uint32(v563)))
	v566 = F_LWLockAcquire(m, v564, int32(0))
	mBase = m.M
	v567 = m.ExcPending
	if v567 != 0 {
		goto L25
	} else {
		goto L97
	}
L97:
	;
	if v526 != 0 {
		goto L99
	} else {
		goto L100
	}
L98:
	;
	v583 = *(*int32)(unsafe.Add(mBase, uint32(v33)+116))
	v584 = *(*int32)(unsafe.Add(mBase, uint32(v33)+148))
	v587 = F_entry_alloc(m, v33+int32(120), v583, v584, v39, base.B2i32(l10 != int32(0)))
	mBase = m.M
	v588 = m.ExcPending
	if v588 != 0 {
		goto L25
	} else {
		goto L105
	}
L99:
	;
	v569 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_store[1]))
	v570 = *(*int32)(unsafe.Add(mBase, uint32(v569)+32))
	v571 = *(*int32)(unsafe.Add(mBase, uint32(v33)+112))
	if v570 == v571 {
		goto L98
	} else {
		goto L102
	}
L100:
	;
	goto L101
L101:
	;
	v573 = *(*int32)(unsafe.Add(mBase, uint32(v33)+148))
	v577 = F_qtext_store(m, v520, v573, v33+int32(116), int32(0))
	mBase = m.M
	v578 = m.ExcPending
	if v578 != 0 {
		goto L25
	} else {
		goto L103
	}
L102:
	;
	goto L101
L103:
	;
	if v577 == int32(0) {
		v1404 = v505
		goto L27
	} else {
		goto L104
	}
L104:
	;
	goto L98
L105:
	;
	if v558 == int32(0) {
		v1099 = v587
		v1111 = v505
		goto L28
	} else {
		goto L106
	}
L106:
	;
	v592 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_store[1]))
	v593 = *(*int32)(unsafe.Add(mBase, uint32(v592)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v592)+20)) = int32(1)
	if v593 != 0 {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	v597 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_store[1]))
	F_s_lock(m, v597+int32(20), int32(_a_F_pgss_store_6), int32(2428), int32(_a_F_pgss_store_7))
	mBase = m.M
	v604 = m.ExcPending
	if v604 != 0 {
		goto L25
	} else {
		goto L110
	}
L108:
	;
	goto L109
L109:
	;
	v606 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_store[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v606)+20)) = int32(0)
	v609 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v606)+24)))
	v611 = int64(*(*int32)(unsafe.Add(mBase, _c_F_pgss_store[6])))
	if base.Ui64(v609) < base.Ui64(v611<<(uint(int64(9))%64)) {
		v1099 = v587
		v1111 = v505
		goto L28
	} else {
		goto L111
	}
L110:
	;
	goto L109
L111:
	;
	v615 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v606)+16)))
	if base.Ui64(v609) < base.Ui64(v611*v615<<(uint(int64(1))%64)) {
		v1099 = v587
		v1111 = v505
		goto L28
	} else {
		goto L112
	}
L112:
	;
	v622 = F_qtext_load_file(m, v33+int32(220))
	mBase = m.M
	v623 = m.ExcPending
	if v623 != 0 {
		goto L25
	} else {
		goto L116
	}
L113:
	;
	v1089 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_store[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v1089)+20)) = int32(0)
	v1092 = *(*int32)(unsafe.Add(mBase, uint32(v1089)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v1089)+32)) = v1092 + int32(1)
	v1099 = v587
	v1111 = v505
	goto L28
L114:
	;
	v1051 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_store[1]))
	F_s_lock(m, v1051+int32(20), int32(_a_F_pgss_store_6), v1049, int32(_a_F_pgss_store_8))
	mBase = m.M
	v1057 = m.ExcPending
	if v1057 != 0 {
		goto L25
	} else {
		goto L207
	}
L115:
	;
	F_emscripten_builtin_free(m, v622)
	mBase = m.M
	v906 = v33 + int32(160)
	v908 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_store[2]))
	F_hash_seq_init(m, v906, v908)
	mBase = m.M
	v910 = m.ExcPending
	if v910 != 0 {
		goto L25
	} else {
		goto L186
	}
L116:
	;
	if v622 == int32(0) {
		goto L115
	} else {
		goto L117
	}
L117:
	;
	v628 = F_AllocateFile(m, int32(_a_F_pgss_store_9), int32(_a_F_pgss_store_10))
	mBase = m.M
	v629 = m.ExcPending
	if v629 != 0 {
		goto L25
	} else {
		goto L118
	}
L118:
	;
	if v628 == int32(0) {
		goto L119
	} else {
		goto L120
	}
L119:
	;
	v634 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v635 = m.ExcPending
	if v635 != 0 {
		goto L25
	} else {
		goto L122
	}
L120:
	;
	goto L121
L121:
	;
	v653 = v33 + int32(160)
	v655 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_store[2]))
	F_hash_seq_init(m, v653, v655)
	mBase = m.M
	v657 = m.ExcPending
	if v657 != 0 {
		goto L25
	} else {
		goto L127
	}
L122:
	;
	if v634 == int32(0) {
		goto L115
	} else {
		goto L123
	}
L123:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v639 = m.ExcPending
	if v639 != 0 {
		goto L25
	} else {
		goto L124
	}
L124:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+16)) = int32(_a_F_pgss_store_9)
	F_errmsg(m, int32(_a_F_pgss_store_11), v33+int32(16))
	mBase = m.M
	v646 = m.ExcPending
	if v646 != 0 {
		goto L25
	} else {
		goto L125
	}
L125:
	;
	F_errfinish(m, int32(_a_F_pgss_store_6), int32(2513), int32(_a_F_pgss_store_8))
	mBase = m.M
	v651 = m.ExcPending
	if v651 != 0 {
		goto L25
	} else {
		goto L126
	}
L126:
	;
	goto L115
L127:
	;
	v658 = F_hash_seq_search(m, v653)
	mBase = m.M
	v659 = m.ExcPending
	if v659 != 0 {
		goto L25
	} else {
		goto L129
	}
L128:
	;
	v785 = *(*int32)(unsafe.Add(mBase, uint32(v628)+60))
	if v785 < int32(0) {
		goto L158
	} else {
		goto L159
	}
L129:
	;
	if v658 == int32(0) {
		goto L130
	} else {
		goto L131
	}
L130:
	;
	v662 = int32(0)
	v769 = v662
	v773 = v662
	goto L128
L131:
	;
	goto L132
L132:
	;
	v664 = *(*int32)(unsafe.Add(mBase, uint32(v33)+220))
	v665 = int32(0)
	v669 = v658
	v681 = v665
	v685 = v665
	goto L133
L133:
	;
	v697 = *(*int32)(unsafe.Add(mBase, uint32(v669)+396))
	if v697 < int32(0) {
		goto L137
	} else {
		goto L138
	}
L134:
	;
	v769 = v746
	v773 = v748
	goto L128
L135:
	;
	v753 = F_hash_seq_search(m, v33+int32(160))
	mBase = m.M
	v754 = m.ExcPending
	if v754 != 0 {
		goto L25
	} else {
		goto L154
	}
L136:
	;
	v712 = int32(1)
	v714 = v697 + v712
	v715 = F_fwrite(m, v622+v700, v712, v714, v628)
	mBase = m.M
	v716 = m.ExcPending
	if v716 != 0 {
		goto L25
	} else {
		goto L141
	}
L137:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v669)+392)) = int64(-4294967296)
	v746 = v681
	v748 = v685
	goto L135
L138:
	;
	v700 = *(*int32)(unsafe.Add(mBase, uint32(v669)+392))
	v701 = v700 + v697
	if base.Ui32(v664) <= base.Ui32(v701) {
		goto L137
	} else {
		goto L139
	}
L139:
	;
	v704 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v622+v701))))
	if v704 == int32(0) {
		goto L136
	} else {
		goto L140
	}
L140:
	;
	goto L137
L141:
	;
	if v715 != v714 {
		goto L142
	} else {
		goto L143
	}
L142:
	;
	v720 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v721 = m.ExcPending
	if v721 != 0 {
		goto L25
	} else {
		goto L145
	}
L143:
	;
	goto L144
L144:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v669)+392)) = v685
	v746 = v681 + int32(1)
	v748 = v714 + v685
	goto L135
L145:
	;
	if v720 != 0 {
		goto L146
	} else {
		goto L147
	}
L146:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v723 = m.ExcPending
	if v723 != 0 {
		goto L25
	} else {
		goto L149
	}
L147:
	;
	goto L148
L148:
	;
	F_hash_seq_term(m, v33+int32(160))
	mBase = m.M
	v739 = m.ExcPending
	if v739 != 0 {
		goto L25
	} else {
		goto L152
	}
L149:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+80)) = int32(_a_F_pgss_store_9)
	F_errmsg(m, int32(_a_F_pgss_store_11), v33+int32(80))
	mBase = m.M
	v730 = m.ExcPending
	if v730 != 0 {
		goto L25
	} else {
		goto L150
	}
L150:
	;
	F_errfinish(m, int32(_a_F_pgss_store_6), int32(2543), int32(_a_F_pgss_store_8))
	mBase = m.M
	v735 = m.ExcPending
	if v735 != 0 {
		goto L25
	} else {
		goto L151
	}
L151:
	;
	goto L148
L152:
	;
	v740 = F_FreeFile(m, v628)
	mBase = m.M
	v741 = m.ExcPending
	if v741 != 0 {
		goto L25
	} else {
		goto L153
	}
L153:
	;
	goto L115
L154:
	;
	if v753 != 0 {
		v669 = v753
		v681 = v746
		v685 = v748
		goto L133
	} else {
		goto L155
	}
L155:
	;
	goto L134
L156:
	;
	v817 = F_FreeFile(m, v628)
	mBase = m.M
	v818 = m.ExcPending
	if v818 != 0 {
		goto L25
	} else {
		goto L167
	}
L157:
	;
	v794 = F_ftruncate(m, v792, base.I64_extend_i32_u(v773))
	mBase = m.M
	if v794 == int32(0) {
		goto L156
	} else {
		goto L161
	}
L158:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_pgss_store[7])) = int32(8)
	v792 = int32(-1)
	goto L160
L159:
	;
	v792 = v785
	goto L160
L160:
	;
	goto L157
L161:
	;
	v799 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v800 = m.ExcPending
	if v800 != 0 {
		goto L25
	} else {
		goto L162
	}
L162:
	;
	if v799 == int32(0) {
		goto L156
	} else {
		goto L163
	}
L163:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v804 = m.ExcPending
	if v804 != 0 {
		goto L25
	} else {
		goto L164
	}
L164:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+64)) = int32(_a_F_pgss_store_9)
	F_errmsg(m, int32(_a_F_pgss_store_12), v33-int32(-64))
	mBase = m.M
	v811 = m.ExcPending
	if v811 != 0 {
		goto L25
	} else {
		goto L165
	}
L165:
	;
	F_errfinish(m, int32(_a_F_pgss_store_6), int32(2561), int32(_a_F_pgss_store_8))
	mBase = m.M
	v816 = m.ExcPending
	if v816 != 0 {
		goto L25
	} else {
		goto L166
	}
L166:
	;
	goto L156
L167:
	;
	if v817 != 0 {
		goto L168
	} else {
		goto L169
	}
L168:
	;
	v821 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v822 = m.ExcPending
	if v822 != 0 {
		goto L25
	} else {
		goto L171
	}
L169:
	;
	goto L170
L170:
	;
	v841 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v842 = m.ExcPending
	if v842 != 0 {
		goto L25
	} else {
		goto L176
	}
L171:
	;
	if v821 == int32(0) {
		goto L115
	} else {
		goto L172
	}
L172:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v826 = m.ExcPending
	if v826 != 0 {
		goto L25
	} else {
		goto L173
	}
L173:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+48)) = int32(_a_F_pgss_store_9)
	F_errmsg(m, int32(_a_F_pgss_store_11), v33+int32(48))
	mBase = m.M
	v833 = m.ExcPending
	if v833 != 0 {
		goto L25
	} else {
		goto L174
	}
L174:
	;
	F_errfinish(m, int32(_a_F_pgss_store_6), int32(2568), int32(_a_F_pgss_store_8))
	mBase = m.M
	v838 = m.ExcPending
	if v838 != 0 {
		goto L25
	} else {
		goto L175
	}
L175:
	;
	goto L115
L176:
	;
	if v841 != 0 {
		goto L177
	} else {
		goto L178
	}
L177:
	;
	v844 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_store[1]))
	v845 = *(*int32)(unsafe.Add(mBase, uint32(v844)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v33)+32)) = v845
	*(*int32)(unsafe.Add(mBase, uint32(v33)+36)) = v773
	F_errmsg_internal(m, int32(_a_F_pgss_store_13), v33+int32(32))
	mBase = m.M
	v852 = m.ExcPending
	if v852 != 0 {
		goto L25
	} else {
		goto L180
	}
L178:
	;
	goto L179
L179:
	;
	v859 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_store[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v859)+24)) = v773
	if v769 <= int32(0) {
		goto L182
	} else {
		goto L183
	}
L180:
	;
	F_errfinish(m, int32(_a_F_pgss_store_6), int32(2574), int32(_a_F_pgss_store_8))
	mBase = m.M
	v857 = m.ExcPending
	if v857 != 0 {
		goto L25
	} else {
		goto L181
	}
L181:
	;
	goto L179
L182:
	;
	v865 = int32(1024)
	goto L184
L183:
	;
	v864 = base.I32_div_u_s(v773, v769)
	v865 = v864
	goto L184
L184:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v859)+16)) = v865
	F_emscripten_builtin_free(m, v622)
	mBase = m.M
	v868 = *(*int32)(unsafe.Add(mBase, uint32(v859)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v859)+20)) = int32(1)
	if v868 == int32(0) {
		goto L113
	} else {
		goto L185
	}
L185:
	;
	v1049 = int32(2597)
	goto L114
L186:
	;
	v911 = F_hash_seq_search(m, v906)
	mBase = m.M
	v912 = m.ExcPending
	if v912 != 0 {
		goto L25
	} else {
		goto L187
	}
L187:
	;
	if v911 != 0 {
		goto L188
	} else {
		goto L189
	}
L188:
	;
	v915 = v911
	goto L191
L189:
	;
	goto L190
L190:
	;
	v979 = int32(_a_F_pgss_store_9)
	v980 = F_unlink(m, v979)
	mBase = m.M
	v983 = F_AllocateFile(m, v979, int32(_a_F_pgss_store_10))
	mBase = m.M
	v984 = m.ExcPending
	if v984 != 0 {
		goto L25
	} else {
		goto L196
	}
L191:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v915)+392)) = int64(-4294967296)
	v947 = F_hash_seq_search(m, v33+int32(160))
	mBase = m.M
	v948 = m.ExcPending
	if v948 != 0 {
		goto L25
	} else {
		goto L193
	}
L192:
	;
	goto L190
L193:
	;
	if v947 != 0 {
		v915 = v947
		goto L191
	} else {
		goto L194
	}
L194:
	;
	goto L192
L195:
	;
	v1008 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_store[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v1008)+16)) = int32(1024)
	v1011 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1008)+24)) = v1011
	v1013 = *(*int32)(unsafe.Add(mBase, uint32(v1008)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v1008)+20)) = int32(1)
	if v1013 == v1011 {
		goto L113
	} else {
		goto L206
	}
L196:
	;
	if v983 == int32(0) {
		goto L197
	} else {
		goto L198
	}
L197:
	;
	v989 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v990 = m.ExcPending
	if v990 != 0 {
		goto L25
	} else {
		goto L200
	}
L198:
	;
	goto L199
L199:
	;
	v1005 = F_FreeFile(m, v983)
	mBase = m.M
	v1006 = m.ExcPending
	if v1006 != 0 {
		goto L25
	} else {
		goto L205
	}
L200:
	;
	if v989 == int32(0) {
		goto L195
	} else {
		goto L201
	}
L201:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v994 = m.ExcPending
	if v994 != 0 {
		goto L25
	} else {
		goto L202
	}
L202:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33))) = int32(_a_F_pgss_store_9)
	F_errmsg(m, int32(_a_F_pgss_store_14), v33)
	mBase = m.M
	v999 = m.ExcPending
	if v999 != 0 {
		goto L25
	} else {
		goto L203
	}
L203:
	;
	F_errfinish(m, int32(_a_F_pgss_store_6), int32(2627), int32(_a_F_pgss_store_8))
	mBase = m.M
	v1004 = m.ExcPending
	if v1004 != 0 {
		goto L25
	} else {
		goto L204
	}
L204:
	;
	goto L195
L205:
	;
	goto L195
L206:
	;
	v1049 = int32(2648)
	goto L114
L207:
	;
	goto L113
L208:
	;
	v1126 = *(*int32)(unsafe.Add(mBase, uint32(v1099)+424))
	*(*int32)(unsafe.Add(mBase, uint32(v1099)+424)) = int32(1)
	if v1126 != 0 {
		goto L209
	} else {
		goto L210
	}
L209:
	;
	F_s_lock(m, v1099+int32(424), int32(_a_F_pgss_store_6), int32(1404), int32(_a_F_pgss_store_15))
	mBase = m.M
	v1135 = m.ExcPending
	if v1135 != 0 {
		goto L25
	} else {
		goto L212
	}
L210:
	;
	goto L211
L211:
	;
	v1136 = *(*int64)(unsafe.Add(mBase, uint32(v1099)+24))
	v1138 = *(*int64)(unsafe.Add(mBase, uint32(v1099)+32))
	if v1136 == int64(0)-v1138 {
		goto L213
	} else {
		goto L214
	}
L212:
	;
	goto L211
L213:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1099)+256)) = int64(4607182418800017408)
	goto L215
L214:
	;
	goto L215
L215:
	;
	v1144 = l4 << (uint(int32(3)) % 32)
	v1147 = v1144 + (v1099 + int32(24))
	v1148 = *(*int64)(unsafe.Add(mBase, uint32(v1147)))
	v1150 = v1148 + int64(1)
	*(*int64)(unsafe.Add(mBase, uint32(v1147))) = v1150
	v1152 = v1144 + v1099
	v1154 = v1152 + int32(40)
	v1155 = *(*float64)(unsafe.Add(mBase, uint32(v1154)))
	*(*float64)(unsafe.Add(mBase, uint32(v1154))) = base.F64_add(l5, v1155)
	if v1148 == int64(0) {
		goto L217
	} else {
		goto L218
	}
L216:
	;
	v1210 = *(*int64)(unsafe.Add(mBase, uint32(v1099)+120))
	*(*int64)(unsafe.Add(mBase, uint32(v1099)+120)) = v1210 + l6
	v1213 = *(*int64)(unsafe.Add(mBase, uint32(v1099)+128))
	v1214 = *(*int64)(unsafe.Add(mBase, uint32(l7)))
	*(*int64)(unsafe.Add(mBase, uint32(v1099)+128)) = v1213 + v1214
	v1217 = *(*int64)(unsafe.Add(mBase, uint32(v1099)+136))
	v1218 = *(*int64)(unsafe.Add(mBase, uint32(l7)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v1099)+136)) = v1217 + v1218
	v1221 = *(*int64)(unsafe.Add(mBase, uint32(v1099)+144))
	v1222 = *(*int64)(unsafe.Add(mBase, uint32(l7)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v1099)+144)) = v1221 + v1222
	v1225 = *(*int64)(unsafe.Add(mBase, uint32(v1099)+152))
	v1226 = *(*int64)(unsafe.Add(mBase, uint32(l7)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v1099)+152)) = v1225 + v1226
	v1229 = *(*int64)(unsafe.Add(mBase, uint32(v1099)+160))
	v1230 = *(*int64)(unsafe.Add(mBase, uint32(l7)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v1099)+160)) = v1229 + v1230
	v1233 = *(*int64)(unsafe.Add(mBase, uint32(v1099)+168))
	v1234 = *(*int64)(unsafe.Add(mBase, uint32(l7)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v1099)+168)) = v1233 + v1234
	v1237 = *(*int64)(unsafe.Add(mBase, uint32(v1099)+176))
	v1238 = *(*int64)(unsafe.Add(mBase, uint32(l7)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v1099)+176)) = v1237 + v1238
	v1241 = *(*int64)(unsafe.Add(mBase, uint32(v1099)+184))
	v1242 = *(*int64)(unsafe.Add(mBase, uint32(l7)+56))
	*(*int64)(unsafe.Add(mBase, uint32(v1099)+184)) = v1241 + v1242
	v1245 = *(*int64)(unsafe.Add(mBase, uint32(v1099)+192))
	v1246 = *(*int64)(unsafe.Add(mBase, uint32(l7)+64))
	*(*int64)(unsafe.Add(mBase, uint32(v1099)+192)) = v1245 + v1246
	v1249 = *(*int64)(unsafe.Add(mBase, uint32(v1099)+200))
	v1250 = *(*int64)(unsafe.Add(mBase, uint32(l7)+72))
	*(*int64)(unsafe.Add(mBase, uint32(v1099)+200)) = v1249 + v1250
	v1253 = *(*float64)(unsafe.Add(mBase, uint32(v1099)+208))
	v1254 = *(*int64)(unsafe.Add(mBase, uint32(l7)+80))
	v1256 = float64(1e+06)
	*(*float64)(unsafe.Add(mBase, uint32(v1099)+208)) = base.F64_add(v1253, base.F64_div(base.F64_convert_i64_s(v1254), v1256))
	v1260 = *(*float64)(unsafe.Add(mBase, uint32(v1099)+216))
	v1261 = *(*int64)(unsafe.Add(mBase, uint32(l7)+88))
	*(*float64)(unsafe.Add(mBase, uint32(v1099)+216)) = base.F64_add(v1260, base.F64_div(base.F64_convert_i64_s(v1261), v1256))
	v1267 = *(*float64)(unsafe.Add(mBase, uint32(v1099)+224))
	v1268 = *(*int64)(unsafe.Add(mBase, uint32(l7)+96))
	*(*float64)(unsafe.Add(mBase, uint32(v1099)+224)) = base.F64_add(v1267, base.F64_div(base.F64_convert_i64_s(v1268), v1256))
	v1274 = *(*float64)(unsafe.Add(mBase, uint32(v1099)+232))
	v1275 = *(*int64)(unsafe.Add(mBase, uint32(l7)+104))
	*(*float64)(unsafe.Add(mBase, uint32(v1099)+232)) = base.F64_add(v1274, base.F64_div(base.F64_convert_i64_s(v1275), v1256))
	v1281 = *(*float64)(unsafe.Add(mBase, uint32(v1099)+240))
	v1282 = *(*int64)(unsafe.Add(mBase, uint32(l7)+112))
	*(*float64)(unsafe.Add(mBase, uint32(v1099)+240)) = base.F64_add(v1281, base.F64_div(base.F64_convert_i64_s(v1282), v1256))
	v1288 = *(*float64)(unsafe.Add(mBase, uint32(v1099)+248))
	v1289 = *(*int64)(unsafe.Add(mBase, uint32(l7)+120))
	*(*float64)(unsafe.Add(mBase, uint32(v1099)+248)) = base.F64_add(v1288, base.F64_div(base.F64_convert_i64_s(v1289), v1256))
	v1295 = *(*float64)(unsafe.Add(mBase, uint32(v1099)+256))
	*(*float64)(unsafe.Add(mBase, uint32(v1099)+256)) = base.F64_add(v1295, float64(1))
	v1299 = *(*int64)(unsafe.Add(mBase, uint32(v1099)+264))
	v1300 = *(*int64)(unsafe.Add(mBase, uint32(l8)))
	*(*int64)(unsafe.Add(mBase, uint32(v1099)+264)) = v1299 + v1300
	v1303 = *(*int64)(unsafe.Add(mBase, uint32(v1099)+272))
	v1304 = *(*int64)(unsafe.Add(mBase, uint32(l8)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v1099)+272)) = v1303 + v1304
	v1307 = *(*int64)(unsafe.Add(mBase, uint32(v1099)+280))
	v1308 = *(*int64)(unsafe.Add(mBase, uint32(l8)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v1099)+280)) = v1307 + v1308
	v1311 = *(*int64)(unsafe.Add(mBase, uint32(v1099)+288))
	v1312 = *(*int64)(unsafe.Add(mBase, uint32(l8)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v1099)+288)) = v1311 + v1312
	if l9 != 0 {
		goto L227
	} else {
		goto L228
	}
L217:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v1152)+88)) = l5
	*(*float64)(unsafe.Add(mBase, uint32(v1152)+72)) = l5
	*(*float64)(unsafe.Add(mBase, uint32(v1152)+56)) = l5
	goto L216
L218:
	;
	goto L219
L219:
	;
	v1164 = v1152 + int32(88)
	v1165 = *(*float64)(unsafe.Add(mBase, uint32(v1164)))
	v1166 = base.F64_sub(l5, v1165)
	v1169 = base.F64_add(v1165, base.F64_div(v1166, base.F64_convert_i64_s(v1150)))
	*(*float64)(unsafe.Add(mBase, uint32(v1164))) = v1169
	v1172 = v1152 + int32(104)
	v1175 = *(*float64)(unsafe.Add(mBase, uint32(v1172)))
	*(*float64)(unsafe.Add(mBase, uint32(v1172))) = base.F64_add(base.F64_mul(v1166, base.F64_sub(l5, v1169)), v1175)
	v1179 = v1152 + int32(56)
	v1180 = *(*float64)(unsafe.Add(mBase, uint32(v1179)))
	if base.F64_ne(v1180, float64(0)) != 0 {
		goto L220
	} else {
		goto L221
	}
L220:
	;
	if base.F64_lt(l5, v1180) != 0 {
		goto L223
	} else {
		goto L224
	}
L221:
	;
	v1187 = v1099 + l4<<(uint(int32(3))%32) + int32(72)
	v1188 = *(*float64)(unsafe.Add(mBase, uint32(v1187)))
	if base.F64_ne(v1188, float64(0)) != 0 {
		goto L220
	} else {
		goto L222
	}
L222:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v1179))) = l5
	*(*float64)(unsafe.Add(mBase, uint32(v1187))) = l5
	goto L216
L223:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v1179))) = l5
	goto L225
L224:
	;
	goto L225
L225:
	;
	v1200 = v1099 + l4<<(uint(int32(3))%32) + int32(72)
	v1201 = *(*float64)(unsafe.Add(mBase, uint32(v1200)))
	if base.F64_lt(v1201, l5) == int32(0) {
		goto L216
	} else {
		goto L226
	}
L226:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v1200))) = l5
	goto L216
L227:
	;
	v1315 = *(*int64)(unsafe.Add(mBase, uint32(v1099)+296))
	v1316 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l9))))
	*(*int64)(unsafe.Add(mBase, uint32(v1099)+296)) = v1315 + v1316
	v1319 = *(*float64)(unsafe.Add(mBase, uint32(v1099)+304))
	v1320 = *(*int64)(unsafe.Add(mBase, uint32(l9)+8))
	v1322 = float64(1e+06)
	*(*float64)(unsafe.Add(mBase, uint32(v1099)+304)) = base.F64_add(v1319, base.F64_div(base.F64_convert_i64_s(v1320), v1322))
	v1326 = *(*int64)(unsafe.Add(mBase, uint32(l9)+16))
	v1329 = base.F64_div(base.F64_convert_i64_s(v1326), v1322)
	if base.F64_ne(v1329, float64(0)) != 0 {
		goto L230
	} else {
		goto L231
	}
L228:
	;
	goto L229
L229:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1099)+424)) = int32(0)
	v1381 = *(*int64)(unsafe.Add(mBase, uint32(v1099)+376))
	*(*int64)(unsafe.Add(mBase, uint32(v1099)+376)) = v1381 + base.I64_extend_i32_s(l11)
	v1385 = *(*int64)(unsafe.Add(mBase, uint32(v1099)+384))
	*(*int64)(unsafe.Add(mBase, uint32(v1099)+384)) = v1385 + base.I64_extend_i32_s(l12)
	v1404 = v1111
	goto L27
L230:
	;
	v1332 = *(*int64)(unsafe.Add(mBase, uint32(v1099)+328))
	*(*int64)(unsafe.Add(mBase, uint32(v1099)+328)) = v1332 + int64(1)
	goto L232
L231:
	;
	goto L232
L232:
	;
	v1336 = *(*float64)(unsafe.Add(mBase, uint32(v1099)+320))
	*(*float64)(unsafe.Add(mBase, uint32(v1099)+320)) = base.F64_add(v1329, v1336)
	v1339 = *(*int64)(unsafe.Add(mBase, uint32(l9)+24))
	v1342 = base.F64_div(base.F64_convert_i64_s(v1339), float64(1e+06))
	if base.F64_ne(v1342, float64(0)) != 0 {
		goto L233
	} else {
		goto L234
	}
L233:
	;
	v1345 = *(*int64)(unsafe.Add(mBase, uint32(v1099)+312))
	*(*int64)(unsafe.Add(mBase, uint32(v1099)+312)) = v1345 + int64(1)
	goto L235
L234:
	;
	goto L235
L235:
	;
	v1349 = *(*float64)(unsafe.Add(mBase, uint32(v1099)+336))
	*(*float64)(unsafe.Add(mBase, uint32(v1099)+336)) = base.F64_add(v1342, v1349)
	v1352 = *(*int64)(unsafe.Add(mBase, uint32(l9)+32))
	v1355 = base.F64_div(base.F64_convert_i64_s(v1352), float64(1e+06))
	if base.F64_ne(v1355, float64(0)) != 0 {
		goto L236
	} else {
		goto L237
	}
L236:
	;
	v1358 = *(*int64)(unsafe.Add(mBase, uint32(v1099)+344))
	*(*int64)(unsafe.Add(mBase, uint32(v1099)+344)) = v1358 + int64(1)
	goto L238
L237:
	;
	goto L238
L238:
	;
	v1362 = *(*float64)(unsafe.Add(mBase, uint32(v1099)+352))
	*(*float64)(unsafe.Add(mBase, uint32(v1099)+352)) = base.F64_add(v1355, v1362)
	v1365 = *(*int64)(unsafe.Add(mBase, uint32(l9)+40))
	v1368 = base.F64_div(base.F64_convert_i64_s(v1365), float64(1e+06))
	if base.F64_ne(v1368, float64(0)) != 0 {
		goto L239
	} else {
		goto L240
	}
L239:
	;
	v1371 = *(*int64)(unsafe.Add(mBase, uint32(v1099)+360))
	*(*int64)(unsafe.Add(mBase, uint32(v1099)+360)) = v1371 + int64(1)
	goto L241
L240:
	;
	goto L241
L241:
	;
	v1375 = *(*float64)(unsafe.Add(mBase, uint32(v1099)+368))
	*(*float64)(unsafe.Add(mBase, uint32(v1099)+368)) = base.F64_add(v1368, v1375)
	goto L229
L242:
	;
	if v1404 == int32(0) {
		goto L2
	} else {
		goto L243
	}
L243:
	;
	F_pfree(m, v1404)
	mBase = m.M
	v1427 = m.ExcPending
	if v1427 != 0 {
		goto L25
	} else {
		goto L244
	}
L244:
	;
	goto L2
}
