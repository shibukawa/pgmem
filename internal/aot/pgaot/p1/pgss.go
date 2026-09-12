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
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int64
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	v5 = *(*int32)(unsafe.Add(mBase, _consts[1473]))
	if v5 != 0 {
		m.T0[v5].(func(*base.Module, int32, int32))(m, l0, l1)
		mBase = m.M
		v7 = m.ExcPending
		if v7 != 0 {
			return
		} else {
			v11 = *(*int32)(unsafe.Add(mBase, _consts[20]))
			if int32(0) <= v11 {
				return
			} else {
				v15 = *(*int32)(unsafe.Add(mBase, _consts[1482]))
				if v15 != int32(2) {
					if v15 != int32(1) {
						return
					} else {
						v23 = *(*int32)(unsafe.Add(mBase, _consts[1483]))
						if v23 != 0 {
							return
						} else {
							v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							v26 = *(*int64)(unsafe.Add(mBase, uint32(v25)+8))
							if v26 == int64(0) {
								return
							} else {
								v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
								if v29 != 0 {
									return
								} else {
									v30 = int32(4562096)
									v31 = *(*int32)(unsafe.Add(mBase, _consts[3]))
									v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
									v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+100))
									*(*int32)(unsafe.Add(mBase, _consts[3])) = v33
									v38 = F_InstrAlloc(m, int32(1), int32(2147483647), int32(0))
									mBase = m.M
									v39 = m.ExcPending
									if v39 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v38
										*(*int32)(unsafe.Add(mBase, _consts[3])) = v31
										return
									}
								}
							}
						}
					}
				} else {
					v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v26 = *(*int64)(unsafe.Add(mBase, uint32(v25)+8))
					if v26 == int64(0) {
						return
					} else {
						v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
						if v29 != 0 {
							return
						} else {
							v30 = int32(4562096)
							v31 = *(*int32)(unsafe.Add(mBase, _consts[3]))
							v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
							v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+100))
							*(*int32)(unsafe.Add(mBase, _consts[3])) = v33
							v38 = F_InstrAlloc(m, int32(1), int32(2147483647), int32(0))
							mBase = m.M
							v39 = m.ExcPending
							if v39 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v38
								*(*int32)(unsafe.Add(mBase, _consts[3])) = v31
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
			v11 = *(*int32)(unsafe.Add(mBase, _consts[20]))
			if int32(0) <= v11 {
				return
			} else {
				v15 = *(*int32)(unsafe.Add(mBase, _consts[1482]))
				if v15 != int32(2) {
					if v15 != int32(1) {
						return
					} else {
						v23 = *(*int32)(unsafe.Add(mBase, _consts[1483]))
						if v23 != 0 {
							return
						} else {
							v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							v26 = *(*int64)(unsafe.Add(mBase, uint32(v25)+8))
							if v26 == int64(0) {
								return
							} else {
								v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
								if v29 != 0 {
									return
								} else {
									v30 = int32(4562096)
									v31 = *(*int32)(unsafe.Add(mBase, _consts[3]))
									v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
									v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+100))
									*(*int32)(unsafe.Add(mBase, _consts[3])) = v33
									v38 = F_InstrAlloc(m, int32(1), int32(2147483647), int32(0))
									mBase = m.M
									v39 = m.ExcPending
									if v39 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v38
										*(*int32)(unsafe.Add(mBase, _consts[3])) = v31
										return
									}
								}
							}
						}
					}
				} else {
					v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v26 = *(*int64)(unsafe.Add(mBase, uint32(v25)+8))
					if v26 == int64(0) {
						return
					} else {
						v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
						if v29 != 0 {
							return
						} else {
							v30 = int32(4562096)
							v31 = *(*int32)(unsafe.Add(mBase, _consts[3]))
							v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
							v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+100))
							*(*int32)(unsafe.Add(mBase, _consts[3])) = v33
							v38 = F_InstrAlloc(m, int32(1), int32(2147483647), int32(0))
							mBase = m.M
							v39 = m.ExcPending
							if v39 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v38
								*(*int32)(unsafe.Add(mBase, _consts[3])) = v31
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
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
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
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
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
	var v89 int32
	_ = v89
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v124 int32
	_ = v124
	var v125 int64
	_ = v125
	var v128 int32
	_ = v128
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v203 int32
	_ = v203
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v307 int32
	_ = v307
	var v321 int32
	_ = v321
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v370 int32
	_ = v370
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v389 int32
	_ = v389
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v403 int32
	_ = v403
	var v406 int32
	_ = v406
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v417 int32
	_ = v417
	var v419 int32
	_ = v419
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v427 int32
	_ = v427
	var v429 int32
	_ = v429
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v438 int32
	_ = v438
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v448 int32
	_ = v448
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v484 int32
	_ = v484
	var v486 int32
	_ = v486
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v538 int32
	_ = v538
	var v545 int32
	_ = v545
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v551 int64
	_ = v551
	var v553 int64
	_ = v553
	var v557 int64
	_ = v557
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
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
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v577 int32
	_ = v577
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v601 int32
	_ = v601
	var v608 int32
	_ = v608
	var v610 int32
	_ = v610
	var v613 int64
	_ = v613
	var v615 int64
	_ = v615
	var v619 int64
	_ = v619
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v632 int32
	_ = v632
	var v633 int32
	_ = v633
	var v638 int32
	_ = v638
	var v639 int32
	_ = v639
	var v643 int32
	_ = v643
	var v650 int32
	_ = v650
	var v655 int32
	_ = v655
	var v659 int32
	_ = v659
	var v661 int32
	_ = v661
	var v664 int32
	_ = v664
	var v665 int32
	_ = v665
	var v668 int32
	_ = v668
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	var v675 int32
	_ = v675
	var v690 int32
	_ = v690
	var v692 int32
	_ = v692
	var v702 int32
	_ = v702
	var v705 int32
	_ = v705
	var v706 int32
	_ = v706
	var v709 int32
	_ = v709
	var v717 int32
	_ = v717
	var v719 int32
	_ = v719
	var v720 int32
	_ = v720
	var v721 int32
	_ = v721
	var v725 int32
	_ = v725
	var v726 int32
	_ = v726
	var v728 int32
	_ = v728
	var v735 int32
	_ = v735
	var v740 int32
	_ = v740
	var v744 int32
	_ = v744
	var v745 int32
	_ = v745
	var v746 int32
	_ = v746
	var v754 int32
	_ = v754
	var v755 int32
	_ = v755
	var v758 int32
	_ = v758
	var v759 int32
	_ = v759
	var v777 int32
	_ = v777
	var v779 int32
	_ = v779
	var v789 int32
	_ = v789
	var v792 int32
	_ = v792
	var v793 int32
	_ = v793
	var v794 int32
	_ = v794
	var v801 int32
	_ = v801
	var v803 int32
	_ = v803
	var v808 int32
	_ = v808
	var v809 int32
	_ = v809
	var v813 int32
	_ = v813
	var v820 int32
	_ = v820
	var v825 int32
	_ = v825
	var v826 int32
	_ = v826
	var v827 int32
	_ = v827
	var v830 int32
	_ = v830
	var v831 int32
	_ = v831
	var v835 int32
	_ = v835
	var v842 int32
	_ = v842
	var v847 int32
	_ = v847
	var v850 int32
	_ = v850
	var v851 int32
	_ = v851
	var v853 int32
	_ = v853
	var v854 int32
	_ = v854
	var v861 int32
	_ = v861
	var v866 int32
	_ = v866
	var v868 int32
	_ = v868
	var v873 int32
	_ = v873
	var v874 int32
	_ = v874
	var v877 int32
	_ = v877
	var v916 int32
	_ = v916
	var v918 int32
	_ = v918
	var v921 int32
	_ = v921
	var v922 int32
	_ = v922
	var v925 int32
	_ = v925
	var v956 int32
	_ = v956
	var v957 int32
	_ = v957
	var v987 int32
	_ = v987
	var v988 int32
	_ = v988
	var v991 int32
	_ = v991
	var v992 int32
	_ = v992
	var v997 int32
	_ = v997
	var v998 int32
	_ = v998
	var v1002 int32
	_ = v1002
	var v1007 int32
	_ = v1007
	var v1012 int32
	_ = v1012
	var v1013 int32
	_ = v1013
	var v1014 int32
	_ = v1014
	var v1016 int32
	_ = v1016
	var v1019 int32
	_ = v1019
	var v1021 int32
	_ = v1021
	var v1056 int32
	_ = v1056
	var v1058 int32
	_ = v1058
	var v1064 int32
	_ = v1064
	var v1095 int32
	_ = v1095
	var v1098 int32
	_ = v1098
	var v1102 int32
	_ = v1102
	var v1105 int32
	_ = v1105
	var v1131 int32
	_ = v1131
	var v1140 int32
	_ = v1140
	var v1143 int64
	_ = v1143
	var v1145 int64
	_ = v1145
	var v1151 int32
	_ = v1151
	var v1152 int32
	_ = v1152
	var v1153 int64
	_ = v1153
	var v1155 int64
	_ = v1155
	var v1157 int32
	_ = v1157
	var v1159 int32
	_ = v1159
	var v1160 float64
	_ = v1160
	var v1169 int32
	_ = v1169
	var v1170 float64
	_ = v1170
	var v1171 float64
	_ = v1171
	var v1174 float64
	_ = v1174
	var v1177 int32
	_ = v1177
	var v1180 float64
	_ = v1180
	var v1184 int32
	_ = v1184
	var v1185 float64
	_ = v1185
	var v1192 int32
	_ = v1192
	var v1193 float64
	_ = v1193
	var v1205 int32
	_ = v1205
	var v1206 float64
	_ = v1206
	var v1215 int64
	_ = v1215
	var v1218 int64
	_ = v1218
	var v1219 int64
	_ = v1219
	var v1222 int64
	_ = v1222
	var v1223 int64
	_ = v1223
	var v1226 int64
	_ = v1226
	var v1227 int64
	_ = v1227
	var v1230 int64
	_ = v1230
	var v1231 int64
	_ = v1231
	var v1234 int64
	_ = v1234
	var v1235 int64
	_ = v1235
	var v1238 int64
	_ = v1238
	var v1239 int64
	_ = v1239
	var v1242 int64
	_ = v1242
	var v1243 int64
	_ = v1243
	var v1246 int64
	_ = v1246
	var v1247 int64
	_ = v1247
	var v1250 int64
	_ = v1250
	var v1251 int64
	_ = v1251
	var v1254 int64
	_ = v1254
	var v1255 int64
	_ = v1255
	var v1258 float64
	_ = v1258
	var v1259 int64
	_ = v1259
	var v1261 float64
	_ = v1261
	var v1265 float64
	_ = v1265
	var v1266 int64
	_ = v1266
	var v1272 float64
	_ = v1272
	var v1273 int64
	_ = v1273
	var v1279 float64
	_ = v1279
	var v1280 int64
	_ = v1280
	var v1286 float64
	_ = v1286
	var v1287 int64
	_ = v1287
	var v1293 float64
	_ = v1293
	var v1294 int64
	_ = v1294
	var v1300 float64
	_ = v1300
	var v1304 int64
	_ = v1304
	var v1305 int64
	_ = v1305
	var v1308 int64
	_ = v1308
	var v1309 int64
	_ = v1309
	var v1312 int64
	_ = v1312
	var v1313 int64
	_ = v1313
	var v1316 int64
	_ = v1316
	var v1317 int64
	_ = v1317
	var v1320 int64
	_ = v1320
	var v1321 int64
	_ = v1321
	var v1324 float64
	_ = v1324
	var v1325 int64
	_ = v1325
	var v1327 float64
	_ = v1327
	var v1331 int64
	_ = v1331
	var v1334 float64
	_ = v1334
	var v1337 int64
	_ = v1337
	var v1341 float64
	_ = v1341
	var v1344 int64
	_ = v1344
	var v1347 float64
	_ = v1347
	var v1350 int64
	_ = v1350
	var v1354 float64
	_ = v1354
	var v1357 int64
	_ = v1357
	var v1360 float64
	_ = v1360
	var v1363 int64
	_ = v1363
	var v1367 float64
	_ = v1367
	var v1370 int64
	_ = v1370
	var v1373 float64
	_ = v1373
	var v1376 int64
	_ = v1376
	var v1380 float64
	_ = v1380
	var v1386 int64
	_ = v1386
	var v1390 int64
	_ = v1390
	var v1394 int32
	_ = v1394
	var v1424 int32
	_ = v1424
	var v1425 int32
	_ = v1425
	var v1427 int32
	_ = v1427
	var v1431 int32
	_ = v1431
	v14 = int32(0)
	v30 = m.G0
	v32 = v30 - int32(224)
	m.G0 = v32
	*(*int32)(unsafe.Add(mBase, uint32(v32)+148)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v32)+152)) = l2
	v37 = *(*int32)(unsafe.Add(mBase, _consts[356]))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
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
	m.G0 = v32 + int32(224)
	return
