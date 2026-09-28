package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_pg_logical_emit_message_bytea(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v10 int64
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
	var v21 int64
	_ = v21
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
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
	var v97 int32
	_ = v97
	var v103 int32
	_ = v103
	v10 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v12 = F_pg_detoast_datum_packed(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int64(0)
	} else {
		v16 = F_text_to_cstring(m, v12)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int64(0)
		} else {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
			v19 = F_pg_detoast_datum_packed(m, v18)
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int64(0)
			} else {
				v21 = *(*int64)(unsafe.Add(mBase, uint32(l0)+72))
				v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
				if v22 == int32(1) {
					v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+1)))
					if v28 == int32(18) {
						v31 = int32(16)
					} else {
						v31 = int32(0)
					}
					if base.Ui32((v28-int32(1))&int32(255)) < base.Ui32(int32(3)) {
						v38 = int32(4)
					} else {
						v38 = v31
					}
					v51 = v38
				} else {
					v39 = int32(1)
					if v22&v39 != 0 {
						v51 = int32(base.Ui32(v22)>>(uint(v39)%32)) - v39
					} else {
						v45 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
						v51 = int32(base.Ui32(v45)>>(uint(int32(2))%32)) - int32(4)
					}
				}
				v52 = m.G0
				v54 = v52 - int32(16)
				m.G0 = v54
				v57 = base.B2i32(v10 != int64(0))
				if v10 != int64(0) {
					v58 = F_GetCurrentTransactionId(m)
					mBase = m.M
					v59 = m.ExcPending
					if v59 != 0 {
						return int64(0)
					} else {
						v60 = int32(1)
						if v22&v60 != 0 {
							v64 = v60
						} else {
							v64 = int32(4)
						}
						*(*uint8)(unsafe.Add(mBase, uint32(v54)+4)) = uint8(v57)
						v70 = *(*int32)(unsafe.Add(mBase, _c_F_pg_logical_emit_message_bytea[0]))
						*(*int32)(unsafe.Add(mBase, uint32(v54))) = v70
						v72 = F_strlen(m, v16)
						mBase = m.M
						*(*int32)(unsafe.Add(mBase, uint32(v54)+12)) = v51
						*(*int32)(unsafe.Add(mBase, uint32(v54)+8)) = v72 + int32(1)
						F_XLogBeginInsert(m)
						mBase = m.M
						v78 = m.ExcPending
						if v78 != 0 {
							return int64(0)
						} else {
							F_XLogRegisterData(m, v54, int32(16))
							mBase = m.M
							v81 = m.ExcPending
							if v81 != 0 {
								return int64(0)
							} else {
								v82 = *(*int32)(unsafe.Add(mBase, uint32(v54)+8))
								F_XLogRegisterData(m, v16, v82)
								mBase = m.M
								v84 = m.ExcPending
								if v84 != 0 {
									return int64(0)
								} else {
									F_XLogRegisterData(m, v19+v64, v51)
									mBase = m.M
									v86 = m.ExcPending
									if v86 != 0 {
										return int64(0)
									} else {
										v88 = int32(_a_F_pg_logical_emit_message_bytea_0)
										v90 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_logical_emit_message_bytea[1])))
										v91 = v90 | int32(1)
										*(*uint8)(unsafe.Add(mBase, _c_F_pg_logical_emit_message_bytea[1])) = uint8(v91)
										v95 = F_XLogInsert(m, int32(21), int32(0))
										mBase = m.M
										v96 = m.ExcPending
										if v96 != 0 {
											return int64(0)
										} else {
											v97 = int32(0)
											if base.B2i32(base.B2i32(v21 != int64(0)) == v97)|v57 == v97 {
												F_XLogFlush(m, v95)
												mBase = m.M
												v103 = m.ExcPending
												if v103 != 0 {
													return int64(0)
												} else {
													m.G0 = v54 + int32(16)
													return v95
												}
											} else {
												m.G0 = v54 + int32(16)
												return v95
											}
										}
									}
								}
							}
						}
					}
				} else {
					v60 = int32(1)
					if v22&v60 != 0 {
						v64 = v60
					} else {
						v64 = int32(4)
					}
					*(*uint8)(unsafe.Add(mBase, uint32(v54)+4)) = uint8(v57)
					v70 = *(*int32)(unsafe.Add(mBase, _c_F_pg_logical_emit_message_bytea[0]))
					*(*int32)(unsafe.Add(mBase, uint32(v54))) = v70
					v72 = F_strlen(m, v16)
					mBase = m.M
					*(*int32)(unsafe.Add(mBase, uint32(v54)+12)) = v51
					*(*int32)(unsafe.Add(mBase, uint32(v54)+8)) = v72 + int32(1)
					F_XLogBeginInsert(m)
					mBase = m.M
					v78 = m.ExcPending
					if v78 != 0 {
						return int64(0)
					} else {
						F_XLogRegisterData(m, v54, int32(16))
						mBase = m.M
						v81 = m.ExcPending
						if v81 != 0 {
							return int64(0)
						} else {
							v82 = *(*int32)(unsafe.Add(mBase, uint32(v54)+8))
							F_XLogRegisterData(m, v16, v82)
							mBase = m.M
							v84 = m.ExcPending
							if v84 != 0 {
								return int64(0)
							} else {
								F_XLogRegisterData(m, v19+v64, v51)
								mBase = m.M
								v86 = m.ExcPending
								if v86 != 0 {
									return int64(0)
								} else {
									v88 = int32(_a_F_pg_logical_emit_message_bytea_0)
									v90 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_logical_emit_message_bytea[1])))
									v91 = v90 | int32(1)
									*(*uint8)(unsafe.Add(mBase, _c_F_pg_logical_emit_message_bytea[1])) = uint8(v91)
									v95 = F_XLogInsert(m, int32(21), int32(0))
									mBase = m.M
									v96 = m.ExcPending
									if v96 != 0 {
										return int64(0)
									} else {
										v97 = int32(0)
										if base.B2i32(base.B2i32(v21 != int64(0)) == v97)|v57 == v97 {
											F_XLogFlush(m, v95)
											mBase = m.M
											v103 = m.ExcPending
											if v103 != 0 {
												return int64(0)
											} else {
												m.G0 = v54 + int32(16)
												return v95
											}
										} else {
											m.G0 = v54 + int32(16)
											return v95
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
	var v18 int64
	_ = v18
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
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
	var v41 int32
	_ = v41
	var v46 int64
	_ = v46
	var v47 int64
	_ = v47
	var v48 int64
	_ = v48
	var v49 int64
	_ = v49
	var v54 int32
	_ = v54
	var v67 int32
	_ = v67
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v99 int32
	_ = v99
	var v113 int32
	_ = v113
	var v128 int32
	_ = v128
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v149 int64
	_ = v149
	var v150 int64
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v174 int32
	_ = v174
	var v188 int32
	_ = v188
	var v203 int32
	_ = v203
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v273 int32
	_ = v273
	var v287 int32
	_ = v287
	var v302 int32
	_ = v302
	var v318 int32
	_ = v318
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v346 int32
	_ = v346
	var v360 int32
	_ = v360
	var v375 int32
	_ = v375
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v417 int32
	_ = v417
	var v429 int32
	_ = v429
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v447 int32
	_ = v447
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v463 int32
	_ = v463
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v530 int32
	_ = v530
	var v532 int32
	_ = v532
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v560 int32
	_ = v560
	var v575 int32
	_ = v575
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v583 int32
	_ = v583
	var v585 int32
	_ = v585
	var v597 int32
	_ = v597
	var v599 int32
	_ = v599
	var v600 int32
	_ = v600
	var v601 int64
	_ = v601
	var v604 int64
	_ = v604
	var v609 int32
	_ = v609
	var v612 int32
	_ = v612
	var v616 int64
	_ = v616
	var v623 int64
	_ = v623
	var v638 int32
	_ = v638
	var v652 int32
	_ = v652
	var v667 int32
	_ = v667
	var v683 int32
	_ = v683
	var v696 int64
	_ = v696
	var v697 int32
	_ = v697
	var v698 int64
	_ = v698
	var v699 int64
	_ = v699
	var v700 int64
	_ = v700
	var v712 int32
	_ = v712
	var v715 int32
	_ = v715
	var v718 int32
	_ = v718
	var v720 int32
	_ = v720
	var v722 int32
	_ = v722
	var v732 int32
	_ = v732
	var v733 int32
	_ = v733
	var v734 int32
	_ = v734
	var v735 int32
	_ = v735
	var v736 int32
	_ = v736
	var v737 int32
	_ = v737
	var v739 int32
	_ = v739
	var v740 int32
	_ = v740
	var v745 int64
	_ = v745
	var v746 int64
	_ = v746
	var v747 int64
	_ = v747
	var v748 int64
	_ = v748
	var v775 int32
	_ = v775
	var v781 int32
	_ = v781
	var v782 int32
	_ = v782
	var v785 int32
	_ = v785
	var v802 int32
	_ = v802
	var v816 int32
	_ = v816
	var v817 int32
	_ = v817
	var v818 int32
	_ = v818
	var v831 int32
	_ = v831
	var v832 int32
	_ = v832
	var v833 int32
	_ = v833
	var v853 int32
	_ = v853
	var v869 int32
	_ = v869
	var v882 int64
	_ = v882
	var v884 int32
	_ = v884
	var v885 int64
	_ = v885
	var v887 int32
	_ = v887
	var v889 int32
	_ = v889
	var v891 int32
	_ = v891
	var v892 int64
	_ = v892
	var v905 int32
	_ = v905
	var v918 int32
	_ = v918
	var v919 int32
	_ = v919
	var v920 int64
	_ = v920
	var v935 int32
	_ = v935
	var v961 int32
	_ = v961
	var v962 int32
	_ = v962
	var v963 int32
	_ = v963
	var v978 int32
	_ = v978
	var v990 int32
	_ = v990
	var v994 int32
	_ = v994
	var v1010 int32
	_ = v1010
	var v1011 int32
	_ = v1011
	var v1024 int32
	_ = v1024
	var v1028 int32
	_ = v1028
	var v1029 int64
	_ = v1029
	var v1034 int64
	_ = v1034
	var v1036 int32
	_ = v1036
	var v1037 int64
	_ = v1037
	var v1039 int32
	_ = v1039
	var v1052 int32
	_ = v1052
	var v1053 int32
	_ = v1053
	var v1054 int64
	_ = v1054
	var v1077 int64
	_ = v1077
	var v1079 int32
	_ = v1079
	var v1098 int32
	_ = v1098
	var v1111 int32
	_ = v1111
	var v1124 int32
	_ = v1124
	var v1137 int32
	_ = v1137
	var v1150 int32
	_ = v1150
	var v1174 int32
	_ = v1174
	var v1187 int32
	_ = v1187
	var v1211 int32
	_ = v1211
	var v1212 int64
	_ = v1212
	var v1216 int32
	_ = v1216
	var v1218 int32
	_ = v1218
	var v1219 int32
	_ = v1219
	var v1222 int32
	_ = v1222
	var v1224 int32
	_ = v1224
	var v1226 int32
	_ = v1226
	var v1227 int64
	_ = v1227
	var v1228 int32
	_ = v1228
	var v1229 int32
	_ = v1229
	var v1230 int32
	_ = v1230
	var v1231 int32
	_ = v1231
	var v1232 int32
	_ = v1232
	var v1233 int64
	_ = v1233
	var v1234 int64
	_ = v1234
	var v1235 int64
	_ = v1235
	var v1236 int32
	_ = v1236
	var v1237 int32
	_ = v1237
	var v1239 int32
	_ = v1239
	v3 = l2
	v4 = int32(0)
	v18 = int64(0)
	v24 = m.G0
	v26 = v24 - int32(288)
	m.G0 = v26
	v33 = v4
	v34 = v4
	v35 = v4
	v36 = v4
	v37 = v4
	v38 = v4
	v39 = v4
	v40 = int32(-1)
	v41 = v4
	v46 = v18
	v47 = v18
	v48 = v18
	v49 = v18
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
	if v40 == int32(1) {
		goto L8
	} else {
		goto L9
	}
L5:
	;
	goto L3
L6:
	;
	v1211 = int32(m.ExcTag)
	v1212 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v1211 == int32(0) {
		goto L144
	} else {
		goto L145
	}
L7:
	;
	if v740 == int32(0) {
		goto L86
	} else {
		goto L87
	}
L8:
	;
	v732 = v33
	v733 = v34
	v734 = v35
	v735 = v36
	v736 = v37
	v737 = v38
	v739 = v39
	v740 = v41
	v745 = v46
	v746 = v47
	v747 = v48
	v748 = v49
	goto L7
L9:
	;
	goto L10
L10:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+228)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v26)+224)) = v36
	*(*int64)(unsafe.Add(mBase, uint32(v26)+232)) = v47
	*(*int64)(unsafe.Add(mBase, uint32(v26)+240)) = v48
	*(*int64)(unsafe.Add(mBase, uint32(v26)+248)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v26)+260)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v26)+264)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v26)+268)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v26)+272)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v26)+276)) = v34
	*(*int64)(unsafe.Add(mBase, uint32(v26)+280)) = v46
	F_CheckSlotPermissions(m)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L6
	} else {
		goto L11
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+228)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v26)+224)) = v36
	*(*int64)(unsafe.Add(mBase, uint32(v26)+232)) = v47
	*(*int64)(unsafe.Add(mBase, uint32(v26)+240)) = v48
	*(*int64)(unsafe.Add(mBase, uint32(v26)+248)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v26)+260)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v26)+264)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v26)+268)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v26)+272)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v26)+276)) = v34
	*(*int64)(unsafe.Add(mBase, uint32(v26)+280)) = v46
	F_CheckLogicalDecodingRequirements(m, int32(0))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L6
	} else {
		goto L12
	}
