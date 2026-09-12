package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_AcquireExecutorLocks(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v10 int32
	_ = v10
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
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
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	v3 = int32(0)
	if l0 == v3 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v10 <= int32(0) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v17 = v3
	goto L4
L4:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v20+v17<<(uint(int32(2))%32))))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	if v25 == int32(6) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L1
L6:
	;
	v102 = v17 + int32(1)
	v103 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v102 < v103 {
		v17 = v102
		goto L4
	} else {
		goto L41
	}
L7:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v24)+88))
	v30 = v28
	goto L11
L8:
	;
	goto L9
L9:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v24)+44))
	if v59 == int32(0) {
		goto L6
	} else {
		goto L27
	}
L10:
	;
	if v54 == int32(0) {
		goto L6
	} else {
		goto L24
	}
L11:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	switch v32 - int32(241) {
	case 0:
		goto L16
	case 1:
		goto L15
	default:
		goto L17
	}
L13:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)+28))
	v30 = v52
	goto L11
L14:
	;
	v54 = v50
	goto L10
L15:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	if v47 == int32(6) {
		v51 = v46
		goto L13
	} else {
		goto L23
	}
L16:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
	if v43 == int32(6) {
		v51 = v42
		goto L13
	} else {
		goto L22
	}
L17:
	;
	if v32 != int32(201) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v54 = int32(0)
	goto L10
L19:
	;
	goto L20
L20:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v30)+12))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
	if v39 != int32(6) {
		v50 = v38
		goto L14
	} else {
		goto L21
	}
L21:
	;
	v51 = v38
	goto L13
L22:
	;
	v50 = v42
	goto L14
L23:
	;
	v50 = v46
	goto L14
L24:
	;
	F_ScanQueryForLocks(m, v54, l1)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	return
L26:
	;
	goto L6
L27:
	;
	v62 = int32(0)
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v59)+4))
	if v63 <= v62 {
		goto L6
	} else {
		goto L28
	}
L28:
	;
	v68 = v62
	goto L29
L29:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v59)+12))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v73+v68<<(uint(int32(2))%32))))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)+12))
	switch v78 {
	case 0:
		goto L32
	case 1:
		goto L33
	default:
		goto L31
	}
L30:
	;
	goto L6
L31:
	;
	v91 = v68 + int32(1)
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v59)+4))
	if v91 < v92 {
		v68 = v91
		goto L29
	} else {
		goto L40
	}
L32:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v77)+24))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v77)+16))
	if l1 != 0 {
		goto L35
	} else {
		goto L36
	}
L33:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v77)+16))
	if v79 == int32(0) {
		goto L31
	} else {
		goto L34
	}
L34:
	;
	goto L32
L35:
	;
	F_LockRelationOid(m, v83, v82)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L25
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	F_UnlockRelationOid(m, v83, v82)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L25
	} else {
		goto L39
	}
L38:
	;
	goto L31
L39:
	;
	goto L31
L40:
	;
	goto L30
