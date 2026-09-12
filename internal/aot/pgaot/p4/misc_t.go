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
	var v102 int32
	_ = v102
	var __phi102 int32
	_ = __phi102
	var v103 int32
	_ = v103
	var __phi103 int32
	_ = __phi103
	var v108 int32
	_ = v108
	var __phi108 int32
	_ = __phi108
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
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v232 int32
	_ = v232
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v291 int32
	_ = v291
	var v294 int32
	_ = v294
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v337 int32
	_ = v337
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v363 int32
	_ = v363
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v401 int32
	_ = v401
	var v403 int32
	_ = v403
	var v405 int32
	_ = v405
	var v412 int32
	_ = v412
	var v414 int32
	_ = v414
	var v416 int32
	_ = v416
	var v418 int32
	_ = v418
	var v420 int32
	_ = v420
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v431 int32
	_ = v431
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v472 int32
	_ = v472
	var v474 int32
	_ = v474
	var v476 int32
	_ = v476
	var v485 int32
	_ = v485
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v530 int32
	_ = v530
	var v532 int32
	_ = v532
	var v534 int32
	_ = v534
	var v541 int32
	_ = v541
	var v543 int32
	_ = v543
	var v545 int32
	_ = v545
	var v547 int32
	_ = v547
	var v549 int32
	_ = v549
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v560 int32
	_ = v560
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v601 int32
	_ = v601
	var v603 int32
	_ = v603
	var v605 int32
	_ = v605
	var v614 int32
	_ = v614
	var v617 int32
	_ = v617
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v670 int32
	_ = v670
	var v714 int32
	_ = v714
	var v720 int32
	_ = v720
	var v749 int32
	_ = v749
	var v750 int32
	_ = v750
	var v784 int32
	_ = v784
	var v791 int32
	_ = v791
	var v798 int32
	_ = v798
	var v834 int32
	_ = v834
	var v836 int32
	_ = v836
	var v838 int32
	_ = v838
	var v846 int32
	_ = v846
	var v847 int32
	_ = v847
	var v848 int32
	_ = v848
	var v851 int32
	_ = v851
	var v858 int32
	_ = v858
	var v888 int32
	_ = v888
	var v889 int32
	_ = v889
	var v893 int32
	_ = v893
	var v904 int32
	_ = v904
	var v906 int32
	_ = v906
	var v909 int32
	_ = v909
	var v913 int32
	_ = v913
	var v950 int32
	_ = v950
	var v953 int32
	_ = v953
	var v956 int32
	_ = v956
	var v957 int32
	_ = v957
	var v963 int32
	_ = v963
	var v1002 int32
	_ = v1002
	var v1044 int32
	_ = v1044
	var v1061 int32
	_ = v1061
	var v1066 int32
	_ = v1066
	var v1086 int32
	_ = v1086
	var v1092 int32
	_ = v1092
	var v1094 int32
	_ = v1094
	var v1128 int32
	_ = v1128
	var v1130 int32
	_ = v1130
	var v1132 int32
	_ = v1132
	var v1134 int32
	_ = v1134
	var v1143 int32
	_ = v1143
	var v1147 int32
	_ = v1147
	var v1148 int32
	_ = v1148
	var v1149 int32
	_ = v1149
	var v1151 int32
	_ = v1151
	var v1153 int32
	_ = v1153
	var v1156 int32
	_ = v1156
	var v1166 int32
	_ = v1166
	var v1173 int32
	_ = v1173
	var v1190 int32
	_ = v1190
	var v1193 int32
	_ = v1193
	var v1196 int32
	_ = v1196
	var v1200 int32
	_ = v1200
	var v1204 int32
	_ = v1204
	var v1206 int32
	_ = v1206
	var v1215 int32
	_ = v1215
	var v1220 int32
	_ = v1220
	var v1224 int32
	_ = v1224
	var v1226 int32
	_ = v1226
	var v1231 int32
	_ = v1231
	var v1248 int32
	_ = v1248
	var v1251 int32
	_ = v1251
	var v1253 int32
	_ = v1253
	var v1255 int32
	_ = v1255
	var v1257 int32
	_ = v1257
	var v1259 int32
	_ = v1259
	var v1268 int32
	_ = v1268
	var v1272 int32
	_ = v1272
	var v1273 int32
	_ = v1273
	var v1274 int32
	_ = v1274
	var v1276 int32
	_ = v1276
	var v1278 int32
	_ = v1278
	var v1281 int32
	_ = v1281
	var v1291 int32
	_ = v1291
	var v1298 int32
	_ = v1298
	var v1315 int32
	_ = v1315
	var v1318 int32
	_ = v1318
	var v1321 int32
	_ = v1321
	var v1325 int32
	_ = v1325
	var v1329 int32
	_ = v1329
	var v1331 int32
	_ = v1331
	var v1340 int32
	_ = v1340
	var v1345 int32
	_ = v1345
	var v1349 int32
	_ = v1349
	var v1351 int32
	_ = v1351
	var v1356 int32
	_ = v1356
	var v1373 int32
	_ = v1373
	var v1376 int32
	_ = v1376
	var v1378 int32
	_ = v1378
	var v1380 int32
	_ = v1380
	var v1382 int32
	_ = v1382
	var v1387 int32
	_ = v1387
	var v1421 int32
	_ = v1421
	var v1430 int32
	_ = v1430
	var v1434 int32
	_ = v1434
	var v1435 int32
	_ = v1435
	var v1436 int32
	_ = v1436
	var v1438 int32
	_ = v1438
	var v1440 int32
	_ = v1440
	var v1443 int32
	_ = v1443
	var v1453 int32
	_ = v1453
	var v1460 int32
	_ = v1460
	var v1477 int32
	_ = v1477
	var v1480 int32
	_ = v1480
	var v1483 int32
	_ = v1483
	var v1487 int32
	_ = v1487
	var v1491 int32
	_ = v1491
	var v1493 int32
	_ = v1493
	var v1502 int32
	_ = v1502
	var v1507 int32
	_ = v1507
	var v1511 int32
	_ = v1511
	var v1513 int32
	_ = v1513
	var v1518 int32
	_ = v1518
	var v1535 int32
	_ = v1535
	var v1538 int32
	_ = v1538
	var v1540 int32
	_ = v1540
	var v1544 int32
	_ = v1544
	v2 = int32(0)
	v37 = m.G0
	v39 = v37 - int32(16)
	m.G0 = v39
	v42 = *(*int32)(unsafe.Add(mBase, _consts[1162]))
	v44 = *(*int32)(unsafe.Add(mBase, _consts[1161]))
	v46 = *(*int32)(unsafe.Add(mBase, _consts[771]))
	if v42 < v44+v46 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v39 + int32(16)
	return v1544
L2:
	;
	v1544 = int32(-1)
	goto L1
L3:
	;
	goto L4
L4:
	;
	v51 = *(*int32)(unsafe.Add(mBase, _consts[1164]))
	v54 = v51 + v44*int32(20)
	v56 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[1165])) = v56
	v59 = *(*int32)(unsafe.Add(mBase, _consts[1159]))
	v61 = v59 - int32(1)
	if v61 < v56 {
		v1387 = v2
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v1421 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[1166])) = v1421
	*(*int32)(unsafe.Add(mBase, _consts[1167])) = v1421
	*(*int32)(unsafe.Add(mBase, uint32(v39)+12)) = v1421
	v1430 = v39 + int32(12)
	v1434 = *(*int32)(unsafe.Add(mBase, uint32(l0)+616))
	if v1434 != 0 {
		goto L220
	} else {
		goto L221
	}
L6:
	;
	v65 = *(*int32)(unsafe.Add(mBase, _consts[1163]))
	v67 = *(*int32)(unsafe.Add(mBase, _consts[1168]))
	v68 = int32(-1)
	v70 = int32(4)
	v71 = v67 + v70
	v72 = int32(3)
	v75 = *(*int32)(unsafe.Add(mBase, _consts[1169]))
	v79 = v75 + v70
	v83 = *(*int32)(unsafe.Add(mBase, _consts[1170]))
	v85 = *(*int32)(unsafe.Add(mBase, _consts[1171]))
	v87 = *(*int32)(unsafe.Add(mBase, _consts[1172]))
	__phi89 = v61
	__phi102 = v59
	__phi103 = v2
	__phi108 = v2
	v89 = __phi89
	v102 = __phi102
	v103 = __phi103
	v108 = __phi108
	goto L7
L7:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v65+v89*int32(20))+8))
	v129 = v103
	goto L10
L8:
	;
	v1086 = int32(0)
	if v59 <= v1086 {
		goto L145
	} else {
		goto L146
	}
L9:
	;
	if int32(0) < v89 {
		__phi89 = v89 - int32(1)
		__phi102 = v89
		__phi103 = v1061
		__phi108 = v1066
		v89 = __phi89
		v102 = __phi102
		v103 = __phi103
		v108 = __phi108
		goto L7
	} else {
		goto L144
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
	v175 = v87 + v103*int32(12)
	*(*int32)(unsafe.Add(mBase, uint32(v175))) = v127
	v179 = v85 + v108<<(uint(int32(2))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v175)+4)) = v179
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
	v1061 = v103
	v1066 = v108
	goto L9
L16:
	;
	v272 = v181 << (uint(int32(2)) % 32)
	if v75&v72 != 0 {
		v285 = v272
		goto L23
	} else {
		goto L24
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
	v194 = int32(0)
	goto L19
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v83+v194<<(uint(int32(2))%32)))) = v191
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v191)+4))
	if v232 != v188 {
		v191 = v232
		v194 = v194 + int32(1)
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
	if v67&v72 != 0 {
		v300 = v272
		goto L32
	} else {
		goto L33
	}
L23:
	;
	v287 = F__emscripten_memset_bulkmem(m, v75, base.I32_extend8_s(int32(0)), v285)
	mBase = m.M
	goto L30
L24:
	;
	if base.Ui32(int32(1024)) < base.Ui32(v272) {
		v285 = v272
		goto L23
	} else {
		goto L25
	}
L25:
	;
	v275 = v272 + v75
	if base.Ui32(v275) <= base.Ui32(v75) {
		goto L22
	} else {
		goto L26
	}
L26:
	;
	if base.Ui32(v79) < base.Ui32(v275) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v278 = v275
	goto L29
L28:
	;
	v278 = v79
	goto L29
L29:
	;
	v285 = (v278+(v75^v68))&int32(-4) + int32(4)
	goto L23
L30:
	;
	goto L22
L31:
	;
	v306 = v181 - int32(1)
	if int32(0) < v102 {
		goto L40
	} else {
		goto L41
	}
L32:
	;
	v303 = F__emscripten_memset_bulkmem(m, v67, base.I32_extend8_s(int32(0)), v300)
	mBase = m.M
	goto L39
L33:
	;
	if base.Ui32(int32(1024)) < base.Ui32(v272) {
		v300 = v272
		goto L32
	} else {
		goto L34
	}
L34:
	;
	v291 = v272 + v67
	if base.Ui32(v291) <= base.Ui32(v67) {
		goto L31
	} else {
		goto L35
	}
L35:
	;
	if base.Ui32(v71) < base.Ui32(v291) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v294 = v291
	goto L38
L37:
	;
	v294 = v71
	goto L38
L38:
	;
	v300 = (v294+(v67^v68))&int32(-4) + int32(4)
	goto L32
L39:
	;
	goto L31
L40:
	;
	v310 = v181 & int32(-2)
	v311 = int32(1)
	v312 = v181 & v311
	v337 = int32(0)
	goto L43
L41:
	;
	goto L42
L42:
	;
	if int32(0) <= v306 {
		goto L106
	} else {
		goto L107
	}
L43:
	;
	if v306 < int32(0) {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	goto L42
L45:
	;
	v670 = v337 + int32(1)
	if v670 != v102 {
		v337 = v670
		goto L43
	} else {
		goto L105
	}
L46:
	;
	v360 = v65 + v337*int32(20)
	v361 = *(*int32)(unsafe.Add(mBase, uint32(v360)))
	v363 = int32(-1)
	if v306 != 0 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v365 = v306
	v366 = v363
	v367 = int32(0)
	goto L50
L48:
	;
	v434 = v306
	v435 = v363
	goto L49
L49:
	;
	if v312 == int32(0) {
		v485 = v435
		goto L67
	} else {
		goto L68
	}
L50:
	;
	v401 = v365 << (uint(int32(2)) % 32)
	v403 = *(*int32)(unsafe.Add(mBase, uint32(v83+v401)))
	if v361 != v403 {
		goto L54
	} else {
		goto L55
	}
L51:
	;
	v434 = v429
	v435 = v427
	goto L49
L52:
	;
	v414 = v365 - int32(1)
	v416 = v414 << (uint(int32(2)) % 32)
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v83+v416)))
	if v361 != v418 {
		goto L61
	} else {
		goto L62
	}
L53:
	;
	v412 = v366
	goto L52
L54:
	;
	v405 = *(*int32)(unsafe.Add(mBase, uint32(v403)+616))
	if v405 != v361 {
		goto L53
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	if v366 == int32(-1) {
		v412 = v365
		goto L52
	} else {
		goto L58
	}
L57:
	;
	goto L56
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v75+v401))) = int32(-1)
	goto L53
L59:
	;
	v428 = int32(2)
	v429 = v365 - v428
	v431 = v367 + v428
	if v431 != v310 {
		v365 = v429
		v366 = v427
		v367 = v431
		goto L50
	} else {
		goto L66
	}
L60:
	;
	v427 = v412
	goto L59
L61:
	;
	v420 = *(*int32)(unsafe.Add(mBase, uint32(v418)+616))
	if v420 != v361 {
		goto L60
	} else {
		goto L64
	}
L62:
	;
	goto L63
L63:
	;
	if v412 == int32(-1) {
		v427 = v414
		goto L59
	} else {
		goto L65
	}
L64:
	;
	goto L63
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v75+v416))) = int32(-1)
	goto L60
L66:
	;
	goto L51
L67:
	;
	if v485 < int32(0) {
		goto L45
	} else {
		goto L75
	}
L68:
	;
	v472 = v434 << (uint(int32(2)) % 32)
	v474 = *(*int32)(unsafe.Add(mBase, uint32(v83+v472)))
	if v361 != v474 {
		goto L70
	} else {
		goto L71
	}
L69:
	;
	v485 = v435
	goto L67
L70:
	;
	v476 = *(*int32)(unsafe.Add(mBase, uint32(v474)+616))
	if v476 != v361 {
		goto L69
	} else {
		goto L73
	}
L71:
	;
	goto L72
L72:
	;
	if v435 == int32(-1) {
		v485 = v434
		goto L67
	} else {
		goto L74
	}
L73:
	;
	goto L72
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v472+v75))) = int32(-1)
	goto L69
L75:
	;
	v488 = *(*int32)(unsafe.Add(mBase, uint32(v360)+4))
	v489 = int32(0)
	v490 = int32(-1)
	if base.B2i32(v181 == v306>>(uint(int32(31))%32)&v306+v311) == v489 {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v494 = v306
	v495 = v490
	v496 = v489
	goto L79
L77:
	;
	v563 = v306
	v564 = v490
	goto L78
L78:
	;
	if v312 == int32(0) {
		v614 = v564
		goto L96
	} else {
		goto L97
	}
L79:
	;
	v530 = v494 << (uint(int32(2)) % 32)
	v532 = *(*int32)(unsafe.Add(mBase, uint32(v83+v530)))
	if v488 != v532 {
		goto L83
	} else {
		goto L84
	}
L80:
	;
	v563 = v558
	v564 = v556
	goto L78
L81:
	;
	v543 = v494 - int32(1)
	v545 = v543 << (uint(int32(2)) % 32)
	v547 = *(*int32)(unsafe.Add(mBase, uint32(v83+v545)))
	if v488 != v547 {
		goto L90
	} else {
		goto L91
	}
L82:
	;
	v541 = v495
	goto L81
L83:
	;
	v534 = *(*int32)(unsafe.Add(mBase, uint32(v532)+616))
	if v534 != v488 {
		goto L82
	} else {
		goto L86
	}
L84:
	;
	goto L85
L85:
	;
	if v495 == int32(-1) {
		v541 = v494
		goto L81
	} else {
		goto L87
	}
L86:
	;
	goto L85
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v75+v530))) = int32(-1)
	goto L82
