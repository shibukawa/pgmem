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
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	v5 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v5 != v6 {
		if base.Ui32(v5) < base.Ui32(v6) {
			v11 = int32(-1)
		} else {
			v11 = int32(1)
		}
		return v11
	} else {
		v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
		v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
		if v13 != v14 {
			if base.Ui32(v13) < base.Ui32(v14) {
				v19 = int32(-1)
			} else {
				v19 = int32(1)
			}
			return v19
		} else {
			v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+2)))
			v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+2)))
			if v22 != v23 {
				if base.Ui32(v22) < base.Ui32(v23) {
					v28 = int32(-1)
				} else {
					v28 = int32(1)
				}
				v29 = v28
			} else {
				v29 = int32(0)
			}
			return v29
		}
	}
}
func F_CheckCmdReplicaIdentity(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int64
	_ = v68
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
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
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v100 int32
	_ = v100
	var v107 int32
	_ = v107
	var v124 int32
	_ = v124
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
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v154 int32
	_ = v154
	var v182 int32
	_ = v182
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v227 int32
	_ = v227
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v272 int32
	_ = v272
	var v285 int32
	_ = v285
	var v303 int32
	_ = v303
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
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
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v335 int32
	_ = v335
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v342 int32
	_ = v342
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v349 int32
	_ = v349
	var v352 int32
	_ = v352
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v386 int32
	_ = v386
	var v389 int32
	_ = v389
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v428 int32
	_ = v428
	var v431 int32
	_ = v431
	var v433 int32
	_ = v433
	var v436 int32
	_ = v436
	var v447 int32
	_ = v447
	var v450 int32
	_ = v450
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v457 int32
	_ = v457
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v462 int32
	_ = v462
	var v464 int32
	_ = v464
	var v468 int32
	_ = v468
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v488 int32
	_ = v488
	var v490 int32
	_ = v490
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v513 int32
	_ = v513
	var v516 int32
	_ = v516
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v535 int32
	_ = v535
	var v552 int32
	_ = v552
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v564 int32
	_ = v564
	var v589 int32
	_ = v589
	var v624 int32
	_ = v624
	var v656 int32
	_ = v656
	var v657 int32
	_ = v657
	var v660 int32
	_ = v660
	var v664 int32
	_ = v664
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
	var v678 int32
	_ = v678
	var v681 int32
	_ = v681
	var v684 int32
	_ = v684
	var v686 int32
	_ = v686
	var v687 int32
	_ = v687
	var v690 int32
	_ = v690
	var v694 int32
	_ = v694
	var v695 int32
	_ = v695
	var v705 int32
	_ = v705
	var v706 int32
	_ = v706
	var v709 int32
	_ = v709
	var v713 int32
	_ = v713
	var v716 int32
	_ = v716
	var v718 int32
	_ = v718
	var v721 int32
	_ = v721
	var v728 int32
	_ = v728
	var v730 int32
	_ = v730
	var v738 int32
	_ = v738
	var v739 int32
	_ = v739
	var v752 int32
	_ = v752
	var v764 int32
	_ = v764
	var v787 int32
	_ = v787
	var v788 int32
	_ = v788
	var v791 int32
	_ = v791
	var v798 int32
	_ = v798
	var v799 int32
	_ = v799
	var v809 int32
	_ = v809
	var v812 int32
	_ = v812
	var v813 int32
	_ = v813
	var v814 int32
	_ = v814
	var v815 int32
	_ = v815
	var v816 int32
	_ = v816
	var v817 int32
	_ = v817
	var v818 int32
	_ = v818
	var v819 int32
	_ = v819
	var v820 int32
	_ = v820
	var v821 int32
	_ = v821
	var v822 int32
	_ = v822
	var v824 int32
	_ = v824
	var v830 int32
	_ = v830
	var v839 int32
	_ = v839
	var v841 int32
	_ = v841
	var v842 int32
	_ = v842
	var v845 int32
	_ = v845
	var v849 int32
	_ = v849
	var v852 int32
	_ = v852
	var v854 int32
	_ = v854
	var v857 int32
	_ = v857
	var v864 int32
	_ = v864
	var v866 int32
	_ = v866
	var v874 int32
	_ = v874
	var v875 int32
	_ = v875
	var v888 int32
	_ = v888
	var v919 int32
	_ = v919
	var v921 int32
	_ = v921
	var v923 int32
	_ = v923
	var v925 int32
	_ = v925
	var v926 int32
	_ = v926
	var v929 int32
	_ = v929
	var v962 int32
	_ = v962
	var v965 int32
	_ = v965
	var v966 int32
	_ = v966
	var v967 int32
	_ = v967
	var v969 int32
	_ = v969
	var v971 int32
	_ = v971
	var v973 int32
	_ = v973
	var v976 int32
	_ = v976
	var v977 int32
	_ = v977
	var v978 int32
	_ = v978
	var v980 int32
	_ = v980
	var v982 int32
	_ = v982
	var v1013 int32
	_ = v1013
	var v1014 int32
	_ = v1014
	var v1017 int32
	_ = v1017
	var v1020 int32
	_ = v1020
	var v1023 int32
	_ = v1023
	var v1026 int32
	_ = v1026
	var v1027 int32
	_ = v1027
	var v1030 int32
	_ = v1030
	var v1033 int32
	_ = v1033
	var v1036 int32
	_ = v1036
	var v1037 int32
	_ = v1037
	var v1040 int32
	_ = v1040
	var v1043 int32
	_ = v1043
	var v1046 int32
	_ = v1046
	var v1047 int32
	_ = v1047
	var v1051 int32
	_ = v1051
	var v1052 int32
	_ = v1052
	var v1082 int32
	_ = v1082
	var v1084 int32
	_ = v1084
	var v1087 int32
	_ = v1087
	var v1088 int32
	_ = v1088
	var v1091 int32
	_ = v1091
	var v1094 int32
	_ = v1094
	var v1095 int32
	_ = v1095
	var v1097 int32
	_ = v1097
	var v1099 int64
	_ = v1099
	var v1137 int32
	_ = v1137
	var v1141 int32
	_ = v1141
	var v1146 int32
	_ = v1146
	var v1149 int32
	_ = v1149
	var v1152 int32
	_ = v1152
	var v1156 int32
	_ = v1156
	var v1160 int32
	_ = v1160
	var v1163 int32
	_ = v1163
	var v1164 int32
	_ = v1164
	var v1172 int32
	_ = v1172
	var v1176 int32
	_ = v1176
	var v1181 int32
	_ = v1181
	var v1185 int32
	_ = v1185
	var v1188 int32
	_ = v1188
	var v1191 int32
	_ = v1191
	var v1195 int32
	_ = v1195
	var v1196 int32
	_ = v1196
	var v1197 int32
	_ = v1197
	var v1198 int32
	_ = v1198
	var v1199 int32
	_ = v1199
	var v1204 int32
	_ = v1204
	var v1209 int32
	_ = v1209
	var v1246 int32
	_ = v1246
	var v1249 int32
	_ = v1249
	var v1250 int32
	_ = v1250
	var v1258 int32
	_ = v1258
	var v1262 int32
	_ = v1262
	var v1267 int32
	_ = v1267
	var v1271 int32
	_ = v1271
	var v1274 int32
	_ = v1274
	var v1275 int32
	_ = v1275
	var v1283 int32
	_ = v1283
	var v1287 int32
	_ = v1287
	var v1292 int32
	_ = v1292
	var v1296 int32
	_ = v1296
	var v1299 int32
	_ = v1299
	var v1300 int32
	_ = v1300
	var v1308 int32
	_ = v1308
	var v1312 int32
	_ = v1312
	var v1317 int32
	_ = v1317
	var v1321 int32
	_ = v1321
	var v1324 int32
	_ = v1324
	var v1325 int32
	_ = v1325
	var v1333 int32
	_ = v1333
	var v1337 int32
	_ = v1337
	var v1342 int32
	_ = v1342
	var v1346 int32
	_ = v1346
	var v1349 int32
	_ = v1349
	var v1350 int32
	_ = v1350
	var v1358 int32
	_ = v1358
	var v1362 int32
	_ = v1362
	var v1367 int32
	_ = v1367
	var v1371 int32
	_ = v1371
	var v1374 int32
	_ = v1374
	var v1375 int32
	_ = v1375
	var v1381 int32
	_ = v1381
	var v1385 int32
	_ = v1385
	var v1390 int32
	_ = v1390
	var v1394 int32
	_ = v1394
	var v1397 int32
	_ = v1397
	var v1398 int32
	_ = v1398
	var v1406 int32
	_ = v1406
	var v1410 int32
	_ = v1410
	var v1415 int32
	_ = v1415
	v3 = int32(0)
	v29 = m.G0
	v31 = v29 - int32(128)
	m.G0 = v31
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+119)))
	if v34 == int32(112) {
		goto L8
	} else {
		goto L9
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1394 = m.ExcPending
	if v1394 != 0 {
		goto L24
	} else {
		goto L280
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1371 = m.ExcPending
	if v1371 != 0 {
		goto L24
	} else {
		goto L275
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1346 = m.ExcPending
	if v1346 != 0 {
		goto L24
	} else {
		goto L270
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1321 = m.ExcPending
	if v1321 != 0 {
		goto L24
	} else {
		goto L265
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1296 = m.ExcPending
	if v1296 != 0 {
		goto L24
	} else {
		goto L260
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1271 = m.ExcPending
	if v1271 != 0 {
		goto L24
	} else {
		goto L255
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1246 = m.ExcPending
	if v1246 != 0 {
		goto L24
	} else {
		goto L250
	}
L8:
	;
	m.G0 = v31 + int32(128)
	return
L9:
	;
	switch l1 - int32(2) {
	case 0, 2:
		goto L10
	default:
		goto L8
	}
L10:
	;
	v40 = v31 + int32(118)
	v41 = m.G0
	v43 = v41 - int32(16)
	m.G0 = v43
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46)+119)))
	switch v47 - int32(112) {
	case 0, 2:
		goto L12
	default:
		v59 = v3
		goto L11
	}
L11:
	;
	if v59 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L12:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	goto L13
L13:
	;
	if base.Ui32(v50) < base.Ui32(int32(12000)) {
		v59 = v3
		goto L11
	} else {
		goto L14
	}
L14:
	;
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46)+118)))
	v59 = base.B2i32(v53 == int32(112)) & base.B2i32(base.Ui32(int32(16383)) < base.Ui32(v50))
	goto L11
L15:
	;
	if l1 == int32(2) {
		goto L226
	} else {
		goto L227
	}
L16:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1137 = m.ExcPending
	if v1137 != 0 {
		goto L24
	} else {
		goto L222
	}
L17:
	;
	m.G0 = v43 + int32(16)
	goto L15
L18:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v40))) = int64(72340172821233664)
	v65 = int32(257)
	*(*uint16)(unsafe.Add(mBase, uint32(v40)+8)) = uint16(v65)
	goto L17
L19:
	;
	goto L20
L20:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+176))
	if v67 != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v68 = *(*int64)(unsafe.Add(mBase, uint32(v67)))
	*(*int64)(unsafe.Add(mBase, uint32(v40))) = v68
	v70 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v67)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(v40)+8)) = uint16(v70)
	goto L17
L22:
	;
	goto L23
L23:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v40))) = int64(72340172821233664)
	v74 = int32(257)
	*(*uint16)(unsafe.Add(mBase, uint32(v40)+8)) = uint16(v74)
	v76 = F_GetRelationPublications(m, v45)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	return
L25:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v78)+68))
	v80 = F_GetSchemaPublications(m, v79)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L24
	} else {
		goto L26
	}
L26:
	;
	v82 = F_list_concat_unique_oid(m, v76, v80)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L24
	} else {
		goto L27
	}
L27:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84)+131)))
	if v85 != int32(1) {
		v182 = v82
		v197 = v3
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v200 = m.G0
	v202 = v200 - int32(48)
	m.G0 = v202
	v206 = F_table_open(m, int32(6104), int32(1))
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L24
	} else {
		goto L43
	}
L29:
	;
	v88 = F_get_partition_ancestors(m, v45)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L24
	} else {
		goto L30
	}
L30:
	;
	if v88 == int32(0) {
		v182 = v82
		v197 = v3
		goto L28
	} else {
		goto L31
	}
L31:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v88)+4))
	if int32(0) < v92 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v100 = int32(0)
	v107 = v82
	goto L35
L33:
	;
	v154 = v82
	goto L34
L34:
	;
	v182 = v154
	v197 = v88
	goto L28
L35:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v88)+12))
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v124+v100<<(uint(int32(2))%32))))
	v129 = F_GetRelationPublications(m, v128)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L24
	} else {
		goto L37
	}
L36:
	;
	v154 = v137
	goto L34
L37:
	;
	v131 = F_list_concat_unique_oid(m, v107, v129)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L24
	} else {
		goto L38
	}
L38:
	;
	v133 = F_get_rel_namespace(m, v128)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L24
	} else {
		goto L39
	}
L39:
	;
	v135 = F_GetSchemaPublications(m, v133)
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L24
	} else {
		goto L40
	}
L40:
	;
	v137 = F_list_concat_unique_oid(m, v131, v135)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L24
	} else {
		goto L41
	}
L41:
	;
	v140 = v100 + int32(1)
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v88)+4))
	if v140 < v141 {
		v100 = v140
		v107 = v137
		goto L35
	} else {
		goto L42
	}
L42:
	;
	goto L36
L43:
	;
	F_ScanKeyInit(m, v202, int32(4), int32(3), int32(60), int32(1))
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L24
	} else {
		goto L44
	}
L44:
	;
	v214 = int32(0)
	v218 = F_systable_beginscan(m, v206, v214, v214, v214, int32(1), v202)
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L24
	} else {
		goto L45
	}
L45:
	;
	v227 = int32(0)
	goto L46
L46:
	;
	v248 = F_systable_getnext(m, v218)
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L24
	} else {
		goto L48
	}
L47:
	;
	F_systable_endscan(m, v218)
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L24
	} else {
		goto L53
	}
L48:
	;
	if v248 != 0 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v248)+16))
	v251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v250)+22)))
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v250+v251)))
	v254 = F_lappend_oid(m, v227, v253)
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L24
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	goto L47
L52:
	;
	v227 = v254
	goto L46
L53:
	;
	F_sequence_close(m, v206, int32(1))
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L24
	} else {
		goto L54
	}
L54:
	;
	m.G0 = v202 + int32(48)
	v264 = F_list_concat_unique_oid(m, v182, v227)
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L24
	} else {
		goto L56
	}
L55:
	;
	v1082 = *(*int32)(unsafe.Add(mBase, uint32(l0)+176))
	if v1082 != 0 {
		goto L217
	} else {
		goto L218
	}
L56:
	;
	if v264 == int32(0) {
		goto L55
	} else {
		goto L57
	}
L57:
	;
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v264)+4))
	if v268 <= int32(0) {
		goto L55
	} else {
		goto L58
	}
L58:
	;
	v272 = v31 + int32(122)
	v285 = int32(0)
	goto L59
L59:
	;
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v264)+12))
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v303+v285<<(uint(int32(2))%32))))
	v308 = F_SearchSysCache1(m, int32(51), v307)
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L24
	} else {
		goto L61
	}
L60:
	;
	goto L55
L61:
	;
	if v308 == int32(0) {
		goto L16
	} else {
		goto L62
	}
L62:
	;
	v312 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40))))
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v308)+16))
	v314 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v313)+22)))
	v315 = v313 + v314
	v316 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v315)+73)))
	v317 = v312 | v316
	*(*uint8)(unsafe.Add(mBase, uint32(v40))) = uint8(v317)
	v319 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+1)))
	v320 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v315)+74)))
	v321 = v319 | v320
	*(*uint8)(unsafe.Add(mBase, uint32(v40)+1)) = uint8(v321)
	v323 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+2)))
	v324 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v315)+75)))
	v325 = v323 | v324
	*(*uint8)(unsafe.Add(mBase, uint32(v40)+2)) = uint8(v325)
	v327 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+3)))
	v328 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v315)+76)))
	v329 = v327 | v328
	*(*uint8)(unsafe.Add(mBase, uint32(v40)+3)) = uint8(v329)
	v331 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v315)+72)))
	if v331 != 0 {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v447 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v315)+74)))
	if v447 == int32(0) {
		goto L105
	} else {
		goto L106
	}
L64:
	;
	v332 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v315)+74)))
	if v332 == int32(0) {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v335 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v315)+75)))
	if v335 != int32(1) {
		goto L63
	} else {
		goto L68
	}
L66:
	;
	goto L67
L67:
	;
	v338 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v315)+77)))
	v339 = int32(0)
	v340 = m.G0
	v342 = v340 - int32(32)
	m.G0 = v342
	v344 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v345 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v344)+130)))
	if v345 == int32(102) {
		v416 = v339
		goto L69
	} else {
		goto L70
	}
L68:
	;
	goto L67
L69:
	;
	m.G0 = v342 + int32(32)
	if v416 == int32(0) {
		goto L63
	} else {
		goto L99
	}
L70:
	;
	v349 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	if v338 == int32(0) {
		v359 = v349
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v360 = F_SearchSysCache2(m, int32(53), v359, v307)
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L24
	} else {
		goto L78
	}
L72:
	;
	v352 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v344)+131)))
	if v352 != int32(1) {
		v359 = v349
		goto L71
	} else {
		goto L73
	}
L73:
	;
	v355 = F_GetTopMostAncestorInPublication(m, v307, v197)
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L24
	} else {
		goto L74
	}
L74:
	;
	if v355 != 0 {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v357 = v355
	goto L77
L76:
	;
	v357 = v349
	goto L77
L77:
	;
	v359 = v357
	goto L71
L78:
	;
	if v360 == int32(0) {
		v416 = v339
		goto L69
	} else {
		goto L79
	}
L79:
	;
	v368 = F_SysCacheGetAttr(m, int32(53), v360, int32(4), v342+int32(31))
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L24
	} else {
		goto L80
	}
L80:
	;
	v370 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v342)+31)))
	if v370 != 0 {
		v410 = v339
		goto L81
	} else {
		goto L82
	}
L81:
	;
	F_ReleaseCatCache(m, v360)
	mBase = m.M
	v415 = m.ExcPending
	if v415 != 0 {
		goto L24
	} else {
		goto L98
	}
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v342)+16)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v342)+24)) = v359
	*(*int32)(unsafe.Add(mBase, uint32(v342)+20)) = v349
	*(*uint8)(unsafe.Add(mBase, uint32(v342)+16)) = uint8(v338)
	v377 = F_RelationGetIndexAttrBitmap(m, l0, int32(2))
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L24
	} else {
		goto L83
	}
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v342)+12)) = v377
	v380 = F_text_to_cstring(m, v368)
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L24
	} else {
		goto L84
	}
L84:
	;
	v382 = F_stringToNode(m, v380)
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L24
	} else {
		goto L85
	}
L85:
	;
	if v382 == int32(0) {
		v410 = v339
		goto L81
	} else {
		goto L86
	}
L86:
	;
	v386 = *(*int32)(unsafe.Add(mBase, uint32(v382)))
	if v386 == int32(6) {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v389 = int32(*(*int16)(unsafe.Add(mBase, uint32(v382)+8)))
	if v338 != 0 {
		goto L90
	} else {
		goto L91
	}
L88:
	;
	goto L89
L89:
	;
	v408 = F_expression_tree_walker_impl(m, v382, int32(567), v342+int32(12))
	mBase = m.M
	v409 = m.ExcPending
	if v409 != 0 {
		goto L24
	} else {
		goto L97
	}
L90:
	;
	v391 = F_get_attname(m, v359, v389, int32(0))
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L24
	} else {
		goto L93
	}
L91:
	;
	v395 = v389
	goto L92
L92:
	;
	v399 = F_bms_is_member(m, v395+int32(7), v377)
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L24
	} else {
		goto L95
	}
L93:
	;
	v393 = F_get_attnum(m, v349, v391)
	mBase = m.M
	v394 = m.ExcPending
	if v394 != 0 {
		goto L24
	} else {
		goto L94
	}
L94:
	;
	v395 = v393
	goto L92
L95:
	;
	if v399 == int32(0) {
		v410 = int32(1)
		goto L81
	} else {
		goto L96
	}
L96:
	;
	goto L89
L97:
	;
	v410 = v408
	goto L81
L98:
	;
	v416 = v410
	goto L69
L99:
	;
	v428 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v315)+74)))
	if v428 == int32(1) {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	v431 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v272))) = uint8(v431)
	goto L102
L101:
	;
	goto L102
L102:
	;
	v433 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v315)+75)))
	if v433 != int32(1) {
		goto L63
	} else {
		goto L103
	}
L103:
	;
	v436 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v40)+5)) = uint8(v436)
	goto L63
L104:
	;
	F_ReleaseCatCache(m, v308)
	mBase = m.M
	v1013 = m.ExcPending
	if v1013 != 0 {
		goto L24
	} else {
		goto L198
	}
L105:
	;
	v450 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v315)+75)))
	if v450 != int32(1) {
		goto L104
	} else {
		goto L108
	}
L106:
	;
	goto L107
L107:
	;
	v453 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v315)+77)))
	v454 = int32(*(*int8)(unsafe.Add(mBase, uint32(v315)+78)))
	v455 = m.G0
	v457 = v455 - int32(16)
	m.G0 = v457
	v459 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v460 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v457)+12)) = v460
	v462 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v464 = v43 + int32(15)
	*(*uint8)(unsafe.Add(mBase, uint32(v464))) = uint8(v460)
	v468 = v43 + int32(14)
	*(*uint8)(unsafe.Add(mBase, uint32(v468))) = uint8(v460)
	if v453 == v460 {
		v481 = v459
		goto L109
	} else {
		goto L110
	}
L108:
	;
	goto L107
L109:
	;
	v482 = F_GetPublication(m, v307)
	mBase = m.M
	v483 = m.ExcPending
	if v483 != 0 {
		goto L24
	} else {
		goto L116
	}
L110:
	;
	v473 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v474 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v473)+131)))
	if v474 != int32(1) {
		v481 = v459
		goto L109
	} else {
		goto L111
	}
L111:
	;
	v477 = F_GetTopMostAncestorInPublication(m, v307, v197)
	mBase = m.M
	v478 = m.ExcPending
	if v478 != 0 {
		goto L24
	} else {
		goto L112
	}
L112:
	;
	if v477 != 0 {
		goto L113
	} else {
		goto L114
	}
L113:
	;
	v479 = v477
	goto L115
L114:
	;
	v479 = v459
	goto L115
L115:
	;
	v481 = v479
	goto L109
L116:
	;
	v485 = v457 + int32(12)
	v486 = m.G0
	v488 = v486 - int32(16)
	m.G0 = v488
	v490 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v482)+8)))
	if v490 != 0 {
		goto L117
	} else {
		goto L118
	}
L117:
	;
	m.G0 = v488 + int32(16)
	v656 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v657 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v656)+130)))
	if v657 != int32(102) {
		goto L138
	} else {
		goto L139
	}
L118:
	;
	v492 = *(*int32)(unsafe.Add(mBase, uint32(v482)))
	v493 = F_SearchSysCache2(m, int32(53), v481, v492)
	mBase = m.M
	v494 = m.ExcPending
	if v494 != 0 {
		goto L24
	} else {
		goto L119
	}
L119:
	;
	if v493 == int32(0) {
		goto L117
	} else {
		goto L120
	}
L120:
	;
	v501 = F_SysCacheGetAttr(m, int32(53), v493, int32(5), v488+int32(15))
	mBase = m.M
	v502 = m.ExcPending
	if v502 != 0 {
		goto L24
	} else {
		goto L121
	}
L121:
	;
	v503 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v488)+15)))
	if v485 == int32(0) {
		goto L122
	} else {
		goto L123
	}
L122:
	;
	F_ReleaseCatCache(m, v493)
	mBase = m.M
	v624 = m.ExcPending
	if v624 != 0 {
		goto L24
	} else {
		goto L136
	}
L123:
	;
	if v503&int32(1) != 0 {
		goto L122
	} else {
		goto L124
	}
L124:
	;
	v508 = *(*int32)(unsafe.Add(mBase, uint32(v485)))
	v509 = F_pg_detoast_datum(m, v501)
	mBase = m.M
	v510 = m.ExcPending
	if v510 != 0 {
		goto L24
	} else {
		goto L125
	}
L125:
	;
	v513 = *(*int32)(unsafe.Add(mBase, uint32(v509)+8))
	if v513 == int32(0) {
		goto L126
	} else {
		goto L127
	}
L126:
	;
	v516 = *(*int32)(unsafe.Add(mBase, uint32(v509)+4))
	v523 = (v516<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L128
L127:
	;
	v523 = v513
	goto L128
L128:
	;
	v524 = *(*int32)(unsafe.Add(mBase, uint32(v509+int32(16))))
	v525 = int32(0)
	if v525 < v524 {
		goto L129
	} else {
		goto L130
	}
L129:
	;
	v535 = v525
	v552 = v508
	goto L132
L130:
	;
	v589 = v508
	goto L131
L131:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v485))) = v589
	goto L122
L132:
	;
	v560 = int32(*(*int16)(unsafe.Add(mBase, uint32(v509+v523+v535<<(uint(int32(1))%32)))))
	v561 = F_bms_add_member(m, v552, v560)
	mBase = m.M
	v562 = m.ExcPending
	if v562 != 0 {
		goto L24
	} else {
		goto L134
	}
L133:
	;
	v589 = v561
	goto L131
L134:
	;
	v564 = v535 + int32(1)
	if v564 != v524 {
		v535 = v564
		v552 = v561
		goto L132
	} else {
		goto L135
	}
L135:
	;
	goto L133
L136:
	;
	goto L117
L137:
	;
	m.G0 = v457 + int32(16)
	if v929&int32(1) == int32(0) {
		goto L104
	} else {
		goto L193
	}
L138:
	;
	v694 = F_RelationGetIndexAttrBitmap(m, l0, int32(2))
	mBase = m.M
	v695 = m.ExcPending
	if v695 != 0 {
		goto L24
	} else {
		goto L150
	}
L139:
	;
	v660 = *(*int32)(unsafe.Add(mBase, uint32(v457)+12))
	*(*uint8)(unsafe.Add(mBase, uint32(v464))) = uint8(base.B2i32(v660 != int32(0)))
	v664 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	if v454 == int32(115) {
		v676 = v664
		goto L140
	} else {
		goto L141
	}
L140:
	;
	v678 = *(*int32)(unsafe.Add(mBase, uint32(v676)+16))
	if v678 == int32(0) {
		goto L144
	} else {
		goto L145
	}
L141:
	;
	v667 = *(*int32)(unsafe.Add(mBase, uint32(v664)+16))
	if v667 == int32(0) {
		v676 = v664
		goto L140
	} else {
		goto L142
	}
L142:
	;
	v670 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v667)+17)))
	if v670 != int32(1) {
		v676 = v664
		goto L140
	} else {
		goto L143
	}
L143:
	;
	v673 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v468))) = uint8(v673)
	v675 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v676 = v675
	goto L140
L144:
	;
	v686 = int32(1)
	v687 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v468))))
	if v687 != v686 {
		goto L138
	} else {
		goto L147
	}
L145:
	;
	v681 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v678)+18)))
	if v681 != int32(1) {
		goto L144
	} else {
		goto L146
	}
L146:
	;
	v684 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v468))) = uint8(v684)
	goto L144
L147:
	;
	v690 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v464))))
	if v690 != 0 {
		v929 = v686
		goto L137
	} else {
		goto L148
	}
L148:
	;
	goto L138
L149:
	;
	v919 = *(*int32)(unsafe.Add(mBase, uint32(v457)+12))
	F_bms_free(m, v919)
	mBase = m.M
	v921 = m.ExcPending
	if v921 != 0 {
		goto L24
	} else {
		goto L190
	}
L150:
	;
	if v694 == int32(0) {
		goto L153
	} else {
		goto L154
	}
L151:
	;
	if v752 < int32(0) {
		goto L149
	} else {
		goto L162
	}
L152:
	;
	v752 = base.I32_ctz(v738) | v739<<(uint(int32(5))%32)
	goto L151
L153:
	;
	v752 = int32(-2)
	goto L151
L154:
	;
	v705 = base.I32_div_s(int32(0), int32(32))
	v706 = *(*int32)(unsafe.Add(mBase, uint32(v694)+4))
	if v706 <= v705 {
		goto L153
	} else {
		goto L155
	}
L155:
	;
	v709 = v694 + int32(8)
	v713 = *(*int32)(unsafe.Add(mBase, uint32(v709+v705<<(uint(int32(2))%32))))
	v716 = v713 & int32(-1)
	if v716 != 0 {
		v738 = v716
		v739 = v705
		goto L152
	} else {
		goto L156
	}
