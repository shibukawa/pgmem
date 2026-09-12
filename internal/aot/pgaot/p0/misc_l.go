package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_LargeObjectExistsWithSnapshot(m *base.Module, l0 int32, l1 int32) int32 {
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
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
	var v31 int32
	_ = v31
	v5 = m.G0
	v7 = v5 - int32(48)
	m.G0 = v7
	F_ScanKeyInit(m, v7, int32(1), int32(3), int32(184), l0)
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v18 = F_table_open(m, int32(2995), int32(1))
		v19 = m.ExcPending
		if v19 != 0 {
			return int32(0)
		} else {
			v21 = int32(1)
			v23 = F_systable_beginscan(m, v18, int32(2996), v21, l1, v21, v7)
			v24 = m.ExcPending
			if v24 != 0 {
				return int32(0)
			} else {
				v25 = F_systable_getnext(m, v23)
				v26 = m.ExcPending
				if v26 != 0 {
					return int32(0)
				} else {
					F_systable_endscan(m, v23)
					v28 = m.ExcPending
					if v28 != 0 {
						return int32(0)
					} else {
						F_sequence_close(m, v18, int32(1))
						v31 = m.ExcPending
						if v31 != 0 {
							return int32(0)
						} else {
							m.G0 = v7 + int32(48)
							return base.B2i32(v25 != int32(0))
						}
					}
				}
			}
		}
	}
}
func F_LexizeExec(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v27 int32
	_ = v27
	var v42 int32
	_ = v42
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
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
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
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v168 int32
	_ = v168
	var v173 int32
	_ = v173
	var v181 int32
	_ = v181
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v216 int32
	_ = v216
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v246 int32
	_ = v246
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v259 int32
	_ = v259
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v284 int32
	_ = v284
	var v288 int32
	_ = v288
	var v302 int32
	_ = v302
	var v312 int32
	_ = v312
	var v318 int32
	_ = v318
	var v331 int32
	_ = v331
	var v342 int32
	_ = v342
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v359 int32
	_ = v359
	var v365 int32
	_ = v365
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v385 int32
	_ = v385
	var v398 int32
	_ = v398
	var v400 int32
	_ = v400
	var v430 int32
	_ = v430
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v444 int32
	_ = v444
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v463 int32
	_ = v463
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v482 int32
	_ = v482
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v489 int32
	_ = v489
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v504 int32
	_ = v504
	var v507 int32
	_ = v507
	var v518 int32
	_ = v518
	var v520 int32
	_ = v520
	var v526 int32
	_ = v526
	var v530 int32
	_ = v530
	var v535 int32
	_ = v535
	var v540 int32
	_ = v540
	var v548 int32
	_ = v548
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v572 int32
	_ = v572
	var v582 int32
	_ = v582
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v588 int32
	_ = v588
	var v591 int32
	_ = v591
	var v602 int32
	_ = v602
	var v604 int32
	_ = v604
	var v610 int32
	_ = v610
	var v614 int32
	_ = v614
	var v619 int32
	_ = v619
	var v624 int32
	_ = v624
	var v632 int64
	_ = v632
	var v636 int32
	_ = v636
	var v641 int32
	_ = v641
	var v651 int32
	_ = v651
	var v653 int32
	_ = v653
	var v654 int32
	_ = v654
	var v656 int32
	_ = v656
	var v658 int32
	_ = v658
	var v667 int32
	_ = v667
	var v671 int32
	_ = v671
	v14 = l0 + int32(12)
	goto L5
L1:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v667)+32)) = int64(0)
	return v671
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v656
	v667 = v654
	v671 = v658
	goto L1
L3:
	;
	v632 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+40)) = v632
	*(*int64)(unsafe.Add(mBase, uint32(l0)+4)) = v632
	v636 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if l1 != 0 {
		v654 = l0
		v656 = v636
		v658 = v624
		goto L2
	} else {
		goto L168
	}
L4:
	;
	v585 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v585 == int32(0) {
		v624 = v439
		goto L3
	} else {
		goto L153
	}
L5:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v27 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	v566 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v567 = int32(0)
	if l1 != 0 {
		v654 = l0
		v656 = v566
		v658 = v567
		goto L2
	} else {
		goto L147
	}
L7:
	;
	goto L6
L8:
	;
	goto L11
L9:
	;
	goto L10
L10:
	;
	v346 = F_lookup_ts_dictionary_cache(m, v27)
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L32
	} else {
		goto L90
	}
L11:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v42 == int32(0) {
		goto L7
	} else {
		goto L13
	}
L13:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
	if v45 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v331
	v342 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v331)+12)) = v342
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v342
	goto L11
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v318
	v331 = v318
	goto L14
L16:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v55 <= v68 {
		v274 = v42
		goto L27
	} else {
		goto L28
	}
L17:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v42)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v58
	if v58 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L18:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v48)+12))
	if v49 <= v45 {
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
	v54 = v51 + v45<<(uint(int32(3))%32)
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
	if v55 != 0 {
		goto L16
	} else {
		goto L20
	}
L20:
	;
	goto L17
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = int32(0)
	goto L23
L22:
	;
	goto L23
L23:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v64 == int32(0) {
		v318 = v42
		goto L15
	} else {
		goto L24
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v64)+12)) = v42
	v331 = v42
	goto L14
L25:
	;
	v312 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v312 == int32(0) {
		v318 = v302
		goto L15
	} else {
		goto L89
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = int32(0)
	v302 = v288
	goto L25
L27:
	;
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v274)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v284
	if v284 != 0 {
		v302 = v274
		goto L25
	} else {
		goto L88
	}
L28:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v42)+8))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
	v75 = v68
	v78 = v71
	v79 = v70
	goto L30
L29:
	;
	v239 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v239 != 0 {
		goto L74
	} else {
		goto L75
	}
L30:
	;
	v85 = v75 << (uint(int32(2)) % 32)
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v54)+4))
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v85+v86)))
	v89 = F_lookup_ts_dictionary_cache(m, v88)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v237 != 0 {
		v274 = v237
		goto L27
	} else {
		goto L72
	}
L32:
	;
	return int32(0)
L33:
	;
	v93 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v93
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)) = uint16(v93)
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v89)+44))
	v101 = F_FunctionCall4Coll(m, v89+int32(12), v93, v100, v78, v79, v14)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L32
	} else {
		goto L34
	}
L34:
	;
	v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+13)))
	if v103 == int32(1) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v54)+4))
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v106+v85)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v75 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v108
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v42)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v113
	if v101 == int32(0) {
		goto L5
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	if v101 != 0 {
		goto L50
	} else {
		goto L51
	}
L38:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v117 != 0 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v117)+4))
	if v118 != 0 {
		goto L42
	} else {
		goto L43
	}
L40:
	;
	goto L41
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v101
	goto L5
L42:
	;
	v122 = v117
	v124 = v118
	goto L45
L43:
	;
	v151 = v117
	goto L44
L44:
	;
	F_pfree(m, v151)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L32
	} else {
		goto L49
	}
L45:
	;
	F_pfree(m, v124)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L32
	} else {
		goto L47
	}
L46:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v151 = v138
	goto L44
L47:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v122+int32(12))))
	if v137 != 0 {
		v122 = v122 + int32(8)
		v124 = v137
		goto L45
	} else {
		goto L48
	}
L48:
	;
	goto L46
L49:
	;
	goto L41
L50:
	;
	v168 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101)+2)))
	if v168&int32(4) == int32(0) {
		goto L29
	} else {
		goto L53
	}
L51:
	;
	v231 = v78
	v232 = v79
	goto L52
L52:
	;
	v234 = v75 + int32(1)
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
	if v234 < v235 {
		v75 = v234
		v78 = v231
		v79 = v232
		goto L30
	} else {
		goto L71
	}
L53:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v101)+4))
	if v173&int32(3) == int32(0) {
		v197 = v173
		goto L56
	} else {
		goto L57
	}
L54:
	;
	v231 = v173
	v232 = v230
	goto L52
L55:
	;
	v230 = v222 - v173
	goto L54
L56:
	;
	v201 = v197
	goto L65
L57:
	;
	v181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173))))
	if v181 == int32(0) {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v230 = int32(0)
	goto L54
L59:
	;
	goto L60
L60:
	;
	v186 = v173
	goto L61
L61:
	;
	v190 = v186 + int32(1)
	if v190&int32(3) == int32(0) {
		v197 = v190
		goto L56
	} else {
		goto L63
	}
L62:
	;
	v222 = v190
	goto L55
L63:
	;
	v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v190))))
	if v195 != 0 {
		v186 = v190
		goto L61
	} else {
		goto L64
	}
L64:
	;
	goto L62
L65:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v201)))
	v210 = int32(-2139062144)
	if (int32(16843008)-v207|v207)&v210 == v210 {
		v201 = v201 + int32(4)
		goto L65
	} else {
		goto L67
	}
L66:
	;
	v216 = v201
	goto L68
L67:
	;
	goto L66
L68:
	;
	v220 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v216))))
	if v220 != 0 {
		v216 = v216 + int32(1)
		goto L68
	} else {
		goto L70
	}
L69:
	;
	v222 = v216
	goto L55
L70:
	;
	goto L69
L71:
	;
	goto L31
L72:
	;
	v288 = int32(0)
	goto L26
L73:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v246 != 0 {
		goto L79
	} else {
		goto L80
	}
L74:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v239)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v240
	if v240 != 0 {
		goto L73
	} else {
		goto L77
	}
L75:
	;
	goto L76
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = int32(0)
	goto L73
L77:
	;
	goto L76
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v239
	v250 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v239)+12)) = v250
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v250
	v254 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if l1 != 0 {
		v654 = l0
		v656 = v254
		v658 = v101
		goto L2
	} else {
		goto L82
	}
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v246)+12)) = v239
	goto L78
L80:
	;
	goto L81
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v239
	goto L78
L82:
	;
	if v254 == int32(0) {
		v667 = l0
		v671 = v101
		goto L1
	} else {
		goto L83
	}
L83:
	;
	v259 = v254
	goto L84
L84:
	;
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v259)+12))
	F_pfree(m, v259)
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L32
	} else {
		goto L86
	}
L85:
	;
	v667 = l0
	v671 = v101
	goto L1
L86:
	;
	if v269 != 0 {
		v259 = v269
		goto L84
	} else {
		goto L87
	}
L87:
	;
	goto L85
L88:
	;
	v288 = v274
	goto L26
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v312)+12)) = v302
	v331 = v302
	goto L14
L90:
	;
	v348 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v348 == int32(0) {
		goto L7
	} else {
		goto L91
	}
L91:
	;
	v359 = v348
	goto L92
L92:
	;
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v359)))
	if v365 == int32(0) {
		goto L96
	} else {
		goto L97
	}
L93:
	;
	goto L7
L94:
	;
	if v548 != 0 {
		v359 = v548
		goto L92
	} else {
		goto L146
	}
L95:
	;
	v540 = *(*int32)(unsafe.Add(mBase, uint32(v359)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v540
	v548 = v540
	goto L94
L96:
	;
	v430 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+13)) = uint8(v430)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)) = uint8(base.B2i32(v365 == v430))
	v436 = *(*int32)(unsafe.Add(mBase, uint32(v346)+44))
	v437 = *(*int32)(unsafe.Add(mBase, uint32(v359)+4))
	v438 = *(*int32)(unsafe.Add(mBase, uint32(v359)+8))
	v439 = F_FunctionCall4Coll(m, v346+int32(12), v430, v436, v437, v438, v14)
	mBase = m.M
	v440 = m.ExcPending
	if v440 != 0 {
		goto L32
	} else {
		goto L109
	}
L97:
	;
	v368 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v368)+12))
	if v369 <= v365 {
		goto L95
	} else {
		goto L98
	}
L98:
	;
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v368)+16))
	v374 = v371 + v365<<(uint(int32(3))%32)
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v374)))
	if v375 == int32(0) {
		goto L95
	} else {
		goto L99
	}
L99:
	;
	if int32(0) < v375 {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	v380 = *(*int32)(unsafe.Add(mBase, uint32(v374)+4))
	v381 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v385 = int32(0)
	goto L103
L101:
	;
	goto L102
L102:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(0)
	goto L5
L103:
	;
	v398 = *(*int32)(unsafe.Add(mBase, uint32(v380+v385<<(uint(int32(2))%32))))
	v400 = v385 + int32(1)
	if v375 <= v400 {
		goto L105
	} else {
		goto L106
	}
L104:
	;
	if v398 == v381 {
		goto L96
	} else {
		goto L108
	}
L105:
	;
	goto L104
L106:
	;
	if v398 != v381 {
		v385 = v400
		goto L103
	} else {
		goto L107
	}
L107:
	;
	goto L105
L108:
	;
	goto L102
L109:
	;
	v441 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+13)))
	if v441 == int32(1) {
		goto L110
	} else {
		goto L111
	}
L110:
	;
	v444 = *(*int32)(unsafe.Add(mBase, uint32(v359)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v444
	if v439 == int32(0) {
		goto L113
	} else {
		goto L114
	}
L111:
	;
	goto L112
L112:
	;
	if v439 != 0 {
		goto L4
	} else {
		goto L127
	}
L113:
	;
	v548 = v444
	goto L94
L114:
	;
	goto L115
L115:
	;
	v448 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v448 != 0 {
		goto L116
	} else {
		goto L117
	}
L116:
	;
	v449 = *(*int32)(unsafe.Add(mBase, uint32(v448)+4))
	if v449 != 0 {
		goto L119
	} else {
		goto L120
	}
L117:
	;
	v489 = v444
	goto L118
L118:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v359
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v439
	v548 = v489
	goto L94
L119:
	;
	v452 = v448
	v453 = v449
	goto L122
L120:
	;
	v482 = v448
	goto L121
L121:
	;
	F_pfree(m, v482)
	mBase = m.M
	v484 = m.ExcPending
	if v484 != 0 {
		goto L32
	} else {
		goto L126
	}
L122:
	;
	F_pfree(m, v453)
	mBase = m.M
	v463 = m.ExcPending
	if v463 != 0 {
		goto L32
	} else {
		goto L124
	}
L123:
	;
	v469 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v482 = v469
	goto L121
L124:
	;
	v468 = *(*int32)(unsafe.Add(mBase, uint32(v452+int32(12))))
	if v468 != 0 {
		v452 = v452 + int32(8)
		v453 = v468
		goto L122
	} else {
		goto L125
	}
L125:
	;
	goto L123
L126:
	;
	v485 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v489 = v485
	goto L118
L127:
	;
	v500 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v500 != 0 {
		goto L128
	} else {
		goto L129
	}
L128:
	;
	v501 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v501 == int32(0) {
		v624 = v500
		goto L3
	} else {
		goto L131
	}
L129:
	;
	goto L130
L130:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(0)
	goto L5
L131:
	;
	v504 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v507 = v501
	goto L132
L132:
	;
	if v507 == v504 {
		goto L134
	} else {
		goto L135
	}
L133:
	;
	v624 = v500
	goto L3
L134:
	;
	v518 = *(*int32)(unsafe.Add(mBase, uint32(v504)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v518
	goto L136
L135:
	;
	goto L136
L136:
	;
	v520 = *(*int32)(unsafe.Add(mBase, uint32(v507)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v520
	if v520 == int32(0) {
		goto L137
	} else {
		goto L138
	}
L137:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = int32(0)
	goto L139
L138:
	;
	goto L139
L139:
	;
	v526 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v526 != 0 {
		goto L141
	} else {
		goto L142
	}
L140:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v507
	v530 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v507)+12)) = v530
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v530
	if v507 == v504 {
		v624 = v500
		goto L3
	} else {
		goto L144
	}
L141:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v526)+12)) = v507
	goto L140
L142:
	;
	goto L143
L143:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v507
	goto L140
L144:
	;
	v535 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v535 != 0 {
		v507 = v535
		goto L132
	} else {
		goto L145
	}
L145:
	;
	goto L133
L146:
	;
	goto L93
L147:
	;
	if v566 == int32(0) {
		v667 = l0
		v671 = v567
		goto L1
	} else {
		goto L148
	}
L148:
	;
	v572 = v566
	goto L149
L149:
	;
	v582 = *(*int32)(unsafe.Add(mBase, uint32(v572)+12))
	F_pfree(m, v572)
	mBase = m.M
	v584 = m.ExcPending
	if v584 != 0 {
		goto L32
	} else {
		goto L151
	}
L150:
	;
	v667 = l0
	v671 = v567
	goto L1
L151:
	;
	if v582 != 0 {
		v572 = v582
		goto L149
	} else {
		goto L152
	}
L152:
	;
	goto L150
L153:
	;
	v588 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v591 = v585
	goto L154
L154:
	;
	if v591 == v588 {
		goto L156
	} else {
		goto L157
	}
L155:
	;
	v624 = v439
	goto L3
L156:
	;
	v602 = *(*int32)(unsafe.Add(mBase, uint32(v588)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v602
	goto L158
L157:
	;
	goto L158
L158:
	;
	v604 = *(*int32)(unsafe.Add(mBase, uint32(v591)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v604
	if v604 == int32(0) {
		goto L159
	} else {
		goto L160
	}
L159:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = int32(0)
	goto L161
L160:
	;
	goto L161
L161:
	;
	v610 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v610 != 0 {
		goto L163
	} else {
		goto L164
	}
L162:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v591
	v614 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v591)+12)) = v614
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v614
	if v591 == v588 {
		v624 = v439
		goto L3
	} else {
		goto L166
	}
L163:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v610)+12)) = v591
	goto L162
L164:
	;
	goto L165
L165:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v591
	goto L162
L166:
	;
	v619 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v619 != 0 {
		v591 = v619
		goto L154
	} else {
		goto L167
	}
L167:
	;
	goto L155
L168:
	;
	if v636 == int32(0) {
		v667 = l0
		v671 = v624
		goto L1
	} else {
		goto L169
	}
L169:
	;
	v641 = v636
	goto L170
L170:
	;
	v651 = *(*int32)(unsafe.Add(mBase, uint32(v641)+12))
	F_pfree(m, v641)
	mBase = m.M
	v653 = m.ExcPending
	if v653 != 0 {
		goto L32
	} else {
		goto L172
	}
L171:
	;
	v667 = l0
	v671 = v624
	goto L1
L172:
	;
	if v651 != 0 {
		v641 = v651
		goto L170
	} else {
		goto L173
	}
L173:
	;
	goto L171
}
func F_LockCheckConflicts(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v53 int32
	_ = v53
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v167 int32
	_ = v167
	var v186 int32
	_ = v186
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v212 int32
	_ = v212
	v5 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(48)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v17+l1<<(uint(int32(2))%32))))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	if v21&v22 == v5 {
		v186 = v5
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L41
	} else {
		goto L42
	}
L2:
	;
	m.G0 = v15 + int32(48)
	return v186
