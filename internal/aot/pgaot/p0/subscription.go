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
	F_LockSharedObject(m, int32(_a_F_AddSubscriptionRelState_0), l0, int32(1))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return
	} else {
		v18 = F_table_open(m, int32(_a_F_AddSubscriptionRelState_1), int32(3))
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
											F_relation_close(m, v18, int32(0))
											mBase = m.M
											v54 = m.ExcPending
											if v54 != 0 {
												return
											} else {
												m.G0 = v10 + int32(48)
												return
											}
										} else {
											F_relation_close(m, v18, int32(3))
											mBase = m.M
											v57 = m.ExcPending
											if v57 != 0 {
												return
											} else {
												F_UnlockSharedObject(m, int32(_a_F_AddSubscriptionRelState_0), l0, int32(1))
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
										F_relation_close(m, v18, int32(0))
										mBase = m.M
										v54 = m.ExcPending
										if v54 != 0 {
											return
										} else {
											m.G0 = v10 + int32(48)
											return
										}
									} else {
										F_relation_close(m, v18, int32(3))
										mBase = m.M
										v57 = m.ExcPending
										if v57 != 0 {
											return
										} else {
											F_UnlockSharedObject(m, int32(_a_F_AddSubscriptionRelState_0), l0, int32(1))
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
						F_errmsg_internal(m, int32(_a_F_AddSubscriptionRelState_2), v10)
						mBase = m.M
						v73 = m.ExcPending
						if v73 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_AddSubscriptionRelState_3), int32(285), int32(_a_F_AddSubscriptionRelState_4))
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
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
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
	var v34 int32
	_ = v34
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
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v115 int32
	_ = v115
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v145 int32
	_ = v145
	var v160 int32
	_ = v160
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v191 int32
	_ = v191
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v240 int32
	_ = v240
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v302 int32
	_ = v302
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v404 int32
	_ = v404
	var v422 int32
	_ = v422
	var v425 int32
	_ = v425
	var v439 int32
	_ = v439
	var v451 int32
	_ = v451
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v483 int32
	_ = v483
	var v492 int32
	_ = v492
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v510 int32
	_ = v510
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v546 int32
	_ = v546
	var v558 int32
	_ = v558
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v573 int32
	_ = v573
	var v575 int32
	_ = v575
	var v578 int32
	_ = v578
	var v579 int32
	_ = v579
	var v593 int32
	_ = v593
	var v602 int32
	_ = v602
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v619 int32
	_ = v619
	var v621 int32
	_ = v621
	var v635 int32
	_ = v635
	var v638 int32
	_ = v638
	var v639 int32
	_ = v639
	var v674 int32
	_ = v674
	var v687 int32
	_ = v687
	var v700 int32
	_ = v700
	var v702 int32
	_ = v702
	var v716 int32
	_ = v716
	var v731 int32
	_ = v731
	var v747 int32
	_ = v747
	var v761 int32
	_ = v761
	var v763 int32
	_ = v763
	var v764 int32
	_ = v764
	var v775 int32
	_ = v775
	var v781 int32
	_ = v781
	var v782 int32
	_ = v782
	var v796 int32
	_ = v796
	var v809 int32
	_ = v809
	var v810 int32
	_ = v810
	var v813 int32
	_ = v813
	var v814 int32
	_ = v814
	var v828 int32
	_ = v828
	var v830 int32
	_ = v830
	var v837 int32
	_ = v837
	var v841 int32
	_ = v841
	var v842 int32
	_ = v842
	var v845 int32
	_ = v845
	var v848 int64
	_ = v848
	var v875 int32
	_ = v875
	var v877 int32
	_ = v877
	var v890 int32
	_ = v890
	var v891 int32
	_ = v891
	var v907 int32
	_ = v907
	var v922 int32
	_ = v922
	var v923 int32
	_ = v923
	var v924 int32
	_ = v924
	var v927 int32
	_ = v927
	var v962 int32
	_ = v962
	var v975 int32
	_ = v975
	var v992 int32
	_ = v992
	var v1011 int32
	_ = v1011
	var v1026 int32
	_ = v1026
	var v1029 int32
	_ = v1029
	var v1031 int32
	_ = v1031
	var v1033 int32
	_ = v1033
	var v1042 int32
	_ = v1042
	var v1043 int32
	_ = v1043
	var v1044 int32
	_ = v1044
	var v1045 int32
	_ = v1045
	var v1046 int32
	_ = v1046
	var v1047 int32
	_ = v1047
	var v1048 int32
	_ = v1048
	var v1049 int32
	_ = v1049
	var v1051 int32
	_ = v1051
	var v1052 int32
	_ = v1052
	var v1053 int32
	_ = v1053
	var v1066 int32
	_ = v1066
	var v1067 int32
	_ = v1067
	var v1068 int32
	_ = v1068
	var v1082 int32
	_ = v1082
	var v1091 int32
	_ = v1091
	var v1095 int32
	_ = v1095
	var v1096 int32
	_ = v1096
	var v1099 int32
	_ = v1099
	var v1102 int64
	_ = v1102
	var v1127 int32
	_ = v1127
	var v1131 int32
	_ = v1131
	var v1133 int32
	_ = v1133
	var v1146 int32
	_ = v1146
	var v1150 int32
	_ = v1150
	var v1151 int32
	_ = v1151
	var v1183 int32
	_ = v1183
	var v1186 int32
	_ = v1186
	var v1201 int32
	_ = v1201
	var v1207 int32
	_ = v1207
	var v1208 int32
	_ = v1208
	var v1220 int32
	_ = v1220
	var v1237 int32
	_ = v1237
	var v1243 int32
	_ = v1243
	var v1244 int32
	_ = v1244
	var v1248 int32
	_ = v1248
	var v1249 int32
	_ = v1249
	var v1255 int32
	_ = v1255
	var v1260 int32
	_ = v1260
	var v1272 int32
	_ = v1272
	var v1293 int32
	_ = v1293
	var v1294 int64
	_ = v1294
	var v1298 int32
	_ = v1298
	var v1300 int32
	_ = v1300
	var v1301 int32
	_ = v1301
	var v1304 int32
	_ = v1304
	var v1306 int32
	_ = v1306
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
	var v1320 int32
	_ = v1320
	v3 = int32(0)
	v21 = m.G0
	v23 = v21 - int32(496)
	m.G0 = v23
	v29 = v3
	v30 = v3
	v31 = int32(-1)
	v32 = v3
	v33 = v3
	v34 = v3
	v35 = v3
	v36 = v3
	v37 = v3
	v38 = v3
	v39 = v3
	v40 = v3
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
	m.G0 = v23 + int32(496)
	return
L4:
	;
	goto L3
L5:
	;
	if v31 == int32(1) {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	v1293 = int32(m.ExcTag)
	v1294 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v1293 == int32(0) {
		goto L145
	} else {
		goto L146
	}
L8:
	;
	if v1052 == int32(0) {
		goto L122
	} else {
		goto L123
	}
L9:
	;
	v1042 = v29
	v1043 = v30
	v1044 = v37
	v1045 = v32
	v1046 = v33
	v1047 = v34
	v1048 = v35
	v1049 = v36
	v1051 = v38
	v1052 = v39
	v1053 = v40
	goto L8
L10:
	;
	goto L11
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+456)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v23)+460)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v23)+464)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v23)+472)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v23)+476)) = v30
	*(*int32)(unsafe.Add(mBase, uint32(v23)+480)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v23)+488)) = v29
	*(*int32)(unsafe.Add(mBase, uint32(v23)+492)) = v32
	v56 = int32(1)
	v57 = v38 & v56
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+471)) = uint8(v57)
	v60 = v40 & v56
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+487)) = uint8(v60)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+300)) = int32(0)
	v66 = F_table_open(m, int32(_a_F_DropSubscription_0), int32(3))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L7
	} else {
		goto L12
	}