L41:
	;
	goto L5
}
func F_ExecutorStart(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v31 int32
	_ = v31
	var v32 int64
	_ = v32
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v45 int64
	_ = v45
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v140 int64
	_ = v140
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v289 int32
	_ = v289
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v326 int32
	_ = v326
	var v328 int32
	_ = v328
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
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
	var v346 int32
	_ = v346
	var v348 int32
	_ = v348
	var v350 int32
	_ = v350
	var v352 int32
	_ = v352
	var v357 int32
	_ = v357
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v366 int32
	_ = v366
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v374 int32
	_ = v374
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v390 int32
	_ = v390
	var v394 int32
	_ = v394
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v402 int32
	_ = v402
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v407 int32
	_ = v407
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v425 int32
	_ = v425
	var v429 int32
	_ = v429
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v437 int32
	_ = v437
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v459 int32
	_ = v459
	var v467 int32
	_ = v467
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
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v514 int32
	_ = v514
	var v518 int32
	_ = v518
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v563 int32
	_ = v563
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v582 int32
	_ = v582
	var v584 int32
	_ = v584
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v599 int32
	_ = v599
	var v600 int32
	_ = v600
	var v601 int32
	_ = v601
	var v604 int32
	_ = v604
	var v605 int32
	_ = v605
	var v607 int32
	_ = v607
	var v612 int32
	_ = v612
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v628 int32
	_ = v628
	var v630 int32
	_ = v630
	var v631 int32
	_ = v631
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v636 int32
	_ = v636
	var v637 int32
	_ = v637
	var v638 int32
	_ = v638
	var v640 int32
	_ = v640
	var v643 int32
	_ = v643
	var v644 int32
	_ = v644
	var v646 int32
	_ = v646
	var v649 int32
	_ = v649
	var v659 int32
	_ = v659
	var v662 int32
	_ = v662
	var v681 int32
	_ = v681
	var v683 int32
	_ = v683
	var v692 int32
	_ = v692
	var v717 int32
	_ = v717
	var v719 int32
	_ = v719
	var v751 int32
	_ = v751
	var v752 int32
	_ = v752
	var v783 int32
	_ = v783
	var v784 int32
	_ = v784
	var v785 int32
	_ = v785
	var v787 int32
	_ = v787
	var v789 int32
	_ = v789
	var v790 int32
	_ = v790
	var v792 int32
	_ = v792
	var v794 int32
	_ = v794
	var v796 int32
	_ = v796
	var v798 int32
	_ = v798
	var v800 int32
	_ = v800
	var v802 int32
	_ = v802
	var v804 int32
	_ = v804
	var v806 int32
	_ = v806
	var v808 int32
	_ = v808
	var v810 int32
	_ = v810
	var v816 int32
	_ = v816
	var v847 int32
	_ = v847
	var v849 int32
	_ = v849
	var v850 int32
	_ = v850
	var v853 int32
	_ = v853
	var v854 int32
	_ = v854
	var v858 int32
	_ = v858
	var v861 int32
	_ = v861
	var v862 int32
	_ = v862
	var v864 int32
	_ = v864
	var v866 int32
	_ = v866
	var v870 int32
	_ = v870
	var v882 int32
	_ = v882
	var v905 int32
	_ = v905
	var v906 int32
	_ = v906
	var v938 int32
	_ = v938
	var v939 int32
	_ = v939
	var v940 int32
	_ = v940
	var v942 int32
	_ = v942
	var v946 int32
	_ = v946
	var v947 int32
	_ = v947
	var v954 int32
	_ = v954
	var v955 int32
	_ = v955
	var v957 int32
	_ = v957
	var v961 int32
	_ = v961
	var v962 int32
	_ = v962
	var v965 int32
	_ = v965
	var v967 int32
	_ = v967
	var v968 int32
	_ = v968
	var v969 int32
	_ = v969
	var v970 int32
	_ = v970
	var v972 int32
	_ = v972
	var v975 int32
	_ = v975
	var v976 int32
	_ = v976
	var v986 int32
	_ = v986
	var v987 int32
	_ = v987
	var v990 int32
	_ = v990
	var v994 int32
	_ = v994
	var v997 int32
	_ = v997
	var v999 int32
	_ = v999
	var v1002 int32
	_ = v1002
	var v1009 int32
	_ = v1009
	var v1011 int32
	_ = v1011
	var v1019 int32
	_ = v1019
	var v1020 int32
	_ = v1020
	var v1033 int32
	_ = v1033
	var v1036 int32
	_ = v1036
	var v1049 int32
	_ = v1049
	var v1066 int32
	_ = v1066
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
	var v1081 int32
	_ = v1081
	var v1083 int32
	_ = v1083
	var v1084 int32
	_ = v1084
	var v1087 int32
	_ = v1087
	var v1091 int32
	_ = v1091
	var v1094 int32
	_ = v1094
	var v1096 int32
	_ = v1096
	var v1099 int32
	_ = v1099
	var v1106 int32
	_ = v1106
	var v1108 int32
	_ = v1108
	var v1116 int32
	_ = v1116
	var v1117 int32
	_ = v1117
	var v1130 int32
	_ = v1130
	var v1146 int32
	_ = v1146
	var v1164 int32
	_ = v1164
	var v1165 int32
	_ = v1165
	var v1180 int32
	_ = v1180
	var v1198 int32
	_ = v1198
	var v1199 int32
	_ = v1199
	var v1214 int32
	_ = v1214
	var v1231 int32
	_ = v1231
	var v1232 int32
	_ = v1232
	var v1233 int32
	_ = v1233
	var v1235 int32
	_ = v1235
	var v1241 int32
	_ = v1241
	var v1242 int32
	_ = v1242
	var v1243 int32
	_ = v1243
	var v1246 int32
	_ = v1246
	var v1247 int32
	_ = v1247
	var v1248 int32
	_ = v1248
	var v1249 int32
	_ = v1249
	var v1250 int32
	_ = v1250
	var v1252 int32
	_ = v1252
	var v1253 int32
	_ = v1253
	var v1254 int32
	_ = v1254
	var v1257 int32
	_ = v1257
	var v1258 int32
	_ = v1258
	var v1261 int32
	_ = v1261
	var v1262 int32
	_ = v1262
	var v1266 int32
	_ = v1266
	var v1269 int32
	_ = v1269
	var v1270 int32
	_ = v1270
	var v1274 int32
	_ = v1274
	var v1284 int32
	_ = v1284
	var v1285 int32
	_ = v1285
	var v1286 int32
	_ = v1286
	var v1293 int32
	_ = v1293
	var v1296 int32
	_ = v1296
	var v1299 int32
	_ = v1299
	var v1300 int32
	_ = v1300
	var v1302 int32
	_ = v1302
	var v1305 int32
	_ = v1305
	var v1313 int32
	_ = v1313
	var v1339 int32
	_ = v1339
	var v1343 int32
	_ = v1343
	var v1344 int32
	_ = v1344
	var v1345 int32
	_ = v1345
	var v1346 int32
	_ = v1346
	var v1347 int32
	_ = v1347
	var v1353 int32
	_ = v1353
	var v1354 int32
	_ = v1354
	var v1357 int32
	_ = v1357
	var v1358 int32
	_ = v1358
	var v1359 int32
	_ = v1359
	var v1362 int32
	_ = v1362
	var v1363 int32
	_ = v1363
	var v1372 int32
	_ = v1372
	var v1373 int32
	_ = v1373
	var v1379 int32
	_ = v1379
	var v1384 int32
	_ = v1384
	var v1386 int32
	_ = v1386
	var v1387 int32
	_ = v1387
	var v1388 int32
	_ = v1388
	var v1391 int32
	_ = v1391
	var v1392 int32
	_ = v1392
	var v1398 int32
	_ = v1398
	var v1401 int32
	_ = v1401
	var v1402 int32
	_ = v1402
	var v1410 int32
	_ = v1410
	var v1415 int32
	_ = v1415
	var v1419 int32
	_ = v1419
	var v1422 int32
	_ = v1422
	var v1423 int32
	_ = v1423
	var v1431 int32
	_ = v1431
	var v1436 int32
	_ = v1436
	var v1440 int32
	_ = v1440
	var v1443 int32
	_ = v1443
	var v1444 int32
	_ = v1444
	var v1452 int32
	_ = v1452
	var v1457 int32
	_ = v1457
	var v1458 int32
	_ = v1458
	var v1464 int32
	_ = v1464
	var v1467 int32
	_ = v1467
	var v1468 int32
	_ = v1468
	var v1476 int32
	_ = v1476
	var v1481 int32
	_ = v1481
	var v1483 int32
	_ = v1483
	var v1484 int32
	_ = v1484
	var v1485 int32
	_ = v1485
	var v1489 int32
	_ = v1489
	var v1491 int32
	_ = v1491
	var v1492 int32
	_ = v1492
	var v1495 int32
	_ = v1495
	var v1497 int32
	_ = v1497
	var v1499 int32
	_ = v1499
	var v1501 int32
	_ = v1501
	var v1503 int32
	_ = v1503
	var v1505 int32
	_ = v1505
	var v1506 int32
	_ = v1506
	var v1515 int32
	_ = v1515
	var v1527 int32
	_ = v1527
	var v1528 int32
	_ = v1528
	var v1560 int32
	_ = v1560
	var v1565 int32
	_ = v1565
	var v1568 int32
	_ = v1568
	var v1572 int32
	_ = v1572
	var v1576 int32
	_ = v1576
	var v1579 int32
	_ = v1579
	var v1606 int32
	_ = v1606
	var v1610 int32
	_ = v1610
	var v1611 int32
	_ = v1611
	var v1612 int32
	_ = v1612
	var v1613 int32
	_ = v1613
	var v1614 int32
	_ = v1614
	var v1615 int32
	_ = v1615
	var v1616 int32
	_ = v1616
	var v1617 int32
	_ = v1617
	var v1618 int32
	_ = v1618
	var v1619 int32
	_ = v1619
	var v1621 int32
	_ = v1621
	var v1624 int32
	_ = v1624
	var v1625 int32
	_ = v1625
	var v1657 int32
	_ = v1657
	var v1658 int32
	_ = v1658
	var v1659 int32
	_ = v1659
	var v1662 int32
	_ = v1662
	var v1665 int32
	_ = v1665
	var v1668 int32
	_ = v1668
	var v1673 int32
	_ = v1673
	var v1676 int32
	_ = v1676
	var v1677 int32
	_ = v1677
	var v1685 int32
	_ = v1685
	var v1690 int32
	_ = v1690
	var v1694 int32
	_ = v1694
	var v1697 int32
	_ = v1697
	var v1698 int32
	_ = v1698
	var v1706 int32
	_ = v1706
	var v1711 int32
	_ = v1711
	var v1712 int32
	_ = v1712
	var v1745 int32
	_ = v1745
	var v1746 int32
	_ = v1746
	var v1750 int32
	_ = v1750
	var v1754 int32
	_ = v1754
	var v1755 int32
	_ = v1755
	var v1756 int32
	_ = v1756
	var v1757 int32
	_ = v1757
	var v1758 int32
	_ = v1758
	var v1759 int32
	_ = v1759
	var v1761 int32
	_ = v1761
	var v1772 int32
	_ = v1772
	v3 = int32(0)
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v32 = *(*int64)(unsafe.Add(mBase, uint32(v31)+8))
	v36 = *(*int32)(unsafe.Add(mBase, _consts[48]))
	if v36 == v3 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v71 = *(*int32)(unsafe.Add(mBase, _consts[340]))
	if v71 != 0 {
		goto L6
	} else {
		goto L7
	}
L2:
	;
	goto L1
L3:
	;
	v40 = int32(*(*uint8)(unsafe.Add(mBase, _consts[49])))
	if v40 != int32(1) {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v45 = *(*int64)(unsafe.Add(mBase, uint32(v36)+392))
	if int32(1)&base.B2i32(v45 != int64(0)) != 0 {
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v49 = int32(4483812)
	v51 = *(*int32)(unsafe.Add(mBase, _consts[13]))
	v52 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[13])) = v51 + v52
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
	*(*int32)(unsafe.Add(mBase, uint32(v36))) = v55 + v52
	*(*int64)(unsafe.Add(mBase, uint32(v36)+392)) = v32
	*(*int32)(unsafe.Add(mBase, uint32(v36))) = v55 + int32(2)
	v66 = *(*int32)(unsafe.Add(mBase, _consts[13]))
	*(*int32)(unsafe.Add(mBase, _consts[13])) = v66 - v52
	goto L2
L6:
	;
	m.T0[v71].(func(*base.Module, int32, int32))(m, l0, l1)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	v74 = m.G0
	v76 = v74 - int32(128)
	m.G0 = v76
	v79 = int32(*(*uint8)(unsafe.Add(mBase, _consts[341])))
	if v79 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L9:
	;
	return
L10:
	;
	return
L11:
	;
	v243 = F_CreateExecutorState(m)
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L9
	} else {
		goto L44
	}
