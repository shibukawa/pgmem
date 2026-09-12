package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_AddSubscriptionRelState(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64, l4 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int64
	_ = v25
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	v8 = m.G0
	v10 = v8 - int32(48)
	m.G0 = v10
	F_LockSharedObject(m, int32(6100), l0, int32(1))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return
	} else {
		v18 = F_table_open(m, int32(6102), int32(3))
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return
		} else {
			v21 = F_SearchSysCacheCopy(m, int32(68), l1, l0)
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return
			} else {
				if v21 == int32(0) {
					v25 = int64(0)
					*(*int64)(unsafe.Add(mBase, uint32(v10)+24)) = v25
					*(*int64)(unsafe.Add(mBase, uint32(v10)+16)) = v25
					*(*int32)(unsafe.Add(mBase, uint32(v10)+44)) = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = l2
					*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = l1
					*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = l0
					if l3 != v25 {
						v36 = F_Int64GetDatum(m, l3)
						mBase = m.M
						v37 = m.ExcPending
						if v37 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v10)+28)) = v36
							v41 = *(*int32)(unsafe.Add(mBase, uint32(v18)+52))
							v46 = F_heap_form_tuple(m, v41, v10+int32(16), v10+int32(44))
							mBase = m.M
							v47 = m.ExcPending
							if v47 != 0 {
								return
							} else {
								F_CatalogTupleInsert(m, v18, v46)
								mBase = m.M
								v49 = m.ExcPending
								if v49 != 0 {
									return
								} else {
									F_pfree(m, v46)
									mBase = m.M
									v51 = m.ExcPending
									if v51 != 0 {
										return
									} else {
										if l4 != 0 {
											F_sequence_close(m, v18, int32(0))
											mBase = m.M
											v54 = m.ExcPending
											if v54 != 0 {
												return
											} else {
												m.G0 = v10 + int32(48)
												return
											}
										} else {
											F_sequence_close(m, v18, int32(3))
											mBase = m.M
											v57 = m.ExcPending
											if v57 != 0 {
												return
											} else {
												F_UnlockSharedObject(m, int32(6100), l0, int32(1))
												mBase = m.M
												v61 = m.ExcPending
												if v61 != 0 {
													return
												} else {
													m.G0 = v10 + int32(48)
													return
												}
											}
										}
									}
								}
							}
						}
					} else {
						v39 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(v10)+47)) = uint8(v39)
						v41 = *(*int32)(unsafe.Add(mBase, uint32(v18)+52))
						v46 = F_heap_form_tuple(m, v41, v10+int32(16), v10+int32(44))
						mBase = m.M
						v47 = m.ExcPending
						if v47 != 0 {
							return
						} else {
							F_CatalogTupleInsert(m, v18, v46)
							mBase = m.M
							v49 = m.ExcPending
							if v49 != 0 {
								return
							} else {
								F_pfree(m, v46)
								mBase = m.M
								v51 = m.ExcPending
								if v51 != 0 {
									return
								} else {
									if l4 != 0 {
										F_sequence_close(m, v18, int32(0))
										mBase = m.M
										v54 = m.ExcPending
										if v54 != 0 {
											return
										} else {
											m.G0 = v10 + int32(48)
											return
										}
									} else {
										F_sequence_close(m, v18, int32(3))
										mBase = m.M
										v57 = m.ExcPending
										if v57 != 0 {
											return
										} else {
											F_UnlockSharedObject(m, int32(6100), l0, int32(1))
											mBase = m.M
											v61 = m.ExcPending
											if v61 != 0 {
												return
											} else {
												m.G0 = v10 + int32(48)
												return
											}
										}
									}
								}
							}
						}
					}
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v68 = m.ExcPending
					if v68 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = l0
						*(*int32)(unsafe.Add(mBase, uint32(v10))) = l1
						F_errmsg_internal(m, int32(123621), v10)
						mBase = m.M
						v73 = m.ExcPending
						if v73 != 0 {
							return
						} else {
							F_errfinish(m, int32(520759), int32(285), int32(371664))
							mBase = m.M
							v78 = m.ExcPending
							if v78 != 0 {
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
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
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
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v97 int32
	_ = v97
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v176 int32
	_ = v176
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v218 int32
	_ = v218
	var v239 int32
	_ = v239
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v282 int32
	_ = v282
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v349 int32
	_ = v349
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v371 int32
	_ = v371
	var v391 int32
	_ = v391
	var v394 int32
	_ = v394
	var v412 int32
	_ = v412
	var v415 int32
	_ = v415
	var v435 int32
	_ = v435
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v599 int32
	_ = v599
	var v602 int32
	_ = v602
	var v622 int32
	_ = v622
	var v640 int32
	_ = v640
	var v659 int32
	_ = v659
	var v660 int32
	_ = v660
	var v663 int32
	_ = v663
	var v664 int32
	_ = v664
	var v685 int32
	_ = v685
	var v695 int32
	_ = v695
	var v699 int32
	_ = v699
	var v700 int32
	_ = v700
	var v701 int32
	_ = v701
	var v719 int32
	_ = v719
	var v721 int32
	_ = v721
	var v722 int32
	_ = v722
	var v769 int32
	_ = v769
	var v787 int32
	_ = v787
	var v805 int32
	_ = v805
	var v806 int32
	_ = v806
	var v808 int32
	_ = v808
	var v810 int32
	_ = v810
	var v813 int32
	_ = v813
	var v814 int32
	_ = v814
	var v835 int32
	_ = v835
	var v845 int32
	_ = v845
	var v849 int32
	_ = v849
	var v850 int32
	_ = v850
	var v868 int32
	_ = v868
	var v888 int32
	_ = v888
	var v890 int32
	_ = v890
	var v891 int32
	_ = v891
	var v940 int32
	_ = v940
	var v959 int32
	_ = v959
	var v978 int32
	_ = v978
	var v998 int32
	_ = v998
	var v1019 int32
	_ = v1019
	var v1041 int32
	_ = v1041
	var v1061 int32
	_ = v1061
	var v1063 int32
	_ = v1063
	var v1064 int32
	_ = v1064
	var v1081 int32
	_ = v1081
	var v1085 int32
	_ = v1085
	var v1086 int32
	_ = v1086
	var v1106 int32
	_ = v1106
	var v1125 int32
	_ = v1125
	var v1126 int32
	_ = v1126
	var v1129 int32
	_ = v1129
	var v1130 int32
	_ = v1130
	var v1151 int32
	_ = v1151
	var v1154 int32
	_ = v1154
	var v1161 int32
	_ = v1161
	var v1165 int32
	_ = v1165
	var v1166 int32
	_ = v1166
	var v1169 int32
	_ = v1169
	var v1172 int64
	_ = v1172
	var v1205 int32
	_ = v1205
	var v1224 int32
	_ = v1224
	var v1225 int32
	_ = v1225
	var v1247 int32
	_ = v1247
	var v1268 int32
	_ = v1268
	var v1269 int32
	_ = v1269
	var v1270 int32
	_ = v1270
	var v1272 int32
	_ = v1272
	var v1276 int32
	_ = v1276
	var v1278 int32
	_ = v1278
	var v1287 int32
	_ = v1287
	var v1288 int32
	_ = v1288
	var v1289 int32
	_ = v1289
	var v1290 int32
	_ = v1290
	var v1291 int32
	_ = v1291
	var v1292 int32
	_ = v1292
	var v1293 int32
	_ = v1293
	var v1294 int32
	_ = v1294
	var v1295 int32
	_ = v1295
	var v1296 int32
	_ = v1296
	var v1297 int32
	_ = v1297
	var v1298 int32
	_ = v1298
	var v1299 int32
	_ = v1299
	var v1300 int32
	_ = v1300
	var v1301 int32
	_ = v1301
	var v1303 int32
	_ = v1303
	var v1304 int32
	_ = v1304
	var v1309 int32
	_ = v1309
	var v1317 int32
	_ = v1317
	var v1318 int32
	_ = v1318
	var v1319 int32
	_ = v1319
	var v1340 int32
	_ = v1340
	var v1350 int32
	_ = v1350
	var v1354 int32
	_ = v1354
	var v1355 int32
	_ = v1355
	var v1358 int32
	_ = v1358
	var v1361 int64
	_ = v1361
	var v1385 int32
	_ = v1385
	var v1396 int32
	_ = v1396
	var v1415 int32
	_ = v1415
	var v1418 int32
	_ = v1418
	var v1419 int32
	_ = v1419
	var v1426 int32
	_ = v1426
	var v1427 int32
	_ = v1427
	var v1431 int32
	_ = v1431
	var v1432 int32
	_ = v1432
	var v1438 int32
	_ = v1438
	var v1449 int32
	_ = v1449
	var v1467 int32
	_ = v1467
	var v1512 int32
	_ = v1512
	var v1515 int32
	_ = v1515
	var v1536 int32
	_ = v1536
	var v1542 int32
	_ = v1542
	var v1543 int32
	_ = v1543
	var v1561 int32
	_ = v1561
	var v1584 int32
	_ = v1584
	var v1632 int32
	_ = v1632
	var v1651 int32
	_ = v1651
	var v1674 int32
	_ = v1674
	var v1699 int32
	_ = v1699
	var v1720 int32
	_ = v1720
	var v1746 int32
	_ = v1746
	var v1749 int32
	_ = v1749
	var v1750 int64
	_ = v1750
	var v1754 int32
	_ = v1754
	var v1756 int32
	_ = v1756
	var v1757 int32
	_ = v1757
	var v1761 int32
	_ = v1761
	var v1763 int32
	_ = v1763
	var v1764 int32
	_ = v1764
	var v1765 int32
	_ = v1765
	var v1766 int32
	_ = v1766
	var v1767 int32
	_ = v1767
	var v1768 int32
	_ = v1768
	var v1769 int32
	_ = v1769
	var v1770 int32
	_ = v1770
	var v1771 int32
	_ = v1771
	var v1772 int32
	_ = v1772
	var v1773 int32
	_ = v1773
	var v1774 int32
	_ = v1774
	var v1775 int32
	_ = v1775
	var v1776 int32
	_ = v1776
	var v1777 int32
	_ = v1777
	var v1778 int32
	_ = v1778
	var v1779 int32
	_ = v1779
	var v1780 int32
	_ = v1780
	var v1782 int32
	_ = v1782
	v3 = int32(0)
	v29 = m.G0
	v31 = v29 - int32(144)
	m.G0 = v31
	v37 = v3
	v38 = v3
	v39 = v3
	v40 = v3
	v41 = v3
	v42 = v3
	v43 = v3
	v44 = v3
	v45 = v3
	v46 = v3
	v47 = v3
	v48 = v3
	v49 = v3
	v50 = v3
	v51 = int32(-1)
	v52 = v3
	v53 = v3
	v54 = v3
	v59 = v31
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
	m.G0 = v31 + int32(144)
	return
L4:
	;
	goto L3
L5:
	;
	if v51 == int32(1) {
		goto L10
	} else {
		goto L11
	}
L7:
	;
	v1749 = int32(m.ExcTag)
	v1750 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v1749 == int32(0) {
		goto L145
	} else {
		goto L146
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+84)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v31)+80)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+88)) = v1085
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+95)) = uint8(v810)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+96)) = v808
	*(*int32)(unsafe.Add(mBase, uint32(v31)+100)) = v805
	*(*int32)(unsafe.Add(mBase, uint32(v31)+104)) = v575
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+111)) = uint8(v576)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+112)) = v308
	*(*int32)(unsafe.Add(mBase, uint32(v31)+116)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v31)+120)) = v83
	*(*int32)(unsafe.Add(mBase, uint32(v31)+124)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v31)+128)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v31)+132)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v31)+136)) = v71
	*(*int32)(unsafe.Add(mBase, uint32(v31)+140)) = v68
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1632 = m.ExcPending
	if v1632 != 0 {
		v1746 = v83
		goto L7
	} else {
		goto L140
	}
L9:
	;
	if v1304 == int32(0) {
		goto L118
	} else {
		goto L119
	}
L10:
	;
	v1287 = v37
	v1288 = v38
	v1289 = v39
	v1290 = v40
	v1291 = v41
	v1292 = v42
	v1293 = v43
	v1294 = v44
	v1295 = v45
	v1296 = v46
	v1297 = v47
	v1298 = v48
	v1299 = v49
	v1300 = v50
	v1301 = v52
	v1303 = v53
	v1304 = v54
	v1309 = v59
	goto L9
L11:
	;
	goto L12
L12:
	;
	v64 = int32(-64)
	v65 = v59 + v64
	m.G0 = v65
	v67 = int32(16)
	v68 = v65 - v67
	m.G0 = v68
	v71 = v68 - v67
	m.G0 = v71
	v74 = v71 + v64
	m.G0 = v74
	v77 = v74 - v67
	m.G0 = v77
	v80 = v77 - int32(160)
	m.G0 = v80
	v83 = v80 + v64
	m.G0 = v83
	*(*int32)(unsafe.Add(mBase, uint32(v77))) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+84)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v31)+80)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+88)) = v43
	v90 = int32(1)
	v91 = v50 & v90
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+95)) = uint8(v91)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+96)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v31)+100)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v31)+104)) = v52
	v97 = v53 & v90
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+111)) = uint8(v97)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+112)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v31)+116)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v31)+120)) = v83
	*(*int32)(unsafe.Add(mBase, uint32(v31)+124)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v31)+128)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v31)+132)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v31)+136)) = v71
	*(*int32)(unsafe.Add(mBase, uint32(v31)+140)) = v68
	v109 = F_table_open(m, int32(6100), int32(3))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		v1746 = v83
		goto L7
	} else {
		goto L13
	}
L13:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+80)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+84)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v31)+88)) = v43
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+95)) = uint8(v91)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+96)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v31)+100)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v31)+104)) = v52
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+111)) = uint8(v97)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+112)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v31)+116)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v31)+120)) = v83
	*(*int32)(unsafe.Add(mBase, uint32(v31)+124)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v31)+128)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v31)+132)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v31)+136)) = v71
	*(*int32)(unsafe.Add(mBase, uint32(v31)+140)) = v68
	v130 = *(*int32)(unsafe.Add(mBase, _consts[223]))
	v131 = F_SearchSysCache2(m, int32(66), v130, v111)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		v1746 = v83
		goto L7
	} else {
		goto L14
	}
L14:
	;
	if v131 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+84)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v31)+80)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+88)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v31)+96)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v31)+100)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v31)+104)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v31)+112)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v31)+116)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v31)+120)) = v83
	*(*int32)(unsafe.Add(mBase, uint32(v31)+124)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v31)+128)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v31)+132)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v31)+136)) = v71
	*(*int32)(unsafe.Add(mBase, uint32(v31)+140)) = v68
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+95)) = uint8(v91)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+111)) = uint8(v97)
	F_sequence_close(m, v109, int32(0))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		v1746 = v83
		goto L7
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v131)+16))
	v305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v304)+22)))
	v306 = v304 + v305
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v306)+80))
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v306)))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+84)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v31)+80)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+88)) = v43
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+95)) = uint8(v91)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+96)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v31)+100)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v31)+104)) = v52
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+111)) = uint8(v97)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+112)) = v308
	*(*int32)(unsafe.Add(mBase, uint32(v31)+116)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v31)+120)) = v83
	*(*int32)(unsafe.Add(mBase, uint32(v31)+124)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v31)+128)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v31)+132)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v31)+136)) = v71
	*(*int32)(unsafe.Add(mBase, uint32(v31)+140)) = v68
	v325 = F_superuser_arg(m, v307)
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		v1746 = v83
		goto L7
	} else {
		goto L30
	}
L18:
	;
	v154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
	if v154 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+84)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v31)+80)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+88)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v31)+96)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v31)+100)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v31)+104)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v31)+112)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v31)+116)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v31)+120)) = v83
	*(*int32)(unsafe.Add(mBase, uint32(v31)+124)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v31)+128)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v31)+132)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v31)+136)) = v71
	*(*int32)(unsafe.Add(mBase, uint32(v31)+140)) = v68
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+95)) = uint8(v91)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+111)) = uint8(v97)
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		v1746 = v83
		goto L7
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+84)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v31)+80)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+88)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v31)+96)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v31)+100)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v31)+104)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v31)+112)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v31)+116)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v31)+120)) = v83
	*(*int32)(unsafe.Add(mBase, uint32(v31)+124)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v31)+128)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v31)+132)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v31)+136)) = v71
	*(*int32)(unsafe.Add(mBase, uint32(v31)+140)) = v68
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+95)) = uint8(v91)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+111)) = uint8(v97)
	v258 = F_errstart(m, int32(18), int32(0))
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		v1746 = v83
		goto L7
	} else {
		goto L26
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+84)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v31)+80)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+88)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v31)+96)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v31)+100)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v31)+104)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v31)+112)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v31)+116)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v31)+120)) = v83
	*(*int32)(unsafe.Add(mBase, uint32(v31)+124)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v31)+128)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v31)+132)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v31)+136)) = v71
	*(*int32)(unsafe.Add(mBase, uint32(v31)+140)) = v68
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+95)) = uint8(v91)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+111)) = uint8(v97)
	F_errcode(m, int32(67137668))
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		v1746 = v83
		goto L7
	} else {
		goto L23
	}
