package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_pg_logical_emit_message_bytea(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v95 int64
	_ = v95
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v12 = F_pg_detoast_datum_packed(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v16 = F_text_to_cstring(m, v12)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
			v19 = F_pg_detoast_datum_packed(m, v18)
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int32(0)
			} else {
				v21 = int32(1)
				v22 = v19 + v21
				v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
				v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
				v28 = v26 & v21
				if v28 != 0 {
					v29 = v22
				} else {
					v29 = v19 + int32(4)
				}
				if v26 == int32(1) {
					v32 = int32(4)
					v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22))))
					if v34&int32(254) == int32(2) {
						v43 = v32
					} else {
						v43 = base.B2i32(v34 == int32(18)) << (uint(v32) % 32)
					}
					if v34 == int32(1) {
						v46 = v32
					} else {
						v46 = v43
					}
					v57 = v46
				} else {
					v47 = int32(1)
					if v28 != 0 {
						v57 = int32(base.Ui32(v26)>>(uint(v47)%32)) - v47
					} else {
						v51 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
						v57 = int32(base.Ui32(v51)>>(uint(int32(2))%32)) - int32(4)
					}
				}
				v58 = int32(0)
				v60 = m.G0
				v62 = v60 - int32(16)
				m.G0 = v62
				v65 = base.B2i32(v10 != v58)
				if v10 != v58 {
					v66 = F_GetCurrentTransactionId(m)
					mBase = m.M
					v67 = m.ExcPending
					if v67 != 0 {
						return int32(0)
					} else {
						*(*uint8)(unsafe.Add(mBase, uint32(v62)+4)) = uint8(v65)
						v70 = *(*int32)(unsafe.Add(mBase, _consts[107]))
						*(*int32)(unsafe.Add(mBase, uint32(v62))) = v70
						v72 = F_strlen(m, v16)
						mBase = m.M
						*(*int32)(unsafe.Add(mBase, uint32(v62)+12)) = v57
						*(*int32)(unsafe.Add(mBase, uint32(v62)+8)) = v72 + int32(1)
						F_XLogBeginInsert(m)
						mBase = m.M
						v78 = m.ExcPending
						if v78 != 0 {
							return int32(0)
						} else {
							F_XLogRegisterData(m, v62, int32(16))
							mBase = m.M
							v81 = m.ExcPending
							if v81 != 0 {
								return int32(0)
							} else {
								v82 = *(*int32)(unsafe.Add(mBase, uint32(v62)+8))
								F_XLogRegisterData(m, v16, v82)
								mBase = m.M
								v84 = m.ExcPending
								if v84 != 0 {
									return int32(0)
								} else {
									F_XLogRegisterData(m, v29, v57)
									mBase = m.M
									v86 = m.ExcPending
									if v86 != 0 {
										return int32(0)
									} else {
										v88 = int32(4411444)
										v90 = int32(*(*uint8)(unsafe.Add(mBase, _consts[43])))
										v91 = v90 | int32(1)
										*(*uint8)(unsafe.Add(mBase, _consts[43])) = uint8(v91)
										v95 = F_XLogInsert(m, int32(21), int32(0))
										mBase = m.M
										v96 = m.ExcPending
										if v96 != 0 {
											return int32(0)
										} else {
											if v10 != v58 {
												m.G0 = v62 + int32(16)
												v104 = F_Int64GetDatum(m, v95)
												mBase = m.M
												v105 = m.ExcPending
												if v105 != 0 {
													return int32(0)
												} else {
													return v104
												}
											} else {
												if base.B2i32(v23 != v58) == int32(0) {
													m.G0 = v62 + int32(16)
													v104 = F_Int64GetDatum(m, v95)
													mBase = m.M
													v105 = m.ExcPending
													if v105 != 0 {
														return int32(0)
													} else {
														return v104
													}
												} else {
													F_XLogFlush(m, v95)
													mBase = m.M
													v100 = m.ExcPending
													if v100 != 0 {
														return int32(0)
													} else {
														m.G0 = v62 + int32(16)
														v104 = F_Int64GetDatum(m, v95)
														mBase = m.M
														v105 = m.ExcPending
														if v105 != 0 {
															return int32(0)
														} else {
															return v104
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
				} else {
					*(*uint8)(unsafe.Add(mBase, uint32(v62)+4)) = uint8(v65)
					v70 = *(*int32)(unsafe.Add(mBase, _consts[107]))
					*(*int32)(unsafe.Add(mBase, uint32(v62))) = v70
					v72 = F_strlen(m, v16)
					mBase = m.M
					*(*int32)(unsafe.Add(mBase, uint32(v62)+12)) = v57
					*(*int32)(unsafe.Add(mBase, uint32(v62)+8)) = v72 + int32(1)
					F_XLogBeginInsert(m)
					mBase = m.M
					v78 = m.ExcPending
					if v78 != 0 {
						return int32(0)
					} else {
						F_XLogRegisterData(m, v62, int32(16))
						mBase = m.M
						v81 = m.ExcPending
						if v81 != 0 {
							return int32(0)
						} else {
							v82 = *(*int32)(unsafe.Add(mBase, uint32(v62)+8))
							F_XLogRegisterData(m, v16, v82)
							mBase = m.M
							v84 = m.ExcPending
							if v84 != 0 {
								return int32(0)
							} else {
								F_XLogRegisterData(m, v29, v57)
								mBase = m.M
								v86 = m.ExcPending
								if v86 != 0 {
									return int32(0)
								} else {
									v88 = int32(4411444)
									v90 = int32(*(*uint8)(unsafe.Add(mBase, _consts[43])))
									v91 = v90 | int32(1)
									*(*uint8)(unsafe.Add(mBase, _consts[43])) = uint8(v91)
									v95 = F_XLogInsert(m, int32(21), int32(0))
									mBase = m.M
									v96 = m.ExcPending
									if v96 != 0 {
										return int32(0)
									} else {
										if v10 != v58 {
											m.G0 = v62 + int32(16)
											v104 = F_Int64GetDatum(m, v95)
											mBase = m.M
											v105 = m.ExcPending
											if v105 != 0 {
												return int32(0)
											} else {
												return v104
											}
										} else {
											if base.B2i32(v23 != v58) == int32(0) {
												m.G0 = v62 + int32(16)
												v104 = F_Int64GetDatum(m, v95)
												mBase = m.M
												v105 = m.ExcPending
												if v105 != 0 {
													return int32(0)
												} else {
													return v104
												}
											} else {
												F_XLogFlush(m, v95)
												mBase = m.M
												v100 = m.ExcPending
												if v100 != 0 {
													return int32(0)
												} else {
													m.G0 = v62 + int32(16)
													v104 = F_Int64GetDatum(m, v95)
													mBase = m.M
													v105 = m.ExcPending
													if v105 != 0 {
														return int32(0)
													} else {
														return v104
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
func F_pg_logical_slot_get_changes_guts(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v32 int64
	_ = v32
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
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
	var v74 int32
	_ = v74
	var v86 int64
	_ = v86
	var v87 int64
	_ = v87
	var v88 int64
	_ = v88
	var v89 int64
	_ = v89
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
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v122 int32
	_ = v122
	var v128 int32
	_ = v128
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v167 int32
	_ = v167
	var v185 int32
	_ = v185
	var v204 int32
	_ = v204
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v230 int64
	_ = v230
	var v231 int64
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v259 int32
	_ = v259
	var v277 int32
	_ = v277
	var v296 int32
	_ = v296
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v382 int32
	_ = v382
	var v400 int32
	_ = v400
	var v419 int32
	_ = v419
	var v439 int32
	_ = v439
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v475 int32
	_ = v475
	var v493 int32
	_ = v493
	var v512 int32
	_ = v512
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v558 int32
	_ = v558
	var v573 int32
	_ = v573
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v599 int32
	_ = v599
	var v600 int32
	_ = v600
	var v602 int32
	_ = v602
	var v618 int32
	_ = v618
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v622 int32
	_ = v622
	var v638 int32
	_ = v638
	var v639 int32
	_ = v639
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v673 int32
	_ = v673
	var v674 int32
	_ = v674
	var v690 int32
	_ = v690
	var v691 int32
	_ = v691
	var v693 int32
	_ = v693
	var v694 int32
	_ = v694
	var v708 int32
	_ = v708
	var v711 int32
	_ = v711
	var v750 int32
	_ = v750
	var v751 int32
	_ = v751
	var v753 int32
	_ = v753
	var v772 int32
	_ = v772
	var v777 int32
	_ = v777
	var v778 int32
	_ = v778
	var v780 int32
	_ = v780
	var v782 int32
	_ = v782
	var v801 int32
	_ = v801
	var v802 int32
	_ = v802
	var v803 int64
	_ = v803
	var v808 int32
	_ = v808
	var v809 int64
	_ = v809
	var v818 int64
	_ = v818
	var v837 int32
	_ = v837
	var v855 int32
	_ = v855
	var v874 int32
	_ = v874
	var v894 int32
	_ = v894
	var v911 int64
	_ = v911
	var v912 int32
	_ = v912
	var v913 int64
	_ = v913
	var v914 int64
	_ = v914
	var v915 int64
	_ = v915
	var v931 int32
	_ = v931
	var v934 int32
	_ = v934
	var v937 int32
	_ = v937
	var v939 int32
	_ = v939
	var v949 int32
	_ = v949
	var v950 int32
	_ = v950
	var v951 int32
	_ = v951
	var v952 int32
	_ = v952
	var v953 int32
	_ = v953
	var v954 int32
	_ = v954
	var v955 int32
	_ = v955
	var v956 int32
	_ = v956
	var v957 int32
	_ = v957
	var v958 int32
	_ = v958
	var v960 int32
	_ = v960
	var v961 int32
	_ = v961
	var v964 int32
	_ = v964
	var v976 int64
	_ = v976
	var v977 int64
	_ = v977
	var v978 int64
	_ = v978
	var v979 int64
	_ = v979
	var v1008 int32
	_ = v1008
	var v1012 int32
	_ = v1012
	var v1013 int32
	_ = v1013
	var v1016 int32
	_ = v1016
	var v1037 int32
	_ = v1037
	var v1055 int32
	_ = v1055
	var v1056 int32
	_ = v1056
	var v1057 int32
	_ = v1057
	var v1074 int32
	_ = v1074
	var v1075 int32
	_ = v1075
	var v1076 int32
	_ = v1076
	var v1100 int32
	_ = v1100
	var v1120 int32
	_ = v1120
	var v1137 int64
	_ = v1137
	var v1139 int32
	_ = v1139
	var v1140 int64
	_ = v1140
	var v1142 int32
	_ = v1142
	var v1144 int32
	_ = v1144
	var v1146 int32
	_ = v1146
	var v1147 int64
	_ = v1147
	var v1164 int32
	_ = v1164
	var v1181 int32
	_ = v1181
	var v1182 int32
	_ = v1182
	var v1183 int64
	_ = v1183
	var v1202 int32
	_ = v1202
	var v1240 int32
	_ = v1240
	var v1241 int32
	_ = v1241
	var v1242 int32
	_ = v1242
	var v1261 int32
	_ = v1261
	var v1262 int32
	_ = v1262
	var v1281 int32
	_ = v1281
	var v1301 int32
	_ = v1301
	var v1302 int32
	_ = v1302
	var v1319 int32
	_ = v1319
	var v1323 int32
	_ = v1323
	var v1324 int64
	_ = v1324
	var v1329 int64
	_ = v1329
	var v1331 int32
	_ = v1331
	var v1332 int64
	_ = v1332
	var v1334 int32
	_ = v1334
	var v1351 int32
	_ = v1351
	var v1352 int32
	_ = v1352
	var v1353 int64
	_ = v1353
	var v1390 int64
	_ = v1390
	var v1414 int32
	_ = v1414
	var v1431 int32
	_ = v1431
	var v1448 int32
	_ = v1448
	var v1465 int32
	_ = v1465
	var v1482 int32
	_ = v1482
	var v1510 int32
	_ = v1510
	var v1527 int32
	_ = v1527
	var v1547 int32
	_ = v1547
	var v1565 int32
	_ = v1565
	var v1566 int64
	_ = v1566
	var v1570 int32
	_ = v1570
	var v1572 int32
	_ = v1572
	var v1573 int32
	_ = v1573
	var v1577 int32
	_ = v1577
	var v1579 int32
	_ = v1579
	var v1580 int32
	_ = v1580
	var v1581 int32
	_ = v1581
	var v1582 int32
	_ = v1582
	var v1583 int32
	_ = v1583
	var v1584 int32
	_ = v1584
	var v1585 int64
	_ = v1585
	var v1586 int32
	_ = v1586
	var v1587 int32
	_ = v1587
	var v1588 int32
	_ = v1588
	var v1589 int32
	_ = v1589
	var v1590 int32
	_ = v1590
	var v1591 int64
	_ = v1591
	var v1592 int64
	_ = v1592
	var v1593 int64
	_ = v1593
	var v1594 int32
	_ = v1594
	var v1595 int32
	_ = v1595
	var v1597 int32
	_ = v1597
	v3 = l2
	v4 = int32(0)
	v32 = int64(0)
	v38 = m.G0
	v40 = v38 - int32(112)
	m.G0 = v40
	v59 = v4
	v60 = v4
	v61 = v4
	v62 = v4
	v63 = v4
	v64 = v4
	v65 = v4
	v66 = v4
	v67 = v4
	v68 = v4
	v69 = v4
	v70 = int32(-1)
	v71 = v4
	v74 = v40
	v86 = v32
	v87 = v32
	v88 = v32
	v89 = v32
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
	if v70 == int32(1) {
		goto L8
	} else {
		goto L9
	}
L5:
	;
	goto L3
L6:
	;
	v1565 = int32(m.ExcTag)
	v1566 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v1565 == int32(0) {
		goto L144
	} else {
		goto L145
	}
L7:
	;
	if v961 == int32(0) {
		goto L86
	} else {
		goto L87
	}
L8:
	;
	v949 = v59
	v950 = v60
	v951 = v61
	v952 = v62
	v953 = v63
	v954 = v64
	v955 = v65
	v956 = v66
	v957 = v67
	v958 = v68
	v960 = v69
	v961 = v71
	v964 = v74
	v976 = v86
	v977 = v87
	v978 = v88
	v979 = v89
	goto L7
L9:
	;
	goto L10
L10:
	;
	v94 = int32(16)
	v95 = v74 - v94
	m.G0 = v95
	v98 = v95 - v94
	m.G0 = v98
	v101 = v98 - int32(160)
	m.G0 = v101
	v104 = v101 - v94
	m.G0 = v104
	v107 = v104 - v94
	m.G0 = v107
	v109 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v40)+32)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v40)+36)) = v65
	*(*int64)(unsafe.Add(mBase, uint32(v40)+40)) = v87
	*(*int64)(unsafe.Add(mBase, uint32(v40)+48)) = v88
	*(*int64)(unsafe.Add(mBase, uint32(v40)+56)) = v89
	*(*int32)(unsafe.Add(mBase, uint32(v40)+68)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v40)+72)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v40)+76)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v40)+80)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v40)+84)) = v62
	*(*int64)(unsafe.Add(mBase, uint32(v40)+88)) = v86
	v122 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	*(*int32)(unsafe.Add(mBase, uint32(v40)+96)) = v122
	*(*int32)(unsafe.Add(mBase, uint32(v40)+100)) = v107
	*(*int32)(unsafe.Add(mBase, uint32(v40)+104)) = v104
	*(*int32)(unsafe.Add(mBase, uint32(v40)+108)) = v101
	F_CheckSlotPermissions(m)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		v1547 = v107
		goto L6
	} else {
		goto L11
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+36)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v40)+32)) = v64
	*(*int64)(unsafe.Add(mBase, uint32(v40)+40)) = v87
	*(*int64)(unsafe.Add(mBase, uint32(v40)+48)) = v88
	*(*int64)(unsafe.Add(mBase, uint32(v40)+56)) = v89
	*(*int32)(unsafe.Add(mBase, uint32(v40)+68)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v40)+72)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v40)+76)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v40)+80)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v40)+84)) = v62
	*(*int64)(unsafe.Add(mBase, uint32(v40)+88)) = v86
	*(*int32)(unsafe.Add(mBase, uint32(v40)+96)) = v122
	*(*int32)(unsafe.Add(mBase, uint32(v40)+100)) = v107
	*(*int32)(unsafe.Add(mBase, uint32(v40)+104)) = v104
	*(*int32)(unsafe.Add(mBase, uint32(v40)+108)) = v101
	F_CheckLogicalDecodingRequirements(m)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		v1547 = v107
		goto L6
	} else {
		goto L12
	}