L3:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v26 <= int32(0) {
		v186 = v5
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	v30 = int32(2)
	v32 = v26 + int32(1)
	if v32 <= v30 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v35 = v30
	goto L7
L6:
	;
	v35 = v32
	goto L7
L7:
	;
	v40 = int32(1)
	v41 = int32(0)
	goto L8
L8:
	;
	v53 = int32(1) << (uint(v40) % 32)
	if v53&v21 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	if v67 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15+v40<<(uint(int32(2))%32)))) = v68
	v74 = v40 + int32(1)
	if v74 != v35 {
		v40 = v74
		v41 = v67
		goto L8
	} else {
		goto L14
	}
L11:
	;
	v67 = v41
	v68 = int32(0)
	goto L10
L12:
	;
	goto L13
L13:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l2+int32(88)+v40<<(uint(int32(2))%32))))
	v65 = v61 - base.B2i32(v53&v29 != int32(0))
	v67 = v65 + v41
	v68 = v65
	goto L10
L14:
	;
	goto L9
L15:
	;
	v186 = int32(0)
	goto L2
L16:
	;
	goto L17
L17:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v81 = *(*int32)(unsafe.Add(mBase, _c_F_LockCheckConflicts[0]))
	if v79 == v81 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v186 = int32(1)
	goto L2
L19:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v81)+616))
	if v83 == int32(0) {
		goto L18
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v86 = int32(1)
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+14)))
	if v87 == v86 {
		v186 = v86
		goto L2
	} else {
		goto L23
	}
L22:
	;
	goto L21
L23:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(l2)+28))
	if v90 == int32(0) {
		v186 = v86
		goto L2
	} else {
		goto L24
	}
L24:
	;
	v94 = l2 + int32(24)
	if v90 == v94 {
		v186 = v86
		goto L2
	} else {
		goto L25
	}
L25:
	;
	v97 = v67
	v101 = v90
	goto L26
L26:
	;
	if l3 == v101-int32(20) {
		v156 = v97
		goto L28
	} else {
		goto L29
	}
L27:
	;
	goto L18
L28:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v101)+4))
	if v167 != v94 {
		v97 = v156
		v101 = v167
		goto L26
	} else {
		goto L40
	}
L29:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v101-int32(12))))
	if v79 != v113 {
		v156 = v97
		goto L28
	} else {
		goto L30
	}
L30:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v101-int32(8))))
	v119 = v118 & v21
	if v119 == int32(0) {
		v156 = v97
		goto L28
	} else {
		goto L31
	}
L31:
	;
	v122 = int32(1)
	v123 = v97
	goto L32
L32:
	;
	if int32(base.Ui32(v119)>>(uint(v122)%32))&int32(1) != 0 {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	if v148 != 0 {
		v156 = v148
		goto L28
	} else {
		goto L39
	}
L34:
	;
	v139 = v15 + v122<<(uint(int32(2))%32)
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v139)))
	if v140 <= int32(0) {
		goto L1
	} else {
		goto L37
	}
L35:
	;
	v148 = v123
	goto L36
L36:
	;
	v152 = v122 + int32(1)
	if v152 <= v26 {
		v122 = v152
		v123 = v148
		goto L32
	} else {
		goto L38
	}
L37:
	;
	v143 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v139))) = v140 - v143
	v148 = v123 - v143
	goto L36
L38:
	;
	goto L33
L39:
	;
	v186 = int32(0)
	goto L2
L40:
	;
	goto L27
L41:
	;
	return int32(0)
L42:
	;
	F_errmsg_internal(m, int32(_a_F_LockCheckConflicts_0), int32(0))
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L41
	} else {
		goto L43
	}
L43:
	;
	F_errfinish(m, int32(_a_F_LockCheckConflicts_1), int32(1625), int32(_a_F_LockCheckConflicts_2))
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L41
	} else {
		goto L44
	}
L44:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_LockViewRecurse_walker(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v105 int64
	_ = v105
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v183 int32
	_ = v183
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v203 int32
	_ = v203
	v3 = int32(0)
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	if l0 == v3 {
		v183 = v3
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L12
	} else {
		goto L62
	}
L2:
	;
	m.G0 = v13 + int32(16)
	return v183
L3:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v18 == int32(67) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	if v21 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	v171 = F_expression_tree_walker_impl(m, l0, int32(560), l1)
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L12
	} else {
		goto L61
	}
L7:
	;
	v168 = F_query_tree_walker_impl(m, l0, int32(560), l1, int32(4))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L12
	} else {
		goto L60
	}
L8:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	if v24 <= int32(0) {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v34 = v3
	goto L10
L10:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v37+v34<<(uint(int32(2))%32))))
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+21)))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v41)+16))
	v44 = F_get_rel_name(m, v43)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	goto L7
L12:
	;
	return int32(0)
L13:
	;
	v49 = v42 - int32(112)
	if base.Ui32(int32(6)) < base.Ui32(v49) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v153 = v34 + int32(1)
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	if v153 < v154 {
		v34 = v153
		goto L10
	} else {
		goto L59
	}
L15:
	;
	if int32(1)<<(uint(v49)%32)&int32(69) == int32(0) {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v59 = int32(0)
	if v58 == v59 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	if v97 != 0 {
		goto L14
	} else {
		goto L30
	}
L18:
	;
	v97 = int32(0)
	goto L17
L19:
	;
	goto L20
L20:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v58)+4))
	if v65 <= int32(0) {
		v90 = v59
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v97 = v90
	goto L17
L22:
	;
	v68 = int32(0)
	if v68 < v65 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v71 = v65
	goto L25
L24:
	;
	v71 = v68
	goto L25
L25:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v58)+12))
	v74 = int32(0)
	goto L26
L26:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v72+v74<<(uint(int32(2))%32))))
	v83 = base.B2i32(v82 == v43)
	if v82 == v43 {
		v90 = v83
		goto L21
	} else {
		goto L28
	}
L27:
	;
	v90 = v83
	goto L21
L28:
	;
	v85 = v74 + int32(1)
	if v85 != v71 {
		v74 = v85
		goto L26
	} else {
		goto L29
	}
L29:
	;
	goto L27
L30:
	;
	v98 = base.I32_extend8_s(v42)
	v99 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v102 < int32(2) {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v105 = int64(16414)
	goto L33
L32:
	;
	v105 = int64(16412)
	goto L33
L33:
	;
	v110 = F_pg_class_aclcheck(m, v43, v99, v105|base.I64_extend_i32_u(base.B2i32(v102 < int32(4))))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L12
	} else {
		goto L34
	}
L34:
	;
	if v110 != 0 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	switch v98 - int32(73) {
	case 0, 32:
		goto L44
	default:
		v121 = int32(41)
		goto L39
	case 10:
		goto L43
	case 29:
		goto L40
	case 36:
		goto L41
	case 45:
		goto L42
	}
L36:
	;
	goto L37
L37:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+4)))
	if v127 == int32(0) {
		goto L47
	} else {
		goto L48
	}
L38:
	;
	F_aclcheck_error(m, v110, v123, v44)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L12
	} else {
		goto L45
	}
L39:
	;
	v123 = v121
	goto L38
L40:
	;
	v121 = int32(18)
	goto L39
L41:
	;
	v123 = int32(23)
	goto L38
L42:
	;
	v123 = int32(51)
	goto L38
L43:
	;
	v123 = int32(37)
	goto L38
L44:
	;
	v123 = int32(20)
	goto L38
L45:
	;
	goto L37
L46:
	;
	if v98 == int32(118) {
		goto L53
	} else {
		goto L54
	}
L47:
	;
	F_LockRelationOid(m, v43, v126)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L12
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	v132 = F_ConditionalLockRelationOid(m, v43, v126)
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L12
	} else {
		goto L51
	}
L50:
	;
	goto L46
L51:
	;
	if v132 == int32(0) {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	goto L46
L53:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+4)))
	v140 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	F_LockViewRecurse(m, v43, v138, v139, v140)
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L12
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+20)))
	if v143 != int32(1) {
		goto L14
	} else {
		goto L57
	}
L56:
	;
	goto L14
L57:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+4)))
	F_LockTableRecurse(m, v43, v146, v147)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L12
	} else {
		goto L58
	}
L58:
	;
	goto L14
L59:
	;
	goto L11
L60:
	;
	v183 = v168
	goto L2
L61:
	;
	v183 = v171
	goto L2
L62:
	;
	F_errcode(m, int32(50463045))
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L12
	} else {
		goto L63
	}
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v44
	F_errmsg(m, int32(_a_F_LockViewRecurse_walker_0), v13)
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L12
	} else {
		goto L64
	}
L64:
	;
	F_errfinish(m, int32(_a_F_LockViewRecurse_walker_1), int32(224), int32(_a_F_LockViewRecurse_walker_2))
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L12
	} else {
		goto L65
	}
L65:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_LookupCreationNamespace(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
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
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = int32(_a_F_LookupCreationNamespace_0)
	v12 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_LookupCreationNamespace[0])))
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v13 == int32(0) {
		v32 = v12
		v33 = v13
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L14
	} else {
		goto L21
	}
L2:
	;
	m.G0 = v7 + int32(16)
	return v63
L3:
	;
	if v33-v32 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L4:
	;
	goto L3
L5:
	;
	if v12 != v13 {
		v32 = v12
		v33 = v13
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v17 = l0
	v18 = v9
	goto L7
L7:
	;
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+1)))
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+1)))
	if v22 == int32(0) {
		v32 = v21
		v33 = v22
		goto L4
	} else {
		goto L9
	}
L8:
	;
	v32 = v21
	v33 = v22
	goto L4
L9:
	;
	v25 = int32(1)
	if v21 == v22 {
		v17 = v17 + v25
		v18 = v18 + v25
		goto L7
	} else {
		goto L10
	}
L10:
	;
	goto L8
L11:
	;
	F_AccessTempTableNamespace(m, int32(0))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	goto L13
L13:
	;
	v45 = int32(0)
	v48 = F_GetSysCacheOid(m, int32(37), l0, v45, v45, v45)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L14
	} else {
		goto L16
	}
L14:
	;
	return int32(0)
L15:
	;
	v43 = *(*int32)(unsafe.Add(mBase, _c_F_LookupCreationNamespace[1]))
	v63 = v43
	goto L2
L16:
	;
	if v48 == int32(0) {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	v54 = *(*int32)(unsafe.Add(mBase, _c_F_LookupCreationNamespace[2]))
	v56 = F_object_aclcheck(m, int32(2615), v48, v54, int64(512))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L14
	} else {
		goto L18
	}
L18:
	;
	if v56 == int32(0) {
		v63 = v48
		goto L2
	} else {
		goto L19
	}
L19:
	;
	F_aclcheck_error(m, v56, int32(36), l0)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L14
	} else {
		goto L20
	}
L20:
	;
	v63 = v48
	goto L2
L21:
	;
	F_errcode(m, int32(1411))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L14
	} else {
		goto L22
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
	F_errmsg(m, int32(_a_F_LookupCreationNamespace_1), v7)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L14
	} else {
		goto L23
	}
L23:
	;
	F_errfinish(m, int32(_a_F_LookupCreationNamespace_2), int32(3547), int32(_a_F_LookupCreationNamespace_3))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L14
	} else {
		goto L24
	}
L24:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F__ltq_regex(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
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
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = F_pg_detoast_datum(m, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v12 = F_pg_detoast_datum(m, v11)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int32(0)
		} else {
			v15 = F_array_iterator(m, v6, int32(_a_F__ltq_regex_0), v12, int32(0))
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int32(0)
			} else {
				v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				if v17 != v6 {
					F_pfree(m, v6)
					mBase = m.M
					v20 = m.ExcPending
					if v20 != 0 {
						return int32(0)
					} else {
						v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
						if v21 != v12 {
							F_pfree(m, v12)
							mBase = m.M
							v24 = m.ExcPending
							if v24 != 0 {
								return int32(0)
							} else {
								return v15
							}
						} else {
							return v15
						}
					}
				} else {
					v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					if v21 != v12 {
						F_pfree(m, v12)
						mBase = m.M
						v24 = m.ExcPending
						if v24 != 0 {
							return int32(0)
						} else {
							return v15
						}
					} else {
						return v15
					}
				}
			}
		}
	}
}
func F_lappend(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
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
	if l0 == int32(0) {
		v9 = F_palloc(m, int32(32))
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = int32(4)
			*(*int64)(unsafe.Add(mBase, uint32(v9))) = int64(4294967297)
			v18 = v9 + int32(16)
			*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v18
			v71 = v9
			v72 = v18
			v73 = int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v72+v73<<(uint(int32(2))%32)-int32(4)))) = l1
			return v71
		}
	} else {
		v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		if v22 <= v21 {
			v24 = int32(1)
			v26 = int32(16)
			v28 = v21 + v24
			if v28 <= v26 {
				v31 = v26
			} else {
				v31 = v28
			}
			if v31&(v31-int32(1)) != 0 {
				v38 = v24 << (uint(int32(32)-base.I32_clz(v31)) % 32)
			} else {
				v38 = v31
			}
			v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v41 = l0 + int32(16)
			if v39 == v41 {
				v43 = F_GetMemoryChunkContext(m, l0)
				mBase = m.M
				v44 = m.ExcPending
				if v44 != 0 {
					return int32(0)
				} else {
					v47 = F_MemoryContextAlloc(m, v43, v38<<(uint(int32(2))%32))
					mBase = m.M
					v48 = m.ExcPending
					if v48 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v47
						v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v52 = v50 << (uint(int32(2)) % 32)
						if v52 != 0 {
							v53 = F__emscripten_memcpy_bulkmem(m, v47, v41, v52)
							mBase = m.M
						} else {
						}
						*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v38
						v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v66 = v62
						v68 = v66 + int32(1)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v68
						v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						v71 = l0
						v72 = v70
						v73 = v68
						*(*int32)(unsafe.Add(mBase, uint32(v72+v73<<(uint(int32(2))%32)-int32(4)))) = l1
						return v71
					}
				}
			} else {
				v57 = F_repalloc(m, v39, v38<<(uint(int32(2))%32))
				mBase = m.M
				v58 = m.ExcPending
				if v58 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v57
					*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v38
					v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v66 = v62
					v68 = v66 + int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v68
					v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					v71 = l0
					v72 = v70
					v73 = v68
					*(*int32)(unsafe.Add(mBase, uint32(v72+v73<<(uint(int32(2))%32)-int32(4)))) = l1
					return v71
				}
			}
		} else {
			v66 = v21
			v68 = v66 + int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v68
			v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v71 = l0
			v72 = v70
			v73 = v68
			*(*int32)(unsafe.Add(mBase, uint32(v72+v73<<(uint(int32(2))%32)-int32(4)))) = l1
			return v71
		}
	}
}
func F_lappend_xid(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
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
	if l0 == int32(0) {
		v9 = F_palloc(m, int32(32))
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = int32(4)
			*(*int64)(unsafe.Add(mBase, uint32(v9))) = int64(4294967769)
			v18 = v9 + int32(16)
			*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v18
			v71 = v9
			v72 = v18
			v73 = int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v72+v73<<(uint(int32(2))%32)-int32(4)))) = l1
			return v71
		}
	} else {
		v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		if v22 <= v21 {
			v24 = int32(1)
			v26 = int32(16)
			v28 = v21 + v24
			if v28 <= v26 {
				v31 = v26
			} else {
				v31 = v28
			}
			if v31&(v31-int32(1)) != 0 {
				v38 = v24 << (uint(int32(32)-base.I32_clz(v31)) % 32)
			} else {
				v38 = v31
			}
			v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v41 = l0 + int32(16)
			if v39 == v41 {
				v43 = F_GetMemoryChunkContext(m, l0)
				mBase = m.M
				v44 = m.ExcPending
				if v44 != 0 {
					return int32(0)
				} else {
					v47 = F_MemoryContextAlloc(m, v43, v38<<(uint(int32(2))%32))
					mBase = m.M
					v48 = m.ExcPending
					if v48 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v47
						v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v52 = v50 << (uint(int32(2)) % 32)
						if v52 != 0 {
							v53 = F__emscripten_memcpy_bulkmem(m, v47, v41, v52)
							mBase = m.M
						} else {
						}
						*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v38
						v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v66 = v62
						v68 = v66 + int32(1)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v68
						v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						v71 = l0
						v72 = v70
						v73 = v68
						*(*int32)(unsafe.Add(mBase, uint32(v72+v73<<(uint(int32(2))%32)-int32(4)))) = l1
						return v71
					}
				}
			} else {
				v57 = F_repalloc(m, v39, v38<<(uint(int32(2))%32))
				mBase = m.M
				v58 = m.ExcPending
				if v58 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v57
					*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v38
					v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v66 = v62
					v68 = v66 + int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v68
					v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					v71 = l0
					v72 = v70
					v73 = v68
					*(*int32)(unsafe.Add(mBase, uint32(v72+v73<<(uint(int32(2))%32)-int32(4)))) = l1
					return v71
				}
			}
		} else {
			v66 = v21
			v68 = v66 + int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v68
			v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v71 = l0
			v72 = v70
			v73 = v68
			*(*int32)(unsafe.Add(mBase, uint32(v72+v73<<(uint(int32(2))%32)-int32(4)))) = l1
			return v71
		}
	}
}
func F_last_dir_separator(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	v4 = l0
	v5 = int32(0)
	for {
		v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4))))
		if v7 != int32(47) {
		} else {
			v11 = v4
			v4 = v4 + int32(1)
			v5 = v11
			continue
		}
		if v7 != 0 {
			v11 = v5
			v4 = v4 + int32(1)
			v5 = v11
			continue
		} else {
			break
		}
		break
	}
	return v5
}
func F_latin2_to_mic(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
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
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_check_encoding_conversion_args(m, v8, v9, v10, int32(9), int32(7))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		v21 = F_latin2mic(m, v6, v5, v10, int32(130), int32(9), base.B2i32(v7 != int32(0)))
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return int32(0)
		} else {
			return v21
		}
	}
}
func F_length_in_encoding(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v48 int32
	_ = v48
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v104 int32
	_ = v104
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v136 int32
	_ = v136
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v272 int32
	_ = v272
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v313 int32
	_ = v313
	var v324 int32
	_ = v324
	var v327 int32
	_ = v327
	var v331 int32
	_ = v331
	var v336 int32
	_ = v336
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = F_pg_detoast_datum_packed(m, v12)
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
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v25 = m.G0
	v27 = v25 + int32(-64)
	m.G0 = v27
	v29 = int32(-1)
	if v17 == int32(0) {
		v104 = v29
		goto L4
	} else {
		goto L5
	}
L3:
	;
	if int32(0) <= v104 {
		goto L29
	} else {
		goto L30
	}
L4:
	;
	m.G0 = v27 - int32(-64)
	goto L3
L5:
	;
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
	if v32 == int32(0) {
		v104 = v29
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v35 = F_strlen(m, v17)
	mBase = m.M
	if base.Ui32(int32(63)) < base.Ui32(v35) {
		v104 = v29
		goto L4
	} else {
		goto L7
	}
L7:
	;
	v38 = v17
	v39 = v32
	v40 = v27
	goto L8
L8:
	;
	v48 = F_isalnum(m, v39&int32(255))
	mBase = m.M
	if v48 != 0 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v65 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v61))) = uint8(v65)
	v69 = int32(*(*int8)(unsafe.Add(mBase, uint32(v27))))
	v70 = int32(_a_F_length_in_encoding_0)
	v71 = int32(_a_F_length_in_encoding_1)
	goto L17