L23:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+84)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v31)+80)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+88)) = v43
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+95)) = uint8(v91)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+96)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v31)+100)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v31)+104)) = v52
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+111)) = uint8(v97)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+112)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v31)+116)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v31)+120)) = v83
	*(*int32)(unsafe.Add(mBase, uint32(v31)+124)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v31)+128)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v31)+132)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v31)+136)) = v71
	*(*int32)(unsafe.Add(mBase, uint32(v31)+140)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v31)+16)) = v196
	F_errmsg(m, int32(77232), v31+int32(16))
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		v1746 = v83
		goto L7
	} else {
		goto L24
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+84)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v31)+80)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+88)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v31)+96)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v31)+100)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v31)+104)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v31)+112)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v31)+116)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v31)+120)) = v83
	*(*int32)(unsafe.Add(mBase, uint32(v31)+124)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v31)+128)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v31)+132)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v31)+136)) = v71
	*(*int32)(unsafe.Add(mBase, uint32(v31)+140)) = v68
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+95)) = uint8(v91)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+111)) = uint8(v97)
	F_errfinish(m, int32(518572), int32(1666), int32(260247))
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		v1746 = v83
		goto L7
	} else {
		goto L25
	}
L25:
	;
	goto L1
L26:
	;
	if v258 == int32(0) {
		goto L4
	} else {
		goto L27
	}
L27:
	;
	v262 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+84)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v31)+80)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+88)) = v43
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+95)) = uint8(v91)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+96)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v31)+100)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v31)+104)) = v52
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+111)) = uint8(v97)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+112)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v31)+116)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v31)+120)) = v83
	*(*int32)(unsafe.Add(mBase, uint32(v31)+124)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v31)+128)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v31)+132)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v31)+136)) = v71
	*(*int32)(unsafe.Add(mBase, uint32(v31)+140)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v31))) = v262
	F_errmsg(m, int32(349164), v31)
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		v1746 = v83
		goto L7
	} else {
		goto L28
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+84)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v31)+80)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+88)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v31)+96)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v31)+100)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v31)+104)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v31)+112)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v31)+116)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v31)+120)) = v83
	*(*int32)(unsafe.Add(mBase, uint32(v31)+124)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v31)+128)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v31)+132)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v31)+136)) = v71
	*(*int32)(unsafe.Add(mBase, uint32(v31)+140)) = v68
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+95)) = uint8(v91)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+111)) = uint8(v97)
	F_errfinish(m, int32(518572), int32(1670), int32(260247))
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		v1746 = v83
		goto L7
	} else {
		goto L29
	}
L29:
	;
	goto L4
L30:
	;
	v327 = int32(0)
	if v325 == v327 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v330 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v306)+89)))
	v331 = v330
	goto L33
L32:
	;
	v331 = v327
	goto L33
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+84)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v31)+80)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+88)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v31)+96)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v31)+100)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v31)+104)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v31)+112)) = v308
	*(*int32)(unsafe.Add(mBase, uint32(v31)+116)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v31)+120)) = v83
	*(*int32)(unsafe.Add(mBase, uint32(v31)+124)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v31)+128)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v31)+132)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v31)+136)) = v71
	*(*int32)(unsafe.Add(mBase, uint32(v31)+140)) = v68
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+95)) = uint8(v91)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+111)) = uint8(v97)
	v349 = *(*int32)(unsafe.Add(mBase, _consts[4]))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+84)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v31)+80)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+88)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v31)+96)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v31)+100)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v31)+104)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v31)+112)) = v308
	*(*int32)(unsafe.Add(mBase, uint32(v31)+116)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v31)+120)) = v83
	*(*int32)(unsafe.Add(mBase, uint32(v31)+124)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v31)+128)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v31)+132)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v31)+136)) = v71
	*(*int32)(unsafe.Add(mBase, uint32(v31)+140)) = v68
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+95)) = uint8(v91)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+111)) = uint8(v97)
	v367 = F_object_ownercheck(m, int32(6100), v308, v349)
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		v1746 = v83
		goto L7
	} else {
		goto L34
	}
L34:
	;
	if v367 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v371 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+84)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v31)+80)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+88)) = v43
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+95)) = uint8(v91)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+96)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v31)+100)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v31)+104)) = v52
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+111)) = uint8(v97)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+112)) = v308
	*(*int32)(unsafe.Add(mBase, uint32(v31)+116)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v31)+120)) = v83
	*(*int32)(unsafe.Add(mBase, uint32(v31)+124)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v31)+128)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v31)+132)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v31)+136)) = v71
	*(*int32)(unsafe.Add(mBase, uint32(v31)+140)) = v68
	F_aclcheck_error(m, int32(2), int32(38), v371)
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		v1746 = v83
		goto L7
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	v394 = *(*int32)(unsafe.Add(mBase, _consts[203]))
	if v394 != 0 {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	goto L37
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+84)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v31)+80)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+88)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v31)+96)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v31)+100)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v31)+104)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v31)+112)) = v308
	*(*int32)(unsafe.Add(mBase, uint32(v31)+116)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v31)+120)) = v83
	*(*int32)(unsafe.Add(mBase, uint32(v31)+124)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v31)+128)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v31)+132)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v31)+136)) = v71
	*(*int32)(unsafe.Add(mBase, uint32(v31)+140)) = v68
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+95)) = uint8(v91)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+111)) = uint8(v97)
	v412 = int32(0)
	F_RunObjectDropHook(m, int32(6100), v308, v412, v412)
	mBase = m.M
	v415 = m.ExcPending
	if v415 != 0 {
		v1746 = v83
		goto L7
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+84)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v31)+80)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+88)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v31)+96)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v31)+100)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v31)+104)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v31)+112)) = v308
	*(*int32)(unsafe.Add(mBase, uint32(v31)+116)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v31)+120)) = v83
	*(*int32)(unsafe.Add(mBase, uint32(v31)+124)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v31)+128)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v31)+132)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v31)+136)) = v71
	*(*int32)(unsafe.Add(mBase, uint32(v31)+140)) = v68
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+95)) = uint8(v91)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+111)) = uint8(v97)
	F_LockSharedObject(m, int32(6100), v308, int32(8))
	mBase = m.M
	v435 = m.ExcPending
	if v435 != 0 {
		v1746 = v83
		goto L7
	} else {
		goto L43
	}
L42:
	;
	goto L41
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+84)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v31)+80)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+88)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v31)+96)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v31)+100)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v31)+104)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v31)+112)) = v308
	*(*int32)(unsafe.Add(mBase, uint32(v31)+116)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v31)+120)) = v83
	*(*int32)(unsafe.Add(mBase, uint32(v31)+124)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v31)+128)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v31)+132)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v31)+136)) = v71
	*(*int32)(unsafe.Add(mBase, uint32(v31)+140)) = v68
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+95)) = uint8(v91)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+111)) = uint8(v97)
	v454 = F_SysCacheGetAttrNotNull(m, int32(67), v131, int32(4))
	mBase = m.M
	v455 = m.ExcPending
	if v455 != 0 {
		v1746 = v83
		goto L7
	} else {
		goto L44
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+84)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v31)+80)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+88)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v31)+96)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v31)+100)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v31)+104)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v31)+112)) = v308
	*(*int32)(unsafe.Add(mBase, uint32(v31)+116)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v31)+120)) = v83
	*(*int32)(unsafe.Add(mBase, uint32(v31)+124)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v31)+128)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v31)+132)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v31)+136)) = v71
	*(*int32)(unsafe.Add(mBase, uint32(v31)+140)) = v68
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+95)) = uint8(v91)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+111)) = uint8(v97)
	v472 = F_pstrdup(m, v454)
	mBase = m.M
	v473 = m.ExcPending
	if v473 != 0 {
		v1746 = v83
		goto L7
	} else {
		goto L45
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+84)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v31)+80)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+88)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v31)+96)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v31)+100)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v31)+104)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v31)+112)) = v308
	*(*int32)(unsafe.Add(mBase, uint32(v31)+116)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v31)+120)) = v83
	*(*int32)(unsafe.Add(mBase, uint32(v31)+124)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v31)+128)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v31)+132)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v31)+136)) = v71
	*(*int32)(unsafe.Add(mBase, uint32(v31)+140)) = v68
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+95)) = uint8(v91)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+111)) = uint8(v97)
	v492 = F_SysCacheGetAttrNotNull(m, int32(67), v131, int32(14))
	mBase = m.M
	v493 = m.ExcPending
	if v493 != 0 {
		v1746 = v83
		goto L7
	} else {
		goto L46
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+84)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v31)+80)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+88)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v31)+96)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v31)+100)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v31)+104)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v31)+112)) = v308
	*(*int32)(unsafe.Add(mBase, uint32(v31)+116)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v31)+120)) = v83
	*(*int32)(unsafe.Add(mBase, uint32(v31)+124)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v31)+128)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v31)+132)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v31)+136)) = v71
	*(*int32)(unsafe.Add(mBase, uint32(v31)+140)) = v68
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+95)) = uint8(v91)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+111)) = uint8(v97)
	v510 = F_text_to_cstring(m, v492)
	mBase = m.M
	v511 = m.ExcPending
	if v511 != 0 {
		v1746 = v83
		goto L7
	} else {
		goto L47
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+84)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v31)+80)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+88)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v31)+96)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v31)+100)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v31)+104)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v31)+112)) = v308
	*(*int32)(unsafe.Add(mBase, uint32(v31)+116)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v31)+120)) = v83
	*(*int32)(unsafe.Add(mBase, uint32(v31)+124)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v31)+128)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v31)+132)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v31)+136)) = v71
	*(*int32)(unsafe.Add(mBase, uint32(v31)+140)) = v68
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+95)) = uint8(v91)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+111)) = uint8(v97)
	v530 = F_SysCacheGetAttr(m, int32(67), v131, int32(15), v71)
	mBase = m.M
	v531 = m.ExcPending
	if v531 != 0 {
		v1746 = v83
		goto L7
	} else {
		goto L48
	}
L48:
	;
	v532 = int32(0)
	v533 = int32(1)
	v534 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71))))
	if v534 != 0 {
		v575 = v532
		v576 = v533
		goto L49
	} else {
		goto L50
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v68)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v68)+4)) = v308
	*(*int32)(unsafe.Add(mBase, uint32(v68))) = int32(6100)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+84)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v31)+80)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+88)) = v43
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+95)) = uint8(v91)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+96)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v31)+100)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v31)+104)) = v575
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+111)) = uint8(v576)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+112)) = v308
	*(*int32)(unsafe.Add(mBase, uint32(v31)+116)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v31)+120)) = v83
	*(*int32)(unsafe.Add(mBase, uint32(v31)+124)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v31)+128)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v31)+132)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v31)+136)) = v71
	*(*int32)(unsafe.Add(mBase, uint32(v31)+140)) = v68
	v599 = int32(1)
	F_EventTriggerSQLDropAddObject(m, v68, v599, v599)
	mBase = m.M
	v602 = m.ExcPending
	if v602 != 0 {
		v1746 = v83
		goto L7
	} else {
		goto L54
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+84)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v31)+80)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+88)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v31)+96)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v31)+100)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v31)+104)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v31)+112)) = v308
	*(*int32)(unsafe.Add(mBase, uint32(v31)+116)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v31)+120)) = v83
	*(*int32)(unsafe.Add(mBase, uint32(v31)+124)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v31)+128)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v31)+132)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v31)+136)) = v71
	*(*int32)(unsafe.Add(mBase, uint32(v31)+140)) = v68
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+95)) = uint8(v91)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+111)) = uint8(v97)
	v551 = F_pstrdup(m, v530)
	mBase = m.M
	v552 = m.ExcPending
	if v552 != 0 {
		v1746 = v83
		goto L7
	} else {
		goto L51
	}
L51:
	;
	if v551 == int32(0) {
		v575 = v532
		v576 = v533
		goto L49
	} else {
		goto L52
	}
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+84)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v31)+80)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+88)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v31)+96)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v31)+100)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v31)+104)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v31)+112)) = v308
	*(*int32)(unsafe.Add(mBase, uint32(v31)+116)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v31)+120)) = v83
	*(*int32)(unsafe.Add(mBase, uint32(v31)+124)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v31)+128)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v31)+132)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v31)+136)) = v71
	*(*int32)(unsafe.Add(mBase, uint32(v31)+140)) = v68
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+95)) = uint8(v91)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+111)) = uint8(v97)
	F_PreventInTransactionBlock(m, l1, int32(555109))
	mBase = m.M
	v574 = m.ExcPending
	if v574 != 0 {
		v1746 = v83
		goto L7
	} else {
		goto L53
	}
L53:
	;
	v575 = v551
	v576 = int32(0)
	goto L49
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+84)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v31)+80)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+88)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v31)+96)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v31)+100)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v31)+104)) = v575
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+111)) = uint8(v576)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+112)) = v308
	*(*int32)(unsafe.Add(mBase, uint32(v31)+116)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v31)+120)) = v83
	*(*int32)(unsafe.Add(mBase, uint32(v31)+124)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v31)+128)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v31)+132)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v31)+136)) = v71
	*(*int32)(unsafe.Add(mBase, uint32(v31)+140)) = v68
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+95)) = uint8(v91)
	F_CatalogTupleDelete(m, v109, v131+int32(4))
	mBase = m.M
	v622 = m.ExcPending
	if v622 != 0 {
		v1746 = v83
		goto L7
	} else {
		goto L55
	}
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+84)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v31)+80)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+88)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v31)+96)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v31)+100)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v31)+104)) = v575
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+111)) = uint8(v576)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+112)) = v308
	*(*int32)(unsafe.Add(mBase, uint32(v31)+116)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v31)+120)) = v83
	*(*int32)(unsafe.Add(mBase, uint32(v31)+124)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v31)+128)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v31)+132)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v31)+136)) = v71
	*(*int32)(unsafe.Add(mBase, uint32(v31)+140)) = v68
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+95)) = uint8(v91)
	F_ReleaseCatCache(m, v131)
	mBase = m.M
	v640 = m.ExcPending
	if v640 != 0 {
		v1746 = v83
		goto L7
	} else {
		goto L56
	}
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+84)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v31)+80)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+88)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v31)+96)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v31)+100)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v31)+104)) = v575
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+111)) = uint8(v576)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+112)) = v308
	*(*int32)(unsafe.Add(mBase, uint32(v31)+116)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v31)+120)) = v83
	*(*int32)(unsafe.Add(mBase, uint32(v31)+124)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v31)+128)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v31)+132)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v31)+136)) = v71
	*(*int32)(unsafe.Add(mBase, uint32(v31)+140)) = v68
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+95)) = uint8(v91)
	v659 = F_logicalrep_workers_find(m, v308, int32(0), int32(1))
	mBase = m.M
	v660 = m.ExcPending
	if v660 != 0 {
		v1746 = v83
		goto L7
	} else {
		goto L58
	}
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+84)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v31)+80)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+88)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v31)+96)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v31)+100)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v31)+104)) = v575
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+111)) = uint8(v576)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+112)) = v308
	*(*int32)(unsafe.Add(mBase, uint32(v31)+116)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v31)+120)) = v83
	*(*int32)(unsafe.Add(mBase, uint32(v31)+124)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v31)+128)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v31)+132)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v31)+136)) = v71
	*(*int32)(unsafe.Add(mBase, uint32(v31)+140)) = v68
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+95)) = uint8(v91)
	F_list_free(m, v659)
	mBase = m.M
	v769 = m.ExcPending
	if v769 != 0 {
		v1746 = v83
		goto L7
	} else {
		goto L65
	}
L58:
	;
	if v659 == int32(0) {
		goto L57
	} else {
		goto L59
	}
L59:
	;
	v663 = int32(0)
	v664 = *(*int32)(unsafe.Add(mBase, uint32(v659)+4))
	if v664 <= v663 {
		goto L57
	} else {
		goto L60
	}
L60:
	;
	v685 = v663
	goto L61
L61:
	;
	v695 = *(*int32)(unsafe.Add(mBase, uint32(v659)+12))
	v699 = *(*int32)(unsafe.Add(mBase, uint32(v695+v685<<(uint(int32(2))%32))))
	v700 = *(*int32)(unsafe.Add(mBase, uint32(v699)+36))
	v701 = *(*int32)(unsafe.Add(mBase, uint32(v699)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+84)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v31)+80)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+88)) = v43
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+95)) = uint8(v91)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+96)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v31)+100)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v31)+104)) = v575
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+111)) = uint8(v576)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+112)) = v308
	*(*int32)(unsafe.Add(mBase, uint32(v31)+116)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v31)+120)) = v83
	*(*int32)(unsafe.Add(mBase, uint32(v31)+124)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v31)+128)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v31)+132)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v31)+136)) = v71
	*(*int32)(unsafe.Add(mBase, uint32(v31)+140)) = v68
	F_logicalrep_worker_stop(m, v701, v700)
	mBase = m.M
	v719 = m.ExcPending
	if v719 != 0 {
		v1746 = v83
		goto L7
	} else {
		goto L63
	}
L62:
	;
	goto L57
L63:
	;
	v721 = v685 + int32(1)
	v722 = *(*int32)(unsafe.Add(mBase, uint32(v659)+4))
	if v721 < v722 {
		v685 = v721
		goto L61
	} else {
		goto L64
	}
