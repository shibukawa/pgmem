package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_JsonbExtractScalar(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v12 int32
	_ = v12
	var v19 int32
	_ = v19
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
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v11 = int32(1342177280)
	v12 = v10 & v11
	if v12 != v11 {
		if v10&int32(1073741824) != 0 {
			v19 = int32(16)
		} else {
			v19 = int32(17)
		}
		*(*int32)(unsafe.Add(mBase, uint32(l1))) = v19
		m.G0 = v8 + int32(32)
		return base.B2i32(v12 == int32(1342177280))
	} else {
		v21 = F_JsonbIteratorInit(m, l0)
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v8)+28)) = v21
			v27 = v8 + int32(28)
			v29 = v8 + int32(8)
			v31 = F_JsonbIteratorNext(m, v27, v29, int32(1))
			mBase = m.M
			v32 = m.ExcPending
			if v32 != 0 {
				return int32(0)
			} else {
				v34 = F_JsonbIteratorNext(m, v27, l1, int32(1))
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return int32(0)
				} else {
					v37 = F_JsonbIteratorNext(m, v27, v29, int32(1))
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
						return int32(0)
					} else {
						v40 = F_JsonbIteratorNext(m, v27, v29, int32(1))
						mBase = m.M
						v41 = m.ExcPending
						if v41 != 0 {
							return int32(0)
						} else {
							m.G0 = v8 + int32(32)
							return base.B2i32(v12 == int32(1342177280))
						}
					}
				}
			}
		}
	}
}
func F_JsonbIteratorNext(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
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
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v87 int32
	_ = v87
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v146 int32
	_ = v146
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v174 int32
	_ = v174
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v231 int32
	_ = v231
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v247 int32
	_ = v247
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v285 int32
	_ = v285
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v298 int32
	_ = v298
	var v304 int32
	_ = v304
	var v307 int32
	_ = v307
	var v311 int32
	_ = v311
	var v328 int32
	_ = v328
	var v332 int32
	_ = v332
	var v341 int32
	_ = v341
	var v344 int32
	_ = v344
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v357 int32
	_ = v357
	var v363 int32
	_ = v363
	var v366 int32
	_ = v366
	var v370 int32
	_ = v370
	var v385 int32
	_ = v385
	var v397 int32
	_ = v397
	var v400 int32
	_ = v400
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
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
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v433 int32
	_ = v433
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v446 int32
	_ = v446
	var v452 int32
	_ = v452
	var v455 int32
	_ = v455
	var v459 int32
	_ = v459
	var v476 int32
	_ = v476
	var v480 int32
	_ = v480
	var v489 int32
	_ = v489
	var v492 int32
	_ = v492
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v505 int32
	_ = v505
	var v511 int32
	_ = v511
	var v514 int32
	_ = v514
	var v518 int32
	_ = v518
	var v533 int32
	_ = v533
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v551 int32
	_ = v551
	var v553 int32
	_ = v553
	var v556 int32
	_ = v556
	var v558 int32
	_ = v558
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v566 int32
	_ = v566
	var v570 int32
	_ = v570
	var v572 int32
	_ = v572
	var v575 int32
	_ = v575
	var v577 int32
	_ = v577
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v589 int32
	_ = v589
	var v593 int32
	_ = v593
	var v598 int32
	_ = v598
	var v602 int32
	_ = v602
	var v606 int32
	_ = v606
	var v611 int32
	_ = v611
	var v615 int32
	_ = v615
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v618 int32
	_ = v618
	var v620 int32
	_ = v620
	var v626 int32
	_ = v626
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v8 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v626
L2:
	;
	v12 = v8
	goto L5
L3:
	;
	goto L4
L4:
	;
	v620 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v620
	v626 = v620
	goto L1
L5:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v12)+32))
	switch v16 {
	case 0:
		goto L14
	case 1:
		goto L13
	case 2:
		goto L12
	case 3:
		goto L11
	case 4:
		goto L10
	default:
		goto L9
	}
L7:
	;
	v615 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v616 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v617 = F_iteratorFromContainer(m, v615, v616)
	mBase = m.M
	v618 = m.ExcPending
	if v618 != 0 {
		goto L18
	} else {
		goto L126
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v602 = m.ExcPending
	if v602 != 0 {
		goto L18
	} else {
		goto L123
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v589 = m.ExcPending
	if v589 != 0 {
		goto L18
	} else {
		goto L120
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = int32(3)
	v407 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v408 = *(*int32)(unsafe.Add(mBase, uint32(v407)))
	v409 = *(*int32)(unsafe.Add(mBase, uint32(v407)+4))
	v410 = *(*int32)(unsafe.Add(mBase, uint32(v407)+20))
	v411 = v409 + v410
	v412 = *(*int32)(unsafe.Add(mBase, uint32(v407)+16))
	v413 = *(*int32)(unsafe.Add(mBase, uint32(v407)+28))
	v421 = v408 + v411<<(uint(int32(2))%32) + int32(4)
	v422 = *(*int32)(unsafe.Add(mBase, uint32(v421)))
	switch int32(base.Ui32(v422)>>(uint(int32(28))%32)) & int32(7) {
	case 0:
		goto L92
	case 1:
		goto L91
	case 2:
		goto L89
	case 3:
		goto L90
	case 4:
		goto L93
	default:
		goto L88
	}
L11:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	if base.Ui32(v253) <= base.Ui32(v252) {
		goto L56
	} else {
		goto L57
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(17)
	v210 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v210)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v211
	v213 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v214 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v213)+20)) = v214
	v216 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v216)+24)) = v214
	v219 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v219)+4))
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v219)))
	v224 = v214
	v226 = v220
	goto L50
L13:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	if base.Ui32(v40) <= base.Ui32(v39) {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(16)
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v20
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+8)))
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+12)) = uint8(v23)
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v26 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v25)+20)) = v26
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+24)) = v26
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+28)) = v26
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v34)+32)) = int32(1)
	return int32(4)
L15:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v12)+36))
	F_pfree(m, v12)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	goto L17
L17:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v12)+24))
	v62 = v52 + v39<<(uint(int32(2))%32) + int32(4)
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)))
	switch int32(base.Ui32(v63)>>(uint(int32(28))%32)) & int32(7) {
	case 0:
		goto L25
	case 1:
		goto L24
	case 2:
		goto L22
	case 3:
		goto L23
	case 4:
		goto L26
	default:
		goto L21
	}
L18:
	;
	return int32(0)
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v42
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(0)
	return int32(5)
L20:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v186)+12))
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v186)+20))
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v187+v188<<(uint(int32(2))%32))))
	v194 = v192 & int32(268435455)
	if int32(0) <= v192 {
		goto L46
	} else {
		goto L47
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(18)
	v130 = (v54 + int32(3)) & int32(-4)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v53 + v130
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v62)))
	if v133 < int32(0) {
		goto L37
	} else {
		goto L38
	}
L22:
	;
	v121 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+4)) = uint8(v121)
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(3)
	goto L20
L23:
	;
	v117 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+4)) = uint8(v117)
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(3)
	goto L20
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v53 + (v54+int32(3))&int32(-4)
	goto L20
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v53 + v54
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v62)))
	if v74 < int32(0) {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(0)
	goto L20
L27:
	;
	v79 = v39
	v80 = int32(0)
	goto L30
L28:
	;
	goto L29
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v74 & int32(268435455)
	goto L20
L30:
	;
	v87 = v79 - int32(1)
	if int32(0) <= v87 {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v74&int32(268435455) - v100
	goto L20
L32:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v52+v79<<(uint(int32(2))%32))))
	v96 = v93&int32(268435455) + v80
	if int32(0) <= v93 {
		v79 = v87
		v80 = v96
		goto L30
	} else {
		goto L35
	}
L33:
	;
	v100 = v80
	goto L34
L34:
	;
	goto L31
