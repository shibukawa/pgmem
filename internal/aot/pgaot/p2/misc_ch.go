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
	if base.Ui32(v50) < base.Ui32(int32(_a_F_CheckCmdReplicaIdentity_0)) {
		v59 = v3
		goto L11
	} else {
		goto L14
	}
L14:
	;
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46)+118)))
	v59 = base.B2i32(v53 == int32(112)) & base.B2i32(base.Ui32(int32(_a_F_CheckCmdReplicaIdentity_1)) < base.Ui32(v50))
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
	v206 = F_table_open(m, int32(_a_F_CheckCmdReplicaIdentity_2), int32(1))
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
	v1087 = int32(_a_F_CheckCmdReplicaIdentity_3)
	v1088 = *(*int32)(unsafe.Add(mBase, _c_F_CheckCmdReplicaIdentity[0]))
	v1091 = *(*int32)(unsafe.Add(mBase, _c_F_CheckCmdReplicaIdentity[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_CheckCmdReplicaIdentity[0])) = v1091
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
	*(*int32)(unsafe.Add(mBase, _c_F_CheckCmdReplicaIdentity[0])) = v1088
	goto L17
L222:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43))) = v307
	F_errmsg_internal(m, int32(_a_F_CheckCmdReplicaIdentity_4), v43)
	mBase = m.M
	v1141 = m.ExcPending
	if v1141 != 0 {
		goto L24
	} else {
		goto L223
	}
L223:
	;
	F_errfinish(m, int32(_a_F_CheckCmdReplicaIdentity_5), int32(_a_F_CheckCmdReplicaIdentity_6), int32(_a_F_CheckCmdReplicaIdentity_7))
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
	F_errcode(m, int32(_a_F_CheckCmdReplicaIdentity_8))
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
	F_errmsg(m, int32(_a_F_CheckCmdReplicaIdentity_9), v31+int32(32))
	mBase = m.M
	v1172 = m.ExcPending
	if v1172 != 0 {
		goto L24
	} else {
		goto L234
	}
L234:
	;
	F_errdetail(m, int32(_a_F_CheckCmdReplicaIdentity_10), int32(0))
	mBase = m.M
	v1176 = m.ExcPending
	if v1176 != 0 {
		goto L24
	} else {
		goto L235
	}
L235:
	;
	F_errfinish(m, int32(_a_F_CheckCmdReplicaIdentity_11), int32(823), int32(_a_F_CheckCmdReplicaIdentity_12))
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
	F_errcode(m, int32(_a_F_CheckCmdReplicaIdentity_8))
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
	F_errmsg(m, int32(_a_F_CheckCmdReplicaIdentity_9), v31-int32(-64))
	mBase = m.M
	v1258 = m.ExcPending
	if v1258 != 0 {
		goto L24
	} else {
		goto L252
	}
L252:
	;
	F_errdetail(m, int32(_a_F_CheckCmdReplicaIdentity_13), int32(0))
	mBase = m.M
	v1262 = m.ExcPending
	if v1262 != 0 {
		goto L24
	} else {
		goto L253
	}
L253:
	;
	F_errfinish(m, int32(_a_F_CheckCmdReplicaIdentity_11), int32(811), int32(_a_F_CheckCmdReplicaIdentity_12))
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
	F_errcode(m, int32(_a_F_CheckCmdReplicaIdentity_8))
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
	F_errmsg(m, int32(_a_F_CheckCmdReplicaIdentity_9), v31+int32(48))
	mBase = m.M
	v1283 = m.ExcPending
	if v1283 != 0 {
		goto L24
	} else {
		goto L257
	}
L257:
	;
	F_errdetail(m, int32(_a_F_CheckCmdReplicaIdentity_14), int32(0))
	mBase = m.M
	v1287 = m.ExcPending
	if v1287 != 0 {
		goto L24
	} else {
		goto L258
	}
L258:
	;
	F_errfinish(m, int32(_a_F_CheckCmdReplicaIdentity_11), int32(817), int32(_a_F_CheckCmdReplicaIdentity_12))
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
	F_errcode(m, int32(_a_F_CheckCmdReplicaIdentity_8))
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
	F_errmsg(m, int32(_a_F_CheckCmdReplicaIdentity_15), v31+int32(112))
	mBase = m.M
	v1308 = m.ExcPending
	if v1308 != 0 {
		goto L24
	} else {
		goto L262
	}