L64:
	;
	goto L62
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+84)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v31)+80)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+88)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v31)+96)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v31)+100)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v31)+104)) = v575
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+111)) = uint8(v576)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+112)) = v308
	*(*int32)(unsafe.Add(mBase, uint32(v31)+116)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v31)+120)) = v83
	*(*int32)(unsafe.Add(mBase, uint32(v31)+124)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v31)+128)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v31)+132)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v31)+136)) = v71
	*(*int32)(unsafe.Add(mBase, uint32(v31)+140)) = v68
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+95)) = uint8(v91)
	F_ApplyLauncherForgetWorkerStartTime(m, v308)
	mBase = m.M
	v787 = m.ExcPending
	if v787 != 0 {
		v1746 = v83
		goto L7
	} else {
		goto L66
	}
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+84)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v31)+80)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+88)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v31)+96)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v31)+100)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v31)+104)) = v575
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+111)) = uint8(v576)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+112)) = v308
	*(*int32)(unsafe.Add(mBase, uint32(v31)+116)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v31)+120)) = v83
	*(*int32)(unsafe.Add(mBase, uint32(v31)+124)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v31)+128)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v31)+132)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v31)+136)) = v71
	*(*int32)(unsafe.Add(mBase, uint32(v31)+140)) = v68
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+95)) = uint8(v91)
	v805 = F_GetSubscriptionRelations(m, v308, int32(1))
	mBase = m.M
	v806 = m.ExcPending
	if v806 != 0 {
		v1746 = v83
		goto L7
	} else {
		goto L67
	}
L67:
	;
	v808 = v805 + int32(4)
	v810 = base.B2i32(v805 == int32(0))
	if v805 == int32(0) {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+84)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v31)+80)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+88)) = v43
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+95)) = uint8(v810)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+96)) = v808
	*(*int32)(unsafe.Add(mBase, uint32(v31)+100)) = v805
	*(*int32)(unsafe.Add(mBase, uint32(v31)+104)) = v575
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+111)) = uint8(v576)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+112)) = v308
	*(*int32)(unsafe.Add(mBase, uint32(v31)+116)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v31)+120)) = v83
	*(*int32)(unsafe.Add(mBase, uint32(v31)+124)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v31)+128)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v31)+132)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v31)+136)) = v71
	*(*int32)(unsafe.Add(mBase, uint32(v31)+140)) = v68
	F_deleteSharedDependencyRecordsFor(m, int32(6100), v308, int32(0))
	mBase = m.M
	v940 = m.ExcPending
	if v940 != 0 {
		v1746 = v83
		goto L7
	} else {
		goto L79
	}
L69:
	;
	v813 = int32(0)
	v814 = *(*int32)(unsafe.Add(mBase, uint32(v805)+4))
	if v814 <= v813 {
		goto L68
	} else {
		goto L70
	}
L70:
	;
	v835 = v813
	goto L71
L71:
	;
	v845 = *(*int32)(unsafe.Add(mBase, uint32(v805)+12))
	v849 = *(*int32)(unsafe.Add(mBase, uint32(v845+v835<<(uint(int32(2))%32))))
	v850 = *(*int32)(unsafe.Add(mBase, uint32(v849)))
	if v850 != 0 {
		goto L73
	} else {
		goto L74
	}
L72:
	;
	goto L68
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+84)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v31)+80)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+88)) = v43
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+95)) = uint8(v810)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+96)) = v808
	*(*int32)(unsafe.Add(mBase, uint32(v31)+100)) = v805
	*(*int32)(unsafe.Add(mBase, uint32(v31)+104)) = v575
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+111)) = uint8(v576)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+112)) = v308
	*(*int32)(unsafe.Add(mBase, uint32(v31)+116)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v31)+120)) = v83
	*(*int32)(unsafe.Add(mBase, uint32(v31)+124)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v31)+128)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v31)+132)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v31)+136)) = v71
	*(*int32)(unsafe.Add(mBase, uint32(v31)+140)) = v68
	F_ReplicationOriginNameForLogicalRep(m, v308, v850, v74)
	mBase = m.M
	v868 = m.ExcPending
	if v868 != 0 {
		v1746 = v83
		goto L7
	} else {
		goto L76
	}
L74:
	;
	goto L75
L75:
	;
	v890 = v835 + int32(1)
	v891 = *(*int32)(unsafe.Add(mBase, uint32(v808)))
	if v890 < v891 {
		v835 = v890
		goto L71
	} else {
		goto L78
	}
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+84)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v31)+80)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+88)) = v43
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+95)) = uint8(v810)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+96)) = v808
	*(*int32)(unsafe.Add(mBase, uint32(v31)+100)) = v805
	*(*int32)(unsafe.Add(mBase, uint32(v31)+104)) = v575
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+111)) = uint8(v576)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+112)) = v308
	*(*int32)(unsafe.Add(mBase, uint32(v31)+116)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v31)+120)) = v83
	*(*int32)(unsafe.Add(mBase, uint32(v31)+124)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v31)+128)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v31)+132)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v31)+136)) = v71
	*(*int32)(unsafe.Add(mBase, uint32(v31)+140)) = v68
	F_replorigin_drop_by_name(m, v74, int32(1), int32(0))
	mBase = m.M
	v888 = m.ExcPending
	if v888 != 0 {
		v1746 = v83
		goto L7
	} else {
		goto L77
	}
L77:
	;
	goto L75
L78:
	;
	goto L72
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+84)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v31)+80)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+88)) = v43
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+95)) = uint8(v810)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+96)) = v808
	*(*int32)(unsafe.Add(mBase, uint32(v31)+100)) = v805
	*(*int32)(unsafe.Add(mBase, uint32(v31)+104)) = v575
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+111)) = uint8(v576)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+112)) = v308
	*(*int32)(unsafe.Add(mBase, uint32(v31)+116)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v31)+120)) = v83
	*(*int32)(unsafe.Add(mBase, uint32(v31)+124)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v31)+128)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v31)+132)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v31)+136)) = v71
	*(*int32)(unsafe.Add(mBase, uint32(v31)+140)) = v68
	F_RemoveSubscriptionRel(m, v308, int32(0))
	mBase = m.M
	v959 = m.ExcPending
	if v959 != 0 {
		v1746 = v83
		goto L7
	} else {
		goto L80
	}
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+84)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v31)+80)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+88)) = v43
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+95)) = uint8(v810)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+96)) = v808
	*(*int32)(unsafe.Add(mBase, uint32(v31)+100)) = v805
	*(*int32)(unsafe.Add(mBase, uint32(v31)+104)) = v575
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+111)) = uint8(v576)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+112)) = v308
	*(*int32)(unsafe.Add(mBase, uint32(v31)+116)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v31)+120)) = v83
	*(*int32)(unsafe.Add(mBase, uint32(v31)+124)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v31)+128)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v31)+132)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v31)+136)) = v71
	*(*int32)(unsafe.Add(mBase, uint32(v31)+140)) = v68
	F_ReplicationOriginNameForLogicalRep(m, v308, int32(0), v74)
	mBase = m.M
	v978 = m.ExcPending
	if v978 != 0 {
		v1746 = v83
		goto L7
	} else {
		goto L81
	}
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+84)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v31)+80)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+88)) = v43
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+95)) = uint8(v810)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+96)) = v808
	*(*int32)(unsafe.Add(mBase, uint32(v31)+100)) = v805
	*(*int32)(unsafe.Add(mBase, uint32(v31)+104)) = v575
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+111)) = uint8(v576)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+112)) = v308
	*(*int32)(unsafe.Add(mBase, uint32(v31)+116)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v31)+120)) = v83
	*(*int32)(unsafe.Add(mBase, uint32(v31)+124)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v31)+128)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v31)+132)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v31)+136)) = v71
	*(*int32)(unsafe.Add(mBase, uint32(v31)+140)) = v68
	F_replorigin_drop_by_name(m, v74, int32(1), int32(0))
	mBase = m.M
	v998 = m.ExcPending
	if v998 != 0 {
		v1746 = v83
		goto L7
	} else {
		goto L82
	}
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+84)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v31)+80)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+88)) = v43
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+95)) = uint8(v810)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+96)) = v808
	*(*int32)(unsafe.Add(mBase, uint32(v31)+100)) = v805
	*(*int32)(unsafe.Add(mBase, uint32(v31)+104)) = v575
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+111)) = uint8(v576)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+112)) = v308
	*(*int32)(unsafe.Add(mBase, uint32(v31)+116)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v31)+120)) = v83
	*(*int32)(unsafe.Add(mBase, uint32(v31)+124)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v31)+128)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v31)+132)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v31)+136)) = v71
	*(*int32)(unsafe.Add(mBase, uint32(v31)+140)) = v68
	F_pgstat_drop_transactional(m, int32(5), int32(0), base.I64_extend_i32_u(v308))
	mBase = m.M
	v1019 = m.ExcPending
	if v1019 != 0 {
		v1746 = v83
		goto L7
	} else {
		goto L83
	}
L83:
	;
	if base.B2i32(v805 == int32(0))&v576 != 0 {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+84)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v31)+80)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+88)) = v43
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+95)) = uint8(v810)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+96)) = v808
	*(*int32)(unsafe.Add(mBase, uint32(v31)+100)) = v805
	*(*int32)(unsafe.Add(mBase, uint32(v31)+104)) = v575
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+111)) = uint8(v576)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+112)) = v308
	*(*int32)(unsafe.Add(mBase, uint32(v31)+116)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v31)+120)) = v83
	*(*int32)(unsafe.Add(mBase, uint32(v31)+124)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v31)+128)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v31)+132)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v31)+136)) = v71
	*(*int32)(unsafe.Add(mBase, uint32(v31)+140)) = v68
	F_sequence_close(m, v109, int32(0))
	mBase = m.M
	v1041 = m.ExcPending
	if v1041 != 0 {
		v1746 = v83
		goto L7
	} else {
		goto L87
	}
L85:
	;
	goto L86
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+84)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v31)+80)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+88)) = v43
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+95)) = uint8(v810)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+96)) = v808
	*(*int32)(unsafe.Add(mBase, uint32(v31)+100)) = v805
	*(*int32)(unsafe.Add(mBase, uint32(v31)+104)) = v575
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+111)) = uint8(v576)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+112)) = v308
	*(*int32)(unsafe.Add(mBase, uint32(v31)+116)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v31)+120)) = v83
	*(*int32)(unsafe.Add(mBase, uint32(v31)+124)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v31)+128)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v31)+132)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v31)+136)) = v71
	*(*int32)(unsafe.Add(mBase, uint32(v31)+140)) = v68
	F_load_file(m, int32(226143), int32(0))
	mBase = m.M
	v1061 = m.ExcPending
	if v1061 != 0 {
		v1746 = v83
		goto L7
	} else {
		goto L88
	}
L87:
	;
	goto L4
L88:
	;
	v1063 = *(*int32)(unsafe.Add(mBase, _consts[290]))
	v1064 = *(*int32)(unsafe.Add(mBase, uint32(v1063)))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+80)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+84)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v31)+88)) = v43
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+95)) = uint8(v810)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+96)) = v808
	*(*int32)(unsafe.Add(mBase, uint32(v31)+100)) = v805
	*(*int32)(unsafe.Add(mBase, uint32(v31)+104)) = v575
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+111)) = uint8(v576)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+112)) = v308
	*(*int32)(unsafe.Add(mBase, uint32(v31)+116)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v31)+120)) = v83
	*(*int32)(unsafe.Add(mBase, uint32(v31)+124)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v31)+128)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v31)+132)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v31)+136)) = v71
	*(*int32)(unsafe.Add(mBase, uint32(v31)+140)) = v68
	v1081 = int32(1)
	v1085 = m.T0[v1064].(func(*base.Module, int32, int32, int32, int32, int32, int32) int32)(m, v510, v1081, v1081, v331&v1081, v472, v77)
	mBase = m.M
	v1086 = m.ExcPending
	if v1086 != 0 {
		v1746 = v83
		goto L7
	} else {
		goto L89
	}
L89:
	;
	if v1085 == int32(0) {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	if v576 != 0 {
		goto L93
	} else {
		goto L94
	}
L91:
	;
	goto L92
L92:
	;
	v1276 = *(*int32)(unsafe.Add(mBase, _consts[74]))
	v1278 = *(*int32)(unsafe.Add(mBase, _consts[291]))
	goto L113
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+84)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v31)+80)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+88)) = v1085
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+95)) = uint8(v810)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+96)) = v808
	*(*int32)(unsafe.Add(mBase, uint32(v31)+100)) = v805
	*(*int32)(unsafe.Add(mBase, uint32(v31)+104)) = v575
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+111)) = uint8(v576)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+112)) = v308
	*(*int32)(unsafe.Add(mBase, uint32(v31)+116)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v31)+120)) = v83
	*(*int32)(unsafe.Add(mBase, uint32(v31)+124)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v31)+128)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v31)+132)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v31)+136)) = v71
	*(*int32)(unsafe.Add(mBase, uint32(v31)+140)) = v68
	F_list_free(m, v805)
	mBase = m.M
	v1106 = m.ExcPending
	if v1106 != 0 {
		v1746 = v83
		goto L7
	} else {
		goto L96
	}
L94:
	;
	goto L95
L95:
	;
	v1126 = *(*int32)(unsafe.Add(mBase, uint32(v77)))
	if v805 == int32(0) {
		goto L8
	} else {
		goto L98
	}
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+84)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v31)+80)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+88)) = v1085
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+95)) = uint8(v810)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+96)) = v808
	*(*int32)(unsafe.Add(mBase, uint32(v31)+100)) = v805
	*(*int32)(unsafe.Add(mBase, uint32(v31)+104)) = v575
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+111)) = uint8(v576)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+112)) = v308
	*(*int32)(unsafe.Add(mBase, uint32(v31)+116)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v31)+120)) = v83
	*(*int32)(unsafe.Add(mBase, uint32(v31)+124)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v31)+128)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v31)+132)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v31)+136)) = v71
	*(*int32)(unsafe.Add(mBase, uint32(v31)+140)) = v68
	F_sequence_close(m, v109, int32(0))
	mBase = m.M
	v1125 = m.ExcPending
	if v1125 != 0 {
		v1746 = v83
		goto L7
	} else {
		goto L97
	}
L97:
	;
	goto L4
L98:
	;
	v1129 = int32(0)
	v1130 = *(*int32)(unsafe.Add(mBase, uint32(v805)+4))
	if v1130 <= v1129 {
		goto L8
	} else {
		goto L99
	}
L99:
	;
	v1151 = v1129
	v1154 = v1130
	goto L100
L100:
	;
	v1161 = *(*int32)(unsafe.Add(mBase, uint32(v805)+12))
	v1165 = *(*int32)(unsafe.Add(mBase, uint32(v1161+v1151<<(uint(int32(2))%32))))
	v1166 = *(*int32)(unsafe.Add(mBase, uint32(v1165)))
	if v1166 == int32(0) {
		v1270 = v1154
		goto L102
	} else {
		goto L103
	}
L101:
	;
	goto L8
L102:
	;
	v1272 = v1151 + int32(1)
	if v1272 < v1270 {
		v1151 = v1272
		v1154 = v1270
		goto L100
	} else {
		goto L112
	}
L103:
	;
	v1169 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1165)+16)))
	if v1169 == int32(115) {
		v1270 = v1154
		goto L102
	} else {
		goto L104
	}
L104:
	;
	v1172 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v65))) = v1172
	*(*int64)(unsafe.Add(mBase, uint32(v65)+56)) = v1172
	*(*int64)(unsafe.Add(mBase, uint32(v65)+48)) = v1172
	*(*int64)(unsafe.Add(mBase, uint32(v65)+40)) = v1172
	*(*int64)(unsafe.Add(mBase, uint32(v65)+32)) = v1172
	*(*int64)(unsafe.Add(mBase, uint32(v65)+24)) = v1172
	*(*int64)(unsafe.Add(mBase, uint32(v65)+16)) = v1172
	*(*int64)(unsafe.Add(mBase, uint32(v65)+8)) = v1172
	*(*int32)(unsafe.Add(mBase, uint32(v31)+84)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v31)+80)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+88)) = v1085
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+95)) = uint8(v810)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+96)) = v808
	*(*int32)(unsafe.Add(mBase, uint32(v31)+100)) = v805
	*(*int32)(unsafe.Add(mBase, uint32(v31)+104)) = v575
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+111)) = uint8(v576)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+112)) = v308
	*(*int32)(unsafe.Add(mBase, uint32(v31)+116)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v31)+120)) = v83
	*(*int32)(unsafe.Add(mBase, uint32(v31)+124)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v31)+128)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v31)+132)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v31)+136)) = v71
	*(*int32)(unsafe.Add(mBase, uint32(v31)+140)) = v68
	F_ReplicationSlotNameForTablesync(m, v308, v1166, v65)
	mBase = m.M
	v1205 = m.ExcPending
	if v1205 != 0 {
		v1746 = v83
		goto L7
	} else {
		goto L105
	}
L105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+84)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v31)+80)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+88)) = v1085
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+95)) = uint8(v810)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+96)) = v808
	*(*int32)(unsafe.Add(mBase, uint32(v31)+100)) = v805
	*(*int32)(unsafe.Add(mBase, uint32(v31)+104)) = v575
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+111)) = uint8(v576)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+112)) = v308
	*(*int32)(unsafe.Add(mBase, uint32(v31)+116)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v31)+120)) = v83
	*(*int32)(unsafe.Add(mBase, uint32(v31)+124)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v31)+128)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v31)+132)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v31)+136)) = v71
	*(*int32)(unsafe.Add(mBase, uint32(v31)+140)) = v68
	v1224 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v1225 = m.ExcPending
	if v1225 != 0 {
		v1746 = v83
		goto L7
	} else {
		goto L106
	}