L35:
	;
	v100 = v96
	goto L34
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v174 + (v54 - v130)
	goto L20
L37:
	;
	v138 = v39
	v139 = int32(0)
	goto L40
L38:
	;
	goto L39
L39:
	;
	v174 = v133 & int32(268435455)
	goto L36
L40:
	;
	v146 = v138 - int32(1)
	if int32(0) <= v146 {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	v174 = v133&int32(268435455) - v159
	goto L36
L42:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v52+v138<<(uint(int32(2))%32))))
	v155 = v152&int32(268435455) + v139
	if int32(0) <= v152 {
		v138 = v146
		v139 = v155
		goto L40
	} else {
		goto L45
	}
L43:
	;
	v159 = v139
	goto L44
L44:
	;
	goto L41
L45:
	;
	v159 = v155
	goto L44
L46:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v186)+24))
	v199 = v197 + v194
	goto L48
L47:
	;
	v199 = v194
	goto L48
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v186)+24)) = v199
	v201 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v201)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v201)+20)) = v202 + int32(1)
	v206 = int32(3)
	if l2 != 0 {
		v626 = v206
		goto L1
	} else {
		goto L49
	}
L49:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	switch v207 {
	case 0, 1, 2, 3, 32:
		v626 = v206
		goto L1
	default:
		goto L7
	}
L50:
	;
	v231 = v226 - int32(1)
	if int32(0) <= v231 {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v219)+28)) = v243
	v247 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v247)+32)) = int32(3)
	return int32(6)
L52:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v221+v226<<(uint(int32(2))%32))))
	v240 = v237&int32(268435455) + v224
	if int32(0) <= v237 {
		v224 = v240
		v226 = v231
		goto L50
	} else {
		goto L55
	}
L53:
	;
	v243 = v224
	goto L54
L54:
	;
	goto L51
L55:
	;
	v243 = v240
	goto L54
L56:
	;
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v12)+36))
	F_pfree(m, v12)
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L18
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v12)+24))
	v273 = v263 + v252<<(uint(int32(2))%32) + int32(4)
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v273)))
	switch int32(base.Ui32(v274)>>(uint(int32(28))%32)) & int32(7) {
	case 0:
		goto L65
	case 1:
		goto L64
	case 2:
		goto L62
	case 3:
		goto L63
	case 4:
		goto L66
	default:
		goto L61
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v255
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(0)
	return int32(7)
L60:
	;
	v397 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v397 != int32(1) {
		goto L8
	} else {
		goto L86
	}
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(18)
	v341 = (v265 + int32(3)) & int32(-4)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v264 + v341
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v273)))
	if v344 < int32(0) {
		goto L77
	} else {
		goto L78
	}
L62:
	;
	v332 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+4)) = uint8(v332)
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(3)
	goto L60
L63:
	;
	v328 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+4)) = uint8(v328)
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(3)
	goto L60
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v264 + (v265+int32(3))&int32(-4)
	goto L60
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v264 + v265
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v273)))
	if v285 < int32(0) {
		goto L67
	} else {
		goto L68
	}
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(0)
	goto L60
L67:
	;
	v290 = v252
	v291 = int32(0)
	goto L70
L68:
	;
	goto L69
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v285 & int32(268435455)
	goto L60
L70:
	;
	v298 = v290 - int32(1)
	if int32(0) <= v298 {
		goto L72
	} else {
		goto L73
	}
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v285&int32(268435455) - v311
	goto L60
L72:
	;
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v263+v290<<(uint(int32(2))%32))))
	v307 = v304&int32(268435455) + v291
	if int32(0) <= v304 {
		v290 = v298
		v291 = v307
		goto L70
	} else {
		goto L75
	}
L73:
	;
	v311 = v291
	goto L74
L74:
	;
	goto L71
L75:
	;
	v311 = v307
	goto L74
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v385 + (v265 - v341)
	goto L60
L77:
	;
	v349 = v252
	v350 = int32(0)
	goto L80
L78:
	;
	goto L79
L79:
	;
	v385 = v344 & int32(268435455)
	goto L76
L80:
	;
	v357 = v349 - int32(1)
	if int32(0) <= v357 {
		goto L82
	} else {
		goto L83
	}
L81:
	;
	v385 = v344&int32(268435455) - v370
	goto L76
L82:
	;
	v363 = *(*int32)(unsafe.Add(mBase, uint32(v263+v349<<(uint(int32(2))%32))))
	v366 = v363&int32(268435455) + v350
	if int32(0) <= v363 {
		v349 = v357
		v350 = v366
		goto L80
	} else {
		goto L85
	}
L83:
	;
	v370 = v350
	goto L84
L84:
	;
	goto L81
L85:
	;
	v370 = v366
	goto L84
L86:
	;
	v400 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v400)+32)) = int32(4)
	return int32(1)
L87:
	;
	v545 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v546 = *(*int32)(unsafe.Add(mBase, uint32(v545)+12))
	v547 = *(*int32)(unsafe.Add(mBase, uint32(v545)+20))
	v551 = *(*int32)(unsafe.Add(mBase, uint32(v546+v547<<(uint(int32(2))%32))))
	v553 = v551 & int32(268435455)
	if int32(0) <= v551 {
		goto L113
	} else {
		goto L114
	}
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(18)
	v489 = (v413 + int32(3)) & int32(-4)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v412 + v489
	v492 = *(*int32)(unsafe.Add(mBase, uint32(v421)))
	if v492 < int32(0) {
		goto L104
	} else {
		goto L105
	}
L89:
	;
	v480 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+4)) = uint8(v480)
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(3)
	goto L87
L90:
	;
	v476 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+4)) = uint8(v476)
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(3)
	goto L87
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v412 + (v413+int32(3))&int32(-4)
	goto L87
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v412 + v413
	v433 = *(*int32)(unsafe.Add(mBase, uint32(v421)))
	if v433 < int32(0) {
		goto L94
	} else {
		goto L95
	}
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(0)
	goto L87
L94:
	;
	v438 = v411
	v439 = int32(0)
	goto L97
L95:
	;
	goto L96
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v433 & int32(268435455)
	goto L87
L97:
	;
	v446 = v438 - int32(1)
	if int32(0) <= v446 {
		goto L99
	} else {
		goto L100
	}
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v433&int32(268435455) - v459
	goto L87
L99:
	;
	v452 = *(*int32)(unsafe.Add(mBase, uint32(v408+v438<<(uint(int32(2))%32))))
	v455 = v452&int32(268435455) + v439
	if int32(0) <= v452 {
		v438 = v446
		v439 = v455
		goto L97
	} else {
		goto L102
	}
L100:
	;
	v459 = v439
	goto L101
L101:
	;
	goto L98
L102:
	;
	v459 = v455
	goto L101
L103:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v533 + (v413 - v489)
	goto L87
L104:
	;
	v497 = v411
	v498 = int32(0)
	goto L107
L105:
	;
	goto L106
L106:
	;
	v533 = v492 & int32(268435455)
	goto L103
L107:
	;
	v505 = v497 - int32(1)
	if int32(0) <= v505 {
		goto L109
	} else {
		goto L110
	}
L108:
	;
	v533 = v492&int32(268435455) - v518
	goto L103
L109:
	;
	v511 = *(*int32)(unsafe.Add(mBase, uint32(v408+v497<<(uint(int32(2))%32))))
	v514 = v511&int32(268435455) + v498
	if int32(0) <= v511 {
		v497 = v505
		v498 = v514
		goto L107
	} else {
		goto L112
	}
L110:
	;
	v518 = v498
	goto L111
L111:
	;
	goto L108
L112:
	;
	v518 = v514
	goto L111
L113:
	;
	v556 = *(*int32)(unsafe.Add(mBase, uint32(v545)+24))
	v558 = v556 + v553
	goto L115
L114:
	;
	v558 = v553
	goto L115