L12:
	;
	v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v146 == int32(1) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+36)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v40)+32)) = v64
	*(*int64)(unsafe.Add(mBase, uint32(v40)+40)) = v87
	*(*int64)(unsafe.Add(mBase, uint32(v40)+48)) = v88
	*(*int64)(unsafe.Add(mBase, uint32(v40)+56)) = v89
	*(*int32)(unsafe.Add(mBase, uint32(v40)+68)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v40)+72)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v40)+76)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v40)+80)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v40)+84)) = v62
	*(*int64)(unsafe.Add(mBase, uint32(v40)+88)) = v86
	*(*int32)(unsafe.Add(mBase, uint32(v40)+96)) = v122
	*(*int32)(unsafe.Add(mBase, uint32(v40)+100)) = v107
	*(*int32)(unsafe.Add(mBase, uint32(v40)+104)) = v104
	*(*int32)(unsafe.Add(mBase, uint32(v40)+108)) = v101
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		v1547 = v107
		goto L6
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(32)))))
	if v226 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+36)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v40)+32)) = v64
	*(*int64)(unsafe.Add(mBase, uint32(v40)+40)) = v87
	*(*int64)(unsafe.Add(mBase, uint32(v40)+48)) = v88
	*(*int64)(unsafe.Add(mBase, uint32(v40)+56)) = v89
	*(*int32)(unsafe.Add(mBase, uint32(v40)+68)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v40)+72)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v40)+76)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v40)+80)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v40)+84)) = v62
	*(*int64)(unsafe.Add(mBase, uint32(v40)+88)) = v86
	*(*int32)(unsafe.Add(mBase, uint32(v40)+96)) = v122
	*(*int32)(unsafe.Add(mBase, uint32(v40)+100)) = v107
	*(*int32)(unsafe.Add(mBase, uint32(v40)+104)) = v104
	*(*int32)(unsafe.Add(mBase, uint32(v40)+108)) = v101
	F_errcode(m, int32(67108994))
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		v1547 = v107
		goto L6
	} else {
		goto L17
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+36)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v40)+32)) = v64
	*(*int64)(unsafe.Add(mBase, uint32(v40)+40)) = v87
	*(*int64)(unsafe.Add(mBase, uint32(v40)+48)) = v88
	*(*int64)(unsafe.Add(mBase, uint32(v40)+56)) = v89
	*(*int32)(unsafe.Add(mBase, uint32(v40)+68)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v40)+72)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v40)+76)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v40)+80)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v40)+84)) = v62
	*(*int64)(unsafe.Add(mBase, uint32(v40)+88)) = v86
	*(*int32)(unsafe.Add(mBase, uint32(v40)+96)) = v122
	*(*int32)(unsafe.Add(mBase, uint32(v40)+100)) = v107
	*(*int32)(unsafe.Add(mBase, uint32(v40)+104)) = v104
	*(*int32)(unsafe.Add(mBase, uint32(v40)+108)) = v101
	F_errmsg(m, int32(303491), int32(0))
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		v1547 = v107
		goto L6
	} else {
		goto L18
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+36)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v40)+32)) = v64
	*(*int64)(unsafe.Add(mBase, uint32(v40)+40)) = v87
	*(*int64)(unsafe.Add(mBase, uint32(v40)+48)) = v88
	*(*int64)(unsafe.Add(mBase, uint32(v40)+56)) = v89
	*(*int32)(unsafe.Add(mBase, uint32(v40)+68)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v40)+72)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v40)+76)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v40)+80)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v40)+84)) = v62
	*(*int64)(unsafe.Add(mBase, uint32(v40)+88)) = v86
	*(*int32)(unsafe.Add(mBase, uint32(v40)+96)) = v122
	*(*int32)(unsafe.Add(mBase, uint32(v40)+100)) = v107
	*(*int32)(unsafe.Add(mBase, uint32(v40)+104)) = v104
	*(*int32)(unsafe.Add(mBase, uint32(v40)+108)) = v101
	F_errfinish(m, int32(495112), int32(123), int32(115526))
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		v1547 = v107
		goto L6
	} else {
		goto L19
	}
L19:
	;
	goto L3
L20:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(28))))
	v230 = *(*int64)(unsafe.Add(mBase, uint32(v229)))
	v231 = v230
	goto L22
L21:
	;
	v231 = int64(0)
	goto L22
L22:
	;
	v232 = int32(0)
	v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(40)))))
	if v233 == v232 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v236 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(36))))
	v237 = v236
	goto L25
L24:
	;
	v237 = v232
	goto L25
L25:
	;
	v238 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(48)))))
	if v238 == int32(1) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+36)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v40)+32)) = v64
	*(*int64)(unsafe.Add(mBase, uint32(v40)+40)) = v87
	*(*int64)(unsafe.Add(mBase, uint32(v40)+48)) = v88
	*(*int64)(unsafe.Add(mBase, uint32(v40)+56)) = v89
	*(*int32)(unsafe.Add(mBase, uint32(v40)+68)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v40)+72)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v40)+76)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v40)+80)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v40)+84)) = v237
	*(*int64)(unsafe.Add(mBase, uint32(v40)+88)) = v231
	*(*int32)(unsafe.Add(mBase, uint32(v40)+96)) = v122
	*(*int32)(unsafe.Add(mBase, uint32(v40)+100)) = v107
	*(*int32)(unsafe.Add(mBase, uint32(v40)+104)) = v104
	*(*int32)(unsafe.Add(mBase, uint32(v40)+108)) = v101
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		v1547 = v107
		goto L6
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	v317 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v318 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(44))))
	*(*int32)(unsafe.Add(mBase, uint32(v40)+36)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v40)+32)) = v64
	*(*int64)(unsafe.Add(mBase, uint32(v40)+40)) = v87
	*(*int64)(unsafe.Add(mBase, uint32(v40)+48)) = v88
	*(*int64)(unsafe.Add(mBase, uint32(v40)+56)) = v89
	*(*int32)(unsafe.Add(mBase, uint32(v40)+68)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v40)+72)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v40)+76)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v40)+80)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v40)+84)) = v237
	*(*int64)(unsafe.Add(mBase, uint32(v40)+88)) = v231
	*(*int32)(unsafe.Add(mBase, uint32(v40)+96)) = v122
	*(*int32)(unsafe.Add(mBase, uint32(v40)+100)) = v107
	*(*int32)(unsafe.Add(mBase, uint32(v40)+104)) = v104
	*(*int32)(unsafe.Add(mBase, uint32(v40)+108)) = v101
	v334 = F_pg_detoast_datum(m, v318)
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		v1547 = v107
		goto L6
	} else {
		goto L33
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+36)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v40)+32)) = v64
	*(*int64)(unsafe.Add(mBase, uint32(v40)+40)) = v87
	*(*int64)(unsafe.Add(mBase, uint32(v40)+48)) = v88
	*(*int64)(unsafe.Add(mBase, uint32(v40)+56)) = v89
	*(*int32)(unsafe.Add(mBase, uint32(v40)+68)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v40)+72)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v40)+76)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v40)+80)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v40)+84)) = v237
	*(*int64)(unsafe.Add(mBase, uint32(v40)+88)) = v231
	*(*int32)(unsafe.Add(mBase, uint32(v40)+96)) = v122
	*(*int32)(unsafe.Add(mBase, uint32(v40)+100)) = v107
	*(*int32)(unsafe.Add(mBase, uint32(v40)+104)) = v104
	*(*int32)(unsafe.Add(mBase, uint32(v40)+108)) = v101
	F_errcode(m, int32(67108994))
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		v1547 = v107
		goto L6
	} else {
		goto L30
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+36)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v40)+32)) = v64
	*(*int64)(unsafe.Add(mBase, uint32(v40)+40)) = v87
	*(*int64)(unsafe.Add(mBase, uint32(v40)+48)) = v88
	*(*int64)(unsafe.Add(mBase, uint32(v40)+56)) = v89
	*(*int32)(unsafe.Add(mBase, uint32(v40)+68)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v40)+72)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v40)+76)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v40)+80)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v40)+84)) = v237
	*(*int64)(unsafe.Add(mBase, uint32(v40)+88)) = v231
	*(*int32)(unsafe.Add(mBase, uint32(v40)+96)) = v122
	*(*int32)(unsafe.Add(mBase, uint32(v40)+100)) = v107
	*(*int32)(unsafe.Add(mBase, uint32(v40)+104)) = v104
	*(*int32)(unsafe.Add(mBase, uint32(v40)+108)) = v101
	F_errmsg(m, int32(302972), int32(0))
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		v1547 = v107
		goto L6
	} else {
		goto L31
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+36)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v40)+32)) = v64
	*(*int64)(unsafe.Add(mBase, uint32(v40)+40)) = v87
	*(*int64)(unsafe.Add(mBase, uint32(v40)+48)) = v88
	*(*int64)(unsafe.Add(mBase, uint32(v40)+56)) = v89
	*(*int32)(unsafe.Add(mBase, uint32(v40)+68)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v40)+72)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v40)+76)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v40)+80)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v40)+84)) = v237
	*(*int64)(unsafe.Add(mBase, uint32(v40)+88)) = v231
	*(*int32)(unsafe.Add(mBase, uint32(v40)+96)) = v122
	*(*int32)(unsafe.Add(mBase, uint32(v40)+100)) = v107
	*(*int32)(unsafe.Add(mBase, uint32(v40)+104)) = v104
	*(*int32)(unsafe.Add(mBase, uint32(v40)+108)) = v101
	F_errfinish(m, int32(495112), int32(139), int32(115526))
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		v1547 = v107
		goto L6
	} else {
		goto L32
	}
