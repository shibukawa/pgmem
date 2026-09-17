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
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v266 int32
	_ = v266
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
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
	var v409 int32
	_ = v409
	var v450 int32
	_ = v450
	var v465 int32
	_ = v465
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v517 int32
	_ = v517
	var v522 int32
	_ = v522
	var v536 int32
	_ = v536
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v554 int32
	_ = v554
	var v556 int32
	_ = v556
	var v559 int32
	_ = v559
	var v562 int32
	_ = v562
	var v564 int32
	_ = v564
	var v578 int32
	_ = v578
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v593 int32
	_ = v593
	var v632 int32
	_ = v632
	var v633 int32
	_ = v633
	var v642 int32
	_ = v642
	var v644 int32
	_ = v644
	var v648 int32
	_ = v648
	var v649 int32
	_ = v649
	var v652 int32
	_ = v652
	var v656 int32
	_ = v656
	var v657 int32
	_ = v657
	var v659 int32
	_ = v659
	var v662 int32
	_ = v662
	var v664 int32
	_ = v664
	var v669 int32
	_ = v669
	var v671 int32
	_ = v671
	var v676 int32
	_ = v676
	var v678 int32
	_ = v678
	var v681 int32
	_ = v681
	var v683 int32
	_ = v683
	var v686 int32
	_ = v686
	var v698 int32
	_ = v698
	var v700 int32
	_ = v700
	var v705 int32
	_ = v705
	var v706 int32
	_ = v706
	var v712 int32
	_ = v712
	var v717 int32
	_ = v717
	var v718 int32
	_ = v718
	var v719 int32
	_ = v719
	var v720 int32
	_ = v720
	var v721 int32
	_ = v721
	var v723 int32
	_ = v723
	var v725 int32
	_ = v725
	var v728 int32
	_ = v728
	var v729 int32
	_ = v729
	var v730 int32
	_ = v730
	var v731 int32
	_ = v731
	var v732 int32
	_ = v732
	var v738 int32
	_ = v738
	var v740 int32
	_ = v740
	var v755 int32
	_ = v755
	var v758 int32
	_ = v758
	var v761 int32
	_ = v761
	var v763 int32
	_ = v763
	var v775 int32
	_ = v775
	var v776 int32
	_ = v776
	var v778 int32
	_ = v778
	var v804 int32
	_ = v804
	var v806 int32
	_ = v806
	var v809 int32
	_ = v809
	var v811 int32
	_ = v811
	var v812 int32
	_ = v812
	var v813 int32
	_ = v813
	var v816 int32
	_ = v816
	var v829 int32
	_ = v829
	var v830 int32
	_ = v830
	var v839 int32
	_ = v839
	var v841 int32
	_ = v841
	var v845 int32
	_ = v845
	var v846 int32
	_ = v846
	var v849 int32
	_ = v849
	var v853 int32
	_ = v853
	var v854 int32
	_ = v854
	var v856 int32
	_ = v856
	var v859 int32
	_ = v859
	var v861 int32
	_ = v861
	var v866 int32
	_ = v866
	var v868 int32
	_ = v868
	var v873 int32
	_ = v873
	var v875 int32
	_ = v875
	var v878 int32
	_ = v878
	var v880 int32
	_ = v880
	var v883 int32
	_ = v883
	var v895 int32
	_ = v895
	var v897 int32
	_ = v897
	var v900 int32
	_ = v900
	var v903 int32
	_ = v903
	var v907 int32
	_ = v907
	var v908 int32
	_ = v908
	var v918 int32
	_ = v918
	var v919 int32
	_ = v919
	var v922 int32
	_ = v922
	var v924 int32
	_ = v924
	var v933 int32
	_ = v933
	var v943 int32
	_ = v943
	var v944 int32
	_ = v944
	var v971 int32
	_ = v971
	var v974 int32
	_ = v974
	var v977 int32
	_ = v977
	var v979 int32
	_ = v979
	var v987 int32
	_ = v987
	var v988 int32
	_ = v988
	var v1001 int32
	_ = v1001
	var v1003 int32
	_ = v1003
	var v1004 int32
	_ = v1004
	var v1017 int32
	_ = v1017
	var v1032 int32
	_ = v1032
	var v1034 int32
	_ = v1034
	var v1037 int32
	_ = v1037
	var v1039 int32
	_ = v1039
	var v1040 int32
	_ = v1040
	var v1041 int32
	_ = v1041
	var v1044 int32
	_ = v1044
	var v1057 int32
	_ = v1057
	var v1058 int32
	_ = v1058
	var v1067 int32
	_ = v1067
	var v1069 int32
	_ = v1069
	var v1073 int32
	_ = v1073
	var v1074 int32
	_ = v1074
	var v1077 int32
	_ = v1077
	var v1081 int32
	_ = v1081
	var v1082 int32
	_ = v1082
	var v1084 int32
	_ = v1084
	var v1087 int32
	_ = v1087
	var v1089 int32
	_ = v1089
	var v1094 int32
	_ = v1094
	var v1096 int32
	_ = v1096
	var v1101 int32
	_ = v1101
	var v1103 int32
	_ = v1103
	var v1106 int32
	_ = v1106
	var v1108 int32
	_ = v1108
	var v1111 int32
	_ = v1111
	var v1123 int32
	_ = v1123
	var v1125 int32
	_ = v1125
	var v1128 int32
	_ = v1128
	var v1130 int32
	_ = v1130
	var v1131 int32
	_ = v1131
	var v1176 int32
	_ = v1176
	var v1190 int32
	_ = v1190
	var v1191 int32
	_ = v1191
	var v1194 int32
	_ = v1194
	var v1196 int32
	_ = v1196
	var v1203 int32
	_ = v1203
	var v1207 int32
	_ = v1207
	var v1209 int32
	_ = v1209
	var v1231 int32
	_ = v1231
	var v1233 int32
	_ = v1233
	var v1242 int32
	_ = v1242
	var v1244 int32
	_ = v1244
	var v1247 int32
	_ = v1247
	var v1262 int32
	_ = v1262
	var v1294 int32
	_ = v1294
	var v1323 int32
	_ = v1323
	var v1328 int32
	_ = v1328
	var v1330 int32
	_ = v1330
	var v1332 int32
	_ = v1332
	var v1333 int32
	_ = v1333
	var v1335 int32
	_ = v1335
	var v1338 int32
	_ = v1338
	var v1343 int32
	_ = v1343
	var v1344 int32
	_ = v1344
	var v1350 int32
	_ = v1350
	var v1354 int32
	_ = v1354
	var v1358 int32
	_ = v1358
	var v1363 int32
	_ = v1363
	var v1365 int32
	_ = v1365
	var v1371 int32
	_ = v1371
	var v1376 int32
	_ = v1376
	var v1379 int32
	_ = v1379
	var v1381 int32
	_ = v1381
	var v1388 int32
	_ = v1388
	var v1391 int32
	_ = v1391
	var v1408 int32
	_ = v1408
	var v1410 int32
	_ = v1410
	var v1429 int32
	_ = v1429
	var v1430 int32
	_ = v1430
	var v1431 int32
	_ = v1431
	var v1436 int32
	_ = v1436
	var v1442 int32
	_ = v1442
	var v1446 int32
	_ = v1446
	var v1451 int32
	_ = v1451
	var v1472 int32
	_ = v1472
	var v1474 int32
	_ = v1474
	var v1477 int32
	_ = v1477
	var v1479 int32
	_ = v1479
	var v1482 int32
	_ = v1482
	var v1496 int32
	_ = v1496
	v18 = m.G0
	v20 = v18 - int32(32)
	m.G0 = v20
	v23 = F_palloc0(m, int32(22))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v27 = int32(1)
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23))))
	v31 = v29 & v27
	if v31 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v32 = v27
	goto L5
