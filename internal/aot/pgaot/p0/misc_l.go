package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"sync/atomic"
	"unsafe"
)

func F_LargeObjectExistsWithSnapshot(m *base.Module, l0 int32, l1 int32) int32 {
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v16 int32
	_ = v16
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
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	v5 = m.G0
	v7 = v5 + int32(-64)
	m.G0 = v7
	F_ScanKeyInit(m, v7, int32(1), int32(3), int32(184), base.I64_extend_i32_u(l0))
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		v19 = F_table_open(m, int32(2995), int32(1))
		v20 = m.ExcPending
		if v20 != 0 {
			return int32(0)
		} else {
			v22 = int32(1)
			v24 = F_systable_beginscan(m, v19, int32(2996), v22, l1, v22, v7)
			v25 = m.ExcPending
			if v25 != 0 {
				return int32(0)
			} else {
				v26 = F_systable_getnext(m, v24)
				v27 = m.ExcPending
				if v27 != 0 {
					return int32(0)
				} else {
					F_systable_endscan(m, v24)
					v29 = m.ExcPending
					if v29 != 0 {
						return int32(0)
					} else {
						F_relation_close(m, v19, int32(1))
						v32 = m.ExcPending
						if v32 != 0 {
							return int32(0)
						} else {
							m.G0 = v7 - int32(-64)
							return base.B2i32(v26 != int32(0))
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
	var v15 int64
	_ = v15
	var v28 int32
	_ = v28
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v101 int64
	_ = v101
	var v104 int64
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v170 int32
	_ = v170
	var v175 int32
	_ = v175
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
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v205 int32
	_ = v205
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v230 int32
	_ = v230
	var v234 int32
	_ = v234
	var v248 int32
	_ = v248
	var v258 int32
	_ = v258
	var v264 int32
	_ = v264
	var v277 int32
	_ = v277
	var v288 int32
	_ = v288
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v302 int32
	_ = v302
	var v311 int32
	_ = v311
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v331 int32
	_ = v331
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v376 int32
	_ = v376
	var v382 int64
	_ = v382
	var v383 int64
	_ = v383
	var v384 int64
	_ = v384
	var v385 int64
	_ = v385
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v391 int32
	_ = v391
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v414 int32
	_ = v414
	var v427 int32
	_ = v427
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v435 int32
	_ = v435
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v449 int32
	_ = v449
	var v452 int32
	_ = v452
	var v463 int32
	_ = v463
	var v465 int32
	_ = v465
	var v471 int32
	_ = v471
	var v475 int32
	_ = v475
	var v480 int32
	_ = v480
	var v485 int32
	_ = v485
	var v490 int32
	_ = v490
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v517 int32
	_ = v517
	var v527 int32
	_ = v527
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v533 int32
	_ = v533
	var v536 int32
	_ = v536
	var v547 int32
	_ = v547
	var v549 int32
	_ = v549
	var v555 int32
	_ = v555
	var v559 int32
	_ = v559
	var v564 int32
	_ = v564
	var v570 int32
	_ = v570
	var v577 int64
	_ = v577
	var v581 int32
	_ = v581
	var v586 int32
	_ = v586
	var v596 int32
	_ = v596
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v601 int32
	_ = v601
	var v604 int32
	_ = v604
	var v612 int32
	_ = v612
	var v617 int32
	_ = v617
	v15 = base.I64_extend_i32_u(l0 + int32(12))
	goto L5
L1:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v612)+32)) = int64(0)
	return v617
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v601
	v612 = v599
	v617 = v604
	goto L1
L3:
	;
	v577 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+40)) = v577
	*(*int64)(unsafe.Add(mBase, uint32(l0)+4)) = v577
	v581 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if l1 != 0 {
		v599 = l0
		v601 = v581
		v604 = v570
		goto L2
	} else {
		goto L151
	}
L4:
	;
	v530 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v530 == int32(0) {
		v570 = v387
		goto L3
	} else {
		goto L136
	}
L5:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v28 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	v511 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v512 = int32(0)
	if l1 != 0 {
		v599 = l0
		v601 = v511
		v604 = v512
		goto L2
	} else {
		goto L130
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
	v292 = F_lookup_ts_dictionary_cache(m, v28)
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L32
	} else {
		goto L73
	}
L11:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v43 == int32(0) {
		goto L7
	} else {
		goto L13
	}
L13:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
	if v46 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v277
	v288 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v277)+12)) = v288
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v288
	goto L11
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v264
	v277 = v264
	goto L14
L16:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v56 <= v69 {
		v220 = v43
		goto L27
	} else {
		goto L28
	}
L17:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v43)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v59
	if v59 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L18:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)+12))
	if v50 <= v46 {
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v49)+16))
	v55 = v52 + v46<<(uint(int32(3))%32)
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
	if v56 != 0 {
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
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v65 == int32(0) {
		v264 = v43
		goto L15
	} else {
		goto L24
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v65)+12)) = v43
	v277 = v43
	goto L14
L25:
	;
	v258 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v258 == int32(0) {
		v264 = v248
		goto L15
	} else {
		goto L72
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = int32(0)
	v248 = v234
	goto L25
L27:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v220)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v230
	if v230 != 0 {
		v248 = v220
		goto L25
	} else {
		goto L71
	}
L28:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v43)+8))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
	v76 = v72
	v77 = v69
	v80 = v71
	goto L30
L29:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v185 != 0 {
		goto L57
	} else {
		goto L58
	}
L30:
	;
	v86 = v77 << (uint(int32(2)) % 32)
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v55)+4))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v86+v87)))
	v90 = F_lookup_ts_dictionary_cache(m, v89)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v183 != 0 {
		v220 = v183
		goto L27
	} else {
		goto L55
	}
L32:
	;
	return int32(0)
L33:
	;
	v94 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v94
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)) = uint16(v94)
	v101 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v90)+44)))
	v104 = F_FunctionCall4Coll(m, v90+int32(12), v94, v101, base.I64_extend_i32_u(v76), base.I64_extend_i32_s(v80), v15)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L32
	} else {
		goto L34
	}
L34:
	;
	v106 = base.I32_wrap_i64(v104)
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+13)))
	if v107 == int32(1) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v55)+4))
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v110+v86)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v77 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v112
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v43)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v117
	if v106 == int32(0) {
		goto L5
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	if v106 != 0 {
		goto L50
	} else {
		goto L51
	}
L38:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v121 != 0 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v121)+4))
	if v122 != 0 {
		goto L42
	} else {
		goto L43
	}
L40:
	;
	goto L41
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v106
	goto L5
L42:
	;
	v126 = v122
	v127 = v121
	goto L45
L43:
	;
	v153 = v121
	goto L44
L44:
	;
	F_pfree(m, v153)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L32
	} else {
		goto L49
	}
L45:
	;
	F_pfree(m, v126)
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L32
	} else {
		goto L47
	}
L46:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v153 = v140
	goto L44
L47:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v127)+12))
	if v137 != 0 {
		v126 = v137
		v127 = v127 + int32(8)
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
	v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106)+2)))
	if v170&int32(4) == int32(0) {
		goto L29
	} else {
		goto L53
	}
L51:
	;
	v177 = v76
	v178 = v80
	goto L52
L52:
	;
	v180 = v77 + int32(1)
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
	if v180 < v181 {
		v76 = v177
		v77 = v180
		v80 = v178
		goto L30
	} else {
		goto L54
	}
L53:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v106)+4))
	v176 = F_strlen(m, v175)
	mBase = m.M
	v177 = v175
	v178 = v176
	goto L52
L54:
	;
	goto L31
L55:
	;
	v234 = int32(0)
	goto L26
L56:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v192 != 0 {
		goto L62
	} else {
		goto L63
	}
L57:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v185)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v186
	if v186 != 0 {
		goto L56
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = int32(0)
	goto L56
L60:
	;
	goto L59
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v185
	v196 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v185)+12)) = v196
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v196
	v200 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if l1 != 0 {
		v599 = l0
		v601 = v200
		v604 = v106
		goto L2
	} else {
		goto L65
	}
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v192)+12)) = v185
	goto L61
L63:
	;
	goto L64
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v185
	goto L61
L65:
	;
	if v200 == int32(0) {
		v612 = l0
		v617 = v106
		goto L1
	} else {
		goto L66
	}
L66:
	;
	v205 = v200
	goto L67
L67:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v205)+12))
	F_pfree(m, v205)
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L32
	} else {
		goto L69
	}
L68:
	;
	v612 = l0
	v617 = v106
	goto L1
L69:
	;
	if v215 != 0 {
		v205 = v215
		goto L67
	} else {
		goto L70
	}
L70:
	;
	goto L68
L71:
	;
	v234 = v220
	goto L26
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v258)+12)) = v248
	v277 = v248
	goto L14
L73:
	;
	v294 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v294 == int32(0) {
		goto L7
	} else {
		goto L74
	}
L74:
	;
	v302 = v294
	goto L75
L75:
	;
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v302)))
	if v311 == int32(0) {
		goto L79
	} else {
		goto L80
	}
L76:
	;
	goto L7
L77:
	;
	if v490 != 0 {
		v302 = v490
		goto L75
	} else {
		goto L129
	}
L78:
	;
	v485 = *(*int32)(unsafe.Add(mBase, uint32(v302)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v485
	v490 = v485
	goto L77
L79:
	;
	v376 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+13)) = uint8(v376)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)) = uint8(base.B2i32(v311 == v376))
	v382 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v292)+44)))
	v383 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v302)+4)))
	v384 = int64(*(*int32)(unsafe.Add(mBase, uint32(v302)+8)))
	v385 = F_FunctionCall4Coll(m, v292+int32(12), v376, v382, v383, v384, v15)
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L32
	} else {
		goto L92
	}
L80:
	;
	v314 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v314)+12))
	if v315 <= v311 {
		goto L78
	} else {
		goto L81
	}
L81:
	;
	v317 = *(*int32)(unsafe.Add(mBase, uint32(v314)+16))
	v320 = v317 + v311<<(uint(int32(3))%32)
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v320)))
	if v321 == int32(0) {
		goto L78
	} else {
		goto L82
	}
L82:
	;
	if int32(0) < v321 {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v320)+4))
	v327 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v331 = int32(0)
	goto L86
L84:
	;
	goto L85
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(0)
	goto L5
L86:
	;
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v326+v331<<(uint(int32(2))%32))))
	v346 = v331 + int32(1)
	if v321 <= v346 {
		goto L88
	} else {
		goto L89
	}
L87:
	;
	if v344 == v327 {
		goto L79
	} else {
		goto L91
	}
L88:
	;
	goto L87
L89:
	;
	if v344 != v327 {
		v331 = v346
		goto L86
	} else {
		goto L90
	}
L90:
	;
	goto L88
L91:
	;
	goto L85
L92:
	;
	v387 = base.I32_wrap_i64(v385)
	v388 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+13)))
	if v388 == int32(1) {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v391 = *(*int32)(unsafe.Add(mBase, uint32(v302)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v391
	if v387 == int32(0) {
		goto L96
	} else {
		goto L97
	}
L94:
	;
	goto L95
L95:
	;
	if v387 != 0 {
		goto L4
	} else {
		goto L110
	}
L96:
	;
	v490 = v391
	goto L77
L97:
	;
	goto L98
L98:
	;
	v395 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v395 != 0 {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	v396 = *(*int32)(unsafe.Add(mBase, uint32(v395)+4))
	if v396 != 0 {
		goto L102
	} else {
		goto L103
	}
L100:
	;
	v435 = v391
	goto L101
L101:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v302
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v387
	v490 = v435
	goto L77
L102:
	;
	v399 = v395
	v401 = v396
	goto L105
L103:
	;
	v427 = v395
	goto L104
L104:
	;
	F_pfree(m, v427)
	mBase = m.M
	v429 = m.ExcPending
	if v429 != 0 {
		goto L32
	} else {
		goto L109
	}
L105:
	;
	F_pfree(m, v401)
	mBase = m.M
	v410 = m.ExcPending
	if v410 != 0 {
		goto L32
	} else {
		goto L107
	}
L106:
	;
	v414 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v427 = v414
	goto L104
L107:
	;
	v411 = *(*int32)(unsafe.Add(mBase, uint32(v399)+12))
	if v411 != 0 {
		v399 = v399 + int32(8)
		v401 = v411
		goto L105
	} else {
		goto L108
	}
L108:
	;
	goto L106
L109:
	;
	v430 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v435 = v430
	goto L101
L110:
	;
	v445 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v445 != 0 {
		goto L111
	} else {
		goto L112
	}
L111:
	;
	v446 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v446 == int32(0) {
		v570 = v445
		goto L3
	} else {
		goto L114
	}
L112:
	;
	goto L113
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(0)
	goto L5
L114:
	;
	v449 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v452 = v446
	goto L115
L115:
	;
	if v452 == v449 {
		goto L117
	} else {
		goto L118
	}
L116:
	;
	v570 = v445
	goto L3
L117:
	;
	v463 = *(*int32)(unsafe.Add(mBase, uint32(v449)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v463
	goto L119
L118:
	;
	goto L119
L119:
	;
	v465 = *(*int32)(unsafe.Add(mBase, uint32(v452)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v465
	if v465 == int32(0) {
		goto L120
	} else {
		goto L121
	}
L120:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = int32(0)
	goto L122
L121:
	;
	goto L122
L122:
	;
	v471 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v471 != 0 {
		goto L124
	} else {
		goto L125
	}
L123:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v452
	v475 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v452)+12)) = v475
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v475
	if v452 == v449 {
		v570 = v445
		goto L3
	} else {
		goto L127
	}
L124:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v471)+12)) = v452
	goto L123
L125:
	;
	goto L126
L126:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v452
	goto L123
L127:
	;
	v480 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v480 != 0 {
		v452 = v480
		goto L115
	} else {
		goto L128
	}
L128:
	;
	goto L116
L129:
	;
	goto L76
L130:
	;
	if v511 == int32(0) {
		v612 = l0
		v617 = v512
		goto L1
	} else {
		goto L131
	}
L131:
	;
	v517 = v511
	goto L132
L132:
	;
	v527 = *(*int32)(unsafe.Add(mBase, uint32(v517)+12))
	F_pfree(m, v517)
	mBase = m.M
	v529 = m.ExcPending
	if v529 != 0 {
		goto L32
	} else {
		goto L134
	}
L133:
	;
	v612 = l0
	v617 = v512
	goto L1
L134:
	;
	if v527 != 0 {
		v517 = v527
		goto L132
	} else {
		goto L135
	}
L135:
	;
	goto L133
L136:
	;
	v533 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v536 = v530
	goto L137
L137:
	;
	if v536 == v533 {
		goto L139
	} else {
		goto L140
	}
L138:
	;
	v570 = v387
	goto L3
L139:
	;
	v547 = *(*int32)(unsafe.Add(mBase, uint32(v533)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v547
	goto L141
L140:
	;
	goto L141
L141:
	;
	v549 = *(*int32)(unsafe.Add(mBase, uint32(v536)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v549
	if v549 == int32(0) {
		goto L142
	} else {
		goto L143
	}
L142:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = int32(0)
	goto L144
L143:
	;
	goto L144
L144:
	;
	v555 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v555 != 0 {
		goto L146
	} else {
		goto L147
	}
L145:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v536
	v559 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v536)+12)) = v559
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v559
	if v536 == v533 {
		v570 = v387
		goto L3
	} else {
		goto L149
	}
L146:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v555)+12)) = v536
	goto L145
L147:
	;
	goto L148
L148:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v536
	goto L145
L149:
	;
	v564 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v564 != 0 {
		v536 = v564
		goto L137
	} else {
		goto L150
	}
L150:
	;
	goto L138
L151:
	;
	if v581 == int32(0) {
		v612 = l0
		v617 = v570
		goto L1
	} else {
		goto L152
	}
L152:
	;
	v586 = v581
	goto L153
L153:
	;
	v596 = *(*int32)(unsafe.Add(mBase, uint32(v586)+12))
	F_pfree(m, v586)
	mBase = m.M
	v598 = m.ExcPending
	if v598 != 0 {
		goto L32
	} else {
		goto L155
	}
L154:
	;
	v612 = l0
	v617 = v570
	goto L1
L155:
	;
	if v596 != 0 {
		v586 = v596
		goto L153
	} else {
		goto L156
	}
L156:
	;
	goto L154
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
	F_errstart_cold(m, int32(24), int32(0))
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
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v81)+364))
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
	F_errfinish(m, int32(_a_F_LockCheckConflicts_1), int32(1666), int32(_a_F_LockCheckConflicts_2))
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
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v92 int32
	_ = v92
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v106 int64
	_ = v106
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v184 int32
	_ = v184
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v204 int32
	_ = v204
	v3 = int32(0)
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	if l0 == v3 {
		v184 = v3
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L12
	} else {
		goto L61
	}