L10:
	;
	if base.Ui32((v39-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	v61 = v40
	goto L12
L12:
	;
	v63 = v38 + int32(1)
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63))))
	if v64 != 0 {
		v38 = v63
		v39 = v64
		v40 = v61
		goto L8
	} else {
		goto L16
	}
L13:
	;
	v57 = v39 | int32(32)
	goto L15
L14:
	;
	v57 = v39
	goto L15
L15:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v40))) = uint8(v57)
	v61 = v40 + int32(1)
	goto L12
L16:
	;
	goto L9
L17:
	;
	v83 = v71 + (v70-v71)>>(uint(int32(4))%32)<<(uint(int32(3))%32)
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
	v85 = int32(*(*int8)(unsafe.Add(mBase, uint32(v84))))
	v86 = v69 - v85
	if v86 != 0 {
		v89 = v86
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v104 = v29
	goto L4
L19:
	;
	v93 = base.B2i32(v89 < int32(0))
	if v89 < int32(0) {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	v87 = F_strcmp(m, v27, v84)
	mBase = m.M
	if v87 != 0 {
		v89 = v87
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v83)+4))
	v104 = v88
	goto L4
L22:
	;
	v94 = v83 - int32(8)
	goto L24
L23:
	;
	v94 = v70
	goto L24
L24:
	;
	if v89 < int32(0) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v97 = v71
	goto L27
L26:
	;
	v97 = v83 + int32(8)
	goto L27
L27:
	;
	if base.Ui32(v97) <= base.Ui32(v94) {
		v70 = v94
		v71 = v97
		goto L17
	} else {
		goto L28
	}
L28:
	;
	goto L18
L29:
	;
	v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
	if v112 == int32(1) {
		goto L33
	} else {
		goto L34
	}
L30:
	;
	goto L31
L31:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L1
	} else {
		goto L101
	}
L32:
	;
	v143 = int32(1)
	if v112&v143 != 0 {
		goto L43
	} else {
		goto L44
	}
L33:
	;
	v115 = int32(4)
	v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+1)))
	if v117&int32(254) == int32(2) {
		goto L36
	} else {
		goto L37
	}
L34:
	;
	goto L35
L35:
	;
	v130 = int32(1)
	if v112&v130 != 0 {
		v142 = int32(base.Ui32(v112)>>(uint(v130)%32)) - v130
		goto L32
	} else {
		goto L42
	}
L36:
	;
	v126 = v115
	goto L38
L37:
	;
	v126 = base.B2i32(v117 == int32(18)) << (uint(v115) % 32)
	goto L38
L38:
	;
	if v117 == int32(1) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v129 = v115
	goto L41
L40:
	;
	v129 = v126
	goto L41
L41:
	;
	v142 = v129
	goto L32
L42:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	v142 = int32(base.Ui32(v136)>>(uint(int32(2))%32)) - int32(4)
	goto L32
L43:
	;
	v147 = v143
	goto L45
L44:
	;
	v147 = int32(4)
	goto L45
L45:
	;
	v148 = v13 + v147
	v149 = int32(0)
	if base.Ui32(v104) <= base.Ui32(int32(41)) {
		goto L48
	} else {
		goto L49
	}
L46:
	;
	m.G0 = v10 + int32(16)
	return v313
L47:
	;
	if v159 <= int32(1) {
		goto L51
	} else {
		goto L52
	}
L48:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v104*int32(28))+uint32(_c_F_length_in_encoding[0])))
	v159 = v158
	goto L50
L49:
	;
	v159 = int32(1)
	goto L50
L50:
	;
	goto L47
L51:
	;
	v163 = int32(0)
	v167 = base.B2i32(v142 != v163)
	if v148&int32(3) == v163 {
		v193 = v148
		v195 = v142
		v196 = v167
		goto L57
	} else {
		goto L58
	}
L52:
	;
	goto L53
L53:
	;
	if v142 <= int32(0) {
		v313 = v149
		goto L46
	} else {
		goto L85
	}
L54:
	;
	if v266 != 0 {
		goto L80
	} else {
		goto L81
	}
L55:
	;
	v266 = int32(0)
	goto L54
L56:
	;
	v244 = v237
	v246 = v239
	goto L74
L57:
	;
	if v196 == int32(0) {
		goto L55
	} else {
		goto L65
	}
L58:
	;
	if v142 == int32(0) {
		v193 = v148
		v195 = v142
		v196 = v167
		goto L57
	} else {
		goto L59
	}
L59:
	;
	v176 = v148
	v178 = v142
	goto L60
L60:
	;
	v181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v176))))
	if v181 == int32(0) {
		v237 = v176
		v239 = v178
		goto L56
	} else {
		goto L62
	}
L61:
	;
	v193 = v188
	v195 = v184
	v196 = v186
	goto L57
L62:
	;
	v183 = int32(1)
	v184 = v178 - v183
	v185 = int32(0)
	v186 = base.B2i32(v184 != v185)
	v188 = v176 + v183
	if v188&int32(3) == v185 {
		v193 = v188
		v195 = v184
		v196 = v186
		goto L57
	} else {
		goto L63
	}
L63:
	;
	if v184 != 0 {
		v176 = v188
		v178 = v184
		goto L60
	} else {
		goto L64
	}
L64:
	;
	goto L61
L65:
	;
	v200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v193))))
	if v200 == int32(0) {
		v230 = v193
		v232 = v195
		goto L66
	} else {
		goto L67
	}
L66:
	;
	if v232 == int32(0) {
		goto L55
	} else {
		goto L73
	}
L67:
	;
	if base.Ui32(v195) < base.Ui32(int32(4)) {
		v230 = v193
		v232 = v195
		goto L66
	} else {
		goto L68
	}
L68:
	;
	v210 = v193
	v212 = v195
	goto L69
L69:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v210)))
	v217 = v216 ^ int32(0)
	v220 = int32(-2139062144)
	if (int32(16843008)-v217|v217)&v220 != v220 {
		v237 = v210
		v239 = v212
		goto L56
	} else {
		goto L71
	}
L70:
	;
	v230 = v225
	v232 = v227
	goto L66
L71:
	;
	v224 = int32(4)
	v225 = v210 + v224
	v227 = v212 - v224
	if base.Ui32(int32(3)) < base.Ui32(v227) {
		v210 = v225
		v212 = v227
		goto L69
	} else {
		goto L72
	}
L72:
	;
	goto L70
L73:
	;
	v237 = v230
	v239 = v232
	goto L56
L74:
	;
	v249 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v244))))
	if int32(0) == v249 {
		goto L76
	} else {
		goto L77
	}
L75:
	;
	goto L55
L76:
	;
	v266 = v244
	goto L54
L77:
	;
	goto L78
L78:
	;
	v251 = int32(1)
	v254 = v246 - v251
	if v254 != 0 {
		v244 = v244 + v251
		v246 = v254
		goto L74
	} else {
		goto L79
	}
L79:
	;
	goto L75
L80:
	;
	v267 = int32(-1)
	goto L82
L81:
	;
	v267 = v142
	goto L82
L82:
	;
	if v266 == int32(0) {
		v313 = v267
		goto L46
	} else {
		goto L83
	}
L83:
	;
	F_report_invalid_encoding(m, v104, v266, int32(1))
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L1
	} else {
		goto L84
	}
L84:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L85:
	;
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v104*int32(28))+uint32(_c_F_length_in_encoding[1])))
	v281 = v142
	v282 = v148
	v283 = v149
	goto L86
L86:
	;
	v287 = int32(*(*int8)(unsafe.Add(mBase, uint32(v282))))
	if int32(0) <= v287 {
		goto L91
	} else {
		goto L92
	}
L87:
	;
	v313 = v307
	goto L46
L88:
	;
	v307 = v283 + int32(1)
	if int32(0) < v303 {
		v281 = v303
		v282 = v305
		v283 = v307
		goto L86
	} else {
		goto L100
	}
L89:
	;
	v303 = v281 - v294
	v305 = v282 + v294
	goto L88
L90:
	;
	F_report_invalid_encoding(m, v104, v282, v281)
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L1
	} else {
		goto L99
	}
L91:
	;
	if v287 != 0 {
		goto L94
	} else {
		goto L95
	}
L92:
	;
	goto L93
L93:
	;
	v294 = m.T0[v279].(func(*base.Module, int32, int32) int32)(m, v282, v281)
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L1
	} else {
		goto L97
	}
L94:
	;
	v290 = int32(1)
	v303 = v281 - v290
	v305 = v282 + v290
	goto L88
L95:
	;
	goto L96
L96:
	;
	goto L90
L97:
	;
	if int32(0) <= v294 {
		goto L89
	} else {
		goto L98
	}
L98:
	;
	goto L90
L99:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L100:
	;
	goto L87
L101:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L1
	} else {
		goto L102
	}
L102:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v17
	F_errmsg(m, int32(_a_F_length_in_encoding_2), v10)
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L1
	} else {
		goto L103
	}
L103:
	;
	F_errfinish(m, int32(_a_F_length_in_encoding_3), int32(637), int32(_a_F_length_in_encoding_4))
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L1
	} else {
		goto L104
	}
L104:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_lexescape(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v74 int32
	_ = v74
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v230 int32
	_ = v230
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v314 int32
	_ = v314
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v350 int32
	_ = v350
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v368 int32
	_ = v368
	var v373 int32
	_ = v373
	var v375 int32
	_ = v375
	var v377 int32
	_ = v377
	var v392 int32
	_ = v392
	var v395 int32
	_ = v395
	var v397 int32
	_ = v397
	var v402 int32
	_ = v402
	var v404 int32
	_ = v404
	var v409 int32
	_ = v409
	var v411 int32
	_ = v411
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v425 int32
	_ = v425
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v456 int32
	_ = v456
	var v458 int32
	_ = v458
	var v460 int32
	_ = v460
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v470 int32
	_ = v470
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v478 int32
	_ = v478
	var v483 int32
	_ = v483
	var v485 int32
	_ = v485
	var v487 int32
	_ = v487
	var v502 int32
	_ = v502
	var v505 int32
	_ = v505
	var v507 int32
	_ = v507
	var v512 int32
	_ = v512
	var v514 int32
	_ = v514
	var v519 int32
	_ = v519
	var v521 int32
	_ = v521
	var v524 int32
	_ = v524
	var v529 int32
	_ = v529
	var v533 int32
	_ = v533
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v553 int32
	_ = v553
	var v555 int32
	_ = v555
	var v557 int32
	_ = v557
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v567 int32
	_ = v567
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v575 int32
	_ = v575
	var v580 int32
	_ = v580
	var v582 int32
	_ = v582
	var v584 int32
	_ = v584
	var v599 int32
	_ = v599
	var v602 int32
	_ = v602
	var v604 int32
	_ = v604
	var v609 int32
	_ = v609
	var v611 int32
	_ = v611
	var v616 int32
	_ = v616
	var v618 int32
	_ = v618
	var v621 int32
	_ = v621
	var v628 int32
	_ = v628
	var v634 int32
	_ = v634
	var v642 int32
	_ = v642
	var v644 int32
	_ = v644
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v4 + int32(4)
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
	v10 = v8 - int32(48)
	if base.Ui32(v10) < base.Ui32(int32(10)) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+8)) = v25 | int32(128)
	switch v10 {
	case 0:
		goto L5
	case 1, 2, 3, 4, 5, 6, 7, 8, 9:
		goto L6
	default:
		goto L4
	case 17:
		goto L30
	case 18:
		goto L28
	case 20:
		goto L25
	case 29:
		goto L21
	case 35:
		goto L17
	case 37:
		goto L14
	case 39:
		goto L11
	case 41:
		goto L8
	case 42:
		goto L7
	case 49:
		goto L31
	case 50:
		goto L29
	case 51:
		goto L27
	case 52:
		goto L26
	case 53:
		goto L24
	case 54:
		goto L23
	case 61:
		goto L22
	case 62:
		goto L20
	case 66:
		goto L19
	case 67:
		goto L18
	case 68:
		goto L16
	case 69:
		goto L15
	case 70:
		goto L13
	case 71:
		goto L12
	case 72:
		goto L10
	case 73:
		goto L9
	}
L2:
	;
	if base.Ui32(v8&int32(-33)-int32(65)) < base.Ui32(int32(26)) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v8
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(112)
	return int32(1)
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
	v642 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v642 != 0 {
		goto L154
	} else {
		goto L155
	}
L5:
	;
	v548 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v549 = *(*int32)(unsafe.Add(mBase, uint32(v548)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v548)+8)) = v549 | int32(512)
	v553 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v555 = v553 - int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v555
	v557 = int32(8)
	v559 = int32(3)
	v560 = int32(0)
	v567 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v572 = v555
	v573 = v560
	v575 = v560
	goto L135
L6:
	;
	v456 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v458 = v456 - int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v458
	v460 = int32(10)
	v462 = int32(255)
	v463 = int32(0)
	v470 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v475 = v458
	v476 = v463
	v478 = v463
	goto L110
L7:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(90)
	return int32(1)
L8:
	;
	v443 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v444 = *(*int32)(unsafe.Add(mBase, uint32(v443)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v443)+8)) = v444 | int32(1024)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(87)
	return int32(1)
L9:
	;
	v434 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v435 = *(*int32)(unsafe.Add(mBase, uint32(v434)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v434)+8)) = v435 | int32(1024)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(119)
	return int32(1)
L10:
	;
	v345 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v345)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v345)+8)) = v346 | int32(512)
	v350 = int32(16)
	v352 = int32(255)
	v353 = int32(0)
	v359 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v360 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v365 = v359
	v366 = v353
	v368 = v353
	goto L88
L11:
	;
	v336 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v336)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v336)+8)) = v337 | int32(1024)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(55834574947)
	return int32(1)
L12:
	;
	v327 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v327)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v327)+8)) = v328 | int32(1024)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(55834574963)
	return int32(1)
L13:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(47244640368)
	return int32(1)
L14:
	;
	v239 = int32(16)
	v240 = int32(8)
	v242 = int32(0)
	v248 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v249 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v254 = v248
	v255 = v242
	v257 = v242
	goto L66
L15:
	;
	v155 = int32(16)
	v156 = int32(4)
	v158 = int32(0)
	v164 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v165 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v170 = v164
	v171 = v158
	v173 = v158
	goto L44
L16:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(38654705776)
	return int32(1)
L17:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v142)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v142)+8)) = v143 | int32(1024)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(42949673059)
	return int32(1)
L18:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v133)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v133)+8)) = v134 | int32(1024)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(42949673075)
	return int32(1)
L19:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(55834574960)
	return int32(1)
L20:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(42949673072)
	return int32(1)
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(62)
	return int32(1)
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(60)
	return int32(1)
L23:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(51539607664)
	return int32(1)
L24:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v98)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v98)+8)) = v99 | int32(512)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(112)
	v108 = F_chrnamed(m, l0, int32(_a_F_lexescape_0), int32(_a_F_lexescape_1), int32(27))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L32
	} else {
		goto L40
	}
L25:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v89)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v89)+8)) = v90 | int32(1024)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(21474836579)
	return int32(1)
L26:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v80)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v80)+8)) = v81 | int32(1024)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(21474836595)
	return int32(1)
L27:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v53)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v53)+8)) = v54 | int32(512)
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if base.Ui32(v59) <= base.Ui32(v58) {
		goto L34
	} else {
		goto L35
	}
L28:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(395136991344)
	return int32(1)
L29:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(34359738480)
	return int32(1)
L30:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(65)
	return int32(1)
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(112)
	v34 = F_chrnamed(m, l0, int32(_a_F_lexescape_2), int32(_a_F_lexescape_0), int32(7))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	return int32(0)
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v34
	return int32(1)
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v63 != 0 {
		goto L37
	} else {
		goto L38
	}
L35:
	;
	goto L36
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(112)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v58 + int32(4)
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v58)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v74 & int32(31)
	return int32(1)
L37:
	;
	v65 = v63
	goto L39
L38:
	;
	v65 = int32(5)
	goto L39
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v65
	return int32(0)
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v108
	return int32(1)
L41:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v220 = int32(0)
	if base.B2i32(v219 == v220)&base.B2i32(base.Ui32(v209) < base.Ui32(int32(2147483647))) == v220 {
		goto L57
	} else {
		goto L58
	}
L42:
	;
	if base.Ui32(v207) < base.Ui32(v156) {
		goto L51
	} else {
		goto L52
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v170
	v207 = v171
	v209 = v173
	goto L42
L44:
	;
	if base.Ui32(v165) <= base.Ui32(v170) {
		v207 = v171
		v209 = v173
		goto L42
	} else {
		goto L46
	}
L45:
	;
	v207 = v156
	v209 = v200
	goto L42
L46:
	;
	v178 = v170 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v178
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v170)))
	v182 = v180 - int32(48)
	if base.Ui32(int32(54)) < base.Ui32(v182) {
		goto L43
	} else {
		goto L47
	}
L47:
	;
	if base.I32_wrap_i64(int64(base.Ui64(int64(35465847073801215))>>(uint(base.I64_extend_i32_u(v182))%64)))&int32(1) == int32(0) {
		goto L43
	} else {
		goto L48
	}
L48:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v182<<(uint(int32(2))%32))+uint32(_c_F_lexescape[0])))
	if base.Ui32(v155) <= base.Ui32(v197) {
		goto L43
	} else {
		goto L49
	}
L49:
	;
	v200 = v197 + v155*v173
	v202 = v171 + int32(1)
	if v202 != v156 {
		v170 = v178
		v171 = v202
		v173 = v200
		goto L44
	} else {
		goto L50
	}
L50:
	;
	goto L45
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
	v214 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v214 != 0 {
		goto L54
	} else {
		goto L55
	}
L52:
	;
	goto L53
L53:
	;
	goto L41
L54:
	;
	v216 = v214
	goto L56
L55:
	;
	v216 = int32(5)
	goto L56
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v216
	goto L53
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
	if v219 != 0 {
		goto L60
	} else {
		goto L61
	}
L58:
	;
	goto L59
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v209
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(112)
	return int32(1)
L60:
	;
	v230 = v219
	goto L62
L61:
	;
	v230 = int32(5)
	goto L62
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v230
	return int32(0)
L63:
	;
	v303 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v304 = int32(0)
	if base.B2i32(v303 == v304)&base.B2i32(base.Ui32(v293) < base.Ui32(int32(2147483647))) == v304 {
		goto L79
	} else {
		goto L80
	}
L64:
	;
	if base.Ui32(v291) < base.Ui32(v240) {
		goto L73
	} else {
		goto L74
	}
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v254
	v291 = v255
	v293 = v257
	goto L64
L66:
	;
	if base.Ui32(v249) <= base.Ui32(v254) {
		v291 = v255
		v293 = v257
		goto L64
	} else {
		goto L68
	}
