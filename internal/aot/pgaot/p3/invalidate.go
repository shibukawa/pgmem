package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"sync/atomic"
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
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
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
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
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
	var v165 int32
	_ = v165
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
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v295 int32
	_ = v295
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v305 int64
	_ = v305
	var v306 int32
	_ = v306
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
	var v319 int64
	_ = v319
	var v321 int64
	_ = v321
	var v323 int32
	_ = v323
	var v327 int32
	_ = v327
	var v330 int32
	_ = v330
	var v333 int64
	_ = v333
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v347 int64
	_ = v347
	var v351 int64
	_ = v351
	var v352 int64
	_ = v352
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v364 int32
	_ = v364
	var v366 int32
	_ = v366
	var v370 int32
	_ = v370
	var v372 int64
	_ = v372
	var v374 int64
	_ = v374
	var v376 int64
	_ = v376
	var v378 int64
	_ = v378
	var v380 int64
	_ = v380
	var v382 int64
	_ = v382
	var v384 int64
	_ = v384
	var v386 int64
	_ = v386
	var v391 int32
	_ = v391
	var v393 int32
	_ = v393
	var v395 int32
	_ = v395
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v403 int32
	_ = v403
	var v407 int32
	_ = v407
	var v409 int32
	_ = v409
	var v412 int32
	_ = v412
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v422 int32
	_ = v422
	var v428 int32
	_ = v428
	var v433 int32
	_ = v433
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v440 int32
	_ = v440
	var v443 int32
	_ = v443
	var v445 int32
	_ = v445
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
	var v456 int64
	_ = v456
	var v458 int64
	_ = v458
	var v460 int64
	_ = v460
	var v462 int32
	_ = v462
	var v464 int32
	_ = v464
	var v466 int32
	_ = v466
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v475 int32
	_ = v475
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v483 int32
	_ = v483
	var v485 int32
	_ = v485
	var v489 int32
	_ = v489
	var v494 int32
	_ = v494
	var v498 int32
	_ = v498
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v505 int32
	_ = v505
	var v508 int32
	_ = v508
	var v514 int32
	_ = v514
	var v519 int32
	_ = v519
	var v530 int32
	_ = v530
	var v543 int32
	_ = v543
	var v555 int32
	_ = v555
	var v559 int32
	_ = v559
	var v564 int32
	_ = v564
	var v566 int32
	_ = v566
	var v569 int32
	_ = v569
	var v573 int32
	_ = v573
	var v582 int32
	_ = v582
	v5 = int32(0)
	v22 = m.G0
	v24 = v22 - int32(1232)
	m.G0 = v24
	v27 = *(*int32)(unsafe.Add(mBase, _c_F_InvalidateObsoleteReplicationSlots[0]))
	if v27 == v5 {
		v582 = v5
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v24 + int32(1232)
	return v582
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
	v60 = v42
	v65 = v5
	goto L8
L6:
	;
	goto L7
L7:
	;
	v569 = *(*int32)(unsafe.Add(mBase, _c_F_InvalidateObsoleteReplicationSlots[2]))
	F_LWLockRelease(m, v569+int32(_a_F_InvalidateObsoleteReplicationSlots_0))
	mBase = m.M
	v573 = m.ExcPending
	if v573 != 0 {
		goto L3
	} else {
		goto L121
	}
L8:
	;
	v76 = *(*int32)(unsafe.Add(mBase, _c_F_InvalidateObsoleteReplicationSlots[3]))
	v78 = v76
	v84 = v60
	v88 = int32(0)
	v89 = v65
	goto L10
L9:
	;
	v555 = *(*int32)(unsafe.Add(mBase, _c_F_InvalidateObsoleteReplicationSlots[2]))
	F_LWLockRelease(m, v555+int32(_a_F_InvalidateObsoleteReplicationSlots_0))
	mBase = m.M
	v559 = m.ExcPending
	if v559 != 0 {
		goto L3
	} else {
		goto L117
	}
L10:
	;
	v101 = v78 + v88*int32(288)
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101)+4)))
	if v102 != int32(1) {
		v508 = v78
		v514 = v84
		v519 = v89
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
	v530 = v88 + int32(1)
	if v530 < v514 {
		v78 = v508
		v84 = v514
		v88 = v530
		v89 = v519
		goto L10
	} else {
		goto L116
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
		v508 = v78
		v514 = v84
		v519 = v89
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
		v508 = v78
		v514 = v84
		v519 = v89
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
	v128 = v120
	v130 = v120
	v133 = v89
	v141 = int64(0)
	goto L21
L20:
	;
	v498 = *(*int32)(unsafe.Add(mBase, _c_F_InvalidateObsoleteReplicationSlots[2]))
	v502 = F_LWLockAcquire(m, v498+int32(_a_F_InvalidateObsoleteReplicationSlots_0), int32(1))
	mBase = m.M
	v503 = m.ExcPending
	if v503 != 0 {
		goto L3
	} else {
		goto L114
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
	v485 = *(*int32)(unsafe.Add(mBase, _c_F_InvalidateObsoleteReplicationSlots[2]))
	F_LWLockRelease(m, v485+int32(_a_F_InvalidateObsoleteReplicationSlots_0))
	mBase = m.M
	v489 = m.ExcPending
	if v489 != 0 {
		goto L3
	} else {
		goto L113
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
	v165 = base.AtomicRmwXchg32(m, v101, int32(0), int32(1))
	if v165 != 0 {
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
	v307 = *(*int64)(unsafe.Add(mBase, uint32(v118)+56))
	*(*int64)(unsafe.Add(mBase, uint32(v24)+200)) = v307
	v309 = *(*int64)(unsafe.Add(mBase, uint32(v118)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v24)+192)) = v309
	v311 = *(*int64)(unsafe.Add(mBase, uint32(v118)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v24)+184)) = v311
	v313 = *(*int64)(unsafe.Add(mBase, uint32(v118)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v24)+176)) = v313
	v315 = *(*int64)(unsafe.Add(mBase, uint32(v118)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v24)+168)) = v315
	v317 = *(*int64)(unsafe.Add(mBase, uint32(v118)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v24)+160)) = v317
	v319 = *(*int64)(unsafe.Add(mBase, uint32(v118)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v24)+152)) = v319
	v321 = *(*int64)(unsafe.Add(mBase, uint32(v118)))
	*(*int64)(unsafe.Add(mBase, uint32(v24)+144)) = v321
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v101)+8))
	if v323 != 0 {
		v337 = v133
		goto L79
	} else {
		goto L80
	}