L4:
	;
	v32 = int32(4)
	goto L5
L5:
	;
	v36 = int32(58)
	v37 = F___strchrnul(m, l0, v36)
	mBase = m.M
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37))))
	if v39 == v36 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	if v43 != 0 {
		goto L10
	} else {
		goto L11
	}
L7:
	;
	v43 = v37
	goto L9
L8:
	;
	v43 = int32(0)
	goto L9
L9:
	;
	goto L6
L10:
	;
	v44 = int32(3)
	goto L12
L11:
	;
	v44 = int32(2)
	goto L12
L12:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v23+v32))) = uint8(v44)
	v47 = v23 + int32(1)
	v49 = v23 + int32(4)
	if v31 != 0 {
		goto L16
	} else {
		goto L17
	}
L13:
	;
	m.G0 = v20 + int32(32)
	return v1496
L14:
	;
	if v1332 != 0 {
		goto L310
	} else {
		goto L311
	}
L15:
	;
	if int32(0) <= v1323 {
		goto L290
	} else {
		goto L291
	}
L16:
	;
	v50 = v47
	goto L18
L17:
	;
	v50 = v49
	goto L18
L18:
	;
	v51 = int32(2)
	v52 = v50 + v51
	if v44 == v51 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v57 = int32(4)
	goto L21
L20:
	;
	v57 = int32(16)
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
	v59 = v57
	goto L24
L23:
	;
	v59 = int32(-1)
	goto L24
L24:
	;
	switch v44 - int32(2) {
	case 0:
		goto L35
	case 1:
		goto L34
	default:
		goto L33
	}
L25:
	;
	v1323 = v1294
	goto L15
L26:
	;
	v1294 = int32(-1)
	goto L25
L27:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_network_in[0])) = v1262
	goto L26
L28:
	;
	if v974 == int32(47) {
		goto L232
	} else {
		goto L233
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
	if v59 == int32(-1) {
		goto L225
	} else {
		goto L226
	}
L35:
	;
	if v59 == int32(-1) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v69 = l0
	v74 = int32(4)
	v75 = v52
	goto L42
L37:
	;
	goto L38
L38:
	;
	v513 = l0 + int32(1)
	v514 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v514 == int32(48) {
		goto L123
	} else {
		goto L124
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
	if v252 != 0 {
		goto L82
	} else {
		goto L83
	}
L42:
	;
	v83 = v69 + int32(1)
	v85 = int32(*(*int8)(unsafe.Add(mBase, uint32(v69))))
	if base.Ui32(int32(9)) < base.Ui32((v85-int32(48))&int32(255)) {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	if v242 == int32(47) {
		v251 = v227
		v252 = v225
		v256 = v240
		v257 = v238
		goto L41
	} else {
		goto L79
	}
L44:
	;
	v251 = v83
	v252 = v85
	v256 = v74
	v257 = v75
	goto L41
L45:
	;
	goto L46
L46:
	;
	v95 = v83
	v97 = v85
	v98 = int32(0)
	goto L47
L47:
	;
	v109 = int32(_a_F_network_in_5)
	v110 = int32(11)
	goto L52
L48:
	;
	if v74 == int32(0) {
		goto L39
	} else {
		goto L77
	}
L49:
	;
	v217 = v215 - int32(_a_F_network_in_5)
	if base.Ui32(int32(10)) <= base.Ui32(v217) {
		goto L32
	} else {
		goto L74
	}
L50:
	;
	v215 = int32(0)
	goto L49
L51:
	;
	v193 = v186
	v195 = v188
	goto L68
L52:
	;
	goto L59
L59:
	;
	v149 = v97 & int32(255)
	v150 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_network_in[1])))
	if base.B2i32(v149 == v150)|int32(0) == int32(0) {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v159 = v109
	v161 = v110
	goto L63
L61:
	;
	v179 = v109
	v181 = v110
	goto L62
L62:
	;
	if v181 == int32(0) {
		goto L50
	} else {
		goto L67
	}
L63:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v159)))
	v166 = v165 ^ v149*int32(16843009)
	v169 = int32(-2139062144)
	if (int32(16843008)-v166|v166)&v169 != v169 {
		v186 = v159
		v188 = v161
		goto L51
	} else {
		goto L65
	}
L64:
	;
	v179 = v174
	v181 = v176
	goto L62
L65:
	;
	v173 = int32(4)
	v174 = v159 + v173
	v176 = v161 - v173
	if base.Ui32(int32(3)) < base.Ui32(v176) {
		v159 = v174
		v161 = v176
		goto L63
	} else {
		goto L66
	}
L66:
	;
	goto L64
L67:
	;
	v186 = v179
	v188 = v181
	goto L51
L68:
	;
	v198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v193))))
	if v97&int32(255) == v198 {
		goto L70
	} else {
		goto L71
	}
L69:
	;
	goto L50
L70:
	;
	v215 = v193
	goto L49