L12:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+456)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v23)+460)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v23)+464)) = v33
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+471)) = uint8(v57)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+472)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v23)+476)) = v30
	*(*int32)(unsafe.Add(mBase, uint32(v23)+480)) = v37
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+487)) = uint8(v60)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+488)) = v29
	*(*int32)(unsafe.Add(mBase, uint32(v23)+492)) = v66
	v81 = *(*int32)(unsafe.Add(mBase, _c_F_DropSubscription[0]))
	v82 = F_SearchSysCache2(m, int32(66), v81, v68)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L7
	} else {
		goto L13
	}
L13:
	;
	if v82 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+460)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v23)+456)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v23)+464)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v23)+472)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v23)+476)) = v30
	*(*int32)(unsafe.Add(mBase, uint32(v23)+480)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v23)+488)) = v29
	*(*int32)(unsafe.Add(mBase, uint32(v23)+492)) = v66
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+471)) = uint8(v57)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+487)) = uint8(v60)
	F_relation_close(m, v66, int32(0))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L7
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v82)+16))
	v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207)+22)))
	v209 = v207 + v208
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v209)+80))
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v209)))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+460)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v23)+456)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v23)+464)) = v33
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+471)) = uint8(v57)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+472)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v23)+476)) = v30
	*(*int32)(unsafe.Add(mBase, uint32(v23)+480)) = v37
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+487)) = uint8(v60)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+488)) = v211
	*(*int32)(unsafe.Add(mBase, uint32(v23)+492)) = v66
	v222 = F_superuser_arg(m, v210)
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L7
	} else {
		goto L29
	}
L17:
	;
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
	if v99 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+460)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v23)+456)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v23)+464)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v23)+472)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v23)+476)) = v30
	*(*int32)(unsafe.Add(mBase, uint32(v23)+480)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v23)+488)) = v29
	*(*int32)(unsafe.Add(mBase, uint32(v23)+492)) = v66
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+471)) = uint8(v57)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+487)) = uint8(v60)
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L7
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+460)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v23)+456)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v23)+464)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v23)+472)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v23)+476)) = v30
	*(*int32)(unsafe.Add(mBase, uint32(v23)+480)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v23)+488)) = v29
	*(*int32)(unsafe.Add(mBase, uint32(v23)+492)) = v66
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+471)) = uint8(v57)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+487)) = uint8(v60)
	v173 = F_errstart(m, int32(18), int32(0))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L7
	} else {
		goto L25
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+460)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v23)+456)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v23)+464)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v23)+472)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v23)+476)) = v30
	*(*int32)(unsafe.Add(mBase, uint32(v23)+480)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v23)+488)) = v29
	*(*int32)(unsafe.Add(mBase, uint32(v23)+492)) = v66
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+471)) = uint8(v57)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+487)) = uint8(v60)
	F_errcode(m, int32(67137668))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L7
	} else {
		goto L22
	}
L22:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+460)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v23)+456)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v23)+464)) = v33
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+471)) = uint8(v57)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+472)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v23)+476)) = v30
	*(*int32)(unsafe.Add(mBase, uint32(v23)+480)) = v37
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+487)) = uint8(v60)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+488)) = v29
	*(*int32)(unsafe.Add(mBase, uint32(v23)+492)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = v129
	F_errmsg(m, int32(_a_F_DropSubscription_1), v23+int32(16))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L7
	} else {
		goto L23
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+460)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v23)+456)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v23)+464)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v23)+472)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v23)+476)) = v30
	*(*int32)(unsafe.Add(mBase, uint32(v23)+480)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v23)+488)) = v29
	*(*int32)(unsafe.Add(mBase, uint32(v23)+492)) = v66
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+471)) = uint8(v57)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+487)) = uint8(v60)
	F_errfinish(m, int32(_a_F_DropSubscription_2), int32(1666), int32(_a_F_DropSubscription_3))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L7
	} else {
		goto L24
	}
L24:
	;
	goto L1
L25:
	;
	if v173 == int32(0) {
		goto L4
	} else {
		goto L26
	}
L26:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+460)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v23)+456)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v23)+464)) = v33
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+471)) = uint8(v57)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+472)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v23)+476)) = v30
	*(*int32)(unsafe.Add(mBase, uint32(v23)+480)) = v37
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+487)) = uint8(v60)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+488)) = v29
	*(*int32)(unsafe.Add(mBase, uint32(v23)+492)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v23))) = v177
	F_errmsg(m, int32(_a_F_DropSubscription_4), v23)
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L7
	} else {
		goto L27
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+460)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v23)+456)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v23)+464)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v23)+472)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v23)+476)) = v30
	*(*int32)(unsafe.Add(mBase, uint32(v23)+480)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v23)+488)) = v29
	*(*int32)(unsafe.Add(mBase, uint32(v23)+492)) = v66
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+471)) = uint8(v57)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+487)) = uint8(v60)
	F_errfinish(m, int32(_a_F_DropSubscription_2), int32(1670), int32(_a_F_DropSubscription_3))
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L7
	} else {
		goto L28
	}
L28:
	;
	goto L4
L29:
	;
	v224 = int32(0)
	if v222 == v224 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209)+89)))
	v228 = v227
	goto L32
L31:
	;
	v228 = v224
	goto L32
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+460)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v23)+456)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v23)+464)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v23)+472)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v23)+476)) = v30
	*(*int32)(unsafe.Add(mBase, uint32(v23)+480)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v23)+488)) = v211
	*(*int32)(unsafe.Add(mBase, uint32(v23)+492)) = v66
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+471)) = uint8(v57)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+487)) = uint8(v60)
	v240 = *(*int32)(unsafe.Add(mBase, _c_F_DropSubscription[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+460)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v23)+456)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v23)+464)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v23)+472)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v23)+476)) = v30
	*(*int32)(unsafe.Add(mBase, uint32(v23)+480)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v23)+488)) = v211
	*(*int32)(unsafe.Add(mBase, uint32(v23)+492)) = v66
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+471)) = uint8(v57)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+487)) = uint8(v60)
	v252 = F_object_ownercheck(m, int32(_a_F_DropSubscription_0), v211, v240)
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L7
	} else {
		goto L33
	}
