package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

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
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int64
	_ = v71
	var v73 int32
	_ = v73
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
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v100 int32
	_ = v100
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v164 int32
	_ = v164
	var v182 int32
	_ = v182
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v272 int32
	_ = v272
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v286 int32
	_ = v286
	var v301 int32
	_ = v301
	var v317 int32
	_ = v317
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v350 int32
	_ = v350
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v364 int32
	_ = v364
	var v367 int32
	_ = v367
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v374 int32
	_ = v374
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v385 int64
	_ = v385
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v404 int32
	_ = v404
	var v407 int32
	_ = v407
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v430 int32
	_ = v430
	var v433 int32
	_ = v433
	var v438 int32
	_ = v438
	var v447 int32
	_ = v447
	var v450 int32
	_ = v450
	var v452 int32
	_ = v452
	var v455 int32
	_ = v455
	var v467 int32
	_ = v467
	var v470 int32
	_ = v470
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v477 int32
	_ = v477
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v482 int32
	_ = v482
	var v484 int32
	_ = v484
	var v488 int32
	_ = v488
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v508 int32
	_ = v508
	var v510 int32
	_ = v510
	var v513 int64
	_ = v513
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v522 int64
	_ = v522
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v526 int32
	_ = v526
	var v532 int32
	_ = v532
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v539 int32
	_ = v539
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v554 int32
	_ = v554
	var v566 int32
	_ = v566
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v587 int32
	_ = v587
	var v603 int32
	_ = v603
	var v647 int32
	_ = v647
	var v679 int32
	_ = v679
	var v680 int32
	_ = v680
	var v683 int32
	_ = v683
	var v687 int32
	_ = v687
	var v690 int32
	_ = v690
	var v693 int32
	_ = v693
	var v696 int32
	_ = v696
	var v698 int32
	_ = v698
	var v700 int32
	_ = v700
	var v701 int32
	_ = v701
	var v704 int32
	_ = v704
	var v707 int32
	_ = v707
	var v709 int32
	_ = v709
	var v710 int32
	_ = v710
	var v713 int32
	_ = v713
	var v717 int32
	_ = v717
	var v718 int32
	_ = v718
	var v726 int32
	_ = v726
	var v729 int32
	_ = v729
	var v732 int32
	_ = v732
	var v736 int32
	_ = v736
	var v739 int32
	_ = v739
	var v740 int32
	_ = v740
	var v744 int32
	_ = v744
	var v751 int32
	_ = v751
	var v753 int32
	_ = v753
	var v761 int32
	_ = v761
	var v762 int32
	_ = v762
	var v775 int32
	_ = v775
	var v788 int32
	_ = v788
	var v808 int32
	_ = v808
	var v809 int32
	_ = v809
	var v812 int32
	_ = v812
	var v819 int32
	_ = v819
	var v820 int32
	_ = v820
	var v830 int32
	_ = v830
	var v833 int32
	_ = v833
	var v834 int32
	_ = v834
	var v835 int32
	_ = v835
	var v836 int32
	_ = v836
	var v837 int32
	_ = v837
	var v838 int32
	_ = v838
	var v839 int32
	_ = v839
	var v840 int32
	_ = v840
	var v841 int32
	_ = v841
	var v842 int32
	_ = v842
	var v843 int32
	_ = v843
	var v845 int32
	_ = v845
	var v851 int32
	_ = v851
	var v860 int32
	_ = v860
	var v862 int32
	_ = v862
	var v863 int32
	_ = v863
	var v866 int32
	_ = v866
	var v870 int32
	_ = v870
	var v873 int32
	_ = v873
	var v875 int32
	_ = v875
	var v878 int32
	_ = v878
	var v885 int32
	_ = v885
	var v887 int32
	_ = v887
	var v895 int32
	_ = v895
	var v896 int32
	_ = v896
	var v909 int32
	_ = v909
	var v940 int32
	_ = v940
	var v942 int32
	_ = v942
	var v944 int32
	_ = v944
	var v946 int32
	_ = v946
	var v947 int32
	_ = v947
	var v951 int32
	_ = v951
	var v983 int32
	_ = v983
	var v986 int32
	_ = v986
	var v987 int32
	_ = v987
	var v988 int32
	_ = v988
	var v990 int32
	_ = v990
	var v992 int32
	_ = v992
	var v994 int32
	_ = v994
	var v997 int32
	_ = v997
	var v998 int32
	_ = v998
	var v999 int32
	_ = v999
	var v1001 int32
	_ = v1001
	var v1003 int32
	_ = v1003
	var v1034 int32
	_ = v1034
	var v1035 int32
	_ = v1035
	var v1038 int32
	_ = v1038
	var v1041 int32
	_ = v1041
	var v1044 int32
	_ = v1044
	var v1047 int32
	_ = v1047
	var v1048 int32
	_ = v1048
	var v1051 int32
	_ = v1051
	var v1054 int32
	_ = v1054
	var v1057 int32
	_ = v1057
	var v1058 int32
	_ = v1058
	var v1061 int32
	_ = v1061
	var v1064 int32
	_ = v1064
	var v1067 int32
	_ = v1067
	var v1068 int32
	_ = v1068
	var v1072 int32
	_ = v1072
	var v1073 int32
	_ = v1073
	var v1103 int32
	_ = v1103
	var v1105 int32
	_ = v1105
	var v1108 int32
	_ = v1108
	var v1109 int32
	_ = v1109
	var v1112 int32
	_ = v1112
	var v1115 int32
	_ = v1115
	var v1116 int32
	_ = v1116
	var v1118 int32
	_ = v1118
	var v1120 int64
	_ = v1120
	var v1158 int32
	_ = v1158
	var v1162 int32
	_ = v1162
	var v1167 int32
	_ = v1167
	var v1169 int32
	_ = v1169
	var v1172 int32
	_ = v1172
	var v1179 int32
	_ = v1179
	var v1186 int32
	_ = v1186
	var v1192 int32
	_ = v1192
	var v1195 int32
	_ = v1195
	var v1202 int32
	_ = v1202
	var v1209 int32
	_ = v1209
	var v1214 int32
	_ = v1214
	var v1215 int32
	_ = v1215
	var v1216 int32
	_ = v1216
	var v1217 int32
	_ = v1217
	var v1222 int32
	_ = v1222
	var v1227 int32
	_ = v1227
	var v1264 int32
	_ = v1264
	var v1267 int32
	_ = v1267
	var v1268 int32
	_ = v1268
	var v1276 int32
	_ = v1276
	var v1279 int32
	_ = v1279
	var v1280 int32
	_ = v1280
	var v1285 int32
	_ = v1285
	var v1289 int32
	_ = v1289
	var v1292 int32
	_ = v1292
	var v1293 int32
	_ = v1293
	var v1301 int32
	_ = v1301
	var v1304 int32
	_ = v1304
	var v1305 int32
	_ = v1305
	var v1310 int32
	_ = v1310
	var v1314 int32
	_ = v1314
	var v1317 int32
	_ = v1317
	var v1318 int32
	_ = v1318
	var v1326 int32
	_ = v1326
	var v1329 int32
	_ = v1329
	var v1330 int32
	_ = v1330
	var v1335 int32
	_ = v1335
	var v1339 int32
	_ = v1339
	var v1342 int32
	_ = v1342
	var v1343 int32
	_ = v1343
	var v1351 int32
	_ = v1351
	var v1354 int32
	_ = v1354
	var v1355 int32
	_ = v1355
	var v1360 int32
	_ = v1360
	var v1364 int32
	_ = v1364
	var v1367 int32
	_ = v1367
	var v1368 int32
	_ = v1368
	var v1376 int32
	_ = v1376
	var v1379 int32
	_ = v1379
	var v1380 int32
	_ = v1380
	var v1385 int32
	_ = v1385
	var v1389 int32
	_ = v1389
	var v1392 int32
	_ = v1392
	var v1393 int32
	_ = v1393
	var v1401 int32
	_ = v1401
	var v1404 int32
	_ = v1404
	var v1405 int32
	_ = v1405
	var v1410 int32
	_ = v1410
	var v1414 int32
	_ = v1414
	var v1417 int32
	_ = v1417
	var v1418 int32
	_ = v1418
	var v1424 int32
	_ = v1424
	var v1428 int32
	_ = v1428
	var v1433 int32
	_ = v1433
	var v1437 int32
	_ = v1437
	var v1440 int32
	_ = v1440
	var v1441 int32
	_ = v1441
	var v1449 int32
	_ = v1449
	var v1453 int32
	_ = v1453
	var v1458 int32
	_ = v1458
	v3 = int32(0)
	v29 = m.G0
	v31 = v29 - int32(128)
	m.G0 = v31
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+119)))
	if v34 == int32(112) {
		goto L9
	} else {
		goto L10
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1437 = m.ExcPending
	if v1437 != 0 {
		goto L27
	} else {
		goto L298
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1414 = m.ExcPending
	if v1414 != 0 {
		goto L27
	} else {
		goto L293
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1389 = m.ExcPending
	if v1389 != 0 {
		goto L27
	} else {
		goto L288
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1364 = m.ExcPending
	if v1364 != 0 {
		goto L27
	} else {
		goto L283
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1339 = m.ExcPending
	if v1339 != 0 {
		goto L27
	} else {
		goto L278
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1314 = m.ExcPending
	if v1314 != 0 {
		goto L27
	} else {
		goto L273
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1289 = m.ExcPending
	if v1289 != 0 {
		goto L27
	} else {
		goto L268
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1264 = m.ExcPending
	if v1264 != 0 {
		goto L27
	} else {
		goto L263
	}
L9:
	;
	m.G0 = v31 + int32(128)
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
	v40 = v31 + int32(118)
	v41 = m.G0
	v43 = v41 - int32(16)
	m.G0 = v43
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+119)))
	switch v48 - int32(112) {
	case 0, 2:
		goto L13
	case 1:
		v61 = v3
		goto L12
	default:
		goto L14
	}
L12:
	;
	if v61 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L13:
	;
	goto L16
L14:
	;
	if v48 != int32(83) {
		v61 = v3
		goto L12
	} else {
		goto L15
	}
L15:
	;
	goto L13
L16:
	;
	if base.Ui32(v45) < base.Ui32(int32(_a_F_CheckCmdReplicaIdentity_0)) {
		v61 = v3
		goto L12
	} else {
		goto L17
	}
L17:
	;
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+118)))
	v61 = base.B2i32(v55 == int32(112)) & base.B2i32(base.Ui32(int32(_a_F_CheckCmdReplicaIdentity_1)) < base.Ui32(v45))
	goto L12
L18:
	;
	v1169 = base.B2i32(l1 != int32(2))
	if v1169 == int32(0) {
		goto L230
	} else {
		goto L231
	}
L19:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1158 = m.ExcPending
	if v1158 != 0 {
		goto L27
	} else {
		goto L227
	}
L20:
	;
	m.G0 = v43 + int32(16)
	goto L18
L21:
	;
	v64 = int32(257)
	*(*uint16)(unsafe.Add(mBase, uint32(v40)+8)) = uint16(v64)
	*(*int64)(unsafe.Add(mBase, uint32(v40))) = int64(72340172821233664)
	goto L20
L22:
	;
	goto L23
L23:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+176))
	if v68 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v69 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v68)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(v40)+8)) = uint16(v69)
	v71 = *(*int64)(unsafe.Add(mBase, uint32(v68)))
	*(*int64)(unsafe.Add(mBase, uint32(v40))) = v71
	goto L20
