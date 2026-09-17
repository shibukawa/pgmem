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
	var v6 int32
	_ = v6
	var v15 int32
	_ = v15
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
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_InvalidateConstraintCacheCallBack[0]))
	v6 = int32(0)
	if base.B2i32(v5 == v6)|base.B2i32(v5 == int32(_a_F_InvalidateConstraintCacheCallBack_0)) == v6 {
		v15 = *(*int32)(unsafe.Add(mBase, _c_F_InvalidateConstraintCacheCallBack[1]))
		if base.Ui32(v15) <= base.Ui32(int32(1000)) {
			v18 = l2
		} else {
			v18 = int32(0)
		}
		v19 = v5
		for {
			v22 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
			if v18 == int32(0) {
				v35 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v19-int32(692)))) = uint8(v35)
				v37 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
				*(*int32)(unsafe.Add(mBase, uint32(v37)+4)) = v22
				v39 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
				*(*int32)(unsafe.Add(mBase, uint32(v22))) = v39
				v41 = int32(_a_F_InvalidateConstraintCacheCallBack_1)
				v43 = *(*int32)(unsafe.Add(mBase, _c_F_InvalidateConstraintCacheCallBack[1]))
				*(*int32)(unsafe.Add(mBase, _c_F_InvalidateConstraintCacheCallBack[1])) = v43 - int32(1)
			} else {
				v27 = *(*int32)(unsafe.Add(mBase, uint32(v19-int32(684))))
				if v27 == v18 {
					v35 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v19-int32(692)))) = uint8(v35)
					v37 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
					*(*int32)(unsafe.Add(mBase, uint32(v37)+4)) = v22
					v39 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
					*(*int32)(unsafe.Add(mBase, uint32(v22))) = v39
					v41 = int32(_a_F_InvalidateConstraintCacheCallBack_1)
					v43 = *(*int32)(unsafe.Add(mBase, _c_F_InvalidateConstraintCacheCallBack[1]))
					*(*int32)(unsafe.Add(mBase, _c_F_InvalidateConstraintCacheCallBack[1])) = v43 - int32(1)
				} else {
					v31 = *(*int32)(unsafe.Add(mBase, uint32(v19-int32(680))))
					if v31 != v18 {
					} else {
						v35 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(v19-int32(692)))) = uint8(v35)
						v37 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
						*(*int32)(unsafe.Add(mBase, uint32(v37)+4)) = v22
						v39 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
						*(*int32)(unsafe.Add(mBase, uint32(v22))) = v39
						v41 = int32(_a_F_InvalidateConstraintCacheCallBack_1)
						v43 = *(*int32)(unsafe.Add(mBase, _c_F_InvalidateConstraintCacheCallBack[1]))
						*(*int32)(unsafe.Add(mBase, _c_F_InvalidateConstraintCacheCallBack[1])) = v43 - int32(1)
					}
				}
			}
			if v22 != int32(_a_F_InvalidateConstraintCacheCallBack_0) {
				v19 = v22
				continue
			} else {
				break
			}
			break
		}
	} else {
	}
	return
}
func F_InvalidateObsoleteReplicationSlots(m *base.Module, l0 int32, l1 int64, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v31 int64
	_ = v31
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v45 int64
	_ = v45
	var v53 int32
	_ = v53
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v141 int64
	_ = v141
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v152 int64
	_ = v152
	var v153 int64
	_ = v153
	var v162 int64
	_ = v162
	var v163 int32
	_ = v163
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v174 int64
	_ = v174
	var v182 int32
	_ = v182
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v210 int32
	_ = v210
	var v224 int32
	_ = v224
	var v229 int32
	_ = v229
	var v236 int32
	_ = v236
	var v239 int64
	_ = v239
	var v242 int64
	_ = v242
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
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v261 int64
	_ = v261
	var v262 int64
	_ = v262
	var v264 int64
	_ = v264
	var v269 int32
	_ = v269
	var v271 int64
	_ = v271
	var v279 int32
	_ = v279
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v292 int32
	_ = v292
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v301 int64
	_ = v301
	var v302 int32
	_ = v302
	var v303 int64
	_ = v303
	var v305 int64
	_ = v305
	var v307 int64
	_ = v307
	var v309 int64
	_ = v309
	var v311 int64
	_ = v311
	var v313 int64
	_ = v313
	var v315 int64
	_ = v315
	var v317 int64
	_ = v317
	var v319 int32
	_ = v319
	var v323 int32
	_ = v323
	var v326 int32
	_ = v326
	var v329 int64
	_ = v329
	var v333 int32
	_ = v333
	var v342 int64
	_ = v342
	var v346 int64
	_ = v346
	var v347 int64
	_ = v347
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v365 int32
	_ = v365
	var v367 int64
	_ = v367
	var v369 int64
	_ = v369
	var v371 int64
	_ = v371
	var v373 int64
	_ = v373
	var v375 int64
	_ = v375
	var v377 int64
	_ = v377
	var v379 int64
	_ = v379
	var v381 int64
	_ = v381
	var v386 int32
	_ = v386
	var v388 int32
	_ = v388
	var v390 int32
	_ = v390
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v398 int32
	_ = v398
	var v402 int32
	_ = v402
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v422 int32
	_ = v422
	var v427 int32
	_ = v427
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v434 int32
	_ = v434
	var v437 int32
	_ = v437
	var v439 int32
	_ = v439
	var v440 int64
	_ = v440
	var v442 int64
	_ = v442
	var v444 int64
	_ = v444
	var v446 int64
	_ = v446
	var v448 int64
	_ = v448
	var v450 int64
	_ = v450
	var v452 int64
	_ = v452
	var v454 int64
	_ = v454
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
	var v464 int32
	_ = v464
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v470 int32
	_ = v470
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v478 int32
	_ = v478
	var v480 int32
	_ = v480
	var v484 int32
	_ = v484
	var v489 int32
	_ = v489
	var v493 int32
	_ = v493
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v500 int32
	_ = v500
	var v503 int32
	_ = v503
	var v510 int32
	_ = v510
	var v514 int32
	_ = v514
	var v525 int32
	_ = v525
	var v538 int32
	_ = v538
	var v550 int32
	_ = v550
	var v554 int32
	_ = v554
	var v559 int32
	_ = v559
	var v561 int32
	_ = v561
	var v564 int32
	_ = v564
	var v568 int32
	_ = v568
	var v574 int32
	_ = v574
	v5 = int32(0)
	v22 = m.G0
	v24 = v22 - int32(1232)
	m.G0 = v24
	v27 = *(*int32)(unsafe.Add(mBase, _c_F_InvalidateObsoleteReplicationSlots[0]))
	if v27 == v5 {
		v574 = v5
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v24 + int32(1232)
	return v574
L2:
	;
	v31 = int64(*(*int32)(unsafe.Add(mBase, _c_F_InvalidateObsoleteReplicationSlots[1])))
	v33 = *(*int32)(unsafe.Add(mBase, _c_F_InvalidateObsoleteReplicationSlots[2]))
	v37 = F_LWLockAcquire(m, v33+int32(_a_F_InvalidateObsoleteReplicationSlots_0), int32(1))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return int32(0)
L4:
	;
	v42 = *(*int32)(unsafe.Add(mBase, _c_F_InvalidateObsoleteReplicationSlots[0]))
	if int32(0) < v42 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v45 = l1 * v31
	v53 = l0 & int32(8)
	v61 = v42
	v65 = v5
	goto L8
L6:
	;
	goto L7
L7:
	;
	v564 = *(*int32)(unsafe.Add(mBase, _c_F_InvalidateObsoleteReplicationSlots[2]))
	F_LWLockRelease(m, v564+int32(_a_F_InvalidateObsoleteReplicationSlots_0))
	mBase = m.M
	v568 = m.ExcPending
	if v568 != 0 {
		goto L3
	} else {
		goto L122
	}
L8:
	;
	v76 = *(*int32)(unsafe.Add(mBase, _c_F_InvalidateObsoleteReplicationSlots[3]))
	v78 = v76
	v85 = v61
	v88 = int32(0)
	v89 = v65
	goto L10
L9:
	;
	v550 = *(*int32)(unsafe.Add(mBase, _c_F_InvalidateObsoleteReplicationSlots[2]))
	F_LWLockRelease(m, v550+int32(_a_F_InvalidateObsoleteReplicationSlots_0))
	mBase = m.M
	v554 = m.ExcPending
	if v554 != 0 {
		goto L3
	} else {
		goto L118
	}
L10:
	;
	v101 = v78 + v88*int32(288)
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101)+4)))
	if v102 != int32(1) {
		v503 = v78
		v510 = v85
		v514 = v89
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
	v525 = v88 + int32(1)
	if v525 < v510 {
		v78 = v503
		v85 = v510
		v88 = v525
		v89 = v514
		goto L10
	} else {
		goto L117
	}
L14:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v101)+88))
	if v105 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v107 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_InvalidateObsoleteReplicationSlots[4])))
	if v107&int32(1) != 0 {
		v503 = v78
		v510 = v85
		v514 = v89
		goto L13
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+140)) = int32(0)
	v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101)+4)))
	if v112 != int32(1) {
		v503 = v78
		v510 = v85
		v514 = v89
		goto L13
	} else {
		goto L19
	}