L67:
	;
	v291 = v240
	v293 = v284
	goto L64
L68:
	;
	v262 = v254 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v262
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v254)))
	v266 = v264 - int32(48)
	if base.Ui32(int32(54)) < base.Ui32(v266) {
		goto L65
	} else {
		goto L69
	}
L69:
	;
	if base.I32_wrap_i64(int64(base.Ui64(int64(35465847073801215))>>(uint(base.I64_extend_i32_u(v266))%64)))&int32(1) == int32(0) {
		goto L65
	} else {
		goto L70
	}
L70:
	;
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v266<<(uint(int32(2))%32))+uint32(_c_F_lexescape[0])))
	if base.Ui32(v239) <= base.Ui32(v281) {
		goto L65
	} else {
		goto L71
	}
L71:
	;
	v284 = v281 + v239*v257
	v286 = v255 + int32(1)
	if v286 != v240 {
		v254 = v262
		v255 = v286
		v257 = v284
		goto L66
	} else {
		goto L72
	}
L72:
	;
	goto L67
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
	v298 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v298 != 0 {
		goto L76
	} else {
		goto L77
	}
L74:
	;
	goto L75
L75:
	;
	goto L63
L76:
	;
	v300 = v298
	goto L78
L77:
	;
	v300 = int32(5)
	goto L78
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v300
	goto L75
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
	if v303 != 0 {
		goto L82
	} else {
		goto L83
	}
L80:
	;
	goto L81
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v293
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(112)
	return int32(1)
L82:
	;
	v314 = v303
	goto L84
L83:
	;
	v314 = int32(5)
	goto L84
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v314
	return int32(0)
L85:
	;
	v414 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v415 = int32(0)
	if base.B2i32(v414 == v415)&base.B2i32(base.Ui32(v404) < base.Ui32(int32(2147483647))) == v415 {
		goto L101
	} else {
		goto L102
	}
L86:
	;
	if base.Ui32(v402) < base.Ui32(int32(1)) {
		goto L95
	} else {
		goto L96
	}
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v365
	v402 = v366
	v404 = v368
	goto L86
L88:
	;
	if base.Ui32(v360) <= base.Ui32(v365) {
		v402 = v366
		v404 = v368
		goto L86
	} else {
		goto L90
	}
L89:
	;
	v402 = v352
	v404 = v395
	goto L86
L90:
	;
	v373 = v365 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v373
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v365)))
	v377 = v375 - int32(48)
	if base.Ui32(int32(54)) < base.Ui32(v377) {
		goto L87
	} else {
		goto L91
	}
L91:
	;
	if base.I32_wrap_i64(int64(base.Ui64(int64(35465847073801215))>>(uint(base.I64_extend_i32_u(v377))%64)))&int32(1) == int32(0) {
		goto L87
	} else {
		goto L92
	}
L92:
	;
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v377<<(uint(int32(2))%32))+uint32(_c_F_lexescape[0])))
	if base.Ui32(v350) <= base.Ui32(v392) {
		goto L87
	} else {
		goto L93
	}
L93:
	;
	v395 = v392 + v350*v368
	v397 = v366 + int32(1)
	if v397 != v352 {
		v365 = v373
		v366 = v397
		v368 = v395
		goto L88
	} else {
		goto L94
	}
L94:
	;
	goto L89
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
	v409 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v409 != 0 {
		goto L98
	} else {
		goto L99
	}
L96:
	;
	goto L97
L97:
	;
	goto L85
L98:
	;
	v411 = v409
	goto L100
L99:
	;
	v411 = int32(5)
	goto L100
L100:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v411
	goto L97
L101:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
	if v414 != 0 {
		goto L104
	} else {
		goto L105
	}
L102:
	;
	goto L103
L103:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v404
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(112)
	return int32(1)
L104:
	;
	v425 = v414
	goto L106
L105:
	;
	v425 = int32(5)
	goto L106
L106:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v425
	return int32(0)
L107:
	;
	v524 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v524 != 0 {
		goto L123
	} else {
		goto L124
	}
L108:
	;
	if base.Ui32(v512) < base.Ui32(int32(1)) {
		goto L117
	} else {
		goto L118
	}
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v475
	v512 = v476
	v514 = v478
	goto L108
L110:
	;
	if base.Ui32(v470) <= base.Ui32(v475) {
		v512 = v476
		v514 = v478
		goto L108
	} else {
		goto L112
	}
L111:
	;
	v512 = v462
	v514 = v505
	goto L108
L112:
	;
	v483 = v475 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v483
	v485 = *(*int32)(unsafe.Add(mBase, uint32(v475)))
	v487 = v485 - int32(48)
	if base.Ui32(int32(54)) < base.Ui32(v487) {
		goto L109
	} else {
		goto L113
	}
L113:
	;
	if base.I32_wrap_i64(int64(base.Ui64(int64(35465847073801215))>>(uint(base.I64_extend_i32_u(v487))%64)))&int32(1) == int32(0) {
		goto L109
	} else {
		goto L114
	}
L114:
	;
	v502 = *(*int32)(unsafe.Add(mBase, uint32(v487<<(uint(int32(2))%32))+uint32(_c_F_lexescape[0])))
	if base.Ui32(v460) <= base.Ui32(v502) {
		goto L109
	} else {
		goto L115
	}
L115:
	;
	v505 = v502 + v460*v478
	v507 = v476 + int32(1)
	if v507 != v462 {
		v475 = v483
		v476 = v507
		v478 = v505
		goto L110
	} else {
		goto L116
	}
L116:
	;
	goto L111
L117:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
	v519 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v519 != 0 {
		goto L120
	} else {
		goto L121
	}
L118:
	;
	goto L119
L119:
	;
	goto L107
L120:
	;
	v521 = v519
	goto L122
L121:
	;
	v521 = int32(5)
	goto L122
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v521
	goto L119
L123:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
	return int32(0)
L124:
	;
	goto L125
L125:
	;
	v529 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v456 != v529 {
		goto L127
	} else {
		goto L128
	}
L126:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v456
	goto L5
L127:
	;
	if v514 <= int32(0) {
		goto L126
	} else {
		goto L130
	}
L128:
	;
	goto L129
L129:
	;
	v535 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v536 = *(*int32)(unsafe.Add(mBase, uint32(v535)+8))
	v537 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v535)+8)) = v536 | v537
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v514
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(98)
	return v537
L130:
	;
	v533 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v533 < v514 {
		goto L126
	} else {
		goto L131
	}
L131:
	;
	goto L129
L132:
	;
	v621 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v621 != 0 {
		goto L148
	} else {
		goto L149
	}
L133:
	;
	if base.Ui32(v609) < base.Ui32(int32(1)) {
		goto L142
	} else {
		goto L143
	}
L134:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v572
	v609 = v573
	v611 = v575
	goto L133
L135:
	;
	if base.Ui32(v567) <= base.Ui32(v572) {
		v609 = v573
		v611 = v575
		goto L133
	} else {
		goto L137
	}
L136:
	;
	v609 = v559
	v611 = v602
	goto L133
L137:
	;
	v580 = v572 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v580
	v582 = *(*int32)(unsafe.Add(mBase, uint32(v572)))
	v584 = v582 - int32(48)
	if base.Ui32(int32(54)) < base.Ui32(v584) {
		goto L134
	} else {
		goto L138
	}
L138:
	;
	if base.I32_wrap_i64(int64(base.Ui64(int64(35465847073801215))>>(uint(base.I64_extend_i32_u(v584))%64)))&int32(1) == int32(0) {
		goto L134
	} else {
		goto L139
	}
L139:
	;
	v599 = *(*int32)(unsafe.Add(mBase, uint32(v584<<(uint(int32(2))%32))+uint32(_c_F_lexescape[0])))
	if base.Ui32(v557) <= base.Ui32(v599) {
		goto L134
	} else {
		goto L140
	}
L140:
	;
	v602 = v599 + v557*v575
	v604 = v573 + int32(1)
	if v604 != v559 {
		v572 = v580
		v573 = v604
		v575 = v602
		goto L135
	} else {
		goto L141
	}
L141:
	;
	goto L136
L142:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
	v616 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v616 != 0 {
		goto L145
	} else {
		goto L146
	}
L143:
	;
	goto L144
L144:
	;
	goto L132
L145:
	;
	v618 = v616
	goto L147
L146:
	;
	v618 = int32(5)
	goto L147
L147:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v618
	goto L144
L148:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
	return int32(0)
L149:
	;
	goto L150
L150:
	;
	if base.Ui32(int32(256)) <= base.Ui32(v611) {
		goto L151
	} else {
		goto L152
	}
L151:
	;
	v628 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v628 - int32(4)
	v634 = int32(base.Ui32(v611) >> (uint(int32(3)) % 32))
	goto L153
L152:
	;
	v634 = v611
	goto L153
L153:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v634
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(112)
	return int32(1)
L154:
	;
	v644 = v642
	goto L156
L155:
	;
	v644 = int32(5)
	goto L156
L156:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v644
	return int32(0)
}
func F_lithuanian_UTF_8_stem(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v179 int32
	_ = v179
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v230 int32
	_ = v230
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v250 int32
	_ = v250
	var v256 int32
	_ = v256
	var v263 int32
	_ = v263
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v301 int32
	_ = v301
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v314 int32
	_ = v314
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v323 int32
	_ = v323
	var v333 int32
	_ = v333
	var v335 int32
	_ = v335
	var v339 int32
	_ = v339
	var v352 int32
	_ = v352
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v372 int32
	_ = v372
	var v378 int32
	_ = v378
	var v386 int32
	_ = v386
	var v397 int32
	_ = v397
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v406 int32
	_ = v406
	var v412 int32
	_ = v412
	var v416 int32
	_ = v416
	var v429 int32
	_ = v429
	var v432 int32
	_ = v432
	var v435 int32
	_ = v435
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v488 int32
	_ = v488
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v494 int32
	_ = v494
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v504 int32
	_ = v504
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v511 int32
	_ = v511
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v517 int32
	_ = v517
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v527 int32
	_ = v527
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v535 int32
	_ = v535
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v541 int32
	_ = v541
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v548 int32
	_ = v548
	var v550 int32
	_ = v550
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v559 int32
	_ = v559
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v568 int32
	_ = v568
	var v572 int32
	_ = v572
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v579 int32
	_ = v579
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v4))) = v5
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v5 == v7 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v170 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v171 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v179 = v169
	goto L50
L2:
	;
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9+v7))))
	if v11 != int32(97) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v14 = int32(0)
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v9-int32(4))))
	if v21 == v14 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	if int32(7) <= v93 {
		goto L21
	} else {
		goto L22
	}
L5:
	;
	v93 = int32(0)
	goto L4
L6:
	;
	goto L7
L7:
	;
	v26 = v21 & int32(3)
	if base.Ui32(v21) < base.Ui32(int32(4)) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	if v26 != 0 {
		goto L15
	} else {
		goto L16
	}
L9:
	;
	v60 = v9
	v61 = int32(0)
	goto L8
L10:
	;
	goto L11
L11:
	;
	v33 = v9
	v34 = int32(0)
	v37 = v14
	goto L12
L12:
	;
	v39 = int32(*(*int8)(unsafe.Add(mBase, uint32(v33))))
	v40 = int32(-65)
	v43 = int32(*(*int8)(unsafe.Add(mBase, uint32(v33)+1)))
	v47 = int32(*(*int8)(unsafe.Add(mBase, uint32(v33)+2)))
	v51 = int32(*(*int8)(unsafe.Add(mBase, uint32(v33)+3)))
	v54 = v34 + base.B2i32(v40 < v39) + base.B2i32(v40 < v43) + base.B2i32(v40 < v47) + base.B2i32(v40 < v51)
	v55 = int32(4)
	v56 = v33 + v55
	v58 = v37 + v55
	if v58 != v21&int32(-4) {
		v33 = v56
		v34 = v54
		v37 = v58
		goto L12
	} else {
		goto L14
	}
L13:
	;
	v60 = v56
	v61 = v54
	goto L8
L14:
	;
	goto L13
L15:
	;
	v66 = v60
	v67 = v61
	v69 = v14
	goto L18
L16:
	;
	v82 = v61
	goto L17
L17:
	;
	v93 = v82
	goto L4
L18:
	;
	v72 = int32(*(*int8)(unsafe.Add(mBase, uint32(v66))))
	v75 = v67 + base.B2i32(int32(-65) < v72)
	v76 = int32(1)
	v79 = v69 + v76
	if v79 != v26 {
		v66 = v66 + v76
		v67 = v75
		v69 = v79
		goto L18
	} else {
		goto L20
	}
L19:
	;
	v82 = v75
	goto L17
L20:
	;
	goto L19
L21:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	goto L26
L22:
	;
	v155 = v7
	goto L23
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v155
	goto L1
L24:
	;
	if v150 < int32(0) {
		goto L44
	} else {
		goto L45
	}
L26:
	;
	goto L27
L27:
	;
	goto L28
L28:
	;
	v105 = v97
	v107 = int32(1)
	goto L31
L30:
	;
	v150 = v135
	goto L24
L31:
	;
	if v98 <= v105 {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	goto L30
L33:
	;
	v150 = int32(-1)
	goto L24
L34:
	;
	goto L35
L35:
	;
	v112 = v105 + int32(1)
	v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96+v105))))
	if base.Ui32(v114) < base.Ui32(int32(192)) {
		v135 = v112
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v136 = int32(1)
	if v136 < v107 {
		v105 = v135
		v107 = v107 - v136
		goto L31
	} else {
		goto L43
	}
L37:
	;
	if v98 <= v112 {
		v135 = v112
		goto L36
	} else {
		goto L38
	}
L38:
	;
	v121 = v112
	goto L39
L39:
	;
	v124 = int32(*(*int8)(unsafe.Add(mBase, uint32(v96+v121))))
	if int32(-65) < v124 {
		v135 = v121
		goto L36
	} else {
		goto L41
	}
L40:
	;
	v135 = v98
	goto L36
L41:
	;
	v128 = v121 + int32(1)
	if v128 != v98 {
		v121 = v128
		goto L39
	} else {
		goto L42
	}
L42:
	;
	goto L40
L43:
	;
	goto L32
L44:
	;
	v153 = v7
	goto L46
L45:
	;
	v153 = v150
	goto L46
L46:
	;
	v155 = v153
	goto L23
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v7
	v406 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v406
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v406
	if v406-int32(3) <= v7 {
		goto L100
	} else {
		goto L101
	}
L48:
	;
	if v274 < int32(0) {
		goto L47
	} else {
		goto L73
	}
L49:
	;
	v274 = v246
	goto L48
L50:
	;
	if v170 <= v179 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v274 = int32(-1)
	goto L48
L53:
	;
	goto L54
L54:
	;
	v186 = int32(1)
	v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v179+v171))))
	if base.Ui32(v188) < base.Ui32(int32(192)) {
		v245 = v188
		v246 = v186
		goto L55
	} else {
		goto L56
	}
L55:
	;
	if int32(371) < v245 {
		goto L68
	} else {
		goto L69
	}
L56:
	;
	v192 = v179 + int32(1)
	if v192 == v170 {
		v245 = v188
		v246 = v186
		goto L55
	} else {
		goto L57
	}
L57:
	;
	v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192+v171))))
	v197 = v195 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v188) {
		goto L59
	} else {
		goto L60
	}
L58:
	;
	v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v201+v171))))
	v213 = v211 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v188) {
		goto L64
	} else {
		goto L65
	}
L59:
	;
	v201 = v179 + int32(2)
	if v201 != v170 {
		goto L58
	} else {
		goto L62
	}
L60:
	;
	goto L61
L61:
	;
	v245 = v188<<(uint(int32(6))%32)&int32(1984) | v197
	v246 = int32(2)
	goto L55
L62:
	;
	goto L61
L63:
	;
	v230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v171+v217))))
	v245 = v230&int32(63) | (v188<<(uint(int32(18))%32)&int32(_a_F_lithuanian_UTF_8_stem_0) | v197<<(uint(int32(12))%32) | v213<<(uint(int32(6))%32))
	v246 = int32(4)
	goto L55
L64:
	;
	v217 = v179 + int32(3)
	if v217 != v170 {
		goto L63
	} else {
		goto L67
	}
L65:
	;
	goto L66
L66:
	;
	v245 = v188<<(uint(int32(12))%32)&int32(_a_F_lithuanian_UTF_8_stem_1) | v197<<(uint(int32(6))%32) | v213
	v246 = int32(3)
	goto L55
L67:
	;
	goto L66
L68:
	;
	v263 = v246 + v179
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v263
	v179 = v263
	goto L50
L69:
	;
	v250 = v245 - int32(97)
	if v250 < int32(0) {
		goto L68
	} else {
		goto L70
	}
L70:
	;
	v256 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v250)>>(uint(int32(3))%32)))+uint32(_c_F_lithuanian_UTF_8_stem[0]))))
	if int32(base.Ui32(v256)>>(uint(v250&int32(7))%32))&int32(1) != 0 {
		goto L49
	} else {
		goto L71
	}
L71:
	;
	goto L68
L73:
	;
	v277 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v278 = v277 + v274
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v278
	v292 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v293 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v301 = v278
	goto L76
L74:
	;
	if v397 < int32(0) {
		goto L47
	} else {
		goto L98
	}
L75:
	;
	v397 = v368
	goto L74
L76:
	;
	if v292 <= v301 {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v397 = int32(-1)
	goto L74
L79:
	;
	goto L80
L80:
	;
	v308 = int32(1)
	v310 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v301+v293))))
	if base.Ui32(v310) < base.Ui32(int32(192)) {
		v367 = v310
		v368 = v308
		goto L81
	} else {
		goto L82
	}
L81:
	;
	if int32(371) < v367 {
		goto L75
	} else {
		goto L94
	}
L82:
	;
	v314 = v301 + int32(1)
	if v314 == v292 {
		v367 = v310
		v368 = v308
		goto L81
	} else {
		goto L83
	}
L83:
	;
	v317 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v314+v293))))
	v319 = v317 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v310) {
		goto L85
	} else {
		goto L86
	}
L84:
	;
	v333 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v323+v293))))
	v335 = v333 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v310) {
		goto L90
	} else {
		goto L91
	}
L85:
	;
	v323 = v301 + int32(2)
	if v323 != v292 {
		goto L84
	} else {
		goto L88
	}
L86:
	;
	goto L87
L87:
	;
	v367 = v310<<(uint(int32(6))%32)&int32(1984) | v319
	v368 = int32(2)
	goto L81
L88:
	;
	goto L87
L89:
	;
	v352 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v293+v339))))
	v367 = v352&int32(63) | (v310<<(uint(int32(18))%32)&int32(_a_F_lithuanian_UTF_8_stem_0) | v319<<(uint(int32(12))%32) | v335<<(uint(int32(6))%32))
	v368 = int32(4)
	goto L81
L90:
	;
	v339 = v301 + int32(3)
	if v339 != v292 {
		goto L89
	} else {
		goto L93
	}
