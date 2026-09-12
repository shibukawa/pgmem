package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_pgstat_count_backend_io_op(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int64) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v24 int32
	_ = v24
	var v27 int64
	_ = v27
	var v33 int64
	_ = v33
	var v37 int32
	_ = v37
	v8 = *(*int32)(unsafe.Add(mBase, _consts[200]))
	if base.Ui32(int32(16)) < base.Ui32(v8) {
	} else {
		if int32(1)<<(uint(v8)%32)&int32(115186) == int32(0) {
		} else {
			v24 = l0*int32(320) + l1<<(uint(int32(6))%32) + l2<<(uint(int32(3))%32)
			v27 = *(*int64)(unsafe.Add(mBase, uint32(v24)+uint32(_consts[927])))
			*(*int64)(unsafe.Add(mBase, uint32(v24)+uint32(_consts[927]))) = v27 + base.I64_extend_i32_u(l3)
			v33 = *(*int64)(unsafe.Add(mBase, uint32(v24)+uint32(_consts[928])))
			*(*int64)(unsafe.Add(mBase, uint32(v24)+uint32(_consts[928]))) = v33 + l4
			v37 = int32(1)
			*(*uint8)(unsafe.Add(mBase, _consts[154])) = uint8(v37)
			*(*uint8)(unsafe.Add(mBase, _consts[202])) = uint8(v37)
		}
	}
	return
}
func F_pgstat_drop_transactional(m *base.Module, l0 int32, l1 int32, l2 int64) {
	var v6 int32
	_ = v6
	F_create_drop_transactional_internal(m, l0, l1, l2, int32(0))
	v6 = m.ExcPending
	if v6 != 0 {
		return
	} else {
		return
	}
}
func F_pgstat_fetch_pending_entry(m *base.Module, l0 int32, l1 int32, l2 int64) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
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
	v4 = int32(0)
	v7 = F_pgstat_get_entry_ref(m, l0, l1, l2, v4, v4)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		if v7 != 0 {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
			if v12 != 0 {
				v13 = v7
			} else {
				v13 = int32(0)
			}
			v14 = v13
		} else {
			v14 = v4
		}
		return v14
	}
}
func F_pgstat_fetch_stat_checkpointer(m *base.Module) int32 {
	var v5 int32
	_ = v5
	F_pgstat_snapshot_fixed(m, int32(9))
	v5 = m.ExcPending
	if v5 != 0 {
		return int32(0)
	} else {
		return int32(4477992)
	}
}
func F_pgstat_get_beentry_by_proc_number(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	v3 = m.G0
	v5 = v3 - int32(432)
	m.G0 = v5
	F_pgstat_read_current_status(m)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v5)+408)) = l0
		v13 = *(*int32)(unsafe.Add(mBase, _consts[916]))
		v15 = *(*int32)(unsafe.Add(mBase, _consts[917]))
		v18 = F_bsearch(m, v5, v13, v15, int32(432), int32(1212))
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int32(0)
		} else {
			m.G0 = v5 + int32(432)
			return v18
		}
	}
}
func F_pgstat_get_entry_ref(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v81 int32
	_ = v81
	var v82 int64
	_ = v82
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v97 int64
	_ = v97
	var v100 int64
	_ = v100
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v106 int64
	_ = v106
	var v110 int32
	_ = v110
	var v111 int64
	_ = v111
	var v114 int32
	_ = v114
	var v125 int32
	_ = v125
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v156 int32
	_ = v156
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v188 int32
	_ = v188
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v236 int64
	_ = v236
	var v238 int64
	_ = v238
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v260 int32
	_ = v260
	var v275 int64
	_ = v275
	var v276 int64
	_ = v276
	var v279 int64
	_ = v279
	var v282 int64
	_ = v282
	var v283 int64
	_ = v283
	var v284 int64
	_ = v284
	var v289 int64
	_ = v289
	var v295 int64
	_ = v295
	var v301 int64
	_ = v301
	var v306 int64
	_ = v306
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v322 int32
	_ = v322
	var v337 int64
	_ = v337
	var v340 int32
	_ = v340
	var v342 int64
	_ = v342
	var v344 int64
	_ = v344
	var v347 int64
	_ = v347
	var v348 int64
	_ = v348
	var v358 int64
	_ = v358
	var v363 int32
	_ = v363
	var v364 int64
	_ = v364
	var v365 int32
	_ = v365
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v373 int64
	_ = v373
	var v383 int64
	_ = v383
	var v391 int32
	_ = v391
	var v398 float64
	_ = v398
	var v404 int32
	_ = v404
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v417 int32
	_ = v417
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v437 int64
	_ = v437
	var v438 int64
	_ = v438
	var v441 int64
	_ = v441
	var v442 int64
	_ = v442
	var v443 int64
	_ = v443
	var v448 int64
	_ = v448
	var v450 int64
	_ = v450
	var v455 int64
	_ = v455
	var v461 int64
	_ = v461
	var v466 int64
	_ = v466
	var v477 int32
	_ = v477
	var v482 int32
	_ = v482
	var v489 int32
	_ = v489
	var v494 int32
	_ = v494
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v509 int64
	_ = v509
	var v510 int64
	_ = v510
	var v513 int64
	_ = v513
	var v514 int64
	_ = v514
	var v515 int64
	_ = v515
	var v520 int64
	_ = v520
	var v523 int32
	_ = v523
	var v524 int64
	_ = v524
	var v529 int64
	_ = v529
	var v535 int64
	_ = v535
	var v540 int64
	_ = v540
	var v548 int32
	_ = v548
	var v550 int32
	_ = v550
	var v569 int32
	_ = v569
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v576 int64
	_ = v576
	var v578 int64
	_ = v578
	var v580 int64
	_ = v580
	var v603 int32
	_ = v603
	var v607 int32
	_ = v607
	var v609 int32
	_ = v609
	var v633 int32
	_ = v633
	var v637 int32
	_ = v637
	var v641 int32
	_ = v641
	var v646 int32
	_ = v646
	var v650 int32
	_ = v650
	var v654 int32
	_ = v654
	var v659 int32
	_ = v659
	var v680 int32
	_ = v680
	var v681 int32
	_ = v681
	var v682 int32
	_ = v682
	var v683 int32
	_ = v683
	var v686 int32
	_ = v686
	var v687 int32
	_ = v687
	var v691 int32
	_ = v691
	var v697 int32
	_ = v697
	var v700 int32
	_ = v700
	var v711 int32
	_ = v711
	var v712 int32
	_ = v712
	var v720 int32
	_ = v720
	var v721 int32
	_ = v721
	var v722 int32
	_ = v722
	var v725 int32
	_ = v725
	var v726 int32
	_ = v726
	var v728 int32
	_ = v728
	var v729 int32
	_ = v729
	var v731 int32
	_ = v731
	var v733 int32
	_ = v733
	var v743 int32
	_ = v743
	var v744 int32
	_ = v744
	var v745 int32
	_ = v745
	var v748 int32
	_ = v748
	var v749 int32
	_ = v749
	var v750 int32
	_ = v750
	var v753 int32
	_ = v753
	var v754 int32
	_ = v754
	var v756 int32
	_ = v756
	var v761 int32
	_ = v761
	var v774 int32
	_ = v774
	var v776 int32
	_ = v776
	var v777 int64
	_ = v777
	var v778 int64
	_ = v778
	var v781 int64
	_ = v781
	var v782 int64
	_ = v782
	var v783 int64
	_ = v783
	var v788 int64
	_ = v788
	var v790 int64
	_ = v790
	var v795 int64
	_ = v795
	var v801 int64
	_ = v801
	var v806 int64
	_ = v806
	var v814 int32
	_ = v814
	var v816 int32
	_ = v816
	var v818 int32
	_ = v818
	var v822 int32
	_ = v822
	var v825 int32
	_ = v825
	var v826 int32
	_ = v826
	var v837 int32
	_ = v837
	var v838 int32
	_ = v838
	var v848 int32
	_ = v848
	var v851 int32
	_ = v851
	var v853 int64
	_ = v853
	var v861 int32
	_ = v861
	var v864 int32
	_ = v864
	var v865 int32
	_ = v865
	var v872 int32
	_ = v872
	var v876 int32
	_ = v876
	var v893 int32
	_ = v893
	var v897 int32
	_ = v897
	var v907 int32
	_ = v907
	var v910 int32
	_ = v910
	var v913 int32
	_ = v913
	var v914 int64
	_ = v914
	var v916 int64
	_ = v916
	var v918 int64
	_ = v918
	var v941 int32
	_ = v941
	var v942 int32
	_ = v942
	var v950 int32
	_ = v950
	var v953 int32
	_ = v953
	var v955 int64
	_ = v955
	var v961 int32
	_ = v961
	var v964 int32
	_ = v964
	var v965 int32
	_ = v965
	var v966 int32
	_ = v966
	var v969 int32
	_ = v969
	var v973 int32
	_ = v973
	var v977 int32
	_ = v977
	var v982 int32
	_ = v982
	var v990 int32
	_ = v990
	var v1003 int32
	_ = v1003
	var v1004 int32
	_ = v1004
	var v1018 int32
	_ = v1018
	var v1032 int32
	_ = v1032
	var v1034 int32
	_ = v1034
	var v1035 int32
	_ = v1035
	var v1041 int32
	_ = v1041
	var v1047 int32
	_ = v1047
	var v1063 int32
	_ = v1063
	var v1067 int32
	_ = v1067
	var v1068 int32
	_ = v1068
	var v1072 int32
	_ = v1072
	var v1077 int32
	_ = v1077
	var v1078 int32
	_ = v1078
	var v1079 int32
	_ = v1079
	var v1080 int32
	_ = v1080
	var v1082 int32
	_ = v1082
	var v1087 int32
	_ = v1087
	var v1102 int32
	_ = v1102
	var v1104 int32
	_ = v1104
	var v1112 int32
	_ = v1112
	var v1114 int32
	_ = v1114
	var v1116 int32
	_ = v1116
	var v1117 int32
	_ = v1117
	var v1119 int32
	_ = v1119
	var v1120 int32
	_ = v1120
	var v1124 int32
	_ = v1124
	var v1125 int32
	_ = v1125
	var v1126 int32
	_ = v1126
	var v1131 int32
	_ = v1131
	var v1132 int32
	_ = v1132
	var v1138 int32
	_ = v1138
	var v1143 int32
	_ = v1143
	var v1145 int32
	_ = v1145
	var v1148 int32
	_ = v1148
	var v1152 int32
	_ = v1152
	var v1154 int32
	_ = v1154
	var v1157 int64
	_ = v1157
	var v1159 int64
	_ = v1159
	var v1163 int32
	_ = v1163
	var v1165 int32
	_ = v1165
	var v1173 int32
	_ = v1173
	var v1174 int32
	_ = v1174
	var v1175 int32
	_ = v1175
	var v1176 int32
	_ = v1176
	var v1177 int32
	_ = v1177
	var v1179 int32
	_ = v1179
	var v1180 int32
	_ = v1180
	var v1183 int32
	_ = v1183
	var v1201 int32
	_ = v1201
	var v1203 int32
	_ = v1203
	var v1211 int32
	_ = v1211
	var v1213 int32
	_ = v1213
	var v1215 int32
	_ = v1215
	var v1216 int32
	_ = v1216
	var v1233 int32
	_ = v1233
	var v1235 int32
	_ = v1235
	var v1243 int32
	_ = v1243
	var v1245 int32
	_ = v1245
	var v1247 int32
	_ = v1247
	var v1248 int32
	_ = v1248
	var v1250 int32
	_ = v1250
	var v1251 int32
	_ = v1251
	var v1256 int32
	_ = v1256
	var v1258 int32
	_ = v1258
	var v1261 int32
	_ = v1261
	var v1265 int32
	_ = v1265
	var v1270 int32
	_ = v1270
	var v1272 int32
	_ = v1272
	var v1273 int64
	_ = v1273
	var v1275 int64
	_ = v1275
	var v1281 int32
	_ = v1281
	var v1284 int32
	_ = v1284
	var v1285 int32
	_ = v1285
	var v1286 int32
	_ = v1286
	var v1287 int32
	_ = v1287
	var v1288 int32
	_ = v1288
	var v1293 int32
	_ = v1293
	var v1295 int32
	_ = v1295
	var v1298 int32
	_ = v1298
	var v1305 int32
	_ = v1305
	var v1325 int32
	_ = v1325
	var v1327 int32
	_ = v1327
	var v1331 int32
	_ = v1331
	var v1334 int32
	_ = v1334
	var v1338 int32
	_ = v1338
	var v1339 int64
	_ = v1339
	var v1341 int64
	_ = v1341
	var v1347 int32
	_ = v1347
	var v1352 int32
	_ = v1352
	var v1359 int32
	_ = v1359
	var v1373 int32
	_ = v1373
	v6 = int32(0)
	v21 = m.G0
	v23 = v21 - int32(96)
	m.G0 = v23
	*(*int64)(unsafe.Add(mBase, uint32(v23)+72)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v23)+68)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v23)+64)) = l0
	v29 = *(*int32)(unsafe.Add(mBase, _consts[931]))
	if v29 == v6 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v34 = *(*int32)(unsafe.Add(mBase, _consts[12]))
	v39 = F_AllocSetContextCreateInternal(m, v34, int32(355453), int32(0), int32(1024), int32(8192))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v45 = *(*int32)(unsafe.Add(mBase, _consts[932]))
	if v45 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	return int32(0)