L18:
	;
	goto L17
L19:
	;
	v116 = v101 + int32(224)
	v118 = v101 + int32(24)
	v120 = int32(0)
	v122 = v120
	v129 = v120
	v133 = v89
	v141 = int64(0)
	goto L21
L20:
	;
	v493 = *(*int32)(unsafe.Add(mBase, _c_F_InvalidateObsoleteReplicationSlots[2]))
	v497 = F_LWLockAcquire(m, v493+int32(_a_F_InvalidateObsoleteReplicationSlots_0), int32(1))
	mBase = m.M
	v498 = m.ExcPending
	if v498 != 0 {
		goto L3
	} else {
		goto L115
	}
L21:
	;
	if v53 != 0 {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v480 = *(*int32)(unsafe.Add(mBase, _c_F_InvalidateObsoleteReplicationSlots[2]))
	F_LWLockRelease(m, v480+int32(_a_F_InvalidateObsoleteReplicationSlots_0))
	mBase = m.M
	v484 = m.ExcPending
	if v484 != 0 {
		goto L3
	} else {
		goto L114
	}
L23:
	;
	v147 = m.G0
	v148 = int32(16)
	v149 = v147 - v148
	m.G0 = v149
	F_gettimeofday(m, v149)
	mBase = m.M
	v152 = *(*int64)(unsafe.Add(mBase, uint32(v149)))
	v153 = int64(*(*int32)(unsafe.Add(mBase, uint32(v149)+8)))
	m.G0 = v149 + v148
	goto L26
L24:
	;
	v162 = int64(0)
	goto L25
L25:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v101)))
	*(*int32)(unsafe.Add(mBase, uint32(v101))) = int32(1)
	if v163 != 0 {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v162 = v153 + v152*int64(1000000) - int64(946684800000000)
	goto L25
L27:
	;
	F_s_lock(m, v101, int32(_a_F_InvalidateObsoleteReplicationSlots_1), int32(1875), int32(_a_F_InvalidateObsoleteReplicationSlots_2))
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L3
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v101)+112))
	if v171 != 0 {
		goto L34
	} else {
		goto L35
	}