L2:
	;
	m.G0 = v13 + int32(16)
	return v184
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
	v172 = F_expression_tree_walker_impl(m, l0, int32(603), l1)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L12
	} else {
		goto L60
	}
L7:
	;
	v169 = F_query_tree_walker_impl(m, l0, int32(603), l1, int32(4))
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L12
	} else {
		goto L59
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
	if base.B2i32(base.Ui32(int32(6)) < base.Ui32(v49))|base.B2i32(int32(1)<<(uint(v49)%32)&int32(69) == int32(0)) != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v154 = v34 + int32(1)
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	if v154 < v155 {
		v34 = v154
		goto L10
	} else {
		goto L58
	}
L15:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v60 = int32(0)
	if v59 == v60 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	if v98 != 0 {
		goto L14
	} else {
		goto L29
	}
L17:
	;
	v98 = int32(0)
	goto L16
L18:
	;
	goto L19
L19:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v59)+4))
	if v66 <= int32(0) {
		v92 = v60
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v98 = v92
	goto L16
L21:
	;
	v69 = int32(0)
	if v69 < v66 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v72 = v66
	goto L24
L23:
	;
	v72 = v69
	goto L24
L24:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v59)+12))
	v75 = int32(0)
	goto L25
L25:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v73+v75<<(uint(int32(2))%32))))
	v84 = base.B2i32(v83 == v43)
	if v83 == v43 {
		v92 = v84
		goto L20
	} else {
		goto L27
	}
L26:
	;
	v92 = v84
	goto L20
L27:
	;
	v86 = v75 + int32(1)
	if v86 != v72 {
		v75 = v86
		goto L25
	} else {
		goto L28
	}
L28:
	;
	goto L26
L29:
	;
	v99 = base.I32_extend8_s(v42)
	v100 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v103 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v103 < int32(2) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v106 = int64(16414)
	goto L32
L31:
	;
	v106 = int64(16412)
	goto L32
L32:
	;
	v111 = F_pg_class_aclcheck(m, v43, v100, v106|base.I64_extend_i32_u(base.B2i32(v103 < int32(4))))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L12
	} else {
		goto L33
	}
L33:
	;
	if v111 != 0 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	switch v99 - int32(73) {
	case 0, 32:
		v122 = int32(20)
		goto L38
	default:
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
L35:
	;
	goto L36
L36:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+4)))
	if v128 == int32(0) {
		goto L46
	} else {
		goto L47
	}
L37:
	;
	F_aclcheck_error(m, v111, v124, v44)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L12
	} else {
		goto L44
	}
L38:
	;
	v124 = v122
	goto L37
L39:
	;
	v122 = int32(42)
	goto L38
L40:
	;
	v124 = int32(18)
	goto L37
L41:
	;
	v124 = int32(23)
	goto L37
L42:
	;
	v124 = int32(52)
	goto L37
L43:
	;
	v124 = int32(38)
	goto L37
L44:
	;
	goto L36
L45:
	;
	if v99 == int32(118) {
		goto L52
	} else {
		goto L53
	}
L46:
	;
	F_LockRelationOid(m, v43, v127)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L12
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	v133 = F_ConditionalLockRelationOid(m, v43, v127)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L12
	} else {
		goto L50
	}
L49:
	;
	goto L45
L50:
	;
	if v133 == int32(0) {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	goto L45
L52:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+4)))
	v141 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	F_LockViewRecurse(m, v43, v139, v140, v141)
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L12
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+20)))
	if v144 != int32(1) {
		goto L14
	} else {
		goto L56
	}
L55:
	;
	goto L14
L56:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+4)))
	F_LockTableRecurse(m, v43, v147, v148)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L12
	} else {
		goto L57
	}
L57:
	;
	goto L14
L58:
	;
	goto L11
L59:
	;
	v184 = v169
	goto L2
L60:
	;
	v184 = v172
	goto L2
L61:
	;
	F_errcode(m, int32(50463045))
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L12
	} else {
		goto L62
	}
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v44
	F_errmsg(m, int32(_a_F_LockViewRecurse_walker_0), v13)
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L12
	} else {
		goto L63
	}
L63:
	;
	F_errfinish(m, int32(_a_F_LockViewRecurse_walker_1), int32(224), int32(_a_F_LockViewRecurse_walker_2))
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L12
	} else {
		goto L64
	}
L64:
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
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v47 int64
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = int32(_a_F_LookupCreationNamespace_0)
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v15 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_LookupCreationNamespace[0])))
	if base.B2i32(v12 == int32(0))|base.B2i32(v12 != v15) != 0 {
		v33 = v12
		v34 = v15
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L13
	} else {
		goto L20
	}
L2:
	;
	m.G0 = v7 + int32(16)
	return v65
L3:
	;
	if v33-v34 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L4:
	;
	goto L3
L5:
	;
	v18 = l0
	v19 = v9
	goto L6
L6:
	;
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+1)))
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+1)))
	if v23 == int32(0) {
		v33 = v23
		v34 = v22
		goto L4
	} else {
		goto L8
	}
L7:
	;
	v33 = v23
	v34 = v22
	goto L4
L8:
	;
	v26 = int32(1)
	if v23 == v22 {
		v18 = v18 + v26
		v19 = v19 + v26
		goto L6
	} else {
		goto L9
	}
L9:
	;
	goto L7
L10:
	;
	F_AccessTempTableNamespace(m, int32(0))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	goto L12
L12:
	;
	v47 = int64(0)
	v50 = F_GetSysCacheOid(m, int32(37), base.I64_extend_i32_u(l0), v47, v47, v47)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L13
	} else {
		goto L15
	}
L13:
	;
	return int32(0)
L14:
	;
	v44 = *(*int32)(unsafe.Add(mBase, _c_F_LookupCreationNamespace[1]))
	v65 = v44
	goto L2
L15:
	;
	if v50 == int32(0) {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	v56 = *(*int32)(unsafe.Add(mBase, _c_F_LookupCreationNamespace[2]))
	v58 = F_object_aclcheck(m, int32(2615), v50, v56, int64(512))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L13
	} else {
		goto L17
	}
L17:
	;
	if v58 == int32(0) {
		v65 = v50
		goto L2
	} else {
		goto L18
	}
L18:
	;
	F_aclcheck_error(m, v58, int32(37), l0)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L13
	} else {
		goto L19
	}
L19:
	;
	v65 = v50
	goto L2
L20:
	;
	F_errcode(m, int32(1411))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L13
	} else {
		goto L21
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
	F_errmsg(m, int32(_a_F_LookupCreationNamespace_1), v7)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L13
	} else {
		goto L22
	}
L22:
	;
	F_errfinish(m, int32(_a_F_LookupCreationNamespace_2), int32(3616), int32(_a_F_LookupCreationNamespace_3))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L13
	} else {
		goto L23
	}
L23:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F__ltq_regex(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14226(m, l0, int32(_a_F__ltq_regex_0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return v3
	}
}
func F_l2_normalize(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v9 float64
	_ = v9
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
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v48 float64
	_ = v48
	var v51 int32
	_ = v51
	var v52 float32
	_ = v52
	var v53 float64
	_ = v53
	var v55 float32
	_ = v55
	var v56 float64
	_ = v56
	var v58 float32
	_ = v58
	var v59 float64
	_ = v59
	var v61 float32
	_ = v61
	var v62 float64
	_ = v62
	var v67 float64
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v86 float64
	_ = v86
	var v89 int32
	_ = v89
	var v96 int32
	_ = v96
	var v98 float64
	_ = v98
	var v102 float32
	_ = v102
	var v103 float64
	_ = v103
	var v105 float64
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v120 float64
	_ = v120
	var v126 int32
	_ = v126
	var v127 float64
	_ = v127
	var v128 int32
	_ = v128
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v148 float32
	_ = v148
	var v154 int32
	_ = v154
	var v157 float32
	_ = v157
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v171 int32
	_ = v171
	var v182 int32
	_ = v182
	var v185 float32
	_ = v185
	var v201 int32
	_ = v201
	var v214 float32
	_ = v214
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	v2 = int32(0)
	v9 = float64(0)
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v14 = F_pg_detoast_datum(m, v13)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v18 = int32(*(*int16)(unsafe.Add(mBase, uint32(v14)+4)))
	v19 = F_mul_size(m, int32(4), v18)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v21 = F_add_size(m, int32(8), v19)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v23 = F_palloc0(m, v21)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v23)+4)) = uint16(v18)
	*(*int32)(unsafe.Add(mBase, uint32(v23))) = v21 << (uint(int32(2)) % 32)
	v29 = int32(*(*int16)(unsafe.Add(mBase, uint32(v14)+4)))
	if v29 <= int32(0) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	return base.I64_extend_i32_u(v23)
L7:
	;
	v33 = v14 + int32(8)
	v34 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v29) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	if base.F64_gt(v120, float64(0)) == int32(0) {
		goto L6
	} else {
		goto L19
	}
L9:
	;
	v39 = v34
	v44 = v2
	v48 = v9
	goto L12
L10:
	;
	v77 = v34
	v86 = v9
	goto L11
L11:
	;
	v89 = v77
	v96 = v2
	v98 = v86
	goto L16
L12:
	;
	v51 = v33 + v39<<(uint(int32(2))%32)
	v52 = *(*float32)(unsafe.Add(mBase, uint32(v51)+12))
	v53 = base.F64_promote_f32(v52)
	v55 = *(*float32)(unsafe.Add(mBase, uint32(v51)+8))
	v56 = base.F64_promote_f32(v55)
	v58 = *(*float32)(unsafe.Add(mBase, uint32(v51)+4))
	v59 = base.F64_promote_f32(v58)
	v61 = *(*float32)(unsafe.Add(mBase, uint32(v51)))
	v62 = base.F64_promote_f32(v61)
	v67 = base.F64_add(base.F64_mul(v53, v53), base.F64_add(base.F64_mul(v56, v56), base.F64_add(base.F64_mul(v59, v59), base.F64_add(base.F64_mul(v62, v62), v48))))
	v68 = int32(4)
	v69 = v39 + v68
	v71 = v44 + v68
	if v71 != v29&int32(_a_F_l2_normalize_0) {
		v39 = v69
		v44 = v71
		v48 = v67
		goto L12
	} else {
		goto L14
	}
L13:
	;
	if v29&int32(3) == int32(0) {
		v120 = v67
		goto L8
	} else {
		goto L15
	}
L14:
	;
	goto L13
L15:
	;
	v77 = v69
	v86 = v67
	goto L11
L16:
	;
	v102 = *(*float32)(unsafe.Add(mBase, uint32(v33+v89<<(uint(int32(2))%32))))
	v103 = base.F64_promote_f32(v102)
	v105 = base.F64_add(base.F64_mul(v103, v103), v98)
	v106 = int32(1)
	v109 = v96 + v106
	if v109 != v29&int32(3) {
		v89 = v89 + v106
		v96 = v109
		v98 = v105
		goto L16
	} else {
		goto L18
	}
L17:
	;
	v120 = v105
	goto L8
L18:
	;
	goto L17
L19:
	;
	v126 = v23 + int32(8)
	v127 = base.F64_sqrt(v120)
	v128 = int32(0)
	if v29 != int32(1) {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	v201 = int32(0)
	goto L28
L21:
	;
	v134 = v128
	v139 = int32(0)
	goto L24
L22:
	;
	v171 = v128
	goto L23
L23:
	;
	v182 = v171 << (uint(int32(2)) % 32)
	v185 = *(*float32)(unsafe.Add(mBase, uint32(v182+v33)))
	*(*float32)(unsafe.Add(mBase, uint32(v126+v182))) = base.F32_demote_f64(base.F64_div(base.F64_promote_f32(v185), v127))
	goto L20
L24:
	;
	v144 = int32(2)
	v145 = v134 << (uint(v144) % 32)
	v148 = *(*float32)(unsafe.Add(mBase, uint32(v33+v145)))
	*(*float32)(unsafe.Add(mBase, uint32(v126+v145))) = base.F32_demote_f64(base.F64_div(base.F64_promote_f32(v148), v127))
	v154 = v145 | int32(4)
	v157 = *(*float32)(unsafe.Add(mBase, uint32(v33+v154)))
	*(*float32)(unsafe.Add(mBase, uint32(v126+v154))) = base.F32_demote_f64(base.F64_div(base.F64_promote_f32(v157), v127))
	v163 = v134 + v144
	v165 = v139 + v144
	if v165 != v29&int32(_a_F_l2_normalize_1) {
		v134 = v163
		v139 = v165
		goto L24
	} else {
		goto L26
	}
L25:
	;
	if v29&int32(1) == int32(0) {
		goto L20
	} else {
		goto L27
	}
L26:
	;
	goto L25
L27:
	;
	v171 = v163
	goto L23
L28:
	;
	v214 = *(*float32)(unsafe.Add(mBase, uint32(v126+v201<<(uint(int32(2))%32))))
	if base.F32_ne(base.F32_abs(v214), math.Float32frombits(uint32(0x7f800000))) != 0 {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	F_float_overflow_error(m)
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L1
	} else {
		goto L34
	}
L30:
	;
	v219 = v201 + int32(1)
	if v29 != v219 {
		v201 = v219
		goto L28
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	goto L29
L33:
	;
	goto L6
L34:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_lappend(m *base.Module, l0 int32, l1 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn14324(m, l0, l1, int64(4294967297))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_lappend_xid(m *base.Module, l0 int32, l1 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn14324(m, l0, l1, int64(4294967777))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_last_dir_separator(m *base.Module, l0 int32) int32 {
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
	v4 = l0
	v6 = int32(0)
	for {
		v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4))))
		if v7 != int32(47) {
		} else {
			v11 = v4
			v4 = v4 + int32(1)
			v6 = v11
			continue
		}
		if v7 != 0 {
			v11 = v6
			v4 = v4 + int32(1)
			v6 = v11
			continue
		} else {
			break
		}
		break
	}
	return v6
}
func F_length_in_encoding(m *base.Module, l0 int32) int64 {
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
	var v62 int32
	_ = v62
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
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v135 int32
	_ = v135
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v275 int32
	_ = v275
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v315 int32
	_ = v315
	var v328 int32
	_ = v328
	var v331 int32
	_ = v331
	var v335 int32
	_ = v335
	var v340 int32
	_ = v340
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
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
	return int64(0)
L2:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
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
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+1)))
	if v62 != 0 {
		v38 = v38 + int32(1)
		v39 = v62
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
	v328 = m.ExcPending
	if v328 != 0 {
		goto L1
	} else {
		goto L100
	}
L32:
	;
	v142 = int32(1)
	if v112&v142 != 0 {
		goto L43
	} else {
		goto L44
	}
L33:
	;
	v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+1)))
	if v118 == int32(18) {
		goto L36
	} else {
		goto L37
	}
L34:
	;
	goto L35
L35:
	;
	v129 = int32(1)
	if v112&v129 != 0 {
		v141 = int32(base.Ui32(v112)>>(uint(v129)%32)) - v129
		goto L32
	} else {
		goto L42
	}
L36:
	;
	v121 = int32(16)
	goto L38
L37:
	;
	v121 = int32(0)
	goto L38
L38:
	;
	if base.Ui32((v118-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v128 = int32(4)
	goto L41
L40:
	;
	v128 = v121
	goto L41
L41:
	;
	v141 = v128
	goto L32
L42:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	v141 = int32(base.Ui32(v135)>>(uint(int32(2))%32)) - int32(4)
	goto L32
L43:
	;
	v146 = v142
	goto L45
L44:
	;
	v146 = int32(4)
	goto L45
L45:
	;
	v147 = v13 + v146
	v148 = int32(0)
	if base.B2i32(v104 == int32(7))|base.B2i32(base.Ui32(int32(41)) < base.Ui32(v104)) != 0 {
		goto L48
	} else {
		goto L49
	}
L46:
	;
	m.G0 = v10 + int32(16)
	return base.I64_extend_i32_s(v315)
L47:
	;
	if v160 <= int32(1) {
		goto L51
	} else {
		goto L52
	}
L48:
	;
	v160 = int32(1)
	goto L50
L49:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v104*int32(28))+uint32(_c_F_length_in_encoding[0])))
	v160 = v159
	goto L50
L50:
	;
	goto L47
L51:
	;
	v164 = int32(0)
	if base.B2i32(v147&int32(3) == v164)|base.B2i32(v141 == v164) != 0 {
		v195 = v147
		v197 = v141
		v198 = base.B2i32(v141 != v164)
		goto L57
	} else {
		goto L58
	}