L3:
	;
	v42 = *(*int32)(unsafe.Add(mBase, _consts[1484]))
	if v42 == int32(0) {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v46 = *(*int32)(unsafe.Add(mBase, _consts[1485]))
	if v46 == int32(0) {
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v50 = v32 + int32(152)
	v52 = v32 + int32(148)
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
	if v57 < int32(0) {
		goto L10
	} else {
		goto L11
	}
L6:
	;
	v124 = v32 + int32(136)
	v125 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v124))) = v125
	v128 = v32 + int32(128)
	*(*int64)(unsafe.Add(mBase, uint32(v128))) = v125
	*(*int64)(unsafe.Add(mBase, uint32(v32)+120)) = v125
	v134 = *(*int32)(unsafe.Add(mBase, _consts[31]))
	*(*int64)(unsafe.Add(mBase, uint32(v128))) = l1
	v137 = *(*int32)(unsafe.Add(mBase, _consts[1483]))
	*(*uint8)(unsafe.Add(mBase, uint32(v124))) = uint8(base.B2i32(v137 == int32(0)))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+120)) = v134
	v143 = *(*int32)(unsafe.Add(mBase, _consts[158]))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+124)) = v143
	v146 = *(*int32)(unsafe.Add(mBase, _consts[1484]))
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v146)))
	v149 = F_LWLockAcquire(m, v147, int32(1))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L24
	} else {
		goto L25
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v50))) = v118
	*(*int32)(unsafe.Add(mBase, uint32(v52))) = v117
	goto L6
L8:
	;
	v75 = v71
	v78 = v72
	v79 = v73
	goto L16
L9:
	;
	v68 = F_strlen(m, v65)
	mBase = m.M
	if v68 <= int32(0) {
		v114 = v65
		v117 = v68
		v118 = v67
		goto L7
	} else {
		goto L14
	}