L25:
	;
	goto L26
L26:
	;
	v73 = int32(257)
	*(*uint16)(unsafe.Add(mBase, uint32(v40)+8)) = uint16(v73)
	*(*int64)(unsafe.Add(mBase, uint32(v40))) = int64(72340172821233664)
	v77 = F_GetRelationIncludedPublications(m, v45)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	return
L28:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)+68))
	v81 = F_GetSchemaPublications(m, v80)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L27
	} else {
		goto L29
	}
L29:
	;
	v83 = F_list_concat_unique_oid(m, v77, v81)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L27
	} else {
		goto L30
	}
L30:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85)+131)))
	if v86 != int32(1) {
		v182 = v45
		v190 = v3
		v192 = v83
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v207 = F_GetRelationExcludedPublications(m, v182)
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L27
	} else {
		goto L46
	}
L32:
	;
	v89 = F_get_partition_ancestors(m, v45)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L27
	} else {
		goto L33
	}
L33:
	;
	if v89 == int32(0) {
		v182 = v45
		v190 = v3
		v192 = v83
		goto L31
	} else {
		goto L34
	}
L34:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v89)+12))
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v89)+4))
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v93+v94<<(uint(int32(2))%32)-int32(4))))
	if int32(0) < v94 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v115 = int32(0)
	v117 = v83
	goto L38
L36:
	;
	v164 = v83
	goto L37
L37:
	;
	v182 = v100
	v190 = v89
	v192 = v164
	goto L31
L38:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v89)+12))
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v132+v115<<(uint(int32(2))%32))))
	v137 = F_GetRelationIncludedPublications(m, v136)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L27
	} else {
		goto L40
	}
L39:
	;
	v164 = v145
	goto L37
L40:
	;
	v139 = F_list_concat_unique_oid(m, v117, v137)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L27
	} else {
		goto L41
	}
L41:
	;
	v141 = F_get_rel_namespace(m, v136)
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L27
	} else {
		goto L42
	}
L42:
	;
	v143 = F_GetSchemaPublications(m, v141)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L27
	} else {
		goto L43
	}
L43:
	;
	v145 = F_list_concat_unique_oid(m, v139, v143)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L27
	} else {
		goto L44
	}
L44:
	;
	v148 = v115 + int32(1)
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v89)+4))
	if v148 < v149 {
		v115 = v148
		v117 = v145
		goto L38
	} else {
		goto L45
	}
L45:
	;
	goto L39
L46:
	;
	v210 = m.G0
	v212 = v210 + int32(-64)
	m.G0 = v212
	v216 = F_table_open(m, int32(_a_F_CheckCmdReplicaIdentity_2), int32(1))
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L27
	} else {
		goto L47
	}
L47:
	;
	v219 = v210 + int32(-56)
	F_ScanKeyInit(m, v219, int32(4), int32(3), int32(60), int64(1))
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L27
	} else {
		goto L48
	}
L48:
	;
	v226 = int32(0)
	v230 = F_systable_beginscan(m, v216, v226, v226, v226, int32(1), v219)
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L27
	} else {
		goto L49
	}
L49:
	;
	v234 = int32(0)
	goto L50
L50:
	;
	v260 = F_systable_getnext(m, v230)
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L27
	} else {
		goto L52
	}
L51:
	;
	F_systable_endscan(m, v230)
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L27
	} else {
		goto L57
	}
L52:
	;
	if v260 != 0 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v260)+16))
	v263 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v262)+22)))
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v262+v263)))
	v266 = F_lappend_oid(m, v234, v265)
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L27
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	goto L51
L56:
	;
	v234 = v266
	goto L50
L57:
	;
	F_relation_close(m, v216, int32(1))
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L27
	} else {
		goto L58
	}
L58:
	;
	m.G0 = v212 - int32(-64)
	v276 = F_list_difference_oid(m, v234, v207)
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L27
	} else {
		goto L60
	}
L59:
	;
	v1103 = *(*int32)(unsafe.Add(mBase, uint32(l0)+176))
	if v1103 != 0 {
		goto L222
	} else {
		goto L223
	}
L60:
	;
	v278 = F_list_concat_unique_oid(m, v192, v276)
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L27
	} else {
		goto L61
	}
L61:
	;
	if v278 == int32(0) {
		goto L59
	} else {
		goto L62
	}
L62:
	;
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v278)+4))
	if v282 <= int32(0) {
		goto L59
	} else {
		goto L63
	}
L63:
	;
	v286 = v31 + int32(122)
	v301 = int32(0)
	goto L64
L64:
	;
	v317 = *(*int32)(unsafe.Add(mBase, uint32(v278)+12))
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v317+v301<<(uint(int32(2))%32))))
	v323 = F_SearchSysCache1(m, int32(51), base.I64_extend_i32_u(v321))
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L27
	} else {
		goto L66
	}
L65:
	;
	goto L59
L66:
	;
	if v323 == int32(0) {
		goto L19
	} else {
		goto L67
	}
