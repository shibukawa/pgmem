package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_TerminateLocalBufferIO(m *base.Module, l0 int32, l1 int64, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int64
	_ = v4
	var v8 int64
	_ = v8
	var v10 int64
	_ = v10
	var v17 int64
	_ = v17
	v4 = int64(0)
	v8 = base.AtomicRmwCmpxchg64(m, l0, int32(24), v4, v4)
	v10 = v8 & int64(-134217729)
	if l2 != 0 {
		*(*int32)(unsafe.Add(mBase, uint32(l0+int32(36)))) = int32(-1)
		v17 = v10 - int64(1)
	} else {
		v17 = v10
	}
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = v17 | l1
	return
}
func F_TestConfiguration(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var __phi89 int32
	_ = __phi89
	var v100 int32
	_ = v100
	var __phi100 int32
	_ = __phi100
	var v109 int32
	_ = v109
	var __phi109 int32
	_ = __phi109
	var v116 int32
	_ = v116
	var __phi116 int32
	_ = __phi116
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v165 int32
	_ = v165
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v234 int32
	_ = v234
	var v273 int32
	_ = v273
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v287 int32
	_ = v287
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v315 int32
	_ = v315
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v347 int32
	_ = v347
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v372 int32
	_ = v372
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v382 int32
	_ = v382
	var v410 int32
	_ = v410
	var v412 int32
	_ = v412
	var v414 int32
	_ = v414
	var v421 int32
	_ = v421
	var v423 int32
	_ = v423
	var v425 int32
	_ = v425
	var v427 int32
	_ = v427
	var v429 int32
	_ = v429
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v440 int32
	_ = v440
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v481 int32
	_ = v481
	var v483 int32
	_ = v483
	var v485 int32
	_ = v485
	var v528 int32
	_ = v528
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v545 int32
	_ = v545
	var v573 int32
	_ = v573
	var v575 int32
	_ = v575
	var v577 int32
	_ = v577
	var v584 int32
	_ = v584
	var v586 int32
	_ = v586
	var v588 int32
	_ = v588
	var v590 int32
	_ = v590
	var v592 int32
	_ = v592
	var v599 int32
	_ = v599
	var v600 int32
	_ = v600
	var v601 int32
	_ = v601
	var v603 int32
	_ = v603
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v644 int32
	_ = v644
	var v646 int32
	_ = v646
	var v648 int32
	_ = v648
	var v691 int32
	_ = v691
	var v694 int32
	_ = v694
	var v696 int32
	_ = v696
	var v697 int32
	_ = v697
	var v698 int32
	_ = v698
	var v704 int32
	_ = v704
	var v705 int32
	_ = v705
	var v747 int32
	_ = v747
	var v790 int32
	_ = v790
	var v800 int32
	_ = v800
	var v826 int32
	_ = v826
	var v827 int32
	_ = v827
	var v861 int32
	_ = v861
	var v868 int32
	_ = v868
	var v875 int32
	_ = v875
	var v911 int32
	_ = v911
	var v913 int32
	_ = v913
	var v915 int32
	_ = v915
	var v923 int32
	_ = v923
	var v924 int32
	_ = v924
	var v925 int32
	_ = v925
	var v928 int32
	_ = v928
	var v936 int32
	_ = v936
	var v965 int32
	_ = v965
	var v966 int32
	_ = v966
	var v970 int32
	_ = v970
	var v981 int32
	_ = v981
	var v983 int32
	_ = v983
	var v986 int32
	_ = v986
	var v990 int32
	_ = v990
	var v1027 int32
	_ = v1027
	var v1030 int32
	_ = v1030
	var v1033 int32
	_ = v1033
	var v1034 int32
	_ = v1034
	var v1040 int32
	_ = v1040
	var v1079 int32
	_ = v1079
	var v1121 int32
	_ = v1121
	var v1144 int32
	_ = v1144
	var v1151 int32
	_ = v1151
	var v1163 int32
	_ = v1163
	var v1168 int32
	_ = v1168
	var v1171 int32
	_ = v1171
	var v1204 int32
	_ = v1204
	var v1206 int32
	_ = v1206
	var v1208 int32
	_ = v1208
	var v1210 int32
	_ = v1210
	var v1219 int32
	_ = v1219
	var v1223 int32
	_ = v1223
	var v1224 int32
	_ = v1224
	var v1225 int32
	_ = v1225
	var v1227 int32
	_ = v1227
	var v1229 int32
	_ = v1229
	var v1232 int32
	_ = v1232
	var v1242 int32
	_ = v1242
	var v1249 int32
	_ = v1249
	var v1266 int32
	_ = v1266
	var v1269 int32
	_ = v1269
	var v1273 int32
	_ = v1273
	var v1277 int32
	_ = v1277
	var v1279 int32
	_ = v1279
	var v1286 int32
	_ = v1286
	var v1289 int32
	_ = v1289
	var v1293 int32
	_ = v1293
	var v1295 int32
	_ = v1295
	var v1300 int32
	_ = v1300
	var v1317 int32
	_ = v1317
	var v1320 int32
	_ = v1320
	var v1322 int32
	_ = v1322
	var v1324 int32
	_ = v1324
	var v1326 int32
	_ = v1326
	var v1328 int32
	_ = v1328
	var v1337 int32
	_ = v1337
	var v1341 int32
	_ = v1341
	var v1342 int32
	_ = v1342
	var v1343 int32
	_ = v1343
	var v1345 int32
	_ = v1345
	var v1347 int32
	_ = v1347
	var v1350 int32
	_ = v1350
	var v1360 int32
	_ = v1360
	var v1367 int32
	_ = v1367
	var v1384 int32
	_ = v1384
	var v1387 int32
	_ = v1387
	var v1391 int32
	_ = v1391
	var v1395 int32
	_ = v1395
	var v1397 int32
	_ = v1397
	var v1404 int32
	_ = v1404
	var v1407 int32
	_ = v1407
	var v1411 int32
	_ = v1411
	var v1413 int32
	_ = v1413
	var v1418 int32
	_ = v1418
	var v1435 int32
	_ = v1435
	var v1438 int32
	_ = v1438
	var v1440 int32
	_ = v1440
	var v1442 int32
	_ = v1442
	var v1444 int32
	_ = v1444
	var v1450 int32
	_ = v1450
	var v1453 int32
	_ = v1453
	var v1483 int32
	_ = v1483
	var v1492 int32
	_ = v1492
	var v1496 int32
	_ = v1496
	var v1497 int32
	_ = v1497
	var v1498 int32
	_ = v1498
	var v1500 int32
	_ = v1500
	var v1502 int32
	_ = v1502
	var v1505 int32
	_ = v1505
	var v1515 int32
	_ = v1515
	var v1522 int32
	_ = v1522
	var v1539 int32
	_ = v1539
	var v1542 int32
	_ = v1542
	var v1546 int32
	_ = v1546
	var v1550 int32
	_ = v1550
	var v1552 int32
	_ = v1552
	var v1559 int32
	_ = v1559
	var v1562 int32
	_ = v1562
	var v1566 int32
	_ = v1566
	var v1568 int32
	_ = v1568
	var v1573 int32
	_ = v1573
	var v1590 int32
	_ = v1590
	var v1593 int32
	_ = v1593
	var v1595 int32
	_ = v1595
	var v1600 int32
	_ = v1600
	var v1603 int32
	_ = v1603
	v2 = int32(0)
	v37 = m.G0
	v39 = v37 - int32(16)
	m.G0 = v39
	v42 = *(*int32)(unsafe.Add(mBase, _c_F_TestConfiguration[0]))
	v44 = *(*int32)(unsafe.Add(mBase, _c_F_TestConfiguration[1]))
	v46 = *(*int32)(unsafe.Add(mBase, _c_F_TestConfiguration[2]))
	if v42 < v44+v46 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v1603 + int32(16)
	return v1600
L2:
	;
	v1600 = int32(-1)
	v1603 = v39
	goto L1
L3:
	;
	goto L4
L4:
	;
	v51 = *(*int32)(unsafe.Add(mBase, _c_F_TestConfiguration[3]))
	v54 = v51 + v44*int32(20)
	v56 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_TestConfiguration[4])) = v56
	v59 = *(*int32)(unsafe.Add(mBase, _c_F_TestConfiguration[5]))
	v61 = v59 - int32(1)
	if v61 < v56 {
		v1450 = v2
		v1453 = v39
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v1483 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_TestConfiguration[6])) = v1483
	*(*int32)(unsafe.Add(mBase, _c_F_TestConfiguration[7])) = v1483
	*(*int32)(unsafe.Add(mBase, uint32(v1453)+12)) = v1483
	v1492 = v1453 + int32(12)
	v1496 = *(*int32)(unsafe.Add(mBase, uint32(l0)+364))
	if v1496 != 0 {
		goto L218
	} else {
		goto L219
	}
L6:
	;
	v65 = *(*int32)(unsafe.Add(mBase, _c_F_TestConfiguration[8]))
	v67 = *(*int32)(unsafe.Add(mBase, _c_F_TestConfiguration[9]))
	v68 = int32(-1)
	v70 = int32(4)
	v71 = v67 + v70
	v72 = int32(3)
	v73 = v67 & v72
	v75 = *(*int32)(unsafe.Add(mBase, _c_F_TestConfiguration[10]))
	v79 = v75 + v70
	v83 = *(*int32)(unsafe.Add(mBase, _c_F_TestConfiguration[11]))
	v85 = *(*int32)(unsafe.Add(mBase, _c_F_TestConfiguration[12]))
	v87 = *(*int32)(unsafe.Add(mBase, _c_F_TestConfiguration[13]))
	__phi89 = v61
	__phi100 = v59
	__phi109 = v2
	__phi116 = v2
	v89 = __phi89
	v100 = __phi100
	v109 = __phi109
	v116 = __phi116
	goto L7
L7:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v65+v89*int32(20))+8))
	v129 = v109
	goto L10
L8:
	;
	v1163 = int32(0)
	if v59 <= v1163 {
		v1450 = v1163
		v1453 = v39
		goto L5
	} else {
		goto L147
	}
L9:
	;
	if int32(0) < v89 {
		__phi89 = v89 - int32(1)
		__phi100 = v89
		__phi109 = v1144
		__phi116 = v1151
		v89 = __phi89
		v100 = __phi100
		v109 = __phi109
		v116 = __phi116
		goto L7
	} else {
		goto L146
	}
L10:
	;
	v165 = v129 - int32(1)
	if int32(0) <= v165 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v175 = v87 + v109*int32(12)
	v178 = v85 + v116<<(uint(int32(2))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v175)+4)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v175))) = v127
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v127)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v175)+8)) = v181
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v127)+36))
	if v183 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L12:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v87+v165*int32(12))))
	if v171 != v127 {
		v129 = v165
		goto L10
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	goto L11
L15:
	;
	v1144 = v109
	v1151 = v116
	goto L9
L16:
	;
	v273 = v181 << (uint(int32(2)) % 32)
	if v75&v72 != 0 {
		v287 = v273
		goto L25
	} else {
		goto L26
	}
L17:
	;
	v188 = v127 + int32(32)
	if v183 == v188 {
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v191 = v183
	v192 = int32(0)
	goto L19
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v83+v192<<(uint(int32(2))%32)))) = v191 - int32(388)
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v191)+4))
	if v234 != v188 {
		v191 = v234
		v192 = v192 + int32(1)
		goto L19
	} else {
		goto L21
	}
L20:
	;
	goto L16
L21:
	;
	goto L20
L22:
	;
	v315 = v181 - int32(1)
	if int32(0) < v100 {
		goto L42
	} else {
		goto L43
	}
L23:
	;
	if v306 == int32(0) {
		goto L22
	} else {
		goto L41
	}
L24:
	;
	v303 = int32(0)
	if v73 == v303 {
		goto L22
	} else {
		goto L40
	}
L25:
	;
	if v287 != 0 {
		goto L32
	} else {
		goto L33
	}
L26:
	;
	if base.Ui32(int32(1024)) < base.Ui32(v273) {
		v287 = v273
		goto L25
	} else {
		goto L27
	}
L27:
	;
	if v273 == int32(0) {
		goto L24
	} else {
		goto L28
	}
L28:
	;
	v278 = v273 + v75
	if base.Ui32(v79) < base.Ui32(v278) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v280 = v278
	goto L31
L30:
	;
	v280 = v79
	goto L31
L31:
	;
	v287 = (v280+(v75^v68))&int32(-4) + int32(4)
	goto L25
L32:
	;
	base.MemoryFill(m, v75, int32(0), v287)
	goto L34
L33:
	;
	goto L34
L34:
	;
	if v73|base.B2i32(base.Ui32(int32(1025)) <= base.Ui32(v273)) != 0 {
		v306 = v273
		goto L23
	} else {
		goto L35
	}
L35:
	;
	if v273 == int32(0) {
		goto L22
	} else {
		goto L36
	}
L36:
	;
	v295 = v273 + v67
	if base.Ui32(v71) < base.Ui32(v295) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v297 = v295
	goto L39
L38:
	;
	v297 = v71
	goto L39
L39:
	;
	v306 = (v297+(v67^v68))&int32(-4) + int32(4)
	goto L23
L40:
	;
	v306 = v303
	goto L23
L41:
	;
	base.MemoryFill(m, v67, int32(0), v306)
	goto L22
L42:
	;
	v319 = v181 & int32(-2)
	v320 = int32(1)
	v321 = v181 & v320
	v347 = int32(0)
	goto L45
L43:
	;
	goto L44
L44:
	;
	if int32(0) <= v315 {
		goto L108
	} else {
		goto L109
	}
L45:
	;
	if v315 < int32(0) {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	goto L44
L47:
	;
	v747 = v347 + int32(1)
	if v747 != v100 {
		v347 = v747
		goto L45
	} else {
		goto L107
	}
L48:
	;
	v369 = v65 + v347*int32(20)
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v369)))
	v372 = int32(-1)
	if v315 != 0 {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	if v528 < int32(0) {
		goto L47
	} else {
		goto L77
	}
L50:
	;
	v374 = v315
	v375 = v372
	v382 = int32(0)
	goto L53
L51:
	;
	v445 = v315
	v446 = v372
	goto L52
L52:
	;
	v481 = v445 << (uint(int32(2)) % 32)
	v483 = *(*int32)(unsafe.Add(mBase, uint32(v83+v481)))
	if v370 != v483 {
		goto L72
	} else {
		goto L73
	}
L53:
	;
	v410 = v374 << (uint(int32(2)) % 32)
	v412 = *(*int32)(unsafe.Add(mBase, uint32(v83+v410)))
	if v370 != v412 {
		goto L57
	} else {
		goto L58
	}
L54:
	;
	if v321 == int32(0) {
		v528 = v436
		goto L49
	} else {
		goto L70
	}
L55:
	;
	v423 = v374 - int32(1)
	v425 = v423 << (uint(int32(2)) % 32)
	v427 = *(*int32)(unsafe.Add(mBase, uint32(v83+v425)))
	if v370 != v427 {
		goto L64
	} else {
		goto L65
	}
L56:
	;
	v421 = v375
	goto L55
L57:
	;
	v414 = *(*int32)(unsafe.Add(mBase, uint32(v412)+364))
	if v414 != v370 {
		goto L56
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	if v375 == int32(-1) {
		v421 = v374
		goto L55
	} else {
		goto L61
	}
L60:
	;
	goto L59
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v410+v75))) = int32(-1)
	goto L56
L62:
	;
	v437 = int32(2)
	v438 = v374 - v437
	v440 = v382 + v437
	if v440 != v319 {
		v374 = v438
		v375 = v436
		v382 = v440
		goto L53
	} else {
		goto L69
	}
L63:
	;
	v436 = v421
	goto L62
L64:
	;
	v429 = *(*int32)(unsafe.Add(mBase, uint32(v427)+364))
	if v429 != v370 {
		goto L63
	} else {
		goto L67
	}
L65:
	;
	goto L66
L66:
	;
	if v421 == int32(-1) {
		v436 = v423
		goto L62
	} else {
		goto L68
	}
L67:
	;
	goto L66
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v75+v425))) = int32(-1)
	goto L63
L69:
	;
	goto L54
L70:
	;
	v445 = v438
	v446 = v436
	goto L52
L71:
	;
	v528 = v446
	goto L49
L72:
	;
	v485 = *(*int32)(unsafe.Add(mBase, uint32(v483)+364))
	if v485 != v370 {
		goto L71
	} else {
		goto L75
	}
L73:
	;
	goto L74
L74:
	;
	if v446 == int32(-1) {
		v528 = v445
		goto L49
	} else {
		goto L76
	}
L75:
	;
	goto L74
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v481+v75))) = int32(-1)
	goto L71
L77:
	;
	v531 = *(*int32)(unsafe.Add(mBase, uint32(v369)+4))
	v532 = int32(0)
	v533 = int32(-1)
	if base.B2i32(v181 == v315>>(uint(int32(31))%32)&v315+v320) == v532 {
		goto L79
	} else {
		goto L80
	}
L78:
	;
	if v691 < int32(0) {
		goto L47
	} else {
		goto L106
	}
L79:
	;
	v537 = v315
	v538 = v533
	v545 = v532
	goto L82
L80:
	;
	v608 = v315
	v609 = v533
	goto L81
L81:
	;
	v644 = v608 << (uint(int32(2)) % 32)
	v646 = *(*int32)(unsafe.Add(mBase, uint32(v83+v644)))
	if v531 != v646 {
		goto L101
	} else {
		goto L102
	}
L82:
	;
	v573 = v537 << (uint(int32(2)) % 32)
	v575 = *(*int32)(unsafe.Add(mBase, uint32(v83+v573)))
	if v531 != v575 {
		goto L86
	} else {
		goto L87
	}
L83:
	;
	if v321 == int32(0) {
		v691 = v599
		goto L78
	} else {
		goto L99
	}
L84:
	;
	v586 = v537 - int32(1)
	v588 = v586 << (uint(int32(2)) % 32)
	v590 = *(*int32)(unsafe.Add(mBase, uint32(v83+v588)))
	if v531 != v590 {
		goto L93
	} else {
		goto L94
	}
L85:
	;
	v584 = v538
	goto L84
L86:
	;
	v577 = *(*int32)(unsafe.Add(mBase, uint32(v575)+364))
	if v577 != v531 {
		goto L85
	} else {
		goto L89
	}
L87:
	;
	goto L88
L88:
	;
	if v538 == int32(-1) {
		v584 = v537
		goto L84
	} else {
		goto L90
	}
L89:
	;
	goto L88
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v573+v75))) = int32(-1)
	goto L85
L91:
	;
	v600 = int32(2)
	v601 = v537 - v600
	v603 = v545 + v600
	if v603 != v319 {
		v537 = v601
		v538 = v599
		v545 = v603
		goto L82
	} else {
		goto L98
	}
L92:
	;
	v599 = v584
	goto L91
L93:
	;
	v592 = *(*int32)(unsafe.Add(mBase, uint32(v590)+364))
	if v592 != v531 {
		goto L92
	} else {
		goto L96
	}
L94:
	;
	goto L95
L95:
	;
	if v584 == int32(-1) {
		v599 = v586
		goto L91
	} else {
		goto L97
	}
L96:
	;
	goto L95
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v75+v588))) = int32(-1)
	goto L92
L98:
	;
	goto L83
L99:
	;
	v608 = v601
	v609 = v599
	goto L81
L100:
	;
	v691 = v609
	goto L78
L101:
	;
	v648 = *(*int32)(unsafe.Add(mBase, uint32(v646)+364))
	if v648 != v531 {
		goto L100
	} else {
		goto L104
	}
L102:
	;
	goto L103
L103:
	;
	if v609 == int32(-1) {
		v691 = v608
		goto L78
	} else {
		goto L105
	}
L104:
	;
	goto L103
L105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v644+v75))) = int32(-1)
	goto L100
L106:
	;
	v694 = int32(2)
	v696 = v75 + v528<<(uint(v694)%32)
	v697 = *(*int32)(unsafe.Add(mBase, uint32(v696)))
	v698 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v696))) = v697 + v698
	*(*int32)(unsafe.Add(mBase, uint32(v369)+12)) = v528
	v704 = v67 + v691<<(uint(v694)%32)
	v705 = *(*int32)(unsafe.Add(mBase, uint32(v704)))
	*(*int32)(unsafe.Add(mBase, uint32(v369)+16)) = v705
	*(*int32)(unsafe.Add(mBase, uint32(v704))) = v347 + v698
	goto L47
L107:
	;
	goto L46
L108:
	;
	v790 = v315
	v800 = v315
	goto L111
L109:
	;
	goto L110
L110:
	;
	v1121 = v109 + int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_TestConfiguration[4])) = v1121
	v1144 = v1121
	v1151 = v181 + v116
	goto L9
L111:
	;
	v826 = v790 + int32(1)
	v827 = v790
	goto L113
L112:
	;
	goto L110
L113:
	;
	v861 = int32(1)
	v868 = *(*int32)(unsafe.Add(mBase, uint32(v83+v827<<(uint(int32(2))%32))))
	if v868 == int32(0) {
		v826 = v826 - v861
		v827 = v827 - v861
		goto L113
	} else {
		goto L115
	}
L114:
	;
	if v827 < int32(0) {
		goto L116
	} else {
		goto L117
	}
L115:
	;
	goto L114
L116:
	;
	v1600 = int32(-1)
	v1603 = v39
	goto L1