L91:
	;
	goto L92
L92:
	;
	v367 = v310<<(uint(int32(12))%32)&int32(_a_F_lithuanian_UTF_8_stem_1) | v319<<(uint(int32(6))%32) | v335
	v368 = int32(3)
	goto L81
L93:
	;
	goto L92
L94:
	;
	v372 = v367 - int32(97)
	if v372 < int32(0) {
		goto L75
	} else {
		goto L95
	}
L95:
	;
	v378 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v372)>>(uint(int32(3))%32)))+uint32(_c_F_lithuanian_UTF_8_stem[0]))))
	if int32(base.Ui32(v378)>>(uint(v372&int32(7))%32))&int32(1) == int32(0) {
		goto L75
	} else {
		goto L96
	}
L96:
	;
	v386 = v368 + v301
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v386
	v301 = v386
	goto L76
L98:
	;
	v400 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v401 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v400))) = v401 + v397
	goto L47
L99:
	;
	return v579
L100:
	;
	v488 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v488
	v490 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v491 = *(*int32)(unsafe.Add(mBase, uint32(v490)))
	if v488 < v491 {
		goto L131
	} else {
		goto L132
	}
L101:
	;
	v412 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v416 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v412+v406-int32(1)))))
	if v416&int32(224) != int32(96) {
		goto L100
	} else {
		goto L102
	}
L102:
	;
	if int32(1)<<(uint(v416)%32)&int32(_a_F_lithuanian_UTF_8_stem_2) == int32(0) {
		goto L100
	} else {
		goto L103
	}
L103:
	;
	v429 = F_find_among_b(m, l0, int32(_a_F_lithuanian_UTF_8_stem_3), int32(11))
	mBase = m.M
	v432 = m.ExcPending
	if v432 != 0 {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	return int32(0)
L105:
	;
	if v429 == int32(0) {
		goto L100
	} else {
		goto L106
	}
L106:
	;
	v435 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v435
	switch v429 - int32(1) {
	case 0:
		goto L114
	case 1:
		goto L113
	case 2:
		goto L112
	case 3:
		goto L111
	case 4:
		goto L110
	case 5:
		goto L109
	case 6:
		goto L108
	case 7:
		goto L107
	default:
		goto L100
	}
L107:
	;
	v483 = F_slice_from_s(m, l0, int32(6), int32(_a_F_lithuanian_UTF_8_stem_4))
	mBase = m.M
	v484 = m.ExcPending
	if v484 != 0 {
		goto L104
	} else {
		goto L129
	}
L108:
	;
	v477 = F_slice_from_s(m, l0, int32(6), int32(_a_F_lithuanian_UTF_8_stem_5))
	mBase = m.M
	v478 = m.ExcPending
	if v478 != 0 {
		goto L104
	} else {
		goto L127
	}
L109:
	;
	v471 = F_slice_from_s(m, l0, int32(6), int32(_a_F_lithuanian_UTF_8_stem_6))
	mBase = m.M
	v472 = m.ExcPending
	if v472 != 0 {
		goto L104
	} else {
		goto L125
	}
L110:
	;
	v465 = F_slice_from_s(m, l0, int32(4), int32(_a_F_lithuanian_UTF_8_stem_7))
	mBase = m.M
	v466 = m.ExcPending
	if v466 != 0 {
		goto L104
	} else {
		goto L123
	}
L111:
	;
	v459 = F_slice_from_s(m, l0, int32(4), int32(_a_F_lithuanian_UTF_8_stem_8))
	mBase = m.M
	v460 = m.ExcPending
	if v460 != 0 {
		goto L104
	} else {
		goto L121
	}
L112:
	;
	v453 = F_slice_from_s(m, l0, int32(7), int32(_a_F_lithuanian_UTF_8_stem_9))
	mBase = m.M
	v454 = m.ExcPending
	if v454 != 0 {
		goto L104
	} else {
		goto L119
	}
L113:
	;
	v447 = F_slice_from_s(m, l0, int32(5), int32(_a_F_lithuanian_UTF_8_stem_10))
	mBase = m.M
	v448 = m.ExcPending
	if v448 != 0 {
		goto L104
	} else {
		goto L117
	}
L114:
	;
	v441 = F_slice_from_s(m, l0, int32(5), int32(_a_F_lithuanian_UTF_8_stem_11))
	mBase = m.M
	v442 = m.ExcPending
	if v442 != 0 {
		goto L104
	} else {
		goto L115
	}
L115:
	;
	if int32(0) <= v441 {
		goto L100
	} else {
		goto L116
	}
L116:
	;
	v579 = v441
	goto L99
L117:
	;
	if int32(0) <= v447 {
		goto L100
	} else {
		goto L118
	}
L118:
	;
	v579 = v447
	goto L99
L119:
	;
	if int32(0) <= v453 {
		goto L100
	} else {
		goto L120
	}
L120:
	;
	v579 = v453
	goto L99
L121:
	;
	if int32(0) <= v459 {
		goto L100
	} else {
		goto L122
	}
L122:
	;
	v579 = v459
	goto L99
L123:
	;
	if int32(0) <= v465 {
		goto L100
	} else {
		goto L124
	}
L124:
	;
	v579 = v465
	goto L99
L125:
	;
	if int32(0) <= v471 {
		goto L100
	} else {
		goto L126
	}
L126:
	;
	v579 = v471
	goto L99
L127:
	;
	if int32(0) <= v477 {
		goto L100
	} else {
		goto L128
	}
L128:
	;
	v579 = v477
	goto L99
L129:
	;
	if v483 < int32(0) {
		v579 = v483
		goto L99
	} else {
		goto L130
	}
L130:
	;
	goto L100
L131:
	;
	v511 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v511
	v513 = F_r_fix_chdz(m, l0)
	mBase = m.M
	v514 = m.ExcPending
	if v514 != 0 {
		goto L104
	} else {
		goto L139
	}
L132:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v488
	v494 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v491
	v498 = F_find_among_b(m, l0, int32(_a_F_lithuanian_UTF_8_stem_12), int32(204))
	mBase = m.M
	v499 = m.ExcPending
	if v499 != 0 {
		goto L104
	} else {
		goto L133
	}
L133:
	;
	if v498 == int32(0) {
		goto L134
	} else {
		goto L135
	}
L134:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v494
	goto L131
L135:
	;
	goto L136
L136:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v494
	v504 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v504
	v506 = F_slice_del(m, l0)
	mBase = m.M
	v507 = m.ExcPending
	if v507 != 0 {
		goto L104
	} else {
		goto L137
	}
L137:
	;
	if v506 < int32(0) {
		v579 = v506
		goto L99
	} else {
		goto L138
	}
L138:
	;
	goto L131
L139:
	;
	if v513 < int32(0) {
		v579 = v513
		goto L99
	} else {
		goto L140
	}
L140:
	;
	v517 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v517
	goto L141
L141:
	;
	v522 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v523 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v524 = *(*int32)(unsafe.Add(mBase, uint32(v523)))
	if v524 <= v522 {
		goto L144
	} else {
		goto L145
	}
L142:
	;
	v579 = v574
	goto L99
L143:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v527
	v572 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v572
	v574 = F_slice_del(m, l0)
	mBase = m.M
	v575 = m.ExcPending
	if v575 != 0 {
		goto L104
	} else {
		goto L158
	}
L144:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v522
	v527 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v524
	v531 = F_find_among_b(m, l0, int32(_a_F_lithuanian_UTF_8_stem_13), int32(62))
	mBase = m.M
	v532 = m.ExcPending
	if v532 != 0 {
		goto L104
	} else {
		goto L147
	}
L145:
	;
	goto L146
L146:
	;
	v535 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v535
	v537 = F_r_fix_chdz(m, l0)
	mBase = m.M
	v538 = m.ExcPending
	if v538 != 0 {
		goto L104
	} else {
		goto L149
	}
L147:
	;
	if v531 != 0 {
		goto L143
	} else {
		goto L148
	}
L148:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v527
	goto L146
L149:
	;
	if v537 < int32(0) {
		v579 = v537
		goto L99
	} else {
		goto L150
	}
L150:
	;
	v541 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v541
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v541
	v545 = v541 - int32(1)
	v546 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v545 <= v546 {
		goto L151
	} else {
		goto L152
	}
L151:
	;
	v568 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v568
	v579 = int32(1)
	goto L99
L152:
	;
	v548 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v550 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v548+v545))))
	if v550 != int32(100) {
		goto L151
	} else {
		goto L153
	}
L153:
	;
	v555 = F_find_among_b(m, l0, int32(_a_F_lithuanian_UTF_8_stem_14), int32(1))
	mBase = m.M
	v556 = m.ExcPending
	if v556 != 0 {
		goto L104
	} else {
		goto L154
	}
L154:
	;
	if v555 == int32(0) {
		goto L151
	} else {
		goto L155
	}
L155:
	;
	v559 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v559
	v563 = F_slice_from_s(m, l0, int32(1), int32(_a_F_lithuanian_UTF_8_stem_15))
	mBase = m.M
	v564 = m.ExcPending
	if v564 != 0 {
		goto L104
	} else {
		goto L156
	}
L156:
	;
	if v563 < int32(0) {
		v579 = v563
		goto L99
	} else {
		goto L157
	}
L157:
	;
	goto L151
L158:
	;
	if int32(0) <= v574 {
		goto L141
	} else {
		goto L159
	}