L30:
	;
	goto L29
L31:
	;
	v303 = *(*int64)(unsafe.Add(mBase, uint32(v118)+56))
	*(*int64)(unsafe.Add(mBase, uint32(v24)+200)) = v303
	v305 = *(*int64)(unsafe.Add(mBase, uint32(v118)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v24)+192)) = v305
	v307 = *(*int64)(unsafe.Add(mBase, uint32(v118)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v24)+184)) = v307
	v309 = *(*int64)(unsafe.Add(mBase, uint32(v118)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v24)+176)) = v309
	v311 = *(*int64)(unsafe.Add(mBase, uint32(v118)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v24)+168)) = v311
	v313 = *(*int64)(unsafe.Add(mBase, uint32(v118)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v24)+160)) = v313
	v315 = *(*int64)(unsafe.Add(mBase, uint32(v118)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v24)+152)) = v315
	v317 = *(*int64)(unsafe.Add(mBase, uint32(v118)))
	*(*int64)(unsafe.Add(mBase, uint32(v24)+144)) = v317
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v101)+8))
	if v319 != 0 {
		v333 = v133
		goto L79
	} else {
		goto L80
	}
L32:
	;
	v297 = int32(0)
	v299 = v296
	v300 = v297
	v301 = v141
	v302 = v297
	goto L31
L33:
	;
	v296 = int32(2)
	goto L32
L34:
	;
	v279 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v101))) = v279
	if v122 == v279 {
		goto L75
	} else {
		goto L76
	}