L52:
	;
	goto L53
L53:
	;
	if v141 <= int32(0) {
		v315 = v148
		goto L46
	} else {
		goto L84
	}
L54:
	;
	if v269 != 0 {
		goto L79
	} else {
		goto L80
	}
L55:
	;
	v269 = int32(0)
	goto L54
L56:
	;
	v247 = v240
	v249 = v242
	goto L73
L57:
	;
	if v198 == int32(0) {
		goto L55
	} else {
		goto L64
	}
L58:
	;
	v178 = v147
	v180 = v141
	goto L59
L59:
	;
	v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178))))
	if v183 == int32(0) {
		v240 = v178
		v242 = v180
		goto L56
	} else {
		goto L61
	}
L60:
	;
	v195 = v190
	v197 = v186
	v198 = v188
	goto L57
L61:
	;
	v185 = int32(1)
	v186 = v180 - v185
	v187 = int32(0)
	v188 = base.B2i32(v186 != v187)
	v190 = v178 + v185
	if v190&int32(3) == v187 {
		v195 = v190
		v197 = v186
		v198 = v188
		goto L57
	} else {
		goto L62
	}
L62:
	;
	if v186 != 0 {
		v178 = v190
		v180 = v186
		goto L59
	} else {
		goto L63
	}
L63:
	;
	goto L60
L64:
	;
	v203 = int32(0)
	v204 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v195))))
	if base.B2i32(v203 == v204)|base.B2i32(base.Ui32(v197) < base.Ui32(int32(4))) == v203 {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v213 = v195
	v215 = v197
	goto L68
L66:
	;
	v233 = v195
	v235 = v197
	goto L67
L67:
	;
	if v235 == int32(0) {
		goto L55
	} else {
		goto L72
	}
L68:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v213)))
	v220 = v219 ^ int32(0)
	v223 = int32(-2139062144)
	if (int32(16843008)-v220|v220)&v223 != v223 {
		v240 = v213
		v242 = v215
		goto L56
	} else {
		goto L70
	}
L69:
	;
	v233 = v228
	v235 = v230
	goto L67
L70:
	;
	v227 = int32(4)
	v228 = v213 + v227
	v230 = v215 - v227
	if base.Ui32(int32(3)) < base.Ui32(v230) {
		v213 = v228
		v215 = v230
		goto L68
	} else {
		goto L71
	}
L71:
	;
	goto L69
L72:
	;
	v240 = v233
	v242 = v235
	goto L56
L73:
	;
	v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v247))))
	if int32(0) == v252 {
		goto L75
	} else {
		goto L76
	}
L74:
	;
	goto L55
L75:
	;
	v269 = v247
	goto L54
L76:
	;
	goto L77
L77:
	;
	v254 = int32(1)
	v257 = v249 - v254
	if v257 != 0 {
		v247 = v247 + v254
		v249 = v257
		goto L73
	} else {
		goto L78
	}
L78:
	;
	goto L74
L79:
	;
	v270 = int32(-1)
	goto L81
L80:
	;
	v270 = v141
	goto L81
L81:
	;
	if v269 == int32(0) {
		v315 = v270
		goto L46
	} else {
		goto L82
	}
L82:
	;
	F_report_invalid_encoding(m, v104, v269, int32(1))
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L1
	} else {
		goto L83
	}
L83:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L84:
	;
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v104*int32(28))+uint32(_c_F_length_in_encoding[1])))
	v283 = v141
	v284 = v147
	v285 = v148
	goto L85
L85:
	;
	v290 = int32(*(*int8)(unsafe.Add(mBase, uint32(v284))))
	if int32(0) <= v290 {
		goto L90
	} else {
		goto L91
	}
L86:
	;
	v315 = v310
	goto L46
L87:
	;
	v310 = v285 + int32(1)
	if int32(0) < v306 {
		v283 = v306
		v284 = v308
		v285 = v310
		goto L85
	} else {
		goto L99
	}
L88:
	;
	v306 = v283 - v297
	v308 = v284 + v297
	goto L87
L89:
	;
	F_report_invalid_encoding(m, v104, v284, v283)
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L1
	} else {
		goto L98
	}
L90:
	;
	if v290 != 0 {
		goto L93
	} else {
		goto L94
	}
L91:
	;
	goto L92
L92:
	;
	v297 = m.T0[v282].(func(*base.Module, int32, int32) int32)(m, v284, v283)
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L1
	} else {
		goto L96
	}
L93:
	;
	v293 = int32(1)
	v306 = v283 - v293
	v308 = v284 + v293
	goto L87
L94:
	;
	goto L95
L95:
	;
	goto L89
L96:
	;
	if int32(0) <= v297 {
		goto L88
	} else {
		goto L97
	}
L97:
	;
	goto L89
L98:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L99:
	;
	goto L86
L100:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L1
	} else {
		goto L101
	}
L101:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v17
	F_errmsg(m, int32(_a_F_length_in_encoding_2), v10)
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L1
	} else {
		goto L102
	}
L102:
	;
	F_errfinish(m, int32(_a_F_length_in_encoding_3), int32(639), int32(_a_F_length_in_encoding_4))
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L1
	} else {
		goto L103
	}
L103:
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
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v309 int32
	_ = v309
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v326 int32
	_ = v326
	var v328 int32
	_ = v328
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v351 int32
	_ = v351
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v361 int32
	_ = v361
	var v366 int32
	_ = v366
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v374 int32
	_ = v374
	var v376 int32
	_ = v376
	var v378 int32
	_ = v378
	var v392 int32
	_ = v392
	var v395 int32
	_ = v395
	var v397 int32
	_ = v397
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v409 int32
	_ = v409
	var v411 int32
	_ = v411
	var v414 int32
	_ = v414
	var v419 int32
	_ = v419
	var v423 int32
	_ = v423
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v441 int32
	_ = v441
	var v443 int32
	_ = v443
	var v445 int32
	_ = v445
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v455 int32
	_ = v455
	var v460 int32
	_ = v460
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v468 int32
	_ = v468
	var v470 int32
	_ = v470
	var v472 int32
	_ = v472
	var v486 int32
	_ = v486
	var v489 int32
	_ = v489
	var v491 int32
	_ = v491
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v503 int32
	_ = v503
	var v505 int32
	_ = v505
	var v508 int32
	_ = v508
	var v515 int32
	_ = v515
	var v521 int32
	_ = v521
	var v527 int32
	_ = v527
	var v529 int32
	_ = v529
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v545 int32
	_ = v545
	var v549 int32
	_ = v549
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v4 + int32(4)
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
	v10 = v8 - int32(48)
	if base.B2i32(base.Ui32(v10) < base.Ui32(int32(10)))|base.B2i32(base.Ui32(v8&int32(-33)-int32(65)) < base.Ui32(int32(26))) == int32(0) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return int32(1)
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v549
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(112)
	goto L1
L3:
	;
	v549 = v8
	goto L2
L4:
	;
	goto L5
L5:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+8)) = v23 | int32(128)
	switch v10 {
	case 0:
		goto L8
	case 1, 2, 3, 4, 5, 6, 7, 8, 9:
		goto L9
	default:
		goto L7
	case 17:
		goto L33
	case 18:
		goto L31
	case 20:
		goto L28
	case 29:
		goto L24
	case 35:
		goto L20
	case 37:
		goto L17
	case 39:
		goto L14
	case 41:
		goto L11
	case 42:
		goto L10
	case 49:
		goto L34
	case 50:
		goto L32
	case 51:
		goto L30
	case 52:
		goto L29
	case 53:
		goto L27
	case 54:
		goto L26
	case 61:
		goto L25
	case 62:
		goto L23
	case 66:
		goto L22
	case 67:
		goto L21
	case 68:
		goto L19
	case 69:
		goto L18
	case 70:
		goto L16
	case 71:
		goto L15
	case 72:
		goto L13
	case 73:
		goto L12
	}
L6:
	;
	v534 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v535 = int32(0)
	if base.B2i32(v534 == v535)&base.B2i32(base.Ui32(v533) < base.Ui32(int32(2147483647))) == v535 {
		goto L137
	} else {
		goto L138
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
	v527 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v527 != 0 {
		goto L134
	} else {
		goto L135
	}
L8:
	;
	v436 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v437 = *(*int32)(unsafe.Add(mBase, uint32(v436)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v436)+8)) = v437 | int32(512)
	v441 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v443 = v441 - int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v443
	v445 = int32(8)
	v447 = int32(3)
	v448 = int32(0)
	v455 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v460 = v443
	v462 = v448
	v463 = v448
	goto L116
L9:
	;
	v347 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v349 = v347 - int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v349
	v351 = int32(10)
	v353 = int32(255)
	v354 = int32(0)
	v361 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v366 = v349
	v368 = v354
	v369 = v354
	goto L92
L10:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(90)
	goto L1
L11:
	;
	v338 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v338)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v338)+8)) = v339 | int32(1024)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(87)
	goto L1
L12:
	;
	v331 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v331)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v331)+8)) = v332 | int32(1024)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(119)
	goto L1
L13:
	;
	v263 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v263)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v263)+8)) = v264 | int32(512)
	v268 = int32(16)
	v270 = int32(255)
	v271 = int32(0)
	v277 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v278 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v283 = v277
	v285 = v271
	v286 = v271
	goto L77
L14:
	;
	v256 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v256)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v256)+8)) = v257 | int32(1024)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(55834574947)
	goto L1
L15:
	;
	v249 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v249)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v249)+8)) = v250 | int32(1024)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(55834574963)
	goto L1
L16:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(47244640368)
	goto L1
L17:
	;
	v184 = int32(16)
	v185 = int32(8)
	v187 = int32(0)
	v193 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v194 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v199 = v193
	v201 = v187
	v202 = v187
	goto L62
L18:
	;
	v121 = int32(16)
	v122 = int32(4)
	v124 = int32(0)
	v130 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v131 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v136 = v130
	v138 = v124
	v139 = v124
	goto L47
L19:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(38654705776)
	goto L1
L20:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v112)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v112)+8)) = v113 | int32(1024)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(42949673059)
	goto L1
L21:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v105)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v105)+8)) = v106 | int32(1024)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(42949673075)
	goto L1
L22:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(55834574960)
	goto L1
L23:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(42949673072)
	goto L1
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(62)
	goto L1
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(60)
	goto L1
L26:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(51539607664)
	goto L1
L27:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v82)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v82)+8)) = v83 | int32(512)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(112)
	v92 = F_chrnamed(m, l0, int32(_a_F_lexescape_0), int32(_a_F_lexescape_1), int32(27))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L35
	} else {
		goto L43
	}
L28:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v75)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v75)+8)) = v76 | int32(1024)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(21474836579)
	goto L1
L29:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v68)+8)) = v69 | int32(1024)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(21474836595)
	goto L1
L30:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v43)+8)) = v44 | int32(512)
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if base.Ui32(v49) <= base.Ui32(v48) {
		goto L37
	} else {
		goto L38
	}
L31:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(395136991344)
	goto L1
L32:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(34359738480)
	goto L1
L33:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(65)
	goto L1
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(112)
	v32 = F_chrnamed(m, l0, int32(_a_F_lexescape_2), int32(_a_F_lexescape_0), int32(7))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	return int32(0)
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v32
	goto L1
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v53 != 0 {
		goto L40
	} else {
		goto L41
	}
L38:
	;
	goto L39
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(112)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v48 + int32(4)
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v64 & int32(31)
	goto L1
L40:
	;
	v55 = v53
	goto L42
L41:
	;
	v55 = int32(5)
	goto L42
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v55
	return int32(0)
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v92
	goto L1
L44:
	;
	v533 = v174
	goto L6
L45:
	;
	if base.Ui32(v173) < base.Ui32(v122) {
		goto L53
	} else {
		goto L54
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v136
	v173 = v138
	v174 = v139
	goto L45
L47:
	;
	if base.Ui32(v131) <= base.Ui32(v136) {
		v173 = v138
		v174 = v139
		goto L45
	} else {
		goto L49
	}
L48:
	;
	v173 = v122
	v174 = v165
	goto L45
L49:
	;
	v144 = v136 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v144
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v136)))
	v148 = v146 - int32(48)
	if base.B2i32(base.Ui32(int32(54)) < base.Ui32(v148))|base.B2i32(base.I32_wrap_i64(int64(base.Ui64(int64(35465847073801215))>>(uint(base.I64_extend_i32_u(v148))%64)))&int32(1) == int32(0)) != 0 {
		goto L46
	} else {
		goto L50
	}
L50:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v148<<(uint(int32(2))%32))+uint32(_c_F_lexescape[0])))
	if base.Ui32(v121) <= base.Ui32(v162) {
		goto L46
	} else {
		goto L51
	}
L51:
	;
	v165 = v162 + v121*v139
	v167 = v138 + int32(1)
	if v167 != v122 {
		v136 = v144
		v138 = v167
		v139 = v165
		goto L47
	} else {
		goto L52
	}
L52:
	;
	goto L48
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
	v179 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v179 != 0 {
		goto L56
	} else {
		goto L57
	}
L54:
	;
	goto L55
L55:
	;
	goto L44
L56:
	;
	v181 = v179
	goto L58
L57:
	;
	v181 = int32(5)
	goto L58
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v181
	goto L55
L59:
	;
	v533 = v237
	goto L6
L60:
	;
	if base.Ui32(v236) < base.Ui32(v185) {
		goto L68
	} else {
		goto L69
	}
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v199
	v236 = v201
	v237 = v202
	goto L60
L62:
	;
	if base.Ui32(v194) <= base.Ui32(v199) {
		v236 = v201
		v237 = v202
		goto L60
	} else {
		goto L64
	}
L63:
	;
	v236 = v185
	v237 = v228
	goto L60
L64:
	;
	v207 = v199 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v207
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v199)))
	v211 = v209 - int32(48)
	if base.B2i32(base.Ui32(int32(54)) < base.Ui32(v211))|base.B2i32(base.I32_wrap_i64(int64(base.Ui64(int64(35465847073801215))>>(uint(base.I64_extend_i32_u(v211))%64)))&int32(1) == int32(0)) != 0 {
		goto L61
	} else {
		goto L65
	}
L65:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v211<<(uint(int32(2))%32))+uint32(_c_F_lexescape[0])))
	if base.Ui32(v184) <= base.Ui32(v225) {
		goto L61
	} else {
		goto L66
	}
L66:
	;
	v228 = v225 + v184*v202
	v230 = v201 + int32(1)
	if v230 != v185 {
		v199 = v207
		v201 = v230
		v202 = v228
		goto L62
	} else {
		goto L67
	}
L67:
	;
	goto L63
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
	v242 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v242 != 0 {
		goto L71
	} else {
		goto L72
	}
L69:
	;
	goto L70
L70:
	;
	goto L59
L71:
	;
	v244 = v242
	goto L73
L72:
	;
	v244 = int32(5)
	goto L73
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v244
	goto L70
L74:
	;
	v533 = v321
	goto L6
L75:
	;
	if base.Ui32(v320) < base.Ui32(int32(1)) {
		goto L83
	} else {
		goto L84
	}
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v283
	v320 = v285
	v321 = v286
	goto L75
L77:
	;
	if base.Ui32(v278) <= base.Ui32(v283) {
		v320 = v285
		v321 = v286
		goto L75
	} else {
		goto L79
	}
L78:
	;
	v320 = v270
	v321 = v312
	goto L75
L79:
	;
	v291 = v283 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v291
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v283)))
	v295 = v293 - int32(48)
	if base.B2i32(base.Ui32(int32(54)) < base.Ui32(v295))|base.B2i32(base.I32_wrap_i64(int64(base.Ui64(int64(35465847073801215))>>(uint(base.I64_extend_i32_u(v295))%64)))&int32(1) == int32(0)) != 0 {
		goto L76
	} else {
		goto L80
	}
L80:
	;
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v295<<(uint(int32(2))%32))+uint32(_c_F_lexescape[0])))
	if base.Ui32(v268) <= base.Ui32(v309) {
		goto L76
	} else {
		goto L81
	}
L81:
	;
	v312 = v309 + v268*v286
	v314 = v285 + int32(1)
	if v314 != v270 {
		v283 = v291
		v285 = v314
		v286 = v312
		goto L77
	} else {
		goto L82
	}
L82:
	;
	goto L78
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
	v326 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v326 != 0 {
		goto L86
	} else {
		goto L87
	}
L84:
	;
	goto L85
L85:
	;
	goto L74
L86:
	;
	v328 = v326
	goto L88
L87:
	;
	v328 = int32(5)
	goto L88
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v328
	goto L85
