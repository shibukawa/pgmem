package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_network_broadcast(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v73 int32
	_ = v73
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v15 = F_pg_detoast_datum_packed(m, v14)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		v20 = F_palloc0(m, int32(22))
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return int32(0)
		} else {
			v22 = int32(1)
			v23 = v20 + v22
			v25 = v20 + int32(4)
			v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
			if v26&v22 != 0 {
				v29 = v23
			} else {
				v29 = v25
			}
			v31 = v29 + int32(2)
			v32 = int32(1)
			v33 = v15 + v32
			v35 = v15 + int32(4)
			v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
			v38 = v36 & v32
			if v38 != 0 {
				v39 = v33
			} else {
				v39 = v35
			}
			v41 = v39 + int32(2)
			v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+1)))
			v43 = int32(4)
			if v38 != 0 {
				v47 = int32(1)
			} else {
				v47 = v43
			}
			v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15+v47))))
			if v49 == int32(2) {
				v52 = v43
			} else {
				v52 = int32(16)
			}
			v54 = int32(0)
			v56 = v42
			for {
				if base.Ui32(int32(8)) <= base.Ui32(v56) {
					v79 = v56 - int32(8)
					v80 = int32(0)
				} else {
					v73 = int32(0)
					if v56 == v73 {
						v79 = v73
						v80 = int32(255)
					} else {
						v79 = v73
						v80 = int32(base.Ui32(int32(255)) >> (uint(v56) % 32))
					}
				}
				v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54+v41))))
				v83 = v80 | v82
				*(*uint8)(unsafe.Add(mBase, uint32(v54+v31))) = uint8(v83)
				v86 = v54 | int32(1)
				if base.Ui32(v79) <= base.Ui32(int32(7)) {
					v90 = int32(0)
					if v79 == v90 {
						v99 = v90
						v100 = int32(255)
					} else {
						v99 = v90
						v100 = int32(base.Ui32(int32(255)) >> (uint(v79) % 32))
					}
				} else {
					v99 = v79 - int32(8)
					v100 = int32(0)
				}
				v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41+v86))))
				v103 = v100 | v102
				*(*uint8)(unsafe.Add(mBase, uint32(v31+v86))) = uint8(v103)
				v106 = v54 + int32(2)
				if v106 != v52 {
					v54 = v106
					v56 = v99
					continue
				} else {
					break
				}
				break
			}
			v108 = int32(1)
			v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
			v112 = v110 & v108
			if v112 != 0 {
				v113 = v108
			} else {
				v113 = int32(4)
			}
			v115 = int32(1)
			v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
			if v117&v115 != 0 {
				v120 = v115
			} else {
				v120 = int32(4)
			}
			v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15+v120))))
			*(*uint8)(unsafe.Add(mBase, uint32(v20+v113))) = uint8(v122)
			if v112 != 0 {
				v124 = v23
			} else {
				v124 = v25
			}
			v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
			if v125&int32(1) != 0 {
				v128 = v33
			} else {
				v128 = v35
			}
			v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128)+1)))
			*(*uint8)(unsafe.Add(mBase, uint32(v124)+1)) = uint8(v129)
			v133 = int32(1)
			v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
			if v135&v133 != 0 {
				v138 = v133
			} else {
				v138 = int32(4)
			}
			v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20+v138))))
			if v140 == int32(2) {
				v143 = int32(40)
			} else {
				v143 = int32(88)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v20))) = v143
			return v20
		}
	}
}
func F_network_in(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v150 int32
	_ = v150
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
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
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v268 int32
	_ = v268
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v335 int32
	_ = v335
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v355 int32
	_ = v355
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v362 int32
	_ = v362
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v372 int32
	_ = v372
	var v374 int32
	_ = v374
	var v379 int32
	_ = v379
	var v381 int32
	_ = v381
	var v384 int32
	_ = v384
	var v386 int32
	_ = v386
	var v389 int32
	_ = v389
	var v401 int32
	_ = v401
	var v403 int32
	_ = v403
	var v406 int32
	_ = v406
	var v408 int32
	_ = v408
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v449 int32
	_ = v449
	var v467 int32
	_ = v467
	var v473 int32
	_ = v473
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v521 int32
	_ = v521
	var v526 int32
	_ = v526
	var v540 int32
	_ = v540
	var v552 int32
	_ = v552
	var v554 int32
	_ = v554
	var v556 int32
	_ = v556
	var v559 int32
	_ = v559
	var v564 int32
	_ = v564
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v569 int32
	_ = v569
	var v583 int32
	_ = v583
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v635 int32
	_ = v635
	var v645 int32
	_ = v645
	var v647 int32
	_ = v647
	var v651 int32
	_ = v651
	var v652 int32
	_ = v652
	var v655 int32
	_ = v655
	var v659 int32
	_ = v659
	var v660 int32
	_ = v660
	var v662 int32
	_ = v662
	var v665 int32
	_ = v665
	var v667 int32
	_ = v667
	var v672 int32
	_ = v672
	var v674 int32
	_ = v674
	var v679 int32
	_ = v679
	var v681 int32
	_ = v681
	var v684 int32
	_ = v684
	var v686 int32
	_ = v686
	var v689 int32
	_ = v689
	var v701 int32
	_ = v701
	var v703 int32
	_ = v703
	var v708 int32
	_ = v708
	var v709 int32
	_ = v709
	var v715 int32
	_ = v715
	var v720 int32
	_ = v720
	var v721 int32
	_ = v721
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v724 int32
	_ = v724
	var v726 int32
	_ = v726
	var v729 int32
	_ = v729
	var v731 int32
	_ = v731
	var v732 int32
	_ = v732
	var v733 int32
	_ = v733
	var v734 int32
	_ = v734
	var v736 int32
	_ = v736
	var v742 int32
	_ = v742
	var v744 int32
	_ = v744
	var v760 int32
	_ = v760
	var v762 int32
	_ = v762
	var v763 int32
	_ = v763
	var v766 int32
	_ = v766
	var v781 int32
	_ = v781
	var v782 int32
	_ = v782
	var v783 int32
	_ = v783
	var v809 int32
	_ = v809
	var v811 int32
	_ = v811
	var v814 int32
	_ = v814
	var v816 int32
	_ = v816
	var v817 int32
	_ = v817
	var v818 int32
	_ = v818
	var v821 int32
	_ = v821
	var v833 int32
	_ = v833
	var v843 int32
	_ = v843
	var v845 int32
	_ = v845
	var v849 int32
	_ = v849
	var v850 int32
	_ = v850
	var v853 int32
	_ = v853
	var v857 int32
	_ = v857
	var v858 int32
	_ = v858
	var v860 int32
	_ = v860
	var v863 int32
	_ = v863
	var v865 int32
	_ = v865
	var v870 int32
	_ = v870
	var v872 int32
	_ = v872
	var v877 int32
	_ = v877
	var v879 int32
	_ = v879
	var v882 int32
	_ = v882
	var v884 int32
	_ = v884
	var v887 int32
	_ = v887
	var v899 int32
	_ = v899
	var v901 int32
	_ = v901
	var v904 int32
	_ = v904
	var v907 int32
	_ = v907
	var v911 int32
	_ = v911
	var v912 int32
	_ = v912
	var v922 int32
	_ = v922
	var v923 int32
	_ = v923
	var v926 int32
	_ = v926
	var v928 int32
	_ = v928
	var v937 int32
	_ = v937
	var v947 int32
	_ = v947
	var v948 int32
	_ = v948
	var v976 int32
	_ = v976
	var v978 int32
	_ = v978
	var v979 int32
	_ = v979
	var v982 int32
	_ = v982
	var v992 int32
	_ = v992
	var v993 int32
	_ = v993
	var v1005 int32
	_ = v1005
	var v1006 int32
	_ = v1006
	var v1008 int32
	_ = v1008
	var v1022 int32
	_ = v1022
	var v1036 int32
	_ = v1036
	var v1038 int32
	_ = v1038
	var v1041 int32
	_ = v1041
	var v1043 int32
	_ = v1043
	var v1044 int32
	_ = v1044
	var v1045 int32
	_ = v1045
	var v1048 int32
	_ = v1048
	var v1060 int32
	_ = v1060
	var v1070 int32
	_ = v1070
	var v1072 int32
	_ = v1072
	var v1076 int32
	_ = v1076
	var v1077 int32
	_ = v1077
	var v1080 int32
	_ = v1080
	var v1084 int32
	_ = v1084
	var v1085 int32
	_ = v1085
	var v1087 int32
	_ = v1087
	var v1090 int32
	_ = v1090
	var v1092 int32
	_ = v1092
	var v1097 int32
	_ = v1097
	var v1099 int32
	_ = v1099
	var v1104 int32
	_ = v1104
	var v1106 int32
	_ = v1106
	var v1109 int32
	_ = v1109
	var v1111 int32
	_ = v1111
	var v1114 int32
	_ = v1114
	var v1126 int32
	_ = v1126
	var v1128 int32
	_ = v1128
	var v1131 int32
	_ = v1131
	var v1133 int32
	_ = v1133
	var v1135 int32
	_ = v1135
	var v1136 int32
	_ = v1136
	var v1148 int32
	_ = v1148
	var v1179 int32
	_ = v1179
	var v1193 int32
	_ = v1193
	var v1194 int32
	_ = v1194
	var v1197 int32
	_ = v1197
	var v1199 int32
	_ = v1199
	var v1206 int32
	_ = v1206
	var v1210 int32
	_ = v1210
	var v1211 int32
	_ = v1211
	var v1233 int32
	_ = v1233
	var v1236 int32
	_ = v1236
	var v1247 int32
	_ = v1247
	var v1249 int32
	_ = v1249
	var v1252 int32
	_ = v1252
	var v1265 int32
	_ = v1265
	var v1299 int32
	_ = v1299
	var v1332 int32
	_ = v1332
	var v1337 int32
	_ = v1337
	var v1339 int32
	_ = v1339
	var v1341 int32
	_ = v1341
	var v1342 int32
	_ = v1342
	var v1344 int32
	_ = v1344
	var v1347 int32
	_ = v1347
	var v1352 int32
	_ = v1352
	var v1353 int32
	_ = v1353
	var v1359 int32
	_ = v1359
	var v1363 int32
	_ = v1363
	var v1367 int32
	_ = v1367
	var v1372 int32
	_ = v1372
	var v1374 int32
	_ = v1374
	var v1379 int32
	_ = v1379
	var v1384 int32
	_ = v1384
	var v1387 int32
	_ = v1387
	var v1389 int32
	_ = v1389
	var v1396 int32
	_ = v1396
	var v1399 int32
	_ = v1399
	var v1417 int32
	_ = v1417
	var v1419 int32
	_ = v1419
	var v1439 int32
	_ = v1439
	var v1440 int32
	_ = v1440
	var v1441 int32
	_ = v1441
	var v1446 int32
	_ = v1446
	var v1452 int32
	_ = v1452
	var v1456 int32
	_ = v1456
	var v1461 int32
	_ = v1461
	var v1483 int32
	_ = v1483
	var v1485 int32
	_ = v1485
	var v1488 int32
	_ = v1488
	var v1490 int32
	_ = v1490
	var v1493 int32
	_ = v1493
	var v1506 int32
	_ = v1506
	v19 = m.G0
	v21 = v19 - int32(32)
	m.G0 = v21
	v24 = F_palloc0(m, int32(22))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v28 = int32(1)
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24))))
	v32 = v30 & v28
	if v32 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v33 = v28
	goto L5
