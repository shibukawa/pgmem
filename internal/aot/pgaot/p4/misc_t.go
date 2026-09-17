package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_TerminateLocalBufferIO(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v14 int32
	_ = v14
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v7 = v5 & int32(-134217729)
	if l2 != 0 {
		*(*int32)(unsafe.Add(mBase, uint32(l0+int32(36)))) = int32(-1)
		v14 = v7 - int32(1)
	} else {
		v14 = v7
	}
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v14 | l1
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
	var v232 int32
	_ = v232
	var v271 int32
	_ = v271
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v285 int32
	_ = v285
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v313 int32
	_ = v313
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v345 int32
	_ = v345
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v370 int32
	_ = v370
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v380 int32
	_ = v380
	var v408 int32
	_ = v408
	var v410 int32
	_ = v410
	var v412 int32
	_ = v412
	var v419 int32
	_ = v419
	var v421 int32
	_ = v421
	var v423 int32
	_ = v423
	var v425 int32
	_ = v425
	var v427 int32
	_ = v427
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v438 int32
	_ = v438
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v479 int32
	_ = v479
	var v481 int32
	_ = v481
	var v483 int32
	_ = v483
	var v526 int32
	_ = v526
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v543 int32
	_ = v543
	var v571 int32
	_ = v571
	var v573 int32
	_ = v573
	var v575 int32
	_ = v575
	var v582 int32
	_ = v582
	var v584 int32
	_ = v584
	var v586 int32
	_ = v586
	var v588 int32
	_ = v588
	var v590 int32
	_ = v590
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v601 int32
	_ = v601
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v642 int32
	_ = v642
	var v644 int32
	_ = v644
	var v646 int32
	_ = v646
	var v689 int32
	_ = v689
	var v692 int32
	_ = v692
	var v694 int32
	_ = v694
	var v695 int32
	_ = v695
	var v696 int32
	_ = v696
	var v702 int32
	_ = v702
	var v703 int32
	_ = v703
	var v745 int32
	_ = v745
	var v788 int32
	_ = v788
	var v798 int32
	_ = v798
	var v824 int32
	_ = v824
	var v825 int32
	_ = v825
	var v859 int32
	_ = v859
	var v866 int32
	_ = v866
	var v873 int32
	_ = v873
	var v909 int32
	_ = v909
	var v911 int32
	_ = v911
	var v913 int32
	_ = v913
	var v921 int32
	_ = v921
	var v922 int32
	_ = v922
	var v923 int32
	_ = v923
	var v926 int32
	_ = v926
	var v934 int32
	_ = v934
	var v963 int32
	_ = v963
	var v964 int32
	_ = v964
	var v968 int32
	_ = v968
	var v979 int32
	_ = v979
	var v981 int32
	_ = v981
	var v984 int32
	_ = v984
	var v988 int32
	_ = v988
	var v1025 int32
	_ = v1025
	var v1028 int32
	_ = v1028
	var v1031 int32
	_ = v1031
	var v1032 int32
	_ = v1032
	var v1038 int32
	_ = v1038
	var v1077 int32
	_ = v1077
	var v1119 int32
	_ = v1119
	var v1142 int32
	_ = v1142
	var v1149 int32
	_ = v1149
	var v1161 int32
	_ = v1161
	var v1166 int32
	_ = v1166
	var v1169 int32
	_ = v1169
	var v1202 int32
	_ = v1202
	var v1204 int32
	_ = v1204
	var v1206 int32
	_ = v1206
	var v1208 int32
	_ = v1208
	var v1217 int32
	_ = v1217
	var v1221 int32
	_ = v1221
	var v1222 int32
	_ = v1222
	var v1223 int32
	_ = v1223
	var v1225 int32
	_ = v1225
	var v1227 int32
	_ = v1227
	var v1230 int32
	_ = v1230
	var v1240 int32
	_ = v1240
	var v1247 int32
	_ = v1247
	var v1264 int32
	_ = v1264
	var v1267 int32
	_ = v1267
	var v1270 int32
	_ = v1270
	var v1274 int32
	_ = v1274
	var v1278 int32
	_ = v1278
	var v1280 int32
	_ = v1280
	var v1289 int32
	_ = v1289
	var v1294 int32
	_ = v1294
	var v1298 int32
	_ = v1298
	var v1300 int32
	_ = v1300
	var v1305 int32
	_ = v1305
	var v1322 int32
	_ = v1322
	var v1325 int32
	_ = v1325
	var v1327 int32
	_ = v1327
	var v1329 int32
	_ = v1329
	var v1331 int32
	_ = v1331
	var v1333 int32
	_ = v1333
	var v1342 int32
	_ = v1342
	var v1346 int32
	_ = v1346
	var v1347 int32
	_ = v1347
	var v1348 int32
	_ = v1348
	var v1350 int32
	_ = v1350
	var v1352 int32
	_ = v1352
	var v1355 int32
	_ = v1355
	var v1365 int32
	_ = v1365
	var v1372 int32
	_ = v1372
	var v1389 int32
	_ = v1389
	var v1392 int32
	_ = v1392
	var v1395 int32
	_ = v1395
	var v1399 int32
	_ = v1399
	var v1403 int32
	_ = v1403
	var v1405 int32
	_ = v1405
	var v1414 int32
	_ = v1414
	var v1419 int32
	_ = v1419
	var v1423 int32
	_ = v1423
	var v1425 int32
	_ = v1425
	var v1430 int32
	_ = v1430
	var v1447 int32
	_ = v1447
	var v1450 int32
	_ = v1450
	var v1452 int32
	_ = v1452
	var v1454 int32
	_ = v1454
	var v1456 int32
	_ = v1456
	var v1462 int32
	_ = v1462
	var v1465 int32
	_ = v1465
	var v1495 int32
	_ = v1495
	var v1504 int32
	_ = v1504
	var v1508 int32
	_ = v1508
	var v1509 int32
	_ = v1509
	var v1510 int32
	_ = v1510
	var v1512 int32
	_ = v1512
	var v1514 int32
	_ = v1514
	var v1517 int32
	_ = v1517
	var v1527 int32
	_ = v1527
	var v1534 int32
	_ = v1534
	var v1551 int32
	_ = v1551
	var v1554 int32
	_ = v1554
	var v1557 int32
	_ = v1557
	var v1561 int32
	_ = v1561
	var v1565 int32
	_ = v1565
	var v1567 int32
	_ = v1567
	var v1576 int32
	_ = v1576
	var v1581 int32
	_ = v1581
	var v1585 int32
	_ = v1585
	var v1587 int32
	_ = v1587
	var v1592 int32
	_ = v1592
	var v1609 int32
	_ = v1609
	var v1612 int32
	_ = v1612
	var v1614 int32
	_ = v1614
	var v1619 int32
	_ = v1619
	var v1622 int32
	_ = v1622
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
	m.G0 = v1622 + int32(16)
	return v1619
L2:
	;
	v1619 = int32(-1)
	v1622 = v39
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
		v1462 = v2
		v1465 = v39
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v1495 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_TestConfiguration[6])) = v1495
	*(*int32)(unsafe.Add(mBase, _c_F_TestConfiguration[7])) = v1495
	*(*int32)(unsafe.Add(mBase, uint32(v1465)+12)) = v1495
	v1504 = v1465 + int32(12)
	v1508 = *(*int32)(unsafe.Add(mBase, uint32(l0)+616))
	if v1508 != 0 {
		goto L220
	} else {
		goto L221
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
	v1161 = int32(0)
	if v59 <= v1161 {
		v1462 = v1161
		v1465 = v39
		goto L5
	} else {
		goto L147
	}
L9:
	;
	if int32(0) < v89 {
		__phi89 = v89 - int32(1)
		__phi100 = v89
		__phi109 = v1142
		__phi116 = v1149
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
	v1142 = v109
	v1149 = v116
	goto L9
L16:
	;
	v271 = v181 << (uint(int32(2)) % 32)
	if v75&v72 != 0 {
		v285 = v271
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
	*(*int32)(unsafe.Add(mBase, uint32(v83+v192<<(uint(int32(2))%32)))) = v191
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v191)+4))
	if v232 != v188 {
		v191 = v232
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
	v313 = v181 - int32(1)
	if int32(0) < v100 {
		goto L42
	} else {
		goto L43
	}
L23:
	;
	if v304 == int32(0) {
		goto L22
	} else {
		goto L41
	}
L24:
	;
	v301 = int32(0)
	if v73 == v301 {
		goto L22
	} else {
		goto L40
	}
L25:
	;
	if v285 != 0 {
		goto L32
	} else {
		goto L33
	}
L26:
	;
	if base.Ui32(int32(1024)) < base.Ui32(v271) {
		v285 = v271
		goto L25
	} else {
		goto L27
	}
L27:
	;
	if v271 == int32(0) {
		goto L24
	} else {
		goto L28
	}
L28:
	;
	v276 = v271 + v75
	if base.Ui32(v79) < base.Ui32(v276) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v278 = v276
	goto L31
L30:
	;
	v278 = v79
	goto L31
L31:
	;
	v285 = (v278+(v75^v68))&int32(-4) + int32(4)
	goto L25
L32:
	;
	base.MemoryFill(m, v75, int32(0), v285)
	goto L34
L33:
	;
	goto L34
L34:
	;
	if v73|base.B2i32(base.Ui32(int32(1025)) <= base.Ui32(v271)) != 0 {
		v304 = v271
		goto L23
	} else {
		goto L35
	}
L35:
	;
	if v271 == int32(0) {
		goto L22
	} else {
		goto L36
	}
L36:
	;
	v293 = v271 + v67
	if base.Ui32(v71) < base.Ui32(v293) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v295 = v293
	goto L39
L38:
	;
	v295 = v71
	goto L39
L39:
	;
	v304 = (v295+(v67^v68))&int32(-4) + int32(4)
	goto L23
L40:
	;
	v304 = v301
	goto L23
L41:
	;
	base.MemoryFill(m, v67, int32(0), v304)
	goto L22
L42:
	;
	v317 = v181 & int32(-2)
	v318 = int32(1)
	v319 = v181 & v318
	v345 = int32(0)
	goto L45
L43:
	;
	goto L44
L44:
	;
	if int32(0) <= v313 {
		goto L108
	} else {
		goto L109
	}
L45:
	;
	if v313 < int32(0) {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	goto L44
L47:
	;
	v745 = v345 + int32(1)
	if v745 != v100 {
		v345 = v745
		goto L45
	} else {
		goto L107
	}
L48:
	;
	v367 = v65 + v345*int32(20)
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v367)))
	v370 = int32(-1)
	if v313 != 0 {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	if v526 < int32(0) {
		goto L47
	} else {
		goto L77
	}
L50:
	;
	v372 = v313
	v373 = v370
	v380 = int32(0)
	goto L53
L51:
	;
	v443 = v313
	v444 = v370
	goto L52
L52:
	;
	v479 = v443 << (uint(int32(2)) % 32)
	v481 = *(*int32)(unsafe.Add(mBase, uint32(v83+v479)))
	if v368 != v481 {
		goto L72
	} else {
		goto L73
	}
L53:
	;
	v408 = v372 << (uint(int32(2)) % 32)
	v410 = *(*int32)(unsafe.Add(mBase, uint32(v83+v408)))
	if v368 != v410 {
		goto L57
	} else {
		goto L58
	}
L54:
	;
	if v319 == int32(0) {
		v526 = v434
		goto L49
	} else {
		goto L70
	}
L55:
	;
	v421 = v372 - int32(1)
	v423 = v421 << (uint(int32(2)) % 32)
	v425 = *(*int32)(unsafe.Add(mBase, uint32(v83+v423)))
	if v368 != v425 {
		goto L64
	} else {
		goto L65
	}
L56:
	;
	v419 = v373
	goto L55
L57:
	;
	v412 = *(*int32)(unsafe.Add(mBase, uint32(v410)+616))
	if v412 != v368 {
		goto L56
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	if v373 == int32(-1) {
		v419 = v372
		goto L55
	} else {
		goto L61
	}
L60:
	;
	goto L59
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v408+v75))) = int32(-1)
	goto L56