L10:
	;
	v65 = l0
	v67 = int32(0)
	goto L9
L11:
	;
	goto L12
L12:
	;
	v61 = l0 + v57
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
	if int32(0) < v62 {
		v71 = v61
		v72 = v62
		v73 = v57
		goto L8
	} else {
		goto L13
	}
L13:
	;
	v65 = v61
	v67 = v57
	goto L9
L14:
	;
	v71 = v65
	v72 = v68
	v73 = v67
	goto L8
L15:
	;
	v100 = v78
	goto L20
L16:
	;
	v82 = int32(*(*int8)(unsafe.Add(mBase, uint32(v75))))
	v83 = F_scanner_isspace(m, v82)
	mBase = m.M
	if v83 == int32(0) {
		goto L15
	} else {
		goto L18
	}
L17:
	;
	v114 = v89
	v117 = int32(0)
	v118 = v72 + v73
	goto L7
L18:
	;
	v86 = int32(1)
	v89 = v75 + v86
	if v86 < v78 {
		v75 = v89
		v78 = v78 - v86
		v79 = v79 + v86
		goto L16
	} else {
		goto L19
	}
L19:
	;
	goto L17
L20:
	;
	v105 = int32(*(*int8)(unsafe.Add(mBase, uint32(v100+(v75-int32(1))))))
	v106 = F_scanner_isspace(m, v105)
	mBase = m.M
	if v106 == int32(0) {
		v114 = v75
		v117 = v100
		v118 = v79
		goto L7
	} else {
		goto L22
	}
L21:
	;
	v114 = v75
	v117 = int32(0)
	v118 = v79
	goto L7
L22:
	;
	v109 = int32(1)
	if v109 < v100 {
		v100 = v100 - v109
		goto L20
	} else {
		goto L23
	}
L23:
	;
	goto L21
L24:
	;
	return
L25:
	;
	v151 = int32(0)
	v153 = *(*int32)(unsafe.Add(mBase, _consts[1485]))
	v158 = F_hash_search(m, v153, v32+int32(120), v151, v151)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L24
	} else {
		goto L28
	}
L26:
	;
	v1424 = *(*int32)(unsafe.Add(mBase, _consts[1484]))
	v1425 = *(*int32)(unsafe.Add(mBase, uint32(v1424)))
	F_LWLockRelease(m, v1425)
	mBase = m.M
	v1427 = m.ExcPending
	if v1427 != 0 {
		goto L24
	} else {
		goto L247
	}
L27:
	;
	if l10 != 0 {
		v1394 = v1102
		goto L26
	} else {
		goto L213
	}
L28:
	;
	if v158 != 0 {
		v1102 = v151
		v1105 = v158
		goto L27
	} else {
		goto L29
	}
L29:
	;
	if l10 != 0 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v161 = *(*int32)(unsafe.Add(mBase, _consts[1484]))
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v161)))
	F_LWLockRelease(m, v162)
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L24
	} else {
		goto L33
	}
L31:
	;
	v495 = v151
	goto L32
L32:
	;
	if v495 != 0 {
		goto L86
	} else {
		goto L87
	}
L33:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v32)+148))
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v32)+152))
	v167 = *(*int32)(unsafe.Add(mBase, uint32(l10)+16))
	if int32(2) <= v167 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(l10)+8))
	F_pg_qsort(m, v170, v167, int32(12), int32(7703))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L24
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(l10)+8))
	v180 = F_scanner_init(m, v114, v32+int32(160), int32(1897756), int32(1354976))
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L24
	} else {
		goto L38
	}
L37:
	;
	goto L36
L38:
	;
	v182 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+180)) = uint8(v182)
	v184 = *(*int32)(unsafe.Add(mBase, uint32(l10)+16))
	if v184 <= v182 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	F_scanner_finish(m, v180)
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L24
	} else {
		goto L59
	}
L40:
	;
	v187 = v151
	v203 = v184
	goto L41
L41:
	;
	if v187 == int32(0) {
		goto L44
	} else {
		goto L45
	}
L42:
	;
	goto L39
L43:
	;
	v321 = v187 + int32(1)
	if v321 < v307 {
		v187 = v321
		v203 = v307
		goto L41
	} else {
		goto L58
	}
L44:
	;
	v231 = v175 + v187*int32(12)
	v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v231)+8)))
	if v232 != 0 {
		v307 = v203
		goto L43
	} else {
		goto L47
	}
L45:
	;
	v218 = int32(12)
	v220 = v175 + v187*v218
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v220)))
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v220-v218)))
	if v221 != v224 {
		goto L44
	} else {
		goto L46
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v220)+4)) = int32(-1)
	v307 = v203
	goto L43
L47:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v231)))
	v234 = v233 - v166
	goto L48
L48:
	;
	v268 = F_core_yylex(m, v32+int32(220), v32+int32(156), v180)
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L24
	} else {
		goto L50
	}
L49:
	;
	v275 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v234+v114))))
	if v275 == int32(45) {
		goto L53
	} else {
		goto L54
	}
L50:
	;
	if v268 == int32(0) {
		goto L39
	} else {
		goto L51
	}
L51:
	;
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v32)+156))
	if v272 < v234 {
		goto L48
	} else {
		goto L52
	}
L52:
	;
	goto L49
L53:
	;
	v282 = F_core_yylex(m, v32+int32(220), v32+int32(156), v180)
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L24
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v32)+160))
	v288 = F_strlen(m, v286+v234)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v231)+4)) = v288
	v290 = *(*int32)(unsafe.Add(mBase, uint32(l10)+16))
	v307 = v290
	goto L43
L56:
	;
	if v282 == int32(0) {
		goto L39
	} else {
		goto L57
	}
L57:
	;
	goto L55
L58:
	;
	goto L42
L59:
	;
	v354 = *(*int32)(unsafe.Add(mBase, uint32(l10)+16))
	v360 = F_palloc(m, v165+v354*int32(10)+int32(1))
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L24
	} else {
		goto L60
	}
L60:
	;
	v362 = int32(0)
	v363 = *(*int32)(unsafe.Add(mBase, uint32(l10)+16))
	if v362 < v363 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v370 = int32(0)
	v382 = v363
	v383 = v362
	v384 = v14
	v386 = v14
	v387 = v14
	v389 = v14
	goto L64
L62:
	;
	v466 = v362
	v467 = v14
	goto L63
L63:
	;
	v481 = v165 - v467
	if v481 != 0 {
		goto L82
	} else {
		goto L83
	}
L64:
	;
	v397 = v370 * int32(12)
	v398 = *(*int32)(unsafe.Add(mBase, uint32(l10)+8))
	v399 = v397 + v398
	v400 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v399)+9)))
	if v400 == int32(1) {
		goto L67
	} else {
		goto L68
	}
L65:
	;
	v466 = v441
	v467 = v442
	goto L63
L66:
	;
	v448 = v370 + int32(1)
	if v448 < v440 {
		v370 = v448
		v382 = v440
		v383 = v441
		v384 = v442
		v386 = v444
		v387 = v445
		v389 = v446
		goto L64
	} else {
		goto L80
	}
L67:
	;
	v403 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l10)+24)))
	if v403 != int32(1) {
		v440 = v382
		v441 = v383
		v442 = v384
		v444 = v386
		v445 = v387
		v446 = v389
		goto L66
	} else {
		goto L70
	}
