package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_AddSubscriptionRelState(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64, l4 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int64
	_ = v23
	var v24 int64
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int64
	_ = v29
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	v10 = m.G0
	v12 = v10 + int32(-64)
	m.G0 = v12
	F_LockSharedObject(m, int32(_a_F_AddSubscriptionRelState_0), l0, int32(1))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return
	} else {
		v20 = F_table_open(m, int32(_a_F_AddSubscriptionRelState_1), int32(3))
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return
		} else {
			v23 = base.I64_extend_i32_u(l1)
			v24 = base.I64_extend_i32_u(l0)
			v25 = F_SearchSysCacheCopy(m, int32(68), v23, v24)
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return
			} else {
				if v25 == int32(0) {
					v29 = int64(0)
					*(*int64)(unsafe.Add(mBase, uint32(v12)+40)) = v29
					*(*int32)(unsafe.Add(mBase, uint32(v12)+60)) = int32(0)
					*(*int64)(unsafe.Add(mBase, uint32(v12)+24)) = v23
					*(*int64)(unsafe.Add(mBase, uint32(v12)+16)) = v24
					*(*int64)(unsafe.Add(mBase, uint32(v12)+32)) = base.I64_extend_i32_s(l2)
					if l3 != v29 {
						*(*int64)(unsafe.Add(mBase, uint32(v12)+40)) = l3
					} else {
						v40 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(v12)+63)) = uint8(v40)
					}
					v42 = *(*int32)(unsafe.Add(mBase, uint32(v20)+52))
					v47 = F_heap_form_tuple(m, v42, v10+int32(-48), v10+int32(-4))
					mBase = m.M
					v48 = m.ExcPending
					if v48 != 0 {
						return
					} else {
						F_CatalogTupleInsert(m, v20, v47)
						mBase = m.M
						v50 = m.ExcPending
						if v50 != 0 {
							return
						} else {
							F_pfree(m, v47)
							mBase = m.M
							v52 = m.ExcPending
							if v52 != 0 {
								return
							} else {
								if l4 != 0 {
									F_relation_close(m, v20, int32(0))
									mBase = m.M
									v55 = m.ExcPending
									if v55 != 0 {
										return
									} else {
										m.G0 = v12 - int32(-64)
										return
									}
								} else {
									F_relation_close(m, v20, int32(3))
									mBase = m.M
									v58 = m.ExcPending
									if v58 != 0 {
										return
									} else {
										F_UnlockSharedObject(m, int32(_a_F_AddSubscriptionRelState_0), l0, int32(1))
										mBase = m.M
										v62 = m.ExcPending
										if v62 != 0 {
											return
										} else {
											m.G0 = v12 - int32(-64)
											return
										}
									}
								}
							}
						}
					}
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v69 = m.ExcPending
					if v69 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = l0
						*(*int32)(unsafe.Add(mBase, uint32(v12))) = l1
						F_errmsg_internal(m, int32(_a_F_AddSubscriptionRelState_2), v12)
						mBase = m.M
						v74 = m.ExcPending
						if v74 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_AddSubscriptionRelState_3), int32(361), int32(_a_F_AddSubscriptionRelState_4))
							mBase = m.M
							v79 = m.ExcPending
							if v79 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_DropSubscription(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
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
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int64
	_ = v82
	var v97 int64
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v135 int32
	_ = v135
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v169 int32
	_ = v169
	var v186 int32
	_ = v186
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v221 int32
	_ = v221
	var v238 int32
	_ = v238
	var v255 int64
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v316 int32
	_ = v316
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v334 int32
	_ = v334
	var v350 int32
	_ = v350
	var v353 int32
	_ = v353
	var v367 int32
	_ = v367
	var v370 int32
	_ = v370
	var v386 int32
	_ = v386
	var v401 int64
	_ = v401
	var v402 int32
	_ = v402
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v434 int64
	_ = v434
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v474 int32
	_ = v474
	var v494 int32
	_ = v494
	var v497 int32
	_ = v497
	var v513 int32
	_ = v513
	var v527 int32
	_ = v527
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v563 int32
	_ = v563
	var v576 int32
	_ = v576
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v597 int32
	_ = v597
	var v599 int32
	_ = v599
	var v600 int32
	_ = v600
	var v641 int32
	_ = v641
	var v655 int32
	_ = v655
	var v668 int32
	_ = v668
	var v671 int32
	_ = v671
	var v672 int32
	_ = v672
	var v674 int32
	_ = v674
	var v676 int32
	_ = v676
	var v679 int32
	_ = v679
	var v680 int32
	_ = v680
	var v696 int32
	_ = v696
	var v709 int32
	_ = v709
	var v713 int32
	_ = v713
	var v714 int32
	_ = v714
	var v728 int32
	_ = v728
	var v730 int32
	_ = v730
	var v746 int32
	_ = v746
	var v749 int32
	_ = v749
	var v750 int32
	_ = v750
	var v792 int32
	_ = v792
	var v793 int32
	_ = v793
	var v809 int32
	_ = v809
	var v824 int32
	_ = v824
	var v839 int32
	_ = v839
	var v841 int32
	_ = v841
	var v857 int32
	_ = v857
	var v874 int32
	_ = v874
	var v892 int32
	_ = v892
	var v908 int32
	_ = v908
	var v923 int32
	_ = v923
	var v924 int32
	_ = v924
	var v939 int32
	_ = v939
	var v940 int32
	_ = v940
	var v954 int32
	_ = v954
	var v955 int32
	_ = v955
	var v956 int32
	_ = v956
	var v974 int32
	_ = v974
	var v975 int32
	_ = v975
	var v989 int32
	_ = v989
	var v990 int32
	_ = v990
	var v991 int32
	_ = v991
	var v993 int32
	_ = v993
	var v997 int32
	_ = v997
	var v998 int32
	_ = v998
	var v1011 int32
	_ = v1011
	var v1017 int32
	_ = v1017
	var v1018 int32
	_ = v1018
	var v1019 int32
	_ = v1019
	var v1020 int32
	_ = v1020
	var v1037 int32
	_ = v1037
	var v1052 int32
	_ = v1052
	var v1053 int32
	_ = v1053
	var v1056 int32
	_ = v1056
	var v1057 int32
	_ = v1057
	var v1073 int32
	_ = v1073
	var v1075 int32
	_ = v1075
	var v1086 int32
	_ = v1086
	var v1090 int32
	_ = v1090
	var v1091 int32
	_ = v1091
	var v1094 int32
	_ = v1094
	var v1097 int64
	_ = v1097
	var v1126 int32
	_ = v1126
	var v1128 int32
	_ = v1128
	var v1143 int32
	_ = v1143
	var v1144 int32
	_ = v1144
	var v1162 int32
	_ = v1162
	var v1179 int32
	_ = v1179
	var v1180 int32
	_ = v1180
	var v1181 int32
	_ = v1181
	var v1184 int32
	_ = v1184
	var v1227 int32
	_ = v1227
	var v1242 int32
	_ = v1242
	var v1261 int32
	_ = v1261
	var v1282 int32
	_ = v1282
	var v1299 int32
	_ = v1299
	var v1302 int32
	_ = v1302
	var v1304 int32
	_ = v1304
	var v1306 int32
	_ = v1306
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
	var v1324 int32
	_ = v1324
	var v1326 int32
	_ = v1326
	var v1327 int32
	_ = v1327
	var v1328 int32
	_ = v1328
	var v1345 int32
	_ = v1345
	var v1346 int32
	_ = v1346
	var v1347 int32
	_ = v1347
	var v1363 int32
	_ = v1363
	var v1376 int32
	_ = v1376
	var v1380 int32
	_ = v1380
	var v1381 int32
	_ = v1381
	var v1384 int32
	_ = v1384
	var v1387 int64
	_ = v1387
	var v1412 int32
	_ = v1412
	var v1418 int32
	_ = v1418
	var v1420 int32
	_ = v1420
	var v1435 int32
	_ = v1435
	var v1439 int32
	_ = v1439
	var v1440 int32
	_ = v1440
	var v1480 int32
	_ = v1480
	var v1483 int32
	_ = v1483
	var v1500 int32
	_ = v1500
	var v1506 int32
	_ = v1506
	var v1507 int32
	_ = v1507
	var v1521 int32
	_ = v1521
	var v1540 int32
	_ = v1540
	var v1546 int32
	_ = v1546
	var v1547 int32
	_ = v1547
	var v1552 int32
	_ = v1552
	var v1553 int32
	_ = v1553
	var v1559 int32
	_ = v1559
	var v1565 int32
	_ = v1565
	var v1579 int32
	_ = v1579
	var v1606 int32
	_ = v1606
	var v1607 int64
	_ = v1607
	var v1611 int32
	_ = v1611
	var v1613 int32
	_ = v1613
	var v1614 int32
	_ = v1614
	var v1617 int32
	_ = v1617
	var v1619 int32
	_ = v1619
	var v1621 int32
	_ = v1621
	var v1622 int32
	_ = v1622
	var v1623 int32
	_ = v1623
	var v1624 int32
	_ = v1624
	var v1625 int32
	_ = v1625
	var v1626 int32
	_ = v1626
	var v1627 int32
	_ = v1627
	var v1628 int32
	_ = v1628
	var v1629 int32
	_ = v1629
	var v1630 int32
	_ = v1630
	var v1631 int32
	_ = v1631
	var v1632 int32
	_ = v1632
	var v1633 int32
	_ = v1633
	var v1635 int32
	_ = v1635
	v3 = int32(0)
	v27 = m.G0
	v29 = v27 - int32(512)
	m.G0 = v29
	v35 = v3
	v36 = v3
	v37 = v3
	v38 = v3
	v39 = v3
	v40 = v3
	v41 = v3
	v42 = int32(-1)
	v43 = v3
	v44 = v3
	v45 = v3
	v46 = v3
	v47 = v3
	v48 = v3
	goto L2
L1:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2:
	;
	goto L5
L3:
	;
	m.G0 = v29 + int32(512)
	return
L4:
	;
	goto L3
L5:
	;
	if v42 == int32(1) {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	v1606 = int32(m.ExcTag)
	v1607 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v1606 == int32(0) {
		goto L161
	} else {
		goto L162
	}
L8:
	;
	if v1327 == int32(0) {
		goto L138
	} else {
		goto L139
	}
L9:
	;
	v1315 = v35
	v1316 = v36
	v1317 = v37
	v1318 = v38
	v1319 = v39
	v1320 = v40
	v1321 = v41
	v1322 = v45
	v1323 = v43
	v1324 = v44
	v1326 = v46
	v1327 = v47
	v1328 = v48
	goto L8
L10:
	;
	goto L11
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+464)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v29)+468)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v29)+472)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v29)+476)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v29)+484)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v29)+488)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v29)+492)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v29)+500)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v29)+504)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v29)+508)) = v37
	v70 = int32(1)
	v71 = v46 & v70
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+483)) = uint8(v71)
	v74 = v48 & v70
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+499)) = uint8(v74)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+316)) = int32(0)
	v80 = F_table_open(m, int32(_a_F_DropSubscription_0), int32(3))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L7
	} else {
		goto L12
	}
L12:
	;
	v82 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+4)))
	*(*int32)(unsafe.Add(mBase, uint32(v29)+464)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v29)+468)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v29)+472)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v29)+476)) = v43
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+483)) = uint8(v71)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+484)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v29)+488)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v29)+492)) = v45
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+499)) = uint8(v74)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+500)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v29)+504)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v29)+508)) = v80
	v97 = int64(*(*uint32)(unsafe.Add(mBase, _c_F_DropSubscription[0])))
	v98 = F_SearchSysCache2(m, int32(66), v97, v82)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L7
	} else {
		goto L13
	}
L13:
	;
	if v98 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+468)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v29)+464)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v29)+472)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v29)+476)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v29)+484)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v29)+488)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v29)+492)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v29)+500)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v29)+504)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v29)+508)) = v80
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+483)) = uint8(v71)
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+499)) = uint8(v74)
	F_relation_close(m, v80, int32(0))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L7
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+468)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v29)+464)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v29)+472)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v29)+476)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v29)+484)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v29)+488)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v29)+492)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v29)+500)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v29)+504)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v29)+508)) = v80
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+483)) = uint8(v71)
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+499)) = uint8(v74)
	v255 = F_SysCacheGetAttr(m, int32(67), v98, int32(18), v29+int32(387))
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L7
	} else {
		goto L29
	}
L17:
	;
	v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
	if v117 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+468)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v29)+464)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v29)+472)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v29)+476)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v29)+484)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v29)+488)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v29)+492)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v29)+500)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v29)+504)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v29)+508)) = v80
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+483)) = uint8(v71)
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+499)) = uint8(v74)
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L7
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+468)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v29)+464)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v29)+472)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v29)+476)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v29)+484)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v29)+488)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v29)+492)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v29)+500)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v29)+504)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v29)+508)) = v80
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+483)) = uint8(v71)
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+499)) = uint8(v74)
	v201 = F_errstart(m, int32(18), int32(0))
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L7
	} else {
		goto L25
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+468)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v29)+464)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v29)+472)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v29)+476)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v29)+484)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v29)+488)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v29)+492)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v29)+500)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v29)+504)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v29)+508)) = v80
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+483)) = uint8(v71)
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+499)) = uint8(v74)
	F_errcode(m, int32(67137668))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L7
	} else {
		goto L22
	}
L22:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v29)+468)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v29)+464)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v29)+472)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v29)+476)) = v43
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+483)) = uint8(v71)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+484)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v29)+488)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v29)+492)) = v45
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+499)) = uint8(v74)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+500)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v29)+504)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v29)+508)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v29)+16)) = v151
	F_errmsg(m, int32(_a_F_DropSubscription_1), v29+int32(16))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L7
	} else {
		goto L23
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+468)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v29)+464)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v29)+472)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v29)+476)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v29)+484)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v29)+488)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v29)+492)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v29)+500)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v29)+504)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v29)+508)) = v80
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+483)) = uint8(v71)
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+499)) = uint8(v74)
	F_errfinish(m, int32(_a_F_DropSubscription_2), int32(2445), int32(_a_F_DropSubscription_3))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L7
	} else {
		goto L24
	}
L24:
	;
	goto L1
L25:
	;
	if v201 == int32(0) {
		goto L4
	} else {
		goto L26
	}
L26:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v29)+468)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v29)+464)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v29)+472)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v29)+476)) = v43
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+483)) = uint8(v71)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+484)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v29)+488)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v29)+492)) = v45
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+499)) = uint8(v74)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+500)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v29)+504)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v29)+508)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v29))) = v205
	F_errmsg(m, int32(_a_F_DropSubscription_4), v29)
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L7
	} else {
		goto L27
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+468)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v29)+464)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v29)+472)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v29)+476)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v29)+484)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v29)+488)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v29)+492)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v29)+500)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v29)+504)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v29)+508)) = v80
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+483)) = uint8(v71)
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+499)) = uint8(v74)
	F_errfinish(m, int32(_a_F_DropSubscription_2), int32(2449), int32(_a_F_DropSubscription_3))
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L7
	} else {
		goto L28
	}
L28:
	;
	goto L4
L29:
	;
	v257 = int32(0)
	v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+387)))
	if v258 == v257 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+468)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v29)+464)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v29)+472)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v29)+476)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v29)+484)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v29)+488)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v29)+492)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v29)+500)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v29)+504)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v29)+508)) = v80
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+483)) = uint8(v71)
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+499)) = uint8(v74)
	v274 = F_text_to_cstring(m, base.I32_wrap_i64(v255))
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L7
	} else {
		goto L33
	}
L31:
	;
	v276 = v44
	v277 = v257
	goto L32
L32:
	;
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v98)+16))
	v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v278)+22)))
	v280 = v278 + v279
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v280)+104))
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v280)+80))
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v280)))
	*(*int32)(unsafe.Add(mBase, uint32(v29)+468)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v29)+464)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v29)+472)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v29)+476)) = v43
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+483)) = uint8(v71)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+484)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v29)+488)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v29)+492)) = v45
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+499)) = uint8(v74)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+500)) = v283
	*(*int32)(unsafe.Add(mBase, uint32(v29)+504)) = v276
	*(*int32)(unsafe.Add(mBase, uint32(v29)+508)) = v80
	v296 = F_superuser_arg(m, v282)
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L7
	} else {
		goto L34
	}
L33:
	;
	v276 = v274
	v277 = v274
	goto L32
L34:
	;
	v298 = int32(0)
	if v296 == v298 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v301 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v280)+89)))
	v302 = v301
	goto L37
L36:
	;
	v302 = v298
	goto L37
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+468)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v29)+464)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v29)+472)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v29)+476)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v29)+484)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v29)+488)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v29)+492)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v29)+500)) = v283
	*(*int32)(unsafe.Add(mBase, uint32(v29)+504)) = v276
	*(*int32)(unsafe.Add(mBase, uint32(v29)+508)) = v80
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+483)) = uint8(v71)
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+499)) = uint8(v74)
	v316 = *(*int32)(unsafe.Add(mBase, _c_F_DropSubscription[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v29)+468)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v29)+464)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v29)+472)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v29)+476)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v29)+484)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v29)+488)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v29)+492)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v29)+500)) = v283
	*(*int32)(unsafe.Add(mBase, uint32(v29)+504)) = v276
	*(*int32)(unsafe.Add(mBase, uint32(v29)+508)) = v80
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+483)) = uint8(v71)
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+499)) = uint8(v74)
	v330 = F_object_ownercheck(m, int32(_a_F_DropSubscription_0), v283, v316)
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L7
	} else {
		goto L38
	}