L117:
	;
	goto L118
L118:
	;
	v875 = v827
	goto L119
L119:
	;
	v911 = v875 << (uint(int32(2)) % 32)
	v913 = *(*int32)(unsafe.Add(mBase, uint32(v83+v911)))
	if v913 != 0 {
		goto L122
	} else {
		goto L123
	}
L120:
	;
	v923 = *(*int32)(unsafe.Add(mBase, uint32(v913)+364))
	if v923 != 0 {
		goto L127
	} else {
		goto L128
	}
L121:
	;
	goto L120
L122:
	;
	v915 = *(*int32)(unsafe.Add(mBase, uint32(v75+v911)))
	if v915 == int32(0) {
		goto L121
	} else {
		goto L125
	}
L123:
	;
	goto L124
L124:
	;
	if int32(0) < v875 {
		v875 = v875 - int32(1)
		goto L119
	} else {
		goto L126
	}
L125:
	;
	goto L124
L126:
	;
	v1600 = int32(-1)
	v1603 = v39
	goto L1
L127:
	;
	v924 = v923
	goto L129
L128:
	;
	v924 = v913
	goto L129
L129:
	;
	v925 = int32(0)
	v928 = v925
	v936 = v925
	goto L130
L130:
	;
	v965 = v83 + v928<<(uint(int32(2))%32)
	v966 = *(*int32)(unsafe.Add(mBase, uint32(v965)))
	if v924 != v966 {
		goto L133
	} else {
		goto L134
	}
L131:
	;
	v986 = *(*int32)(unsafe.Add(mBase, uint32(v67+v911)))
	if int32(0) < v986 {
		goto L139
	} else {
		goto L140
	}
L132:
	;
	v983 = v928 + int32(1)
	if v983 != v826 {
		v928 = v983
		v936 = v981
		goto L130
	} else {
		goto L138
	}
L133:
	;
	if v966 == int32(0) {
		v981 = v936
		goto L132
	} else {
		goto L136
	}
L134:
	;
	goto L135
L135:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v178+(v800-v936)<<(uint(int32(2))%32)))) = v966
	*(*int32)(unsafe.Add(mBase, uint32(v965))) = int32(0)
	v981 = v936 + int32(1)
	goto L132
L136:
	;
	v970 = *(*int32)(unsafe.Add(mBase, uint32(v966)+364))
	if v970 != v924 {
		v981 = v936
		goto L132
	} else {
		goto L137
	}
L137:
	;
	goto L135
L138:
	;
	goto L131
L139:
	;
	v990 = v986
	goto L142
L140:
	;
	goto L141
L141:
	;
	v1079 = v800 - v981
	if int32(0) <= v1079 {
		v790 = v827
		v800 = v1079
		goto L111
	} else {
		goto L145
	}
L142:
	;
	v1027 = v65 + v990*int32(20)
	v1030 = *(*int32)(unsafe.Add(mBase, uint32(v1027-int32(8))))
	v1033 = v75 + v1030<<(uint(int32(2))%32)
	v1034 = *(*int32)(unsafe.Add(mBase, uint32(v1033)))
	*(*int32)(unsafe.Add(mBase, uint32(v1033))) = v1034 - int32(1)
	v1040 = *(*int32)(unsafe.Add(mBase, uint32(v1027-int32(4))))
	if int32(0) < v1040 {
		v990 = v1040
		goto L142
	} else {
		goto L144
	}
L143:
	;
	goto L141
L144:
	;
	goto L143
L145:
	;
	goto L112
L146:
	;
	goto L8
L147:
	;
	v1168 = int32(0)
	v1171 = v1163
	goto L148
L148:
	;
	v1204 = v1168 * int32(20)
	v1206 = *(*int32)(unsafe.Add(mBase, _c_F_TestConfiguration[8]))
	v1208 = *(*int32)(unsafe.Add(mBase, uint32(v1204+v1206)))
	v1210 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_TestConfiguration[7])) = v1210
	*(*int32)(unsafe.Add(mBase, _c_F_TestConfiguration[6])) = v1210
	*(*int32)(unsafe.Add(mBase, uint32(v39)+12)) = v1210
	v1219 = v39 + int32(12)
	v1223 = *(*int32)(unsafe.Add(mBase, uint32(v1208)+364))
	if v1223 != 0 {
		goto L152
	} else {
		goto L153
	}
L149:
	;
	v1450 = v1440
	v1453 = v39
	goto L5
L150:
	;
	v1324 = *(*int32)(unsafe.Add(mBase, _c_F_TestConfiguration[8]))
	v1326 = *(*int32)(unsafe.Add(mBase, uint32(v1324+v1204)+4))
	v1328 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_TestConfiguration[7])) = v1328
	*(*int32)(unsafe.Add(mBase, _c_F_TestConfiguration[6])) = v1328
	*(*int32)(unsafe.Add(mBase, uint32(v39)+12)) = v1328
	v1337 = v39 + int32(12)
	v1341 = *(*int32)(unsafe.Add(mBase, uint32(v1326)+364))
	if v1341 != 0 {
		goto L185
	} else {
		goto L186
	}
L151:
	;
	if v1317 == int32(0) {
		v1322 = v1171
		goto L150
	} else {
		goto L181
	}
L152:
	;
	v1224 = v1223
	goto L154
L153:
	;
	v1224 = v1208
	goto L154
L154:
	;
	v1225 = int32(0)
	v1227 = *(*int32)(unsafe.Add(mBase, _c_F_TestConfiguration[14]))
	v1229 = *(*int32)(unsafe.Add(mBase, _c_F_TestConfiguration[7]))
	if v1225 < v1229 {
		goto L155
	} else {
		goto L156
	}
L155:
	;
	v1232 = v1225
	goto L158
L156:
	;
	goto L157
L157:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_TestConfiguration[7])) = v1229 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1227+v1229<<(uint(int32(2))%32)))) = v1224
	v1266 = *(*int32)(unsafe.Add(mBase, uint32(v1224)+392))
	if v1266 == int32(0) {
		goto L167
	} else {
		goto L168
	}
L158:
	;
	v1242 = *(*int32)(unsafe.Add(mBase, uint32(v1227+v1232<<(uint(int32(2))%32))))
	if v1224 == v1242 {
		goto L160
	} else {
		goto L161
	}
L159:
	;
	goto L157
L160:
	;
	if v1232 != 0 {
		goto L163
	} else {
		goto L164
	}
L161:
	;
	goto L162
L162:
	;
	v1249 = v1232 + int32(1)
	if v1249 != v1229 {
		v1232 = v1249
		goto L158
	} else {
		goto L166
	}
L163:
	;
	v1317 = int32(0)
	goto L151
L164:
	;
	goto L165
L165:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_TestConfiguration[6])) = v1210
	v1317 = int32(1)
	goto L151
L166:
	;
	goto L159
L167:
	;
	v1273 = *(*int32)(unsafe.Add(mBase, uint32(v1224)+372))
	if v1273 == int32(0) {
		goto L170
	} else {
		goto L171
	}
L168:
	;
	v1269 = F_FindLockCycleRecurseMember(m, v1224, v1224, v1210, v54, v1219)
	mBase = m.M
	if v1269 == int32(0) {
		goto L167
	} else {
		goto L169
	}
L169:
	;
	v1317 = int32(1)
	goto L151
L170:
	;
	v1317 = int32(0)
	goto L151
L171:
	;
	v1277 = v1224 + int32(368)
	if v1273 == v1277 {
		goto L170
	} else {
		goto L172
	}
L172:
	;
	v1279 = v1273
	goto L173
L173:
	;
	v1286 = *(*int32)(unsafe.Add(mBase, uint32(v1279)+16))
	if v1286 == int32(0) {
		goto L175
	} else {
		goto L176
	}
L174:
	;
	goto L170
L175:
	;
	v1300 = *(*int32)(unsafe.Add(mBase, uint32(v1279)+4))
	if v1300 != v1277 {
		v1279 = v1300
		goto L173
	} else {
		goto L180
	}
L176:
	;
	v1289 = *(*int32)(unsafe.Add(mBase, uint32(v1279)+8))
	if v1289 == int32(0) {
		goto L175
	} else {
		goto L177
	}
L177:
	;
	v1293 = v1279 - int32(376)
	if v1293 == v1224 {
		goto L175
	} else {
		goto L178
	}
L178:
	;
	v1295 = F_FindLockCycleRecurseMember(m, v1293, v1224, v1210, v54, v1219)
	mBase = m.M
	if v1295 == int32(0) {
		goto L175
	} else {
		goto L179
	}
L179:
	;
	v1317 = int32(1)
	goto L151
L180:
	;
	goto L174
L181:
	;
	v1320 = *(*int32)(unsafe.Add(mBase, uint32(v39)+12))
	if v1320 != 0 {
		v1322 = v1320
		goto L150
	} else {
		goto L182
	}
L182:
	;
	v1600 = int32(-1)
	v1603 = v39
	goto L1
L183:
	;
	v1442 = v1168 + int32(1)
	v1444 = *(*int32)(unsafe.Add(mBase, _c_F_TestConfiguration[5]))
	if v1442 < v1444 {
		v1168 = v1442
		v1171 = v1440
		goto L148
	} else {
		goto L216
	}
L184:
	;
	if v1435 == int32(0) {
		v1440 = v1322
		goto L183
	} else {
		goto L214
	}
L185:
	;
	v1342 = v1341
	goto L187
L186:
	;
	v1342 = v1326
	goto L187
L187:
	;
	v1343 = int32(0)
	v1345 = *(*int32)(unsafe.Add(mBase, _c_F_TestConfiguration[14]))
	v1347 = *(*int32)(unsafe.Add(mBase, _c_F_TestConfiguration[7]))
	if v1343 < v1347 {
		goto L188
	} else {
		goto L189
	}
L188:
	;
	v1350 = v1343
	goto L191
L189:
	;
	goto L190
L190:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_TestConfiguration[7])) = v1347 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1345+v1347<<(uint(int32(2))%32)))) = v1342
	v1384 = *(*int32)(unsafe.Add(mBase, uint32(v1342)+392))
	if v1384 == int32(0) {
		goto L200
	} else {
		goto L201
	}
L191:
	;
	v1360 = *(*int32)(unsafe.Add(mBase, uint32(v1345+v1350<<(uint(int32(2))%32))))
	if v1342 == v1360 {
		goto L193
	} else {
		goto L194
	}
L192:
	;
	goto L190
L193:
	;
	if v1350 != 0 {
		goto L196
	} else {
		goto L197
	}
L194:
	;
	goto L195
L195:
	;
	v1367 = v1350 + int32(1)
	if v1367 != v1347 {
		v1350 = v1367
		goto L191
	} else {
		goto L199
	}
L196:
	;
	v1435 = int32(0)
	goto L184
L197:
	;
	goto L198
L198:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_TestConfiguration[6])) = v1328
	v1435 = int32(1)
	goto L184
L199:
	;
	goto L192
L200:
	;
	v1391 = *(*int32)(unsafe.Add(mBase, uint32(v1342)+372))
	if v1391 == int32(0) {
		goto L203
	} else {
		goto L204
	}
L201:
	;
	v1387 = F_FindLockCycleRecurseMember(m, v1342, v1342, v1328, v54, v1337)
	mBase = m.M
	if v1387 == int32(0) {
		goto L200
	} else {
		goto L202
	}
L202:
	;
	v1435 = int32(1)
	goto L184
L203:
	;
	v1435 = int32(0)
	goto L184
L204:
	;
	v1395 = v1342 + int32(368)
	if v1391 == v1395 {
		goto L203
	} else {
		goto L205
	}
L205:
	;
	v1397 = v1391
	goto L206
L206:
	;
	v1404 = *(*int32)(unsafe.Add(mBase, uint32(v1397)+16))
	if v1404 == int32(0) {
		goto L208
	} else {
		goto L209
	}
L207:
	;
	goto L203
L208:
	;
	v1418 = *(*int32)(unsafe.Add(mBase, uint32(v1397)+4))
	if v1418 != v1395 {
		v1397 = v1418
		goto L206
	} else {
		goto L213
	}
L209:
	;
	v1407 = *(*int32)(unsafe.Add(mBase, uint32(v1397)+8))
	if v1407 == int32(0) {
		goto L208
	} else {
		goto L210
	}
L210:
	;
	v1411 = v1397 - int32(376)
	if v1411 == v1342 {
		goto L208
	} else {
		goto L211
	}
L211:
	;
	v1413 = F_FindLockCycleRecurseMember(m, v1411, v1342, v1328, v54, v1337)
	mBase = m.M
	if v1413 == int32(0) {
		goto L208
	} else {
		goto L212
	}
L212:
	;
	v1435 = int32(1)
	goto L184
L213:
	;
	goto L207
L214:
	;
	v1438 = *(*int32)(unsafe.Add(mBase, uint32(v39)+12))
	if v1438 != 0 {
		v1440 = v1438
		goto L183
	} else {
		goto L215
	}
L215:
	;
	v1600 = int32(-1)
	v1603 = v39
	goto L1
L216:
	;
	goto L149
L217:
	;
	if v1590 == int32(0) {
		v1600 = v1450
		v1603 = v1453
		goto L1
	} else {
		goto L247
	}
L218:
	;
	v1497 = v1496
	goto L220
L219:
	;
	v1497 = l0
	goto L220
L220:
	;
	v1498 = int32(0)
	v1500 = *(*int32)(unsafe.Add(mBase, _c_F_TestConfiguration[14]))
	v1502 = *(*int32)(unsafe.Add(mBase, _c_F_TestConfiguration[7]))
	if v1498 < v1502 {
		goto L221
	} else {
		goto L222
	}
L221:
	;
	v1505 = v1498
	goto L224
L222:
	;
	goto L223
L223:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_TestConfiguration[7])) = v1502 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1500+v1502<<(uint(int32(2))%32)))) = v1497
	v1539 = *(*int32)(unsafe.Add(mBase, uint32(v1497)+392))
	if v1539 == int32(0) {
		goto L233
	} else {
		goto L234
	}
L224:
	;
	v1515 = *(*int32)(unsafe.Add(mBase, uint32(v1500+v1505<<(uint(int32(2))%32))))
	if v1497 == v1515 {
		goto L226
	} else {
		goto L227
	}
L225:
	;
	goto L223
L226:
	;
	if v1505 != 0 {
		goto L229
	} else {
		goto L230
	}
L227:
	;
	goto L228
L228:
	;
	v1522 = v1505 + int32(1)
	if v1522 != v1502 {
		v1505 = v1522
		goto L224
	} else {
		goto L232
	}
L229:
	;
	v1590 = int32(0)
	goto L217
L230:
	;
	goto L231
L231:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_TestConfiguration[6])) = v1483
	v1590 = int32(1)
	goto L217
L232:
	;
	goto L225
L233:
	;
	v1546 = *(*int32)(unsafe.Add(mBase, uint32(v1497)+372))
	if v1546 == int32(0) {
		goto L236
	} else {
		goto L237
	}
L234:
	;
	v1542 = F_FindLockCycleRecurseMember(m, v1497, v1497, v1483, v54, v1492)
	mBase = m.M
	if v1542 == int32(0) {
		goto L233
	} else {
		goto L235
	}
L235:
	;
	v1590 = int32(1)
	goto L217
L236:
	;
	v1590 = int32(0)
	goto L217
L237:
	;
	v1550 = v1497 + int32(368)
	if v1546 == v1550 {
		goto L236
	} else {
		goto L238
	}
L238:
	;
	v1552 = v1546
	goto L239
L239:
	;
	v1559 = *(*int32)(unsafe.Add(mBase, uint32(v1552)+16))
	if v1559 == int32(0) {
		goto L241
	} else {
		goto L242
	}
L240:
	;
	goto L236
L241:
	;
	v1573 = *(*int32)(unsafe.Add(mBase, uint32(v1552)+4))
	if v1573 != v1550 {
		v1552 = v1573
		goto L239
	} else {
		goto L246
	}
L242:
	;
	v1562 = *(*int32)(unsafe.Add(mBase, uint32(v1552)+8))
	if v1562 == int32(0) {
		goto L241
	} else {
		goto L243
	}
L243:
	;
	v1566 = v1552 - int32(376)
	if v1566 == v1497 {
		goto L241
	} else {
		goto L244
	}
L244:
	;
	v1568 = F_FindLockCycleRecurseMember(m, v1566, v1497, v1483, v54, v1492)
	mBase = m.M
	if v1568 == int32(0) {
		goto L241
	} else {
		goto L245
	}
L245:
	;
	v1590 = int32(1)
	goto L217
L246:
	;
	goto L240
L247:
	;
	v1593 = *(*int32)(unsafe.Add(mBase, uint32(v1453)+12))
	if v1593 != 0 {
		goto L248
	} else {
		goto L249
	}
L248:
	;
	v1595 = v1593
	goto L250
L249:
	;
	v1595 = int32(-1)
	goto L250
L250:
	;
	v1600 = v1595
	v1603 = v1453
	goto L1
}
func F_TransformPubWhereClauses(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
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
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
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
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	v4 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(32)
	m.G0 = v11
	if l0 == v4 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L14
	} else {
		goto L25
	}
L2:
	;
	m.G0 = v11 + int32(32)
	return
L3:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v15 <= int32(0) {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v25 = v4
	goto L5
L5:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v26+v25<<(uint(int32(2))%32))))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	if v31 != 0 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	goto L2
L7:
	;
	if l2 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	goto L9
L9:
	;
	v76 = v25 + int32(1)
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v76 < v77 {
		v25 = v76
		goto L5
	} else {
		goto L24
	}
L10:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)+48))
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+119)))
	if v36 == int32(112) {
		goto L1
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	v40 = F_make_parsestate(m, int32(0))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	goto L12
L14:
	;
	return
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+4)) = l1
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	v45 = int32(0)
	v48 = F_addRangeTableEntryForRelation(m, v40, v43, int32(1), v45, v45, v45)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v51 = int32(1)
	F_addNSItemToQuery(m, v40, v48, int32(0), v51, v51)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L14
	} else {
		goto L17
	}
L17:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	v56 = F_copyObjectImpl(m, v55)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L14
	} else {
		goto L18
	}
L18:
	;
	v60 = F_transformWhereClause(m, v40, v56, int32(6), int32(_a_F_TransformPubWhereClauses_0))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L14
	} else {
		goto L19
	}
L19:
	;
	F_assign_expr_collations(m, v40, v60)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L14
	} else {
		goto L20
	}
L20:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	v66 = F_expand_generated_columns_in_expr(m, v60, v64, int32(1))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L14
	} else {
		goto L21
	}
L21:
	;
	v68 = F_check_simple_rowfilter_expr_walker(m, v66, v40)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L14
	} else {
		goto L22
	}
L22:
	;
	F_free_parsestate(m, v40)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L14
	} else {
		goto L23
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+4)) = v66
	goto L9
L24:
	;
	goto L6
L25:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L14
	} else {
		goto L26
	}
L26:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v97)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v98 + int32(4)
	F_errmsg(m, int32(_a_F_TransformPubWhereClauses_1), v11+int32(16))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L14
	} else {
		goto L27
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = int32(_a_F_TransformPubWhereClauses_2)
	v110 = F_errdetail(m, int32(_a_F_TransformPubWhereClauses_3), v11)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L14
	} else {
		goto L28
	}
L28:
	;
	F_errfinish(m, int32(_a_F_TransformPubWhereClauses_4), int32(738), int32(_a_F_TransformPubWhereClauses_5))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L14
	} else {
		goto L29
	}