L106:
	;
	if v1224 != 0 {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+84)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v31)+80)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+88)) = v1085
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+95)) = uint8(v810)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+96)) = v808
	*(*int32)(unsafe.Add(mBase, uint32(v31)+100)) = v805
	*(*int32)(unsafe.Add(mBase, uint32(v31)+104)) = v575
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+111)) = uint8(v576)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+112)) = v308
	*(*int32)(unsafe.Add(mBase, uint32(v31)+116)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v31)+120)) = v83
	*(*int32)(unsafe.Add(mBase, uint32(v31)+124)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v31)+128)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v31)+132)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v31)+136)) = v71
	*(*int32)(unsafe.Add(mBase, uint32(v31)+140)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v31)+64)) = v65
	F_errmsg_internal(m, int32(732320), v31-int32(-64))
	mBase = m.M
	v1247 = m.ExcPending
	if v1247 != 0 {
		v1746 = v83
		goto L7
	} else {
		goto L110
	}
L108:
	;
	goto L109
L109:
	;
	v1269 = *(*int32)(unsafe.Add(mBase, uint32(v808)))
	v1270 = v1269
	goto L102
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+84)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v31)+80)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+88)) = v1085
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+95)) = uint8(v810)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+96)) = v808
	*(*int32)(unsafe.Add(mBase, uint32(v31)+100)) = v805
	*(*int32)(unsafe.Add(mBase, uint32(v31)+104)) = v575
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+111)) = uint8(v576)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+112)) = v308
	*(*int32)(unsafe.Add(mBase, uint32(v31)+116)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v31)+120)) = v83
	*(*int32)(unsafe.Add(mBase, uint32(v31)+124)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v31)+128)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v31)+132)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v31)+136)) = v71
	*(*int32)(unsafe.Add(mBase, uint32(v31)+140)) = v68
	F_errfinish(m, int32(518572), int32(2342), int32(223701))
	mBase = m.M
	v1268 = m.ExcPending
	if v1268 != 0 {
		v1746 = v83
		goto L7
	} else {
		goto L111
	}
L111:
	;
	goto L109
L112:
	;
	goto L101
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v80)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v80))) = v31 + int32(76)
	goto L116
L114:
	;
	v1287 = v83
	v1288 = v308
	v1289 = v805
	v1290 = v109
	v1291 = v68
	v1292 = v74
	v1293 = v1085
	v1294 = v808
	v1295 = v77
	v1296 = v1278
	v1297 = v1276
	v1298 = v71
	v1299 = v80
	v1300 = v810
	v1301 = v575
	v1303 = v576
	v1304 = int32(0)
	v1309 = v83
	goto L9
L116:
	;
	goto L114
L117:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+84)) = v1296
	*(*int32)(unsafe.Add(mBase, uint32(v31)+80)) = v1297
	*(*int32)(unsafe.Add(mBase, uint32(v31)+88)) = v1293
	*(*int32)(unsafe.Add(mBase, uint32(v31)+96)) = v1294
	*(*int32)(unsafe.Add(mBase, uint32(v31)+100)) = v1289
	*(*int32)(unsafe.Add(mBase, uint32(v31)+104)) = v1301
	*(*int32)(unsafe.Add(mBase, uint32(v31)+112)) = v1288
	*(*int32)(unsafe.Add(mBase, uint32(v31)+116)) = v1290
	*(*int32)(unsafe.Add(mBase, uint32(v31)+120)) = v1287
	*(*int32)(unsafe.Add(mBase, uint32(v31)+124)) = v1299
	*(*int32)(unsafe.Add(mBase, uint32(v31)+128)) = v1295
	*(*int32)(unsafe.Add(mBase, uint32(v31)+132)) = v1292
	*(*int32)(unsafe.Add(mBase, uint32(v31)+136)) = v1298
	*(*int32)(unsafe.Add(mBase, uint32(v31)+140)) = v1291
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+95)) = uint8(v1317)
	v1512 = v1303 & int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+111)) = uint8(v1512)
	F_list_free(m, v1289)
	mBase = m.M
	v1515 = m.ExcPending
	if v1515 != 0 {
		v1746 = v1309
		goto L7
	} else {
		goto L133
	}
L118:
	;
	*(*int32)(unsafe.Add(mBase, _consts[291])) = v1299
	v1317 = v1300 & int32(1)
	if v1317 != 0 {
		goto L117
	} else {
		goto L121
	}
L119:
	;
	goto L120
L120:
	;
	*(*int32)(unsafe.Add(mBase, _consts[74])) = v1297
	*(*int32)(unsafe.Add(mBase, _consts[291])) = v1296
	v1426 = *(*int32)(unsafe.Add(mBase, _consts[290]))
	v1427 = *(*int32)(unsafe.Add(mBase, uint32(v1426)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+80)) = v1297
	*(*int32)(unsafe.Add(mBase, uint32(v31)+84)) = v1296
	*(*int32)(unsafe.Add(mBase, uint32(v31)+88)) = v1293
	v1431 = int32(1)
	v1432 = v1300 & v1431
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+95)) = uint8(v1432)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+96)) = v1294
	*(*int32)(unsafe.Add(mBase, uint32(v31)+100)) = v1289
	*(*int32)(unsafe.Add(mBase, uint32(v31)+104)) = v1301
	v1438 = v1303 & v1431
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+111)) = uint8(v1438)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+112)) = v1288
	*(*int32)(unsafe.Add(mBase, uint32(v31)+116)) = v1290
	*(*int32)(unsafe.Add(mBase, uint32(v31)+120)) = v1287
	*(*int32)(unsafe.Add(mBase, uint32(v31)+124)) = v1299
	*(*int32)(unsafe.Add(mBase, uint32(v31)+128)) = v1295
	*(*int32)(unsafe.Add(mBase, uint32(v31)+132)) = v1292
	*(*int32)(unsafe.Add(mBase, uint32(v31)+136)) = v1298
	*(*int32)(unsafe.Add(mBase, uint32(v31)+140)) = v1291
	m.T0[v1427].(func(*base.Module, int32))(m, v1293)
	mBase = m.M
	v1449 = m.ExcPending
	if v1449 != 0 {
		v1746 = v1309
		goto L7
	} else {
		goto L131
	}
L121:
	;
	v1318 = int32(0)
	v1319 = *(*int32)(unsafe.Add(mBase, uint32(v1294)))
	if v1319 <= v1318 {
		goto L117
	} else {
		goto L122
	}
L122:
	;
	v1340 = v1318
	goto L123
L123:
	;
	v1350 = *(*int32)(unsafe.Add(mBase, uint32(v1289)+12))
	v1354 = *(*int32)(unsafe.Add(mBase, uint32(v1350+v1340<<(uint(int32(2))%32))))
	v1355 = *(*int32)(unsafe.Add(mBase, uint32(v1354)))
	if v1355 == int32(0) {
		goto L125
	} else {
		goto L126
	}
L124:
	;
	goto L117
L125:
	;
	v1418 = v1340 + int32(1)
	v1419 = *(*int32)(unsafe.Add(mBase, uint32(v1294)))
	if v1418 < v1419 {
		v1340 = v1418
		goto L123
	} else {
		goto L130
	}
L126:
	;
	v1358 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1354)+16)))
	if v1358 == int32(115) {
		goto L125
	} else {
		goto L127
	}
L127:
	;
	v1361 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v1287))) = v1361
	*(*int64)(unsafe.Add(mBase, uint32(v1287)+56)) = v1361
	*(*int64)(unsafe.Add(mBase, uint32(v1287)+48)) = v1361
	*(*int64)(unsafe.Add(mBase, uint32(v1287)+40)) = v1361
	*(*int64)(unsafe.Add(mBase, uint32(v1287)+32)) = v1361
	*(*int64)(unsafe.Add(mBase, uint32(v1287)+24)) = v1361
	*(*int64)(unsafe.Add(mBase, uint32(v1287)+16)) = v1361
	*(*int64)(unsafe.Add(mBase, uint32(v1287)+8)) = v1361
	*(*int32)(unsafe.Add(mBase, uint32(v31)+84)) = v1296
	*(*int32)(unsafe.Add(mBase, uint32(v31)+80)) = v1297
	*(*int32)(unsafe.Add(mBase, uint32(v31)+88)) = v1293
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+95)) = uint8(v1317)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+96)) = v1294
	*(*int32)(unsafe.Add(mBase, uint32(v31)+100)) = v1289
	*(*int32)(unsafe.Add(mBase, uint32(v31)+104)) = v1301
	v1385 = v1303 & int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+111)) = uint8(v1385)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+112)) = v1288
	*(*int32)(unsafe.Add(mBase, uint32(v31)+116)) = v1290
	*(*int32)(unsafe.Add(mBase, uint32(v31)+120)) = v1287
	*(*int32)(unsafe.Add(mBase, uint32(v31)+124)) = v1299
	*(*int32)(unsafe.Add(mBase, uint32(v31)+128)) = v1295
	*(*int32)(unsafe.Add(mBase, uint32(v31)+132)) = v1292
	*(*int32)(unsafe.Add(mBase, uint32(v31)+136)) = v1298
	*(*int32)(unsafe.Add(mBase, uint32(v31)+140)) = v1291
	F_ReplicationSlotNameForTablesync(m, v1288, v1355, v1287)
	mBase = m.M
	v1396 = m.ExcPending
	if v1396 != 0 {
		v1746 = v1309
		goto L7
	} else {
		goto L128
	}
L128:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+84)) = v1296
	*(*int32)(unsafe.Add(mBase, uint32(v31)+80)) = v1297
	*(*int32)(unsafe.Add(mBase, uint32(v31)+88)) = v1293
	*(*int32)(unsafe.Add(mBase, uint32(v31)+96)) = v1294
	*(*int32)(unsafe.Add(mBase, uint32(v31)+100)) = v1289
	*(*int32)(unsafe.Add(mBase, uint32(v31)+104)) = v1301
	*(*int32)(unsafe.Add(mBase, uint32(v31)+112)) = v1288
	*(*int32)(unsafe.Add(mBase, uint32(v31)+116)) = v1290
	*(*int32)(unsafe.Add(mBase, uint32(v31)+120)) = v1287
	*(*int32)(unsafe.Add(mBase, uint32(v31)+124)) = v1299
	*(*int32)(unsafe.Add(mBase, uint32(v31)+128)) = v1295
	*(*int32)(unsafe.Add(mBase, uint32(v31)+132)) = v1292
	*(*int32)(unsafe.Add(mBase, uint32(v31)+136)) = v1298
	*(*int32)(unsafe.Add(mBase, uint32(v31)+140)) = v1291
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+95)) = uint8(v1317)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+111)) = uint8(v1385)
	F_ReplicationSlotDropAtPubNode(m, v1293, v1287, int32(1))
	mBase = m.M
	v1415 = m.ExcPending
	if v1415 != 0 {
		v1746 = v1309
		goto L7
	} else {
		goto L129
	}
L129:
	;
	goto L125
L130:
	;
	goto L124
L131:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+84)) = v1296
	*(*int32)(unsafe.Add(mBase, uint32(v31)+80)) = v1297
	*(*int32)(unsafe.Add(mBase, uint32(v31)+88)) = v1293
	*(*int32)(unsafe.Add(mBase, uint32(v31)+96)) = v1294
	*(*int32)(unsafe.Add(mBase, uint32(v31)+100)) = v1289
	*(*int32)(unsafe.Add(mBase, uint32(v31)+104)) = v1301
	*(*int32)(unsafe.Add(mBase, uint32(v31)+112)) = v1288
	*(*int32)(unsafe.Add(mBase, uint32(v31)+116)) = v1290
	*(*int32)(unsafe.Add(mBase, uint32(v31)+120)) = v1287
	*(*int32)(unsafe.Add(mBase, uint32(v31)+124)) = v1299
	*(*int32)(unsafe.Add(mBase, uint32(v31)+128)) = v1295
	*(*int32)(unsafe.Add(mBase, uint32(v31)+132)) = v1292
	*(*int32)(unsafe.Add(mBase, uint32(v31)+136)) = v1298
	*(*int32)(unsafe.Add(mBase, uint32(v31)+140)) = v1291
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+95)) = uint8(v1432)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+111)) = uint8(v1438)
	F_pg_re_throw(m)
	mBase = m.M
	v1467 = m.ExcPending
	if v1467 != 0 {
		v1746 = v1309
		goto L7
	} else {
		goto L132
	}
L132:
	;
	goto L1
L133:
	;
	if v1512 == int32(0) {
		goto L134
	} else {
		goto L135
	}
L134:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+84)) = v1296
	*(*int32)(unsafe.Add(mBase, uint32(v31)+80)) = v1297
	*(*int32)(unsafe.Add(mBase, uint32(v31)+88)) = v1293
	*(*int32)(unsafe.Add(mBase, uint32(v31)+96)) = v1294
	*(*int32)(unsafe.Add(mBase, uint32(v31)+100)) = v1289
	*(*int32)(unsafe.Add(mBase, uint32(v31)+104)) = v1301
	*(*int32)(unsafe.Add(mBase, uint32(v31)+112)) = v1288
	*(*int32)(unsafe.Add(mBase, uint32(v31)+116)) = v1290
	*(*int32)(unsafe.Add(mBase, uint32(v31)+120)) = v1287
	*(*int32)(unsafe.Add(mBase, uint32(v31)+124)) = v1299
	*(*int32)(unsafe.Add(mBase, uint32(v31)+128)) = v1295
	*(*int32)(unsafe.Add(mBase, uint32(v31)+132)) = v1292
	*(*int32)(unsafe.Add(mBase, uint32(v31)+136)) = v1298
	*(*int32)(unsafe.Add(mBase, uint32(v31)+140)) = v1291
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+95)) = uint8(v1317)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+111)) = uint8(v1512)
	F_ReplicationSlotDropAtPubNode(m, v1293, v1301, int32(0))
	mBase = m.M
	v1536 = m.ExcPending
	if v1536 != 0 {
		v1746 = v1309
		goto L7
	} else {
		goto L137
	}
L135:
	;
	goto L136
L136:
	;
	*(*int32)(unsafe.Add(mBase, _consts[74])) = v1297
	*(*int32)(unsafe.Add(mBase, _consts[291])) = v1296
	v1542 = *(*int32)(unsafe.Add(mBase, _consts[290]))
	v1543 = *(*int32)(unsafe.Add(mBase, uint32(v1542)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+80)) = v1297
	*(*int32)(unsafe.Add(mBase, uint32(v31)+84)) = v1296
	*(*int32)(unsafe.Add(mBase, uint32(v31)+88)) = v1293
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+95)) = uint8(v1317)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+96)) = v1294
	*(*int32)(unsafe.Add(mBase, uint32(v31)+100)) = v1289
	*(*int32)(unsafe.Add(mBase, uint32(v31)+104)) = v1301
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+111)) = uint8(v1512)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+112)) = v1288
	*(*int32)(unsafe.Add(mBase, uint32(v31)+116)) = v1290
	*(*int32)(unsafe.Add(mBase, uint32(v31)+120)) = v1287
	*(*int32)(unsafe.Add(mBase, uint32(v31)+124)) = v1299
	*(*int32)(unsafe.Add(mBase, uint32(v31)+128)) = v1295
	*(*int32)(unsafe.Add(mBase, uint32(v31)+132)) = v1292
	*(*int32)(unsafe.Add(mBase, uint32(v31)+136)) = v1298
	*(*int32)(unsafe.Add(mBase, uint32(v31)+140)) = v1291
	m.T0[v1543].(func(*base.Module, int32))(m, v1293)
	mBase = m.M
	v1561 = m.ExcPending
	if v1561 != 0 {
		v1746 = v1309
		goto L7
	} else {
		goto L138
	}
L137:
	;
	goto L136
L138:
	;
	*(*int32)(unsafe.Add(mBase, _consts[74])) = v1297
	*(*int32)(unsafe.Add(mBase, _consts[291])) = v1296
	*(*int32)(unsafe.Add(mBase, uint32(v31)+80)) = v1297
	*(*int32)(unsafe.Add(mBase, uint32(v31)+84)) = v1296
	*(*int32)(unsafe.Add(mBase, uint32(v31)+88)) = v1293
	*(*int32)(unsafe.Add(mBase, uint32(v31)+96)) = v1294
	*(*int32)(unsafe.Add(mBase, uint32(v31)+100)) = v1289
	*(*int32)(unsafe.Add(mBase, uint32(v31)+104)) = v1301
	*(*int32)(unsafe.Add(mBase, uint32(v31)+112)) = v1288
	*(*int32)(unsafe.Add(mBase, uint32(v31)+116)) = v1290
	*(*int32)(unsafe.Add(mBase, uint32(v31)+120)) = v1287
	*(*int32)(unsafe.Add(mBase, uint32(v31)+124)) = v1299
	*(*int32)(unsafe.Add(mBase, uint32(v31)+128)) = v1295
	*(*int32)(unsafe.Add(mBase, uint32(v31)+132)) = v1292
	*(*int32)(unsafe.Add(mBase, uint32(v31)+136)) = v1298
	*(*int32)(unsafe.Add(mBase, uint32(v31)+140)) = v1291
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+95)) = uint8(v1317)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+111)) = uint8(v1512)
	F_sequence_close(m, v1290, int32(0))
	mBase = m.M
	v1584 = m.ExcPending
	if v1584 != 0 {
		v1746 = v1309
		goto L7
	} else {
		goto L139
	}
L139:
	;
	goto L4