L88:
	;
	v557 = int32(2)
	v558 = v494 - v557
	v560 = v496 + v557
	if v560 != v310 {
		v494 = v558
		v495 = v556
		v496 = v560
		goto L79
	} else {
		goto L95
	}
L89:
	;
	v556 = v541
	goto L88
L90:
	;
	v549 = *(*int32)(unsafe.Add(mBase, uint32(v547)+616))
	if v549 != v488 {
		goto L89
	} else {
		goto L93
	}
L91:
	;
	goto L92
L92:
	;
	if v541 == int32(-1) {
		v556 = v543
		goto L88
	} else {
		goto L94
	}
L93:
	;
	goto L92
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v75+v545))) = int32(-1)
	goto L89
L95:
	;
	goto L80
L96:
	;
	if v614 < int32(0) {
		goto L45
	} else {
		goto L104
	}
L97:
	;
	v601 = v563 << (uint(int32(2)) % 32)
	v603 = *(*int32)(unsafe.Add(mBase, uint32(v83+v601)))
	if v488 != v603 {
		goto L99
	} else {
		goto L100
	}
L98:
	;
	v614 = v564
	goto L96
L99:
	;
	v605 = *(*int32)(unsafe.Add(mBase, uint32(v603)+616))
	if v605 != v488 {
		goto L98
	} else {
		goto L102
	}
L100:
	;
	goto L101
L101:
	;
	if v564 == int32(-1) {
		v614 = v563
		goto L96
	} else {
		goto L103
	}
L102:
	;
	goto L101
L103:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v601+v75))) = int32(-1)
	goto L98
L104:
	;
	v617 = int32(2)
	v619 = v75 + v485<<(uint(v617)%32)
	v620 = *(*int32)(unsafe.Add(mBase, uint32(v619)))
	v621 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v619))) = v620 + v621
	*(*int32)(unsafe.Add(mBase, uint32(v360)+12)) = v485
	v627 = v67 + v614<<(uint(v617)%32)
	v628 = *(*int32)(unsafe.Add(mBase, uint32(v627)))
	*(*int32)(unsafe.Add(mBase, uint32(v360)+16)) = v628
	*(*int32)(unsafe.Add(mBase, uint32(v627))) = v337 + v621
	goto L45
L105:
	;
	goto L44
L106:
	;
	v714 = v306
	v720 = v306
	goto L109
L107:
	;
	goto L108
L108:
	;
	v1044 = v103 + int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[1165])) = v1044
	v1061 = v1044
	v1066 = v181 + v108
	goto L9
L109:
	;
	v749 = v714 + int32(1)
	v750 = v714
	goto L111
L110:
	;
	goto L108
L111:
	;
	v784 = int32(1)
	v791 = *(*int32)(unsafe.Add(mBase, uint32(v83+v750<<(uint(int32(2))%32))))
	if v791 == int32(0) {
		v749 = v749 - v784
		v750 = v750 - v784
		goto L111
	} else {
		goto L113
	}
L112:
	;
	if v750 < int32(0) {
		goto L114
	} else {
		goto L115
	}
L113:
	;
	goto L112
L114:
	;
	v1544 = int32(-1)
	goto L1
L115:
	;
	goto L116
L116:
	;
	v798 = v750
	goto L117
L117:
	;
	v834 = v798 << (uint(int32(2)) % 32)
	v836 = *(*int32)(unsafe.Add(mBase, uint32(v83+v834)))
	if v836 != 0 {
		goto L120
	} else {
		goto L121
	}
L118:
	;
	v846 = *(*int32)(unsafe.Add(mBase, uint32(v836)+616))
	if v846 != 0 {
		goto L125
	} else {
		goto L126
	}
L119:
	;
	goto L118
L120:
	;
	v838 = *(*int32)(unsafe.Add(mBase, uint32(v834+v75)))
	if v838 == int32(0) {
		goto L119
	} else {
		goto L123
	}
L121:
	;
	goto L122
L122:
	;
	if int32(0) < v798 {
		v798 = v798 - int32(1)
		goto L117
	} else {
		goto L124
	}
L123:
	;
	goto L122
L124:
	;
	v1544 = int32(-1)
	goto L1
L125:
	;
	v847 = v846
	goto L127
L126:
	;
	v847 = v836
	goto L127
L127:
	;
	v848 = int32(0)
	v851 = v848
	v858 = v848
	goto L128
L128:
	;
	v888 = v83 + v851<<(uint(int32(2))%32)
	v889 = *(*int32)(unsafe.Add(mBase, uint32(v888)))
	if v847 != v889 {
		goto L131
	} else {
		goto L132
	}
L129:
	;
	v909 = *(*int32)(unsafe.Add(mBase, uint32(v834+v67)))
	if int32(0) < v909 {
		goto L137
	} else {
		goto L138
	}
L130:
	;
	v906 = v851 + int32(1)
	if v906 != v749 {
		v851 = v906
		v858 = v904
		goto L128
	} else {
		goto L136
	}
L131:
	;
	if v889 == int32(0) {
		v904 = v858
		goto L130
	} else {
		goto L134
	}
L132:
	;
	goto L133
L133:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v179+(v720-v858)<<(uint(int32(2))%32)))) = v889
	*(*int32)(unsafe.Add(mBase, uint32(v888))) = int32(0)
	v904 = v858 + int32(1)
	goto L130
L134:
	;
	v893 = *(*int32)(unsafe.Add(mBase, uint32(v889)+616))
	if v893 != v847 {
		v904 = v858
		goto L130
	} else {
		goto L135
	}
L135:
	;
	goto L133
L136:
	;
	goto L129
L137:
	;
	v913 = v909
	goto L140
L138:
	;
	goto L139
L139:
	;
	v1002 = v720 - v904
	if int32(0) <= v1002 {
		v714 = v750
		v720 = v1002
		goto L109
	} else {
		goto L143
	}
L140:
	;
	v950 = v65 + v913*int32(20)
	v953 = *(*int32)(unsafe.Add(mBase, uint32(v950-int32(8))))
	v956 = v75 + v953<<(uint(int32(2))%32)
	v957 = *(*int32)(unsafe.Add(mBase, uint32(v956)))
	*(*int32)(unsafe.Add(mBase, uint32(v956))) = v957 - int32(1)
	v963 = *(*int32)(unsafe.Add(mBase, uint32(v950-int32(4))))
	if int32(0) < v963 {
		v913 = v963
		goto L140
	} else {
		goto L142
	}
L141:
	;
	goto L139
L142:
	;
	goto L141
L143:
	;
	goto L110
L144:
	;
	goto L8
L145:
	;
	v1387 = int32(0)
	goto L5
L146:
	;
	goto L147
L147:
	;
	v1092 = v1086
	v1094 = int32(0)
	goto L148
L148:
	;
	v1128 = v1092 * int32(20)
	v1130 = *(*int32)(unsafe.Add(mBase, _consts[1163]))
	v1132 = *(*int32)(unsafe.Add(mBase, uint32(v1128+v1130)))
	v1134 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[1167])) = v1134
	*(*int32)(unsafe.Add(mBase, _consts[1166])) = v1134
	*(*int32)(unsafe.Add(mBase, uint32(v39)+12)) = v1134
	v1143 = v39 + int32(12)
	v1147 = *(*int32)(unsafe.Add(mBase, uint32(v1132)+616))
	if v1147 != 0 {
		goto L152
	} else {
		goto L153
	}
L149:
	;
	v1387 = v1378
	goto L5
L150:
	;
	v1255 = *(*int32)(unsafe.Add(mBase, _consts[1163]))
	v1257 = *(*int32)(unsafe.Add(mBase, uint32(v1255+v1128)+4))
	v1259 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[1167])) = v1259
	*(*int32)(unsafe.Add(mBase, _consts[1166])) = v1259
	*(*int32)(unsafe.Add(mBase, uint32(v39)+12)) = v1259
	v1268 = v39 + int32(12)
	v1272 = *(*int32)(unsafe.Add(mBase, uint32(v1257)+616))
	if v1272 != 0 {
		goto L186
	} else {
		goto L187
	}
L151:
	;
	if v1248 == int32(0) {
		v1253 = v1094
		goto L150
	} else {
		goto L182
	}
L152:
	;
	v1148 = v1147
	goto L154
L153:
	;
	v1148 = v1132
	goto L154
L154:
	;
	v1149 = int32(0)
	v1151 = *(*int32)(unsafe.Add(mBase, _consts[1173]))
	v1153 = *(*int32)(unsafe.Add(mBase, _consts[1167]))
	if v1149 < v1153 {
		goto L155
	} else {
		goto L156
	}
L155:
	;
	v1156 = v1149
	goto L158
L156:
	;
	goto L157
L157:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1167])) = v1153 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1151+v1153<<(uint(int32(2))%32)))) = v1148
	v1190 = *(*int32)(unsafe.Add(mBase, uint32(v1148)+4))
	if v1190 == int32(0) {
		goto L167
	} else {
		goto L168
	}
L158:
	;
	v1166 = *(*int32)(unsafe.Add(mBase, uint32(v1151+v1156<<(uint(int32(2))%32))))
	if v1148 == v1166 {
		goto L160
	} else {
		goto L161
	}
L159:
	;
	goto L157
L160:
	;
	if v1156 != 0 {
		goto L163
	} else {
		goto L164
	}
L161:
	;
	goto L162
L162:
	;
	v1173 = v1156 + int32(1)
	if v1173 != v1153 {
		v1156 = v1173
		goto L158
	} else {
		goto L166
	}
L163:
	;
	v1248 = int32(0)
	goto L151
L164:
	;
	goto L165
L165:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1166])) = v1134
	v1248 = int32(1)
	goto L151
L166:
	;
	goto L159
L167:
	;
	v1200 = *(*int32)(unsafe.Add(mBase, uint32(v1148)+624))
	if v1200 == int32(0) {
		goto L171
	} else {
		goto L172
	}
L168:
	;
	v1193 = *(*int32)(unsafe.Add(mBase, uint32(v1148)+92))
	if v1193 == int32(0) {
		goto L167
	} else {
		goto L169
	}
L169:
	;
	v1196 = F_FindLockCycleRecurseMember(m, v1148, v1148, v1134, v54, v1143)
	mBase = m.M
	if v1196 == int32(0) {
		goto L167
	} else {
		goto L170
	}
L170:
	;
	v1248 = int32(1)
	goto L151
L171:
	;
	v1248 = int32(0)
	goto L151
L172:
	;
	v1204 = v1148 + int32(620)
	if v1200 == v1204 {
		goto L171
	} else {
		goto L173
	}
L173:
	;
	v1206 = v1200
	goto L174
L174:
	;
	v1215 = *(*int32)(unsafe.Add(mBase, uint32(v1206-int32(624))))
	if v1215 == int32(0) {
		goto L176
	} else {
		goto L177
	}
L175:
	;
	goto L171
L176:
	;
	v1231 = *(*int32)(unsafe.Add(mBase, uint32(v1206)+4))
	if v1231 != v1204 {
		v1206 = v1231
		goto L174
	} else {
		goto L181
	}
L177:
	;
	v1220 = *(*int32)(unsafe.Add(mBase, uint32(v1206-int32(536))))
	if v1220 == int32(0) {
		goto L176
	} else {
		goto L178
	}
L178:
	;
	v1224 = v1206 - int32(628)
	if v1224 == v1148 {
		goto L176
	} else {
		goto L179
	}
L179:
	;
	v1226 = F_FindLockCycleRecurseMember(m, v1224, v1148, v1134, v54, v1143)
	mBase = m.M
	if v1226 == int32(0) {
		goto L176
	} else {
		goto L180
	}
L180:
	;
	v1248 = int32(1)
	goto L151
L181:
	;
	goto L175
L182:
	;
	v1251 = *(*int32)(unsafe.Add(mBase, uint32(v39)+12))
	if v1251 != 0 {
		v1253 = v1251
		goto L150
	} else {
		goto L183
	}
L183:
	;
	v1544 = int32(-1)
	goto L1
L184:
	;
	v1380 = v1092 + int32(1)
	v1382 = *(*int32)(unsafe.Add(mBase, _consts[1159]))
	if v1380 < v1382 {
		v1092 = v1380
		v1094 = v1378
		goto L148
	} else {
		goto L218
	}
L185:
	;
	if v1373 == int32(0) {
		v1378 = v1253
		goto L184
	} else {
		goto L216
	}
L186:
	;
	v1273 = v1272
	goto L188
L187:
	;
	v1273 = v1257
	goto L188
L188:
	;
	v1274 = int32(0)
	v1276 = *(*int32)(unsafe.Add(mBase, _consts[1173]))
	v1278 = *(*int32)(unsafe.Add(mBase, _consts[1167]))
	if v1274 < v1278 {
		goto L189
	} else {
		goto L190
	}
L189:
	;
	v1281 = v1274
	goto L192
L190:
	;
	goto L191
L191:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1167])) = v1278 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1276+v1278<<(uint(int32(2))%32)))) = v1273
	v1315 = *(*int32)(unsafe.Add(mBase, uint32(v1273)+4))
	if v1315 == int32(0) {
		goto L201
	} else {
		goto L202
	}
L192:
	;
	v1291 = *(*int32)(unsafe.Add(mBase, uint32(v1276+v1281<<(uint(int32(2))%32))))
	if v1273 == v1291 {
		goto L194
	} else {
		goto L195
	}
L193:
	;
	goto L191
L194:
	;
	if v1281 != 0 {
		goto L197
	} else {
		goto L198
	}
L195:
	;
	goto L196
L196:
	;
	v1298 = v1281 + int32(1)
	if v1298 != v1278 {
		v1281 = v1298
		goto L192
	} else {
		goto L200
	}
L197:
	;
	v1373 = int32(0)
	goto L185
L198:
	;
	goto L199
L199:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1166])) = v1259
	v1373 = int32(1)
	goto L185
L200:
	;
	goto L193
L201:
	;
	v1325 = *(*int32)(unsafe.Add(mBase, uint32(v1273)+624))
	if v1325 == int32(0) {
		goto L205
	} else {
		goto L206
	}
L202:
	;
	v1318 = *(*int32)(unsafe.Add(mBase, uint32(v1273)+92))
	if v1318 == int32(0) {
		goto L201
	} else {
		goto L203
	}
L203:
	;
	v1321 = F_FindLockCycleRecurseMember(m, v1273, v1273, v1259, v54, v1268)
	mBase = m.M
	if v1321 == int32(0) {
		goto L201
	} else {
		goto L204
	}
L204:
	;
	v1373 = int32(1)
	goto L185
L205:
	;
	v1373 = int32(0)
	goto L185
L206:
	;
	v1329 = v1273 + int32(620)
	if v1325 == v1329 {
		goto L205
	} else {
		goto L207
	}
L207:
	;
	v1331 = v1325
	goto L208
L208:
	;
	v1340 = *(*int32)(unsafe.Add(mBase, uint32(v1331-int32(624))))
	if v1340 == int32(0) {
		goto L210
	} else {
		goto L211
	}
L209:
	;
	goto L205
L210:
	;
	v1356 = *(*int32)(unsafe.Add(mBase, uint32(v1331)+4))
	if v1356 != v1329 {
		v1331 = v1356
		goto L208
	} else {
		goto L215
	}
L211:
	;
	v1345 = *(*int32)(unsafe.Add(mBase, uint32(v1331-int32(536))))
	if v1345 == int32(0) {
		goto L210
	} else {
		goto L212
	}
L212:
	;
	v1349 = v1331 - int32(628)
	if v1349 == v1273 {
		goto L210
	} else {
		goto L213
	}
L213:
	;
	v1351 = F_FindLockCycleRecurseMember(m, v1349, v1273, v1259, v54, v1268)
	mBase = m.M
	if v1351 == int32(0) {
		goto L210
	} else {
		goto L214
	}
L214:
	;
	v1373 = int32(1)
	goto L185
L215:
	;
	goto L209
L216:
	;
	v1376 = *(*int32)(unsafe.Add(mBase, uint32(v39)+12))
	if v1376 != 0 {
		v1378 = v1376
		goto L184
	} else {
		goto L217
	}
L217:
	;
	v1544 = int32(-1)
	goto L1
L218:
	;
	goto L149
L219:
	;
	if v1535 == int32(0) {
		v1544 = v1387
		goto L1
	} else {
		goto L250
	}
L220:
	;
	v1435 = v1434
	goto L222
L221:
	;
	v1435 = l0
	goto L222
