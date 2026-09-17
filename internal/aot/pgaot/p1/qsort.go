package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_qsortCompareItemPointers(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int64
	_ = v5
	var v6 int64
	_ = v6
	var v7 int64
	_ = v7
	var v9 int64
	_ = v9
	var v10 int64
	_ = v10
	var v13 int64
	_ = v13
	var v14 int64
	_ = v14
	var v15 int64
	_ = v15
	var v18 int64
	_ = v18
	var v22 int64
	_ = v22
	v5 = int64(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
	v6 = int64(*(*uint16)(unsafe.Add(mBase, uint32(l0)+2)))
	v7 = int64(32)
	v9 = int64(*(*uint16)(unsafe.Add(mBase, uint32(l0))))
	v10 = int64(48)
	v13 = v5 | (v6<<(uint(v7)%64) | v9<<(uint(v10)%64))
	v14 = int64(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
	v15 = int64(*(*uint16)(unsafe.Add(mBase, uint32(l1)+2)))
	v18 = int64(*(*uint16)(unsafe.Add(mBase, uint32(l1))))
	v22 = v14 | (v15<<(uint(v7)%64) | v18<<(uint(v10)%64))
	return base.B2i32(base.Ui64(v22) < base.Ui64(v13)) - base.B2i32(base.Ui64(v13) < base.Ui64(v22))
}
func F_qsort_ssup(m *base.Module, l0 int32, l1 int32, l2 int32) {
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
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v71 int32
	_ = v71
	var v78 int32
	_ = v78
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v117 int32
	_ = v117
	var v121 int64
	_ = v121
	var v123 int64
	_ = v123
	var v126 int32
	_ = v126
	var v127 int64
	_ = v127
	var v129 int64
	_ = v129
	var v131 int64
	_ = v131
	var v133 int64
	_ = v133
	var v151 int32
	_ = v151
	var v157 int32
	_ = v157
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v179 int32
	_ = v179
	var v184 int32
	_ = v184
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v201 int32
	_ = v201
	var v206 int32
	_ = v206
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v250 int64
	_ = v250
	var v252 int64
	_ = v252
	var v254 int64
	_ = v254
	var v256 int64
	_ = v256
	var v258 int64
	_ = v258
	var v260 int64
	_ = v260
	var v263 int32
	_ = v263
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v299 int32
	_ = v299
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v317 int32
	_ = v317
	var v321 int64
	_ = v321
	var v323 int64
	_ = v323
	var v325 int64
	_ = v325
	var v327 int64
	_ = v327
	var v329 int64
	_ = v329
	var v331 int64
	_ = v331
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v342 int32
	_ = v342
	var v347 int32
	_ = v347
	var v350 int32
	_ = v350
	var v363 int32
	_ = v363
	var v366 int32
	_ = v366
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v379 int32
	_ = v379
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v395 int32
	_ = v395
	var v399 int64
	_ = v399
	var v401 int64
	_ = v401
	var v403 int64
	_ = v403
	var v405 int64
	_ = v405
	var v407 int64
	_ = v407
	var v409 int64
	_ = v409
	var v414 int32
	_ = v414
	var v416 int32
	_ = v416
	var v418 int32
	_ = v418
	var v420 int32
	_ = v420
	var v426 int32
	_ = v426
	var v429 int32
	_ = v429
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v441 int32
	_ = v441
	var v443 int32
	_ = v443
	var v454 int32
	_ = v454
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v465 int64
	_ = v465
	var v467 int64
	_ = v467
	var v469 int32
	_ = v469
	var v470 int64
	_ = v470
	var v472 int64
	_ = v472
	var v474 int64
	_ = v474
	var v476 int64
	_ = v476
	var v479 int32
	_ = v479
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v502 int32
	_ = v502
	var v504 int32
	_ = v504
	var v515 int32
	_ = v515
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v526 int64
	_ = v526
	var v528 int64
	_ = v528
	var v530 int32
	_ = v530
	var v531 int64
	_ = v531
	var v533 int64
	_ = v533
	var v535 int64
	_ = v535
	var v537 int64
	_ = v537
	var v540 int32
	_ = v540
	var v560 int32
	_ = v560
	var v572 int32
	_ = v572
	var v576 int64
	_ = v576
	var v578 int64
	_ = v578
	var v580 int64
	_ = v580
	var v582 int64
	_ = v582
	var v584 int64
	_ = v584
	var v586 int64
	_ = v586
	var v588 int32
	_ = v588
	v15 = m.G0
	v17 = v15 - int32(16)
	m.G0 = v17
	v19 = l0
	v20 = l1
	goto L1
L1:
	;
	v34 = v19 + int32(16)
	v36 = v20
	goto L3
L2:
	;
	m.G0 = v17 + int32(16)
	return
L3:
	;
	v50 = *(*int32)(unsafe.Add(mBase, _c_F_qsort_ssup[0]))
	if v50 != 0 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	goto L2
L5:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	v55 = v19 + v36<<(uint(int32(4))%32)
	if base.Ui32(v36) <= base.Ui32(int32(6)) {
		goto L11
	} else {
		goto L12
	}
L8:
	;
	return
L9:
	;
	goto L7
L10:
	;
	goto L4
L11:
	;
	if base.Ui32(v36) < base.Ui32(int32(2)) {
		goto L10
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	v157 = v34
	goto L39
L14:
	;
	v71 = v34
	goto L15
L15:
	;
	if base.Ui32(v71) <= base.Ui32(v19) {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	goto L10
L17:
	;
	v151 = v71 + int32(16)
	if base.Ui32(v151) < base.Ui32(v55) {
		v71 = v151
		goto L15
	} else {
		goto L38
	}
L18:
	;
	v78 = v71
	goto L19
L19:
	;
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78)+8)))
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78-int32(8)))))
	if v92 == int32(1) {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	goto L17
L21:
	;
	v121 = *(*int64)(unsafe.Add(mBase, uint32(v78)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+8)) = v121
	v123 = *(*int64)(unsafe.Add(mBase, uint32(v78)))
	*(*int64)(unsafe.Add(mBase, uint32(v17))) = v123
	v126 = v78 - int32(16)
	v127 = *(*int64)(unsafe.Add(mBase, uint32(v126)))
	*(*int64)(unsafe.Add(mBase, uint32(v78))) = v127
	v129 = *(*int64)(unsafe.Add(mBase, uint32(v126)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v78)+8)) = v129
	v131 = *(*int64)(unsafe.Add(mBase, uint32(v17)))
	*(*int64)(unsafe.Add(mBase, uint32(v126))) = v131
	v133 = *(*int64)(unsafe.Add(mBase, uint32(v17)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v126)+8)) = v133
	if base.Ui32(v19) < base.Ui32(v126) {
		v78 = v126
		goto L19
	} else {
		goto L37
	}
