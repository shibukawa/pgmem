package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_CreateEmptyBlockRefTable(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	v5 = F_palloc(m, int32(8))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, _c_F_CreateEmptyBlockRefTable[0]))
		*(*int32)(unsafe.Add(mBase, uint32(v5)+4)) = v10
		v13 = F_MemoryContextAllocZero(m, v10, int32(32))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v13)+28)) = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v13)+24)) = v10
			v20 = F_MemoryContextAllocExtended(m, v10, int32(_a_F_CreateEmptyBlockRefTable_0), int32(5))
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return int32(0)
			} else {
				*(*int64)(unsafe.Add(mBase, uint32(v13)+12)) = int64(31662498914303)
				*(*int64)(unsafe.Add(mBase, uint32(v13))) = int64(8192)
				*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v20
				*(*int32)(unsafe.Add(mBase, uint32(v5))) = v13
				return v5
			}
		}
	}
}
func F_crc32_bytea(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
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
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v7 = F_pg_detoast_datum_packed(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int64(0)
	} else {
		v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
		v12 = int32(1)
		v13 = v11 & v12
		if v11 == v12 {
			v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+1)))
			if base.Ui32((v17-int32(1))&int32(255)) < base.Ui32(int32(3)) {
				v43 = int32(4)
				if v13 != 0 {
					v47 = int32(1)
				} else {
					v47 = int32(4)
				}
				v48 = v7 + v47
				v49 = int32(-1)
				if v43 != int32(1) {
					v57 = v48
					v58 = v49
					v59 = int32(0)
					for {
						v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57))))
						v64 = int32(255)
						v66 = int32(2)
						v68 = *(*int32)(unsafe.Add(mBase, uint32((v62^v58)&v64<<(uint(v66)%32))+uint32(_c_F_crc32_bytea[0])))
						v69 = int32(8)
						v71 = v68 ^ int32(base.Ui32(v58)>>(uint(v69)%32))
						v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57)+1)))
						v78 = *(*int32)(unsafe.Add(mBase, uint32((v71^v72)&v64<<(uint(v66)%32))+uint32(_c_F_crc32_bytea[0])))
						v81 = v78 ^ int32(base.Ui32(v71)>>(uint(v69)%32))
						v83 = v57 + v66
						v85 = v59 + v66
						if v85 != v43&int32(-2) {
							v57 = v83
							v58 = v81
							v59 = v85
							continue
						} else {
							break
						}
						break
					}
					if v43&int32(1) == int32(0) {
						v105 = v81
					} else {
						v89 = v83
						v90 = v81
						v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89))))
						v100 = *(*int32)(unsafe.Add(mBase, uint32((v94^v90)&int32(255)<<(uint(int32(2))%32))+uint32(_c_F_crc32_bytea[0])))
						v105 = v100 ^ int32(base.Ui32(v90)>>(uint(int32(8))%32))
					}
				} else {
					v89 = v48
					v90 = v49
					v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89))))
					v100 = *(*int32)(unsafe.Add(mBase, uint32((v94^v90)&int32(255)<<(uint(int32(2))%32))+uint32(_c_F_crc32_bytea[0])))
					v105 = v100 ^ int32(base.Ui32(v90)>>(uint(int32(8))%32))
				}
				return base.I64_extend_i32_u(v105 ^ int32(-1))
			} else {
				if v17 == int32(18) {
					v28 = int32(16)
				} else {
					v28 = int32(0)
				}
				v40 = v28
				if v40 != 0 {
					v43 = v40
					if v13 != 0 {
						v47 = int32(1)
					} else {
						v47 = int32(4)
					}
					v48 = v7 + v47
					v49 = int32(-1)
					if v43 != int32(1) {
						v57 = v48
						v58 = v49
						v59 = int32(0)
						for {
							v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57))))
							v64 = int32(255)
							v66 = int32(2)
							v68 = *(*int32)(unsafe.Add(mBase, uint32((v62^v58)&v64<<(uint(v66)%32))+uint32(_c_F_crc32_bytea[0])))
							v69 = int32(8)
							v71 = v68 ^ int32(base.Ui32(v58)>>(uint(v69)%32))
							v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57)+1)))
							v78 = *(*int32)(unsafe.Add(mBase, uint32((v71^v72)&v64<<(uint(v66)%32))+uint32(_c_F_crc32_bytea[0])))
							v81 = v78 ^ int32(base.Ui32(v71)>>(uint(v69)%32))
							v83 = v57 + v66
							v85 = v59 + v66
							if v85 != v43&int32(-2) {
								v57 = v83
								v58 = v81
								v59 = v85
								continue
							} else {
								break
							}
							break
						}
						if v43&int32(1) == int32(0) {
							v105 = v81
						} else {
							v89 = v83
							v90 = v81
							v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89))))
							v100 = *(*int32)(unsafe.Add(mBase, uint32((v94^v90)&int32(255)<<(uint(int32(2))%32))+uint32(_c_F_crc32_bytea[0])))
							v105 = v100 ^ int32(base.Ui32(v90)>>(uint(int32(8))%32))
						}
					} else {
						v89 = v48
						v90 = v49
						v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89))))
						v100 = *(*int32)(unsafe.Add(mBase, uint32((v94^v90)&int32(255)<<(uint(int32(2))%32))+uint32(_c_F_crc32_bytea[0])))
						v105 = v100 ^ int32(base.Ui32(v90)>>(uint(int32(8))%32))
					}
					return base.I64_extend_i32_u(v105 ^ int32(-1))
				} else {
					return int64(0)
				}
			}
		} else {
			v29 = int32(1)
			if v13 != 0 {
				v40 = int32(base.Ui32(v11)>>(uint(v29)%32)) - v29
			} else {
				v33 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
				v40 = int32(base.Ui32(v33)>>(uint(int32(2))%32)) - int32(4)
			}
			if v40 != 0 {
				v43 = v40
				if v13 != 0 {
					v47 = int32(1)
				} else {
					v47 = int32(4)
				}
				v48 = v7 + v47
				v49 = int32(-1)
				if v43 != int32(1) {
					v57 = v48
					v58 = v49
					v59 = int32(0)
					for {
						v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57))))
						v64 = int32(255)
						v66 = int32(2)
						v68 = *(*int32)(unsafe.Add(mBase, uint32((v62^v58)&v64<<(uint(v66)%32))+uint32(_c_F_crc32_bytea[0])))
						v69 = int32(8)
						v71 = v68 ^ int32(base.Ui32(v58)>>(uint(v69)%32))
						v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57)+1)))
						v78 = *(*int32)(unsafe.Add(mBase, uint32((v71^v72)&v64<<(uint(v66)%32))+uint32(_c_F_crc32_bytea[0])))
						v81 = v78 ^ int32(base.Ui32(v71)>>(uint(v69)%32))
						v83 = v57 + v66
						v85 = v59 + v66
						if v85 != v43&int32(-2) {
							v57 = v83
							v58 = v81
							v59 = v85
							continue
						} else {
							break
						}
						break
					}
					if v43&int32(1) == int32(0) {
						v105 = v81
					} else {
						v89 = v83
						v90 = v81
						v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89))))
						v100 = *(*int32)(unsafe.Add(mBase, uint32((v94^v90)&int32(255)<<(uint(int32(2))%32))+uint32(_c_F_crc32_bytea[0])))
						v105 = v100 ^ int32(base.Ui32(v90)>>(uint(int32(8))%32))
					}
				} else {
					v89 = v48
					v90 = v49
					v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89))))
					v100 = *(*int32)(unsafe.Add(mBase, uint32((v94^v90)&int32(255)<<(uint(int32(2))%32))+uint32(_c_F_crc32_bytea[0])))
					v105 = v100 ^ int32(base.Ui32(v90)>>(uint(int32(8))%32))
				}
				return base.I64_extend_i32_u(v105 ^ int32(-1))
			} else {
				return int64(0)
			}
		}
	}
}
func F_createTrgmNFA(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v134 int32
	_ = v134
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v260 int32
	_ = v260
	var v264 int32
	_ = v264
	var v268 int32
	_ = v268
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v281 int32
	_ = v281
	var v285 int32
	_ = v285
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v339 int32
	_ = v339
	var v377 int32
	_ = v377
	var v381 int32
	_ = v381
	var v386 int32
	_ = v386
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v397 int32
	_ = v397
	var v400 int32
	_ = v400
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v411 int32
	_ = v411
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v435 int32
	_ = v435
	var v439 int32
	_ = v439
	var v441 int32
	_ = v441
	var v447 int32
	_ = v447
	var v489 int32
	_ = v489
	var v530 int32
	_ = v530
	var v571 int64
	_ = v571
	var v578 int32
	_ = v578
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v607 int32
	_ = v607
	var v612 int32
	_ = v612
	var v614 int64
	_ = v614
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v626 int32
	_ = v626
	var v633 int32
	_ = v633
	var v650 int32
	_ = v650
	var v675 int32
	_ = v675
	var v679 int32
	_ = v679
	var v680 int32
	_ = v680
	var v683 int32
	_ = v683
	var v687 int32
	_ = v687
	var v693 int32
	_ = v693
	var v694 int32
	_ = v694
	var v697 int32
	_ = v697
	var v704 int32
	_ = v704
	var v739 int32
	_ = v739
	var v746 int32
	_ = v746
	var v750 int32
	_ = v750
	var v752 int32
	_ = v752
	var v754 int32
	_ = v754
	var v755 int32
	_ = v755
	var v758 int32
	_ = v758
	var v763 int32
	_ = v763
	var v799 int32
	_ = v799
	var v802 int32
	_ = v802
	var v805 int32
	_ = v805
	var v809 int32
	_ = v809
	var v812 int32
	_ = v812
	var v813 int32
	_ = v813
	var v828 int32
	_ = v828
	var v856 int32
	_ = v856
	var v857 int32
	_ = v857
	var v861 int32
	_ = v861
	var v862 int32
	_ = v862
	var v863 int32
	_ = v863
	var v864 int32
	_ = v864
	var v865 int32
	_ = v865
	var v866 int32
	_ = v866
	var v867 int32
	_ = v867
	var v869 int32
	_ = v869
	var v877 int32
	_ = v877
	var v914 int32
	_ = v914
	var v915 int32
	_ = v915
	var v918 int32
	_ = v918
	var v921 int32
	_ = v921
	var v922 int32
	_ = v922
	var v925 int32
	_ = v925
	var v928 int32
	_ = v928
	var v929 int32
	_ = v929
	var v932 int32
	_ = v932
	var v940 int32
	_ = v940
	var v942 int32
	_ = v942
	var v945 int32
	_ = v945
	var v947 int32
	_ = v947
	var v949 int32
	_ = v949
	var v956 int32
	_ = v956
	var v960 int32
	_ = v960
	var v1002 int32
	_ = v1002
	var v1004 int32
	_ = v1004
	var v1005 int32
	_ = v1005
	var v1046 int32
	_ = v1046
	var v1049 int32
	_ = v1049
	var v1051 int32
	_ = v1051
	var v1052 int64
	_ = v1052
	var v1053 int64
	_ = v1053
	var v1056 int64
	_ = v1056
	var v1057 int64
	_ = v1057
	var v1058 int64
	_ = v1058
	var v1059 int64
	_ = v1059
	var v1060 int64
	_ = v1060
	var v1061 int64
	_ = v1061
	var v1062 int64
	_ = v1062
	var v1063 int64
	_ = v1063
	var v1064 int64
	_ = v1064
	var v1065 int64
	_ = v1065
	var v1066 int64
	_ = v1066
	var v1067 int64
	_ = v1067
	var v1068 int64
	_ = v1068
	var v1069 int64
	_ = v1069
	var v1070 int64
	_ = v1070
	var v1071 int64
	_ = v1071
	var v1072 int64
	_ = v1072
	var v1073 int64
	_ = v1073
	var v1074 int64
	_ = v1074
	var v1075 int64
	_ = v1075
	var v1076 int64
	_ = v1076
	var v1077 int64
	_ = v1077
	var v1078 int64
	_ = v1078
	var v1079 int64
	_ = v1079
	var v1080 int64
	_ = v1080
	var v1081 int64
	_ = v1081
	var v1082 int64
	_ = v1082
	var v1083 int64
	_ = v1083
	var v1084 int64
	_ = v1084
	var v1085 int64
	_ = v1085
	var v1086 int64
	_ = v1086
	var v1118 int64
	_ = v1118
	var v1121 int32
	_ = v1121
	var v1124 int32
	_ = v1124
	var v1125 int32
	_ = v1125
	var v1127 int32
	_ = v1127
	var v1132 int32
	_ = v1132
	var v1139 int32
	_ = v1139
	var v1167 int32
	_ = v1167
	var v1171 int32
	_ = v1171
	var v1172 int32
	_ = v1172
	var v1175 int32
	_ = v1175
	var v1176 int32
	_ = v1176
	var v1178 int32
	_ = v1178
	var v1179 int32
	_ = v1179
	var v1180 int32
	_ = v1180
	var v1183 int32
	_ = v1183
	var v1190 int32
	_ = v1190
	var v1221 int32
	_ = v1221
	var v1224 int32
	_ = v1224
	var v1225 int32
	_ = v1225
	var v1229 int32
	_ = v1229
	var v1235 int32
	_ = v1235
	var v1267 int32
	_ = v1267
	var v1271 int32
	_ = v1271
	var v1273 int32
	_ = v1273
	var v1274 int32
	_ = v1274
	var v1276 int32
	_ = v1276
	var v1280 int32
	_ = v1280
	var v1281 int32
	_ = v1281
	var v1283 int64
	_ = v1283
	var v1285 int32
	_ = v1285
	var v1294 int32
	_ = v1294
	var v1295 int32
	_ = v1295
	var v1297 int32
	_ = v1297
	var v1298 int32
	_ = v1298
	var v1300 int32
	_ = v1300
	var v1301 int32
	_ = v1301
	var v1304 int32
	_ = v1304
	var v1344 int32
	_ = v1344
	var v1345 int32
	_ = v1345
	var v1390 int32
	_ = v1390
	var v1396 int32
	_ = v1396
	var __phi1396 int32
	_ = __phi1396
	var v1400 int32
	_ = v1400
	var __phi1400 int32
	_ = __phi1400
	var v1402 int32
	_ = v1402
	var __phi1402 int32
	_ = __phi1402
	var v1435 int64
	_ = v1435
	var v1436 int64
	_ = v1436
	var v1438 int64
	_ = v1438
	var v1440 int64
	_ = v1440
	var v1443 int64
	_ = v1443
	var v1445 int64
	_ = v1445
	var v1447 int64
	_ = v1447
	var v1449 int64
	_ = v1449
	var v1470 int64
	_ = v1470
	var v1471 int64
	_ = v1471
	var v1506 int64
	_ = v1506
	var v1509 int32
	_ = v1509
	var v1510 int32
	_ = v1510
	var v1512 int32
	_ = v1512
	var v1514 int32
	_ = v1514
	var v1519 int64
	_ = v1519
	var v1520 int32
	_ = v1520
	var v1530 int64
	_ = v1530
	var v1533 int64
	_ = v1533
	var v1534 int64
	_ = v1534
	var v1538 int32
	_ = v1538
	var v1542 int32
	_ = v1542
	var v1545 int64
	_ = v1545
	var v1547 int64
	_ = v1547
	var v1549 int64
	_ = v1549
	var v1551 int64
	_ = v1551
	var v1555 int32
	_ = v1555
	var v1556 int32
	_ = v1556
	var v1557 int32
	_ = v1557
	var v1558 int32
	_ = v1558
	var v1560 int32
	_ = v1560
	var v1562 int32
	_ = v1562
	var v1580 int32
	_ = v1580
	var v1613 int32
	_ = v1613
	var v1621 int32
	_ = v1621
	var v1649 int64
	_ = v1649
	var v1651 float32
	_ = v1651
	var v1657 int32
	_ = v1657
	var v1658 int32
	_ = v1658
	var v1665 int32
	_ = v1665
	var v1666 int32
	_ = v1666
	var v1667 int32
	_ = v1667
	var v1668 int32
	_ = v1668
	var v1674 int32
	_ = v1674
	var v1678 int32
	_ = v1678
	var v1679 int32
	_ = v1679
	var v1681 int32
	_ = v1681
	var v1682 int32
	_ = v1682
	var v1688 int32
	_ = v1688
	var v1692 int32
	_ = v1692
	var v1693 int32
	_ = v1693
	var v1699 float32
	_ = v1699
	var v1701 float32
	_ = v1701
	var v1703 float32
	_ = v1703
	var v1705 int64
	_ = v1705
	var v1707 int32
	_ = v1707
	var v1712 int32
	_ = v1712
	var v1730 int32
	_ = v1730
	var v1748 int64
	_ = v1748
	var v1750 float32
	_ = v1750
	var v1758 int32
	_ = v1758
	var v1759 int32
	_ = v1759
	var v1762 int32
	_ = v1762
	var v1765 int32
	_ = v1765
	var v1775 int32
	_ = v1775
	var v1779 int32
	_ = v1779
	var v1809 int32
	_ = v1809
	var v1810 int32
	_ = v1810
	var v1811 int32
	_ = v1811
	var v1812 int32
	_ = v1812
	var v1851 int32
	_ = v1851
	var v1858 int32
	_ = v1858
	var v1891 int32
	_ = v1891
	var v1892 int32
	_ = v1892
	var v1893 int32
	_ = v1893
	var v1899 int32
	_ = v1899
	var v1902 int32
	_ = v1902
	var v1904 int32
	_ = v1904
	var v1934 int32
	_ = v1934
	var v1935 int32
	_ = v1935
	var v1937 int32
	_ = v1937
	var v1939 int32
	_ = v1939
	var v1940 int32
	_ = v1940
	var v1942 int32
	_ = v1942
	var v1948 int32
	_ = v1948
	var v1981 int32
	_ = v1981
	var v1982 int32
	_ = v1982
	var v1984 int32
	_ = v1984
	var v1987 int32
	_ = v1987
	var v1990 int32
	_ = v1990
	var v1997 int32
	_ = v1997
	var v1998 int32
	_ = v1998
	var v2000 int32
	_ = v2000
	var v2002 int32
	_ = v2002
	var v2004 int32
	_ = v2004
	var v2014 int32
	_ = v2014
	var v2049 int32
	_ = v2049
	var v2050 int32
	_ = v2050
	var v2051 int32
	_ = v2051
	var v2052 int32
	_ = v2052
	var v2091 int32
	_ = v2091
	var v2098 int32
	_ = v2098
	var v2131 int32
	_ = v2131
	var v2136 int32
	_ = v2136
	var v2173 int32
	_ = v2173
	var v2174 int32
	_ = v2174
	var v2213 int32
	_ = v2213
	var v2217 int32
	_ = v2217
	var v2218 int32
	_ = v2218
	var v2228 int32
	_ = v2228
	var v2259 int32
	_ = v2259
	var v2269 int32
	_ = v2269
	var v2270 int32
	_ = v2270
	var v2304 int32
	_ = v2304
	var v2305 int32
	_ = v2305
	var v2306 int32
	_ = v2306
	var v2307 int32
	_ = v2307
	var v2346 int32
	_ = v2346
	var v2351 int32
	_ = v2351
	var v2386 int32
	_ = v2386
	var v2388 int32
	_ = v2388
	var v2389 int32
	_ = v2389
	var v2393 int32
	_ = v2393
	var v2394 int32
	_ = v2394
	var v2396 int32
	_ = v2396
	var v2437 int32
	_ = v2437
	var v2439 float32
	_ = v2439
	var v2441 int64
	_ = v2441
	var v2478 int64
	_ = v2478
	var v2480 float32
	_ = v2480
	var v2483 int32
	_ = v2483
	var v2520 int64
	_ = v2520
	var v2526 int32
	_ = v2526
	var v2531 int32
	_ = v2531
	var v2533 int32
	_ = v2533
	var v2534 int32
	_ = v2534
	var v2542 int32
	_ = v2542
	var v2543 int32
	_ = v2543
	var v2546 int32
	_ = v2546
	var v2583 int32
	_ = v2583
	var v2584 int32
	_ = v2584
	var v2590 int32
	_ = v2590
	var v2591 int32
	_ = v2591
	var v2597 int32
	_ = v2597
	var v2598 int32
	_ = v2598
	var v2604 int32
	_ = v2604
	var v2605 int32
	_ = v2605
	var v2611 int32
	_ = v2611
	var v2612 int32
	_ = v2612
	var v2613 int32
	_ = v2613
	var v2615 int32
	_ = v2615
	var v2619 int32
	_ = v2619
	var v2623 int32
	_ = v2623
	var v2658 int32
	_ = v2658
	var v2662 int32
	_ = v2662
	var v2665 int32
	_ = v2665
	var v2699 int32
	_ = v2699
	var v2700 int32
	_ = v2700
	var v2706 int32
	_ = v2706
	var v2707 int32
	_ = v2707
	var v2710 int32
	_ = v2710
	var v2715 int32
	_ = v2715
	var v2716 int32
	_ = v2716
	var v2722 int32
	_ = v2722
	var v2724 int32
	_ = v2724
	var v2726 int32
	_ = v2726
	var v2730 int32
	_ = v2730
	var v2733 int32
	_ = v2733
	var v2739 int32
	_ = v2739
	var v2744 int32
	_ = v2744
	var v2757 int32
	_ = v2757
	var v2790 int32
	_ = v2790
	var v2791 int32
	_ = v2791
	var v2797 int32
	_ = v2797
	var v2810 int32
	_ = v2810
	var v2818 int32
	_ = v2818
	var v2822 int32
	_ = v2822
	var v2832 int32
	_ = v2832
	var v2850 int32
	_ = v2850
	var v2853 int32
	_ = v2853
	var v2854 int32
	_ = v2854
	var v2857 int32
	_ = v2857
	var v2858 int32
	_ = v2858
	var v2861 int32
	_ = v2861
	var v2862 int32
	_ = v2862
	var v2865 int32
	_ = v2865
	var v2868 int32
	_ = v2868
	var v2871 int32
	_ = v2871
	var v2872 int32
	_ = v2872
	var v2873 int32
	_ = v2873
	var v2875 int32
	_ = v2875
	var v2876 int32
	_ = v2876
	var v2880 int32
	_ = v2880
	var v2883 int32
	_ = v2883
	var v2884 int32
	_ = v2884
	var v2885 int32
	_ = v2885
	var v2887 int32
	_ = v2887
	var v2891 int32
	_ = v2891
	var v2894 int32
	_ = v2894
	var v2900 int32
	_ = v2900
	var v2930 int32
	_ = v2930
	var v2931 int32
	_ = v2931
	var v2935 int32
	_ = v2935
	var v2938 int32
	_ = v2938
	var v2941 int32
	_ = v2941
	var v2942 int32
	_ = v2942
	var v2944 int32
	_ = v2944
	var v2946 int32
	_ = v2946
	var v2948 int32
	_ = v2948
	var v2952 int32
	_ = v2952
	var v2956 int32
	_ = v2956
	var v2959 int32
	_ = v2959
	var v2966 int32
	_ = v2966
	var v2995 int32
	_ = v2995
	var v2996 int32
	_ = v2996
	var v3000 int32
	_ = v3000
	var v3002 int32
	_ = v3002
	var v3004 int32
	_ = v3004
	var v3006 int32
	_ = v3006
	var v3009 int32
	_ = v3009
	var v3015 int32
	_ = v3015
	var v3049 int32
	_ = v3049
	var v3050 int32
	_ = v3050
	var v3054 int32
	_ = v3054
	var v3069 int32
	_ = v3069
	var v3071 int32
	_ = v3071
	var v3073 int32
	_ = v3073
	var v3096 int32
	_ = v3096
	var v3098 int32
	_ = v3098
	var v3100 int32
	_ = v3100
	var v3105 int32
	_ = v3105
	var v3112 int32
	_ = v3112
	var v3121 int32
	_ = v3121
	var v3129 int32
	_ = v3129
	var v3131 int32
	_ = v3131
	var v3134 int32
	_ = v3134
	var v3135 int32
	_ = v3135
	var v3141 int32
	_ = v3141
	var v3145 int32
	_ = v3145
	var v3147 int32
	_ = v3147
	var v3149 int32
	_ = v3149
	var v3152 int32
	_ = v3152
	var v3153 int32
	_ = v3153
	var v3155 int32
	_ = v3155
	var v3156 int32
	_ = v3156
	var v3157 int32
	_ = v3157
	var v3158 int32
	_ = v3158
	var v3159 int32
	_ = v3159
	var v3161 int32
	_ = v3161
	var v3167 int32
	_ = v3167
	var v3170 int32
	_ = v3170
	var v3175 int32
	_ = v3175
	var v3178 int32
	_ = v3178
	var v3185 int32
	_ = v3185
	var v3188 int32
	_ = v3188
	var v3193 int32
	_ = v3193
	var v3201 int32
	_ = v3201
	var v3212 int32
	_ = v3212
	var v3216 int32
	_ = v3216
	var v3217 int32
	_ = v3217
	var v3219 int32
	_ = v3219
	var v3221 int32
	_ = v3221
	var v3222 int32
	_ = v3222
	var v3226 int32
	_ = v3226
	var v3229 int32
	_ = v3229
	var v3262 int32
	_ = v3262
	var v3268 int32
	_ = v3268
	var v3271 int32
	_ = v3271
	var v3303 int32
	_ = v3303
	var v3304 int32
	_ = v3304
	var v3308 int32
	_ = v3308
	var v3311 int32
	_ = v3311
	var v3344 int32
	_ = v3344
	var v3353 int32
	_ = v3353
	var v3385 int32
	_ = v3385
	var v3393 int32
	_ = v3393
	var v3397 int32
	_ = v3397
	var v3426 int32
	_ = v3426
	var v3428 int32
	_ = v3428
	var v3443 int32
	_ = v3443
	var v3469 int32
	_ = v3469
	var v3471 int32
	_ = v3471
	var v3472 int32
	_ = v3472
	var v3473 int32
	_ = v3473
	var v3474 int32
	_ = v3474
	var v3479 int32
	_ = v3479
	var v3486 int32
	_ = v3486
	var v3514 int32
	_ = v3514
	var v3515 int32
	_ = v3515
	var v3518 int32
	_ = v3518
	var v3531 int32
	_ = v3531
	var v3534 int32
	_ = v3534
	var v3535 int32
	_ = v3535
	var v3547 int32
	_ = v3547
	var v3576 int32
	_ = v3576
	var v3577 int32
	_ = v3577
	var v3578 int32
	_ = v3578
	var v3580 int32
	_ = v3580
	var v3582 int32
	_ = v3582
	var v3583 int32
	_ = v3583
	var v3584 int32
	_ = v3584
	var v3586 int32
	_ = v3586
	var v3587 int32
	_ = v3587
	var v3588 int32
	_ = v3588
	var v3589 int32
	_ = v3589
	var v3631 int32
	_ = v3631
	var v3666 int32
	_ = v3666
	var v3667 int32
	_ = v3667
	var v3670 int32
	_ = v3670
	var v3671 int32
	_ = v3671
	var v3675 int32
	_ = v3675
	var v3681 int32
	_ = v3681
	var v3683 int32
	_ = v3683
	var v3713 int32
	_ = v3713
	var v3717 int32
	_ = v3717
	var v3718 int32
	_ = v3718
	var v3723 int32
	_ = v3723
	var v3758 int32
	_ = v3758
	var v3759 int32
	_ = v3759
	var v3760 int32
	_ = v3760
	var v3764 int32
	_ = v3764
	var v3765 int32
	_ = v3765
	var v3768 int32
	_ = v3768
	var v3769 int32
	_ = v3769
	var v3771 int32
	_ = v3771
	var v3773 int32
	_ = v3773
	var v3775 int32
	_ = v3775
	var v3778 int32
	_ = v3778
	var v3780 int32
	_ = v3780
	var v3782 int32
	_ = v3782
	var v3785 int32
	_ = v3785
	var v3825 int32
	_ = v3825
	var v3826 int32
	_ = v3826
	var v3830 int32
	_ = v3830
	var v3833 int32
	_ = v3833
	var v3838 int32
	_ = v3838
	var __phi3838 int32
	_ = __phi3838
	var v3842 int32
	_ = v3842
	var __phi3842 int32
	_ = __phi3842
	var v3845 int32
	_ = v3845
	var __phi3845 int32
	_ = __phi3845
	var v3877 int32
	_ = v3877
	var v3878 int32
	_ = v3878
	var v3881 int32
	_ = v3881
	var v3882 int32
	_ = v3882
	var v3885 int32
	_ = v3885
	var v3886 int32
	_ = v3886
	var v3890 int32
	_ = v3890
	var v3892 int64
	_ = v3892
	var v3896 int32
	_ = v3896
	var v3900 int32
	_ = v3900
	var v3904 int32
	_ = v3904
	var v3907 int32
	_ = v3907
	var v3912 int32
	_ = v3912
	var v3914 int32
	_ = v3914
	var v3952 int32
	_ = v3952
	var v3954 int32
	_ = v3954
	var v3955 int32
	_ = v3955
	var v3956 int32
	_ = v3956
	var v3958 int32
	_ = v3958
	var v3962 int32
	_ = v3962
	var v3963 int32
	_ = v3963
	var v3966 int32
	_ = v3966
	var v3967 int32
	_ = v3967
	var v3968 int32
	_ = v3968
	var v3974 int32
	_ = v3974
	var v3978 int32
	_ = v3978
	var v3983 int32
	_ = v3983
	var v4015 int32
	_ = v4015
	var v4016 int32
	_ = v4016
	var v4020 int32
	_ = v4020
	var v4022 int32
	_ = v4022
	var v4023 int32
	_ = v4023
	var v4027 int32
	_ = v4027
	var v4029 int32
	_ = v4029
	var v4030 int32
	_ = v4030
	var v4034 int32
	_ = v4034
	var v4036 int32
	_ = v4036
	var v4037 int32
	_ = v4037
	var v4041 int32
	_ = v4041
	var v4043 int32
	_ = v4043
	var v4044 int32
	_ = v4044
	var v4045 int32
	_ = v4045
	var v4047 int32
	_ = v4047
	var v4051 int32
	_ = v4051
	var v4055 int32
	_ = v4055
	var v4091 int32
	_ = v4091
	var v4095 int32
	_ = v4095
	var v4097 int32
	_ = v4097
	var v4133 int32
	_ = v4133
	var v4137 int32
	_ = v4137
	var v4139 int32
	_ = v4139
	var v4140 int32
	_ = v4140
	var v4143 int32
	_ = v4143
	var v4145 int32
	_ = v4145
	var v4186 int32
	_ = v4186
	var v4187 int32
	_ = v4187
	var v4189 int32
	_ = v4189
	var v4198 int32
	_ = v4198
	var v4202 int32
	_ = v4202
	var v4207 int32
	_ = v4207
	var v4239 int32
	_ = v4239
	var v4240 int32
	_ = v4240
	var v4246 int32
	_ = v4246
	var v4250 int32
	_ = v4250
	var v4251 int32
	_ = v4251
	var v4257 int32
	_ = v4257
	var v4261 int32
	_ = v4261
	var v4262 int32
	_ = v4262
	var v4263 int32
	_ = v4263
	var v4265 int32
	_ = v4265
	var v4269 int32
	_ = v4269
	var v4273 int32
	_ = v4273
	var v4310 int32
	_ = v4310
	var v4311 int32
	_ = v4311
	var v4317 int32
	_ = v4317
	var v4361 int32
	_ = v4361
	var v4362 int32
	_ = v4362
	var v4366 int32
	_ = v4366
	var v4367 int32
	_ = v4367
	var v4370 int32
	_ = v4370
	var v4371 int32
	_ = v4371
	var v4377 int32
	_ = v4377
	var v4382 int32
	_ = v4382
	var v4412 int32
	_ = v4412
	var v4414 int32
	_ = v4414
	var v4421 int32
	_ = v4421
	var v4427 int32
	_ = v4427
	var v4429 int32
	_ = v4429
	var v4464 int32
	_ = v4464
	var v4465 int32
	_ = v4465
	var v4469 int32
	_ = v4469
	var v4470 int32
	_ = v4470
	var v4472 int32
	_ = v4472
	var v4474 int32
	_ = v4474
	var v4477 int32
	_ = v4477
	var v4483 int32
	_ = v4483
	var v4485 int32
	_ = v4485
	var v4520 int32
	_ = v4520
	var v4561 int32
	_ = v4561
	var v4562 int32
	_ = v4562
	var v4563 int32
	_ = v4563
	var v4565 int32
	_ = v4565
	var v4566 int32
	_ = v4566
	var v4567 int32
	_ = v4567
	var v4569 int32
	_ = v4569
	var v4572 int32
	_ = v4572
	var v4573 int32
	_ = v4573
	var v4595 int32
	_ = v4595
	var v4618 int32
	_ = v4618
	v5 = int32(0)
	v40 = m.G0
	v42 = v40 - int32(240)
	m.G0 = v42
	v45 = *(*int32)(unsafe.Add(mBase, _c_F_createTrgmNFA[0]))
	v50 = F_AllocSetContextCreateInternal(m, v45, int32(_a_F_createTrgmNFA_0), v5, int32(_a_F_createTrgmNFA_1), int32(_a_F_createTrgmNFA_2))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v54 = int32(_a_F_createTrgmNFA_3)
	v55 = *(*int32)(unsafe.Add(mBase, _c_F_createTrgmNFA[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_createTrgmNFA[0])) = v50
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v60 == int32(1) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v94 = F_palloc(m, v89<<(uint(int32(2))%32)+int32(4))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L1
	} else {
		goto L14
	}
L4:
	;
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
	if v66 == int32(18) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	v77 = int32(1)
	if v60&v77 != 0 {
		v89 = int32(base.Ui32(v60)>>(uint(v77)%32)) - v77
		goto L3
	} else {
		goto L13
	}
L7:
	;
	v69 = int32(16)
	goto L9
L8:
	;
	v69 = int32(0)
	goto L9
L9:
	;
	if base.Ui32((v66-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v76 = int32(4)
	goto L12
L11:
	;
	v76 = v69
	goto L12
L12:
	;
	v89 = v76
	goto L3
L13:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v89 = int32(base.Ui32(v83)>>(uint(int32(2))%32)) - int32(4)
	goto L3
L14:
	;
	v96 = int32(1)
	if v60&v96 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v100 = v96
	goto L17
L16:
	;
	v100 = int32(4)
	goto L17
L17:
	;
	v102 = F_pg_mb2wchar_with_len(m, l0+v100, v94, v89)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v105 = F_pg_regcomp(m, v42+int32(32), v94, v102, int32(27), l1)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	F_pfree(m, v94)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	if v105 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_createTrgmNFA[0])) = v55
	F_MemoryContextDelete(m, v50)
	mBase = m.M
	v4618 = m.ExcPending
	if v4618 != 0 {
		goto L1
	} else {
		goto L521
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v42)+224)) = int32(0)
	v2790 = F_MemoryContextAllocZero(m, l3, v2757*int32(3)+int32(5))
	mBase = m.M
	v2791 = m.ExcPending
	if v2791 != 0 {
		goto L1
	} else {
		goto L310
	}