L29:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_TruncateSUBTRANS(m *base.Module, l0 int32) {
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v12 int32
	_ = v12
	v2 = l0
	for {
		v4 = v2 - int32(1)
		if base.Ui32(v4) < base.Ui32(int32(3)) {
			v2 = v4
			continue
		} else {
			break
		}
		break
	}
	F_SimpleLruTruncate(m, int32(_a_F_TruncateSUBTRANS_0), base.I64_extend_i32_u(int32(base.Ui32(v4)>>(uint(int32(11))%32))))
	v12 = m.ExcPending
	if v12 != 0 {
		return
	} else {
		return
	}
}
func F___towrite(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+72)) = v3 - int32(1) | v3
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v8&int32(8) != 0 {
		*(*int32)(unsafe.Add(mBase, uint32(l0))) = v8 | int32(32)
		return int32(-1)
	} else {
		*(*int64)(unsafe.Add(mBase, uint32(l0)+4)) = int64(0)
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v18
		*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v18
		v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v18 + v21
		return int32(0)
	}
}
func F_textgename(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int64
	_ = v4
	var v5 int64
	_ = v5
	var v6 int64
	_ = v6
	var v9 int32
	_ = v9
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v4 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v5 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v6 = F_DirectFunctionCall2Coll(m, int32(1756), v3, v4, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int64(0)
	} else {
		return int64(base.Ui64(v6^int64(-1))>>(uint(int64(31))%64)) & int64(1)
	}
}
func F_textlike_support(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v4 = F_like_regex_support(m, v2, int32(0))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		return base.I64_extend_i32_u(v4)
	}
}
func F_textregexeq(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14378(m, l0, int32(19))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return v3
	}
}
func F_textregexeq_support(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v4 = F_like_regex_support(m, v2, int32(2))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		return base.I64_extend_i32_u(v4)
	}
}
func F_textregexreplace(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
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
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v46 int32
	_ = v46
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v15 = F_pg_detoast_datum_packed(m, v14)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int64(0)
	} else {
		v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v20 = F_pg_detoast_datum_packed(m, v19)
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return int64(0)
		} else {
			v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
			v23 = F_pg_detoast_datum_packed(m, v22)
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int64(0)
			} else {
				v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
				v26 = F_pg_detoast_datum_packed(m, v25)
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return int64(0)
				} else {
					v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26))))
					if v28 == int32(1) {
						v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+1)))
						if base.B2i32(base.Ui32(int32(18)) < base.Ui32(v31))|base.B2i32(int32(1)<<(uint(v31)%32)&int32(_a_F_textregexreplace_0) == int32(0)) != 0 {
							F_parse_re_flags(m, v12+int32(8), v26)
							mBase = m.M
							v129 = m.ExcPending
							if v129 != 0 {
								return int64(0)
							} else {
								v130 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
								v131 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
								v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+12)))
								v136 = F_replace_text_regexp(m, v15, v20, v23, v130, v131, int32(0), v133^int32(1))
								mBase = m.M
								v137 = m.ExcPending
								if v137 != 0 {
									return int64(0)
								} else {
									m.G0 = v12 + int32(16)
									return base.I64_extend_i32_u(v136)
								}
							}
						} else {
							if v31 == int32(18) {
								v46 = int32(16)
							} else {
								v46 = int32(0)
							}
							if base.Ui32((v31-int32(1))&int32(255)) < base.Ui32(int32(3)) {
								v53 = int32(4)
							} else {
								v53 = v46
							}
							v89 = v53
							v91 = v26 + int32(1)
							v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91))))
							if base.Ui32(int32(9)) < base.Ui32((v92-int32(48))&int32(255)) {
								F_parse_re_flags(m, v12+int32(8), v26)
								mBase = m.M
								v129 = m.ExcPending
								if v129 != 0 {
									return int64(0)
								} else {
									v130 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
									v131 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
									v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+12)))
									v136 = F_replace_text_regexp(m, v15, v20, v23, v130, v131, int32(0), v133^int32(1))
									mBase = m.M
									v137 = m.ExcPending
									if v137 != 0 {
										return int64(0)
									} else {
										m.G0 = v12 + int32(16)
										return base.I64_extend_i32_u(v136)
									}
								}
							} else {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v102 = m.ExcPending
								if v102 != 0 {
									return int64(0)
								} else {
									F_errcode(m, int32(50856066))
									mBase = m.M
									v105 = m.ExcPending
									if v105 != 0 {
										return int64(0)
									} else {
										v107 = F_pg_mblen_range(m, v91, v91+v89)
										mBase = m.M
										v108 = m.ExcPending
										if v108 != 0 {
											return int64(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v91
											*(*int32)(unsafe.Add(mBase, uint32(v12))) = v107
											F_errmsg(m, int32(_a_F_textregexreplace_1), v12)
											mBase = m.M
											v113 = m.ExcPending
											if v113 != 0 {
												return int64(0)
											} else {
												F_errhint(m, int32(_a_F_textregexreplace_2), int32(0))
												mBase = m.M
												v117 = m.ExcPending
												if v117 != 0 {
													return int64(0)
												} else {
													F_errfinish(m, int32(_a_F_textregexreplace_3), int32(683), int32(_a_F_textregexreplace_4))
													mBase = m.M
													v122 = m.ExcPending
													if v122 != 0 {
														return int64(0)
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
						v57 = v28 & int32(1)
						if v57 == int32(0) {
							v60 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
							if v60&int32(-4) == int32(16) {
								F_parse_re_flags(m, v12+int32(8), v26)
								mBase = m.M
								v129 = m.ExcPending
								if v129 != 0 {
									return int64(0)
								} else {
									v130 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
									v131 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
									v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+12)))
									v136 = F_replace_text_regexp(m, v15, v20, v23, v130, v131, int32(0), v133^int32(1))
									mBase = m.M
									v137 = m.ExcPending
									if v137 != 0 {
										return int64(0)
									} else {
										m.G0 = v12 + int32(16)
										return base.I64_extend_i32_u(v136)
									}
								}
							} else {
								v65 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
								v68 = int32(4)
								if v57 != 0 {
									v72 = int32(1)
								} else {
									v72 = v68
								}
								v89 = int32(base.Ui32(v65)>>(uint(int32(2))%32)) - v68
								v91 = v26 + v72
								v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91))))
								if base.Ui32(int32(9)) < base.Ui32((v92-int32(48))&int32(255)) {
									F_parse_re_flags(m, v12+int32(8), v26)
									mBase = m.M
									v129 = m.ExcPending
									if v129 != 0 {
										return int64(0)
									} else {
										v130 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
										v131 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
										v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+12)))
										v136 = F_replace_text_regexp(m, v15, v20, v23, v130, v131, int32(0), v133^int32(1))
										mBase = m.M
										v137 = m.ExcPending
										if v137 != 0 {
											return int64(0)
										} else {
											m.G0 = v12 + int32(16)
											return base.I64_extend_i32_u(v136)
										}
									}
								} else {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v102 = m.ExcPending
									if v102 != 0 {
										return int64(0)
									} else {
										F_errcode(m, int32(50856066))
										mBase = m.M
										v105 = m.ExcPending
										if v105 != 0 {
											return int64(0)
										} else {
											v107 = F_pg_mblen_range(m, v91, v91+v89)
											mBase = m.M
											v108 = m.ExcPending
											if v108 != 0 {
												return int64(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v91
												*(*int32)(unsafe.Add(mBase, uint32(v12))) = v107
												F_errmsg(m, int32(_a_F_textregexreplace_1), v12)
												mBase = m.M
												v113 = m.ExcPending
												if v113 != 0 {
													return int64(0)
												} else {
													F_errhint(m, int32(_a_F_textregexreplace_2), int32(0))
													mBase = m.M
													v117 = m.ExcPending
													if v117 != 0 {
														return int64(0)
													} else {
														F_errfinish(m, int32(_a_F_textregexreplace_3), int32(683), int32(_a_F_textregexreplace_4))
														mBase = m.M
														v122 = m.ExcPending
														if v122 != 0 {
															return int64(0)
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
							if v28&int32(254) == int32(2) {
								F_parse_re_flags(m, v12+int32(8), v26)
								mBase = m.M
								v129 = m.ExcPending
								if v129 != 0 {
									return int64(0)
								} else {
									v130 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
									v131 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
									v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+12)))
									v136 = F_replace_text_regexp(m, v15, v20, v23, v130, v131, int32(0), v133^int32(1))
									mBase = m.M
									v137 = m.ExcPending
									if v137 != 0 {
										return int64(0)
									} else {
										m.G0 = v12 + int32(16)
										return base.I64_extend_i32_u(v136)
									}
								}
							} else {
								v78 = int32(1)
								if v28&v78 != 0 {
									v86 = v78
								} else {
									v86 = int32(4)
								}
								v89 = int32(base.Ui32(v28)>>(uint(v78)%32)) - v78
								v91 = v26 + v86
								v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91))))
								if base.Ui32(int32(9)) < base.Ui32((v92-int32(48))&int32(255)) {
									F_parse_re_flags(m, v12+int32(8), v26)
									mBase = m.M
									v129 = m.ExcPending
									if v129 != 0 {
										return int64(0)
									} else {
										v130 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
										v131 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
										v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+12)))
										v136 = F_replace_text_regexp(m, v15, v20, v23, v130, v131, int32(0), v133^int32(1))
										mBase = m.M
										v137 = m.ExcPending
										if v137 != 0 {
											return int64(0)
										} else {
											m.G0 = v12 + int32(16)
											return base.I64_extend_i32_u(v136)
										}
									}
								} else {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v102 = m.ExcPending
									if v102 != 0 {
										return int64(0)
									} else {
										F_errcode(m, int32(50856066))
										mBase = m.M
										v105 = m.ExcPending
										if v105 != 0 {
											return int64(0)
										} else {
											v107 = F_pg_mblen_range(m, v91, v91+v89)
											mBase = m.M
											v108 = m.ExcPending
											if v108 != 0 {
												return int64(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v91
												*(*int32)(unsafe.Add(mBase, uint32(v12))) = v107
												F_errmsg(m, int32(_a_F_textregexreplace_1), v12)
												mBase = m.M
												v113 = m.ExcPending
												if v113 != 0 {
													return int64(0)
												} else {
													F_errhint(m, int32(_a_F_textregexreplace_2), int32(0))
													mBase = m.M
													v117 = m.ExcPending
													if v117 != 0 {
														return int64(0)
													} else {
														F_errfinish(m, int32(_a_F_textregexreplace_3), int32(683), int32(_a_F_textregexreplace_4))
														mBase = m.M
														v122 = m.ExcPending
														if v122 != 0 {
															return int64(0)
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
		}
	}
}
func F_textregexreplace_extended(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
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
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
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
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
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
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	v10 = m.G0
	v12 = v10 - int32(32)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v15 = F_pg_detoast_datum_packed(m, v14)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int64(0)
	} else {
		v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v20 = F_pg_detoast_datum_packed(m, v19)
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return int64(0)
		} else {
			v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
			v23 = F_pg_detoast_datum_packed(m, v22)
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int64(0)
			} else {
				v25 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
				if int32(6) <= v25 {
					v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
					v29 = F_pg_detoast_datum_packed(m, v28)
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
						return int64(0)
					} else {
						v31 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
						v32 = v31
						v33 = v29
						v34 = int32(1)
						if v32 < int32(4) {
							v48 = v34
							v49 = v34
							F_parse_re_flags(m, v12+int32(24), v33)
							mBase = m.M
							v53 = m.ExcPending
							if v53 != 0 {
								return int64(0)
							} else {
								v54 = *(*int32)(unsafe.Add(mBase, uint32(v12)+24))
								v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
								v56 = int32(1)
								v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+28)))
								v61 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
								if v61 < int32(5) {
									v64 = v58 ^ v56
								} else {
									v64 = v49
								}
								v65 = F_replace_text_regexp(m, v15, v20, v23, v54, v55, v48-v56, v64)
								mBase = m.M
								v66 = m.ExcPending
								if v66 != 0 {
									return int64(0)
								} else {
									m.G0 = v12 + int32(32)
									return base.I64_extend_i32_u(v65)
								}
							}
						} else {
							v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
							if v38 <= int32(0) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v75 = m.ExcPending
								if v75 != 0 {
									return int64(0)
								} else {
									F_errcode(m, int32(50856066))
									mBase = m.M
									v78 = m.ExcPending
									if v78 != 0 {
										return int64(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v38
										*(*int32)(unsafe.Add(mBase, uint32(v12))) = int32(_a_F_textregexreplace_extended_0)
										F_errmsg(m, int32(_a_F_textregexreplace_extended_1), v12)
										mBase = m.M
										v84 = m.ExcPending
										if v84 != 0 {
											return int64(0)
										} else {
											F_errfinish(m, int32(_a_F_textregexreplace_extended_2), int32(718), int32(_a_F_textregexreplace_extended_3))
											mBase = m.M
											v89 = m.ExcPending
											if v89 != 0 {
												return int64(0)
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									}
								}
							} else {
								if v32&int32(_a_F_textregexreplace_extended_4) == int32(4) {
									v48 = v38
									v49 = v34
									F_parse_re_flags(m, v12+int32(24), v33)
									mBase = m.M
									v53 = m.ExcPending
									if v53 != 0 {
										return int64(0)
									} else {
										v54 = *(*int32)(unsafe.Add(mBase, uint32(v12)+24))
										v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
										v56 = int32(1)
										v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+28)))
										v61 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
										if v61 < int32(5) {
											v64 = v58 ^ v56
										} else {
											v64 = v49
										}
										v65 = F_replace_text_regexp(m, v15, v20, v23, v54, v55, v48-v56, v64)
										mBase = m.M
										v66 = m.ExcPending
										if v66 != 0 {
											return int64(0)
										} else {
											m.G0 = v12 + int32(32)
											return base.I64_extend_i32_u(v65)
										}
									}
								} else {
									v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
									if v45 < int32(0) {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v93 = m.ExcPending
										if v93 != 0 {
											return int64(0)
										} else {
											F_errcode(m, int32(50856066))
											mBase = m.M
											v96 = m.ExcPending
											if v96 != 0 {
												return int64(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v45
												*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = int32(_a_F_textregexreplace_extended_5)
												F_errmsg(m, int32(_a_F_textregexreplace_extended_1), v12+int32(16))
												mBase = m.M
												v104 = m.ExcPending
												if v104 != 0 {
													return int64(0)
												} else {
													F_errfinish(m, int32(_a_F_textregexreplace_extended_2), int32(727), int32(_a_F_textregexreplace_extended_3))
													mBase = m.M
													v109 = m.ExcPending
													if v109 != 0 {
														return int64(0)
													} else {
														base.Wasm_trap_unreachable()
														for {
														}
													}
												}
											}
										}
									} else {
										v48 = v38
										v49 = v45
										F_parse_re_flags(m, v12+int32(24), v33)
										mBase = m.M
										v53 = m.ExcPending
										if v53 != 0 {
											return int64(0)
										} else {
											v54 = *(*int32)(unsafe.Add(mBase, uint32(v12)+24))
											v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
											v56 = int32(1)
											v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+28)))
											v61 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
											if v61 < int32(5) {
												v64 = v58 ^ v56
											} else {
												v64 = v49
											}
											v65 = F_replace_text_regexp(m, v15, v20, v23, v54, v55, v48-v56, v64)
											mBase = m.M
											v66 = m.ExcPending
											if v66 != 0 {
												return int64(0)
											} else {
												m.G0 = v12 + int32(32)
												return base.I64_extend_i32_u(v65)
											}
										}
									}
								}
							}
						}
					}
				} else {
					v32 = v25
					v33 = int32(0)
					v34 = int32(1)
					if v32 < int32(4) {
						v48 = v34
						v49 = v34
						F_parse_re_flags(m, v12+int32(24), v33)
						mBase = m.M
						v53 = m.ExcPending
						if v53 != 0 {
							return int64(0)
						} else {
							v54 = *(*int32)(unsafe.Add(mBase, uint32(v12)+24))
							v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							v56 = int32(1)
							v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+28)))
							v61 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
							if v61 < int32(5) {
								v64 = v58 ^ v56
							} else {
								v64 = v49
							}
							v65 = F_replace_text_regexp(m, v15, v20, v23, v54, v55, v48-v56, v64)
							mBase = m.M
							v66 = m.ExcPending
							if v66 != 0 {
								return int64(0)
							} else {
								m.G0 = v12 + int32(32)
								return base.I64_extend_i32_u(v65)
							}
						}
					} else {
						v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
						if v38 <= int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v75 = m.ExcPending
							if v75 != 0 {
								return int64(0)
							} else {
								F_errcode(m, int32(50856066))
								mBase = m.M
								v78 = m.ExcPending
								if v78 != 0 {
									return int64(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v38
									*(*int32)(unsafe.Add(mBase, uint32(v12))) = int32(_a_F_textregexreplace_extended_0)
									F_errmsg(m, int32(_a_F_textregexreplace_extended_1), v12)
									mBase = m.M
									v84 = m.ExcPending
									if v84 != 0 {
										return int64(0)
									} else {
										F_errfinish(m, int32(_a_F_textregexreplace_extended_2), int32(718), int32(_a_F_textregexreplace_extended_3))
										mBase = m.M
										v89 = m.ExcPending
										if v89 != 0 {
											return int64(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							}
						} else {
							if v32&int32(_a_F_textregexreplace_extended_4) == int32(4) {
								v48 = v38
								v49 = v34
								F_parse_re_flags(m, v12+int32(24), v33)
								mBase = m.M
								v53 = m.ExcPending
								if v53 != 0 {
									return int64(0)
								} else {
									v54 = *(*int32)(unsafe.Add(mBase, uint32(v12)+24))
									v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
									v56 = int32(1)
									v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+28)))
									v61 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
									if v61 < int32(5) {
										v64 = v58 ^ v56
									} else {
										v64 = v49
									}
									v65 = F_replace_text_regexp(m, v15, v20, v23, v54, v55, v48-v56, v64)
									mBase = m.M
									v66 = m.ExcPending
									if v66 != 0 {
										return int64(0)
									} else {
										m.G0 = v12 + int32(32)
										return base.I64_extend_i32_u(v65)
									}
								}
							} else {
								v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
								if v45 < int32(0) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v93 = m.ExcPending
									if v93 != 0 {
										return int64(0)
									} else {
										F_errcode(m, int32(50856066))
										mBase = m.M
										v96 = m.ExcPending
										if v96 != 0 {
											return int64(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v45
											*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = int32(_a_F_textregexreplace_extended_5)
											F_errmsg(m, int32(_a_F_textregexreplace_extended_1), v12+int32(16))
											mBase = m.M
											v104 = m.ExcPending
											if v104 != 0 {
												return int64(0)
											} else {
												F_errfinish(m, int32(_a_F_textregexreplace_extended_2), int32(727), int32(_a_F_textregexreplace_extended_3))
												mBase = m.M
												v109 = m.ExcPending
												if v109 != 0 {
													return int64(0)
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											}
										}
									}
								} else {
									v48 = v38
									v49 = v45
									F_parse_re_flags(m, v12+int32(24), v33)
									mBase = m.M
									v53 = m.ExcPending
									if v53 != 0 {
										return int64(0)
									} else {
										v54 = *(*int32)(unsafe.Add(mBase, uint32(v12)+24))
										v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
										v56 = int32(1)
										v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+28)))
										v61 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
										if v61 < int32(5) {
											v64 = v58 ^ v56
										} else {
											v64 = v49
										}
										v65 = F_replace_text_regexp(m, v15, v20, v23, v54, v55, v48-v56, v64)
										mBase = m.M
										v66 = m.ExcPending
										if v66 != 0 {
											return int64(0)
										} else {
											m.G0 = v12 + int32(32)
											return base.I64_extend_i32_u(v65)
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
func F_thesaurus_lexize(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v48 int64
	_ = v48
	var v49 int64
	_ = v49
	var v50 int64
	_ = v50
	var v52 int64
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
	var v159 int32
	_ = v159
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v205 int32
	_ = v205
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v232 int32
	_ = v232
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v253 int32
	_ = v253
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v289 int32
	_ = v289
	var v295 int32
	_ = v295
	var v304 int32
	_ = v304
	var v313 int32
	_ = v313
	var v321 int32
	_ = v321
	var v329 int32
	_ = v329
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v360 int32
	_ = v360
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v387 int32
	_ = v387
	var v403 int32
	_ = v403
	var v414 int32
	_ = v414
	var v416 int32
	_ = v416
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v440 int32
	_ = v440
	var v450 int32
	_ = v450
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v478 int32
	_ = v478
	var v481 int32
	_ = v481
	var v489 int32
	_ = v489
	var v497 int32
	_ = v497
	var v502 int32
	_ = v502
	var v506 int32
	_ = v506
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v518 int32
	_ = v518
	var v520 int32
	_ = v520
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v552 int32
	_ = v552
	var v566 int32
	_ = v566
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v593 int32
	_ = v593
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v605 int32
	_ = v605
	var v614 int32
	_ = v614
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v626 int32
	_ = v626
	var v628 int32
	_ = v628
	var v630 int32
	_ = v630
	var v650 int32
	_ = v650
	var v656 int32
	_ = v656
	var v665 int32
	_ = v665
	var v682 int32
	_ = v682
	var v690 int32
	_ = v690
	var v699 int32
	_ = v699
	var v701 int32
	_ = v701
	var v721 int32
	_ = v721
	var v730 int32
	_ = v730
	var v732 int32
	_ = v732
	var v748 int32
	_ = v748
	var v764 int32
	_ = v764
	var v825 int32
	_ = v825
	var v826 int32
	_ = v826
	var v835 int32
	_ = v835
	var v836 int32
	_ = v836
	var v837 int32
	_ = v837
	var v839 int32
	_ = v839
	var v842 int32
	_ = v842
	var v850 int32
	_ = v850
	var v873 int32
	_ = v873
	var v875 int32
	_ = v875
	var v877 int32
	_ = v877
	var v879 int32
	_ = v879
	var v880 int32
	_ = v880
	var v881 int32
	_ = v881
	var v884 int32
	_ = v884
	var v885 int32
	_ = v885
	var v889 int32
	_ = v889
	var v890 int32
	_ = v890
	var v902 int32
	_ = v902
	var v903 int32
	_ = v903
	var v905 int32
	_ = v905
	var v908 int32
	_ = v908
	var v909 int32
	_ = v909
	var v912 int32
	_ = v912
	var v914 int32
	_ = v914
	var v926 int32
	_ = v926
	var v928 int32
	_ = v928
	var v930 int32
	_ = v930
	var v933 int32
	_ = v933
	var v934 int32
	_ = v934
	var v935 int32
	_ = v935
	var v938 int32
	_ = v938
	var v952 int32
	_ = v952
	var v953 int32
	_ = v953
	var v954 int32
	_ = v954
	var v956 int64
	_ = v956
	var v958 int32
	_ = v958
	var v960 int32
	_ = v960
	var v961 int32
	_ = v961
	var v962 int32
	_ = v962
	var v965 int32
	_ = v965
	var v966 int32
	_ = v966
	var v985 int32
	_ = v985
	var v990 int32
	_ = v990
	var v1021 int64
	_ = v1021
	v15 = m.G0
	v17 = v15 - int32(16)
	m.G0 = v17
	v19 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+18)))
	if v19 != int32(4) {
		goto L7
	} else {
		goto L8
	}
L1:
	;
	m.G0 = v17 + int32(16)
	return v1021
L2:
	;
	v1021 = int64(0)
	goto L1
L3:
	;
	v875 = *(*int32)(unsafe.Add(mBase, uint32(v850)+12))
	v877 = base.B2i32(v875 == int32(0))
	v879 = v33 & int32(_a_F_thesaurus_lexize_0)
	v880 = *(*int32)(unsafe.Add(mBase, uint32(v27)+20))
	v881 = *(*int32)(unsafe.Add(mBase, uint32(v850)))
	v884 = v880 + v881<<(uint(int32(3))%32)
	v885 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v884))))
	if v879 == v885 {
		goto L173
	} else {
		goto L174
	}
