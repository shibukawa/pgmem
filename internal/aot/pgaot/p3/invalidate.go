package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_InvalidateConstraintCacheCallBack(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	v5 = *(*int32)(unsafe.Add(mBase, _consts[1053]))
	if v5 == int32(0) {
	} else {
		if v5 == int32(4548648) {
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, _consts[1054]))
			if base.Ui32(v12) <= base.Ui32(int32(1000)) {
				v15 = l2
			} else {
				v15 = int32(0)
			}
			v16 = v5
			for {
				v19 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
				if v15 == int32(0) {
					v32 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v16-int32(692)))) = uint8(v32)
					v34 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
					*(*int32)(unsafe.Add(mBase, uint32(v34)+4)) = v19
					v36 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
					*(*int32)(unsafe.Add(mBase, uint32(v19))) = v36
					v38 = int32(4548656)
					v40 = *(*int32)(unsafe.Add(mBase, _consts[1054]))
					*(*int32)(unsafe.Add(mBase, _consts[1054])) = v40 - int32(1)
				} else {
					v24 = *(*int32)(unsafe.Add(mBase, uint32(v16-int32(684))))
					if v24 == v15 {
						v32 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(v16-int32(692)))) = uint8(v32)
						v34 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
						*(*int32)(unsafe.Add(mBase, uint32(v34)+4)) = v19
						v36 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
						*(*int32)(unsafe.Add(mBase, uint32(v19))) = v36
						v38 = int32(4548656)
						v40 = *(*int32)(unsafe.Add(mBase, _consts[1054]))
						*(*int32)(unsafe.Add(mBase, _consts[1054])) = v40 - int32(1)
					} else {
						v28 = *(*int32)(unsafe.Add(mBase, uint32(v16-int32(680))))
						if v28 != v15 {
						} else {
							v32 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v16-int32(692)))) = uint8(v32)
							v34 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
							*(*int32)(unsafe.Add(mBase, uint32(v34)+4)) = v19
							v36 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
							*(*int32)(unsafe.Add(mBase, uint32(v19))) = v36
							v38 = int32(4548656)
							v40 = *(*int32)(unsafe.Add(mBase, _consts[1054]))
							*(*int32)(unsafe.Add(mBase, _consts[1054])) = v40 - int32(1)
						}
					}
				}
				if v19 != int32(4548648) {
					v16 = v19
					continue
				} else {
					break
				}
				break
			}
		}
	}
	return
}
func F_InvalidateObsoleteReplicationSlots(m *base.Module, l0 int32, l1 int64, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v38 int64
	_ = v38
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v52 int64
	_ = v52
	var v60 int32
	_ = v60
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v167 int64
	_ = v167
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v178 int64
	_ = v178
	var v179 int64
	_ = v179
	var v188 int64
	_ = v188
	var v189 int32
	_ = v189
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int64
	_ = v198
	var v204 int32
	_ = v204
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v231 int32
	_ = v231
	var v245 int32
	_ = v245
	var v250 int32
	_ = v250
	var v257 int32
	_ = v257
	var v260 int64
	_ = v260
	var v263 int64
	_ = v263
	var v268 int32
	_ = v268
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v282 int64
	_ = v282
	var v283 int64
	_ = v283
	var v285 int64
	_ = v285
	var v290 int32
	_ = v290
	var v292 int64
	_ = v292
	var v300 int32
	_ = v300
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v313 int32
	_ = v313
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v323 int64
	_ = v323
	var v324 int32
	_ = v324
	var v326 int32
	_ = v326
	var v327 int64
	_ = v327
	var v330 int32
	_ = v330
	var v331 int64
	_ = v331
	var v334 int32
	_ = v334
	var v335 int64
	_ = v335
	var v338 int32
	_ = v338
	var v339 int64
	_ = v339
	var v342 int32
	_ = v342
	var v343 int64
	_ = v343
	var v346 int32
	_ = v346
	var v347 int64
	_ = v347
	var v350 int32
	_ = v350
	var v351 int64
	_ = v351
	var v353 int64
	_ = v353
	var v355 int32
	_ = v355
	var v359 int32
	_ = v359
	var v362 int32
	_ = v362
	var v365 int64
	_ = v365
	var v369 int32
	_ = v369
	var v378 int64
	_ = v378
	var v382 int64
	_ = v382
	var v383 int64
	_ = v383
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v395 int32
	_ = v395
	var v397 int32
	_ = v397
	var v401 int32
	_ = v401
	var v403 int64
	_ = v403
	var v405 int64
	_ = v405
	var v407 int64
	_ = v407
	var v409 int64
	_ = v409
	var v411 int64
	_ = v411
	var v413 int64
	_ = v413
	var v415 int64
	_ = v415
	var v417 int64
	_ = v417
	var v422 int32
	_ = v422
	var v424 int32
	_ = v424
	var v426 int32
	_ = v426
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v434 int32
	_ = v434
	var v438 int32
	_ = v438
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v458 int32
	_ = v458
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v470 int32
	_ = v470
	var v475 int32
	_ = v475
	var v477 int32
	_ = v477
	var v478 int64
	_ = v478
	var v480 int64
	_ = v480
	var v482 int64
	_ = v482
	var v484 int64
	_ = v484
	var v486 int64
	_ = v486
	var v488 int64
	_ = v488
	var v490 int64
	_ = v490
	var v492 int64
	_ = v492
	var v494 int32
	_ = v494
	var v496 int32
	_ = v496
	var v498 int32
	_ = v498
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v508 int32
	_ = v508
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v516 int32
	_ = v516
	var v518 int32
	_ = v518
	var v522 int32
	_ = v522
	var v529 int32
	_ = v529
	var v541 int32
	_ = v541
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v548 int32
	_ = v548
	var v551 int32
	_ = v551
	var v557 int32
	_ = v557
	var v562 int32
	_ = v562
	var v580 int32
	_ = v580
	var v593 int32
	_ = v593
	var v612 int32
	_ = v612
	var v616 int32
	_ = v616
	var v623 int32
	_ = v623
	var v625 int32
	_ = v625
	var v628 int32
	_ = v628
	var v632 int32
	_ = v632
	var v638 int32
	_ = v638
	v5 = int32(0)
	v29 = m.G0
	v31 = v29 - int32(1232)
	m.G0 = v31
	v34 = *(*int32)(unsafe.Add(mBase, _consts[269]))
	if v34 == v5 {
		v638 = v5
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v31 + int32(1232)
	return v638
L2:
	;
	v38 = int64(*(*int32)(unsafe.Add(mBase, _consts[191])))
	v40 = *(*int32)(unsafe.Add(mBase, _consts[47]))
	v44 = F_LWLockAcquire(m, v40+int32(4736), int32(1))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return int32(0)
L4:
	;
	v49 = *(*int32)(unsafe.Add(mBase, _consts[269]))
	if int32(0) < v49 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v52 = l1 * v38
	v60 = l0 & int32(8)
	v67 = v49
	v72 = v5
	goto L8
L6:
	;
	goto L7
L7:
	;
	v628 = *(*int32)(unsafe.Add(mBase, _consts[47]))
	F_LWLockRelease(m, v628+int32(4736))
	mBase = m.M
	v632 = m.ExcPending
	if v632 != 0 {
		goto L3
	} else {
		goto L125
	}
L8:
	;
	v90 = *(*int32)(unsafe.Add(mBase, _consts[270]))
	v92 = v90
	v98 = v67
	v101 = int32(0)
	v103 = v72
	goto L10
L9:
	;
	v612 = *(*int32)(unsafe.Add(mBase, _consts[47]))
	F_LWLockRelease(m, v612+int32(4736))
	mBase = m.M
	v616 = m.ExcPending
	if v616 != 0 {
		goto L3
	} else {
		goto L121
	}
L10:
	;
	v122 = v92 + v101*int32(288)
	v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122)+4)))
	if v123 != int32(1) {
		v551 = v92
		v557 = v98
		v562 = v103
		goto L13
	} else {
		goto L14
	}