L62:
	;
	v435 = int32(2)
	v436 = v372 - v435
	v438 = v380 + v435
	if v438 != v317 {
		v372 = v436
		v373 = v434
		v380 = v438
		goto L53
	} else {
		goto L69
	}
L63:
	;
	v434 = v419
	goto L62
L64:
	;
	v427 = *(*int32)(unsafe.Add(mBase, uint32(v425)+616))
	if v427 != v368 {
		goto L63
	} else {
		goto L67
	}
L65:
	;
	goto L66
L66:
	;
	if v419 == int32(-1) {
		v434 = v421
		goto L62
	} else {
		goto L68
	}
L67:
	;
	goto L66
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v75+v423))) = int32(-1)
	goto L63
L69:
	;
	goto L54
L70:
	;
	v443 = v436
	v444 = v434
	goto L52
L71:
	;
	v526 = v444
	goto L49
L72:
	;
	v483 = *(*int32)(unsafe.Add(mBase, uint32(v481)+616))
	if v483 != v368 {
		goto L71
	} else {
		goto L75
	}
L73:
	;
	goto L74
L74:
	;
	if v444 == int32(-1) {
		v526 = v443
		goto L49
	} else {
		goto L76
	}
L75:
	;
	goto L74
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v479+v75))) = int32(-1)
	goto L71
L77:
	;
	v529 = *(*int32)(unsafe.Add(mBase, uint32(v367)+4))
	v530 = int32(0)
	v531 = int32(-1)
	if base.B2i32(v181 == v313>>(uint(int32(31))%32)&v313+v318) == v530 {
		goto L79
	} else {
		goto L80
	}
L78:
	;
	if v689 < int32(0) {
		goto L47
	} else {
		goto L106
	}
L79:
	;
	v535 = v313
	v536 = v531
	v543 = v530
	goto L82
L80:
	;
	v606 = v313
	v607 = v531
	goto L81
L81:
	;
	v642 = v606 << (uint(int32(2)) % 32)
	v644 = *(*int32)(unsafe.Add(mBase, uint32(v83+v642)))
	if v529 != v644 {
		goto L101
	} else {
		goto L102
	}
L82:
	;
	v571 = v535 << (uint(int32(2)) % 32)
	v573 = *(*int32)(unsafe.Add(mBase, uint32(v83+v571)))
	if v529 != v573 {
		goto L86
	} else {
		goto L87
	}
L83:
	;
	if v319 == int32(0) {
		v689 = v597
		goto L78
	} else {
		goto L99
	}
L84:
	;
	v584 = v535 - int32(1)
	v586 = v584 << (uint(int32(2)) % 32)
	v588 = *(*int32)(unsafe.Add(mBase, uint32(v83+v586)))
	if v529 != v588 {
		goto L93
	} else {
		goto L94
	}
L85:
	;
	v582 = v536
	goto L84
L86:
	;
	v575 = *(*int32)(unsafe.Add(mBase, uint32(v573)+616))
	if v575 != v529 {
		goto L85
	} else {
		goto L89
	}
L87:
	;
	goto L88
L88:
	;
	if v536 == int32(-1) {
		v582 = v535
		goto L84
	} else {
		goto L90
	}
L89:
	;
	goto L88
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v571+v75))) = int32(-1)
	goto L85
L91:
	;
	v598 = int32(2)
	v599 = v535 - v598
	v601 = v543 + v598
	if v601 != v317 {
		v535 = v599
		v536 = v597
		v543 = v601
		goto L82
	} else {
		goto L98
	}
L92:
	;
	v597 = v582
	goto L91
L93:
	;
	v590 = *(*int32)(unsafe.Add(mBase, uint32(v588)+616))
	if v590 != v529 {
		goto L92
	} else {
		goto L96
	}
L94:
	;
	goto L95
L95:
	;
	if v582 == int32(-1) {
		v597 = v584
		goto L91
	} else {
		goto L97
	}
L96:
	;
	goto L95
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v75+v586))) = int32(-1)
	goto L92
L98:
	;
	goto L83
L99:
	;
	v606 = v599
	v607 = v597
	goto L81
L100:
	;
	v689 = v607
	goto L78
L101:
	;
	v646 = *(*int32)(unsafe.Add(mBase, uint32(v644)+616))
	if v646 != v529 {
		goto L100
	} else {
		goto L104
	}
L102:
	;
	goto L103
L103:
	;
	if v607 == int32(-1) {
		v689 = v606
		goto L78
	} else {
		goto L105
	}
L104:
	;
	goto L103
L105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v642+v75))) = int32(-1)
	goto L100
L106:
	;
	v692 = int32(2)
	v694 = v75 + v526<<(uint(v692)%32)
	v695 = *(*int32)(unsafe.Add(mBase, uint32(v694)))
	v696 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v694))) = v695 + v696
	*(*int32)(unsafe.Add(mBase, uint32(v367)+12)) = v526
	v702 = v67 + v689<<(uint(v692)%32)
	v703 = *(*int32)(unsafe.Add(mBase, uint32(v702)))
	*(*int32)(unsafe.Add(mBase, uint32(v367)+16)) = v703
	*(*int32)(unsafe.Add(mBase, uint32(v702))) = v345 + v696
	goto L47
L107:
	;
	goto L46
L108:
	;
	v788 = v313
	v798 = v313
	goto L111
L109:
	;
	goto L110
L110:
	;
	v1119 = v109 + int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_TestConfiguration[4])) = v1119
	v1142 = v1119
	v1149 = v181 + v116
	goto L9
L111:
	;
	v824 = v788 + int32(1)
	v825 = v788
	goto L113
L112:
	;
	goto L110
L113:
	;
	v859 = int32(1)
	v866 = *(*int32)(unsafe.Add(mBase, uint32(v83+v825<<(uint(int32(2))%32))))
	if v866 == int32(0) {
		v824 = v824 - v859
		v825 = v825 - v859
		goto L113
	} else {
		goto L115
	}
L114:
	;
	if v825 < int32(0) {
		goto L116
	} else {
		goto L117
	}
L115:
	;
	goto L114
L116:
	;
	v1619 = int32(-1)
	v1622 = v39
	goto L1
L117:
	;
	goto L118
L118:
	;
	v873 = v825
	goto L119
L119:
	;
	v909 = v873 << (uint(int32(2)) % 32)
	v911 = *(*int32)(unsafe.Add(mBase, uint32(v83+v909)))
	if v911 != 0 {
		goto L122
	} else {
		goto L123
	}
L120:
	;
	v921 = *(*int32)(unsafe.Add(mBase, uint32(v911)+616))
	if v921 != 0 {
		goto L127
	} else {
		goto L128
	}
L121:
	;
	goto L120
L122:
	;
	v913 = *(*int32)(unsafe.Add(mBase, uint32(v75+v909)))
	if v913 == int32(0) {
		goto L121
	} else {
		goto L125
	}
L123:
	;
	goto L124
L124:
	;
	if int32(0) < v873 {
		v873 = v873 - int32(1)
		goto L119
	} else {
		goto L126
	}
L125:
	;
	goto L124
L126:
	;
	v1619 = int32(-1)
	v1622 = v39
	goto L1
L127:
	;
	v922 = v921
	goto L129
L128:
	;
	v922 = v911
	goto L129
L129:
	;
	v923 = int32(0)
	v926 = v923
	v934 = v923
	goto L130
L130:
	;
	v963 = v83 + v926<<(uint(int32(2))%32)
	v964 = *(*int32)(unsafe.Add(mBase, uint32(v963)))
	if v922 != v964 {
		goto L133
	} else {
		goto L134
	}
L131:
	;
	v984 = *(*int32)(unsafe.Add(mBase, uint32(v67+v909)))
	if int32(0) < v984 {
		goto L139
	} else {
		goto L140
	}
L132:
	;
	v981 = v926 + int32(1)
	if v981 != v824 {
		v926 = v981
		v934 = v979
		goto L130
	} else {
		goto L138
	}
L133:
	;
	if v964 == int32(0) {
		v979 = v934
		goto L132
	} else {
		goto L136
	}
L134:
	;
	goto L135
L135:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v178+(v798-v934)<<(uint(int32(2))%32)))) = v964
	*(*int32)(unsafe.Add(mBase, uint32(v963))) = int32(0)
	v979 = v934 + int32(1)
	goto L132
L136:
	;
	v968 = *(*int32)(unsafe.Add(mBase, uint32(v964)+616))
	if v968 != v922 {
		v979 = v934
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
	v988 = v984
	goto L142
L140:
	;
	goto L141
L141:
	;
	v1077 = v798 - v979
	if int32(0) <= v1077 {
		v788 = v825
		v798 = v1077
		goto L111
	} else {
		goto L145
	}
L142:
	;
	v1025 = v65 + v988*int32(20)
	v1028 = *(*int32)(unsafe.Add(mBase, uint32(v1025-int32(8))))
	v1031 = v75 + v1028<<(uint(int32(2))%32)
	v1032 = *(*int32)(unsafe.Add(mBase, uint32(v1031)))
	*(*int32)(unsafe.Add(mBase, uint32(v1031))) = v1032 - int32(1)
	v1038 = *(*int32)(unsafe.Add(mBase, uint32(v1025-int32(4))))
	if int32(0) < v1038 {
		v988 = v1038
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
	v1166 = int32(0)
	v1169 = v1161
	goto L148
L148:
	;
	v1202 = v1166 * int32(20)
	v1204 = *(*int32)(unsafe.Add(mBase, _c_F_TestConfiguration[8]))
	v1206 = *(*int32)(unsafe.Add(mBase, uint32(v1202+v1204)))
	v1208 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_TestConfiguration[7])) = v1208
	*(*int32)(unsafe.Add(mBase, _c_F_TestConfiguration[6])) = v1208
	*(*int32)(unsafe.Add(mBase, uint32(v39)+12)) = v1208
	v1217 = v39 + int32(12)
	v1221 = *(*int32)(unsafe.Add(mBase, uint32(v1206)+616))
	if v1221 != 0 {
		goto L152
	} else {
		goto L153
	}
L149:
	;
	v1462 = v1452
	v1465 = v39
	goto L5