L35:
	;
	v172 = int32(0)
	v174 = *(*int64)(unsafe.Add(mBase, uint32(v101)+104))
	if base.B2i32(l0&int32(1) == v172)|base.B2i32(v174 == int64(0))|base.B2i32(base.Ui64(v45) <= base.Ui64(v174)) == v172 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v182 = int32(1)
	v299 = v182
	v300 = v182
	v301 = v141
	v302 = int32(0)
	goto L31
L37:
	;
	goto L38
L38:
	;
	if l0&int32(2) == int32(0) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	if l0&int32(4) == int32(0) {
		goto L56
	} else {
		goto L57
	}
L40:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v101)+88))
	v188 = int32(0)
	if base.B2i32(v187 == v188)|base.B2i32(base.B2i32(l2 == v188)|base.B2i32(l2 == v187) == v188) != 0 {
		goto L39
	} else {
		goto L41
	}
L41:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v101)+20))
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v101)+16))
	if v198 != 0 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(l3))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v198)) == int32(0) {
		goto L46
	} else {
		goto L47
	}
L43:
	;
	goto L44
L44:
	;
	if v197 == int32(0) {
		goto L39
	} else {
		goto L50
	}
L45:
	;
	if v210 != 0 {
		goto L33
	} else {
		goto L49
	}
L46:
	;
	v210 = base.B2i32(base.Ui32(v198) <= base.Ui32(l3))
	goto L45
L47:
	;
	goto L48
L48:
	;
	v210 = base.B2i32(v198-l3 <= int32(0))
	goto L45
L49:
	;
	goto L44
L50:
	;
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(l3))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v197)) == int32(0) {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	if v224 != 0 {
		goto L33
	} else {
		goto L55
	}
L52:
	;
	v224 = base.B2i32(base.Ui32(v197) <= base.Ui32(l3))
	goto L51
L53:
	;
	goto L54
L54:
	;
	v224 = base.B2i32(v197-l3 <= int32(0))
	goto L51
L55:
	;
	goto L39
L56:
	;
	if v53 == int32(0) {
		goto L34
	} else {
		goto L59
	}
L57:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v101)+88))
	if v229 == int32(0) {
		goto L56
	} else {
		goto L58
	}
L58:
	;
	v296 = int32(4)
	goto L32
L59:
	;
	v236 = *(*int32)(unsafe.Add(mBase, _c_F_InvalidateObsoleteReplicationSlots[5]))
	if v236 == int32(0) {
		goto L34
	} else {
		goto L60
	}
L60:
	;
	v239 = *(*int64)(unsafe.Add(mBase, uint32(v101)+104))
	if v239 == int64(0) {
		goto L34
	} else {
		goto L61
	}
L61:
	;
	v242 = *(*int64)(unsafe.Add(mBase, uint32(v101)+272))
	if v242 <= int64(0) {
		goto L34
	} else {
		goto L62
	}
L62:
	;
	v247 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_InvalidateObsoleteReplicationSlots[6])))
	if v247 == int32(1) {
		goto L64
	} else {
		goto L65
	}
L63:
	;
	if v257 != 0 {
		goto L67
	} else {
		goto L68
	}
L64:
	;
	v252 = *(*int32)(unsafe.Add(mBase, _c_F_InvalidateObsoleteReplicationSlots[7]))
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v252)+316))
	v255 = base.B2i32(v253 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _c_F_InvalidateObsoleteReplicationSlots[6])) = uint8(v255)
	v257 = v255
	goto L66
L65:
	;
	v257 = int32(0)
	goto L66
L66:
	;
	goto L63
L67:
	;
	v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101)+201)))
	if v258 != 0 {
		goto L34
	} else {
		goto L70
	}
L68:
	;
	goto L69
L69:
	;
	v260 = *(*int32)(unsafe.Add(mBase, _c_F_InvalidateObsoleteReplicationSlots[5]))
	v261 = *(*int64)(unsafe.Add(mBase, uint32(v101)+272))
	v262 = v162 - v261
	v264 = base.I64_div_u_s(v262, int64(1000000))
	if int64(0) < v262 {
		goto L71
	} else {
		goto L72
	}
L70:
	;
	goto L69
L71:
	;
	v269 = base.I32_wrap_i64(v264)
	goto L73
L72:
	;
	v269 = int32(0)
	goto L73
L73:
	;
	if v269 < v260 {
		goto L34
	} else {
		goto L74
	}