L22:
	;
	if v89&int32(1) != 0 {
		goto L17
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	if v89&int32(1) != 0 {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+9)))
	if v97 == int32(0) {
		goto L21
	} else {
		goto L26
	}
L26:
	;
	goto L17
L27:
	;
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+9)))
	if v102 != 0 {
		goto L21
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v78-int32(12))))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v78)+4))
	v107 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v108 = m.T0[v107].(func(*base.Module, int32, int32, int32) int32)(m, v105, v106, l2)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L8
	} else {
		goto L31
	}
L30:
	;
	goto L17
L31:
	;
	v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+8)))
	if v110 == int32(1) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	if v108 < int32(0) {
		goto L21
	} else {
		goto L35
	}
L33:
	;
	v117 = v108
	goto L34
L34:
	;
	if v117 <= int32(0) {
		goto L17
	} else {
		goto L36
	}
L35:
	;
	v117 = int32(0) - v108
	goto L34
L36:
	;
	goto L21
L37:
	;
	goto L20
L38:
	;
	goto L16
L39:
	;
	v168 = *(*int32)(unsafe.Add(mBase, _c_F_qsort_ssup[0]))
	if v168 != 0 {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	v213 = v19 + v36<<(uint(int32(3))%32)&int32(-16)
	if v36 != int32(7) {
		goto L63
	} else {
		goto L64
	}
L41:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L8
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	v171 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157)+8)))
	v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157-int32(8)))))
	if v174 == int32(1) {
		goto L47
	} else {
		goto L48
	}
