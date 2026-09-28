package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_tts_buffer_is_current_xact_tuple(m *base.Module, l0 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn14395(m, l0, int32(_a_F_tts_buffer_is_current_xact_tuple_0), int32(796))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_tts_heap_copy_minimal_tuple(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
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
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v5 != 0 {
		v33 = v5
		v35 = F_minimal_tuple_from_heap_tuple(m, v33, l1)
		mBase = m.M
		v36 = m.ExcPending
		if v36 != 0 {
			return int32(0)
		} else {
			return v35
		}
	} else {
		v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)))
		if v7&int32(4) != 0 {
			v33 = int32(0)
			v35 = F_minimal_tuple_from_heap_tuple(m, v33, l1)
			mBase = m.M
			v36 = m.ExcPending
			if v36 != 0 {
				return int32(0)
			} else {
				return v35
			}
		} else {
			v10 = int32(_a_F_tts_heap_copy_minimal_tuple_0)
			v11 = *(*int32)(unsafe.Add(mBase, _c_F_tts_heap_copy_minimal_tuple[0]))
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
			*(*int32)(unsafe.Add(mBase, _c_F_tts_heap_copy_minimal_tuple[0])) = v13
			v15 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = v15
			*(*uint16)(unsafe.Add(mBase, uint32(l0)+6)) = uint16(v15)
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			v22 = F_heap_form_tuple(m, v19, v20, v21)
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v22
				v27 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
				v29 = v27 | int32(4)
				*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)) = uint16(v29)
				*(*int32)(unsafe.Add(mBase, _c_F_tts_heap_copy_minimal_tuple[0])) = v11
				v33 = v22
				v35 = F_minimal_tuple_from_heap_tuple(m, v33, l1)
				mBase = m.M
				v36 = m.ExcPending
				if v36 != 0 {
					return int32(0)
				} else {
					return v35
				}
			}
		}
	}
}
func F_tts_heap_getsomeattrs(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v19 int32
	_ = v19
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
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v99 int32
	_ = v99
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v221 int64
	_ = v221
	var v222 int64
	_ = v222
	var v223 int64
	_ = v223
	var v224 int64
	_ = v224
	var v225 int64
	_ = v225
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v241 int32
	_ = v241
	var v254 int32
	_ = v254
	var v262 int32
	_ = v262
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v291 int64
	_ = v291
	var v292 int64
	_ = v292
	var v293 int64
	_ = v293
	var v294 int64
	_ = v294
	var v296 int64
	_ = v296
	var v299 int32
	_ = v299
	var v305 int32
	_ = v305
	var v324 int32
	_ = v324
	var v341 int32
	_ = v341
	var v344 int32
	_ = v344
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
	var v354 int32
	_ = v354
	var v360 int32
	_ = v360
	var v363 int32
	_ = v363
	var v366 int64
	_ = v366
	var v367 int64
	_ = v367
	var v368 int64
	_ = v368
	var v369 int64
	_ = v369
	var v371 int32
	_ = v371
	var v375 int32
	_ = v375
	var v385 int32
	_ = v385
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v395 int32
	_ = v395
	var v397 int32
	_ = v397
	var v400 int32
	_ = v400
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v412 int32
	_ = v412
	var v416 int32
	_ = v416
	var v421 int32
	_ = v421
	var v425 int32
	_ = v425
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v438 int64
	_ = v438
	var v441 int32
	_ = v441
	var v445 int32
	_ = v445
	var v464 int32
	_ = v464
	var v484 int32
	_ = v484
	var v487 int32
	_ = v487
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v496 int32
	_ = v496
	var v502 int32
	_ = v502
	var v505 int32
	_ = v505
	var v508 int64
	_ = v508
	var v509 int64
	_ = v509
	var v510 int64
	_ = v510
	var v511 int64
	_ = v511
	var v513 int32
	_ = v513
	var v517 int32
	_ = v517
	var v527 int32
	_ = v527
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v537 int32
	_ = v537
	var v539 int32
	_ = v539
	var v542 int32
	_ = v542
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v554 int32
	_ = v554
	var v558 int32
	_ = v558
	var v563 int32
	_ = v563
	var v567 int32
	_ = v567
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v580 int64
	_ = v580
	var v582 int64
	_ = v582
	var v585 int32
	_ = v585
	var v589 int32
	_ = v589
	var v605 int32
	_ = v605
	var v608 int32
	_ = v608
	var v612 int32
	_ = v612
	var v616 int32
	_ = v616
	v2 = l1
	v3 = int32(0)
	v19 = m.G0
	v21 = v19 - int32(16)
	m.G0 = v21
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v2 < v23 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v25 = v2
	goto L3
L2:
	;
	v25 = v23
	goto L3
L3:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+16))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+16))
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+20)))
	if v31&int32(1) != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	v185 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+6)))
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+6)) = uint16(v2)
	v188 = v26 + int32(28)
	v189 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if base.Ui32(v185) < base.Ui32(v25) {
		goto L35
	} else {
		goto L36
	}
L5:
	;
	v171 = v165
	v172 = v165
	v181 = v166
	goto L4