L68:
	;
	goto L69
L69:
	;
	v406 = *(*int32)(unsafe.Add(mBase, uint32(v399)+4))
	if v406 < int32(0) {
		v440 = v382
		v441 = v383
		v442 = v384
		v444 = v386
		v445 = v387
		v446 = v389
		goto L66
	} else {
		goto L71
	}
L70:
	;
	goto L69
L71:
	;
	v412 = *(*int32)(unsafe.Add(mBase, uint32(v399)))
	v413 = v412 - v166
	v414 = v413 - (v387 + v389)
	if v414 != 0 {
		goto L73
	} else {
		goto L74
	}
L72:
	;
	v417 = *(*int32)(unsafe.Add(mBase, uint32(l10)+8))
	v419 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v417+v397)+8)))
	v421 = v386 + int32(1)
	v422 = *(*int32)(unsafe.Add(mBase, uint32(l10)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+96)) = v421 + v422
	if v419 != 0 {
		goto L76
	} else {
		goto L77
	}
L73:
	;
	v415 = F__emscripten_memcpy_bulkmem(m, v360+v383, v384+v114, v414)
	mBase = m.M
	goto L75
L74:
	;
	goto L75
L75:
	;
	goto L72
L76:
	;
	v427 = int32(600151)
	goto L78
L77:
	;
	v427 = int32(794587)
	goto L78
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+100)) = v427
	v429 = v414 + v383
	v434 = F_pg_sprintf(m, v360+v429, int32(187011), v32+int32(96))
	mBase = m.M
	v435 = m.ExcPending
	if v435 != 0 {
		goto L24
	} else {
		goto L79
	}
L79:
	;
	v438 = *(*int32)(unsafe.Add(mBase, uint32(l10)+16))
	v440 = v438
	v441 = v434 + v429
	v442 = v406 + v413
	v444 = v421
	v445 = v406
	v446 = v413
	goto L66
L80:
	;
	goto L65
L81:
	;
	v484 = v481 + v466
	v486 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v360+v484))) = uint8(v486)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+148)) = v484
	v490 = *(*int32)(unsafe.Add(mBase, _consts[1484]))
	v491 = *(*int32)(unsafe.Add(mBase, uint32(v490)))
	v493 = F_LWLockAcquire(m, v491, int32(1))
	mBase = m.M
	v494 = m.ExcPending
	if v494 != 0 {
		goto L24
	} else {
		goto L85
	}
L82:
	;
	v482 = F__emscripten_memcpy_bulkmem(m, v360+v466, v467+v114, v481)
	mBase = m.M
	goto L84
L83:
	;
	goto L84
L84:
	;
	goto L81
L85:
	;
	v495 = v360
	goto L32
L86:
	;
	v524 = v495
	goto L88
L87:
	;
	v524 = v114
	goto L88
L88:
	;
	v525 = *(*int32)(unsafe.Add(mBase, uint32(v32)+148))
	v530 = F_qtext_store(m, v524, v525, v32+int32(116), v32+int32(112))
	mBase = m.M
	v531 = m.ExcPending
	if v531 != 0 {
		goto L24
	} else {
		goto L89
	}
L89:
	;
	v533 = *(*int32)(unsafe.Add(mBase, _consts[1484]))
	v534 = *(*int32)(unsafe.Add(mBase, uint32(v533)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v533)+20)) = int32(1)
	if v534 != 0 {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v538 = *(*int32)(unsafe.Add(mBase, _consts[1484]))
	F_s_lock(m, v538+int32(20), int32(518750), int32(2428), int32(123253))
	mBase = m.M
	v545 = m.ExcPending
	if v545 != 0 {
		goto L24
	} else {
		goto L93
	}
L91:
	;
	goto L92
L92:
	;
	v547 = *(*int32)(unsafe.Add(mBase, _consts[1484]))
	v548 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v547)+20)) = v548
	v551 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v547)+24)))
	v553 = int64(*(*int32)(unsafe.Add(mBase, _consts[1486])))
	if base.Ui64(v553<<(uint(int64(9))%64)) <= base.Ui64(v551) {
		goto L94
	} else {
		goto L95
	}
L93:
	;
	goto L92
L94:
	;
	v557 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v547)+16)))
	v562 = base.B2i32(base.Ui64(v553*v557<<(uint(int64(1))%64)) <= base.Ui64(v551))
	goto L96
L95:
	;
	v562 = v548
	goto L96
L96:
	;
	v563 = *(*int32)(unsafe.Add(mBase, uint32(v547)))
	F_LWLockRelease(m, v563)
	mBase = m.M
	v565 = m.ExcPending
	if v565 != 0 {
		goto L24
	} else {
		goto L97
	}
L97:
	;
	v567 = *(*int32)(unsafe.Add(mBase, _consts[1484]))
	v568 = *(*int32)(unsafe.Add(mBase, uint32(v567)))
	v570 = F_LWLockAcquire(m, v568, int32(0))
	mBase = m.M
	v571 = m.ExcPending
	if v571 != 0 {
		goto L24
	} else {
		goto L98
	}
L98:
	;
	if v530 != 0 {
		goto L100
	} else {
		goto L101
	}
L99:
	;
	v587 = *(*int32)(unsafe.Add(mBase, uint32(v32)+116))
	v588 = *(*int32)(unsafe.Add(mBase, uint32(v32)+148))
	v591 = F_entry_alloc(m, v32+int32(120), v587, v588, v38, base.B2i32(l10 != int32(0)))
	mBase = m.M
	v592 = m.ExcPending
	if v592 != 0 {
		goto L24
	} else {
		goto L106
	}
L100:
	;
	v573 = *(*int32)(unsafe.Add(mBase, _consts[1484]))
	v574 = *(*int32)(unsafe.Add(mBase, uint32(v573)+32))
	v575 = *(*int32)(unsafe.Add(mBase, uint32(v32)+112))
	if v574 == v575 {
		goto L99
	} else {
		goto L103
	}
L101:
	;
	goto L102
L102:
	;
	v577 = *(*int32)(unsafe.Add(mBase, uint32(v32)+148))
	v581 = F_qtext_store(m, v524, v577, v32+int32(116), int32(0))
	mBase = m.M
	v582 = m.ExcPending
	if v582 != 0 {
		goto L24
	} else {
		goto L104
	}
L103:
	;
	goto L102
L104:
	;
	if v581 == int32(0) {
		v1394 = v495
		goto L26
	} else {
		goto L105
	}
L105:
	;
	goto L99
L106:
	;
	if v562 == int32(0) {
		v1102 = v495
		v1105 = v591
		goto L27
	} else {
		goto L107
	}
L107:
	;
	v596 = *(*int32)(unsafe.Add(mBase, _consts[1484]))
	v597 = *(*int32)(unsafe.Add(mBase, uint32(v596)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v596)+20)) = int32(1)
	if v597 != 0 {
		goto L108
	} else {
		goto L109
	}
L108:
	;
	v601 = *(*int32)(unsafe.Add(mBase, _consts[1484]))
	F_s_lock(m, v601+int32(20), int32(518750), int32(2428), int32(123253))
	mBase = m.M
	v608 = m.ExcPending
	if v608 != 0 {
		goto L24
	} else {
		goto L111
	}
L109:
	;
	goto L110
