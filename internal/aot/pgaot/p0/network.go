package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_network_broadcast(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
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
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v118 int32
	_ = v118
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v12 = F_pg_detoast_datum_packed(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int64(0)
	} else {
		v17 = F_palloc0(m, int32(22))
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int64(0)
		} else {
			v19 = int32(4)
			v21 = int32(1)
			v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
			if v23&v21 != 0 {
				v26 = v21
			} else {
				v26 = v19
			}
			v27 = v12 + v26
			v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27))))
			if v28 == int32(2) {
				v31 = v19
			} else {
				v31 = int32(16)
			}
			v32 = int32(1)
			v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
			if v34&v32 != 0 {
				v37 = v32
			} else {
				v37 = int32(4)
			}
			v38 = v17 + v37
			v39 = int32(2)
			v40 = v38 + v39
			v42 = v27 + v39
			v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+1)))
			v45 = int32(0)
			v47 = v43
			for {
				if base.Ui32(int32(8)) <= base.Ui32(v47) {
					v67 = v47 - int32(8)
					v68 = int32(0)
				} else {
					v61 = int32(0)
					if v47 == v61 {
						v67 = v61
						v68 = int32(255)
					} else {
						v67 = v61
						v68 = int32(base.Ui32(int32(255)) >> (uint(v47) % 32))
					}
				}
				v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45+v42))))
				v71 = v68 | v70
				*(*uint8)(unsafe.Add(mBase, uint32(v45+v40))) = uint8(v71)
				v74 = v45 | int32(1)
				if base.Ui32(v67) <= base.Ui32(int32(7)) {
					v78 = int32(0)
					if v67 == v78 {
						v87 = v78
						v88 = int32(255)
					} else {
						v87 = v78
						v88 = int32(base.Ui32(int32(255)) >> (uint(v67) % 32))
					}
				} else {
					v87 = v67 - int32(8)
					v88 = int32(0)
				}
				v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42+v74))))
				v91 = v88 | v90
				*(*uint8)(unsafe.Add(mBase, uint32(v40+v74))) = uint8(v91)
				v94 = v45 + int32(2)
				if v94 != v31 {
					v45 = v94
					v47 = v87
					continue
				} else {
					break
				}
				break
			}
			v96 = int32(1)
			v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
			if v98&v96 != 0 {
				v101 = v96
			} else {
				v101 = int32(4)
			}
			v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12+v101))))
			*(*uint8)(unsafe.Add(mBase, uint32(v38))) = uint8(v103)
			v105 = int32(1)
			v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
			if v107&v105 != 0 {
				v110 = v105
			} else {
				v110 = int32(4)
			}
			v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12+v110)+1)))
			*(*uint8)(unsafe.Add(mBase, uint32(v38)+1)) = uint8(v112)
			if v103 == int32(2) {
				v118 = int32(40)
			} else {
				v118 = int32(88)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v17))) = v118
			return base.I64_extend_i32_u(v17)
		}
	}
}
func F_network_in(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v251 int32
	_ = v251
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v338 int32
	_ = v338
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v348 int32
	_ = v348
	var v350 int32
	_ = v350
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v367 int32
	_ = v367
	var v369 int32
	_ = v369
	var v372 int32
	_ = v372
	var v384 int32
	_ = v384
	var v386 int32
	_ = v386
	var v389 int32
	_ = v389
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v430 int32
	_ = v430
	var v444 int32
	_ = v444
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v492 int32
	_ = v492
	var v497 int32
	_ = v497
	var v511 int32
	_ = v511
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v534 int32
	_ = v534
	var v536 int32
	_ = v536
	var v550 int32
	_ = v550
	var v559 int32
	_ = v559
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v611 int32
	_ = v611
	var v613 int32
	_ = v613
	var v617 int32
	_ = v617
	var v618 int32
	_ = v618
	var v621 int32
	_ = v621
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v628 int32
	_ = v628
	var v631 int32
	_ = v631
	var v633 int32
	_ = v633
	var v638 int32
	_ = v638
	var v640 int32
	_ = v640
	var v645 int32
	_ = v645
	var v647 int32
	_ = v647
	var v650 int32
	_ = v650
	var v652 int32
	_ = v652
	var v655 int32
	_ = v655
	var v667 int32
	_ = v667
	var v669 int32
	_ = v669
	var v674 int32
	_ = v674
	var v675 int32
	_ = v675
	var v681 int32
	_ = v681
	var v686 int32
	_ = v686
	var v687 int32
	_ = v687
	var v688 int32
	_ = v688
	var v689 int32
	_ = v689
	var v690 int32
	_ = v690
	var v692 int32
	_ = v692
	var v694 int32
	_ = v694
	var v695 int32
	_ = v695
	var v698 int32
	_ = v698
	var v699 int32
	_ = v699
	var v700 int32
	_ = v700
	var v701 int32
	_ = v701
	var v702 int32
	_ = v702
	var v708 int32
	_ = v708
	var v710 int32
	_ = v710
	var v725 int32
	_ = v725
	var v728 int32
	_ = v728
	var v729 int32
	_ = v729
	var v730 int32
	_ = v730
	var v743 int32
	_ = v743
	var v745 int32
	_ = v745
	var v746 int32
	_ = v746
	var v770 int32
	_ = v770
	var v772 int32
	_ = v772
	var v775 int32
	_ = v775
	var v777 int32
	_ = v777
	var v778 int32
	_ = v778
	var v779 int32
	_ = v779
	var v782 int32
	_ = v782
	var v795 int32
	_ = v795
	var v796 int32
	_ = v796
	var v805 int32
	_ = v805
	var v807 int32
	_ = v807
	var v811 int32
	_ = v811
	var v812 int32
	_ = v812
	var v815 int32
	_ = v815
	var v819 int32
	_ = v819
	var v820 int32
	_ = v820
	var v822 int32
	_ = v822
	var v825 int32
	_ = v825
	var v827 int32
	_ = v827
	var v832 int32
	_ = v832
	var v834 int32
	_ = v834
	var v839 int32
	_ = v839
	var v841 int32
	_ = v841
	var v844 int32
	_ = v844
	var v846 int32
	_ = v846
	var v849 int32
	_ = v849
	var v861 int32
	_ = v861
	var v863 int32
	_ = v863
	var v866 int32
	_ = v866
	var v869 int32
	_ = v869
	var v873 int32
	_ = v873
	var v874 int32
	_ = v874
	var v884 int32
	_ = v884
	var v885 int32
	_ = v885
	var v888 int32
	_ = v888
	var v890 int32
	_ = v890
	var v899 int32
	_ = v899
	var v909 int32
	_ = v909
	var v910 int32
	_ = v910
	var v937 int32
	_ = v937
	var v939 int32
	_ = v939
	var v941 int32
	_ = v941
	var v942 int32
	_ = v942
	var v951 int32
	_ = v951
	var v953 int32
	_ = v953
	var v965 int32
	_ = v965
	var v966 int32
	_ = v966
	var v968 int32
	_ = v968
	var v979 int32
	_ = v979
	var v994 int32
	_ = v994
	var v996 int32
	_ = v996
	var v999 int32
	_ = v999
	var v1001 int32
	_ = v1001
	var v1002 int32
	_ = v1002
	var v1003 int32
	_ = v1003
	var v1006 int32
	_ = v1006
	var v1019 int32
	_ = v1019
	var v1020 int32
	_ = v1020
	var v1029 int32
	_ = v1029
	var v1031 int32
	_ = v1031
	var v1035 int32
	_ = v1035
	var v1036 int32
	_ = v1036
	var v1039 int32
	_ = v1039
	var v1043 int32
	_ = v1043
	var v1044 int32
	_ = v1044
	var v1046 int32
	_ = v1046
	var v1049 int32
	_ = v1049
	var v1051 int32
	_ = v1051
	var v1056 int32
	_ = v1056
	var v1058 int32
	_ = v1058
	var v1063 int32
	_ = v1063
	var v1065 int32
	_ = v1065
	var v1068 int32
	_ = v1068
	var v1070 int32
	_ = v1070
	var v1073 int32
	_ = v1073
	var v1085 int32
	_ = v1085
	var v1087 int32
	_ = v1087
	var v1090 int32
	_ = v1090
	var v1092 int32
	_ = v1092
	var v1093 int32
	_ = v1093
	var v1136 int32
	_ = v1136
	var v1150 int32
	_ = v1150
	var v1151 int32
	_ = v1151
	var v1154 int32
	_ = v1154
	var v1156 int32
	_ = v1156
	var v1163 int32
	_ = v1163
	var v1167 int32
	_ = v1167
	var v1168 int32
	_ = v1168
	var v1187 int32
	_ = v1187
	var v1188 int32
	_ = v1188
	var v1198 int32
	_ = v1198
	var v1200 int32
	_ = v1200
	var v1203 int32
	_ = v1203
	var v1219 int32
	_ = v1219
	var v1245 int32
	_ = v1245
	var v1271 int32
	_ = v1271
	var v1276 int32
	_ = v1276
	var v1278 int32
	_ = v1278
	var v1281 int32
	_ = v1281
	var v1282 int32
	_ = v1282
	var v1283 int32
	_ = v1283
	var v1286 int32
	_ = v1286
	var v1291 int32
	_ = v1291
	var v1292 int32
	_ = v1292
	var v1298 int32
	_ = v1298
	var v1302 int32
	_ = v1302
	var v1306 int32
	_ = v1306
	var v1311 int32
	_ = v1311
	var v1318 int32
	_ = v1318
	var v1323 int32
	_ = v1323
	var v1326 int32
	_ = v1326
	var v1328 int32
	_ = v1328
	var v1335 int32
	_ = v1335
	var v1338 int32
	_ = v1338
	var v1353 int32
	_ = v1353
	var v1355 int32
	_ = v1355
	var v1372 int32
	_ = v1372
	var v1373 int32
	_ = v1373
	var v1374 int32
	_ = v1374
	var v1379 int32
	_ = v1379
	var v1385 int32
	_ = v1385
	var v1388 int32
	_ = v1388
	var v1389 int32
	_ = v1389
	var v1394 int32
	_ = v1394
	var v1415 int32
	_ = v1415
	var v1429 int32
	_ = v1429
	v16 = m.G0
	v18 = v16 - int32(32)
	m.G0 = v18
	v21 = F_palloc0(m, int32(22))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v25 = int32(1)
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
	if v27&v25 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v30 = v25
	goto L5
L4:
	;
	v30 = int32(4)
	goto L5
L5:
	;
	v31 = v21 + v30
	v34 = int32(58)
	v35 = F___strchrnul(m, l0, v34)
	mBase = m.M
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35))))
	if v37 == v34 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	if v41 != 0 {
		goto L10
	} else {
		goto L11
	}