L74:
	;
	v271 = *(*int64)(unsafe.Add(mBase, uint32(v101)+272))
	v299 = int32(8)
	v300 = int32(0)
	v301 = v271
	v302 = int32(1)
	goto L31
L75:
	;
	v284 = *(*int32)(unsafe.Add(mBase, _c_F_InvalidateObsoleteReplicationSlots[0]))
	v286 = *(*int32)(unsafe.Add(mBase, _c_F_InvalidateObsoleteReplicationSlots[3]))
	v503 = v286
	v510 = v284
	v514 = v133
	goto L13
L76:
	;
	goto L77
L77:
	;
	v288 = *(*int32)(unsafe.Add(mBase, _c_F_InvalidateObsoleteReplicationSlots[2]))
	F_LWLockRelease(m, v288+int32(_a_F_InvalidateObsoleteReplicationSlots_0))
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L3
	} else {
		goto L78
	}
L78:
	;
	v489 = v133
	goto L20
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v101))) = int32(0)
	if v302 != 0 {
		goto L82
	} else {
		goto L83
	}
L80:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_InvalidateObsoleteReplicationSlots[8])) = v101
	v323 = *(*int32)(unsafe.Add(mBase, _c_F_InvalidateObsoleteReplicationSlots[9]))
	*(*int32)(unsafe.Add(mBase, uint32(v101)+112)) = v299
	*(*int32)(unsafe.Add(mBase, uint32(v101)+8)) = v323
	v326 = int32(1)
	if v300 == int32(0) {
		v333 = v326
		goto L79
	} else {
		goto L81
	}
L81:
	;
	v329 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v101)+280)) = v329
	*(*int64)(unsafe.Add(mBase, uint32(v101)+104)) = v329
	v333 = v326
	goto L79
L82:
	;
	v342 = v162 - v301
	if v342 <= int64(0) {
		goto L86
	} else {
		goto L87
	}
L83:
	;
	goto L84
L84:
	;
	if v319 != 0 {
		goto L92
	} else {
		goto L93
	}
L85:
	;
	goto L84
L86:
	;
	v354 = int32(0)
	v355 = int32(0)
	goto L88
L87:
	;
	v346 = int64(1000000)
	v347 = base.I64_div_u_s(v342, v346)
	v354 = base.I32_wrap_i64(v347)
	v355 = base.I32_wrap_i64(v342 - v347*v346)
	goto L88
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24+int32(140)))) = v354
	*(*int32)(unsafe.Add(mBase, uint32(v24+int32(208)))) = v355
	goto L85
L89:
	;
	F_ConditionVariableSleep(m, v116, int32(134217777))
	mBase = m.M
	v467 = m.ExcPending
	if v467 != 0 {
		goto L3
	} else {
		goto L111
	}
L90:
	;
	v464 = v319
	goto L89
L91:
	;
	v462 = F_kill(m, v319, int32(15))
	mBase = m.M
	v463 = m.ExcPending
	if v463 != 0 {
		goto L3
	} else {
		goto L110
	}
L92:
	;
	F_ConditionVariablePrepareToSleep(m, v116)
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L3
	} else {
		goto L95
	}
L93:
	;
	goto L94
L94:
	;
	v398 = *(*int32)(unsafe.Add(mBase, _c_F_InvalidateObsoleteReplicationSlots[2]))
	F_LWLockRelease(m, v398+int32(_a_F_InvalidateObsoleteReplicationSlots_0))
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		goto L3
	} else {
		goto L101
	}
L95:
	;
	v361 = *(*int32)(unsafe.Add(mBase, _c_F_InvalidateObsoleteReplicationSlots[2]))
	F_LWLockRelease(m, v361+int32(_a_F_InvalidateObsoleteReplicationSlots_0))
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L3
	} else {
		goto L96
	}
L96:
	;
	if v319 == v129 {
		v464 = v129
		goto L89
	} else {
		goto L97
	}