L23:
	;
	v112 = v42 + int32(32)
	*(*int32)(unsafe.Add(mBase, uint32(v42)+64)) = v112
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v112)+24))
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v114)+84))
	v117 = v115 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v42)+72)) = v117
	v121 = F_palloc0(m, v117*int32(12))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L1
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	v2724 = v42 - int32(-64)
	F_pg_regerror(m, v105, v2724)
	mBase = m.M
	v2726 = m.ExcPending
	if v2726 != 0 {
		goto L1
	} else {
		goto L305
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v42)+68)) = v121
	if int32(0) < v117 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v134 = v5
	goto L30
L28:
	;
	goto L29
L29:
	;
	v571 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v42)+88)) = v571
	*(*int64)(unsafe.Add(mBase, uint32(v42)+93)) = v571
	*(*int64)(unsafe.Add(mBase, uint32(v42)+184)) = int64(171798691852)
	v578 = *(*int32)(unsafe.Add(mBase, _c_F_createTrgmNFA[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v42)+212)) = v578
	v585 = F_hash_create(m, int32(_a_F_createTrgmNFA_4), int64(1024), v42+int32(176), int32(1064))
	mBase = m.M
	v586 = m.ExcPending
	if v586 != 0 {
		goto L1
	} else {
		goto L84
	}
