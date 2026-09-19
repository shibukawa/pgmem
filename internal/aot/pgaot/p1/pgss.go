package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"sync/atomic"
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
	var v532 int32
	_ = v532
	var v534 int32
	_ = v534
	var v541 int32
	_ = v541
	var v543 int32
	_ = v543
	var v544 int64
	_ = v544
	var v545 int32
	_ = v545
	var v550 int64
	_ = v550
	var v554 int64
	_ = v554
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v562 int32
	_ = v562
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v574 int32
	_ = v574
	var v578 int32
	_ = v578
	var v579 int32
	_ = v579
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	var v593 int32
	_ = v593
	var v596 int32
	_ = v596
	var v598 int32
	_ = v598
	var v605 int32
	_ = v605
	var v607 int32
	_ = v607
	var v608 int64
	_ = v608
	var v609 int32
	_ = v609
	var v613 int64
	_ = v613
	var v617 int64
	_ = v617
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v630 int32
	_ = v630
	var v631 int32
	_ = v631
	var v636 int32
	_ = v636
	var v637 int32
	_ = v637
	var v641 int32
	_ = v641
	var v648 int32
	_ = v648
	var v653 int32
	_ = v653
	var v655 int32
	_ = v655
	var v657 int32
	_ = v657
	var v659 int32
	_ = v659
	var v660 int32
	_ = v660
	var v661 int32
	_ = v661
	var v664 int32
	_ = v664
	var v666 int32
	_ = v666
	var v667 int32
	_ = v667
	var v671 int32
	_ = v671
	var v683 int32
	_ = v683
	var v687 int32
	_ = v687
	var v699 int32
	_ = v699
	var v702 int32
	_ = v702
	var v703 int32
	_ = v703
	var v706 int32
	_ = v706
	var v714 int32
	_ = v714
	var v716 int32
	_ = v716
	var v717 int32
	_ = v717
	var v718 int32
	_ = v718
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v725 int32
	_ = v725
	var v732 int32
	_ = v732
	var v737 int32
	_ = v737
	var v741 int32
	_ = v741
	var v742 int32
	_ = v742
	var v743 int32
	_ = v743
	var v748 int32
	_ = v748
	var v750 int32
	_ = v750
	var v755 int32
	_ = v755
	var v756 int32
	_ = v756
	var v771 int32
	_ = v771
	var v775 int32
	_ = v775
	var v787 int32
	_ = v787
	var v794 int32
	_ = v794
	var v796 int32
	_ = v796
	var v801 int32
	_ = v801
	var v802 int32
	_ = v802
	var v806 int32
	_ = v806
	var v813 int32
	_ = v813
	var v818 int32
	_ = v818
	var v819 int32
	_ = v819
	var v820 int32
	_ = v820
	var v823 int32
	_ = v823
	var v824 int32
	_ = v824
	var v828 int32
	_ = v828
	var v835 int32
	_ = v835
	var v840 int32
	_ = v840
	var v843 int32
	_ = v843
	var v844 int32
	_ = v844
	var v846 int32
	_ = v846
	var v847 int32
	_ = v847
	var v854 int32
	_ = v854
	var v859 int32
	_ = v859
	var v861 int32
	_ = v861
	var v866 int32
	_ = v866
	var v867 int32
	_ = v867
	var v872 int32
	_ = v872
	var v908 int32
	_ = v908
	var v910 int32
	_ = v910
	var v912 int32
	_ = v912
	var v913 int32
	_ = v913
	var v914 int32
	_ = v914
	var v917 int32
	_ = v917
	var v949 int32
	_ = v949
	var v950 int32
	_ = v950
	var v981 int32
	_ = v981
	var v982 int32
	_ = v982
	var v985 int32
	_ = v985
	var v986 int32
	_ = v986
	var v991 int32
	_ = v991
	var v992 int32
	_ = v992
	var v996 int32
	_ = v996
	var v1001 int32
	_ = v1001
	var v1006 int32
	_ = v1006
	var v1007 int32
	_ = v1007
	var v1008 int32
	_ = v1008
	var v1010 int32
	_ = v1010
	var v1013 int32
	_ = v1013
	var v1017 int32
	_ = v1017
	var v1051 int32
	_ = v1051
	var v1053 int32
	_ = v1053
	var v1059 int32
	_ = v1059
	var v1091 int32
	_ = v1091
	var v1092 int32
	_ = v1092
	var v1096 int32
	_ = v1096
	var v1102 int32
	_ = v1102
	var v1114 int32
	_ = v1114
	var v1131 int32
	_ = v1131
	var v1138 int32
	_ = v1138
	var v1139 int64
	_ = v1139
	var v1141 int64
	_ = v1141
	var v1147 int32
	_ = v1147
	var v1150 int32
	_ = v1150
	var v1151 int64
	_ = v1151
	var v1153 int64
	_ = v1153
	var v1155 int32
	_ = v1155
	var v1157 int32
	_ = v1157
	var v1158 float64
	_ = v1158
	var v1167 int32
	_ = v1167
	var v1168 float64
	_ = v1168
	var v1169 float64
	_ = v1169
	var v1172 float64
	_ = v1172
	var v1175 int32
	_ = v1175
	var v1178 float64
	_ = v1178
	var v1182 int32
	_ = v1182
	var v1183 float64
	_ = v1183
	var v1190 int32
	_ = v1190
	var v1191 float64
	_ = v1191
	var v1203 int32
	_ = v1203
	var v1204 float64
	_ = v1204
	var v1213 int64
	_ = v1213
	var v1216 int64
	_ = v1216
	var v1217 int64
	_ = v1217
	var v1220 int64
	_ = v1220
	var v1221 int64
	_ = v1221
	var v1224 int64
	_ = v1224
	var v1225 int64
	_ = v1225
	var v1228 int64
	_ = v1228
	var v1229 int64
	_ = v1229
	var v1232 int64
	_ = v1232
	var v1233 int64
	_ = v1233
	var v1236 int64
	_ = v1236
	var v1237 int64
	_ = v1237
	var v1240 int64
	_ = v1240
	var v1241 int64
	_ = v1241
	var v1244 int64
	_ = v1244
	var v1245 int64
	_ = v1245
	var v1248 int64
	_ = v1248
	var v1249 int64
	_ = v1249
	var v1252 int64
	_ = v1252
	var v1253 int64
	_ = v1253
	var v1256 float64
	_ = v1256
	var v1257 int64
	_ = v1257
	var v1259 float64
	_ = v1259
	var v1263 float64
	_ = v1263
	var v1264 int64
	_ = v1264
	var v1270 float64
	_ = v1270
	var v1271 int64
	_ = v1271
	var v1277 float64
	_ = v1277
	var v1278 int64
	_ = v1278
	var v1284 float64
	_ = v1284
	var v1285 int64
	_ = v1285
	var v1291 float64
	_ = v1291
	var v1292 int64
	_ = v1292
	var v1298 float64
	_ = v1298
	var v1302 int64
	_ = v1302
	var v1303 int64
	_ = v1303
	var v1306 int64
	_ = v1306
	var v1307 int64
	_ = v1307
	var v1310 int64
	_ = v1310
	var v1311 int64
	_ = v1311
	var v1314 int64
	_ = v1314
	var v1315 int64
	_ = v1315
	var v1318 int64
	_ = v1318
	var v1319 int64
	_ = v1319
	var v1322 float64
	_ = v1322
	var v1323 int64
	_ = v1323
	var v1325 float64
	_ = v1325
	var v1329 int64
	_ = v1329
	var v1332 float64
	_ = v1332
	var v1335 int64
	_ = v1335
	var v1339 float64
	_ = v1339
	var v1342 int64
	_ = v1342
	var v1345 float64
	_ = v1345
	var v1348 int64
	_ = v1348
	var v1352 float64
	_ = v1352
	var v1355 int64
	_ = v1355
	var v1358 float64
	_ = v1358
	var v1361 int64
	_ = v1361
	var v1365 float64
	_ = v1365
	var v1368 int64
	_ = v1368
	var v1371 float64
	_ = v1371
	var v1374 int64
	_ = v1374
	var v1378 float64
	_ = v1378
	var v1382 int64
	_ = v1382
	var v1386 int64
	_ = v1386
	var v1390 int32
	_ = v1390
	var v1408 int32
	_ = v1408
	var v1424 int32
	_ = v1424
	var v1425 int32
	_ = v1425
	var v1427 int32
	_ = v1427
	var v1431 int32
	_ = v1431
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
	v1424 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_store[1]))
	v1425 = *(*int32)(unsafe.Add(mBase, uint32(v1424)))
	F_LWLockRelease(m, v1425)
	mBase = m.M
	v1427 = m.ExcPending
	if v1427 != 0 {
		goto L25
	} else {
		goto L242
	}