L110:
	;
	v610 = *(*int32)(unsafe.Add(mBase, _consts[1484]))
	*(*int32)(unsafe.Add(mBase, uint32(v610)+20)) = int32(0)
	v613 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v610)+24)))
	v615 = int64(*(*int32)(unsafe.Add(mBase, _consts[1486])))
	if base.Ui64(v613) < base.Ui64(v615<<(uint(int64(9))%64)) {
		v1102 = v495
		v1105 = v591
		goto L27
	} else {
		goto L112
	}
L111:
	;
	goto L110
L112:
	;
	v619 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v610)+16)))
	if base.Ui64(v613) < base.Ui64(v615*v619<<(uint(int64(1))%64)) {
		v1102 = v495
		v1105 = v591
		goto L27
	} else {
		goto L113
	}
L113:
	;
	v626 = F_qtext_load_file(m, v32+int32(220))
	mBase = m.M
	v627 = m.ExcPending
	if v627 != 0 {
		goto L24
	} else {
		goto L117
	}
L114:
	;
	v1095 = *(*int32)(unsafe.Add(mBase, _consts[1484]))
	*(*int32)(unsafe.Add(mBase, uint32(v1095)+20)) = int32(0)
	v1098 = *(*int32)(unsafe.Add(mBase, uint32(v1095)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v1095)+32)) = v1098 + int32(1)
	v1102 = v495
	v1105 = v591
	goto L27
L115:
	;
	v1058 = *(*int32)(unsafe.Add(mBase, _consts[1484]))
	F_s_lock(m, v1058+int32(20), int32(518750), v1056, int32(123258))
	mBase = m.M
	v1064 = m.ExcPending
	if v1064 != 0 {
		goto L24
	} else {
		goto L212
	}
L116:
	;
	F_emscripten_builtin_free(m, v626)
	mBase = m.M
	v916 = *(*int32)(unsafe.Add(mBase, _consts[1485]))
	F_hash_seq_init(m, v32+int32(160), v916)
	mBase = m.M
	v918 = m.ExcPending
	if v918 != 0 {
		goto L24
	} else {
		goto L191
	}
L117:
	;
	if v626 == int32(0) {
		goto L116
	} else {
		goto L118
	}
L118:
	;
	v632 = F_AllocateFile(m, int32(119677), int32(34101))
	mBase = m.M
	v633 = m.ExcPending
	if v633 != 0 {
		goto L24
	} else {
		goto L119
	}
L119:
	;
	if v632 == int32(0) {
		goto L120
	} else {
		goto L121
	}
L120:
	;
	v638 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v639 = m.ExcPending
	if v639 != 0 {
		goto L24
	} else {
		goto L123
	}
L121:
	;
	goto L122
L122:
	;
	v659 = *(*int32)(unsafe.Add(mBase, _consts[1485]))
	F_hash_seq_init(m, v32+int32(160), v659)
	mBase = m.M
	v661 = m.ExcPending
	if v661 != 0 {
		goto L24
	} else {
		goto L128
	}
L123:
	;
	if v638 == int32(0) {
		goto L116
	} else {
		goto L124
	}
L124:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v643 = m.ExcPending
	if v643 != 0 {
		goto L24
	} else {
		goto L125
	}
L125:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+16)) = int32(119677)
	F_errmsg(m, int32(314307), v32+int32(16))
	mBase = m.M
	v650 = m.ExcPending
	if v650 != 0 {
		goto L24
	} else {
		goto L126
	}
L126:
	;
	F_errfinish(m, int32(518750), int32(2513), int32(123258))
	mBase = m.M
	v655 = m.ExcPending
	if v655 != 0 {
		goto L24
	} else {
		goto L127
	}
L127:
	;
	goto L116
L128:
	;
	v664 = F_hash_seq_search(m, v32+int32(160))
	mBase = m.M
	v665 = m.ExcPending
	if v665 != 0 {
		goto L24
	} else {
		goto L130
	}
L129:
	;
	v789 = *(*int32)(unsafe.Add(mBase, uint32(v632)+76))
	if v789 < int32(0) {
		goto L160
	} else {
		goto L161
	}
L130:
	;
	if v664 == int32(0) {
		goto L131
	} else {
		goto L132
	}
L131:
	;
	v668 = int32(0)
	v777 = v668
	v779 = v668
	goto L129
L132:
	;
	goto L133
L133:
	;
	v670 = *(*int32)(unsafe.Add(mBase, uint32(v32)+220))
	v671 = int32(0)
	v675 = v664
	v690 = v671
	v692 = v671
	goto L134
L134:
	;
	v702 = *(*int32)(unsafe.Add(mBase, uint32(v675)+396))
	if v702 < int32(0) {
		goto L138
	} else {
		goto L139
	}
L135:
	;
	v777 = v754
	v779 = v755
	goto L129
L136:
	;
	v758 = F_hash_seq_search(m, v32+int32(160))
	mBase = m.M
	v759 = m.ExcPending
	if v759 != 0 {
		goto L24
	} else {
		goto L155
	}
L137:
	;
	v717 = int32(1)
	v719 = v702 + v717
	v720 = F_fwrite(m, v705+v626, v717, v719, v632)
	mBase = m.M
	v721 = m.ExcPending
	if v721 != 0 {
		goto L24
	} else {
		goto L142
	}
L138:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v675)+392)) = int64(-4294967296)
	v754 = v690
	v755 = v692
	goto L136
L139:
	;
	v705 = *(*int32)(unsafe.Add(mBase, uint32(v675)+392))
	v706 = v705 + v702
	if base.Ui32(v670) <= base.Ui32(v706) {
		goto L138
	} else {
		goto L140
	}
L140:
	;
	v709 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v706+v626))))
	if v709 == int32(0) {
		goto L137
	} else {
		goto L141
	}
L141:
	;
	goto L138
L142:
	;
	if v720 != v719 {
		goto L143
	} else {
		goto L144
	}
L143:
	;
	v725 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v726 = m.ExcPending
	if v726 != 0 {
		goto L24
	} else {
		goto L146
	}
L144:
	;
	goto L145
L145:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v675)+392)) = v690
	v754 = v719 + v690
	v755 = v692 + int32(1)
	goto L136
L146:
	;
	if v725 != 0 {
		goto L147
	} else {
		goto L148
	}
L147:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v728 = m.ExcPending
	if v728 != 0 {
		goto L24
	} else {
		goto L150
	}
L148:
	;
	goto L149
L149:
	;
	F_hash_seq_term(m, v32+int32(160))
	mBase = m.M
	v744 = m.ExcPending
	if v744 != 0 {
		goto L24
	} else {
		goto L153
	}
L150:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+80)) = int32(119677)
	F_errmsg(m, int32(314307), v32+int32(80))
	mBase = m.M
	v735 = m.ExcPending
	if v735 != 0 {
		goto L24
	} else {
		goto L151
	}
L151:
	;
	F_errfinish(m, int32(518750), int32(2543), int32(123258))
	mBase = m.M
	v740 = m.ExcPending
	if v740 != 0 {
		goto L24
	} else {
		goto L152
	}
L152:
	;
	goto L149
L153:
	;
	v745 = F_FreeFile(m, v632)
	mBase = m.M
	v746 = m.ExcPending
	if v746 != 0 {
		goto L24
	} else {
		goto L154
	}
L154:
	;
	goto L116
L155:
	;
	if v758 != 0 {
		v675 = v758
		v690 = v754
		v692 = v755
		goto L134
	} else {
		goto L156
	}
L156:
	;
	goto L135