L12:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v97)+52))
	if v98 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L13:
	;
	v86 = *(*int32)(unsafe.Add(mBase, _consts[39]))
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v86)+72))
	if v87 != 0 {
		goto L17
	} else {
		goto L18
	}
L14:
	;
	goto L15
L15:
	;
	if l1&int32(1) != 0 {
		goto L11
	} else {
		goto L22
	}
L16:
	;
	if l1&int32(1) != 0 {
		goto L11
	} else {
		goto L20
	}
L17:
	;
	v89 = int32(1)
	goto L19
L18:
	;
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86)+76)))
	v89 = v88
	goto L19
L19:
	;
	goto L16
L20:
	;
	if v89&int32(1) != 0 {
		goto L12
	} else {
		goto L21
	}
L21:
	;
	goto L11
L22:
	;
	goto L12
L23:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v97)+4))
	if v198 == int32(1) {
		goto L37
	} else {
		goto L38
	}
L24:
	;
	v101 = int32(0)
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v98)+4))
	if v102 <= v101 {
		goto L23
	} else {
		goto L25
	}
L25:
	;
	v107 = v101
	goto L26
L26:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v98)+12))
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v135+v107<<(uint(int32(2))%32))))
	v140 = *(*int64)(unsafe.Add(mBase, uint32(v139)+16))
	if v140&int64(-3) == int64(0) {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	goto L23
L28:
	;
	v165 = v107 + int32(1)
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v98)+4))
	if v165 < v166 {
		v107 = v165
		goto L26
	} else {
		goto L36
	}
L29:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v139)+4))
	v146 = F_get_rel_namespace(m, v145)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L9
	} else {
		goto L30
	}
L30:
	;
	v150 = *(*int32)(unsafe.Add(mBase, _consts[249]))
	goto L31
L31:
	;
	if base.B2i32(v150 != int32(0))&base.B2i32(v146 == v150) != 0 {
		goto L28
	} else {
		goto L32
	}
L32:
	;
	v155 = F_CreateCommandTag(m, v97)
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L9
	} else {
		goto L33
	}
L33:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v155<<(uint(int32(3))%32))+uint32(_consts[288])))
	goto L34
L34:
	;
	F_PreventCommandIfReadOnly(m, v161)
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L9
	} else {
		goto L35
	}
L35:
	;
	goto L28
L36:
	;
	goto L27
L37:
	;
	v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97)+25)))
	if v201 != int32(1) {
		goto L11
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	v204 = F_CreateCommandTag(m, v97)
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L9
	} else {
		goto L41
	}
L40:
	;
	goto L39
L41:
	;
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v204<<(uint(int32(3))%32))+uint32(_consts[288])))
	goto L42
L42:
	;
	F_PreventCommandIfParallelMode(m, v210)
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L9
	} else {
		goto L43
	}
L43:
	;
	goto L11
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v243
	v246 = int32(4489152)
	v247 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v243)+100))
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v249
	v251 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v243)+88)) = v251
	v253 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v253)+84))
	if v254 != 0 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v254)+4))
	v258 = F_palloc0(m, v255*int32(12))
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L9
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	v261 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v243)+56)) = v261
	v263 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v243)+96)) = v263
	v265 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if base.Ui32(int32(4)) <= base.Ui32(v265-int32(2)) {
		goto L52
	} else {
		goto L53
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v243)+92)) = v258
	goto L47
L49:
	;
	v308 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v309 = F_RegisterSnapshot(m, v308)
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L9
	} else {
		goto L66
	}
L50:
	;
	v296 = F_GetCurrentCommandId(m, int32(1))
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L9
	} else {
		goto L62
	}
L51:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L9
	} else {
		goto L59
	}
L52:
	;
	if v265 != int32(1) {
		goto L51
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	v278 = F_GetCurrentCommandId(m, int32(1))
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L9
	} else {
		goto L58
	}
L55:
	;
	v272 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v272)+72))
	if v273 != 0 {
		goto L50
	} else {
		goto L56
	}
L56:
	;
	v274 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v272)+25)))
	if v274 != 0 {
		goto L50
	} else {
		goto L57
	}
L57:
	;
	v307 = l1 | int32(32)
	goto L49
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v243)+64)) = v278
	v307 = l1
	goto L49
L59:
	;
	v285 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v76))) = v285
	F_errmsg_internal(m, int32(485074), v76)
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L9
	} else {
		goto L60
	}
L60:
	;
	F_errfinish(m, int32(494876), int32(238), int32(82725))
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L9
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
	*(*int32)(unsafe.Add(mBase, uint32(v243)+64)) = v296
	v301 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v302 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v301)+25)))
	if v302&int32(1) != 0 {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v305 = l1
	goto L65
L64:
	;
	v305 = l1 | int32(32)
	goto L65
L65:
	;
	v307 = v305
	goto L49
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v243)+8)) = v309
	v312 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v313 = F_RegisterSnapshot(m, v312)
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L9
	} else {
		goto L67
	}
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v243)+128)) = v307
	*(*int32)(unsafe.Add(mBase, uint32(v243)+12)) = v313
	v317 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v243)+132)) = v317
	v319 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v319)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v243)+176)) = v320
	if v307&int32(33) == int32(0) {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v326 = int32(4386248)
	v328 = *(*int32)(unsafe.Add(mBase, _consts[285]))
	*(*int32)(unsafe.Add(mBase, _consts[285])) = v328 + int32(1)
	goto L71
L69:
	;
	goto L70
L70:
	;
	v332 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v333 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v333)+36))
	v335 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v336 = *(*int32)(unsafe.Add(mBase, uint32(v333)+44))
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v333)+52))
	v339 = F_ExecCheckPermissions(m, v336, v337, int32(1))
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L9
	} else {
		goto L72
	}
L71:
	;
	goto L70
L72:
	;
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v333)+52))
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v333)+48))
	v343 = F_bms_copy(m, v342)
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L9
	} else {
		goto L73
	}
L73:
	;
	F_ExecInitRangeTable(m, v335, v336, v341, v343)
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L9
	} else {
		goto L74
	}
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v335)+36)) = v333
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v333)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v335)+40)) = v348
	v350 = m.G0
	v352 = v350 - int32(16)
	m.G0 = v352
	if v348 == int32(0) {
		v1261 = v307
		v1262 = v335
		v1266 = l0
		v1269 = v76
		v1270 = v333
		v1274 = v352
		v1284 = v334
		v1285 = v247
		v1286 = v332
		goto L75
	} else {
		goto L76
	}
L75:
	;
	m.G0 = v1274 + int32(16)
	v1293 = *(*int32)(unsafe.Add(mBase, uint32(v1270)+72))
	if v1293 == int32(0) {
		goto L218
	} else {
		goto L219
	}
L76:
	;
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v348)+4))
	if v357 <= int32(0) {
		v1261 = v307
		v1262 = v335
		v1266 = l0
		v1269 = v76
		v1270 = v333
		v1274 = v352
		v1284 = v334
		v1285 = v247
		v1286 = v332
		goto L75
	} else {
		goto L77
	}
L77:
	;
	v361 = v307
	v362 = v335
	v366 = l0
	v369 = v76
	v370 = v333
	v374 = v352
	v378 = v3
	v379 = v348
	v384 = v334
	v385 = v247
	v386 = v332
	goto L78
L78:
	;
	v390 = *(*int32)(unsafe.Add(mBase, uint32(v379)+12))
	v394 = *(*int32)(unsafe.Add(mBase, uint32(v390+v378<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v374)+12)) = int32(0)
	v397 = F_CreateExprContext(m, v362)
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L9
	} else {
		goto L80
	}
L79:
	;
	v1261 = v361
	v1262 = v362
	v1266 = v366
	v1269 = v369
	v1270 = v370
	v1274 = v374
	v1284 = v384
	v1285 = v385
	v1286 = v386
	goto L75
L80:
	;
	v399 = *(*int32)(unsafe.Add(mBase, uint32(v362)+76))
	if v399 == int32(0) {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	v402 = *(*int32)(unsafe.Add(mBase, uint32(v362)+100))
	v404 = F_CreatePartitionDirectory(m, v402, int32(0))
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
		goto L9
	} else {
		goto L84
	}