L7:
	;
	v41 = v35
	goto L9
L8:
	;
	v41 = int32(0)
	goto L9
L9:
	;
	goto L6
L10:
	;
	v42 = int32(3)
	goto L12
L11:
	;
	v42 = int32(2)
	goto L12
L12:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v31))) = uint8(v42)
	v45 = v31 + int32(2)
	if v41 != 0 {
		goto L16
	} else {
		goto L17
	}
L13:
	;
	m.G0 = v18 + int32(32)
	return v1429
L14:
	;
	if base.B2i32(l1 == int32(0))|base.B2i32(v1271 == v1286) != 0 {
		goto L303
	} else {
		goto L304
	}
L15:
	;
	if int32(0) <= v1271 {
		goto L283
	} else {
		goto L284
	}
L16:
	;
	v48 = int32(16)
	goto L18
L17:
	;
	v48 = int32(4)
	goto L18
L18:
	;
	if l1 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v50 = v48
	goto L21
L20:
	;
	v50 = int32(-1)
	goto L21
L21:
	;
	switch v42 - int32(2) {
	case 0:
		goto L32
	case 1:
		goto L31
	default:
		goto L30
	}
L22:
	;
	v1271 = v1245
	goto L15
L23:
	;
	v1245 = int32(-1)
	goto L22
L24:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_network_in[0])) = v1219
	goto L23
L25:
	;
	if v939 == int32(47) {
		goto L225
	} else {
		goto L226
	}
L26:
	;
	m.Env.X__assert_fail(m, int32(_a_F_network_in_0), int32(_a_F_network_in_1), int32(149), int32(_a_F_network_in_2))
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L27:
	;
	m.Env.X__assert_fail(m, int32(_a_F_network_in_3), int32(_a_F_network_in_1), int32(120), int32(_a_F_network_in_2))
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L28:
	;
	m.Env.X__assert_fail(m, int32(_a_F_network_in_0), int32(_a_F_network_in_1), int32(300), int32(_a_F_network_in_4))
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L29:
	;
	m.Env.X__assert_fail(m, int32(_a_F_network_in_0), int32(_a_F_network_in_1), int32(275), int32(_a_F_network_in_4))
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_network_in[0])) = int32(5)
	goto L23
L31:
	;
	if v50 == int32(-1) {
		goto L218
	} else {
		goto L219
	}
L32:
	;
	if v50 == int32(-1) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v61 = l0
	v64 = int32(4)
	v67 = v45
	goto L39
L34:
	;
	goto L35
L35:
	;
	v488 = l0 + int32(1)
	v489 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v489 == int32(48) {
		goto L120
	} else {
		goto L121
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_network_in[0])) = int32(35)
	goto L23
L37:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_network_in[0])) = int32(44)
	goto L23
L38:
	;
	if v238 != 0 {
		goto L79
	} else {
		goto L80
	}