L67:
	;
	v327 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40))))
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v323)+16))
	v329 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v328)+22)))
	v330 = v328 + v329
	v331 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v330)+74)))
	v332 = v327 | v331
	*(*uint8)(unsafe.Add(mBase, uint32(v40))) = uint8(v332)
	v334 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+1)))
	v335 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v330)+75)))
	v336 = v334 | v335
	*(*uint8)(unsafe.Add(mBase, uint32(v40)+1)) = uint8(v336)
	v338 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+2)))
	v339 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v330)+76)))
	v340 = v338 | v339
	*(*uint8)(unsafe.Add(mBase, uint32(v40)+2)) = uint8(v340)
	v342 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+3)))
	v343 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v330)+77)))
	v344 = v342 | v343
	*(*uint8)(unsafe.Add(mBase, uint32(v40)+3)) = uint8(v344)
	v346 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v330)+72)))
	if v346 != 0 {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v467 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v330)+75)))
	if v467 == int32(0) {
		goto L110
	} else {
		goto L111
	}
L69:
	;
	v347 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v330)+75)))
	if v347 == int32(0) {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v350 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v330)+76)))
	if v350 != int32(1) {
		goto L68
	} else {
		goto L73
	}
L71:
	;
	goto L72
L72:
	;
	v353 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v330)+78)))
	v354 = int32(0)
	v355 = m.G0
	v357 = v355 - int32(32)
	m.G0 = v357
	v359 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v360 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v359)+130)))
	if v360 == int32(102) {
		v438 = v354
		goto L74
	} else {
		goto L75
	}
L73:
	;
	goto L72
L74:
	;
	m.G0 = v357 + int32(32)
	if v438 == int32(0) {
		goto L68
	} else {
		goto L104
	}
L75:
	;
	v364 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	if v353 == int32(0) {
		v374 = v364
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v377 = F_SearchSysCache2(m, int32(53), base.I64_extend_i32_u(v374), base.I64_extend_i32_u(v321))
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L27
	} else {
		goto L83
	}
L77:
	;
	v367 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v359)+131)))
	if v367 != int32(1) {
		v374 = v364
		goto L76
	} else {
		goto L78
	}
L78:
	;
	v370 = F_GetTopMostAncestorInPublication(m, v321, v190)
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
		goto L27
	} else {
		goto L79
	}
L79:
	;
	if v370 != 0 {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v372 = v370
	goto L82
L81:
	;
	v372 = v364
	goto L82
L82:
	;
	v374 = v372
	goto L76
L83:
	;
	if v377 == int32(0) {
		v438 = v354
		goto L74
	} else {
		goto L84
	}
L84:
	;
	v385 = F_SysCacheGetAttr(m, int32(53), v377, int32(5), v357+int32(31))
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L27
	} else {
		goto L85
	}
L85:
	;
	v387 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v357)+31)))
	if v387 != 0 {
		v430 = v354
		goto L86
	} else {
		goto L87
	}
L86:
	;
	F_ReleaseCatCache(m, v377)
	mBase = m.M
	v433 = m.ExcPending
	if v433 != 0 {
		goto L27
	} else {
		goto L103
	}
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v357)+16)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v357)+24)) = v374
	*(*int32)(unsafe.Add(mBase, uint32(v357)+20)) = v364
	*(*uint8)(unsafe.Add(mBase, uint32(v357)+16)) = uint8(v353)
	v394 = F_RelationGetIndexAttrBitmap(m, l0, int32(2))
	mBase = m.M
	v395 = m.ExcPending
	if v395 != 0 {
		goto L27
	} else {
		goto L88
	}
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v357)+12)) = v394
	v398 = F_text_to_cstring(m, base.I32_wrap_i64(v385))
	mBase = m.M
	v399 = m.ExcPending
	if v399 != 0 {
		goto L27
	} else {
		goto L89
	}
L89:
	;
	v400 = F_stringToNode(m, v398)
	mBase = m.M
	v401 = m.ExcPending
	if v401 != 0 {
		goto L27
	} else {
		goto L90
	}
L90:
	;
	if v400 == int32(0) {
		v430 = v354
		goto L86
	} else {
		goto L91
	}
L91:
	;
	v404 = *(*int32)(unsafe.Add(mBase, uint32(v400)))
	if v404 == int32(6) {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	v407 = int32(*(*int16)(unsafe.Add(mBase, uint32(v400)+8)))
	if v353 != 0 {
		goto L95
	} else {
		goto L96
	}
L93:
	;
	goto L94
L94:
	;
	v426 = F_expression_tree_walker_impl(m, v400, int32(609), v357+int32(12))
	mBase = m.M
	v427 = m.ExcPending
	if v427 != 0 {
		goto L27
	} else {
		goto L102
	}
L95:
	;
	v409 = F_get_attname(m, v374, v407, int32(0))
	mBase = m.M
	v410 = m.ExcPending
	if v410 != 0 {
		goto L27
	} else {
		goto L98
	}
L96:
	;
	v413 = v407
	goto L97
L97:
	;
	v417 = F_bms_is_member(m, v413+int32(7), v394)
	mBase = m.M
	v418 = m.ExcPending
	if v418 != 0 {
		goto L27
	} else {
		goto L100
	}
L98:
	;
	v411 = F_get_attnum(m, v364, v409)
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L27
	} else {
		goto L99
	}
L99:
	;
	v413 = v411
	goto L97
L100:
	;
	if v417 == int32(0) {
		v430 = int32(1)
		goto L86
	} else {
		goto L101
	}
L101:
	;
	goto L94
L102:
	;
	v430 = v426
	goto L86
L103:
	;
	v438 = v430
	goto L74
L104:
	;
	v447 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v330)+75)))
	if v447 == int32(1) {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	v450 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v286))) = uint8(v450)
	goto L107
L106:
	;
	goto L107
L107:
	;
	v452 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v330)+76)))
	if v452 != int32(1) {
		goto L68
	} else {
		goto L108
	}
L108:
	;
	v455 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v40)+5)) = uint8(v455)
	goto L68
L109:
	;
	F_ReleaseCatCache(m, v323)
	mBase = m.M
	v1034 = m.ExcPending
	if v1034 != 0 {
		goto L27
	} else {
		goto L203
	}
L110:
	;
	v470 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v330)+76)))
	if v470 != int32(1) {
		goto L109
	} else {
		goto L113
	}
L111:
	;
	goto L112
L112:
	;
	v473 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v330)+78)))
	v474 = int32(*(*int8)(unsafe.Add(mBase, uint32(v330)+79)))
	v475 = m.G0
	v477 = v475 - int32(16)
	m.G0 = v477
	v479 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v480 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v477)+12)) = v480
	v482 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v484 = v43 + int32(15)
	*(*uint8)(unsafe.Add(mBase, uint32(v484))) = uint8(v480)
	v488 = v43 + int32(14)
	*(*uint8)(unsafe.Add(mBase, uint32(v488))) = uint8(v480)
	if v473 == v480 {
		v501 = v479
		goto L114
	} else {
		goto L115
	}
L113:
	;
	goto L112
L114:
	;
	v502 = F_GetPublication(m, v321)
	mBase = m.M
	v503 = m.ExcPending
	if v503 != 0 {
		goto L27
	} else {
		goto L121
	}
L115:
	;
	v493 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v494 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v493)+131)))
	if v494 != int32(1) {
		v501 = v479
		goto L114
	} else {
		goto L116
	}
L116:
	;
	v497 = F_GetTopMostAncestorInPublication(m, v321, v190)
	mBase = m.M
	v498 = m.ExcPending
	if v498 != 0 {
		goto L27
	} else {
		goto L117
	}
L117:
	;
	if v497 != 0 {
		goto L118
	} else {
		goto L119
	}
L118:
	;
	v499 = v497
	goto L120
L119:
	;
	v499 = v479
	goto L120
L120:
	;
	v501 = v499
	goto L114
L121:
	;
	v505 = v477 + int32(12)
	v506 = m.G0
	v508 = v506 - int32(16)
	m.G0 = v508
	v510 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v502)+8)))
	if v510 != 0 {
		goto L122
	} else {
		goto L123
	}
L122:
	;
	m.G0 = v508 + int32(16)
	v679 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v680 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v679)+130)))
	if v680 != int32(102) {
		goto L143
	} else {
		goto L144
	}