L222:
	;
	v1436 = int32(0)
	v1438 = *(*int32)(unsafe.Add(mBase, _consts[1173]))
	v1440 = *(*int32)(unsafe.Add(mBase, _consts[1167]))
	if v1436 < v1440 {
		goto L223
	} else {
		goto L224
	}
L223:
	;
	v1443 = v1436
	goto L226
L224:
	;
	goto L225
L225:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1167])) = v1440 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1438+v1440<<(uint(int32(2))%32)))) = v1435
	v1477 = *(*int32)(unsafe.Add(mBase, uint32(v1435)+4))
	if v1477 == int32(0) {
		goto L235
	} else {
		goto L236
	}
L226:
	;
	v1453 = *(*int32)(unsafe.Add(mBase, uint32(v1438+v1443<<(uint(int32(2))%32))))
	if v1435 == v1453 {
		goto L228
	} else {
		goto L229
	}
L227:
	;
	goto L225
L228:
	;
	if v1443 != 0 {
		goto L231
	} else {
		goto L232
	}
L229:
	;
	goto L230
L230:
	;
	v1460 = v1443 + int32(1)
	if v1460 != v1440 {
		v1443 = v1460
		goto L226
	} else {
		goto L234
	}
L231:
	;
	v1535 = int32(0)
	goto L219
L232:
	;
	goto L233
L233:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1166])) = v1421
	v1535 = int32(1)
	goto L219
L234:
	;
	goto L227
L235:
	;
	v1487 = *(*int32)(unsafe.Add(mBase, uint32(v1435)+624))
	if v1487 == int32(0) {
		goto L239
	} else {
		goto L240
	}
L236:
	;
	v1480 = *(*int32)(unsafe.Add(mBase, uint32(v1435)+92))
	if v1480 == int32(0) {
		goto L235
	} else {
		goto L237
	}
L237:
	;
	v1483 = F_FindLockCycleRecurseMember(m, v1435, v1435, v1421, v54, v1430)
	mBase = m.M
	if v1483 == int32(0) {
		goto L235
	} else {
		goto L238
	}
L238:
	;
	v1535 = int32(1)
	goto L219
L239:
	;
	v1535 = int32(0)
	goto L219
L240:
	;
	v1491 = v1435 + int32(620)
	if v1487 == v1491 {
		goto L239
	} else {
		goto L241
	}
L241:
	;
	v1493 = v1487
	goto L242
L242:
	;
	v1502 = *(*int32)(unsafe.Add(mBase, uint32(v1493-int32(624))))
	if v1502 == int32(0) {
		goto L244
	} else {
		goto L245
	}
L243:
	;
	goto L239
L244:
	;
	v1518 = *(*int32)(unsafe.Add(mBase, uint32(v1493)+4))
	if v1518 != v1491 {
		v1493 = v1518
		goto L242
	} else {
		goto L249
	}
L245:
	;
	v1507 = *(*int32)(unsafe.Add(mBase, uint32(v1493-int32(536))))
	if v1507 == int32(0) {
		goto L244
	} else {
		goto L246
	}
L246:
	;
	v1511 = v1493 - int32(628)
	if v1511 == v1435 {
		goto L244
	} else {
		goto L247
	}
L247:
	;
	v1513 = F_FindLockCycleRecurseMember(m, v1511, v1435, v1421, v54, v1430)
	mBase = m.M
	if v1513 == int32(0) {
		goto L244
	} else {
		goto L248
	}
L248:
	;
	v1535 = int32(1)
	goto L219
L249:
	;
	goto L243
L250:
	;
	v1538 = *(*int32)(unsafe.Add(mBase, uint32(v39)+12))
	if v1538 != 0 {
		goto L251
	} else {
		goto L252
	}
L251:
	;
	v1540 = v1538
	goto L253
L252:
	;
	v1540 = int32(-1)
	goto L253
L253:
	;
	v1544 = v1540
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
	var v24 int32
	_ = v24
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
	v24 = v4
	goto L5
L5:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v26+v24<<(uint(int32(2))%32))))
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
	v76 = v24 + int32(1)
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v76 < v77 {
		v24 = v76
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
	v60 = F_transformWhereClause(m, v40, v56, int32(6), int32(536163))
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
	F_errmsg(m, int32(686364), v11+int32(16))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L14
	} else {
		goto L27
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = int32(84060)
	F_errdetail(m, int32(613894), v11)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L14
	} else {
		goto L28
	}
L28:
	;
	F_errfinish(m, int32(491199), int32(734), int32(160235))
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
	F_SimpleLruTruncate(m, int32(4381832), base.I64_extend_i32_u(int32(base.Ui32(v4)>>(uint(int32(11))%32))))
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
	v6 = F_DirectFunctionCall2Coll(m, int32(1559), v3, v4, v5)
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
	mBase := m.M
	_ = mBase
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
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
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
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = F_pg_detoast_datum_packed(m, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		v13 = v8 + int32(1)
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v15 = F_pg_detoast_datum_packed(m, v14)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8))))
			v18 = int32(1)
			v19 = v17 & v18
			if v17 == v18 {
				v22 = int32(4)
				v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
				if v24&int32(254) == int32(2) {
					v33 = v22
				} else {
					v33 = base.B2i32(v24 == int32(18)) << (uint(v22) % 32)
				}
				if v24 == int32(1) {
					v36 = v22
				} else {
					v36 = v33
				}
				v47 = v36
			} else {
				v37 = int32(1)
				if v19 != 0 {
					v47 = int32(base.Ui32(v17)>>(uint(v37)%32)) - v37
				} else {
					v41 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
					v47 = int32(base.Ui32(v41)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v50 = F_RE_compile_and_cache(m, v15, int32(19), v49)
			mBase = m.M
			v51 = m.ExcPending
			if v51 != 0 {
				return int32(0)
			} else {
				v56 = F_palloc(m, v47<<(uint(int32(2))%32)+int32(4))
				mBase = m.M
				v57 = m.ExcPending
				if v57 != 0 {
					return int32(0)
				} else {
					if v19 != 0 {
						v60 = v13
					} else {
						v60 = v8 + int32(4)
					}
					v61 = F_pg_mb2wchar_with_len(m, v60, v56, v47)
					mBase = m.M
					v62 = m.ExcPending
					if v62 != 0 {
						return int32(0)
					} else {
						v63 = int32(0)
						v66 = F_RE_wchar_execute(m, v56, v61, v63, v63, v63)
						mBase = m.M
						v67 = m.ExcPending
						if v67 != 0 {
							return int32(0)
						} else {
							F_pfree(m, v56)
							mBase = m.M
							v69 = m.ExcPending
							if v69 != 0 {
								return int32(0)
							} else {
								return v66
							}
						}
					}
				}
			}
		}
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
	var v41 int32
	_ = v41
	var v51 int32
	_ = v51
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
						if base.Ui32(int32(18)) < base.Ui32(v30) {
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
							if int32(1)<<(uint(v30)%32)&int32(262158) == int32(0) {
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
								v41 = int32(4)
								if v30&int32(254) == int32(2) {
									v51 = v41
								} else {
									v51 = base.B2i32(v30 == int32(18)) << (uint(v41) % 32)
								}
								if v30 == int32(1) {
									v54 = v41
								} else {
									v54 = v51
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
												F_errmsg(m, int32(671818), v11)
												mBase = m.M
												v105 = m.ExcPending
												if v105 != 0 {
													return int32(0)
												} else {
													F_errhint(m, int32(556344), int32(0))
													mBase = m.M
													v109 = m.ExcPending
													if v109 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(492464), int32(683), int32(417113))
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
												F_errmsg(m, int32(671818), v11)
												mBase = m.M
												v105 = m.ExcPending
												if v105 != 0 {
													return int32(0)
												} else {
													F_errhint(m, int32(556344), int32(0))
													mBase = m.M
													v109 = m.ExcPending
													if v109 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(492464), int32(683), int32(417113))
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
												F_errmsg(m, int32(671818), v11)
												mBase = m.M
												v105 = m.ExcPending
												if v105 != 0 {
													return int32(0)
												} else {
													F_errhint(m, int32(556344), int32(0))
													mBase = m.M
													v109 = m.ExcPending
													if v109 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(492464), int32(683), int32(417113))
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
									v64 = v48
								}
								v65 = F_replace_text_regexp(m, v15, v20, v23, v54, v55, v49-v56, v64)
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
										*(*int32)(unsafe.Add(mBase, uint32(v12))) = int32(82511)
										F_errmsg(m, int32(484897), v12)
										mBase = m.M
										v83 = m.ExcPending
										if v83 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(492464), int32(718), int32(458403))
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
								if v32&int32(65535) == int32(4) {
									v48 = v34
									v49 = v38
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
											v64 = v48
										}
										v65 = F_replace_text_regexp(m, v15, v20, v23, v54, v55, v49-v56, v64)
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
												*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = int32(283947)
												F_errmsg(m, int32(484897), v12+int32(16))
												mBase = m.M
												v103 = m.ExcPending
												if v103 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(492464), int32(727), int32(458403))
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
										v48 = v45
										v49 = v38
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
												v64 = v48
											}
											v65 = F_replace_text_regexp(m, v15, v20, v23, v54, v55, v49-v56, v64)
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
								v64 = v48
							}
							v65 = F_replace_text_regexp(m, v15, v20, v23, v54, v55, v49-v56, v64)
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
									*(*int32)(unsafe.Add(mBase, uint32(v12))) = int32(82511)
									F_errmsg(m, int32(484897), v12)
									mBase = m.M
									v83 = m.ExcPending
									if v83 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(492464), int32(718), int32(458403))
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
							if v32&int32(65535) == int32(4) {
								v48 = v34
								v49 = v38
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
										v64 = v48
									}
									v65 = F_replace_text_regexp(m, v15, v20, v23, v54, v55, v49-v56, v64)
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
											*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = int32(283947)
											F_errmsg(m, int32(484897), v12+int32(16))
											mBase = m.M
											v103 = m.ExcPending
											if v103 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(492464), int32(727), int32(458403))
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
									v48 = v45
									v49 = v38
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
											v64 = v48
										}
										v65 = F_replace_text_regexp(m, v15, v20, v23, v54, v55, v49-v56, v64)
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
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v147 int32
	_ = v147
	var v152 int32
	_ = v152
	var v157 int32
	_ = v157
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v232 int32
	_ = v232
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v254 int32
	_ = v254
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
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
	var v301 int32
	_ = v301
	var v308 int32
	_ = v308
	var v318 int32
	_ = v318
	var v322 int32
	_ = v322
	var v328 int32
	_ = v328
	var v336 int32
	_ = v336
	var v346 int32
	_ = v346
	var v348 int32
	_ = v348
	var v369 int32
	_ = v369
	var v379 int32
	_ = v379
	var v381 int32
	_ = v381
	var v398 int32
	_ = v398
	var v415 int32
	_ = v415
	var v425 int32
	_ = v425
	var v429 int32
	_ = v429
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v457 int32
	_ = v457
	var v459 int32
	_ = v459
	var v466 int32
	_ = v466
	var v483 int32
	_ = v483
	var v485 int32
	_ = v485
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v499 int32
	_ = v499
	var v502 int32
	_ = v502
	var v512 int32
	_ = v512
	var v517 int32
	_ = v517
	var v522 int32
	_ = v522
	var v526 int32
	_ = v526
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v538 int32
	_ = v538
	var v540 int32
	_ = v540
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v573 int32
	_ = v573
	var v588 int32
	_ = v588
	var v596 int32
	_ = v596
	var v600 int32
	_ = v600
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v616 int32
	_ = v616
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v629 int32
	_ = v629
	var v638 int32
	_ = v638
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
	var v656 int32
	_ = v656
	var v660 int32
	_ = v660
	var v661 int32
	_ = v661
	var v662 int32
	_ = v662
	var v663 int32
	_ = v663
	var v685 int32
	_ = v685
	var v692 int32
	_ = v692
	var v712 int32
	_ = v712
	var v720 int32
	_ = v720
	var v730 int32
	_ = v730
	var v732 int32
	_ = v732
	var v753 int32
	_ = v753
	var v763 int32
	_ = v763
	var v765 int32
	_ = v765
	var v782 int32
	_ = v782
	var v799 int32
	_ = v799
	var v867 int32
	_ = v867
	var v869 int32
	_ = v869
	var v879 int32
	_ = v879
	var v880 int32
	_ = v880
	var v881 int32
	_ = v881
	var v883 int32
	_ = v883
	var v886 int32
	_ = v886
	var v896 int32
	_ = v896
	var v915 int32
	_ = v915
	var v918 int32
	_ = v918
	var v920 int32
	_ = v920
	var v922 int32
	_ = v922
	var v923 int32
	_ = v923
	var v924 int32
	_ = v924
	var v927 int32
	_ = v927
	var v928 int32
	_ = v928
	var v932 int32
	_ = v932
	var v934 int32
	_ = v934
	var v946 int32
	_ = v946
	var v947 int32
	_ = v947
	var v949 int32
	_ = v949
	var v952 int32
	_ = v952
	var v953 int32
	_ = v953
	var v956 int32
	_ = v956
	var v959 int32
	_ = v959
	var v969 int32
	_ = v969
	var v971 int32
	_ = v971
	var v972 int32
	_ = v972
	var v977 int32
	_ = v977
	var v978 int32
	_ = v978
	var v979 int32
	_ = v979
	var v984 int32
	_ = v984
	var v997 int32
	_ = v997
	var v998 int32
	_ = v998
	var v999 int32
	_ = v999
	var v1001 int64
	_ = v1001
	var v1003 int32
	_ = v1003
	var v1005 int32
	_ = v1005
	var v1006 int32
	_ = v1006
	var v1007 int32
	_ = v1007
	var v1010 int32
	_ = v1010
	var v1011 int32
	_ = v1011
	var v1018 int32
	_ = v1018
	var v1039 int32
	_ = v1039
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
	return v1039
L2:
	;
	v918 = *(*int32)(unsafe.Add(mBase, uint32(v896)+12))
	v920 = base.B2i32(v918 == int32(0))
	v922 = v31 & int32(65535)
	v923 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
	v924 = *(*int32)(unsafe.Add(mBase, uint32(v896)))
	v927 = v923 + v924<<(uint(int32(3))%32)
	v928 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v927))))
	if v922 == v928 {
		goto L180
	} else {
		goto L181
	}
L3:
	;
	v915 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+1)) = uint8(v915)
	v1039 = v915
	goto L1
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+4)) = v896
	if v896 != 0 {
		goto L2
	} else {
		goto L176
	}
L5:
	;
	v532 = int32(0)
	v533 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
	if v533 == v532 {
		v548 = v532
		goto L112
	} else {
		goto L113
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v522 = m.ExcPending
	if v522 != 0 {
		goto L16
	} else {
		goto L109
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
		v1039 = v2
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
	v64 = v50
	v67 = v2
	goto L23
L23:
	;
	v72 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v64))))
	v74 = v64
	v75 = int32(0)
	goto L25
L24:
	;
	v896 = v512
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
	v102 = v97 & int32(65535)
	v105 = F_palloc(m, v102<<(uint(int32(2))%32))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L16
	} else {
		goto L32
	}
L27:
	;
	goto L26
L28:
	;
	v97 = v75
	v99 = v74
	goto L27
L29:
	;
	goto L30
L30:
	;
	v90 = v75 + int32(1)
	v94 = v74 + int32(8)
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v74+int32(12))))
	if v95 != 0 {
		v74 = v94
		v75 = v90
		goto L25
	} else {
		goto L31
	}
L31:
	;
	v97 = v90
	v99 = v94
	goto L27
L32:
	;
	if v102 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	v517 = *(*int32)(unsafe.Add(mBase, uint32(v99)+4))
	if v517 != 0 {
		v64 = v99
		v67 = v512
		goto L23
	} else {
		goto L108
	}
L34:
	;
	v186 = v102 & int32(3)
	v189 = v67
	goto L49
L35:
	;
	v109 = int32(0)
	goto L36
L36:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
	if v122 != 0 {
		goto L40
	} else {
		goto L41
	}
L37:
	;
	F_pfree(m, v105)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L16
	} else {
		goto L47
	}
L38:
	;
	goto L37
L39:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v135)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v105+v109<<(uint(int32(2))%32)))) = v147
	if v147 == int32(0) {
		goto L38
	} else {
		goto L45
	}