L39:
	;
	v72 = v61 + int32(1)
	v74 = int32(*(*int8)(unsafe.Add(mBase, uint32(v61))))
	if base.Ui32(int32(9)) < base.Ui32((v74-int32(48))&int32(255)) {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	if v229 == int32(47) {
		v238 = v212
		v239 = v214
		v242 = v227
		v245 = v225
		goto L38
	} else {
		goto L76
	}
L41:
	;
	v238 = v74
	v239 = v72
	v242 = v64
	v245 = v67
	goto L38
L42:
	;
	goto L43
L43:
	;
	v84 = v72
	v85 = v74
	v87 = int32(0)
	goto L44
L44:
	;
	v96 = int32(_a_F_network_in_5)
	v97 = int32(11)
	goto L49
L45:
	;
	if v64 == int32(0) {
		goto L36
	} else {
		goto L74
	}
L46:
	;
	v204 = v202 - int32(_a_F_network_in_5)
	if base.Ui32(int32(10)) <= base.Ui32(v204) {
		goto L29
	} else {
		goto L71
	}
L47:
	;
	v202 = int32(0)
	goto L46
L48:
	;
	v180 = v173
	v182 = v175
	goto L65
L49:
	;
	goto L56
L56:
	;
	v136 = v85 & int32(255)
	v137 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_network_in[1])))
	if base.B2i32(v136 == v137)|int32(0) == int32(0) {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v146 = v96
	v148 = v97
	goto L60
L58:
	;
	v166 = v96
	v168 = v97
	goto L59
L59:
	;
	if v168 == int32(0) {
		goto L47
	} else {
		goto L64
	}
L60:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v146)))
	v153 = v152 ^ v136*int32(16843009)
	v156 = int32(-2139062144)
	if (int32(16843008)-v153|v153)&v156 != v156 {
		v173 = v146
		v175 = v148
		goto L48
	} else {
		goto L62
	}
L61:
	;
	v166 = v161
	v168 = v163
	goto L59
L62:
	;
	v160 = int32(4)
	v161 = v146 + v160
	v163 = v148 - v160
	if base.Ui32(int32(3)) < base.Ui32(v163) {
		v146 = v161
		v148 = v163
		goto L60
	} else {
		goto L63
	}
L63:
	;
	goto L61
L64:
	;
	v173 = v166
	v175 = v168
	goto L48
L65:
	;
	v185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180))))
	if v85&int32(255) == v185 {
		goto L67
	} else {
		goto L68
	}
L66:
	;
	goto L47
L67:
	;
	v202 = v180
	goto L46
L68:
	;
	goto L69
L69:
	;
	v187 = int32(1)
	v190 = v182 - v187
	if v190 != 0 {
		v180 = v180 + v187
		v182 = v190
		goto L65
	} else {
		goto L70
	}
L70:
	;
	goto L66
L71:
	;
	v209 = v204 + v87*int32(10)
	if int32(255) < v209 {
		goto L37
	} else {
		goto L72
	}
L72:
	;
	v212 = int32(*(*int8)(unsafe.Add(mBase, uint32(v84))))
	v214 = v84 + int32(1)
	if base.Ui32((v212-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v84 = v214
		v85 = v212
		v87 = v209
		goto L44
	} else {
		goto L73
	}
L73:
	;
	goto L45
L74:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v67))) = uint8(v209)
	v224 = int32(1)
	v225 = v67 + v224
	v227 = v64 - v224
	v229 = v212 & int32(255)
	if v229 == int32(46) {
		v61 = v214
		v64 = v227
		v67 = v225
		goto L39
	} else {
		goto L75
	}
L75:
	;
	goto L40
L76:
	;
	if v229 != 0 {
		goto L37
	} else {
		goto L77
	}
L77:
	;
	v238 = v212
	v239 = v214
	v242 = v227
	v245 = v225
	goto L38
L78:
	;
	v444 = base.I32_div_s(v430, int32(8))
	if base.B2i32(v45 == v245)|base.B2i32(v245-v45 < v444) != 0 {
		goto L37
	} else {
		goto L117
	}
L79:
	;
	if v238 != int32(47) {
		goto L37
	} else {
		goto L82
	}
L80:
	;
	goto L81
L81:
	;
	if v245-v45 != int32(4) {
		goto L37
	} else {
		goto L116
	}
L82:
	;
	v251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v239))))
	if base.B2i32(base.Ui32(int32(9)) < base.Ui32((v251-int32(48))&int32(255)))|base.B2i32(base.Ui32(v245) <= base.Ui32(v45)) != 0 {
		goto L37
	} else {
		goto L83
	}
L83:
	;
	v264 = v251
	v265 = int32(0)
	v266 = v239
	goto L84
L84:
	;
	v276 = int32(_a_F_network_in_5)
	v278 = v264 & int32(255)
	v279 = int32(11)
	goto L89
L85:
	;
	if v392&int32(255) != 0 {
		goto L37
	} else {
		goto L113
	}
L86:
	;
	v386 = v384 - int32(_a_F_network_in_5)
	if base.Ui32(int32(10)) <= base.Ui32(v386) {
		goto L28
	} else {
		goto L111
	}
L87:
	;
	v384 = int32(0)
	goto L86
L88:
	;
	v362 = v355
	v364 = v357
	goto L105
L89:
	;
	goto L96
L96:
	;
	v318 = v278 & int32(255)
	v319 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_network_in[1])))
	if base.B2i32(v318 == v319)|int32(0) == int32(0) {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	v328 = v276
	v330 = v279
	goto L100
L98:
	;
	v348 = v276
	v350 = v279
	goto L99
L99:
	;
	if v350 == int32(0) {
		goto L87
	} else {
		goto L104
	}
L100:
	;
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v328)))
	v335 = v334 ^ v318*int32(16843009)
	v338 = int32(-2139062144)
	if (int32(16843008)-v335|v335)&v338 != v338 {
		v355 = v328
		v357 = v330
		goto L88
	} else {
		goto L102
	}
L101:
	;
	v348 = v343
	v350 = v345
	goto L99
L102:
	;
	v342 = int32(4)
	v343 = v328 + v342
	v345 = v330 - v342
	if base.Ui32(int32(3)) < base.Ui32(v345) {
		v328 = v343
		v330 = v345
		goto L100
	} else {
		goto L103
	}
L103:
	;
	goto L101
L104:
	;
	v355 = v348
	v357 = v350
	goto L88
L105:
	;
	v367 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v362))))
	if v278&int32(255) == v367 {
		goto L107
	} else {
		goto L108
	}
L106:
	;
	goto L87
L107:
	;
	v384 = v362
	goto L86
L108:
	;
	goto L109
L109:
	;
	v369 = int32(1)
	v372 = v364 - v369
	if v372 != 0 {
		v362 = v362 + v369
		v364 = v372
		goto L105
	} else {
		goto L110
	}
L110:
	;
	goto L106
L111:
	;
	v389 = int32(10)
	v391 = v386 + v265*v389
	v392 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v266)+1)))
	if base.Ui32((v392-int32(48))&int32(255)) < base.Ui32(v389) {
		v264 = v392
		v265 = v391
		v266 = v266 + int32(1)
		goto L84
	} else {
		goto L112
	}