L5:
	;
	*(*int32)(unsafe.Add(mBase, _consts[931])) = v39
	goto L3
L6:
	;
	v50 = *(*int32)(unsafe.Add(mBase, _consts[12]))
	v55 = F_AllocSetContextCreateInternal(m, v50, int32(338443), int32(0), int32(1024), int32(8192))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L4
	} else {
		goto L9
	}
L7:
	;
	v58 = v45
	goto L8
L8:
	;
	v60 = *(*int32)(unsafe.Add(mBase, _consts[933]))
	if v60 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, _consts[932])) = v55
	v58 = v55
	goto L8
L10:
	;
	v64 = F_MemoryContextAllocZero(m, v58, int32(32))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L4
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	if l4 != 0 {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v64)+28)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v64)+24)) = v58
	v71 = F_MemoryContextAllocExtended(m, v58, int32(6144), int32(5))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L4
	} else {
		goto L14
	}
L14:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v64)+12)) = int64(987842478335)
	*(*int64)(unsafe.Add(mBase, uint32(v64))) = int64(256)
	*(*int32)(unsafe.Add(mBase, uint32(v64)+20)) = v71
	*(*int32)(unsafe.Add(mBase, _consts[933])) = v64
	v81 = *(*int32)(unsafe.Add(mBase, _consts[278]))
	v82 = *(*int64)(unsafe.Add(mBase, uint32(v81)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v81)+16)) = v82
	*(*uint32)(unsafe.Add(mBase, _consts[934])) = uint32(v82)
	goto L12
L15:
	;
	v89 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v89)
	goto L17
L16:
	;
	goto L17
L17:
	;
	v92 = *(*int32)(unsafe.Add(mBase, _consts[933]))
	if v92 == int32(0) {
		v260 = v6
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v275 = *(*int64)(unsafe.Add(mBase, uint32(v23)+64))
	v276 = *(*int64)(unsafe.Add(mBase, uint32(v23)+72))
	*(*int64)(unsafe.Add(mBase, uint32(v23)+88)) = v276
	*(*int64)(unsafe.Add(mBase, uint32(v23)+80)) = v275
	v279 = int64(23)
	v282 = int64(2388976653695081527)
	v283 = (v275 ^ int64(base.Ui64(v275)>>(uint(v279)%64))) * v282
	v284 = int64(47)
	v289 = int64(-8645972361240307355)
	v295 = (v276 ^ int64(base.Ui64(v276)>>(uint(v279)%64))) * v282
	v301 = ((v283^int64(base.Ui64(v283)>>(uint(v284)%64))^int64(-9208349263878056368))*v289 ^ int64(base.Ui64(v295)>>(uint(v284)%64)) ^ v295) * v289
	v306 = (int64(base.Ui64(v301)>>(uint(v279)%64)) ^ v301) * v282
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v260)+16))
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v260)+8))
	v317 = v314
	v322 = v315
	goto L43
L19:
	;
	v96 = *(*int32)(unsafe.Add(mBase, _consts[278]))
	v97 = *(*int64)(unsafe.Add(mBase, uint32(v96)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v96)+16)) = v97
	v100 = int64(*(*int32)(unsafe.Add(mBase, _consts[934])))
	if v100 == v97 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v103 = *(*int32)(unsafe.Add(mBase, _consts[933]))
	v260 = v103
	goto L18
L21:
	;
	goto L22
L22:
	;
	v105 = *(*int32)(unsafe.Add(mBase, _consts[278]))
	v106 = *(*int64)(unsafe.Add(mBase, uint32(v105)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v105)+16)) = v106
	v110 = *(*int32)(unsafe.Add(mBase, _consts[933]))
	v111 = *(*int64)(unsafe.Add(mBase, uint32(v110)))
	if v111 == int64(0) {
		v156 = int32(-1)
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v168 = v156
	v172 = v110
	v174 = v6
	goto L29
L24:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v110)+20))
	v125 = int32(0)
	goto L25
L25:
	;
	v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114+v125*int32(24))+16)))
	if v139 != int32(1) {
		v156 = v125
		goto L23
	} else {
		goto L27
	}
L26:
	;
	v156 = int32(-1)
	goto L23
L27:
	;
	v143 = v125 + int32(1)
	if base.Ui64(base.I64_extend_i32_u(v143)) < base.Ui64(v111) {
		v125 = v143
		goto L25
	} else {
		goto L28
	}
L28:
	;
	goto L26
L29:
	;
	v188 = v168
	v194 = v174
	v197 = v174
	goto L32
L30:
	;
	*(*uint32)(unsafe.Add(mBase, _consts[934])) = uint32(v106)
	v260 = v172
	goto L18
L31:
	;
	goto L30
L32:
	;
	if v197&int32(1) != 0 {
		goto L31
	} else {
		goto L34
	}
L33:
	;
	if v221 == int32(0) {
		goto L31
	} else {
		goto L36
	}
L34:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v172)+12))
	v210 = int32(1)
	v211 = v188 - v210
	v215 = base.B2i32(v209&(v211^v156) == int32(0))
	v216 = v215 | v194
	v219 = v209 & v211
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v172)+20))
	v221 = v188*int32(24) + v220
	v222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v221)+16)))
	if v222 != v210 {
		v188 = v219
		v194 = v216
		v197 = v215
		goto L32
	} else {
		goto L35
	}
L35:
	;
	goto L33
L36:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v221)+20))
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v227)))
	v229 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v228)+16)))
	if v229 == int32(0) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v228)+24))
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v227)+8))
	if v232 == v233 {
		v168 = v219
		v174 = v216
		goto L29
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v227)+12))
	if v235 != 0 {
		v168 = v219
		v174 = v216
		goto L29
	} else {
		goto L41
	}
L40:
	;
	goto L39
L41:
	;
	v236 = *(*int64)(unsafe.Add(mBase, uint32(v221)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v23)+56)) = v236
	v238 = *(*int64)(unsafe.Add(mBase, uint32(v221)))
	*(*int64)(unsafe.Add(mBase, uint32(v23)+48)) = v238
	F_pgstat_release_entry_ref(m, v23+int32(48), v227, int32(0))
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L4
	} else {
		goto L42
	}