L32:
	;
	v300 = int32(0)
	v303 = v299
	v304 = v300
	v305 = v141
	v306 = v300
	goto L31
L33:
	;
	v299 = int32(2)
	goto L32
L34:
	;
	v279 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v101))), uint32(v279))
	if v128&int32(1) == v279 {
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
	v303 = v182
	v304 = v182
	v305 = v141
	v306 = int32(0)
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
	if base.B2i32(v187 == v188)|base.B2i32(base.B2i32(l2 == v188)|base.B2i32(v187 == l2) == v188) != 0 {
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
	v299 = int32(4)
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
	v303 = int32(8)
	v304 = int32(0)
	v305 = v271
	v306 = int32(1)
	goto L31
L75:
	;
	v287 = *(*int32)(unsafe.Add(mBase, _c_F_InvalidateObsoleteReplicationSlots[0]))
	v289 = *(*int32)(unsafe.Add(mBase, _c_F_InvalidateObsoleteReplicationSlots[3]))
	v508 = v289
	v514 = v287
	v519 = v133
	goto L13
L76:
	;
	goto L77
L77:
	;
	v291 = *(*int32)(unsafe.Add(mBase, _c_F_InvalidateObsoleteReplicationSlots[2]))
	F_LWLockRelease(m, v291+int32(_a_F_InvalidateObsoleteReplicationSlots_0))
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L3
	} else {
		goto L78
	}
L78:
	;
	v494 = v133
	goto L20
L79:
	;
	v338 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v101))), uint32(v338))
	if v306 != 0 {
		goto L82
	} else {
		goto L83
	}
L80:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_InvalidateObsoleteReplicationSlots[8])) = v101
	v327 = *(*int32)(unsafe.Add(mBase, _c_F_InvalidateObsoleteReplicationSlots[9]))
	*(*int32)(unsafe.Add(mBase, uint32(v101)+112)) = v303
	*(*int32)(unsafe.Add(mBase, uint32(v101)+8)) = v327
	v330 = int32(1)
	if v304 == int32(0) {
		v337 = v330
		goto L79
	} else {
		goto L81
	}
L81:
	;
	v333 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v101)+280)) = v333
	*(*int64)(unsafe.Add(mBase, uint32(v101)+104)) = v333
	v337 = v330
	goto L79
L82:
	;
	v347 = v162 - v305
	if v347 <= int64(0) {
		goto L86
	} else {
		goto L87
	}
L83:
	;
	goto L84
L84:
	;
	if v323 != 0 {
		goto L92
	} else {
		goto L93
	}
L85:
	;
	goto L84
L86:
	;
	v359 = int32(0)
	v360 = int32(0)
	goto L88
L87:
	;
	v351 = int64(1000000)
	v352 = base.I64_div_u_s(v347, v351)
	v359 = base.I32_wrap_i64(v352)
	v360 = base.I32_wrap_i64(v347 - v352*v351)
	goto L88
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24+int32(140)))) = v359
	*(*int32)(unsafe.Add(mBase, uint32(v24+int32(208)))) = v360
	goto L85
