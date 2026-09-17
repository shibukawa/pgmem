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
	var v636 int32
	_ = v636
	var v637 int32
	_ = v637
	var v638 int64
	_ = v638
	var v643 int32
	_ = v643
	var v644 int64
	_ = v644
	var v653 int64
	_ = v653
	var v669 int32
	_ = v669
	var v684 int32
	_ = v684
	var v700 int32
	_ = v700
	var v717 int32
	_ = v717
	var v731 int64
	_ = v731
	var v732 int32
	_ = v732
	var v733 int64
	_ = v733
	var v734 int64
	_ = v734
	var v735 int64
	_ = v735
	var v748 int32
	_ = v748
	var v751 int32
	_ = v751
	var v754 int32
	_ = v754
	var v756 int32
	_ = v756
	var v758 int32
	_ = v758
	var v768 int32
	_ = v768
	var v769 int32
	_ = v769
	var v770 int32
	_ = v770
	var v771 int32
	_ = v771
	var v772 int32
	_ = v772
	var v773 int32
	_ = v773
	var v774 int32
	_ = v774
	var v776 int32
	_ = v776
	var v777 int32
	_ = v777
	var v782 int64
	_ = v782
	var v783 int64
	_ = v783
	var v784 int64
	_ = v784
	var v785 int64
	_ = v785
	var v813 int32
	_ = v813
	var v819 int32
	_ = v819
	var v820 int32
	_ = v820
	var v823 int32
	_ = v823
	var v841 int32
	_ = v841
	var v856 int32
	_ = v856
	var v857 int32
	_ = v857
	var v858 int32
	_ = v858
	var v872 int32
	_ = v872
	var v873 int32
	_ = v873
	var v874 int32
	_ = v874
	var v895 int32
	_ = v895
	var v912 int32
	_ = v912
	var v926 int64
	_ = v926
	var v928 int32
	_ = v928
	var v929 int64
	_ = v929
	var v931 int32
	_ = v931
	var v933 int32
	_ = v933
	var v935 int32
	_ = v935
	var v936 int64
	_ = v936
	var v950 int32
	_ = v950
	var v964 int32
	_ = v964
	var v965 int32
	_ = v965
	var v966 int64
	_ = v966
	var v982 int32
	_ = v982
	var v1009 int32
	_ = v1009
	var v1010 int32
	_ = v1010
	var v1011 int32
	_ = v1011
	var v1027 int32
	_ = v1027
	var v1040 int32
	_ = v1040
	var v1044 int32
	_ = v1044
	var v1061 int32
	_ = v1061
	var v1062 int32
	_ = v1062
	var v1076 int32
	_ = v1076
	var v1080 int32
	_ = v1080
	var v1081 int64
	_ = v1081
	var v1086 int64
	_ = v1086
	var v1088 int32
	_ = v1088
	var v1089 int64
	_ = v1089
	var v1091 int32
	_ = v1091
	var v1105 int32
	_ = v1105
	var v1106 int32
	_ = v1106
	var v1107 int64
	_ = v1107
	var v1131 int64
	_ = v1131
	var v1135 int32
	_ = v1135
	var v1155 int32
	_ = v1155
	var v1169 int32
	_ = v1169
	var v1183 int32
	_ = v1183
	var v1197 int32
	_ = v1197
	var v1211 int32
	_ = v1211
	var v1236 int32
	_ = v1236
	var v1250 int32
	_ = v1250
	var v1275 int32
	_ = v1275
	var v1276 int64
	_ = v1276
	var v1280 int32
	_ = v1280
	var v1282 int32
	_ = v1282
	var v1283 int32
	_ = v1283
	var v1286 int32
	_ = v1286
	var v1288 int32
	_ = v1288
	var v1290 int32
	_ = v1290
	var v1291 int32
	_ = v1291
	var v1292 int64
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
	var v1298 int64
	_ = v1298
	var v1299 int64
	_ = v1299
	var v1300 int64
	_ = v1300
	var v1301 int32
	_ = v1301
	var v1302 int32
	_ = v1302
	var v1304 int32
	_ = v1304
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
	v1275 = int32(m.ExcTag)
	v1276 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v1275 == int32(0) {
		goto L144
	} else {
		goto L145
	}
L7:
	;
	if v777 == int32(0) {
		goto L86
	} else {
		goto L87
	}