L71:
	;
	goto L72
L72:
	;
	v200 = int32(1)
	v203 = v195 - v200
	if v203 != 0 {
		v193 = v193 + v200
		v195 = v203
		goto L68
	} else {
		goto L73
	}
L73:
	;
	goto L69
L74:
	;
	v222 = v217 + v98*int32(10)
	if int32(255) < v222 {
		goto L40
	} else {
		goto L75
	}
L75:
	;
	v225 = int32(*(*int8)(unsafe.Add(mBase, uint32(v95))))
	v227 = v95 + int32(1)
	if base.Ui32((v225-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v95 = v227
		v97 = v225
		v98 = v222
		goto L47
	} else {
		goto L76
	}
L76:
	;
	goto L48
L77:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v75))) = uint8(v222)
	v237 = int32(1)
	v238 = v75 + v237
	v240 = v74 - v237
	v242 = v225 & int32(255)
	if v242 == int32(46) {
		v69 = v227
		v74 = v240
		v75 = v238
		goto L42
	} else {
		goto L78
	}
L78:
	;
	goto L43
L79:
	;
	if v242 != 0 {
		goto L40
	} else {
		goto L80
	}
L80:
	;
	v251 = v227
	v252 = v225
	v256 = v240
	v257 = v238
	goto L41
L81:
	;
	v465 = base.I32_div_s(v450, int32(8))
	if base.B2i32(v52 == v257)|base.B2i32(v257-v52 < v465) != 0 {
		goto L40
	} else {
		goto L120
	}
L82:
	;
	if v252 != int32(47) {
		goto L40
	} else {
		goto L85
	}
L83:
	;
	goto L84
L84:
	;
	if v257-v52 != int32(4) {
		goto L40
	} else {
		goto L119
	}
L85:
	;
	v266 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v251))))
	if base.B2i32(base.Ui32(int32(9)) < base.Ui32((v266-int32(48))&int32(255)))|base.B2i32(base.Ui32(v257) <= base.Ui32(v52)) != 0 {
		goto L40
	} else {
		goto L86
	}
L86:
	;
	v279 = v266
	v280 = v251
	v281 = int32(0)
	goto L87
L87:
	;
	v293 = int32(_a_F_network_in_5)
	v295 = v279 & int32(255)
	v296 = int32(11)
	goto L92
L88:
	;
	if v409&int32(255) != 0 {
		goto L40
	} else {
		goto L116
	}
L89:
	;
	v403 = v401 - int32(_a_F_network_in_5)
	if base.Ui32(int32(10)) <= base.Ui32(v403) {
		goto L31
	} else {
		goto L114
	}
L90:
	;
	v401 = int32(0)
	goto L89
L91:
	;
	v379 = v372
	v381 = v374
	goto L108
L92:
	;
	goto L99
L99:
	;
	v335 = v295 & int32(255)
	v336 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_network_in[1])))
	if base.B2i32(v335 == v336)|int32(0) == int32(0) {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	v345 = v293
	v347 = v296
	goto L103
L101:
	;
	v365 = v293
	v367 = v296
	goto L102
L102:
	;
	if v367 == int32(0) {
		goto L90
	} else {
		goto L107
	}
L103:
	;
	v351 = *(*int32)(unsafe.Add(mBase, uint32(v345)))
	v352 = v351 ^ v335*int32(16843009)
	v355 = int32(-2139062144)
	if (int32(16843008)-v352|v352)&v355 != v355 {
		v372 = v345
		v374 = v347
		goto L91
	} else {
		goto L105
	}
L104:
	;
	v365 = v360
	v367 = v362
	goto L102
L105:
	;
	v359 = int32(4)
	v360 = v345 + v359
	v362 = v347 - v359
	if base.Ui32(int32(3)) < base.Ui32(v362) {
		v345 = v360
		v347 = v362
		goto L103
	} else {
		goto L106
	}
L106:
	;
	goto L104
L107:
	;
	v372 = v365
	v374 = v367
	goto L91
L108:
	;
	v384 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v379))))
	if v295&int32(255) == v384 {
		goto L110
	} else {
		goto L111
	}
L109:
	;
	goto L90
L110:
	;
	v401 = v379
	goto L89
L111:
	;
	goto L112
L112:
	;
	v386 = int32(1)
	v389 = v381 - v386
	if v389 != 0 {
		v379 = v379 + v386
		v381 = v389
		goto L108
	} else {
		goto L113
	}
L113:
	;
	goto L109
L114:
	;
	v406 = int32(10)
	v408 = v403 + v281*v406
	v409 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v280)+1)))
	if base.Ui32((v409-int32(48))&int32(255)) < base.Ui32(v406) {
		v279 = v409
		v280 = v280 + int32(1)
		v281 = v408
		goto L87
	} else {
		goto L115
	}
L115:
	;
	goto L88
L116:
	;
	if int32(32) < v408 {
		goto L39
	} else {
		goto L117
	}
L117:
	;
	if v408 != int32(-1) {
		v450 = v408
		goto L81
	} else {
		goto L118
	}
L118:
	;
	goto L84
L119:
	;
	v450 = int32(32)
	goto L81
L120:
	;
	if v256 == int32(0) {
		v1294 = v450
		goto L25
	} else {
		goto L121
	}
L121:
	;
	base.MemoryFill(m, v257, int32(0), v256)
	v1323 = v450
	goto L15
L122:
	;
	v755 = v513
	v758 = v514
	v761 = v59
	v763 = v52
	goto L184
L123:
	;
	v517 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v513))))
	if v517|int32(32) != int32(120) {
		goto L122
	} else {
		goto L126
	}
L124:
	;
	goto L125
L125:
	;
	if base.Ui32((v514-int32(48))&int32(255)) <= base.Ui32(int32(9)) {
		goto L122
	} else {
		goto L183
	}
L126:
	;
	v522 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+2)))
	goto L127
L127:
	;
	if base.B2i32(base.Ui32(v522-int32(48)) < base.Ui32(int32(10)))|base.B2i32(base.Ui32(v522|int32(32)-int32(97)) < base.Ui32(int32(6))) == int32(0) {
		goto L122
	} else {
		goto L128
	}
L128:
	;
	v536 = int32(35)
	if v59 == int32(0) {
		v1262 = v536
		goto L27
	} else {
		goto L129
	}