L140:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+84)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v31)+80)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+88)) = v1085
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+95)) = uint8(v810)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+96)) = v808
	*(*int32)(unsafe.Add(mBase, uint32(v31)+100)) = v805
	*(*int32)(unsafe.Add(mBase, uint32(v31)+104)) = v575
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+111)) = uint8(v576)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+112)) = v308
	*(*int32)(unsafe.Add(mBase, uint32(v31)+116)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v31)+120)) = v83
	*(*int32)(unsafe.Add(mBase, uint32(v31)+124)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v31)+128)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v31)+132)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v31)+136)) = v71
	*(*int32)(unsafe.Add(mBase, uint32(v31)+140)) = v68
	F_errcode(m, int32(100663808))
	mBase = m.M
	v1651 = m.ExcPending
	if v1651 != 0 {
		v1746 = v83
		goto L7
	} else {
		goto L141
	}
L141:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+84)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v31)+80)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+88)) = v1085
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+95)) = uint8(v810)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+96)) = v808
	*(*int32)(unsafe.Add(mBase, uint32(v31)+100)) = v805
	*(*int32)(unsafe.Add(mBase, uint32(v31)+104)) = v575
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+111)) = uint8(v576)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+112)) = v308
	*(*int32)(unsafe.Add(mBase, uint32(v31)+116)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v31)+120)) = v83
	*(*int32)(unsafe.Add(mBase, uint32(v31)+124)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v31)+128)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v31)+132)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v31)+136)) = v71
	*(*int32)(unsafe.Add(mBase, uint32(v31)+140)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v31)+52)) = v1126
	*(*int32)(unsafe.Add(mBase, uint32(v31)+48)) = v575
	F_errmsg(m, int32(215998), v31+int32(48))
	mBase = m.M
	v1674 = m.ExcPending
	if v1674 != 0 {
		v1746 = v83
		goto L7
	} else {
		goto L142
	}
L142:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+84)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v31)+80)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+88)) = v1085
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+95)) = uint8(v810)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+96)) = v808
	*(*int32)(unsafe.Add(mBase, uint32(v31)+100)) = v805
	*(*int32)(unsafe.Add(mBase, uint32(v31)+104)) = v575
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+111)) = uint8(v576)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+112)) = v308
	*(*int32)(unsafe.Add(mBase, uint32(v31)+116)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v31)+120)) = v83
	*(*int32)(unsafe.Add(mBase, uint32(v31)+124)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v31)+128)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v31)+132)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v31)+136)) = v71
	*(*int32)(unsafe.Add(mBase, uint32(v31)+140)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v31)+36)) = int32(713730)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+32)) = int32(567196)
	F_errhint(m, int32(607726), v31+int32(32))
	mBase = m.M
	v1699 = m.ExcPending
	if v1699 != 0 {
		v1746 = v83
		goto L7
	} else {
		goto L143
	}
L143:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+84)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v31)+80)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+88)) = v1085
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+95)) = uint8(v810)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+96)) = v808
	*(*int32)(unsafe.Add(mBase, uint32(v31)+100)) = v805
	*(*int32)(unsafe.Add(mBase, uint32(v31)+104)) = v575
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+111)) = uint8(v576)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+112)) = v308
	*(*int32)(unsafe.Add(mBase, uint32(v31)+116)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v31)+120)) = v83
	*(*int32)(unsafe.Add(mBase, uint32(v31)+124)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v31)+128)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v31)+132)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v31)+136)) = v71
	*(*int32)(unsafe.Add(mBase, uint32(v31)+140)) = v68
	F_errfinish(m, int32(518572), int32(2353), int32(223701))
	mBase = m.M
	v1720 = m.ExcPending
	if v1720 != 0 {
		v1746 = v83
		goto L7
	} else {
		goto L144
	}
L144:
	;
	goto L1
L145:
	;
	v1754 = int32(v1750)
	m.G0 = v1746
	v1756 = *(*int32)(unsafe.Add(mBase, uint32(v1754)+4))
	v1757 = *(*int32)(unsafe.Add(mBase, uint32(v1754)))
	v1761 = *(*int32)(unsafe.Add(mBase, uint32(v1757)))
	if v31+int32(76) == v1761 {
		goto L148
	} else {
		goto L149
	}
L146:
	;
	m.ExcPending = 1
	goto L154
L147:
	;
	if v1764 != 0 {
		goto L151
	} else {
		goto L152
	}
L148:
	;
	v1763 = *(*int32)(unsafe.Add(mBase, uint32(v1757)+4))
	v1764 = v1763
	goto L150
L149:
	;
	v1764 = int32(0)
	goto L150
L150:
	;
	goto L147
L151:
	;
	v1765 = *(*int32)(unsafe.Add(mBase, uint32(v31)+140))
	v1766 = *(*int32)(unsafe.Add(mBase, uint32(v31)+136))
	v1767 = *(*int32)(unsafe.Add(mBase, uint32(v31)+132))
	v1768 = *(*int32)(unsafe.Add(mBase, uint32(v31)+128))
	v1769 = *(*int32)(unsafe.Add(mBase, uint32(v31)+124))
	v1770 = *(*int32)(unsafe.Add(mBase, uint32(v31)+120))
	v1771 = *(*int32)(unsafe.Add(mBase, uint32(v31)+116))
	v1772 = *(*int32)(unsafe.Add(mBase, uint32(v31)+112))
	v1773 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+111)))
	v1774 = *(*int32)(unsafe.Add(mBase, uint32(v31)+104))
	v1775 = *(*int32)(unsafe.Add(mBase, uint32(v31)+100))
	v1776 = *(*int32)(unsafe.Add(mBase, uint32(v31)+96))
	v1777 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+95)))
	v1778 = *(*int32)(unsafe.Add(mBase, uint32(v31)+88))
	v1779 = *(*int32)(unsafe.Add(mBase, uint32(v31)+84))
	v1780 = *(*int32)(unsafe.Add(mBase, uint32(v31)+80))
	v37 = v1770
	v38 = v1772
	v39 = v1775
	v40 = v1771
	v41 = v1765
	v42 = v1767
	v43 = v1778
	v44 = v1776
	v45 = v1768
	v46 = v1779
	v47 = v1780
	v48 = v1766
	v49 = v1769
	v50 = v1777
	v51 = v1764
	v52 = v1774
	v53 = v1773
	v54 = v1756
	v59 = v1746
	goto L2
L152:
	;
	goto L153
L153:
	;
	F___wasm_longjmp(m, v1757, v1756)
	mBase = m.M
	v1782 = m.ExcPending
	if v1782 != 0 {
		goto L154
	} else {
		goto L155
	}
L154:
	;
	return
L155:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_GetSubscriptionRelations(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
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
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int64
	_ = v74
	var v75 int64
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	v10 = m.G0
	v12 = v10 - int32(112)
	m.G0 = v12
	v14 = int32(1)
	v17 = F_table_open(m, int32(6102), v14)
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
	F_ScanKeyInit(m, v12+int32(16), int32(1), int32(3), int32(184), l0)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	if l1 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v30 = int32(3)
	F_ScanKeyInit(m, v12-int32(-64), v30, v30, int32(70), int32(114))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L1
	} else {
		goto L7
	}
L5:
	;
	v37 = v14
	goto L6
L6:
	;
	v38 = int32(0)
	v44 = F_systable_beginscan(m, v17, v38, v38, v38, v37, v12+int32(16))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L1
	} else {
		goto L8
	}
L7:
	;
	v37 = int32(2)
	goto L6
L8:
	;
	v46 = F_systable_getnext(m, v44)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	if v46 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v48 = v46
	v49 = v38
	goto L13
L11:
	;
	v82 = v38
	goto L12
L12:
	;
	F_systable_endscan(m, v44)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L1
	} else {
		goto L23
	}
L13:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57)+22)))
	v60 = F_palloc(m, int32(24))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L1
	} else {
		goto L15
	}
L14:
	;
	v82 = v77
	goto L12
L15:
	;
	v62 = v57 + v58
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v60))) = v63
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62)+8)))
	*(*uint8)(unsafe.Add(mBase, uint32(v60)+16)) = uint8(v65)
	v71 = F_SysCacheGetAttr(m, int32(68), v48, int32(4), v12+int32(15))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+15)))
	if v73 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v75 = int64(0)
	goto L19
L18:
	;
	v74 = *(*int64)(unsafe.Add(mBase, uint32(v71)))
	v75 = v74
	goto L19
L19:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v60)+8)) = v75
	v77 = F_lappend(m, v49, v60)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	v79 = F_systable_getnext(m, v44)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	if v79 != 0 {
		v48 = v79
		v49 = v77
		goto L13
	} else {
		goto L22
	}
L22:
	;
	goto L14
L23:
	;
	F_sequence_close(m, v17, int32(1))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	m.G0 = v12 + int32(112)
	return v82
}
func F_parse_subscription_options(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v26 int64
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
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
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
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
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v117 int32
	_ = v117
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
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
	var v278 int32
	_ = v278
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v339 int32
	_ = v339
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v377 int32
	_ = v377
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v387 int32
	_ = v387
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v399 int32
	_ = v399
	var v402 int32
	_ = v402
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v422 int32
	_ = v422
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v432 int32
	_ = v432
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v460 int32
	_ = v460
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v470 int32
	_ = v470
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v498 int32
	_ = v498
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v508 int32
	_ = v508
	var v514 int32
	_ = v514
	var v516 int32
	_ = v516
	var v518 int32
	_ = v518
	var v521 int32
	_ = v521
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v547 int32
	_ = v547
	var v556 int32
	_ = v556
	var v559 int32
	_ = v559
	var v561 int32
	_ = v561
	var v570 int32
	_ = v570
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v591 int32
	_ = v591
	var v600 int32
	_ = v600
	var v603 int32
	_ = v603
	var v605 int32
	_ = v605
	var v614 int32
	_ = v614
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v635 int32
	_ = v635
	var v644 int32
	_ = v644
	var v647 int32
	_ = v647
	var v649 int32
	_ = v649
	var v658 int32
	_ = v658
	var v665 int32
	_ = v665
	var v666 int32
	_ = v666
	var v669 int32
	_ = v669
	var v670 int32
	_ = v670
	var v680 int32
	_ = v680
	var v689 int32
	_ = v689
	var v692 int32
	_ = v692
	var v694 int32
	_ = v694
	var v703 int32
	_ = v703
	var v709 int32
	_ = v709
	var v710 int32
	_ = v710
	var v713 int32
	_ = v713
	var v714 int32
	_ = v714
	var v724 int32
	_ = v724
	var v733 int32
	_ = v733
	var v736 int32
	_ = v736
	var v738 int32
	_ = v738
	var v747 int32
	_ = v747
	var v754 int32
	_ = v754
	var v757 int32
	_ = v757
	var v758 int32
	_ = v758
	var v762 int32
	_ = v762
	var v767 int32
	_ = v767
	var v772 int32
	_ = v772
	var v779 int32
	_ = v779
	var v780 int32
	_ = v780
	var v783 int32
	_ = v783
	var v784 int32
	_ = v784
	var v788 int32
	_ = v788
	var v789 int32
	_ = v789
	var v792 int32
	_ = v792
	var v793 int32
	_ = v793
	var v796 int32
	_ = v796
	var v803 int32
	_ = v803
	var v804 int32
	_ = v804
	var v806 int32
	_ = v806
	var v812 int32
	_ = v812
	var v813 int32
	_ = v813
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
	var v827 int32
	_ = v827
	var v830 int32
	_ = v830
	var v831 int32
	_ = v831
	var v834 int32
	_ = v834
	var v841 int32
	_ = v841
	var v842 int32
	_ = v842
	var v844 int32
	_ = v844
	var v850 int32
	_ = v850
	var v851 int32
	_ = v851
	var v855 int32
	_ = v855
	var v856 int32
	_ = v856
	var v859 int32
	_ = v859
	var v860 int32
	_ = v860
	var v864 int32
	_ = v864
	var v865 int32
	_ = v865
	var v868 int32
	_ = v868
	var v869 int32
	_ = v869
	var v872 int32
	_ = v872
	var v879 int32
	_ = v879
	var v880 int32
	_ = v880
	var v882 int32
	_ = v882
	var v888 int32
	_ = v888
	var v889 int32
	_ = v889
	var v893 int32
	_ = v893
	var v894 int32
	_ = v894
	var v897 int32
	_ = v897
	var v898 int32
	_ = v898
	var v902 int32
	_ = v902
	var v903 int32
	_ = v903
	var v906 int32
	_ = v906
	var v907 int32
	_ = v907
	var v910 int32
	_ = v910
	var v917 int32
	_ = v917
	var v918 int32
	_ = v918
	var v920 int32
	_ = v920
	var v926 int32
	_ = v926
	var v927 int32
	_ = v927
	var v931 int32
	_ = v931
	var v932 int32
	_ = v932
	var v935 int32
	_ = v935
	var v936 int32
	_ = v936
	var v940 int32
	_ = v940
	var v941 int32
	_ = v941
	var v944 int32
	_ = v944
	var v945 int32
	_ = v945
	var v948 int32
	_ = v948
	var v955 int32
	_ = v955
	var v956 int32
	_ = v956
	var v958 int32
	_ = v958
	var v964 int32
	_ = v964
	var v965 int32
	_ = v965
	var v967 int32
	_ = v967
	var v968 int32
	_ = v968
	var v971 int32
	_ = v971
	var v972 int32
	_ = v972
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
	var v1002 int32
	_ = v1002
	var v1003 int32
	_ = v1003
	var v1004 int32
	_ = v1004
	var v1009 int32
	_ = v1009
	var v1010 int32
	_ = v1010
	var v1013 int32
	_ = v1013
	var v1014 int32
	_ = v1014
	var v1024 int32
	_ = v1024
	var v1033 int32
	_ = v1033
	var v1036 int32
	_ = v1036
	var v1038 int32
	_ = v1038
	var v1047 int32
	_ = v1047
	var v1050 int32
	_ = v1050
	var v1054 int32
	_ = v1054
	var v1055 int32
	_ = v1055
	var v1058 int32
	_ = v1058
	var v1059 int32
	_ = v1059
	var v1069 int32
	_ = v1069
	var v1078 int32
	_ = v1078
	var v1081 int32
	_ = v1081
	var v1083 int32
	_ = v1083
	var v1092 int32
	_ = v1092
	var v1098 int32
	_ = v1098
	var v1101 int32
	_ = v1101
	var v1102 int32
	_ = v1102
	var v1108 int32
	_ = v1108
	var v1113 int32
	_ = v1113
	var v1116 int32
	_ = v1116
	var v1117 int32
	_ = v1117
	var v1120 int32
	_ = v1120
	var v1121 int32
	_ = v1121
	var v1125 int32
	_ = v1125
	var v1126 int32
	_ = v1126
	var v1129 int32
	_ = v1129
	var v1130 int32
	_ = v1130
	var v1133 int32
	_ = v1133
	var v1140 int32
	_ = v1140
	var v1141 int32
	_ = v1141
	var v1143 int32
	_ = v1143
	var v1144 int32
	_ = v1144
	var v1145 int32
	_ = v1145
	var v1148 int32
	_ = v1148
	var v1151 int32
	_ = v1151
	var v1152 int32
	_ = v1152
	var v1156 int32
	_ = v1156
	var v1157 int32
	_ = v1157
	var v1160 int32
	_ = v1160
	var v1161 int32
	_ = v1161
	var v1164 int32
	_ = v1164
	var v1171 int32
	_ = v1171
	var v1172 int32
	_ = v1172
	var v1179 int32
	_ = v1179
	var v1180 int32
	_ = v1180
	var v1181 int64
	_ = v1181
	var v1184 int32
	_ = v1184
	var v1185 int32
	_ = v1185
	var v1186 int64
	_ = v1186
	var v1194 int32
	_ = v1194
	var v1197 int32
	_ = v1197
	var v1198 int32
	_ = v1198
	var v1204 int32
	_ = v1204
	var v1209 int32
	_ = v1209
	var v1213 int32
	_ = v1213
	var v1214 int32
	_ = v1214
	var v1222 int32
	_ = v1222
	var v1223 int32
	_ = v1223
	var v1228 int32
	_ = v1228
	var v1231 int32
	_ = v1231
	var v1237 int32
	_ = v1237
	var v1242 int32
	_ = v1242
	var v1271 int32
	_ = v1271
	var v1274 int32
	_ = v1274
	var v1277 int32
	_ = v1277
	var v1285 int32
	_ = v1285
	var v1288 int32
	_ = v1288
	var v1297 int32
	_ = v1297
	var v1302 int32
	_ = v1302
	var v1303 int32
	_ = v1303
	var v1306 int32
	_ = v1306
	var v1309 int32
	_ = v1309
	var v1312 int32
	_ = v1312
	var v1315 int32
	_ = v1315
	var v1319 int32
	_ = v1319
	var v1320 int32
	_ = v1320
	var v1325 int32
	_ = v1325
	var v1331 int32
	_ = v1331
	var v1334 int32
	_ = v1334
	var v1343 int32
	_ = v1343
	var v1348 int32
	_ = v1348
	var v1349 int32
	_ = v1349
	var v1355 int32
	_ = v1355
	var v1358 int32
	_ = v1358
	var v1369 int32
	_ = v1369
	var v1374 int32
	_ = v1374
	var v1382 int32
	_ = v1382
	var v1385 int32
	_ = v1385
	var v1394 int32
	_ = v1394
	var v1399 int32
	_ = v1399
	var v1403 int32
	_ = v1403
	var v1406 int32
	_ = v1406
	var v1415 int32
	_ = v1415
	var v1420 int32
	_ = v1420
	var v1429 int32
	_ = v1429
	var v1434 int32
	_ = v1434
	var v1443 int32
	_ = v1443
	var v1448 int32
	_ = v1448
	var v1452 int32
	_ = v1452
	v26 = int64(0)
	v27 = m.G0
	v29 = v27 - int32(160)
	m.G0 = v29
	*(*int64)(unsafe.Add(mBase, uint32(l3))) = v26
	*(*int64)(unsafe.Add(mBase, uint32(l3)+32)) = v26
	*(*int64)(unsafe.Add(mBase, uint32(l3)+24)) = v26
	*(*int64)(unsafe.Add(mBase, uint32(l3)+16)) = v26
	*(*int64)(unsafe.Add(mBase, uint32(l3)+8)) = v26
	v42 = l2 & int32(1)
	if v42 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v43 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+12)) = uint8(v43)
	goto L3