L32:
	;
	goto L3
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+36)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v40)+32)) = v64
	*(*int64)(unsafe.Add(mBase, uint32(v40)+40)) = v87
	*(*int64)(unsafe.Add(mBase, uint32(v40)+48)) = v88
	*(*int64)(unsafe.Add(mBase, uint32(v40)+56)) = v89
	*(*int32)(unsafe.Add(mBase, uint32(v40)+68)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v40)+72)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v40)+76)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v40)+80)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v40)+84)) = v237
	*(*int64)(unsafe.Add(mBase, uint32(v40)+88)) = v231
	*(*int32)(unsafe.Add(mBase, uint32(v40)+96)) = v122
	*(*int32)(unsafe.Add(mBase, uint32(v40)+100)) = v107
	*(*int32)(unsafe.Add(mBase, uint32(v40)+104)) = v104
	*(*int32)(unsafe.Add(mBase, uint32(v40)+108)) = v101
	v352 = F_palloc0(m, int32(24))
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		v1547 = v107
		goto L6
	} else {
		goto L34
	}
L34:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v352)+8)) = uint8(v3)
	v355 = int32(4515600)
	v356 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v358 = *(*int32)(unsafe.Add(mBase, uint32(v109)+4))
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v358)+16))
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v359
	v361 = *(*int32)(unsafe.Add(mBase, uint32(v334)+4))
	if base.Ui32(int32(2)) <= base.Ui32(v361) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+36)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v40)+32)) = v64
	*(*int64)(unsafe.Add(mBase, uint32(v40)+40)) = v87
	*(*int64)(unsafe.Add(mBase, uint32(v40)+48)) = v88
	*(*int64)(unsafe.Add(mBase, uint32(v40)+56)) = v89
	*(*int32)(unsafe.Add(mBase, uint32(v40)+68)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v40)+72)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v40)+76)) = v356
	*(*int32)(unsafe.Add(mBase, uint32(v40)+80)) = v352
	*(*int32)(unsafe.Add(mBase, uint32(v40)+84)) = v237
	*(*int64)(unsafe.Add(mBase, uint32(v40)+88)) = v231
	*(*int32)(unsafe.Add(mBase, uint32(v40)+96)) = v122
	*(*int32)(unsafe.Add(mBase, uint32(v40)+100)) = v107
	*(*int32)(unsafe.Add(mBase, uint32(v40)+104)) = v104
	*(*int32)(unsafe.Add(mBase, uint32(v40)+108)) = v101
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		v1547 = v107
		goto L6
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+36)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v40)+32)) = v64
	*(*int64)(unsafe.Add(mBase, uint32(v40)+40)) = v87
	*(*int64)(unsafe.Add(mBase, uint32(v40)+48)) = v88
	*(*int64)(unsafe.Add(mBase, uint32(v40)+56)) = v89
	*(*int32)(unsafe.Add(mBase, uint32(v40)+68)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v40)+72)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v40)+76)) = v356
	*(*int32)(unsafe.Add(mBase, uint32(v40)+80)) = v352
	*(*int32)(unsafe.Add(mBase, uint32(v40)+84)) = v237
	*(*int64)(unsafe.Add(mBase, uint32(v40)+88)) = v231
	*(*int32)(unsafe.Add(mBase, uint32(v40)+96)) = v122
	*(*int32)(unsafe.Add(mBase, uint32(v40)+100)) = v107
	*(*int32)(unsafe.Add(mBase, uint32(v40)+104)) = v104
	*(*int32)(unsafe.Add(mBase, uint32(v40)+108)) = v101
	v455 = F_array_contains_nulls(m, v334)
	mBase = m.M
	v456 = m.ExcPending
	if v456 != 0 {
		v1547 = v107
		goto L6
	} else {
		goto L42
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+36)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v40)+32)) = v64
	*(*int64)(unsafe.Add(mBase, uint32(v40)+40)) = v87
	*(*int64)(unsafe.Add(mBase, uint32(v40)+48)) = v88
	*(*int64)(unsafe.Add(mBase, uint32(v40)+56)) = v89
	*(*int32)(unsafe.Add(mBase, uint32(v40)+68)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v40)+72)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v40)+76)) = v356
	*(*int32)(unsafe.Add(mBase, uint32(v40)+80)) = v352
	*(*int32)(unsafe.Add(mBase, uint32(v40)+84)) = v237
	*(*int64)(unsafe.Add(mBase, uint32(v40)+88)) = v231
	*(*int32)(unsafe.Add(mBase, uint32(v40)+96)) = v122
	*(*int32)(unsafe.Add(mBase, uint32(v40)+100)) = v107
	*(*int32)(unsafe.Add(mBase, uint32(v40)+104)) = v104
	*(*int32)(unsafe.Add(mBase, uint32(v40)+108)) = v101
	F_errcode(m, int32(1088))
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		v1547 = v107
		goto L6
	} else {
		goto L39
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+36)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v40)+32)) = v64
	*(*int64)(unsafe.Add(mBase, uint32(v40)+40)) = v87
	*(*int64)(unsafe.Add(mBase, uint32(v40)+48)) = v88
	*(*int64)(unsafe.Add(mBase, uint32(v40)+56)) = v89
	*(*int32)(unsafe.Add(mBase, uint32(v40)+68)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v40)+72)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v40)+76)) = v356
	*(*int32)(unsafe.Add(mBase, uint32(v40)+80)) = v352
	*(*int32)(unsafe.Add(mBase, uint32(v40)+84)) = v237
	*(*int64)(unsafe.Add(mBase, uint32(v40)+88)) = v231
	*(*int32)(unsafe.Add(mBase, uint32(v40)+96)) = v122
	*(*int32)(unsafe.Add(mBase, uint32(v40)+100)) = v107
	*(*int32)(unsafe.Add(mBase, uint32(v40)+104)) = v104
	*(*int32)(unsafe.Add(mBase, uint32(v40)+108)) = v101
	F_errmsg(m, int32(313225), int32(0))
	mBase = m.M
	v419 = m.ExcPending
	if v419 != 0 {
		v1547 = v107
		goto L6
	} else {
		goto L40
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+36)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v40)+32)) = v64
	*(*int64)(unsafe.Add(mBase, uint32(v40)+40)) = v87
	*(*int64)(unsafe.Add(mBase, uint32(v40)+48)) = v88
	*(*int64)(unsafe.Add(mBase, uint32(v40)+56)) = v89
	*(*int32)(unsafe.Add(mBase, uint32(v40)+68)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v40)+72)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v40)+76)) = v356
	*(*int32)(unsafe.Add(mBase, uint32(v40)+80)) = v352
	*(*int32)(unsafe.Add(mBase, uint32(v40)+84)) = v237
	*(*int64)(unsafe.Add(mBase, uint32(v40)+88)) = v231
	*(*int32)(unsafe.Add(mBase, uint32(v40)+96)) = v122
	*(*int32)(unsafe.Add(mBase, uint32(v40)+100)) = v107
	*(*int32)(unsafe.Add(mBase, uint32(v40)+104)) = v104
	*(*int32)(unsafe.Add(mBase, uint32(v40)+108)) = v101
	F_errfinish(m, int32(495112), int32(156), int32(115526))
	mBase = m.M
	v439 = m.ExcPending
	if v439 != 0 {
		v1547 = v107
		goto L6
	} else {
		goto L41
	}
L41:
	;
	goto L3
L42:
	;
	if v455 != 0 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+36)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v40)+32)) = v64
	*(*int64)(unsafe.Add(mBase, uint32(v40)+40)) = v87
	*(*int64)(unsafe.Add(mBase, uint32(v40)+48)) = v88
	*(*int64)(unsafe.Add(mBase, uint32(v40)+56)) = v89
	*(*int32)(unsafe.Add(mBase, uint32(v40)+68)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v40)+72)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v40)+76)) = v356
	*(*int32)(unsafe.Add(mBase, uint32(v40)+80)) = v352
	*(*int32)(unsafe.Add(mBase, uint32(v40)+84)) = v237
	*(*int64)(unsafe.Add(mBase, uint32(v40)+88)) = v231
	*(*int32)(unsafe.Add(mBase, uint32(v40)+96)) = v122
	*(*int32)(unsafe.Add(mBase, uint32(v40)+100)) = v107
	*(*int32)(unsafe.Add(mBase, uint32(v40)+104)) = v104
	*(*int32)(unsafe.Add(mBase, uint32(v40)+108)) = v101
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v475 = m.ExcPending
	if v475 != 0 {
		v1547 = v107
		goto L6
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	v533 = int32(0)
	if v361 != int32(1) {
		v708 = v67
		v711 = v533
		goto L53
	} else {
		goto L54
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+36)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v40)+32)) = v64
	*(*int64)(unsafe.Add(mBase, uint32(v40)+40)) = v87
	*(*int64)(unsafe.Add(mBase, uint32(v40)+48)) = v88
	*(*int64)(unsafe.Add(mBase, uint32(v40)+56)) = v89
	*(*int32)(unsafe.Add(mBase, uint32(v40)+68)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v40)+72)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v40)+76)) = v356
	*(*int32)(unsafe.Add(mBase, uint32(v40)+80)) = v352
	*(*int32)(unsafe.Add(mBase, uint32(v40)+84)) = v237
	*(*int64)(unsafe.Add(mBase, uint32(v40)+88)) = v231
	*(*int32)(unsafe.Add(mBase, uint32(v40)+96)) = v122
	*(*int32)(unsafe.Add(mBase, uint32(v40)+100)) = v107
	*(*int32)(unsafe.Add(mBase, uint32(v40)+104)) = v104
	*(*int32)(unsafe.Add(mBase, uint32(v40)+108)) = v101
	F_errcode(m, int32(1088))
	mBase = m.M
	v493 = m.ExcPending
	if v493 != 0 {
		v1547 = v107
		goto L6
	} else {
		goto L47
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+36)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v40)+32)) = v64
	*(*int64)(unsafe.Add(mBase, uint32(v40)+40)) = v87
	*(*int64)(unsafe.Add(mBase, uint32(v40)+48)) = v88
	*(*int64)(unsafe.Add(mBase, uint32(v40)+56)) = v89
	*(*int32)(unsafe.Add(mBase, uint32(v40)+68)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v40)+72)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v40)+76)) = v356
	*(*int32)(unsafe.Add(mBase, uint32(v40)+80)) = v352
	*(*int32)(unsafe.Add(mBase, uint32(v40)+84)) = v237
	*(*int64)(unsafe.Add(mBase, uint32(v40)+88)) = v231
	*(*int32)(unsafe.Add(mBase, uint32(v40)+96)) = v122
	*(*int32)(unsafe.Add(mBase, uint32(v40)+100)) = v107
	*(*int32)(unsafe.Add(mBase, uint32(v40)+104)) = v104
	*(*int32)(unsafe.Add(mBase, uint32(v40)+108)) = v101
	F_errmsg(m, int32(152677), int32(0))
	mBase = m.M
	v512 = m.ExcPending
	if v512 != 0 {
		v1547 = v107
		goto L6
	} else {
		goto L48
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+36)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v40)+32)) = v64
	*(*int64)(unsafe.Add(mBase, uint32(v40)+40)) = v87
	*(*int64)(unsafe.Add(mBase, uint32(v40)+48)) = v88
	*(*int64)(unsafe.Add(mBase, uint32(v40)+56)) = v89
	*(*int32)(unsafe.Add(mBase, uint32(v40)+68)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v40)+72)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v40)+76)) = v356
	*(*int32)(unsafe.Add(mBase, uint32(v40)+80)) = v352
	*(*int32)(unsafe.Add(mBase, uint32(v40)+84)) = v237
	*(*int64)(unsafe.Add(mBase, uint32(v40)+88)) = v231
	*(*int32)(unsafe.Add(mBase, uint32(v40)+96)) = v122
	*(*int32)(unsafe.Add(mBase, uint32(v40)+100)) = v107
	*(*int32)(unsafe.Add(mBase, uint32(v40)+104)) = v104
	*(*int32)(unsafe.Add(mBase, uint32(v40)+108)) = v101
	F_errfinish(m, int32(495112), int32(162), int32(115526))
	mBase = m.M
	v532 = m.ExcPending
	if v532 != 0 {
		v1547 = v107
		goto L6
	} else {
		goto L49
	}