L40:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v64+v109<<(uint(int32(3))%32))+4))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+12)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = v126
	v130 = int32(8)
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
	v135 = F_bsearch(m, v16+v130, v132, v122, v130, int32(1166))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L16
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v105+v109<<(uint(int32(2))%32)))) = int32(0)
	goto L38
L43:
	;
	if v135 != 0 {
		goto L39
	} else {
		goto L44
	}
L44:
	;
	goto L42
L45:
	;
	v152 = v109 + int32(1)
	if v102 != v152 {
		v109 = v152
		goto L36
	} else {
		goto L46
	}
L46:
	;
	goto L34
L47:
	;
	v512 = v67
	goto L33
L48:
	;
	v512 = v189
	goto L33
L49:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v105)))
	v205 = int32(0)
	if v102 != 0 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v212 = v204
	v216 = v205
	goto L54
L52:
	;
	v318 = v204
	v322 = v205
	goto L53
L53:
	;
	if v102 != v322 {
		v398 = v189
		goto L82
	} else {
		goto L83
	}
L54:
	;
	v223 = v105 + v216<<(uint(int32(2))%32)
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v223)))
	if v224 == int32(0) {
		goto L57
	} else {
		goto L58
	}
L55:
	;
	v318 = v301
	v322 = v308
	goto L53
L56:
	;
	if v308 < v102 {
		v212 = v301
		v216 = v308
		goto L54
	} else {
		goto L81
	}
L57:
	;
	goto L48
L58:
	;
	v232 = v224
	goto L59
L59:
	;
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v232)))
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v212)))
	if base.Ui32(v242) < base.Ui32(v243) {
		goto L61
	} else {
		goto L62
	}
L60:
	;
	if base.Ui32(v243) < base.Ui32(v242) {
		v301 = v232
		v308 = int32(0)
		goto L56
	} else {
		goto L65
	}
L61:
	;
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v232)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v223))) = v245
	if v245 != 0 {
		v232 = v245
		goto L59
	} else {
		goto L64
	}
L62:
	;
	goto L63
L63:
	;
	goto L60
L64:
	;
	goto L57
L65:
	;
	v254 = v232
	goto L66
L66:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v254)))
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v212)))
	if v264 == v265 {
		goto L69
	} else {
		goto L70
	}
L67:
	;
	goto L57
L68:
	;
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v254)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v223))) = v279
	if v279 != 0 {
		v254 = v279
		goto L66
	} else {
		goto L80
	}
L69:
	;
	v267 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v254)+4)))
	if v267 != v31&int32(65535) {
		goto L68
	} else {
		goto L72
	}
L70:
	;
	v271 = v212
	v272 = v265
	goto L71
L71:
	;
	v276 = base.B2i32(v272 == v264)
	if v272 == v264 {
		goto L74
	} else {
		goto L75
	}
L72:
	;
	v269 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v254)+6)))
	if v102 != v269 {
		goto L68
	} else {
		goto L73
	}
L73:
	;
	v271 = v254
	v272 = v264
	goto L71
L74:
	;
	v277 = v216 + int32(1)
	goto L76
L75:
	;
	v277 = int32(0)
	goto L76
L76:
	;
	if v272 == v264 {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v278 = v271
	goto L79
L78:
	;
	v278 = v254
	goto L79
L79:
	;
	v301 = v278
	v308 = v277
	goto L56
L80:
	;
	goto L67
L81:
	;
	goto L55
L82:
	;
	if v102 == int32(0) {
		v189 = v398
		goto L49
	} else {
		goto L97
	}
L83:
	;
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v318)))
	if v26 == int32(0) {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	if v189 != 0 {
		goto L90
	} else {
		goto L91
	}
L85:
	;
	v336 = v26
	goto L86
L86:
	;
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v336)))
	if v346 == v328 {
		goto L84
	} else {
		goto L88
	}
L87:
	;
	v398 = v189
	goto L82
L88:
	;
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v336)+12))
	if v348 != 0 {
		v336 = v348
		goto L86
	} else {
		goto L89
	}
L89:
	;
	goto L87
L90:
	;
	v369 = v189
	goto L93
L91:
	;
	goto L92
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v318)+12)) = v189
	v398 = v318
	goto L82
L93:
	;
	v379 = *(*int32)(unsafe.Add(mBase, uint32(v369)))
	if v379 == v328 {
		v398 = v189
		goto L82
	} else {
		goto L95
	}
L94:
	;
	goto L92
L95:
	;
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v369)+12))
	if v381 != 0 {
		v369 = v381
		goto L93
	} else {
		goto L96
	}
L96:
	;
	goto L94
L97:
	;
	v415 = int32(0)
	if base.B2i32(base.Ui32(v102) < base.Ui32(int32(4))) == v415 {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v425 = v415
	v429 = v415
	goto L101
L99:
	;
	v466 = v415
	goto L100
L100:
	;
	if v186 == int32(0) {
		v189 = v398
		goto L49
	} else {
		goto L104
	}
L101:
	;
	v437 = v105 + v425<<(uint(int32(2))%32)
	v438 = *(*int32)(unsafe.Add(mBase, uint32(v437)))
	v439 = *(*int32)(unsafe.Add(mBase, uint32(v438)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v437))) = v439
	v441 = int32(4)
	v442 = v437 + v441
	v443 = *(*int32)(unsafe.Add(mBase, uint32(v442)))
	v444 = *(*int32)(unsafe.Add(mBase, uint32(v443)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v442))) = v444
	v447 = v437 + int32(8)
	v448 = *(*int32)(unsafe.Add(mBase, uint32(v447)))
	v449 = *(*int32)(unsafe.Add(mBase, uint32(v448)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v447))) = v449
	v452 = v437 + int32(12)
	v453 = *(*int32)(unsafe.Add(mBase, uint32(v452)))
	v454 = *(*int32)(unsafe.Add(mBase, uint32(v453)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v452))) = v454
	v457 = v425 + v441
	v459 = v429 + v441
	if v459 != v102&int32(65532) {
		v425 = v457
		v429 = v459
		goto L101
	} else {
		goto L103
	}
L102:
	;
	v466 = v457
	goto L100
L103:
	;
	goto L102
L104:
	;
	v483 = v466
	v485 = v415
	goto L105
L105:
	;
	v495 = v105 + v483<<(uint(int32(2))%32)
	v496 = *(*int32)(unsafe.Add(mBase, uint32(v495)))
	v497 = *(*int32)(unsafe.Add(mBase, uint32(v496)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v495))) = v497
	v499 = int32(1)
	v502 = v485 + v499
	if v502 != v186 {
		v483 = v483 + v499
		v485 = v502
		goto L105
	} else {
		goto L107
	}
L106:
	;
	v189 = v398
	goto L49
L107:
	;
	goto L106
L108:
	;
	goto L24
L109:
	;
	F_errmsg_internal(m, int32(303075), int32(0))
	mBase = m.M
	v526 = m.ExcPending
	if v526 != 0 {
		goto L16
	} else {
		goto L110
	}
L110:
	;
	F_errfinish(m, int32(490439), int32(799), int32(338330))
	mBase = m.M
	v531 = m.ExcPending
	if v531 != 0 {
		goto L16
	} else {
		goto L111
	}
L111:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = v548
	v555 = v16 + int32(4)
	v556 = int32(1)
	v573 = int32(0)
	goto L117
L113:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v16)+8)) = int64(0)
	v538 = int32(8)
	v540 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
	v543 = F_bsearch(m, v16+v538, v540, v533, v538, int32(1166))
	mBase = m.M
	v544 = m.ExcPending
	if v544 != 0 {
		goto L16
	} else {
		goto L114
	}
L114:
	;
	if v543 == int32(0) {
		v548 = v532
		goto L112
	} else {
		goto L115
	}
L115:
	;
	v547 = *(*int32)(unsafe.Add(mBase, uint32(v543)+4))
	v548 = v547
	goto L112
L116:
	;
	v896 = v573
	goto L4
L117:
	;
	v588 = *(*int32)(unsafe.Add(mBase, uint32(v555)))
	goto L119
L119:
	;
	v596 = v588
	v600 = int32(0)
	goto L122
L121:
	;
	if v556 != v692 {
		v782 = v573
		goto L150
	} else {
		goto L151
	}
L122:
	;
	v607 = v555 + v600<<(uint(int32(2))%32)
	v608 = *(*int32)(unsafe.Add(mBase, uint32(v607)))
	if v608 == int32(0) {
		goto L125
	} else {
		goto L126
	}
L123:
	;
	goto L121
L124:
	;
	if v692 < v556 {
		v596 = v685
		v600 = v692
		goto L122
	} else {
		goto L149
	}
L125:
	;
	goto L116
L126:
	;
	v616 = v608
	goto L127
L127:
	;
	v626 = *(*int32)(unsafe.Add(mBase, uint32(v616)))
	v627 = *(*int32)(unsafe.Add(mBase, uint32(v596)))
	if base.Ui32(v626) < base.Ui32(v627) {
		goto L129
	} else {
		goto L130
	}
L128:
	;
	if base.Ui32(v627) < base.Ui32(v626) {
		v685 = v616
		v692 = int32(0)
		goto L124
	} else {
		goto L133
	}
L129:
	;
	v629 = *(*int32)(unsafe.Add(mBase, uint32(v616)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v607))) = v629
	if v629 != 0 {
		v616 = v629
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
	v638 = v616
	goto L134
L134:
	;
	v648 = *(*int32)(unsafe.Add(mBase, uint32(v638)))
	v649 = *(*int32)(unsafe.Add(mBase, uint32(v596)))
	if v648 == v649 {
		goto L137
	} else {
		goto L138
	}
L135:
	;
	goto L125
L136:
	;
	v663 = *(*int32)(unsafe.Add(mBase, uint32(v638)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v607))) = v663
	if v663 != 0 {
		v638 = v663
		goto L134
	} else {
		goto L148
	}
L137:
	;
	v651 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v638)+4)))
	if v651 != v31&int32(65535) {
		goto L136
	} else {
		goto L140
	}
L138:
	;
	v655 = v596
	v656 = v649
	goto L139
L139:
	;
	v660 = base.B2i32(v656 == v648)
	if v656 == v648 {
		goto L142
	} else {
		goto L143
	}
L140:
	;
	v653 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v638)+6)))
	if v556 != v653 {
		goto L136
	} else {
		goto L141
	}
L141:
	;
	v655 = v638
	v656 = v648
	goto L139
L142:
	;
	v661 = v600 + int32(1)
	goto L144
L143:
	;
	v661 = int32(0)
	goto L144
L144:
	;
	if v656 == v648 {
		goto L145
	} else {
		goto L146
	}
L145:
	;
	v662 = v655
	goto L147
L146:
	;
	v662 = v638
	goto L147
L147:
	;
	v685 = v662
	v692 = v661
	goto L124
L148:
	;
	goto L135
L149:
	;
	goto L123
L150:
	;
	goto L165
L151:
	;
	v712 = *(*int32)(unsafe.Add(mBase, uint32(v685)))
	if v26 == int32(0) {
		goto L152
	} else {
		goto L153
	}
L152:
	;
	if v573 != 0 {
		goto L158
	} else {
		goto L159
	}
L153:
	;
	v720 = v26
	goto L154
L154:
	;
	v730 = *(*int32)(unsafe.Add(mBase, uint32(v720)))
	if v730 == v712 {
		goto L152
	} else {
		goto L156
	}
L155:
	;
	v782 = v573
	goto L150
L156:
	;
	v732 = *(*int32)(unsafe.Add(mBase, uint32(v720)+12))
	if v732 != 0 {
		v720 = v732
		goto L154
	} else {
		goto L157
	}
L157:
	;
	goto L155
L158:
	;
	v753 = v573
	goto L161
L159:
	;
	goto L160
L160:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v685)+12)) = v573
	v782 = v685
	goto L150
L161:
	;
	v763 = *(*int32)(unsafe.Add(mBase, uint32(v753)))
	if v763 == v712 {
		v782 = v573
		goto L150
	} else {
		goto L163
	}
L162:
	;
	goto L160
L163:
	;
	v765 = *(*int32)(unsafe.Add(mBase, uint32(v753)+12))
	if v765 != 0 {
		v753 = v765
		goto L161
	} else {
		goto L164
	}
L164:
	;
	goto L162
L165:
	;
	v799 = int32(0)
	goto L167
L167:
	;
	goto L168
L168:
	;
	goto L172
L172:
	;
	v867 = v799
	v869 = v799
	goto L173
L173:
	;
	v879 = v555 + v867<<(uint(int32(2))%32)
	v880 = *(*int32)(unsafe.Add(mBase, uint32(v879)))
	v881 = *(*int32)(unsafe.Add(mBase, uint32(v880)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v879))) = v881
	v883 = int32(1)
	v886 = v869 + v883
	if v886 != v556 {
		v867 = v867 + v883
		v869 = v886
		goto L173
	} else {
		goto L175
	}
L174:
	;
	v573 = v782
	goto L117
L175:
	;
	goto L174
L176:
	;
	goto L3
L177:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+1)) = uint8(v971)
	v1039 = v977
	goto L1
L178:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v977)+4)) = int32(0)
	goto L177
L179:
	;
	v1018 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+1)) = uint8(v1018)
	v1039 = int32(0)
	goto L1
L180:
	;
	v956 = v927
	v959 = v920
	v969 = int32(0)
	goto L182
L181:
	;
	v932 = v918
	v934 = v920
	goto L183
L182:
	;
	if v959 != 0 {
		goto L187
	} else {
		goto L188
	}
L183:
	;
	if v934&int32(1) != 0 {
		goto L179
	} else {
		goto L185
	}
L184:
	;
	v956 = v952
	v959 = v949
	v969 = int32(1)
	goto L182
L185:
	;
	v946 = *(*int32)(unsafe.Add(mBase, uint32(v932)))
	v947 = *(*int32)(unsafe.Add(mBase, uint32(v932)+12))
	v949 = base.B2i32(v947 == int32(0))
	v952 = v923 + v946<<(uint(int32(3))%32)
	v953 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v952))))
	if v953 != v922 {
		v932 = v947
		v934 = v949
		goto L183
	} else {
		goto L186
	}
L186:
	;
	goto L184
L187:
	;
	v971 = v969
	goto L189
L188:
	;
	v971 = int32(1)
	goto L189
L189:
	;
	v972 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v956)+2)))
	v977 = F_palloc(m, v972<<(uint(int32(3))%32)+int32(8))
	mBase = m.M
	v978 = m.ExcPending
	if v978 != 0 {
		goto L16
	} else {
		goto L190
	}
L190:
	;
	v979 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v956)+2)))
	if v979 == int32(0) {
		goto L178
	} else {
		goto L191
	}
L191:
	;
	v984 = int32(0)
	goto L192
L192:
	;
	v997 = v984 << (uint(int32(3)) % 32)
	v998 = v977 + v997
	v999 = *(*int32)(unsafe.Add(mBase, uint32(v956)+4))
	v1001 = *(*int64)(unsafe.Add(mBase, uint32(v999+v997)))
	*(*int64)(unsafe.Add(mBase, uint32(v998))) = v1001
	v1003 = *(*int32)(unsafe.Add(mBase, uint32(v956)+4))
	v1005 = *(*int32)(unsafe.Add(mBase, uint32(v1003+v997)+4))
	v1006 = F_pstrdup(m, v1005)
	mBase = m.M
	v1007 = m.ExcPending
	if v1007 != 0 {
		goto L16
	} else {
		goto L194
	}
L193:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v977+v1011<<(uint(int32(3))%32))+4)) = int32(0)
	goto L177
L194:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v998)+4)) = v1006
	v1010 = v984 + int32(1)
	v1011 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v956)+2)))
	if base.Ui32(v1010) < base.Ui32(v1011) {
		v984 = v1010
		goto L192
	} else {
		goto L195
	}