L6:
	;
	v34 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v30)+18)))
	v36 = v34 & int32(2047)
	v45 = v30 + (int32(base.Ui32(v36+int32(7))>>(uint(int32(3))%32))+int32(30))&int32(1016)
	if v2 < v36 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	v158 = v30 + int32(24)
	if v2 <= v23 {
		goto L26
	} else {
		goto L27
	}
L9:
	;
	v47 = v2
	goto L11
L10:
	;
	v47 = v36
	goto L11
L11:
	;
	if v47 <= v23 {
		v165 = v47
		v166 = v45
		goto L5
	} else {
		goto L12
	}
L12:
	;
	v50 = v30 + int32(23)
	v52 = v47 >> (uint(int32(3)) % 32)
	if v52 <= int32(0) {
		v83 = v3
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83+v50))))
	v105 = base.I32_ctz(v99^int32(-1)) + v83<<(uint(int32(3))%32)
	if v105 < v47 {
		goto L19
	} else {
		goto L20
	}
L14:
	;
	v58 = v3
	goto L15
L15:
	;
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58+v50))))
	if v74 != int32(255) {
		v83 = v58
		goto L13
	} else {
		goto L17
	}
L16:
	;
	v83 = v52
	goto L13
L17:
	;
	v78 = v58 + int32(1)
	if v78 != v52 {
		v58 = v78
		goto L15
	} else {
		goto L18
	}
L18:
	;
	goto L16
L19:
	;
	v107 = v105
	goto L21
L20:
	;
	v107 = v47
	goto L21
L21:
	;
	v111 = (v47 + int32(7)) >> (uint(int32(3)) % 32)
	if v111 <= int32(0) {
		v171 = v47
		v172 = v107
		v181 = v45
		goto L4
	} else {
		goto L22
	}
L22:
	;
	v116 = v3
	v117 = v28
	goto L23
L23:
	;
	v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116+v50))))
	v135 = v133 ^ int32(255)
	v138 = int32(_a_F_tts_heap_getsomeattrs_0)
	*(*int64)(unsafe.Add(mBase, uint32(v117))) = (base.I64_extend_i32_u(int32(base.Ui32(v135)>>(uint(int32(4))%32))*v138)<<(uint(int64(32))%64) | base.I64_extend_i32_u(v135&int32(15)*v138)) & int64(72340172838076673)
	v155 = v116 + int32(1)
	if v155 != v111 {
		v116 = v155
		v117 = v117 + int32(8)
		goto L23
	} else {
		goto L25
	}
L24:
	;
	v171 = v47
	v172 = v107
	v181 = v45
	goto L4
L25:
	;
	goto L24
L26:
	;
	v171 = v2
	v172 = v2
	v181 = v158
	goto L4
L27:
	;
	goto L28
L28:
	;
	v160 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v30)+18)))
	v162 = v160 & int32(2047)
	if v2 < v162 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v164 = v2
	goto L31
L30:
	;
	v164 = v162
	goto L31
L31:
	;
	v165 = v164
	v166 = v158
	goto L5
L32:
	;
	m.G0 = v21 + int32(16)
	return
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = v616
	goto L32
L34:
	;
	if v27 < v172 {
		goto L48
	} else {
		goto L49
	}
L35:
	;
	v193 = v185
	goto L38
L36:
	;
	goto L37
L37:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+12)) = v233
	v241 = v185
	goto L34
L38:
	;
	v210 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v193+v28))) = uint8(v210)
	v213 = v193 << (uint(int32(3)) % 32)
	v214 = v188 + v213
	v215 = int32(*(*int16)(unsafe.Add(mBase, uint32(v214))))
	v216 = v181 + v215
	v218 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v214)+2)))
	switch v218 - int32(1) {
	case 0:
		goto L42
	case 1:
		goto L43
	default:
		goto L41
	case 3:
		goto L44
	}
L39:
	;
	v230 = v218 + v215
	*(*int32)(unsafe.Add(mBase, uint32(v21)+12)) = v230
	if v23 < v2 {
		v241 = v25
		goto L34
	} else {
		goto L46
	}
L40:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v189+v213))) = v225
	v228 = v193 + int32(1)
	if v228 != v25 {
		v193 = v228
		goto L38
	} else {
		goto L45
	}
L41:
	;
	v224 = *(*int64)(unsafe.Add(mBase, uint32(v216)))
	v225 = v224
	goto L40
L42:
	;
	v223 = int64(*(*int8)(unsafe.Add(mBase, uint32(v216))))
	v225 = v223
	goto L40
L43:
	;
	v222 = int64(*(*int16)(unsafe.Add(mBase, uint32(v216))))
	v225 = v222
	goto L40
L44:
	;
	v221 = int64(*(*int32)(unsafe.Add(mBase, uint32(v216))))
	v225 = v221
	goto L40
L45:
	;
	goto L39
L46:
	;
	v616 = v230
	goto L33
L47:
	;
	if base.Ui32(v305) < base.Ui32(v172) {
		goto L65
	} else {
		goto L66
	}
L48:
	;
	v254 = v27
	goto L50
L49:
	;
	v254 = v172
	goto L50