L8:
	;
	v768 = v34
	v769 = v35
	v770 = v36
	v771 = v37
	v772 = v38
	v773 = v39
	v774 = v40
	v776 = v41
	v777 = v43
	v782 = v48
	v783 = v49
	v784 = v50
	v785 = v51
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
	*(*int64)(unsafe.Add(mBase, uint32(v27)+224)) = v735
	*(*int64)(unsafe.Add(mBase, uint32(v27)+232)) = v733
	*(*int64)(unsafe.Add(mBase, uint32(v27)+240)) = v734
	*(*int32)(unsafe.Add(mBase, uint32(v27)+252)) = v565
	*(*int32)(unsafe.Add(mBase, uint32(v27)+256)) = v562
	*(*int32)(unsafe.Add(mBase, uint32(v27)+260)) = v267
	*(*int32)(unsafe.Add(mBase, uint32(v27)+264)) = v263
	*(*int32)(unsafe.Add(mBase, uint32(v27)+268)) = v166
	*(*int64)(unsafe.Add(mBase, uint32(v27)+272)) = v160
	*(*int32)(unsafe.Add(mBase, uint32(v27)+284)) = v69
	v748 = int32(1)
	F_ReplicationSlotAcquire(m, v234, v748, v748)
	mBase = m.M
	v751 = m.ExcPending
	if v751 != 0 {
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
	v731 = F_GetXLogReplayRecPtr(m, int32(0))
	mBase = m.M
	v732 = m.ExcPending
	if v732 != 0 {
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
	v669 = m.ExcPending
	if v669 != 0 {
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
	v636 = int32(_a_F_pg_logical_slot_get_changes_guts_7)
	v637 = *(*int32)(unsafe.Add(mBase, _c_F_pg_logical_slot_get_changes_guts[3]))
	v638 = *(*int64)(unsafe.Add(mBase, uint32(v637)+280))
	*(*int64)(unsafe.Add(mBase, uint32(v637)+280)) = v638
	*(*int64)(unsafe.Add(mBase, _c_F_pg_logical_slot_get_changes_guts[4])) = v638
	v643 = *(*int32)(unsafe.Add(mBase, _c_F_pg_logical_slot_get_changes_guts[3]))
	v644 = *(*int64)(unsafe.Add(mBase, uint32(v643)+272))
	*(*int64)(unsafe.Add(mBase, uint32(v643)+272)) = v644
	*(*int64)(unsafe.Add(mBase, _c_F_pg_logical_slot_get_changes_guts[5])) = v644
	goto L74
L72:
	;
	v733 = v50
	v734 = v653
	v735 = v653
	goto L50
L74:
	;
	goto L75
L75:
	;
	v653 = *(*int64)(unsafe.Add(mBase, _c_F_pg_logical_slot_get_changes_guts[4]))
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
	v684 = m.ExcPending
	if v684 != 0 {
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
	F_errmsg(m, int32(_a_F_pg_logical_slot_get_changes_guts_8), int32(0))
	mBase = m.M
	v700 = m.ExcPending
	if v700 != 0 {
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
	v717 = m.ExcPending
	if v717 != 0 {
		goto L6
	} else {
		goto L79
	}
L79:
	;
	goto L3
L80:
	;
	v733 = v731
	v734 = v51
	v735 = v731
	goto L50
L81:
	;
	v754 = *(*int32)(unsafe.Add(mBase, _c_F_pg_logical_slot_get_changes_guts[6]))
	v756 = *(*int32)(unsafe.Add(mBase, _c_F_pg_logical_slot_get_changes_guts[7]))
	goto L82
L82:
	;
	v758 = v27 + int32(48)
	*(*int32)(unsafe.Add(mBase, uint32(v758)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v758))) = v27 + int32(28)
	goto L85
L83:
	;
	v768 = v263
	v769 = v166
	v770 = v754
	v771 = v756
	v772 = v267
	v773 = v562
	v774 = v69
	v776 = v565
	v777 = int32(0)
	v782 = v160
	v783 = v735
	v784 = v733
	v785 = v734
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
	*(*int32)(unsafe.Add(mBase, uint32(v27)+216)) = v770
	*(*int32)(unsafe.Add(mBase, uint32(v27)+220)) = v771
	*(*int64)(unsafe.Add(mBase, uint32(v27)+224)) = v783
	*(*int64)(unsafe.Add(mBase, uint32(v27)+232)) = v784
	*(*int64)(unsafe.Add(mBase, uint32(v27)+240)) = v785
	*(*int32)(unsafe.Add(mBase, uint32(v27)+252)) = v776
	*(*int32)(unsafe.Add(mBase, uint32(v27)+256)) = v773
	*(*int32)(unsafe.Add(mBase, uint32(v27)+260)) = v772
	*(*int32)(unsafe.Add(mBase, uint32(v27)+264)) = v768
	*(*int32)(unsafe.Add(mBase, uint32(v27)+268)) = v769
	*(*int64)(unsafe.Add(mBase, uint32(v27)+272)) = v782
	*(*int32)(unsafe.Add(mBase, uint32(v27)+284)) = v774
	v813 = int32(0)
	v819 = F_CreateDecodingContext(m, int64(0), v776, v813, v27+int32(36), int32(1012), int32(1013), v813)
	mBase = m.M
	v820 = m.ExcPending
	if v820 != 0 {
		goto L6
	} else {
		goto L89
	}
L87:
	;
	goto L88
L88:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_pg_logical_slot_get_changes_guts[6])) = v770
	*(*int32)(unsafe.Add(mBase, _c_F_pg_logical_slot_get_changes_guts[7])) = v771
	*(*int32)(unsafe.Add(mBase, uint32(v27)+216)) = v770
	*(*int32)(unsafe.Add(mBase, uint32(v27)+220)) = v771
	*(*int64)(unsafe.Add(mBase, uint32(v27)+224)) = v783
	*(*int64)(unsafe.Add(mBase, uint32(v27)+232)) = v784
	*(*int64)(unsafe.Add(mBase, uint32(v27)+240)) = v785
	*(*int32)(unsafe.Add(mBase, uint32(v27)+252)) = v776
	*(*int32)(unsafe.Add(mBase, uint32(v27)+256)) = v773
	*(*int32)(unsafe.Add(mBase, uint32(v27)+260)) = v772
	*(*int32)(unsafe.Add(mBase, uint32(v27)+264)) = v768
	*(*int32)(unsafe.Add(mBase, uint32(v27)+268)) = v769
	*(*int64)(unsafe.Add(mBase, uint32(v27)+272)) = v782
	*(*int32)(unsafe.Add(mBase, uint32(v27)+284)) = v774
	F_InvalidateSystemCachesExtended(m)
	mBase = m.M
	v1236 = m.ExcPending
	if v1236 != 0 {
		goto L6
	} else {
		goto L142
	}
L89:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_pg_logical_slot_get_changes_guts[1])) = v772
	if v3 != 0 {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+220)) = v771
	*(*int32)(unsafe.Add(mBase, uint32(v27)+216)) = v770
	*(*int64)(unsafe.Add(mBase, uint32(v27)+232)) = v784
	*(*int64)(unsafe.Add(mBase, uint32(v27)+240)) = v785
	*(*int32)(unsafe.Add(mBase, uint32(v27)+252)) = v776
	*(*int32)(unsafe.Add(mBase, uint32(v27)+256)) = v773
	*(*int32)(unsafe.Add(mBase, uint32(v27)+260)) = v772
	*(*int32)(unsafe.Add(mBase, uint32(v27)+264)) = v768
	*(*int32)(unsafe.Add(mBase, uint32(v27)+268)) = v769
	*(*int32)(unsafe.Add(mBase, uint32(v27)+284)) = v774
	*(*int64)(unsafe.Add(mBase, uint32(v27)+272)) = v782
	*(*int64)(unsafe.Add(mBase, uint32(v27)+224)) = v783
	if base.Ui64(v782) < base.Ui64(v783) {
		goto L98
	} else {
		goto L99
	}