L42:
	;
	v246 = *(*int32)(unsafe.Add(mBase, _consts[933]))
	v168 = v219
	v172 = v246
	v174 = v216
	goto L29
L43:
	;
	if base.Ui32(v317) <= base.Ui32(v322) {
		goto L52
	} else {
		goto L53
	}
L45:
	;
	v1373 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v260)+16)) = v1373
	v317 = v1373
	v322 = v1359
	goto L43
L46:
	;
	v1325 = *(*int32)(unsafe.Add(mBase, _consts[280]))
	F_dshash_delete_entry(m, v1325, v1077)
	mBase = m.M
	v1327 = m.ExcPending
	if v1327 != 0 {
		goto L4
	} else {
		goto L215
	}
L47:
	;
	m.G0 = v23 + int32(96)
	return v1305
L48:
	;
	v1063 = *(*int32)(unsafe.Add(mBase, _consts[280]))
	v1067 = F_dshash_find(m, v1063, v23-int32(-64), int32(0))
	mBase = m.M
	v1068 = m.ExcPending
	if v1068 != 0 {
		goto L4
	} else {
		goto L164
	}
L49:
	;
	v1032 = *(*int32)(unsafe.Add(mBase, _consts[931]))
	v1034 = F_MemoryContextAlloc(m, v1032, int32(24))
	mBase = m.M
	v1035 = m.ExcPending
	if v1035 != 0 {
		goto L4
	} else {
		goto L163
	}
L50:
	;
	v1003 = *(*int32)(unsafe.Add(mBase, uint32(v260)+8))
	v1004 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v260)+8)) = v1003 + v1004
	*(*uint8)(unsafe.Add(mBase, uint32(v990)+16)) = uint8(v1004)
	*(*int64)(unsafe.Add(mBase, uint32(v990)+8)) = v276
	*(*int64)(unsafe.Add(mBase, uint32(v990))) = v275
	v1018 = v990
	goto L49
L51:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v973 = m.ExcPending
	if v973 != 0 {
		goto L4
	} else {
		goto L160
	}
L52:
	;
	v337 = *(*int64)(unsafe.Add(mBase, uint32(v260)))
	if v337 == int64(4294967296) {
		goto L51
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	v680 = int32(0)
	v681 = *(*int32)(unsafe.Add(mBase, uint32(v260)+20))
	v682 = *(*int32)(unsafe.Add(mBase, uint32(v260)+12))
	v683 = v682 & base.I32_wrap_i64(int64(base.Ui64(v306)>>(uint(v284)%64))^v306-int64(base.Ui64(v306)>>(uint(int64(32))%64)))
	v686 = v681 + v683*int32(24)
	v687 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v686)+16)))
	if v687 == v680 {
		v990 = v686
		goto L50
	} else {
		goto L107
	}
L55:
	;
	v340 = int32(0)
	v342 = int64(2)
	v344 = v337 << (uint(int64(1)) % 64)
	if base.Ui64(v344) <= base.Ui64(v342) {
		goto L58
	} else {
		goto L59
	}
L56:
	;
	goto L54
L57:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v650 = m.ExcPending
	if v650 != 0 {
		goto L4
	} else {
		goto L104
	}
L58:
	;
	v347 = v342
	goto L60
L59:
	;
	v347 = v344
	goto L60
L60:
	;
	v348 = int64(1)
	if v347&(v347-v348) == int64(0) {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v358 = v347
	goto L63
L62:
	;
	v358 = v348 << (uint(int64(64)-base.I64_clz(v347)) % 64)
	goto L63
L63:
	;
	if base.Ui64(v358*int64(24)) < base.Ui64(int64(2147483647)) {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v363 = *(*int32)(unsafe.Add(mBase, uint32(v260)+20))
	v364 = *(*int64)(unsafe.Add(mBase, uint32(v260)))
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v260)+24))
	v370 = F_MemoryContextAllocExtended(m, v365, base.I32_wrap_i64(v358)*int32(24), int32(5))
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
		goto L4
	} else {
		goto L67
	}
L65:
	;
	goto L66
L66:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v637 = m.ExcPending
	if v637 != 0 {
		goto L4
	} else {
		goto L101
	}
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v260)+20)) = v370
	v373 = int64(1)
	if v358&(v358-v373) == int64(0) {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v383 = v358
	goto L70
L69:
	;
	v383 = v373 << (uint(int64(64)-base.I64_clz(v358)) % 64)
	goto L70
L70:
	;
	if base.Ui64(int64(2147483647)) <= base.Ui64(v383*int64(24)) {
		goto L57
	} else {
		goto L71
	}
L71:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v260))) = v383
	v391 = base.I32_wrap_i64(v383) - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v260)+12)) = v391
	v398 = base.F64_mul(base.F64_convert_i64_u(v383), float64(0.9))
	if base.F64_lt(v398, float64(4.294967296e+09))&base.F64_ge(v398, float64(0)) != 0 {
		goto L73
	} else {
		goto L74
	}
L72:
	;
	if v383 == int64(4294967296) {
		goto L76
	} else {
		goto L77
	}
L73:
	;
	v404 = base.I32_trunc_f64_u(v398)
	v406 = v404
	goto L72
L74:
	;
	goto L75
L75:
	;
	v406 = int32(0)
	goto L72
L76:
	;
	v407 = int32(-85899346)
	goto L78
L77:
	;
	v407 = v406
	goto L78
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v260)+16)) = v407
	if v364 != int64(0) {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v417 = v340
	goto L83
L80:
	;
	goto L81
L81:
	;
	F_pfree(m, v363)
	mBase = m.M
	v633 = m.ExcPending
	if v633 != 0 {
		goto L4
	} else {
		goto L100
	}
L82:
	;
	v489 = v482
	v494 = v340
	goto L88
L83:
	;
	v433 = v363 + v417*int32(24)
	v434 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v433)+16)))
	if v434 != int32(1) {
		v482 = v417
		goto L82
	} else {
		goto L85
	}
L84:
	;
	v482 = int32(0)
	goto L82
L85:
	;
	v437 = *(*int64)(unsafe.Add(mBase, uint32(v433)))
	v438 = int64(23)
	v441 = int64(2388976653695081527)
	v442 = (int64(base.Ui64(v437)>>(uint(v438)%64)) ^ v437) * v441
	v443 = int64(47)
	v448 = int64(-8645972361240307355)
	v450 = *(*int64)(unsafe.Add(mBase, uint32(v433)+8))
	v455 = (int64(base.Ui64(v450)>>(uint(v438)%64)) ^ v450) * v441
	v461 = ((v442^int64(base.Ui64(v442)>>(uint(v443)%64))^int64(-9208349263878056368))*v448 ^ int64(base.Ui64(v455)>>(uint(v443)%64)) ^ v455) * v448
	v466 = (int64(base.Ui64(v461)>>(uint(v438)%64)) ^ v461) * v441
	if v391&base.I32_wrap_i64(int64(base.Ui64(v466)>>(uint(v443)%64))^v466-int64(base.Ui64(v466)>>(uint(int64(32))%64))) == v417 {
		v482 = v417
		goto L82
	} else {
		goto L86
	}
L86:
	;
	v477 = v417 + int32(1)
	if base.Ui64(base.I64_extend_i32_u(v477)) < base.Ui64(v364) {
		v417 = v477
		goto L83
	} else {
		goto L87
	}
L87:
	;
	goto L84
L88:
	;
	v505 = v363 + v489*int32(24)
	v506 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v505)+16)))
	if v506 == int32(1) {
		goto L90
	} else {
		goto L91
	}
L89:
	;
	goto L81
L90:
	;
	v509 = *(*int64)(unsafe.Add(mBase, uint32(v505)))
	v510 = int64(23)
	v513 = int64(2388976653695081527)
	v514 = (int64(base.Ui64(v509)>>(uint(v510)%64)) ^ v509) * v513
	v515 = int64(47)
	v520 = int64(-8645972361240307355)
	v523 = v505 + int32(8)
	v524 = *(*int64)(unsafe.Add(mBase, uint32(v523)))
	v529 = (int64(base.Ui64(v524)>>(uint(v510)%64)) ^ v524) * v513
	v535 = ((v514^int64(base.Ui64(v514)>>(uint(v515)%64))^int64(-9208349263878056368))*v520 ^ int64(base.Ui64(v529)>>(uint(v515)%64)) ^ v529) * v520
	v540 = (int64(base.Ui64(v535)>>(uint(v510)%64)) ^ v535) * v513
	v548 = *(*int32)(unsafe.Add(mBase, uint32(v260)+12))
	v550 = base.I32_wrap_i64(int64(base.Ui64(v540)>>(uint(v515)%64)) ^ v540 - int64(base.Ui64(v540)>>(uint(int64(32))%64)))
	goto L93
L91:
	;
	goto L92
L92:
	;
	v603 = v489 + int32(1)
	if base.Ui64(base.I64_extend_i32_u(v603)) < base.Ui64(v364) {
		goto L96
	} else {
		goto L97
	}
L93:
	;
	v569 = v550 & v548
	v574 = v370 + v569*int32(24)
	v575 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v574)+16)))
	if v575 != 0 {
		v550 = v569 + int32(1)
		goto L93
	} else {
		goto L95
	}
L94:
	;
	v576 = *(*int64)(unsafe.Add(mBase, uint32(v505)))
	*(*int64)(unsafe.Add(mBase, uint32(v574))) = v576
	v578 = *(*int64)(unsafe.Add(mBase, uint32(v505)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v574)+16)) = v578
	v580 = *(*int64)(unsafe.Add(mBase, uint32(v523)))
	*(*int64)(unsafe.Add(mBase, uint32(v574)+8)) = v580
	goto L92
L95:
	;
	goto L94
L96:
	;
	v607 = v603
	goto L98
L97:
	;
	v607 = int32(0)
	goto L98