L50:
	;
	if base.Ui32(v254) <= base.Ui32(v241) {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v305 = v241
	goto L47
L52:
	;
	goto L53
L53:
	;
	v262 = v241
	goto L54
L54:
	;
	v275 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v262+v28))) = uint8(v275)
	v278 = v262 << (uint(int32(3)) % 32)
	v279 = v188 + v278
	v280 = int32(*(*int16)(unsafe.Add(mBase, uint32(v279))))
	v281 = v181 + v280
	v282 = int32(*(*int16)(unsafe.Add(mBase, uint32(v279)+2)))
	v284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v279)+4)))
	if v284 == int32(1) {
		goto L57
	} else {
		goto L58
	}
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+12)) = v280 + v282
	v305 = v254
	goto L47
L56:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v189+v278))) = v296
	v299 = v262 + int32(1)
	if v299 != v254 {
		v262 = v299
		goto L54
	} else {
		goto L64
	}
L57:
	;
	switch v282&int32(_a_F_tts_heap_getsomeattrs_1) - int32(1) {
	case 0:
		goto L61
	case 1:
		goto L62
	default:
		goto L60
	case 3:
		goto L63
	}
L58:
	;
	goto L59
L59:
	;
	v296 = base.I64_extend_i32_u(v281)
	goto L56
L60:
	;
	v294 = *(*int64)(unsafe.Add(mBase, uint32(v281)))
	v296 = v294
	goto L56
L61:
	;
	v293 = int64(*(*int8)(unsafe.Add(mBase, uint32(v281))))
	v296 = v293
	goto L56
L62:
	;
	v292 = int64(*(*int16)(unsafe.Add(mBase, uint32(v281))))
	v296 = v292
	goto L56
L63:
	;
	v291 = int64(*(*int32)(unsafe.Add(mBase, uint32(v281))))
	v296 = v291
	goto L56
L64:
	;
	goto L55
L65:
	;
	v324 = v305
	goto L68
L66:
	;
	v445 = v305
	goto L67
L67:
	;
	if base.Ui32(v445) < base.Ui32(v171) {
		goto L99
	} else {
		goto L100
	}
L68:
	;
	v341 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v324+v28))) = uint8(v341)
	v344 = v324 << (uint(int32(3)) % 32)
	v347 = v21 + int32(12)
	v348 = v344 + v188
	v349 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v348)+4)))
	v350 = int32(*(*int16)(unsafe.Add(mBase, uint32(v348)+2)))
	v351 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v348)+5)))
	if v341 < v350 {
		goto L71
	} else {
		goto L72
	}
L69:
	;
	v445 = v172
	goto L67
L70:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v189+v344))) = v438
	v441 = v324 + int32(1)
	if v441 != v172 {
		v324 = v441
		goto L68
	} else {
		goto L98
	}
L71:
	;
	v354 = *(*int32)(unsafe.Add(mBase, uint32(v347)))
	v360 = (v351 + v354 - int32(1)) & (int32(0) - v351)
	*(*int32)(unsafe.Add(mBase, uint32(v347))) = v360 + v350
	v363 = v181 + v360
	if v349 != 0 {
		goto L74
	} else {
		goto L75
	}
L72:
	;
	goto L73
L73:
	;
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v347)))
	if v350 == int32(-1) {
		goto L81
	} else {
		goto L82
	}
L74:
	;
	switch v350 - int32(1) {
	case 0:
		goto L80
	case 1:
		goto L79
	default:
		goto L77
	case 3:
		goto L78
	}
L75:
	;
	goto L76
L76:
	;
	v438 = base.I64_extend_i32_u(v363)
	goto L70
L77:
	;
	v369 = *(*int64)(unsafe.Add(mBase, uint32(v363)))
	v438 = v369
	goto L70
L78:
	;
	v368 = int64(*(*int32)(unsafe.Add(mBase, uint32(v363))))
	v438 = v368
	goto L70
L79:
	;
	v367 = int64(*(*int16)(unsafe.Add(mBase, uint32(v363))))
	v438 = v367
	goto L70
L80:
	;
	v366 = int64(*(*int8)(unsafe.Add(mBase, uint32(v363))))
	v438 = v366
	goto L70
L81:
	;
	v375 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v181+v371))))
	if v375&int32(1) == int32(0) {
		goto L84
	} else {
		goto L85
	}
L82:
	;
	goto L83
L83:
	;
	v421 = int32(1)
	v425 = (v371 + v351 - v421) & (int32(0) - v351)
	*(*int32)(unsafe.Add(mBase, uint32(v347))) = v425
	v427 = v181 + v425
	v428 = F_strlen(m, v427)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v347))) = v428 + v425 + v421
	v438 = base.I64_extend_i32_u(v427)
	goto L70
L84:
	;
	v385 = (v371 + v351 - int32(1)) & (int32(0) - v351)
	*(*int32)(unsafe.Add(mBase, uint32(v347))) = v385
	v388 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v181+v385))))
	v389 = v385
	v390 = v388
	goto L86
L85:
	;
	v389 = v371
	v390 = v375
	goto L86
L86:
	;
	v391 = v181 + v389
	if v390 == int32(1) {
		goto L88
	} else {
		goto L89
	}
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v347))) = v416 + v389
	v438 = base.I64_extend_i32_u(v391)
	goto L70