L91:
	;
	v823 = *(*int32)(unsafe.Add(mBase, uint32(v819)+108))
	if v823 == int32(1) {
		goto L90
	} else {
		goto L92
	}
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+220)) = v771
	*(*int32)(unsafe.Add(mBase, uint32(v27)+216)) = v770
	*(*int64)(unsafe.Add(mBase, uint32(v27)+224)) = v783
	*(*int64)(unsafe.Add(mBase, uint32(v27)+232)) = v784
	*(*int64)(unsafe.Add(mBase, uint32(v27)+240)) = v785
	*(*int32)(unsafe.Add(mBase, uint32(v27)+252)) = v776
	*(*int32)(unsafe.Add(mBase, uint32(v27)+256)) = v773
	*(*int32)(unsafe.Add(mBase, uint32(v27)+260)) = v772
	*(*int32)(unsafe.Add(mBase, uint32(v27)+264)) = v768
	*(*int32)(unsafe.Add(mBase, uint32(v27)+268)) = v769
	*(*int64)(unsafe.Add(mBase, uint32(v27)+272)) = v782
	*(*int32)(unsafe.Add(mBase, uint32(v27)+284)) = v774
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v841 = m.ExcPending
	if v841 != 0 {
		goto L6
	} else {
		goto L93
	}
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+220)) = v771
	*(*int32)(unsafe.Add(mBase, uint32(v27)+216)) = v770
	*(*int64)(unsafe.Add(mBase, uint32(v27)+224)) = v783
	*(*int64)(unsafe.Add(mBase, uint32(v27)+232)) = v784
	*(*int64)(unsafe.Add(mBase, uint32(v27)+240)) = v785
	*(*int32)(unsafe.Add(mBase, uint32(v27)+252)) = v776
	*(*int32)(unsafe.Add(mBase, uint32(v27)+256)) = v773
	*(*int32)(unsafe.Add(mBase, uint32(v27)+260)) = v772
	*(*int32)(unsafe.Add(mBase, uint32(v27)+264)) = v768
	*(*int32)(unsafe.Add(mBase, uint32(v27)+268)) = v769
	*(*int64)(unsafe.Add(mBase, uint32(v27)+272)) = v782
	*(*int32)(unsafe.Add(mBase, uint32(v27)+284)) = v774
	F_errcode(m, int32(1088))
	mBase = m.M
	v856 = m.ExcPending
	if v856 != 0 {
		goto L6
	} else {
		goto L94
	}