L123:
	;
	v513 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v502))))
	v514 = F_SearchSysCache2(m, int32(53), base.I64_extend_i32_u(v501), v513)
	mBase = m.M
	v515 = m.ExcPending
	if v515 != 0 {
		goto L27
	} else {
		goto L124
	}
L124:
	;
	if v514 == int32(0) {
		goto L122
	} else {
		goto L125
	}
L125:
	;
	v522 = F_SysCacheGetAttr(m, int32(53), v514, int32(6), v508+int32(15))
	mBase = m.M
	v523 = m.ExcPending
	if v523 != 0 {
		goto L27
	} else {
		goto L126
	}
L126:
	;
	v524 = int32(0)
	v526 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v508)+15)))
	if base.B2i32(v505 == v524)|v526&int32(1) == v524 {
		goto L127
	} else {
		goto L128
	}
L127:
	;
	v532 = *(*int32)(unsafe.Add(mBase, uint32(v505)))
	v534 = F_pg_detoast_datum(m, base.I32_wrap_i64(v522))
	mBase = m.M
	v535 = m.ExcPending
	if v535 != 0 {
		goto L27
	} else {
		goto L130
	}
L128:
	;
	goto L129
L129:
	;
	F_ReleaseCatCache(m, v514)
	mBase = m.M
	v647 = m.ExcPending
	if v647 != 0 {
		goto L27
	} else {
		goto L141
	}
L130:
	;
	v536 = *(*int32)(unsafe.Add(mBase, uint32(v534)+8))
	if v536 == int32(0) {
		goto L131
	} else {
		goto L132
	}
L131:
	;
	v539 = *(*int32)(unsafe.Add(mBase, uint32(v534)+4))
	v546 = (v539<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L133
L132:
	;
	v546 = v536
	goto L133
L133:
	;
	v547 = int32(0)
	v548 = *(*int32)(unsafe.Add(mBase, uint32(v534)+16))
	if v547 < v548 {
		goto L134
	} else {
		goto L135
	}
L134:
	;
	v554 = v547
	v566 = v532
	goto L137
L135:
	;
	v603 = v532
	goto L136
L136:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v505))) = v603
	goto L129
L137:
	;
	v583 = int32(*(*int16)(unsafe.Add(mBase, uint32(v534+v546+v554<<(uint(int32(1))%32)))))
	v584 = F_bms_add_member(m, v566, v583)
	mBase = m.M
	v585 = m.ExcPending
	if v585 != 0 {
		goto L27
	} else {
		goto L139
	}
L138:
	;
	v603 = v584
	goto L136
L139:
	;
	v587 = v554 + int32(1)
	if v587 != v548 {
		v554 = v587
		v566 = v584
		goto L137
	} else {
		goto L140
	}
L140:
	;
	goto L138
L141:
	;
	goto L122
L142:
	;
	m.G0 = v477 + int32(16)
	if v951&int32(1) == int32(0) {
		goto L109
	} else {
		goto L198
	}
L143:
	;
	v717 = F_RelationGetIndexAttrBitmap(m, l0, int32(2))
	mBase = m.M
	v718 = m.ExcPending
	if v718 != 0 {
		goto L27
	} else {
		goto L155
	}
L144:
	;
	v683 = *(*int32)(unsafe.Add(mBase, uint32(v477)+12))
	*(*uint8)(unsafe.Add(mBase, uint32(v484))) = uint8(base.B2i32(v683 != int32(0)))
	v687 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	if v474 == int32(115) {
		v700 = v687
		goto L145
	} else {
		goto L146
	}
L145:
	;
	v701 = *(*int32)(unsafe.Add(mBase, uint32(v700)+24))
	if v701 == int32(0) {
		goto L149
	} else {
		goto L150
	}
L146:
	;
	v690 = *(*int32)(unsafe.Add(mBase, uint32(v687)+24))
	if v690 == int32(0) {
		v700 = v687
		goto L145
	} else {
		goto L147
	}
L147:
	;
	v693 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v690)+17)))
	if v693 != int32(1) {
		v700 = v687
		goto L145
	} else {
		goto L148
	}
L148:
	;
	v696 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v488))) = uint8(v696)
	v698 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v700 = v698
	goto L145
L149:
	;
	v709 = int32(1)
	v710 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v488))))
	if v710 != v709 {
		goto L143
	} else {
		goto L152
	}
L150:
	;
	v704 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v701)+18)))
	if v704 != int32(1) {
		goto L149
	} else {
		goto L151
	}
L151:
	;
	v707 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v488))) = uint8(v707)
	goto L149
L152:
	;
	v713 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v484))))
	if v713 != 0 {
		v951 = v709
		goto L142
	} else {
		goto L153
	}
L153:
	;
	goto L143
L154:
	;
	v940 = *(*int32)(unsafe.Add(mBase, uint32(v477)+12))
	F_bms_free(m, v940)
	mBase = m.M
	v942 = m.ExcPending
	if v942 != 0 {
		goto L27
	} else {
		goto L195
	}
L155:
	;
	if v717 == int32(0) {
		goto L158
	} else {
		goto L159
	}
L156:
	;
	if v775 < int32(0) {
		goto L154
	} else {
		goto L167
	}
L157:
	;
	v775 = base.I32_ctz(v761) | v762<<(uint(int32(5))%32)
	goto L156
L158:
	;
	v775 = int32(-2)
	goto L156
L159:
	;
	v726 = int32(0)
	v729 = *(*int32)(unsafe.Add(mBase, uint32(v717)+4))
	if v729 <= v726 {
		goto L158
	} else {
		goto L160
	}
L160:
	;
	v732 = v717 + int32(8)
	v736 = *(*int32)(unsafe.Add(mBase, uint32(v732)))
	v739 = v736 & int32(-1)
	if v739 != 0 {
		v761 = v739
		v762 = v726
		goto L157
	} else {
		goto L161
	}
L161:
	;
	v740 = int32(1)
	if v740 == v729 {
		goto L158
	} else {
		goto L162
	}
L162:
	;
	v744 = v740
	goto L163
L163:
	;
	v751 = *(*int32)(unsafe.Add(mBase, uint32(v732+v744<<(uint(int32(2))%32))))
	if v751 != 0 {
		v761 = v751
		v762 = v744
		goto L157
	} else {
		goto L165
	}
L164:
	;
	goto L158
L165:
	;
	v753 = v744 + int32(1)
	if v753 != v729 {
		v744 = v753
		goto L163
	} else {
		goto L166
	}
L166:
	;
	goto L164
L167:
	;
	v788 = v775
	goto L168
L168:
	;
	v808 = base.I32_extend16_s(v788 - int32(7))
	v809 = *(*int32)(unsafe.Add(mBase, uint32(v477)+12))
	if v809 == int32(0) {
		goto L171
	} else {
		goto L172
	}
L169:
	;
	goto L154
L170:
	;
	if v717 == int32(0) {
		goto L185
	} else {
		goto L186
	}
L171:
	;
	v812 = *(*int32)(unsafe.Add(mBase, uint32(v482)))
	v819 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v482+v812<<(uint(int32(3))%32)+v808*int32(100))+18)))
	v820 = int32(115)
	if base.B2i32(base.B2i32(v819 == v820)&base.B2i32(v474 != v820) == int32(0))&base.B2i32(v819 != int32(118)) != 0 {
		goto L170
	} else {
		goto L174
	}
L172:
	;
	goto L173
L173:
	;
	if v473 != 0 {
		goto L175
	} else {
		goto L176
	}
L174:
	;
	v830 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v488))) = uint8(v830)
	goto L154
L175:
	;
	v833 = F_get_attname(m, v479, v808, int32(0))
	mBase = m.M
	v834 = m.ExcPending
	if v834 != 0 {
		goto L27
	} else {
		goto L178
	}
L176:
	;
	v838 = v809
	v839 = v808
	goto L177
L177:
	;
	v840 = F_bms_is_member(m, v839, v838)
	mBase = m.M
	v841 = m.ExcPending
	if v841 != 0 {
		goto L27
	} else {
		goto L180
	}
L178:
	;
	v835 = F_get_attnum(m, v501, v833)
	mBase = m.M
	v836 = m.ExcPending
	if v836 != 0 {
		goto L27
	} else {
		goto L179
	}
L179:
	;
	v837 = *(*int32)(unsafe.Add(mBase, uint32(v477)+12))
	v838 = v837
	v839 = v835
	goto L177
