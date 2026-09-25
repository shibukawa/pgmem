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
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v94 int64
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
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
					v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22))))
					if v35 == int32(18) {
						v38 = int32(16)
					} else {
						v38 = int32(0)
					}
					if base.Ui32((v35-int32(1))&int32(255)) < base.Ui32(int32(3)) {
						v45 = int32(4)
					} else {
						v45 = v38
					}
					v56 = v45
				} else {
					v46 = int32(1)
					if v28 != 0 {
						v56 = int32(base.Ui32(v26)>>(uint(v46)%32)) - v46
					} else {
						v50 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
						v56 = int32(base.Ui32(v50)>>(uint(int32(2))%32)) - int32(4)
					}
				}
				v57 = m.G0
				v59 = v57 - int32(16)
				m.G0 = v59
				v62 = base.B2i32(v10 != int32(0))
				if v10 != int32(0) {
					v63 = F_GetCurrentTransactionId(m)
					mBase = m.M
					v64 = m.ExcPending
					if v64 != 0 {
						return int32(0)
					} else {
						*(*uint8)(unsafe.Add(mBase, uint32(v59)+4)) = uint8(v62)
						v69 = *(*int32)(unsafe.Add(mBase, _c_F_pg_logical_emit_message_bytea[0]))
						*(*int32)(unsafe.Add(mBase, uint32(v59))) = v69
						v71 = F_strlen(m, v16)
						mBase = m.M
						*(*int32)(unsafe.Add(mBase, uint32(v59)+12)) = v56
						*(*int32)(unsafe.Add(mBase, uint32(v59)+8)) = v71 + int32(1)
						F_XLogBeginInsert(m)
						mBase = m.M
						v77 = m.ExcPending
						if v77 != 0 {
							return int32(0)
						} else {
							F_XLogRegisterData(m, v59, int32(16))
							mBase = m.M
							v80 = m.ExcPending
							if v80 != 0 {
								return int32(0)
							} else {
								v81 = *(*int32)(unsafe.Add(mBase, uint32(v59)+8))
								F_XLogRegisterData(m, v16, v81)
								mBase = m.M
								v83 = m.ExcPending
								if v83 != 0 {
									return int32(0)
								} else {
									F_XLogRegisterData(m, v29, v56)
									mBase = m.M
									v85 = m.ExcPending
									if v85 != 0 {
										return int32(0)
									} else {
										v87 = int32(_a_F_pg_logical_emit_message_bytea_0)
										v89 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_logical_emit_message_bytea[1])))
										v90 = v89 | int32(1)
										*(*uint8)(unsafe.Add(mBase, _c_F_pg_logical_emit_message_bytea[1])) = uint8(v90)
										v94 = F_XLogInsert(m, int32(21), int32(0))
										mBase = m.M
										v95 = m.ExcPending
										if v95 != 0 {
											return int32(0)
										} else {
											v96 = int32(0)
											if base.B2i32(base.B2i32(v23 != int32(0)) == v96)|v62 == v96 {
												F_XLogFlush(m, v94)
												mBase = m.M
												v102 = m.ExcPending
												if v102 != 0 {
													return int32(0)
												} else {
													m.G0 = v59 + int32(16)
													v106 = F_Int64GetDatum(m, v94)
													mBase = m.M
													v107 = m.ExcPending
													if v107 != 0 {
														return int32(0)
													} else {
														return v106
													}
												}
											} else {
												m.G0 = v59 + int32(16)
												v106 = F_Int64GetDatum(m, v94)
												mBase = m.M
												v107 = m.ExcPending
												if v107 != 0 {
													return int32(0)
												} else {
													return v106
												}
											}
										}
									}
								}
							}
						}
					}
				} else {
					*(*uint8)(unsafe.Add(mBase, uint32(v59)+4)) = uint8(v62)
					v69 = *(*int32)(unsafe.Add(mBase, _c_F_pg_logical_emit_message_bytea[0]))
					*(*int32)(unsafe.Add(mBase, uint32(v59))) = v69
					v71 = F_strlen(m, v16)
					mBase = m.M
					*(*int32)(unsafe.Add(mBase, uint32(v59)+12)) = v56
					*(*int32)(unsafe.Add(mBase, uint32(v59)+8)) = v71 + int32(1)
					F_XLogBeginInsert(m)
					mBase = m.M
					v77 = m.ExcPending
					if v77 != 0 {
						return int32(0)
					} else {
						F_XLogRegisterData(m, v59, int32(16))
						mBase = m.M
						v80 = m.ExcPending
						if v80 != 0 {
							return int32(0)
						} else {
							v81 = *(*int32)(unsafe.Add(mBase, uint32(v59)+8))
							F_XLogRegisterData(m, v16, v81)
							mBase = m.M
							v83 = m.ExcPending
							if v83 != 0 {
								return int32(0)
							} else {
								F_XLogRegisterData(m, v29, v56)
								mBase = m.M
								v85 = m.ExcPending
								if v85 != 0 {
									return int32(0)
								} else {
									v87 = int32(_a_F_pg_logical_emit_message_bytea_0)
									v89 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_logical_emit_message_bytea[1])))
									v90 = v89 | int32(1)
									*(*uint8)(unsafe.Add(mBase, _c_F_pg_logical_emit_message_bytea[1])) = uint8(v90)
									v94 = F_XLogInsert(m, int32(21), int32(0))
									mBase = m.M
									v95 = m.ExcPending
									if v95 != 0 {
										return int32(0)
									} else {
										v96 = int32(0)
										if base.B2i32(base.B2i32(v23 != int32(0)) == v96)|v62 == v96 {
											F_XLogFlush(m, v94)
											mBase = m.M
											v102 = m.ExcPending
											if v102 != 0 {
												return int32(0)
											} else {
												m.G0 = v59 + int32(16)
												v106 = F_Int64GetDatum(m, v94)
												mBase = m.M
												v107 = m.ExcPending
												if v107 != 0 {
													return int32(0)
												} else {
													return v106
												}
											}
										} else {
											m.G0 = v59 + int32(16)
											v106 = F_Int64GetDatum(m, v94)
											mBase = m.M
											v107 = m.ExcPending
											if v107 != 0 {
												return int32(0)
											} else {
												return v106
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
	var v19 int64
	_ = v19
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
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
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v48 int64
	_ = v48
	var v49 int64
	_ = v49
	var v50 int64
	_ = v50
	var v51 int64
	_ = v51
	var v56 int32
	_ = v56
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v105 int32
	_ = v105
	var v120 int32
	_ = v120
	var v136 int32
	_ = v136
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v159 int64
	_ = v159
	var v160 int64
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v185 int32
	_ = v185
	var v200 int32
	_ = v200
	var v216 int32
	_ = v216
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v290 int32
	_ = v290
	var v305 int32
	_ = v305
	var v321 int32
	_ = v321
	var v338 int32
	_ = v338
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v368 int32
	_ = v368
	var v383 int32
	_ = v383
	var v399 int32
	_ = v399
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v443 int32
	_ = v443
	var v455 int32
	_ = v455
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v474 int32
	_ = v474
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v491 int32
	_ = v491
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v562 int32
	_ = v562
	var v565 int32
	_ = v565
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v594 int32
	_ = v594
	var v610 int32
	_ = v610
	var v615 int32
	_ = v615
	var v616 int32
	_ = v616
	var v618 int32
	_ = v618
	var v620 int32
	_ = v620
	var v633 int32
	_ = v633
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v637 int64
	_ = v637
	var v640 int64
	_ = v640
	var v645 int32
	_ = v645
	var v648 int32
	_ = v648
	var v652 int64
	_ = v652
	var v659 int64
	_ = v659
	var v675 int32
	_ = v675
	var v690 int32
	_ = v690
	var v706 int32
	_ = v706
	var v723 int32
	_ = v723
	var v737 int64
	_ = v737
	var v738 int32
	_ = v738
	var v739 int64
	_ = v739
	var v740 int64
	_ = v740
	var v741 int64
	_ = v741
	var v754 int32
	_ = v754
	var v757 int32
	_ = v757
	var v760 int32
	_ = v760
	var v762 int32
	_ = v762
	var v764 int32
	_ = v764
	var v774 int32
	_ = v774
	var v775 int32
	_ = v775
	var v776 int32
	_ = v776
	var v777 int32
	_ = v777
	var v778 int32
	_ = v778
	var v779 int32
	_ = v779
	var v780 int32
	_ = v780
	var v782 int32
	_ = v782
	var v783 int32
	_ = v783
	var v788 int64
	_ = v788
	var v789 int64
	_ = v789
	var v790 int64
	_ = v790
	var v791 int64
	_ = v791
	var v819 int32
	_ = v819
	var v825 int32
	_ = v825
	var v826 int32
	_ = v826
	var v829 int32
	_ = v829
	var v847 int32
	_ = v847
	var v862 int32
	_ = v862
	var v863 int32
	_ = v863
	var v864 int32
	_ = v864
	var v878 int32
	_ = v878
	var v879 int32
	_ = v879
	var v880 int32
	_ = v880
	var v901 int32
	_ = v901
	var v918 int32
	_ = v918
	var v932 int64
	_ = v932
	var v934 int32
	_ = v934
	var v935 int64
	_ = v935
	var v937 int32
	_ = v937
	var v939 int32
	_ = v939
	var v941 int32
	_ = v941
	var v942 int64
	_ = v942
	var v956 int32
	_ = v956
	var v970 int32
	_ = v970
	var v971 int32
	_ = v971
	var v972 int64
	_ = v972
	var v988 int32
	_ = v988
	var v1015 int32
	_ = v1015
	var v1016 int32
	_ = v1016
	var v1017 int32
	_ = v1017
	var v1033 int32
	_ = v1033
	var v1046 int32
	_ = v1046
	var v1050 int32
	_ = v1050
	var v1067 int32
	_ = v1067
	var v1068 int32
	_ = v1068
	var v1082 int32
	_ = v1082
	var v1086 int32
	_ = v1086
	var v1087 int64
	_ = v1087
	var v1092 int64
	_ = v1092
	var v1094 int32
	_ = v1094
	var v1095 int64
	_ = v1095
	var v1097 int32
	_ = v1097
	var v1111 int32
	_ = v1111
	var v1112 int32
	_ = v1112
	var v1113 int64
	_ = v1113
	var v1137 int64
	_ = v1137
	var v1141 int32
	_ = v1141
	var v1161 int32
	_ = v1161
	var v1175 int32
	_ = v1175
	var v1189 int32
	_ = v1189
	var v1203 int32
	_ = v1203
	var v1217 int32
	_ = v1217
	var v1242 int32
	_ = v1242
	var v1256 int32
	_ = v1256
	var v1281 int32
	_ = v1281
	var v1282 int64
	_ = v1282
	var v1286 int32
	_ = v1286
	var v1288 int32
	_ = v1288
	var v1289 int32
	_ = v1289
	var v1292 int32
	_ = v1292
	var v1294 int32
	_ = v1294
	var v1296 int32
	_ = v1296
	var v1297 int32
	_ = v1297
	var v1298 int64
	_ = v1298
	var v1299 int32
	_ = v1299
	var v1300 int32
	_ = v1300
	var v1301 int32
	_ = v1301
	var v1302 int32
	_ = v1302
	var v1303 int32
	_ = v1303
	var v1304 int64
	_ = v1304
	var v1305 int64
	_ = v1305
	var v1306 int64
	_ = v1306
	var v1307 int32
	_ = v1307
	var v1308 int32
	_ = v1308
	var v1310 int32
	_ = v1310
	v3 = l2
	v4 = int32(0)
	v19 = int64(0)
	v25 = m.G0
	v27 = v25 - int32(288)
	m.G0 = v27
	v34 = v4
	v35 = v4
	v36 = v4
	v37 = v4
	v38 = v4
	v39 = v4
	v40 = v4
	v41 = v4
	v42 = int32(-1)
	v43 = v4
	v48 = v19
	v49 = v19
	v50 = v19
	v51 = v19
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
	if v42 == int32(1) {
		goto L8
	} else {
		goto L9
	}
L5:
	;
	goto L3
L6:
	;
	v1281 = int32(m.ExcTag)
	v1282 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v1281 == int32(0) {
		goto L144
	} else {
		goto L145
	}
L7:
	;
	if v783 == int32(0) {
		goto L86
	} else {
		goto L87
	}
L8:
	;
	v774 = v34
	v775 = v35
	v776 = v36
	v777 = v37
	v778 = v38
	v779 = v39
	v780 = v40
	v782 = v41
	v783 = v43
	v788 = v48
	v789 = v49
	v790 = v50
	v791 = v51
	goto L7
L9:
	;
	goto L10
L10:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+216)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v27)+220)) = v37
	*(*int64)(unsafe.Add(mBase, uint32(v27)+224)) = v49
	*(*int64)(unsafe.Add(mBase, uint32(v27)+232)) = v50
	*(*int64)(unsafe.Add(mBase, uint32(v27)+240)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v27)+252)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v27)+256)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v27)+260)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v27)+264)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v27)+268)) = v35
	*(*int64)(unsafe.Add(mBase, uint32(v27)+272)) = v48
	v69 = *(*int32)(unsafe.Add(mBase, _c_F_pg_logical_slot_get_changes_guts[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+284)) = v69
	F_CheckSlotPermissions(m)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L6
	} else {
		goto L11
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+220)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v27)+216)) = v36
	*(*int64)(unsafe.Add(mBase, uint32(v27)+224)) = v49
	*(*int64)(unsafe.Add(mBase, uint32(v27)+232)) = v50
	*(*int64)(unsafe.Add(mBase, uint32(v27)+240)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v27)+252)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v27)+256)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v27)+260)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v27)+264)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v27)+268)) = v35
	*(*int64)(unsafe.Add(mBase, uint32(v27)+272)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v27)+284)) = v69
	F_CheckLogicalDecodingRequirements(m)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L6
	} else {
		goto L12
	}