L4:
	;
	v873 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+1)) = uint8(v873)
	goto L2
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+4)) = v850
	if v850 != 0 {
		goto L3
	} else {
		goto L171
	}
L6:
	;
	v512 = int32(0)
	v513 = *(*int32)(unsafe.Add(mBase, uint32(v27)+12))
	if v513 == v512 {
		v528 = v512
		goto L110
	} else {
		goto L111
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v502 = m.ExcPending
	if v502 != 0 {
		goto L17
	} else {
		goto L107
	}
L8:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	if v22 == int32(0) {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22))))
	if v26 != 0 {
		v1021 = int64(0)
		goto L1
	} else {
		goto L10
	}
L10:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	if v28 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v29 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v28)+4)))
	v33 = v29 + int32(1)
	goto L13
L12:
	;
	v33 = int32(0)
	goto L13
L13:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+4)))
	if v35 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
	v39 = F_lookup_ts_dictionary_cache(m, v38)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	v44 = v34
	goto L16
L16:
	;
	v48 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v44)+44)))
	v49 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v50 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
	v52 = F_FunctionCall4Coll(m, v44+int32(12), int32(0), v48, v49, v50, int64(0))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L17
	} else {
		goto L19
	}
L17:
	;
	return int64(0)
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+4)) = v39
	v44 = v39
	goto L16
L19:
	;
	v54 = base.I32_wrap_i64(v52)
	if v54 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+4)) = int32(0)
	goto L4
L21:
	;
	goto L22
L22:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v54)+4))
	if v59 == int32(0) {
		goto L6
	} else {
		goto L23
	}
L23:
	;
	v64 = v54
	v68 = int32(0)
	goto L24
L24:
	;
	v76 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v64))))
	v78 = v64
	v79 = int32(0)
	goto L26
L25:
	;
	v850 = v489
	goto L5
L26:
	;
	v92 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v78))))
	if v92 != v76 {
		goto L29
	} else {
		goto L30
	}
L27:
	;
	v105 = v100 & int32(_a_F_thesaurus_lexize_0)
	v106 = F_palloc_mul(m, int32(4), v105)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L17
	} else {
		goto L33
	}
L28:
	;
	goto L27
L29:
	;
	v100 = v79
	v101 = v78
	goto L28
L30:
	;
	goto L31
L31:
	;
	v95 = v79 + int32(1)
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v78)+12))
	v98 = v78 + int32(8)
	if v96 != 0 {
		v78 = v98
		v79 = v95
		goto L26
	} else {
		goto L32
	}
L32:
	;
	v100 = v95
	v101 = v98
	goto L28
L33:
	;
	if v105 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	v497 = *(*int32)(unsafe.Add(mBase, uint32(v101)+4))
	if v497 != 0 {
		v64 = v101
		v68 = v489
		goto L24
	} else {
		goto L106
	}
L35:
	;
	v188 = v105 & int32(3)
	v191 = v68
	goto L50
L36:
	;
	v110 = int32(0)
	goto L37
L37:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v27)+12))
	if v124 != 0 {
		goto L41
	} else {
		goto L42
	}
L38:
	;
	F_pfree(m, v106)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L17
	} else {
		goto L48
	}
L39:
	;
	goto L38
L40:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v137)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v106+v110<<(uint(int32(2))%32)))) = v149
	if v149 == int32(0) {
		goto L39
	} else {
		goto L46
	}
L41:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v64+v110<<(uint(int32(3))%32))+4))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+12)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+8)) = v128
	v132 = int32(8)
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v27)+8))
	v137 = F_bsearch(m, v17+v132, v134, v124, v132, int32(1265))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L17
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v106+v110<<(uint(int32(2))%32)))) = int32(0)
	goto L39
L44:
	;
	if v137 != 0 {
		goto L40
	} else {
		goto L45
	}
L45:
	;
	goto L43
L46:
	;
	v154 = v110 + int32(1)
	if v105 != v154 {
		v110 = v154
		goto L37
	} else {
		goto L47
	}
L47:
	;
	goto L35
L48:
	;
	v489 = v68
	goto L34
L49:
	;
	v489 = v191
	goto L34
L50:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v106)))
	if v105 != 0 {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	if v105 == int32(0) {
		v191 = v387
		goto L50
	} else {
		goto L95
	}
L53:
	;
	v213 = v205
	v214 = int32(0)
	goto L56
L54:
	;
	v313 = v205
	goto L55
L55:
	;
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v313)))
	if v28 == int32(0) {
		goto L82
	} else {
		goto L83
	}
L56:
	;
	v223 = v106 + v214<<(uint(int32(2))%32)
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v223)))
	if v224 == int32(0) {
		goto L60
	} else {
		goto L61
	}
L57:
	;
	if v105 != v304 {
		v387 = v191
		goto L52
	} else {
		goto L81
	}
L58:
	;
	if v304 < v105 {
		v213 = v295
		v214 = v304
		goto L56
	} else {
		goto L80
	}
L59:
	;
	if v262 == v263 {
		goto L77
	} else {
		goto L78
	}
L60:
	;
	goto L49
L61:
	;
	v232 = v224
	goto L62
L62:
	;
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v232)))
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v213)))
	if base.Ui32(v241) < base.Ui32(v242) {
		goto L64
	} else {
		goto L65
	}
L63:
	;
	if base.Ui32(v242) < base.Ui32(v241) {
		v295 = v232
		v304 = int32(0)
		goto L58
	} else {
		goto L68
	}
L64:
	;
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v232)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v223))) = v244
	if v244 != 0 {
		v232 = v244
		goto L62
	} else {
		goto L67
	}
L65:
	;
	goto L66
L66:
	;
	goto L63
L67:
	;
	goto L60
L68:
	;
	v253 = v232
	goto L69
L69:
	;
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v253)))
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v213)))
	if v262 != v263 {
		goto L59
	} else {
		goto L71
	}
L70:
	;
	goto L60
L71:
	;
	v265 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v253)+4)))
	if v33&int32(_a_F_thesaurus_lexize_0) == v265 {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v267 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v253)+6)))
	if v105 == v267 {
		goto L59
	} else {
		goto L75
	}
L73:
	;
	goto L74
L74:
	;
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v253)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v223))) = v269
	if v269 != 0 {
		v253 = v269
		goto L69
	} else {
		goto L76
	}
L75:
	;
	goto L74
L76:
	;
	goto L70
L77:
	;
	v289 = v214 + int32(1)
	goto L79
L78:
	;
	v289 = int32(0)
	goto L79
L79:
	;
	v295 = v253
	v304 = v289
	goto L58
L80:
	;
	goto L57
L81:
	;
	v313 = v295
	goto L55
L82:
	;
	if v191 != 0 {
		goto L88
	} else {
		goto L89
	}
L83:
	;
	v329 = v28
	goto L84
L84:
	;
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v329)))
	if v338 == v321 {
		goto L82
	} else {
		goto L86
	}
L85:
	;
	v387 = v191
	goto L52
L86:
	;
	v340 = *(*int32)(unsafe.Add(mBase, uint32(v329)+12))
	if v340 != 0 {
		v329 = v340
		goto L84
	} else {
		goto L87
	}
L87:
	;
	goto L85
L88:
	;
	v360 = v191
	goto L91
L89:
	;
	goto L90
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v313)+12)) = v191
	v387 = v313
	goto L52
L91:
	;
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v360)))
	if v369 == v321 {
		v387 = v191
		goto L52
	} else {
		goto L93
	}
L92:
	;
	goto L90
L93:
	;
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v360)+12))
	if v371 != 0 {
		v360 = v371
		goto L91
	} else {
		goto L94
	}
L94:
	;
	goto L92
L95:
	;
	v403 = int32(0)
	if base.B2i32(base.Ui32(v105) < base.Ui32(int32(4))) == v403 {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v414 = v403
	v416 = v403
	goto L99
L97:
	;
	v450 = v403
	goto L98
L98:
	;
	v464 = v450
	v465 = v403
	goto L103
L99:
	;
	v424 = v106 + v414<<(uint(int32(2))%32)
	v425 = *(*int32)(unsafe.Add(mBase, uint32(v424)))
	v426 = *(*int32)(unsafe.Add(mBase, uint32(v425)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v424))) = v426
	v428 = *(*int32)(unsafe.Add(mBase, uint32(v424)+4))
	v429 = *(*int32)(unsafe.Add(mBase, uint32(v428)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v424)+4)) = v429
	v431 = *(*int32)(unsafe.Add(mBase, uint32(v424)+8))
	v432 = *(*int32)(unsafe.Add(mBase, uint32(v431)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v424)+8)) = v432
	v434 = *(*int32)(unsafe.Add(mBase, uint32(v424)+12))
	v435 = *(*int32)(unsafe.Add(mBase, uint32(v434)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v424)+12)) = v435
	v437 = int32(4)
	v438 = v414 + v437
	v440 = v416 + v437
	if v440 != v105&int32(_a_F_thesaurus_lexize_1) {
		v414 = v438
		v416 = v440
		goto L99
	} else {
		goto L101
	}
L100:
	;
	if v188 == int32(0) {
		v191 = v387
		goto L50
	} else {
		goto L102
	}
L101:
	;
	goto L100
L102:
	;
	v450 = v438
	goto L98
L103:
	;
	v474 = v106 + v464<<(uint(int32(2))%32)
	v475 = *(*int32)(unsafe.Add(mBase, uint32(v474)))
	v476 = *(*int32)(unsafe.Add(mBase, uint32(v475)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v474))) = v476
	v478 = int32(1)
	v481 = v465 + v478
	if v481 != v188 {
		v464 = v464 + v478
		v465 = v481
		goto L103
	} else {
		goto L105
	}
L104:
	;
	v191 = v387
	goto L50
L105:
	;
	goto L104
L106:
	;
	goto L25
L107:
	;
	F_errmsg_internal(m, int32(_a_F_thesaurus_lexize_2), int32(0))
	mBase = m.M
	v506 = m.ExcPending
	if v506 != 0 {
		goto L17
	} else {
		goto L108
	}
L108:
	;
	F_errfinish(m, int32(_a_F_thesaurus_lexize_3), int32(800), int32(_a_F_thesaurus_lexize_4))
	mBase = m.M
	v511 = m.ExcPending
	if v511 != 0 {
		goto L17
	} else {
		goto L109
	}
L109:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+4)) = v528
	v535 = v17 + int32(4)
	v536 = int32(1)
	v552 = int32(0)
	goto L115
L111:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v17)+8)) = int64(0)
	v518 = int32(8)
	v520 = *(*int32)(unsafe.Add(mBase, uint32(v27)+8))
	v523 = F_bsearch(m, v17+v518, v520, v513, v518, int32(1265))
	mBase = m.M
	v524 = m.ExcPending
	if v524 != 0 {
		goto L17
	} else {
		goto L112
	}
L112:
	;
	if v523 == int32(0) {
		v528 = v512
		goto L110
	} else {
		goto L113
	}
L113:
	;
	v527 = *(*int32)(unsafe.Add(mBase, uint32(v523)+4))
	v528 = v527
	goto L110
L114:
	;
	v850 = v552
	goto L5
L115:
	;
	v566 = *(*int32)(unsafe.Add(mBase, uint32(v535)))
	goto L118
L117:
	;
	goto L160
L118:
	;
	v574 = v566
	v575 = int32(0)
	goto L121
L120:
	;
	v682 = *(*int32)(unsafe.Add(mBase, uint32(v656)))
	if v28 == int32(0) {
		goto L147
	} else {
		goto L148
	}
L121:
	;
	v584 = v535 + v575<<(uint(int32(2))%32)
	v585 = *(*int32)(unsafe.Add(mBase, uint32(v584)))
	if v585 == int32(0) {
		goto L125
	} else {
		goto L126
	}
L122:
	;
	if v536 != v665 {
		v748 = v552
		goto L117
	} else {
		goto L146
	}
L123:
	;
	if v665 < v536 {
		v574 = v656
		v575 = v665
		goto L121
	} else {
		goto L145
	}
L124:
	;
	if v623 == v624 {
		goto L142
	} else {
		goto L143
	}
L125:
	;
	goto L114
L126:
	;
	v593 = v585
	goto L127
L127:
	;
	v602 = *(*int32)(unsafe.Add(mBase, uint32(v593)))
	v603 = *(*int32)(unsafe.Add(mBase, uint32(v574)))
	if base.Ui32(v602) < base.Ui32(v603) {
		goto L129
	} else {
		goto L130
	}
L128:
	;
	if base.Ui32(v603) < base.Ui32(v602) {
		v656 = v593
		v665 = int32(0)
		goto L123
	} else {
		goto L133
	}
L129:
	;
	v605 = *(*int32)(unsafe.Add(mBase, uint32(v593)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v584))) = v605
	if v605 != 0 {
		v593 = v605
		goto L127
	} else {
		goto L132
	}
L130:
	;
	goto L131
L131:
	;
	goto L128
L132:
	;
	goto L125
L133:
	;
	v614 = v593
	goto L134
L134:
	;
	v623 = *(*int32)(unsafe.Add(mBase, uint32(v614)))
	v624 = *(*int32)(unsafe.Add(mBase, uint32(v574)))
	if v623 != v624 {
		goto L124
	} else {
		goto L136
	}
L135:
	;
	goto L125
L136:
	;
	v626 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v614)+4)))
	if v33&int32(_a_F_thesaurus_lexize_0) == v626 {
		goto L137
	} else {
		goto L138
	}
L137:
	;
	v628 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v614)+6)))
	if v536 == v628 {
		goto L124
	} else {
		goto L140
	}
L138:
	;
	goto L139
L139:
	;
	v630 = *(*int32)(unsafe.Add(mBase, uint32(v614)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v584))) = v630
	if v630 != 0 {
		v614 = v630
		goto L134
	} else {
		goto L141
	}
L140:
	;
	goto L139
L141:
	;
	goto L135
L142:
	;
	v650 = v575 + int32(1)
	goto L144
L143:
	;
	v650 = int32(0)
	goto L144
L144:
	;
	v656 = v614
	v665 = v650
	goto L123
L145:
	;
	goto L122
L146:
	;
	goto L120
L147:
	;
	if v552 != 0 {
		goto L153
	} else {
		goto L154
	}
L148:
	;
	v690 = v28
	goto L149
L149:
	;
	v699 = *(*int32)(unsafe.Add(mBase, uint32(v690)))
	if v699 == v682 {
		goto L147
	} else {
		goto L151
	}
L150:
	;
	v748 = v552
	goto L117
L151:
	;
	v701 = *(*int32)(unsafe.Add(mBase, uint32(v690)+12))
	if v701 != 0 {
		v690 = v701
		goto L149
	} else {
		goto L152
	}
L152:
	;
	goto L150
L153:
	;
	v721 = v552
	goto L156
L154:
	;
	goto L155
L155:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v656)+12)) = v552
	v748 = v656
	goto L117
L156:
	;
	v730 = *(*int32)(unsafe.Add(mBase, uint32(v721)))
	if v730 == v682 {
		v748 = v552
		goto L117
	} else {
		goto L158
	}
L157:
	;
	goto L155
L158:
	;
	v732 = *(*int32)(unsafe.Add(mBase, uint32(v721)+12))
	if v732 != 0 {
		v721 = v732
		goto L156
	} else {
		goto L159
	}
L159:
	;
	goto L157
L160:
	;
	v764 = int32(0)
	goto L162
L162:
	;
	goto L163
L163:
	;
	v825 = v764
	v826 = v764
	goto L168
L168:
	;
	v835 = v535 + v825<<(uint(int32(2))%32)
	v836 = *(*int32)(unsafe.Add(mBase, uint32(v835)))
	v837 = *(*int32)(unsafe.Add(mBase, uint32(v836)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v835))) = v837
	v839 = int32(1)
	v842 = v826 + v839
	if v842 != v536 {
		v825 = v825 + v839
		v826 = v842
		goto L168
	} else {
		goto L170
	}
L169:
	;
	v552 = v748
	goto L115
L170:
	;
	goto L169
L171:
	;
	goto L4
L172:
	;
	v990 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+1)) = uint8(v990)
	goto L2
L173:
	;
	v912 = v884
	v914 = v877
	v926 = int32(0)
	goto L175
L174:
	;
	v889 = v875
	v890 = v877
	goto L176
L175:
	;
	if v914 != 0 {
		goto L180
	} else {
		goto L181
	}
L176:
	;
	if v890 != 0 {
		goto L172
	} else {
		goto L178
	}
L177:
	;
	v912 = v908
	v914 = v905
	v926 = int32(1)
	goto L175
L178:
	;
	v902 = *(*int32)(unsafe.Add(mBase, uint32(v889)))
	v903 = *(*int32)(unsafe.Add(mBase, uint32(v889)+12))
	v905 = base.B2i32(v903 == int32(0))
	v908 = v880 + v902<<(uint(int32(3))%32)
	v909 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v908))))
	if v909 != v879 {
		v889 = v903
		v890 = v905
		goto L176
	} else {
		goto L179
	}
L179:
	;
	goto L177
L180:
	;
	v928 = v926
	goto L182
L181:
	;
	v928 = int32(1)
	goto L182
L182:
	;
	v930 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v912)+2)))
	v933 = F_palloc_mul(m, int32(8), v930+int32(1))
	mBase = m.M
	v934 = m.ExcPending
	if v934 != 0 {
		goto L17
	} else {
		goto L183
	}
L183:
	;
	v935 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v912)+2)))
	if v935 != 0 {
		goto L184
	} else {
		goto L185
	}
L184:
	;
	v938 = int32(0)
	goto L187
L185:
	;
	v985 = v933
	goto L186
L186:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v985)+4)) = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+1)) = uint8(v928)
	v1021 = base.I64_extend_i32_u(v933)
	goto L1
L187:
	;
	v952 = v938 << (uint(int32(3)) % 32)
	v953 = v933 + v952
	v954 = *(*int32)(unsafe.Add(mBase, uint32(v912)+4))
	v956 = *(*int64)(unsafe.Add(mBase, uint32(v954+v952)))
	*(*int64)(unsafe.Add(mBase, uint32(v953))) = v956
	v958 = *(*int32)(unsafe.Add(mBase, uint32(v912)+4))
	v960 = *(*int32)(unsafe.Add(mBase, uint32(v958+v952)+4))
	v961 = F_pstrdup(m, v960)
	mBase = m.M
	v962 = m.ExcPending
	if v962 != 0 {
		goto L17
	} else {
		goto L189
	}
L188:
	;
	v985 = v933 + v966<<(uint(int32(3))%32)
	goto L186
L189:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v953)+4)) = v961
	v965 = v938 + int32(1)
	v966 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v912)+2)))
	if base.Ui32(v965) < base.Ui32(v966) {
		v938 = v965
		goto L187
	} else {
		goto L190
	}