L112:
	;
	goto L85
L113:
	;
	if int32(32) < v391 {
		goto L36
	} else {
		goto L114
	}
L114:
	;
	if v391 != int32(-1) {
		v430 = v391
		goto L78
	} else {
		goto L115
	}
L115:
	;
	goto L81
L116:
	;
	v430 = int32(32)
	goto L78
L117:
	;
	if v242 == int32(0) {
		v1245 = v430
		goto L22
	} else {
		goto L118
	}
L118:
	;
	base.MemoryFill(m, v245, int32(0), v242)
	v1271 = v430
	goto L15
L119:
	;
	v725 = v488
	v728 = v489
	v729 = v45
	v730 = v50
	goto L177
L120:
	;
	v492 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v488))))
	if v492|int32(32) != int32(120) {
		goto L119
	} else {
		goto L123
	}
L121:
	;
	goto L122
L122:
	;
	if base.Ui32((v489-int32(48))&int32(255)) <= base.Ui32(int32(9)) {
		goto L119
	} else {
		goto L176
	}
L123:
	;
	v497 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+2)))
	goto L124
L124:
	;
	if base.B2i32(base.Ui32(v497-int32(48)) < base.Ui32(int32(10)))|base.B2i32(base.Ui32(v497|int32(32)-int32(97)) < base.Ui32(int32(6))) == int32(0) {
		goto L119
	} else {
		goto L125
	}
L125:
	;
	v511 = int32(35)
	if v50 == int32(0) {
		v1219 = v511
		goto L24
	} else {
		goto L126
	}
L126:
	;
	if v497 == int32(0) {
		goto L127
	} else {
		goto L128
	}
L127:
	;
	v1219 = int32(44)
	goto L24
L128:
	;
	goto L129
L129:
	;
	v524 = v497
	v525 = l0 + int32(3)
	v526 = int32(0)
	v527 = v45
	v528 = v50
	v534 = int32(0)
	goto L131
L130:
	;
	v702 = base.I32_extend8_s(v695)
	if v700 == int32(0) {
		v937 = v694
		v939 = v702
		v941 = v698
		v942 = v699
		goto L25
	} else {
		goto L174
	}
L131:
	;
	v536 = v524 & int32(255)
	goto L133
L132:
	;
	v694 = v692
	v695 = int32(0)
	v698 = v686
	v699 = v687
	v700 = v688
	v701 = v689
	goto L130
L133:
	;
	if base.B2i32(base.Ui32(v536-int32(48)) < base.Ui32(int32(10)))|base.B2i32(base.Ui32(v536|int32(32)-int32(97)) < base.Ui32(int32(6))) == int32(0) {
		goto L134
	} else {
		goto L135
	}
L134:
	;
	v694 = v525
	v695 = v524
	v698 = v527
	v699 = v528
	v700 = v526
	v701 = v534
	goto L130
L135:
	;
	goto L136
L136:
	;
	v550 = int32(_a_F_network_in_6)
	if base.Ui32((v524-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L137
	} else {
		goto L138
	}
L137:
	;
	v559 = v524 | int32(32)
	goto L139
L138:
	;
	v559 = v524
	goto L139
L139:
	;
	v561 = v559 & int32(255)
	v562 = int32(17)
	goto L143
L140:
	;
	v669 = v667 - int32(_a_F_network_in_6)
	if base.Ui32(int32(16)) <= base.Ui32(v669) {
		goto L27
	} else {
		goto L165
	}
L141:
	;
	v667 = int32(0)
	goto L140
L142:
	;
	v645 = v638
	v647 = v640
	goto L159
L143:
	;
	goto L150
L150:
	;
	v601 = v561 & int32(255)
	v602 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_network_in[2])))
	if base.B2i32(v601 == v602)|int32(0) == int32(0) {
		goto L151
	} else {
		goto L152
	}
L151:
	;
	v611 = v550
	v613 = v562
	goto L154
L152:
	;
	v631 = v550
	v633 = v562
	goto L153
L153:
	;
	if v633 == int32(0) {
		goto L141
	} else {
		goto L158
	}
L154:
	;
	v617 = *(*int32)(unsafe.Add(mBase, uint32(v611)))
	v618 = v617 ^ v601*int32(16843009)
	v621 = int32(-2139062144)
	if (int32(16843008)-v618|v618)&v621 != v621 {
		v638 = v611
		v640 = v613
		goto L142
	} else {
		goto L156
	}
L155:
	;
	v631 = v626
	v633 = v628
	goto L153
L156:
	;
	v625 = int32(4)
	v626 = v611 + v625
	v628 = v613 - v625
	if base.Ui32(int32(3)) < base.Ui32(v628) {
		v611 = v626
		v613 = v628
		goto L154
	} else {
		goto L157
	}
L157:
	;
	goto L155
L158:
	;
	v638 = v631
	v640 = v633
	goto L142
L159:
	;
	v650 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v645))))
	if v561&int32(255) == v650 {
		goto L161
	} else {
		goto L162
	}
L160:
	;
	goto L141
L161:
	;
	v667 = v645
	goto L140
L162:
	;
	goto L163
L163:
	;
	v652 = int32(1)
	v655 = v647 - v652
	if v655 != 0 {
		v645 = v645 + v652
		v647 = v655
		goto L159
	} else {
		goto L164
	}
L164:
	;
	goto L160
L165:
	;
	v674 = v669 | v534<<(uint(int32(4))%32)
	v675 = int32(1)
	if v526 == v675 {
		goto L166
	} else {
		goto L167
	}
L166:
	;
	if v528 == int32(0) {
		v1219 = v511
		goto L24
	} else {
		goto L169
	}
L167:
	;
	v686 = v527
	v687 = v528
	v688 = v675
	goto L168
L168:
	;
	if v526 != 0 {
		goto L170
	} else {
		goto L171
	}
L169:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v527))) = uint8(v674)
	v681 = int32(1)
	v686 = v527 + v681
	v687 = v528 - v681
	v688 = int32(0)
	goto L168
L170:
	;
	v689 = v674
	goto L172
L171:
	;
	v689 = v669
	goto L172
L172:
	;
	v690 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v525))))
	v692 = v525 + int32(1)
	if v690 != 0 {
		v524 = v690
		v525 = v692
		v526 = v688
		v527 = v686
		v528 = v687
		v534 = v689
		goto L131
	} else {
		goto L173
	}
L173:
	;
	goto L132
L174:
	;
	if v699 == int32(0) {
		v1219 = v511
		goto L24
	} else {
		goto L175
	}
L175:
	;
	v708 = v701 << (uint(int32(4)) % 32)
	*(*uint8)(unsafe.Add(mBase, uint32(v698))) = uint8(v708)
	v710 = int32(1)
	v937 = v694
	v939 = v702
	v941 = v698 + v710
	v942 = v699 - v710
	goto L25
L176:
	;
	v1219 = int32(44)
	goto L24