L4:
	;
	v33 = int32(4)
	goto L5
L5:
	;
	v37 = int32(58)
	v38 = F___strchrnul(m, l0, v37)
	mBase = m.M
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38))))
	if v40 == v37 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	if v44 != 0 {
		goto L10
	} else {
		goto L11
	}
L7:
	;
	v44 = v38
	goto L9
L8:
	;
	v44 = int32(0)
	goto L9
L9:
	;
	goto L6
L10:
	;
	v45 = int32(3)
	goto L12
L11:
	;
	v45 = int32(2)
	goto L12
L12:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v24+v33))) = uint8(v45)
	v48 = v24 + int32(1)
	v50 = v24 + int32(4)
	if v32 != 0 {
		goto L16
	} else {
		goto L17
	}
L13:
	;
	m.G0 = v21 + int32(32)
	return v1506
L14:
	;
	if v1341 != 0 {
		goto L320
	} else {
		goto L321
	}
L15:
	;
	if int32(0) <= v1332 {
		goto L300
	} else {
		goto L301
	}
L16:
	;
	v51 = v48
	goto L18
L17:
	;
	v51 = v50
	goto L18
L18:
	;
	v52 = int32(2)
	v53 = v51 + v52
	if v45 == v52 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v58 = int32(4)
	goto L21
L20:
	;
	v58 = int32(16)
	goto L21
L21:
	;
	if l1 != 0 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v60 = v58
	goto L24
L23:
	;
	v60 = int32(-1)
	goto L24
L24:
	;
	switch v45 - int32(2) {
	case 0:
		goto L35
	case 1:
		goto L34
	default:
		goto L33
	}
L25:
	;
	v1332 = v1299
	goto L15
L26:
	;
	v1299 = int32(-1)
	goto L25
L27:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_network_in[0])) = v1265
	goto L26
L28:
	;
	if v978 == int32(47) {
		goto L239
	} else {
		goto L240
	}
L29:
	;
	m.Env.X__assert_fail(m, int32(_a_F_network_in_0), int32(_a_F_network_in_1), int32(150), int32(_a_F_network_in_2))
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L30:
	;
	m.Env.X__assert_fail(m, int32(_a_F_network_in_3), int32(_a_F_network_in_1), int32(121), int32(_a_F_network_in_2))
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L31:
	;
	m.Env.X__assert_fail(m, int32(_a_F_network_in_0), int32(_a_F_network_in_1), int32(301), int32(_a_F_network_in_4))
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L32:
	;
	m.Env.X__assert_fail(m, int32(_a_F_network_in_0), int32(_a_F_network_in_1), int32(276), int32(_a_F_network_in_4))
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_network_in[0])) = int32(5)
	goto L26
L34:
	;
	if v60 == int32(-1) {
		goto L232
	} else {
		goto L233
	}
L35:
	;
	if v60 == int32(-1) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v71 = l0
	v73 = int32(4)
	v74 = v53
	goto L42
L37:
	;
	goto L38
L38:
	;
	v517 = l0 + int32(1)
	v518 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v518 == int32(48) {
		goto L128
	} else {
		goto L129
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_network_in[0])) = int32(35)
	goto L26
L40:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_network_in[0])) = int32(44)
	goto L26
L41:
	;
	if v251 != 0 {
		goto L83
	} else {
		goto L84
	}
