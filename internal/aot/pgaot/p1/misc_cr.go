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
func F_crc32_bytea(m *base.Module, l0 int32) int32 {
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
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = F_pg_detoast_datum_packed(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		v11 = int32(1)
		v12 = v7 + v11
		v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
		v15 = v13 & v11
		if v13 == v11 {
			v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
			if base.Ui32((v19-int32(1))&int32(255)) < base.Ui32(int32(3)) {
				v46 = int32(4)
				if v15 != 0 {
					v50 = v12
				} else {
					v50 = v7 + int32(4)
				}
				v51 = int32(1)
				if v46 == v51 {
					v94 = v50
					v95 = int32(-1)
				} else {
					v60 = v50
					v61 = int32(-1)
					v62 = int32(0)
					for {
						v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60))))
						v67 = int32(255)
						v69 = int32(2)
						v73 = *(*int32)(unsafe.Add(mBase, uint32((v65^v61)&v67<<(uint(v69)%32))+uint32(_c_F_crc32_bytea[0])))
						v74 = int32(8)
						v76 = v73 ^ int32(base.Ui32(v61)>>(uint(v74)%32))
						v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+1)))
						v85 = *(*int32)(unsafe.Add(mBase, uint32((v76^v77)&v67<<(uint(v69)%32))+uint32(_c_F_crc32_bytea[0])))
						v88 = v85 ^ int32(base.Ui32(v76)>>(uint(v74)%32))
						v90 = v60 + v69
						v92 = v62 + v69
						if v92 != v46&int32(-2) {
							v60 = v90
							v61 = v88
							v62 = v92
							continue
						} else {
							break
						}
						break
					}
					v94 = v90
					v95 = v88
				}
				if v46&v51 != 0 {
					v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94))))
					v107 = *(*int32)(unsafe.Add(mBase, uint32((v99^v95)&int32(255)<<(uint(int32(2))%32))+uint32(_c_F_crc32_bytea[0])))
					v111 = v107 ^ int32(base.Ui32(v95)>>(uint(int32(8))%32))
				} else {
					v111 = v95
				}
				v115 = F_Int64GetDatum(m, base.I64_extend_i32_u(v111^int32(-1)))
				mBase = m.M
				v116 = m.ExcPending
				if v116 != 0 {
					return int32(0)
				} else {
					return v115
				}
			} else {
				v41 = base.B2i32(v19 == int32(18)) << (uint(int32(4)) % 32)
				if v41 != 0 {
					v46 = v41
					if v15 != 0 {
						v50 = v12
					} else {
						v50 = v7 + int32(4)
					}
					v51 = int32(1)
					if v46 == v51 {
						v94 = v50
						v95 = int32(-1)
					} else {
						v60 = v50
						v61 = int32(-1)
						v62 = int32(0)
						for {
							v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60))))
							v67 = int32(255)
							v69 = int32(2)
							v73 = *(*int32)(unsafe.Add(mBase, uint32((v65^v61)&v67<<(uint(v69)%32))+uint32(_c_F_crc32_bytea[0])))
							v74 = int32(8)
							v76 = v73 ^ int32(base.Ui32(v61)>>(uint(v74)%32))
							v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+1)))
							v85 = *(*int32)(unsafe.Add(mBase, uint32((v76^v77)&v67<<(uint(v69)%32))+uint32(_c_F_crc32_bytea[0])))
							v88 = v85 ^ int32(base.Ui32(v76)>>(uint(v74)%32))
							v90 = v60 + v69
							v92 = v62 + v69
							if v92 != v46&int32(-2) {
								v60 = v90
								v61 = v88
								v62 = v92
								continue
							} else {
								break
							}
							break
						}
						v94 = v90
						v95 = v88
					}
					if v46&v51 != 0 {
						v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94))))
						v107 = *(*int32)(unsafe.Add(mBase, uint32((v99^v95)&int32(255)<<(uint(int32(2))%32))+uint32(_c_F_crc32_bytea[0])))
						v111 = v107 ^ int32(base.Ui32(v95)>>(uint(int32(8))%32))
					} else {
						v111 = v95
					}
					v115 = F_Int64GetDatum(m, base.I64_extend_i32_u(v111^int32(-1)))
					mBase = m.M
					v116 = m.ExcPending
					if v116 != 0 {
						return int32(0)
					} else {
						return v115
					}
				} else {
					v43 = F_Int64GetDatum(m, int64(0))
					mBase = m.M
					v44 = m.ExcPending
					if v44 != 0 {
						return int32(0)
					} else {
						return v43
					}
				}
			}
		} else {
			v30 = int32(1)
			if v15 != 0 {
				v41 = int32(base.Ui32(v13)>>(uint(v30)%32)) - v30
			} else {
				v34 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
				v41 = int32(base.Ui32(v34)>>(uint(int32(2))%32)) - int32(4)
			}
			if v41 != 0 {
				v46 = v41
				if v15 != 0 {
					v50 = v12
				} else {
					v50 = v7 + int32(4)
				}
				v51 = int32(1)
				if v46 == v51 {
					v94 = v50
					v95 = int32(-1)
				} else {
					v60 = v50
					v61 = int32(-1)
					v62 = int32(0)
					for {
						v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60))))
						v67 = int32(255)
						v69 = int32(2)
						v73 = *(*int32)(unsafe.Add(mBase, uint32((v65^v61)&v67<<(uint(v69)%32))+uint32(_c_F_crc32_bytea[0])))
						v74 = int32(8)
						v76 = v73 ^ int32(base.Ui32(v61)>>(uint(v74)%32))
						v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+1)))
						v85 = *(*int32)(unsafe.Add(mBase, uint32((v76^v77)&v67<<(uint(v69)%32))+uint32(_c_F_crc32_bytea[0])))
						v88 = v85 ^ int32(base.Ui32(v76)>>(uint(v74)%32))
						v90 = v60 + v69
						v92 = v62 + v69
						if v92 != v46&int32(-2) {
							v60 = v90
							v61 = v88
							v62 = v92
							continue
						} else {
							break
						}
						break
					}
					v94 = v90
					v95 = v88
				}
				if v46&v51 != 0 {
					v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94))))
					v107 = *(*int32)(unsafe.Add(mBase, uint32((v99^v95)&int32(255)<<(uint(int32(2))%32))+uint32(_c_F_crc32_bytea[0])))
					v111 = v107 ^ int32(base.Ui32(v95)>>(uint(int32(8))%32))
				} else {
					v111 = v95
				}
				v115 = F_Int64GetDatum(m, base.I64_extend_i32_u(v111^int32(-1)))
				mBase = m.M
				v116 = m.ExcPending
				if v116 != 0 {
					return int32(0)
				} else {
					return v115
				}
			} else {
				v43 = F_Int64GetDatum(m, int64(0))
				mBase = m.M
				v44 = m.ExcPending
				if v44 != 0 {
					return int32(0)
				} else {
					return v43
				}
			}
		}
	}
}
func F_createTrgmNFA(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v73 int32
	_ = v73
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
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v140 int32
	_ = v140
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v261 int32
	_ = v261
	var v265 int32
	_ = v265
	var v269 int32
	_ = v269
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
	var v282 int32
	_ = v282
	var v286 int32
	_ = v286
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v339 int32
	_ = v339
	var v376 int32
	_ = v376
	var v380 int32
	_ = v380
	var v385 int32
	_ = v385
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v397 int32
	_ = v397
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v413 int32
	_ = v413
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v437 int32
	_ = v437
	var v441 int32
	_ = v441
	var v443 int32
	_ = v443
	var v449 int32
	_ = v449
	var v490 int32
	_ = v490
	var v530 int32
	_ = v530
	var v570 int64
	_ = v570
	var v577 int32
	_ = v577
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v606 int32
	_ = v606
	var v611 int32
	_ = v611
	var v613 int64
	_ = v613
	var v622 int32
	_ = v622
	var v623 int32
	_ = v623
	var v625 int32
	_ = v625
	var v632 int32
	_ = v632
	var v633 int32
	_ = v633
	var v649 int32
	_ = v649
	var v676 int32
	_ = v676
	var v680 int32
	_ = v680
	var v681 int32
	_ = v681
	var v684 int32
	_ = v684
	var v688 int32
	_ = v688
	var v694 int32
	_ = v694
	var v695 int32
	_ = v695
	var v698 int32
	_ = v698
	var v705 int32
	_ = v705
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
	var v798 int32
	_ = v798
	var v801 int32
	_ = v801
	var v804 int32
	_ = v804
	var v808 int32
	_ = v808
	var v811 int32
	_ = v811
	var v812 int32
	_ = v812
	var v825 int32
	_ = v825
	var v853 int32
	_ = v853
	var v854 int32
	_ = v854
	var v858 int32
	_ = v858
	var v859 int32
	_ = v859
	var v860 int32
	_ = v860
	var v861 int32
	_ = v861
	var v864 int32
	_ = v864
	var v865 int32
	_ = v865
	var v866 int32
	_ = v866
	var v868 int32
	_ = v868
	var v876 int32
	_ = v876
	var v912 int32
	_ = v912
	var v913 int32
	_ = v913
	var v916 int32
	_ = v916
	var v919 int32
	_ = v919
	var v920 int32
	_ = v920
	var v923 int32
	_ = v923
	var v926 int32
	_ = v926
	var v927 int32
	_ = v927
	var v930 int32
	_ = v930
	var v938 int32
	_ = v938
	var v940 int32
	_ = v940
	var v943 int32
	_ = v943
	var v945 int32
	_ = v945
	var v947 int32
	_ = v947
	var v954 int32
	_ = v954
	var v958 int32
	_ = v958
	var v999 int32
	_ = v999
	var v1001 int32
	_ = v1001
	var v1002 int32
	_ = v1002
	var v1042 int32
	_ = v1042
	var v1045 int32
	_ = v1045
	var v1047 int32
	_ = v1047
	var v1048 int32
	_ = v1048
	var v1049 int32
	_ = v1049
	var v1050 int32
	_ = v1050
	var v1051 int32
	_ = v1051
	var v1052 int32
	_ = v1052
	var v1053 int32
	_ = v1053
	var v1054 int32
	_ = v1054
	var v1055 int32
	_ = v1055
	var v1056 int32
	_ = v1056
	var v1057 int32
	_ = v1057
	var v1058 int32
	_ = v1058
	var v1059 int32
	_ = v1059
	var v1060 int32
	_ = v1060
	var v1061 int32
	_ = v1061
	var v1062 int32
	_ = v1062
	var v1063 int32
	_ = v1063
	var v1064 int32
	_ = v1064
	var v1065 int32
	_ = v1065
	var v1066 int32
	_ = v1066
	var v1067 int32
	_ = v1067
	var v1068 int32
	_ = v1068
	var v1069 int32
	_ = v1069
	var v1070 int32
	_ = v1070
	var v1071 int32
	_ = v1071
	var v1072 int32
	_ = v1072
	var v1073 int32
	_ = v1073
	var v1074 int32
	_ = v1074
	var v1075 int32
	_ = v1075
	var v1078 int32
	_ = v1078
	var v1079 int32
	_ = v1079
	var v1080 int32
	_ = v1080
	var v1081 int32
	_ = v1081
	var v1082 int32
	_ = v1082
	var v1114 int32
	_ = v1114
	var v1117 int32
	_ = v1117
	var v1120 int32
	_ = v1120
	var v1121 int32
	_ = v1121
	var v1123 int32
	_ = v1123
	var v1128 int32
	_ = v1128
	var v1137 int32
	_ = v1137
	var v1162 int32
	_ = v1162
	var v1166 int32
	_ = v1166
	var v1167 int32
	_ = v1167
	var v1168 int32
	_ = v1168
	var v1172 int32
	_ = v1172
	var v1174 int32
	_ = v1174
	var v1177 int32
	_ = v1177
	var v1178 int32
	_ = v1178
	var v1181 int32
	_ = v1181
	var v1192 int32
	_ = v1192
	var v1218 int32
	_ = v1218
	var v1221 int32
	_ = v1221
	var v1222 int32
	_ = v1222
	var v1226 int32
	_ = v1226
	var v1232 int32
	_ = v1232
	var v1263 int32
	_ = v1263
	var v1267 int32
	_ = v1267
	var v1269 int32
	_ = v1269
	var v1270 int32
	_ = v1270
	var v1272 int32
	_ = v1272
	var v1276 int32
	_ = v1276
	var v1277 int32
	_ = v1277
	var v1279 int64
	_ = v1279
	var v1281 int32
	_ = v1281
	var v1290 int32
	_ = v1290
	var v1291 int32
	_ = v1291
	var v1293 int32
	_ = v1293
	var v1294 int32
	_ = v1294
	var v1296 int32
	_ = v1296
	var v1297 int32
	_ = v1297
	var v1300 int32
	_ = v1300
	var v1339 int32
	_ = v1339
	var v1340 int32
	_ = v1340
	var v1384 int32
	_ = v1384
	var v1388 int32
	_ = v1388
	var __phi1388 int32
	_ = __phi1388
	var v1392 int32
	_ = v1392
	var __phi1392 int32
	_ = __phi1392
	var v1394 int32
	_ = v1394
	var __phi1394 int32
	_ = __phi1394
	var v1426 int32
	_ = v1426
	var v1434 int32
	_ = v1434
	var v1435 int32
	_ = v1435
	var v1436 int32
	_ = v1436
	var v1439 int32
	_ = v1439
	var v1440 int32
	_ = v1440
	var v1442 int32
	_ = v1442
	var v1443 int32
	_ = v1443
	var v1445 int32
	_ = v1445
	var v1447 int32
	_ = v1447
	var v1457 int32
	_ = v1457
	var v1458 int32
	_ = v1458
	var v1459 int32
	_ = v1459
	var v1462 int32
	_ = v1462
	var v1463 int32
	_ = v1463
	var v1464 int32
	_ = v1464
	var v1467 int32
	_ = v1467
	var v1468 int32
	_ = v1468
	var v1470 int32
	_ = v1470
	var v1475 int32
	_ = v1475
	var v1488 int32
	_ = v1488
	var v1491 int64
	_ = v1491
	var v1493 int64
	_ = v1493
	var v1495 int64
	_ = v1495
	var v1497 int64
	_ = v1497
	var v1501 int32
	_ = v1501
	var v1502 int32
	_ = v1502
	var v1503 int32
	_ = v1503
	var v1504 int32
	_ = v1504
	var v1506 int32
	_ = v1506
	var v1508 int32
	_ = v1508
	var v1528 int32
	_ = v1528
	var v1557 int32
	_ = v1557
	var v1565 int32
	_ = v1565
	var v1593 int64
	_ = v1593
	var v1594 float32
	_ = v1594
	var v1600 int32
	_ = v1600
	var v1601 int32
	_ = v1601
	var v1608 int32
	_ = v1608
	var v1609 int32
	_ = v1609
	var v1610 int32
	_ = v1610
	var v1611 int32
	_ = v1611
	var v1617 int32
	_ = v1617
	var v1621 int32
	_ = v1621
	var v1622 int32
	_ = v1622
	var v1624 int32
	_ = v1624
	var v1625 int32
	_ = v1625
	var v1631 int32
	_ = v1631
	var v1635 int32
	_ = v1635
	var v1636 int32
	_ = v1636
	var v1642 float32
	_ = v1642
	var v1644 float32
	_ = v1644
	var v1646 float32
	_ = v1646
	var v1648 int64
	_ = v1648
	var v1650 int32
	_ = v1650
	var v1655 int32
	_ = v1655
	var v1673 int32
	_ = v1673
	var v1692 int64
	_ = v1692
	var v1693 float32
	_ = v1693
	var v1701 int32
	_ = v1701
	var v1702 int32
	_ = v1702
	var v1705 int32
	_ = v1705
	var v1708 int32
	_ = v1708
	var v1720 int32
	_ = v1720
	var v1722 int32
	_ = v1722
	var v1751 int32
	_ = v1751
	var v1752 int32
	_ = v1752
	var v1753 int32
	_ = v1753
	var v1754 int32
	_ = v1754
	var v1792 int32
	_ = v1792
	var v1799 int32
	_ = v1799
	var v1831 int32
	_ = v1831
	var v1832 int32
	_ = v1832
	var v1833 int32
	_ = v1833
	var v1839 int32
	_ = v1839
	var v1842 int32
	_ = v1842
	var v1843 int32
	_ = v1843
	var v1873 int32
	_ = v1873
	var v1874 int32
	_ = v1874
	var v1876 int32
	_ = v1876
	var v1878 int32
	_ = v1878
	var v1879 int32
	_ = v1879
	var v1881 int32
	_ = v1881
	var v1887 int32
	_ = v1887
	var v1919 int32
	_ = v1919
	var v1920 int32
	_ = v1920
	var v1922 int32
	_ = v1922
	var v1925 int32
	_ = v1925
	var v1928 int32
	_ = v1928
	var v1935 int32
	_ = v1935
	var v1936 int32
	_ = v1936
	var v1938 int32
	_ = v1938
	var v1941 int32
	_ = v1941
	var v1942 int32
	_ = v1942
	var v1952 int32
	_ = v1952
	var v1986 int32
	_ = v1986
	var v1987 int32
	_ = v1987
	var v1988 int32
	_ = v1988
	var v1989 int32
	_ = v1989
	var v2027 int32
	_ = v2027
	var v2034 int32
	_ = v2034
	var v2066 int32
	_ = v2066
	var v2071 int32
	_ = v2071
	var v2107 int32
	_ = v2107
	var v2146 int32
	_ = v2146
	var v2147 int32
	_ = v2147
	var __phi2147 int32
	_ = __phi2147
	var v2151 int32
	_ = v2151
	var __phi2151 int32
	_ = __phi2151
	var v2187 int32
	_ = v2187
	var v2227 int32
	_ = v2227
	var v2228 int32
	_ = v2228
	var v2242 int32
	_ = v2242
	var v2268 int32
	_ = v2268
	var v2278 int32
	_ = v2278
	var v2283 int32
	_ = v2283
	var v2312 int32
	_ = v2312
	var v2313 int32
	_ = v2313
	var v2314 int32
	_ = v2314
	var v2315 int32
	_ = v2315
	var v2353 int32
	_ = v2353
	var v2358 int32
	_ = v2358
	var v2392 int32
	_ = v2392
	var v2394 int32
	_ = v2394
	var v2395 int32
	_ = v2395
	var v2399 int32
	_ = v2399
	var v2400 int32
	_ = v2400
	var v2402 int32
	_ = v2402
	var v2442 int32
	_ = v2442
	var v2444 float32
	_ = v2444
	var v2446 int64
	_ = v2446
	var v2483 int64
	_ = v2483
	var v2484 float32
	_ = v2484
	var v2487 int32
	_ = v2487
	var v2524 int64
	_ = v2524
	var v2529 int32
	_ = v2529
	var v2534 int32
	_ = v2534
	var v2536 int32
	_ = v2536
	var v2537 int32
	_ = v2537
	var v2545 int32
	_ = v2545
	var v2546 int32
	_ = v2546
	var v2549 int32
	_ = v2549
	var v2585 int32
	_ = v2585
	var v2586 int32
	_ = v2586
	var v2592 int32
	_ = v2592
	var v2593 int32
	_ = v2593
	var v2599 int32
	_ = v2599
	var v2600 int32
	_ = v2600
	var v2608 int32
	_ = v2608
	var v2609 int32
	_ = v2609
	var v2615 int32
	_ = v2615
	var v2616 int32
	_ = v2616
	var v2617 int32
	_ = v2617
	var v2619 int32
	_ = v2619
	var v2621 int32
	_ = v2621
	var v2625 int32
	_ = v2625
	var v2661 int32
	_ = v2661
	var v2665 int32
	_ = v2665
	var v2668 int32
	_ = v2668
	var v2701 int32
	_ = v2701
	var v2702 int32
	_ = v2702
	var v2708 int32
	_ = v2708
	var v2709 int32
	_ = v2709
	var v2712 int32
	_ = v2712
	var v2717 int32
	_ = v2717
	var v2718 int32
	_ = v2718
	var v2724 int32
	_ = v2724
	var v2727 int32
	_ = v2727
	var v2728 int32
	_ = v2728
	var v2732 int32
	_ = v2732
	var v2735 int32
	_ = v2735
	var v2743 int32
	_ = v2743
	var v2748 int32
	_ = v2748
	var v2759 int32
	_ = v2759
	var v2793 int32
	_ = v2793
	var v2794 int32
	_ = v2794
	var v2800 int32
	_ = v2800
	var v2813 int32
	_ = v2813
	var v2821 int32
	_ = v2821
	var v2827 int32
	_ = v2827
	var v2838 int32
	_ = v2838
	var v2852 int32
	_ = v2852
	var v2855 int32
	_ = v2855
	var v2856 int32
	_ = v2856
	var v2859 int32
	_ = v2859
	var v2860 int32
	_ = v2860
	var v2863 int32
	_ = v2863
	var v2864 int32
	_ = v2864
	var v2867 int32
	_ = v2867
	var v2870 int32
	_ = v2870
	var v2873 int32
	_ = v2873
	var v2874 int32
	_ = v2874
	var v2875 int32
	_ = v2875
	var v2877 int32
	_ = v2877
	var v2878 int32
	_ = v2878
	var v2882 int32
	_ = v2882
	var v2885 int32
	_ = v2885
	var v2886 int32
	_ = v2886
	var v2887 int32
	_ = v2887
	var v2889 int32
	_ = v2889
	var v2893 int32
	_ = v2893
	var v2896 int32
	_ = v2896
	var v2910 int32
	_ = v2910
	var v2931 int32
	_ = v2931
	var v2932 int32
	_ = v2932
	var v2936 int32
	_ = v2936
	var v2939 int32
	_ = v2939
	var v2942 int32
	_ = v2942
	var v2943 int32
	_ = v2943
	var v2945 int32
	_ = v2945
	var v2947 int32
	_ = v2947
	var v2949 int32
	_ = v2949
	var v2953 int32
	_ = v2953
	var v2957 int32
	_ = v2957
	var v2960 int32
	_ = v2960
	var v2975 int32
	_ = v2975
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
	var v3048 int32
	_ = v3048
	var v3049 int32
	_ = v3049
	var v3053 int32
	_ = v3053
	var v3068 int32
	_ = v3068
	var v3070 int32
	_ = v3070
	var v3072 int32
	_ = v3072
	var v3095 int32
	_ = v3095
	var v3097 int32
	_ = v3097
	var v3099 int32
	_ = v3099
	var v3104 int32
	_ = v3104
	var v3111 int32
	_ = v3111
	var v3124 int32
	_ = v3124
	var v3128 int32
	_ = v3128
	var v3130 int32
	_ = v3130
	var v3133 int32
	_ = v3133
	var v3136 int32
	_ = v3136
	var v3142 int32
	_ = v3142
	var v3146 int32
	_ = v3146
	var v3148 int32
	_ = v3148
	var v3150 int32
	_ = v3150
	var v3153 int32
	_ = v3153
	var v3154 int32
	_ = v3154
	var v3156 int32
	_ = v3156
	var v3157 int32
	_ = v3157
	var v3158 int32
	_ = v3158
	var v3159 int32
	_ = v3159
	var v3160 int32
	_ = v3160
	var v3162 int32
	_ = v3162
	var v3168 int32
	_ = v3168
	var v3171 int32
	_ = v3171
	var v3176 int32
	_ = v3176
	var v3179 int32
	_ = v3179
	var v3186 int32
	_ = v3186
	var v3189 int32
	_ = v3189
	var v3194 int32
	_ = v3194
	var v3202 int32
	_ = v3202
	var v3213 int32
	_ = v3213
	var v3217 int32
	_ = v3217
	var v3218 int32
	_ = v3218
	var v3220 int32
	_ = v3220
	var v3222 int32
	_ = v3222
	var v3223 int32
	_ = v3223
	var v3227 int32
	_ = v3227
	var v3230 int32
	_ = v3230
	var v3262 int32
	_ = v3262
	var v3268 int32
	_ = v3268
	var v3271 int32
	_ = v3271
	var v3302 int32
	_ = v3302
	var v3303 int32
	_ = v3303
	var v3307 int32
	_ = v3307
	var v3310 int32
	_ = v3310
	var v3342 int32
	_ = v3342
	var v3351 int32
	_ = v3351
	var v3382 int32
	_ = v3382
	var v3390 int32
	_ = v3390
	var v3396 int32
	_ = v3396
	var v3422 int32
	_ = v3422
	var v3424 int32
	_ = v3424
	var v3442 int32
	_ = v3442
	var v3466 int32
	_ = v3466
	var v3467 int32
	_ = v3467
	var v3470 int32
	_ = v3470
	var v3471 int32
	_ = v3471
	var v3476 int32
	_ = v3476
	var v3485 int32
	_ = v3485
	var v3510 int32
	_ = v3510
	var v3511 int32
	_ = v3511
	var v3514 int32
	_ = v3514
	var v3527 int32
	_ = v3527
	var v3530 int32
	_ = v3530
	var v3531 int32
	_ = v3531
	var v3545 int32
	_ = v3545
	var v3570 int32
	_ = v3570
	var v3573 int32
	_ = v3573
	var v3574 int32
	_ = v3574
	var v3578 int32
	_ = v3578
	var v3581 int32
	_ = v3581
	var v3582 int32
	_ = v3582
	var v3584 int32
	_ = v3584
	var v3585 int32
	_ = v3585
	var v3586 int32
	_ = v3586
	var v3587 int32
	_ = v3587
	var v3628 int32
	_ = v3628
	var v3662 int32
	_ = v3662
	var v3663 int32
	_ = v3663
	var v3666 int32
	_ = v3666
	var v3667 int32
	_ = v3667
	var v3671 int32
	_ = v3671
	var v3677 int32
	_ = v3677
	var v3679 int32
	_ = v3679
	var v3708 int32
	_ = v3708
	var v3712 int32
	_ = v3712
	var v3713 int32
	_ = v3713
	var v3718 int32
	_ = v3718
	var v3752 int32
	_ = v3752
	var v3753 int32
	_ = v3753
	var v3754 int32
	_ = v3754
	var v3758 int32
	_ = v3758
	var v3759 int32
	_ = v3759
	var v3762 int32
	_ = v3762
	var v3763 int32
	_ = v3763
	var v3765 int32
	_ = v3765
	var v3767 int32
	_ = v3767
	var v3769 int32
	_ = v3769
	var v3772 int32
	_ = v3772
	var v3774 int32
	_ = v3774
	var v3776 int32
	_ = v3776
	var v3779 int32
	_ = v3779
	var v3818 int32
	_ = v3818
	var v3819 int32
	_ = v3819
	var v3823 int32
	_ = v3823
	var v3826 int32
	_ = v3826
	var v3831 int32
	_ = v3831
	var __phi3831 int32
	_ = __phi3831
	var v3835 int32
	_ = v3835
	var __phi3835 int32
	_ = __phi3835
	var v3838 int32
	_ = v3838
	var __phi3838 int32
	_ = __phi3838
	var v3869 int32
	_ = v3869
	var v3870 int32
	_ = v3870
	var v3873 int32
	_ = v3873
	var v3874 int32
	_ = v3874
	var v3877 int32
	_ = v3877
	var v3878 int32
	_ = v3878
	var v3882 int64
	_ = v3882
	var v3884 int32
	_ = v3884
	var v3888 int32
	_ = v3888
	var v3892 int32
	_ = v3892
	var v3896 int32
	_ = v3896
	var v3899 int32
	_ = v3899
	var v3904 int32
	_ = v3904
	var v3906 int32
	_ = v3906
	var v3943 int32
	_ = v3943
	var v3945 int32
	_ = v3945
	var v3946 int32
	_ = v3946
	var v3947 int32
	_ = v3947
	var v3949 int32
	_ = v3949
	var v3953 int32
	_ = v3953
	var v3954 int32
	_ = v3954
	var v3957 int32
	_ = v3957
	var v3958 int32
	_ = v3958
	var v3959 int32
	_ = v3959
	var v3965 int32
	_ = v3965
	var v3969 int32
	_ = v3969
	var v3973 int32
	_ = v3973
	var v4005 int32
	_ = v4005
	var v4006 int32
	_ = v4006
	var v4010 int32
	_ = v4010
	var v4012 int32
	_ = v4012
	var v4013 int32
	_ = v4013
	var v4017 int32
	_ = v4017
	var v4019 int32
	_ = v4019
	var v4020 int32
	_ = v4020
	var v4024 int32
	_ = v4024
	var v4026 int32
	_ = v4026
	var v4027 int32
	_ = v4027
	var v4031 int32
	_ = v4031
	var v4033 int32
	_ = v4033
	var v4034 int32
	_ = v4034
	var v4035 int32
	_ = v4035
	var v4037 int32
	_ = v4037
	var v4039 int32
	_ = v4039
	var v4043 int32
	_ = v4043
	var v4078 int32
	_ = v4078
	var v4082 int32
	_ = v4082
	var v4084 int32
	_ = v4084
	var v4119 int32
	_ = v4119
	var v4123 int32
	_ = v4123
	var v4125 int32
	_ = v4125
	var v4126 int32
	_ = v4126
	var v4129 int32
	_ = v4129
	var v4131 int32
	_ = v4131
	var v4171 int32
	_ = v4171
	var v4172 int32
	_ = v4172
	var v4174 int32
	_ = v4174
	var v4176 int32
	_ = v4176
	var v4183 int32
	_ = v4183
	var v4187 int32
	_ = v4187
	var v4191 int32
	_ = v4191
	var v4223 int32
	_ = v4223
	var v4224 int32
	_ = v4224
	var v4230 int32
	_ = v4230
	var v4234 int32
	_ = v4234
	var v4235 int32
	_ = v4235
	var v4241 int32
	_ = v4241
	var v4245 int32
	_ = v4245
	var v4246 int32
	_ = v4246
	var v4247 int32
	_ = v4247
	var v4249 int32
	_ = v4249
	var v4251 int32
	_ = v4251
	var v4255 int32
	_ = v4255
	var v4293 int32
	_ = v4293
	var v4294 int32
	_ = v4294
	var v4300 int32
	_ = v4300
	var v4343 int32
	_ = v4343
	var v4344 int32
	_ = v4344
	var v4348 int32
	_ = v4348
	var v4349 int32
	_ = v4349
	var v4352 int32
	_ = v4352
	var v4353 int32
	_ = v4353
	var v4359 int32
	_ = v4359
	var v4363 int32
	_ = v4363
	var v4393 int32
	_ = v4393
	var v4395 int32
	_ = v4395
	var v4402 int32
	_ = v4402
	var v4408 int32
	_ = v4408
	var v4410 int32
	_ = v4410
	var v4444 int32
	_ = v4444
	var v4445 int32
	_ = v4445
	var v4449 int32
	_ = v4449
	var v4450 int32
	_ = v4450
	var v4452 int32
	_ = v4452
	var v4454 int32
	_ = v4454
	var v4457 int32
	_ = v4457
	var v4463 int32
	_ = v4463
	var v4465 int32
	_ = v4465
	var v4499 int32
	_ = v4499
	var v4539 int32
	_ = v4539
	var v4540 int32
	_ = v4540
	var v4541 int32
	_ = v4541
	var v4543 int32
	_ = v4543
	var v4544 int32
	_ = v4544
	var v4545 int32
	_ = v4545
	var v4547 int32
	_ = v4547
	var v4550 int32
	_ = v4550
	var v4551 int32
	_ = v4551
	var v4572 int32
	_ = v4572
	var v4595 int32
	_ = v4595
	v5 = int32(0)
	v39 = m.G0
	v41 = v39 - int32(240)
	m.G0 = v41
	v44 = *(*int32)(unsafe.Add(mBase, _c_F_createTrgmNFA[0]))
	v49 = F_AllocSetContextCreateInternal(m, v44, int32(_a_F_createTrgmNFA_0), v5, int32(_a_F_createTrgmNFA_1), int32(_a_F_createTrgmNFA_2))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v53 = int32(_a_F_createTrgmNFA_3)
	v54 = *(*int32)(unsafe.Add(mBase, _c_F_createTrgmNFA[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_createTrgmNFA[0])) = v49
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v59 == int32(1) {
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
	v62 = int32(4)
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
	if v64&int32(254) == int32(2) {
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
	if v59&v77 != 0 {
		v89 = int32(base.Ui32(v59)>>(uint(v77)%32)) - v77
		goto L3
	} else {
		goto L13
	}
L7:
	;
	v73 = v62
	goto L9
L8:
	;
	v73 = base.B2i32(v64 == int32(18)) << (uint(v62) % 32)
	goto L9
L9:
	;
	if v64 == int32(1) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v76 = v62
	goto L12
L11:
	;
	v76 = v73
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
	if v59&v96 != 0 {
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
	v105 = F_pg_regcomp(m, v41+int32(32), v94, v102, int32(27), l1)
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
	*(*int32)(unsafe.Add(mBase, _c_F_createTrgmNFA[0])) = v54
	F_MemoryContextDelete(m, v49)
	mBase = m.M
	v4595 = m.ExcPending
	if v4595 != 0 {
		goto L1
	} else {
		goto L538
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+224)) = int32(0)
	v2793 = F_MemoryContextAllocZero(m, l3, v2759*int32(3)+int32(5))
	mBase = m.M
	v2794 = m.ExcPending
	if v2794 != 0 {
		goto L1
	} else {
		goto L326
	}
L23:
	;
	v112 = v41 + int32(32)
	*(*int32)(unsafe.Add(mBase, uint32(v41)+64)) = v112
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v112)+24))
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v116)+84))
	v119 = v117 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v41)+72)) = v119
	v123 = F_palloc0(m, v119*int32(12))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L1
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	v2727 = F_pg_regerror(m, v105, v41-int32(-64))
	mBase = m.M
	v2728 = m.ExcPending
	if v2728 != 0 {
		goto L1
	} else {
		goto L321
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+68)) = v123
	if int32(0) < v119 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v140 = v5
	goto L30