L12:
	;
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v87 == int32(1) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+220)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v27)+216)) = v36
	*(*int64)(unsafe.Add(mBase, uint32(v27)+224)) = v49
	*(*int64)(unsafe.Add(mBase, uint32(v27)+232)) = v50
	*(*int64)(unsafe.Add(mBase, uint32(v27)+240)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v27)+252)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v27)+256)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v27)+260)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v27)+264)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v27)+268)) = v35
	*(*int64)(unsafe.Add(mBase, uint32(v27)+272)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v27)+284)) = v69
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L6
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v155 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+220)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v27)+216)) = v36
	*(*int64)(unsafe.Add(mBase, uint32(v27)+224)) = v49
	*(*int64)(unsafe.Add(mBase, uint32(v27)+232)) = v50
	*(*int64)(unsafe.Add(mBase, uint32(v27)+240)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v27)+252)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v27)+256)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v27)+260)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v27)+264)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v27)+268)) = v35
	*(*int64)(unsafe.Add(mBase, uint32(v27)+272)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v27)+284)) = v69
	F_errcode(m, int32(67108994))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L6
	} else {
		goto L17
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+220)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v27)+216)) = v36
	*(*int64)(unsafe.Add(mBase, uint32(v27)+224)) = v49
	*(*int64)(unsafe.Add(mBase, uint32(v27)+232)) = v50
	*(*int64)(unsafe.Add(mBase, uint32(v27)+240)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v27)+252)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v27)+256)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v27)+260)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v27)+264)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v27)+268)) = v35
	*(*int64)(unsafe.Add(mBase, uint32(v27)+272)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v27)+284)) = v69
	F_errmsg(m, int32(_a_F_pg_logical_slot_get_changes_guts_0), int32(0))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L6
	} else {
		goto L18
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+220)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v27)+216)) = v36
	*(*int64)(unsafe.Add(mBase, uint32(v27)+224)) = v49
	*(*int64)(unsafe.Add(mBase, uint32(v27)+232)) = v50
	*(*int64)(unsafe.Add(mBase, uint32(v27)+240)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v27)+252)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v27)+256)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v27)+260)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v27)+264)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v27)+268)) = v35
	*(*int64)(unsafe.Add(mBase, uint32(v27)+272)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v27)+284)) = v69
	F_errfinish(m, int32(_a_F_pg_logical_slot_get_changes_guts_1), int32(123), int32(_a_F_pg_logical_slot_get_changes_guts_2))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L6
	} else {
		goto L19
	}
L19:
	;
	goto L3
L20:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v159 = *(*int64)(unsafe.Add(mBase, uint32(v158)))
	v160 = v159
	goto L22
L21:
	;
	v160 = int64(0)
	goto L22
L22:
	;
	v161 = int32(0)
	v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+40)))
	if v162 == v161 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v166 = v165
	goto L25
L24:
	;
	v166 = v161
	goto L25
L25:
	;
	v167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)))
	if v167 == int32(1) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+220)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v27)+216)) = v36
	*(*int64)(unsafe.Add(mBase, uint32(v27)+224)) = v49
	*(*int64)(unsafe.Add(mBase, uint32(v27)+232)) = v50
	*(*int64)(unsafe.Add(mBase, uint32(v27)+240)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v27)+252)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v27)+256)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v27)+260)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v27)+264)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v27)+268)) = v166
	*(*int64)(unsafe.Add(mBase, uint32(v27)+272)) = v160
	*(*int32)(unsafe.Add(mBase, uint32(v27)+284)) = v69
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L6
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	v234 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v235 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+220)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v27)+216)) = v36
	*(*int64)(unsafe.Add(mBase, uint32(v27)+224)) = v49
	*(*int64)(unsafe.Add(mBase, uint32(v27)+232)) = v50
	*(*int64)(unsafe.Add(mBase, uint32(v27)+240)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v27)+252)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v27)+256)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v27)+260)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v27)+264)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v27)+268)) = v166
	*(*int64)(unsafe.Add(mBase, uint32(v27)+272)) = v160
	*(*int32)(unsafe.Add(mBase, uint32(v27)+284)) = v69
	v248 = F_pg_detoast_datum(m, v235)
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L6
	} else {
		goto L33
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+220)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v27)+216)) = v36
	*(*int64)(unsafe.Add(mBase, uint32(v27)+224)) = v49
	*(*int64)(unsafe.Add(mBase, uint32(v27)+232)) = v50
	*(*int64)(unsafe.Add(mBase, uint32(v27)+240)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v27)+252)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v27)+256)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v27)+260)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v27)+264)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v27)+268)) = v166
	*(*int64)(unsafe.Add(mBase, uint32(v27)+272)) = v160
	*(*int32)(unsafe.Add(mBase, uint32(v27)+284)) = v69
	F_errcode(m, int32(67108994))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L6
	} else {
		goto L30
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+220)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v27)+216)) = v36
	*(*int64)(unsafe.Add(mBase, uint32(v27)+224)) = v49
	*(*int64)(unsafe.Add(mBase, uint32(v27)+232)) = v50
	*(*int64)(unsafe.Add(mBase, uint32(v27)+240)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v27)+252)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v27)+256)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v27)+260)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v27)+264)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v27)+268)) = v166
	*(*int64)(unsafe.Add(mBase, uint32(v27)+272)) = v160
	*(*int32)(unsafe.Add(mBase, uint32(v27)+284)) = v69
	F_errmsg(m, int32(_a_F_pg_logical_slot_get_changes_guts_3), int32(0))
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L6
	} else {
		goto L31
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+220)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v27)+216)) = v36
	*(*int64)(unsafe.Add(mBase, uint32(v27)+224)) = v49
	*(*int64)(unsafe.Add(mBase, uint32(v27)+232)) = v50
	*(*int64)(unsafe.Add(mBase, uint32(v27)+240)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v27)+252)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v27)+256)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v27)+260)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v27)+264)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v27)+268)) = v166
	*(*int64)(unsafe.Add(mBase, uint32(v27)+272)) = v160
	*(*int32)(unsafe.Add(mBase, uint32(v27)+284)) = v69
	F_errfinish(m, int32(_a_F_pg_logical_slot_get_changes_guts_1), int32(139), int32(_a_F_pg_logical_slot_get_changes_guts_2))
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L6
	} else {
		goto L32
	}