L28:
	;
	if l10 != 0 {
		v1408 = v1114
		goto L27
	} else {
		goto L208
	}
L29:
	;
	if v154 != 0 {
		v1102 = v154
		v1114 = v14
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
	v532 = base.AtomicRmwXchg32(m, v529, int32(20), int32(1))
	if v532 != 0 {
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
	v544 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v543)+24)))
	v545 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v543)+20)), uint32(v545))
	v550 = int64(*(*int32)(unsafe.Add(mBase, _c_F_pgss_store[6])))
	if base.Ui64(v550<<(uint(int64(9))%64)) <= base.Ui64(v544) {
		goto L93
	} else {
		goto L94
	}
L92:
	;
	goto L91
L93:
	;
	v554 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v543)+16)))
	v559 = base.B2i32(base.Ui64(v550*v554<<(uint(int64(1))%64)) <= base.Ui64(v544))
	goto L95
L94:
	;
	v559 = v545
	goto L95
L95:
	;
	v560 = *(*int32)(unsafe.Add(mBase, uint32(v543)))
	F_LWLockRelease(m, v560)
	mBase = m.M
	v562 = m.ExcPending
	if v562 != 0 {
		goto L25
	} else {
		goto L96
	}
L96:
	;
	v564 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_store[1]))
	v565 = *(*int32)(unsafe.Add(mBase, uint32(v564)))
	v567 = F_LWLockAcquire(m, v565, int32(0))
	mBase = m.M
	v568 = m.ExcPending
	if v568 != 0 {
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
	v584 = *(*int32)(unsafe.Add(mBase, uint32(v33)+116))
	v585 = *(*int32)(unsafe.Add(mBase, uint32(v33)+148))
	v588 = F_entry_alloc(m, v33+int32(120), v584, v585, v39, base.B2i32(l10 != int32(0)))
	mBase = m.M
	v589 = m.ExcPending
	if v589 != 0 {
		goto L25
	} else {
		goto L105
	}
L99:
	;
	v570 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_store[1]))
	v571 = *(*int32)(unsafe.Add(mBase, uint32(v570)+32))
	v572 = *(*int32)(unsafe.Add(mBase, uint32(v33)+112))
	if v571 == v572 {
		goto L98
	} else {
		goto L102
	}