L115:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v545)+24)) = v558
	v560 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v561 = *(*int32)(unsafe.Add(mBase, uint32(v560)+12))
	v562 = *(*int32)(unsafe.Add(mBase, uint32(v560)+20))
	v563 = int32(2)
	v566 = *(*int32)(unsafe.Add(mBase, uint32(v560)+4))
	v570 = *(*int32)(unsafe.Add(mBase, uint32(v561+v562<<(uint(v563)%32)+v566<<(uint(v563)%32))))
	v572 = v570 & int32(268435455)
	if int32(0) <= v570 {
		goto L116
	} else {
		goto L117
	}
L116:
	;
	v575 = *(*int32)(unsafe.Add(mBase, uint32(v560)+28))
	v577 = v575 + v572
	goto L118
L117:
	;
	v577 = v572
	goto L118
L118:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v560)+28)) = v577
	v579 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v580 = *(*int32)(unsafe.Add(mBase, uint32(v579)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v579)+20)) = v580 + int32(1)
	v584 = int32(2)
	if l2 != 0 {
		v626 = v584
		goto L1
	} else {
		goto L119
	}
L119:
	;
	v585 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	switch v585 {
	case 0, 1, 2, 3, 32:
		v626 = v584
		goto L1
	default:
		goto L7
	}
L120:
	;
	F_errmsg_internal(m, int32(_a_F_JsonbIteratorNext_0), int32(0))
	mBase = m.M
	v593 = m.ExcPending
	if v593 != 0 {
		goto L18
	} else {
		goto L121
	}
L121:
	;
	F_errfinish(m, int32(_a_F_JsonbIteratorNext_1), int32(1002), int32(_a_F_JsonbIteratorNext_2))
	mBase = m.M
	v598 = m.ExcPending
	if v598 != 0 {
		goto L18
	} else {
		goto L122
	}
L122:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L123:
	;
	F_errmsg_internal(m, int32(_a_F_JsonbIteratorNext_3), int32(0))
	mBase = m.M
	v606 = m.ExcPending
	if v606 != 0 {
		goto L18
	} else {
		goto L124
	}