L12:
	;
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v82 == int32(1) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+228)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v26)+224)) = v36
	*(*int64)(unsafe.Add(mBase, uint32(v26)+232)) = v47
	*(*int64)(unsafe.Add(mBase, uint32(v26)+240)) = v48
	*(*int64)(unsafe.Add(mBase, uint32(v26)+248)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v26)+260)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v26)+264)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v26)+268)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v26)+272)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v26)+276)) = v34
	*(*int64)(unsafe.Add(mBase, uint32(v26)+280)) = v46
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L6
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)))
	if v146 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+228)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v26)+224)) = v36
	*(*int64)(unsafe.Add(mBase, uint32(v26)+232)) = v47
	*(*int64)(unsafe.Add(mBase, uint32(v26)+240)) = v48
	*(*int64)(unsafe.Add(mBase, uint32(v26)+248)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v26)+260)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v26)+264)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v26)+268)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v26)+272)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v26)+276)) = v34
	*(*int64)(unsafe.Add(mBase, uint32(v26)+280)) = v46
	F_errcode(m, int32(67108994))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L6
	} else {
		goto L17
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+228)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v26)+224)) = v36
	*(*int64)(unsafe.Add(mBase, uint32(v26)+232)) = v47
	*(*int64)(unsafe.Add(mBase, uint32(v26)+240)) = v48
	*(*int64)(unsafe.Add(mBase, uint32(v26)+248)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v26)+260)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v26)+264)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v26)+268)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v26)+272)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v26)+276)) = v34
	*(*int64)(unsafe.Add(mBase, uint32(v26)+280)) = v46
	F_errmsg(m, int32(_a_F_pg_logical_slot_get_changes_guts_0), int32(0))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L6
	} else {
		goto L18
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+228)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v26)+224)) = v36
	*(*int64)(unsafe.Add(mBase, uint32(v26)+232)) = v47
	*(*int64)(unsafe.Add(mBase, uint32(v26)+240)) = v48
	*(*int64)(unsafe.Add(mBase, uint32(v26)+248)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v26)+260)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v26)+264)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v26)+268)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v26)+272)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v26)+276)) = v34
	*(*int64)(unsafe.Add(mBase, uint32(v26)+280)) = v46
	F_errfinish(m, int32(_a_F_pg_logical_slot_get_changes_guts_1), int32(123), int32(_a_F_pg_logical_slot_get_changes_guts_2))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L6
	} else {
		goto L19
	}
L19:
	;
	goto L3
L20:
	;
	v149 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v150 = v149
	goto L22
L21:
	;
	v150 = int64(0)
	goto L22
L22:
	;
	v151 = int32(0)
	v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+64)))
	if v152 == v151 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v156 = v155
	goto L25
L24:
	;
	v156 = v151
	goto L25
L25:
	;
	v157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+80)))
	if v157 == int32(1) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+228)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v26)+224)) = v36
	*(*int64)(unsafe.Add(mBase, uint32(v26)+232)) = v47
	*(*int64)(unsafe.Add(mBase, uint32(v26)+240)) = v48
	*(*int64)(unsafe.Add(mBase, uint32(v26)+248)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v26)+260)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v26)+264)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v26)+268)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v26)+272)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v26)+276)) = v156
	*(*int64)(unsafe.Add(mBase, uint32(v26)+280)) = v150
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L6
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v221 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+228)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v26)+224)) = v36
	*(*int64)(unsafe.Add(mBase, uint32(v26)+232)) = v47
	*(*int64)(unsafe.Add(mBase, uint32(v26)+240)) = v48
	*(*int64)(unsafe.Add(mBase, uint32(v26)+248)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v26)+260)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v26)+264)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v26)+268)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v26)+272)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v26)+276)) = v156
	*(*int64)(unsafe.Add(mBase, uint32(v26)+280)) = v150
	v233 = F_pg_detoast_datum(m, v221)
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L6
	} else {
		goto L33
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+228)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v26)+224)) = v36
	*(*int64)(unsafe.Add(mBase, uint32(v26)+232)) = v47
	*(*int64)(unsafe.Add(mBase, uint32(v26)+240)) = v48
	*(*int64)(unsafe.Add(mBase, uint32(v26)+248)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v26)+260)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v26)+264)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v26)+268)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v26)+272)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v26)+276)) = v156
	*(*int64)(unsafe.Add(mBase, uint32(v26)+280)) = v150
	F_errcode(m, int32(67108994))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L6
	} else {
		goto L30
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+228)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v26)+224)) = v36
	*(*int64)(unsafe.Add(mBase, uint32(v26)+232)) = v47
	*(*int64)(unsafe.Add(mBase, uint32(v26)+240)) = v48
	*(*int64)(unsafe.Add(mBase, uint32(v26)+248)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v26)+260)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v26)+264)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v26)+268)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v26)+272)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v26)+276)) = v156
	*(*int64)(unsafe.Add(mBase, uint32(v26)+280)) = v150
	F_errmsg(m, int32(_a_F_pg_logical_slot_get_changes_guts_3), int32(0))
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L6
	} else {
		goto L31
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+228)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v26)+224)) = v36
	*(*int64)(unsafe.Add(mBase, uint32(v26)+232)) = v47
	*(*int64)(unsafe.Add(mBase, uint32(v26)+240)) = v48
	*(*int64)(unsafe.Add(mBase, uint32(v26)+248)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v26)+260)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v26)+264)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v26)+268)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v26)+272)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v26)+276)) = v156
	*(*int64)(unsafe.Add(mBase, uint32(v26)+280)) = v150
	F_errfinish(m, int32(_a_F_pg_logical_slot_get_changes_guts_1), int32(139), int32(_a_F_pg_logical_slot_get_changes_guts_2))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L6
	} else {
		goto L32
	}