L156:
	;
	v718 = v705 + int32(1)
	if v718 == v706 {
		goto L153
	} else {
		goto L157
	}
L157:
	;
	v721 = v718
	goto L158
L158:
	;
	v728 = *(*int32)(unsafe.Add(mBase, uint32(v709+v721<<(uint(int32(2))%32))))
	if v728 != 0 {
		v738 = v728
		v739 = v721
		goto L152
	} else {
		goto L160
	}
L159:
	;
	goto L153
L160:
	;
	v730 = v721 + int32(1)
	if v730 != v706 {
		v721 = v730
		goto L158
	} else {
		goto L161
	}
L161:
	;
	goto L159
L162:
	;
	v764 = v752
	goto L163
L163:
	;
	v787 = base.I32_extend16_s(v764 - int32(7))
	v788 = *(*int32)(unsafe.Add(mBase, uint32(v457)+12))
	if v788 == int32(0) {
		goto L166
	} else {
		goto L167
	}
L164:
	;
	goto L149
L165:
	;
	if v694 == int32(0) {
		goto L180
	} else {
		goto L181
	}
L166:
	;
	v791 = *(*int32)(unsafe.Add(mBase, uint32(v462)))
	v798 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v462+int32(10)+v791<<(uint(int32(4))%32)+v787*int32(100)))))
	v799 = int32(115)
	if base.B2i32(base.B2i32(v798 == v799)&base.B2i32(v454 != v799) == int32(0))&base.B2i32(v798 != int32(118)) != 0 {
		goto L165
	} else {
		goto L169
	}
L167:
	;
	goto L168
L168:
	;
	if v453 != 0 {
		goto L170
	} else {
		goto L171
	}
L169:
	;
	v809 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v468))) = uint8(v809)
	goto L149
L170:
	;
	v812 = F_get_attname(m, v459, v787, int32(0))
	mBase = m.M
	v813 = m.ExcPending
	if v813 != 0 {
		goto L24
	} else {
		goto L173
	}
L171:
	;
	v817 = v787
	v818 = v788
	goto L172
L172:
	;
	v819 = F_bms_is_member(m, v817, v818)
	mBase = m.M
	v820 = m.ExcPending
	if v820 != 0 {
		goto L24
	} else {
		goto L175
	}
L173:
	;
	v814 = F_get_attnum(m, v481, v812)
	mBase = m.M
	v815 = m.ExcPending
	if v815 != 0 {
		goto L24
	} else {
		goto L174
	}
L174:
	;
	v816 = *(*int32)(unsafe.Add(mBase, uint32(v457)+12))
	v817 = v814
	v818 = v816
	goto L172
L175:
	;
	v821 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v464))))
	v822 = int32(1)
	v824 = v821 | (v819 ^ v822)
	*(*uint8)(unsafe.Add(mBase, uint32(v464))) = uint8(v824)
	if v824&v822 == int32(0) {
		goto L165
	} else {
		goto L176
	}
L176:
	;
	v830 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v468))))
	if v830 != 0 {
		goto L149
	} else {
		goto L177
	}
L177:
	;
	goto L165
L178:
	;
	if int32(0) <= v888 {
		v764 = v888
		goto L163
	} else {
		goto L189
	}
L179:
	;
	v888 = base.I32_ctz(v874) | v875<<(uint(int32(5))%32)
	goto L178
L180:
	;
	v888 = int32(-2)
	goto L178
L181:
	;
	v839 = v764 + int32(1)
	v841 = base.I32_div_s(v839, int32(32))
	v842 = *(*int32)(unsafe.Add(mBase, uint32(v694)+4))
	if v842 <= v841 {
		goto L180
	} else {
		goto L182
	}
L182:
	;
	v845 = v694 + int32(8)
	v849 = *(*int32)(unsafe.Add(mBase, uint32(v845+v841<<(uint(int32(2))%32))))
	v852 = v849 & (int32(-1) << (uint(v839) % 32))
	if v852 != 0 {
		v874 = v852
		v875 = v841
		goto L179
	} else {
		goto L183
	}
L183:
	;
	v854 = v841 + int32(1)
	if v854 == v842 {
		goto L180
	} else {
		goto L184
	}
L184:
	;
	v857 = v854
	goto L185
L185:
	;
	v864 = *(*int32)(unsafe.Add(mBase, uint32(v845+v857<<(uint(int32(2))%32))))
	if v864 != 0 {
		v874 = v864
		v875 = v857
		goto L179
	} else {
		goto L187
	}
L186:
	;
	goto L180
L187:
	;
	v866 = v857 + int32(1)
	if v866 != v842 {
		v857 = v866
		goto L185
	} else {
		goto L188
	}
L188:
	;
	goto L186
L189:
	;
	goto L164
L190:
	;
	F_bms_free(m, v694)
	mBase = m.M
	v923 = m.ExcPending
	if v923 != 0 {
		goto L24
	} else {
		goto L191
	}
L191:
	;
	v925 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v464))))
	if v925 != 0 {
		v929 = int32(1)
		goto L137
	} else {
		goto L192
	}
L192:
	;
	v926 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v468))))
	v929 = v926
	goto L137
L193:
	;
	v962 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v315)+74)))
	if v962 == int32(1) {
		goto L194
	} else {
		goto L195
	}
L194:
	;
	v965 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+15)))
	v966 = int32(1)
	v967 = v965 ^ v966
	*(*uint8)(unsafe.Add(mBase, uint32(v40)+6)) = uint8(v967)
	v969 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+14)))
	v971 = v969 ^ v966
	*(*uint8)(unsafe.Add(mBase, uint32(v40)+8)) = uint8(v971)
	goto L196
L195:
	;
	goto L196
L196:
	;
	v973 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v315)+75)))
	if v973 != int32(1) {
		goto L104
	} else {
		goto L197
	}
L197:
	;
	v976 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+15)))
	v977 = int32(1)
	v978 = v976 ^ v977
	*(*uint8)(unsafe.Add(mBase, uint32(v40)+7)) = uint8(v978)
	v980 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+14)))
	v982 = v980 ^ v977
	*(*uint8)(unsafe.Add(mBase, uint32(v40)+9)) = uint8(v982)
	goto L104
L198:
	;
	v1014 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40))))
	if v1014 != int32(1) {
		goto L199
	} else {
		goto L200
	}
L199:
	;
	v1051 = v285 + int32(1)
	v1052 = *(*int32)(unsafe.Add(mBase, uint32(v264)+4))
	if v1051 < v1052 {
		v285 = v1051
		goto L59
	} else {
		goto L216
	}
L200:
	;
	v1017 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+1)))
	if v1017 != int32(1) {
		goto L199
	} else {
		goto L201
	}
L201:
	;
	v1020 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+2)))
	if v1020 != int32(1) {
		goto L202
	} else {
		goto L203
	}
L202:
	;
	v1030 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+2)))
	if v1030 != int32(1) {
		goto L207
	} else {
		goto L208
	}
L203:
	;
	v1023 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+3)))
	if v1023 != int32(1) {
		goto L202
	} else {
		goto L204
	}
L204:
	;
	v1026 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v272))))
	if v1026 != 0 {
		goto L202
	} else {
		goto L205
	}
L205:
	;
	v1027 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+5)))
	if v1027 != int32(1) {
		goto L55
	} else {
		goto L206
	}
L206:
	;
	goto L202
L207:
	;
	v1040 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+2)))
	if v1040 != int32(1) {
		goto L199
	} else {
		goto L212
	}
L208:
	;
	v1033 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+3)))
	if v1033 != int32(1) {
		goto L207
	} else {
		goto L209
	}
L209:
	;
	v1036 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+6)))
	if v1036 != 0 {
		goto L207
	} else {
		goto L210
	}
L210:
	;
	v1037 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+7)))
	if v1037 != int32(1) {
		goto L55
	} else {
		goto L211
	}
L211:
	;
	goto L207
L212:
	;
	v1043 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+3)))
	if v1043 != int32(1) {
		goto L199
	} else {
		goto L213
	}
L213:
	;
	v1046 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+8)))
	if v1046 != 0 {
		goto L199
	} else {
		goto L214
	}
L214:
	;
	v1047 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+9)))
	if v1047 != int32(1) {
		goto L55
	} else {
		goto L215
	}
L215:
	;
	goto L199
L216:
	;
	goto L60
L217:
	;
	F_pfree(m, v1082)
	mBase = m.M
	v1084 = m.ExcPending
	if v1084 != 0 {
		goto L24
	} else {
		goto L220
	}
L218:
	;
	goto L219
L219:
	;
	v1087 = int32(4486928)
	v1088 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v1091 = *(*int32)(unsafe.Add(mBase, _consts[400]))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v1091
	v1094 = F_palloc(m, int32(10))
	mBase = m.M
	v1095 = m.ExcPending
	if v1095 != 0 {
		goto L24
	} else {
		goto L221
	}
L220:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+176)) = int32(0)
	goto L219
L221:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+176)) = v1094
	v1097 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v40)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(v1094)+8)) = uint16(v1097)
	v1099 = *(*int64)(unsafe.Add(mBase, uint32(v40)))
	*(*int64)(unsafe.Add(mBase, uint32(v1094))) = v1099
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v1088
	goto L17
L222:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43))) = v307
	F_errmsg_internal(m, int32(46623), v43)
	mBase = m.M
	v1141 = m.ExcPending
	if v1141 != 0 {
		goto L24
	} else {
		goto L223
	}
L223:
	;
	F_errfinish(m, int32(496016), int32(5867), int32(485470))
	mBase = m.M
	v1146 = m.ExcPending
	if v1146 != 0 {
		goto L24
	} else {
		goto L224
	}
L224:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L225:
	;
	v1196 = F_RelationGetReplicaIndex(m, l0)
	mBase = m.M
	v1197 = m.ExcPending
	if v1197 != 0 {
		goto L24
	} else {
		goto L241
	}
L226:
	;
	v1149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+122)))
	if v1149 == int32(0) {
		goto L7
	} else {
		goto L229
	}
L227:
	;
	goto L228
L228:
	;
	if l1 != int32(4) {
		v1195 = int32(0)
		goto L225
	} else {
		goto L237
	}
L229:
	;
	v1152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+124)))
	if v1152 == int32(0) {
		goto L6
	} else {
		goto L230
	}
L230:
	;
	v1156 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+126)))
	if v1156 != 0 {
		v1195 = int32(0)
		goto L225
	} else {
		goto L231
	}
L231:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1160 = m.ExcPending
	if v1160 != 0 {
		goto L24
	} else {
		goto L232
	}
L232:
	;
	F_errcode(m, int32(393348))
	mBase = m.M
	v1163 = m.ExcPending
	if v1163 != 0 {
		goto L24
	} else {
		goto L233
	}
L233:
	;
	v1164 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+32)) = v1164 + int32(4)
	F_errmsg(m, int32(701818), v31+int32(32))
	mBase = m.M
	v1172 = m.ExcPending
	if v1172 != 0 {
		goto L24
	} else {
		goto L234
	}
L234:
	;
	F_errdetail(m, int32(574308), int32(0))
	mBase = m.M
	v1176 = m.ExcPending
	if v1176 != 0 {
		goto L24
	} else {
		goto L235
	}
L235:
	;
	F_errfinish(m, int32(493175), int32(823), int32(10823))
	mBase = m.M
	v1181 = m.ExcPending
	if v1181 != 0 {
		goto L24
	} else {
		goto L236
	}
L236:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L237:
	;
	v1185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+123)))
	if v1185 == int32(0) {
		goto L5
	} else {
		goto L238
	}
L238:
	;
	v1188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+125)))
	if v1188 == int32(0) {
		goto L4
	} else {
		goto L239
	}
L239:
	;
	v1191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+127)))
	if v1191 == int32(0) {
		goto L3
	} else {
		goto L240
	}
L240:
	;
	v1195 = int32(1)
	goto L225
L241:
	;
	if v1196 != 0 {
		goto L8
	} else {
		goto L242
	}
L242:
	;
	v1198 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v1199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1198)+130)))
	if v1199 == int32(102) {
		goto L8
	} else {
		goto L243
	}
L243:
	;
	if l1 == int32(2) {
		goto L244
	} else {
		goto L245
	}
L244:
	;
	v1204 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+119)))
	if v1204 == int32(1) {
		goto L2
	} else {
		goto L247
	}
L245:
	;
	goto L246
L246:
	;
	if v1195 == int32(0) {
		goto L8
	} else {
		goto L248
	}
L247:
	;
	goto L246
L248:
	;
	v1209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+120)))
	if v1209 == int32(1) {
		goto L1
	} else {
		goto L249
	}
L249:
	;
	goto L8
L250:
	;
	F_errcode(m, int32(393348))
	mBase = m.M
	v1249 = m.ExcPending
	if v1249 != 0 {
		goto L24
	} else {
		goto L251
	}
L251:
	;
	v1250 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+64)) = v1250 + int32(4)
	F_errmsg(m, int32(701818), v31-int32(-64))
	mBase = m.M
	v1258 = m.ExcPending
	if v1258 != 0 {
		goto L24
	} else {
		goto L252
	}
L252:
	;
	F_errdetail(m, int32(553142), int32(0))
	mBase = m.M
	v1262 = m.ExcPending
	if v1262 != 0 {
		goto L24
	} else {
		goto L253
	}
L253:
	;
	F_errfinish(m, int32(493175), int32(811), int32(10823))
	mBase = m.M
	v1267 = m.ExcPending
	if v1267 != 0 {
		goto L24
	} else {
		goto L254
	}
L254:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L255:
	;
	F_errcode(m, int32(393348))
	mBase = m.M
	v1274 = m.ExcPending
	if v1274 != 0 {
		goto L24
	} else {
		goto L256
	}
L256:
	;
	v1275 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+48)) = v1275 + int32(4)
	F_errmsg(m, int32(701818), v31+int32(48))
	mBase = m.M
	v1283 = m.ExcPending
	if v1283 != 0 {
		goto L24
	} else {
		goto L257
	}
L257:
	;
	F_errdetail(m, int32(553069), int32(0))
	mBase = m.M
	v1287 = m.ExcPending
	if v1287 != 0 {
		goto L24
	} else {
		goto L258
	}
L258:
	;
	F_errfinish(m, int32(493175), int32(817), int32(10823))
	mBase = m.M
	v1292 = m.ExcPending
	if v1292 != 0 {
		goto L24
	} else {
		goto L259
	}
L259:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L260:
	;
	F_errcode(m, int32(393348))
	mBase = m.M
	v1299 = m.ExcPending
	if v1299 != 0 {
		goto L24
	} else {
		goto L261
	}
L261:
	;
	v1300 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+112)) = v1300 + int32(4)
	F_errmsg(m, int32(700985), v31+int32(112))
	mBase = m.M
	v1308 = m.ExcPending
	if v1308 != 0 {
		goto L24
	} else {
		goto L262
	}
L262:
	;
	F_errdetail(m, int32(553142), int32(0))
	mBase = m.M
	v1312 = m.ExcPending
	if v1312 != 0 {
		goto L24
	} else {
		goto L263
	}
L263:
	;
	F_errfinish(m, int32(493175), int32(829), int32(10823))
	mBase = m.M
	v1317 = m.ExcPending
	if v1317 != 0 {
		goto L24
	} else {
		goto L264
	}
L264:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L265:
	;
	F_errcode(m, int32(393348))
	mBase = m.M
	v1324 = m.ExcPending
	if v1324 != 0 {
		goto L24
	} else {
		goto L266
	}
L266:
	;
	v1325 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+96)) = v1325 + int32(4)
	F_errmsg(m, int32(700985), v31+int32(96))
	mBase = m.M
	v1333 = m.ExcPending
	if v1333 != 0 {
		goto L24
	} else {
		goto L267
	}
L267:
	;
	F_errdetail(m, int32(553069), int32(0))
	mBase = m.M
	v1337 = m.ExcPending
	if v1337 != 0 {
		goto L24
	} else {
		goto L268
	}
L268:
	;
	F_errfinish(m, int32(493175), int32(835), int32(10823))
	mBase = m.M
	v1342 = m.ExcPending
	if v1342 != 0 {
		goto L24
	} else {
		goto L269
	}
L269:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L270:
	;
	F_errcode(m, int32(393348))
	mBase = m.M
	v1349 = m.ExcPending
	if v1349 != 0 {
		goto L24
	} else {
		goto L271
	}
L271:
	;
	v1350 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+80)) = v1350 + int32(4)
	F_errmsg(m, int32(700985), v31+int32(80))
	mBase = m.M
	v1358 = m.ExcPending
	if v1358 != 0 {
		goto L24
	} else {
		goto L272
	}
L272:
	;
	F_errdetail(m, int32(574308), int32(0))
	mBase = m.M
	v1362 = m.ExcPending
	if v1362 != 0 {
		goto L24
	} else {
		goto L273
	}
L273:
	;
	F_errfinish(m, int32(493175), int32(841), int32(10823))
	mBase = m.M
	v1367 = m.ExcPending
	if v1367 != 0 {
		goto L24
	} else {
		goto L274
	}
L274:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L275:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v1374 = m.ExcPending
	if v1374 != 0 {
		goto L24
	} else {
		goto L276
	}
L276:
	;
	v1375 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v31))) = v1375 + int32(4)
	F_errmsg(m, int32(159664), v31)
	mBase = m.M
	v1381 = m.ExcPending
	if v1381 != 0 {
		goto L24
	} else {
		goto L277
	}
L277:
	;
	F_errhint(m, int32(640794), int32(0))
	mBase = m.M
	v1385 = m.ExcPending
	if v1385 != 0 {
		goto L24
	} else {
		goto L278
	}
L278:
	;
	F_errfinish(m, int32(493175), int32(861), int32(10823))
	mBase = m.M
	v1390 = m.ExcPending
	if v1390 != 0 {
		goto L24
	} else {
		goto L279
	}
L279:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L280:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v1397 = m.ExcPending
	if v1397 != 0 {
		goto L24
	} else {
		goto L281
	}
L281:
	;
	v1398 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+16)) = v1398 + int32(4)
	F_errmsg(m, int32(159281), v31+int32(16))
	mBase = m.M
	v1406 = m.ExcPending
	if v1406 != 0 {
		goto L24
	} else {
		goto L282
	}
L282:
	;
	F_errhint(m, int32(640719), int32(0))
	mBase = m.M
	v1410 = m.ExcPending
	if v1410 != 0 {
		goto L24
	} else {
		goto L283
	}
L283:
	;
	F_errfinish(m, int32(493175), int32(867), int32(10823))
	mBase = m.M
	v1415 = m.ExcPending
	if v1415 != 0 {
		goto L24
	} else {
		goto L284
	}
L284:
	;
	base.Wasm_trap_unreachable()
	for {
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
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v76 int32
	_ = v76
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
	goto L10
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
	v59 = v29
	goto L15
L10:
	;
	v33 = int32(*(*int8)(unsafe.Add(mBase, uint32(v26))))
	goto L12
L11:
	;
	v81 = v48
	v84 = int32(0)
	v85 = v23 + v24
	goto L1
L12:
	;
	if base.B2i32(v33 == int32(32))|base.B2i32(base.Ui32((v33-int32(9))&int32(255)) < base.Ui32(int32(5))) == int32(0) {
		goto L9
	} else {
		goto L13
	}
L13:
	;
	v45 = int32(1)
	v48 = v26 + v45
	if v45 < v29 {
		v26 = v48
		v29 = v29 - v45
		v30 = v30 + v45
		goto L10
	} else {
		goto L14
	}
L14:
	;
	goto L11
L15:
	;
	v64 = int32(*(*int8)(unsafe.Add(mBase, uint32(v59+(v26-int32(1))))))
	goto L17
L16:
	;
	v81 = v26
	v84 = int32(0)
	v85 = v30
	goto L1
L17:
	;
	if base.B2i32(v64 == int32(32))|base.B2i32(base.Ui32((v64-int32(9))&int32(255)) < base.Ui32(int32(5))) == int32(0) {
		v81 = v26
		v84 = v59
		v85 = v30
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v76 = int32(1)
	if v76 < v59 {
		v59 = v59 - v76
		goto L15
	} else {
		goto L19
	}
L19:
	;
	goto L16
}
func F_ComputeXidHorizons(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
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
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int64
	_ = v46
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v78 int32
	_ = v78
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v175 int32
	_ = v175
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v309 int32
	_ = v309
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v388 int32
	_ = v388
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v411 int32
	_ = v411
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v441 int32
	_ = v441
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v463 int32
	_ = v463
	var v466 int32
	_ = v466
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v504 int32
	_ = v504
	var v506 int32
	_ = v506
	var v509 int32
	_ = v509
	var v511 int32
	_ = v511
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v534 int64
	_ = v534
	var v535 int32
	_ = v535
	var v539 int64
	_ = v539
	var v542 int64
	_ = v542
	var v543 int32
	_ = v543
	var v547 int64
	_ = v547
	var v550 int64
	_ = v550
	var v554 int64
	_ = v554
	var v557 int64
	_ = v557
	var v558 int32
	_ = v558
	var v562 int64
	_ = v562
	var v564 int32
	_ = v564
	var v565 int64
	_ = v565
	var v568 int64
	_ = v568
	var v570 int64
	_ = v570
	var v572 int64
	_ = v572
	var v574 int32
	_ = v574
	var v575 int64
	_ = v575
	var v578 int64
	_ = v578
	var v580 int64
	_ = v580
	var v582 int64
	_ = v582
	var v584 int32
	_ = v584
	var v585 int64
	_ = v585
	var v590 int64
	_ = v590
	var v592 int64
	_ = v592
	var v594 int64
	_ = v594
	var v598 int32
	_ = v598
	v18 = *(*int32)(unsafe.Add(mBase, _consts[193]))
	v21 = int32(*(*uint8)(unsafe.Add(mBase, _consts[182])))
	if v21 == int32(1) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v33 = *(*int32)(unsafe.Add(mBase, _consts[153]))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = int32(0)
	v38 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v42 = F_LWLockAcquire(m, v38+int32(512), int32(1))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	v26 = *(*int32)(unsafe.Add(mBase, _consts[175]))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+316))
	v29 = base.B2i32(v27 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _consts[182])) = uint8(v29)
	v31 = v29
	goto L4
L3:
	;
	v31 = int32(0)
	goto L4
L4:
	;
	goto L1
L5:
	;
	return
L6:
	;
	v45 = *(*int32)(unsafe.Add(mBase, _consts[139]))
	v46 = *(*int64)(unsafe.Add(mBase, uint32(v45)+48))
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v46
	v48 = int32(3)
	v51 = base.I32_wrap_i64(v46) + int32(1)
	if base.Ui32(v51) <= base.Ui32(v48) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v54 = v48
	goto L9
L8:
	;
	v54 = v51
	goto L9
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v54
	v59 = *(*int32)(unsafe.Add(mBase, _consts[185]))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)+36))
	if v60 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v61 = v60
	goto L12
L11:
	;
	v61 = v54
	goto L12
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v61
	v64 = *(*int32)(unsafe.Add(mBase, _consts[193]))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+28))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v65
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v64)+32))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v67
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	if int32(0) < v69 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v78 = int32(0)
	goto L16
L14:
	;
	goto L15
L15:
	;
	if v31 == int32(0) {
		goto L68
	} else {
		goto L69
	}
L16:
	;
	v92 = *(*int32)(unsafe.Add(mBase, _consts[153]))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v92)+12))
	v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v93+v78))))
	v97 = v78 << (uint(int32(2)) % 32)
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v18+int32(36)+v97)))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v97+v34)))
	v103 = *(*int32)(unsafe.Add(mBase, _consts[194]))
	v106 = v103 + v99*int32(640)
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v106)+40))
	if v107 != 0 {
		goto L20
	} else {
		goto L21
	}
L17:
	;
	goto L15
L18:
	;
	v195 = v78 + int32(1)
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	if v195 < v196 {
		v78 = v195
		goto L16
	} else {
		goto L66
	}
L19:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v126 != 0 {
		goto L34
	} else {
		goto L35
	}
L20:
	;
	if v101 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	goto L22
L22:
	;
	if v101 == int32(0) {
		goto L18
	} else {
		goto L33
	}
L23:
	;
	v125 = v107
	goto L19
L24:
	;
	goto L25
L25:
	;
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v101))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v107)) == int32(0) {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	if v121 != 0 {
		goto L30
	} else {
		goto L31
	}
L27:
	;
	v121 = base.B2i32(base.Ui32(v107) < base.Ui32(v101))
	goto L26
L28:
	;
	goto L29
L29:
	;
	v121 = int32(base.Ui32(v107-v101) >> (uint(int32(31)) % 32))
	goto L26
L30:
	;
	v122 = v107
	goto L32
L31:
	;
	v122 = v101
	goto L32
L32:
	;
	v125 = v122
	goto L19
L33:
	;
	v125 = v101
	goto L19
L34:
	;
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v125))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v126)) == int32(0) {
		goto L38
	} else {
		goto L39
	}
L35:
	;
	v140 = v125
	goto L36
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v140
	if v95&int32(18) != 0 {
		goto L18
	} else {
		goto L44
	}
L37:
	;
	if v138 != 0 {
		goto L41
	} else {
		goto L42
	}
L38:
	;
	v138 = base.B2i32(base.Ui32(v126) < base.Ui32(v125))
	goto L37
L39:
	;
	goto L40
L40:
	;
	v138 = int32(base.Ui32(v126-v125) >> (uint(int32(31)) % 32))
	goto L37
L41:
	;
	v139 = v126
	goto L43
L42:
	;
	v139 = v125
	goto L43
L43:
	;
	v140 = v139
	goto L36
L44:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v144 != 0 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v125))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v144)) == int32(0) {
		goto L49
	} else {
		goto L50
	}
L46:
	;
	v158 = v125
	goto L47
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v158
	v161 = *(*int32)(unsafe.Add(mBase, _consts[4]))
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v106)+60))
	v164 = int32(0)
	if (base.B2i32(v161 == v162)|base.B2i32(v161 == v164)|int32(base.Ui32(v95)>>(uint(int32(5))%32))|v31)&int32(1) == v164 {
		goto L18
	} else {
		goto L55
	}
L48:
	;
	if v156 != 0 {
		goto L52
	} else {
		goto L53
	}
L49:
	;
	v156 = base.B2i32(base.Ui32(v144) < base.Ui32(v125))
	goto L48
L50:
	;
	goto L51
L51:
	;
	v156 = int32(base.Ui32(v144-v125) >> (uint(int32(31)) % 32))
	goto L48
L52:
	;
	v157 = v144
	goto L54
L53:
	;
	v157 = v125
	goto L54
L54:
	;
	v158 = v157
	goto L47
L55:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v175 != 0 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v125))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v175)) == int32(0) {
		goto L60
	} else {
		goto L61
	}
L57:
	;
	v189 = v125
	goto L58
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v189
	goto L18
L59:
	;
	if v187 != 0 {
		goto L63
	} else {
		goto L64
	}
L60:
	;
	v187 = base.B2i32(base.Ui32(v175) < base.Ui32(v125))
	goto L59
L61:
	;
	goto L62
L62:
	;
	v187 = int32(base.Ui32(v175-v125) >> (uint(int32(31)) % 32))
	goto L59
L63:
	;
	v188 = v175
	goto L65
L64:
	;
	v188 = v125
	goto L65
L65:
	;
	v189 = v188
	goto L58
L66:
	;
	goto L17
L67:
	;
	v365 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v366 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v366 == int32(0) {
		goto L128
	} else {
		goto L129
	}
L68:
	;
	v217 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	F_LWLockRelease(m, v217+int32(512))
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L5
	} else {
		goto L71
	}
L69:
	;
	goto L70
L70:
	;
	v224 = *(*int32)(unsafe.Add(mBase, _consts[193]))
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v224)+20))
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v224)+16))
	if v226 < v225 {
		goto L74
	} else {
		goto L75
	}
L71:
	;
	goto L67
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v293
	v309 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v309 != 0 {
		goto L99
	} else {
		goto L100
	}
L73:
	;
	if v261 == int32(0) {
		goto L87
	} else {
		goto L88
	}
L74:
	;
	v229 = *(*int32)(unsafe.Add(mBase, _consts[770]))
	v231 = v226
	goto L78
L75:
	;
	goto L76
L76:
	;
	v270 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	F_LWLockRelease(m, v270+int32(512))
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L5
	} else {
		goto L86
	}
L77:
	;
	v263 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	F_LWLockRelease(m, v263+int32(512))
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L5
	} else {
		goto L84
	}
L78:
	;
	v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v231+v229))))
	if v247 == int32(1) {
		goto L80
	} else {
		goto L81
	}
L79:
	;
	v261 = int32(0)
	goto L77
L80:
	;
	v251 = *(*int32)(unsafe.Add(mBase, _consts[771]))
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v251+v231<<(uint(int32(2))%32))))
	v261 = v255
	goto L77
L81:
	;
	goto L82