L129:
	;
	if v522 == int32(0) {
		goto L130
	} else {
		goto L131
	}
L130:
	;
	v1262 = int32(44)
	goto L27
L131:
	;
	goto L132
L132:
	;
	v548 = l0 + int32(3)
	v549 = int32(0)
	v550 = v522
	v554 = v59
	v556 = v52
	v559 = int32(0)
	goto L134
L133:
	;
	if v729 == int32(0) {
		v971 = v725
		v974 = v728
		v977 = v730
		v979 = v731
		goto L28
	} else {
		goto L181
	}
L134:
	;
	v562 = base.I32_extend8_s(v550)
	v564 = v550 & int32(255)
	goto L136
L135:
	;
	v725 = v723
	v728 = int32(0)
	v729 = v717
	v730 = v718
	v731 = v719
	v732 = v720
	goto L133
L136:
	;
	if base.B2i32(base.Ui32(v564-int32(48)) < base.Ui32(int32(10)))|base.B2i32(base.Ui32(v564|int32(32)-int32(97)) < base.Ui32(int32(6))) == int32(0) {
		goto L137
	} else {
		goto L138
	}
L137:
	;
	v725 = v548
	v728 = v562
	v729 = v549
	v730 = v554
	v731 = v556
	v732 = v559
	goto L133
L138:
	;
	goto L139
L139:
	;
	v578 = int32(_a_F_network_in_6)
	if base.Ui32((v550-int32(65))&int32(255)) <= base.Ui32(int32(25)) {
		goto L140
	} else {
		goto L141
	}
L140:
	;
	if base.Ui32(v564-int32(65)) < base.Ui32(int32(26)) {
		goto L144
	} else {
		goto L145
	}
L141:
	;
	v592 = v562
	goto L142
L142:
	;
	v593 = int32(17)
	goto L150
L143:
	;
	v592 = v591
	goto L142
L144:
	;
	v591 = v564 | int32(32)
	goto L146
L145:
	;
	v591 = v564
	goto L146
L146:
	;
	goto L143
L147:
	;
	v700 = v698 - int32(_a_F_network_in_6)
	if base.Ui32(int32(16)) <= base.Ui32(v700) {
		goto L30
	} else {
		goto L172
	}
L148:
	;
	v698 = int32(0)
	goto L147
L149:
	;
	v676 = v669
	v678 = v671
	goto L166
L150:
	;
	goto L157
L157:
	;
	v632 = v592 & int32(255)
	v633 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_network_in[2])))
	if base.B2i32(v632 == v633)|int32(0) == int32(0) {
		goto L158
	} else {
		goto L159
	}
L158:
	;
	v642 = v578
	v644 = v593
	goto L161
L159:
	;
	v662 = v578
	v664 = v593
	goto L160
L160:
	;
	if v664 == int32(0) {
		goto L148
	} else {
		goto L165
	}
L161:
	;
	v648 = *(*int32)(unsafe.Add(mBase, uint32(v642)))
	v649 = v648 ^ v632*int32(16843009)
	v652 = int32(-2139062144)
	if (int32(16843008)-v649|v649)&v652 != v652 {
		v669 = v642
		v671 = v644
		goto L149
	} else {
		goto L163
	}
L162:
	;
	v662 = v657
	v664 = v659
	goto L160
L163:
	;
	v656 = int32(4)
	v657 = v642 + v656
	v659 = v644 - v656
	if base.Ui32(int32(3)) < base.Ui32(v659) {
		v642 = v657
		v644 = v659
		goto L161
	} else {
		goto L164
	}
L164:
	;
	goto L162
L165:
	;
	v669 = v662
	v671 = v664
	goto L149
L166:
	;
	v681 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v676))))
	if v592&int32(255) == v681 {
		goto L168
	} else {
		goto L169
	}
L167:
	;
	goto L148
L168:
	;
	v698 = v676
	goto L147
L169:
	;
	goto L170
L170:
	;
	v683 = int32(1)
	v686 = v678 - v683
	if v686 != 0 {
		v676 = v676 + v683
		v678 = v686
		goto L166
	} else {
		goto L171
	}
L171:
	;
	goto L167
L172:
	;
	v705 = v700 | v559<<(uint(int32(4))%32)
	v706 = int32(1)
	if v549 == v706 {
		goto L173
	} else {
		goto L174
	}
L173:
	;
	if v554 == int32(0) {
		v1262 = v536
		goto L27
	} else {
		goto L176
	}
L174:
	;
	v717 = v706
	v718 = v554
	v719 = v556
	goto L175
L175:
	;
	if v549 != 0 {
		goto L177
	} else {
		goto L178
	}
L176:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v556))) = uint8(v705)
	v712 = int32(1)
	v717 = int32(0)
	v718 = v554 - v712
	v719 = v556 + v712
	goto L175
L177:
	;
	v720 = v705
	goto L179
L178:
	;
	v720 = v700
	goto L179
L179:
	;
	v721 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v548))))
	v723 = v548 + int32(1)
	if v721 != 0 {
		v548 = v723
		v549 = v717
		v550 = v721
		v554 = v718
		v556 = v719
		v559 = v720
		goto L134
	} else {
		goto L180
	}
L180:
	;
	goto L135
L181:
	;
	if v730 == int32(0) {
		v1262 = v536
		goto L27
	} else {
		goto L182
	}
L182:
	;
	v738 = v732 << (uint(int32(4)) % 32)
	*(*uint8)(unsafe.Add(mBase, uint32(v731))) = uint8(v738)
	v740 = int32(1)
	v971 = v725
	v974 = v728
	v977 = v730 - v740
	v979 = v731 + v740
	goto L28
L183:
	;
	v1262 = int32(44)
	goto L27
L184:
	;
	v775 = v755
	v776 = int32(0)
	v778 = v758 & int32(255)
	goto L186
L185:
	;
	v1262 = v900
	goto L27
L186:
	;
	goto L192
L187:
	;
	if v761 == int32(0) {
		goto L216
	} else {
		goto L217
	}
L188:
	;
	v897 = v895 - int32(_a_F_network_in_7)
	if base.Ui32(int32(10)) <= base.Ui32(v897) {
		goto L29
	} else {
		goto L213
	}
L189:
	;
	v895 = int32(0)
	goto L188
L190:
	;
	v873 = v866
	v875 = v868
	goto L207