L49:
	;
	goto L3
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+36)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v40)+32)) = v64
	*(*int64)(unsafe.Add(mBase, uint32(v40)+40)) = v915
	*(*int64)(unsafe.Add(mBase, uint32(v40)+48)) = v913
	*(*int64)(unsafe.Add(mBase, uint32(v40)+56)) = v914
	*(*int32)(unsafe.Add(mBase, uint32(v40)+68)) = v711
	*(*int32)(unsafe.Add(mBase, uint32(v40)+72)) = v708
	*(*int32)(unsafe.Add(mBase, uint32(v40)+76)) = v356
	*(*int32)(unsafe.Add(mBase, uint32(v40)+80)) = v352
	*(*int32)(unsafe.Add(mBase, uint32(v40)+84)) = v237
	*(*int64)(unsafe.Add(mBase, uint32(v40)+88)) = v231
	*(*int32)(unsafe.Add(mBase, uint32(v40)+96)) = v122
	*(*int32)(unsafe.Add(mBase, uint32(v40)+100)) = v107
	*(*int32)(unsafe.Add(mBase, uint32(v40)+104)) = v104
	*(*int32)(unsafe.Add(mBase, uint32(v40)+108)) = v101
	v931 = int32(1)
	F_ReplicationSlotAcquire(m, v317, v931, v931)
	mBase = m.M
	v934 = m.ExcPending
	if v934 != 0 {
		v1547 = v107
		goto L6
	} else {
		goto L81
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+36)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v40)+32)) = v64
	*(*int64)(unsafe.Add(mBase, uint32(v40)+40)) = v87
	*(*int64)(unsafe.Add(mBase, uint32(v40)+48)) = v88
	*(*int64)(unsafe.Add(mBase, uint32(v40)+56)) = v89
	*(*int32)(unsafe.Add(mBase, uint32(v40)+68)) = v711
	*(*int32)(unsafe.Add(mBase, uint32(v40)+72)) = v708
	*(*int32)(unsafe.Add(mBase, uint32(v40)+76)) = v356
	*(*int32)(unsafe.Add(mBase, uint32(v40)+80)) = v352
	*(*int32)(unsafe.Add(mBase, uint32(v40)+84)) = v237
	*(*int64)(unsafe.Add(mBase, uint32(v40)+88)) = v231
	*(*int32)(unsafe.Add(mBase, uint32(v40)+96)) = v122
	*(*int32)(unsafe.Add(mBase, uint32(v40)+100)) = v107
	*(*int32)(unsafe.Add(mBase, uint32(v40)+104)) = v104
	*(*int32)(unsafe.Add(mBase, uint32(v40)+108)) = v101
	v911 = F_GetXLogReplayRecPtr(m, int32(0))
	mBase = m.M
	v912 = m.ExcPending
	if v912 != 0 {
		v1547 = v107
		goto L6
	} else {
		goto L80
	}
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+36)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v40)+32)) = v64
	*(*int64)(unsafe.Add(mBase, uint32(v40)+40)) = v87
	*(*int64)(unsafe.Add(mBase, uint32(v40)+48)) = v88
	*(*int64)(unsafe.Add(mBase, uint32(v40)+56)) = v89
	*(*int32)(unsafe.Add(mBase, uint32(v40)+68)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v40)+72)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v40)+76)) = v356
	*(*int32)(unsafe.Add(mBase, uint32(v40)+80)) = v352
	*(*int32)(unsafe.Add(mBase, uint32(v40)+84)) = v237
	*(*int64)(unsafe.Add(mBase, uint32(v40)+88)) = v231
	*(*int32)(unsafe.Add(mBase, uint32(v40)+96)) = v122
	*(*int32)(unsafe.Add(mBase, uint32(v40)+100)) = v107
	*(*int32)(unsafe.Add(mBase, uint32(v40)+104)) = v104
	*(*int32)(unsafe.Add(mBase, uint32(v40)+108)) = v101
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v837 = m.ExcPending
	if v837 != 0 {
		v1547 = v107
		goto L6
	} else {
		goto L76
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+36)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v40)+32)) = v64
	*(*int64)(unsafe.Add(mBase, uint32(v40)+40)) = v87
	*(*int64)(unsafe.Add(mBase, uint32(v40)+48)) = v88
	*(*int64)(unsafe.Add(mBase, uint32(v40)+56)) = v89
	*(*int32)(unsafe.Add(mBase, uint32(v40)+68)) = v711
	*(*int32)(unsafe.Add(mBase, uint32(v40)+72)) = v708
	*(*int32)(unsafe.Add(mBase, uint32(v40)+76)) = v356
	*(*int32)(unsafe.Add(mBase, uint32(v40)+80)) = v352
	*(*int32)(unsafe.Add(mBase, uint32(v40)+84)) = v237
	*(*int64)(unsafe.Add(mBase, uint32(v40)+88)) = v231
	*(*int32)(unsafe.Add(mBase, uint32(v40)+96)) = v122
	*(*int32)(unsafe.Add(mBase, uint32(v40)+100)) = v107
	*(*int32)(unsafe.Add(mBase, uint32(v40)+104)) = v104
	*(*int32)(unsafe.Add(mBase, uint32(v40)+108)) = v101
	F_InitMaterializedSRF(m, l0, int32(0))
	mBase = m.M
	v750 = m.ExcPending
	if v750 != 0 {
		v1547 = v107
		goto L6
	} else {
		goto L66
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+36)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v40)+32)) = v64
	*(*int64)(unsafe.Add(mBase, uint32(v40)+40)) = v87
	*(*int64)(unsafe.Add(mBase, uint32(v40)+48)) = v88
	*(*int64)(unsafe.Add(mBase, uint32(v40)+56)) = v89
	*(*int32)(unsafe.Add(mBase, uint32(v40)+68)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v40)+72)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v40)+76)) = v356
	*(*int32)(unsafe.Add(mBase, uint32(v40)+80)) = v352
	*(*int32)(unsafe.Add(mBase, uint32(v40)+84)) = v237
	*(*int64)(unsafe.Add(mBase, uint32(v40)+88)) = v231
	*(*int32)(unsafe.Add(mBase, uint32(v40)+96)) = v122
	*(*int32)(unsafe.Add(mBase, uint32(v40)+100)) = v107
	*(*int32)(unsafe.Add(mBase, uint32(v40)+104)) = v104
	*(*int32)(unsafe.Add(mBase, uint32(v40)+108)) = v101
	F_deconstruct_array_builtin(m, v334, int32(25), v98, int32(0), v95)
	mBase = m.M
	v554 = m.ExcPending
	if v554 != 0 {
		v1547 = v107
		goto L6
	} else {
		goto L55
	}
L55:
	;
	v555 = *(*int32)(unsafe.Add(mBase, uint32(v95)))
	if v555&int32(1) != 0 {
		goto L52
	} else {
		goto L56
	}
L56:
	;
	v558 = int32(0)
	if v555 <= v558 {
		v708 = v67
		v711 = v533
		goto L53
	} else {
		goto L57
	}
L57:
	;
	v573 = v67
	v576 = v533
	v577 = v558
	goto L58
L58:
	;
	v599 = v577 << (uint(int32(2)) % 32)
	v600 = *(*int32)(unsafe.Add(mBase, uint32(v98)))
	v602 = *(*int32)(unsafe.Add(mBase, uint32(v599+v600)))
	*(*int32)(unsafe.Add(mBase, uint32(v40)+36)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v40)+32)) = v64
	*(*int64)(unsafe.Add(mBase, uint32(v40)+40)) = v87
	*(*int64)(unsafe.Add(mBase, uint32(v40)+48)) = v88
	*(*int64)(unsafe.Add(mBase, uint32(v40)+56)) = v89
	*(*int32)(unsafe.Add(mBase, uint32(v40)+68)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v40)+72)) = v573
	*(*int32)(unsafe.Add(mBase, uint32(v40)+76)) = v356
	*(*int32)(unsafe.Add(mBase, uint32(v40)+80)) = v352
	*(*int32)(unsafe.Add(mBase, uint32(v40)+84)) = v237
	*(*int64)(unsafe.Add(mBase, uint32(v40)+88)) = v231
	*(*int32)(unsafe.Add(mBase, uint32(v40)+96)) = v122
	*(*int32)(unsafe.Add(mBase, uint32(v40)+100)) = v107
	*(*int32)(unsafe.Add(mBase, uint32(v40)+104)) = v104
	*(*int32)(unsafe.Add(mBase, uint32(v40)+108)) = v101
	v618 = F_text_to_cstring(m, v602)
	mBase = m.M
	v619 = m.ExcPending
	if v619 != 0 {
		v1547 = v107
		goto L6
	} else {
		goto L60
	}
L59:
	;
	v708 = v690
	v711 = v690
	goto L53
L60:
	;
	v620 = *(*int32)(unsafe.Add(mBase, uint32(v98)))
	v622 = *(*int32)(unsafe.Add(mBase, uint32(v620+v599)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v40)+36)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v40)+32)) = v64
	*(*int64)(unsafe.Add(mBase, uint32(v40)+40)) = v87
	*(*int64)(unsafe.Add(mBase, uint32(v40)+48)) = v88
	*(*int64)(unsafe.Add(mBase, uint32(v40)+56)) = v89
	*(*int32)(unsafe.Add(mBase, uint32(v40)+68)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v40)+72)) = v573
	*(*int32)(unsafe.Add(mBase, uint32(v40)+76)) = v356
	*(*int32)(unsafe.Add(mBase, uint32(v40)+80)) = v352
	*(*int32)(unsafe.Add(mBase, uint32(v40)+84)) = v237
	*(*int64)(unsafe.Add(mBase, uint32(v40)+88)) = v231
	*(*int32)(unsafe.Add(mBase, uint32(v40)+96)) = v122
	*(*int32)(unsafe.Add(mBase, uint32(v40)+100)) = v107
	*(*int32)(unsafe.Add(mBase, uint32(v40)+104)) = v104
	*(*int32)(unsafe.Add(mBase, uint32(v40)+108)) = v101
	v638 = F_text_to_cstring(m, v622)
	mBase = m.M
	v639 = m.ExcPending
	if v639 != 0 {
		v1547 = v107
		goto L6
	} else {
		goto L61
	}
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+36)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v40)+32)) = v64
	*(*int64)(unsafe.Add(mBase, uint32(v40)+40)) = v87
	*(*int64)(unsafe.Add(mBase, uint32(v40)+48)) = v88
	*(*int64)(unsafe.Add(mBase, uint32(v40)+56)) = v89
	*(*int32)(unsafe.Add(mBase, uint32(v40)+68)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v40)+72)) = v573
	*(*int32)(unsafe.Add(mBase, uint32(v40)+76)) = v356
	*(*int32)(unsafe.Add(mBase, uint32(v40)+80)) = v352
	*(*int32)(unsafe.Add(mBase, uint32(v40)+84)) = v237
	*(*int64)(unsafe.Add(mBase, uint32(v40)+88)) = v231
	*(*int32)(unsafe.Add(mBase, uint32(v40)+96)) = v122
	*(*int32)(unsafe.Add(mBase, uint32(v40)+100)) = v107
	*(*int32)(unsafe.Add(mBase, uint32(v40)+104)) = v104
	*(*int32)(unsafe.Add(mBase, uint32(v40)+108)) = v101
	v655 = F_makeString(m, v638)
	mBase = m.M
	v656 = m.ExcPending
	if v656 != 0 {
		v1547 = v107
		goto L6
	} else {
		goto L62
	}
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+36)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v40)+32)) = v64
	*(*int64)(unsafe.Add(mBase, uint32(v40)+40)) = v87
	*(*int64)(unsafe.Add(mBase, uint32(v40)+48)) = v88
	*(*int64)(unsafe.Add(mBase, uint32(v40)+56)) = v89
	*(*int32)(unsafe.Add(mBase, uint32(v40)+68)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v40)+72)) = v573
	*(*int32)(unsafe.Add(mBase, uint32(v40)+76)) = v356
	*(*int32)(unsafe.Add(mBase, uint32(v40)+80)) = v352
	*(*int32)(unsafe.Add(mBase, uint32(v40)+84)) = v237
	*(*int64)(unsafe.Add(mBase, uint32(v40)+88)) = v231
	*(*int32)(unsafe.Add(mBase, uint32(v40)+96)) = v122
	*(*int32)(unsafe.Add(mBase, uint32(v40)+100)) = v107
	*(*int32)(unsafe.Add(mBase, uint32(v40)+104)) = v104
	*(*int32)(unsafe.Add(mBase, uint32(v40)+108)) = v101
	v673 = F_makeDefElem(m, v618, v655, int32(-1))
	mBase = m.M
	v674 = m.ExcPending
	if v674 != 0 {
		v1547 = v107
		goto L6
	} else {
		goto L63
	}
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+36)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v40)+32)) = v64
	*(*int64)(unsafe.Add(mBase, uint32(v40)+40)) = v87
	*(*int64)(unsafe.Add(mBase, uint32(v40)+48)) = v88
	*(*int64)(unsafe.Add(mBase, uint32(v40)+56)) = v89
	*(*int32)(unsafe.Add(mBase, uint32(v40)+68)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v40)+72)) = v573
	*(*int32)(unsafe.Add(mBase, uint32(v40)+76)) = v356
	*(*int32)(unsafe.Add(mBase, uint32(v40)+80)) = v352
	*(*int32)(unsafe.Add(mBase, uint32(v40)+84)) = v237
	*(*int64)(unsafe.Add(mBase, uint32(v40)+88)) = v231
	*(*int32)(unsafe.Add(mBase, uint32(v40)+96)) = v122
	*(*int32)(unsafe.Add(mBase, uint32(v40)+100)) = v107
	*(*int32)(unsafe.Add(mBase, uint32(v40)+104)) = v104
	*(*int32)(unsafe.Add(mBase, uint32(v40)+108)) = v101
	v690 = F_lappend(m, v576, v673)
	mBase = m.M
	v691 = m.ExcPending
	if v691 != 0 {
		v1547 = v107
		goto L6
	} else {
		goto L64
	}