L11:
	;
	goto L9
L12:
	;
	goto L11
L13:
	;
	v580 = v101 + int32(1)
	if v580 < v557 {
		v92 = v551
		v98 = v557
		v101 = v580
		v103 = v562
		goto L10
	} else {
		goto L120
	}
L14:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v122)+88))
	if v126 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v128 = int32(*(*uint8)(unsafe.Add(mBase, _consts[714])))
	if v128 != 0 {
		v551 = v92
		v557 = v98
		v562 = v103
		goto L13
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+140)) = int32(0)
	v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122)+4)))
	if v131 != int32(1) {
		v551 = v92
		v557 = v98
		v562 = v103
		goto L13
	} else {
		goto L19
	}
L18:
	;
	goto L17
L19:
	;
	v135 = v122 + int32(224)
	v137 = v122 + int32(24)
	v139 = int32(0)
	v149 = v139
	v152 = v103
	v155 = v139
	v167 = int64(0)
	goto L21
L20:
	;
	v541 = *(*int32)(unsafe.Add(mBase, _consts[47]))
	v545 = F_LWLockAcquire(m, v541+int32(4736), int32(1))
	mBase = m.M
	v546 = m.ExcPending
	if v546 != 0 {
		goto L3
	} else {
		goto L118
	}
L21:
	;
	if v60 != 0 {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v518 = *(*int32)(unsafe.Add(mBase, _consts[47]))
	F_LWLockRelease(m, v518+int32(4736))
	mBase = m.M
	v522 = m.ExcPending
	if v522 != 0 {
		goto L3
	} else {
		goto L117
	}
L23:
	;
	v173 = m.G0
	v174 = int32(16)
	v175 = v173 - v174
	m.G0 = v175
	F___gettimeofday(m, v175)
	mBase = m.M
	v178 = *(*int64)(unsafe.Add(mBase, uint32(v175)))
	v179 = int64(*(*int32)(unsafe.Add(mBase, uint32(v175)+8)))
	m.G0 = v175 + v174
	goto L26
L24:
	;
	v188 = int64(0)
	goto L25
L25:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v122)))
	*(*int32)(unsafe.Add(mBase, uint32(v122))) = int32(1)
	if v189 != 0 {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v188 = v179 + v178*int64(1000000) - int64(946684800000000)
	goto L25
L27:
	;
	F_s_lock(m, v122, int32(518163), int32(1875), int32(92708))
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L3
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v122)+112))
	if v197 != 0 {
		goto L34
	} else {
		goto L35
	}