L82:
	;
	v257 = v231 + int32(1)
	if v257 != v225 {
		v231 = v257
		goto L78
	} else {
		goto L83
	}
L83:
	;
	goto L79
L84:
	;
	v268 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v268 != 0 {
		goto L73
	} else {
		goto L85
	}
L85:
	;
	v293 = v261
	v294 = v261
	goto L72
L86:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v293 = v275
	v294 = int32(0)
	goto L72
L87:
	;
	v293 = v268
	v294 = int32(0)
	goto L72
L88:
	;
	goto L89
L89:
	;
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v261))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v268)) == int32(0) {
		goto L91
	} else {
		goto L92
	}
L90:
	;
	if v290 != 0 {
		goto L94
	} else {
		goto L95
	}
L91:
	;
	v290 = base.B2i32(base.Ui32(v268) < base.Ui32(v261))
	goto L90
L92:
	;
	goto L93
L93:
	;
	v290 = int32(base.Ui32(v268-v261) >> (uint(int32(31)) % 32))
	goto L90
L94:
	;
	v291 = v268
	goto L96
L95:
	;
	v291 = v261
	goto L96
L96:
	;
	v293 = v291
	v294 = v261
	goto L72
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v347
	goto L67
L98:
	;
	v345 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v347 = v345
	goto L97
L99:
	;
	if v294 == int32(0) {
		goto L98
	} else {
		goto L102
	}
L100:
	;
	v325 = v294
	goto L101
L101:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v325
	v327 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v327 == int32(0) {
		v347 = v294
		goto L97
	} else {
		goto L110
	}
L102:
	;
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v294))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v309)) == int32(0) {
		goto L104
	} else {
		goto L105
	}
L103:
	;
	if v323 != 0 {
		goto L107
	} else {
		goto L108
	}
L104:
	;
	v323 = base.B2i32(base.Ui32(v309) < base.Ui32(v294))
	goto L103
L105:
	;
	goto L106
L106:
	;
	v323 = int32(base.Ui32(v309-v294) >> (uint(int32(31)) % 32))
	goto L103
L107:
	;
	v324 = v309
	goto L109
L108:
	;
	v324 = v294
	goto L109
L109:
	;
	v325 = v324
	goto L101
L110:
	;
	if v294 == int32(0) {
		goto L111
	} else {
		goto L112
	}
L111:
	;
	v347 = v327
	goto L97
L112:
	;
	goto L113
L113:
	;
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v294))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v327)) == int32(0) {
		goto L115
	} else {
		goto L116
	}
L114:
	;
	if v343 != 0 {
		goto L118
	} else {
		goto L119
	}
L115:
	;
	v343 = base.B2i32(base.Ui32(v327) < base.Ui32(v294))
	goto L114
L116:
	;
	goto L117
L117:
	;
	v343 = int32(base.Ui32(v327-v294) >> (uint(int32(31)) % 32))
	goto L114
L118:
	;
	v344 = v327
	goto L120
L119:
	;
	v344 = v294
	goto L120
L120:
	;
	v347 = v344
	goto L97
L121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v461
	v466 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v466 != 0 {
		goto L175
	} else {
		goto L176
	}
L122:
	;
	if v439 == int32(0) {
		goto L161
	} else {
		goto L162
	}
L123:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v419
	v461 = v417
	v463 = v419
	goto L121
L124:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v438
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v441
	if v438 != 0 {
		goto L122
	} else {
		goto L160
	}
L125:
	;
	if v418 == int32(0) {
		goto L123
	} else {
		goto L152
	}
L126:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v366
	v415 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v416 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v417 = v416
	v418 = v415
	v419 = v366
	goto L125
L127:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v386
	v388 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v388 == int32(0) {
		v407 = v385
		v408 = v386
		goto L139
	} else {
		goto L140
	}
L128:
	;
	v385 = v365
	v386 = v365
	goto L127
L129:
	;
	goto L130
L130:
	;
	if v365 == int32(0) {
		goto L126
	} else {
		goto L131
	}
L131:
	;
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v365))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v366)) == int32(0) {
		goto L133
	} else {
		goto L134
	}
L132:
	;
	if v382 != 0 {
		goto L136
	} else {
		goto L137
	}
L133:
	;
	v382 = base.B2i32(base.Ui32(v366) < base.Ui32(v365))
	goto L132
L134:
	;
	goto L135
L135:
	;
	v382 = int32(base.Ui32(v366-v365) >> (uint(int32(31)) % 32))
	goto L132
L136:
	;
	v383 = v366
	goto L138
L137:
	;
	v383 = v365
	goto L138
L138:
	;
	v384 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v385 = v384
	v386 = v383
	goto L127
L139:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v408
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v407
	v411 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v408 == int32(0) {
		v438 = v407
		v439 = v411
		v441 = v411
		goto L124
	} else {
		goto L151
	}
L140:
	;
	if v385 == int32(0) {
		goto L141
	} else {
		goto L142
	}
L141:
	;
	v407 = v388
	v408 = v386
	goto L139
L142:
	;
	goto L143
L143:
	;
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v385))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v388)) == int32(0) {
		goto L145
	} else {
		goto L146
	}
L144:
	;
	if v404 != 0 {
		goto L148
	} else {
		goto L149
	}
L145:
	;
	v404 = base.B2i32(base.Ui32(v388) < base.Ui32(v385))
	goto L144
L146:
	;
	goto L147
L147:
	;
	v404 = int32(base.Ui32(v388-v385) >> (uint(int32(31)) % 32))
	goto L144
L148:
	;
	v405 = v388
	goto L150
L149:
	;
	v405 = v385
	goto L150
L150:
	;
	v406 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v407 = v405
	v408 = v406
	goto L139
L151:
	;
	v417 = v407
	v418 = v411
	v419 = v408
	goto L125
L152:
	;
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v418))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v419)) == int32(0) {
		goto L154
	} else {
		goto L155
	}
L153:
	;
	if v434 != 0 {
		goto L157
	} else {
		goto L158
	}
L154:
	;
	v434 = base.B2i32(base.Ui32(v419) < base.Ui32(v418))
	goto L153
L155:
	;
	goto L156
L156:
	;
	v434 = int32(base.Ui32(v419-v418) >> (uint(int32(31)) % 32))
	goto L153
L157:
	;
	v435 = v419
	goto L159
L158:
	;
	v435 = v418
	goto L159
L159:
	;
	v436 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v437 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v438 = v437
	v439 = v436
	v441 = v435
	goto L124
L160:
	;
	v461 = v439
	v463 = v441
	goto L121
L161:
	;
	v461 = v438
	v463 = v441
	goto L121
L162:
	;
	goto L163
L163:
	;
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v439))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v438)) == int32(0) {
		goto L165
	} else {
		goto L166
	}
L164:
	;
	if v458 != 0 {
		goto L168
	} else {
		goto L169
	}
L165:
	;
	v458 = base.B2i32(base.Ui32(v438) < base.Ui32(v439))
	goto L164
L166:
	;
	goto L167
L167:
	;
	v458 = int32(base.Ui32(v438-v439) >> (uint(int32(31)) % 32))
	goto L164
L168:
	;
	v459 = v438
	goto L170
L169:
	;
	v459 = v439
	goto L170
L170:
	;
	v460 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v461 = v459
	v463 = v460
	goto L121
L171:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v530
	v534 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	v535 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v539 = v534 + base.I64_extend_i32_s(v535-base.I32_wrap_i64(v534))
	*(*int64)(unsafe.Add(mBase, _consts[776])) = v539
	v542 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	v543 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v547 = v542 + base.I64_extend_i32_s(v543-base.I32_wrap_i64(v542))
	*(*int64)(unsafe.Add(mBase, _consts[777])) = v547
	v550 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	v554 = v550 + base.I64_extend_i32_s(v531-base.I32_wrap_i64(v550))
	*(*int64)(unsafe.Add(mBase, _consts[778])) = v554
	v557 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	v558 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v562 = v557 + base.I64_extend_i32_s(v558-base.I32_wrap_i64(v557))
	*(*int64)(unsafe.Add(mBase, _consts[779])) = v562
	v564 = int32(4403592)
	v565 = *(*int64)(unsafe.Add(mBase, _consts[780]))
	if base.Ui64(v565) < base.Ui64(v539) {
		goto L210
	} else {
		goto L211
	}
L172:
	;
	if v511 == int32(0) {
		goto L200
	} else {
		goto L201
	}
L173:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v461
	v506 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v461 == int32(0) {
		v530 = v506
		v531 = v506
		goto L171
	} else {
		goto L199
	}
L174:
	;
	if v487 != 0 {
		goto L189
	} else {
		goto L190
	}
L175:
	;
	if v463 != 0 {
		goto L178
	} else {
		goto L179
	}
L176:
	;
	goto L177
L177:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v463
	if v463 == int32(0) {
		goto L173
	} else {
		goto L188
	}
L178:
	;
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v463))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v466)) == int32(0) {
		goto L182
	} else {
		goto L183
	}
L179:
	;
	v481 = v461
	v482 = v466
	goto L180
L180:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v482
	v487 = v481
	v488 = v482
	goto L174
L181:
	;
	if v478 != 0 {
		goto L185
	} else {
		goto L186
	}
L182:
	;
	v478 = base.B2i32(base.Ui32(v466) < base.Ui32(v463))
	goto L181
L183:
	;
	goto L184
L184:
	;
	v478 = int32(base.Ui32(v466-v463) >> (uint(int32(31)) % 32))
	goto L181
L185:
	;
	v479 = v466
	goto L187
L186:
	;
	v479 = v463
	goto L187
L187:
	;
	v480 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v481 = v480
	v482 = v479
	goto L180
L188:
	;
	v487 = v461
	v488 = v463
	goto L174
L189:
	;
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v487))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v488)) == int32(0) {
		goto L193
	} else {
		goto L194
	}
L190:
	;
	v502 = v488
	goto L191
L191:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v502
	v504 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v509 = v502
	v511 = v504
	goto L172
L192:
	;
	if v500 != 0 {
		goto L196
	} else {
		goto L197
	}
L193:
	;
	v500 = base.B2i32(base.Ui32(v488) < base.Ui32(v487))
	goto L192
L194:
	;
	goto L195
L195:
	;
	v500 = int32(base.Ui32(v488-v487) >> (uint(int32(31)) % 32))
	goto L192
L196:
	;
	v501 = v488
	goto L198
L197:
	;
	v501 = v487
	goto L198
L198:
	;
	v502 = v501
	goto L191
L199:
	;
	v509 = v461
	v511 = v506
	goto L172
L200:
	;
	v530 = v509
	v531 = int32(0)
	goto L171
L201:
	;
	goto L202
L202:
	;
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v511))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v509)) == int32(0) {
		goto L204
	} else {
		goto L205
	}
L203:
	;
	if v526 != 0 {
		goto L207
	} else {
		goto L208
	}
L204:
	;
	v526 = base.B2i32(base.Ui32(v509) < base.Ui32(v511))
	goto L203
L205:
	;
	goto L206
L206:
	;
	v526 = int32(base.Ui32(v509-v511) >> (uint(int32(31)) % 32))
	goto L203
L207:
	;
	v527 = v509
	goto L209
L208:
	;
	v527 = v511
	goto L209
L209:
	;
	v528 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v530 = v527
	v531 = v528
	goto L171
L210:
	;
	v568 = v539
	goto L212
L211:
	;
	v568 = v565
	goto L212
L212:
	;
	if base.I32_wrap_i64(v565) != 0 {
		goto L213
	} else {
		goto L214
	}
L213:
	;
	v570 = v568
	goto L215
L214:
	;
	v570 = v539
	goto L215
L215:
	;
	if base.I32_wrap_i64(v539) != 0 {
		goto L216
	} else {
		goto L217
	}
L216:
	;
	v572 = v570
	goto L218
L217:
	;
	v572 = v565
	goto L218
L218:
	;
	*(*int64)(unsafe.Add(mBase, _consts[780])) = v572
	v574 = int32(4403608)
	v575 = *(*int64)(unsafe.Add(mBase, _consts[781]))
	if base.Ui64(v575) < base.Ui64(v547) {
		goto L219
	} else {
		goto L220
	}
L219:
	;
	v578 = v547
	goto L221
L220:
	;
	v578 = v575
	goto L221
L221:
	;
	if base.I32_wrap_i64(v575) != 0 {
		goto L222
	} else {
		goto L223
	}
L222:
	;
	v580 = v578
	goto L224
L223:
	;
	v580 = v547
	goto L224
L224:
	;
	if base.I32_wrap_i64(v547) != 0 {
		goto L225
	} else {
		goto L226
	}
L225:
	;
	v582 = v580
	goto L227
L226:
	;
	v582 = v575
	goto L227
L227:
	;
	*(*int64)(unsafe.Add(mBase, _consts[781])) = v582
	v584 = int32(4403624)
	v585 = *(*int64)(unsafe.Add(mBase, _consts[782]))
	*(*int64)(unsafe.Add(mBase, _consts[783])) = v562
	if base.Ui64(v585) < base.Ui64(v554) {
		goto L228
	} else {
		goto L229
	}
L228:
	;
	v590 = v554
	goto L230
L229:
	;
	v590 = v585
	goto L230
L230:
	;
	if base.I32_wrap_i64(v585) != 0 {
		goto L231
	} else {
		goto L232
	}
L231:
	;
	v592 = v590
	goto L233
L232:
	;
	v592 = v554
	goto L233
L233:
	;
	if base.I32_wrap_i64(v554) != 0 {
		goto L234
	} else {
		goto L235
	}
L234:
	;
	v594 = v592
	goto L236
L235:
	;
	v594 = v585
	goto L236
L236:
	;
	*(*int64)(unsafe.Add(mBase, _consts[782])) = v594
	v598 = *(*int32)(unsafe.Add(mBase, _consts[784]))
	*(*int32)(unsafe.Add(mBase, _consts[785])) = v598
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
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v56 int32
	_ = v56
	v7 = *(*int32)(unsafe.Add(mBase, _consts[798]))
	if v7 != 0 {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
		*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(1)
		if v8 != 0 {
			F_s_lock(m, v7, int32(495856), int32(238), int32(236597))
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return
			} else {
				v17 = *(*int32)(unsafe.Add(mBase, _consts[153]))
				v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
				v20 = *(*int32)(unsafe.Add(mBase, _consts[743]))
				v23 = v18 + v20*int32(640)
				v25 = v23 + int32(84)
				v26 = *(*int32)(unsafe.Add(mBase, uint32(v23)+88))
				if v26 == int32(0) {
					v29 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
					if v29 != 0 {
						v35 = v29
						*(*int32)(unsafe.Add(mBase, uint32(v18+v26*int32(640))+84)) = v35
						v40 = v26
						v41 = v35
						if v41 == int32(-1) {
							*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v40
						} else {
							v46 = *(*int32)(unsafe.Add(mBase, _consts[153]))
							v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
							*(*int32)(unsafe.Add(mBase, uint32(v47+v41*int32(640))+88)) = v40
						}
						*(*int64)(unsafe.Add(mBase, uint32(v25))) = int64(0)
					} else {
					}
				} else {
					v30 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
					if v26 != int32(-1) {
						v35 = v30
						*(*int32)(unsafe.Add(mBase, uint32(v18+v26*int32(640))+84)) = v35
						v40 = v26
						v41 = v35
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v30
						v34 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
						v40 = v34
						v41 = v30
					}
					if v41 == int32(-1) {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v40
					} else {
						v46 = *(*int32)(unsafe.Add(mBase, _consts[153]))
						v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
						*(*int32)(unsafe.Add(mBase, uint32(v47+v41*int32(640))+88)) = v40
					}
					*(*int64)(unsafe.Add(mBase, uint32(v25))) = int64(0)
				}
				v56 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v7))) = v56
				*(*int32)(unsafe.Add(mBase, _consts[798])) = v56
				return
			}
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, _consts[153]))
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
			v20 = *(*int32)(unsafe.Add(mBase, _consts[743]))
			v23 = v18 + v20*int32(640)
			v25 = v23 + int32(84)
			v26 = *(*int32)(unsafe.Add(mBase, uint32(v23)+88))
			if v26 == int32(0) {
				v29 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
				if v29 != 0 {
					v35 = v29
					*(*int32)(unsafe.Add(mBase, uint32(v18+v26*int32(640))+84)) = v35
					v40 = v26
					v41 = v35
					if v41 == int32(-1) {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v40
					} else {
						v46 = *(*int32)(unsafe.Add(mBase, _consts[153]))
						v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
						*(*int32)(unsafe.Add(mBase, uint32(v47+v41*int32(640))+88)) = v40
					}
					*(*int64)(unsafe.Add(mBase, uint32(v25))) = int64(0)
				} else {
				}
			} else {
				v30 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
				if v26 != int32(-1) {
					v35 = v30
					*(*int32)(unsafe.Add(mBase, uint32(v18+v26*int32(640))+84)) = v35
					v40 = v26
					v41 = v35
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v30
					v34 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
					v40 = v34
					v41 = v30
				}
				if v41 == int32(-1) {
					*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v40
				} else {
					v46 = *(*int32)(unsafe.Add(mBase, _consts[153]))
					v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
					*(*int32)(unsafe.Add(mBase, uint32(v47+v41*int32(640))+88)) = v40
				}
				*(*int64)(unsafe.Add(mBase, uint32(v25))) = int64(0)
			}
			v56 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v7))) = v56
			*(*int32)(unsafe.Add(mBase, _consts[798])) = v56
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
	v62 = *(*int32)(unsafe.Add(mBase, _consts[1]))
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
	var v31 int32
	_ = v31
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
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v144 int32
	_ = v144
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v188 int32
	_ = v188
	v14 = m.G0
	v16 = v14 - int32(48)
	m.G0 = v16
	v19 = *(*int32)(unsafe.Add(mBase, _consts[193]))
	v31 = int32(0)
	goto L1
L1:
	;
	v36 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	if v36 != 0 {
		goto L3
	} else {
		goto L4
	}
L2:
	;
	m.G0 = v16 + int32(48)
	return v188
L3:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	v41 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v41
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v41
	v46 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v50 = F_LWLockAcquire(m, v46+int32(512), int32(1))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L6
	} else {
		goto L8
	}
L6:
	;
	return int32(0)
L7:
	;
	goto L5
L8:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	if v52 <= int32(0) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	goto L2
L10:
	;
	v57 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	F_LWLockRelease(m, v57+int32(512))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L6
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	v62 = int32(0)
	v64 = *(*int32)(unsafe.Add(mBase, _consts[194]))
	v70 = v62
	v71 = v62
	v72 = v62
	goto L14
L13:
	;
	v188 = int32(0)
	goto L9
L14:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v19+int32(36)+v70<<(uint(int32(2))%32))))
	v86 = v64 + v83*int32(640)
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v86)+60))
	if v87 != l0 {
		v123 = v71
		v124 = v72
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v131 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	F_LWLockRelease(m, v131+int32(512))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L6
	} else {
		goto L25
	}
L16:
	;
	v127 = v70 + int32(1)
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	if v127 < v128 {
		v70 = v127
		v71 = v123
		v72 = v124
		goto L14
	} else {
		goto L24
	}
L17:
	;
	v90 = *(*int32)(unsafe.Add(mBase, _consts[185]))
	if v86 == v90 {
		v123 = v71
		v124 = v72
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v86)+44))
	if v92 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v95 = int32(1)
	v96 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v96 + v95
	v123 = v95
	v124 = v72
	goto L16
L20:
	;
	goto L21
L21:
	;
	v101 = *(*int32)(unsafe.Add(mBase, _consts[153]))
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v101)+12))
	v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102+v70))))
	v105 = int32(1)
	v106 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v106 + v105
	if v104&v105 == int32(0) {
		v123 = v105
		v124 = v72
		goto L16
	} else {
		goto L22
	}
L22:
	;
	if int32(9) < v72 {
		v123 = v105
		v124 = v72
		goto L16
	} else {
		goto L23
	}
L23:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v86)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v16+v72<<(uint(int32(2))%32)))) = v119
	v123 = v105
	v124 = v72 + int32(1)
	goto L16
L24:
	;
	goto L15
L25:
	;
	if v123 == int32(0) {
		v188 = v123
		goto L9
	} else {
		goto L26
	}
L26:
	;
	v138 = int32(0)
	if v138 < v124 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v144 = v138
	goto L30
L28:
	;
	goto L29
L29:
	;
	F_pg_usleep(m, int32(100000))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L6
	} else {
		goto L34
	}
L30:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v16+v144<<(uint(int32(2))%32))))
	v159 = F_kill(m, v157, int32(15))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L6
	} else {
		goto L32
	}
L31:
	;
	goto L29
L32:
	;
	v162 = v144 + int32(1)
	if v162 != v124 {
		v144 = v162
		goto L30
	} else {
		goto L33
	}
L33:
	;
	goto L31
L34:
	;
	v181 = v31 + int32(1)
	if v181 != int32(50) {
		v31 = v181
		goto L1
	} else {
		goto L35
	}
L35:
	;
	v188 = v123
	goto L9
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
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
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
	F_ScanKeyInit(m, v11+int32(32), int32(1), int32(3), int32(184), l0)
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
	F_systable_endscan(m, v66)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L4
	} else {
		goto L20
	}
L10:
	;
	v66 = F_systable_beginscan(m, v58, int32(2675), int32(1), int32(0), int32(3), v11+int32(32))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L4
	} else {
		goto L11
	}
L11:
	;
	v68 = F_systable_getnext(m, v66)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L4
	} else {
		goto L12
	}
L12:
	;
	if v68 == int32(0) {
		v89 = v5
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
	F_CatalogTupleDelete(m, v58, v68+int32(4))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L4
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v58)+52))
	v85 = F_heap_modify_tuple(m, v68, v78, v11+int32(16), v11+int32(12), v11+int32(8))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L4
	} else {
		goto L18
	}
L17:
	;
	v89 = v5
	goto L9
L18:
	;
	F_CatalogTupleUpdate(m, v58, v68+int32(4), v85)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L4
	} else {
		goto L19
	}
L19:
	;
	v89 = v85
	goto L9
L20:
	;
	v92 = int32(0)
	if base.B2i32(v34 == v92)|base.B2i32(v89 != v92) == v92 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v58)+52))
	v104 = F_heap_form_tuple(m, v99, v11+int32(16), v11+int32(12))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L4
	} else {
		goto L24
	}
L22:
	;
	v108 = v89
	goto L23
L23:
	;
	if v108 != 0 {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	F_CatalogTupleInsert(m, v58, v104)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L4
	} else {
		goto L25
	}
L25:
	;
	v108 = v104
	goto L23
L26:
	;
	F_pfree(m, v108)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L4
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	F_sequence_close(m, v58, int32(0))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
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
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v19 int64
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v43 int64
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v12 = F_palloc0(m, int32(420))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = v12 + int32(276)
		v19 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
		*(*int64)(unsafe.Add(mBase, uint32(v12))) = v19
		v22 = l0 + int32(8)
		v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
		*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v23
		v25 = int32(112)
		*(*uint8)(unsafe.Add(mBase, uint32(v12)+394)) = uint8(v25)
		*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = int32(-1)
		v29 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
		*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v29
		v36 = F_pg_sprintf(m, v12+int32(280), int32(59295), v9+int32(16))
		mBase = m.M
		v37 = m.ExcPending
		if v37 != 0 {
			return int32(0)
		} else {
			v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			*(*int32)(unsafe.Add(mBase, uint32(v12)+60)) = v29
			*(*int32)(unsafe.Add(mBase, uint32(v12)+64)) = v38
			v41 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
			*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v41
			v43 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
			*(*int64)(unsafe.Add(mBase, uint32(v9))) = v43
			v46 = F_smgropen(m, v9, int32(-1))
			mBase = m.M
			v47 = m.ExcPending
			if v47 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v46
				m.G0 = v9 + int32(32)
				return v12
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
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v109 int32
	_ = v109
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v162 int32
	_ = v162
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v208 int32
	_ = v208
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v229 int32
	_ = v229
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v242 int32
	_ = v242
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v258 int32
	_ = v258
	var v263 int32
	_ = v263
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
	v20 = *(*int32)(unsafe.Add(mBase, _consts[642]))
	if v20 != 0 {
		goto L7
	} else {
		goto L8
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L1
	} else {
		goto L66
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L1
	} else {
		goto L62
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L1
	} else {
		goto L58
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
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
	v191 = m.ExcPending
	if v191 != 0 {
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
	v27 = *(*int32)(unsafe.Add(mBase, _consts[4]))
	if v23 != v27 {
		goto L4
	} else {
		goto L12
	}
L12:
	;
	v30 = *(*int32)(unsafe.Add(mBase, _consts[61]))
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
	v35 = *(*int32)(unsafe.Add(mBase, _consts[63]))
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
	F_s_lock(m, v20, int32(494412), int32(387), int32(62156))
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
	v50 = *(*int64)(unsafe.Add(mBase, uint32(v13)+4))
	*(*int64)(unsafe.Add(mBase, uint32(v20)+137)) = v50
	v52 = *(*int64)(unsafe.Add(mBase, uint32(v13)+60))
	*(*int64)(unsafe.Add(mBase, uint32(v20)+193)) = v52
	v54 = *(*int64)(unsafe.Add(mBase, uint32(v13)+52))
	*(*int64)(unsafe.Add(mBase, uint32(v20)+185)) = v54
	v56 = *(*int64)(unsafe.Add(mBase, uint32(v13)+44))
	*(*int64)(unsafe.Add(mBase, uint32(v20)+177)) = v56
	v58 = *(*int64)(unsafe.Add(mBase, uint32(v13)+36))
	*(*int64)(unsafe.Add(mBase, uint32(v20)+169)) = v58
	v60 = *(*int64)(unsafe.Add(mBase, uint32(v13)+28))
	*(*int64)(unsafe.Add(mBase, uint32(v20)+161)) = v60
	v62 = *(*int64)(unsafe.Add(mBase, uint32(v13)+20))
	*(*int64)(unsafe.Add(mBase, uint32(v20)+153)) = v62
	v64 = *(*int64)(unsafe.Add(mBase, uint32(v13)+12))
	*(*int64)(unsafe.Add(mBase, uint32(v20)+145)) = v64
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
	v85 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v89 = F_LWLockAcquire(m, v85+int32(4736), int32(0))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
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
	F_s_lock(m, v20, int32(494412), int32(395), int32(62156))
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
	v92 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v96 = F_LWLockAcquire(m, v92+int32(512), int32(0))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	v100 = F_GetOldestSafeDecodingTransactionId(m, l1^int32(1))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = int32(1)
	if v102 != 0 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	F_s_lock(m, v20, int32(494412), int32(430), int32(62156))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L1
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+100)) = v100
	*(*int32)(unsafe.Add(mBase, uint32(v20)+20)) = v100
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
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = v100
	goto L41
L40:
	;
	goto L41
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = int32(0)
	F_ReplicationSlotsComputeRequiredXmin(m, int32(1))
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	v119 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	F_LWLockRelease(m, v119+int32(512))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	v125 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	F_LWLockRelease(m, v125+int32(4736))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	F_ReplicationSlotMarkDirty(m)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	F_ReplicationSlotSave(m)
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	v134 = int32(0)
	v137 = F_StartupDecodingContext(m, v134, l2, v100, l1, v134, int32(1), l3, l4, l5, l6)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	v139 = int32(4486928)
	v140 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v137)))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v142
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v137)+24))
	if v144 != 0 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v13)+88)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+84)) = int32(230936)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+80)) = v137
	*(*int32)(unsafe.Add(mBase, uint32(v13)+72)) = int32(993)
	v152 = int32(4479832)
	v153 = *(*int32)(unsafe.Add(mBase, _consts[84]))
	*(*int32)(unsafe.Add(mBase, _consts[84])) = v13 + int32(68)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+68)) = v153
	*(*int32)(unsafe.Add(mBase, uint32(v13)+76)) = v13 + int32(80)
	v162 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v137)+164)) = uint8(v162)
	*(*uint8)(unsafe.Add(mBase, uint32(v137)+147)) = uint8(v162)
	m.T0[v144].(func(*base.Module, int32, int32, int32))(m, v137, v137+int32(108), int32(1))
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L1
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v140
	v177 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v137)+145)))
	v178 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+136)))
	v179 = v177 & v178
	*(*uint8)(unsafe.Add(mBase, uint32(v137)+145)) = uint8(v179)
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v137)+12))
	v182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v137)+112)))
	*(*uint8)(unsafe.Add(mBase, uint32(v181)+116)) = uint8(v182)
	m.G0 = v13 + int32(96)
	return v137
L51:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v13)+68))
	*(*int32)(unsafe.Add(mBase, _consts[84])) = v172
	goto L50