L191:
	;
	if base.B2i32(v812 != v813) == int32(0) {
		goto L189
	} else {
		goto L198
	}
L192:
	;
	v804 = int32(_a_F_network_in_7)
	v806 = int32(11)
	goto L193
L193:
	;
	v809 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v804))))
	if v809 == v778&int32(255) {
		v866 = v804
		v868 = v806
		goto L190
	} else {
		goto L195
	}
L194:
	;
	goto L191
L195:
	;
	v811 = int32(1)
	v812 = v806 - v811
	v813 = int32(0)
	v816 = v804 + v811
	if v816&int32(3) == v813 {
		goto L191
	} else {
		goto L196
	}
L196:
	;
	if v812 != 0 {
		v804 = v816
		v806 = v812
		goto L193
	} else {
		goto L197
	}
L197:
	;
	goto L194
L198:
	;
	v829 = v778 & int32(255)
	v830 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v816))))
	if base.B2i32(v829 == v830)|base.B2i32(base.Ui32(v812) < base.Ui32(int32(4))) == int32(0) {
		goto L199
	} else {
		goto L200
	}
L199:
	;
	v839 = v816
	v841 = v812
	goto L202
L200:
	;
	v859 = v816
	v861 = v812
	goto L201
L201:
	;
	if v861 == int32(0) {
		goto L189
	} else {
		goto L206
	}
L202:
	;
	v845 = *(*int32)(unsafe.Add(mBase, uint32(v839)))
	v846 = v845 ^ v829*int32(16843009)
	v849 = int32(-2139062144)
	if (int32(16843008)-v846|v846)&v849 != v849 {
		v866 = v839
		v868 = v841
		goto L190
	} else {
		goto L204
	}
L203:
	;
	v859 = v854
	v861 = v856
	goto L201
L204:
	;
	v853 = int32(4)
	v854 = v839 + v853
	v856 = v841 - v853
	if base.Ui32(int32(3)) < base.Ui32(v856) {
		v839 = v854
		v841 = v856
		goto L202
	} else {
		goto L205
	}
L205:
	;
	goto L203
L206:
	;
	v866 = v859
	v868 = v861
	goto L190
L207:
	;
	v878 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v873))))
	if v778&int32(255) == v878 {
		goto L209
	} else {
		goto L210
	}
L208:
	;
	goto L189
L209:
	;
	v895 = v873
	goto L188
L210:
	;
	goto L211
L211:
	;
	v880 = int32(1)
	v883 = v875 - v880
	if v883 != 0 {
		v873 = v873 + v880
		v875 = v883
		goto L207
	} else {
		goto L212
	}
L212:
	;
	goto L208
L213:
	;
	v900 = int32(44)
	v903 = v897 + v776*int32(10)
	if int32(255) < v903 {
		v1262 = v900
		goto L27
	} else {
		goto L214
	}
L214:
	;
	v907 = v775 + int32(1)
	v908 = int32(*(*int8)(unsafe.Add(mBase, uint32(v775))))
	if base.Ui32((v908-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v775 = v907
		v776 = v903
		v778 = v908
		goto L186
	} else {
		goto L215
	}
L215:
	;
	goto L187
L216:
	;
	v1262 = int32(35)
	goto L27
L217:
	;
	goto L218
L218:
	;
	v918 = int32(1)
	v919 = v761 - v918
	*(*uint8)(unsafe.Add(mBase, uint32(v763))) = uint8(v903)
	v922 = v763 + v918
	v924 = v908 & int32(255)
	if v924 != int32(46) {
		goto L219
	} else {
		goto L220
	}
L219:
	;
	if v924 == int32(0) {
		v971 = v907
		v974 = v908
		v977 = v919
		v979 = v922
		goto L28
	} else {
		goto L222
	}
L220:
	;
	goto L221
L221:
	;
	v933 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v775)+1)))
	if base.Ui32((v933-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v755 = v775 + int32(2)
		v758 = v933
		v761 = v919
		v763 = v922
		goto L184
	} else {
		goto L224
	}
L222:
	;
	if v924 != int32(47) {
		v1262 = v900
		goto L27
	} else {
		goto L223
	}
L223:
	;
	v971 = v907
	v974 = v908
	v977 = v919
	v979 = v922
	goto L28
L224:
	;
	goto L185
L225:
	;
	v943 = F_inet_cidr_pton_ipv6(m, l0, v52, int32(16))
	mBase = m.M
	v1323 = v943
	goto L15
L226:
	;
	goto L227
L227:
	;
	v944 = F_inet_cidr_pton_ipv6(m, l0, v52, v59)
	mBase = m.M
	v1323 = v944
	goto L15
L228:
	;
	if v1209 <= v1207 {
		v1294 = v1209
		goto L25
	} else {
		goto L283
	}
L229:
	;
	v1176 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52))))
	if base.Ui32(int32(239)) < base.Ui32(v1176) {
		v1191 = int32(32)
		goto L269
	} else {
		goto L270
	}
L230:
	;
	m.Env.X__assert_fail(m, int32(_a_F_network_in_0), int32(_a_F_network_in_1), int32(181), int32(_a_F_network_in_2))
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L231:
	;
	if v1130 == int32(-1) {
		goto L229
	} else {
		goto L268
	}
L232:
	;
	v987 = int32(44)
	v988 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v971))))
	if base.B2i32(base.Ui32(int32(9)) < base.Ui32((v988-int32(48))&int32(255)))|base.B2i32(base.Ui32(v979) <= base.Ui32(v52)) != 0 {
		v1262 = v987
		goto L27
	} else {
		goto L235
	}
L233:
	;
	goto L234
L234:
	;
	if v974|base.B2i32(v52 == v979) != 0 {
		v1262 = int32(44)
		goto L27
	} else {
		goto L267
	}
L235:
	;
	v1001 = v971
	v1003 = int32(0)
	v1004 = v988
	goto L236
L236:
	;
	v1017 = v1004 & int32(255)
	goto L242
L237:
	;
	if v1131&int32(255) != 0 {
		v1262 = v987
		goto L27
	} else {
		goto L265
	}
L238:
	;
	v1125 = v1123 - int32(_a_F_network_in_7)
	if base.Ui32(int32(10)) <= base.Ui32(v1125) {
		goto L230
	} else {
		goto L263
	}
L239:
	;
	v1123 = int32(0)
	goto L238