L30:
	;
	v167 = v121 + v134*int32(12)
	v170 = int32(-1)
	if v134 <= int32(0) {
		v186 = v170
		goto L32
	} else {
		goto L33
	}
L31:
	;
	goto L29
L32:
	;
	if base.Ui32(int32(257)) <= base.Ui32(v186) {
		goto L38
	} else {
		goto L39
	}
L33:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v42+int32(32))+24))
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v173)+84))
	if base.Ui32(v174) < base.Ui32(v134) {
		v186 = v170
		goto L32
	} else {
		goto L34
	}
L34:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v173)+92))
	v179 = v176 + v134*int32(24)
	v180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v179)+20)))
	if v180&int32(2) != 0 {
		v186 = v170
		goto L32
	} else {
		goto L35
	}
L35:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v179)+4))
	if v183 != 0 {
		v186 = v170
		goto L32
	} else {
		goto L36
	}
L36:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v179)))
	v186 = v184
	goto L32
L37:
	;
	v530 = v134 + int32(1)
	if v530 != v117 {
		v134 = v530
		goto L30
	} else {
		goto L83
	}
L38:
	;
	v189 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v167))) = uint8(v189)
	goto L37
L39:
	;
	goto L40
L40:
	;
	v191 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v167))) = uint16(v191)
	v195 = F_palloc_mul(m, int32(4), v186)
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v167)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v167)+8)) = v195
	v203 = F_palloc_mul(m, int32(4), v186)
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	v205 = int32(0)
	if base.B2i32(v134 <= v205)|base.B2i32(v186 <= v205) != 0 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	if v186 != 0 {
		goto L58
	} else {
		goto L59
	}
L44:
	;
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v42+int32(32))+24))
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v210)+84))
	if base.Ui32(v211) < base.Ui32(v134) {
		goto L43
	} else {
		goto L45
	}
L45:
	;
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v210)+92))
	v217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v213+v134*int32(24))+20)))
	if v217&int32(2) != 0 {
		goto L43
	} else {
		goto L46
	}
L46:
	;
	v221 = v186
	v228 = int32(0)
	v230 = v203
	goto L47
L47:
	;
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v210)+96))
	v264 = int32(*(*int16)(unsafe.Add(mBase, uint32(v260+v228<<(uint(int32(1))%32)))))
	if v264 == v134 {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	goto L43
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v230))) = v228
	v268 = v221 - int32(1)
	if v268 == int32(0) {
		goto L43
	} else {
		goto L52
	}
L50:
	;
	v273 = v221
	v274 = v230
	goto L51
L51:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v210)+96))
	v276 = int32(1)
	v277 = v228 | v276
	v281 = int32(*(*int16)(unsafe.Add(mBase, uint32(v275+v277<<(uint(v276)%32)))))
	if v281 == v134 {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	v273 = v268
	v274 = v230 + int32(4)
	goto L51
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v274))) = v277
	v285 = v273 - int32(1)
	if v285 == int32(0) {
		goto L43
	} else {
		goto L56
	}
L54:
	;
	v290 = v273
	v291 = v274
	goto L55
L55:
	;
	v293 = v228 + int32(2)
	if v293 != int32(2048) {
		v221 = v290
		v228 = v293
		v230 = v291
		goto L47
	} else {
		goto L57
	}
L56:
	;
	v290 = v285
	v291 = v274 + int32(4)
	goto L55
L57:
	;
	goto L48
L58:
	;
	v339 = int32(0)
	goto L61
L59:
	;
	goto L60
L60:
	;
	F_pfree(m, v203)
	mBase = m.M
	v489 = m.ExcPending
	if v489 != 0 {
		goto L1
	} else {
		goto L82
	}
L61:
	;
	v377 = *(*int32)(unsafe.Add(mBase, uint32(v203+v339<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v42)+224)) = v377
	if v377 == int32(0) {
		goto L63
	} else {
		goto L64
	}
L62:
	;
	goto L60
L63:
	;
	v447 = v339 + int32(1)
	if v447 != v186 {
		v339 = v447
		goto L61
	} else {
		goto L81
	}
L64:
	;
	v381 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v42)+180)) = uint8(v381)
	*(*int32)(unsafe.Add(mBase, uint32(v42)+176)) = v381
	v386 = v42 + int32(176)
	v390 = F_pg_wchar2mb_with_len(m, v42+int32(224), v386, int32(1))
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	v393 = F_str_tolower(m, v386, v390, int32(100))
	mBase = m.M
	v394 = m.ExcPending
	if v394 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	v397 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v393))))
	v400 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v386))))
	if base.B2i32(v397 == int32(0))|base.B2i32(v397 != v400) != 0 {
		v418 = v397
		v419 = v400
		goto L68
	} else {
		goto L69
	}
L67:
	;
	F_pfree(m, v393)
	mBase = m.M
	v422 = m.ExcPending
	if v422 != 0 {
		goto L1
	} else {
		goto L74
	}
L68:
	;
	goto L67
L69:
	;
	v403 = v393
	v404 = v386
	goto L70
L70:
	;
	v407 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v404)+1)))
	v408 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v403)+1)))
	if v408 == int32(0) {
		v418 = v408
		v419 = v407
		goto L68
	} else {
		goto L72
	}
L71:
	;
	v418 = v408
	v419 = v407
	goto L68
L72:
	;
	v411 = int32(1)
	if v408 == v407 {
		v403 = v403 + v411
		v404 = v404 + v411
		goto L70
	} else {
		goto L73
	}
L73:
	;
	goto L71
L74:
	;
	if v418-v419 != 0 {
		goto L63
	} else {
		goto L75
	}
L75:
	;
	v423 = *(*int32)(unsafe.Add(mBase, uint32(v42)+176))
	*(*int32)(unsafe.Add(mBase, uint32(v42)+164)) = v423
	if v390 == int32(0) {
		goto L63
	} else {
		goto L76
	}
L76:
	;
	v429 = F_t_isalnum_with_len(m, v42+int32(164), v390)
	mBase = m.M
	v430 = m.ExcPending
	if v430 != 0 {
		goto L1
	} else {
		goto L77
	}
L77:
	;
	if v429 != 0 {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v431 = *(*int32)(unsafe.Add(mBase, uint32(v167)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v167)+4)) = v431 + int32(1)
	v435 = *(*int32)(unsafe.Add(mBase, uint32(v167)+8))
	v439 = *(*int32)(unsafe.Add(mBase, uint32(v42)+164))
	*(*int32)(unsafe.Add(mBase, uint32(v435+v431<<(uint(int32(2))%32)))) = v439
	goto L63
L79:
	;
	goto L80
L80:
	;
	v441 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v167)+1)) = uint8(v441)
	goto L63
L81:
	;
	goto L62
L82:
	;
	goto L37
L83:
	;
	goto L31
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v42)+84)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v42)+76)) = v585
	*(*int64)(unsafe.Add(mBase, uint32(v42)+164)) = int64(-8589934595)
	v594 = *(*int32)(unsafe.Add(mBase, uint32(v42+int32(32))+24))
	v595 = *(*int32)(unsafe.Add(mBase, uint32(v594)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v42)+172)) = v595
	v602 = F_hash_search(m, v585, v42+int32(164), int32(1), v42+int32(224))
	mBase = m.M
	v603 = m.ExcPending
	if v603 != 0 {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	v604 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+224)))
	if v604 == int32(1) {
		goto L87
	} else {
		goto L88
	}
L86:
	;
	v1167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1132)+20)))
	if v1167&int32(2) != 0 {
		v4595 = v5
		goto L21
	} else {
		goto L148
	}
L87:
	;
	v607 = *(*int32)(unsafe.Add(mBase, uint32(v602)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v602)+20)) = v607 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v42)+80)) = v602
	v1132 = v602
	v1139 = v5
	goto L86
L88:
	;
	goto L89
L89:
	;
	v612 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v602)+20)) = v612
	v614 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v602)+12)) = v614
	*(*int32)(unsafe.Add(mBase, uint32(v42)+84)) = int32(1)
	*(*int64)(unsafe.Add(mBase, uint32(v602)+32)) = v614
	*(*int64)(unsafe.Add(mBase, uint32(v602)+24)) = int64(4294967295)
	v623 = F_lappend(m, v612, v602)
	mBase = m.M
	v624 = m.ExcPending
	if v624 != 0 {
		goto L1
	} else {
		goto L90
	}
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v42)+88)) = v623
	v626 = *(*int32)(unsafe.Add(mBase, uint32(v602)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v602)+20)) = v626 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v42)+80)) = v602
	if v623 == int32(0) {
		v1132 = v602
		v1139 = v5
		goto L86
	} else {
		goto L91
	}
L91:
	;
	v633 = *(*int32)(unsafe.Add(mBase, uint32(v623)+4))
	if v633 <= int32(0) {
		v1132 = v602
		v1139 = v5
		goto L86
	} else {
		goto L92
	}
L92:
	;
	v650 = v5
	goto L93
L93:
	;
	v675 = *(*int32)(unsafe.Add(mBase, uint32(v623)+12))
	v679 = *(*int32)(unsafe.Add(mBase, uint32(v675+v650<<(uint(int32(2))%32))))
	v680 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+100)))
	if v680 == int32(1) {
		goto L96
	} else {
		goto L97
	}
L94:
	;
	v1127 = *(*int32)(unsafe.Add(mBase, uint32(v42)+80))
	v1132 = v1127
	v1139 = v1046
	goto L86
L95:
	;
	v1046 = *(*int32)(unsafe.Add(mBase, uint32(v42)+96))
	if v1046 <= int32(1024) {
		goto L139
	} else {
		goto L140
	}
L96:
	;
	v683 = *(*int32)(unsafe.Add(mBase, uint32(v679)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v679)+20)) = v683 | int32(2)
	goto L95
L97:
	;
	goto L98
L98:
	;
	v687 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v42)+92)) = v687
	F_addKey(m, v42-int32(-64), v679, v679)
	mBase = m.M
	v693 = m.ExcPending
	if v693 != 0 {
		goto L1
	} else {
		goto L99
	}
L99:
	;
	v694 = *(*int32)(unsafe.Add(mBase, uint32(v42)+92))
	if v694 == int32(0) {
		v763 = v687
		goto L100
	} else {
		goto L101
	}
L100:
	;
	F_list_free(m, v763)
	mBase = m.M
	v799 = m.ExcPending
	if v799 != 0 {
		goto L1
	} else {
		goto L112
	}
L101:
	;
	v697 = *(*int32)(unsafe.Add(mBase, uint32(v694)+4))
	if v697 <= int32(0) {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	v763 = v694
	goto L100
L103:
	;
	goto L104
L104:
	;
	v704 = v687
	goto L105
L105:
	;
	v739 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v679)+20)))
	if v739&int32(2) == int32(0) {
		goto L107
	} else {
		goto L108
	}
L106:
	;
	v758 = *(*int32)(unsafe.Add(mBase, uint32(v42)+92))
	v763 = v758
	goto L100
L107:
	;
	v746 = *(*int32)(unsafe.Add(mBase, uint32(v694)+12))
	v750 = *(*int32)(unsafe.Add(mBase, uint32(v746+v704<<(uint(int32(2))%32))))
	F_addKey(m, v42-int32(-64), v679, v750)
	mBase = m.M
	v752 = m.ExcPending
	if v752 != 0 {
		goto L1
	} else {
		goto L110
	}
L108:
	;
	goto L109
L109:
	;
	goto L106
L110:
	;
	v754 = v704 + int32(1)
	v755 = *(*int32)(unsafe.Add(mBase, uint32(v694)+4))
	if v754 < v755 {
		v704 = v754
		goto L105
	} else {
		goto L111
	}
L111:
	;
	goto L109
L112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v42)+92)) = int32(0)
	v802 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v679)+20)))
	if v802&int32(2) != 0 {
		goto L95
	} else {
		goto L113
	}
L113:
	;
	v805 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v42)+232)) = v805
	*(*int64)(unsafe.Add(mBase, uint32(v42)+224)) = int64(0)
	v809 = *(*int32)(unsafe.Add(mBase, uint32(v679)+16))
	if v809 == v805 {
		goto L95
	} else {
		goto L114
	}
L114:
	;
	v812 = int32(0)
	v813 = *(*int32)(unsafe.Add(mBase, uint32(v809)+4))
	if v813 <= v812 {
		goto L95
	} else {
		goto L115
	}
L115:
	;
	v828 = v812
	goto L116
L116:
	;
	v856 = *(*int32)(unsafe.Add(mBase, uint32(v42)+64))
	v857 = *(*int32)(unsafe.Add(mBase, uint32(v809)+12))
	v861 = *(*int32)(unsafe.Add(mBase, uint32(v857+v828<<(uint(int32(2))%32))))
	v862 = *(*int32)(unsafe.Add(mBase, uint32(v861)+8))
	v863 = F_pg_reg_getnumoutarcs(m, v856, v862)
	mBase = m.M
	v864 = m.ExcPending
	if v864 != 0 {
		goto L1
	} else {
		goto L118
	}
L117:
	;
	goto L95
L118:
	;
	v865 = F_palloc_mul(m, int32(8), v863)
	mBase = m.M
	v866 = m.ExcPending
	if v866 != 0 {
		goto L1
	} else {
		goto L119
	}
L119:
	;
	v867 = *(*int32)(unsafe.Add(mBase, uint32(v861)+8))
	F_pg_reg_getoutarcs(m, v856, v867, v865, v863)
	mBase = m.M
	v869 = m.ExcPending
	if v869 != 0 {
		goto L1
	} else {
		goto L120
	}
L120:
	;
	if int32(0) < v863 {
		goto L121
	} else {
		goto L122
	}
L121:
	;
	v877 = int32(0)
	goto L124
L122:
	;
	goto L123
L123:
	;
	F_pfree(m, v865)
	mBase = m.M
	v1002 = m.ExcPending
	if v1002 != 0 {
		goto L1
	} else {
		goto L136
	}