L32:
	;
	goto L3
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+228)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v26)+224)) = v36
	*(*int64)(unsafe.Add(mBase, uint32(v26)+232)) = v47
	*(*int64)(unsafe.Add(mBase, uint32(v26)+240)) = v48
	*(*int64)(unsafe.Add(mBase, uint32(v26)+248)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v26)+260)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v26)+264)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v26)+268)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v26)+272)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v26)+276)) = v156
	*(*int64)(unsafe.Add(mBase, uint32(v26)+280)) = v150
	v247 = F_palloc0(m, int32(24))
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L6
	} else {
		goto L34
	}
L34:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v247)+8)) = uint8(v3)
	v250 = int32(_a_F_pg_logical_slot_get_changes_guts_4)
	v251 = *(*int32)(unsafe.Add(mBase, _c_F_pg_logical_slot_get_changes_guts[0]))
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v54)+4))
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v253)+16))
	*(*int32)(unsafe.Add(mBase, _c_F_pg_logical_slot_get_changes_guts[0])) = v254
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v233)+4))
	if base.Ui32(int32(2)) <= base.Ui32(v256) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+228)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v26)+224)) = v36
	*(*int64)(unsafe.Add(mBase, uint32(v26)+232)) = v47
	*(*int64)(unsafe.Add(mBase, uint32(v26)+240)) = v48
	*(*int64)(unsafe.Add(mBase, uint32(v26)+248)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v26)+260)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v26)+264)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v26)+268)) = v251
	*(*int32)(unsafe.Add(mBase, uint32(v26)+272)) = v247
	*(*int32)(unsafe.Add(mBase, uint32(v26)+276)) = v156
	*(*int64)(unsafe.Add(mBase, uint32(v26)+280)) = v150
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L6
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+228)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v26)+224)) = v36
	*(*int64)(unsafe.Add(mBase, uint32(v26)+232)) = v47
	*(*int64)(unsafe.Add(mBase, uint32(v26)+240)) = v48
	*(*int64)(unsafe.Add(mBase, uint32(v26)+248)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v26)+260)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v26)+264)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v26)+268)) = v251
	*(*int32)(unsafe.Add(mBase, uint32(v26)+272)) = v247
	*(*int32)(unsafe.Add(mBase, uint32(v26)+276)) = v156
	*(*int64)(unsafe.Add(mBase, uint32(v26)+280)) = v150
	v330 = F_array_contains_nulls(m, v233)
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L6
	} else {
		goto L42
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+228)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v26)+224)) = v36
	*(*int64)(unsafe.Add(mBase, uint32(v26)+232)) = v47
	*(*int64)(unsafe.Add(mBase, uint32(v26)+240)) = v48
	*(*int64)(unsafe.Add(mBase, uint32(v26)+248)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v26)+260)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v26)+264)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v26)+268)) = v251
	*(*int32)(unsafe.Add(mBase, uint32(v26)+272)) = v247
	*(*int32)(unsafe.Add(mBase, uint32(v26)+276)) = v156
	*(*int64)(unsafe.Add(mBase, uint32(v26)+280)) = v150
	F_errcode(m, int32(1088))
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L6
	} else {
		goto L39
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+228)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v26)+224)) = v36
	*(*int64)(unsafe.Add(mBase, uint32(v26)+232)) = v47
	*(*int64)(unsafe.Add(mBase, uint32(v26)+240)) = v48
	*(*int64)(unsafe.Add(mBase, uint32(v26)+248)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v26)+260)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v26)+264)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v26)+268)) = v251
	*(*int32)(unsafe.Add(mBase, uint32(v26)+272)) = v247
	*(*int32)(unsafe.Add(mBase, uint32(v26)+276)) = v156
	*(*int64)(unsafe.Add(mBase, uint32(v26)+280)) = v150
	F_errmsg(m, int32(_a_F_pg_logical_slot_get_changes_guts_5), int32(0))
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L6
	} else {
		goto L40
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+228)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v26)+224)) = v36
	*(*int64)(unsafe.Add(mBase, uint32(v26)+232)) = v47
	*(*int64)(unsafe.Add(mBase, uint32(v26)+240)) = v48
	*(*int64)(unsafe.Add(mBase, uint32(v26)+248)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v26)+260)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v26)+264)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v26)+268)) = v251
	*(*int32)(unsafe.Add(mBase, uint32(v26)+272)) = v247
	*(*int32)(unsafe.Add(mBase, uint32(v26)+276)) = v156
	*(*int64)(unsafe.Add(mBase, uint32(v26)+280)) = v150
	F_errfinish(m, int32(_a_F_pg_logical_slot_get_changes_guts_1), int32(156), int32(_a_F_pg_logical_slot_get_changes_guts_2))
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L6
	} else {
		goto L41
	}
L41:
	;
	goto L3
L42:
	;
	if v330 != 0 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+228)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v26)+224)) = v36
	*(*int64)(unsafe.Add(mBase, uint32(v26)+232)) = v47
	*(*int64)(unsafe.Add(mBase, uint32(v26)+240)) = v48
	*(*int64)(unsafe.Add(mBase, uint32(v26)+248)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v26)+260)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v26)+264)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v26)+268)) = v251
	*(*int32)(unsafe.Add(mBase, uint32(v26)+272)) = v247
	*(*int32)(unsafe.Add(mBase, uint32(v26)+276)) = v156
	*(*int64)(unsafe.Add(mBase, uint32(v26)+280)) = v150
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L6
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	v392 = int32(0)
	if v256 != int32(1) {
		v530 = v38
		v532 = v392
		goto L53
	} else {
		goto L54
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+228)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v26)+224)) = v36
	*(*int64)(unsafe.Add(mBase, uint32(v26)+232)) = v47
	*(*int64)(unsafe.Add(mBase, uint32(v26)+240)) = v48
	*(*int64)(unsafe.Add(mBase, uint32(v26)+248)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v26)+260)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v26)+264)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v26)+268)) = v251
	*(*int32)(unsafe.Add(mBase, uint32(v26)+272)) = v247
	*(*int32)(unsafe.Add(mBase, uint32(v26)+276)) = v156
	*(*int64)(unsafe.Add(mBase, uint32(v26)+280)) = v150
	F_errcode(m, int32(1088))
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L6
	} else {
		goto L47
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+228)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v26)+224)) = v36
	*(*int64)(unsafe.Add(mBase, uint32(v26)+232)) = v47
	*(*int64)(unsafe.Add(mBase, uint32(v26)+240)) = v48
	*(*int64)(unsafe.Add(mBase, uint32(v26)+248)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v26)+260)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v26)+264)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v26)+268)) = v251
	*(*int32)(unsafe.Add(mBase, uint32(v26)+272)) = v247
	*(*int32)(unsafe.Add(mBase, uint32(v26)+276)) = v156
	*(*int64)(unsafe.Add(mBase, uint32(v26)+280)) = v150
	F_errmsg(m, int32(_a_F_pg_logical_slot_get_changes_guts_6), int32(0))
	mBase = m.M
	v375 = m.ExcPending
	if v375 != 0 {
		goto L6
	} else {
		goto L48
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+228)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v26)+224)) = v36
	*(*int64)(unsafe.Add(mBase, uint32(v26)+232)) = v47
	*(*int64)(unsafe.Add(mBase, uint32(v26)+240)) = v48
	*(*int64)(unsafe.Add(mBase, uint32(v26)+248)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v26)+260)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v26)+264)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v26)+268)) = v251
	*(*int32)(unsafe.Add(mBase, uint32(v26)+272)) = v247
	*(*int32)(unsafe.Add(mBase, uint32(v26)+276)) = v156
	*(*int64)(unsafe.Add(mBase, uint32(v26)+280)) = v150
	F_errfinish(m, int32(_a_F_pg_logical_slot_get_changes_guts_1), int32(162), int32(_a_F_pg_logical_slot_get_changes_guts_2))
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L6
	} else {
		goto L49
	}