L30:
	;
	goto L29
L31:
	;
	v326 = v31 + int32(200)
	v327 = *(*int64)(unsafe.Add(mBase, uint32(v137)+56))
	*(*int64)(unsafe.Add(mBase, uint32(v326))) = v327
	v330 = v31 + int32(192)
	v331 = *(*int64)(unsafe.Add(mBase, uint32(v137)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v330))) = v331
	v334 = v31 + int32(184)
	v335 = *(*int64)(unsafe.Add(mBase, uint32(v137)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v334))) = v335
	v338 = v31 + int32(176)
	v339 = *(*int64)(unsafe.Add(mBase, uint32(v137)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v338))) = v339
	v342 = v31 + int32(168)
	v343 = *(*int64)(unsafe.Add(mBase, uint32(v137)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v342))) = v343
	v346 = v31 + int32(160)
	v347 = *(*int64)(unsafe.Add(mBase, uint32(v137)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v346))) = v347
	v350 = v31 + int32(152)
	v351 = *(*int64)(unsafe.Add(mBase, uint32(v137)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v350))) = v351
	v353 = *(*int64)(unsafe.Add(mBase, uint32(v137)))
	*(*int64)(unsafe.Add(mBase, uint32(v31)+144)) = v353
	v355 = *(*int32)(unsafe.Add(mBase, uint32(v122)+8))
	if v355 != 0 {
		v369 = v152
		goto L81
	} else {
		goto L82
	}
L32:
	;
	v318 = int32(0)
	v321 = v317
	v322 = v318
	v323 = v167
	v324 = v318
	goto L31
L33:
	;
	v317 = int32(2)
	goto L32
L34:
	;
	v300 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v122))) = v300
	if v149 == v300 {
		goto L77
	} else {
		goto L78
	}
L35:
	;
	v198 = *(*int64)(unsafe.Add(mBase, uint32(v122)+104))
	if l0&int32(1) == int32(0) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	if l0&int32(2) == int32(0) {
		goto L40
	} else {
		goto L41
	}
L37:
	;
	if v198 == int64(0) {
		goto L36
	} else {
		goto L38
	}
L38:
	;
	if base.Ui64(v52) <= base.Ui64(v198) {
		goto L36
	} else {
		goto L39
	}
L39:
	;
	v204 = int32(1)
	v321 = v204
	v322 = v204
	v323 = v167
	v324 = int32(0)
	goto L31
L40:
	;
	if l0&int32(4) == int32(0) {
		goto L58
	} else {
		goto L59
	}
L41:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v122)+88))
	if v209 == int32(0) {
		goto L40
	} else {
		goto L42
	}