L28:
	;
	goto L29
L29:
	;
	v570 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v41)+93)) = v570
	*(*int64)(unsafe.Add(mBase, uint32(v41)+88)) = v570
	*(*int64)(unsafe.Add(mBase, uint32(v41)+192)) = int64(171798691852)
	v577 = *(*int32)(unsafe.Add(mBase, _c_F_createTrgmNFA[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v41)+216)) = v577
	v584 = F_hash_create(m, int32(_a_F_createTrgmNFA_4), int32(1024), v41+int32(176), int32(1064))
	mBase = m.M
	v585 = m.ExcPending
	if v585 != 0 {
		goto L1
	} else {
		goto L86
	}
L30:
	;
	v170 = v123 + v140*int32(12)
	v173 = int32(-1)
	if v140 <= int32(0) {
		v189 = v173
		goto L32
	} else {
		goto L33
	}
L31:
	;
	goto L29
L32:
	;
	if base.Ui32(int32(257)) <= base.Ui32(v189) {
		goto L38
	} else {
		goto L39
	}
L33:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v41+int32(32))+24))
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v176)+84))
	if base.Ui32(v177) < base.Ui32(v140) {
		v189 = v173
		goto L32
	} else {
		goto L34
	}
L34:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v176)+92))
	v182 = v179 + v140*int32(24)
	v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182)+20)))
	if v183&int32(2) != 0 {
		v189 = v173
		goto L32
	} else {
		goto L35
	}