L124:
	;
	F_errfinish(m, int32(_a_F_JsonbIteratorNext_1), int32(967), int32(_a_F_JsonbIteratorNext_2))
	mBase = m.M
	v611 = m.ExcPending
	if v611 != 0 {
		goto L18
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
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v617
	v12 = v617
	goto L5
}
func F_JsonbToCString(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v5 = F_JsonbToCStringWorker(m, l0, l1, l2, int32(0))
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
func F_jsonb_array_element(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
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
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = F_pg_detoast_datum(m, v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(v5)+4))
		if v9&int32(1073741824) == int32(0) {
			v14 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v14)
			return int32(0)
		} else {
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
			if v20 < int32(0) {
				v24 = v9 & int32(268435455)
				if base.Ui32(v24) < base.Ui32(int32(0)-v20) {
					v28 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v28)
					return int32(0)
				} else {
					v34 = v20 + v24
					v35 = F_getIthJsonbValueFromContainer(m, v5+int32(4), v34)
					mBase = m.M
					v36 = m.ExcPending
					if v36 != 0 {
						return int32(0)
					} else {
						if v35 != 0 {
							v37 = F_JsonbValueToJsonb(m, v35)
							mBase = m.M
							v38 = m.ExcPending
							if v38 != 0 {
								return int32(0)
							} else {
								return v37
							}
						} else {
							v40 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v40)
							return int32(0)
						}
					}
				}
			} else {
				v34 = v20
				v35 = F_getIthJsonbValueFromContainer(m, v5+int32(4), v34)
				mBase = m.M
				v36 = m.ExcPending
				if v36 != 0 {
					return int32(0)
				} else {
					if v35 != 0 {
						v37 = F_JsonbValueToJsonb(m, v35)
						mBase = m.M
						v38 = m.ExcPending
						if v38 != 0 {
							return int32(0)
						} else {
							return v37
						}
					} else {
						v40 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v40)
						return int32(0)
					}
				}
			}
		}
	}
}
func F_jsonb_array_element_text(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = F_pg_detoast_datum(m, v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(v5)+4))
		if v9&int32(1073741824) == int32(0) {
			v14 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v14)
			return int32(0)
		} else {
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
			if v20 < int32(0) {
				v24 = v9 & int32(268435455)
				if base.Ui32(v24) < base.Ui32(int32(0)-v20) {
					v28 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v28)
					return int32(0)
				} else {
					v34 = v20 + v24
					v35 = F_getIthJsonbValueFromContainer(m, v5+int32(4), v34)
					mBase = m.M
					v36 = m.ExcPending
					if v36 != 0 {
						return int32(0)
					} else {
						if v35 == int32(0) {
							v45 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v45)
							return int32(0)
						} else {
							v39 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
							if v39 == int32(0) {
								v45 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v45)
								return int32(0)
							} else {
								v42 = F_JsonbValueAsText(m, v35)
								mBase = m.M
								v43 = m.ExcPending
								if v43 != 0 {
									return int32(0)
								} else {
									return v42
								}
							}
						}
					}
				}
			} else {
				v34 = v20
				v35 = F_getIthJsonbValueFromContainer(m, v5+int32(4), v34)
				mBase = m.M
				v36 = m.ExcPending
				if v36 != 0 {
					return int32(0)
				} else {
					if v35 == int32(0) {
						v45 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v45)
						return int32(0)
					} else {
						v39 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
						if v39 == int32(0) {
							v45 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v45)
							return int32(0)
						} else {
							v42 = F_JsonbValueAsText(m, v35)
							mBase = m.M
							v43 = m.ExcPending
							if v43 != 0 {
								return int32(0)
							} else {
								return v42
							}
						}
					}
				}
			}
		}
	}
}
func F_jsonb_float4(m *base.Module, l0 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn13934(m, l0, int32(_a_F_jsonb_float4_0), int32(1319))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_jsonb_float8(m *base.Module, l0 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn13934(m, l0, int32(_a_F_jsonb_float8_0), int32(17))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_jsonb_in_scalar(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
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
	var v62 int32
	_ = v62
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	v5 = m.G0
	v7 = v5 - int32(48)
	m.G0 = v7
	switch l2 - int32(1) {
	case 0:
		*(*int32)(unsafe.Add(mBase, uint32(v7)+28)) = int32(1)
		v13 = F_strlen(m, l1)
		mBase = m.M
		*(*int32)(unsafe.Add(mBase, uint32(v7)+32)) = v13
		if base.Ui32(v13) < base.Ui32(int32(268435456)) {
			*(*int32)(unsafe.Add(mBase, uint32(v7)+36)) = l1
			v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			if v83 == int32(0) {
				v86 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(v7)+16)) = uint8(v86)
				*(*int64)(unsafe.Add(mBase, uint32(v7)+4)) = int64(4294967312)
				v90 = int32(4)
				v93 = F_pushJsonbValue(m, l0, v90, v7+v90)
				mBase = m.M
				v94 = m.ExcPending
				if v94 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v93
					v99 = F_pushJsonbValue(m, l0, int32(3), v7+int32(28))
					mBase = m.M
					v100 = m.ExcPending
					if v100 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v99
						v102 = int32(0)
						v105 = F_pushJsonbValue(m, l0, int32(5), v102)
						mBase = m.M
						v106 = m.ExcPending
						if v106 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v105
							v139 = v102
							m.G0 = v7 + int32(48)
							return v139
						}
					}
				}
			} else {
				v108 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
				switch v108 - int32(16) {
				case 0:
					v133 = F_pushJsonbValue(m, l0, int32(3), v7+int32(28))
					mBase = m.M
					v134 = m.ExcPending
					if v134 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v133
						v139 = int32(0)
						m.G0 = v7 + int32(48)
						return v139
					}
				case 1:
					v114 = F_pushJsonbValue(m, l0, int32(2), v7+int32(28))
					mBase = m.M
					v115 = m.ExcPending
					if v115 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v114
						v139 = int32(0)
						m.G0 = v7 + int32(48)
						return v139
					}
				default:
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v120 = m.ExcPending
					if v120 != 0 {
						return int32(0)
					} else {
						F_errmsg_internal(m, int32(_a_F_jsonb_in_scalar_0), int32(0))
						mBase = m.M
						v124 = m.ExcPending
						if v124 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_jsonb_in_scalar_1), int32(454), int32(_a_F_jsonb_in_scalar_2))
							mBase = m.M
							v129 = m.ExcPending
							if v129 != 0 {
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
			v17 = int32(23)
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v19 = F_errsave_start(m, v18)
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int32(0)
			} else {
				if v19 == int32(0) {
					v139 = v17
					m.G0 = v7 + int32(48)
					return v139
				} else {
					F_errcode(m, int32(261))
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(_a_F_jsonb_in_scalar_3), int32(0))
						mBase = m.M
						v31 = m.ExcPending
						if v31 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(268435455)
							F_errdetail(m, int32(_a_F_jsonb_in_scalar_4), v7)
							mBase = m.M
							v36 = m.ExcPending
							if v36 != 0 {
								return int32(0)
							} else {
								F_errsave_finish(m, v18, int32(_a_F_jsonb_in_scalar_1), int32(284), int32(_a_F_jsonb_in_scalar_5))
								mBase = m.M
								v41 = m.ExcPending
								if v41 != 0 {
									return int32(0)
								} else {
									v139 = v17
									m.G0 = v7 + int32(48)
									return v139
								}
							}
						}
					}
				}
			}
		}
	case 1:
		*(*int32)(unsafe.Add(mBase, uint32(v7)+28)) = int32(2)
		v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v50 = F_DirectInputFunctionCallSafe(m, int32(408), l1, int32(-1), v47, v7+int32(24))
		mBase = m.M
		v51 = m.ExcPending
		if v51 != 0 {
			return int32(0)
		} else {
			if v50 == int32(0) {
				v139 = int32(23)
				m.G0 = v7 + int32(48)
				return v139
			} else {
				v54 = *(*int32)(unsafe.Add(mBase, uint32(v7)+24))
				v55 = F_pg_detoast_datum(m, v54)
				mBase = m.M
				v56 = m.ExcPending
				if v56 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v7)+32)) = v55
					v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					if v83 == int32(0) {
						v86 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(v7)+16)) = uint8(v86)
						*(*int64)(unsafe.Add(mBase, uint32(v7)+4)) = int64(4294967312)
						v90 = int32(4)
						v93 = F_pushJsonbValue(m, l0, v90, v7+v90)
						mBase = m.M
						v94 = m.ExcPending
						if v94 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v93
							v99 = F_pushJsonbValue(m, l0, int32(3), v7+int32(28))
							mBase = m.M
							v100 = m.ExcPending
							if v100 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v99
								v102 = int32(0)
								v105 = F_pushJsonbValue(m, l0, int32(5), v102)
								mBase = m.M
								v106 = m.ExcPending
								if v106 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v105
									v139 = v102
									m.G0 = v7 + int32(48)
									return v139
								}
							}
						}
					} else {
						v108 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
						switch v108 - int32(16) {
						case 0:
							v133 = F_pushJsonbValue(m, l0, int32(3), v7+int32(28))
							mBase = m.M
							v134 = m.ExcPending
							if v134 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v133
								v139 = int32(0)
								m.G0 = v7 + int32(48)
								return v139
							}
						case 1:
							v114 = F_pushJsonbValue(m, l0, int32(2), v7+int32(28))
							mBase = m.M
							v115 = m.ExcPending
							if v115 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v114
								v139 = int32(0)
								m.G0 = v7 + int32(48)
								return v139
							}
						default:
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v120 = m.ExcPending
							if v120 != 0 {
								return int32(0)
							} else {
								F_errmsg_internal(m, int32(_a_F_jsonb_in_scalar_0), int32(0))
								mBase = m.M
								v124 = m.ExcPending
								if v124 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_jsonb_in_scalar_1), int32(454), int32(_a_F_jsonb_in_scalar_2))
									mBase = m.M
									v129 = m.ExcPending
									if v129 != 0 {
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
	default:
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v71 = m.ExcPending
		if v71 != 0 {
			return int32(0)
		} else {
			F_errmsg_internal(m, int32(_a_F_jsonb_in_scalar_6), int32(0))
			mBase = m.M
			v75 = m.ExcPending
			if v75 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(_a_F_jsonb_in_scalar_1), int32(424), int32(_a_F_jsonb_in_scalar_2))
				mBase = m.M
				v80 = m.ExcPending
				if v80 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	case 8:
		v58 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(v7)+32)) = uint8(v58)
		*(*int32)(unsafe.Add(mBase, uint32(v7)+28)) = int32(3)
		v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if v83 == int32(0) {
			v86 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(v7)+16)) = uint8(v86)
			*(*int64)(unsafe.Add(mBase, uint32(v7)+4)) = int64(4294967312)
			v90 = int32(4)
			v93 = F_pushJsonbValue(m, l0, v90, v7+v90)
			mBase = m.M
			v94 = m.ExcPending
			if v94 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v93
				v99 = F_pushJsonbValue(m, l0, int32(3), v7+int32(28))
				mBase = m.M
				v100 = m.ExcPending
				if v100 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v99
					v102 = int32(0)
					v105 = F_pushJsonbValue(m, l0, int32(5), v102)
					mBase = m.M
					v106 = m.ExcPending
					if v106 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v105
						v139 = v102
						m.G0 = v7 + int32(48)
						return v139
					}
				}
			}
		} else {
			v108 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
			switch v108 - int32(16) {
			case 0:
				v133 = F_pushJsonbValue(m, l0, int32(3), v7+int32(28))
				mBase = m.M
				v134 = m.ExcPending
				if v134 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v133
					v139 = int32(0)
					m.G0 = v7 + int32(48)
					return v139
				}
			case 1:
				v114 = F_pushJsonbValue(m, l0, int32(2), v7+int32(28))
				mBase = m.M
				v115 = m.ExcPending
				if v115 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v114
					v139 = int32(0)
					m.G0 = v7 + int32(48)
					return v139
				}
			default:
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v120 = m.ExcPending
				if v120 != 0 {
					return int32(0)
				} else {
					F_errmsg_internal(m, int32(_a_F_jsonb_in_scalar_0), int32(0))
					mBase = m.M
					v124 = m.ExcPending
					if v124 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_jsonb_in_scalar_1), int32(454), int32(_a_F_jsonb_in_scalar_2))
						mBase = m.M
						v129 = m.ExcPending
						if v129 != 0 {
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
	case 9:
		v62 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v7)+32)) = uint8(v62)
		*(*int32)(unsafe.Add(mBase, uint32(v7)+28)) = int32(3)
		v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if v83 == int32(0) {
			v86 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(v7)+16)) = uint8(v86)
			*(*int64)(unsafe.Add(mBase, uint32(v7)+4)) = int64(4294967312)
			v90 = int32(4)
			v93 = F_pushJsonbValue(m, l0, v90, v7+v90)
			mBase = m.M
			v94 = m.ExcPending
			if v94 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v93
				v99 = F_pushJsonbValue(m, l0, int32(3), v7+int32(28))
				mBase = m.M
				v100 = m.ExcPending
				if v100 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v99
					v102 = int32(0)
					v105 = F_pushJsonbValue(m, l0, int32(5), v102)
					mBase = m.M
					v106 = m.ExcPending
					if v106 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v105
						v139 = v102
						m.G0 = v7 + int32(48)
						return v139
					}
				}
			}
		} else {
			v108 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
			switch v108 - int32(16) {
			case 0:
				v133 = F_pushJsonbValue(m, l0, int32(3), v7+int32(28))
				mBase = m.M
				v134 = m.ExcPending
				if v134 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v133
					v139 = int32(0)
					m.G0 = v7 + int32(48)
					return v139
				}
			case 1:
				v114 = F_pushJsonbValue(m, l0, int32(2), v7+int32(28))
				mBase = m.M
				v115 = m.ExcPending
				if v115 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v114
					v139 = int32(0)
					m.G0 = v7 + int32(48)
					return v139
				}
			default:
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v120 = m.ExcPending
				if v120 != 0 {
					return int32(0)
				} else {
					F_errmsg_internal(m, int32(_a_F_jsonb_in_scalar_0), int32(0))
					mBase = m.M
					v124 = m.ExcPending
					if v124 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_jsonb_in_scalar_1), int32(454), int32(_a_F_jsonb_in_scalar_2))
						mBase = m.M
						v129 = m.ExcPending
						if v129 != 0 {
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
	case 10:
		*(*int32)(unsafe.Add(mBase, uint32(v7)+28)) = int32(0)
		v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if v83 == int32(0) {
			v86 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(v7)+16)) = uint8(v86)
			*(*int64)(unsafe.Add(mBase, uint32(v7)+4)) = int64(4294967312)
			v90 = int32(4)
			v93 = F_pushJsonbValue(m, l0, v90, v7+v90)
			mBase = m.M
			v94 = m.ExcPending
			if v94 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v93
				v99 = F_pushJsonbValue(m, l0, int32(3), v7+int32(28))
				mBase = m.M
				v100 = m.ExcPending
				if v100 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v99
					v102 = int32(0)
					v105 = F_pushJsonbValue(m, l0, int32(5), v102)
					mBase = m.M
					v106 = m.ExcPending
					if v106 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v105
						v139 = v102
						m.G0 = v7 + int32(48)
						return v139
					}
				}
			}
		} else {
			v108 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
			switch v108 - int32(16) {
			case 0:
				v133 = F_pushJsonbValue(m, l0, int32(3), v7+int32(28))
				mBase = m.M
				v134 = m.ExcPending
				if v134 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v133
					v139 = int32(0)
					m.G0 = v7 + int32(48)
					return v139
				}
			case 1:
				v114 = F_pushJsonbValue(m, l0, int32(2), v7+int32(28))
				mBase = m.M
				v115 = m.ExcPending
				if v115 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v114
					v139 = int32(0)
					m.G0 = v7 + int32(48)
					return v139
				}
			default:
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v120 = m.ExcPending
				if v120 != 0 {
					return int32(0)
				} else {
					F_errmsg_internal(m, int32(_a_F_jsonb_in_scalar_0), int32(0))
					mBase = m.M
					v124 = m.ExcPending
					if v124 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_jsonb_in_scalar_1), int32(454), int32(_a_F_jsonb_in_scalar_2))
						mBase = m.M
						v129 = m.ExcPending
						if v129 != 0 {
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
func F_jsonb_lt(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
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
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = F_pg_detoast_datum(m, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v13 = F_pg_detoast_datum(m, v12)
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			v17 = F_compareJsonbContainers(m, v6+int32(4), v13+int32(4))
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				if v19 != v6 {
					F_pfree(m, v6)
					mBase = m.M
					v22 = m.ExcPending
					if v22 != 0 {
						return int32(0)
					} else {
						v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
						if v23 != v13 {
							F_pfree(m, v13)
							mBase = m.M
							v26 = m.ExcPending
							if v26 != 0 {
								return int32(0)
							} else {
								return int32(base.Ui32(v17) >> (uint(int32(31)) % 32))
							}
						} else {
							return int32(base.Ui32(v17) >> (uint(int32(31)) % 32))
						}
					}
				} else {
					v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					if v23 != v13 {
						F_pfree(m, v13)
						mBase = m.M
						v26 = m.ExcPending
						if v26 != 0 {
							return int32(0)
						} else {
							return int32(base.Ui32(v17) >> (uint(int32(31)) % 32))
						}
					} else {
						return int32(base.Ui32(v17) >> (uint(int32(31)) % 32))
					}
				}
			}
		}
	}
}
func F_jsonb_object_agg_strict_transfn(m *base.Module, l0 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = F_jsonb_object_agg_transfn_worker(m, l0, int32(1), int32(0))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_jsonb_object_agg_unique_strict_transfn(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v2 = int32(1)
	v4 = F_jsonb_object_agg_transfn_worker(m, l0, v2, v2)
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_jsonb_object_agg_unique_transfn(m *base.Module, l0 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = F_jsonb_object_agg_transfn_worker(m, l0, int32(0), int32(1))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_jsonb_object_field_text(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
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
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v83 int32
	_ = v83
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = F_pg_detoast_datum(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v17 = F_pg_detoast_datum_packed(m, v16)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+7)))
			if v19&int32(32) == int32(0) {
				v76 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v76)
				v83 = int32(0)
				m.G0 = v9 + int32(32)
				return v83
			} else {
				v24 = int32(4)
				v26 = int32(1)
				v27 = v17 + v26
				v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
				v32 = v30 & v26
				if v32 != 0 {
					v33 = v27
				} else {
					v33 = v17 + v24
				}
				if v30 == int32(1) {
					v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27))))
					if v39 == int32(18) {
						v42 = int32(16)
					} else {
						v42 = int32(0)
					}
					if base.Ui32((v39-int32(1))&int32(255)) < base.Ui32(int32(3)) {
						v49 = int32(4)
					} else {
						v49 = v42
					}
					v60 = v49
				} else {
					v50 = int32(1)
					if v32 != 0 {
						v60 = int32(base.Ui32(v30)>>(uint(v50)%32)) - v50
					} else {
						v54 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
						v60 = int32(base.Ui32(v54)>>(uint(int32(2))%32)) - int32(4)
					}
				}
				v63 = F_getKeyJsonValueFromContainer(m, v12+v24, v33, v60, v9+int32(12))
				mBase = m.M
				v64 = m.ExcPending
				if v64 != 0 {
					return int32(0)
				} else {
					if v63 == int32(0) {
						v76 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v76)
						v83 = int32(0)
						m.G0 = v9 + int32(32)
						return v83
					} else {
						v67 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
						if v67 == int32(0) {
							v76 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v76)
							v83 = int32(0)
							m.G0 = v9 + int32(32)
							return v83
						} else {
							v70 = F_JsonbValueAsText(m, v63)
							mBase = m.M
							v71 = m.ExcPending
							if v71 != 0 {
								return int32(0)
							} else {
								v83 = v70
								m.G0 = v9 + int32(32)
								return v83
							}
						}
					}
				}
			}
		}
	}
}
func F_jsonb_path_query_array_internal(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
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
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
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
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = F_pg_detoast_datum(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v17 = F_pg_detoast_datum(m, v16)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v9))) = int64(0)
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v22 = F_pg_detoast_datum(m, v21)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v29 = F_executeJsonPath(m, v17, v22, int32(1394), int32(1395), v12, base.B2i32(v26 == int32(0)), v9, l1)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v31 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v31
	v37 = F_pushJsonbValue(m, v9+int32(12), int32(4), v31)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v39 = int32(0)
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	if v40 != 0 {
		v55 = v39
		v56 = v40
		v57 = int32(0)
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v58 = v55
	v59 = v56
	goto L15
L8:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	if v41 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v44 = int32(0)
	v55 = v39
	v56 = v44
	v57 = v44
	goto L7
L10:
	;
	goto L11
L11:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v41)+12))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v41)+4))
	if int32(1) < v50 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v53 = v46 + int32(4)
	goto L14