L124:
	;
	v914 = v865 + v877<<(uint(int32(3))%32)
	v915 = *(*int32)(unsafe.Add(mBase, uint32(v914)))
	if v915 < int32(0) {
		goto L126
	} else {
		goto L127
	}
L125:
	;
	goto L123
L126:
	;
	v960 = v877 + int32(1)
	if v960 != v863 {
		v877 = v960
		goto L124
	} else {
		goto L135
	}
L127:
	;
	v918 = *(*int32)(unsafe.Add(mBase, uint32(v42)+68))
	v921 = v918 + v915*int32(12)
	v922 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v921))))
	if v922 != int32(1) {
		goto L126
	} else {
		goto L128
	}
L128:
	;
	v925 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v921)+1)))
	if v925 == int32(1) {
		goto L129
	} else {
		goto L130
	}
L129:
	;
	v928 = *(*int32)(unsafe.Add(mBase, uint32(v861)+4))
	v929 = int32(-4)
	*(*int32)(unsafe.Add(mBase, uint32(v42)+228)) = v929
	*(*int32)(unsafe.Add(mBase, uint32(v42)+224)) = v928
	v932 = *(*int32)(unsafe.Add(mBase, uint32(v914)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v42)+232)) = v932
	F_addArc(m, v42-int32(-64), v679, v861, v929, v42+int32(224))
	mBase = m.M
	v940 = m.ExcPending
	if v940 != 0 {
		goto L1
	} else {
		goto L132
	}
L130:
	;
	goto L131
L131:
	;
	v942 = *(*int32)(unsafe.Add(mBase, uint32(v921)+4))
	if v942 <= int32(0) {
		goto L126
	} else {
		goto L133
	}
L132:
	;
	goto L131
L133:
	;
	v945 = *(*int32)(unsafe.Add(mBase, uint32(v861)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v42)+224)) = v945
	v947 = *(*int32)(unsafe.Add(mBase, uint32(v914)))
	*(*int32)(unsafe.Add(mBase, uint32(v42)+228)) = v947
	v949 = *(*int32)(unsafe.Add(mBase, uint32(v914)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v42)+232)) = v949
	F_addArc(m, v42-int32(-64), v679, v861, v947, v42+int32(224))
	mBase = m.M
	v956 = m.ExcPending
	if v956 != 0 {
		goto L1
	} else {
		goto L134
	}
L134:
	;
	goto L126
L135:
	;
	goto L125
L136:
	;
	v1004 = v828 + int32(1)
	v1005 = *(*int32)(unsafe.Add(mBase, uint32(v809)+4))
	if v1004 < v1005 {
		v828 = v1004
		goto L116
	} else {
		goto L137
	}
L137:
	;
	goto L117
L138:
	;
	v1124 = v650 + int32(1)
	v1125 = *(*int32)(unsafe.Add(mBase, uint32(v623)+4))
	if v1124 < v1125 {
		v650 = v1124
		goto L93
	} else {
		goto L147
	}
L139:
	;
	v1049 = *(*int32)(unsafe.Add(mBase, uint32(v42)+76))
	v1051 = *(*int32)(unsafe.Add(mBase, uint32(v1049)))
	v1052 = *(*int64)(unsafe.Add(mBase, uint32(v1051)+8))
	v1053 = *(*int64)(unsafe.Add(mBase, uint32(v1051)+808))
	if v1053 != int64(0) {
		goto L143
	} else {
		goto L144
	}
L140:
	;
	goto L141
L141:
	;
	v1121 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v42)+100)) = uint8(v1121)
	goto L138
L142:
	;
	if v1118 < int64(129) {
		goto L138
	} else {
		goto L146
	}
L143:
	;
	v1056 = *(*int64)(unsafe.Add(mBase, uint32(v1051)+752))
	v1057 = *(*int64)(unsafe.Add(mBase, uint32(v1051)+728))
	v1058 = *(*int64)(unsafe.Add(mBase, uint32(v1051)+704))
	v1059 = *(*int64)(unsafe.Add(mBase, uint32(v1051)+680))
	v1060 = *(*int64)(unsafe.Add(mBase, uint32(v1051)+656))
	v1061 = *(*int64)(unsafe.Add(mBase, uint32(v1051)+632))
	v1062 = *(*int64)(unsafe.Add(mBase, uint32(v1051)+608))
	v1063 = *(*int64)(unsafe.Add(mBase, uint32(v1051)+584))
	v1064 = *(*int64)(unsafe.Add(mBase, uint32(v1051)+560))
	v1065 = *(*int64)(unsafe.Add(mBase, uint32(v1051)+536))
	v1066 = *(*int64)(unsafe.Add(mBase, uint32(v1051)+512))
	v1067 = *(*int64)(unsafe.Add(mBase, uint32(v1051)+488))
	v1068 = *(*int64)(unsafe.Add(mBase, uint32(v1051)+464))
	v1069 = *(*int64)(unsafe.Add(mBase, uint32(v1051)+440))
	v1070 = *(*int64)(unsafe.Add(mBase, uint32(v1051)+416))
	v1071 = *(*int64)(unsafe.Add(mBase, uint32(v1051)+392))
	v1072 = *(*int64)(unsafe.Add(mBase, uint32(v1051)+368))
	v1073 = *(*int64)(unsafe.Add(mBase, uint32(v1051)+344))
	v1074 = *(*int64)(unsafe.Add(mBase, uint32(v1051)+320))
	v1075 = *(*int64)(unsafe.Add(mBase, uint32(v1051)+296))
	v1076 = *(*int64)(unsafe.Add(mBase, uint32(v1051)+272))
	v1077 = *(*int64)(unsafe.Add(mBase, uint32(v1051)+248))
	v1078 = *(*int64)(unsafe.Add(mBase, uint32(v1051)+224))
	v1079 = *(*int64)(unsafe.Add(mBase, uint32(v1051)+200))
	v1080 = *(*int64)(unsafe.Add(mBase, uint32(v1051)+176))
	v1081 = *(*int64)(unsafe.Add(mBase, uint32(v1051)+152))
	v1082 = *(*int64)(unsafe.Add(mBase, uint32(v1051)+128))
	v1083 = *(*int64)(unsafe.Add(mBase, uint32(v1051)+104))
	v1084 = *(*int64)(unsafe.Add(mBase, uint32(v1051)+80))
	v1085 = *(*int64)(unsafe.Add(mBase, uint32(v1051)+56))
	v1086 = *(*int64)(unsafe.Add(mBase, uint32(v1051)+32))
	v1118 = v1056 + (v1057 + (v1058 + (v1059 + (v1060 + (v1061 + (v1062 + (v1063 + (v1064 + (v1065 + (v1066 + (v1067 + (v1068 + (v1069 + (v1070 + (v1071 + (v1072 + (v1073 + (v1074 + (v1075 + (v1076 + (v1077 + (v1078 + (v1079 + (v1080 + (v1081 + (v1082 + (v1083 + (v1084 + (v1085 + (v1086 + v1052))))))))))))))))))))))))))))))
	goto L145
L144:
	;
	v1118 = v1052
	goto L145
L145:
	;
	goto L142
L146:
	;
	goto L141
L147:
	;
	goto L94
L148:
	;
	v1171 = F_palloc0_mul(m, int32(32), v1139)
	mBase = m.M
	v1172 = m.ExcPending
	if v1172 != 0 {
		goto L1
	} else {
		goto L149
	}
L149:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v42)+104)) = v1171
	v1175 = v42 + int32(176)
	v1176 = *(*int32)(unsafe.Add(mBase, uint32(v42)+76))
	F_hash_seq_init(m, v1175, v1176)
	mBase = m.M
	v1178 = m.ExcPending
	if v1178 != 0 {
		goto L1
	} else {
		goto L150
	}
L150:
	;
	v1179 = F_hash_seq_search(m, v1175)
	mBase = m.M
	v1180 = m.ExcPending
	if v1180 != 0 {
		goto L1
	} else {
		goto L151
	}
L151:
	;
	if v1179 != 0 {
		goto L152
	} else {
		goto L153
	}
L152:
	;
	v1183 = int32(0)
	v1190 = v1179
	goto L155
L153:
	;
	goto L154
L154:
	;
	if int32(2) <= v1139 {
		goto L167
	} else {
		goto L168
	}
L155:
	;
	v1221 = *(*int32)(unsafe.Add(mBase, uint32(v1190)+12))
	if v1221 == int32(0) {
		v1304 = v1183
		goto L157
	} else {
		goto L158
	}
L156:
	;
	goto L154
L157:
	;
	v1344 = F_hash_seq_search(m, v42+int32(176))
	mBase = m.M
	v1345 = m.ExcPending
	if v1345 != 0 {
		goto L1
	} else {
		goto L165
	}
L158:
	;
	v1224 = int32(0)
	v1225 = *(*int32)(unsafe.Add(mBase, uint32(v1221)+4))
	if v1225 <= v1224 {
		v1304 = v1183
		goto L157
	} else {
		goto L159
	}
L159:
	;
	v1229 = v1183
	v1235 = v1224
	goto L160
L160:
	;
	v1267 = *(*int32)(unsafe.Add(mBase, uint32(v1221)+12))
	v1271 = *(*int32)(unsafe.Add(mBase, uint32(v1267+v1235<<(uint(int32(2))%32))))
	v1273 = F_palloc(m, int32(8))
	mBase = m.M
	v1274 = m.ExcPending
	if v1274 != 0 {
		goto L1
	} else {
		goto L162
	}
L161:
	;
	v1304 = v1298
	goto L157
L162:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1273))) = v1190
	v1276 = *(*int32)(unsafe.Add(mBase, uint32(v1271)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v1273)+4)) = v1276
	v1280 = v1171 + v1229<<(uint(int32(5))%32)
	v1281 = *(*int32)(unsafe.Add(mBase, uint32(v1271)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1280)+8)) = v1281
	v1283 = *(*int64)(unsafe.Add(mBase, uint32(v1271)))
	*(*int64)(unsafe.Add(mBase, uint32(v1280))) = v1283
	v1285 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1280)+24)) = uint8(v1285)
	*(*int32)(unsafe.Add(mBase, uint32(v1280)+12)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v42)+12)) = v1273
	*(*int32)(unsafe.Add(mBase, uint32(v42)+224)) = v1273
	v1294 = F_list_make1_impl(m, v1285, v42+int32(12))
	mBase = m.M
	v1295 = m.ExcPending
	if v1295 != 0 {
		goto L1
	} else {
		goto L163
	}
L163:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1280)+28)) = v1294
	v1297 = int32(1)
	v1298 = v1229 + v1297
	v1300 = v1235 + v1297
	v1301 = *(*int32)(unsafe.Add(mBase, uint32(v1221)+4))
	if v1300 < v1301 {
		v1229 = v1298
		v1235 = v1300
		goto L160
	} else {
		goto L164
	}
L164:
	;
	goto L161
L165:
	;
	if v1344 != 0 {
		v1183 = v1304
		v1190 = v1344
		goto L155
	} else {
		goto L166
	}
L166:
	;
	goto L156
L167:
	;
	F_pg_qsort(m, v1171, v1139, int32(32), int32(_a_F_createTrgmNFA_5))
	mBase = m.M
	v1390 = m.ExcPending
	if v1390 != 0 {
		goto L1
	} else {
		goto L170
	}
L168:
	;
	v1580 = v1139
	goto L169
L169:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v42)+108)) = v1580
	if int32(0) < v1580 {
		goto L187
	} else {
		goto L188
	}
L170:
	;
	__phi1396 = v1171 + int32(32)
	__phi1400 = v1171
	__phi1402 = v1171
	v1396 = __phi1396
	v1400 = __phi1400
	v1402 = __phi1402
	goto L171
L171:
	;
	v1435 = *(*int64)(unsafe.Add(mBase, uint32(v1402)+32))
	v1436 = int64(56)
	v1438 = int64(65280)
	v1440 = int64(40)
	v1443 = int64(16711680)
	v1445 = int64(24)
	v1447 = int64(4278190080)
	v1449 = int64(8)
	v1470 = v1435<<(uint(v1436)%64) | v1435&v1438<<(uint(v1440)%64) | (v1435&v1443<<(uint(v1445)%64) | v1435&v1447<<(uint(v1449)%64)) | (int64(base.Ui64(v1435)>>(uint(v1449)%64))&v1447 | int64(base.Ui64(v1435)>>(uint(v1445)%64))&v1443 | (int64(base.Ui64(v1435)>>(uint(v1440)%64))&v1438 | int64(base.Ui64(v1435)>>(uint(v1436)%64))))
	v1471 = *(*int64)(unsafe.Add(mBase, uint32(v1400)))
	v1506 = v1471<<(uint(v1436)%64) | v1471&v1438<<(uint(v1440)%64) | (v1471&v1443<<(uint(v1445)%64) | v1471&v1447<<(uint(v1449)%64)) | (int64(base.Ui64(v1471)>>(uint(v1449)%64))&v1447 | int64(base.Ui64(v1471)>>(uint(v1445)%64))&v1443 | (int64(base.Ui64(v1471)>>(uint(v1440)%64))&v1438 | int64(base.Ui64(v1471)>>(uint(v1436)%64))))
	if v1470 == v1506 {
		goto L175
	} else {
		goto L176
	}
L172:
	;
	v1580 = (v1560-v1171)>>(uint(int32(5))%32) + int32(1)
	goto L169
L173:
	;
	v1562 = v1396 + int32(32)
	if base.Ui32(v1562) < base.Ui32(v1171+v1139<<(uint(int32(5))%32)) {
		__phi1396 = v1562
		__phi1400 = v1560
		__phi1402 = v1396
		v1396 = __phi1396
		v1400 = __phi1400
		v1402 = __phi1402
		goto L171
	} else {
		goto L186
	}
L174:
	;
	if int32(0) < v1542 {
		goto L182
	} else {
		goto L183
	}
L175:
	;
	v1509 = *(*int32)(unsafe.Add(mBase, uint32(v1396)+8))
	v1510 = int32(16711935)
	v1512 = int32(8)
	v1514 = int32(24)
	v1519 = base.I64_extend_i32_u(base.I32_rotr(v1509&v1510, v1512) | base.I32_rotr(v1509, v1514)&v1510)
	v1520 = *(*int32)(unsafe.Add(mBase, uint32(v1400)+8))
	v1530 = base.I64_extend_i32_u(base.I32_rotr(v1520&v1510, v1512) | base.I32_rotr(v1520, v1514)&v1510)
	if v1519 == v1530 {
		v1542 = int32(0)
		goto L174
	} else {
		goto L178
	}
L176:
	;
	v1533 = v1506
	v1534 = v1470
	goto L177
L177:
	;
	if base.Ui64(v1534) < base.Ui64(v1533) {
		goto L179
	} else {
		goto L180
	}
L178:
	;
	v1533 = v1530
	v1534 = v1519
	goto L177
L179:
	;
	v1538 = int32(-1)
	goto L181
L180:
	;
	v1538 = int32(1)
	goto L181
L181:
	;
	v1542 = v1538
	goto L174
L182:
	;
	v1545 = *(*int64)(unsafe.Add(mBase, uint32(v1396)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v1400)+56)) = v1545
	v1547 = *(*int64)(unsafe.Add(mBase, uint32(v1396)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v1400)+48)) = v1547
	v1549 = *(*int64)(unsafe.Add(mBase, uint32(v1396)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v1400)+40)) = v1549
	v1551 = *(*int64)(unsafe.Add(mBase, uint32(v1396)))
	*(*int64)(unsafe.Add(mBase, uint32(v1400)+32)) = v1551
	v1560 = v1400 + int32(32)
	goto L173
L183:
	;
	goto L184
L184:
	;
	v1555 = *(*int32)(unsafe.Add(mBase, uint32(v1400)+28))
	v1556 = *(*int32)(unsafe.Add(mBase, uint32(v1402)+60))
	v1557 = F_list_concat(m, v1555, v1556)
	mBase = m.M
	v1558 = m.ExcPending
	if v1558 != 0 {
		goto L1
	} else {
		goto L185
	}
L185:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1400)+28)) = v1557
	v1560 = v1400
	goto L173
L186:
	;
	goto L172
L187:
	;
	v1613 = *(*int32)(unsafe.Add(mBase, uint32(v42)+68))
	v1621 = int32(0)
	v1649 = int64(0)
	v1651 = float32(0)
	goto L190
L188:
	;
	goto L189
L189:
	;
	F_pg_qsort(m, v1171, v1580, int32(32), int32(_a_F_createTrgmNFA_6))
	mBase = m.M
	v2715 = m.ExcPending
	if v2715 != 0 {
		goto L1
	} else {
		goto L303
	}
L190:
	;
	v1657 = v1171 + v1621<<(uint(int32(5))%32)
	v1658 = *(*int32)(unsafe.Add(mBase, uint32(v1657)))
	if v1658 != int32(-4) {
		goto L192
	} else {
		goto L193
	}