L89:
	;
	v414 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v414 != 0 {
		goto L104
	} else {
		goto L105
	}
L90:
	;
	if base.Ui32(v403) < base.Ui32(int32(1)) {
		goto L98
	} else {
		goto L99
	}
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v366
	v403 = v368
	v404 = v369
	goto L90
L92:
	;
	if base.Ui32(v361) <= base.Ui32(v366) {
		v403 = v368
		v404 = v369
		goto L90
	} else {
		goto L94
	}
L93:
	;
	v403 = v353
	v404 = v395
	goto L90
L94:
	;
	v374 = v366 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v374
	v376 = *(*int32)(unsafe.Add(mBase, uint32(v366)))
	v378 = v376 - int32(48)
	if base.B2i32(base.Ui32(int32(54)) < base.Ui32(v378))|base.B2i32(base.I32_wrap_i64(int64(base.Ui64(int64(35465847073801215))>>(uint(base.I64_extend_i32_u(v378))%64)))&int32(1) == int32(0)) != 0 {
		goto L91
	} else {
		goto L95
	}
L95:
	;
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v378<<(uint(int32(2))%32))+uint32(_c_F_lexescape[0])))
	if base.Ui32(v351) <= base.Ui32(v392) {
		goto L91
	} else {
		goto L96
	}
L96:
	;
	v395 = v392 + v351*v369
	v397 = v368 + int32(1)
	if v397 != v353 {
		v366 = v374
		v368 = v397
		v369 = v395
		goto L92
	} else {
		goto L97
	}
L97:
	;
	goto L93
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
	v409 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v409 != 0 {
		goto L101
	} else {
		goto L102
	}
L99:
	;
	goto L100
L100:
	;
	goto L89
L101:
	;
	v411 = v409
	goto L103
L102:
	;
	v411 = int32(5)
	goto L103
L103:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v411
	goto L100
L104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
	return int32(0)
L105:
	;
	goto L106
L106:
	;
	v419 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v347 != v419 {
		goto L108
	} else {
		goto L109
	}
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v347
	goto L8
L108:
	;
	if v404 <= int32(0) {
		goto L107
	} else {
		goto L111
	}
L109:
	;
	goto L110
L110:
	;
	v425 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v426 = *(*int32)(unsafe.Add(mBase, uint32(v425)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v425)+8)) = v426 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v404
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(98)
	goto L1
L111:
	;
	v423 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v423 < v404 {
		goto L107
	} else {
		goto L112
	}
L112:
	;
	goto L110
L113:
	;
	v508 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v508 != 0 {
		goto L128
	} else {
		goto L129
	}
L114:
	;
	if base.Ui32(v497) < base.Ui32(int32(1)) {
		goto L122
	} else {
		goto L123
	}
L115:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v460
	v497 = v462
	v498 = v463
	goto L114
L116:
	;
	if base.Ui32(v455) <= base.Ui32(v460) {
		v497 = v462
		v498 = v463
		goto L114
	} else {
		goto L118
	}
L117:
	;
	v497 = v447
	v498 = v489
	goto L114
L118:
	;
	v468 = v460 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v468
	v470 = *(*int32)(unsafe.Add(mBase, uint32(v460)))
	v472 = v470 - int32(48)
	if base.B2i32(base.Ui32(int32(54)) < base.Ui32(v472))|base.B2i32(base.I32_wrap_i64(int64(base.Ui64(int64(35465847073801215))>>(uint(base.I64_extend_i32_u(v472))%64)))&int32(1) == int32(0)) != 0 {
		goto L115
	} else {
		goto L119
	}
L119:
	;
	v486 = *(*int32)(unsafe.Add(mBase, uint32(v472<<(uint(int32(2))%32))+uint32(_c_F_lexescape[0])))
	if base.Ui32(v445) <= base.Ui32(v486) {
		goto L115
	} else {
		goto L120
	}
L120:
	;
	v489 = v486 + v445*v463
	v491 = v462 + int32(1)
	if v491 != v447 {
		v460 = v468
		v462 = v491
		v463 = v489
		goto L116
	} else {
		goto L121
	}
L121:
	;
	goto L117
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
	v503 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v503 != 0 {
		goto L125
	} else {
		goto L126
	}
L123:
	;
	goto L124
L124:
	;
	goto L113
L125:
	;
	v505 = v503
	goto L127
L126:
	;
	v505 = int32(5)
	goto L127
L127:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v505
	goto L124
L128:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
	return int32(0)
L129:
	;
	goto L130
L130:
	;
	if base.Ui32(int32(256)) <= base.Ui32(v498) {
		goto L131
	} else {
		goto L132
	}
L131:
	;
	v515 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v515 - int32(4)
	v521 = int32(base.Ui32(v498) >> (uint(int32(3)) % 32))
	goto L133
L132:
	;
	v521 = v498
	goto L133
L133:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v521
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(112)
	goto L1
L134:
	;
	v529 = v527
	goto L136
L135:
	;
	v529 = int32(5)
	goto L136
L136:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v529
	return int32(0)
L137:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
	if v534 != 0 {
		goto L140
	} else {
		goto L141
	}
L138:
	;
	goto L139
L139:
	;
	v549 = v533
	goto L2
L140:
	;
	v545 = v534
	goto L142
L141:
	;
	v545 = int32(5)
	goto L142
L142:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v545
	return int32(0)
}
func F_lithuanian_UTF_8_stem(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v16 int32
	_ = v16
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v97 int32
	_ = v97
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v123 int32
	_ = v123
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v174 int32
	_ = v174
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v200 int32
	_ = v200
	var v207 int32
	_ = v207
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v245 int32
	_ = v245
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v267 int32
	_ = v267
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v283 int32
	_ = v283
	var v296 int32
	_ = v296
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v316 int32
	_ = v316
	var v322 int32
	_ = v322
	var v330 int32
	_ = v330
	var v341 int32
	_ = v341
	var v344 int32
	_ = v344
	var v349 int32
	_ = v349
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v374 int32
	_ = v374
	var v377 int32
	_ = v377
	var v380 int32
	_ = v380
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v433 int32
	_ = v433
	var v435 int32
	_ = v435
	var v438 int32
	_ = v438
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v449 int32
	_ = v449
	var v451 int32
	_ = v451
	var v455 int32
	_ = v455
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v461 int32
	_ = v461
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v470 int32
	_ = v470
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v479 int32
	_ = v479
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v485 int32
	_ = v485
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v495 int32
	_ = v495
	var v498 int32
	_ = v498
	var v501 int32
	_ = v501
	var v505 int32
	_ = v505
	var v508 int32
	_ = v508
	var v510 int32
	_ = v510
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v517 int32
	_ = v517
	var v519 int32
	_ = v519
	var v522 int32
	_ = v522
	var v526 int32
	_ = v526
	var v528 int32
	_ = v528
	var v532 int32
	_ = v532
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v4
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v4 == v6 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v114 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v115 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v123 = v113
	goto L24
L2:
	;
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8+v6))))
	if v10 != int32(97) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v6 + int32(1)
	v16 = int32(0)
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v8-int32(4))))
	if v23 == v16 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	if int32(6) < v97 {
		goto L1
	} else {
		goto L20
	}
L5:
	;
	v97 = int32(0)
	goto L4
L6:
	;
	goto L7
L7:
	;
	v28 = v23 & int32(3)
	if base.Ui32(v23) < base.Ui32(int32(4)) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v97 = v86
	goto L4
L9:
	;
	v70 = v64
	v71 = v65
	v75 = v16
	goto L17
L10:
	;
	v64 = v8
	v65 = int32(0)
	goto L9
L11:
	;
	goto L12
L12:
	;
	v35 = v8
	v36 = int32(0)
	v39 = v16
	goto L13
L13:
	;
	v41 = int32(*(*int8)(unsafe.Add(mBase, uint32(v35))))
	v42 = int32(-65)
	v45 = int32(*(*int8)(unsafe.Add(mBase, uint32(v35)+1)))
	v49 = int32(*(*int8)(unsafe.Add(mBase, uint32(v35)+2)))
	v53 = int32(*(*int8)(unsafe.Add(mBase, uint32(v35)+3)))
	v56 = v36 + base.B2i32(v42 < v41) + base.B2i32(v42 < v45) + base.B2i32(v42 < v49) + base.B2i32(v42 < v53)
	v57 = int32(4)
	v58 = v35 + v57
	v60 = v39 + v57
	if v60 != v23&int32(-4) {
		v35 = v58
		v36 = v56
		v39 = v60
		goto L13
	} else {
		goto L15
	}
L14:
	;
	if v28 == int32(0) {
		v86 = v56
		goto L8
	} else {
		goto L16
	}
L15:
	;
	goto L14
L16:
	;
	v64 = v58
	v65 = v56
	goto L9
L17:
	;
	v76 = int32(*(*int8)(unsafe.Add(mBase, uint32(v70))))
	v79 = v71 + base.B2i32(int32(-65) < v76)
	v80 = int32(1)
	v83 = v75 + v80
	if v83 != v28 {
		v70 = v70 + v80
		v71 = v79
		v75 = v83
		goto L17
	} else {
		goto L19
	}
L18:
	;
	v86 = v79
	goto L8
L19:
	;
	goto L18
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v6
	goto L1
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v6
	v349 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v349
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v349
	if v349-int32(3) <= v6 {
		goto L74
	} else {
		goto L75
	}
L22:
	;
	if v218 < int32(0) {
		goto L21
	} else {
		goto L47
	}
L23:
	;
	v218 = v190
	goto L22
L24:
	;
	if v114 <= v123 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v218 = int32(-1)
	goto L22
L27:
	;
	goto L28
L28:
	;
	v130 = int32(1)
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123+v115))))
	if base.Ui32(v132) < base.Ui32(int32(192)) {
		v189 = v132
		v190 = v130
		goto L29
	} else {
		goto L30
	}
L29:
	;
	if int32(371) < v189 {
		goto L42
	} else {
		goto L43
	}
L30:
	;
	v136 = v123 + int32(1)
	if v136 == v114 {
		v189 = v132
		v190 = v130
		goto L29
	} else {
		goto L31
	}
L31:
	;
	v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v136+v115))))
	v141 = v139 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v132) {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v145+v115))))
	v157 = v155 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v132) {
		goto L38
	} else {
		goto L39
	}
L33:
	;
	v145 = v123 + int32(2)
	if v145 != v114 {
		goto L32
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	v189 = v132<<(uint(int32(6))%32)&int32(1984) | v141
	v190 = int32(2)
	goto L29
L36:
	;
	goto L35
L37:
	;
	v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115+v161))))
	v189 = v174&int32(63) | (v132<<(uint(int32(18))%32)&int32(_a_F_lithuanian_UTF_8_stem_0) | v141<<(uint(int32(12))%32) | v157<<(uint(int32(6))%32))
	v190 = int32(4)
	goto L29
L38:
	;
	v161 = v123 + int32(3)
	if v161 != v114 {
		goto L37
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	v189 = v132<<(uint(int32(12))%32)&int32(_a_F_lithuanian_UTF_8_stem_1) | v141<<(uint(int32(6))%32) | v157
	v190 = int32(3)
	goto L29
L41:
	;
	goto L40
L42:
	;
	v207 = v190 + v123
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v207
	v123 = v207
	goto L24
L43:
	;
	v194 = v189 - int32(97)
	if v194 < int32(0) {
		goto L42
	} else {
		goto L44
	}
L44:
	;
	v200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v194)>>(uint(int32(3))%32)))+uint32(_c_F_lithuanian_UTF_8_stem[0]))))
	if int32(base.Ui32(v200)>>(uint(v194&int32(7))%32))&int32(1) != 0 {
		goto L23
	} else {
		goto L45
	}
L45:
	;
	goto L42
L47:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v222 = v221 + v218
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v222
	v236 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v237 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v245 = v222
	goto L50
L48:
	;
	if v341 < int32(0) {
		goto L21
	} else {
		goto L72
	}
L49:
	;
	v341 = v312
	goto L48
L50:
	;
	if v236 <= v245 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v341 = int32(-1)
	goto L48
L53:
	;
	goto L54
L54:
	;
	v252 = int32(1)
	v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v245+v237))))
	if base.Ui32(v254) < base.Ui32(int32(192)) {
		v311 = v254
		v312 = v252
		goto L55
	} else {
		goto L56
	}
L55:
	;
	if int32(371) < v311 {
		goto L49
	} else {
		goto L68
	}
L56:
	;
	v258 = v245 + int32(1)
	if v258 == v236 {
		v311 = v254
		v312 = v252
		goto L55
	} else {
		goto L57
	}
L57:
	;
	v261 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v258+v237))))
	v263 = v261 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v254) {
		goto L59
	} else {
		goto L60
	}
L58:
	;
	v277 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v267+v237))))
	v279 = v277 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v254) {
		goto L64
	} else {
		goto L65
	}
L59:
	;
	v267 = v245 + int32(2)
	if v267 != v236 {
		goto L58
	} else {
		goto L62
	}
L60:
	;
	goto L61
L61:
	;
	v311 = v254<<(uint(int32(6))%32)&int32(1984) | v263
	v312 = int32(2)
	goto L55
L62:
	;
	goto L61
L63:
	;
	v296 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v237+v283))))
	v311 = v296&int32(63) | (v254<<(uint(int32(18))%32)&int32(_a_F_lithuanian_UTF_8_stem_0) | v263<<(uint(int32(12))%32) | v279<<(uint(int32(6))%32))
	v312 = int32(4)
	goto L55
L64:
	;
	v283 = v245 + int32(3)
	if v283 != v236 {
		goto L63
	} else {
		goto L67
	}
L65:
	;
	goto L66
L66:
	;
	v311 = v254<<(uint(int32(12))%32)&int32(_a_F_lithuanian_UTF_8_stem_1) | v263<<(uint(int32(6))%32) | v279
	v312 = int32(3)
	goto L55
L67:
	;
	goto L66
L68:
	;
	v316 = v311 - int32(97)
	if v316 < int32(0) {
		goto L49
	} else {
		goto L69
	}
L69:
	;
	v322 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v316)>>(uint(int32(3))%32)))+uint32(_c_F_lithuanian_UTF_8_stem[0]))))
	if int32(base.Ui32(v322)>>(uint(v316&int32(7))%32))&int32(1) == int32(0) {
		goto L49
	} else {
		goto L70
	}
L70:
	;
	v330 = v312 + v245
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v330
	v245 = v330
	goto L50
L72:
	;
	v344 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v344 + v341
	goto L21
L73:
	;
	return v532
L74:
	;
	v433 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v433
	v435 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v433 < v435 {
		goto L104
	} else {
		goto L105
	}
L75:
	;
	v355 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v357 = int32(1)
	v359 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v355+v349-v357))))
	if base.B2i32(v359&int32(224) != int32(96))|base.B2i32(v357<<(uint(v359)%32)&int32(_a_F_lithuanian_UTF_8_stem_2) == int32(0)) != 0 {
		goto L74
	} else {
		goto L76
	}
L76:
	;
	v374 = F_find_among_b(m, l0, int32(_a_F_lithuanian_UTF_8_stem_3), int32(11), int32(0))
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	return int32(0)
L78:
	;
	if v374 == int32(0) {
		goto L74
	} else {
		goto L79
	}
L79:
	;
	v380 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v380
	switch v374 - int32(1) {
	case 0:
		goto L87
	case 1:
		goto L86
	case 2:
		goto L85
	case 3:
		goto L84
	case 4:
		goto L83
	case 5:
		goto L82
	case 6:
		goto L81
	case 7:
		goto L80
	default:
		goto L74
	}
L80:
	;
	v428 = F_slice_from_s(m, l0, int32(6), int32(_a_F_lithuanian_UTF_8_stem_4))
	mBase = m.M
	v429 = m.ExcPending
	if v429 != 0 {
		goto L77
	} else {
		goto L102
	}
L81:
	;
	v422 = F_slice_from_s(m, l0, int32(6), int32(_a_F_lithuanian_UTF_8_stem_5))
	mBase = m.M
	v423 = m.ExcPending
	if v423 != 0 {
		goto L77
	} else {
		goto L100
	}
L82:
	;
	v416 = F_slice_from_s(m, l0, int32(6), int32(_a_F_lithuanian_UTF_8_stem_6))
	mBase = m.M
	v417 = m.ExcPending
	if v417 != 0 {
		goto L77
	} else {
		goto L98
	}
L83:
	;
	v410 = F_slice_from_s(m, l0, int32(4), int32(_a_F_lithuanian_UTF_8_stem_7))
	mBase = m.M
	v411 = m.ExcPending
	if v411 != 0 {
		goto L77
	} else {
		goto L96
	}