L38:
	;
	if v330 == int32(0) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v334 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v29)+468)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v29)+464)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v29)+472)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v29)+476)) = v43
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+483)) = uint8(v71)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+484)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v29)+488)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v29)+492)) = v45
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+499)) = uint8(v74)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+500)) = v283
	*(*int32)(unsafe.Add(mBase, uint32(v29)+504)) = v276
	*(*int32)(unsafe.Add(mBase, uint32(v29)+508)) = v80
	F_aclcheck_error(m, int32(2), int32(39), v334)
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L7
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	v353 = *(*int32)(unsafe.Add(mBase, _c_F_DropSubscription[2]))
	if v353 != 0 {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	goto L41
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+468)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v29)+464)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v29)+472)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v29)+476)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v29)+484)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v29)+488)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v29)+492)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v29)+500)) = v283
	*(*int32)(unsafe.Add(mBase, uint32(v29)+504)) = v276
	*(*int32)(unsafe.Add(mBase, uint32(v29)+508)) = v80
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+483)) = uint8(v71)
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+499)) = uint8(v74)
	v367 = int32(0)
	F_RunObjectDropHook(m, int32(_a_F_DropSubscription_0), v283, v367, v367)
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
		goto L7
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+468)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v29)+464)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v29)+472)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v29)+476)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v29)+484)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v29)+488)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v29)+492)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v29)+500)) = v283
	*(*int32)(unsafe.Add(mBase, uint32(v29)+504)) = v276
	*(*int32)(unsafe.Add(mBase, uint32(v29)+508)) = v80
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+483)) = uint8(v71)
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+499)) = uint8(v74)
	F_LockSharedObject(m, int32(_a_F_DropSubscription_0), v283, int32(8))
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L7
	} else {
		goto L47
	}
L46:
	;
	goto L45
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+468)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v29)+464)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v29)+472)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v29)+476)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v29)+484)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v29)+488)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v29)+492)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v29)+500)) = v283
	*(*int32)(unsafe.Add(mBase, uint32(v29)+504)) = v276
	*(*int32)(unsafe.Add(mBase, uint32(v29)+508)) = v80
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+483)) = uint8(v71)
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+499)) = uint8(v74)
	v401 = F_SysCacheGetAttrNotNull(m, int32(67), v98, int32(4))
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		goto L7
	} else {
		goto L48
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+468)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v29)+464)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v29)+472)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v29)+476)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v29)+484)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v29)+488)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v29)+492)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v29)+500)) = v283
	*(*int32)(unsafe.Add(mBase, uint32(v29)+504)) = v276
	*(*int32)(unsafe.Add(mBase, uint32(v29)+508)) = v80
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+483)) = uint8(v71)
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+499)) = uint8(v74)
	v416 = F_pstrdup(m, base.I32_wrap_i64(v401))
	mBase = m.M
	v417 = m.ExcPending
	if v417 != 0 {
		goto L7
	} else {
		goto L49
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+468)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v29)+464)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v29)+472)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v29)+476)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v29)+484)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v29)+488)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v29)+492)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v29)+500)) = v283
	*(*int32)(unsafe.Add(mBase, uint32(v29)+504)) = v276
	*(*int32)(unsafe.Add(mBase, uint32(v29)+508)) = v80
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+483)) = uint8(v71)
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+499)) = uint8(v74)
	v434 = F_SysCacheGetAttr(m, int32(67), v98, int32(19), v29+int32(387))
	mBase = m.M
	v435 = m.ExcPending
	if v435 != 0 {
		goto L7
	} else {
		goto L50
	}
L50:
	;
	v436 = int32(0)
	v437 = int32(1)
	v438 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+387)))
	if v438 != 0 {
		v472 = v436
		v474 = v437
		goto L51
	} else {
		goto L52
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+396)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+392)) = v283
	*(*int32)(unsafe.Add(mBase, uint32(v29)+388)) = int32(_a_F_DropSubscription_0)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+464)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v29)+468)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v29)+472)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v29)+476)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v29)+484)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v29)+488)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v29)+492)) = v472
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+499)) = uint8(v474)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+500)) = v283
	*(*int32)(unsafe.Add(mBase, uint32(v29)+504)) = v276
	*(*int32)(unsafe.Add(mBase, uint32(v29)+508)) = v80
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+483)) = uint8(v71)
	v494 = int32(1)
	F_EventTriggerSQLDropAddObject(m, v29+int32(388), v494, v494)
	mBase = m.M
	v497 = m.ExcPending
	if v497 != 0 {
		goto L7
	} else {
		goto L56
	}
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+468)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v29)+464)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v29)+472)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v29)+476)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v29)+484)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v29)+488)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v29)+492)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v29)+500)) = v283
	*(*int32)(unsafe.Add(mBase, uint32(v29)+504)) = v276
	*(*int32)(unsafe.Add(mBase, uint32(v29)+508)) = v80
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+483)) = uint8(v71)
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+499)) = uint8(v74)
	v452 = F_pstrdup(m, base.I32_wrap_i64(v434))
	mBase = m.M
	v453 = m.ExcPending
	if v453 != 0 {
		goto L7
	} else {
		goto L53
	}
L53:
	;
	if v452 == int32(0) {
		v472 = v436
		v474 = v437
		goto L51
	} else {
		goto L54
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+468)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v29)+464)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v29)+472)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v29)+476)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v29)+484)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v29)+488)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v29)+492)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v29)+500)) = v283
	*(*int32)(unsafe.Add(mBase, uint32(v29)+504)) = v276
	*(*int32)(unsafe.Add(mBase, uint32(v29)+508)) = v80
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+483)) = uint8(v71)
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+499)) = uint8(v74)
	F_PreventInTransactionBlock(m, l1, int32(_a_F_DropSubscription_5))
	mBase = m.M
	v471 = m.ExcPending
	if v471 != 0 {
		goto L7
	} else {
		goto L55
	}
L55:
	;
	v472 = v452
	v474 = int32(0)
	goto L51
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+468)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v29)+464)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v29)+472)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v29)+476)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v29)+484)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v29)+488)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v29)+492)) = v472
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+499)) = uint8(v474)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+500)) = v283
	*(*int32)(unsafe.Add(mBase, uint32(v29)+504)) = v276
	*(*int32)(unsafe.Add(mBase, uint32(v29)+508)) = v80
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+483)) = uint8(v71)
	F_simple_heap_delete(m, v80, v98+int32(4))
	mBase = m.M
	v513 = m.ExcPending
	if v513 != 0 {
		goto L7
	} else {
		goto L57
	}
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+468)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v29)+464)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v29)+472)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v29)+476)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v29)+484)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v29)+488)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v29)+492)) = v472
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+499)) = uint8(v474)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+500)) = v283
	*(*int32)(unsafe.Add(mBase, uint32(v29)+504)) = v276
	*(*int32)(unsafe.Add(mBase, uint32(v29)+508)) = v80
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+483)) = uint8(v71)
	F_ReleaseCatCache(m, v98)
	mBase = m.M
	v527 = m.ExcPending
	if v527 != 0 {
		goto L7
	} else {
		goto L58
	}
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+468)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v29)+464)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v29)+472)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v29)+476)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v29)+484)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v29)+488)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v29)+492)) = v472
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+499)) = uint8(v474)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+500)) = v283
	*(*int32)(unsafe.Add(mBase, uint32(v29)+504)) = v276
	*(*int32)(unsafe.Add(mBase, uint32(v29)+508)) = v80
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+483)) = uint8(v71)
	v542 = F_logicalrep_workers_find(m, v283, int32(0), int32(1))
	mBase = m.M
	v543 = m.ExcPending
	if v543 != 0 {
		goto L7
	} else {
		goto L60
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+468)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v29)+464)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v29)+472)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v29)+476)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v29)+484)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v29)+488)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v29)+492)) = v472
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+499)) = uint8(v474)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+500)) = v283
	*(*int32)(unsafe.Add(mBase, uint32(v29)+504)) = v276
	*(*int32)(unsafe.Add(mBase, uint32(v29)+508)) = v80
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+483)) = uint8(v71)
	F_list_free(m, v542)
	mBase = m.M
	v641 = m.ExcPending
	if v641 != 0 {
		goto L7
	} else {
		goto L67
	}
L60:
	;
	if v542 == int32(0) {
		goto L59
	} else {
		goto L61
	}
L61:
	;
	v546 = int32(0)
	v547 = *(*int32)(unsafe.Add(mBase, uint32(v542)+4))
	if v547 <= v546 {
		goto L59
	} else {
		goto L62
	}
L62:
	;
	v563 = v546
	goto L63
L63:
	;
	v576 = *(*int32)(unsafe.Add(mBase, uint32(v542)+12))
	v580 = *(*int32)(unsafe.Add(mBase, uint32(v576+v563<<(uint(int32(2))%32))))
	v581 = *(*int32)(unsafe.Add(mBase, uint32(v580)+36))
	v582 = *(*int32)(unsafe.Add(mBase, uint32(v580)+32))
	v583 = *(*int32)(unsafe.Add(mBase, uint32(v580)))
	*(*int32)(unsafe.Add(mBase, uint32(v29)+468)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v29)+464)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v29)+472)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v29)+476)) = v43
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+483)) = uint8(v71)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+484)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v29)+488)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v29)+492)) = v472
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+499)) = uint8(v474)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+500)) = v283
	*(*int32)(unsafe.Add(mBase, uint32(v29)+504)) = v276
	*(*int32)(unsafe.Add(mBase, uint32(v29)+508)) = v80
	F_logicalrep_worker_stop(m, v583, v582, v581)
	mBase = m.M
	v597 = m.ExcPending
	if v597 != 0 {
		goto L7
	} else {
		goto L65
	}
L64:
	;
	goto L59
L65:
	;
	v599 = v563 + int32(1)
	v600 = *(*int32)(unsafe.Add(mBase, uint32(v542)+4))
	if v599 < v600 {
		v563 = v599
		goto L63
	} else {
		goto L66
	}
L66:
	;
	goto L64
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+468)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v29)+464)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v29)+472)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v29)+476)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v29)+484)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v29)+488)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v29)+492)) = v472
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+499)) = uint8(v474)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+500)) = v283
	*(*int32)(unsafe.Add(mBase, uint32(v29)+504)) = v276
	*(*int32)(unsafe.Add(mBase, uint32(v29)+508)) = v80
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+483)) = uint8(v71)
	F_ApplyLauncherForgetWorkerStartTime(m, v283)
	mBase = m.M
	v655 = m.ExcPending
	if v655 != 0 {
		goto L7
	} else {
		goto L68
	}
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+468)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v29)+464)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v29)+472)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v29)+476)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v29)+484)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v29)+488)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v29)+492)) = v472
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+499)) = uint8(v474)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+500)) = v283
	*(*int32)(unsafe.Add(mBase, uint32(v29)+504)) = v276
	*(*int32)(unsafe.Add(mBase, uint32(v29)+508)) = v80
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+483)) = uint8(v71)
	v668 = int32(1)
	v671 = F_GetSubscriptionRelations(m, v283, v668, int32(0), v668)
	mBase = m.M
	v672 = m.ExcPending
	if v672 != 0 {
		goto L7
	} else {
		goto L69
	}
L69:
	;
	v674 = v671 + int32(4)
	v676 = base.B2i32(v671 == int32(0))
	if v671 == int32(0) {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+468)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v29)+464)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v29)+472)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v29)+476)) = v43
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+483)) = uint8(v676)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+484)) = v674
	*(*int32)(unsafe.Add(mBase, uint32(v29)+488)) = v671
	*(*int32)(unsafe.Add(mBase, uint32(v29)+492)) = v472
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+499)) = uint8(v474)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+500)) = v283
	*(*int32)(unsafe.Add(mBase, uint32(v29)+504)) = v276
	*(*int32)(unsafe.Add(mBase, uint32(v29)+508)) = v80
	v792 = F_deleteDependencyRecordsFor(m, int32(_a_F_DropSubscription_0), v283, int32(0))
	mBase = m.M
	v793 = m.ExcPending
	if v793 != 0 {
		goto L7
	} else {
		goto L81
	}
L71:
	;
	v679 = int32(0)
	v680 = *(*int32)(unsafe.Add(mBase, uint32(v671)+4))
	if v680 <= v679 {
		goto L70
	} else {
		goto L72
	}
L72:
	;
	v696 = v679
	goto L73
L73:
	;
	v709 = *(*int32)(unsafe.Add(mBase, uint32(v671)+12))
	v713 = *(*int32)(unsafe.Add(mBase, uint32(v709+v696<<(uint(int32(2))%32))))
	v714 = *(*int32)(unsafe.Add(mBase, uint32(v713)))
	if v714 != 0 {
		goto L75
	} else {
		goto L76
	}
L74:
	;
	goto L70
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+468)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v29)+464)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v29)+472)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v29)+476)) = v43
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+483)) = uint8(v676)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+484)) = v674
	*(*int32)(unsafe.Add(mBase, uint32(v29)+488)) = v671
	*(*int32)(unsafe.Add(mBase, uint32(v29)+492)) = v472
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+499)) = uint8(v474)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+500)) = v283
	*(*int32)(unsafe.Add(mBase, uint32(v29)+504)) = v276
	*(*int32)(unsafe.Add(mBase, uint32(v29)+508)) = v80
	v728 = v29 + int32(320)
	F_ReplicationOriginNameForLogicalRep(m, v283, v714, v728)
	mBase = m.M
	v730 = m.ExcPending
	if v730 != 0 {
		goto L7
	} else {
		goto L78
	}
L76:
	;
	goto L77
L77:
	;
	v749 = v696 + int32(1)
	v750 = *(*int32)(unsafe.Add(mBase, uint32(v674)))
	if v749 < v750 {
		v696 = v749
		goto L73
	} else {
		goto L80
	}
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+468)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v29)+464)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v29)+472)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v29)+476)) = v43
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+483)) = uint8(v676)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+484)) = v674
	*(*int32)(unsafe.Add(mBase, uint32(v29)+488)) = v671
	*(*int32)(unsafe.Add(mBase, uint32(v29)+492)) = v472
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+499)) = uint8(v474)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+500)) = v283
	*(*int32)(unsafe.Add(mBase, uint32(v29)+504)) = v276
	*(*int32)(unsafe.Add(mBase, uint32(v29)+508)) = v80
	F_replorigin_drop_by_name(m, v728, int32(1), int32(0))
	mBase = m.M
	v746 = m.ExcPending
	if v746 != 0 {
		goto L7
	} else {
		goto L79
	}
L79:
	;
	goto L77
L80:
	;
	goto L74
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+468)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v29)+464)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v29)+472)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v29)+476)) = v43
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+483)) = uint8(v676)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+484)) = v674
	*(*int32)(unsafe.Add(mBase, uint32(v29)+488)) = v671
	*(*int32)(unsafe.Add(mBase, uint32(v29)+492)) = v472
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+499)) = uint8(v474)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+500)) = v283
	*(*int32)(unsafe.Add(mBase, uint32(v29)+504)) = v276
	*(*int32)(unsafe.Add(mBase, uint32(v29)+508)) = v80
	F_deleteSharedDependencyRecordsFor(m, int32(_a_F_DropSubscription_0), v283, int32(0))
	mBase = m.M
	v809 = m.ExcPending
	if v809 != 0 {
		goto L7
	} else {
		goto L82
	}
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+468)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v29)+464)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v29)+472)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v29)+476)) = v43
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+483)) = uint8(v676)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+484)) = v674
	*(*int32)(unsafe.Add(mBase, uint32(v29)+488)) = v671
	*(*int32)(unsafe.Add(mBase, uint32(v29)+492)) = v472
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+499)) = uint8(v474)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+500)) = v283
	*(*int32)(unsafe.Add(mBase, uint32(v29)+504)) = v276
	*(*int32)(unsafe.Add(mBase, uint32(v29)+508)) = v80
	F_RemoveSubscriptionRel(m, v283, int32(0))
	mBase = m.M
	v824 = m.ExcPending
	if v824 != 0 {
		goto L7
	} else {
		goto L83
	}
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+468)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v29)+464)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v29)+472)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v29)+476)) = v43
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+483)) = uint8(v676)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+484)) = v674
	*(*int32)(unsafe.Add(mBase, uint32(v29)+488)) = v671
	*(*int32)(unsafe.Add(mBase, uint32(v29)+492)) = v472
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+499)) = uint8(v474)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+500)) = v283
	*(*int32)(unsafe.Add(mBase, uint32(v29)+504)) = v276
	*(*int32)(unsafe.Add(mBase, uint32(v29)+508)) = v80
	v839 = v29 + int32(320)
	F_ReplicationOriginNameForLogicalRep(m, v283, int32(0), v839)
	mBase = m.M
	v841 = m.ExcPending
	if v841 != 0 {
		goto L7
	} else {
		goto L84
	}
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+468)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v29)+464)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v29)+472)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v29)+476)) = v43
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+483)) = uint8(v676)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+484)) = v674
	*(*int32)(unsafe.Add(mBase, uint32(v29)+488)) = v671
	*(*int32)(unsafe.Add(mBase, uint32(v29)+492)) = v472
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+499)) = uint8(v474)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+500)) = v283
	*(*int32)(unsafe.Add(mBase, uint32(v29)+504)) = v276
	*(*int32)(unsafe.Add(mBase, uint32(v29)+508)) = v80
	F_replorigin_drop_by_name(m, v839, int32(1), int32(0))
	mBase = m.M
	v857 = m.ExcPending
	if v857 != 0 {
		goto L7
	} else {
		goto L85
	}
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+468)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v29)+464)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v29)+472)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v29)+476)) = v43
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+483)) = uint8(v676)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+484)) = v674
	*(*int32)(unsafe.Add(mBase, uint32(v29)+488)) = v671
	*(*int32)(unsafe.Add(mBase, uint32(v29)+492)) = v472
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+499)) = uint8(v474)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+500)) = v283
	*(*int32)(unsafe.Add(mBase, uint32(v29)+504)) = v276
	*(*int32)(unsafe.Add(mBase, uint32(v29)+508)) = v80
	F_pgstat_drop_transactional(m, int32(5), int32(0), base.I64_extend_i32_u(v283))
	mBase = m.M
	v874 = m.ExcPending
	if v874 != 0 {
		goto L7
	} else {
		goto L86
	}
L86:
	;
	if base.B2i32(v671 == int32(0))&v474 != 0 {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+468)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v29)+464)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v29)+472)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v29)+476)) = v43
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+483)) = uint8(v676)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+484)) = v674
	*(*int32)(unsafe.Add(mBase, uint32(v29)+488)) = v671
	*(*int32)(unsafe.Add(mBase, uint32(v29)+492)) = v472
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+499)) = uint8(v474)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+500)) = v283
	*(*int32)(unsafe.Add(mBase, uint32(v29)+504)) = v276
	*(*int32)(unsafe.Add(mBase, uint32(v29)+508)) = v80
	F_relation_close(m, v80, int32(0))
	mBase = m.M
	v892 = m.ExcPending
	if v892 != 0 {
		goto L7
	} else {
		goto L90
	}