L100:
	;
	goto L101
L101:
	;
	v574 = *(*int32)(unsafe.Add(mBase, uint32(v33)+148))
	v578 = F_qtext_store(m, v520, v574, v33+int32(116), int32(0))
	mBase = m.M
	v579 = m.ExcPending
	if v579 != 0 {
		goto L25
	} else {
		goto L103
	}
L102:
	;
	goto L101
L103:
	;
	if v578 == int32(0) {
		v1408 = v505
		goto L27
	} else {
		goto L104
	}
L104:
	;
	goto L98
L105:
	;
	if v559 == int32(0) {
		v1102 = v588
		v1114 = v505
		goto L28
	} else {
		goto L106
	}
L106:
	;
	v593 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_store[1]))
	v596 = base.AtomicRmwXchg32(m, v593, int32(20), int32(1))
	if v596 != 0 {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	v598 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_store[1]))
	F_s_lock(m, v598+int32(20), int32(_a_F_pgss_store_6), int32(2428), int32(_a_F_pgss_store_7))
	mBase = m.M
	v605 = m.ExcPending
	if v605 != 0 {
		goto L25
	} else {
		goto L110
	}
L108:
	;
	goto L109
L109:
	;
	v607 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_store[1]))
	v608 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v607)+24)))
	v609 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v607)+20)), uint32(v609))
	v613 = int64(*(*int32)(unsafe.Add(mBase, _c_F_pgss_store[6])))
	if base.Ui64(v608) < base.Ui64(v613<<(uint(int64(9))%64)) {
		v1102 = v588
		v1114 = v505
		goto L28
	} else {
		goto L111
	}
L110:
	;
	goto L109
L111:
	;
	v617 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v607)+16)))
	if base.Ui64(v608) < base.Ui64(v613*v617<<(uint(int64(1))%64)) {
		v1102 = v588
		v1114 = v505
		goto L28
	} else {
		goto L112
	}
L112:
	;
	v624 = F_qtext_load_file(m, v33+int32(220))
	mBase = m.M
	v625 = m.ExcPending
	if v625 != 0 {
		goto L25
	} else {
		goto L116
	}
L113:
	;
	v1091 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_store[1]))
	v1092 = *(*int32)(unsafe.Add(mBase, uint32(v1091)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v1091)+32)) = v1092 + int32(1)
	v1096 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v1091)+20)), uint32(v1096))
	v1102 = v588
	v1114 = v505
	goto L28
L114:
	;
	v1053 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_store[1]))
	F_s_lock(m, v1053+int32(20), int32(_a_F_pgss_store_6), v1051, int32(_a_F_pgss_store_8))
	mBase = m.M
	v1059 = m.ExcPending
	if v1059 != 0 {
		goto L25
	} else {
		goto L207
	}
L115:
	;
	F_emscripten_builtin_free(m, v624)
	mBase = m.M
	v908 = v33 + int32(160)
	v910 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_store[2]))
	F_hash_seq_init(m, v908, v910)
	mBase = m.M
	v912 = m.ExcPending
	if v912 != 0 {
		goto L25
	} else {
		goto L186
	}
L116:
	;
	if v624 == int32(0) {
		goto L115
	} else {
		goto L117
	}
L117:
	;
	v630 = F_AllocateFile(m, int32(_a_F_pgss_store_9), int32(_a_F_pgss_store_10))
	mBase = m.M
	v631 = m.ExcPending
	if v631 != 0 {
		goto L25
	} else {
		goto L118
	}
L118:
	;
	if v630 == int32(0) {
		goto L119
	} else {
		goto L120
	}
L119:
	;
	v636 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v637 = m.ExcPending
	if v637 != 0 {
		goto L25
	} else {
		goto L122
	}
L120:
	;
	goto L121
L121:
	;
	v655 = v33 + int32(160)
	v657 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_store[2]))
	F_hash_seq_init(m, v655, v657)
	mBase = m.M
	v659 = m.ExcPending
	if v659 != 0 {
		goto L25
	} else {
		goto L127
	}
L122:
	;
	if v636 == int32(0) {
		goto L115
	} else {
		goto L123
	}
L123:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v641 = m.ExcPending
	if v641 != 0 {
		goto L25
	} else {
		goto L124
	}
L124:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+16)) = int32(_a_F_pgss_store_9)
	F_errmsg(m, int32(_a_F_pgss_store_11), v33+int32(16))
	mBase = m.M
	v648 = m.ExcPending
	if v648 != 0 {
		goto L25
	} else {
		goto L125
	}