L82:
	;
	goto L83
L83:
	;
	v407 = int32(0)
	v409 = *(*int32)(unsafe.Add(mBase, uint32(v394)+8))
	if v409 != 0 {
		goto L85
	} else {
		goto L86
	}
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v362)+76)) = v404
	goto L83
L85:
	;
	v410 = *(*int32)(unsafe.Add(mBase, uint32(v409)+4))
	v411 = v410
	goto L87
L86:
	;
	v411 = v407
	goto L87
L87:
	;
	v416 = F_palloc(m, v411<<(uint(int32(2))%32)+int32(24))
	mBase = m.M
	v417 = m.ExcPending
	if v417 != 0 {
		goto L9
	} else {
		goto L88
	}
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v416)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v416))) = v397
	v421 = *(*int32)(unsafe.Add(mBase, uint32(v394)+12))
	v422 = F_bms_copy(m, v421)
	mBase = m.M
	v423 = m.ExcPending
	if v423 != 0 {
		goto L9
	} else {
		goto L89
	}
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v416)+20)) = v411
	v425 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v416)+16)) = uint16(v425)
	*(*int32)(unsafe.Add(mBase, uint32(v416)+8)) = v422
	v429 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v434 = F_AllocSetContextCreateInternal(m, v429, int32(370968), v425, int32(8192), int32(8388608))
	mBase = m.M
	v435 = m.ExcPending
	if v435 != 0 {
		goto L9
	} else {
		goto L90
	}
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v416)+12)) = v434
	v437 = *(*int32)(unsafe.Add(mBase, uint32(v394)+8))
	if v437 == int32(0) {
		v1214 = v407
		goto L91
	} else {
		goto L92
	}
L91:
	;
	v1231 = *(*int32)(unsafe.Add(mBase, uint32(v362)+44))
	v1232 = F_lappend(m, v1231, v416)
	mBase = m.M
	v1233 = m.ExcPending
	if v1233 != 0 {
		goto L9
	} else {
		goto L205
	}
L92:
	;
	v440 = int32(0)
	v441 = *(*int32)(unsafe.Add(mBase, uint32(v437)+4))
	if v441 <= v440 {
		v1214 = v407
		goto L91
	} else {
		goto L93
	}
L93:
	;
	v459 = v407
	v467 = v440
	goto L94
L94:
	;
	v478 = v467 << (uint(int32(2)) % 32)
	v479 = *(*int32)(unsafe.Add(mBase, uint32(v437)+12))
	v481 = *(*int32)(unsafe.Add(mBase, uint32(v478+v479)))
	if v481 != 0 {
		goto L96
	} else {
		goto L97
	}
L95:
	;
	v1214 = v1180
	goto L91
L96:
	;
	v482 = *(*int32)(unsafe.Add(mBase, uint32(v481)+4))
	v483 = v482
	goto L98
L97:
	;
	v483 = int32(0)
	goto L98
L98:
	;
	v489 = F_palloc(m, v483*int32(120)|int32(4))
	mBase = m.M
	v490 = m.ExcPending
	if v490 != 0 {
		goto L9
	} else {
		goto L99
	}
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v478+(v416+int32(24))))) = v489
	*(*int32)(unsafe.Add(mBase, uint32(v489))) = v483
	if v481 == int32(0) {
		v1180 = v459
		goto L100
	} else {
		goto L101
	}
L100:
	;
	v1198 = v467 + int32(1)
	v1199 = *(*int32)(unsafe.Add(mBase, uint32(v437)+4))
	if v1198 < v1199 {
		v459 = v1180
		v467 = v1198
		goto L94
	} else {
		goto L204
	}
L101:
	;
	v495 = int32(0)
	v496 = *(*int32)(unsafe.Add(mBase, uint32(v481)+4))
	if v496 <= v495 {
		v1180 = v459
		goto L100
	} else {
		goto L102
	}
L102:
	;
	v514 = v459
	v518 = v495
	goto L103
L103:
	;
	v533 = v489 + int32(4) + v518*int32(120)
	v534 = *(*int32)(unsafe.Add(mBase, uint32(v481)+12))
	v538 = *(*int32)(unsafe.Add(mBase, uint32(v534+v518<<(uint(int32(2))%32))))
	v539 = *(*int32)(unsafe.Add(mBase, uint32(v538)+4))
	v540 = F_ExecGetRangeTableRelation(m, v362, v539)
	mBase = m.M
	v541 = m.ExcPending
	if v541 != 0 {
		goto L9
	} else {
		goto L105
	}
L104:
	;
	v1180 = v1146
	goto L100
L105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v533))) = v540
	v543 = F_RelationGetPartitionKey(m, v540)
	mBase = m.M
	v544 = m.ExcPending
	if v544 != 0 {
		goto L9
	} else {
		goto L106
	}
L106:
	;
	v545 = *(*int32)(unsafe.Add(mBase, uint32(v362)+76))
	v546 = F_PartitionDirectoryLookup(m, v545, v540)
	mBase = m.M
	v547 = m.ExcPending
	if v547 != 0 {
		goto L9
	} else {
		goto L107
	}
L107:
	;
	v548 = *(*int32)(unsafe.Add(mBase, uint32(v546)))
	*(*int32)(unsafe.Add(mBase, uint32(v533)+4)) = v548
	v552 = F_palloc(m, v548<<(uint(int32(2))%32))
	mBase = m.M
	v553 = m.ExcPending
	if v553 != 0 {
		goto L9
	} else {
		goto L108
	}
L108:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v533)+8)) = v552
	v555 = *(*int32)(unsafe.Add(mBase, uint32(v546)))
	v556 = *(*int32)(unsafe.Add(mBase, uint32(v538)+12))
	if v555 != v556 {
		goto L111
	} else {
		goto L112
	}
L109:
	;
	v938 = *(*int32)(unsafe.Add(mBase, uint32(v538)+8))
	v939 = F_bms_copy(m, v938)
	mBase = m.M
	v940 = m.ExcPending
	if v940 != 0 {
		goto L9
	} else {
		goto L161
	}
L110:
	;
	v637 = F_palloc(m, v636)
	mBase = m.M
	v638 = m.ExcPending
	if v638 != 0 {
		goto L9
	} else {
		goto L137
	}
L111:
	;
	v636 = v555 << (uint(int32(2)) % 32)
	goto L110
L112:
	;
	goto L113
L113:
	;
	v560 = *(*int32)(unsafe.Add(mBase, uint32(v546)+8))
	v561 = *(*int32)(unsafe.Add(mBase, uint32(v538)+28))
	v563 = v555 << (uint(int32(2)) % 32)
	if base.Ui32(int32(4)) <= base.Ui32(v563) {
		goto L117
	} else {
		goto L118
	}
L114:
	;
	if v625 != 0 {
		v636 = v563
		goto L110
	} else {
		goto L132
	}
L115:
	;
	v625 = int32(0)
	goto L114
L116:
	;
	v599 = v594
	v600 = v595
	v601 = v596
	goto L126
L117:
	;
	if (v560|v561)&int32(3) != 0 {
		v594 = v560
		v595 = v561
		v596 = v563
		goto L116
	} else {
		goto L120
	}
L118:
	;
	v587 = v560
	v588 = v561
	v589 = v563
	goto L119
L119:
	;
	if v589 == int32(0) {
		goto L115
	} else {
		goto L125
	}
L120:
	;
	v571 = v560
	v572 = v561
	v573 = v563
	goto L121
L121:
	;
	v576 = *(*int32)(unsafe.Add(mBase, uint32(v571)))
	v577 = *(*int32)(unsafe.Add(mBase, uint32(v572)))
	if v576 != v577 {
		v594 = v571
		v595 = v572
		v596 = v573
		goto L116
	} else {
		goto L123
	}
L122:
	;
	v587 = v582
	v588 = v580
	v589 = v584
	goto L119