L88:
	;
	v395 = int32(18)
	v397 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v391)+1)))
	if v397 == v395 {
		goto L91
	} else {
		goto L92
	}
L89:
	;
	goto L90
L90:
	;
	v408 = int32(1)
	if v390&v408 != 0 {
		v416 = int32(base.Ui32(v390) >> (uint(v408) % 32))
		goto L87
	} else {
		goto L97
	}
L91:
	;
	v400 = v395
	goto L93
L92:
	;
	v400 = int32(2)
	goto L93
L93:
	;
	if base.Ui32((v397-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	v407 = int32(6)
	goto L96
L95:
	;
	v407 = v400
	goto L96
L96:
	;
	v416 = v407
	goto L87
L97:
	;
	v412 = *(*int32)(unsafe.Add(mBase, uint32(v391)))
	v416 = int32(base.Ui32(v412) >> (uint(int32(2)) % 32))
	goto L87
L98:
	;
	goto L69
L99:
	;
	v464 = v445
	goto L102
L100:
	;
	v589 = v445
	goto L101
L101:
	;
	v605 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
	if base.Ui32(v2) <= base.Ui32(v589) {
		v616 = v605
		goto L33
	} else {
		goto L136
	}
L102:
	;
	v484 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v464+v28))))
	if v484 != 0 {
		goto L104
	} else {
		goto L105
	}
L103:
	;
	v589 = v171
	goto L101
L104:
	;
	v582 = int64(0)
	goto L106
L105:
	;
	v487 = v21 + int32(12)
	v490 = v188 + v464<<(uint(int32(3))%32)
	v491 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v490)+4)))
	v492 = int32(*(*int16)(unsafe.Add(mBase, uint32(v490)+2)))
	v493 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v490)+5)))
	if int32(0) < v492 {
		goto L108
	} else {
		goto L109
	}
L106:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v189+v464<<(uint(int32(3))%32)))) = v582
	v585 = v464 + int32(1)
	if v585 != v171 {
		v464 = v585
		goto L102
	} else {
		goto L135
	}
L107:
	;
	v582 = v580
	goto L106
L108:
	;
	v496 = *(*int32)(unsafe.Add(mBase, uint32(v487)))
	v502 = (v493 + v496 - int32(1)) & (int32(0) - v493)
	*(*int32)(unsafe.Add(mBase, uint32(v487))) = v502 + v492
	v505 = v181 + v502
	if v491 != 0 {
		goto L111
	} else {
		goto L112
	}
L109:
	;
	goto L110
L110:
	;
	v513 = *(*int32)(unsafe.Add(mBase, uint32(v487)))
	if v492 == int32(-1) {
		goto L118
	} else {
		goto L119
	}
L111:
	;
	switch v492 - int32(1) {
	case 0:
		goto L117
	case 1:
		goto L116
	default:
		goto L114
	case 3:
		goto L115
	}
L112:
	;
	goto L113
L113:
	;
	v580 = base.I64_extend_i32_u(v505)
	goto L107
L114:
	;
	v511 = *(*int64)(unsafe.Add(mBase, uint32(v505)))
	v580 = v511
	goto L107
L115:
	;
	v510 = int64(*(*int32)(unsafe.Add(mBase, uint32(v505))))
	v580 = v510
	goto L107
L116:
	;
	v509 = int64(*(*int16)(unsafe.Add(mBase, uint32(v505))))
	v580 = v509
	goto L107
L117:
	;
	v508 = int64(*(*int8)(unsafe.Add(mBase, uint32(v505))))
	v580 = v508
	goto L107
L118:
	;
	v517 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v181+v513))))
	if v517&int32(1) == int32(0) {
		goto L121
	} else {
		goto L122
	}
L119:
	;
	goto L120
L120:
	;
	v563 = int32(1)
	v567 = (v513 + v493 - v563) & (int32(0) - v493)
	*(*int32)(unsafe.Add(mBase, uint32(v487))) = v567
	v569 = v181 + v567
	v570 = F_strlen(m, v569)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v487))) = v570 + v567 + v563
	v580 = base.I64_extend_i32_u(v569)
	goto L107
L121:
	;
	v527 = (v513 + v493 - int32(1)) & (int32(0) - v493)
	*(*int32)(unsafe.Add(mBase, uint32(v487))) = v527
	v530 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v181+v527))))
	v531 = v527
	v532 = v530
	goto L123
L122:
	;
	v531 = v513
	v532 = v517
	goto L123
L123:
	;
	v533 = v181 + v531
	if v532 == int32(1) {
		goto L125
	} else {
		goto L126
	}
L124:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v487))) = v558 + v531
	v580 = base.I64_extend_i32_u(v533)
	goto L107
L125:
	;
	v537 = int32(18)
	v539 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v533)+1)))
	if v539 == v537 {
		goto L128
	} else {
		goto L129
	}
L126:
	;
	goto L127
L127:
	;
	v550 = int32(1)
	if v532&v550 != 0 {
		v558 = int32(base.Ui32(v532) >> (uint(v550) % 32))
		goto L124
	} else {
		goto L134
	}
L128:
	;
	v542 = v537
	goto L130
L129:
	;
	v542 = int32(2)
	goto L130