L190:
	;
	goto L188
}
func F_tidlarger(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int64
	_ = v4
	var v5 int64
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v34 int64
	_ = v34
	v4 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v5 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v6 = base.I32_wrap_i64(v5)
	v7 = base.I32_wrap_i64(v4)
	v11 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v6)+2)))
	v12 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v6))))
	v13 = int32(16)
	v15 = v11 | v12<<(uint(v13)%32)
	v16 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v7)+2)))
	v17 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v7))))
	v20 = v16 | v17<<(uint(v13)%32)
	if base.Ui32(v15) < base.Ui32(v20) {
		v31 = int32(-1)
	} else {
		if base.Ui32(v20) < base.Ui32(v15) {
			v31 = int32(1)
		} else {
			v25 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v6)+4)))
			v26 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v7)+4)))
			if base.Ui32(v25) < base.Ui32(v26) {
				v31 = int32(-1)
			} else {
				v31 = base.B2i32(base.Ui32(v26) < base.Ui32(v25))
			}
		}
	}
	if v31 < int32(0) {
		v34 = v4
	} else {
		v34 = v5
	}
	return v34 & int64(4294967295)
}
func F_timestamp2date_safe(m *base.Module, l0 int64, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
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
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v80 int32
	_ = v80
	v5 = m.G0
	v7 = v5 - int32(48)
	m.G0 = v7
	if l0 == int64(-9223372036854775807-1) {
		v80 = int32(-2147483648)
		m.G0 = v7 + int32(48)
		return v80
	} else {
		if l0 == int64(9223372036854775807) {
			v80 = int32(2147483647)
			m.G0 = v7 + int32(48)
			return v80
		} else {
			v15 = int32(0)
			v20 = F_timestamp2tm(m, l0, v15, v7+int32(4), v7, v15, v15)
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int32(0)
			} else {
				if v20 != 0 {
					v28 = base.I32_wrap_i64(l0>>(uint(int64(63))%64)) ^ int32(2147483647)
					v29 = F_errsave_start(m, l1)
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
						return int32(0)
					} else {
						if v29 == int32(0) {
							v80 = v28
							m.G0 = v7 + int32(48)
							return v80
						} else {
							F_errcode(m, int32(134217858))
							mBase = m.M
							v35 = m.ExcPending
							if v35 != 0 {
								return int32(0)
							} else {
								F_errmsg(m, int32(_a_F_timestamp2date_safe_0), int32(0))
								mBase = m.M
								v39 = m.ExcPending
								if v39 != 0 {
									return int32(0)
								} else {
									F_errsave_finish(m, l1, int32(_a_F_timestamp2date_safe_1), int32(1378), int32(_a_F_timestamp2date_safe_2))
									mBase = m.M
									v44 = m.ExcPending
									if v44 != 0 {
										return int32(0)
									} else {
										v80 = v28
										m.G0 = v7 + int32(48)
										return v80
									}
								}
							}
						}
					}
				} else {
					v45 = *(*int32)(unsafe.Add(mBase, uint32(v7)+24))
					v46 = *(*int32)(unsafe.Add(mBase, uint32(v7)+20))
					v47 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
					v52 = base.B2i32(int32(2) < v46)
					if int32(2) < v46 {
						v53 = int32(_a_F_timestamp2date_safe_3)
					} else {
						v53 = int32(_a_F_timestamp2date_safe_4)
					}
					v54 = v53 + v45
					v59 = base.I32_div_s(v54, int32(4))
					v62 = base.I32_div_s(v54, int32(-100))
					v65 = base.I32_div_s(v54, int32(400))
					if int32(2) < v46 {
						v69 = int32(1)
					} else {
						v69 = int32(13)
					}
					v74 = base.I32_div_s((v69+v46)*int32(_a_F_timestamp2date_safe_5), int32(256))
					v80 = v47 + v54*int32(365) + v59 + v62 + v65 + v74 - int32(_a_F_timestamp2date_safe_6) - int32(_a_F_timestamp2date_safe_7)
					m.G0 = v7 + int32(48)
					return v80
				}
			}
		}
	}
}
func F_timestamp2timestamptz_safe(m *base.Module, l0 int64, l1 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v15 int64
	_ = v15
	var v23 int64
	_ = v23
	var v24 int64
	_ = v24
	var v27 int64
	_ = v27
	var v30 int32
	_ = v30
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v101 int32
	_ = v101
	var v111 int64
	_ = v111
	var v113 int64
	_ = v113
	var v118 int64
	_ = v118
	var v120 int64
	_ = v120
	var v127 int64
	_ = v127
	var v129 int32
	_ = v129
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v151 int64
	_ = v151
	var v161 int64
	_ = v161
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v179 int32
	_ = v179
	var v180 int64
	_ = v180
	v6 = m.G0
	v8 = v6 - int32(48)
	m.G0 = v8
	if base.Ui64(l0-int64(9223372036854775807)) < base.Ui64(int64(2)) {
		v180 = l0
		m.G0 = v8 + int32(48)
		return v180
	} else {
		v15 = base.I64_div_s(l0, int64(86400000000))
		if base.Ui64(int64(172799999999)) <= base.Ui64(l0+int64(86399999999)) {
			v23 = v15 * int64(-86400000000)
		} else {
			v23 = int64(0)
		}
		v24 = v23 + l0
		v27 = v24>>(uint(int64(63))%64) + v15
		if int64(-2451545) <= v27 {
			v30 = base.I32_wrap_i64(v27)
			v42 = v30 + int32(_a_F_timestamp2timestamptz_safe_0)
			v43 = int32(_a_F_timestamp2timestamptz_safe_1)
			v44 = base.I32_div_u_s(v42, v43)
			v45 = int32(3)
			v51 = int32(2)
			v56 = base.I32_div_u_s((v44*int32(1073595727)+v42)<<(uint(v51)%32)|v45, v43)
			v59 = v30 + int32(_a_F_timestamp2timestamptz_safe_2) + v44*v45 + v56 + int32(_a_F_timestamp2timestamptz_safe_3)
			v60 = int32(1461)
			v61 = base.I32_div_u_s(v59, v60)
			v64 = v61*int32(-1461) + v59
			v66 = v64 << (uint(v51) % 32)
			if base.Ui32(v60) <= base.Ui32(v66) {
				v72 = base.I32_rem_u_s(v64+int32(305), int32(365))
				v77 = v72
			} else {
				v76 = base.I32_rem_u_s(v64+int32(306), int32(366))
				v77 = v76
			}
			v79 = base.I32_div_u_s(v66, int32(1461))
			*(*int32)(unsafe.Add(mBase, uint32(v8+int32(24)))) = v79 + v61<<(uint(int32(2))%32) - int32(_a_F_timestamp2timestamptz_safe_4)
			v87 = v77 + int32(123)
			v91 = int32(base.Ui32(v87*int32(2141)) >> (uint(int32(16)) % 32))
			*(*int32)(unsafe.Add(mBase, uint32(v8+int32(16)))) = v87 - int32(base.Ui32(v91*int32(_a_F_timestamp2timestamptz_safe_5))>>(uint(int32(8))%32))
			v101 = base.I32_rem_u_s(v91+int32(10), int32(12))
			*(*int32)(unsafe.Add(mBase, uint32(v8+int32(20)))) = v101 + int32(1)
			*(*int64)(unsafe.Add(mBase, uint32(v8)+36)) = int64(4294967295)
			if v24 < int64(0) {
				v111 = v24 + int64(86400000000)
			} else {
				v111 = v24
			}
			v113 = base.I64_div_s(v111, int64(3600000000))
			*(*uint32)(unsafe.Add(mBase, uint32(v8)+12)) = uint32(v113)
			v118 = base.I64_extend32_s(v113)*int64(-3600000000) + v111
			v120 = base.I64_div_s(v118, int64(60000000))
			*(*uint32)(unsafe.Add(mBase, uint32(v8)+8)) = uint32(v120)
			v127 = base.I64_div_s(base.I64_extend32_s(v120)*int64(-60000000)+v118, int64(1000000))
			*(*uint32)(unsafe.Add(mBase, uint32(v8)+4)) = uint32(v127)
			v129 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v8)+44)) = v129
			v135 = *(*int32)(unsafe.Add(mBase, _c_F_timestamp2timestamptz_safe[0]))
			v137 = m.G0
			v138 = int32(16)
			v139 = v137 - v138
			m.G0 = v139
			v143 = F_DetermineTimeZoneOffsetInternal(m, v8+int32(4), v135, v139+int32(8))
			mBase = m.M
			m.G0 = v139 + v138
			v151 = base.I64_extend_i32_s(v129-v143)*int64(-1000000) + l0
			if base.Ui64(v151+int64(211813488000000000)) < base.Ui64(int64(-9011559254509551616)) {
				v180 = v151
				m.G0 = v8 + int32(48)
				return v180
			} else {
				v161 = l0>>(uint(int64(63))%64) ^ int64(9223372036854775807)
				v162 = F_errsave_start(m, l1)
				mBase = m.M
				v165 = m.ExcPending
				if v165 != 0 {
					return int64(0)
				} else {
					if v162 == int32(0) {
						v180 = v161
						m.G0 = v8 + int32(48)
						return v180
					} else {
						F_errcode(m, int32(134217858))
						mBase = m.M
						v170 = m.ExcPending
						if v170 != 0 {
							return int64(0)
						} else {
							F_errmsg(m, int32(_a_F_timestamp2timestamptz_safe_6), int32(0))
							mBase = m.M
							v174 = m.ExcPending
							if v174 != 0 {
								return int64(0)
							} else {
								F_errsave_finish(m, l1, int32(_a_F_timestamp2timestamptz_safe_7), int32(_a_F_timestamp2timestamptz_safe_8), int32(_a_F_timestamp2timestamptz_safe_9))
								mBase = m.M
								v179 = m.ExcPending
								if v179 != 0 {
									return int64(0)
								} else {
									v180 = v161
									m.G0 = v8 + int32(48)
									return v180
								}
							}
						}
					}
				}
			}
		} else {
			v161 = l0>>(uint(int64(63))%64) ^ int64(9223372036854775807)
			v162 = F_errsave_start(m, l1)
			mBase = m.M
			v165 = m.ExcPending
			if v165 != 0 {
				return int64(0)
			} else {
				if v162 == int32(0) {
					v180 = v161
					m.G0 = v8 + int32(48)
					return v180
				} else {
					F_errcode(m, int32(134217858))
					mBase = m.M
					v170 = m.ExcPending
					if v170 != 0 {
						return int64(0)
					} else {
						F_errmsg(m, int32(_a_F_timestamp2timestamptz_safe_6), int32(0))
						mBase = m.M
						v174 = m.ExcPending
						if v174 != 0 {
							return int64(0)
						} else {
							F_errsave_finish(m, l1, int32(_a_F_timestamp2timestamptz_safe_7), int32(_a_F_timestamp2timestamptz_safe_8), int32(_a_F_timestamp2timestamptz_safe_9))
							mBase = m.M
							v179 = m.ExcPending
							if v179 != 0 {
								return int64(0)
							} else {
								v180 = v161
								m.G0 = v8 + int32(48)
								return v180
							}
						}
					}
				}
			}
		}
	}
}
func F_timestamptypmodin(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14384(m, l0, int32(0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return v3
	}
}
func F_timestamptztypmodin(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14384(m, l0, int32(1))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return v3
	}
}
func F_timetypmodout(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14386(m, l0, int32(_a_F_timetypmodout_0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return v3
	}
}
func F_timingsafe_bcmp(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v13 int32
	_ = v13
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	v4 = int32(0)
	if l2 == v4 {
		return int32(0)
	} else {
		v13 = l2 & int32(3)
		if base.Ui32(l2) < base.Ui32(int32(4)) {
			v52 = l0
			v53 = l1
			v54 = int32(0)
			v59 = v52
			v60 = v53
			v61 = v54
			v65 = v4
			for {
				v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60))))
				v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59))))
				v69 = v61 | (v66 ^ v67)
				v70 = int32(1)
				v75 = v65 + v70
				if v75 != v13 {
					v59 = v59 + v70
					v60 = v60 + v70
					v61 = v69
					v65 = v75
					continue
				} else {
					break
				}
				break
			}
			v79 = v69
		} else {
			v20 = l0
			v21 = l1
			v22 = int32(0)
			v25 = v4
			for {
				v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
				v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
				v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+1)))
				v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+1)))
				v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+2)))
				v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+2)))
				v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+3)))
				v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+3)))
				v42 = v22 | (v27 ^ v28) | (v31 ^ v32) | (v35 ^ v36) | (v39 ^ v40)
				v43 = int32(4)
				v44 = v21 + v43
				v46 = v20 + v43
				v48 = v25 + v43
				if v48 != l2&int32(-4) {
					v20 = v46
					v21 = v44
					v22 = v42
					v25 = v48
					continue
				} else {
					break
				}
				break
			}
			if v13 == int32(0) {
				v79 = v42
			} else {
				v52 = v46
				v53 = v44
				v54 = v42
				v59 = v52
				v60 = v53
				v61 = v54
				v65 = v4
				for {
					v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60))))
					v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59))))
					v69 = v61 | (v66 ^ v67)
					v70 = int32(1)
					v75 = v65 + v70
					if v75 != v13 {
						v59 = v59 + v70
						v60 = v60 + v70
						v61 = v69
						v65 = v75
						continue
					} else {
						break
					}
					break
				}
				v79 = v69
			}
		}
		return base.B2i32(v79 != int32(0))
	}
}
func F_tliInHistory(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v10 int32
	_ = v10
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
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	v3 = int32(0)
	if l1 == v3 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	goto L3
L3:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v10 <= int32(0) {
		v37 = v3
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return v37
L5:
	;
	v13 = int32(0)
	if v13 < v10 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v16 = v10
	goto L8
L7:
	;
	v16 = v13
	goto L8
L8:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v20 = int32(0)
	goto L9
L9:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v17+v20<<(uint(int32(2))%32))))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
	v29 = base.B2i32(v28 == l0)
	if v28 == l0 {
		v37 = v29
		goto L4
	} else {
		goto L11
	}
L10:
	;
	v37 = v29
	goto L4
L11:
	;
	v31 = v20 + int32(1)
	if v31 != v16 {
		v20 = v31
		goto L9
	} else {
		goto L12
	}
L12:
	;
	goto L10
}
func F_tlist_same_exprs(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
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
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v65 int32
	_ = v65
	v3 = int32(0)
	if l0 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v8 = v7
	goto L3
L2:
	;
	v8 = v3
	goto L3
L3:
	;
	if l1 != 0 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	return v65
L5:
	;
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v11 = v9
	goto L7
L6:
	;
	v11 = int32(0)
	goto L7
L7:
	;
	if v11 == v8 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v16 = v3
	goto L11
L9:
	;
	goto L10
L10:
	;
	v65 = int32(0)
	goto L4
L11:
	;
	v19 = int32(0)
	if l0 == v19 {
		v29 = v19
		goto L13
	} else {
		goto L14
	}
L12:
	;
	goto L10
L13:
	;
	v30 = int32(1)
	if l1 == int32(0) {
		v65 = v30
		goto L4
	} else {
		goto L16
	}
L14:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v23 <= v16 {
		v29 = int32(0)
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v29 = v25 + v16<<(uint(int32(2))%32)
	goto L13
L16:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if base.B2i32(v29 == int32(0))|base.B2i32(v35 <= v16) != 0 {
		v65 = v30
		goto L4
	} else {
		goto L17
	}
L17:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v38 == int32(0) {
		v65 = v30
		goto L4
	} else {
		goto L18
	}
L18:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v16<<(uint(int32(2))%32)+v38)))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
	v50 = F_equal(m, v46, v49)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	return int32(0)
L20:
	;
	if v50 != 0 {
		v16 = v16 + int32(1)
		goto L11
	} else {
		goto L21
	}
L21:
	;
	goto L12
}
func F_toastrel_valueid_exists(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
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
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v65 int32
	_ = v65
	v6 = m.G0
	v8 = v6 + int32(-64)
	m.G0 = v8
	v13 = F_toast_open_indexes(m, l0, int32(3), v8, v6+int32(-60))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v18 = v6 + int32(-56)
	F_ScanKeyInit(m, v18, int32(1), int32(3), int32(184), base.I64_extend_i32_u(l1))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v25+v13<<(uint(int32(2))%32))))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+56))
	v31 = int32(1)
	v34 = F_systable_beginscan(m, l0, v30, v31, int32(_a_F_toastrel_valueid_exists_0), v31, v18)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v36 = F_systable_getnext(m, v34)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	F_systable_endscan(m, v34)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	if int32(0) < v40 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v44 = int32(0)
	goto L10
L8:
	;
	goto L9
L9:
	;
	F_pfree(m, v25)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L1
	} else {
		goto L14
	}
L10:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v25+v44<<(uint(int32(2))%32))))
	F_relation_close(m, v52, int32(3))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L1
	} else {
		goto L12
	}
L11:
	;
	goto L9
L12:
	;
	v57 = v44 + int32(1)
	if v57 != v40 {
		v44 = v57
		goto L10
	} else {
		goto L13
	}
L13:
	;
	goto L11
L14:
	;
	m.G0 = v8 - int32(-64)
	return base.B2i32(v36 != int32(0))
}
func F_towupper(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
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
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v55 int32
	_ = v55
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v119 int32
	_ = v119
	v2 = int32(1)
	if base.Ui32(int32(_a_F_towupper_0)) < base.Ui32(l0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v119
L2:
	;
	v119 = l0
	goto L1
L3:
	;
	v12 = int32(255)
	v13 = l0 & v12
	v14 = int32(3)
	v15 = base.I32_div_u_s(v13, v14)
	v21 = int32(2)
	v23 = *(*int32)(unsafe.Add(mBase, uint32((l0-v15*v14)&v12<<(uint(v21)%32))+uint32(_c_F_towupper[0])))
	v24 = int32(8)
	v25 = int32(base.Ui32(l0) >> (uint(v24) % 32))
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+uint32(_c_F_towupper[1]))))
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15+v26*int32(86))+uint32(_c_F_towupper[1]))))
	v35 = base.I32_rem_u_s(int32(base.Ui32(v23*v30)>>(uint(int32(11))%32)), int32(6))
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+uint32(_c_F_towupper[2]))))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v35<<(uint(v21)%32)+v38<<(uint(v21)%32))+uint32(_c_F_towupper[3])))
	v44 = v42 >> (uint(v24) % 32)
	v46 = v42 & v12
	if base.Ui32(v46) <= base.Ui32(int32(1)) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v119 = v44&(int32(0)-(v2^v46)) + l0
	goto L1
L5:
	;
	goto L6
L6:
	;
	v55 = v44 & int32(255)
	if v55 == int32(0) {
		goto L2
	} else {
		goto L7
	}
L7:
	;
	v62 = int32(base.Ui32(v44) >> (uint(int32(8)) % 32))
	v63 = v55
	goto L8
L8:
	;
	v69 = int32(1)
	v70 = int32(base.Ui32(v63) >> (uint(v69) % 32))
	v71 = v70 + v62
	v73 = v71 << (uint(v69) % 32)
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73)+uint32(_c_F_towupper[4]))))
	if v74 == v13 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	goto L2
L10:
	;
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73)+uint32(_c_F_towupper[5]))))
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v78<<(uint(int32(2))%32))+uint32(_c_F_towupper[3])))
	v83 = v81 & int32(255)
	if base.Ui32(v83) <= base.Ui32(int32(1)) {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	goto L12
L12:
	;
	v97 = base.B2i32(base.Ui32(v13) < base.Ui32(v74))
	if base.Ui32(v13) < base.Ui32(v74) {
		goto L19
	} else {
		goto L20
	}
L13:
	;
	v119 = (int32(0)-(v2^v83))&(v81>>(uint(int32(8))%32)) + l0
	goto L1
L14:
	;
	goto L15
L15:
	;
	goto L16
L16:
	;
	goto L18
L18:
	;
	v119 = int32(-1) + l0
	goto L1
L19:
	;
	v98 = v62
	goto L21
L20:
	;
	v98 = v71
	goto L21
L21:
	;
	if base.Ui32(v13) < base.Ui32(v74) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v100 = v70
	goto L24
L23:
	;
	v100 = v63 - v70
	goto L24
L24:
	;
	if v100 != 0 {
		v62 = v98
		v63 = v100
		goto L8
	} else {
		goto L25
	}
L25:
	;
	goto L9
}
func F_transformLimitClause(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	if l1 == int32(0) {
		v45 = int32(0)
		return v45
	} else {
		v9 = F_transformExpr(m, l0, l1, l2)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			v14 = F_coerce_to_specific_type(m, l0, v9, int32(20), l3)
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				F_checkExprIsVarFree(m, l0, v14, l3)
				mBase = m.M
				v17 = m.ExcPending
				if v17 != 0 {
					return int32(0)
				} else {
					if base.B2i32(l2 != int32(22))|base.B2i32(l4 != int32(1)) != 0 {
						v45 = v14
						return v45
					} else {
						v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
						if v23 != int32(72) {
							v45 = v14
							return v45
						} else {
							v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+12)))
							if v26 != int32(1) {
								v45 = v14
								return v45
							} else {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v32 = m.ExcPending
								if v32 != 0 {
									return int32(0)
								} else {
									F_errcode(m, int32(654573698))
									mBase = m.M
									v35 = m.ExcPending
									if v35 != 0 {
										return int32(0)
									} else {
										F_errmsg(m, int32(_a_F_transformLimitClause_0), int32(0))
										mBase = m.M
										v39 = m.ExcPending
										if v39 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_transformLimitClause_1), int32(1909), int32(_a_F_transformLimitClause_2))
											mBase = m.M
											v44 = m.ExcPending
											if v44 != 0 {
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
	}
}
func F_transformLockingClause(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
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
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v109 int32
	_ = v109
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int64
	_ = v238
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v312 int32
	_ = v312
	var v315 int32
	_ = v315
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v327 int32
	_ = v327
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v354 int32
	_ = v354
	var v358 int32
	_ = v358
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v380 int32
	_ = v380
	var v385 int32
	_ = v385
	var v389 int32
	_ = v389
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v395 int32
	_ = v395
	var v400 int32
	_ = v400
	var v402 int32
	_ = v402
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v411 int32
	_ = v411
	var v416 int32
	_ = v416
	var v420 int32
	_ = v420
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v426 int32
	_ = v426
	var v431 int32
	_ = v431
	var v433 int32
	_ = v433
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v442 int32
	_ = v442
	var v447 int32
	_ = v447
	var v451 int32
	_ = v451
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v457 int32
	_ = v457
	var v462 int32
	_ = v462
	var v464 int32
	_ = v464
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v473 int32
	_ = v473
	var v478 int32
	_ = v478
	var v482 int32
	_ = v482
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v488 int32
	_ = v488
	var v493 int32
	_ = v493
	var v495 int32
	_ = v495
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v504 int32
	_ = v504
	var v509 int32
	_ = v509
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v520 int32
	_ = v520
	var v525 int32
	_ = v525
	var v545 int32
	_ = v545
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v552 int32
	_ = v552
	var v557 int32
	_ = v557
	var v559 int32
	_ = v559
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v567 int32
	_ = v567
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v576 int32
	_ = v576
	var v585 int32
	_ = v585
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v599 int32
	_ = v599
	var v600 int32
	_ = v600
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v605 int32
	_ = v605
	var v608 int32
	_ = v608
	var v613 int32
	_ = v613
	var v614 int32
	_ = v614
	var v617 int32
	_ = v617
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v623 int32
	_ = v623
	var v631 int32
	_ = v631
	var v632 int32
	_ = v632
	var v635 int32
	_ = v635
	var v646 int32
	_ = v646
	var v648 int32
	_ = v648
	var v649 int32
	_ = v649
	var v651 int32
	_ = v651
	var v653 int32
	_ = v653
	var v655 int32
	_ = v655
	var v657 int32
	_ = v657
	var v660 int32
	_ = v660
	var v661 int32
	_ = v661
	var v668 int32
	_ = v668
	var v669 int32
	_ = v669
	var v670 int32
	_ = v670
	var v674 int32
	_ = v674
	var v675 int32
	_ = v675
	var v676 int32
	_ = v676
	var v677 int64
	_ = v677
	var v681 int32
	_ = v681
	var v682 int32
	_ = v682
	var v685 int32
	_ = v685
	var v690 int32
	_ = v690
	var v691 int32
	_ = v691
	var v694 int32
	_ = v694
	var v697 int32
	_ = v697
	var v698 int32
	_ = v698
	var v700 int32
	_ = v700
	var v708 int32
	_ = v708
	var v709 int32
	_ = v709
	var v712 int32
	_ = v712
	var v723 int32
	_ = v723
	var v725 int32
	_ = v725
	var v726 int32
	_ = v726
	var v728 int32
	_ = v728
	var v730 int32
	_ = v730
	var v732 int32
	_ = v732
	var v734 int32
	_ = v734
	var v737 int32
	_ = v737
	var v738 int32
	_ = v738
	var v745 int32
	_ = v745
	var v746 int32
	_ = v746
	var v747 int32
	_ = v747
	var v751 int32
	_ = v751
	var v754 int32
	_ = v754
	var v759 int32
	_ = v759
	v4 = l3
	v5 = int32(0)
	v15 = m.G0
	v17 = v15 - int32(144)
	m.G0 = v17
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	F_CheckSelectLocking(m, l1, v20)
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
	v24 = F_palloc0(m, int32(16))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v24))) = int64(94)
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+8)) = v28
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+12)) = v30
	if v19 != 0 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	m.G0 = v17 + int32(144)
	return