L180:
	;
	v842 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v484))))
	v843 = int32(1)
	v845 = v842 | (v840 ^ v843)
	*(*uint8)(unsafe.Add(mBase, uint32(v484))) = uint8(v845)
	if v845&v843 == int32(0) {
		goto L170
	} else {
		goto L181
	}
L181:
	;
	v851 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v488))))
	if v851 != 0 {
		goto L154
	} else {
		goto L182
	}
L182:
	;
	goto L170
L183:
	;
	if int32(0) <= v909 {
		v788 = v909
		goto L168
	} else {
		goto L194
	}
L184:
	;
	v909 = base.I32_ctz(v895) | v896<<(uint(int32(5))%32)
	goto L183
L185:
	;
	v909 = int32(-2)
	goto L183
L186:
	;
	v860 = v788 + int32(1)
	v862 = int32(base.Ui32(v860) >> (uint(int32(5)) % 32))
	v863 = *(*int32)(unsafe.Add(mBase, uint32(v717)+4))
	if v863 <= v862 {
		goto L185
	} else {
		goto L187
	}
L187:
	;
	v866 = v717 + int32(8)
	v870 = *(*int32)(unsafe.Add(mBase, uint32(v866+v862<<(uint(int32(2))%32))))
	v873 = v870 & (int32(-1) << (uint(v860) % 32))
	if v873 != 0 {
		v895 = v873
		v896 = v862
		goto L184
	} else {
		goto L188
	}
L188:
	;
	v875 = v862 + int32(1)
	if v875 == v863 {
		goto L185
	} else {
		goto L189
	}
L189:
	;
	v878 = v875
	goto L190
L190:
	;
	v885 = *(*int32)(unsafe.Add(mBase, uint32(v866+v878<<(uint(int32(2))%32))))
	if v885 != 0 {
		v895 = v885
		v896 = v878
		goto L184
	} else {
		goto L192
	}
L191:
	;
	goto L185
L192:
	;
	v887 = v878 + int32(1)
	if v887 != v863 {
		v878 = v887
		goto L190
	} else {
		goto L193
	}
L193:
	;
	goto L191
L194:
	;
	goto L169
L195:
	;
	F_bms_free(m, v717)
	mBase = m.M
	v944 = m.ExcPending
	if v944 != 0 {
		goto L27
	} else {
		goto L196
	}
L196:
	;
	v946 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v484))))
	if v946 != 0 {
		v951 = int32(1)
		goto L142
	} else {
		goto L197
	}
L197:
	;
	v947 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v488))))
	v951 = v947
	goto L142
L198:
	;
	v983 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v330)+75)))
	if v983 == int32(1) {
		goto L199
	} else {
		goto L200
	}
L199:
	;
	v986 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+15)))
	v987 = int32(1)
	v988 = v986 ^ v987
	*(*uint8)(unsafe.Add(mBase, uint32(v40)+6)) = uint8(v988)
	v990 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+14)))
	v992 = v990 ^ v987
	*(*uint8)(unsafe.Add(mBase, uint32(v40)+8)) = uint8(v992)
	goto L201
L200:
	;
	goto L201
L201:
	;
	v994 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v330)+76)))
	if v994 != int32(1) {
		goto L109
	} else {
		goto L202
	}
L202:
	;
	v997 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+15)))
	v998 = int32(1)
	v999 = v997 ^ v998
	*(*uint8)(unsafe.Add(mBase, uint32(v40)+7)) = uint8(v999)
	v1001 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+14)))
	v1003 = v1001 ^ v998
	*(*uint8)(unsafe.Add(mBase, uint32(v40)+9)) = uint8(v1003)
	goto L109
L203:
	;
	v1035 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40))))
	if v1035 != int32(1) {
		goto L204
	} else {
		goto L205
	}
L204:
	;
	v1072 = v301 + int32(1)
	v1073 = *(*int32)(unsafe.Add(mBase, uint32(v278)+4))
	if v1072 < v1073 {
		v301 = v1072
		goto L64
	} else {
		goto L221
	}
L205:
	;
	v1038 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+1)))
	if v1038 != int32(1) {
		goto L204
	} else {
		goto L206
	}
L206:
	;
	v1041 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+2)))
	if v1041 != int32(1) {
		goto L207
	} else {
		goto L208
	}
L207:
	;
	v1051 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+2)))
	if v1051 != int32(1) {
		goto L212
	} else {
		goto L213
	}
L208:
	;
	v1044 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+3)))
	if v1044 != int32(1) {
		goto L207
	} else {
		goto L209
	}
L209:
	;
	v1047 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v286))))
	if v1047 != 0 {
		goto L207
	} else {
		goto L210
	}
L210:
	;
	v1048 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+5)))
	if v1048 != int32(1) {
		goto L59
	} else {
		goto L211
	}
L211:
	;
	goto L207
L212:
	;
	v1061 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+2)))
	if v1061 != int32(1) {
		goto L204
	} else {
		goto L217
	}
L213:
	;
	v1054 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+3)))
	if v1054 != int32(1) {
		goto L212
	} else {
		goto L214
	}
L214:
	;
	v1057 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+6)))
	if v1057 != 0 {
		goto L212
	} else {
		goto L215
	}
L215:
	;
	v1058 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+7)))
	if v1058 != int32(1) {
		goto L59
	} else {
		goto L216
	}
L216:
	;
	goto L212
L217:
	;
	v1064 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+3)))
	if v1064 != int32(1) {
		goto L204
	} else {
		goto L218
	}
L218:
	;
	v1067 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+8)))
	if v1067 != 0 {
		goto L204
	} else {
		goto L219
	}
L219:
	;
	v1068 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+9)))
	if v1068 != int32(1) {
		goto L59
	} else {
		goto L220
	}
L220:
	;
	goto L204
L221:
	;
	goto L65
L222:
	;
	F_pfree(m, v1103)
	mBase = m.M
	v1105 = m.ExcPending
	if v1105 != 0 {
		goto L27
	} else {
		goto L225
	}
L223:
	;
	goto L224
L224:
	;
	v1108 = int32(_a_F_CheckCmdReplicaIdentity_3)
	v1109 = *(*int32)(unsafe.Add(mBase, _c_F_CheckCmdReplicaIdentity[0]))
	v1112 = *(*int32)(unsafe.Add(mBase, _c_F_CheckCmdReplicaIdentity[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_CheckCmdReplicaIdentity[0])) = v1112
	v1115 = F_palloc(m, int32(10))
	mBase = m.M
	v1116 = m.ExcPending
	if v1116 != 0 {
		goto L27
	} else {
		goto L226
	}
L225:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+176)) = int32(0)
	goto L224
L226:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+176)) = v1115
	v1118 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v40)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(v1115)+8)) = uint16(v1118)
	v1120 = *(*int64)(unsafe.Add(mBase, uint32(v40)))
	*(*int64)(unsafe.Add(mBase, uint32(v1115))) = v1120
	*(*int32)(unsafe.Add(mBase, _c_F_CheckCmdReplicaIdentity[0])) = v1109
	goto L20
L227:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43))) = v321
	F_errmsg_internal(m, int32(_a_F_CheckCmdReplicaIdentity_4), v43)
	mBase = m.M
	v1162 = m.ExcPending
	if v1162 != 0 {
		goto L27
	} else {
		goto L228
	}
L228:
	;
	F_errfinish(m, int32(_a_F_CheckCmdReplicaIdentity_5), int32(_a_F_CheckCmdReplicaIdentity_6), int32(_a_F_CheckCmdReplicaIdentity_7))
	mBase = m.M
	v1167 = m.ExcPending
	if v1167 != 0 {
		goto L27
	} else {
		goto L229
	}
L229:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L230:
	;
	v1172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+122)))
	if v1172&int32(1) == int32(0) {
		goto L8
	} else {
		goto L233
	}
L231:
	;
	goto L232
L232:
	;
	if v1169 == int32(0) {
		goto L234
	} else {
		goto L235
	}
L233:
	;
	goto L232
L234:
	;
	v1179 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+124)))
	if v1179&int32(1) == int32(0) {
		goto L7
	} else {
		goto L237
	}
L235:
	;
	goto L236
L236:
	;
	if l1 == int32(2) {
		goto L238
	} else {
		goto L239
	}
L237:
	;
	goto L236