L42:
	;
	v85 = v71 + int32(1)
	v87 = int32(*(*int8)(unsafe.Add(mBase, uint32(v71))))
	if base.Ui32(int32(9)) < base.Ui32((v87-int32(48))&int32(255)) {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	if v243 == int32(47) {
		v251 = v226
		v253 = v228
		v255 = v241
		v256 = v239
		goto L41
	} else {
		goto L80
	}
L44:
	;
	v251 = v87
	v253 = v85
	v255 = v73
	v256 = v74
	goto L41
L45:
	;
	goto L46
L46:
	;
	v97 = v87
	v98 = v85
	v100 = int32(0)
	goto L47
L47:
	;
	v112 = int32(_a_F_network_in_5)
	v113 = int32(11)
	goto L52
L48:
	;
	if v73 == int32(0) {
		goto L39
	} else {
		goto L78
	}
L49:
	;
	v218 = v216 - int32(_a_F_network_in_5)
	if base.Ui32(int32(10)) <= base.Ui32(v218) {
		goto L32
	} else {
		goto L75
	}
L50:
	;
	v216 = int32(0)
	goto L49
L51:
	;
	v194 = v187
	v196 = v189
	goto L69
L52:
	;
	goto L60
L60:
	;
	v150 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_network_in[1])))
	if v150 == v97&int32(255) {
		v180 = v112
		v182 = v113
		goto L61
	} else {
		goto L62
	}
L61:
	;
	if v182 == int32(0) {
		goto L50
	} else {
		goto L68
	}
L62:
	;
	goto L63
L63:
	;
	v160 = v112
	v162 = v113
	goto L64
L64:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v160)))
	v167 = v166 ^ v97&int32(255)*int32(16843009)
	v170 = int32(-2139062144)
	if (int32(16843008)-v167|v167)&v170 != v170 {
		v187 = v160
		v189 = v162
		goto L51
	} else {
		goto L66
	}
L65:
	;
	v180 = v175
	v182 = v177
	goto L61
L66:
	;
	v174 = int32(4)
	v175 = v160 + v174
	v177 = v162 - v174
	if base.Ui32(int32(3)) < base.Ui32(v177) {
		v160 = v175
		v162 = v177
		goto L64
	} else {
		goto L67
	}
L67:
	;
	goto L65
L68:
	;
	v187 = v180
	v189 = v182
	goto L51
L69:
	;
	v199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v194))))
	if v97&int32(255) == v199 {
		goto L71
	} else {
		goto L72
	}
L70:
	;
	goto L50
L71:
	;
	v216 = v194
	goto L49
L72:
	;
	goto L73
L73:
	;
	v201 = int32(1)
	v204 = v196 - v201
	if v204 != 0 {
		v194 = v194 + v201
		v196 = v204
		goto L69
	} else {
		goto L74
	}
L74:
	;
	goto L70
L75:
	;
	v223 = v218 + v100*int32(10)
	if int32(255) < v223 {
		goto L40
	} else {
		goto L76
	}
L76:
	;
	v226 = int32(*(*int8)(unsafe.Add(mBase, uint32(v98))))
	v228 = v98 + int32(1)
	if base.Ui32((v226-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v97 = v226
		v98 = v228
		v100 = v223
		goto L47
	} else {
		goto L77
	}
L77:
	;
	goto L48
L78:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v74))) = uint8(v223)
	v238 = int32(1)
	v239 = v74 + v238
	v241 = v73 - v238
	v243 = v226 & int32(255)
	if v243 == int32(46) {
		v71 = v228
		v73 = v241
		v74 = v239
		goto L42
	} else {
		goto L79
	}
L79:
	;
	goto L43
L80:
	;
	if v243 != 0 {
		goto L40
	} else {
		goto L81
	}
L81:
	;
	v251 = v226
	v253 = v228
	v255 = v241
	v256 = v239
	goto L41
L82:
	;
	if v256 == v53 {
		goto L40
	} else {
		goto L123
	}
L83:
	;
	if v251 != int32(47) {
		goto L40
	} else {
		goto L86
	}
L84:
	;
	goto L85
L85:
	;
	if v256-v53 != int32(4) {
		goto L40
	} else {
		goto L122
	}
L86:
	;
	v268 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v253))))
	if base.Ui32(int32(9)) < base.Ui32((v268-int32(48))&int32(255)) {
		goto L40
	} else {
		goto L87
	}
L87:
	;
	if base.Ui32(v256) <= base.Ui32(v53) {
		goto L40
	} else {
		goto L88
	}
L88:
	;
	v280 = int32(0)
	v281 = v268
	v282 = v253
	goto L89
L89:
	;
	v295 = int32(_a_F_network_in_5)
	v297 = v281 & int32(255)
	v298 = int32(11)
	goto L94
L90:
	;
	if v411&int32(255) != 0 {
		goto L40
	} else {
		goto L119
	}
L91:
	;
	v403 = v401 - int32(_a_F_network_in_5)
	if base.Ui32(int32(10)) <= base.Ui32(v403) {
		goto L31
	} else {
		goto L117
	}
L92:
	;
	v401 = int32(0)
	goto L91
L93:
	;
	v379 = v372
	v381 = v374
	goto L111
L94:
	;
	goto L102
L102:
	;
	v335 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_network_in[1])))
	if v335 == v297&int32(255) {
		v365 = v295
		v367 = v298
		goto L103
	} else {
		goto L104
	}
L103:
	;
	if v367 == int32(0) {
		goto L92
	} else {
		goto L110
	}
L104:
	;
	goto L105
L105:
	;
	v345 = v295
	v347 = v298
	goto L106
L106:
	;
	v351 = *(*int32)(unsafe.Add(mBase, uint32(v345)))
	v352 = v351 ^ v297&int32(255)*int32(16843009)
	v355 = int32(-2139062144)
	if (int32(16843008)-v352|v352)&v355 != v355 {
		v372 = v345
		v374 = v347
		goto L93
	} else {
		goto L108
	}
L107:
	;
	v365 = v360
	v367 = v362
	goto L103
L108:
	;
	v359 = int32(4)
	v360 = v345 + v359
	v362 = v347 - v359
	if base.Ui32(int32(3)) < base.Ui32(v362) {
		v345 = v360
		v347 = v362
		goto L106
	} else {
		goto L109
	}
L109:
	;
	goto L107
L110:
	;
	v372 = v365
	v374 = v367
	goto L93
L111:
	;
	v384 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v379))))
	if v297&int32(255) == v384 {
		goto L113
	} else {
		goto L114
	}
L112:
	;
	goto L92
L113:
	;
	v401 = v379
	goto L91
L114:
	;
	goto L115
L115:
	;
	v386 = int32(1)
	v389 = v381 - v386
	if v389 != 0 {
		v379 = v379 + v386
		v381 = v389
		goto L111
	} else {
		goto L116
	}
L116:
	;
	goto L112
L117:
	;
	v406 = int32(10)
	v408 = v403 + v280*v406
	v410 = v282 + int32(1)
	v411 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v410))))
	if base.Ui32((v411-int32(48))&int32(255)) < base.Ui32(v406) {
		v280 = v408
		v281 = v411
		v282 = v410
		goto L89
	} else {
		goto L118
	}
L118:
	;
	goto L90
L119:
	;
	if int32(32) < v408 {
		goto L39
	} else {
		goto L120
	}
L120:
	;
	if v408 != int32(-1) {
		v449 = v408
		goto L82
	} else {
		goto L121
	}
L121:
	;
	goto L85
L122:
	;
	v449 = int32(32)
	goto L82
L123:
	;
	v467 = base.I32_div_s(v449, int32(8))
	if v256-v53 < v467 {
		goto L40
	} else {
		goto L124
	}
L124:
	;
	if v255 == int32(0) {
		v1299 = v449
		goto L25
	} else {
		goto L125
	}
L125:
	;
	v473 = F__emscripten_memset_bulkmem(m, v256, base.I32_extend8_s(int32(0)), v255)
	mBase = m.M
	goto L126