L88:
	;
	goto L89
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+468)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v29)+464)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v29)+472)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v29)+476)) = v43
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+483)) = uint8(v676)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+484)) = v674
	*(*int32)(unsafe.Add(mBase, uint32(v29)+488)) = v671
	*(*int32)(unsafe.Add(mBase, uint32(v29)+492)) = v472
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+499)) = uint8(v474)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+500)) = v283
	*(*int32)(unsafe.Add(mBase, uint32(v29)+504)) = v276
	*(*int32)(unsafe.Add(mBase, uint32(v29)+508)) = v80
	F_load_file(m, int32(_a_F_DropSubscription_6), int32(0))
	mBase = m.M
	v908 = m.ExcPending
	if v908 != 0 {
		goto L7
	} else {
		goto L91
	}
L90:
	;
	goto L4
L91:
	;
	if v281 != 0 {
		goto L94
	} else {
		goto L95
	}
L92:
	;
	v1302 = *(*int32)(unsafe.Add(mBase, _c_F_DropSubscription[3]))
	v1304 = *(*int32)(unsafe.Add(mBase, _c_F_DropSubscription[4]))
	goto L134
L93:
	;
	if v474 != 0 {
		goto L108
	} else {
		goto L109
	}
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+464)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v29)+316)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+468)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v29)+472)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v29)+476)) = v43
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+483)) = uint8(v676)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+484)) = v674
	*(*int32)(unsafe.Add(mBase, uint32(v29)+488)) = v671
	*(*int32)(unsafe.Add(mBase, uint32(v29)+492)) = v472
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+499)) = uint8(v474)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+500)) = v283
	*(*int32)(unsafe.Add(mBase, uint32(v29)+504)) = v276
	*(*int32)(unsafe.Add(mBase, uint32(v29)+508)) = v80
	v923 = F_GetForeignServer(m, v281)
	mBase = m.M
	v924 = m.ExcPending
	if v924 != 0 {
		goto L7
	} else {
		goto L97
	}
L95:
	;
	v991 = v43
	v993 = v277
	goto L96
L96:
	;
	if v993 == int32(0) {
		v1019 = v38
		v1020 = v991
		goto L93
	} else {
		goto L105
	}
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+468)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v29)+464)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v29)+472)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v29)+476)) = v43
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+483)) = uint8(v676)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+484)) = v674
	*(*int32)(unsafe.Add(mBase, uint32(v29)+488)) = v671
	*(*int32)(unsafe.Add(mBase, uint32(v29)+492)) = v472
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+499)) = uint8(v474)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+500)) = v283
	*(*int32)(unsafe.Add(mBase, uint32(v29)+504)) = v276
	*(*int32)(unsafe.Add(mBase, uint32(v29)+508)) = v80
	v939 = F_object_aclcheck(m, int32(1417), v281, v282, int64(256))
	mBase = m.M
	v940 = m.ExcPending
	if v940 != 0 {
		goto L7
	} else {
		goto L98
	}
L98:
	;
	if v939 != 0 {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+468)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v29)+464)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v29)+472)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v29)+476)) = v43
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+483)) = uint8(v676)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+484)) = v674
	*(*int32)(unsafe.Add(mBase, uint32(v29)+488)) = v671
	*(*int32)(unsafe.Add(mBase, uint32(v29)+492)) = v472
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+499)) = uint8(v474)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+500)) = v283
	*(*int32)(unsafe.Add(mBase, uint32(v29)+504)) = v276
	*(*int32)(unsafe.Add(mBase, uint32(v29)+508)) = v80
	v954 = F_GetUserNameFromId(m, v282, int32(0))
	mBase = m.M
	v955 = m.ExcPending
	if v955 != 0 {
		goto L7
	} else {
		goto L102
	}
L100:
	;
	goto L101
L101:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+468)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v29)+464)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v29)+472)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v29)+476)) = v43
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+483)) = uint8(v676)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+484)) = v674
	*(*int32)(unsafe.Add(mBase, uint32(v29)+488)) = v671
	*(*int32)(unsafe.Add(mBase, uint32(v29)+492)) = v472
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+499)) = uint8(v474)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+500)) = v283
	*(*int32)(unsafe.Add(mBase, uint32(v29)+504)) = v276
	*(*int32)(unsafe.Add(mBase, uint32(v29)+508)) = v80
	v989 = F_ForeignServerConnectionString(m, v282, v923)
	mBase = m.M
	v990 = m.ExcPending
	if v990 != 0 {
		goto L7
	} else {
		goto L104
	}
L102:
	;
	v956 = *(*int32)(unsafe.Add(mBase, uint32(v923)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v29)+468)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v29)+464)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v29)+472)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v29)+476)) = v43
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+483)) = uint8(v676)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+484)) = v674
	*(*int32)(unsafe.Add(mBase, uint32(v29)+488)) = v671
	*(*int32)(unsafe.Add(mBase, uint32(v29)+492)) = v472
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+499)) = uint8(v474)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+500)) = v283
	*(*int32)(unsafe.Add(mBase, uint32(v29)+504)) = v276
	*(*int32)(unsafe.Add(mBase, uint32(v29)+508)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v29)+84)) = v956
	*(*int32)(unsafe.Add(mBase, uint32(v29)+80)) = v954
	v974 = F_psprintf(m, int32(_a_F_DropSubscription_7), v29+int32(80))
	mBase = m.M
	v975 = m.ExcPending
	if v975 != 0 {
		goto L7
	} else {
		goto L103
	}
L103:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+316)) = v974
	v1019 = v38
	v1020 = v43
	goto L93
L104:
	;
	v991 = v989
	v993 = v989
	goto L96
L105:
	;
	v997 = *(*int32)(unsafe.Add(mBase, _c_F_DropSubscription[5]))
	v998 = *(*int32)(unsafe.Add(mBase, uint32(v997)))
	*(*int32)(unsafe.Add(mBase, uint32(v29)+464)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v29)+468)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v29)+472)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v29)+476)) = v991
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+483)) = uint8(v676)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+484)) = v674
	*(*int32)(unsafe.Add(mBase, uint32(v29)+488)) = v671
	*(*int32)(unsafe.Add(mBase, uint32(v29)+492)) = v472
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+499)) = uint8(v474)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+500)) = v283
	*(*int32)(unsafe.Add(mBase, uint32(v29)+504)) = v276
	*(*int32)(unsafe.Add(mBase, uint32(v29)+508)) = v80
	v1011 = int32(1)
	v1017 = m.T0[v998].(func(*base.Module, int32, int32, int32, int32, int32, int32) int32)(m, v993, v1011, v1011, v302&v1011, v416, v29+int32(316))
	mBase = m.M
	v1018 = m.ExcPending
	if v1018 != 0 {
		goto L7
	} else {
		goto L106
	}
L106:
	;
	if v1017 != 0 {
		goto L92
	} else {
		goto L107
	}
L107:
	;
	v1019 = v1017
	v1020 = v991
	goto L93
L108:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+468)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v29)+464)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v29)+472)) = v1019
	*(*int32)(unsafe.Add(mBase, uint32(v29)+476)) = v1020
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+483)) = uint8(v676)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+484)) = v674
	*(*int32)(unsafe.Add(mBase, uint32(v29)+488)) = v671
	*(*int32)(unsafe.Add(mBase, uint32(v29)+492)) = v472
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+499)) = uint8(v474)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+500)) = v283
	*(*int32)(unsafe.Add(mBase, uint32(v29)+504)) = v276
	*(*int32)(unsafe.Add(mBase, uint32(v29)+508)) = v80
	F_list_free(m, v671)
	mBase = m.M
	v1037 = m.ExcPending
	if v1037 != 0 {
		goto L7
	} else {
		goto L111
	}
L109:
	;
	goto L110
L110:
	;
	v1053 = *(*int32)(unsafe.Add(mBase, uint32(v29)+316))
	if v671 == int32(0) {
		goto L113
	} else {
		goto L114
	}
L111:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+468)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v29)+464)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v29)+472)) = v1019
	*(*int32)(unsafe.Add(mBase, uint32(v29)+476)) = v1020
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+483)) = uint8(v676)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+484)) = v674
	*(*int32)(unsafe.Add(mBase, uint32(v29)+488)) = v671
	*(*int32)(unsafe.Add(mBase, uint32(v29)+492)) = v472
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+499)) = uint8(v474)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+500)) = v283
	*(*int32)(unsafe.Add(mBase, uint32(v29)+504)) = v276
	*(*int32)(unsafe.Add(mBase, uint32(v29)+508)) = v80
	F_relation_close(m, v80, int32(0))
	mBase = m.M
	v1052 = m.ExcPending
	if v1052 != 0 {
		goto L7
	} else {
		goto L112
	}
L112:
	;
	goto L4
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+468)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v29)+464)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v29)+472)) = v1019
	*(*int32)(unsafe.Add(mBase, uint32(v29)+476)) = v1020
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+483)) = uint8(v676)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+484)) = v674
	*(*int32)(unsafe.Add(mBase, uint32(v29)+488)) = v671
	*(*int32)(unsafe.Add(mBase, uint32(v29)+492)) = v472
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+499)) = uint8(v474)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+500)) = v283
	*(*int32)(unsafe.Add(mBase, uint32(v29)+504)) = v276
	*(*int32)(unsafe.Add(mBase, uint32(v29)+508)) = v80
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1227 = m.ExcPending
	if v1227 != 0 {
		goto L7
	} else {
		goto L129
	}
L114:
	;
	v1056 = int32(0)
	v1057 = *(*int32)(unsafe.Add(mBase, uint32(v671)+4))
	if v1057 <= v1056 {
		goto L113
	} else {
		goto L115
	}
L115:
	;
	v1073 = v1056
	v1075 = v1057
	goto L116
L116:
	;
	v1086 = *(*int32)(unsafe.Add(mBase, uint32(v671)+12))
	v1090 = *(*int32)(unsafe.Add(mBase, uint32(v1086+v1073<<(uint(int32(2))%32))))
	v1091 = *(*int32)(unsafe.Add(mBase, uint32(v1090)))
	if v1091 == int32(0) {
		v1181 = v1075
		goto L118
	} else {
		goto L119
	}
L117:
	;
	goto L113
L118:
	;
	v1184 = v1073 + int32(1)
	if v1184 < v1181 {
		v1073 = v1184
		v1075 = v1181
		goto L116
	} else {
		goto L128
	}
L119:
	;
	v1094 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1090)+16)))
	if v1094 == int32(115) {
		v1181 = v1075
		goto L118
	} else {
		goto L120
	}
L120:
	;
	v1097 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v29)+456)) = v1097
	*(*int64)(unsafe.Add(mBase, uint32(v29)+448)) = v1097
	*(*int64)(unsafe.Add(mBase, uint32(v29)+440)) = v1097
	*(*int64)(unsafe.Add(mBase, uint32(v29)+432)) = v1097
	*(*int64)(unsafe.Add(mBase, uint32(v29)+424)) = v1097
	*(*int64)(unsafe.Add(mBase, uint32(v29)+416)) = v1097
	*(*int64)(unsafe.Add(mBase, uint32(v29)+408)) = v1097
	*(*int64)(unsafe.Add(mBase, uint32(v29)+400)) = v1097
	*(*int32)(unsafe.Add(mBase, uint32(v29)+464)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v29)+468)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v29)+472)) = v1019
	*(*int32)(unsafe.Add(mBase, uint32(v29)+476)) = v1020
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+483)) = uint8(v676)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+484)) = v674
	*(*int32)(unsafe.Add(mBase, uint32(v29)+488)) = v671
	*(*int32)(unsafe.Add(mBase, uint32(v29)+492)) = v472
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+499)) = uint8(v474)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+504)) = v276
	*(*int32)(unsafe.Add(mBase, uint32(v29)+500)) = v283
	*(*int32)(unsafe.Add(mBase, uint32(v29)+508)) = v80
	v1126 = v29 + int32(400)
	F_ReplicationSlotNameForTablesync(m, v283, v1091, v1126)
	mBase = m.M
	v1128 = m.ExcPending
	if v1128 != 0 {
		goto L7
	} else {
		goto L121
	}
L121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+468)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v29)+464)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v29)+472)) = v1019
	*(*int32)(unsafe.Add(mBase, uint32(v29)+476)) = v1020
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+483)) = uint8(v676)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+484)) = v674
	*(*int32)(unsafe.Add(mBase, uint32(v29)+488)) = v671
	*(*int32)(unsafe.Add(mBase, uint32(v29)+492)) = v472
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+499)) = uint8(v474)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+500)) = v283
	*(*int32)(unsafe.Add(mBase, uint32(v29)+504)) = v276
	*(*int32)(unsafe.Add(mBase, uint32(v29)+508)) = v80
	v1143 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v1144 = m.ExcPending
	if v1144 != 0 {
		goto L7
	} else {
		goto L122
	}
L122:
	;
	if v1143 != 0 {
		goto L123
	} else {
		goto L124
	}
L123:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+468)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v29)+464)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v29)+472)) = v1019
	*(*int32)(unsafe.Add(mBase, uint32(v29)+476)) = v1020
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+483)) = uint8(v676)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+484)) = v674
	*(*int32)(unsafe.Add(mBase, uint32(v29)+488)) = v671
	*(*int32)(unsafe.Add(mBase, uint32(v29)+492)) = v472
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+499)) = uint8(v474)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+500)) = v283
	*(*int32)(unsafe.Add(mBase, uint32(v29)+504)) = v276
	*(*int32)(unsafe.Add(mBase, uint32(v29)+508)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v29)+64)) = v1126
	F_errmsg_internal(m, int32(_a_F_DropSubscription_8), v29-int32(-64))
	mBase = m.M
	v1162 = m.ExcPending
	if v1162 != 0 {
		goto L7
	} else {
		goto L126
	}
L124:
	;
	goto L125
L125:
	;
	v1180 = *(*int32)(unsafe.Add(mBase, uint32(v674)))
	v1181 = v1180
	goto L118
L126:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+468)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v29)+464)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v29)+472)) = v1019
	*(*int32)(unsafe.Add(mBase, uint32(v29)+476)) = v1020
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+483)) = uint8(v676)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+484)) = v674
	*(*int32)(unsafe.Add(mBase, uint32(v29)+488)) = v671
	*(*int32)(unsafe.Add(mBase, uint32(v29)+492)) = v472
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+499)) = uint8(v474)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+500)) = v283
	*(*int32)(unsafe.Add(mBase, uint32(v29)+504)) = v276
	*(*int32)(unsafe.Add(mBase, uint32(v29)+508)) = v80
	F_errfinish(m, int32(_a_F_DropSubscription_2), int32(3491), int32(_a_F_DropSubscription_9))
	mBase = m.M
	v1179 = m.ExcPending
	if v1179 != 0 {
		goto L7
	} else {
		goto L127
	}
L127:
	;
	goto L125
L128:
	;
	goto L117
L129:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+468)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v29)+464)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v29)+472)) = v1019
	*(*int32)(unsafe.Add(mBase, uint32(v29)+476)) = v1020
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+483)) = uint8(v676)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+484)) = v674
	*(*int32)(unsafe.Add(mBase, uint32(v29)+488)) = v671
	*(*int32)(unsafe.Add(mBase, uint32(v29)+492)) = v472
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+499)) = uint8(v474)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+500)) = v283
	*(*int32)(unsafe.Add(mBase, uint32(v29)+504)) = v276
	*(*int32)(unsafe.Add(mBase, uint32(v29)+508)) = v80
	F_errcode(m, int32(100663808))
	mBase = m.M
	v1242 = m.ExcPending
	if v1242 != 0 {
		goto L7
	} else {
		goto L130
	}
L130:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+468)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v29)+464)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v29)+472)) = v1019
	*(*int32)(unsafe.Add(mBase, uint32(v29)+476)) = v1020
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+483)) = uint8(v676)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+484)) = v674
	*(*int32)(unsafe.Add(mBase, uint32(v29)+488)) = v671
	*(*int32)(unsafe.Add(mBase, uint32(v29)+492)) = v472
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+499)) = uint8(v474)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+500)) = v283
	*(*int32)(unsafe.Add(mBase, uint32(v29)+504)) = v276
	*(*int32)(unsafe.Add(mBase, uint32(v29)+508)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v29)+52)) = v1053
	*(*int32)(unsafe.Add(mBase, uint32(v29)+48)) = v472
	F_errmsg(m, int32(_a_F_DropSubscription_10), v29+int32(48))
	mBase = m.M
	v1261 = m.ExcPending
	if v1261 != 0 {
		goto L7
	} else {
		goto L131
	}
L131:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+468)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v29)+464)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v29)+472)) = v1019
	*(*int32)(unsafe.Add(mBase, uint32(v29)+476)) = v1020
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+483)) = uint8(v676)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+484)) = v674
	*(*int32)(unsafe.Add(mBase, uint32(v29)+488)) = v671
	*(*int32)(unsafe.Add(mBase, uint32(v29)+492)) = v472
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+499)) = uint8(v474)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+500)) = v283
	*(*int32)(unsafe.Add(mBase, uint32(v29)+504)) = v276
	*(*int32)(unsafe.Add(mBase, uint32(v29)+508)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v29)+36)) = int32(_a_F_DropSubscription_11)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+32)) = int32(_a_F_DropSubscription_12)
	F_errhint(m, int32(_a_F_DropSubscription_13), v29+int32(32))
	mBase = m.M
	v1282 = m.ExcPending
	if v1282 != 0 {
		goto L7
	} else {
		goto L132
	}
L132:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+468)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v29)+464)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v29)+472)) = v1019
	*(*int32)(unsafe.Add(mBase, uint32(v29)+476)) = v1020
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+483)) = uint8(v676)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+484)) = v674
	*(*int32)(unsafe.Add(mBase, uint32(v29)+488)) = v671
	*(*int32)(unsafe.Add(mBase, uint32(v29)+492)) = v472
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+499)) = uint8(v474)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+500)) = v283
	*(*int32)(unsafe.Add(mBase, uint32(v29)+504)) = v276
	*(*int32)(unsafe.Add(mBase, uint32(v29)+508)) = v80
	F_errfinish(m, int32(_a_F_DropSubscription_2), int32(3502), int32(_a_F_DropSubscription_9))
	mBase = m.M
	v1299 = m.ExcPending
	if v1299 != 0 {
		goto L7
	} else {
		goto L133
	}
L133:
	;
	goto L1