L177:
	;
	v743 = v725
	v745 = v728 & int32(255)
	v746 = int32(0)
	goto L179
L178:
	;
	v1219 = v866
	goto L24
L179:
	;
	goto L185
L180:
	;
	if v730 == int32(0) {
		goto L209
	} else {
		goto L210
	}
L181:
	;
	v863 = v861 - int32(_a_F_network_in_7)
	if base.Ui32(int32(10)) <= base.Ui32(v863) {
		goto L26
	} else {
		goto L206
	}
L182:
	;
	v861 = int32(0)
	goto L181
L183:
	;
	v839 = v832
	v841 = v834
	goto L200
L184:
	;
	if base.B2i32(v778 != v779) == int32(0) {
		goto L182
	} else {
		goto L191
	}
L185:
	;
	v770 = int32(_a_F_network_in_7)
	v772 = int32(11)
	goto L186
L186:
	;
	v775 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v770))))
	if v775 == v745&int32(255) {
		v832 = v770
		v834 = v772
		goto L183
	} else {
		goto L188
	}
L187:
	;
	goto L184
L188:
	;
	v777 = int32(1)
	v778 = v772 - v777
	v779 = int32(0)
	v782 = v770 + v777
	if v782&int32(3) == v779 {
		goto L184
	} else {
		goto L189
	}
L189:
	;
	if v778 != 0 {
		v770 = v782
		v772 = v778
		goto L186
	} else {
		goto L190
	}
L190:
	;
	goto L187
L191:
	;
	v795 = v745 & int32(255)
	v796 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v782))))
	if base.B2i32(v795 == v796)|base.B2i32(base.Ui32(v778) < base.Ui32(int32(4))) == int32(0) {
		goto L192
	} else {
		goto L193
	}
L192:
	;
	v805 = v782
	v807 = v778
	goto L195
L193:
	;
	v825 = v782
	v827 = v778
	goto L194
L194:
	;
	if v827 == int32(0) {
		goto L182
	} else {
		goto L199
	}
L195:
	;
	v811 = *(*int32)(unsafe.Add(mBase, uint32(v805)))
	v812 = v811 ^ v795*int32(16843009)
	v815 = int32(-2139062144)
	if (int32(16843008)-v812|v812)&v815 != v815 {
		v832 = v805
		v834 = v807
		goto L183
	} else {
		goto L197
	}
L196:
	;
	v825 = v820
	v827 = v822
	goto L194
L197:
	;
	v819 = int32(4)
	v820 = v805 + v819
	v822 = v807 - v819
	if base.Ui32(int32(3)) < base.Ui32(v822) {
		v805 = v820
		v807 = v822
		goto L195
	} else {
		goto L198
	}
L198:
	;
	goto L196
L199:
	;
	v832 = v825
	v834 = v827
	goto L183
L200:
	;
	v844 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v839))))
	if v745&int32(255) == v844 {
		goto L202
	} else {
		goto L203
	}
L201:
	;
	goto L182
L202:
	;
	v861 = v839
	goto L181
L203:
	;
	goto L204
L204:
	;
	v846 = int32(1)
	v849 = v841 - v846
	if v849 != 0 {
		v839 = v839 + v846
		v841 = v849
		goto L200
	} else {
		goto L205
	}
L205:
	;
	goto L201
L206:
	;
	v866 = int32(44)
	v869 = v863 + v746*int32(10)
	if int32(255) < v869 {
		v1219 = v866
		goto L24
	} else {
		goto L207
	}
L207:
	;
	v873 = v743 + int32(1)
	v874 = int32(*(*int8)(unsafe.Add(mBase, uint32(v743))))
	if base.Ui32((v874-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v743 = v873
		v745 = v874
		v746 = v869
		goto L179
	} else {
		goto L208
	}
L208:
	;
	goto L180
L209:
	;
	v1219 = int32(35)
	goto L24
L210:
	;
	goto L211
L211:
	;
	v884 = int32(1)
	v885 = v730 - v884
	*(*uint8)(unsafe.Add(mBase, uint32(v729))) = uint8(v869)
	v888 = v729 + v884
	v890 = v874 & int32(255)
	if v890 != int32(46) {
		goto L212
	} else {
		goto L213
	}
L212:
	;
	if v890 == int32(0) {
		v937 = v873
		v939 = v874
		v941 = v888
		v942 = v885
		goto L25
	} else {
		goto L215
	}
L213:
	;
	goto L214
L214:
	;
	v899 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v743)+1)))
	if base.Ui32((v899-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v725 = v743 + int32(2)
		v728 = v899
		v729 = v888
		v730 = v885
		goto L177
	} else {
		goto L217
	}
L215:
	;
	if v890 != int32(47) {
		v1219 = v866
		goto L24
	} else {
		goto L216
	}
L216:
	;
	v937 = v873
	v939 = v874
	v941 = v888
	v942 = v885
	goto L25
L217:
	;
	goto L178
L218:
	;
	v909 = F_inet_cidr_pton_ipv6(m, l0, v45, int32(16))
	mBase = m.M
	v1271 = v909
	goto L15
L219:
	;
	goto L220
L220:
	;
	v910 = F_inet_cidr_pton_ipv6(m, l0, v45, v50)
	mBase = m.M
	v1271 = v910
	goto L15
L221:
	;
	if v1168 <= v1167 {
		v1245 = v1168
		goto L22
	} else {
		goto L276
	}
L222:
	;
	v1136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45))))
	if base.Ui32(int32(239)) < base.Ui32(v1136) {
		v1151 = int32(32)
		goto L262
	} else {
		goto L263
	}
L223:
	;
	m.Env.X__assert_fail(m, int32(_a_F_network_in_0), int32(_a_F_network_in_1), int32(180), int32(_a_F_network_in_2))
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L224:
	;
	if v1092 == int32(-1) {
		goto L222
	} else {
		goto L261
	}
L225:
	;
	v951 = int32(44)
	v953 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v937))))
	if base.B2i32(base.Ui32(v941) <= base.Ui32(v45))|base.B2i32(base.Ui32(int32(9)) < base.Ui32((v953-int32(48))&int32(255))) != 0 {
		v1219 = v951
		goto L24
	} else {
		goto L228
	}
L226:
	;
	goto L227
L227:
	;
	if v939|base.B2i32(v941 == v45) != 0 {
		v1219 = int32(44)
		goto L24
	} else {
		goto L260
	}
L228:
	;
	v965 = v937
	v966 = int32(0)
	v968 = v953
	goto L229
L229:
	;
	v979 = v968 & int32(255)
	goto L235
L230:
	;
	if v1093&int32(255) != 0 {
		v1219 = v951
		goto L24
	} else {
		goto L258
	}
L231:
	;
	v1087 = v1085 - int32(_a_F_network_in_7)
	if base.Ui32(int32(10)) <= base.Ui32(v1087) {
		goto L223
	} else {
		goto L256
	}
L232:
	;
	v1085 = int32(0)
	goto L231