L150:
	;
	v1329 = *(*int32)(unsafe.Add(mBase, _c_F_TestConfiguration[8]))
	v1331 = *(*int32)(unsafe.Add(mBase, uint32(v1329+v1202)+4))
	v1333 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_TestConfiguration[7])) = v1333
	*(*int32)(unsafe.Add(mBase, _c_F_TestConfiguration[6])) = v1333
	*(*int32)(unsafe.Add(mBase, uint32(v39)+12)) = v1333
	v1342 = v39 + int32(12)
	v1346 = *(*int32)(unsafe.Add(mBase, uint32(v1331)+616))
	if v1346 != 0 {
		goto L186
	} else {
		goto L187
	}
L151:
	;
	if v1322 == int32(0) {
		v1327 = v1169
		goto L150
	} else {
		goto L182
	}
L152:
	;
	v1222 = v1221
	goto L154
L153:
	;
	v1222 = v1206
	goto L154
L154:
	;
	v1223 = int32(0)
	v1225 = *(*int32)(unsafe.Add(mBase, _c_F_TestConfiguration[14]))
	v1227 = *(*int32)(unsafe.Add(mBase, _c_F_TestConfiguration[7]))
	if v1223 < v1227 {
		goto L155
	} else {
		goto L156
	}
L155:
	;
	v1230 = v1223
	goto L158
L156:
	;
	goto L157
L157:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_TestConfiguration[7])) = v1227 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1225+v1227<<(uint(int32(2))%32)))) = v1222
	v1264 = *(*int32)(unsafe.Add(mBase, uint32(v1222)+4))
	if v1264 == int32(0) {
		goto L167
	} else {
		goto L168
	}
L158:
	;
	v1240 = *(*int32)(unsafe.Add(mBase, uint32(v1225+v1230<<(uint(int32(2))%32))))
	if v1222 == v1240 {
		goto L160
	} else {
		goto L161
	}
L159:
	;
	goto L157
L160:
	;
	if v1230 != 0 {
		goto L163
	} else {
		goto L164
	}
L161:
	;
	goto L162
L162:
	;
	v1247 = v1230 + int32(1)
	if v1247 != v1227 {
		v1230 = v1247
		goto L158
	} else {
		goto L166
	}
L163:
	;
	v1322 = int32(0)
	goto L151
L164:
	;
	goto L165
L165:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_TestConfiguration[6])) = v1208
	v1322 = int32(1)
	goto L151
L166:
	;
	goto L159
L167:
	;
	v1274 = *(*int32)(unsafe.Add(mBase, uint32(v1222)+624))
	if v1274 == int32(0) {
		goto L171
	} else {
		goto L172
	}
L168:
	;
	v1267 = *(*int32)(unsafe.Add(mBase, uint32(v1222)+92))
	if v1267 == int32(0) {
		goto L167
	} else {
		goto L169
	}
L169:
	;
	v1270 = F_FindLockCycleRecurseMember(m, v1222, v1222, v1208, v54, v1217)
	mBase = m.M
	if v1270 == int32(0) {
		goto L167
	} else {
		goto L170
	}
L170:
	;
	v1322 = int32(1)
	goto L151
L171:
	;
	v1322 = int32(0)
	goto L151
L172:
	;
	v1278 = v1222 + int32(620)
	if v1274 == v1278 {
		goto L171
	} else {
		goto L173
	}
L173:
	;
	v1280 = v1274
	goto L174
L174:
	;
	v1289 = *(*int32)(unsafe.Add(mBase, uint32(v1280-int32(624))))
	if v1289 == int32(0) {
		goto L176
	} else {
		goto L177
	}
L175:
	;
	goto L171
L176:
	;
	v1305 = *(*int32)(unsafe.Add(mBase, uint32(v1280)+4))
	if v1305 != v1278 {
		v1280 = v1305
		goto L174
	} else {
		goto L181
	}
L177:
	;
	v1294 = *(*int32)(unsafe.Add(mBase, uint32(v1280-int32(536))))
	if v1294 == int32(0) {
		goto L176
	} else {
		goto L178
	}
L178:
	;
	v1298 = v1280 - int32(628)
	if v1298 == v1222 {
		goto L176
	} else {
		goto L179
	}
L179:
	;
	v1300 = F_FindLockCycleRecurseMember(m, v1298, v1222, v1208, v54, v1217)
	mBase = m.M
	if v1300 == int32(0) {
		goto L176
	} else {
		goto L180
	}
L180:
	;
	v1322 = int32(1)
	goto L151
L181:
	;
	goto L175
L182:
	;
	v1325 = *(*int32)(unsafe.Add(mBase, uint32(v39)+12))
	if v1325 != 0 {
		v1327 = v1325
		goto L150
	} else {
		goto L183
	}
L183:
	;
	v1619 = int32(-1)
	v1622 = v39
	goto L1
L184:
	;
	v1454 = v1166 + int32(1)
	v1456 = *(*int32)(unsafe.Add(mBase, _c_F_TestConfiguration[5]))
	if v1454 < v1456 {
		v1166 = v1454
		v1169 = v1452
		goto L148
	} else {
		goto L218
	}
L185:
	;
	if v1447 == int32(0) {
		v1452 = v1327
		goto L184
	} else {
		goto L216
	}
L186:
	;
	v1347 = v1346
	goto L188
L187:
	;
	v1347 = v1331
	goto L188
L188:
	;
	v1348 = int32(0)
	v1350 = *(*int32)(unsafe.Add(mBase, _c_F_TestConfiguration[14]))
	v1352 = *(*int32)(unsafe.Add(mBase, _c_F_TestConfiguration[7]))
	if v1348 < v1352 {
		goto L189
	} else {
		goto L190
	}
L189:
	;
	v1355 = v1348
	goto L192
L190:
	;
	goto L191
L191:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_TestConfiguration[7])) = v1352 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1350+v1352<<(uint(int32(2))%32)))) = v1347
	v1389 = *(*int32)(unsafe.Add(mBase, uint32(v1347)+4))
	if v1389 == int32(0) {
		goto L201
	} else {
		goto L202
	}
L192:
	;
	v1365 = *(*int32)(unsafe.Add(mBase, uint32(v1350+v1355<<(uint(int32(2))%32))))
	if v1347 == v1365 {
		goto L194
	} else {
		goto L195
	}
L193:
	;
	goto L191
L194:
	;
	if v1355 != 0 {
		goto L197
	} else {
		goto L198
	}
L195:
	;
	goto L196
L196:
	;
	v1372 = v1355 + int32(1)
	if v1372 != v1352 {
		v1355 = v1372
		goto L192
	} else {
		goto L200
	}
L197:
	;
	v1447 = int32(0)
	goto L185
L198:
	;
	goto L199
L199:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_TestConfiguration[6])) = v1333
	v1447 = int32(1)
	goto L185
L200:
	;
	goto L193
L201:
	;
	v1399 = *(*int32)(unsafe.Add(mBase, uint32(v1347)+624))
	if v1399 == int32(0) {
		goto L205
	} else {
		goto L206
	}
L202:
	;
	v1392 = *(*int32)(unsafe.Add(mBase, uint32(v1347)+92))
	if v1392 == int32(0) {
		goto L201
	} else {
		goto L203
	}
L203:
	;
	v1395 = F_FindLockCycleRecurseMember(m, v1347, v1347, v1333, v54, v1342)
	mBase = m.M
	if v1395 == int32(0) {
		goto L201
	} else {
		goto L204
	}
L204:
	;
	v1447 = int32(1)
	goto L185
L205:
	;
	v1447 = int32(0)
	goto L185
L206:
	;
	v1403 = v1347 + int32(620)
	if v1399 == v1403 {
		goto L205
	} else {
		goto L207
	}
L207:
	;
	v1405 = v1399
	goto L208
L208:
	;
	v1414 = *(*int32)(unsafe.Add(mBase, uint32(v1405-int32(624))))
	if v1414 == int32(0) {
		goto L210
	} else {
		goto L211
	}
L209:
	;
	goto L205
L210:
	;
	v1430 = *(*int32)(unsafe.Add(mBase, uint32(v1405)+4))
	if v1430 != v1403 {
		v1405 = v1430
		goto L208
	} else {
		goto L215
	}
L211:
	;
	v1419 = *(*int32)(unsafe.Add(mBase, uint32(v1405-int32(536))))
	if v1419 == int32(0) {
		goto L210
	} else {
		goto L212
	}
L212:
	;
	v1423 = v1405 - int32(628)
	if v1423 == v1347 {
		goto L210
	} else {
		goto L213
	}
L213:
	;
	v1425 = F_FindLockCycleRecurseMember(m, v1423, v1347, v1333, v54, v1342)
	mBase = m.M
	if v1425 == int32(0) {
		goto L210
	} else {
		goto L214
	}
L214:
	;
	v1447 = int32(1)
	goto L185
L215:
	;
	goto L209
L216:
	;
	v1450 = *(*int32)(unsafe.Add(mBase, uint32(v39)+12))
	if v1450 != 0 {
		v1452 = v1450
		goto L184
	} else {
		goto L217
	}
L217:
	;
	v1619 = int32(-1)
	v1622 = v39
	goto L1
L218:
	;
	goto L149
L219:
	;
	if v1609 == int32(0) {
		v1619 = v1462
		v1622 = v1465
		goto L1
	} else {
		goto L250
	}
L220:
	;
	v1509 = v1508
	goto L222
L221:
	;
	v1509 = l0
	goto L222
L222:
	;
	v1510 = int32(0)
	v1512 = *(*int32)(unsafe.Add(mBase, _c_F_TestConfiguration[14]))
	v1514 = *(*int32)(unsafe.Add(mBase, _c_F_TestConfiguration[7]))
	if v1510 < v1514 {
		goto L223
	} else {
		goto L224
	}
L223:
	;
	v1517 = v1510
	goto L226
L224:
	;
	goto L225
L225:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_TestConfiguration[7])) = v1514 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1512+v1514<<(uint(int32(2))%32)))) = v1509
	v1551 = *(*int32)(unsafe.Add(mBase, uint32(v1509)+4))
	if v1551 == int32(0) {
		goto L235
	} else {
		goto L236
	}
L226:
	;
	v1527 = *(*int32)(unsafe.Add(mBase, uint32(v1512+v1517<<(uint(int32(2))%32))))
	if v1509 == v1527 {
		goto L228
	} else {
		goto L229
	}
L227:
	;
	goto L225
L228:
	;
	if v1517 != 0 {
		goto L231
	} else {
		goto L232
	}
L229:
	;
	goto L230
L230:
	;
	v1534 = v1517 + int32(1)
	if v1534 != v1514 {
		v1517 = v1534
		goto L226
	} else {
		goto L234
	}
L231:
	;
	v1609 = int32(0)
	goto L219
L232:
	;
	goto L233
L233:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_TestConfiguration[6])) = v1495
	v1609 = int32(1)
	goto L219
L234:
	;
	goto L227
L235:
	;
	v1561 = *(*int32)(unsafe.Add(mBase, uint32(v1509)+624))
	if v1561 == int32(0) {
		goto L239
	} else {
		goto L240
	}
L236:
	;
	v1554 = *(*int32)(unsafe.Add(mBase, uint32(v1509)+92))
	if v1554 == int32(0) {
		goto L235
	} else {
		goto L237
	}
L237:
	;
	v1557 = F_FindLockCycleRecurseMember(m, v1509, v1509, v1495, v54, v1504)
	mBase = m.M
	if v1557 == int32(0) {
		goto L235
	} else {
		goto L238
	}
L238:
	;
	v1609 = int32(1)
	goto L219