L89:
	;
	F_ConditionVariableSleep(m, v116, int32(134217777))
	mBase = m.M
	v472 = m.ExcPending
	if v472 != 0 {
		goto L3
	} else {
		goto L110
	}
L90:
	;
	v469 = v323
	goto L89
L91:
	;
	v468 = F_pgmem_kill(m, v323, int32(15))
	mBase = m.M
	goto L90
L92:
	;
	F_ConditionVariablePrepareToSleep(m, v116)
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L3
	} else {
		goto L95
	}
L93:
	;
	goto L94
L94:
	;
	v403 = *(*int32)(unsafe.Add(mBase, _c_F_InvalidateObsoleteReplicationSlots[2]))
	F_LWLockRelease(m, v403+int32(_a_F_InvalidateObsoleteReplicationSlots_0))
	mBase = m.M
	v407 = m.ExcPending
	if v407 != 0 {
		goto L3
	} else {
		goto L101
	}
L95:
	;
	v366 = *(*int32)(unsafe.Add(mBase, _c_F_InvalidateObsoleteReplicationSlots[2]))
	F_LWLockRelease(m, v366+int32(_a_F_InvalidateObsoleteReplicationSlots_0))
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
		goto L3
	} else {
		goto L96
	}
L96:
	;
	if v323 == v130 {
		v469 = v130
		goto L89
	} else {
		goto L97
	}
L97:
	;
	v372 = *(*int64)(unsafe.Add(mBase, uint32(v24)+144))
	*(*int64)(unsafe.Add(mBase, uint32(v24)+72)) = v372
	v374 = *(*int64)(unsafe.Add(mBase, uint32(v24)+152))
	*(*int64)(unsafe.Add(mBase, uint32(v24)+80)) = v374
	v376 = *(*int64)(unsafe.Add(mBase, uint32(v24)+160))
	*(*int64)(unsafe.Add(mBase, uint32(v24)+88)) = v376
	v378 = *(*int64)(unsafe.Add(mBase, uint32(v24)+168))
	*(*int64)(unsafe.Add(mBase, uint32(v24)+96)) = v378
	v380 = *(*int64)(unsafe.Add(mBase, uint32(v24)+176))
	*(*int64)(unsafe.Add(mBase, uint32(v24)+104)) = v380
	v382 = *(*int64)(unsafe.Add(mBase, uint32(v24)+184))
	*(*int64)(unsafe.Add(mBase, uint32(v24)+112)) = v382
	v384 = *(*int64)(unsafe.Add(mBase, uint32(v24)+192))
	*(*int64)(unsafe.Add(mBase, uint32(v24)+120)) = v384
	v386 = *(*int64)(unsafe.Add(mBase, uint32(v24)+200))
	*(*int64)(unsafe.Add(mBase, uint32(v24)+128)) = v386
	v391 = *(*int32)(unsafe.Add(mBase, uint32(v24)+140))
	F_ReportSlotInvalidation(m, v303, int32(1), v323, v24+int32(72), v174, v45, l3, v391)
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L3
	} else {
		goto L98
	}
L98:
	;
	v395 = *(*int32)(unsafe.Add(mBase, _c_F_InvalidateObsoleteReplicationSlots[10]))
	if v395 != int32(13) {
		goto L91
	} else {
		goto L99
	}
L99:
	;
	v400 = F_SendProcSignal(m, v323, int32(11), int32(-1))
	mBase = m.M
	v401 = m.ExcPending
	if v401 != 0 {
		goto L3
	} else {
		goto L100
	}
L100:
	;
	goto L90
L101:
	;
	v409 = *(*int32)(unsafe.Add(mBase, _c_F_InvalidateObsoleteReplicationSlots[8]))
	v412 = base.AtomicRmwXchg32(m, v409, int32(0), int32(1))
	if v412 != 0 {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	F_s_lock(m, v409, int32(_a_F_InvalidateObsoleteReplicationSlots_1), int32(1107), int32(_a_F_InvalidateObsoleteReplicationSlots_3))
	mBase = m.M
	v417 = m.ExcPending
	if v417 != 0 {
		goto L3
	} else {
		goto L105
	}
L103:
	;
	goto L104
L104:
	;
	v418 = int32(_a_F_InvalidateObsoleteReplicationSlots_4)
	v419 = *(*int32)(unsafe.Add(mBase, _c_F_InvalidateObsoleteReplicationSlots[8]))
	v420 = int32(257)
	*(*uint16)(unsafe.Add(mBase, uint32(v419)+12)) = uint16(v420)
	v422 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v409))), uint32(v422))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+64)) = int32(_a_F_InvalidateObsoleteReplicationSlots_5)
	v428 = *(*int32)(unsafe.Add(mBase, _c_F_InvalidateObsoleteReplicationSlots[8]))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+68)) = v428 + int32(24)
	v433 = v24 + int32(208)
	v437 = F_pg_sprintf(m, v433, int32(_a_F_InvalidateObsoleteReplicationSlots_6), v24-int32(-64))
	mBase = m.M
	v438 = m.ExcPending
	if v438 != 0 {
		goto L3
	} else {
		goto L106
	}