L13:
	;
	v53 = int32(0)
	goto L14
L14:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
	v55 = v53
	v56 = v54
	v57 = v41
	goto L7
L15:
	;
	if v58 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	v91 = F_pushJsonbValue(m, v9+int32(12), int32(5), int32(0))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L1
	} else {
		goto L28
	}
L17:
	;
	if v59 != 0 {
		goto L24
	} else {
		goto L25
	}
L18:
	;
	v66 = int32(0)
	v80 = v66
	v81 = v66
	goto L17
L19:
	;
	goto L20
L20:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v58)))
	v70 = v58 + int32(4)
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v57)+12))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
	if base.Ui32(v70) < base.Ui32(v72+v73<<(uint(int32(2))%32)) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v78 = v70
	goto L23
L22:
	;
	v78 = int32(0)
	goto L23
L23:
	;
	v80 = v68
	v81 = v78
	goto L17
L24:
	;
	v85 = F_pushJsonbValue(m, v9+int32(12), int32(3), v59)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L1
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	goto L16
L27:
	;
	v58 = v81
	v59 = v80
	goto L15
L28:
	;
	v93 = F_JsonbValueToJsonb(m, v91)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	m.G0 = v9 + int32(16)
	return v93
}
func F_jsonb_path_query_first_tz(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_jsonb_path_query_first_internal(m, l0, int32(1))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_jsonb_populate_record(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v3 = int32(0)
	v6 = F_populate_record_worker(m, l0, int32(_a_F_jsonb_populate_record_0), v3, int32(1), v3)
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return v6
	}
}
func F_jsonb_populate_recordset(m *base.Module, l0 int32) int32 {
	var v8 int32
	_ = v8
	F_populate_recordset_worker(m, l0, int32(_a_F_jsonb_populate_recordset_0), int32(0), int32(1))
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return int32(0)
	}
}
func F_jsonb_send(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
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
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
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
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = F_pg_detoast_datum(m, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		v13 = F_makeStringInfo(m)
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
			v21 = F_JsonbToCStringWorker(m, v13, v9+int32(4), int32(base.Ui32(v17)>>(uint(int32(2))%32)), int32(0))
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int32(0)
			} else {
				F_pq_begintypsend(m, v6)
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					F_enlargeStringInfo(m, v6, int32(1))
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return int32(0)
					} else {
						v28 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
						v29 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
						v31 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(v28+v29))) = uint8(v31)
						*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v28 + v31
						v36 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
						v37 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
						F_pq_sendtext(m, v6, v36, v37)
						mBase = m.M
						v39 = m.ExcPending
						if v39 != 0 {
							return int32(0)
						} else {
							F_free_attrmap(m, v13)
							mBase = m.M
							v41 = m.ExcPending
							if v41 != 0 {
								return int32(0)
							} else {
								v43 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
								v44 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
								*(*int32)(unsafe.Add(mBase, uint32(v43))) = v44 << (uint(int32(2)) % 32)
								m.G0 = v6 + int32(16)
								return v43
							}
						}
					}
				}
			}
		}
	}
}
func F_jsonb_subscript_fetch(m *base.Module, l0 int32, l1 int32, l2 int32) {
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
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)+4))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
	v8 = F_pg_detoast_datum(m, v7)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
		v11 = *(*int32)(unsafe.Add(mBase, uint32(v4)+8))
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
		v14 = F_jsonb_get_element(m, v8, v10, v11, v12, int32(0))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return
		} else {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			*(*int32)(unsafe.Add(mBase, uint32(v16))) = v14
			return
		}
	}
}
func F_jsonb_subscript_transform(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
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
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
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
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v187 int32
	_ = v187
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v222 int32
	_ = v222
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v242 int32
	_ = v242
	v6 = int32(0)
	v12 = m.G0
	v14 = v12 - int32(48)
	m.G0 = v14
	if l1 == v6 {
		v187 = v6
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L14
	} else {
		goto L62
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L14
	} else {
		goto L55
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v187
	*(*int64)(unsafe.Add(mBase, uint32(l0)+12)) = int64(-4294963494)
	m.G0 = v14 + int32(48)
	return
L4:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v18 <= int32(0) {
		v187 = v6
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v24 = v14 + int32(32) | int32(4)
	v34 = v6
	v35 = v6
	goto L6
L6:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v36+v34<<(uint(int32(2))%32))))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+12))
	if l3 != 0 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v187 = v171
	goto L3