L42:
	;
	v212 = int32(0)
	if base.B2i32(l2 == v212)|base.B2i32(l2 == v209) == v212 {
		goto L40
	} else {
		goto L43
	}
L43:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v122)+20))
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v122)+16))
	if v219 != 0 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(l3))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v219)) == int32(0) {
		goto L48
	} else {
		goto L49
	}
L45:
	;
	goto L46
L46:
	;
	if v218 == int32(0) {
		goto L40
	} else {
		goto L52
	}
L47:
	;
	if v231 != 0 {
		goto L33
	} else {
		goto L51
	}
L48:
	;
	v231 = base.B2i32(base.Ui32(v219) <= base.Ui32(l3))
	goto L47
L49:
	;
	goto L50
L50:
	;
	v231 = base.B2i32(v219-l3 <= int32(0))
	goto L47
L51:
	;
	goto L46
L52:
	;
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(l3))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v218)) == int32(0) {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	if v245 != 0 {
		goto L33
	} else {
		goto L57
	}
L54:
	;
	v245 = base.B2i32(base.Ui32(v218) <= base.Ui32(l3))
	goto L53
L55:
	;
	goto L56
L56:
	;
	v245 = base.B2i32(v218-l3 <= int32(0))
	goto L53
L57:
	;
	goto L40
L58:
	;
	if v60 == int32(0) {
		goto L34
	} else {
		goto L61
	}
L59:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v122)+88))
	if v250 == int32(0) {
		goto L58
	} else {
		goto L60
	}
L60:
	;
	v317 = int32(4)
	goto L32
L61:
	;
	v257 = *(*int32)(unsafe.Add(mBase, _consts[715]))
	if v257 == int32(0) {
		goto L34
	} else {
		goto L62
	}
L62:
	;
	v260 = *(*int64)(unsafe.Add(mBase, uint32(v122)+104))
	if v260 == int64(0) {
		goto L34
	} else {
		goto L63
	}
L63:
	;
	v263 = *(*int64)(unsafe.Add(mBase, uint32(v122)+272))
	if v263 <= int64(0) {
		goto L34
	} else {
		goto L64
	}
L64:
	;
	v268 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1])))
	if v268 == int32(1) {
		goto L66
	} else {
		goto L67
	}
L65:
	;
	if v278 != 0 {
		goto L69
	} else {
		goto L70
	}
L66:
	;
	v273 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v273)+316))
	v276 = base.B2i32(v274 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _consts[1])) = uint8(v276)
	v278 = v276
	goto L68
L67:
	;
	v278 = int32(0)
	goto L68
L68:
	;
	goto L65
L69:
	;
	v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122)+201)))
	if v279 != 0 {
		goto L34
	} else {
		goto L72
	}
L70:
	;
	goto L71
L71:
	;
	v281 = *(*int32)(unsafe.Add(mBase, _consts[715]))
	v282 = *(*int64)(unsafe.Add(mBase, uint32(v122)+272))
	v283 = v188 - v282
	v285 = base.I64_div_u_s(v283, int64(1000000))
	if int64(0) < v283 {
		goto L73
	} else {
		goto L74
	}
L72:
	;
	goto L71
L73:
	;
	v290 = base.I32_wrap_i64(v285)
	goto L75
L74:
	;
	v290 = int32(0)
	goto L75
L75:
	;
	if v290 < v281 {
		goto L34
	} else {
		goto L76
	}
L76:
	;
	v292 = *(*int64)(unsafe.Add(mBase, uint32(v122)+272))
	v321 = int32(8)
	v322 = int32(0)
	v323 = v292
	v324 = int32(1)
	goto L31
L77:
	;
	v305 = *(*int32)(unsafe.Add(mBase, _consts[269]))
	v307 = *(*int32)(unsafe.Add(mBase, _consts[270]))
	v551 = v307
	v557 = v305
	v562 = v152
	goto L13
L78:
	;
	goto L79
L79:
	;
	v309 = *(*int32)(unsafe.Add(mBase, _consts[47]))
	F_LWLockRelease(m, v309+int32(4736))
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L3
	} else {
		goto L80
	}
L80:
	;
	v529 = v152
	goto L20
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v122))) = int32(0)
	if v324 != 0 {
		goto L84
	} else {
		goto L85
	}