L125:
	;
	F_errfinish(m, int32(_a_F_pgss_store_6), int32(2513), int32(_a_F_pgss_store_8))
	mBase = m.M
	v653 = m.ExcPending
	if v653 != 0 {
		goto L25
	} else {
		goto L126
	}
L126:
	;
	goto L115
L127:
	;
	v660 = F_hash_seq_search(m, v655)
	mBase = m.M
	v661 = m.ExcPending
	if v661 != 0 {
		goto L25
	} else {
		goto L129
	}
L128:
	;
	v787 = *(*int32)(unsafe.Add(mBase, uint32(v630)+60))
	if v787 < int32(0) {
		goto L158
	} else {
		goto L159
	}
L129:
	;
	if v660 == int32(0) {
		goto L130
	} else {
		goto L131
	}
L130:
	;
	v664 = int32(0)
	v771 = v664
	v775 = v664
	goto L128
L131:
	;
	goto L132
L132:
	;
	v666 = *(*int32)(unsafe.Add(mBase, uint32(v33)+220))
	v667 = int32(0)
	v671 = v660
	v683 = v667
	v687 = v667
	goto L133
L133:
	;
	v699 = *(*int32)(unsafe.Add(mBase, uint32(v671)+396))
	if v699 < int32(0) {
		goto L137
	} else {
		goto L138
	}
L134:
	;
	v771 = v748
	v775 = v750
	goto L128
L135:
	;
	v755 = F_hash_seq_search(m, v33+int32(160))
	mBase = m.M
	v756 = m.ExcPending
	if v756 != 0 {
		goto L25
	} else {
		goto L154
	}
L136:
	;
	v714 = int32(1)
	v716 = v699 + v714
	v717 = F_fwrite(m, v624+v702, v714, v716, v630)
	mBase = m.M
	v718 = m.ExcPending
	if v718 != 0 {
		goto L25
	} else {
		goto L141
	}
L137:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v671)+392)) = int64(-4294967296)
	v748 = v683
	v750 = v687
	goto L135
L138:
	;
	v702 = *(*int32)(unsafe.Add(mBase, uint32(v671)+392))
	v703 = v702 + v699
	if base.Ui32(v666) <= base.Ui32(v703) {
		goto L137
	} else {
		goto L139
	}
L139:
	;
	v706 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v624+v703))))
	if v706 == int32(0) {
		goto L136
	} else {
		goto L140
	}
L140:
	;
	goto L137
L141:
	;
	if v717 != v716 {
		goto L142
	} else {
		goto L143
	}
L142:
	;
	v722 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v723 = m.ExcPending
	if v723 != 0 {
		goto L25
	} else {
		goto L145
	}
L143:
	;
	goto L144
L144:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v671)+392)) = v687
	v748 = v683 + int32(1)
	v750 = v716 + v687
	goto L135
L145:
	;
	if v722 != 0 {
		goto L146
	} else {
		goto L147
	}
L146:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v725 = m.ExcPending
	if v725 != 0 {
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
	v741 = m.ExcPending
	if v741 != 0 {
		goto L25
	} else {
		goto L152
	}
L149:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+80)) = int32(_a_F_pgss_store_9)
	F_errmsg(m, int32(_a_F_pgss_store_11), v33+int32(80))
	mBase = m.M
	v732 = m.ExcPending
	if v732 != 0 {
		goto L25
	} else {
		goto L150
	}
L150:
	;
	F_errfinish(m, int32(_a_F_pgss_store_6), int32(2543), int32(_a_F_pgss_store_8))
	mBase = m.M
	v737 = m.ExcPending
	if v737 != 0 {
		goto L25
	} else {
		goto L151
	}
L151:
	;
	goto L148
L152:
	;
	v742 = F_FreeFile(m, v630)
	mBase = m.M
	v743 = m.ExcPending
	if v743 != 0 {
		goto L25
	} else {
		goto L153
	}
L153:
	;
	goto L115
L154:
	;
	if v755 != 0 {
		v671 = v755
		v683 = v748
		v687 = v750
		goto L133
	} else {
		goto L155
	}
L155:
	;
	goto L134
L156:
	;
	v819 = F_FreeFile(m, v630)
	mBase = m.M
	v820 = m.ExcPending
	if v820 != 0 {
		goto L25
	} else {
		goto L167
	}
L157:
	;
	v796 = F_ftruncate(m, v794, base.I64_extend_i32_u(v775))
	mBase = m.M
	if v796 == int32(0) {
		goto L156
	} else {
		goto L161
	}
L158:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_pgss_store[7])) = int32(8)
	v794 = int32(-1)
	goto L160
L159:
	;
	v794 = v787
	goto L160
L160:
	;
	goto L157
L161:
	;
	v801 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v802 = m.ExcPending
	if v802 != 0 {
		goto L25
	} else {
		goto L162
	}
L162:
	;
	if v801 == int32(0) {
		goto L156
	} else {
		goto L163
	}
L163:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v806 = m.ExcPending
	if v806 != 0 {
		goto L25
	} else {
		goto L164
	}
L164:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+64)) = int32(_a_F_pgss_store_9)
	F_errmsg(m, int32(_a_F_pgss_store_12), v33-int32(-64))
	mBase = m.M
	v813 = m.ExcPending
	if v813 != 0 {
		goto L25
	} else {
		goto L165
	}