L5:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	if v32 <= int32(0) {
		goto L4
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	v573 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	if v573 == int32(0) {
		goto L4
	} else {
		goto L183
	}
L8:
	;
	v46 = v5
	goto L9
L9:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v49+v46<<(uint(int32(2))%32))))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
	if v54 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v545 = m.ExcPending
	if v545 != 0 {
		goto L1
	} else {
		goto L174
	}
L11:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	if v91 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L12:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v53)+8))
	if v57 == int32(0) {
		goto L11
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L1
	} else {
		goto L16
	}
L15:
	;
	goto L14
L16:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v69 = v67 - int32(1)
	if base.Ui32(v69) <= base.Ui32(int32(3)) {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+128)) = v76
	F_errmsg(m, int32(_a_F_transformLockingClause_0), v17+int32(128))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L1
	} else {
		goto L22
	}
L19:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v69<<(uint(int32(2))%32))+uint32(_c_F_transformLockingClause[0])))
	v76 = v74
	goto L21
L20:
	;
	v76 = int32(_a_F_transformLockingClause_1)
	goto L21
L21:
	;
	goto L18
L22:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v53)+24))
	F_parser_errposition(m, l0, v83)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	F_errfinish(m, int32(_a_F_transformLockingClause_2), int32(3548), int32(_a_F_transformLockingClause_3))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L25:
	;
	goto L10
L26:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v91)+4))
	if v94 <= int32(0) {
		goto L25
	} else {
		goto L27
	}
L27:
	;
	v97 = int32(0)
	if v97 < v94 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v100 = v94
	goto L30
L29:
	;
	v100 = v97
	goto L30
L30:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v91)+12))
	v109 = int32(0)
	goto L31
L31:
	;
	v117 = int32(1)
	v118 = v109 + v117
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v101+v109<<(uint(int32(2))%32))))
	v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122)+125)))
	if v123 != v117 {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	goto L25
L33:
	;
	if v118 != v100 {
		v109 = v118
		goto L31
	} else {
		goto L173
	}
L34:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v122)+8))
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v126)+4))
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v122)+4))
	if v128 != 0 {
		v136 = v127
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v53)+12))
	v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v136))))
	v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v137))))
	if base.B2i32(v140 == int32(0))|base.B2i32(v140 != v143) != 0 {
		v161 = v140
		v162 = v143
		goto L40
	} else {
		goto L41
	}
L36:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v122)+12))
	switch v129 - int32(1) {
	case 0, 4:
		goto L33
	case 1:
		goto L37
	default:
		v136 = v127
		goto L35
	}
L37:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v122)+64))
	if v132 == int32(0) {
		goto L33
	} else {
		goto L38
	}
L38:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v132)+4))
	v136 = v135
	goto L35
L39:
	;
	if v161-v162 != 0 {
		goto L33
	} else {
		goto L46
	}
L40:
	;
	goto L39
L41:
	;
	v146 = v136
	v147 = v137
	goto L42
L42:
	;
	v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v147)+1)))
	v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146)+1)))
	if v151 == int32(0) {
		v161 = v151
		v162 = v150
		goto L40
	} else {
		goto L44
	}
L43:
	;
	v161 = v151
	v162 = v150
	goto L40
L44:
	;
	v154 = int32(1)
	if v151 == v150 {
		v146 = v146 + v154
		v147 = v147 + v154
		goto L42
	} else {
		goto L45
	}
L45:
	;
	goto L43
L46:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v122)+12))
	switch v164 {
	case 0:
		goto L56
	case 1:
		goto L55
	case 2:
		goto L53
	case 3:
		goto L52
	case 4:
		goto L51
	case 5:
		goto L50
	case 6:
		goto L49
	case 7:
		goto L48
	default:
		goto L47
	}
L47:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v513 = m.ExcPending
	if v513 != 0 {
		goto L1
	} else {
		goto L170
	}
L48:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v482 = m.ExcPending
	if v482 != 0 {
		goto L1
	} else {
		goto L161
	}
L49:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v451 = m.ExcPending
	if v451 != 0 {
		goto L1
	} else {
		goto L152
	}
L50:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v420 = m.ExcPending
	if v420 != 0 {
		goto L1
	} else {
		goto L143
	}
L51:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v389 = m.ExcPending
	if v389 != 0 {
		goto L1
	} else {
		goto L134
	}
L52:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L1
	} else {
		goto L125
	}
L53:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L1
	} else {
		goto L116
	}
L54:
	;
	v321 = v46 + int32(1)
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	if v321 < v322 {
		v46 = v321
		goto L9
	} else {
		goto L115
	}
L55:
	;
	v242 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v243 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	if v4 == int32(0) {
		goto L86
	} else {
		goto L87
	}
L56:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v166 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	if v4 == int32(0) {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v169 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+43)) = uint8(v169)
	goto L59
L58:
	;
	goto L59
L59:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(l1)+140))
	if v174 != 0 {
		goto L63
	} else {
		goto L64
	}
L60:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v236 = F_getRTEPermissionInfo(m, v235, v122)
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L1
	} else {
		goto L85
	}
L61:
	;
	if v207 != 0 {
		goto L74
	} else {
		goto L75
	}
L62:
	;
	goto L61
L63:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v174)+4))
	if v175 <= int32(0) {
		v207 = int32(0)
		goto L62
	} else {
		goto L66
	}
L64:
	;
	goto L65
L65:
	;
	v207 = int32(0)
	goto L62
L66:
	;
	v178 = int32(0)
	if v178 < v175 {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v181 = v175
	goto L69
L68:
	;
	v181 = v178
	goto L69
L69:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v174)+12))
	v184 = int32(0)
	goto L70
L70:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v182+v184<<(uint(int32(2))%32))))
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v192)+4))
	if v193 == v118 {
		v207 = v192
		goto L62
	} else {
		goto L72
	}
L71:
	;
	goto L65
L72:
	;
	v196 = v184 + int32(1)
	if v196 != v181 {
		v184 = v196
		goto L70
	} else {
		goto L73
	}
L73:
	;
	goto L71
L74:
	;
	v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207)+16)))
	v210 = v4 & v209
	*(*uint8)(unsafe.Add(mBase, uint32(v207)+16)) = uint8(v210)
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v207)+8))
	if base.Ui32(v166) < base.Ui32(v212) {
		goto L77
	} else {
		goto L78
	}
L75:
	;
	goto L76
L76:
	;
	v221 = F_palloc0(m, int32(20))
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L1
	} else {
		goto L83
	}
L77:
	;
	v214 = v212
	goto L79
L78:
	;
	v214 = v166
	goto L79
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v207)+8)) = v214
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v207)+12))
	if base.Ui32(v165) < base.Ui32(v216) {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v218 = v216
	goto L82
L81:
	;
	v218 = v165
	goto L82
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v207)+12)) = v218
	goto L60
L83:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v221)+16)) = uint8(v4)
	*(*int32)(unsafe.Add(mBase, uint32(v221)+12)) = v165
	*(*int32)(unsafe.Add(mBase, uint32(v221)+8)) = v166
	*(*int32)(unsafe.Add(mBase, uint32(v221)+4)) = v118
	*(*int32)(unsafe.Add(mBase, uint32(v221))) = int32(109)
	v229 = *(*int32)(unsafe.Add(mBase, uint32(l1)+140))
	v230 = F_lappend(m, v229, v221)
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L1
	} else {
		goto L84
	}
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+140)) = v230
	goto L60
L85:
	;
	v238 = *(*int64)(unsafe.Add(mBase, uint32(v236)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v236)+16)) = v238 | int64(4)
	goto L54
L86:
	;
	v246 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+43)) = uint8(v246)
	goto L88
L87:
	;
	goto L88
L88:
	;
	v251 = *(*int32)(unsafe.Add(mBase, uint32(l1)+140))
	if v251 != 0 {
		goto L92
	} else {
		goto L93
	}
L89:
	;
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v122)+36))
	F_transformLockingClause(m, l0, v312, v24, int32(1))
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L1
	} else {
		goto L114
	}
L90:
	;
	if v284 != 0 {
		goto L103
	} else {
		goto L104
	}
L91:
	;
	goto L90
L92:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v251)+4))
	if v252 <= int32(0) {
		v284 = int32(0)
		goto L91
	} else {
		goto L95
	}
L93:
	;
	goto L94
L94:
	;
	v284 = int32(0)
	goto L91
L95:
	;
	v255 = int32(0)
	if v255 < v252 {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v258 = v252
	goto L98
L97:
	;
	v258 = v255
	goto L98
L98:
	;
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v251)+12))
	v261 = int32(0)
	goto L99
L99:
	;
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v259+v261<<(uint(int32(2))%32))))
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v269)+4))
	if v270 == v118 {
		v284 = v269
		goto L91
	} else {
		goto L101
	}
L100:
	;
	goto L94
L101:
	;
	v273 = v261 + int32(1)
	if v273 != v258 {
		v261 = v273
		goto L99
	} else {
		goto L102
	}
L102:
	;
	goto L100
L103:
	;
	v286 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v284)+16)))
	v287 = v4 & v286
	*(*uint8)(unsafe.Add(mBase, uint32(v284)+16)) = uint8(v287)
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v284)+8))
	if base.Ui32(v243) < base.Ui32(v289) {
		goto L106
	} else {
		goto L107
	}
L104:
	;
	goto L105
L105:
	;
	v298 = F_palloc0(m, int32(20))
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L1
	} else {
		goto L112
	}
L106:
	;
	v291 = v289
	goto L108
L107:
	;
	v291 = v243
	goto L108
L108:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v284)+8)) = v291
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v284)+12))
	if base.Ui32(v242) < base.Ui32(v293) {
		goto L109
	} else {
		goto L110
	}
L109:
	;
	v295 = v293
	goto L111
L110:
	;
	v295 = v242
	goto L111
L111:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v284)+12)) = v295
	goto L89
L112:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v298)+16)) = uint8(v4)
	*(*int32)(unsafe.Add(mBase, uint32(v298)+12)) = v242
	*(*int32)(unsafe.Add(mBase, uint32(v298)+8)) = v243
	*(*int32)(unsafe.Add(mBase, uint32(v298)+4)) = v118
	*(*int32)(unsafe.Add(mBase, uint32(v298))) = int32(109)
	v306 = *(*int32)(unsafe.Add(mBase, uint32(l1)+140))
	v307 = F_lappend(m, v306, v298)
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L1
	} else {
		goto L113
	}
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+140)) = v307
	goto L89
L114:
	;
	goto L54
L115:
	;
	goto L4
L116:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L1
	} else {
		goto L117
	}
L117:
	;
	v331 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v333 = v331 - int32(1)
	if base.Ui32(v333) <= base.Ui32(int32(3)) {
		goto L119
	} else {
		goto L120
	}
L118:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+32)) = v340
	F_errmsg(m, int32(_a_F_transformLockingClause_4), v17+int32(32))
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L1
	} else {
		goto L122
	}
L119:
	;
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v333<<(uint(int32(2))%32))+uint32(_c_F_transformLockingClause[0])))
	v340 = v338
	goto L121
L120:
	;
	v340 = int32(_a_F_transformLockingClause_1)
	goto L121
L121:
	;
	goto L118
L122:
	;
	v347 = *(*int32)(unsafe.Add(mBase, uint32(v53)+24))
	F_parser_errposition(m, l0, v347)
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L1
	} else {
		goto L123
	}
L123:
	;
	F_errfinish(m, int32(_a_F_transformLockingClause_2), int32(3612), int32(_a_F_transformLockingClause_3))
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L1
	} else {
		goto L124
	}
L124:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L125:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L1
	} else {
		goto L126
	}
L126:
	;
	v362 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v364 = v362 - int32(1)
	if base.Ui32(v364) <= base.Ui32(int32(3)) {
		goto L128
	} else {
		goto L129
	}
L127:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+48)) = v371
	F_errmsg(m, int32(_a_F_transformLockingClause_5), v17+int32(48))
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L1
	} else {
		goto L131
	}
L128:
	;
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v364<<(uint(int32(2))%32))+uint32(_c_F_transformLockingClause[0])))
	v371 = v369
	goto L130
L129:
	;
	v371 = int32(_a_F_transformLockingClause_1)
	goto L130
L130:
	;
	goto L127
L131:
	;
	v378 = *(*int32)(unsafe.Add(mBase, uint32(v53)+24))
	F_parser_errposition(m, l0, v378)
	mBase = m.M
	v380 = m.ExcPending
	if v380 != 0 {
		goto L1
	} else {
		goto L132
	}
L132:
	;
	F_errfinish(m, int32(_a_F_transformLockingClause_2), int32(3621), int32(_a_F_transformLockingClause_3))
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L1
	} else {
		goto L133
	}
L133:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L134:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L1
	} else {
		goto L135
	}
L135:
	;
	v393 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v395 = v393 - int32(1)
	if base.Ui32(v395) <= base.Ui32(int32(3)) {
		goto L137
	} else {
		goto L138
	}
L136:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+64)) = v402
	F_errmsg(m, int32(_a_F_transformLockingClause_6), v17-int32(-64))
	mBase = m.M
	v408 = m.ExcPending
	if v408 != 0 {
		goto L1
	} else {
		goto L140
	}
L137:
	;
	v400 = *(*int32)(unsafe.Add(mBase, uint32(v395<<(uint(int32(2))%32))+uint32(_c_F_transformLockingClause[0])))
	v402 = v400
	goto L139
L138:
	;
	v402 = int32(_a_F_transformLockingClause_1)
	goto L139
L139:
	;
	goto L136
L140:
	;
	v409 = *(*int32)(unsafe.Add(mBase, uint32(v53)+24))
	F_parser_errposition(m, l0, v409)
	mBase = m.M
	v411 = m.ExcPending
	if v411 != 0 {
		goto L1
	} else {
		goto L141
	}
L141:
	;
	F_errfinish(m, int32(_a_F_transformLockingClause_2), int32(3630), int32(_a_F_transformLockingClause_3))
	mBase = m.M
	v416 = m.ExcPending
	if v416 != 0 {
		goto L1
	} else {
		goto L142
	}
L142:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L143:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v423 = m.ExcPending
	if v423 != 0 {
		goto L1
	} else {
		goto L144
	}
L144:
	;
	v424 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v426 = v424 - int32(1)
	if base.Ui32(v426) <= base.Ui32(int32(3)) {
		goto L146
	} else {
		goto L147
	}
L145:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+80)) = v433
	F_errmsg(m, int32(_a_F_transformLockingClause_7), v17+int32(80))
	mBase = m.M
	v439 = m.ExcPending
	if v439 != 0 {
		goto L1
	} else {
		goto L149
	}
L146:
	;
	v431 = *(*int32)(unsafe.Add(mBase, uint32(v426<<(uint(int32(2))%32))+uint32(_c_F_transformLockingClause[0])))
	v433 = v431
	goto L148
L147:
	;
	v433 = int32(_a_F_transformLockingClause_1)
	goto L148
L148:
	;
	goto L145
L149:
	;
	v440 = *(*int32)(unsafe.Add(mBase, uint32(v53)+24))
	F_parser_errposition(m, l0, v440)
	mBase = m.M
	v442 = m.ExcPending
	if v442 != 0 {
		goto L1
	} else {
		goto L150
	}
L150:
	;
	F_errfinish(m, int32(_a_F_transformLockingClause_2), int32(3639), int32(_a_F_transformLockingClause_3))
	mBase = m.M
	v447 = m.ExcPending
	if v447 != 0 {
		goto L1
	} else {
		goto L151
	}
L151:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L152:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v454 = m.ExcPending
	if v454 != 0 {
		goto L1
	} else {
		goto L153
	}
L153:
	;
	v455 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v457 = v455 - int32(1)
	if base.Ui32(v457) <= base.Ui32(int32(3)) {
		goto L155
	} else {
		goto L156
	}
L154:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+96)) = v464
	F_errmsg(m, int32(_a_F_transformLockingClause_8), v17+int32(96))
	mBase = m.M
	v470 = m.ExcPending
	if v470 != 0 {
		goto L1
	} else {
		goto L158
	}
L155:
	;
	v462 = *(*int32)(unsafe.Add(mBase, uint32(v457<<(uint(int32(2))%32))+uint32(_c_F_transformLockingClause[0])))
	v464 = v462
	goto L157
L156:
	;
	v464 = int32(_a_F_transformLockingClause_1)
	goto L157
L157:
	;
	goto L154
L158:
	;
	v471 = *(*int32)(unsafe.Add(mBase, uint32(v53)+24))
	F_parser_errposition(m, l0, v471)
	mBase = m.M
	v473 = m.ExcPending
	if v473 != 0 {
		goto L1
	} else {
		goto L159
	}
L159:
	;
	F_errfinish(m, int32(_a_F_transformLockingClause_2), int32(3648), int32(_a_F_transformLockingClause_3))
	mBase = m.M
	v478 = m.ExcPending
	if v478 != 0 {
		goto L1
	} else {
		goto L160
	}
L160:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L161:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v485 = m.ExcPending
	if v485 != 0 {
		goto L1
	} else {
		goto L162
	}
L162:
	;
	v486 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v488 = v486 - int32(1)
	if base.Ui32(v488) <= base.Ui32(int32(3)) {
		goto L164
	} else {
		goto L165
	}
L163:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+112)) = v495
	F_errmsg(m, int32(_a_F_transformLockingClause_9), v17+int32(112))
	mBase = m.M
	v501 = m.ExcPending
	if v501 != 0 {
		goto L1
	} else {
		goto L167
	}
L164:
	;
	v493 = *(*int32)(unsafe.Add(mBase, uint32(v488<<(uint(int32(2))%32))+uint32(_c_F_transformLockingClause[0])))
	v495 = v493
	goto L166
L165:
	;
	v495 = int32(_a_F_transformLockingClause_1)
	goto L166
L166:
	;
	goto L163
L167:
	;
	v502 = *(*int32)(unsafe.Add(mBase, uint32(v53)+24))
	F_parser_errposition(m, l0, v502)
	mBase = m.M
	v504 = m.ExcPending
	if v504 != 0 {
		goto L1
	} else {
		goto L168
	}
L168:
	;
	F_errfinish(m, int32(_a_F_transformLockingClause_2), int32(3657), int32(_a_F_transformLockingClause_3))
	mBase = m.M
	v509 = m.ExcPending
	if v509 != 0 {
		goto L1
	} else {
		goto L169
	}