L238:
	;
	v1186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+126)))
	if v1186&int32(1) == int32(0) {
		goto L6
	} else {
		goto L241
	}
L239:
	;
	goto L240
L240:
	;
	v1192 = base.B2i32(l1 != int32(4))
	if v1192 == int32(0) {
		goto L242
	} else {
		goto L243
	}
L241:
	;
	goto L240
L242:
	;
	v1195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+123)))
	if v1195&int32(1) == int32(0) {
		goto L5
	} else {
		goto L245
	}
L243:
	;
	goto L244
L244:
	;
	if v1192 == int32(0) {
		goto L246
	} else {
		goto L247
	}
L245:
	;
	goto L244
L246:
	;
	v1202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+125)))
	if v1202&int32(1) == int32(0) {
		goto L4
	} else {
		goto L249
	}
L247:
	;
	goto L248
L248:
	;
	if l1 == int32(4) {
		goto L250
	} else {
		goto L251
	}
L249:
	;
	goto L248
L250:
	;
	v1209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+127)))
	if v1209&int32(1) == int32(0) {
		goto L3
	} else {
		goto L253
	}
L251:
	;
	goto L252
L252:
	;
	v1214 = F_RelationGetReplicaIndex(m, l0)
	mBase = m.M
	v1215 = m.ExcPending
	if v1215 != 0 {
		goto L27
	} else {
		goto L254
	}
L253:
	;
	goto L252
L254:
	;
	if v1214 != 0 {
		goto L9
	} else {
		goto L255
	}
L255:
	;
	v1216 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v1217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1216)+130)))
	if v1217 == int32(102) {
		goto L9
	} else {
		goto L256
	}
L256:
	;
	if l1 == int32(2) {
		goto L257
	} else {
		goto L258
	}
L257:
	;
	v1222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+119)))
	if v1222&int32(1) != 0 {
		goto L2
	} else {
		goto L260
	}
L258:
	;
	goto L259
L259:
	;
	if l1 != int32(4) {
		goto L9
	} else {
		goto L261
	}
L260:
	;
	goto L259
L261:
	;
	v1227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+120)))
	if v1227&int32(1) != 0 {
		goto L1
	} else {
		goto L262
	}
L262:
	;
	goto L9
L263:
	;
	F_errcode(m, int32(_a_F_CheckCmdReplicaIdentity_8))
	mBase = m.M
	v1267 = m.ExcPending
	if v1267 != 0 {
		goto L27
	} else {
		goto L264
	}
L264:
	;
	v1268 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+112)) = v1268 + int32(4)
	F_errmsg(m, int32(_a_F_CheckCmdReplicaIdentity_9), v31+int32(112))
	mBase = m.M
	v1276 = m.ExcPending
	if v1276 != 0 {
		goto L27
	} else {
		goto L265
	}
L265:
	;
	v1279 = F_errdetail(m, int32(_a_F_CheckCmdReplicaIdentity_10), int32(0))
	mBase = m.M
	v1280 = m.ExcPending
	if v1280 != 0 {
		goto L27
	} else {
		goto L266
	}
L266:
	;
	F_errfinish(m, int32(_a_F_CheckCmdReplicaIdentity_11), int32(1073), int32(_a_F_CheckCmdReplicaIdentity_12))
	mBase = m.M
	v1285 = m.ExcPending
	if v1285 != 0 {
		goto L27
	} else {
		goto L267
	}
L267:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L268:
	;
	F_errcode(m, int32(_a_F_CheckCmdReplicaIdentity_8))
	mBase = m.M
	v1292 = m.ExcPending
	if v1292 != 0 {
		goto L27
	} else {
		goto L269
	}
L269:
	;
	v1293 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+96)) = v1293 + int32(4)
	F_errmsg(m, int32(_a_F_CheckCmdReplicaIdentity_9), v31+int32(96))
	mBase = m.M
	v1301 = m.ExcPending
	if v1301 != 0 {
		goto L27
	} else {
		goto L270
	}
L270:
	;
	v1304 = F_errdetail(m, int32(_a_F_CheckCmdReplicaIdentity_13), int32(0))
	mBase = m.M
	v1305 = m.ExcPending
	if v1305 != 0 {
		goto L27
	} else {
		goto L271
	}
L271:
	;
	F_errfinish(m, int32(_a_F_CheckCmdReplicaIdentity_11), int32(1079), int32(_a_F_CheckCmdReplicaIdentity_12))
	mBase = m.M
	v1310 = m.ExcPending
	if v1310 != 0 {
		goto L27
	} else {
		goto L272
	}
L272:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L273:
	;
	F_errcode(m, int32(_a_F_CheckCmdReplicaIdentity_8))
	mBase = m.M
	v1317 = m.ExcPending
	if v1317 != 0 {
		goto L27
	} else {
		goto L274
	}
L274:
	;
	v1318 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+80)) = v1318 + int32(4)
	F_errmsg(m, int32(_a_F_CheckCmdReplicaIdentity_9), v31+int32(80))
	mBase = m.M
	v1326 = m.ExcPending
	if v1326 != 0 {
		goto L27
	} else {
		goto L275
	}
L275:
	;
	v1329 = F_errdetail(m, int32(_a_F_CheckCmdReplicaIdentity_14), int32(0))
	mBase = m.M
	v1330 = m.ExcPending
	if v1330 != 0 {
		goto L27
	} else {
		goto L276
	}
L276:
	;
	F_errfinish(m, int32(_a_F_CheckCmdReplicaIdentity_11), int32(1085), int32(_a_F_CheckCmdReplicaIdentity_12))
	mBase = m.M
	v1335 = m.ExcPending
	if v1335 != 0 {
		goto L27
	} else {
		goto L277
	}
L277:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L278:
	;
	F_errcode(m, int32(_a_F_CheckCmdReplicaIdentity_8))
	mBase = m.M
	v1342 = m.ExcPending
	if v1342 != 0 {
		goto L27
	} else {
		goto L279
	}
L279:
	;
	v1343 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+64)) = v1343 + int32(4)
	F_errmsg(m, int32(_a_F_CheckCmdReplicaIdentity_15), v31-int32(-64))
	mBase = m.M
	v1351 = m.ExcPending
	if v1351 != 0 {
		goto L27
	} else {
		goto L280
	}
L280:
	;
	v1354 = F_errdetail(m, int32(_a_F_CheckCmdReplicaIdentity_10), int32(0))
	mBase = m.M
	v1355 = m.ExcPending
	if v1355 != 0 {
		goto L27
	} else {
		goto L281
	}
L281:
	;
	F_errfinish(m, int32(_a_F_CheckCmdReplicaIdentity_11), int32(1091), int32(_a_F_CheckCmdReplicaIdentity_12))
	mBase = m.M
	v1360 = m.ExcPending
	if v1360 != 0 {
		goto L27
	} else {
		goto L282
	}
L282:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L283:
	;
	F_errcode(m, int32(_a_F_CheckCmdReplicaIdentity_8))
	mBase = m.M
	v1367 = m.ExcPending
	if v1367 != 0 {
		goto L27
	} else {
		goto L284
	}
L284:
	;
	v1368 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+48)) = v1368 + int32(4)
	F_errmsg(m, int32(_a_F_CheckCmdReplicaIdentity_15), v31+int32(48))
	mBase = m.M
	v1376 = m.ExcPending
	if v1376 != 0 {
		goto L27
	} else {
		goto L285
	}
L285:
	;
	v1379 = F_errdetail(m, int32(_a_F_CheckCmdReplicaIdentity_13), int32(0))
	mBase = m.M
	v1380 = m.ExcPending
	if v1380 != 0 {
		goto L27
	} else {
		goto L286
	}
L286:
	;
	F_errfinish(m, int32(_a_F_CheckCmdReplicaIdentity_11), int32(1097), int32(_a_F_CheckCmdReplicaIdentity_12))
	mBase = m.M
	v1385 = m.ExcPending
	if v1385 != 0 {
		goto L27
	} else {
		goto L287
	}
L287:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L288:
	;
	F_errcode(m, int32(_a_F_CheckCmdReplicaIdentity_8))
	mBase = m.M
	v1392 = m.ExcPending
	if v1392 != 0 {
		goto L27
	} else {
		goto L289
	}