L165:
	;
	F_errfinish(m, int32(_a_F_pgss_store_6), int32(2561), int32(_a_F_pgss_store_8))
	mBase = m.M
	v818 = m.ExcPending
	if v818 != 0 {
		goto L25
	} else {
		goto L166
	}
L166:
	;
	goto L156
L167:
	;
	if v819 != 0 {
		goto L168
	} else {
		goto L169
	}
L168:
	;
	v823 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v824 = m.ExcPending
	if v824 != 0 {
		goto L25
	} else {
		goto L171
	}
L169:
	;
	goto L170
L170:
	;
	v843 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v844 = m.ExcPending
	if v844 != 0 {
		goto L25
	} else {
		goto L176
	}
L171:
	;
	if v823 == int32(0) {
		goto L115
	} else {
		goto L172
	}
L172:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v828 = m.ExcPending
	if v828 != 0 {
		goto L25
	} else {
		goto L173
	}
L173:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+48)) = int32(_a_F_pgss_store_9)
	F_errmsg(m, int32(_a_F_pgss_store_11), v33+int32(48))
	mBase = m.M
	v835 = m.ExcPending
	if v835 != 0 {
		goto L25
	} else {
		goto L174
	}
L174:
	;
	F_errfinish(m, int32(_a_F_pgss_store_6), int32(2568), int32(_a_F_pgss_store_8))
	mBase = m.M
	v840 = m.ExcPending
	if v840 != 0 {
		goto L25
	} else {
		goto L175
	}
L175:
	;
	goto L115
L176:
	;
	if v843 != 0 {
		goto L177
	} else {
		goto L178
	}
L177:
	;
	v846 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_store[1]))
	v847 = *(*int32)(unsafe.Add(mBase, uint32(v846)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v33)+32)) = v847
	*(*int32)(unsafe.Add(mBase, uint32(v33)+36)) = v775
	F_errmsg_internal(m, int32(_a_F_pgss_store_13), v33+int32(32))
	mBase = m.M
	v854 = m.ExcPending
	if v854 != 0 {
		goto L25
	} else {
		goto L180
	}
L178:
	;
	goto L179
L179:
	;
	v861 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_store[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v861)+24)) = v775
	if v771 <= int32(0) {
		goto L182
	} else {
		goto L183
	}
L180:
	;
	F_errfinish(m, int32(_a_F_pgss_store_6), int32(2574), int32(_a_F_pgss_store_8))
	mBase = m.M
	v859 = m.ExcPending
	if v859 != 0 {
		goto L25
	} else {
		goto L181
	}
L181:
	;
	goto L179
L182:
	;
	v867 = int32(1024)
	goto L184
L183:
	;
	v866 = base.I32_div_u_s(v775, v771)
	v867 = v866
	goto L184
L184:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v861)+16)) = v867
	F_emscripten_builtin_free(m, v624)
	mBase = m.M
	v872 = base.AtomicRmwXchg32(m, v861, int32(20), int32(1))
	if v872 == int32(0) {
		goto L113
	} else {
		goto L185
	}
L185:
	;
	v1051 = int32(2597)
	goto L114
L186:
	;
	v913 = F_hash_seq_search(m, v908)
	mBase = m.M
	v914 = m.ExcPending
	if v914 != 0 {
		goto L25
	} else {
		goto L187
	}
L187:
	;
	if v913 != 0 {
		goto L188
	} else {
		goto L189
	}
L188:
	;
	v917 = v913
	goto L191
L189:
	;
	goto L190
L190:
	;
	v981 = int32(_a_F_pgss_store_9)
	v982 = F_unlink(m, v981)
	mBase = m.M
	v985 = F_AllocateFile(m, v981, int32(_a_F_pgss_store_10))
	mBase = m.M
	v986 = m.ExcPending
	if v986 != 0 {
		goto L25
	} else {
		goto L196
	}
L191:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v917)+392)) = int64(-4294967296)
	v949 = F_hash_seq_search(m, v33+int32(160))
	mBase = m.M
	v950 = m.ExcPending
	if v950 != 0 {
		goto L25
	} else {
		goto L193
	}
L192:
	;
	goto L190
L193:
	;
	if v949 != 0 {
		v917 = v949
		goto L191
	} else {
		goto L194
	}
L194:
	;
	goto L192
L195:
	;
	v1010 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_store[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v1010)+16)) = int32(1024)
	v1013 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1010)+24)) = v1013
	v1017 = base.AtomicRmwXchg32(m, v1010, int32(20), int32(1))
	if v1017 == v1013 {
		goto L113
	} else {
		goto L206
	}
L196:
	;
	if v985 == int32(0) {
		goto L197
	} else {
		goto L198
	}
L197:
	;
	v991 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v992 = m.ExcPending
	if v992 != 0 {
		goto L25
	} else {
		goto L200
	}
L198:
	;
	goto L199
L199:
	;
	v1007 = F_FreeFile(m, v985)
	mBase = m.M
	v1008 = m.ExcPending
	if v1008 != 0 {
		goto L25
	} else {
		goto L205
	}