L239:
	;
	v1609 = int32(0)
	goto L219
L240:
	;
	v1565 = v1509 + int32(620)
	if v1561 == v1565 {
		goto L239
	} else {
		goto L241
	}
L241:
	;
	v1567 = v1561
	goto L242
L242:
	;
	v1576 = *(*int32)(unsafe.Add(mBase, uint32(v1567-int32(624))))
	if v1576 == int32(0) {
		goto L244
	} else {
		goto L245
	}
L243:
	;
	goto L239
L244:
	;
	v1592 = *(*int32)(unsafe.Add(mBase, uint32(v1567)+4))
	if v1592 != v1565 {
		v1567 = v1592
		goto L242
	} else {
		goto L249
	}
L245:
	;
	v1581 = *(*int32)(unsafe.Add(mBase, uint32(v1567-int32(536))))
	if v1581 == int32(0) {
		goto L244
	} else {
		goto L246
	}
L246:
	;
	v1585 = v1567 - int32(628)
	if v1585 == v1509 {
		goto L244
	} else {
		goto L247
	}
L247:
	;
	v1587 = F_FindLockCycleRecurseMember(m, v1585, v1509, v1495, v54, v1504)
	mBase = m.M
	if v1587 == int32(0) {
		goto L244
	} else {
		goto L248
	}
L248:
	;
	v1609 = int32(1)
	goto L219
L249:
	;
	goto L243
L250:
	;
	v1612 = *(*int32)(unsafe.Add(mBase, uint32(v1465)+12))
	if v1612 != 0 {
		goto L251
	} else {
		goto L252
	}
L251:
	;
	v1614 = v1612
	goto L253
L252:
	;
	v1614 = int32(-1)
	goto L253
L253:
	;
	v1619 = v1614
	v1622 = v1465
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
	F_errdetail(m, int32(_a_F_TransformPubWhereClauses_3), v11)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L14
	} else {
		goto L28
	}
L28:
	;
	F_errfinish(m, int32(_a_F_TransformPubWhereClauses_4), int32(734), int32(_a_F_TransformPubWhereClauses_5))
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
func F_textgename(m *base.Module, l0 int32) int32 {
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
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v6 = F_DirectFunctionCall2Coll(m, int32(1543), v3, v4, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return int32(base.Ui32(v6^int32(-1)) >> (uint(int32(31)) % 32))
	}
}
func F_textlike_support(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = F_like_regex_support(m, v2, int32(0))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_textregexeq(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn13995(m, l0, int32(19))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_textregexeq_support(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = F_like_regex_support(m, v2, int32(2))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_textregexreplace(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
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
	var v30 int32
	_ = v30
	var v47 int32
	_ = v47
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v75 int32
	_ = v75
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v14 = F_pg_detoast_datum_packed(m, v13)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v19 = F_pg_detoast_datum_packed(m, v18)
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int32(0)
		} else {
			v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
			v22 = F_pg_detoast_datum_packed(m, v21)
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int32(0)
			} else {
				v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
				v25 = F_pg_detoast_datum_packed(m, v24)
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int32(0)
				} else {
					v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25))))
					if v27 == int32(1) {
						v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+1)))
						if base.B2i32(base.Ui32(int32(18)) < base.Ui32(v30))|base.B2i32(int32(1)<<(uint(v30)%32)&int32(_a_F_textregexreplace_0) == int32(0)) != 0 {
							F_parse_re_flags(m, v11+int32(8), v25)
							mBase = m.M
							v120 = m.ExcPending
							if v120 != 0 {
								return int32(0)
							} else {
								v121 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
								v122 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
								v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+12)))
								v127 = F_replace_text_regexp(m, v14, v19, v22, v121, v122, int32(0), v124^int32(1))
								mBase = m.M
								v128 = m.ExcPending
								if v128 != 0 {
									return int32(0)
								} else {
									m.G0 = v11 + int32(16)
									return v127
								}
							}
						} else {
							if v30 == int32(18) {
								v47 = int32(16)
							} else {
								v47 = int32(0)
							}
							if base.Ui32((v30-int32(1))&int32(255)) < base.Ui32(int32(3)) {
								v54 = int32(4)
							} else {
								v54 = v47
							}
							v82 = v25 + int32(1)
							v83 = v54
							v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82))))
							if base.Ui32(int32(9)) < base.Ui32((v84-int32(48))&int32(255)) {
								F_parse_re_flags(m, v11+int32(8), v25)
								mBase = m.M
								v120 = m.ExcPending
								if v120 != 0 {
									return int32(0)
								} else {
									v121 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
									v122 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
									v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+12)))
									v127 = F_replace_text_regexp(m, v14, v19, v22, v121, v122, int32(0), v124^int32(1))
									mBase = m.M
									v128 = m.ExcPending
									if v128 != 0 {
										return int32(0)
									} else {
										m.G0 = v11 + int32(16)
										return v127
									}
								}
							} else {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v94 = m.ExcPending
								if v94 != 0 {
									return int32(0)
								} else {
									F_errcode(m, int32(50856066))
									mBase = m.M
									v97 = m.ExcPending
									if v97 != 0 {
										return int32(0)
									} else {
										v99 = F_pg_mblen_range(m, v82, v83+v82)
										mBase = m.M
										v100 = m.ExcPending
										if v100 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v82
											*(*int32)(unsafe.Add(mBase, uint32(v11))) = v99
											F_errmsg(m, int32(_a_F_textregexreplace_1), v11)
											mBase = m.M
											v105 = m.ExcPending
											if v105 != 0 {
												return int32(0)
											} else {
												F_errhint(m, int32(_a_F_textregexreplace_2), int32(0))
												mBase = m.M
												v109 = m.ExcPending
												if v109 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(_a_F_textregexreplace_3), int32(683), int32(_a_F_textregexreplace_4))
													mBase = m.M
													v114 = m.ExcPending
													if v114 != 0 {
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
					} else {
						if v27&int32(1) == int32(0) {
							v59 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
							if v59&int32(-4) == int32(16) {
								F_parse_re_flags(m, v11+int32(8), v25)
								mBase = m.M
								v120 = m.ExcPending
								if v120 != 0 {
									return int32(0)
								} else {
									v121 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
									v122 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
									v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+12)))
									v127 = F_replace_text_regexp(m, v14, v19, v22, v121, v122, int32(0), v124^int32(1))
									mBase = m.M
									v128 = m.ExcPending
									if v128 != 0 {
										return int32(0)
									} else {
										m.G0 = v11 + int32(16)
										return v127
									}
								}
							} else {
								v64 = int32(4)
								v66 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
								v82 = v25 + v64
								v83 = int32(base.Ui32(v66)>>(uint(int32(2))%32)) - v64
								v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82))))
								if base.Ui32(int32(9)) < base.Ui32((v84-int32(48))&int32(255)) {
									F_parse_re_flags(m, v11+int32(8), v25)
									mBase = m.M
									v120 = m.ExcPending
									if v120 != 0 {
										return int32(0)
									} else {
										v121 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
										v122 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
										v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+12)))
										v127 = F_replace_text_regexp(m, v14, v19, v22, v121, v122, int32(0), v124^int32(1))
										mBase = m.M
										v128 = m.ExcPending
										if v128 != 0 {
											return int32(0)
										} else {
											m.G0 = v11 + int32(16)
											return v127
										}
									}
								} else {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v94 = m.ExcPending
									if v94 != 0 {
										return int32(0)
									} else {
										F_errcode(m, int32(50856066))
										mBase = m.M
										v97 = m.ExcPending
										if v97 != 0 {
											return int32(0)
										} else {
											v99 = F_pg_mblen_range(m, v82, v83+v82)
											mBase = m.M
											v100 = m.ExcPending
											if v100 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v82
												*(*int32)(unsafe.Add(mBase, uint32(v11))) = v99
												F_errmsg(m, int32(_a_F_textregexreplace_1), v11)
												mBase = m.M
												v105 = m.ExcPending
												if v105 != 0 {
													return int32(0)
												} else {
													F_errhint(m, int32(_a_F_textregexreplace_2), int32(0))
													mBase = m.M
													v109 = m.ExcPending
													if v109 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(_a_F_textregexreplace_3), int32(683), int32(_a_F_textregexreplace_4))
														mBase = m.M
														v114 = m.ExcPending
														if v114 != 0 {
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
						} else {
							if v27&int32(254) == int32(2) {
								F_parse_re_flags(m, v11+int32(8), v25)
								mBase = m.M
								v120 = m.ExcPending
								if v120 != 0 {
									return int32(0)
								} else {
									v121 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
									v122 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
									v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+12)))
									v127 = F_replace_text_regexp(m, v14, v19, v22, v121, v122, int32(0), v124^int32(1))
									mBase = m.M
									v128 = m.ExcPending
									if v128 != 0 {
										return int32(0)
									} else {
										m.G0 = v11 + int32(16)
										return v127
									}
								}
							} else {
								v75 = int32(1)
								v82 = v25 + v75
								v83 = int32(base.Ui32(v27)>>(uint(v75)%32)) - v75
								v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82))))
								if base.Ui32(int32(9)) < base.Ui32((v84-int32(48))&int32(255)) {
									F_parse_re_flags(m, v11+int32(8), v25)
									mBase = m.M
									v120 = m.ExcPending
									if v120 != 0 {
										return int32(0)
									} else {
										v121 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
										v122 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
										v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+12)))
										v127 = F_replace_text_regexp(m, v14, v19, v22, v121, v122, int32(0), v124^int32(1))
										mBase = m.M
										v128 = m.ExcPending
										if v128 != 0 {
											return int32(0)
										} else {
											m.G0 = v11 + int32(16)
											return v127
										}
									}
								} else {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v94 = m.ExcPending
									if v94 != 0 {
										return int32(0)
									} else {
										F_errcode(m, int32(50856066))
										mBase = m.M
										v97 = m.ExcPending
										if v97 != 0 {
											return int32(0)
										} else {
											v99 = F_pg_mblen_range(m, v82, v83+v82)
											mBase = m.M
											v100 = m.ExcPending
											if v100 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v82
												*(*int32)(unsafe.Add(mBase, uint32(v11))) = v99
												F_errmsg(m, int32(_a_F_textregexreplace_1), v11)
												mBase = m.M
												v105 = m.ExcPending
												if v105 != 0 {
													return int32(0)
												} else {
													F_errhint(m, int32(_a_F_textregexreplace_2), int32(0))
													mBase = m.M
													v109 = m.ExcPending
													if v109 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(_a_F_textregexreplace_3), int32(683), int32(_a_F_textregexreplace_4))
														mBase = m.M
														v114 = m.ExcPending
														if v114 != 0 {
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
		}
	}
}
func F_textregexreplace_extended(m *base.Module, l0 int32) int32 {
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
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	v10 = m.G0
	v12 = v10 - int32(32)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v15 = F_pg_detoast_datum_packed(m, v14)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v20 = F_pg_detoast_datum_packed(m, v19)
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return int32(0)
		} else {
			v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
			v23 = F_pg_detoast_datum_packed(m, v22)
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int32(0)
			} else {
				v25 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
				if int32(6) <= v25 {
					v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
					v29 = F_pg_detoast_datum_packed(m, v28)
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
						return int32(0)
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
								return int32(0)
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
									return int32(0)
								} else {
									m.G0 = v12 + int32(32)
									return v65
								}
							}
						} else {
							v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
							if v38 <= int32(0) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v74 = m.ExcPending
								if v74 != 0 {
									return int32(0)
								} else {
									F_errcode(m, int32(50856066))
									mBase = m.M
									v77 = m.ExcPending
									if v77 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v38
										*(*int32)(unsafe.Add(mBase, uint32(v12))) = int32(_a_F_textregexreplace_extended_0)
										F_errmsg(m, int32(_a_F_textregexreplace_extended_1), v12)
										mBase = m.M
										v83 = m.ExcPending
										if v83 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_textregexreplace_extended_2), int32(718), int32(_a_F_textregexreplace_extended_3))
											mBase = m.M
											v88 = m.ExcPending
											if v88 != 0 {
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
								if v32&int32(_a_F_textregexreplace_extended_4) == int32(4) {
									v48 = v38
									v49 = v34
									F_parse_re_flags(m, v12+int32(24), v33)
									mBase = m.M
									v53 = m.ExcPending
									if v53 != 0 {
										return int32(0)
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
											return int32(0)
										} else {
											m.G0 = v12 + int32(32)
											return v65
										}
									}
								} else {
									v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
									if v45 < int32(0) {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v92 = m.ExcPending
										if v92 != 0 {
											return int32(0)
										} else {
											F_errcode(m, int32(50856066))
											mBase = m.M
											v95 = m.ExcPending
											if v95 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v45
												*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = int32(_a_F_textregexreplace_extended_5)
												F_errmsg(m, int32(_a_F_textregexreplace_extended_1), v12+int32(16))
												mBase = m.M
												v103 = m.ExcPending
												if v103 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(_a_F_textregexreplace_extended_2), int32(727), int32(_a_F_textregexreplace_extended_3))
													mBase = m.M
													v108 = m.ExcPending
													if v108 != 0 {
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
										v48 = v38
										v49 = v45
										F_parse_re_flags(m, v12+int32(24), v33)
										mBase = m.M
										v53 = m.ExcPending
										if v53 != 0 {
											return int32(0)
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
												return int32(0)
											} else {
												m.G0 = v12 + int32(32)
												return v65
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
							return int32(0)
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
								return int32(0)
							} else {
								m.G0 = v12 + int32(32)
								return v65
							}
						}
					} else {
						v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
						if v38 <= int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v74 = m.ExcPending
							if v74 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(50856066))
								mBase = m.M
								v77 = m.ExcPending
								if v77 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v38
									*(*int32)(unsafe.Add(mBase, uint32(v12))) = int32(_a_F_textregexreplace_extended_0)
									F_errmsg(m, int32(_a_F_textregexreplace_extended_1), v12)
									mBase = m.M
									v83 = m.ExcPending
									if v83 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_textregexreplace_extended_2), int32(718), int32(_a_F_textregexreplace_extended_3))
										mBase = m.M
										v88 = m.ExcPending
										if v88 != 0 {
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
							if v32&int32(_a_F_textregexreplace_extended_4) == int32(4) {
								v48 = v38
								v49 = v34
								F_parse_re_flags(m, v12+int32(24), v33)
								mBase = m.M
								v53 = m.ExcPending
								if v53 != 0 {
									return int32(0)
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
										return int32(0)
									} else {
										m.G0 = v12 + int32(32)
										return v65
									}
								}
							} else {
								v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
								if v45 < int32(0) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v92 = m.ExcPending
									if v92 != 0 {
										return int32(0)
									} else {
										F_errcode(m, int32(50856066))
										mBase = m.M
										v95 = m.ExcPending
										if v95 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v45
											*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = int32(_a_F_textregexreplace_extended_5)
											F_errmsg(m, int32(_a_F_textregexreplace_extended_1), v12+int32(16))
											mBase = m.M
											v103 = m.ExcPending
											if v103 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(_a_F_textregexreplace_extended_2), int32(727), int32(_a_F_textregexreplace_extended_3))
												mBase = m.M
												v108 = m.ExcPending
												if v108 != 0 {
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
									v48 = v38
									v49 = v45
									F_parse_re_flags(m, v12+int32(24), v33)
									mBase = m.M
									v53 = m.ExcPending
									if v53 != 0 {
										return int32(0)
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
											return int32(0)
										} else {
											m.G0 = v12 + int32(32)
											return v65
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
func F_thesaurus_lexize(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v14 int32
	_ = v14
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
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v199 int32
	_ = v199
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v226 int32
	_ = v226
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v247 int32
	_ = v247
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v283 int32
	_ = v283
	var v289 int32
	_ = v289
	var v298 int32
	_ = v298
	var v307 int32
	_ = v307
	var v315 int32
	_ = v315
	var v323 int32
	_ = v323
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v354 int32
	_ = v354
	var v363 int32
	_ = v363
	var v365 int32
	_ = v365
	var v381 int32
	_ = v381
	var v397 int32
	_ = v397
	var v408 int32
	_ = v408
	var v410 int32
	_ = v410
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
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
	var v444 int32
	_ = v444
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v472 int32
	_ = v472
	var v475 int32
	_ = v475
	var v483 int32
	_ = v483
	var v490 int32
	_ = v490
	var v495 int32
	_ = v495
	var v499 int32
	_ = v499
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v511 int32
	_ = v511
	var v513 int32
	_ = v513
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v545 int32
	_ = v545
	var v559 int32
	_ = v559
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v586 int32
	_ = v586
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v598 int32
	_ = v598
	var v607 int32
	_ = v607
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v619 int32
	_ = v619
	var v621 int32
	_ = v621
	var v623 int32
	_ = v623
	var v643 int32
	_ = v643
	var v649 int32
	_ = v649
	var v658 int32
	_ = v658
	var v675 int32
	_ = v675
	var v683 int32
	_ = v683
	var v692 int32
	_ = v692
	var v694 int32
	_ = v694
	var v714 int32
	_ = v714
	var v723 int32
	_ = v723
	var v725 int32
	_ = v725
	var v741 int32
	_ = v741
	var v757 int32
	_ = v757
	var v818 int32
	_ = v818
	var v819 int32
	_ = v819
	var v828 int32
	_ = v828
	var v829 int32
	_ = v829
	var v830 int32
	_ = v830
	var v832 int32
	_ = v832
	var v835 int32
	_ = v835
	var v843 int32
	_ = v843
	var v864 int32
	_ = v864
	var v867 int32
	_ = v867
	var v869 int32
	_ = v869
	var v871 int32
	_ = v871
	var v872 int32
	_ = v872
	var v873 int32
	_ = v873
	var v876 int32
	_ = v876
	var v877 int32
	_ = v877
	var v881 int32
	_ = v881
	var v882 int32
	_ = v882
	var v893 int32
	_ = v893
	var v894 int32
	_ = v894
	var v896 int32
	_ = v896
	var v899 int32
	_ = v899
	var v900 int32
	_ = v900
	var v903 int32
	_ = v903
	var v904 int32
	_ = v904
	var v916 int32
	_ = v916
	var v918 int32
	_ = v918
	var v919 int32
	_ = v919
	var v924 int32
	_ = v924
	var v925 int32
	_ = v925
	var v926 int32
	_ = v926
	var v930 int32
	_ = v930
	var v942 int32
	_ = v942
	var v943 int32
	_ = v943
	var v944 int32
	_ = v944
	var v946 int64
	_ = v946
	var v948 int32
	_ = v948
	var v950 int32
	_ = v950
	var v951 int32
	_ = v951
	var v952 int32
	_ = v952
	var v955 int32
	_ = v955
	var v956 int32
	_ = v956
	var v974 int32
	_ = v974
	var v978 int32
	_ = v978
	var v982 int32
	_ = v982
	v2 = int32(0)
	v14 = m.G0
	v16 = v14 - int32(16)
	m.G0 = v16
	v18 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+18)))
	if v18 != int32(4) {
		goto L6
	} else {
		goto L7
	}
L1:
	;
	m.G0 = v16 + int32(16)
	return v982
L2:
	;
	v867 = *(*int32)(unsafe.Add(mBase, uint32(v843)+12))
	v869 = base.B2i32(v867 == int32(0))
	v871 = v31 & int32(_a_F_thesaurus_lexize_0)
	v872 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
	v873 = *(*int32)(unsafe.Add(mBase, uint32(v843)))
	v876 = v872 + v873<<(uint(int32(3))%32)
	v877 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v876))))
	if v871 == v877 {
		goto L172
	} else {
		goto L173
	}