L33:
	;
	if v252 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v256 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+460)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v23)+456)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v23)+464)) = v33
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+471)) = uint8(v57)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+472)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v23)+476)) = v30
	*(*int32)(unsafe.Add(mBase, uint32(v23)+480)) = v37
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+487)) = uint8(v60)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+488)) = v211
	*(*int32)(unsafe.Add(mBase, uint32(v23)+492)) = v66
	F_aclcheck_error(m, int32(2), int32(38), v256)
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L7
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	v273 = *(*int32)(unsafe.Add(mBase, _c_F_DropSubscription[2]))
	if v273 != 0 {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	goto L36
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+460)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v23)+456)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v23)+464)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v23)+472)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v23)+476)) = v30
	*(*int32)(unsafe.Add(mBase, uint32(v23)+480)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v23)+488)) = v211
	*(*int32)(unsafe.Add(mBase, uint32(v23)+492)) = v66
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+471)) = uint8(v57)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+487)) = uint8(v60)
	v285 = int32(0)
	F_RunObjectDropHook(m, int32(_a_F_DropSubscription_0), v211, v285, v285)
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L7
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+460)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v23)+456)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v23)+464)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v23)+472)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v23)+476)) = v30
	*(*int32)(unsafe.Add(mBase, uint32(v23)+480)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v23)+488)) = v211
	*(*int32)(unsafe.Add(mBase, uint32(v23)+492)) = v66
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+471)) = uint8(v57)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+487)) = uint8(v60)
	F_LockSharedObject(m, int32(_a_F_DropSubscription_0), v211, int32(8))
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L7
	} else {
		goto L42
	}
L41:
	;
	goto L40
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+460)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v23)+456)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v23)+464)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v23)+472)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v23)+476)) = v30
	*(*int32)(unsafe.Add(mBase, uint32(v23)+480)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v23)+488)) = v211
	*(*int32)(unsafe.Add(mBase, uint32(v23)+492)) = v66
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+471)) = uint8(v57)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+487)) = uint8(v60)
	v315 = F_SysCacheGetAttrNotNull(m, int32(67), v82, int32(4))
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L7
	} else {
		goto L43
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+460)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v23)+456)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v23)+464)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v23)+472)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v23)+476)) = v30
	*(*int32)(unsafe.Add(mBase, uint32(v23)+480)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v23)+488)) = v211
	*(*int32)(unsafe.Add(mBase, uint32(v23)+492)) = v66
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+471)) = uint8(v57)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+487)) = uint8(v60)
	v327 = F_pstrdup(m, v315)
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L7
	} else {
		goto L44
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+460)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v23)+456)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v23)+464)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v23)+472)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v23)+476)) = v30
	*(*int32)(unsafe.Add(mBase, uint32(v23)+480)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v23)+488)) = v211
	*(*int32)(unsafe.Add(mBase, uint32(v23)+492)) = v66
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+471)) = uint8(v57)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+487)) = uint8(v60)
	v341 = F_SysCacheGetAttrNotNull(m, int32(67), v82, int32(14))
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L7
	} else {
		goto L45
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+460)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v23)+456)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v23)+464)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v23)+472)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v23)+476)) = v30
	*(*int32)(unsafe.Add(mBase, uint32(v23)+480)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v23)+488)) = v211
	*(*int32)(unsafe.Add(mBase, uint32(v23)+492)) = v66
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+471)) = uint8(v57)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+487)) = uint8(v60)
	v353 = F_text_to_cstring(m, v341)
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L7
	} else {
		goto L46
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+460)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v23)+456)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v23)+464)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v23)+472)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v23)+476)) = v30
	*(*int32)(unsafe.Add(mBase, uint32(v23)+480)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v23)+488)) = v211
	*(*int32)(unsafe.Add(mBase, uint32(v23)+492)) = v66
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+471)) = uint8(v57)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+487)) = uint8(v60)
	v369 = F_SysCacheGetAttr(m, int32(67), v82, int32(15), v23+int32(371))
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
		goto L7
	} else {
		goto L47
	}
L47:
	;
	v371 = int32(0)
	v372 = int32(1)
	v373 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+371)))
	if v373 != 0 {
		v402 = v371
		v404 = v372
		goto L48
	} else {
		goto L49
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+380)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+376)) = v211
	*(*int32)(unsafe.Add(mBase, uint32(v23)+372)) = int32(_a_F_DropSubscription_0)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+456)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v23)+460)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v23)+464)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v23)+472)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v23)+476)) = v30
	*(*int32)(unsafe.Add(mBase, uint32(v23)+480)) = v402
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+487)) = uint8(v404)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+488)) = v211
	*(*int32)(unsafe.Add(mBase, uint32(v23)+492)) = v66
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+471)) = uint8(v57)
	v422 = int32(1)
	F_EventTriggerSQLDropAddObject(m, v23+int32(372), v422, v422)
	mBase = m.M
	v425 = m.ExcPending
	if v425 != 0 {
		goto L7
	} else {
		goto L53
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+460)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v23)+456)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v23)+464)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v23)+472)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v23)+476)) = v30
	*(*int32)(unsafe.Add(mBase, uint32(v23)+480)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v23)+488)) = v211
	*(*int32)(unsafe.Add(mBase, uint32(v23)+492)) = v66
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+471)) = uint8(v57)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+487)) = uint8(v60)
	v384 = F_pstrdup(m, v369)
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L7
	} else {
		goto L50
	}
L50:
	;
	if v384 == int32(0) {
		v402 = v371
		v404 = v372
		goto L48
	} else {
		goto L51
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+460)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v23)+456)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v23)+464)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v23)+472)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v23)+476)) = v30
	*(*int32)(unsafe.Add(mBase, uint32(v23)+480)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v23)+488)) = v211
	*(*int32)(unsafe.Add(mBase, uint32(v23)+492)) = v66
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+471)) = uint8(v57)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+487)) = uint8(v60)
	F_PreventInTransactionBlock(m, l1, int32(_a_F_DropSubscription_5))
	mBase = m.M
	v401 = m.ExcPending
	if v401 != 0 {
		goto L7
	} else {
		goto L52
	}
L52:
	;
	v402 = v384
	v404 = int32(0)
	goto L48
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+460)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v23)+456)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v23)+464)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v23)+472)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v23)+476)) = v30
	*(*int32)(unsafe.Add(mBase, uint32(v23)+480)) = v402
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+487)) = uint8(v404)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+488)) = v211
	*(*int32)(unsafe.Add(mBase, uint32(v23)+492)) = v66
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+471)) = uint8(v57)
	F_simple_heap_delete(m, v66, v82+int32(4))
	mBase = m.M
	v439 = m.ExcPending
	if v439 != 0 {
		goto L7
	} else {
		goto L54
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+460)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v23)+456)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v23)+464)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v23)+472)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v23)+476)) = v30
	*(*int32)(unsafe.Add(mBase, uint32(v23)+480)) = v402
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+487)) = uint8(v404)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+488)) = v211
	*(*int32)(unsafe.Add(mBase, uint32(v23)+492)) = v66
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+471)) = uint8(v57)
	F_ReleaseCatCache(m, v82)
	mBase = m.M
	v451 = m.ExcPending
	if v451 != 0 {
		goto L7
	} else {
		goto L55
	}
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+460)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v23)+456)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v23)+464)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v23)+472)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v23)+476)) = v30
	*(*int32)(unsafe.Add(mBase, uint32(v23)+480)) = v402
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+487)) = uint8(v404)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+488)) = v211
	*(*int32)(unsafe.Add(mBase, uint32(v23)+492)) = v66
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+471)) = uint8(v57)
	v464 = F_logicalrep_workers_find(m, v211, int32(0), int32(1))
	mBase = m.M
	v465 = m.ExcPending
	if v465 != 0 {
		goto L7
	} else {
		goto L57
	}
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+460)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v23)+456)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v23)+464)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v23)+472)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v23)+476)) = v30
	*(*int32)(unsafe.Add(mBase, uint32(v23)+480)) = v402
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+487)) = uint8(v404)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+488)) = v211
	*(*int32)(unsafe.Add(mBase, uint32(v23)+492)) = v66
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+471)) = uint8(v57)
	F_list_free(m, v464)
	mBase = m.M
	v546 = m.ExcPending
	if v546 != 0 {
		goto L7
	} else {
		goto L64
	}