L2:
	;
	goto L3
L3:
	;
	v46 = l2 & int32(2)
	if v46 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v47 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+13)) = uint8(v47)
	goto L6
L5:
	;
	goto L6
L6:
	;
	v50 = l2 & int32(4)
	if v50 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v51 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+14)) = uint8(v51)
	goto L9
L8:
	;
	goto L9
L9:
	;
	v54 = l2 & int32(16)
	if v54 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v55 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+15)) = uint8(v55)
	goto L12
L11:
	;
	goto L12
L12:
	;
	v58 = l2 & int32(64)
	if v58 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v59 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+16)) = uint8(v59)
	goto L15
L14:
	;
	goto L15
L15:
	;
	v62 = l2 & int32(128)
	if v62 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v63 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+17)) = uint8(v63)
	goto L18
L17:
	;
	goto L18
L18:
	;
	v66 = l2 & int32(256)
	if v66 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v67 = int32(112)
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+18)) = uint8(v67)
	goto L21
L20:
	;
	goto L21
L21:
	;
	v70 = l2 & int32(512)
	if v70 != 0 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v71 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+19)) = uint8(v71)
	goto L24
L23:
	;
	goto L24
L24:
	;
	v74 = l2 & int32(1024)
	if v74 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v75 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+20)) = uint8(v75)
	goto L27
L26:
	;
	goto L27
L27:
	;
	v78 = l2 & int32(2048)
	if v78 != 0 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v79 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+21)) = uint8(v79)
	goto L30
L29:
	;
	goto L30
L30:
	;
	v82 = l2 & int32(4096)
	if v82 != 0 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v83 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+22)) = uint8(v83)
	goto L33
L32:
	;
	goto L33
L33:
	;
	v86 = l2 & int32(8192)
	if v86 != 0 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v87 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+23)) = uint8(v87)
	goto L36
L35:
	;
	goto L36
L36:
	;
	if base.Ui32(int32(32768)) <= base.Ui32(l2) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v92 = F_pstrdup(m, int32(18822))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L40
	} else {
		goto L41
	}
L38:
	;
	goto L39
L39:
	;
	if l1 == int32(0) {
		goto L43
	} else {
		goto L44
	}
L40:
	;
	return
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+24)) = v92
	goto L39
L42:
	;
	F_errorConflictingDefElem(m, v138, l0)
	mBase = m.M
	v1452 = m.ExcPending
	if v1452 != 0 {
		goto L40
	} else {
		goto L460
	}
L43:
	;
	if v42 == int32(0) {
		goto L413
	} else {
		goto L414
	}
L44:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v97 <= int32(0) {
		goto L43
	} else {
		goto L45
	}
L45:
	;
	v117 = int32(0)
	goto L46
L46:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v134+v117<<(uint(int32(2))%32))))
	if v42 == int32(0) {
		goto L51
	} else {
		goto L52
	}
L47:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1228 = m.ExcPending
	if v1228 != 0 {
		goto L40
	} else {
		goto L405
	}
L48:
	;
	goto L47
L49:
	;
	v1222 = v117 + int32(1)
	v1223 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v1222 < v1223 {
		v117 = v1222
		goto L46
	} else {
		goto L404
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v168 | int32(1)
	v1213 = F_defGetBoolean(m, v138)
	mBase = m.M
	v1214 = m.ExcPending
	if v1214 != 0 {
		goto L40
	} else {
		goto L403
	}
L51:
	;
	if v46 == int32(0) {
		goto L63
	} else {
		goto L64
	}
L52:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v138)+8))
	v142 = int32(117493)
	v145 = int32(*(*uint8)(unsafe.Add(mBase, _consts[273])))
	v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v141))))
	if v146 == int32(0) {
		v165 = v145
		v166 = v146
		goto L54
	} else {
		goto L55
	}
L53:
	;
	if v166-v165 != 0 {
		goto L51
	} else {
		goto L61
	}
L54:
	;
	goto L53
L55:
	;
	if v145 != v146 {
		v165 = v145
		v166 = v146
		goto L54
	} else {
		goto L56
	}
L56:
	;
	v150 = v141
	v151 = v142
	goto L57
L57:
	;
	v154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+1)))
	v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v150)+1)))
	if v155 == int32(0) {
		v165 = v154
		v166 = v155
		goto L54
	} else {
		goto L59
	}
L58:
	;
	v165 = v154
	v166 = v155
	goto L54
L59:
	;
	v158 = int32(1)
	if v154 == v155 {
		v150 = v150 + v158
		v151 = v151 + v158
		goto L57
	} else {
		goto L60
	}
L60:
	;
	goto L58
L61:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v168&int32(1) == int32(0) {
		goto L50
	} else {
		goto L62
	}
L62:
	;
	goto L42
L63:
	;
	if v50 == int32(0) {
		goto L76
	} else {
		goto L77
	}
L64:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v138)+8))
	v176 = int32(477547)
	v179 = int32(*(*uint8)(unsafe.Add(mBase, _consts[274])))
	v180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175))))
	if v180 == int32(0) {
		v199 = v179
		v200 = v180
		goto L66
	} else {
		goto L67
	}
L65:
	;
	if v200-v199 != 0 {
		goto L63
	} else {
		goto L73
	}
L66:
	;
	goto L65
L67:
	;
	if v179 != v180 {
		v199 = v179
		v200 = v180
		goto L66
	} else {
		goto L68
	}
L68:
	;
	v184 = v175
	v185 = v176
	goto L69
L69:
	;
	v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v185)+1)))
	v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v184)+1)))
	if v189 == int32(0) {
		v199 = v188
		v200 = v189
		goto L66
	} else {
		goto L71
	}
L70:
	;
	v199 = v188
	v200 = v189
	goto L66
L71:
	;
	v192 = int32(1)
	if v188 == v189 {
		v184 = v184 + v192
		v185 = v185 + v192
		goto L69
	} else {
		goto L72
	}
L72:
	;
	goto L70
L73:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v202&int32(2) != 0 {
		goto L42
	} else {
		goto L74
	}
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v202 | int32(2)
	v208 = F_defGetBoolean(m, v138)
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L40
	} else {
		goto L75
	}
L75:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+13)) = uint8(v208)
	goto L49
L76:
	;
	if l2&int32(8) == int32(0) {
		goto L89
	} else {
		goto L90
	}
L77:
	;
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v138)+8))
	v214 = int32(91378)
	v217 = int32(*(*uint8)(unsafe.Add(mBase, _consts[275])))
	v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v213))))
	if v218 == int32(0) {
		v237 = v217
		v238 = v218
		goto L79
	} else {
		goto L80
	}
L78:
	;
	if v238-v237 != 0 {
		goto L76
	} else {
		goto L86
	}
L79:
	;
	goto L78
L80:
	;
	if v217 != v218 {
		v237 = v217
		v238 = v218
		goto L79
	} else {
		goto L81
	}
L81:
	;
	v222 = v213
	v223 = v214
	goto L82
L82:
	;
	v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v223)+1)))
	v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v222)+1)))
	if v227 == int32(0) {
		v237 = v226
		v238 = v227
		goto L79
	} else {
		goto L84
	}
L83:
	;
	v237 = v226
	v238 = v227
	goto L79
L84:
	;
	v230 = int32(1)
	if v226 == v227 {
		v222 = v222 + v230
		v223 = v223 + v230
		goto L82
	} else {
		goto L85
	}
L85:
	;
	goto L83
L86:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v240&int32(4) != 0 {
		goto L42
	} else {
		goto L87
	}
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v240 | int32(4)
	v246 = F_defGetBoolean(m, v138)
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L40
	} else {
		goto L88
	}
L88:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+14)) = uint8(v246)
	goto L49
L89:
	;
	if v54 == int32(0) {
		goto L114
	} else {
		goto L115
	}
L90:
	;
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v138)+8))
	v252 = int32(397025)
	v255 = int32(*(*uint8)(unsafe.Add(mBase, _consts[276])))
	v256 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v251))))
	if v256 == int32(0) {
		v275 = v255
		v276 = v256
		goto L92
	} else {
		goto L93
	}
L91:
	;
	if v276-v275 != 0 {
		goto L89
	} else {
		goto L99
	}
L92:
	;
	goto L91
L93:
	;
	if v255 != v256 {
		v275 = v255
		v276 = v256
		goto L92
	} else {
		goto L94
	}
L94:
	;
	v260 = v251
	v261 = v252
	goto L95
L95:
	;
	v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v261)+1)))
	v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v260)+1)))
	if v265 == int32(0) {
		v275 = v264
		v276 = v265
		goto L92
	} else {
		goto L97
	}
L96:
	;
	v275 = v264
	v276 = v265
	goto L92
L97:
	;
	v268 = int32(1)
	if v264 == v265 {
		v260 = v260 + v268
		v261 = v261 + v268
		goto L95
	} else {
		goto L98
	}
L98:
	;
	goto L96
L99:
	;
	v278 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v278&int32(8) != 0 {
		goto L42
	} else {
		goto L100
	}
L100:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v278 | int32(8)
	v284 = F_defGetString(m, v138)
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L40
	} else {
		goto L101
	}
L101:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+4)) = v284
	v287 = int32(391048)
	v290 = int32(*(*uint8)(unsafe.Add(mBase, _consts[277])))
	v291 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v284))))
	if v291 == int32(0) {
		v310 = v290
		v311 = v291
		goto L103
	} else {
		goto L104
	}
L102:
	;
	if v311-v310 == int32(0) {
		goto L110
	} else {
		goto L111
	}
L103:
	;
	goto L102
L104:
	;
	if v290 != v291 {
		v310 = v290
		v311 = v291
		goto L103
	} else {
		goto L105
	}
L105:
	;
	v295 = v284
	v296 = v287
	goto L106
L106:
	;
	v299 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v296)+1)))
	v300 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v295)+1)))
	if v300 == int32(0) {
		v310 = v299
		v311 = v300
		goto L103
	} else {
		goto L108
	}
L107:
	;
	v310 = v299
	v311 = v300
	goto L103
L108:
	;
	v303 = int32(1)
	if v299 == v300 {
		v295 = v295 + v303
		v296 = v296 + v303
		goto L106
	} else {
		goto L109
	}
L109:
	;
	goto L107
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+4)) = int32(0)
	goto L49
L111:
	;
	goto L112
L112:
	;
	v318 = F_ReplicationSlotValidateName(m, v284, int32(21))
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L40
	} else {
		goto L113
	}
L113:
	;
	goto L49
L114:
	;
	if l2&int32(32) == int32(0) {
		goto L127
	} else {
		goto L128
	}
L115:
	;
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v138)+8))
	v323 = int32(529779)
	v326 = int32(*(*uint8)(unsafe.Add(mBase, _consts[278])))
	v327 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v322))))
	if v327 == int32(0) {
		v346 = v326
		v347 = v327
		goto L117
	} else {
		goto L118
	}
L116:
	;
	if v347-v346 != 0 {
		goto L114
	} else {
		goto L124
	}
L117:
	;
	goto L116
L118:
	;
	if v326 != v327 {
		v346 = v326
		v347 = v327
		goto L117
	} else {
		goto L119
	}
L119:
	;
	v331 = v322
	v332 = v323
	goto L120
L120:
	;
	v335 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v332)+1)))
	v336 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v331)+1)))
	if v336 == int32(0) {
		v346 = v335
		v347 = v336
		goto L117
	} else {
		goto L122
	}
L121:
	;
	v346 = v335
	v347 = v336
	goto L117
L122:
	;
	v339 = int32(1)
	if v335 == v336 {
		v331 = v331 + v339
		v332 = v332 + v339
		goto L120
	} else {
		goto L123
	}
L123:
	;
	goto L121
L124:
	;
	v349 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v349&int32(16) != 0 {
		goto L42
	} else {
		goto L125
	}
L125:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v349 | int32(16)
	v355 = F_defGetBoolean(m, v138)
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L40
	} else {
		goto L126
	}
L126:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+15)) = uint8(v355)
	goto L49
L127:
	;
	if v58 == int32(0) {
		goto L141
	} else {
		goto L142
	}
L128:
	;
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v138)+8))
	v361 = int32(107394)
	v364 = int32(*(*uint8)(unsafe.Add(mBase, _consts[279])))
	v365 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v360))))
	if v365 == int32(0) {
		v384 = v364
		v385 = v365
		goto L130
	} else {
		goto L131
	}
L129:
	;
	if v385-v384 != 0 {
		goto L127
	} else {
		goto L137
	}
L130:
	;
	goto L129
L131:
	;
	if v364 != v365 {
		v384 = v364
		v385 = v365
		goto L130
	} else {
		goto L132
	}
L132:
	;
	v369 = v360
	v370 = v361
	goto L133
L133:
	;
	v373 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v370)+1)))
	v374 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v369)+1)))
	if v374 == int32(0) {
		v384 = v373
		v385 = v374
		goto L130
	} else {
		goto L135
	}
L134:
	;
	v384 = v373
	v385 = v374
	goto L130
L135:
	;
	v377 = int32(1)
	if v373 == v374 {
		v369 = v369 + v377
		v370 = v370 + v377
		goto L133
	} else {
		goto L136
	}
L136:
	;
	goto L134
L137:
	;
	v387 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v387&int32(32) != 0 {
		goto L42
	} else {
		goto L138
	}
L138:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v387 | int32(32)
	v393 = F_defGetString(m, v138)
	mBase = m.M
	v394 = m.ExcPending
	if v394 != 0 {
		goto L40
	} else {
		goto L139
	}
L139:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+8)) = v393
	v399 = int32(0)
	F_set_config_option(m, int32(107394), v393, int32(4), int32(12), v399, v399)
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		goto L40
	} else {
		goto L140
	}
L140:
	;
	goto L49
L141:
	;
	if v62 == int32(0) {
		goto L154
	} else {
		goto L155
	}
L142:
	;
	v405 = *(*int32)(unsafe.Add(mBase, uint32(v138)+8))
	v406 = int32(338698)
	v409 = int32(*(*uint8)(unsafe.Add(mBase, _consts[280])))
	v410 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v405))))
	if v410 == int32(0) {
		v429 = v409
		v430 = v410
		goto L144
	} else {
		goto L145
	}
L143:
	;
	if v430-v429 != 0 {
		goto L141
	} else {
		goto L151
	}
L144:
	;
	goto L143
L145:
	;
	if v409 != v410 {
		v429 = v409
		v430 = v410
		goto L144
	} else {
		goto L146
	}
L146:
	;
	v414 = v405
	v415 = v406
	goto L147
L147:
	;
	v418 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v415)+1)))
	v419 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v414)+1)))
	if v419 == int32(0) {
		v429 = v418
		v430 = v419
		goto L144
	} else {
		goto L149
	}
L148:
	;
	v429 = v418
	v430 = v419
	goto L144
L149:
	;
	v422 = int32(1)
	if v418 == v419 {
		v414 = v414 + v422
		v415 = v415 + v422
		goto L147
	} else {
		goto L150
	}
L150:
	;
	goto L148
L151:
	;
	v432 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v432&int32(64) != 0 {
		goto L42
	} else {
		goto L152
	}
L152:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v432 | int32(64)
	v438 = F_defGetBoolean(m, v138)
	mBase = m.M
	v439 = m.ExcPending
	if v439 != 0 {
		goto L40
	} else {
		goto L153
	}
L153:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+16)) = uint8(v438)
	goto L49
L154:
	;
	if v66 == int32(0) {
		goto L167
	} else {
		goto L168
	}
L155:
	;
	v443 = *(*int32)(unsafe.Add(mBase, uint32(v138)+8))
	v444 = int32(18016)
	v447 = int32(*(*uint8)(unsafe.Add(mBase, _consts[281])))
	v448 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v443))))
	if v448 == int32(0) {
		v467 = v447
		v468 = v448
		goto L157
	} else {
		goto L158
	}
L156:
	;
	if v468-v467 != 0 {
		goto L154
	} else {
		goto L164
	}
L157:
	;
	goto L156
L158:
	;
	if v447 != v448 {
		v467 = v447
		v468 = v448
		goto L157
	} else {
		goto L159
	}
L159:
	;
	v452 = v443
	v453 = v444
	goto L160
L160:
	;
	v456 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v453)+1)))
	v457 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v452)+1)))
	if v457 == int32(0) {
		v467 = v456
		v468 = v457
		goto L157
	} else {
		goto L162
	}
L161:
	;
	v467 = v456
	v468 = v457
	goto L157
L162:
	;
	v460 = int32(1)
	if v456 == v457 {
		v452 = v452 + v460
		v453 = v453 + v460
		goto L160
	} else {
		goto L163
	}
L163:
	;
	goto L161
L164:
	;
	v470 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v470&int32(128) != 0 {
		goto L42
	} else {
		goto L165
	}
L165:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v470 | int32(128)
	v476 = F_defGetBoolean(m, v138)
	mBase = m.M
	v477 = m.ExcPending
	if v477 != 0 {
		goto L40
	} else {
		goto L166
	}