L32:
	;
	goto L3
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+220)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v27)+216)) = v36
	*(*int64)(unsafe.Add(mBase, uint32(v27)+224)) = v49
	*(*int64)(unsafe.Add(mBase, uint32(v27)+232)) = v50
	*(*int64)(unsafe.Add(mBase, uint32(v27)+240)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v27)+252)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v27)+256)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v27)+260)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v27)+264)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v27)+268)) = v166
	*(*int64)(unsafe.Add(mBase, uint32(v27)+272)) = v160
	*(*int32)(unsafe.Add(mBase, uint32(v27)+284)) = v69
	v263 = F_palloc0(m, int32(24))
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L6
	} else {
		goto L34
	}
L34:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v263)+8)) = uint8(v3)
	v266 = int32(_a_F_pg_logical_slot_get_changes_guts_4)
	v267 = *(*int32)(unsafe.Add(mBase, _c_F_pg_logical_slot_get_changes_guts[1]))
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v56)+4))
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v269)+16))
	*(*int32)(unsafe.Add(mBase, _c_F_pg_logical_slot_get_changes_guts[1])) = v270
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v248)+4))
	if base.Ui32(int32(2)) <= base.Ui32(v272) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+220)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v27)+216)) = v36
	*(*int64)(unsafe.Add(mBase, uint32(v27)+224)) = v49
	*(*int64)(unsafe.Add(mBase, uint32(v27)+232)) = v50
	*(*int64)(unsafe.Add(mBase, uint32(v27)+240)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v27)+252)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v27)+256)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v27)+260)) = v267
	*(*int32)(unsafe.Add(mBase, uint32(v27)+264)) = v263
	*(*int32)(unsafe.Add(mBase, uint32(v27)+268)) = v166
	*(*int64)(unsafe.Add(mBase, uint32(v27)+272)) = v160
	*(*int32)(unsafe.Add(mBase, uint32(v27)+284)) = v69
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L6
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+220)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v27)+216)) = v36
	*(*int64)(unsafe.Add(mBase, uint32(v27)+224)) = v49
	*(*int64)(unsafe.Add(mBase, uint32(v27)+232)) = v50
	*(*int64)(unsafe.Add(mBase, uint32(v27)+240)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v27)+252)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v27)+256)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v27)+260)) = v267
	*(*int32)(unsafe.Add(mBase, uint32(v27)+264)) = v263
	*(*int32)(unsafe.Add(mBase, uint32(v27)+268)) = v166
	*(*int64)(unsafe.Add(mBase, uint32(v27)+272)) = v160
	*(*int32)(unsafe.Add(mBase, uint32(v27)+284)) = v69
	v351 = F_array_contains_nulls(m, v248)
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L6
	} else {
		goto L42
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+220)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v27)+216)) = v36
	*(*int64)(unsafe.Add(mBase, uint32(v27)+224)) = v49
	*(*int64)(unsafe.Add(mBase, uint32(v27)+232)) = v50
	*(*int64)(unsafe.Add(mBase, uint32(v27)+240)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v27)+252)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v27)+256)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v27)+260)) = v267
	*(*int32)(unsafe.Add(mBase, uint32(v27)+264)) = v263
	*(*int32)(unsafe.Add(mBase, uint32(v27)+268)) = v166
	*(*int64)(unsafe.Add(mBase, uint32(v27)+272)) = v160
	*(*int32)(unsafe.Add(mBase, uint32(v27)+284)) = v69
	F_errcode(m, int32(1088))
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L6
	} else {
		goto L39
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+220)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v27)+216)) = v36
	*(*int64)(unsafe.Add(mBase, uint32(v27)+224)) = v49
	*(*int64)(unsafe.Add(mBase, uint32(v27)+232)) = v50
	*(*int64)(unsafe.Add(mBase, uint32(v27)+240)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v27)+252)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v27)+256)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v27)+260)) = v267
	*(*int32)(unsafe.Add(mBase, uint32(v27)+264)) = v263
	*(*int32)(unsafe.Add(mBase, uint32(v27)+268)) = v166
	*(*int64)(unsafe.Add(mBase, uint32(v27)+272)) = v160
	*(*int32)(unsafe.Add(mBase, uint32(v27)+284)) = v69
	F_errmsg(m, int32(_a_F_pg_logical_slot_get_changes_guts_5), int32(0))
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L6
	} else {
		goto L40
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+220)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v27)+216)) = v36
	*(*int64)(unsafe.Add(mBase, uint32(v27)+224)) = v49
	*(*int64)(unsafe.Add(mBase, uint32(v27)+232)) = v50
	*(*int64)(unsafe.Add(mBase, uint32(v27)+240)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v27)+252)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v27)+256)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v27)+260)) = v267
	*(*int32)(unsafe.Add(mBase, uint32(v27)+264)) = v263
	*(*int32)(unsafe.Add(mBase, uint32(v27)+268)) = v166
	*(*int64)(unsafe.Add(mBase, uint32(v27)+272)) = v160
	*(*int32)(unsafe.Add(mBase, uint32(v27)+284)) = v69
	F_errfinish(m, int32(_a_F_pg_logical_slot_get_changes_guts_1), int32(156), int32(_a_F_pg_logical_slot_get_changes_guts_2))
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L6
	} else {
		goto L41
	}
L41:
	;
	goto L3
L42:
	;
	if v351 != 0 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+220)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v27)+216)) = v36
	*(*int64)(unsafe.Add(mBase, uint32(v27)+224)) = v49
	*(*int64)(unsafe.Add(mBase, uint32(v27)+232)) = v50
	*(*int64)(unsafe.Add(mBase, uint32(v27)+240)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v27)+252)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v27)+256)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v27)+260)) = v267
	*(*int32)(unsafe.Add(mBase, uint32(v27)+264)) = v263
	*(*int32)(unsafe.Add(mBase, uint32(v27)+268)) = v166
	*(*int64)(unsafe.Add(mBase, uint32(v27)+272)) = v160
	*(*int32)(unsafe.Add(mBase, uint32(v27)+284)) = v69
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L6
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	v417 = int32(0)
	if v272 != int32(1) {
		v562 = v39
		v565 = v417
		goto L53
	} else {
		goto L54
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+220)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v27)+216)) = v36
	*(*int64)(unsafe.Add(mBase, uint32(v27)+224)) = v49
	*(*int64)(unsafe.Add(mBase, uint32(v27)+232)) = v50
	*(*int64)(unsafe.Add(mBase, uint32(v27)+240)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v27)+252)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v27)+256)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v27)+260)) = v267
	*(*int32)(unsafe.Add(mBase, uint32(v27)+264)) = v263
	*(*int32)(unsafe.Add(mBase, uint32(v27)+268)) = v166
	*(*int64)(unsafe.Add(mBase, uint32(v27)+272)) = v160
	*(*int32)(unsafe.Add(mBase, uint32(v27)+284)) = v69
	F_errcode(m, int32(1088))
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L6
	} else {
		goto L47
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+220)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v27)+216)) = v36
	*(*int64)(unsafe.Add(mBase, uint32(v27)+224)) = v49
	*(*int64)(unsafe.Add(mBase, uint32(v27)+232)) = v50
	*(*int64)(unsafe.Add(mBase, uint32(v27)+240)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v27)+252)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v27)+256)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v27)+260)) = v267
	*(*int32)(unsafe.Add(mBase, uint32(v27)+264)) = v263
	*(*int32)(unsafe.Add(mBase, uint32(v27)+268)) = v166
	*(*int64)(unsafe.Add(mBase, uint32(v27)+272)) = v160
	*(*int32)(unsafe.Add(mBase, uint32(v27)+284)) = v69
	F_errmsg(m, int32(_a_F_pg_logical_slot_get_changes_guts_6), int32(0))
	mBase = m.M
	v399 = m.ExcPending
	if v399 != 0 {
		goto L6
	} else {
		goto L48
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+220)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v27)+216)) = v36
	*(*int64)(unsafe.Add(mBase, uint32(v27)+224)) = v49
	*(*int64)(unsafe.Add(mBase, uint32(v27)+232)) = v50
	*(*int64)(unsafe.Add(mBase, uint32(v27)+240)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v27)+252)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v27)+256)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v27)+260)) = v267
	*(*int32)(unsafe.Add(mBase, uint32(v27)+264)) = v263
	*(*int32)(unsafe.Add(mBase, uint32(v27)+268)) = v166
	*(*int64)(unsafe.Add(mBase, uint32(v27)+272)) = v160
	*(*int32)(unsafe.Add(mBase, uint32(v27)+284)) = v69
	F_errfinish(m, int32(_a_F_pg_logical_slot_get_changes_guts_1), int32(162), int32(_a_F_pg_logical_slot_get_changes_guts_2))
	mBase = m.M
	v416 = m.ExcPending
	if v416 != 0 {
		goto L6
	} else {
		goto L49
	}