L57:
	;
	if v464 == int32(0) {
		goto L56
	} else {
		goto L58
	}
L58:
	;
	v468 = int32(0)
	v469 = *(*int32)(unsafe.Add(mBase, uint32(v464)+4))
	if v469 <= v468 {
		goto L56
	} else {
		goto L59
	}
L59:
	;
	v483 = v468
	goto L60
L60:
	;
	v492 = *(*int32)(unsafe.Add(mBase, uint32(v464)+12))
	v496 = *(*int32)(unsafe.Add(mBase, uint32(v492+v483<<(uint(int32(2))%32))))
	v497 = *(*int32)(unsafe.Add(mBase, uint32(v496)+36))
	v498 = *(*int32)(unsafe.Add(mBase, uint32(v496)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+460)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v23)+456)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v23)+464)) = v33
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+471)) = uint8(v57)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+472)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v23)+476)) = v30
	*(*int32)(unsafe.Add(mBase, uint32(v23)+480)) = v402
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+487)) = uint8(v404)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+488)) = v211
	*(*int32)(unsafe.Add(mBase, uint32(v23)+492)) = v66
	F_logicalrep_worker_stop(m, v498, v497)
	mBase = m.M
	v510 = m.ExcPending
	if v510 != 0 {
		goto L7
	} else {
		goto L62
	}
L61:
	;
	goto L56
L62:
	;
	v512 = v483 + int32(1)
	v513 = *(*int32)(unsafe.Add(mBase, uint32(v464)+4))
	if v512 < v513 {
		v483 = v512
		goto L60
	} else {
		goto L63
	}
L63:
	;
	goto L61
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+460)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v23)+456)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v23)+464)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v23)+472)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v23)+476)) = v30
	*(*int32)(unsafe.Add(mBase, uint32(v23)+480)) = v402
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+487)) = uint8(v404)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+488)) = v211
	*(*int32)(unsafe.Add(mBase, uint32(v23)+492)) = v66
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+471)) = uint8(v57)
	F_ApplyLauncherForgetWorkerStartTime(m, v211)
	mBase = m.M
	v558 = m.ExcPending
	if v558 != 0 {
		goto L7
	} else {
		goto L65
	}
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+460)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v23)+456)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v23)+464)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v23)+472)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v23)+476)) = v30
	*(*int32)(unsafe.Add(mBase, uint32(v23)+480)) = v402
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+487)) = uint8(v404)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+488)) = v211
	*(*int32)(unsafe.Add(mBase, uint32(v23)+492)) = v66
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+471)) = uint8(v57)
	v570 = F_GetSubscriptionRelations(m, v211, int32(1))
	mBase = m.M
	v571 = m.ExcPending
	if v571 != 0 {
		goto L7
	} else {
		goto L66
	}
L66:
	;
	v573 = v570 + int32(4)
	v575 = base.B2i32(v570 == int32(0))
	if v570 == int32(0) {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+460)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v23)+456)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v23)+464)) = v33
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+471)) = uint8(v575)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+472)) = v573
	*(*int32)(unsafe.Add(mBase, uint32(v23)+476)) = v570
	*(*int32)(unsafe.Add(mBase, uint32(v23)+480)) = v402
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+487)) = uint8(v404)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+488)) = v211
	*(*int32)(unsafe.Add(mBase, uint32(v23)+492)) = v66
	F_deleteSharedDependencyRecordsFor(m, int32(_a_F_DropSubscription_0), v211, int32(0))
	mBase = m.M
	v674 = m.ExcPending
	if v674 != 0 {
		goto L7
	} else {
		goto L78
	}
L68:
	;
	v578 = int32(0)
	v579 = *(*int32)(unsafe.Add(mBase, uint32(v570)+4))
	if v579 <= v578 {
		goto L67
	} else {
		goto L69
	}
L69:
	;
	v593 = v578
	goto L70
L70:
	;
	v602 = *(*int32)(unsafe.Add(mBase, uint32(v570)+12))
	v606 = *(*int32)(unsafe.Add(mBase, uint32(v602+v593<<(uint(int32(2))%32))))
	v607 = *(*int32)(unsafe.Add(mBase, uint32(v606)))
	if v607 != 0 {
		goto L72
	} else {
		goto L73
	}
L71:
	;
	goto L67
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+460)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v23)+456)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v23)+464)) = v33
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+471)) = uint8(v575)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+472)) = v573
	*(*int32)(unsafe.Add(mBase, uint32(v23)+476)) = v570
	*(*int32)(unsafe.Add(mBase, uint32(v23)+480)) = v402
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+487)) = uint8(v404)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+488)) = v211
	*(*int32)(unsafe.Add(mBase, uint32(v23)+492)) = v66
	v619 = v23 + int32(304)
	F_ReplicationOriginNameForLogicalRep(m, v211, v607, v619)
	mBase = m.M
	v621 = m.ExcPending
	if v621 != 0 {
		goto L7
	} else {
		goto L75
	}
L73:
	;
	goto L74
L74:
	;
	v638 = v593 + int32(1)
	v639 = *(*int32)(unsafe.Add(mBase, uint32(v573)))
	if v638 < v639 {
		v593 = v638
		goto L70
	} else {
		goto L77
	}
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+460)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v23)+456)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v23)+464)) = v33
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+471)) = uint8(v575)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+472)) = v573
	*(*int32)(unsafe.Add(mBase, uint32(v23)+476)) = v570
	*(*int32)(unsafe.Add(mBase, uint32(v23)+480)) = v402
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+487)) = uint8(v404)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+488)) = v211
	*(*int32)(unsafe.Add(mBase, uint32(v23)+492)) = v66
	F_replorigin_drop_by_name(m, v619, int32(1), int32(0))
	mBase = m.M
	v635 = m.ExcPending
	if v635 != 0 {
		goto L7
	} else {
		goto L76
	}
L76:
	;
	goto L74