L84:
	;
	v404 = F_slice_from_s(m, l0, int32(4), int32(_a_F_lithuanian_UTF_8_stem_8))
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
		goto L77
	} else {
		goto L94
	}
L85:
	;
	v398 = F_slice_from_s(m, l0, int32(7), int32(_a_F_lithuanian_UTF_8_stem_9))
	mBase = m.M
	v399 = m.ExcPending
	if v399 != 0 {
		goto L77
	} else {
		goto L92
	}
L86:
	;
	v392 = F_slice_from_s(m, l0, int32(5), int32(_a_F_lithuanian_UTF_8_stem_10))
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L77
	} else {
		goto L90
	}
L87:
	;
	v386 = F_slice_from_s(m, l0, int32(5), int32(_a_F_lithuanian_UTF_8_stem_11))
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L77
	} else {
		goto L88
	}
L88:
	;
	if int32(0) <= v386 {
		goto L74
	} else {
		goto L89
	}
L89:
	;
	v532 = v386
	goto L73
L90:
	;
	if int32(0) <= v392 {
		goto L74
	} else {
		goto L91
	}
L91:
	;
	v532 = v392
	goto L73
L92:
	;
	if int32(0) <= v398 {
		goto L74
	} else {
		goto L93
	}
L93:
	;
	v532 = v398
	goto L73
L94:
	;
	if int32(0) <= v404 {
		goto L74
	} else {
		goto L95
	}
L95:
	;
	v532 = v404
	goto L73
L96:
	;
	if int32(0) <= v410 {
		goto L74
	} else {
		goto L97
	}
L97:
	;
	v532 = v410
	goto L73
L98:
	;
	if int32(0) <= v416 {
		goto L74
	} else {
		goto L99
	}
L99:
	;
	v532 = v416
	goto L73
L100:
	;
	if int32(0) <= v422 {
		goto L74
	} else {
		goto L101
	}
L101:
	;
	v532 = v422
	goto L73
L102:
	;
	if v428 < int32(0) {
		v532 = v428
		goto L73
	} else {
		goto L103
	}
L103:
	;
	goto L74
L104:
	;
	v455 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v455
	v457 = F_r_fix_chdz(m, l0)
	mBase = m.M
	v458 = m.ExcPending
	if v458 != 0 {
		goto L77
	} else {
		goto L111
	}
L105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v433
	v438 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v435
	v443 = F_find_among_b(m, l0, int32(_a_F_lithuanian_UTF_8_stem_12), int32(193), int32(0))
	mBase = m.M
	v444 = m.ExcPending
	if v444 != 0 {
		goto L77
	} else {
		goto L106
	}
L106:
	;
	if v443 == int32(0) {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v438
	goto L104
L108:
	;
	goto L109
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v438
	v449 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v449
	v451 = F_slice_del(m, l0)
	mBase = m.M
	if v451 < int32(0) {
		v532 = v451
		goto L73
	} else {
		goto L110
	}
L110:
	;
	goto L104
L111:
	;
	if v457 < int32(0) {
		v532 = v457
		goto L73
	} else {
		goto L112
	}
L112:
	;
	v461 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v461
	goto L113
L113:
	;
	v466 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v467 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v467 <= v466 {
		goto L116
	} else {
		goto L117
	}
L114:
	;
	v532 = v528
	goto L73
L115:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v470
	v526 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v526
	v528 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v528 {
		goto L113
	} else {
		goto L134
	}
L116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v466
	v470 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v467
	v475 = F_find_among_b(m, l0, int32(_a_F_lithuanian_UTF_8_stem_13), int32(62), int32(0))
	mBase = m.M
	v476 = m.ExcPending
	if v476 != 0 {
		goto L77
	} else {
		goto L119
	}
L117:
	;
	goto L118
L118:
	;
	v479 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v479
	v481 = F_r_fix_chdz(m, l0)
	mBase = m.M
	v482 = m.ExcPending
	if v482 != 0 {
		goto L77
	} else {
		goto L121
	}
L119:
	;
	if v475 != 0 {
		goto L115
	} else {
		goto L120
	}
L120:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v470
	goto L118
L121:
	;
	if v481 < int32(0) {
		v532 = v481
		goto L73
	} else {
		goto L122
	}
L122:
	;
	v485 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v485
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v485
	v489 = int32(0)
	v490 = int32(2)
	v495 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v485-v495 < v490 {
		v505 = v489
		goto L125
	} else {
		goto L126
	}
L123:
	;
	if v519 < int32(0) {
		v532 = v519
		goto L73
	} else {
		goto L133
	}
L124:
	;
	if v505 == int32(0) {
		v519 = v489
		goto L123
	} else {
		goto L128
	}
L125:
	;
	goto L124
L126:
	;
	v498 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v501 = F_memcmp(m, v498+v485-v490, int32(_a_F_lithuanian_UTF_8_stem_14), v490)
	mBase = m.M
	if v501 != 0 {
		v505 = v489
		goto L125
	} else {
		goto L127
	}
L127:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v485 - v490
	v505 = int32(1)
	goto L125
L128:
	;
	v508 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v508
	v510 = int32(1)
	v513 = F_slice_from_s(m, l0, v510, int32(_a_F_lithuanian_UTF_8_stem_15))
	mBase = m.M
	v514 = m.ExcPending
	if v514 != 0 {
		goto L77
	} else {
		goto L129
	}
L129:
	;
	if int32(0) <= v513 {
		goto L130
	} else {
		goto L131
	}
L130:
	;
	v517 = v510
	goto L132
L131:
	;
	v517 = v513
	goto L132
L132:
	;
	v519 = v517
	goto L123
L133:
	;
	v522 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v522
	v532 = int32(1)
	goto L73
L134:
	;
	goto L114
}
func F_lock_twophase_recover(m *base.Module, l0 int64, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v65 int64
	_ = v65
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v99 int32
	_ = v99
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v142 int32
	_ = v142
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v265 int32
	_ = v265
	var v272 int32
	_ = v272
	var v276 int32
	_ = v276
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v294 int32
	_ = v294
	var v301 int32
	_ = v301
	var v306 int32
	_ = v306
	var v310 int32
	_ = v310
	var v314 int32
	_ = v314
	var v319 int32
	_ = v319
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v326 int32
	_ = v326
	var v327 int64
	_ = v327
	var v328 int32
	_ = v328
	var v336 int32
	_ = v336
	var v341 int32
	_ = v341
	v13 = m.G0
	v15 = v13 - int32(80)
	m.G0 = v15
	v18 = F_TwoPhaseGetDummyProc(m, l0, int32(0))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return
	} else {
		v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+15)))
		if base.Ui32(int32(253)) < base.Ui32((v20-int32(3))&int32(255)) {
			v27 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
			v29 = *(*int32)(unsafe.Add(mBase, _c_F_lock_twophase_recover[0]))
			v30 = F_get_hash_value(m, v29, l2)
			mBase = m.M
			v31 = m.ExcPending
			if v31 != 0 {
				return
			} else {
				v33 = *(*int32)(unsafe.Add(mBase, _c_F_lock_twophase_recover[1]))
				v35 = v30 & int32(15)
				v40 = v33 + v35<<(uint(int32(7))%32) + int32(_a_F_lock_twophase_recover_0)
				v42 = F_LWLockAcquire(m, v40, int32(0))
				mBase = m.M
				v43 = m.ExcPending
				if v43 != 0 {
					return
				} else {
					v46 = *(*int32)(unsafe.Add(mBase, uint32(v20<<(uint(int32(2))%32))+uint32(_c_F_lock_twophase_recover[2])))
					v48 = *(*int32)(unsafe.Add(mBase, _c_F_lock_twophase_recover[0]))
					v52 = F_hash_search_with_hash_value(m, v48, l2, v30, int32(3), v15+int32(71))
					mBase = m.M
					v53 = m.ExcPending
					if v53 != 0 {
						return
					} else {
						if v52 == int32(0) {
							F_LWLockRelease(m, v40)
							mBase = m.M
							v283 = m.ExcPending
							if v283 != 0 {
								return
							} else {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v287 = m.ExcPending
								if v287 != 0 {
									return
								} else {
									F_errcode(m, int32(_a_F_lock_twophase_recover_1))
									mBase = m.M
									v290 = m.ExcPending
									if v290 != 0 {
										return
									} else {
										F_errmsg(m, int32(_a_F_lock_twophase_recover_2), int32(0))
										mBase = m.M
										v294 = m.ExcPending
										if v294 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = int32(_a_F_lock_twophase_recover_3)
											F_errhint(m, int32(_a_F_lock_twophase_recover_4), v15+int32(16))
											mBase = m.M
											v301 = m.ExcPending
											if v301 != 0 {
												return
											} else {
												F_errfinish(m, int32(_a_F_lock_twophase_recover_5), int32(_a_F_lock_twophase_recover_6), int32(_a_F_lock_twophase_recover_7))
												mBase = m.M
												v306 = m.ExcPending
												if v306 != 0 {
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
							v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+71)))
							if v56 == int32(0) {
								v59 = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v52)+128)) = v59
								*(*int32)(unsafe.Add(mBase, uint32(v52)+84)) = v59
								*(*int32)(unsafe.Add(mBase, uint32(v52)+40)) = v59
								v65 = int64(0)
								*(*int64)(unsafe.Add(mBase, uint32(v52)+16)) = v65
								v68 = v52 + int32(32)
								*(*int32)(unsafe.Add(mBase, uint32(v52)+36)) = v68
								*(*int32)(unsafe.Add(mBase, uint32(v52)+32)) = v68
								v72 = v52 + int32(24)
								*(*int32)(unsafe.Add(mBase, uint32(v52)+28)) = v72
								*(*int32)(unsafe.Add(mBase, uint32(v52)+24)) = v72
								*(*int64)(unsafe.Add(mBase, uint32(v52)+44)) = v65
								*(*int64)(unsafe.Add(mBase, uint32(v52)+52)) = v65
								*(*int64)(unsafe.Add(mBase, uint32(v52)+60)) = v65
								*(*int64)(unsafe.Add(mBase, uint32(v52)+68)) = v65
								*(*int64)(unsafe.Add(mBase, uint32(v52)+76)) = v65
								*(*int64)(unsafe.Add(mBase, uint32(v52)+88)) = v65
								*(*int64)(unsafe.Add(mBase, uint32(v52)+96)) = v65
								*(*int64)(unsafe.Add(mBase, uint32(v52)+104)) = v65
								*(*int64)(unsafe.Add(mBase, uint32(v52)+112)) = v65
								*(*int64)(unsafe.Add(mBase, uint32(v52)+120)) = v65
							} else {
							}
							*(*int32)(unsafe.Add(mBase, uint32(v15)+72)) = v52
							*(*int32)(unsafe.Add(mBase, uint32(v15)+76)) = v18
							v99 = *(*int32)(unsafe.Add(mBase, _c_F_lock_twophase_recover[3]))
							v108 = F_hash_search_with_hash_value(m, v99, v15+int32(72), v30^v18<<(uint(int32(4))%32), int32(3), v15+int32(71))
							mBase = m.M
							v109 = m.ExcPending
							if v109 != 0 {
								return
							} else {
								if v108 == int32(0) {
									v112 = *(*int32)(unsafe.Add(mBase, uint32(v52)+84))
									if v112 == int32(0) {
										v116 = *(*int32)(unsafe.Add(mBase, _c_F_lock_twophase_recover[0]))
										v119 = F_hash_search_with_hash_value(m, v116, v52, v30, int32(2), int32(0))
										mBase = m.M
										v120 = m.ExcPending
										if v120 != 0 {
											return
										} else {
											if v119 == int32(0) {
												F_errstart_cold(m, int32(24), int32(0))
												mBase = m.M
												v310 = m.ExcPending
												if v310 != 0 {
													return
												} else {
													F_errmsg_internal(m, int32(_a_F_lock_twophase_recover_8), int32(0))
													mBase = m.M
													v314 = m.ExcPending
													if v314 != 0 {
														return
													} else {
														F_errfinish(m, int32(_a_F_lock_twophase_recover_5), int32(_a_F_lock_twophase_recover_9), int32(_a_F_lock_twophase_recover_7))
														mBase = m.M
														v319 = m.ExcPending
														if v319 != 0 {
															return
														} else {
															base.Wasm_trap_unreachable()
															for {
															}
														}
													}
												}
											} else {
												F_LWLockRelease(m, v40)
												mBase = m.M
												v124 = m.ExcPending
												if v124 != 0 {
													return
												} else {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v128 = m.ExcPending
													if v128 != 0 {
														return
													} else {
														F_errcode(m, int32(_a_F_lock_twophase_recover_1))
														mBase = m.M
														v131 = m.ExcPending
														if v131 != 0 {
															return
														} else {
															F_errmsg(m, int32(_a_F_lock_twophase_recover_2), int32(0))
															mBase = m.M
															v135 = m.ExcPending
															if v135 != 0 {
																return
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = int32(_a_F_lock_twophase_recover_3)
																F_errhint(m, int32(_a_F_lock_twophase_recover_4), v15+int32(32))
																mBase = m.M
																v142 = m.ExcPending
																if v142 != 0 {
																	return
																} else {
																	F_errfinish(m, int32(_a_F_lock_twophase_recover_5), int32(_a_F_lock_twophase_recover_10), int32(_a_F_lock_twophase_recover_7))
																	mBase = m.M
																	v147 = m.ExcPending
																	if v147 != 0 {
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
										F_LWLockRelease(m, v40)
										mBase = m.M
										v124 = m.ExcPending
										if v124 != 0 {
											return
										} else {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v128 = m.ExcPending
											if v128 != 0 {
												return
											} else {
												F_errcode(m, int32(_a_F_lock_twophase_recover_1))
												mBase = m.M
												v131 = m.ExcPending
												if v131 != 0 {
													return
												} else {
													F_errmsg(m, int32(_a_F_lock_twophase_recover_2), int32(0))
													mBase = m.M
													v135 = m.ExcPending
													if v135 != 0 {
														return
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = int32(_a_F_lock_twophase_recover_3)
														F_errhint(m, int32(_a_F_lock_twophase_recover_4), v15+int32(32))
														mBase = m.M
														v142 = m.ExcPending
														if v142 != 0 {
															return
														} else {
															F_errfinish(m, int32(_a_F_lock_twophase_recover_5), int32(_a_F_lock_twophase_recover_10), int32(_a_F_lock_twophase_recover_7))
															mBase = m.M
															v147 = m.ExcPending
															if v147 != 0 {
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
									v148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+71)))
									if v148 == int32(0) {
										*(*int64)(unsafe.Add(mBase, uint32(v108)+12)) = int64(0)
										*(*int32)(unsafe.Add(mBase, uint32(v108)+8)) = v18
										v155 = v52 + int32(24)
										v156 = *(*int32)(unsafe.Add(mBase, uint32(v52)+28))
										if v156 == int32(0) {
											*(*int32)(unsafe.Add(mBase, uint32(v52)+28)) = v155
											*(*int32)(unsafe.Add(mBase, uint32(v52)+24)) = v155
										} else {
										}
										*(*int32)(unsafe.Add(mBase, uint32(v108)+24)) = v155
										v162 = *(*int32)(unsafe.Add(mBase, uint32(v155)))
										*(*int32)(unsafe.Add(mBase, uint32(v108)+20)) = v162
										v165 = v108 + int32(20)
										*(*int32)(unsafe.Add(mBase, uint32(v162)+4)) = v165
										*(*int32)(unsafe.Add(mBase, uint32(v155))) = v165
										v170 = v18 + v35<<(uint(int32(3))%32)
										v172 = v170 + int32(420)
										v173 = *(*int32)(unsafe.Add(mBase, uint32(v170)+424))
										if v173 == int32(0) {
											*(*int32)(unsafe.Add(mBase, uint32(v172)+4)) = v172
											*(*int32)(unsafe.Add(mBase, uint32(v172))) = v172
										} else {
										}
										*(*int32)(unsafe.Add(mBase, uint32(v108)+32)) = v172
										v179 = *(*int32)(unsafe.Add(mBase, uint32(v172)))
										*(*int32)(unsafe.Add(mBase, uint32(v108)+28)) = v179
										v182 = v108 + int32(28)
										*(*int32)(unsafe.Add(mBase, uint32(v179)+4)) = v182
										*(*int32)(unsafe.Add(mBase, uint32(v172))) = v182
									} else {
									}
									v188 = *(*int32)(unsafe.Add(mBase, uint32(v52)+84))
									v189 = int32(1)
									*(*int32)(unsafe.Add(mBase, uint32(v52)+84)) = v188 + v189
									v193 = v27 << (uint(int32(2)) % 32)
									v194 = v52 + v193
									v196 = v194 + int32(44)
									v197 = *(*int32)(unsafe.Add(mBase, uint32(v196)))
									*(*int32)(unsafe.Add(mBase, uint32(v196))) = v197 + v189
									v202 = v189 << (uint(v27) % 32)
									v203 = *(*int32)(unsafe.Add(mBase, uint32(v108)+12))
									if v202&v203 != 0 {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v323 = m.ExcPending
										if v323 != 0 {
											return
										} else {
											v324 = *(*int32)(unsafe.Add(mBase, uint32(v46)+8))
											v326 = *(*int32)(unsafe.Add(mBase, uint32(v324+v193)))
											v327 = *(*int64)(unsafe.Add(mBase, uint32(v52)))
											v328 = *(*int32)(unsafe.Add(mBase, uint32(v52)+8))
											*(*int32)(unsafe.Add(mBase, uint32(v15)+60)) = v328
											*(*int64)(unsafe.Add(mBase, uint32(v15)+52)) = v327
											*(*int32)(unsafe.Add(mBase, uint32(v15)+48)) = v326
											F_errmsg_internal(m, int32(_a_F_lock_twophase_recover_11), v15+int32(48))
											mBase = m.M
											v336 = m.ExcPending
											if v336 != 0 {
												return
											} else {
												F_errfinish(m, int32(_a_F_lock_twophase_recover_5), int32(_a_F_lock_twophase_recover_12), int32(_a_F_lock_twophase_recover_7))
												mBase = m.M
												v341 = m.ExcPending
												if v341 != 0 {
													return
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											}
										}
									} else {
										v205 = *(*int32)(unsafe.Add(mBase, uint32(v52)+128))
										v206 = int32(1)
										*(*int32)(unsafe.Add(mBase, uint32(v52)+128)) = v205 + v206
										v210 = v194 + int32(88)
										v211 = *(*int32)(unsafe.Add(mBase, uint32(v210)))
										*(*int32)(unsafe.Add(mBase, uint32(v210))) = v211 + v206
										v215 = *(*int32)(unsafe.Add(mBase, uint32(v52)+16))
										*(*int32)(unsafe.Add(mBase, uint32(v52)+16)) = v215 | v202
										v218 = *(*int32)(unsafe.Add(mBase, uint32(v210)))
										v219 = *(*int32)(unsafe.Add(mBase, uint32(v196)))
										if v218 == v219 {
											v221 = *(*int32)(unsafe.Add(mBase, uint32(v52)+20))
											*(*int32)(unsafe.Add(mBase, uint32(v52)+20)) = v221 & (v202 ^ int32(-1))
										} else {
										}
										v226 = *(*int32)(unsafe.Add(mBase, uint32(v108)+12))
										*(*int32)(unsafe.Add(mBase, uint32(v108)+12)) = v226 | v202
										v229 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+15)))
										if v229 != int32(1) {
											F_LWLockRelease(m, v40)
											mBase = m.M
											v265 = m.ExcPending
											if v265 != 0 {
												return
											} else {
												m.G0 = v15 + int32(80)
												return
											}
										} else {
											v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+14)))
											if v232 != 0 {
												F_LWLockRelease(m, v40)
												mBase = m.M
												v265 = m.ExcPending
												if v265 != 0 {
													return
												} else {
													m.G0 = v15 + int32(80)
													return
												}
											} else {
												v233 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
												if base.B2i32(v233 == int32(0))|base.B2i32(v27 < int32(5)) != 0 {
													F_LWLockRelease(m, v40)
													mBase = m.M
													v265 = m.ExcPending
													if v265 != 0 {
														return
													} else {
														m.G0 = v15 + int32(80)
														return
													}
												} else {
													v240 = *(*int32)(unsafe.Add(mBase, _c_F_lock_twophase_recover[4]))
													v243 = base.AtomicRmwXchg32(m, v240, int32(0), int32(1))
													if v243 != 0 {
														F_s_lock(m, v240, int32(_a_F_lock_twophase_recover_13))
														mBase = m.M
														v246 = m.ExcPending
														if v246 != 0 {
															return
														} else {
															v247 = int32(_a_F_lock_twophase_recover_14)
															v248 = *(*int32)(unsafe.Add(mBase, _c_F_lock_twophase_recover[4]))
															v253 = v248 + v30&int32(1023)<<(uint(int32(2))%32)
															v254 = *(*int32)(unsafe.Add(mBase, uint32(v253)+4))
															*(*int32)(unsafe.Add(mBase, uint32(v253)+4)) = v254 + int32(1)
															v259 = *(*int32)(unsafe.Add(mBase, _c_F_lock_twophase_recover[4]))
															v260 = int32(0)
															atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v259))), uint32(v260))
															F_LWLockRelease(m, v40)
															mBase = m.M
															v265 = m.ExcPending
															if v265 != 0 {
																return
															} else {
																m.G0 = v15 + int32(80)
																return
															}
														}
													} else {
														v247 = int32(_a_F_lock_twophase_recover_14)
														v248 = *(*int32)(unsafe.Add(mBase, _c_F_lock_twophase_recover[4]))
														v253 = v248 + v30&int32(1023)<<(uint(int32(2))%32)
														v254 = *(*int32)(unsafe.Add(mBase, uint32(v253)+4))
														*(*int32)(unsafe.Add(mBase, uint32(v253)+4)) = v254 + int32(1)
														v259 = *(*int32)(unsafe.Add(mBase, _c_F_lock_twophase_recover[4]))
														v260 = int32(0)
														atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v259))), uint32(v260))
														F_LWLockRelease(m, v40)
														mBase = m.M
														v265 = m.ExcPending
														if v265 != 0 {
															return
														} else {
															m.G0 = v15 + int32(80)
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
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v272 = m.ExcPending
			if v272 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v15))) = v20
				F_errmsg_internal(m, int32(_a_F_lock_twophase_recover_15), v15)
				mBase = m.M
				v276 = m.ExcPending
				if v276 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_lock_twophase_recover_5), int32(_a_F_lock_twophase_recover_16), int32(_a_F_lock_twophase_recover_7))
					mBase = m.M
					v281 = m.ExcPending
					if v281 != 0 {
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
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
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
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v139 int32
	_ = v139
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
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
	var v182 int32
	_ = v182
	var v193 int32
	_ = v193
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v252 int32
	_ = v252
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v293 int32
	_ = v293
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v302 int32
	_ = v302
	var v307 int32
	_ = v307
	var v318 int32
	_ = v318
	var v325 int32
	_ = v325
	var v333 int32
	_ = v333
	var v344 int32
	_ = v344
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v358 int32
	_ = v358
	var v379 int32
	_ = v379
	var v382 int32
	_ = v382
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v388 int32
	_ = v388
	var v393 int32
	_ = v393
	var v396 int32
	_ = v396
	var v399 int32
	_ = v399
	var v403 int32
	_ = v403
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
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v432 int32
	_ = v432
	var v435 int32
	_ = v435
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v448 int32
	_ = v448
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v455 int32
	_ = v455
	var v463 int32
	_ = v463
	var v465 int32
	_ = v465
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v490 int32
	_ = v490
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v500 int32
	_ = v500
	var v503 int32
	_ = v503
	var v508 int32
	_ = v508
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v514 int32
	_ = v514
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v530 int32
	_ = v530
	var v537 int32
	_ = v537
	var v542 int32
	_ = v542
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v554 int32
	_ = v554
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v572 int32
	_ = v572
	var v582 int32
	_ = v582
	var v591 int32
	_ = v591
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
	return v591
L5:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117)+8)))
	if v118&int32(2) != 0 {
		goto L35
	} else {
		goto L36
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
	v28 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+66)))
	v29 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+64)))
	if v29 <= v28 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	if l4 == int32(0) {
		v591 = v105
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
	if base.Ui32(v39) < base.Ui32(v28) {
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
	v41 = v28
	goto L15
L15:
	;
	if v28 == int32(256) {
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
	if base.Ui32(v44) < base.Ui32(v29) {
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
	if base.Ui32(v29) <= base.Ui32(v83) {
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
	if base.B2i32(l3 != v108)|base.B2i32(v108 != v105) != 0 {
		v591 = v105
		goto L4
	} else {
		goto L34
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = int32(1)
	return v105
L35:
	;
	v124 = (l3 - l2) >> (uint(int32(2)) % 32)
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v117)+40))
	if base.Ui32(v124) < base.Ui32(v125) {
		v591 = int32(0)
		goto L4
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	v157 = int32(0)
	v166 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v157 < v166 {
		goto L57
	} else {
		goto L58
	}
L38:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v117)+44))
	if v128 == int32(256) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	if l4 == int32(0) {
		goto L42
	} else {
		goto L43
	}
L40:
	;
	goto L41
L41:
	;
	v139 = int32(0)
	if base.B2i32(l4 == v139)|base.B2i32(l3 != v127)|base.B2i32(base.Ui32(v128+int32(1)) < base.Ui32(v124)) == v139 {
		goto L48
	} else {
		goto L49
	}
L42:
	;
	return l3
L43:
	;
	goto L44
L44:
	;
	if l3 != v127 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	return l3
L46:
	;
	goto L47
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = int32(1)
	return l3
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = int32(1)
	goto L50
L49:
	;
	goto L50
L50:
	;
	if base.Ui32(v128) < base.Ui32(v124) {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v155 = l2 + v128<<(uint(int32(2))%32)
	goto L53
L52:
	;
	v155 = l3
	goto L53
L53:
	;
	return v155
L54:
	;
	if v379 == int32(0) {
		v591 = v157
		goto L4
	} else {
		goto L91
	}
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v358)+20)) = l2
	*(*int64)(unsafe.Add(mBase, uint32(l1)+48)) = int64(0)
	v379 = v358
	goto L54