L49:
	;
	goto L3
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+228)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v26)+224)) = v36
	*(*int64)(unsafe.Add(mBase, uint32(v26)+232)) = v700
	*(*int64)(unsafe.Add(mBase, uint32(v26)+240)) = v698
	*(*int64)(unsafe.Add(mBase, uint32(v26)+248)) = v699
	*(*int32)(unsafe.Add(mBase, uint32(v26)+260)) = v532
	*(*int32)(unsafe.Add(mBase, uint32(v26)+264)) = v530
	*(*int32)(unsafe.Add(mBase, uint32(v26)+268)) = v251
	*(*int32)(unsafe.Add(mBase, uint32(v26)+272)) = v247
	*(*int32)(unsafe.Add(mBase, uint32(v26)+276)) = v156
	*(*int64)(unsafe.Add(mBase, uint32(v26)+280)) = v150
	v712 = int32(1)
	F_ReplicationSlotAcquire(m, v220, v712, v712)
	mBase = m.M
	v715 = m.ExcPending
	if v715 != 0 {
		goto L6
	} else {
		goto L81
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+228)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v26)+224)) = v36
	*(*int64)(unsafe.Add(mBase, uint32(v26)+232)) = v47
	*(*int64)(unsafe.Add(mBase, uint32(v26)+240)) = v48
	*(*int64)(unsafe.Add(mBase, uint32(v26)+248)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v26)+260)) = v532
	*(*int32)(unsafe.Add(mBase, uint32(v26)+264)) = v530
	*(*int32)(unsafe.Add(mBase, uint32(v26)+268)) = v251
	*(*int32)(unsafe.Add(mBase, uint32(v26)+272)) = v247
	*(*int32)(unsafe.Add(mBase, uint32(v26)+276)) = v156
	*(*int64)(unsafe.Add(mBase, uint32(v26)+280)) = v150
	v696 = F_GetXLogReplayRecPtr(m, int32(0))
	mBase = m.M
	v697 = m.ExcPending
	if v697 != 0 {
		goto L6
	} else {
		goto L80
	}
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+228)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v26)+224)) = v36
	*(*int64)(unsafe.Add(mBase, uint32(v26)+232)) = v47
	*(*int64)(unsafe.Add(mBase, uint32(v26)+240)) = v48
	*(*int64)(unsafe.Add(mBase, uint32(v26)+248)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v26)+260)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v26)+264)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v26)+268)) = v251
	*(*int32)(unsafe.Add(mBase, uint32(v26)+272)) = v247
	*(*int32)(unsafe.Add(mBase, uint32(v26)+276)) = v156
	*(*int64)(unsafe.Add(mBase, uint32(v26)+280)) = v150
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v638 = m.ExcPending
	if v638 != 0 {
		goto L6
	} else {
		goto L76
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+228)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v26)+224)) = v36
	*(*int64)(unsafe.Add(mBase, uint32(v26)+232)) = v47
	*(*int64)(unsafe.Add(mBase, uint32(v26)+240)) = v48
	*(*int64)(unsafe.Add(mBase, uint32(v26)+248)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v26)+260)) = v532
	*(*int32)(unsafe.Add(mBase, uint32(v26)+264)) = v530
	*(*int32)(unsafe.Add(mBase, uint32(v26)+268)) = v251
	*(*int32)(unsafe.Add(mBase, uint32(v26)+272)) = v247
	*(*int32)(unsafe.Add(mBase, uint32(v26)+276)) = v156
	*(*int64)(unsafe.Add(mBase, uint32(v26)+280)) = v150
	F_InitMaterializedSRF(m, l0, int32(0))
	mBase = m.M
	v557 = m.ExcPending
	if v557 != 0 {
		goto L6
	} else {
		goto L66
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+224)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v26)+228)) = v35
	*(*int64)(unsafe.Add(mBase, uint32(v26)+232)) = v47
	*(*int64)(unsafe.Add(mBase, uint32(v26)+240)) = v48
	*(*int64)(unsafe.Add(mBase, uint32(v26)+248)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v26)+260)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v26)+264)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v26)+268)) = v251
	*(*int32)(unsafe.Add(mBase, uint32(v26)+272)) = v247
	*(*int32)(unsafe.Add(mBase, uint32(v26)+276)) = v156
	*(*int64)(unsafe.Add(mBase, uint32(v26)+280)) = v150
	F_deconstruct_array_builtin(m, v233, int32(25), v26+int32(216), int32(0), v26+int32(220))
	mBase = m.M
	v413 = m.ExcPending
	if v413 != 0 {
		goto L6
	} else {
		goto L55
	}
L55:
	;
	v414 = *(*int32)(unsafe.Add(mBase, uint32(v26)+220))
	if v414&int32(1) != 0 {
		goto L52
	} else {
		goto L56
	}
L56:
	;
	v417 = int32(0)
	if v414 <= v417 {
		v530 = v38
		v532 = v392
		goto L53
	} else {
		goto L57
	}
L57:
	;
	v429 = v38
	v431 = v392
	v432 = v417
	goto L58
L58:
	;
	v444 = v432 << (uint(int32(3)) % 32)
	v445 = *(*int32)(unsafe.Add(mBase, uint32(v26)+216))
	v447 = *(*int32)(unsafe.Add(mBase, uint32(v444+v445)))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+224)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v26)+228)) = v35
	*(*int64)(unsafe.Add(mBase, uint32(v26)+232)) = v47
	*(*int64)(unsafe.Add(mBase, uint32(v26)+240)) = v48
	*(*int64)(unsafe.Add(mBase, uint32(v26)+248)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v26)+260)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v26)+264)) = v429
	*(*int32)(unsafe.Add(mBase, uint32(v26)+268)) = v251
	*(*int32)(unsafe.Add(mBase, uint32(v26)+272)) = v247
	*(*int32)(unsafe.Add(mBase, uint32(v26)+276)) = v156
	*(*int64)(unsafe.Add(mBase, uint32(v26)+280)) = v150
	v459 = F_text_to_cstring(m, v447)
	mBase = m.M
	v460 = m.ExcPending
	if v460 != 0 {
		goto L6
	} else {
		goto L60
	}
L59:
	;
	v530 = v515
	v532 = v515
	goto L53
L60:
	;
	v461 = *(*int32)(unsafe.Add(mBase, uint32(v26)+216))
	v463 = *(*int32)(unsafe.Add(mBase, uint32(v461+v444)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+224)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v26)+228)) = v35
	*(*int64)(unsafe.Add(mBase, uint32(v26)+232)) = v47
	*(*int64)(unsafe.Add(mBase, uint32(v26)+240)) = v48
	*(*int64)(unsafe.Add(mBase, uint32(v26)+248)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v26)+260)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v26)+264)) = v429
	*(*int32)(unsafe.Add(mBase, uint32(v26)+268)) = v251
	*(*int32)(unsafe.Add(mBase, uint32(v26)+272)) = v247
	*(*int32)(unsafe.Add(mBase, uint32(v26)+276)) = v156
	*(*int64)(unsafe.Add(mBase, uint32(v26)+280)) = v150
	v475 = F_text_to_cstring(m, v463)
	mBase = m.M
	v476 = m.ExcPending
	if v476 != 0 {
		goto L6
	} else {
		goto L61
	}
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+228)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v26)+224)) = v36
	*(*int64)(unsafe.Add(mBase, uint32(v26)+232)) = v47
	*(*int64)(unsafe.Add(mBase, uint32(v26)+240)) = v48
	*(*int64)(unsafe.Add(mBase, uint32(v26)+248)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v26)+260)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v26)+264)) = v429
	*(*int32)(unsafe.Add(mBase, uint32(v26)+268)) = v251
	*(*int32)(unsafe.Add(mBase, uint32(v26)+272)) = v247
	*(*int32)(unsafe.Add(mBase, uint32(v26)+276)) = v156
	*(*int64)(unsafe.Add(mBase, uint32(v26)+280)) = v150
	v488 = F_makeString(m, v475)
	mBase = m.M
	v489 = m.ExcPending
	if v489 != 0 {
		goto L6
	} else {
		goto L62
	}
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+228)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v26)+224)) = v36
	*(*int64)(unsafe.Add(mBase, uint32(v26)+232)) = v47
	*(*int64)(unsafe.Add(mBase, uint32(v26)+240)) = v48
	*(*int64)(unsafe.Add(mBase, uint32(v26)+248)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v26)+260)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v26)+264)) = v429
	*(*int32)(unsafe.Add(mBase, uint32(v26)+268)) = v251
	*(*int32)(unsafe.Add(mBase, uint32(v26)+272)) = v247
	*(*int32)(unsafe.Add(mBase, uint32(v26)+276)) = v156
	*(*int64)(unsafe.Add(mBase, uint32(v26)+280)) = v150
	v502 = F_makeDefElem(m, v459, v488, int32(-1))
	mBase = m.M
	v503 = m.ExcPending
	if v503 != 0 {
		goto L6
	} else {
		goto L63
	}
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+228)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v26)+224)) = v36
	*(*int64)(unsafe.Add(mBase, uint32(v26)+232)) = v47
	*(*int64)(unsafe.Add(mBase, uint32(v26)+240)) = v48
	*(*int64)(unsafe.Add(mBase, uint32(v26)+248)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v26)+260)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v26)+264)) = v429
	*(*int32)(unsafe.Add(mBase, uint32(v26)+268)) = v251
	*(*int32)(unsafe.Add(mBase, uint32(v26)+272)) = v247
	*(*int32)(unsafe.Add(mBase, uint32(v26)+276)) = v156
	*(*int64)(unsafe.Add(mBase, uint32(v26)+280)) = v150
	v515 = F_lappend(m, v431, v502)
	mBase = m.M
	v516 = m.ExcPending
	if v516 != 0 {
		goto L6
	} else {
		goto L64
	}
L64:
	;
	v518 = v432 + int32(2)
	v519 = *(*int32)(unsafe.Add(mBase, uint32(v26)+220))
	if v518 < v519 {
		v429 = v515
		v431 = v515
		v432 = v518
		goto L58
	} else {
		goto L65
	}
L65:
	;
	goto L59
L66:
	;
	v558 = *(*int32)(unsafe.Add(mBase, uint32(v54)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v247))) = v558
	v560 = *(*int32)(unsafe.Add(mBase, uint32(v54)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v247)+4)) = v560
	*(*int32)(unsafe.Add(mBase, uint32(v26)+228)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v26)+224)) = v36
	*(*int64)(unsafe.Add(mBase, uint32(v26)+232)) = v47
	*(*int64)(unsafe.Add(mBase, uint32(v26)+240)) = v48
	*(*int64)(unsafe.Add(mBase, uint32(v26)+248)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v26)+260)) = v532
	*(*int32)(unsafe.Add(mBase, uint32(v26)+264)) = v530
	*(*int32)(unsafe.Add(mBase, uint32(v26)+268)) = v251
	*(*int32)(unsafe.Add(mBase, uint32(v26)+272)) = v247
	*(*int32)(unsafe.Add(mBase, uint32(v26)+276)) = v156
	*(*int64)(unsafe.Add(mBase, uint32(v26)+280)) = v150
	v575 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_logical_slot_get_changes_guts[1])))
	if v575 == int32(1) {
		goto L68
	} else {
		goto L69
	}