L134:
	;
	v1306 = v29 + int32(160)
	*(*int32)(unsafe.Add(mBase, uint32(v1306)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1306))) = v29 + int32(92)
	goto L137
L135:
	;
	v1315 = v283
	v1316 = v671
	v1317 = v80
	v1318 = v1017
	v1319 = v674
	v1320 = v1302
	v1321 = v1304
	v1322 = v472
	v1323 = v991
	v1324 = v276
	v1326 = v676
	v1327 = int32(0)
	v1328 = v474
	goto L8
L137:
	;
	goto L135
L138:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_DropSubscription[4])) = v29 + int32(160)
	v1345 = v1326 & int32(1)
	if v1345 != 0 {
		goto L141
	} else {
		goto L142
	}
L139:
	;
	goto L140
L140:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_DropSubscription[3])) = v1320
	*(*int32)(unsafe.Add(mBase, _c_F_DropSubscription[4])) = v1321
	v1546 = *(*int32)(unsafe.Add(mBase, _c_F_DropSubscription[5]))
	v1547 = *(*int32)(unsafe.Add(mBase, uint32(v1546)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v29)+464)) = v1320
	*(*int32)(unsafe.Add(mBase, uint32(v29)+468)) = v1321
	*(*int32)(unsafe.Add(mBase, uint32(v29)+472)) = v1318
	*(*int32)(unsafe.Add(mBase, uint32(v29)+476)) = v1323
	v1552 = int32(1)
	v1553 = v1326 & v1552
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+483)) = uint8(v1553)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+484)) = v1319
	*(*int32)(unsafe.Add(mBase, uint32(v29)+488)) = v1316
	*(*int32)(unsafe.Add(mBase, uint32(v29)+492)) = v1322
	v1559 = v1328 & v1552
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+499)) = uint8(v1559)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+500)) = v1315
	*(*int32)(unsafe.Add(mBase, uint32(v29)+504)) = v1324
	*(*int32)(unsafe.Add(mBase, uint32(v29)+508)) = v1317
	m.T0[v1547].(func(*base.Module, int32))(m, v1318)
	mBase = m.M
	v1565 = m.ExcPending
	if v1565 != 0 {
		goto L7
	} else {
		goto L159
	}
L141:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+468)) = v1321
	*(*int32)(unsafe.Add(mBase, uint32(v29)+464)) = v1320
	*(*int32)(unsafe.Add(mBase, uint32(v29)+472)) = v1318
	*(*int32)(unsafe.Add(mBase, uint32(v29)+476)) = v1323
	*(*int32)(unsafe.Add(mBase, uint32(v29)+484)) = v1319
	*(*int32)(unsafe.Add(mBase, uint32(v29)+488)) = v1316
	*(*int32)(unsafe.Add(mBase, uint32(v29)+492)) = v1322
	*(*int32)(unsafe.Add(mBase, uint32(v29)+500)) = v1315
	*(*int32)(unsafe.Add(mBase, uint32(v29)+504)) = v1324
	*(*int32)(unsafe.Add(mBase, uint32(v29)+508)) = v1317
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+483)) = uint8(v1345)
	v1480 = v1328 & int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+499)) = uint8(v1480)
	F_list_free(m, v1316)
	mBase = m.M
	v1483 = m.ExcPending
	if v1483 != 0 {
		goto L7
	} else {
		goto L152
	}
L142:
	;
	v1346 = int32(0)
	v1347 = *(*int32)(unsafe.Add(mBase, uint32(v1319)))
	if v1347 <= v1346 {
		goto L141
	} else {
		goto L143
	}
L143:
	;
	v1363 = v1346
	goto L144
L144:
	;
	v1376 = *(*int32)(unsafe.Add(mBase, uint32(v1316)+12))
	v1380 = *(*int32)(unsafe.Add(mBase, uint32(v1376+v1363<<(uint(int32(2))%32))))
	v1381 = *(*int32)(unsafe.Add(mBase, uint32(v1380)))
	if v1381 == int32(0) {
		goto L146
	} else {
		goto L147
	}
L145:
	;
	goto L141
L146:
	;
	v1439 = v1363 + int32(1)
	v1440 = *(*int32)(unsafe.Add(mBase, uint32(v1319)))
	if v1439 < v1440 {
		v1363 = v1439
		goto L144
	} else {
		goto L151
	}
L147:
	;
	v1384 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1380)+16)))
	if v1384 == int32(115) {
		goto L146
	} else {
		goto L148
	}
L148:
	;
	v1387 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v29)+152)) = v1387
	*(*int64)(unsafe.Add(mBase, uint32(v29)+144)) = v1387
	*(*int64)(unsafe.Add(mBase, uint32(v29)+136)) = v1387
	*(*int64)(unsafe.Add(mBase, uint32(v29)+128)) = v1387
	*(*int64)(unsafe.Add(mBase, uint32(v29)+120)) = v1387
	*(*int64)(unsafe.Add(mBase, uint32(v29)+112)) = v1387
	*(*int64)(unsafe.Add(mBase, uint32(v29)+104)) = v1387
	*(*int64)(unsafe.Add(mBase, uint32(v29)+96)) = v1387
	*(*int32)(unsafe.Add(mBase, uint32(v29)+464)) = v1320
	*(*int32)(unsafe.Add(mBase, uint32(v29)+468)) = v1321
	*(*int32)(unsafe.Add(mBase, uint32(v29)+472)) = v1318
	*(*int32)(unsafe.Add(mBase, uint32(v29)+476)) = v1323
	*(*int32)(unsafe.Add(mBase, uint32(v29)+484)) = v1319
	*(*int32)(unsafe.Add(mBase, uint32(v29)+488)) = v1316
	*(*int32)(unsafe.Add(mBase, uint32(v29)+492)) = v1322
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+483)) = uint8(v1345)
	v1412 = v1328 & int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+499)) = uint8(v1412)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+504)) = v1324
	*(*int32)(unsafe.Add(mBase, uint32(v29)+500)) = v1315
	*(*int32)(unsafe.Add(mBase, uint32(v29)+508)) = v1317
	v1418 = v29 + int32(96)
	F_ReplicationSlotNameForTablesync(m, v1315, v1381, v1418)
	mBase = m.M
	v1420 = m.ExcPending
	if v1420 != 0 {
		goto L7
	} else {
		goto L149
	}
L149:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+468)) = v1321
	*(*int32)(unsafe.Add(mBase, uint32(v29)+464)) = v1320
	*(*int32)(unsafe.Add(mBase, uint32(v29)+472)) = v1318
	*(*int32)(unsafe.Add(mBase, uint32(v29)+476)) = v1323
	*(*int32)(unsafe.Add(mBase, uint32(v29)+484)) = v1319
	*(*int32)(unsafe.Add(mBase, uint32(v29)+488)) = v1316
	*(*int32)(unsafe.Add(mBase, uint32(v29)+492)) = v1322
	*(*int32)(unsafe.Add(mBase, uint32(v29)+500)) = v1315
	*(*int32)(unsafe.Add(mBase, uint32(v29)+504)) = v1324
	*(*int32)(unsafe.Add(mBase, uint32(v29)+508)) = v1317
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+483)) = uint8(v1345)
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+499)) = uint8(v1412)
	F_ReplicationSlotDropAtPubNode(m, v1318, v1418, int32(1))
	mBase = m.M
	v1435 = m.ExcPending
	if v1435 != 0 {
		goto L7
	} else {
		goto L150
	}
L150:
	;
	goto L146
L151:
	;
	goto L145
L152:
	;
	if v1480 == int32(0) {
		goto L153
	} else {
		goto L154
	}
L153:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+468)) = v1321
	*(*int32)(unsafe.Add(mBase, uint32(v29)+464)) = v1320
	*(*int32)(unsafe.Add(mBase, uint32(v29)+472)) = v1318
	*(*int32)(unsafe.Add(mBase, uint32(v29)+476)) = v1323
	*(*int32)(unsafe.Add(mBase, uint32(v29)+484)) = v1319
	*(*int32)(unsafe.Add(mBase, uint32(v29)+488)) = v1316
	*(*int32)(unsafe.Add(mBase, uint32(v29)+492)) = v1322
	*(*int32)(unsafe.Add(mBase, uint32(v29)+500)) = v1315
	*(*int32)(unsafe.Add(mBase, uint32(v29)+504)) = v1324
	*(*int32)(unsafe.Add(mBase, uint32(v29)+508)) = v1317
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+483)) = uint8(v1345)
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+499)) = uint8(v1480)
	F_ReplicationSlotDropAtPubNode(m, v1318, v1322, int32(0))
	mBase = m.M
	v1500 = m.ExcPending
	if v1500 != 0 {
		goto L7
	} else {
		goto L156
	}
L154:
	;
	goto L155
L155:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_DropSubscription[3])) = v1320
	*(*int32)(unsafe.Add(mBase, _c_F_DropSubscription[4])) = v1321
	v1506 = *(*int32)(unsafe.Add(mBase, _c_F_DropSubscription[5]))
	v1507 = *(*int32)(unsafe.Add(mBase, uint32(v1506)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v29)+464)) = v1320
	*(*int32)(unsafe.Add(mBase, uint32(v29)+468)) = v1321
	*(*int32)(unsafe.Add(mBase, uint32(v29)+472)) = v1318
	*(*int32)(unsafe.Add(mBase, uint32(v29)+476)) = v1323
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+483)) = uint8(v1345)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+484)) = v1319
	*(*int32)(unsafe.Add(mBase, uint32(v29)+488)) = v1316
	*(*int32)(unsafe.Add(mBase, uint32(v29)+492)) = v1322
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+499)) = uint8(v1480)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+500)) = v1315
	*(*int32)(unsafe.Add(mBase, uint32(v29)+504)) = v1324
	*(*int32)(unsafe.Add(mBase, uint32(v29)+508)) = v1317
	m.T0[v1507].(func(*base.Module, int32))(m, v1318)
	mBase = m.M
	v1521 = m.ExcPending
	if v1521 != 0 {
		goto L7
	} else {
		goto L157
	}
L156:
	;
	goto L155
L157:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_DropSubscription[3])) = v1320
	*(*int32)(unsafe.Add(mBase, _c_F_DropSubscription[4])) = v1321
	*(*int32)(unsafe.Add(mBase, uint32(v29)+464)) = v1320
	*(*int32)(unsafe.Add(mBase, uint32(v29)+468)) = v1321
	*(*int32)(unsafe.Add(mBase, uint32(v29)+472)) = v1318
	*(*int32)(unsafe.Add(mBase, uint32(v29)+476)) = v1323
	*(*int32)(unsafe.Add(mBase, uint32(v29)+484)) = v1319
	*(*int32)(unsafe.Add(mBase, uint32(v29)+488)) = v1316
	*(*int32)(unsafe.Add(mBase, uint32(v29)+492)) = v1322
	*(*int32)(unsafe.Add(mBase, uint32(v29)+500)) = v1315
	*(*int32)(unsafe.Add(mBase, uint32(v29)+504)) = v1324
	*(*int32)(unsafe.Add(mBase, uint32(v29)+508)) = v1317
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+483)) = uint8(v1345)
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+499)) = uint8(v1480)
	F_relation_close(m, v1317, int32(0))
	mBase = m.M
	v1540 = m.ExcPending
	if v1540 != 0 {
		goto L7
	} else {
		goto L158
	}
L158:
	;
	goto L4
L159:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+468)) = v1321
	*(*int32)(unsafe.Add(mBase, uint32(v29)+464)) = v1320
	*(*int32)(unsafe.Add(mBase, uint32(v29)+472)) = v1318
	*(*int32)(unsafe.Add(mBase, uint32(v29)+476)) = v1323
	*(*int32)(unsafe.Add(mBase, uint32(v29)+484)) = v1319
	*(*int32)(unsafe.Add(mBase, uint32(v29)+488)) = v1316
	*(*int32)(unsafe.Add(mBase, uint32(v29)+492)) = v1322
	*(*int32)(unsafe.Add(mBase, uint32(v29)+500)) = v1315
	*(*int32)(unsafe.Add(mBase, uint32(v29)+504)) = v1324
	*(*int32)(unsafe.Add(mBase, uint32(v29)+508)) = v1317
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+483)) = uint8(v1553)
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+499)) = uint8(v1559)
	F_pg_re_throw(m)
	mBase = m.M
	v1579 = m.ExcPending
	if v1579 != 0 {
		goto L7
	} else {
		goto L160
	}
L160:
	;
	goto L1
L161:
	;
	v1611 = int32(v1607)
	m.G0 = v29
	v1613 = *(*int32)(unsafe.Add(mBase, uint32(v1611)+4))
	v1614 = *(*int32)(unsafe.Add(mBase, uint32(v1611)))
	v1617 = *(*int32)(unsafe.Add(mBase, uint32(v1614)))
	if v29+int32(92) == v1617 {
		goto L164
	} else {
		goto L165
	}
L162:
	;
	m.ExcPending = 1
	goto L170
L163:
	;
	if v1621 != 0 {
		goto L167
	} else {
		goto L168
	}
L164:
	;
	v1619 = *(*int32)(unsafe.Add(mBase, uint32(v1614)+4))
	v1621 = v1619
	goto L166
L165:
	;
	v1621 = int32(0)
	goto L166
L166:
	;
	goto L163
L167:
	;
	v1622 = *(*int32)(unsafe.Add(mBase, uint32(v29)+508))
	v1623 = *(*int32)(unsafe.Add(mBase, uint32(v29)+504))
	v1624 = *(*int32)(unsafe.Add(mBase, uint32(v29)+500))
	v1625 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+499)))
	v1626 = *(*int32)(unsafe.Add(mBase, uint32(v29)+492))
	v1627 = *(*int32)(unsafe.Add(mBase, uint32(v29)+488))
	v1628 = *(*int32)(unsafe.Add(mBase, uint32(v29)+484))
	v1629 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+483)))
	v1630 = *(*int32)(unsafe.Add(mBase, uint32(v29)+476))
	v1631 = *(*int32)(unsafe.Add(mBase, uint32(v29)+472))
	v1632 = *(*int32)(unsafe.Add(mBase, uint32(v29)+468))
	v1633 = *(*int32)(unsafe.Add(mBase, uint32(v29)+464))
	v35 = v1624
	v36 = v1627
	v37 = v1622
	v38 = v1631
	v39 = v1628
	v40 = v1633
	v41 = v1632
	v42 = v1621
	v43 = v1630
	v44 = v1623
	v45 = v1626
	v46 = v1629
	v47 = v1613
	v48 = v1625
	goto L2
L168:
	;
	goto L169
L169:
	;
	F___wasm_longjmp(m, v1614, v1613)
	mBase = m.M
	v1635 = m.ExcPending
	if v1635 != 0 {
		goto L170
	} else {
		goto L171
	}
L170:
	;
	return
L171:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_GetSubscriptionRelations(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
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
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v89 int64
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int64
	_ = v92
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	v5 = int32(0)
	v11 = m.G0
	v13 = v11 - int32(128)
	m.G0 = v13
	v17 = F_table_open(m, int32(_a_F_GetSubscriptionRelations_0), int32(1))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	F_ScanKeyInit(m, v13+int32(16), int32(1), int32(3), int32(184), base.I64_extend_i32_u(l0))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v29 = int32(0)
	if l3 != 0 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L1
	} else {
		goto L36
	}
L5:
	;
	v34 = int32(3)
	F_ScanKeyInit(m, v13+int32(72), v34, v34, int32(70), int64(114))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L1
	} else {
		goto L8
	}
L6:
	;
	v42 = int32(1)
	goto L7
L7:
	;
	v45 = F_systable_beginscan(m, v17, v29, v29, v29, v42, v13+int32(16))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L1
	} else {
		goto L9
	}
L8:
	;
	v42 = int32(2)
	goto L7
L9:
	;
	v47 = F_systable_getnext(m, v45)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	if v47 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v52 = v47
	v58 = v5
	goto L14
L12:
	;
	v109 = v5
	goto L13
L13:
	;
	F_systable_endscan(m, v45)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L1
	} else {
		goto L34
	}
L14:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v52)+16))
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+22)))
	v61 = v59 + v60
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)+4))
	v63 = F_get_rel_relkind(m, v62)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L1
	} else {
		goto L17
	}
L15:
	;
	v109 = v97
	goto L13
L16:
	;
	v98 = F_systable_getnext(m, v45)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L1
	} else {
		goto L32
	}
L17:
	;
	if v63 == int32(0) {
		v97 = v58
		goto L16
	} else {
		goto L18
	}
L18:
	;
	if v63&int32(255) == int32(83) {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v78 = F_palloc(m, int32(24))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L1
	} else {
		goto L26
	}
L20:
	;
	if l2 != 0 {
		goto L19
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	if v63&int32(-3) != int32(112) {
		goto L4
	} else {
		goto L24
	}
L23:
	;
	v97 = v58
	goto L16
L24:
	;
	if l1 == int32(0) {
		v97 = v58
		goto L16
	} else {
		goto L25
	}
L25:
	;
	goto L19
L26:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v61)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v78))) = v80
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61)+8)))
	*(*uint8)(unsafe.Add(mBase, uint32(v78)+16)) = uint8(v82)
	v89 = F_SysCacheGetAttr(m, int32(68), v52, int32(4), v13+int32(15))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+15)))
	if v91 != 0 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v92 = int64(0)
	goto L30
L29:
	;
	v92 = v89
	goto L30
L30:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v78)+8)) = v92
	v94 = F_lappend(m, v58, v78)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	v97 = v94
	goto L16
L32:
	;
	if v98 != 0 {
		v52 = v98
		v58 = v97
		goto L14
	} else {
		goto L33
	}
L33:
	;
	goto L15
L34:
	;
	F_relation_close(m, v17, int32(1))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	m.G0 = v13 + int32(128)
	return v109