L105:
	;
	goto L104
L106:
	;
	v440 = *(*int32)(unsafe.Add(mBase, _c_F_InvalidateObsoleteReplicationSlots[8]))
	F_SaveSlotToPath(m, v440, v433, int32(21))
	mBase = m.M
	v443 = m.ExcPending
	if v443 != 0 {
		goto L3
	} else {
		goto L107
	}
L107:
	;
	F_ReplicationSlotRelease(m)
	mBase = m.M
	v445 = m.ExcPending
	if v445 != 0 {
		goto L3
	} else {
		goto L108
	}
L108:
	;
	v446 = *(*int64)(unsafe.Add(mBase, uint32(v24)+144))
	*(*int64)(unsafe.Add(mBase, uint32(v24))) = v446
	v448 = *(*int64)(unsafe.Add(mBase, uint32(v24)+152))
	*(*int64)(unsafe.Add(mBase, uint32(v24)+8)) = v448
	v450 = *(*int64)(unsafe.Add(mBase, uint32(v24)+160))
	*(*int64)(unsafe.Add(mBase, uint32(v24)+16)) = v450
	v452 = *(*int64)(unsafe.Add(mBase, uint32(v24)+168))
	*(*int64)(unsafe.Add(mBase, uint32(v24)+24)) = v452
	v454 = *(*int64)(unsafe.Add(mBase, uint32(v24)+176))
	*(*int64)(unsafe.Add(mBase, uint32(v24)+32)) = v454
	v456 = *(*int64)(unsafe.Add(mBase, uint32(v24)+184))
	*(*int64)(unsafe.Add(mBase, uint32(v24)+40)) = v456
	v458 = *(*int64)(unsafe.Add(mBase, uint32(v24)+192))
	*(*int64)(unsafe.Add(mBase, uint32(v24)+48)) = v458
	v460 = *(*int64)(unsafe.Add(mBase, uint32(v24)+200))
	*(*int64)(unsafe.Add(mBase, uint32(v24)+56)) = v460
	v462 = int32(0)
	v464 = *(*int32)(unsafe.Add(mBase, uint32(v24)+140))
	F_ReportSlotInvalidation(m, v303, v462, v462, v24, v174, v45, l3, v464)
	mBase = m.M
	v466 = m.ExcPending
	if v466 != 0 {
		goto L3
	} else {
		goto L109
	}
L109:
	;
	v494 = v337
	goto L20
L110:
	;
	v473 = int32(1)
	v475 = *(*int32)(unsafe.Add(mBase, _c_F_InvalidateObsoleteReplicationSlots[2]))
	v479 = F_LWLockAcquire(m, v475+int32(_a_F_InvalidateObsoleteReplicationSlots_0), v473)
	mBase = m.M
	v480 = m.ExcPending
	if v480 != 0 {
		goto L3
	} else {
		goto L111
	}
L111:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+140)) = int32(0)
	v483 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101)+4)))
	if v483 != 0 {
		v128 = v473
		v130 = v469
		v133 = v337
		v141 = v305
		goto L21
	} else {
		goto L112
	}
L112:
	;
	goto L22
L113:
	;
	v494 = v337
	goto L20
L114:
	;
	v505 = *(*int32)(unsafe.Add(mBase, _c_F_InvalidateObsoleteReplicationSlots[0]))
	if int32(0) < v505 {
		v60 = v505
		v65 = v494
		goto L8
	} else {
		goto L115
	}
L115:
	;
	v543 = v494
	goto L12
L116:
	;
	v543 = v519
	goto L12
L117:
	;
	if v543 == int32(0) {
		v582 = int32(0)
		goto L1
	} else {
		goto L118
	}
L118:
	;
	F_ReplicationSlotsComputeRequiredXmin(m, int32(0))
	mBase = m.M
	v564 = m.ExcPending
	if v564 != 0 {
		goto L3
	} else {
		goto L119
	}
L119:
	;
	F_ReplicationSlotsComputeRequiredLSN(m)
	mBase = m.M
	v566 = m.ExcPending
	if v566 != 0 {
		goto L3
	} else {
		goto L120
	}
L120:
	;
	v582 = int32(1)
	goto L1
L121:
	;
	v582 = v5
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