L35:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v182)+4))
	if v186 != 0 {
		v189 = v173
		goto L32
	} else {
		goto L36
	}
L36:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v182)))
	v189 = v187
	goto L32
L37:
	;
	v530 = v140 + int32(1)
	if v530 != v119 {
		v140 = v530
		goto L30
	} else {
		goto L85
	}
L38:
	;
	v192 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v170))) = uint8(v192)
	goto L37
L39:
	;
	goto L40
L40:
	;
	v194 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v170))) = uint16(v194)
	v198 = v189 << (uint(int32(2)) % 32)
	v199 = F_palloc(m, v198)
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v170)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v170)+8)) = v199
	v206 = F_palloc(m, v198)
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	if v140 <= int32(0) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	if v189 != 0 {
		goto L59
	} else {
		goto L60
	}
L44:
	;
	if v189 <= int32(0) {
		goto L43
	} else {
		goto L45
	}
L45:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v41+int32(32))+24))
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v212)+84))
	if base.Ui32(v213) < base.Ui32(v140) {
		goto L43
	} else {
		goto L46
	}
L46:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v212)+92))
	v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v215+v140*int32(24))+20)))
	if v219&int32(2) != 0 {
		goto L43
	} else {
		goto L47
	}
L47:
	;
	v223 = v189
	v230 = int32(0)
	v231 = v206
	goto L48
L48:
	;
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v212)+96))
	v265 = int32(*(*int16)(unsafe.Add(mBase, uint32(v261+v230<<(uint(int32(1))%32)))))
	if v265 == v140 {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	goto L43
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v231))) = v230
	v269 = v223 - int32(1)
	if v269 == int32(0) {
		goto L43
	} else {
		goto L53
	}
L51:
	;
	v274 = v223
	v275 = v231
	goto L52
L52:
	;
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v212)+96))
	v277 = int32(1)
	v278 = v230 | v277
	v282 = int32(*(*int16)(unsafe.Add(mBase, uint32(v276+v278<<(uint(v277)%32)))))
	if v282 == v140 {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	v274 = v269
	v275 = v231 + int32(4)
	goto L52
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v275))) = v278
	v286 = v274 - int32(1)
	if v286 == int32(0) {
		goto L43
	} else {
		goto L57
	}
L55:
	;
	v291 = v274
	v292 = v275
	goto L56
L56:
	;
	v294 = v230 + int32(2)
	if v294 != int32(2048) {
		v223 = v291
		v230 = v294
		v231 = v292
		goto L48
	} else {
		goto L58
	}
L57:
	;
	v291 = v286
	v292 = v275 + int32(4)
	goto L56
L58:
	;
	goto L49
L59:
	;
	v339 = int32(0)
	goto L62
L60:
	;
	goto L61
L61:
	;
	F_pfree(m, v206)
	mBase = m.M
	v490 = m.ExcPending
	if v490 != 0 {
		goto L1
	} else {
		goto L84
	}
L62:
	;
	v376 = *(*int32)(unsafe.Add(mBase, uint32(v206+v339<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v41)+224)) = v376
	if v376 == int32(0) {
		goto L64
	} else {
		goto L65
	}
L63:
	;
	goto L61
L64:
	;
	v449 = v339 + int32(1)
	if v449 != v189 {
		v339 = v449
		goto L62
	} else {
		goto L83
	}
L65:
	;
	v380 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v41+int32(180)))) = uint8(v380)
	*(*int32)(unsafe.Add(mBase, uint32(v41)+176)) = v380
	v385 = v41 + int32(176)
	v391 = F_pg_wchar2mb_with_len(m, v41+int32(224), v385, int32(1))
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	v394 = F_str_tolower(m, v385, v391, int32(100))
	mBase = m.M
	v395 = m.ExcPending
	if v395 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	v397 = v41 + int32(176)
	v400 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v397))))
	v401 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v394))))
	if v401 == int32(0) {
		v420 = v400
		v421 = v401
		goto L69
	} else {
		goto L70
	}
L68:
	;
	F_pfree(m, v394)
	mBase = m.M
	v424 = m.ExcPending
	if v424 != 0 {
		goto L1
	} else {
		goto L76
	}
L69:
	;
	goto L68
L70:
	;
	if v400 != v401 {
		v420 = v400
		v421 = v401
		goto L69
	} else {
		goto L71
	}
L71:
	;
	v405 = v394
	v406 = v397
	goto L72
L72:
	;
	v409 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v406)+1)))
	v410 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v405)+1)))
	if v410 == int32(0) {
		v420 = v409
		v421 = v410
		goto L69
	} else {
		goto L74
	}
L73:
	;
	v420 = v409
	v421 = v410
	goto L69
L74:
	;
	v413 = int32(1)
	if v409 == v410 {
		v405 = v405 + v413
		v406 = v406 + v413
		goto L72
	} else {
		goto L75
	}
L75:
	;
	goto L73
L76:
	;
	if v421-v420 != 0 {
		goto L64
	} else {
		goto L77
	}
L77:
	;
	v425 = *(*int32)(unsafe.Add(mBase, uint32(v41)+176))
	*(*int32)(unsafe.Add(mBase, uint32(v41)+164)) = v425
	if v391 == int32(0) {
		goto L64
	} else {
		goto L78
	}
L78:
	;
	v431 = F_t_isalnum_with_len(m, v41+int32(164), v391)
	mBase = m.M
	v432 = m.ExcPending
	if v432 != 0 {
		goto L1
	} else {
		goto L79
	}
L79:
	;
	if v431 != 0 {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v433 = *(*int32)(unsafe.Add(mBase, uint32(v170)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v170)+4)) = v433 + int32(1)
	v437 = *(*int32)(unsafe.Add(mBase, uint32(v170)+8))
	v441 = *(*int32)(unsafe.Add(mBase, uint32(v41)+164))
	*(*int32)(unsafe.Add(mBase, uint32(v437+v433<<(uint(int32(2))%32)))) = v441
	goto L64
L81:
	;
	goto L82
L82:
	;
	v443 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v170)+1)) = uint8(v443)
	goto L64
L83:
	;
	goto L63
L84:
	;
	goto L37
L85:
	;
	goto L31
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+84)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v41)+76)) = v584
	*(*int64)(unsafe.Add(mBase, uint32(v41)+164)) = int64(-8589934595)
	v593 = *(*int32)(unsafe.Add(mBase, uint32(v41+int32(32))+24))
	v594 = *(*int32)(unsafe.Add(mBase, uint32(v593)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v41)+172)) = v594
	v601 = F_hash_search(m, v584, v41+int32(164), int32(1), v41+int32(224))
	mBase = m.M
	v602 = m.ExcPending
	if v602 != 0 {
		goto L1
	} else {
		goto L87
	}
L87:
	;
	v603 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+224)))
	if v603 == int32(1) {
		goto L89
	} else {
		goto L90
	}
L88:
	;
	v1162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1128)+20)))
	if v1162&int32(2) != 0 {
		v4572 = v5
		goto L21
	} else {
		goto L150
	}
L89:
	;
	v606 = *(*int32)(unsafe.Add(mBase, uint32(v601)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v601)+20)) = v606 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v41)+80)) = v601
	v1128 = v601
	v1137 = v5
	goto L88
L90:
	;
	goto L91
L91:
	;
	v611 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v601)+20)) = v611
	v613 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v601)+12)) = v613
	*(*int32)(unsafe.Add(mBase, uint32(v41)+84)) = int32(1)
	*(*int64)(unsafe.Add(mBase, uint32(v601)+32)) = v613
	*(*int64)(unsafe.Add(mBase, uint32(v601)+24)) = int64(4294967295)
	v622 = F_lappend(m, v611, v601)
	mBase = m.M
	v623 = m.ExcPending
	if v623 != 0 {
		goto L1
	} else {
		goto L92
	}
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+88)) = v622
	v625 = *(*int32)(unsafe.Add(mBase, uint32(v601)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v601)+20)) = v625 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v41)+80)) = v601
	if v622 == int32(0) {
		v1128 = v601
		v1137 = v5
		goto L88
	} else {
		goto L93
	}
L93:
	;
	v632 = int32(0)
	v633 = *(*int32)(unsafe.Add(mBase, uint32(v622)+4))
	if v633 <= v632 {
		v1128 = v601
		v1137 = v5
		goto L88
	} else {
		goto L94
	}
L94:
	;
	v649 = v632
	goto L95
L95:
	;
	v676 = *(*int32)(unsafe.Add(mBase, uint32(v622)+12))
	v680 = *(*int32)(unsafe.Add(mBase, uint32(v676+v649<<(uint(int32(2))%32))))
	v681 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+100)))
	if v681 == int32(1) {
		goto L98
	} else {
		goto L99
	}
L96:
	;
	v1123 = *(*int32)(unsafe.Add(mBase, uint32(v41)+80))
	v1128 = v1123
	v1137 = v1042
	goto L88
L97:
	;
	v1042 = *(*int32)(unsafe.Add(mBase, uint32(v41)+96))
	if v1042 <= int32(1024) {
		goto L141
	} else {
		goto L142
	}
L98:
	;
	v684 = *(*int32)(unsafe.Add(mBase, uint32(v680)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v680)+20)) = v684 | int32(2)
	goto L97
L99:
	;
	goto L100
L100:
	;
	v688 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v41)+92)) = v688
	F_addKey(m, v41-int32(-64), v680, v680)
	mBase = m.M
	v694 = m.ExcPending
	if v694 != 0 {
		goto L1
	} else {
		goto L101
	}
L101:
	;
	v695 = *(*int32)(unsafe.Add(mBase, uint32(v41)+92))
	if v695 == int32(0) {
		v763 = v688
		goto L102
	} else {
		goto L103
	}
L102:
	;
	F_list_free(m, v763)
	mBase = m.M
	v798 = m.ExcPending
	if v798 != 0 {
		goto L1
	} else {
		goto L114
	}
L103:
	;
	v698 = *(*int32)(unsafe.Add(mBase, uint32(v695)+4))
	if v698 <= int32(0) {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	v763 = v695
	goto L102
L105:
	;
	goto L106
L106:
	;
	v705 = v688
	goto L107
L107:
	;
	v739 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v680)+20)))
	if v739&int32(2) == int32(0) {
		goto L109
	} else {
		goto L110
	}
L108:
	;
	v758 = *(*int32)(unsafe.Add(mBase, uint32(v41)+92))
	v763 = v758
	goto L102
L109:
	;
	v746 = *(*int32)(unsafe.Add(mBase, uint32(v695)+12))
	v750 = *(*int32)(unsafe.Add(mBase, uint32(v746+v705<<(uint(int32(2))%32))))
	F_addKey(m, v41-int32(-64), v680, v750)
	mBase = m.M
	v752 = m.ExcPending
	if v752 != 0 {
		goto L1
	} else {
		goto L112
	}
L110:
	;
	goto L111
L111:
	;
	goto L108
L112:
	;
	v754 = v705 + int32(1)
	v755 = *(*int32)(unsafe.Add(mBase, uint32(v695)+4))
	if v754 < v755 {
		v705 = v754
		goto L107
	} else {
		goto L113
	}
L113:
	;
	goto L111
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+92)) = int32(0)
	v801 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v680)+20)))
	if v801&int32(2) != 0 {
		goto L97
	} else {
		goto L115
	}
L115:
	;
	v804 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v41+int32(232)))) = v804
	*(*int64)(unsafe.Add(mBase, uint32(v41)+224)) = int64(0)
	v808 = *(*int32)(unsafe.Add(mBase, uint32(v680)+16))
	if v808 == v804 {
		goto L97
	} else {
		goto L116
	}