L49:
	;
	goto L3
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+220)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v27)+216)) = v36
	*(*int64)(unsafe.Add(mBase, uint32(v27)+224)) = v741
	*(*int64)(unsafe.Add(mBase, uint32(v27)+232)) = v739
	*(*int64)(unsafe.Add(mBase, uint32(v27)+240)) = v740
	*(*int32)(unsafe.Add(mBase, uint32(v27)+252)) = v565
	*(*int32)(unsafe.Add(mBase, uint32(v27)+256)) = v562
	*(*int32)(unsafe.Add(mBase, uint32(v27)+260)) = v267
	*(*int32)(unsafe.Add(mBase, uint32(v27)+264)) = v263
	*(*int32)(unsafe.Add(mBase, uint32(v27)+268)) = v166
	*(*int64)(unsafe.Add(mBase, uint32(v27)+272)) = v160
	*(*int32)(unsafe.Add(mBase, uint32(v27)+284)) = v69
	v754 = int32(1)
	F_ReplicationSlotAcquire(m, v234, v754, v754)
	mBase = m.M
	v757 = m.ExcPending
	if v757 != 0 {
		goto L6
	} else {
		goto L81
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+220)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v27)+216)) = v36
	*(*int64)(unsafe.Add(mBase, uint32(v27)+224)) = v49
	*(*int64)(unsafe.Add(mBase, uint32(v27)+232)) = v50
	*(*int64)(unsafe.Add(mBase, uint32(v27)+240)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v27)+252)) = v565
	*(*int32)(unsafe.Add(mBase, uint32(v27)+256)) = v562
	*(*int32)(unsafe.Add(mBase, uint32(v27)+260)) = v267
	*(*int32)(unsafe.Add(mBase, uint32(v27)+264)) = v263
	*(*int32)(unsafe.Add(mBase, uint32(v27)+268)) = v166
	*(*int64)(unsafe.Add(mBase, uint32(v27)+272)) = v160
	*(*int32)(unsafe.Add(mBase, uint32(v27)+284)) = v69
	v737 = F_GetXLogReplayRecPtr(m, int32(0))
	mBase = m.M
	v738 = m.ExcPending
	if v738 != 0 {
		goto L6
	} else {
		goto L80
	}
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+220)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v27)+216)) = v36
	*(*int64)(unsafe.Add(mBase, uint32(v27)+224)) = v49
	*(*int64)(unsafe.Add(mBase, uint32(v27)+232)) = v50
	*(*int64)(unsafe.Add(mBase, uint32(v27)+240)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v27)+252)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v27)+256)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v27)+260)) = v267
	*(*int32)(unsafe.Add(mBase, uint32(v27)+264)) = v263
	*(*int32)(unsafe.Add(mBase, uint32(v27)+268)) = v166
	*(*int64)(unsafe.Add(mBase, uint32(v27)+272)) = v160
	*(*int32)(unsafe.Add(mBase, uint32(v27)+284)) = v69
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v675 = m.ExcPending
	if v675 != 0 {
		goto L6
	} else {
		goto L76
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+220)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v27)+216)) = v36
	*(*int64)(unsafe.Add(mBase, uint32(v27)+224)) = v49
	*(*int64)(unsafe.Add(mBase, uint32(v27)+232)) = v50
	*(*int64)(unsafe.Add(mBase, uint32(v27)+240)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v27)+252)) = v565
	*(*int32)(unsafe.Add(mBase, uint32(v27)+256)) = v562
	*(*int32)(unsafe.Add(mBase, uint32(v27)+260)) = v267
	*(*int32)(unsafe.Add(mBase, uint32(v27)+264)) = v263
	*(*int32)(unsafe.Add(mBase, uint32(v27)+268)) = v166
	*(*int64)(unsafe.Add(mBase, uint32(v27)+272)) = v160
	*(*int32)(unsafe.Add(mBase, uint32(v27)+284)) = v69
	F_InitMaterializedSRF(m, l0, int32(0))
	mBase = m.M
	v591 = m.ExcPending
	if v591 != 0 {
		goto L6
	} else {
		goto L66
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+216)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v27)+220)) = v37
	*(*int64)(unsafe.Add(mBase, uint32(v27)+224)) = v49
	*(*int64)(unsafe.Add(mBase, uint32(v27)+232)) = v50
	*(*int64)(unsafe.Add(mBase, uint32(v27)+240)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v27)+252)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v27)+256)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v27)+260)) = v267
	*(*int32)(unsafe.Add(mBase, uint32(v27)+264)) = v263
	*(*int32)(unsafe.Add(mBase, uint32(v27)+268)) = v166
	*(*int64)(unsafe.Add(mBase, uint32(v27)+272)) = v160
	*(*int32)(unsafe.Add(mBase, uint32(v27)+284)) = v69
	F_deconstruct_array_builtin(m, v248, int32(25), v27+int32(208), int32(0), v27+int32(212))
	mBase = m.M
	v439 = m.ExcPending
	if v439 != 0 {
		goto L6
	} else {
		goto L55
	}
L55:
	;
	v440 = *(*int32)(unsafe.Add(mBase, uint32(v27)+212))
	if v440&int32(1) != 0 {
		goto L52
	} else {
		goto L56
	}
L56:
	;
	v443 = int32(0)
	if v440 <= v443 {
		v562 = v39
		v565 = v417
		goto L53
	} else {
		goto L57
	}
L57:
	;
	v455 = v39
	v458 = v417
	v459 = v443
	goto L58
L58:
	;
	v471 = v459 << (uint(int32(2)) % 32)
	v472 = *(*int32)(unsafe.Add(mBase, uint32(v27)+208))
	v474 = *(*int32)(unsafe.Add(mBase, uint32(v471+v472)))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+216)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v27)+220)) = v37
	*(*int64)(unsafe.Add(mBase, uint32(v27)+224)) = v49
	*(*int64)(unsafe.Add(mBase, uint32(v27)+232)) = v50
	*(*int64)(unsafe.Add(mBase, uint32(v27)+240)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v27)+252)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v27)+256)) = v455
	*(*int32)(unsafe.Add(mBase, uint32(v27)+260)) = v267
	*(*int32)(unsafe.Add(mBase, uint32(v27)+264)) = v263
	*(*int32)(unsafe.Add(mBase, uint32(v27)+268)) = v166
	*(*int64)(unsafe.Add(mBase, uint32(v27)+272)) = v160
	*(*int32)(unsafe.Add(mBase, uint32(v27)+284)) = v69
	v487 = F_text_to_cstring(m, v474)
	mBase = m.M
	v488 = m.ExcPending
	if v488 != 0 {
		goto L6
	} else {
		goto L60
	}
L59:
	;
	v562 = v547
	v565 = v547
	goto L53
L60:
	;
	v489 = *(*int32)(unsafe.Add(mBase, uint32(v27)+208))
	v491 = *(*int32)(unsafe.Add(mBase, uint32(v489+v471)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+216)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v27)+220)) = v37
	*(*int64)(unsafe.Add(mBase, uint32(v27)+224)) = v49
	*(*int64)(unsafe.Add(mBase, uint32(v27)+232)) = v50
	*(*int64)(unsafe.Add(mBase, uint32(v27)+240)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v27)+252)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v27)+256)) = v455
	*(*int32)(unsafe.Add(mBase, uint32(v27)+260)) = v267
	*(*int32)(unsafe.Add(mBase, uint32(v27)+264)) = v263
	*(*int32)(unsafe.Add(mBase, uint32(v27)+268)) = v166
	*(*int64)(unsafe.Add(mBase, uint32(v27)+272)) = v160
	*(*int32)(unsafe.Add(mBase, uint32(v27)+284)) = v69
	v504 = F_text_to_cstring(m, v491)
	mBase = m.M
	v505 = m.ExcPending
	if v505 != 0 {
		goto L6
	} else {
		goto L61
	}
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+220)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v27)+216)) = v36
	*(*int64)(unsafe.Add(mBase, uint32(v27)+224)) = v49
	*(*int64)(unsafe.Add(mBase, uint32(v27)+232)) = v50
	*(*int64)(unsafe.Add(mBase, uint32(v27)+240)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v27)+252)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v27)+256)) = v455
	*(*int32)(unsafe.Add(mBase, uint32(v27)+260)) = v267
	*(*int32)(unsafe.Add(mBase, uint32(v27)+264)) = v263
	*(*int32)(unsafe.Add(mBase, uint32(v27)+268)) = v166
	*(*int64)(unsafe.Add(mBase, uint32(v27)+272)) = v160
	*(*int32)(unsafe.Add(mBase, uint32(v27)+284)) = v69
	v518 = F_makeString(m, v504)
	mBase = m.M
	v519 = m.ExcPending
	if v519 != 0 {
		goto L6
	} else {
		goto L62
	}
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+220)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v27)+216)) = v36
	*(*int64)(unsafe.Add(mBase, uint32(v27)+224)) = v49
	*(*int64)(unsafe.Add(mBase, uint32(v27)+232)) = v50
	*(*int64)(unsafe.Add(mBase, uint32(v27)+240)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v27)+252)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v27)+256)) = v455
	*(*int32)(unsafe.Add(mBase, uint32(v27)+260)) = v267
	*(*int32)(unsafe.Add(mBase, uint32(v27)+264)) = v263
	*(*int32)(unsafe.Add(mBase, uint32(v27)+268)) = v166
	*(*int64)(unsafe.Add(mBase, uint32(v27)+272)) = v160
	*(*int32)(unsafe.Add(mBase, uint32(v27)+284)) = v69
	v533 = F_makeDefElem(m, v487, v518, int32(-1))
	mBase = m.M
	v534 = m.ExcPending
	if v534 != 0 {
		goto L6
	} else {
		goto L63
	}
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+220)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v27)+216)) = v36
	*(*int64)(unsafe.Add(mBase, uint32(v27)+224)) = v49
	*(*int64)(unsafe.Add(mBase, uint32(v27)+232)) = v50
	*(*int64)(unsafe.Add(mBase, uint32(v27)+240)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v27)+252)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v27)+256)) = v455
	*(*int32)(unsafe.Add(mBase, uint32(v27)+260)) = v267
	*(*int32)(unsafe.Add(mBase, uint32(v27)+264)) = v263
	*(*int32)(unsafe.Add(mBase, uint32(v27)+268)) = v166
	*(*int64)(unsafe.Add(mBase, uint32(v27)+272)) = v160
	*(*int32)(unsafe.Add(mBase, uint32(v27)+284)) = v69
	v547 = F_lappend(m, v458, v533)
	mBase = m.M
	v548 = m.ExcPending
	if v548 != 0 {
		goto L6
	} else {
		goto L64
	}
L64:
	;
	v550 = v459 + int32(2)
	v551 = *(*int32)(unsafe.Add(mBase, uint32(v27)+212))
	if v550 < v551 {
		v455 = v547
		v458 = v547
		v459 = v550
		goto L58
	} else {
		goto L65
	}
L65:
	;
	goto L59