L3:
	;
	v864 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+1)) = uint8(v864)
	v982 = v864
	goto L1
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+4)) = v843
	if v843 != 0 {
		goto L2
	} else {
		goto L170
	}
L5:
	;
	v505 = int32(0)
	v506 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
	if v506 == v505 {
		v521 = v505
		goto L109
	} else {
		goto L110
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v495 = m.ExcPending
	if v495 != 0 {
		goto L16
	} else {
		goto L106
	}
L7:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v21 == int32(0) {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
	if v24 != 0 {
		v982 = v2
		goto L1
	} else {
		goto L9
	}
L9:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	if v26 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v27 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v26)+4)))
	v31 = v27 + int32(1)
	goto L12
L11:
	;
	v31 = int32(0)
	goto L12
L12:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+4)))
	if v33 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	v37 = F_lookup_ts_dictionary_cache(m, v36)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	v42 = v32
	goto L15
L15:
	;
	v45 = int32(0)
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v42)+44))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v50 = F_FunctionCall4Coll(m, v42+int32(12), v45, v46, v47, v48, v45)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L16
	} else {
		goto L18
	}
L16:
	;
	return int32(0)
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+4)) = v37
	v42 = v37
	goto L15
L18:
	;
	if v50 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+4)) = int32(0)
	goto L3
L20:
	;
	goto L21
L21:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
	if v56 == int32(0) {
		goto L5
	} else {
		goto L22
	}
L22:
	;
	v60 = v50
	v65 = v2
	goto L23
L23:
	;
	v72 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v60))))
	v74 = v60
	v76 = int32(0)
	goto L25
L24:
	;
	v843 = v483
	goto L4
L25:
	;
	v87 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v74))))
	if v87 != v72 {
		goto L28
	} else {
		goto L29
	}
L26:
	;
	v99 = v96 & int32(_a_F_thesaurus_lexize_0)
	v102 = F_palloc(m, v99<<(uint(int32(2))%32))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L16
	} else {
		goto L32
	}
L27:
	;
	goto L26
L28:
	;
	v95 = v74
	v96 = v76
	goto L27
L29:
	;
	goto L30
L30:
	;
	v90 = v76 + int32(1)
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v74)+12))
	v93 = v74 + int32(8)
	if v91 != 0 {
		v74 = v93
		v76 = v90
		goto L25
	} else {
		goto L31
	}
L31:
	;
	v95 = v93
	v96 = v90
	goto L27
L32:
	;
	if v99 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	v490 = *(*int32)(unsafe.Add(mBase, uint32(v95)+4))
	if v490 != 0 {
		v60 = v95
		v65 = v483
		goto L23
	} else {
		goto L105
	}
L34:
	;
	v182 = v99 & int32(3)
	v185 = v65
	goto L49
L35:
	;
	v106 = int32(0)
	goto L36
L36:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
	if v119 != 0 {
		goto L40
	} else {
		goto L41
	}