L82:
	;
	*(*int32)(unsafe.Add(mBase, _consts[678])) = v122
	v359 = *(*int32)(unsafe.Add(mBase, _consts[140]))
	*(*int32)(unsafe.Add(mBase, uint32(v122)+112)) = v321
	*(*int32)(unsafe.Add(mBase, uint32(v122)+8)) = v359
	v362 = int32(1)
	if v322 == int32(0) {
		v369 = v362
		goto L81
	} else {
		goto L83
	}
L83:
	;
	v365 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v122)+280)) = v365
	*(*int64)(unsafe.Add(mBase, uint32(v122)+104)) = v365
	v369 = v362
	goto L81
L84:
	;
	v378 = v188 - v323
	if v378 <= int64(0) {
		goto L89
	} else {
		goto L90
	}
L85:
	;
	goto L86
L86:
	;
	if v355 != 0 {
		goto L95
	} else {
		goto L96
	}
L87:
	;
	goto L86
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31+int32(140)))) = v390
	*(*int32)(unsafe.Add(mBase, uint32(v31+int32(208)))) = v391
	goto L87
L89:
	;
	v390 = int32(0)
	v391 = int32(0)
	goto L88
L90:
	;
	goto L91
L91:
	;
	v382 = int64(1000000)
	v383 = base.I64_div_u_s(v378, v382)
	v390 = base.I32_wrap_i64(v383)
	v391 = base.I32_wrap_i64(v378 - v383*v382)
	goto L88
L92:
	;
	F_ConditionVariableSleep(m, v135, int32(134217777))
	mBase = m.M
	v505 = m.ExcPending
	if v505 != 0 {
		goto L3
	} else {
		goto L114
	}
L93:
	;
	v502 = v355
	goto L92
L94:
	;
	v500 = F_kill(m, v355, int32(15))
	mBase = m.M
	v501 = m.ExcPending
	if v501 != 0 {
		goto L3
	} else {
		goto L113
	}
L95:
	;
	F_ConditionVariablePrepareToSleep(m, v135)
	mBase = m.M
	v395 = m.ExcPending
	if v395 != 0 {
		goto L3
	} else {
		goto L98
	}
L96:
	;
	goto L97
L97:
	;
	v434 = *(*int32)(unsafe.Add(mBase, _consts[47]))
	F_LWLockRelease(m, v434+int32(4736))
	mBase = m.M
	v438 = m.ExcPending
	if v438 != 0 {
		goto L3
	} else {
		goto L104
	}
L98:
	;
	v397 = *(*int32)(unsafe.Add(mBase, _consts[47]))
	F_LWLockRelease(m, v397+int32(4736))
	mBase = m.M
	v401 = m.ExcPending
	if v401 != 0 {
		goto L3
	} else {
		goto L99
	}
L99:
	;
	if v355 == v155 {
		v502 = v155
		goto L92
	} else {
		goto L100
	}
L100:
	;
	v403 = *(*int64)(unsafe.Add(mBase, uint32(v350)))
	*(*int64)(unsafe.Add(mBase, uint32(v31)+80)) = v403
	v405 = *(*int64)(unsafe.Add(mBase, uint32(v346)))
	*(*int64)(unsafe.Add(mBase, uint32(v31)+88)) = v405
	v407 = *(*int64)(unsafe.Add(mBase, uint32(v342)))
	*(*int64)(unsafe.Add(mBase, uint32(v31)+96)) = v407
	v409 = *(*int64)(unsafe.Add(mBase, uint32(v338)))
	*(*int64)(unsafe.Add(mBase, uint32(v31)+104)) = v409
	v411 = *(*int64)(unsafe.Add(mBase, uint32(v334)))
	*(*int64)(unsafe.Add(mBase, uint32(v31)+112)) = v411
	v413 = *(*int64)(unsafe.Add(mBase, uint32(v330)))
	*(*int64)(unsafe.Add(mBase, uint32(v31)+120)) = v413
	v415 = *(*int64)(unsafe.Add(mBase, uint32(v326)))
	*(*int64)(unsafe.Add(mBase, uint32(v31)+128)) = v415
	v417 = *(*int64)(unsafe.Add(mBase, uint32(v31)+144))
	*(*int64)(unsafe.Add(mBase, uint32(v31)+72)) = v417
	v422 = *(*int32)(unsafe.Add(mBase, uint32(v31)+140))
	F_ReportSlotInvalidation(m, v321, int32(1), v355, v31+int32(72), v198, v52, l3, v422)
	mBase = m.M
	v424 = m.ExcPending
	if v424 != 0 {
		goto L3
	} else {
		goto L101
	}