L77:
	;
	goto L71
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+460)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v23)+456)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v23)+464)) = v33
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+471)) = uint8(v575)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+472)) = v573
	*(*int32)(unsafe.Add(mBase, uint32(v23)+476)) = v570
	*(*int32)(unsafe.Add(mBase, uint32(v23)+480)) = v402
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+487)) = uint8(v404)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+488)) = v211
	*(*int32)(unsafe.Add(mBase, uint32(v23)+492)) = v66
	F_RemoveSubscriptionRel(m, v211, int32(0))
	mBase = m.M
	v687 = m.ExcPending
	if v687 != 0 {
		goto L7
	} else {
		goto L79
	}
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+460)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v23)+456)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v23)+464)) = v33
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+471)) = uint8(v575)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+472)) = v573
	*(*int32)(unsafe.Add(mBase, uint32(v23)+476)) = v570
	*(*int32)(unsafe.Add(mBase, uint32(v23)+480)) = v402
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+487)) = uint8(v404)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+488)) = v211
	*(*int32)(unsafe.Add(mBase, uint32(v23)+492)) = v66
	v700 = v23 + int32(304)
	F_ReplicationOriginNameForLogicalRep(m, v211, int32(0), v700)
	mBase = m.M
	v702 = m.ExcPending
	if v702 != 0 {
		goto L7
	} else {
		goto L80
	}
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+460)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v23)+456)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v23)+464)) = v33
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+471)) = uint8(v575)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+472)) = v573
	*(*int32)(unsafe.Add(mBase, uint32(v23)+476)) = v570
	*(*int32)(unsafe.Add(mBase, uint32(v23)+480)) = v402
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+487)) = uint8(v404)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+488)) = v211
	*(*int32)(unsafe.Add(mBase, uint32(v23)+492)) = v66
	F_replorigin_drop_by_name(m, v700, int32(1), int32(0))
	mBase = m.M
	v716 = m.ExcPending
	if v716 != 0 {
		goto L7
	} else {
		goto L81
	}
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+460)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v23)+456)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v23)+464)) = v33
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+471)) = uint8(v575)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+472)) = v573
	*(*int32)(unsafe.Add(mBase, uint32(v23)+476)) = v570
	*(*int32)(unsafe.Add(mBase, uint32(v23)+480)) = v402
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+487)) = uint8(v404)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+488)) = v211
	*(*int32)(unsafe.Add(mBase, uint32(v23)+492)) = v66
	F_pgstat_drop_transactional(m, int32(5), int32(0), base.I64_extend_i32_u(v211))
	mBase = m.M
	v731 = m.ExcPending
	if v731 != 0 {
		goto L7
	} else {
		goto L82
	}
L82:
	;
	if base.B2i32(v570 == int32(0))&v404 != 0 {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+460)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v23)+456)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v23)+464)) = v33
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+471)) = uint8(v575)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+472)) = v573
	*(*int32)(unsafe.Add(mBase, uint32(v23)+476)) = v570
	*(*int32)(unsafe.Add(mBase, uint32(v23)+480)) = v402
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+487)) = uint8(v404)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+488)) = v211
	*(*int32)(unsafe.Add(mBase, uint32(v23)+492)) = v66
	F_relation_close(m, v66, int32(0))
	mBase = m.M
	v747 = m.ExcPending
	if v747 != 0 {
		goto L7
	} else {
		goto L86
	}
L84:
	;
	goto L85
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+460)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v23)+456)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v23)+464)) = v33
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+471)) = uint8(v575)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+472)) = v573
	*(*int32)(unsafe.Add(mBase, uint32(v23)+476)) = v570
	*(*int32)(unsafe.Add(mBase, uint32(v23)+480)) = v402
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+487)) = uint8(v404)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+488)) = v211
	*(*int32)(unsafe.Add(mBase, uint32(v23)+492)) = v66
	F_load_file(m, int32(_a_F_DropSubscription_6), int32(0))
	mBase = m.M
	v761 = m.ExcPending
	if v761 != 0 {
		goto L7
	} else {
		goto L87
	}
L86:
	;
	goto L4
L87:
	;
	v763 = *(*int32)(unsafe.Add(mBase, _c_F_DropSubscription[3]))
	v764 = *(*int32)(unsafe.Add(mBase, uint32(v763)))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+456)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v23)+460)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v23)+464)) = v33
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+471)) = uint8(v575)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+472)) = v573
	*(*int32)(unsafe.Add(mBase, uint32(v23)+476)) = v570
	*(*int32)(unsafe.Add(mBase, uint32(v23)+480)) = v402
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+487)) = uint8(v404)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+488)) = v211
	*(*int32)(unsafe.Add(mBase, uint32(v23)+492)) = v66
	v775 = int32(1)
	v781 = m.T0[v764].(func(*base.Module, int32, int32, int32, int32, int32, int32) int32)(m, v353, v775, v775, v228&v775, v327, v23+int32(300))
	mBase = m.M
	v782 = m.ExcPending
	if v782 != 0 {
		goto L7
	} else {
		goto L88
	}
L88:
	;
	if v781 == int32(0) {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	if v404 != 0 {
		goto L92
	} else {
		goto L93
	}
L90:
	;
	goto L91
L91:
	;
	v1029 = *(*int32)(unsafe.Add(mBase, _c_F_DropSubscription[4]))
	v1031 = *(*int32)(unsafe.Add(mBase, _c_F_DropSubscription[5]))
	goto L118
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+460)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v23)+456)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v23)+464)) = v781
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+471)) = uint8(v575)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+472)) = v573
	*(*int32)(unsafe.Add(mBase, uint32(v23)+476)) = v570
	*(*int32)(unsafe.Add(mBase, uint32(v23)+480)) = v402
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+487)) = uint8(v404)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+488)) = v211
	*(*int32)(unsafe.Add(mBase, uint32(v23)+492)) = v66
	F_list_free(m, v570)
	mBase = m.M
	v796 = m.ExcPending
	if v796 != 0 {
		goto L7
	} else {
		goto L95
	}
L93:
	;
	goto L94
L94:
	;
	v810 = *(*int32)(unsafe.Add(mBase, uint32(v23)+300))
	if v570 == int32(0) {
		goto L97
	} else {
		goto L98
	}
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+460)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v23)+456)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v23)+464)) = v781
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+471)) = uint8(v575)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+472)) = v573
	*(*int32)(unsafe.Add(mBase, uint32(v23)+476)) = v570
	*(*int32)(unsafe.Add(mBase, uint32(v23)+480)) = v402
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+487)) = uint8(v404)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+488)) = v211
	*(*int32)(unsafe.Add(mBase, uint32(v23)+492)) = v66
	F_relation_close(m, v66, int32(0))
	mBase = m.M
	v809 = m.ExcPending
	if v809 != 0 {
		goto L7
	} else {
		goto L96
	}
L96:
	;
	goto L4
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+460)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v23)+456)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v23)+464)) = v781
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+471)) = uint8(v575)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+472)) = v573
	*(*int32)(unsafe.Add(mBase, uint32(v23)+476)) = v570
	*(*int32)(unsafe.Add(mBase, uint32(v23)+480)) = v402
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+487)) = uint8(v404)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+488)) = v211
	*(*int32)(unsafe.Add(mBase, uint32(v23)+492)) = v66
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v962 = m.ExcPending
	if v962 != 0 {
		goto L7
	} else {
		goto L113
	}
L98:
	;
	v813 = int32(0)
	v814 = *(*int32)(unsafe.Add(mBase, uint32(v570)+4))
	if v814 <= v813 {
		goto L97
	} else {
		goto L99
	}
L99:
	;
	v828 = v813
	v830 = v814
	goto L100
L100:
	;
	v837 = *(*int32)(unsafe.Add(mBase, uint32(v570)+12))
	v841 = *(*int32)(unsafe.Add(mBase, uint32(v837+v828<<(uint(int32(2))%32))))
	v842 = *(*int32)(unsafe.Add(mBase, uint32(v841)))
	if v842 == int32(0) {
		v924 = v830
		goto L102
	} else {
		goto L103
	}
L101:
	;
	goto L97
L102:
	;
	v927 = v828 + int32(1)
	if v927 < v924 {
		v828 = v927
		v830 = v924
		goto L100
	} else {
		goto L112
	}