L233:
	;
	v1063 = v1056
	v1065 = v1058
	goto L250
L234:
	;
	if base.B2i32(v1002 != v1003) == int32(0) {
		goto L232
	} else {
		goto L241
	}
L235:
	;
	v994 = int32(_a_F_network_in_7)
	v996 = int32(11)
	goto L236
L236:
	;
	v999 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v994))))
	if v999 == v979&int32(255) {
		v1056 = v994
		v1058 = v996
		goto L233
	} else {
		goto L238
	}
L237:
	;
	goto L234
L238:
	;
	v1001 = int32(1)
	v1002 = v996 - v1001
	v1003 = int32(0)
	v1006 = v994 + v1001
	if v1006&int32(3) == v1003 {
		goto L234
	} else {
		goto L239
	}
L239:
	;
	if v1002 != 0 {
		v994 = v1006
		v996 = v1002
		goto L236
	} else {
		goto L240
	}
L240:
	;
	goto L237
L241:
	;
	v1019 = v979 & int32(255)
	v1020 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1006))))
	if base.B2i32(v1019 == v1020)|base.B2i32(base.Ui32(v1002) < base.Ui32(int32(4))) == int32(0) {
		goto L242
	} else {
		goto L243
	}
L242:
	;
	v1029 = v1006
	v1031 = v1002
	goto L245
L243:
	;
	v1049 = v1006
	v1051 = v1002
	goto L244
L244:
	;
	if v1051 == int32(0) {
		goto L232
	} else {
		goto L249
	}
L245:
	;
	v1035 = *(*int32)(unsafe.Add(mBase, uint32(v1029)))
	v1036 = v1035 ^ v1019*int32(16843009)
	v1039 = int32(-2139062144)
	if (int32(16843008)-v1036|v1036)&v1039 != v1039 {
		v1056 = v1029
		v1058 = v1031
		goto L233
	} else {
		goto L247
	}
L246:
	;
	v1049 = v1044
	v1051 = v1046
	goto L244
L247:
	;
	v1043 = int32(4)
	v1044 = v1029 + v1043
	v1046 = v1031 - v1043
	if base.Ui32(int32(3)) < base.Ui32(v1046) {
		v1029 = v1044
		v1031 = v1046
		goto L245
	} else {
		goto L248
	}
L248:
	;
	goto L246
L249:
	;
	v1056 = v1049
	v1058 = v1051
	goto L233
L250:
	;
	v1068 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1063))))
	if v979&int32(255) == v1068 {
		goto L252
	} else {
		goto L253
	}
L251:
	;
	goto L232
L252:
	;
	v1085 = v1063
	goto L231
L253:
	;
	goto L254
L254:
	;
	v1070 = int32(1)
	v1073 = v1065 - v1070
	if v1073 != 0 {
		v1063 = v1063 + v1070
		v1065 = v1073
		goto L250
	} else {
		goto L255
	}
L255:
	;
	goto L251
L256:
	;
	v1090 = int32(10)
	v1092 = v1087 + v966*v1090
	v1093 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v965)+1)))
	if base.Ui32((v1093-int32(48))&int32(255)) < base.Ui32(v1090) {
		v965 = v965 + int32(1)
		v966 = v1092
		v968 = v1093
		goto L229
	} else {
		goto L257
	}
L257:
	;
	goto L230
L258:
	;
	if v1092 <= int32(32) {
		goto L224
	} else {
		goto L259
	}
L259:
	;
	v1219 = int32(35)
	goto L24
L260:
	;
	goto L222
L261:
	;
	v1167 = (v941 - v45) << (uint(int32(3)) % 32)
	v1168 = v1092
	goto L221
L262:
	;
	v1154 = (v941 - v45) << (uint(int32(3)) % 32)
	if v1154 < v1151 {
		goto L269
	} else {
		goto L270
	}
L263:
	;
	if base.Ui32(int32(223)) < base.Ui32(v1136) {
		v1151 = int32(8)
		goto L262
	} else {
		goto L264
	}
L264:
	;
	if base.Ui32(int32(191)) < base.Ui32(v1136) {
		v1151 = int32(24)
		goto L262
	} else {
		goto L265
	}
L265:
	;
	if base.I32_extend8_s(v1136) < int32(0) {
		goto L266
	} else {
		goto L267
	}
L266:
	;
	v1150 = int32(16)
	goto L268
L267:
	;
	v1150 = int32(8)
	goto L268
L268:
	;
	v1151 = v1150
	goto L262
L269:
	;
	v1156 = v1151
	goto L271
L270:
	;
	v1156 = v1154
	goto L271
L271:
	;
	if v1156 != int32(8) {
		v1167 = v1154
		v1168 = v1156
		goto L221
	} else {
		goto L272
	}
L272:
	;
	if v1136 == int32(224) {
		goto L273
	} else {
		goto L274
	}
L273:
	;
	v1163 = int32(4)
	goto L275
L274:
	;
	v1163 = int32(8)
	goto L275
L275:
	;
	v1167 = v1154
	v1168 = v1163
	goto L221
L276:
	;
	v1187 = v941
	v1188 = v942
	goto L277
L277:
	;
	if v1188 == int32(0) {
		goto L279
	} else {
		goto L280
	}
L278:
	;
	v1245 = v1168
	goto L22
L279:
	;
	v1219 = int32(35)
	goto L24
L280:
	;
	goto L281
L281:
	;
	v1198 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1187))) = uint8(v1198)
	v1200 = int32(1)
	v1203 = v1187 + v1200
	if (v1203-v45)<<(uint(int32(3))%32) < v1168 {
		v1187 = v1203
		v1188 = v1188 - v1200
		goto L277
	} else {
		goto L282
	}
L282:
	;
	goto L278
L283:
	;
	v1276 = int32(1)
	v1278 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
	if v1278&v1276 != 0 {
		goto L286
	} else {
		goto L287
	}
L284:
	;
	goto L285
L285:
	;
	v1291 = F_errsave_start(m, l2)
	mBase = m.M
	v1292 = m.ExcPending
	if v1292 != 0 {
		goto L1
	} else {
		goto L293
	}
L286:
	;
	v1281 = v1276
	goto L288
L287:
	;
	v1281 = int32(4)
	goto L288
L288:
	;
	v1282 = v21 + v1281
	v1283 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1282))))
	if v1283 == int32(2) {
		goto L289
	} else {
		goto L290
	}
L289:
	;
	v1286 = int32(32)
	goto L291
L290:
	;
	v1286 = int32(128)
	goto L291
L291:
	;
	if base.Ui32(v1271) <= base.Ui32(v1286) {
		goto L14
	} else {
		goto L292
	}
L292:
	;
	goto L285
L293:
	;
	if v1291 == int32(0) {
		goto L294
	} else {
		goto L295
	}
L294:
	;
	v1429 = int32(0)
	goto L13
L295:
	;
	goto L296