L66:
	;
	v592 = *(*int32)(unsafe.Add(mBase, uint32(v56)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v263))) = v592
	v594 = *(*int32)(unsafe.Add(mBase, uint32(v56)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v263)+4)) = v594
	*(*int32)(unsafe.Add(mBase, uint32(v27)+220)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v27)+216)) = v36
	*(*int64)(unsafe.Add(mBase, uint32(v27)+224)) = v49
	*(*int64)(unsafe.Add(mBase, uint32(v27)+232)) = v50
	*(*int64)(unsafe.Add(mBase, uint32(v27)+240)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v27)+252)) = v565
	*(*int32)(unsafe.Add(mBase, uint32(v27)+256)) = v562
	*(*int32)(unsafe.Add(mBase, uint32(v27)+260)) = v267
	*(*int32)(unsafe.Add(mBase, uint32(v27)+264)) = v263
	*(*int32)(unsafe.Add(mBase, uint32(v27)+268)) = v166
	*(*int64)(unsafe.Add(mBase, uint32(v27)+272)) = v160
	*(*int32)(unsafe.Add(mBase, uint32(v27)+284)) = v69
	v610 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_logical_slot_get_changes_guts[2])))
	if v610 == int32(1) {
		goto L68
	} else {
		goto L69
	}
L67:
	;
	if v620 != 0 {
		goto L51
	} else {
		goto L71
	}
L68:
	;
	v615 = *(*int32)(unsafe.Add(mBase, _c_F_pg_logical_slot_get_changes_guts[3]))
	v616 = *(*int32)(unsafe.Add(mBase, uint32(v615)+316))
	v618 = base.B2i32(v616 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _c_F_pg_logical_slot_get_changes_guts[2])) = uint8(v618)
	v620 = v618
	goto L70
L69:
	;
	v620 = int32(0)
	goto L70
L70:
	;
	goto L67
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+220)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v27)+216)) = v36
	*(*int64)(unsafe.Add(mBase, uint32(v27)+224)) = v49
	*(*int64)(unsafe.Add(mBase, uint32(v27)+232)) = v50
	*(*int64)(unsafe.Add(mBase, uint32(v27)+240)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v27)+252)) = v565
	*(*int32)(unsafe.Add(mBase, uint32(v27)+256)) = v562
	*(*int32)(unsafe.Add(mBase, uint32(v27)+260)) = v267
	*(*int32)(unsafe.Add(mBase, uint32(v27)+264)) = v263
	*(*int32)(unsafe.Add(mBase, uint32(v27)+268)) = v166
	*(*int64)(unsafe.Add(mBase, uint32(v27)+272)) = v160
	*(*int32)(unsafe.Add(mBase, uint32(v27)+284)) = v69
	v633 = int32(0)
	v635 = int32(_a_F_pg_logical_slot_get_changes_guts_7)
	v636 = *(*int32)(unsafe.Add(mBase, _c_F_pg_logical_slot_get_changes_guts[3]))
	v637 = int64(0)
	v640 = base.AtomicRmwCmpxchg64(m, v636, int32(280), v637, v637)
	*(*int64)(unsafe.Add(mBase, _c_F_pg_logical_slot_get_changes_guts[4])) = v640
	v645 = base.AtomicRmwOr32(m, v633, int32(_a_F_pg_logical_slot_get_changes_guts_8), v633)
	v648 = *(*int32)(unsafe.Add(mBase, _c_F_pg_logical_slot_get_changes_guts[3]))
	v652 = base.AtomicRmwCmpxchg64(m, v648, int32(272), v637, v637)
	*(*int64)(unsafe.Add(mBase, _c_F_pg_logical_slot_get_changes_guts[5])) = v652
	goto L74
L72:
	;
	v739 = v50
	v740 = v659
	v741 = v659
	goto L50
L74:
	;
	goto L75
L75:
	;
	v659 = *(*int64)(unsafe.Add(mBase, _c_F_pg_logical_slot_get_changes_guts[4]))
	goto L72
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+220)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v27)+216)) = v36
	*(*int64)(unsafe.Add(mBase, uint32(v27)+224)) = v49
	*(*int64)(unsafe.Add(mBase, uint32(v27)+232)) = v50
	*(*int64)(unsafe.Add(mBase, uint32(v27)+240)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v27)+252)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v27)+256)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v27)+260)) = v267
	*(*int32)(unsafe.Add(mBase, uint32(v27)+264)) = v263
	*(*int32)(unsafe.Add(mBase, uint32(v27)+268)) = v166
	*(*int64)(unsafe.Add(mBase, uint32(v27)+272)) = v160
	*(*int32)(unsafe.Add(mBase, uint32(v27)+284)) = v69
	F_errcode(m, int32(1088))
	mBase = m.M
	v690 = m.ExcPending
	if v690 != 0 {
		goto L6
	} else {
		goto L77
	}
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+220)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v27)+216)) = v36
	*(*int64)(unsafe.Add(mBase, uint32(v27)+224)) = v49
	*(*int64)(unsafe.Add(mBase, uint32(v27)+232)) = v50
	*(*int64)(unsafe.Add(mBase, uint32(v27)+240)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v27)+252)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v27)+256)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v27)+260)) = v267
	*(*int32)(unsafe.Add(mBase, uint32(v27)+264)) = v263
	*(*int32)(unsafe.Add(mBase, uint32(v27)+268)) = v166
	*(*int64)(unsafe.Add(mBase, uint32(v27)+272)) = v160
	*(*int32)(unsafe.Add(mBase, uint32(v27)+284)) = v69
	F_errmsg(m, int32(_a_F_pg_logical_slot_get_changes_guts_9), int32(0))
	mBase = m.M
	v706 = m.ExcPending
	if v706 != 0 {
		goto L6
	} else {
		goto L78
	}
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+220)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v27)+216)) = v36
	*(*int64)(unsafe.Add(mBase, uint32(v27)+224)) = v49
	*(*int64)(unsafe.Add(mBase, uint32(v27)+232)) = v50
	*(*int64)(unsafe.Add(mBase, uint32(v27)+240)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v27)+252)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v27)+256)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v27)+260)) = v267
	*(*int32)(unsafe.Add(mBase, uint32(v27)+264)) = v263
	*(*int32)(unsafe.Add(mBase, uint32(v27)+268)) = v166
	*(*int64)(unsafe.Add(mBase, uint32(v27)+272)) = v160
	*(*int32)(unsafe.Add(mBase, uint32(v27)+284)) = v69
	F_errfinish(m, int32(_a_F_pg_logical_slot_get_changes_guts_1), int32(177), int32(_a_F_pg_logical_slot_get_changes_guts_2))
	mBase = m.M
	v723 = m.ExcPending
	if v723 != 0 {
		goto L6
	} else {
		goto L79
	}
L79:
	;
	goto L3
L80:
	;
	v739 = v737
	v740 = v51
	v741 = v737
	goto L50
L81:
	;
	v760 = *(*int32)(unsafe.Add(mBase, _c_F_pg_logical_slot_get_changes_guts[6]))
	v762 = *(*int32)(unsafe.Add(mBase, _c_F_pg_logical_slot_get_changes_guts[7]))
	goto L82
L82:
	;
	v764 = v27 + int32(48)
	*(*int32)(unsafe.Add(mBase, uint32(v764)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v764))) = v27 + int32(28)
	goto L85
L83:
	;
	v774 = v263
	v775 = v166
	v776 = v760
	v777 = v762
	v778 = v267
	v779 = v562
	v780 = v69
	v782 = v565
	v783 = int32(0)
	v788 = v160
	v789 = v741
	v790 = v739
	v791 = v740
	goto L7
L85:
	;
	goto L83
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+44)) = int32(394)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+40)) = int32(395)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+36)) = int32(396)
	*(*int32)(unsafe.Add(mBase, _c_F_pg_logical_slot_get_changes_guts[7])) = v27 + int32(48)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+216)) = v776
	*(*int32)(unsafe.Add(mBase, uint32(v27)+220)) = v777
	*(*int64)(unsafe.Add(mBase, uint32(v27)+224)) = v789
	*(*int64)(unsafe.Add(mBase, uint32(v27)+232)) = v790
	*(*int64)(unsafe.Add(mBase, uint32(v27)+240)) = v791
	*(*int32)(unsafe.Add(mBase, uint32(v27)+252)) = v782
	*(*int32)(unsafe.Add(mBase, uint32(v27)+256)) = v779
	*(*int32)(unsafe.Add(mBase, uint32(v27)+260)) = v778
	*(*int32)(unsafe.Add(mBase, uint32(v27)+264)) = v774
	*(*int32)(unsafe.Add(mBase, uint32(v27)+268)) = v775
	*(*int64)(unsafe.Add(mBase, uint32(v27)+272)) = v788
	*(*int32)(unsafe.Add(mBase, uint32(v27)+284)) = v780
	v819 = int32(0)
	v825 = F_CreateDecodingContext(m, int64(0), v782, v819, v27+int32(36), int32(1012), int32(1013), v819)
	mBase = m.M
	v826 = m.ExcPending
	if v826 != 0 {
		goto L6
	} else {
		goto L89
	}
L87:
	;
	goto L88
L88:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_pg_logical_slot_get_changes_guts[6])) = v776
	*(*int32)(unsafe.Add(mBase, _c_F_pg_logical_slot_get_changes_guts[7])) = v777
	*(*int32)(unsafe.Add(mBase, uint32(v27)+216)) = v776
	*(*int32)(unsafe.Add(mBase, uint32(v27)+220)) = v777
	*(*int64)(unsafe.Add(mBase, uint32(v27)+224)) = v789
	*(*int64)(unsafe.Add(mBase, uint32(v27)+232)) = v790
	*(*int64)(unsafe.Add(mBase, uint32(v27)+240)) = v791
	*(*int32)(unsafe.Add(mBase, uint32(v27)+252)) = v782
	*(*int32)(unsafe.Add(mBase, uint32(v27)+256)) = v779
	*(*int32)(unsafe.Add(mBase, uint32(v27)+260)) = v778
	*(*int32)(unsafe.Add(mBase, uint32(v27)+264)) = v774
	*(*int32)(unsafe.Add(mBase, uint32(v27)+268)) = v775
	*(*int64)(unsafe.Add(mBase, uint32(v27)+272)) = v788
	*(*int32)(unsafe.Add(mBase, uint32(v27)+284)) = v780
	F_InvalidateSystemCachesExtended(m)
	mBase = m.M
	v1242 = m.ExcPending
	if v1242 != 0 {
		goto L6
	} else {
		goto L142
	}