L101:
	;
	v426 = *(*int32)(unsafe.Add(mBase, _consts[200]))
	if v426 != int32(13) {
		goto L94
	} else {
		goto L102
	}
L102:
	;
	v431 = F_SendProcSignal(m, v355, int32(11), int32(-1))
	mBase = m.M
	v432 = m.ExcPending
	if v432 != 0 {
		goto L3
	} else {
		goto L103
	}
L103:
	;
	goto L93
L104:
	;
	v440 = *(*int32)(unsafe.Add(mBase, _consts[678]))
	v441 = *(*int32)(unsafe.Add(mBase, uint32(v440)))
	*(*int32)(unsafe.Add(mBase, uint32(v440))) = int32(1)
	if v441 != 0 {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	F_s_lock(m, v440, int32(518163), int32(1107), int32(8439))
	mBase = m.M
	v448 = m.ExcPending
	if v448 != 0 {
		goto L3
	} else {
		goto L108
	}
L106:
	;
	goto L107
L107:
	;
	v449 = int32(4472308)
	v450 = *(*int32)(unsafe.Add(mBase, _consts[678]))
	v451 = int32(257)
	*(*uint16)(unsafe.Add(mBase, uint32(v450)+12)) = uint16(v451)
	*(*int32)(unsafe.Add(mBase, uint32(v440))) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+64)) = int32(91132)
	v458 = *(*int32)(unsafe.Add(mBase, _consts[678]))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+68)) = v458 + int32(24)
	v467 = F_pg_sprintf(m, v31+int32(208), int32(188238), v31-int32(-64))
	mBase = m.M
	v468 = m.ExcPending
	if v468 != 0 {
		goto L3
	} else {
		goto L109
	}
L108:
	;
	goto L107
L109:
	;
	v470 = *(*int32)(unsafe.Add(mBase, _consts[678]))
	F_SaveSlotToPath(m, v470, v31+int32(208), int32(21))
	mBase = m.M
	v475 = m.ExcPending
	if v475 != 0 {
		goto L3
	} else {
		goto L110
	}
L110:
	;
	F_ReplicationSlotRelease(m)
	mBase = m.M
	v477 = m.ExcPending
	if v477 != 0 {
		goto L3
	} else {
		goto L111
	}
L111:
	;
	v478 = *(*int64)(unsafe.Add(mBase, uint32(v350)))
	*(*int64)(unsafe.Add(mBase, uint32(v31)+8)) = v478
	v480 = *(*int64)(unsafe.Add(mBase, uint32(v346)))
	*(*int64)(unsafe.Add(mBase, uint32(v31)+16)) = v480
	v482 = *(*int64)(unsafe.Add(mBase, uint32(v342)))
	*(*int64)(unsafe.Add(mBase, uint32(v31)+24)) = v482
	v484 = *(*int64)(unsafe.Add(mBase, uint32(v338)))
	*(*int64)(unsafe.Add(mBase, uint32(v31)+32)) = v484
	v486 = *(*int64)(unsafe.Add(mBase, uint32(v334)))
	*(*int64)(unsafe.Add(mBase, uint32(v31)+40)) = v486
	v488 = *(*int64)(unsafe.Add(mBase, uint32(v330)))
	*(*int64)(unsafe.Add(mBase, uint32(v31)+48)) = v488
	v490 = *(*int64)(unsafe.Add(mBase, uint32(v326)))
	*(*int64)(unsafe.Add(mBase, uint32(v31)+56)) = v490
	v492 = *(*int64)(unsafe.Add(mBase, uint32(v31)+144))
	*(*int64)(unsafe.Add(mBase, uint32(v31))) = v492
	v494 = int32(0)
	v496 = *(*int32)(unsafe.Add(mBase, uint32(v31)+140))
	F_ReportSlotInvalidation(m, v321, v494, v494, v31, v198, v52, l3, v496)
	mBase = m.M
	v498 = m.ExcPending
	if v498 != 0 {
		goto L3
	} else {
		goto L112
	}
L112:
	;
	v529 = v369
	goto L20
L113:
	;
	goto L93
L114:
	;
	v506 = int32(1)
	v508 = *(*int32)(unsafe.Add(mBase, _consts[47]))
	v512 = F_LWLockAcquire(m, v508+int32(4736), v506)
	mBase = m.M
	v513 = m.ExcPending
	if v513 != 0 {
		goto L3
	} else {
		goto L115
	}