L159:
	;
	goto L142
}
func F_lock_twophase_recover(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
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
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v135 int32
	_ = v135
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v200 int32
	_ = v200
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v217 int32
	_ = v217
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v284 int32
	_ = v284
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v306 int32
	_ = v306
	var v313 int32
	_ = v313
	var v317 int32
	_ = v317
	var v322 int32
	_ = v322
	var v324 int32
	_ = v324
	var v328 int32
	_ = v328
	var v331 int32
	_ = v331
	var v335 int32
	_ = v335
	var v342 int32
	_ = v342
	var v347 int32
	_ = v347
	var v351 int32
	_ = v351
	var v355 int32
	_ = v355
	var v360 int32
	_ = v360
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v368 int64
	_ = v368
	var v369 int32
	_ = v369
	var v377 int32
	_ = v377
	var v382 int32
	_ = v382
	v14 = m.G0
	v16 = v14 - int32(80)
	m.G0 = v16
	v19 = F_TwoPhaseGetDummyProc(m, l0, int32(0))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		return
	} else {
		v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+15)))
		if base.Ui32(int32(253)) < base.Ui32((v21-int32(3))&int32(255)) {
			v28 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
			v30 = *(*int32)(unsafe.Add(mBase, _c_F_lock_twophase_recover[0]))
			v31 = F_get_hash_value(m, v30, l2)
			mBase = m.M
			v32 = m.ExcPending
			if v32 != 0 {
				return
			} else {
				v34 = *(*int32)(unsafe.Add(mBase, _c_F_lock_twophase_recover[1]))
				v36 = v31 & int32(15)
				v41 = v34 + v36<<(uint(int32(7))%32) + int32(_a_F_lock_twophase_recover_0)
				v43 = F_LWLockAcquire(m, v41, int32(0))
				mBase = m.M
				v44 = m.ExcPending
				if v44 != 0 {
					return
				} else {
					v49 = *(*int32)(unsafe.Add(mBase, uint32(v21<<(uint(int32(2))%32))+uint32(_c_F_lock_twophase_recover[2])))
					v51 = *(*int32)(unsafe.Add(mBase, _c_F_lock_twophase_recover[0]))
					v55 = F_hash_search_with_hash_value(m, v51, l2, v31, int32(3), v16+int32(71))
					mBase = m.M
					v56 = m.ExcPending
					if v56 != 0 {
						return
					} else {
						if v55 == int32(0) {
							F_LWLockRelease(m, v41)
							mBase = m.M
							v324 = m.ExcPending
							if v324 != 0 {
								return
							} else {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v328 = m.ExcPending
								if v328 != 0 {
									return
								} else {
									F_errcode(m, int32(_a_F_lock_twophase_recover_1))
									mBase = m.M
									v331 = m.ExcPending
									if v331 != 0 {
										return
									} else {
										F_errmsg(m, int32(_a_F_lock_twophase_recover_2), int32(0))
										mBase = m.M
										v335 = m.ExcPending
										if v335 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = int32(_a_F_lock_twophase_recover_3)
											F_errhint(m, int32(_a_F_lock_twophase_recover_4), v16+int32(16))
											mBase = m.M
											v342 = m.ExcPending
											if v342 != 0 {
												return
											} else {
												F_errfinish(m, int32(_a_F_lock_twophase_recover_5), int32(_a_F_lock_twophase_recover_6), int32(_a_F_lock_twophase_recover_7))
												mBase = m.M
												v347 = m.ExcPending
												if v347 != 0 {
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
							v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+71)))
							if v59 == int32(0) {
								v62 = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v55)+128)) = v62
								*(*int32)(unsafe.Add(mBase, uint32(v55)+84)) = v62
								*(*int32)(unsafe.Add(mBase, uint32(v55)+40)) = v62
								*(*int64)(unsafe.Add(mBase, uint32(v55)+16)) = int64(0)
								v71 = v55 + int32(32)
								*(*int32)(unsafe.Add(mBase, uint32(v55)+36)) = v71
								*(*int32)(unsafe.Add(mBase, uint32(v55)+32)) = v71
								v75 = v55 + int32(24)
								*(*int32)(unsafe.Add(mBase, uint32(v55)+28)) = v75
								*(*int32)(unsafe.Add(mBase, uint32(v55)+24)) = v75
								v78 = int32(40)
								v81 = v55 + int32(44)
								if v81&int32(3) == v62 {
									v87 = v55 + int32(84)
									v89 = v55 + int32(48)
									if base.Ui32(v89) < base.Ui32(v87) {
										v91 = v87
									} else {
										v91 = v89
									}
									v99 = (v91-v55-int32(45))&int32(-4) + int32(4)
								} else {
									v99 = v78
								}
								v103 = F__emscripten_memset_bulkmem(m, v81, base.I32_extend8_s(int32(0)), v99)
								mBase = m.M
								v105 = v55 + int32(88)
								if v105&int32(3) == int32(0) {
									v111 = v55 + int32(128)
									v113 = v55 + int32(92)
									if base.Ui32(v113) < base.Ui32(v111) {
										v115 = v111
									} else {
										v115 = v113
									}
									v123 = (v115-v55-int32(89))&int32(-4) + int32(4)
								} else {
									v123 = v78
								}
								v127 = F__emscripten_memset_bulkmem(m, v105, base.I32_extend8_s(int32(0)), v123)
								mBase = m.M
							} else {
							}
							*(*int32)(unsafe.Add(mBase, uint32(v16)+72)) = v55
							*(*int32)(unsafe.Add(mBase, uint32(v16)+76)) = v19
							v135 = *(*int32)(unsafe.Add(mBase, _c_F_lock_twophase_recover[3]))
							v144 = F_hash_search_with_hash_value(m, v135, v16+int32(72), v31^v19<<(uint(int32(4))%32), int32(3), v16+int32(71))
							mBase = m.M
							v145 = m.ExcPending
							if v145 != 0 {
								return
							} else {
								if v144 == int32(0) {
									v148 = *(*int32)(unsafe.Add(mBase, uint32(v55)+84))
									if v148 == int32(0) {
										v152 = *(*int32)(unsafe.Add(mBase, _c_F_lock_twophase_recover[0]))
										v155 = F_hash_search_with_hash_value(m, v152, v55, v31, int32(2), int32(0))
										mBase = m.M
										v156 = m.ExcPending
										if v156 != 0 {
											return
										} else {
											if v155 == int32(0) {
												F_errstart_cold(m, int32(23), int32(0))
												mBase = m.M
												v351 = m.ExcPending
												if v351 != 0 {
													return
												} else {
													F_errmsg_internal(m, int32(_a_F_lock_twophase_recover_8), int32(0))
													mBase = m.M
													v355 = m.ExcPending
													if v355 != 0 {
														return
													} else {
														F_errfinish(m, int32(_a_F_lock_twophase_recover_5), int32(_a_F_lock_twophase_recover_9), int32(_a_F_lock_twophase_recover_7))
														mBase = m.M
														v360 = m.ExcPending
														if v360 != 0 {
															return
														} else {
															base.Wasm_trap_unreachable()
															for {
															}
														}
													}
												}
											} else {
												F_LWLockRelease(m, v41)
												mBase = m.M
												v160 = m.ExcPending
												if v160 != 0 {
													return
												} else {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v164 = m.ExcPending
													if v164 != 0 {
														return
													} else {
														F_errcode(m, int32(_a_F_lock_twophase_recover_1))
														mBase = m.M
														v167 = m.ExcPending
														if v167 != 0 {
															return
														} else {
															F_errmsg(m, int32(_a_F_lock_twophase_recover_2), int32(0))
															mBase = m.M
															v171 = m.ExcPending
															if v171 != 0 {
																return
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v16)+32)) = int32(_a_F_lock_twophase_recover_3)
																F_errhint(m, int32(_a_F_lock_twophase_recover_4), v16+int32(32))
																mBase = m.M
																v178 = m.ExcPending
																if v178 != 0 {
																	return
																} else {
																	F_errfinish(m, int32(_a_F_lock_twophase_recover_5), int32(_a_F_lock_twophase_recover_10), int32(_a_F_lock_twophase_recover_7))
																	mBase = m.M
																	v183 = m.ExcPending
																	if v183 != 0 {
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
										F_LWLockRelease(m, v41)
										mBase = m.M
										v160 = m.ExcPending
										if v160 != 0 {
											return
										} else {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v164 = m.ExcPending
											if v164 != 0 {
												return
											} else {
												F_errcode(m, int32(_a_F_lock_twophase_recover_1))
												mBase = m.M
												v167 = m.ExcPending
												if v167 != 0 {
													return
												} else {
													F_errmsg(m, int32(_a_F_lock_twophase_recover_2), int32(0))
													mBase = m.M
													v171 = m.ExcPending
													if v171 != 0 {
														return
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v16)+32)) = int32(_a_F_lock_twophase_recover_3)
														F_errhint(m, int32(_a_F_lock_twophase_recover_4), v16+int32(32))
														mBase = m.M
														v178 = m.ExcPending
														if v178 != 0 {
															return
														} else {
															F_errfinish(m, int32(_a_F_lock_twophase_recover_5), int32(_a_F_lock_twophase_recover_10), int32(_a_F_lock_twophase_recover_7))
															mBase = m.M
															v183 = m.ExcPending
															if v183 != 0 {
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
								} else {
									v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+71)))
									if v184 == int32(0) {
										*(*int64)(unsafe.Add(mBase, uint32(v144)+12)) = int64(0)
										*(*int32)(unsafe.Add(mBase, uint32(v144)+8)) = v19
										v191 = v144 + int32(20)
										v193 = v55 + int32(24)
										v194 = *(*int32)(unsafe.Add(mBase, uint32(v55)+28))
										if v194 == int32(0) {
											*(*int32)(unsafe.Add(mBase, uint32(v55)+28)) = v193
											*(*int32)(unsafe.Add(mBase, uint32(v55)+24)) = v193
										} else {
										}
										*(*int32)(unsafe.Add(mBase, uint32(v144)+24)) = v193
										v200 = *(*int32)(unsafe.Add(mBase, uint32(v193)))
										*(*int32)(unsafe.Add(mBase, uint32(v144)+20)) = v200
										*(*int32)(unsafe.Add(mBase, uint32(v200)+4)) = v191
										*(*int32)(unsafe.Add(mBase, uint32(v193))) = v191
										v205 = v144 + int32(28)
										v208 = v19 + v36<<(uint(int32(3))%32)
										v210 = v208 + int32(148)
										v211 = *(*int32)(unsafe.Add(mBase, uint32(v208)+152))
										if v211 == int32(0) {
											*(*int32)(unsafe.Add(mBase, uint32(v210)+4)) = v210
											*(*int32)(unsafe.Add(mBase, uint32(v210))) = v210
										} else {
										}
										*(*int32)(unsafe.Add(mBase, uint32(v144)+32)) = v210
										v217 = *(*int32)(unsafe.Add(mBase, uint32(v210)))
										*(*int32)(unsafe.Add(mBase, uint32(v144)+28)) = v217
										*(*int32)(unsafe.Add(mBase, uint32(v217)+4)) = v205
										*(*int32)(unsafe.Add(mBase, uint32(v210))) = v205
									} else {
									}
									v226 = *(*int32)(unsafe.Add(mBase, uint32(v55)+84))
									v227 = int32(1)
									*(*int32)(unsafe.Add(mBase, uint32(v55)+84)) = v226 + v227
									v231 = v28 << (uint(int32(2)) % 32)
									v232 = v55 + v231
									v234 = v232 + int32(44)
									v235 = *(*int32)(unsafe.Add(mBase, uint32(v234)))
									*(*int32)(unsafe.Add(mBase, uint32(v234))) = v235 + v227
									v240 = v227 << (uint(v28) % 32)
									v241 = *(*int32)(unsafe.Add(mBase, uint32(v144)+12))
									if v240&v241 != 0 {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v364 = m.ExcPending
										if v364 != 0 {
											return
										} else {
											v365 = *(*int32)(unsafe.Add(mBase, uint32(v49)+8))
											v367 = *(*int32)(unsafe.Add(mBase, uint32(v365+v231)))
											v368 = *(*int64)(unsafe.Add(mBase, uint32(v55)))
											v369 = *(*int32)(unsafe.Add(mBase, uint32(v55)+8))
											*(*int32)(unsafe.Add(mBase, uint32(v16)+60)) = v369
											*(*int64)(unsafe.Add(mBase, uint32(v16)+52)) = v368
											*(*int32)(unsafe.Add(mBase, uint32(v16)+48)) = v367
											F_errmsg_internal(m, int32(_a_F_lock_twophase_recover_11), v16+int32(48))
											mBase = m.M
											v377 = m.ExcPending
											if v377 != 0 {
												return
											} else {
												F_errfinish(m, int32(_a_F_lock_twophase_recover_5), int32(_a_F_lock_twophase_recover_12), int32(_a_F_lock_twophase_recover_7))
												mBase = m.M
												v382 = m.ExcPending
												if v382 != 0 {
													return
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											}
										}
									} else {
										v243 = *(*int32)(unsafe.Add(mBase, uint32(v55)+128))
										v244 = int32(1)
										*(*int32)(unsafe.Add(mBase, uint32(v55)+128)) = v243 + v244
										v248 = v232 + int32(88)
										v249 = *(*int32)(unsafe.Add(mBase, uint32(v248)))
										*(*int32)(unsafe.Add(mBase, uint32(v248))) = v249 + v244
										v253 = *(*int32)(unsafe.Add(mBase, uint32(v55)+16))
										*(*int32)(unsafe.Add(mBase, uint32(v55)+16)) = v253 | v240
										v256 = *(*int32)(unsafe.Add(mBase, uint32(v248)))
										v257 = *(*int32)(unsafe.Add(mBase, uint32(v234)))
										if v256 == v257 {
											v259 = *(*int32)(unsafe.Add(mBase, uint32(v55)+20))
											*(*int32)(unsafe.Add(mBase, uint32(v55)+20)) = v259 & (v240 ^ int32(-1))
										} else {
										}
										v264 = *(*int32)(unsafe.Add(mBase, uint32(v144)+12))
										*(*int32)(unsafe.Add(mBase, uint32(v144)+12)) = v264 | v240
										v267 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+15)))
										if v267 != int32(1) {
											F_LWLockRelease(m, v41)
											mBase = m.M
											v306 = m.ExcPending
											if v306 != 0 {
												return
											} else {
												m.G0 = v16 + int32(80)
												return
											}
										} else {
											v270 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+14)))
											if v270 != 0 {
												F_LWLockRelease(m, v41)
												mBase = m.M
												v306 = m.ExcPending
												if v306 != 0 {
													return
												} else {
													m.G0 = v16 + int32(80)
													return
												}
											} else {
												v271 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
												if v271 == int32(0) {
													F_LWLockRelease(m, v41)
													mBase = m.M
													v306 = m.ExcPending
													if v306 != 0 {
														return
													} else {
														m.G0 = v16 + int32(80)
														return
													}
												} else {
													if v28 < int32(5) {
														F_LWLockRelease(m, v41)
														mBase = m.M
														v306 = m.ExcPending
														if v306 != 0 {
															return
														} else {
															m.G0 = v16 + int32(80)
															return
														}
													} else {
														v277 = *(*int32)(unsafe.Add(mBase, _c_F_lock_twophase_recover[4]))
														v278 = *(*int32)(unsafe.Add(mBase, uint32(v277)))
														*(*int32)(unsafe.Add(mBase, uint32(v277))) = int32(1)
														if v278 != 0 {
															v284 = *(*int32)(unsafe.Add(mBase, _c_F_lock_twophase_recover[4]))
															F_s_lock(m, v284, int32(_a_F_lock_twophase_recover_5), int32(_a_F_lock_twophase_recover_13), int32(_a_F_lock_twophase_recover_7))
															mBase = m.M
															v289 = m.ExcPending
															if v289 != 0 {
																return
															} else {
																v291 = *(*int32)(unsafe.Add(mBase, _c_F_lock_twophase_recover[4]))
																v296 = v291 + v31&int32(1023)<<(uint(int32(2))%32) + int32(4)
																v297 = *(*int32)(unsafe.Add(mBase, uint32(v296)))
																*(*int32)(unsafe.Add(mBase, uint32(v296))) = v297 + int32(1)
																*(*int32)(unsafe.Add(mBase, uint32(v291))) = int32(0)
																F_LWLockRelease(m, v41)
																mBase = m.M
																v306 = m.ExcPending
																if v306 != 0 {
																	return
																} else {
																	m.G0 = v16 + int32(80)
																	return
																}
															}
														} else {
															v291 = *(*int32)(unsafe.Add(mBase, _c_F_lock_twophase_recover[4]))
															v296 = v291 + v31&int32(1023)<<(uint(int32(2))%32) + int32(4)
															v297 = *(*int32)(unsafe.Add(mBase, uint32(v296)))
															*(*int32)(unsafe.Add(mBase, uint32(v296))) = v297 + int32(1)
															*(*int32)(unsafe.Add(mBase, uint32(v291))) = int32(0)
															F_LWLockRelease(m, v41)
															mBase = m.M
															v306 = m.ExcPending
															if v306 != 0 {
																return
															} else {
																m.G0 = v16 + int32(80)
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
					}
				}
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v313 = m.ExcPending
			if v313 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v16))) = v21
				F_errmsg_internal(m, int32(_a_F_lock_twophase_recover_14), v16)
				mBase = m.M
				v317 = m.ExcPending
				if v317 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_lock_twophase_recover_5), int32(_a_F_lock_twophase_recover_15), int32(_a_F_lock_twophase_recover_7))
					mBase = m.M
					v322 = m.ExcPending
					if v322 != 0 {
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
func F_longest(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
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
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v92 int32
	_ = v92
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v188 int32
	_ = v188
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v286 int32
	_ = v286
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v297 int32
	_ = v297
	var v302 int32
	_ = v302
	var v313 int32
	_ = v313
	var v320 int32
	_ = v320
	var v328 int32
	_ = v328
	var v339 int32
	_ = v339
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v353 int32
	_ = v353
	var v374 int32
	_ = v374
	var v377 int32
	_ = v377
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v383 int32
	_ = v383
	var v388 int32
	_ = v388
	var v391 int32
	_ = v391
	var v394 int32
	_ = v394
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v411 int32
	_ = v411
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v425 int32
	_ = v425
	var v428 int32
	_ = v428
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v440 int32
	_ = v440
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v447 int32
	_ = v447
	var v455 int32
	_ = v455
	var v457 int32
	_ = v457
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v481 int32
	_ = v481
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v491 int32
	_ = v491
	var v494 int32
	_ = v494
	var v499 int32
	_ = v499
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v505 int32
	_ = v505
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v521 int32
	_ = v521
	var v528 int32
	_ = v528
	var v533 int32
	_ = v533
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v545 int32
	_ = v545
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v563 int32
	_ = v563
	var v573 int32
	_ = v573
	var v582 int32
	_ = v582
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
	if l4 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = int32(0)
	goto L3
L2:
	;
	goto L3
L3:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+60))
	if v17 < int32(0) {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	return v582
L5:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116)+8)))
	if v117&int32(2) != 0 {
		goto L36
	} else {
		goto L37
	}
L6:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v23 = v20 + v17<<(uint(int32(3))%32)
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	if v24 < int32(0) {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v27 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+66)))
	v28 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+64)))
	if v28 <= v27 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	if l4 == int32(0) {
		v582 = v105
		goto L4
	} else {
		goto L33
	}
L9:
	;
	v31 = l2
	goto L11
L10:
	;
	v31 = int32(0)
	goto L11
L11:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	if v24 == v32 {
		v105 = v31
		goto L8
	} else {
		goto L12
	}
L12:
	;
	v38 = v32 - v24
	v39 = base.I32_div_u_s((l3-l2)>>(uint(int32(2))%32), v38)
	if base.Ui32(v39) < base.Ui32(v27) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v41 = v39
	goto L15
L14:
	;
	v41 = v27
	goto L15
L15:
	;
	if v27 == int32(256) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v44 = v39
	goto L18
L17:
	;
	v44 = v41
	goto L18
L18:
	;
	if base.Ui32(v44) < base.Ui32(v28) {
		v105 = int32(0)
		goto L8
	} else {
		goto L19
	}
L19:
	;
	if v44 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	if base.Ui32(v28) <= base.Ui32(v83) {
		goto L30
	} else {
		goto L31
	}
L21:
	;
	v80 = l2
	v83 = int32(0)
	goto L20
L22:
	;
	goto L23
L23:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v50 = int32(2)
	v58 = l2
	v61 = int32(0)
	goto L24
L24:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)+420))
	v70 = m.T0[v69].(func(*base.Module, int32, int32, int32) int32)(m, v49+v24<<(uint(v50)%32), v58, v38)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	v80 = v74
	v83 = v44
	goto L20
L26:
	;
	return int32(0)
L27:
	;
	if v70 != 0 {
		v80 = v58
		v83 = v61
		goto L20
	} else {
		goto L28
	}
L28:
	;
	v74 = v58 + v38<<(uint(v50)%32)
	v76 = v61 + int32(1)
	if v76 != v44 {
		v58 = v74
		v61 = v76
		goto L24
	} else {
		goto L29
	}
L29:
	;
	goto L25
L30:
	;
	v92 = v80
	goto L32
L31:
	;
	v92 = int32(0)
	goto L32
L32:
	;
	v105 = v92
	goto L8
L33:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if l3 != v108 {
		v582 = v105
		goto L4
	} else {
		goto L34
	}
L34:
	;
	if v108 != v105 {
		v582 = v105
		goto L4
	} else {
		goto L35
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = int32(1)
	return v105
L36:
	;
	v123 = (l3 - l2) >> (uint(int32(2)) % 32)
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v116)+40))
	if base.Ui32(v123) < base.Ui32(v124) {
		v582 = int32(0)
		goto L4
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	v152 = int32(0)
	v161 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v152 < v161 {
		goto L59
	} else {
		goto L60
	}
L39:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v116)+44))
	if v127 == int32(256) {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	if l4 == int32(0) {
		goto L43
	} else {
		goto L44
	}
L41:
	;
	goto L42
L42:
	;
	if v126 != l3 {
		goto L49
	} else {
		goto L50
	}
L43:
	;
	return l3
L44:
	;
	goto L45
L45:
	;
	if v126 != l3 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	return l3
L47:
	;
	goto L48
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = int32(1)
	return l3
L49:
	;
	if base.Ui32(v127) < base.Ui32(v123) {
		goto L53
	} else {
		goto L54
	}
L50:
	;
	if l4 == int32(0) {
		goto L49
	} else {
		goto L51
	}
L51:
	;
	if base.Ui32(v127+int32(1)) < base.Ui32(v123) {
		goto L49
	} else {
		goto L52
	}
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = int32(1)
	goto L49
L53:
	;
	v150 = l2 + v127<<(uint(int32(2))%32)
	goto L55
L54:
	;
	v150 = l3
	goto L55
L55:
	;
	return v150
L56:
	;
	if v374 == int32(0) {
		v582 = v152
		goto L4
	} else {
		goto L93
	}
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v353)+20)) = l2
	*(*int64)(unsafe.Add(mBase, uint32(l1)+48)) = int64(0)
	v374 = v353
	goto L56
L58:
	;
	v328 = int32(0)
	goto L90
L59:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v164)+8)))
	if v165&int32(1) != 0 {
		v320 = v164
		goto L58
	} else {
		goto L62
	}
L60:
	;
	goto L61
L61:
	;
	v169 = F_getvacant(m, l0, l1, l2, l2)
	mBase = m.M
	if v169 == int32(0) {
		goto L63
	} else {
		goto L64
	}
L62:
	;
	goto L61
L63:
	;
	v374 = int32(0)
	goto L56
L64:
	;
	goto L65
L65:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if int32(0) < v173 {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v177 = int32(0)
	goto L69
L67:
	;
	goto L68
L68:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v169)))
	v210 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v210)+12))
	v216 = v209 + int32(base.Ui32(v211)>>(uint(int32(3))%32))&int32(536870908)
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v216)))
	v218 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v216))) = v217 | v218<<(uint(v211)%32)
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v169)))
	v223 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v223 == v218 {
		goto L73
	} else {
		goto L74
	}
L69:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v169)))
	*(*int32)(unsafe.Add(mBase, uint32(v188+v177<<(uint(int32(2))%32)))) = int32(0)
	v195 = v177 + int32(1)
	v196 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v195 < v196 {
		v177 = v195
		goto L69
	} else {
		goto L71
	}
L70:
	;
	goto L68
L71:
	;
	goto L70
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v169)+8)) = int32(13)
	*(*int32)(unsafe.Add(mBase, uint32(v169)+4)) = v302
	v313 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v313 <= int32(0) {
		v353 = v169
		goto L57
	} else {
		goto L89
	}
L73:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v222)))
	v302 = v226
	goto L72
L74:
	;
	goto L75
L75:
	;
	if v223 <= int32(0) {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v302 = int32(0)
	goto L72
L77:
	;
	goto L78
L78:
	;
	v231 = v223 & int32(3)
	v232 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v223) {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v238 = v232
	v241 = v232
	v247 = v152
	goto L82
L80:
	;
	v265 = v232
	v268 = v232
	goto L81
L81:
	;
	if v231 == int32(0) {
		v302 = v268
		goto L72
	} else {
		goto L85
	}
L82:
	;
	v251 = v222 + v238<<(uint(int32(2))%32)
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v251)+12))
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v251)+8))
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v251)+4))
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v251)))
	v259 = v252 ^ (v253 ^ (v254 ^ (v255 ^ v241)))
	v260 = int32(4)
	v261 = v238 + v260
	v263 = v247 + v260
	if v263 != v223&int32(2147483644) {
		v238 = v261
		v241 = v259
		v247 = v263
		goto L82
	} else {
		goto L84
	}
L83:
	;
	v265 = v261
	v268 = v259
	goto L81
L84:
	;
	goto L83
L85:
	;
	v278 = v265
	v281 = v268
	v286 = v152
	goto L86
L86:
	;
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v222+v278<<(uint(int32(2))%32))))
	v293 = v292 ^ v281
	v294 = int32(1)
	v297 = v286 + v294
	if v297 != v231 {
		v278 = v278 + v294
		v281 = v293
		v286 = v297
		goto L86
	} else {
		goto L88
	}
L87:
	;
	v302 = v293
	goto L72
L88:
	;
	goto L87
L89:
	;
	v320 = v169
	goto L58
L90:
	;
	v339 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v339+v328<<(uint(int32(5))%32))+20)) = int32(0)
	v346 = v328 + int32(1)
	v347 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v346 < v347 {
		v328 = v346
		goto L90
	} else {
		goto L92
	}
L91:
	;
	v353 = v320
	goto L57
L92:
	;
	goto L91
L93:
	;
	v377 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v377 == l2 {
		goto L95
	} else {
		goto L96
	}
L94:
	;
	v403 = F_miss(m, l0, l1, v374, base.I32_extend16_s(v401), l2, l2)
	mBase = m.M
	v404 = m.ExcPending
	if v404 != 0 {
		goto L26
	} else {
		goto L101
	}
L95:
	;
	v379 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	v380 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v383 = int32(1)
	v388 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v379+(v380^int32(-1))&v383<<(uint(v383)%32))+20)))
	v401 = v388
	goto L94
L96:
	;
	goto L97
L97:
	;
	v391 = *(*int32)(unsafe.Add(mBase, uint32(l2-int32(4))))
	if base.Ui32(v391) <= base.Ui32(int32(2047)) {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v394 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
	v398 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v394+v391<<(uint(int32(1))%32)))))
	v401 = v398
	goto L94
L99:
	;
	goto L100
L100:
	;
	v399 = F_pg_reg_getcolor(m, v14, v391)
	mBase = m.M
	v401 = v399
	goto L94
L101:
	;
	if v403 == int32(0) {
		v582 = v152
		goto L4
	} else {
		goto L102
	}
L102:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v403)+20)) = l2
	v411 = l3 + base.B2i32(l3 != v13)<<(uint(int32(2))%32)
	if base.Ui32(v411) <= base.Ui32(l2) {
		goto L104
	} else {
		goto L105
	}
L103:
	;
	v462 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v462 != 0 {
		v582 = v152
		goto L4
	} else {
		goto L118
	}
L104:
	;
	v455 = l2
	v457 = v403
	goto L103
L105:
	;
	goto L106
L106:
	;
	v418 = l2
	v419 = v403
	goto L107
L107:
	;
	v425 = *(*int32)(unsafe.Add(mBase, uint32(v418)))
	if base.Ui32(v425) <= base.Ui32(int32(2047)) {
		goto L110
	} else {
		goto L111
	}