L123:
	;
	v579 = int32(4)
	v580 = v572 + v579
	v582 = v571 + v579
	v584 = v573 - v579
	if base.Ui32(int32(3)) < base.Ui32(v584) {
		v571 = v582
		v572 = v580
		v573 = v584
		goto L121
	} else {
		goto L124
	}
L124:
	;
	goto L122
L125:
	;
	v594 = v587
	v595 = v588
	v596 = v589
	goto L116
L126:
	;
	v604 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v599))))
	v605 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v600))))
	if v604 == v605 {
		goto L128
	} else {
		goto L129
	}
L127:
	;
	v625 = v604 - v605
	goto L114
L128:
	;
	v607 = int32(1)
	v612 = v601 - v607
	if v612 != 0 {
		v599 = v599 + v607
		v600 = v600 + v607
		v601 = v612
		goto L126
	} else {
		goto L131
	}
L129:
	;
	goto L130
L130:
	;
	goto L127
L131:
	;
	goto L115
L132:
	;
	v626 = *(*int32)(unsafe.Add(mBase, uint32(v538)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v533)+12)) = v626
	v628 = *(*int32)(unsafe.Add(mBase, uint32(v538)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v533)+16)) = v628
	v630 = *(*int32)(unsafe.Add(mBase, uint32(v538)+16))
	v631 = *(*int32)(unsafe.Add(mBase, uint32(v538)+12))
	v633 = v631 << (uint(int32(2)) % 32)
	if v633 != 0 {
		goto L134
	} else {
		goto L135
	}
L133:
	;
	goto L109
L134:
	;
	v634 = F__emscripten_memcpy_bulkmem(m, v552, v630, v633)
	mBase = m.M
	goto L136
L135:
	;
	goto L136
L136:
	;
	goto L133
L137:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v533)+12)) = v637
	v640 = *(*int32)(unsafe.Add(mBase, uint32(v546)))
	v643 = F_palloc(m, v640<<(uint(int32(2))%32))
	mBase = m.M
	v644 = m.ExcPending
	if v644 != 0 {
		goto L9
	} else {
		goto L138
	}
L138:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v533)+16)) = v643
	v646 = *(*int32)(unsafe.Add(mBase, uint32(v546)))
	if v646 <= int32(0) {
		goto L109
	} else {
		goto L139
	}
L139:
	;
	v649 = int32(0)
	v659 = v649
	v662 = v649
	goto L140
L140:
	;
	v681 = *(*int32)(unsafe.Add(mBase, uint32(v538)+12))
	if v681 <= v659 {
		v751 = v659
		goto L142
	} else {
		goto L143
	}
L141:
	;
	goto L109
L142:
	;
	v752 = v751
	goto L149
L143:
	;
	v683 = *(*int32)(unsafe.Add(mBase, uint32(v538)+28))
	v692 = v659
	goto L144
L144:
	;
	v717 = *(*int32)(unsafe.Add(mBase, uint32(v683+v692<<(uint(int32(2))%32))))
	if v717 != 0 {
		v751 = v692
		goto L142
	} else {
		goto L146
	}
L145:
	;
	v751 = v681
	goto L142
L146:
	;
	v719 = v692 + int32(1)
	if v719 != v681 {
		v692 = v719
		goto L144
	} else {
		goto L147
	}
L147:
	;
	goto L145
L148:
	;
	v905 = v662 + int32(1)
	v906 = *(*int32)(unsafe.Add(mBase, uint32(v546)))
	if v905 < v906 {
		v659 = v882
		v662 = v905
		goto L140
	} else {
		goto L160
	}
L149:
	;
	if v681 <= v752 {
		goto L151
	} else {
		goto L152
	}
L150:
	;
	v861 = v662 << (uint(int32(2)) % 32)
	v862 = *(*int32)(unsafe.Add(mBase, uint32(v533)+12))
	v864 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v861+v862))) = v864
	v866 = *(*int32)(unsafe.Add(mBase, uint32(v533)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v866+v861))) = v864
	v870 = *(*int32)(unsafe.Add(mBase, uint32(v533)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v870+v861))) = int32(0)
	v882 = v752
	goto L148
L151:
	;
	v816 = v752
	goto L154
L152:
	;
	v783 = int32(2)
	v784 = v752 << (uint(v783) % 32)
	v785 = *(*int32)(unsafe.Add(mBase, uint32(v538)+28))
	v787 = *(*int32)(unsafe.Add(mBase, uint32(v784+v785)))
	v789 = v662 << (uint(v783) % 32)
	v790 = *(*int32)(unsafe.Add(mBase, uint32(v546)+8))
	v792 = *(*int32)(unsafe.Add(mBase, uint32(v789+v790)))
	if v787 != v792 {
		goto L151
	} else {
		goto L153
	}
L153:
	;
	v794 = *(*int32)(unsafe.Add(mBase, uint32(v533)+8))
	v796 = *(*int32)(unsafe.Add(mBase, uint32(v538)+16))
	v798 = *(*int32)(unsafe.Add(mBase, uint32(v796+v784)))
	*(*int32)(unsafe.Add(mBase, uint32(v794+v789))) = v798
	v800 = *(*int32)(unsafe.Add(mBase, uint32(v533)+12))
	v802 = *(*int32)(unsafe.Add(mBase, uint32(v538)+20))
	v804 = *(*int32)(unsafe.Add(mBase, uint32(v802+v784)))
	*(*int32)(unsafe.Add(mBase, uint32(v800+v789))) = v804
	v806 = *(*int32)(unsafe.Add(mBase, uint32(v533)+16))
	v808 = *(*int32)(unsafe.Add(mBase, uint32(v538)+24))
	v810 = *(*int32)(unsafe.Add(mBase, uint32(v808+v784)))
	*(*int32)(unsafe.Add(mBase, uint32(v806+v789))) = v810
	v882 = v752 + int32(1)
	goto L148
L154:
	;
	v847 = v816 + int32(1)
	if v847 < v681 {
		goto L156
	} else {
		goto L157
	}
L155:
	;
	goto L150
L156:
	;
	v849 = *(*int32)(unsafe.Add(mBase, uint32(v538)+28))
	v850 = int32(2)
	v853 = *(*int32)(unsafe.Add(mBase, uint32(v849+v847<<(uint(v850)%32))))
	v854 = *(*int32)(unsafe.Add(mBase, uint32(v546)+8))
	v858 = *(*int32)(unsafe.Add(mBase, uint32(v854+v662<<(uint(v850)%32))))
	if v853 != v858 {
		v816 = v847
		goto L154
	} else {
		goto L159
	}
L157:
	;
	goto L158
L158:
	;
	goto L155
L159:
	;
	v752 = v847
	goto L149
L160:
	;
	goto L141
L161:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v533)+20)) = v939
	v942 = *(*int32)(unsafe.Add(mBase, uint32(v538)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v533)+24)) = v942
	if v942 == int32(0) {
		goto L162
	} else {
		goto L163
	}
L162:
	;
	v957 = *(*int32)(unsafe.Add(mBase, uint32(v538)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v533)+28)) = v957
	if v957 == int32(0) {
		goto L166
	} else {
		goto L167
	}
L163:
	;
	v946 = *(*int32)(unsafe.Add(mBase, uint32(v397)+64))
	v947 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v946)+128)))
	if v947&int32(2) != 0 {
		goto L162
	} else {
		goto L164
	}
L164:
	;
	F_InitPartitionPruneContext(m, v533+int32(32), v942, v546, v543, int32(0), v397)
	mBase = m.M
	v954 = m.ExcPending
	if v954 != 0 {
		goto L9
	} else {
		goto L165
	}
L165:
	;
	v955 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v416)+16)) = uint8(v955)
	goto L162
L166:
	;
	v967 = *(*int32)(unsafe.Add(mBase, uint32(v416)+4))
	v968 = *(*int32)(unsafe.Add(mBase, uint32(v538)+40))
	v969 = F_bms_add_members(m, v967, v968)
	mBase = m.M
	v970 = m.ExcPending
	if v970 != 0 {
		goto L9
	} else {
		goto L169
	}