L157:
	;
	v826 = F_FreeFile(m, v632)
	mBase = m.M
	v827 = m.ExcPending
	if v827 != 0 {
		goto L24
	} else {
		goto L172
	}
L158:
	;
	v803 = F_ftruncate(m, v801, base.I64_extend_i32_u(v777))
	mBase = m.M
	if v803 == int32(0) {
		goto L157
	} else {
		goto L166
	}
L159:
	;
	if v794 < int32(0) {
		goto L163
	} else {
		goto L164
	}
L160:
	;
	v792 = *(*int32)(unsafe.Add(mBase, uint32(v632)+60))
	v794 = v792
	goto L159
L161:
	;
	goto L162
L162:
	;
	v793 = *(*int32)(unsafe.Add(mBase, uint32(v632)+60))
	v794 = v793
	goto L159
L163:
	;
	*(*int32)(unsafe.Add(mBase, _consts[137])) = int32(8)
	v801 = int32(-1)
	goto L165
L164:
	;
	v801 = v794
	goto L165
L165:
	;
	goto L158
L166:
	;
	v808 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v809 = m.ExcPending
	if v809 != 0 {
		goto L24
	} else {
		goto L167
	}
L167:
	;
	if v808 == int32(0) {
		goto L157
	} else {
		goto L168
	}
L168:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v813 = m.ExcPending
	if v813 != 0 {
		goto L24
	} else {
		goto L169
	}
L169:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+64)) = int32(119677)
	F_errmsg(m, int32(314432), v32-int32(-64))
	mBase = m.M
	v820 = m.ExcPending
	if v820 != 0 {
		goto L24
	} else {
		goto L170
	}
L170:
	;
	F_errfinish(m, int32(518750), int32(2561), int32(123258))
	mBase = m.M
	v825 = m.ExcPending
	if v825 != 0 {
		goto L24
	} else {
		goto L171
	}
L171:
	;
	goto L157
L172:
	;
	if v826 != 0 {
		goto L173
	} else {
		goto L174
	}
L173:
	;
	v830 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v831 = m.ExcPending
	if v831 != 0 {
		goto L24
	} else {
		goto L176
	}
L174:
	;
	goto L175
L175:
	;
	v850 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v851 = m.ExcPending
	if v851 != 0 {
		goto L24
	} else {
		goto L181
	}
L176:
	;
	if v830 == int32(0) {
		goto L116
	} else {
		goto L177
	}
L177:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v835 = m.ExcPending
	if v835 != 0 {
		goto L24
	} else {
		goto L178
	}
L178:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+48)) = int32(119677)
	F_errmsg(m, int32(314307), v32+int32(48))
	mBase = m.M
	v842 = m.ExcPending
	if v842 != 0 {
		goto L24
	} else {
		goto L179
	}
L179:
	;
	F_errfinish(m, int32(518750), int32(2568), int32(123258))
	mBase = m.M
	v847 = m.ExcPending
	if v847 != 0 {
		goto L24
	} else {
		goto L180
	}
L180:
	;
	goto L116
L181:
	;
	if v850 != 0 {
		goto L182
	} else {
		goto L183
	}
L182:
	;
	v853 = *(*int32)(unsafe.Add(mBase, _consts[1484]))
	v854 = *(*int32)(unsafe.Add(mBase, uint32(v853)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+32)) = v854
	*(*int32)(unsafe.Add(mBase, uint32(v32)+36)) = v777
	F_errmsg_internal(m, int32(38582), v32+int32(32))
	mBase = m.M
	v861 = m.ExcPending
	if v861 != 0 {
		goto L24
	} else {
		goto L185
	}
L183:
	;
	goto L184
L184:
	;
	v868 = *(*int32)(unsafe.Add(mBase, _consts[1484]))
	*(*int32)(unsafe.Add(mBase, uint32(v868)+24)) = v777
	if v779 <= int32(0) {
		goto L187
	} else {
		goto L188
	}
L185:
	;
	F_errfinish(m, int32(518750), int32(2574), int32(123258))
	mBase = m.M
	v866 = m.ExcPending
	if v866 != 0 {
		goto L24
	} else {
		goto L186
	}
L186:
	;
	goto L184
L187:
	;
	v874 = int32(1024)
	goto L189
L188:
	;
	v873 = base.I32_div_u_s(v777, v779)
	v874 = v873
	goto L189
L189:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v868)+16)) = v874
	F_emscripten_builtin_free(m, v626)
	mBase = m.M
	v877 = *(*int32)(unsafe.Add(mBase, uint32(v868)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v868)+20)) = int32(1)
	if v877 == int32(0) {
		goto L114
	} else {
		goto L190
	}
L190:
	;
	v1056 = int32(2597)
	goto L115
L191:
	;
	v921 = F_hash_seq_search(m, v32+int32(160))
	mBase = m.M
	v922 = m.ExcPending
	if v922 != 0 {
		goto L24
	} else {
		goto L192
	}
L192:
	;
	if v921 != 0 {
		goto L193
	} else {
		goto L194
	}
L193:
	;
	v925 = v921
	goto L196
L194:
	;
	goto L195
L195:
	;
	v987 = int32(119677)
	v988 = F_unlink(m, v987)
	mBase = m.M
	v991 = F_AllocateFile(m, v987, int32(34101))
	mBase = m.M
	v992 = m.ExcPending
	if v992 != 0 {
		goto L24
	} else {
		goto L201
	}
L196:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v925)+392)) = int64(-4294967296)
	v956 = F_hash_seq_search(m, v32+int32(160))
	mBase = m.M
	v957 = m.ExcPending
	if v957 != 0 {
		goto L24
	} else {
		goto L198
	}
L197:
	;
	goto L195
L198:
	;
	if v956 != 0 {
		v925 = v956
		goto L196
	} else {
		goto L199
	}
L199:
	;
	goto L197
L200:
	;
	v1016 = *(*int32)(unsafe.Add(mBase, _consts[1484]))
	*(*int32)(unsafe.Add(mBase, uint32(v1016)+16)) = int32(1024)
	v1019 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1016)+24)) = v1019
	v1021 = *(*int32)(unsafe.Add(mBase, uint32(v1016)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v1016)+20)) = int32(1)
	if v1021 == v1019 {
		goto L114
	} else {
		goto L211
	}
L201:
	;
	if v991 == int32(0) {
		goto L202
	} else {
		goto L203
	}
L202:
	;
	v997 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v998 = m.ExcPending
	if v998 != 0 {
		goto L24
	} else {
		goto L205
	}
L203:
	;
	goto L204
L204:
	;
	v1013 = F_FreeFile(m, v991)
	mBase = m.M
	v1014 = m.ExcPending
	if v1014 != 0 {
		goto L24
	} else {
		goto L210
	}
L205:
	;
	if v997 == int32(0) {
		goto L200
	} else {
		goto L206
	}
L206:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v1002 = m.ExcPending
	if v1002 != 0 {
		goto L24
	} else {
		goto L207
	}
L207:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32))) = int32(119677)
	F_errmsg(m, int32(314368), v32)
	mBase = m.M
	v1007 = m.ExcPending
	if v1007 != 0 {
		goto L24
	} else {
		goto L208
	}
L208:
	;
	F_errfinish(m, int32(518750), int32(2627), int32(123258))
	mBase = m.M
	v1012 = m.ExcPending
	if v1012 != 0 {
		goto L24
	} else {
		goto L209
	}
L209:
	;
	goto L200
L210:
	;
	goto L200