L130:
	;
	if base.Ui32((v539-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L131
	} else {
		goto L132
	}
L131:
	;
	v549 = int32(6)
	goto L133
L132:
	;
	v549 = v542
	goto L133
L133:
	;
	v558 = v549
	goto L124
L134:
	;
	v554 = *(*int32)(unsafe.Add(mBase, uint32(v533)))
	v558 = int32(base.Ui32(v554) >> (uint(int32(2)) % 32))
	goto L124
L135:
	;
	goto L103
L136:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = v605
	v608 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v30)+18)))
	F_slot_getmissingattrs(m, l0, v608&int32(2047), v2)
	mBase = m.M
	v612 = m.ExcPending
	if v612 != 0 {
		goto L137
	} else {
		goto L138
	}
L137:
	;
	return
L138:
	;
	goto L32
}
func F_tts_heap_getsysattr(m *base.Module, l0 int32, l1 int32, l2 int32) int64 {
	var v6 int64
	_ = v6
	var v9 int32
	_ = v9
	v6 = Fn14394(m, l0, l1, l2, int32(_a_F_tts_heap_getsysattr_0), int32(369))
	v9 = m.ExcPending
	if v9 != 0 {
		return int64(0)
	} else {
		return v6
	}
}
func F_tts_minimal_copy_heap_tuple(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v67 int64
	_ = v67
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	if v5 != 0 {
		v41 = v5
		v44 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
		v47 = F_palloc(m, v44+int32(32))
		mBase = m.M
		v48 = m.ExcPending
		if v48 != 0 {
			return int32(0)
		} else {
			v49 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v47)+12)) = v49
			*(*uint16)(unsafe.Add(mBase, uint32(v47)+8)) = uint16(v49)
			*(*int32)(unsafe.Add(mBase, uint32(v47)+4)) = int32(-1)
			*(*int32)(unsafe.Add(mBase, uint32(v47))) = v44 + int32(8)
			*(*int32)(unsafe.Add(mBase, uint32(v47)+16)) = v47 + int32(24)
			v61 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
			if v61 != 0 {
				base.MemoryCopy(m, v47+int32(32), v41, v61)
			} else {
			}
			v65 = int32(0)
			*(*uint16)(unsafe.Add(mBase, uint32(v47)+40)) = uint16(v65)
			v67 = int64(0)
			*(*int64)(unsafe.Add(mBase, uint32(v47)+32)) = v67
			*(*int64)(unsafe.Add(mBase, uint32(v47)+24)) = v67
			return v47
		}
	} else {
		v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)))
		if v7&int32(4) != 0 {
			v41 = int32(0)
			v44 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
			v47 = F_palloc(m, v44+int32(32))
			mBase = m.M
			v48 = m.ExcPending
			if v48 != 0 {
				return int32(0)
			} else {
				v49 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v47)+12)) = v49
				*(*uint16)(unsafe.Add(mBase, uint32(v47)+8)) = uint16(v49)
				*(*int32)(unsafe.Add(mBase, uint32(v47)+4)) = int32(-1)
				*(*int32)(unsafe.Add(mBase, uint32(v47))) = v44 + int32(8)
				*(*int32)(unsafe.Add(mBase, uint32(v47)+16)) = v47 + int32(24)
				v61 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
				if v61 != 0 {
					base.MemoryCopy(m, v47+int32(32), v41, v61)
				} else {
				}
				v65 = int32(0)
				*(*uint16)(unsafe.Add(mBase, uint32(v47)+40)) = uint16(v65)
				v67 = int64(0)
				*(*int64)(unsafe.Add(mBase, uint32(v47)+32)) = v67
				*(*int64)(unsafe.Add(mBase, uint32(v47)+24)) = v67
				return v47
			}
		} else {
			v10 = int32(_a_F_tts_minimal_copy_heap_tuple_0)
			v11 = *(*int32)(unsafe.Add(mBase, _c_F_tts_minimal_copy_heap_tuple[0]))
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
			*(*int32)(unsafe.Add(mBase, _c_F_tts_minimal_copy_heap_tuple[0])) = v13
			v15 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+72)) = v15
			*(*uint16)(unsafe.Add(mBase, uint32(l0)+6)) = uint16(v15)
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			v23 = F_heap_form_minimal_tuple(m, v19, v20, v21, v15)
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = v23
				v28 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
				v30 = v28 | int32(4)
				*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)) = uint16(v30)
				v32 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
				v33 = int32(8)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+68)) = v23 - v33
				*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v32 + v33
				*(*int32)(unsafe.Add(mBase, _c_F_tts_minimal_copy_heap_tuple[0])) = v11
				v41 = v23
				v44 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
				v47 = F_palloc(m, v44+int32(32))
				mBase = m.M
				v48 = m.ExcPending
				if v48 != 0 {
					return int32(0)
				} else {
					v49 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v47)+12)) = v49
					*(*uint16)(unsafe.Add(mBase, uint32(v47)+8)) = uint16(v49)
					*(*int32)(unsafe.Add(mBase, uint32(v47)+4)) = int32(-1)
					*(*int32)(unsafe.Add(mBase, uint32(v47))) = v44 + int32(8)
					*(*int32)(unsafe.Add(mBase, uint32(v47)+16)) = v47 + int32(24)
					v61 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
					if v61 != 0 {
						base.MemoryCopy(m, v47+int32(32), v41, v61)
					} else {
					}
					v65 = int32(0)
					*(*uint16)(unsafe.Add(mBase, uint32(v47)+40)) = uint16(v65)
					v67 = int64(0)
					*(*int64)(unsafe.Add(mBase, uint32(v47)+32)) = v67
					*(*int64)(unsafe.Add(mBase, uint32(v47)+24)) = v67
					return v47
				}
			}
		}
	}
}
func F_tts_minimal_copyslot(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	v4 = int32(_a_F_tts_minimal_copyslot_0)
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_tts_minimal_copyslot[0]))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, _c_F_tts_minimal_copyslot[0])) = v7
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+48))
	v12 = m.T0[v11].(func(*base.Module, int32, int32) int32)(m, l1, int32(0))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, _c_F_tts_minimal_copyslot[0])) = v5
		v17 = F_ExecStoreMinimalTuple(m, v12, l0, int32(1))
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return
		} else {
			return
		}
	}
}
func F_tts_virtual_clear(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v22 int32
	_ = v22
	v3 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
	if v3&int32(4) != 0 {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
		F_pfree(m, v6)
		mBase = m.M
		v8 = m.ExcPending
		if v8 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = int32(0)
			v11 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
			v14 = v11 & int32(-5)
			v15 = int32(0)
			*(*uint16)(unsafe.Add(mBase, uint32(l0)+36)) = uint16(v15)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = int32(-1)
			*(*uint16)(unsafe.Add(mBase, uint32(l0)+6)) = uint16(v15)
			v22 = v14 | int32(2)
			*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)) = uint16(v22)
			return
		}
	} else {
		v14 = v3
		v15 = int32(0)
		*(*uint16)(unsafe.Add(mBase, uint32(l0)+36)) = uint16(v15)
		*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = int32(-1)
		*(*uint16)(unsafe.Add(mBase, uint32(l0)+6)) = uint16(v15)
		v22 = v14 | int32(2)
		*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)) = uint16(v22)
		return
	}
}
func F_tts_virtual_copyslot(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v51 int64
	_ = v51
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v7 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
	if v7&int32(4) != 0 {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
		F_pfree(m, v10)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = int32(0)
			v15 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
			v18 = v15 & int32(-5)
			v19 = int32(0)
			*(*uint16)(unsafe.Add(mBase, uint32(l0)+36)) = uint16(v19)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = int32(-1)
			*(*uint16)(unsafe.Add(mBase, uint32(l0)+6)) = uint16(v19)
			v26 = v18 | int32(2)
			*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)) = uint16(v26)
			v28 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
			v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
			v30 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+6)))
			if v30 < v29 {
				v32 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
				v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+16))
				m.T0[v33].(func(*base.Module, int32, int32))(m, l1, v29)
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return
				} else {
					v36 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
					if int32(0) < v36 {
						v42 = int32(0)
						for {
							v46 = v42 << (uint(int32(3)) % 32)
							v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
							v49 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
							v51 = *(*int64)(unsafe.Add(mBase, uint32(v49+v46)))
							*(*int64)(unsafe.Add(mBase, uint32(v46+v47))) = v51
							v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
							v55 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
							v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55+v42))))
							*(*uint8)(unsafe.Add(mBase, uint32(v53+v42))) = uint8(v57)
							v60 = v42 + int32(1)
							v61 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
							if v60 < v61 {
								v42 = v60
								continue
							} else {
								break
							}
							break
						}
						v66 = v61
					} else {
						v66 = v36
					}
					*(*uint16)(unsafe.Add(mBase, uint32(l0)+6)) = uint16(v66)
					v69 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
					v71 = v69 & int32(_a_F_tts_virtual_copyslot_0)
					*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)) = uint16(v71)
					F_tts_virtual_materialize(m, l0)
					mBase = m.M
					v74 = m.ExcPending
					if v74 != 0 {
						return
					} else {
						return
					}
				}
			} else {
				v36 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
				if int32(0) < v36 {
					v42 = int32(0)
					for {
						v46 = v42 << (uint(int32(3)) % 32)
						v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
						v49 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
						v51 = *(*int64)(unsafe.Add(mBase, uint32(v49+v46)))
						*(*int64)(unsafe.Add(mBase, uint32(v46+v47))) = v51
						v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						v55 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
						v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55+v42))))
						*(*uint8)(unsafe.Add(mBase, uint32(v53+v42))) = uint8(v57)
						v60 = v42 + int32(1)
						v61 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
						if v60 < v61 {
							v42 = v60
							continue
						} else {
							break
						}
						break
					}
					v66 = v61
				} else {
					v66 = v36
				}
				*(*uint16)(unsafe.Add(mBase, uint32(l0)+6)) = uint16(v66)
				v69 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
				v71 = v69 & int32(_a_F_tts_virtual_copyslot_0)
				*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)) = uint16(v71)
				F_tts_virtual_materialize(m, l0)
				mBase = m.M
				v74 = m.ExcPending
				if v74 != 0 {
					return
				} else {
					return
				}
			}
		}
	} else {
		v18 = v7
		v19 = int32(0)
		*(*uint16)(unsafe.Add(mBase, uint32(l0)+36)) = uint16(v19)
		*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = int32(-1)
		*(*uint16)(unsafe.Add(mBase, uint32(l0)+6)) = uint16(v19)
		v26 = v18 | int32(2)
		*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)) = uint16(v26)
		v28 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
		v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
		v30 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+6)))
		if v30 < v29 {
			v32 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
			v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+16))
			m.T0[v33].(func(*base.Module, int32, int32))(m, l1, v29)
			mBase = m.M
			v35 = m.ExcPending
			if v35 != 0 {
				return
			} else {
				v36 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
				if int32(0) < v36 {
					v42 = int32(0)
					for {
						v46 = v42 << (uint(int32(3)) % 32)
						v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
						v49 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
						v51 = *(*int64)(unsafe.Add(mBase, uint32(v49+v46)))
						*(*int64)(unsafe.Add(mBase, uint32(v46+v47))) = v51
						v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						v55 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
						v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55+v42))))
						*(*uint8)(unsafe.Add(mBase, uint32(v53+v42))) = uint8(v57)
						v60 = v42 + int32(1)
						v61 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
						if v60 < v61 {
							v42 = v60
							continue
						} else {
							break
						}
						break
					}
					v66 = v61
				} else {
					v66 = v36
				}
				*(*uint16)(unsafe.Add(mBase, uint32(l0)+6)) = uint16(v66)
				v69 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
				v71 = v69 & int32(_a_F_tts_virtual_copyslot_0)
				*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)) = uint16(v71)
				F_tts_virtual_materialize(m, l0)
				mBase = m.M
				v74 = m.ExcPending
				if v74 != 0 {
					return
				} else {
					return
				}
			}
		} else {
			v36 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
			if int32(0) < v36 {
				v42 = int32(0)
				for {
					v46 = v42 << (uint(int32(3)) % 32)
					v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
					v49 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
					v51 = *(*int64)(unsafe.Add(mBase, uint32(v49+v46)))
					*(*int64)(unsafe.Add(mBase, uint32(v46+v47))) = v51
					v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					v55 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
					v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55+v42))))
					*(*uint8)(unsafe.Add(mBase, uint32(v53+v42))) = uint8(v57)
					v60 = v42 + int32(1)
					v61 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
					if v60 < v61 {
						v42 = v60
						continue
					} else {
						break
					}
					break
				}
				v66 = v61
			} else {
				v66 = v36
			}
			*(*uint16)(unsafe.Add(mBase, uint32(l0)+6)) = uint16(v66)
			v69 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
			v71 = v69 & int32(_a_F_tts_virtual_copyslot_0)
			*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)) = uint16(v71)
			F_tts_virtual_materialize(m, l0)
			mBase = m.M
			v74 = m.ExcPending
			if v74 != 0 {
				return
			} else {
				return
			}
		}
	}
}
func F_tts_virtual_materialize(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
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
	var v39 int64
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v155 int64
	_ = v155
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
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
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v245 int32
	_ = v245
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	v2 = int32(0)
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)))
	if v11&int32(4) != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	if v15 <= int32(0) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v22 = v2
	v23 = v15
	v25 = v2
	goto L4
