package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_LWLockRelease(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int64
	_ = v73
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v96 int32
	_ = v96
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v112 int64
	_ = v112
	var v116 int64
	_ = v116
	var v120 int64
	_ = v120
	var v123 int32
	_ = v123
	var v126 int64
	_ = v126
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v299 int32
	_ = v299
	var v304 int32
	_ = v304
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v333 int32
	_ = v333
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v356 int32
	_ = v356
	var v360 int32
	_ = v360
	var v366 int32
	_ = v366
	var v371 int32
	_ = v371
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v377 int32
	_ = v377
	var v380 int32
	_ = v380
	var v385 int32
	_ = v385
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v398 int32
	_ = v398
	var v401 int32
	_ = v401
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v421 int32
	_ = v421
	var v424 int32
	_ = v424
	var v443 int32
	_ = v443
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v479 int32
	_ = v479
	var v482 int32
	_ = v482
	var v485 int32
	_ = v485
	var v487 int32
	_ = v487
	var v490 int32
	_ = v490
	var v493 int32
	_ = v493
	var v496 int32
	_ = v496
	var v498 int32
	_ = v498
	var v512 int32
	_ = v512
	var v514 int32
	_ = v514
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v531 int32
	_ = v531
	var v536 int32
	_ = v536
	v2 = int32(0)
	v14 = m.G0
	v16 = v14 - int32(32)
	m.G0 = v16
	v19 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockRelease[0]))
	v21 = v19
	v22 = v2
	goto L2
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v524 = m.ExcPending
	if v524 != 0 {
		goto L36
	} else {
		goto L102
	}
L2:
	;
	v34 = v21 - int32(1)
	if v34 < int32(0) {
		goto L1
	} else {
		goto L4
	}
L3:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v40)+uint32(_c_F_LWLockRelease[1])))
	v49 = v19 - int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_LWLockRelease[0])) = v49
	if v49 <= v34 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	v40 = v34 << (uint(int32(3)) % 32)
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+uint32(_c_F_LWLockRelease[2])))
	if l0 != v41 {
		v21 = v34
		v22 = v22 + int32(1)
		goto L2
	} else {
		goto L5
	}
L5:
	;
	goto L3
L6:
	;
	if v45 != 0 {
		goto L19
	} else {
		goto L20
	}
L7:
	;
	v53 = v22 & int32(3)
	if v53 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v55 = v34
	v56 = int32(0)
	goto L11
L9:
	;
	v79 = v34
	goto L10
L10:
	;
	if base.Ui32(v22-int32(1)) < base.Ui32(int32(3)) {
		goto L6
	} else {
		goto L14
	}
L11:
	;
	v67 = int32(3)
	v69 = int32(1)
	v70 = v55 + v69
	v73 = *(*int64)(unsafe.Add(mBase, uint32(v70<<(uint(v67)%32))+uint32(_c_F_LWLockRelease[2])))
	*(*int64)(unsafe.Add(mBase, uint32(v55<<(uint(v67)%32))+uint32(_c_F_LWLockRelease[2]))) = v73
	v76 = v56 + v69
	if v76 != v53 {
		v55 = v70
		v56 = v76
		goto L11
	} else {
		goto L13
	}
L12:
	;
	v79 = v70
	goto L10
L13:
	;
	goto L12
L14:
	;
	v96 = v79
	goto L15
L15:
	;
	v108 = int32(3)
	v109 = v96 << (uint(v108) % 32)
	v112 = *(*int64)(unsafe.Add(mBase, uint32(v109)+uint32(_c_F_LWLockRelease[3])))
	*(*int64)(unsafe.Add(mBase, uint32(v109)+uint32(_c_F_LWLockRelease[2]))) = v112
	v116 = *(*int64)(unsafe.Add(mBase, uint32(v109)+uint32(_c_F_LWLockRelease[4])))
	*(*int64)(unsafe.Add(mBase, uint32(v109)+uint32(_c_F_LWLockRelease[3]))) = v116
	v120 = *(*int64)(unsafe.Add(mBase, uint32(v109)+uint32(_c_F_LWLockRelease[5])))
	*(*int64)(unsafe.Add(mBase, uint32(v109)+uint32(_c_F_LWLockRelease[4]))) = v120
	v123 = v96 + int32(4)
	v126 = *(*int64)(unsafe.Add(mBase, uint32(v123<<(uint(v108)%32))+uint32(_c_F_LWLockRelease[2])))
	*(*int64)(unsafe.Add(mBase, uint32(v109)+uint32(_c_F_LWLockRelease[5]))) = v126
	if v123 != v49 {
		v96 = v123
		goto L15
	} else {
		goto L17
	}
L16:
	;
	goto L6
L17:
	;
	goto L16
L18:
	;
	v512 = int32(_a_F_LWLockRelease_0)
	v514 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockRelease[6]))
	*(*int32)(unsafe.Add(mBase, _c_F_LWLockRelease[6])) = v514 - int32(1)
	m.G0 = v16 + int32(32)
	return
L19:
	;
	v144 = int32(-1)
	goto L21