L116:
	;
	v811 = int32(0)
	v812 = *(*int32)(unsafe.Add(mBase, uint32(v808)+4))
	if v812 <= v811 {
		goto L97
	} else {
		goto L117
	}
L117:
	;
	v825 = v811
	goto L118
L118:
	;
	v853 = *(*int32)(unsafe.Add(mBase, uint32(v41)+64))
	v854 = *(*int32)(unsafe.Add(mBase, uint32(v808)+12))
	v858 = *(*int32)(unsafe.Add(mBase, uint32(v854+v825<<(uint(int32(2))%32))))
	v859 = *(*int32)(unsafe.Add(mBase, uint32(v858)+8))
	v860 = F_pg_reg_getnumoutarcs(m, v853, v859)
	mBase = m.M
	v861 = m.ExcPending
	if v861 != 0 {
		goto L1
	} else {
		goto L120
	}
L119:
	;
	goto L97
L120:
	;
	v864 = F_palloc(m, v860<<(uint(int32(3))%32))
	mBase = m.M
	v865 = m.ExcPending
	if v865 != 0 {
		goto L1
	} else {
		goto L121
	}
L121:
	;
	v866 = *(*int32)(unsafe.Add(mBase, uint32(v858)+8))
	F_pg_reg_getoutarcs(m, v853, v866, v864, v860)
	mBase = m.M
	v868 = m.ExcPending
	if v868 != 0 {
		goto L1
	} else {
		goto L122
	}
L122:
	;
	if int32(0) < v860 {
		goto L123
	} else {
		goto L124
	}
L123:
	;
	v876 = int32(0)
	goto L126
L124:
	;
	goto L125
L125:
	;
	F_pfree(m, v864)
	mBase = m.M
	v999 = m.ExcPending
	if v999 != 0 {
		goto L1
	} else {
		goto L138
	}
L126:
	;
	v912 = v864 + v876<<(uint(int32(3))%32)
	v913 = *(*int32)(unsafe.Add(mBase, uint32(v912)))
	if v913 < int32(0) {
		goto L128
	} else {
		goto L129
	}
L127:
	;
	goto L125
L128:
	;
	v958 = v876 + int32(1)
	if v958 != v860 {
		v876 = v958
		goto L126
	} else {
		goto L137
	}
L129:
	;
	v916 = *(*int32)(unsafe.Add(mBase, uint32(v41)+68))
	v919 = v916 + v913*int32(12)
	v920 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v919))))
	if v920 != int32(1) {
		goto L128
	} else {
		goto L130
	}
L130:
	;
	v923 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v919)+1)))
	if v923 == int32(1) {
		goto L131
	} else {
		goto L132
	}
L131:
	;
	v926 = *(*int32)(unsafe.Add(mBase, uint32(v858)+4))
	v927 = int32(-4)
	*(*int32)(unsafe.Add(mBase, uint32(v41)+228)) = v927
	*(*int32)(unsafe.Add(mBase, uint32(v41)+224)) = v926
	v930 = *(*int32)(unsafe.Add(mBase, uint32(v912)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v41)+232)) = v930
	F_addArc(m, v41-int32(-64), v680, v858, v927, v41+int32(224))
	mBase = m.M
	v938 = m.ExcPending
	if v938 != 0 {
		goto L1
	} else {
		goto L134
	}
L132:
	;
	goto L133
L133:
	;
	v940 = *(*int32)(unsafe.Add(mBase, uint32(v919)+4))
	if v940 <= int32(0) {
		goto L128
	} else {
		goto L135
	}
L134:
	;
	goto L133
L135:
	;
	v943 = *(*int32)(unsafe.Add(mBase, uint32(v858)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v41)+224)) = v943
	v945 = *(*int32)(unsafe.Add(mBase, uint32(v912)))
	*(*int32)(unsafe.Add(mBase, uint32(v41)+228)) = v945
	v947 = *(*int32)(unsafe.Add(mBase, uint32(v912)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v41)+232)) = v947
	F_addArc(m, v41-int32(-64), v680, v858, v945, v41+int32(224))
	mBase = m.M
	v954 = m.ExcPending
	if v954 != 0 {
		goto L1
	} else {
		goto L136
	}
L136:
	;
	goto L128
L137:
	;
	goto L127
L138:
	;
	v1001 = v825 + int32(1)
	v1002 = *(*int32)(unsafe.Add(mBase, uint32(v808)+4))
	if v1001 < v1002 {
		v825 = v1001
		goto L118
	} else {
		goto L139
	}
L139:
	;
	goto L119
L140:
	;
	v1120 = v649 + int32(1)
	v1121 = *(*int32)(unsafe.Add(mBase, uint32(v622)+4))
	if v1120 < v1121 {
		v649 = v1120
		goto L95
	} else {
		goto L149
	}
L141:
	;
	v1045 = *(*int32)(unsafe.Add(mBase, uint32(v41)+76))
	v1047 = *(*int32)(unsafe.Add(mBase, uint32(v1045)))
	v1048 = *(*int32)(unsafe.Add(mBase, uint32(v1047)+4))
	v1049 = *(*int32)(unsafe.Add(mBase, uint32(v1047)+412))
	if v1049 != 0 {
		goto L145
	} else {
		goto L146
	}
L142:
	;
	goto L143
L143:
	;
	v1117 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v41)+100)) = uint8(v1117)
	goto L140
L144:
	;
	if v1114 < int32(129) {
		goto L140
	} else {
		goto L148
	}
L145:
	;
	v1050 = *(*int32)(unsafe.Add(mBase, uint32(v1047)+376))
	v1051 = *(*int32)(unsafe.Add(mBase, uint32(v1047)+364))
	v1052 = *(*int32)(unsafe.Add(mBase, uint32(v1047)+352))
	v1053 = *(*int32)(unsafe.Add(mBase, uint32(v1047)+340))
	v1054 = *(*int32)(unsafe.Add(mBase, uint32(v1047)+328))
	v1055 = *(*int32)(unsafe.Add(mBase, uint32(v1047)+316))
	v1056 = *(*int32)(unsafe.Add(mBase, uint32(v1047)+304))
	v1057 = *(*int32)(unsafe.Add(mBase, uint32(v1047)+292))
	v1058 = *(*int32)(unsafe.Add(mBase, uint32(v1047)+280))
	v1059 = *(*int32)(unsafe.Add(mBase, uint32(v1047)+268))
	v1060 = *(*int32)(unsafe.Add(mBase, uint32(v1047)+256))
	v1061 = *(*int32)(unsafe.Add(mBase, uint32(v1047)+244))
	v1062 = *(*int32)(unsafe.Add(mBase, uint32(v1047)+232))
	v1063 = *(*int32)(unsafe.Add(mBase, uint32(v1047)+220))
	v1064 = *(*int32)(unsafe.Add(mBase, uint32(v1047)+208))
	v1065 = *(*int32)(unsafe.Add(mBase, uint32(v1047)+196))
	v1066 = *(*int32)(unsafe.Add(mBase, uint32(v1047)+184))
	v1067 = *(*int32)(unsafe.Add(mBase, uint32(v1047)+172))
	v1068 = *(*int32)(unsafe.Add(mBase, uint32(v1047)+160))
	v1069 = *(*int32)(unsafe.Add(mBase, uint32(v1047)+148))
	v1070 = *(*int32)(unsafe.Add(mBase, uint32(v1047)+136))
	v1071 = *(*int32)(unsafe.Add(mBase, uint32(v1047)+124))
	v1072 = *(*int32)(unsafe.Add(mBase, uint32(v1047)+112))
	v1073 = *(*int32)(unsafe.Add(mBase, uint32(v1047)+100))
	v1074 = *(*int32)(unsafe.Add(mBase, uint32(v1047)+88))
	v1075 = *(*int32)(unsafe.Add(mBase, uint32(v1047)+76))
	v1078 = *(*int32)(unsafe.Add(mBase, uint32(v1047-int32(-64))))
	v1079 = *(*int32)(unsafe.Add(mBase, uint32(v1047)+52))
	v1080 = *(*int32)(unsafe.Add(mBase, uint32(v1047)+40))
	v1081 = *(*int32)(unsafe.Add(mBase, uint32(v1047)+28))
	v1082 = *(*int32)(unsafe.Add(mBase, uint32(v1047)+16))
	v1114 = v1050 + (v1051 + (v1052 + (v1053 + (v1054 + (v1055 + (v1056 + (v1057 + (v1058 + (v1059 + (v1060 + (v1061 + (v1062 + (v1063 + (v1064 + (v1065 + (v1066 + (v1067 + (v1068 + (v1069 + (v1070 + (v1071 + (v1072 + (v1073 + (v1074 + (v1075 + (v1078 + (v1079 + (v1080 + (v1081 + (v1082 + v1048))))))))))))))))))))))))))))))
	goto L147
L146:
	;
	v1114 = v1048
	goto L147
L147:
	;
	goto L144
L148:
	;
	goto L143
L149:
	;
	goto L96
L150:
	;
	v1166 = v1137 << (uint(int32(5)) % 32)
	v1167 = F_palloc0(m, v1166)
	mBase = m.M
	v1168 = m.ExcPending
	if v1168 != 0 {
		goto L1
	} else {
		goto L151
	}
L151:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+104)) = v1167
	v1172 = *(*int32)(unsafe.Add(mBase, uint32(v41)+76))
	F_hash_seq_init(m, v41+int32(176), v1172)
	mBase = m.M
	v1174 = m.ExcPending
	if v1174 != 0 {
		goto L1
	} else {
		goto L152
	}
L152:
	;
	v1177 = F_hash_seq_search(m, v41+int32(176))
	mBase = m.M
	v1178 = m.ExcPending
	if v1178 != 0 {
		goto L1
	} else {
		goto L153
	}
L153:
	;
	if v1177 != 0 {
		goto L154
	} else {
		goto L155
	}
L154:
	;
	v1181 = int32(0)
	v1192 = v1177
	goto L157
L155:
	;
	goto L156
L156:
	;
	if int32(2) <= v1137 {
		goto L169
	} else {
		goto L170
	}
L157:
	;
	v1218 = *(*int32)(unsafe.Add(mBase, uint32(v1192)+12))
	if v1218 == int32(0) {
		v1300 = v1181
		goto L159
	} else {
		goto L160
	}
L158:
	;
	goto L156
L159:
	;
	v1339 = F_hash_seq_search(m, v41+int32(176))
	mBase = m.M
	v1340 = m.ExcPending
	if v1340 != 0 {
		goto L1
	} else {
		goto L167
	}
L160:
	;
	v1221 = int32(0)
	v1222 = *(*int32)(unsafe.Add(mBase, uint32(v1218)+4))
	if v1222 <= v1221 {
		v1300 = v1181
		goto L159
	} else {
		goto L161
	}
L161:
	;
	v1226 = v1181
	v1232 = v1221
	goto L162
L162:
	;
	v1263 = *(*int32)(unsafe.Add(mBase, uint32(v1218)+12))
	v1267 = *(*int32)(unsafe.Add(mBase, uint32(v1263+v1232<<(uint(int32(2))%32))))
	v1269 = F_palloc(m, int32(8))
	mBase = m.M
	v1270 = m.ExcPending
	if v1270 != 0 {
		goto L1
	} else {
		goto L164
	}
L163:
	;
	v1300 = v1294
	goto L159
L164:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1269))) = v1192
	v1272 = *(*int32)(unsafe.Add(mBase, uint32(v1267)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v1269)+4)) = v1272
	v1276 = v1167 + v1226<<(uint(int32(5))%32)
	v1277 = *(*int32)(unsafe.Add(mBase, uint32(v1267)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1276)+8)) = v1277
	v1279 = *(*int64)(unsafe.Add(mBase, uint32(v1267)))
	*(*int64)(unsafe.Add(mBase, uint32(v1276))) = v1279
	v1281 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1276)+24)) = uint8(v1281)
	*(*int32)(unsafe.Add(mBase, uint32(v1276)+12)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v41)+12)) = v1269
	*(*int32)(unsafe.Add(mBase, uint32(v41)+224)) = v1269
	v1290 = F_list_make1_impl(m, v1281, v41+int32(12))
	mBase = m.M
	v1291 = m.ExcPending
	if v1291 != 0 {
		goto L1
	} else {
		goto L165
	}
L165:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1276)+28)) = v1290
	v1293 = int32(1)
	v1294 = v1226 + v1293
	v1296 = v1232 + v1293
	v1297 = *(*int32)(unsafe.Add(mBase, uint32(v1218)+4))
	if v1296 < v1297 {
		v1226 = v1294
		v1232 = v1296
		goto L162
	} else {
		goto L166
	}
L166:
	;
	goto L163
L167:
	;
	if v1339 != 0 {
		v1181 = v1300
		v1192 = v1339
		goto L157
	} else {
		goto L168
	}
L168:
	;
	goto L158
L169:
	;
	F_pg_qsort(m, v1167, v1137, int32(32), int32(_a_F_createTrgmNFA_5))
	mBase = m.M
	v1384 = m.ExcPending
	if v1384 != 0 {
		goto L1
	} else {
		goto L172
	}
L170:
	;
	v1528 = v1137
	goto L171
L171:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+108)) = v1528
	if int32(0) < v1528 {
		goto L199
	} else {
		goto L200
	}
L172:
	;
	__phi1388 = v1167 + int32(32)
	__phi1392 = v1167
	__phi1394 = v1167
	v1388 = __phi1388
	v1392 = __phi1392
	v1394 = __phi1394
	goto L173
L173:
	;
	v1426 = int32(12)
	goto L179
L174:
	;
	v1528 = (v1506-v1167)>>(uint(int32(5))%32) + int32(1)
	goto L171
L175:
	;
	v1508 = v1388 + int32(32)
	if base.Ui32(v1508) < base.Ui32(v1167+v1166) {
		__phi1388 = v1508
		__phi1392 = v1506
		__phi1394 = v1388
		v1388 = __phi1388
		v1392 = __phi1392
		v1394 = __phi1394
		goto L173
	} else {
		goto L198
	}
L176:
	;
	if int32(0) < v1488 {
		goto L194
	} else {
		goto L195
	}
L177:
	;
	v1488 = int32(0)
	goto L176
L178:
	;
	v1462 = v1457
	v1463 = v1458
	v1464 = v1459
	goto L188
L179:
	;
	if (v1388|v1392)&int32(3) != 0 {
		v1457 = v1388
		v1458 = v1392
		v1459 = v1426
		goto L178
	} else {
		goto L182
	}
L181:
	;
	if v1447 == int32(0) {
		goto L177
	} else {
		goto L187
	}
L182:
	;
	v1434 = v1388
	v1435 = v1392
	v1436 = v1426
	goto L183
L183:
	;
	v1439 = *(*int32)(unsafe.Add(mBase, uint32(v1434)))
	v1440 = *(*int32)(unsafe.Add(mBase, uint32(v1435)))
	if v1439 != v1440 {
		v1457 = v1434
		v1458 = v1435
		v1459 = v1436
		goto L178
	} else {
		goto L185
	}
L184:
	;
	goto L181
L185:
	;
	v1442 = int32(4)
	v1443 = v1435 + v1442
	v1445 = v1434 + v1442
	v1447 = v1436 - v1442
	if base.Ui32(int32(3)) < base.Ui32(v1447) {
		v1434 = v1445
		v1435 = v1443
		v1436 = v1447
		goto L183
	} else {
		goto L186
	}
L186:
	;
	goto L184
L187:
	;
	v1457 = v1445
	v1458 = v1443
	v1459 = v1447
	goto L178
L188:
	;
	v1467 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1462))))
	v1468 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1463))))
	if v1467 == v1468 {
		goto L190
	} else {
		goto L191
	}
L189:
	;
	v1488 = v1467 - v1468
	goto L176