L44:
	;
	goto L43
L45:
	;
	goto L40
L46:
	;
	v206 = v157 + int32(16)
	if base.Ui32(v206) < base.Ui32(v55) {
		v157 = v206
		goto L39
	} else {
		goto L62
	}
L47:
	;
	if v171&int32(1) != 0 {
		goto L46
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	if v171&int32(1) != 0 {
		goto L52
	} else {
		goto L53
	}
L50:
	;
	v179 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+9)))
	if v179 == int32(0) {
		goto L45
	} else {
		goto L51
	}
L51:
	;
	goto L46
L52:
	;
	v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+9)))
	if v184 == int32(0) {
		goto L46
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v157-int32(12))))
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v157)+4))
	v191 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v192 = m.T0[v191].(func(*base.Module, int32, int32, int32) int32)(m, v189, v190, l2)
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L8
	} else {
		goto L56
	}
L55:
	;
	goto L45
L56:
	;
	v194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+8)))
	if v194 == int32(1) {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	if v192 < int32(0) {
		goto L45
	} else {
		goto L60
	}
L58:
	;
	v201 = v192
	goto L59
L59:
	;
	if int32(0) < v201 {
		goto L45
	} else {
		goto L61
	}
L60:
	;
	v201 = int32(0) - v192
	goto L59
L61:
	;
	goto L46
L62:
	;
	goto L10
L63:
	;
	v217 = v55 - int32(16)
	if base.Ui32(v36) < base.Ui32(int32(41)) {
		goto L67
	} else {
		goto L68
	}
L64:
	;
	v247 = v213
	goto L65
L65:
	;
	v250 = *(*int64)(unsafe.Add(mBase, uint32(v19)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+8)) = v250
	v252 = *(*int64)(unsafe.Add(mBase, uint32(v19)))
	*(*int64)(unsafe.Add(mBase, uint32(v17))) = v252
	v254 = *(*int64)(unsafe.Add(mBase, uint32(v247)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v19)+8)) = v254
	v256 = *(*int64)(unsafe.Add(mBase, uint32(v247)))
	*(*int64)(unsafe.Add(mBase, uint32(v19))) = v256
	v258 = *(*int64)(unsafe.Add(mBase, uint32(v17)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v247)+8)) = v258
	v260 = *(*int64)(unsafe.Add(mBase, uint32(v17)))
	*(*int64)(unsafe.Add(mBase, uint32(v247))) = v260
	v263 = v55 - int32(16)
	v267 = v34
	v268 = v263
	v270 = v34
	v271 = v263
	goto L74
L66:
	;
	v243 = F_qsort_ssup_med3(m, v238, v240, v239, l2)
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L8
	} else {
		goto L73
	}
L67:
	;
	v238 = v19
	v239 = v217
	v240 = v213
	goto L66
L68:
	;
	goto L69
L69:
	;
	v221 = int32(base.Ui32(v36) >> (uint(int32(3)) % 32))
	v223 = v221 << (uint(int32(4)) % 32)
	v226 = v221 << (uint(int32(5)) % 32)
	v228 = F_qsort_ssup_med3(m, v19, v19+v223, v19+v226, l2)
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L8
	} else {
		goto L70
	}
L70:
	;
	v232 = F_qsort_ssup_med3(m, v213-v223, v213, v213+v223, l2)
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L8
	} else {
		goto L71
	}
L71:
	;
	v236 = F_qsort_ssup_med3(m, v217-v226, v217-v223, v217, l2)
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L8
	} else {
		goto L72
	}
L72:
	;
	v238 = v228
	v239 = v236
	v240 = v232
	goto L66
L73:
	;
	v247 = v243
	goto L65
L74:
	;
	if base.Ui32(v268) < base.Ui32(v267) {
		v347 = v267
		v350 = v270
		goto L76
	} else {
		goto L77
	}
L76:
	;
	if base.Ui32(v347) <= base.Ui32(v268) {
		goto L104
	} else {
		goto L105
	}
L77:
	;
	v282 = v267
	v285 = v270
	goto L78
L78:
	;
	v293 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+8)))
	v294 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v282)+8)))
	if v294 == int32(1) {
		goto L82
	} else {
		goto L83
	}