L289:
	;
	v1393 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+32)) = v1393 + int32(4)
	F_errmsg(m, int32(_a_F_CheckCmdReplicaIdentity_15), v31+int32(32))
	mBase = m.M
	v1401 = m.ExcPending
	if v1401 != 0 {
		goto L27
	} else {
		goto L290
	}
L290:
	;
	v1404 = F_errdetail(m, int32(_a_F_CheckCmdReplicaIdentity_14), int32(0))
	mBase = m.M
	v1405 = m.ExcPending
	if v1405 != 0 {
		goto L27
	} else {
		goto L291
	}
L291:
	;
	F_errfinish(m, int32(_a_F_CheckCmdReplicaIdentity_11), int32(1103), int32(_a_F_CheckCmdReplicaIdentity_12))
	mBase = m.M
	v1410 = m.ExcPending
	if v1410 != 0 {
		goto L27
	} else {
		goto L292
	}
L292:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L293:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v1417 = m.ExcPending
	if v1417 != 0 {
		goto L27
	} else {
		goto L294
	}
L294:
	;
	v1418 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v31))) = v1418 + int32(4)
	F_errmsg(m, int32(_a_F_CheckCmdReplicaIdentity_16), v31)
	mBase = m.M
	v1424 = m.ExcPending
	if v1424 != 0 {
		goto L27
	} else {
		goto L295
	}
L295:
	;
	F_errhint(m, int32(_a_F_CheckCmdReplicaIdentity_17), int32(0))
	mBase = m.M
	v1428 = m.ExcPending
	if v1428 != 0 {
		goto L27
	} else {
		goto L296
	}
L296:
	;
	F_errfinish(m, int32(_a_F_CheckCmdReplicaIdentity_11), int32(1123), int32(_a_F_CheckCmdReplicaIdentity_12))
	mBase = m.M
	v1433 = m.ExcPending
	if v1433 != 0 {
		goto L27
	} else {
		goto L297
	}
L297:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L298:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v1440 = m.ExcPending
	if v1440 != 0 {
		goto L27
	} else {
		goto L299
	}
L299:
	;
	v1441 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+16)) = v1441 + int32(4)
	F_errmsg(m, int32(_a_F_CheckCmdReplicaIdentity_18), v31+int32(16))
	mBase = m.M
	v1449 = m.ExcPending
	if v1449 != 0 {
		goto L27
	} else {
		goto L300
	}
L300:
	;
	F_errhint(m, int32(_a_F_CheckCmdReplicaIdentity_19), int32(0))
	mBase = m.M
	v1453 = m.ExcPending
	if v1453 != 0 {
		goto L27
	} else {
		goto L301
	}
L301:
	;
	F_errfinish(m, int32(_a_F_CheckCmdReplicaIdentity_11), int32(1129), int32(_a_F_CheckCmdReplicaIdentity_12))
	mBase = m.M
	v1458 = m.ExcPending
	if v1458 != 0 {
		goto L27
	} else {
		goto L302
	}
L302:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_CheckDim_1(m *base.Module, l0 int32) {
	var v10 int32
	_ = v10
	Fn14208(m, l0, int32(109), int32(_a_F_CheckDim_1_0), int32(_a_F_CheckDim_1_1), int32(_a_F_CheckDim_1_2), int32(104), int32(_a_F_CheckDim_1_3), int32(_a_F_CheckDim_1_4))
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
func F_CheckUsageOnTypesInSingleRelExpr(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
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
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v102 int64
	_ = v102
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v139 int32
	_ = v139
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v170 int32
	_ = v170
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	v4 = int32(0)
	v11 = m.G0
	v13 = v11 - int32(160)
	m.G0 = v13
	v15 = int32(16)
	v16 = v13 + v15
	base.MemoryFill(m, v16, v4, int32(136))
	v21 = F_palloc(m, v15)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v21)+8)) = int64(137438953472)
	v27 = F_palloc_mul(m, int32(12), int32(32))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v29 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+4)) = v29
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = v27
	v32 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+40)) = v32
	v35 = int32(114)
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+37)) = uint8(v35)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+32)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v13)+28)) = v29
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = int32(101)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+152)) = v21
	*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = v16
	*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v16
	v48 = F_list_make1_impl(m, v32, v13+int32(4))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v48
	v53 = F_list_make1_impl(m, int32(1), v13)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+156)) = v53
	v58 = F_find_expr_references_walker(m, l0, v13+int32(152))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v13)+152))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)+8))
	if int32(2) <= v61 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
	F_pg_qsort(m, v64, v61, int32(12), int32(497))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L1
	} else {
		goto L10
	}
L8:
	;
	v129 = v60
	v139 = v61
	goto L9
L9:
	;
	if int32(0) < v139 {
		goto L23
	} else {
		goto L24
	}
L10:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v60)+8))
	if int32(2) <= v69 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
	v75 = v72
	v79 = int32(1)
	v80 = v32
	goto L14
L12:
	;
	v122 = v32
	goto L13
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v60)+8)) = v122
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v13)+152))
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v127)+8))
	v129 = v127
	v139 = v128
	goto L9
L14:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
	v88 = v85 + v79*int32(12)
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v88)))
	if v84 != v89 {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	v122 = v109
	goto L13
L16:
	;
	v113 = v79 + int32(1)
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v60)+8))
	if v113 < v114 {
		v75 = v108
		v79 = v113
		v80 = v109
		goto L14
	} else {
		goto L22
	}
L17:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v88)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v75)+20)) = v100
	v102 = *(*int64)(unsafe.Add(mBase, uint32(v88)))
	*(*int64)(unsafe.Add(mBase, uint32(v75)+12)) = v102
	v108 = v75 + int32(12)
	v109 = v80 + int32(1)
	goto L16
L18:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v75)+4))
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v88)+4))
	if v91 != v92 {
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v75)+8))
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v88)+8))
	if v94 == v95 {
		v108 = v75
		v109 = v80
		goto L16
	} else {
		goto L20
	}
L20:
	;
	if v94 != 0 {
		goto L17
	} else {
		goto L21
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v75)+8)) = v95
	v108 = v75
	v109 = v80
	goto L16
L22:
	;
	goto L15
L23:
	;
	v149 = v4
	goto L26
L24:
	;
	v194 = v129
	goto L25
L25:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v194)))
	F_pfree(m, v204)
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L1
	} else {
		goto L36
	}
L26:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v129)))
	v155 = v152 + v149*int32(12)
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v155)))
	if v156 != int32(1247) {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v13)+152))
	v194 = v193
	goto L25
L28:
	;
	v190 = v149 + int32(1)
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v129)+8))
	if v190 < v191 {
		v149 = v190
		goto L26
	} else {
		goto L35
	}
L29:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v155)+4))
	v170 = int32(1)
	goto L30
L30:
	;
	if base.B2i32(int32(0)|base.B2i32(base.Ui32(int32(_a_F_CheckUsageOnTypesInSingleRelExpr_0)) < base.Ui32(v160)) == int32(0))&((v170|base.B2i32(v160 != int32(2200)))&v170) != 0 {
		goto L28
	} else {
		goto L31
	}
L31:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v155)))
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v155)+4))
	v181 = F_object_aclcheck(m, v178, v179, l2, int64(256))
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	if v181 == int32(0) {
		goto L28
	} else {
		goto L33
	}
L33:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v155)+4))
	F_aclcheck_error_type(m, v181, v185)
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	goto L28
L35:
	;
	goto L27
L36:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v194)+4))
	if v207 != 0 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	F_pfree(m, v207)
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L1
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	F_pfree(m, v194)
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L1
	} else {
		goto L41
	}
L40:
	;
	goto L39