L169:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L170:
	;
	v514 = *(*int32)(unsafe.Add(mBase, uint32(v122)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v514
	F_errmsg_internal(m, int32(_a_F_transformLockingClause_10), v17+int32(16))
	mBase = m.M
	v520 = m.ExcPending
	if v520 != 0 {
		goto L1
	} else {
		goto L171
	}
L171:
	;
	F_errfinish(m, int32(_a_F_transformLockingClause_2), int32(3664), int32(_a_F_transformLockingClause_3))
	mBase = m.M
	v525 = m.ExcPending
	if v525 != 0 {
		goto L1
	} else {
		goto L172
	}
L172:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L173:
	;
	goto L32
L174:
	;
	F_errcode(m, int32(16908420))
	mBase = m.M
	v548 = m.ExcPending
	if v548 != 0 {
		goto L1
	} else {
		goto L175
	}
L175:
	;
	v549 = *(*int32)(unsafe.Add(mBase, uint32(v53)+12))
	v550 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v552 = v550 - int32(1)
	if base.Ui32(v552) <= base.Ui32(int32(3)) {
		goto L177
	} else {
		goto L178
	}
L176:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+4)) = v559
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v549
	F_errmsg(m, int32(_a_F_transformLockingClause_11), v17)
	mBase = m.M
	v564 = m.ExcPending
	if v564 != 0 {
		goto L1
	} else {
		goto L180
	}
L177:
	;
	v557 = *(*int32)(unsafe.Add(mBase, uint32(v552<<(uint(int32(2))%32))+uint32(_c_F_transformLockingClause[0])))
	v559 = v557
	goto L179
L178:
	;
	v559 = int32(_a_F_transformLockingClause_1)
	goto L179
L179:
	;
	goto L176
L180:
	;
	v565 = *(*int32)(unsafe.Add(mBase, uint32(v53)+24))
	F_parser_errposition(m, l0, v565)
	mBase = m.M
	v567 = m.ExcPending
	if v567 != 0 {
		goto L1
	} else {
		goto L181
	}
L181:
	;
	F_errfinish(m, int32(_a_F_transformLockingClause_2), int32(3678), int32(_a_F_transformLockingClause_3))
	mBase = m.M
	v572 = m.ExcPending
	if v572 != 0 {
		goto L1
	} else {
		goto L182
	}
L182:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L183:
	;
	v576 = *(*int32)(unsafe.Add(mBase, uint32(v573)+4))
	if v576 <= int32(0) {
		goto L4
	} else {
		goto L184
	}
L184:
	;
	v585 = v5
	goto L185
L185:
	;
	v593 = int32(1)
	v594 = v585 + v593
	v595 = *(*int32)(unsafe.Add(mBase, uint32(v573)+12))
	v599 = *(*int32)(unsafe.Add(mBase, uint32(v595+v585<<(uint(int32(2))%32))))
	v600 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v599)+125)))
	if v600 != v593 {
		goto L187
	} else {
		goto L188
	}
L186:
	;
	goto L4
L187:
	;
	v759 = *(*int32)(unsafe.Add(mBase, uint32(v573)+4))
	if v594 < v759 {
		v585 = v594
		goto L185
	} else {
		goto L249
	}
L188:
	;
	v603 = *(*int32)(unsafe.Add(mBase, uint32(v599)+12))
	switch v603 {
	case 0:
		goto L190
	case 1:
		goto L189
	default:
		goto L187
	}
L189:
	;
	v681 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v682 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	if v4 == int32(0) {
		goto L220
	} else {
		goto L221
	}
L190:
	;
	v604 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v605 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	if v4 == int32(0) {
		goto L191
	} else {
		goto L192
	}
L191:
	;
	v608 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+43)) = uint8(v608)
	goto L193
L192:
	;
	goto L193
L193:
	;
	v613 = *(*int32)(unsafe.Add(mBase, uint32(l1)+140))
	if v613 != 0 {
		goto L197
	} else {
		goto L198
	}
L194:
	;
	v674 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v675 = F_getRTEPermissionInfo(m, v674, v599)
	mBase = m.M
	v676 = m.ExcPending
	if v676 != 0 {
		goto L1
	} else {
		goto L219
	}
L195:
	;
	if v646 != 0 {
		goto L208
	} else {
		goto L209
	}
L196:
	;
	goto L195
L197:
	;
	v614 = *(*int32)(unsafe.Add(mBase, uint32(v613)+4))
	if v614 <= int32(0) {
		v646 = int32(0)
		goto L196
	} else {
		goto L200
	}
L198:
	;
	goto L199
L199:
	;
	v646 = int32(0)
	goto L196
L200:
	;
	v617 = int32(0)
	if v617 < v614 {
		goto L201
	} else {
		goto L202
	}
L201:
	;
	v620 = v614
	goto L203
L202:
	;
	v620 = v617
	goto L203
L203:
	;
	v621 = *(*int32)(unsafe.Add(mBase, uint32(v613)+12))
	v623 = int32(0)
	goto L204
L204:
	;
	v631 = *(*int32)(unsafe.Add(mBase, uint32(v621+v623<<(uint(int32(2))%32))))
	v632 = *(*int32)(unsafe.Add(mBase, uint32(v631)+4))
	if v632 == v594 {
		v646 = v631
		goto L196
	} else {
		goto L206
	}
L205:
	;
	goto L199
L206:
	;
	v635 = v623 + int32(1)
	if v635 != v620 {
		v623 = v635
		goto L204
	} else {
		goto L207
	}
L207:
	;
	goto L205
L208:
	;
	v648 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v646)+16)))
	v649 = v4 & v648
	*(*uint8)(unsafe.Add(mBase, uint32(v646)+16)) = uint8(v649)
	v651 = *(*int32)(unsafe.Add(mBase, uint32(v646)+8))
	if base.Ui32(v605) < base.Ui32(v651) {
		goto L211
	} else {
		goto L212
	}
L209:
	;
	goto L210
L210:
	;
	v660 = F_palloc0(m, int32(20))
	mBase = m.M
	v661 = m.ExcPending
	if v661 != 0 {
		goto L1
	} else {
		goto L217
	}
L211:
	;
	v653 = v651
	goto L213
L212:
	;
	v653 = v605
	goto L213
L213:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v646)+8)) = v653
	v655 = *(*int32)(unsafe.Add(mBase, uint32(v646)+12))
	if base.Ui32(v604) < base.Ui32(v655) {
		goto L214
	} else {
		goto L215
	}
L214:
	;
	v657 = v655
	goto L216
L215:
	;
	v657 = v604
	goto L216
L216:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v646)+12)) = v657
	goto L194
L217:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v660)+16)) = uint8(v4)
	*(*int32)(unsafe.Add(mBase, uint32(v660)+12)) = v604
	*(*int32)(unsafe.Add(mBase, uint32(v660)+8)) = v605
	*(*int32)(unsafe.Add(mBase, uint32(v660)+4)) = v594
	*(*int32)(unsafe.Add(mBase, uint32(v660))) = int32(109)
	v668 = *(*int32)(unsafe.Add(mBase, uint32(l1)+140))
	v669 = F_lappend(m, v668, v660)
	mBase = m.M
	v670 = m.ExcPending
	if v670 != 0 {
		goto L1
	} else {
		goto L218
	}
L218:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+140)) = v669
	goto L194
L219:
	;
	v677 = *(*int64)(unsafe.Add(mBase, uint32(v675)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v675)+16)) = v677 | int64(4)
	goto L187
L220:
	;
	v685 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+43)) = uint8(v685)
	goto L222
L221:
	;
	goto L222
L222:
	;
	v690 = *(*int32)(unsafe.Add(mBase, uint32(l1)+140))
	if v690 != 0 {
		goto L226
	} else {
		goto L227
	}
L223:
	;
	v751 = *(*int32)(unsafe.Add(mBase, uint32(v599)+36))
	F_transformLockingClause(m, l0, v751, v24, int32(1))
	mBase = m.M
	v754 = m.ExcPending
	if v754 != 0 {
		goto L1
	} else {
		goto L248
	}
L224:
	;
	if v723 != 0 {
		goto L237
	} else {
		goto L238
	}
L225:
	;
	goto L224
L226:
	;
	v691 = *(*int32)(unsafe.Add(mBase, uint32(v690)+4))
	if v691 <= int32(0) {
		v723 = int32(0)
		goto L225
	} else {
		goto L229
	}
L227:
	;
	goto L228
L228:
	;
	v723 = int32(0)
	goto L225
L229:
	;
	v694 = int32(0)
	if v694 < v691 {
		goto L230
	} else {
		goto L231
	}
L230:
	;
	v697 = v691
	goto L232
L231:
	;
	v697 = v694
	goto L232
L232:
	;
	v698 = *(*int32)(unsafe.Add(mBase, uint32(v690)+12))
	v700 = int32(0)
	goto L233
L233:
	;
	v708 = *(*int32)(unsafe.Add(mBase, uint32(v698+v700<<(uint(int32(2))%32))))
	v709 = *(*int32)(unsafe.Add(mBase, uint32(v708)+4))
	if v709 == v594 {
		v723 = v708
		goto L225
	} else {
		goto L235
	}
L234:
	;
	goto L228
L235:
	;
	v712 = v700 + int32(1)
	if v712 != v697 {
		v700 = v712
		goto L233
	} else {
		goto L236
	}
L236:
	;
	goto L234
L237:
	;
	v725 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v723)+16)))
	v726 = v4 & v725
	*(*uint8)(unsafe.Add(mBase, uint32(v723)+16)) = uint8(v726)
	v728 = *(*int32)(unsafe.Add(mBase, uint32(v723)+8))
	if base.Ui32(v682) < base.Ui32(v728) {
		goto L240
	} else {
		goto L241
	}
L238:
	;
	goto L239
L239:
	;
	v737 = F_palloc0(m, int32(20))
	mBase = m.M
	v738 = m.ExcPending
	if v738 != 0 {
		goto L1
	} else {
		goto L246
	}
L240:
	;
	v730 = v728
	goto L242
L241:
	;
	v730 = v682
	goto L242
L242:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v723)+8)) = v730
	v732 = *(*int32)(unsafe.Add(mBase, uint32(v723)+12))
	if base.Ui32(v681) < base.Ui32(v732) {
		goto L243
	} else {
		goto L244
	}
L243:
	;
	v734 = v732
	goto L245
L244:
	;
	v734 = v681
	goto L245
L245:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v723)+12)) = v734
	goto L223
L246:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v737)+16)) = uint8(v4)
	*(*int32)(unsafe.Add(mBase, uint32(v737)+12)) = v681
	*(*int32)(unsafe.Add(mBase, uint32(v737)+8)) = v682
	*(*int32)(unsafe.Add(mBase, uint32(v737)+4)) = v594
	*(*int32)(unsafe.Add(mBase, uint32(v737))) = int32(109)
	v745 = *(*int32)(unsafe.Add(mBase, uint32(l1)+140))
	v746 = F_lappend(m, v745, v737)
	mBase = m.M
	v747 = m.ExcPending
	if v747 != 0 {
		goto L1
	} else {
		goto L247
	}
L247:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+140)) = v746
	goto L223
L248:
	;
	goto L187
L249:
	;
	goto L186
}
func F_transformOptionalSelectInto(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
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
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v5 != int32(141) {
		v8 = F_transformStmt(m, l0, l1)
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			return v8
		}
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
		if v13 != 0 {
			v16 = l1
			for {
				v18 = *(*int32)(unsafe.Add(mBase, uint32(v16)+76))
				v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+68))
				if v19 != 0 {
					v16 = v18
					continue
				} else {
					break
				}
				break
			}
			v22 = v18
		} else {
			v22 = l1
		}
		v24 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
		if v24 == int32(0) {
			v27 = F_transformStmt(m, l0, l1)
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return int32(0)
			} else {
				return v27
			}
		} else {
			v31 = F_palloc0(m, int32(20))
			mBase = m.M
			v32 = m.ExcPending
			if v32 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v31)+4)) = l1
				*(*int32)(unsafe.Add(mBase, uint32(v31))) = int32(242)
				v36 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
				v37 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(v31)+16)) = uint8(v37)
				*(*int32)(unsafe.Add(mBase, uint32(v31)+12)) = int32(42)
				*(*int32)(unsafe.Add(mBase, uint32(v31)+8)) = v36
				*(*int32)(unsafe.Add(mBase, uint32(v22)+8)) = int32(0)
				v44 = F_transformStmt(m, l0, v31)
				mBase = m.M
				v45 = m.ExcPending
				if v45 != 0 {
					return int32(0)
				} else {
					return v44
				}
			}
		}
	}
}
func F_transformWhereClause(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	if l1 == int32(0) {
		return int32(0)
	} else {
		v9 = F_transformExpr(m, l0, l1, l2)
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			v13 = F_coerce_to_boolean(m, l0, v9, l3)
			v14 = m.ExcPending
			if v14 != 0 {
				return int32(0)
			} else {
				return v13
			}
		}
	}
}
func F_transientrel_shutdown(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	F_FreeBulkInsertState(m, v4)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
		v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+188))
		if v8 == int32(0) {
			v18 = v7
			F_relation_close(m, v18, int32(0))
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(0)
				return
			}
		} else {
			v11 = *(*int32)(unsafe.Add(mBase, uint32(v8)+108))
			if v11 == int32(0) {
				v18 = v7
				F_relation_close(m, v18, int32(0))
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(0)
					return
				}
			} else {
				v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
				m.T0[v11].(func(*base.Module, int32, int32))(m, v7, v14)
				mBase = m.M
				v16 = m.ExcPending
				if v16 != 0 {
					return
				} else {
					v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
					v18 = v17
					F_relation_close(m, v18, int32(0))
					mBase = m.M
					v22 = m.ExcPending
					if v22 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(0)
						return
					}
				}
			}
		}
	}
}
func F_translate(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
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
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v119 int32
	_ = v119
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
	var v136 int64
	_ = v136
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v181 int32
	_ = v181
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v215 int32
	_ = v215
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v277 int32
	_ = v277
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v296 int32
	_ = v296
	var v307 int32
	_ = v307
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v324 int32
	_ = v324
	var v327 int32
	_ = v327
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v353 int32
	_ = v353
	var v380 int32
	_ = v380
	var v390 int32
	_ = v390
	var v399 int32
	_ = v399
	var v411 int32
	_ = v411
	var v433 int32
	_ = v433
	var v436 int32
	_ = v436
	var v440 int32
	_ = v440
	var v445 int32
	_ = v445
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v22 = F_pg_detoast_datum_packed(m, v21)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v27 = F_pg_detoast_datum_packed(m, v26)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v30 = F_pg_detoast_datum_packed(m, v29)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22))))
	if v32 == int32(1) {
		goto L9
	} else {
		goto L10
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v433 = m.ExcPending
	if v433 != 0 {
		goto L1
	} else {
		goto L102
	}
L6:
	;
	return base.I64_extend_i32_u(v411)
L7:
	;
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27))))
	if v66 == int32(1) {
		goto L19
	} else {
		goto L20
	}
L8:
	;
	if v61 <= int32(0) {
		v411 = v22
		goto L6
	} else {
		goto L17
	}
L9:
	;
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+1)))
	if base.Ui32((v36-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		v65 = int32(4)
		goto L7
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v48 = int32(1)
	if v32&v48 != 0 {
		v61 = int32(base.Ui32(v32)>>(uint(v48)%32)) - v48
		goto L8
	} else {
		goto L16
	}
L12:
	;
	if v36 == int32(18) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v47 = int32(16)
	goto L15
L14:
	;
	v47 = int32(0)
	goto L15
L15:
	;
	v61 = v47
	goto L8
L16:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	v61 = int32(base.Ui32(v54)>>(uint(int32(2))%32)) - int32(4)
	goto L8
L17:
	;
	v65 = v61
	goto L7
L18:
	;
	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30))))
	if v96 == int32(1) {
		goto L30
	} else {
		goto L31
	}
L19:
	;
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+1)))
	if v72 == int32(18) {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	goto L21
L21:
	;
	v83 = int32(1)
	if v66&v83 != 0 {
		v95 = int32(base.Ui32(v66)>>(uint(v83)%32)) - v83
		goto L18
	} else {
		goto L28
	}
L22:
	;
	v75 = int32(16)
	goto L24
L23:
	;
	v75 = int32(0)
	goto L24
L24:
	;
	if base.Ui32((v72-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v82 = int32(4)
	goto L27
L26:
	;
	v82 = v75
	goto L27
L27:
	;
	v95 = v82
	goto L18
L28:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
	v95 = int32(base.Ui32(v89)>>(uint(int32(2))%32)) - int32(4)
	goto L18
L29:
	;
	v127 = *(*int32)(unsafe.Add(mBase, _c_F_translate[0]))
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v127)+4))
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v128*int32(28))+uint32(_c_F_translate[1])))
	goto L40
L30:
	;
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+1)))
	if v102 == int32(18) {
		goto L33
	} else {
		goto L34
	}
L31:
	;
	goto L32
L32:
	;
	v113 = int32(1)
	if v96&v113 != 0 {
		v125 = int32(base.Ui32(v96)>>(uint(v113)%32)) - v113
		goto L29
	} else {
		goto L39
	}
L33:
	;
	v105 = int32(16)
	goto L35
L34:
	;
	v105 = int32(0)
	goto L35
L35:
	;
	if base.Ui32((v102-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v112 = int32(4)
	goto L38
L37:
	;
	v112 = v105
	goto L38
L38:
	;
	v125 = v112
	goto L29
L39:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	v125 = int32(base.Ui32(v119)>>(uint(int32(2))%32)) - int32(4)
	goto L29
L40:
	;
	v136 = base.I64_extend_i32_s(v133) * base.I64_extend_i32_s(v65)
	v140 = base.I32_wrap_i64(v136)
	if base.I32_wrap_i64(int64(base.Ui64(v136)>>(uint(int64(32))%64))) != v140>>(uint(int32(31))%32) {
		goto L5
	} else {
		goto L41
	}
L41:
	;
	v145 = v140 + int32(4)
	if base.B2i32(v145 < v140)|base.B2i32(base.Ui32(int32(1073741824)) <= base.Ui32(v145)) != 0 {
		goto L5
	} else {
		goto L42
	}
L42:
	;
	v150 = int32(1)
	if v32&v150 != 0 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v154 = v150
	goto L45
L44:
	;
	v154 = int32(4)
	goto L45
L45:
	;
	v155 = v22 + v154
	v157 = int32(1)
	if v66&v157 != 0 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v161 = v157
	goto L48
L47:
	;
	v161 = int32(4)
	goto L48
L48:
	;
	v162 = v27 + v161
	v164 = int32(1)
	if v96&v164 != 0 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v168 = v164
	goto L51
L50:
	;
	v168 = int32(4)
	goto L51
L51:
	;
	v169 = v30 + v168
	v170 = v169 + v125
	v172 = base.B2i32(int32(0) < v125)
	v173 = F_palloc(m, v145)
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	v181 = v173 + int32(4)
	v186 = v65
	v187 = v155
	v191 = int32(0)
	goto L53
L53:
	;
	v199 = F_pg_mblen_range(m, v187, v155+v65)
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L1
	} else {
		goto L55
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v173))) = v390<<(uint(int32(2))%32) + int32(16)
	v411 = v173
	goto L6
L55:
	;
	v201 = int32(0)
	if base.B2i32(v95 <= int32(0)) == v201 {
		goto L57
	} else {
		goto L58
	}
L56:
	;
	v399 = v186 - v199
	if int32(0) < v399 {
		v181 = v380
		v186 = v399
		v187 = v199 + v187
		v191 = v390
		goto L53
	} else {
		goto L101
	}
L57:
	;
	v205 = v201
	v215 = v201
	goto L60
L58:
	;
	goto L59
L59:
	;
	if v199 != 0 {
		goto L98
	} else {
		goto L99
	}
L60:
	;
	v225 = v205 + v162
	v226 = F_pg_mblen_range(m, v225, v162+v95)
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L1
	} else {
		goto L63
	}
L61:
	;
	goto L59
L62:
	;
	v353 = v205 + v226
	if v353 < v95 {
		v205 = v353
		v215 = v215 + int32(1)
		goto L60
	} else {
		goto L97
	}
L63:
	;
	if v226 != v199 {
		goto L62
	} else {
		goto L64
	}
L64:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v199) {
		goto L68
	} else {
		goto L69
	}
L65:
	;
	if v290 != 0 {
		goto L62
	} else {
		goto L83
	}
L66:
	;
	v290 = int32(0)
	goto L65
L67:
	;
	v264 = v259
	v265 = v260
	v266 = v261
	goto L77
L68:
	;
	if (v187|v225)&int32(3) != 0 {
		v259 = v187
		v260 = v225
		v261 = v199
		goto L67
	} else {
		goto L71
	}
L69:
	;
	v252 = v187
	v253 = v225
	v254 = v199
	goto L70
L70:
	;
	if v254 == int32(0) {
		goto L66
	} else {
		goto L76
	}
L71:
	;
	v236 = v187
	v237 = v225
	v238 = v199
	goto L72
L72:
	;
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v236)))
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v237)))
	if v241 != v242 {
		v259 = v236
		v260 = v237
		v261 = v238
		goto L67
	} else {
		goto L74
	}
L73:
	;
	v252 = v247
	v253 = v245
	v254 = v249
	goto L70
L74:
	;
	v244 = int32(4)
	v245 = v237 + v244
	v247 = v236 + v244
	v249 = v238 - v244
	if base.Ui32(int32(3)) < base.Ui32(v249) {
		v236 = v247
		v237 = v245
		v238 = v249
		goto L72
	} else {
		goto L75
	}
L75:
	;
	goto L73
L76:
	;
	v259 = v252
	v260 = v253
	v261 = v254
	goto L67
L77:
	;
	v269 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v264))))
	v270 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v265))))
	if v269 == v270 {
		goto L79
	} else {
		goto L80
	}
L78:
	;
	v290 = v269 - v270
	goto L65