L166:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+17)) = uint8(v476)
	goto L49
L167:
	;
	if v70 == int32(0) {
		goto L261
	} else {
		goto L262
	}
L168:
	;
	v481 = *(*int32)(unsafe.Add(mBase, uint32(v138)+8))
	v482 = int32(352562)
	v485 = int32(*(*uint8)(unsafe.Add(mBase, _consts[282])))
	v486 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v481))))
	if v486 == int32(0) {
		v505 = v485
		v506 = v486
		goto L170
	} else {
		goto L171
	}
L169:
	;
	if v506-v505 != 0 {
		goto L167
	} else {
		goto L177
	}
L170:
	;
	goto L169
L171:
	;
	if v485 != v486 {
		v505 = v485
		v506 = v486
		goto L170
	} else {
		goto L172
	}
L172:
	;
	v490 = v481
	v491 = v482
	goto L173
L173:
	;
	v494 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v491)+1)))
	v495 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v490)+1)))
	if v495 == int32(0) {
		v505 = v494
		v506 = v495
		goto L170
	} else {
		goto L175
	}
L174:
	;
	v505 = v494
	v506 = v495
	goto L170
L175:
	;
	v498 = int32(1)
	if v494 == v495 {
		v490 = v490 + v498
		v491 = v491 + v498
		goto L173
	} else {
		goto L176
	}
L176:
	;
	goto L174
L177:
	;
	v508 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v508&int32(256) != 0 {
		goto L42
	} else {
		goto L178
	}
L178:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v508 | int32(256)
	v514 = m.G0
	v516 = v514 - int32(16)
	m.G0 = v516
	v518 = *(*int32)(unsafe.Add(mBase, uint32(v138)+12))
	if v518 == int32(0) {
		goto L180
	} else {
		goto L181
	}
L179:
	;
	m.G0 = v516 + int32(16)
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+18)) = uint8(v772)
	goto L49
L180:
	;
	v772 = int32(116)
	goto L179
L181:
	;
	v521 = *(*int32)(unsafe.Add(mBase, uint32(v518)))
	if v521 == int32(465) {
		goto L183
	} else {
		goto L184
	}
L182:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v754 = m.ExcPending
	if v754 != 0 {
		goto L40
	} else {
		goto L257
	}
L183:
	;
	v525 = *(*int32)(unsafe.Add(mBase, uint32(v518)+4))
	switch v525 {
	case 0:
		v772 = int32(102)
		goto L179
	case 1:
		goto L180
	default:
		goto L182
	}
L184:
	;
	goto L185
L185:
	;
	v526 = int32(102)
	v527 = F_defGetString(m, v138)
	mBase = m.M
	v528 = m.ExcPending
	if v528 != 0 {
		goto L40
	} else {
		goto L186
	}
L186:
	;
	v532 = v527
	v533 = int32(379468)
	goto L188
L187:
	;
	if v570 == int32(0) {
		v772 = v526
		goto L179
	} else {
		goto L200
	}
L188:
	;
	v536 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v532))))
	v537 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v533))))
	if v536 == v537 {
		v559 = v536
		goto L190
	} else {
		goto L191
	}
L189:
	;
	v570 = int32(0)
	goto L187
L190:
	;
	v561 = int32(1)
	if v559 != 0 {
		v532 = v532 + v561
		v533 = v533 + v561
		goto L188
	} else {
		goto L199
	}
L191:
	;
	if base.Ui32((v536-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L192
	} else {
		goto L193
	}
L192:
	;
	v547 = v536 | int32(32)
	goto L194
L193:
	;
	v547 = v536
	goto L194
L194:
	;
	if base.Ui32((v537-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L195
	} else {
		goto L196
	}
L195:
	;
	v556 = v537 | int32(32)
	goto L197
L196:
	;
	v556 = v537
	goto L197
L197:
	;
	if v547 == v556 {
		v559 = v547
		goto L190
	} else {
		goto L198
	}
L198:
	;
	v570 = v547 - v556
	goto L187
L199:
	;
	goto L189
L200:
	;
	v576 = v527
	v577 = int32(356102)
	goto L202
L201:
	;
	if v614 == int32(0) {
		v772 = v526
		goto L179
	} else {
		goto L214
	}
L202:
	;
	v580 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v576))))
	v581 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v577))))
	if v580 == v581 {
		v603 = v580
		goto L204
	} else {
		goto L205
	}
L203:
	;
	v614 = int32(0)
	goto L201
L204:
	;
	v605 = int32(1)
	if v603 != 0 {
		v576 = v576 + v605
		v577 = v577 + v605
		goto L202
	} else {
		goto L213
	}
L205:
	;
	if base.Ui32((v580-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L206
	} else {
		goto L207
	}
L206:
	;
	v591 = v580 | int32(32)
	goto L208
L207:
	;
	v591 = v580
	goto L208
L208:
	;
	if base.Ui32((v581-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L209
	} else {
		goto L210
	}
L209:
	;
	v600 = v581 | int32(32)
	goto L211
L210:
	;
	v600 = v581
	goto L211
L211:
	;
	if v591 == v600 {
		v603 = v591
		goto L204
	} else {
		goto L212
	}
L212:
	;
	v614 = v591 - v600
	goto L201
L213:
	;
	goto L203
L214:
	;
	v620 = v527
	v621 = int32(361949)
	goto L216
L215:
	;
	if v658 == int32(0) {
		goto L180
	} else {
		goto L228
	}
L216:
	;
	v624 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v620))))
	v625 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v621))))
	if v624 == v625 {
		v647 = v624
		goto L218
	} else {
		goto L219
	}
L217:
	;
	v658 = int32(0)
	goto L215
L218:
	;
	v649 = int32(1)
	if v647 != 0 {
		v620 = v620 + v649
		v621 = v621 + v649
		goto L216
	} else {
		goto L227
	}
L219:
	;
	if base.Ui32((v624-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L220
	} else {
		goto L221
	}
L220:
	;
	v635 = v624 | int32(32)
	goto L222
L221:
	;
	v635 = v624
	goto L222
L222:
	;
	if base.Ui32((v625-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L223
	} else {
		goto L224
	}
L223:
	;
	v644 = v625 | int32(32)
	goto L225
L224:
	;
	v644 = v625
	goto L225
L225:
	;
	if v635 == v644 {
		v647 = v635
		goto L218
	} else {
		goto L226
	}
L226:
	;
	v658 = v635 - v644
	goto L215
L227:
	;
	goto L217
L228:
	;
	v665 = v527
	v666 = int32(286459)
	goto L230
L229:
	;
	if v703 == int32(0) {
		v772 = int32(116)
		goto L179
	} else {
		goto L242
	}
L230:
	;
	v669 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v665))))
	v670 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v666))))
	if v669 == v670 {
		v692 = v669
		goto L232
	} else {
		goto L233
	}
L231:
	;
	v703 = int32(0)
	goto L229
L232:
	;
	v694 = int32(1)
	if v692 != 0 {
		v665 = v665 + v694
		v666 = v666 + v694
		goto L230
	} else {
		goto L241
	}
L233:
	;
	if base.Ui32((v669-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L234
	} else {
		goto L235
	}
L234:
	;
	v680 = v669 | int32(32)
	goto L236
L235:
	;
	v680 = v669
	goto L236
L236:
	;
	if base.Ui32((v670-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L237
	} else {
		goto L238
	}
L237:
	;
	v689 = v670 | int32(32)
	goto L239
L238:
	;
	v689 = v670
	goto L239
L239:
	;
	if v680 == v689 {
		v692 = v680
		goto L232
	} else {
		goto L240
	}
L240:
	;
	v703 = v680 - v689
	goto L229
L241:
	;
	goto L231
L242:
	;
	v709 = v527
	v710 = int32(322898)
	goto L244
L243:
	;
	if v747 != 0 {
		goto L182
	} else {
		goto L256
	}
L244:
	;
	v713 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v709))))
	v714 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v710))))
	if v713 == v714 {
		v736 = v713
		goto L246
	} else {
		goto L247
	}
L245:
	;
	v747 = int32(0)
	goto L243
L246:
	;
	v738 = int32(1)
	if v736 != 0 {
		v709 = v709 + v738
		v710 = v710 + v738
		goto L244
	} else {
		goto L255
	}
L247:
	;
	if base.Ui32((v713-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L248
	} else {
		goto L249
	}
L248:
	;
	v724 = v713 | int32(32)
	goto L250
L249:
	;
	v724 = v713
	goto L250
L250:
	;
	if base.Ui32((v714-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L251
	} else {
		goto L252
	}
L251:
	;
	v733 = v714 | int32(32)
	goto L253
L252:
	;
	v733 = v714
	goto L253
L253:
	;
	if v724 == v733 {
		v736 = v724
		goto L246
	} else {
		goto L254
	}
L254:
	;
	v747 = v724 - v733
	goto L243
L255:
	;
	goto L245
L256:
	;
	v772 = int32(112)
	goto L179
L257:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v757 = m.ExcPending
	if v757 != 0 {
		goto L40
	} else {
		goto L258
	}
L258:
	;
	v758 = *(*int32)(unsafe.Add(mBase, uint32(v138)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v516))) = v758
	F_errmsg(m, int32(762776), v516)
	mBase = m.M
	v762 = m.ExcPending
	if v762 != 0 {
		goto L40
	} else {
		goto L259
	}
L259:
	;
	F_errfinish(m, int32(518572), int32(2509), int32(434230))
	mBase = m.M
	v767 = m.ExcPending
	if v767 != 0 {
		goto L40
	} else {
		goto L260
	}
L260:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L261:
	;
	if v74 == int32(0) {
		goto L274
	} else {
		goto L275
	}
L262:
	;
	v779 = *(*int32)(unsafe.Add(mBase, uint32(v138)+8))
	v780 = int32(379552)
	v783 = int32(*(*uint8)(unsafe.Add(mBase, _consts[283])))
	v784 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v779))))
	if v784 == int32(0) {
		v803 = v783
		v804 = v784
		goto L264
	} else {
		goto L265
	}
L263:
	;
	if v804-v803 != 0 {
		goto L261
	} else {
		goto L271
	}
L264:
	;
	goto L263
L265:
	;
	if v783 != v784 {
		v803 = v783
		v804 = v784
		goto L264
	} else {
		goto L266
	}
L266:
	;
	v788 = v779
	v789 = v780
	goto L267
L267:
	;
	v792 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v789)+1)))
	v793 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v788)+1)))
	if v793 == int32(0) {
		v803 = v792
		v804 = v793
		goto L264
	} else {
		goto L269
	}
L268:
	;
	v803 = v792
	v804 = v793
	goto L264
L269:
	;
	v796 = int32(1)
	if v792 == v793 {
		v788 = v788 + v796
		v789 = v789 + v796
		goto L267
	} else {
		goto L270
	}
L270:
	;
	goto L268
L271:
	;
	v806 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v806&int32(512) != 0 {
		goto L42
	} else {
		goto L272
	}
L272:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v806 | int32(512)
	v812 = F_defGetBoolean(m, v138)
	mBase = m.M
	v813 = m.ExcPending
	if v813 != 0 {
		goto L40
	} else {
		goto L273
	}
L273:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+19)) = uint8(v812)
	goto L49
L274:
	;
	if v78 == int32(0) {
		goto L287
	} else {
		goto L288
	}
L275:
	;
	v817 = *(*int32)(unsafe.Add(mBase, uint32(v138)+8))
	v818 = int32(223009)
	v821 = int32(*(*uint8)(unsafe.Add(mBase, _consts[284])))
	v822 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v817))))
	if v822 == int32(0) {
		v841 = v821
		v842 = v822
		goto L277
	} else {
		goto L278
	}
L276:
	;
	if v842-v841 != 0 {
		goto L274
	} else {
		goto L284
	}
L277:
	;
	goto L276
L278:
	;
	if v821 != v822 {
		v841 = v821
		v842 = v822
		goto L277
	} else {
		goto L279
	}
L279:
	;
	v826 = v817
	v827 = v818
	goto L280
L280:
	;
	v830 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v827)+1)))
	v831 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v826)+1)))
	if v831 == int32(0) {
		v841 = v830
		v842 = v831
		goto L277
	} else {
		goto L282
	}
L281:
	;
	v841 = v830
	v842 = v831
	goto L277
L282:
	;
	v834 = int32(1)
	if v830 == v831 {
		v826 = v826 + v834
		v827 = v827 + v834
		goto L280
	} else {
		goto L283
	}
L283:
	;
	goto L281
L284:
	;
	v844 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v844&int32(1024) != 0 {
		goto L42
	} else {
		goto L285
	}
L285:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v844 | int32(1024)
	v850 = F_defGetBoolean(m, v138)
	mBase = m.M
	v851 = m.ExcPending
	if v851 != 0 {
		goto L40
	} else {
		goto L286
	}
L286:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+20)) = uint8(v850)
	goto L49
L287:
	;
	if v82 == int32(0) {
		goto L300
	} else {
		goto L301
	}
L288:
	;
	v855 = *(*int32)(unsafe.Add(mBase, uint32(v138)+8))
	v856 = int32(471901)
	v859 = int32(*(*uint8)(unsafe.Add(mBase, _consts[285])))
	v860 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v855))))
	if v860 == int32(0) {
		v879 = v859
		v880 = v860
		goto L290
	} else {
		goto L291
	}
L289:
	;
	if v880-v879 != 0 {
		goto L287
	} else {
		goto L297
	}
L290:
	;
	goto L289
L291:
	;
	if v859 != v860 {
		v879 = v859
		v880 = v860
		goto L290
	} else {
		goto L292
	}
L292:
	;
	v864 = v855
	v865 = v856
	goto L293
L293:
	;
	v868 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v865)+1)))
	v869 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v864)+1)))
	if v869 == int32(0) {
		v879 = v868
		v880 = v869
		goto L290
	} else {
		goto L295
	}
L294:
	;
	v879 = v868
	v880 = v869
	goto L290
L295:
	;
	v872 = int32(1)
	if v868 == v869 {
		v864 = v864 + v872
		v865 = v865 + v872
		goto L293
	} else {
		goto L296
	}
L296:
	;
	goto L294
L297:
	;
	v882 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v882&int32(2048) != 0 {
		goto L42
	} else {
		goto L298
	}
L298:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v882 | int32(2048)
	v888 = F_defGetBoolean(m, v138)
	mBase = m.M
	v889 = m.ExcPending
	if v889 != 0 {
		goto L40
	} else {
		goto L299
	}
L299:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+21)) = uint8(v888)
	goto L49
L300:
	;
	if v86 == int32(0) {
		goto L313
	} else {
		goto L314
	}
L301:
	;
	v893 = *(*int32)(unsafe.Add(mBase, uint32(v138)+8))
	v894 = int32(229517)
	v897 = int32(*(*uint8)(unsafe.Add(mBase, _consts[286])))
	v898 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v893))))
	if v898 == int32(0) {
		v917 = v897
		v918 = v898
		goto L303
	} else {
		goto L304
	}
L302:
	;
	if v918-v917 != 0 {
		goto L300
	} else {
		goto L310
	}
L303:
	;
	goto L302
L304:
	;
	if v897 != v898 {
		v917 = v897
		v918 = v898
		goto L303
	} else {
		goto L305
	}
L305:
	;
	v902 = v893
	v903 = v894
	goto L306
L306:
	;
	v906 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v903)+1)))
	v907 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v902)+1)))
	if v907 == int32(0) {
		v917 = v906
		v918 = v907
		goto L303
	} else {
		goto L308
	}
L307:
	;
	v917 = v906
	v918 = v907
	goto L303
L308:
	;
	v910 = int32(1)
	if v906 == v907 {
		v902 = v902 + v910
		v903 = v903 + v910
		goto L306
	} else {
		goto L309
	}
L309:
	;
	goto L307
L310:
	;
	v920 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v920&int32(4096) != 0 {
		goto L42
	} else {
		goto L311
	}
L311:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v920 | int32(4096)
	v926 = F_defGetBoolean(m, v138)
	mBase = m.M
	v927 = m.ExcPending
	if v927 != 0 {
		goto L40
	} else {
		goto L312
	}
L312:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+22)) = uint8(v926)
	goto L49
L313:
	;
	if base.Ui32(l2) < base.Ui32(int32(32768)) {
		goto L326
	} else {
		goto L327
	}
L314:
	;
	v931 = *(*int32)(unsafe.Add(mBase, uint32(v138)+8))
	v932 = int32(226082)
	v935 = int32(*(*uint8)(unsafe.Add(mBase, _consts[287])))
	v936 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v931))))
	if v936 == int32(0) {
		v955 = v935
		v956 = v936
		goto L316
	} else {
		goto L317
	}
L315:
	;
	if v956-v955 != 0 {
		goto L313
	} else {
		goto L323
	}
L316:
	;
	goto L315
L317:
	;
	if v935 != v936 {
		v955 = v935
		v956 = v936
		goto L316
	} else {
		goto L318
	}
L318:
	;
	v940 = v931
	v941 = v932
	goto L319
L319:
	;
	v944 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v941)+1)))
	v945 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v940)+1)))
	if v945 == int32(0) {
		v955 = v944
		v956 = v945
		goto L316
	} else {
		goto L321
	}
L320:
	;
	v955 = v944
	v956 = v945
	goto L316
L321:
	;
	v948 = int32(1)
	if v944 == v945 {
		v940 = v940 + v948
		v941 = v941 + v948
		goto L319
	} else {
		goto L322
	}
L322:
	;
	goto L320