L36:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v61)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v123
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v63
	F_errmsg_internal(m, int32(_a_F_GetSubscriptionRelations_1), v13)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	F_errfinish(m, int32(_a_F_GetSubscriptionRelations_2), int32(705), int32(_a_F_GetSubscriptionRelations_3))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_parse_subscription_options(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v29 int64
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v134 int32
	_ = v134
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
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
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v292 int32
	_ = v292
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v314 int32
	_ = v314
	var v317 int32
	_ = v317
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v328 int32
	_ = v328
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v352 int32
	_ = v352
	var v355 int32
	_ = v355
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
	var v376 int32
	_ = v376
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v391 int32
	_ = v391
	var v394 int32
	_ = v394
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v405 int32
	_ = v405
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v415 int32
	_ = v415
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v427 int32
	_ = v427
	var v430 int32
	_ = v430
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v437 int32
	_ = v437
	var v440 int32
	_ = v440
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v451 int32
	_ = v451
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v461 int32
	_ = v461
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v476 int32
	_ = v476
	var v479 int32
	_ = v479
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v490 int32
	_ = v490
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v500 int32
	_ = v500
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v515 int32
	_ = v515
	var v518 int32
	_ = v518
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v529 int32
	_ = v529
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v539 int32
	_ = v539
	var v545 int32
	_ = v545
	var v547 int32
	_ = v547
	var v549 int32
	_ = v549
	var v552 int32
	_ = v552
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v578 int32
	_ = v578
	var v587 int32
	_ = v587
	var v590 int32
	_ = v590
	var v592 int32
	_ = v592
	var v601 int32
	_ = v601
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v622 int32
	_ = v622
	var v631 int32
	_ = v631
	var v634 int32
	_ = v634
	var v636 int32
	_ = v636
	var v645 int32
	_ = v645
	var v651 int32
	_ = v651
	var v652 int32
	_ = v652
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v666 int32
	_ = v666
	var v675 int32
	_ = v675
	var v678 int32
	_ = v678
	var v680 int32
	_ = v680
	var v689 int32
	_ = v689
	var v696 int32
	_ = v696
	var v697 int32
	_ = v697
	var v700 int32
	_ = v700
	var v701 int32
	_ = v701
	var v711 int32
	_ = v711
	var v720 int32
	_ = v720
	var v723 int32
	_ = v723
	var v725 int32
	_ = v725
	var v734 int32
	_ = v734
	var v740 int32
	_ = v740
	var v741 int32
	_ = v741
	var v744 int32
	_ = v744
	var v745 int32
	_ = v745
	var v755 int32
	_ = v755
	var v764 int32
	_ = v764
	var v767 int32
	_ = v767
	var v769 int32
	_ = v769
	var v778 int32
	_ = v778
	var v785 int32
	_ = v785
	var v788 int32
	_ = v788
	var v789 int32
	_ = v789
	var v793 int32
	_ = v793
	var v798 int32
	_ = v798
	var v802 int32
	_ = v802
	var v810 int32
	_ = v810
	var v811 int32
	_ = v811
	var v814 int32
	_ = v814
	var v817 int32
	_ = v817
	var v820 int32
	_ = v820
	var v821 int32
	_ = v821
	var v824 int32
	_ = v824
	var v825 int32
	_ = v825
	var v828 int32
	_ = v828
	var v835 int32
	_ = v835
	var v836 int32
	_ = v836
	var v838 int32
	_ = v838
	var v844 int32
	_ = v844
	var v845 int32
	_ = v845
	var v849 int32
	_ = v849
	var v850 int32
	_ = v850
	var v853 int32
	_ = v853
	var v856 int32
	_ = v856
	var v859 int32
	_ = v859
	var v860 int32
	_ = v860
	var v863 int32
	_ = v863
	var v864 int32
	_ = v864
	var v867 int32
	_ = v867
	var v874 int32
	_ = v874
	var v875 int32
	_ = v875
	var v877 int32
	_ = v877
	var v883 int32
	_ = v883
	var v884 int32
	_ = v884
	var v888 int32
	_ = v888
	var v889 int32
	_ = v889
	var v892 int32
	_ = v892
	var v895 int32
	_ = v895
	var v898 int32
	_ = v898
	var v899 int32
	_ = v899
	var v902 int32
	_ = v902
	var v903 int32
	_ = v903
	var v906 int32
	_ = v906
	var v913 int32
	_ = v913
	var v914 int32
	_ = v914
	var v916 int32
	_ = v916
	var v922 int32
	_ = v922
	var v923 int32
	_ = v923
	var v927 int32
	_ = v927
	var v928 int32
	_ = v928
	var v931 int32
	_ = v931
	var v934 int32
	_ = v934
	var v937 int32
	_ = v937
	var v938 int32
	_ = v938
	var v941 int32
	_ = v941
	var v942 int32
	_ = v942
	var v945 int32
	_ = v945
	var v952 int32
	_ = v952
	var v953 int32
	_ = v953
	var v955 int32
	_ = v955
	var v961 int32
	_ = v961
	var v962 int32
	_ = v962
	var v966 int32
	_ = v966
	var v967 int32
	_ = v967
	var v970 int32
	_ = v970
	var v973 int32
	_ = v973
	var v976 int32
	_ = v976
	var v977 int32
	_ = v977
	var v980 int32
	_ = v980
	var v981 int32
	_ = v981
	var v984 int32
	_ = v984
	var v991 int32
	_ = v991
	var v992 int32
	_ = v992
	var v994 int32
	_ = v994
	var v1000 int32
	_ = v1000
	var v1001 int32
	_ = v1001
	var v1005 int32
	_ = v1005
	var v1006 int32
	_ = v1006
	var v1009 int32
	_ = v1009
	var v1012 int32
	_ = v1012
	var v1015 int32
	_ = v1015
	var v1016 int32
	_ = v1016
	var v1019 int32
	_ = v1019
	var v1020 int32
	_ = v1020
	var v1023 int32
	_ = v1023
	var v1030 int32
	_ = v1030
	var v1031 int32
	_ = v1031
	var v1033 int32
	_ = v1033
	var v1039 int32
	_ = v1039
	var v1040 int32
	_ = v1040
	var v1044 int32
	_ = v1044
	var v1045 int32
	_ = v1045
	var v1048 int32
	_ = v1048
	var v1051 int32
	_ = v1051
	var v1054 int32
	_ = v1054
	var v1055 int32
	_ = v1055
	var v1058 int32
	_ = v1058
	var v1059 int32
	_ = v1059
	var v1062 int32
	_ = v1062
	var v1069 int32
	_ = v1069
	var v1070 int32
	_ = v1070
	var v1072 int32
	_ = v1072
	var v1078 int32
	_ = v1078
	var v1079 int32
	_ = v1079
	var v1086 int32
	_ = v1086
	var v1089 int32
	_ = v1089
	var v1096 int32
	_ = v1096
	var v1101 int32
	_ = v1101
	var v1102 int32
	_ = v1102
	var v1103 int32
	_ = v1103
	var v1106 int32
	_ = v1106
	var v1109 int32
	_ = v1109
	var v1112 int32
	_ = v1112
	var v1113 int32
	_ = v1113
	var v1116 int32
	_ = v1116
	var v1117 int32
	_ = v1117
	var v1120 int32
	_ = v1120
	var v1127 int32
	_ = v1127
	var v1128 int32
	_ = v1128
	var v1130 int32
	_ = v1130
	var v1136 int32
	_ = v1136
	var v1138 int32
	_ = v1138
	var v1139 int32
	_ = v1139
	var v1140 int32
	_ = v1140
	var v1145 int32
	_ = v1145
	var v1146 int32
	_ = v1146
	var v1149 int32
	_ = v1149
	var v1150 int32
	_ = v1150
	var v1160 int32
	_ = v1160
	var v1169 int32
	_ = v1169
	var v1172 int32
	_ = v1172
	var v1174 int32
	_ = v1174
	var v1183 int32
	_ = v1183
	var v1186 int32
	_ = v1186
	var v1190 int32
	_ = v1190
	var v1191 int32
	_ = v1191
	var v1194 int32
	_ = v1194
	var v1195 int32
	_ = v1195
	var v1205 int32
	_ = v1205
	var v1214 int32
	_ = v1214
	var v1217 int32
	_ = v1217
	var v1219 int32
	_ = v1219
	var v1228 int32
	_ = v1228
	var v1234 int32
	_ = v1234
	var v1237 int32
	_ = v1237
	var v1238 int32
	_ = v1238
	var v1244 int32
	_ = v1244
	var v1249 int32
	_ = v1249
	var v1252 int32
	_ = v1252
	var v1253 int32
	_ = v1253
	var v1256 int32
	_ = v1256
	var v1259 int32
	_ = v1259
	var v1262 int32
	_ = v1262
	var v1263 int32
	_ = v1263
	var v1266 int32
	_ = v1266
	var v1267 int32
	_ = v1267
	var v1270 int32
	_ = v1270
	var v1277 int32
	_ = v1277
	var v1278 int32
	_ = v1278
	var v1280 int32
	_ = v1280
	var v1281 int32
	_ = v1281
	var v1282 int32
	_ = v1282
	var v1285 int32
	_ = v1285
	var v1288 int32
	_ = v1288
	var v1291 int32
	_ = v1291
	var v1294 int32
	_ = v1294
	var v1295 int32
	_ = v1295
	var v1298 int32
	_ = v1298
	var v1299 int32
	_ = v1299
	var v1302 int32
	_ = v1302
	var v1309 int32
	_ = v1309
	var v1310 int32
	_ = v1310
	var v1318 int64
	_ = v1318
	var v1319 int32
	_ = v1319
	var v1322 int32
	_ = v1322
	var v1323 int32
	_ = v1323
	var v1324 int64
	_ = v1324
	var v1331 int32
	_ = v1331
	var v1332 int32
	_ = v1332
	var v1335 int32
	_ = v1335
	var v1338 int32
	_ = v1338
	var v1341 int32
	_ = v1341
	var v1342 int32
	_ = v1342
	var v1345 int32
	_ = v1345
	var v1346 int32
	_ = v1346
	var v1349 int32
	_ = v1349
	var v1356 int32
	_ = v1356
	var v1357 int32
	_ = v1357
	var v1359 int32
	_ = v1359
	var v1365 int32
	_ = v1365
	var v1366 int32
	_ = v1366
	var v1370 int32
	_ = v1370
	var v1372 int32
	_ = v1372
	var v1373 int32
	_ = v1373
	var v1374 int32
	_ = v1374
	var v1378 int32
	_ = v1378
	var v1381 int32
	_ = v1381
	var v1384 int32
	_ = v1384
	var v1388 int32
	_ = v1388
	var v1391 int32
	_ = v1391
	var v1392 int32
	_ = v1392
	var v1398 int32
	_ = v1398
	var v1403 int32
	_ = v1403
	var v1407 int32
	_ = v1407
	var v1408 int32
	_ = v1408
	var v1416 int32
	_ = v1416
	var v1417 int32
	_ = v1417
	var v1422 int32
	_ = v1422
	var v1425 int32
	_ = v1425
	var v1431 int32
	_ = v1431
	var v1436 int32
	_ = v1436
	var v1468 int32
	_ = v1468
	var v1471 int32
	_ = v1471
	var v1474 int32
	_ = v1474
	var v1482 int32
	_ = v1482
	var v1485 int32
	_ = v1485
	var v1494 int32
	_ = v1494
	var v1499 int32
	_ = v1499
	var v1500 int32
	_ = v1500
	var v1503 int32
	_ = v1503
	var v1506 int32
	_ = v1506
	var v1509 int32
	_ = v1509
	var v1512 int32
	_ = v1512
	var v1516 int32
	_ = v1516
	var v1517 int32
	_ = v1517
	var v1522 int32
	_ = v1522
	var v1528 int32
	_ = v1528
	var v1531 int32
	_ = v1531
	var v1540 int32
	_ = v1540
	var v1545 int32
	_ = v1545
	var v1546 int32
	_ = v1546
	var v1552 int32
	_ = v1552
	var v1555 int32
	_ = v1555
	var v1566 int32
	_ = v1566
	var v1571 int32
	_ = v1571
	var v1579 int32
	_ = v1579
	var v1582 int32
	_ = v1582
	var v1591 int32
	_ = v1591
	var v1596 int32
	_ = v1596
	var v1600 int32
	_ = v1600
	var v1603 int32
	_ = v1603
	var v1612 int32
	_ = v1612
	var v1617 int32
	_ = v1617
	var v1626 int32
	_ = v1626
	var v1631 int32
	_ = v1631
	var v1640 int32
	_ = v1640
	var v1645 int32
	_ = v1645
	var v1649 int32
	_ = v1649
	v29 = int64(0)
	v30 = m.G0
	v32 = v30 - int32(176)
	m.G0 = v32
	*(*int64)(unsafe.Add(mBase, uint32(l3)+48)) = v29
	*(*int64)(unsafe.Add(mBase, uint32(l3)+40)) = v29
	*(*int64)(unsafe.Add(mBase, uint32(l3)+32)) = v29
	*(*int64)(unsafe.Add(mBase, uint32(l3)+24)) = v29
	*(*int64)(unsafe.Add(mBase, uint32(l3)+16)) = v29
	*(*int64)(unsafe.Add(mBase, uint32(l3)+8)) = v29
	*(*int64)(unsafe.Add(mBase, uint32(l3))) = v29
	v49 = l2 & int32(1)
	if v49 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v50 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+12)) = uint8(v50)
	goto L3
L2:
	;
	goto L3
L3:
	;
	v53 = l2 & int32(2)
	if v53 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v54 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+13)) = uint8(v54)
	goto L6
L5:
	;
	goto L6
L6:
	;
	v57 = l2 & int32(4)
	if v57 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v58 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+14)) = uint8(v58)
	goto L9
L8:
	;
	goto L9
L9:
	;
	v61 = l2 & int32(16)
	if v61 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v62 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+15)) = uint8(v62)
	goto L12
L11:
	;
	goto L12
L12:
	;
	v65 = l2 & int32(64)
	if v65 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v66 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+16)) = uint8(v66)
	goto L15
L14:
	;
	goto L15
L15:
	;
	v69 = l2 & int32(128)
	if v69 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v70 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+17)) = uint8(v70)
	goto L18
L17:
	;
	goto L18
L18:
	;
	v73 = l2 & int32(256)
	if v73 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v74 = int32(112)
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+18)) = uint8(v74)
	goto L21
L20:
	;
	goto L21
L21:
	;
	v77 = l2 & int32(512)
	if v77 != 0 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v78 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+19)) = uint8(v78)
	goto L24
L23:
	;
	goto L24
L24:
	;
	v81 = l2 & int32(1024)
	if v81 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v82 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+20)) = uint8(v82)
	goto L27
L26:
	;
	goto L27
L27:
	;
	v85 = l2 & int32(2048)
	if v85 != 0 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v86 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+21)) = uint8(v86)
	goto L30
L29:
	;
	goto L30
L30:
	;
	v89 = l2 & int32(_a_F_parse_subscription_options_0)
	if v89 != 0 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v90 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+22)) = uint8(v90)
	goto L33
L32:
	;
	goto L33
L33:
	;
	v93 = l2 & int32(_a_F_parse_subscription_options_1)
	if v93 != 0 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v94 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+23)) = uint8(v94)
	goto L36
L35:
	;
	goto L36
L36:
	;
	v97 = l2 & int32(_a_F_parse_subscription_options_2)
	if v97 != 0 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v98 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+24)) = uint8(v98)
	goto L39
L38:
	;
	goto L39
L39:
	;
	v101 = l2 & int32(_a_F_parse_subscription_options_3)
	if v101 != 0 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+28)) = int32(0)
	goto L42
L41:
	;
	goto L42
L42:
	;
	if base.Ui32(int32(_a_F_parse_subscription_options_4)) <= base.Ui32(l2) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v107 = F_pstrdup(m, int32(_a_F_parse_subscription_options_5))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L46
	} else {
		goto L47
	}
L44:
	;
	goto L45
L45:
	;
	if l1 == int32(0) {
		goto L49
	} else {
		goto L50
	}
L46:
	;
	return
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+32)) = v107
	goto L45
L48:
	;
	F_errorConflictingDefElem(m, v158, l0)
	mBase = m.M
	v1649 = m.ExcPending
	if v1649 != 0 {
		goto L46
	} else {
		goto L495
	}
L49:
	;
	if v49 == int32(0) {
		goto L448
	} else {
		goto L449
	}
L50:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v112 <= int32(0) {
		goto L49
	} else {
		goto L51
	}
L51:
	;
	v134 = int32(0)
	goto L52
L52:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v154+v134<<(uint(int32(2))%32))))
	if v49 == int32(0) {
		goto L57
	} else {
		goto L58
	}
L53:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1422 = m.ExcPending
	if v1422 != 0 {
		goto L46
	} else {
		goto L440
	}
L54:
	;
	goto L53
L55:
	;
	v1416 = v134 + int32(1)
	v1417 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v1416 < v1417 {
		v134 = v1416
		goto L52
	} else {
		goto L439
	}
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v189 | int32(1)
	v1407 = F_defGetBoolean(m, v158)
	mBase = m.M
	v1408 = m.ExcPending
	if v1408 != 0 {
		goto L46
	} else {
		goto L438
	}
L57:
	;
	if v53 == int32(0) {
		goto L68
	} else {
		goto L69
	}
L58:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v158)+8))
	v162 = int32(_a_F_parse_subscription_options_6)
	v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161))))
	v168 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_parse_subscription_options[0])))
	if base.B2i32(v165 == int32(0))|base.B2i32(v165 != v168) != 0 {
		v186 = v165
		v187 = v168
		goto L60
	} else {
		goto L61
	}
L59:
	;
	if v186-v187 != 0 {
		goto L57
	} else {
		goto L66
	}
L60:
	;
	goto L59
L61:
	;
	v171 = v161
	v172 = v162
	goto L62
L62:
	;
	v175 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v172)+1)))
	v176 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v171)+1)))
	if v176 == int32(0) {
		v186 = v176
		v187 = v175
		goto L60
	} else {
		goto L64
	}
L63:
	;
	v186 = v176
	v187 = v175
	goto L60
L64:
	;
	v179 = int32(1)
	if v176 == v175 {
		v171 = v171 + v179
		v172 = v172 + v179
		goto L62
	} else {
		goto L65
	}
L65:
	;
	goto L63
L66:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v189&int32(1) == int32(0) {
		goto L56
	} else {
		goto L67
	}
L67:
	;
	goto L48
L68:
	;
	if v57 == int32(0) {
		goto L80
	} else {
		goto L81
	}
L69:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v158)+8))
	v197 = int32(_a_F_parse_subscription_options_7)
	v200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v196))))
	v203 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_parse_subscription_options[1])))
	if base.B2i32(v200 == int32(0))|base.B2i32(v200 != v203) != 0 {
		v221 = v200
		v222 = v203
		goto L71
	} else {
		goto L72
	}