L4:
	;
	v31 = v25 << (uint(int32(3)) % 32)
	v32 = v14 + int32(28) + v31
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+4)))
	if v33 != 0 {
		v112 = v22
		v113 = v23
		goto L6
	} else {
		goto L7
	}
L5:
	;
	if v112 == int32(0) {
		goto L1
	} else {
		goto L31
	}
L6:
	;
	v118 = v25 + int32(1)
	if v118 < v113 {
		v22 = v112
		v23 = v113
		v25 = v118
		goto L4
	} else {
		goto L30
	}
L7:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34+v25))))
	if v36 != 0 {
		v112 = v22
		v113 = v23
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v39 = *(*int64)(unsafe.Add(mBase, uint32(v37+v31)))
	v40 = int32(*(*int16)(unsafe.Add(mBase, uint32(v32)+2)))
	if v40 == int32(-1) {
		goto L12
	} else {
		goto L13
	}
L9:
	;
	v112 = v106 + v107
	v113 = v23
	goto L6
L10:
	;
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+5)))
	v94 = int32(1)
	v98 = (v22 + v92 - v94) & (int32(0) - v92)
	if v44&v94 != 0 {
		goto L27
	} else {
		goto L28
	}
L11:
	;
	v79 = int32(18)
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+1)))
	if v81 == v79 {
		goto L21
	} else {
		goto L22
	}