L126:
	;
	v1332 = v449
	goto L15
L127:
	;
	v760 = v517
	v762 = v518
	v763 = v60
	v766 = v53
	goto L190
L128:
	;
	v521 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v517))))
	if v521|int32(32) != int32(120) {
		goto L127
	} else {
		goto L131
	}
L129:
	;
	goto L130
L130:
	;
	if base.Ui32((v518-int32(48))&int32(255)) <= base.Ui32(int32(9)) {
		goto L127
	} else {
		goto L189
	}
L131:
	;
	v526 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+2)))
	goto L132
L132:
	;
	if base.B2i32(base.Ui32(v526-int32(48)) < base.Ui32(int32(10)))|base.B2i32(base.Ui32(v526|int32(32)-int32(97)) < base.Ui32(int32(6))) == int32(0) {
		goto L127
	} else {
		goto L133
	}
L133:
	;
	v540 = int32(35)
	if v60 == int32(0) {
		v1265 = v540
		goto L27
	} else {
		goto L134
	}
L134:
	;
	if v526 == int32(0) {
		goto L135
	} else {
		goto L136
	}
L135:
	;
	v1265 = int32(44)
	goto L27
L136:
	;
	goto L137
L137:
	;
	v552 = v526
	v554 = int32(0)
	v556 = v60
	v559 = v53
	v564 = l0 + int32(3)
	v566 = int32(0)
	goto L139
L138:
	;
	if v734 == int32(0) {
		v976 = v729
		v978 = v731
		v979 = v732
		v982 = v733
		goto L28
	} else {
		goto L187
	}
L139:
	;
	v567 = base.I32_extend8_s(v552)
	v569 = v552 & int32(255)
	goto L141
L140:
	;
	v729 = v726
	v731 = int32(0)
	v732 = v720
	v733 = v721
	v734 = v722
	v736 = v723
	goto L138
L141:
	;
	if base.B2i32(base.Ui32(v569-int32(48)) < base.Ui32(int32(10)))|base.B2i32(base.Ui32(v569|int32(32)-int32(97)) < base.Ui32(int32(6))) == int32(0) {
		goto L142
	} else {
		goto L143
	}
L142:
	;
	v729 = v564
	v731 = v567
	v732 = v556
	v733 = v559
	v734 = v554
	v736 = v566
	goto L138
L143:
	;
	goto L144
L144:
	;
	v583 = int32(_a_F_network_in_6)
	if base.Ui32((v552-int32(65))&int32(255)) <= base.Ui32(int32(25)) {
		goto L145
	} else {
		goto L146
	}
L145:
	;
	if base.Ui32(v569-int32(65)) < base.Ui32(int32(26)) {
		goto L149
	} else {
		goto L150
	}
L146:
	;
	v597 = v567
	goto L147
L147:
	;
	v598 = int32(17)
	goto L155
L148:
	;
	v597 = v596
	goto L147
L149:
	;
	v596 = v569 | int32(32)
	goto L151
L150:
	;
	v596 = v569
	goto L151
L151:
	;
	goto L148
L152:
	;
	v703 = v701 - int32(_a_F_network_in_6)
	if base.Ui32(int32(16)) <= base.Ui32(v703) {
		goto L30
	} else {
		goto L178
	}
L153:
	;
	v701 = int32(0)
	goto L152
L154:
	;
	v679 = v672
	v681 = v674
	goto L172
L155:
	;
	goto L163
L163:
	;
	v635 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_network_in[2])))
	if v635 == v597&int32(255) {
		v665 = v583
		v667 = v598
		goto L164
	} else {
		goto L165
	}
L164:
	;
	if v667 == int32(0) {
		goto L153
	} else {
		goto L171
	}
L165:
	;
	goto L166
L166:
	;
	v645 = v583
	v647 = v598
	goto L167
L167:
	;
	v651 = *(*int32)(unsafe.Add(mBase, uint32(v645)))
	v652 = v651 ^ v597&int32(255)*int32(16843009)
	v655 = int32(-2139062144)
	if (int32(16843008)-v652|v652)&v655 != v655 {
		v672 = v645
		v674 = v647
		goto L154
	} else {
		goto L169
	}
L168:
	;
	v665 = v660
	v667 = v662
	goto L164
L169:
	;
	v659 = int32(4)
	v660 = v645 + v659
	v662 = v647 - v659
	if base.Ui32(int32(3)) < base.Ui32(v662) {
		v645 = v660
		v647 = v662
		goto L167
	} else {
		goto L170
	}
L170:
	;
	goto L168
L171:
	;
	v672 = v665
	v674 = v667
	goto L154
L172:
	;
	v684 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v679))))
	if v597&int32(255) == v684 {
		goto L174
	} else {
		goto L175
	}
L173:
	;
	goto L153
L174:
	;
	v701 = v679
	goto L152
L175:
	;
	goto L176
L176:
	;
	v686 = int32(1)
	v689 = v681 - v686
	if v689 != 0 {
		v679 = v679 + v686
		v681 = v689
		goto L172
	} else {
		goto L177
	}
L177:
	;
	goto L173
L178:
	;
	v708 = v703 | v566<<(uint(int32(4))%32)
	v709 = int32(1)
	if v554 == v709 {
		goto L179
	} else {
		goto L180
	}
L179:
	;
	if v556 == int32(0) {
		v1265 = v540
		goto L27
	} else {
		goto L182
	}
L180:
	;
	v720 = v556
	v721 = v559
	v722 = v709
	goto L181
L181:
	;
	if v554 != 0 {
		goto L183
	} else {
		goto L184
	}
L182:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v559))) = uint8(v708)
	v715 = int32(1)
	v720 = v556 - v715
	v721 = v559 + v715
	v722 = int32(0)
	goto L181
L183:
	;
	v723 = v708
	goto L185
L184:
	;
	v723 = v703
	goto L185
L185:
	;
	v724 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v564))))
	v726 = v564 + int32(1)
	if v724 != 0 {
		v552 = v724
		v554 = v722
		v556 = v720
		v559 = v721
		v564 = v726
		v566 = v723
		goto L139
	} else {
		goto L186
	}
L186:
	;
	goto L140
L187:
	;
	if v732 == int32(0) {
		v1265 = v540
		goto L27
	} else {
		goto L188
	}
L188:
	;
	v742 = v736 << (uint(int32(4)) % 32)
	*(*uint8)(unsafe.Add(mBase, uint32(v733))) = uint8(v742)
	v744 = int32(1)
	v976 = v729
	v978 = v731
	v979 = v732 - v744
	v982 = v733 + v744
	goto L28
L189:
	;
	v1265 = int32(44)
	goto L27
L190:
	;
	v781 = v760
	v782 = int32(0)
	v783 = v762 & int32(255)
	goto L192
L191:
	;
	v1265 = v904
	goto L27
L192:
	;
	goto L198
L193:
	;
	if v763 == int32(0) {
		goto L223
	} else {
		goto L224
	}
L194:
	;
	v901 = v899 - int32(_a_F_network_in_7)
	if base.Ui32(int32(10)) <= base.Ui32(v901) {
		goto L29
	} else {
		goto L220
	}
L195:
	;
	v899 = int32(0)
	goto L194
L196:
	;
	v877 = v870
	v879 = v872
	goto L214
L197:
	;
	if base.B2i32(v817 != v818) == int32(0) {
		goto L195
	} else {
		goto L205
	}
L198:
	;
	goto L199
L199:
	;
	v809 = int32(_a_F_network_in_7)
	v811 = int32(11)
	goto L200