L98:
	;
	v609 = v494 + int32(1)
	if base.Ui64(base.I64_extend_i32_u(v609)) < base.Ui64(v364) {
		v489 = v607
		v494 = v609
		goto L88
	} else {
		goto L99
	}
L99:
	;
	goto L89
L100:
	;
	goto L56
L101:
	;
	F_errmsg_internal(m, int32(418308), int32(0))
	mBase = m.M
	v641 = m.ExcPending
	if v641 != 0 {
		goto L4
	} else {
		goto L102
	}
L102:
	;
	F_errfinish(m, int32(342023), int32(327), int32(357131))
	mBase = m.M
	v646 = m.ExcPending
	if v646 != 0 {
		goto L4
	} else {
		goto L103
	}
L103:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L104:
	;
	F_errmsg_internal(m, int32(418308), int32(0))
	mBase = m.M
	v654 = m.ExcPending
	if v654 != 0 {
		goto L4
	} else {
		goto L105
	}
L105:
	;
	F_errfinish(m, int32(342023), int32(327), int32(357131))
	mBase = m.M
	v659 = m.ExcPending
	if v659 != 0 {
		goto L4
	} else {
		goto L106
	}
L106:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L107:
	;
	v691 = v683
	v697 = v686
	v700 = v680
	goto L108
L108:
	;
	v711 = v23 + int32(80)
	v712 = int32(16)
	goto L113
L109:
	;
	v966 = *(*int32)(unsafe.Add(mBase, uint32(v697)+20))
	if v966 == int32(0) {
		v1018 = v697
		goto L49
	} else {
		goto L158
	}
L110:
	;
	if v774 != 0 {
		goto L128
	} else {
		goto L129
	}
L111:
	;
	v774 = int32(0)
	goto L110
L112:
	;
	v748 = v743
	v749 = v744
	v750 = v745
	goto L122
L113:
	;
	if (v697|v711)&int32(3) != 0 {
		v743 = v697
		v744 = v711
		v745 = v712
		goto L112
	} else {
		goto L116
	}
L115:
	;
	if v733 == int32(0) {
		goto L111
	} else {
		goto L121
	}
L116:
	;
	v720 = v697
	v721 = v711
	v722 = v712
	goto L117
L117:
	;
	v725 = *(*int32)(unsafe.Add(mBase, uint32(v720)))
	v726 = *(*int32)(unsafe.Add(mBase, uint32(v721)))
	if v725 != v726 {
		v743 = v720
		v744 = v721
		v745 = v722
		goto L112
	} else {
		goto L119
	}
L118:
	;
	goto L115
L119:
	;
	v728 = int32(4)
	v729 = v721 + v728
	v731 = v720 + v728
	v733 = v722 - v728
	if base.Ui32(int32(3)) < base.Ui32(v733) {
		v720 = v731
		v721 = v729
		v722 = v733
		goto L117
	} else {
		goto L120
	}
L120:
	;
	goto L118
L121:
	;
	v743 = v731
	v744 = v729
	v745 = v733
	goto L112
L122:
	;
	v753 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v748))))
	v754 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v749))))
	if v753 == v754 {
		goto L124
	} else {
		goto L125
	}
L123:
	;
	v774 = v753 - v754
	goto L110
L124:
	;
	v756 = int32(1)
	v761 = v750 - v756
	if v761 != 0 {
		v748 = v748 + v756
		v749 = v749 + v756
		v750 = v761
		goto L122
	} else {
		goto L127
	}
L125:
	;
	goto L126
L126:
	;
	goto L123
L127:
	;
	goto L111
L128:
	;
	v776 = v691 + int32(1)
	v777 = *(*int64)(unsafe.Add(mBase, uint32(v697)))
	v778 = int64(23)
	v781 = int64(2388976653695081527)
	v782 = (int64(base.Ui64(v777)>>(uint(v778)%64)) ^ v777) * v781
	v783 = int64(47)
	v788 = int64(-8645972361240307355)
	v790 = *(*int64)(unsafe.Add(mBase, uint32(v697)+8))
	v795 = (int64(base.Ui64(v790)>>(uint(v778)%64)) ^ v790) * v781
	v801 = ((v782^int64(base.Ui64(v782)>>(uint(v783)%64))^int64(-9208349263878056368))*v788 ^ int64(base.Ui64(v795)>>(uint(v783)%64)) ^ v795) * v788
	v806 = (int64(base.Ui64(v801)>>(uint(v778)%64)) ^ v801) * v781
	v814 = v682 & base.I32_wrap_i64(int64(base.Ui64(v806)>>(uint(v783)%64))^v806-int64(base.Ui64(v806)>>(uint(int64(32))%64)))
	if base.Ui32(v691) < base.Ui32(v814) {
		goto L131
	} else {
		goto L132
	}
L129:
	;
	goto L130
L130:
	;
	goto L109
L131:
	;
	v816 = *(*int32)(unsafe.Add(mBase, uint32(v260)))
	v818 = v691 + v816
	goto L133
L132:
	;
	v818 = v691
	goto L133
L133:
	;
	if base.Ui32(v818-v814) < base.Ui32(v700) {
		goto L134
	} else {
		goto L135
	}
L134:
	;
	v822 = v682 & v776
	v825 = v681 + v822*int32(24)
	v826 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v825)+16)))
	if v826 != 0 {
		goto L137
	} else {
		goto L138
	}
L135:
	;
	goto L136
L136:
	;
	v950 = v700 + int32(1)
	if base.Ui32(int32(26)) <= base.Ui32(v950) {
		goto L153
	} else {
		goto L154
	}
L137:
	;
	v837 = v822
	v838 = int32(0)
	goto L140
L138:
	;
	v872 = v825
	v876 = v822
	goto L139
L139:
	;
	if v691 != v876 {
		goto L147
	} else {
		goto L148
	}
L140:
	;
	v848 = v838 + int32(1)
	if int32(151) <= v848 {
		goto L142
	} else {
		goto L143
	}
L141:
	;
	v872 = v864
	v876 = v861
	goto L139
L142:
	;
	v851 = *(*int32)(unsafe.Add(mBase, uint32(v260)+8))
	v853 = *(*int64)(unsafe.Add(mBase, uint32(v260)))
	if base.F64_ge(base.F64_div(base.F64_convert_i32_u(v851), base.F64_convert_i64_u(v853)), float64(0.1)) != 0 {
		v1359 = v851
		goto L45
	} else {
		goto L145
	}
L143:
	;
	goto L144
L144:
	;
	v861 = (v837 + int32(1)) & v682
	v864 = v681 + v861*int32(24)
	v865 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v864)+16)))
	if v865 != 0 {
		v837 = v861
		v838 = v848
		goto L140
	} else {
		goto L146
	}
L145:
	;
	goto L144
L146:
	;
	goto L141
L147:
	;
	v893 = v872
	v897 = v876
	goto L150
L148:
	;
	goto L149
L149:
	;
	v941 = *(*int32)(unsafe.Add(mBase, uint32(v260)+8))
	v942 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v260)+8)) = v941 + v942
	*(*uint8)(unsafe.Add(mBase, uint32(v697)+16)) = uint8(v942)
	*(*int64)(unsafe.Add(mBase, uint32(v697)+8)) = v276
	*(*int64)(unsafe.Add(mBase, uint32(v697))) = v275
	v1018 = v697
	goto L49
L150:
	;
	v907 = *(*int32)(unsafe.Add(mBase, uint32(v260)+12))
	v910 = v907 & (v897 - int32(1))
	v913 = v681 + v910*int32(24)
	v914 = *(*int64)(unsafe.Add(mBase, uint32(v913)))
	*(*int64)(unsafe.Add(mBase, uint32(v893))) = v914
	v916 = *(*int64)(unsafe.Add(mBase, uint32(v913)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v893)+16)) = v916
	v918 = *(*int64)(unsafe.Add(mBase, uint32(v913)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v893)+8)) = v918
	if v691 != v910 {
		v893 = v913
		v897 = v910
		goto L150
	} else {
		goto L152
	}
L151:
	;
	goto L149
L152:
	;
	goto L151
L153:
	;
	v953 = *(*int32)(unsafe.Add(mBase, uint32(v260)+8))
	v955 = *(*int64)(unsafe.Add(mBase, uint32(v260)))
	if base.F64_ge(base.F64_div(base.F64_convert_i32_u(v953), base.F64_convert_i64_u(v955)), float64(0.1)) != 0 {
		v1359 = v953
		goto L45
	} else {
		goto L156
	}
L154:
	;
	goto L155
L155:
	;
	v961 = v682 & v776
	v964 = v681 + v961*int32(24)
	v965 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v964)+16)))
	if v965 != 0 {
		v691 = v961
		v697 = v964
		v700 = v950
		goto L108
	} else {
		goto L157
	}
L156:
	;
	goto L155
L157:
	;
	v990 = v964
	goto L50
L158:
	;
	v969 = *(*int32)(unsafe.Add(mBase, uint32(v966)+4))
	if v969 != 0 {
		v1305 = v966
		goto L47
	} else {
		goto L159
	}
L159:
	;
	v1047 = v966
	goto L48
L160:
	;
	F_errmsg_internal(m, int32(482457), int32(0))
	mBase = m.M
	v977 = m.ExcPending
	if v977 != 0 {
		goto L4
	} else {
		goto L161
	}
L161:
	;
	F_errfinish(m, int32(342023), int32(630), int32(325730))
	mBase = m.M
	v982 = m.ExcPending
	if v982 != 0 {
		goto L4
	} else {
		goto L162
	}
L162:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L163:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1018)+20)) = v1034
	*(*int32)(unsafe.Add(mBase, uint32(v1034)+12)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v1034))) = int64(0)
	v1041 = *(*int32)(unsafe.Add(mBase, uint32(v1018)+20))
	v1047 = v1041
	goto L48