L103:
	;
	v845 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v841)+16)))
	if v845 == int32(115) {
		v924 = v830
		goto L102
	} else {
		goto L104
	}
L104:
	;
	v848 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v23)+440)) = v848
	*(*int64)(unsafe.Add(mBase, uint32(v23)+432)) = v848
	*(*int64)(unsafe.Add(mBase, uint32(v23)+424)) = v848
	*(*int64)(unsafe.Add(mBase, uint32(v23)+416)) = v848
	*(*int64)(unsafe.Add(mBase, uint32(v23)+408)) = v848
	*(*int64)(unsafe.Add(mBase, uint32(v23)+400)) = v848
	*(*int64)(unsafe.Add(mBase, uint32(v23)+392)) = v848
	*(*int64)(unsafe.Add(mBase, uint32(v23)+384)) = v848
	*(*int32)(unsafe.Add(mBase, uint32(v23)+456)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v23)+460)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v23)+464)) = v781
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+471)) = uint8(v575)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+472)) = v573
	*(*int32)(unsafe.Add(mBase, uint32(v23)+476)) = v570
	*(*int32)(unsafe.Add(mBase, uint32(v23)+480)) = v402
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+487)) = uint8(v404)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+488)) = v211
	*(*int32)(unsafe.Add(mBase, uint32(v23)+492)) = v66
	v875 = v23 + int32(384)
	F_ReplicationSlotNameForTablesync(m, v211, v842, v875)
	mBase = m.M
	v877 = m.ExcPending
	if v877 != 0 {
		goto L7
	} else {
		goto L105
	}
L105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+460)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v23)+456)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v23)+464)) = v781
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+471)) = uint8(v575)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+472)) = v573
	*(*int32)(unsafe.Add(mBase, uint32(v23)+476)) = v570
	*(*int32)(unsafe.Add(mBase, uint32(v23)+480)) = v402
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+487)) = uint8(v404)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+488)) = v211
	*(*int32)(unsafe.Add(mBase, uint32(v23)+492)) = v66
	v890 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v891 = m.ExcPending
	if v891 != 0 {
		goto L7
	} else {
		goto L106
	}
L106:
	;
	if v890 != 0 {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+460)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v23)+456)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v23)+464)) = v781
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+471)) = uint8(v575)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+472)) = v573
	*(*int32)(unsafe.Add(mBase, uint32(v23)+476)) = v570
	*(*int32)(unsafe.Add(mBase, uint32(v23)+480)) = v402
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+487)) = uint8(v404)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+488)) = v211
	*(*int32)(unsafe.Add(mBase, uint32(v23)+492)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v23)+64)) = v875
	F_errmsg_internal(m, int32(_a_F_DropSubscription_7), v23-int32(-64))
	mBase = m.M
	v907 = m.ExcPending
	if v907 != 0 {
		goto L7
	} else {
		goto L110
	}
L108:
	;
	goto L109
L109:
	;
	v923 = *(*int32)(unsafe.Add(mBase, uint32(v573)))
	v924 = v923
	goto L102
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+460)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v23)+456)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v23)+464)) = v781
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+471)) = uint8(v575)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+472)) = v573
	*(*int32)(unsafe.Add(mBase, uint32(v23)+476)) = v570
	*(*int32)(unsafe.Add(mBase, uint32(v23)+480)) = v402
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+487)) = uint8(v404)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+488)) = v211
	*(*int32)(unsafe.Add(mBase, uint32(v23)+492)) = v66
	F_errfinish(m, int32(_a_F_DropSubscription_2), int32(2342), int32(_a_F_DropSubscription_8))
	mBase = m.M
	v922 = m.ExcPending
	if v922 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v23)+460)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v23)+456)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v23)+464)) = v781
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+471)) = uint8(v575)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+472)) = v573
	*(*int32)(unsafe.Add(mBase, uint32(v23)+476)) = v570
	*(*int32)(unsafe.Add(mBase, uint32(v23)+480)) = v402
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+487)) = uint8(v404)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+488)) = v211
	*(*int32)(unsafe.Add(mBase, uint32(v23)+492)) = v66
	F_errcode(m, int32(100663808))
	mBase = m.M
	v975 = m.ExcPending
	if v975 != 0 {
		goto L7
	} else {
		goto L114
	}
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+460)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v23)+456)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v23)+464)) = v781
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+471)) = uint8(v575)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+472)) = v573
	*(*int32)(unsafe.Add(mBase, uint32(v23)+476)) = v570
	*(*int32)(unsafe.Add(mBase, uint32(v23)+480)) = v402
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+487)) = uint8(v404)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+488)) = v211
	*(*int32)(unsafe.Add(mBase, uint32(v23)+492)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v23)+52)) = v810
	*(*int32)(unsafe.Add(mBase, uint32(v23)+48)) = v402
	F_errmsg(m, int32(_a_F_DropSubscription_9), v23+int32(48))
	mBase = m.M
	v992 = m.ExcPending
	if v992 != 0 {
		goto L7
	} else {
		goto L115
	}
L115:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+460)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v23)+456)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v23)+464)) = v781
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+471)) = uint8(v575)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+472)) = v573
	*(*int32)(unsafe.Add(mBase, uint32(v23)+476)) = v570
	*(*int32)(unsafe.Add(mBase, uint32(v23)+480)) = v402
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+487)) = uint8(v404)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+488)) = v211
	*(*int32)(unsafe.Add(mBase, uint32(v23)+492)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v23)+36)) = int32(_a_F_DropSubscription_10)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+32)) = int32(_a_F_DropSubscription_11)
	F_errhint(m, int32(_a_F_DropSubscription_12), v23+int32(32))
	mBase = m.M
	v1011 = m.ExcPending
	if v1011 != 0 {
		goto L7
	} else {
		goto L116
	}
L116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+460)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v23)+456)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v23)+464)) = v781
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+471)) = uint8(v575)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+472)) = v573
	*(*int32)(unsafe.Add(mBase, uint32(v23)+476)) = v570
	*(*int32)(unsafe.Add(mBase, uint32(v23)+480)) = v402
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+487)) = uint8(v404)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+488)) = v211
	*(*int32)(unsafe.Add(mBase, uint32(v23)+492)) = v66
	F_errfinish(m, int32(_a_F_DropSubscription_2), int32(2353), int32(_a_F_DropSubscription_8))
	mBase = m.M
	v1026 = m.ExcPending
	if v1026 != 0 {
		goto L7
	} else {
		goto L117
	}
L117:
	;
	goto L1
L118:
	;
	v1033 = v23 + int32(144)
	*(*int32)(unsafe.Add(mBase, uint32(v1033)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1033))) = v23 + int32(76)
	goto L121
L119:
	;
	v1042 = v211
	v1043 = v570
	v1044 = v402
	v1045 = v66
	v1046 = v781
	v1047 = v573
	v1048 = v1029
	v1049 = v1031
	v1051 = v575
	v1052 = int32(0)
	v1053 = v404
	goto L8
L121:
	;
	goto L119
L122:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_DropSubscription[5])) = v23 + int32(144)
	v1066 = v1051 & int32(1)
	if v1066 != 0 {
		goto L125
	} else {
		goto L126
	}
L123:
	;
	goto L124