L200:
	;
	v814 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v809))))
	if v814 == v783&int32(255) {
		v870 = v809
		v872 = v811
		goto L196
	} else {
		goto L202
	}
L201:
	;
	goto L197
L202:
	;
	v816 = int32(1)
	v817 = v811 - v816
	v818 = int32(0)
	v821 = v809 + v816
	if v821&int32(3) == v818 {
		goto L197
	} else {
		goto L203
	}
L203:
	;
	if v817 != 0 {
		v809 = v821
		v811 = v817
		goto L200
	} else {
		goto L204
	}
L204:
	;
	goto L201
L205:
	;
	v833 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v821))))
	if v833 == v783&int32(255) {
		v863 = v821
		v865 = v817
		goto L206
	} else {
		goto L207
	}
L206:
	;
	if v865 == int32(0) {
		goto L195
	} else {
		goto L213
	}
L207:
	;
	if base.Ui32(v817) < base.Ui32(int32(4)) {
		v863 = v821
		v865 = v817
		goto L206
	} else {
		goto L208
	}
L208:
	;
	v843 = v821
	v845 = v817
	goto L209
L209:
	;
	v849 = *(*int32)(unsafe.Add(mBase, uint32(v843)))
	v850 = v849 ^ v783&int32(255)*int32(16843009)
	v853 = int32(-2139062144)
	if (int32(16843008)-v850|v850)&v853 != v853 {
		v870 = v843
		v872 = v845
		goto L196
	} else {
		goto L211
	}
L210:
	;
	v863 = v858
	v865 = v860
	goto L206
L211:
	;
	v857 = int32(4)
	v858 = v843 + v857
	v860 = v845 - v857
	if base.Ui32(int32(3)) < base.Ui32(v860) {
		v843 = v858
		v845 = v860
		goto L209
	} else {
		goto L212
	}
L212:
	;
	goto L210
L213:
	;
	v870 = v863
	v872 = v865
	goto L196
L214:
	;
	v882 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v877))))
	if v783&int32(255) == v882 {
		goto L216
	} else {
		goto L217
	}
L215:
	;
	goto L195
L216:
	;
	v899 = v877
	goto L194
L217:
	;
	goto L218
L218:
	;
	v884 = int32(1)
	v887 = v879 - v884
	if v887 != 0 {
		v877 = v877 + v884
		v879 = v887
		goto L214
	} else {
		goto L219
	}
L219:
	;
	goto L215
L220:
	;
	v904 = int32(44)
	v907 = v901 + v782*int32(10)
	if int32(255) < v907 {
		v1265 = v904
		goto L27
	} else {
		goto L221
	}
L221:
	;
	v911 = v781 + int32(1)
	v912 = int32(*(*int8)(unsafe.Add(mBase, uint32(v781))))
	if base.Ui32((v912-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v781 = v911
		v782 = v907
		v783 = v912
		goto L192
	} else {
		goto L222
	}
L222:
	;
	goto L193
L223:
	;
	v1265 = int32(35)
	goto L27
L224:
	;
	goto L225
L225:
	;
	v922 = int32(1)
	v923 = v763 - v922
	*(*uint8)(unsafe.Add(mBase, uint32(v766))) = uint8(v907)
	v926 = v766 + v922
	v928 = v912 & int32(255)
	if v928 != int32(46) {
		goto L226
	} else {
		goto L227
	}
L226:
	;
	if v928 == int32(0) {
		v976 = v911
		v978 = v912
		v979 = v923
		v982 = v926
		goto L28
	} else {
		goto L229
	}
L227:
	;
	goto L228
L228:
	;
	v937 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v781)+1)))
	if base.Ui32((v937-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v760 = v781 + int32(2)
		v762 = v937
		v763 = v923
		v766 = v926
		goto L190
	} else {
		goto L231
	}
L229:
	;
	if v928 != int32(47) {
		v1265 = v904
		goto L27
	} else {
		goto L230
	}
L230:
	;
	v976 = v911
	v978 = v912
	v979 = v923
	v982 = v926
	goto L28
L231:
	;
	goto L191
L232:
	;
	v947 = F_inet_cidr_pton_ipv6(m, l0, v53, int32(16))
	mBase = m.M
	v1332 = v947
	goto L15
L233:
	;
	goto L234
L234:
	;
	v948 = F_inet_cidr_pton_ipv6(m, l0, v53, v60)
	mBase = m.M
	v1332 = v948
	goto L15
L235:
	;
	if v1210 <= v1211 {
		v1299 = v1210
		goto L25
	} else {
		goto L293
	}
L236:
	;
	v1179 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53))))
	if base.Ui32(int32(239)) < base.Ui32(v1179) {
		v1194 = int32(32)
		goto L279
	} else {
		goto L280
	}
L237:
	;
	m.Env.X__assert_fail(m, int32(_a_F_network_in_0), int32(_a_F_network_in_1), int32(181), int32(_a_F_network_in_2))
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L238:
	;
	if v1133 == int32(-1) {
		goto L236
	} else {
		goto L278
	}
L239:
	;
	v992 = int32(44)
	v993 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v976))))
	if base.Ui32(int32(9)) < base.Ui32((v993-int32(48))&int32(255)) {
		v1265 = v992
		goto L27
	} else {
		goto L242
	}
L240:
	;
	goto L241
L241:
	;
	v1148 = int32(44)
	if v978 != 0 {
		v1265 = v1148
		goto L27
	} else {
		goto L276
	}
L242:
	;
	if base.Ui32(v982) <= base.Ui32(v53) {
		v1265 = v992
		goto L27
	} else {
		goto L243
	}
L243:
	;
	v1005 = int32(0)
	v1006 = v976
	v1008 = v993
	goto L244
L244:
	;
	v1022 = v1008 & int32(255)
	goto L250
L245:
	;
	if v1136&int32(255) != 0 {
		v1265 = v992
		goto L27
	} else {
		goto L274
	}
L246:
	;
	v1128 = v1126 - int32(_a_F_network_in_7)
	if base.Ui32(int32(10)) <= base.Ui32(v1128) {
		goto L237
	} else {
		goto L272
	}
L247:
	;
	v1126 = int32(0)
	goto L246
L248:
	;
	v1104 = v1097
	v1106 = v1099
	goto L266
L249:
	;
	if base.B2i32(v1044 != v1045) == int32(0) {
		goto L247
	} else {
		goto L257
	}
L250:
	;
	goto L251
L251:
	;
	v1036 = int32(_a_F_network_in_7)
	v1038 = int32(11)
	goto L252
L252:
	;
	v1041 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1036))))
	if v1041 == v1022&int32(255) {
		v1097 = v1036
		v1099 = v1038
		goto L248
	} else {
		goto L254
	}
L253:
	;
	goto L249
L254:
	;
	v1043 = int32(1)
	v1044 = v1038 - v1043
	v1045 = int32(0)
	v1048 = v1036 + v1043
	if v1048&int32(3) == v1045 {
		goto L249
	} else {
		goto L255
	}
L255:
	;
	if v1044 != 0 {
		v1036 = v1048
		v1038 = v1044
		goto L252
	} else {
		goto L256
	}
L256:
	;
	goto L253
L257:
	;
	v1060 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1048))))
	if v1060 == v1022&int32(255) {
		v1090 = v1048
		v1092 = v1044
		goto L258
	} else {
		goto L259
	}
L258:
	;
	if v1092 == int32(0) {
		goto L247
	} else {
		goto L265
	}
L259:
	;
	if base.Ui32(v1044) < base.Ui32(int32(4)) {
		v1090 = v1048
		v1092 = v1044
		goto L258
	} else {
		goto L260
	}