L20:
	;
	v144 = int32(-262144)
	goto L21
L21:
	;
	if v45 != 0 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v147 = int32(1)
	goto L24
L23:
	;
	v147 = int32(_a_F_LWLockRelease_1)
	goto L24
L24:
	;
	v149 = base.AtomicRmwSub32(m, l0, int32(4), v147)
	if (v144+v149)&int32(-1073217537) != int32(-2147483648) {
		goto L18
	} else {
		goto L25
	}
L25:
	;
	v155 = int32(536870912)
	v157 = base.AtomicRmwOr32(m, l0, int32(4), v155)
	if v157&v155 != 0 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v161 = v157
	goto L29
L27:
	;
	goto L28
L28:
	;
	v259 = int32(-1)
	v260 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v260 == v259 {
		goto L52
	} else {
		goto L53
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+28)) = int32(_a_F_LWLockRelease_2)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+24)) = int32(860)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+20)) = int32(_a_F_LWLockRelease_3)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v16)+8)) = int64(0)
	if v161&int32(536870912) != 0 {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	goto L28
L31:
	;
	goto L34
L32:
	;
	goto L33
L33:
	;
	v221 = int32(_a_F_LWLockRelease_4)
	v222 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockRelease[7]))
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v16+int32(8))+8))
	if v224 == int32(0) {
		goto L42
	} else {
		goto L43
	}
L34:
	;
	F_perform_spin_delay(m, v16+int32(8))
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	goto L33
L36:
	;
	return
L37:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v202&int32(536870912) != 0 {
		goto L34
	} else {
		goto L38
	}
L38:
	;
	goto L35
L39:
	;
	v241 = int32(536870912)
	v243 = base.AtomicRmwOr32(m, l0, int32(4), v241)
	if v243&v241 != 0 {
		v161 = v243
		goto L29
	} else {
		goto L50
	}
L40:
	;
	goto L39
L41:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_LWLockRelease[7])) = v239
	goto L40
L42:
	;
	if int32(999) < v222 {
		goto L40
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	if v222 < int32(11) {
		goto L40
	} else {
		goto L49
	}
L45:
	;
	v229 = int32(900)
	if v229 <= v222 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v232 = v229
	goto L48
L47:
	;
	v232 = v222
	goto L48
L48:
	;
	v239 = v232 + int32(100)
	goto L41
L49:
	;
	v239 = v222 - int32(1)
	goto L41
L50:
	;
	goto L30
L51:
	;
	v392 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v398 = base.AtomicRmwCmpxchg32(m, l0, int32(4), v392, (v392&int32(-1610612737)|v385)&v391)
	if v392 != v398 {
		goto L82
	} else {
		goto L83
	}
L52:
	;
	v380 = v259
	v385 = v2
	v391 = int32(1610612735)
	goto L51
L53:
	;
	goto L54
L54:
	;
	v265 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockRelease[8]))
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v265)))
	v272 = v260
	v273 = v259
	v274 = v265
	v276 = int32(-1)
	v278 = v266 + v260*int32(768)
	v280 = v266
	v281 = v2
	v282 = v2
	goto L55
L55:
	;
	v285 = v272 * int32(768)
	v286 = v280 + v285
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v278)+348))
	if v281&int32(1) == int32(0) {
		goto L59
	} else {
		goto L60
	}
L56:
	;
	if v366&int32(1) != 0 {
		goto L76
	} else {
		goto L77
	}
L57:
	;
	goto L56
L58:
	;
	if v287 == int32(-1) {
		v360 = v347
		v366 = v352
		goto L57
	} else {
		goto L75
	}
L59:
	;
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v286)+348))
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v286)+352))
	if v295 == int32(-1) {
		goto L63
	} else {
		goto L64
	}
L60:
	;
	v292 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v286)+345)))
	if v292 != 0 {
		goto L59
	} else {
		goto L61
	}
L61:
	;
	v347 = v273
	v348 = v274
	v349 = v276
	v352 = v282
	v353 = int32(1)
	goto L58
L62:
	;
	if v294 == int32(-1) {
		goto L67
	} else {
		goto L68
	}
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v294
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v286)+352))
	v304 = v299
	goto L62
L64:
	;
	goto L65
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v280+v295*int32(768))+348)) = v294
	v304 = v295
	goto L62
L66:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v286)+348)) = int64(0)
	v318 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockRelease[8]))
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v318)))
	v320 = v319 + v285
	if v276 == int32(-1) {
		goto L71
	} else {
		goto L72
	}
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v304
	goto L66
L68:
	;
	goto L69
L69:
	;
	v309 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockRelease[8]))
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v309)))
	*(*int32)(unsafe.Add(mBase, uint32(v310+v294*int32(768))+352)) = v304
	goto L66
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v320)+348)) = int32(-1)
	v336 = int32(2)
	*(*uint8)(unsafe.Add(mBase, uint32(v286)+344)) = uint8(v336)
	v338 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v286)+345)))
	v340 = base.B2i32(v338 != v336)
	v341 = v340 | v282
	if v338 == int32(0) {
		v360 = v333
		v366 = v341
		goto L57
	} else {
		goto L74
	}
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v320)+352)) = int32(-1)
	v333 = v272
	goto L70