L323:
	;
	v958 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v958&int32(8192) != 0 {
		goto L42
	} else {
		goto L324
	}
L324:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v958 | int32(8192)
	v964 = F_defGetBoolean(m, v138)
	mBase = m.M
	v965 = m.ExcPending
	if v965 != 0 {
		goto L40
	} else {
		goto L325
	}
L325:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+23)) = uint8(v964)
	goto L49
L326:
	;
	if l2&int32(16384) == int32(0) {
		goto L372
	} else {
		goto L373
	}
L327:
	;
	v967 = *(*int32)(unsafe.Add(mBase, uint32(v138)+8))
	v968 = int32(290156)
	v971 = int32(*(*uint8)(unsafe.Add(mBase, _consts[288])))
	v972 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v967))))
	if v972 == int32(0) {
		v991 = v971
		v992 = v972
		goto L329
	} else {
		goto L330
	}
L328:
	;
	if v992-v991 != 0 {
		goto L326
	} else {
		goto L336
	}
L329:
	;
	goto L328
L330:
	;
	if v971 != v972 {
		v991 = v971
		v992 = v972
		goto L329
	} else {
		goto L331
	}
L331:
	;
	v976 = v967
	v977 = v968
	goto L332
L332:
	;
	v980 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v977)+1)))
	v981 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v976)+1)))
	if v981 == int32(0) {
		v991 = v980
		v992 = v981
		goto L329
	} else {
		goto L334
	}
L333:
	;
	v991 = v980
	v992 = v981
	goto L329
L334:
	;
	v984 = int32(1)
	if v980 == v981 {
		v976 = v976 + v984
		v977 = v977 + v984
		goto L332
	} else {
		goto L335
	}
L335:
	;
	goto L333
L336:
	;
	v994 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v994&int32(32768) != 0 {
		goto L42
	} else {
		goto L337
	}
L337:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v994 | int32(32768)
	v1000 = *(*int32)(unsafe.Add(mBase, uint32(l3)+24))
	F_pfree(m, v1000)
	mBase = m.M
	v1002 = m.ExcPending
	if v1002 != 0 {
		goto L40
	} else {
		goto L338
	}
L338:
	;
	v1003 = F_defGetString(m, v138)
	mBase = m.M
	v1004 = m.ExcPending
	if v1004 != 0 {
		goto L40
	} else {
		goto L339
	}
L339:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+24)) = v1003
	v1009 = v1003
	v1010 = int32(391048)
	goto L341
L340:
	;
	if v1047 == int32(0) {
		goto L49
	} else {
		goto L353
	}
L341:
	;
	v1013 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1009))))
	v1014 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1010))))
	if v1013 == v1014 {
		v1036 = v1013
		goto L343
	} else {
		goto L344
	}
L342:
	;
	v1047 = int32(0)
	goto L340
L343:
	;
	v1038 = int32(1)
	if v1036 != 0 {
		v1009 = v1009 + v1038
		v1010 = v1010 + v1038
		goto L341
	} else {
		goto L352
	}
L344:
	;
	if base.Ui32((v1013-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L345
	} else {
		goto L346
	}
L345:
	;
	v1024 = v1013 | int32(32)
	goto L347
L346:
	;
	v1024 = v1013
	goto L347
L347:
	;
	if base.Ui32((v1014-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L348
	} else {
		goto L349
	}
L348:
	;
	v1033 = v1014 | int32(32)
	goto L350
L349:
	;
	v1033 = v1014
	goto L350
L350:
	;
	if v1024 == v1033 {
		v1036 = v1024
		goto L343
	} else {
		goto L351
	}
L351:
	;
	v1047 = v1024 - v1033
	goto L340
L352:
	;
	goto L342
L353:
	;
	v1050 = *(*int32)(unsafe.Add(mBase, uint32(l3)+24))
	v1054 = v1050
	v1055 = int32(18822)
	goto L355
L354:
	;
	if v1092 == int32(0) {
		goto L49
	} else {
		goto L367
	}
L355:
	;
	v1058 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1054))))
	v1059 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1055))))
	if v1058 == v1059 {
		v1081 = v1058
		goto L357
	} else {
		goto L358
	}
L356:
	;
	v1092 = int32(0)
	goto L354
L357:
	;
	v1083 = int32(1)
	if v1081 != 0 {
		v1054 = v1054 + v1083
		v1055 = v1055 + v1083
		goto L355
	} else {
		goto L366
	}
L358:
	;
	if base.Ui32((v1058-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L359
	} else {
		goto L360
	}
L359:
	;
	v1069 = v1058 | int32(32)
	goto L361
L360:
	;
	v1069 = v1058
	goto L361
L361:
	;
	if base.Ui32((v1059-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L362
	} else {
		goto L363
	}
L362:
	;
	v1078 = v1059 | int32(32)
	goto L364
L363:
	;
	v1078 = v1059
	goto L364
L364:
	;
	if v1069 == v1078 {
		v1081 = v1069
		goto L357
	} else {
		goto L365
	}
L365:
	;
	v1092 = v1069 - v1078
	goto L354
L366:
	;
	goto L356
L367:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1098 = m.ExcPending
	if v1098 != 0 {
		goto L40
	} else {
		goto L368
	}
L368:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v1101 = m.ExcPending
	if v1101 != 0 {
		goto L40
	} else {
		goto L369
	}
L369:
	;
	v1102 = *(*int32)(unsafe.Add(mBase, uint32(l3)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v29)+144)) = v1102
	F_errmsg(m, int32(761032), v29+int32(144))
	mBase = m.M
	v1108 = m.ExcPending
	if v1108 != 0 {
		goto L40
	} else {
		goto L370
	}
L370:
	;
	F_errfinish(m, int32(518572), int32(331), int32(145929))
	mBase = m.M
	v1113 = m.ExcPending
	if v1113 != 0 {
		goto L40
	} else {
		goto L371
	}
L371:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L372:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1194 = m.ExcPending
	if v1194 != 0 {
		goto L40
	} else {
		goto L399
	}
L373:
	;
	v1116 = *(*int32)(unsafe.Add(mBase, uint32(v138)+8))
	v1117 = int32(257509)
	v1120 = int32(*(*uint8)(unsafe.Add(mBase, _consts[289])))
	v1121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1116))))
	if v1121 == int32(0) {
		v1140 = v1120
		v1141 = v1121
		goto L375
	} else {
		goto L376
	}
L374:
	;
	if v1141-v1140 != 0 {
		goto L372
	} else {
		goto L382
	}
L375:
	;
	goto L374
L376:
	;
	if v1120 != v1121 {
		v1140 = v1120
		v1141 = v1121
		goto L375
	} else {
		goto L377
	}
L377:
	;
	v1125 = v1116
	v1126 = v1117
	goto L378
L378:
	;
	v1129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1126)+1)))
	v1130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1125)+1)))
	if v1130 == int32(0) {
		v1140 = v1129
		v1141 = v1130
		goto L375
	} else {
		goto L380
	}
L379:
	;
	v1140 = v1129
	v1141 = v1130
	goto L375
L380:
	;
	v1133 = int32(1)
	if v1129 == v1130 {
		v1125 = v1125 + v1133
		v1126 = v1126 + v1133
		goto L378
	} else {
		goto L381
	}
L381:
	;
	goto L379
L382:
	;
	v1143 = F_defGetString(m, v138)
	mBase = m.M
	v1144 = m.ExcPending
	if v1144 != 0 {
		goto L40
	} else {
		goto L383
	}
L383:
	;
	v1145 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v1145&int32(16384) != 0 {
		goto L42
	} else {
		goto L384
	}
L384:
	;
	v1148 = int32(391048)
	v1151 = int32(*(*uint8)(unsafe.Add(mBase, _consts[277])))
	v1152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1143))))
	if v1152 == int32(0) {
		v1171 = v1151
		v1172 = v1152
		goto L387
	} else {
		goto L388
	}
L385:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l3)+32)) = v1186
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v1185 | int32(16384)
	goto L49
L386:
	;
	if v1172-v1171 == int32(0) {
		goto L394
	} else {
		goto L395
	}
L387:
	;
	goto L386
L388:
	;
	if v1151 != v1152 {
		v1171 = v1151
		v1172 = v1152
		goto L387
	} else {
		goto L389
	}
L389:
	;
	v1156 = v1143
	v1157 = v1148
	goto L390
L390:
	;
	v1160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1157)+1)))
	v1161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1156)+1)))
	if v1161 == int32(0) {
		v1171 = v1160
		v1172 = v1161
		goto L387
	} else {
		goto L392
	}
L391:
	;
	v1171 = v1160
	v1172 = v1161
	goto L387
L392:
	;
	v1164 = int32(1)
	if v1160 == v1161 {
		v1156 = v1156 + v1164
		v1157 = v1157 + v1164
		goto L390
	} else {
		goto L393
	}
L393:
	;
	goto L391
L394:
	;
	v1185 = v1145
	v1186 = int64(0)
	goto L385
L395:
	;
	goto L396
L396:
	;
	v1179 = F_DirectFunctionCall1Coll(m, int32(572), int32(0), v1143)
	mBase = m.M
	v1180 = m.ExcPending
	if v1180 != 0 {
		goto L40
	} else {
		goto L397
	}
L397:
	;
	v1181 = *(*int64)(unsafe.Add(mBase, uint32(v1179)))
	if v1181 == int64(0) {
		goto L48
	} else {
		goto L398
	}
L398:
	;
	v1184 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v1185 = v1184
	v1186 = v1181
	goto L385
L399:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v1197 = m.ExcPending
	if v1197 != 0 {
		goto L40
	} else {
		goto L400
	}
L400:
	;
	v1198 = *(*int32)(unsafe.Add(mBase, uint32(v138)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v29)+112)) = v1198
	F_errmsg(m, int32(759838), v29+int32(112))
	mBase = m.M
	v1204 = m.ExcPending
	if v1204 != 0 {
		goto L40
	} else {
		goto L401
	}
L401:
	;
	F_errfinish(m, int32(518572), int32(363), int32(145929))
	mBase = m.M
	v1209 = m.ExcPending
	if v1209 != 0 {
		goto L40
	} else {
		goto L402
	}
L402:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L403:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+12)) = uint8(v1213)
	goto L49
L404:
	;
	goto L43
L405:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v1231 = m.ExcPending
	if v1231 != 0 {
		goto L40
	} else {
		goto L406
	}
L406:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+128)) = v1143
	F_errmsg(m, int32(215650), v29+int32(128))
	mBase = m.M
	v1237 = m.ExcPending
	if v1237 != 0 {
		goto L40
	} else {
		goto L407
	}
L407:
	;
	F_errfinish(m, int32(518572), int32(354), int32(145929))
	mBase = m.M
	v1242 = m.ExcPending
	if v1242 != 0 {
		goto L40
	} else {
		goto L408
	}
L408:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L409:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+52)) = int32(361903)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+48)) = int32(566020)
	F_errmsg(m, int32(146380), v29+int32(48))
	mBase = m.M
	v1443 = m.ExcPending
	if v1443 != 0 {
		goto L40
	} else {
		goto L458
	}
L410:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+20)) = int32(361922)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+16)) = int32(566020)
	F_errmsg(m, int32(146380), v29+int32(16))
	mBase = m.M
	v1429 = m.ExcPending
	if v1429 != 0 {
		goto L40
	} else {
		goto L456
	}
L411:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1403 = m.ExcPending
	if v1403 != 0 {
		goto L40
	} else {
		goto L452
	}
L412:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1382 = m.ExcPending
	if v1382 != 0 {
		goto L40
	} else {
		goto L448
	}
L413:
	;
	v1319 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v1319 != 0 {
		goto L431
	} else {
		goto L432
	}
L414:
	;
	v1271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+12)))
	if v1271&int32(1) != 0 {
		goto L413
	} else {
		goto L415
	}
L415:
	;
	v1274 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+13)))
	if v1274 != int32(1) {
		goto L416
	} else {
		goto L417
	}
L416:
	;
	v1303 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+14)))
	if v1303 == int32(1) {
		goto L423
	} else {
		goto L424
	}
L417:
	;
	v1277 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3))))
	if v1277&int32(2) == int32(0) {
		goto L416
	} else {
		goto L418
	}
L418:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1285 = m.ExcPending
	if v1285 != 0 {
		goto L40
	} else {
		goto L419
	}
L419:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v1288 = m.ExcPending
	if v1288 != 0 {
		goto L40
	} else {
		goto L420
	}
L420:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+100)) = int32(361922)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+96)) = int32(379442)
	F_errmsg(m, int32(146380), v29+int32(96))
	mBase = m.M
	v1297 = m.ExcPending
	if v1297 != 0 {
		goto L40
	} else {
		goto L421
	}
L421:
	;
	F_errfinish(m, int32(518572), int32(379), int32(145929))
	mBase = m.M
	v1302 = m.ExcPending
	if v1302 != 0 {
		goto L40
	} else {
		goto L422
	}
L422:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L423:
	;
	v1306 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3))))
	if v1306&int32(4) != 0 {
		goto L412
	} else {
		goto L426
	}
L424:
	;
	goto L425
L425:
	;
	v1309 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+15)))
	if v1309 == int32(1) {
		goto L427
	} else {
		goto L428
	}
L426:
	;
	goto L425
L427:
	;
	v1312 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3))))
	if v1312&int32(16) != 0 {
		goto L411
	} else {
		goto L430
	}
L428:
	;
	goto L429
L429:
	;
	v1315 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+15)) = uint8(v1315)
	*(*uint16)(unsafe.Add(mBase, uint32(l3)+13)) = uint16(v1315)
	goto L413
L430:
	;
	goto L429
L431:
	;
	m.G0 = v29 + int32(160)
	return
L432:
	;
	v1320 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v1320&int32(8) == int32(0) {
		goto L431
	} else {
		goto L433
	}
L433:
	;
	v1325 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+13)))
	if v1325 == int32(1) {
		goto L434
	} else {
		goto L435
	}
L434:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1331 = m.ExcPending
	if v1331 != 0 {
		goto L40
	} else {
		goto L437
	}
L435:
	;
	goto L436
L436:
	;
	v1349 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+14)))
	if v1349 != int32(1) {
		goto L431
	} else {
		goto L442
	}
L437:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v1334 = m.ExcPending
	if v1334 != 0 {
		goto L40
	} else {
		goto L438
	}
L438:
	;
	if v1320&int32(2) != 0 {
		goto L410
	} else {
		goto L439
	}
L439:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+4)) = int32(379458)
	*(*int32)(unsafe.Add(mBase, uint32(v29))) = int32(566020)
	F_errmsg(m, int32(189739), v29)
	mBase = m.M
	v1343 = m.ExcPending
	if v1343 != 0 {
		goto L40
	} else {
		goto L440
	}
L440:
	;
	F_errfinish(m, int32(518572), int32(421), int32(145929))
	mBase = m.M
	v1348 = m.ExcPending
	if v1348 != 0 {
		goto L40
	} else {
		goto L441
	}
L441:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L442:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1355 = m.ExcPending
	if v1355 != 0 {
		goto L40
	} else {
		goto L443
	}
L443:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v1358 = m.ExcPending
	if v1358 != 0 {
		goto L40
	} else {
		goto L444
	}
L444:
	;
	if v1320&int32(4) != 0 {
		goto L409
	} else {
		goto L445
	}
L445:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+36)) = int32(379422)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+32)) = int32(566020)
	F_errmsg(m, int32(189739), v29+int32(32))
	mBase = m.M
	v1369 = m.ExcPending
	if v1369 != 0 {
		goto L40
	} else {
		goto L446
	}
L446:
	;
	F_errfinish(m, int32(518572), int32(437), int32(145929))
	mBase = m.M
	v1374 = m.ExcPending
	if v1374 != 0 {
		goto L40
	} else {
		goto L447
	}
L447:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L448:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v1385 = m.ExcPending
	if v1385 != 0 {
		goto L40
	} else {
		goto L449
	}
L449:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+84)) = int32(361903)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+80)) = int32(379442)
	F_errmsg(m, int32(146380), v29+int32(80))
	mBase = m.M
	v1394 = m.ExcPending
	if v1394 != 0 {
		goto L40
	} else {
		goto L450
	}
L450:
	;
	F_errfinish(m, int32(518572), int32(386), int32(145929))
	mBase = m.M
	v1399 = m.ExcPending
	if v1399 != 0 {
		goto L40
	} else {
		goto L451
	}
L451:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L452:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v1406 = m.ExcPending
	if v1406 != 0 {
		goto L40
	} else {
		goto L453
	}
L453:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+68)) = int32(361937)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+64)) = int32(379442)
	F_errmsg(m, int32(146380), v29-int32(-64))
	mBase = m.M
	v1415 = m.ExcPending
	if v1415 != 0 {
		goto L40
	} else {
		goto L454
	}
L454:
	;
	F_errfinish(m, int32(518572), int32(393), int32(145929))
	mBase = m.M
	v1420 = m.ExcPending
	if v1420 != 0 {
		goto L40
	} else {
		goto L455
	}
L455:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L456:
	;
	F_errfinish(m, int32(518572), int32(415), int32(145929))
	mBase = m.M
	v1434 = m.ExcPending
	if v1434 != 0 {
		goto L40
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
	F_errfinish(m, int32(518572), int32(431), int32(145929))
	mBase = m.M
	v1448 = m.ExcPending
	if v1448 != 0 {
		goto L40
	} else {
		goto L459
	}
L459:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L460:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_subscription_change_cb(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	v5 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _consts[539])) = uint8(v5)
	return
}