L190:
	;
	v1470 = int32(1)
	v1475 = v1464 - v1470
	if v1475 != 0 {
		v1462 = v1462 + v1470
		v1463 = v1463 + v1470
		v1464 = v1475
		goto L188
	} else {
		goto L193
	}
L191:
	;
	goto L192
L192:
	;
	goto L189
L193:
	;
	goto L177
L194:
	;
	v1491 = *(*int64)(unsafe.Add(mBase, uint32(v1388)))
	*(*int64)(unsafe.Add(mBase, uint32(v1392)+32)) = v1491
	v1493 = *(*int64)(unsafe.Add(mBase, uint32(v1388)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v1392)+56)) = v1493
	v1495 = *(*int64)(unsafe.Add(mBase, uint32(v1388)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v1392)+48)) = v1495
	v1497 = *(*int64)(unsafe.Add(mBase, uint32(v1388)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v1392)+40)) = v1497
	v1506 = v1392 + int32(32)
	goto L175
L195:
	;
	goto L196
L196:
	;
	v1501 = *(*int32)(unsafe.Add(mBase, uint32(v1392)+28))
	v1502 = *(*int32)(unsafe.Add(mBase, uint32(v1394)+60))
	v1503 = F_list_concat(m, v1501, v1502)
	mBase = m.M
	v1504 = m.ExcPending
	if v1504 != 0 {
		goto L1
	} else {
		goto L197
	}
L197:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1392)+28)) = v1503
	v1506 = v1392
	goto L175
L198:
	;
	goto L174
L199:
	;
	v1557 = *(*int32)(unsafe.Add(mBase, uint32(v41)+68))
	v1565 = int32(0)
	v1593 = int64(0)
	v1594 = float32(0)
	goto L202
L200:
	;
	goto L201
L201:
	;
	F_pg_qsort(m, v1167, v1528, int32(32), int32(_a_F_createTrgmNFA_6))
	mBase = m.M
	v2717 = m.ExcPending
	if v2717 != 0 {
		goto L1
	} else {
		goto L319
	}
L202:
	;
	v1600 = v1167 + v1565<<(uint(int32(5))%32)
	v1601 = *(*int32)(unsafe.Add(mBase, uint32(v1600)))
	if v1601 != int32(-4) {
		goto L204
	} else {
		goto L205
	}
L203:
	;
	F_pg_qsort(m, v1167, v1528, int32(32), int32(_a_F_createTrgmNFA_6))
	mBase = m.M
	v1655 = m.ExcPending
	if v1655 != 0 {
		goto L1
	} else {
		goto L216
	}
L204:
	;
	v1608 = *(*int32)(unsafe.Add(mBase, uint32(v1557+v1601*int32(12))+4))
	v1609 = v1608
	v1610 = int32(0)
	goto L206
L205:
	;
	v1609 = int32(1)
	v1610 = int32(2)
	goto L206
L206:
	;
	v1611 = *(*int32)(unsafe.Add(mBase, uint32(v1600)+4))
	if v1611 != int32(-4) {
		goto L208
	} else {
		goto L209
	}
L207:
	;
	v1624 = v1622 << (uint(int32(1)) % 32)
	v1625 = *(*int32)(unsafe.Add(mBase, uint32(v1600)+8))
	if v1625 != int32(-4) {
		goto L212
	} else {
		goto L213
	}
L208:
	;
	v1617 = *(*int32)(unsafe.Add(mBase, uint32(v1557+v1611*int32(12))+4))
	v1621 = v1617 * v1609
	v1622 = v1610
	goto L207
L209:
	;
	goto L210
L210:
	;
	v1621 = v1609
	v1622 = v1610 | int32(1)
	goto L207
L211:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1600)+16)) = v1635
	v1642 = *(*float32)(unsafe.Add(mBase, uint32(v1636<<(uint(int32(2))%32))+uint32(_c_F_createTrgmNFA[1])))
	v1644 = base.F32_mul(v1642, base.F32_convert_i32_s(v1635))
	*(*float32)(unsafe.Add(mBase, uint32(v1600)+20)) = v1644
	v1646 = base.F32_add(v1594, v1644)
	v1648 = v1593 + base.I64_extend_i32_s(v1635)
	v1650 = v1565 + int32(1)
	if v1650 != v1528 {
		v1565 = v1650
		v1593 = v1648
		v1594 = v1646
		goto L202
	} else {
		goto L215
	}
L212:
	;
	v1631 = *(*int32)(unsafe.Add(mBase, uint32(v1557+v1625*int32(12))+4))
	v1635 = v1631 * v1621
	v1636 = v1624
	goto L211
L213:
	;
	goto L214
L214:
	;
	v1635 = v1621
	v1636 = v1624 | int32(1)
	goto L211
L215:
	;
	goto L203
L216:
	;
	v1673 = int32(0)
	v1692 = v1648
	v1693 = v1646
	goto L217
L217:
	;
	if base.F32_le(v1693, float32(16)) == int32(0) {
		goto L219
	} else {
		goto L220
	}
L218:
	;
	if v2524 <= int64(256) {
		goto L290
	} else {
		goto L291
	}
L219:
	;
	v1701 = v1167 + v1673<<(uint(int32(5))%32)
	v1702 = *(*int32)(unsafe.Add(mBase, uint32(v1701)+28))
	if v1702 == int32(0) {
		goto L223
	} else {
		goto L224
	}
L220:
	;
	v2524 = v1692
	goto L221
L221:
	;
	goto L218
L222:
	;
	v2487 = v1673 + int32(1)
	if v2487 != v1528 {
		v1673 = v2487
		v1692 = v2483
		v1693 = v2484
		goto L217
	} else {
		goto L289
	}
L223:
	;
	v2442 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1701)+24)) = uint8(v2442)
	v2444 = *(*float32)(unsafe.Add(mBase, uint32(v1701)+20))
	v2446 = int64(*(*int32)(unsafe.Add(mBase, uint32(v1701)+16)))
	v2483 = v1692 - v2446
	v2484 = base.F32_sub(v1693, v2444)
	goto L222
L224:
	;
	v1705 = *(*int32)(unsafe.Add(mBase, uint32(v1702)+4))
	if v1705 <= int32(0) {
		goto L223
	} else {
		goto L225
	}
L225:
	;
	v1708 = *(*int32)(unsafe.Add(mBase, uint32(v1702)+12))
	v1720 = int32(0)
	v1722 = v1705
	goto L226
L226:
	;
	v1751 = *(*int32)(unsafe.Add(mBase, uint32(v1708+v1720<<(uint(int32(2))%32))))
	v1752 = *(*int32)(unsafe.Add(mBase, uint32(v1751)+4))
	v1753 = *(*int32)(unsafe.Add(mBase, uint32(v1751)))
	v1754 = v1753
	goto L228
L227:
	;
	v1942 = int32(0)
	if v1942 < v1941 {
		goto L251
	} else {
		goto L252
	}
L228:
	;
	v1792 = *(*int32)(unsafe.Add(mBase, uint32(v1754)+28))
	if v1792 != 0 {
		v1754 = v1792
		goto L228
	} else {
		goto L230
	}
L229:
	;
	v1799 = v1752
	goto L231
L230:
	;
	goto L229
L231:
	;
	v1831 = *(*int32)(unsafe.Add(mBase, uint32(v1799)+28))
	if v1831 != 0 {
		v1799 = v1831
		goto L231
	} else {
		goto L233
	}
L232:
	;
	v1832 = *(*int32)(unsafe.Add(mBase, uint32(v1754)+32))
	v1833 = *(*int32)(unsafe.Add(mBase, uint32(v1754)+20))
	v1839 = v1754
	v1842 = v1832 | v1833
	v1843 = v1832
	goto L234
L233:
	;
	goto L232
L234:
	;
	v1873 = *(*int32)(unsafe.Add(mBase, uint32(v1839)+36))
	if v1873 != 0 {
		goto L236
	} else {
		goto L237
	}
L235:
	;
	v1878 = *(*int32)(unsafe.Add(mBase, uint32(v1799)+32))
	v1879 = *(*int32)(unsafe.Add(mBase, uint32(v1799)+20))
	v1881 = v1799
	v1887 = v1878 | v1879
	goto L239
L236:
	;
	v1874 = *(*int32)(unsafe.Add(mBase, uint32(v1873)+20))
	v1876 = *(*int32)(unsafe.Add(mBase, uint32(v1873)+32))
	v1839 = v1873
	v1842 = v1874 | v1842 | v1876
	v1843 = v1876
	goto L234
L237:
	;
	goto L238
L238:
	;
	goto L235
L239:
	;
	v1919 = *(*int32)(unsafe.Add(mBase, uint32(v1881)+36))
	if v1919 != 0 {
		goto L241
	} else {
		goto L242
	}
L240:
	;
	v1925 = int32(3)
	v1928 = base.B2i32((v1887|v1842)&v1925 == v1925)
	if v1928 == int32(0) {
		goto L244
	} else {
		goto L245
	}
L241:
	;
	v1920 = *(*int32)(unsafe.Add(mBase, uint32(v1919)+20))
	v1922 = *(*int32)(unsafe.Add(mBase, uint32(v1919)+32))
	v1881 = v1919
	v1887 = v1920 | v1887 | v1922
	goto L239
L242:
	;
	goto L243
L243:
	;
	goto L240
L244:
	;
	if v1881 != v1839 {
		goto L247
	} else {
		goto L248
	}
L245:
	;
	v1941 = v1722
	goto L246
L246:
	;
	goto L227
L247:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1881)+36)) = v1839
	*(*int32)(unsafe.Add(mBase, uint32(v1839)+32)) = v1887 | v1843
	v1935 = *(*int32)(unsafe.Add(mBase, uint32(v1702)+4))
	v1936 = v1935
	goto L249
L248:
	;
	v1936 = v1722
	goto L249
L249:
	;
	v1938 = v1720 + int32(1)
	if v1938 < v1936 {
		v1720 = v1938
		v1722 = v1936
		goto L226
	} else {
		goto L250
	}
L250:
	;
	v1941 = v1936
	goto L246
L251:
	;
	v1952 = v1942
	goto L254
L252:
	;
	v2242 = v1941
	goto L253
L253:
	;
	if (v1887|v1842)&v1925 == v1925 {
		v2483 = v1692
		v2484 = v1693
		goto L222
	} else {
		goto L275
	}
L254:
	;
	v1986 = *(*int32)(unsafe.Add(mBase, uint32(v1708+v1952<<(uint(int32(2))%32))))
	v1987 = *(*int32)(unsafe.Add(mBase, uint32(v1986)+4))
	v1988 = *(*int32)(unsafe.Add(mBase, uint32(v1986)))
	v1989 = v1988
	goto L256
L255:
	;
	v2242 = v2228
	goto L253
L256:
	;
	v2027 = *(*int32)(unsafe.Add(mBase, uint32(v1989)+28))
	if v2027 != 0 {
		v1989 = v2027
		goto L256
	} else {
		goto L258
	}
L257:
	;
	v2034 = v1987
	goto L259
L258:
	;
	goto L257
L259:
	;
	v2066 = *(*int32)(unsafe.Add(mBase, uint32(v2034)+28))
	if v2066 != 0 {
		v2034 = v2066
		goto L259
	} else {
		goto L261
	}
L260:
	;
	if v1989 != 0 {
		goto L262
	} else {
		goto L263
	}
L261:
	;
	goto L260
L262:
	;
	v2071 = v1989
	goto L265
L263:
	;
	goto L264
L264:
	;
	v2146 = *(*int32)(unsafe.Add(mBase, uint32(v2034)+36))
	if v2146 != 0 {
		goto L268
	} else {
		goto L269
	}
L265:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2071)+32)) = int32(0)
	v2107 = *(*int32)(unsafe.Add(mBase, uint32(v2071)+36))
	if v2107 != 0 {
		v2071 = v2107
		goto L265
	} else {
		goto L267
	}
L266:
	;
	goto L264
L267:
	;
	goto L266
L268:
	;
	__phi2147 = v2034
	__phi2151 = v2146
	v2147 = __phi2147
	v2151 = __phi2151
	goto L271
L269:
	;
	goto L270
L270:
	;
	v2227 = v1952 + int32(1)
	v2228 = *(*int32)(unsafe.Add(mBase, uint32(v1702)+4))
	if v2227 < v2228 {
		v1952 = v2227
		goto L254
	} else {
		goto L274
	}
L271:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v2147)+32)) = int64(0)
	v2187 = *(*int32)(unsafe.Add(mBase, uint32(v2151)+36))
	if v2187 != 0 {
		__phi2147 = v2151
		__phi2151 = v2187
		v2147 = __phi2147
		v2151 = __phi2151
		goto L271
	} else {
		goto L273
	}
L272:
	;
	goto L270
L273:
	;
	goto L272
L274:
	;
	goto L255
L275:
	;
	v2268 = int32(0)
	if v2242 <= v2268 {
		goto L223
	} else {
		goto L276
	}
L276:
	;
	v2278 = v2268
	v2283 = v2242
	goto L277
L277:
	;
	v2312 = *(*int32)(unsafe.Add(mBase, uint32(v1708+v2278<<(uint(int32(2))%32))))
	v2313 = *(*int32)(unsafe.Add(mBase, uint32(v2312)+4))
	v2314 = *(*int32)(unsafe.Add(mBase, uint32(v2312)))
	v2315 = v2314
	goto L279
L278:
	;
	goto L223
L279:
	;
	v2353 = *(*int32)(unsafe.Add(mBase, uint32(v2315)+28))
	if v2353 != 0 {
		v2315 = v2353
		goto L279
	} else {
		goto L281
	}
L280:
	;
	v2358 = v2313
	goto L282
L281:
	;
	goto L280
L282:
	;
	v2392 = *(*int32)(unsafe.Add(mBase, uint32(v2358)+28))
	if v2392 != 0 {
		v2358 = v2392
		goto L282
	} else {
		goto L284
	}
L283:
	;
	if v2358 != v2315 {
		goto L285
	} else {
		goto L286
	}
L284:
	;
	goto L283
L285:
	;
	v2394 = *(*int32)(unsafe.Add(mBase, uint32(v2315)+20))
	v2395 = *(*int32)(unsafe.Add(mBase, uint32(v2358)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v2315)+20)) = v2394 | v2395
	*(*int32)(unsafe.Add(mBase, uint32(v2358)+28)) = v2315
	v2399 = *(*int32)(unsafe.Add(mBase, uint32(v1702)+4))
	v2400 = v2399
	goto L287
L286:
	;
	v2400 = v2283
	goto L287
L287:
	;
	v2402 = v2278 + int32(1)
	if v2402 < v2400 {
		v2278 = v2402
		v2283 = v2400
		goto L277
	} else {
		goto L288
	}
L288:
	;
	goto L278
L289:
	;
	v2524 = v2483
	goto L221
L290:
	;
	v2529 = base.I32_wrap_i64(v2524)
	*(*int32)(unsafe.Add(mBase, uint32(v41)+112)) = v2529
	F_pg_qsort(m, v1167, v1528, int32(32), int32(_a_F_createTrgmNFA_5))
	mBase = m.M
	v2534 = m.ExcPending
	if v2534 != 0 {
		goto L1
	} else {
		goto L293
	}
L291:
	;
	goto L292
L292:
	;
	v4572 = v5
	goto L21
L293:
	;
	v2536 = v1528 & int32(3)
	v2537 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v1528) {
		goto L294
	} else {
		goto L295
	}
L294:
	;
	v2545 = v2537
	v2546 = int32(0)
	v2549 = v2537
	goto L297
L295:
	;
	v2621 = v2537
	v2625 = v2537
	goto L296
L296:
	;
	if v2536 == int32(0) {
		v2759 = v2529
		goto L22
	} else {
		goto L312
	}