L8:
	;
	if v41 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L10
L10:
	;
	if v41 != 0 {
		goto L25
	} else {
		goto L26
	}
L11:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v40)+8))
	v45 = v44
	goto L13
L12:
	;
	v45 = v41
	goto L13
L13:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	return
L15:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	F_errmsg(m, int32(_a_F_jsonb_subscript_transform_0), int32(0))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L14
	} else {
		goto L17
	}
L17:
	;
	v57 = F_exprLocation(m, v45)
	mBase = m.M
	F_parser_errposition(m, l2, v57)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L14
	} else {
		goto L18
	}
L18:
	;
	F_errfinish(m, int32(_a_F_jsonb_subscript_transform_1), int32(68), int32(_a_F_jsonb_subscript_transform_2))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L14
	} else {
		goto L19
	}
L19:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L20:
	;
	v163 = int32(-1)
	v167 = F_coerce_type(m, l2, v68, v162, v160, v163, int32(0), int32(2), v163)
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L14
	} else {
		goto L51
	}
L21:
	;
	if v156 == int32(705) {
		goto L2
	} else {
		goto L50
	}
L22:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v14)+36))
	v156 = v155
	goto L21
L23:
	;
	v150 = F_can_coerce_type(m, int32(1), v14+int32(44), v24, int32(0))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L14
	} else {
		goto L48
	}
L24:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L14
	} else {
		goto L41
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+44)) = int32(0)
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l2)+68))
	v68 = F_transformExpr(m, l2, v41, v67)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L14
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L14
	} else {
		goto L36
	}
L28:
	;
	v70 = F_exprType(m, v68)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L14
	} else {
		goto L29
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+44)) = v70
	v74 = int32(705)
	if v70 == v74 {
		v160 = int32(25)
		v162 = v74
		goto L20
	} else {
		goto L30
	}
L30:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v14)+32)) = int64(107374182423)
	v81 = v14 + int32(44)
	v85 = F_can_coerce_type(m, int32(1), v81, v14+int32(32), int32(0))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L14
	} else {
		goto L31
	}
L31:
	;
	if v85 == int32(0) {
		goto L23
	} else {
		goto L32
	}
L32:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v14)+32))
	v92 = F_can_coerce_type(m, int32(1), v81, v24, int32(0))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L14
	} else {
		goto L33
	}
L33:
	;
	if v92 == int32(0) {
		v156 = v89
		goto L21
	} else {
		goto L34
	}
L34:
	;
	if v89 != int32(705) {
		goto L24
	} else {
		goto L35
	}
L35:
	;
	goto L22
L36:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L14
	} else {
		goto L37
	}
L37:
	;
	F_errmsg(m, int32(_a_F_jsonb_subscript_transform_0), int32(0))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L14
	} else {
		goto L38
	}
L38:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v40)+12))
	v110 = F_exprLocation(m, v109)
	mBase = m.M
	F_parser_errposition(m, l2, v110)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L14
	} else {
		goto L39
	}
L39:
	;
	F_errfinish(m, int32(_a_F_jsonb_subscript_transform_1), int32(149), int32(_a_F_jsonb_subscript_transform_2))
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L14
	} else {
		goto L40
	}
L40:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L41:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L14
	} else {
		goto L42
	}
L42:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v14)+44))
	v126 = F_format_type_be(m, v125)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L14
	} else {
		goto L43
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v126
	F_errmsg(m, int32(_a_F_jsonb_subscript_transform_3), v14+int32(16))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L14
	} else {
		goto L44
	}
L44:
	;
	F_errhint(m, int32(_a_F_jsonb_subscript_transform_4), int32(0))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L14
	} else {
		goto L45
	}
L45:
	;
	v138 = F_exprLocation(m, v68)
	mBase = m.M
	F_parser_errposition(m, l2, v138)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L14
	} else {
		goto L46
	}