L164:
	;
	if l3 == int32(0) {
		v1154 = v1067
		goto L165
	} else {
		goto L166
	}
L165:
	;
	if v1154 == int32(0) {
		goto L183
	} else {
		goto L184
	}
L166:
	;
	if v1067 != 0 {
		v1154 = v1067
		goto L165
	} else {
		goto L167
	}
L167:
	;
	v1072 = *(*int32)(unsafe.Add(mBase, _consts[280]))
	v1077 = F_dshash_find_or_insert(m, v1072, v23-int32(-64), v23+int32(80))
	mBase = m.M
	v1078 = m.ExcPending
	if v1078 != 0 {
		goto L4
	} else {
		goto L168
	}
L168:
	;
	v1079 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+80)))
	if v1079 != 0 {
		v1154 = v1077
		goto L165
	} else {
		goto L169
	}
L169:
	;
	v1080 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1077)+20)) = v1080
	v1082 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1077)+16)) = uint8(v1082)
	*(*int32)(unsafe.Add(mBase, uint32(v1077)+24)) = v1082
	v1087 = *(*int32)(unsafe.Add(mBase, _consts[281]))
	if base.Ui32(l0-v1080) <= base.Ui32(int32(11)) {
		goto L171
	} else {
		goto L172
	}
L170:
	;
	v1117 = *(*int32)(unsafe.Add(mBase, uint32(v1116)+4))
	v1119 = F_dsa_allocate_extended(m, v1087, v1117, int32(6))
	mBase = m.M
	v1120 = m.ExcPending
	if v1120 != 0 {
		goto L4
	} else {
		goto L177
	}
L171:
	;
	v1116 = l0*int32(72) + int32(1683808)
	goto L170
L172:
	;
	goto L173
L173:
	;
	if base.Ui32(int32(8)) < base.Ui32(l0-int32(24)) {
		v1114 = int32(0)
		goto L174
	} else {
		goto L175
	}
L174:
	;
	v1116 = v1114
	goto L170
L175:
	;
	v1102 = int32(0)
	v1104 = *(*int32)(unsafe.Add(mBase, _consts[279]))
	if v1104 == v1102 {
		v1114 = v1102
		goto L174
	} else {
		goto L176
	}
L176:
	;
	v1112 = *(*int32)(unsafe.Add(mBase, uint32(v1104+l0<<(uint(int32(2))%32)-int32(96))))
	v1114 = v1112
	goto L174
L177:
	;
	if v1119 == int32(0) {
		goto L46
	} else {
		goto L178
	}
L178:
	;
	v1124 = *(*int32)(unsafe.Add(mBase, _consts[281]))
	v1125 = F_dsa_get_address(m, v1124, v1119)
	mBase = m.M
	v1126 = m.ExcPending
	if v1126 != 0 {
		goto L4
	} else {
		goto L179
	}
L179:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1125))) = int32(-559038737)
	*(*int32)(unsafe.Add(mBase, uint32(v1077)+28)) = v1119
	v1131 = v1125 + int32(4)
	v1132 = int32(81)
	*(*uint16)(unsafe.Add(mBase, uint32(v1131))) = uint16(v1132)
	*(*int32)(unsafe.Add(mBase, uint32(v1131)+4)) = int32(1073741824)
	*(*int64)(unsafe.Add(mBase, uint32(v1131)+8)) = int64(-1)
	goto L180
L180:
	;
	v1138 = *(*int32)(unsafe.Add(mBase, uint32(v1077)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v1077)+20)) = v1138 + int32(1)
	v1143 = *(*int32)(unsafe.Add(mBase, _consts[280]))
	F_dshash_release_lock(m, v1143, v1077)
	mBase = m.M
	v1145 = m.ExcPending
	if v1145 != 0 {
		goto L4
	} else {
		goto L181
	}
L181:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1047))) = v1077
	*(*int32)(unsafe.Add(mBase, uint32(v1047)+4)) = v1125
	v1148 = *(*int32)(unsafe.Add(mBase, uint32(v1077)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v1047)+8)) = v1148
	if l4 == int32(0) {
		v1305 = v1047
		goto L47
	} else {
		goto L182
	}
L182:
	;
	v1152 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v1152)
	v1305 = v1047
	goto L47
L183:
	;
	v1157 = *(*int64)(unsafe.Add(mBase, uint32(v23)+72))
	*(*int64)(unsafe.Add(mBase, uint32(v23)+8)) = v1157
	v1159 = *(*int64)(unsafe.Add(mBase, uint32(v23)+64))
	*(*int64)(unsafe.Add(mBase, uint32(v23))) = v1159
	F_pgstat_release_entry_ref(m, v23, v1047, int32(0))
	mBase = m.M
	v1163 = m.ExcPending
	if v1163 != 0 {
		goto L4
	} else {
		goto L186
	}
L184:
	;
	goto L185
L185:
	;
	v1165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1154)+16)))
	if l3 == int32(0) {
		goto L187
	} else {
		goto L188
	}
L186:
	;
	v1305 = int32(0)
	goto L47
L187:
	;
	if v1165&int32(1) != 0 {
		goto L208
	} else {
		goto L209
	}
L188:
	;
	if v1165&int32(1) == int32(0) {
		goto L187
	} else {
		goto L189
	}
L189:
	;
	v1173 = *(*int32)(unsafe.Add(mBase, _consts[281]))
	v1174 = *(*int32)(unsafe.Add(mBase, uint32(v1154)+28))
	v1175 = F_dsa_get_address(m, v1173, v1174)
	mBase = m.M
	v1176 = m.ExcPending
	if v1176 != 0 {
		goto L4
	} else {
		goto L190
	}
L190:
	;
	v1177 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1154)+16)) = uint8(v1177)
	v1179 = *(*int32)(unsafe.Add(mBase, uint32(v1154)+20))
	v1180 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1154)+20)) = v1179 + v1180
	v1183 = *(*int32)(unsafe.Add(mBase, uint32(v1154)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v1154)+24)) = v1183 + v1180
	if base.Ui32(l0-v1180) <= base.Ui32(int32(11)) {
		goto L192
	} else {
		goto L193
	}
L191:
	;
	v1216 = *(*int32)(unsafe.Add(mBase, uint32(v1215)+16))
	if base.Ui32(l0-int32(1)) <= base.Ui32(int32(11)) {
		goto L199
	} else {
		goto L200
	}
L192:
	;
	v1215 = l0*int32(72) + int32(1683808)
	goto L191
L193:
	;
	goto L194
L194:
	;
	if base.Ui32(int32(8)) < base.Ui32(l0-int32(24)) {
		v1213 = int32(0)
		goto L195
	} else {
		goto L196
	}
L195:
	;
	v1215 = v1213
	goto L191
L196:
	;
	v1201 = int32(0)
	v1203 = *(*int32)(unsafe.Add(mBase, _consts[279]))
	if v1203 == v1201 {
		v1213 = v1201
		goto L195
	} else {
		goto L197
	}
L197:
	;
	v1211 = *(*int32)(unsafe.Add(mBase, uint32(v1203+l0<<(uint(int32(2))%32)-int32(96))))
	v1213 = v1211
	goto L195
L198:
	;
	v1248 = *(*int32)(unsafe.Add(mBase, uint32(v1247)+20))
	v1250 = F__emscripten_memset_bulkmem(m, v1175+v1216, base.I32_extend8_s(int32(0)), v1248)
	mBase = m.M
	goto L205
L199:
	;
	v1247 = l0*int32(72) + int32(1683808)
	goto L198
L200:
	;
	goto L201
L201:
	;
	if base.Ui32(int32(8)) < base.Ui32(l0-int32(24)) {
		v1245 = int32(0)
		goto L202
	} else {
		goto L203
	}
L202:
	;
	v1247 = v1245
	goto L198
L203:
	;
	v1233 = int32(0)
	v1235 = *(*int32)(unsafe.Add(mBase, _consts[279]))
	if v1235 == v1233 {
		v1245 = v1233
		goto L202
	} else {
		goto L204
	}
L204:
	;
	v1243 = *(*int32)(unsafe.Add(mBase, uint32(v1235+l0<<(uint(int32(2))%32)-int32(96))))
	v1245 = v1243
	goto L202
L205:
	;
	v1251 = *(*int32)(unsafe.Add(mBase, uint32(v1154)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v1154)+20)) = v1251 + int32(1)
	v1256 = *(*int32)(unsafe.Add(mBase, _consts[280]))
	F_dshash_release_lock(m, v1256, v1154)
	mBase = m.M
	v1258 = m.ExcPending
	if v1258 != 0 {
		goto L4
	} else {
		goto L206
	}
L206:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1047))) = v1154
	*(*int32)(unsafe.Add(mBase, uint32(v1047)+4)) = v1175
	v1261 = *(*int32)(unsafe.Add(mBase, uint32(v1154)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v1047)+8)) = v1261
	if l4 == int32(0) {
		v1305 = v1047
		goto L47
	} else {
		goto L207
	}
L207:
	;
	v1265 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v1265)
	v1305 = v1047
	goto L47
L208:
	;
	v1270 = *(*int32)(unsafe.Add(mBase, _consts[280]))
	F_dshash_release_lock(m, v1270, v1154)
	mBase = m.M
	v1272 = m.ExcPending
	if v1272 != 0 {
		goto L4
	} else {
		goto L211
	}
L209:
	;
	goto L210
L210:
	;
	v1284 = *(*int32)(unsafe.Add(mBase, _consts[281]))
	v1285 = *(*int32)(unsafe.Add(mBase, uint32(v1154)+28))
	v1286 = F_dsa_get_address(m, v1284, v1285)
	mBase = m.M
	v1287 = m.ExcPending
	if v1287 != 0 {
		goto L4
	} else {
		goto L213
	}