L115:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+140)) = int32(0)
	v516 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122)+4)))
	if v516 != 0 {
		v149 = v506
		v152 = v369
		v155 = v502
		v167 = v323
		goto L21
	} else {
		goto L116
	}
L116:
	;
	goto L22
L117:
	;
	v529 = v369
	goto L20
L118:
	;
	v548 = *(*int32)(unsafe.Add(mBase, _consts[269]))
	if int32(0) < v548 {
		v67 = v548
		v72 = v529
		goto L8
	} else {
		goto L119
	}
L119:
	;
	v593 = v529
	goto L12
L120:
	;
	v593 = v562
	goto L12
L121:
	;
	if v593&int32(1) == int32(0) {
		v638 = int32(0)
		goto L1
	} else {
		goto L122
	}
L122:
	;
	F_ReplicationSlotsComputeRequiredXmin(m, int32(0))
	mBase = m.M
	v623 = m.ExcPending
	if v623 != 0 {
		goto L3
	} else {
		goto L123
	}
L123:
	;
	F_ReplicationSlotsComputeRequiredLSN(m)
	mBase = m.M
	v625 = m.ExcPending
	if v625 != 0 {
		goto L3
	} else {
		goto L124
	}
L124:
	;
	v638 = int32(1)
	goto L1
L125:
	;
	v638 = v5
	goto L1
}
func F_InvalidateOprCacheCallBack(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
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
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	v4 = m.G0
	v6 = v4 - int32(32)
	m.G0 = v6
	v11 = *(*int32)(unsafe.Add(mBase, _consts[456]))
	F_hash_seq_init(m, v6+int32(12), v11)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	goto L4
L3:
	;
	m.G0 = v6 + int32(32)
	return
L4:
	;
	v19 = F_hash_seq_search(m, v6+int32(12))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L6
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L1
	} else {
		goto L10
	}
L6:
	;
	if v19 == int32(0) {
		goto L3
	} else {
		goto L7
	}
L7:
	;
	v24 = *(*int32)(unsafe.Add(mBase, _consts[456]))
	v27 = F_hash_search(m, v24, v19, int32(2), int32(0))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	if v27 != 0 {
		goto L4
	} else {
		goto L9
	}
L9:
	;
	goto L5
L10:
	;
	F_errmsg_internal(m, int32(467571), int32(0))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	F_errfinish(m, int32(521082), int32(1051), int32(336357))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_InvalidateOprProofCacheCallBack(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	v4 = m.G0
	v6 = v4 - int32(32)
	m.G0 = v6
	v11 = *(*int32)(unsafe.Add(mBase, _consts[617]))
	F_hash_seq_init(m, v6+int32(12), v11)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v16 = F_hash_seq_search(m, v6+int32(12))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	if v16 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v19 = v16
	goto L7
L5:
	;
	goto L6
L6:
	;
	m.G0 = v6 + int32(32)
	return
L7:
	;
	v21 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v19)+8)) = uint16(v21)
	v25 = F_hash_seq_search(m, v6+int32(12))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L9
	}
L8:
	;
	goto L6
L9:
	;
	if v25 != 0 {
		v19 = v25
		goto L7
	} else {
		goto L10
	}
L10:
	;
	goto L8
}
func F_InvalidatePublicationRels(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	v2 = int32(0)
	if l0 == v2 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v5 <= int32(4095) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v8 <= int32(0) {
		goto L1
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	F_CacheInvalidateRelcacheAll(m)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L9
	} else {
		goto L12
	}
L6:
	;
	v12 = v2
	goto L7
L7:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v13+v12<<(uint(int32(2))%32))))
	F_CacheInvalidateRelcacheByRelid(m, v17)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	goto L1
L9:
	;
	return
L10:
	;
	v21 = v12 + int32(1)
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v21 < v22 {
		v12 = v21
		goto L7
	} else {
		goto L11
	}
L11:
	;
	goto L8
L12:
	;
	goto L1
}
func F_InvalidateSystemCaches(m *base.Module) {
	var v2 int32
	_ = v2
	F_InvalidateSystemCachesExtended(m)
	v2 = m.ExcPending
	if v2 != 0 {
		return
	} else {
		return
	}
}
func F_invalidate_syncing_table_states(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	*(*int32)(unsafe.Add(mBase, _consts[686])) = int32(0)
	return
}