L52:
	;
	F_errmsg_internal(m, int32(85669), int32(0))
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	F_errfinish(m, int32(494412), int32(358), int32(62156))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
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
	F_errmsg_internal(m, int32(274822), int32(0))
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	F_errfinish(m, int32(494412), int32(361), int32(62156))
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
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
	v220 = m.ExcPending
	if v220 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	F_errmsg(m, int32(333999), int32(0))
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	F_errfinish(m, int32(494412), int32(367), int32(62156))
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
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
	v236 = m.ExcPending
	if v236 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v20 + int32(24)
	F_errmsg(m, int32(359582), v13)
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	F_errfinish(m, int32(494412), int32(373), int32(62156))
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
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
	v254 = m.ExcPending
	if v254 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	F_errmsg(m, int32(159159), int32(0))
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L1
	} else {
		goto L68
	}
L68:
	;
	F_errfinish(m, int32(494412), int32(379), int32(62156))
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
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
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int64
	_ = v41
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v64 int64
	_ = v64
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v97 int64
	_ = v97
	var v98 int32
	_ = v98
	var v99 int64
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v108 int64
	_ = v108
	var v110 int32
	_ = v110
	v6 = m.G0
	v8 = v6 - int32(2128)
	m.G0 = v8
	v12 = *(*int32)(unsafe.Add(mBase, _consts[276]))
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
	v37 = F_pg_snprintf(m, v8+int32(32), int32(1061), int32(39033), v8+int32(16))
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
	v21 = *(*int32)(unsafe.Add(mBase, _consts[276]))
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
	v41 = F_db_dir_size(m, v8+int32(32))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L2
	} else {
		goto L10
	}
L10:
	;
	v48 = F_pg_snprintf(m, v8+int32(1104), int32(1024), int32(485645), int32(0))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L2
	} else {
		goto L11
	}
L11:
	;
	v52 = F_AllocateDir(m, v8+int32(1104))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L2
	} else {
		goto L13
	}
L12:
	;
	F_FreeDir(m, v52)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L2
	} else {
		goto L35
	}
L13:
	;
	v56 = F_ReadDir(m, v52, v8+int32(1104))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L2
	} else {
		goto L14
	}
L14:
	;
	if v56 == int32(0) {
		v108 = v41
		goto L12
	} else {
		goto L15
	}
L15:
	;
	v62 = v56
	v64 = v41
	goto L16
L16:
	;
	v66 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	if v66 != 0 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v108 = v99
	goto L12
L18:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L2
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62)+19)))
	if v69 != int32(46) {
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
	*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = int32(550670)
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = int32(485645)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v62 + int32(19)
	v93 = F_pg_snprintf(m, v8+int32(32), int32(1061), int32(39011), v8)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L2
	} else {
		goto L31
	}
L23:
	;
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62)+20)))
	if v72 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62)+20)))
	if v73 != int32(46) {
		goto L22
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	v79 = F_ReadDir(m, v52, v8+int32(1104))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L2
	} else {
		goto L29
	}
L27:
	;
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62)+21)))
	if v76 != 0 {
		goto L22
	} else {
		goto L28
	}
L28:
	;
	goto L26
L29:
	;
	if v79 != 0 {
		v62 = v79
		goto L16
	} else {
		goto L30
	}
L30:
	;
	v108 = v64
	goto L12
L31:
	;
	v97 = F_db_dir_size(m, v8+int32(32))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L2
	} else {
		goto L32
	}
L32:
	;
	v99 = v97 + v64
	v102 = F_ReadDir(m, v52, v8+int32(1104))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L2
	} else {
		goto L33
	}
L33:
	;
	if v102 != 0 {
		v62 = v102
		v64 = v99
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
	return v108
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
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v11 = *(*int32)(unsafe.Add(mBase, _consts[765]))
	if v11 <= int32(0) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v34 = m.ExcPending
		if v34 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = l1
			*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
			F_errmsg_internal(m, int32(12161), v8)
			mBase = m.M
			v39 = m.ExcPending
			if v39 != 0 {
				return
			} else {
				F_errfinish(m, int32(496590), int32(410), int32(98961))
				mBase = m.M
				v44 = m.ExcPending
				if v44 != 0 {
					return
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	} else {
		v15 = v11 - int32(1)
		v17 = v15 << (uint(int32(3)) % 32)
		v20 = *(*int32)(unsafe.Add(mBase, uint32(v17)+uint32(_consts[766])))
		if v20 != l0 {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v34 = m.ExcPending
			if v34 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = l1
				*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
				F_errmsg_internal(m, int32(12161), v8)
				mBase = m.M
				v39 = m.ExcPending
				if v39 != 0 {
					return
				} else {
					F_errfinish(m, int32(496590), int32(410), int32(98961))
					mBase = m.M
					v44 = m.ExcPending
					if v44 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v22 = *(*int32)(unsafe.Add(mBase, uint32(v17)+uint32(_consts[767])))
			if v22 != l1 {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = l1
					*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
					F_errmsg_internal(m, int32(12161), v8)
					mBase = m.M
					v39 = m.ExcPending
					if v39 != 0 {
						return
					} else {
						F_errfinish(m, int32(496590), int32(410), int32(98961))
						mBase = m.M
						v44 = m.ExcPending
						if v44 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				*(*int32)(unsafe.Add(mBase, _consts[765])) = v15
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
	v14 = *(*int32)(unsafe.Add(mBase, _consts[586]))
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
	v68 = *(*int32)(unsafe.Add(mBase, _consts[586]))
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
	v47 = *(*int32)(unsafe.Add(mBase, _consts[587]))
	if base.Ui32(int32(127)) < base.Ui32(v9) {
		goto L24
	} else {
		goto L25
	}
L11:
	;
	v35 = *(*int32)(unsafe.Add(mBase, _consts[587]))
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
	v31 = v9<<(uint(int32(2))%32) + int32(1854276)
	goto L17
L16:
	;
	v26 = F_case_index(m, v9)
	mBase = m.M
	v31 = v26<<(uint(int32(2))%32) + int32(1854272)
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
	v101 = *(*int32)(unsafe.Add(mBase, _consts[587]))
	if base.Ui32(int32(127)) < base.Ui32(v65) {
		goto L44
	} else {
		goto L45
	}
L31:
	;
	v89 = *(*int32)(unsafe.Add(mBase, _consts[587]))
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
	v85 = v65<<(uint(int32(2))%32) + int32(1854276)
	goto L37
L36:
	;
	v80 = F_case_index(m, v65)
	mBase = m.M
	v85 = v80<<(uint(int32(2))%32) + int32(1854272)
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
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v217 int32
	_ = v217
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v243 int32
	_ = v243
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v263 int32
	_ = v263
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v318 int32
	_ = v318
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v352 int32
	_ = v352
	var v354 int32
	_ = v354
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v364 int32
	_ = v364
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
				v359 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				if v359 != 0 {
					v361 = v359
				} else {
					v361 = int32(12)
				}
				*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v361
				v364 = int32(0)
			} else {
				v364 = v23
			}
			return v364
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
				v359 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				if v359 != 0 {
					v361 = v359
				} else {
					v361 = int32(12)
				}
				*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v361
				v364 = int32(0)
			} else {
				v364 = v29
			}
			return v364
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
							v352 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							if v352 != 0 {
								v354 = v352
							} else {
								v354 = int32(12)
							}
							*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v354
							*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
							v359 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							if v359 != 0 {
								v361 = v359
							} else {
								v361 = int32(12)
							}
							*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v361
							v364 = int32(0)
							return v364
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
							v89 = *(*int32)(unsafe.Add(mBase, uint32(v75)+12))
							*(*int32)(unsafe.Add(mBase, uint32(v75)+12)) = v89 + int32(1)
							return v75
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
								v352 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
								if v352 != 0 {
									v354 = v352
								} else {
									v354 = int32(12)
								}
								*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v354
								*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
								v359 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
								if v359 != 0 {
									v361 = v359
								} else {
									v361 = int32(12)
								}
								*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v361
								v364 = int32(0)
								return v364
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
								v89 = *(*int32)(unsafe.Add(mBase, uint32(v75)+12))
								*(*int32)(unsafe.Add(mBase, uint32(v75)+12)) = v89 + int32(1)
								return v75
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
					v89 = *(*int32)(unsafe.Add(mBase, uint32(v75)+12))
					*(*int32)(unsafe.Add(mBase, uint32(v75)+12)) = v89 + int32(1)
					return v75
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
					v352 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					if v352 != 0 {
						v354 = v352
					} else {
						v354 = int32(12)
					}
					*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v354
					*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
					v359 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					if v359 != 0 {
						v361 = v359
					} else {
						v361 = int32(12)
					}
					*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v361
					v364 = int32(0)
					return v364
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
					v89 = *(*int32)(unsafe.Add(mBase, uint32(v75)+12))
					*(*int32)(unsafe.Add(mBase, uint32(v75)+12)) = v89 + int32(1)
					return v75
				}
			}
		}
	case 3:
		v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
		if v94 != 0 {
			v95 = *(*int32)(unsafe.Add(mBase, uint32(v94)+4))
			if v95 < int32(2) {
				F_pfree(m, v94)
				mBase = m.M
				v110 = m.ExcPending
				if v110 != 0 {
					return int32(0)
				} else {
					v114 = F_palloc_extended(m, int32(36), int32(2))
					mBase = m.M
					v115 = m.ExcPending
					if v115 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v114)+24)) = int32(-1)
						*(*int64)(unsafe.Add(mBase, uint32(v114)+12)) = int64(0)
						*(*int64)(unsafe.Add(mBase, uint32(v114))) = int64(8589934592)
						*(*int32)(unsafe.Add(mBase, uint32(v114)+20)) = v114 + int32(36)
						*(*int32)(unsafe.Add(mBase, uint32(v114)+8)) = v114 + int32(28)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = v114
						v129 = *(*int32)(unsafe.Add(mBase, uint32(v114)))
						v130 = v114
						v131 = v129
						v132 = int32(1)
						*(*int32)(unsafe.Add(mBase, uint32(v130))) = v131 + v132
						v135 = *(*int32)(unsafe.Add(mBase, uint32(v130)+8))
						v136 = int32(2)
						*(*int32)(unsafe.Add(mBase, uint32(v135+v131<<(uint(v136)%32)))) = int32(9)
						v141 = *(*int32)(unsafe.Add(mBase, uint32(v130)))
						*(*int32)(unsafe.Add(mBase, uint32(v130))) = v141 + v132
						v145 = *(*int32)(unsafe.Add(mBase, uint32(v130)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v145+v141<<(uint(v136)%32)))) = int32(32)
						return v130
					}
				}
			} else {
				v98 = int32(0)
				v99 = *(*int32)(unsafe.Add(mBase, uint32(v94)+16))
				if v99 < v98 {
					F_pfree(m, v94)
					mBase = m.M
					v110 = m.ExcPending
					if v110 != 0 {
						return int32(0)
					} else {
						v114 = F_palloc_extended(m, int32(36), int32(2))
						mBase = m.M
						v115 = m.ExcPending
						if v115 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v114)+24)) = int32(-1)
							*(*int64)(unsafe.Add(mBase, uint32(v114)+12)) = int64(0)
							*(*int64)(unsafe.Add(mBase, uint32(v114))) = int64(8589934592)
							*(*int32)(unsafe.Add(mBase, uint32(v114)+20)) = v114 + int32(36)
							*(*int32)(unsafe.Add(mBase, uint32(v114)+8)) = v114 + int32(28)
							*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = v114
							v129 = *(*int32)(unsafe.Add(mBase, uint32(v114)))
							v130 = v114
							v131 = v129
							v132 = int32(1)
							*(*int32)(unsafe.Add(mBase, uint32(v130))) = v131 + v132
							v135 = *(*int32)(unsafe.Add(mBase, uint32(v130)+8))
							v136 = int32(2)
							*(*int32)(unsafe.Add(mBase, uint32(v135+v131<<(uint(v136)%32)))) = int32(9)
							v141 = *(*int32)(unsafe.Add(mBase, uint32(v130)))
							*(*int32)(unsafe.Add(mBase, uint32(v130))) = v141 + v132
							v145 = *(*int32)(unsafe.Add(mBase, uint32(v130)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v145+v141<<(uint(v136)%32)))) = int32(32)
							return v130
						}
					}
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v94)+24)) = int32(-1)
					v104 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v94)+12)) = v104
					*(*int32)(unsafe.Add(mBase, uint32(v94))) = v104
					v130 = v94
					v131 = v98
					v132 = int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(v130))) = v131 + v132
					v135 = *(*int32)(unsafe.Add(mBase, uint32(v130)+8))
					v136 = int32(2)
					*(*int32)(unsafe.Add(mBase, uint32(v135+v131<<(uint(v136)%32)))) = int32(9)
					v141 = *(*int32)(unsafe.Add(mBase, uint32(v130)))
					*(*int32)(unsafe.Add(mBase, uint32(v130))) = v141 + v132
					v145 = *(*int32)(unsafe.Add(mBase, uint32(v130)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v145+v141<<(uint(v136)%32)))) = int32(32)
					return v130
				}
			}
		} else {
			v114 = F_palloc_extended(m, int32(36), int32(2))
			mBase = m.M
			v115 = m.ExcPending
			if v115 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v114)+24)) = int32(-1)
				*(*int64)(unsafe.Add(mBase, uint32(v114)+12)) = int64(0)
				*(*int64)(unsafe.Add(mBase, uint32(v114))) = int64(8589934592)
				*(*int32)(unsafe.Add(mBase, uint32(v114)+20)) = v114 + int32(36)
				*(*int32)(unsafe.Add(mBase, uint32(v114)+8)) = v114 + int32(28)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = v114
				v129 = *(*int32)(unsafe.Add(mBase, uint32(v114)))
				v130 = v114
				v131 = v129
				v132 = int32(1)
				*(*int32)(unsafe.Add(mBase, uint32(v130))) = v131 + v132
				v135 = *(*int32)(unsafe.Add(mBase, uint32(v130)+8))
				v136 = int32(2)
				*(*int32)(unsafe.Add(mBase, uint32(v135+v131<<(uint(v136)%32)))) = int32(9)
				v141 = *(*int32)(unsafe.Add(mBase, uint32(v130)))
				*(*int32)(unsafe.Add(mBase, uint32(v130))) = v141 + v132
				v145 = *(*int32)(unsafe.Add(mBase, uint32(v130)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v145+v141<<(uint(v136)%32)))) = int32(32)
				return v130
			}
		}
	case 4:
		v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
		if v152 != 0 {
			v153 = *(*int32)(unsafe.Add(mBase, uint32(v152)+4))
			if v153 < int32(0) {
				F_pfree(m, v152)
				mBase = m.M
				v167 = m.ExcPending
				if v167 != 0 {
					return int32(0)
				} else {
					v170 = F_palloc_extended(m, int32(44), int32(2))
					mBase = m.M
					v171 = m.ExcPending
					if v171 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v170)+24)) = int32(-1)
						*(*int64)(unsafe.Add(mBase, uint32(v170)+12)) = int64(8589934592)
						*(*int64)(unsafe.Add(mBase, uint32(v170))) = int64(0)
						v179 = v170 + int32(28)
						*(*int32)(unsafe.Add(mBase, uint32(v170)+20)) = v179
						*(*int32)(unsafe.Add(mBase, uint32(v170)+8)) = v179
						*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = v170
						v183 = *(*int32)(unsafe.Add(mBase, uint32(v170)+12))
						v186 = v170
						v188 = v183 << (uint(int32(3)) % 32)
						v189 = *(*int32)(unsafe.Add(mBase, uint32(v186)+20))
						*(*int32)(unsafe.Add(mBase, uint32(v188+v189))) = int32(0)
						v193 = *(*int32)(unsafe.Add(mBase, uint32(v186)+20))
						v194 = *(*int32)(unsafe.Add(mBase, uint32(v186)+12))
						v195 = int32(3)
						*(*int32)(unsafe.Add(mBase, uint32(v193+v194<<(uint(v195)%32))+4)) = int32(31)
						v200 = *(*int32)(unsafe.Add(mBase, uint32(v186)+12))
						v201 = int32(1)
						v202 = v200 + v201
						*(*int32)(unsafe.Add(mBase, uint32(v186)+12)) = v202
						v204 = *(*int32)(unsafe.Add(mBase, uint32(v186)+20))
						*(*int32)(unsafe.Add(mBase, uint32(v204+v202<<(uint(v195)%32)))) = int32(127)
						v210 = *(*int32)(unsafe.Add(mBase, uint32(v186)+20))
						v211 = *(*int32)(unsafe.Add(mBase, uint32(v186)+12))
						*(*int32)(unsafe.Add(mBase, uint32(v210+v211<<(uint(v195)%32))+4)) = int32(159)
						v217 = *(*int32)(unsafe.Add(mBase, uint32(v186)+12))
						*(*int32)(unsafe.Add(mBase, uint32(v186)+12)) = v217 + v201
						return v186
					}
				}
			} else {
				v156 = *(*int32)(unsafe.Add(mBase, uint32(v152)+16))
				if v156 < int32(2) {
					F_pfree(m, v152)
					mBase = m.M
					v167 = m.ExcPending
					if v167 != 0 {
						return int32(0)
					} else {
						v170 = F_palloc_extended(m, int32(44), int32(2))
						mBase = m.M
						v171 = m.ExcPending
						if v171 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v170)+24)) = int32(-1)
							*(*int64)(unsafe.Add(mBase, uint32(v170)+12)) = int64(8589934592)
							*(*int64)(unsafe.Add(mBase, uint32(v170))) = int64(0)
							v179 = v170 + int32(28)
							*(*int32)(unsafe.Add(mBase, uint32(v170)+20)) = v179
							*(*int32)(unsafe.Add(mBase, uint32(v170)+8)) = v179
							*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = v170
							v183 = *(*int32)(unsafe.Add(mBase, uint32(v170)+12))
							v186 = v170
							v188 = v183 << (uint(int32(3)) % 32)
							v189 = *(*int32)(unsafe.Add(mBase, uint32(v186)+20))
							*(*int32)(unsafe.Add(mBase, uint32(v188+v189))) = int32(0)
							v193 = *(*int32)(unsafe.Add(mBase, uint32(v186)+20))
							v194 = *(*int32)(unsafe.Add(mBase, uint32(v186)+12))
							v195 = int32(3)
							*(*int32)(unsafe.Add(mBase, uint32(v193+v194<<(uint(v195)%32))+4)) = int32(31)
							v200 = *(*int32)(unsafe.Add(mBase, uint32(v186)+12))
							v201 = int32(1)
							v202 = v200 + v201
							*(*int32)(unsafe.Add(mBase, uint32(v186)+12)) = v202
							v204 = *(*int32)(unsafe.Add(mBase, uint32(v186)+20))
							*(*int32)(unsafe.Add(mBase, uint32(v204+v202<<(uint(v195)%32)))) = int32(127)
							v210 = *(*int32)(unsafe.Add(mBase, uint32(v186)+20))
							v211 = *(*int32)(unsafe.Add(mBase, uint32(v186)+12))
							*(*int32)(unsafe.Add(mBase, uint32(v210+v211<<(uint(v195)%32))+4)) = int32(159)
							v217 = *(*int32)(unsafe.Add(mBase, uint32(v186)+12))
							*(*int32)(unsafe.Add(mBase, uint32(v186)+12)) = v217 + v201
							return v186
						}
					}
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v152)+24)) = int32(-1)
					v161 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v152)+12)) = v161
					*(*int32)(unsafe.Add(mBase, uint32(v152))) = v161
					v186 = v152
					v188 = v161
					v189 = *(*int32)(unsafe.Add(mBase, uint32(v186)+20))
					*(*int32)(unsafe.Add(mBase, uint32(v188+v189))) = int32(0)
					v193 = *(*int32)(unsafe.Add(mBase, uint32(v186)+20))
					v194 = *(*int32)(unsafe.Add(mBase, uint32(v186)+12))
					v195 = int32(3)
					*(*int32)(unsafe.Add(mBase, uint32(v193+v194<<(uint(v195)%32))+4)) = int32(31)
					v200 = *(*int32)(unsafe.Add(mBase, uint32(v186)+12))
					v201 = int32(1)
					v202 = v200 + v201
					*(*int32)(unsafe.Add(mBase, uint32(v186)+12)) = v202
					v204 = *(*int32)(unsafe.Add(mBase, uint32(v186)+20))
					*(*int32)(unsafe.Add(mBase, uint32(v204+v202<<(uint(v195)%32)))) = int32(127)
					v210 = *(*int32)(unsafe.Add(mBase, uint32(v186)+20))
					v211 = *(*int32)(unsafe.Add(mBase, uint32(v186)+12))
					*(*int32)(unsafe.Add(mBase, uint32(v210+v211<<(uint(v195)%32))+4)) = int32(159)
					v217 = *(*int32)(unsafe.Add(mBase, uint32(v186)+12))
					*(*int32)(unsafe.Add(mBase, uint32(v186)+12)) = v217 + v201
					return v186
				}
			}
		} else {
			v170 = F_palloc_extended(m, int32(44), int32(2))
			mBase = m.M
			v171 = m.ExcPending
			if v171 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v170)+24)) = int32(-1)
				*(*int64)(unsafe.Add(mBase, uint32(v170)+12)) = int64(8589934592)
				*(*int64)(unsafe.Add(mBase, uint32(v170))) = int64(0)
				v179 = v170 + int32(28)
				*(*int32)(unsafe.Add(mBase, uint32(v170)+20)) = v179
				*(*int32)(unsafe.Add(mBase, uint32(v170)+8)) = v179
				*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = v170
				v183 = *(*int32)(unsafe.Add(mBase, uint32(v170)+12))
				v186 = v170
				v188 = v183 << (uint(int32(3)) % 32)
				v189 = *(*int32)(unsafe.Add(mBase, uint32(v186)+20))
				*(*int32)(unsafe.Add(mBase, uint32(v188+v189))) = int32(0)
				v193 = *(*int32)(unsafe.Add(mBase, uint32(v186)+20))
				v194 = *(*int32)(unsafe.Add(mBase, uint32(v186)+12))
				v195 = int32(3)
				*(*int32)(unsafe.Add(mBase, uint32(v193+v194<<(uint(v195)%32))+4)) = int32(31)
				v200 = *(*int32)(unsafe.Add(mBase, uint32(v186)+12))
				v201 = int32(1)
				v202 = v200 + v201
				*(*int32)(unsafe.Add(mBase, uint32(v186)+12)) = v202
				v204 = *(*int32)(unsafe.Add(mBase, uint32(v186)+20))
				*(*int32)(unsafe.Add(mBase, uint32(v204+v202<<(uint(v195)%32)))) = int32(127)
				v210 = *(*int32)(unsafe.Add(mBase, uint32(v186)+20))
				v211 = *(*int32)(unsafe.Add(mBase, uint32(v186)+12))
				*(*int32)(unsafe.Add(mBase, uint32(v210+v211<<(uint(v195)%32))+4)) = int32(159)
				v217 = *(*int32)(unsafe.Add(mBase, uint32(v186)+12))
				*(*int32)(unsafe.Add(mBase, uint32(v186)+12)) = v217 + v201
				return v186
			}
		}
	case 5:
		v224 = F_pg_ctype_get_cache(m, int32(976), int32(5))
		mBase = m.M
		v225 = m.ExcPending
		if v225 != 0 {
			return int32(0)
		} else {
			if v224 == int32(0) {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
				v359 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				if v359 != 0 {
					v361 = v359
				} else {
					v361 = int32(12)
				}
				*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v361
				v364 = int32(0)
			} else {
				v364 = v224
			}
			return v364
		}
	case 6:
		v343 = F_pg_ctype_get_cache(m, int32(981), int32(6))
		mBase = m.M
		v344 = m.ExcPending
		if v344 != 0 {
			return int32(0)
		} else {
			if v343 == int32(0) {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
				v359 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				if v359 != 0 {
					v361 = v359
				} else {
					v361 = int32(12)
				}
				*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v361
				v364 = int32(0)
			} else {
				v364 = v343
			}
			return v364
		}
	case 7:
		v331 = F_pg_ctype_get_cache(m, int32(979), int32(7))
		mBase = m.M
		v332 = m.ExcPending
		if v332 != 0 {
			return int32(0)
		} else {
			if v331 == int32(0) {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
				v359 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				if v359 != 0 {
					v361 = v359
				} else {
					v361 = int32(12)
				}
				*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v361
				v364 = int32(0)
			} else {
				v364 = v331
			}
			return v364
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
				v359 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				if v359 != 0 {
					v361 = v359
				} else {
					v361 = int32(12)
				}
				*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v361
				v364 = int32(0)
			} else {
				v364 = v15
			}
			return v364
		}
	case 9:
		v230 = F_pg_ctype_get_cache(m, int32(977), int32(9))
		mBase = m.M
		v231 = m.ExcPending
		if v231 != 0 {
			return int32(0)
		} else {
			if v230 == int32(0) {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
				v359 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				if v359 != 0 {
					v361 = v359
				} else {
					v361 = int32(12)
				}
				*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v361
				v364 = int32(0)
			} else {
				v364 = v230
			}
			return v364
		}
	case 10:
		v325 = F_pg_ctype_get_cache(m, int32(978), int32(10))
		mBase = m.M
		v326 = m.ExcPending
		if v326 != 0 {
			return int32(0)
		} else {
			if v325 == int32(0) {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
				v359 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				if v359 != 0 {
					v361 = v359
				} else {
					v361 = int32(12)
				}
				*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v361
				v364 = int32(0)
			} else {
				v364 = v325
			}
			return v364
		}
	case 11:
		v337 = F_pg_ctype_get_cache(m, int32(980), int32(11))
		mBase = m.M
		v338 = m.ExcPending
		if v338 != 0 {
			return int32(0)
		} else {
			if v337 == int32(0) {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
				v359 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				if v359 != 0 {
					v361 = v359
				} else {
					v361 = int32(12)
				}
				*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v361
				v364 = int32(0)
			} else {
				v364 = v337
			}
			return v364
		}
	case 12:
		v234 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
		if v234 != 0 {
			v235 = *(*int32)(unsafe.Add(mBase, uint32(v234)+4))
			if v235 < int32(0) {
				F_pfree(m, v234)
				mBase = m.M
				v249 = m.ExcPending
				if v249 != 0 {
					return int32(0)
				} else {
					v252 = F_palloc_extended(m, int32(52), int32(2))
					mBase = m.M
					v253 = m.ExcPending
					if v253 != 0 {
						return int32(0)
					} else {
						if v252 == int32(0) {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = int32(0)
							v352 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							if v352 != 0 {
								v354 = v352
							} else {
								v354 = int32(12)
							}
							*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v354
							*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
							v359 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							if v359 != 0 {
								v361 = v359
							} else {
								v361 = int32(12)
							}
							*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v361
							v364 = int32(0)
							return v364
						} else {
							*(*int64)(unsafe.Add(mBase, uint32(v252))) = int64(0)
							*(*int32)(unsafe.Add(mBase, uint32(v252)+24)) = int32(-1)
							*(*int64)(unsafe.Add(mBase, uint32(v252)+12)) = int64(12884901888)
							v263 = v252 + int32(28)
							*(*int32)(unsafe.Add(mBase, uint32(v252)+20)) = v263
							*(*int32)(unsafe.Add(mBase, uint32(v252)+8)) = v263
							*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = v252
							v267 = *(*int32)(unsafe.Add(mBase, uint32(v252)+12))
							v270 = v252
							v272 = v267 << (uint(int32(3)) % 32)
							v273 = *(*int32)(unsafe.Add(mBase, uint32(v270)+20))
							*(*int32)(unsafe.Add(mBase, uint32(v272+v273))) = int32(48)
							v277 = *(*int32)(unsafe.Add(mBase, uint32(v270)+20))
							v278 = *(*int32)(unsafe.Add(mBase, uint32(v270)+12))
							v279 = int32(3)
							*(*int32)(unsafe.Add(mBase, uint32(v277+v278<<(uint(v279)%32))+4)) = int32(57)
							v284 = *(*int32)(unsafe.Add(mBase, uint32(v270)+12))
							v285 = int32(1)
							v286 = v284 + v285
							*(*int32)(unsafe.Add(mBase, uint32(v270)+12)) = v286
							v288 = *(*int32)(unsafe.Add(mBase, uint32(v270)+20))
							*(*int32)(unsafe.Add(mBase, uint32(v288+v286<<(uint(v279)%32)))) = int32(97)
							v294 = *(*int32)(unsafe.Add(mBase, uint32(v270)+20))
							v295 = *(*int32)(unsafe.Add(mBase, uint32(v270)+12))
							*(*int32)(unsafe.Add(mBase, uint32(v294+v295<<(uint(v279)%32))+4)) = int32(102)
							v301 = *(*int32)(unsafe.Add(mBase, uint32(v270)+12))
							v303 = v301 + v285
							*(*int32)(unsafe.Add(mBase, uint32(v270)+12)) = v303
							v305 = *(*int32)(unsafe.Add(mBase, uint32(v270)+20))
							*(*int32)(unsafe.Add(mBase, uint32(v305+v303<<(uint(v279)%32)))) = int32(65)
							v311 = *(*int32)(unsafe.Add(mBase, uint32(v270)+20))
							v312 = *(*int32)(unsafe.Add(mBase, uint32(v270)+12))
							*(*int32)(unsafe.Add(mBase, uint32(v311+v312<<(uint(v279)%32))+4)) = int32(70)
							v318 = *(*int32)(unsafe.Add(mBase, uint32(v270)+12))
							*(*int32)(unsafe.Add(mBase, uint32(v270)+12)) = v318 + v285
							return v270
						}
					}
				}
			} else {
				v238 = *(*int32)(unsafe.Add(mBase, uint32(v234)+16))
				if v238 < int32(3) {
					F_pfree(m, v234)
					mBase = m.M
					v249 = m.ExcPending
					if v249 != 0 {
						return int32(0)
					} else {
						v252 = F_palloc_extended(m, int32(52), int32(2))
						mBase = m.M
						v253 = m.ExcPending
						if v253 != 0 {
							return int32(0)
						} else {
							if v252 == int32(0) {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = int32(0)
								v352 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
								if v352 != 0 {
									v354 = v352
								} else {
									v354 = int32(12)
								}
								*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v354
								*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
								v359 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
								if v359 != 0 {
									v361 = v359
								} else {
									v361 = int32(12)
								}
								*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v361
								v364 = int32(0)
								return v364
							} else {
								*(*int64)(unsafe.Add(mBase, uint32(v252))) = int64(0)
								*(*int32)(unsafe.Add(mBase, uint32(v252)+24)) = int32(-1)
								*(*int64)(unsafe.Add(mBase, uint32(v252)+12)) = int64(12884901888)
								v263 = v252 + int32(28)
								*(*int32)(unsafe.Add(mBase, uint32(v252)+20)) = v263
								*(*int32)(unsafe.Add(mBase, uint32(v252)+8)) = v263
								*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = v252
								v267 = *(*int32)(unsafe.Add(mBase, uint32(v252)+12))
								v270 = v252
								v272 = v267 << (uint(int32(3)) % 32)
								v273 = *(*int32)(unsafe.Add(mBase, uint32(v270)+20))
								*(*int32)(unsafe.Add(mBase, uint32(v272+v273))) = int32(48)
								v277 = *(*int32)(unsafe.Add(mBase, uint32(v270)+20))
								v278 = *(*int32)(unsafe.Add(mBase, uint32(v270)+12))
								v279 = int32(3)
								*(*int32)(unsafe.Add(mBase, uint32(v277+v278<<(uint(v279)%32))+4)) = int32(57)
								v284 = *(*int32)(unsafe.Add(mBase, uint32(v270)+12))
								v285 = int32(1)
								v286 = v284 + v285
								*(*int32)(unsafe.Add(mBase, uint32(v270)+12)) = v286
								v288 = *(*int32)(unsafe.Add(mBase, uint32(v270)+20))
								*(*int32)(unsafe.Add(mBase, uint32(v288+v286<<(uint(v279)%32)))) = int32(97)
								v294 = *(*int32)(unsafe.Add(mBase, uint32(v270)+20))
								v295 = *(*int32)(unsafe.Add(mBase, uint32(v270)+12))
								*(*int32)(unsafe.Add(mBase, uint32(v294+v295<<(uint(v279)%32))+4)) = int32(102)
								v301 = *(*int32)(unsafe.Add(mBase, uint32(v270)+12))
								v303 = v301 + v285
								*(*int32)(unsafe.Add(mBase, uint32(v270)+12)) = v303
								v305 = *(*int32)(unsafe.Add(mBase, uint32(v270)+20))
								*(*int32)(unsafe.Add(mBase, uint32(v305+v303<<(uint(v279)%32)))) = int32(65)
								v311 = *(*int32)(unsafe.Add(mBase, uint32(v270)+20))
								v312 = *(*int32)(unsafe.Add(mBase, uint32(v270)+12))
								*(*int32)(unsafe.Add(mBase, uint32(v311+v312<<(uint(v279)%32))+4)) = int32(70)
								v318 = *(*int32)(unsafe.Add(mBase, uint32(v270)+12))
								*(*int32)(unsafe.Add(mBase, uint32(v270)+12)) = v318 + v285
								return v270
							}
						}
					}
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v234)+24)) = int32(-1)
					v243 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v234)+12)) = v243
					*(*int32)(unsafe.Add(mBase, uint32(v234))) = v243
					v270 = v234
					v272 = v243
					v273 = *(*int32)(unsafe.Add(mBase, uint32(v270)+20))
					*(*int32)(unsafe.Add(mBase, uint32(v272+v273))) = int32(48)
					v277 = *(*int32)(unsafe.Add(mBase, uint32(v270)+20))
					v278 = *(*int32)(unsafe.Add(mBase, uint32(v270)+12))
					v279 = int32(3)
					*(*int32)(unsafe.Add(mBase, uint32(v277+v278<<(uint(v279)%32))+4)) = int32(57)
					v284 = *(*int32)(unsafe.Add(mBase, uint32(v270)+12))
					v285 = int32(1)
					v286 = v284 + v285
					*(*int32)(unsafe.Add(mBase, uint32(v270)+12)) = v286
					v288 = *(*int32)(unsafe.Add(mBase, uint32(v270)+20))
					*(*int32)(unsafe.Add(mBase, uint32(v288+v286<<(uint(v279)%32)))) = int32(97)
					v294 = *(*int32)(unsafe.Add(mBase, uint32(v270)+20))
					v295 = *(*int32)(unsafe.Add(mBase, uint32(v270)+12))
					*(*int32)(unsafe.Add(mBase, uint32(v294+v295<<(uint(v279)%32))+4)) = int32(102)
					v301 = *(*int32)(unsafe.Add(mBase, uint32(v270)+12))
					v303 = v301 + v285
					*(*int32)(unsafe.Add(mBase, uint32(v270)+12)) = v303
					v305 = *(*int32)(unsafe.Add(mBase, uint32(v270)+20))
					*(*int32)(unsafe.Add(mBase, uint32(v305+v303<<(uint(v279)%32)))) = int32(65)
					v311 = *(*int32)(unsafe.Add(mBase, uint32(v270)+20))
					v312 = *(*int32)(unsafe.Add(mBase, uint32(v270)+12))
					*(*int32)(unsafe.Add(mBase, uint32(v311+v312<<(uint(v279)%32))+4)) = int32(70)
					v318 = *(*int32)(unsafe.Add(mBase, uint32(v270)+12))
					*(*int32)(unsafe.Add(mBase, uint32(v270)+12)) = v318 + v285
					return v270
				}
			}
		} else {
			v252 = F_palloc_extended(m, int32(52), int32(2))
			mBase = m.M
			v253 = m.ExcPending
			if v253 != 0 {
				return int32(0)
			} else {
				if v252 == int32(0) {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = int32(0)
					v352 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					if v352 != 0 {
						v354 = v352
					} else {
						v354 = int32(12)
					}
					*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v354
					*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
					v359 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					if v359 != 0 {
						v361 = v359
					} else {
						v361 = int32(12)
					}
					*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v361
					v364 = int32(0)
					return v364
				} else {
					*(*int64)(unsafe.Add(mBase, uint32(v252))) = int64(0)
					*(*int32)(unsafe.Add(mBase, uint32(v252)+24)) = int32(-1)
					*(*int64)(unsafe.Add(mBase, uint32(v252)+12)) = int64(12884901888)
					v263 = v252 + int32(28)
					*(*int32)(unsafe.Add(mBase, uint32(v252)+20)) = v263
					*(*int32)(unsafe.Add(mBase, uint32(v252)+8)) = v263
					*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = v252
					v267 = *(*int32)(unsafe.Add(mBase, uint32(v252)+12))
					v270 = v252
					v272 = v267 << (uint(int32(3)) % 32)
					v273 = *(*int32)(unsafe.Add(mBase, uint32(v270)+20))
					*(*int32)(unsafe.Add(mBase, uint32(v272+v273))) = int32(48)
					v277 = *(*int32)(unsafe.Add(mBase, uint32(v270)+20))
					v278 = *(*int32)(unsafe.Add(mBase, uint32(v270)+12))
					v279 = int32(3)
					*(*int32)(unsafe.Add(mBase, uint32(v277+v278<<(uint(v279)%32))+4)) = int32(57)
					v284 = *(*int32)(unsafe.Add(mBase, uint32(v270)+12))
					v285 = int32(1)
					v286 = v284 + v285
					*(*int32)(unsafe.Add(mBase, uint32(v270)+12)) = v286
					v288 = *(*int32)(unsafe.Add(mBase, uint32(v270)+20))
					*(*int32)(unsafe.Add(mBase, uint32(v288+v286<<(uint(v279)%32)))) = int32(97)
					v294 = *(*int32)(unsafe.Add(mBase, uint32(v270)+20))
					v295 = *(*int32)(unsafe.Add(mBase, uint32(v270)+12))
					*(*int32)(unsafe.Add(mBase, uint32(v294+v295<<(uint(v279)%32))+4)) = int32(102)
					v301 = *(*int32)(unsafe.Add(mBase, uint32(v270)+12))
					v303 = v301 + v285
					*(*int32)(unsafe.Add(mBase, uint32(v270)+12)) = v303
					v305 = *(*int32)(unsafe.Add(mBase, uint32(v270)+20))
					*(*int32)(unsafe.Add(mBase, uint32(v305+v303<<(uint(v279)%32)))) = int32(65)
					v311 = *(*int32)(unsafe.Add(mBase, uint32(v270)+20))
					v312 = *(*int32)(unsafe.Add(mBase, uint32(v270)+12))
					*(*int32)(unsafe.Add(mBase, uint32(v311+v312<<(uint(v279)%32))+4)) = int32(70)
					v318 = *(*int32)(unsafe.Add(mBase, uint32(v270)+12))
					*(*int32)(unsafe.Add(mBase, uint32(v270)+12)) = v318 + v285
					return v270
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
				v359 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				if v359 != 0 {
					v361 = v359
				} else {
					v361 = int32(12)
				}
				*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v361
				v364 = int32(0)
			} else {
				v364 = v35
			}
			return v364
		}
	default:
		*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
		v359 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		if v359 != 0 {
			v361 = v359
		} else {
			v361 = int32(12)
		}
		*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v361
		v364 = int32(0)
		return v364
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
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
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
	var v104 int32
	_ = v104
	var v110 int64
	_ = v110
	var v116 int32
	_ = v116
	var v125 int32
	_ = v125
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
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
	var v186 int32
	_ = v186
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
													v49 = *(*int32)(unsafe.Add(mBase, uint32(v43)+12))
													v50 = *(*int32)(unsafe.Add(mBase, uint32(v43)+8))
													v51 = int32(*(*int16)(unsafe.Add(mBase, uint32(v43)+4)))
													if v51 < int32(0) {
													} else {
														v54 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
														v56 = v54 - int32(97)
														if base.Ui32(int32(17)) < base.Ui32(v56) {
														} else {
															if int32(1)<<(uint(v56)%32)&int32(163841) == int32(0) {
															} else {
																v65 = *(*int32)(unsafe.Add(mBase, uint32(v36)+80))
																if v65 != 0 {
																} else {
																	v66 = *(*int32)(unsafe.Add(mBase, uint32(v43)+36))
																	if v66 == int32(0) {
																		v69 = *(*int32)(unsafe.Add(mBase, uint32(v36)+52))
																		v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)+20))
																		v74 = *(*int32)(unsafe.Add(mBase, uint32(v43)+32))
																		*(*int32)(unsafe.Add(mBase, uint32(v70+v51*int32(24))+12)) = v74
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
													}
													v84 = *(*int32)(unsafe.Add(mBase, uint32(v43)+16))
													v85 = *(*int32)(unsafe.Add(mBase, uint32(v43)+20))
													if v85 == int32(0) {
														*(*int32)(unsafe.Add(mBase, uint32(v50)+20)) = v84
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v85)+16)) = v84
													}
													if v84 != 0 {
														*(*int32)(unsafe.Add(mBase, uint32(v84)+20)) = v85
													} else {
													}
													v91 = *(*int32)(unsafe.Add(mBase, uint32(v50)+12))
													*(*int32)(unsafe.Add(mBase, uint32(v50)+12)) = v91 - int32(1)
													v95 = *(*int32)(unsafe.Add(mBase, uint32(v43)+24))
													v96 = *(*int32)(unsafe.Add(mBase, uint32(v43)+28))
													if v96 == int32(0) {
														*(*int32)(unsafe.Add(mBase, uint32(v49)+16)) = v95
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v96)+24)) = v95
													}
													v102 = v43 + int32(8)
													if v95 != 0 {
														*(*int32)(unsafe.Add(mBase, uint32(v95)+28)) = v96
													} else {
													}
													v104 = *(*int32)(unsafe.Add(mBase, uint32(v49)+8))
													*(*int32)(unsafe.Add(mBase, uint32(v49)+8)) = v104 - int32(1)
													*(*int32)(unsafe.Add(mBase, uint32(v43))) = int32(0)
													v110 = int64(0)
													*(*int64)(unsafe.Add(mBase, uint32(v102)+16)) = v110
													*(*int64)(unsafe.Add(mBase, uint32(v102)+8)) = v110
													*(*int64)(unsafe.Add(mBase, uint32(v102))) = v110
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
													v131 = *(*int32)(unsafe.Add(mBase, uint32(v125)+12))
													v132 = *(*int32)(unsafe.Add(mBase, uint32(v125)+8))
													v133 = int32(*(*int16)(unsafe.Add(mBase, uint32(v125)+4)))
													if v133 < int32(0) {
													} else {
														v136 = *(*int32)(unsafe.Add(mBase, uint32(v125)))
														v138 = v136 - int32(97)
														if base.Ui32(int32(17)) < base.Ui32(v138) {
														} else {
															if int32(1)<<(uint(v138)%32)&int32(163841) == int32(0) {
															} else {
																v147 = *(*int32)(unsafe.Add(mBase, uint32(v36)+80))
																if v147 != 0 {
																} else {
																	v148 = *(*int32)(unsafe.Add(mBase, uint32(v125)+36))
																	if v148 == int32(0) {
																		v151 = *(*int32)(unsafe.Add(mBase, uint32(v36)+52))
																		v152 = *(*int32)(unsafe.Add(mBase, uint32(v151)+20))
																		v156 = *(*int32)(unsafe.Add(mBase, uint32(v125)+32))
																		*(*int32)(unsafe.Add(mBase, uint32(v152+v133*int32(24))+12)) = v156
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
													}
													v166 = *(*int32)(unsafe.Add(mBase, uint32(v125)+16))
													v167 = *(*int32)(unsafe.Add(mBase, uint32(v125)+20))
													if v167 == int32(0) {
														*(*int32)(unsafe.Add(mBase, uint32(v132)+20)) = v166
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v167)+16)) = v166
													}
													if v166 != 0 {
														*(*int32)(unsafe.Add(mBase, uint32(v166)+20)) = v167
													} else {
													}
													v173 = *(*int32)(unsafe.Add(mBase, uint32(v132)+12))
													*(*int32)(unsafe.Add(mBase, uint32(v132)+12)) = v173 - int32(1)
													v177 = *(*int32)(unsafe.Add(mBase, uint32(v125)+24))
													v178 = *(*int32)(unsafe.Add(mBase, uint32(v125)+28))
													if v178 == int32(0) {
														*(*int32)(unsafe.Add(mBase, uint32(v131)+16)) = v177
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v178)+24)) = v177
													}
													v184 = v125 + int32(8)
													if v177 != 0 {
														*(*int32)(unsafe.Add(mBase, uint32(v177)+28)) = v178
													} else {
													}
													v186 = *(*int32)(unsafe.Add(mBase, uint32(v131)+8))
													*(*int32)(unsafe.Add(mBase, uint32(v131)+8)) = v186 - int32(1)
													*(*int32)(unsafe.Add(mBase, uint32(v125))) = int32(0)
													v192 = int64(0)
													*(*int64)(unsafe.Add(mBase, uint32(v184)+16)) = v192
													*(*int64)(unsafe.Add(mBase, uint32(v184)+8)) = v192
													*(*int64)(unsafe.Add(mBase, uint32(v184))) = v192
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
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	if l0 == int32(0) {
		return int32(0)
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if v7 == int32(22) {
			return int32(1)
		} else {
			v13 = F_expression_tree_walker_impl(m, l0, int32(1047), l1)
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int32(0)
			} else {
				return v13
			}
		}
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
	var v46 int32
	_ = v46
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
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
						v83 = m.ExcPending
						if v83 != 0 {
							return
						} else {
							F_errcode(m, int32(50364548))
							mBase = m.M
							v86 = m.ExcPending
							if v86 != 0 {
								return
							} else {
								v96 = *(*int32)(unsafe.Add(mBase, _consts[297]))
								*(*int32)(unsafe.Add(mBase, uint32(v7))) = v96
								F_errmsg(m, int32(184362), v7)
								mBase = m.M
								v101 = m.ExcPending
								if v101 != 0 {
									return
								} else {
									v102 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
									v104 = F_locate_agg_of_level(m, v102, int32(0))
									mBase = m.M
									v105 = m.ExcPending
									if v105 != 0 {
										return
									} else {
										F_parser_errposition(m, l0, v104)
										mBase = m.M
										v107 = m.ExcPending
										if v107 != 0 {
											return
										} else {
											F_errfinish(m, int32(495469), int32(1965), int32(549966))
											mBase = m.M
											v112 = m.ExcPending
											if v112 != 0 {
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
										F_errcode(m, int32(655492))
										mBase = m.M
										v36 = m.ExcPending
										if v36 != 0 {
											return
										} else {
											v46 = *(*int32)(unsafe.Add(mBase, _consts[297]))
											*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v46
											F_errmsg(m, int32(184277), v7+int32(16))
											mBase = m.M
											v53 = m.ExcPending
											if v53 != 0 {
												return
											} else {
												v54 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
												v55 = F_locate_windowfunc(m, v54)
												mBase = m.M
												v56 = m.ExcPending
												if v56 != 0 {
													return
												} else {
													F_parser_errposition(m, l0, v55)
													mBase = m.M
													v58 = m.ExcPending
													if v58 != 0 {
														return
													} else {
														F_errfinish(m, int32(495469), int32(1974), int32(549966))
														mBase = m.M
														v63 = m.ExcPending
														if v63 != 0 {
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
								F_errcode(m, int32(655492))
								mBase = m.M
								v36 = m.ExcPending
								if v36 != 0 {
									return
								} else {
									v46 = *(*int32)(unsafe.Add(mBase, _consts[297]))
									*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v46
									F_errmsg(m, int32(184277), v7+int32(16))
									mBase = m.M
									v53 = m.ExcPending
									if v53 != 0 {
										return
									} else {
										v54 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
										v55 = F_locate_windowfunc(m, v54)
										mBase = m.M
										v56 = m.ExcPending
										if v56 != 0 {
											return
										} else {
											F_parser_errposition(m, l0, v55)
											mBase = m.M
											v58 = m.ExcPending
											if v58 != 0 {
												return
											} else {
												F_errfinish(m, int32(495469), int32(1974), int32(549966))
												mBase = m.M
												v63 = m.ExcPending
												if v63 != 0 {
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
			v67 = m.ExcPending
			if v67 != 0 {
				return
			} else {
				F_errmsg_internal(m, int32(549943), int32(0))
				mBase = m.M
				v71 = m.ExcPending
				if v71 != 0 {
					return
				} else {
					F_errfinish(m, int32(495469), int32(1983), int32(549966))
					mBase = m.M
					v76 = m.ExcPending
					if v76 != 0 {
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
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v192 int32
	_ = v192
	var v204 int32
	_ = v204
	var v208 int32
	_ = v208
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v275 int32
	_ = v275
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v299 int32
	_ = v299
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v310 int32
	_ = v310
	var v317 int32
	_ = v317
	var v320 int32
	_ = v320
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v338 int32
	_ = v338
	var v342 int32
	_ = v342
	var v345 int32
	_ = v345
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v354 int32
	_ = v354
	var v359 int32
	_ = v359
	var v363 int32
	_ = v363
	var v366 int32
	_ = v366
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v375 int32
	_ = v375
	var v380 int32
	_ = v380
	var v384 int32
	_ = v384
	var v387 int32
	_ = v387
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v400 int32
	_ = v400
	var v402 int32
	_ = v402
	var v407 int32
	_ = v407
	v3 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(80)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v17 == int32(9) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l1+v30)))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+76)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v15)+68)) = int64(4294967295)
	*(*int64)(unsafe.Add(mBase, uint32(v15)+60)) = int64(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+56)) = l0
	v47 = F_check_agg_arguments_walker(m, v37, v15+int32(56))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v30 = int32(68)
	v31 = v21
	v32 = v20
	v33 = int32(52)
	v34 = l1 + int32(32)
	goto L1
L3:
	;
	goto L4
L4:
	;
	v30 = int32(20)
	v31 = v3
	v32 = v3
	v33 = int32(16)
	v34 = l1 + int32(4)
	goto L1
L5:
	;
	return
L6:
	;
	v51 = F_check_agg_arguments_walker(m, v32, v15+int32(56))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v15)+64))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v15)+60))
	if v54 < int32(0) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v57 = int32(0)
	if v57 < v53 {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	if v54 < v53 {
		goto L14
	} else {
		goto L15
	}
L10:
	;
	if v53 == v66 {
		goto L20
	} else {
		goto L21
	}
L11:
	;
	v60 = v53
	goto L13
L12:
	;
	v60 = v57
	goto L13
L13:
	;
	v66 = v60
	goto L10
L14:
	;
	v62 = v54
	goto L16
L15:
	;
	v62 = v53
	goto L16
L16:
	;
	if v53 < int32(0) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v65 = v54
	goto L19
L18:
	;
	v65 = v62
	goto L19
L19:
	;
	v66 = v65
	goto L10
L20:
	;
	v68 = F_locate_agg_of_level(m, v37, v53)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L5
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v15)+68))
	v94 = int32(0)
	if base.B2i32(v94 <= v93)&base.B2i32(v93 < v66) == v94 {
		goto L36
	} else {
		goto L37
	}
L23:
	;
	if v68 < int32(0) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v72 = F_locate_agg_of_level(m, v32, v53)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L5
	} else {
		goto L27
	}
L25:
	;
	v74 = v68
	goto L26
L26:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L5
	} else {
		goto L28
	}
L27:
	;
	v74 = v72
	goto L26
L28:
	;
	F_errcode(m, int32(50364548))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L5
	} else {
		goto L29
	}