L64:
	;
	v693 = v577 + int32(2)
	v694 = *(*int32)(unsafe.Add(mBase, uint32(v95)))
	if v693 < v694 {
		v573 = v690
		v576 = v690
		v577 = v693
		goto L58
	} else {
		goto L65
	}
L65:
	;
	goto L59
L66:
	;
	v751 = *(*int32)(unsafe.Add(mBase, uint32(v109)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v352))) = v751
	v753 = *(*int32)(unsafe.Add(mBase, uint32(v109)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v352)+4)) = v753
	*(*int32)(unsafe.Add(mBase, uint32(v40)+36)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v40)+32)) = v64
	*(*int64)(unsafe.Add(mBase, uint32(v40)+40)) = v87
	*(*int64)(unsafe.Add(mBase, uint32(v40)+48)) = v88
	*(*int64)(unsafe.Add(mBase, uint32(v40)+56)) = v89
	*(*int32)(unsafe.Add(mBase, uint32(v40)+68)) = v711
	*(*int32)(unsafe.Add(mBase, uint32(v40)+72)) = v708
	*(*int32)(unsafe.Add(mBase, uint32(v40)+76)) = v356
	*(*int32)(unsafe.Add(mBase, uint32(v40)+80)) = v352
	*(*int32)(unsafe.Add(mBase, uint32(v40)+84)) = v237
	*(*int64)(unsafe.Add(mBase, uint32(v40)+88)) = v231
	*(*int32)(unsafe.Add(mBase, uint32(v40)+96)) = v122
	*(*int32)(unsafe.Add(mBase, uint32(v40)+100)) = v107
	*(*int32)(unsafe.Add(mBase, uint32(v40)+104)) = v104
	*(*int32)(unsafe.Add(mBase, uint32(v40)+108)) = v101
	v772 = int32(*(*uint8)(unsafe.Add(mBase, _consts[112])))
	if v772 == int32(1) {
		goto L68
	} else {
		goto L69
	}
L67:
	;
	if v782 != 0 {
		goto L51
	} else {
		goto L71
	}
L68:
	;
	v777 = *(*int32)(unsafe.Add(mBase, _consts[113]))
	v778 = *(*int32)(unsafe.Add(mBase, uint32(v777)+316))
	v780 = base.B2i32(v778 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _consts[112])) = uint8(v780)
	v782 = v780
	goto L70
L69:
	;
	v782 = int32(0)
	goto L70
L70:
	;
	goto L67
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+36)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v40)+32)) = v64
	*(*int64)(unsafe.Add(mBase, uint32(v40)+40)) = v87
	*(*int64)(unsafe.Add(mBase, uint32(v40)+48)) = v88
	*(*int64)(unsafe.Add(mBase, uint32(v40)+56)) = v89
	*(*int32)(unsafe.Add(mBase, uint32(v40)+68)) = v711
	*(*int32)(unsafe.Add(mBase, uint32(v40)+72)) = v708
	*(*int32)(unsafe.Add(mBase, uint32(v40)+76)) = v356
	*(*int32)(unsafe.Add(mBase, uint32(v40)+80)) = v352
	*(*int32)(unsafe.Add(mBase, uint32(v40)+84)) = v237
	*(*int64)(unsafe.Add(mBase, uint32(v40)+88)) = v231
	*(*int32)(unsafe.Add(mBase, uint32(v40)+96)) = v122
	*(*int32)(unsafe.Add(mBase, uint32(v40)+100)) = v107
	*(*int32)(unsafe.Add(mBase, uint32(v40)+104)) = v104
	*(*int32)(unsafe.Add(mBase, uint32(v40)+108)) = v101
	v801 = int32(4411168)
	v802 = *(*int32)(unsafe.Add(mBase, _consts[113]))
	v803 = *(*int64)(unsafe.Add(mBase, uint32(v802)+280))
	*(*int64)(unsafe.Add(mBase, uint32(v802)+280)) = v803
	*(*int64)(unsafe.Add(mBase, _consts[118])) = v803
	v808 = *(*int32)(unsafe.Add(mBase, _consts[113]))
	v809 = *(*int64)(unsafe.Add(mBase, uint32(v808)+272))
	*(*int64)(unsafe.Add(mBase, uint32(v808)+272)) = v809
	*(*int64)(unsafe.Add(mBase, _consts[117])) = v809
	goto L74
L72:
	;
	v913 = v88
	v914 = v818
	v915 = v818
	goto L50
L74:
	;
	goto L75
L75:
	;
	v818 = *(*int64)(unsafe.Add(mBase, _consts[118]))
	goto L72
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+36)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v40)+32)) = v64
	*(*int64)(unsafe.Add(mBase, uint32(v40)+40)) = v87
	*(*int64)(unsafe.Add(mBase, uint32(v40)+48)) = v88
	*(*int64)(unsafe.Add(mBase, uint32(v40)+56)) = v89
	*(*int32)(unsafe.Add(mBase, uint32(v40)+68)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v40)+72)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v40)+76)) = v356
	*(*int32)(unsafe.Add(mBase, uint32(v40)+80)) = v352
	*(*int32)(unsafe.Add(mBase, uint32(v40)+84)) = v237
	*(*int64)(unsafe.Add(mBase, uint32(v40)+88)) = v231
	*(*int32)(unsafe.Add(mBase, uint32(v40)+96)) = v122
	*(*int32)(unsafe.Add(mBase, uint32(v40)+100)) = v107
	*(*int32)(unsafe.Add(mBase, uint32(v40)+104)) = v104
	*(*int32)(unsafe.Add(mBase, uint32(v40)+108)) = v101
	F_errcode(m, int32(1088))
	mBase = m.M
	v855 = m.ExcPending
	if v855 != 0 {
		v1547 = v107
		goto L6
	} else {
		goto L77
	}
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+36)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v40)+32)) = v64
	*(*int64)(unsafe.Add(mBase, uint32(v40)+40)) = v87
	*(*int64)(unsafe.Add(mBase, uint32(v40)+48)) = v88
	*(*int64)(unsafe.Add(mBase, uint32(v40)+56)) = v89
	*(*int32)(unsafe.Add(mBase, uint32(v40)+68)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v40)+72)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v40)+76)) = v356
	*(*int32)(unsafe.Add(mBase, uint32(v40)+80)) = v352
	*(*int32)(unsafe.Add(mBase, uint32(v40)+84)) = v237
	*(*int64)(unsafe.Add(mBase, uint32(v40)+88)) = v231
	*(*int32)(unsafe.Add(mBase, uint32(v40)+96)) = v122
	*(*int32)(unsafe.Add(mBase, uint32(v40)+100)) = v107
	*(*int32)(unsafe.Add(mBase, uint32(v40)+104)) = v104
	*(*int32)(unsafe.Add(mBase, uint32(v40)+108)) = v101
	F_errmsg(m, int32(123134), int32(0))
	mBase = m.M
	v874 = m.ExcPending
	if v874 != 0 {
		v1547 = v107
		goto L6
	} else {
		goto L78
	}
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+36)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v40)+32)) = v64
	*(*int64)(unsafe.Add(mBase, uint32(v40)+40)) = v87
	*(*int64)(unsafe.Add(mBase, uint32(v40)+48)) = v88
	*(*int64)(unsafe.Add(mBase, uint32(v40)+56)) = v89
	*(*int32)(unsafe.Add(mBase, uint32(v40)+68)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v40)+72)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v40)+76)) = v356
	*(*int32)(unsafe.Add(mBase, uint32(v40)+80)) = v352
	*(*int32)(unsafe.Add(mBase, uint32(v40)+84)) = v237
	*(*int64)(unsafe.Add(mBase, uint32(v40)+88)) = v231
	*(*int32)(unsafe.Add(mBase, uint32(v40)+96)) = v122
	*(*int32)(unsafe.Add(mBase, uint32(v40)+100)) = v107
	*(*int32)(unsafe.Add(mBase, uint32(v40)+104)) = v104
	*(*int32)(unsafe.Add(mBase, uint32(v40)+108)) = v101
	F_errfinish(m, int32(495112), int32(177), int32(115526))
	mBase = m.M
	v894 = m.ExcPending
	if v894 != 0 {
		v1547 = v107
		goto L6
	} else {
		goto L79
	}
L79:
	;
	goto L3
L80:
	;
	v913 = v911
	v914 = v89
	v915 = v911
	goto L50
L81:
	;
	v937 = *(*int32)(unsafe.Add(mBase, _consts[141]))
	v939 = *(*int32)(unsafe.Add(mBase, _consts[142]))
	goto L82
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v101)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v101))) = v40 + int32(28)
	goto L85
L83:
	;
	v949 = v107
	v950 = v104
	v951 = v352
	v952 = v237
	v953 = v101
	v954 = v937
	v955 = v939
	v956 = v356
	v957 = v708
	v958 = v122
	v960 = v711
	v961 = int32(0)
	v964 = v107
	v976 = v231
	v977 = v915
	v978 = v913
	v979 = v914
	goto L7
L85:
	;
	goto L83
L86:
	;
	*(*int32)(unsafe.Add(mBase, _consts[142])) = v953
	*(*int32)(unsafe.Add(mBase, uint32(v950)+8)) = int32(394)
	*(*int32)(unsafe.Add(mBase, uint32(v950)+4)) = int32(395)
	*(*int32)(unsafe.Add(mBase, uint32(v950))) = int32(396)
	*(*int32)(unsafe.Add(mBase, uint32(v40)+36)) = v955
	*(*int32)(unsafe.Add(mBase, uint32(v40)+32)) = v954
	*(*int64)(unsafe.Add(mBase, uint32(v40)+40)) = v977
	*(*int64)(unsafe.Add(mBase, uint32(v40)+48)) = v978
	*(*int64)(unsafe.Add(mBase, uint32(v40)+56)) = v979
	*(*int32)(unsafe.Add(mBase, uint32(v40)+68)) = v960
	*(*int32)(unsafe.Add(mBase, uint32(v40)+72)) = v957
	*(*int32)(unsafe.Add(mBase, uint32(v40)+76)) = v956
	*(*int32)(unsafe.Add(mBase, uint32(v40)+80)) = v951
	*(*int32)(unsafe.Add(mBase, uint32(v40)+84)) = v952
	*(*int64)(unsafe.Add(mBase, uint32(v40)+88)) = v976
	*(*int32)(unsafe.Add(mBase, uint32(v40)+96)) = v958
	*(*int32)(unsafe.Add(mBase, uint32(v40)+100)) = v949
	*(*int32)(unsafe.Add(mBase, uint32(v40)+104)) = v950
	*(*int32)(unsafe.Add(mBase, uint32(v40)+108)) = v953
	v1008 = int32(0)
	v1012 = F_CreateDecodingContext(m, int64(0), v960, v1008, v950, int32(1012), int32(1013), v1008)
	mBase = m.M
	v1013 = m.ExcPending
	if v1013 != 0 {
		v1547 = v964
		goto L6
	} else {
		goto L89
	}