L262:
	;
	F_errdetail(m, int32(_a_F_CheckCmdReplicaIdentity_13), int32(0))
	mBase = m.M
	v1312 = m.ExcPending
	if v1312 != 0 {
		goto L24
	} else {
		goto L263
	}
L263:
	;
	F_errfinish(m, int32(_a_F_CheckCmdReplicaIdentity_11), int32(829), int32(_a_F_CheckCmdReplicaIdentity_12))
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
	F_errcode(m, int32(_a_F_CheckCmdReplicaIdentity_8))
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
	F_errmsg(m, int32(_a_F_CheckCmdReplicaIdentity_15), v31+int32(96))
	mBase = m.M
	v1333 = m.ExcPending
	if v1333 != 0 {
		goto L24
	} else {
		goto L267
	}
L267:
	;
	F_errdetail(m, int32(_a_F_CheckCmdReplicaIdentity_14), int32(0))
	mBase = m.M
	v1337 = m.ExcPending
	if v1337 != 0 {
		goto L24
	} else {
		goto L268
	}
L268:
	;
	F_errfinish(m, int32(_a_F_CheckCmdReplicaIdentity_11), int32(835), int32(_a_F_CheckCmdReplicaIdentity_12))
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
	F_errcode(m, int32(_a_F_CheckCmdReplicaIdentity_8))
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
	F_errmsg(m, int32(_a_F_CheckCmdReplicaIdentity_15), v31+int32(80))
	mBase = m.M
	v1358 = m.ExcPending
	if v1358 != 0 {
		goto L24
	} else {
		goto L272
	}
L272:
	;
	F_errdetail(m, int32(_a_F_CheckCmdReplicaIdentity_10), int32(0))
	mBase = m.M
	v1362 = m.ExcPending
	if v1362 != 0 {
		goto L24
	} else {
		goto L273
	}
L273:
	;
	F_errfinish(m, int32(_a_F_CheckCmdReplicaIdentity_11), int32(841), int32(_a_F_CheckCmdReplicaIdentity_12))
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
	F_errmsg(m, int32(_a_F_CheckCmdReplicaIdentity_16), v31)
	mBase = m.M
	v1381 = m.ExcPending
	if v1381 != 0 {
		goto L24
	} else {
		goto L277
	}
L277:
	;
	F_errhint(m, int32(_a_F_CheckCmdReplicaIdentity_17), int32(0))
	mBase = m.M
	v1385 = m.ExcPending
	if v1385 != 0 {
		goto L24
	} else {
		goto L278
	}
L278:
	;
	F_errfinish(m, int32(_a_F_CheckCmdReplicaIdentity_11), int32(861), int32(_a_F_CheckCmdReplicaIdentity_12))
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
	F_errmsg(m, int32(_a_F_CheckCmdReplicaIdentity_18), v31+int32(16))
	mBase = m.M
	v1406 = m.ExcPending
	if v1406 != 0 {
		goto L24
	} else {
		goto L282
	}
L282:
	;
	F_errhint(m, int32(_a_F_CheckCmdReplicaIdentity_19), int32(0))
	mBase = m.M
	v1410 = m.ExcPending
	if v1410 != 0 {
		goto L24
	} else {
		goto L283
	}