L200:
	;
	if v991 == int32(0) {
		goto L195
	} else {
		goto L201
	}
L201:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v996 = m.ExcPending
	if v996 != 0 {
		goto L25
	} else {
		goto L202
	}
L202:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33))) = int32(_a_F_pgss_store_9)
	F_errmsg(m, int32(_a_F_pgss_store_14), v33)
	mBase = m.M
	v1001 = m.ExcPending
	if v1001 != 0 {
		goto L25
	} else {
		goto L203
	}
L203:
	;
	F_errfinish(m, int32(_a_F_pgss_store_6), int32(2627), int32(_a_F_pgss_store_8))
	mBase = m.M
	v1006 = m.ExcPending
	if v1006 != 0 {
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
	v1051 = int32(2648)
	goto L114
L207:
	;
	goto L113
L208:
	;
	v1131 = base.AtomicRmwXchg32(m, v1102, int32(424), int32(1))
	if v1131 != 0 {
		goto L209
	} else {
		goto L210
	}
L209:
	;
	F_s_lock(m, v1102+int32(424), int32(_a_F_pgss_store_6), int32(1404), int32(_a_F_pgss_store_15))
	mBase = m.M
	v1138 = m.ExcPending
	if v1138 != 0 {
		goto L25
	} else {
		goto L212
	}
L210:
	;
	goto L211
L211:
	;
	v1139 = *(*int64)(unsafe.Add(mBase, uint32(v1102)+24))
	v1141 = *(*int64)(unsafe.Add(mBase, uint32(v1102)+32))
	if v1139 == int64(0)-v1141 {
		goto L213
	} else {
		goto L214
	}
L212:
	;
	goto L211
L213:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1102)+256)) = int64(4607182418800017408)
	goto L215
L214:
	;
	goto L215
L215:
	;
	v1147 = l4 << (uint(int32(3)) % 32)
	v1150 = v1147 + (v1102 + int32(24))
	v1151 = *(*int64)(unsafe.Add(mBase, uint32(v1150)))
	v1153 = v1151 + int64(1)
	*(*int64)(unsafe.Add(mBase, uint32(v1150))) = v1153
	v1155 = v1147 + v1102
	v1157 = v1155 + int32(40)
	v1158 = *(*float64)(unsafe.Add(mBase, uint32(v1157)))
	*(*float64)(unsafe.Add(mBase, uint32(v1157))) = base.F64_add(l5, v1158)
	if v1151 == int64(0) {
		goto L217
	} else {
		goto L218
	}