L87:
	;
	goto L88
L88:
	;
	*(*int32)(unsafe.Add(mBase, _consts[141])) = v954
	*(*int32)(unsafe.Add(mBase, _consts[142])) = v955
	*(*int32)(unsafe.Add(mBase, uint32(v40)+32)) = v954
	*(*int32)(unsafe.Add(mBase, uint32(v40)+36)) = v955
	*(*int64)(unsafe.Add(mBase, uint32(v40)+40)) = v977
	*(*int64)(unsafe.Add(mBase, uint32(v40)+48)) = v978
	*(*int64)(unsafe.Add(mBase, uint32(v40)+56)) = v979
	*(*int32)(unsafe.Add(mBase, uint32(v40)+68)) = v960
	*(*int32)(unsafe.Add(mBase, uint32(v40)+72)) = v957
	*(*int32)(unsafe.Add(mBase, uint32(v40)+76)) = v956
	*(*int32)(unsafe.Add(mBase, uint32(v40)+80)) = v951
	*(*int32)(unsafe.Add(mBase, uint32(v40)+84)) = v952
	*(*int64)(unsafe.Add(mBase, uint32(v40)+88)) = v976
	*(*int32)(unsafe.Add(mBase, uint32(v40)+96)) = v958
	*(*int32)(unsafe.Add(mBase, uint32(v40)+100)) = v949
	*(*int32)(unsafe.Add(mBase, uint32(v40)+104)) = v950
	*(*int32)(unsafe.Add(mBase, uint32(v40)+108)) = v953
	F_InvalidateSystemCachesExtended(m)
	mBase = m.M
	v1510 = m.ExcPending
	if v1510 != 0 {
		v1547 = v964
		goto L6
	} else {
		goto L142
	}
L89:
	;
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v956
	if v3 != 0 {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+36)) = v955
	*(*int32)(unsafe.Add(mBase, uint32(v40)+32)) = v954
	*(*int64)(unsafe.Add(mBase, uint32(v40)+48)) = v978
	*(*int64)(unsafe.Add(mBase, uint32(v40)+56)) = v979
	*(*int32)(unsafe.Add(mBase, uint32(v40)+68)) = v960
	*(*int32)(unsafe.Add(mBase, uint32(v40)+72)) = v957
	*(*int32)(unsafe.Add(mBase, uint32(v40)+76)) = v956
	*(*int32)(unsafe.Add(mBase, uint32(v40)+80)) = v951
	*(*int32)(unsafe.Add(mBase, uint32(v40)+84)) = v952
	*(*int32)(unsafe.Add(mBase, uint32(v40)+96)) = v958
	*(*int32)(unsafe.Add(mBase, uint32(v40)+100)) = v949
	*(*int32)(unsafe.Add(mBase, uint32(v40)+104)) = v950
	*(*int32)(unsafe.Add(mBase, uint32(v40)+108)) = v953
	*(*int64)(unsafe.Add(mBase, uint32(v40)+88)) = v976
	*(*int64)(unsafe.Add(mBase, uint32(v40)+40)) = v977
	if base.Ui64(v976) < base.Ui64(v977) {
		goto L98
	} else {
		goto L99
	}
L91:
	;
	v1016 = *(*int32)(unsafe.Add(mBase, uint32(v1012)+108))
	if v1016 == int32(1) {
		goto L90
	} else {
		goto L92
	}
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+36)) = v955
	*(*int32)(unsafe.Add(mBase, uint32(v40)+32)) = v954
	*(*int64)(unsafe.Add(mBase, uint32(v40)+40)) = v977
	*(*int64)(unsafe.Add(mBase, uint32(v40)+48)) = v978
	*(*int64)(unsafe.Add(mBase, uint32(v40)+56)) = v979
	*(*int32)(unsafe.Add(mBase, uint32(v40)+68)) = v960
	*(*int32)(unsafe.Add(mBase, uint32(v40)+72)) = v957
	*(*int32)(unsafe.Add(mBase, uint32(v40)+76)) = v956
	*(*int32)(unsafe.Add(mBase, uint32(v40)+80)) = v951
	*(*int32)(unsafe.Add(mBase, uint32(v40)+84)) = v952
	*(*int64)(unsafe.Add(mBase, uint32(v40)+88)) = v976
	*(*int32)(unsafe.Add(mBase, uint32(v40)+96)) = v958
	*(*int32)(unsafe.Add(mBase, uint32(v40)+100)) = v949
	*(*int32)(unsafe.Add(mBase, uint32(v40)+104)) = v950
	*(*int32)(unsafe.Add(mBase, uint32(v40)+108)) = v953
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1037 = m.ExcPending
	if v1037 != 0 {
		v1547 = v964
		goto L6
	} else {
		goto L93
	}
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+36)) = v955
	*(*int32)(unsafe.Add(mBase, uint32(v40)+32)) = v954
	*(*int64)(unsafe.Add(mBase, uint32(v40)+40)) = v977
	*(*int64)(unsafe.Add(mBase, uint32(v40)+48)) = v978
	*(*int64)(unsafe.Add(mBase, uint32(v40)+56)) = v979
	*(*int32)(unsafe.Add(mBase, uint32(v40)+68)) = v960
	*(*int32)(unsafe.Add(mBase, uint32(v40)+72)) = v957
	*(*int32)(unsafe.Add(mBase, uint32(v40)+76)) = v956
	*(*int32)(unsafe.Add(mBase, uint32(v40)+80)) = v951
	*(*int32)(unsafe.Add(mBase, uint32(v40)+84)) = v952
	*(*int64)(unsafe.Add(mBase, uint32(v40)+88)) = v976
	*(*int32)(unsafe.Add(mBase, uint32(v40)+96)) = v958
	*(*int32)(unsafe.Add(mBase, uint32(v40)+100)) = v949
	*(*int32)(unsafe.Add(mBase, uint32(v40)+104)) = v950
	*(*int32)(unsafe.Add(mBase, uint32(v40)+108)) = v953
	F_errcode(m, int32(1088))
	mBase = m.M
	v1055 = m.ExcPending
	if v1055 != 0 {
		v1547 = v964
		goto L6
	} else {
		goto L94
	}
L94:
	;
	v1056 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1057 = *(*int32)(unsafe.Add(mBase, uint32(v1056)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v40)+32)) = v954
	*(*int32)(unsafe.Add(mBase, uint32(v40)+36)) = v955
	*(*int64)(unsafe.Add(mBase, uint32(v40)+40)) = v977
	*(*int64)(unsafe.Add(mBase, uint32(v40)+48)) = v978
	*(*int64)(unsafe.Add(mBase, uint32(v40)+56)) = v979
	*(*int32)(unsafe.Add(mBase, uint32(v40)+68)) = v960
	*(*int32)(unsafe.Add(mBase, uint32(v40)+72)) = v957
	*(*int32)(unsafe.Add(mBase, uint32(v40)+76)) = v956
	*(*int32)(unsafe.Add(mBase, uint32(v40)+80)) = v951
	*(*int32)(unsafe.Add(mBase, uint32(v40)+84)) = v952
	*(*int64)(unsafe.Add(mBase, uint32(v40)+88)) = v976
	*(*int32)(unsafe.Add(mBase, uint32(v40)+96)) = v958
	*(*int32)(unsafe.Add(mBase, uint32(v40)+100)) = v949
	*(*int32)(unsafe.Add(mBase, uint32(v40)+104)) = v950
	*(*int32)(unsafe.Add(mBase, uint32(v40)+108)) = v953
	v1074 = *(*int32)(unsafe.Add(mBase, _consts[517]))
	v1075 = F_format_procedure(m, v1057)
	mBase = m.M
	v1076 = m.ExcPending
	if v1076 != 0 {
		v1547 = v964
		goto L6
	} else {
		goto L95
	}
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+36)) = v955
	*(*int32)(unsafe.Add(mBase, uint32(v40)+32)) = v954
	*(*int64)(unsafe.Add(mBase, uint32(v40)+40)) = v977
	*(*int64)(unsafe.Add(mBase, uint32(v40)+48)) = v978
	*(*int64)(unsafe.Add(mBase, uint32(v40)+56)) = v979
	*(*int32)(unsafe.Add(mBase, uint32(v40)+68)) = v960
	*(*int32)(unsafe.Add(mBase, uint32(v40)+72)) = v957
	*(*int32)(unsafe.Add(mBase, uint32(v40)+76)) = v956
	*(*int32)(unsafe.Add(mBase, uint32(v40)+80)) = v951
	*(*int32)(unsafe.Add(mBase, uint32(v40)+84)) = v952
	*(*int64)(unsafe.Add(mBase, uint32(v40)+88)) = v976
	*(*int32)(unsafe.Add(mBase, uint32(v40)+96)) = v958
	*(*int32)(unsafe.Add(mBase, uint32(v40)+100)) = v949
	*(*int32)(unsafe.Add(mBase, uint32(v40)+104)) = v950
	*(*int32)(unsafe.Add(mBase, uint32(v40)+108)) = v953
	*(*int32)(unsafe.Add(mBase, uint32(v40)+20)) = v1075
	*(*int32)(unsafe.Add(mBase, uint32(v40)+16)) = v1074 + int32(137)
	F_errmsg(m, int32(505098), v40+int32(16))
	mBase = m.M
	v1100 = m.ExcPending
	if v1100 != 0 {
		v1547 = v964
		goto L6
	} else {
		goto L96
	}
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+36)) = v955
	*(*int32)(unsafe.Add(mBase, uint32(v40)+32)) = v954
	*(*int64)(unsafe.Add(mBase, uint32(v40)+40)) = v977
	*(*int64)(unsafe.Add(mBase, uint32(v40)+48)) = v978
	*(*int64)(unsafe.Add(mBase, uint32(v40)+56)) = v979
	*(*int32)(unsafe.Add(mBase, uint32(v40)+68)) = v960
	*(*int32)(unsafe.Add(mBase, uint32(v40)+72)) = v957
	*(*int32)(unsafe.Add(mBase, uint32(v40)+76)) = v956
	*(*int32)(unsafe.Add(mBase, uint32(v40)+80)) = v951
	*(*int32)(unsafe.Add(mBase, uint32(v40)+84)) = v952
	*(*int64)(unsafe.Add(mBase, uint32(v40)+88)) = v976
	*(*int32)(unsafe.Add(mBase, uint32(v40)+96)) = v958
	*(*int32)(unsafe.Add(mBase, uint32(v40)+100)) = v949
	*(*int32)(unsafe.Add(mBase, uint32(v40)+104)) = v950
	*(*int32)(unsafe.Add(mBase, uint32(v40)+108)) = v953
	F_errfinish(m, int32(495112), int32(226), int32(115526))
	mBase = m.M
	v1120 = m.ExcPending
	if v1120 != 0 {
		v1547 = v964
		goto L6
	} else {
		goto L97
	}
L97:
	;
	goto L3
L98:
	;
	v1137 = v976
	goto L100
L99:
	;
	v1137 = v977
	goto L100
L100:
	;
	v1139 = base.B2i32(v976 == int64(0))
	if v976 == int64(0) {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	v1140 = v977
	goto L103
L102:
	;
	v1140 = v1137
	goto L103
L103:
	;
	F_WaitForStandbyConfirmation(m, v1140)
	mBase = m.M
	v1142 = m.ExcPending
	if v1142 != 0 {
		v1547 = v964
		goto L6
	} else {
		goto L104
	}
L104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1012)+140)) = v951
	v1144 = *(*int32)(unsafe.Add(mBase, uint32(v1012)+8))
	v1146 = *(*int32)(unsafe.Add(mBase, _consts[517]))
	v1147 = *(*int64)(unsafe.Add(mBase, uint32(v1146)+104))
	*(*int32)(unsafe.Add(mBase, uint32(v40)+32)) = v954
	*(*int32)(unsafe.Add(mBase, uint32(v40)+36)) = v955
	*(*int64)(unsafe.Add(mBase, uint32(v40)+40)) = v977
	*(*int64)(unsafe.Add(mBase, uint32(v40)+48)) = v978
	*(*int64)(unsafe.Add(mBase, uint32(v40)+56)) = v979
	*(*int32)(unsafe.Add(mBase, uint32(v40)+68)) = v960
	*(*int32)(unsafe.Add(mBase, uint32(v40)+72)) = v957
	*(*int32)(unsafe.Add(mBase, uint32(v40)+76)) = v956
	*(*int32)(unsafe.Add(mBase, uint32(v40)+80)) = v951
	*(*int32)(unsafe.Add(mBase, uint32(v40)+84)) = v952
	*(*int64)(unsafe.Add(mBase, uint32(v40)+88)) = v976
	*(*int32)(unsafe.Add(mBase, uint32(v40)+96)) = v958
	*(*int32)(unsafe.Add(mBase, uint32(v40)+100)) = v949
	*(*int32)(unsafe.Add(mBase, uint32(v40)+104)) = v950
	*(*int32)(unsafe.Add(mBase, uint32(v40)+108)) = v953
	F_XLogBeginRead(m, v1144, v1147)
	mBase = m.M
	v1164 = m.ExcPending
	if v1164 != 0 {
		v1547 = v964
		goto L6
	} else {
		goto L105
	}