L70:
	;
	if v221-v222 != 0 {
		goto L68
	} else {
		goto L77
	}
L71:
	;
	goto L70
L72:
	;
	v206 = v196
	v207 = v197
	goto L73
L73:
	;
	v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207)+1)))
	v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206)+1)))
	if v211 == int32(0) {
		v221 = v211
		v222 = v210
		goto L71
	} else {
		goto L75
	}
L74:
	;
	v221 = v211
	v222 = v210
	goto L71
L75:
	;
	v214 = int32(1)
	if v211 == v210 {
		v206 = v206 + v214
		v207 = v207 + v214
		goto L73
	} else {
		goto L76
	}
L76:
	;
	goto L74
L77:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v224&int32(2) != 0 {
		goto L48
	} else {
		goto L78
	}
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v224 | int32(2)
	v230 = F_defGetBoolean(m, v158)
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L46
	} else {
		goto L79
	}
L79:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+13)) = uint8(v230)
	goto L55
L80:
	;
	if l2&int32(8) == int32(0) {
		goto L92
	} else {
		goto L93
	}
L81:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v158)+8))
	v236 = int32(_a_F_parse_subscription_options_8)
	v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v235))))
	v242 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_parse_subscription_options[2])))
	if base.B2i32(v239 == int32(0))|base.B2i32(v239 != v242) != 0 {
		v260 = v239
		v261 = v242
		goto L83
	} else {
		goto L84
	}
L82:
	;
	if v260-v261 != 0 {
		goto L80
	} else {
		goto L89
	}
L83:
	;
	goto L82
L84:
	;
	v245 = v235
	v246 = v236
	goto L85
L85:
	;
	v249 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v246)+1)))
	v250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v245)+1)))
	if v250 == int32(0) {
		v260 = v250
		v261 = v249
		goto L83
	} else {
		goto L87
	}
L86:
	;
	v260 = v250
	v261 = v249
	goto L83
L87:
	;
	v253 = int32(1)
	if v250 == v249 {
		v245 = v245 + v253
		v246 = v246 + v253
		goto L85
	} else {
		goto L88
	}
L88:
	;
	goto L86
L89:
	;
	v263 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v263&int32(4) != 0 {
		goto L48
	} else {
		goto L90
	}
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v263 | int32(4)
	v269 = F_defGetBoolean(m, v158)
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L46
	} else {
		goto L91
	}
L91:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+14)) = uint8(v269)
	goto L55
L92:
	;
	if v61 == int32(0) {
		goto L115
	} else {
		goto L116
	}
L93:
	;
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v158)+8))
	v275 = int32(_a_F_parse_subscription_options_9)
	v278 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v274))))
	v281 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_parse_subscription_options[3])))
	if base.B2i32(v278 == int32(0))|base.B2i32(v278 != v281) != 0 {
		v299 = v278
		v300 = v281
		goto L95
	} else {
		goto L96
	}
L94:
	;
	if v299-v300 != 0 {
		goto L92
	} else {
		goto L101
	}
L95:
	;
	goto L94
L96:
	;
	v284 = v274
	v285 = v275
	goto L97
L97:
	;
	v288 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v285)+1)))
	v289 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v284)+1)))
	if v289 == int32(0) {
		v299 = v289
		v300 = v288
		goto L95
	} else {
		goto L99
	}
L98:
	;
	v299 = v289
	v300 = v288
	goto L95
L99:
	;
	v292 = int32(1)
	if v289 == v288 {
		v284 = v284 + v292
		v285 = v285 + v292
		goto L97
	} else {
		goto L100
	}
L100:
	;
	goto L98
L101:
	;
	v302 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v302&int32(8) != 0 {
		goto L48
	} else {
		goto L102
	}
L102:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v302 | int32(8)
	v308 = F_defGetString(m, v158)
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L46
	} else {
		goto L103
	}
L103:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+4)) = v308
	v311 = int32(_a_F_parse_subscription_options_10)
	v314 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v308))))
	v317 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_parse_subscription_options[4])))
	if base.B2i32(v314 == int32(0))|base.B2i32(v314 != v317) != 0 {
		v335 = v314
		v336 = v317
		goto L105
	} else {
		goto L106
	}
L104:
	;
	if v335-v336 == int32(0) {
		goto L111
	} else {
		goto L112
	}
L105:
	;
	goto L104
L106:
	;
	v320 = v308
	v321 = v311
	goto L107
L107:
	;
	v324 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v321)+1)))
	v325 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v320)+1)))
	if v325 == int32(0) {
		v335 = v325
		v336 = v324
		goto L105
	} else {
		goto L109
	}
L108:
	;
	v335 = v325
	v336 = v324
	goto L105
L109:
	;
	v328 = int32(1)
	if v325 == v324 {
		v320 = v320 + v328
		v321 = v321 + v328
		goto L107
	} else {
		goto L110
	}
L110:
	;
	goto L108
L111:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+4)) = int32(0)
	goto L55
L112:
	;
	goto L113
L113:
	;
	v344 = F_ReplicationSlotValidateName(m, v308, int32(0), int32(21))
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L46
	} else {
		goto L114
	}
L114:
	;
	goto L55
L115:
	;
	if l2&int32(32) == int32(0) {
		goto L127
	} else {
		goto L128
	}
L116:
	;
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v158)+8))
	v349 = int32(_a_F_parse_subscription_options_11)
	v352 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v348))))
	v355 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_parse_subscription_options[5])))
	if base.B2i32(v352 == int32(0))|base.B2i32(v352 != v355) != 0 {
		v373 = v352
		v374 = v355
		goto L118
	} else {
		goto L119
	}
L117:
	;
	if v373-v374 != 0 {
		goto L115
	} else {
		goto L124
	}
L118:
	;
	goto L117
L119:
	;
	v358 = v348
	v359 = v349
	goto L120
L120:
	;
	v362 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v359)+1)))
	v363 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v358)+1)))
	if v363 == int32(0) {
		v373 = v363
		v374 = v362
		goto L118
	} else {
		goto L122
	}
L121:
	;
	v373 = v363
	v374 = v362
	goto L118
L122:
	;
	v366 = int32(1)
	if v363 == v362 {
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
	v376 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v376&int32(16) != 0 {
		goto L48
	} else {
		goto L125
	}
L125:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v376 | int32(16)
	v382 = F_defGetBoolean(m, v158)
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L46
	} else {
		goto L126
	}
L126:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+15)) = uint8(v382)
	goto L55
L127:
	;
	if v65 == int32(0) {
		goto L140
	} else {
		goto L141
	}
L128:
	;
	v387 = *(*int32)(unsafe.Add(mBase, uint32(v158)+8))
	v388 = int32(_a_F_parse_subscription_options_12)
	v391 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v387))))
	v394 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_parse_subscription_options[6])))
	if base.B2i32(v391 == int32(0))|base.B2i32(v391 != v394) != 0 {
		v412 = v391
		v413 = v394
		goto L130
	} else {
		goto L131
	}
L129:
	;
	if v412-v413 != 0 {
		goto L127
	} else {
		goto L136
	}
L130:
	;
	goto L129
L131:
	;
	v397 = v387
	v398 = v388
	goto L132
L132:
	;
	v401 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v398)+1)))
	v402 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v397)+1)))
	if v402 == int32(0) {
		v412 = v402
		v413 = v401
		goto L130
	} else {
		goto L134
	}
L133:
	;
	v412 = v402
	v413 = v401
	goto L130
L134:
	;
	v405 = int32(1)
	if v402 == v401 {
		v397 = v397 + v405
		v398 = v398 + v405
		goto L132
	} else {
		goto L135
	}
L135:
	;
	goto L133
L136:
	;
	v415 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v415&int32(32) != 0 {
		goto L48
	} else {
		goto L137
	}
L137:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v415 | int32(32)
	v421 = F_defGetString(m, v158)
	mBase = m.M
	v422 = m.ExcPending
	if v422 != 0 {
		goto L46
	} else {
		goto L138
	}
L138:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+8)) = v421
	v427 = int32(0)
	F_set_config_option(m, int32(_a_F_parse_subscription_options_12), v421, int32(4), int32(12), v427, v427)
	mBase = m.M
	v430 = m.ExcPending
	if v430 != 0 {
		goto L46
	} else {
		goto L139
	}
L139:
	;
	goto L55
L140:
	;
	if v69 == int32(0) {
		goto L152
	} else {
		goto L153
	}
L141:
	;
	v433 = *(*int32)(unsafe.Add(mBase, uint32(v158)+8))
	v434 = int32(_a_F_parse_subscription_options_13)
	v437 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v433))))
	v440 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_parse_subscription_options[7])))
	if base.B2i32(v437 == int32(0))|base.B2i32(v437 != v440) != 0 {
		v458 = v437
		v459 = v440
		goto L143
	} else {
		goto L144
	}
L142:
	;
	if v458-v459 != 0 {
		goto L140
	} else {
		goto L149
	}
L143:
	;
	goto L142
L144:
	;
	v443 = v433
	v444 = v434
	goto L145
L145:
	;
	v447 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v444)+1)))
	v448 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v443)+1)))
	if v448 == int32(0) {
		v458 = v448
		v459 = v447
		goto L143
	} else {
		goto L147
	}
L146:
	;
	v458 = v448
	v459 = v447
	goto L143
L147:
	;
	v451 = int32(1)
	if v448 == v447 {
		v443 = v443 + v451
		v444 = v444 + v451
		goto L145
	} else {
		goto L148
	}
L148:
	;
	goto L146
L149:
	;
	v461 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v461&int32(64) != 0 {
		goto L48
	} else {
		goto L150
	}
L150:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v461 | int32(64)
	v467 = F_defGetBoolean(m, v158)
	mBase = m.M
	v468 = m.ExcPending
	if v468 != 0 {
		goto L46
	} else {
		goto L151
	}
L151:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+16)) = uint8(v467)
	goto L55
L152:
	;
	if v73 == int32(0) {
		goto L164
	} else {
		goto L165
	}
L153:
	;
	v472 = *(*int32)(unsafe.Add(mBase, uint32(v158)+8))
	v473 = int32(_a_F_parse_subscription_options_14)
	v476 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v472))))
	v479 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_parse_subscription_options[8])))
	if base.B2i32(v476 == int32(0))|base.B2i32(v476 != v479) != 0 {
		v497 = v476
		v498 = v479
		goto L155
	} else {
		goto L156
	}
L154:
	;
	if v497-v498 != 0 {
		goto L152
	} else {
		goto L161
	}
L155:
	;
	goto L154
L156:
	;
	v482 = v472
	v483 = v473
	goto L157
L157:
	;
	v486 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v483)+1)))
	v487 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v482)+1)))
	if v487 == int32(0) {
		v497 = v487
		v498 = v486
		goto L155
	} else {
		goto L159
	}
L158:
	;
	v497 = v487
	v498 = v486
	goto L155
L159:
	;
	v490 = int32(1)
	if v487 == v486 {
		v482 = v482 + v490
		v483 = v483 + v490
		goto L157
	} else {
		goto L160
	}
L160:
	;
	goto L158
L161:
	;
	v500 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v500&int32(128) != 0 {
		goto L48
	} else {
		goto L162
	}
L162:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v500 | int32(128)
	v506 = F_defGetBoolean(m, v158)
	mBase = m.M
	v507 = m.ExcPending
	if v507 != 0 {
		goto L46
	} else {
		goto L163
	}
L163:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+17)) = uint8(v506)
	goto L55
L164:
	;
	if v77 == int32(0) {
		goto L257
	} else {
		goto L258
	}
L165:
	;
	v511 = *(*int32)(unsafe.Add(mBase, uint32(v158)+8))
	v512 = int32(_a_F_parse_subscription_options_15)
	v515 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v511))))
	v518 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_parse_subscription_options[9])))
	if base.B2i32(v515 == int32(0))|base.B2i32(v515 != v518) != 0 {
		v536 = v515
		v537 = v518
		goto L167
	} else {
		goto L168
	}
L166:
	;
	if v536-v537 != 0 {
		goto L164
	} else {
		goto L173
	}
L167:
	;
	goto L166
L168:
	;
	v521 = v511
	v522 = v512
	goto L169
L169:
	;
	v525 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v522)+1)))
	v526 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v521)+1)))
	if v526 == int32(0) {
		v536 = v526
		v537 = v525
		goto L167
	} else {
		goto L171
	}
L170:
	;
	v536 = v526
	v537 = v525
	goto L167
L171:
	;
	v529 = int32(1)
	if v526 == v525 {
		v521 = v521 + v529
		v522 = v522 + v529
		goto L169
	} else {
		goto L172
	}
L172:
	;
	goto L170
L173:
	;
	v539 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v539&int32(256) != 0 {
		goto L48
	} else {
		goto L174
	}
L174:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v539 | int32(256)
	v545 = m.G0
	v547 = v545 - int32(16)
	m.G0 = v547
	v549 = *(*int32)(unsafe.Add(mBase, uint32(v158)+12))
	if v549 == int32(0) {
		goto L176
	} else {
		goto L177
	}
L175:
	;
	m.G0 = v547 + int32(16)
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+18)) = uint8(v802)
	goto L55
L176:
	;
	v802 = int32(116)
	goto L175
L177:
	;
	v552 = *(*int32)(unsafe.Add(mBase, uint32(v549)))
	if v552 == int32(473) {
		goto L179
	} else {
		goto L180
	}
L178:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v785 = m.ExcPending
	if v785 != 0 {
		goto L46
	} else {
		goto L253
	}
L179:
	;
	v556 = *(*int32)(unsafe.Add(mBase, uint32(v549)+4))
	switch v556 {
	case 0:
		v802 = int32(102)
		goto L175
	case 1:
		goto L176
	default:
		goto L178
	}
L180:
	;
	goto L181
L181:
	;
	v557 = int32(102)
	v558 = F_defGetString(m, v158)
	mBase = m.M
	v559 = m.ExcPending
	if v559 != 0 {
		goto L46
	} else {
		goto L182
	}
L182:
	;
	v563 = v558
	v564 = int32(_a_F_parse_subscription_options_16)
	goto L184
L183:
	;
	if v601 == int32(0) {
		v802 = v557
		goto L175
	} else {
		goto L196
	}
L184:
	;
	v567 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v563))))
	v568 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v564))))
	if v567 == v568 {
		v590 = v567
		goto L186
	} else {
		goto L187
	}
L185:
	;
	v601 = int32(0)
	goto L183
L186:
	;
	v592 = int32(1)
	if v590 != 0 {
		v563 = v563 + v592
		v564 = v564 + v592
		goto L184
	} else {
		goto L195
	}
L187:
	;
	if base.Ui32((v567-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L188
	} else {
		goto L189
	}
L188:
	;
	v578 = v567 | int32(32)
	goto L190
L189:
	;
	v578 = v567
	goto L190
L190:
	;
	if base.Ui32((v568-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L191
	} else {
		goto L192
	}
L191:
	;
	v587 = v568 | int32(32)
	goto L193
L192:
	;
	v587 = v568
	goto L193
L193:
	;
	if v578 == v587 {
		v590 = v578
		goto L186
	} else {
		goto L194
	}
L194:
	;
	v601 = v578 - v587
	goto L183
L195:
	;
	goto L185
L196:
	;
	v607 = v558
	v608 = int32(_a_F_parse_subscription_options_17)
	goto L198
L197:
	;
	if v645 == int32(0) {
		v802 = v557
		goto L175
	} else {
		goto L210
	}
L198:
	;
	v611 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v607))))
	v612 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v608))))
	if v611 == v612 {
		v634 = v611
		goto L200
	} else {
		goto L201
	}
L199:
	;
	v645 = int32(0)
	goto L197
L200:
	;
	v636 = int32(1)
	if v634 != 0 {
		v607 = v607 + v636
		v608 = v608 + v636
		goto L198
	} else {
		goto L209
	}
L201:
	;
	if base.Ui32((v611-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L202
	} else {
		goto L203
	}
L202:
	;
	v622 = v611 | int32(32)
	goto L204
L203:
	;
	v622 = v611
	goto L204
L204:
	;
	if base.Ui32((v612-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L205
	} else {
		goto L206
	}
L205:
	;
	v631 = v612 | int32(32)
	goto L207
L206:
	;
	v631 = v612
	goto L207
L207:
	;
	if v622 == v631 {
		v634 = v622
		goto L200
	} else {
		goto L208
	}
L208:
	;
	v645 = v622 - v631
	goto L197
L209:
	;
	goto L199
L210:
	;
	v651 = v558
	v652 = int32(_a_F_parse_subscription_options_18)
	goto L212
L211:
	;
	if v689 == int32(0) {
		goto L176
	} else {
		goto L224
	}
L212:
	;
	v655 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v651))))
	v656 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v652))))
	if v655 == v656 {
		v678 = v655
		goto L214
	} else {
		goto L215
	}
L213:
	;
	v689 = int32(0)
	goto L211
L214:
	;
	v680 = int32(1)
	if v678 != 0 {
		v651 = v651 + v680
		v652 = v652 + v680
		goto L212
	} else {
		goto L223
	}
L215:
	;
	if base.Ui32((v655-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L216
	} else {
		goto L217
	}
L216:
	;
	v666 = v655 | int32(32)
	goto L218
L217:
	;
	v666 = v655
	goto L218
L218:
	;
	if base.Ui32((v656-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L219
	} else {
		goto L220
	}
L219:
	;
	v675 = v656 | int32(32)
	goto L221
L220:
	;
	v675 = v656
	goto L221
L221:
	;
	if v666 == v675 {
		v678 = v666
		goto L214
	} else {
		goto L222
	}
L222:
	;
	v689 = v666 - v675
	goto L211
L223:
	;
	goto L213
L224:
	;
	v696 = v558
	v697 = int32(_a_F_parse_subscription_options_19)
	goto L226
L225:
	;
	if v734 == int32(0) {
		v802 = int32(116)
		goto L175
	} else {
		goto L238
	}
L226:
	;
	v700 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v696))))
	v701 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v697))))
	if v700 == v701 {
		v723 = v700
		goto L228
	} else {
		goto L229
	}
L227:
	;
	v734 = int32(0)
	goto L225
L228:
	;
	v725 = int32(1)
	if v723 != 0 {
		v696 = v696 + v725
		v697 = v697 + v725
		goto L226
	} else {
		goto L237
	}