L296:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v1298 = m.ExcPending
	if v1298 != 0 {
		goto L1
	} else {
		goto L297
	}
L297:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+4)) = l0
	if l1 != 0 {
		goto L298
	} else {
		goto L299
	}
L298:
	;
	v1302 = int32(_a_F_network_in_8)
	goto L300
L299:
	;
	v1302 = int32(_a_F_network_in_9)
	goto L300
L300:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v1302
	F_errmsg(m, int32(_a_F_network_in_10), v18)
	mBase = m.M
	v1306 = m.ExcPending
	if v1306 != 0 {
		goto L1
	} else {
		goto L301
	}
L301:
	;
	F_errsave_finish(m, l2, int32(_a_F_network_in_11), int32(98), int32(_a_F_network_in_12))
	mBase = m.M
	v1311 = m.ExcPending
	if v1311 != 0 {
		goto L1
	} else {
		goto L302
	}
L302:
	;
	v1429 = int32(0)
	goto L13
L303:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1282)+1)) = uint8(v1271)
	if v1283 == int32(2) {
		goto L322
	} else {
		goto L323
	}
L304:
	;
	v1318 = int32(base.Ui32(v1271) >> (uint(int32(3)) % 32))
	if v1283 == int32(2) {
		goto L305
	} else {
		goto L306
	}
L305:
	;
	v1323 = int32(4)
	goto L307
L306:
	;
	v1323 = int32(16)
	goto L307
L307:
	;
	if base.Ui32(v1323) <= base.Ui32(v1318) {
		goto L303
	} else {
		goto L308
	}
L308:
	;
	v1326 = v1282 + int32(2)
	v1328 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1326+v1318))))
	if v1328<<(uint(v1271&int32(7))%32)&int32(255) != 0 {
		goto L309
	} else {
		goto L310
	}
L309:
	;
	v1372 = int32(0)
	v1373 = F_errsave_start(m, l2)
	mBase = m.M
	v1374 = m.ExcPending
	if v1374 != 0 {
		goto L1
	} else {
		goto L316
	}
L310:
	;
	v1335 = v1318 + int32(1)
	if v1335 == v1323 {
		goto L303
	} else {
		goto L311
	}
L311:
	;
	v1338 = v1335
	goto L312
L312:
	;
	v1353 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1338+v1326))))
	if v1353 != 0 {
		goto L309
	} else {
		goto L314
	}
L313:
	;
	goto L303
L314:
	;
	v1355 = v1338 + int32(1)
	if v1323 != v1355 {
		v1338 = v1355
		goto L312
	} else {
		goto L315
	}
L315:
	;
	goto L313
L316:
	;
	if v1373 == int32(0) {
		v1429 = v1372
		goto L13
	} else {
		goto L317
	}
L317:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v1379 = m.ExcPending
	if v1379 != 0 {
		goto L1
	} else {
		goto L318
	}
L318:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = l0
	F_errmsg(m, int32(_a_F_network_in_13), v18+int32(16))
	mBase = m.M
	v1385 = m.ExcPending
	if v1385 != 0 {
		goto L1
	} else {
		goto L319
	}
L319:
	;
	v1388 = F_errdetail(m, int32(_a_F_network_in_14), int32(0))
	mBase = m.M
	v1389 = m.ExcPending
	if v1389 != 0 {
		goto L1
	} else {
		goto L320
	}
L320:
	;
	F_errsave_finish(m, l2, int32(_a_F_network_in_11), int32(109), int32(_a_F_network_in_12))
	mBase = m.M
	v1394 = m.ExcPending
	if v1394 != 0 {
		goto L1
	} else {
		goto L321
	}
L321:
	;
	v1429 = v1372
	goto L13
L322:
	;
	v1415 = int32(40)
	goto L324
L323:
	;
	v1415 = int32(88)
	goto L324
L324:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = v1415
	v1429 = v21
	goto L13
}
func F_network_masklen(m *base.Module, l0 int32) int64 {
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
	var v14 int64
	_ = v14
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = F_pg_detoast_datum_packed(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		v7 = int32(1)
		v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3))))
		if v9&v7 != 0 {
			v12 = v7
		} else {
			v12 = int32(4)
		}
		v14 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v3+v12)+1)))
		return v14
	}
}
func F_network_netmask(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v95 int32
	_ = v95
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v7 = F_pg_detoast_datum_packed(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int64(0)
	} else {
		v12 = F_palloc0(m, int32(22))
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int64(0)
		} else {
			v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
			v15 = int32(1)
			v16 = v14 & v15
			v18 = int32(4)
			v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
			v23 = v21 & v15
			if v23 != 0 {
				v24 = v15
			} else {
				v24 = v18
			}
			v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7+v24)+1)))
			if v26 != 0 {
				if v16 != 0 {
					v29 = int32(1)
				} else {
					v29 = int32(4)
				}
				v34 = v26
				v35 = int32(0)
				for {
					if base.Ui32(int32(7)) < base.Ui32(v34) {
						v47 = int32(-1)
					} else {
						v47 = int32(255) << (uint(int32(8)-v34) % 32)
					}
					*(*uint8)(unsafe.Add(mBase, uint32(v35+(v12+v29+int32(2))))) = uint8(v47)
					v51 = int32(8)
					if v34 <= v51 {
						v54 = v51
					} else {
						v54 = v34
					}
					v56 = v54 - int32(8)
					if v56 != 0 {
						v34 = v56
						v35 = v35 + int32(1)
						continue
					} else {
						break
					}
					break
				}
				v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
				v58 = int32(1)
				v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
				v67 = v57 & v58
				v68 = v60 & v58
			} else {
				v67 = v23
				v68 = v16
			}
			if v68 != 0 {
				v69 = v15
			} else {
				v69 = v18
			}
			v70 = v12 + v69
			if v67 != 0 {
				v73 = int32(1)
			} else {
				v73 = int32(4)
			}
			v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7+v73))))
			*(*uint8)(unsafe.Add(mBase, uint32(v70))) = uint8(v75)
			v79 = int32(1)
			v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
			if v81&v79 != 0 {
				v84 = v79
			} else {
				v84 = int32(4)
			}
			v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7+v84))))
			if v86 == int32(2) {
				v89 = int32(32)
			} else {
				v89 = int32(-128)
			}
			*(*uint8)(unsafe.Add(mBase, uint32(v70)+1)) = uint8(v89)
			if v75 == int32(2) {
				v95 = int32(40)
			} else {
				v95 = int32(88)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v12))) = v95
			return base.I64_extend_i32_u(v12)
		}
	}
}
func F_network_overlap(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int64
	_ = v5
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
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
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
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v157 int64
	_ = v157
	v5 = int64(0)
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v7 = F_pg_detoast_datum_packed(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v12 = F_pg_detoast_datum_packed(m, v11)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v14 = int32(1)
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
	if v16&v14 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	return int64(1)
L5:
	;
	return v157
L6:
	;
	v19 = v14
	goto L8
L7:
	;
	v19 = int32(4)
	goto L8
L8:
	;
	v20 = v7 + v19
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
	v22 = int32(1)
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
	if v24&v22 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v27 = v22
	goto L11
L10:
	;
	v27 = int32(4)
	goto L11
L11:
	;
	v28 = v12 + v27
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28))))
	if v21 != v29 {
		v157 = v5
		goto L5
	} else {
		goto L12
	}