L94:
	;
	v857 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v858 = *(*int32)(unsafe.Add(mBase, uint32(v857)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+216)) = v770
	*(*int32)(unsafe.Add(mBase, uint32(v27)+220)) = v771
	*(*int64)(unsafe.Add(mBase, uint32(v27)+224)) = v783
	*(*int64)(unsafe.Add(mBase, uint32(v27)+232)) = v784
	*(*int64)(unsafe.Add(mBase, uint32(v27)+240)) = v785
	*(*int32)(unsafe.Add(mBase, uint32(v27)+252)) = v776
	*(*int32)(unsafe.Add(mBase, uint32(v27)+256)) = v773
	*(*int32)(unsafe.Add(mBase, uint32(v27)+260)) = v772
	*(*int32)(unsafe.Add(mBase, uint32(v27)+264)) = v768
	*(*int32)(unsafe.Add(mBase, uint32(v27)+268)) = v769
	*(*int64)(unsafe.Add(mBase, uint32(v27)+272)) = v782
	*(*int32)(unsafe.Add(mBase, uint32(v27)+284)) = v774
	v872 = *(*int32)(unsafe.Add(mBase, _c_F_pg_logical_slot_get_changes_guts[8]))
	v873 = F_format_procedure(m, v858)
	mBase = m.M
	v874 = m.ExcPending
	if v874 != 0 {
		goto L6
	} else {
		goto L95
	}
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+220)) = v771
	*(*int32)(unsafe.Add(mBase, uint32(v27)+216)) = v770
	*(*int64)(unsafe.Add(mBase, uint32(v27)+224)) = v783
	*(*int64)(unsafe.Add(mBase, uint32(v27)+232)) = v784
	*(*int64)(unsafe.Add(mBase, uint32(v27)+240)) = v785
	*(*int32)(unsafe.Add(mBase, uint32(v27)+252)) = v776
	*(*int32)(unsafe.Add(mBase, uint32(v27)+256)) = v773
	*(*int32)(unsafe.Add(mBase, uint32(v27)+260)) = v772
	*(*int32)(unsafe.Add(mBase, uint32(v27)+264)) = v768
	*(*int32)(unsafe.Add(mBase, uint32(v27)+268)) = v769
	*(*int64)(unsafe.Add(mBase, uint32(v27)+272)) = v782
	*(*int32)(unsafe.Add(mBase, uint32(v27)+284)) = v774
	*(*int32)(unsafe.Add(mBase, uint32(v27)+20)) = v873
	*(*int32)(unsafe.Add(mBase, uint32(v27)+16)) = v872 + int32(137)
	F_errmsg(m, int32(_a_F_pg_logical_slot_get_changes_guts_9), v27+int32(16))
	mBase = m.M
	v895 = m.ExcPending
	if v895 != 0 {
		goto L6
	} else {
		goto L96
	}
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+220)) = v771
	*(*int32)(unsafe.Add(mBase, uint32(v27)+216)) = v770
	*(*int64)(unsafe.Add(mBase, uint32(v27)+224)) = v783
	*(*int64)(unsafe.Add(mBase, uint32(v27)+232)) = v784
	*(*int64)(unsafe.Add(mBase, uint32(v27)+240)) = v785
	*(*int32)(unsafe.Add(mBase, uint32(v27)+252)) = v776
	*(*int32)(unsafe.Add(mBase, uint32(v27)+256)) = v773
	*(*int32)(unsafe.Add(mBase, uint32(v27)+260)) = v772
	*(*int32)(unsafe.Add(mBase, uint32(v27)+264)) = v768
	*(*int32)(unsafe.Add(mBase, uint32(v27)+268)) = v769
	*(*int64)(unsafe.Add(mBase, uint32(v27)+272)) = v782
	*(*int32)(unsafe.Add(mBase, uint32(v27)+284)) = v774
	F_errfinish(m, int32(_a_F_pg_logical_slot_get_changes_guts_1), int32(226), int32(_a_F_pg_logical_slot_get_changes_guts_2))
	mBase = m.M
	v912 = m.ExcPending
	if v912 != 0 {
		goto L6
	} else {
		goto L97
	}
L97:
	;
	goto L3
L98:
	;
	v926 = v782
	goto L100
L99:
	;
	v926 = v783
	goto L100
L100:
	;
	v928 = base.B2i32(v782 == int64(0))
	if v782 == int64(0) {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	v929 = v783
	goto L103
L102:
	;
	v929 = v926
	goto L103
L103:
	;
	F_WaitForStandbyConfirmation(m, v929)
	mBase = m.M
	v931 = m.ExcPending
	if v931 != 0 {
		goto L6
	} else {
		goto L104
	}
L104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v819)+140)) = v768
	v933 = *(*int32)(unsafe.Add(mBase, uint32(v819)+8))
	v935 = *(*int32)(unsafe.Add(mBase, _c_F_pg_logical_slot_get_changes_guts[8]))
	v936 = *(*int64)(unsafe.Add(mBase, uint32(v935)+104))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+216)) = v770
	*(*int32)(unsafe.Add(mBase, uint32(v27)+220)) = v771
	*(*int64)(unsafe.Add(mBase, uint32(v27)+224)) = v783
	*(*int64)(unsafe.Add(mBase, uint32(v27)+232)) = v784
	*(*int64)(unsafe.Add(mBase, uint32(v27)+240)) = v785
	*(*int32)(unsafe.Add(mBase, uint32(v27)+252)) = v776
	*(*int32)(unsafe.Add(mBase, uint32(v27)+256)) = v773
	*(*int32)(unsafe.Add(mBase, uint32(v27)+260)) = v772
	*(*int32)(unsafe.Add(mBase, uint32(v27)+264)) = v768
	*(*int32)(unsafe.Add(mBase, uint32(v27)+268)) = v769
	*(*int64)(unsafe.Add(mBase, uint32(v27)+272)) = v782
	*(*int32)(unsafe.Add(mBase, uint32(v27)+284)) = v774
	F_XLogBeginRead(m, v933, v936)
	mBase = m.M
	v950 = m.ExcPending
	if v950 != 0 {
		goto L6
	} else {
		goto L105
	}