L108:
	;
	v455 = v447
	v457 = v445
	goto L103
L109:
	;
	v435 = *(*int32)(unsafe.Add(mBase, uint32(v419)+24))
	v436 = base.I32_extend16_s(v434)
	v440 = *(*int32)(unsafe.Add(mBase, uint32(v435+v436<<(uint(int32(2))%32))))
	if v440 != 0 {
		v445 = v440
		goto L113
	} else {
		goto L114
	}
L110:
	;
	v428 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
	v432 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v428+v425<<(uint(int32(1))%32)))))
	v434 = v432
	goto L109
L111:
	;
	goto L112
L112:
	;
	v433 = F_pg_reg_getcolor(m, v14, v425)
	mBase = m.M
	v434 = v433
	goto L109
L113:
	;
	v447 = v418 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v445)+20)) = v447
	if base.Ui32(v447) < base.Ui32(v411) {
		v418 = v447
		v419 = v445
		goto L107
	} else {
		goto L117
	}
L114:
	;
	v443 = F_miss(m, l0, l1, v419, v436, v418+int32(4), l2)
	mBase = m.M
	v444 = m.ExcPending
	if v444 != 0 {
		goto L26
	} else {
		goto L115
	}
L115:
	;
	if v443 != 0 {
		v445 = v443
		goto L113
	} else {
		goto L116
	}
L116:
	;
	v455 = v418
	v457 = v419
	goto L103
L117:
	;
	goto L108
L118:
	;
	v463 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v455 != v463 {
		goto L119
	} else {
		goto L120
	}
L119:
	;
	v487 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v488 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v488 <= int32(0) {
		v563 = v487
		goto L131
	} else {
		goto L132
	}
L120:
	;
	if l3 != v463 {
		goto L119
	} else {
		goto L121
	}
L121:
	;
	if l4 != 0 {
		goto L122
	} else {
		goto L123
	}
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = int32(1)
	goto L124
L123:
	;
	goto L124
L124:
	;
	v468 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	v469 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v475 = int32(*(*int16)(unsafe.Add(mBase, uint32(v468+(v469^int32(-1))&int32(2))+24)))
	v476 = F_miss(m, l0, l1, v457, v475, v455, l2)
	mBase = m.M
	v477 = m.ExcPending
	if v477 != 0 {
		goto L26
	} else {
		goto L125
	}
L125:
	;
	v478 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v478 != 0 {
		v582 = v152
		goto L4
	} else {
		goto L126
	}
L126:
	;
	if v476 == int32(0) {
		goto L119
	} else {
		goto L127
	}
L127:
	;
	v481 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v476)+8)))
	if v481&int32(2) != 0 {
		goto L128
	} else {
		goto L129
	}
L128:
	;
	return v455
L129:
	;
	goto L130
L130:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v476)+20)) = v455
	goto L119
L131:
	;
	if v563 != 0 {
		goto L167
	} else {
		goto L168
	}
L132:
	;
	v491 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v488&int32(1) != 0 {
		goto L133
	} else {
		goto L134
	}
L133:
	;
	v494 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v491)+8)))
	if v494&int32(2) == int32(0) {
		v505 = v487
		goto L136
	} else {
		goto L137
	}
L134:
	;
	v511 = v491
	v512 = v487
	v513 = v488
	goto L135
L135:
	;
	if v488 == int32(1) {
		v563 = v512
		goto L131
	} else {
		goto L145
	}
L136:
	;
	v511 = v491 + int32(32)
	v512 = v505
	v513 = v488 - int32(1)
	goto L135
L137:
	;
	v499 = *(*int32)(unsafe.Add(mBase, uint32(v491)+20))
	if v487 == v499 {
		v505 = v487
		goto L136
	} else {
		goto L138
	}
L138:
	;
	if base.Ui32(v487) < base.Ui32(v499) {
		goto L139
	} else {
		goto L140
	}
L139:
	;
	v502 = v499
	goto L141
L140:
	;
	v502 = v487
	goto L141
L141:
	;
	if v487 != 0 {
		goto L142
	} else {
		goto L143
	}
L142:
	;
	v503 = v502
	goto L144
L143:
	;
	v503 = v499
	goto L144
L144:
	;
	v505 = v503
	goto L136
L145:
	;
	v517 = v513
	v518 = v511
	v521 = v512
	goto L146
L146:
	;
	v528 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v518)+8)))
	if v528&int32(2) == int32(0) {
		v539 = v521
		goto L148
	} else {
		goto L149
	}
L147:
	;
	v563 = v551
	goto L131
L148:
	;
	v540 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v518)+40)))
	if v540&int32(2) == int32(0) {
		v551 = v539
		goto L157
	} else {
		goto L158
	}
L149:
	;
	v533 = *(*int32)(unsafe.Add(mBase, uint32(v518)+20))
	if v521 == v533 {
		v539 = v521
		goto L148
	} else {
		goto L150
	}
L150:
	;
	if base.Ui32(v521) < base.Ui32(v533) {
		goto L151
	} else {
		goto L152
	}
L151:
	;
	v536 = v533
	goto L153
L152:
	;
	v536 = v521
	goto L153
L153:
	;
	if v521 != 0 {
		goto L154
	} else {
		goto L155
	}
L154:
	;
	v537 = v536
	goto L156
L155:
	;
	v537 = v533
	goto L156
L156:
	;
	v539 = v537
	goto L148
L157:
	;
	v552 = int32(2)
	if v552 < v517 {
		v517 = v517 - v552
		v518 = v518 - int32(-64)
		v521 = v551
		goto L146
	} else {
		goto L166
	}
L158:
	;
	v545 = *(*int32)(unsafe.Add(mBase, uint32(v518)+52))
	if v539 == v545 {
		v551 = v539
		goto L157
	} else {
		goto L159
	}
L159:
	;
	if base.Ui32(v539) < base.Ui32(v545) {
		goto L160
	} else {
		goto L161
	}
L160:
	;
	v548 = v545
	goto L162
L161:
	;
	v548 = v539
	goto L162
L162:
	;
	if v539 != 0 {
		goto L163
	} else {
		goto L164
	}
L163:
	;
	v549 = v548
	goto L165
L164:
	;
	v549 = v545
	goto L165
L165:
	;
	v551 = v549
	goto L157
L166:
	;
	goto L147
L167:
	;
	v573 = v563 - int32(4)
	goto L169
L168:
	;
	v573 = int32(0)
	goto L169
L169:
	;
	v582 = v573
	goto L4
}
func F_lookup_proof_cache(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
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
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v118 int32
	_ = v118
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v190 int32
	_ = v190
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v240 int32
	_ = v240
	v4 = int32(0)
	v18 = m.G0
	v20 = v18 + int32(-64)
	m.G0 = v20
	v23 = *(*int32)(unsafe.Add(mBase, _c_F_lookup_proof_cache[0]))
	if v23 == v4 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v20)+28)) = int64(85899345928)
	v34 = F_hash_create(m, int32(_a_F_lookup_proof_cache_0), int32(256), v18+int32(-52), int32(40))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v46 = v23
	goto L3
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v20)+12)) = l0
	v54 = F_hash_search(m, v46, v18+int32(-52), int32(1), v18+int32(-1))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L4
	} else {
		goto L7
	}
L4:
	;
	return int32(0)
L5:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_lookup_proof_cache[0])) = v34
	F_CacheRegisterSyscacheCallback(m, int32(3), int32(891), int32(0))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v45 = *(*int32)(unsafe.Add(mBase, _c_F_lookup_proof_cache[0]))
	v46 = v45
	goto L3
L7:
	;
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+63)))
	if v56 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	m.G0 = v20 - int32(-64)
	return v54
L9:
	;
	v65 = F_get_op_index_interpretation(m, l1)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L4
	} else {
		goto L21
	}
L10:
	;
	v59 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v54)+8)) = uint16(v59)
	goto L9
L11:
	;
	goto L12
L12:
	;
	if l2 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54)+9)))
	if v61 == int32(0) {
		goto L9
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54)+8)))
	if v64 != 0 {
		goto L8
	} else {
		goto L17
	}
L16:
	;
	goto L8
L17:
	;
	goto L9
L18:
	;
	F_list_free_deep(m, v213)
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L4
	} else {
		goto L53
	}
L19:
	;
	v209 = v4
	v213 = v205
	v215 = v4
	goto L18
L20:
	;
	v205 = int32(0)
	goto L19
L21:
	;
	if v65 == int32(0) {
		goto L20
	} else {
		goto L22
	}
L22:
	;
	v69 = F_get_op_index_interpretation(m, l0)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L4
	} else {
		goto L23
	}
L23:
	;
	if v69 == int32(0) {
		goto L20
	} else {
		goto L24
	}
L24:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v69)+4))
	if v73 <= int32(0) {
		v205 = v69
		goto L19
	} else {
		goto L25
	}
L25:
	;
	if l2 != 0 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v78 = int32(_a_F_lookup_proof_cache_1)
	goto L28
L27:
	;
	v78 = int32(_a_F_lookup_proof_cache_2)
	goto L28
L28:
	;
	if l2 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v81 = int32(_a_F_lookup_proof_cache_3)
	goto L31
L30:
	;
	v81 = int32(_a_F_lookup_proof_cache_4)
	goto L31
L31:
	;
	v91 = v4
	v92 = v4
	goto L32
L32:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v65)+4))
	if int32(0) < v99 {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	v209 = int32(0)
	v213 = v69
	v215 = v190
	goto L18
L34:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v69)+12))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v102+v92<<(uint(int32(2))%32))))
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v106)))
	v109 = int32(0)
	v118 = v91
	goto L37
L35:
	;
	v190 = v91
	goto L36
L36:
	;
	v199 = v92 + int32(1)
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v69)+4))
	if v199 < v200 {
		v91 = v190
		v92 = v199
		goto L32
	} else {
		goto L52
	}
L37:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v65)+12))
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v126+v109<<(uint(int32(2))%32))))
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v130)))
	if v107 != v131 {
		v174 = v118
		goto L39
	} else {
		goto L40
	}
L38:
	;
	v190 = v174
	goto L36
L39:
	;
	v178 = v109 + int32(1)
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v65)+4))
	if v178 < v179 {
		v109 = v178
		v118 = v174
		goto L37
	} else {
		goto L51
	}
L40:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v106)+4))
	v134 = int32(1)
	v135 = v133 - v134
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v130)+4))
	v138 = v136 - v134
	v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135+(v81+v138*int32(6))))))
	v144 = v118 | v143
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v78+v138*int32(24)+v135<<(uint(int32(2))%32))))
	switch v151 {
	case 0:
		v174 = v144
		goto L39
	default:
		goto L42
	case 6:
		goto L43
	}
L41:
	;
	if v166 == int32(0) {
		v174 = v144
		goto L39
	} else {
		goto L48
	}
L42:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v106)+12))
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v130)+12))
	v163 = F_get_opfamily_member_for_cmptype(m, v107, v161, v162, v151)
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L4
	} else {
		goto L47
	}
L43:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v106)+12))
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v130)+12))
	v155 = F_get_opfamily_member_for_cmptype(m, v107, v152, v153, int32(3))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L4
	} else {
		goto L44
	}
L44:
	;
	if v155 == int32(0) {
		v174 = v144
		goto L39
	} else {
		goto L45
	}
L45:
	;
	v159 = F_get_negator(m, v155)
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L4
	} else {
		goto L46
	}
L46:
	;
	v166 = v159
	goto L41
L47:
	;
	v166 = v163
	goto L41
L48:
	;
	v169 = F_op_volatile(m, v166)
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L4
	} else {
		goto L49
	}
L49:
	;
	if v169 == int32(105) {
		v209 = v166
		v213 = v69
		v215 = v144
		goto L18
	} else {
		goto L50
	}
L50:
	;
	v174 = v144
	goto L39
L51:
	;
	goto L38
L52:
	;
	goto L33
L53:
	;
	F_list_free_deep(m, v65)
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L4
	} else {
		goto L54
	}
L54:
	;
	if v215&int32(1) != 0 {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v229 = F_op_volatile(m, l1)
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L4
	} else {
		goto L58
	}
L56:
	;
	v233 = v4
	goto L57
L57:
	;
	if l2 != 0 {
		goto L59
	} else {
		goto L60
	}
L58:
	;
	v233 = base.B2i32(v229 == int32(105))
	goto L57
L59:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v54)+11)) = uint8(v233)
	*(*int32)(unsafe.Add(mBase, uint32(v54)+16)) = v209
	v236 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v54)+9)) = uint8(v236)
	goto L8
L60:
	;
	goto L61