L211:
	;
	v1273 = *(*int64)(unsafe.Add(mBase, uint32(v23)+72))
	*(*int64)(unsafe.Add(mBase, uint32(v23)+24)) = v1273
	v1275 = *(*int64)(unsafe.Add(mBase, uint32(v23)+64))
	*(*int64)(unsafe.Add(mBase, uint32(v23)+16)) = v1275
	F_pgstat_release_entry_ref(m, v23+int32(16), v1047, int32(0))
	mBase = m.M
	v1281 = m.ExcPending
	if v1281 != 0 {
		goto L4
	} else {
		goto L212
	}
L212:
	;
	v1305 = int32(0)
	goto L47
L213:
	;
	v1288 = *(*int32)(unsafe.Add(mBase, uint32(v1154)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v1154)+20)) = v1288 + int32(1)
	v1293 = *(*int32)(unsafe.Add(mBase, _consts[280]))
	F_dshash_release_lock(m, v1293, v1154)
	mBase = m.M
	v1295 = m.ExcPending
	if v1295 != 0 {
		goto L4
	} else {
		goto L214
	}
L214:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1047))) = v1154
	*(*int32)(unsafe.Add(mBase, uint32(v1047)+4)) = v1286
	v1298 = *(*int32)(unsafe.Add(mBase, uint32(v1154)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v1047)+8)) = v1298
	v1305 = v1047
	goto L47
L215:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1331 = m.ExcPending
	if v1331 != 0 {
		goto L4
	} else {
		goto L216
	}
L216:
	;
	F_errcode(m, int32(8389))
	mBase = m.M
	v1334 = m.ExcPending
	if v1334 != 0 {
		goto L4
	} else {
		goto L217
	}
L217:
	;
	F_errmsg(m, int32(14012), int32(0))
	mBase = m.M
	v1338 = m.ExcPending
	if v1338 != 0 {
		goto L4
	} else {
		goto L218
	}
L218:
	;
	v1339 = *(*int64)(unsafe.Add(mBase, uint32(v23)+64))
	*(*int64)(unsafe.Add(mBase, uint32(v23)+32)) = v1339
	v1341 = *(*int64)(unsafe.Add(mBase, uint32(v23)+72))
	*(*int64)(unsafe.Add(mBase, uint32(v23)+40)) = v1341
	F_errdetail(m, int32(602933), v23+int32(32))
	mBase = m.M
	v1347 = m.ExcPending
	if v1347 != 0 {
		goto L4
	} else {
		goto L219
	}
L219:
	;
	F_errfinish(m, int32(519971), int32(537), int32(354922))
	mBase = m.M
	v1352 = m.ExcPending
	if v1352 != 0 {
		goto L4
	} else {
		goto L220
	}
L220:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pgstat_get_slru_name(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	if base.Ui32(l0) <= base.Ui32(int32(7)) {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0<<(uint(int32(2))%32))+uint32(_consts[935])))
		v10 = v9
	} else {
		v10 = int32(0)
	}
	return v10
}
func F_pgstat_get_transactional_drops(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int64
	_ = v50
	var v52 int64
	_ = v52
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	v3 = int32(0)
	v9 = *(*int32)(unsafe.Add(mBase, _consts[946]))
	if v9 == v3 {
		v64 = v3
		return v64
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
		v15 = F_palloc(m, v12<<(uint(int32(4))%32))
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l1))) = v15
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
			if v20 == int32(0) {
				v64 = v3
			} else {
				v24 = v9 + int32(8)
				if v20 == v24 {
					v64 = v3
				} else {
					v29 = v3
					v30 = v20
					for {
						v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30-int32(4)))))
						if l0 != 0 {
							if v35&int32(1) == int32(0) {
								v44 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
								v47 = v44 + v29<<(uint(int32(4))%32)
								v49 = v30 - int32(20)
								v50 = *(*int64)(unsafe.Add(mBase, uint32(v49)))
								*(*int64)(unsafe.Add(mBase, uint32(v47))) = v50
								v52 = *(*int64)(unsafe.Add(mBase, uint32(v49)+8))
								*(*int64)(unsafe.Add(mBase, uint32(v47)+8)) = v52
								v57 = v29 + int32(1)
							} else {
								v57 = v29
							}
						} else {
							if v35&int32(1) == int32(0) {
								v57 = v29
							} else {
								v44 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
								v47 = v44 + v29<<(uint(int32(4))%32)
								v49 = v30 - int32(20)
								v50 = *(*int64)(unsafe.Add(mBase, uint32(v49)))
								*(*int64)(unsafe.Add(mBase, uint32(v47))) = v50
								v52 = *(*int64)(unsafe.Add(mBase, uint32(v49)+8))
								*(*int64)(unsafe.Add(mBase, uint32(v47)+8)) = v52
								v57 = v29 + int32(1)
							}
						}
						v59 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
						if v59 != v24 {
							v29 = v57
							v30 = v59
							continue
						} else {
							break
						}
						break
					}
					v64 = v57
				}
			}
			return v64
		}
	}
}
func F_pgstat_get_wait_event(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v16 int32
	_ = v16
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v156 int32
	_ = v156
	v2 = int32(0)
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	if l0 == v2 {
		v138 = v2
		m.G0 = v6 + int32(16)
		return v138
	} else {
		v10 = int32(97313)
		switch int32(base.Ui32(l0-int32(16777216)) >> (uint(int32(24)) % 32)) {
		case 0:
			v16 = l0 & int32(65535)
			if base.Ui32(v16) <= base.Ui32(int32(94)) {
				v23 = *(*int32)(unsafe.Add(mBase, uint32(v16<<(uint(int32(2))%32))+uint32(_consts[947])))
				v43 = v23
			} else {
				v27 = (v16 - int32(95)) & int32(65535)
				v29 = *(*int32)(unsafe.Add(mBase, _consts[948]))
				if v27 < v29 {
					v32 = *(*int32)(unsafe.Add(mBase, _consts[949]))
					v36 = *(*int32)(unsafe.Add(mBase, uint32(v32+v27<<(uint(int32(2))%32))))
					if v36 != 0 {
						v38 = v36
					} else {
						v38 = int32(283744)
					}
					v41 = v38
				} else {
					v41 = int32(283744)
				}
				v43 = v41
			}
			v138 = v43
			m.G0 = v6 + int32(16)
			return v138
		default:
			v138 = v10
			m.G0 = v6 + int32(16)
			return v138
		case 2:
			v45 = l0 & int32(65535)
			if base.Ui32(v45) <= base.Ui32(int32(11)) {
				v52 = *(*int32)(unsafe.Add(mBase, uint32(v45<<(uint(int32(2))%32))+uint32(_consts[950])))
				v54 = v52
			} else {
				v54 = int32(570470)
			}
			v138 = v54
			m.G0 = v6 + int32(16)
			return v138
		case 3:
			if l0 == int32(67108864) {
				v91 = int32(292465)
			} else {
				v91 = int32(97313)
			}
			v138 = v91
			m.G0 = v6 + int32(16)
			return v138
		case 4:
			v93 = l0 - int32(83886080)
			if base.Ui32(int32(18)) <= base.Ui32(v93) {
				v138 = v10
			} else {
				v100 = *(*int32)(unsafe.Add(mBase, uint32(v93<<(uint(int32(2))%32))+uint32(_consts[951])))
				v138 = v100
			}
			m.G0 = v6 + int32(16)
			return v138
		case 5:
			v102 = l0 - int32(100663296)
			if base.Ui32(int32(9)) <= base.Ui32(v102) {
				v138 = v10
			} else {
				v109 = *(*int32)(unsafe.Add(mBase, uint32(v102<<(uint(int32(2))%32))+uint32(_consts[952])))
				v138 = v109
			}
			m.G0 = v6 + int32(16)
			return v138
		case 6, 10:
			*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = l0
			if l0 == int32(117440512) {
				v138 = int32(283829)
				m.G0 = v6 + int32(16)
				return v138
			} else {
				v60 = *(*int32)(unsafe.Add(mBase, _consts[47]))
				v64 = F_LWLockAcquire(m, v60+int32(6144), int32(1))
				mBase = m.M
				v67 = m.ExcPending
				if v67 != 0 {
					return int32(0)
				} else {
					v69 = *(*int32)(unsafe.Add(mBase, _consts[953]))
					v75 = F_hash_search(m, v69, v6+int32(12), int32(0), v6+int32(11))
					mBase = m.M
					v76 = m.ExcPending
					if v76 != 0 {
						return int32(0)
					} else {
						v78 = *(*int32)(unsafe.Add(mBase, _consts[47]))
						F_LWLockRelease(m, v78+int32(6144))
						mBase = m.M
						v82 = m.ExcPending
						if v82 != 0 {
							return int32(0)
						} else {
							if v75 == int32(0) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v146 = m.ExcPending
								if v146 != 0 {
									return int32(0)
								} else {
									v147 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
									*(*int32)(unsafe.Add(mBase, uint32(v6))) = v147
									F_errmsg_internal(m, int32(49333), v6)
									mBase = m.M
									v151 = m.ExcPending
									if v151 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(514823), int32(295), int32(232872))
										mBase = m.M
										v156 = m.ExcPending
										if v156 != 0 {
											return int32(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							} else {
								v138 = v75 + int32(4)
								m.G0 = v6 + int32(16)
								return v138
							}
						}
					}
				}
			}
		case 7:
			v111 = l0 + int32(-134217728)
			if base.Ui32(int32(57)) <= base.Ui32(v111) {
				v138 = v10
			} else {
				v118 = *(*int32)(unsafe.Add(mBase, uint32(v111<<(uint(int32(2))%32))+uint32(_consts[954])))
				v138 = v118
			}
			m.G0 = v6 + int32(16)
			return v138
		case 8:
			v120 = l0 - int32(150994944)
			if base.Ui32(int32(10)) <= base.Ui32(v120) {
				v138 = v10
			} else {
				v127 = *(*int32)(unsafe.Add(mBase, uint32(v120<<(uint(int32(2))%32))+uint32(_consts[955])))
				v138 = v127
			}
			m.G0 = v6 + int32(16)
			return v138
		case 9:
			v129 = l0 - int32(167772160)
			if base.Ui32(int32(81)) <= base.Ui32(v129) {
				v138 = v10
			} else {
				v136 = *(*int32)(unsafe.Add(mBase, uint32(v129<<(uint(int32(2))%32))+uint32(_consts[956])))
				v138 = v136
			}
			m.G0 = v6 + int32(16)
			return v138
		}
	}
}
func F_pgstat_io_snapshot_cb(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int64
	_ = v15
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	v6 = *(*int32)(unsafe.Add(mBase, _consts[278]))
	v8 = v6 + int32(608)
	v10 = F_LWLockAcquire(m, v8, int32(1))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v14 = *(*int32)(unsafe.Add(mBase, _consts[278]))
	v15 = *(*int64)(unsafe.Add(mBase, uint32(v14)+896))
	*(*int64)(unsafe.Add(mBase, _consts[930])) = v15
	goto L4
L3:
	;
	F_LWLockRelease(m, v8)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L7
	}
L4:
	;
	v21 = F__emscripten_memcpy_bulkmem(m, int32(4478088), v6+int32(904), int32(2880))
	mBase = m.M
	goto L6
L6:
	;
	goto L3
L7:
	;
	v26 = int32(1)
	goto L8
L8:
	;
	v31 = *(*int32)(unsafe.Add(mBase, _consts[278]))
	v36 = v31 + v26<<(uint(int32(4))%32) + int32(608)
	v38 = F_LWLockAcquire(m, v36, int32(1))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L1
	} else {
		goto L10
	}