L216:
	;
	v1213 = *(*int64)(unsafe.Add(mBase, uint32(v1102)+120))
	*(*int64)(unsafe.Add(mBase, uint32(v1102)+120)) = v1213 + l6
	v1216 = *(*int64)(unsafe.Add(mBase, uint32(v1102)+128))
	v1217 = *(*int64)(unsafe.Add(mBase, uint32(l7)))
	*(*int64)(unsafe.Add(mBase, uint32(v1102)+128)) = v1216 + v1217
	v1220 = *(*int64)(unsafe.Add(mBase, uint32(v1102)+136))
	v1221 = *(*int64)(unsafe.Add(mBase, uint32(l7)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v1102)+136)) = v1220 + v1221
	v1224 = *(*int64)(unsafe.Add(mBase, uint32(v1102)+144))
	v1225 = *(*int64)(unsafe.Add(mBase, uint32(l7)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v1102)+144)) = v1224 + v1225
	v1228 = *(*int64)(unsafe.Add(mBase, uint32(v1102)+152))
	v1229 = *(*int64)(unsafe.Add(mBase, uint32(l7)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v1102)+152)) = v1228 + v1229
	v1232 = *(*int64)(unsafe.Add(mBase, uint32(v1102)+160))
	v1233 = *(*int64)(unsafe.Add(mBase, uint32(l7)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v1102)+160)) = v1232 + v1233
	v1236 = *(*int64)(unsafe.Add(mBase, uint32(v1102)+168))
	v1237 = *(*int64)(unsafe.Add(mBase, uint32(l7)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v1102)+168)) = v1236 + v1237
	v1240 = *(*int64)(unsafe.Add(mBase, uint32(v1102)+176))
	v1241 = *(*int64)(unsafe.Add(mBase, uint32(l7)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v1102)+176)) = v1240 + v1241
	v1244 = *(*int64)(unsafe.Add(mBase, uint32(v1102)+184))
	v1245 = *(*int64)(unsafe.Add(mBase, uint32(l7)+56))
	*(*int64)(unsafe.Add(mBase, uint32(v1102)+184)) = v1244 + v1245
	v1248 = *(*int64)(unsafe.Add(mBase, uint32(v1102)+192))
	v1249 = *(*int64)(unsafe.Add(mBase, uint32(l7)+64))
	*(*int64)(unsafe.Add(mBase, uint32(v1102)+192)) = v1248 + v1249
	v1252 = *(*int64)(unsafe.Add(mBase, uint32(v1102)+200))
	v1253 = *(*int64)(unsafe.Add(mBase, uint32(l7)+72))
	*(*int64)(unsafe.Add(mBase, uint32(v1102)+200)) = v1252 + v1253
	v1256 = *(*float64)(unsafe.Add(mBase, uint32(v1102)+208))
	v1257 = *(*int64)(unsafe.Add(mBase, uint32(l7)+80))
	v1259 = float64(1e+06)
	*(*float64)(unsafe.Add(mBase, uint32(v1102)+208)) = base.F64_add(v1256, base.F64_div(base.F64_convert_i64_s(v1257), v1259))
	v1263 = *(*float64)(unsafe.Add(mBase, uint32(v1102)+216))
	v1264 = *(*int64)(unsafe.Add(mBase, uint32(l7)+88))
	*(*float64)(unsafe.Add(mBase, uint32(v1102)+216)) = base.F64_add(v1263, base.F64_div(base.F64_convert_i64_s(v1264), v1259))
	v1270 = *(*float64)(unsafe.Add(mBase, uint32(v1102)+224))
	v1271 = *(*int64)(unsafe.Add(mBase, uint32(l7)+96))
	*(*float64)(unsafe.Add(mBase, uint32(v1102)+224)) = base.F64_add(v1270, base.F64_div(base.F64_convert_i64_s(v1271), v1259))
	v1277 = *(*float64)(unsafe.Add(mBase, uint32(v1102)+232))
	v1278 = *(*int64)(unsafe.Add(mBase, uint32(l7)+104))
	*(*float64)(unsafe.Add(mBase, uint32(v1102)+232)) = base.F64_add(v1277, base.F64_div(base.F64_convert_i64_s(v1278), v1259))
	v1284 = *(*float64)(unsafe.Add(mBase, uint32(v1102)+240))
	v1285 = *(*int64)(unsafe.Add(mBase, uint32(l7)+112))
	*(*float64)(unsafe.Add(mBase, uint32(v1102)+240)) = base.F64_add(v1284, base.F64_div(base.F64_convert_i64_s(v1285), v1259))
	v1291 = *(*float64)(unsafe.Add(mBase, uint32(v1102)+248))
	v1292 = *(*int64)(unsafe.Add(mBase, uint32(l7)+120))
	*(*float64)(unsafe.Add(mBase, uint32(v1102)+248)) = base.F64_add(v1291, base.F64_div(base.F64_convert_i64_s(v1292), v1259))
	v1298 = *(*float64)(unsafe.Add(mBase, uint32(v1102)+256))
	*(*float64)(unsafe.Add(mBase, uint32(v1102)+256)) = base.F64_add(v1298, float64(1))
	v1302 = *(*int64)(unsafe.Add(mBase, uint32(v1102)+264))
	v1303 = *(*int64)(unsafe.Add(mBase, uint32(l8)))
	*(*int64)(unsafe.Add(mBase, uint32(v1102)+264)) = v1302 + v1303
	v1306 = *(*int64)(unsafe.Add(mBase, uint32(v1102)+272))
	v1307 = *(*int64)(unsafe.Add(mBase, uint32(l8)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v1102)+272)) = v1306 + v1307
	v1310 = *(*int64)(unsafe.Add(mBase, uint32(v1102)+280))
	v1311 = *(*int64)(unsafe.Add(mBase, uint32(l8)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v1102)+280)) = v1310 + v1311
	v1314 = *(*int64)(unsafe.Add(mBase, uint32(v1102)+288))
	v1315 = *(*int64)(unsafe.Add(mBase, uint32(l8)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v1102)+288)) = v1314 + v1315
	if l9 != 0 {
		goto L227
	} else {
		goto L228
	}
L217:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v1155)+88)) = l5
	*(*float64)(unsafe.Add(mBase, uint32(v1155)+72)) = l5
	*(*float64)(unsafe.Add(mBase, uint32(v1155)+56)) = l5
	goto L216
L218:
	;
	goto L219
L219:
	;
	v1167 = v1155 + int32(88)
	v1168 = *(*float64)(unsafe.Add(mBase, uint32(v1167)))
	v1169 = base.F64_sub(l5, v1168)
	v1172 = base.F64_add(v1168, base.F64_div(v1169, base.F64_convert_i64_s(v1153)))
	*(*float64)(unsafe.Add(mBase, uint32(v1167))) = v1172
	v1175 = v1155 + int32(104)
	v1178 = *(*float64)(unsafe.Add(mBase, uint32(v1175)))
	*(*float64)(unsafe.Add(mBase, uint32(v1175))) = base.F64_add(base.F64_mul(v1169, base.F64_sub(l5, v1172)), v1178)
	v1182 = v1155 + int32(56)
	v1183 = *(*float64)(unsafe.Add(mBase, uint32(v1182)))
	if base.F64_ne(v1183, float64(0)) != 0 {
		goto L220
	} else {
		goto L221
	}
L220:
	;
	if base.F64_lt(l5, v1183) != 0 {
		goto L223
	} else {
		goto L224
	}
L221:
	;
	v1190 = v1102 + l4<<(uint(int32(3))%32) + int32(72)
	v1191 = *(*float64)(unsafe.Add(mBase, uint32(v1190)))
	if base.F64_ne(v1191, float64(0)) != 0 {
		goto L220
	} else {
		goto L222
	}
L222:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v1182))) = l5
	*(*float64)(unsafe.Add(mBase, uint32(v1190))) = l5
	goto L216
L223:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v1182))) = l5
	goto L225