L167:
	;
	v961 = *(*int32)(unsafe.Add(mBase, uint32(v397)+64))
	v962 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v961)+128)))
	if v962&int32(2) != 0 {
		goto L166
	} else {
		goto L168
	}
L168:
	;
	v965 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v416)+17)) = uint8(v965)
	goto L166
L169:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v416)+4)) = v969
	v972 = *(*int32)(unsafe.Add(mBase, uint32(v538)+32))
	if v972 == int32(0) {
		v1146 = v514
		goto L170
	} else {
		goto L171
	}
L170:
	;
	v1164 = v518 + int32(1)
	v1165 = *(*int32)(unsafe.Add(mBase, uint32(v481)+4))
	if v1164 < v1165 {
		v514 = v1146
		v518 = v1164
		goto L103
	} else {
		goto L203
	}
L171:
	;
	v975 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v416)+16)))
	if v975 != 0 {
		v1146 = v514
		goto L170
	} else {
		goto L172
	}
L172:
	;
	v976 = *(*int32)(unsafe.Add(mBase, uint32(v533)+20))
	if v976 == int32(0) {
		goto L175
	} else {
		goto L176
	}
L173:
	;
	if v1033 < int32(0) {
		v1146 = v514
		goto L170
	} else {
		goto L184
	}
L174:
	;
	v1033 = base.I32_ctz(v1019) | v1020<<(uint(int32(5))%32)
	goto L173
L175:
	;
	v1033 = int32(-2)
	goto L173
L176:
	;
	v986 = base.I32_div_s(int32(0), int32(32))
	v987 = *(*int32)(unsafe.Add(mBase, uint32(v976)+4))
	if v987 <= v986 {
		goto L175
	} else {
		goto L177
	}
L177:
	;
	v990 = v976 + int32(8)
	v994 = *(*int32)(unsafe.Add(mBase, uint32(v990+v986<<(uint(int32(2))%32))))
	v997 = v994 & int32(-1)
	if v997 != 0 {
		v1019 = v997
		v1020 = v986
		goto L174
	} else {
		goto L178
	}
L178:
	;
	v999 = v986 + int32(1)
	if v999 == v987 {
		goto L175
	} else {
		goto L179
	}
L179:
	;
	v1002 = v999
	goto L180
L180:
	;
	v1009 = *(*int32)(unsafe.Add(mBase, uint32(v990+v1002<<(uint(int32(2))%32))))
	if v1009 != 0 {
		v1019 = v1009
		v1020 = v1002
		goto L174
	} else {
		goto L182
	}
L181:
	;
	goto L175
L182:
	;
	v1011 = v1002 + int32(1)
	if v1011 != v987 {
		v1002 = v1011
		goto L180
	} else {
		goto L183
	}
L183:
	;
	goto L181
L184:
	;
	v1036 = v1033
	v1049 = v514
	goto L185
L185:
	;
	v1066 = *(*int32)(unsafe.Add(mBase, uint32(v533)+16))
	v1070 = *(*int32)(unsafe.Add(mBase, uint32(v1066+v1036<<(uint(int32(2))%32))))
	if v1070 != 0 {
		goto L187
	} else {
		goto L188
	}
L186:
	;
	v1146 = v1073
	goto L170
L187:
	;
	v1071 = F_bms_add_member(m, v1049, v1070)
	mBase = m.M
	v1072 = m.ExcPending
	if v1072 != 0 {
		goto L9
	} else {
		goto L190
	}
L188:
	;
	v1073 = v1049
	goto L189
L189:
	;
	v1074 = *(*int32)(unsafe.Add(mBase, uint32(v533)+20))
	if v1074 == int32(0) {
		goto L193
	} else {
		goto L194
	}
L190:
	;
	v1073 = v1071
	goto L189
L191:
	;
	if int32(0) <= v1130 {
		v1036 = v1130
		v1049 = v1073
		goto L185
	} else {
		goto L202
	}
L192:
	;
	v1130 = base.I32_ctz(v1116) | v1117<<(uint(int32(5))%32)
	goto L191
L193:
	;
	v1130 = int32(-2)
	goto L191
L194:
	;
	v1081 = v1036 + int32(1)
	v1083 = base.I32_div_s(v1081, int32(32))
	v1084 = *(*int32)(unsafe.Add(mBase, uint32(v1074)+4))
	if v1084 <= v1083 {
		goto L193
	} else {
		goto L195
	}
L195:
	;
	v1087 = v1074 + int32(8)
	v1091 = *(*int32)(unsafe.Add(mBase, uint32(v1087+v1083<<(uint(int32(2))%32))))
	v1094 = v1091 & (int32(-1) << (uint(v1081) % 32))
	if v1094 != 0 {
		v1116 = v1094
		v1117 = v1083
		goto L192
	} else {
		goto L196
	}
L196:
	;
	v1096 = v1083 + int32(1)
	if v1096 == v1084 {
		goto L193
	} else {
		goto L197
	}
L197:
	;
	v1099 = v1096
	goto L198
L198:
	;
	v1106 = *(*int32)(unsafe.Add(mBase, uint32(v1087+v1099<<(uint(int32(2))%32))))
	if v1106 != 0 {
		v1116 = v1106
		v1117 = v1099
		goto L192
	} else {
		goto L200
	}
L199:
	;
	goto L193
L200:
	;
	v1108 = v1099 + int32(1)
	if v1108 != v1084 {
		v1099 = v1108
		goto L198
	} else {
		goto L201
	}
L201:
	;
	goto L199
L202:
	;
	goto L186
L203:
	;
	goto L104
L204:
	;
	goto L95
L205:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v362)+44)) = v1232
	v1235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v416)+16)))
	if v1235 == int32(1) {
		goto L207
	} else {
		goto L208
	}
L206:
	;
	v1248 = *(*int32)(unsafe.Add(mBase, uint32(v362)+52))
	v1249 = F_bms_add_members(m, v1248, v1247)
	mBase = m.M
	v1250 = m.ExcPending
	if v1250 != 0 {
		goto L9
	} else {
		goto L211
	}
L207:
	;
	v1241 = F_ExecFindMatchingSubPlans(m, v416, int32(1), v374+int32(12))
	mBase = m.M
	v1242 = m.ExcPending
	if v1242 != 0 {
		goto L9
	} else {
		goto L210
	}
L208:
	;
	goto L209
L209:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v374)+12)) = v1214
	v1246 = int32(0)
	v1247 = v1214
	goto L206
L210:
	;
	v1243 = *(*int32)(unsafe.Add(mBase, uint32(v374)+12))
	v1246 = v1241
	v1247 = v1243
	goto L206
L211:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v362)+52)) = v1249
	v1252 = *(*int32)(unsafe.Add(mBase, uint32(v362)+48))
	v1253 = F_lappend(m, v1252, v1246)
	mBase = m.M
	v1254 = m.ExcPending
	if v1254 != 0 {
		goto L9
	} else {
		goto L212
	}
L212:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v362)+48)) = v1253
	v1257 = v378 + int32(1)
	v1258 = *(*int32)(unsafe.Add(mBase, uint32(v379)+4))
	if v1257 < v1258 {
		v378 = v1257
		goto L78
	} else {
		goto L213
	}
L213:
	;
	goto L79
L214:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1266)+44)) = v1657
	*(*int32)(unsafe.Add(mBase, uint32(v1266)+36)) = v1772
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v1285
	m.G0 = v1269 + int32(128)
	return
L215:
	;
	v1712 = int32(0)
	goto L293
L216:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1694 = m.ExcPending
	if v1694 != 0 {
		goto L9
	} else {
		goto L289
	}
L217:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1673 = m.ExcPending
	if v1673 != 0 {
		goto L9
	} else {
		goto L285
	}
L218:
	;
	v1560 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1262)+156)) = v1560
	*(*int32)(unsafe.Add(mBase, uint32(v1262)+104)) = v1560
	v1565 = *(*int32)(unsafe.Add(mBase, uint32(v1270)+64))
	if v1565 == v1560 {
		goto L269
	} else {
		goto L270
	}