L67:
	;
	if v585 != 0 {
		goto L51
	} else {
		goto L71
	}
L68:
	;
	v580 = *(*int32)(unsafe.Add(mBase, _c_F_pg_logical_slot_get_changes_guts[2]))
	v581 = *(*int32)(unsafe.Add(mBase, uint32(v580)+308))
	v583 = base.B2i32(v581 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _c_F_pg_logical_slot_get_changes_guts[1])) = uint8(v583)
	v585 = v583
	goto L70
L69:
	;
	v585 = int32(0)
	goto L70
L70:
	;
	goto L67
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+228)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v26)+224)) = v36
	*(*int64)(unsafe.Add(mBase, uint32(v26)+232)) = v47
	*(*int64)(unsafe.Add(mBase, uint32(v26)+240)) = v48
	*(*int64)(unsafe.Add(mBase, uint32(v26)+248)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v26)+260)) = v532
	*(*int32)(unsafe.Add(mBase, uint32(v26)+264)) = v530
	*(*int32)(unsafe.Add(mBase, uint32(v26)+268)) = v251
	*(*int32)(unsafe.Add(mBase, uint32(v26)+272)) = v247
	*(*int32)(unsafe.Add(mBase, uint32(v26)+276)) = v156
	*(*int64)(unsafe.Add(mBase, uint32(v26)+280)) = v150
	v597 = int32(0)
	v599 = int32(_a_F_pg_logical_slot_get_changes_guts_7)
	v600 = *(*int32)(unsafe.Add(mBase, _c_F_pg_logical_slot_get_changes_guts[2]))
	v601 = int64(0)
	v604 = base.AtomicRmwCmpxchg64(m, v600, int32(272), v601, v601)
	*(*int64)(unsafe.Add(mBase, _c_F_pg_logical_slot_get_changes_guts[3])) = v604
	v609 = base.AtomicRmwOr32(m, v597, int32(_a_F_pg_logical_slot_get_changes_guts_8), v597)
	v612 = *(*int32)(unsafe.Add(mBase, _c_F_pg_logical_slot_get_changes_guts[2]))
	v616 = base.AtomicRmwCmpxchg64(m, v612, int32(264), v601, v601)
	*(*int64)(unsafe.Add(mBase, _c_F_pg_logical_slot_get_changes_guts[4])) = v616
	goto L74
L72:
	;
	v698 = v48
	v699 = v623
	v700 = v623
	goto L50
L74:
	;
	goto L75
L75:
	;
	v623 = *(*int64)(unsafe.Add(mBase, _c_F_pg_logical_slot_get_changes_guts[3]))
	goto L72
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+228)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v26)+224)) = v36
	*(*int64)(unsafe.Add(mBase, uint32(v26)+232)) = v47
	*(*int64)(unsafe.Add(mBase, uint32(v26)+240)) = v48
	*(*int64)(unsafe.Add(mBase, uint32(v26)+248)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v26)+260)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v26)+264)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v26)+268)) = v251
	*(*int32)(unsafe.Add(mBase, uint32(v26)+272)) = v247
	*(*int32)(unsafe.Add(mBase, uint32(v26)+276)) = v156
	*(*int64)(unsafe.Add(mBase, uint32(v26)+280)) = v150
	F_errcode(m, int32(1088))
	mBase = m.M
	v652 = m.ExcPending
	if v652 != 0 {
		goto L6
	} else {
		goto L77
	}
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+228)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v26)+224)) = v36
	*(*int64)(unsafe.Add(mBase, uint32(v26)+232)) = v47
	*(*int64)(unsafe.Add(mBase, uint32(v26)+240)) = v48
	*(*int64)(unsafe.Add(mBase, uint32(v26)+248)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v26)+260)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v26)+264)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v26)+268)) = v251
	*(*int32)(unsafe.Add(mBase, uint32(v26)+272)) = v247
	*(*int32)(unsafe.Add(mBase, uint32(v26)+276)) = v156
	*(*int64)(unsafe.Add(mBase, uint32(v26)+280)) = v150
	F_errmsg(m, int32(_a_F_pg_logical_slot_get_changes_guts_9), int32(0))
	mBase = m.M
	v667 = m.ExcPending
	if v667 != 0 {
		goto L6
	} else {
		goto L78
	}
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+228)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v26)+224)) = v36
	*(*int64)(unsafe.Add(mBase, uint32(v26)+232)) = v47
	*(*int64)(unsafe.Add(mBase, uint32(v26)+240)) = v48
	*(*int64)(unsafe.Add(mBase, uint32(v26)+248)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v26)+260)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v26)+264)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v26)+268)) = v251
	*(*int32)(unsafe.Add(mBase, uint32(v26)+272)) = v247
	*(*int32)(unsafe.Add(mBase, uint32(v26)+276)) = v156
	*(*int64)(unsafe.Add(mBase, uint32(v26)+280)) = v150
	F_errfinish(m, int32(_a_F_pg_logical_slot_get_changes_guts_1), int32(177), int32(_a_F_pg_logical_slot_get_changes_guts_2))
	mBase = m.M
	v683 = m.ExcPending
	if v683 != 0 {
		goto L6
	} else {
		goto L79
	}
L79:
	;
	goto L3
L80:
	;
	v698 = v696
	v699 = v49
	v700 = v696
	goto L50
L81:
	;
	v718 = *(*int32)(unsafe.Add(mBase, _c_F_pg_logical_slot_get_changes_guts[5]))
	v720 = *(*int32)(unsafe.Add(mBase, _c_F_pg_logical_slot_get_changes_guts[6]))
	goto L82
L82:
	;
	v722 = v26 + int32(48)
	*(*int32)(unsafe.Add(mBase, uint32(v722)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v722))) = v26 + int32(28)
	goto L85
L83:
	;
	v732 = v247
	v733 = v156
	v734 = v720
	v735 = v718
	v736 = v251
	v737 = v530
	v739 = v532
	v740 = int32(0)
	v745 = v150
	v746 = v700
	v747 = v698
	v748 = v699
	goto L7
L85:
	;
	goto L83
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+44)) = int32(414)
	*(*int32)(unsafe.Add(mBase, uint32(v26)+40)) = int32(415)
	*(*int32)(unsafe.Add(mBase, uint32(v26)+36)) = int32(416)
	*(*int32)(unsafe.Add(mBase, _c_F_pg_logical_slot_get_changes_guts[6])) = v26 + int32(48)
	*(*int32)(unsafe.Add(mBase, uint32(v26)+224)) = v735
	*(*int32)(unsafe.Add(mBase, uint32(v26)+228)) = v734
	*(*int64)(unsafe.Add(mBase, uint32(v26)+232)) = v746
	*(*int64)(unsafe.Add(mBase, uint32(v26)+240)) = v747
	*(*int64)(unsafe.Add(mBase, uint32(v26)+248)) = v748
	*(*int32)(unsafe.Add(mBase, uint32(v26)+260)) = v739
	*(*int32)(unsafe.Add(mBase, uint32(v26)+264)) = v737
	*(*int32)(unsafe.Add(mBase, uint32(v26)+268)) = v736
	*(*int32)(unsafe.Add(mBase, uint32(v26)+272)) = v732
	*(*int32)(unsafe.Add(mBase, uint32(v26)+276)) = v733
	*(*int64)(unsafe.Add(mBase, uint32(v26)+280)) = v745
	v775 = int32(0)
	v781 = F_CreateDecodingContext(m, int64(0), v739, v775, v26+int32(36), int32(1079), int32(1080), v775)
	mBase = m.M
	v782 = m.ExcPending
	if v782 != 0 {
		goto L6
	} else {
		goto L89
	}
L87:
	;
	goto L88
L88:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_pg_logical_slot_get_changes_guts[5])) = v735
	*(*int32)(unsafe.Add(mBase, _c_F_pg_logical_slot_get_changes_guts[6])) = v734
	*(*int32)(unsafe.Add(mBase, uint32(v26)+224)) = v735
	*(*int32)(unsafe.Add(mBase, uint32(v26)+228)) = v734
	*(*int64)(unsafe.Add(mBase, uint32(v26)+232)) = v746
	*(*int64)(unsafe.Add(mBase, uint32(v26)+240)) = v747
	*(*int64)(unsafe.Add(mBase, uint32(v26)+248)) = v748
	*(*int32)(unsafe.Add(mBase, uint32(v26)+260)) = v739
	*(*int32)(unsafe.Add(mBase, uint32(v26)+264)) = v737
	*(*int32)(unsafe.Add(mBase, uint32(v26)+268)) = v736
	*(*int32)(unsafe.Add(mBase, uint32(v26)+272)) = v732
	*(*int32)(unsafe.Add(mBase, uint32(v26)+276)) = v733
	*(*int64)(unsafe.Add(mBase, uint32(v26)+280)) = v745
	F_InvalidateSystemCachesExtended(m)
	mBase = m.M
	v1174 = m.ExcPending
	if v1174 != 0 {
		goto L6
	} else {
		goto L142
	}