L105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+220)) = v771
	*(*int32)(unsafe.Add(mBase, uint32(v27)+216)) = v770
	*(*int64)(unsafe.Add(mBase, uint32(v27)+224)) = v783
	*(*int64)(unsafe.Add(mBase, uint32(v27)+232)) = v784
	*(*int64)(unsafe.Add(mBase, uint32(v27)+240)) = v785
	*(*int32)(unsafe.Add(mBase, uint32(v27)+252)) = v776
	*(*int32)(unsafe.Add(mBase, uint32(v27)+256)) = v773
	*(*int32)(unsafe.Add(mBase, uint32(v27)+260)) = v772
	*(*int32)(unsafe.Add(mBase, uint32(v27)+264)) = v768
	*(*int32)(unsafe.Add(mBase, uint32(v27)+268)) = v769
	*(*int64)(unsafe.Add(mBase, uint32(v27)+272)) = v782
	*(*int32)(unsafe.Add(mBase, uint32(v27)+284)) = v774
	F_InvalidateSystemCachesExtended(m)
	mBase = m.M
	v964 = m.ExcPending
	if v964 != 0 {
		goto L6
	} else {
		goto L106
	}
L106:
	;
	v965 = *(*int32)(unsafe.Add(mBase, uint32(v819)+8))
	v966 = *(*int64)(unsafe.Add(mBase, uint32(v965)+40))
	if base.Ui64(v783) <= base.Ui64(v966) {
		v1131 = v966
		goto L107
	} else {
		goto L108
	}
L107:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_pg_logical_slot_get_changes_guts[0])) = v774
	v1135 = int32(0)
	if base.B2i32(l1 == v1135)|base.B2i32(v1131 == int64(0)) == v1135 {
		goto L134
	} else {
		goto L135
	}
L108:
	;
	v982 = v965
	goto L109
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+216)) = v770
	*(*int32)(unsafe.Add(mBase, uint32(v27)+32)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+220)) = v771
	*(*int64)(unsafe.Add(mBase, uint32(v27)+224)) = v783
	*(*int64)(unsafe.Add(mBase, uint32(v27)+232)) = v784
	*(*int64)(unsafe.Add(mBase, uint32(v27)+240)) = v785
	*(*int32)(unsafe.Add(mBase, uint32(v27)+252)) = v776
	*(*int32)(unsafe.Add(mBase, uint32(v27)+256)) = v773
	*(*int32)(unsafe.Add(mBase, uint32(v27)+260)) = v772
	*(*int32)(unsafe.Add(mBase, uint32(v27)+264)) = v768
	*(*int32)(unsafe.Add(mBase, uint32(v27)+268)) = v769
	*(*int64)(unsafe.Add(mBase, uint32(v27)+272)) = v782
	*(*int32)(unsafe.Add(mBase, uint32(v27)+284)) = v774
	v1009 = F_XLogReadRecord(m, v982, v27+int32(32))
	mBase = m.M
	v1010 = m.ExcPending
	if v1010 != 0 {
		goto L6
	} else {
		goto L111
	}
L110:
	;
	v1131 = v1107
	goto L107
L111:
	;
	v1011 = *(*int32)(unsafe.Add(mBase, uint32(v27)+32))
	if v1011 != 0 {
		goto L112
	} else {
		goto L113
	}
L112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+220)) = v771
	*(*int32)(unsafe.Add(mBase, uint32(v27)+216)) = v770
	*(*int64)(unsafe.Add(mBase, uint32(v27)+224)) = v783
	*(*int64)(unsafe.Add(mBase, uint32(v27)+232)) = v784
	*(*int64)(unsafe.Add(mBase, uint32(v27)+240)) = v785
	*(*int32)(unsafe.Add(mBase, uint32(v27)+252)) = v776
	*(*int32)(unsafe.Add(mBase, uint32(v27)+256)) = v773
	*(*int32)(unsafe.Add(mBase, uint32(v27)+260)) = v772
	*(*int32)(unsafe.Add(mBase, uint32(v27)+264)) = v768
	*(*int32)(unsafe.Add(mBase, uint32(v27)+268)) = v769
	*(*int64)(unsafe.Add(mBase, uint32(v27)+272)) = v782
	*(*int32)(unsafe.Add(mBase, uint32(v27)+284)) = v774
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1027 = m.ExcPending
	if v1027 != 0 {
		goto L6
	} else {
		goto L115
	}
L113:
	;
	goto L114
L114:
	;
	if v1009 != 0 {
		goto L118
	} else {
		goto L119
	}
L115:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+216)) = v770
	*(*int32)(unsafe.Add(mBase, uint32(v27)+220)) = v771
	*(*int64)(unsafe.Add(mBase, uint32(v27)+224)) = v783
	*(*int64)(unsafe.Add(mBase, uint32(v27)+232)) = v784
	*(*int64)(unsafe.Add(mBase, uint32(v27)+240)) = v785
	*(*int32)(unsafe.Add(mBase, uint32(v27)+252)) = v776
	*(*int32)(unsafe.Add(mBase, uint32(v27)+256)) = v773
	*(*int32)(unsafe.Add(mBase, uint32(v27)+260)) = v772
	*(*int32)(unsafe.Add(mBase, uint32(v27)+264)) = v768
	*(*int32)(unsafe.Add(mBase, uint32(v27)+268)) = v769
	*(*int64)(unsafe.Add(mBase, uint32(v27)+272)) = v782
	*(*int32)(unsafe.Add(mBase, uint32(v27)+284)) = v774
	v1040 = *(*int32)(unsafe.Add(mBase, uint32(v27)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v27))) = v1040
	F_errmsg_internal(m, int32(_a_F_pg_logical_slot_get_changes_guts_10), v27)
	mBase = m.M
	v1044 = m.ExcPending
	if v1044 != 0 {
		goto L6
	} else {
		goto L116
	}