L260:
	;
	v1070 = v1048
	v1072 = v1044
	goto L261
L261:
	;
	v1076 = *(*int32)(unsafe.Add(mBase, uint32(v1070)))
	v1077 = v1076 ^ v1022&int32(255)*int32(16843009)
	v1080 = int32(-2139062144)
	if (int32(16843008)-v1077|v1077)&v1080 != v1080 {
		v1097 = v1070
		v1099 = v1072
		goto L248
	} else {
		goto L263
	}
L262:
	;
	v1090 = v1085
	v1092 = v1087
	goto L258
L263:
	;
	v1084 = int32(4)
	v1085 = v1070 + v1084
	v1087 = v1072 - v1084
	if base.Ui32(int32(3)) < base.Ui32(v1087) {
		v1070 = v1085
		v1072 = v1087
		goto L261
	} else {
		goto L264
	}
L264:
	;
	goto L262
L265:
	;
	v1097 = v1090
	v1099 = v1092
	goto L248
L266:
	;
	v1109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1104))))
	if v1022&int32(255) == v1109 {
		goto L268
	} else {
		goto L269
	}
L267:
	;
	goto L247
L268:
	;
	v1126 = v1104
	goto L246
L269:
	;
	goto L270
L270:
	;
	v1111 = int32(1)
	v1114 = v1106 - v1111
	if v1114 != 0 {
		v1104 = v1104 + v1111
		v1106 = v1114
		goto L266
	} else {
		goto L271
	}
L271:
	;
	goto L267
L272:
	;
	v1131 = int32(10)
	v1133 = v1128 + v1005*v1131
	v1135 = v1006 + int32(1)
	v1136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1135))))
	if base.Ui32((v1136-int32(48))&int32(255)) < base.Ui32(v1131) {
		v1005 = v1133
		v1006 = v1135
		v1008 = v1136
		goto L244
	} else {
		goto L273
	}
L273:
	;
	goto L245
L274:
	;
	if v1133 <= int32(32) {
		goto L238
	} else {
		goto L275
	}
L275:
	;
	v1265 = int32(35)
	goto L27
L276:
	;
	if v53 == v982 {
		v1265 = v1148
		goto L27
	} else {
		goto L277
	}
L277:
	;
	goto L236
L278:
	;
	v1210 = v1133
	v1211 = (v982 - v53) << (uint(int32(3)) % 32)
	goto L235
L279:
	;
	v1197 = (v982 - v53) << (uint(int32(3)) % 32)
	if v1197 < v1194 {
		goto L286
	} else {
		goto L287
	}
L280:
	;
	if base.Ui32(int32(223)) < base.Ui32(v1179) {
		v1194 = int32(8)
		goto L279
	} else {
		goto L281
	}
L281:
	;
	if base.Ui32(int32(191)) < base.Ui32(v1179) {
		v1194 = int32(24)
		goto L279
	} else {
		goto L282
	}
L282:
	;
	if base.I32_extend8_s(v1179) < int32(0) {
		goto L283
	} else {
		goto L284
	}
L283:
	;
	v1193 = int32(16)
	goto L285
L284:
	;
	v1193 = int32(8)
	goto L285
L285:
	;
	v1194 = v1193
	goto L279
L286:
	;
	v1199 = v1194
	goto L288
L287:
	;
	v1199 = v1197
	goto L288
L288:
	;
	if v1199 != int32(8) {
		v1210 = v1199
		v1211 = v1197
		goto L235
	} else {
		goto L289
	}
L289:
	;
	if v1179 == int32(224) {
		goto L290
	} else {
		goto L291
	}
L290:
	;
	v1206 = int32(4)
	goto L292
L291:
	;
	v1206 = int32(8)
	goto L292
L292:
	;
	v1210 = v1206
	v1211 = v1197
	goto L235
L293:
	;
	v1233 = v979
	v1236 = v982
	goto L294
L294:
	;
	if v1233 == int32(0) {
		goto L296
	} else {
		goto L297
	}
L295:
	;
	v1299 = v1210
	goto L25
L296:
	;
	v1265 = int32(35)
	goto L27
L297:
	;
	goto L298
L298:
	;
	v1247 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1236))) = uint8(v1247)
	v1249 = int32(1)
	v1252 = v1236 + v1249
	if (v1252-v53)<<(uint(int32(3))%32) < v1210 {
		v1233 = v1233 - v1249
		v1236 = v1252
		goto L294
	} else {
		goto L299
	}
L299:
	;
	goto L295
L300:
	;
	v1337 = int32(1)
	v1339 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24))))
	v1341 = v1339 & v1337
	if v1341 != 0 {
		goto L303
	} else {
		goto L304
	}
L301:
	;
	goto L302
L302:
	;
	v1352 = F_errsave_start(m, l2)
	mBase = m.M
	v1353 = m.ExcPending
	if v1353 != 0 {
		goto L1
	} else {
		goto L310
	}
L303:
	;
	v1342 = v1337
	goto L305
L304:
	;
	v1342 = int32(4)
	goto L305
L305:
	;
	v1344 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24+v1342))))
	if v1344 == int32(2) {
		goto L306
	} else {
		goto L307
	}
L306:
	;
	v1347 = int32(32)
	goto L308
L307:
	;
	v1347 = int32(128)
	goto L308
L308:
	;
	if base.Ui32(v1332) <= base.Ui32(v1347) {
		goto L14
	} else {
		goto L309
	}
L309:
	;
	goto L302
L310:
	;
	if v1352 == int32(0) {
		goto L311
	} else {
		goto L312
	}
L311:
	;
	v1506 = int32(0)
	goto L13
L312:
	;
	goto L313
L313:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v1359 = m.ExcPending
	if v1359 != 0 {
		goto L1
	} else {
		goto L314
	}
L314:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+4)) = l0
	if l1 != 0 {
		goto L315
	} else {
		goto L316
	}
L315:
	;
	v1363 = int32(_a_F_network_in_8)
	goto L317
L316:
	;
	v1363 = int32(_a_F_network_in_9)
	goto L317
L317:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = v1363
	F_errmsg(m, int32(_a_F_network_in_10), v21)
	mBase = m.M
	v1367 = m.ExcPending
	if v1367 != 0 {
		goto L1
	} else {
		goto L318
	}
L318:
	;
	F_errsave_finish(m, l2, int32(_a_F_network_in_11), int32(100), int32(_a_F_network_in_12))
	mBase = m.M
	v1372 = m.ExcPending
	if v1372 != 0 {
		goto L1
	} else {
		goto L319
	}
L319:
	;
	v1506 = int32(0)
	goto L13
L320:
	;
	v1374 = v48
	goto L322
L321:
	;
	v1374 = v50
	goto L322
L322:
	;
	if l1 == int32(0) {
		goto L323
	} else {
		goto L324
	}
L323:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1374)+1)) = uint8(v1332)
	v1483 = int32(1)
	v1485 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24))))
	if v1485&v1483 != 0 {
		goto L343
	} else {
		goto L344
	}
L324:
	;
	if v1332 == v1347 {
		goto L323
	} else {
		goto L325
	}
L325:
	;
	v1379 = int32(base.Ui32(v1332) >> (uint(int32(3)) % 32))
	if v1344 == int32(2) {
		goto L326
	} else {
		goto L327
	}
L326:
	;
	v1384 = int32(4)
	goto L328
L327:
	;
	v1384 = int32(16)
	goto L328
L328:
	;
	if base.Ui32(v1384) <= base.Ui32(v1379) {
		goto L323
	} else {
		goto L329
	}