L240:
	;
	v1101 = v1094
	v1103 = v1096
	goto L257
L241:
	;
	if base.B2i32(v1040 != v1041) == int32(0) {
		goto L239
	} else {
		goto L248
	}
L242:
	;
	v1032 = int32(_a_F_network_in_7)
	v1034 = int32(11)
	goto L243
L243:
	;
	v1037 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1032))))
	if v1037 == v1017&int32(255) {
		v1094 = v1032
		v1096 = v1034
		goto L240
	} else {
		goto L245
	}
L244:
	;
	goto L241
L245:
	;
	v1039 = int32(1)
	v1040 = v1034 - v1039
	v1041 = int32(0)
	v1044 = v1032 + v1039
	if v1044&int32(3) == v1041 {
		goto L241
	} else {
		goto L246
	}
L246:
	;
	if v1040 != 0 {
		v1032 = v1044
		v1034 = v1040
		goto L243
	} else {
		goto L247
	}
L247:
	;
	goto L244
L248:
	;
	v1057 = v1017 & int32(255)
	v1058 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1044))))
	if base.B2i32(v1057 == v1058)|base.B2i32(base.Ui32(v1040) < base.Ui32(int32(4))) == int32(0) {
		goto L249
	} else {
		goto L250
	}
L249:
	;
	v1067 = v1044
	v1069 = v1040
	goto L252
L250:
	;
	v1087 = v1044
	v1089 = v1040
	goto L251
L251:
	;
	if v1089 == int32(0) {
		goto L239
	} else {
		goto L256
	}
L252:
	;
	v1073 = *(*int32)(unsafe.Add(mBase, uint32(v1067)))
	v1074 = v1073 ^ v1057*int32(16843009)
	v1077 = int32(-2139062144)
	if (int32(16843008)-v1074|v1074)&v1077 != v1077 {
		v1094 = v1067
		v1096 = v1069
		goto L240
	} else {
		goto L254
	}
L253:
	;
	v1087 = v1082
	v1089 = v1084
	goto L251
L254:
	;
	v1081 = int32(4)
	v1082 = v1067 + v1081
	v1084 = v1069 - v1081
	if base.Ui32(int32(3)) < base.Ui32(v1084) {
		v1067 = v1082
		v1069 = v1084
		goto L252
	} else {
		goto L255
	}
L255:
	;
	goto L253
L256:
	;
	v1094 = v1087
	v1096 = v1089
	goto L240
L257:
	;
	v1106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1101))))
	if v1017&int32(255) == v1106 {
		goto L259
	} else {
		goto L260
	}
L258:
	;
	goto L239
L259:
	;
	v1123 = v1101
	goto L238
L260:
	;
	goto L261
L261:
	;
	v1108 = int32(1)
	v1111 = v1103 - v1108
	if v1111 != 0 {
		v1101 = v1101 + v1108
		v1103 = v1111
		goto L257
	} else {
		goto L262
	}
L262:
	;
	goto L258
L263:
	;
	v1128 = int32(10)
	v1130 = v1125 + v1003*v1128
	v1131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1001)+1)))
	if base.Ui32((v1131-int32(48))&int32(255)) < base.Ui32(v1128) {
		v1001 = v1001 + int32(1)
		v1003 = v1130
		v1004 = v1131
		goto L236
	} else {
		goto L264
	}
L264:
	;
	goto L237
L265:
	;
	if v1130 <= int32(32) {
		goto L231
	} else {
		goto L266
	}
L266:
	;
	v1262 = int32(35)
	goto L27
L267:
	;
	goto L229
L268:
	;
	v1207 = (v979 - v52) << (uint(int32(3)) % 32)
	v1209 = v1130
	goto L228
L269:
	;
	v1194 = (v979 - v52) << (uint(int32(3)) % 32)
	if v1194 < v1191 {
		goto L276
	} else {
		goto L277
	}
L270:
	;
	if base.Ui32(int32(223)) < base.Ui32(v1176) {
		v1191 = int32(8)
		goto L269
	} else {
		goto L271
	}
L271:
	;
	if base.Ui32(int32(191)) < base.Ui32(v1176) {
		v1191 = int32(24)
		goto L269
	} else {
		goto L272
	}
L272:
	;
	if base.I32_extend8_s(v1176) < int32(0) {
		goto L273
	} else {
		goto L274
	}
L273:
	;
	v1190 = int32(16)
	goto L275
L274:
	;
	v1190 = int32(8)
	goto L275
L275:
	;
	v1191 = v1190
	goto L269
L276:
	;
	v1196 = v1191
	goto L278
L277:
	;
	v1196 = v1194
	goto L278
L278:
	;
	if v1196 != int32(8) {
		v1207 = v1194
		v1209 = v1196
		goto L228
	} else {
		goto L279
	}
L279:
	;
	if v1176 == int32(224) {
		goto L280
	} else {
		goto L281
	}
L280:
	;
	v1203 = int32(4)
	goto L282
L281:
	;
	v1203 = int32(8)
	goto L282
L282:
	;
	v1207 = v1194
	v1209 = v1203
	goto L228
L283:
	;
	v1231 = v977
	v1233 = v979
	goto L284
L284:
	;
	if v1231 == int32(0) {
		goto L286
	} else {
		goto L287
	}
L285:
	;
	v1294 = v1209
	goto L25
L286:
	;
	v1262 = int32(35)
	goto L27
L287:
	;
	goto L288
L288:
	;
	v1242 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1233))) = uint8(v1242)
	v1244 = int32(1)
	v1247 = v1233 + v1244
	if (v1247-v52)<<(uint(int32(3))%32) < v1209 {
		v1231 = v1231 - v1244
		v1233 = v1247
		goto L284
	} else {
		goto L289
	}
L289:
	;
	goto L285
L290:
	;
	v1328 = int32(1)
	v1330 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23))))
	v1332 = v1330 & v1328
	if v1332 != 0 {
		goto L293
	} else {
		goto L294
	}
L291:
	;
	goto L292
L292:
	;
	v1343 = F_errsave_start(m, l2)
	mBase = m.M
	v1344 = m.ExcPending
	if v1344 != 0 {
		goto L1
	} else {
		goto L300
	}
L293:
	;
	v1333 = v1328
	goto L295
L294:
	;
	v1333 = int32(4)
	goto L295