L116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+220)) = v771
	*(*int32)(unsafe.Add(mBase, uint32(v27)+216)) = v770
	*(*int64)(unsafe.Add(mBase, uint32(v27)+224)) = v783
	*(*int64)(unsafe.Add(mBase, uint32(v27)+232)) = v784
	*(*int64)(unsafe.Add(mBase, uint32(v27)+240)) = v785
	*(*int32)(unsafe.Add(mBase, uint32(v27)+252)) = v776
	*(*int32)(unsafe.Add(mBase, uint32(v27)+256)) = v773
	*(*int32)(unsafe.Add(mBase, uint32(v27)+260)) = v772
	*(*int32)(unsafe.Add(mBase, uint32(v27)+264)) = v768
	*(*int32)(unsafe.Add(mBase, uint32(v27)+268)) = v769
	*(*int64)(unsafe.Add(mBase, uint32(v27)+272)) = v782
	*(*int32)(unsafe.Add(mBase, uint32(v27)+284)) = v774
	F_errfinish(m, int32(_a_F_pg_logical_slot_get_changes_guts_1), int32(259), int32(_a_F_pg_logical_slot_get_changes_guts_2))
	mBase = m.M
	v1061 = m.ExcPending
	if v1061 != 0 {
		goto L6
	} else {
		goto L117
	}
L117:
	;
	goto L3
L118:
	;
	v1062 = *(*int32)(unsafe.Add(mBase, uint32(v819)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+220)) = v771
	*(*int32)(unsafe.Add(mBase, uint32(v27)+216)) = v770
	*(*int64)(unsafe.Add(mBase, uint32(v27)+224)) = v783
	*(*int64)(unsafe.Add(mBase, uint32(v27)+232)) = v784
	*(*int64)(unsafe.Add(mBase, uint32(v27)+240)) = v785
	*(*int32)(unsafe.Add(mBase, uint32(v27)+252)) = v776
	*(*int32)(unsafe.Add(mBase, uint32(v27)+256)) = v773
	*(*int32)(unsafe.Add(mBase, uint32(v27)+260)) = v772
	*(*int32)(unsafe.Add(mBase, uint32(v27)+264)) = v768
	*(*int32)(unsafe.Add(mBase, uint32(v27)+268)) = v769
	*(*int64)(unsafe.Add(mBase, uint32(v27)+272)) = v782
	*(*int32)(unsafe.Add(mBase, uint32(v27)+284)) = v774
	F_LogicalDecodingProcessRecord(m, v819, v1062)
	mBase = m.M
	v1076 = m.ExcPending
	if v1076 != 0 {
		goto L6
	} else {
		goto L121
	}
L119:
	;
	goto L120
L120:
	;
	if v928 == int32(0) {
		goto L122
	} else {
		goto L123
	}
L121:
	;
	goto L120
L122:
	;
	v1080 = *(*int32)(unsafe.Add(mBase, uint32(v819)+8))
	v1081 = *(*int64)(unsafe.Add(mBase, uint32(v1080)+40))
	if base.Ui64(v782) <= base.Ui64(v1081) {
		v1131 = v1081
		goto L107
	} else {
		goto L125
	}
L123:
	;
	goto L124
L124:
	;
	if v769 == int32(0) {
		goto L126
	} else {
		goto L127
	}
L125:
	;
	goto L124
L126:
	;
	v1091 = *(*int32)(unsafe.Add(mBase, _c_F_pg_logical_slot_get_changes_guts[9]))
	if v1091 != 0 {
		goto L129
	} else {
		goto L130
	}
L127:
	;
	v1086 = *(*int64)(unsafe.Add(mBase, uint32(v768)+16))
	if v1086 < base.I64_extend_i32_s(v769) {
		goto L126
	} else {
		goto L128
	}
L128:
	;
	v1088 = *(*int32)(unsafe.Add(mBase, uint32(v819)+8))
	v1089 = *(*int64)(unsafe.Add(mBase, uint32(v1088)+40))
	v1131 = v1089
	goto L107
L129:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+220)) = v771
	*(*int32)(unsafe.Add(mBase, uint32(v27)+216)) = v770
	*(*int64)(unsafe.Add(mBase, uint32(v27)+224)) = v783
	*(*int64)(unsafe.Add(mBase, uint32(v27)+232)) = v784
	*(*int64)(unsafe.Add(mBase, uint32(v27)+240)) = v785
	*(*int32)(unsafe.Add(mBase, uint32(v27)+252)) = v776
	*(*int32)(unsafe.Add(mBase, uint32(v27)+256)) = v773
	*(*int32)(unsafe.Add(mBase, uint32(v27)+260)) = v772
	*(*int32)(unsafe.Add(mBase, uint32(v27)+264)) = v768
	*(*int32)(unsafe.Add(mBase, uint32(v27)+268)) = v769
	*(*int64)(unsafe.Add(mBase, uint32(v27)+272)) = v782
	*(*int32)(unsafe.Add(mBase, uint32(v27)+284)) = v774
	F_ProcessInterrupts(m)
	mBase = m.M
	v1105 = m.ExcPending
	if v1105 != 0 {
		goto L6
	} else {
		goto L132
	}
