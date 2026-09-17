package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_CMPTRGM_UNSIGNED(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	v5 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v5 != v6 {
		v21 = v5
		v22 = v6
		if base.Ui32(v21) < base.Ui32(v22) {
			v26 = int32(-1)
		} else {
			v26 = int32(1)
		}
		return v26
	} else {
		v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
		v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
		if v8 != v9 {
			v21 = v8
			v22 = v9
			if base.Ui32(v21) < base.Ui32(v22) {
				v26 = int32(-1)
			} else {
				v26 = int32(1)
			}
			return v26
		} else {
			v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+2)))
			v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+2)))
			if v11 != v12 {
				if base.Ui32(v11) < base.Ui32(v12) {
					v17 = int32(-1)
				} else {
					v17 = int32(1)
				}
				v19 = v17
			} else {
				v19 = int32(0)
			}
			return v19
		}
	}
}
func F_CheckCmdReplicaIdentity(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int64
	_ = v69
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v152 int32
	_ = v152
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v265 int32
	_ = v265
	var v279 int32
	_ = v279
	var v295 int32
	_ = v295
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v327 int32
	_ = v327
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v341 int32
	_ = v341
	var v344 int32
	_ = v344
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v378 int32
	_ = v378
	var v381 int32
	_ = v381
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v420 int32
	_ = v420
	var v423 int32
	_ = v423
	var v425 int32
	_ = v425
	var v428 int32
	_ = v428
	var v439 int32
	_ = v439
	var v442 int32
	_ = v442
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v449 int32
	_ = v449
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v454 int32
	_ = v454
	var v456 int32
	_ = v456
	var v460 int32
	_ = v460
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v480 int32
	_ = v480
	var v482 int32
	_ = v482
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v497 int32
	_ = v497
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v509 int32
	_ = v509
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v529 int32
	_ = v529
	var v535 int32
	_ = v535
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v556 int32
	_ = v556
	var v571 int32
	_ = v571
	var v614 int32
	_ = v614
	var v645 int32
	_ = v645
	var v646 int32
	_ = v646
	var v649 int32
	_ = v649
	var v653 int32
	_ = v653
	var v656 int32
	_ = v656
	var v659 int32
	_ = v659
	var v662 int32
	_ = v662
	var v664 int32
	_ = v664
	var v665 int32
	_ = v665
	var v667 int32
	_ = v667
	var v670 int32
	_ = v670
	var v673 int32
	_ = v673
	var v675 int32
	_ = v675
	var v676 int32
	_ = v676
	var v679 int32
	_ = v679
	var v683 int32
	_ = v683
	var v684 int32
	_ = v684
	var v694 int32
	_ = v694
	var v695 int32
	_ = v695
	var v698 int32
	_ = v698
	var v702 int32
	_ = v702
	var v705 int32
	_ = v705
	var v707 int32
	_ = v707
	var v710 int32
	_ = v710
	var v717 int32
	_ = v717
	var v719 int32
	_ = v719
	var v727 int32
	_ = v727
	var v728 int32
	_ = v728
	var v741 int32
	_ = v741
	var v747 int32
	_ = v747
	var v773 int32
	_ = v773
	var v774 int32
	_ = v774
	var v777 int32
	_ = v777
	var v784 int32
	_ = v784
	var v785 int32
	_ = v785
	var v795 int32
	_ = v795
	var v798 int32
	_ = v798
	var v799 int32
	_ = v799
	var v800 int32
	_ = v800
	var v801 int32
	_ = v801
	var v802 int32
	_ = v802
	var v803 int32
	_ = v803
	var v804 int32
	_ = v804
	var v805 int32
	_ = v805
	var v806 int32
	_ = v806
	var v807 int32
	_ = v807
	var v808 int32
	_ = v808
	var v810 int32
	_ = v810
	var v816 int32
	_ = v816
	var v825 int32
	_ = v825
	var v827 int32
	_ = v827
	var v828 int32
	_ = v828
	var v831 int32
	_ = v831
	var v835 int32
	_ = v835
	var v838 int32
	_ = v838
	var v840 int32
	_ = v840
	var v843 int32
	_ = v843
	var v850 int32
	_ = v850
	var v852 int32
	_ = v852
	var v860 int32
	_ = v860
	var v861 int32
	_ = v861
	var v874 int32
	_ = v874
	var v904 int32
	_ = v904
	var v906 int32
	_ = v906
	var v908 int32
	_ = v908
	var v910 int32
	_ = v910
	var v911 int32
	_ = v911
	var v914 int32
	_ = v914
	var v946 int32
	_ = v946
	var v949 int32
	_ = v949
	var v950 int32
	_ = v950
	var v951 int32
	_ = v951
	var v953 int32
	_ = v953
	var v955 int32
	_ = v955
	var v957 int32
	_ = v957
	var v960 int32
	_ = v960
	var v961 int32
	_ = v961
	var v962 int32
	_ = v962
	var v964 int32
	_ = v964
	var v966 int32
	_ = v966
	var v996 int32
	_ = v996
	var v997 int32
	_ = v997
	var v1000 int32
	_ = v1000
	var v1003 int32
	_ = v1003
	var v1006 int32
	_ = v1006
	var v1009 int32
	_ = v1009
	var v1010 int32
	_ = v1010
	var v1013 int32
	_ = v1013
	var v1016 int32
	_ = v1016
	var v1019 int32
	_ = v1019
	var v1020 int32
	_ = v1020
	var v1023 int32
	_ = v1023
	var v1026 int32
	_ = v1026
	var v1029 int32
	_ = v1029
	var v1030 int32
	_ = v1030
	var v1034 int32
	_ = v1034
	var v1035 int32
	_ = v1035
	var v1064 int32
	_ = v1064
	var v1066 int32
	_ = v1066
	var v1069 int32
	_ = v1069
	var v1070 int32
	_ = v1070
	var v1073 int32
	_ = v1073
	var v1076 int32
	_ = v1076
	var v1077 int32
	_ = v1077
	var v1079 int32
	_ = v1079
	var v1081 int64
	_ = v1081
	var v1118 int32
	_ = v1118
	var v1122 int32
	_ = v1122
	var v1127 int32
	_ = v1127
	var v1129 int32
	_ = v1129
	var v1132 int32
	_ = v1132
	var v1139 int32
	_ = v1139
	var v1146 int32
	_ = v1146
	var v1152 int32
	_ = v1152
	var v1155 int32
	_ = v1155
	var v1162 int32
	_ = v1162
	var v1169 int32
	_ = v1169
	var v1174 int32
	_ = v1174
	var v1175 int32
	_ = v1175
	var v1176 int32
	_ = v1176
	var v1177 int32
	_ = v1177
	var v1182 int32
	_ = v1182
	var v1187 int32
	_ = v1187
	var v1223 int32
	_ = v1223
	var v1226 int32
	_ = v1226
	var v1227 int32
	_ = v1227
	var v1235 int32
	_ = v1235
	var v1239 int32
	_ = v1239
	var v1244 int32
	_ = v1244
	var v1248 int32
	_ = v1248
	var v1251 int32
	_ = v1251
	var v1252 int32
	_ = v1252
	var v1260 int32
	_ = v1260
	var v1264 int32
	_ = v1264
	var v1269 int32
	_ = v1269
	var v1273 int32
	_ = v1273
	var v1276 int32
	_ = v1276
	var v1277 int32
	_ = v1277
	var v1285 int32
	_ = v1285
	var v1289 int32
	_ = v1289
	var v1294 int32
	_ = v1294
	var v1298 int32
	_ = v1298
	var v1301 int32
	_ = v1301
	var v1302 int32
	_ = v1302
	var v1310 int32
	_ = v1310
	var v1314 int32
	_ = v1314
	var v1319 int32
	_ = v1319
	var v1323 int32
	_ = v1323
	var v1326 int32
	_ = v1326
	var v1327 int32
	_ = v1327
	var v1335 int32
	_ = v1335
	var v1339 int32
	_ = v1339
	var v1344 int32
	_ = v1344
	var v1348 int32
	_ = v1348
	var v1351 int32
	_ = v1351
	var v1352 int32
	_ = v1352
	var v1360 int32
	_ = v1360
	var v1364 int32
	_ = v1364
	var v1369 int32
	_ = v1369
	var v1373 int32
	_ = v1373
	var v1376 int32
	_ = v1376
	var v1377 int32
	_ = v1377
	var v1383 int32
	_ = v1383
	var v1387 int32
	_ = v1387
	var v1392 int32
	_ = v1392
	var v1396 int32
	_ = v1396
	var v1399 int32
	_ = v1399
	var v1400 int32
	_ = v1400
	var v1408 int32
	_ = v1408
	var v1412 int32
	_ = v1412
	var v1417 int32
	_ = v1417
	v3 = int32(0)
	v28 = m.G0
	v30 = v28 - int32(128)
	m.G0 = v30
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+119)))
	if v33 == int32(112) {
		goto L9
	} else {
		goto L10
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1396 = m.ExcPending
	if v1396 != 0 {
		goto L25
	} else {
		goto L294
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1373 = m.ExcPending
	if v1373 != 0 {
		goto L25
	} else {
		goto L289
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1348 = m.ExcPending
	if v1348 != 0 {
		goto L25
	} else {
		goto L284
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1323 = m.ExcPending
	if v1323 != 0 {
		goto L25
	} else {
		goto L279
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1298 = m.ExcPending
	if v1298 != 0 {
		goto L25
	} else {
		goto L274
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1273 = m.ExcPending
	if v1273 != 0 {
		goto L25
	} else {
		goto L269
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1248 = m.ExcPending
	if v1248 != 0 {
		goto L25
	} else {
		goto L264
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1223 = m.ExcPending
	if v1223 != 0 {
		goto L25
	} else {
		goto L259
	}
L9:
	;
	m.G0 = v30 + int32(128)
	return
L10:
	;
	switch l1 - int32(2) {
	case 0, 2:
		goto L11
	default:
		goto L9
	}
L11:
	;
	v39 = v30 + int32(118)
	v40 = m.G0
	v42 = v40 - int32(16)
	m.G0 = v42
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45)+119)))
	switch v46 - int32(112) {
	case 0, 2:
		goto L13
	default:
		v59 = v3
		goto L12
	}
L12:
	;
	if v59 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L13:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	goto L14
L14:
	;
	if base.Ui32(v49) < base.Ui32(int32(_a_F_CheckCmdReplicaIdentity_0)) {
		v59 = v3
		goto L12
	} else {
		goto L15
	}
L15:
	;
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45)+118)))
	v59 = base.B2i32(v52 == int32(112)) & base.B2i32(base.Ui32(int32(_a_F_CheckCmdReplicaIdentity_1)) < base.Ui32(v49))
	goto L12
L16:
	;
	v1129 = base.B2i32(l1 != int32(2))
	if v1129 == int32(0) {
		goto L226
	} else {
		goto L227
	}
L17:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1118 = m.ExcPending
	if v1118 != 0 {
		goto L25
	} else {
		goto L223
	}
L18:
	;
	m.G0 = v42 + int32(16)
	goto L16
L19:
	;
	v62 = int32(257)
	*(*uint16)(unsafe.Add(mBase, uint32(v39)+8)) = uint16(v62)
	*(*int64)(unsafe.Add(mBase, uint32(v39))) = int64(72340172821233664)
	goto L18
L20:
	;
	goto L21
L21:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+176))
	if v66 != 0 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v67 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v66)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(v39)+8)) = uint16(v67)
	v69 = *(*int64)(unsafe.Add(mBase, uint32(v66)))
	*(*int64)(unsafe.Add(mBase, uint32(v39))) = v69
	goto L18
L23:
	;
	goto L24
L24:
	;
	v71 = int32(257)
	*(*uint16)(unsafe.Add(mBase, uint32(v39)+8)) = uint16(v71)
	*(*int64)(unsafe.Add(mBase, uint32(v39))) = int64(72340172821233664)
	v75 = F_GetRelationPublications(m, v44)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	return
L26:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)+68))
	v79 = F_GetSchemaPublications(m, v78)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L25
	} else {
		goto L27
	}
L27:
	;
	v81 = F_list_concat_unique_oid(m, v75, v79)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L25
	} else {
		goto L28
	}
L28:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83)+131)))
	if v84 != int32(1) {
		v179 = v81
		v182 = v3
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v194 = m.G0
	v196 = v194 - int32(48)
	m.G0 = v196
	v200 = F_table_open(m, int32(_a_F_CheckCmdReplicaIdentity_2), int32(1))
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L25
	} else {
		goto L44
	}
L30:
	;
	v87 = F_get_partition_ancestors(m, v44)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L25
	} else {
		goto L31
	}
L31:
	;
	if v87 == int32(0) {
		v179 = v81
		v182 = v3
		goto L29
	} else {
		goto L32
	}
L32:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v87)+4))
	if int32(0) < v91 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v106 = v81
	v109 = v3
	goto L36
L34:
	;
	v152 = v81
	goto L35
L35:
	;
	v179 = v152
	v182 = v87
	goto L29
L36:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v87)+12))
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v121+v109<<(uint(int32(2))%32))))
	v126 = F_GetRelationPublications(m, v125)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L25
	} else {
		goto L38
	}
L37:
	;
	v152 = v134
	goto L35
L38:
	;
	v128 = F_list_concat_unique_oid(m, v106, v126)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L25
	} else {
		goto L39
	}
L39:
	;
	v130 = F_get_rel_namespace(m, v125)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L25
	} else {
		goto L40
	}
L40:
	;
	v132 = F_GetSchemaPublications(m, v130)
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L25
	} else {
		goto L41
	}
L41:
	;
	v134 = F_list_concat_unique_oid(m, v128, v132)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L25
	} else {
		goto L42
	}
L42:
	;
	v137 = v109 + int32(1)
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v87)+4))
	if v137 < v138 {
		v106 = v134
		v109 = v137
		goto L36
	} else {
		goto L43
	}
L43:
	;
	goto L37
L44:
	;
	F_ScanKeyInit(m, v196, int32(4), int32(3), int32(60), int32(1))
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L25
	} else {
		goto L45
	}
L45:
	;
	v208 = int32(0)
	v212 = F_systable_beginscan(m, v200, v208, v208, v208, int32(1), v196)
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L25
	} else {
		goto L46
	}
L46:
	;
	v216 = v3
	goto L47
L47:
	;
	v241 = F_systable_getnext(m, v212)
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L25
	} else {
		goto L49
	}
L48:
	;
	F_systable_endscan(m, v212)
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L25
	} else {
		goto L54
	}
L49:
	;
	if v241 != 0 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v241)+16))
	v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v243)+22)))
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v243+v244)))
	v247 = F_lappend_oid(m, v216, v246)
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L25
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	goto L48
L53:
	;
	v216 = v247
	goto L47
L54:
	;
	F_relation_close(m, v200, int32(1))
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L25
	} else {
		goto L55
	}
L55:
	;
	m.G0 = v196 + int32(48)
	v257 = F_list_concat_unique_oid(m, v179, v216)
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L25
	} else {
		goto L57
	}
L56:
	;
	v1064 = *(*int32)(unsafe.Add(mBase, uint32(l0)+176))
	if v1064 != 0 {
		goto L218
	} else {
		goto L219
	}
L57:
	;
	if v257 == int32(0) {
		goto L56
	} else {
		goto L58
	}
L58:
	;
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v257)+4))
	if v261 <= int32(0) {
		goto L56
	} else {
		goto L59
	}
L59:
	;
	v265 = v30 + int32(122)
	v279 = int32(0)
	goto L60
L60:
	;
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v257)+12))
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v295+v279<<(uint(int32(2))%32))))
	v300 = F_SearchSysCache1(m, int32(51), v299)
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L25
	} else {
		goto L62
	}
L61:
	;
	goto L56
L62:
	;
	if v300 == int32(0) {
		goto L17
	} else {
		goto L63
	}
L63:
	;
	v304 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39))))
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v300)+16))
	v306 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v305)+22)))
	v307 = v305 + v306
	v308 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v307)+73)))
	v309 = v304 | v308
	*(*uint8)(unsafe.Add(mBase, uint32(v39))) = uint8(v309)
	v311 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+1)))
	v312 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v307)+74)))
	v313 = v311 | v312
	*(*uint8)(unsafe.Add(mBase, uint32(v39)+1)) = uint8(v313)
	v315 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+2)))
	v316 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v307)+75)))
	v317 = v315 | v316
	*(*uint8)(unsafe.Add(mBase, uint32(v39)+2)) = uint8(v317)
	v319 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+3)))
	v320 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v307)+76)))
	v321 = v319 | v320
	*(*uint8)(unsafe.Add(mBase, uint32(v39)+3)) = uint8(v321)
	v323 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v307)+72)))
	if v323 != 0 {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v439 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v307)+74)))
	if v439 == int32(0) {
		goto L106
	} else {
		goto L107
	}
L65:
	;
	v324 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v307)+74)))
	if v324 == int32(0) {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v327 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v307)+75)))
	if v327 != int32(1) {
		goto L64
	} else {
		goto L69
	}
L67:
	;
	goto L68
L68:
	;
	v330 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v307)+77)))
	v331 = int32(0)
	v332 = m.G0
	v334 = v332 - int32(32)
	m.G0 = v334
	v336 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v337 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v336)+130)))
	if v337 == int32(102) {
		v408 = v331
		goto L70
	} else {
		goto L71
	}
L69:
	;
	goto L68
L70:
	;
	m.G0 = v334 + int32(32)
	if v408 == int32(0) {
		goto L64
	} else {
		goto L100
	}
L71:
	;
	v341 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	if v330 == int32(0) {
		v351 = v341
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v352 = F_SearchSysCache2(m, int32(53), v351, v299)
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L25
	} else {
		goto L79
	}
L73:
	;
	v344 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v336)+131)))
	if v344 != int32(1) {
		v351 = v341
		goto L72
	} else {
		goto L74
	}
L74:
	;
	v347 = F_GetTopMostAncestorInPublication(m, v299, v182)
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L25
	} else {
		goto L75
	}
L75:
	;
	if v347 != 0 {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v349 = v347
	goto L78
L77:
	;
	v349 = v341
	goto L78
L78:
	;
	v351 = v349
	goto L72
L79:
	;
	if v352 == int32(0) {
		v408 = v331
		goto L70
	} else {
		goto L80
	}
L80:
	;
	v360 = F_SysCacheGetAttr(m, int32(53), v352, int32(4), v334+int32(31))
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L25
	} else {
		goto L81
	}
L81:
	;
	v362 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v334)+31)))
	if v362 != 0 {
		v402 = v331
		goto L82
	} else {
		goto L83
	}
L82:
	;
	F_ReleaseCatCache(m, v352)
	mBase = m.M
	v407 = m.ExcPending
	if v407 != 0 {
		goto L25
	} else {
		goto L99
	}
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v334)+16)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v334)+24)) = v351
	*(*int32)(unsafe.Add(mBase, uint32(v334)+20)) = v341
	*(*uint8)(unsafe.Add(mBase, uint32(v334)+16)) = uint8(v330)
	v369 = F_RelationGetIndexAttrBitmap(m, l0, int32(2))
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
		goto L25
	} else {
		goto L84
	}
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v334)+12)) = v369
	v372 = F_text_to_cstring(m, v360)
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L25
	} else {
		goto L85
	}
L85:
	;
	v374 = F_stringToNode(m, v372)
	mBase = m.M
	v375 = m.ExcPending
	if v375 != 0 {
		goto L25
	} else {
		goto L86
	}
L86:
	;
	if v374 == int32(0) {
		v402 = v331
		goto L82
	} else {
		goto L87
	}
L87:
	;
	v378 = *(*int32)(unsafe.Add(mBase, uint32(v374)))
	if v378 == int32(6) {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v381 = int32(*(*int16)(unsafe.Add(mBase, uint32(v374)+8)))
	if v330 != 0 {
		goto L91
	} else {
		goto L92
	}
L89:
	;
	goto L90
L90:
	;
	v400 = F_expression_tree_walker_impl(m, v374, int32(567), v334+int32(12))
	mBase = m.M
	v401 = m.ExcPending
	if v401 != 0 {
		goto L25
	} else {
		goto L98
	}
L91:
	;
	v383 = F_get_attname(m, v351, v381, int32(0))
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L25
	} else {
		goto L94
	}
L92:
	;
	v387 = v381
	goto L93
L93:
	;
	v391 = F_bms_is_member(m, v387+int32(7), v369)
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L25
	} else {
		goto L96
	}
L94:
	;
	v385 = F_get_attnum(m, v341, v383)
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L25
	} else {
		goto L95
	}
L95:
	;
	v387 = v385
	goto L93
L96:
	;
	if v391 == int32(0) {
		v402 = int32(1)
		goto L82
	} else {
		goto L97
	}
L97:
	;
	goto L90
L98:
	;
	v402 = v400
	goto L82
L99:
	;
	v408 = v402
	goto L70
L100:
	;
	v420 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v307)+74)))
	if v420 == int32(1) {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	v423 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v265))) = uint8(v423)
	goto L103
L102:
	;
	goto L103
L103:
	;
	v425 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v307)+75)))
	if v425 != int32(1) {
		goto L64
	} else {
		goto L104
	}
L104:
	;
	v428 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v39)+5)) = uint8(v428)
	goto L64
L105:
	;
	F_ReleaseCatCache(m, v300)
	mBase = m.M
	v996 = m.ExcPending
	if v996 != 0 {
		goto L25
	} else {
		goto L199
	}
L106:
	;
	v442 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v307)+75)))
	if v442 != int32(1) {
		goto L105
	} else {
		goto L109
	}
L107:
	;
	goto L108
L108:
	;
	v445 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v307)+77)))
	v446 = int32(*(*int8)(unsafe.Add(mBase, uint32(v307)+78)))
	v447 = m.G0
	v449 = v447 - int32(16)
	m.G0 = v449
	v451 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v452 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v449)+12)) = v452
	v454 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v456 = v42 + int32(15)
	*(*uint8)(unsafe.Add(mBase, uint32(v456))) = uint8(v452)
	v460 = v42 + int32(14)
	*(*uint8)(unsafe.Add(mBase, uint32(v460))) = uint8(v452)
	if v445 == v452 {
		v473 = v451
		goto L110
	} else {
		goto L111
	}
L109:
	;
	goto L108
L110:
	;
	v474 = F_GetPublication(m, v299)
	mBase = m.M
	v475 = m.ExcPending
	if v475 != 0 {
		goto L25
	} else {
		goto L117
	}
L111:
	;
	v465 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v466 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v465)+131)))
	if v466 != int32(1) {
		v473 = v451
		goto L110
	} else {
		goto L112
	}
L112:
	;
	v469 = F_GetTopMostAncestorInPublication(m, v299, v182)
	mBase = m.M
	v470 = m.ExcPending
	if v470 != 0 {
		goto L25
	} else {
		goto L113
	}
L113:
	;
	if v469 != 0 {
		goto L114
	} else {
		goto L115
	}
L114:
	;
	v471 = v469
	goto L116
L115:
	;
	v471 = v451
	goto L116
L116:
	;
	v473 = v471
	goto L110
L117:
	;
	v477 = v449 + int32(12)
	v478 = m.G0
	v480 = v478 - int32(16)
	m.G0 = v480
	v482 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v474)+8)))
	if v482 != 0 {
		goto L118
	} else {
		goto L119
	}
L118:
	;
	m.G0 = v480 + int32(16)
	v645 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v646 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v645)+130)))
	if v646 != int32(102) {
		goto L139
	} else {
		goto L140
	}
L119:
	;
	v484 = *(*int32)(unsafe.Add(mBase, uint32(v474)))
	v485 = F_SearchSysCache2(m, int32(53), v473, v484)
	mBase = m.M
	v486 = m.ExcPending
	if v486 != 0 {
		goto L25
	} else {
		goto L120
	}
L120:
	;
	if v485 == int32(0) {
		goto L118
	} else {
		goto L121
	}
L121:
	;
	v493 = F_SysCacheGetAttr(m, int32(53), v485, int32(5), v480+int32(15))
	mBase = m.M
	v494 = m.ExcPending
	if v494 != 0 {
		goto L25
	} else {
		goto L122
	}
L122:
	;
	v495 = int32(0)
	v497 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v480)+15)))
	if base.B2i32(v477 == v495)|v497&int32(1) == v495 {
		goto L123
	} else {
		goto L124
	}
L123:
	;
	v503 = *(*int32)(unsafe.Add(mBase, uint32(v477)))
	v504 = F_pg_detoast_datum(m, v493)
	mBase = m.M
	v505 = m.ExcPending
	if v505 != 0 {
		goto L25
	} else {
		goto L126
	}
L124:
	;
	goto L125
L125:
	;
	F_ReleaseCatCache(m, v485)
	mBase = m.M
	v614 = m.ExcPending
	if v614 != 0 {
		goto L25
	} else {
		goto L137
	}
L126:
	;
	v506 = *(*int32)(unsafe.Add(mBase, uint32(v504)+8))
	if v506 == int32(0) {
		goto L127
	} else {
		goto L128
	}
L127:
	;
	v509 = *(*int32)(unsafe.Add(mBase, uint32(v504)+4))
	v516 = (v509<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L129
L128:
	;
	v516 = v506
	goto L129
L129:
	;
	v517 = int32(0)
	v518 = *(*int32)(unsafe.Add(mBase, uint32(v504)+16))
	if v517 < v518 {
		goto L130
	} else {
		goto L131
	}
L130:
	;
	v529 = v517
	v535 = v503
	goto L133
L131:
	;
	v571 = v503
	goto L132
L132:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v477))) = v571
	goto L125
L133:
	;
	v552 = int32(*(*int16)(unsafe.Add(mBase, uint32(v504+v516+v529<<(uint(int32(1))%32)))))
	v553 = F_bms_add_member(m, v535, v552)
	mBase = m.M
	v554 = m.ExcPending
	if v554 != 0 {
		goto L25
	} else {
		goto L135
	}
L134:
	;
	v571 = v553
	goto L132
L135:
	;
	v556 = v529 + int32(1)
	if v556 != v518 {
		v529 = v556
		v535 = v553
		goto L133
	} else {
		goto L136
	}
L136:
	;
	goto L134
L137:
	;
	goto L118
L138:
	;
	m.G0 = v449 + int32(16)
	if v914&int32(1) == int32(0) {
		goto L105
	} else {
		goto L194
	}
L139:
	;
	v683 = F_RelationGetIndexAttrBitmap(m, l0, int32(2))
	mBase = m.M
	v684 = m.ExcPending
	if v684 != 0 {
		goto L25
	} else {
		goto L151
	}
L140:
	;
	v649 = *(*int32)(unsafe.Add(mBase, uint32(v449)+12))
	*(*uint8)(unsafe.Add(mBase, uint32(v456))) = uint8(base.B2i32(v649 != int32(0)))
	v653 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	if v446 == int32(115) {
		v665 = v653
		goto L141
	} else {
		goto L142
	}
L141:
	;
	v667 = *(*int32)(unsafe.Add(mBase, uint32(v665)+16))
	if v667 == int32(0) {
		goto L145
	} else {
		goto L146
	}
L142:
	;
	v656 = *(*int32)(unsafe.Add(mBase, uint32(v653)+16))
	if v656 == int32(0) {
		v665 = v653
		goto L141
	} else {
		goto L143
	}
L143:
	;
	v659 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v656)+17)))
	if v659 != int32(1) {
		v665 = v653
		goto L141
	} else {
		goto L144
	}
L144:
	;
	v662 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v460))) = uint8(v662)
	v664 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v665 = v664
	goto L141
L145:
	;
	v675 = int32(1)
	v676 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v460))))
	if v676 != v675 {
		goto L139
	} else {
		goto L148
	}
L146:
	;
	v670 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v667)+18)))
	if v670 != int32(1) {
		goto L145
	} else {
		goto L147
	}
L147:
	;
	v673 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v460))) = uint8(v673)
	goto L145
L148:
	;
	v679 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v456))))
	if v679 != 0 {
		v914 = v675
		goto L138
	} else {
		goto L149
	}
L149:
	;
	goto L139
L150:
	;
	v904 = *(*int32)(unsafe.Add(mBase, uint32(v449)+12))
	F_bms_free(m, v904)
	mBase = m.M
	v906 = m.ExcPending
	if v906 != 0 {
		goto L25
	} else {
		goto L191
	}
L151:
	;
	if v683 == int32(0) {
		goto L154
	} else {
		goto L155
	}
L152:
	;
	if v741 < int32(0) {
		goto L150
	} else {
		goto L163
	}
L153:
	;
	v741 = base.I32_ctz(v727) | v728<<(uint(int32(5))%32)
	goto L152
L154:
	;
	v741 = int32(-2)
	goto L152
L155:
	;
	v694 = base.I32_div_s(int32(0), int32(32))
	v695 = *(*int32)(unsafe.Add(mBase, uint32(v683)+4))
	if v695 <= v694 {
		goto L154
	} else {
		goto L156
	}
L156:
	;
	v698 = v683 + int32(8)
	v702 = *(*int32)(unsafe.Add(mBase, uint32(v698+v694<<(uint(int32(2))%32))))
	v705 = v702 & int32(-1)
	if v705 != 0 {
		v727 = v705
		v728 = v694
		goto L153
	} else {
		goto L157
	}
L157:
	;
	v707 = v694 + int32(1)
	if v707 == v695 {
		goto L154
	} else {
		goto L158
	}
L158:
	;
	v710 = v707
	goto L159
L159:
	;
	v717 = *(*int32)(unsafe.Add(mBase, uint32(v698+v710<<(uint(int32(2))%32))))
	if v717 != 0 {
		v727 = v717
		v728 = v710
		goto L153
	} else {
		goto L161
	}
L160:
	;
	goto L154
L161:
	;
	v719 = v710 + int32(1)
	if v719 != v695 {
		v710 = v719
		goto L159
	} else {
		goto L162
	}
L162:
	;
	goto L160
L163:
	;
	v747 = v741
	goto L164
L164:
	;
	v773 = base.I32_extend16_s(v747 - int32(7))
	v774 = *(*int32)(unsafe.Add(mBase, uint32(v449)+12))
	if v774 == int32(0) {
		goto L167
	} else {
		goto L168
	}
L165:
	;
	goto L150
L166:
	;
	if v683 == int32(0) {
		goto L181
	} else {
		goto L182
	}
L167:
	;
	v777 = *(*int32)(unsafe.Add(mBase, uint32(v454)))
	v784 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v454+v777<<(uint(int32(4))%32)+v773*int32(100))+10)))
	v785 = int32(115)
	if base.B2i32(base.B2i32(v784 == v785)&base.B2i32(v446 != v785) == int32(0))&base.B2i32(v784 != int32(118)) != 0 {
		goto L166
	} else {
		goto L170
	}
L168:
	;
	goto L169
L169:
	;
	if v445 != 0 {
		goto L171
	} else {
		goto L172
	}
L170:
	;
	v795 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v460))) = uint8(v795)
	goto L150
L171:
	;
	v798 = F_get_attname(m, v451, v773, int32(0))
	mBase = m.M
	v799 = m.ExcPending
	if v799 != 0 {
		goto L25
	} else {
		goto L174
	}
L172:
	;
	v803 = v773
	v804 = v774
	goto L173
L173:
	;
	v805 = F_bms_is_member(m, v803, v804)
	mBase = m.M
	v806 = m.ExcPending
	if v806 != 0 {
		goto L25
	} else {
		goto L176
	}
L174:
	;
	v800 = F_get_attnum(m, v473, v798)
	mBase = m.M
	v801 = m.ExcPending
	if v801 != 0 {
		goto L25
	} else {
		goto L175
	}
L175:
	;
	v802 = *(*int32)(unsafe.Add(mBase, uint32(v449)+12))
	v803 = v800
	v804 = v802
	goto L173
L176:
	;
	v807 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v456))))
	v808 = int32(1)
	v810 = v807 | (v805 ^ v808)
	*(*uint8)(unsafe.Add(mBase, uint32(v456))) = uint8(v810)
	if v810&v808 == int32(0) {
		goto L166
	} else {
		goto L177
	}
L177:
	;
	v816 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v460))))
	if v816 != 0 {
		goto L150
	} else {
		goto L178
	}
L178:
	;
	goto L166
L179:
	;
	if int32(0) <= v874 {
		v747 = v874
		goto L164
	} else {
		goto L190
	}
L180:
	;
	v874 = base.I32_ctz(v860) | v861<<(uint(int32(5))%32)
	goto L179
L181:
	;
	v874 = int32(-2)
	goto L179
L182:
	;
	v825 = v747 + int32(1)
	v827 = base.I32_div_s(v825, int32(32))
	v828 = *(*int32)(unsafe.Add(mBase, uint32(v683)+4))
	if v828 <= v827 {
		goto L181
	} else {
		goto L183
	}
L183:
	;
	v831 = v683 + int32(8)
	v835 = *(*int32)(unsafe.Add(mBase, uint32(v831+v827<<(uint(int32(2))%32))))
	v838 = v835 & (int32(-1) << (uint(v825) % 32))
	if v838 != 0 {
		v860 = v838
		v861 = v827
		goto L180
	} else {
		goto L184
	}
L184:
	;
	v840 = v827 + int32(1)
	if v840 == v828 {
		goto L181
	} else {
		goto L185
	}
L185:
	;
	v843 = v840
	goto L186
L186:
	;
	v850 = *(*int32)(unsafe.Add(mBase, uint32(v831+v843<<(uint(int32(2))%32))))
	if v850 != 0 {
		v860 = v850
		v861 = v843
		goto L180
	} else {
		goto L188
	}
L187:
	;
	goto L181
L188:
	;
	v852 = v843 + int32(1)
	if v852 != v828 {
		v843 = v852
		goto L186
	} else {
		goto L189
	}
L189:
	;
	goto L187
L190:
	;
	goto L165
L191:
	;
	F_bms_free(m, v683)
	mBase = m.M
	v908 = m.ExcPending
	if v908 != 0 {
		goto L25
	} else {
		goto L192
	}
L192:
	;
	v910 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v456))))
	if v910 != 0 {
		v914 = int32(1)
		goto L138
	} else {
		goto L193
	}
L193:
	;
	v911 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v460))))
	v914 = v911
	goto L138
L194:
	;
	v946 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v307)+74)))
	if v946 == int32(1) {
		goto L195
	} else {
		goto L196
	}
L195:
	;
	v949 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+15)))
	v950 = int32(1)
	v951 = v949 ^ v950
	*(*uint8)(unsafe.Add(mBase, uint32(v39)+6)) = uint8(v951)
	v953 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+14)))
	v955 = v953 ^ v950
	*(*uint8)(unsafe.Add(mBase, uint32(v39)+8)) = uint8(v955)
	goto L197
L196:
	;
	goto L197
L197:
	;
	v957 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v307)+75)))
	if v957 != int32(1) {
		goto L105
	} else {
		goto L198
	}
L198:
	;
	v960 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+15)))
	v961 = int32(1)
	v962 = v960 ^ v961
	*(*uint8)(unsafe.Add(mBase, uint32(v39)+7)) = uint8(v962)
	v964 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+14)))
	v966 = v964 ^ v961
	*(*uint8)(unsafe.Add(mBase, uint32(v39)+9)) = uint8(v966)
	goto L105
L199:
	;
	v997 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39))))
	if v997 != int32(1) {
		goto L200
	} else {
		goto L201
	}
L200:
	;
	v1034 = v279 + int32(1)
	v1035 = *(*int32)(unsafe.Add(mBase, uint32(v257)+4))
	if v1034 < v1035 {
		v279 = v1034
		goto L60
	} else {
		goto L217
	}
L201:
	;
	v1000 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+1)))
	if v1000 != int32(1) {
		goto L200
	} else {
		goto L202
	}
L202:
	;
	v1003 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+2)))
	if v1003 != int32(1) {
		goto L203
	} else {
		goto L204
	}
L203:
	;
	v1013 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+2)))
	if v1013 != int32(1) {
		goto L208
	} else {
		goto L209
	}
L204:
	;
	v1006 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+3)))
	if v1006 != int32(1) {
		goto L203
	} else {
		goto L205
	}
L205:
	;
	v1009 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v265))))
	if v1009 != 0 {
		goto L203
	} else {
		goto L206
	}
L206:
	;
	v1010 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+5)))
	if v1010 != int32(1) {
		goto L56
	} else {
		goto L207
	}
L207:
	;
	goto L203
L208:
	;
	v1023 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+2)))
	if v1023 != int32(1) {
		goto L200
	} else {
		goto L213
	}
L209:
	;
	v1016 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+3)))
	if v1016 != int32(1) {
		goto L208
	} else {
		goto L210
	}
L210:
	;
	v1019 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+6)))
	if v1019 != 0 {
		goto L208
	} else {
		goto L211
	}
L211:
	;
	v1020 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+7)))
	if v1020 != int32(1) {
		goto L56
	} else {
		goto L212
	}
L212:
	;
	goto L208
L213:
	;
	v1026 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+3)))
	if v1026 != int32(1) {
		goto L200
	} else {
		goto L214
	}
L214:
	;
	v1029 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+8)))
	if v1029 != 0 {
		goto L200
	} else {
		goto L215
	}
L215:
	;
	v1030 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+9)))
	if v1030 != int32(1) {
		goto L56
	} else {
		goto L216
	}
L216:
	;
	goto L200
L217:
	;
	goto L61
L218:
	;
	F_pfree(m, v1064)
	mBase = m.M
	v1066 = m.ExcPending
	if v1066 != 0 {
		goto L25
	} else {
		goto L221
	}
L219:
	;
	goto L220
L220:
	;
	v1069 = int32(_a_F_CheckCmdReplicaIdentity_3)
	v1070 = *(*int32)(unsafe.Add(mBase, _c_F_CheckCmdReplicaIdentity[0]))
	v1073 = *(*int32)(unsafe.Add(mBase, _c_F_CheckCmdReplicaIdentity[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_CheckCmdReplicaIdentity[0])) = v1073
	v1076 = F_palloc(m, int32(10))
	mBase = m.M
	v1077 = m.ExcPending
	if v1077 != 0 {
		goto L25
	} else {
		goto L222
	}
L221:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+176)) = int32(0)
	goto L220
L222:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+176)) = v1076
	v1079 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v39)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(v1076)+8)) = uint16(v1079)
	v1081 = *(*int64)(unsafe.Add(mBase, uint32(v39)))
	*(*int64)(unsafe.Add(mBase, uint32(v1076))) = v1081
	*(*int32)(unsafe.Add(mBase, _c_F_CheckCmdReplicaIdentity[0])) = v1070
	goto L18
L223:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v42))) = v299
	F_errmsg_internal(m, int32(_a_F_CheckCmdReplicaIdentity_4), v42)
	mBase = m.M
	v1122 = m.ExcPending
	if v1122 != 0 {
		goto L25
	} else {
		goto L224
	}
L224:
	;
	F_errfinish(m, int32(_a_F_CheckCmdReplicaIdentity_5), int32(_a_F_CheckCmdReplicaIdentity_6), int32(_a_F_CheckCmdReplicaIdentity_7))
	mBase = m.M
	v1127 = m.ExcPending
	if v1127 != 0 {
		goto L25
	} else {
		goto L225
	}
L225:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L226:
	;
	v1132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+122)))
	if v1132&int32(1) == int32(0) {
		goto L8
	} else {
		goto L229
	}
L227:
	;
	goto L228
L228:
	;
	if v1129 == int32(0) {
		goto L230
	} else {
		goto L231
	}
L229:
	;
	goto L228
L230:
	;
	v1139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+124)))
	if v1139&int32(1) == int32(0) {
		goto L7
	} else {
		goto L233
	}
L231:
	;
	goto L232
L232:
	;
	if l1 == int32(2) {
		goto L234
	} else {
		goto L235
	}
L233:
	;
	goto L232
L234:
	;
	v1146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+126)))
	if v1146&int32(1) == int32(0) {
		goto L6
	} else {
		goto L237
	}
L235:
	;
	goto L236
L236:
	;
	v1152 = base.B2i32(l1 != int32(4))
	if v1152 == int32(0) {
		goto L238
	} else {
		goto L239
	}
L237:
	;
	goto L236
L238:
	;
	v1155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+123)))
	if v1155&int32(1) == int32(0) {
		goto L5
	} else {
		goto L241
	}
L239:
	;
	goto L240
L240:
	;
	if v1152 == int32(0) {
		goto L242
	} else {
		goto L243
	}
L241:
	;
	goto L240
L242:
	;
	v1162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+125)))
	if v1162&int32(1) == int32(0) {
		goto L4
	} else {
		goto L245
	}
L243:
	;
	goto L244
L244:
	;
	if l1 == int32(4) {
		goto L246
	} else {
		goto L247
	}
L245:
	;
	goto L244
L246:
	;
	v1169 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+127)))
	if v1169&int32(1) == int32(0) {
		goto L3
	} else {
		goto L249
	}
L247:
	;
	goto L248
L248:
	;
	v1174 = F_RelationGetReplicaIndex(m, l0)
	mBase = m.M
	v1175 = m.ExcPending
	if v1175 != 0 {
		goto L25
	} else {
		goto L250
	}
L249:
	;
	goto L248
L250:
	;
	if v1174 != 0 {
		goto L9
	} else {
		goto L251
	}
L251:
	;
	v1176 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v1177 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1176)+130)))
	if v1177 == int32(102) {
		goto L9
	} else {
		goto L252
	}
L252:
	;
	if l1 == int32(2) {
		goto L253
	} else {
		goto L254
	}
L253:
	;
	v1182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+119)))
	if v1182&int32(1) != 0 {
		goto L2
	} else {
		goto L256
	}
L254:
	;
	goto L255
L255:
	;
	if l1 != int32(4) {
		goto L9
	} else {
		goto L257
	}
L256:
	;
	goto L255
L257:
	;
	v1187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+120)))
	if v1187&int32(1) != 0 {
		goto L1
	} else {
		goto L258
	}
L258:
	;
	goto L9
L259:
	;
	F_errcode(m, int32(_a_F_CheckCmdReplicaIdentity_8))
	mBase = m.M
	v1226 = m.ExcPending
	if v1226 != 0 {
		goto L25
	} else {
		goto L260
	}
L260:
	;
	v1227 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v30)+112)) = v1227 + int32(4)
	F_errmsg(m, int32(_a_F_CheckCmdReplicaIdentity_9), v30+int32(112))
	mBase = m.M
	v1235 = m.ExcPending
	if v1235 != 0 {
		goto L25
	} else {
		goto L261
	}
L261:
	;
	F_errdetail(m, int32(_a_F_CheckCmdReplicaIdentity_10), int32(0))
	mBase = m.M
	v1239 = m.ExcPending
	if v1239 != 0 {
		goto L25
	} else {
		goto L262
	}
L262:
	;
	F_errfinish(m, int32(_a_F_CheckCmdReplicaIdentity_11), int32(811), int32(_a_F_CheckCmdReplicaIdentity_12))
	mBase = m.M
	v1244 = m.ExcPending
	if v1244 != 0 {
		goto L25
	} else {
		goto L263
	}
L263:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L264:
	;
	F_errcode(m, int32(_a_F_CheckCmdReplicaIdentity_8))
	mBase = m.M
	v1251 = m.ExcPending
	if v1251 != 0 {
		goto L25
	} else {
		goto L265
	}
L265:
	;
	v1252 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v30)+96)) = v1252 + int32(4)
	F_errmsg(m, int32(_a_F_CheckCmdReplicaIdentity_9), v30+int32(96))
	mBase = m.M
	v1260 = m.ExcPending
	if v1260 != 0 {
		goto L25
	} else {
		goto L266
	}
L266:
	;
	F_errdetail(m, int32(_a_F_CheckCmdReplicaIdentity_13), int32(0))
	mBase = m.M
	v1264 = m.ExcPending
	if v1264 != 0 {
		goto L25
	} else {
		goto L267
	}
L267:
	;
	F_errfinish(m, int32(_a_F_CheckCmdReplicaIdentity_11), int32(817), int32(_a_F_CheckCmdReplicaIdentity_12))
	mBase = m.M
	v1269 = m.ExcPending
	if v1269 != 0 {
		goto L25
	} else {
		goto L268
	}
L268:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L269:
	;
	F_errcode(m, int32(_a_F_CheckCmdReplicaIdentity_8))
	mBase = m.M
	v1276 = m.ExcPending
	if v1276 != 0 {
		goto L25
	} else {
		goto L270
	}
L270:
	;
	v1277 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v30)+80)) = v1277 + int32(4)
	F_errmsg(m, int32(_a_F_CheckCmdReplicaIdentity_9), v30+int32(80))
	mBase = m.M
	v1285 = m.ExcPending
	if v1285 != 0 {
		goto L25
	} else {
		goto L271
	}
L271:
	;
	F_errdetail(m, int32(_a_F_CheckCmdReplicaIdentity_14), int32(0))
	mBase = m.M
	v1289 = m.ExcPending
	if v1289 != 0 {
		goto L25
	} else {
		goto L272
	}
L272:
	;
	F_errfinish(m, int32(_a_F_CheckCmdReplicaIdentity_11), int32(823), int32(_a_F_CheckCmdReplicaIdentity_12))
	mBase = m.M
	v1294 = m.ExcPending
	if v1294 != 0 {
		goto L25
	} else {
		goto L273
	}
L273:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L274:
	;
	F_errcode(m, int32(_a_F_CheckCmdReplicaIdentity_8))
	mBase = m.M
	v1301 = m.ExcPending
	if v1301 != 0 {
		goto L25
	} else {
		goto L275
	}
L275:
	;
	v1302 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v30)+64)) = v1302 + int32(4)
	F_errmsg(m, int32(_a_F_CheckCmdReplicaIdentity_15), v30-int32(-64))
	mBase = m.M
	v1310 = m.ExcPending
	if v1310 != 0 {
		goto L25
	} else {
		goto L276
	}
L276:
	;
	F_errdetail(m, int32(_a_F_CheckCmdReplicaIdentity_10), int32(0))
	mBase = m.M
	v1314 = m.ExcPending
	if v1314 != 0 {
		goto L25
	} else {
		goto L277
	}
L277:
	;
	F_errfinish(m, int32(_a_F_CheckCmdReplicaIdentity_11), int32(829), int32(_a_F_CheckCmdReplicaIdentity_12))
	mBase = m.M
	v1319 = m.ExcPending
	if v1319 != 0 {
		goto L25
	} else {
		goto L278
	}
L278:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L279:
	;
	F_errcode(m, int32(_a_F_CheckCmdReplicaIdentity_8))
	mBase = m.M
	v1326 = m.ExcPending
	if v1326 != 0 {
		goto L25
	} else {
		goto L280
	}
L280:
	;
	v1327 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v30)+48)) = v1327 + int32(4)
	F_errmsg(m, int32(_a_F_CheckCmdReplicaIdentity_15), v30+int32(48))
	mBase = m.M
	v1335 = m.ExcPending
	if v1335 != 0 {
		goto L25
	} else {
		goto L281
	}
L281:
	;
	F_errdetail(m, int32(_a_F_CheckCmdReplicaIdentity_13), int32(0))
	mBase = m.M
	v1339 = m.ExcPending
	if v1339 != 0 {
		goto L25
	} else {
		goto L282
	}
L282:
	;
	F_errfinish(m, int32(_a_F_CheckCmdReplicaIdentity_11), int32(835), int32(_a_F_CheckCmdReplicaIdentity_12))
	mBase = m.M
	v1344 = m.ExcPending
	if v1344 != 0 {
		goto L25
	} else {
		goto L283
	}
L283:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L284:
	;
	F_errcode(m, int32(_a_F_CheckCmdReplicaIdentity_8))
	mBase = m.M
	v1351 = m.ExcPending
	if v1351 != 0 {
		goto L25
	} else {
		goto L285
	}
L285:
	;
	v1352 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v30)+32)) = v1352 + int32(4)
	F_errmsg(m, int32(_a_F_CheckCmdReplicaIdentity_15), v30+int32(32))
	mBase = m.M
	v1360 = m.ExcPending
	if v1360 != 0 {
		goto L25
	} else {
		goto L286
	}
L286:
	;
	F_errdetail(m, int32(_a_F_CheckCmdReplicaIdentity_14), int32(0))
	mBase = m.M
	v1364 = m.ExcPending
	if v1364 != 0 {
		goto L25
	} else {
		goto L287
	}
L287:
	;
	F_errfinish(m, int32(_a_F_CheckCmdReplicaIdentity_11), int32(841), int32(_a_F_CheckCmdReplicaIdentity_12))
	mBase = m.M
	v1369 = m.ExcPending
	if v1369 != 0 {
		goto L25
	} else {
		goto L288
	}
L288:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L289:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v1376 = m.ExcPending
	if v1376 != 0 {
		goto L25
	} else {
		goto L290
	}
L290:
	;
	v1377 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v30))) = v1377 + int32(4)
	F_errmsg(m, int32(_a_F_CheckCmdReplicaIdentity_16), v30)
	mBase = m.M
	v1383 = m.ExcPending
	if v1383 != 0 {
		goto L25
	} else {
		goto L291
	}
L291:
	;
	F_errhint(m, int32(_a_F_CheckCmdReplicaIdentity_17), int32(0))
	mBase = m.M
	v1387 = m.ExcPending
	if v1387 != 0 {
		goto L25
	} else {
		goto L292
	}
L292:
	;
	F_errfinish(m, int32(_a_F_CheckCmdReplicaIdentity_11), int32(861), int32(_a_F_CheckCmdReplicaIdentity_12))
	mBase = m.M
	v1392 = m.ExcPending
	if v1392 != 0 {
		goto L25
	} else {
		goto L293
	}
L293:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L294:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v1399 = m.ExcPending
	if v1399 != 0 {
		goto L25
	} else {
		goto L295
	}
L295:
	;
	v1400 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v30)+16)) = v1400 + int32(4)
	F_errmsg(m, int32(_a_F_CheckCmdReplicaIdentity_18), v30+int32(16))
	mBase = m.M
	v1408 = m.ExcPending
	if v1408 != 0 {
		goto L25
	} else {
		goto L296
	}
L296:
	;
	F_errhint(m, int32(_a_F_CheckCmdReplicaIdentity_19), int32(0))
	mBase = m.M
	v1412 = m.ExcPending
	if v1412 != 0 {
		goto L25
	} else {
		goto L297
	}
L297:
	;
	F_errfinish(m, int32(_a_F_CheckCmdReplicaIdentity_11), int32(867), int32(_a_F_CheckCmdReplicaIdentity_12))
	mBase = m.M
	v1417 = m.ExcPending
	if v1417 != 0 {
		goto L25
	} else {
		goto L298
	}
L298:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_CheckDim_1(m *base.Module, l0 int32) {
	var v10 int32
	_ = v10
	Fn13823(m, l0, int32(109), int32(_a_F_CheckDim_1_0), int32(_a_F_CheckDim_1_1), int32(_a_F_CheckDim_1_2), int32(104), int32(_a_F_CheckDim_1_3), int32(_a_F_CheckDim_1_4))
	v10 = m.ExcPending
	if v10 != 0 {
		return
	} else {
		return
	}
}
func F_CheckElement_1(m *base.Module, l0 int32) {
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	v2 = int32(_a_F_CheckElement_1_0)
	v7 = l0 & int32(_a_F_CheckElement_1_1)
	if base.B2i32(l0&v2 == v2)&base.B2i32(v7 != v2) == int32(0) {
		if v7 == int32(_a_F_CheckElement_1_0) {
			F_errstart_cold(m, int32(21), int32(0))
			v34 = m.ExcPending
			if v34 != 0 {
				return
			} else {
				F_errcode(m, int32(130))
				v37 = m.ExcPending
				if v37 != 0 {
					return
				} else {
					F_errmsg(m, int32(_a_F_CheckElement_1_2), int32(0))
					v41 = m.ExcPending
					if v41 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_CheckElement_1_3), int32(126), int32(_a_F_CheckElement_1_4))
						v46 = m.ExcPending
						if v46 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		} else {
			return
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		v18 = m.ExcPending
		if v18 != 0 {
			return
		} else {
			F_errcode(m, int32(130))
			v21 = m.ExcPending
			if v21 != 0 {
				return
			} else {
				F_errmsg(m, int32(_a_F_CheckElement_1_5), int32(0))
				v25 = m.ExcPending
				if v25 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_CheckElement_1_3), int32(121), int32(_a_F_CheckElement_1_4))
					v30 = m.ExcPending
					if v30 != 0 {
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
func F_CheckNnz(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	if int32(0) <= l0 {
		if base.Ui32(int32(_a_F_CheckNnz_0)) <= base.Ui32(l0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v35 = m.ExcPending
			if v35 != 0 {
				return
			} else {
				F_errcode(m, int32(261))
				mBase = m.M
				v38 = m.ExcPending
				if v38 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v6))) = int32(_a_F_CheckNnz_1)
					F_errmsg(m, int32(_a_F_CheckNnz_2), v6)
					mBase = m.M
					v43 = m.ExcPending
					if v43 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_CheckNnz_3), int32(96), int32(_a_F_CheckNnz_4))
						mBase = m.M
						v48 = m.ExcPending
						if v48 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		} else {
			if l1 < l0 {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v52 = m.ExcPending
				if v52 != 0 {
					return
				} else {
					F_errcode(m, int32(261))
					mBase = m.M
					v55 = m.ExcPending
					if v55 != 0 {
						return
					} else {
						F_errmsg(m, int32(_a_F_CheckNnz_5), int32(0))
						mBase = m.M
						v59 = m.ExcPending
						if v59 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_CheckNnz_3), int32(101), int32(_a_F_CheckNnz_4))
							mBase = m.M
							v64 = m.ExcPending
							if v64 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			} else {
				m.G0 = v6 + int32(16)
				return
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return
		} else {
			F_errcode(m, int32(130))
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return
			} else {
				F_errmsg(m, int32(_a_F_CheckNnz_6), int32(0))
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_CheckNnz_3), int32(91), int32(_a_F_CheckNnz_4))
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
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
func F_CleanQuerytext(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v48 int32
	_ = v48
	var v55 int32
	_ = v55
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v8 < int32(0) {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v85
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v84
	return v81
L2:
	;
	v26 = v22
	v29 = v23
	v30 = v24
	goto L9
L3:
	;
	v19 = F_strlen(m, v16)
	mBase = m.M
	if v19 <= int32(0) {
		v81 = v16
		v84 = v19
		v85 = v18
		goto L1
	} else {
		goto L8
	}
L4:
	;
	v16 = l0
	v18 = int32(0)
	goto L3
L5:
	;
	goto L6
L6:
	;
	v12 = l0 + v8
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if int32(0) < v13 {
		v22 = v12
		v23 = v13
		v24 = v8
		goto L2
	} else {
		goto L7
	}
L7:
	;
	v16 = v12
	v18 = v8
	goto L3
L8:
	;
	v22 = v16
	v23 = v19
	v24 = v18
	goto L2
L9:
	;
	v33 = int32(*(*int8)(unsafe.Add(mBase, uint32(v26))))
	goto L11
L10:
	;
	v81 = v75
	v84 = int32(0)
	v85 = v23 + v24
	goto L1
L11:
	;
	if base.B2i32(v33 == int32(32))|base.B2i32(base.Ui32((v33-int32(9))&int32(255)) < base.Ui32(int32(5))) == int32(0) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v48 = v29
	goto L15
L13:
	;
	goto L14
L14:
	;
	v72 = int32(1)
	v75 = v26 + v72
	if v72 < v29 {
		v26 = v75
		v29 = v29 - v72
		v30 = v30 + v72
		goto L9
	} else {
		goto L20
	}
L15:
	;
	v55 = int32(*(*int8)(unsafe.Add(mBase, uint32(v26+v48-int32(1)))))
	goto L17
L16:
	;
	v81 = v26
	v84 = int32(0)
	v85 = v30
	goto L1
L17:
	;
	if base.B2i32(v55 == int32(32))|base.B2i32(base.Ui32((v55-int32(9))&int32(255)) < base.Ui32(int32(5))) == int32(0) {
		v81 = v26
		v84 = v48
		v85 = v30
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v67 = int32(1)
	if v67 < v48 {
		v48 = v48 - v67
		goto L15
	} else {
		goto L19
	}
L19:
	;
	goto L16
L20:
	;
	goto L10
}
func F_CompareCandidateDistances(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 float32
	_ = v7
	var v8 int32
	_ = v8
	var v9 float32
	_ = v9
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v7 = *(*float32)(unsafe.Add(mBase, uint32(v6)+4))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v9 = *(*float32)(unsafe.Add(mBase, uint32(v8)+4))
	if base.F32_lt(v7, v9) != 0 {
		v23 = int32(1)
	} else {
		if base.F32_gt(v7, v9) != 0 {
			v23 = int32(-1)
		} else {
			v14 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
			v15 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
			if base.Ui32(v14) < base.Ui32(v15) {
				v23 = int32(1)
			} else {
				if base.Ui32(v15) < base.Ui32(v14) {
					v20 = int32(-1)
				} else {
					v20 = int32(0)
				}
				v23 = v20
			}
		}
	}
	return v23
}
func F_ComputeXidHorizons(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
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
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int64
	_ = v45
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v76 int32
	_ = v76
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v172 int32
	_ = v172
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v394 int32
	_ = v394
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v424 int32
	_ = v424
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v446 int32
	_ = v446
	var v449 int32
	_ = v449
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v487 int32
	_ = v487
	var v489 int32
	_ = v489
	var v492 int32
	_ = v492
	var v494 int32
	_ = v494
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v517 int64
	_ = v517
	var v518 int32
	_ = v518
	var v522 int64
	_ = v522
	var v525 int64
	_ = v525
	var v526 int32
	_ = v526
	var v530 int64
	_ = v530
	var v533 int64
	_ = v533
	var v537 int64
	_ = v537
	var v540 int64
	_ = v540
	var v541 int32
	_ = v541
	var v545 int64
	_ = v545
	var v547 int32
	_ = v547
	var v549 int64
	_ = v549
	var v551 int64
	_ = v551
	var v553 int64
	_ = v553
	var v555 int64
	_ = v555
	var v557 int32
	_ = v557
	var v559 int64
	_ = v559
	var v561 int64
	_ = v561
	var v563 int64
	_ = v563
	var v565 int64
	_ = v565
	var v567 int32
	_ = v567
	var v568 int64
	_ = v568
	var v573 int64
	_ = v573
	var v575 int64
	_ = v575
	var v577 int64
	_ = v577
	var v581 int32
	_ = v581
	v2 = int32(0)
	v17 = *(*int32)(unsafe.Add(mBase, _c_F_ComputeXidHorizons[0]))
	v20 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ComputeXidHorizons[1])))
	if v20 == int32(1) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v32 = *(*int32)(unsafe.Add(mBase, _c_F_ComputeXidHorizons[2]))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = int32(0)
	v37 = *(*int32)(unsafe.Add(mBase, _c_F_ComputeXidHorizons[3]))
	v41 = F_LWLockAcquire(m, v37+int32(512), int32(1))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	v25 = *(*int32)(unsafe.Add(mBase, _c_F_ComputeXidHorizons[4]))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+316))
	v28 = base.B2i32(v26 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _c_F_ComputeXidHorizons[1])) = uint8(v28)
	v30 = v28
	goto L4
L3:
	;
	v30 = v2
	goto L4
L4:
	;
	goto L1
L5:
	;
	return
L6:
	;
	v44 = *(*int32)(unsafe.Add(mBase, _c_F_ComputeXidHorizons[5]))
	v45 = *(*int64)(unsafe.Add(mBase, uint32(v44)+48))
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v45
	v47 = int32(3)
	v50 = base.I32_wrap_i64(v45) + int32(1)
	if base.Ui32(v50) <= base.Ui32(v47) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v53 = v47
	goto L9
L8:
	;
	v53 = v50
	goto L9
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v53
	v58 = *(*int32)(unsafe.Add(mBase, _c_F_ComputeXidHorizons[6]))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)+36))
	if v59 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v60 = v59
	goto L12
L11:
	;
	v60 = v53
	goto L12
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v60
	v63 = *(*int32)(unsafe.Add(mBase, _c_F_ComputeXidHorizons[0]))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)+28))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v64
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v63)+32))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v66
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	if int32(0) < v68 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v76 = v2
	goto L16
L14:
	;
	goto L15
L15:
	;
	if v30 != 0 {
		goto L70
	} else {
		goto L71
	}
L16:
	;
	v89 = *(*int32)(unsafe.Add(mBase, _c_F_ComputeXidHorizons[2]))
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v89)+12))
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90+v76))))
	v94 = v76 << (uint(int32(2)) % 32)
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v17+int32(36)+v94)))
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v94+v33)))
	v100 = *(*int32)(unsafe.Add(mBase, _c_F_ComputeXidHorizons[7]))
	v103 = v100 + v96*int32(640)
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v103)+40))
	if v104 != 0 {
		goto L20
	} else {
		goto L21
	}
L17:
	;
	goto L15
L18:
	;
	v191 = v76 + int32(1)
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	if v191 < v192 {
		v76 = v191
		goto L16
	} else {
		goto L66
	}
L19:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v123 != 0 {
		goto L34
	} else {
		goto L35
	}
L20:
	;
	if v98 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	goto L22
L22:
	;
	if v98 == int32(0) {
		goto L18
	} else {
		goto L33
	}
L23:
	;
	v122 = v104
	goto L19
L24:
	;
	goto L25
L25:
	;
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v98))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v104)) == int32(0) {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	if v118 != 0 {
		goto L30
	} else {
		goto L31
	}
L27:
	;
	v118 = base.B2i32(base.Ui32(v104) < base.Ui32(v98))
	goto L26
L28:
	;
	goto L29
L29:
	;
	v118 = int32(base.Ui32(v104-v98) >> (uint(int32(31)) % 32))
	goto L26
L30:
	;
	v119 = v104
	goto L32
L31:
	;
	v119 = v98
	goto L32
L32:
	;
	v122 = v119
	goto L19
L33:
	;
	v122 = v98
	goto L19
L34:
	;
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v122))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v123)) == int32(0) {
		goto L38
	} else {
		goto L39
	}
L35:
	;
	v137 = v122
	goto L36
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v137
	if v92&int32(18) != 0 {
		goto L18
	} else {
		goto L44
	}
L37:
	;
	if v135 != 0 {
		goto L41
	} else {
		goto L42
	}
L38:
	;
	v135 = base.B2i32(base.Ui32(v123) < base.Ui32(v122))
	goto L37
L39:
	;
	goto L40
L40:
	;
	v135 = int32(base.Ui32(v123-v122) >> (uint(int32(31)) % 32))
	goto L37
L41:
	;
	v136 = v123
	goto L43
L42:
	;
	v136 = v122
	goto L43
L43:
	;
	v137 = v136
	goto L36
L44:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v141 != 0 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v122))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v141)) == int32(0) {
		goto L49
	} else {
		goto L50
	}
L46:
	;
	v155 = v122
	goto L47
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v155
	v158 = *(*int32)(unsafe.Add(mBase, _c_F_ComputeXidHorizons[8]))
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v103)+60))
	v161 = int32(0)
	if base.B2i32(v158 == v159)|base.B2i32(v158 == v161)|(int32(base.Ui32(v92)>>(uint(int32(5))%32))|v30)&int32(1) == v161 {
		goto L18
	} else {
		goto L55
	}
L48:
	;
	if v153 != 0 {
		goto L52
	} else {
		goto L53
	}
L49:
	;
	v153 = base.B2i32(base.Ui32(v141) < base.Ui32(v122))
	goto L48
L50:
	;
	goto L51
L51:
	;
	v153 = int32(base.Ui32(v141-v122) >> (uint(int32(31)) % 32))
	goto L48
L52:
	;
	v154 = v141
	goto L54
L53:
	;
	v154 = v122
	goto L54
L54:
	;
	v155 = v154
	goto L47
L55:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v172 != 0 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v122))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v172)) == int32(0) {
		goto L60
	} else {
		goto L61
	}
L57:
	;
	v186 = v122
	goto L58
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v186
	goto L18
L59:
	;
	if v184 != 0 {
		goto L63
	} else {
		goto L64
	}
L60:
	;
	v184 = base.B2i32(base.Ui32(v172) < base.Ui32(v122))
	goto L59
L61:
	;
	goto L62
L62:
	;
	v184 = int32(base.Ui32(v172-v122) >> (uint(int32(31)) % 32))
	goto L59
L63:
	;
	v185 = v172
	goto L65
L64:
	;
	v185 = v122
	goto L65
L65:
	;
	v186 = v185
	goto L58
L66:
	;
	goto L17
L67:
	;
	v348 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v349 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v349 == int32(0) {
		goto L123
	} else {
		goto L124
	}
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v312
	v314 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v314 != 0 {
		goto L103
	} else {
		goto L104
	}
L69:
	;
	v310 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v312 = v310
	goto L68
L70:
	;
	v210 = *(*int32)(unsafe.Add(mBase, _c_F_ComputeXidHorizons[0]))
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v210)+20))
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v210)+16))
	if v211 <= v213 {
		v261 = int32(0)
		goto L73
	} else {
		goto L74
	}
L71:
	;
	goto L72
L72:
	;
	v305 = *(*int32)(unsafe.Add(mBase, _c_F_ComputeXidHorizons[3]))
	F_LWLockRelease(m, v305+int32(512))
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L5
	} else {
		goto L102
	}
L73:
	;
	v263 = *(*int32)(unsafe.Add(mBase, _c_F_ComputeXidHorizons[3]))
	F_LWLockRelease(m, v263+int32(512))
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L5
	} else {
		goto L81
	}
L74:
	;
	v216 = *(*int32)(unsafe.Add(mBase, _c_F_ComputeXidHorizons[9]))
	v218 = v213
	goto L75
L75:
	;
	v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v218+v216))))
	if v233 == int32(1) {
		goto L77
	} else {
		goto L78
	}
L76:
	;
	v261 = int32(0)
	goto L73
L77:
	;
	v237 = *(*int32)(unsafe.Add(mBase, _c_F_ComputeXidHorizons[10]))
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v237+v218<<(uint(int32(2))%32))))
	v261 = v241
	goto L73
L78:
	;
	goto L79
L79:
	;
	v243 = v218 + int32(1)
	if v243 != v211 {
		v218 = v243
		goto L75
	} else {
		goto L80
	}
L80:
	;
	goto L76
L81:
	;
	v268 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v268 != 0 {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	if v261 == int32(0) {
		goto L69
	} else {
		goto L85
	}
L83:
	;
	v284 = v261
	goto L84
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v284
	v286 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v286 == int32(0) {
		v312 = v261
		goto L68
	} else {
		goto L93
	}
L85:
	;
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v261))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v268)) == int32(0) {
		goto L87
	} else {
		goto L88
	}
L86:
	;
	if v282 != 0 {
		goto L90
	} else {
		goto L91
	}
L87:
	;
	v282 = base.B2i32(base.Ui32(v268) < base.Ui32(v261))
	goto L86
L88:
	;
	goto L89
L89:
	;
	v282 = int32(base.Ui32(v268-v261) >> (uint(int32(31)) % 32))
	goto L86
L90:
	;
	v283 = v268
	goto L92
L91:
	;
	v283 = v261
	goto L92
L92:
	;
	v284 = v283
	goto L84
L93:
	;
	if v261 == int32(0) {
		v312 = v286
		goto L68
	} else {
		goto L94
	}
L94:
	;
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v261))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v286)) == int32(0) {
		goto L96
	} else {
		goto L97
	}
L95:
	;
	if v302 != 0 {
		goto L99
	} else {
		goto L100
	}
L96:
	;
	v302 = base.B2i32(base.Ui32(v286) < base.Ui32(v261))
	goto L95
L97:
	;
	goto L98
L98:
	;
	v302 = int32(base.Ui32(v286-v261) >> (uint(int32(31)) % 32))
	goto L95
L99:
	;
	v303 = v286
	goto L101
L100:
	;
	v303 = v261
	goto L101
L101:
	;
	v312 = v303
	goto L68
L102:
	;
	goto L67
L103:
	;
	if v261 == int32(0) {
		goto L106
	} else {
		goto L107
	}
L104:
	;
	v331 = v261
	goto L105
L105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v331
	goto L67
L106:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v314
	goto L67
L107:
	;
	goto L108
L108:
	;
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v261))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v314)) == int32(0) {
		goto L110
	} else {
		goto L111
	}
L109:
	;
	if v329 != 0 {
		goto L113
	} else {
		goto L114
	}
L110:
	;
	v329 = base.B2i32(base.Ui32(v314) < base.Ui32(v261))
	goto L109
L111:
	;
	goto L112
L112:
	;
	v329 = int32(base.Ui32(v314-v261) >> (uint(int32(31)) % 32))
	goto L109
L113:
	;
	v330 = v314
	goto L115
L114:
	;
	v330 = v261
	goto L115
L115:
	;
	v331 = v330
	goto L105
L116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v444
	v449 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v449 != 0 {
		goto L170
	} else {
		goto L171
	}
L117:
	;
	if v422 == int32(0) {
		goto L156
	} else {
		goto L157
	}
L118:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v402
	v444 = v400
	v446 = v402
	goto L116
L119:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v421
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v424
	if v421 != 0 {
		goto L117
	} else {
		goto L155
	}
L120:
	;
	if v401 == int32(0) {
		goto L118
	} else {
		goto L147
	}
L121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v349
	v398 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v399 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v400 = v399
	v401 = v398
	v402 = v349
	goto L120
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v369
	v371 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v371 == int32(0) {
		v390 = v368
		v391 = v369
		goto L134
	} else {
		goto L135
	}
L123:
	;
	v368 = v348
	v369 = v348
	goto L122
L124:
	;
	goto L125
L125:
	;
	if v348 == int32(0) {
		goto L121
	} else {
		goto L126
	}
L126:
	;
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v348))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v349)) == int32(0) {
		goto L128
	} else {
		goto L129
	}
L127:
	;
	if v365 != 0 {
		goto L131
	} else {
		goto L132
	}
L128:
	;
	v365 = base.B2i32(base.Ui32(v349) < base.Ui32(v348))
	goto L127
L129:
	;
	goto L130
L130:
	;
	v365 = int32(base.Ui32(v349-v348) >> (uint(int32(31)) % 32))
	goto L127
L131:
	;
	v366 = v349
	goto L133
L132:
	;
	v366 = v348
	goto L133
L133:
	;
	v367 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v368 = v367
	v369 = v366
	goto L122
L134:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v391
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v390
	v394 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v391 == int32(0) {
		v421 = v390
		v422 = v394
		v424 = v394
		goto L119
	} else {
		goto L146
	}
L135:
	;
	if v368 == int32(0) {
		goto L136
	} else {
		goto L137
	}
L136:
	;
	v390 = v371
	v391 = v369
	goto L134
L137:
	;
	goto L138
L138:
	;
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v368))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v371)) == int32(0) {
		goto L140
	} else {
		goto L141
	}
L139:
	;
	if v387 != 0 {
		goto L143
	} else {
		goto L144
	}
L140:
	;
	v387 = base.B2i32(base.Ui32(v371) < base.Ui32(v368))
	goto L139
L141:
	;
	goto L142
L142:
	;
	v387 = int32(base.Ui32(v371-v368) >> (uint(int32(31)) % 32))
	goto L139
L143:
	;
	v388 = v371
	goto L145
L144:
	;
	v388 = v368
	goto L145
L145:
	;
	v389 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v390 = v388
	v391 = v389
	goto L134
L146:
	;
	v400 = v390
	v401 = v394
	v402 = v391
	goto L120
L147:
	;
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v401))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v402)) == int32(0) {
		goto L149
	} else {
		goto L150
	}
L148:
	;
	if v417 != 0 {
		goto L152
	} else {
		goto L153
	}
L149:
	;
	v417 = base.B2i32(base.Ui32(v402) < base.Ui32(v401))
	goto L148
L150:
	;
	goto L151
L151:
	;
	v417 = int32(base.Ui32(v402-v401) >> (uint(int32(31)) % 32))
	goto L148
L152:
	;
	v418 = v402
	goto L154
L153:
	;
	v418 = v401
	goto L154
L154:
	;
	v419 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v420 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v421 = v420
	v422 = v419
	v424 = v418
	goto L119
L155:
	;
	v444 = v422
	v446 = v424
	goto L116
L156:
	;
	v444 = v421
	v446 = v424
	goto L116
L157:
	;
	goto L158
L158:
	;
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v422))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v421)) == int32(0) {
		goto L160
	} else {
		goto L161
	}
L159:
	;
	if v441 != 0 {
		goto L163
	} else {
		goto L164
	}
L160:
	;
	v441 = base.B2i32(base.Ui32(v421) < base.Ui32(v422))
	goto L159
L161:
	;
	goto L162
L162:
	;
	v441 = int32(base.Ui32(v421-v422) >> (uint(int32(31)) % 32))
	goto L159
L163:
	;
	v442 = v421
	goto L165
L164:
	;
	v442 = v422
	goto L165
L165:
	;
	v443 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v444 = v442
	v446 = v443
	goto L116
L166:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v513
	v517 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	v518 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v522 = v517 + base.I64_extend_i32_s(v518-base.I32_wrap_i64(v517))
	*(*int64)(unsafe.Add(mBase, _c_F_ComputeXidHorizons[11])) = v522
	v525 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	v526 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v530 = v525 + base.I64_extend_i32_s(v526-base.I32_wrap_i64(v525))
	*(*int64)(unsafe.Add(mBase, _c_F_ComputeXidHorizons[12])) = v530
	v533 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	v537 = v533 + base.I64_extend_i32_s(v514-base.I32_wrap_i64(v533))
	*(*int64)(unsafe.Add(mBase, _c_F_ComputeXidHorizons[13])) = v537
	v540 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	v541 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v545 = v540 + base.I64_extend_i32_s(v541-base.I32_wrap_i64(v540))
	*(*int64)(unsafe.Add(mBase, _c_F_ComputeXidHorizons[14])) = v545
	v547 = int32(_a_F_ComputeXidHorizons_0)
	v549 = *(*int64)(unsafe.Add(mBase, _c_F_ComputeXidHorizons[15]))
	if base.Ui64(v549) < base.Ui64(v522) {
		goto L205
	} else {
		goto L206
	}
L167:
	;
	if v494 == int32(0) {
		goto L195
	} else {
		goto L196
	}
L168:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v444
	v489 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v444 == int32(0) {
		v513 = v489
		v514 = v489
		goto L166
	} else {
		goto L194
	}
L169:
	;
	if v470 != 0 {
		goto L184
	} else {
		goto L185
	}
L170:
	;
	if v446 != 0 {
		goto L173
	} else {
		goto L174
	}
L171:
	;
	goto L172
L172:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v446
	if v446 == int32(0) {
		goto L168
	} else {
		goto L183
	}
L173:
	;
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v446))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v449)) == int32(0) {
		goto L177
	} else {
		goto L178
	}
L174:
	;
	v464 = v444
	v465 = v449
	goto L175
L175:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v465
	v470 = v464
	v471 = v465
	goto L169
L176:
	;
	if v461 != 0 {
		goto L180
	} else {
		goto L181
	}
L177:
	;
	v461 = base.B2i32(base.Ui32(v449) < base.Ui32(v446))
	goto L176
L178:
	;
	goto L179
L179:
	;
	v461 = int32(base.Ui32(v449-v446) >> (uint(int32(31)) % 32))
	goto L176
L180:
	;
	v462 = v449
	goto L182
L181:
	;
	v462 = v446
	goto L182
L182:
	;
	v463 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v464 = v463
	v465 = v462
	goto L175
L183:
	;
	v470 = v444
	v471 = v446
	goto L169
L184:
	;
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v470))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v471)) == int32(0) {
		goto L188
	} else {
		goto L189
	}
L185:
	;
	v485 = v471
	goto L186
L186:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v485
	v487 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v492 = v485
	v494 = v487
	goto L167
L187:
	;
	if v483 != 0 {
		goto L191
	} else {
		goto L192
	}
L188:
	;
	v483 = base.B2i32(base.Ui32(v471) < base.Ui32(v470))
	goto L187
L189:
	;
	goto L190
L190:
	;
	v483 = int32(base.Ui32(v471-v470) >> (uint(int32(31)) % 32))
	goto L187
L191:
	;
	v484 = v471
	goto L193
L192:
	;
	v484 = v470
	goto L193
L193:
	;
	v485 = v484
	goto L186
L194:
	;
	v492 = v444
	v494 = v489
	goto L167
L195:
	;
	v513 = v492
	v514 = int32(0)
	goto L166
L196:
	;
	goto L197
L197:
	;
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v494))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v492)) == int32(0) {
		goto L199
	} else {
		goto L200
	}
L198:
	;
	if v509 != 0 {
		goto L202
	} else {
		goto L203
	}
L199:
	;
	v509 = base.B2i32(base.Ui32(v492) < base.Ui32(v494))
	goto L198
L200:
	;
	goto L201
L201:
	;
	v509 = int32(base.Ui32(v492-v494) >> (uint(int32(31)) % 32))
	goto L198
L202:
	;
	v510 = v492
	goto L204
L203:
	;
	v510 = v494
	goto L204
L204:
	;
	v511 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v513 = v510
	v514 = v511
	goto L166
L205:
	;
	v551 = v522
	goto L207
L206:
	;
	v551 = v549
	goto L207
L207:
	;
	if base.I32_wrap_i64(v549) != 0 {
		goto L208
	} else {
		goto L209
	}
L208:
	;
	v553 = v551
	goto L210
L209:
	;
	v553 = v522
	goto L210
L210:
	;
	if base.I32_wrap_i64(v522) != 0 {
		goto L211
	} else {
		goto L212
	}
L211:
	;
	v555 = v553
	goto L213
L212:
	;
	v555 = v549
	goto L213
L213:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_ComputeXidHorizons[15])) = v555
	v557 = int32(_a_F_ComputeXidHorizons_1)
	v559 = *(*int64)(unsafe.Add(mBase, _c_F_ComputeXidHorizons[16]))
	if base.Ui64(v559) < base.Ui64(v530) {
		goto L214
	} else {
		goto L215
	}
L214:
	;
	v561 = v530
	goto L216
L215:
	;
	v561 = v559
	goto L216
L216:
	;
	if base.I32_wrap_i64(v559) != 0 {
		goto L217
	} else {
		goto L218
	}
L217:
	;
	v563 = v561
	goto L219
L218:
	;
	v563 = v530
	goto L219
L219:
	;
	if base.I32_wrap_i64(v530) != 0 {
		goto L220
	} else {
		goto L221
	}
L220:
	;
	v565 = v563
	goto L222
L221:
	;
	v565 = v559
	goto L222
L222:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_ComputeXidHorizons[16])) = v565
	v567 = int32(_a_F_ComputeXidHorizons_2)
	v568 = *(*int64)(unsafe.Add(mBase, _c_F_ComputeXidHorizons[17]))
	*(*int64)(unsafe.Add(mBase, _c_F_ComputeXidHorizons[18])) = v545
	if base.Ui64(v568) < base.Ui64(v537) {
		goto L223
	} else {
		goto L224
	}
L223:
	;
	v573 = v537
	goto L225
L224:
	;
	v573 = v568
	goto L225
L225:
	;
	if base.I32_wrap_i64(v568) != 0 {
		goto L226
	} else {
		goto L227
	}
L226:
	;
	v575 = v573
	goto L228
L227:
	;
	v575 = v537
	goto L228
L228:
	;
	if base.I32_wrap_i64(v537) != 0 {
		goto L229
	} else {
		goto L230
	}
L229:
	;
	v577 = v575
	goto L231
L230:
	;
	v577 = v568
	goto L231
L231:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_ComputeXidHorizons[17])) = v577
	v581 = *(*int32)(unsafe.Add(mBase, _c_F_ComputeXidHorizons[19]))
	*(*int32)(unsafe.Add(mBase, _c_F_ComputeXidHorizons[20])) = v581
	return
}
func F_ConditionVariableCancelSleep(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v54 int32
	_ = v54
	v7 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariableCancelSleep[0]))
	if v7 != 0 {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
		*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(1)
		if v8 != 0 {
			F_s_lock(m, v7, int32(_a_F_ConditionVariableCancelSleep_0), int32(238), int32(_a_F_ConditionVariableCancelSleep_1))
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return
			} else {
				v17 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariableCancelSleep[1]))
				v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
				v20 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariableCancelSleep[2]))
				v23 = v18 + v20*int32(640)
				v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+88))
				if v24 == int32(0) {
					v27 = *(*int32)(unsafe.Add(mBase, uint32(v23)+84))
					if v27 != 0 {
						v33 = v27
						*(*int32)(unsafe.Add(mBase, uint32(v18+v24*int32(640))+84)) = v33
						v38 = v24
						v39 = v33
						if v39 == int32(-1) {
							*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v38
						} else {
							v44 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariableCancelSleep[1]))
							v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
							*(*int32)(unsafe.Add(mBase, uint32(v45+v39*int32(640))+88)) = v38
						}
						*(*int64)(unsafe.Add(mBase, uint32(v23)+84)) = int64(0)
					} else {
					}
				} else {
					v28 = *(*int32)(unsafe.Add(mBase, uint32(v23)+84))
					if v24 != int32(-1) {
						v33 = v28
						*(*int32)(unsafe.Add(mBase, uint32(v18+v24*int32(640))+84)) = v33
						v38 = v24
						v39 = v33
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v28
						v32 = *(*int32)(unsafe.Add(mBase, uint32(v23)+88))
						v38 = v32
						v39 = v28
					}
					if v39 == int32(-1) {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v38
					} else {
						v44 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariableCancelSleep[1]))
						v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
						*(*int32)(unsafe.Add(mBase, uint32(v45+v39*int32(640))+88)) = v38
					}
					*(*int64)(unsafe.Add(mBase, uint32(v23)+84)) = int64(0)
				}
				v54 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v7))) = v54
				*(*int32)(unsafe.Add(mBase, _c_F_ConditionVariableCancelSleep[0])) = v54
				return
			}
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariableCancelSleep[1]))
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
			v20 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariableCancelSleep[2]))
			v23 = v18 + v20*int32(640)
			v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+88))
			if v24 == int32(0) {
				v27 = *(*int32)(unsafe.Add(mBase, uint32(v23)+84))
				if v27 != 0 {
					v33 = v27
					*(*int32)(unsafe.Add(mBase, uint32(v18+v24*int32(640))+84)) = v33
					v38 = v24
					v39 = v33
					if v39 == int32(-1) {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v38
					} else {
						v44 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariableCancelSleep[1]))
						v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
						*(*int32)(unsafe.Add(mBase, uint32(v45+v39*int32(640))+88)) = v38
					}
					*(*int64)(unsafe.Add(mBase, uint32(v23)+84)) = int64(0)
				} else {
				}
			} else {
				v28 = *(*int32)(unsafe.Add(mBase, uint32(v23)+84))
				if v24 != int32(-1) {
					v33 = v28
					*(*int32)(unsafe.Add(mBase, uint32(v18+v24*int32(640))+84)) = v33
					v38 = v24
					v39 = v33
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v28
					v32 = *(*int32)(unsafe.Add(mBase, uint32(v23)+88))
					v38 = v32
					v39 = v28
				}
				if v39 == int32(-1) {
					*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v38
				} else {
					v44 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariableCancelSleep[1]))
					v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
					*(*int32)(unsafe.Add(mBase, uint32(v45+v39*int32(640))+88)) = v38
				}
				*(*int64)(unsafe.Add(mBase, uint32(v23)+84)) = int64(0)
			}
			v54 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v7))) = v54
			*(*int32)(unsafe.Add(mBase, _c_F_ConditionVariableCancelSleep[0])) = v54
			return
		}
	} else {
		return
	}
}
func F_ConditionalXactLockTableWait(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v87 int32
	_ = v87
	v3 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = int32(17104896)
	*(*int64)(unsafe.Add(mBase, uint32(v7)+4)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
	v18 = F_LockAcquireExtended(m, v7, int32(5), v3, int32(1), v3, l1)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v7 + int32(16)
	return v87
L2:
	;
	return int32(0)
L3:
	;
	if v18 == int32(0) {
		v87 = v3
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v26 = F_LockRelease(m, v7, int32(5), int32(0))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v29 = F_TransactionIdIsInProgress(m, l0)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L2
	} else {
		goto L6
	}
L6:
	;
	if v29 == int32(0) {
		v87 = int32(1)
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v33 = F_SubTransGetTopmostTransaction(m, l0)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L2
	} else {
		goto L8
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = int32(17104896)
	*(*int64)(unsafe.Add(mBase, uint32(v7)+4)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v33
	v40 = int32(0)
	v45 = F_LockAcquireExtended(m, v7, int32(5), v40, int32(1), v40, l1)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L2
	} else {
		goto L9
	}
L9:
	;
	if v45 == int32(0) {
		v87 = v40
		goto L1
	} else {
		goto L10
	}
L10:
	;
	v49 = v33
	goto L11
L11:
	;
	v55 = F_LockRelease(m, v7, int32(5), int32(0))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L2
	} else {
		goto L14
	}
L12:
	;
	v87 = v57 ^ int32(1)
	goto L1
L13:
	;
	goto L12
L14:
	;
	v57 = F_TransactionIdIsInProgress(m, v49)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L2
	} else {
		goto L15
	}
L15:
	;
	if v57 == int32(0) {
		goto L13
	} else {
		goto L16
	}
L16:
	;
	v62 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionalXactLockTableWait[0]))
	if v62 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L2
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	F_pg_usleep(m, int32(1000))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L2
	} else {
		goto L21
	}
L20:
	;
	goto L19
L21:
	;
	v68 = F_SubTransGetTopmostTransaction(m, v49)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L2
	} else {
		goto L22
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = int32(17104896)
	*(*int64)(unsafe.Add(mBase, uint32(v7)+4)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v68
	v76 = int32(0)
	v79 = F_LockAcquireExtended(m, v7, int32(5), v76, int32(1), v76, l1)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L2
	} else {
		goto L23
	}
L23:
	;
	if v79 != 0 {
		v49 = v68
		goto L11
	} else {
		goto L24
	}
L24:
	;
	goto L13
}
func F_CountOtherDBBackends(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v137 int32
	_ = v137
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v189 int32
	_ = v189
	v14 = m.G0
	v16 = v14 - int32(48)
	m.G0 = v16
	v19 = *(*int32)(unsafe.Add(mBase, _c_F_CountOtherDBBackends[0]))
	v34 = int32(0)
	goto L2
L1:
	;
	m.G0 = v16 + int32(48)
	return v189
L2:
	;
	v36 = *(*int32)(unsafe.Add(mBase, _c_F_CountOtherDBBackends[1]))
	if v36 != 0 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v179 = *(*int32)(unsafe.Add(mBase, _c_F_CountOtherDBBackends[2]))
	F_LWLockRelease(m, v179+int32(512))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L7
	} else {
		goto L34
	}
L4:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	v41 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v41
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v41
	v46 = *(*int32)(unsafe.Add(mBase, _c_F_CountOtherDBBackends[2]))
	v50 = F_LWLockAcquire(m, v46+int32(512), int32(1))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L7
	} else {
		goto L9
	}
L7:
	;
	return int32(0)
L8:
	;
	goto L6
L9:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	if int32(0) < v52 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v55 = int32(0)
	v57 = *(*int32)(unsafe.Add(mBase, _c_F_CountOtherDBBackends[3]))
	v63 = v55
	v64 = v55
	v65 = v55
	goto L13
L11:
	;
	goto L12
L12:
	;
	goto L3
L13:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v19+int32(36)+v63<<(uint(int32(2))%32))))
	v79 = v57 + v76*int32(640)
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)+60))
	if v80 != l0 {
		v117 = v64
		v118 = v65
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v124 = *(*int32)(unsafe.Add(mBase, _c_F_CountOtherDBBackends[2]))
	F_LWLockRelease(m, v124+int32(512))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L7
	} else {
		goto L23
	}
L15:
	;
	v120 = v63 + int32(1)
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	if v120 < v121 {
		v63 = v120
		v64 = v117
		v65 = v118
		goto L13
	} else {
		goto L22
	}
L16:
	;
	v83 = *(*int32)(unsafe.Add(mBase, _c_F_CountOtherDBBackends[4]))
	if v79 == v83 {
		v117 = v64
		v118 = v65
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v79)+44))
	if v85 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v88 = int32(1)
	v89 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v89 + v88
	v117 = v64
	v118 = v88
	goto L15
L19:
	;
	goto L20
L20:
	;
	v94 = *(*int32)(unsafe.Add(mBase, _c_F_CountOtherDBBackends[5]))
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v94)+12))
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95+v63))))
	v98 = int32(1)
	v99 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v99 + v98
	if base.B2i32(v97&v98 == int32(0))|base.B2i32(int32(9) < v64) != 0 {
		v117 = v64
		v118 = v98
		goto L15
	} else {
		goto L21
	}
L21:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v79)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v16+v64<<(uint(int32(2))%32)))) = v113
	v117 = v64 + int32(1)
	v118 = v98
	goto L15
L22:
	;
	goto L14
L23:
	;
	if v118 == int32(0) {
		v189 = v118
		goto L1
	} else {
		goto L24
	}
L24:
	;
	v131 = int32(0)
	if v131 < v117 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v137 = v131
	goto L28
L26:
	;
	goto L27
L27:
	;
	F_pg_usleep(m, int32(_a_F_CountOtherDBBackends_0))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L7
	} else {
		goto L32
	}
L28:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v16+v137<<(uint(int32(2))%32))))
	v152 = F_kill(m, v150, int32(15))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L7
	} else {
		goto L30
	}
L29:
	;
	goto L27
L30:
	;
	v155 = v137 + int32(1)
	if v155 != v117 {
		v137 = v155
		goto L28
	} else {
		goto L31
	}
L31:
	;
	goto L29
L32:
	;
	v174 = v34 + int32(1)
	if v174 != int32(50) {
		v34 = v174
		goto L2
	} else {
		goto L33
	}
L33:
	;
	v189 = v118
	goto L1
L34:
	;
	v189 = int32(0)
	goto L1
}
func F_CreateComments(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v97 int32
	_ = v97
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
	var v111 int32
	_ = v111
	v5 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(176)
	m.G0 = v11
	v13 = int32(1)
	if l3 == v5 {
		v33 = v13
		v34 = v5
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v36 = v11 + int32(32)
	F_ScanKeyInit(m, v36, int32(1), int32(3), int32(184), l0)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L4
	} else {
		goto L6
	}
L2:
	;
	v17 = int32(0)
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3))))
	if v18 == v17 {
		v33 = v13
		v34 = v17
		goto L1
	} else {
		goto L3
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = int32(16843009)
	v23 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v23
	*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = l0
	v29 = F_cstring_to_text(m, l3)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+28)) = v29
	v33 = v23
	v34 = int32(1)
	goto L1
L6:
	;
	F_ScanKeyInit(m, v11+int32(80), int32(2), int32(3), int32(184), l1)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L4
	} else {
		goto L7
	}
L7:
	;
	v51 = int32(3)
	F_ScanKeyInit(m, v11+int32(128), v51, v51, int32(65), l2)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L4
	} else {
		goto L8
	}
L8:
	;
	v58 = F_table_open(m, int32(2609), int32(3))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L4
	} else {
		goto L10
	}
L9:
	;
	F_systable_endscan(m, v64)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L4
	} else {
		goto L20
	}
L10:
	;
	v64 = F_systable_beginscan(m, v58, int32(2675), int32(1), int32(0), int32(3), v36)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L4
	} else {
		goto L11
	}
L11:
	;
	v66 = F_systable_getnext(m, v64)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L4
	} else {
		goto L12
	}
L12:
	;
	if v66 == int32(0) {
		v87 = v5
		goto L9
	} else {
		goto L13
	}
L13:
	;
	if v33 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	F_simple_heap_delete(m, v58, v66+int32(4))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L4
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v58)+52))
	v83 = F_heap_modify_tuple(m, v66, v76, v11+int32(16), v11+int32(12), v11+int32(8))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L4
	} else {
		goto L18
	}
L17:
	;
	v87 = v5
	goto L9
L18:
	;
	F_CatalogTupleUpdate(m, v58, v66+int32(4), v83)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L4
	} else {
		goto L19
	}
L19:
	;
	v87 = v83
	goto L9
L20:
	;
	v90 = int32(0)
	if base.B2i32(v34 == v90)|base.B2i32(v87 != v90) == v90 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v58)+52))
	v102 = F_heap_form_tuple(m, v97, v11+int32(16), v11+int32(12))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L4
	} else {
		goto L24
	}
L22:
	;
	v106 = v87
	goto L23
L23:
	;
	if v106 != 0 {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	F_CatalogTupleInsert(m, v58, v102)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L4
	} else {
		goto L25
	}
L25:
	;
	v106 = v102
	goto L23
L26:
	;
	F_pfree(m, v106)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L4
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	F_relation_close(m, v58, int32(0))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L4
	} else {
		goto L30
	}
L29:
	;
	goto L28
L30:
	;
	m.G0 = v11 + int32(176)
	return
}
func F_CreateFakeRelcacheEntry(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v18 int64
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v40 int64
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v11 = F_palloc0(m, int32(420))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = v11 + int32(276)
		v18 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
		*(*int64)(unsafe.Add(mBase, uint32(v11))) = v18
		v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v20
		v22 = int32(112)
		*(*uint8)(unsafe.Add(mBase, uint32(v11)+394)) = uint8(v22)
		*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = int32(-1)
		v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v26
		v33 = F_pg_sprintf(m, v11+int32(280), int32(_a_F_CreateFakeRelcacheEntry_0), v8+int32(16))
		mBase = m.M
		v34 = m.ExcPending
		if v34 != 0 {
			return int32(0)
		} else {
			v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			*(*int32)(unsafe.Add(mBase, uint32(v11)+60)) = v26
			*(*int32)(unsafe.Add(mBase, uint32(v11)+64)) = v35
			v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v38
			v40 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
			*(*int64)(unsafe.Add(mBase, uint32(v8))) = v40
			v43 = F_smgropen(m, v8, int32(-1))
			mBase = m.M
			v44 = m.ExcPending
			if v44 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v43
				m.G0 = v8 + int32(32)
				return v11
			}
		}
	}
}
func F_CreateInitDecodingContext(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32, l4 int32, l5 int32, l6 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v49 int32
	_ = v49
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
	var v60 int64
	_ = v60
	var v62 int64
	_ = v62
	var v64 int64
	_ = v64
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v108 int32
	_ = v108
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v161 int32
	_ = v161
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v199 int32
	_ = v199
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v228 int32
	_ = v228
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v241 int32
	_ = v241
	var v246 int32
	_ = v246
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v257 int32
	_ = v257
	var v262 int32
	_ = v262
	v11 = m.G0
	v13 = v11 - int32(96)
	m.G0 = v13
	F_CheckLogicalDecodingRequirements(m)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v20 = *(*int32)(unsafe.Add(mBase, _c_F_CreateInitDecodingContext[0]))
	if v20 != 0 {
		goto L7
	} else {
		goto L8
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L1
	} else {
		goto L66
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L1
	} else {
		goto L62
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L1
	} else {
		goto L58
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L1
	} else {
		goto L55
	}
L7:
	;
	if l0 == int32(0) {
		goto L6
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L1
	} else {
		goto L52
	}
L10:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v20)+88))
	if v23 == int32(0) {
		goto L5
	} else {
		goto L11
	}
L11:
	;
	v27 = *(*int32)(unsafe.Add(mBase, _c_F_CreateInitDecodingContext[1]))
	if v23 != v27 {
		goto L4
	} else {
		goto L12
	}
L12:
	;
	v30 = *(*int32)(unsafe.Add(mBase, _c_F_CreateInitDecodingContext[2]))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+20))
	goto L13
L13:
	;
	if v31 == int32(2) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v35 = *(*int32)(unsafe.Add(mBase, _c_F_CreateInitDecodingContext[3]))
	if v35 != 0 {
		goto L3
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v39 = F_strncpy(m, v13+int32(4), l0, int32(64))
	mBase = m.M
	v40 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v39)+63)) = uint8(v40)
	goto L18
L17:
	;
	goto L16
L18:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = int32(1)
	if v42 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	F_s_lock(m, v20, int32(_a_F_CreateInitDecodingContext_1), int32(387), int32(_a_F_CreateInitDecodingContext_2))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L1
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v50 = *(*int64)(unsafe.Add(mBase, uint32(v13)+60))
	*(*int64)(unsafe.Add(mBase, uint32(v20)+193)) = v50
	v52 = *(*int64)(unsafe.Add(mBase, uint32(v13)+52))
	*(*int64)(unsafe.Add(mBase, uint32(v20)+185)) = v52
	v54 = *(*int64)(unsafe.Add(mBase, uint32(v13)+44))
	*(*int64)(unsafe.Add(mBase, uint32(v20)+177)) = v54
	v56 = *(*int64)(unsafe.Add(mBase, uint32(v13)+36))
	*(*int64)(unsafe.Add(mBase, uint32(v20)+169)) = v56
	v58 = *(*int64)(unsafe.Add(mBase, uint32(v13)+28))
	*(*int64)(unsafe.Add(mBase, uint32(v20)+161)) = v58
	v60 = *(*int64)(unsafe.Add(mBase, uint32(v13)+20))
	*(*int64)(unsafe.Add(mBase, uint32(v20)+153)) = v60
	v62 = *(*int64)(unsafe.Add(mBase, uint32(v13)+12))
	*(*int64)(unsafe.Add(mBase, uint32(v20)+145)) = v62
	v64 = *(*int64)(unsafe.Add(mBase, uint32(v13)+4))
	*(*int64)(unsafe.Add(mBase, uint32(v20)+137)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = int32(0)
	if l2 == int64(0) {
		goto L24
	} else {
		goto L25
	}
L22:
	;
	goto L21
L23:
	;
	v84 = *(*int32)(unsafe.Add(mBase, _c_F_CreateInitDecodingContext[4]))
	v88 = F_LWLockAcquire(m, v84+int32(_a_F_CreateInitDecodingContext_6), int32(0))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L1
	} else {
		goto L32
	}
L24:
	;
	F_ReplicationSlotReserveWal(m)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L1
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = int32(1)
	if v72 != 0 {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	goto L23
L28:
	;
	F_s_lock(m, v20, int32(_a_F_CreateInitDecodingContext_1), int32(395), int32(_a_F_CreateInitDecodingContext_2))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L1
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v20)+104)) = l2
	goto L23
L31:
	;
	goto L30
L32:
	;
	v91 = *(*int32)(unsafe.Add(mBase, _c_F_CreateInitDecodingContext[4]))
	v95 = F_LWLockAcquire(m, v91+int32(512), int32(0))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	v99 = F_GetOldestSafeDecodingTransactionId(m, l1^int32(1))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = int32(1)
	if v101 != 0 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	F_s_lock(m, v20, int32(_a_F_CreateInitDecodingContext_1), int32(430), int32(_a_F_CreateInitDecodingContext_2))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L1
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+100)) = v99
	*(*int32)(unsafe.Add(mBase, uint32(v20)+20)) = v99
	if l1 != 0 {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	goto L37
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = v99
	goto L41
L40:
	;
	goto L41
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = int32(0)
	F_ReplicationSlotsComputeRequiredXmin(m, int32(1))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	v118 = *(*int32)(unsafe.Add(mBase, _c_F_CreateInitDecodingContext[4]))
	F_LWLockRelease(m, v118+int32(512))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	v124 = *(*int32)(unsafe.Add(mBase, _c_F_CreateInitDecodingContext[4]))
	F_LWLockRelease(m, v124+int32(_a_F_CreateInitDecodingContext_6))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	F_ReplicationSlotMarkDirty(m)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	F_ReplicationSlotSave(m)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	v133 = int32(0)
	v136 = F_StartupDecodingContext(m, v133, l2, v99, l1, v133, int32(1), l3, l4, l5, l6)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	v138 = int32(_a_F_CreateInitDecodingContext_7)
	v139 = *(*int32)(unsafe.Add(mBase, _c_F_CreateInitDecodingContext[5]))
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v136)))
	*(*int32)(unsafe.Add(mBase, _c_F_CreateInitDecodingContext[5])) = v141
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v136)+24))
	if v143 != 0 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v13)+88)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+84)) = int32(_a_F_CreateInitDecodingContext_8)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+80)) = v136
	*(*int32)(unsafe.Add(mBase, uint32(v13)+72)) = int32(993)
	v151 = int32(_a_F_CreateInitDecodingContext_9)
	v152 = *(*int32)(unsafe.Add(mBase, _c_F_CreateInitDecodingContext[6]))
	*(*int32)(unsafe.Add(mBase, _c_F_CreateInitDecodingContext[6])) = v13 + int32(68)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+68)) = v152
	*(*int32)(unsafe.Add(mBase, uint32(v13)+76)) = v13 + int32(80)
	v161 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v136)+164)) = uint8(v161)
	*(*uint8)(unsafe.Add(mBase, uint32(v136)+147)) = uint8(v161)
	m.T0[v143].(func(*base.Module, int32, int32, int32))(m, v136, v136+int32(108), int32(1))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L1
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CreateInitDecodingContext[5])) = v139
	v176 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v136)+145)))
	v177 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+136)))
	v178 = v176 & v177
	*(*uint8)(unsafe.Add(mBase, uint32(v136)+145)) = uint8(v178)
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v136)+12))
	v181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v136)+112)))
	*(*uint8)(unsafe.Add(mBase, uint32(v180)+116)) = uint8(v181)
	m.G0 = v13 + int32(96)
	return v136
L51:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v13)+68))
	*(*int32)(unsafe.Add(mBase, _c_F_CreateInitDecodingContext[6])) = v171
	goto L50
L52:
	;
	F_errmsg_internal(m, int32(_a_F_CreateInitDecodingContext_10), int32(0))
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	F_errfinish(m, int32(_a_F_CreateInitDecodingContext_1), int32(358), int32(_a_F_CreateInitDecodingContext_2))
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L55:
	;
	F_errmsg_internal(m, int32(_a_F_CreateInitDecodingContext_0), int32(0))
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	F_errfinish(m, int32(_a_F_CreateInitDecodingContext_1), int32(361), int32(_a_F_CreateInitDecodingContext_2))
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L58:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	F_errmsg(m, int32(_a_F_CreateInitDecodingContext_3), int32(0))
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	F_errfinish(m, int32(_a_F_CreateInitDecodingContext_1), int32(367), int32(_a_F_CreateInitDecodingContext_2))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L62:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v20 + int32(24)
	F_errmsg(m, int32(_a_F_CreateInitDecodingContext_4), v13)
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	F_errfinish(m, int32(_a_F_CreateInitDecodingContext_1), int32(373), int32(_a_F_CreateInitDecodingContext_2))
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L66:
	;
	F_errcode(m, int32(16777538))
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	F_errmsg(m, int32(_a_F_CreateInitDecodingContext_5), int32(0))
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L1
	} else {
		goto L68
	}
L68:
	;
	F_errfinish(m, int32(_a_F_CreateInitDecodingContext_1), int32(379), int32(_a_F_CreateInitDecodingContext_2))
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L1
	} else {
		goto L69
	}
L69:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_calculate_database_size(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int64
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v58 int64
	_ = v58
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int64
	_ = v89
	var v90 int32
	_ = v90
	var v91 int64
	_ = v91
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v100 int64
	_ = v100
	var v102 int32
	_ = v102
	v6 = m.G0
	v8 = v6 - int32(2128)
	m.G0 = v8
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_calculate_database_size[0]))
	v14 = F_object_aclcheck(m, int32(1262), l0, v12, int64(2048))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = l0
	v32 = v8 + int32(32)
	v37 = F_pg_snprintf(m, v32, int32(1061), int32(_a_F_calculate_database_size_0), v8+int32(16))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L2
	} else {
		goto L9
	}
L2:
	;
	return int64(0)
L3:
	;
	if v14 == int32(0) {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v21 = *(*int32)(unsafe.Add(mBase, _c_F_calculate_database_size[0]))
	v23 = F_has_privs_of_role(m, v21, int32(3375))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L2
	} else {
		goto L5
	}
L5:
	;
	if v23 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v26 = F_get_database_name(m, l0)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L2
	} else {
		goto L7
	}
L7:
	;
	F_aclcheck_error(m, v14, int32(9), v26)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L2
	} else {
		goto L8
	}
L8:
	;
	goto L1
L9:
	;
	v39 = F_db_dir_size(m, v32)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L2
	} else {
		goto L10
	}
L10:
	;
	v42 = v8 + int32(1104)
	v46 = F_pg_snprintf(m, v42, int32(1024), int32(_a_F_calculate_database_size_1), int32(0))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L2
	} else {
		goto L11
	}
L11:
	;
	v48 = F_AllocateDir(m, v42)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L2
	} else {
		goto L13
	}
L12:
	;
	F_FreeDir(m, v48)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L2
	} else {
		goto L35
	}
L13:
	;
	v50 = F_ReadDir(m, v48, v42)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L2
	} else {
		goto L14
	}
L14:
	;
	if v50 == int32(0) {
		v100 = v39
		goto L12
	} else {
		goto L15
	}
L15:
	;
	v55 = v50
	v58 = v39
	goto L16
L16:
	;
	v60 = *(*int32)(unsafe.Add(mBase, _c_F_calculate_database_size[1]))
	if v60 != 0 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v100 = v91
	goto L12
L18:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L2
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+19)))
	if v63 != int32(46) {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	goto L20
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = int32(_a_F_calculate_database_size_2)
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = int32(_a_F_calculate_database_size_1)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v55 + int32(19)
	v84 = v8 + int32(32)
	v87 = F_pg_snprintf(m, v84, int32(1061), int32(_a_F_calculate_database_size_3), v8)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L2
	} else {
		goto L31
	}
L23:
	;
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+20)))
	if v66 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+20)))
	if v67 != int32(46) {
		goto L22
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	v73 = F_ReadDir(m, v48, v8+int32(1104))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L2
	} else {
		goto L29
	}
L27:
	;
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+21)))
	if v70 != 0 {
		goto L22
	} else {
		goto L28
	}
L28:
	;
	goto L26
L29:
	;
	if v73 != 0 {
		v55 = v73
		goto L16
	} else {
		goto L30
	}
L30:
	;
	v100 = v58
	goto L12
L31:
	;
	v89 = F_db_dir_size(m, v84)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L2
	} else {
		goto L32
	}
L32:
	;
	v91 = v89 + v58
	v94 = F_ReadDir(m, v48, v8+int32(1104))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L2
	} else {
		goto L33
	}
L33:
	;
	if v94 != 0 {
		v55 = v94
		v58 = v91
		goto L16
	} else {
		goto L34
	}
L34:
	;
	goto L17
L35:
	;
	m.G0 = v8 + int32(2128)
	return v100
}
func F_cancel_before_shmem_exit(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_cancel_before_shmem_exit[0]))
	if v11 <= int32(0) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v35 = m.ExcPending
		if v35 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = l1
			*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
			F_errmsg_internal(m, int32(_a_F_cancel_before_shmem_exit_0), v8)
			mBase = m.M
			v40 = m.ExcPending
			if v40 != 0 {
				return
			} else {
				F_errfinish(m, int32(_a_F_cancel_before_shmem_exit_1), int32(410), int32(_a_F_cancel_before_shmem_exit_2))
				mBase = m.M
				v45 = m.ExcPending
				if v45 != 0 {
					return
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	} else {
		v15 = v11 << (uint(int32(3)) % 32)
		v18 = *(*int32)(unsafe.Add(mBase, uint32(v15)+uint32(_c_F_cancel_before_shmem_exit[1])))
		if v18 != l0 {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v35 = m.ExcPending
			if v35 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = l1
				*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
				F_errmsg_internal(m, int32(_a_F_cancel_before_shmem_exit_0), v8)
				mBase = m.M
				v40 = m.ExcPending
				if v40 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_cancel_before_shmem_exit_1), int32(410), int32(_a_F_cancel_before_shmem_exit_2))
					mBase = m.M
					v45 = m.ExcPending
					if v45 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v22 = *(*int32)(unsafe.Add(mBase, uint32(v15)+uint32(_c_F_cancel_before_shmem_exit[2])))
			if v22 != l1 {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = l1
					*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
					F_errmsg_internal(m, int32(_a_F_cancel_before_shmem_exit_0), v8)
					mBase = m.M
					v40 = m.ExcPending
					if v40 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_cancel_before_shmem_exit_1), int32(410), int32(_a_F_cancel_before_shmem_exit_2))
						mBase = m.M
						v45 = m.ExcPending
						if v45 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				*(*int32)(unsafe.Add(mBase, _c_F_cancel_before_shmem_exit[0])) = v11 - int32(1)
				m.G0 = v8 + int32(16)
				return
			}
		}
	}
}
func F_carc_cmp(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	v6 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0))))
	v7 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1))))
	if v6 < v7 {
		v18 = int32(-1)
	} else {
		if v7 < v6 {
			v18 = int32(1)
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			if v12 < v13 {
				v18 = int32(-1)
			} else {
				v18 = base.B2i32(v13 < v12)
			}
		}
	}
	return v18
}
func F_casecmp(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	if l2 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v5 = l0
	v6 = l1
	v7 = l2
	goto L4
L2:
	;
	goto L3
L3:
	;
	return int32(0)
L4:
	;
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
	if v9 == v10 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L3
L6:
	;
	v122 = int32(4)
	v127 = v7 - int32(1)
	if v127 != 0 {
		v5 = v5 + v122
		v6 = v6 + v122
		v7 = v127
		goto L4
	} else {
		goto L49
	}
L7:
	;
	v14 = *(*int32)(unsafe.Add(mBase, _c_F_casecmp[0]))
	switch v14 - int32(1) {
	case 0:
		goto L12
	case 1:
		goto L11
	case 2:
		goto L10
	default:
		goto L13
	}
L8:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
	v68 = *(*int32)(unsafe.Add(mBase, _c_F_casecmp[0]))
	switch v68 - int32(1) {
	case 0:
		goto L32
	case 1:
		goto L31
	case 2:
		goto L30
	default:
		goto L33
	}
L9:
	;
	v64 = v60
	goto L8
L10:
	;
	v47 = *(*int32)(unsafe.Add(mBase, _c_F_casecmp[1]))
	if base.Ui32(int32(127)) < base.Ui32(v9) {
		goto L24
	} else {
		goto L25
	}
L11:
	;
	v35 = *(*int32)(unsafe.Add(mBase, _c_F_casecmp[1]))
	if base.Ui32(int32(127)) < base.Ui32(v9) {
		goto L21
	} else {
		goto L22
	}
L12:
	;
	if base.Ui32(v9) <= base.Ui32(int32(127)) {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	if base.Ui32(int32(127)) < base.Ui32(v9) {
		v60 = v9
		goto L9
	} else {
		goto L14
	}
L14:
	;
	v19 = F_pg_tolower(m, v9)
	mBase = m.M
	v64 = v19
	goto L8
L15:
	;
	v31 = v9<<(uint(int32(2))%32) + int32(_a_F_casecmp_0)
	goto L17
L16:
	;
	v26 = F_case_index(m, v9)
	mBase = m.M
	v31 = v26<<(uint(int32(2))%32) + int32(_a_F_casecmp_1)
	goto L17
L17:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
	if v32 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v33 = v32
	goto L20
L19:
	;
	v33 = v9
	goto L20
L20:
	;
	v64 = v33
	goto L8
L21:
	;
	v45 = F_towlower(m, v9)
	mBase = m.M
	v64 = v45
	goto L8
L22:
	;
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+4)))
	if v38&int32(1) == int32(0) {
		goto L21
	} else {
		goto L23
	}
L23:
	;
	v43 = F_pg_tolower(m, v9)
	mBase = m.M
	v64 = v43
	goto L8
L24:
	;
	if base.Ui32(int32(255)) < base.Ui32(v9) {
		v60 = v9
		goto L9
	} else {
		goto L27
	}
L25:
	;
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+4)))
	if v50&int32(1) == int32(0) {
		goto L24
	} else {
		goto L26
	}
L26:
	;
	v55 = F_pg_tolower(m, v9)
	mBase = m.M
	v64 = v55
	goto L8
L27:
	;
	v59 = F_tolower(m, v9)
	mBase = m.M
	v60 = v59
	goto L9
L28:
	;
	if v64 == v118 {
		goto L6
	} else {
		goto L48
	}
L29:
	;
	v118 = v114
	goto L28
L30:
	;
	v101 = *(*int32)(unsafe.Add(mBase, _c_F_casecmp[1]))
	if base.Ui32(int32(127)) < base.Ui32(v65) {
		goto L44
	} else {
		goto L45
	}
L31:
	;
	v89 = *(*int32)(unsafe.Add(mBase, _c_F_casecmp[1]))
	if base.Ui32(int32(127)) < base.Ui32(v65) {
		goto L41
	} else {
		goto L42
	}
L32:
	;
	if base.Ui32(v65) <= base.Ui32(int32(127)) {
		goto L35
	} else {
		goto L36
	}
L33:
	;
	if base.Ui32(int32(127)) < base.Ui32(v65) {
		v114 = v65
		goto L29
	} else {
		goto L34
	}
L34:
	;
	v73 = F_pg_tolower(m, v65)
	mBase = m.M
	v118 = v73
	goto L28
L35:
	;
	v85 = v65<<(uint(int32(2))%32) + int32(_a_F_casecmp_0)
	goto L37
L36:
	;
	v80 = F_case_index(m, v65)
	mBase = m.M
	v85 = v80<<(uint(int32(2))%32) + int32(_a_F_casecmp_1)
	goto L37
L37:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
	if v86 != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v87 = v86
	goto L40
L39:
	;
	v87 = v65
	goto L40
L40:
	;
	v118 = v87
	goto L28
L41:
	;
	v99 = F_towlower(m, v65)
	mBase = m.M
	v118 = v99
	goto L28
L42:
	;
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89)+4)))
	if v92&int32(1) == int32(0) {
		goto L41
	} else {
		goto L43
	}
L43:
	;
	v97 = F_pg_tolower(m, v65)
	mBase = m.M
	v118 = v97
	goto L28
L44:
	;
	if base.Ui32(int32(255)) < base.Ui32(v65) {
		v114 = v65
		goto L29
	} else {
		goto L47
	}
L45:
	;
	v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101)+4)))
	if v104&int32(1) == int32(0) {
		goto L44
	} else {
		goto L46
	}
L46:
	;
	v109 = F_pg_tolower(m, v65)
	mBase = m.M
	v118 = v109
	goto L28
L47:
	;
	v113 = F_tolower(m, v65)
	mBase = m.M
	v114 = v113
	goto L29
L48:
	;
	return int32(1)
L49:
	;
	goto L5
}
func F_catalan_ISO_8859_1_create_env(m *base.Module) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_SN_create_env(m, int32(0), int32(2))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_cclasscvec(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v156 int32
	_ = v156
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v233 int32
	_ = v233
	var v239 int32
	_ = v239
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v253 int32
	_ = v253
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v342 int32
	_ = v342
	var v344 int32
	_ = v344
	var v348 int32
	_ = v348
	var v351 int32
	_ = v351
	var v353 int32
	_ = v353
	v4 = int32(1)
	if l1 == int32(11) {
		v8 = v4
	} else {
		v8 = l1
	}
	if l1 == int32(7) {
		v11 = v4
	} else {
		v11 = v8
	}
	if l2 != 0 {
		v12 = v11
	} else {
		v12 = l1
	}
	switch v12 {
	case 0:
		v23 = F_pg_ctype_get_cache(m, int32(973), int32(0))
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return int32(0)
		} else {
			if v23 == int32(0) {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
				v342 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				if v342 != 0 {
					v344 = v342
				} else {
					v344 = int32(12)
				}
				*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v344
				v348 = int32(0)
			} else {
				v348 = v23
			}
			return v348
		}
	case 1:
		v29 = F_pg_ctype_get_cache(m, int32(974), int32(1))
		mBase = m.M
		v30 = m.ExcPending
		if v30 != 0 {
			return int32(0)
		} else {
			if v29 == int32(0) {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
				v342 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				if v342 != 0 {
					v344 = v342
				} else {
					v344 = int32(12)
				}
				*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v344
				v348 = int32(0)
			} else {
				v348 = v29
			}
			return v348
		}
	case 2:
		v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
		if v39 != 0 {
			v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
			if v40 < int32(0) {
				F_pfree(m, v39)
				mBase = m.M
				v54 = m.ExcPending
				if v54 != 0 {
					return int32(0)
				} else {
					v57 = F_palloc_extended(m, int32(36), int32(2))
					mBase = m.M
					v58 = m.ExcPending
					if v58 != 0 {
						return int32(0)
					} else {
						if v57 == int32(0) {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = int32(0)
							v335 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							if v335 != 0 {
								v337 = v335
							} else {
								v337 = int32(12)
							}
							*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v337
							*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
							v342 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							if v342 != 0 {
								v344 = v342
							} else {
								v344 = int32(12)
							}
							*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v344
							v348 = int32(0)
							return v348
						} else {
							*(*int64)(unsafe.Add(mBase, uint32(v57))) = int64(0)
							*(*int32)(unsafe.Add(mBase, uint32(v57)+24)) = int32(-1)
							*(*int64)(unsafe.Add(mBase, uint32(v57)+12)) = int64(4294967296)
							v68 = v57 + int32(28)
							*(*int32)(unsafe.Add(mBase, uint32(v57)+20)) = v68
							*(*int32)(unsafe.Add(mBase, uint32(v57)+8)) = v68
							*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = v57
							v72 = *(*int32)(unsafe.Add(mBase, uint32(v57)+12))
							v75 = v57
							v77 = v72 << (uint(int32(3)) % 32)
							v78 = *(*int32)(unsafe.Add(mBase, uint32(v75)+20))
							*(*int32)(unsafe.Add(mBase, uint32(v77+v78))) = int32(0)
							v82 = *(*int32)(unsafe.Add(mBase, uint32(v75)+20))
							v83 = *(*int32)(unsafe.Add(mBase, uint32(v75)+12))
							*(*int32)(unsafe.Add(mBase, uint32(v82+v83<<(uint(int32(3))%32))+4)) = int32(127)
							v351 = v75
							v353 = *(*int32)(unsafe.Add(mBase, uint32(v351)+12))
							*(*int32)(unsafe.Add(mBase, uint32(v351)+12)) = v353 + int32(1)
							return v351
						}
					}
				}
			} else {
				v43 = *(*int32)(unsafe.Add(mBase, uint32(v39)+16))
				if v43 <= int32(0) {
					F_pfree(m, v39)
					mBase = m.M
					v54 = m.ExcPending
					if v54 != 0 {
						return int32(0)
					} else {
						v57 = F_palloc_extended(m, int32(36), int32(2))
						mBase = m.M
						v58 = m.ExcPending
						if v58 != 0 {
							return int32(0)
						} else {
							if v57 == int32(0) {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = int32(0)
								v335 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
								if v335 != 0 {
									v337 = v335
								} else {
									v337 = int32(12)
								}
								*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v337
								*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
								v342 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
								if v342 != 0 {
									v344 = v342
								} else {
									v344 = int32(12)
								}
								*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v344
								v348 = int32(0)
								return v348
							} else {
								*(*int64)(unsafe.Add(mBase, uint32(v57))) = int64(0)
								*(*int32)(unsafe.Add(mBase, uint32(v57)+24)) = int32(-1)
								*(*int64)(unsafe.Add(mBase, uint32(v57)+12)) = int64(4294967296)
								v68 = v57 + int32(28)
								*(*int32)(unsafe.Add(mBase, uint32(v57)+20)) = v68
								*(*int32)(unsafe.Add(mBase, uint32(v57)+8)) = v68
								*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = v57
								v72 = *(*int32)(unsafe.Add(mBase, uint32(v57)+12))
								v75 = v57
								v77 = v72 << (uint(int32(3)) % 32)
								v78 = *(*int32)(unsafe.Add(mBase, uint32(v75)+20))
								*(*int32)(unsafe.Add(mBase, uint32(v77+v78))) = int32(0)
								v82 = *(*int32)(unsafe.Add(mBase, uint32(v75)+20))
								v83 = *(*int32)(unsafe.Add(mBase, uint32(v75)+12))
								*(*int32)(unsafe.Add(mBase, uint32(v82+v83<<(uint(int32(3))%32))+4)) = int32(127)
								v351 = v75
								v353 = *(*int32)(unsafe.Add(mBase, uint32(v351)+12))
								*(*int32)(unsafe.Add(mBase, uint32(v351)+12)) = v353 + int32(1)
								return v351
							}
						}
					}
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v39)+24)) = int32(-1)
					v48 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v39)+12)) = v48
					*(*int32)(unsafe.Add(mBase, uint32(v39))) = v48
					v75 = v39
					v77 = v48
					v78 = *(*int32)(unsafe.Add(mBase, uint32(v75)+20))
					*(*int32)(unsafe.Add(mBase, uint32(v77+v78))) = int32(0)
					v82 = *(*int32)(unsafe.Add(mBase, uint32(v75)+20))
					v83 = *(*int32)(unsafe.Add(mBase, uint32(v75)+12))
					*(*int32)(unsafe.Add(mBase, uint32(v82+v83<<(uint(int32(3))%32))+4)) = int32(127)
					v351 = v75
					v353 = *(*int32)(unsafe.Add(mBase, uint32(v351)+12))
					*(*int32)(unsafe.Add(mBase, uint32(v351)+12)) = v353 + int32(1)
					return v351
				}
			}
		} else {
			v57 = F_palloc_extended(m, int32(36), int32(2))
			mBase = m.M
			v58 = m.ExcPending
			if v58 != 0 {
				return int32(0)
			} else {
				if v57 == int32(0) {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = int32(0)
					v335 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					if v335 != 0 {
						v337 = v335
					} else {
						v337 = int32(12)
					}
					*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v337
					*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
					v342 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					if v342 != 0 {
						v344 = v342
					} else {
						v344 = int32(12)
					}
					*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v344
					v348 = int32(0)
					return v348
				} else {
					*(*int64)(unsafe.Add(mBase, uint32(v57))) = int64(0)
					*(*int32)(unsafe.Add(mBase, uint32(v57)+24)) = int32(-1)
					*(*int64)(unsafe.Add(mBase, uint32(v57)+12)) = int64(4294967296)
					v68 = v57 + int32(28)
					*(*int32)(unsafe.Add(mBase, uint32(v57)+20)) = v68
					*(*int32)(unsafe.Add(mBase, uint32(v57)+8)) = v68
					*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = v57
					v72 = *(*int32)(unsafe.Add(mBase, uint32(v57)+12))
					v75 = v57
					v77 = v72 << (uint(int32(3)) % 32)
					v78 = *(*int32)(unsafe.Add(mBase, uint32(v75)+20))
					*(*int32)(unsafe.Add(mBase, uint32(v77+v78))) = int32(0)
					v82 = *(*int32)(unsafe.Add(mBase, uint32(v75)+20))
					v83 = *(*int32)(unsafe.Add(mBase, uint32(v75)+12))
					*(*int32)(unsafe.Add(mBase, uint32(v82+v83<<(uint(int32(3))%32))+4)) = int32(127)
					v351 = v75
					v353 = *(*int32)(unsafe.Add(mBase, uint32(v351)+12))
					*(*int32)(unsafe.Add(mBase, uint32(v351)+12)) = v353 + int32(1)
					return v351
				}
			}
		}
	case 3:
		v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
		if v89 != 0 {
			v90 = *(*int32)(unsafe.Add(mBase, uint32(v89)+4))
			if v90 < int32(2) {
				F_pfree(m, v89)
				mBase = m.M
				v105 = m.ExcPending
				if v105 != 0 {
					return int32(0)
				} else {
					v109 = F_palloc_extended(m, int32(36), int32(2))
					mBase = m.M
					v110 = m.ExcPending
					if v110 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v109)+24)) = int32(-1)
						*(*int64)(unsafe.Add(mBase, uint32(v109)+12)) = int64(0)
						*(*int64)(unsafe.Add(mBase, uint32(v109))) = int64(8589934592)
						*(*int32)(unsafe.Add(mBase, uint32(v109)+20)) = v109 + int32(36)
						*(*int32)(unsafe.Add(mBase, uint32(v109)+8)) = v109 + int32(28)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = v109
						v124 = *(*int32)(unsafe.Add(mBase, uint32(v109)))
						v125 = v109
						v126 = v124
						v127 = int32(1)
						*(*int32)(unsafe.Add(mBase, uint32(v125))) = v126 + v127
						v130 = *(*int32)(unsafe.Add(mBase, uint32(v125)+8))
						v131 = int32(2)
						*(*int32)(unsafe.Add(mBase, uint32(v130+v126<<(uint(v131)%32)))) = int32(9)
						v136 = *(*int32)(unsafe.Add(mBase, uint32(v125)))
						*(*int32)(unsafe.Add(mBase, uint32(v125))) = v136 + v127
						v140 = *(*int32)(unsafe.Add(mBase, uint32(v125)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v140+v136<<(uint(v131)%32)))) = int32(32)
						return v125
					}
				}
			} else {
				v93 = int32(0)
				v94 = *(*int32)(unsafe.Add(mBase, uint32(v89)+16))
				if v94 < v93 {
					F_pfree(m, v89)
					mBase = m.M
					v105 = m.ExcPending
					if v105 != 0 {
						return int32(0)
					} else {
						v109 = F_palloc_extended(m, int32(36), int32(2))
						mBase = m.M
						v110 = m.ExcPending
						if v110 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v109)+24)) = int32(-1)
							*(*int64)(unsafe.Add(mBase, uint32(v109)+12)) = int64(0)
							*(*int64)(unsafe.Add(mBase, uint32(v109))) = int64(8589934592)
							*(*int32)(unsafe.Add(mBase, uint32(v109)+20)) = v109 + int32(36)
							*(*int32)(unsafe.Add(mBase, uint32(v109)+8)) = v109 + int32(28)
							*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = v109
							v124 = *(*int32)(unsafe.Add(mBase, uint32(v109)))
							v125 = v109
							v126 = v124
							v127 = int32(1)
							*(*int32)(unsafe.Add(mBase, uint32(v125))) = v126 + v127
							v130 = *(*int32)(unsafe.Add(mBase, uint32(v125)+8))
							v131 = int32(2)
							*(*int32)(unsafe.Add(mBase, uint32(v130+v126<<(uint(v131)%32)))) = int32(9)
							v136 = *(*int32)(unsafe.Add(mBase, uint32(v125)))
							*(*int32)(unsafe.Add(mBase, uint32(v125))) = v136 + v127
							v140 = *(*int32)(unsafe.Add(mBase, uint32(v125)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v140+v136<<(uint(v131)%32)))) = int32(32)
							return v125
						}
					}
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v89)+24)) = int32(-1)
					v99 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v89)+12)) = v99
					*(*int32)(unsafe.Add(mBase, uint32(v89))) = v99
					v125 = v89
					v126 = v93
					v127 = int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(v125))) = v126 + v127
					v130 = *(*int32)(unsafe.Add(mBase, uint32(v125)+8))
					v131 = int32(2)
					*(*int32)(unsafe.Add(mBase, uint32(v130+v126<<(uint(v131)%32)))) = int32(9)
					v136 = *(*int32)(unsafe.Add(mBase, uint32(v125)))
					*(*int32)(unsafe.Add(mBase, uint32(v125))) = v136 + v127
					v140 = *(*int32)(unsafe.Add(mBase, uint32(v125)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v140+v136<<(uint(v131)%32)))) = int32(32)
					return v125
				}
			}
		} else {
			v109 = F_palloc_extended(m, int32(36), int32(2))
			mBase = m.M
			v110 = m.ExcPending
			if v110 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v109)+24)) = int32(-1)
				*(*int64)(unsafe.Add(mBase, uint32(v109)+12)) = int64(0)
				*(*int64)(unsafe.Add(mBase, uint32(v109))) = int64(8589934592)
				*(*int32)(unsafe.Add(mBase, uint32(v109)+20)) = v109 + int32(36)
				*(*int32)(unsafe.Add(mBase, uint32(v109)+8)) = v109 + int32(28)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = v109
				v124 = *(*int32)(unsafe.Add(mBase, uint32(v109)))
				v125 = v109
				v126 = v124
				v127 = int32(1)
				*(*int32)(unsafe.Add(mBase, uint32(v125))) = v126 + v127
				v130 = *(*int32)(unsafe.Add(mBase, uint32(v125)+8))
				v131 = int32(2)
				*(*int32)(unsafe.Add(mBase, uint32(v130+v126<<(uint(v131)%32)))) = int32(9)
				v136 = *(*int32)(unsafe.Add(mBase, uint32(v125)))
				*(*int32)(unsafe.Add(mBase, uint32(v125))) = v136 + v127
				v140 = *(*int32)(unsafe.Add(mBase, uint32(v125)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v140+v136<<(uint(v131)%32)))) = int32(32)
				return v125
			}
		}
	case 4:
		v147 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
		if v147 != 0 {
			v148 = *(*int32)(unsafe.Add(mBase, uint32(v147)+4))
			if v148 < int32(0) {
				F_pfree(m, v147)
				mBase = m.M
				v162 = m.ExcPending
				if v162 != 0 {
					return int32(0)
				} else {
					v165 = F_palloc_extended(m, int32(44), int32(2))
					mBase = m.M
					v166 = m.ExcPending
					if v166 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v165)+24)) = int32(-1)
						*(*int64)(unsafe.Add(mBase, uint32(v165)+12)) = int64(8589934592)
						*(*int64)(unsafe.Add(mBase, uint32(v165))) = int64(0)
						v174 = v165 + int32(28)
						*(*int32)(unsafe.Add(mBase, uint32(v165)+20)) = v174
						*(*int32)(unsafe.Add(mBase, uint32(v165)+8)) = v174
						*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = v165
						v178 = *(*int32)(unsafe.Add(mBase, uint32(v165)+12))
						v181 = v165
						v183 = v178 << (uint(int32(3)) % 32)
						v184 = *(*int32)(unsafe.Add(mBase, uint32(v181)+20))
						*(*int32)(unsafe.Add(mBase, uint32(v183+v184))) = int32(0)
						v188 = *(*int32)(unsafe.Add(mBase, uint32(v181)+20))
						v189 = *(*int32)(unsafe.Add(mBase, uint32(v181)+12))
						v190 = int32(3)
						*(*int32)(unsafe.Add(mBase, uint32(v188+v189<<(uint(v190)%32))+4)) = int32(31)
						v195 = *(*int32)(unsafe.Add(mBase, uint32(v181)+12))
						v197 = v195 + int32(1)
						*(*int32)(unsafe.Add(mBase, uint32(v181)+12)) = v197
						v199 = *(*int32)(unsafe.Add(mBase, uint32(v181)+20))
						*(*int32)(unsafe.Add(mBase, uint32(v199+v197<<(uint(v190)%32)))) = int32(127)
						v205 = *(*int32)(unsafe.Add(mBase, uint32(v181)+20))
						v206 = *(*int32)(unsafe.Add(mBase, uint32(v181)+12))
						*(*int32)(unsafe.Add(mBase, uint32(v205+v206<<(uint(v190)%32))+4)) = int32(159)
						v351 = v181
						v353 = *(*int32)(unsafe.Add(mBase, uint32(v351)+12))
						*(*int32)(unsafe.Add(mBase, uint32(v351)+12)) = v353 + int32(1)
						return v351
					}
				}
			} else {
				v151 = *(*int32)(unsafe.Add(mBase, uint32(v147)+16))
				if v151 < int32(2) {
					F_pfree(m, v147)
					mBase = m.M
					v162 = m.ExcPending
					if v162 != 0 {
						return int32(0)
					} else {
						v165 = F_palloc_extended(m, int32(44), int32(2))
						mBase = m.M
						v166 = m.ExcPending
						if v166 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v165)+24)) = int32(-1)
							*(*int64)(unsafe.Add(mBase, uint32(v165)+12)) = int64(8589934592)
							*(*int64)(unsafe.Add(mBase, uint32(v165))) = int64(0)
							v174 = v165 + int32(28)
							*(*int32)(unsafe.Add(mBase, uint32(v165)+20)) = v174
							*(*int32)(unsafe.Add(mBase, uint32(v165)+8)) = v174
							*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = v165
							v178 = *(*int32)(unsafe.Add(mBase, uint32(v165)+12))
							v181 = v165
							v183 = v178 << (uint(int32(3)) % 32)
							v184 = *(*int32)(unsafe.Add(mBase, uint32(v181)+20))
							*(*int32)(unsafe.Add(mBase, uint32(v183+v184))) = int32(0)
							v188 = *(*int32)(unsafe.Add(mBase, uint32(v181)+20))
							v189 = *(*int32)(unsafe.Add(mBase, uint32(v181)+12))
							v190 = int32(3)
							*(*int32)(unsafe.Add(mBase, uint32(v188+v189<<(uint(v190)%32))+4)) = int32(31)
							v195 = *(*int32)(unsafe.Add(mBase, uint32(v181)+12))
							v197 = v195 + int32(1)
							*(*int32)(unsafe.Add(mBase, uint32(v181)+12)) = v197
							v199 = *(*int32)(unsafe.Add(mBase, uint32(v181)+20))
							*(*int32)(unsafe.Add(mBase, uint32(v199+v197<<(uint(v190)%32)))) = int32(127)
							v205 = *(*int32)(unsafe.Add(mBase, uint32(v181)+20))
							v206 = *(*int32)(unsafe.Add(mBase, uint32(v181)+12))
							*(*int32)(unsafe.Add(mBase, uint32(v205+v206<<(uint(v190)%32))+4)) = int32(159)
							v351 = v181
							v353 = *(*int32)(unsafe.Add(mBase, uint32(v351)+12))
							*(*int32)(unsafe.Add(mBase, uint32(v351)+12)) = v353 + int32(1)
							return v351
						}
					}
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v147)+24)) = int32(-1)
					v156 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v147)+12)) = v156
					*(*int32)(unsafe.Add(mBase, uint32(v147))) = v156
					v181 = v147
					v183 = v156
					v184 = *(*int32)(unsafe.Add(mBase, uint32(v181)+20))
					*(*int32)(unsafe.Add(mBase, uint32(v183+v184))) = int32(0)
					v188 = *(*int32)(unsafe.Add(mBase, uint32(v181)+20))
					v189 = *(*int32)(unsafe.Add(mBase, uint32(v181)+12))
					v190 = int32(3)
					*(*int32)(unsafe.Add(mBase, uint32(v188+v189<<(uint(v190)%32))+4)) = int32(31)
					v195 = *(*int32)(unsafe.Add(mBase, uint32(v181)+12))
					v197 = v195 + int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(v181)+12)) = v197
					v199 = *(*int32)(unsafe.Add(mBase, uint32(v181)+20))
					*(*int32)(unsafe.Add(mBase, uint32(v199+v197<<(uint(v190)%32)))) = int32(127)
					v205 = *(*int32)(unsafe.Add(mBase, uint32(v181)+20))
					v206 = *(*int32)(unsafe.Add(mBase, uint32(v181)+12))
					*(*int32)(unsafe.Add(mBase, uint32(v205+v206<<(uint(v190)%32))+4)) = int32(159)
					v351 = v181
					v353 = *(*int32)(unsafe.Add(mBase, uint32(v351)+12))
					*(*int32)(unsafe.Add(mBase, uint32(v351)+12)) = v353 + int32(1)
					return v351
				}
			}
		} else {
			v165 = F_palloc_extended(m, int32(44), int32(2))
			mBase = m.M
			v166 = m.ExcPending
			if v166 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v165)+24)) = int32(-1)
				*(*int64)(unsafe.Add(mBase, uint32(v165)+12)) = int64(8589934592)
				*(*int64)(unsafe.Add(mBase, uint32(v165))) = int64(0)
				v174 = v165 + int32(28)
				*(*int32)(unsafe.Add(mBase, uint32(v165)+20)) = v174
				*(*int32)(unsafe.Add(mBase, uint32(v165)+8)) = v174
				*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = v165
				v178 = *(*int32)(unsafe.Add(mBase, uint32(v165)+12))
				v181 = v165
				v183 = v178 << (uint(int32(3)) % 32)
				v184 = *(*int32)(unsafe.Add(mBase, uint32(v181)+20))
				*(*int32)(unsafe.Add(mBase, uint32(v183+v184))) = int32(0)
				v188 = *(*int32)(unsafe.Add(mBase, uint32(v181)+20))
				v189 = *(*int32)(unsafe.Add(mBase, uint32(v181)+12))
				v190 = int32(3)
				*(*int32)(unsafe.Add(mBase, uint32(v188+v189<<(uint(v190)%32))+4)) = int32(31)
				v195 = *(*int32)(unsafe.Add(mBase, uint32(v181)+12))
				v197 = v195 + int32(1)
				*(*int32)(unsafe.Add(mBase, uint32(v181)+12)) = v197
				v199 = *(*int32)(unsafe.Add(mBase, uint32(v181)+20))
				*(*int32)(unsafe.Add(mBase, uint32(v199+v197<<(uint(v190)%32)))) = int32(127)
				v205 = *(*int32)(unsafe.Add(mBase, uint32(v181)+20))
				v206 = *(*int32)(unsafe.Add(mBase, uint32(v181)+12))
				*(*int32)(unsafe.Add(mBase, uint32(v205+v206<<(uint(v190)%32))+4)) = int32(159)
				v351 = v181
				v353 = *(*int32)(unsafe.Add(mBase, uint32(v351)+12))
				*(*int32)(unsafe.Add(mBase, uint32(v351)+12)) = v353 + int32(1)
				return v351
			}
		}
	case 5:
		v214 = F_pg_ctype_get_cache(m, int32(976), int32(5))
		mBase = m.M
		v215 = m.ExcPending
		if v215 != 0 {
			return int32(0)
		} else {
			if v214 == int32(0) {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
				v342 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				if v342 != 0 {
					v344 = v342
				} else {
					v344 = int32(12)
				}
				*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v344
				v348 = int32(0)
			} else {
				v348 = v214
			}
			return v348
		}
	case 6:
		v328 = F_pg_ctype_get_cache(m, int32(981), int32(6))
		mBase = m.M
		v329 = m.ExcPending
		if v329 != 0 {
			return int32(0)
		} else {
			if v328 == int32(0) {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
				v342 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				if v342 != 0 {
					v344 = v342
				} else {
					v344 = int32(12)
				}
				*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v344
				v348 = int32(0)
			} else {
				v348 = v328
			}
			return v348
		}
	case 7:
		v316 = F_pg_ctype_get_cache(m, int32(979), int32(7))
		mBase = m.M
		v317 = m.ExcPending
		if v317 != 0 {
			return int32(0)
		} else {
			if v316 == int32(0) {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
				v342 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				if v342 != 0 {
					v344 = v342
				} else {
					v344 = int32(12)
				}
				*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v344
				v348 = int32(0)
			} else {
				v348 = v316
			}
			return v348
		}
	case 8:
		v15 = F_pg_ctype_get_cache(m, int32(972), int32(8))
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			if v15 == int32(0) {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
				v342 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				if v342 != 0 {
					v344 = v342
				} else {
					v344 = int32(12)
				}
				*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v344
				v348 = int32(0)
			} else {
				v348 = v15
			}
			return v348
		}
	case 9:
		v220 = F_pg_ctype_get_cache(m, int32(977), int32(9))
		mBase = m.M
		v221 = m.ExcPending
		if v221 != 0 {
			return int32(0)
		} else {
			if v220 == int32(0) {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
				v342 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				if v342 != 0 {
					v344 = v342
				} else {
					v344 = int32(12)
				}
				*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v344
				v348 = int32(0)
			} else {
				v348 = v220
			}
			return v348
		}
	case 10:
		v310 = F_pg_ctype_get_cache(m, int32(978), int32(10))
		mBase = m.M
		v311 = m.ExcPending
		if v311 != 0 {
			return int32(0)
		} else {
			if v310 == int32(0) {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
				v342 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				if v342 != 0 {
					v344 = v342
				} else {
					v344 = int32(12)
				}
				*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v344
				v348 = int32(0)
			} else {
				v348 = v310
			}
			return v348
		}
	case 11:
		v322 = F_pg_ctype_get_cache(m, int32(980), int32(11))
		mBase = m.M
		v323 = m.ExcPending
		if v323 != 0 {
			return int32(0)
		} else {
			if v322 == int32(0) {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
				v342 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				if v342 != 0 {
					v344 = v342
				} else {
					v344 = int32(12)
				}
				*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v344
				v348 = int32(0)
			} else {
				v348 = v322
			}
			return v348
		}
	case 12:
		v224 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
		if v224 != 0 {
			v225 = *(*int32)(unsafe.Add(mBase, uint32(v224)+4))
			if v225 < int32(0) {
				F_pfree(m, v224)
				mBase = m.M
				v239 = m.ExcPending
				if v239 != 0 {
					return int32(0)
				} else {
					v242 = F_palloc_extended(m, int32(52), int32(2))
					mBase = m.M
					v243 = m.ExcPending
					if v243 != 0 {
						return int32(0)
					} else {
						if v242 == int32(0) {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = int32(0)
							v335 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							if v335 != 0 {
								v337 = v335
							} else {
								v337 = int32(12)
							}
							*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v337
							*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
							v342 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							if v342 != 0 {
								v344 = v342
							} else {
								v344 = int32(12)
							}
							*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v344
							v348 = int32(0)
							return v348
						} else {
							*(*int64)(unsafe.Add(mBase, uint32(v242))) = int64(0)
							*(*int32)(unsafe.Add(mBase, uint32(v242)+24)) = int32(-1)
							*(*int64)(unsafe.Add(mBase, uint32(v242)+12)) = int64(12884901888)
							v253 = v242 + int32(28)
							*(*int32)(unsafe.Add(mBase, uint32(v242)+20)) = v253
							*(*int32)(unsafe.Add(mBase, uint32(v242)+8)) = v253
							*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = v242
							v257 = *(*int32)(unsafe.Add(mBase, uint32(v242)+12))
							v260 = v242
							v262 = v257 << (uint(int32(3)) % 32)
							v263 = *(*int32)(unsafe.Add(mBase, uint32(v260)+20))
							*(*int32)(unsafe.Add(mBase, uint32(v262+v263))) = int32(48)
							v267 = *(*int32)(unsafe.Add(mBase, uint32(v260)+20))
							v268 = *(*int32)(unsafe.Add(mBase, uint32(v260)+12))
							v269 = int32(3)
							*(*int32)(unsafe.Add(mBase, uint32(v267+v268<<(uint(v269)%32))+4)) = int32(57)
							v274 = *(*int32)(unsafe.Add(mBase, uint32(v260)+12))
							v275 = int32(1)
							v276 = v274 + v275
							*(*int32)(unsafe.Add(mBase, uint32(v260)+12)) = v276
							v278 = *(*int32)(unsafe.Add(mBase, uint32(v260)+20))
							*(*int32)(unsafe.Add(mBase, uint32(v278+v276<<(uint(v269)%32)))) = int32(97)
							v284 = *(*int32)(unsafe.Add(mBase, uint32(v260)+20))
							v285 = *(*int32)(unsafe.Add(mBase, uint32(v260)+12))
							*(*int32)(unsafe.Add(mBase, uint32(v284+v285<<(uint(v269)%32))+4)) = int32(102)
							v291 = *(*int32)(unsafe.Add(mBase, uint32(v260)+12))
							v293 = v291 + v275
							*(*int32)(unsafe.Add(mBase, uint32(v260)+12)) = v293
							v295 = *(*int32)(unsafe.Add(mBase, uint32(v260)+20))
							*(*int32)(unsafe.Add(mBase, uint32(v295+v293<<(uint(v269)%32)))) = int32(65)
							v301 = *(*int32)(unsafe.Add(mBase, uint32(v260)+20))
							v302 = *(*int32)(unsafe.Add(mBase, uint32(v260)+12))
							*(*int32)(unsafe.Add(mBase, uint32(v301+v302<<(uint(v269)%32))+4)) = int32(70)
							v351 = v260
							v353 = *(*int32)(unsafe.Add(mBase, uint32(v351)+12))
							*(*int32)(unsafe.Add(mBase, uint32(v351)+12)) = v353 + int32(1)
							return v351
						}
					}
				}
			} else {
				v228 = *(*int32)(unsafe.Add(mBase, uint32(v224)+16))
				if v228 < int32(3) {
					F_pfree(m, v224)
					mBase = m.M
					v239 = m.ExcPending
					if v239 != 0 {
						return int32(0)
					} else {
						v242 = F_palloc_extended(m, int32(52), int32(2))
						mBase = m.M
						v243 = m.ExcPending
						if v243 != 0 {
							return int32(0)
						} else {
							if v242 == int32(0) {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = int32(0)
								v335 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
								if v335 != 0 {
									v337 = v335
								} else {
									v337 = int32(12)
								}
								*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v337
								*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
								v342 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
								if v342 != 0 {
									v344 = v342
								} else {
									v344 = int32(12)
								}
								*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v344
								v348 = int32(0)
								return v348
							} else {
								*(*int64)(unsafe.Add(mBase, uint32(v242))) = int64(0)
								*(*int32)(unsafe.Add(mBase, uint32(v242)+24)) = int32(-1)
								*(*int64)(unsafe.Add(mBase, uint32(v242)+12)) = int64(12884901888)
								v253 = v242 + int32(28)
								*(*int32)(unsafe.Add(mBase, uint32(v242)+20)) = v253
								*(*int32)(unsafe.Add(mBase, uint32(v242)+8)) = v253
								*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = v242
								v257 = *(*int32)(unsafe.Add(mBase, uint32(v242)+12))
								v260 = v242
								v262 = v257 << (uint(int32(3)) % 32)
								v263 = *(*int32)(unsafe.Add(mBase, uint32(v260)+20))
								*(*int32)(unsafe.Add(mBase, uint32(v262+v263))) = int32(48)
								v267 = *(*int32)(unsafe.Add(mBase, uint32(v260)+20))
								v268 = *(*int32)(unsafe.Add(mBase, uint32(v260)+12))
								v269 = int32(3)
								*(*int32)(unsafe.Add(mBase, uint32(v267+v268<<(uint(v269)%32))+4)) = int32(57)
								v274 = *(*int32)(unsafe.Add(mBase, uint32(v260)+12))
								v275 = int32(1)
								v276 = v274 + v275
								*(*int32)(unsafe.Add(mBase, uint32(v260)+12)) = v276
								v278 = *(*int32)(unsafe.Add(mBase, uint32(v260)+20))
								*(*int32)(unsafe.Add(mBase, uint32(v278+v276<<(uint(v269)%32)))) = int32(97)
								v284 = *(*int32)(unsafe.Add(mBase, uint32(v260)+20))
								v285 = *(*int32)(unsafe.Add(mBase, uint32(v260)+12))
								*(*int32)(unsafe.Add(mBase, uint32(v284+v285<<(uint(v269)%32))+4)) = int32(102)
								v291 = *(*int32)(unsafe.Add(mBase, uint32(v260)+12))
								v293 = v291 + v275
								*(*int32)(unsafe.Add(mBase, uint32(v260)+12)) = v293
								v295 = *(*int32)(unsafe.Add(mBase, uint32(v260)+20))
								*(*int32)(unsafe.Add(mBase, uint32(v295+v293<<(uint(v269)%32)))) = int32(65)
								v301 = *(*int32)(unsafe.Add(mBase, uint32(v260)+20))
								v302 = *(*int32)(unsafe.Add(mBase, uint32(v260)+12))
								*(*int32)(unsafe.Add(mBase, uint32(v301+v302<<(uint(v269)%32))+4)) = int32(70)
								v351 = v260
								v353 = *(*int32)(unsafe.Add(mBase, uint32(v351)+12))
								*(*int32)(unsafe.Add(mBase, uint32(v351)+12)) = v353 + int32(1)
								return v351
							}
						}
					}
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v224)+24)) = int32(-1)
					v233 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v224)+12)) = v233
					*(*int32)(unsafe.Add(mBase, uint32(v224))) = v233
					v260 = v224
					v262 = v233
					v263 = *(*int32)(unsafe.Add(mBase, uint32(v260)+20))
					*(*int32)(unsafe.Add(mBase, uint32(v262+v263))) = int32(48)
					v267 = *(*int32)(unsafe.Add(mBase, uint32(v260)+20))
					v268 = *(*int32)(unsafe.Add(mBase, uint32(v260)+12))
					v269 = int32(3)
					*(*int32)(unsafe.Add(mBase, uint32(v267+v268<<(uint(v269)%32))+4)) = int32(57)
					v274 = *(*int32)(unsafe.Add(mBase, uint32(v260)+12))
					v275 = int32(1)
					v276 = v274 + v275
					*(*int32)(unsafe.Add(mBase, uint32(v260)+12)) = v276
					v278 = *(*int32)(unsafe.Add(mBase, uint32(v260)+20))
					*(*int32)(unsafe.Add(mBase, uint32(v278+v276<<(uint(v269)%32)))) = int32(97)
					v284 = *(*int32)(unsafe.Add(mBase, uint32(v260)+20))
					v285 = *(*int32)(unsafe.Add(mBase, uint32(v260)+12))
					*(*int32)(unsafe.Add(mBase, uint32(v284+v285<<(uint(v269)%32))+4)) = int32(102)
					v291 = *(*int32)(unsafe.Add(mBase, uint32(v260)+12))
					v293 = v291 + v275
					*(*int32)(unsafe.Add(mBase, uint32(v260)+12)) = v293
					v295 = *(*int32)(unsafe.Add(mBase, uint32(v260)+20))
					*(*int32)(unsafe.Add(mBase, uint32(v295+v293<<(uint(v269)%32)))) = int32(65)
					v301 = *(*int32)(unsafe.Add(mBase, uint32(v260)+20))
					v302 = *(*int32)(unsafe.Add(mBase, uint32(v260)+12))
					*(*int32)(unsafe.Add(mBase, uint32(v301+v302<<(uint(v269)%32))+4)) = int32(70)
					v351 = v260
					v353 = *(*int32)(unsafe.Add(mBase, uint32(v351)+12))
					*(*int32)(unsafe.Add(mBase, uint32(v351)+12)) = v353 + int32(1)
					return v351
				}
			}
		} else {
			v242 = F_palloc_extended(m, int32(52), int32(2))
			mBase = m.M
			v243 = m.ExcPending
			if v243 != 0 {
				return int32(0)
			} else {
				if v242 == int32(0) {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = int32(0)
					v335 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					if v335 != 0 {
						v337 = v335
					} else {
						v337 = int32(12)
					}
					*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v337
					*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
					v342 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					if v342 != 0 {
						v344 = v342
					} else {
						v344 = int32(12)
					}
					*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v344
					v348 = int32(0)
					return v348
				} else {
					*(*int64)(unsafe.Add(mBase, uint32(v242))) = int64(0)
					*(*int32)(unsafe.Add(mBase, uint32(v242)+24)) = int32(-1)
					*(*int64)(unsafe.Add(mBase, uint32(v242)+12)) = int64(12884901888)
					v253 = v242 + int32(28)
					*(*int32)(unsafe.Add(mBase, uint32(v242)+20)) = v253
					*(*int32)(unsafe.Add(mBase, uint32(v242)+8)) = v253
					*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = v242
					v257 = *(*int32)(unsafe.Add(mBase, uint32(v242)+12))
					v260 = v242
					v262 = v257 << (uint(int32(3)) % 32)
					v263 = *(*int32)(unsafe.Add(mBase, uint32(v260)+20))
					*(*int32)(unsafe.Add(mBase, uint32(v262+v263))) = int32(48)
					v267 = *(*int32)(unsafe.Add(mBase, uint32(v260)+20))
					v268 = *(*int32)(unsafe.Add(mBase, uint32(v260)+12))
					v269 = int32(3)
					*(*int32)(unsafe.Add(mBase, uint32(v267+v268<<(uint(v269)%32))+4)) = int32(57)
					v274 = *(*int32)(unsafe.Add(mBase, uint32(v260)+12))
					v275 = int32(1)
					v276 = v274 + v275
					*(*int32)(unsafe.Add(mBase, uint32(v260)+12)) = v276
					v278 = *(*int32)(unsafe.Add(mBase, uint32(v260)+20))
					*(*int32)(unsafe.Add(mBase, uint32(v278+v276<<(uint(v269)%32)))) = int32(97)
					v284 = *(*int32)(unsafe.Add(mBase, uint32(v260)+20))
					v285 = *(*int32)(unsafe.Add(mBase, uint32(v260)+12))
					*(*int32)(unsafe.Add(mBase, uint32(v284+v285<<(uint(v269)%32))+4)) = int32(102)
					v291 = *(*int32)(unsafe.Add(mBase, uint32(v260)+12))
					v293 = v291 + v275
					*(*int32)(unsafe.Add(mBase, uint32(v260)+12)) = v293
					v295 = *(*int32)(unsafe.Add(mBase, uint32(v260)+20))
					*(*int32)(unsafe.Add(mBase, uint32(v295+v293<<(uint(v269)%32)))) = int32(65)
					v301 = *(*int32)(unsafe.Add(mBase, uint32(v260)+20))
					v302 = *(*int32)(unsafe.Add(mBase, uint32(v260)+12))
					*(*int32)(unsafe.Add(mBase, uint32(v301+v302<<(uint(v269)%32))+4)) = int32(70)
					v351 = v260
					v353 = *(*int32)(unsafe.Add(mBase, uint32(v351)+12))
					*(*int32)(unsafe.Add(mBase, uint32(v351)+12)) = v353 + int32(1)
					return v351
				}
			}
		}
	case 13:
		v35 = F_pg_ctype_get_cache(m, int32(975), int32(13))
		mBase = m.M
		v36 = m.ExcPending
		if v36 != 0 {
			return int32(0)
		} else {
			if v35 == int32(0) {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
				v342 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				if v342 != 0 {
					v344 = v342
				} else {
					v344 = int32(12)
				}
				*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v344
				v348 = int32(0)
			} else {
				v348 = v35
			}
			return v348
		}
	default:
		*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
		v342 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		if v342 != 0 {
			v344 = v342
		} else {
			v344 = int32(12)
		}
		*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v344
		v348 = int32(0)
		return v348
	}
}
func F_char_decrement(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v16 int32
	_ = v16
	v5 = l1 & int32(255)
	v6 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(base.B2i32(v5 == v6))
	v9 = int32(24)
	if v5 != 0 {
		v16 = (l1<<(uint(v9)%32) - int32(16777216)) >> (uint(v9) % 32)
	} else {
		v16 = v6
	}
	return v16
}
func F_charclasscomplement(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v102 int32
	_ = v102
	var v109 int32
	_ = v109
	var v110 int64
	_ = v110
	var v116 int32
	_ = v116
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v184 int32
	_ = v184
	var v191 int32
	_ = v191
	var v192 int64
	_ = v192
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v214 int32
	_ = v214
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	v8 = F_newstate(m, v7)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		if v10 != 0 {
			return
		} else {
			v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v12 | int32(1024)
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			v19 = F_cclasscvec(m, l0, l1, v16&int32(8))
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return
			} else {
				v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				if v21 != 0 {
					return
				} else {
					F_subcolorcvec(m, l0, v19, v8, v8)
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
						return
					} else {
						v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						if v24 != 0 {
							return
						} else {
							v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
							v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
							F_okcolors(m, v25, v26)
							mBase = m.M
							v28 = m.ExcPending
							if v28 != 0 {
								return
							} else {
								v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
								if v29 != 0 {
									return
								} else {
									v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
									v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
									F_colorcomplement(m, v30, v31, int32(112), v8, l2, l3)
									mBase = m.M
									v34 = m.ExcPending
									if v34 != 0 {
										return
									} else {
										v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
										if v35 != 0 {
										} else {
											v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
											for {
												v43 = *(*int32)(unsafe.Add(mBase, uint32(v8)+16))
												if v43 != 0 {
													v48 = *(*int32)(unsafe.Add(mBase, uint32(v43)+12))
													v49 = *(*int32)(unsafe.Add(mBase, uint32(v43)+8))
													v50 = int32(*(*int16)(unsafe.Add(mBase, uint32(v43)+4)))
													if v50 < int32(0) {
													} else {
														v53 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
														v55 = v53 - int32(97)
														if base.B2i32(base.Ui32(int32(17)) < base.Ui32(v55))|base.B2i32(int32(1)<<(uint(v55)%32)&int32(_a_F_charclasscomplement_0) == int32(0)) != 0 {
														} else {
															v65 = *(*int32)(unsafe.Add(mBase, uint32(v36)+80))
															if v65 != 0 {
															} else {
																v66 = *(*int32)(unsafe.Add(mBase, uint32(v43)+36))
																if v66 == int32(0) {
																	v69 = *(*int32)(unsafe.Add(mBase, uint32(v36)+52))
																	v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)+20))
																	v74 = *(*int32)(unsafe.Add(mBase, uint32(v43)+32))
																	*(*int32)(unsafe.Add(mBase, uint32(v70+v50*int32(24))+12)) = v74
																	v78 = v74
																} else {
																	v76 = *(*int32)(unsafe.Add(mBase, uint32(v43)+32))
																	*(*int32)(unsafe.Add(mBase, uint32(v66)+32)) = v76
																	v78 = v76
																}
																if v78 != 0 {
																	*(*int32)(unsafe.Add(mBase, uint32(v78)+36)) = v66
																} else {
																}
																*(*int64)(unsafe.Add(mBase, uint32(v43)+32)) = int64(0)
															}
														}
													}
													v84 = *(*int32)(unsafe.Add(mBase, uint32(v43)+16))
													v85 = *(*int32)(unsafe.Add(mBase, uint32(v43)+20))
													if v85 == int32(0) {
														*(*int32)(unsafe.Add(mBase, uint32(v49)+20)) = v84
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v85)+16)) = v84
													}
													if v84 != 0 {
														*(*int32)(unsafe.Add(mBase, uint32(v84)+20)) = v85
													} else {
													}
													v91 = *(*int32)(unsafe.Add(mBase, uint32(v49)+12))
													*(*int32)(unsafe.Add(mBase, uint32(v49)+12)) = v91 - int32(1)
													v95 = *(*int32)(unsafe.Add(mBase, uint32(v43)+24))
													v96 = *(*int32)(unsafe.Add(mBase, uint32(v43)+28))
													if v96 == int32(0) {
														*(*int32)(unsafe.Add(mBase, uint32(v48)+16)) = v95
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v96)+24)) = v95
													}
													if v95 != 0 {
														*(*int32)(unsafe.Add(mBase, uint32(v95)+28)) = v96
													} else {
													}
													v102 = *(*int32)(unsafe.Add(mBase, uint32(v48)+8))
													*(*int32)(unsafe.Add(mBase, uint32(v48)+8)) = v102 - int32(1)
													*(*int32)(unsafe.Add(mBase, uint32(v43))) = int32(0)
													v109 = v43 + int32(8)
													v110 = int64(0)
													*(*int64)(unsafe.Add(mBase, uint32(v109)+16)) = v110
													*(*int64)(unsafe.Add(mBase, uint32(v109)+8)) = v110
													*(*int64)(unsafe.Add(mBase, uint32(v109))) = v110
													v116 = *(*int32)(unsafe.Add(mBase, uint32(v36)+32))
													*(*int32)(unsafe.Add(mBase, uint32(v43)+16)) = v116
													*(*int32)(unsafe.Add(mBase, uint32(v36)+32)) = v43
													continue
												} else {
													break
												}
												break
											}
											for {
												v125 = *(*int32)(unsafe.Add(mBase, uint32(v8)+20))
												if v125 != 0 {
													v130 = *(*int32)(unsafe.Add(mBase, uint32(v125)+12))
													v131 = *(*int32)(unsafe.Add(mBase, uint32(v125)+8))
													v132 = int32(*(*int16)(unsafe.Add(mBase, uint32(v125)+4)))
													if v132 < int32(0) {
													} else {
														v135 = *(*int32)(unsafe.Add(mBase, uint32(v125)))
														v137 = v135 - int32(97)
														if base.B2i32(base.Ui32(int32(17)) < base.Ui32(v137))|base.B2i32(int32(1)<<(uint(v137)%32)&int32(_a_F_charclasscomplement_0) == int32(0)) != 0 {
														} else {
															v147 = *(*int32)(unsafe.Add(mBase, uint32(v36)+80))
															if v147 != 0 {
															} else {
																v148 = *(*int32)(unsafe.Add(mBase, uint32(v125)+36))
																if v148 == int32(0) {
																	v151 = *(*int32)(unsafe.Add(mBase, uint32(v36)+52))
																	v152 = *(*int32)(unsafe.Add(mBase, uint32(v151)+20))
																	v156 = *(*int32)(unsafe.Add(mBase, uint32(v125)+32))
																	*(*int32)(unsafe.Add(mBase, uint32(v152+v132*int32(24))+12)) = v156
																	v160 = v156
																} else {
																	v158 = *(*int32)(unsafe.Add(mBase, uint32(v125)+32))
																	*(*int32)(unsafe.Add(mBase, uint32(v148)+32)) = v158
																	v160 = v158
																}
																if v160 != 0 {
																	*(*int32)(unsafe.Add(mBase, uint32(v160)+36)) = v148
																} else {
																}
																*(*int64)(unsafe.Add(mBase, uint32(v125)+32)) = int64(0)
															}
														}
													}
													v166 = *(*int32)(unsafe.Add(mBase, uint32(v125)+16))
													v167 = *(*int32)(unsafe.Add(mBase, uint32(v125)+20))
													if v167 == int32(0) {
														*(*int32)(unsafe.Add(mBase, uint32(v131)+20)) = v166
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v167)+16)) = v166
													}
													if v166 != 0 {
														*(*int32)(unsafe.Add(mBase, uint32(v166)+20)) = v167
													} else {
													}
													v173 = *(*int32)(unsafe.Add(mBase, uint32(v131)+12))
													*(*int32)(unsafe.Add(mBase, uint32(v131)+12)) = v173 - int32(1)
													v177 = *(*int32)(unsafe.Add(mBase, uint32(v125)+24))
													v178 = *(*int32)(unsafe.Add(mBase, uint32(v125)+28))
													if v178 == int32(0) {
														*(*int32)(unsafe.Add(mBase, uint32(v130)+16)) = v177
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v178)+24)) = v177
													}
													if v177 != 0 {
														*(*int32)(unsafe.Add(mBase, uint32(v177)+28)) = v178
													} else {
													}
													v184 = *(*int32)(unsafe.Add(mBase, uint32(v130)+8))
													*(*int32)(unsafe.Add(mBase, uint32(v130)+8)) = v184 - int32(1)
													*(*int32)(unsafe.Add(mBase, uint32(v125))) = int32(0)
													v191 = v125 + int32(8)
													v192 = int64(0)
													*(*int64)(unsafe.Add(mBase, uint32(v191)+16)) = v192
													*(*int64)(unsafe.Add(mBase, uint32(v191)+8)) = v192
													*(*int64)(unsafe.Add(mBase, uint32(v191))) = v192
													v198 = *(*int32)(unsafe.Add(mBase, uint32(v36)+32))
													*(*int32)(unsafe.Add(mBase, uint32(v125)+16)) = v198
													*(*int32)(unsafe.Add(mBase, uint32(v36)+32)) = v125
													continue
												} else {
													break
												}
												break
											}
											v201 = int32(0)
											*(*uint8)(unsafe.Add(mBase, uint32(v8)+4)) = uint8(v201)
											*(*int32)(unsafe.Add(mBase, uint32(v8))) = int32(-1)
											v205 = *(*int32)(unsafe.Add(mBase, uint32(v8)+32))
											v206 = *(*int32)(unsafe.Add(mBase, uint32(v8)+28))
											if v206 != 0 {
												*(*int32)(unsafe.Add(mBase, uint32(v206)+32)) = v205
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v36)+24)) = v205
											}
											v209 = *(*int32)(unsafe.Add(mBase, uint32(v8)+28))
											if v205 != 0 {
												*(*int32)(unsafe.Add(mBase, uint32(v205)+28)) = v209
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v36)+20)) = v209
											}
											*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = int32(0)
											v214 = *(*int32)(unsafe.Add(mBase, uint32(v36)+28))
											*(*int32)(unsafe.Add(mBase, uint32(v8)+28)) = v214
											*(*int32)(unsafe.Add(mBase, uint32(v36)+28)) = v8
										}
										return
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
func F_charhashfast(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	v2 = base.I32_extend8_s(l0)
	v3 = int32(16)
	v7 = (int32(base.Ui32(v2)>>(uint(v3)%32)) ^ v2) * int32(-2048144789)
	v12 = (int32(base.Ui32(v7)>>(uint(int32(13))%32)) ^ v7) * int32(-1028477387)
	return int32(base.Ui32(v12)>>(uint(v3)%32)) ^ v12
}
func F_chartoi4(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0)+20)))
	return v2
}
func F_checkExprHasSubLink_walker(m *base.Module, l0 int32, l1 int32) int32 {
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v5 = Fn13866(m, l0, l1, int32(1047), int32(22))
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
func F_checkTargetlistEntrySQL92(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
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
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	if base.Ui32(l2-int32(20)) < base.Ui32(int32(2)) {
		m.G0 = v7 + int32(32)
		return
	} else {
		if l2 == int32(19) {
			v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+92)))
			if v15 == int32(1) {
				v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
				v20 = F_contain_aggs_of_level(m, v18, int32(0))
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return
				} else {
					if v20 != 0 {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v80 = m.ExcPending
						if v80 != 0 {
							return
						} else {
							F_errcode(m, int32(50364548))
							mBase = m.M
							v83 = m.ExcPending
							if v83 != 0 {
								return
							} else {
								v89 = *(*int32)(unsafe.Add(mBase, _c_F_checkTargetlistEntrySQL92[0]))
								*(*int32)(unsafe.Add(mBase, uint32(v7))) = v89
								F_errmsg(m, int32(_a_F_checkTargetlistEntrySQL92_0), v7)
								mBase = m.M
								v95 = m.ExcPending
								if v95 != 0 {
									return
								} else {
									v96 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
									v98 = F_locate_agg_of_level(m, v96, int32(0))
									mBase = m.M
									v99 = m.ExcPending
									if v99 != 0 {
										return
									} else {
										F_parser_errposition(m, l0, v98)
										mBase = m.M
										v101 = m.ExcPending
										if v101 != 0 {
											return
										} else {
											F_errfinish(m, int32(_a_F_checkTargetlistEntrySQL92_1), int32(1965), int32(_a_F_checkTargetlistEntrySQL92_2))
											mBase = m.M
											v106 = m.ExcPending
											if v106 != 0 {
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
					} else {
						v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+93)))
						if v22 != int32(1) {
							m.G0 = v7 + int32(32)
							return
						} else {
							v25 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
							v26 = F_contain_windowfuncs(m, v25)
							mBase = m.M
							v27 = m.ExcPending
							if v27 != 0 {
								return
							} else {
								if v26 == int32(0) {
									m.G0 = v7 + int32(32)
									return
								} else {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v33 = m.ExcPending
									if v33 != 0 {
										return
									} else {
										F_errcode(m, int32(_a_F_checkTargetlistEntrySQL92_3))
										mBase = m.M
										v36 = m.ExcPending
										if v36 != 0 {
											return
										} else {
											v42 = *(*int32)(unsafe.Add(mBase, _c_F_checkTargetlistEntrySQL92[0]))
											*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v42
											F_errmsg(m, int32(_a_F_checkTargetlistEntrySQL92_4), v7+int32(16))
											mBase = m.M
											v50 = m.ExcPending
											if v50 != 0 {
												return
											} else {
												v51 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
												v52 = F_locate_windowfunc(m, v51)
												mBase = m.M
												v53 = m.ExcPending
												if v53 != 0 {
													return
												} else {
													F_parser_errposition(m, l0, v52)
													mBase = m.M
													v55 = m.ExcPending
													if v55 != 0 {
														return
													} else {
														F_errfinish(m, int32(_a_F_checkTargetlistEntrySQL92_1), int32(1974), int32(_a_F_checkTargetlistEntrySQL92_2))
														mBase = m.M
														v60 = m.ExcPending
														if v60 != 0 {
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
						}
					}
				}
			} else {
				v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+93)))
				if v22 != int32(1) {
					m.G0 = v7 + int32(32)
					return
				} else {
					v25 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
					v26 = F_contain_windowfuncs(m, v25)
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return
					} else {
						if v26 == int32(0) {
							m.G0 = v7 + int32(32)
							return
						} else {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v33 = m.ExcPending
							if v33 != 0 {
								return
							} else {
								F_errcode(m, int32(_a_F_checkTargetlistEntrySQL92_3))
								mBase = m.M
								v36 = m.ExcPending
								if v36 != 0 {
									return
								} else {
									v42 = *(*int32)(unsafe.Add(mBase, _c_F_checkTargetlistEntrySQL92[0]))
									*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v42
									F_errmsg(m, int32(_a_F_checkTargetlistEntrySQL92_4), v7+int32(16))
									mBase = m.M
									v50 = m.ExcPending
									if v50 != 0 {
										return
									} else {
										v51 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
										v52 = F_locate_windowfunc(m, v51)
										mBase = m.M
										v53 = m.ExcPending
										if v53 != 0 {
											return
										} else {
											F_parser_errposition(m, l0, v52)
											mBase = m.M
											v55 = m.ExcPending
											if v55 != 0 {
												return
											} else {
												F_errfinish(m, int32(_a_F_checkTargetlistEntrySQL92_1), int32(1974), int32(_a_F_checkTargetlistEntrySQL92_2))
												mBase = m.M
												v60 = m.ExcPending
												if v60 != 0 {
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
				}
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v64 = m.ExcPending
			if v64 != 0 {
				return
			} else {
				F_errmsg_internal(m, int32(_a_F_checkTargetlistEntrySQL92_5), int32(0))
				mBase = m.M
				v68 = m.ExcPending
				if v68 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_checkTargetlistEntrySQL92_1), int32(1983), int32(_a_F_checkTargetlistEntrySQL92_2))
					mBase = m.M
					v73 = m.ExcPending
					if v73 != 0 {
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
func F_check_agglevels_and_constraints(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
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
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v266 int32
	_ = v266
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v287 int32
	_ = v287
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v298 int32
	_ = v298
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v326 int32
	_ = v326
	var v330 int32
	_ = v330
	var v333 int32
	_ = v333
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v342 int32
	_ = v342
	var v347 int32
	_ = v347
	var v351 int32
	_ = v351
	var v354 int32
	_ = v354
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v363 int32
	_ = v363
	var v368 int32
	_ = v368
	var v372 int32
	_ = v372
	var v375 int32
	_ = v375
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v388 int32
	_ = v388
	var v390 int32
	_ = v390
	var v395 int32
	_ = v395
	v3 = int32(0)
	v12 = m.G0
	v14 = v12 - int32(80)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v16 == int32(9) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l1+v29)))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+76)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v14)+68)) = int64(4294967295)
	*(*int64)(unsafe.Add(mBase, uint32(v14)+60)) = int64(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+56)) = l0
	v45 = v14 + int32(56)
	v46 = F_check_agg_arguments_walker(m, v36, v45)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v29 = int32(68)
	v30 = v20
	v31 = v19
	v32 = int32(52)
	v33 = l1 + int32(32)
	goto L1
L3:
	;
	goto L4
L4:
	;
	v29 = int32(20)
	v30 = v3
	v31 = v3
	v32 = int32(16)
	v33 = l1 + int32(4)
	goto L1
L5:
	;
	return
L6:
	;
	v48 = F_check_agg_arguments_walker(m, v31, v45)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v14)+64))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v14)+60))
	if v51 < int32(0) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v54 = int32(0)
	if v54 < v50 {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	if v51 < v50 {
		goto L14
	} else {
		goto L15
	}
L10:
	;
	if v50 == v63 {
		goto L20
	} else {
		goto L21
	}
L11:
	;
	v57 = v50
	goto L13
L12:
	;
	v57 = v54
	goto L13
L13:
	;
	v63 = v57
	goto L10
L14:
	;
	v59 = v51
	goto L16
L15:
	;
	v59 = v50
	goto L16
L16:
	;
	if v50 < int32(0) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v62 = v51
	goto L19
L18:
	;
	v62 = v59
	goto L19
L19:
	;
	v63 = v62
	goto L10
L20:
	;
	v65 = F_locate_agg_of_level(m, v36, v50)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L5
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v14)+68))
	v91 = int32(0)
	if base.B2i32(v91 <= v90)&base.B2i32(v90 < v63) == v91 {
		goto L36
	} else {
		goto L37
	}
L23:
	;
	if v65 < int32(0) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v69 = F_locate_agg_of_level(m, v31, v50)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L5
	} else {
		goto L27
	}
L25:
	;
	v71 = v65
	goto L26
L26:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L5
	} else {
		goto L28
	}
L27:
	;
	v71 = v69
	goto L26
L28:
	;
	F_errcode(m, int32(50364548))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L5
	} else {
		goto L29
	}
L29:
	;
	F_errmsg(m, int32(_a_F_check_agglevels_and_constraints_0), int32(0))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L5
	} else {
		goto L30
	}
L30:
	;
	F_parser_errposition(m, l0, v71)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L5
	} else {
		goto L31
	}
L31:
	;
	F_errfinish(m, int32(_a_F_check_agglevels_and_constraints_1), int32(693), int32(_a_F_check_agglevels_and_constraints_2))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L5
	} else {
		goto L32
	}
L32:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L33:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L5
	} else {
		goto L120
	}
L34:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L5
	} else {
		goto L114
	}
L35:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L5
	} else {
		goto L108
	}
L36:
	;
	if v30 == int32(0) {
		goto L39
	} else {
		goto L40
	}
L37:
	;
	goto L38
L38:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L5
	} else {
		goto L102
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1+v32))) = v63
	if v63 <= int32(0) {
		v184 = l0
		goto L46
	} else {
		goto L47
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+68)) = int32(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v14)+60)) = int64(-1)
	v105 = F_check_agg_arguments_walker(m, v30, v14+int32(56))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L5
	} else {
		goto L41
	}
L41:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v14)+60))
	if base.B2i32(int32(0) <= v107)&base.B2i32(v107 < v63) != 0 {
		goto L35
	} else {
		goto L42
	}
L42:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v14)+64))
	if base.B2i32(int32(0) <= v112)&base.B2i32(v112 <= v63) != 0 {
		goto L34
	} else {
		goto L43
	}
L43:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v14)+68))
	if v117 < int32(0) {
		goto L39
	} else {
		goto L44
	}
L44:
	;
	if v117 < v63 {
		goto L33
	} else {
		goto L45
	}
L45:
	;
	goto L39
L46:
	;
	v195 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v184)+92)) = uint8(v195)
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v184)+68))
	switch v199 - int32(2) {
	case 0, 1:
		v242 = int32(_a_F_check_agglevels_and_constraints_3)
		v243 = int32(_a_F_check_agglevels_and_constraints_4)
		goto L61
	case 2:
		goto L81
	case 3:
		goto L80
	case 4, 6, 13, 14, 15, 17, 20, 21, 22, 23, 24, 25, 42:
		goto L60
	default:
		goto L59
	case 9:
		goto L78
	case 10:
		goto L77
	case 11:
		goto L76
	case 16:
		goto L75
	case 26, 27:
		goto L74
	case 28, 29:
		goto L73
	case 30:
		goto L72
	case 31:
		goto L71
	case 32:
		goto L70
	case 33:
		goto L69
	case 34:
		goto L68
	case 35:
		goto L67
	case 36:
		goto L79
	case 37:
		goto L66
	case 38:
		goto L65
	case 39:
		goto L63
	case 40:
		goto L62
	case 41:
		goto L64
	}
L47:
	;
	v127 = v63 & int32(7)
	if v127 == int32(0) {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	if base.Ui32(v63) < base.Ui32(int32(8)) {
		v184 = v148
		goto L46
	} else {
		goto L55
	}
L49:
	;
	v148 = l0
	v149 = v63
	goto L48
L50:
	;
	goto L51
L51:
	;
	v131 = l0
	v132 = v63
	v133 = int32(0)
	goto L52
L52:
	;
	v142 = int32(1)
	v143 = v132 - v142
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v131)))
	v146 = v133 + v142
	if v146 != v127 {
		v131 = v144
		v132 = v143
		v133 = v146
		goto L52
	} else {
		goto L54
	}
L53:
	;
	v148 = v144
	v149 = v143
	goto L48
L54:
	;
	goto L53
L55:
	;
	v161 = v148
	v162 = v149
	goto L56
L56:
	;
	v172 = int32(8)
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v161)))
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v174)))
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v175)))
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v176)))
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v177)))
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v178)))
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v179)))
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v180)))
	if v172 < v162 {
		v161 = v181
		v162 = v162 - v172
		goto L56
	} else {
		goto L58
	}
L57:
	;
	v184 = v181
	goto L46
L58:
	;
	goto L57
L59:
	;
	m.G0 = v14 + int32(80)
	return
L60:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L5
	} else {
		goto L90
	}
L61:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L5
	} else {
		goto L82
	}
L62:
	;
	v242 = int32(_a_F_check_agglevels_and_constraints_46)
	v243 = int32(_a_F_check_agglevels_and_constraints_47)
	goto L61
L63:
	;
	v242 = int32(_a_F_check_agglevels_and_constraints_44)
	v243 = int32(_a_F_check_agglevels_and_constraints_45)
	goto L61
L64:
	;
	v242 = int32(_a_F_check_agglevels_and_constraints_48)
	v243 = int32(_a_F_check_agglevels_and_constraints_49)
	goto L61
L65:
	;
	v242 = int32(_a_F_check_agglevels_and_constraints_42)
	v243 = int32(_a_F_check_agglevels_and_constraints_43)
	goto L61
L66:
	;
	v242 = int32(_a_F_check_agglevels_and_constraints_40)
	v243 = int32(_a_F_check_agglevels_and_constraints_41)
	goto L61
L67:
	;
	v242 = int32(_a_F_check_agglevels_and_constraints_36)
	v243 = int32(_a_F_check_agglevels_and_constraints_37)
	goto L61
L68:
	;
	v242 = int32(_a_F_check_agglevels_and_constraints_34)
	v243 = int32(_a_F_check_agglevels_and_constraints_35)
	goto L61
L69:
	;
	v242 = int32(_a_F_check_agglevels_and_constraints_32)
	v243 = int32(_a_F_check_agglevels_and_constraints_33)
	goto L61
L70:
	;
	v242 = int32(_a_F_check_agglevels_and_constraints_30)
	v243 = int32(_a_F_check_agglevels_and_constraints_31)
	goto L61
L71:
	;
	v242 = int32(_a_F_check_agglevels_and_constraints_28)
	v243 = int32(_a_F_check_agglevels_and_constraints_29)
	goto L61
L72:
	;
	v242 = int32(_a_F_check_agglevels_and_constraints_26)
	v243 = int32(_a_F_check_agglevels_and_constraints_27)
	goto L61
L73:
	;
	v242 = int32(_a_F_check_agglevels_and_constraints_24)
	v243 = int32(_a_F_check_agglevels_and_constraints_25)
	goto L61
L74:
	;
	v242 = int32(_a_F_check_agglevels_and_constraints_22)
	v243 = int32(_a_F_check_agglevels_and_constraints_23)
	goto L61
L75:
	;
	v242 = int32(_a_F_check_agglevels_and_constraints_20)
	v243 = int32(_a_F_check_agglevels_and_constraints_21)
	goto L61
L76:
	;
	v242 = int32(_a_F_check_agglevels_and_constraints_18)
	v243 = int32(_a_F_check_agglevels_and_constraints_19)
	goto L61
L77:
	;
	v242 = int32(_a_F_check_agglevels_and_constraints_16)
	v243 = int32(_a_F_check_agglevels_and_constraints_17)
	goto L61
L78:
	;
	v242 = int32(_a_F_check_agglevels_and_constraints_14)
	v243 = int32(_a_F_check_agglevels_and_constraints_15)
	goto L61
L79:
	;
	v242 = int32(_a_F_check_agglevels_and_constraints_38)
	v243 = int32(_a_F_check_agglevels_and_constraints_39)
	goto L61
L80:
	;
	v242 = int32(_a_F_check_agglevels_and_constraints_9)
	v243 = int32(_a_F_check_agglevels_and_constraints_10)
	goto L61
L81:
	;
	v242 = int32(_a_F_check_agglevels_and_constraints_7)
	v243 = int32(_a_F_check_agglevels_and_constraints_8)
	goto L61
L82:
	;
	F_errcode(m, int32(50364548))
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L5
	} else {
		goto L83
	}
L83:
	;
	if v16 == int32(9) {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	v253 = v242
	goto L86
L85:
	;
	v253 = v243
	goto L86
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v253
	F_errmsg_internal(m, int32(_a_F_check_agglevels_and_constraints_5), v14+int32(16))
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L5
	} else {
		goto L87
	}
L87:
	;
	F_parser_errposition(m, v184, v35)
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L5
	} else {
		goto L88
	}
L88:
	;
	F_errfinish(m, int32(_a_F_check_agglevels_and_constraints_1), int32(600), int32(_a_F_check_agglevels_and_constraints_6))
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L5
	} else {
		goto L89
	}
L89:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L90:
	;
	F_errcode(m, int32(50364548))
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L5
	} else {
		goto L91
	}
L91:
	;
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v184)+68))
	if base.Ui32(v274) <= base.Ui32(int32(44)) {
		goto L93
	} else {
		goto L94
	}
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v281
	if v16 == int32(9) {
		goto L96
	} else {
		goto L97
	}
L93:
	;
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v274<<(uint(int32(2))%32))+uint32(_c_F_check_agglevels_and_constraints[0])))
	v281 = v279
	goto L95
L94:
	;
	v281 = int32(_a_F_check_agglevels_and_constraints_11)
	goto L95
L95:
	;
	goto L92
L96:
	;
	v287 = int32(_a_F_check_agglevels_and_constraints_12)
	goto L98
L97:
	;
	v287 = int32(_a_F_check_agglevels_and_constraints_13)
	goto L98
L98:
	;
	F_errmsg_internal(m, v287, v14+int32(32))
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L5
	} else {
		goto L99
	}
L99:
	;
	F_parser_errposition(m, v184, v35)
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L5
	} else {
		goto L100
	}
L100:
	;
	F_errfinish(m, int32(_a_F_check_agglevels_and_constraints_1), int32(615), int32(_a_F_check_agglevels_and_constraints_6))
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L5
	} else {
		goto L101
	}
L101:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L102:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L5
	} else {
		goto L103
	}
L103:
	;
	F_errmsg(m, int32(_a_F_check_agglevels_and_constraints_51), int32(0))
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L5
	} else {
		goto L104
	}
L104:
	;
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v14)+72))
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v313)+8))
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v314)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v315
	F_errdetail(m, int32(_a_F_check_agglevels_and_constraints_52), v14)
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L5
	} else {
		goto L105
	}
L105:
	;
	F_parser_errposition(m, l0, v35)
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L5
	} else {
		goto L106
	}
L106:
	;
	F_errfinish(m, int32(_a_F_check_agglevels_and_constraints_1), int32(708), int32(_a_F_check_agglevels_and_constraints_2))
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L5
	} else {
		goto L107
	}
L107:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L108:
	;
	F_errcode(m, int32(50364548))
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L5
	} else {
		goto L109
	}
L109:
	;
	F_errmsg(m, int32(_a_F_check_agglevels_and_constraints_50), int32(0))
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L5
	} else {
		goto L110
	}
L110:
	;
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v14)+60))
	v339 = F_locate_var_of_level(m, v30, v338)
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L5
	} else {
		goto L111
	}
L111:
	;
	F_parser_errposition(m, l0, v339)
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L5
	} else {
		goto L112
	}
L112:
	;
	F_errfinish(m, int32(_a_F_check_agglevels_and_constraints_1), int32(731), int32(_a_F_check_agglevels_and_constraints_2))
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L5
	} else {
		goto L113
	}
L113:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L114:
	;
	F_errcode(m, int32(50364548))
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L5
	} else {
		goto L115
	}
L115:
	;
	F_errmsg(m, int32(_a_F_check_agglevels_and_constraints_0), int32(0))
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L5
	} else {
		goto L116
	}
L116:
	;
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v14)+64))
	v360 = F_locate_agg_of_level(m, v30, v359)
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L5
	} else {
		goto L117
	}
L117:
	;
	F_parser_errposition(m, l0, v360)
	mBase = m.M
	v363 = m.ExcPending
	if v363 != 0 {
		goto L5
	} else {
		goto L118
	}
L118:
	;
	F_errfinish(m, int32(_a_F_check_agglevels_and_constraints_1), int32(738), int32(_a_F_check_agglevels_and_constraints_2))
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L5
	} else {
		goto L119
	}
L119:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L120:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v375 = m.ExcPending
	if v375 != 0 {
		goto L5
	} else {
		goto L121
	}
L121:
	;
	F_errmsg(m, int32(_a_F_check_agglevels_and_constraints_51), int32(0))
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L5
	} else {
		goto L122
	}
L122:
	;
	v380 = *(*int32)(unsafe.Add(mBase, uint32(v14)+72))
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v380)+8))
	v382 = *(*int32)(unsafe.Add(mBase, uint32(v381)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = v382
	F_errdetail(m, int32(_a_F_check_agglevels_and_constraints_52), v14+int32(48))
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L5
	} else {
		goto L123
	}
L123:
	;
	F_parser_errposition(m, l0, v35)
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L5
	} else {
		goto L124
	}
L124:
	;
	F_errfinish(m, int32(_a_F_check_agglevels_and_constraints_1), int32(745), int32(_a_F_check_agglevels_and_constraints_2))
	mBase = m.M
	v395 = m.ExcPending
	if v395 != 0 {
		goto L5
	} else {
		goto L125
	}
L125:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_check_amoptsproc_signature(m *base.Module, l0 int32) int32 {
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
	var v16 int32
	_ = v16
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	*(*int32)(unsafe.Add(mBase, uint32(v5))) = int32(2281)
	v10 = int32(1)
	v13 = F_check_amproc_signature(m, l0, int32(2278), v10, v10, v10, v5)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		m.G0 = v5 + int32(16)
		return v13
	}
}
func F_check_amproc_signature(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
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
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v82 int32
	_ = v82
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	v16 = F_SearchSysCache1(m, int32(47), l0)
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
	if v16 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v16)+16))
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+22)))
	v22 = v20 + v21
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+108))
	if v23 != l1 {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	goto L5
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L1
	} else {
		goto L28
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = l5
	if int32(0) < l4 {
		goto L12
	} else {
		goto L13
	}
L7:
	;
	v33 = int32(0)
	goto L6
L8:
	;
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+100)))
	if v25 != 0 {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v26 = int32(*(*int16)(unsafe.Add(mBase, uint32(v22)+104)))
	if v26 < l3 {
		goto L7
	} else {
		goto L10
	}
L10:
	;
	if v26 <= l4 {
		v33 = int32(1)
		goto L6
	} else {
		goto L11
	}
L11:
	;
	goto L7
L12:
	;
	v41 = v33
	v45 = int32(0)
	goto L15
L13:
	;
	v72 = v33
	goto L14
L14:
	;
	F_ReleaseCatCache(m, v16)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L1
	} else {
		goto L27
	}
L15:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v50 + int32(4)
	v54 = int32(*(*int16)(unsafe.Add(mBase, uint32(v22)+104)))
	if v54 <= v45 {
		v65 = v41
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v72 = v65
	goto L14
L17:
	;
	v69 = v45 + int32(1)
	if v69 != l4 {
		v41 = v65
		v45 = v69
		goto L15
	} else {
		goto L26
	}
L18:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v22+int32(136)+v45<<(uint(int32(2))%32))))
	if l2 != 0 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v65 = int32(0)
	goto L17
L20:
	;
	if v56 != v60 {
		goto L19
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v62 = F_IsBinaryCoercible(m, v56, v60)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L1
	} else {
		goto L24
	}
L23:
	;
	v65 = v41
	goto L17
L24:
	;
	if v62 != 0 {
		v65 = v41
		goto L17
	} else {
		goto L25
	}
L25:
	;
	goto L19
L26:
	;
	goto L16
L27:
	;
	m.G0 = v13 + int32(16)
	return v72
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = l0
	F_errmsg_internal(m, int32(_a_F_check_amproc_signature_0), v13)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	F_errfinish(m, int32(_a_F_check_amproc_signature_1), int32(163), int32(_a_F_check_amproc_signature_2))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_check_bonjour(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	v4 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v4 == int32(1) {
		v8 = *(*int32)(unsafe.Add(mBase, _c_F_check_bonjour[0]))
		*(*int32)(unsafe.Add(mBase, _c_F_check_bonjour[1])) = v8
		v14 = F_format_elog_string(m, int32(_a_F_check_bonjour_0), int32(0))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, _c_F_check_bonjour[2])) = v14
			return v4 ^ int32(1)
		}
	} else {
		return v4 ^ int32(1)
	}
}
func F_check_parameter_resolution_walker(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	v3 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	if l0 == v3 {
		v67 = v3
		m.G0 = v9 + int32(32)
		return v67
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if v13 != int32(67) {
			if v13 != int32(8) {
				v64 = F_expression_tree_walker_impl(m, l0, int32(495), l1)
				mBase = m.M
				v65 = m.ExcPending
				if v65 != 0 {
					return int32(0)
				} else {
					v67 = v64
					m.G0 = v9 + int32(32)
					return v67
				}
			} else {
				v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				if v18 != 0 {
					v67 = v3
					m.G0 = v9 + int32(32)
					return v67
				} else {
					v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					if v19 <= int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v77 = m.ExcPending
						if v77 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(33685636))
							mBase = m.M
							v80 = m.ExcPending
							if v80 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v9))) = v19
								F_errmsg(m, int32(_a_F_check_parameter_resolution_walker_0), v9)
								mBase = m.M
								v84 = m.ExcPending
								if v84 != 0 {
									return int32(0)
								} else {
									v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
									F_parser_errposition(m, l1, v85)
									mBase = m.M
									v87 = m.ExcPending
									if v87 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_check_parameter_resolution_walker_1), int32(305), int32(_a_F_check_parameter_resolution_walker_2))
										mBase = m.M
										v92 = m.ExcPending
										if v92 != 0 {
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
					} else {
						v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)+120))
						v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
						v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
						if v24 < v19 {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v77 = m.ExcPending
							if v77 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(33685636))
								mBase = m.M
								v80 = m.ExcPending
								if v80 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v9))) = v19
									F_errmsg(m, int32(_a_F_check_parameter_resolution_walker_0), v9)
									mBase = m.M
									v84 = m.ExcPending
									if v84 != 0 {
										return int32(0)
									} else {
										v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
										F_parser_errposition(m, l1, v85)
										mBase = m.M
										v87 = m.ExcPending
										if v87 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_check_parameter_resolution_walker_1), int32(305), int32(_a_F_check_parameter_resolution_walker_2))
											mBase = m.M
											v92 = m.ExcPending
											if v92 != 0 {
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
						} else {
							v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							v27 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
							v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
							v34 = *(*int32)(unsafe.Add(mBase, uint32(v28+v19<<(uint(int32(2))%32)-int32(4))))
							if v26 == v34 {
								v67 = v3
								m.G0 = v9 + int32(32)
								return v67
							} else {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v41 = m.ExcPending
								if v41 != 0 {
									return int32(0)
								} else {
									F_errcode(m, int32(134348932))
									mBase = m.M
									v44 = m.ExcPending
									if v44 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v19
										F_errmsg(m, int32(_a_F_check_parameter_resolution_walker_3), v9+int32(16))
										mBase = m.M
										v50 = m.ExcPending
										if v50 != 0 {
											return int32(0)
										} else {
											v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
											F_parser_errposition(m, l1, v51)
											mBase = m.M
											v53 = m.ExcPending
											if v53 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(_a_F_check_parameter_resolution_walker_1), int32(312), int32(_a_F_check_parameter_resolution_walker_2))
												mBase = m.M
												v58 = m.ExcPending
												if v58 != 0 {
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
				}
			}
		} else {
			v61 = F_query_tree_walker_impl(m, l0, int32(495), l1, int32(0))
			mBase = m.M
			v62 = m.ExcPending
			if v62 != 0 {
				return int32(0)
			} else {
				v67 = v61
				m.G0 = v9 + int32(32)
				return v67
			}
		}
	}
}
func F_check_primary_slot_name(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	v4 = m.G0
	v6 = v4 - int32(32)
	m.G0 = v6
	v8 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v6)+24)) = v8
	*(*int32)(unsafe.Add(mBase, uint32(v6)+20)) = v8
	v12 = int32(1)
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v13 == v8 {
		v60 = v12
		m.G0 = v6 + int32(32)
		return v60
	} else {
		v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
		if v16 == int32(0) {
			v60 = v12
			m.G0 = v6 + int32(32)
			return v60
		} else {
			v25 = F_ReplicationSlotValidateNameInternal(m, v13, v6+int32(28), v6+int32(24), v6+int32(20))
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return int32(0)
			} else {
				if v25 != 0 {
					v60 = v12
					m.G0 = v6 + int32(32)
					return v60
				} else {
					v29 = *(*int32)(unsafe.Add(mBase, uint32(v6)+28))
					*(*int32)(unsafe.Add(mBase, _c_F_check_primary_slot_name[0])) = v29
					v32 = int32(0)
					v34 = *(*int32)(unsafe.Add(mBase, _c_F_check_primary_slot_name[1]))
					*(*int32)(unsafe.Add(mBase, _c_F_check_primary_slot_name[2])) = v34
					v37 = *(*int32)(unsafe.Add(mBase, uint32(v6)+24))
					*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = v37
					v43 = F_format_elog_string(m, int32(_a_F_check_primary_slot_name_0), v6+int32(16))
					mBase = m.M
					v44 = m.ExcPending
					if v44 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, _c_F_check_primary_slot_name[3])) = v43
						v46 = *(*int32)(unsafe.Add(mBase, uint32(v6)+20))
						if v46 == int32(0) {
							v60 = v32
							m.G0 = v6 + int32(32)
							return v60
						} else {
							v50 = *(*int32)(unsafe.Add(mBase, _c_F_check_primary_slot_name[1]))
							*(*int32)(unsafe.Add(mBase, _c_F_check_primary_slot_name[2])) = v50
							v53 = *(*int32)(unsafe.Add(mBase, uint32(v6)+20))
							*(*int32)(unsafe.Add(mBase, uint32(v6))) = v53
							v57 = F_format_elog_string(m, int32(_a_F_check_primary_slot_name_0), v6)
							mBase = m.M
							v58 = m.ExcPending
							if v58 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, _c_F_check_primary_slot_name[4])) = v57
								v60 = v32
								m.G0 = v6 + int32(32)
								return v60
							}
						}
					}
				}
			}
		}
	}
}
func F_check_random_seed(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v5 = F_guc_malloc(m, int32(4))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l1))) = v5
		if v5 != 0 {
			*(*int32)(unsafe.Add(mBase, uint32(v5))) = base.B2i32(base.Ui32(int32(10)) < base.Ui32(l2))
		} else {
		}
		return base.B2i32(v5 != int32(0))
	}
}
func F_check_safe_enum_use(m *base.Module, l0 int32) {
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
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+21)))
	if v11&int32(1) != 0 {
		m.G0 = v8 + int32(16)
		return
	} else {
		v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+22)))
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
		v16 = F_TransactionIdIsInProgress(m, v15)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return
		} else {
			if v16 == int32(0) {
				v20 = F_TransactionIdDidCommit(m, v15)
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return
				} else {
					if v20 != 0 {
						m.G0 = v8 + int32(16)
						return
					} else {
						v22 = v10 + v14
						v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
						v24 = m.G0
						v26 = v24 - int32(16)
						m.G0 = v26
						*(*int32)(unsafe.Add(mBase, uint32(v26)+12)) = v23
						v31 = *(*int32)(unsafe.Add(mBase, _c_F_check_safe_enum_use[0]))
						if v31 != 0 {
							v37 = F_hash_search(m, v31, v26+int32(12), int32(0), v26+int32(11))
							mBase = m.M
							v38 = m.ExcPending
							if v38 != 0 {
								return
							} else {
								v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+11)))
								v40 = v39
								m.G0 = v26 + int32(16)
								if v40&int32(1) == int32(0) {
									m.G0 = v8 + int32(16)
									return
								} else {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v51 = m.ExcPending
									if v51 != 0 {
										return
									} else {
										F_errcode(m, int32(67240261))
										mBase = m.M
										v54 = m.ExcPending
										if v54 != 0 {
											return
										} else {
											v55 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
											v56 = F_format_type_be(m, v55)
											mBase = m.M
											v57 = m.ExcPending
											if v57 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v56
												*(*int32)(unsafe.Add(mBase, uint32(v8))) = v22 + int32(12)
												F_errmsg(m, int32(_a_F_check_safe_enum_use_0), v8)
												mBase = m.M
												v64 = m.ExcPending
												if v64 != 0 {
													return
												} else {
													F_errhint(m, int32(_a_F_check_safe_enum_use_1), int32(0))
													mBase = m.M
													v68 = m.ExcPending
													if v68 != 0 {
														return
													} else {
														F_errfinish(m, int32(_a_F_check_safe_enum_use_2), int32(102), int32(_a_F_check_safe_enum_use_3))
														mBase = m.M
														v73 = m.ExcPending
														if v73 != 0 {
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
						} else {
							v40 = int32(0)
							m.G0 = v26 + int32(16)
							if v40&int32(1) == int32(0) {
								m.G0 = v8 + int32(16)
								return
							} else {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v51 = m.ExcPending
								if v51 != 0 {
									return
								} else {
									F_errcode(m, int32(67240261))
									mBase = m.M
									v54 = m.ExcPending
									if v54 != 0 {
										return
									} else {
										v55 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
										v56 = F_format_type_be(m, v55)
										mBase = m.M
										v57 = m.ExcPending
										if v57 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v56
											*(*int32)(unsafe.Add(mBase, uint32(v8))) = v22 + int32(12)
											F_errmsg(m, int32(_a_F_check_safe_enum_use_0), v8)
											mBase = m.M
											v64 = m.ExcPending
											if v64 != 0 {
												return
											} else {
												F_errhint(m, int32(_a_F_check_safe_enum_use_1), int32(0))
												mBase = m.M
												v68 = m.ExcPending
												if v68 != 0 {
													return
												} else {
													F_errfinish(m, int32(_a_F_check_safe_enum_use_2), int32(102), int32(_a_F_check_safe_enum_use_3))
													mBase = m.M
													v73 = m.ExcPending
													if v73 != 0 {
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
					}
				}
			} else {
				v22 = v10 + v14
				v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
				v24 = m.G0
				v26 = v24 - int32(16)
				m.G0 = v26
				*(*int32)(unsafe.Add(mBase, uint32(v26)+12)) = v23
				v31 = *(*int32)(unsafe.Add(mBase, _c_F_check_safe_enum_use[0]))
				if v31 != 0 {
					v37 = F_hash_search(m, v31, v26+int32(12), int32(0), v26+int32(11))
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
						return
					} else {
						v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+11)))
						v40 = v39
						m.G0 = v26 + int32(16)
						if v40&int32(1) == int32(0) {
							m.G0 = v8 + int32(16)
							return
						} else {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v51 = m.ExcPending
							if v51 != 0 {
								return
							} else {
								F_errcode(m, int32(67240261))
								mBase = m.M
								v54 = m.ExcPending
								if v54 != 0 {
									return
								} else {
									v55 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
									v56 = F_format_type_be(m, v55)
									mBase = m.M
									v57 = m.ExcPending
									if v57 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v56
										*(*int32)(unsafe.Add(mBase, uint32(v8))) = v22 + int32(12)
										F_errmsg(m, int32(_a_F_check_safe_enum_use_0), v8)
										mBase = m.M
										v64 = m.ExcPending
										if v64 != 0 {
											return
										} else {
											F_errhint(m, int32(_a_F_check_safe_enum_use_1), int32(0))
											mBase = m.M
											v68 = m.ExcPending
											if v68 != 0 {
												return
											} else {
												F_errfinish(m, int32(_a_F_check_safe_enum_use_2), int32(102), int32(_a_F_check_safe_enum_use_3))
												mBase = m.M
												v73 = m.ExcPending
												if v73 != 0 {
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
				} else {
					v40 = int32(0)
					m.G0 = v26 + int32(16)
					if v40&int32(1) == int32(0) {
						m.G0 = v8 + int32(16)
						return
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v51 = m.ExcPending
						if v51 != 0 {
							return
						} else {
							F_errcode(m, int32(67240261))
							mBase = m.M
							v54 = m.ExcPending
							if v54 != 0 {
								return
							} else {
								v55 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
								v56 = F_format_type_be(m, v55)
								mBase = m.M
								v57 = m.ExcPending
								if v57 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v56
									*(*int32)(unsafe.Add(mBase, uint32(v8))) = v22 + int32(12)
									F_errmsg(m, int32(_a_F_check_safe_enum_use_0), v8)
									mBase = m.M
									v64 = m.ExcPending
									if v64 != 0 {
										return
									} else {
										F_errhint(m, int32(_a_F_check_safe_enum_use_1), int32(0))
										mBase = m.M
										v68 = m.ExcPending
										if v68 != 0 {
											return
										} else {
											F_errfinish(m, int32(_a_F_check_safe_enum_use_2), int32(102), int32(_a_F_check_safe_enum_use_3))
											mBase = m.M
											v73 = m.ExcPending
											if v73 != 0 {
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
			}
		}
	}
}
func F_check_stage_log_stats(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	v4 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v6 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_check_stage_log_stats[0])))
	v7 = v4 & v6
	if v7&int32(1) != 0 {
		v11 = *(*int32)(unsafe.Add(mBase, _c_F_check_stage_log_stats[1]))
		*(*int32)(unsafe.Add(mBase, _c_F_check_stage_log_stats[2])) = v11
		v17 = F_format_elog_string(m, int32(_a_F_check_stage_log_stats_0), int32(0))
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, _c_F_check_stage_log_stats[3])) = v17
			return (v7 ^ int32(-1)) & int32(1)
		}
	} else {
		return (v7 ^ int32(-1)) & int32(1)
	}
}
func F_checkdig(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
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
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	v2 = int32(0)
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v12 = base.B2i32(v10 == int32(77))
	if v10 == int32(77) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v13 = int32(3)
	goto L3
L2:
	;
	v13 = v2
	goto L3
L3:
	;
	if v10 == int32(0) {
		v58 = v13
		v60 = v2
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v62 = int32(10)
	v67 = base.I32_rem_u_s(v58*int32(3)+v60, v62)
	if v67 != 0 {
		goto L16
	} else {
		goto L17
	}
L5:
	;
	v17 = l0
	v18 = v10
	v19 = v12
	v20 = v13
	v21 = int32(13)
	v22 = v2
	goto L6
L6:
	;
	v27 = (v18 - int32(48)) & int32(255)
	if base.Ui32(v27) <= base.Ui32(int32(9)) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v58 = v44
	v60 = v46
	goto L4
L8:
	;
	v32 = v19 & int32(1)
	if v32 != 0 {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	v43 = v19
	v44 = v20
	v45 = v21
	v46 = v22
	goto L10
L10:
	;
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+1)))
	if v48 == int32(0) {
		v58 = v44
		v60 = v46
		goto L4
	} else {
		goto L14
	}
L11:
	;
	v33 = int32(0)
	goto L13
L12:
	;
	v33 = v27
	goto L13
L13:
	;
	v39 = int32(1)
	v43 = v19 + v39
	v44 = (int32(0)-v32)&v27 + v20
	v45 = v21 - v39
	v46 = v33 + v22
	goto L10
L14:
	;
	v51 = int32(1)
	if base.Ui32(v51) < base.Ui32(v45) {
		v17 = v17 + v51
		v18 = v48
		v19 = v43
		v20 = v44
		v21 = v45
		v22 = v46
		goto L6
	} else {
		goto L15
	}
L15:
	;
	goto L7
L16:
	;
	v70 = v62 - v67
	goto L18
L17:
	;
	v70 = int32(0)
	goto L18
L18:
	;
	return v70
}
func F_clearerr(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v2 & int32(-49)
	return
}
func F_clog_identify(m *base.Module, l0 int32) int32 {
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	v5 = l0 & int32(240)
	if v5 == int32(16) {
		v8 = int32(_a_F_clog_identify_0)
	} else {
		v8 = int32(0)
	}
	if v5 != 0 {
		v10 = v8
	} else {
		v10 = int32(_a_F_clog_identify_1)
	}
	return v10
}
func F_clog_redo(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int64
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int64
	_ = v40
	var v42 int64
	_ = v42
	var v43 int64
	_ = v43
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+48)))
	v12 = v10 & int32(240)
	if v12 != 0 {
		if v12 == int32(16) {
			v15 = *(*int32)(unsafe.Add(mBase, uint32(v9)+64))
			v16 = *(*int64)(unsafe.Add(mBase, uint32(v15)))
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
			F_AdvanceOldestClogXid(m, v17)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return
			} else {
				F_SimpleLruTruncate(m, int32(_a_F_clog_redo_0), v16)
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return
				} else {
					m.G0 = v7 + int32(16)
					return
				}
			}
		} else {
			F_errstart_cold(m, int32(23), int32(0))
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v7))) = v12
				F_errmsg_internal(m, int32(_a_F_clog_redo_1), v7)
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_clog_redo_2), int32(1142), int32(_a_F_clog_redo_3))
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	} else {
		v37 = *(*int32)(unsafe.Add(mBase, _c_F_clog_redo[0]))
		v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)+28))
		v39 = *(*int32)(unsafe.Add(mBase, uint32(v9)+64))
		v40 = *(*int64)(unsafe.Add(mBase, uint32(v39)))
		v42 = int64(*(*uint16)(unsafe.Add(mBase, _c_F_clog_redo[1])))
		v43 = base.I64_rem_s(v40, v42)
		v47 = v38 + base.I32_wrap_i64(v43)<<(uint(int32(7))%32)
		v49 = F_LWLockAcquire(m, v47, int32(0))
		mBase = m.M
		v50 = m.ExcPending
		if v50 != 0 {
			return
		} else {
			v51 = int32(_a_F_clog_redo_0)
			v53 = F_SimpleLruZeroPage(m, v51, v40)
			mBase = m.M
			v54 = m.ExcPending
			if v54 != 0 {
				return
			} else {
				F_SimpleLruWritePage(m, v51, v53)
				mBase = m.M
				v56 = m.ExcPending
				if v56 != 0 {
					return
				} else {
					F_LWLockRelease(m, v47)
					mBase = m.M
					v58 = m.ExcPending
					if v58 != 0 {
						return
					} else {
						m.G0 = v7 + int32(16)
						return
					}
				}
			}
		}
	}
}
func F_close(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v13 int32
	_ = v13
	v2 = m.Wasi_snapshot_preview1.Fd_close(m, l0)
	mBase = m.M
	if v2 != int32(27) {
		v6 = v2
	} else {
		v6 = int32(0)
	}
	if v6 == int32(0) {
		v13 = int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, _c_F_close[0])) = v6
		v13 = int32(-1)
	}
	return v13
}
func F_close_ps(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 float64
	_ = v12
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = F_palloc(m, int32(16))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		v12 = F_lseg_closept_point(m, v8, v5, v6)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int32(0)
		} else {
			if base.Ui64(int64(9218868437227405313)) <= base.Ui64(base.I64_reinterpret_f64(v12)&int64(9223372036854775807)) {
				v19 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v19)
				v22 = int32(0)
			} else {
				v22 = v8
			}
			return v22
		}
	}
}
func F_cmpaliases(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3))))
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4))))
	if base.B2i32(v7 == int32(0))|base.B2i32(v7 != v10) != 0 {
		v28 = v7
		v29 = v10
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v28 - v29
L2:
	;
	goto L1
L3:
	;
	v13 = v3
	v14 = v4
	goto L4
L4:
	;
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+1)))
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+1)))
	if v18 == int32(0) {
		v28 = v18
		v29 = v17
		goto L2
	} else {
		goto L6
	}
L5:
	;
	v28 = v18
	v29 = v17
	goto L2
L6:
	;
	v21 = int32(1)
	if v18 == v17 {
		v13 = v13 + v21
		v14 = v14 + v21
		goto L4
	} else {
		goto L7
	}
L7:
	;
	goto L5
}
func F_cmpcmdflag(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v3 == int32(2) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v45
L2:
	;
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v7 == v8 {
		v45 = int32(0)
		goto L1
	} else {
		goto L5
	}
L3:
	;
	goto L4
L4:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
	if base.B2i32(v19 == int32(0))|base.B2i32(v19 != v22) != 0 {
		v40 = v19
		v41 = v22
		goto L10
	} else {
		goto L11
	}
L5:
	;
	if base.Ui32(v8) < base.Ui32(v7) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v13 = int32(1)
	goto L8
L7:
	;
	v13 = int32(-1)
	goto L8
L8:
	;
	return v13
L9:
	;
	v45 = v40 - v41
	goto L1
L10:
	;
	goto L9
L11:
	;
	v25 = v15
	v26 = v16
	goto L12
L12:
	;
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+1)))
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+1)))
	if v30 == int32(0) {
		v40 = v30
		v41 = v29
		goto L10
	} else {
		goto L14
	}
L13:
	;
	v40 = v30
	v41 = v29
	goto L10
L14:
	;
	v33 = int32(1)
	if v30 == v29 {
		v25 = v25 + v33
		v26 = v26 + v33
		goto L12
	} else {
		goto L15
	}
L15:
	;
	goto L13
}
func F_cmpspellaffix(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4))))
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6))))
	if base.B2i32(v9 == int32(0))|base.B2i32(v9 != v12) != 0 {
		v30 = v9
		v31 = v12
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v30 - v31
L2:
	;
	goto L1
L3:
	;
	v15 = v4
	v16 = v6
	goto L4
L4:
	;
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+1)))
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+1)))
	if v20 == int32(0) {
		v30 = v20
		v31 = v19
		goto L2
	} else {
		goto L6
	}
L5:
	;
	v30 = v20
	v31 = v19
	goto L2
L6:
	;
	v23 = int32(1)
	if v20 == v19 {
		v15 = v15 + v23
		v16 = v16 + v23
		goto L4
	} else {
		goto L7
	}
L7:
	;
	goto L5
}
func F_coerce_null_to_domain(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
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
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = l1
	v15 = F_getBaseTypeAndTypmod(m, l0, v10+int32(12))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		v19 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
		v22 = F_makeConst(m, v15, v19, l2, l3, int32(0), int32(1), l4)
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return int32(0)
		} else {
			if l0 != v15 {
				v25 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
				v26 = int32(0)
				v30 = F_coerce_to_domain(m, v22, v15, v25, l0, v26, int32(2), int32(-1), v26)
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return int32(0)
				} else {
					v32 = v30
					m.G0 = v10 + int32(16)
					return v32
				}
			} else {
				v32 = v22
				m.G0 = v10 + int32(16)
				return v32
			}
		}
	}
}
func F_colorTrgmInfoPenaltyCmp(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 float32
	_ = v8
	var v9 float32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	v8 = *(*float32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = *(*float32)(unsafe.Add(mBase, uint32(l1)+20))
	if base.F32_ne(v8, v9) != 0 {
		v11 = int32(-1)
	} else {
		v11 = int32(0)
	}
	if base.F32_lt(v8, v9) != 0 {
		v13 = int32(1)
	} else {
		v13 = v11
	}
	return v13
}
func F_combo_decrypt(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v7)+20))
	v10 = m.T0[v9].(func(*base.Module, int32, int32, int32, int32, int32, int32) int32)(m, v7, v8, l1, l2, l3, l4)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		return v10
	}
}
func F_combo_encrypt(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
	v10 = m.T0[v9].(func(*base.Module, int32, int32, int32, int32, int32, int32) int32)(m, v7, v8, l1, l2, l3, l4)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		return v10
	}
}
func F_combo_encrypt_len(m *base.Module, l0 int32, l1 int32) int32 {
	return l1 + int32(512)
}
func F_comp_ptrgm(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v17 int32
	_ = v17
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_comp_ptrgm[0]))
	v6 = m.T0[v5].(func(*base.Module, int32, int32) int32)(m, l0, l1)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		if v6 != 0 {
			v17 = v6
		} else {
			v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			v17 = base.B2i32(v11 < v10) - base.B2i32(v10 < v11)
		}
		return v17
	}
}
func F_comparePairs(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v5 == v6 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v85
L2:
	;
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if base.Ui32(int32(4)) <= base.Ui32(v5) {
		goto L8
	} else {
		goto L9
	}
L3:
	;
	goto L4
L4:
	;
	if base.Ui32(v6) < base.Ui32(v5) {
		goto L28
	} else {
		goto L29
	}
L5:
	;
	if v71 != 0 {
		v85 = v71
		goto L1
	} else {
		goto L23
	}
L6:
	;
	v71 = int32(0)
	goto L5
L7:
	;
	v45 = v40
	v46 = v41
	v47 = v42
	goto L17
L8:
	;
	if (v8|v9)&int32(3) != 0 {
		v40 = v8
		v41 = v9
		v42 = v5
		goto L7
	} else {
		goto L11
	}
L9:
	;
	v33 = v8
	v34 = v9
	v35 = v5
	goto L10
L10:
	;
	if v35 == int32(0) {
		goto L6
	} else {
		goto L16
	}
L11:
	;
	v17 = v8
	v18 = v9
	v19 = v5
	goto L12
L12:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	if v22 != v23 {
		v40 = v17
		v41 = v18
		v42 = v19
		goto L7
	} else {
		goto L14
	}
L13:
	;
	v33 = v28
	v34 = v26
	v35 = v30
	goto L10
L14:
	;
	v25 = int32(4)
	v26 = v18 + v25
	v28 = v17 + v25
	v30 = v19 - v25
	if base.Ui32(int32(3)) < base.Ui32(v30) {
		v17 = v28
		v18 = v26
		v19 = v30
		goto L12
	} else {
		goto L15
	}
L15:
	;
	goto L13
L16:
	;
	v40 = v33
	v41 = v34
	v42 = v35
	goto L7
L17:
	;
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45))))
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46))))
	if v50 == v51 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v71 = v50 - v51
	goto L5
L19:
	;
	v53 = int32(1)
	v58 = v47 - v53
	if v58 != 0 {
		v45 = v45 + v53
		v46 = v46 + v53
		v47 = v58
		goto L17
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	goto L18
L22:
	;
	goto L6
L23:
	;
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+17)))
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+17)))
	if v73 == v74 {
		v85 = int32(0)
		goto L1
	} else {
		goto L24
	}
L24:
	;
	if v73 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v78 = int32(1)
	goto L27
L26:
	;
	v78 = int32(-1)
	goto L27
L27:
	;
	return v78
L28:
	;
	v83 = int32(1)
	goto L30
L29:
	;
	v83 = int32(-1)
	goto L30
L30:
	;
	v85 = v83
	goto L1
}
func F_compare_distances(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 float64
	_ = v8
	var v9 float64
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	v8 = *(*float64)(unsafe.Add(mBase, uint32(l0)+8))
	v9 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
	if base.F64_gt(v8, v9) != 0 {
		v11 = int32(-1)
	} else {
		v11 = int32(0)
	}
	if base.F64_lt(v8, v9) != 0 {
		v13 = int32(1)
	} else {
		v13 = v11
	}
	return v13
}
func F_compare_values(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
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
	var v21 int32
	_ = v21
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v8 = F_FunctionCall2Coll(m, v4, v5, v6, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		if v8 != 0 {
			v21 = int32(-1)
			return v21
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
			v14 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v17 = F_FunctionCall2Coll(m, v13, v14, v15, v16)
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				v21 = base.B2i32(v17 != int32(0))
				return v21
			}
		}
	}
}
func F_comparecost_3(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 float32
	_ = v7
	var v8 float32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	v7 = *(*float32)(unsafe.Add(mBase, uint32(l0)+4))
	v8 = *(*float32)(unsafe.Add(mBase, uint32(l1)+4))
	if base.F32_gt(v7, v8) != 0 {
		v10 = int32(1)
	} else {
		v10 = int32(-1)
	}
	if base.F32_ne(v7, v8) != 0 {
		v13 = v10
	} else {
		v13 = int32(0)
	}
	return v13
}
func F_compatible_oper_opid(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
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
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	v4 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v15 = F_oper(m, v4, l0, l1, l2, v4, int32(-1))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		if v15 == int32(0) {
			v58 = v4
			m.G0 = v10 + int32(16)
			if v58 != 0 {
				v62 = *(*int32)(unsafe.Add(mBase, uint32(v58)+16))
				v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62)+22)))
				v65 = *(*int32)(unsafe.Add(mBase, uint32(v62+v63)))
				F_ReleaseCatCache(m, v58)
				mBase = m.M
				v67 = m.ExcPending
				if v67 != 0 {
					return int32(0)
				} else {
					v69 = v65
					return v69
				}
			} else {
				v69 = v4
				return v69
			}
		} else {
			v21 = *(*int32)(unsafe.Add(mBase, uint32(v15)+16))
			v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+22)))
			v23 = v21 + v22
			v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+80))
			v25 = F_IsBinaryCoercible(m, l1, v24)
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return int32(0)
			} else {
				if v25 == int32(0) {
					F_ReleaseCatCache(m, v15)
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return int32(0)
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v39 = m.ExcPending
						if v39 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(52461700))
							mBase = m.M
							v42 = m.ExcPending
							if v42 != 0 {
								return int32(0)
							} else {
								v43 = F_op_signature_string(m, l0, l1, l2)
								mBase = m.M
								v44 = m.ExcPending
								if v44 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v10))) = v43
									F_errmsg(m, int32(_a_F_compatible_oper_opid_0), v10)
									mBase = m.M
									v48 = m.ExcPending
									if v48 != 0 {
										return int32(0)
									} else {
										F_parser_errposition(m, int32(0), int32(-1))
										mBase = m.M
										v52 = m.ExcPending
										if v52 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_compatible_oper_opid_1), int32(475), int32(_a_F_compatible_oper_opid_2))
											mBase = m.M
											v57 = m.ExcPending
											if v57 != 0 {
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
				} else {
					v29 = *(*int32)(unsafe.Add(mBase, uint32(v23)+84))
					v30 = F_IsBinaryCoercible(m, l2, v29)
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return int32(0)
					} else {
						if v30 == int32(0) {
							F_ReleaseCatCache(m, v15)
							mBase = m.M
							v35 = m.ExcPending
							if v35 != 0 {
								return int32(0)
							} else {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v39 = m.ExcPending
								if v39 != 0 {
									return int32(0)
								} else {
									F_errcode(m, int32(52461700))
									mBase = m.M
									v42 = m.ExcPending
									if v42 != 0 {
										return int32(0)
									} else {
										v43 = F_op_signature_string(m, l0, l1, l2)
										mBase = m.M
										v44 = m.ExcPending
										if v44 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v10))) = v43
											F_errmsg(m, int32(_a_F_compatible_oper_opid_0), v10)
											mBase = m.M
											v48 = m.ExcPending
											if v48 != 0 {
												return int32(0)
											} else {
												F_parser_errposition(m, int32(0), int32(-1))
												mBase = m.M
												v52 = m.ExcPending
												if v52 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(_a_F_compatible_oper_opid_1), int32(475), int32(_a_F_compatible_oper_opid_2))
													mBase = m.M
													v57 = m.ExcPending
													if v57 != 0 {
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
						} else {
							v58 = v15
							m.G0 = v10 + int32(16)
							if v58 != 0 {
								v62 = *(*int32)(unsafe.Add(mBase, uint32(v58)+16))
								v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62)+22)))
								v65 = *(*int32)(unsafe.Add(mBase, uint32(v62+v63)))
								F_ReleaseCatCache(m, v58)
								mBase = m.M
								v67 = m.ExcPending
								if v67 != 0 {
									return int32(0)
								} else {
									v69 = v65
									return v69
								}
							} else {
								v69 = v4
								return v69
							}
						}
					}
				}
			}
		}
	}
}
func F_composite_to_json(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
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
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v142 int32
	_ = v142
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v163 int32
	_ = v163
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
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
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	v13 = m.G0
	v15 = v13 - int32(48)
	m.G0 = v15
	v17 = F_pg_detoast_datum(m, l0)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	v21 = F_lookup_rowtype_tupdesc(m, v19, v20)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+44)) = v17
	*(*int32)(unsafe.Add(mBase, uint32(v15)+28)) = int32(base.Ui32(v23) >> (uint(int32(2)) % 32))
	F_appendStringInfoChar(m, l1, int32(123))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	if int32(0) < v31 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	if l2 != 0 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	F_appendStringInfoChar(m, l1, int32(125))
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L1
	} else {
		goto L56
	}
L8:
	;
	v36 = int32(3)
	goto L10
L9:
	;
	v36 = int32(1)
	goto L10
L10:
	;
	if l2 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v39 = int32(_a_F_composite_to_json_0)
	goto L13
L12:
	;
	v39 = int32(_a_F_composite_to_json_1)
	goto L13
L13:
	;
	v43 = int32(0)
	v47 = v31
	v49 = int32(0)
	goto L14
L14:
	;
	v60 = v21 + v47<<(uint(int32(4))%32) + v43*int32(100)
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+111)))
	if v61 == int32(1) {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	goto L7
L16:
	;
	v186 = v47
	v187 = v49
	v189 = v43 + int32(1)
	goto L18
L17:
	;
	if v49 != 0 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	if v189 < v186 {
		v43 = v189
		v47 = v186
		v49 = v187
		goto L14
	} else {
		goto L55
	}
L19:
	;
	F_appendBinaryStringInfo(m, l1, v39, v36)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L1
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	F_escape_json(m, l1, v60+int32(24))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L1
	} else {
		goto L23
	}
L22:
	;
	goto L21
L23:
	;
	F_appendStringInfoChar(m, l1, int32(58))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	v78 = v43 + int32(1)
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v15)+44))
	v80 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v79)+18)))
	if base.Ui32(v80&int32(2047)) <= base.Ui32(v43) {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	v153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+27)))
	if v153 == int32(1) {
		goto L50
	} else {
		goto L51
	}
L26:
	;
	v86 = F_getmissingattr(m, v21, v78, v15+int32(27))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L1
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	v88 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+27)) = uint8(v88)
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79)+20)))
	if v90&int32(1) == v88 {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	v152 = v86
	goto L25
L30:
	;
	v97 = v21 + int32(4) + v78<<(uint(int32(4))%32)
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v97)))
	if int32(0) <= v98 {
		goto L33
	} else {
		goto L34
	}
L31:
	;
	goto L32
L32:
	;
	v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79+int32(base.Ui32(v43)>>(uint(int32(3))%32)))+23)))
	if int32(base.Ui32(v134)>>(uint(v43&int32(7))%32))&int32(1) == int32(0) {
		goto L45
	} else {
		goto L46
	}
L33:
	;
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79)+22)))
	v103 = v79 + v101 + v98
	v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97)+6)))
	if v104 != int32(1) {
		v152 = v103
		goto L25
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	v129 = F_nocachegetattr(m, v15+int32(28), v78, v21)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L1
	} else {
		goto L44
	}
L36:
	;
	v107 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v97)+4)))
	switch v107 - int32(1) {
	case 0:
		goto L40
	case 1:
		goto L39
	default:
		goto L37
	case 3:
		goto L38
	}
L37:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L1
	} else {
		goto L41
	}
L38:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v103)))
	v152 = v112
	goto L25
L39:
	;
	v111 = int32(*(*int16)(unsafe.Add(mBase, uint32(v103))))
	v152 = v111
	goto L25
L40:
	;
	v110 = int32(*(*int8)(unsafe.Add(mBase, uint32(v103))))
	v152 = v110
	goto L25
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = base.I32_extend16_s(v107)
	F_errmsg_internal(m, int32(_a_F_composite_to_json_2), v15)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	F_errfinish(m, int32(_a_F_composite_to_json_3), int32(70), int32(_a_F_composite_to_json_4))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L44:
	;
	v152 = v129
	goto L25
L45:
	;
	v142 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+27)) = uint8(v142)
	v152 = int32(0)
	goto L25
L46:
	;
	goto L47
L47:
	;
	v147 = F_nocachegetattr(m, v15+int32(28), v78, v21)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	v152 = v147
	goto L25
L49:
	;
	v177 = int32(1)
	F_datum_to_json_internal(m, v152, v174&v177, l1, v175, v176, int32(0))
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L1
	} else {
		goto L54
	}
L50:
	;
	v156 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v156
	*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = v156
	v174 = int32(1)
	v175 = v156
	v176 = v156
	goto L49
L51:
	;
	goto L52
L52:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v60+int32(20))+68))
	F_json_categorize_type(m, v163, int32(0), v15+int32(20), v15+int32(16))
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	v171 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+27)))
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v15)+20))
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v15)+16))
	v174 = v171
	v175 = v172
	v176 = v173
	goto L49
L54:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	v186 = v183
	v187 = v177
	v189 = v78
	goto L18
L55:
	;
	goto L15
L56:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
	if int32(0) <= v206 {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	F_DecrTupleDescRefCount(m, v21)
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L1
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	m.G0 = v15 + int32(48)
	return
L60:
	;
	goto L59
}
func F_compress_process(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
	v8 = m.Env.Pgmem_deflate_write(m, v7, l2, l3)
	mBase = m.M
	if v8 < int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(-105)
L2:
	;
	goto L3
L3:
	;
	v14 = l1 + int32(8)
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	goto L4
L4:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v23 = m.Env.Pgmem_zstream_read(m, v22, v14, v15)
	mBase = m.M
	if v23 < int32(0) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	return v32
L6:
	;
	return int32(-105)
L7:
	;
	goto L8
L8:
	;
	if v23 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	return int32(0)
L10:
	;
	goto L11
L11:
	;
	v32 = F_pushf_write(m, l0, v14, v23)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	return int32(0)
L13:
	;
	if int32(0) <= v32 {
		goto L4
	} else {
		goto L14
	}
L14:
	;
	goto L5
}
func F_computeDistance(m *base.Module, l0 int32, l1 int32, l2 int32) float64 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 float64
	_ = v21
	var v22 float64
	_ = v22
	var v23 float64
	_ = v23
	var v27 float64
	_ = v27
	var v31 float64
	_ = v31
	var v32 float64
	_ = v32
	var v34 float64
	_ = v34
	var v36 float64
	_ = v36
	var v37 float64
	_ = v37
	var v40 float64
	_ = v40
	var v42 float64
	_ = v42
	var v51 float64
	_ = v51
	var v54 float64
	_ = v54
	var v56 float64
	_ = v56
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v77 float64
	_ = v77
	var v78 float64
	_ = v78
	var v82 float64
	_ = v82
	var v88 float64
	_ = v88
	var v90 float64
	_ = v90
	var v98 float64
	_ = v98
	var v101 float64
	_ = v101
	var v103 float64
	_ = v103
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 float64
	_ = v130
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 float64
	_ = v135
	var v136 float64
	_ = v136
	var v138 float64
	_ = v138
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 float64
	_ = v144
	var v145 float64
	_ = v145
	var v147 float64
	_ = v147
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 float64
	_ = v153
	var v155 float64
	_ = v155
	var v157 float64
	_ = v157
	var v159 float64
	_ = v159
	var v162 float64
	_ = v162
	var v172 int32
	_ = v172
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	if l0 != 0 {
		v17 = F_DirectFunctionCall2Coll(m, int32(112), int32(0), l2, l1+int32(16))
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return float64(0)
		} else {
			v21 = *(*float64)(unsafe.Add(mBase, uint32(v17)))
			v162 = v21
			m.G0 = v11 + int32(16)
			return v162
		}
	} else {
		v22 = *(*float64)(unsafe.Add(mBase, uint32(l2)))
		v23 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
		if base.F64_le(v22, v23) == int32(0) {
			v77 = *(*float64)(unsafe.Add(mBase, uint32(l2)+8))
			v78 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
			if base.F64_le(v77, v78) == int32(0) {
				v128 = F_DirectFunctionCall2Coll(m, int32(112), int32(0), l2, l1+int32(16))
				mBase = m.M
				v129 = m.ExcPending
				if v129 != 0 {
					return float64(0)
				} else {
					v130 = *(*float64)(unsafe.Add(mBase, uint32(v128)))
					v133 = F_DirectFunctionCall2Coll(m, int32(112), int32(0), l2, l1)
					mBase = m.M
					v134 = m.ExcPending
					if v134 != 0 {
						return float64(0)
					} else {
						v135 = *(*float64)(unsafe.Add(mBase, uint32(v133)))
						v136 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
						*(*float64)(unsafe.Add(mBase, uint32(v11))) = v136
						v138 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
						*(*float64)(unsafe.Add(mBase, uint32(v11)+8)) = v138
						v142 = F_DirectFunctionCall2Coll(m, int32(112), int32(0), l2, v11)
						mBase = m.M
						v143 = m.ExcPending
						if v143 != 0 {
							return float64(0)
						} else {
							v144 = *(*float64)(unsafe.Add(mBase, uint32(v142)))
							v145 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
							*(*float64)(unsafe.Add(mBase, uint32(v11))) = v145
							v147 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
							*(*float64)(unsafe.Add(mBase, uint32(v11)+8)) = v147
							v151 = F_DirectFunctionCall2Coll(m, int32(112), int32(0), l2, v11)
							mBase = m.M
							v152 = m.ExcPending
							if v152 != 0 {
								return float64(0)
							} else {
								v153 = *(*float64)(unsafe.Add(mBase, uint32(v151)))
								if base.F64_gt(v130, v135) != 0 {
									v155 = v135
								} else {
									v155 = v130
								}
								if base.F64_gt(v155, v144) != 0 {
									v157 = v144
								} else {
									v157 = v155
								}
								if base.F64_gt(v157, v153) != 0 {
									v159 = v153
								} else {
									v159 = v157
								}
								v162 = v159
								m.G0 = v11 + int32(16)
								return v162
							}
						}
					}
				}
			} else {
				v82 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
				if base.F64_ge(v77, v82) == int32(0) {
					v128 = F_DirectFunctionCall2Coll(m, int32(112), int32(0), l2, l1+int32(16))
					mBase = m.M
					v129 = m.ExcPending
					if v129 != 0 {
						return float64(0)
					} else {
						v130 = *(*float64)(unsafe.Add(mBase, uint32(v128)))
						v133 = F_DirectFunctionCall2Coll(m, int32(112), int32(0), l2, l1)
						mBase = m.M
						v134 = m.ExcPending
						if v134 != 0 {
							return float64(0)
						} else {
							v135 = *(*float64)(unsafe.Add(mBase, uint32(v133)))
							v136 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
							*(*float64)(unsafe.Add(mBase, uint32(v11))) = v136
							v138 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
							*(*float64)(unsafe.Add(mBase, uint32(v11)+8)) = v138
							v142 = F_DirectFunctionCall2Coll(m, int32(112), int32(0), l2, v11)
							mBase = m.M
							v143 = m.ExcPending
							if v143 != 0 {
								return float64(0)
							} else {
								v144 = *(*float64)(unsafe.Add(mBase, uint32(v142)))
								v145 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
								*(*float64)(unsafe.Add(mBase, uint32(v11))) = v145
								v147 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
								*(*float64)(unsafe.Add(mBase, uint32(v11)+8)) = v147
								v151 = F_DirectFunctionCall2Coll(m, int32(112), int32(0), l2, v11)
								mBase = m.M
								v152 = m.ExcPending
								if v152 != 0 {
									return float64(0)
								} else {
									v153 = *(*float64)(unsafe.Add(mBase, uint32(v151)))
									if base.F64_gt(v130, v135) != 0 {
										v155 = v135
									} else {
										v155 = v130
									}
									if base.F64_gt(v155, v144) != 0 {
										v157 = v144
									} else {
										v157 = v155
									}
									if base.F64_gt(v157, v153) != 0 {
										v159 = v153
									} else {
										v159 = v157
									}
									v162 = v159
									m.G0 = v11 + int32(16)
									return v162
								}
							}
						}
					}
				} else {
					if base.F64_gt(v22, v23) != 0 {
						v88 = math.Float64frombits(uint64(0x7ff0000000000000))
						v90 = base.F64_sub(v22, v23)
						if base.F64_eq(base.F64_abs(v22), v88)|base.F64_ne(base.F64_abs(v90), v88) != 0 {
							v162 = v90
							m.G0 = v11 + int32(16)
							return v162
						} else {
							if base.F64_ne(base.F64_abs(v23), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
								F_float_overflow_error(m)
								mBase = m.M
								v172 = m.ExcPending
								if v172 != 0 {
									return float64(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							} else {
								v162 = v90
								m.G0 = v11 + int32(16)
								return v162
							}
						}
					} else {
						v98 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
						if base.F64_gt(v98, v22) != 0 {
							v101 = math.Float64frombits(uint64(0x7ff0000000000000))
							v103 = base.F64_sub(v98, v22)
							if base.F64_eq(base.F64_abs(v22), v101)|base.F64_ne(base.F64_abs(v103), v101) != 0 {
								v162 = v103
								m.G0 = v11 + int32(16)
								return v162
							} else {
								if base.F64_ne(base.F64_abs(v98), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
									F_float_overflow_error(m)
									mBase = m.M
									v172 = m.ExcPending
									if v172 != 0 {
										return float64(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								} else {
									v162 = v103
									m.G0 = v11 + int32(16)
									return v162
								}
							}
						} else {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v114 = m.ExcPending
							if v114 != 0 {
								return float64(0)
							} else {
								F_errmsg_internal(m, int32(_a_F_computeDistance_0), int32(0))
								mBase = m.M
								v118 = m.ExcPending
								if v118 != 0 {
									return float64(0)
								} else {
									F_errfinish(m, int32(_a_F_computeDistance_1), int32(1256), int32(_a_F_computeDistance_2))
									mBase = m.M
									v123 = m.ExcPending
									if v123 != 0 {
										return float64(0)
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
		} else {
			v27 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
			if base.F64_ge(v22, v27) == int32(0) {
				v77 = *(*float64)(unsafe.Add(mBase, uint32(l2)+8))
				v78 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
				if base.F64_le(v77, v78) == int32(0) {
					v128 = F_DirectFunctionCall2Coll(m, int32(112), int32(0), l2, l1+int32(16))
					mBase = m.M
					v129 = m.ExcPending
					if v129 != 0 {
						return float64(0)
					} else {
						v130 = *(*float64)(unsafe.Add(mBase, uint32(v128)))
						v133 = F_DirectFunctionCall2Coll(m, int32(112), int32(0), l2, l1)
						mBase = m.M
						v134 = m.ExcPending
						if v134 != 0 {
							return float64(0)
						} else {
							v135 = *(*float64)(unsafe.Add(mBase, uint32(v133)))
							v136 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
							*(*float64)(unsafe.Add(mBase, uint32(v11))) = v136
							v138 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
							*(*float64)(unsafe.Add(mBase, uint32(v11)+8)) = v138
							v142 = F_DirectFunctionCall2Coll(m, int32(112), int32(0), l2, v11)
							mBase = m.M
							v143 = m.ExcPending
							if v143 != 0 {
								return float64(0)
							} else {
								v144 = *(*float64)(unsafe.Add(mBase, uint32(v142)))
								v145 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
								*(*float64)(unsafe.Add(mBase, uint32(v11))) = v145
								v147 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
								*(*float64)(unsafe.Add(mBase, uint32(v11)+8)) = v147
								v151 = F_DirectFunctionCall2Coll(m, int32(112), int32(0), l2, v11)
								mBase = m.M
								v152 = m.ExcPending
								if v152 != 0 {
									return float64(0)
								} else {
									v153 = *(*float64)(unsafe.Add(mBase, uint32(v151)))
									if base.F64_gt(v130, v135) != 0 {
										v155 = v135
									} else {
										v155 = v130
									}
									if base.F64_gt(v155, v144) != 0 {
										v157 = v144
									} else {
										v157 = v155
									}
									if base.F64_gt(v157, v153) != 0 {
										v159 = v153
									} else {
										v159 = v157
									}
									v162 = v159
									m.G0 = v11 + int32(16)
									return v162
								}
							}
						}
					}
				} else {
					v82 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
					if base.F64_ge(v77, v82) == int32(0) {
						v128 = F_DirectFunctionCall2Coll(m, int32(112), int32(0), l2, l1+int32(16))
						mBase = m.M
						v129 = m.ExcPending
						if v129 != 0 {
							return float64(0)
						} else {
							v130 = *(*float64)(unsafe.Add(mBase, uint32(v128)))
							v133 = F_DirectFunctionCall2Coll(m, int32(112), int32(0), l2, l1)
							mBase = m.M
							v134 = m.ExcPending
							if v134 != 0 {
								return float64(0)
							} else {
								v135 = *(*float64)(unsafe.Add(mBase, uint32(v133)))
								v136 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
								*(*float64)(unsafe.Add(mBase, uint32(v11))) = v136
								v138 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
								*(*float64)(unsafe.Add(mBase, uint32(v11)+8)) = v138
								v142 = F_DirectFunctionCall2Coll(m, int32(112), int32(0), l2, v11)
								mBase = m.M
								v143 = m.ExcPending
								if v143 != 0 {
									return float64(0)
								} else {
									v144 = *(*float64)(unsafe.Add(mBase, uint32(v142)))
									v145 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
									*(*float64)(unsafe.Add(mBase, uint32(v11))) = v145
									v147 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
									*(*float64)(unsafe.Add(mBase, uint32(v11)+8)) = v147
									v151 = F_DirectFunctionCall2Coll(m, int32(112), int32(0), l2, v11)
									mBase = m.M
									v152 = m.ExcPending
									if v152 != 0 {
										return float64(0)
									} else {
										v153 = *(*float64)(unsafe.Add(mBase, uint32(v151)))
										if base.F64_gt(v130, v135) != 0 {
											v155 = v135
										} else {
											v155 = v130
										}
										if base.F64_gt(v155, v144) != 0 {
											v157 = v144
										} else {
											v157 = v155
										}
										if base.F64_gt(v157, v153) != 0 {
											v159 = v153
										} else {
											v159 = v157
										}
										v162 = v159
										m.G0 = v11 + int32(16)
										return v162
									}
								}
							}
						}
					} else {
						if base.F64_gt(v22, v23) != 0 {
							v88 = math.Float64frombits(uint64(0x7ff0000000000000))
							v90 = base.F64_sub(v22, v23)
							if base.F64_eq(base.F64_abs(v22), v88)|base.F64_ne(base.F64_abs(v90), v88) != 0 {
								v162 = v90
								m.G0 = v11 + int32(16)
								return v162
							} else {
								if base.F64_ne(base.F64_abs(v23), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
									F_float_overflow_error(m)
									mBase = m.M
									v172 = m.ExcPending
									if v172 != 0 {
										return float64(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								} else {
									v162 = v90
									m.G0 = v11 + int32(16)
									return v162
								}
							}
						} else {
							v98 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
							if base.F64_gt(v98, v22) != 0 {
								v101 = math.Float64frombits(uint64(0x7ff0000000000000))
								v103 = base.F64_sub(v98, v22)
								if base.F64_eq(base.F64_abs(v22), v101)|base.F64_ne(base.F64_abs(v103), v101) != 0 {
									v162 = v103
									m.G0 = v11 + int32(16)
									return v162
								} else {
									if base.F64_ne(base.F64_abs(v98), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
										F_float_overflow_error(m)
										mBase = m.M
										v172 = m.ExcPending
										if v172 != 0 {
											return float64(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									} else {
										v162 = v103
										m.G0 = v11 + int32(16)
										return v162
									}
								}
							} else {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v114 = m.ExcPending
								if v114 != 0 {
									return float64(0)
								} else {
									F_errmsg_internal(m, int32(_a_F_computeDistance_0), int32(0))
									mBase = m.M
									v118 = m.ExcPending
									if v118 != 0 {
										return float64(0)
									} else {
										F_errfinish(m, int32(_a_F_computeDistance_1), int32(1256), int32(_a_F_computeDistance_2))
										mBase = m.M
										v123 = m.ExcPending
										if v123 != 0 {
											return float64(0)
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
			} else {
				v31 = *(*float64)(unsafe.Add(mBase, uint32(l2)+8))
				v32 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
				if base.F64_le(v31, v32) != 0 {
					v34 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
					if base.F64_ge(v31, v34) != 0 {
						v162 = float64(0)
						m.G0 = v11 + int32(16)
						return v162
					} else {
						v36 = *(*float64)(unsafe.Add(mBase, uint32(l2)+8))
						v37 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
						if base.F64_gt(v36, v37) != 0 {
							v40 = math.Float64frombits(uint64(0x7ff0000000000000))
							v42 = base.F64_sub(v36, v37)
							if base.F64_eq(base.F64_abs(v36), v40)|base.F64_ne(base.F64_abs(v42), v40)|base.F64_eq(base.F64_abs(v37), v40) != 0 {
								v162 = v42
								m.G0 = v11 + int32(16)
								return v162
							} else {
								F_float_overflow_error(m)
								mBase = m.M
								v172 = m.ExcPending
								if v172 != 0 {
									return float64(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						} else {
							v51 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
							if base.F64_gt(v51, v36) != 0 {
								v54 = math.Float64frombits(uint64(0x7ff0000000000000))
								v56 = base.F64_sub(v51, v36)
								if base.F64_eq(base.F64_abs(v36), v54)|base.F64_ne(base.F64_abs(v56), v54) != 0 {
									v162 = v56
									m.G0 = v11 + int32(16)
									return v162
								} else {
									if base.F64_ne(base.F64_abs(v51), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
										F_float_overflow_error(m)
										mBase = m.M
										v172 = m.ExcPending
										if v172 != 0 {
											return float64(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									} else {
										v162 = v56
										m.G0 = v11 + int32(16)
										return v162
									}
								}
							} else {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v67 = m.ExcPending
								if v67 != 0 {
									return float64(0)
								} else {
									F_errmsg_internal(m, int32(_a_F_computeDistance_0), int32(0))
									mBase = m.M
									v71 = m.ExcPending
									if v71 != 0 {
										return float64(0)
									} else {
										F_errfinish(m, int32(_a_F_computeDistance_1), int32(1245), int32(_a_F_computeDistance_2))
										mBase = m.M
										v76 = m.ExcPending
										if v76 != 0 {
											return float64(0)
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
				} else {
					v36 = *(*float64)(unsafe.Add(mBase, uint32(l2)+8))
					v37 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
					if base.F64_gt(v36, v37) != 0 {
						v40 = math.Float64frombits(uint64(0x7ff0000000000000))
						v42 = base.F64_sub(v36, v37)
						if base.F64_eq(base.F64_abs(v36), v40)|base.F64_ne(base.F64_abs(v42), v40)|base.F64_eq(base.F64_abs(v37), v40) != 0 {
							v162 = v42
							m.G0 = v11 + int32(16)
							return v162
						} else {
							F_float_overflow_error(m)
							mBase = m.M
							v172 = m.ExcPending
							if v172 != 0 {
								return float64(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					} else {
						v51 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
						if base.F64_gt(v51, v36) != 0 {
							v54 = math.Float64frombits(uint64(0x7ff0000000000000))
							v56 = base.F64_sub(v51, v36)
							if base.F64_eq(base.F64_abs(v36), v54)|base.F64_ne(base.F64_abs(v56), v54) != 0 {
								v162 = v56
								m.G0 = v11 + int32(16)
								return v162
							} else {
								if base.F64_ne(base.F64_abs(v51), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
									F_float_overflow_error(m)
									mBase = m.M
									v172 = m.ExcPending
									if v172 != 0 {
										return float64(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								} else {
									v162 = v56
									m.G0 = v11 + int32(16)
									return v162
								}
							}
						} else {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v67 = m.ExcPending
							if v67 != 0 {
								return float64(0)
							} else {
								F_errmsg_internal(m, int32(_a_F_computeDistance_0), int32(0))
								mBase = m.M
								v71 = m.ExcPending
								if v71 != 0 {
									return float64(0)
								} else {
									F_errfinish(m, int32(_a_F_computeDistance_1), int32(1245), int32(_a_F_computeDistance_2))
									mBase = m.M
									v76 = m.ExcPending
									if v76 != 0 {
										return float64(0)
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
	}
}
func F_computeLeafRecompressWALData(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	v2 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if base.B2i32(v14 == v2)|base.B2i32(l0 == v14) == v2 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v22 = v2
	v23 = v14
	goto L4
L2:
	;
	v37 = v2
	goto L3
L3:
	;
	v49 = F_palloc(m, v37<<(uint(int32(1))%32)+int32(_a_F_computeLeafRecompressWALData_0))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+8)))
	v33 = v22 + base.B2i32(v30 != int32(0))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	if v34 != l0 {
		v22 = v33
		v23 = v34
		goto L4
	} else {
		goto L6
	}
L5:
	;
	v37 = v33
	goto L3
L6:
	;
	goto L5
L7:
	;
	return
L8:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v49))) = uint16(v37)
	v53 = v49 + int32(2)
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v55 = int32(0)
	if base.B2i32(v54 == v55)|base.B2i32(l0 == v54) == v55 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v62 = v53
	v63 = v54
	v66 = v2
	goto L12
L10:
	;
	v155 = v53
	goto L11
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v155 - v49
	m.G0 = v12 + int32(16)
	return
L12:
	;
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63)+8)))
	switch v70 {
	case 0:
		goto L18
	case 1:
		goto L16
	default:
		goto L17
	}
L13:
	;
	v155 = v147
	goto L11
L14:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v63)+4))
	if v152 != l0 {
		v62 = v147
		v63 = v152
		v66 = v151
		goto L12
	} else {
		goto L37
	}
L15:
	;
	v147 = v142 + v62 + int32(2)
	v151 = v139 + v66
	goto L14
L16:
	;
	v134 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v62)+1)) = uint8(v134)
	*(*uint8)(unsafe.Add(mBase, uint32(v62))) = uint8(v66)
	v139 = v134
	v142 = int32(0)
	goto L15
L17:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v63)+20))
	v74 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v73)+6)))
	v78 = (v74 + int32(1)) & int32(_a_F_computeLeafRecompressWALData_1)
	v80 = v78 + int32(8)
	if v70 == int32(4) {
		goto L21
	} else {
		goto L22
	}
L18:
	;
	v147 = v62
	v151 = v66 + int32(1)
	goto L14
L19:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L7
	} else {
		goto L34
	}
L20:
	;
	if v80 != 0 {
		goto L31
	} else {
		goto L32
	}
L21:
	;
	v83 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v63)+16)))
	*(*uint8)(unsafe.Add(mBase, uint32(v62))) = uint8(v66)
	if base.Ui32(v83*int32(6)) <= base.Ui32(v80) {
		goto L24
	} else {
		goto L25
	}
L22:
	;
	goto L23
L23:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v62)+1)) = uint8(v70)
	*(*uint8)(unsafe.Add(mBase, uint32(v62))) = uint8(v66)
	if v70&int32(254) != int32(2) {
		goto L19
	} else {
		goto L30
	}
L24:
	;
	v88 = int32(4)
	*(*uint8)(unsafe.Add(mBase, uint32(v62)+1)) = uint8(v88)
	v90 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v63)+16)))
	*(*uint16)(unsafe.Add(mBase, uint32(v62)+2)) = uint16(v90)
	v93 = v90 * int32(6)
	if v93 != 0 {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	goto L26
L26:
	;
	v101 = int32(3)
	*(*uint8)(unsafe.Add(mBase, uint32(v62)+1)) = uint8(v101)
	v110 = v101
	goto L20
L27:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v63)+12))
	base.MemoryCopy(m, v62+int32(4), v96, v93)
	goto L29
L28:
	;
	goto L29
L29:
	;
	v139 = int32(1)
	v142 = v93 + int32(2)
	goto L15
L30:
	;
	v110 = v70
	goto L20
L31:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v63)+20))
	base.MemoryCopy(m, v62+int32(2), v113, v80)
	goto L33
L32:
	;
	goto L33
L33:
	;
	v139 = base.B2i32(v110 != int32(2))
	v142 = (v78 + int32(9)) & int32(_a_F_computeLeafRecompressWALData_2)
	goto L15
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v70
	F_errmsg_internal(m, int32(_a_F_computeLeafRecompressWALData_3), v12)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L7
	} else {
		goto L35
	}
L35:
	;
	F_errfinish(m, int32(_a_F_computeLeafRecompressWALData_4), int32(955), int32(_a_F_computeLeafRecompressWALData_5))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L7
	} else {
		goto L36
	}
L36:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L37:
	;
	goto L13
}
func F_connect(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v15 int32
	_ = v15
	v4 = int32(0)
	v7 = m.Env.X__syscall_connect(m, l0, int32(_a_F_connect_0), int32(12), v4, v4, v4)
	mBase = m.M
	if base.Ui32(int32(-4095)) <= base.Ui32(v7) {
		*(*int32)(unsafe.Add(mBase, _c_F_connect[0])) = int32(0) - v7
		v15 = int32(-1)
	} else {
		v15 = v7
	}
	return v15
}
func F_consider_index_join_outer_rels(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v22 int32
	_ = v22
	var v39 int32
	_ = v39
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v77 int32
	_ = v77
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v164 int32
	_ = v164
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v184 int32
	_ = v184
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v201 int32
	_ = v201
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v242 int32
	_ = v242
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v267 int32
	_ = v267
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v328 int32
	_ = v328
	var v350 int32
	_ = v350
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	v11 = int32(0)
	if l7 == v11 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l7)+4))
	if v22 <= int32(0) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v39 = v11
	goto L4
L4:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l7)+12))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v46+v39<<(uint(int32(2))%32))))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)+60))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l9)))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v51)+28))
	v55 = F_list_member(m, v53, v54)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L1
L6:
	;
	return
L7:
	;
	if v55 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l9)))
	if v59 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L10
L10:
	;
	v371 = v39 + int32(1)
	v372 = *(*int32)(unsafe.Add(mBase, uint32(l7)+4))
	if v371 < v372 {
		v39 = v371
		goto L4
	} else {
		goto L85
	}
L11:
	;
	F_get_join_index_paths(m, l0, l1, l2, l3, l4, l5, l6, v54, l9)
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L6
	} else {
		goto L84
	}
L12:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v59)+4))
	if v62 <= int32(0) {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v77 = int32(0)
	goto L14
L14:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l9)))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)+12))
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v86+v77<<(uint(int32(2))%32))))
	if v54 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L15:
	;
	goto L11
L16:
	;
	v328 = v77 + int32(1)
	if v328 != v62 {
		v77 = v328
		goto L14
	} else {
		goto L83
	}
L17:
	;
	if v184 != int32(3) {
		goto L16
	} else {
		goto L52
	}
L18:
	;
	v184 = base.B2i32(v90 != int32(0))
	goto L17
L19:
	;
	goto L20
L20:
	;
	if v90 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v184 = int32(2)
	goto L17
L22:
	;
	goto L23
L23:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v54)+4))
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v90)+4))
	if v107 < v108 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v110 = v107
	goto L26
L25:
	;
	v110 = v108
	goto L26
L26:
	;
	if v110 <= int32(1) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v113 = int32(1)
	goto L29
L28:
	;
	v113 = v110
	goto L29
L29:
	;
	v114 = int32(8)
	v118 = int32(0)
	v120 = v118
	v121 = v118
	goto L32
L30:
	;
	v184 = int32(3)
	goto L17
L31:
	;
	v184 = v171
	goto L17
L32:
	;
	v131 = v121 << (uint(int32(2)) % 32)
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v54+v114+v131)))
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v131+(v90+v114))))
	if v133&(v135^int32(-1)) != 0 {
		goto L35
	} else {
		goto L36
	}
L33:
	;
	if v108 < v107 {
		goto L42
	} else {
		goto L43
	}
L34:
	;
	v157 = v121 + int32(1)
	if v157 != v113 {
		v120 = v155
		v121 = v157
		goto L32
	} else {
		goto L41
	}
L35:
	;
	if base.B2i32(v120 == int32(1))|v135&(v133^int32(-1)) != 0 {
		v171 = int32(3)
		goto L31
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	if v135&(v133^int32(-1)) == int32(0) {
		v155 = v120
		goto L34
	} else {
		goto L39
	}
L38:
	;
	v155 = int32(2)
	goto L34
L39:
	;
	if v120 == int32(2) {
		goto L30
	} else {
		goto L40
	}
L40:
	;
	v155 = int32(1)
	goto L34
L41:
	;
	goto L33
L42:
	;
	if v155 == int32(1) {
		goto L45
	} else {
		goto L46
	}
L43:
	;
	goto L44
L44:
	;
	if v108 <= v107 {
		v171 = v155
		goto L31
	} else {
		goto L48
	}
L45:
	;
	v164 = int32(3)
	goto L47
L46:
	;
	v164 = int32(2)
	goto L47
L47:
	;
	v184 = v164
	goto L17
L48:
	;
	if v155 == int32(2) {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v170 = int32(3)
	goto L51
L50:
	;
	v170 = int32(1)
	goto L51
L51:
	;
	v171 = v170
	goto L31
L52:
	;
	if v52 == int32(0) {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v299 = *(*int32)(unsafe.Add(mBase, uint32(l9)))
	if v299 != 0 {
		goto L77
	} else {
		goto L78
	}
L54:
	;
	v189 = int32(0)
	v190 = *(*int32)(unsafe.Add(mBase, uint32(l7)+4))
	if v190 <= v189 {
		goto L53
	} else {
		goto L55
	}
L55:
	;
	v201 = v189
	v208 = v190
	goto L56
L56:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(l7)+12))
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v212+v201<<(uint(int32(2))%32))))
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v216)+4))
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v217)+60))
	if v52 == v218 {
		goto L58
	} else {
		goto L59
	}
L57:
	;
	goto L53
L58:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v217)+28))
	v221 = int32(0)
	if v220 == v221 {
		goto L62
	} else {
		goto L63
	}
L59:
	;
	v276 = v208
	goto L60
L60:
	;
	v278 = v201 + int32(1)
	if v278 < v276 {
		v201 = v278
		v208 = v276
		goto L56
	} else {
		goto L76
	}
L61:
	;
	if v274 != 0 {
		goto L16
	} else {
		goto L75
	}
L62:
	;
	v274 = int32(1)
	goto L61
L63:
	;
	goto L64
L64:
	;
	if v90 == int32(0) {
		v267 = v221
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v274 = v267
	goto L61
L66:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v220)+4))
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v90)+4))
	if v231 < v230 {
		v267 = v221
		goto L65
	} else {
		goto L67
	}
L67:
	;
	v233 = int32(1)
	if v230 <= v233 {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v236 = v233
	goto L70
L69:
	;
	v236 = v230
	goto L70
L70:
	;
	v237 = int32(8)
	v242 = int32(0)
	goto L71
L71:
	;
	v249 = v242 << (uint(int32(2)) % 32)
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v220+v237+v249)))
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v90+v237+v249)))
	v256 = v251 & (v253 ^ int32(-1))
	v258 = base.B2i32(v256 == int32(0))
	if v256 != 0 {
		v267 = v258
		goto L65
	} else {
		goto L73
	}
L72:
	;
	v267 = v258
	goto L65
L73:
	;
	v260 = v242 + int32(1)
	if v260 != v236 {
		v242 = v260
		goto L71
	} else {
		goto L74
	}
L74:
	;
	goto L72
L75:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(l7)+4))
	v276 = v275
	goto L60
L76:
	;
	goto L57
L77:
	;
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v299)+4))
	v302 = v300
	goto L79
L78:
	;
	v302 = int32(0)
	goto L79
L79:
	;
	if l8*int32(10) <= v302 {
		goto L11
	} else {
		goto L80
	}
L80:
	;
	v304 = F_bms_union(m, v54, v90)
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L6
	} else {
		goto L81
	}
L81:
	;
	F_get_join_index_paths(m, l0, l1, l2, l3, l4, l5, l6, v304, l9)
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L6
	} else {
		goto L82
	}
L82:
	;
	goto L16
L83:
	;
	goto L15
L84:
	;
	goto L10
L85:
	;
	goto L5
}
func F_conv_utf8_to(m *base.Module, l0 int32) int32 {
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v385 int32
	_ = v385
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v398 int32
	_ = v398
	var v400 int32
	_ = v400
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v424 int32
	_ = v424
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v437 int32
	_ = v437
	var v439 int32
	_ = v439
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v463 int32
	_ = v463
	var v473 int32
	_ = v473
	var v477 int32
	_ = v477
	var v481 int32
	_ = v481
	var v493 int32
	_ = v493
	if base.Ui32(l0) < base.Ui32(int32(128)) {
		v45 = l0
	} else {
		if base.Ui32(l0) <= base.Ui32(int32(_a_F_conv_utf8_to_0)) {
			v45 = int32(base.Ui32(l0)>>(uint(int32(2))%32))&int32(1984) | l0&int32(63)
		} else {
			if base.Ui32(l0) <= base.Ui32(int32(16777215)) {
				v45 = int32(base.Ui32(l0)>>(uint(int32(4))%32))&int32(_a_F_conv_utf8_to_1) | (int32(base.Ui32(l0)>>(uint(int32(2))%32))&int32(4032) | l0&int32(63))
			} else {
				v45 = int32(base.Ui32(l0)>>(uint(int32(2))%32))&int32(4032) | (int32(base.Ui32(l0)>>(uint(int32(6))%32))&int32(_a_F_conv_utf8_to_2) | (int32(base.Ui32(l0)>>(uint(int32(4))%32))&int32(_a_F_conv_utf8_to_3) | l0&int32(63)))
			}
		}
	}
	if base.Ui32(v45-int32(1106)) <= base.Ui32(int32(_a_F_conv_utf8_to_4)) {
		v51 = v45 - int32(286)
		v52 = int32(_a_F_conv_utf8_to_0)
		v53 = v51 & v52
		v55 = base.I32_div_u_s(v53, int32(1260))
		v58 = int32(10)
		v59 = base.I32_div_u_s(v53, v58)
		v61 = base.I32_rem_u_s(v59, int32(126))
		return v55<<(uint(int32(16))%32) | (v61<<(uint(int32(8))%32) + (v51-v59*v58)&v52 + int32(_a_F_conv_utf8_to_5)) | int32(-2127560656)
	} else {
		if base.Ui32(v45-int32(_a_F_conv_utf8_to_6)) <= base.Ui32(int32(2109)) {
			v81 = v45 - int32(576)
			v82 = int32(_a_F_conv_utf8_to_0)
			v83 = v81 & v82
			v85 = base.I32_div_u_s(v83, int32(1260))
			v88 = int32(10)
			v89 = base.I32_div_u_s(v83, v88)
			v91 = base.I32_rem_u_s(v89, int32(126))
			return v85<<(uint(int32(16))%32) | (v91<<(uint(int32(8))%32) + (v81-v89*v88)&v82 + int32(_a_F_conv_utf8_to_5)) | int32(-2127560656)
		} else {
			if base.Ui32(v45-int32(_a_F_conv_utf8_to_7)) <= base.Ui32(int32(764)) {
				v111 = v45 - int32(878)
				v112 = int32(_a_F_conv_utf8_to_0)
				v113 = v111 & v112
				v115 = base.I32_div_u_s(v113, int32(1260))
				v120 = int32(10)
				v121 = base.I32_div_u_s(v113, v120)
				v123 = base.I32_rem_u_s(v121, int32(126))
				return v115<<(uint(int32(16))%32) + int32(2146828288) | (v123<<(uint(int32(8))%32) + (v111-v121*v120)&v112 + int32(_a_F_conv_utf8_to_5)) | int32(-2110783440)
			} else {
				if base.Ui32(v45-int32(_a_F_conv_utf8_to_8)) <= base.Ui32(int32(884)) {
					v143 = v45 - int32(887)
					v144 = int32(_a_F_conv_utf8_to_0)
					v145 = v143 & v144
					v147 = base.I32_div_u_s(v145, int32(1260))
					v152 = int32(10)
					v153 = base.I32_div_u_s(v145, v152)
					v155 = base.I32_rem_u_s(v153, int32(126))
					return v147<<(uint(int32(16))%32) + int32(2146828288) | (v155<<(uint(int32(8))%32) + (v143-v153*v152)&v144 + int32(_a_F_conv_utf8_to_5)) | int32(-2110783440)
				} else {
					if base.Ui32(v45-int32(_a_F_conv_utf8_to_9)) <= base.Ui32(int32(470)) {
						v175 = v45 - int32(889)
						v176 = int32(_a_F_conv_utf8_to_0)
						v177 = v175 & v176
						v179 = base.I32_div_u_s(v177, int32(1260))
						v184 = int32(10)
						v185 = base.I32_div_u_s(v177, v184)
						v187 = base.I32_rem_u_s(v185, int32(126))
						return v179<<(uint(int32(16))%32) + int32(2146828288) | (v187<<(uint(int32(8))%32) + (v175-v185*v184)&v176 + int32(_a_F_conv_utf8_to_5)) | int32(-2110783440)
					} else {
						if base.Ui32(v45-int32(_a_F_conv_utf8_to_10)) <= base.Ui32(int32(372)) {
							v207 = v45 - int32(894)
							v208 = int32(_a_F_conv_utf8_to_0)
							v209 = v207 & v208
							v211 = base.I32_div_u_s(v209, int32(1260))
							v216 = int32(10)
							v217 = base.I32_div_u_s(v209, v216)
							v219 = base.I32_rem_u_s(v217, int32(126))
							return v211<<(uint(int32(16))%32) + int32(2146828288) | (v219<<(uint(int32(8))%32) + (v207-v217*v216)&v208 + int32(_a_F_conv_utf8_to_5)) | int32(-2110783440)
						} else {
							if base.Ui32(v45-int32(_a_F_conv_utf8_to_11)) <= base.Ui32(int32(440)) {
								v239 = v45 - int32(900)
								v240 = int32(_a_F_conv_utf8_to_0)
								v241 = v239 & v240
								v243 = base.I32_div_u_s(v241, int32(1260))
								v248 = int32(10)
								v249 = base.I32_div_u_s(v241, v248)
								v251 = base.I32_rem_u_s(v249, int32(126))
								return v243<<(uint(int32(16))%32) + int32(2146828288) | (v251<<(uint(int32(8))%32) + (v239-v249*v248)&v240 + int32(_a_F_conv_utf8_to_5)) | int32(-2110783440)
							} else {
								if base.Ui32(v45-int32(_a_F_conv_utf8_to_12)) <= base.Ui32(int32(702)) {
									v271 = v45 - int32(911)
									v272 = int32(_a_F_conv_utf8_to_0)
									v273 = v271 & v272
									v275 = base.I32_div_u_s(v273, int32(1260))
									v280 = int32(10)
									v281 = base.I32_div_u_s(v273, v280)
									v283 = base.I32_rem_u_s(v281, int32(126))
									return v275<<(uint(int32(16))%32) + int32(2146828288) | (v283<<(uint(int32(8))%32) + (v271-v281*v280)&v272 + int32(_a_F_conv_utf8_to_5)) | int32(-2110783440)
								} else {
									if base.Ui32(v45-int32(_a_F_conv_utf8_to_13)) <= base.Ui32(int32(_a_F_conv_utf8_to_14)) {
										v303 = v45 - int32(_a_F_conv_utf8_to_15)
										v304 = int32(_a_F_conv_utf8_to_0)
										v305 = v303 & v304
										v307 = base.I32_div_u_s(v305, int32(_a_F_conv_utf8_to_16))
										v311 = base.I32_div_u_s(v305, int32(1260))
										v312 = int32(10)
										v313 = base.I32_rem_u_s(v311, v312)
										v320 = base.I32_div_u_s(v305, v312)
										v322 = base.I32_rem_u_s(v320, int32(126))
										return v307<<(uint(int32(24))%32) | v313<<(uint(int32(16))%32) - int32(2130706432) | (v322<<(uint(int32(8))%32) + (v303-v320*v312)&v304 + int32(_a_F_conv_utf8_to_5)) | int32(_a_F_conv_utf8_to_17)
									} else {
										if base.Ui32(v45-int32(_a_F_conv_utf8_to_18)) <= base.Ui32(int32(_a_F_conv_utf8_to_19)) {
											v342 = v45 - int32(_a_F_conv_utf8_to_20)
											v343 = int32(_a_F_conv_utf8_to_0)
											v344 = v342 & v343
											v346 = base.I32_div_u_s(v344, int32(_a_F_conv_utf8_to_16))
											v350 = base.I32_div_u_s(v344, int32(1260))
											v351 = int32(10)
											v352 = base.I32_rem_u_s(v350, v351)
											v359 = base.I32_div_u_s(v344, v351)
											v361 = base.I32_rem_u_s(v359, int32(126))
											return v346<<(uint(int32(24))%32) | v352<<(uint(int32(16))%32) - int32(2130706432) | (v361<<(uint(int32(8))%32) + (v342-v359*v351)&v343 + int32(_a_F_conv_utf8_to_5)) | int32(_a_F_conv_utf8_to_17)
										} else {
											if base.Ui32(v45-int32(_a_F_conv_utf8_to_21)) <= base.Ui32(int32(1029)) {
												v381 = v45 - int32(_a_F_conv_utf8_to_22)
												v382 = int32(_a_F_conv_utf8_to_0)
												v383 = v381 & v382
												v385 = base.I32_div_u_s(v383, int32(_a_F_conv_utf8_to_16))
												v389 = base.I32_div_u_s(v383, int32(1260))
												v390 = int32(10)
												v391 = base.I32_rem_u_s(v389, v390)
												v398 = base.I32_div_u_s(v383, v390)
												v400 = base.I32_rem_u_s(v398, int32(126))
												return v385<<(uint(int32(24))%32) | v391<<(uint(int32(16))%32) - int32(2130706432) | (v400<<(uint(int32(8))%32) + (v381-v398*v390)&v382 + int32(_a_F_conv_utf8_to_5)) | int32(_a_F_conv_utf8_to_17)
											} else {
												if base.Ui32(v45-int32(_a_F_conv_utf8_to_23)) <= base.Ui32(int32(25)) {
													v420 = v45 - int32(_a_F_conv_utf8_to_24)
													v421 = int32(_a_F_conv_utf8_to_0)
													v422 = v420 & v421
													v424 = base.I32_div_u_s(v422, int32(_a_F_conv_utf8_to_16))
													v428 = base.I32_div_u_s(v422, int32(1260))
													v429 = int32(10)
													v430 = base.I32_rem_u_s(v428, v429)
													v437 = base.I32_div_u_s(v422, v429)
													v439 = base.I32_rem_u_s(v437, int32(126))
													return v424<<(uint(int32(24))%32) | v430<<(uint(int32(16))%32) - int32(2130706432) | (v439<<(uint(int32(8))%32) + (v420-v437*v429)&v421 + int32(_a_F_conv_utf8_to_5)) | int32(_a_F_conv_utf8_to_17)
												} else {
													if base.Ui32(v45-int32(_a_F_conv_utf8_to_25)) <= base.Ui32(int32(_a_F_conv_utf8_to_26)) {
														v459 = v45 + int32(_a_F_conv_utf8_to_27)
														v460 = int32(10)
														v461 = base.I32_div_u_s(v459, v460)
														v463 = base.I32_rem_u_s(v461, int32(126))
														v473 = base.I32_div_u_s(v459, int32(_a_F_conv_utf8_to_16))
														v477 = base.I32_div_u_s(v459, int32(1260))
														v481 = base.I32_rem_u_s(v477&int32(_a_F_conv_utf8_to_0), v460)
														v493 = v463<<(uint(int32(8))%32) + (v459 - v461*v460) + int32(_a_F_conv_utf8_to_5) | (v473<<(uint(int32(24))%32) | v481<<(uint(int32(16))%32) - int32(2130706432)) | int32(_a_F_conv_utf8_to_17)
													} else {
														v493 = int32(0)
													}
													return v493
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
}
func F_convert_EXISTS_sublink_to_join(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
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
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v111 int32
	_ = v111
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v182 int32
	_ = v182
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v218 int32
	_ = v218
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v243 int32
	_ = v243
	var v250 int32
	_ = v250
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v270 int32
	_ = v270
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 int64
	_ = v284
	var v294 int32
	_ = v294
	v5 = int32(0)
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+48))
	if v11 != 0 {
		v294 = v5
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v294
L2:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v13 = F_copyObjectImpl(m, v10)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return int32(0)
L4:
	;
	v17 = F_simplify_EXISTS_query(m, l0, v13)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	if v17 == int32(0) {
		v294 = v5
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v13)+60))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+8)) = int32(0)
	v26 = F_contain_vars_of_level(m, v13, int32(1))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L3
	} else {
		goto L7
	}
L7:
	;
	if v26 != 0 {
		v294 = v5
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v29 = F_contain_vars_of_level(m, v22, int32(1))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L3
	} else {
		goto L9
	}
L9:
	;
	if v29 == int32(0) {
		v294 = v5
		goto L1
	} else {
		goto L10
	}
L10:
	;
	v33 = F_contain_volatile_functions(m, v22)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L3
	} else {
		goto L11
	}
L11:
	;
	if v33 != 0 {
		v294 = v5
		goto L1
	} else {
		goto L12
	}
L12:
	;
	F_replace_empty_jointree(m, v13)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L3
	} else {
		goto L13
	}
L13:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v12)+52))
	if v38 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
	v40 = v39
	goto L16
L15:
	;
	v40 = int32(0)
	goto L16
L16:
	;
	F_OffsetVarNodes(m, v13, v40)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L3
	} else {
		goto L17
	}
L17:
	;
	F_OffsetVarNodes(m, v22, v40)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L3
	} else {
		goto L18
	}
L18:
	;
	F_IncrementVarSublevelsUp(m, v13, int32(-1), int32(1))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L3
	} else {
		goto L19
	}
L19:
	;
	F_IncrementVarSublevelsUp(m, v22, int32(-1), int32(1))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L3
	} else {
		goto L20
	}
L20:
	;
	v53 = F_pull_varnos(m, l0, v22)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L3
	} else {
		goto L21
	}
L21:
	;
	if v53 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L22:
	;
	if int32(0) <= v111 {
		goto L33
	} else {
		goto L34
	}
L23:
	;
	v111 = base.I32_ctz(v97) | v98<<(uint(int32(5))%32)
	goto L22
L24:
	;
	v111 = int32(-2)
	goto L22
L25:
	;
	v64 = base.I32_div_s(int32(0), int32(32))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
	if v65 <= v64 {
		goto L24
	} else {
		goto L26
	}
L26:
	;
	v68 = v53 + int32(8)
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v68+v64<<(uint(int32(2))%32))))
	v75 = v72 & int32(-1)
	if v75 != 0 {
		v97 = v75
		v98 = v64
		goto L23
	} else {
		goto L27
	}
L27:
	;
	v77 = v64 + int32(1)
	if v77 == v65 {
		goto L24
	} else {
		goto L28
	}
L28:
	;
	v80 = v77
	goto L29
L29:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v68+v80<<(uint(int32(2))%32))))
	if v87 != 0 {
		v97 = v87
		v98 = v80
		goto L23
	} else {
		goto L31
	}
L30:
	;
	goto L24
L31:
	;
	v89 = v80 + int32(1)
	if v89 != v65 {
		v80 = v89
		goto L29
	} else {
		goto L32
	}
L32:
	;
	goto L30
L33:
	;
	v118 = v111
	v122 = v5
	goto L36
L34:
	;
	v193 = v5
	goto L35
L35:
	;
	F_bms_free(m, v53)
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L3
	} else {
		goto L54
	}
L36:
	;
	if v118 <= v40 {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	v193 = v126
	goto L35
L38:
	;
	v124 = F_bms_add_member(m, v122, v118)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L3
	} else {
		goto L41
	}
L39:
	;
	v126 = v122
	goto L40
L40:
	;
	if v53 == int32(0) {
		goto L44
	} else {
		goto L45
	}
L41:
	;
	v126 = v124
	goto L40
L42:
	;
	if int32(0) <= v182 {
		v118 = v182
		v122 = v126
		goto L36
	} else {
		goto L53
	}
L43:
	;
	v182 = base.I32_ctz(v168) | v169<<(uint(int32(5))%32)
	goto L42
L44:
	;
	v182 = int32(-2)
	goto L42
L45:
	;
	v133 = v118 + int32(1)
	v135 = base.I32_div_s(v133, int32(32))
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
	if v136 <= v135 {
		goto L44
	} else {
		goto L46
	}
L46:
	;
	v139 = v53 + int32(8)
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v139+v135<<(uint(int32(2))%32))))
	v146 = v143 & (int32(-1) << (uint(v133) % 32))
	if v146 != 0 {
		v168 = v146
		v169 = v135
		goto L43
	} else {
		goto L47
	}
L47:
	;
	v148 = v135 + int32(1)
	if v148 == v136 {
		goto L44
	} else {
		goto L48
	}
L48:
	;
	v151 = v148
	goto L49
L49:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v139+v151<<(uint(int32(2))%32))))
	if v158 != 0 {
		v168 = v158
		v169 = v151
		goto L43
	} else {
		goto L51
	}
L50:
	;
	goto L44
L51:
	;
	v160 = v151 + int32(1)
	if v160 != v136 {
		v151 = v160
		goto L49
	} else {
		goto L52
	}
L52:
	;
	goto L50
L53:
	;
	goto L37
L54:
	;
	v196 = int32(0)
	if v193 == v196 {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	if v250 == int32(0) {
		v294 = v196
		goto L1
	} else {
		goto L69
	}
L56:
	;
	v250 = int32(1)
	goto L55
L57:
	;
	goto L58
L58:
	;
	if l3 == int32(0) {
		v243 = v196
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v250 = v243
	goto L55
L60:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v193)+4))
	v207 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v207 < v206 {
		v243 = v196
		goto L59
	} else {
		goto L61
	}
L61:
	;
	v209 = int32(1)
	if v206 <= v209 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v212 = v209
	goto L64
L63:
	;
	v212 = v206
	goto L64
L64:
	;
	v213 = int32(8)
	v218 = int32(0)
	goto L65
L65:
	;
	v225 = v218 << (uint(int32(2)) % 32)
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v193+v213+v225)))
	v229 = *(*int32)(unsafe.Add(mBase, uint32(l3+v213+v225)))
	v232 = v227 & (v229 ^ int32(-1))
	v234 = base.B2i32(v232 == int32(0))
	if v232 != 0 {
		v243 = v234
		goto L59
	} else {
		goto L67
	}
L66:
	;
	v243 = v234
	goto L59
L67:
	;
	v236 = v218 + int32(1)
	if v236 != v212 {
		v218 = v236
		goto L65
	} else {
		goto L68
	}
L68:
	;
	goto L66
L69:
	;
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v13)+52))
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v13)+56))
	F_CombineRangeTables(m, v12+int32(52), v12+int32(56), v257, v258)
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L3
	} else {
		goto L70
	}
L70:
	;
	v262 = F_palloc0(m, int32(40))
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L3
	} else {
		goto L71
	}
L71:
	;
	v264 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v262)+12)) = v264
	*(*uint8)(unsafe.Add(mBase, uint32(v262)+8)) = uint8(v264)
	if l2 != 0 {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v270 = int32(5)
	goto L74
L73:
	;
	v270 = int32(4)
	goto L74
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v262)+4)) = v270
	*(*int32)(unsafe.Add(mBase, uint32(v262))) = int32(64)
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v13)+60))
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v274)+4))
	if v275 == int32(0) {
		v283 = v274
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v284 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v262)+32)) = v284
	*(*int32)(unsafe.Add(mBase, uint32(v262)+28)) = v22
	*(*int64)(unsafe.Add(mBase, uint32(v262)+20)) = v284
	*(*int32)(unsafe.Add(mBase, uint32(v262)+16)) = v283
	v294 = v262
	goto L1
L76:
	;
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v275)+4))
	if v278 != int32(1) {
		v283 = v274
		goto L75
	} else {
		goto L77
	}
L77:
	;
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v275)+12))
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v281)))
	v283 = v282
	goto L75
}
func F_convert_combining_aggrefs(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	v3 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	if l0 == v3 {
		v79 = v3
		m.G0 = v7 + int32(16)
		return v79
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if v11 == int32(9) {
			v15 = F_palloc0(m, int32(72))
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v15))) = int32(9)
				base.MemoryCopy(m, v15, l0, int32(72))
				v23 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v15)+44)) = v23
				*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v23
				v27 = F_copyObjectImpl(m, v15)
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return int32(0)
				} else {
					v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
					*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v29
					v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
					*(*int32)(unsafe.Add(mBase, uint32(v15)+44)) = v31
					*(*int32)(unsafe.Add(mBase, uint32(v15)+56)) = int32(6)
					v38 = *(*int32)(unsafe.Add(mBase, uint32(v15)+20))
					if v38 == int32(2281) {
						v41 = int32(17)
					} else {
						v41 = v38
					}
					*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = v41
					v48 = int32(0)
					v50 = F_makeTargetEntry(m, v15, int32(1), v48, v48)
					mBase = m.M
					v51 = m.ExcPending
					if v51 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v50
						*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = v50
						v57 = F_list_make1_impl(m, int32(1), v7+int32(8))
						mBase = m.M
						v58 = m.ExcPending
						if v58 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v27)+32)) = v57
							*(*int32)(unsafe.Add(mBase, uint32(v27)+56)) = int32(9)
							v79 = v27
							m.G0 = v7 + int32(16)
							return v79
						}
					}
				}
			}
		} else {
			v75 = F_expression_tree_mutator_impl(m, l0, int32(837), l1)
			mBase = m.M
			v76 = m.ExcPending
			if v76 != 0 {
				return int32(0)
			} else {
				v79 = v75
				m.G0 = v7 + int32(16)
				return v79
			}
		}
	}
}
func F_copy_dest_startup(m *base.Module, l0 int32, l1 int32, l2 int32) {
	return
}
func F_copydir(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
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
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v188 int32
	_ = v188
	var v193 int32
	_ = v193
	v9 = m.G0
	v11 = v9 - int32(_a_F_copydir_0)
	m.G0 = v11
	v14 = *(*int32)(unsafe.Add(mBase, _c_F_copydir[0]))
	v15 = F_mkdir(m, l1, v14)
	mBase = m.M
	goto L1
L1:
	;
	if v15 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L2:
	;
	v18 = F_AllocateDir(m, l0)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L5
	} else {
		goto L58
	}
L5:
	;
	return
L6:
	;
	v20 = F_ReadDir(m, v18, l0)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	if v20 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v26 = v20
	goto L11
L9:
	;
	goto L10
L10:
	;
	F_FreeDir(m, v18)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L5
	} else {
		goto L33
	}
L11:
	;
	v31 = *(*int32)(unsafe.Add(mBase, _c_F_copydir[1]))
	if v31 != 0 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	goto L10
L13:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L5
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+19)))
	if v34 != int32(46) {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	goto L15
L17:
	;
	v91 = F_ReadDir(m, v18, l0)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L5
	} else {
		goto L31
	}
L18:
	;
	v47 = v26 + int32(19)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+36)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = l0
	v51 = v11 + int32(2112)
	v56 = F_pg_snprintf(m, v51, int32(2048), int32(_a_F_copydir_1), v11+int32(32))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L5
	} else {
		goto L23
	}
L19:
	;
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+20)))
	if v37 == int32(0) {
		goto L17
	} else {
		goto L20
	}
L20:
	;
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+20)))
	if v40 != int32(46) {
		goto L18
	} else {
		goto L21
	}
L21:
	;
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+21)))
	if v43 == int32(0) {
		goto L17
	} else {
		goto L22
	}
L22:
	;
	goto L18
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = l1
	v66 = F_pg_snprintf(m, v11-int32(-64), int32(2048), int32(_a_F_copydir_1), v11+int32(16))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L5
	} else {
		goto L24
	}
L24:
	;
	v70 = F_get_dirent_type(m, v51, v26, int32(0), int32(21))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L5
	} else {
		goto L27
	}
L25:
	;
	F_copy_file(m, v11+int32(2112), v11-int32(-64))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L5
	} else {
		goto L30
	}
L26:
	;
	if l2 == int32(0) {
		goto L17
	} else {
		goto L28
	}
L27:
	;
	switch v70 - int32(2) {
	case 0:
		goto L25
	case 1:
		goto L26
	default:
		goto L17
	}
L28:
	;
	F_copydir(m, v11+int32(2112), v11-int32(-64), int32(1))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L5
	} else {
		goto L29
	}
L29:
	;
	goto L17
L30:
	;
	goto L17
L31:
	;
	if v91 != 0 {
		v26 = v91
		goto L11
	} else {
		goto L32
	}
L32:
	;
	goto L12
L33:
	;
	v104 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_copydir[2])))
	if v104 == int32(1) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v107 = F_AllocateDir(m, l1)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L5
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	m.G0 = v11 + int32(_a_F_copydir_0)
	return
L37:
	;
	v109 = F_ReadDir(m, v107, l1)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L5
	} else {
		goto L38
	}
L38:
	;
	if v109 != 0 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v115 = v109
	goto L42
L40:
	;
	goto L41
L41:
	;
	F_FreeDir(m, v107)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L5
	} else {
		goto L56
	}
L42:
	;
	v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115)+19)))
	if v119 != int32(46) {
		goto L45
	} else {
		goto L46
	}
L43:
	;
	goto L41
L44:
	;
	v151 = F_ReadDir(m, v107, l1)
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L5
	} else {
		goto L54
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v115 + int32(19)
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = l1
	v136 = v11 - int32(-64)
	v139 = F_pg_snprintf(m, v136, int32(2048), int32(_a_F_copydir_1), v11)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L5
	} else {
		goto L50
	}
L46:
	;
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115)+20)))
	if v122 == int32(0) {
		goto L44
	} else {
		goto L47
	}
L47:
	;
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115)+20)))
	if v125 != int32(46) {
		goto L45
	} else {
		goto L48
	}
L48:
	;
	v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115)+21)))
	if v128 == int32(0) {
		goto L44
	} else {
		goto L49
	}
L49:
	;
	goto L45
L50:
	;
	v143 = F_get_dirent_type(m, v136, v115, int32(0), int32(21))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L5
	} else {
		goto L51
	}
L51:
	;
	if v143 != int32(2) {
		goto L44
	} else {
		goto L52
	}
L52:
	;
	F_fsync_fname(m, v136, int32(0))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L5
	} else {
		goto L53
	}
L53:
	;
	goto L44
L54:
	;
	if v151 != 0 {
		v115 = v151
		goto L42
	} else {
		goto L55
	}
L55:
	;
	goto L43
L56:
	;
	F_fsync_fname(m, l1, int32(1))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L5
	} else {
		goto L57
	}
L57:
	;
	goto L36
L58:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L5
	} else {
		goto L59
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = l1
	F_errmsg(m, int32(_a_F_copydir_2), v11+int32(48))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L5
	} else {
		goto L60
	}
L60:
	;
	F_errfinish(m, int32(_a_F_copydir_3), int32(58), int32(_a_F_copydir_4))
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L5
	} else {
		goto L61
	}
L61:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_core_yy_create_buffer(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
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
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	v8 = F_palloc(m, int32(48))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		if v8 != 0 {
			*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = int32(_a_F_core_yy_create_buffer_0)
			v15 = F_palloc(m, int32(_a_F_core_yy_create_buffer_1))
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v15
				if v15 == int32(0) {
					F_yy_fatal_error_2(m, int32(_a_F_core_yy_create_buffer_2))
					mBase = m.M
					v81 = m.ExcPending
					if v81 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				} else {
					v20 = int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = v20
					v23 = *(*int32)(unsafe.Add(mBase, _c_F_core_yy_create_buffer[0]))
					v24 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v24
					*(*uint8)(unsafe.Add(mBase, uint32(v15))) = uint8(v24)
					v28 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
					*(*uint8)(unsafe.Add(mBase, uint32(v28)+1)) = uint8(v24)
					*(*int32)(unsafe.Add(mBase, uint32(v8)+44)) = v24
					*(*int32)(unsafe.Add(mBase, uint32(v8)+28)) = v20
					v35 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
					*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v35
					v37 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
					if v37 == v24 {
					} else {
						v40 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
						v43 = v37 + v40<<(uint(int32(2))%32)
						v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
						if v8 != v44 {
						} else {
							v46 = *(*int32)(unsafe.Add(mBase, uint32(v44)+16))
							*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v46
							v48 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
							v49 = *(*int32)(unsafe.Add(mBase, uint32(v48)+8))
							*(*int32)(unsafe.Add(mBase, uint32(l1)+80)) = v49
							*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = v49
							v52 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
							v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
							*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v53
							v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49))))
							*(*uint8)(unsafe.Add(mBase, uint32(l1)+24)) = uint8(v55)
						}
					}
					*(*int32)(unsafe.Add(mBase, uint32(v8)+40)) = int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
					v62 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
					if v62 != 0 {
						v63 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
						v67 = *(*int32)(unsafe.Add(mBase, uint32(v62+v63<<(uint(int32(2))%32))))
						if v8 == v67 {
						} else {
							*(*int64)(unsafe.Add(mBase, uint32(v8)+32)) = int64(1)
						}
					} else {
						*(*int64)(unsafe.Add(mBase, uint32(v8)+32)) = int64(1)
					}
					*(*int32)(unsafe.Add(mBase, uint32(v8)+24)) = int32(0)
					*(*int32)(unsafe.Add(mBase, _c_F_core_yy_create_buffer[0])) = v23
					return v8
				}
			}
		} else {
			F_yy_fatal_error_2(m, int32(_a_F_core_yy_create_buffer_2))
			mBase = m.M
			v78 = m.ExcPending
			if v78 != 0 {
				return int32(0)
			} else {
				base.Wasm_trap_unreachable()
				for {
				}
			}
		}
	}
}
func F_cost_sort(m *base.Module, l0 int32, l1 int32, l2 float64, l3 float64, l4 int32, l5 float64, l6 int32, l7 float64) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v20 float64
	_ = v20
	var v23 float64
	_ = v23
	var v25 float64
	_ = v25
	var v28 float64
	_ = v28
	var v31 int64
	_ = v31
	var v32 float64
	_ = v32
	var v39 float64
	_ = v39
	var v41 float64
	_ = v41
	var v45 int32
	_ = v45
	var v46 float64
	_ = v46
	var v48 float64
	_ = v48
	var v49 int32
	_ = v49
	var v52 float64
	_ = v52
	var v56 float64
	_ = v56
	var v58 float64
	_ = v58
	var v59 float64
	_ = v59
	var v61 float64
	_ = v61
	var v62 float64
	_ = v62
	var v66 float64
	_ = v66
	var v69 float64
	_ = v69
	var v73 float64
	_ = v73
	var v79 float64
	_ = v79
	var v80 float64
	_ = v80
	var v84 float64
	_ = v84
	var v88 float64
	_ = v88
	var v96 float64
	_ = v96
	var v99 float64
	_ = v99
	var v102 float64
	_ = v102
	var v105 int32
	_ = v105
	var v106 float64
	_ = v106
	var v112 float64
	_ = v112
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v15 = v12 + int32(8)
	v20 = float64(2)
	if base.F64_lt(l3, v20) != 0 {
		v23 = v20
	} else {
		v23 = l3
	}
	v25 = *(*float64)(unsafe.Add(mBase, _c_F_cost_sort[0]))
	v28 = base.F64_mul(v23, base.F64_add(base.F64_add(v25, v25), l5))
	v31 = base.I64_extend_i32_s(l6) << (uint(int64(10)) % 64)
	v32 = base.F64_convert_i64_s(v31)
	v39 = base.F64_convert_i32_u((l4+int32(7))&int32(-8) + int32(24))
	v41 = base.F64_mul(l3, v39)
	v45 = base.F64_lt(l7, v23) & base.F64_gt(l7, float64(0))
	if v45 != 0 {
		v46 = base.F64_mul(l7, v39)
	} else {
		v46 = v41
	}
	if base.F64_lt(v32, v46) != 0 {
		v48 = F_log(m, v23)
		mBase = m.M
		v49 = F_tuplesort_merge_order(m, v31)
		mBase = m.M
		v52 = base.F64_mul(base.F64_div(v48, float64(0.693147180559945)), v28)
		*(*float64)(unsafe.Add(mBase, uint32(v15))) = v52
		v56 = base.F64_ceil(base.F64_mul(v41, float64(0.0001220703125)))
		v58 = base.F64_div(v41, v32)
		v59 = base.F64_convert_i32_s(v49)
		if base.F64_gt(v58, v59) != 0 {
			v61 = F_log(m, v58)
			mBase = m.M
			v62 = F_log(m, v59)
			mBase = m.M
			v66 = base.F64_ceil(base.F64_div(v61, v62))
		} else {
			v66 = float64(1)
		}
		v69 = *(*float64)(unsafe.Add(mBase, _c_F_cost_sort[1]))
		v73 = *(*float64)(unsafe.Add(mBase, _c_F_cost_sort[2]))
		v96 = base.F64_add(base.F64_mul(base.F64_mul(base.F64_add(v56, v56), v66), base.F64_add(base.F64_mul(v69, float64(0.75)), base.F64_mul(v73, float64(0.25)))), v52)
	} else {
		if v45 != 0 {
			v79 = l7
		} else {
			v79 = v23
		}
		v80 = base.F64_add(v79, v79)
		if base.F64_gt(v23, v80)|base.F64_gt(v41, v32) != 0 {
			v84 = F_log(m, v80)
			mBase = m.M
			v96 = base.F64_mul(base.F64_div(v84, float64(0.693147180559945)), v28)
		} else {
			v88 = F_log(m, v23)
			mBase = m.M
			v96 = base.F64_mul(base.F64_div(v88, float64(0.693147180559945)), v28)
		}
	}
	*(*float64)(unsafe.Add(mBase, uint32(v15))) = v96
	v99 = *(*float64)(unsafe.Add(mBase, _c_F_cost_sort[0]))
	*(*float64)(unsafe.Add(mBase, uint32(v12))) = base.F64_mul(v23, v99)
	v102 = *(*float64)(unsafe.Add(mBase, uint32(v12)+8))
	*(*float64)(unsafe.Add(mBase, uint32(l0)+32)) = l3
	v105 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_cost_sort[3])))
	v106 = base.F64_add(l2, v102)
	*(*float64)(unsafe.Add(mBase, uint32(l0)+48)) = v106
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = l1 + (v105 ^ int32(1))
	v112 = *(*float64)(unsafe.Add(mBase, uint32(v12)))
	*(*float64)(unsafe.Add(mBase, uint32(l0)+56)) = base.F64_add(v106, v112)
	m.G0 = v12 + int32(16)
	return
}
func F_crc32c_bytea(m *base.Module, l0 int32) int32 {
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
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = F_pg_detoast_datum_packed(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		v11 = int32(1)
		v12 = v7 + v11
		v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
		v17 = v15 & v11
		if v17 != 0 {
			v18 = v12
		} else {
			v18 = v7 + int32(4)
		}
		if v15 == int32(1) {
			v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
			if v24 == int32(18) {
				v27 = int32(16)
			} else {
				v27 = int32(0)
			}
			if base.Ui32((v24-int32(1))&int32(255)) < base.Ui32(int32(3)) {
				v34 = int32(4)
			} else {
				v34 = v27
			}
			v45 = v34
		} else {
			v35 = int32(1)
			if v17 != 0 {
				v45 = int32(base.Ui32(v15)>>(uint(v35)%32)) - v35
			} else {
				v39 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
				v45 = int32(base.Ui32(v39)>>(uint(int32(2))%32)) - int32(4)
			}
		}
		v46 = m.Env.Pgmem_crc32c(m, int32(-1), v18, v45)
		mBase = m.M
		v50 = F_Int64GetDatum(m, base.I64_extend_i32_u(v46^int32(-1)))
		mBase = m.M
		v51 = m.ExcPending
		if v51 != 0 {
			return int32(0)
		} else {
			return v50
		}
	}
}
func F_create_incremental_sort_path(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 float64) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v37 float64
	_ = v37
	var v38 float64
	_ = v38
	var v39 float64
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	v10 = F_palloc0(m, int32(88))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = l1
		*(*int64)(unsafe.Add(mBase, uint32(v10))) = int64(1559073128752)
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
		v18 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v10)+20)) = uint8(v18)
		*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v18
		*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = v17
		v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+26)))
		if v23 == int32(1) {
			v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+21)))
			v28 = v26
		} else {
			v28 = int32(0)
		}
		v30 = v28 & int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(v10)+21)) = uint8(v30)
		v32 = *(*int32)(unsafe.Add(mBase, uint32(l2)+24))
		*(*int32)(unsafe.Add(mBase, uint32(v10)+72)) = l2
		*(*int32)(unsafe.Add(mBase, uint32(v10)+64)) = l3
		*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = v32
		v36 = *(*int32)(unsafe.Add(mBase, uint32(l2)+40))
		v37 = *(*float64)(unsafe.Add(mBase, uint32(l2)+48))
		v38 = *(*float64)(unsafe.Add(mBase, uint32(l2)+56))
		v39 = *(*float64)(unsafe.Add(mBase, uint32(l2)+32))
		v40 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
		v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+32))
		v43 = *(*int32)(unsafe.Add(mBase, _c_F_create_incremental_sort_path[0]))
		F_cost_incremental_sort(m, v10, l0, l3, l4, v36, v37, v38, v39, v41, v43, l5)
		mBase = m.M
		v45 = m.ExcPending
		if v45 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v10)+80)) = l4
			return v10
		}
	}
}
func F_create_indexscan_plan(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v74 int32
	_ = v74
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
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
	var v114 int32
	_ = v114
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v141 int32
	_ = v141
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v265 int32
	_ = v265
	var v270 int32
	_ = v270
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v292 int32
	_ = v292
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v319 int32
	_ = v319
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v373 int32
	_ = v373
	var v377 int32
	_ = v377
	var v379 int32
	_ = v379
	var v383 int32
	_ = v383
	var v386 int32
	_ = v386
	var v389 int32
	_ = v389
	var v392 int32
	_ = v392
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v433 int32
	_ = v433
	var v438 int32
	_ = v438
	var v446 int32
	_ = v446
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v469 int32
	_ = v469
	var v484 int32
	_ = v484
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v519 int32
	_ = v519
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v559 int32
	_ = v559
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v582 int32
	_ = v582
	var v597 int32
	_ = v597
	var v599 float64
	_ = v599
	var v601 float64
	_ = v601
	var v603 float64
	_ = v603
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v608 int32
	_ = v608
	var v610 int32
	_ = v610
	v6 = int32(0)
	v21 = m.G0
	v23 = v21 - int32(32)
	m.G0 = v23
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+68))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	if v30 == v6 {
		v141 = v29
		v148 = v6
		v150 = v6
		v154 = v25
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(l1)+84))
	v156 = int32(0)
	v163 = v156
	v164 = v156
	goto L18
L2:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	if v33 <= int32(0) {
		v141 = v29
		v148 = v6
		v150 = v6
		v154 = v25
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v41 = v33
	v45 = v6
	v50 = v6
	v52 = v6
	goto L4
L4:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v30)+12))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v56+v45<<(uint(int32(2))%32))))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)+8))
	if v61 == int32(0) {
		v114 = v41
		v123 = v50
		v125 = v52
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
	v133 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v141 = v132
	v148 = v123
	v150 = v125
	v154 = v133
	goto L1
L6:
	;
	v130 = v45 + int32(1)
	if v130 < v114 {
		v41 = v114
		v45 = v130
		v50 = v123
		v52 = v125
		goto L4
	} else {
		goto L16
	}
L7:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v61)+4))
	if v64 <= int32(0) {
		v114 = v41
		v123 = v50
		v125 = v52
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v67 = int32(*(*int16)(unsafe.Add(mBase, uint32(v60)+14)))
	v74 = int32(0)
	v83 = v50
	v85 = v52
	goto L9
L9:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v61)+12))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v89+v74<<(uint(int32(2))%32))))
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v93)+4))
	v95 = F_lappend(m, v83, v94)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	v114 = v108
	v123 = v95
	v125 = v102
	goto L6
L11:
	;
	return int32(0)
L12:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v60)+16))
	v100 = F_fix_indexqual_clause(m, l0, v25, v67, v94, v99)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v102 = F_lappend(m, v85, v100)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L11
	} else {
		goto L14
	}
L14:
	;
	v105 = v74 + int32(1)
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v61)+4))
	if v105 < v106 {
		v74 = v105
		v83 = v95
		v85 = v102
		goto L9
	} else {
		goto L15
	}
L15:
	;
	goto L10
L16:
	;
	goto L5
L17:
	;
	if l4 != 0 {
		goto L94
	} else {
		goto L95
	}
L18:
	;
	v178 = int32(0)
	if v141 == v178 {
		v189 = v178
		goto L20
	} else {
		goto L21
	}
L19:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v424 = m.ExcPending
	if v424 != 0 {
		goto L11
	} else {
		goto L90
	}
L20:
	;
	if v155 != 0 {
		goto L25
	} else {
		goto L26
	}
L21:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v141)+4))
	if v183 <= v163 {
		v189 = int32(0)
		goto L20
	} else {
		goto L22
	}
L22:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v141)+12))
	v189 = v185 + v163<<(uint(int32(2))%32)
	goto L20
L23:
	;
	goto L19
L24:
	;
	v412 = *(*int32)(unsafe.Add(mBase, uint32(v197+v163<<(uint(int32(2))%32))))
	v413 = *(*int32)(unsafe.Add(mBase, uint32(v189)))
	v415 = F_fix_indexqual_clause(m, l0, v154, v412, v413, int32(0))
	mBase = m.M
	v416 = m.ExcPending
	if v416 != 0 {
		goto L11
	} else {
		goto L88
	}
L25:
	;
	v190 = int32(0)
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v155)+4))
	if base.B2i32(v189 == v190)|base.B2i32(v192 <= v163) == v190 {
		goto L28
	} else {
		goto L29
	}
L26:
	;
	v199 = v178
	goto L27
L27:
	;
	v200 = int32(0)
	if l3 == v200 {
		v319 = v200
		goto L32
	} else {
		goto L33
	}
L28:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v155)+12))
	if v197 != 0 {
		goto L24
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	v199 = v164
	goto L27
L31:
	;
	goto L30
L32:
	;
	v332 = F_order_qual_clauses(m, l0, v319)
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L11
	} else {
		goto L66
	}
L33:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v203 <= int32(0) {
		v319 = v200
		goto L32
	} else {
		goto L34
	}
L34:
	;
	v212 = int32(0)
	v214 = v200
	goto L35
L35:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v227+v212<<(uint(int32(2))%32))))
	v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v231)+10)))
	if v232 != 0 {
		v306 = v214
		goto L37
	} else {
		goto L38
	}
L36:
	;
	v319 = v306
	goto L32
L37:
	;
	v309 = v212 + int32(1)
	v310 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v309 < v310 {
		v212 = v309
		v214 = v306
		goto L35
	} else {
		goto L65
	}
L38:
	;
	if v30 != 0 {
		goto L41
	} else {
		goto L42
	}
L39:
	;
	if v285 != 0 {
		v306 = v214
		goto L37
	} else {
		goto L56
	}
L40:
	;
	goto L39
L41:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	if v238 <= int32(0) {
		v285 = int32(0)
		goto L40
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	v285 = int32(0)
	goto L40
L44:
	;
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v231)+60))
	v242 = int32(0)
	if v242 < v238 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v245 = v238
	goto L47
L46:
	;
	v245 = v242
	goto L47
L47:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v30)+12))
	v249 = int32(0)
	goto L48
L48:
	;
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v246+v249<<(uint(int32(2))%32))))
	v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v258)+12)))
	if v259 != 0 {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	goto L43
L50:
	;
	v270 = v249 + int32(1)
	if v270 != v245 {
		v249 = v270
		goto L48
	} else {
		goto L55
	}
L51:
	;
	v260 = int32(1)
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v258)+4))
	if v231 == v261 {
		v285 = v260
		goto L40
	} else {
		goto L52
	}
L52:
	;
	if v241 == int32(0) {
		goto L50
	} else {
		goto L53
	}
L53:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v261)+60))
	if v265 == v241 {
		v285 = v260
		goto L40
	} else {
		goto L54
	}
L54:
	;
	goto L50
L55:
	;
	goto L49
L56:
	;
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v231)+4))
	v288 = F_contain_mutable_functions(m, v287)
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L11
	} else {
		goto L57
	}
L57:
	;
	if v288 == int32(0) {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v231)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+24)) = v292
	*(*int32)(unsafe.Add(mBase, uint32(v23)+28)) = v292
	v298 = F_list_make1_impl(m, int32(1), v23+int32(24))
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L11
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	v304 = F_lappend(m, v214, v231)
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L11
	} else {
		goto L64
	}
L61:
	;
	v301 = F_predicate_implied_by(m, v298, v148, int32(0))
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L11
	} else {
		goto L62
	}
L62:
	;
	if v301 != 0 {
		v306 = v214
		goto L37
	} else {
		goto L63
	}
L63:
	;
	goto L60
L64:
	;
	v306 = v304
	goto L37
L65:
	;
	goto L36
L66:
	;
	v335 = F_extract_actual_clauses(m, v332, int32(0))
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L11
	} else {
		goto L67
	}
L67:
	;
	v337 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v337 != 0 {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v338 = F_replace_nestloop_params_mutator(m, v148, l0)
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L11
	} else {
		goto L71
	}
L69:
	;
	v344 = v335
	v345 = v29
	v346 = v148
	goto L70
L70:
	;
	if v345 == int32(0) {
		goto L74
	} else {
		goto L75
	}
L71:
	;
	v340 = F_replace_nestloop_params_mutator(m, v335, l0)
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		goto L11
	} else {
		goto L72
	}
L72:
	;
	v342 = F_replace_nestloop_params_mutator(m, v29, l0)
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L11
	} else {
		goto L73
	}
L73:
	;
	v344 = v340
	v345 = v342
	v346 = v338
	goto L70
L74:
	;
	v446 = int32(0)
	goto L17
L75:
	;
	goto L76
L76:
	;
	v350 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
	v351 = int32(0)
	v358 = v351
	v360 = v351
	goto L77
L77:
	;
	v373 = int32(0)
	if v350 == v373 {
		v383 = v373
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v386 = *(*int32)(unsafe.Add(mBase, uint32(v345)+4))
	if base.B2i32(v383 == int32(0))|base.B2i32(v386 <= v358) != 0 {
		v446 = v360
		goto L17
	} else {
		goto L82
	}
L80:
	;
	v377 = *(*int32)(unsafe.Add(mBase, uint32(v350)+4))
	if v377 <= v358 {
		v383 = int32(0)
		goto L79
	} else {
		goto L81
	}
L81:
	;
	v379 = *(*int32)(unsafe.Add(mBase, uint32(v350)+12))
	v383 = v379 + v358<<(uint(int32(2))%32)
	goto L79
L82:
	;
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v345)+12))
	if v389 == int32(0) {
		v446 = v360
		goto L17
	} else {
		goto L83
	}
L83:
	;
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v383)))
	v396 = *(*int32)(unsafe.Add(mBase, uint32(v389+v358<<(uint(int32(2))%32))))
	v397 = F_exprType(m, v396)
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L11
	} else {
		goto L84
	}
L84:
	;
	v399 = *(*int32)(unsafe.Add(mBase, uint32(v392)+8))
	v400 = *(*int32)(unsafe.Add(mBase, uint32(v392)+12))
	v401 = F_get_opfamily_member_for_cmptype(m, v399, v397, v397, v400)
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		goto L11
	} else {
		goto L85
	}
L85:
	;
	if v401 == int32(0) {
		goto L23
	} else {
		goto L86
	}
L86:
	;
	v407 = F_lappend_oid(m, v360, v401)
	mBase = m.M
	v408 = m.ExcPending
	if v408 != 0 {
		goto L11
	} else {
		goto L87
	}
L87:
	;
	v358 = v358 + int32(1)
	v360 = v407
	goto L77
L88:
	;
	v417 = F_lappend(m, v164, v415)
	mBase = m.M
	v418 = m.ExcPending
	if v418 != 0 {
		goto L11
	} else {
		goto L89
	}
L89:
	;
	v163 = v163 + int32(1)
	v164 = v417
	goto L18
L90:
	;
	v425 = *(*int32)(unsafe.Add(mBase, uint32(v392)+12))
	v426 = *(*int32)(unsafe.Add(mBase, uint32(v392)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+12)) = v426
	*(*int32)(unsafe.Add(mBase, uint32(v23)+8)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v23)+4)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v23))) = v425
	F_errmsg_internal(m, int32(_a_F_create_indexscan_plan_0), v23)
	mBase = m.M
	v433 = m.ExcPending
	if v433 != 0 {
		goto L11
	} else {
		goto L91
	}
L91:
	;
	F_errfinish(m, int32(_a_F_create_indexscan_plan_1), int32(3138), int32(_a_F_create_indexscan_plan_2))
	mBase = m.M
	v438 = m.ExcPending
	if v438 != 0 {
		goto L11
	} else {
		goto L92
	}
L92:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L93:
	;
	v597 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v582)+4)) = v597
	v599 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v582)+8)) = v599
	v601 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v582)+16)) = v601
	v603 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v582)+24)) = v603
	v605 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v606 = *(*int32)(unsafe.Add(mBase, uint32(v605)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v582)+32)) = v606
	v608 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)))
	*(*uint8)(unsafe.Add(mBase, uint32(v582)+36)) = uint8(v608)
	v610 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
	*(*uint8)(unsafe.Add(mBase, uint32(v582)+37)) = uint8(v610)
	m.G0 = v23 + int32(32)
	return v582
L94:
	;
	v459 = *(*int32)(unsafe.Add(mBase, uint32(v25)+92))
	if v459 != 0 {
		goto L97
	} else {
		goto L98
	}
L95:
	;
	goto L96
L96:
	;
	v559 = *(*int32)(unsafe.Add(mBase, uint32(l1)+88))
	v561 = F_palloc0(m, int32(112))
	mBase = m.M
	v562 = m.ExcPending
	if v562 != 0 {
		goto L11
	} else {
		goto L107
	}
L97:
	;
	v460 = *(*int32)(unsafe.Add(mBase, uint32(v459)+4))
	if int32(0) < v460 {
		goto L100
	} else {
		goto L101
	}
L98:
	;
	v541 = int32(0)
	goto L99
L99:
	;
	v542 = *(*int32)(unsafe.Add(mBase, uint32(l1)+88))
	v544 = F_palloc0(m, int32(104))
	mBase = m.M
	v545 = m.ExcPending
	if v545 != 0 {
		goto L11
	} else {
		goto L106
	}
L100:
	;
	v469 = int32(0)
	goto L103
L101:
	;
	goto L102
L102:
	;
	v519 = *(*int32)(unsafe.Add(mBase, uint32(v25)+92))
	v541 = v519
	goto L99
L103:
	;
	v484 = *(*int32)(unsafe.Add(mBase, uint32(v459)+12))
	v488 = *(*int32)(unsafe.Add(mBase, uint32(v484+v469<<(uint(int32(2))%32))))
	v489 = *(*int32)(unsafe.Add(mBase, uint32(v25)+76))
	v491 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v489+v469))))
	v492 = int32(1)
	v493 = v491 ^ v492
	*(*uint8)(unsafe.Add(mBase, uint32(v488)+26)) = uint8(v493)
	v496 = v469 + v492
	v497 = *(*int32)(unsafe.Add(mBase, uint32(v459)+4))
	if v496 < v497 {
		v469 = v496
		goto L103
	} else {
		goto L105
	}
L104:
	;
	goto L102
L105:
	;
	goto L104
L106:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v544)+100)) = v542
	*(*int32)(unsafe.Add(mBase, uint32(v544)+96)) = v541
	*(*int32)(unsafe.Add(mBase, uint32(v544)+92)) = v199
	*(*int32)(unsafe.Add(mBase, uint32(v544)+88)) = v346
	*(*int32)(unsafe.Add(mBase, uint32(v544)+84)) = v150
	*(*int32)(unsafe.Add(mBase, uint32(v544)+80)) = v26
	*(*int32)(unsafe.Add(mBase, uint32(v544)+72)) = v28
	*(*int64)(unsafe.Add(mBase, uint32(v544)+52)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v544)+48)) = v344
	*(*int32)(unsafe.Add(mBase, uint32(v544)+44)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v544))) = int32(342)
	v582 = v544
	goto L93
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v561)+104)) = v559
	*(*int32)(unsafe.Add(mBase, uint32(v561)+100)) = v446
	*(*int32)(unsafe.Add(mBase, uint32(v561)+96)) = v345
	*(*int32)(unsafe.Add(mBase, uint32(v561)+92)) = v199
	*(*int32)(unsafe.Add(mBase, uint32(v561)+88)) = v346
	*(*int32)(unsafe.Add(mBase, uint32(v561)+84)) = v150
	*(*int32)(unsafe.Add(mBase, uint32(v561)+80)) = v26
	*(*int32)(unsafe.Add(mBase, uint32(v561)+72)) = v28
	*(*int64)(unsafe.Add(mBase, uint32(v561)+52)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v561)+48)) = v344
	*(*int32)(unsafe.Add(mBase, uint32(v561)+44)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v561))) = int32(341)
	v582 = v561
	goto L93
}
func F_create_limit_path(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int64, l6 int64) int32 {
	mBase := m.M
	_ = mBase
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v40 float64
	_ = v40
	var v42 int32
	_ = v42
	var v44 float64
	_ = v44
	var v46 float64
	_ = v46
	var v48 int32
	_ = v48
	var v55 int32
	_ = v55
	var v56 float64
	_ = v56
	var v58 int32
	_ = v58
	var v59 float64
	_ = v59
	var v61 int32
	_ = v61
	var v62 float64
	_ = v62
	var v69 float64
	_ = v69
	var v70 float64
	_ = v70
	var v79 float64
	_ = v79
	var v83 float64
	_ = v83
	var v84 float64
	_ = v84
	var v85 float64
	_ = v85
	var v86 float64
	_ = v86
	var v88 float64
	_ = v88
	var v95 float64
	_ = v95
	var v98 float64
	_ = v98
	var v99 float64
	_ = v99
	var v100 float64
	_ = v100
	var v103 float64
	_ = v103
	var v106 float64
	_ = v106
	var v113 float64
	_ = v113
	var v114 float64
	_ = v114
	var v123 float64
	_ = v123
	var v127 float64
	_ = v127
	var v128 float64
	_ = v128
	var v129 float64
	_ = v129
	var v130 float64
	_ = v130
	var v132 float64
	_ = v132
	var v138 float64
	_ = v138
	var v141 float64
	_ = v141
	var v144 float64
	_ = v144
	v16 = F_palloc0(m, int32(88))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = l0
		*(*int64)(unsafe.Add(mBase, uint32(v16))) = int64(1602022801725)
		v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
		v24 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v16)+20)) = uint8(v24)
		*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v24
		*(*int32)(unsafe.Add(mBase, uint32(v16)+12)) = v23
		v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+26)))
		if v29 == int32(1) {
			v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
			v34 = v32
		} else {
			v34 = int32(0)
		}
		v36 = v34 & int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(v16)+21)) = uint8(v36)
		v38 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
		*(*int32)(unsafe.Add(mBase, uint32(v16)+24)) = v38
		v40 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
		*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v40
		v42 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
		*(*int32)(unsafe.Add(mBase, uint32(v16)+40)) = v42
		v44 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
		*(*float64)(unsafe.Add(mBase, uint32(v16)+48)) = v44
		v46 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
		*(*float64)(unsafe.Add(mBase, uint32(v16)+56)) = v46
		v48 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
		*(*int32)(unsafe.Add(mBase, uint32(v16)+84)) = l4
		*(*int32)(unsafe.Add(mBase, uint32(v16)+80)) = l3
		*(*int32)(unsafe.Add(mBase, uint32(v16)+76)) = l2
		*(*int32)(unsafe.Add(mBase, uint32(v16)+72)) = l1
		*(*int32)(unsafe.Add(mBase, uint32(v16)+64)) = v48
		v55 = v16 + int32(56)
		v56 = *(*float64)(unsafe.Add(mBase, uint32(v55)))
		v58 = v16 + int32(48)
		v59 = *(*float64)(unsafe.Add(mBase, uint32(v58)))
		v61 = v16 + int32(32)
		v62 = *(*float64)(unsafe.Add(mBase, uint32(v61)))
		if l5 != int64(0) {
			if int64(0) < l5 {
				v85 = base.F64_convert_i64_u(l5)
				v86 = v62
			} else {
				v69 = base.F64_mul(v62, float64(0.1))
				v70 = float64(1e+100)
				if base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v69)&int64(9223372036854775807)))|base.F64_gt(v69, v70) != 0 {
					v83 = v70
				} else {
					v79 = float64(1)
					if base.F64_le(v69, v79) != 0 {
						v83 = v79
					} else {
						v83 = base.F64_nearest(v69)
					}
				}
				v84 = *(*float64)(unsafe.Add(mBase, uint32(v61)))
				v85 = v83
				v86 = v84
			}
			if base.F64_gt(v85, v86) != 0 {
				v88 = v86
			} else {
				v88 = v85
			}
			if base.F64_gt(v62, float64(0)) != 0 {
				v95 = *(*float64)(unsafe.Add(mBase, uint32(v58)))
				*(*float64)(unsafe.Add(mBase, uint32(v58))) = base.F64_add(base.F64_div(base.F64_mul(base.F64_sub(v56, v59), v88), v62), v95)
				v98 = *(*float64)(unsafe.Add(mBase, uint32(v61)))
				v99 = v98
			} else {
				v99 = v86
			}
			v100 = base.F64_sub(v99, v88)
			if base.F64_lt(v100, float64(1)) != 0 {
				v103 = float64(1)
			} else {
				v103 = v100
			}
			*(*float64)(unsafe.Add(mBase, uint32(v61))) = v103
			v106 = v103
		} else {
			v106 = v62
		}
		if l6 != int64(0) {
			if int64(0) < l6 {
				v129 = base.F64_convert_i64_u(l6)
				v130 = v106
			} else {
				v113 = base.F64_mul(v62, float64(0.1))
				v114 = float64(1e+100)
				if base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v113)&int64(9223372036854775807)))|base.F64_gt(v113, v114) != 0 {
					v127 = v114
				} else {
					v123 = float64(1)
					if base.F64_le(v113, v123) != 0 {
						v127 = v123
					} else {
						v127 = base.F64_nearest(v113)
					}
				}
				v128 = *(*float64)(unsafe.Add(mBase, uint32(v61)))
				v129 = v127
				v130 = v128
			}
			if base.F64_gt(v129, v130) != 0 {
				v132 = v130
			} else {
				v132 = v129
			}
			if base.F64_gt(v62, float64(0)) != 0 {
				v138 = *(*float64)(unsafe.Add(mBase, uint32(v58)))
				*(*float64)(unsafe.Add(mBase, uint32(v55))) = base.F64_add(base.F64_div(base.F64_mul(base.F64_sub(v56, v59), v132), v62), v138)
			} else {
			}
			v141 = float64(1)
			if base.F64_lt(v132, v141) != 0 {
				v144 = v141
			} else {
				v144 = v132
			}
			*(*float64)(unsafe.Add(mBase, uint32(v61))) = v144
		} else {
		}
		return v16
	}
}
func F_create_s(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_palloc(m, int32(10))
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		if v3 == int32(0) {
			return int32(0)
		} else {
			*(*int64)(unsafe.Add(mBase, uint32(v3))) = int64(1)
			return v3 + int32(8)
		}
	}
}
func F_create_setop_path(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 float64, l7 float64) int32 {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v62 float64
	_ = v62
	var v63 float64
	_ = v63
	var v66 float64
	_ = v66
	var v67 float64
	_ = v67
	var v70 float64
	_ = v70
	var v71 float64
	_ = v71
	var v72 float64
	_ = v72
	var v76 int32
	_ = v76
	var v79 float64
	_ = v79
	var v86 float64
	_ = v86
	var v87 float64
	_ = v87
	var v90 float64
	_ = v90
	var v91 float64
	_ = v91
	var v92 float64
	_ = v92
	var v95 int32
	_ = v95
	var v98 float64
	_ = v98
	var v100 float64
	_ = v100
	var v106 int32
	_ = v106
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v124 float64
	_ = v124
	var v126 int32
	_ = v126
	var v130 float64
	_ = v130
	var v131 float64
	_ = v131
	var v134 float64
	_ = v134
	var v138 int32
	_ = v138
	v15 = F_palloc0(m, int32(104))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = l0
		*(*int64)(unsafe.Add(mBase, uint32(v15))) = int64(1593432867129)
		v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v23 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v15)+20)) = uint8(v23)
		*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v23
		*(*int32)(unsafe.Add(mBase, uint32(v15)+12)) = v22
		v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+26)))
		if v29 != int32(1) {
			v37 = v23
		} else {
			v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
			if v33 != int32(1) {
				v37 = int32(0)
			} else {
				v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+21)))
				v37 = v36
			}
		}
		v39 = v37 & int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(v15)+21)) = uint8(v39)
		v41 = *(*int32)(unsafe.Add(mBase, uint32(l2)+24))
		v42 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
		*(*int32)(unsafe.Add(mBase, uint32(v15)+24)) = v41 + v42
		if l4 == int32(0) {
			v47 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
			v48 = v47
		} else {
			v48 = int32(0)
		}
		*(*float64)(unsafe.Add(mBase, uint32(v15)+96)) = l6
		*(*int32)(unsafe.Add(mBase, uint32(v15)+88)) = l5
		*(*int32)(unsafe.Add(mBase, uint32(v15)+84)) = l4
		*(*int32)(unsafe.Add(mBase, uint32(v15)+80)) = l3
		*(*int32)(unsafe.Add(mBase, uint32(v15)+76)) = l2
		*(*int32)(unsafe.Add(mBase, uint32(v15)+72)) = l1
		*(*int32)(unsafe.Add(mBase, uint32(v15)+64)) = v48
		v56 = *(*int32)(unsafe.Add(mBase, uint32(l2)+40))
		v57 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
		v58 = v56 + v57
		*(*int32)(unsafe.Add(mBase, uint32(v15)+40)) = v58
		if l4 == int32(0) {
			v62 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
			v63 = *(*float64)(unsafe.Add(mBase, uint32(l2)+48))
			*(*float64)(unsafe.Add(mBase, uint32(v15)+48)) = base.F64_add(v62, v63)
			v66 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
			v67 = *(*float64)(unsafe.Add(mBase, uint32(l2)+56))
			v70 = *(*float64)(unsafe.Add(mBase, _c_F_create_setop_path[0]))
			v71 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
			v72 = *(*float64)(unsafe.Add(mBase, uint32(l2)+32))
			if l5 != 0 {
				v76 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
				v79 = base.F64_convert_i32_s(v76)
			} else {
				v79 = float64(0)
			}
			*(*float64)(unsafe.Add(mBase, uint32(v15)+56)) = base.F64_add(base.F64_mul(v70, l7), base.F64_add(base.F64_mul(base.F64_mul(v70, base.F64_add(v71, v72)), v79), base.F64_add(v66, v67)))
			*(*float64)(unsafe.Add(mBase, uint32(v15)+32)) = l7
			return v15
		} else {
			v86 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
			v87 = *(*float64)(unsafe.Add(mBase, uint32(l2)+56))
			v90 = *(*float64)(unsafe.Add(mBase, _c_F_create_setop_path[0]))
			v91 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
			v92 = *(*float64)(unsafe.Add(mBase, uint32(l2)+32))
			if l5 != 0 {
				v95 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
				v98 = base.F64_convert_i32_s(v95)
			} else {
				v98 = float64(0)
			}
			v100 = base.F64_add(base.F64_mul(base.F64_mul(v90, base.F64_add(v91, v92)), v98), base.F64_add(v86, v87))
			*(*float64)(unsafe.Add(mBase, uint32(v15)+48)) = v100
			*(*float64)(unsafe.Add(mBase, uint32(v15)+56)) = base.F64_add(base.F64_mul(v90, l7), v100)
			v106 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_create_setop_path[1])))
			if v106 == int32(0) {
				*(*int32)(unsafe.Add(mBase, uint32(v15)+40)) = v58 + int32(1)
			} else {
			}
			v112 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
			v113 = *(*int32)(unsafe.Add(mBase, uint32(v112)+32))
			v124 = *(*float64)(unsafe.Add(mBase, _c_F_create_setop_path[2]))
			v126 = *(*int32)(unsafe.Add(mBase, _c_F_create_setop_path[3]))
			v130 = base.F64_mul(base.F64_mul(v124, base.F64_convert_i32_s(v126)), float64(1024))
			v131 = float64(4.294967295e+09)
			if base.F64_lt(v130, v131) != 0 {
				v134 = v130
			} else {
				v134 = v131
			}
			if base.F64_gt(base.F64_mul(l6, base.F64_convert_i32_u((v113+int32(7))&int32(-8)+int32(16))), base.F64_convert_i32_u(base.I32_trunc_sat_f64_u(v134))) != 0 {
				v138 = *(*int32)(unsafe.Add(mBase, uint32(v15)+40))
				*(*int32)(unsafe.Add(mBase, uint32(v15)+40)) = v138 + int32(1)
			} else {
			}
			*(*float64)(unsafe.Add(mBase, uint32(v15)+32)) = l7
			return v15
		}
	}
}
func F_create_subqueryscan_path(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
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
	var v38 int32
	_ = v38
	v10 = F_palloc0(m, int32(80))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = l1
		*(*int64)(unsafe.Add(mBase, uint32(v10))) = int64(1490353651999)
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
		*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = v17
		v19 = F_get_baserel_parampathinfo(m, l0, l1, l5)
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int32(0)
		} else {
			v21 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v10)+20)) = uint8(v21)
			*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v19
			v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+26)))
			if v24 == int32(1) {
				v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+21)))
				v29 = v27
			} else {
				v29 = int32(0)
			}
			v31 = v29 & int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(v10)+21)) = uint8(v31)
			v33 = *(*int32)(unsafe.Add(mBase, uint32(l2)+24))
			*(*int32)(unsafe.Add(mBase, uint32(v10)+72)) = l2
			*(*int32)(unsafe.Add(mBase, uint32(v10)+64)) = l4
			*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = v33
			F_cost_subqueryscan(m, v10, l0, l1, v19, l3)
			mBase = m.M
			v38 = m.ExcPending
			if v38 != 0 {
				return int32(0)
			} else {
				return v10
			}
		}
	}
}
func F_currval_oid(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v29 int64
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	v4 = m.G0
	v6 = v4 - int32(32)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	F_init_sequence(m, v8, v6+int32(28), v6+int32(24))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(v6)+28))
		v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
		v20 = *(*int32)(unsafe.Add(mBase, _c_F_currval_oid[0]))
		v22 = F_pg_class_aclcheck(m, v18, v20, int64(258))
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return int32(0)
		} else {
			if v22 == int32(0) {
				v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+12)))
				if v26 == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v65 = m.ExcPending
					if v65 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(325))
						mBase = m.M
						v68 = m.ExcPending
						if v68 != 0 {
							return int32(0)
						} else {
							v69 = *(*int32)(unsafe.Add(mBase, uint32(v6)+24))
							v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)+48))
							*(*int32)(unsafe.Add(mBase, uint32(v6))) = v70 + int32(4)
							F_errmsg(m, int32(_a_F_currval_oid_0), v6)
							mBase = m.M
							v76 = m.ExcPending
							if v76 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_currval_oid_1), int32(887), int32(_a_F_currval_oid_2))
								mBase = m.M
								v81 = m.ExcPending
								if v81 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					}
				} else {
					v29 = *(*int64)(unsafe.Add(mBase, uint32(v17)+16))
					v30 = *(*int32)(unsafe.Add(mBase, uint32(v6)+24))
					F_relation_close(m, v30, int32(0))
					mBase = m.M
					v33 = m.ExcPending
					if v33 != 0 {
						return int32(0)
					} else {
						v34 = F_Int64GetDatum(m, v29)
						mBase = m.M
						v35 = m.ExcPending
						if v35 != 0 {
							return int32(0)
						} else {
							m.G0 = v6 + int32(32)
							return v34
						}
					}
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v43 = m.ExcPending
				if v43 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(16797828))
					mBase = m.M
					v46 = m.ExcPending
					if v46 != 0 {
						return int32(0)
					} else {
						v47 = *(*int32)(unsafe.Add(mBase, uint32(v6)+24))
						v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)+48))
						*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = v48 + int32(4)
						F_errmsg(m, int32(_a_F_currval_oid_3), v6+int32(16))
						mBase = m.M
						v56 = m.ExcPending
						if v56 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_currval_oid_1), int32(881), int32(_a_F_currval_oid_2))
							mBase = m.M
							v61 = m.ExcPending
							if v61 != 0 {
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