L89:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_pg_logical_slot_get_changes_guts[1])) = v778
	if v3 != 0 {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+220)) = v777
	*(*int32)(unsafe.Add(mBase, uint32(v27)+216)) = v776
	*(*int64)(unsafe.Add(mBase, uint32(v27)+232)) = v790
	*(*int64)(unsafe.Add(mBase, uint32(v27)+240)) = v791
	*(*int32)(unsafe.Add(mBase, uint32(v27)+252)) = v782
	*(*int32)(unsafe.Add(mBase, uint32(v27)+256)) = v779
	*(*int32)(unsafe.Add(mBase, uint32(v27)+260)) = v778
	*(*int32)(unsafe.Add(mBase, uint32(v27)+264)) = v774
	*(*int32)(unsafe.Add(mBase, uint32(v27)+268)) = v775
	*(*int32)(unsafe.Add(mBase, uint32(v27)+284)) = v780
	*(*int64)(unsafe.Add(mBase, uint32(v27)+272)) = v788
	*(*int64)(unsafe.Add(mBase, uint32(v27)+224)) = v789
	if base.Ui64(v788) < base.Ui64(v789) {
		goto L98
	} else {
		goto L99
	}
L91:
	;
	v829 = *(*int32)(unsafe.Add(mBase, uint32(v825)+108))
	if v829 == int32(1) {
		goto L90
	} else {
		goto L92
	}
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+220)) = v777
	*(*int32)(unsafe.Add(mBase, uint32(v27)+216)) = v776
	*(*int64)(unsafe.Add(mBase, uint32(v27)+224)) = v789
	*(*int64)(unsafe.Add(mBase, uint32(v27)+232)) = v790
	*(*int64)(unsafe.Add(mBase, uint32(v27)+240)) = v791
	*(*int32)(unsafe.Add(mBase, uint32(v27)+252)) = v782
	*(*int32)(unsafe.Add(mBase, uint32(v27)+256)) = v779
	*(*int32)(unsafe.Add(mBase, uint32(v27)+260)) = v778
	*(*int32)(unsafe.Add(mBase, uint32(v27)+264)) = v774
	*(*int32)(unsafe.Add(mBase, uint32(v27)+268)) = v775
	*(*int64)(unsafe.Add(mBase, uint32(v27)+272)) = v788
	*(*int32)(unsafe.Add(mBase, uint32(v27)+284)) = v780
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v847 = m.ExcPending
	if v847 != 0 {
		goto L6
	} else {
		goto L93
	}
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+220)) = v777
	*(*int32)(unsafe.Add(mBase, uint32(v27)+216)) = v776
	*(*int64)(unsafe.Add(mBase, uint32(v27)+224)) = v789
	*(*int64)(unsafe.Add(mBase, uint32(v27)+232)) = v790
	*(*int64)(unsafe.Add(mBase, uint32(v27)+240)) = v791
	*(*int32)(unsafe.Add(mBase, uint32(v27)+252)) = v782
	*(*int32)(unsafe.Add(mBase, uint32(v27)+256)) = v779
	*(*int32)(unsafe.Add(mBase, uint32(v27)+260)) = v778
	*(*int32)(unsafe.Add(mBase, uint32(v27)+264)) = v774
	*(*int32)(unsafe.Add(mBase, uint32(v27)+268)) = v775
	*(*int64)(unsafe.Add(mBase, uint32(v27)+272)) = v788
	*(*int32)(unsafe.Add(mBase, uint32(v27)+284)) = v780
	F_errcode(m, int32(1088))
	mBase = m.M
	v862 = m.ExcPending
	if v862 != 0 {
		goto L6
	} else {
		goto L94
	}
L94:
	;
	v863 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v864 = *(*int32)(unsafe.Add(mBase, uint32(v863)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+216)) = v776
	*(*int32)(unsafe.Add(mBase, uint32(v27)+220)) = v777
	*(*int64)(unsafe.Add(mBase, uint32(v27)+224)) = v789
	*(*int64)(unsafe.Add(mBase, uint32(v27)+232)) = v790
	*(*int64)(unsafe.Add(mBase, uint32(v27)+240)) = v791
	*(*int32)(unsafe.Add(mBase, uint32(v27)+252)) = v782
	*(*int32)(unsafe.Add(mBase, uint32(v27)+256)) = v779
	*(*int32)(unsafe.Add(mBase, uint32(v27)+260)) = v778
	*(*int32)(unsafe.Add(mBase, uint32(v27)+264)) = v774
	*(*int32)(unsafe.Add(mBase, uint32(v27)+268)) = v775
	*(*int64)(unsafe.Add(mBase, uint32(v27)+272)) = v788
	*(*int32)(unsafe.Add(mBase, uint32(v27)+284)) = v780
	v878 = *(*int32)(unsafe.Add(mBase, _c_F_pg_logical_slot_get_changes_guts[8]))
	v879 = F_format_procedure(m, v864)
	mBase = m.M
	v880 = m.ExcPending
	if v880 != 0 {
		goto L6
	} else {
		goto L95
	}
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+220)) = v777
	*(*int32)(unsafe.Add(mBase, uint32(v27)+216)) = v776
	*(*int64)(unsafe.Add(mBase, uint32(v27)+224)) = v789
	*(*int64)(unsafe.Add(mBase, uint32(v27)+232)) = v790
	*(*int64)(unsafe.Add(mBase, uint32(v27)+240)) = v791
	*(*int32)(unsafe.Add(mBase, uint32(v27)+252)) = v782
	*(*int32)(unsafe.Add(mBase, uint32(v27)+256)) = v779
	*(*int32)(unsafe.Add(mBase, uint32(v27)+260)) = v778
	*(*int32)(unsafe.Add(mBase, uint32(v27)+264)) = v774
	*(*int32)(unsafe.Add(mBase, uint32(v27)+268)) = v775
	*(*int64)(unsafe.Add(mBase, uint32(v27)+272)) = v788
	*(*int32)(unsafe.Add(mBase, uint32(v27)+284)) = v780
	*(*int32)(unsafe.Add(mBase, uint32(v27)+20)) = v879
	*(*int32)(unsafe.Add(mBase, uint32(v27)+16)) = v878 + int32(137)
	F_errmsg(m, int32(_a_F_pg_logical_slot_get_changes_guts_10), v27+int32(16))
	mBase = m.M
	v901 = m.ExcPending
	if v901 != 0 {
		goto L6
	} else {
		goto L96
	}
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+220)) = v777
	*(*int32)(unsafe.Add(mBase, uint32(v27)+216)) = v776
	*(*int64)(unsafe.Add(mBase, uint32(v27)+224)) = v789
	*(*int64)(unsafe.Add(mBase, uint32(v27)+232)) = v790
	*(*int64)(unsafe.Add(mBase, uint32(v27)+240)) = v791
	*(*int32)(unsafe.Add(mBase, uint32(v27)+252)) = v782
	*(*int32)(unsafe.Add(mBase, uint32(v27)+256)) = v779
	*(*int32)(unsafe.Add(mBase, uint32(v27)+260)) = v778
	*(*int32)(unsafe.Add(mBase, uint32(v27)+264)) = v774
	*(*int32)(unsafe.Add(mBase, uint32(v27)+268)) = v775
	*(*int64)(unsafe.Add(mBase, uint32(v27)+272)) = v788
	*(*int32)(unsafe.Add(mBase, uint32(v27)+284)) = v780
	F_errfinish(m, int32(_a_F_pg_logical_slot_get_changes_guts_1), int32(226), int32(_a_F_pg_logical_slot_get_changes_guts_2))
	mBase = m.M
	v918 = m.ExcPending
	if v918 != 0 {
		goto L6
	} else {
		goto L97
	}
L97:
	;
	goto L3
L98:
	;
	v932 = v788
	goto L100
L99:
	;
	v932 = v789
	goto L100
L100:
	;
	v934 = base.B2i32(v788 == int64(0))
	if v788 == int64(0) {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	v935 = v789
	goto L103
L102:
	;
	v935 = v932
	goto L103
L103:
	;
	F_WaitForStandbyConfirmation(m, v935)
	mBase = m.M
	v937 = m.ExcPending
	if v937 != 0 {
		goto L6
	} else {
		goto L104
	}
L104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v825)+140)) = v774
	v939 = *(*int32)(unsafe.Add(mBase, uint32(v825)+8))
	v941 = *(*int32)(unsafe.Add(mBase, _c_F_pg_logical_slot_get_changes_guts[8]))
	v942 = *(*int64)(unsafe.Add(mBase, uint32(v941)+104))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+216)) = v776
	*(*int32)(unsafe.Add(mBase, uint32(v27)+220)) = v777
	*(*int64)(unsafe.Add(mBase, uint32(v27)+224)) = v789
	*(*int64)(unsafe.Add(mBase, uint32(v27)+232)) = v790
	*(*int64)(unsafe.Add(mBase, uint32(v27)+240)) = v791
	*(*int32)(unsafe.Add(mBase, uint32(v27)+252)) = v782
	*(*int32)(unsafe.Add(mBase, uint32(v27)+256)) = v779
	*(*int32)(unsafe.Add(mBase, uint32(v27)+260)) = v778
	*(*int32)(unsafe.Add(mBase, uint32(v27)+264)) = v774
	*(*int32)(unsafe.Add(mBase, uint32(v27)+268)) = v775
	*(*int64)(unsafe.Add(mBase, uint32(v27)+272)) = v788
	*(*int32)(unsafe.Add(mBase, uint32(v27)+284)) = v780
	F_XLogBeginRead(m, v939, v942)
	mBase = m.M
	v956 = m.ExcPending
	if v956 != 0 {
		goto L6
	} else {
		goto L105
	}
L105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+220)) = v777
	*(*int32)(unsafe.Add(mBase, uint32(v27)+216)) = v776
	*(*int64)(unsafe.Add(mBase, uint32(v27)+224)) = v789
	*(*int64)(unsafe.Add(mBase, uint32(v27)+232)) = v790
	*(*int64)(unsafe.Add(mBase, uint32(v27)+240)) = v791
	*(*int32)(unsafe.Add(mBase, uint32(v27)+252)) = v782
	*(*int32)(unsafe.Add(mBase, uint32(v27)+256)) = v779
	*(*int32)(unsafe.Add(mBase, uint32(v27)+260)) = v778
	*(*int32)(unsafe.Add(mBase, uint32(v27)+264)) = v774
	*(*int32)(unsafe.Add(mBase, uint32(v27)+268)) = v775
	*(*int64)(unsafe.Add(mBase, uint32(v27)+272)) = v788
	*(*int32)(unsafe.Add(mBase, uint32(v27)+284)) = v780
	F_InvalidateSystemCachesExtended(m)
	mBase = m.M
	v970 = m.ExcPending
	if v970 != 0 {
		goto L6
	} else {
		goto L106
	}