L56:
	;
	v333 = int32(0)
	goto L88
L57:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v169)+8)))
	if v170&int32(1) != 0 {
		v325 = v169
		goto L56
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	v174 = F_getvacant(m, l0, l1, l2, l2)
	mBase = m.M
	if v174 == int32(0) {
		goto L61
	} else {
		goto L62
	}
L60:
	;
	goto L59
L61:
	;
	v379 = int32(0)
	goto L54
L62:
	;
	goto L63
L63:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if int32(0) < v178 {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v182 = int32(0)
	goto L67
L65:
	;
	goto L66
L66:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v174)))
	v215 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v215)+12))
	v221 = v214 + int32(base.Ui32(v216)>>(uint(int32(3))%32))&int32(536870908)
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v221)))
	v223 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v221))) = v222 | v223<<(uint(v216)%32)
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v174)))
	v228 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v228 == v223 {
		goto L71
	} else {
		goto L72
	}
L67:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v174)))
	*(*int32)(unsafe.Add(mBase, uint32(v193+v182<<(uint(int32(2))%32)))) = int32(0)
	v200 = v182 + int32(1)
	v201 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v200 < v201 {
		v182 = v200
		goto L67
	} else {
		goto L69
	}
L68:
	;
	goto L66
L69:
	;
	goto L68
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v174)+8)) = int32(13)
	*(*int32)(unsafe.Add(mBase, uint32(v174)+4)) = v307
	v318 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v318 <= int32(0) {
		v358 = v174
		goto L55
	} else {
		goto L87
	}
L71:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v227)))
	v307 = v231
	goto L70
L72:
	;
	goto L73
L73:
	;
	if v228 <= int32(0) {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v307 = int32(0)
	goto L70
L75:
	;
	goto L76
L76:
	;
	v236 = v228 & int32(3)
	v237 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v228) {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v243 = v237
	v246 = v237
	v252 = v157
	goto L80
L78:
	;
	v272 = v237
	v275 = v237
	goto L79
L79:
	;
	v283 = v272
	v286 = v275
	v293 = v157
	goto L84
L80:
	;
	v256 = v227 + v243<<(uint(int32(2))%32)
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v256)+12))
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v256)+8))
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v256)+4))
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v256)))
	v264 = v257 ^ (v258 ^ (v259 ^ (v260 ^ v246)))
	v265 = int32(4)
	v266 = v243 + v265
	v268 = v252 + v265
	if v268 != v228&int32(2147483644) {
		v243 = v266
		v246 = v264
		v252 = v268
		goto L80
	} else {
		goto L82
	}
L81:
	;
	if v236 == int32(0) {
		v307 = v264
		goto L70
	} else {
		goto L83
	}
L82:
	;
	goto L81
L83:
	;
	v272 = v266
	v275 = v264
	goto L79
L84:
	;
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v227+v283<<(uint(int32(2))%32))))
	v298 = v297 ^ v286
	v299 = int32(1)
	v302 = v293 + v299
	if v302 != v236 {
		v283 = v283 + v299
		v286 = v298
		v293 = v302
		goto L84
	} else {
		goto L86
	}
L85:
	;
	v307 = v298
	goto L70
L86:
	;
	goto L85
L87:
	;
	v325 = v174
	goto L56
L88:
	;
	v344 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v344+v333<<(uint(int32(5))%32))+20)) = int32(0)
	v351 = v333 + int32(1)
	v352 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v351 < v352 {
		v333 = v351
		goto L88
	} else {
		goto L90
	}
L89:
	;
	v358 = v325
	goto L55
L90:
	;
	goto L89
L91:
	;
	v382 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v382 == l2 {
		goto L93
	} else {
		goto L94
	}
L92:
	;
	v409 = F_miss(m, l0, l1, v379, base.I32_extend16_s(v407), l2, l2)
	mBase = m.M
	v410 = m.ExcPending
	if v410 != 0 {
		goto L26
	} else {
		goto L100
	}
L93:
	;
	v384 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	v385 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v388 = int32(1)
	v393 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v384+(v385^int32(-1))&v388<<(uint(v388)%32))+20)))
	v407 = v393
	goto L92
L94:
	;
	goto L95
L95:
	;
	v396 = *(*int32)(unsafe.Add(mBase, uint32(l2-int32(4))))
	if base.Ui32(v396) <= base.Ui32(int32(2047)) {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v399 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
	v403 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v399+v396<<(uint(int32(1))%32)))))
	v407 = v403
	goto L92
L97:
	;
	goto L98
L98:
	;
	v404 = F_pg_reg_getcolor(m, v14, v396)
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
		goto L26
	} else {
		goto L99
	}
L99:
	;
	v407 = v404
	goto L92
L100:
	;
	if v409 == int32(0) {
		v591 = v157
		goto L4
	} else {
		goto L101
	}
L101:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v409)+20)) = l2
	if l3 != v13 {
		goto L103
	} else {
		goto L104
	}
L102:
	;
	v470 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v470 != 0 {
		v591 = v157
		goto L4
	} else {
		goto L121
	}
L103:
	;
	v417 = int32(4)
	goto L105
L104:
	;
	v417 = int32(0)
	goto L105
L105:
	;
	v418 = l3 + v417
	if base.Ui32(v418) <= base.Ui32(l2) {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	v463 = l2
	v465 = v409
	goto L102
L107:
	;
	goto L108
L108:
	;
	v425 = l2
	v426 = v409
	goto L109
L109:
	;
	v432 = *(*int32)(unsafe.Add(mBase, uint32(v425)))
	if base.Ui32(v432) <= base.Ui32(int32(2047)) {
		goto L112
	} else {
		goto L113
	}
L110:
	;
	v463 = v455
	v465 = v453
	goto L102
L111:
	;
	v443 = *(*int32)(unsafe.Add(mBase, uint32(v426)+24))
	v444 = base.I32_extend16_s(v442)
	v448 = *(*int32)(unsafe.Add(mBase, uint32(v443+v444<<(uint(int32(2))%32))))
	if v448 != 0 {
		v453 = v448
		goto L116
	} else {
		goto L117
	}
L112:
	;
	v435 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
	v439 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v435+v432<<(uint(int32(1))%32)))))
	v442 = v439
	goto L111
L113:
	;
	goto L114
L114:
	;
	v440 = F_pg_reg_getcolor(m, v14, v432)
	mBase = m.M
	v441 = m.ExcPending
	if v441 != 0 {
		goto L26
	} else {
		goto L115
	}
L115:
	;
	v442 = v440
	goto L111
L116:
	;
	v455 = v425 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v453)+20)) = v455
	if base.Ui32(v455) < base.Ui32(v418) {
		v425 = v455
		v426 = v453
		goto L109
	} else {
		goto L120
	}
L117:
	;
	v451 = F_miss(m, l0, l1, v426, v444, v425+int32(4), l2)
	mBase = m.M
	v452 = m.ExcPending
	if v452 != 0 {
		goto L26
	} else {
		goto L118
	}