L219:
	;
	v1296 = *(*int32)(unsafe.Add(mBase, uint32(v1262)+20))
	v1299 = F_palloc0(m, v1296<<(uint(int32(2))%32))
	mBase = m.M
	v1300 = m.ExcPending
	if v1300 != 0 {
		goto L9
	} else {
		goto L220
	}
L220:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1262)+28)) = v1299
	v1302 = *(*int32)(unsafe.Add(mBase, uint32(v1270)+72))
	if v1302 == int32(0) {
		goto L218
	} else {
		goto L221
	}
L221:
	;
	v1305 = *(*int32)(unsafe.Add(mBase, uint32(v1302)+4))
	if v1305 <= int32(0) {
		goto L218
	} else {
		goto L222
	}
L222:
	;
	v1313 = int32(0)
	goto L223
L223:
	;
	v1339 = *(*int32)(unsafe.Add(mBase, uint32(v1302)+12))
	v1343 = *(*int32)(unsafe.Add(mBase, uint32(v1339+v1313<<(uint(int32(2))%32))))
	v1344 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1343)+32)))
	if v1344 != 0 {
		goto L225
	} else {
		goto L226
	}
L224:
	;
	goto L218
L225:
	;
	v1527 = v1313 + int32(1)
	v1528 = *(*int32)(unsafe.Add(mBase, uint32(v1302)+4))
	if v1527 < v1528 {
		v1313 = v1527
		goto L223
	} else {
		goto L268
	}
L226:
	;
	v1345 = *(*int32)(unsafe.Add(mBase, uint32(v1262)+16))
	v1346 = *(*int32)(unsafe.Add(mBase, uint32(v1345)+12))
	v1347 = *(*int32)(unsafe.Add(mBase, uint32(v1343)+4))
	v1353 = *(*int32)(unsafe.Add(mBase, uint32(v1346+v1347<<(uint(int32(2))%32)-int32(4))))
	v1354 = *(*int32)(unsafe.Add(mBase, uint32(v1353)+12))
	if v1354 == int32(0) {
		goto L227
	} else {
		goto L228
	}
L227:
	;
	v1357 = *(*int32)(unsafe.Add(mBase, uint32(v1262)+52))
	v1358 = F_bms_is_member(m, v1347, v1357)
	mBase = m.M
	v1359 = m.ExcPending
	if v1359 != 0 {
		goto L9
	} else {
		goto L230
	}
L228:
	;
	goto L229
L229:
	;
	v1362 = *(*int32)(unsafe.Add(mBase, uint32(v1353)+16))
	v1363 = *(*int32)(unsafe.Add(mBase, uint32(v1343)+16))
	if base.Ui32(int32(5)) <= base.Ui32(v1363) {
		goto L233
	} else {
		goto L234
	}
L230:
	;
	if v1358 == int32(0) {
		goto L225
	} else {
		goto L231
	}
L231:
	;
	goto L229
L232:
	;
	v1491 = F_palloc(m, int32(44))
	mBase = m.M
	v1492 = m.ExcPending
	if v1492 != 0 {
		goto L9
	} else {
		goto L267
	}
L233:
	;
	if v1363 == int32(5) {
		v1489 = int32(0)
		goto L232
	} else {
		goto L236
	}
L234:
	;
	goto L235
L235:
	;
	v1386 = *(*int32)(unsafe.Add(mBase, uint32(v1343)+4))
	v1387 = F_ExecGetRangeTableRelation(m, v1262, v1386)
	mBase = m.M
	v1388 = m.ExcPending
	if v1388 != 0 {
		goto L9
	} else {
		goto L240
	}
L236:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1372 = m.ExcPending
	if v1372 != 0 {
		goto L9
	} else {
		goto L237
	}
L237:
	;
	v1373 = *(*int32)(unsafe.Add(mBase, uint32(v1343)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v1269)+16)) = v1373
	F_errmsg_internal(m, int32(484536), v1269+int32(16))
	mBase = m.M
	v1379 = m.ExcPending
	if v1379 != 0 {
		goto L9
	} else {
		goto L238
	}
L238:
	;
	F_errfinish(m, int32(494876), int32(918), int32(282653))
	mBase = m.M
	v1384 = m.ExcPending
	if v1384 != 0 {
		goto L9
	} else {
		goto L239
	}
L239:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L240:
	;
	if v1387 == int32(0) {
		v1489 = int32(0)
		goto L232
	} else {
		goto L241
	}
L241:
	;
	v1391 = *(*int32)(unsafe.Add(mBase, uint32(v1387)+48))
	v1392 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1391)+119)))
	switch v1392 - int32(83) {
	case 0:
		goto L247
	default:
		goto L216
	case 19:
		goto L243
	case 26:
		goto L244
	case 29, 31:
		v1489 = v1387
		goto L232
	case 33:
		goto L246
	case 35:
		goto L245
	}
L242:
	;
	v1489 = v1387
	goto L232
L243:
	;
	v1483 = F_GetFdwRoutineForRelation(m, v1387, int32(0))
	mBase = m.M
	v1484 = m.ExcPending
	if v1484 != 0 {
		goto L9
	} else {
		goto L265
	}
L244:
	;
	v1458 = *(*int32)(unsafe.Add(mBase, uint32(v1343)+16))
	if v1458 == int32(4) {
		goto L242
	} else {
		goto L260
	}
L245:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1440 = m.ExcPending
	if v1440 != 0 {
		goto L9
	} else {
		goto L256
	}
L246:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1419 = m.ExcPending
	if v1419 != 0 {
		goto L9
	} else {
		goto L252
	}
L247:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1398 = m.ExcPending
	if v1398 != 0 {
		goto L9
	} else {
		goto L248
	}
L248:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v1401 = m.ExcPending
	if v1401 != 0 {
		goto L9
	} else {
		goto L249
	}
L249:
	;
	v1402 = *(*int32)(unsafe.Add(mBase, uint32(v1387)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v1269)+48)) = v1402 + int32(4)
	F_errmsg(m, int32(704617), v1269+int32(48))
	mBase = m.M
	v1410 = m.ExcPending
	if v1410 != 0 {
		goto L9
	} else {
		goto L250
	}
L250:
	;
	F_errfinish(m, int32(494876), int32(1197), int32(307494))
	mBase = m.M
	v1415 = m.ExcPending
	if v1415 != 0 {
		goto L9
	} else {
		goto L251
	}
L251:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L252:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v1422 = m.ExcPending
	if v1422 != 0 {
		goto L9
	} else {
		goto L253
	}
L253:
	;
	v1423 = *(*int32)(unsafe.Add(mBase, uint32(v1387)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v1269)+64)) = v1423 + int32(4)
	F_errmsg(m, int32(691966), v1269-int32(-64))
	mBase = m.M
	v1431 = m.ExcPending
	if v1431 != 0 {
		goto L9
	} else {
		goto L254
	}
L254:
	;
	F_errfinish(m, int32(494876), int32(1204), int32(307494))
	mBase = m.M
	v1436 = m.ExcPending
	if v1436 != 0 {
		goto L9
	} else {
		goto L255
	}
L255:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L256:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v1443 = m.ExcPending
	if v1443 != 0 {
		goto L9
	} else {
		goto L257
	}
L257:
	;
	v1444 = *(*int32)(unsafe.Add(mBase, uint32(v1387)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v1269)+80)) = v1444 + int32(4)
	F_errmsg(m, int32(680079), v1269+int32(80))
	mBase = m.M
	v1452 = m.ExcPending
	if v1452 != 0 {
		goto L9
	} else {
		goto L258
	}
L258:
	;
	F_errfinish(m, int32(494876), int32(1211), int32(307494))
	mBase = m.M
	v1457 = m.ExcPending
	if v1457 != 0 {
		goto L9
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
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1464 = m.ExcPending
	if v1464 != 0 {
		goto L9
	} else {
		goto L261
	}
L261:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v1467 = m.ExcPending
	if v1467 != 0 {
		goto L9
	} else {
		goto L262
	}
L262:
	;
	v1468 = *(*int32)(unsafe.Add(mBase, uint32(v1387)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v1269)+96)) = v1468 + int32(4)
	F_errmsg(m, int32(680416), v1269+int32(96))
	mBase = m.M
	v1476 = m.ExcPending
	if v1476 != 0 {
		goto L9
	} else {
		goto L263
	}