L106:
	;
	v971 = *(*int32)(unsafe.Add(mBase, uint32(v825)+8))
	v972 = *(*int64)(unsafe.Add(mBase, uint32(v971)+40))
	if base.Ui64(v789) <= base.Ui64(v972) {
		v1137 = v972
		goto L107
	} else {
		goto L108
	}
L107:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_pg_logical_slot_get_changes_guts[0])) = v780
	v1141 = int32(0)
	if base.B2i32(l1 == v1141)|base.B2i32(v1137 == int64(0)) == v1141 {
		goto L134
	} else {
		goto L135
	}
L108:
	;
	v988 = v971
	goto L109
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+216)) = v776
	*(*int32)(unsafe.Add(mBase, uint32(v27)+32)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+220)) = v777
	*(*int64)(unsafe.Add(mBase, uint32(v27)+224)) = v789
	*(*int64)(unsafe.Add(mBase, uint32(v27)+232)) = v790
	*(*int64)(unsafe.Add(mBase, uint32(v27)+240)) = v791
	*(*int32)(unsafe.Add(mBase, uint32(v27)+252)) = v782
	*(*int32)(unsafe.Add(mBase, uint32(v27)+256)) = v779
	*(*int32)(unsafe.Add(mBase, uint32(v27)+260)) = v778
	*(*int32)(unsafe.Add(mBase, uint32(v27)+264)) = v774
	*(*int32)(unsafe.Add(mBase, uint32(v27)+268)) = v775
	*(*int64)(unsafe.Add(mBase, uint32(v27)+272)) = v788
	*(*int32)(unsafe.Add(mBase, uint32(v27)+284)) = v780
	v1015 = F_XLogReadRecord(m, v988, v27+int32(32))
	mBase = m.M
	v1016 = m.ExcPending
	if v1016 != 0 {
		goto L6
	} else {
		goto L111
	}
L110:
	;
	v1137 = v1113
	goto L107
L111:
	;
	v1017 = *(*int32)(unsafe.Add(mBase, uint32(v27)+32))
	if v1017 != 0 {
		goto L112
	} else {
		goto L113
	}
L112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+220)) = v777
	*(*int32)(unsafe.Add(mBase, uint32(v27)+216)) = v776
	*(*int64)(unsafe.Add(mBase, uint32(v27)+224)) = v789
	*(*int64)(unsafe.Add(mBase, uint32(v27)+232)) = v790
	*(*int64)(unsafe.Add(mBase, uint32(v27)+240)) = v791
	*(*int32)(unsafe.Add(mBase, uint32(v27)+252)) = v782
	*(*int32)(unsafe.Add(mBase, uint32(v27)+256)) = v779
	*(*int32)(unsafe.Add(mBase, uint32(v27)+260)) = v778
	*(*int32)(unsafe.Add(mBase, uint32(v27)+264)) = v774
	*(*int32)(unsafe.Add(mBase, uint32(v27)+268)) = v775
	*(*int64)(unsafe.Add(mBase, uint32(v27)+272)) = v788
	*(*int32)(unsafe.Add(mBase, uint32(v27)+284)) = v780
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1033 = m.ExcPending
	if v1033 != 0 {
		goto L6
	} else {
		goto L115
	}
L113:
	;
	goto L114
L114:
	;
	if v1015 != 0 {
		goto L118
	} else {
		goto L119
	}
L115:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+216)) = v776
	*(*int32)(unsafe.Add(mBase, uint32(v27)+220)) = v777
	*(*int64)(unsafe.Add(mBase, uint32(v27)+224)) = v789
	*(*int64)(unsafe.Add(mBase, uint32(v27)+232)) = v790
	*(*int64)(unsafe.Add(mBase, uint32(v27)+240)) = v791
	*(*int32)(unsafe.Add(mBase, uint32(v27)+252)) = v782
	*(*int32)(unsafe.Add(mBase, uint32(v27)+256)) = v779
	*(*int32)(unsafe.Add(mBase, uint32(v27)+260)) = v778
	*(*int32)(unsafe.Add(mBase, uint32(v27)+264)) = v774
	*(*int32)(unsafe.Add(mBase, uint32(v27)+268)) = v775
	*(*int64)(unsafe.Add(mBase, uint32(v27)+272)) = v788
	*(*int32)(unsafe.Add(mBase, uint32(v27)+284)) = v780
	v1046 = *(*int32)(unsafe.Add(mBase, uint32(v27)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v27))) = v1046
	F_errmsg_internal(m, int32(_a_F_pg_logical_slot_get_changes_guts_11), v27)
	mBase = m.M
	v1050 = m.ExcPending
	if v1050 != 0 {
		goto L6
	} else {
		goto L116
	}
L116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+220)) = v777
	*(*int32)(unsafe.Add(mBase, uint32(v27)+216)) = v776
	*(*int64)(unsafe.Add(mBase, uint32(v27)+224)) = v789
	*(*int64)(unsafe.Add(mBase, uint32(v27)+232)) = v790
	*(*int64)(unsafe.Add(mBase, uint32(v27)+240)) = v791
	*(*int32)(unsafe.Add(mBase, uint32(v27)+252)) = v782
	*(*int32)(unsafe.Add(mBase, uint32(v27)+256)) = v779
	*(*int32)(unsafe.Add(mBase, uint32(v27)+260)) = v778
	*(*int32)(unsafe.Add(mBase, uint32(v27)+264)) = v774
	*(*int32)(unsafe.Add(mBase, uint32(v27)+268)) = v775
	*(*int64)(unsafe.Add(mBase, uint32(v27)+272)) = v788
	*(*int32)(unsafe.Add(mBase, uint32(v27)+284)) = v780
	F_errfinish(m, int32(_a_F_pg_logical_slot_get_changes_guts_1), int32(259), int32(_a_F_pg_logical_slot_get_changes_guts_2))
	mBase = m.M
	v1067 = m.ExcPending
	if v1067 != 0 {
		goto L6
	} else {
		goto L117
	}
L117:
	;
	goto L3
L118:
	;
	v1068 = *(*int32)(unsafe.Add(mBase, uint32(v825)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+220)) = v777
	*(*int32)(unsafe.Add(mBase, uint32(v27)+216)) = v776
	*(*int64)(unsafe.Add(mBase, uint32(v27)+224)) = v789
	*(*int64)(unsafe.Add(mBase, uint32(v27)+232)) = v790
	*(*int64)(unsafe.Add(mBase, uint32(v27)+240)) = v791
	*(*int32)(unsafe.Add(mBase, uint32(v27)+252)) = v782
	*(*int32)(unsafe.Add(mBase, uint32(v27)+256)) = v779
	*(*int32)(unsafe.Add(mBase, uint32(v27)+260)) = v778
	*(*int32)(unsafe.Add(mBase, uint32(v27)+264)) = v774
	*(*int32)(unsafe.Add(mBase, uint32(v27)+268)) = v775
	*(*int64)(unsafe.Add(mBase, uint32(v27)+272)) = v788
	*(*int32)(unsafe.Add(mBase, uint32(v27)+284)) = v780
	F_LogicalDecodingProcessRecord(m, v825, v1068)
	mBase = m.M
	v1082 = m.ExcPending
	if v1082 != 0 {
		goto L6
	} else {
		goto L121
	}
L119:
	;
	goto L120
L120:
	;
	if v934 == int32(0) {
		goto L122
	} else {
		goto L123
	}
L121:
	;
	goto L120
L122:
	;
	v1086 = *(*int32)(unsafe.Add(mBase, uint32(v825)+8))
	v1087 = *(*int64)(unsafe.Add(mBase, uint32(v1086)+40))
	if base.Ui64(v788) <= base.Ui64(v1087) {
		v1137 = v1087
		goto L107
	} else {
		goto L125
	}
L123:
	;
	goto L124
L124:
	;
	if v775 == int32(0) {
		goto L126
	} else {
		goto L127
	}
L125:
	;
	goto L124
L126:
	;
	v1097 = *(*int32)(unsafe.Add(mBase, _c_F_pg_logical_slot_get_changes_guts[9]))
	if v1097 != 0 {
		goto L129
	} else {
		goto L130
	}
L127:
	;
	v1092 = *(*int64)(unsafe.Add(mBase, uint32(v774)+16))
	if v1092 < base.I64_extend_i32_s(v775) {
		goto L126
	} else {
		goto L128
	}
L128:
	;
	v1094 = *(*int32)(unsafe.Add(mBase, uint32(v825)+8))
	v1095 = *(*int64)(unsafe.Add(mBase, uint32(v1094)+40))
	v1137 = v1095
	goto L107
L129:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+220)) = v777
	*(*int32)(unsafe.Add(mBase, uint32(v27)+216)) = v776
	*(*int64)(unsafe.Add(mBase, uint32(v27)+224)) = v789
	*(*int64)(unsafe.Add(mBase, uint32(v27)+232)) = v790
	*(*int64)(unsafe.Add(mBase, uint32(v27)+240)) = v791
	*(*int32)(unsafe.Add(mBase, uint32(v27)+252)) = v782
	*(*int32)(unsafe.Add(mBase, uint32(v27)+256)) = v779
	*(*int32)(unsafe.Add(mBase, uint32(v27)+260)) = v778
	*(*int32)(unsafe.Add(mBase, uint32(v27)+264)) = v774
	*(*int32)(unsafe.Add(mBase, uint32(v27)+268)) = v775
	*(*int64)(unsafe.Add(mBase, uint32(v27)+272)) = v788
	*(*int32)(unsafe.Add(mBase, uint32(v27)+284)) = v780
	F_ProcessInterrupts(m)
	mBase = m.M
	v1111 = m.ExcPending
	if v1111 != 0 {
		goto L6
	} else {
		goto L132
	}