L195:
	;
	goto L193
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
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = F_pg_detoast_datum(m, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		v14 = F_ArrayGetIntegerTypmods(m, v8, v5+int32(12))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v5)+12))
			if v16 != int32(1) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(50856066))
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(221384), int32(0))
						mBase = m.M
						v29 = m.ExcPending
						if v29 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(492706), int32(118), int32(275828))
							mBase = m.M
							v34 = m.ExcPending
							if v34 != 0 {
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
				v36 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
				v37 = F_anytimestamp_typmod_check(m, int32(0), v36)
				mBase = m.M
				v38 = m.ExcPending
				if v38 != 0 {
					return int32(0)
				} else {
					m.G0 = v5 + int32(16)
					return v37
				}
			}
		}
	}
}
func F_timestamptztypmodin(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = F_pg_detoast_datum(m, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		v14 = F_ArrayGetIntegerTypmods(m, v8, v5+int32(12))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v5)+12))
			if v16 != int32(1) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(50856066))
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(221384), int32(0))
						mBase = m.M
						v29 = m.ExcPending
						if v29 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(492706), int32(118), int32(275828))
							mBase = m.M
							v34 = m.ExcPending
							if v34 != 0 {
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
				v36 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
				v37 = F_anytimestamp_typmod_check(m, int32(1), v36)
				mBase = m.M
				v38 = m.ExcPending
				if v38 != 0 {
					return int32(0)
				} else {
					m.G0 = v5 + int32(16)
					return v37
				}
			}
		}
	}
}
func F_timetypmodout(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if int32(0) <= v7 {
		*(*int32)(unsafe.Add(mBase, uint32(v5)+4)) = int32(370226)
		*(*int32)(unsafe.Add(mBase, uint32(v5))) = v7
		v14 = F_psprintf(m, int32(176454), v5)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			v21 = v14
			m.G0 = v5 + int32(16)
			return v21
		}
	} else {
		v19 = F_pstrdup(m, int32(370226))
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int32(0)
		} else {
			v21 = v19
			m.G0 = v5 + int32(16)
			return v21
		}
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
	var v36 int32
	_ = v36
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
		v36 = v3
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return v36
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
		v36 = v29
		goto L4
	} else {
		goto L11
	}
L10:
	;
	v36 = v29
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
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
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
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v63 int32
	_ = v63
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
	return v63
L5:
	;
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v10 = v9
	goto L7
L6:
	;
	v10 = v3
	goto L7
L7:
	;
	if v10 == v8 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v14 = v3
	goto L11
L9:
	;
	goto L10
L10:
	;
	v63 = int32(0)
	goto L4
L11:
	;
	v18 = int32(0)
	if l0 == v18 {
		v28 = v18
		goto L13
	} else {
		goto L14
	}
L12:
	;
	goto L10
L13:
	;
	v29 = int32(1)
	if l1 == int32(0) {
		v63 = v29
		goto L4
	} else {
		goto L16
	}
L14:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v22 <= v14 {
		v28 = int32(0)
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v28 = v24 + v14<<(uint(int32(2))%32)
	goto L13
L16:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v32 <= v14 {
		v63 = v29
		goto L4
	} else {
		goto L17
	}
L17:
	;
	if v28 == int32(0) {
		v63 = v29
		goto L4
	} else {
		goto L18
	}
L18:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v39 = v36 + v14<<(uint(int32(2))%32)
	if v39 == int32(0) {
		v63 = v29
		goto L4
	} else {
		goto L19
	}
L19:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	v48 = F_equal(m, v45, v47)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	return int32(0)
L21:
	;
	if v48 != 0 {
		v14 = v14 + int32(1)
		goto L11
	} else {
		goto L22
	}
L22:
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
	var v47 int32
	_ = v47
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v68 int32
	_ = v68
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
	F_ScanKeyInit(m, v6+int32(-48), int32(1), int32(3), int32(184), l1)
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
	v37 = F_systable_beginscan(m, l0, v31, v32, int32(4154240), v32, v6+int32(-48))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v39 = F_systable_getnext(m, v37)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	F_systable_endscan(m, v37)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	if int32(0) < v43 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v47 = int32(0)
	goto L10
L8:
	;
	goto L9
L9:
	;
	F_pfree(m, v26)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L1
	} else {
		goto L14
	}
L10:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v26+v47<<(uint(int32(2))%32))))
	F_relation_close(m, v55, int32(3))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L1
	} else {
		goto L12
	}
L11:
	;
	goto L9
L12:
	;
	v60 = v47 + int32(1)
	if v60 != v43 {
		v47 = v60
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
	return base.B2i32(v39 != int32(0))
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
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v63 int32
	_ = v63
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v129 int32
	_ = v129
	v2 = int32(1)
	if base.Ui32(int32(131071)) < base.Ui32(l0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v129
L2:
	;
	v129 = l0
	goto L1
L3:
	;
	v12 = int32(255)
	v13 = l0 & v12
	v14 = int32(3)
	v15 = base.I32_div_u_s(v13, v14)
	v21 = int32(2)
	v25 = *(*int32)(unsafe.Add(mBase, uint32((l0-v15*v14)&v12<<(uint(v21)%32))+uint32(_consts[1609])))
	v26 = int32(8)
	v27 = int32(base.Ui32(l0) >> (uint(v26) % 32))
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1610]))))
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30*int32(86)+v15)+uint32(_consts[1610]))))
	v41 = base.I32_rem_u_s(int32(base.Ui32(v25*v36)>>(uint(int32(11))%32)), int32(6))
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1611]))))
	v50 = *(*int32)(unsafe.Add(mBase, uint32((v41+v44)<<(uint(v21)%32))+uint32(_consts[1612])))
	v52 = v50 >> (uint(v26) % 32)
	v54 = v50 & v12
	if base.Ui32(v54) <= base.Ui32(int32(1)) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v129 = v52&(int32(0)-(v2^v54)) + l0
	goto L1
L5:
	;
	goto L6
L6:
	;
	v63 = v52 & int32(255)
	if v63 == int32(0) {
		goto L2
	} else {
		goto L7
	}
L7:
	;
	v70 = v63
	v71 = int32(base.Ui32(v52) >> (uint(int32(8)) % 32))
	goto L8
L8:
	;
	v77 = int32(1)
	v78 = int32(base.Ui32(v70) >> (uint(v77) % 32))
	v79 = v78 + v71
	v81 = v79 << (uint(v77) % 32)
	v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81)+uint32(_consts[1613]))))
	if v84 == v13 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	goto L2
L10:
	;
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81)+uint32(_consts[1614]))))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v86<<(uint(int32(2))%32))+uint32(_consts[1612])))
	v93 = v91 & int32(255)
	if base.Ui32(v93) <= base.Ui32(int32(1)) {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	goto L12
L12:
	;
	v107 = base.B2i32(base.Ui32(v13) < base.Ui32(v84))
	if base.Ui32(v13) < base.Ui32(v84) {
		goto L19
	} else {
		goto L20
	}
L13:
	;
	v129 = (int32(0)-(v2^v93))&(v91>>(uint(int32(8))%32)) + l0
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
	v129 = int32(-1) + l0
	goto L1
L19:
	;
	v108 = v71
	goto L21
L20:
	;
	v108 = v79
	goto L21
L21:
	;
	if base.Ui32(v13) < base.Ui32(v84) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v110 = v78
	goto L24
L23:
	;
	v110 = v70 - v78
	goto L24
L24:
	;
	if v110 != 0 {
		v70 = v110
		v71 = v108
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
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	if l1 == int32(0) {
		v44 = int32(0)
		return v44
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
					if l2 != int32(22) {
						v44 = v14
						return v44
					} else {
						if l4 != int32(1) {
							v44 = v14
							return v44
						} else {
							v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
							if v22 != int32(72) {
								v44 = v14
								return v44
							} else {
								v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+12)))
								if v25 != int32(1) {
									v44 = v14
									return v44
								} else {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v31 = m.ExcPending
									if v31 != 0 {
										return int32(0)
									} else {
										F_errcode(m, int32(654573698))
										mBase = m.M
										v34 = m.ExcPending
										if v34 != 0 {
											return int32(0)
										} else {
											F_errmsg(m, int32(356398), int32(0))
											mBase = m.M
											v38 = m.ExcPending
											if v38 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(495469), int32(1907), int32(357521))
												mBase = m.M
												v43 = m.ExcPending
												if v43 != 0 {
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
	var v71 int32
	_ = v71
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v112 int32
	_ = v112
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
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
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int64
	_ = v240
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v314 int32
	_ = v314
	var v317 int32
	_ = v317
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v329 int32
	_ = v329
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v337 int32
	_ = v337
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
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
	var v367 int32
	_ = v367
	var v371 int32
	_ = v371
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v388 int32
	_ = v388
	var v393 int32
	_ = v393
	var v397 int32
	_ = v397
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v405 int32
	_ = v405
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v422 int32
	_ = v422
	var v427 int32
	_ = v427
	var v431 int32
	_ = v431
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v439 int32
	_ = v439
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v456 int32
	_ = v456
	var v461 int32
	_ = v461
	var v465 int32
	_ = v465
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v473 int32
	_ = v473
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v490 int32
	_ = v490
	var v495 int32
	_ = v495
	var v499 int32
	_ = v499
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v507 int32
	_ = v507
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v524 int32
	_ = v524
	var v529 int32
	_ = v529
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v540 int32
	_ = v540
	var v545 int32
	_ = v545
	var v565 int32
	_ = v565
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v574 int32
	_ = v574
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v590 int32
	_ = v590
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v599 int32
	_ = v599
	var v608 int32
	_ = v608
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v618 int32
	_ = v618
	var v622 int32
	_ = v622
	var v623 int32
	_ = v623
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v631 int32
	_ = v631
	var v636 int32
	_ = v636
	var v637 int32
	_ = v637
	var v640 int32
	_ = v640
	var v643 int32
	_ = v643
	var v644 int32
	_ = v644
	var v646 int32
	_ = v646
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v658 int32
	_ = v658
	var v669 int32
	_ = v669
	var v671 int32
	_ = v671
	var v672 int32
	_ = v672
	var v674 int32
	_ = v674
	var v676 int32
	_ = v676
	var v678 int32
	_ = v678
	var v680 int32
	_ = v680
	var v683 int32
	_ = v683
	var v684 int32
	_ = v684
	var v691 int32
	_ = v691
	var v692 int32
	_ = v692
	var v693 int32
	_ = v693
	var v698 int32
	_ = v698
	var v699 int32
	_ = v699
	var v700 int32
	_ = v700
	var v701 int64
	_ = v701
	var v705 int32
	_ = v705
	var v706 int32
	_ = v706
	var v709 int32
	_ = v709
	var v714 int32
	_ = v714
	var v715 int32
	_ = v715
	var v718 int32
	_ = v718
	var v721 int32
	_ = v721
	var v722 int32
	_ = v722
	var v724 int32
	_ = v724
	var v732 int32
	_ = v732
	var v733 int32
	_ = v733
	var v736 int32
	_ = v736
	var v747 int32
	_ = v747
	var v749 int32
	_ = v749
	var v750 int32
	_ = v750
	var v752 int32
	_ = v752
	var v754 int32
	_ = v754
	var v756 int32
	_ = v756
	var v758 int32
	_ = v758
	var v761 int32
	_ = v761
	var v762 int32
	_ = v762
	var v769 int32
	_ = v769
	var v770 int32
	_ = v770
	var v771 int32
	_ = v771
	var v776 int32
	_ = v776
	var v779 int32
	_ = v779
	var v785 int32
	_ = v785
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
	v596 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	if v596 == int32(0) {
		goto L4
	} else {
		goto L184
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
	v565 = m.ExcPending
	if v565 != 0 {
		goto L1
	} else {
		goto L175
	}
L11:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	if v94 == int32(0) {
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
	v71 = v67 - int32(1)
	if base.Ui32(v71) <= base.Ui32(int32(3)) {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+128)) = v79
	F_errmsg(m, int32(162545), v17+int32(128))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L1
	} else {
		goto L22
	}
L19:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v71<<(uint(int32(2))%32))+uint32(_consts[383])))
	v79 = v78
	goto L21
L20:
	;
	v79 = int32(372185)
	goto L21
L21:
	;
	goto L18
L22:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v53)+24))
	F_parser_errposition(m, l0, v86)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	F_errfinish(m, int32(495167), int32(3539), int32(357659))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
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
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v94)+4))
	if v97 <= int32(0) {
		goto L25
	} else {
		goto L27
	}
L27:
	;
	v100 = int32(0)
	if v100 < v97 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v103 = v97
	goto L30
L29:
	;
	v103 = v100
	goto L30
L30:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v94)+12))
	v112 = int32(0)
	goto L31
L31:
	;
	v120 = int32(1)
	v121 = v112 + v120
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v104+v112<<(uint(int32(2))%32))))
	v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v125)+125)))
	if v126 != v120 {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	goto L25
L33:
	;
	if v103 != v121 {
		v112 = v121
		goto L31
	} else {
		goto L174
	}
L34:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v125)+8))
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v129)+4))
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v125)+4))
	if v131 != 0 {
		v139 = v130
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v53)+12))
	v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v140))))
	v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v139))))
	if v144 == int32(0) {
		v163 = v143
		v164 = v144
		goto L40
	} else {
		goto L41
	}
L36:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v125)+12))
	switch v132 - int32(1) {
	case 0, 4:
		goto L33
	case 1:
		goto L37
	default:
		v139 = v130
		goto L35
	}
L37:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v125)+64))
	if v135 == int32(0) {
		goto L33
	} else {
		goto L38
	}
L38:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v135)+4))
	v139 = v138
	goto L35
L39:
	;
	if v164-v163 != 0 {
		goto L33
	} else {
		goto L47
	}
L40:
	;
	goto L39
L41:
	;
	if v143 != v144 {
		v163 = v143
		v164 = v144
		goto L40
	} else {
		goto L42
	}
L42:
	;
	v148 = v139
	v149 = v140
	goto L43
L43:
	;
	v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149)+1)))
	v153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148)+1)))
	if v153 == int32(0) {
		v163 = v152
		v164 = v153
		goto L40
	} else {
		goto L45
	}
L44:
	;
	v163 = v152
	v164 = v153
	goto L40
L45:
	;
	v156 = int32(1)
	if v152 == v153 {
		v148 = v148 + v156
		v149 = v149 + v156
		goto L43
	} else {
		goto L46
	}
L46:
	;
	goto L44
L47:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v125)+12))
	switch v166 {
	case 0:
		goto L57
	case 1:
		goto L56
	case 2:
		goto L54
	case 3:
		goto L53
	case 4:
		goto L52
	case 5:
		goto L51
	case 6:
		goto L50
	case 7:
		goto L49
	default:
		goto L48
	}
L48:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v533 = m.ExcPending
	if v533 != 0 {
		goto L1
	} else {
		goto L171
	}
L49:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v499 = m.ExcPending
	if v499 != 0 {
		goto L1
	} else {
		goto L162
	}
L50:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v465 = m.ExcPending
	if v465 != 0 {
		goto L1
	} else {
		goto L153
	}
L51:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v431 = m.ExcPending
	if v431 != 0 {
		goto L1
	} else {
		goto L144
	}
L52:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
		goto L1
	} else {
		goto L135
	}
L53:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v363 = m.ExcPending
	if v363 != 0 {
		goto L1
	} else {
		goto L126
	}
L54:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L1
	} else {
		goto L117
	}
L55:
	;
	v323 = v46 + int32(1)
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	if v323 < v324 {
		v46 = v323
		goto L9
	} else {
		goto L116
	}
L56:
	;
	v244 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v245 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	if v4 == int32(0) {
		goto L87
	} else {
		goto L88
	}
L57:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v168 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	if v4 == int32(0) {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v171 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+43)) = uint8(v171)
	goto L60
L59:
	;
	goto L60
L60:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(l1)+140))
	if v176 != 0 {
		goto L64
	} else {
		goto L65
	}
L61:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v238 = F_getRTEPermissionInfo(m, v237, v125)
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L1
	} else {
		goto L86
	}
L62:
	;
	if v209 != 0 {
		goto L75
	} else {
		goto L76
	}
L63:
	;
	goto L62
L64:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v176)+4))
	if v177 <= int32(0) {
		v209 = int32(0)
		goto L63
	} else {
		goto L67
	}
L65:
	;
	goto L66
L66:
	;
	v209 = int32(0)
	goto L63
L67:
	;
	v180 = int32(0)
	if v180 < v177 {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v183 = v177
	goto L70
L69:
	;
	v183 = v180
	goto L70
L70:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v176)+12))
	v186 = int32(0)
	goto L71
L71:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v184+v186<<(uint(int32(2))%32))))
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v194)+4))
	if v195 == v121 {
		v209 = v194
		goto L63
	} else {
		goto L73
	}
L72:
	;
	goto L66
L73:
	;
	v198 = v186 + int32(1)
	if v198 != v183 {
		v186 = v198
		goto L71
	} else {
		goto L74
	}
L74:
	;
	goto L72