L295:
	;
	v1335 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23+v1333))))
	if v1335 == int32(2) {
		goto L296
	} else {
		goto L297
	}
L296:
	;
	v1338 = int32(32)
	goto L298
L297:
	;
	v1338 = int32(128)
	goto L298
L298:
	;
	if base.Ui32(v1323) <= base.Ui32(v1338) {
		goto L14
	} else {
		goto L299
	}
L299:
	;
	goto L292
L300:
	;
	if v1343 == int32(0) {
		goto L301
	} else {
		goto L302
	}
L301:
	;
	v1496 = int32(0)
	goto L13
L302:
	;
	goto L303
L303:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v1350 = m.ExcPending
	if v1350 != 0 {
		goto L1
	} else {
		goto L304
	}
L304:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+4)) = l0
	if l1 != 0 {
		goto L305
	} else {
		goto L306
	}
L305:
	;
	v1354 = int32(_a_F_network_in_8)
	goto L307
L306:
	;
	v1354 = int32(_a_F_network_in_9)
	goto L307
L307:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v1354
	F_errmsg(m, int32(_a_F_network_in_10), v20)
	mBase = m.M
	v1358 = m.ExcPending
	if v1358 != 0 {
		goto L1
	} else {
		goto L308
	}
L308:
	;
	F_errsave_finish(m, l2, int32(_a_F_network_in_11), int32(100), int32(_a_F_network_in_12))
	mBase = m.M
	v1363 = m.ExcPending
	if v1363 != 0 {
		goto L1
	} else {
		goto L309
	}
L309:
	;
	v1496 = int32(0)
	goto L13
L310:
	;
	v1365 = v47
	goto L312
L311:
	;
	v1365 = v49
	goto L312
L312:
	;
	if base.B2i32(l1 == int32(0))|base.B2i32(v1338 == v1323) != 0 {
		goto L313
	} else {
		goto L314
	}
L313:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1365)+1)) = uint8(v1323)
	v1472 = int32(1)
	v1474 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23))))
	if v1474&v1472 != 0 {
		goto L332
	} else {
		goto L333
	}
L314:
	;
	v1371 = int32(base.Ui32(v1323) >> (uint(int32(3)) % 32))
	if v1335 == int32(2) {
		goto L315
	} else {
		goto L316
	}
L315:
	;
	v1376 = int32(4)
	goto L317
L316:
	;
	v1376 = int32(16)
	goto L317
L317:
	;
	if base.Ui32(v1376) <= base.Ui32(v1371) {
		goto L313
	} else {
		goto L318
	}
L318:
	;
	v1379 = v1365 + int32(2)
	v1381 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1379+v1371))))
	if v1381<<(uint(v1323&int32(7))%32)&int32(255) != 0 {
		goto L319
	} else {
		goto L320
	}
L319:
	;
	v1429 = int32(0)
	v1430 = F_errsave_start(m, l2)
	mBase = m.M
	v1431 = m.ExcPending
	if v1431 != 0 {
		goto L1
	} else {
		goto L326
	}
L320:
	;
	v1388 = v1371 + int32(1)
	if v1388 == v1376 {
		goto L313
	} else {
		goto L321
	}
L321:
	;
	v1391 = v1388
	goto L322
L322:
	;
	v1408 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1391+v1379))))
	if v1408 != 0 {
		goto L319
	} else {
		goto L324
	}
L323:
	;
	goto L313
L324:
	;
	v1410 = v1391 + int32(1)
	if v1376 != v1410 {
		v1391 = v1410
		goto L322
	} else {
		goto L325
	}
L325:
	;
	goto L323
L326:
	;
	if v1430 == int32(0) {
		v1496 = v1429
		goto L13
	} else {
		goto L327
	}
L327:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v1436 = m.ExcPending
	if v1436 != 0 {
		goto L1
	} else {
		goto L328
	}
L328:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = l0
	F_errmsg(m, int32(_a_F_network_in_13), v20+int32(16))
	mBase = m.M
	v1442 = m.ExcPending
	if v1442 != 0 {
		goto L1
	} else {
		goto L329
	}
L329:
	;
	F_errdetail(m, int32(_a_F_network_in_14), int32(0))
	mBase = m.M
	v1446 = m.ExcPending
	if v1446 != 0 {
		goto L1
	} else {
		goto L330
	}
L330:
	;
	F_errsave_finish(m, l2, int32(_a_F_network_in_11), int32(111), int32(_a_F_network_in_12))
	mBase = m.M
	v1451 = m.ExcPending
	if v1451 != 0 {
		goto L1
	} else {
		goto L331
	}
L331:
	;
	v1496 = v1429
	goto L13
L332:
	;
	v1477 = v1472
	goto L334
L333:
	;
	v1477 = int32(4)
	goto L334
L334:
	;
	v1479 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23+v1477))))
	if v1479 == int32(2) {
		goto L335
	} else {
		goto L336
	}
L335:
	;
	v1482 = int32(40)
	goto L337
L336:
	;
	v1482 = int32(88)
	goto L337
L337:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23))) = v1482
	v1496 = v23
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
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
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
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v154 int32
	_ = v154
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
		goto L5
	} else {
		goto L6
	}
L4:
	;
	return int32(0)
L5:
	;
	v18 = v13
	goto L7
L6:
	;
	v18 = int32(4)
	goto L7
L7:
	;
	v19 = v6 + v18
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
	v21 = int32(1)
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
	if v23&v21 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v26 = v21
	goto L10
L9:
	;
	v26 = int32(4)
	goto L10
L10:
	;
	v27 = v11 + v26
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27))))
	if v20 != v28 {
		goto L4
	} else {
		goto L11
	}
L11:
	;
	v30 = int32(2)
	v31 = v19 + v30
	v33 = v27 + v30
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+1)))
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+1)))
	if base.Ui32(v34) < base.Ui32(v35) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v37 = v34
	goto L14
L13:
	;
	v37 = v35
	goto L14
L14:
	;
	v39 = int32(base.Ui32(v37) >> (uint(int32(3)) % 32))
	if base.Ui32(int32(4)) <= base.Ui32(v39) {
		goto L18
	} else {
		goto L19
	}
L15:
	;
	if v101 != 0 {
		goto L4
	} else {
		goto L33
	}
L16:
	;
	v101 = int32(0)
	goto L15
L17:
	;
	v75 = v70
	v76 = v71
	v77 = v72
	goto L27