L124:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_DropSubscription[4])) = v1048
	*(*int32)(unsafe.Add(mBase, _c_F_DropSubscription[5])) = v1049
	v1243 = *(*int32)(unsafe.Add(mBase, _c_F_DropSubscription[3]))
	v1244 = *(*int32)(unsafe.Add(mBase, uint32(v1243)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+456)) = v1048
	*(*int32)(unsafe.Add(mBase, uint32(v23)+460)) = v1049
	*(*int32)(unsafe.Add(mBase, uint32(v23)+464)) = v1046
	v1248 = int32(1)
	v1249 = v1051 & v1248
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+471)) = uint8(v1249)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+472)) = v1047
	*(*int32)(unsafe.Add(mBase, uint32(v23)+476)) = v1043
	*(*int32)(unsafe.Add(mBase, uint32(v23)+480)) = v1044
	v1255 = v1053 & v1248
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+487)) = uint8(v1255)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+488)) = v1042
	*(*int32)(unsafe.Add(mBase, uint32(v23)+492)) = v1045
	m.T0[v1244].(func(*base.Module, int32))(m, v1046)
	mBase = m.M
	v1260 = m.ExcPending
	if v1260 != 0 {
		goto L7
	} else {
		goto L143
	}
L125:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+460)) = v1049
	*(*int32)(unsafe.Add(mBase, uint32(v23)+456)) = v1048
	*(*int32)(unsafe.Add(mBase, uint32(v23)+464)) = v1046
	*(*int32)(unsafe.Add(mBase, uint32(v23)+472)) = v1047
	*(*int32)(unsafe.Add(mBase, uint32(v23)+476)) = v1043
	*(*int32)(unsafe.Add(mBase, uint32(v23)+480)) = v1044
	*(*int32)(unsafe.Add(mBase, uint32(v23)+488)) = v1042
	*(*int32)(unsafe.Add(mBase, uint32(v23)+492)) = v1045
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+471)) = uint8(v1066)
	v1183 = v1053 & int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+487)) = uint8(v1183)
	F_list_free(m, v1043)
	mBase = m.M
	v1186 = m.ExcPending
	if v1186 != 0 {
		goto L7
	} else {
		goto L136
	}
L126:
	;
	v1067 = int32(0)
	v1068 = *(*int32)(unsafe.Add(mBase, uint32(v1047)))
	if v1068 <= v1067 {
		goto L125
	} else {
		goto L127
	}
L127:
	;
	v1082 = v1067
	goto L128
L128:
	;
	v1091 = *(*int32)(unsafe.Add(mBase, uint32(v1043)+12))
	v1095 = *(*int32)(unsafe.Add(mBase, uint32(v1091+v1082<<(uint(int32(2))%32))))
	v1096 = *(*int32)(unsafe.Add(mBase, uint32(v1095)))
	if v1096 == int32(0) {
		goto L130
	} else {
		goto L131
	}
L129:
	;
	goto L125
L130:
	;
	v1150 = v1082 + int32(1)
	v1151 = *(*int32)(unsafe.Add(mBase, uint32(v1047)))
	if v1150 < v1151 {
		v1082 = v1150
		goto L128
	} else {
		goto L135
	}
L131:
	;
	v1099 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1095)+16)))
	if v1099 == int32(115) {
		goto L130
	} else {
		goto L132
	}
L132:
	;
	v1102 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v23)+136)) = v1102
	*(*int64)(unsafe.Add(mBase, uint32(v23)+128)) = v1102
	*(*int64)(unsafe.Add(mBase, uint32(v23)+120)) = v1102
	*(*int64)(unsafe.Add(mBase, uint32(v23)+112)) = v1102
	*(*int64)(unsafe.Add(mBase, uint32(v23)+104)) = v1102
	*(*int64)(unsafe.Add(mBase, uint32(v23)+96)) = v1102
	*(*int64)(unsafe.Add(mBase, uint32(v23)+88)) = v1102
	*(*int64)(unsafe.Add(mBase, uint32(v23)+80)) = v1102
	*(*int32)(unsafe.Add(mBase, uint32(v23)+456)) = v1048
	*(*int32)(unsafe.Add(mBase, uint32(v23)+460)) = v1049
	*(*int32)(unsafe.Add(mBase, uint32(v23)+464)) = v1046
	*(*int32)(unsafe.Add(mBase, uint32(v23)+472)) = v1047
	*(*int32)(unsafe.Add(mBase, uint32(v23)+476)) = v1043
	*(*int32)(unsafe.Add(mBase, uint32(v23)+480)) = v1044
	*(*int32)(unsafe.Add(mBase, uint32(v23)+488)) = v1042
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+471)) = uint8(v1066)
	v1127 = v1053 & int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+487)) = uint8(v1127)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+492)) = v1045
	v1131 = v23 + int32(80)
	F_ReplicationSlotNameForTablesync(m, v1042, v1096, v1131)
	mBase = m.M
	v1133 = m.ExcPending
	if v1133 != 0 {
		goto L7
	} else {
		goto L133
	}
L133:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+460)) = v1049
	*(*int32)(unsafe.Add(mBase, uint32(v23)+456)) = v1048
	*(*int32)(unsafe.Add(mBase, uint32(v23)+464)) = v1046
	*(*int32)(unsafe.Add(mBase, uint32(v23)+472)) = v1047
	*(*int32)(unsafe.Add(mBase, uint32(v23)+476)) = v1043
	*(*int32)(unsafe.Add(mBase, uint32(v23)+480)) = v1044
	*(*int32)(unsafe.Add(mBase, uint32(v23)+488)) = v1042
	*(*int32)(unsafe.Add(mBase, uint32(v23)+492)) = v1045
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+471)) = uint8(v1066)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+487)) = uint8(v1127)
	F_ReplicationSlotDropAtPubNode(m, v1046, v1131, int32(1))
	mBase = m.M
	v1146 = m.ExcPending
	if v1146 != 0 {
		goto L7
	} else {
		goto L134
	}
L134:
	;
	goto L130
L135:
	;
	goto L129
L136:
	;
	if v1183 == int32(0) {
		goto L137
	} else {
		goto L138
	}
L137:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+460)) = v1049
	*(*int32)(unsafe.Add(mBase, uint32(v23)+456)) = v1048
	*(*int32)(unsafe.Add(mBase, uint32(v23)+464)) = v1046
	*(*int32)(unsafe.Add(mBase, uint32(v23)+472)) = v1047
	*(*int32)(unsafe.Add(mBase, uint32(v23)+476)) = v1043
	*(*int32)(unsafe.Add(mBase, uint32(v23)+480)) = v1044
	*(*int32)(unsafe.Add(mBase, uint32(v23)+488)) = v1042
	*(*int32)(unsafe.Add(mBase, uint32(v23)+492)) = v1045
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+471)) = uint8(v1066)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+487)) = uint8(v1183)
	F_ReplicationSlotDropAtPubNode(m, v1046, v1044, int32(0))
	mBase = m.M
	v1201 = m.ExcPending
	if v1201 != 0 {
		goto L7
	} else {
		goto L140
	}
L138:
	;
	goto L139