L29:
	;
	F_errmsg(m, int32(438680), int32(0))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L5
	} else {
		goto L30
	}
L30:
	;
	F_parser_errposition(m, l0, v74)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L5
	} else {
		goto L31
	}
L31:
	;
	F_errfinish(m, int32(495055), int32(693), int32(120199))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
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
	v384 = m.ExcPending
	if v384 != 0 {
		goto L5
	} else {
		goto L120
	}
L34:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v363 = m.ExcPending
	if v363 != 0 {
		goto L5
	} else {
		goto L114
	}
L35:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L5
	} else {
		goto L108
	}
L36:
	;
	if v31 == int32(0) {
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
	v317 = m.ExcPending
	if v317 != 0 {
		goto L5
	} else {
		goto L102
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1+v33))) = v66
	if v66 <= int32(0) {
		v192 = l0
		goto L46
	} else {
		goto L47
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+68)) = int32(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v15)+60)) = int64(-1)
	v108 = F_check_agg_arguments_walker(m, v31, v15+int32(56))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L5
	} else {
		goto L41
	}
L41:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v15)+60))
	if base.B2i32(int32(0) <= v110)&base.B2i32(v110 < v66) != 0 {
		goto L35
	} else {
		goto L42
	}
L42:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v15)+64))
	if base.B2i32(int32(0) <= v115)&base.B2i32(v115 <= v66) != 0 {
		goto L34
	} else {
		goto L43
	}
L43:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v15)+68))
	if v120 < int32(0) {
		goto L39
	} else {
		goto L44
	}
L44:
	;
	if v120 < v66 {
		goto L33
	} else {
		goto L45
	}
L45:
	;
	goto L39
L46:
	;
	v204 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v192)+92)) = uint8(v204)
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v192)+68))
	switch v208 - int32(2) {
	case 0, 1:
		v251 = int32(138788)
		v252 = int32(138843)
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
	v130 = v66 & int32(7)
	if v130 == int32(0) {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	if base.Ui32(v66) < base.Ui32(int32(8)) {
		v192 = v152
		goto L46
	} else {
		goto L55
	}
L49:
	;
	v152 = l0
	v153 = v66
	goto L48
L50:
	;
	goto L51
L51:
	;
	v134 = l0
	v135 = v66
	v136 = int32(0)
	goto L52
L52:
	;
	v146 = int32(1)
	v147 = v135 - v146
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v134)))
	v150 = v136 + v146
	if v150 != v130 {
		v134 = v148
		v135 = v147
		v136 = v150
		goto L52
	} else {
		goto L54
	}
L53:
	;
	v152 = v148
	v153 = v147
	goto L48
L54:
	;
	goto L53
L55:
	;
	v166 = v152
	v167 = v153
	goto L56
L56:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v166)))
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v180)))
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v181)))
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v182)))
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v183)))
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v184)))
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v185)))
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v186)))
	if base.Ui32(v167-int32(9)) < base.Ui32(int32(-2)) {
		v166 = v187
		v167 = v167 - int32(8)
		goto L56
	} else {
		goto L58
	}
L57:
	;
	v192 = v187
	goto L46
L58:
	;
	goto L57
L59:
	;
	m.G0 = v15 + int32(80)
	return
L60:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L5
	} else {
		goto L90
	}
L61:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L5
	} else {
		goto L82
	}
L62:
	;
	v251 = int32(139577)
	v252 = int32(139643)
	goto L61
L63:
	;
	v251 = int32(121899)
	v252 = int32(121953)
	goto L61
L64:
	;
	v251 = int32(145158)
	v252 = int32(145227)
	goto L61
L65:
	;
	v251 = int32(143684)
	v252 = int32(143749)
	goto L61
L66:
	;
	v251 = int32(422255)
	v252 = int32(422310)
	goto L61
L67:
	;
	v251 = int32(139025)
	v252 = int32(139088)
	goto L61
L68:
	;
	v251 = int32(132357)
	v252 = int32(132415)
	goto L61
L69:
	;
	v251 = int32(145523)
	v252 = int32(145584)
	goto L61
L70:
	;
	v251 = int32(144768)
	v252 = int32(144830)
	goto L61
L71:
	;
	v251 = int32(160007)
	v252 = int32(160063)
	goto L61
L72:
	;
	v251 = int32(144204)
	v252 = int32(144261)
	goto L61
L73:
	;
	v251 = int32(145935)
	v252 = int32(145994)
	goto L61
L74:
	;
	v251 = int32(119053)
	v252 = int32(119110)
	goto L61
L75:
	;
	v251 = int32(139322)
	v252 = int32(139383)
	goto L61
L76:
	;
	v251 = int32(519570)
	v252 = int32(519623)
	goto L61
L77:
	;
	v251 = int32(519002)
	v252 = int32(519053)
	goto L61
L78:
	;
	v251 = int32(538026)
	v252 = int32(538078)
	goto L61
L79:
	;
	v251 = int32(143973)
	v252 = int32(144031)
	goto L61
L80:
	;
	v251 = int32(528075)
	v252 = int32(528132)
	goto L61
L81:
	;
	v251 = int32(304281)
	v252 = int32(304357)
	goto L61
L82:
	;
	F_errcode(m, int32(50364548))
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L5
	} else {
		goto L83
	}
L83:
	;
	if v17 == int32(9) {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	v262 = v251
	goto L86
L85:
	;
	v262 = v252
	goto L86
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v262
	F_errmsg_internal(m, int32(205224), v15+int32(16))
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L5
	} else {
		goto L87
	}
L87:
	;
	F_parser_errposition(m, v192, v36)
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L5
	} else {
		goto L88
	}
L88:
	;
	F_errfinish(m, int32(495055), int32(600), int32(118766))
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
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
	v282 = m.ExcPending
	if v282 != 0 {
		goto L5
	} else {
		goto L91
	}
L91:
	;
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v192)+68))
	if base.Ui32(v283) <= base.Ui32(int32(44)) {
		goto L93
	} else {
		goto L94
	}
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v293
	if v17 == int32(9) {
		goto L96
	} else {
		goto L97
	}
L93:
	;
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v283<<(uint(int32(2))%32))+uint32(_consts[296])))
	v293 = v292
	goto L95
L94:
	;
	v293 = int32(422936)
	goto L95