L75:
	;
	v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209)+16)))
	v212 = v4 & v211
	*(*uint8)(unsafe.Add(mBase, uint32(v209)+16)) = uint8(v212)
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v209)+8))
	if base.Ui32(v168) < base.Ui32(v214) {
		goto L78
	} else {
		goto L79
	}
L76:
	;
	goto L77
L77:
	;
	v223 = F_palloc0(m, int32(20))
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L1
	} else {
		goto L84
	}
L78:
	;
	v216 = v214
	goto L80
L79:
	;
	v216 = v168
	goto L80
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v209)+8)) = v216
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v209)+12))
	if base.Ui32(v167) < base.Ui32(v218) {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	v220 = v218
	goto L83
L82:
	;
	v220 = v167
	goto L83
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v209)+12)) = v220
	goto L61
L84:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v223)+16)) = uint8(v4)
	*(*int32)(unsafe.Add(mBase, uint32(v223)+12)) = v167
	*(*int32)(unsafe.Add(mBase, uint32(v223)+8)) = v168
	*(*int32)(unsafe.Add(mBase, uint32(v223)+4)) = v121
	*(*int32)(unsafe.Add(mBase, uint32(v223))) = int32(109)
	v231 = *(*int32)(unsafe.Add(mBase, uint32(l1)+140))
	v232 = F_lappend(m, v231, v223)
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+140)) = v232
	goto L61
L86:
	;
	v240 = *(*int64)(unsafe.Add(mBase, uint32(v238)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v238)+16)) = v240 | int64(4)
	goto L55
L87:
	;
	v248 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+43)) = uint8(v248)
	goto L89
L88:
	;
	goto L89
L89:
	;
	v253 = *(*int32)(unsafe.Add(mBase, uint32(l1)+140))
	if v253 != 0 {
		goto L93
	} else {
		goto L94
	}
L90:
	;
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v125)+36))
	F_transformLockingClause(m, l0, v314, v24, int32(1))
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L1
	} else {
		goto L115
	}
L91:
	;
	if v286 != 0 {
		goto L104
	} else {
		goto L105
	}
L92:
	;
	goto L91
L93:
	;
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v253)+4))
	if v254 <= int32(0) {
		v286 = int32(0)
		goto L92
	} else {
		goto L96
	}
L94:
	;
	goto L95
L95:
	;
	v286 = int32(0)
	goto L92
L96:
	;
	v257 = int32(0)
	if v257 < v254 {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	v260 = v254
	goto L99
L98:
	;
	v260 = v257
	goto L99
L99:
	;
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v253)+12))
	v263 = int32(0)
	goto L100
L100:
	;
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v261+v263<<(uint(int32(2))%32))))
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v271)+4))
	if v272 == v121 {
		v286 = v271
		goto L92
	} else {
		goto L102
	}
L101:
	;
	goto L95
L102:
	;
	v275 = v263 + int32(1)
	if v275 != v260 {
		v263 = v275
		goto L100
	} else {
		goto L103
	}
L103:
	;
	goto L101
L104:
	;
	v288 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v286)+16)))
	v289 = v4 & v288
	*(*uint8)(unsafe.Add(mBase, uint32(v286)+16)) = uint8(v289)
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v286)+8))
	if base.Ui32(v245) < base.Ui32(v291) {
		goto L107
	} else {
		goto L108
	}
L105:
	;
	goto L106
L106:
	;
	v300 = F_palloc0(m, int32(20))
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L1
	} else {
		goto L113
	}
L107:
	;
	v293 = v291
	goto L109
L108:
	;
	v293 = v245
	goto L109
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v286)+8)) = v293
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v286)+12))
	if base.Ui32(v244) < base.Ui32(v295) {
		goto L110
	} else {
		goto L111
	}
L110:
	;
	v297 = v295
	goto L112
L111:
	;
	v297 = v244
	goto L112
L112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v286)+12)) = v297
	goto L90
L113:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v300)+16)) = uint8(v4)
	*(*int32)(unsafe.Add(mBase, uint32(v300)+12)) = v244
	*(*int32)(unsafe.Add(mBase, uint32(v300)+8)) = v245
	*(*int32)(unsafe.Add(mBase, uint32(v300)+4)) = v121
	*(*int32)(unsafe.Add(mBase, uint32(v300))) = int32(109)
	v308 = *(*int32)(unsafe.Add(mBase, uint32(l1)+140))
	v309 = F_lappend(m, v308, v300)
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L1
	} else {
		goto L114
	}
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+140)) = v309
	goto L90
L115:
	;
	goto L55
L116:
	;
	goto L4
L117:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L1
	} else {
		goto L118
	}
L118:
	;
	v333 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v337 = v333 - int32(1)
	if base.Ui32(v337) <= base.Ui32(int32(3)) {
		goto L120
	} else {
		goto L121
	}
L119:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+32)) = v345
	F_errmsg(m, int32(274476), v17+int32(32))
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L1
	} else {
		goto L123
	}
L120:
	;
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v337<<(uint(int32(2))%32))+uint32(_consts[383])))
	v345 = v344
	goto L122
L121:
	;
	v345 = int32(372185)
	goto L122
L122:
	;
	goto L119
L123:
	;
	v352 = *(*int32)(unsafe.Add(mBase, uint32(v53)+24))
	F_parser_errposition(m, l0, v352)
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L1
	} else {
		goto L124
	}
L124:
	;
	F_errfinish(m, int32(495167), int32(3603), int32(357659))
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L1
	} else {
		goto L125
	}
L125:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L126:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L1
	} else {
		goto L127
	}
L127:
	;
	v367 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v371 = v367 - int32(1)
	if base.Ui32(v371) <= base.Ui32(int32(3)) {
		goto L129
	} else {
		goto L130
	}
L128:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+48)) = v379
	F_errmsg(m, int32(252718), v17+int32(48))
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L1
	} else {
		goto L132
	}
L129:
	;
	v378 = *(*int32)(unsafe.Add(mBase, uint32(v371<<(uint(int32(2))%32))+uint32(_consts[383])))
	v379 = v378
	goto L131
L130:
	;
	v379 = int32(372185)
	goto L131
L131:
	;
	goto L128
L132:
	;
	v386 = *(*int32)(unsafe.Add(mBase, uint32(v53)+24))
	F_parser_errposition(m, l0, v386)
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L1
	} else {
		goto L133
	}
L133:
	;
	F_errfinish(m, int32(495167), int32(3612), int32(357659))
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L1
	} else {
		goto L134
	}
L134:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L135:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L1
	} else {
		goto L136
	}
L136:
	;
	v401 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v405 = v401 - int32(1)
	if base.Ui32(v405) <= base.Ui32(int32(3)) {
		goto L138
	} else {
		goto L139
	}
L137:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+64)) = v413
	F_errmsg(m, int32(252422), v17-int32(-64))
	mBase = m.M
	v419 = m.ExcPending
	if v419 != 0 {
		goto L1
	} else {
		goto L141
	}
L138:
	;
	v412 = *(*int32)(unsafe.Add(mBase, uint32(v405<<(uint(int32(2))%32))+uint32(_consts[383])))
	v413 = v412
	goto L140
L139:
	;
	v413 = int32(372185)
	goto L140
L140:
	;
	goto L137
L141:
	;
	v420 = *(*int32)(unsafe.Add(mBase, uint32(v53)+24))
	F_parser_errposition(m, l0, v420)
	mBase = m.M
	v422 = m.ExcPending
	if v422 != 0 {
		goto L1
	} else {
		goto L142
	}
L142:
	;
	F_errfinish(m, int32(495167), int32(3621), int32(357659))
	mBase = m.M
	v427 = m.ExcPending
	if v427 != 0 {
		goto L1
	} else {
		goto L143
	}
L143:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L144:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v434 = m.ExcPending
	if v434 != 0 {
		goto L1
	} else {
		goto L145
	}
L145:
	;
	v435 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v439 = v435 - int32(1)
	if base.Ui32(v439) <= base.Ui32(int32(3)) {
		goto L147
	} else {
		goto L148
	}
L146:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+80)) = v447
	F_errmsg(m, int32(519986), v17+int32(80))
	mBase = m.M
	v453 = m.ExcPending
	if v453 != 0 {
		goto L1
	} else {
		goto L150
	}
L147:
	;
	v446 = *(*int32)(unsafe.Add(mBase, uint32(v439<<(uint(int32(2))%32))+uint32(_consts[383])))
	v447 = v446
	goto L149
L148:
	;
	v447 = int32(372185)
	goto L149
L149:
	;
	goto L146
L150:
	;
	v454 = *(*int32)(unsafe.Add(mBase, uint32(v53)+24))
	F_parser_errposition(m, l0, v454)
	mBase = m.M
	v456 = m.ExcPending
	if v456 != 0 {
		goto L1
	} else {
		goto L151
	}
L151:
	;
	F_errfinish(m, int32(495167), int32(3630), int32(357659))
	mBase = m.M
	v461 = m.ExcPending
	if v461 != 0 {
		goto L1
	} else {
		goto L152
	}
L152:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L153:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v468 = m.ExcPending
	if v468 != 0 {
		goto L1
	} else {
		goto L154
	}
L154:
	;
	v469 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v473 = v469 - int32(1)
	if base.Ui32(v473) <= base.Ui32(int32(3)) {
		goto L156
	} else {
		goto L157
	}
L155:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+96)) = v481
	F_errmsg(m, int32(17020), v17+int32(96))
	mBase = m.M
	v487 = m.ExcPending
	if v487 != 0 {
		goto L1
	} else {
		goto L159
	}
L156:
	;
	v480 = *(*int32)(unsafe.Add(mBase, uint32(v473<<(uint(int32(2))%32))+uint32(_consts[383])))
	v481 = v480
	goto L158
L157:
	;
	v481 = int32(372185)
	goto L158
L158:
	;
	goto L155
L159:
	;
	v488 = *(*int32)(unsafe.Add(mBase, uint32(v53)+24))
	F_parser_errposition(m, l0, v488)
	mBase = m.M
	v490 = m.ExcPending
	if v490 != 0 {
		goto L1
	} else {
		goto L160
	}
L160:
	;
	F_errfinish(m, int32(495167), int32(3639), int32(357659))
	mBase = m.M
	v495 = m.ExcPending
	if v495 != 0 {
		goto L1
	} else {
		goto L161
	}
L161:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L162:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v502 = m.ExcPending
	if v502 != 0 {
		goto L1
	} else {
		goto L163
	}
L163:
	;
	v503 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v507 = v503 - int32(1)
	if base.Ui32(v507) <= base.Ui32(int32(3)) {
		goto L165
	} else {
		goto L166
	}
L164:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+112)) = v515
	F_errmsg(m, int32(362069), v17+int32(112))
	mBase = m.M
	v521 = m.ExcPending
	if v521 != 0 {
		goto L1
	} else {
		goto L168
	}
L165:
	;
	v514 = *(*int32)(unsafe.Add(mBase, uint32(v507<<(uint(int32(2))%32))+uint32(_consts[383])))
	v515 = v514
	goto L167
L166:
	;
	v515 = int32(372185)
	goto L167
L167:
	;
	goto L164
L168:
	;
	v522 = *(*int32)(unsafe.Add(mBase, uint32(v53)+24))
	F_parser_errposition(m, l0, v522)
	mBase = m.M
	v524 = m.ExcPending
	if v524 != 0 {
		goto L1
	} else {
		goto L169
	}
L169:
	;
	F_errfinish(m, int32(495167), int32(3648), int32(357659))
	mBase = m.M
	v529 = m.ExcPending
	if v529 != 0 {
		goto L1
	} else {
		goto L170
	}
L170:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L171:
	;
	v534 = *(*int32)(unsafe.Add(mBase, uint32(v125)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v534
	F_errmsg_internal(m, int32(483081), v17+int32(16))
	mBase = m.M
	v540 = m.ExcPending
	if v540 != 0 {
		goto L1
	} else {
		goto L172
	}
L172:
	;
	F_errfinish(m, int32(495167), int32(3655), int32(357659))
	mBase = m.M
	v545 = m.ExcPending
	if v545 != 0 {
		goto L1
	} else {
		goto L173
	}
L173:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L174:
	;
	goto L32
L175:
	;
	F_errcode(m, int32(16908420))
	mBase = m.M
	v568 = m.ExcPending
	if v568 != 0 {
		goto L1
	} else {
		goto L176
	}
L176:
	;
	v569 = *(*int32)(unsafe.Add(mBase, uint32(v53)+12))
	v570 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v574 = v570 - int32(1)
	if base.Ui32(v574) <= base.Ui32(int32(3)) {
		goto L178
	} else {
		goto L179
	}
L177:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+4)) = v582
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v569
	F_errmsg(m, int32(356757), v17)
	mBase = m.M
	v587 = m.ExcPending
	if v587 != 0 {
		goto L1
	} else {
		goto L181
	}
L178:
	;
	v581 = *(*int32)(unsafe.Add(mBase, uint32(v574<<(uint(int32(2))%32))+uint32(_consts[383])))
	v582 = v581
	goto L180
L179:
	;
	v582 = int32(372185)
	goto L180
L180:
	;
	goto L177
L181:
	;
	v588 = *(*int32)(unsafe.Add(mBase, uint32(v53)+24))
	F_parser_errposition(m, l0, v588)
	mBase = m.M
	v590 = m.ExcPending
	if v590 != 0 {
		goto L1
	} else {
		goto L182
	}
L182:
	;
	F_errfinish(m, int32(495167), int32(3669), int32(357659))
	mBase = m.M
	v595 = m.ExcPending
	if v595 != 0 {
		goto L1
	} else {
		goto L183
	}
L183:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L184:
	;
	v599 = *(*int32)(unsafe.Add(mBase, uint32(v596)+4))
	if v599 <= int32(0) {
		goto L4
	} else {
		goto L185
	}
L185:
	;
	v608 = v5
	goto L186
L186:
	;
	v616 = int32(1)
	v617 = v608 + v616
	v618 = *(*int32)(unsafe.Add(mBase, uint32(v596)+12))
	v622 = *(*int32)(unsafe.Add(mBase, uint32(v618+v608<<(uint(int32(2))%32))))
	v623 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v622)+125)))
	if v623 != v616 {
		goto L188
	} else {
		goto L189
	}
L187:
	;
	goto L4
L188:
	;
	v785 = *(*int32)(unsafe.Add(mBase, uint32(v596)+4))
	if v617 < v785 {
		v608 = v617
		goto L186
	} else {
		goto L250
	}
L189:
	;
	v626 = *(*int32)(unsafe.Add(mBase, uint32(v622)+12))
	switch v626 {
	case 0:
		goto L191
	case 1:
		goto L190
	default:
		goto L188
	}
L190:
	;
	v705 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v706 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	if v4 == int32(0) {
		goto L221
	} else {
		goto L222
	}
L191:
	;
	v627 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v628 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	if v4 == int32(0) {
		goto L192
	} else {
		goto L193
	}
L192:
	;
	v631 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+43)) = uint8(v631)
	goto L194
L193:
	;
	goto L194
L194:
	;
	v636 = *(*int32)(unsafe.Add(mBase, uint32(l1)+140))
	if v636 != 0 {
		goto L198
	} else {
		goto L199
	}
L195:
	;
	v698 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v699 = F_getRTEPermissionInfo(m, v698, v622)
	mBase = m.M
	v700 = m.ExcPending
	if v700 != 0 {
		goto L1
	} else {
		goto L220
	}
L196:
	;
	if v669 != 0 {
		goto L209
	} else {
		goto L210
	}
L197:
	;
	goto L196
L198:
	;
	v637 = *(*int32)(unsafe.Add(mBase, uint32(v636)+4))
	if v637 <= int32(0) {
		v669 = int32(0)
		goto L197
	} else {
		goto L201
	}
L199:
	;
	goto L200
L200:
	;
	v669 = int32(0)
	goto L197
L201:
	;
	v640 = int32(0)
	if v640 < v637 {
		goto L202
	} else {
		goto L203
	}
L202:
	;
	v643 = v637
	goto L204
L203:
	;
	v643 = v640
	goto L204
L204:
	;
	v644 = *(*int32)(unsafe.Add(mBase, uint32(v636)+12))
	v646 = int32(0)
	goto L205
L205:
	;
	v654 = *(*int32)(unsafe.Add(mBase, uint32(v644+v646<<(uint(int32(2))%32))))
	v655 = *(*int32)(unsafe.Add(mBase, uint32(v654)+4))
	if v655 == v617 {
		v669 = v654
		goto L197
	} else {
		goto L207
	}