L224:
	;
	goto L225
L225:
	;
	v1203 = v1102 + l4<<(uint(int32(3))%32) + int32(72)
	v1204 = *(*float64)(unsafe.Add(mBase, uint32(v1203)))
	if base.F64_lt(v1204, l5) == int32(0) {
		goto L216
	} else {
		goto L226
	}
L226:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v1203))) = l5
	goto L216
L227:
	;
	v1318 = *(*int64)(unsafe.Add(mBase, uint32(v1102)+296))
	v1319 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l9))))
	*(*int64)(unsafe.Add(mBase, uint32(v1102)+296)) = v1318 + v1319
	v1322 = *(*float64)(unsafe.Add(mBase, uint32(v1102)+304))
	v1323 = *(*int64)(unsafe.Add(mBase, uint32(l9)+8))
	v1325 = float64(1e+06)
	*(*float64)(unsafe.Add(mBase, uint32(v1102)+304)) = base.F64_add(v1322, base.F64_div(base.F64_convert_i64_s(v1323), v1325))
	v1329 = *(*int64)(unsafe.Add(mBase, uint32(l9)+16))
	v1332 = base.F64_div(base.F64_convert_i64_s(v1329), v1325)
	if base.F64_ne(v1332, float64(0)) != 0 {
		goto L230
	} else {
		goto L231
	}
L228:
	;
	goto L229
L229:
	;
	v1382 = *(*int64)(unsafe.Add(mBase, uint32(v1102)+376))
	*(*int64)(unsafe.Add(mBase, uint32(v1102)+376)) = v1382 + base.I64_extend_i32_s(l11)
	v1386 = *(*int64)(unsafe.Add(mBase, uint32(v1102)+384))
	*(*int64)(unsafe.Add(mBase, uint32(v1102)+384)) = v1386 + base.I64_extend_i32_s(l12)
	v1390 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v1102)+424)), uint32(v1390))
	v1408 = v1114
	goto L27
L230:
	;
	v1335 = *(*int64)(unsafe.Add(mBase, uint32(v1102)+328))
	*(*int64)(unsafe.Add(mBase, uint32(v1102)+328)) = v1335 + int64(1)
	goto L232
L231:
	;
	goto L232
L232:
	;
	v1339 = *(*float64)(unsafe.Add(mBase, uint32(v1102)+320))
	*(*float64)(unsafe.Add(mBase, uint32(v1102)+320)) = base.F64_add(v1332, v1339)
	v1342 = *(*int64)(unsafe.Add(mBase, uint32(l9)+24))
	v1345 = base.F64_div(base.F64_convert_i64_s(v1342), float64(1e+06))
	if base.F64_ne(v1345, float64(0)) != 0 {
		goto L233
	} else {
		goto L234
	}
L233:
	;
	v1348 = *(*int64)(unsafe.Add(mBase, uint32(v1102)+312))
	*(*int64)(unsafe.Add(mBase, uint32(v1102)+312)) = v1348 + int64(1)
	goto L235
L234:
	;
	goto L235
L235:
	;
	v1352 = *(*float64)(unsafe.Add(mBase, uint32(v1102)+336))
	*(*float64)(unsafe.Add(mBase, uint32(v1102)+336)) = base.F64_add(v1345, v1352)
	v1355 = *(*int64)(unsafe.Add(mBase, uint32(l9)+32))
	v1358 = base.F64_div(base.F64_convert_i64_s(v1355), float64(1e+06))
	if base.F64_ne(v1358, float64(0)) != 0 {
		goto L236
	} else {
		goto L237
	}
L236:
	;
	v1361 = *(*int64)(unsafe.Add(mBase, uint32(v1102)+344))
	*(*int64)(unsafe.Add(mBase, uint32(v1102)+344)) = v1361 + int64(1)
	goto L238
L237:
	;
	goto L238
L238:
	;
	v1365 = *(*float64)(unsafe.Add(mBase, uint32(v1102)+352))
	*(*float64)(unsafe.Add(mBase, uint32(v1102)+352)) = base.F64_add(v1358, v1365)
	v1368 = *(*int64)(unsafe.Add(mBase, uint32(l9)+40))
	v1371 = base.F64_div(base.F64_convert_i64_s(v1368), float64(1e+06))
	if base.F64_ne(v1371, float64(0)) != 0 {
		goto L239
	} else {
		goto L240
	}
L239:
	;
	v1374 = *(*int64)(unsafe.Add(mBase, uint32(v1102)+360))
	*(*int64)(unsafe.Add(mBase, uint32(v1102)+360)) = v1374 + int64(1)
	goto L241
L240:
	;
	goto L241
L241:
	;
	v1378 = *(*float64)(unsafe.Add(mBase, uint32(v1102)+368))
	*(*float64)(unsafe.Add(mBase, uint32(v1102)+368)) = base.F64_add(v1371, v1378)
	goto L229
L242:
	;
	if v1408 == int32(0) {
		goto L2
	} else {
		goto L243
	}
L243:
	;
	F_pfree(m, v1408)
	mBase = m.M
	v1431 = m.ExcPending
	if v1431 != 0 {
		goto L25
	} else {
		goto L244
	}
L244:
	;
	goto L2
}