L329:
	;
	v1387 = v1374 + int32(2)
	v1389 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1387+v1379))))
	if v1389<<(uint(v1332&int32(7))%32)&int32(255) != 0 {
		goto L330
	} else {
		goto L331
	}
L330:
	;
	v1439 = int32(0)
	v1440 = F_errsave_start(m, l2)
	mBase = m.M
	v1441 = m.ExcPending
	if v1441 != 0 {
		goto L1
	} else {
		goto L337
	}
L331:
	;
	v1396 = v1379 + int32(1)
	if v1396 == v1384 {
		goto L323
	} else {
		goto L332
	}
L332:
	;
	v1399 = v1396
	goto L333
L333:
	;
	v1417 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1399+v1387))))
	if v1417 != 0 {
		goto L330
	} else {
		goto L335
	}
L334:
	;
	goto L323
L335:
	;
	v1419 = v1399 + int32(1)
	if v1384 != v1419 {
		v1399 = v1419
		goto L333
	} else {
		goto L336
	}
L336:
	;
	goto L334
L337:
	;
	if v1440 == int32(0) {
		v1506 = v1439
		goto L13
	} else {
		goto L338
	}
L338:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v1446 = m.ExcPending
	if v1446 != 0 {
		goto L1
	} else {
		goto L339
	}
L339:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+16)) = l0
	F_errmsg(m, int32(_a_F_network_in_13), v21+int32(16))
	mBase = m.M
	v1452 = m.ExcPending
	if v1452 != 0 {
		goto L1
	} else {
		goto L340
	}
L340:
	;
	F_errdetail(m, int32(_a_F_network_in_14), int32(0))
	mBase = m.M
	v1456 = m.ExcPending
	if v1456 != 0 {
		goto L1
	} else {
		goto L341
	}
L341:
	;
	F_errsave_finish(m, l2, int32(_a_F_network_in_11), int32(111), int32(_a_F_network_in_12))
	mBase = m.M
	v1461 = m.ExcPending
	if v1461 != 0 {
		goto L1
	} else {
		goto L342
	}
L342:
	;
	v1506 = v1439
	goto L13
L343:
	;
	v1488 = v1483
	goto L345
L344:
	;
	v1488 = int32(4)
	goto L345
L345:
	;
	v1490 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24+v1488))))
	if v1490 == int32(2) {
		goto L346
	} else {
		goto L347
	}
L346:
	;
	v1493 = int32(40)
	goto L348
L347:
	;
	v1493 = int32(88)
	goto L348
L348:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24))) = v1493
	v1506 = v24
	goto L13
}
func F_network_masklen(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = F_pg_detoast_datum_packed(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v7 = int32(1)
		v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3))))
		if v9&v7 != 0 {
			v12 = v7
		} else {
			v12 = int32(4)
		}
		v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3+v12)+1)))
		return v14
	}
}
func F_network_netmask(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
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
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = F_pg_detoast_datum_packed(m, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		v14 = F_palloc0(m, int32(22))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			v16 = int32(4)
			v17 = v14 + v16
			v18 = int32(1)
			v19 = v14 + v18
			v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
			v22 = v20 & v18
			v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9))))
			v27 = v25 & v18
			if v27 != 0 {
				v28 = v18
			} else {
				v28 = v16
			}
			v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9+v28)+1)))
			if v30 != 0 {
				if v22 != 0 {
					v31 = v19
				} else {
					v31 = v17
				}
				v35 = v30
				v37 = int32(0)
				for {
					if base.Ui32(int32(7)) < base.Ui32(v35) {
						v50 = int32(-1)
					} else {
						v50 = int32(255) << (uint(int32(8)-v35) % 32)
					}
					*(*uint8)(unsafe.Add(mBase, uint32(v37+(v31+int32(2))))) = uint8(v50)
					v54 = int32(8)
					if v35 <= v54 {
						v57 = v54
					} else {
						v57 = v35
					}
					v59 = v57 - int32(8)
					if v59 != 0 {
						v35 = v59
						v37 = v37 + int32(1)
						continue
					} else {
						break
					}
					break
				}
				v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9))))
				v61 = int32(1)
				v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
				v68 = v63 & v61
				v70 = v60 & v61
			} else {
				v68 = v22
				v70 = v27
			}
			if v68 != 0 {
				v75 = int32(1)
			} else {
				v75 = int32(4)
			}
			if v70 != 0 {
				v79 = int32(1)
			} else {
				v79 = int32(4)
			}
			v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9+v79))))
			*(*uint8)(unsafe.Add(mBase, uint32(v14+v75))) = uint8(v81)
			if v68 != 0 {
				v83 = v19
			} else {
				v83 = v17
			}
			v86 = int32(1)
			v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9))))
			if v88&v86 != 0 {
				v91 = v86
			} else {
				v91 = int32(4)
			}
			v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9+v91))))
			if v93 == int32(2) {
				v96 = int32(32)
			} else {
				v96 = int32(-128)
			}
			*(*uint8)(unsafe.Add(mBase, uint32(v83)+1)) = uint8(v96)
			v100 = int32(1)
			v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
			if v102&v100 != 0 {
				v105 = v100
			} else {
				v105 = int32(4)
			}
			v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14+v105))))
			if v107 == int32(2) {
				v110 = int32(40)
			} else {
				v110 = int32(88)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v14))) = v110
			return v14
		}
	}
}
func F_network_overlap(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
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
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
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
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	var v168 int32
	_ = v168
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = F_pg_detoast_datum_packed(m, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v11 = F_pg_detoast_datum_packed(m, v10)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v13 = int32(1)
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6))))
	if v15&v13 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v18 = v13
	goto L6
L5:
	;
	v18 = int32(4)
	goto L6
L6:
	;
	v19 = v6 + v18
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
	v21 = int32(1)
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
	if v23&v21 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v26 = v21
	goto L9
L8:
	;
	v26 = int32(4)
	goto L9
L9:
	;
	v27 = v11 + v26
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27))))
	if v20 != v28 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	return int32(0)
L11:
	;
	goto L12
L12:
	;
	v32 = int32(2)
	v33 = v19 + v32
	v35 = v27 + v32
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+1)))
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+1)))
	if base.Ui32(v36) < base.Ui32(v37) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v39 = v36
	goto L15
L14:
	;
	v39 = v37
	goto L15
L15:
	;
	v41 = int32(base.Ui32(v39) >> (uint(int32(3)) % 32))
	if base.Ui32(int32(4)) <= base.Ui32(v41) {
		goto L19
	} else {
		goto L20
	}
L16:
	;
	if v103 != 0 {
		goto L34
	} else {
		goto L35
	}
L17:
	;
	v103 = int32(0)
	goto L16
L18:
	;
	v77 = v72
	v78 = v73
	v79 = v74
	goto L28
L19:
	;
	if (v33|v35)&int32(3) != 0 {
		v72 = v33
		v73 = v35
		v74 = v41
		goto L18
	} else {
		goto L22
	}
L20:
	;
	v65 = v33
	v66 = v35
	v67 = v41
	goto L21
L21:
	;
	if v67 == int32(0) {
		goto L17
	} else {
		goto L27
	}
L22:
	;
	v49 = v33
	v50 = v35
	v51 = v41
	goto L23
L23:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
	if v54 != v55 {
		v72 = v49
		v73 = v50
		v74 = v51
		goto L18
	} else {
		goto L25
	}
L24:
	;
	v65 = v60
	v66 = v58
	v67 = v62
	goto L21
L25:
	;
	v57 = int32(4)
	v58 = v50 + v57
	v60 = v49 + v57
	v62 = v51 - v57
	if base.Ui32(int32(3)) < base.Ui32(v62) {
		v49 = v60
		v50 = v58
		v51 = v62
		goto L23
	} else {
		goto L26
	}