L283:
	;
	F_errfinish(m, int32(_a_F_CheckCmdReplicaIdentity_11), int32(867), int32(_a_F_CheckCmdReplicaIdentity_12))
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
															if int32(1)<<(uint(v56)%32)&int32(_a_F_charclasscomplement_0) == int32(0) {
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
															if int32(1)<<(uint(v138)%32)&int32(_a_F_charclasscomplement_0) == int32(0) {
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
								v96 = *(*int32)(unsafe.Add(mBase, _c_F_checkTargetlistEntrySQL92[0]))
								*(*int32)(unsafe.Add(mBase, uint32(v7))) = v96
								F_errmsg(m, int32(_a_F_checkTargetlistEntrySQL92_0), v7)
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
											F_errfinish(m, int32(_a_F_checkTargetlistEntrySQL92_1), int32(1965), int32(_a_F_checkTargetlistEntrySQL92_2))
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
										F_errcode(m, int32(_a_F_checkTargetlistEntrySQL92_3))
										mBase = m.M
										v36 = m.ExcPending
										if v36 != 0 {
											return
										} else {
											v46 = *(*int32)(unsafe.Add(mBase, _c_F_checkTargetlistEntrySQL92[0]))
											*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v46
											F_errmsg(m, int32(_a_F_checkTargetlistEntrySQL92_4), v7+int32(16))
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
														F_errfinish(m, int32(_a_F_checkTargetlistEntrySQL92_1), int32(1974), int32(_a_F_checkTargetlistEntrySQL92_2))
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
								F_errcode(m, int32(_a_F_checkTargetlistEntrySQL92_3))
								mBase = m.M
								v36 = m.ExcPending
								if v36 != 0 {
									return
								} else {
									v46 = *(*int32)(unsafe.Add(mBase, _c_F_checkTargetlistEntrySQL92[0]))
									*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v46
									F_errmsg(m, int32(_a_F_checkTargetlistEntrySQL92_4), v7+int32(16))
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
												F_errfinish(m, int32(_a_F_checkTargetlistEntrySQL92_1), int32(1974), int32(_a_F_checkTargetlistEntrySQL92_2))
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
				F_errmsg_internal(m, int32(_a_F_checkTargetlistEntrySQL92_5), int32(0))
				mBase = m.M
				v71 = m.ExcPending
				if v71 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_checkTargetlistEntrySQL92_1), int32(1983), int32(_a_F_checkTargetlistEntrySQL92_2))
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
	F_errmsg(m, int32(_a_F_check_agglevels_and_constraints_0), int32(0))
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
	F_errfinish(m, int32(_a_F_check_agglevels_and_constraints_1), int32(693), int32(_a_F_check_agglevels_and_constraints_2))
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
		v251 = int32(_a_F_check_agglevels_and_constraints_3)
		v252 = int32(_a_F_check_agglevels_and_constraints_4)
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
	v251 = int32(_a_F_check_agglevels_and_constraints_46)
	v252 = int32(_a_F_check_agglevels_and_constraints_47)
	goto L61
L63:
	;
	v251 = int32(_a_F_check_agglevels_and_constraints_44)
	v252 = int32(_a_F_check_agglevels_and_constraints_45)
	goto L61
L64:
	;
	v251 = int32(_a_F_check_agglevels_and_constraints_48)
	v252 = int32(_a_F_check_agglevels_and_constraints_49)
	goto L61
L65:
	;
	v251 = int32(_a_F_check_agglevels_and_constraints_42)
	v252 = int32(_a_F_check_agglevels_and_constraints_43)
	goto L61
L66:
	;
	v251 = int32(_a_F_check_agglevels_and_constraints_40)
	v252 = int32(_a_F_check_agglevels_and_constraints_41)
	goto L61
L67:
	;
	v251 = int32(_a_F_check_agglevels_and_constraints_36)
	v252 = int32(_a_F_check_agglevels_and_constraints_37)
	goto L61
L68:
	;
	v251 = int32(_a_F_check_agglevels_and_constraints_34)
	v252 = int32(_a_F_check_agglevels_and_constraints_35)
	goto L61
L69:
	;
	v251 = int32(_a_F_check_agglevels_and_constraints_32)
	v252 = int32(_a_F_check_agglevels_and_constraints_33)
	goto L61
L70:
	;
	v251 = int32(_a_F_check_agglevels_and_constraints_30)
	v252 = int32(_a_F_check_agglevels_and_constraints_31)
	goto L61
L71:
	;
	v251 = int32(_a_F_check_agglevels_and_constraints_28)
	v252 = int32(_a_F_check_agglevels_and_constraints_29)
	goto L61
L72:
	;
	v251 = int32(_a_F_check_agglevels_and_constraints_26)
	v252 = int32(_a_F_check_agglevels_and_constraints_27)
	goto L61
L73:
	;
	v251 = int32(_a_F_check_agglevels_and_constraints_24)
	v252 = int32(_a_F_check_agglevels_and_constraints_25)
	goto L61
L74:
	;
	v251 = int32(_a_F_check_agglevels_and_constraints_22)
	v252 = int32(_a_F_check_agglevels_and_constraints_23)
	goto L61
L75:
	;
	v251 = int32(_a_F_check_agglevels_and_constraints_20)
	v252 = int32(_a_F_check_agglevels_and_constraints_21)
	goto L61
L76:
	;
	v251 = int32(_a_F_check_agglevels_and_constraints_18)
	v252 = int32(_a_F_check_agglevels_and_constraints_19)
	goto L61
L77:
	;
	v251 = int32(_a_F_check_agglevels_and_constraints_16)
	v252 = int32(_a_F_check_agglevels_and_constraints_17)
	goto L61
L78:
	;
	v251 = int32(_a_F_check_agglevels_and_constraints_14)
	v252 = int32(_a_F_check_agglevels_and_constraints_15)
	goto L61
L79:
	;
	v251 = int32(_a_F_check_agglevels_and_constraints_38)
	v252 = int32(_a_F_check_agglevels_and_constraints_39)
	goto L61
L80:
	;
	v251 = int32(_a_F_check_agglevels_and_constraints_9)
	v252 = int32(_a_F_check_agglevels_and_constraints_10)
	goto L61
L81:
	;
	v251 = int32(_a_F_check_agglevels_and_constraints_7)
	v252 = int32(_a_F_check_agglevels_and_constraints_8)
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
	F_errmsg_internal(m, int32(_a_F_check_agglevels_and_constraints_5), v15+int32(16))
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
	F_errfinish(m, int32(_a_F_check_agglevels_and_constraints_1), int32(600), int32(_a_F_check_agglevels_and_constraints_6))
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
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v283<<(uint(int32(2))%32))+uint32(_c_F_check_agglevels_and_constraints[0])))
	v293 = v292
	goto L95