L297:
	;
	v2585 = v1167 + v2545<<(uint(int32(5))%32)
	v2586 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2585)+24)))
	if v2586 == int32(1) {
		goto L299
	} else {
		goto L300
	}
L298:
	;
	v2621 = v2617
	v2625 = v2615
	goto L296
L299:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2585)+12)) = v2549
	v2592 = v2549 + int32(1)
	goto L301
L300:
	;
	v2592 = v2549
	goto L301
L301:
	;
	v2593 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2585)+56)))
	if v2593 == int32(1) {
		goto L302
	} else {
		goto L303
	}
L302:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2585)+44)) = v2592
	v2599 = v2592 + int32(1)
	goto L304
L303:
	;
	v2599 = v2592
	goto L304
L304:
	;
	v2600 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2585)+88)))
	if v2600 == int32(1) {
		goto L305
	} else {
		goto L306
	}
L305:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2585-int32(-64))+12)) = v2599
	v2608 = v2599 + int32(1)
	goto L307
L306:
	;
	v2608 = v2599
	goto L307
L307:
	;
	v2609 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2585)+120)))
	if v2609 == int32(1) {
		goto L308
	} else {
		goto L309
	}
L308:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2585)+108)) = v2608
	v2615 = v2608 + int32(1)
	goto L310
L309:
	;
	v2615 = v2608
	goto L310
L310:
	;
	v2616 = int32(4)
	v2617 = v2545 + v2616
	v2619 = v2546 + v2616
	if v2619 != v1528&int32(2147483644) {
		v2545 = v2617
		v2546 = v2619
		v2549 = v2615
		goto L297
	} else {
		goto L311
	}
L311:
	;
	goto L298
L312:
	;
	v2661 = v2621
	v2665 = v2625
	v2668 = v2537
	goto L313
L313:
	;
	v2701 = v1167 + v2661<<(uint(int32(5))%32)
	v2702 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2701)+24)))
	if v2702 == int32(1) {
		goto L315
	} else {
		goto L316
	}
L314:
	;
	v2759 = v2529
	goto L22
L315:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2701)+12)) = v2665
	v2708 = v2665 + int32(1)
	goto L317
L316:
	;
	v2708 = v2665
	goto L317
L317:
	;
	v2709 = int32(1)
	v2712 = v2668 + v2709
	if v2712 != v2536 {
		v2661 = v2661 + v2709
		v2665 = v2708
		v2668 = v2712
		goto L313
	} else {
		goto L318
	}
L318:
	;
	goto L314
L319:
	;
	v2718 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v41)+112)) = v2718
	F_pg_qsort(m, v1167, v1528, int32(32), int32(_a_F_createTrgmNFA_5))
	mBase = m.M
	v2724 = m.ExcPending
	if v2724 != 0 {
		goto L1
	} else {
		goto L320
	}
L320:
	;
	v2759 = v2718
	goto L22
L321:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2732 = m.ExcPending
	if v2732 != 0 {
		goto L1
	} else {
		goto L322
	}
L322:
	;
	F_errcode(m, int32(302252162))
	mBase = m.M
	v2735 = m.ExcPending
	if v2735 != 0 {
		goto L1
	} else {
		goto L323
	}
L323:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+16)) = v41 - int32(-64)
	F_errmsg(m, int32(_a_F_createTrgmNFA_7), v41+int32(16))
	mBase = m.M
	v2743 = m.ExcPending
	if v2743 != 0 {
		goto L1
	} else {
		goto L324
	}
L324:
	;
	F_errfinish(m, int32(_a_F_createTrgmNFA_8), int32(751), int32(_a_F_createTrgmNFA_9))
	mBase = m.M
	v2748 = m.ExcPending
	if v2748 != 0 {
		goto L1
	} else {
		goto L325
	}
L325:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L326:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2793))) = v2759*int32(12) + int32(20)
	v2800 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v2793)+4)) = uint8(v2800)
	if int32(0) < v1528 {
		goto L327
	} else {
		goto L328
	}
L327:
	;
	v2813 = v41 + int32(177)
	v2821 = v2793 + int32(5)
	v2827 = v1528
	v2838 = v5
	goto L330
L328:
	;
	v3442 = v1172
	goto L329
L329:
	;
	F_hash_seq_init(m, v41+int32(176), v3442)
	mBase = m.M
	v3466 = m.ExcPending
	if v3466 != 0 {
		goto L1
	} else {
		goto L413
	}
L330:
	;
	v2852 = *(*int32)(unsafe.Add(mBase, uint32(v41)+104))
	v2855 = v2852 + v2838<<(uint(int32(5))%32)
	v2856 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2855)+24)))
	if v2856 != int32(1) {
		v3390 = v2821
		v3396 = v2827
		goto L332
	} else {
		goto L333
	}
L331:
	;
	v3424 = *(*int32)(unsafe.Add(mBase, uint32(v41)+76))
	v3442 = v3424
	goto L329
L332:
	;
	v3422 = v2838 + int32(1)
	if v3422 < v3396 {
		v2821 = v3390
		v2827 = v3396
		v2838 = v3422
		goto L330
	} else {
		goto L412
	}
L333:
	;
	v2859 = *(*int32)(unsafe.Add(mBase, uint32(v41)+68))
	v2860 = *(*int32)(unsafe.Add(mBase, uint32(v2855)))
	v2863 = v2859 + v2860*int32(12)
	v2864 = int32(1)
	v2867 = base.B2i32(v2860 == int32(-4))
	if v2867 == int32(0) {
		goto L334
	} else {
		goto L335
	}
L334:
	;
	v2870 = *(*int32)(unsafe.Add(mBase, uint32(v2863)+4))
	if v2870 <= int32(0) {
		v3390 = v2821
		v3396 = v2827
		goto L332
	} else {
		goto L337
	}
L335:
	;
	v2873 = v2864
	goto L336
L336:
	;
	v2874 = *(*int32)(unsafe.Add(mBase, uint32(v2855)+4))
	v2875 = int32(12)
	v2877 = v2859 + v2874*v2875
	v2878 = *(*int32)(unsafe.Add(mBase, uint32(v2855)+8))
	v2882 = base.B2i32(v2874 == int32(-4))
	if v2882 == int32(0) {
		goto L338
	} else {
		goto L339
	}
L337:
	;
	v2873 = v2870
	goto L336
L338:
	;
	v2885 = *(*int32)(unsafe.Add(mBase, uint32(v2877)+4))
	v2886 = v2885
	goto L340
L339:
	;
	v2886 = v2864
	goto L340
L340:
	;
	v2887 = v2859 + v2878*v2875
	v2889 = v2873
	v2893 = v2886
	v2896 = v2821
	v2910 = int32(0)
	goto L341
L341:
	;
	if int32(0) < v2893 {
		goto L344
	} else {
		goto L345
	}
L342:
	;
	v3382 = *(*int32)(unsafe.Add(mBase, uint32(v41)+108))
	v3390 = v3351
	v3396 = v3382
	goto L332
L343:
	;
	goto L342
L344:
	;
	if v2860 == int32(-4) {
		goto L347
	} else {
		goto L348
	}
L345:
	;
	v3303 = v2889
	v3307 = v2893
	v3310 = v2896
	goto L346
L346:
	;
	v3342 = v2910 + int32(1)
	if v3342 < v3303 {
		v2889 = v3303
		v2893 = v3307
		v2896 = v3310
		v2910 = v3342
		goto L341
	} else {
		goto L411
	}
L347:
	;
	v2932 = v41 + int32(224)
	goto L349
L348:
	;
	v2931 = *(*int32)(unsafe.Add(mBase, uint32(v2863)+8))
	v2932 = v2931
	goto L349
L349:
	;
	v2936 = *(*int32)(unsafe.Add(mBase, uint32(v2932+v2910<<(uint(int32(2))%32))))
	v2939 = base.B2i32(v2878 == int32(-4))
	if v2939 == int32(0) {
		goto L350
	} else {
		goto L351
	}
L350:
	;
	v2942 = *(*int32)(unsafe.Add(mBase, uint32(v2887)+4))
	v2943 = v2942
	goto L352
L351:
	;
	v2943 = int32(1)
	goto L352
L352:
	;
	v2945 = int32(base.Ui32(v2936) >> (uint(int32(24)) % 32))
	v2947 = int32(base.Ui32(v2936) >> (uint(int32(16)) % 32))
	v2949 = int32(base.Ui32(v2936) >> (uint(int32(8)) % 32))
	v2953 = v2943
	v2957 = v2893
	v2960 = v2896
	v2975 = int32(0)
	goto L353
L353:
	;
	if int32(0) < v2953 {
		goto L356
	} else {
		goto L357
	}
L354:
	;
	if v2860 == int32(-4) {
		v3351 = v3271
		goto L343
	} else {
		goto L410
	}
L355:
	;
	goto L354
L356:
	;
	if v2874 == int32(-4) {
		goto L359
	} else {
		goto L360
	}
L357:
	;
	v3223 = v2953
	v3227 = v2957
	v3230 = v2960
	goto L358
L358:
	;
	v3262 = v2975 + int32(1)
	if v3262 < v3227 {
		v2953 = v3223
		v2957 = v3227
		v2960 = v3230
		v2975 = v3262
		goto L353
	} else {
		goto L409
	}
L359:
	;
	v2996 = v41 + int32(224)
	goto L361
L360:
	;
	v2995 = *(*int32)(unsafe.Add(mBase, uint32(v2877)+8))
	v2996 = v2995
	goto L361
L361:
	;
	v3000 = *(*int32)(unsafe.Add(mBase, uint32(v2996+v2975<<(uint(int32(2))%32))))
	v3002 = int32(base.Ui32(v3000) >> (uint(int32(24)) % 32))
	v3004 = int32(base.Ui32(v3000) >> (uint(int32(16)) % 32))
	v3006 = int32(base.Ui32(v3000) >> (uint(int32(8)) % 32))
	v3009 = int32(0)
	v3015 = v2960
	goto L362
L362:
	;
	if v2878 == int32(-4) {
		goto L364
	} else {
		goto L365
	}
L363:
	;
	if v2874 == int32(-4) {
		v3268 = int32(1)
		v3271 = v3213
		goto L355
	} else {
		goto L408
	}
L364:
	;
	v3049 = v41 + int32(224)
	goto L366
L365:
	;
	v3048 = *(*int32)(unsafe.Add(mBase, uint32(v2887)+8))
	v3049 = v3048
	goto L366
L366:
	;
	v3053 = *(*int32)(unsafe.Add(mBase, uint32(v3049+v3009<<(uint(int32(2))%32))))
	if v2936&int32(255) != 0 {
		goto L368
	} else {
		goto L369
	}
L367:
	;
	v3072 = v3070 + int32(1)
	if v3000&int32(255) != 0 {
		goto L375
	} else {
		goto L376
	}
L368:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v41)+176)) = uint8(v2936)
	if v2949&int32(255) == int32(0) {
		v3070 = v2813
		goto L367
	} else {
		goto L371
	}
L369:
	;
	goto L370
L370:
	;
	v3068 = int32(32)
	*(*uint8)(unsafe.Add(mBase, uint32(v41)+176)) = uint8(v3068)
	v3070 = v2813
	goto L367
L371:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v41)+177)) = uint8(v2949)
	if v2947&int32(255) == int32(0) {
		v3070 = v41 + int32(178)
		goto L367
	} else {
		goto L372
	}
L372:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v41)+178)) = uint8(v2947)
	if base.Ui32(v2936) < base.Ui32(int32(16777216)) {
		v3070 = v41 + int32(179)
		goto L367
	} else {
		goto L373
	}
L373:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v41)+179)) = uint8(v2945)
	v3070 = v41 + int32(180)
	goto L367
L374:
	;
	v3099 = v3097 + int32(1)
	if v3053&int32(255) != 0 {
		goto L386
	} else {
		goto L387
	}
L375:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v3070))) = uint8(v3000)
	if v3006&int32(255) == int32(0) {
		v3097 = v3072
		goto L374
	} else {
		goto L378
	}
L376:
	;
	goto L377
L377:
	;
	v3095 = int32(32)
	*(*uint8)(unsafe.Add(mBase, uint32(v3070))) = uint8(v3095)
	v3097 = v3072
	goto L374
L378:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v3070)+1)) = uint8(v3006)
	if v3004&int32(255) == int32(0) {
		goto L379
	} else {
		goto L380
	}
L379:
	;
	v3097 = v3070 + int32(2)
	goto L374
L380:
	;
	goto L381
L381:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v3070)+2)) = uint8(v3004)
	if base.Ui32(v3000) < base.Ui32(int32(16777216)) {
		goto L382
	} else {
		goto L383
	}
L382:
	;
	v3097 = v3070 + int32(3)
	goto L374
L383:
	;
	goto L384
L384:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v3070)+3)) = uint8(v3002)
	v3097 = v3070 + int32(4)
	goto L374
L385:
	;
	v3133 = v41 + int32(176)
	v3136 = v3130 - v3133
	v3142 = int32(255)
	switch v3136 {
	case 0:
		v3194 = v3136
		goto L397
	default:
		goto L398
	case 3:
		goto L399
	}
L386:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v3097))) = uint8(v3053)
	v3104 = int32(base.Ui32(v3053) >> (uint(int32(8)) % 32))
	if v3104&int32(255) == int32(0) {
		v3130 = v3099
		goto L385
	} else {
		goto L389
	}
L387:
	;
	goto L388
L388:
	;
	v3128 = int32(32)
	*(*uint8)(unsafe.Add(mBase, uint32(v3097))) = uint8(v3128)
	v3130 = v3099
	goto L385
L389:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v3097)+1)) = uint8(v3104)
	v3111 = int32(base.Ui32(v3053) >> (uint(int32(16)) % 32))
	if v3111&int32(255) == int32(0) {
		goto L390
	} else {
		goto L391
	}
L390:
	;
	v3130 = v3097 + int32(2)
	goto L385
L391:
	;
	goto L392
L392:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v3097)+2)) = uint8(v3111)
	if base.Ui32(v3053) < base.Ui32(int32(16777216)) {
		goto L393
	} else {
		goto L394
	}
L393:
	;
	v3130 = v3097 + int32(3)
	goto L385
L394:
	;
	goto L395
L395:
	;
	v3124 = int32(base.Ui32(v3053) >> (uint(int32(24)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v3097)+3)) = uint8(v3124)
	v3130 = v3097 + int32(4)
	goto L385
L396:
	;
	v3213 = v3015 + int32(3)
	if v2878 == int32(-4) {
		goto L404
	} else {
		goto L405
	}
L397:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v3015))) = uint16(v3194)
	v3202 = int32(base.Ui32(v3194) >> (uint(int32(16)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v3015)+2)) = uint8(v3202)
	goto L396
L398:
	;
	v3153 = v3133
	v3154 = v3136
	v3156 = v3142
	v3157 = v3142
	v3158 = v3142
	v3159 = v3142
	goto L400
L399:
	;
	v3146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3133))))
	*(*uint8)(unsafe.Add(mBase, uint32(v3015))) = uint8(v3146)
	v3148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3133)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v3015)+1)) = uint8(v3148)
	v3150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3133)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v3015)+2)) = uint8(v3150)
	goto L396
L400:
	;
	v3160 = int32(8)
	v3162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3153))))
	v3168 = *(*int32)(unsafe.Add(mBase, uint32((v3162^v3156)<<(uint(int32(2))%32))+uint32(_c_F_createTrgmNFA[2])))
	v3171 = int32(16)
	v3176 = int32(24)
	v3179 = v3168 ^ (v3157<<(uint(v3160)%32)&int32(_a_F_createTrgmNFA_10) | v3158<<(uint(v3171)%32)&int32(16711680) | v3159<<(uint(v3176)%32))
	v3186 = int32(1)
	v3189 = v3154 - v3186
	if v3189 != 0 {
		v3153 = v3153 + v3186
		v3154 = v3189
		v3156 = int32(base.Ui32(v3179) >> (uint(v3176) % 32))
		v3157 = v3168
		v3158 = int32(base.Ui32(v3179) >> (uint(v3160) % 32))
		v3159 = int32(base.Ui32(v3179) >> (uint(v3171) % 32))
		goto L400
	} else {
		goto L402
	}