L191:
	;
	F_pg_qsort(m, v1171, v1580, int32(32), int32(_a_F_createTrgmNFA_6))
	mBase = m.M
	v1712 = m.ExcPending
	if v1712 != 0 {
		goto L1
	} else {
		goto L204
	}
L192:
	;
	v1665 = *(*int32)(unsafe.Add(mBase, uint32(v1613+v1658*int32(12))+4))
	v1666 = v1665
	v1667 = int32(0)
	goto L194
L193:
	;
	v1666 = int32(1)
	v1667 = int32(2)
	goto L194
L194:
	;
	v1668 = *(*int32)(unsafe.Add(mBase, uint32(v1657)+4))
	if v1668 != int32(-4) {
		goto L196
	} else {
		goto L197
	}
L195:
	;
	v1681 = v1679 << (uint(int32(1)) % 32)
	v1682 = *(*int32)(unsafe.Add(mBase, uint32(v1657)+8))
	if v1682 != int32(-4) {
		goto L200
	} else {
		goto L201
	}
L196:
	;
	v1674 = *(*int32)(unsafe.Add(mBase, uint32(v1613+v1668*int32(12))+4))
	v1678 = v1674 * v1666
	v1679 = v1667
	goto L195
L197:
	;
	goto L198
L198:
	;
	v1678 = v1666
	v1679 = v1667 | int32(1)
	goto L195
L199:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1657)+16)) = v1692
	v1699 = *(*float32)(unsafe.Add(mBase, uint32(v1693<<(uint(int32(2))%32))+uint32(_c_F_createTrgmNFA[1])))
	v1701 = base.F32_mul(v1699, base.F32_convert_i32_s(v1692))
	*(*float32)(unsafe.Add(mBase, uint32(v1657)+20)) = v1701
	v1703 = base.F32_add(v1651, v1701)
	v1705 = v1649 + base.I64_extend_i32_s(v1692)
	v1707 = v1621 + int32(1)
	if v1707 != v1580 {
		v1621 = v1707
		v1649 = v1705
		v1651 = v1703
		goto L190
	} else {
		goto L203
	}
L200:
	;
	v1688 = *(*int32)(unsafe.Add(mBase, uint32(v1613+v1682*int32(12))+4))
	v1692 = v1688 * v1678
	v1693 = v1681
	goto L199
L201:
	;
	goto L202
L202:
	;
	v1692 = v1678
	v1693 = v1681 | int32(1)
	goto L199
L203:
	;
	goto L191
L204:
	;
	v1730 = v5
	v1748 = v1705
	v1750 = v1703
	goto L205
L205:
	;
	if base.F32_le(v1750, float32(16)) == int32(0) {
		goto L207
	} else {
		goto L208
	}
L206:
	;
	if v2520 <= int64(256) {
		goto L274
	} else {
		goto L275
	}
L207:
	;
	v1758 = v1171 + v1730<<(uint(int32(5))%32)
	v1759 = *(*int32)(unsafe.Add(mBase, uint32(v1758)+28))
	if v1759 == int32(0) {
		goto L211
	} else {
		goto L212
	}
L208:
	;
	v2520 = v1748
	goto L209
L209:
	;
	goto L206
L210:
	;
	v2483 = v1730 + int32(1)
	if v2483 != v1580 {
		v1730 = v2483
		v1748 = v2478
		v1750 = v2480
		goto L205
	} else {
		goto L273
	}
L211:
	;
	v2437 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1758)+24)) = uint8(v2437)
	v2439 = *(*float32)(unsafe.Add(mBase, uint32(v1758)+20))
	v2441 = int64(*(*int32)(unsafe.Add(mBase, uint32(v1758)+16)))
	v2478 = v1748 - v2441
	v2480 = base.F32_sub(v1750, v2439)
	goto L210
L212:
	;
	v1762 = *(*int32)(unsafe.Add(mBase, uint32(v1759)+4))
	if v1762 <= int32(0) {
		goto L211
	} else {
		goto L213
	}
L213:
	;
	v1765 = *(*int32)(unsafe.Add(mBase, uint32(v1759)+12))
	v1775 = v1762
	v1779 = int32(0)
	goto L214
L214:
	;
	v1809 = *(*int32)(unsafe.Add(mBase, uint32(v1765+v1779<<(uint(int32(2))%32))))
	v1810 = *(*int32)(unsafe.Add(mBase, uint32(v1809)+4))
	v1811 = *(*int32)(unsafe.Add(mBase, uint32(v1809)))
	v1812 = v1811
	goto L216
L215:
	;
	v2004 = int32(0)
	if v2004 < v2002 {
		goto L239
	} else {
		goto L240
	}
L216:
	;
	v1851 = *(*int32)(unsafe.Add(mBase, uint32(v1812)+28))
	if v1851 != 0 {
		v1812 = v1851
		goto L216
	} else {
		goto L218
	}
L217:
	;
	v1858 = v1810
	goto L219
L218:
	;
	goto L217
L219:
	;
	v1891 = *(*int32)(unsafe.Add(mBase, uint32(v1858)+28))
	if v1891 != 0 {
		v1858 = v1891
		goto L219
	} else {
		goto L221
	}
L220:
	;
	v1892 = *(*int32)(unsafe.Add(mBase, uint32(v1812)+32))
	v1893 = *(*int32)(unsafe.Add(mBase, uint32(v1812)+20))
	v1899 = v1812
	v1902 = v1892 | v1893
	v1904 = v1892
	goto L222
L221:
	;
	goto L220
L222:
	;
	v1934 = *(*int32)(unsafe.Add(mBase, uint32(v1899)+36))
	if v1934 != 0 {
		goto L224
	} else {
		goto L225
	}
L223:
	;
	v1939 = *(*int32)(unsafe.Add(mBase, uint32(v1858)+32))
	v1940 = *(*int32)(unsafe.Add(mBase, uint32(v1858)+20))
	v1942 = v1858
	v1948 = v1939 | v1940
	goto L227
L224:
	;
	v1935 = *(*int32)(unsafe.Add(mBase, uint32(v1934)+20))
	v1937 = *(*int32)(unsafe.Add(mBase, uint32(v1934)+32))
	v1899 = v1934
	v1902 = v1935 | v1902 | v1937
	v1904 = v1937
	goto L222
L225:
	;
	goto L226
L226:
	;
	goto L223
L227:
	;
	v1981 = *(*int32)(unsafe.Add(mBase, uint32(v1942)+36))
	if v1981 != 0 {
		goto L229
	} else {
		goto L230
	}
L228:
	;
	v1987 = int32(3)
	v1990 = base.B2i32((v1948|v1902)&v1987 == v1987)
	if v1990 == int32(0) {
		goto L232
	} else {
		goto L233
	}
L229:
	;
	v1982 = *(*int32)(unsafe.Add(mBase, uint32(v1981)+20))
	v1984 = *(*int32)(unsafe.Add(mBase, uint32(v1981)+32))
	v1942 = v1981
	v1948 = v1982 | v1948 | v1984
	goto L227
L230:
	;
	goto L231
L231:
	;
	goto L228
L232:
	;
	if v1942 != v1899 {
		goto L235
	} else {
		goto L236
	}
L233:
	;
	v2002 = v1775
	goto L234
L234:
	;
	goto L215
L235:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1942)+36)) = v1899
	*(*int32)(unsafe.Add(mBase, uint32(v1899)+32)) = v1948 | v1904
	v1997 = *(*int32)(unsafe.Add(mBase, uint32(v1759)+4))
	v1998 = v1997
	goto L237
L236:
	;
	v1998 = v1775
	goto L237
L237:
	;
	v2000 = v1779 + int32(1)
	if v2000 < v1998 {
		v1775 = v1998
		v1779 = v2000
		goto L214
	} else {
		goto L238
	}
L238:
	;
	v2002 = v1998
	goto L234
L239:
	;
	v2014 = v2004
	goto L242
L240:
	;
	v2228 = v2002
	goto L241
L241:
	;
	if (v1948|v1902)&v1987 == v1987 {
		v2478 = v1748
		v2480 = v1750
		goto L210
	} else {
		goto L259
	}
L242:
	;
	v2049 = *(*int32)(unsafe.Add(mBase, uint32(v1765+v2014<<(uint(int32(2))%32))))
	v2050 = *(*int32)(unsafe.Add(mBase, uint32(v2049)+4))
	v2051 = *(*int32)(unsafe.Add(mBase, uint32(v2049)))
	v2052 = v2051
	goto L244
L243:
	;
	v2228 = v2218
	goto L241
L244:
	;
	v2091 = *(*int32)(unsafe.Add(mBase, uint32(v2052)+28))
	if v2091 != 0 {
		v2052 = v2091
		goto L244
	} else {
		goto L246
	}
L245:
	;
	v2098 = v2050
	goto L247
L246:
	;
	goto L245
L247:
	;
	v2131 = *(*int32)(unsafe.Add(mBase, uint32(v2098)+28))
	if v2131 != 0 {
		v2098 = v2131
		goto L247
	} else {
		goto L249
	}
L248:
	;
	v2136 = v2052
	goto L250
L249:
	;
	goto L248
L250:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2136)+32)) = int32(0)
	v2173 = *(*int32)(unsafe.Add(mBase, uint32(v2136)+36))
	if v2173 != 0 {
		v2136 = v2173
		goto L250
	} else {
		goto L252
	}
L251:
	;
	v2174 = v2098
	goto L253
L252:
	;
	goto L251
L253:
	;
	v2213 = *(*int32)(unsafe.Add(mBase, uint32(v2174)+36))
	if v2213 != 0 {
		goto L255
	} else {
		goto L256
	}
L254:
	;
	v2217 = v2014 + int32(1)
	v2218 = *(*int32)(unsafe.Add(mBase, uint32(v1759)+4))
	if v2217 < v2218 {
		v2014 = v2217
		goto L242
	} else {
		goto L258
	}
L255:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v2174)+32)) = int64(0)
	v2174 = v2213
	goto L253
L256:
	;
	goto L257
L257:
	;
	goto L254
L258:
	;
	goto L243
L259:
	;
	v2259 = int32(0)
	if v2228 <= v2259 {
		goto L211
	} else {
		goto L260
	}
L260:
	;
	v2269 = v2259
	v2270 = v2228
	goto L261
L261:
	;
	v2304 = *(*int32)(unsafe.Add(mBase, uint32(v1765+v2269<<(uint(int32(2))%32))))
	v2305 = *(*int32)(unsafe.Add(mBase, uint32(v2304)+4))
	v2306 = *(*int32)(unsafe.Add(mBase, uint32(v2304)))
	v2307 = v2306
	goto L263
L262:
	;
	goto L211
L263:
	;
	v2346 = *(*int32)(unsafe.Add(mBase, uint32(v2307)+28))
	if v2346 != 0 {
		v2307 = v2346
		goto L263
	} else {
		goto L265
	}
L264:
	;
	v2351 = v2305
	goto L266
L265:
	;
	goto L264
L266:
	;
	v2386 = *(*int32)(unsafe.Add(mBase, uint32(v2351)+28))
	if v2386 != 0 {
		v2351 = v2386
		goto L266
	} else {
		goto L268
	}
L267:
	;
	if v2351 != v2307 {
		goto L269
	} else {
		goto L270
	}
L268:
	;
	goto L267
L269:
	;
	v2388 = *(*int32)(unsafe.Add(mBase, uint32(v2307)+20))
	v2389 = *(*int32)(unsafe.Add(mBase, uint32(v2351)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v2307)+20)) = v2388 | v2389
	*(*int32)(unsafe.Add(mBase, uint32(v2351)+28)) = v2307
	v2393 = *(*int32)(unsafe.Add(mBase, uint32(v1759)+4))
	v2394 = v2393
	goto L271
L270:
	;
	v2394 = v2270
	goto L271
L271:
	;
	v2396 = v2269 + int32(1)
	if v2396 < v2394 {
		v2269 = v2396
		v2270 = v2394
		goto L261
	} else {
		goto L272
	}
L272:
	;
	goto L262
L273:
	;
	v2520 = v2478
	goto L209
L274:
	;
	v2526 = base.I32_wrap_i64(v2520)
	*(*int32)(unsafe.Add(mBase, uint32(v42)+112)) = v2526
	F_pg_qsort(m, v1171, v1580, int32(32), int32(_a_F_createTrgmNFA_5))
	mBase = m.M
	v2531 = m.ExcPending
	if v2531 != 0 {
		goto L1
	} else {
		goto L277
	}
L275:
	;
	goto L276
L276:
	;
	v4595 = v5
	goto L21
L277:
	;
	v2533 = v1580 & int32(3)
	v2534 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v1580) {
		goto L278
	} else {
		goto L279
	}
L278:
	;
	v2542 = v2534
	v2543 = int32(0)
	v2546 = v2534
	goto L281
L279:
	;
	v2619 = v2534
	v2623 = v2534
	goto L280
L280:
	;
	v2658 = v2619
	v2662 = v2623
	v2665 = v2534
	goto L297
L281:
	;
	v2583 = v1171 + v2542<<(uint(int32(5))%32)
	v2584 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2583)+24)))
	if v2584 == int32(1) {
		goto L283
	} else {
		goto L284
	}
L282:
	;
	if v2533 == int32(0) {
		v2757 = v2526
		goto L22
	} else {
		goto L296
	}
L283:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2583)+12)) = v2546
	v2590 = v2546 + int32(1)
	goto L285
L284:
	;
	v2590 = v2546
	goto L285
L285:
	;
	v2591 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2583)+56)))
	if v2591 == int32(1) {
		goto L286
	} else {
		goto L287
	}
L286:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2583)+44)) = v2590
	v2597 = v2590 + int32(1)
	goto L288
L287:
	;
	v2597 = v2590
	goto L288
L288:
	;
	v2598 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2583)+88)))
	if v2598 == int32(1) {
		goto L289
	} else {
		goto L290
	}
L289:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2583)+76)) = v2597
	v2604 = v2597 + int32(1)
	goto L291
L290:
	;
	v2604 = v2597
	goto L291
L291:
	;
	v2605 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2583)+120)))
	if v2605 == int32(1) {
		goto L292
	} else {
		goto L293
	}
L292:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2583)+108)) = v2604
	v2611 = v2604 + int32(1)
	goto L294
L293:
	;
	v2611 = v2604
	goto L294
L294:
	;
	v2612 = int32(4)
	v2613 = v2542 + v2612
	v2615 = v2543 + v2612
	if v2615 != v1580&int32(2147483644) {
		v2542 = v2613
		v2543 = v2615
		v2546 = v2611
		goto L281
	} else {
		goto L295
	}
L295:
	;
	goto L282
L296:
	;
	v2619 = v2613
	v2623 = v2611
	goto L280
L297:
	;
	v2699 = v1171 + v2658<<(uint(int32(5))%32)
	v2700 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2699)+24)))
	if v2700 == int32(1) {
		goto L299
	} else {
		goto L300
	}
L298:
	;
	v2757 = v2526
	goto L22
L299:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2699)+12)) = v2662
	v2706 = v2662 + int32(1)
	goto L301
L300:
	;
	v2706 = v2662
	goto L301
L301:
	;
	v2707 = int32(1)
	v2710 = v2665 + v2707
	if v2710 != v2533 {
		v2658 = v2658 + v2707
		v2662 = v2706
		v2665 = v2710
		goto L297
	} else {
		goto L302
	}
L302:
	;
	goto L298
L303:
	;
	v2716 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v42)+112)) = v2716
	F_pg_qsort(m, v1171, v1580, int32(32), int32(_a_F_createTrgmNFA_5))
	mBase = m.M
	v2722 = m.ExcPending
	if v2722 != 0 {
		goto L1
	} else {
		goto L304
	}
L304:
	;
	v2757 = v2716
	goto L22
L305:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2730 = m.ExcPending
	if v2730 != 0 {
		goto L1
	} else {
		goto L306
	}
L306:
	;
	F_errcode(m, int32(302252162))
	mBase = m.M
	v2733 = m.ExcPending
	if v2733 != 0 {
		goto L1
	} else {
		goto L307
	}
L307:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v42)+16)) = v2724
	F_errmsg(m, int32(_a_F_createTrgmNFA_7), v42+int32(16))
	mBase = m.M
	v2739 = m.ExcPending
	if v2739 != 0 {
		goto L1
	} else {
		goto L308
	}
L308:
	;
	F_errfinish(m, int32(_a_F_createTrgmNFA_8), int32(751), int32(_a_F_createTrgmNFA_9))
	mBase = m.M
	v2744 = m.ExcPending
	if v2744 != 0 {
		goto L1
	} else {
		goto L309
	}
L309:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L310:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2790))) = v2757*int32(12) + int32(20)
	v2797 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v2790)+4)) = uint8(v2797)
	if int32(0) < v1580 {
		goto L311
	} else {
		goto L312
	}