L12:
	;
	v43 = base.I32_wrap_i64(v39)
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43))))
	if v44 != int32(1) {
		goto L10
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+5)))
	v69 = int32(0)
	v71 = (v22 + v65 - int32(1)) & (v69 - v65)
	if v69 < v40 {
		v106 = v40
		v107 = v71
		goto L9
	} else {
		goto L20
	}
L15:
	;
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+5)))
	v53 = (v22 + v47 - int32(1)) & (int32(0) - v47)
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+1)))
	if v54&int32(254) != int32(2) {
		goto L11
	} else {
		goto L16
	}
L16:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v39))+2))
	goto L17
L17:
	;
	v61 = F_EOH_get_flat_size(m, v60)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	return
L19:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v112 = v61 + v53
	v113 = v64
	goto L6
L20:
	;
	v75 = F_strlen(m, base.I32_wrap_i64(v39))
	mBase = m.M
	v106 = v75 + int32(1)
	v107 = v71
	goto L9
L21:
	;
	v84 = v79
	goto L23
L22:
	;
	v84 = int32(2)
	goto L23
L23:
	;
	if base.Ui32((v81-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v91 = int32(6)
	goto L26
L25:
	;
	v91 = v84
	goto L26
L26:
	;
	v106 = v91
	v107 = v53
	goto L9
L27:
	;
	v106 = int32(base.Ui32(v44) >> (uint(int32(1)) % 32))
	v107 = v98
	goto L9
L28:
	;
	goto L29
L29:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
	v106 = int32(base.Ui32(v103) >> (uint(int32(2)) % 32))
	v107 = v98
	goto L9
L30:
	;
	goto L5
L31:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v123 = F_MemoryContextAlloc(m, v122, v112)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L18
	} else {
		goto L32
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v123
	v126 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
	v128 = v126 | int32(4)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)) = uint16(v128)
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	if v130 <= int32(0) {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	v138 = v123
	v141 = int32(0)
	goto L34
L34:
	;
	v147 = v141 << (uint(int32(3)) % 32)
	v148 = v14 + int32(28) + v147
	v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148)+4)))
	if v149 != 0 {
		v245 = v138
		goto L36
	} else {
		goto L37
	}