L61:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v54)+10)) = uint8(v233)
	*(*int32)(unsafe.Add(mBase, uint32(v54)+12)) = v209
	v240 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v54)+8)) = uint8(v240)
	goto L8
}
func F_lookup_rowtype_tupdesc_internal(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
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
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v118 int64
	_ = v118
	var v120 int64
	_ = v120
	var v122 int32
	_ = v122
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v166 int32
	_ = v166
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = l1
	if l0 != int32(2249) {
		v16 = F_lookup_type_cache(m, l0, int32(256))
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int32(0)
		} else {
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v16)+188))
			if v20 != 0 {
				v140 = v20
				m.G0 = v10 + int32(16)
				return v140
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(151027844))
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return int32(0)
					} else {
						v28 = F_format_type_be(m, l0)
						mBase = m.M
						v29 = m.ExcPending
						if v29 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v10))) = v28
							F_errmsg(m, int32(_a_F_lookup_rowtype_tupdesc_internal_0), v10)
							mBase = m.M
							v33 = m.ExcPending
							if v33 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_lookup_rowtype_tupdesc_internal_1), int32(1842), int32(_a_F_lookup_rowtype_tupdesc_internal_2))
								mBase = m.M
								v38 = m.ExcPending
								if v38 != 0 {
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
		if l1 < int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v154 = m.ExcPending
			if v154 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(151027844))
				mBase = m.M
				v157 = m.ExcPending
				if v157 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(_a_F_lookup_rowtype_tupdesc_internal_3), int32(0))
					mBase = m.M
					v161 = m.ExcPending
					if v161 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_lookup_rowtype_tupdesc_internal_1), int32(1901), int32(_a_F_lookup_rowtype_tupdesc_internal_2))
						mBase = m.M
						v166 = m.ExcPending
						if v166 != 0 {
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
			v42 = *(*int32)(unsafe.Add(mBase, _c_F_lookup_rowtype_tupdesc_internal[0]))
			if l1 < v42 {
				v45 = *(*int32)(unsafe.Add(mBase, _c_F_lookup_rowtype_tupdesc_internal[1]))
				v49 = *(*int32)(unsafe.Add(mBase, uint32(v45+l1<<(uint(int32(4))%32))+8))
				if v49 != 0 {
					v140 = v49
					m.G0 = v10 + int32(16)
					return v140
				} else {
					v52 = *(*int32)(unsafe.Add(mBase, _c_F_lookup_rowtype_tupdesc_internal[2]))
					v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)+8))
					if v53 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v154 = m.ExcPending
						if v154 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(151027844))
							mBase = m.M
							v157 = m.ExcPending
							if v157 != 0 {
								return int32(0)
							} else {
								F_errmsg(m, int32(_a_F_lookup_rowtype_tupdesc_internal_3), int32(0))
								mBase = m.M
								v161 = m.ExcPending
								if v161 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_lookup_rowtype_tupdesc_internal_1), int32(1901), int32(_a_F_lookup_rowtype_tupdesc_internal_2))
									mBase = m.M
									v166 = m.ExcPending
									if v166 != 0 {
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
						v56 = *(*int32)(unsafe.Add(mBase, uint32(v52)+16))
						v60 = F_dshash_find(m, v56, v10+int32(12), int32(0))
						mBase = m.M
						v61 = m.ExcPending
						if v61 != 0 {
							return int32(0)
						} else {
							if v60 == int32(0) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v154 = m.ExcPending
								if v154 != 0 {
									return int32(0)
								} else {
									F_errcode(m, int32(151027844))
									mBase = m.M
									v157 = m.ExcPending
									if v157 != 0 {
										return int32(0)
									} else {
										F_errmsg(m, int32(_a_F_lookup_rowtype_tupdesc_internal_3), int32(0))
										mBase = m.M
										v161 = m.ExcPending
										if v161 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_lookup_rowtype_tupdesc_internal_1), int32(1901), int32(_a_F_lookup_rowtype_tupdesc_internal_2))
											mBase = m.M
											v166 = m.ExcPending
											if v166 != 0 {
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
								v65 = *(*int32)(unsafe.Add(mBase, _c_F_lookup_rowtype_tupdesc_internal[2]))
								v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)+4))
								v67 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
								v68 = F_dsa_get_address(m, v66, v67)
								mBase = m.M
								v69 = m.ExcPending
								if v69 != 0 {
									return int32(0)
								} else {
									v70 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
									v72 = *(*int32)(unsafe.Add(mBase, _c_F_lookup_rowtype_tupdesc_internal[1]))
									if v72 != 0 {
										v74 = *(*int32)(unsafe.Add(mBase, _c_F_lookup_rowtype_tupdesc_internal[0]))
										v86 = v72
										v87 = v74
										if v87 <= v70 {
											v91 = int32(1)
											v94 = v70 + v91
											if v94&v70 != 0 {
												v99 = v91 << (uint(int32(32)-base.I32_clz(v94)) % 32)
											} else {
												v99 = v94
											}
											v102 = F_repalloc0(m, v86, v87<<(uint(int32(4))%32), v99<<(uint(int32(4))%32))
											mBase = m.M
											v103 = m.ExcPending
											if v103 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, _c_F_lookup_rowtype_tupdesc_internal[0])) = v99
												*(*int32)(unsafe.Add(mBase, _c_F_lookup_rowtype_tupdesc_internal[1])) = v102
												v108 = v102
												v111 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
												v112 = int32(4)
												*(*int32)(unsafe.Add(mBase, uint32(v108+v111<<(uint(v112)%32))+8)) = v68
												v116 = int32(_a_F_lookup_rowtype_tupdesc_internal_4)
												v118 = *(*int64)(unsafe.Add(mBase, _c_F_lookup_rowtype_tupdesc_internal[3]))
												v120 = v118 + int64(1)
												*(*int64)(unsafe.Add(mBase, _c_F_lookup_rowtype_tupdesc_internal[3])) = v120
												v122 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
												*(*int64)(unsafe.Add(mBase, uint32(v108+v122<<(uint(v112)%32)))) = v120
												v128 = *(*int32)(unsafe.Add(mBase, _c_F_lookup_rowtype_tupdesc_internal[2]))
												v129 = *(*int32)(unsafe.Add(mBase, uint32(v128)+16))
												F_dshash_release_lock(m, v129, v60)
												mBase = m.M
												v131 = m.ExcPending
												if v131 != 0 {
													return int32(0)
												} else {
													v133 = *(*int32)(unsafe.Add(mBase, _c_F_lookup_rowtype_tupdesc_internal[1]))
													v134 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
													v138 = *(*int32)(unsafe.Add(mBase, uint32(v133+v134<<(uint(int32(4))%32))+8))
													v140 = v138
													m.G0 = v10 + int32(16)
													return v140
												}
											}
										} else {
											v108 = v86
											v111 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
											v112 = int32(4)
											*(*int32)(unsafe.Add(mBase, uint32(v108+v111<<(uint(v112)%32))+8)) = v68
											v116 = int32(_a_F_lookup_rowtype_tupdesc_internal_4)
											v118 = *(*int64)(unsafe.Add(mBase, _c_F_lookup_rowtype_tupdesc_internal[3]))
											v120 = v118 + int64(1)
											*(*int64)(unsafe.Add(mBase, _c_F_lookup_rowtype_tupdesc_internal[3])) = v120
											v122 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
											*(*int64)(unsafe.Add(mBase, uint32(v108+v122<<(uint(v112)%32)))) = v120
											v128 = *(*int32)(unsafe.Add(mBase, _c_F_lookup_rowtype_tupdesc_internal[2]))
											v129 = *(*int32)(unsafe.Add(mBase, uint32(v128)+16))
											F_dshash_release_lock(m, v129, v60)
											mBase = m.M
											v131 = m.ExcPending
											if v131 != 0 {
												return int32(0)
											} else {
												v133 = *(*int32)(unsafe.Add(mBase, _c_F_lookup_rowtype_tupdesc_internal[1]))
												v134 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
												v138 = *(*int32)(unsafe.Add(mBase, uint32(v133+v134<<(uint(int32(4))%32))+8))
												v140 = v138
												m.G0 = v10 + int32(16)
												return v140
											}
										}
									} else {
										v77 = *(*int32)(unsafe.Add(mBase, _c_F_lookup_rowtype_tupdesc_internal[4]))
										v79 = F_MemoryContextAllocZero(m, v77, int32(1024))
										mBase = m.M
										v80 = m.ExcPending
										if v80 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, _c_F_lookup_rowtype_tupdesc_internal[0])) = int32(64)
											*(*int32)(unsafe.Add(mBase, _c_F_lookup_rowtype_tupdesc_internal[1])) = v79
											v86 = v79
											v87 = int32(64)
											if v87 <= v70 {
												v91 = int32(1)
												v94 = v70 + v91
												if v94&v70 != 0 {
													v99 = v91 << (uint(int32(32)-base.I32_clz(v94)) % 32)
												} else {
													v99 = v94
												}
												v102 = F_repalloc0(m, v86, v87<<(uint(int32(4))%32), v99<<(uint(int32(4))%32))
												mBase = m.M
												v103 = m.ExcPending
												if v103 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, _c_F_lookup_rowtype_tupdesc_internal[0])) = v99
													*(*int32)(unsafe.Add(mBase, _c_F_lookup_rowtype_tupdesc_internal[1])) = v102
													v108 = v102
													v111 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
													v112 = int32(4)
													*(*int32)(unsafe.Add(mBase, uint32(v108+v111<<(uint(v112)%32))+8)) = v68
													v116 = int32(_a_F_lookup_rowtype_tupdesc_internal_4)
													v118 = *(*int64)(unsafe.Add(mBase, _c_F_lookup_rowtype_tupdesc_internal[3]))
													v120 = v118 + int64(1)
													*(*int64)(unsafe.Add(mBase, _c_F_lookup_rowtype_tupdesc_internal[3])) = v120
													v122 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
													*(*int64)(unsafe.Add(mBase, uint32(v108+v122<<(uint(v112)%32)))) = v120
													v128 = *(*int32)(unsafe.Add(mBase, _c_F_lookup_rowtype_tupdesc_internal[2]))
													v129 = *(*int32)(unsafe.Add(mBase, uint32(v128)+16))
													F_dshash_release_lock(m, v129, v60)
													mBase = m.M
													v131 = m.ExcPending
													if v131 != 0 {
														return int32(0)
													} else {
														v133 = *(*int32)(unsafe.Add(mBase, _c_F_lookup_rowtype_tupdesc_internal[1]))
														v134 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
														v138 = *(*int32)(unsafe.Add(mBase, uint32(v133+v134<<(uint(int32(4))%32))+8))
														v140 = v138
														m.G0 = v10 + int32(16)
														return v140
													}
												}
											} else {
												v108 = v86
												v111 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
												v112 = int32(4)
												*(*int32)(unsafe.Add(mBase, uint32(v108+v111<<(uint(v112)%32))+8)) = v68
												v116 = int32(_a_F_lookup_rowtype_tupdesc_internal_4)
												v118 = *(*int64)(unsafe.Add(mBase, _c_F_lookup_rowtype_tupdesc_internal[3]))
												v120 = v118 + int64(1)
												*(*int64)(unsafe.Add(mBase, _c_F_lookup_rowtype_tupdesc_internal[3])) = v120
												v122 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
												*(*int64)(unsafe.Add(mBase, uint32(v108+v122<<(uint(v112)%32)))) = v120
												v128 = *(*int32)(unsafe.Add(mBase, _c_F_lookup_rowtype_tupdesc_internal[2]))
												v129 = *(*int32)(unsafe.Add(mBase, uint32(v128)+16))
												F_dshash_release_lock(m, v129, v60)
												mBase = m.M
												v131 = m.ExcPending
												if v131 != 0 {
													return int32(0)
												} else {
													v133 = *(*int32)(unsafe.Add(mBase, _c_F_lookup_rowtype_tupdesc_internal[1]))
													v134 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
													v138 = *(*int32)(unsafe.Add(mBase, uint32(v133+v134<<(uint(int32(4))%32))+8))
													v140 = v138
													m.G0 = v10 + int32(16)
													return v140
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
				v52 = *(*int32)(unsafe.Add(mBase, _c_F_lookup_rowtype_tupdesc_internal[2]))
				v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)+8))
				if v53 == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v154 = m.ExcPending
					if v154 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(151027844))
						mBase = m.M
						v157 = m.ExcPending
						if v157 != 0 {
							return int32(0)
						} else {
							F_errmsg(m, int32(_a_F_lookup_rowtype_tupdesc_internal_3), int32(0))
							mBase = m.M
							v161 = m.ExcPending
							if v161 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_lookup_rowtype_tupdesc_internal_1), int32(1901), int32(_a_F_lookup_rowtype_tupdesc_internal_2))
								mBase = m.M
								v166 = m.ExcPending
								if v166 != 0 {
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
					v56 = *(*int32)(unsafe.Add(mBase, uint32(v52)+16))
					v60 = F_dshash_find(m, v56, v10+int32(12), int32(0))
					mBase = m.M
					v61 = m.ExcPending
					if v61 != 0 {
						return int32(0)
					} else {
						if v60 == int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v154 = m.ExcPending
							if v154 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(151027844))
								mBase = m.M
								v157 = m.ExcPending
								if v157 != 0 {
									return int32(0)
								} else {
									F_errmsg(m, int32(_a_F_lookup_rowtype_tupdesc_internal_3), int32(0))
									mBase = m.M
									v161 = m.ExcPending
									if v161 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_lookup_rowtype_tupdesc_internal_1), int32(1901), int32(_a_F_lookup_rowtype_tupdesc_internal_2))
										mBase = m.M
										v166 = m.ExcPending
										if v166 != 0 {
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
							v65 = *(*int32)(unsafe.Add(mBase, _c_F_lookup_rowtype_tupdesc_internal[2]))
							v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)+4))
							v67 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
							v68 = F_dsa_get_address(m, v66, v67)
							mBase = m.M
							v69 = m.ExcPending
							if v69 != 0 {
								return int32(0)
							} else {
								v70 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
								v72 = *(*int32)(unsafe.Add(mBase, _c_F_lookup_rowtype_tupdesc_internal[1]))
								if v72 != 0 {
									v74 = *(*int32)(unsafe.Add(mBase, _c_F_lookup_rowtype_tupdesc_internal[0]))
									v86 = v72
									v87 = v74
									if v87 <= v70 {
										v91 = int32(1)
										v94 = v70 + v91
										if v94&v70 != 0 {
											v99 = v91 << (uint(int32(32)-base.I32_clz(v94)) % 32)
										} else {
											v99 = v94
										}
										v102 = F_repalloc0(m, v86, v87<<(uint(int32(4))%32), v99<<(uint(int32(4))%32))
										mBase = m.M
										v103 = m.ExcPending
										if v103 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, _c_F_lookup_rowtype_tupdesc_internal[0])) = v99
											*(*int32)(unsafe.Add(mBase, _c_F_lookup_rowtype_tupdesc_internal[1])) = v102
											v108 = v102
											v111 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
											v112 = int32(4)
											*(*int32)(unsafe.Add(mBase, uint32(v108+v111<<(uint(v112)%32))+8)) = v68
											v116 = int32(_a_F_lookup_rowtype_tupdesc_internal_4)
											v118 = *(*int64)(unsafe.Add(mBase, _c_F_lookup_rowtype_tupdesc_internal[3]))
											v120 = v118 + int64(1)
											*(*int64)(unsafe.Add(mBase, _c_F_lookup_rowtype_tupdesc_internal[3])) = v120
											v122 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
											*(*int64)(unsafe.Add(mBase, uint32(v108+v122<<(uint(v112)%32)))) = v120
											v128 = *(*int32)(unsafe.Add(mBase, _c_F_lookup_rowtype_tupdesc_internal[2]))
											v129 = *(*int32)(unsafe.Add(mBase, uint32(v128)+16))
											F_dshash_release_lock(m, v129, v60)
											mBase = m.M
											v131 = m.ExcPending
											if v131 != 0 {
												return int32(0)
											} else {
												v133 = *(*int32)(unsafe.Add(mBase, _c_F_lookup_rowtype_tupdesc_internal[1]))
												v134 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
												v138 = *(*int32)(unsafe.Add(mBase, uint32(v133+v134<<(uint(int32(4))%32))+8))
												v140 = v138
												m.G0 = v10 + int32(16)
												return v140
											}
										}
									} else {
										v108 = v86
										v111 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
										v112 = int32(4)
										*(*int32)(unsafe.Add(mBase, uint32(v108+v111<<(uint(v112)%32))+8)) = v68
										v116 = int32(_a_F_lookup_rowtype_tupdesc_internal_4)
										v118 = *(*int64)(unsafe.Add(mBase, _c_F_lookup_rowtype_tupdesc_internal[3]))
										v120 = v118 + int64(1)
										*(*int64)(unsafe.Add(mBase, _c_F_lookup_rowtype_tupdesc_internal[3])) = v120
										v122 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
										*(*int64)(unsafe.Add(mBase, uint32(v108+v122<<(uint(v112)%32)))) = v120
										v128 = *(*int32)(unsafe.Add(mBase, _c_F_lookup_rowtype_tupdesc_internal[2]))
										v129 = *(*int32)(unsafe.Add(mBase, uint32(v128)+16))
										F_dshash_release_lock(m, v129, v60)
										mBase = m.M
										v131 = m.ExcPending
										if v131 != 0 {
											return int32(0)
										} else {
											v133 = *(*int32)(unsafe.Add(mBase, _c_F_lookup_rowtype_tupdesc_internal[1]))
											v134 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
											v138 = *(*int32)(unsafe.Add(mBase, uint32(v133+v134<<(uint(int32(4))%32))+8))
											v140 = v138
											m.G0 = v10 + int32(16)
											return v140
										}
									}
								} else {
									v77 = *(*int32)(unsafe.Add(mBase, _c_F_lookup_rowtype_tupdesc_internal[4]))
									v79 = F_MemoryContextAllocZero(m, v77, int32(1024))
									mBase = m.M
									v80 = m.ExcPending
									if v80 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, _c_F_lookup_rowtype_tupdesc_internal[0])) = int32(64)
										*(*int32)(unsafe.Add(mBase, _c_F_lookup_rowtype_tupdesc_internal[1])) = v79
										v86 = v79
										v87 = int32(64)
										if v87 <= v70 {
											v91 = int32(1)
											v94 = v70 + v91
											if v94&v70 != 0 {
												v99 = v91 << (uint(int32(32)-base.I32_clz(v94)) % 32)
											} else {
												v99 = v94
											}
											v102 = F_repalloc0(m, v86, v87<<(uint(int32(4))%32), v99<<(uint(int32(4))%32))
											mBase = m.M
											v103 = m.ExcPending
											if v103 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, _c_F_lookup_rowtype_tupdesc_internal[0])) = v99
												*(*int32)(unsafe.Add(mBase, _c_F_lookup_rowtype_tupdesc_internal[1])) = v102
												v108 = v102
												v111 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
												v112 = int32(4)
												*(*int32)(unsafe.Add(mBase, uint32(v108+v111<<(uint(v112)%32))+8)) = v68
												v116 = int32(_a_F_lookup_rowtype_tupdesc_internal_4)
												v118 = *(*int64)(unsafe.Add(mBase, _c_F_lookup_rowtype_tupdesc_internal[3]))
												v120 = v118 + int64(1)
												*(*int64)(unsafe.Add(mBase, _c_F_lookup_rowtype_tupdesc_internal[3])) = v120
												v122 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
												*(*int64)(unsafe.Add(mBase, uint32(v108+v122<<(uint(v112)%32)))) = v120
												v128 = *(*int32)(unsafe.Add(mBase, _c_F_lookup_rowtype_tupdesc_internal[2]))
												v129 = *(*int32)(unsafe.Add(mBase, uint32(v128)+16))
												F_dshash_release_lock(m, v129, v60)
												mBase = m.M
												v131 = m.ExcPending
												if v131 != 0 {
													return int32(0)
												} else {
													v133 = *(*int32)(unsafe.Add(mBase, _c_F_lookup_rowtype_tupdesc_internal[1]))
													v134 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
													v138 = *(*int32)(unsafe.Add(mBase, uint32(v133+v134<<(uint(int32(4))%32))+8))
													v140 = v138
													m.G0 = v10 + int32(16)
													return v140
												}
											}
										} else {
											v108 = v86
											v111 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
											v112 = int32(4)
											*(*int32)(unsafe.Add(mBase, uint32(v108+v111<<(uint(v112)%32))+8)) = v68
											v116 = int32(_a_F_lookup_rowtype_tupdesc_internal_4)
											v118 = *(*int64)(unsafe.Add(mBase, _c_F_lookup_rowtype_tupdesc_internal[3]))
											v120 = v118 + int64(1)
											*(*int64)(unsafe.Add(mBase, _c_F_lookup_rowtype_tupdesc_internal[3])) = v120
											v122 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
											*(*int64)(unsafe.Add(mBase, uint32(v108+v122<<(uint(v112)%32)))) = v120
											v128 = *(*int32)(unsafe.Add(mBase, _c_F_lookup_rowtype_tupdesc_internal[2]))
											v129 = *(*int32)(unsafe.Add(mBase, uint32(v128)+16))
											F_dshash_release_lock(m, v129, v60)
											mBase = m.M
											v131 = m.ExcPending
											if v131 != 0 {
												return int32(0)
											} else {
												v133 = *(*int32)(unsafe.Add(mBase, _c_F_lookup_rowtype_tupdesc_internal[1]))
												v134 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
												v138 = *(*int32)(unsafe.Add(mBase, uint32(v133+v134<<(uint(int32(4))%32))+8))
												v140 = v138
												m.G0 = v10 + int32(16)
												return v140
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
func F_lstat(m *base.Module, l0 int32, l1 int32) int32 {
	var v5 int32
	_ = v5
	v5 = F___fstatat(m, int32(-100), l0, l1, int32(256))
	return v5
}
func F_lt_q_regex(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
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
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v55 int32
	_ = v55
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v94 int32
	_ = v94
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v114 int32
	_ = v114
	var v121 int32
	_ = v121
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = F_pg_detoast_datum(m, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v13 = F_pg_detoast_datum(m, v12)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	v19 = F_ArrayGetNItems(m, v16, v13+int32(16))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if v21 < int32(2) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L1
	} else {
		goto L35
	}
L6:
	;
	v24 = F_array_contains_nulls(m, v13)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L1
	} else {
		goto L31
	}
L9:
	;
	if v24 != 0 {
		goto L5
	} else {
		goto L10
	}
L10:
	;
	if int32(0) < v19 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v73 != v8 {
		goto L23
	} else {
		goto L24
	}
L12:
	;
	if v15 != 0 {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	goto L14
L14:
	;
	v72 = int32(0)
	goto L11
L15:
	;
	v34 = v15
	goto L17
L16:
	;
	v34 = (v16<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L17
L17:
	;
	v37 = v13 + v34
	v39 = v19
	goto L18
L18:
	;
	v45 = F_DirectFunctionCall2Coll(m, int32(_a_F_lt_q_regex_0), int32(0), v8, v37)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L1
	} else {
		goto L20
	}
L19:
	;
	goto L14
L20:
	;
	if v45 != 0 {
		v72 = int32(1)
		goto L11
	} else {
		goto L21
	}
L21:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
	v55 = int32(1)
	if v55 < v39 {
		v37 = v37 + (int32(base.Ui32(v47)>>(uint(int32(2))%32))+int32(3))&int32(2147483644)
		v39 = v39 - v55
		goto L18
	} else {
		goto L22
	}
L22:
	;
	goto L19
L23:
	;
	F_pfree(m, v8)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L1
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v77 != v13 {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	goto L25
L27:
	;
	F_pfree(m, v13)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L1
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	return v72
L30:
	;
	goto L29
L31:
	;
	F_errcode(m, int32(352845954))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	F_errmsg(m, int32(_a_F_lt_q_regex_1), int32(0))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	F_errfinish(m, int32(_a_F_lt_q_regex_2), int32(316), int32(_a_F_lt_q_regex_3))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L35:
	;
	F_errcode(m, int32(67108994))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	F_errmsg(m, int32(_a_F_lt_q_regex_4), int32(0))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	F_errfinish(m, int32(_a_F_lt_q_regex_2), int32(320), int32(_a_F_lt_q_regex_3))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_lt_q_rregex(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = F_DirectFunctionCall2Coll(m, int32(_a_F_lt_q_rregex_0), int32(0), v4, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return v6
	}
}
func F_ltreeparentsel(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v8 float64
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v8 = F_generic_restriction_selectivity(m, v2, v3, int32(0), v5, v6, float64(0.001))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		v12 = F_Float8GetDatum(m, v8)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int32(0)
		} else {
			return v12
		}
	}
}
func F_lz4_decompress_datum(m *base.Module) int32 {
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	F_errstart_cold(m, int32(21), int32(0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		F_errcode(m, int32(1088))
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			F_errmsg(m, int32(_a_F_lz4_decompress_datum_0), int32(0))
			v13 = m.ExcPending
			if v13 != 0 {
				return int32(0)
			} else {
				F_errdetail(m, int32(_a_F_lz4_decompress_datum_1), int32(0))
				v17 = m.ExcPending
				if v17 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_lz4_decompress_datum_2), int32(185), int32(_a_F_lz4_decompress_datum_3))
					v22 = m.ExcPending
					if v22 != 0 {
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