L311:
	;
	v2810 = v42 + int32(177)
	v2818 = v2790 + int32(5)
	v2822 = v1580
	v2832 = v5
	goto L314
L312:
	;
	v3443 = v1176
	goto L313
L313:
	;
	v3469 = v42 + int32(176)
	F_hash_seq_init(m, v3469, v3443)
	mBase = m.M
	v3471 = m.ExcPending
	if v3471 != 0 {
		goto L1
	} else {
		goto L397
	}
L314:
	;
	v2850 = *(*int32)(unsafe.Add(mBase, uint32(v42)+104))
	v2853 = v2850 + v2832<<(uint(int32(5))%32)
	v2854 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2853)+24)))
	if v2854 != int32(1) {
		v3393 = v2818
		v3397 = v2822
		goto L316
	} else {
		goto L317
	}
L315:
	;
	v3428 = *(*int32)(unsafe.Add(mBase, uint32(v42)+76))
	v3443 = v3428
	goto L313
L316:
	;
	v3426 = v2832 + int32(1)
	if v3426 < v3397 {
		v2818 = v3393
		v2822 = v3397
		v2832 = v3426
		goto L314
	} else {
		goto L396
	}
L317:
	;
	v2857 = *(*int32)(unsafe.Add(mBase, uint32(v42)+68))
	v2858 = *(*int32)(unsafe.Add(mBase, uint32(v2853)))
	v2861 = v2857 + v2858*int32(12)
	v2862 = int32(1)
	v2865 = base.B2i32(v2858 == int32(-4))
	if v2865 == int32(0) {
		goto L318
	} else {
		goto L319
	}
L318:
	;
	v2868 = *(*int32)(unsafe.Add(mBase, uint32(v2861)+4))
	if v2868 <= int32(0) {
		v3393 = v2818
		v3397 = v2822
		goto L316
	} else {
		goto L321
	}
L319:
	;
	v2871 = v2862
	goto L320
L320:
	;
	v2872 = *(*int32)(unsafe.Add(mBase, uint32(v2853)+4))
	v2873 = int32(12)
	v2875 = v2857 + v2872*v2873
	v2876 = *(*int32)(unsafe.Add(mBase, uint32(v2853)+8))
	v2880 = base.B2i32(v2872 == int32(-4))
	if v2880 == int32(0) {
		goto L322
	} else {
		goto L323
	}
L321:
	;
	v2871 = v2868
	goto L320
L322:
	;
	v2883 = *(*int32)(unsafe.Add(mBase, uint32(v2875)+4))
	v2884 = v2883
	goto L324
L323:
	;
	v2884 = v2862
	goto L324
L324:
	;
	v2885 = v2876*v2873 + v2857
	v2887 = v2871
	v2891 = v2884
	v2894 = v2818
	v2900 = int32(0)
	goto L325
L325:
	;
	if int32(0) < v2891 {
		goto L328
	} else {
		goto L329
	}
L326:
	;
	v3385 = *(*int32)(unsafe.Add(mBase, uint32(v42)+108))
	v3393 = v3353
	v3397 = v3385
	goto L316
L327:
	;
	goto L326
L328:
	;
	if v2858 == int32(-4) {
		goto L331
	} else {
		goto L332
	}
L329:
	;
	v3304 = v2887
	v3308 = v2891
	v3311 = v2894
	goto L330
L330:
	;
	v3344 = v2900 + int32(1)
	if v3344 < v3304 {
		v2887 = v3304
		v2891 = v3308
		v2894 = v3311
		v2900 = v3344
		goto L325
	} else {
		goto L395
	}
L331:
	;
	v2931 = v42 + int32(224)
	goto L333
L332:
	;
	v2930 = *(*int32)(unsafe.Add(mBase, uint32(v2861)+8))
	v2931 = v2930
	goto L333
L333:
	;
	v2935 = *(*int32)(unsafe.Add(mBase, uint32(v2931+v2900<<(uint(int32(2))%32))))
	v2938 = base.B2i32(v2876 == int32(-4))
	if v2938 == int32(0) {
		goto L334
	} else {
		goto L335
	}
L334:
	;
	v2941 = *(*int32)(unsafe.Add(mBase, uint32(v2885)+4))
	v2942 = v2941
	goto L336
L335:
	;
	v2942 = int32(1)
	goto L336
L336:
	;
	v2944 = int32(base.Ui32(v2935) >> (uint(int32(24)) % 32))
	v2946 = int32(base.Ui32(v2935) >> (uint(int32(16)) % 32))
	v2948 = int32(base.Ui32(v2935) >> (uint(int32(8)) % 32))
	v2952 = v2942
	v2956 = v2891
	v2959 = v2894
	v2966 = int32(0)
	goto L337
L337:
	;
	if int32(0) < v2952 {
		goto L340
	} else {
		goto L341
	}
L338:
	;
	if v2858 == int32(-4) {
		v3353 = v3271
		goto L327
	} else {
		goto L394
	}
L339:
	;
	goto L338
L340:
	;
	if v2872 == int32(-4) {
		goto L343
	} else {
		goto L344
	}
L341:
	;
	v3222 = v2952
	v3226 = v2956
	v3229 = v2959
	goto L342
L342:
	;
	v3262 = v2966 + int32(1)
	if v3262 < v3226 {
		v2952 = v3222
		v2956 = v3226
		v2959 = v3229
		v2966 = v3262
		goto L337
	} else {
		goto L393
	}
L343:
	;
	v2996 = v42 + int32(224)
	goto L345
L344:
	;
	v2995 = *(*int32)(unsafe.Add(mBase, uint32(v2875)+8))
	v2996 = v2995
	goto L345
L345:
	;
	v3000 = *(*int32)(unsafe.Add(mBase, uint32(v2996+v2966<<(uint(int32(2))%32))))
	v3002 = int32(base.Ui32(v3000) >> (uint(int32(24)) % 32))
	v3004 = int32(base.Ui32(v3000) >> (uint(int32(16)) % 32))
	v3006 = int32(base.Ui32(v3000) >> (uint(int32(8)) % 32))
	v3009 = int32(0)
	v3015 = v2959
	goto L346
L346:
	;
	if v2876 == int32(-4) {
		goto L348
	} else {
		goto L349
	}
L347:
	;
	if v2872 == int32(-4) {
		v3268 = int32(1)
		v3271 = v3212
		goto L339
	} else {
		goto L392
	}
L348:
	;
	v3050 = v42 + int32(224)
	goto L350
L349:
	;
	v3049 = *(*int32)(unsafe.Add(mBase, uint32(v2885)+8))
	v3050 = v3049
	goto L350
L350:
	;
	v3054 = *(*int32)(unsafe.Add(mBase, uint32(v3050+v3009<<(uint(int32(2))%32))))
	if v2935&int32(255) != 0 {
		goto L352
	} else {
		goto L353
	}
L351:
	;
	v3073 = v3071 + int32(1)
	if v3000&int32(255) != 0 {
		goto L359
	} else {
		goto L360
	}
L352:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v42)+176)) = uint8(v2935)
	if v2948&int32(255) == int32(0) {
		v3071 = v2810
		goto L351
	} else {
		goto L355
	}
L353:
	;
	goto L354
L354:
	;
	v3069 = int32(32)
	*(*uint8)(unsafe.Add(mBase, uint32(v42)+176)) = uint8(v3069)
	v3071 = v2810
	goto L351
L355:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v42)+177)) = uint8(v2948)
	if v2946&int32(255) == int32(0) {
		v3071 = v42 + int32(178)
		goto L351
	} else {
		goto L356
	}
L356:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v42)+178)) = uint8(v2946)
	if v2944 == int32(0) {
		v3071 = v42 + int32(179)
		goto L351
	} else {
		goto L357
	}
L357:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v42)+179)) = uint8(v2944)
	v3071 = v42 + int32(180)
	goto L351
L358:
	;
	v3100 = v3098 + int32(1)
	if v3054&int32(255) != 0 {
		goto L370
	} else {
		goto L371
	}
L359:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v3071))) = uint8(v3000)
	if v3006&int32(255) == int32(0) {
		v3098 = v3073
		goto L358
	} else {
		goto L362
	}
L360:
	;
	goto L361
L361:
	;
	v3096 = int32(32)
	*(*uint8)(unsafe.Add(mBase, uint32(v3071))) = uint8(v3096)
	v3098 = v3073
	goto L358
L362:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v3071)+1)) = uint8(v3006)
	if v3004&int32(255) == int32(0) {
		goto L363
	} else {
		goto L364
	}
L363:
	;
	v3098 = v3071 + int32(2)
	goto L358
L364:
	;
	goto L365
L365:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v3071)+2)) = uint8(v3004)
	if v3002 == int32(0) {
		goto L366
	} else {
		goto L367
	}
L366:
	;
	v3098 = v3071 + int32(3)
	goto L358
L367:
	;
	goto L368
L368:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v3071)+3)) = uint8(v3002)
	v3098 = v3071 + int32(4)
	goto L358
L369:
	;
	v3134 = v42 + int32(176)
	v3135 = v3131 - v3134
	v3141 = int32(255)
	switch v3135 {
	case 0:
		v3193 = v3135
		goto L381
	default:
		goto L382
	case 3:
		goto L383
	}
L370:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v3098))) = uint8(v3054)
	v3105 = int32(base.Ui32(v3054) >> (uint(int32(8)) % 32))
	if v3105&int32(255) == int32(0) {
		v3131 = v3100
		goto L369
	} else {
		goto L373
	}
L371:
	;
	goto L372
L372:
	;
	v3129 = int32(32)
	*(*uint8)(unsafe.Add(mBase, uint32(v3098))) = uint8(v3129)
	v3131 = v3100
	goto L369
L373:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v3098)+1)) = uint8(v3105)
	v3112 = int32(base.Ui32(v3054) >> (uint(int32(16)) % 32))
	if v3112&int32(255) == int32(0) {
		goto L374
	} else {
		goto L375
	}
L374:
	;
	v3131 = v3098 + int32(2)
	goto L369
L375:
	;
	goto L376
L376:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v3098)+2)) = uint8(v3112)
	v3121 = int32(base.Ui32(v3054) >> (uint(int32(24)) % 32))
	if v3121 == int32(0) {
		goto L377
	} else {
		goto L378
	}
L377:
	;
	v3131 = v3098 + int32(3)
	goto L369
L378:
	;
	goto L379
L379:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v3098)+3)) = uint8(v3121)
	v3131 = v3098 + int32(4)
	goto L369
L380:
	;
	v3212 = v3015 + int32(3)
	if v2876 == int32(-4) {
		goto L388
	} else {
		goto L389
	}
L381:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v3015))) = uint16(v3193)
	v3201 = int32(base.Ui32(v3193) >> (uint(int32(16)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v3015)+2)) = uint8(v3201)
	goto L380
L382:
	;
	v3152 = v3134
	v3153 = v3135
	v3155 = v3141
	v3156 = v3141
	v3157 = v3141
	v3158 = v3141
	goto L384
L383:
	;
	v3145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3134))))
	*(*uint8)(unsafe.Add(mBase, uint32(v3015))) = uint8(v3145)
	v3147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3134)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v3015)+1)) = uint8(v3147)
	v3149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3134)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v3015)+2)) = uint8(v3149)
	goto L380
L384:
	;
	v3159 = int32(8)
	v3161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3152))))
	v3167 = *(*int32)(unsafe.Add(mBase, uint32((v3161^v3155)<<(uint(int32(2))%32))+uint32(_c_F_createTrgmNFA[2])))
	v3170 = int32(16)
	v3175 = int32(24)
	v3178 = v3167 ^ (v3156<<(uint(v3159)%32)&int32(_a_F_createTrgmNFA_10) | v3157<<(uint(v3170)%32)&int32(16711680) | v3158<<(uint(v3175)%32))
	v3185 = int32(1)
	v3188 = v3153 - v3185
	if v3188 != 0 {
		v3152 = v3152 + v3185
		v3153 = v3188
		v3155 = int32(base.Ui32(v3178) >> (uint(v3175) % 32))
		v3156 = v3167
		v3157 = int32(base.Ui32(v3178) >> (uint(v3159) % 32))
		v3158 = int32(base.Ui32(v3178) >> (uint(v3170) % 32))
		goto L384
	} else {
		goto L386
	}
L385:
	;
	v3193 = v3178 ^ int32(-1)
	goto L381
L386:
	;
	goto L385
L387:
	;
	goto L347
L388:
	;
	v3219 = int32(1)
	goto L387
L389:
	;
	goto L390
L390:
	;
	v3216 = v3009 + int32(1)
	v3217 = *(*int32)(unsafe.Add(mBase, uint32(v2885)+4))
	if v3216 < v3217 {
		v3009 = v3216
		v3015 = v3212
		goto L346
	} else {
		goto L391
	}
L391:
	;
	v3219 = v3217
	goto L387
L392:
	;
	v3221 = *(*int32)(unsafe.Add(mBase, uint32(v2875)+4))
	v3222 = v3219
	v3226 = v3221
	v3229 = v3212
	goto L342
L393:
	;
	v3268 = v3226
	v3271 = v3229
	goto L339
L394:
	;
	v3303 = *(*int32)(unsafe.Add(mBase, uint32(v2861)+4))
	v3304 = v3303
	v3308 = v3268
	v3311 = v3271
	goto L330
L395:
	;
	v3353 = v3311
	goto L327
L396:
	;
	goto L315
L397:
	;
	v3472 = int32(2)
	v3473 = F_hash_seq_search(m, v3469)
	mBase = m.M
	v3474 = m.ExcPending
	if v3474 != 0 {
		goto L1
	} else {
		goto L398
	}
L398:
	;
	if v3473 != 0 {
		goto L399
	} else {
		goto L400
	}
L399:
	;
	v3479 = v3473
	v3486 = v3472
	goto L402
L400:
	;
	v3547 = v3472
	goto L401
L401:
	;
	v3576 = *(*int32)(unsafe.Add(mBase, uint32(v42)+96))
	v3577 = F_palloc_mul(m, int32(12), v3576)
	mBase = m.M
	v3578 = m.ExcPending
	if v3578 != 0 {
		goto L1
	} else {
		goto L415
	}
L402:
	;
	v3514 = *(*int32)(unsafe.Add(mBase, uint32(v3479)+28))
	if v3514 != 0 {
		v3479 = v3514
		goto L402
	} else {
		goto L404
	}
L403:
	;
	v3547 = v3531
	goto L401
L404:
	;
	v3515 = *(*int32)(unsafe.Add(mBase, uint32(v3479)+24))
	if int32(0) <= v3515 {
		v3531 = v3486
		goto L405
	} else {
		goto L406
	}
L405:
	;
	v3534 = F_hash_seq_search(m, v42+int32(176))
	mBase = m.M
	v3535 = m.ExcPending
	if v3535 != 0 {
		goto L1
	} else {
		goto L413
	}
L406:
	;
	v3518 = *(*int32)(unsafe.Add(mBase, uint32(v3479)+20))
	if v3518&int32(1) != 0 {
		goto L407
	} else {
		goto L408
	}
L407:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3479)+24)) = int32(0)
	v3531 = v3486
	goto L405
L408:
	;
	goto L409
L409:
	;
	if v3518&int32(2) != 0 {
		goto L410
	} else {
		goto L411
	}
L410:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3479)+24)) = int32(1)
	v3531 = v3486
	goto L405
L411:
	;
	goto L412
L412:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3479)+24)) = v3486
	v3531 = v3486 + int32(1)
	goto L405
L413:
	;
	if v3534 != 0 {
		v3479 = v3534
		v3486 = v3531
		goto L402
	} else {
		goto L414
	}
L414:
	;
	goto L403
L415:
	;
	v3580 = v42 + int32(176)
	F_hash_seq_init(m, v3580, v3443)
	mBase = m.M
	v3582 = m.ExcPending
	if v3582 != 0 {
		goto L1
	} else {
		goto L416
	}
L416:
	;
	v3583 = F_hash_seq_search(m, v3580)
	mBase = m.M
	v3584 = m.ExcPending
	if v3584 != 0 {
		goto L1
	} else {
		goto L418
	}
L417:
	;
	v3952 = int32(0)
	v3954 = F_MemoryContextAlloc(m, l3, int32(28))
	mBase = m.M
	v3955 = m.ExcPending
	if v3955 != 0 {
		goto L1
	} else {
		goto L455
	}
L418:
	;
	if v3583 != 0 {
		goto L419
	} else {
		goto L420
	}
L419:
	;
	v3586 = *(*int32)(unsafe.Add(mBase, uint32(v42)+108))
	v3587 = *(*int32)(unsafe.Add(mBase, uint32(v42)+104))
	v3588 = v3583
	v3589 = int32(0)
	goto L422
L420:
	;
	goto L421
L421:
	;
	v3907 = int32(0)
	F_pg_qsort(m, v3577, v3907, int32(12), int32(_a_F_createTrgmNFA_11))
	mBase = m.M
	v3912 = m.ExcPending
	if v3912 != 0 {
		goto L1
	} else {
		goto L454
	}