L206:
	;
	goto L200
L207:
	;
	v658 = v646 + int32(1)
	if v658 != v643 {
		v646 = v658
		goto L205
	} else {
		goto L208
	}
L208:
	;
	goto L206
L209:
	;
	v671 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v669)+16)))
	v672 = v4 & v671
	*(*uint8)(unsafe.Add(mBase, uint32(v669)+16)) = uint8(v672)
	v674 = *(*int32)(unsafe.Add(mBase, uint32(v669)+8))
	if base.Ui32(v628) < base.Ui32(v674) {
		goto L212
	} else {
		goto L213
	}
L210:
	;
	goto L211
L211:
	;
	v683 = F_palloc0(m, int32(20))
	mBase = m.M
	v684 = m.ExcPending
	if v684 != 0 {
		goto L1
	} else {
		goto L218
	}
L212:
	;
	v676 = v674
	goto L214
L213:
	;
	v676 = v628
	goto L214
L214:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v669)+8)) = v676
	v678 = *(*int32)(unsafe.Add(mBase, uint32(v669)+12))
	if base.Ui32(v627) < base.Ui32(v678) {
		goto L215
	} else {
		goto L216
	}
L215:
	;
	v680 = v678
	goto L217
L216:
	;
	v680 = v627
	goto L217
L217:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v669)+12)) = v680
	goto L195
L218:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v683)+16)) = uint8(v4)
	*(*int32)(unsafe.Add(mBase, uint32(v683)+12)) = v627
	*(*int32)(unsafe.Add(mBase, uint32(v683)+8)) = v628
	*(*int32)(unsafe.Add(mBase, uint32(v683)+4)) = v617
	*(*int32)(unsafe.Add(mBase, uint32(v683))) = int32(109)
	v691 = *(*int32)(unsafe.Add(mBase, uint32(l1)+140))
	v692 = F_lappend(m, v691, v683)
	mBase = m.M
	v693 = m.ExcPending
	if v693 != 0 {
		goto L1
	} else {
		goto L219
	}
L219:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+140)) = v692
	goto L195
L220:
	;
	v701 = *(*int64)(unsafe.Add(mBase, uint32(v699)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v699)+16)) = v701 | int64(4)
	goto L188
L221:
	;
	v709 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+43)) = uint8(v709)
	goto L223
L222:
	;
	goto L223
L223:
	;
	v714 = *(*int32)(unsafe.Add(mBase, uint32(l1)+140))
	if v714 != 0 {
		goto L227
	} else {
		goto L228
	}
L224:
	;
	v776 = *(*int32)(unsafe.Add(mBase, uint32(v622)+36))
	F_transformLockingClause(m, l0, v776, v24, int32(1))
	mBase = m.M
	v779 = m.ExcPending
	if v779 != 0 {
		goto L1
	} else {
		goto L249
	}
L225:
	;
	if v747 != 0 {
		goto L238
	} else {
		goto L239
	}
L226:
	;
	goto L225
L227:
	;
	v715 = *(*int32)(unsafe.Add(mBase, uint32(v714)+4))
	if v715 <= int32(0) {
		v747 = int32(0)
		goto L226
	} else {
		goto L230
	}
L228:
	;
	goto L229
L229:
	;
	v747 = int32(0)
	goto L226
L230:
	;
	v718 = int32(0)
	if v718 < v715 {
		goto L231
	} else {
		goto L232
	}
L231:
	;
	v721 = v715
	goto L233
L232:
	;
	v721 = v718
	goto L233
L233:
	;
	v722 = *(*int32)(unsafe.Add(mBase, uint32(v714)+12))
	v724 = int32(0)
	goto L234
L234:
	;
	v732 = *(*int32)(unsafe.Add(mBase, uint32(v722+v724<<(uint(int32(2))%32))))
	v733 = *(*int32)(unsafe.Add(mBase, uint32(v732)+4))
	if v733 == v617 {
		v747 = v732
		goto L226
	} else {
		goto L236
	}
L235:
	;
	goto L229
L236:
	;
	v736 = v724 + int32(1)
	if v736 != v721 {
		v724 = v736
		goto L234
	} else {
		goto L237
	}
L237:
	;
	goto L235
L238:
	;
	v749 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v747)+16)))
	v750 = v4 & v749
	*(*uint8)(unsafe.Add(mBase, uint32(v747)+16)) = uint8(v750)
	v752 = *(*int32)(unsafe.Add(mBase, uint32(v747)+8))
	if base.Ui32(v706) < base.Ui32(v752) {
		goto L241
	} else {
		goto L242
	}
L239:
	;
	goto L240
L240:
	;
	v761 = F_palloc0(m, int32(20))
	mBase = m.M
	v762 = m.ExcPending
	if v762 != 0 {
		goto L1
	} else {
		goto L247
	}
L241:
	;
	v754 = v752
	goto L243
L242:
	;
	v754 = v706
	goto L243
L243:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v747)+8)) = v754
	v756 = *(*int32)(unsafe.Add(mBase, uint32(v747)+12))
	if base.Ui32(v705) < base.Ui32(v756) {
		goto L244
	} else {
		goto L245
	}
L244:
	;
	v758 = v756
	goto L246
L245:
	;
	v758 = v705
	goto L246
L246:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v747)+12)) = v758
	goto L224
L247:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v761)+16)) = uint8(v4)
	*(*int32)(unsafe.Add(mBase, uint32(v761)+12)) = v705
	*(*int32)(unsafe.Add(mBase, uint32(v761)+8)) = v706
	*(*int32)(unsafe.Add(mBase, uint32(v761)+4)) = v617
	*(*int32)(unsafe.Add(mBase, uint32(v761))) = int32(109)
	v769 = *(*int32)(unsafe.Add(mBase, uint32(l1)+140))
	v770 = F_lappend(m, v769, v761)
	mBase = m.M
	v771 = m.ExcPending
	if v771 != 0 {
		goto L1
	} else {
		goto L248
	}
L248:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+140)) = v770
	goto L224
L249:
	;
	goto L188
L250:
	;
	goto L187
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
			F_sequence_close(m, v18, int32(0))
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
				F_sequence_close(m, v18, int32(0))
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
					F_sequence_close(m, v18, int32(0))
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
	var v20 int32
	_ = v20
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
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v109 int32
	_ = v109
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
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v274 int32
	_ = v274
	var v287 int32
	_ = v287
	var v292 int32
	_ = v292
	var v296 int32
	_ = v296
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v316 int32
	_ = v316
	var v319 int32
	_ = v319
	var v326 int32
	_ = v326
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v348 int32
	_ = v348
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v379 int32
	_ = v379
	var v384 int32
	_ = v384
	var v393 int32
	_ = v393
	var v406 int32
	_ = v406
	var v425 int32
	_ = v425
	var v428 int32
	_ = v428
	var v432 int32
	_ = v432
	var v437 int32
	_ = v437
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v21 = F_pg_detoast_datum_packed(m, v20)
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
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v26 = F_pg_detoast_datum_packed(m, v25)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v29 = F_pg_detoast_datum_packed(m, v28)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
	if v31 == int32(1) {
		goto L9
	} else {
		goto L10
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v425 = m.ExcPending
	if v425 != 0 {
		goto L1
	} else {
		goto L102
	}
L6:
	;
	return v406
L7:
	;
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26))))
	if v64 == int32(1) {
		goto L16
	} else {
		goto L17
	}
L8:
	;
	if v59 <= int32(0) {
		v406 = v21
		goto L6
	} else {
		goto L14
	}
L9:
	;
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+1)))
	if base.Ui32((v35-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		v63 = int32(4)
		goto L7
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v46 = int32(1)
	if v31&v46 != 0 {
		v59 = int32(base.Ui32(v31)>>(uint(v46)%32)) - v46
		goto L8
	} else {
		goto L13
	}
L12:
	;
	v59 = base.B2i32(v35 == int32(18)) << (uint(int32(4)) % 32)
	goto L8
L13:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	v59 = int32(base.Ui32(v52)>>(uint(int32(2))%32)) - int32(4)
	goto L8
L14:
	;
	v63 = v59
	goto L7
L15:
	;
	v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29))))
	if v95 == int32(1) {
		goto L27
	} else {
		goto L28
	}
L16:
	;
	v67 = int32(4)
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+1)))
	if v69&int32(254) == int32(2) {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	goto L18
L18:
	;
	v82 = int32(1)
	if v64&v82 != 0 {
		v94 = int32(base.Ui32(v64)>>(uint(v82)%32)) - v82
		goto L15
	} else {
		goto L25
	}
L19:
	;
	v78 = v67
	goto L21
L20:
	;
	v78 = base.B2i32(v69 == int32(18)) << (uint(v67) % 32)
	goto L21
L21:
	;
	if v69 == int32(1) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v81 = v67
	goto L24
L23:
	;
	v81 = v78
	goto L24
L24:
	;
	v94 = v81
	goto L15
L25:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	v94 = int32(base.Ui32(v88)>>(uint(int32(2))%32)) - int32(4)
	goto L15
L26:
	;
	v127 = *(*int32)(unsafe.Add(mBase, _consts[486]))
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v127)+4))
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v128*int32(28))+uint32(_consts[1295])))
	goto L37
L27:
	;
	v98 = int32(4)
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+1)))
	if v100&int32(254) == int32(2) {
		goto L30
	} else {
		goto L31
	}
L28:
	;
	goto L29
L29:
	;
	v113 = int32(1)
	if v95&v113 != 0 {
		v125 = int32(base.Ui32(v95)>>(uint(v113)%32)) - v113
		goto L26
	} else {
		goto L36
	}
L30:
	;
	v109 = v98
	goto L32
L31:
	;
	v109 = base.B2i32(v100 == int32(18)) << (uint(v98) % 32)
	goto L32
L32:
	;
	if v100 == int32(1) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v112 = v98
	goto L35
L34:
	;
	v112 = v109
	goto L35
L35:
	;
	v125 = v112
	goto L26
L36:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
	v125 = int32(base.Ui32(v119)>>(uint(int32(2))%32)) - int32(4)
	goto L26
L37:
	;
	v136 = base.I64_extend_i32_s(v133) * base.I64_extend_i32_s(v63)
	v140 = base.I32_wrap_i64(v136)
	if base.I32_wrap_i64(int64(base.Ui64(v136)>>(uint(int64(32))%64))) != v140>>(uint(int32(31))%32) {
		goto L5
	} else {
		goto L38
	}
L38:
	;
	v145 = v140 + int32(4)
	if v145 < v140 {
		goto L5
	} else {
		goto L39
	}
L39:
	;
	if base.Ui32(int32(1073741824)) <= base.Ui32(v145) {
		goto L5
	} else {
		goto L40
	}
L40:
	;
	v149 = int32(1)
	if v31&v149 != 0 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v153 = v149
	goto L43
L42:
	;
	v153 = int32(4)
	goto L43
L43:
	;
	v154 = v21 + v153
	v156 = int32(1)
	if v64&v156 != 0 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v160 = v156
	goto L46
L45:
	;
	v160 = int32(4)
	goto L46
L46:
	;
	v161 = v26 + v160
	v163 = int32(1)
	if v95&v163 != 0 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v167 = v163
	goto L49
L48:
	;
	v167 = int32(4)
	goto L49
L49:
	;
	v168 = v29 + v167
	v169 = v125 + v168
	v170 = base.B2i32(base.Ui32(v168) < base.Ui32(v169))
	v171 = F_palloc(m, v145)
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	v175 = int32(0)
	v184 = v175
	v186 = v63
	v188 = v154
	v189 = v171 + int32(4)
	goto L51
L51:
	;
	v197 = F_pg_mblen_range(m, v188, v154+v63)
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L1
	} else {
		goto L53
	}
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v171))) = v379<<(uint(int32(2))%32) + int32(16)
	v406 = v171
	goto L6
L53:
	;
	v199 = int32(0)
	if base.B2i32(v94 <= v175) == v199 {
		goto L55
	} else {
		goto L56
	}
L54:
	;
	v393 = v186 - v197
	if int32(0) < v393 {
		v184 = v379
		v186 = v393
		v188 = v197 + v188
		v189 = v384
		goto L51
	} else {
		goto L101
	}
L55:
	;
	v203 = v199
	v205 = v199
	goto L58
L56:
	;
	goto L57
L57:
	;
	if v197 != 0 {
		goto L98
	} else {
		goto L99
	}
L58:
	;
	v222 = v203 + v161
	v223 = F_pg_mblen_range(m, v222, v161+v94)
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L1
	} else {
		goto L61
	}
L59:
	;
	goto L57
L60:
	;
	v348 = v203 + v223
	if v348 < v94 {
		v203 = v348
		v205 = v205 + int32(1)
		goto L58
	} else {
		goto L96
	}
L61:
	;
	if v223 != v197 {
		goto L60
	} else {
		goto L62
	}
L62:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v197) {
		goto L66
	} else {
		goto L67
	}
L63:
	;
	if v287 != 0 {
		goto L60
	} else {
		goto L81
	}
L64:
	;
	v287 = int32(0)
	goto L63
L65:
	;
	v261 = v256
	v262 = v257
	v263 = v258
	goto L75
L66:
	;
	if (v188|v222)&int32(3) != 0 {
		v256 = v188
		v257 = v222
		v258 = v197
		goto L65
	} else {
		goto L69
	}
L67:
	;
	v249 = v188
	v250 = v222
	v251 = v197
	goto L68
L68:
	;
	if v251 == int32(0) {
		goto L64
	} else {
		goto L74
	}
L69:
	;
	v233 = v188
	v234 = v222
	v235 = v197
	goto L70
L70:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v233)))
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v234)))
	if v238 != v239 {
		v256 = v233
		v257 = v234
		v258 = v235
		goto L65
	} else {
		goto L72
	}
L71:
	;
	v249 = v244
	v250 = v242
	v251 = v246
	goto L68
L72:
	;
	v241 = int32(4)
	v242 = v234 + v241
	v244 = v233 + v241
	v246 = v235 - v241
	if base.Ui32(int32(3)) < base.Ui32(v246) {
		v233 = v244
		v234 = v242
		v235 = v246
		goto L70
	} else {
		goto L73
	}
L73:
	;
	goto L71
L74:
	;
	v256 = v249
	v257 = v250
	v258 = v251
	goto L65
L75:
	;
	v266 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v261))))
	v267 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v262))))
	if v266 == v267 {
		goto L77
	} else {
		goto L78
	}
L76:
	;
	v287 = v266 - v267
	goto L63
L77:
	;
	v269 = int32(1)
	v274 = v263 - v269
	if v274 != 0 {
		v261 = v261 + v269
		v262 = v262 + v269
		v263 = v274
		goto L75
	} else {
		goto L80
	}
L78:
	;
	goto L79
L79:
	;
	goto L76
L80:
	;
	goto L64
L81:
	;
	if v205 <= int32(0) {
		v319 = v168
		v326 = v170
		goto L82
	} else {
		goto L83
	}
L82:
	;
	if v326 == int32(0) {
		v379 = v184
		v384 = v189
		goto L54
	} else {
		goto L90
	}
L83:
	;
	if base.Ui32(v169) <= base.Ui32(v168) {
		v319 = v168
		v326 = v170
		goto L82
	} else {
		goto L84
	}
L84:
	;
	v292 = v168
	v296 = int32(0)
	goto L85
L85:
	;
	v311 = F_pg_mblen_range(m, v292, v169)
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L1
	} else {
		goto L87
	}
L86:
	;
	v319 = v313
	v326 = v314
	goto L82
L87:
	;
	v313 = v311 + v292
	v314 = base.B2i32(base.Ui32(v313) < base.Ui32(v169))
	v316 = v296 + int32(1)
	if v205 <= v316 {
		v319 = v313
		v326 = v314
		goto L82
	} else {
		goto L88
	}
L88:
	;
	if base.Ui32(v313) < base.Ui32(v169) {
		v292 = v313
		v296 = v316
		goto L85
	} else {
		goto L89
	}
L89:
	;
	goto L86
L90:
	;
	v340 = F_pg_mblen_range(m, v319, v169)
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		goto L1
	} else {
		goto L91
	}
L91:
	;
	if v340 != 0 {
		goto L93
	} else {
		goto L94
	}
L92:
	;
	v379 = v340 + v184
	v384 = v344 + v340
	goto L54