L211:
	;
	v1056 = int32(2648)
	goto L115
L212:
	;
	goto L114
L213:
	;
	v1131 = *(*int32)(unsafe.Add(mBase, uint32(v1105)+424))
	*(*int32)(unsafe.Add(mBase, uint32(v1105)+424)) = int32(1)
	if v1131 != 0 {
		goto L214
	} else {
		goto L215
	}
L214:
	;
	F_s_lock(m, v1105+int32(424), int32(518750), int32(1404), int32(383665))
	mBase = m.M
	v1140 = m.ExcPending
	if v1140 != 0 {
		goto L24
	} else {
		goto L217
	}
L215:
	;
	goto L216
L216:
	;
	v1143 = *(*int64)(unsafe.Add(mBase, uint32(v1105)+24))
	v1145 = *(*int64)(unsafe.Add(mBase, uint32(v1105)+32))
	if v1143 == int64(0)-v1145 {
		goto L218
	} else {
		goto L219
	}
L217:
	;
	goto L216
L218:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1105)+256)) = int64(4607182418800017408)
	goto L220
L219:
	;
	goto L220
L220:
	;
	v1151 = l4 << (uint(int32(3)) % 32)
	v1152 = v1105 + int32(24) + v1151
	v1153 = *(*int64)(unsafe.Add(mBase, uint32(v1152)))
	v1155 = v1153 + int64(1)
	*(*int64)(unsafe.Add(mBase, uint32(v1152))) = v1155
	v1157 = v1151 + v1105
	v1159 = v1157 + int32(40)
	v1160 = *(*float64)(unsafe.Add(mBase, uint32(v1159)))
	*(*float64)(unsafe.Add(mBase, uint32(v1159))) = base.F64_add(l5, v1160)
	if v1153 == int64(0) {
		goto L222
	} else {
		goto L223
	}
L221:
	;
	v1215 = *(*int64)(unsafe.Add(mBase, uint32(v1105)+120))
	*(*int64)(unsafe.Add(mBase, uint32(v1105)+120)) = v1215 + l6
	v1218 = *(*int64)(unsafe.Add(mBase, uint32(v1105)+128))
	v1219 = *(*int64)(unsafe.Add(mBase, uint32(l7)))
	*(*int64)(unsafe.Add(mBase, uint32(v1105)+128)) = v1218 + v1219
	v1222 = *(*int64)(unsafe.Add(mBase, uint32(v1105)+136))
	v1223 = *(*int64)(unsafe.Add(mBase, uint32(l7)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v1105)+136)) = v1222 + v1223
	v1226 = *(*int64)(unsafe.Add(mBase, uint32(v1105)+144))
	v1227 = *(*int64)(unsafe.Add(mBase, uint32(l7)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v1105)+144)) = v1226 + v1227
	v1230 = *(*int64)(unsafe.Add(mBase, uint32(v1105)+152))
	v1231 = *(*int64)(unsafe.Add(mBase, uint32(l7)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v1105)+152)) = v1230 + v1231
	v1234 = *(*int64)(unsafe.Add(mBase, uint32(v1105)+160))
	v1235 = *(*int64)(unsafe.Add(mBase, uint32(l7)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v1105)+160)) = v1234 + v1235
	v1238 = *(*int64)(unsafe.Add(mBase, uint32(v1105)+168))
	v1239 = *(*int64)(unsafe.Add(mBase, uint32(l7)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v1105)+168)) = v1238 + v1239
	v1242 = *(*int64)(unsafe.Add(mBase, uint32(v1105)+176))
	v1243 = *(*int64)(unsafe.Add(mBase, uint32(l7)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v1105)+176)) = v1242 + v1243
	v1246 = *(*int64)(unsafe.Add(mBase, uint32(v1105)+184))
	v1247 = *(*int64)(unsafe.Add(mBase, uint32(l7)+56))
	*(*int64)(unsafe.Add(mBase, uint32(v1105)+184)) = v1246 + v1247
	v1250 = *(*int64)(unsafe.Add(mBase, uint32(v1105)+192))
	v1251 = *(*int64)(unsafe.Add(mBase, uint32(l7)+64))
	*(*int64)(unsafe.Add(mBase, uint32(v1105)+192)) = v1250 + v1251
	v1254 = *(*int64)(unsafe.Add(mBase, uint32(v1105)+200))
	v1255 = *(*int64)(unsafe.Add(mBase, uint32(l7)+72))
	*(*int64)(unsafe.Add(mBase, uint32(v1105)+200)) = v1254 + v1255
	v1258 = *(*float64)(unsafe.Add(mBase, uint32(v1105)+208))
	v1259 = *(*int64)(unsafe.Add(mBase, uint32(l7)+80))
	v1261 = float64(1e+06)
	*(*float64)(unsafe.Add(mBase, uint32(v1105)+208)) = base.F64_add(v1258, base.F64_div(base.F64_convert_i64_s(v1259), v1261))
	v1265 = *(*float64)(unsafe.Add(mBase, uint32(v1105)+216))
	v1266 = *(*int64)(unsafe.Add(mBase, uint32(l7)+88))
	*(*float64)(unsafe.Add(mBase, uint32(v1105)+216)) = base.F64_add(v1265, base.F64_div(base.F64_convert_i64_s(v1266), v1261))
	v1272 = *(*float64)(unsafe.Add(mBase, uint32(v1105)+224))
	v1273 = *(*int64)(unsafe.Add(mBase, uint32(l7)+96))
	*(*float64)(unsafe.Add(mBase, uint32(v1105)+224)) = base.F64_add(v1272, base.F64_div(base.F64_convert_i64_s(v1273), v1261))
	v1279 = *(*float64)(unsafe.Add(mBase, uint32(v1105)+232))
	v1280 = *(*int64)(unsafe.Add(mBase, uint32(l7)+104))
	*(*float64)(unsafe.Add(mBase, uint32(v1105)+232)) = base.F64_add(v1279, base.F64_div(base.F64_convert_i64_s(v1280), v1261))
	v1286 = *(*float64)(unsafe.Add(mBase, uint32(v1105)+240))
	v1287 = *(*int64)(unsafe.Add(mBase, uint32(l7)+112))
	*(*float64)(unsafe.Add(mBase, uint32(v1105)+240)) = base.F64_add(v1286, base.F64_div(base.F64_convert_i64_s(v1287), v1261))
	v1293 = *(*float64)(unsafe.Add(mBase, uint32(v1105)+248))
	v1294 = *(*int64)(unsafe.Add(mBase, uint32(l7)+120))
	*(*float64)(unsafe.Add(mBase, uint32(v1105)+248)) = base.F64_add(v1293, base.F64_div(base.F64_convert_i64_s(v1294), v1261))
	v1300 = *(*float64)(unsafe.Add(mBase, uint32(v1105)+256))
	*(*float64)(unsafe.Add(mBase, uint32(v1105)+256)) = base.F64_add(v1300, float64(1))
	v1304 = *(*int64)(unsafe.Add(mBase, uint32(v1105)+264))
	v1305 = *(*int64)(unsafe.Add(mBase, uint32(l8)))
	*(*int64)(unsafe.Add(mBase, uint32(v1105)+264)) = v1304 + v1305
	v1308 = *(*int64)(unsafe.Add(mBase, uint32(v1105)+272))
	v1309 = *(*int64)(unsafe.Add(mBase, uint32(l8)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v1105)+272)) = v1308 + v1309
	v1312 = *(*int64)(unsafe.Add(mBase, uint32(v1105)+280))
	v1313 = *(*int64)(unsafe.Add(mBase, uint32(l8)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v1105)+280)) = v1312 + v1313
	v1316 = *(*int64)(unsafe.Add(mBase, uint32(v1105)+288))
	v1317 = *(*int64)(unsafe.Add(mBase, uint32(l8)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v1105)+288)) = v1316 + v1317
	if l9 != 0 {
		goto L232
	} else {
		goto L233
	}