L401:
	;
	v3194 = v3179 ^ int32(-1)
	goto L397
L402:
	;
	goto L401
L403:
	;
	goto L363
L404:
	;
	v3220 = int32(1)
	goto L403
L405:
	;
	goto L406
L406:
	;
	v3217 = v3009 + int32(1)
	v3218 = *(*int32)(unsafe.Add(mBase, uint32(v2887)+4))
	if v3217 < v3218 {
		v3009 = v3217
		v3015 = v3213
		goto L362
	} else {
		goto L407
	}
L407:
	;
	v3220 = v3218
	goto L403
L408:
	;
	v3222 = *(*int32)(unsafe.Add(mBase, uint32(v2877)+4))
	v3223 = v3220
	v3227 = v3222
	v3230 = v3213
	goto L358
L409:
	;
	v3268 = v3227
	v3271 = v3230
	goto L355
L410:
	;
	v3302 = *(*int32)(unsafe.Add(mBase, uint32(v2863)+4))
	v3303 = v3302
	v3307 = v3268
	v3310 = v3271
	goto L346
L411:
	;
	v3351 = v3310
	goto L343
L412:
	;
	goto L331
L413:
	;
	v3467 = int32(2)
	v3470 = F_hash_seq_search(m, v41+int32(176))
	mBase = m.M
	v3471 = m.ExcPending
	if v3471 != 0 {
		goto L1
	} else {
		goto L414
	}
L414:
	;
	if v3470 != 0 {
		goto L415
	} else {
		goto L416
	}
L415:
	;
	v3476 = v3470
	v3485 = v3467
	goto L418
L416:
	;
	v3545 = v3467
	goto L417
L417:
	;
	v3570 = *(*int32)(unsafe.Add(mBase, uint32(v41)+96))
	v3573 = F_palloc(m, v3570*int32(12))
	mBase = m.M
	v3574 = m.ExcPending
	if v3574 != 0 {
		goto L1
	} else {
		goto L431
	}
L418:
	;
	v3510 = *(*int32)(unsafe.Add(mBase, uint32(v3476)+28))
	if v3510 != 0 {
		v3476 = v3510
		goto L418
	} else {
		goto L420
	}
L419:
	;
	v3545 = v3527
	goto L417
L420:
	;
	v3511 = *(*int32)(unsafe.Add(mBase, uint32(v3476)+24))
	if int32(0) <= v3511 {
		v3527 = v3485
		goto L421
	} else {
		goto L422
	}
L421:
	;
	v3530 = F_hash_seq_search(m, v41+int32(176))
	mBase = m.M
	v3531 = m.ExcPending
	if v3531 != 0 {
		goto L1
	} else {
		goto L429
	}
L422:
	;
	v3514 = *(*int32)(unsafe.Add(mBase, uint32(v3476)+20))
	if v3514&int32(1) != 0 {
		goto L423
	} else {
		goto L424
	}
L423:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3476)+24)) = int32(0)
	v3527 = v3485
	goto L421
L424:
	;
	goto L425
L425:
	;
	if v3514&int32(2) != 0 {
		goto L426
	} else {
		goto L427
	}
L426:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3476)+24)) = int32(1)
	v3527 = v3485
	goto L421
L427:
	;
	goto L428
L428:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3476)+24)) = v3485
	v3527 = v3485 + int32(1)
	goto L421
L429:
	;
	if v3530 != 0 {
		v3476 = v3530
		v3485 = v3527
		goto L418
	} else {
		goto L430
	}
L430:
	;
	goto L419
L431:
	;
	F_hash_seq_init(m, v41+int32(176), v3442)
	mBase = m.M
	v3578 = m.ExcPending
	if v3578 != 0 {
		goto L1
	} else {
		goto L432
	}
L432:
	;
	v3581 = F_hash_seq_search(m, v41+int32(176))
	mBase = m.M
	v3582 = m.ExcPending
	if v3582 != 0 {
		goto L1
	} else {
		goto L434
	}
L433:
	;
	v3943 = int32(0)
	v3945 = F_MemoryContextAlloc(m, l3, int32(28))
	mBase = m.M
	v3946 = m.ExcPending
	if v3946 != 0 {
		goto L1
	} else {
		goto L471
	}
L434:
	;
	if v3581 != 0 {
		goto L435
	} else {
		goto L436
	}
L435:
	;
	v3584 = *(*int32)(unsafe.Add(mBase, uint32(v41)+108))
	v3585 = *(*int32)(unsafe.Add(mBase, uint32(v41)+104))
	v3586 = v3581
	v3587 = int32(0)
	goto L438
L436:
	;
	goto L437
L437:
	;
	v3899 = int32(0)
	F_pg_qsort(m, v3573, v3899, int32(12), int32(_a_F_createTrgmNFA_11))
	mBase = m.M
	v3904 = m.ExcPending
	if v3904 != 0 {
		goto L1
	} else {
		goto L470
	}
L438:
	;
	v3628 = v3586
	goto L440
L439:
	;
	F_pg_qsort(m, v3573, v3779, int32(12), int32(_a_F_createTrgmNFA_11))
	mBase = m.M
	v3823 = m.ExcPending
	if v3823 != 0 {
		goto L1
	} else {
		goto L458
	}
L440:
	;
	v3662 = *(*int32)(unsafe.Add(mBase, uint32(v3628)+28))
	if v3662 != 0 {
		v3628 = v3662
		goto L440
	} else {
		goto L442
	}
L441:
	;
	v3663 = *(*int32)(unsafe.Add(mBase, uint32(v3586)+12))
	if v3663 == int32(0) {
		v3779 = v3587
		goto L443
	} else {
		goto L444
	}
L442:
	;
	goto L441
L443:
	;
	v3818 = F_hash_seq_search(m, v41+int32(176))
	mBase = m.M
	v3819 = m.ExcPending
	if v3819 != 0 {
		goto L1
	} else {
		goto L456
	}
L444:
	;
	v3666 = int32(0)
	v3667 = *(*int32)(unsafe.Add(mBase, uint32(v3663)+4))
	if v3667 <= v3666 {
		v3779 = v3587
		goto L443
	} else {
		goto L445
	}
L445:
	;
	v3671 = v3587
	v3677 = v3666
	v3679 = v3667
	goto L446
L446:
	;
	v3708 = *(*int32)(unsafe.Add(mBase, uint32(v3663)+12))
	v3712 = *(*int32)(unsafe.Add(mBase, uint32(v3708+v3677<<(uint(int32(2))%32))))
	v3713 = *(*int32)(unsafe.Add(mBase, uint32(v3712)+12))
	v3718 = v3713
	goto L448
L447:
	;
	v3779 = v3772
	goto L443
L448:
	;
	v3752 = *(*int32)(unsafe.Add(mBase, uint32(v3718)+28))
	if v3752 != 0 {
		v3718 = v3752
		goto L448
	} else {
		goto L450
	}
L449:
	;
	v3753 = *(*int32)(unsafe.Add(mBase, uint32(v3628)+24))
	v3754 = *(*int32)(unsafe.Add(mBase, uint32(v3718)+24))
	if v3753 != v3754 {
		goto L451
	} else {
		goto L452
	}
L450:
	;
	goto L449
L451:
	;
	v3758 = F_bsearch(m, v3712, v3585, v3584, int32(32), int32(_a_F_createTrgmNFA_5))
	mBase = m.M
	v3759 = m.ExcPending
	if v3759 != 0 {
		goto L1
	} else {
		goto L454
	}
L452:
	;
	v3772 = v3671
	v3774 = v3679
	goto L453
L453:
	;
	v3776 = v3677 + int32(1)
	if v3776 < v3774 {
		v3671 = v3772
		v3677 = v3776
		v3679 = v3774
		goto L446
	} else {
		goto L455
	}
L454:
	;
	v3762 = v3573 + v3671*int32(12)
	v3763 = *(*int32)(unsafe.Add(mBase, uint32(v3628)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v3762))) = v3763
	v3765 = *(*int32)(unsafe.Add(mBase, uint32(v3718)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v3762)+4)) = v3765
	v3767 = *(*int32)(unsafe.Add(mBase, uint32(v3758)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v3762)+8)) = v3767
	v3769 = *(*int32)(unsafe.Add(mBase, uint32(v3663)+4))
	v3772 = v3671 + int32(1)
	v3774 = v3769
	goto L453
L455:
	;
	goto L447
L456:
	;
	if v3818 != 0 {
		v3586 = v3818
		v3587 = v3779
		goto L438
	} else {
		goto L457
	}
L457:
	;
	goto L439
L458:
	;
	if v3779 < int32(2) {
		v3906 = v3779
		goto L433
	} else {
		goto L459
	}
L459:
	;
	v3826 = int32(12)
	__phi3831 = v3573
	__phi3835 = v3573
	__phi3838 = v3573 + v3826
	v3831 = __phi3831
	v3835 = __phi3835
	v3838 = __phi3838
	goto L460
L460:
	;
	v3869 = *(*int32)(unsafe.Add(mBase, uint32(v3835)+12))
	v3870 = *(*int32)(unsafe.Add(mBase, uint32(v3831)))
	if v3869 < v3870 {
		v3888 = v3831
		goto L462
	} else {
		goto L463
	}
L461:
	;
	v3896 = base.I32_div_s(v3888-v3573, int32(12))
	v3906 = v3896 + int32(1)
	goto L433
L462:
	;
	v3892 = v3838 + int32(12)
	if base.Ui32(v3892) < base.Ui32(v3573+v3779*v3826) {
		__phi3831 = v3888
		__phi3835 = v3838
		__phi3838 = v3892
		v3831 = __phi3831
		v3835 = __phi3835
		v3838 = __phi3838
		goto L460
	} else {
		goto L469
	}
L463:
	;
	if v3870 < v3869 {
		goto L464
	} else {
		goto L465
	}
L464:
	;
	v3882 = *(*int64)(unsafe.Add(mBase, uint32(v3838)))
	*(*int64)(unsafe.Add(mBase, uint32(v3831)+12)) = v3882
	v3884 = *(*int32)(unsafe.Add(mBase, uint32(v3838)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v3831)+20)) = v3884
	v3888 = v3831 + int32(12)
	goto L462
L465:
	;
	v3873 = *(*int32)(unsafe.Add(mBase, uint32(v3835)+20))
	v3874 = *(*int32)(unsafe.Add(mBase, uint32(v3831)+8))
	if v3873 < v3874 {
		v3888 = v3831
		goto L462
	} else {
		goto L466
	}
L466:
	;
	if v3874 < v3873 {
		goto L464
	} else {
		goto L467
	}
L467:
	;
	v3877 = *(*int32)(unsafe.Add(mBase, uint32(v3835)+16))
	v3878 = *(*int32)(unsafe.Add(mBase, uint32(v3831)+4))
	if v3877 <= v3878 {
		v3888 = v3831
		goto L462
	} else {
		goto L468
	}
L468:
	;
	goto L464
L469:
	;
	goto L461
L470:
	;
	v3906 = v3899
	goto L433
L471:
	;
	v3947 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v3945))) = v3947
	v3949 = *(*int32)(unsafe.Add(mBase, uint32(v41)+108))
	if v3949 <= v3947 {
		goto L473
	} else {
		goto L474
	}
L472:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3945)+8)) = v3545
	v4343 = F_MemoryContextAlloc(m, l3, v3545<<(uint(int32(3))%32))
	mBase = m.M
	v4344 = m.ExcPending
	if v4344 != 0 {
		goto L1
	} else {
		goto L519
	}
L473:
	;
	v3953 = F_MemoryContextAlloc(m, l3, int32(0))
	mBase = m.M
	v3954 = m.ExcPending
	if v3954 != 0 {
		goto L1
	} else {
		goto L476
	}
L474:
	;
	goto L475
L475:
	;
	v3957 = v3949 & int32(3)
	v3958 = *(*int32)(unsafe.Add(mBase, uint32(v41)+104))
	v3959 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v3949) {
		goto L477
	} else {
		goto L478
	}
L476:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3945)+4)) = v3953
	goto L472
L477:
	;
	v3965 = v3943
	v3969 = v3959
	v3973 = int32(0)
	goto L480
L478:
	;
	v4039 = v3943
	v4043 = v3959
	goto L479
L479:
	;
	if v3957 != 0 {
		goto L495
	} else {
		goto L496
	}
L480:
	;
	v4005 = v3958 + v3969<<(uint(int32(5))%32)
	v4006 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4005)+24)))
	if v4006 == int32(1) {
		goto L482
	} else {
		goto L483
	}
L481:
	;
	v4039 = v4033
	v4043 = v4035
	goto L479
L482:
	;
	v4010 = v3965 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v3945))) = v4010
	v4012 = v4010
	goto L484
L483:
	;
	v4012 = v3965
	goto L484
L484:
	;
	v4013 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4005)+56)))
	if v4013 == int32(1) {
		goto L485
	} else {
		goto L486
	}
L485:
	;
	v4017 = v4012 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v3945))) = v4017
	v4019 = v4017
	goto L487
L486:
	;
	v4019 = v4012
	goto L487
L487:
	;
	v4020 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4005)+88)))
	if v4020 == int32(1) {
		goto L488
	} else {
		goto L489
	}
L488:
	;
	v4024 = v4019 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v3945))) = v4024
	v4026 = v4024
	goto L490
L489:
	;
	v4026 = v4019
	goto L490
L490:
	;
	v4027 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4005)+120)))
	if v4027 == int32(1) {
		goto L491
	} else {
		goto L492
	}
L491:
	;
	v4031 = v4026 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v3945))) = v4031
	v4033 = v4031
	goto L493
L492:
	;
	v4033 = v4026
	goto L493
L493:
	;
	v4034 = int32(4)
	v4035 = v3969 + v4034
	v4037 = v3973 + v4034
	if v4037 != v3949&int32(2147483644) {
		v3965 = v4033
		v3969 = v4035
		v3973 = v4037
		goto L480
	} else {
		goto L494
	}
L494:
	;
	goto L481
L495:
	;
	v4078 = v4039
	v4082 = v4043
	v4084 = int32(0)
	goto L498
L496:
	;
	v4131 = v4039
	goto L497
L497:
	;
	v4171 = F_MemoryContextAlloc(m, l3, v4131<<(uint(int32(2))%32))
	mBase = m.M
	v4172 = m.ExcPending
	if v4172 != 0 {
		goto L1
	} else {
		goto L504
	}
L498:
	;
	v4119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3958+v4082<<(uint(int32(5))%32))+24)))
	if v4119 == int32(1) {
		goto L500
	} else {
		goto L501
	}
L499:
	;
	v4131 = v4125
	goto L497
L500:
	;
	v4123 = v4078 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v3945))) = v4123
	v4125 = v4123
	goto L502
L501:
	;
	v4125 = v4078
	goto L502
L502:
	;
	v4126 = int32(1)
	v4129 = v4084 + v4126
	if v4129 != v3957 {
		v4078 = v4125
		v4082 = v4082 + v4126
		v4084 = v4129
		goto L498
	} else {
		goto L503
	}
L503:
	;
	goto L499
L504:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3945)+4)) = v4171
	v4174 = int32(1)
	v4176 = int32(0)
	if v3949 != v4174 {
		goto L505
	} else {
		goto L506
	}
L505:
	;
	v4183 = v4176
	v4187 = v4176
	v4191 = int32(0)
	goto L508