L422:
	;
	v3631 = v3588
	goto L424
L423:
	;
	F_pg_qsort(m, v3577, v3785, int32(12), int32(_a_F_createTrgmNFA_11))
	mBase = m.M
	v3830 = m.ExcPending
	if v3830 != 0 {
		goto L1
	} else {
		goto L442
	}
L424:
	;
	v3666 = *(*int32)(unsafe.Add(mBase, uint32(v3631)+28))
	if v3666 != 0 {
		v3631 = v3666
		goto L424
	} else {
		goto L426
	}
L425:
	;
	v3667 = *(*int32)(unsafe.Add(mBase, uint32(v3588)+12))
	if v3667 == int32(0) {
		v3785 = v3589
		goto L427
	} else {
		goto L428
	}
L426:
	;
	goto L425
L427:
	;
	v3825 = F_hash_seq_search(m, v42+int32(176))
	mBase = m.M
	v3826 = m.ExcPending
	if v3826 != 0 {
		goto L1
	} else {
		goto L440
	}
L428:
	;
	v3670 = int32(0)
	v3671 = *(*int32)(unsafe.Add(mBase, uint32(v3667)+4))
	if v3671 <= v3670 {
		v3785 = v3589
		goto L427
	} else {
		goto L429
	}
L429:
	;
	v3675 = v3589
	v3681 = v3670
	v3683 = v3671
	goto L430
L430:
	;
	v3713 = *(*int32)(unsafe.Add(mBase, uint32(v3667)+12))
	v3717 = *(*int32)(unsafe.Add(mBase, uint32(v3713+v3681<<(uint(int32(2))%32))))
	v3718 = *(*int32)(unsafe.Add(mBase, uint32(v3717)+12))
	v3723 = v3718
	goto L432
L431:
	;
	v3785 = v3778
	goto L427
L432:
	;
	v3758 = *(*int32)(unsafe.Add(mBase, uint32(v3723)+28))
	if v3758 != 0 {
		v3723 = v3758
		goto L432
	} else {
		goto L434
	}
L433:
	;
	v3759 = *(*int32)(unsafe.Add(mBase, uint32(v3631)+24))
	v3760 = *(*int32)(unsafe.Add(mBase, uint32(v3723)+24))
	if v3759 != v3760 {
		goto L435
	} else {
		goto L436
	}
L434:
	;
	goto L433
L435:
	;
	v3764 = F_bsearch(m, v3717, v3587, v3586, int32(32), int32(_a_F_createTrgmNFA_5))
	mBase = m.M
	v3765 = m.ExcPending
	if v3765 != 0 {
		goto L1
	} else {
		goto L438
	}
L436:
	;
	v3778 = v3675
	v3780 = v3683
	goto L437
L437:
	;
	v3782 = v3681 + int32(1)
	if v3782 < v3780 {
		v3675 = v3778
		v3681 = v3782
		v3683 = v3780
		goto L430
	} else {
		goto L439
	}
L438:
	;
	v3768 = v3577 + v3675*int32(12)
	v3769 = *(*int32)(unsafe.Add(mBase, uint32(v3631)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v3768))) = v3769
	v3771 = *(*int32)(unsafe.Add(mBase, uint32(v3723)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v3768)+4)) = v3771
	v3773 = *(*int32)(unsafe.Add(mBase, uint32(v3764)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v3768)+8)) = v3773
	v3775 = *(*int32)(unsafe.Add(mBase, uint32(v3667)+4))
	v3778 = v3675 + int32(1)
	v3780 = v3775
	goto L437
L439:
	;
	goto L431
L440:
	;
	if v3825 != 0 {
		v3588 = v3825
		v3589 = v3785
		goto L422
	} else {
		goto L441
	}
L441:
	;
	goto L423
L442:
	;
	if v3785 < int32(2) {
		v3914 = v3785
		goto L417
	} else {
		goto L443
	}
L443:
	;
	v3833 = int32(12)
	__phi3838 = v3577
	__phi3842 = v3577
	__phi3845 = v3577 + v3833
	v3838 = __phi3838
	v3842 = __phi3842
	v3845 = __phi3845
	goto L444
L444:
	;
	v3877 = *(*int32)(unsafe.Add(mBase, uint32(v3842)+12))
	v3878 = *(*int32)(unsafe.Add(mBase, uint32(v3838)))
	if v3877 < v3878 {
		v3896 = v3838
		goto L446
	} else {
		goto L447
	}
L445:
	;
	v3904 = base.I32_div_s(v3896-v3577, int32(12))
	v3914 = v3904 + int32(1)
	goto L417
L446:
	;
	v3900 = v3845 + int32(12)
	if base.Ui32(v3900) < base.Ui32(v3577+v3785*v3833) {
		__phi3838 = v3896
		__phi3842 = v3845
		__phi3845 = v3900
		v3838 = __phi3838
		v3842 = __phi3842
		v3845 = __phi3845
		goto L444
	} else {
		goto L453
	}
L447:
	;
	if v3878 < v3877 {
		goto L448
	} else {
		goto L449
	}
L448:
	;
	v3890 = *(*int32)(unsafe.Add(mBase, uint32(v3845)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v3838)+20)) = v3890
	v3892 = *(*int64)(unsafe.Add(mBase, uint32(v3845)))
	*(*int64)(unsafe.Add(mBase, uint32(v3838)+12)) = v3892
	v3896 = v3838 + int32(12)
	goto L446
L449:
	;
	v3881 = *(*int32)(unsafe.Add(mBase, uint32(v3842)+20))
	v3882 = *(*int32)(unsafe.Add(mBase, uint32(v3838)+8))
	if v3881 < v3882 {
		v3896 = v3838
		goto L446
	} else {
		goto L450
	}
L450:
	;
	if v3882 < v3881 {
		goto L448
	} else {
		goto L451
	}
L451:
	;
	v3885 = *(*int32)(unsafe.Add(mBase, uint32(v3842)+16))
	v3886 = *(*int32)(unsafe.Add(mBase, uint32(v3838)+4))
	if v3885 <= v3886 {
		v3896 = v3838
		goto L446
	} else {
		goto L452
	}
L452:
	;
	goto L448
L453:
	;
	goto L445
L454:
	;
	v3914 = v3907
	goto L417
L455:
	;
	v3956 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v3954))) = v3956
	v3958 = *(*int32)(unsafe.Add(mBase, uint32(v42)+108))
	if v3958 <= v3956 {
		goto L457
	} else {
		goto L458
	}
L456:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3954)+8)) = v3547
	v4361 = F_MemoryContextAlloc(m, l3, v3547<<(uint(int32(3))%32))
	mBase = m.M
	v4362 = m.ExcPending
	if v4362 != 0 {
		goto L1
	} else {
		goto L502
	}
L457:
	;
	v3962 = F_MemoryContextAlloc(m, l3, int32(0))
	mBase = m.M
	v3963 = m.ExcPending
	if v3963 != 0 {
		goto L1
	} else {
		goto L460
	}
L458:
	;
	goto L459
L459:
	;
	v3966 = v3958 & int32(3)
	v3967 = *(*int32)(unsafe.Add(mBase, uint32(v42)+104))
	v3968 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v3958) {
		goto L462
	} else {
		goto L463
	}
L460:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3954)+4)) = v3962
	goto L456
L461:
	;
	v4186 = F_MemoryContextAlloc(m, l3, v4145<<(uint(int32(2))%32))
	mBase = m.M
	v4187 = m.ExcPending
	if v4187 != 0 {
		goto L1
	} else {
		goto L487
	}
L462:
	;
	v3974 = v3952
	v3978 = v3968
	v3983 = int32(0)
	goto L465
L463:
	;
	v4051 = v3952
	v4055 = v3968
	goto L464
L464:
	;
	v4091 = v4051
	v4095 = v4055
	v4097 = int32(0)
	goto L481
L465:
	;
	v4015 = v3967 + v3978<<(uint(int32(5))%32)
	v4016 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4015)+24)))
	if v4016 == int32(1) {
		goto L467
	} else {
		goto L468
	}
L466:
	;
	if v3966 == int32(0) {
		v4145 = v4043
		goto L461
	} else {
		goto L480
	}
L467:
	;
	v4020 = v3974 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v3954))) = v4020
	v4022 = v4020
	goto L469
L468:
	;
	v4022 = v3974
	goto L469
L469:
	;
	v4023 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4015)+56)))
	if v4023 == int32(1) {
		goto L470
	} else {
		goto L471
	}
L470:
	;
	v4027 = v4022 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v3954))) = v4027
	v4029 = v4027
	goto L472
L471:
	;
	v4029 = v4022
	goto L472
L472:
	;
	v4030 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4015)+88)))
	if v4030 == int32(1) {
		goto L473
	} else {
		goto L474
	}
L473:
	;
	v4034 = v4029 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v3954))) = v4034
	v4036 = v4034
	goto L475
L474:
	;
	v4036 = v4029
	goto L475
L475:
	;
	v4037 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4015)+120)))
	if v4037 == int32(1) {
		goto L476
	} else {
		goto L477
	}
L476:
	;
	v4041 = v4036 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v3954))) = v4041
	v4043 = v4041
	goto L478
L477:
	;
	v4043 = v4036
	goto L478
L478:
	;
	v4044 = int32(4)
	v4045 = v3978 + v4044
	v4047 = v3983 + v4044
	if v4047 != v3958&int32(2147483644) {
		v3974 = v4043
		v3978 = v4045
		v3983 = v4047
		goto L465
	} else {
		goto L479
	}
L479:
	;
	goto L466
L480:
	;
	v4051 = v4043
	v4055 = v4045
	goto L464
L481:
	;
	v4133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3967+v4095<<(uint(int32(5))%32))+24)))
	if v4133 == int32(1) {
		goto L483
	} else {
		goto L484
	}
L482:
	;
	v4145 = v4139
	goto L461
L483:
	;
	v4137 = v4091 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v3954))) = v4137
	v4139 = v4137
	goto L485
L484:
	;
	v4139 = v4091
	goto L485
L485:
	;
	v4140 = int32(1)
	v4143 = v4097 + v4140
	if v4143 != v3966 {
		v4091 = v4139
		v4095 = v4095 + v4140
		v4097 = v4143
		goto L481
	} else {
		goto L486
	}
L486:
	;
	goto L482
L487:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3954)+4)) = v4186
	v4189 = int32(0)
	if v3958 != int32(1) {
		goto L488
	} else {
		goto L489
	}
L488:
	;
	v4198 = v4189
	v4202 = v4189
	v4207 = int32(0)
	goto L491
L489:
	;
	v4269 = v4189
	v4273 = v4189
	goto L490
L490:
	;
	v4310 = v3967 + v4273<<(uint(int32(5))%32)
	v4311 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4310)+24)))
	if v4311 != int32(1) {
		goto L456
	} else {
		goto L501
	}
L491:
	;
	v4239 = v3967 + v4202<<(uint(int32(5))%32)
	v4240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4239)+24)))
	if v4240 == int32(1) {
		goto L493
	} else {
		goto L494
	}
L492:
	;
	if v3958&int32(1) == int32(0) {
		goto L456
	} else {
		goto L500
	}
L493:
	;
	v4246 = *(*int32)(unsafe.Add(mBase, uint32(v4239)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v4186+v4198<<(uint(int32(2))%32)))) = v4246
	v4250 = v4198 + int32(1)
	goto L495
L494:
	;
	v4250 = v4198
	goto L495
L495:
	;
	v4251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4239)+56)))
	if v4251 == int32(1) {
		goto L496
	} else {
		goto L497
	}
L496:
	;
	v4257 = *(*int32)(unsafe.Add(mBase, uint32(v4239)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v4186+v4250<<(uint(int32(2))%32)))) = v4257
	v4261 = v4250 + int32(1)
	goto L498
L497:
	;
	v4261 = v4250
	goto L498
L498:
	;
	v4262 = int32(2)
	v4263 = v4202 + v4262
	v4265 = v4207 + v4262
	if v4265 != v3958&int32(2147483646) {
		v4198 = v4261
		v4202 = v4263
		v4207 = v4265
		goto L491
	} else {
		goto L499
	}
L499:
	;
	goto L492
L500:
	;
	v4269 = v4261
	v4273 = v4263
	goto L490
L501:
	;
	v4317 = *(*int32)(unsafe.Add(mBase, uint32(v4310)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v4186+v4269<<(uint(int32(2))%32)))) = v4317
	goto L456
L502:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3954)+12)) = v4361
	v4366 = F_MemoryContextAlloc(m, l3, v3914<<(uint(int32(3))%32))
	mBase = m.M
	v4367 = m.ExcPending
	if v4367 != 0 {
		goto L1
	} else {
		goto L503
	}
L503:
	;
	if int32(0) < v3547 {
		goto L504
	} else {
		goto L505
	}
L504:
	;
	v4370 = *(*int32)(unsafe.Add(mBase, uint32(v3954)+12))
	v4371 = int32(0)
	v4377 = v4371
	v4382 = v4371
	goto L507
L505:
	;
	goto L506
L506:
	;
	v4561 = *(*int32)(unsafe.Add(mBase, uint32(v3954)))
	v4562 = F_MemoryContextAlloc(m, l3, v4561)
	mBase = m.M
	v4563 = m.ExcPending
	if v4563 != 0 {
		goto L1
	} else {
		goto L518
	}
L507:
	;
	v4412 = int32(3)
	v4414 = v4370 + v4382<<(uint(v4412)%32)
	*(*int32)(unsafe.Add(mBase, uint32(v4414)+4)) = v4366 + v4377<<(uint(v4412)%32)
	if v3914 <= v4377 {
		goto L510
	} else {
		goto L511
	}
L508:
	;
	goto L506
L509:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4414))) = v4485
	v4520 = v4382 + int32(1)
	if v4520 != v3547 {
		v4377 = v4483
		v4382 = v4520
		goto L507
	} else {
		goto L517
	}
L510:
	;
	v4483 = v4377
	v4485 = int32(0)
	goto L509
L511:
	;
	goto L512
L512:
	;
	v4421 = v3914 - v4377
	v4427 = v4377
	v4429 = int32(0)
	goto L513
L513:
	;
	v4464 = v3577 + v4427*int32(12)
	v4465 = *(*int32)(unsafe.Add(mBase, uint32(v4464)))
	if v4465 != v4382 {
		v4483 = v4427
		v4485 = v4429
		goto L509
	} else {
		goto L515
	}
L514:
	;
	v4483 = v3914
	v4485 = v4421
	goto L509
L515:
	;
	v4469 = v4366 + v4427<<(uint(int32(3))%32)
	v4470 = *(*int32)(unsafe.Add(mBase, uint32(v4464)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v4469))) = v4470
	v4472 = *(*int32)(unsafe.Add(mBase, uint32(v4464)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v4469)+4)) = v4472
	v4474 = int32(1)
	v4477 = v4429 + v4474
	if v4477 != v4421 {
		v4427 = v4427 + v4474
		v4429 = v4477
		goto L513
	} else {
		goto L516
	}
L516:
	;
	goto L514
L517:
	;
	goto L508
L518:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3954)+16)) = v4562
	v4565 = *(*int32)(unsafe.Add(mBase, uint32(v3954)+8))
	v4566 = F_MemoryContextAlloc(m, l3, v4565)
	mBase = m.M
	v4567 = m.ExcPending
	if v4567 != 0 {
		goto L1
	} else {
		goto L519
	}
L519:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3954)+20)) = v4566
	v4569 = *(*int32)(unsafe.Add(mBase, uint32(v3954)+8))
	v4572 = F_MemoryContextAlloc(m, l3, v4569<<(uint(int32(2))%32))
	mBase = m.M
	v4573 = m.ExcPending
	if v4573 != 0 {
		goto L1
	} else {
		goto L520
	}
L520:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3954)+24)) = v4572
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v3954
	v4595 = v2790
	goto L21