L105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+36)) = v955
	*(*int32)(unsafe.Add(mBase, uint32(v40)+32)) = v954
	*(*int64)(unsafe.Add(mBase, uint32(v40)+40)) = v977
	*(*int64)(unsafe.Add(mBase, uint32(v40)+48)) = v978
	*(*int64)(unsafe.Add(mBase, uint32(v40)+56)) = v979
	*(*int32)(unsafe.Add(mBase, uint32(v40)+68)) = v960
	*(*int32)(unsafe.Add(mBase, uint32(v40)+72)) = v957
	*(*int32)(unsafe.Add(mBase, uint32(v40)+76)) = v956
	*(*int32)(unsafe.Add(mBase, uint32(v40)+80)) = v951
	*(*int32)(unsafe.Add(mBase, uint32(v40)+84)) = v952
	*(*int64)(unsafe.Add(mBase, uint32(v40)+88)) = v976
	*(*int32)(unsafe.Add(mBase, uint32(v40)+96)) = v958
	*(*int32)(unsafe.Add(mBase, uint32(v40)+100)) = v949
	*(*int32)(unsafe.Add(mBase, uint32(v40)+104)) = v950
	*(*int32)(unsafe.Add(mBase, uint32(v40)+108)) = v953
	F_InvalidateSystemCachesExtended(m)
	mBase = m.M
	v1181 = m.ExcPending
	if v1181 != 0 {
		v1547 = v964
		goto L6
	} else {
		goto L106
	}
L106:
	;
	v1182 = *(*int32)(unsafe.Add(mBase, uint32(v1012)+8))
	v1183 = *(*int64)(unsafe.Add(mBase, uint32(v1182)+40))
	if base.Ui64(v977) <= base.Ui64(v1183) {
		v1390 = v1183
		goto L107
	} else {
		goto L108
	}
L107:
	;
	*(*int32)(unsafe.Add(mBase, _consts[10])) = v958
	if l1 == int32(0) {
		goto L134
	} else {
		goto L135
	}
L108:
	;
	v1202 = v1182
	goto L109
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v949))) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v40)+36)) = v955
	*(*int32)(unsafe.Add(mBase, uint32(v40)+32)) = v954
	*(*int64)(unsafe.Add(mBase, uint32(v40)+40)) = v977
	*(*int64)(unsafe.Add(mBase, uint32(v40)+48)) = v978
	*(*int64)(unsafe.Add(mBase, uint32(v40)+56)) = v979
	*(*int32)(unsafe.Add(mBase, uint32(v40)+68)) = v960
	*(*int32)(unsafe.Add(mBase, uint32(v40)+72)) = v957
	*(*int32)(unsafe.Add(mBase, uint32(v40)+76)) = v956
	*(*int32)(unsafe.Add(mBase, uint32(v40)+80)) = v951
	*(*int32)(unsafe.Add(mBase, uint32(v40)+84)) = v952
	*(*int64)(unsafe.Add(mBase, uint32(v40)+88)) = v976
	*(*int32)(unsafe.Add(mBase, uint32(v40)+96)) = v958
	*(*int32)(unsafe.Add(mBase, uint32(v40)+100)) = v949
	*(*int32)(unsafe.Add(mBase, uint32(v40)+104)) = v950
	*(*int32)(unsafe.Add(mBase, uint32(v40)+108)) = v953
	v1240 = F_XLogReadRecord(m, v1202, v949)
	mBase = m.M
	v1241 = m.ExcPending
	if v1241 != 0 {
		v1547 = v964
		goto L6
	} else {
		goto L111
	}
L110:
	;
	v1390 = v1353
	goto L107
L111:
	;
	v1242 = *(*int32)(unsafe.Add(mBase, uint32(v949)))
	if v1242 != 0 {
		goto L112
	} else {
		goto L113
	}
L112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+36)) = v955
	*(*int32)(unsafe.Add(mBase, uint32(v40)+32)) = v954
	*(*int64)(unsafe.Add(mBase, uint32(v40)+40)) = v977
	*(*int64)(unsafe.Add(mBase, uint32(v40)+48)) = v978
	*(*int64)(unsafe.Add(mBase, uint32(v40)+56)) = v979
	*(*int32)(unsafe.Add(mBase, uint32(v40)+68)) = v960
	*(*int32)(unsafe.Add(mBase, uint32(v40)+72)) = v957
	*(*int32)(unsafe.Add(mBase, uint32(v40)+76)) = v956
	*(*int32)(unsafe.Add(mBase, uint32(v40)+80)) = v951
	*(*int32)(unsafe.Add(mBase, uint32(v40)+84)) = v952
	*(*int64)(unsafe.Add(mBase, uint32(v40)+88)) = v976
	*(*int32)(unsafe.Add(mBase, uint32(v40)+96)) = v958
	*(*int32)(unsafe.Add(mBase, uint32(v40)+100)) = v949
	*(*int32)(unsafe.Add(mBase, uint32(v40)+104)) = v950
	*(*int32)(unsafe.Add(mBase, uint32(v40)+108)) = v953
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1261 = m.ExcPending
	if v1261 != 0 {
		v1547 = v964
		goto L6
	} else {
		goto L115
	}
L113:
	;
	goto L114
L114:
	;
	if v1240 != 0 {
		goto L118
	} else {
		goto L119
	}
L115:
	;
	v1262 = *(*int32)(unsafe.Add(mBase, uint32(v949)))
	*(*int32)(unsafe.Add(mBase, uint32(v40)+36)) = v955
	*(*int32)(unsafe.Add(mBase, uint32(v40)+32)) = v954
	*(*int64)(unsafe.Add(mBase, uint32(v40)+40)) = v977
	*(*int64)(unsafe.Add(mBase, uint32(v40)+48)) = v978
	*(*int64)(unsafe.Add(mBase, uint32(v40)+56)) = v979
	*(*int32)(unsafe.Add(mBase, uint32(v40)+68)) = v960
	*(*int32)(unsafe.Add(mBase, uint32(v40)+72)) = v957
	*(*int32)(unsafe.Add(mBase, uint32(v40)+76)) = v956
	*(*int32)(unsafe.Add(mBase, uint32(v40)+80)) = v951
	*(*int32)(unsafe.Add(mBase, uint32(v40)+84)) = v952
	*(*int64)(unsafe.Add(mBase, uint32(v40)+88)) = v976
	*(*int32)(unsafe.Add(mBase, uint32(v40)+96)) = v958
	*(*int32)(unsafe.Add(mBase, uint32(v40)+100)) = v949
	*(*int32)(unsafe.Add(mBase, uint32(v40)+104)) = v950
	*(*int32)(unsafe.Add(mBase, uint32(v40)+108)) = v953
	*(*int32)(unsafe.Add(mBase, uint32(v40))) = v1262
	F_errmsg_internal(m, int32(202944), v40)
	mBase = m.M
	v1281 = m.ExcPending
	if v1281 != 0 {
		v1547 = v964
		goto L6
	} else {
		goto L116
	}
L116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+36)) = v955
	*(*int32)(unsafe.Add(mBase, uint32(v40)+32)) = v954
	*(*int64)(unsafe.Add(mBase, uint32(v40)+40)) = v977
	*(*int64)(unsafe.Add(mBase, uint32(v40)+48)) = v978
	*(*int64)(unsafe.Add(mBase, uint32(v40)+56)) = v979
	*(*int32)(unsafe.Add(mBase, uint32(v40)+68)) = v960
	*(*int32)(unsafe.Add(mBase, uint32(v40)+72)) = v957
	*(*int32)(unsafe.Add(mBase, uint32(v40)+76)) = v956
	*(*int32)(unsafe.Add(mBase, uint32(v40)+80)) = v951
	*(*int32)(unsafe.Add(mBase, uint32(v40)+84)) = v952
	*(*int64)(unsafe.Add(mBase, uint32(v40)+88)) = v976
	*(*int32)(unsafe.Add(mBase, uint32(v40)+96)) = v958
	*(*int32)(unsafe.Add(mBase, uint32(v40)+100)) = v949
	*(*int32)(unsafe.Add(mBase, uint32(v40)+104)) = v950
	*(*int32)(unsafe.Add(mBase, uint32(v40)+108)) = v953
	F_errfinish(m, int32(495112), int32(259), int32(115526))
	mBase = m.M
	v1301 = m.ExcPending
	if v1301 != 0 {
		v1547 = v964
		goto L6
	} else {
		goto L117
	}
L117:
	;
	goto L3
L118:
	;
	v1302 = *(*int32)(unsafe.Add(mBase, uint32(v1012)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v40)+36)) = v955
	*(*int32)(unsafe.Add(mBase, uint32(v40)+32)) = v954
	*(*int64)(unsafe.Add(mBase, uint32(v40)+40)) = v977
	*(*int64)(unsafe.Add(mBase, uint32(v40)+48)) = v978
	*(*int64)(unsafe.Add(mBase, uint32(v40)+56)) = v979
	*(*int32)(unsafe.Add(mBase, uint32(v40)+68)) = v960
	*(*int32)(unsafe.Add(mBase, uint32(v40)+72)) = v957
	*(*int32)(unsafe.Add(mBase, uint32(v40)+76)) = v956
	*(*int32)(unsafe.Add(mBase, uint32(v40)+80)) = v951
	*(*int32)(unsafe.Add(mBase, uint32(v40)+84)) = v952
	*(*int64)(unsafe.Add(mBase, uint32(v40)+88)) = v976
	*(*int32)(unsafe.Add(mBase, uint32(v40)+96)) = v958
	*(*int32)(unsafe.Add(mBase, uint32(v40)+100)) = v949
	*(*int32)(unsafe.Add(mBase, uint32(v40)+104)) = v950
	*(*int32)(unsafe.Add(mBase, uint32(v40)+108)) = v953
	F_LogicalDecodingProcessRecord(m, v1012, v1302)
	mBase = m.M
	v1319 = m.ExcPending
	if v1319 != 0 {
		v1547 = v964
		goto L6
	} else {
		goto L121
	}
L119:
	;
	goto L120
L120:
	;
	if v1139 == int32(0) {
		goto L122
	} else {
		goto L123
	}
L121:
	;
	goto L120
L122:
	;
	v1323 = *(*int32)(unsafe.Add(mBase, uint32(v1012)+8))
	v1324 = *(*int64)(unsafe.Add(mBase, uint32(v1323)+40))
	if base.Ui64(v976) <= base.Ui64(v1324) {
		v1390 = v1324
		goto L107
	} else {
		goto L125
	}
L123:
	;
	goto L124
L124:
	;
	if v952 == int32(0) {
		goto L126
	} else {
		goto L127
	}
L125:
	;
	goto L124
L126:
	;
	v1334 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	if v1334 != 0 {
		goto L129
	} else {
		goto L130
	}
L127:
	;
	v1329 = *(*int64)(unsafe.Add(mBase, uint32(v951)+16))
	if v1329 < base.I64_extend_i32_s(v952) {
		goto L126
	} else {
		goto L128
	}
L128:
	;
	v1331 = *(*int32)(unsafe.Add(mBase, uint32(v1012)+8))
	v1332 = *(*int64)(unsafe.Add(mBase, uint32(v1331)+40))
	v1390 = v1332
	goto L107
L129:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+36)) = v955
	*(*int32)(unsafe.Add(mBase, uint32(v40)+32)) = v954
	*(*int64)(unsafe.Add(mBase, uint32(v40)+40)) = v977
	*(*int64)(unsafe.Add(mBase, uint32(v40)+48)) = v978
	*(*int64)(unsafe.Add(mBase, uint32(v40)+56)) = v979
	*(*int32)(unsafe.Add(mBase, uint32(v40)+68)) = v960
	*(*int32)(unsafe.Add(mBase, uint32(v40)+72)) = v957
	*(*int32)(unsafe.Add(mBase, uint32(v40)+76)) = v956
	*(*int32)(unsafe.Add(mBase, uint32(v40)+80)) = v951
	*(*int32)(unsafe.Add(mBase, uint32(v40)+84)) = v952
	*(*int64)(unsafe.Add(mBase, uint32(v40)+88)) = v976
	*(*int32)(unsafe.Add(mBase, uint32(v40)+96)) = v958
	*(*int32)(unsafe.Add(mBase, uint32(v40)+100)) = v949
	*(*int32)(unsafe.Add(mBase, uint32(v40)+104)) = v950
	*(*int32)(unsafe.Add(mBase, uint32(v40)+108)) = v953
	F_ProcessInterrupts(m)
	mBase = m.M
	v1351 = m.ExcPending
	if v1351 != 0 {
		v1547 = v964
		goto L6
	} else {
		goto L132
	}