L506:
	;
	v4251 = v4176
	v4255 = v4176
	goto L507
L507:
	;
	if v3949&v4174 == int32(0) {
		goto L472
	} else {
		goto L517
	}
L508:
	;
	v4223 = v3958 + v4187<<(uint(int32(5))%32)
	v4224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4223)+24)))
	if v4224 == int32(1) {
		goto L510
	} else {
		goto L511
	}
L509:
	;
	v4251 = v4245
	v4255 = v4247
	goto L507
L510:
	;
	v4230 = *(*int32)(unsafe.Add(mBase, uint32(v4223)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v4171+v4183<<(uint(int32(2))%32)))) = v4230
	v4234 = v4183 + int32(1)
	goto L512
L511:
	;
	v4234 = v4183
	goto L512
L512:
	;
	v4235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4223)+56)))
	if v4235 == int32(1) {
		goto L513
	} else {
		goto L514
	}
L513:
	;
	v4241 = *(*int32)(unsafe.Add(mBase, uint32(v4223)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v4171+v4234<<(uint(int32(2))%32)))) = v4241
	v4245 = v4234 + int32(1)
	goto L515
L514:
	;
	v4245 = v4234
	goto L515
L515:
	;
	v4246 = int32(2)
	v4247 = v4187 + v4246
	v4249 = v4191 + v4246
	if v4249 != v3949&int32(2147483646) {
		v4183 = v4245
		v4187 = v4247
		v4191 = v4249
		goto L508
	} else {
		goto L516
	}
L516:
	;
	goto L509
L517:
	;
	v4293 = v3958 + v4255<<(uint(int32(5))%32)
	v4294 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4293)+24)))
	if v4294 != int32(1) {
		goto L472
	} else {
		goto L518
	}
L518:
	;
	v4300 = *(*int32)(unsafe.Add(mBase, uint32(v4293)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v4171+v4251<<(uint(int32(2))%32)))) = v4300
	goto L472
L519:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3945)+12)) = v4343
	v4348 = F_MemoryContextAlloc(m, l3, v3906<<(uint(int32(3))%32))
	mBase = m.M
	v4349 = m.ExcPending
	if v4349 != 0 {
		goto L1
	} else {
		goto L520
	}
L520:
	;
	if int32(0) < v3545 {
		goto L521
	} else {
		goto L522
	}
L521:
	;
	v4352 = *(*int32)(unsafe.Add(mBase, uint32(v3945)+12))
	v4353 = int32(0)
	v4359 = v4353
	v4363 = v4353
	goto L524
L522:
	;
	goto L523
L523:
	;
	v4539 = *(*int32)(unsafe.Add(mBase, uint32(v3945)))
	v4540 = F_MemoryContextAlloc(m, l3, v4539)
	mBase = m.M
	v4541 = m.ExcPending
	if v4541 != 0 {
		goto L1
	} else {
		goto L535
	}
L524:
	;
	v4393 = int32(3)
	v4395 = v4352 + v4363<<(uint(v4393)%32)
	*(*int32)(unsafe.Add(mBase, uint32(v4395)+4)) = v4348 + v4359<<(uint(v4393)%32)
	if v3906 <= v4359 {
		goto L527
	} else {
		goto L528
	}
L525:
	;
	goto L523
L526:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4395))) = v4465
	v4499 = v4363 + int32(1)
	if v4499 != v3545 {
		v4359 = v4463
		v4363 = v4499
		goto L524
	} else {
		goto L534
	}
L527:
	;
	v4463 = v4359
	v4465 = int32(0)
	goto L526
L528:
	;
	goto L529
L529:
	;
	v4402 = v3906 - v4359
	v4408 = v4359
	v4410 = int32(0)
	goto L530
L530:
	;
	v4444 = v3573 + v4408*int32(12)
	v4445 = *(*int32)(unsafe.Add(mBase, uint32(v4444)))
	if v4445 != v4363 {
		v4463 = v4408
		v4465 = v4410
		goto L526
	} else {
		goto L532
	}
L531:
	;
	v4463 = v3906
	v4465 = v4402
	goto L526
L532:
	;
	v4449 = v4348 + v4408<<(uint(int32(3))%32)
	v4450 = *(*int32)(unsafe.Add(mBase, uint32(v4444)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v4449))) = v4450
	v4452 = *(*int32)(unsafe.Add(mBase, uint32(v4444)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v4449)+4)) = v4452
	v4454 = int32(1)
	v4457 = v4410 + v4454
	if v4457 != v4402 {
		v4408 = v4408 + v4454
		v4410 = v4457
		goto L530
	} else {
		goto L533
	}
L533:
	;
	goto L531
L534:
	;
	goto L525
L535:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3945)+16)) = v4540
	v4543 = *(*int32)(unsafe.Add(mBase, uint32(v3945)+8))
	v4544 = F_MemoryContextAlloc(m, l3, v4543)
	mBase = m.M
	v4545 = m.ExcPending
	if v4545 != 0 {
		goto L1
	} else {
		goto L536
	}
L536:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3945)+20)) = v4544
	v4547 = *(*int32)(unsafe.Add(mBase, uint32(v3945)+8))
	v4550 = F_MemoryContextAlloc(m, l3, v4547<<(uint(int32(2))%32))
	mBase = m.M
	v4551 = m.ExcPending
	if v4551 != 0 {
		goto L1
	} else {
		goto L537
	}
L537:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3945)+24)) = v4550
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v3945
	v4572 = v2793
	goto L21
L538:
	;
	m.G0 = v41 + int32(240)
	return v4572
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
									*(*int32)(unsafe.Add(mBase, uint32(v56))) = int32(388)
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
		*(*int64)(unsafe.Add(mBase, uint32(v8))) = int64(1455993913623)
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
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
	var v125 int32
	_ = v125
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
	var v243 float64
	_ = v243
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
	var v267 float64
	_ = v267
	var v275 float64
	_ = v275
	var v276 int32
	_ = v276
	var v277 float64
	_ = v277
	var v279 float64
	_ = v279
	var v280 float64
	_ = v280
	var v284 float64
	_ = v284
	var v286 float64
	_ = v286
	var v288 float64
	_ = v288
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
	*(*int64)(unsafe.Add(mBase, uint32(v19))) = int64(1481763717405)
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
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
	v183 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
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
	*(*int64)(unsafe.Add(mBase, uint32(v42)+24)) = v111
	*(*int64)(unsafe.Add(mBase, uint32(v42)+16)) = v111
	*(*int32)(unsafe.Add(mBase, uint32(v42)+8)) = l0
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
	v125 = v5
	goto L22
L22:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v135+v125<<(uint(int32(2))%32))))
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
	v145 = v125 + int32(1)
	v146 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v145 < v146 {
		v125 = v145
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
	v280 = *(*float64)(unsafe.Add(mBase, uint32(v276)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+40)) = int32(0)
	v284 = float64(0)
	v286 = base.F64_add(v280, base.F64_add(base.F64_add(v178, v275), v284))
	*(*float64)(unsafe.Add(mBase, uint32(v19)+48)) = v286
	v288 = *(*float64)(unsafe.Add(mBase, uint32(v19)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v19)+56)) = base.F64_add(v286, base.F64_add(base.F64_mul(v277, v288), base.F64_add(base.F64_mul(base.F64_sub(base.F64_add(v267, v279), v178), v177), base.F64_add(base.F64_mul(v187, v177), v284))))
	m.G0 = v42 + int32(32)
	return v19
L28:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v28)+16))
	v189 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v42)+24)) = v189
	*(*int64)(unsafe.Add(mBase, uint32(v42)+16)) = v189
	*(*int32)(unsafe.Add(mBase, uint32(v42)+8)) = l0
	if v188 == int32(0) {
		v243 = v10
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
	v256 = *(*float64)(unsafe.Add(mBase, uint32(l1)+200))
	v257 = *(*float64)(unsafe.Add(mBase, uint32(l1)+192))
	v267 = v256
	v275 = v257
	goto L27
L31:
	;
	v252 = *(*float64)(unsafe.Add(mBase, uint32(l1)+200))
	v254 = *(*float64)(unsafe.Add(mBase, uint32(l1)+192))
	v267 = base.F64_add(v243, v252)
	v275 = base.F64_add(v251, v254)
	goto L27
L32:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v188)+4))
	if v198 <= int32(0) {
		v243 = v10
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
	v243 = v232
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
}
func F_create_upper_unique_path(m *base.Module, l0 int32, l1 int32, l2 int32, l3 float64) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v41 float64
	_ = v41
	var v43 float64
	_ = v43
	var v44 float64
	_ = v44
	var v46 float64
	_ = v46
	v11 = F_palloc0(m, int32(80))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = l0
		*(*int64)(unsafe.Add(mBase, uint32(v11))) = int64(1576252997938)
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
		v19 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v11)+20)) = uint8(v19)
		*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v19
		*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v18
		v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+26)))
		if v24 == int32(1) {
			v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
			v29 = v27
		} else {
			v29 = int32(0)
		}
		v31 = v29 & int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(v11)+21)) = uint8(v31)
		v33 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
		*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = v33
		v35 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
		*(*int32)(unsafe.Add(mBase, uint32(v11)+76)) = l2
		*(*int32)(unsafe.Add(mBase, uint32(v11)+72)) = l1
		*(*int32)(unsafe.Add(mBase, uint32(v11)+64)) = v35
		v39 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
		*(*int32)(unsafe.Add(mBase, uint32(v11)+40)) = v39
		v41 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
		*(*float64)(unsafe.Add(mBase, uint32(v11)+48)) = v41
		v43 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
		v44 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
		v46 = *(*float64)(unsafe.Add(mBase, _c_F_create_upper_unique_path[0]))
		*(*float64)(unsafe.Add(mBase, uint32(v11)+32)) = l3
		*(*float64)(unsafe.Add(mBase, uint32(v11)+56)) = base.F64_add(v43, base.F64_mul(base.F64_mul(v46, v44), base.F64_convert_i32_s(l2)))
		return v11
	}
}
func F_createdb_failure_callback(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v3 == int32(0) {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
		F_DropDatabaseBuffers(m, v6)
		mBase = m.M
		v8 = m.ExcPending
		if v8 != 0 {
			return
		} else {
			v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			F_ForgetDatabaseSyncRequests(m, v9)
			mBase = m.M
			v11 = m.ExcPending
			if v11 != 0 {
				return
			} else {
				v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
				F_UnlockSharedObject(m, int32(1262), v13, int32(1))
				mBase = m.M
				v16 = m.ExcPending
				if v16 != 0 {
					return
				} else {
					v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
					F_UnlockSharedObject(m, int32(1262), v18, int32(5))
					mBase = m.M
					v21 = m.ExcPending
					if v21 != 0 {
						return
					} else {
						v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
						F_remove_dbtablespaces(m, v22)
						mBase = m.M
						v24 = m.ExcPending
						if v24 != 0 {
							return
						} else {
							return
						}
					}
				}
			}
		}
	} else {
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
		F_UnlockSharedObject(m, int32(1262), v18, int32(5))
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return
		} else {
			v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			F_remove_dbtablespaces(m, v22)
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
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
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v90 int32
	_ = v90
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v172 int32
	_ = v172
	var v177 int32
	_ = v177
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
			v30 = int32(4)
			v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
			if v32&int32(254) == int32(2) {
				v41 = v30
			} else {
				v41 = base.B2i32(v32 == int32(18)) << (uint(v30) % 32)
			}
			if v32 == int32(1) {
				v44 = v30
			} else {
				v44 = v41
			}
			v57 = v44
		} else {
			v45 = int32(1)
			if v27&v45 != 0 {
				v57 = int32(base.Ui32(v27)>>(uint(v45)%32)) - v45
			} else {
				v51 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				v57 = int32(base.Ui32(v51)>>(uint(int32(2))%32)) - int32(4)
			}
		}
		v58 = *(*int32)(unsafe.Add(mBase, uint32(v15)+uint32(_c_F_cryptohash_internal[1])))
		v59 = F_pg_cryptohash_create(m, l0)
		mBase = m.M
		v60 = m.ExcPending
		if v60 != 0 {
			return int32(0)
		} else {
			v61 = F_pg_cryptohash_init(m, v59)
			mBase = m.M
			if int32(0) <= v61 {
				v64 = int32(1)
				if v27&v64 != 0 {
					v68 = v64
				} else {
					v68 = int32(4)
				}
				v70 = F_pg_cryptohash_update(m, v59, l1+v68, v57)
				mBase = m.M
				if v70 < int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v119 = m.ExcPending
					if v119 != 0 {
						return int32(0)
					} else {
						if v59 == int32(0) {
							v134 = int32(_a_F_cryptohash_internal_0)
						} else {
							v126 = *(*int32)(unsafe.Add(mBase, uint32(v59)+4))
							if v126 == int32(1) {
								v129 = int32(_a_F_cryptohash_internal_1)
							} else {
								v129 = int32(_a_F_cryptohash_internal_2)
							}
							if v126 == int32(2) {
								v132 = int32(_a_F_cryptohash_internal_0)
							} else {
								v132 = v129
							}
							v134 = v132
						}
						*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v134
						*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v58
						F_errmsg_internal(m, int32(_a_F_cryptohash_internal_3), v12+int32(16))
						mBase = m.M
						v141 = m.ExcPending
						if v141 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_cryptohash_internal_4), int32(123), int32(_a_F_cryptohash_internal_5))
							mBase = m.M
							v146 = m.ExcPending
							if v146 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v75 = F_pg_cryptohash_final(m, v59, v21+int32(4), v18)
					mBase = m.M
					if v75 < int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v150 = m.ExcPending
						if v150 != 0 {
							return int32(0)
						} else {
							if v59 == int32(0) {
								v165 = int32(_a_F_cryptohash_internal_0)
							} else {
								v157 = *(*int32)(unsafe.Add(mBase, uint32(v59)+4))
								if v157 == int32(1) {
									v160 = int32(_a_F_cryptohash_internal_1)
								} else {
									v160 = int32(_a_F_cryptohash_internal_2)
								}
								if v157 == int32(2) {
									v163 = int32(_a_F_cryptohash_internal_0)
								} else {
									v163 = v160
								}
								v165 = v163
							}
							*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = v165
							*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v58
							F_errmsg_internal(m, int32(_a_F_cryptohash_internal_6), v12+int32(32))
							mBase = m.M
							v172 = m.ExcPending
							if v172 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_cryptohash_internal_4), int32(127), int32(_a_F_cryptohash_internal_5))
								mBase = m.M
								v177 = m.ExcPending
								if v177 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						F_pg_cryptohash_free(m, v59)
						mBase = m.M
						v79 = m.ExcPending
						if v79 != 0 {
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
				v90 = m.ExcPending
				if v90 != 0 {
					return int32(0)
				} else {
					if v59 == int32(0) {
						v105 = int32(_a_F_cryptohash_internal_0)
					} else {
						v97 = *(*int32)(unsafe.Add(mBase, uint32(v59)+4))
						if v97 == int32(1) {
							v100 = int32(_a_F_cryptohash_internal_1)
						} else {
							v100 = int32(_a_F_cryptohash_internal_2)
						}
						if v97 == int32(2) {
							v103 = int32(_a_F_cryptohash_internal_0)
						} else {
							v103 = v100
						}
						v105 = v103
					}
					*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v105
					*(*int32)(unsafe.Add(mBase, uint32(v12))) = v58
					F_errmsg_internal(m, int32(_a_F_cryptohash_internal_7), v12)
					mBase = m.M
					v110 = m.ExcPending
					if v110 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_cryptohash_internal_4), int32(120), int32(_a_F_cryptohash_internal_5))
						mBase = m.M
						v115 = m.ExcPending
						if v115 != 0 {
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