L229:
	;
	if base.Ui32((v700-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L230
	} else {
		goto L231
	}
L230:
	;
	v711 = v700 | int32(32)
	goto L232
L231:
	;
	v711 = v700
	goto L232
L232:
	;
	if base.Ui32((v701-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L233
	} else {
		goto L234
	}
L233:
	;
	v720 = v701 | int32(32)
	goto L235
L234:
	;
	v720 = v701
	goto L235
L235:
	;
	if v711 == v720 {
		v723 = v711
		goto L228
	} else {
		goto L236
	}
L236:
	;
	v734 = v711 - v720
	goto L225
L237:
	;
	goto L227
L238:
	;
	v740 = v558
	v741 = int32(_a_F_parse_subscription_options_20)
	goto L240
L239:
	;
	if v778 != 0 {
		goto L178
	} else {
		goto L252
	}
L240:
	;
	v744 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v740))))
	v745 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v741))))
	if v744 == v745 {
		v767 = v744
		goto L242
	} else {
		goto L243
	}
L241:
	;
	v778 = int32(0)
	goto L239
L242:
	;
	v769 = int32(1)
	if v767 != 0 {
		v740 = v740 + v769
		v741 = v741 + v769
		goto L240
	} else {
		goto L251
	}
L243:
	;
	if base.Ui32((v744-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L244
	} else {
		goto L245
	}
L244:
	;
	v755 = v744 | int32(32)
	goto L246
L245:
	;
	v755 = v744
	goto L246
L246:
	;
	if base.Ui32((v745-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L247
	} else {
		goto L248
	}
L247:
	;
	v764 = v745 | int32(32)
	goto L249
L248:
	;
	v764 = v745
	goto L249
L249:
	;
	if v755 == v764 {
		v767 = v755
		goto L242
	} else {
		goto L250
	}
L250:
	;
	v778 = v755 - v764
	goto L239
L251:
	;
	goto L241
L252:
	;
	v802 = int32(112)
	goto L175
L253:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v788 = m.ExcPending
	if v788 != 0 {
		goto L46
	} else {
		goto L254
	}
L254:
	;
	v789 = *(*int32)(unsafe.Add(mBase, uint32(v158)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v547))) = v789
	F_errmsg(m, int32(_a_F_parse_subscription_options_21), v547)
	mBase = m.M
	v793 = m.ExcPending
	if v793 != 0 {
		goto L46
	} else {
		goto L255
	}
L255:
	;
	F_errfinish(m, int32(_a_F_parse_subscription_options_22), int32(3658), int32(_a_F_parse_subscription_options_23))
	mBase = m.M
	v798 = m.ExcPending
	if v798 != 0 {
		goto L46
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
	if v81 == int32(0) {
		goto L269
	} else {
		goto L270
	}
L258:
	;
	v810 = *(*int32)(unsafe.Add(mBase, uint32(v158)+8))
	v811 = int32(_a_F_parse_subscription_options_24)
	v814 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v810))))
	v817 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_parse_subscription_options[10])))
	if base.B2i32(v814 == int32(0))|base.B2i32(v814 != v817) != 0 {
		v835 = v814
		v836 = v817
		goto L260
	} else {
		goto L261
	}
L259:
	;
	if v835-v836 != 0 {
		goto L257
	} else {
		goto L266
	}
L260:
	;
	goto L259
L261:
	;
	v820 = v810
	v821 = v811
	goto L262
L262:
	;
	v824 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v821)+1)))
	v825 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v820)+1)))
	if v825 == int32(0) {
		v835 = v825
		v836 = v824
		goto L260
	} else {
		goto L264
	}
L263:
	;
	v835 = v825
	v836 = v824
	goto L260
L264:
	;
	v828 = int32(1)
	if v825 == v824 {
		v820 = v820 + v828
		v821 = v821 + v828
		goto L262
	} else {
		goto L265
	}
L265:
	;
	goto L263
L266:
	;
	v838 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v838&int32(512) != 0 {
		goto L48
	} else {
		goto L267
	}
L267:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v838 | int32(512)
	v844 = F_defGetBoolean(m, v158)
	mBase = m.M
	v845 = m.ExcPending
	if v845 != 0 {
		goto L46
	} else {
		goto L268
	}
L268:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+19)) = uint8(v844)
	goto L55
L269:
	;
	if v85 == int32(0) {
		goto L281
	} else {
		goto L282
	}
L270:
	;
	v849 = *(*int32)(unsafe.Add(mBase, uint32(v158)+8))
	v850 = int32(_a_F_parse_subscription_options_25)
	v853 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v849))))
	v856 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_parse_subscription_options[11])))
	if base.B2i32(v853 == int32(0))|base.B2i32(v853 != v856) != 0 {
		v874 = v853
		v875 = v856
		goto L272
	} else {
		goto L273
	}
L271:
	;
	if v874-v875 != 0 {
		goto L269
	} else {
		goto L278
	}
L272:
	;
	goto L271
L273:
	;
	v859 = v849
	v860 = v850
	goto L274
L274:
	;
	v863 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v860)+1)))
	v864 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v859)+1)))
	if v864 == int32(0) {
		v874 = v864
		v875 = v863
		goto L272
	} else {
		goto L276
	}
L275:
	;
	v874 = v864
	v875 = v863
	goto L272
L276:
	;
	v867 = int32(1)
	if v864 == v863 {
		v859 = v859 + v867
		v860 = v860 + v867
		goto L274
	} else {
		goto L277
	}
L277:
	;
	goto L275
L278:
	;
	v877 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v877&int32(1024) != 0 {
		goto L48
	} else {
		goto L279
	}
L279:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v877 | int32(1024)
	v883 = F_defGetBoolean(m, v158)
	mBase = m.M
	v884 = m.ExcPending
	if v884 != 0 {
		goto L46
	} else {
		goto L280
	}
L280:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+20)) = uint8(v883)
	goto L55
L281:
	;
	if v89 == int32(0) {
		goto L293
	} else {
		goto L294
	}
L282:
	;
	v888 = *(*int32)(unsafe.Add(mBase, uint32(v158)+8))
	v889 = int32(_a_F_parse_subscription_options_26)
	v892 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v888))))
	v895 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_parse_subscription_options[12])))
	if base.B2i32(v892 == int32(0))|base.B2i32(v892 != v895) != 0 {
		v913 = v892
		v914 = v895
		goto L284
	} else {
		goto L285
	}
L283:
	;
	if v913-v914 != 0 {
		goto L281
	} else {
		goto L290
	}
L284:
	;
	goto L283
L285:
	;
	v898 = v888
	v899 = v889
	goto L286
L286:
	;
	v902 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v899)+1)))
	v903 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v898)+1)))
	if v903 == int32(0) {
		v913 = v903
		v914 = v902
		goto L284
	} else {
		goto L288
	}
L287:
	;
	v913 = v903
	v914 = v902
	goto L284
L288:
	;
	v906 = int32(1)
	if v903 == v902 {
		v898 = v898 + v906
		v899 = v899 + v906
		goto L286
	} else {
		goto L289
	}
L289:
	;
	goto L287
L290:
	;
	v916 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v916&int32(2048) != 0 {
		goto L48
	} else {
		goto L291
	}
L291:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v916 | int32(2048)
	v922 = F_defGetBoolean(m, v158)
	mBase = m.M
	v923 = m.ExcPending
	if v923 != 0 {
		goto L46
	} else {
		goto L292
	}
L292:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+21)) = uint8(v922)
	goto L55
L293:
	;
	if v93 == int32(0) {
		goto L305
	} else {
		goto L306
	}
L294:
	;
	v927 = *(*int32)(unsafe.Add(mBase, uint32(v158)+8))
	v928 = int32(_a_F_parse_subscription_options_27)
	v931 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v927))))
	v934 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_parse_subscription_options[13])))
	if base.B2i32(v931 == int32(0))|base.B2i32(v931 != v934) != 0 {
		v952 = v931
		v953 = v934
		goto L296
	} else {
		goto L297
	}
L295:
	;
	if v952-v953 != 0 {
		goto L293
	} else {
		goto L302
	}
L296:
	;
	goto L295
L297:
	;
	v937 = v927
	v938 = v928
	goto L298
L298:
	;
	v941 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v938)+1)))
	v942 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v937)+1)))
	if v942 == int32(0) {
		v952 = v942
		v953 = v941
		goto L296
	} else {
		goto L300
	}
L299:
	;
	v952 = v942
	v953 = v941
	goto L296
L300:
	;
	v945 = int32(1)
	if v942 == v941 {
		v937 = v937 + v945
		v938 = v938 + v945
		goto L298
	} else {
		goto L301
	}
L301:
	;
	goto L299
L302:
	;
	v955 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v955&int32(_a_F_parse_subscription_options_0) != 0 {
		goto L48
	} else {
		goto L303
	}
L303:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v955 | int32(_a_F_parse_subscription_options_0)
	v961 = F_defGetBoolean(m, v158)
	mBase = m.M
	v962 = m.ExcPending
	if v962 != 0 {
		goto L46
	} else {
		goto L304
	}
L304:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+22)) = uint8(v961)
	goto L55
L305:
	;
	if v97 == int32(0) {
		goto L317
	} else {
		goto L318
	}
L306:
	;
	v966 = *(*int32)(unsafe.Add(mBase, uint32(v158)+8))
	v967 = int32(_a_F_parse_subscription_options_28)
	v970 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v966))))
	v973 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_parse_subscription_options[14])))
	if base.B2i32(v970 == int32(0))|base.B2i32(v970 != v973) != 0 {
		v991 = v970
		v992 = v973
		goto L308
	} else {
		goto L309
	}
L307:
	;
	if v991-v992 != 0 {
		goto L305
	} else {
		goto L314
	}
L308:
	;
	goto L307
L309:
	;
	v976 = v966
	v977 = v967
	goto L310
L310:
	;
	v980 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v977)+1)))
	v981 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v976)+1)))
	if v981 == int32(0) {
		v991 = v981
		v992 = v980
		goto L308
	} else {
		goto L312
	}
L311:
	;
	v991 = v981
	v992 = v980
	goto L308
L312:
	;
	v984 = int32(1)
	if v981 == v980 {
		v976 = v976 + v984
		v977 = v977 + v984
		goto L310
	} else {
		goto L313
	}
L313:
	;
	goto L311
L314:
	;
	v994 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v994&int32(_a_F_parse_subscription_options_1) != 0 {
		goto L48
	} else {
		goto L315
	}
L315:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v994 | int32(_a_F_parse_subscription_options_1)
	v1000 = F_defGetBoolean(m, v158)
	mBase = m.M
	v1001 = m.ExcPending
	if v1001 != 0 {
		goto L46
	} else {
		goto L316
	}
L316:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+23)) = uint8(v1000)
	goto L55
L317:
	;
	if v101 == int32(0) {
		goto L329
	} else {
		goto L330
	}
L318:
	;
	v1005 = *(*int32)(unsafe.Add(mBase, uint32(v158)+8))
	v1006 = int32(_a_F_parse_subscription_options_29)
	v1009 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1005))))
	v1012 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_parse_subscription_options[15])))
	if base.B2i32(v1009 == int32(0))|base.B2i32(v1009 != v1012) != 0 {
		v1030 = v1009
		v1031 = v1012
		goto L320
	} else {
		goto L321
	}
L319:
	;
	if v1030-v1031 != 0 {
		goto L317
	} else {
		goto L326
	}
L320:
	;
	goto L319
L321:
	;
	v1015 = v1005
	v1016 = v1006
	goto L322
L322:
	;
	v1019 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1016)+1)))
	v1020 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1015)+1)))
	if v1020 == int32(0) {
		v1030 = v1020
		v1031 = v1019
		goto L320
	} else {
		goto L324
	}
L323:
	;
	v1030 = v1020
	v1031 = v1019
	goto L320
L324:
	;
	v1023 = int32(1)
	if v1020 == v1019 {
		v1015 = v1015 + v1023
		v1016 = v1016 + v1023
		goto L322
	} else {
		goto L325
	}
L325:
	;
	goto L323
L326:
	;
	v1033 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v1033&int32(_a_F_parse_subscription_options_2) != 0 {
		goto L48
	} else {
		goto L327
	}
L327:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v1033 | int32(_a_F_parse_subscription_options_2)
	v1039 = F_defGetBoolean(m, v158)
	mBase = m.M
	v1040 = m.ExcPending
	if v1040 != 0 {
		goto L46
	} else {
		goto L328
	}
L328:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+24)) = uint8(v1039)
	goto L55
L329:
	;
	if base.Ui32(l2) < base.Ui32(int32(_a_F_parse_subscription_options_4)) {
		goto L346
	} else {
		goto L347
	}
L330:
	;
	v1044 = *(*int32)(unsafe.Add(mBase, uint32(v158)+8))
	v1045 = int32(_a_F_parse_subscription_options_30)
	v1048 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1044))))
	v1051 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_parse_subscription_options[16])))
	if base.B2i32(v1048 == int32(0))|base.B2i32(v1048 != v1051) != 0 {
		v1069 = v1048
		v1070 = v1051
		goto L332
	} else {
		goto L333
	}
L331:
	;
	if v1069-v1070 != 0 {
		goto L329
	} else {
		goto L338
	}
L332:
	;
	goto L331
L333:
	;
	v1054 = v1044
	v1055 = v1045
	goto L334
L334:
	;
	v1058 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1055)+1)))
	v1059 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1054)+1)))
	if v1059 == int32(0) {
		v1069 = v1059
		v1070 = v1058
		goto L332
	} else {
		goto L336
	}
L335:
	;
	v1069 = v1059
	v1070 = v1058
	goto L332
L336:
	;
	v1062 = int32(1)
	if v1059 == v1058 {
		v1054 = v1054 + v1062
		v1055 = v1055 + v1062
		goto L334
	} else {
		goto L337
	}
L337:
	;
	goto L335
L338:
	;
	v1072 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v1072&int32(_a_F_parse_subscription_options_3) != 0 {
		goto L48
	} else {
		goto L339
	}
L339:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v1072 | int32(_a_F_parse_subscription_options_3)
	v1078 = F_defGetInt32(m, v158)
	mBase = m.M
	v1079 = m.ExcPending
	if v1079 != 0 {
		goto L46
	} else {
		goto L340
	}
L340:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+28)) = v1078
	if int32(0) <= v1078 {
		goto L55
	} else {
		goto L341
	}
L341:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1086 = m.ExcPending
	if v1086 != 0 {
		goto L46
	} else {
		goto L342
	}
L342:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v1089 = m.ExcPending
	if v1089 != 0 {
		goto L46
	} else {
		goto L343
	}
L343:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+160)) = int32(_a_F_parse_subscription_options_30)
	F_errmsg(m, int32(_a_F_parse_subscription_options_31), v32+int32(160))
	mBase = m.M
	v1096 = m.ExcPending
	if v1096 != 0 {
		goto L46
	} else {
		goto L344
	}
L344:
	;
	F_errfinish(m, int32(_a_F_parse_subscription_options_22), int32(364), int32(_a_F_parse_subscription_options_32))
	mBase = m.M
	v1101 = m.ExcPending
	if v1101 != 0 {
		goto L46
	} else {
		goto L345
	}
L345:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L346:
	;
	if l2&int32(_a_F_parse_subscription_options_33) == int32(0) {
		goto L391
	} else {
		goto L392
	}
L347:
	;
	v1102 = *(*int32)(unsafe.Add(mBase, uint32(v158)+8))
	v1103 = int32(_a_F_parse_subscription_options_34)
	v1106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1102))))
	v1109 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_parse_subscription_options[17])))
	if base.B2i32(v1106 == int32(0))|base.B2i32(v1106 != v1109) != 0 {
		v1127 = v1106
		v1128 = v1109
		goto L349
	} else {
		goto L350
	}
L348:
	;
	if v1127-v1128 != 0 {
		goto L346
	} else {
		goto L355
	}
L349:
	;
	goto L348
L350:
	;
	v1112 = v1102
	v1113 = v1103
	goto L351
L351:
	;
	v1116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1113)+1)))
	v1117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1112)+1)))
	if v1117 == int32(0) {
		v1127 = v1117
		v1128 = v1116
		goto L349
	} else {
		goto L353
	}
L352:
	;
	v1127 = v1117
	v1128 = v1116
	goto L349
L353:
	;
	v1120 = int32(1)
	if v1117 == v1116 {
		v1112 = v1112 + v1120
		v1113 = v1113 + v1120
		goto L351
	} else {
		goto L354
	}
L354:
	;
	goto L352
L355:
	;
	v1130 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v1130&int32(_a_F_parse_subscription_options_4) != 0 {
		goto L48
	} else {
		goto L356
	}
L356:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v1130 | int32(_a_F_parse_subscription_options_4)
	v1136 = *(*int32)(unsafe.Add(mBase, uint32(l3)+32))
	F_pfree(m, v1136)
	mBase = m.M
	v1138 = m.ExcPending
	if v1138 != 0 {
		goto L46
	} else {
		goto L357
	}
L357:
	;
	v1139 = F_defGetString(m, v158)
	mBase = m.M
	v1140 = m.ExcPending
	if v1140 != 0 {
		goto L46
	} else {
		goto L358
	}
L358:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+32)) = v1139
	v1145 = v1139
	v1146 = int32(_a_F_parse_subscription_options_10)
	goto L360
L359:
	;
	if v1183 == int32(0) {
		goto L55
	} else {
		goto L372
	}
L360:
	;
	v1149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1145))))
	v1150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1146))))
	if v1149 == v1150 {
		v1172 = v1149
		goto L362
	} else {
		goto L363
	}
L361:
	;
	v1183 = int32(0)
	goto L359
L362:
	;
	v1174 = int32(1)
	if v1172 != 0 {
		v1145 = v1145 + v1174
		v1146 = v1146 + v1174
		goto L360
	} else {
		goto L371
	}