L130:
	;
	goto L131
L131:
	;
	v1352 = *(*int32)(unsafe.Add(mBase, uint32(v1012)+8))
	v1353 = *(*int64)(unsafe.Add(mBase, uint32(v1352)+40))
	if base.Ui64(v1353) < base.Ui64(v977) {
		v1202 = v1352
		goto L109
	} else {
		goto L133
	}
L132:
	;
	goto L131
L133:
	;
	goto L110
L134:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+36)) = v955
	*(*int32)(unsafe.Add(mBase, uint32(v40)+32)) = v954
	*(*int64)(unsafe.Add(mBase, uint32(v40)+40)) = v977
	*(*int64)(unsafe.Add(mBase, uint32(v40)+48)) = v978
	*(*int64)(unsafe.Add(mBase, uint32(v40)+56)) = v979
	*(*int32)(unsafe.Add(mBase, uint32(v40)+68)) = v960
	*(*int32)(unsafe.Add(mBase, uint32(v40)+72)) = v957
	*(*int32)(unsafe.Add(mBase, uint32(v40)+76)) = v956
	*(*int32)(unsafe.Add(mBase, uint32(v40)+80)) = v951
	*(*int32)(unsafe.Add(mBase, uint32(v40)+84)) = v952
	*(*int64)(unsafe.Add(mBase, uint32(v40)+88)) = v976
	*(*int32)(unsafe.Add(mBase, uint32(v40)+96)) = v958
	*(*int32)(unsafe.Add(mBase, uint32(v40)+100)) = v949
	*(*int32)(unsafe.Add(mBase, uint32(v40)+104)) = v950
	*(*int32)(unsafe.Add(mBase, uint32(v40)+108)) = v953
	F_FreeDecodingContext(m, v1012)
	mBase = m.M
	v1448 = m.ExcPending
	if v1448 != 0 {
		v1547 = v964
		goto L6
	} else {
		goto L139
	}
L135:
	;
	if v1390 == int64(0) {
		goto L134
	} else {
		goto L136
	}
L136:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+36)) = v955
	*(*int32)(unsafe.Add(mBase, uint32(v40)+32)) = v954
	*(*int64)(unsafe.Add(mBase, uint32(v40)+40)) = v977
	*(*int64)(unsafe.Add(mBase, uint32(v40)+48)) = v978
	*(*int64)(unsafe.Add(mBase, uint32(v40)+56)) = v979
	*(*int32)(unsafe.Add(mBase, uint32(v40)+68)) = v960
	*(*int32)(unsafe.Add(mBase, uint32(v40)+72)) = v957
	*(*int32)(unsafe.Add(mBase, uint32(v40)+76)) = v956
	*(*int32)(unsafe.Add(mBase, uint32(v40)+80)) = v951
	*(*int32)(unsafe.Add(mBase, uint32(v40)+84)) = v952
	*(*int64)(unsafe.Add(mBase, uint32(v40)+88)) = v976
	*(*int32)(unsafe.Add(mBase, uint32(v40)+96)) = v958
	*(*int32)(unsafe.Add(mBase, uint32(v40)+100)) = v949
	*(*int32)(unsafe.Add(mBase, uint32(v40)+104)) = v950
	*(*int32)(unsafe.Add(mBase, uint32(v40)+108)) = v953
	F_LogicalConfirmReceivedLocation(m, v1390)
	mBase = m.M
	v1414 = m.ExcPending
	if v1414 != 0 {
		v1547 = v964
		goto L6
	} else {
		goto L137
	}
L137:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+36)) = v955
	*(*int32)(unsafe.Add(mBase, uint32(v40)+32)) = v954
	*(*int64)(unsafe.Add(mBase, uint32(v40)+40)) = v977
	*(*int64)(unsafe.Add(mBase, uint32(v40)+48)) = v978
	*(*int64)(unsafe.Add(mBase, uint32(v40)+56)) = v979
	*(*int32)(unsafe.Add(mBase, uint32(v40)+68)) = v960
	*(*int32)(unsafe.Add(mBase, uint32(v40)+72)) = v957
	*(*int32)(unsafe.Add(mBase, uint32(v40)+76)) = v956
	*(*int32)(unsafe.Add(mBase, uint32(v40)+80)) = v951
	*(*int32)(unsafe.Add(mBase, uint32(v40)+84)) = v952
	*(*int64)(unsafe.Add(mBase, uint32(v40)+88)) = v976
	*(*int32)(unsafe.Add(mBase, uint32(v40)+96)) = v958
	*(*int32)(unsafe.Add(mBase, uint32(v40)+100)) = v949
	*(*int32)(unsafe.Add(mBase, uint32(v40)+104)) = v950
	*(*int32)(unsafe.Add(mBase, uint32(v40)+108)) = v953
	F_ReplicationSlotMarkDirty(m)
	mBase = m.M
	v1431 = m.ExcPending
	if v1431 != 0 {
		v1547 = v964
		goto L6
	} else {
		goto L138
	}
L138:
	;
	goto L134
L139:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+36)) = v955
	*(*int32)(unsafe.Add(mBase, uint32(v40)+32)) = v954
	*(*int64)(unsafe.Add(mBase, uint32(v40)+40)) = v977
	*(*int64)(unsafe.Add(mBase, uint32(v40)+48)) = v978
	*(*int64)(unsafe.Add(mBase, uint32(v40)+56)) = v979
	*(*int32)(unsafe.Add(mBase, uint32(v40)+68)) = v960
	*(*int32)(unsafe.Add(mBase, uint32(v40)+72)) = v957
	*(*int32)(unsafe.Add(mBase, uint32(v40)+76)) = v956
	*(*int32)(unsafe.Add(mBase, uint32(v40)+80)) = v951
	*(*int32)(unsafe.Add(mBase, uint32(v40)+84)) = v952
	*(*int64)(unsafe.Add(mBase, uint32(v40)+88)) = v976
	*(*int32)(unsafe.Add(mBase, uint32(v40)+96)) = v958
	*(*int32)(unsafe.Add(mBase, uint32(v40)+100)) = v949
	*(*int32)(unsafe.Add(mBase, uint32(v40)+104)) = v950
	*(*int32)(unsafe.Add(mBase, uint32(v40)+108)) = v953
	F_ReplicationSlotRelease(m)
	mBase = m.M
	v1465 = m.ExcPending
	if v1465 != 0 {
		v1547 = v964
		goto L6
	} else {
		goto L140
	}
L140:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+36)) = v955
	*(*int32)(unsafe.Add(mBase, uint32(v40)+32)) = v954
	*(*int64)(unsafe.Add(mBase, uint32(v40)+40)) = v977
	*(*int64)(unsafe.Add(mBase, uint32(v40)+48)) = v978
	*(*int64)(unsafe.Add(mBase, uint32(v40)+56)) = v979
	*(*int32)(unsafe.Add(mBase, uint32(v40)+68)) = v960
	*(*int32)(unsafe.Add(mBase, uint32(v40)+72)) = v957
	*(*int32)(unsafe.Add(mBase, uint32(v40)+76)) = v956
	*(*int32)(unsafe.Add(mBase, uint32(v40)+80)) = v951
	*(*int32)(unsafe.Add(mBase, uint32(v40)+84)) = v952
	*(*int64)(unsafe.Add(mBase, uint32(v40)+88)) = v976
	*(*int32)(unsafe.Add(mBase, uint32(v40)+96)) = v958
	*(*int32)(unsafe.Add(mBase, uint32(v40)+100)) = v949
	*(*int32)(unsafe.Add(mBase, uint32(v40)+104)) = v950
	*(*int32)(unsafe.Add(mBase, uint32(v40)+108)) = v953
	F_InvalidateSystemCachesExtended(m)
	mBase = m.M
	v1482 = m.ExcPending
	if v1482 != 0 {
		v1547 = v964
		goto L6
	} else {
		goto L141
	}
L141:
	;
	*(*int32)(unsafe.Add(mBase, _consts[141])) = v954
	*(*int32)(unsafe.Add(mBase, _consts[142])) = v955
	m.G0 = v40 + int32(112)
	return
L142:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+36)) = v955
	*(*int32)(unsafe.Add(mBase, uint32(v40)+32)) = v954
	*(*int64)(unsafe.Add(mBase, uint32(v40)+40)) = v977
	*(*int64)(unsafe.Add(mBase, uint32(v40)+48)) = v978
	*(*int64)(unsafe.Add(mBase, uint32(v40)+56)) = v979
	*(*int32)(unsafe.Add(mBase, uint32(v40)+68)) = v960
	*(*int32)(unsafe.Add(mBase, uint32(v40)+72)) = v957
	*(*int32)(unsafe.Add(mBase, uint32(v40)+76)) = v956
	*(*int32)(unsafe.Add(mBase, uint32(v40)+80)) = v951
	*(*int32)(unsafe.Add(mBase, uint32(v40)+84)) = v952
	*(*int64)(unsafe.Add(mBase, uint32(v40)+88)) = v976
	*(*int32)(unsafe.Add(mBase, uint32(v40)+96)) = v958
	*(*int32)(unsafe.Add(mBase, uint32(v40)+100)) = v949
	*(*int32)(unsafe.Add(mBase, uint32(v40)+104)) = v950
	*(*int32)(unsafe.Add(mBase, uint32(v40)+108)) = v953
	F_pg_re_throw(m)
	mBase = m.M
	v1527 = m.ExcPending
	if v1527 != 0 {
		v1547 = v964
		goto L6
	} else {
		goto L143
	}
L143:
	;
	goto L5
L144:
	;
	v1570 = int32(v1566)
	m.G0 = v1547
	v1572 = *(*int32)(unsafe.Add(mBase, uint32(v1570)+4))
	v1573 = *(*int32)(unsafe.Add(mBase, uint32(v1570)))
	v1577 = *(*int32)(unsafe.Add(mBase, uint32(v1573)))
	if v40+int32(28) == v1577 {
		goto L147
	} else {
		goto L148
	}
L145:
	;
	m.ExcPending = 1
	goto L153
L146:
	;
	if v1580 != 0 {
		goto L150
	} else {
		goto L151
	}
L147:
	;
	v1579 = *(*int32)(unsafe.Add(mBase, uint32(v1573)+4))
	v1580 = v1579
	goto L149
L148:
	;
	v1580 = int32(0)
	goto L149
L149:
	;
	goto L146
L150:
	;
	v1581 = *(*int32)(unsafe.Add(mBase, uint32(v40)+108))
	v1582 = *(*int32)(unsafe.Add(mBase, uint32(v40)+104))
	v1583 = *(*int32)(unsafe.Add(mBase, uint32(v40)+100))
	v1584 = *(*int32)(unsafe.Add(mBase, uint32(v40)+96))
	v1585 = *(*int64)(unsafe.Add(mBase, uint32(v40)+88))
	v1586 = *(*int32)(unsafe.Add(mBase, uint32(v40)+84))
	v1587 = *(*int32)(unsafe.Add(mBase, uint32(v40)+80))
	v1588 = *(*int32)(unsafe.Add(mBase, uint32(v40)+76))
	v1589 = *(*int32)(unsafe.Add(mBase, uint32(v40)+72))
	v1590 = *(*int32)(unsafe.Add(mBase, uint32(v40)+68))
	v1591 = *(*int64)(unsafe.Add(mBase, uint32(v40)+56))
	v1592 = *(*int64)(unsafe.Add(mBase, uint32(v40)+48))
	v1593 = *(*int64)(unsafe.Add(mBase, uint32(v40)+40))
	v1594 = *(*int32)(unsafe.Add(mBase, uint32(v40)+36))
	v1595 = *(*int32)(unsafe.Add(mBase, uint32(v40)+32))
	v59 = v1583
	v60 = v1582
	v61 = v1587
	v62 = v1586
	v63 = v1581
	v64 = v1595
	v65 = v1594
	v66 = v1588
	v67 = v1589
	v68 = v1584
	v69 = v1590
	v70 = v1580
	v71 = v1572
	v74 = v1547
	v86 = v1585
	v87 = v1593
	v88 = v1592
	v89 = v1591
	goto L1
L151:
	;
	goto L152
L152:
	;
	F___wasm_longjmp(m, v1573, v1572)
	mBase = m.M
	v1597 = m.ExcPending
	if v1597 != 0 {
		goto L153
	} else {
		goto L154
	}
L153:
	;
	return
L154:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