L94:
	;
	v293 = int32(_a_F_check_agglevels_and_constraints_11)
	goto L95
L95:
	;
	goto L92
L96:
	;
	v299 = int32(_a_F_check_agglevels_and_constraints_12)
	goto L98
L97:
	;
	v299 = int32(_a_F_check_agglevels_and_constraints_13)
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
	F_errfinish(m, int32(_a_F_check_agglevels_and_constraints_1), int32(615), int32(_a_F_check_agglevels_and_constraints_6))
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
	F_errmsg(m, int32(_a_F_check_agglevels_and_constraints_51), int32(0))
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
	F_errdetail(m, int32(_a_F_check_agglevels_and_constraints_52), v15)
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
	F_errfinish(m, int32(_a_F_check_agglevels_and_constraints_1), int32(708), int32(_a_F_check_agglevels_and_constraints_2))
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
	F_errmsg(m, int32(_a_F_check_agglevels_and_constraints_50), int32(0))
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
	F_errfinish(m, int32(_a_F_check_agglevels_and_constraints_1), int32(731), int32(_a_F_check_agglevels_and_constraints_2))
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
	F_errmsg(m, int32(_a_F_check_agglevels_and_constraints_0), int32(0))
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
	F_errfinish(m, int32(_a_F_check_agglevels_and_constraints_1), int32(738), int32(_a_F_check_agglevels_and_constraints_2))
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
	F_errmsg(m, int32(_a_F_check_agglevels_and_constraints_51), int32(0))
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
	F_errdetail(m, int32(_a_F_check_agglevels_and_constraints_52), v15+int32(48))
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
	F_errfinish(m, int32(_a_F_check_agglevels_and_constraints_1), int32(745), int32(_a_F_check_agglevels_and_constraints_2))
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
		v10 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_check_stage_log_stats[0])))
		if v10 != v8 {
			v26 = v8
			return v26
		} else {
			v14 = *(*int32)(unsafe.Add(mBase, _c_F_check_stage_log_stats[1]))
			*(*int32)(unsafe.Add(mBase, _c_F_check_stage_log_stats[2])) = v14
			v20 = F_format_elog_string(m, int32(_a_F_check_stage_log_stats_0), int32(0))
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, _c_F_check_stage_log_stats[3])) = v20
				v26 = int32(0)
				return v26
			}
		}
	}
}