L521:
	;
	m.G0 = v42 + int32(240)
	return v4595
}
func F_create_edata_for_relation(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = int32(0)
	v14 = F_palloc0(m, int32(20))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = l0
		v19 = F_CreateExecutorState(m)
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v14))) = v19
			v23 = F_palloc0(m, int32(136))
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v23)+12)) = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v23))) = int32(101)
				v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
				v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+56))
				*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = v30
				v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
				v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+48))
				v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+119)))
				*(*int32)(unsafe.Add(mBase, uint32(v23)+24)) = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(v23)+21)) = uint8(v34)
				v40 = F_addRTEPermissionInfo(m, v9+int32(12), v23)
				mBase = m.M
				v41 = m.ExcPending
				if v41 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v23
					*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v23
					v47 = F_list_make1_impl(m, int32(1), v9+int32(4))
					mBase = m.M
					v48 = m.ExcPending
					if v48 != 0 {
						return int32(0)
					} else {
						v49 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
						v51 = F_bms_make_singleton(m, int32(1))
						mBase = m.M
						v52 = m.ExcPending
						if v52 != 0 {
							return int32(0)
						} else {
							F_ExecInitRangeTable(m, v19, v47, v49, v51)
							mBase = m.M
							v54 = m.ExcPending
							if v54 != 0 {
								return int32(0)
							} else {
								v56 = F_palloc0(m, int32(216))
								mBase = m.M
								v57 = m.ExcPending
								if v57 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v56))) = int32(394)
									*(*int32)(unsafe.Add(mBase, uint32(v14)+8)) = v56
									v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
									v63 = int32(0)
									F_InitResultRelInfo(m, v56, v61, int32(1), v63, v63)
									mBase = m.M
									v66 = m.ExcPending
									if v66 != 0 {
										return int32(0)
									} else {
										v67 = *(*int32)(unsafe.Add(mBase, uint32(v19)+72))
										v68 = F_lappend(m, v67, v56)
										mBase = m.M
										v69 = m.ExcPending
										if v69 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v19)+72)) = v68
											v72 = F_GetCurrentCommandId(m, int32(1))
											mBase = m.M
											v73 = m.ExcPending
											if v73 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v19)+64)) = v72
												v75 = int32(_a_F_create_edata_for_relation_0)
												v77 = *(*int32)(unsafe.Add(mBase, _c_F_create_edata_for_relation[0]))
												*(*int32)(unsafe.Add(mBase, _c_F_create_edata_for_relation[0])) = v77 + int32(1)
												m.G0 = v9 + int32(16)
												return v14
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
func F_create_seqscan_path(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	v8 = F_palloc0(m, int32(72))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = l1
		*(*int64)(unsafe.Add(mBase, uint32(v8))) = int64(1473173782810)
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
		*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v15
		v17 = F_get_baserel_parampathinfo(m, l0, l1, l2)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			v19 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v8)+20)) = uint8(base.B2i32(v19 < l3))
			*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v17
			v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+26)))
			*(*int32)(unsafe.Add(mBase, uint32(v8)+64)) = v19
			*(*int32)(unsafe.Add(mBase, uint32(v8)+24)) = l3
			*(*uint8)(unsafe.Add(mBase, uint32(v8)+21)) = uint8(v23)
			F_cost_seqscan(m, v8, l0, l1, v17)
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return int32(0)
			} else {
				return v8
			}
		}
	}
}
func F_create_tidscan_path(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v10 float64
	_ = v10
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v49 float64
	_ = v49
	var v51 int32
	_ = v51
	var v60 int32
	_ = v60
	var v65 float64
	_ = v65
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 float64
	_ = v83
	var v84 int32
	_ = v84
	var v86 float64
	_ = v86
	var v87 float64
	_ = v87
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v100 int32
	_ = v100
	var v105 float64
	_ = v105
	var v111 int64
	_ = v111
	var v126 int32
	_ = v126
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v148 float64
	_ = v148
	var v161 float64
	_ = v161
	var v177 float64
	_ = v177
	var v178 float64
	_ = v178
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v187 float64
	_ = v187
	var v188 int32
	_ = v188
	var v189 int64
	_ = v189
	var v198 int32
	_ = v198
	var v204 int32
	_ = v204
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v232 float64
	_ = v232
	var v233 float64
	_ = v233
	var v244 float64
	_ = v244
	var v251 float64
	_ = v251
	var v252 float64
	_ = v252
	var v254 float64
	_ = v254
	var v256 float64
	_ = v256
	var v257 float64
	_ = v257
	var v268 float64
	_ = v268
	var v275 float64
	_ = v275
	var v276 int32
	_ = v276
	var v277 float64
	_ = v277
	var v279 float64
	_ = v279
	var v280 int64
	_ = v280
	var v284 float64
	_ = v284
	var v285 float64
	_ = v285
	var v289 int32
	_ = v289
	var v290 int64
	_ = v290
	var v295 float64
	_ = v295
	v5 = int32(0)
	v10 = float64(0)
	v19 = F_palloc0(m, int32(80))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+8)) = l1
	*(*int64)(unsafe.Add(mBase, uint32(v19))) = int64(1498943586592)
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+12)) = v26
	v28 = F_get_baserel_parampathinfo(m, l0, l1, l3)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v30 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+20)) = uint8(v30)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = v28
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+26)))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+72)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v19)+64)) = v30
	*(*int32)(unsafe.Add(mBase, uint32(v19)+24)) = v30
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+21)) = uint8(v33)
	v40 = m.G0
	v42 = v40 - int32(32)
	m.G0 = v42
	if v28 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v48 = v28 + int32(8)
	goto L6
L5:
	;
	v48 = l1 + int32(16)
	goto L6
L6:
	;
	v49 = *(*float64)(unsafe.Add(mBase, uint32(v48)))
	*(*float64)(unsafe.Add(mBase, uint32(v19)+32)) = v49
	if l2 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if int32(0) < v51 {
		goto L11
	} else {
		goto L12
	}
L8:
	;
	v177 = v10
	v178 = v10
	goto L9
L9:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
	F_get_tablespace_page_costs(m, v183, v42, int32(0))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L1
	} else {
		goto L26
	}
L10:
	;
	v177 = v105
	v178 = v161
	goto L9
L11:
	;
	v60 = v5
	v65 = v10
	goto L14
L12:
	;
	v100 = v5
	v105 = v10
	goto L13
L13:
	;
	v111 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v42)+16)) = v111
	*(*int32)(unsafe.Add(mBase, uint32(v42)+8)) = l0
	*(*int64)(unsafe.Add(mBase, uint32(v42)+24)) = v111
	if v100 == int32(0) {
		v161 = v10
		goto L10
	} else {
		goto L21
	}
L14:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v71+v60<<(uint(int32(2))%32))))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v75)+4))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)))
	if v77 == int32(20) {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v100 = base.B2i32(int32(0) < v90)
	v105 = v87
	goto L13
L16:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v76)+28))
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v80)+12))
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v81)+4))
	v83 = F_estimate_array_length(m, l0, v82)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L1
	} else {
		goto L19
	}
L17:
	;
	v86 = float64(1)
	goto L18
L18:
	;
	v87 = base.F64_add(v65, v86)
	v89 = v60 + int32(1)
	v90 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v89 < v90 {
		v60 = v89
		v65 = v87
		goto L14
	} else {
		goto L20
	}
L19:
	;
	v86 = v83
	goto L18
L20:
	;
	goto L15
L21:
	;
	v126 = v5
	goto L22
L22:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v135+v126<<(uint(int32(2))%32))))
	v142 = F_cost_qual_eval_walker(m, v139, v42+int32(8))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L1
	} else {
		goto L24
	}
L23:
	;
	v148 = *(*float64)(unsafe.Add(mBase, uint32(v42)+24))
	v161 = v148
	goto L10
L24:
	;
	v145 = v126 + int32(1)
	v146 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v145 < v146 {
		v126 = v145
		goto L22
	} else {
		goto L25
	}
L25:
	;
	goto L23
L26:
	;
	v187 = *(*float64)(unsafe.Add(mBase, uint32(v42)))
	if v28 != 0 {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
	v277 = *(*float64)(unsafe.Add(mBase, uint32(v276)+24))
	v279 = *(*float64)(unsafe.Add(mBase, _c_F_create_tidscan_path[0]))
	v280 = *(*int64)(unsafe.Add(mBase, uint32(l1)+32))
	v284 = *(*float64)(unsafe.Add(mBase, uint32(v276)+16))
	v285 = base.F64_add(base.F64_add(base.F64_add(v178, v275), float64(0)), v284)
	*(*float64)(unsafe.Add(mBase, uint32(v19)+48)) = v285
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v19)+24))
	if v289 != 0 {
		goto L38
	} else {
		goto L39
	}
L28:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v28)+16))
	v189 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v42)+16)) = v189
	*(*int32)(unsafe.Add(mBase, uint32(v42)+8)) = l0
	*(*int64)(unsafe.Add(mBase, uint32(v42)+24)) = v189
	if v188 == int32(0) {
		v244 = v10
		v251 = float64(0)
		goto L31
	} else {
		goto L32
	}
L29:
	;
	goto L30
L30:
	;
	v256 = *(*float64)(unsafe.Add(mBase, uint32(l1)+216))
	v257 = *(*float64)(unsafe.Add(mBase, uint32(l1)+208))
	v268 = v256
	v275 = v257
	goto L27
L31:
	;
	v252 = *(*float64)(unsafe.Add(mBase, uint32(l1)+216))
	v254 = *(*float64)(unsafe.Add(mBase, uint32(l1)+208))
	v268 = base.F64_add(v244, v252)
	v275 = base.F64_add(v251, v254)
	goto L27
L32:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v188)+4))
	if v198 <= int32(0) {
		v244 = v10
		v251 = float64(0)
		goto L31
	} else {
		goto L33
	}
L33:
	;
	v204 = int32(0)
	goto L34
L34:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v188)+12))
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v219+v204<<(uint(int32(2))%32))))
	v226 = F_cost_qual_eval_walker(m, v223, v42+int32(8))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L1
	} else {
		goto L36
	}
L35:
	;
	v232 = *(*float64)(unsafe.Add(mBase, uint32(v42)+24))
	v233 = *(*float64)(unsafe.Add(mBase, uint32(v42)+16))
	v244 = v232
	v251 = v233
	goto L31
L36:
	;
	v229 = v204 + int32(1)
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v188)+4))
	if v229 < v230 {
		v204 = v229
		goto L34
	} else {
		goto L37
	}
L37:
	;
	goto L35
L38:
	;
	v290 = int64(-1)
	goto L40
L39:
	;
	v290 = int64(-262145)
	goto L40
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+40)) = base.B2i32(v280|v290 != int64(-1))
	v295 = *(*float64)(unsafe.Add(mBase, uint32(v19)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v19)+56)) = base.F64_add(v285, base.F64_add(base.F64_mul(v277, v295), base.F64_add(base.F64_mul(base.F64_sub(base.F64_add(v268, v279), v178), v177), base.F64_add(base.F64_mul(v187, v177), float64(0)))))
	m.G0 = v42 + int32(32)
	return v19
}
func F_createdb_failure_callback(m *base.Module, l0 int32, l1 int64) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	v3 = base.I32_wrap_i64(l1)
	v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)+8))
	if v4 == int32(0) {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(v3)+4))
		F_DropDatabaseBuffers(m, v7)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return
		} else {
			v10 = *(*int32)(unsafe.Add(mBase, uint32(v3)+4))
			F_ForgetDatabaseSyncRequests(m, v10)
			mBase = m.M
			v12 = m.ExcPending
			if v12 != 0 {
				return
			} else {
				v14 = *(*int32)(unsafe.Add(mBase, uint32(v3)+4))
				F_UnlockSharedObject(m, int32(1262), v14, int32(1))
				mBase = m.M
				v17 = m.ExcPending
				if v17 != 0 {
					return
				} else {
					v19 = *(*int32)(unsafe.Add(mBase, uint32(v3)))
					F_UnlockSharedObject(m, int32(1262), v19, int32(5))
					mBase = m.M
					v22 = m.ExcPending
					if v22 != 0 {
						return
					} else {
						v23 = *(*int32)(unsafe.Add(mBase, uint32(v3)+4))
						F_remove_dbtablespaces(m, v23)
						mBase = m.M
						v25 = m.ExcPending
						if v25 != 0 {
							return
						} else {
							return
						}
					}
				}
			}
		}
	} else {
		v19 = *(*int32)(unsafe.Add(mBase, uint32(v3)))
		F_UnlockSharedObject(m, int32(1262), v19, int32(5))
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return
		} else {
			v23 = *(*int32)(unsafe.Add(mBase, uint32(v3)+4))
			F_remove_dbtablespaces(m, v23)
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return
			} else {
				return
			}
		}
	}
}
func F_cryptohash_internal(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v89 int32
	_ = v89
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	v10 = m.G0
	v12 = v10 - int32(48)
	m.G0 = v12
	v15 = l0 << (uint(int32(2)) % 32)
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v15)+uint32(_c_F_cryptohash_internal[0])))
	v20 = v18 + int32(4)
	v21 = F_palloc0(m, v20)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		return int32(0)
	} else {
		v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
		if v27 == int32(1) {
			v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
			if v33 == int32(18) {
				v36 = int32(16)
			} else {
				v36 = int32(0)
			}
			if base.Ui32((v33-int32(1))&int32(255)) < base.Ui32(int32(3)) {
				v43 = int32(4)
			} else {
				v43 = v36
			}
			v56 = v43
		} else {
			v44 = int32(1)
			if v27&v44 != 0 {
				v56 = int32(base.Ui32(v27)>>(uint(v44)%32)) - v44
			} else {
				v50 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				v56 = int32(base.Ui32(v50)>>(uint(int32(2))%32)) - int32(4)
			}
		}
		v57 = *(*int32)(unsafe.Add(mBase, uint32(v15)+uint32(_c_F_cryptohash_internal[1])))
		v58 = F_pg_cryptohash_create(m, l0)
		mBase = m.M
		v59 = m.ExcPending
		if v59 != 0 {
			return int32(0)
		} else {
			v60 = F_pg_cryptohash_init(m, v58)
			mBase = m.M
			if int32(0) <= v60 {
				v63 = int32(1)
				if v27&v63 != 0 {
					v67 = v63
				} else {
					v67 = int32(4)
				}
				v69 = F_pg_cryptohash_update(m, v58, l1+v67, v56)
				mBase = m.M
				if v69 < int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v118 = m.ExcPending
					if v118 != 0 {
						return int32(0)
					} else {
						if v58 == int32(0) {
							v133 = int32(_a_F_cryptohash_internal_0)
						} else {
							v125 = *(*int32)(unsafe.Add(mBase, uint32(v58)+4))
							if v125 == int32(1) {
								v128 = int32(_a_F_cryptohash_internal_1)
							} else {
								v128 = int32(_a_F_cryptohash_internal_2)
							}
							if v125 == int32(2) {
								v131 = int32(_a_F_cryptohash_internal_0)
							} else {
								v131 = v128
							}
							v133 = v131
						}
						*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v133
						*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v57
						F_errmsg_internal(m, int32(_a_F_cryptohash_internal_3), v12+int32(16))
						mBase = m.M
						v140 = m.ExcPending
						if v140 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_cryptohash_internal_4), int32(123), int32(_a_F_cryptohash_internal_5))
							mBase = m.M
							v145 = m.ExcPending
							if v145 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v74 = F_pg_cryptohash_final(m, v58, v21+int32(4), v18)
					mBase = m.M
					if v74 < int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v149 = m.ExcPending
						if v149 != 0 {
							return int32(0)
						} else {
							if v58 == int32(0) {
								v164 = int32(_a_F_cryptohash_internal_0)
							} else {
								v156 = *(*int32)(unsafe.Add(mBase, uint32(v58)+4))
								if v156 == int32(1) {
									v159 = int32(_a_F_cryptohash_internal_1)
								} else {
									v159 = int32(_a_F_cryptohash_internal_2)
								}
								if v156 == int32(2) {
									v162 = int32(_a_F_cryptohash_internal_0)
								} else {
									v162 = v159
								}
								v164 = v162
							}
							*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = v164
							*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v57
							F_errmsg_internal(m, int32(_a_F_cryptohash_internal_6), v12+int32(32))
							mBase = m.M
							v171 = m.ExcPending
							if v171 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_cryptohash_internal_4), int32(127), int32(_a_F_cryptohash_internal_5))
								mBase = m.M
								v176 = m.ExcPending
								if v176 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						F_pg_cryptohash_free(m, v58)
						mBase = m.M
						v78 = m.ExcPending
						if v78 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v21))) = v20 << (uint(int32(2)) % 32)
							m.G0 = v12 + int32(48)
							return v21
						}
					}
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v89 = m.ExcPending
				if v89 != 0 {
					return int32(0)
				} else {
					if v58 == int32(0) {
						v104 = int32(_a_F_cryptohash_internal_0)
					} else {
						v96 = *(*int32)(unsafe.Add(mBase, uint32(v58)+4))
						if v96 == int32(1) {
							v99 = int32(_a_F_cryptohash_internal_1)
						} else {
							v99 = int32(_a_F_cryptohash_internal_2)
						}
						if v96 == int32(2) {
							v102 = int32(_a_F_cryptohash_internal_0)
						} else {
							v102 = v99
						}
						v104 = v102
					}
					*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v104
					*(*int32)(unsafe.Add(mBase, uint32(v12))) = v57
					F_errmsg_internal(m, int32(_a_F_cryptohash_internal_7), v12)
					mBase = m.M
					v109 = m.ExcPending
					if v109 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_cryptohash_internal_4), int32(120), int32(_a_F_cryptohash_internal_5))
						mBase = m.M
						v114 = m.ExcPending
						if v114 != 0 {
							return int32(0)
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
