package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_tts_buffer_heap_getsysattr(m *base.Module, l0 int32, l1 int32, l2 int32) int64 {
	var v6 int64
	_ = v6
	var v9 int32
	_ = v9
	v6 = Fn14394(m, l0, l1, l2, int32(_a_F_tts_buffer_heap_getsysattr_0), int32(774))
	v9 = m.ExcPending
	if v9 != 0 {
		return int64(0)
	} else {
		return v6
	}
}
func F_tts_minimal_getsomeattrs(m *base.Module, l0 int32, l1 int32) {
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
	v138 = int32(_a_F_tts_minimal_getsomeattrs_0)
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
	*(*int32)(unsafe.Add(mBase, uint32(l0)+72)) = v616
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
	v233 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
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
	switch v282&int32(_a_F_tts_minimal_getsomeattrs_1) - int32(1) {
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
	*(*int32)(unsafe.Add(mBase, uint32(l0)+72)) = v605
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
func F_tts_minimal_getsysattr(m *base.Module, l0 int32, l1 int32, l2 int32) int64 {
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	F_errstart_cold(m, int32(21), int32(0))
	v9 = m.ExcPending
	if v9 != 0 {
		return int64(0)
	} else {
		F_errcode(m, int32(1088))
		v12 = m.ExcPending
		if v12 != 0 {
			return int64(0)
		} else {
			F_errmsg(m, int32(_a_F_tts_minimal_getsysattr_0), int32(0))
			v16 = m.ExcPending
			if v16 != 0 {
				return int64(0)
			} else {
				F_errfinish(m, int32(_a_F_tts_minimal_getsysattr_1), int32(564), int32(_a_F_tts_minimal_getsysattr_2))
				v21 = m.ExcPending
				if v21 != 0 {
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
func F_tts_minimal_init(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = l0 + int32(52)
	return
}
func F_tts_virtual_copy_minimal_tuple(m *base.Module, l0 int32, l1 int32) int32 {
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
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = F_heap_form_minimal_tuple(m, v3, v4, v5, l1)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return v6
	}
}