L72:
	;
	goto L73
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v320)+352)) = v276
	v327 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockRelease[8]))
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v327)))
	*(*int32)(unsafe.Add(mBase, uint32(v328+v276*int32(768))+348)) = v272
	v333 = v273
	goto L70
L74:
	;
	v345 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockRelease[8]))
	v347 = v333
	v348 = v345
	v349 = v272
	v352 = v341
	v353 = v340 | v281
	goto L58
L75:
	;
	v356 = *(*int32)(unsafe.Add(mBase, uint32(v348)))
	v272 = v287
	v273 = v347
	v274 = v348
	v276 = v349
	v278 = v356 + v287*int32(768)
	v280 = v356
	v281 = v353
	v282 = v352
	goto L55
L76:
	;
	v371 = int32(1073741824)
	goto L78
L77:
	;
	v371 = int32(0)
	goto L78
L78:
	;
	v373 = int32(-1)
	v374 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v374 == v373 {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v377 = int32(1610612735)
	goto L81
L80:
	;
	v377 = v373
	goto L81
L81:
	;
	v380 = v360
	v385 = v371
	v391 = v377
	goto L51
L82:
	;
	v401 = v398
	goto L85
L83:
	;
	goto L84
L84:
	;
	if v380 == int32(-1) {
		goto L18
	} else {
		goto L91
	}
L85:
	;
	v417 = int32(-1)
	v418 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v418 == v417 {
		goto L87
	} else {
		goto L88
	}
L86:
	;
	goto L84
L87:
	;
	v421 = int32(1610612735)
	goto L89
L88:
	;
	v421 = v417
	goto L89
L89:
	;
	v424 = base.AtomicRmwCmpxchg32(m, l0, int32(4), v401, (v401&int32(-1610612737)|v385)&v421)
	if v401 != v424 {
		v401 = v424
		goto L85
	} else {
		goto L90
	}
L90:
	;
	goto L86
L91:
	;
	v443 = v380
	goto L92
L92:
	;
	v455 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockRelease[8]))
	v456 = *(*int32)(unsafe.Add(mBase, uint32(v455)))
	v459 = v456 + v443*int32(768)
	v460 = *(*int32)(unsafe.Add(mBase, uint32(v459)+348))
	v461 = *(*int32)(unsafe.Add(mBase, uint32(v459)+352))
	if v461 != int32(-1) {
		goto L94
	} else {
		goto L95
	}
L93:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v459)+348)) = int64(0)
	v490 = int32(0)
	v493 = base.AtomicRmwOr32(m, v490, int32(_a_F_LWLockRelease_5), v490)
	*(*uint8)(unsafe.Add(mBase, uint32(v459)+344)) = uint8(v490)
	v496 = *(*int32)(unsafe.Add(mBase, uint32(v459)+332))
	F_PGSemaphoreUnlock(m, v496)
	mBase = m.M
	v498 = m.ExcPending
	if v498 != 0 {
		goto L36
	} else {
		goto L101
	}
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v456+v461*int32(768))+348)) = v460
	goto L96
L95:
	;
	goto L96
L96:
	;
	if v460 != int32(-1) {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	v471 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockRelease[8]))
	v472 = *(*int32)(unsafe.Add(mBase, uint32(v471)))
	*(*int32)(unsafe.Add(mBase, uint32(v472+v460*int32(768))+352)) = v461
	*(*int64)(unsafe.Add(mBase, uint32(v459)+348)) = int64(0)
	v479 = int32(0)
	v482 = base.AtomicRmwOr32(m, v479, int32(_a_F_LWLockRelease_5), v479)
	*(*uint8)(unsafe.Add(mBase, uint32(v459)+344)) = uint8(v479)
	v485 = *(*int32)(unsafe.Add(mBase, uint32(v459)+332))
	F_PGSemaphoreUnlock(m, v485)
	mBase = m.M
	v487 = m.ExcPending
	if v487 != 0 {
		goto L36
	} else {
		goto L100
	}
L98:
	;
	goto L99
L99:
	;
	goto L93
L100:
	;
	v443 = v460
	goto L92
L101:
	;
	goto L18
L102:
	;
	v525 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0))))
	v526 = F_GetLWTrancheName(m, v525)
	mBase = m.M
	v527 = m.ExcPending
	if v527 != 0 {
		goto L36
	} else {
		goto L103
	}
L103:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v526
	F_errmsg_internal(m, int32(_a_F_LWLockRelease_6), v16)
	mBase = m.M
	v531 = m.ExcPending
	if v531 != 0 {
		goto L36
	} else {
		goto L104
	}
L104:
	;
	F_errfinish(m, int32(_a_F_LWLockRelease_3), int32(1783), int32(_a_F_LWLockRelease_7))
	mBase = m.M
	v536 = m.ExcPending
	if v536 != 0 {
		goto L36
	} else {
		goto L105
	}
L105:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