L222:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v1157)+88)) = l5
	*(*float64)(unsafe.Add(mBase, uint32(v1157)+72)) = l5
	*(*float64)(unsafe.Add(mBase, uint32(v1157)+56)) = l5
	goto L221
L223:
	;
	goto L224
L224:
	;
	v1169 = v1157 + int32(88)
	v1170 = *(*float64)(unsafe.Add(mBase, uint32(v1169)))
	v1171 = base.F64_sub(l5, v1170)
	v1174 = base.F64_add(v1170, base.F64_div(v1171, base.F64_convert_i64_s(v1155)))
	*(*float64)(unsafe.Add(mBase, uint32(v1169))) = v1174
	v1177 = v1157 + int32(104)
	v1180 = *(*float64)(unsafe.Add(mBase, uint32(v1177)))
	*(*float64)(unsafe.Add(mBase, uint32(v1177))) = base.F64_add(base.F64_mul(v1171, base.F64_sub(l5, v1174)), v1180)
	v1184 = v1157 + int32(56)
	v1185 = *(*float64)(unsafe.Add(mBase, uint32(v1184)))
	if base.F64_ne(v1185, float64(0)) != 0 {
		goto L225
	} else {
		goto L226
	}
L225:
	;
	if base.F64_lt(l5, v1185) != 0 {
		goto L228
	} else {
		goto L229
	}
L226:
	;
	v1192 = v1105 + l4<<(uint(int32(3))%32) + int32(72)
	v1193 = *(*float64)(unsafe.Add(mBase, uint32(v1192)))
	if base.F64_ne(v1193, float64(0)) != 0 {
		goto L225
	} else {
		goto L227
	}
L227:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v1184))) = l5
	*(*float64)(unsafe.Add(mBase, uint32(v1192))) = l5
	goto L221
L228:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v1184))) = l5
	goto L230
L229:
	;
	goto L230
L230:
	;
	v1205 = v1105 + l4<<(uint(int32(3))%32) + int32(72)
	v1206 = *(*float64)(unsafe.Add(mBase, uint32(v1205)))
	if base.F64_lt(v1206, l5) == int32(0) {
		goto L221
	} else {
		goto L231
	}
L231:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v1205))) = l5
	goto L221
L232:
	;
	v1320 = *(*int64)(unsafe.Add(mBase, uint32(v1105)+296))
	v1321 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l9))))
	*(*int64)(unsafe.Add(mBase, uint32(v1105)+296)) = v1320 + v1321
	v1324 = *(*float64)(unsafe.Add(mBase, uint32(v1105)+304))
	v1325 = *(*int64)(unsafe.Add(mBase, uint32(l9)+8))
	v1327 = float64(1e+06)
	*(*float64)(unsafe.Add(mBase, uint32(v1105)+304)) = base.F64_add(v1324, base.F64_div(base.F64_convert_i64_s(v1325), v1327))
	v1331 = *(*int64)(unsafe.Add(mBase, uint32(l9)+16))
	v1334 = base.F64_div(base.F64_convert_i64_s(v1331), v1327)
	if base.F64_ne(v1334, float64(0)) != 0 {
		goto L235
	} else {
		goto L236
	}
L233:
	;
	goto L234
L234:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1105)+424)) = int32(0)
	v1386 = *(*int64)(unsafe.Add(mBase, uint32(v1105)+376))
	*(*int64)(unsafe.Add(mBase, uint32(v1105)+376)) = v1386 + base.I64_extend_i32_s(l11)
	v1390 = *(*int64)(unsafe.Add(mBase, uint32(v1105)+384))
	*(*int64)(unsafe.Add(mBase, uint32(v1105)+384)) = v1390 + base.I64_extend_i32_s(l12)
	v1394 = v1102
	goto L26
L235:
	;
	v1337 = *(*int64)(unsafe.Add(mBase, uint32(v1105)+328))
	*(*int64)(unsafe.Add(mBase, uint32(v1105)+328)) = v1337 + int64(1)
	goto L237
L236:
	;
	goto L237
L237:
	;
	v1341 = *(*float64)(unsafe.Add(mBase, uint32(v1105)+320))
	*(*float64)(unsafe.Add(mBase, uint32(v1105)+320)) = base.F64_add(v1334, v1341)
	v1344 = *(*int64)(unsafe.Add(mBase, uint32(l9)+24))
	v1347 = base.F64_div(base.F64_convert_i64_s(v1344), float64(1e+06))
	if base.F64_ne(v1347, float64(0)) != 0 {
		goto L238
	} else {
		goto L239
	}
L238:
	;
	v1350 = *(*int64)(unsafe.Add(mBase, uint32(v1105)+312))
	*(*int64)(unsafe.Add(mBase, uint32(v1105)+312)) = v1350 + int64(1)
	goto L240
L239:
	;
	goto L240
L240:
	;
	v1354 = *(*float64)(unsafe.Add(mBase, uint32(v1105)+336))
	*(*float64)(unsafe.Add(mBase, uint32(v1105)+336)) = base.F64_add(v1347, v1354)
	v1357 = *(*int64)(unsafe.Add(mBase, uint32(l9)+32))
	v1360 = base.F64_div(base.F64_convert_i64_s(v1357), float64(1e+06))
	if base.F64_ne(v1360, float64(0)) != 0 {
		goto L241
	} else {
		goto L242
	}
L241:
	;
	v1363 = *(*int64)(unsafe.Add(mBase, uint32(v1105)+344))
	*(*int64)(unsafe.Add(mBase, uint32(v1105)+344)) = v1363 + int64(1)
	goto L243
L242:
	;
	goto L243
L243:
	;
	v1367 = *(*float64)(unsafe.Add(mBase, uint32(v1105)+352))
	*(*float64)(unsafe.Add(mBase, uint32(v1105)+352)) = base.F64_add(v1360, v1367)
	v1370 = *(*int64)(unsafe.Add(mBase, uint32(l9)+40))
	v1373 = base.F64_div(base.F64_convert_i64_s(v1370), float64(1e+06))
	if base.F64_ne(v1373, float64(0)) != 0 {
		goto L244
	} else {
		goto L245
	}
L244:
	;
	v1376 = *(*int64)(unsafe.Add(mBase, uint32(v1105)+360))
	*(*int64)(unsafe.Add(mBase, uint32(v1105)+360)) = v1376 + int64(1)
	goto L246
L245:
	;
	goto L246
L246:
	;
	v1380 = *(*float64)(unsafe.Add(mBase, uint32(v1105)+368))
	*(*float64)(unsafe.Add(mBase, uint32(v1105)+368)) = base.F64_add(v1373, v1380)
	goto L234
L247:
	;
	if v1394 == int32(0) {
		goto L2
	} else {
		goto L248
	}
L248:
	;
	F_pfree(m, v1394)
	mBase = m.M
	v1431 = m.ExcPending
	if v1431 != 0 {
		goto L24
	} else {
		goto L249
	}
L249:
	;
	goto L2
}