L12:
	;
	v31 = int32(2)
	v32 = v20 + v31
	v34 = v28 + v31
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+1)))
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+1)))
	if base.Ui32(v35) < base.Ui32(v36) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v38 = v20
	goto L15
L14:
	;
	v38 = v28
	goto L15
L15:
	;
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+1)))
	v41 = int32(base.Ui32(v39) >> (uint(int32(3)) % 32))
	if base.Ui32(int32(4)) <= base.Ui32(v41) {
		goto L19
	} else {
		goto L20
	}
L16:
	;
	if v103 != 0 {
		v157 = v5
		goto L5
	} else {
		goto L34
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
	if (v32|v34)&int32(3) != 0 {
		v72 = v32
		v73 = v34
		v74 = v41
		goto L18
	} else {
		goto L22
	}
L20:
	;
	v65 = v32
	v66 = v34
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
	v49 = v32
	v50 = v34
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
	v105 = v39 & int32(7)
	if v105 == int32(0) {
		goto L4
	} else {
		goto L35
	}
L35:
	;
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41+v32))))
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41+v34))))
	v112 = v109 ^ v111
	if base.Ui32(int32(127)) < base.Ui32(v112) {
		v157 = v5
		goto L5
	} else {
		goto L36
	}
L36:
	;
	if v105 == int32(1) {
		goto L4
	} else {
		goto L37
	}
L37:
	;
	if v112<<(uint(int32(1))%32)&int32(128) != 0 {
		v157 = v5
		goto L5
	} else {
		goto L38
	}
L38:
	;
	if base.Ui32(v105) < base.Ui32(int32(3)) {
		goto L4
	} else {
		goto L39
	}
L39:
	;
	if v112<<(uint(int32(2))%32)&int32(128) != 0 {
		v157 = v5
		goto L5
	} else {
		goto L40
	}
L40:
	;
	if v105 == int32(3) {
		goto L4
	} else {
		goto L41
	}
L41:
	;
	if v112<<(uint(int32(3))%32)&int32(128) != 0 {
		v157 = v5
		goto L5
	} else {
		goto L42
	}
L42:
	;
	if base.Ui32(v105) < base.Ui32(int32(5)) {
		goto L4
	} else {
		goto L43
	}
L43:
	;
	if v112<<(uint(int32(4))%32)&int32(128) != 0 {
		v157 = v5
		goto L5
	} else {
		goto L44
	}
L44:
	;
	if v105 == int32(5) {
		goto L4
	} else {
		goto L45
	}
L45:
	;
	if v112<<(uint(int32(5))%32)&int32(128) != 0 {
		v157 = v5
		goto L5
	} else {
		goto L46
	}
L46:
	;
	if v105 != int32(7) {
		v157 = int64(1)
		goto L5
	} else {
		goto L47
	}
L47:
	;
	v157 = base.I64_extend_i32_u(base.B2i32(v112&int32(2) == int32(0)))
	goto L5
}
func F_network_show(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int64
	_ = v6
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
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v90 int64
	_ = v90
	v6 = int64(0)
	v7 = m.G0
	v9 = v7 - int32(80)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v12 = F_pg_detoast_datum_packed(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int64(0)
	} else {
		v16 = int32(1)
		v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
		if v18&v16 != 0 {
			v21 = v16
		} else {
			v21 = int32(4)
		}
		v22 = v12 + v21
		v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22))))
		v24 = int32(2)
		if v23 == v24 {
			v30 = int32(32)
		} else {
			v30 = int32(128)
		}
		v33 = F_pg_inet_net_ntop(m, v23, v22+v24, v30, v9+int32(16))
		mBase = m.M
		v34 = m.ExcPending
		if v34 != 0 {
			return int64(0)
		} else {
			if v33 == int32(0) {
				v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v38 = F_errsave_start(m, v37)
				mBase = m.M
				v39 = m.ExcPending
				if v39 != 0 {
					return int64(0)
				} else {
					if v38 == int32(0) {
						v90 = v6
						m.G0 = v9 + int32(80)
						return v90
					} else {
						F_errcode(m, int32(50462850))
						mBase = m.M
						v44 = m.ExcPending
						if v44 != 0 {
							return int64(0)
						} else {
							F_errmsg(m, int32(_a_F_network_show_0), int32(0))
							mBase = m.M
							v48 = m.ExcPending
							if v48 != 0 {
								return int64(0)
							} else {
								F_errsave_finish(m, v37, int32(_a_F_network_show_1), int32(1142), int32(_a_F_network_show_2))
								mBase = m.M
								v53 = m.ExcPending
								if v53 != 0 {
									return int64(0)
								} else {
									v90 = v6
									m.G0 = v9 + int32(80)
									return v90
								}
							}
						}
					}
				}
			} else {
				v55 = v9 + int32(16)
				v56 = int32(47)
				v57 = F___strchrnul(m, v55, v56)
				mBase = m.M
				v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57))))
				if v59 == v56 {
					v63 = v57
				} else {
					v63 = int32(0)
				}
				if v63 == int32(0) {
					v66 = F_strlen(m, v55)
					mBase = m.M
					v67 = int32(1)
					v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
					if v69&v67 != 0 {
						v72 = v67
					} else {
						v72 = int32(4)
					}
					v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12+v72)+1)))
					*(*int32)(unsafe.Add(mBase, uint32(v9))) = v74
					v80 = F_pg_snprintf(m, v66+v55, int32(50)-v66, int32(_a_F_network_show_3), v9)
					mBase = m.M
					v81 = m.ExcPending
					if v81 != 0 {
						return int64(0)
					} else {
						v85 = F_cstring_to_text(m, v9+int32(16))
						mBase = m.M
						v86 = m.ExcPending
						if v86 != 0 {
							return int64(0)
						} else {
							v90 = base.I64_extend_i32_u(v85)
							m.G0 = v9 + int32(80)
							return v90
						}
					}
				} else {
					v85 = F_cstring_to_text(m, v9+int32(16))
					mBase = m.M
					v86 = m.ExcPending
					if v86 != 0 {
						return int64(0)
					} else {
						v90 = base.I64_extend_i32_u(v85)
						m.G0 = v9 + int32(80)
						return v90
					}
				}
			}
		}
	}
}
func F_network_sortsupport(m *base.Module, l0 int32) int64 {
	var v5 int64
	_ = v5
	var v8 int32
	_ = v8
	v5 = Fn14333(m, l0, int32(1592), int32(1591), int32(1590))
	v8 = m.ExcPending
	if v8 != 0 {
		return int64(0)
	} else {
		return v5
	}
}