L118:
	;
	if v451 != 0 {
		v453 = v451
		goto L116
	} else {
		goto L119
	}
L119:
	;
	v463 = v425
	v465 = v426
	goto L102
L120:
	;
	goto L110
L121:
	;
	v471 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if base.B2i32(v463 != v471)|base.B2i32(l3 != v471) != 0 {
		goto L122
	} else {
		goto L123
	}
L122:
	;
	v496 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v497 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v497 <= int32(0) {
		v572 = v496
		goto L133
	} else {
		goto L134
	}
L123:
	;
	if l4 != 0 {
		goto L124
	} else {
		goto L125
	}
L124:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = int32(1)
	goto L126
L125:
	;
	goto L126
L126:
	;
	v477 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	v478 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v484 = int32(*(*int16)(unsafe.Add(mBase, uint32(v477+(v478^int32(-1))&int32(2))+24)))
	v485 = F_miss(m, l0, l1, v465, v484, v463, l2)
	mBase = m.M
	v486 = m.ExcPending
	if v486 != 0 {
		goto L26
	} else {
		goto L127
	}
L127:
	;
	v487 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v487 != 0 {
		v591 = v157
		goto L4
	} else {
		goto L128
	}
L128:
	;
	if v485 == int32(0) {
		goto L122
	} else {
		goto L129
	}
L129:
	;
	v490 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v485)+8)))
	if v490&int32(2) != 0 {
		goto L130
	} else {
		goto L131
	}
L130:
	;
	return v463
L131:
	;
	goto L132
L132:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v485)+20)) = v463
	goto L122
L133:
	;
	if v572 != 0 {
		goto L169
	} else {
		goto L170
	}
L134:
	;
	v500 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v497&int32(1) != 0 {
		goto L135
	} else {
		goto L136
	}
L135:
	;
	v503 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v500)+8)))
	if v503&int32(2) == int32(0) {
		v514 = v496
		goto L138
	} else {
		goto L139
	}
L136:
	;
	v520 = v500
	v521 = v496
	v522 = v497
	goto L137
L137:
	;
	if v497 == int32(1) {
		v572 = v521
		goto L133
	} else {
		goto L147
	}
L138:
	;
	v520 = v500 + int32(32)
	v521 = v514
	v522 = v497 - int32(1)
	goto L137
L139:
	;
	v508 = *(*int32)(unsafe.Add(mBase, uint32(v500)+20))
	if v496 == v508 {
		v514 = v496
		goto L138
	} else {
		goto L140
	}
L140:
	;
	if base.Ui32(v496) < base.Ui32(v508) {
		goto L141
	} else {
		goto L142
	}
L141:
	;
	v511 = v508
	goto L143
L142:
	;
	v511 = v496
	goto L143
L143:
	;
	if v496 != 0 {
		goto L144
	} else {
		goto L145
	}
L144:
	;
	v512 = v511
	goto L146
L145:
	;
	v512 = v508
	goto L146
L146:
	;
	v514 = v512
	goto L138
L147:
	;
	v526 = v522
	v527 = v520
	v530 = v521
	goto L148
L148:
	;
	v537 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v527)+8)))
	if v537&int32(2) == int32(0) {
		v548 = v530
		goto L150
	} else {
		goto L151
	}
L149:
	;
	v572 = v560
	goto L133
L150:
	;
	v549 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v527)+40)))
	if v549&int32(2) == int32(0) {
		v560 = v548
		goto L159
	} else {
		goto L160
	}
L151:
	;
	v542 = *(*int32)(unsafe.Add(mBase, uint32(v527)+20))
	if v530 == v542 {
		v548 = v530
		goto L150
	} else {
		goto L152
	}
L152:
	;
	if base.Ui32(v530) < base.Ui32(v542) {
		goto L153
	} else {
		goto L154
	}
L153:
	;
	v545 = v542
	goto L155
L154:
	;
	v545 = v530
	goto L155
L155:
	;
	if v530 != 0 {
		goto L156
	} else {
		goto L157
	}
L156:
	;
	v546 = v545
	goto L158
L157:
	;
	v546 = v542
	goto L158
L158:
	;
	v548 = v546
	goto L150
L159:
	;
	v561 = int32(2)
	if v561 < v526 {
		v526 = v526 - v561
		v527 = v527 - int32(-64)
		v530 = v560
		goto L148
	} else {
		goto L168
	}
L160:
	;
	v554 = *(*int32)(unsafe.Add(mBase, uint32(v527)+52))
	if v548 == v554 {
		v560 = v548
		goto L159
	} else {
		goto L161
	}
L161:
	;
	if base.Ui32(v548) < base.Ui32(v554) {
		goto L162
	} else {
		goto L163
	}
L162:
	;
	v557 = v554
	goto L164
L163:
	;
	v557 = v548
	goto L164
L164:
	;
	if v548 != 0 {
		goto L165
	} else {
		goto L166
	}
L165:
	;
	v558 = v557
	goto L167
L166:
	;
	v558 = v554
	goto L167
L167:
	;
	v560 = v558
	goto L159
L168:
	;
	goto L149
L169:
	;
	v582 = v572 - int32(4)
	goto L171
L170:
	;
	v582 = int32(0)
	goto L171
L171:
	;
	v591 = v582
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
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
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
	var v122 int32
	_ = v122
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
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
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
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v194 int32
	_ = v194
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v218 int32
	_ = v218
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
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
	*(*int64)(unsafe.Add(mBase, uint32(v20)+16)) = int64(85899345928)
	v34 = F_hash_create(m, int32(_a_F_lookup_proof_cache_0), int64(256), v18+int32(-56), int32(40))
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
	*(*int32)(unsafe.Add(mBase, uint32(v20)+12)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v20)+8)) = l0
	v54 = F_hash_search(m, v46, v18+int32(-56), int32(1), v18+int32(-1))
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
	F_CacheRegisterSyscacheCallback(m, int32(3), int32(943), int64(0))
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
		goto L20
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
	F_list_free_deep(m, v211)
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L4
	} else {
		goto L52
	}
L19:
	;
	v209 = v4
	v211 = int32(0)
	v218 = v4
	goto L18
L20:
	;
	if v65 == int32(0) {
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v69 = F_get_op_index_interpretation(m, l0)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L4
	} else {
		goto L22
	}
L22:
	;
	if v69 == int32(0) {
		goto L19
	} else {
		goto L23
	}
L23:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v69)+4))
	if v73 <= int32(0) {
		v209 = v4
		v211 = v69
		v218 = v4
		goto L18
	} else {
		goto L24
	}
L24:
	;
	if l2 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v78 = int32(_a_F_lookup_proof_cache_1)
	goto L27
L26:
	;
	v78 = int32(_a_F_lookup_proof_cache_2)
	goto L27
L27:
	;
	if l2 != 0 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v81 = int32(_a_F_lookup_proof_cache_3)
	goto L30
L29:
	;
	v81 = int32(_a_F_lookup_proof_cache_4)
	goto L30
L30:
	;
	v94 = v4
	v95 = v4
	goto L31
L31:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v65)+4))
	if int32(0) < v99 {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	v209 = int32(0)
	v211 = v69
	v218 = v194
	goto L18
L33:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v69)+12))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v102+v94<<(uint(int32(2))%32))))
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v106)))
	v109 = int32(0)
	v122 = v95
	goto L36
L34:
	;
	v194 = v95
	goto L35
L35:
	;
	v199 = v94 + int32(1)
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v69)+4))
	if v199 < v200 {
		v94 = v199
		v95 = v194
		goto L31
	} else {
		goto L51
	}
L36:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v65)+12))
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v126+v109<<(uint(int32(2))%32))))
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v130)))
	if v107 != v131 {
		v175 = v122
		goto L38
	} else {
		goto L39
	}
L37:
	;
	v194 = v175
	goto L35
L38:
	;
	v178 = v109 + int32(1)
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v65)+4))
	if v178 < v179 {
		v109 = v178
		v122 = v175
		goto L36
	} else {
		goto L50
	}
L39:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v106)+4))
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v130)+4))
	v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v133+(v81+v134*int32(6))-int32(7)))))
	v142 = v122 | v141
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v78+v134*int32(24)+v133<<(uint(int32(2))%32)-int32(28))))
	switch v151 {
	case 0:
		v175 = v142
		goto L38
	default:
		goto L41
	case 6:
		goto L42
	}
L40:
	;
	if v166 == int32(0) {
		v175 = v142
		goto L38
	} else {
		goto L47
	}
L41:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v106)+12))
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v130)+12))
	v163 = F_get_opfamily_member_for_cmptype(m, v107, v161, v162, v151)
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L4
	} else {
		goto L46
	}
L42:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v106)+12))
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v130)+12))
	v155 = F_get_opfamily_member_for_cmptype(m, v107, v152, v153, int32(3))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L4
	} else {
		goto L43
	}
L43:
	;
	if v155 == int32(0) {
		v175 = v142
		goto L38
	} else {
		goto L44
	}
L44:
	;
	v159 = F_get_negator(m, v155)
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L4
	} else {
		goto L45
	}
L45:
	;
	v166 = v159
	goto L40
L46:
	;
	v166 = v163
	goto L40
L47:
	;
	v169 = F_op_volatile(m, v166)
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L4
	} else {
		goto L48
	}
L48:
	;
	if v169 == int32(105) {
		v209 = v166
		v211 = v69
		v218 = v142
		goto L18
	} else {
		goto L49
	}
L49:
	;
	v175 = v142
	goto L38
L50:
	;
	goto L37
L51:
	;
	goto L32
L52:
	;
	F_list_free_deep(m, v65)
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L4
	} else {
		goto L53
	}
L53:
	;
	if v218&int32(1) != 0 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v228 = F_op_volatile(m, l1)
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L4
	} else {
		goto L57
	}
L55:
	;
	v232 = v4
	goto L56
L56:
	;
	if l2 != 0 {
		goto L58
	} else {
		goto L59
	}
L57:
	;
	v232 = base.B2i32(v228 == int32(105))
	goto L56
L58:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v54)+11)) = uint8(v232)
	*(*int32)(unsafe.Add(mBase, uint32(v54)+16)) = v209
	v235 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v54)+9)) = uint8(v235)
	goto L8
L59:
	;
	goto L60