L46:
	;
	F_errfinish(m, int32(_a_F_jsonb_subscript_transform_1), int32(102), int32(_a_F_jsonb_subscript_transform_2))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L14
	} else {
		goto L47
	}
L47:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L48:
	;
	if v150 == int32(0) {
		goto L2
	} else {
		goto L49
	}
L49:
	;
	goto L22
L50:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v14)+44))
	v160 = v156
	v162 = v159
	goto L20
L51:
	;
	if v167 == int32(0) {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	v171 = F_lappend(m, v35, v167)
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L14
	} else {
		goto L53
	}
L53:
	;
	v174 = v34 + int32(1)
	v175 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v174 < v175 {
		v34 = v174
		v35 = v171
		goto L6
	} else {
		goto L54
	}
L54:
	;
	goto L7
L55:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L14
	} else {
		goto L56
	}
L56:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v14)+44))
	v205 = F_format_type_be(m, v204)
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L14
	} else {
		goto L57
	}
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v205
	F_errmsg(m, int32(_a_F_jsonb_subscript_transform_3), v14)
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L14
	} else {
		goto L58
	}
L58:
	;
	F_errhint(m, int32(_a_F_jsonb_subscript_transform_5), int32(0))
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L14
	} else {
		goto L59
	}
L59:
	;
	v215 = F_exprLocation(m, v68)
	mBase = m.M
	F_parser_errposition(m, l2, v215)
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L14
	} else {
		goto L60
	}
L60:
	;
	F_errfinish(m, int32(_a_F_jsonb_subscript_transform_1), int32(116), int32(_a_F_jsonb_subscript_transform_2))
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L14
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
	F_errcode(m, int32(67141764))
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L14
	} else {
		goto L63
	}
L63:
	;
	F_errmsg(m, int32(_a_F_jsonb_subscript_transform_6), int32(0))
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L14
	} else {
		goto L64
	}
L64:
	;
	v235 = F_exprLocation(m, int32(0))
	mBase = m.M
	F_parser_errposition(m, l2, v235)
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L14
	} else {
		goto L65
	}
L65:
	;
	F_errfinish(m, int32(_a_F_jsonb_subscript_transform_1), int32(137), int32(_a_F_jsonb_subscript_transform_2))
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L14
	} else {
		goto L66
	}
L66:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_parse_jsonb_index_flags(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
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
	var v80 int32
	_ = v80
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v138 int32
	_ = v138
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v169 int32
	_ = v169
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v199 int32
	_ = v199
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v258 int32
	_ = v258
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v289 int32
	_ = v289
	var v294 int32
	_ = v294
	var v297 int32
	_ = v297
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v318 int32
	_ = v318
	var v327 int32
	_ = v327
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v349 int32
	_ = v349
	var v355 int32
	_ = v355
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v366 int32
	_ = v366
	var v370 int32
	_ = v370
	var v375 int32
	_ = v375
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v390 int32
	_ = v390
	var v393 int32
	_ = v393
	var v397 int32
	_ = v397
	var v401 int32
	_ = v401
	var v406 int32
	_ = v406
	var v412 int32
	_ = v412
	var v416 int32
	_ = v416
	var v421 int32
	_ = v421
	var v425 int32
	_ = v425
	var v428 int32
	_ = v428
	var v432 int32
	_ = v432
	var v437 int32
	_ = v437
	v4 = m.G0
	v6 = v4 - int32(32)
	m.G0 = v6
	v10 = F_JsonbIteratorInit(m, l0+int32(4))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6)+28)) = v10
	v20 = F_JsonbIteratorNext(m, v6+int32(28), v6+int32(8), int32(0))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	if v20 == int32(4) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v25 = int32(0)
	goto L7
L5:
	;
	goto L6
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v425 = m.ExcPending
	if v425 != 0 {
		goto L1
	} else {
		goto L136
	}
L7:
	;
	v33 = F_JsonbIteratorNext(m, v6+int32(28), v6+int32(8), int32(0))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L13
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L1
	} else {
		goto L133
	}
L9:
	;
	goto L8
L10:
	;
	v25 = v25 | int32(2)
	goto L7
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L1
	} else {
		goto L128
	}
L12:
	;
	v381 = F_JsonbIteratorNext(m, v6+int32(28), v6+int32(8), int32(0))
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L1
	} else {
		goto L126
	}
L13:
	;
	if v33 != int32(3) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	if v33 == int32(5) {
		goto L12
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
	if v52 != int32(1) {
		goto L11
	} else {
		goto L21
	}
L17:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	F_errmsg_internal(m, int32(_a_F_parse_jsonb_index_flags_0), int32(0))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	F_errfinish(m, int32(_a_F_parse_jsonb_index_flags_1), int32(_a_F_parse_jsonb_index_flags_2), int32(_a_F_parse_jsonb_index_flags_3))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L21:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
	if v55 != int32(3) {
		v175 = v55
		goto L22
	} else {
		goto L23
	}
L22:
	;
	if v175 == int32(6) {
		goto L62
	} else {
		goto L63
	}
L23:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v6)+16))
	v64 = v59
	v65 = int32(_a_F_parse_jsonb_index_flags_4)
	v66 = int32(3)
	goto L25
L24:
	;
	if v111 == int32(0) {
		v25 = int32(15)
		goto L7
	} else {
		goto L40
	}
L25:
	;
	if v66 != 0 {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v111 = int32(0)
	goto L24
L27:
	;
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64))))
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65))))
	if v69 == v70 {
		v92 = v69
		goto L30
	} else {
		goto L31
	}
L28:
	;
	goto L29
L29:
	;
	goto L26
L30:
	;
	v94 = int32(1)
	if v92 != 0 {
		v64 = v64 + v94
		v65 = v65 + v94
		v66 = v66 - v94
		goto L25
	} else {
		goto L39
	}
L31:
	;
	if base.Ui32((v69-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v80 = v69 | int32(32)
	goto L34
L33:
	;
	v80 = v69
	goto L34
L34:
	;
	if base.Ui32((v70-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v89 = v70 | int32(32)
	goto L37
L36:
	;
	v89 = v70
	goto L37
L37:
	;
	if v80 == v89 {
		v92 = v80
		goto L30
	} else {
		goto L38
	}
L38:
	;
	v111 = v80 - v89
	goto L24
L39:
	;
	goto L29
L40:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
	if v114 != int32(3) {
		v175 = v114
		goto L22
	} else {
		goto L41
	}
L41:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v6)+16))
	v122 = v117
	v123 = int32(_a_F_parse_jsonb_index_flags_5)
	v124 = int32(3)
	goto L43
L42:
	;
	if v169 == int32(0) {
		goto L58
	} else {
		goto L59
	}
L43:
	;
	if v124 != 0 {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	v169 = int32(0)
	goto L42
L45:
	;
	v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122))))
	v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123))))
	if v127 == v128 {
		v150 = v127
		goto L48
	} else {
		goto L49
	}
L46:
	;
	goto L47
L47:
	;
	goto L44
L48:
	;
	v152 = int32(1)
	if v150 != 0 {
		v122 = v122 + v152
		v123 = v123 + v152
		v124 = v124 - v152
		goto L43
	} else {
		goto L57
	}
L49:
	;
	if base.Ui32((v127-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v138 = v127 | int32(32)
	goto L52
L51:
	;
	v138 = v127
	goto L52
L52:
	;
	if base.Ui32((v128-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v147 = v128 | int32(32)
	goto L55
L54:
	;
	v147 = v128
	goto L55
L55:
	;
	if v138 == v147 {
		v150 = v138
		goto L48
	} else {
		goto L56
	}
L56:
	;
	v169 = v138 - v147
	goto L42
L57:
	;
	goto L47
L58:
	;
	v25 = v25 | int32(1)
	goto L7
L59:
	;
	goto L60
L60:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
	v175 = v174
	goto L22
L61:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L1
	} else {
		goto L120
	}
L62:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v6)+16))
	v183 = v178
	v184 = int32(_a_F_parse_jsonb_index_flags_6)
	v185 = int32(6)
	goto L66