L89:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_pg_logical_slot_get_changes_guts[0])) = v736
	if v3 != 0 {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+228)) = v734
	*(*int32)(unsafe.Add(mBase, uint32(v26)+224)) = v735
	*(*int64)(unsafe.Add(mBase, uint32(v26)+240)) = v747
	*(*int64)(unsafe.Add(mBase, uint32(v26)+248)) = v748
	*(*int32)(unsafe.Add(mBase, uint32(v26)+260)) = v739
	*(*int32)(unsafe.Add(mBase, uint32(v26)+264)) = v737
	*(*int32)(unsafe.Add(mBase, uint32(v26)+268)) = v736
	*(*int32)(unsafe.Add(mBase, uint32(v26)+272)) = v732
	*(*int32)(unsafe.Add(mBase, uint32(v26)+276)) = v733
	*(*int64)(unsafe.Add(mBase, uint32(v26)+280)) = v745
	*(*int64)(unsafe.Add(mBase, uint32(v26)+232)) = v746
	if base.Ui64(v745) < base.Ui64(v746) {
		goto L98
	} else {
		goto L99
	}
L91:
	;
	v785 = *(*int32)(unsafe.Add(mBase, uint32(v781)+108))
	if v785 == int32(1) {
		goto L90
	} else {
		goto L92
	}
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+228)) = v734
	*(*int32)(unsafe.Add(mBase, uint32(v26)+224)) = v735
	*(*int64)(unsafe.Add(mBase, uint32(v26)+232)) = v746
	*(*int64)(unsafe.Add(mBase, uint32(v26)+240)) = v747
	*(*int64)(unsafe.Add(mBase, uint32(v26)+248)) = v748
	*(*int32)(unsafe.Add(mBase, uint32(v26)+260)) = v739
	*(*int32)(unsafe.Add(mBase, uint32(v26)+264)) = v737
	*(*int32)(unsafe.Add(mBase, uint32(v26)+268)) = v736
	*(*int32)(unsafe.Add(mBase, uint32(v26)+272)) = v732
	*(*int32)(unsafe.Add(mBase, uint32(v26)+276)) = v733
	*(*int64)(unsafe.Add(mBase, uint32(v26)+280)) = v745
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v802 = m.ExcPending
	if v802 != 0 {
		goto L6
	} else {
		goto L93
	}
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+228)) = v734
	*(*int32)(unsafe.Add(mBase, uint32(v26)+224)) = v735
	*(*int64)(unsafe.Add(mBase, uint32(v26)+232)) = v746
	*(*int64)(unsafe.Add(mBase, uint32(v26)+240)) = v747
	*(*int64)(unsafe.Add(mBase, uint32(v26)+248)) = v748
	*(*int32)(unsafe.Add(mBase, uint32(v26)+260)) = v739
	*(*int32)(unsafe.Add(mBase, uint32(v26)+264)) = v737
	*(*int32)(unsafe.Add(mBase, uint32(v26)+268)) = v736
	*(*int32)(unsafe.Add(mBase, uint32(v26)+272)) = v732
	*(*int32)(unsafe.Add(mBase, uint32(v26)+276)) = v733
	*(*int64)(unsafe.Add(mBase, uint32(v26)+280)) = v745
	F_errcode(m, int32(1088))
	mBase = m.M
	v816 = m.ExcPending
	if v816 != 0 {
		goto L6
	} else {
		goto L94
	}
L94:
	;
	v817 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v818 = *(*int32)(unsafe.Add(mBase, uint32(v817)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+224)) = v735
	*(*int32)(unsafe.Add(mBase, uint32(v26)+228)) = v734
	*(*int64)(unsafe.Add(mBase, uint32(v26)+232)) = v746
	*(*int64)(unsafe.Add(mBase, uint32(v26)+240)) = v747
	*(*int64)(unsafe.Add(mBase, uint32(v26)+248)) = v748
	*(*int32)(unsafe.Add(mBase, uint32(v26)+260)) = v739
	*(*int32)(unsafe.Add(mBase, uint32(v26)+264)) = v737
	*(*int32)(unsafe.Add(mBase, uint32(v26)+268)) = v736
	*(*int32)(unsafe.Add(mBase, uint32(v26)+272)) = v732
	*(*int32)(unsafe.Add(mBase, uint32(v26)+276)) = v733
	*(*int64)(unsafe.Add(mBase, uint32(v26)+280)) = v745
	v831 = *(*int32)(unsafe.Add(mBase, _c_F_pg_logical_slot_get_changes_guts[7]))
	v832 = F_format_procedure(m, v818)
	mBase = m.M
	v833 = m.ExcPending
	if v833 != 0 {
		goto L6
	} else {
		goto L95
	}
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+228)) = v734
	*(*int32)(unsafe.Add(mBase, uint32(v26)+224)) = v735
	*(*int64)(unsafe.Add(mBase, uint32(v26)+232)) = v746
	*(*int64)(unsafe.Add(mBase, uint32(v26)+240)) = v747
	*(*int64)(unsafe.Add(mBase, uint32(v26)+248)) = v748
	*(*int32)(unsafe.Add(mBase, uint32(v26)+260)) = v739
	*(*int32)(unsafe.Add(mBase, uint32(v26)+264)) = v737
	*(*int32)(unsafe.Add(mBase, uint32(v26)+268)) = v736
	*(*int32)(unsafe.Add(mBase, uint32(v26)+272)) = v732
	*(*int32)(unsafe.Add(mBase, uint32(v26)+276)) = v733
	*(*int64)(unsafe.Add(mBase, uint32(v26)+280)) = v745
	*(*int32)(unsafe.Add(mBase, uint32(v26)+20)) = v832
	*(*int32)(unsafe.Add(mBase, uint32(v26)+16)) = v831 + int32(137)
	F_errmsg(m, int32(_a_F_pg_logical_slot_get_changes_guts_10), v26+int32(16))
	mBase = m.M
	v853 = m.ExcPending
	if v853 != 0 {
		goto L6
	} else {
		goto L96
	}
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+228)) = v734
	*(*int32)(unsafe.Add(mBase, uint32(v26)+224)) = v735
	*(*int64)(unsafe.Add(mBase, uint32(v26)+232)) = v746
	*(*int64)(unsafe.Add(mBase, uint32(v26)+240)) = v747
	*(*int64)(unsafe.Add(mBase, uint32(v26)+248)) = v748
	*(*int32)(unsafe.Add(mBase, uint32(v26)+260)) = v739
	*(*int32)(unsafe.Add(mBase, uint32(v26)+264)) = v737
	*(*int32)(unsafe.Add(mBase, uint32(v26)+268)) = v736
	*(*int32)(unsafe.Add(mBase, uint32(v26)+272)) = v732
	*(*int32)(unsafe.Add(mBase, uint32(v26)+276)) = v733
	*(*int64)(unsafe.Add(mBase, uint32(v26)+280)) = v745
	F_errfinish(m, int32(_a_F_pg_logical_slot_get_changes_guts_1), int32(226), int32(_a_F_pg_logical_slot_get_changes_guts_2))
	mBase = m.M
	v869 = m.ExcPending
	if v869 != 0 {
		goto L6
	} else {
		goto L97
	}
L97:
	;
	goto L3
L98:
	;
	v882 = v745
	goto L100
L99:
	;
	v882 = v746
	goto L100
L100:
	;
	v884 = base.B2i32(v745 == int64(0))
	if v745 == int64(0) {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	v885 = v746
	goto L103
L102:
	;
	v885 = v882
	goto L103
L103:
	;
	F_WaitForStandbyConfirmation(m, v885)
	mBase = m.M
	v887 = m.ExcPending
	if v887 != 0 {
		goto L6
	} else {
		goto L104
	}
L104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v781)+140)) = v732
	v889 = *(*int32)(unsafe.Add(mBase, uint32(v781)+8))
	v891 = *(*int32)(unsafe.Add(mBase, _c_F_pg_logical_slot_get_changes_guts[7]))
	v892 = *(*int64)(unsafe.Add(mBase, uint32(v891)+104))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+224)) = v735
	*(*int32)(unsafe.Add(mBase, uint32(v26)+228)) = v734
	*(*int64)(unsafe.Add(mBase, uint32(v26)+232)) = v746
	*(*int64)(unsafe.Add(mBase, uint32(v26)+240)) = v747
	*(*int64)(unsafe.Add(mBase, uint32(v26)+248)) = v748
	*(*int32)(unsafe.Add(mBase, uint32(v26)+260)) = v739
	*(*int32)(unsafe.Add(mBase, uint32(v26)+264)) = v737
	*(*int32)(unsafe.Add(mBase, uint32(v26)+268)) = v736
	*(*int32)(unsafe.Add(mBase, uint32(v26)+272)) = v732
	*(*int32)(unsafe.Add(mBase, uint32(v26)+276)) = v733
	*(*int64)(unsafe.Add(mBase, uint32(v26)+280)) = v745
	F_XLogBeginRead(m, v889, v892)
	mBase = m.M
	v905 = m.ExcPending
	if v905 != 0 {
		goto L6
	} else {
		goto L105
	}