L37:
	;
	F_pfree(m, v102)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L16
	} else {
		goto L47
	}
L38:
	;
	goto L37
L39:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v132)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v102+v106<<(uint(int32(2))%32)))) = v144
	if v144 == int32(0) {
		goto L38
	} else {
		goto L45
	}
L40:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v60+v106<<(uint(int32(3))%32))+4))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+12)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = v123
	v127 = int32(8)
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
	v132 = F_bsearch(m, v16+v127, v129, v119, v127, int32(1150))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L16
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v102+v106<<(uint(int32(2))%32)))) = int32(0)
	goto L38
L43:
	;
	if v132 != 0 {
		goto L39
	} else {
		goto L44
	}
L44:
	;
	goto L42
L45:
	;
	v149 = v106 + int32(1)
	if v99 != v149 {
		v106 = v149
		goto L36
	} else {
		goto L46
	}
L46:
	;
	goto L34
L47:
	;
	v483 = v65
	goto L33
L48:
	;
	v483 = v185
	goto L33
L49:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v102)))
	if v99 != 0 {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	if v99 == int32(0) {
		v185 = v381
		goto L49
	} else {
		goto L94
	}
L52:
	;
	v207 = v199
	v208 = int32(0)
	goto L55
L53:
	;
	v307 = v199
	goto L54
L54:
	;
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v307)))
	if v26 == int32(0) {
		goto L81
	} else {
		goto L82
	}
L55:
	;
	v217 = v102 + v208<<(uint(int32(2))%32)
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v217)))
	if v218 == int32(0) {
		goto L59
	} else {
		goto L60
	}
L56:
	;
	if v99 != v298 {
		v381 = v185
		goto L51
	} else {
		goto L80
	}
L57:
	;
	if v298 < v99 {
		v207 = v289
		v208 = v298
		goto L55
	} else {
		goto L79
	}
L58:
	;
	if v256 == v257 {
		goto L76
	} else {
		goto L77
	}
L59:
	;
	goto L48
L60:
	;
	v226 = v218
	goto L61
L61:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v226)))
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v207)))
	if base.Ui32(v235) < base.Ui32(v236) {
		goto L63
	} else {
		goto L64
	}
L62:
	;
	if base.Ui32(v236) < base.Ui32(v235) {
		v289 = v226
		v298 = int32(0)
		goto L57
	} else {
		goto L67
	}
L63:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v226)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v217))) = v238
	if v238 != 0 {
		v226 = v238
		goto L61
	} else {
		goto L66
	}
L64:
	;
	goto L65
L65:
	;
	goto L62
L66:
	;
	goto L59
L67:
	;
	v247 = v226
	goto L68
L68:
	;
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v247)))
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v207)))
	if v256 != v257 {
		goto L58
	} else {
		goto L70
	}
L69:
	;
	goto L59
L70:
	;
	v259 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v247)+4)))
	if v31&int32(_a_F_thesaurus_lexize_0) == v259 {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v261 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v247)+6)))
	if v99 == v261 {
		goto L58
	} else {
		goto L74
	}
L72:
	;
	goto L73
L73:
	;
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v247)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v217))) = v263
	if v263 != 0 {
		v247 = v263
		goto L68
	} else {
		goto L75
	}
L74:
	;
	goto L73
L75:
	;
	goto L69
L76:
	;
	v283 = v208 + int32(1)
	goto L78
L77:
	;
	v283 = int32(0)
	goto L78
L78:
	;
	v289 = v247
	v298 = v283
	goto L57
L79:
	;
	goto L56
L80:
	;
	v307 = v289
	goto L54
L81:
	;
	if v185 != 0 {
		goto L87
	} else {
		goto L88
	}
L82:
	;
	v323 = v26
	goto L83
L83:
	;
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v323)))
	if v332 == v315 {
		goto L81
	} else {
		goto L85
	}
L84:
	;
	v381 = v185
	goto L51
L85:
	;
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v323)+12))
	if v334 != 0 {
		v323 = v334
		goto L83
	} else {
		goto L86
	}
L86:
	;
	goto L84
L87:
	;
	v354 = v185
	goto L90
L88:
	;
	goto L89
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v307)+12)) = v185
	v381 = v307
	goto L51
L90:
	;
	v363 = *(*int32)(unsafe.Add(mBase, uint32(v354)))
	if v363 == v315 {
		v381 = v185
		goto L51
	} else {
		goto L92
	}
L91:
	;
	goto L89
L92:
	;
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v354)+12))
	if v365 != 0 {
		v354 = v365
		goto L90
	} else {
		goto L93
	}
L93:
	;
	goto L91
L94:
	;
	v397 = int32(0)
	if base.B2i32(base.Ui32(v99) < base.Ui32(int32(4))) == v397 {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	v408 = v397
	v410 = v397
	goto L98
L96:
	;
	v444 = v397
	goto L97
L97:
	;
	v458 = v444
	v459 = v397
	goto L102
L98:
	;
	v418 = v102 + v408<<(uint(int32(2))%32)
	v419 = *(*int32)(unsafe.Add(mBase, uint32(v418)))
	v420 = *(*int32)(unsafe.Add(mBase, uint32(v419)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v418))) = v420
	v422 = *(*int32)(unsafe.Add(mBase, uint32(v418)+4))
	v423 = *(*int32)(unsafe.Add(mBase, uint32(v422)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v418)+4)) = v423
	v425 = *(*int32)(unsafe.Add(mBase, uint32(v418)+8))
	v426 = *(*int32)(unsafe.Add(mBase, uint32(v425)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v418)+8)) = v426
	v428 = *(*int32)(unsafe.Add(mBase, uint32(v418)+12))
	v429 = *(*int32)(unsafe.Add(mBase, uint32(v428)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v418)+12)) = v429
	v431 = int32(4)
	v432 = v408 + v431
	v434 = v410 + v431
	if v434 != v99&int32(_a_F_thesaurus_lexize_1) {
		v408 = v432
		v410 = v434
		goto L98
	} else {
		goto L100
	}
L99:
	;
	if v182 == int32(0) {
		v185 = v381
		goto L49
	} else {
		goto L101
	}
L100:
	;
	goto L99
L101:
	;
	v444 = v432
	goto L97
L102:
	;
	v468 = v102 + v458<<(uint(int32(2))%32)
	v469 = *(*int32)(unsafe.Add(mBase, uint32(v468)))
	v470 = *(*int32)(unsafe.Add(mBase, uint32(v469)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v468))) = v470
	v472 = int32(1)
	v475 = v459 + v472
	if v475 != v182 {
		v458 = v458 + v472
		v459 = v475
		goto L102
	} else {
		goto L104
	}
L103:
	;
	v185 = v381
	goto L49
L104:
	;
	goto L103
L105:
	;
	goto L24
L106:
	;
	F_errmsg_internal(m, int32(_a_F_thesaurus_lexize_2), int32(0))
	mBase = m.M
	v499 = m.ExcPending
	if v499 != 0 {
		goto L16
	} else {
		goto L107
	}
L107:
	;
	F_errfinish(m, int32(_a_F_thesaurus_lexize_3), int32(799), int32(_a_F_thesaurus_lexize_4))
	mBase = m.M
	v504 = m.ExcPending
	if v504 != 0 {
		goto L16
	} else {
		goto L108
	}
L108:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = v521
	v528 = v16 + int32(4)
	v529 = int32(1)
	v545 = int32(0)
	goto L114
L110:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v16)+8)) = int64(0)
	v511 = int32(8)
	v513 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
	v516 = F_bsearch(m, v16+v511, v513, v506, v511, int32(1150))
	mBase = m.M
	v517 = m.ExcPending
	if v517 != 0 {
		goto L16
	} else {
		goto L111
	}
L111:
	;
	if v516 == int32(0) {
		v521 = v505
		goto L109
	} else {
		goto L112
	}
L112:
	;
	v520 = *(*int32)(unsafe.Add(mBase, uint32(v516)+4))
	v521 = v520
	goto L109
L113:
	;
	v843 = v545
	goto L4
L114:
	;
	v559 = *(*int32)(unsafe.Add(mBase, uint32(v528)))
	goto L117
L116:
	;
	goto L159
L117:
	;
	v567 = v559
	v568 = int32(0)
	goto L120
L119:
	;
	v675 = *(*int32)(unsafe.Add(mBase, uint32(v649)))
	if v26 == int32(0) {
		goto L146
	} else {
		goto L147
	}
L120:
	;
	v577 = v528 + v568<<(uint(int32(2))%32)
	v578 = *(*int32)(unsafe.Add(mBase, uint32(v577)))
	if v578 == int32(0) {
		goto L124
	} else {
		goto L125
	}
L121:
	;
	if v529 != v658 {
		v741 = v545
		goto L116
	} else {
		goto L145
	}
L122:
	;
	if v658 < v529 {
		v567 = v649
		v568 = v658
		goto L120
	} else {
		goto L144
	}
L123:
	;
	if v616 == v617 {
		goto L141
	} else {
		goto L142
	}
L124:
	;
	goto L113
L125:
	;
	v586 = v578
	goto L126
L126:
	;
	v595 = *(*int32)(unsafe.Add(mBase, uint32(v586)))
	v596 = *(*int32)(unsafe.Add(mBase, uint32(v567)))
	if base.Ui32(v595) < base.Ui32(v596) {
		goto L128
	} else {
		goto L129
	}
L127:
	;
	if base.Ui32(v596) < base.Ui32(v595) {
		v649 = v586
		v658 = int32(0)
		goto L122
	} else {
		goto L132
	}
L128:
	;
	v598 = *(*int32)(unsafe.Add(mBase, uint32(v586)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v577))) = v598
	if v598 != 0 {
		v586 = v598
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
	goto L124
L132:
	;
	v607 = v586
	goto L133
L133:
	;
	v616 = *(*int32)(unsafe.Add(mBase, uint32(v607)))
	v617 = *(*int32)(unsafe.Add(mBase, uint32(v567)))
	if v616 != v617 {
		goto L123
	} else {
		goto L135
	}
L134:
	;
	goto L124
L135:
	;
	v619 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v607)+4)))
	if v31&int32(_a_F_thesaurus_lexize_0) == v619 {
		goto L136
	} else {
		goto L137
	}
L136:
	;
	v621 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v607)+6)))
	if v529 == v621 {
		goto L123
	} else {
		goto L139
	}
L137:
	;
	goto L138
L138:
	;
	v623 = *(*int32)(unsafe.Add(mBase, uint32(v607)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v577))) = v623
	if v623 != 0 {
		v607 = v623
		goto L133
	} else {
		goto L140
	}
L139:
	;
	goto L138
L140:
	;
	goto L134
L141:
	;
	v643 = v568 + int32(1)
	goto L143
L142:
	;
	v643 = int32(0)
	goto L143
L143:
	;
	v649 = v607
	v658 = v643
	goto L122
L144:
	;
	goto L121
L145:
	;
	goto L119
L146:
	;
	if v545 != 0 {
		goto L152
	} else {
		goto L153
	}
L147:
	;
	v683 = v26
	goto L148
L148:
	;
	v692 = *(*int32)(unsafe.Add(mBase, uint32(v683)))
	if v692 == v675 {
		goto L146
	} else {
		goto L150
	}
L149:
	;
	v741 = v545
	goto L116