L79:
	;
	v272 = int32(1)
	v277 = v266 - v272
	if v277 != 0 {
		v264 = v264 + v272
		v265 = v265 + v272
		v266 = v277
		goto L77
	} else {
		goto L82
	}
L80:
	;
	goto L81
L81:
	;
	goto L78
L82:
	;
	goto L66
L83:
	;
	if v215 <= int32(0) {
		v324 = v169
		v327 = v172
		goto L84
	} else {
		goto L85
	}
L84:
	;
	if v327 == int32(0) {
		v380 = v181
		v390 = v191
		goto L56
	} else {
		goto L92
	}
L85:
	;
	v293 = int32(0)
	if v125 <= v293 {
		v324 = v169
		v327 = v172
		goto L84
	} else {
		goto L86
	}
L86:
	;
	v296 = v169
	v307 = v293
	goto L87
L87:
	;
	v316 = F_pg_mblen_range(m, v296, v170)
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L1
	} else {
		goto L89
	}
L88:
	;
	v324 = v318
	v327 = v319
	goto L84
L89:
	;
	v318 = v316 + v296
	v319 = base.B2i32(base.Ui32(v318) < base.Ui32(v170))
	v321 = v307 + int32(1)
	if v215 <= v321 {
		v324 = v318
		v327 = v319
		goto L84
	} else {
		goto L90
	}
L90:
	;
	if base.Ui32(v318) < base.Ui32(v170) {
		v296 = v318
		v307 = v321
		goto L87
	} else {
		goto L91
	}
L91:
	;
	goto L88
L92:
	;
	v346 = F_pg_mblen_range(m, v324, v170)
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L1
	} else {
		goto L93
	}
L93:
	;
	if v346 != 0 {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	base.MemoryCopy(m, v181, v324, v346)
	goto L96
L95:
	;
	goto L96
L96:
	;
	v380 = v181 + v346
	v390 = v346 + v191
	goto L56
L97:
	;
	goto L61
L98:
	;
	base.MemoryCopy(m, v181, v187, v199)
	goto L100
L99:
	;
	goto L100
L100:
	;
	v380 = v199 + v181
	v390 = v199 + v191
	goto L56
L101:
	;
	goto L54
L102:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v436 = m.ExcPending
	if v436 != 0 {
		goto L1
	} else {
		goto L103
	}
L103:
	;
	F_errmsg(m, int32(_a_F_translate_0), int32(0))
	mBase = m.M
	v440 = m.ExcPending
	if v440 != 0 {
		goto L1
	} else {
		goto L104
	}
L104:
	;
	F_errfinish(m, int32(_a_F_translate_1), int32(864), int32(_a_F_translate_2))
	mBase = m.M
	v445 = m.ExcPending
	if v445 != 0 {
		goto L1
	} else {
		goto L105
	}
L105:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_truncate_identifier(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	if int32(64) <= l1 {
		v12 = F_pg_mbcliplen(m, l0, l1, int32(63))
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return
		} else {
			if l2 == int32(0) {
				v37 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(l0+v12))) = uint8(v37)
				m.G0 = v7 + int32(16)
				return
			} else {
				v18 = F_errstart(m, int32(18), int32(0))
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return
				} else {
					if v18 == int32(0) {
						v37 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(l0+v12))) = uint8(v37)
						m.G0 = v7 + int32(16)
						return
					} else {
						F_errcode(m, int32(34103428))
						mBase = m.M
						v24 = m.ExcPending
						if v24 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = l0
							*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v12
							*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
							F_errmsg(m, int32(_a_F_truncate_identifier_0), v7)
							mBase = m.M
							v30 = m.ExcPending
							if v30 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_truncate_identifier_1), int32(102), int32(_a_F_truncate_identifier_2))
								mBase = m.M
								v35 = m.ExcPending
								if v35 != 0 {
									return
								} else {
									v37 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(l0+v12))) = uint8(v37)
									m.G0 = v7 + int32(16)
									return
								}
							}
						}
					}
				}
			}
		}
	} else {
		m.G0 = v7 + int32(16)
		return
	}
}
func F_tsm_handler_out(m *base.Module, l0 int32) int64 {
	var v7 int64
	_ = v7
	var v10 int32
	_ = v10
	v7 = Fn14235(m, l0, int32(_a_F_tsm_handler_out_0), int32(372), int32(_a_F_tsm_handler_out_1), int32(_a_F_tsm_handler_out_2), int32(_a_F_tsm_handler_out_3))
	v10 = m.ExcPending
	if v10 != 0 {
		return int64(0)
	} else {
		return v7
	}
}
func F_tsqueryin(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v4 = int32(0)
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v7 = F_parse_tsquery(m, v2, int32(1723), v4, v4, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int64(0)
	} else {
		return base.I64_extend_i32_u(v7)
	}
}
func F_tsvectorout(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v162 int32
	_ = v162
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v262 int32
	_ = v262
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v287 int32
	_ = v287
	var v299 int32
	_ = v299
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v311 int32
	_ = v311
	var v324 int32
	_ = v324
	var v326 int32
	_ = v326
	var v328 int32
	_ = v328
	var v337 int32
	_ = v337
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v353 int32
	_ = v353
	var v357 int32
	_ = v357
	var v361 int32
	_ = v361
	var v365 int32
	_ = v365
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v372 int32
	_ = v372
	var v374 int32
	_ = v374
	var v388 int32
	_ = v388
	var v391 int32
	_ = v391
	var v404 int32
	_ = v404
	var v406 int32
	_ = v406
	var v409 int32
	_ = v409
	v2 = int32(0)
	v15 = m.G0
	v17 = v15 - int32(16)
	m.G0 = v17
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v20 = F_pg_detoast_datum(m, v19)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v25 = v20 + int32(8)
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	v30 = v26*int32(3) + int32(1)
	if int32(0) < v26 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v37 = v30
	v38 = v2
	goto L6
L4:
	;
	v82 = v30
	goto L5
L5:
	;
	v95 = F_palloc(m, v82)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L1
	} else {
		goto L12
	}
L6:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v25+v38<<(uint(int32(2))%32))))
	v54 = int32(1)
	v57 = int32(base.Ui32(v53)>>(uint(v54)%32)) & int32(2047)
	v60 = v57<<(uint(v54)%32) + v37
	if v53&v54 != 0 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v82 = v77
	goto L5
L8:
	;
	v66 = int32(1)
	v71 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v25+v26<<(uint(int32(2))%32)+(int32(base.Ui32(v53)>>(uint(int32(12))%32))+v57+v66)&int32(_a_F_tsvectorout_0)))))
	v77 = v60 + v71*int32(7) + v66
	goto L10
L9:
	;
	v77 = v60
	goto L10
L10:
	;
	v79 = v38 + int32(1)
	if v79 != v26 {
		v37 = v77
		v38 = v79
		goto L6
	} else {
		goto L11
	}
L11:
	;
	goto L7
L12:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	if int32(0) < v97 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v101 = v95
	v103 = v97
	v106 = v25
	v110 = v2
	goto L16
L14:
	;
	v391 = v95
	goto L15
L15:
	;
	v404 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v391))) = uint8(v404)
	v406 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v406 != v20 {
		goto L58
	} else {
		goto L59
	}
L16:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v106)))
	v118 = int32(base.Ui32(v114)>>(uint(int32(1))%32)) & int32(2047)
	if v110 != 0 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v391 = v372
	goto L15
L18:
	;
	v119 = int32(32)
	*(*uint8)(unsafe.Add(mBase, uint32(v101))) = uint8(v119)
	v123 = v101 + int32(1)
	goto L20
L19:
	;
	v123 = v101
	goto L20
L20:
	;
	v124 = int32(39)
	*(*uint8)(unsafe.Add(mBase, uint32(v123))) = uint8(v124)
	v127 = v123 + int32(1)
	if v118 != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v133 = v25 + v103<<(uint(int32(2))%32) + int32(base.Ui32(v114)>>(uint(int32(12))%32))
	v134 = v133 + v118
	v136 = v133
	v137 = v127
	goto L24
L22:
	;
	v262 = v127
	goto L23
L23:
	;
	v274 = int32(39)
	*(*uint8)(unsafe.Add(mBase, uint32(v262))) = uint8(v274)
	v276 = int32(1)
	v277 = v262 + v276
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v106)))
	if v279&v276 == int32(0) {
		v372 = v277
		v374 = v278
		goto L43
	} else {
		goto L44
	}
L24:
	;
	v149 = F_pg_mblen_range(m, v136, v134)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L1
	} else {
		goto L26
	}
L25:
	;
	v262 = v247
	goto L23
L26:
	;
	v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v136))))
	if base.B2i32(v151 != int32(92))&base.B2i32(v151 != int32(39)) == int32(0) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v137))) = uint8(v151)
	v162 = v137 + int32(1)
	goto L29
L28:
	;
	v162 = v137
	goto L29
L29:
	;
	if v149 == int32(0) {
		v246 = v136
		v247 = v162
		goto L30
	} else {
		goto L31
	}
L30:
	;
	if base.Ui32(v246) < base.Ui32(v134) {
		v136 = v246
		v137 = v247
		goto L24
	} else {
		goto L42
	}
L31:
	;
	v167 = v149 & int32(7)
	if v167 != 0 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v169 = v136
	v170 = v162
	v171 = v149
	v173 = int32(0)
	goto L35
L33:
	;
	v194 = v136
	v195 = v162
	v196 = v149
	goto L34
L34:
	;
	if base.Ui32(v149) < base.Ui32(int32(8)) {
		v246 = v194
		v247 = v195
		goto L30
	} else {
		goto L38
	}
L35:
	;
	v182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v169))))
	*(*uint8)(unsafe.Add(mBase, uint32(v170))) = uint8(v182)
	v184 = int32(1)
	v185 = v170 + v184
	v187 = v169 + v184
	v189 = v171 - v184
	v191 = v173 + v184
	if v191 != v167 {
		v169 = v187
		v170 = v185
		v171 = v189
		v173 = v191
		goto L35
	} else {
		goto L37
	}
L36:
	;
	v194 = v187
	v195 = v185
	v196 = v189
	goto L34
L37:
	;
	goto L36
L38:
	;
	v210 = v194
	v211 = v195
	v212 = v196
	goto L39
L39:
	;
	v223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v210))))
	*(*uint8)(unsafe.Add(mBase, uint32(v211))) = uint8(v223)
	v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v210)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v211)+1)) = uint8(v225)
	v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v210)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v211)+2)) = uint8(v227)
	v229 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v210)+3)))
	*(*uint8)(unsafe.Add(mBase, uint32(v211)+3)) = uint8(v229)
	v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v210)+4)))
	*(*uint8)(unsafe.Add(mBase, uint32(v211)+4)) = uint8(v231)
	v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v210)+5)))
	*(*uint8)(unsafe.Add(mBase, uint32(v211)+5)) = uint8(v233)
	v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v210)+6)))
	*(*uint8)(unsafe.Add(mBase, uint32(v211)+6)) = uint8(v235)
	v237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v210)+7)))
	*(*uint8)(unsafe.Add(mBase, uint32(v211)+7)) = uint8(v237)
	v239 = int32(8)
	v240 = v211 + v239
	v242 = v210 + v239
	v244 = v212 - v239
	if v244 != 0 {
		v210 = v242
		v211 = v240
		v212 = v244
		goto L39
	} else {
		goto L41
	}
L40:
	;
	v246 = v242
	v247 = v240
	goto L30
L41:
	;
	goto L40
L42:
	;
	goto L25
L43:
	;
	v388 = v110 + int32(1)
	if v388 < v374 {
		v101 = v372
		v103 = v374
		v106 = v106 + int32(4)
		v110 = v388
		goto L16
	} else {
		goto L57
	}
L44:
	;
	v287 = int32(1)
	v299 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v25+v278<<(uint(int32(2))%32)+(int32(base.Ui32(v279)>>(uint(v287)%32))&int32(2047)+int32(base.Ui32(v279)>>(uint(int32(12))%32))+v287)&int32(_a_F_tsvectorout_0)))))
	if v299 == int32(0) {
		v372 = v277
		v374 = v278
		goto L43
	} else {
		goto L45
	}
L45:
	;
	v302 = int32(58)
	*(*uint8)(unsafe.Add(mBase, uint32(v262)+1)) = uint8(v302)
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	v305 = int32(2)
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v106)))
	v311 = int32(1)
	v324 = v262 + v305
	v326 = v25 + v304<<(uint(v305)%32) + (int32(base.Ui32(v308)>>(uint(int32(12))%32))+int32(base.Ui32(v308)>>(uint(v311)%32))&int32(2047)+v311)&int32(_a_F_tsvectorout_0)
	v328 = v299
	goto L46
L46:
	;
	v337 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v326)+2)))
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v337 & int32(_a_F_tsvectorout_1)
	v342 = F_pg_sprintf(m, v324, int32(_a_F_tsvectorout_2), v17)
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L1
	} else {
		goto L48
	}
L47:
	;
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	v372 = v365
	v374 = v370
	goto L43
L48:
	;
	v344 = v342 + v324
	v346 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v326)+2)))
	switch int32(base.Ui32(v346)>>(uint(int32(14))%32)) - int32(1) {
	case 0:
		goto L51
	case 1:
		goto L52
	case 2:
		v353 = int32(65)
		goto L50
	default:
		v357 = v344
		goto L49
	}
L49:
	;
	if int32(2) <= v328 {
		goto L53
	} else {
		goto L54
	}
L50:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v344))) = uint8(v353)
	v357 = v344 + int32(1)
	goto L49
L51:
	;
	v353 = int32(67)
	goto L50
L52:
	;
	v353 = int32(66)
	goto L50
L53:
	;
	v361 = int32(44)
	*(*uint8)(unsafe.Add(mBase, uint32(v357))) = uint8(v361)
	v365 = v357 + int32(1)
	goto L55
L54:
	;
	v365 = v357
	goto L55
L55:
	;
	v369 = v328 - int32(1)
	if v369 != 0 {
		v324 = v365
		v326 = v326 + int32(2)
		v328 = v369
		goto L46
	} else {
		goto L56
	}
L56:
	;
	goto L47
L57:
	;
	goto L17
L58:
	;
	F_pfree(m, v20)
	mBase = m.M
	v409 = m.ExcPending
	if v409 != 0 {
		goto L1
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	m.G0 = v17 + int32(16)
	return base.I64_extend_i32_u(v95)
L61:
	;
	goto L60
}
func F_tsvectorsend(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
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
	var v31 int32
	_ = v31
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v140 int32
	_ = v140
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v170 int32
	_ = v170
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v16 = F_pg_detoast_datum(m, v15)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	F_pq_begintypsend(m, v13)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	F_enlargeStringInfo(m, v13, int32(4))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	v31 = int32(16711935)
	*(*int32)(unsafe.Add(mBase, uint32(v26+v27))) = base.I32_rotr(v22, int32(24))&v31 | base.I32_rotr(v22&v31, int32(8))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v26 + int32(4)
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	if int32(0) < v42 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v46 = v16 + int32(8)
	v47 = v42
	v49 = v46
	v56 = int32(0)
	goto L8
L6:
	;
	goto L7
L7:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v199))) = v200 << (uint(int32(2)) % 32)
	goto L24
L8:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
	F_pq_sendtext(m, v13, v46+v47<<(uint(int32(2))%32)+int32(base.Ui32(v60)>>(uint(int32(12))%32)), int32(base.Ui32(v60)>>(uint(int32(1))%32))&int32(2047))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L1
	} else {
		goto L10
	}
L9:
	;
	goto L7
L10:
	;
	F_enlargeStringInfo(m, v13, int32(1))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v73 = int32(0)
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	*(*uint8)(unsafe.Add(mBase, uint32(v74+v75))) = uint8(v73)
	v79 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v74 + v79
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
	if v82&v79 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	v89 = int32(1)
	v101 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v46+v85<<(uint(int32(2))%32)+(int32(base.Ui32(v82)>>(uint(v89)%32))&int32(2047)+int32(base.Ui32(v82)>>(uint(int32(12))%32))+v89)&int32(_a_F_tsvectorsend_0)))))
	v102 = v101
	goto L14
L13:
	;
	v102 = v73
	goto L14
L14:
	;
	F_enlargeStringInfo(m, v13, int32(2))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	v109 = int32(8)
	v112 = v102 & int32(_a_F_tsvectorsend_1)
	v115 = v102<<(uint(v109)%32) | int32(base.Ui32(v112)>>(uint(v109)%32))
	*(*uint16)(unsafe.Add(mBase, uint32(v106+v107))) = uint16(v115)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v106 + int32(2)
	if v112 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	v121 = int32(2)
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
	v127 = int32(1)
	v140 = int32(0)
	goto L19
L17:
	;
	goto L18
L18:
	;
	v185 = v56 + int32(1)
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	if v185 < v186 {
		v47 = v186
		v49 = v49 + int32(4)
		v56 = v185
		goto L8
	} else {
		goto L23
	}
L19:
	;
	v153 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v46+v120<<(uint(v121)%32)+(int32(base.Ui32(v124)>>(uint(int32(12))%32))+int32(base.Ui32(v124)>>(uint(v127)%32))&int32(2047)+v127)&int32(_a_F_tsvectorsend_0)+v121+v140<<(uint(int32(1))%32)))))
	F_enlargeStringInfo(m, v13, int32(2))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L1
	} else {
		goto L21
	}
L20:
	;
	goto L18
L21:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	v160 = int32(8)
	v164 = v153<<(uint(v160)%32) | int32(base.Ui32(v153)>>(uint(v160)%32))
	*(*uint16)(unsafe.Add(mBase, uint32(v157+v158))) = uint16(v164)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v157 + int32(2)
	v170 = v140 + int32(1)
	if v170 != v112 {
		v140 = v170
		goto L19
	} else {
		goto L22
	}
L22:
	;
	goto L20
L23:
	;
	goto L9
L24:
	;
	m.G0 = v13 + int32(16)
	return base.I64_extend_i32_u(v199)
}
func F_tuples_equal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int64
	_ = v96
	var v97 int32
	_ = v97
	var v99 int64
	_ = v99
	var v100 int64
	_ = v100
	var v101 int32
	_ = v101
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v122 int32
	_ = v122
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	v17 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+6)))
	if v17 < v16 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+16))
	m.T0[v20].(func(*base.Module, int32, int32))(m, l0, v16)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	v27 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+6)))
	if v27 < v26 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	return int32(0)
L5:
	;
	goto L3
L6:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+16))
	m.T0[v30].(func(*base.Module, int32, int32))(m, l1, v26)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L4
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
	if v35 <= int32(0) {
		v122 = int32(1)
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L8
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L4
	} else {
		goto L33
	}
L11:
	;
	m.G0 = v13 + int32(16)
	return v122
L12:
	;
	v43 = v35
	v44 = int32(0)
	v45 = v34
	goto L13
L13:
	;
	v54 = v45 + v43<<(uint(int32(3))%32) + v44*int32(100)
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54)+119)))
	if v55 != 0 {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v122 = v108
	goto L11
L15:
	;
	v108 = int32(1)
	v110 = v44 + v108
	v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v111)))
	if v110 < v112 {
		v43 = v112
		v44 = v110
		v45 = v111
		goto L13
	} else {
		goto L32
	}
L16:
	;
	v57 = v54 + int32(28)
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57)+90)))
	if v58 != 0 {
		goto L15
	} else {
		goto L17
	}
L17:
	;
	if l3 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v59 = int32(*(*int16)(unsafe.Add(mBase, uint32(v57)+74)))
	v62 = F_bms_is_member(m, v59+int32(7), l3)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L4
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v66 = int32(0)
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67+v44))))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70+v44))))
	if v69 != v72 {
		v122 = v66
		goto L11
	} else {
		goto L23
	}
L21:
	;
	if v62 == int32(0) {
		goto L15
	} else {
		goto L22
	}
L22:
	;
	goto L20
L23:
	;
	if v69 != 0 {
		goto L15
	} else {
		goto L24
	}
L24:
	;
	v76 = l2 + v44<<(uint(int32(2))%32)
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)))
	if v77 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v57)+68))
	v82 = F_lookup_type_cache(m, v80, int32(32))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L4
	} else {
		goto L28
	}
L26:
	;
	v88 = v77
	goto L27
L27:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v57)+96))
	v93 = v44 << (uint(int32(3)) % 32)
	v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v96 = *(*int64)(unsafe.Add(mBase, uint32(v93+v94)))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v99 = *(*int64)(unsafe.Add(mBase, uint32(v97+v93)))
	v100 = F_FunctionCall2Coll(m, v88+int32(76), v91, v96, v99)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L4
	} else {
		goto L30
	}
L28:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v82)+80))
	if v84 == int32(0) {
		goto L10
	} else {
		goto L29
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v76))) = v82
	v88 = v82
	goto L27
L30:
	;
	if v100 == int64(0) {
		v122 = v66
		goto L11
	} else {
		goto L31
	}
L31:
	;
	goto L15
L32:
	;
	goto L14
L33:
	;
	F_errcode(m, int32(52461700))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L4
	} else {
		goto L34
	}
L34:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v57)+68))
	v136 = F_format_type_be(m, v135)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L4
	} else {
		goto L35
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v136
	F_errmsg(m, int32(_a_F_tuples_equal_0), v13)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L4
	} else {
		goto L36
	}
L36:
	;
	F_errfinish(m, int32(_a_F_tuples_equal_1), int32(345), int32(_a_F_tuples_equal_2))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L4
	} else {
		goto L37
	}
L37:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