L79:
	;
	v347 = v342
	v350 = v336
	goto L76
L80:
	;
	v338 = *(*int32)(unsafe.Add(mBase, _c_F_qsort_ssup[0]))
	if v338 != 0 {
		goto L98
	} else {
		goto L99
	}
L81:
	;
	v321 = *(*int64)(unsafe.Add(mBase, uint32(v285)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+8)) = v321
	v323 = *(*int64)(unsafe.Add(mBase, uint32(v285)))
	*(*int64)(unsafe.Add(mBase, uint32(v17))) = v323
	v325 = *(*int64)(unsafe.Add(mBase, uint32(v282)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v285)+8)) = v325
	v327 = *(*int64)(unsafe.Add(mBase, uint32(v282)))
	*(*int64)(unsafe.Add(mBase, uint32(v285))) = v327
	v329 = *(*int64)(unsafe.Add(mBase, uint32(v17)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v282)+8)) = v329
	v331 = *(*int64)(unsafe.Add(mBase, uint32(v17)))
	*(*int64)(unsafe.Add(mBase, uint32(v282))) = v331
	v336 = v285 + int32(16)
	goto L80
L82:
	;
	if v293&int32(1) != 0 {
		goto L81
	} else {
		goto L85
	}
L83:
	;
	goto L84
L84:
	;
	if v293&int32(1) != 0 {
		goto L87
	} else {
		goto L88
	}
L85:
	;
	v299 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+9)))
	if v299 != 0 {
		v336 = v285
		goto L80
	} else {
		goto L86
	}
L86:
	;
	v347 = v282
	v350 = v285
	goto L76
L87:
	;
	v302 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+9)))
	if v302 == int32(0) {
		v336 = v285
		goto L80
	} else {
		goto L90
	}
L88:
	;
	goto L89
L89:
	;
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v282)+4))
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	v307 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v308 = m.T0[v307].(func(*base.Module, int32, int32, int32) int32)(m, v305, v306, l2)
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L8
	} else {
		goto L91
	}
L90:
	;
	v347 = v282
	v350 = v285
	goto L76
L91:
	;
	v310 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+8)))
	if v310 == int32(1) {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	if v308 < int32(0) {
		v347 = v282
		v350 = v285
		goto L76
	} else {
		goto L95
	}
L93:
	;
	v317 = v308
	goto L94
L94:
	;
	if int32(0) < v317 {
		v347 = v282
		v350 = v285
		goto L76
	} else {
		goto L96
	}
L95:
	;
	v317 = int32(0) - v308
	goto L94
L96:
	;
	if v317 != 0 {
		v336 = v285
		goto L80
	} else {
		goto L97
	}
L97:
	;
	goto L81
L98:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L8
	} else {
		goto L101
	}
L99:
	;
	goto L100
L100:
	;
	v342 = v282 + int32(16)
	if base.Ui32(v342) <= base.Ui32(v268) {
		v282 = v342
		v285 = v336
		goto L78
	} else {
		goto L102
	}
L101:
	;
	goto L100
L102:
	;
	goto L79
L103:
	;
	v576 = *(*int64)(unsafe.Add(mBase, uint32(v347)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+8)) = v576
	v578 = *(*int64)(unsafe.Add(mBase, uint32(v347)))
	*(*int64)(unsafe.Add(mBase, uint32(v17))) = v578
	v580 = *(*int64)(unsafe.Add(mBase, uint32(v363)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v347)+8)) = v580
	v582 = *(*int64)(unsafe.Add(mBase, uint32(v363)))
	*(*int64)(unsafe.Add(mBase, uint32(v347))) = v582
	v584 = *(*int64)(unsafe.Add(mBase, uint32(v17)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v363)+8)) = v584
	v586 = *(*int64)(unsafe.Add(mBase, uint32(v17)))
	*(*int64)(unsafe.Add(mBase, uint32(v363))) = v586
	v588 = int32(16)
	v267 = v347 + v588
	v268 = v363 - v588
	v270 = v350
	v271 = v366
	goto L74
L104:
	;
	v363 = v268
	v366 = v271
	goto L107