L9:
	;
	return
L10:
	;
	v40 = int32(2880)
	v41 = v26 * v40
	goto L12
L11:
	;
	F_LWLockRelease(m, v36)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L1
	} else {
		goto L15
	}
L12:
	;
	v48 = F__emscripten_memcpy_bulkmem(m, v41+int32(4478088), v31+v41+int32(904), v40)
	mBase = m.M
	goto L14
L14:
	;
	goto L11
L15:
	;
	v53 = v26 + int32(1)
	if v53 != int32(18) {
		v26 = v53
		goto L8
	} else {
		goto L16
	}
L16:
	;
	goto L9
}
func F_pgstat_progress_update_param(m *base.Module, l0 int32, l1 int64) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	v5 = *(*int32)(unsafe.Add(mBase, _consts[465]))
	if v5 == int32(0) {
	} else {
		v9 = int32(*(*uint8)(unsafe.Add(mBase, _consts[466])))
		if v9 != int32(1) {
		} else {
			v12 = int32(4548548)
			v14 = *(*int32)(unsafe.Add(mBase, _consts[17]))
			v15 = int32(1)
			*(*int32)(unsafe.Add(mBase, _consts[17])) = v14 + v15
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
			*(*int32)(unsafe.Add(mBase, uint32(v5))) = v18 + v15
			*(*int64)(unsafe.Add(mBase, uint32(v5+l0<<(uint(int32(3))%32))+232)) = l1
			v26 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
			*(*int32)(unsafe.Add(mBase, uint32(v5))) = v26 + v15
			v32 = *(*int32)(unsafe.Add(mBase, _consts[17]))
			*(*int32)(unsafe.Add(mBase, _consts[17])) = v32 - v15
		}
	}
	return
}
func F_pgstat_report_archiver(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int64
	_ = v16
	var v17 int64
	_ = v17
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int64
	_ = v40
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int64
	_ = v50
	var v52 int64
	_ = v52
	var v54 int64
	_ = v54
	var v56 int64
	_ = v56
	var v58 int64
	_ = v58
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	v7 = *(*int32)(unsafe.Add(mBase, _consts[278]))
	v11 = m.G0
	v12 = int32(16)
	v13 = v11 - v12
	m.G0 = v13
	F___gettimeofday(m, v13)
	mBase = m.M
	v16 = *(*int64)(unsafe.Add(mBase, uint32(v13)))
	v17 = int64(*(*int32)(unsafe.Add(mBase, uint32(v13)+8)))
	m.G0 = v13 + v12
	v26 = int32(4548548)
	v28 = *(*int32)(unsafe.Add(mBase, _consts[17]))
	v29 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[17])) = v28 + v29
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v7)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+40)) = v32 + v29
	if l1 != 0 {
		v38 = int32(112)
	} else {
		v38 = int32(48)
	}
	v39 = v7 + v38
	v40 = *(*int64)(unsafe.Add(mBase, uint32(v39)))
	*(*int64)(unsafe.Add(mBase, uint32(v39))) = v40 + int64(1)
	if l1 != 0 {
		v46 = int32(120)
	} else {
		v46 = int32(56)
	}
	v47 = v7 + v46
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+40)))
	*(*uint8)(unsafe.Add(mBase, uint32(v47)+40)) = uint8(v48)
	v50 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v47)+32)) = v50
	v52 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v47)+24)) = v52
	v54 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v47)+16)) = v54
	v56 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v47)+8)) = v56
	v58 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v47))) = v58
	if l1 != 0 {
		v62 = int32(168)
	} else {
		v62 = int32(104)
	}
	*(*int64)(unsafe.Add(mBase, uint32(v7+v62))) = v17 + v16*int64(1000000) - int64(946684800000000)
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v7)+40))
	v66 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+40)) = v65 + v66
	v69 = int32(4548548)
	v71 = *(*int32)(unsafe.Add(mBase, _consts[17]))
	*(*int32)(unsafe.Add(mBase, _consts[17])) = v71 - v66
	return
}
func F_pgstat_report_subscription_error(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int64
	_ = v10
	var v14 int64
	_ = v14
	v4 = int32(0)
	v7 = F_pgstat_prep_pending_entry(m, int32(5), v4, base.I64_extend_i32_u(l0), v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
		if l1 != 0 {
			v10 = *(*int64)(unsafe.Add(mBase, uint32(v9)))
			*(*int64)(unsafe.Add(mBase, uint32(v9))) = v10 + int64(1)
			return
		} else {
			v14 = *(*int64)(unsafe.Add(mBase, uint32(v9)+8))
			*(*int64)(unsafe.Add(mBase, uint32(v9)+8)) = v14 + int64(1)
			return
		}
	}
}
func F_pgstat_report_tempfile(m *base.Module, l0 int32) {
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
	var v14 int32
	_ = v14
	var v15 int64
	_ = v15
	var v19 int64
	_ = v19
	v4 = int32(*(*uint8)(unsafe.Add(mBase, _consts[929])))
	if v4 == int32(1) {
		v9 = *(*int32)(unsafe.Add(mBase, _consts[130]))
		v12 = F_pgstat_prep_pending_entry(m, int32(1), v9, int64(0), int32(0))
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return
		} else {
			v14 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
			v15 = *(*int64)(unsafe.Add(mBase, uint32(v14)+136))
			*(*int64)(unsafe.Add(mBase, uint32(v14)+136)) = v15 + base.I64_extend_i32_u(l0)
			v19 = *(*int64)(unsafe.Add(mBase, uint32(v14)+128))
			*(*int64)(unsafe.Add(mBase, uint32(v14)+128)) = v19 + int64(1)
			return
		}
	} else {
		return
	}
}
func F_pgstat_snapshot_fixed(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int64
	_ = v9
	var v18 int32
	_ = v18
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	v7 = int32(*(*uint8)(unsafe.Add(mBase, _consts[918])))
	if v7 != 0 {
		v9 = int64(0)
		*(*int64)(unsafe.Add(mBase, _consts[919])) = v9
		*(*int64)(unsafe.Add(mBase, _consts[920])) = v9
		*(*int64)(unsafe.Add(mBase, _consts[921])) = v9
		v18 = int32(0)
		*(*uint8)(unsafe.Add(mBase, _consts[922])) = uint8(v18)
		*(*int32)(unsafe.Add(mBase, _consts[923])) = v18
		*(*int32)(unsafe.Add(mBase, _consts[924])) = v18
		v27 = *(*int32)(unsafe.Add(mBase, _consts[925]))
		if v27 != 0 {
			F_MemoryContextDelete(m, v27)
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, _consts[925])) = int32(0)
				F_pgstat_clear_backend_activity_snapshot(m)
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return
				} else {
					v36 = int32(0)
					*(*uint8)(unsafe.Add(mBase, _consts[918])) = uint8(v36)
					v40 = *(*int32)(unsafe.Add(mBase, _consts[926]))
					if v40 == int32(2) {
						F_pgstat_build_snapshot(m)
						mBase = m.M
						v44 = m.ExcPending
						if v44 != 0 {
							return
						} else {
							return
						}
					} else {
						if base.Ui32(l0-int32(1)) <= base.Ui32(int32(11)) {
							v73 = l0
							v74 = l0*int32(72) + int32(1683808)
							v75 = int32(4477808)
						} else {
							v57 = l0 - int32(24)
							if base.Ui32(int32(8)) < base.Ui32(v57) {
								v72 = int32(0)
							} else {
								v60 = int32(0)
								v62 = *(*int32)(unsafe.Add(mBase, _consts[279]))
								if v62 == v60 {
									v72 = v60
								} else {
									v70 = *(*int32)(unsafe.Add(mBase, uint32(v62+l0<<(uint(int32(2))%32)-int32(96))))
									v72 = v70
								}
							}
							v73 = v57
							v74 = v72
							v75 = int32(4530480)
						}
						v77 = v73 + v75
						if v40 == int32(0) {
							v80 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v77))) = uint8(v80)
							v83 = *(*int32)(unsafe.Add(mBase, uint32(v74)+64))
							m.T0[v83].(func(*base.Module))(m)
							mBase = m.M
							v85 = m.ExcPending
							if v85 != 0 {
								return
							} else {
								v86 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(v77))) = uint8(v86)
								return
							}
						} else {
							v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77))))
							if v82 != 0 {
								return
							} else {
								v83 = *(*int32)(unsafe.Add(mBase, uint32(v74)+64))
								m.T0[v83].(func(*base.Module))(m)
								mBase = m.M
								v85 = m.ExcPending
								if v85 != 0 {
									return
								} else {
									v86 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(v77))) = uint8(v86)
									return
								}
							}
						}
					}
				}
			}
		} else {
			F_pgstat_clear_backend_activity_snapshot(m)
			mBase = m.M
			v34 = m.ExcPending
			if v34 != 0 {
				return
			} else {
				v36 = int32(0)
				*(*uint8)(unsafe.Add(mBase, _consts[918])) = uint8(v36)
				v40 = *(*int32)(unsafe.Add(mBase, _consts[926]))
				if v40 == int32(2) {
					F_pgstat_build_snapshot(m)
					mBase = m.M
					v44 = m.ExcPending
					if v44 != 0 {
						return
					} else {
						return
					}
				} else {
					if base.Ui32(l0-int32(1)) <= base.Ui32(int32(11)) {
						v73 = l0
						v74 = l0*int32(72) + int32(1683808)
						v75 = int32(4477808)
					} else {
						v57 = l0 - int32(24)
						if base.Ui32(int32(8)) < base.Ui32(v57) {
							v72 = int32(0)
						} else {
							v60 = int32(0)
							v62 = *(*int32)(unsafe.Add(mBase, _consts[279]))
							if v62 == v60 {
								v72 = v60
							} else {
								v70 = *(*int32)(unsafe.Add(mBase, uint32(v62+l0<<(uint(int32(2))%32)-int32(96))))
								v72 = v70
							}
						}
						v73 = v57
						v74 = v72
						v75 = int32(4530480)
					}
					v77 = v73 + v75
					if v40 == int32(0) {
						v80 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(v77))) = uint8(v80)
						v83 = *(*int32)(unsafe.Add(mBase, uint32(v74)+64))
						m.T0[v83].(func(*base.Module))(m)
						mBase = m.M
						v85 = m.ExcPending
						if v85 != 0 {
							return
						} else {
							v86 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(v77))) = uint8(v86)
							return
						}
					} else {
						v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77))))
						if v82 != 0 {
							return
						} else {
							v83 = *(*int32)(unsafe.Add(mBase, uint32(v74)+64))
							m.T0[v83].(func(*base.Module))(m)
							mBase = m.M
							v85 = m.ExcPending
							if v85 != 0 {
								return
							} else {
								v86 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(v77))) = uint8(v86)
								return
							}
						}
					}
				}
			}
		}
	} else {
		v40 = *(*int32)(unsafe.Add(mBase, _consts[926]))
		if v40 == int32(2) {
			F_pgstat_build_snapshot(m)
			mBase = m.M
			v44 = m.ExcPending
			if v44 != 0 {
				return
			} else {
				return
			}
		} else {
			if base.Ui32(l0-int32(1)) <= base.Ui32(int32(11)) {
				v73 = l0
				v74 = l0*int32(72) + int32(1683808)
				v75 = int32(4477808)
			} else {
				v57 = l0 - int32(24)
				if base.Ui32(int32(8)) < base.Ui32(v57) {
					v72 = int32(0)
				} else {
					v60 = int32(0)
					v62 = *(*int32)(unsafe.Add(mBase, _consts[279]))
					if v62 == v60 {
						v72 = v60
					} else {
						v70 = *(*int32)(unsafe.Add(mBase, uint32(v62+l0<<(uint(int32(2))%32)-int32(96))))
						v72 = v70
					}
				}
				v73 = v57
				v74 = v72
				v75 = int32(4530480)
			}
			v77 = v73 + v75
			if v40 == int32(0) {
				v80 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v77))) = uint8(v80)
				v83 = *(*int32)(unsafe.Add(mBase, uint32(v74)+64))
				m.T0[v83].(func(*base.Module))(m)
				mBase = m.M
				v85 = m.ExcPending
				if v85 != 0 {
					return
				} else {
					v86 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(v77))) = uint8(v86)
					return
				}
			} else {
				v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77))))
				if v82 != 0 {
					return
				} else {
					v83 = *(*int32)(unsafe.Add(mBase, uint32(v74)+64))
					m.T0[v83].(func(*base.Module))(m)
					mBase = m.M
					v85 = m.ExcPending
					if v85 != 0 {
						return
					} else {
						v86 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(v77))) = uint8(v86)
						return
					}
				}
			}
		}
	}
}
func F_pgstat_twophase_postcommit(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int64
	_ = v21
	var v22 int64
	_ = v22
	var v25 int64
	_ = v25
	var v26 int64
	_ = v26
	var v29 int64
	_ = v29
	var v30 int64
	_ = v30
	var v33 int32
	_ = v33
	var v37 int64
	_ = v37
	var v38 int64
	_ = v38
	var v40 int32
	_ = v40
	var v41 int64
	_ = v41
	var v47 int64
	_ = v47
	var v48 int64
	_ = v48
	var v49 int64
	_ = v49
	var v50 int64
	_ = v50
	var v54 int64
	_ = v54
	var v55 int64
	_ = v55
	var v59 int64
	_ = v59
	var v60 int64
	_ = v60
	var v61 int64
	_ = v61
	var v62 int64
	_ = v62
	v10 = *(*int32)(unsafe.Add(mBase, _consts[130]))
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+52)))
	if v11 != 0 {
		v12 = int32(0)
	} else {
		v12 = v10
	}
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l2)+48))
	v16 = F_pgstat_prep_pending_entry(m, int32(2), v12, base.I64_extend_i32_u(v13), int32(0))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return
	} else {
		v18 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
		*(*uint8)(unsafe.Add(mBase, uint32(v18)+4)) = uint8(v11)
		*(*int32)(unsafe.Add(mBase, uint32(v18))) = v13
		v21 = *(*int64)(unsafe.Add(mBase, uint32(v18)+40))
		v22 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
		*(*int64)(unsafe.Add(mBase, uint32(v18)+40)) = v21 + v22
		v25 = *(*int64)(unsafe.Add(mBase, uint32(v18)+48))
		v26 = *(*int64)(unsafe.Add(mBase, uint32(l2)+8))
		*(*int64)(unsafe.Add(mBase, uint32(v18)+48)) = v25 + v26
		v29 = *(*int64)(unsafe.Add(mBase, uint32(v18)+56))
		v30 = *(*int64)(unsafe.Add(mBase, uint32(l2)+16))
		*(*int64)(unsafe.Add(mBase, uint32(v18)+56)) = v29 + v30
		v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+53)))
		*(*uint8)(unsafe.Add(mBase, uint32(v18)+80)) = uint8(v33)
		if v33 == int32(0) {
			v37 = *(*int64)(unsafe.Add(mBase, uint32(v18)+88))
			v38 = *(*int64)(unsafe.Add(mBase, uint32(v18)+96))
			v47 = v37
			v48 = v38
		} else {
			v40 = v18 + int32(88)
			v41 = int64(0)
			*(*int64)(unsafe.Add(mBase, uint32(v40))) = v41
			*(*int64)(unsafe.Add(mBase, uint32(v40)+8)) = v41
			v47 = int64(0)
			v48 = v41
		}
		v49 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
		v50 = *(*int64)(unsafe.Add(mBase, uint32(l2)+16))
		*(*int64)(unsafe.Add(mBase, uint32(v18)+88)) = v49 - v50 + v47
		v54 = *(*int64)(unsafe.Add(mBase, uint32(l2)+16))
		v55 = *(*int64)(unsafe.Add(mBase, uint32(l2)+8))
		*(*int64)(unsafe.Add(mBase, uint32(v18)+96)) = v54 + v55 + v48
		v59 = *(*int64)(unsafe.Add(mBase, uint32(v18)+104))
		v60 = *(*int64)(unsafe.Add(mBase, uint32(l2)+16))
		v61 = *(*int64)(unsafe.Add(mBase, uint32(l2)+8))
		v62 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
		*(*int64)(unsafe.Add(mBase, uint32(v18)+104)) = v59 + (v60 + (v61 + v62))
		return
	}
}
func F_pgstat_wal_snapshot_cb(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v13 int64
	_ = v13
	var v18 int64
	_ = v18
	var v23 int64
	_ = v23
	var v28 int64
	_ = v28
	var v31 int64
	_ = v31
	var v34 int32
	_ = v34
	v4 = *(*int32)(unsafe.Add(mBase, _consts[278]))
	v6 = v4 + int32(53272)
	v8 = F_LWLockAcquire(m, v6, int32(1))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return
	} else {
		v13 = *(*int64)(unsafe.Add(mBase, uint32(v4)+uint32(_consts[936])))
		*(*int64)(unsafe.Add(mBase, _consts[937])) = v13
		v18 = *(*int64)(unsafe.Add(mBase, uint32(v4)+uint32(_consts[938])))
		*(*int64)(unsafe.Add(mBase, _consts[939])) = v18
		v23 = *(*int64)(unsafe.Add(mBase, uint32(v4)+uint32(_consts[940])))
		*(*int64)(unsafe.Add(mBase, _consts[941])) = v23
		v28 = *(*int64)(unsafe.Add(mBase, uint32(v4)+uint32(_consts[942])))
		*(*int64)(unsafe.Add(mBase, _consts[943])) = v28
		v31 = *(*int64)(unsafe.Add(mBase, uint32(v4)+uint32(_consts[944])))
		*(*int64)(unsafe.Add(mBase, _consts[945])) = v31
		F_LWLockRelease(m, v6)
		mBase = m.M
		v34 = m.ExcPending
		if v34 != 0 {
			return
		} else {
			return
		}
	}
}