L139:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_DropSubscription[4])) = v1048
	*(*int32)(unsafe.Add(mBase, _c_F_DropSubscription[5])) = v1049
	v1207 = *(*int32)(unsafe.Add(mBase, _c_F_DropSubscription[3]))
	v1208 = *(*int32)(unsafe.Add(mBase, uint32(v1207)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+456)) = v1048
	*(*int32)(unsafe.Add(mBase, uint32(v23)+460)) = v1049
	*(*int32)(unsafe.Add(mBase, uint32(v23)+464)) = v1046
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+471)) = uint8(v1066)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+472)) = v1047
	*(*int32)(unsafe.Add(mBase, uint32(v23)+476)) = v1043
	*(*int32)(unsafe.Add(mBase, uint32(v23)+480)) = v1044
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+487)) = uint8(v1183)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+488)) = v1042
	*(*int32)(unsafe.Add(mBase, uint32(v23)+492)) = v1045
	m.T0[v1208].(func(*base.Module, int32))(m, v1046)
	mBase = m.M
	v1220 = m.ExcPending
	if v1220 != 0 {
		goto L7
	} else {
		goto L141
	}
L140:
	;
	goto L139
L141:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_DropSubscription[4])) = v1048
	*(*int32)(unsafe.Add(mBase, _c_F_DropSubscription[5])) = v1049
	*(*int32)(unsafe.Add(mBase, uint32(v23)+456)) = v1048
	*(*int32)(unsafe.Add(mBase, uint32(v23)+460)) = v1049
	*(*int32)(unsafe.Add(mBase, uint32(v23)+464)) = v1046
	*(*int32)(unsafe.Add(mBase, uint32(v23)+472)) = v1047
	*(*int32)(unsafe.Add(mBase, uint32(v23)+476)) = v1043
	*(*int32)(unsafe.Add(mBase, uint32(v23)+480)) = v1044
	*(*int32)(unsafe.Add(mBase, uint32(v23)+488)) = v1042
	*(*int32)(unsafe.Add(mBase, uint32(v23)+492)) = v1045
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+471)) = uint8(v1066)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+487)) = uint8(v1183)
	F_relation_close(m, v1045, int32(0))
	mBase = m.M
	v1237 = m.ExcPending
	if v1237 != 0 {
		goto L7
	} else {
		goto L142
	}
L142:
	;
	goto L4
L143:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+460)) = v1049
	*(*int32)(unsafe.Add(mBase, uint32(v23)+456)) = v1048
	*(*int32)(unsafe.Add(mBase, uint32(v23)+464)) = v1046
	*(*int32)(unsafe.Add(mBase, uint32(v23)+472)) = v1047
	*(*int32)(unsafe.Add(mBase, uint32(v23)+476)) = v1043
	*(*int32)(unsafe.Add(mBase, uint32(v23)+480)) = v1044
	*(*int32)(unsafe.Add(mBase, uint32(v23)+488)) = v1042
	*(*int32)(unsafe.Add(mBase, uint32(v23)+492)) = v1045
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+471)) = uint8(v1249)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+487)) = uint8(v1255)
	F_pg_re_throw(m)
	mBase = m.M
	v1272 = m.ExcPending
	if v1272 != 0 {
		goto L7
	} else {
		goto L144
	}
L144:
	;
	goto L1
L145:
	;
	v1298 = int32(v1294)
	m.G0 = v23
	v1300 = *(*int32)(unsafe.Add(mBase, uint32(v1298)+4))
	v1301 = *(*int32)(unsafe.Add(mBase, uint32(v1298)))
	v1304 = *(*int32)(unsafe.Add(mBase, uint32(v1301)))
	if v23+int32(76) == v1304 {
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
	if v1308 != 0 {
		goto L151
	} else {
		goto L152
	}
L148:
	;
	v1306 = *(*int32)(unsafe.Add(mBase, uint32(v1301)+4))
	v1308 = v1306
	goto L150
L149:
	;
	v1308 = int32(0)
	goto L150
L150:
	;
	goto L147
L151:
	;
	v1309 = *(*int32)(unsafe.Add(mBase, uint32(v23)+492))
	v1310 = *(*int32)(unsafe.Add(mBase, uint32(v23)+488))
	v1311 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+487)))
	v1312 = *(*int32)(unsafe.Add(mBase, uint32(v23)+480))
	v1313 = *(*int32)(unsafe.Add(mBase, uint32(v23)+476))
	v1314 = *(*int32)(unsafe.Add(mBase, uint32(v23)+472))
	v1315 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+471)))
	v1316 = *(*int32)(unsafe.Add(mBase, uint32(v23)+464))
	v1317 = *(*int32)(unsafe.Add(mBase, uint32(v23)+460))
	v1318 = *(*int32)(unsafe.Add(mBase, uint32(v23)+456))
	v29 = v1310
	v30 = v1313
	v31 = v1308
	v32 = v1309
	v33 = v1316
	v34 = v1314
	v35 = v1318
	v36 = v1317
	v37 = v1312
	v38 = v1315
	v39 = v1300
	v40 = v1311
	goto L2
L152:
	;
	goto L153
L153:
	;
	F___wasm_longjmp(m, v1301, v1300)
	mBase = m.M
	v1320 = m.ExcPending
	if v1320 != 0 {
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
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
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
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int64
	_ = v73
	var v74 int64
	_ = v74
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	v9 = m.G0
	v11 = v9 - int32(112)
	m.G0 = v11
	v13 = int32(1)
	v16 = F_table_open(m, int32(_a_F_GetSubscriptionRelations_0), v13)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	F_ScanKeyInit(m, v11+int32(16), int32(1), int32(3), int32(184), l0)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
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
	v29 = int32(3)
	F_ScanKeyInit(m, v11-int32(-64), v29, v29, int32(70), int32(114))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L7
	}
L5:
	;
	v36 = v13
	goto L6
L6:
	;
	v37 = int32(0)
	v43 = F_systable_beginscan(m, v16, v37, v37, v37, v36, v11+int32(16))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L1
	} else {
		goto L8
	}
L7:
	;
	v36 = int32(2)
	goto L6
L8:
	;
	v45 = F_systable_getnext(m, v43)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	if v45 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v47 = v45
	v48 = v37
	goto L13
L11:
	;
	v81 = v37
	goto L12
L12:
	;
	F_systable_endscan(m, v43)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L1
	} else {
		goto L23
	}
L13:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v47)+16))
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+22)))
	v58 = F_palloc(m, int32(24))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L1
	} else {
		goto L15
	}
L14:
	;
	v81 = v76
	goto L12
L15:
	;
	v60 = v55 + v56
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v58))) = v61
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+8)))
	*(*uint8)(unsafe.Add(mBase, uint32(v58)+16)) = uint8(v63)
	v69 = F_SysCacheGetAttr(m, int32(68), v47, int32(4), v11+int32(15))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+15)))
	if v71 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v74 = int64(0)
	goto L19
L18:
	;
	v73 = *(*int64)(unsafe.Add(mBase, uint32(v69)))
	v74 = v73
	goto L19
L19:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v58)+8)) = v74
	v76 = F_lappend(m, v48, v58)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	v78 = F_systable_getnext(m, v43)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	if v78 != 0 {
		v47 = v78
		v48 = v76
		goto L13
	} else {
		goto L22
	}
L22:
	;
	goto L14
L23:
	;
	F_relation_close(m, v16, int32(1))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	m.G0 = v11 + int32(112)
	return v81
}
func F_subscription_change_cb(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	v5 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_subscription_change_cb[0])) = uint8(v5)
	return
}