L263:
	;
	F_errfinish(m, int32(494876), int32(1219), int32(307494))
	mBase = m.M
	v1481 = m.ExcPending
	if v1481 != 0 {
		goto L9
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
	v1485 = *(*int32)(unsafe.Add(mBase, uint32(v1483)+108))
	if v1485 == int32(0) {
		goto L217
	} else {
		goto L266
	}
L266:
	;
	goto L242
L267:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1491)+4)) = v1362
	*(*int32)(unsafe.Add(mBase, uint32(v1491))) = v1489
	v1495 = *(*int32)(unsafe.Add(mBase, uint32(v1343)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1491)+8)) = v1495
	v1497 = *(*int32)(unsafe.Add(mBase, uint32(v1343)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1491)+12)) = v1497
	v1499 = *(*int32)(unsafe.Add(mBase, uint32(v1343)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v1491)+16)) = v1499
	v1501 = *(*int32)(unsafe.Add(mBase, uint32(v1343)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v1491)+20)) = v1501
	v1503 = *(*int32)(unsafe.Add(mBase, uint32(v1343)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v1491)+24)) = v1503
	v1505 = *(*int32)(unsafe.Add(mBase, uint32(v1343)+28))
	v1506 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1491)+40)) = v1506
	*(*uint16)(unsafe.Add(mBase, uint32(v1491)+38)) = uint16(v1506)
	*(*int32)(unsafe.Add(mBase, uint32(v1491)+34)) = int32(-1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1491)+32)) = uint8(v1506)
	*(*int32)(unsafe.Add(mBase, uint32(v1491)+28)) = v1505
	v1515 = *(*int32)(unsafe.Add(mBase, uint32(v1262)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v1515+v1495<<(uint(int32(2))%32)-int32(4)))) = v1491
	goto L225
L268:
	;
	goto L224
L269:
	;
	v1657 = F_ExecInitNode(m, v1284, v1262, v1261)
	mBase = m.M
	v1658 = m.ExcPending
	if v1658 != 0 {
		goto L9
	} else {
		goto L281
	}
L270:
	;
	v1568 = *(*int32)(unsafe.Add(mBase, uint32(v1565)+4))
	if v1568 <= int32(0) {
		goto L269
	} else {
		goto L271
	}
L271:
	;
	v1572 = v1261 & int32(-29)
	v1576 = v1560
	v1579 = int32(1)
	goto L272
L272:
	;
	v1606 = *(*int32)(unsafe.Add(mBase, uint32(v1565)+12))
	v1610 = *(*int32)(unsafe.Add(mBase, uint32(v1606+v1576<<(uint(int32(2))%32))))
	v1611 = *(*int32)(unsafe.Add(mBase, uint32(v1270)+68))
	v1612 = F_bms_is_member(m, v1579, v1611)
	mBase = m.M
	v1613 = m.ExcPending
	if v1613 != 0 {
		goto L9
	} else {
		goto L274
	}
L273:
	;
	goto L269
L274:
	;
	if v1612 != 0 {
		goto L275
	} else {
		goto L276
	}
L275:
	;
	v1614 = v1572 | int32(4)
	goto L277
L276:
	;
	v1614 = v1572
	goto L277
L277:
	;
	v1615 = F_ExecInitNode(m, v1610, v1262, v1614)
	mBase = m.M
	v1616 = m.ExcPending
	if v1616 != 0 {
		goto L9
	} else {
		goto L278
	}
L278:
	;
	v1617 = *(*int32)(unsafe.Add(mBase, uint32(v1262)+144))
	v1618 = F_lappend(m, v1617, v1615)
	mBase = m.M
	v1619 = m.ExcPending
	if v1619 != 0 {
		goto L9
	} else {
		goto L279
	}
L279:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1262)+144)) = v1618
	v1621 = int32(1)
	v1624 = v1576 + v1621
	v1625 = *(*int32)(unsafe.Add(mBase, uint32(v1565)+4))
	if v1624 < v1625 {
		v1576 = v1624
		v1579 = v1579 + v1621
		goto L272
	} else {
		goto L280
	}
L280:
	;
	goto L273
L281:
	;
	v1659 = *(*int32)(unsafe.Add(mBase, uint32(v1657)+56))
	if v1286 != int32(1) {
		v1772 = v1659
		goto L214
	} else {
		goto L282
	}
L282:
	;
	v1662 = *(*int32)(unsafe.Add(mBase, uint32(v1284)+44))
	if v1662 == int32(0) {
		v1772 = v1659
		goto L214
	} else {
		goto L283
	}
L283:
	;
	v1665 = *(*int32)(unsafe.Add(mBase, uint32(v1662)+4))
	if v1665 <= int32(0) {
		v1772 = v1659
		goto L214
	} else {
		goto L284
	}
L284:
	;
	v1668 = *(*int32)(unsafe.Add(mBase, uint32(v1662)+12))
	goto L215
L285:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v1676 = m.ExcPending
	if v1676 != 0 {
		goto L9
	} else {
		goto L286
	}
L286:
	;
	v1677 = *(*int32)(unsafe.Add(mBase, uint32(v1387)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v1269)+112)) = v1677 + int32(4)
	F_errmsg(m, int32(702223), v1269+int32(112))
	mBase = m.M
	v1685 = m.ExcPending
	if v1685 != 0 {
		goto L9
	} else {
		goto L287
	}
L287:
	;
	F_errfinish(m, int32(494876), int32(1228), int32(307494))
	mBase = m.M
	v1690 = m.ExcPending
	if v1690 != 0 {
		goto L9
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
	F_errcode(m, int32(151027844))
	mBase = m.M
	v1697 = m.ExcPending
	if v1697 != 0 {
		goto L9
	} else {
		goto L290
	}
L290:
	;
	v1698 = *(*int32)(unsafe.Add(mBase, uint32(v1387)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v1269)+32)) = v1698 + int32(4)
	F_errmsg(m, int32(689208), v1269+int32(32))
	mBase = m.M
	v1706 = m.ExcPending
	if v1706 != 0 {
		goto L9
	} else {
		goto L291
	}
L291:
	;
	F_errfinish(m, int32(494876), int32(1234), int32(307494))
	mBase = m.M
	v1711 = m.ExcPending
	if v1711 != 0 {
		goto L9
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
	v1745 = *(*int32)(unsafe.Add(mBase, uint32(v1668+v1712<<(uint(int32(2))%32))))
	v1746 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1745)+26)))
	if v1746 != int32(1) {
		goto L295
	} else {
		goto L296
	}
L294:
	;
	v1754 = F_ExecInitExtraTupleSlot(m, v1262, int32(0), int32(1600612))
	mBase = m.M
	v1755 = m.ExcPending
	if v1755 != 0 {
		goto L9
	} else {
		goto L299
	}
L295:
	;
	v1750 = v1712 + int32(1)
	if v1750 != v1665 {
		v1712 = v1750
		goto L293
	} else {
		goto L298
	}
L296:
	;
	goto L297
L297:
	;
	goto L294
L298:
	;
	v1772 = v1659
	goto L214
L299:
	;
	v1756 = *(*int32)(unsafe.Add(mBase, uint32(v1657)+4))
	v1757 = *(*int32)(unsafe.Add(mBase, uint32(v1756)+44))
	v1758 = F_ExecInitJunkFilter(m, v1757, v1754)
	mBase = m.M
	v1759 = m.ExcPending
	if v1759 != 0 {
		goto L9
	} else {
		goto L300
	}
L300:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1262)+60)) = v1758
	v1761 = *(*int32)(unsafe.Add(mBase, uint32(v1758)+8))
	v1772 = v1761
	goto L214
}