L41:
	;
	m.G0 = v13 + int32(160)
	return
}
func F_CheckpointerShmemRequest(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v27 int32
	_ = v27
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v8 = int32(_a_F_CheckpointerShmemRequest_0)
	v10 = *(*int32)(unsafe.Add(mBase, _c_F_CheckpointerShmemRequest[0]))
	if v8 <= v10 {
		v13 = v8
	} else {
		v13 = v10
	}
	v15 = F_mul_size(m, v13, int32(32))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return
	} else {
		v17 = F_add_size(m, int32(64), v15)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v5)+12)) = int32(_a_F_CheckpointerShmemRequest_1)
			*(*int32)(unsafe.Add(mBase, uint32(v5)+8)) = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v5)+4)) = v17
			*(*int32)(unsafe.Add(mBase, uint32(v5))) = int32(_a_F_CheckpointerShmemRequest_2)
			F_ShmemRequestStructWithOpts(m, v5)
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return
			} else {
				m.G0 = v5 + int32(16)
				return
			}
		}
	}
}
func F_char_decrement(m *base.Module, l0 int32, l1 int64, l2 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int64
	_ = v6
	var v7 int32
	_ = v7
	var v10 int64
	_ = v10
	var v16 int64
	_ = v16
	v6 = int64(0)
	v7 = base.B2i32(l1&int64(255) == v6)
	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v7)
	v10 = int64(56)
	if l1&int64(255) == v6 {
		v16 = v6
	} else {
		v16 = (l1<<(uint(v10)%64) - int64(72057594037927936)) >> (uint(v10) % 64)
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
func F_charhashfast(m *base.Module, l0 int64) int32 {
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	v4 = base.I32_extend8_s(base.I32_wrap_i64(l0))
	v5 = int32(16)
	v9 = (int32(base.Ui32(v4)>>(uint(v5)%32)) ^ v4) * int32(-2048144789)
	v14 = (int32(base.Ui32(v9)>>(uint(int32(13))%32)) ^ v9) * int32(-1028477387)
	return int32(base.Ui32(v14)>>(uint(v5)%32)) ^ v14
}
func F_chartoi4(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int64
	_ = v2
	v2 = int64(*(*int8)(unsafe.Add(mBase, uint32(l0)+24)))
	return v2
}
func F_checkExprHasSubLink_walker(m *base.Module, l0 int32, l1 int32) int32 {
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v5 = Fn14251(m, l0, l1, int32(1125), int32(22))
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
			v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+88)))
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
											F_errfinish(m, int32(_a_F_checkTargetlistEntrySQL92_1), int32(1967), int32(_a_F_checkTargetlistEntrySQL92_2))
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
						v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+89)))
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
														F_errfinish(m, int32(_a_F_checkTargetlistEntrySQL92_1), int32(1976), int32(_a_F_checkTargetlistEntrySQL92_2))
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
				v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+89)))
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
												F_errfinish(m, int32(_a_F_checkTargetlistEntrySQL92_1), int32(1976), int32(_a_F_checkTargetlistEntrySQL92_2))
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
					F_errfinish(m, int32(_a_F_checkTargetlistEntrySQL92_1), int32(1985), int32(_a_F_checkTargetlistEntrySQL92_2))
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
	var v318 int32
	_ = v318
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
	var v387 int32
	_ = v387
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
	F_errfinish(m, int32(_a_F_check_agglevels_and_constraints_1), int32(694), int32(_a_F_check_agglevels_and_constraints_2))
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
	*(*uint8)(unsafe.Add(mBase, uint32(v184)+88)) = uint8(v195)
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v184)+64))
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
	F_errfinish(m, int32(_a_F_check_agglevels_and_constraints_1), int32(601), int32(_a_F_check_agglevels_and_constraints_6))
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
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v184)+64))
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
	F_errfinish(m, int32(_a_F_check_agglevels_and_constraints_1), int32(616), int32(_a_F_check_agglevels_and_constraints_6))
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
	v318 = F_errdetail(m, int32(_a_F_check_agglevels_and_constraints_52), v14)
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
	F_errfinish(m, int32(_a_F_check_agglevels_and_constraints_1), int32(709), int32(_a_F_check_agglevels_and_constraints_2))
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
	F_errfinish(m, int32(_a_F_check_agglevels_and_constraints_1), int32(732), int32(_a_F_check_agglevels_and_constraints_2))
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
	F_errfinish(m, int32(_a_F_check_agglevels_and_constraints_1), int32(739), int32(_a_F_check_agglevels_and_constraints_2))
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
	v387 = F_errdetail(m, int32(_a_F_check_agglevels_and_constraints_52), v14+int32(48))
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
	F_errfinish(m, int32(_a_F_check_agglevels_and_constraints_1), int32(746), int32(_a_F_check_agglevels_and_constraints_2))
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
	var v27 int32
	_ = v27
	var v34 int32
	_ = v34
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v83 int32
	_ = v83
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	v17 = F_SearchSysCache1(m, int32(47), base.I64_extend_i32_u(l0))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	if v17 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v17)+16))
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+22)))
	v23 = v21 + v22
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+108))
	if v24 != l1 {
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
	v91 = m.ExcPending
	if v91 != 0 {
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
	v34 = int32(0)
	goto L6
L8:
	;
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+100)))
	if v26 != 0 {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v27 = int32(*(*int16)(unsafe.Add(mBase, uint32(v23)+104)))
	if v27 < l3 {
		goto L7
	} else {
		goto L10
	}
L10:
	;
	if v27 <= l4 {
		v34 = int32(1)
		goto L6
	} else {
		goto L11
	}
L11:
	;
	goto L7
L12:
	;
	v42 = v34
	v46 = int32(0)
	goto L15
L13:
	;
	v73 = v34
	goto L14
L14:
	;
	F_ReleaseCatCache(m, v17)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L1
	} else {
		goto L27
	}
L15:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v51 + int32(4)
	v55 = int32(*(*int16)(unsafe.Add(mBase, uint32(v23)+104)))
	if v55 <= v46 {
		v66 = v42
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v73 = v66
	goto L14
L17:
	;
	v70 = v46 + int32(1)
	if v70 != l4 {
		v42 = v66
		v46 = v70
		goto L15
	} else {
		goto L26
	}
L18:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v23+int32(136)+v46<<(uint(int32(2))%32))))
	if l2 != 0 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v66 = int32(0)
	goto L17
L20:
	;
	if v57 != v61 {
		goto L19
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v63 = F_IsBinaryCoercible(m, v57, v61)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L1
	} else {
		goto L24
	}
L23:
	;
	v66 = v42
	goto L17
L24:
	;
	if v63 != 0 {
		v66 = v42
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
	return v73
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = l0
	F_errmsg_internal(m, int32(_a_F_check_amproc_signature_0), v13)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	F_errfinish(m, int32(_a_F_check_amproc_signature_1), int32(163), int32(_a_F_check_amproc_signature_2))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
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
				v64 = F_expression_tree_walker_impl(m, l0, int32(530), l1)
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
						v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)+116))
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
			v61 = F_query_tree_walker_impl(m, l0, int32(530), l1, int32(0))
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
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
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
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	v4 = m.G0
	v6 = v4 - int32(32)
	m.G0 = v6
	v8 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v6)+24)) = v8
	*(*int32)(unsafe.Add(mBase, uint32(v6)+20)) = v8
	v12 = int32(1)
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v13 == v8 {
		v61 = v12
		m.G0 = v6 + int32(32)
		return v61
	} else {
		v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
		if v16 == int32(0) {
			v61 = v12
			m.G0 = v6 + int32(32)
			return v61
		} else {
			v26 = F_ReplicationSlotValidateNameInternal(m, v13, int32(0), v6+int32(28), v6+int32(24), v6+int32(20))
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return int32(0)
			} else {
				if v26 != 0 {
					v61 = v12
					m.G0 = v6 + int32(32)
					return v61
				} else {
					v30 = *(*int32)(unsafe.Add(mBase, uint32(v6)+28))
					*(*int32)(unsafe.Add(mBase, _c_F_check_primary_slot_name[0])) = v30
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
						v46 = int32(0)
						v47 = *(*int32)(unsafe.Add(mBase, uint32(v6)+20))
						if v47 == v46 {
							v61 = v46
							m.G0 = v6 + int32(32)
							return v61
						} else {
							v51 = *(*int32)(unsafe.Add(mBase, _c_F_check_primary_slot_name[1]))
							*(*int32)(unsafe.Add(mBase, _c_F_check_primary_slot_name[2])) = v51
							v54 = *(*int32)(unsafe.Add(mBase, uint32(v6)+20))
							*(*int32)(unsafe.Add(mBase, uint32(v6))) = v54
							v58 = F_format_elog_string(m, int32(_a_F_check_primary_slot_name_0), v6)
							mBase = m.M
							v59 = m.ExcPending
							if v59 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, _c_F_check_primary_slot_name[4])) = v58
								v61 = v46
								m.G0 = v6 + int32(32)
								return v61
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