L130:
	;
	goto L131
L131:
	;
	v1106 = *(*int32)(unsafe.Add(mBase, uint32(v819)+8))
	v1107 = *(*int64)(unsafe.Add(mBase, uint32(v1106)+40))
	if base.Ui64(v1107) < base.Ui64(v783) {
		v982 = v1106
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
	*(*int32)(unsafe.Add(mBase, uint32(v27)+220)) = v771
	*(*int32)(unsafe.Add(mBase, uint32(v27)+216)) = v770
	*(*int64)(unsafe.Add(mBase, uint32(v27)+224)) = v783
	*(*int64)(unsafe.Add(mBase, uint32(v27)+232)) = v784
	*(*int64)(unsafe.Add(mBase, uint32(v27)+240)) = v785
	*(*int32)(unsafe.Add(mBase, uint32(v27)+252)) = v776
	*(*int32)(unsafe.Add(mBase, uint32(v27)+256)) = v773
	*(*int32)(unsafe.Add(mBase, uint32(v27)+260)) = v772
	*(*int32)(unsafe.Add(mBase, uint32(v27)+264)) = v768
	*(*int32)(unsafe.Add(mBase, uint32(v27)+268)) = v769
	*(*int64)(unsafe.Add(mBase, uint32(v27)+272)) = v782
	*(*int32)(unsafe.Add(mBase, uint32(v27)+284)) = v774
	F_LogicalConfirmReceivedLocation(m, v1131)
	mBase = m.M
	v1155 = m.ExcPending
	if v1155 != 0 {
		goto L6
	} else {
		goto L137
	}
L135:
	;
	goto L136
L136:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+220)) = v771
	*(*int32)(unsafe.Add(mBase, uint32(v27)+216)) = v770
	*(*int64)(unsafe.Add(mBase, uint32(v27)+224)) = v783
	*(*int64)(unsafe.Add(mBase, uint32(v27)+232)) = v784
	*(*int64)(unsafe.Add(mBase, uint32(v27)+240)) = v785
	*(*int32)(unsafe.Add(mBase, uint32(v27)+252)) = v776
	*(*int32)(unsafe.Add(mBase, uint32(v27)+256)) = v773
	*(*int32)(unsafe.Add(mBase, uint32(v27)+260)) = v772
	*(*int32)(unsafe.Add(mBase, uint32(v27)+264)) = v768
	*(*int32)(unsafe.Add(mBase, uint32(v27)+268)) = v769
	*(*int64)(unsafe.Add(mBase, uint32(v27)+272)) = v782
	*(*int32)(unsafe.Add(mBase, uint32(v27)+284)) = v774
	F_FreeDecodingContext(m, v819)
	mBase = m.M
	v1183 = m.ExcPending
	if v1183 != 0 {
		goto L6
	} else {
		goto L139
	}
L137:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+220)) = v771
	*(*int32)(unsafe.Add(mBase, uint32(v27)+216)) = v770
	*(*int64)(unsafe.Add(mBase, uint32(v27)+224)) = v783
	*(*int64)(unsafe.Add(mBase, uint32(v27)+232)) = v784
	*(*int64)(unsafe.Add(mBase, uint32(v27)+240)) = v785
	*(*int32)(unsafe.Add(mBase, uint32(v27)+252)) = v776
	*(*int32)(unsafe.Add(mBase, uint32(v27)+256)) = v773
	*(*int32)(unsafe.Add(mBase, uint32(v27)+260)) = v772
	*(*int32)(unsafe.Add(mBase, uint32(v27)+264)) = v768
	*(*int32)(unsafe.Add(mBase, uint32(v27)+268)) = v769
	*(*int64)(unsafe.Add(mBase, uint32(v27)+272)) = v782
	*(*int32)(unsafe.Add(mBase, uint32(v27)+284)) = v774
	F_ReplicationSlotMarkDirty(m)
	mBase = m.M
	v1169 = m.ExcPending
	if v1169 != 0 {
		goto L6
	} else {
		goto L138
	}
L138:
	;
	goto L136
L139:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+220)) = v771
	*(*int32)(unsafe.Add(mBase, uint32(v27)+216)) = v770
	*(*int64)(unsafe.Add(mBase, uint32(v27)+224)) = v783
	*(*int64)(unsafe.Add(mBase, uint32(v27)+232)) = v784
	*(*int64)(unsafe.Add(mBase, uint32(v27)+240)) = v785
	*(*int32)(unsafe.Add(mBase, uint32(v27)+252)) = v776
	*(*int32)(unsafe.Add(mBase, uint32(v27)+256)) = v773
	*(*int32)(unsafe.Add(mBase, uint32(v27)+260)) = v772
	*(*int32)(unsafe.Add(mBase, uint32(v27)+264)) = v768
	*(*int32)(unsafe.Add(mBase, uint32(v27)+268)) = v769
	*(*int64)(unsafe.Add(mBase, uint32(v27)+272)) = v782
	*(*int32)(unsafe.Add(mBase, uint32(v27)+284)) = v774
	F_ReplicationSlotRelease(m)
	mBase = m.M
	v1197 = m.ExcPending
	if v1197 != 0 {
		goto L6
	} else {
		goto L140
	}