L105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+228)) = v734
	*(*int32)(unsafe.Add(mBase, uint32(v26)+224)) = v735
	*(*int64)(unsafe.Add(mBase, uint32(v26)+232)) = v746
	*(*int64)(unsafe.Add(mBase, uint32(v26)+240)) = v747
	*(*int64)(unsafe.Add(mBase, uint32(v26)+248)) = v748
	*(*int32)(unsafe.Add(mBase, uint32(v26)+260)) = v739
	*(*int32)(unsafe.Add(mBase, uint32(v26)+264)) = v737
	*(*int32)(unsafe.Add(mBase, uint32(v26)+268)) = v736
	*(*int32)(unsafe.Add(mBase, uint32(v26)+272)) = v732
	*(*int32)(unsafe.Add(mBase, uint32(v26)+276)) = v733
	*(*int64)(unsafe.Add(mBase, uint32(v26)+280)) = v745
	F_InvalidateSystemCachesExtended(m)
	mBase = m.M
	v918 = m.ExcPending
	if v918 != 0 {
		goto L6
	} else {
		goto L106
	}
L106:
	;
	v919 = *(*int32)(unsafe.Add(mBase, uint32(v781)+8))
	v920 = *(*int64)(unsafe.Add(mBase, uint32(v919)+40))
	if base.Ui64(v746) <= base.Ui64(v920) {
		v1077 = v920
		goto L107
	} else {
		goto L108
	}
L107:
	;
	v1079 = int32(0)
	if base.B2i32(l1 == v1079)|base.B2i32(v1077 == int64(0)) == v1079 {
		goto L134
	} else {
		goto L135
	}
L108:
	;
	v935 = v919
	goto L109
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+224)) = v735
	*(*int32)(unsafe.Add(mBase, uint32(v26)+32)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v26)+228)) = v734
	*(*int64)(unsafe.Add(mBase, uint32(v26)+232)) = v746
	*(*int64)(unsafe.Add(mBase, uint32(v26)+240)) = v747
	*(*int64)(unsafe.Add(mBase, uint32(v26)+248)) = v748
	*(*int32)(unsafe.Add(mBase, uint32(v26)+260)) = v739
	*(*int32)(unsafe.Add(mBase, uint32(v26)+264)) = v737
	*(*int32)(unsafe.Add(mBase, uint32(v26)+268)) = v736
	*(*int32)(unsafe.Add(mBase, uint32(v26)+272)) = v732
	*(*int32)(unsafe.Add(mBase, uint32(v26)+276)) = v733
	*(*int64)(unsafe.Add(mBase, uint32(v26)+280)) = v745
	v961 = F_XLogReadRecord(m, v935, v26+int32(32))
	mBase = m.M
	v962 = m.ExcPending
	if v962 != 0 {
		goto L6
	} else {
		goto L111
	}
L110:
	;
	v1077 = v1054
	goto L107
L111:
	;
	v963 = *(*int32)(unsafe.Add(mBase, uint32(v26)+32))
	if v963 != 0 {
		goto L112
	} else {
		goto L113
	}
L112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+228)) = v734
	*(*int32)(unsafe.Add(mBase, uint32(v26)+224)) = v735
	*(*int64)(unsafe.Add(mBase, uint32(v26)+232)) = v746
	*(*int64)(unsafe.Add(mBase, uint32(v26)+240)) = v747
	*(*int64)(unsafe.Add(mBase, uint32(v26)+248)) = v748
	*(*int32)(unsafe.Add(mBase, uint32(v26)+260)) = v739
	*(*int32)(unsafe.Add(mBase, uint32(v26)+264)) = v737
	*(*int32)(unsafe.Add(mBase, uint32(v26)+268)) = v736
	*(*int32)(unsafe.Add(mBase, uint32(v26)+272)) = v732
	*(*int32)(unsafe.Add(mBase, uint32(v26)+276)) = v733
	*(*int64)(unsafe.Add(mBase, uint32(v26)+280)) = v745
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v978 = m.ExcPending
	if v978 != 0 {
		goto L6
	} else {
		goto L115
	}
L113:
	;
	goto L114
L114:
	;
	if v961 != 0 {
		goto L118
	} else {
		goto L119
	}
L115:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+224)) = v735
	*(*int32)(unsafe.Add(mBase, uint32(v26)+228)) = v734
	*(*int64)(unsafe.Add(mBase, uint32(v26)+232)) = v746
	*(*int64)(unsafe.Add(mBase, uint32(v26)+240)) = v747
	*(*int64)(unsafe.Add(mBase, uint32(v26)+248)) = v748
	*(*int32)(unsafe.Add(mBase, uint32(v26)+260)) = v739
	*(*int32)(unsafe.Add(mBase, uint32(v26)+264)) = v737
	*(*int32)(unsafe.Add(mBase, uint32(v26)+268)) = v736
	*(*int32)(unsafe.Add(mBase, uint32(v26)+272)) = v732
	*(*int32)(unsafe.Add(mBase, uint32(v26)+276)) = v733
	*(*int64)(unsafe.Add(mBase, uint32(v26)+280)) = v745
	v990 = *(*int32)(unsafe.Add(mBase, uint32(v26)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v26))) = v990
	F_errmsg_internal(m, int32(_a_F_pg_logical_slot_get_changes_guts_11), v26)
	mBase = m.M
	v994 = m.ExcPending
	if v994 != 0 {
		goto L6
	} else {
		goto L116
	}
L116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+228)) = v734
	*(*int32)(unsafe.Add(mBase, uint32(v26)+224)) = v735
	*(*int64)(unsafe.Add(mBase, uint32(v26)+232)) = v746
	*(*int64)(unsafe.Add(mBase, uint32(v26)+240)) = v747
	*(*int64)(unsafe.Add(mBase, uint32(v26)+248)) = v748
	*(*int32)(unsafe.Add(mBase, uint32(v26)+260)) = v739
	*(*int32)(unsafe.Add(mBase, uint32(v26)+264)) = v737
	*(*int32)(unsafe.Add(mBase, uint32(v26)+268)) = v736
	*(*int32)(unsafe.Add(mBase, uint32(v26)+272)) = v732
	*(*int32)(unsafe.Add(mBase, uint32(v26)+276)) = v733
	*(*int64)(unsafe.Add(mBase, uint32(v26)+280)) = v745
	F_errfinish(m, int32(_a_F_pg_logical_slot_get_changes_guts_1), int32(259), int32(_a_F_pg_logical_slot_get_changes_guts_2))
	mBase = m.M
	v1010 = m.ExcPending
	if v1010 != 0 {
		goto L6
	} else {
		goto L117
	}
L117:
	;
	goto L3
L118:
	;
	v1011 = *(*int32)(unsafe.Add(mBase, uint32(v781)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+228)) = v734
	*(*int32)(unsafe.Add(mBase, uint32(v26)+224)) = v735
	*(*int64)(unsafe.Add(mBase, uint32(v26)+232)) = v746
	*(*int64)(unsafe.Add(mBase, uint32(v26)+240)) = v747
	*(*int64)(unsafe.Add(mBase, uint32(v26)+248)) = v748
	*(*int32)(unsafe.Add(mBase, uint32(v26)+260)) = v739
	*(*int32)(unsafe.Add(mBase, uint32(v26)+264)) = v737
	*(*int32)(unsafe.Add(mBase, uint32(v26)+268)) = v736
	*(*int32)(unsafe.Add(mBase, uint32(v26)+272)) = v732
	*(*int32)(unsafe.Add(mBase, uint32(v26)+276)) = v733
	*(*int64)(unsafe.Add(mBase, uint32(v26)+280)) = v745
	F_LogicalDecodingProcessRecord(m, v781, v1011)
	mBase = m.M
	v1024 = m.ExcPending
	if v1024 != 0 {
		goto L6
	} else {
		goto L121
	}
L119:
	;
	goto L120
L120:
	;
	if v884 == int32(0) {
		goto L122
	} else {
		goto L123
	}
L121:
	;
	goto L120
L122:
	;
	v1028 = *(*int32)(unsafe.Add(mBase, uint32(v781)+8))
	v1029 = *(*int64)(unsafe.Add(mBase, uint32(v1028)+40))
	if base.Ui64(v745) <= base.Ui64(v1029) {
		v1077 = v1029
		goto L107
	} else {
		goto L125
	}
L123:
	;
	goto L124
L124:
	;
	if v733 == int32(0) {
		goto L126
	} else {
		goto L127
	}
L125:
	;
	goto L124
L126:
	;
	v1039 = *(*int32)(unsafe.Add(mBase, _c_F_pg_logical_slot_get_changes_guts[8]))
	if v1039 != 0 {
		goto L129
	} else {
		goto L130
	}
L127:
	;
	v1034 = *(*int64)(unsafe.Add(mBase, uint32(v732)+16))
	if v1034 < base.I64_extend_i32_s(v733) {
		goto L126
	} else {
		goto L128
	}
L128:
	;
	v1036 = *(*int32)(unsafe.Add(mBase, uint32(v781)+8))
	v1037 = *(*int64)(unsafe.Add(mBase, uint32(v1036)+40))
	v1077 = v1037
	goto L107
L129:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+228)) = v734
	*(*int32)(unsafe.Add(mBase, uint32(v26)+224)) = v735
	*(*int64)(unsafe.Add(mBase, uint32(v26)+232)) = v746
	*(*int64)(unsafe.Add(mBase, uint32(v26)+240)) = v747
	*(*int64)(unsafe.Add(mBase, uint32(v26)+248)) = v748
	*(*int32)(unsafe.Add(mBase, uint32(v26)+260)) = v739
	*(*int32)(unsafe.Add(mBase, uint32(v26)+264)) = v737
	*(*int32)(unsafe.Add(mBase, uint32(v26)+268)) = v736
	*(*int32)(unsafe.Add(mBase, uint32(v26)+272)) = v732
	*(*int32)(unsafe.Add(mBase, uint32(v26)+276)) = v733
	*(*int64)(unsafe.Add(mBase, uint32(v26)+280)) = v745
	F_ProcessInterrupts(m)
	mBase = m.M
	v1052 = m.ExcPending
	if v1052 != 0 {
		goto L6
	} else {
		goto L132
	}