L95:
	;
	goto L92
L96:
	;
	v299 = int32(184362)
	goto L98
L97:
	;
	v299 = int32(184404)
	goto L98
L98:
	;
	F_errmsg_internal(m, v299, v15+int32(32))
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L5
	} else {
		goto L99
	}
L99:
	;
	F_parser_errposition(m, v192, v36)
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L5
	} else {
		goto L100
	}
L100:
	;
	F_errfinish(m, int32(495055), int32(615), int32(118766))
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
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
	v320 = m.ExcPending
	if v320 != 0 {
		goto L5
	} else {
		goto L103
	}
L103:
	;
	F_errmsg(m, int32(534908), int32(0))
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L5
	} else {
		goto L104
	}
L104:
	;
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v15)+72))
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v325)+8))
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v326)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v327
	F_errdetail(m, int32(604016), v15)
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L5
	} else {
		goto L105
	}
L105:
	;
	F_parser_errposition(m, l0, v36)
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L5
	} else {
		goto L106
	}
L106:
	;
	F_errfinish(m, int32(495055), int32(708), int32(120199))
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
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
	v345 = m.ExcPending
	if v345 != 0 {
		goto L5
	} else {
		goto L109
	}
L109:
	;
	F_errmsg(m, int32(120516), int32(0))
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L5
	} else {
		goto L110
	}
L110:
	;
	v350 = *(*int32)(unsafe.Add(mBase, uint32(v15)+60))
	v351 = F_locate_var_of_level(m, v31, v350)
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L5
	} else {
		goto L111
	}
L111:
	;
	F_parser_errposition(m, l0, v351)
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L5
	} else {
		goto L112
	}
L112:
	;
	F_errfinish(m, int32(495055), int32(731), int32(120199))
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
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
	v366 = m.ExcPending
	if v366 != 0 {
		goto L5
	} else {
		goto L115
	}
L115:
	;
	F_errmsg(m, int32(438680), int32(0))
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
		goto L5
	} else {
		goto L116
	}
L116:
	;
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v15)+64))
	v372 = F_locate_agg_of_level(m, v31, v371)
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L5
	} else {
		goto L117
	}
L117:
	;
	F_parser_errposition(m, l0, v372)
	mBase = m.M
	v375 = m.ExcPending
	if v375 != 0 {
		goto L5
	} else {
		goto L118
	}
L118:
	;
	F_errfinish(m, int32(495055), int32(738), int32(120199))
	mBase = m.M
	v380 = m.ExcPending
	if v380 != 0 {
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
	v387 = m.ExcPending
	if v387 != 0 {
		goto L5
	} else {
		goto L121
	}
L121:
	;
	F_errmsg(m, int32(534908), int32(0))
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L5
	} else {
		goto L122
	}
L122:
	;
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v15)+72))
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v392)+8))
	v394 = *(*int32)(unsafe.Add(mBase, uint32(v393)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+48)) = v394
	F_errdetail(m, int32(604016), v15+int32(48))
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L5
	} else {
		goto L123
	}
L123:
	;
	F_parser_errposition(m, l0, v36)
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		goto L5
	} else {
		goto L124
	}
L124:
	;
	F_errfinish(m, int32(495055), int32(745), int32(120199))
	mBase = m.M
	v407 = m.ExcPending
	if v407 != 0 {
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
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
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
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v79 int32
	_ = v79
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
	v45 = int32(0)
	v48 = v33
	goto L15
L13:
	;
	v79 = v33
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
		v67 = v48
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v79 = v67
	goto L14
L17:
	;
	v69 = v45 + int32(1)
	if v69 != l4 {
		v45 = v69
		v48 = v67
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
	v67 = int32(0)
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
	v67 = v48
	goto L17
L24:
	;
	if v62 != 0 {
		v67 = v48
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
	return v79
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = l0
	F_errmsg_internal(m, int32(44572), v13)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	F_errfinish(m, int32(495399), int32(163), int32(360650))
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
		v8 = *(*int32)(unsafe.Add(mBase, _consts[155]))
		*(*int32)(unsafe.Add(mBase, _consts[250])) = v8
		v14 = F_format_elog_string(m, int32(428023), int32(0))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, _consts[422])) = v14
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
								F_errmsg(m, int32(464160), v9)
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
										F_errfinish(m, int32(494048), int32(305), int32(220609))
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
									F_errmsg(m, int32(464160), v9)
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
											F_errfinish(m, int32(494048), int32(305), int32(220609))
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
										F_errmsg(m, int32(464215), v9+int32(16))
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
												F_errfinish(m, int32(494048), int32(312), int32(220609))
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
					*(*int32)(unsafe.Add(mBase, _consts[249])) = v29
					v32 = int32(0)
					v34 = *(*int32)(unsafe.Add(mBase, _consts[155]))
					*(*int32)(unsafe.Add(mBase, _consts[250])) = v34
					v37 = *(*int32)(unsafe.Add(mBase, uint32(v6)+24))
					*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = v37
					v43 = F_format_elog_string(m, int32(205224), v6+int32(16))
					mBase = m.M
					v44 = m.ExcPending
					if v44 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, _consts[251])) = v43
						v46 = *(*int32)(unsafe.Add(mBase, uint32(v6)+20))
						if v46 == int32(0) {
							v60 = v32
							m.G0 = v6 + int32(32)
							return v60
						} else {
							v50 = *(*int32)(unsafe.Add(mBase, _consts[155]))
							*(*int32)(unsafe.Add(mBase, _consts[250])) = v50
							v53 = *(*int32)(unsafe.Add(mBase, uint32(v6)+20))
							*(*int32)(unsafe.Add(mBase, uint32(v6))) = v53
							v57 = F_format_elog_string(m, int32(205224), v6)
							mBase = m.M
							v58 = m.ExcPending
							if v58 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, _consts[252])) = v57
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
						v31 = *(*int32)(unsafe.Add(mBase, _consts[929]))
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
												F_errmsg(m, int32(190779), v8)
												mBase = m.M
												v64 = m.ExcPending
												if v64 != 0 {
													return
												} else {
													F_errhint(m, int32(630057), int32(0))
													mBase = m.M
													v68 = m.ExcPending
													if v68 != 0 {
														return
													} else {
														F_errfinish(m, int32(493891), int32(102), int32(357783))
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
											F_errmsg(m, int32(190779), v8)
											mBase = m.M
											v64 = m.ExcPending
											if v64 != 0 {
												return
											} else {
												F_errhint(m, int32(630057), int32(0))
												mBase = m.M
												v68 = m.ExcPending
												if v68 != 0 {
													return
												} else {
													F_errfinish(m, int32(493891), int32(102), int32(357783))
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
				v31 = *(*int32)(unsafe.Add(mBase, _consts[929]))
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
										F_errmsg(m, int32(190779), v8)
										mBase = m.M
										v64 = m.ExcPending
										if v64 != 0 {
											return
										} else {
											F_errhint(m, int32(630057), int32(0))
											mBase = m.M
											v68 = m.ExcPending
											if v68 != 0 {
												return
											} else {
												F_errfinish(m, int32(493891), int32(102), int32(357783))
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
									F_errmsg(m, int32(190779), v8)
									mBase = m.M
									v64 = m.ExcPending
									if v64 != 0 {
										return
									} else {
										F_errhint(m, int32(630057), int32(0))
										mBase = m.M
										v68 = m.ExcPending
										if v68 != 0 {
											return
										} else {
											F_errfinish(m, int32(493891), int32(102), int32(357783))
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
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	v4 = int32(1)
	v5 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v5 != v4 {
		v26 = v4
		return v26
	} else {
		v8 = int32(1)
		v10 = int32(*(*uint8)(unsafe.Add(mBase, _consts[852])))
		if v10 != v8 {
			v26 = v8
			return v26
		} else {
			v14 = *(*int32)(unsafe.Add(mBase, _consts[155]))
			*(*int32)(unsafe.Add(mBase, _consts[250])) = v14
			v20 = F_format_elog_string(m, int32(609811), int32(0))
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, _consts[251])) = v20
				v26 = int32(0)
				return v26
			}
		}
	}
}
func F_clear_setitimer(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v6 int64
	_ = v6
	var v14 int32
	_ = v14
	v2 = m.G0
	v3 = int32(32)
	v4 = v2 - v3
	m.G0 = v4
	v6 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v4)+24)) = v6
	*(*int64)(unsafe.Add(mBase, uint32(v4)+16)) = v6
	*(*int64)(unsafe.Add(mBase, uint32(v4)+8)) = v6
	*(*int64)(unsafe.Add(mBase, uint32(v4))) = v6
	v14 = F_setitimer(m, v4)
	mBase = m.M
	m.G0 = v4 + v3
	return
}
func F_clearerr(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	if int32(0) <= v2 {
		v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		*(*int32)(unsafe.Add(mBase, uint32(l0))) = v5 & int32(-49)
		return
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		*(*int32)(unsafe.Add(mBase, uint32(l0))) = v9 & int32(-49)
		return
	}
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
		v8 = int32(535553)
	} else {
		v8 = int32(0)
	}
	if v5 != 0 {
		v10 = v8
	} else {
		v10 = int32(538498)
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
	var v13 int32
	_ = v13
	var v14 int64
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int64
	_ = v38
	var v40 int64
	_ = v40
	var v41 int64
	_ = v41
	var v45 int32
	_ = v45
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
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+48)))
	v12 = v10 & int32(240)
	switch v12 {
	case 0:
		v35 = *(*int32)(unsafe.Add(mBase, _consts[134]))
		v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)+28))
		v37 = *(*int32)(unsafe.Add(mBase, uint32(v9)+64))
		v38 = *(*int64)(unsafe.Add(mBase, uint32(v37)))
		v40 = int64(*(*uint16)(unsafe.Add(mBase, _consts[135])))
		v41 = base.I64_rem_s(v38, v40)
		v45 = v36 + base.I32_wrap_i64(v41)<<(uint(int32(7))%32)
		v47 = F_LWLockAcquire(m, v45, int32(0))
		mBase = m.M
		v48 = m.ExcPending
		if v48 != 0 {
			return
		} else {
			v49 = int32(4381456)
			v51 = F_SimpleLruZeroPage(m, v49, v38)
			mBase = m.M
			v52 = m.ExcPending
			if v52 != 0 {
				return
			} else {
				F_SimpleLruWritePage(m, v49, v51)
				mBase = m.M
				v54 = m.ExcPending
				if v54 != 0 {
					return
				} else {
					F_LWLockRelease(m, v45)
					mBase = m.M
					v56 = m.ExcPending
					if v56 != 0 {
						return
					} else {
						m.G0 = v7 + int32(16)
						return
					}
				}
			}
		}
	default:
		F_errstart_cold(m, int32(23), int32(0))
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v7))) = v12
			F_errmsg_internal(m, int32(52500), v7)
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return
			} else {
				F_errfinish(m, int32(494953), int32(1142), int32(241410))
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	case 16:
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v9)+64))
		v14 = *(*int64)(unsafe.Add(mBase, uint32(v13)))
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
		F_AdvanceOldestClogXid(m, v15)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return
		} else {
			F_SimpleLruTruncate(m, int32(4381456), v14)
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return
			} else {
				m.G0 = v7 + int32(16)
				return
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
		*(*int32)(unsafe.Add(mBase, _consts[155])) = v6
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
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4))))
	v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3))))
	if v8 == int32(0) {
		v27 = v7
		v28 = v8
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v28 - v27
L2:
	;
	goto L1
L3:
	;
	if v7 != v8 {
		v27 = v7
		v28 = v8
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v12 = v3
	v13 = v4
	goto L5
L5:
	;
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+1)))
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+1)))
	if v17 == int32(0) {
		v27 = v16
		v28 = v17
		goto L2
	} else {
		goto L7
	}
L6:
	;
	v27 = v16
	v28 = v17
	goto L2
L7:
	;
	v20 = int32(1)
	if v16 == v17 {
		v12 = v12 + v20
		v13 = v13 + v20
		goto L5
	} else {
		goto L8
	}
L8:
	;
	goto L6
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
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v3 == int32(2) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v44
L2:
	;
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v7 == v8 {
		v44 = int32(0)
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
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
	if v20 == int32(0) {
		v39 = v19
		v40 = v20
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
	v44 = v40 - v39
	goto L1
L10:
	;
	goto L9
L11:
	;
	if v19 != v20 {
		v39 = v19
		v40 = v20
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v24 = v15
	v25 = v16
	goto L13
L13:
	;
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+1)))
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+1)))
	if v29 == int32(0) {
		v39 = v28
		v40 = v29
		goto L10
	} else {
		goto L15
	}
L14:
	;
	v39 = v28
	v40 = v29
	goto L10
L15:
	;
	v32 = int32(1)
	if v28 == v29 {
		v24 = v24 + v32
		v25 = v25 + v32
		goto L13
	} else {
		goto L16
	}
L16:
	;
	goto L14
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
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6))))
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4))))
	if v10 == int32(0) {
		v29 = v9
		v30 = v10
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v30 - v29
L2:
	;
	goto L1
L3:
	;
	if v9 != v10 {
		v29 = v9
		v30 = v10
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v14 = v4
	v15 = v6
	goto L5
L5:
	;
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+1)))
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+1)))
	if v19 == int32(0) {
		v29 = v18
		v30 = v19
		goto L2
	} else {
		goto L7
	}
L6:
	;
	v29 = v18
	v30 = v19
	goto L2
L7:
	;
	v22 = int32(1)
	if v18 == v19 {
		v14 = v14 + v22
		v15 = v15 + v22
		goto L5
	} else {
		goto L8
	}
L8:
	;
	goto L6
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
	v5 = *(*int32)(unsafe.Add(mBase, _consts[1394]))
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
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+17)))
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+17)))
	if v73 == v74 {
		v85 = int32(0)
		goto L1
	} else {
		goto L24
	}
L24:
	;
	if v74 != 0 {
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
	var v6 int32
	_ = v6
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
	var v22 int32
	_ = v22
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v10 = F_FunctionCall2Coll(m, v6, v7, v8, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		if v10 != 0 {
			v22 = int32(-1)
			return v22
		} else {
			v14 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v18 = F_FunctionCall2Coll(m, v14, v15, v16, v17)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				v22 = base.B2i32(v18 != int32(0))
				return v22
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
	var v68 int32
	_ = v68
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
					v68 = v65
					return v68
				}
			} else {
				v68 = v4
				return v68
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
									F_errmsg(m, int32(201120), v10)
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
											F_errfinish(m, int32(492135), int32(475), int32(217230))
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
											F_errmsg(m, int32(201120), v10)
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
													F_errfinish(m, int32(492135), int32(475), int32(217230))
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
									v68 = v65
									return v68
								}
							} else {
								v68 = v4
								return v68
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
	var v41 int32
	_ = v41
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
	var v69 int32
	_ = v69
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
	var v190 int32
	_ = v190
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
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
	v206 = m.ExcPending
	if v206 != 0 {
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
	v39 = int32(727618)
	goto L13
L12:
	;
	v39 = int32(651469)
	goto L13
L13:
	;
	v41 = v21 + int32(20)
	v43 = int32(0)
	v47 = v31
	v49 = int32(0)
	goto L14
L14:
	;
	v60 = v41 + v47<<(uint(int32(4))%32) + v43*int32(100)
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+91)))
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
	v190 = v43 + int32(1)
	goto L18
L17:
	;
	if v49&int32(1) != 0 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	if v190 < v186 {
		v43 = v190
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
	v69 = m.ExcPending
	if v69 != 0 {
		goto L1
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	F_escape_json(m, l1, v60+int32(4))
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
	v97 = v41 + v43<<(uint(int32(4))%32)
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
	F_errmsg_internal(m, int32(480197), v15)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	F_errfinish(m, int32(324700), int32(70), int32(67479))
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
	F_datum_to_json_internal(m, v152, v176&v177, l1, v175, v174, int32(0))
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
	v174 = v156
	v175 = v156
	v176 = int32(1)
	goto L49
L51:
	;
	goto L52
L52:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v60)+68))
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
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v15)+16))
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v15)+20))
	v173 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+27)))
	v174 = v171
	v175 = v172
	v176 = v173
	goto L49
L54:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	v186 = v183
	v187 = v177
	v190 = v78
	goto L18
L55:
	;
	goto L15
L56:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
	if int32(0) <= v207 {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	F_DecrTupleDescRefCount(m, v21)
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
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
	var v39 float64
	_ = v39
	var v49 float64
	_ = v49
	var v51 float64
	_ = v51
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v74 float64
	_ = v74
	var v75 float64
	_ = v75
	var v79 float64
	_ = v79
	var v84 float64
	_ = v84
	var v94 float64
	_ = v94
	var v96 float64
	_ = v96
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 float64
	_ = v125
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 float64
	_ = v130
	var v131 float64
	_ = v131
	var v133 float64
	_ = v133
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 float64
	_ = v139
	var v140 float64
	_ = v140
	var v142 float64
	_ = v142
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 float64
	_ = v148
	var v150 float64
	_ = v150
	var v152 float64
	_ = v152
	var v154 float64
	_ = v154
	var v156 float64
	_ = v156
	var v167 int32
	_ = v167
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
			v156 = v21
			m.G0 = v11 + int32(16)
			return v156
		}
	} else {
		v22 = *(*float64)(unsafe.Add(mBase, uint32(l2)))
		v23 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
		if base.F64_le(v22, v23) == int32(0) {
			v74 = *(*float64)(unsafe.Add(mBase, uint32(l2)+8))
			v75 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
			if base.F64_le(v74, v75) == int32(0) {
				v123 = F_DirectFunctionCall2Coll(m, int32(112), int32(0), l2, l1+int32(16))
				mBase = m.M
				v124 = m.ExcPending
				if v124 != 0 {
					return float64(0)
				} else {
					v125 = *(*float64)(unsafe.Add(mBase, uint32(v123)))
					v128 = F_DirectFunctionCall2Coll(m, int32(112), int32(0), l2, l1)
					mBase = m.M
					v129 = m.ExcPending
					if v129 != 0 {
						return float64(0)
					} else {
						v130 = *(*float64)(unsafe.Add(mBase, uint32(v128)))
						v131 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
						*(*float64)(unsafe.Add(mBase, uint32(v11))) = v131
						v133 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
						*(*float64)(unsafe.Add(mBase, uint32(v11)+8)) = v133
						v137 = F_DirectFunctionCall2Coll(m, int32(112), int32(0), l2, v11)
						mBase = m.M
						v138 = m.ExcPending
						if v138 != 0 {
							return float64(0)
						} else {
							v139 = *(*float64)(unsafe.Add(mBase, uint32(v137)))
							v140 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
							*(*float64)(unsafe.Add(mBase, uint32(v11))) = v140
							v142 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
							*(*float64)(unsafe.Add(mBase, uint32(v11)+8)) = v142
							v146 = F_DirectFunctionCall2Coll(m, int32(112), int32(0), l2, v11)
							mBase = m.M
							v147 = m.ExcPending
							if v147 != 0 {
								return float64(0)
							} else {
								v148 = *(*float64)(unsafe.Add(mBase, uint32(v146)))
								if base.F64_lt(v130, v125) != 0 {
									v150 = v130
								} else {
									v150 = v125
								}
								if base.F64_gt(v150, v139) != 0 {
									v152 = v139
								} else {
									v152 = v150
								}
								if base.F64_gt(v152, v148) != 0 {
									v154 = v148
								} else {
									v154 = v152
								}
								v156 = v154
								m.G0 = v11 + int32(16)
								return v156
							}
						}
					}
				}
			} else {
				v79 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
				if base.F64_ge(v74, v79) == int32(0) {
					v123 = F_DirectFunctionCall2Coll(m, int32(112), int32(0), l2, l1+int32(16))
					mBase = m.M
					v124 = m.ExcPending
					if v124 != 0 {
						return float64(0)
					} else {
						v125 = *(*float64)(unsafe.Add(mBase, uint32(v123)))
						v128 = F_DirectFunctionCall2Coll(m, int32(112), int32(0), l2, l1)
						mBase = m.M
						v129 = m.ExcPending
						if v129 != 0 {
							return float64(0)
						} else {
							v130 = *(*float64)(unsafe.Add(mBase, uint32(v128)))
							v131 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
							*(*float64)(unsafe.Add(mBase, uint32(v11))) = v131
							v133 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
							*(*float64)(unsafe.Add(mBase, uint32(v11)+8)) = v133
							v137 = F_DirectFunctionCall2Coll(m, int32(112), int32(0), l2, v11)
							mBase = m.M
							v138 = m.ExcPending
							if v138 != 0 {
								return float64(0)
							} else {
								v139 = *(*float64)(unsafe.Add(mBase, uint32(v137)))
								v140 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
								*(*float64)(unsafe.Add(mBase, uint32(v11))) = v140
								v142 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
								*(*float64)(unsafe.Add(mBase, uint32(v11)+8)) = v142
								v146 = F_DirectFunctionCall2Coll(m, int32(112), int32(0), l2, v11)
								mBase = m.M
								v147 = m.ExcPending
								if v147 != 0 {
									return float64(0)
								} else {
									v148 = *(*float64)(unsafe.Add(mBase, uint32(v146)))
									if base.F64_lt(v130, v125) != 0 {
										v150 = v130
									} else {
										v150 = v125
									}
									if base.F64_gt(v150, v139) != 0 {
										v152 = v139
									} else {
										v152 = v150
									}
									if base.F64_gt(v152, v148) != 0 {
										v154 = v148
									} else {
										v154 = v152
									}
									v156 = v154
									m.G0 = v11 + int32(16)
									return v156
								}
							}
						}
					}
				} else {
					if base.F64_gt(v22, v23) != 0 {
						v84 = base.F64_sub(v22, v23)
						if base.F64_ne(base.F64_abs(v84), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
							v156 = v84
							m.G0 = v11 + int32(16)
							return v156
						} else {
							if base.F64_eq(base.F64_abs(v22), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
								v156 = v84
								m.G0 = v11 + int32(16)
								return v156
							} else {
								if base.F64_ne(base.F64_abs(v23), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
									F_float_overflow_error(m)
									mBase = m.M
									v167 = m.ExcPending
									if v167 != 0 {
										return float64(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								} else {
									v156 = v84
									m.G0 = v11 + int32(16)
									return v156
								}
							}
						}
					} else {
						v94 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
						if base.F64_gt(v94, v22) != 0 {
							v96 = base.F64_sub(v94, v22)
							if base.F64_ne(base.F64_abs(v96), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
								v156 = v96
								m.G0 = v11 + int32(16)
								return v156
							} else {
								if base.F64_eq(base.F64_abs(v22), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
									v156 = v96
									m.G0 = v11 + int32(16)
									return v156
								} else {
									if base.F64_ne(base.F64_abs(v94), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
										F_float_overflow_error(m)
										mBase = m.M
										v167 = m.ExcPending
										if v167 != 0 {
											return float64(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									} else {
										v156 = v96
										m.G0 = v11 + int32(16)
										return v156
									}
								}
							}
						} else {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v109 = m.ExcPending
							if v109 != 0 {
								return float64(0)
							} else {
								F_errmsg_internal(m, int32(157009), int32(0))
								mBase = m.M
								v113 = m.ExcPending
								if v113 != 0 {
									return float64(0)
								} else {
									F_errfinish(m, int32(496635), int32(1256), int32(415189))
									mBase = m.M
									v118 = m.ExcPending
									if v118 != 0 {
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
				v74 = *(*float64)(unsafe.Add(mBase, uint32(l2)+8))
				v75 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
				if base.F64_le(v74, v75) == int32(0) {
					v123 = F_DirectFunctionCall2Coll(m, int32(112), int32(0), l2, l1+int32(16))
					mBase = m.M
					v124 = m.ExcPending
					if v124 != 0 {
						return float64(0)
					} else {
						v125 = *(*float64)(unsafe.Add(mBase, uint32(v123)))
						v128 = F_DirectFunctionCall2Coll(m, int32(112), int32(0), l2, l1)
						mBase = m.M
						v129 = m.ExcPending
						if v129 != 0 {
							return float64(0)
						} else {
							v130 = *(*float64)(unsafe.Add(mBase, uint32(v128)))
							v131 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
							*(*float64)(unsafe.Add(mBase, uint32(v11))) = v131
							v133 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
							*(*float64)(unsafe.Add(mBase, uint32(v11)+8)) = v133
							v137 = F_DirectFunctionCall2Coll(m, int32(112), int32(0), l2, v11)
							mBase = m.M
							v138 = m.ExcPending
							if v138 != 0 {
								return float64(0)
							} else {
								v139 = *(*float64)(unsafe.Add(mBase, uint32(v137)))
								v140 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
								*(*float64)(unsafe.Add(mBase, uint32(v11))) = v140
								v142 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
								*(*float64)(unsafe.Add(mBase, uint32(v11)+8)) = v142
								v146 = F_DirectFunctionCall2Coll(m, int32(112), int32(0), l2, v11)
								mBase = m.M
								v147 = m.ExcPending
								if v147 != 0 {
									return float64(0)
								} else {
									v148 = *(*float64)(unsafe.Add(mBase, uint32(v146)))
									if base.F64_lt(v130, v125) != 0 {
										v150 = v130
									} else {
										v150 = v125
									}
									if base.F64_gt(v150, v139) != 0 {
										v152 = v139
									} else {
										v152 = v150
									}
									if base.F64_gt(v152, v148) != 0 {
										v154 = v148
									} else {
										v154 = v152
									}
									v156 = v154
									m.G0 = v11 + int32(16)
									return v156
								}
							}
						}
					}
				} else {
					v79 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
					if base.F64_ge(v74, v79) == int32(0) {
						v123 = F_DirectFunctionCall2Coll(m, int32(112), int32(0), l2, l1+int32(16))
						mBase = m.M
						v124 = m.ExcPending
						if v124 != 0 {
							return float64(0)
						} else {
							v125 = *(*float64)(unsafe.Add(mBase, uint32(v123)))
							v128 = F_DirectFunctionCall2Coll(m, int32(112), int32(0), l2, l1)
							mBase = m.M
							v129 = m.ExcPending
							if v129 != 0 {
								return float64(0)
							} else {
								v130 = *(*float64)(unsafe.Add(mBase, uint32(v128)))
								v131 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
								*(*float64)(unsafe.Add(mBase, uint32(v11))) = v131
								v133 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
								*(*float64)(unsafe.Add(mBase, uint32(v11)+8)) = v133
								v137 = F_DirectFunctionCall2Coll(m, int32(112), int32(0), l2, v11)
								mBase = m.M
								v138 = m.ExcPending
								if v138 != 0 {
									return float64(0)
								} else {
									v139 = *(*float64)(unsafe.Add(mBase, uint32(v137)))
									v140 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
									*(*float64)(unsafe.Add(mBase, uint32(v11))) = v140
									v142 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
									*(*float64)(unsafe.Add(mBase, uint32(v11)+8)) = v142
									v146 = F_DirectFunctionCall2Coll(m, int32(112), int32(0), l2, v11)
									mBase = m.M
									v147 = m.ExcPending
									if v147 != 0 {
										return float64(0)
									} else {
										v148 = *(*float64)(unsafe.Add(mBase, uint32(v146)))
										if base.F64_lt(v130, v125) != 0 {
											v150 = v130
										} else {
											v150 = v125
										}
										if base.F64_gt(v150, v139) != 0 {
											v152 = v139
										} else {
											v152 = v150
										}
										if base.F64_gt(v152, v148) != 0 {
											v154 = v148
										} else {
											v154 = v152
										}
										v156 = v154
										m.G0 = v11 + int32(16)
										return v156
									}
								}
							}
						}
					} else {
						if base.F64_gt(v22, v23) != 0 {
							v84 = base.F64_sub(v22, v23)
							if base.F64_ne(base.F64_abs(v84), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
								v156 = v84
								m.G0 = v11 + int32(16)
								return v156
							} else {
								if base.F64_eq(base.F64_abs(v22), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
									v156 = v84
									m.G0 = v11 + int32(16)
									return v156
								} else {
									if base.F64_ne(base.F64_abs(v23), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
										F_float_overflow_error(m)
										mBase = m.M
										v167 = m.ExcPending
										if v167 != 0 {
											return float64(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									} else {
										v156 = v84
										m.G0 = v11 + int32(16)
										return v156
									}
								}
							}
						} else {
							v94 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
							if base.F64_gt(v94, v22) != 0 {
								v96 = base.F64_sub(v94, v22)
								if base.F64_ne(base.F64_abs(v96), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
									v156 = v96
									m.G0 = v11 + int32(16)
									return v156
								} else {
									if base.F64_eq(base.F64_abs(v22), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
										v156 = v96
										m.G0 = v11 + int32(16)
										return v156
									} else {
										if base.F64_ne(base.F64_abs(v94), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
											F_float_overflow_error(m)
											mBase = m.M
											v167 = m.ExcPending
											if v167 != 0 {
												return float64(0)
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										} else {
											v156 = v96
											m.G0 = v11 + int32(16)
											return v156
										}
									}
								}
							} else {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v109 = m.ExcPending
								if v109 != 0 {
									return float64(0)
								} else {
									F_errmsg_internal(m, int32(157009), int32(0))
									mBase = m.M
									v113 = m.ExcPending
									if v113 != 0 {
										return float64(0)
									} else {
										F_errfinish(m, int32(496635), int32(1256), int32(415189))
										mBase = m.M
										v118 = m.ExcPending
										if v118 != 0 {
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
						v156 = float64(0)
						m.G0 = v11 + int32(16)
						return v156
					} else {
						v36 = *(*float64)(unsafe.Add(mBase, uint32(l2)+8))
						v37 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
						if base.F64_gt(v36, v37) != 0 {
							v39 = base.F64_sub(v36, v37)
							if base.F64_ne(base.F64_abs(v39), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
								v156 = v39
								m.G0 = v11 + int32(16)
								return v156
							} else {
								if base.F64_eq(base.F64_abs(v36), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
									v156 = v39
									m.G0 = v11 + int32(16)
									return v156
								} else {
									if base.F64_eq(base.F64_abs(v37), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
										v156 = v39
										m.G0 = v11 + int32(16)
										return v156
									} else {
										F_float_overflow_error(m)
										mBase = m.M
										v167 = m.ExcPending
										if v167 != 0 {
											return float64(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							}
						} else {
							v49 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
							if base.F64_gt(v49, v36) != 0 {
								v51 = base.F64_sub(v49, v36)
								if base.F64_ne(base.F64_abs(v51), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
									v156 = v51
									m.G0 = v11 + int32(16)
									return v156
								} else {
									if base.F64_eq(base.F64_abs(v36), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
										v156 = v51
										m.G0 = v11 + int32(16)
										return v156
									} else {
										if base.F64_ne(base.F64_abs(v49), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
											F_float_overflow_error(m)
											mBase = m.M
											v167 = m.ExcPending
											if v167 != 0 {
												return float64(0)
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										} else {
											v156 = v51
											m.G0 = v11 + int32(16)
											return v156
										}
									}
								}
							} else {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v64 = m.ExcPending
								if v64 != 0 {
									return float64(0)
								} else {
									F_errmsg_internal(m, int32(157009), int32(0))
									mBase = m.M
									v68 = m.ExcPending
									if v68 != 0 {
										return float64(0)
									} else {
										F_errfinish(m, int32(496635), int32(1245), int32(415189))
										mBase = m.M
										v73 = m.ExcPending
										if v73 != 0 {
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
						v39 = base.F64_sub(v36, v37)
						if base.F64_ne(base.F64_abs(v39), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
							v156 = v39
							m.G0 = v11 + int32(16)
							return v156
						} else {
							if base.F64_eq(base.F64_abs(v36), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
								v156 = v39
								m.G0 = v11 + int32(16)
								return v156
							} else {
								if base.F64_eq(base.F64_abs(v37), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
									v156 = v39
									m.G0 = v11 + int32(16)
									return v156
								} else {
									F_float_overflow_error(m)
									mBase = m.M
									v167 = m.ExcPending
									if v167 != 0 {
										return float64(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						}
					} else {
						v49 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
						if base.F64_gt(v49, v36) != 0 {
							v51 = base.F64_sub(v49, v36)
							if base.F64_ne(base.F64_abs(v51), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
								v156 = v51
								m.G0 = v11 + int32(16)
								return v156
							} else {
								if base.F64_eq(base.F64_abs(v36), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
									v156 = v51
									m.G0 = v11 + int32(16)
									return v156
								} else {
									if base.F64_ne(base.F64_abs(v49), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
										F_float_overflow_error(m)
										mBase = m.M
										v167 = m.ExcPending
										if v167 != 0 {
											return float64(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									} else {
										v156 = v51
										m.G0 = v11 + int32(16)
										return v156
									}
								}
							}
						} else {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v64 = m.ExcPending
							if v64 != 0 {
								return float64(0)
							} else {
								F_errmsg_internal(m, int32(157009), int32(0))
								mBase = m.M
								v68 = m.ExcPending
								if v68 != 0 {
									return float64(0)
								} else {
									F_errfinish(m, int32(496635), int32(1245), int32(415189))
									mBase = m.M
									v73 = m.ExcPending
									if v73 != 0 {
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
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	v2 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v14 == v2 {
		v34 = v2
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v46 = F_palloc(m, v34<<(uint(int32(1))%32)+int32(8194))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L7
	} else {
		goto L8
	}
L2:
	;
	if l0 == v14 {
		v34 = v2
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v19 = v2
	v20 = v14
	goto L4
L4:
	;
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+8)))
	v30 = v19 + base.B2i32(v27 != int32(0))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	if v31 != l0 {
		v19 = v30
		v20 = v31
		goto L4
	} else {
		goto L6
	}
L5:
	;
	v34 = v30
	goto L1
L6:
	;
	goto L5
L7:
	;
	return
L8:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v46))) = uint16(v34)
	v50 = v46 + int32(2)
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v51 == int32(0) {
		v151 = v50
		goto L9
	} else {
		goto L10
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v151 - v46
	m.G0 = v12 + int32(16)
	return
L10:
	;
	if l0 == v51 {
		v151 = v50
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v56 = v50
	v57 = v51
	v59 = v2
	goto L12
L12:
	;
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57)+8)))
	switch v64 {
	case 0:
		goto L18
	case 1:
		goto L16
	default:
		goto L17
	}
L13:
	;
	v151 = v143
	goto L9
L14:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
	if v148 != l0 {
		v56 = v143
		v57 = v148
		v59 = v147
		goto L12
	} else {
		goto L39
	}
L15:
	;
	v143 = v138 + v56 + int32(2)
	v147 = v135 + v59
	goto L14
L16:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v56))) = uint8(v59)
	v131 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v56)+1)) = uint8(v131)
	v135 = v131
	v138 = int32(0)
	goto L15
L17:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v57)+20))
	v68 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v67)+6)))
	v72 = (v68 + int32(1)) & int32(131070)
	v74 = v72 + int32(8)
	if v64 == int32(4) {
		goto L21
	} else {
		goto L22
	}
L18:
	;
	v143 = v56
	v147 = v59 + int32(1)
	goto L14
L19:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L7
	} else {
		goto L36
	}
L20:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v57)+20))
	if v74 != 0 {
		goto L33
	} else {
		goto L34
	}
L21:
	;
	v77 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v57)+16)))
	*(*uint8)(unsafe.Add(mBase, uint32(v56))) = uint8(v59)
	if base.Ui32(v77*int32(6)) <= base.Ui32(v74) {
		goto L24
	} else {
		goto L25
	}
L22:
	;
	goto L23
L23:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v56))) = uint8(v59)
	*(*uint8)(unsafe.Add(mBase, uint32(v56)+1)) = uint8(v64)
	if v64&int32(254) != int32(2) {
		goto L19
	} else {
		goto L31
	}
L24:
	;
	v83 = int32(4)
	*(*uint8)(unsafe.Add(mBase, uint32(v56)+1)) = uint8(v83)
	v85 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v57)+16)))
	*(*uint16)(unsafe.Add(mBase, uint32(v56)+2)) = uint16(v85)
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v57)+12))
	v91 = v85 * int32(6)
	if v91 != 0 {
		goto L28
	} else {
		goto L29
	}
L25:
	;
	goto L26
L26:
	;
	v96 = int32(3)
	*(*uint8)(unsafe.Add(mBase, uint32(v56)+1)) = uint8(v96)
	v105 = v96
	goto L20
L27:
	;
	v135 = int32(1)
	v138 = v91 + int32(2)
	goto L15
L28:
	;
	v92 = F__emscripten_memcpy_bulkmem(m, v56+v83, v89, v91)
	mBase = m.M
	goto L30
L29:
	;
	goto L30
L30:
	;
	goto L27
L31:
	;
	v105 = v64
	goto L20
L32:
	;
	v135 = base.B2i32(v105 != int32(2))
	v138 = (v72 + int32(9)) & int32(262142)
	goto L15
L33:
	;
	v109 = F__emscripten_memcpy_bulkmem(m, v56+int32(2), v108, v74)
	mBase = m.M
	goto L35
L34:
	;
	goto L35
L35:
	;
	goto L32
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v64
	F_errmsg_internal(m, int32(469919), v12)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L7
	} else {
		goto L37
	}
L37:
	;
	F_errfinish(m, int32(496162), int32(955), int32(502209))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L7
	} else {
		goto L38
	}
L38:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L39:
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
	v7 = m.Env.X__syscall_connect(m, l0, int32(4085574), int32(12), v4, v4, v4)
	mBase = m.M
	if base.Ui32(int32(-4095)) <= base.Ui32(v7) {
		*(*int32)(unsafe.Add(mBase, _consts[155])) = int32(0) - v7
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
	var v37 int32
	_ = v37
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
	var v78 int32
	_ = v78
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
	var v141 int32
	_ = v141
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v164 int32
	_ = v164
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v185 int32
	_ = v185
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v202 int32
	_ = v202
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v243 int32
	_ = v243
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v266 int32
	_ = v266
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v329 int32
	_ = v329
	var v351 int32
	_ = v351
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
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
	v37 = v11
	goto L4
L4:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l7)+12))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v46+v37<<(uint(int32(2))%32))))
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
	v372 = v37 + int32(1)
	v373 = *(*int32)(unsafe.Add(mBase, uint32(l7)+4))
	if v372 < v373 {
		v37 = v372
		goto L4
	} else {
		goto L86
	}