L60:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v54)+10)) = uint8(v232)
	*(*int32)(unsafe.Add(mBase, uint32(v54)+12)) = v209
	v239 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v54)+8)) = uint8(v239)
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
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
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
	var v94 int32
	_ = v94
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
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v117 int64
	_ = v117
	var v119 int64
	_ = v119
	var v121 int32
	_ = v121
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
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
				v139 = v20
				m.G0 = v10 + int32(16)
				return v139
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
								F_errfinish(m, int32(_a_F_lookup_rowtype_tupdesc_internal_1), int32(1869), int32(_a_F_lookup_rowtype_tupdesc_internal_2))
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
						F_errfinish(m, int32(_a_F_lookup_rowtype_tupdesc_internal_1), int32(1928), int32(_a_F_lookup_rowtype_tupdesc_internal_2))
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
					v139 = v49
					m.G0 = v10 + int32(16)
					return v139
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
									F_errfinish(m, int32(_a_F_lookup_rowtype_tupdesc_internal_1), int32(1928), int32(_a_F_lookup_rowtype_tupdesc_internal_2))
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
							if v60 != 0 {
								v63 = *(*int32)(unsafe.Add(mBase, _c_F_lookup_rowtype_tupdesc_internal[2]))
								v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)+4))
								v65 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
								v66 = F_dsa_get_address(m, v64, v65)
								mBase = m.M
								v67 = m.ExcPending
								if v67 != 0 {
									return int32(0)
								} else {
									v68 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
									v70 = *(*int32)(unsafe.Add(mBase, _c_F_lookup_rowtype_tupdesc_internal[1]))
									if v70 != 0 {
										v72 = *(*int32)(unsafe.Add(mBase, _c_F_lookup_rowtype_tupdesc_internal[0]))
										v84 = v70
										v85 = v72
										if v85 <= v68 {
											v88 = F_mul_size(m, int32(16), v85)
											mBase = m.M
											v89 = m.ExcPending
											if v89 != 0 {
												return int32(0)
											} else {
												v91 = int32(1)
												v94 = v68 + v91
												if v68&v94 != 0 {
													v99 = v91 << (uint(int32(32)-base.I32_clz(v94)) % 32)
												} else {
													v99 = v94
												}
												v100 = F_mul_size(m, int32(16), v99)
												mBase = m.M
												v101 = m.ExcPending
												if v101 != 0 {
													return int32(0)
												} else {
													v102 = F_repalloc0(m, v84, v88, v100)
													mBase = m.M
													v103 = m.ExcPending
													if v103 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, _c_F_lookup_rowtype_tupdesc_internal[0])) = v99
														*(*int32)(unsafe.Add(mBase, _c_F_lookup_rowtype_tupdesc_internal[1])) = v102
														v109 = v102
														v110 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
														v111 = int32(4)
														*(*int32)(unsafe.Add(mBase, uint32(v109+v110<<(uint(v111)%32))+8)) = v66
														v115 = int32(_a_F_lookup_rowtype_tupdesc_internal_4)
														v117 = *(*int64)(unsafe.Add(mBase, _c_F_lookup_rowtype_tupdesc_internal[3]))
														v119 = v117 + int64(1)
														*(*int64)(unsafe.Add(mBase, _c_F_lookup_rowtype_tupdesc_internal[3])) = v119
														v121 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
														*(*int64)(unsafe.Add(mBase, uint32(v109+v121<<(uint(v111)%32)))) = v119
														v127 = *(*int32)(unsafe.Add(mBase, _c_F_lookup_rowtype_tupdesc_internal[2]))
														v128 = *(*int32)(unsafe.Add(mBase, uint32(v127)+16))
														F_dshash_release_lock(m, v128, v60)
														mBase = m.M
														v130 = m.ExcPending
														if v130 != 0 {
															return int32(0)
														} else {
															v132 = *(*int32)(unsafe.Add(mBase, _c_F_lookup_rowtype_tupdesc_internal[1]))
															v133 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
															v137 = *(*int32)(unsafe.Add(mBase, uint32(v132+v133<<(uint(int32(4))%32))+8))
															v139 = v137
															m.G0 = v10 + int32(16)
															return v139
														}
													}
												}
											}
										} else {
											v109 = v84
											v110 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
											v111 = int32(4)
											*(*int32)(unsafe.Add(mBase, uint32(v109+v110<<(uint(v111)%32))+8)) = v66
											v115 = int32(_a_F_lookup_rowtype_tupdesc_internal_4)
											v117 = *(*int64)(unsafe.Add(mBase, _c_F_lookup_rowtype_tupdesc_internal[3]))
											v119 = v117 + int64(1)
											*(*int64)(unsafe.Add(mBase, _c_F_lookup_rowtype_tupdesc_internal[3])) = v119
											v121 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
											*(*int64)(unsafe.Add(mBase, uint32(v109+v121<<(uint(v111)%32)))) = v119
											v127 = *(*int32)(unsafe.Add(mBase, _c_F_lookup_rowtype_tupdesc_internal[2]))
											v128 = *(*int32)(unsafe.Add(mBase, uint32(v127)+16))
											F_dshash_release_lock(m, v128, v60)
											mBase = m.M
											v130 = m.ExcPending
											if v130 != 0 {
												return int32(0)
											} else {
												v132 = *(*int32)(unsafe.Add(mBase, _c_F_lookup_rowtype_tupdesc_internal[1]))
												v133 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
												v137 = *(*int32)(unsafe.Add(mBase, uint32(v132+v133<<(uint(int32(4))%32))+8))
												v139 = v137
												m.G0 = v10 + int32(16)
												return v139
											}
										}
									} else {
										v75 = *(*int32)(unsafe.Add(mBase, _c_F_lookup_rowtype_tupdesc_internal[4]))
										v77 = F_MemoryContextAllocZero(m, v75, int32(1024))
										mBase = m.M
										v78 = m.ExcPending
										if v78 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, _c_F_lookup_rowtype_tupdesc_internal[0])) = int32(64)
											*(*int32)(unsafe.Add(mBase, _c_F_lookup_rowtype_tupdesc_internal[1])) = v77
											v84 = v77
											v85 = int32(64)
											if v85 <= v68 {
												v88 = F_mul_size(m, int32(16), v85)
												mBase = m.M
												v89 = m.ExcPending
												if v89 != 0 {
													return int32(0)
												} else {
													v91 = int32(1)
													v94 = v68 + v91
													if v68&v94 != 0 {
														v99 = v91 << (uint(int32(32)-base.I32_clz(v94)) % 32)
													} else {
														v99 = v94
													}
													v100 = F_mul_size(m, int32(16), v99)
													mBase = m.M
													v101 = m.ExcPending
													if v101 != 0 {
														return int32(0)
													} else {
														v102 = F_repalloc0(m, v84, v88, v100)
														mBase = m.M
														v103 = m.ExcPending
														if v103 != 0 {
															return int32(0)
														} else {
															*(*int32)(unsafe.Add(mBase, _c_F_lookup_rowtype_tupdesc_internal[0])) = v99
															*(*int32)(unsafe.Add(mBase, _c_F_lookup_rowtype_tupdesc_internal[1])) = v102
															v109 = v102
															v110 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
															v111 = int32(4)
															*(*int32)(unsafe.Add(mBase, uint32(v109+v110<<(uint(v111)%32))+8)) = v66
															v115 = int32(_a_F_lookup_rowtype_tupdesc_internal_4)
															v117 = *(*int64)(unsafe.Add(mBase, _c_F_lookup_rowtype_tupdesc_internal[3]))
															v119 = v117 + int64(1)
															*(*int64)(unsafe.Add(mBase, _c_F_lookup_rowtype_tupdesc_internal[3])) = v119
															v121 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
															*(*int64)(unsafe.Add(mBase, uint32(v109+v121<<(uint(v111)%32)))) = v119
															v127 = *(*int32)(unsafe.Add(mBase, _c_F_lookup_rowtype_tupdesc_internal[2]))
															v128 = *(*int32)(unsafe.Add(mBase, uint32(v127)+16))
															F_dshash_release_lock(m, v128, v60)
															mBase = m.M
															v130 = m.ExcPending
															if v130 != 0 {
																return int32(0)
															} else {
																v132 = *(*int32)(unsafe.Add(mBase, _c_F_lookup_rowtype_tupdesc_internal[1]))
																v133 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
																v137 = *(*int32)(unsafe.Add(mBase, uint32(v132+v133<<(uint(int32(4))%32))+8))
																v139 = v137
																m.G0 = v10 + int32(16)
																return v139
															}
														}
													}
												}
											} else {
												v109 = v84
												v110 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
												v111 = int32(4)
												*(*int32)(unsafe.Add(mBase, uint32(v109+v110<<(uint(v111)%32))+8)) = v66
												v115 = int32(_a_F_lookup_rowtype_tupdesc_internal_4)
												v117 = *(*int64)(unsafe.Add(mBase, _c_F_lookup_rowtype_tupdesc_internal[3]))
												v119 = v117 + int64(1)
												*(*int64)(unsafe.Add(mBase, _c_F_lookup_rowtype_tupdesc_internal[3])) = v119
												v121 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
												*(*int64)(unsafe.Add(mBase, uint32(v109+v121<<(uint(v111)%32)))) = v119
												v127 = *(*int32)(unsafe.Add(mBase, _c_F_lookup_rowtype_tupdesc_internal[2]))
												v128 = *(*int32)(unsafe.Add(mBase, uint32(v127)+16))
												F_dshash_release_lock(m, v128, v60)
												mBase = m.M
												v130 = m.ExcPending
												if v130 != 0 {
													return int32(0)
												} else {
													v132 = *(*int32)(unsafe.Add(mBase, _c_F_lookup_rowtype_tupdesc_internal[1]))
													v133 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
													v137 = *(*int32)(unsafe.Add(mBase, uint32(v132+v133<<(uint(int32(4))%32))+8))
													v139 = v137
													m.G0 = v10 + int32(16)
													return v139
												}
											}
										}
									}
								}
							} else {
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
											F_errfinish(m, int32(_a_F_lookup_rowtype_tupdesc_internal_1), int32(1928), int32(_a_F_lookup_rowtype_tupdesc_internal_2))
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
								F_errfinish(m, int32(_a_F_lookup_rowtype_tupdesc_internal_1), int32(1928), int32(_a_F_lookup_rowtype_tupdesc_internal_2))
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
						if v60 != 0 {
							v63 = *(*int32)(unsafe.Add(mBase, _c_F_lookup_rowtype_tupdesc_internal[2]))
							v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)+4))
							v65 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
							v66 = F_dsa_get_address(m, v64, v65)
							mBase = m.M
							v67 = m.ExcPending
							if v67 != 0 {
								return int32(0)
							} else {
								v68 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
								v70 = *(*int32)(unsafe.Add(mBase, _c_F_lookup_rowtype_tupdesc_internal[1]))
								if v70 != 0 {
									v72 = *(*int32)(unsafe.Add(mBase, _c_F_lookup_rowtype_tupdesc_internal[0]))
									v84 = v70
									v85 = v72
									if v85 <= v68 {
										v88 = F_mul_size(m, int32(16), v85)
										mBase = m.M
										v89 = m.ExcPending
										if v89 != 0 {
											return int32(0)
										} else {
											v91 = int32(1)
											v94 = v68 + v91
											if v68&v94 != 0 {
												v99 = v91 << (uint(int32(32)-base.I32_clz(v94)) % 32)
											} else {
												v99 = v94
											}
											v100 = F_mul_size(m, int32(16), v99)
											mBase = m.M
											v101 = m.ExcPending
											if v101 != 0 {
												return int32(0)
											} else {
												v102 = F_repalloc0(m, v84, v88, v100)
												mBase = m.M
												v103 = m.ExcPending
												if v103 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, _c_F_lookup_rowtype_tupdesc_internal[0])) = v99
													*(*int32)(unsafe.Add(mBase, _c_F_lookup_rowtype_tupdesc_internal[1])) = v102
													v109 = v102
													v110 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
													v111 = int32(4)
													*(*int32)(unsafe.Add(mBase, uint32(v109+v110<<(uint(v111)%32))+8)) = v66
													v115 = int32(_a_F_lookup_rowtype_tupdesc_internal_4)
													v117 = *(*int64)(unsafe.Add(mBase, _c_F_lookup_rowtype_tupdesc_internal[3]))
													v119 = v117 + int64(1)
													*(*int64)(unsafe.Add(mBase, _c_F_lookup_rowtype_tupdesc_internal[3])) = v119
													v121 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
													*(*int64)(unsafe.Add(mBase, uint32(v109+v121<<(uint(v111)%32)))) = v119
													v127 = *(*int32)(unsafe.Add(mBase, _c_F_lookup_rowtype_tupdesc_internal[2]))
													v128 = *(*int32)(unsafe.Add(mBase, uint32(v127)+16))
													F_dshash_release_lock(m, v128, v60)
													mBase = m.M
													v130 = m.ExcPending
													if v130 != 0 {
														return int32(0)
													} else {
														v132 = *(*int32)(unsafe.Add(mBase, _c_F_lookup_rowtype_tupdesc_internal[1]))
														v133 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
														v137 = *(*int32)(unsafe.Add(mBase, uint32(v132+v133<<(uint(int32(4))%32))+8))
														v139 = v137
														m.G0 = v10 + int32(16)
														return v139
													}
												}
											}
										}
									} else {
										v109 = v84
										v110 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
										v111 = int32(4)
										*(*int32)(unsafe.Add(mBase, uint32(v109+v110<<(uint(v111)%32))+8)) = v66
										v115 = int32(_a_F_lookup_rowtype_tupdesc_internal_4)
										v117 = *(*int64)(unsafe.Add(mBase, _c_F_lookup_rowtype_tupdesc_internal[3]))
										v119 = v117 + int64(1)
										*(*int64)(unsafe.Add(mBase, _c_F_lookup_rowtype_tupdesc_internal[3])) = v119
										v121 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
										*(*int64)(unsafe.Add(mBase, uint32(v109+v121<<(uint(v111)%32)))) = v119
										v127 = *(*int32)(unsafe.Add(mBase, _c_F_lookup_rowtype_tupdesc_internal[2]))
										v128 = *(*int32)(unsafe.Add(mBase, uint32(v127)+16))
										F_dshash_release_lock(m, v128, v60)
										mBase = m.M
										v130 = m.ExcPending
										if v130 != 0 {
											return int32(0)
										} else {
											v132 = *(*int32)(unsafe.Add(mBase, _c_F_lookup_rowtype_tupdesc_internal[1]))
											v133 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
											v137 = *(*int32)(unsafe.Add(mBase, uint32(v132+v133<<(uint(int32(4))%32))+8))
											v139 = v137
											m.G0 = v10 + int32(16)
											return v139
										}
									}
								} else {
									v75 = *(*int32)(unsafe.Add(mBase, _c_F_lookup_rowtype_tupdesc_internal[4]))
									v77 = F_MemoryContextAllocZero(m, v75, int32(1024))
									mBase = m.M
									v78 = m.ExcPending
									if v78 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, _c_F_lookup_rowtype_tupdesc_internal[0])) = int32(64)
										*(*int32)(unsafe.Add(mBase, _c_F_lookup_rowtype_tupdesc_internal[1])) = v77
										v84 = v77
										v85 = int32(64)
										if v85 <= v68 {
											v88 = F_mul_size(m, int32(16), v85)
											mBase = m.M
											v89 = m.ExcPending
											if v89 != 0 {
												return int32(0)
											} else {
												v91 = int32(1)
												v94 = v68 + v91
												if v68&v94 != 0 {
													v99 = v91 << (uint(int32(32)-base.I32_clz(v94)) % 32)
												} else {
													v99 = v94
												}
												v100 = F_mul_size(m, int32(16), v99)
												mBase = m.M
												v101 = m.ExcPending
												if v101 != 0 {
													return int32(0)
												} else {
													v102 = F_repalloc0(m, v84, v88, v100)
													mBase = m.M
													v103 = m.ExcPending
													if v103 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, _c_F_lookup_rowtype_tupdesc_internal[0])) = v99
														*(*int32)(unsafe.Add(mBase, _c_F_lookup_rowtype_tupdesc_internal[1])) = v102
														v109 = v102
														v110 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
														v111 = int32(4)
														*(*int32)(unsafe.Add(mBase, uint32(v109+v110<<(uint(v111)%32))+8)) = v66
														v115 = int32(_a_F_lookup_rowtype_tupdesc_internal_4)
														v117 = *(*int64)(unsafe.Add(mBase, _c_F_lookup_rowtype_tupdesc_internal[3]))
														v119 = v117 + int64(1)
														*(*int64)(unsafe.Add(mBase, _c_F_lookup_rowtype_tupdesc_internal[3])) = v119
														v121 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
														*(*int64)(unsafe.Add(mBase, uint32(v109+v121<<(uint(v111)%32)))) = v119
														v127 = *(*int32)(unsafe.Add(mBase, _c_F_lookup_rowtype_tupdesc_internal[2]))
														v128 = *(*int32)(unsafe.Add(mBase, uint32(v127)+16))
														F_dshash_release_lock(m, v128, v60)
														mBase = m.M
														v130 = m.ExcPending
														if v130 != 0 {
															return int32(0)
														} else {
															v132 = *(*int32)(unsafe.Add(mBase, _c_F_lookup_rowtype_tupdesc_internal[1]))
															v133 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
															v137 = *(*int32)(unsafe.Add(mBase, uint32(v132+v133<<(uint(int32(4))%32))+8))
															v139 = v137
															m.G0 = v10 + int32(16)
															return v139
														}
													}
												}
											}
										} else {
											v109 = v84
											v110 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
											v111 = int32(4)
											*(*int32)(unsafe.Add(mBase, uint32(v109+v110<<(uint(v111)%32))+8)) = v66
											v115 = int32(_a_F_lookup_rowtype_tupdesc_internal_4)
											v117 = *(*int64)(unsafe.Add(mBase, _c_F_lookup_rowtype_tupdesc_internal[3]))
											v119 = v117 + int64(1)
											*(*int64)(unsafe.Add(mBase, _c_F_lookup_rowtype_tupdesc_internal[3])) = v119
											v121 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
											*(*int64)(unsafe.Add(mBase, uint32(v109+v121<<(uint(v111)%32)))) = v119
											v127 = *(*int32)(unsafe.Add(mBase, _c_F_lookup_rowtype_tupdesc_internal[2]))
											v128 = *(*int32)(unsafe.Add(mBase, uint32(v127)+16))
											F_dshash_release_lock(m, v128, v60)
											mBase = m.M
											v130 = m.ExcPending
											if v130 != 0 {
												return int32(0)
											} else {
												v132 = *(*int32)(unsafe.Add(mBase, _c_F_lookup_rowtype_tupdesc_internal[1]))
												v133 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
												v137 = *(*int32)(unsafe.Add(mBase, uint32(v132+v133<<(uint(int32(4))%32))+8))
												v139 = v137
												m.G0 = v10 + int32(16)
												return v139
											}
										}
									}
								}
							}
						} else {
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
										F_errfinish(m, int32(_a_F_lookup_rowtype_tupdesc_internal_1), int32(1928), int32(_a_F_lookup_rowtype_tupdesc_internal_2))
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
func F_lt_q_regex(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
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
	var v17 int32
	_ = v17
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
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v49 int64
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v61 int32
	_ = v61
	var v80 int64
	_ = v80
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v9 = F_pg_detoast_datum(m, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v14 = F_pg_detoast_datum(m, v13)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	v20 = F_ArrayGetNItemsSafe(m, v17, v14+int32(16))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	if v22 < int32(2) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L1
	} else {
		goto L35
	}
L6:
	;
	v25 = F_array_contains_nulls(m, v14)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
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
	v93 = m.ExcPending
	if v93 != 0 {
		goto L1
	} else {
		goto L31
	}
L9:
	;
	if v25 != 0 {
		goto L5
	} else {
		goto L10
	}
L10:
	;
	if int32(0) < v20 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v81 != v9 {
		goto L23
	} else {
		goto L24
	}
L12:
	;
	if v16 != 0 {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	goto L14
L14:
	;
	v80 = int64(0)
	goto L11
L15:
	;
	v36 = v16
	goto L17
L16:
	;
	v36 = (v17<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L17
L17:
	;
	v40 = v14 + v36
	v41 = v20
	goto L18
L18:
	;
	v49 = F_DirectFunctionCall2Coll(m, int32(_a_F_lt_q_regex_0), int32(0), base.I64_extend_i32_u(v9), base.I64_extend_i32_u(v40))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L1
	} else {
		goto L20
	}
L19:
	;
	goto L14
L20:
	;
	if v49 != int64(0) {
		v80 = int64(1)
		goto L11
	} else {
		goto L21
	}
L21:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
	v61 = int32(1)
	if v61 < v41 {
		v40 = v40 + (int32(base.Ui32(v53)>>(uint(int32(2))%32))+int32(3))&int32(2147483644)
		v41 = v41 - v61
		goto L18
	} else {
		goto L22
	}
L22:
	;
	goto L19
L23:
	;
	F_pfree(m, v9)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L1
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v85 != v14 {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	goto L25
L27:
	;
	F_pfree(m, v14)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L1
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	return v80
L30:
	;
	goto L29
L31:
	;
	F_errcode(m, int32(352845954))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	F_errmsg(m, int32(_a_F_lt_q_regex_1), int32(0))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	F_errfinish(m, int32(_a_F_lt_q_regex_2), int32(298), int32(_a_F_lt_q_regex_3))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
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
	v112 = m.ExcPending
	if v112 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	F_errmsg(m, int32(_a_F_lt_q_regex_4), int32(0))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	F_errfinish(m, int32(_a_F_lt_q_regex_2), int32(302), int32(_a_F_lt_q_regex_3))
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
func F_lt_q_rregex(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int64
	_ = v4
	var v5 int64
	_ = v5
	var v6 int64
	_ = v6
	var v9 int32
	_ = v9
	v4 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v5 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v6 = F_DirectFunctionCall2Coll(m, int32(_a_F_lt_q_rregex_0), int32(0), v4, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int64(0)
	} else {
		return v6
	}
}
func F_ltreeparentsel(m *base.Module, l0 int32) int64 {
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
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	v8 = F_generic_restriction_selectivity(m, v2, v3, int32(0), v5, v6, float64(0.001))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int64(0)
	} else {
		return base.I64_reinterpret_f64(v8)
	}
}
func F_lz4_decompress_datum(m *base.Module) int32 {
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
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	v2 = m.G0
	v4 = v2 - int32(32)
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
			*(*int32)(unsafe.Add(mBase, uint32(v4)+16)) = int32(_a_F_lz4_decompress_datum_0)
			F_errmsg(m, int32(_a_F_lz4_decompress_datum_1), v4+int32(16))
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v4))) = int32(_a_F_lz4_decompress_datum_0)
				v25 = F_errdetail(m, int32(_a_F_lz4_decompress_datum_2), v4)
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_lz4_decompress_datum_3), int32(185), int32(_a_F_lz4_decompress_datum_4))
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
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