L105:
	;
	v426 = v268
	v429 = v271
	goto L106
L106:
	;
	v437 = int32(4)
	v438 = (v350 - v19) >> (uint(v437) % 32)
	v441 = (v347 - v350) >> (uint(v437) % 32)
	if v438 < v441 {
		goto L132
	} else {
		goto L133
	}
L107:
	;
	v373 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+8)))
	v374 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v363)+8)))
	if v374 == int32(1) {
		goto L111
	} else {
		goto L112
	}
L108:
	;
	v426 = v420
	v429 = v414
	goto L106
L109:
	;
	v416 = *(*int32)(unsafe.Add(mBase, _c_F_qsort_ssup[0]))
	if v416 != 0 {
		goto L127
	} else {
		goto L128
	}
L110:
	;
	v399 = *(*int64)(unsafe.Add(mBase, uint32(v363)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+8)) = v399
	v401 = *(*int64)(unsafe.Add(mBase, uint32(v363)))
	*(*int64)(unsafe.Add(mBase, uint32(v17))) = v401
	v403 = *(*int64)(unsafe.Add(mBase, uint32(v366)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v363)+8)) = v403
	v405 = *(*int64)(unsafe.Add(mBase, uint32(v366)))
	*(*int64)(unsafe.Add(mBase, uint32(v363))) = v405
	v407 = *(*int64)(unsafe.Add(mBase, uint32(v17)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v366)+8)) = v407
	v409 = *(*int64)(unsafe.Add(mBase, uint32(v17)))
	*(*int64)(unsafe.Add(mBase, uint32(v366))) = v409
	v414 = v366 - int32(16)
	goto L109
L111:
	;
	if v373&int32(1) != 0 {
		goto L110
	} else {
		goto L114
	}
L112:
	;
	goto L113
L113:
	;
	if v373&int32(1) != 0 {
		goto L116
	} else {
		goto L117
	}
L114:
	;
	v379 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+9)))
	if v379 != 0 {
		goto L103
	} else {
		goto L115
	}
L115:
	;
	v414 = v366
	goto L109
L116:
	;
	v382 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+9)))
	if v382 != 0 {
		v414 = v366
		goto L109
	} else {
		goto L119
	}
L117:
	;
	goto L118
L118:
	;
	v383 = *(*int32)(unsafe.Add(mBase, uint32(v363)+4))
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	v385 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v386 = m.T0[v385].(func(*base.Module, int32, int32, int32) int32)(m, v383, v384, l2)
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L8
	} else {
		goto L120
	}
L119:
	;
	goto L103
L120:
	;
	v388 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+8)))
	if v388 == int32(1) {
		goto L121
	} else {
		goto L122
	}
L121:
	;
	if v386 < int32(0) {
		v414 = v366
		goto L109
	} else {
		goto L124
	}
L122:
	;
	v395 = v386
	goto L123
L123:
	;
	if v395 < int32(0) {
		goto L103
	} else {
		goto L125
	}
L124:
	;
	v395 = int32(0) - v386
	goto L123
L125:
	;
	if v395 != 0 {
		v414 = v366
		goto L109
	} else {
		goto L126
	}
L126:
	;
	goto L110
L127:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v418 = m.ExcPending
	if v418 != 0 {
		goto L8
	} else {
		goto L130
	}
L128:
	;
	goto L129
L129:
	;
	v420 = v363 - int32(16)
	if base.Ui32(v347) <= base.Ui32(v420) {
		v363 = v420
		v366 = v414
		goto L107
	} else {
		goto L131
	}
L130:
	;
	goto L129
L131:
	;
	goto L108
L132:
	;
	v443 = v438
	goto L134
L133:
	;
	v443 = v441
	goto L134
L134:
	;
	if v443 != 0 {
		goto L135
	} else {
		goto L136
	}
L135:
	;
	v454 = int32(0)
	goto L138
L136:
	;
	goto L137
L137:
	;
	v496 = int32(4)
	v497 = (v429 - v426) >> (uint(v496) % 32)
	v502 = (v55-v429)>>(uint(v496)%32) - int32(1)
	if v497 < v502 {
		goto L141
	} else {
		goto L142
	}