L93:
	;
	v343 = F__emscripten_memcpy_bulkmem(m, v189, v319, v340)
	mBase = m.M
	v344 = v343
	goto L95
L94:
	;
	v344 = v189
	goto L95
L95:
	;
	goto L92
L96:
	;
	goto L59
L97:
	;
	v379 = v197 + v184
	v384 = v371 + v197
	goto L54
L98:
	;
	v370 = F__emscripten_memcpy_bulkmem(m, v189, v188, v197)
	mBase = m.M
	v371 = v370
	goto L100
L99:
	;
	v371 = v189
	goto L100
L100:
	;
	goto L97
L101:
	;
	goto L52
L102:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v428 = m.ExcPending
	if v428 != 0 {
		goto L1
	} else {
		goto L103
	}
L103:
	;
	F_errmsg(m, int32(398222), int32(0))
	mBase = m.M
	v432 = m.ExcPending
	if v432 != 0 {
		goto L1
	} else {
		goto L104
	}
L104:
	;
	F_errfinish(m, int32(490348), int32(864), int32(352447))
	mBase = m.M
	v437 = m.ExcPending
	if v437 != 0 {
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
							F_errmsg(m, int32(671711), v7)
							mBase = m.M
							v30 = m.ExcPending
							if v30 != 0 {
								return
							} else {
								F_errfinish(m, int32(492511), int32(102), int32(221113))
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
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	v2 = m.G0
	v4 = v2 - int32(16)
	m.G0 = v4
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		F_errcode(m, int32(1088))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v4))) = int32(218366)
			F_errmsg(m, int32(191527), v4)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(490975), int32(372), int32(66642))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
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
	v7 = F_parse_tsquery(m, v2, int32(1526), v4, v4, v6)
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
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v71 int32
	_ = v71
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v175 int32
	_ = v175
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
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
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v274 int32
	_ = v274
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v300 int32
	_ = v300
	var v312 int32
	_ = v312
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v321 int32
	_ = v321
	var v324 int32
	_ = v324
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v368 int32
	_ = v368
	var v373 int32
	_ = v373
	var v376 int32
	_ = v376
	var v380 int32
	_ = v380
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v386 int32
	_ = v386
	var v388 int32
	_ = v388
	var v401 int32
	_ = v401
	var v405 int32
	_ = v405
	var v417 int32
	_ = v417
	var v419 int32
	_ = v419
	var v422 int32
	_ = v422
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
	v35 = int32(0)
	v37 = v30
	goto L6
L4:
	;
	v96 = v30
	goto L5
L5:
	;
	v107 = F_palloc(m, v96)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L1
	} else {
		goto L13
	}
L6:
	;
	v50 = v25 + v35<<(uint(int32(2))%32)
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
	v55 = *(*int32)(unsafe.Add(mBase, _consts[486]))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)+4))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v56*int32(28))+uint32(_consts[1295])))
	goto L8
L7:
	;
	v96 = v89
	goto L5
L8:
	;
	v63 = v51&int32(4094)*v61 + v37
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
	if v65&int32(1) != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v71 = int32(1)
	v83 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v25+v64<<(uint(int32(2))%32)+(int32(base.Ui32(v65)>>(uint(v71)%32))&int32(2047)+int32(base.Ui32(v65)>>(uint(int32(12))%32))+v71)&int32(4194302)))))
	v89 = v63 + v83*int32(7) + v71
	goto L11
L10:
	;
	v89 = v63
	goto L11
L11:
	;
	v91 = v35 + int32(1)
	if v91 < v64 {
		v35 = v91
		v37 = v89
		goto L6
	} else {
		goto L12
	}
L12:
	;
	goto L7
L13:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	if int32(0) < v109 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v114 = v107
	v116 = v109
	v119 = v25
	v121 = int32(0)
	goto L17
L15:
	;
	v405 = v107
	goto L16
L16:
	;
	v417 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v405))) = uint8(v417)
	v419 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v419 != v20 {
		goto L59
	} else {
		goto L60
	}
L17:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v119)))
	v132 = v25 + v116<<(uint(int32(2))%32) + int32(base.Ui32(v129)>>(uint(int32(12))%32))
	v137 = v132 + int32(base.Ui32(v129)>>(uint(int32(1))%32))&int32(2047)
	if v121 != 0 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v405 = v386
	goto L16
L19:
	;
	v138 = int32(32)
	*(*uint8)(unsafe.Add(mBase, uint32(v114))) = uint8(v138)
	v142 = v114 + int32(1)
	goto L21
L20:
	;
	v142 = v114
	goto L21
L21:
	;
	v143 = int32(39)
	*(*uint8)(unsafe.Add(mBase, uint32(v142))) = uint8(v143)
	v146 = v142 + int32(1)
	if base.Ui32(v132) < base.Ui32(v137) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v149 = v146
	v151 = v132
	goto L25
L23:
	;
	v274 = v146
	goto L24
L24:
	;
	v287 = int32(39)
	*(*uint8)(unsafe.Add(mBase, uint32(v274))) = uint8(v287)
	v289 = int32(1)
	v290 = v274 + v289
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v119)))
	if v292&v289 == int32(0) {
		v386 = v290
		v388 = v291
		goto L44
	} else {
		goto L45
	}
L25:
	;
	v162 = F_pg_mblen_range(m, v151, v137)
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L1
	} else {
		goto L27
	}
L26:
	;
	v274 = v259
	goto L24
L27:
	;
	v164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151))))
	if base.B2i32(v164 != int32(92))&base.B2i32(v164 != int32(39)) == int32(0) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v149))) = uint8(v164)
	v175 = v149 + int32(1)
	goto L30
L29:
	;
	v175 = v149
	goto L30
L30:
	;
	if v162 == int32(0) {
		v259 = v175
		v261 = v151
		goto L31
	} else {
		goto L32
	}
L31:
	;
	if base.Ui32(v261) < base.Ui32(v137) {
		v149 = v259
		v151 = v261
		goto L25
	} else {
		goto L43
	}
L32:
	;
	v180 = v162 & int32(7)
	if v180 != 0 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v182 = v175
	v183 = v162
	v184 = v151
	v185 = int32(0)
	goto L36
L34:
	;
	v207 = v175
	v208 = v162
	v209 = v151
	goto L35
L35:
	;
	if base.Ui32(v162) < base.Ui32(int32(8)) {
		v259 = v207
		v261 = v209
		goto L31
	} else {
		goto L39
	}
L36:
	;
	v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v184))))
	*(*uint8)(unsafe.Add(mBase, uint32(v182))) = uint8(v195)
	v197 = int32(1)
	v198 = v182 + v197
	v200 = v184 + v197
	v202 = v183 - v197
	v204 = v185 + v197
	if v204 != v180 {
		v182 = v198
		v183 = v202
		v184 = v200
		v185 = v204
		goto L36
	} else {
		goto L38
	}
L37:
	;
	v207 = v198
	v208 = v202
	v209 = v200
	goto L35
L38:
	;
	goto L37
L39:
	;
	v223 = v207
	v224 = v208
	v225 = v209
	goto L40
L40:
	;
	v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v225))))
	*(*uint8)(unsafe.Add(mBase, uint32(v223))) = uint8(v236)
	v238 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v225)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v223)+1)) = uint8(v238)
	v240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v225)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v223)+2)) = uint8(v240)
	v242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v225)+3)))
	*(*uint8)(unsafe.Add(mBase, uint32(v223)+3)) = uint8(v242)
	v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v225)+4)))
	*(*uint8)(unsafe.Add(mBase, uint32(v223)+4)) = uint8(v244)
	v246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v225)+5)))
	*(*uint8)(unsafe.Add(mBase, uint32(v223)+5)) = uint8(v246)
	v248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v225)+6)))
	*(*uint8)(unsafe.Add(mBase, uint32(v223)+6)) = uint8(v248)
	v250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v225)+7)))
	*(*uint8)(unsafe.Add(mBase, uint32(v223)+7)) = uint8(v250)
	v252 = int32(8)
	v253 = v223 + v252
	v255 = v225 + v252
	v257 = v224 - v252
	if v257 != 0 {
		v223 = v253
		v224 = v257
		v225 = v255
		goto L40
	} else {
		goto L42
	}
L41:
	;
	v259 = v253
	v261 = v255
	goto L31
L42:
	;
	goto L41
L43:
	;
	goto L26
L44:
	;
	v401 = v121 + int32(1)
	if v401 < v388 {
		v114 = v386
		v116 = v388
		v119 = v119 + int32(4)
		v121 = v401
		goto L17
	} else {
		goto L58
	}
L45:
	;
	v300 = int32(1)
	v312 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v25+v291<<(uint(int32(2))%32)+(int32(base.Ui32(v292)>>(uint(v300)%32))&int32(2047)+int32(base.Ui32(v292)>>(uint(int32(12))%32))+v300)&int32(4194302)))))
	if v312 == int32(0) {
		v386 = v290
		v388 = v291
		goto L44
	} else {
		goto L46
	}
L46:
	;
	v315 = int32(58)
	*(*uint8)(unsafe.Add(mBase, uint32(v274)+1)) = uint8(v315)
	v317 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	v318 = int32(2)
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v119)))
	v324 = int32(1)
	v338 = v274 + v318
	v339 = v312
	v340 = v25 + v317<<(uint(v318)%32) + (int32(base.Ui32(v321)>>(uint(int32(12))%32))+int32(base.Ui32(v321)>>(uint(v324)%32))&int32(2047)+v324)&int32(4194302)
	goto L47
L47:
	;
	v351 = v340 + int32(2)
	v352 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v351))))
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v352 & int32(16383)
	v357 = F_pg_sprintf(m, v338, int32(485141), v17)
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L1
	} else {
		goto L49
	}
L48:
	;
	v383 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	v386 = v380
	v388 = v383
	goto L44
L49:
	;
	v359 = v357 + v338
	v361 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v351))))
	switch int32(base.Ui32(v361)>>(uint(int32(14))%32)) - int32(1) {
	case 0:
		goto L52
	case 1:
		goto L53
	case 2:
		v368 = int32(65)
		goto L51
	default:
		v373 = v359
		goto L50
	}
L50:
	;
	if int32(2) <= v339 {
		goto L54
	} else {
		goto L55
	}
L51:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v359))) = uint8(v368)
	v373 = v359 + int32(1)
	goto L50
L52:
	;
	v368 = int32(67)
	goto L51
L53:
	;
	v368 = int32(66)
	goto L51
L54:
	;
	v376 = int32(44)
	*(*uint8)(unsafe.Add(mBase, uint32(v373))) = uint8(v376)
	v380 = v373 + int32(1)
	goto L56
L55:
	;
	v380 = v373
	goto L56
L56:
	;
	v382 = v339 - int32(1)
	if v382 != 0 {
		v338 = v380
		v339 = v382
		v340 = v351
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
	v422 = m.ExcPending
	if v422 != 0 {
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
	return v107
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
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v66 int32
	_ = v66
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v146 int32
	_ = v146
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v176 int32
	_ = v176
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
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
	v29 = int32(24)
	v31 = int32(65280)
	v33 = int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v26+v27))) = v22<<(uint(v29)%32) | v22&v31<<(uint(v33)%32) | (int32(base.Ui32(v22)>>(uint(v33)%32))&v31 | int32(base.Ui32(v22)>>(uint(v29)%32)))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v26 + int32(4)
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	if int32(0) < v48 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v52 = v16 + int32(8)
	v53 = v48
	v57 = v52
	v59 = int32(0)
	goto L8
L6:
	;
	goto L7
L7:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v205))) = v206 << (uint(int32(2)) % 32)
	goto L24
L8:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
	F_pq_sendtext(m, v13, v52+v53<<(uint(int32(2))%32)+int32(base.Ui32(v66)>>(uint(int32(12))%32)), int32(base.Ui32(v66)>>(uint(int32(1))%32))&int32(2047))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
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
	v78 = m.ExcPending
	if v78 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v79 = int32(0)
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	*(*uint8)(unsafe.Add(mBase, uint32(v80+v81))) = uint8(v79)
	v85 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v80 + v85
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
	if v88&v85 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	v95 = int32(1)
	v107 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v52+v91<<(uint(int32(2))%32)+(int32(base.Ui32(v88)>>(uint(v95)%32))&int32(2047)+int32(base.Ui32(v88)>>(uint(int32(12))%32))+v95)&int32(4194302)))))
	v108 = v107
	goto L14
L13:
	;
	v108 = v79
	goto L14
L14:
	;
	F_enlargeStringInfo(m, v13, int32(2))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	v115 = int32(8)
	v118 = v108 & int32(65535)
	v121 = v108<<(uint(v115)%32) | int32(base.Ui32(v118)>>(uint(v115)%32))
	*(*uint16)(unsafe.Add(mBase, uint32(v112+v113))) = uint16(v121)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v112 + int32(2)
	if v118 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	v127 = int32(2)
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
	v133 = int32(1)
	v146 = int32(0)
	goto L19
L17:
	;
	goto L18
L18:
	;
	v191 = v59 + int32(1)
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	if v191 < v192 {
		v53 = v192
		v57 = v57 + int32(4)
		v59 = v191
		goto L8
	} else {
		goto L23
	}
L19:
	;
	v159 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v52+v126<<(uint(v127)%32)+(int32(base.Ui32(v130)>>(uint(int32(12))%32))+int32(base.Ui32(v130)>>(uint(v133)%32))&int32(2047)+v133)&int32(4194302)+v127+v146<<(uint(int32(1))%32)))))
	F_enlargeStringInfo(m, v13, int32(2))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L1
	} else {
		goto L21
	}
L20:
	;
	goto L18
L21:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	v166 = int32(8)
	v170 = v159<<(uint(v166)%32) | int32(base.Ui32(v159)>>(uint(v166)%32))
	*(*uint16)(unsafe.Add(mBase, uint32(v163+v164))) = uint16(v170)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v163 + int32(2)
	v176 = v146 + int32(1)
	if v176 != v118 {
		v146 = v176
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
	return v205
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
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
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
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v108 int32
	_ = v108
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
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
		v108 = int32(1)
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
	v119 = m.ExcPending
	if v119 != 0 {
		goto L4
	} else {
		goto L28
	}
L11:
	;
	m.G0 = v13 + int32(16)
	return v108
L12:
	;
	v38 = int32(0)
	v39 = v30
	v40 = v31
	goto L13
L13:
	;
	v50 = v39 + v40<<(uint(int32(4))%32) + v38*int32(100)
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50)+111)))
	if v51 != 0 {
		v93 = v39
		v94 = v40
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v108 = v98
	goto L11
L15:
	;
	v98 = int32(1)
	v100 = v38 + v98
	if v100 < v94 {
		v38 = v100
		v39 = v93
		v40 = v94
		goto L13
	} else {
		goto L27
	}
L16:
	;
	v53 = v50 + int32(20)
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+90)))
	if v54 != 0 {
		v93 = v39
		v94 = v40
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v55 = int32(0)
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56+v38))))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59+v38))))
	if v58 != v61 {
		v108 = v55
		goto L11
	} else {
		goto L18
	}
L18:
	;
	if v58 != 0 {
		v93 = v39
		v94 = v40
		goto L15
	} else {
		goto L19
	}
L19:
	;
	v64 = v38 << (uint(int32(2)) % 32)
	v65 = l2 + v64
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
	if v66 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v53)+68))
	v71 = F_lookup_type_cache(m, v69, int32(32))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L4
	} else {
		goto L23
	}
L21:
	;
	v77 = v66
	goto L22
L22:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v53)+96))
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v81+v64)))
	v84 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v84+v64)))
	v87 = F_FunctionCall2Coll(m, v77+int32(76), v80, v83, v86)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L4
	} else {
		goto L25
	}
L23:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v71)+80))
	if v73 == int32(0) {
		goto L10
	} else {
		goto L24
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v65))) = v71
	v77 = v71
	goto L22
L25:
	;
	if v87 == int32(0) {
		v108 = v55
		goto L11
	} else {
		goto L26
	}
L26:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v93 = v91
	v94 = v92
	goto L15
L27:
	;
	goto L14
L28:
	;
	F_errcode(m, int32(52461700))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L4
	} else {
		goto L29
	}
L29:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v53)+68))
	v124 = F_format_type_be(m, v123)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L4
	} else {
		goto L30
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v124
	F_errmsg(m, int32(188458), v13)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L4
	} else {
		goto L31
	}
L31:
	;
	F_errfinish(m, int32(493175), int32(330), int32(308039))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
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