L26:
	;
	goto L24
L27:
	;
	v72 = v65
	v73 = v66
	v74 = v67
	goto L18
L28:
	;
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77))))
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78))))
	if v82 == v83 {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	v103 = v82 - v83
	goto L16
L30:
	;
	v85 = int32(1)
	v90 = v79 - v85
	if v90 != 0 {
		v77 = v77 + v85
		v78 = v78 + v85
		v79 = v90
		goto L28
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	goto L29
L33:
	;
	goto L17
L34:
	;
	return int32(0)
L35:
	;
	goto L36
L36:
	;
	v107 = v39 & int32(7)
	if v107 == int32(0) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	return int32(1)
L38:
	;
	goto L39
L39:
	;
	v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41+v33))))
	v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41+v35))))
	v116 = v113 ^ v115
	if base.Ui32(int32(127)) < base.Ui32(v116) {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	return int32(0)
L41:
	;
	goto L42
L42:
	;
	v121 = int32(1)
	if v107 == v121 {
		v168 = v121
		goto L43
	} else {
		goto L44
	}
L43:
	;
	return v168
L44:
	;
	if v116<<(uint(int32(1))%32)&int32(128) != 0 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	return int32(0)
L46:
	;
	goto L47
L47:
	;
	if base.Ui32(v107) < base.Ui32(int32(3)) {
		v168 = v121
		goto L43
	} else {
		goto L48
	}
L48:
	;
	if v116<<(uint(int32(2))%32)&int32(128) != 0 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	return int32(0)
L50:
	;
	goto L51
L51:
	;
	if v107 == int32(3) {
		v168 = v121
		goto L43
	} else {
		goto L52
	}
L52:
	;
	if v116<<(uint(int32(3))%32)&int32(128) != 0 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	return int32(0)
L54:
	;
	goto L55
L55:
	;
	if base.Ui32(v107) < base.Ui32(int32(5)) {
		v168 = v121
		goto L43
	} else {
		goto L56
	}
L56:
	;
	if v116<<(uint(int32(4))%32)&int32(128) != 0 {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	return int32(0)
L58:
	;
	goto L59
L59:
	;
	if v107 == int32(5) {
		v168 = v121
		goto L43
	} else {
		goto L60
	}
L60:
	;
	if v116<<(uint(int32(5))%32)&int32(128) != 0 {
		v168 = int32(0)
		goto L43
	} else {
		goto L61
	}
L61:
	;
	if v107 != int32(7) {
		v168 = int32(1)
		goto L43
	} else {
		goto L62
	}
L62:
	;
	v168 = base.B2i32(v116&int32(2) == int32(0))
	goto L43
}
func F_network_show(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	v7 = m.G0
	v9 = v7 - int32(80)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = F_pg_detoast_datum_packed(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v16 = int32(1)
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
	v20 = v18 & v16
	if v20 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v21 = v16
	goto L5
L4:
	;
	v21 = int32(4)
	goto L5
L5:
	;
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12+v21))))
	v25 = v12 + int32(1)
	v27 = v12 + int32(4)
	if v20 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v28 = v25
	goto L8
L7:
	;
	v28 = v27
	goto L8
L8:
	;
	v29 = int32(2)
	if v23 == v29 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v35 = int32(32)
	goto L11
L10:
	;
	v35 = int32(128)
	goto L11
L11:
	;
	v38 = F_pg_inet_net_ntop(m, v23, v28+v29, v35, v9+int32(16))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	if v38 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v42 = int32(47)
	v43 = F___strchrnul(m, v9+int32(16), v42)
	mBase = m.M
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43))))
	if v45 == v42 {
		goto L17
	} else {
		goto L18
	}
L14:
	;
	goto L15
L15:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L1
	} else {
		goto L45
	}
L16:
	;
	if v49 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L17:
	;
	v49 = v43
	goto L19
L18:
	;
	v49 = int32(0)
	goto L19
L19:
	;
	goto L16
L20:
	;
	v53 = v9 + int32(16)
	if v53&int32(3) == int32(0) {
		v77 = v53
		goto L25
	} else {
		goto L26
	}
L21:
	;
	goto L22
L22:
	;
	v128 = F_cstring_to_text(m, v9+int32(16))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L1
	} else {
		goto L44
	}
L23:
	;
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
	if v111&int32(1) != 0 {
		goto L40
	} else {
		goto L41
	}
L24:
	;
	v110 = v102 - v53
	goto L23
L25:
	;
	v81 = v77
	goto L34
L26:
	;
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53))))
	if v61 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v110 = int32(0)
	goto L23
L28:
	;
	goto L29
L29:
	;
	v66 = v53
	goto L30
L30:
	;
	v70 = v66 + int32(1)
	if v70&int32(3) == int32(0) {
		v77 = v70
		goto L25
	} else {
		goto L32
	}
L31:
	;
	v102 = v70
	goto L24
L32:
	;
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70))))
	if v75 != 0 {
		v66 = v70
		goto L30
	} else {
		goto L33
	}
L33:
	;
	goto L31
L34:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v81)))
	v90 = int32(-2139062144)
	if (int32(16843008)-v87|v87)&v90 == v90 {
		v81 = v81 + int32(4)
		goto L34
	} else {
		goto L36
	}
L35:
	;
	v96 = v81
	goto L37
L36:
	;
	goto L35
L37:
	;
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96))))
	if v100 != 0 {
		v96 = v96 + int32(1)
		goto L37
	} else {
		goto L39
	}
L38:
	;
	v102 = v96
	goto L24
L39:
	;
	goto L38
L40:
	;
	v114 = v25
	goto L42
L41:
	;
	v114 = v27
	goto L42
L42:
	;
	v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114)+1)))
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v115
	v123 = F_pg_snprintf(m, v110+(v9+int32(16)), int32(50)-v110, int32(_a_F_network_show_0), v9)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	goto L22
L44:
	;
	m.G0 = v9 + int32(80)
	return v128
L45:
	;
	F_errcode(m, int32(50462850))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	F_errmsg(m, int32(_a_F_network_show_1), int32(0))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	F_errfinish(m, int32(_a_F_network_show_2), int32(1174), int32(_a_F_network_show_3))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_network_sortsupport(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v30 int32
	_ = v30
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v4)+12)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v4)+16)) = int32(1458)
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4)+20)))
	if v9 == int32(1) {
		v12 = int32(_a_F_network_sortsupport_0)
		v13 = *(*int32)(unsafe.Add(mBase, _c_F_network_sortsupport[0]))
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
		*(*int32)(unsafe.Add(mBase, _c_F_network_sortsupport[0])) = v15
		v18 = F_palloc(m, int32(40))
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return int32(0)
		} else {
			v22 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(v18)+8)) = uint8(v22)
			*(*int64)(unsafe.Add(mBase, uint32(v18))) = int64(0)
			F_initHyperLogLog(m, v18+int32(16), int32(10))
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v4)+32)) = int32(1458)
				*(*int32)(unsafe.Add(mBase, uint32(v4)+28)) = int32(1459)
				*(*int32)(unsafe.Add(mBase, uint32(v4)+24)) = int32(1460)
				*(*int32)(unsafe.Add(mBase, uint32(v4)+16)) = int32(116)
				*(*int32)(unsafe.Add(mBase, uint32(v4)+12)) = v18
				*(*int32)(unsafe.Add(mBase, _c_F_network_sortsupport[0])) = v13
				return int32(0)
			}
		}
	} else {
		return int32(0)
	}
}