L363:
	;
	if base.Ui32((v1149-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L364
	} else {
		goto L365
	}
L364:
	;
	v1160 = v1149 | int32(32)
	goto L366
L365:
	;
	v1160 = v1149
	goto L366
L366:
	;
	if base.Ui32((v1150-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L367
	} else {
		goto L368
	}
L367:
	;
	v1169 = v1150 | int32(32)
	goto L369
L368:
	;
	v1169 = v1150
	goto L369
L369:
	;
	if v1160 == v1169 {
		v1172 = v1160
		goto L362
	} else {
		goto L370
	}
L370:
	;
	v1183 = v1160 - v1169
	goto L359
L371:
	;
	goto L361
L372:
	;
	v1186 = *(*int32)(unsafe.Add(mBase, uint32(l3)+32))
	v1190 = v1186
	v1191 = int32(_a_F_parse_subscription_options_5)
	goto L374
L373:
	;
	if v1228 == int32(0) {
		goto L55
	} else {
		goto L386
	}
L374:
	;
	v1194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1190))))
	v1195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1191))))
	if v1194 == v1195 {
		v1217 = v1194
		goto L376
	} else {
		goto L377
	}
L375:
	;
	v1228 = int32(0)
	goto L373
L376:
	;
	v1219 = int32(1)
	if v1217 != 0 {
		v1190 = v1190 + v1219
		v1191 = v1191 + v1219
		goto L374
	} else {
		goto L385
	}
L377:
	;
	if base.Ui32((v1194-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L378
	} else {
		goto L379
	}
L378:
	;
	v1205 = v1194 | int32(32)
	goto L380
L379:
	;
	v1205 = v1194
	goto L380
L380:
	;
	if base.Ui32((v1195-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L381
	} else {
		goto L382
	}
L381:
	;
	v1214 = v1195 | int32(32)
	goto L383
L382:
	;
	v1214 = v1195
	goto L383
L383:
	;
	if v1205 == v1214 {
		v1217 = v1205
		goto L376
	} else {
		goto L384
	}
L384:
	;
	v1228 = v1205 - v1214
	goto L373
L385:
	;
	goto L375
L386:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1234 = m.ExcPending
	if v1234 != 0 {
		goto L46
	} else {
		goto L387
	}
L387:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v1237 = m.ExcPending
	if v1237 != 0 {
		goto L46
	} else {
		goto L388
	}
L388:
	;
	v1238 = *(*int32)(unsafe.Add(mBase, uint32(l3)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+144)) = v1238
	F_errmsg(m, int32(_a_F_parse_subscription_options_35), v32+int32(144))
	mBase = m.M
	v1244 = m.ExcPending
	if v1244 != 0 {
		goto L46
	} else {
		goto L389
	}
L389:
	;
	F_errfinish(m, int32(_a_F_parse_subscription_options_22), int32(387), int32(_a_F_parse_subscription_options_32))
	mBase = m.M
	v1249 = m.ExcPending
	if v1249 != 0 {
		goto L46
	} else {
		goto L390
	}
L390:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L391:
	;
	if l2&int32(_a_F_parse_subscription_options_36) == int32(0) {
		goto L416
	} else {
		goto L417
	}
L392:
	;
	v1252 = *(*int32)(unsafe.Add(mBase, uint32(v158)+8))
	v1253 = int32(_a_F_parse_subscription_options_37)
	v1256 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1252))))
	v1259 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_parse_subscription_options[18])))
	if base.B2i32(v1256 == int32(0))|base.B2i32(v1256 != v1259) != 0 {
		v1277 = v1256
		v1278 = v1259
		goto L394
	} else {
		goto L395
	}
L393:
	;
	if v1277-v1278 != 0 {
		goto L391
	} else {
		goto L400
	}
L394:
	;
	goto L393
L395:
	;
	v1262 = v1252
	v1263 = v1253
	goto L396
L396:
	;
	v1266 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1263)+1)))
	v1267 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1262)+1)))
	if v1267 == int32(0) {
		v1277 = v1267
		v1278 = v1266
		goto L394
	} else {
		goto L398
	}
L397:
	;
	v1277 = v1267
	v1278 = v1266
	goto L394
L398:
	;
	v1270 = int32(1)
	if v1267 == v1266 {
		v1262 = v1262 + v1270
		v1263 = v1263 + v1270
		goto L396
	} else {
		goto L399
	}
L399:
	;
	goto L397
L400:
	;
	v1280 = F_defGetString(m, v158)
	mBase = m.M
	v1281 = m.ExcPending
	if v1281 != 0 {
		goto L46
	} else {
		goto L401
	}
L401:
	;
	v1282 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v1282&int32(_a_F_parse_subscription_options_33) != 0 {
		goto L48
	} else {
		goto L402
	}
L402:
	;
	v1285 = int32(_a_F_parse_subscription_options_10)
	v1288 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1280))))
	v1291 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_parse_subscription_options[4])))
	if base.B2i32(v1288 == int32(0))|base.B2i32(v1288 != v1291) != 0 {
		v1309 = v1288
		v1310 = v1291
		goto L405
	} else {
		goto L406
	}
L403:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l3)+40)) = v1324
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v1323 | int32(_a_F_parse_subscription_options_33)
	goto L55
L404:
	;
	if v1309-v1310 == int32(0) {
		goto L411
	} else {
		goto L412
	}
L405:
	;
	goto L404
L406:
	;
	v1294 = v1280
	v1295 = v1285
	goto L407
L407:
	;
	v1298 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1295)+1)))
	v1299 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1294)+1)))
	if v1299 == int32(0) {
		v1309 = v1299
		v1310 = v1298
		goto L405
	} else {
		goto L409
	}
L408:
	;
	v1309 = v1299
	v1310 = v1298
	goto L405
L409:
	;
	v1302 = int32(1)
	if v1299 == v1298 {
		v1294 = v1294 + v1302
		v1295 = v1295 + v1302
		goto L407
	} else {
		goto L410
	}
L410:
	;
	goto L408
L411:
	;
	v1323 = v1282
	v1324 = int64(0)
	goto L403
L412:
	;
	goto L413
L413:
	;
	v1318 = F_DirectFunctionCall1Coll(m, int32(617), int32(0), base.I64_extend_i32_u(v1280))
	mBase = m.M
	v1319 = m.ExcPending
	if v1319 != 0 {
		goto L46
	} else {
		goto L414
	}
L414:
	;
	if v1318 == int64(0) {
		goto L54
	} else {
		goto L415
	}
L415:
	;
	v1322 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v1323 = v1322
	v1324 = v1318
	goto L403
L416:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1388 = m.ExcPending
	if v1388 != 0 {
		goto L46
	} else {
		goto L434
	}
L417:
	;
	v1331 = *(*int32)(unsafe.Add(mBase, uint32(v158)+8))
	v1332 = int32(_a_F_parse_subscription_options_38)
	v1335 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1331))))
	v1338 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_parse_subscription_options[19])))
	if base.B2i32(v1335 == int32(0))|base.B2i32(v1335 != v1338) != 0 {
		v1356 = v1335
		v1357 = v1338
		goto L419
	} else {
		goto L420
	}
L418:
	;
	if v1356-v1357 != 0 {
		goto L416
	} else {
		goto L425
	}
L419:
	;
	goto L418
L420:
	;
	v1341 = v1331
	v1342 = v1332
	goto L421
L421:
	;
	v1345 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1342)+1)))
	v1346 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1341)+1)))
	if v1346 == int32(0) {
		v1356 = v1346
		v1357 = v1345
		goto L419
	} else {
		goto L423
	}
L422:
	;
	v1356 = v1346
	v1357 = v1345
	goto L419
L423:
	;
	v1349 = int32(1)
	if v1346 == v1345 {
		v1341 = v1341 + v1349
		v1342 = v1342 + v1349
		goto L421
	} else {
		goto L424
	}
L424:
	;
	goto L422
L425:
	;
	v1359 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v1359&int32(_a_F_parse_subscription_options_36) != 0 {
		goto L48
	} else {
		goto L426
	}
L426:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v1359 | int32(_a_F_parse_subscription_options_36)
	v1365 = F_defGetString(m, v158)
	mBase = m.M
	v1366 = m.ExcPending
	if v1366 != 0 {
		goto L46
	} else {
		goto L427
	}
L427:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+48)) = v1365
	v1370 = int32(0)
	v1372 = F_parse_int(m, v1365, v32+int32(172), v1370, v1370)
	mBase = m.M
	v1373 = m.ExcPending
	if v1373 != 0 {
		goto L46
	} else {
		goto L428
	}
L428:
	;
	if v1372 != 0 {
		goto L429
	} else {
		goto L430
	}
L429:
	;
	v1374 = *(*int32)(unsafe.Add(mBase, uint32(v32)+172))
	if v1374 == int32(-1) {
		goto L55
	} else {
		goto L432
	}
L430:
	;
	goto L431
L431:
	;
	v1378 = *(*int32)(unsafe.Add(mBase, uint32(l3)+48))
	v1381 = int32(0)
	F_set_config_option(m, int32(_a_F_parse_subscription_options_38), v1378, int32(4), int32(12), v1381, v1381)
	mBase = m.M
	v1384 = m.ExcPending
	if v1384 != 0 {
		goto L46
	} else {
		goto L433
	}
L432:
	;
	goto L431
L433:
	;
	goto L55
L434:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v1391 = m.ExcPending
	if v1391 != 0 {
		goto L46
	} else {
		goto L435
	}
L435:
	;
	v1392 = *(*int32)(unsafe.Add(mBase, uint32(v158)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+112)) = v1392
	F_errmsg(m, int32(_a_F_parse_subscription_options_39), v32+int32(112))
	mBase = m.M
	v1398 = m.ExcPending
	if v1398 != 0 {
		goto L46
	} else {
		goto L436
	}
L436:
	;
	F_errfinish(m, int32(_a_F_parse_subscription_options_22), int32(443), int32(_a_F_parse_subscription_options_32))
	mBase = m.M
	v1403 = m.ExcPending
	if v1403 != 0 {
		goto L46
	} else {
		goto L437
	}
L437:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L438:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+12)) = uint8(v1407)
	goto L55
L439:
	;
	goto L49
L440:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v1425 = m.ExcPending
	if v1425 != 0 {
		goto L46
	} else {
		goto L441
	}
L441:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+128)) = v1280
	F_errmsg(m, int32(_a_F_parse_subscription_options_40), v32+int32(128))
	mBase = m.M
	v1431 = m.ExcPending
	if v1431 != 0 {
		goto L46
	} else {
		goto L442
	}
L442:
	;
	F_errfinish(m, int32(_a_F_parse_subscription_options_22), int32(410), int32(_a_F_parse_subscription_options_32))
	mBase = m.M
	v1436 = m.ExcPending
	if v1436 != 0 {
		goto L46
	} else {
		goto L443
	}
L443:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L444:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+52)) = int32(_a_F_parse_subscription_options_41)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+48)) = int32(_a_F_parse_subscription_options_42)
	F_errmsg(m, int32(_a_F_parse_subscription_options_43), v32+int32(48))
	mBase = m.M
	v1640 = m.ExcPending
	if v1640 != 0 {
		goto L46
	} else {
		goto L493
	}
L445:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+20)) = int32(_a_F_parse_subscription_options_44)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+16)) = int32(_a_F_parse_subscription_options_42)
	F_errmsg(m, int32(_a_F_parse_subscription_options_43), v32+int32(16))
	mBase = m.M
	v1626 = m.ExcPending
	if v1626 != 0 {
		goto L46
	} else {
		goto L491
	}
L446:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1600 = m.ExcPending
	if v1600 != 0 {
		goto L46
	} else {
		goto L487
	}
L447:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1579 = m.ExcPending
	if v1579 != 0 {
		goto L46
	} else {
		goto L483
	}
L448:
	;
	v1516 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v1516 != 0 {
		goto L466
	} else {
		goto L467
	}
L449:
	;
	v1468 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+12)))
	if v1468&int32(1) != 0 {
		goto L448
	} else {
		goto L450
	}
L450:
	;
	v1471 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+13)))
	if v1471 != int32(1) {
		goto L451
	} else {
		goto L452
	}
L451:
	;
	v1500 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+14)))
	if v1500 == int32(1) {
		goto L458
	} else {
		goto L459
	}
L452:
	;
	v1474 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3))))
	if v1474&int32(2) == int32(0) {
		goto L451
	} else {
		goto L453
	}
L453:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1482 = m.ExcPending
	if v1482 != 0 {
		goto L46
	} else {
		goto L454
	}
L454:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v1485 = m.ExcPending
	if v1485 != 0 {
		goto L46
	} else {
		goto L455
	}
L455:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+100)) = int32(_a_F_parse_subscription_options_44)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+96)) = int32(_a_F_parse_subscription_options_45)
	F_errmsg(m, int32(_a_F_parse_subscription_options_43), v32+int32(96))
	mBase = m.M
	v1494 = m.ExcPending
	if v1494 != 0 {
		goto L46
	} else {
		goto L456
	}
L456:
	;
	F_errfinish(m, int32(_a_F_parse_subscription_options_22), int32(459), int32(_a_F_parse_subscription_options_32))
	mBase = m.M
	v1499 = m.ExcPending
	if v1499 != 0 {
		goto L46
	} else {
		goto L457
	}
L457:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L458:
	;
	v1503 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3))))
	if v1503&int32(4) != 0 {
		goto L447
	} else {
		goto L461
	}
L459:
	;
	goto L460
L460:
	;
	v1506 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+15)))
	if v1506 == int32(1) {
		goto L462
	} else {
		goto L463
	}
L461:
	;
	goto L460
L462:
	;
	v1509 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3))))
	if v1509&int32(16) != 0 {
		goto L446
	} else {
		goto L465
	}
L463:
	;
	goto L464
L464:
	;
	v1512 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+15)) = uint8(v1512)
	*(*uint16)(unsafe.Add(mBase, uint32(l3)+13)) = uint16(v1512)
	goto L448
L465:
	;
	goto L464
L466:
	;
	m.G0 = v32 + int32(176)
	return
L467:
	;
	v1517 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v1517&int32(8) == int32(0) {
		goto L466
	} else {
		goto L468
	}
L468:
	;
	v1522 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+13)))
	if v1522 == int32(1) {
		goto L469
	} else {
		goto L470
	}
L469:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1528 = m.ExcPending
	if v1528 != 0 {
		goto L46
	} else {
		goto L472
	}
L470:
	;
	goto L471
L471:
	;
	v1546 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+14)))
	if v1546 != int32(1) {
		goto L466
	} else {
		goto L477
	}
L472:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v1531 = m.ExcPending
	if v1531 != 0 {
		goto L46
	} else {
		goto L473
	}
L473:
	;
	if v1517&int32(2) != 0 {
		goto L445
	} else {
		goto L474
	}
L474:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+4)) = int32(_a_F_parse_subscription_options_46)
	*(*int32)(unsafe.Add(mBase, uint32(v32))) = int32(_a_F_parse_subscription_options_42)
	F_errmsg(m, int32(_a_F_parse_subscription_options_47), v32)
	mBase = m.M
	v1540 = m.ExcPending
	if v1540 != 0 {
		goto L46
	} else {
		goto L475
	}
L475:
	;
	F_errfinish(m, int32(_a_F_parse_subscription_options_22), int32(501), int32(_a_F_parse_subscription_options_32))
	mBase = m.M
	v1545 = m.ExcPending
	if v1545 != 0 {
		goto L46
	} else {
		goto L476
	}
L476:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L477:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1552 = m.ExcPending
	if v1552 != 0 {
		goto L46
	} else {
		goto L478
	}
L478:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v1555 = m.ExcPending
	if v1555 != 0 {
		goto L46
	} else {
		goto L479
	}
L479:
	;
	if v1517&int32(4) != 0 {
		goto L444
	} else {
		goto L480
	}
L480:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+36)) = int32(_a_F_parse_subscription_options_48)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+32)) = int32(_a_F_parse_subscription_options_42)
	F_errmsg(m, int32(_a_F_parse_subscription_options_47), v32+int32(32))
	mBase = m.M
	v1566 = m.ExcPending
	if v1566 != 0 {
		goto L46
	} else {
		goto L481
	}
L481:
	;
	F_errfinish(m, int32(_a_F_parse_subscription_options_22), int32(517), int32(_a_F_parse_subscription_options_32))
	mBase = m.M
	v1571 = m.ExcPending
	if v1571 != 0 {
		goto L46
	} else {
		goto L482
	}
L482:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L483:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v1582 = m.ExcPending
	if v1582 != 0 {
		goto L46
	} else {
		goto L484
	}
L484:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+84)) = int32(_a_F_parse_subscription_options_41)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+80)) = int32(_a_F_parse_subscription_options_45)
	F_errmsg(m, int32(_a_F_parse_subscription_options_43), v32+int32(80))
	mBase = m.M
	v1591 = m.ExcPending
	if v1591 != 0 {
		goto L46
	} else {
		goto L485
	}
L485:
	;
	F_errfinish(m, int32(_a_F_parse_subscription_options_22), int32(466), int32(_a_F_parse_subscription_options_32))
	mBase = m.M
	v1596 = m.ExcPending
	if v1596 != 0 {
		goto L46
	} else {
		goto L486
	}
L486:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L487:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v1603 = m.ExcPending
	if v1603 != 0 {
		goto L46
	} else {
		goto L488
	}
L488:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+68)) = int32(_a_F_parse_subscription_options_49)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+64)) = int32(_a_F_parse_subscription_options_45)
	F_errmsg(m, int32(_a_F_parse_subscription_options_43), v32-int32(-64))
	mBase = m.M
	v1612 = m.ExcPending
	if v1612 != 0 {
		goto L46
	} else {
		goto L489
	}
L489:
	;
	F_errfinish(m, int32(_a_F_parse_subscription_options_22), int32(473), int32(_a_F_parse_subscription_options_32))
	mBase = m.M
	v1617 = m.ExcPending
	if v1617 != 0 {
		goto L46
	} else {
		goto L490
	}
L490:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L491:
	;
	F_errfinish(m, int32(_a_F_parse_subscription_options_22), int32(495), int32(_a_F_parse_subscription_options_32))
	mBase = m.M
	v1631 = m.ExcPending
	if v1631 != 0 {
		goto L46
	} else {
		goto L492
	}
L492:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L493:
	;
	F_errfinish(m, int32(_a_F_parse_subscription_options_22), int32(511), int32(_a_F_parse_subscription_options_32))
	mBase = m.M
	v1645 = m.ExcPending
	if v1645 != 0 {
		goto L46
	} else {
		goto L494
	}
L494:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L495:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_subscription_change_cb(m *base.Module, l0 int64, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	v5 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_subscription_change_cb[0])) = uint8(v5)
	return
}