L18:
	;
	if (v31|v33)&int32(3) != 0 {
		v70 = v31
		v71 = v33
		v72 = v39
		goto L17
	} else {
		goto L21
	}
L19:
	;
	v63 = v31
	v64 = v33
	v65 = v39
	goto L20
L20:
	;
	if v65 == int32(0) {
		goto L16
	} else {
		goto L26
	}
L21:
	;
	v47 = v31
	v48 = v33
	v49 = v39
	goto L22
L22:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
	if v52 != v53 {
		v70 = v47
		v71 = v48
		v72 = v49
		goto L17
	} else {
		goto L24
	}
L23:
	;
	v63 = v58
	v64 = v56
	v65 = v60
	goto L20
L24:
	;
	v55 = int32(4)
	v56 = v48 + v55
	v58 = v47 + v55
	v60 = v49 - v55
	if base.Ui32(int32(3)) < base.Ui32(v60) {
		v47 = v58
		v48 = v56
		v49 = v60
		goto L22
	} else {
		goto L25
	}
L25:
	;
	goto L23
L26:
	;
	v70 = v63
	v71 = v64
	v72 = v65
	goto L17
L27:
	;
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75))))
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76))))
	if v80 == v81 {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	v101 = v80 - v81
	goto L15
L29:
	;
	v83 = int32(1)
	v88 = v77 - v83
	if v88 != 0 {
		v75 = v75 + v83
		v76 = v76 + v83
		v77 = v88
		goto L27
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	goto L28
L32:
	;
	goto L16
L33:
	;
	v103 = v37 & int32(7)
	if v103 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	return int32(1)
L35:
	;
	goto L36
L36:
	;
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39+v31))))
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39+v33))))
	v112 = v109 ^ v111
	if base.Ui32(int32(127)) < base.Ui32(v112) {
		goto L4
	} else {
		goto L37
	}
L37:
	;
	v115 = int32(1)
	if v103 == v115 {
		v154 = v115
		goto L38
	} else {
		goto L39
	}
L38:
	;
	return v154
L39:
	;
	if v112<<(uint(int32(1))%32)&int32(128) != 0 {
		goto L4
	} else {
		goto L40
	}
L40:
	;
	if base.Ui32(v103) < base.Ui32(int32(3)) {
		v154 = v115
		goto L38
	} else {
		goto L41
	}
L41:
	;
	if v112<<(uint(int32(2))%32)&int32(128) != 0 {
		goto L4
	} else {
		goto L42
	}
L42:
	;
	if v103 == int32(3) {
		v154 = v115
		goto L38
	} else {
		goto L43
	}
L43:
	;
	if v112<<(uint(int32(3))%32)&int32(128) != 0 {
		goto L4
	} else {
		goto L44
	}
L44:
	;
	if base.Ui32(v103) < base.Ui32(int32(5)) {
		v154 = v115
		goto L38
	} else {
		goto L45
	}
L45:
	;
	if v112<<(uint(int32(4))%32)&int32(128) != 0 {
		goto L4
	} else {
		goto L46
	}
L46:
	;
	if v103 == int32(5) {
		v154 = v115
		goto L38
	} else {
		goto L47
	}
L47:
	;
	if v112<<(uint(int32(5))%32)&int32(128) != 0 {
		v154 = int32(0)
		goto L38
	} else {
		goto L48
	}
L48:
	;
	if v103 != int32(7) {
		v154 = int32(1)
		goto L38
	} else {
		goto L49
	}
L49:
	;
	v154 = base.B2i32(v112&int32(2) == int32(0))
	goto L38
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
	var v43 int32
	_ = v43
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
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	v7 = m.G0
	v9 = v7 - int32(80)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = F_pg_detoast_datum_packed(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v16 = int32(1)
		v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
		v20 = v18 & v16
		if v20 != 0 {
			v21 = v16
		} else {
			v21 = int32(4)
		}
		v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12+v21))))
		v25 = v12 + int32(1)
		v27 = v12 + int32(4)
		if v20 != 0 {
			v28 = v25
		} else {
			v28 = v27
		}
		v29 = int32(2)
		if v23 == v29 {
			v35 = int32(32)
		} else {
			v35 = int32(128)
		}
		v37 = v9 + int32(16)
		v38 = F_pg_inet_net_ntop(m, v23, v28+v29, v35, v37)
		mBase = m.M
		v39 = m.ExcPending
		if v39 != 0 {
			return int32(0)
		} else {
			if v38 != 0 {
				v40 = int32(47)
				v41 = F___strchrnul(m, v37, v40)
				mBase = m.M
				v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41))))
				if v43 == v40 {
					v47 = v41
				} else {
					v47 = int32(0)
				}
				if v47 == int32(0) {
					v50 = F_strlen(m, v37)
					mBase = m.M
					v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
					if v51&int32(1) != 0 {
						v54 = v25
					} else {
						v54 = v27
					}
					v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54)+1)))
					*(*int32)(unsafe.Add(mBase, uint32(v9))) = v55
					v61 = F_pg_snprintf(m, v37+v50, int32(50)-v50, int32(_a_F_network_show_0), v9)
					mBase = m.M
					v62 = m.ExcPending
					if v62 != 0 {
						return int32(0)
					} else {
						v66 = F_cstring_to_text(m, v9+int32(16))
						mBase = m.M
						v67 = m.ExcPending
						if v67 != 0 {
							return int32(0)
						} else {
							m.G0 = v9 + int32(80)
							return v66
						}
					}
				} else {
					v66 = F_cstring_to_text(m, v9+int32(16))
					mBase = m.M
					v67 = m.ExcPending
					if v67 != 0 {
						return int32(0)
					} else {
						m.G0 = v9 + int32(80)
						return v66
					}
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v75 = m.ExcPending
				if v75 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(50462850))
					mBase = m.M
					v78 = m.ExcPending
					if v78 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(_a_F_network_show_1), int32(0))
						mBase = m.M
						v82 = m.ExcPending
						if v82 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_network_show_2), int32(1174), int32(_a_F_network_show_3))
							mBase = m.M
							v87 = m.ExcPending
							if v87 != 0 {
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
}
func F_network_sortsupport(m *base.Module, l0 int32) int32 {
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v5 = Fn13942(m, l0, int32(1445), int32(1444), int32(1443))
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