L35:
	;
	goto L1
L36:
	;
	v250 = v141 + int32(1)
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	if v250 < v251 {
		v138 = v245
		v141 = v250
		goto L34
	} else {
		goto L64
	}
L37:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v150+v141))))
	if v152 != 0 {
		v245 = v138
		goto L36
	} else {
		goto L38
	}
L38:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v155 = *(*int64)(unsafe.Add(mBase, uint32(v153+v147)))
	v156 = int32(*(*int16)(unsafe.Add(mBase, uint32(v148)+2)))
	if v156 == int32(-1) {
		goto L43
	} else {
		goto L44
	}
L39:
	;
	if v236 != 0 {
		goto L61
	} else {
		goto L62
	}
L40:
	;
	v231 = base.I32_wrap_i64(v155)
	v232 = F_strlen(m, v231)
	mBase = m.M
	v235 = v192
	v236 = v232 + int32(1)
	v237 = v231
	goto L39
L41:
	;
	v217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148)+5)))
	v219 = int32(1)
	v223 = (v138 + v217 - v219) & (int32(0) - v217)
	if v160&v219 != 0 {
		goto L58
	} else {
		goto L59
	}
L42:
	;
	v197 = int32(18)
	v199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v159)+1)))
	if v199 == v197 {
		goto L52
	} else {
		goto L53
	}
L43:
	;
	v159 = base.I32_wrap_i64(v155)
	v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v159))))
	if v160 != int32(1) {
		goto L41
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148)+5)))
	v190 = int32(0)
	v192 = (v138 + v186 - int32(1)) & (v190 - v186)
	if v156 <= v190 {
		goto L40
	} else {
		goto L51
	}
L46:
	;
	v163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v159)+1)))
	if v163&int32(254) != int32(2) {
		goto L42
	} else {
		goto L47
	}
L47:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v155))+2))
	goto L48
L48:
	;
	v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148)+5)))
	v176 = (v138 + v170 - int32(1)) & (int32(0) - v170)
	v177 = F_EOH_get_flat_size(m, v169)
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L18
	} else {
		goto L49
	}
L49:
	;
	F_EOH_flatten_into(m, v169, v176, v177)
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L18
	} else {
		goto L50
	}
L50:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v181+v147))) = base.I64_extend_i32_u(v176)
	v245 = v177 + v176
	goto L36
L51:
	;
	v235 = v192
	v236 = v156
	v237 = base.I32_wrap_i64(v155)
	goto L39
L52:
	;
	v202 = v197
	goto L54
L53:
	;
	v202 = int32(2)
	goto L54
L54:
	;
	if base.Ui32((v199-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v209 = int32(6)
	goto L57
L56:
	;
	v209 = v202
	goto L57
L57:
	;
	v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148)+5)))
	v235 = (v138 + v210 - int32(1)) & (int32(0) - v210)
	v236 = v209
	v237 = v159
	goto L39
L58:
	;
	v235 = v223
	v236 = int32(base.Ui32(v160) >> (uint(int32(1)) % 32))
	v237 = v159
	goto L39
L59:
	;
	goto L60
L60:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v159)))
	v235 = v223
	v236 = int32(base.Ui32(v228) >> (uint(int32(2)) % 32))
	v237 = v159
	goto L39
L61:
	;
	base.MemoryCopy(m, v235, v237, v236)
	goto L63
L62:
	;
	goto L63
L63:
	;
	v239 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v239+v147))) = base.I64_extend_i32_u(v235)
	v245 = v235 + v236
	goto L36
L64:
	;
	goto L35
}