L63:
	;
	v234 = v175
	goto L64
L64:
	;
	if v234 != int32(7) {
		goto L61
	} else {
		goto L82
	}
L65:
	;
	if v230 == int32(0) {
		goto L10
	} else {
		goto L81
	}
L66:
	;
	if v185 != 0 {
		goto L68
	} else {
		goto L69
	}
L67:
	;
	v230 = int32(0)
	goto L65
L68:
	;
	v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v183))))
	v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v184))))
	if v188 == v189 {
		v211 = v188
		goto L71
	} else {
		goto L72
	}
L69:
	;
	goto L70
L70:
	;
	goto L67
L71:
	;
	v213 = int32(1)
	if v211 != 0 {
		v183 = v183 + v213
		v184 = v184 + v213
		v185 = v185 - v213
		goto L66
	} else {
		goto L80
	}
L72:
	;
	if base.Ui32((v188-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v199 = v188 | int32(32)
	goto L75
L74:
	;
	v199 = v188
	goto L75
L75:
	;
	if base.Ui32((v189-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v208 = v189 | int32(32)
	goto L78
L77:
	;
	v208 = v189
	goto L78
L78:
	;
	if v199 == v208 {
		v211 = v199
		goto L71
	} else {
		goto L79
	}
L79:
	;
	v230 = v199 - v208
	goto L65
L80:
	;
	goto L70
L81:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
	v234 = v233
	goto L64
L82:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v6)+16))
	v242 = v237
	v243 = int32(_a_F_parse_jsonb_index_flags_7)
	v244 = int32(7)
	goto L84
L83:
	;
	if v289 == int32(0) {
		goto L99
	} else {
		goto L100
	}
L84:
	;
	if v244 != 0 {
		goto L86
	} else {
		goto L87
	}
L85:
	;
	v289 = int32(0)
	goto L83
L86:
	;
	v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v242))))
	v248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v243))))
	if v247 == v248 {
		v270 = v247
		goto L89
	} else {
		goto L90
	}
L87:
	;
	goto L88
L88:
	;
	goto L85
L89:
	;
	v272 = int32(1)
	if v270 != 0 {
		v242 = v242 + v272
		v243 = v243 + v272
		v244 = v244 - v272
		goto L84
	} else {
		goto L98
	}
L90:
	;
	if base.Ui32((v247-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	v258 = v247 | int32(32)
	goto L93
L92:
	;
	v258 = v247
	goto L93
L93:
	;
	if base.Ui32((v248-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	v267 = v248 | int32(32)
	goto L96
L95:
	;
	v267 = v248
	goto L96
L96:
	;
	if v258 == v267 {
		v270 = v258
		goto L89
	} else {
		goto L97
	}
L97:
	;
	v289 = v258 - v267
	goto L83
L98:
	;
	goto L88
L99:
	;
	v25 = v25 | int32(4)
	goto L7
L100:
	;
	goto L101
L101:
	;
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
	if v294 != int32(7) {
		goto L61
	} else {
		goto L102
	}
L102:
	;
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v6)+16))
	v302 = v297
	v303 = int32(_a_F_parse_jsonb_index_flags_8)
	v304 = int32(7)
	goto L104
L103:
	;
	if v349 != 0 {
		goto L61
	} else {
		goto L119
	}
L104:
	;
	if v304 != 0 {
		goto L106
	} else {
		goto L107
	}
L105:
	;
	v349 = int32(0)
	goto L103
L106:
	;
	v307 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v302))))
	v308 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v303))))
	if v307 == v308 {
		v330 = v307
		goto L109
	} else {
		goto L110
	}
L107:
	;
	goto L108
L108:
	;
	goto L105
L109:
	;
	v332 = int32(1)
	if v330 != 0 {
		v302 = v302 + v332
		v303 = v303 + v332
		v304 = v304 - v332
		goto L104
	} else {
		goto L118
	}
L110:
	;
	if base.Ui32((v307-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L111
	} else {
		goto L112
	}
L111:
	;
	v318 = v307 | int32(32)
	goto L113
L112:
	;
	v318 = v307
	goto L113
L113:
	;
	if base.Ui32((v308-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L114
	} else {
		goto L115
	}
L114:
	;
	v327 = v308 | int32(32)
	goto L116
L115:
	;
	v327 = v308
	goto L116
L116:
	;
	if v318 == v327 {
		v330 = v318
		goto L109
	} else {
		goto L117
	}
L117:
	;
	v349 = v318 - v327
	goto L103
L118:
	;
	goto L108
L119:
	;
	v25 = v25 | int32(8)
	goto L7
L120:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L1
	} else {
		goto L121
	}
L121:
	;
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v6)+16))
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
	v361 = F_pnstrdup(m, v359, v360)
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L1
	} else {
		goto L122
	}
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6))) = v361
	F_errmsg(m, int32(_a_F_parse_jsonb_index_flags_9), v6)
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L1
	} else {
		goto L123
	}
L123:
	;
	F_errhint(m, int32(_a_F_parse_jsonb_index_flags_10), int32(0))
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
		goto L1
	} else {
		goto L124
	}
L124:
	;
	F_errfinish(m, int32(_a_F_parse_jsonb_index_flags_1), int32(_a_F_parse_jsonb_index_flags_11), int32(_a_F_parse_jsonb_index_flags_3))
	mBase = m.M
	v375 = m.ExcPending
	if v375 != 0 {
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
	if v381 != 0 {
		goto L9
	} else {
		goto L127
	}
L127:
	;
	m.G0 = v6 + int32(32)
	return v25
L128:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L1
	} else {
		goto L129
	}
L129:
	;
	F_errmsg(m, int32(_a_F_parse_jsonb_index_flags_12), int32(0))
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
		goto L1
	} else {
		goto L130
	}
L130:
	;
	F_errhint(m, int32(_a_F_parse_jsonb_index_flags_10), int32(0))
	mBase = m.M
	v401 = m.ExcPending
	if v401 != 0 {
		goto L1
	} else {
		goto L131
	}
L131:
	;
	F_errfinish(m, int32(_a_F_parse_jsonb_index_flags_1), int32(_a_F_parse_jsonb_index_flags_13), int32(_a_F_parse_jsonb_index_flags_3))
	mBase = m.M
	v406 = m.ExcPending
	if v406 != 0 {
		goto L1
	} else {
		goto L132
	}
L132:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L133:
	;
	F_errmsg_internal(m, int32(_a_F_parse_jsonb_index_flags_0), int32(0))
	mBase = m.M
	v416 = m.ExcPending
	if v416 != 0 {
		goto L1
	} else {
		goto L134
	}
L134:
	;
	F_errfinish(m, int32(_a_F_parse_jsonb_index_flags_1), int32(_a_F_parse_jsonb_index_flags_14), int32(_a_F_parse_jsonb_index_flags_3))
	mBase = m.M
	v421 = m.ExcPending
	if v421 != 0 {
		goto L1
	} else {
		goto L135
	}
L135:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L136:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v428 = m.ExcPending
	if v428 != 0 {
		goto L1
	} else {
		goto L137
	}
L137:
	;
	F_errmsg(m, int32(_a_F_parse_jsonb_index_flags_15), int32(0))
	mBase = m.M
	v432 = m.ExcPending
	if v432 != 0 {
		goto L1
	} else {
		goto L138
	}
L138:
	;
	F_errfinish(m, int32(_a_F_parse_jsonb_index_flags_1), int32(_a_F_parse_jsonb_index_flags_16), int32(_a_F_parse_jsonb_index_flags_3))
	mBase = m.M
	v437 = m.ExcPending
	if v437 != 0 {
		goto L1
	} else {
		goto L139
	}
L139:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