L130:
	;
	goto L131
L131:
	;
	v1112 = *(*int32)(unsafe.Add(mBase, uint32(v825)+8))
	v1113 = *(*int64)(unsafe.Add(mBase, uint32(v1112)+40))
	if base.Ui64(v1113) < base.Ui64(v789) {
		v988 = v1112
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
	*(*int32)(unsafe.Add(mBase, uint32(v27)+220)) = v777
	*(*int32)(unsafe.Add(mBase, uint32(v27)+216)) = v776
	*(*int64)(unsafe.Add(mBase, uint32(v27)+224)) = v789
	*(*int64)(unsafe.Add(mBase, uint32(v27)+232)) = v790
	*(*int64)(unsafe.Add(mBase, uint32(v27)+240)) = v791
	*(*int32)(unsafe.Add(mBase, uint32(v27)+252)) = v782
	*(*int32)(unsafe.Add(mBase, uint32(v27)+256)) = v779
	*(*int32)(unsafe.Add(mBase, uint32(v27)+260)) = v778
	*(*int32)(unsafe.Add(mBase, uint32(v27)+264)) = v774
	*(*int32)(unsafe.Add(mBase, uint32(v27)+268)) = v775
	*(*int64)(unsafe.Add(mBase, uint32(v27)+272)) = v788
	*(*int32)(unsafe.Add(mBase, uint32(v27)+284)) = v780
	F_LogicalConfirmReceivedLocation(m, v1137)
	mBase = m.M
	v1161 = m.ExcPending
	if v1161 != 0 {
		goto L6
	} else {
		goto L137
	}
L135:
	;
	goto L136
L136:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+220)) = v777
	*(*int32)(unsafe.Add(mBase, uint32(v27)+216)) = v776
	*(*int64)(unsafe.Add(mBase, uint32(v27)+224)) = v789
	*(*int64)(unsafe.Add(mBase, uint32(v27)+232)) = v790
	*(*int64)(unsafe.Add(mBase, uint32(v27)+240)) = v791
	*(*int32)(unsafe.Add(mBase, uint32(v27)+252)) = v782
	*(*int32)(unsafe.Add(mBase, uint32(v27)+256)) = v779
	*(*int32)(unsafe.Add(mBase, uint32(v27)+260)) = v778
	*(*int32)(unsafe.Add(mBase, uint32(v27)+264)) = v774
	*(*int32)(unsafe.Add(mBase, uint32(v27)+268)) = v775
	*(*int64)(unsafe.Add(mBase, uint32(v27)+272)) = v788
	*(*int32)(unsafe.Add(mBase, uint32(v27)+284)) = v780
	F_FreeDecodingContext(m, v825)
	mBase = m.M
	v1189 = m.ExcPending
	if v1189 != 0 {
		goto L6
	} else {
		goto L139
	}
L137:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+220)) = v777
	*(*int32)(unsafe.Add(mBase, uint32(v27)+216)) = v776
	*(*int64)(unsafe.Add(mBase, uint32(v27)+224)) = v789
	*(*int64)(unsafe.Add(mBase, uint32(v27)+232)) = v790
	*(*int64)(unsafe.Add(mBase, uint32(v27)+240)) = v791
	*(*int32)(unsafe.Add(mBase, uint32(v27)+252)) = v782
	*(*int32)(unsafe.Add(mBase, uint32(v27)+256)) = v779
	*(*int32)(unsafe.Add(mBase, uint32(v27)+260)) = v778
	*(*int32)(unsafe.Add(mBase, uint32(v27)+264)) = v774
	*(*int32)(unsafe.Add(mBase, uint32(v27)+268)) = v775
	*(*int64)(unsafe.Add(mBase, uint32(v27)+272)) = v788
	*(*int32)(unsafe.Add(mBase, uint32(v27)+284)) = v780
	F_ReplicationSlotMarkDirty(m)
	mBase = m.M
	v1175 = m.ExcPending
	if v1175 != 0 {
		goto L6
	} else {
		goto L138
	}
L138:
	;
	goto L136
L139:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+220)) = v777
	*(*int32)(unsafe.Add(mBase, uint32(v27)+216)) = v776
	*(*int64)(unsafe.Add(mBase, uint32(v27)+224)) = v789
	*(*int64)(unsafe.Add(mBase, uint32(v27)+232)) = v790
	*(*int64)(unsafe.Add(mBase, uint32(v27)+240)) = v791
	*(*int32)(unsafe.Add(mBase, uint32(v27)+252)) = v782
	*(*int32)(unsafe.Add(mBase, uint32(v27)+256)) = v779
	*(*int32)(unsafe.Add(mBase, uint32(v27)+260)) = v778
	*(*int32)(unsafe.Add(mBase, uint32(v27)+264)) = v774
	*(*int32)(unsafe.Add(mBase, uint32(v27)+268)) = v775
	*(*int64)(unsafe.Add(mBase, uint32(v27)+272)) = v788
	*(*int32)(unsafe.Add(mBase, uint32(v27)+284)) = v780
	F_ReplicationSlotRelease(m)
	mBase = m.M
	v1203 = m.ExcPending
	if v1203 != 0 {
		goto L6
	} else {
		goto L140
	}
L140:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+220)) = v777
	*(*int32)(unsafe.Add(mBase, uint32(v27)+216)) = v776
	*(*int64)(unsafe.Add(mBase, uint32(v27)+224)) = v789
	*(*int64)(unsafe.Add(mBase, uint32(v27)+232)) = v790
	*(*int64)(unsafe.Add(mBase, uint32(v27)+240)) = v791
	*(*int32)(unsafe.Add(mBase, uint32(v27)+252)) = v782
	*(*int32)(unsafe.Add(mBase, uint32(v27)+256)) = v779
	*(*int32)(unsafe.Add(mBase, uint32(v27)+260)) = v778
	*(*int32)(unsafe.Add(mBase, uint32(v27)+264)) = v774
	*(*int32)(unsafe.Add(mBase, uint32(v27)+268)) = v775
	*(*int64)(unsafe.Add(mBase, uint32(v27)+272)) = v788
	*(*int32)(unsafe.Add(mBase, uint32(v27)+284)) = v780
	F_InvalidateSystemCachesExtended(m)
	mBase = m.M
	v1217 = m.ExcPending
	if v1217 != 0 {
		goto L6
	} else {
		goto L141
	}
L141:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_pg_logical_slot_get_changes_guts[6])) = v776
	*(*int32)(unsafe.Add(mBase, _c_F_pg_logical_slot_get_changes_guts[7])) = v777
	m.G0 = v27 + int32(288)
	return
L142:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+220)) = v777
	*(*int32)(unsafe.Add(mBase, uint32(v27)+216)) = v776
	*(*int64)(unsafe.Add(mBase, uint32(v27)+224)) = v789
	*(*int64)(unsafe.Add(mBase, uint32(v27)+232)) = v790
	*(*int64)(unsafe.Add(mBase, uint32(v27)+240)) = v791
	*(*int32)(unsafe.Add(mBase, uint32(v27)+252)) = v782
	*(*int32)(unsafe.Add(mBase, uint32(v27)+256)) = v779
	*(*int32)(unsafe.Add(mBase, uint32(v27)+260)) = v778
	*(*int32)(unsafe.Add(mBase, uint32(v27)+264)) = v774
	*(*int32)(unsafe.Add(mBase, uint32(v27)+268)) = v775
	*(*int64)(unsafe.Add(mBase, uint32(v27)+272)) = v788
	*(*int32)(unsafe.Add(mBase, uint32(v27)+284)) = v780
	F_pg_re_throw(m)
	mBase = m.M
	v1256 = m.ExcPending
	if v1256 != 0 {
		goto L6
	} else {
		goto L143
	}
L143:
	;
	goto L5
L144:
	;
	v1286 = int32(v1282)
	m.G0 = v27
	v1288 = *(*int32)(unsafe.Add(mBase, uint32(v1286)+4))
	v1289 = *(*int32)(unsafe.Add(mBase, uint32(v1286)))
	v1292 = *(*int32)(unsafe.Add(mBase, uint32(v1289)))
	if v27+int32(28) == v1292 {
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
	if v1296 != 0 {
		goto L150
	} else {
		goto L151
	}
L147:
	;
	v1294 = *(*int32)(unsafe.Add(mBase, uint32(v1289)+4))
	v1296 = v1294
	goto L149
L148:
	;
	v1296 = int32(0)
	goto L149
L149:
	;
	goto L146
L150:
	;
	v1297 = *(*int32)(unsafe.Add(mBase, uint32(v27)+284))
	v1298 = *(*int64)(unsafe.Add(mBase, uint32(v27)+272))
	v1299 = *(*int32)(unsafe.Add(mBase, uint32(v27)+268))
	v1300 = *(*int32)(unsafe.Add(mBase, uint32(v27)+264))
	v1301 = *(*int32)(unsafe.Add(mBase, uint32(v27)+260))
	v1302 = *(*int32)(unsafe.Add(mBase, uint32(v27)+256))
	v1303 = *(*int32)(unsafe.Add(mBase, uint32(v27)+252))
	v1304 = *(*int64)(unsafe.Add(mBase, uint32(v27)+240))
	v1305 = *(*int64)(unsafe.Add(mBase, uint32(v27)+232))
	v1306 = *(*int64)(unsafe.Add(mBase, uint32(v27)+224))
	v1307 = *(*int32)(unsafe.Add(mBase, uint32(v27)+220))
	v1308 = *(*int32)(unsafe.Add(mBase, uint32(v27)+216))
	v34 = v1300
	v35 = v1299
	v36 = v1308
	v37 = v1307
	v38 = v1301
	v39 = v1302
	v40 = v1297
	v41 = v1303
	v42 = v1296
	v43 = v1288
	v48 = v1298
	v49 = v1306
	v50 = v1305
	v51 = v1304
	goto L1
L151:
	;
	goto L152
L152:
	;
	F___wasm_longjmp(m, v1289, v1288)
	mBase = m.M
	v1310 = m.ExcPending
	if v1310 != 0 {
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