L130:
	;
	goto L131
L131:
	;
	v1053 = *(*int32)(unsafe.Add(mBase, uint32(v781)+8))
	v1054 = *(*int64)(unsafe.Add(mBase, uint32(v1053)+40))
	if base.Ui64(v1054) < base.Ui64(v746) {
		v935 = v1053
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
	*(*int32)(unsafe.Add(mBase, uint32(v26)+228)) = v734
	*(*int32)(unsafe.Add(mBase, uint32(v26)+224)) = v735
	*(*int64)(unsafe.Add(mBase, uint32(v26)+232)) = v746
	*(*int64)(unsafe.Add(mBase, uint32(v26)+240)) = v747
	*(*int64)(unsafe.Add(mBase, uint32(v26)+248)) = v748
	*(*int32)(unsafe.Add(mBase, uint32(v26)+260)) = v739
	*(*int32)(unsafe.Add(mBase, uint32(v26)+264)) = v737
	*(*int32)(unsafe.Add(mBase, uint32(v26)+268)) = v736
	*(*int32)(unsafe.Add(mBase, uint32(v26)+272)) = v732
	*(*int32)(unsafe.Add(mBase, uint32(v26)+276)) = v733
	*(*int64)(unsafe.Add(mBase, uint32(v26)+280)) = v745
	F_LogicalConfirmReceivedLocation(m, v1077)
	mBase = m.M
	v1098 = m.ExcPending
	if v1098 != 0 {
		goto L6
	} else {
		goto L137
	}
L135:
	;
	goto L136
L136:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+228)) = v734
	*(*int32)(unsafe.Add(mBase, uint32(v26)+224)) = v735
	*(*int64)(unsafe.Add(mBase, uint32(v26)+232)) = v746
	*(*int64)(unsafe.Add(mBase, uint32(v26)+240)) = v747
	*(*int64)(unsafe.Add(mBase, uint32(v26)+248)) = v748
	*(*int32)(unsafe.Add(mBase, uint32(v26)+260)) = v739
	*(*int32)(unsafe.Add(mBase, uint32(v26)+264)) = v737
	*(*int32)(unsafe.Add(mBase, uint32(v26)+268)) = v736
	*(*int32)(unsafe.Add(mBase, uint32(v26)+272)) = v732
	*(*int32)(unsafe.Add(mBase, uint32(v26)+276)) = v733
	*(*int64)(unsafe.Add(mBase, uint32(v26)+280)) = v745
	F_FreeDecodingContext(m, v781)
	mBase = m.M
	v1124 = m.ExcPending
	if v1124 != 0 {
		goto L6
	} else {
		goto L139
	}
L137:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+228)) = v734
	*(*int32)(unsafe.Add(mBase, uint32(v26)+224)) = v735
	*(*int64)(unsafe.Add(mBase, uint32(v26)+232)) = v746
	*(*int64)(unsafe.Add(mBase, uint32(v26)+240)) = v747
	*(*int64)(unsafe.Add(mBase, uint32(v26)+248)) = v748
	*(*int32)(unsafe.Add(mBase, uint32(v26)+260)) = v739
	*(*int32)(unsafe.Add(mBase, uint32(v26)+264)) = v737
	*(*int32)(unsafe.Add(mBase, uint32(v26)+268)) = v736
	*(*int32)(unsafe.Add(mBase, uint32(v26)+272)) = v732
	*(*int32)(unsafe.Add(mBase, uint32(v26)+276)) = v733
	*(*int64)(unsafe.Add(mBase, uint32(v26)+280)) = v745
	F_ReplicationSlotMarkDirty(m)
	mBase = m.M
	v1111 = m.ExcPending
	if v1111 != 0 {
		goto L6
	} else {
		goto L138
	}
L138:
	;
	goto L136
L139:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+228)) = v734
	*(*int32)(unsafe.Add(mBase, uint32(v26)+224)) = v735
	*(*int64)(unsafe.Add(mBase, uint32(v26)+232)) = v746
	*(*int64)(unsafe.Add(mBase, uint32(v26)+240)) = v747
	*(*int64)(unsafe.Add(mBase, uint32(v26)+248)) = v748
	*(*int32)(unsafe.Add(mBase, uint32(v26)+260)) = v739
	*(*int32)(unsafe.Add(mBase, uint32(v26)+264)) = v737
	*(*int32)(unsafe.Add(mBase, uint32(v26)+268)) = v736
	*(*int32)(unsafe.Add(mBase, uint32(v26)+272)) = v732
	*(*int32)(unsafe.Add(mBase, uint32(v26)+276)) = v733
	*(*int64)(unsafe.Add(mBase, uint32(v26)+280)) = v745
	F_ReplicationSlotRelease(m)
	mBase = m.M
	v1137 = m.ExcPending
	if v1137 != 0 {
		goto L6
	} else {
		goto L140
	}
L140:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+228)) = v734
	*(*int32)(unsafe.Add(mBase, uint32(v26)+224)) = v735
	*(*int64)(unsafe.Add(mBase, uint32(v26)+232)) = v746
	*(*int64)(unsafe.Add(mBase, uint32(v26)+240)) = v747
	*(*int64)(unsafe.Add(mBase, uint32(v26)+248)) = v748
	*(*int32)(unsafe.Add(mBase, uint32(v26)+260)) = v739
	*(*int32)(unsafe.Add(mBase, uint32(v26)+264)) = v737
	*(*int32)(unsafe.Add(mBase, uint32(v26)+268)) = v736
	*(*int32)(unsafe.Add(mBase, uint32(v26)+272)) = v732
	*(*int32)(unsafe.Add(mBase, uint32(v26)+276)) = v733
	*(*int64)(unsafe.Add(mBase, uint32(v26)+280)) = v745
	F_InvalidateSystemCachesExtended(m)
	mBase = m.M
	v1150 = m.ExcPending
	if v1150 != 0 {
		goto L6
	} else {
		goto L141
	}
L141:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_pg_logical_slot_get_changes_guts[5])) = v735
	*(*int32)(unsafe.Add(mBase, _c_F_pg_logical_slot_get_changes_guts[6])) = v734
	m.G0 = v26 + int32(288)
	return
L142:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+228)) = v734
	*(*int32)(unsafe.Add(mBase, uint32(v26)+224)) = v735
	*(*int64)(unsafe.Add(mBase, uint32(v26)+232)) = v746
	*(*int64)(unsafe.Add(mBase, uint32(v26)+240)) = v747
	*(*int64)(unsafe.Add(mBase, uint32(v26)+248)) = v748
	*(*int32)(unsafe.Add(mBase, uint32(v26)+260)) = v739
	*(*int32)(unsafe.Add(mBase, uint32(v26)+264)) = v737
	*(*int32)(unsafe.Add(mBase, uint32(v26)+268)) = v736
	*(*int32)(unsafe.Add(mBase, uint32(v26)+272)) = v732
	*(*int32)(unsafe.Add(mBase, uint32(v26)+276)) = v733
	*(*int64)(unsafe.Add(mBase, uint32(v26)+280)) = v745
	F_pg_re_throw(m)
	mBase = m.M
	v1187 = m.ExcPending
	if v1187 != 0 {
		goto L6
	} else {
		goto L143
	}
L143:
	;
	goto L5
L144:
	;
	v1216 = int32(v1212)
	m.G0 = v26
	v1218 = *(*int32)(unsafe.Add(mBase, uint32(v1216)+4))
	v1219 = *(*int32)(unsafe.Add(mBase, uint32(v1216)))
	v1222 = *(*int32)(unsafe.Add(mBase, uint32(v1219)))
	if v26+int32(28) == v1222 {
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
	if v1226 != 0 {
		goto L150
	} else {
		goto L151
	}
L147:
	;
	v1224 = *(*int32)(unsafe.Add(mBase, uint32(v1219)+4))
	v1226 = v1224
	goto L149
L148:
	;
	v1226 = int32(0)
	goto L149
L149:
	;
	goto L146
L150:
	;
	v1227 = *(*int64)(unsafe.Add(mBase, uint32(v26)+280))
	v1228 = *(*int32)(unsafe.Add(mBase, uint32(v26)+276))
	v1229 = *(*int32)(unsafe.Add(mBase, uint32(v26)+272))
	v1230 = *(*int32)(unsafe.Add(mBase, uint32(v26)+268))
	v1231 = *(*int32)(unsafe.Add(mBase, uint32(v26)+264))
	v1232 = *(*int32)(unsafe.Add(mBase, uint32(v26)+260))
	v1233 = *(*int64)(unsafe.Add(mBase, uint32(v26)+248))
	v1234 = *(*int64)(unsafe.Add(mBase, uint32(v26)+240))
	v1235 = *(*int64)(unsafe.Add(mBase, uint32(v26)+232))
	v1236 = *(*int32)(unsafe.Add(mBase, uint32(v26)+228))
	v1237 = *(*int32)(unsafe.Add(mBase, uint32(v26)+224))
	v33 = v1229
	v34 = v1228
	v35 = v1236
	v36 = v1237
	v37 = v1230
	v38 = v1231
	v39 = v1232
	v40 = v1226
	v41 = v1218
	v46 = v1227
	v47 = v1235
	v48 = v1234
	v49 = v1233
	goto L1
L151:
	;
	goto L152
L152:
	;
	F___wasm_longjmp(m, v1219, v1218)
	mBase = m.M
	v1239 = m.ExcPending
	if v1239 != 0 {
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