L140:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+220)) = v771
	*(*int32)(unsafe.Add(mBase, uint32(v27)+216)) = v770
	*(*int64)(unsafe.Add(mBase, uint32(v27)+224)) = v783
	*(*int64)(unsafe.Add(mBase, uint32(v27)+232)) = v784
	*(*int64)(unsafe.Add(mBase, uint32(v27)+240)) = v785
	*(*int32)(unsafe.Add(mBase, uint32(v27)+252)) = v776
	*(*int32)(unsafe.Add(mBase, uint32(v27)+256)) = v773
	*(*int32)(unsafe.Add(mBase, uint32(v27)+260)) = v772
	*(*int32)(unsafe.Add(mBase, uint32(v27)+264)) = v768
	*(*int32)(unsafe.Add(mBase, uint32(v27)+268)) = v769
	*(*int64)(unsafe.Add(mBase, uint32(v27)+272)) = v782
	*(*int32)(unsafe.Add(mBase, uint32(v27)+284)) = v774
	F_InvalidateSystemCachesExtended(m)
	mBase = m.M
	v1211 = m.ExcPending
	if v1211 != 0 {
		goto L6
	} else {
		goto L141
	}
L141:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_pg_logical_slot_get_changes_guts[6])) = v770
	*(*int32)(unsafe.Add(mBase, _c_F_pg_logical_slot_get_changes_guts[7])) = v771
	m.G0 = v27 + int32(288)
	return
L142:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+220)) = v771
	*(*int32)(unsafe.Add(mBase, uint32(v27)+216)) = v770
	*(*int64)(unsafe.Add(mBase, uint32(v27)+224)) = v783
	*(*int64)(unsafe.Add(mBase, uint32(v27)+232)) = v784
	*(*int64)(unsafe.Add(mBase, uint32(v27)+240)) = v785
	*(*int32)(unsafe.Add(mBase, uint32(v27)+252)) = v776
	*(*int32)(unsafe.Add(mBase, uint32(v27)+256)) = v773
	*(*int32)(unsafe.Add(mBase, uint32(v27)+260)) = v772
	*(*int32)(unsafe.Add(mBase, uint32(v27)+264)) = v768
	*(*int32)(unsafe.Add(mBase, uint32(v27)+268)) = v769
	*(*int64)(unsafe.Add(mBase, uint32(v27)+272)) = v782
	*(*int32)(unsafe.Add(mBase, uint32(v27)+284)) = v774
	F_pg_re_throw(m)
	mBase = m.M
	v1250 = m.ExcPending
	if v1250 != 0 {
		goto L6
	} else {
		goto L143
	}
L143:
	;
	goto L5
L144:
	;
	v1280 = int32(v1276)
	m.G0 = v27
	v1282 = *(*int32)(unsafe.Add(mBase, uint32(v1280)+4))
	v1283 = *(*int32)(unsafe.Add(mBase, uint32(v1280)))
	v1286 = *(*int32)(unsafe.Add(mBase, uint32(v1283)))
	if v27+int32(28) == v1286 {
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
	if v1290 != 0 {
		goto L150
	} else {
		goto L151
	}
L147:
	;
	v1288 = *(*int32)(unsafe.Add(mBase, uint32(v1283)+4))
	v1290 = v1288
	goto L149
L148:
	;
	v1290 = int32(0)
	goto L149
L149:
	;
	goto L146
L150:
	;
	v1291 = *(*int32)(unsafe.Add(mBase, uint32(v27)+284))
	v1292 = *(*int64)(unsafe.Add(mBase, uint32(v27)+272))
	v1293 = *(*int32)(unsafe.Add(mBase, uint32(v27)+268))
	v1294 = *(*int32)(unsafe.Add(mBase, uint32(v27)+264))
	v1295 = *(*int32)(unsafe.Add(mBase, uint32(v27)+260))
	v1296 = *(*int32)(unsafe.Add(mBase, uint32(v27)+256))
	v1297 = *(*int32)(unsafe.Add(mBase, uint32(v27)+252))
	v1298 = *(*int64)(unsafe.Add(mBase, uint32(v27)+240))
	v1299 = *(*int64)(unsafe.Add(mBase, uint32(v27)+232))
	v1300 = *(*int64)(unsafe.Add(mBase, uint32(v27)+224))
	v1301 = *(*int32)(unsafe.Add(mBase, uint32(v27)+220))
	v1302 = *(*int32)(unsafe.Add(mBase, uint32(v27)+216))
	v34 = v1294
	v35 = v1293
	v36 = v1302
	v37 = v1301
	v38 = v1295
	v39 = v1296
	v40 = v1291
	v41 = v1297
	v42 = v1290
	v43 = v1282
	v48 = v1292
	v49 = v1300
	v50 = v1299
	v51 = v1298
	goto L1
L151:
	;
	goto L152
L152:
	;
	F___wasm_longjmp(m, v1283, v1282)
	mBase = m.M
	v1304 = m.ExcPending
	if v1304 != 0 {
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