L150:
	;
	v694 = *(*int32)(unsafe.Add(mBase, uint32(v683)+12))
	if v694 != 0 {
		v683 = v694
		goto L148
	} else {
		goto L151
	}
L151:
	;
	goto L149
L152:
	;
	v714 = v545
	goto L155
L153:
	;
	goto L154
L154:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v649)+12)) = v545
	v741 = v649
	goto L116
L155:
	;
	v723 = *(*int32)(unsafe.Add(mBase, uint32(v714)))
	if v723 == v675 {
		v741 = v545
		goto L116
	} else {
		goto L157
	}
L156:
	;
	goto L154
L157:
	;
	v725 = *(*int32)(unsafe.Add(mBase, uint32(v714)+12))
	if v725 != 0 {
		v714 = v725
		goto L155
	} else {
		goto L158
	}
L158:
	;
	goto L156
L159:
	;
	v757 = int32(0)
	goto L161
L161:
	;
	goto L162
L162:
	;
	v818 = v757
	v819 = v757
	goto L167
L167:
	;
	v828 = v528 + v818<<(uint(int32(2))%32)
	v829 = *(*int32)(unsafe.Add(mBase, uint32(v828)))
	v830 = *(*int32)(unsafe.Add(mBase, uint32(v829)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v828))) = v830
	v832 = int32(1)
	v835 = v819 + v832
	if v835 != v529 {
		v818 = v818 + v832
		v819 = v835
		goto L167
	} else {
		goto L169
	}
L168:
	;
	v545 = v741
	goto L114
L169:
	;
	goto L168
L170:
	;
	goto L3
L171:
	;
	v978 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+1)) = uint8(v978)
	v982 = int32(0)
	goto L1
L172:
	;
	v903 = v876
	v904 = v869
	v916 = int32(0)
	goto L174
L173:
	;
	v881 = v869
	v882 = v867
	goto L175
L174:
	;
	if v904 != 0 {
		goto L179
	} else {
		goto L180
	}
L175:
	;
	if v881 != 0 {
		goto L171
	} else {
		goto L177
	}
L176:
	;
	v903 = v899
	v904 = v896
	v916 = int32(1)
	goto L174
L177:
	;
	v893 = *(*int32)(unsafe.Add(mBase, uint32(v882)))
	v894 = *(*int32)(unsafe.Add(mBase, uint32(v882)+12))
	v896 = base.B2i32(v894 == int32(0))
	v899 = v872 + v893<<(uint(int32(3))%32)
	v900 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v899))))
	if v900 != v871 {
		v881 = v896
		v882 = v894
		goto L175
	} else {
		goto L178
	}
L178:
	;
	goto L176
L179:
	;
	v918 = v916
	goto L181
L180:
	;
	v918 = int32(1)
	goto L181
L181:
	;
	v919 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v903)+2)))
	v924 = F_palloc(m, v919<<(uint(int32(3))%32)+int32(8))
	mBase = m.M
	v925 = m.ExcPending
	if v925 != 0 {
		goto L16
	} else {
		goto L182
	}
L182:
	;
	v926 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v903)+2)))
	if v926 != 0 {
		goto L183
	} else {
		goto L184
	}
L183:
	;
	v930 = int32(0)
	goto L186
L184:
	;
	v974 = v924
	goto L185
L185:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v974)+4)) = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+1)) = uint8(v918)
	v982 = v924
	goto L1
L186:
	;
	v942 = v930 << (uint(int32(3)) % 32)
	v943 = v924 + v942
	v944 = *(*int32)(unsafe.Add(mBase, uint32(v903)+4))
	v946 = *(*int64)(unsafe.Add(mBase, uint32(v944+v942)))
	*(*int64)(unsafe.Add(mBase, uint32(v943))) = v946
	v948 = *(*int32)(unsafe.Add(mBase, uint32(v903)+4))
	v950 = *(*int32)(unsafe.Add(mBase, uint32(v948+v942)+4))
	v951 = F_pstrdup(m, v950)
	mBase = m.M
	v952 = m.ExcPending
	if v952 != 0 {
		goto L16
	} else {
		goto L188
	}
L187:
	;
	v974 = v924 + v956<<(uint(int32(3))%32)
	goto L185
L188:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v943)+4)) = v951
	v955 = v930 + int32(1)
	v956 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v903)+2)))
	if base.Ui32(v955) < base.Ui32(v956) {
		v930 = v955
		goto L186
	} else {
		goto L189
	}
L189:
	;
	goto L187
}
func F_tidlarger(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4)+2)))
	v9 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4))))
	v10 = int32(16)
	v12 = v8 | v9<<(uint(v10)%32)
	v13 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3)+2)))
	v14 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3))))
	v17 = v13 | v14<<(uint(v10)%32)
	if base.Ui32(v12) < base.Ui32(v17) {
		v28 = int32(-1)
	} else {
		if base.Ui32(v17) < base.Ui32(v12) {
			v28 = int32(1)
		} else {
			v22 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4)+4)))
			v23 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3)+4)))
			if base.Ui32(v22) < base.Ui32(v23) {
				v28 = int32(-1)
			} else {
				v28 = base.B2i32(base.Ui32(v23) < base.Ui32(v22))
			}
		}
	}
	if v28 < int32(0) {
		v31 = v3
	} else {
		v31 = v4
	}
	return v31
}
func F_timestamptypmodin(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14000(m, l0, int32(0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_timestamptztypmodin(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14000(m, l0, int32(1))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_timetypmodout(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14002(m, l0, int32(_a_F_timetypmodout_0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
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
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
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
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v66 int32
	_ = v66
	v6 = m.G0
	v8 = v6 + int32(-64)
	m.G0 = v8
	v15 = F_toast_open_indexes(m, l0, int32(3), v6+int32(-56), v6+int32(-52))
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
	v20 = v6 + int32(-48)
	F_ScanKeyInit(m, v20, int32(1), int32(3), int32(184), l1)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v26+v15<<(uint(int32(2))%32))))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+56))
	v32 = int32(1)
	v35 = F_systable_beginscan(m, l0, v31, v32, int32(_a_F_toastrel_valueid_exists_0), v32, v20)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v37 = F_systable_getnext(m, v35)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	F_systable_endscan(m, v35)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	if int32(0) < v41 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v45 = int32(0)
	goto L10
L8:
	;
	goto L9
L9:
	;
	F_pfree(m, v26)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L1
	} else {
		goto L14
	}
L10:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v26+v45<<(uint(int32(2))%32))))
	F_relation_close(m, v53, int32(3))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L1
	} else {
		goto L12
	}
L11:
	;
	goto L9
L12:
	;
	v58 = v45 + int32(1)
	if v58 != v41 {
		v45 = v58
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
	return base.B2i32(v37 != int32(0))
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
											F_errfinish(m, int32(_a_F_transformLimitClause_1), int32(1907), int32(_a_F_transformLimitClause_2))
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
	F_errfinish(m, int32(_a_F_transformLockingClause_2), int32(3539), int32(_a_F_transformLockingClause_3))
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
	F_errfinish(m, int32(_a_F_transformLockingClause_2), int32(3603), int32(_a_F_transformLockingClause_3))
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
	F_errfinish(m, int32(_a_F_transformLockingClause_2), int32(3612), int32(_a_F_transformLockingClause_3))
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
	F_errfinish(m, int32(_a_F_transformLockingClause_2), int32(3621), int32(_a_F_transformLockingClause_3))
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
	F_errfinish(m, int32(_a_F_transformLockingClause_2), int32(3630), int32(_a_F_transformLockingClause_3))
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
	F_errfinish(m, int32(_a_F_transformLockingClause_2), int32(3639), int32(_a_F_transformLockingClause_3))
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
	F_errfinish(m, int32(_a_F_transformLockingClause_2), int32(3648), int32(_a_F_transformLockingClause_3))
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
	F_errfinish(m, int32(_a_F_transformLockingClause_2), int32(3655), int32(_a_F_transformLockingClause_3))
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
	F_errfinish(m, int32(_a_F_transformLockingClause_2), int32(3669), int32(_a_F_transformLockingClause_3))
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
				*(*int32)(unsafe.Add(mBase, uint32(v31)+12)) = int32(41)
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
func F_translate(m *base.Module, l0 int32) int32 {
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
	var v432 int32
	_ = v432
	var v435 int32
	_ = v435
	var v439 int32
	_ = v439
	var v444 int32
	_ = v444
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
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
	return int32(0)
L2:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
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
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
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
	v432 = m.ExcPending
	if v432 != 0 {
		goto L1
	} else {
		goto L102
	}
L6:
	;
	return v411
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
	v435 = m.ExcPending
	if v435 != 0 {
		goto L1
	} else {
		goto L103
	}
L103:
	;
	F_errmsg(m, int32(_a_F_translate_0), int32(0))
	mBase = m.M
	v439 = m.ExcPending
	if v439 != 0 {
		goto L1
	} else {
		goto L104
	}
L104:
	;
	F_errfinish(m, int32(_a_F_translate_1), int32(864), int32(_a_F_translate_2))
	mBase = m.M
	v444 = m.ExcPending
	if v444 != 0 {
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
func F_tsm_handler_out(m *base.Module, l0 int32) int32 {
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v7 = Fn13852(m, l0, int32(_a_F_tsm_handler_out_0), int32(372), int32(_a_F_tsm_handler_out_1), int32(_a_F_tsm_handler_out_2), int32(_a_F_tsm_handler_out_3))
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
	}
}
func F_tsqueryin(m *base.Module, l0 int32) int32 {
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
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = int32(0)
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v7 = F_parse_tsquery(m, v2, int32(1510), v4, v4, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
	}
}
func F_tsvectorout(m *base.Module, l0 int32) int32 {
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
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v173 int32
	_ = v173
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v273 int32
	_ = v273
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v298 int32
	_ = v298
	var v310 int32
	_ = v310
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v319 int32
	_ = v319
	var v322 int32
	_ = v322
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v341 int32
	_ = v341
	var v348 int32
	_ = v348
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v364 int32
	_ = v364
	var v368 int32
	_ = v368
	var v372 int32
	_ = v372
	var v376 int32
	_ = v376
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v383 int32
	_ = v383
	var v385 int32
	_ = v385
	var v399 int32
	_ = v399
	var v402 int32
	_ = v402
	var v415 int32
	_ = v415
	var v417 int32
	_ = v417
	var v420 int32
	_ = v420
	v2 = int32(0)
	v15 = m.G0
	v17 = v15 - int32(16)
	m.G0 = v17
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
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
	return int32(0)
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
	v34 = v30
	v35 = v2
	goto L6
L4:
	;
	v93 = v30
	goto L5
L5:
	;
	v106 = F_palloc(m, v93)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L1
	} else {
		goto L13
	}
L6:
	;
	v49 = v25 + v35<<(uint(int32(2))%32)
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
	v54 = *(*int32)(unsafe.Add(mBase, _c_F_tsvectorout[0]))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v54)+4))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v55*int32(28))+uint32(_c_F_tsvectorout[1])))
	goto L8
L7:
	;
	v93 = v88
	goto L5