L11:
	;
	F_get_join_index_paths(m, l0, l1, l2, l3, l4, l5, l6, v54, l9)
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L6
	} else {
		goto L85
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
	v78 = int32(0)
	goto L14
L14:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l9)))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)+12))
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v86+v78<<(uint(int32(2))%32))))
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
	v329 = v78 + int32(1)
	if v329 != v62 {
		v78 = v329
		goto L14
	} else {
		goto L84
	}
L17:
	;
	if v185 != int32(3) {
		goto L16
	} else {
		goto L53
	}
L18:
	;
	v185 = base.B2i32(v90 != int32(0))
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
	v185 = int32(2)
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
	v185 = int32(3)
	goto L17
L31:
	;
	v185 = v171
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
		goto L43
	} else {
		goto L44
	}
L34:
	;
	v157 = v121 + int32(1)
	if v157 != v113 {
		v120 = v154
		v121 = v157
		goto L32
	} else {
		goto L42
	}
L35:
	;
	v141 = int32(3)
	if v120 == int32(1) {
		v171 = v141
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
		v154 = v120
		goto L34
	} else {
		goto L40
	}
L38:
	;
	if v135&(v133^int32(-1)) != 0 {
		v171 = v141
		goto L31
	} else {
		goto L39
	}
L39:
	;
	v154 = int32(2)
	goto L34
L40:
	;
	if v120 == int32(2) {
		goto L30
	} else {
		goto L41
	}
L41:
	;
	v154 = int32(1)
	goto L34
L42:
	;
	goto L33
L43:
	;
	if v154 == int32(1) {
		goto L46
	} else {
		goto L47
	}
L44:
	;
	goto L45
L45:
	;
	if v108 <= v107 {
		v171 = v154
		goto L31
	} else {
		goto L49
	}
L46:
	;
	v164 = int32(3)
	goto L48
L47:
	;
	v164 = int32(2)
	goto L48
L48:
	;
	v185 = v164
	goto L17
L49:
	;
	if v154 == int32(2) {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v170 = int32(3)
	goto L52
L51:
	;
	v170 = int32(1)
	goto L52
L52:
	;
	v171 = v170
	goto L31
L53:
	;
	if v52 == int32(0) {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v300 = *(*int32)(unsafe.Add(mBase, uint32(l9)))
	if v300 != 0 {
		goto L78
	} else {
		goto L79
	}
L55:
	;
	v190 = int32(0)
	v191 = *(*int32)(unsafe.Add(mBase, uint32(l7)+4))
	if v191 <= v190 {
		goto L54
	} else {
		goto L56
	}
L56:
	;
	v202 = v190
	v209 = v191
	goto L57
L57:
	;
	v213 = *(*int32)(unsafe.Add(mBase, uint32(l7)+12))
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v213+v202<<(uint(int32(2))%32))))
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v217)+4))
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v218)+60))
	if v52 == v219 {
		goto L59
	} else {
		goto L60
	}
L58:
	;
	goto L54
L59:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v218)+28))
	v222 = int32(0)
	if v221 == v222 {
		goto L63
	} else {
		goto L64
	}
L60:
	;
	v277 = v209
	goto L61
L61:
	;
	v279 = v202 + int32(1)
	if v279 < v277 {
		v202 = v279
		v209 = v277
		goto L57
	} else {
		goto L77
	}
L62:
	;
	if v275 != 0 {
		goto L16
	} else {
		goto L76
	}
L63:
	;
	v275 = int32(1)
	goto L62
L64:
	;
	goto L65
L65:
	;
	if v90 == int32(0) {
		v266 = v222
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v275 = v266
	goto L62
L67:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v221)+4))
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v90)+4))
	if v232 < v231 {
		v266 = v222
		goto L66
	} else {
		goto L68
	}
L68:
	;
	v234 = int32(1)
	if v231 <= v234 {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v237 = v234
	goto L71
L70:
	;
	v237 = v231
	goto L71
L71:
	;
	v238 = int32(8)
	v243 = int32(0)
	goto L72
L72:
	;
	v250 = v243 << (uint(int32(2)) % 32)
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v221+v238+v250)))
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v250+(v90+v238))))
	v257 = v252 & (v254 ^ int32(-1))
	v259 = base.B2i32(v257 == int32(0))
	if v257 != 0 {
		v266 = v259
		goto L66
	} else {
		goto L74
	}
L73:
	;
	v266 = v259
	goto L66
L74:
	;
	v261 = v243 + int32(1)
	if v261 != v237 {
		v243 = v261
		goto L72
	} else {
		goto L75
	}
L75:
	;
	goto L73
L76:
	;
	v276 = *(*int32)(unsafe.Add(mBase, uint32(l7)+4))
	v277 = v276
	goto L61
L77:
	;
	goto L58
L78:
	;
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v300)+4))
	v303 = v301
	goto L80
L79:
	;
	v303 = int32(0)
	goto L80
L80:
	;
	if l8*int32(10) <= v303 {
		goto L11
	} else {
		goto L81
	}
L81:
	;
	v305 = F_bms_union(m, v54, v90)
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L6
	} else {
		goto L82
	}
L82:
	;
	F_get_join_index_paths(m, l0, l1, l2, l3, l4, l5, l6, v305, l9)
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L6
	} else {
		goto L83
	}
L83:
	;
	goto L16
L84:
	;
	goto L15
L85:
	;
	goto L10
L86:
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
	var v492 int32
	_ = v492
	if base.Ui32(l0) < base.Ui32(int32(128)) {
		v45 = l0
	} else {
		if base.Ui32(l0) <= base.Ui32(int32(65535)) {
			v45 = int32(base.Ui32(l0)>>(uint(int32(2))%32))&int32(1984) | l0&int32(63)
		} else {
			if base.Ui32(l0) <= base.Ui32(int32(16777215)) {
				v45 = int32(base.Ui32(l0)>>(uint(int32(4))%32))&int32(61440) | (int32(base.Ui32(l0)>>(uint(int32(2))%32))&int32(4032) | l0&int32(63))
			} else {
				v45 = int32(base.Ui32(l0)>>(uint(int32(2))%32))&int32(4032) | (int32(base.Ui32(l0)>>(uint(int32(6))%32))&int32(1835008) | (int32(base.Ui32(l0)>>(uint(int32(4))%32))&int32(258048) | l0&int32(63)))
			}
		}
	}
	if base.Ui32(v45-int32(1106)) <= base.Ui32(int32(7101)) {
		v51 = v45 - int32(286)
		v52 = int32(65535)
		v53 = v51 & v52
		v55 = base.I32_div_u_s(v53, int32(1260))
		v58 = int32(10)
		v59 = base.I32_div_u_s(v53, v58)
		v61 = base.I32_rem_u_s(v59, int32(126))
		return v55<<(uint(int32(16))%32) | (v61<<(uint(int32(8))%32) + (v51-v59*v58)&v52 + int32(33024)) | int32(-2127560656)
	} else {
		if base.Ui32(v45-int32(9795)) <= base.Ui32(int32(2109)) {
			v81 = v45 - int32(576)
			v82 = int32(65535)
			v83 = v81 & v82
			v85 = base.I32_div_u_s(v83, int32(1260))
			v88 = int32(10)
			v89 = base.I32_div_u_s(v83, v88)
			v91 = base.I32_rem_u_s(v89, int32(126))
			return v85<<(uint(int32(16))%32) | (v91<<(uint(int32(8))%32) + (v81-v89*v88)&v82 + int32(33024)) | int32(-2127560656)
		} else {
			if base.Ui32(v45-int32(13851)) <= base.Ui32(int32(764)) {
				v111 = v45 - int32(878)
				v112 = int32(65535)
				v113 = v111 & v112
				v115 = base.I32_div_u_s(v113, int32(1260))
				v120 = int32(10)
				v121 = base.I32_div_u_s(v113, v120)
				v123 = base.I32_rem_u_s(v121, int32(126))
				return v115<<(uint(int32(16))%32) + int32(2146828288) | (v123<<(uint(int32(8))%32) + (v111-v121*v120)&v112 + int32(33024)) | int32(-2110783440)
			} else {
				if base.Ui32(v45-int32(15585)) <= base.Ui32(int32(884)) {
					v143 = v45 - int32(887)
					v144 = int32(65535)
					v145 = v143 & v144
					v147 = base.I32_div_u_s(v145, int32(1260))
					v152 = int32(10)
					v153 = base.I32_div_u_s(v145, v152)
					v155 = base.I32_rem_u_s(v153, int32(126))
					return v147<<(uint(int32(16))%32) + int32(2146828288) | (v155<<(uint(int32(8))%32) + (v143-v153*v152)&v144 + int32(33024)) | int32(-2110783440)
				} else {
					if base.Ui32(v45-int32(16736)) <= base.Ui32(int32(470)) {
						v175 = v45 - int32(889)
						v176 = int32(65535)
						v177 = v175 & v176
						v179 = base.I32_div_u_s(v177, int32(1260))
						v184 = int32(10)
						v185 = base.I32_div_u_s(v177, v184)
						v187 = base.I32_rem_u_s(v185, int32(126))
						return v179<<(uint(int32(16))%32) + int32(2146828288) | (v187<<(uint(int32(8))%32) + (v175-v185*v184)&v176 + int32(33024)) | int32(-2110783440)
					} else {
						if base.Ui32(v45-int32(17623)) <= base.Ui32(int32(372)) {
							v207 = v45 - int32(894)
							v208 = int32(65535)
							v209 = v207 & v208
							v211 = base.I32_div_u_s(v209, int32(1260))
							v216 = int32(10)
							v217 = base.I32_div_u_s(v209, v216)
							v219 = base.I32_rem_u_s(v217, int32(126))
							return v211<<(uint(int32(16))%32) + int32(2146828288) | (v219<<(uint(int32(8))%32) + (v207-v217*v216)&v208 + int32(33024)) | int32(-2110783440)
						} else {
							if base.Ui32(v45-int32(18318)) <= base.Ui32(int32(440)) {
								v239 = v45 - int32(900)
								v240 = int32(65535)
								v241 = v239 & v240
								v243 = base.I32_div_u_s(v241, int32(1260))
								v248 = int32(10)
								v249 = base.I32_div_u_s(v241, v248)
								v251 = base.I32_rem_u_s(v249, int32(126))
								return v243<<(uint(int32(16))%32) + int32(2146828288) | (v251<<(uint(int32(8))%32) + (v239-v249*v248)&v240 + int32(33024)) | int32(-2110783440)
							} else {
								if base.Ui32(v45-int32(18872)) <= base.Ui32(int32(702)) {
									v271 = v45 - int32(911)
									v272 = int32(65535)
									v273 = v271 & v272
									v275 = base.I32_div_u_s(v273, int32(1260))
									v280 = int32(10)
									v281 = base.I32_div_u_s(v273, v280)
									v283 = base.I32_rem_u_s(v281, int32(126))
									return v275<<(uint(int32(16))%32) + int32(2146828288) | (v283<<(uint(int32(8))%32) + (v271-v281*v280)&v272 + int32(33024)) | int32(-2110783440)
								} else {
									if base.Ui32(v45-int32(40870)) <= base.Ui32(int32(14425)) {
										v303 = v45 - int32(21827)
										v304 = int32(65535)
										v305 = v303 & v304
										v307 = base.I32_div_u_s(v305, int32(12600))
										v311 = base.I32_div_u_s(v305, int32(1260))
										v312 = int32(10)
										v313 = base.I32_rem_u_s(v311, v312)
										v320 = base.I32_div_u_s(v305, v312)
										v322 = base.I32_rem_u_s(v320, int32(126))
										return v307<<(uint(int32(24))%32) | v313<<(uint(int32(16))%32) - int32(2130706432) | (v322<<(uint(int32(8))%32) + (v303-v320*v312)&v304 + int32(33024)) | int32(3145776)
									} else {
										if base.Ui32(v45-int32(59493)) <= base.Ui32(int32(4294)) {
											v342 = v45 - int32(25943)
											v343 = int32(65535)
											v344 = v342 & v343
											v346 = base.I32_div_u_s(v344, int32(12600))
											v350 = base.I32_div_u_s(v344, int32(1260))
											v351 = int32(10)
											v352 = base.I32_rem_u_s(v350, v351)
											v359 = base.I32_div_u_s(v344, v351)
											v361 = base.I32_rem_u_s(v359, int32(126))
											return v346<<(uint(int32(24))%32) | v352<<(uint(int32(16))%32) - int32(2130706432) | (v361<<(uint(int32(8))%32) + (v342-v359*v351)&v343 + int32(33024)) | int32(3145776)
										} else {
											if base.Ui32(v45-int32(64042)) <= base.Ui32(int32(1029)) {
												v381 = v45 - int32(25964)
												v382 = int32(65535)
												v383 = v381 & v382
												v385 = base.I32_div_u_s(v383, int32(12600))
												v389 = base.I32_div_u_s(v383, int32(1260))
												v390 = int32(10)
												v391 = base.I32_rem_u_s(v389, v390)
												v398 = base.I32_div_u_s(v383, v390)
												v400 = base.I32_rem_u_s(v398, int32(126))
												return v385<<(uint(int32(24))%32) | v391<<(uint(int32(16))%32) - int32(2130706432) | (v400<<(uint(int32(8))%32) + (v381-v398*v390)&v382 + int32(33024)) | int32(3145776)
											} else {
												if base.Ui32(v45-int32(65510)) <= base.Ui32(int32(25)) {
													v420 = v45 - int32(26116)
													v421 = int32(65535)
													v422 = v420 & v421
													v424 = base.I32_div_u_s(v422, int32(12600))
													v428 = base.I32_div_u_s(v422, int32(1260))
													v429 = int32(10)
													v430 = base.I32_rem_u_s(v428, v429)
													v437 = base.I32_div_u_s(v422, v429)
													v439 = base.I32_rem_u_s(v437, int32(126))
													return v424<<(uint(int32(24))%32) | v430<<(uint(int32(16))%32) - int32(2130706432) | (v439<<(uint(int32(8))%32) + (v420-v437*v429)&v421 + int32(33024)) | int32(3145776)
												} else {
													if base.Ui32(v45-int32(65536)) <= base.Ui32(int32(1048575)) {
														v459 = v45 + int32(123464)
														v460 = int32(10)
														v461 = base.I32_div_u_s(v459, v460)
														v463 = base.I32_rem_u_s(v461, int32(126))
														v473 = base.I32_div_u_s(v459, int32(12600))
														v477 = base.I32_div_u_s(v459, int32(1260))
														v481 = base.I32_rem_u_s(v477&int32(65535), v460)
														v492 = v463<<(uint(int32(8))%32) + (v459 - v461*v460) + int32(33024) | (v473<<(uint(int32(24))%32) | v481<<(uint(int32(16))%32) - int32(2130706432)) | int32(3145776)
													} else {
														v492 = int32(0)
													}
													return v492
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
	var v121 int32
	_ = v121
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
	var v192 int32
	_ = v192
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
	var v241 int32
	_ = v241
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
	v121 = v5
	goto L36
L34:
	;
	v192 = v5
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
	v192 = v126
	goto L35
L38:
	;
	v124 = F_bms_add_member(m, v121, v118)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L3
	} else {
		goto L41
	}
L39:
	;
	v126 = v121
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
		v121 = v126
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
	if v192 == v196 {
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
		v241 = v196
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v250 = v241
	goto L55
L60:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v192)+4))
	v207 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v207 < v206 {
		v241 = v196
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
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v192+v213+v225)))
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v225+(l3+v213))))
	v232 = v227 & (v229 ^ int32(-1))
	v234 = base.B2i32(v232 == int32(0))
	if v232 != 0 {
		v241 = v234
		goto L59
	} else {
		goto L67
	}
L66:
	;
	v241 = v234
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
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	if l0 == int32(0) {
		v80 = int32(0)
		m.G0 = v7 + int32(16)
		return v80
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if v12 == int32(9) {
			v16 = F_palloc0(m, int32(72))
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v16))) = int32(9)
				v23 = F__emscripten_memcpy_bulkmem(m, v16, l0, int32(72))
				mBase = m.M
				v25 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v23)+44)) = v25
				*(*int32)(unsafe.Add(mBase, uint32(v23)+32)) = v25
				v29 = F_copyObjectImpl(m, v23)
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return int32(0)
				} else {
					v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
					*(*int32)(unsafe.Add(mBase, uint32(v23)+32)) = v31
					v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
					*(*int32)(unsafe.Add(mBase, uint32(v23)+44)) = v33
					*(*int32)(unsafe.Add(mBase, uint32(v23)+56)) = int32(6)
					v40 = *(*int32)(unsafe.Add(mBase, uint32(v23)+20))
					if v40 == int32(2281) {
						v43 = int32(17)
					} else {
						v43 = v40
					}
					*(*int32)(unsafe.Add(mBase, uint32(v23)+8)) = v43
					v50 = int32(0)
					v52 = F_makeTargetEntry(m, v23, int32(1), v50, v50)
					mBase = m.M
					v53 = m.ExcPending
					if v53 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v52
						*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = v52
						v59 = F_list_make1_impl(m, int32(1), v7+int32(8))
						mBase = m.M
						v60 = m.ExcPending
						if v60 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v29)+32)) = v59
							*(*int32)(unsafe.Add(mBase, uint32(v29)+56)) = int32(9)
							v80 = v29
							m.G0 = v7 + int32(16)
							return v80
						}
					}
				}
			}
		} else {
			v77 = F_expression_tree_mutator_impl(m, l0, int32(837), l1)
			mBase = m.M
			v78 = m.ExcPending
			if v78 != 0 {
				return int32(0)
			} else {
				v80 = v77
				m.G0 = v7 + int32(16)
				return v80
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
	var v8 int32
	_ = v8
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
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v186 int32
	_ = v186
	var v191 int32
	_ = v191
	v8 = m.G0
	v10 = v8 - int32(4160)
	m.G0 = v10
	v13 = *(*int32)(unsafe.Add(mBase, _consts[383]))
	v14 = F_mkdir(m, l1, v13)
	mBase = m.M
	goto L1
L1:
	;
	if v14 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L2:
	;
	v17 = F_AllocateDir(m, l0)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
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
	v178 = m.ExcPending
	if v178 != 0 {
		goto L5
	} else {
		goto L58
	}
L5:
	;
	return
L6:
	;
	v19 = F_ReadDir(m, v17, l0)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	if v19 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v25 = v19
	goto L11
L9:
	;
	goto L10
L10:
	;
	F_FreeDir(m, v17)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L5
	} else {
		goto L33
	}
L11:
	;
	v29 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	if v29 != 0 {
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
	v31 = m.ExcPending
	if v31 != 0 {
		goto L5
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+19)))
	if v32 != int32(46) {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	goto L15
L17:
	;
	v90 = F_ReadDir(m, v17, l0)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L5
	} else {
		goto L31
	}
L18:
	;
	v45 = v25 + int32(19)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+36)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = l0
	v54 = F_pg_snprintf(m, v10+int32(2112), int32(2048), int32(176209), v10+int32(32))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L5
	} else {
		goto L23
	}
L19:
	;
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+20)))
	if v35 == int32(0) {
		goto L17
	} else {
		goto L20
	}
L20:
	;
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+20)))
	if v38 != int32(46) {
		goto L18
	} else {
		goto L21
	}
L21:
	;
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+21)))
	if v41 == int32(0) {
		goto L17
	} else {
		goto L22
	}