L97:
	;
	v367 = *(*int64)(unsafe.Add(mBase, uint32(v24)+144))
	*(*int64)(unsafe.Add(mBase, uint32(v24)+72)) = v367
	v369 = *(*int64)(unsafe.Add(mBase, uint32(v24)+152))
	*(*int64)(unsafe.Add(mBase, uint32(v24)+80)) = v369
	v371 = *(*int64)(unsafe.Add(mBase, uint32(v24)+160))
	*(*int64)(unsafe.Add(mBase, uint32(v24)+88)) = v371
	v373 = *(*int64)(unsafe.Add(mBase, uint32(v24)+168))
	*(*int64)(unsafe.Add(mBase, uint32(v24)+96)) = v373
	v375 = *(*int64)(unsafe.Add(mBase, uint32(v24)+176))
	*(*int64)(unsafe.Add(mBase, uint32(v24)+104)) = v375
	v377 = *(*int64)(unsafe.Add(mBase, uint32(v24)+184))
	*(*int64)(unsafe.Add(mBase, uint32(v24)+112)) = v377
	v379 = *(*int64)(unsafe.Add(mBase, uint32(v24)+192))
	*(*int64)(unsafe.Add(mBase, uint32(v24)+120)) = v379
	v381 = *(*int64)(unsafe.Add(mBase, uint32(v24)+200))
	*(*int64)(unsafe.Add(mBase, uint32(v24)+128)) = v381
	v386 = *(*int32)(unsafe.Add(mBase, uint32(v24)+140))
	F_ReportSlotInvalidation(m, v299, int32(1), v319, v24+int32(72), v174, v45, l3, v386)
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L3
	} else {
		goto L98
	}
L98:
	;
	v390 = *(*int32)(unsafe.Add(mBase, _c_F_InvalidateObsoleteReplicationSlots[10]))
	if v390 != int32(13) {
		goto L91
	} else {
		goto L99
	}
L99:
	;
	v395 = F_SendProcSignal(m, v319, int32(11), int32(-1))
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L3
	} else {
		goto L100
	}
L100:
	;
	goto L90
L101:
	;
	v404 = *(*int32)(unsafe.Add(mBase, _c_F_InvalidateObsoleteReplicationSlots[8]))
	v405 = *(*int32)(unsafe.Add(mBase, uint32(v404)))
	*(*int32)(unsafe.Add(mBase, uint32(v404))) = int32(1)
	if v405 != 0 {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	F_s_lock(m, v404, int32(_a_F_InvalidateObsoleteReplicationSlots_1), int32(1107), int32(_a_F_InvalidateObsoleteReplicationSlots_3))
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L3
	} else {
		goto L105
	}
L103:
	;
	goto L104
L104:
	;
	v413 = int32(_a_F_InvalidateObsoleteReplicationSlots_4)
	v414 = *(*int32)(unsafe.Add(mBase, _c_F_InvalidateObsoleteReplicationSlots[8]))
	v415 = int32(257)
	*(*uint16)(unsafe.Add(mBase, uint32(v414)+12)) = uint16(v415)
	*(*int32)(unsafe.Add(mBase, uint32(v404))) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+64)) = int32(_a_F_InvalidateObsoleteReplicationSlots_5)
	v422 = *(*int32)(unsafe.Add(mBase, _c_F_InvalidateObsoleteReplicationSlots[8]))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+68)) = v422 + int32(24)
	v427 = v24 + int32(208)
	v431 = F_pg_sprintf(m, v427, int32(_a_F_InvalidateObsoleteReplicationSlots_6), v24-int32(-64))
	mBase = m.M
	v432 = m.ExcPending
	if v432 != 0 {
		goto L3
	} else {
		goto L106
	}
L105:
	;
	goto L104
L106:
	;
	v434 = *(*int32)(unsafe.Add(mBase, _c_F_InvalidateObsoleteReplicationSlots[8]))
	F_SaveSlotToPath(m, v434, v427, int32(21))
	mBase = m.M
	v437 = m.ExcPending
	if v437 != 0 {
		goto L3
	} else {
		goto L107
	}
L107:
	;
	F_ReplicationSlotRelease(m)
	mBase = m.M
	v439 = m.ExcPending
	if v439 != 0 {
		goto L3
	} else {
		goto L108
	}