L8:
	;
	v62 = v50&int32(4094)*v60 + v34
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
	if v64&int32(1) != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v70 = int32(1)
	v82 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v25+v63<<(uint(int32(2))%32)+(int32(base.Ui32(v64)>>(uint(v70)%32))&int32(2047)+int32(base.Ui32(v64)>>(uint(int32(12))%32))+v70)&int32(_a_F_tsvectorout_0)))))
	v88 = v62 + v82*int32(7) + v70
	goto L11
L10:
	;
	v88 = v62
	goto L11
L11:
	;
	v90 = v35 + int32(1)
	if v90 < v63 {
		v34 = v88
		v35 = v90
		goto L6
	} else {
		goto L12
	}
L12:
	;
	goto L7
L13:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	if int32(0) < v108 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v112 = v106
	v114 = v108
	v115 = v25
	v121 = v2
	goto L17
L15:
	;
	v402 = v106
	goto L16
L16:
	;
	v415 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v402))) = uint8(v415)
	v417 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v417 != v20 {
		goto L59
	} else {
		goto L60
	}
L17:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v115)))
	v129 = int32(base.Ui32(v125)>>(uint(int32(1))%32)) & int32(2047)
	if v121 != 0 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v402 = v383
	goto L16
L19:
	;
	v130 = int32(32)
	*(*uint8)(unsafe.Add(mBase, uint32(v112))) = uint8(v130)
	v134 = v112 + int32(1)
	goto L21
L20:
	;
	v134 = v112
	goto L21
L21:
	;
	v135 = int32(39)
	*(*uint8)(unsafe.Add(mBase, uint32(v134))) = uint8(v135)
	v138 = v134 + int32(1)
	if v129 != 0 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v144 = v25 + v114<<(uint(int32(2))%32) + int32(base.Ui32(v125)>>(uint(int32(12))%32))
	v145 = v144 + v129
	v147 = v144
	v148 = v138
	goto L25
L23:
	;
	v273 = v138
	goto L24
L24:
	;
	v285 = int32(39)
	*(*uint8)(unsafe.Add(mBase, uint32(v273))) = uint8(v285)
	v287 = int32(1)
	v288 = v273 + v287
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v115)))
	if v290&v287 == int32(0) {
		v383 = v288
		v385 = v289
		goto L44
	} else {
		goto L45
	}
L25:
	;
	v160 = F_pg_mblen_range(m, v147, v145)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L1
	} else {
		goto L27
	}
L26:
	;
	v273 = v258
	goto L24
L27:
	;
	v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v147))))
	if base.B2i32(v162 != int32(92))&base.B2i32(v162 != int32(39)) == int32(0) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v148))) = uint8(v162)
	v173 = v148 + int32(1)
	goto L30
L29:
	;
	v173 = v148
	goto L30
L30:
	;
	if v160 == int32(0) {
		v257 = v147
		v258 = v173
		goto L31
	} else {
		goto L32
	}
L31:
	;
	if base.Ui32(v257) < base.Ui32(v145) {
		v147 = v257
		v148 = v258
		goto L25
	} else {
		goto L43
	}
L32:
	;
	v178 = v160 & int32(7)
	if v178 != 0 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v180 = v147
	v181 = v173
	v182 = v160
	v186 = int32(0)
	goto L36
L34:
	;
	v205 = v147
	v206 = v173
	v207 = v160
	goto L35
L35:
	;
	if base.Ui32(v160) < base.Ui32(int32(8)) {
		v257 = v205
		v258 = v206
		goto L31
	} else {
		goto L39
	}
L36:
	;
	v193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180))))
	*(*uint8)(unsafe.Add(mBase, uint32(v181))) = uint8(v193)
	v195 = int32(1)
	v196 = v181 + v195
	v198 = v180 + v195
	v200 = v182 - v195
	v202 = v186 + v195
	if v202 != v178 {
		v180 = v198
		v181 = v196
		v182 = v200
		v186 = v202
		goto L36
	} else {
		goto L38
	}
L37:
	;
	v205 = v198
	v206 = v196
	v207 = v200
	goto L35
L38:
	;
	goto L37
L39:
	;
	v221 = v205
	v222 = v206
	v223 = v207
	goto L40
L40:
	;
	v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v221))))
	*(*uint8)(unsafe.Add(mBase, uint32(v222))) = uint8(v234)
	v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v221)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v222)+1)) = uint8(v236)
	v238 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v221)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v222)+2)) = uint8(v238)
	v240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v221)+3)))
	*(*uint8)(unsafe.Add(mBase, uint32(v222)+3)) = uint8(v240)
	v242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v221)+4)))
	*(*uint8)(unsafe.Add(mBase, uint32(v222)+4)) = uint8(v242)
	v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v221)+5)))
	*(*uint8)(unsafe.Add(mBase, uint32(v222)+5)) = uint8(v244)
	v246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v221)+6)))
	*(*uint8)(unsafe.Add(mBase, uint32(v222)+6)) = uint8(v246)
	v248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v221)+7)))
	*(*uint8)(unsafe.Add(mBase, uint32(v222)+7)) = uint8(v248)
	v250 = int32(8)
	v251 = v222 + v250
	v253 = v221 + v250
	v255 = v223 - v250
	if v255 != 0 {
		v221 = v253
		v222 = v251
		v223 = v255
		goto L40
	} else {
		goto L42
	}
L41:
	;
	v257 = v253
	v258 = v251
	goto L31
L42:
	;
	goto L41
L43:
	;
	goto L26
L44:
	;
	v399 = v121 + int32(1)
	if v399 < v385 {
		v112 = v383
		v114 = v385
		v115 = v115 + int32(4)
		v121 = v399
		goto L17
	} else {
		goto L58
	}
L45:
	;
	v298 = int32(1)
	v310 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v25+v289<<(uint(int32(2))%32)+(int32(base.Ui32(v290)>>(uint(v298)%32))&int32(2047)+int32(base.Ui32(v290)>>(uint(int32(12))%32))+v298)&int32(_a_F_tsvectorout_0)))))
	if v310 == int32(0) {
		v383 = v288
		v385 = v289
		goto L44
	} else {
		goto L46
	}
L46:
	;
	v313 = int32(58)
	*(*uint8)(unsafe.Add(mBase, uint32(v273)+1)) = uint8(v313)
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	v316 = int32(2)
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v115)))
	v322 = int32(1)
	v335 = v273 + v316
	v337 = v25 + v315<<(uint(v316)%32) + (int32(base.Ui32(v319)>>(uint(int32(12))%32))+int32(base.Ui32(v319)>>(uint(v322)%32))&int32(2047)+v322)&int32(_a_F_tsvectorout_0)
	v341 = v310
	goto L47
L47:
	;
	v348 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v337)+2)))
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v348 & int32(_a_F_tsvectorout_1)
	v353 = F_pg_sprintf(m, v335, int32(_a_F_tsvectorout_2), v17)
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L1
	} else {
		goto L49
	}
L48:
	;
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	v383 = v376
	v385 = v381
	goto L44
L49:
	;
	v355 = v353 + v335
	v357 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v337)+2)))
	switch int32(base.Ui32(v357)>>(uint(int32(14))%32)) - int32(1) {
	case 0:
		goto L52
	case 1:
		goto L53
	case 2:
		v364 = int32(65)
		goto L51
	default:
		v368 = v355
		goto L50
	}
L50:
	;
	if int32(2) <= v341 {
		goto L54
	} else {
		goto L55
	}
L51:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v355))) = uint8(v364)
	v368 = v355 + int32(1)
	goto L50
L52:
	;
	v364 = int32(67)
	goto L51
L53:
	;
	v364 = int32(66)
	goto L51
L54:
	;
	v372 = int32(44)
	*(*uint8)(unsafe.Add(mBase, uint32(v368))) = uint8(v372)
	v376 = v368 + int32(1)
	goto L56
L55:
	;
	v376 = v368
	goto L56
L56:
	;
	v380 = v341 - int32(1)
	if v380 != 0 {
		v335 = v376
		v337 = v337 + int32(2)
		v341 = v380
		goto L47
	} else {
		goto L57
	}
L57:
	;
	goto L48
L58:
	;
	goto L18
L59:
	;
	F_pfree(m, v20)
	mBase = m.M
	v420 = m.ExcPending
	if v420 != 0 {
		goto L1
	} else {
		goto L62
	}
L60:
	;
	goto L61
L61:
	;
	m.G0 = v17 + int32(16)
	return v106
L62:
	;
	goto L61
}
func F_tsvectorsend(m *base.Module, l0 int32) int32 {
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
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
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
	return int32(0)
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
	return v199
}
func F_tuples_equal(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
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
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
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
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v107 int32
	_ = v107
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
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
	F_slot_getsomeattrs_int(m, l0, v16)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	v25 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+6)))
	if v25 < v24 {
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
	F_slot_getsomeattrs_int(m, l1, v24)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L4
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	if v31 <= int32(0) {
		v107 = int32(1)
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
	v118 = m.ExcPending
	if v118 != 0 {
		goto L4
	} else {
		goto L28
	}
L11:
	;
	m.G0 = v13 + int32(16)
	return v107
L12:
	;
	v37 = v30
	v38 = v31
	v39 = int32(0)
	goto L13
L13:
	;
	v49 = v37 + v38<<(uint(int32(4))%32) + v39*int32(100)
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49)+111)))
	if v50 != 0 {
		v92 = v37
		v93 = v38
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v107 = v97
	goto L11
L15:
	;
	v97 = int32(1)
	v99 = v39 + v97
	if v99 < v93 {
		v37 = v92
		v38 = v93
		v39 = v99
		goto L13
	} else {
		goto L27
	}
L16:
	;
	v52 = v49 + int32(20)
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+90)))
	if v53 != 0 {
		v92 = v37
		v93 = v38
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v54 = int32(0)
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55+v39))))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58+v39))))
	if v57 != v60 {
		v107 = v54
		goto L11
	} else {
		goto L18
	}
L18:
	;
	if v57 != 0 {
		v92 = v37
		v93 = v38
		goto L15
	} else {
		goto L19
	}
L19:
	;
	v63 = v39 << (uint(int32(2)) % 32)
	v64 = l2 + v63
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
	if v65 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v52)+68))
	v70 = F_lookup_type_cache(m, v68, int32(32))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L4
	} else {
		goto L23
	}
L21:
	;
	v76 = v65
	goto L22
L22:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v52)+96))
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v80+v63)))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v83+v63)))
	v86 = F_FunctionCall2Coll(m, v76+int32(76), v79, v82, v85)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L4
	} else {
		goto L25
	}
L23:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v70)+80))
	if v72 == int32(0) {
		goto L10
	} else {
		goto L24
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v64))) = v70
	v76 = v70
	goto L22
L25:
	;
	if v86 == int32(0) {
		v107 = v54
		goto L11
	} else {
		goto L26
	}
L26:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v90)))
	v92 = v90
	v93 = v91
	goto L15
L27:
	;
	goto L14
L28:
	;
	F_errcode(m, int32(52461700))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L4
	} else {
		goto L29
	}
L29:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v52)+68))
	v123 = F_format_type_be(m, v122)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L4
	} else {
		goto L30
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v123
	F_errmsg(m, int32(_a_F_tuples_equal_0), v13)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L4
	} else {
		goto L31
	}
L31:
	;
	F_errfinish(m, int32(_a_F_tuples_equal_1), int32(330), int32(_a_F_tuples_equal_2))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L4
	} else {
		goto L32
	}
L32:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