L22:
	;
	goto L18
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = l1
	v64 = F_pg_snprintf(m, v10-int32(-64), int32(2048), int32(176209), v10+int32(16))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L5
	} else {
		goto L24
	}
L24:
	;
	v70 = F_get_dirent_type(m, v10+int32(2112), v25, int32(0), int32(21))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L5
	} else {
		goto L27
	}
L25:
	;
	F_copy_file(m, v10+int32(2112), v10-int32(-64))
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
	F_copydir(m, v10+int32(2112), v10-int32(-64), int32(1))
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
	if v90 != 0 {
		v25 = v90
		goto L11
	} else {
		goto L32
	}
L32:
	;
	goto L12
L33:
	;
	v102 = int32(*(*uint8)(unsafe.Add(mBase, _consts[191])))
	if v102 == int32(1) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v105 = F_AllocateDir(m, l1)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L5
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	m.G0 = v10 + int32(4160)
	return
L37:
	;
	v107 = F_ReadDir(m, v105, l1)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L5
	} else {
		goto L38
	}
L38:
	;
	if v107 != 0 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v113 = v107
	goto L42
L40:
	;
	goto L41
L41:
	;
	F_FreeDir(m, v105)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L5
	} else {
		goto L56
	}
L42:
	;
	v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v113)+19)))
	if v116 != int32(46) {
		goto L45
	} else {
		goto L46
	}
L43:
	;
	goto L41
L44:
	;
	v151 = F_ReadDir(m, v105, l1)
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L5
	} else {
		goto L54
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v113 + int32(19)
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = l1
	v136 = F_pg_snprintf(m, v10-int32(-64), int32(2048), int32(176209), v10)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L5
	} else {
		goto L50
	}
L46:
	;
	v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v113)+20)))
	if v119 == int32(0) {
		goto L44
	} else {
		goto L47
	}
L47:
	;
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v113)+20)))
	if v122 != int32(46) {
		goto L45
	} else {
		goto L48
	}
L48:
	;
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v113)+21)))
	if v125 == int32(0) {
		goto L44
	} else {
		goto L49
	}
L49:
	;
	goto L45
L50:
	;
	v142 = F_get_dirent_type(m, v10-int32(-64), v113, int32(0), int32(21))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L5
	} else {
		goto L51
	}
L51:
	;
	if v142 != int32(2) {
		goto L44
	} else {
		goto L52
	}
L52:
	;
	F_fsync_fname(m, v10-int32(-64), int32(0))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
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
		v113 = v151
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
	v164 = m.ExcPending
	if v164 != 0 {
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
	v180 = m.ExcPending
	if v180 != 0 {
		goto L5
	} else {
		goto L59
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+48)) = l1
	F_errmsg(m, int32(294861), v10+int32(48))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L5
	} else {
		goto L60
	}
L60:
	;
	F_errfinish(m, int32(491941), int32(58), int32(212156))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
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
			*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = int32(16384)
			v15 = F_palloc(m, int32(16386))
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v15
				if v15 == int32(0) {
					F_yy_fatal_error_2(m, int32(665354))
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
					v23 = *(*int32)(unsafe.Add(mBase, _consts[155]))
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
					*(*int32)(unsafe.Add(mBase, _consts[155])) = v23
					return v8
				}
			}
		} else {
			F_yy_fatal_error_2(m, int32(665354))
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
	var v35 float64
	_ = v35
	var v37 float64
	_ = v37
	var v41 int32
	_ = v41
	var v42 float64
	_ = v42
	var v45 int64
	_ = v45
	var v46 float64
	_ = v46
	var v48 float64
	_ = v48
	var v49 int32
	_ = v49
	var v52 float64
	_ = v52
	var v57 float64
	_ = v57
	var v59 float64
	_ = v59
	var v60 float64
	_ = v60
	var v62 float64
	_ = v62
	var v63 float64
	_ = v63
	var v66 float64
	_ = v66
	var v69 float64
	_ = v69
	var v73 float64
	_ = v73
	var v80 float64
	_ = v80
	var v81 float64
	_ = v81
	var v84 float64
	_ = v84
	var v88 float64
	_ = v88
	var v98 float64
	_ = v98
	var v101 float64
	_ = v101
	var v104 float64
	_ = v104
	var v107 int32
	_ = v107
	var v108 float64
	_ = v108
	var v114 float64
	_ = v114
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
	v25 = *(*float64)(unsafe.Add(mBase, _consts[481]))
	v28 = base.F64_mul(v23, base.F64_add(base.F64_add(v25, v25), l5))
	v35 = base.F64_convert_i32_u((l4+int32(7))&int32(-8) + int32(24))
	v37 = base.F64_mul(l3, v35)
	v41 = base.F64_lt(l7, v23) & base.F64_gt(l7, float64(0))
	if v41 != 0 {
		v42 = base.F64_mul(l7, v35)
	} else {
		v42 = v37
	}
	v45 = base.I64_extend_i32_s(l6) << (uint(int64(10)) % 64)
	v46 = base.F64_convert_i64_s(v45)
	if base.F64_gt(v42, v46) != 0 {
		v48 = F_log(m, v23)
		mBase = m.M
		v49 = F_tuplesort_merge_order(m, v45)
		mBase = m.M
		v52 = base.F64_mul(base.F64_div(v48, float64(0.693147180559945)), v28)
		*(*float64)(unsafe.Add(mBase, uint32(v15))) = v52
		v57 = base.F64_ceil(base.F64_mul(v37, float64(0.0001220703125)))
		v59 = base.F64_div(v37, v46)
		v60 = base.F64_convert_i32_s(v49)
		if base.F64_gt(v59, v60) != 0 {
			v62 = F_log(m, v59)
			mBase = m.M
			v63 = F_log(m, v60)
			mBase = m.M
			v66 = base.F64_ceil(base.F64_div(v62, v63))
		} else {
			v66 = float64(1)
		}
		v69 = *(*float64)(unsafe.Add(mBase, _consts[482]))
		v73 = *(*float64)(unsafe.Add(mBase, _consts[483]))
		v98 = base.F64_add(base.F64_mul(base.F64_mul(base.F64_add(v57, v57), v66), base.F64_add(base.F64_mul(v69, float64(0.75)), base.F64_mul(v73, float64(0.25)))), v52)
	} else {
		if v41 != 0 {
			v80 = l7
		} else {
			v80 = v23
		}
		v81 = base.F64_add(v80, v80)
		if base.F64_gt(v37, v46)|base.F64_gt(v23, v81) != 0 {
			v84 = F_log(m, v81)
			mBase = m.M
			v98 = base.F64_mul(base.F64_div(v84, float64(0.693147180559945)), v28)
		} else {
			v88 = F_log(m, v23)
			mBase = m.M
			v98 = base.F64_mul(base.F64_div(v88, float64(0.693147180559945)), v28)
		}
	}
	*(*float64)(unsafe.Add(mBase, uint32(v15))) = v98
	v101 = *(*float64)(unsafe.Add(mBase, _consts[481]))
	*(*float64)(unsafe.Add(mBase, uint32(v12))) = base.F64_mul(v23, v101)
	v104 = *(*float64)(unsafe.Add(mBase, uint32(v12)+8))
	*(*float64)(unsafe.Add(mBase, uint32(l0)+32)) = l3
	v107 = int32(*(*uint8)(unsafe.Add(mBase, _consts[484])))
	v108 = base.F64_add(l2, v104)
	*(*float64)(unsafe.Add(mBase, uint32(l0)+48)) = v108
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = l1 + (v107 ^ int32(1))
	v114 = *(*float64)(unsafe.Add(mBase, uint32(v12)))
	*(*float64)(unsafe.Add(mBase, uint32(l0)+56)) = base.F64_add(v108, v114)
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
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
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
			v21 = int32(4)
			v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
			if v23&int32(254) == int32(2) {
				v32 = v21
			} else {
				v32 = base.B2i32(v23 == int32(18)) << (uint(v21) % 32)
			}
			if v23 == int32(1) {
				v35 = v21
			} else {
				v35 = v32
			}
			v46 = v35
		} else {
			v36 = int32(1)
			if v17 != 0 {
				v46 = int32(base.Ui32(v15)>>(uint(v36)%32)) - v36
			} else {
				v40 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
				v46 = int32(base.Ui32(v40)>>(uint(int32(2))%32)) - int32(4)
			}
		}
		v47 = m.Env.Pgmem_crc32c(m, int32(-1), v18, v46)
		mBase = m.M
		v51 = F_Int64GetDatum(m, base.I64_extend_i32_u(v47^int32(-1)))
		mBase = m.M
		v52 = m.ExcPending
		if v52 != 0 {
			return int32(0)
		} else {
			return v51
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
		v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+26)))
		if v24 == int32(1) {
			v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+21)))
			v28 = v27
		} else {
			v28 = v18
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
		v43 = *(*int32)(unsafe.Add(mBase, _consts[127]))
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
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v74 int32
	_ = v74
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
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
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v163 int32
	_ = v163
	var v171 int32
	_ = v171
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
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
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
	var v283 int32
	_ = v283
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
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v328 int32
	_ = v328
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v367 int32
	_ = v367
	var v370 int32
	_ = v370
	var v382 int32
	_ = v382
	var v386 int32
	_ = v386
	var v388 int32
	_ = v388
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v397 int32
	_ = v397
	var v400 int32
	_ = v400
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v427 int32
	_ = v427
	var v432 int32
	_ = v432
	var v441 int32
	_ = v441
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v463 int32
	_ = v463
	var v478 int32
	_ = v478
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v513 int32
	_ = v513
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v553 int32
	_ = v553
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v576 int32
	_ = v576
	var v591 int32
	_ = v591
	var v593 float64
	_ = v593
	var v595 float64
	_ = v595
	var v597 float64
	_ = v597
	var v599 int32
	_ = v599
	var v600 int32
	_ = v600
	var v602 int32
	_ = v602
	var v604 int32
	_ = v604
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
		v142 = v29
		v145 = v6
		v149 = v6
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
	v171 = v156
	goto L19
L2:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	if v33 <= int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v142 = v29
	v145 = v6
	v149 = v6
	v154 = v25
	goto L1
L4:
	;
	goto L5
L5:
	;
	v43 = v33
	v45 = v6
	v47 = v6
	v51 = v6
	goto L6
L6:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v30)+12))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v56+v45<<(uint(int32(2))%32))))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)+8))
	if v61 == int32(0) {
		v116 = v43
		v120 = v47
		v124 = v51
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
	v133 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v142 = v132
	v145 = v120
	v149 = v124
	v154 = v133
	goto L1
L8:
	;
	v130 = v45 + int32(1)
	if v130 < v116 {
		v43 = v116
		v45 = v130
		v47 = v120
		v51 = v124
		goto L6
	} else {
		goto L18
	}
L9:
	;
	v64 = int32(0)
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v61)+4))
	if v65 <= v64 {
		v116 = v43
		v120 = v47
		v124 = v51
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v68 = int32(*(*int16)(unsafe.Add(mBase, uint32(v60)+14)))
	v74 = v64
	v80 = v47
	v84 = v51
	goto L11
L11:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v61)+12))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v89+v74<<(uint(int32(2))%32))))
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v93)+4))
	v95 = F_lappend(m, v80, v94)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	v116 = v108
	v120 = v95
	v124 = v102
	goto L8
L13:
	;
	return int32(0)
L14:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v60)+16))
	v100 = F_fix_indexqual_clause(m, l0, v25, v68, v94, v99)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v102 = F_lappend(m, v84, v100)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L13
	} else {
		goto L16
	}
L16:
	;
	v105 = v74 + int32(1)
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v61)+4))
	if v105 < v106 {
		v74 = v105
		v80 = v95
		v84 = v102
		goto L11
	} else {
		goto L17
	}
L17:
	;
	goto L12
L18:
	;
	goto L7
L19:
	;
	v178 = int32(0)
	if v142 == v178 {
		v189 = v178
		goto L21
	} else {
		goto L22
	}
L20:
	;
	v341 = F_order_qual_clauses(m, l0, v328)
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L13
	} else {
		goto L70
	}
L21:
	;
	if v155 != 0 {
		goto L26
	} else {
		goto L27
	}
L22:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v142)+4))
	if v183 <= v163 {
		v189 = int32(0)
		goto L21
	} else {
		goto L23
	}
L23:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v142)+12))
	v189 = v185 + v163<<(uint(int32(2))%32)
	goto L21
L24:
	;
	goto L20
L25:
	;
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v197)))
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v189)))
	v317 = F_fix_indexqual_clause(m, l0, v154, v314, v315, int32(0))
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L13
	} else {
		goto L68
	}
L26:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v155)+4))
	if v190 <= v163 {
		goto L29
	} else {
		goto L30
	}
L27:
	;
	v199 = v178
	goto L28
L28:
	;
	v200 = int32(0)
	if l3 == v200 {
		v328 = v200
		goto L24
	} else {
		goto L33
	}
L29:
	;
	v199 = v171
	goto L28
L30:
	;
	if v189 == int32(0) {
		goto L29
	} else {
		goto L31
	}
L31:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v155)+12))
	v197 = v194 + v163<<(uint(int32(2))%32)
	if v197 != 0 {
		goto L25
	} else {
		goto L32
	}
L32:
	;
	goto L29
L33:
	;
	v203 = int32(0)
	v204 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v204 <= v203 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v328 = v200
	goto L24
L35:
	;
	goto L36
L36:
	;
	v212 = v203
	v214 = v200
	goto L37
L37:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v227+v212<<(uint(int32(2))%32))))
	v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v231)+10)))
	if v232 != 0 {
		v306 = v214
		goto L39
	} else {
		goto L40
	}
L38:
	;
	v328 = v306
	goto L24
L39:
	;
	v309 = v212 + int32(1)
	v310 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v309 < v310 {
		v212 = v309
		v214 = v306
		goto L37
	} else {
		goto L67
	}
L40:
	;
	if v30 != 0 {
		goto L43
	} else {
		goto L44
	}
L41:
	;
	if v283 != 0 {
		v306 = v214
		goto L39
	} else {
		goto L58
	}
L42:
	;
	goto L41
L43:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	if v238 <= int32(0) {
		v283 = int32(0)
		goto L42
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	v283 = int32(0)
	goto L42
L46:
	;
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v231)+60))
	v242 = int32(0)
	if v242 < v238 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v245 = v238
	goto L49
L48:
	;
	v245 = v242
	goto L49
L49:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v30)+12))
	v249 = int32(0)
	goto L50
L50:
	;
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v246+v249<<(uint(int32(2))%32))))
	v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v258)+12)))
	if v259 != 0 {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	goto L45
L52:
	;
	v270 = v249 + int32(1)
	if v270 != v245 {
		v249 = v270
		goto L50
	} else {
		goto L57
	}
L53:
	;
	v260 = int32(1)
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v258)+4))
	if v231 == v261 {
		v283 = v260
		goto L42
	} else {
		goto L54
	}
L54:
	;
	if v241 == int32(0) {
		goto L52
	} else {
		goto L55
	}
L55:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v261)+60))
	if v265 == v241 {
		v283 = v260
		goto L42
	} else {
		goto L56
	}
L56:
	;
	goto L52
L57:
	;
	goto L51
L58:
	;
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v231)+4))
	v288 = F_contain_mutable_functions(m, v287)
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L13
	} else {
		goto L59
	}
L59:
	;
	if v288 == int32(0) {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v231)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+24)) = v292
	*(*int32)(unsafe.Add(mBase, uint32(v23)+28)) = v292
	v298 = F_list_make1_impl(m, int32(1), v23+int32(24))
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L13
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	v304 = F_lappend(m, v214, v231)
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L13
	} else {
		goto L66
	}
L63:
	;
	v301 = F_predicate_implied_by(m, v298, v145, int32(0))
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L13
	} else {
		goto L64
	}
L64:
	;
	if v301 != 0 {
		v306 = v214
		goto L39
	} else {
		goto L65
	}
L65:
	;
	goto L62
L66:
	;
	v306 = v304
	goto L39
L67:
	;
	goto L38
L68:
	;
	v319 = F_lappend(m, v171, v317)
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L13
	} else {
		goto L69
	}
L69:
	;
	v163 = v163 + int32(1)
	v171 = v319
	goto L19
L70:
	;
	v344 = F_extract_actual_clauses(m, v341, int32(0))
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L13
	} else {
		goto L71
	}
L71:
	;
	v346 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v346 != 0 {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v347 = F_replace_nestloop_params_mutator(m, v145, l0)
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L13
	} else {
		goto L75
	}
L73:
	;
	v353 = v344
	v354 = v145
	v355 = v29
	goto L74
L74:
	;
	if v355 == int32(0) {
		goto L79
	} else {
		goto L80
	}
L75:
	;
	v349 = F_replace_nestloop_params_mutator(m, v344, l0)
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L13
	} else {
		goto L76
	}
L76:
	;
	v351 = F_replace_nestloop_params_mutator(m, v29, l0)
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L13
	} else {
		goto L77
	}
L77:
	;
	v353 = v349
	v354 = v347
	v355 = v351
	goto L74
L78:
	;
	if l4 != 0 {
		goto L100
	} else {
		goto L101
	}
L79:
	;
	v441 = int32(0)
	goto L78
L80:
	;
	goto L81
L81:
	;
	v359 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
	v360 = int32(0)
	v367 = v360
	v370 = v360
	goto L82
L82:
	;
	v382 = int32(0)
	if v359 == v382 {
		v392 = v382
		goto L84
	} else {
		goto L85
	}
L83:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v418 = m.ExcPending
	if v418 != 0 {
		goto L13
	} else {
		goto L96
	}
L84:
	;
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v355)+4))
	if v393 <= v367 {
		v441 = v370
		goto L78
	} else {
		goto L87
	}
L85:
	;
	v386 = *(*int32)(unsafe.Add(mBase, uint32(v359)+4))
	if v386 <= v367 {
		v392 = int32(0)
		goto L84
	} else {
		goto L86
	}
L86:
	;
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v359)+12))
	v392 = v388 + v367<<(uint(int32(2))%32)
	goto L84
L87:
	;
	if v392 == int32(0) {
		v441 = v370
		goto L78
	} else {
		goto L88
	}
L88:
	;
	v397 = *(*int32)(unsafe.Add(mBase, uint32(v355)+12))
	v400 = v397 + v367<<(uint(int32(2))%32)
	if v400 == int32(0) {
		v441 = v370
		goto L78
	} else {
		goto L89
	}
L89:
	;
	v403 = *(*int32)(unsafe.Add(mBase, uint32(v392)))
	v404 = *(*int32)(unsafe.Add(mBase, uint32(v400)))
	v405 = F_exprType(m, v404)
	mBase = m.M
	v406 = m.ExcPending
	if v406 != 0 {
		goto L13
	} else {
		goto L90
	}
L90:
	;
	v407 = *(*int32)(unsafe.Add(mBase, uint32(v403)+8))
	v408 = *(*int32)(unsafe.Add(mBase, uint32(v403)+12))
	v409 = F_get_opfamily_member_for_cmptype(m, v407, v405, v405, v408)
	mBase = m.M
	v410 = m.ExcPending
	if v410 != 0 {
		goto L13
	} else {
		goto L91
	}
L91:
	;
	if v409 != 0 {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	v413 = F_lappend_oid(m, v370, v409)
	mBase = m.M
	v414 = m.ExcPending
	if v414 != 0 {
		goto L13
	} else {
		goto L95
	}
L93:
	;
	goto L94
L94:
	;
	goto L83
L95:
	;
	v367 = v367 + int32(1)
	v370 = v413
	goto L82
L96:
	;
	v419 = *(*int32)(unsafe.Add(mBase, uint32(v403)+12))
	v420 = *(*int32)(unsafe.Add(mBase, uint32(v403)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+12)) = v420
	*(*int32)(unsafe.Add(mBase, uint32(v23)+8)) = v405
	*(*int32)(unsafe.Add(mBase, uint32(v23)+4)) = v405
	*(*int32)(unsafe.Add(mBase, uint32(v23))) = v419
	F_errmsg_internal(m, int32(39557), v23)
	mBase = m.M
	v427 = m.ExcPending
	if v427 != 0 {
		goto L13
	} else {
		goto L97
	}
L97:
	;
	F_errfinish(m, int32(493534), int32(3138), int32(281368))
	mBase = m.M
	v432 = m.ExcPending
	if v432 != 0 {
		goto L13
	} else {
		goto L98
	}
L98:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L99:
	;
	v591 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v576)+4)) = v591
	v593 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v576)+8)) = v593
	v595 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v576)+16)) = v595
	v597 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v576)+24)) = v597
	v599 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v600 = *(*int32)(unsafe.Add(mBase, uint32(v599)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v576)+32)) = v600
	v602 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)))
	*(*uint8)(unsafe.Add(mBase, uint32(v576)+36)) = uint8(v602)
	v604 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
	*(*uint8)(unsafe.Add(mBase, uint32(v576)+37)) = uint8(v604)
	m.G0 = v23 + int32(32)
	return v576
L100:
	;
	v453 = *(*int32)(unsafe.Add(mBase, uint32(v25)+92))
	if v453 != 0 {
		goto L103
	} else {
		goto L104
	}
L101:
	;
	goto L102
L102:
	;
	v553 = *(*int32)(unsafe.Add(mBase, uint32(l1)+88))
	v555 = F_palloc0(m, int32(112))
	mBase = m.M
	v556 = m.ExcPending
	if v556 != 0 {
		goto L13
	} else {
		goto L113
	}
L103:
	;
	v454 = *(*int32)(unsafe.Add(mBase, uint32(v453)+4))
	if int32(0) < v454 {
		goto L106
	} else {
		goto L107
	}
L104:
	;
	v535 = int32(0)
	goto L105
L105:
	;
	v536 = *(*int32)(unsafe.Add(mBase, uint32(l1)+88))
	v538 = F_palloc0(m, int32(104))
	mBase = m.M
	v539 = m.ExcPending
	if v539 != 0 {
		goto L13
	} else {
		goto L112
	}
L106:
	;
	v463 = int32(0)
	goto L109
L107:
	;
	goto L108
L108:
	;
	v513 = *(*int32)(unsafe.Add(mBase, uint32(v25)+92))
	v535 = v513
	goto L105
L109:
	;
	v478 = *(*int32)(unsafe.Add(mBase, uint32(v453)+12))
	v482 = *(*int32)(unsafe.Add(mBase, uint32(v478+v463<<(uint(int32(2))%32))))
	v483 = *(*int32)(unsafe.Add(mBase, uint32(v25)+76))
	v485 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v483+v463))))
	v486 = int32(1)
	v487 = v485 ^ v486
	*(*uint8)(unsafe.Add(mBase, uint32(v482)+26)) = uint8(v487)
	v490 = v463 + v486
	v491 = *(*int32)(unsafe.Add(mBase, uint32(v453)+4))
	if v490 < v491 {
		v463 = v490
		goto L109
	} else {
		goto L111
	}
L110:
	;
	goto L108
L111:
	;
	goto L110
L112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v538)+100)) = v536
	*(*int32)(unsafe.Add(mBase, uint32(v538)+96)) = v535
	*(*int32)(unsafe.Add(mBase, uint32(v538)+92)) = v199
	*(*int32)(unsafe.Add(mBase, uint32(v538)+88)) = v354
	*(*int32)(unsafe.Add(mBase, uint32(v538)+84)) = v149
	*(*int32)(unsafe.Add(mBase, uint32(v538)+80)) = v26
	*(*int32)(unsafe.Add(mBase, uint32(v538)+72)) = v28
	*(*int64)(unsafe.Add(mBase, uint32(v538)+52)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v538)+48)) = v353
	*(*int32)(unsafe.Add(mBase, uint32(v538)+44)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v538))) = int32(342)
	v576 = v538
	goto L99
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v555)+104)) = v553
	*(*int32)(unsafe.Add(mBase, uint32(v555)+100)) = v441
	*(*int32)(unsafe.Add(mBase, uint32(v555)+96)) = v355
	*(*int32)(unsafe.Add(mBase, uint32(v555)+92)) = v199
	*(*int32)(unsafe.Add(mBase, uint32(v555)+88)) = v354
	*(*int32)(unsafe.Add(mBase, uint32(v555)+84)) = v149
	*(*int32)(unsafe.Add(mBase, uint32(v555)+80)) = v26
	*(*int32)(unsafe.Add(mBase, uint32(v555)+72)) = v28
	*(*int64)(unsafe.Add(mBase, uint32(v555)+52)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v555)+48)) = v353
	*(*int32)(unsafe.Add(mBase, uint32(v555)+44)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v555))) = int32(341)
	v576 = v555
	goto L99
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
	var v71 float64
	_ = v71
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
	var v105 float64
	_ = v105
	var v113 float64
	_ = v113
	var v115 float64
	_ = v115
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
				v71 = float64(1e+100)
				if base.F64_gt(v69, v71) != 0 {
					v83 = v71
				} else {
					if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v69)&int64(9223372036854775807)) {
						v83 = v71
					} else {
						v79 = float64(1)
						if base.F64_le(v69, v79) != 0 {
							v83 = v79
						} else {
							v83 = base.F64_nearest(v69)
						}
					}
				}
				v84 = *(*float64)(unsafe.Add(mBase, uint32(v61)))
				v85 = v83
				v86 = v84
			}
			if base.F64_lt(v86, v85) != 0 {
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
			v105 = v103
		} else {
			v105 = v62
		}
		if l6 != int64(0) {
			if int64(0) < l6 {
				v129 = v105
				v130 = base.F64_convert_i64_u(l6)
			} else {
				v113 = base.F64_mul(v62, float64(0.1))
				v115 = float64(1e+100)
				if base.F64_gt(v113, v115) != 0 {
					v127 = v115
				} else {
					if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v113)&int64(9223372036854775807)) {
						v127 = v115
					} else {
						v123 = float64(1)
						if base.F64_le(v113, v123) != 0 {
							v127 = v123
						} else {
							v127 = base.F64_nearest(v113)
						}
					}
				}
				v128 = *(*float64)(unsafe.Add(mBase, uint32(v61)))
				v129 = v128
				v130 = v127
			}
			if base.F64_lt(v129, v130) != 0 {
				v132 = v129
			} else {
				v132 = v130
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
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
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
			v70 = *(*float64)(unsafe.Add(mBase, _consts[481]))
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
			v90 = *(*float64)(unsafe.Add(mBase, _consts[481]))
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
			v106 = int32(*(*uint8)(unsafe.Add(mBase, _consts[509])))
			if v106 == int32(0) {
				*(*int32)(unsafe.Add(mBase, uint32(v15)+40)) = v58 + int32(1)
			} else {
			}
			v112 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
			v113 = *(*int32)(unsafe.Add(mBase, uint32(v112)+32))
			v124 = *(*float64)(unsafe.Add(mBase, _consts[507]))
			v126 = *(*int32)(unsafe.Add(mBase, _consts[127]))
			v130 = base.F64_mul(base.F64_mul(v124, base.F64_convert_i32_s(v126)), float64(1024))
			v131 = float64(4.294967295e+09)
			if base.F64_lt(v130, v131) != 0 {
				v134 = v130
			} else {
				v134 = v131
			}
			if base.F64_lt(v134, float64(4.294967296e+09))&base.F64_ge(v134, float64(0)) != 0 {
				v140 = base.I32_trunc_f64_u(v134)
				v142 = v140
			} else {
				v142 = int32(0)
			}
			if base.F64_gt(base.F64_mul(l6, base.F64_convert_i32_u((v113+int32(7))&int32(-8)+int32(16))), base.F64_convert_i32_u(v142)) != 0 {
				v145 = *(*int32)(unsafe.Add(mBase, uint32(v15)+40))
				*(*int32)(unsafe.Add(mBase, uint32(v15)+40)) = v145 + int32(1)
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
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
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
				v28 = v27
			} else {
				v28 = int32(0)
			}
			v30 = v28 & int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(v10)+21)) = uint8(v30)
			v32 = *(*int32)(unsafe.Add(mBase, uint32(l2)+24))
			*(*int32)(unsafe.Add(mBase, uint32(v10)+72)) = l2
			*(*int32)(unsafe.Add(mBase, uint32(v10)+64)) = l4
			*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = v32
			F_cost_subqueryscan(m, v10, l0, l1, v19, l3)
			mBase = m.M
			v37 = m.ExcPending
			if v37 != 0 {
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
		v20 = *(*int32)(unsafe.Add(mBase, _consts[276]))
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
							F_errmsg(m, int32(267007), v6)
							mBase = m.M
							v76 = m.ExcPending
							if v76 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(496302), int32(887), int32(431887))
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
					F_sequence_close(m, v30, int32(0))
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
						F_errmsg(m, int32(194843), v6+int32(16))
						mBase = m.M
						v56 = m.ExcPending
						if v56 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(496302), int32(881), int32(431887))
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