L108:
	;
	v440 = *(*int64)(unsafe.Add(mBase, uint32(v24)+144))
	*(*int64)(unsafe.Add(mBase, uint32(v24))) = v440
	v442 = *(*int64)(unsafe.Add(mBase, uint32(v24)+152))
	*(*int64)(unsafe.Add(mBase, uint32(v24)+8)) = v442
	v444 = *(*int64)(unsafe.Add(mBase, uint32(v24)+160))
	*(*int64)(unsafe.Add(mBase, uint32(v24)+16)) = v444
	v446 = *(*int64)(unsafe.Add(mBase, uint32(v24)+168))
	*(*int64)(unsafe.Add(mBase, uint32(v24)+24)) = v446
	v448 = *(*int64)(unsafe.Add(mBase, uint32(v24)+176))
	*(*int64)(unsafe.Add(mBase, uint32(v24)+32)) = v448
	v450 = *(*int64)(unsafe.Add(mBase, uint32(v24)+184))
	*(*int64)(unsafe.Add(mBase, uint32(v24)+40)) = v450
	v452 = *(*int64)(unsafe.Add(mBase, uint32(v24)+192))
	*(*int64)(unsafe.Add(mBase, uint32(v24)+48)) = v452
	v454 = *(*int64)(unsafe.Add(mBase, uint32(v24)+200))
	*(*int64)(unsafe.Add(mBase, uint32(v24)+56)) = v454
	v456 = int32(0)
	v458 = *(*int32)(unsafe.Add(mBase, uint32(v24)+140))
	F_ReportSlotInvalidation(m, v299, v456, v456, v24, v174, v45, l3, v458)
	mBase = m.M
	v460 = m.ExcPending
	if v460 != 0 {
		goto L3
	} else {
		goto L109
	}
L109:
	;
	v489 = v333
	goto L20
L110:
	;
	goto L90
L111:
	;
	v468 = int32(1)
	v470 = *(*int32)(unsafe.Add(mBase, _c_F_InvalidateObsoleteReplicationSlots[2]))
	v474 = F_LWLockAcquire(m, v470+int32(_a_F_InvalidateObsoleteReplicationSlots_0), v468)
	mBase = m.M
	v475 = m.ExcPending
	if v475 != 0 {
		goto L3
	} else {
		goto L112
	}
L112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+140)) = int32(0)
	v478 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101)+4)))
	if v478 != 0 {
		v122 = v468
		v129 = v464
		v133 = v333
		v141 = v301
		goto L21
	} else {
		goto L113
	}
L113:
	;
	goto L22
L114:
	;
	v489 = v333
	goto L20
L115:
	;
	v500 = *(*int32)(unsafe.Add(mBase, _c_F_InvalidateObsoleteReplicationSlots[0]))
	if int32(0) < v500 {
		v61 = v500
		v65 = v489
		goto L8
	} else {
		goto L116
	}
L116:
	;
	v538 = v489
	goto L12
L117:
	;
	v538 = v514
	goto L12
L118:
	;
	if v538 == int32(0) {
		v574 = int32(0)
		goto L1
	} else {
		goto L119
	}
L119:
	;
	F_ReplicationSlotsComputeRequiredXmin(m, int32(0))
	mBase = m.M
	v559 = m.ExcPending
	if v559 != 0 {
		goto L3
	} else {
		goto L120
	}
L120:
	;
	F_ReplicationSlotsComputeRequiredLSN(m)
	mBase = m.M
	v561 = m.ExcPending
	if v561 != 0 {
		goto L3
	} else {
		goto L121
	}
L121:
	;
	v574 = int32(1)
	goto L1
L122:
	;
	v574 = v5
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
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_InvalidateOprCacheCallBack[0]))
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
	v24 = *(*int32)(unsafe.Add(mBase, _c_F_InvalidateOprCacheCallBack[0]))
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
	F_errmsg_internal(m, int32(_a_F_InvalidateOprCacheCallBack_0), int32(0))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	F_errfinish(m, int32(_a_F_InvalidateOprCacheCallBack_1), int32(1051), int32(_a_F_InvalidateOprCacheCallBack_2))
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
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	v4 = m.G0
	v6 = v4 - int32(32)
	m.G0 = v6
	v9 = v6 + int32(12)
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_InvalidateOprProofCacheCallBack[0]))
	F_hash_seq_init(m, v9, v11)
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
	v14 = F_hash_seq_search(m, v9)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	if v14 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v17 = v14
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
	v19 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v17)+8)) = uint16(v19)
	v23 = F_hash_seq_search(m, v6+int32(12))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L9
	}
L8:
	;
	goto L6
L9:
	;
	if v23 != 0 {
		v17 = v23
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
	*(*int32)(unsafe.Add(mBase, _c_F_invalidate_syncing_table_states[0])) = int32(0)
	return
}