L138:
	;
	v463 = v454 << (uint(int32(4)) % 32)
	v464 = v19 + v463
	v465 = *(*int64)(unsafe.Add(mBase, uint32(v464)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+8)) = v465
	v467 = *(*int64)(unsafe.Add(mBase, uint32(v464)))
	*(*int64)(unsafe.Add(mBase, uint32(v17))) = v467
	v469 = v463 + (v347 - v443<<(uint(int32(4))%32))
	v470 = *(*int64)(unsafe.Add(mBase, uint32(v469)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v464)+8)) = v470
	v472 = *(*int64)(unsafe.Add(mBase, uint32(v469)))
	*(*int64)(unsafe.Add(mBase, uint32(v464))) = v472
	v474 = *(*int64)(unsafe.Add(mBase, uint32(v17)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v469)+8)) = v474
	v476 = *(*int64)(unsafe.Add(mBase, uint32(v17)))
	*(*int64)(unsafe.Add(mBase, uint32(v469))) = v476
	v479 = v454 + int32(1)
	if v479 != v443 {
		v454 = v479
		goto L138
	} else {
		goto L140
	}
L139:
	;
	goto L137
L140:
	;
	goto L139
L141:
	;
	v504 = v497
	goto L143
L142:
	;
	v504 = v502
	goto L143
L143:
	;
	if v504 != 0 {
		goto L144
	} else {
		goto L145
	}
L144:
	;
	v515 = int32(0)
	goto L147
L145:
	;
	goto L146
L146:
	;
	if base.Ui32(v441) <= base.Ui32(v497) {
		goto L150
	} else {
		goto L151
	}
L147:
	;
	v524 = v515 << (uint(int32(4)) % 32)
	v525 = v347 + v524
	v526 = *(*int64)(unsafe.Add(mBase, uint32(v525)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+8)) = v526
	v528 = *(*int64)(unsafe.Add(mBase, uint32(v525)))
	*(*int64)(unsafe.Add(mBase, uint32(v17))) = v528
	v530 = v524 + (v55 - v504<<(uint(int32(4))%32))
	v531 = *(*int64)(unsafe.Add(mBase, uint32(v530)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v525)+8)) = v531
	v533 = *(*int64)(unsafe.Add(mBase, uint32(v530)))
	*(*int64)(unsafe.Add(mBase, uint32(v525))) = v533
	v535 = *(*int64)(unsafe.Add(mBase, uint32(v17)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v530)+8)) = v535
	v537 = *(*int64)(unsafe.Add(mBase, uint32(v17)))
	*(*int64)(unsafe.Add(mBase, uint32(v530))) = v537
	v540 = v515 + int32(1)
	if v540 != v504 {
		v515 = v540
		goto L147
	} else {
		goto L149
	}
L148:
	;
	goto L146
L149:
	;
	goto L148
L150:
	;
	if base.Ui32(int32(2)) <= base.Ui32(v441) {
		goto L153
	} else {
		goto L154
	}
L151:
	;
	goto L152
L152:
	;
	if base.Ui32(int32(2)) <= base.Ui32(v497) {
		goto L158
	} else {
		goto L159
	}
L153:
	;
	F_qsort_ssup(m, v19, v441, l2)
	mBase = m.M
	v560 = m.ExcPending
	if v560 != 0 {
		goto L8
	} else {
		goto L156
	}
L154:
	;
	goto L155
L155:
	;
	if base.Ui32(v497) < base.Ui32(int32(2)) {
		goto L10
	} else {
		goto L157
	}
L156:
	;
	goto L155
L157:
	;
	v19 = v55 - v497<<(uint(int32(4))%32)
	v20 = v497
	goto L1
L158:
	;
	F_qsort_ssup(m, v55-v497<<(uint(int32(4))%32), v497, l2)
	mBase = m.M
	v572 = m.ExcPending
	if v572 != 0 {
		goto L8
	} else {
		goto L161
	}
L159:
	;
	goto L160
L160:
	;
	if base.Ui32(int32(1)) < base.Ui32(v441) {
		v36 = v441
		goto L3
	} else {
		goto L162
	}
L161:
	;
	goto L160
L162:
	;
	goto L10
}
