package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_AbortTransaction(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v180 int32
	_ = v180
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v242 int32
	_ = v242
	var v254 int32
	_ = v254
	var v260 int32
	_ = v260
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v296 int64
	_ = v296
	var v297 int64
	_ = v297
	var v300 int64
	_ = v300
	var v301 int64
	_ = v301
	var v302 int64
	_ = v302
	var v303 int64
	_ = v303
	var v304 int64
	_ = v304
	var v305 int64
	_ = v305
	var v306 int64
	_ = v306
	var v307 int64
	_ = v307
	var v308 int64
	_ = v308
	var v309 int64
	_ = v309
	var v310 int64
	_ = v310
	var v311 int64
	_ = v311
	var v312 int64
	_ = v312
	var v313 int64
	_ = v313
	var v314 int64
	_ = v314
	var v315 int64
	_ = v315
	var v316 int64
	_ = v316
	var v317 int64
	_ = v317
	var v318 int64
	_ = v318
	var v319 int64
	_ = v319
	var v320 int64
	_ = v320
	var v321 int64
	_ = v321
	var v322 int64
	_ = v322
	var v323 int64
	_ = v323
	var v324 int64
	_ = v324
	var v325 int64
	_ = v325
	var v326 int64
	_ = v326
	var v327 int64
	_ = v327
	var v328 int64
	_ = v328
	var v329 int64
	_ = v329
	var v330 int64
	_ = v330
	var v362 int64
	_ = v362
	var v366 int32
	_ = v366
	var v369 int32
	_ = v369
	var v379 int32
	_ = v379
	var v381 int32
	_ = v381
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v388 int64
	_ = v388
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v393 int32
	_ = v393
	var v395 int32
	_ = v395
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
	var v404 int32
	_ = v404
	var v409 int32
	_ = v409
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v414 int32
	_ = v414
	var v417 int32
	_ = v417
	var v422 int32
	_ = v422
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v427 int32
	_ = v427
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v439 int32
	_ = v439
	var v441 int32
	_ = v441
	var v444 int32
	_ = v444
	var v446 int32
	_ = v446
	var v449 int32
	_ = v449
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v457 int32
	_ = v457
	var v460 int32
	_ = v460
	var v462 int32
	_ = v462
	var v469 int32
	_ = v469
	var v481 int32
	_ = v481
	var v486 int32
	_ = v486
	var v488 int32
	_ = v488
	var v493 int32
	_ = v493
	var v496 int32
	_ = v496
	var v500 int32
	_ = v500
	var v503 int32
	_ = v503
	var v505 int32
	_ = v505
	var v512 int32
	_ = v512
	var v515 int32
	_ = v515
	var v517 int32
	_ = v517
	var v519 int32
	_ = v519
	var v522 int32
	_ = v522
	var v524 int32
	_ = v524
	var v537 int32
	_ = v537
	var v539 int32
	_ = v539
	var v542 int32
	_ = v542
	var v561 int32
	_ = v561
	var v565 int32
	_ = v565
	var v567 int32
	_ = v567
	var v571 int32
	_ = v571
	var v575 int32
	_ = v575
	var v578 int32
	_ = v578
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v584 int32
	_ = v584
	var v588 int32
	_ = v588
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v603 int32
	_ = v603
	var v613 int32
	_ = v613
	var v615 int32
	_ = v615
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = int32(_a_F_AbortTransaction_0)
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_AbortTransaction[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_AbortTransaction[0])) = v12 + int32(1)
	v17 = *(*int32)(unsafe.Add(mBase, _c_F_AbortTransaction[1]))
	v19 = *(*int32)(unsafe.Add(mBase, _c_F_AbortTransaction[2]))
	if int32(0) < v19 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_disable_timeout(m, int32(8))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v27 = *(*int32)(unsafe.Add(mBase, _c_F_AbortTransaction[3]))
	*(*int32)(unsafe.Add(mBase, _c_F_AbortTransaction[4])) = v27
	v31 = *(*int32)(unsafe.Add(mBase, _c_F_AbortTransaction[5]))
	v33 = *(*int32)(unsafe.Add(mBase, _c_F_AbortTransaction[6]))
	if v31 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	return
L5:
	;
	goto L3
L6:
	;
	v34 = v31
	goto L8
L7:
	;
	v34 = v33
	goto L8
L8:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_AbortTransaction[7])) = v34
	F_LWLockReleaseAll(m)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L4
	} else {
		goto L9
	}
L9:
	;
	F_WaitLSNCleanup(m)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L4
	} else {
		goto L10
	}
L10:
	;
	v41 = *(*int32)(unsafe.Add(mBase, _c_F_AbortTransaction[8]))
	v42 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v41))) = v42
	v46 = *(*int32)(unsafe.Add(mBase, _c_F_AbortTransaction[9]))
	if v46 == v42 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	F_pgaio_error_cleanup(m)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L4
	} else {
		goto L16
	}
L12:
	;
	goto L11
L13:
	;
	v50 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_AbortTransaction[10])))
	if v50&int32(1) == int32(0) {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v46)+220))
	if v55 == int32(0) {
		goto L12
	} else {
		goto L15
	}
L15:
	;
	v58 = int32(_a_F_AbortTransaction_1)
	v60 = *(*int32)(unsafe.Add(mBase, _c_F_AbortTransaction[11]))
	v61 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_AbortTransaction[11])) = v60 + v61
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
	*(*int32)(unsafe.Add(mBase, uint32(v46))) = v64 + v61
	v68 = int32(0)
	v70 = int32(_a_F_AbortTransaction_2)
	v71 = base.AtomicRmwOr32(m, v68, v70, v68)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+220)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v46)+224)) = v68
	v79 = base.AtomicRmwOr32(m, v68, v70, v68)
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
	*(*int32)(unsafe.Add(mBase, uint32(v46))) = v80 + v61
	v86 = *(*int32)(unsafe.Add(mBase, _c_F_AbortTransaction[11]))
	*(*int32)(unsafe.Add(mBase, _c_F_AbortTransaction[11])) = v86 - v61
	goto L12
L16:
	;
	F_UnlockBuffers(m)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L4
	} else {
		goto L17
	}
L17:
	;
	v94 = int32(0)
	v101 = *(*int32)(unsafe.Add(mBase, _c_F_AbortTransaction[12]))
	if v101 <= v94 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	F_ConditionVariableCancelSleep(m)
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L4
	} else {
		goto L31
	}
L19:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_AbortTransaction[13])) = int64(0)
	*(*int32)(unsafe.Add(mBase, _c_F_AbortTransaction[14])) = int32(_a_F_AbortTransaction_3)
	v180 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_AbortTransaction[12])) = v180
	*(*int32)(unsafe.Add(mBase, _c_F_AbortTransaction[15])) = v180
	*(*uint8)(unsafe.Add(mBase, _c_F_AbortTransaction[16])) = uint8(v180)
	*(*uint8)(unsafe.Add(mBase, _c_F_AbortTransaction[17])) = uint8(v180)
	goto L18
L20:
	;
	v105 = v101 & int32(7)
	v107 = *(*int32)(unsafe.Add(mBase, _c_F_AbortTransaction[18]))
	if base.Ui32(int32(8)) <= base.Ui32(v101) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v113 = v94
	v117 = v94
	goto L24
L22:
	;
	v145 = v94
	goto L23
L23:
	;
	v151 = int32(0)
	v152 = v145
	goto L28
L24:
	;
	v120 = v107 + v113*int32(_a_F_AbortTransaction_4)
	v121 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v120)+uint32(_c_F_AbortTransaction[19]))) = uint8(v121)
	*(*uint8)(unsafe.Add(mBase, uint32(v120)+uint32(_c_F_AbortTransaction[20]))) = uint8(v121)
	*(*uint8)(unsafe.Add(mBase, uint32(v120)+uint32(_c_F_AbortTransaction[21]))) = uint8(v121)
	*(*uint8)(unsafe.Add(mBase, uint32(v120)+uint32(_c_F_AbortTransaction[22]))) = uint8(v121)
	*(*uint8)(unsafe.Add(mBase, uint32(v120)+uint32(_c_F_AbortTransaction[23]))) = uint8(v121)
	*(*uint8)(unsafe.Add(mBase, uint32(v120)+uint32(_c_F_AbortTransaction[24]))) = uint8(v121)
	*(*uint8)(unsafe.Add(mBase, uint32(v120)+uint32(_c_F_AbortTransaction[25]))) = uint8(v121)
	*(*uint8)(unsafe.Add(mBase, uint32(v120))) = uint8(v121)
	v137 = int32(8)
	v138 = v113 + v137
	v140 = v117 + v137
	if v140 != v101&int32(2147483640) {
		v113 = v138
		v117 = v140
		goto L24
	} else {
		goto L26
	}
L25:
	;
	if v105 == int32(0) {
		goto L19
	} else {
		goto L27
	}
L26:
	;
	goto L25
L27:
	;
	v145 = v138
	goto L23
L28:
	;
	v160 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v107+v152*int32(_a_F_AbortTransaction_4)))) = uint8(v160)
	v162 = int32(1)
	v165 = v151 + v162
	if v165 != v105 {
		v151 = v165
		v152 = v152 + v162
		goto L28
	} else {
		goto L30
	}
L29:
	;
	goto L19
L30:
	;
	goto L29
L31:
	;
	F_LockErrorCleanup(m)
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L4
	} else {
		goto L32
	}
L32:
	;
	F_reschedule_timeouts(m)
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L4
	} else {
		goto L33
	}
L33:
	;
	F_pgmem_sigprocmask(m, int32(_a_F_AbortTransaction_5), int32(0))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L4
	} else {
		goto L34
	}
L34:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v17)+24))
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v17)+20))
	switch v202 - int32(2) {
	case 0, 3:
		goto L35
	default:
		goto L36
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+20)) = int32(4)
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v17)+60))
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v17)+64))
	*(*int32)(unsafe.Add(mBase, _c_F_AbortTransaction[26])) = v232
	*(*int32)(unsafe.Add(mBase, _c_F_AbortTransaction[27])) = v231
	goto L44
L36:
	;
	v207 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L4
	} else {
		goto L37
	}
L37:
	;
	if v207 == int32(0) {
		goto L35
	} else {
		goto L38
	}
L38:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v17)+20))
	if base.Ui32(v211) <= base.Ui32(int32(5)) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v211<<(uint(int32(2))%32))+uint32(_c_F_AbortTransaction[28])))
	v218 = v216
	goto L41
L40:
	;
	v218 = int32(_a_F_AbortTransaction_6)
	goto L41
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v218
	F_errmsg_internal(m, int32(_a_F_AbortTransaction_7), v8)
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L4
	} else {
		goto L42
	}
L42:
	;
	F_errfinish(m, int32(_a_F_AbortTransaction_8), int32(2928), int32(_a_F_AbortTransaction_9))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L4
	} else {
		goto L43
	}
L43:
	;
	goto L35
L44:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v17)+28))
	v239 = *(*int32)(unsafe.Add(mBase, _c_F_AbortTransaction[29]))
	if v237 <= v239 {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	v254 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_AbortTransaction[30])) = uint8(v254)
	*(*int32)(unsafe.Add(mBase, _c_F_AbortTransaction[31])) = v254
	goto L49
L46:
	;
	v242 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_AbortTransaction[32])) = v242
	*(*int32)(unsafe.Add(mBase, _c_F_AbortTransaction[33])) = v242
	*(*int32)(unsafe.Add(mBase, _c_F_AbortTransaction[34])) = v242
	*(*int32)(unsafe.Add(mBase, _c_F_AbortTransaction[29])) = v242
	goto L48
L47:
	;
	goto L48
L48:
	;
	goto L45
L49:
	;
	v260 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_AbortTransaction[35])) = uint8(v260)
	*(*int32)(unsafe.Add(mBase, _c_F_AbortTransaction[36])) = v260
	F_AtEOXact_Parallel(m, v260)
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L4
	} else {
		goto L50
	}
L50:
	;
	v269 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+76)) = uint8(v269)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+72)) = v269
	F_AfterTriggerEndXact(m)
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L4
	} else {
		goto L51
	}
L51:
	;
	F_AtAbort_Portals(m)
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L4
	} else {
		goto L52
	}
L52:
	;
	v279 = base.B2i32(v201 == int32(5))
	F_smgrDoPendingSyncs(m, int32(0), v279)
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L4
	} else {
		goto L53
	}
L53:
	;
	F_AtEOXact_LargeObject(m, int32(0))
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L4
	} else {
		goto L54
	}
L54:
	;
	F_ApplyPendingListenActions(m, int32(0))
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L4
	} else {
		goto L55
	}
L55:
	;
	v289 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_AbortTransaction[37])))
	if v289 == int32(0) {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v369 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_AbortTransaction[38])) = v369
	*(*int32)(unsafe.Add(mBase, _c_F_AbortTransaction[39])) = v369
	*(*int32)(unsafe.Add(mBase, _c_F_AbortTransaction[40])) = v369
	F_AtEOXact_RelationMap(m, v369, v279)
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L4
	} else {
		goto L67
	}
L57:
	;
	v293 = *(*int32)(unsafe.Add(mBase, _c_F_AbortTransaction[41]))
	if v293 != 0 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v293)))
	v296 = *(*int64)(unsafe.Add(mBase, uint32(v295)+8))
	v297 = *(*int64)(unsafe.Add(mBase, uint32(v295)+808))
	if v297 != int64(0) {
		goto L62
	} else {
		goto L63
	}
L59:
	;
	goto L60
L60:
	;
	F_asyncQueueUnregister(m)
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L4
	} else {
		goto L66
	}
L61:
	;
	if v362 != int64(0) {
		goto L56
	} else {
		goto L65
	}
L62:
	;
	v300 = *(*int64)(unsafe.Add(mBase, uint32(v295)+752))
	v301 = *(*int64)(unsafe.Add(mBase, uint32(v295)+728))
	v302 = *(*int64)(unsafe.Add(mBase, uint32(v295)+704))
	v303 = *(*int64)(unsafe.Add(mBase, uint32(v295)+680))
	v304 = *(*int64)(unsafe.Add(mBase, uint32(v295)+656))
	v305 = *(*int64)(unsafe.Add(mBase, uint32(v295)+632))
	v306 = *(*int64)(unsafe.Add(mBase, uint32(v295)+608))
	v307 = *(*int64)(unsafe.Add(mBase, uint32(v295)+584))
	v308 = *(*int64)(unsafe.Add(mBase, uint32(v295)+560))
	v309 = *(*int64)(unsafe.Add(mBase, uint32(v295)+536))
	v310 = *(*int64)(unsafe.Add(mBase, uint32(v295)+512))
	v311 = *(*int64)(unsafe.Add(mBase, uint32(v295)+488))
	v312 = *(*int64)(unsafe.Add(mBase, uint32(v295)+464))
	v313 = *(*int64)(unsafe.Add(mBase, uint32(v295)+440))
	v314 = *(*int64)(unsafe.Add(mBase, uint32(v295)+416))
	v315 = *(*int64)(unsafe.Add(mBase, uint32(v295)+392))
	v316 = *(*int64)(unsafe.Add(mBase, uint32(v295)+368))
	v317 = *(*int64)(unsafe.Add(mBase, uint32(v295)+344))
	v318 = *(*int64)(unsafe.Add(mBase, uint32(v295)+320))
	v319 = *(*int64)(unsafe.Add(mBase, uint32(v295)+296))
	v320 = *(*int64)(unsafe.Add(mBase, uint32(v295)+272))
	v321 = *(*int64)(unsafe.Add(mBase, uint32(v295)+248))
	v322 = *(*int64)(unsafe.Add(mBase, uint32(v295)+224))
	v323 = *(*int64)(unsafe.Add(mBase, uint32(v295)+200))
	v324 = *(*int64)(unsafe.Add(mBase, uint32(v295)+176))
	v325 = *(*int64)(unsafe.Add(mBase, uint32(v295)+152))
	v326 = *(*int64)(unsafe.Add(mBase, uint32(v295)+128))
	v327 = *(*int64)(unsafe.Add(mBase, uint32(v295)+104))
	v328 = *(*int64)(unsafe.Add(mBase, uint32(v295)+80))
	v329 = *(*int64)(unsafe.Add(mBase, uint32(v295)+56))
	v330 = *(*int64)(unsafe.Add(mBase, uint32(v295)+32))
	v362 = v300 + (v301 + (v302 + (v303 + (v304 + (v305 + (v306 + (v307 + (v308 + (v309 + (v310 + (v311 + (v312 + (v313 + (v314 + (v315 + (v316 + (v317 + (v318 + (v319 + (v320 + (v321 + (v322 + (v323 + (v324 + (v325 + (v326 + (v327 + (v328 + (v329 + (v330 + v296))))))))))))))))))))))))))))))
	goto L64
L63:
	;
	v362 = v296
	goto L64
L64:
	;
	goto L61
L65:
	;
	goto L60
L66:
	;
	goto L56
L67:
	;
	F_AtAbort_Twophase(m)
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L4
	} else {
		goto L68
	}
L68:
	;
	if v279 == int32(0) {
		goto L70
	} else {
		goto L71
	}
L69:
	;
	v393 = *(*int32)(unsafe.Add(mBase, _c_F_AbortTransaction[42]))
	F_ProcArrayEndTransaction(m, v393, v391)
	mBase = m.M
	v395 = m.ExcPending
	if v395 != 0 {
		goto L4
	} else {
		goto L75
	}
L70:
	;
	v385 = F_RecordTransactionAbort(m, int32(0))
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L4
	} else {
		goto L73
	}
L71:
	;
	goto L72
L72:
	;
	v388 = *(*int64)(unsafe.Add(mBase, _c_F_AbortTransaction[43]))
	F_XLogSetAsyncXactLSN(m, v388)
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L4
	} else {
		goto L74
	}
L73:
	;
	v391 = v385
	goto L69
L74:
	;
	v391 = v260
	goto L69
L75:
	;
	v397 = *(*int32)(unsafe.Add(mBase, _c_F_AbortTransaction[3]))
	if v397 != 0 {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v399 = *(*int32)(unsafe.Add(mBase, _c_F_AbortTransaction[44]))
	if v201 == int32(5) {
		goto L80
	} else {
		goto L81
	}
L77:
	;
	goto L78
L78:
	;
	v613 = int32(_a_F_AbortTransaction_0)
	v615 = *(*int32)(unsafe.Add(mBase, _c_F_AbortTransaction[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_AbortTransaction[0])) = v615 - int32(1)
	m.G0 = v8 + int32(16)
	return
L79:
	;
	v434 = *(*int32)(unsafe.Add(mBase, _c_F_AbortTransaction[3]))
	v435 = int32(1)
	F_ResourceOwnerReleaseInternal(m, v434, v435, int32(0), v435)
	mBase = m.M
	v439 = m.ExcPending
	if v439 != 0 {
		goto L4
	} else {
		goto L93
	}
L80:
	;
	if v399 == int32(0) {
		goto L79
	} else {
		goto L83
	}
L81:
	;
	goto L82
L82:
	;
	if v399 == int32(0) {
		goto L79
	} else {
		goto L88
	}
L83:
	;
	v404 = v399
	goto L84
L84:
	;
	v409 = *(*int32)(unsafe.Add(mBase, uint32(v404)))
	v411 = *(*int32)(unsafe.Add(mBase, uint32(v404)+8))
	v412 = *(*int32)(unsafe.Add(mBase, uint32(v404)+4))
	m.T0[v412].(func(*base.Module, int32, int32))(m, int32(3), v411)
	mBase = m.M
	v414 = m.ExcPending
	if v414 != 0 {
		goto L4
	} else {
		goto L86
	}
L85:
	;
	goto L79
L86:
	;
	if v409 != 0 {
		v404 = v409
		goto L84
	} else {
		goto L87
	}
L87:
	;
	goto L85
L88:
	;
	v417 = v399
	goto L89
L89:
	;
	v422 = *(*int32)(unsafe.Add(mBase, uint32(v417)))
	v424 = *(*int32)(unsafe.Add(mBase, uint32(v417)+8))
	v425 = *(*int32)(unsafe.Add(mBase, uint32(v417)+4))
	m.T0[v425].(func(*base.Module, int32, int32))(m, int32(2), v424)
	mBase = m.M
	v427 = m.ExcPending
	if v427 != 0 {
		goto L4
	} else {
		goto L91
	}
L90:
	;
	goto L79
L91:
	;
	if v422 != 0 {
		v417 = v422
		goto L89
	} else {
		goto L92
	}
L92:
	;
	goto L90
L93:
	;
	F_AtEOXact_Aio(m)
	mBase = m.M
	v441 = m.ExcPending
	if v441 != 0 {
		goto L4
	} else {
		goto L94
	}
L94:
	;
	F_AtEOXact_RelationCache(m, int32(0))
	mBase = m.M
	v444 = m.ExcPending
	if v444 != 0 {
		goto L4
	} else {
		goto L95
	}
L95:
	;
	F_AtEOXact_TypeCache(m)
	mBase = m.M
	v446 = m.ExcPending
	if v446 != 0 {
		goto L4
	} else {
		goto L96
	}
L96:
	;
	F_AtEOXact_Inval(m, int32(0))
	mBase = m.M
	v449 = m.ExcPending
	if v449 != 0 {
		goto L4
	} else {
		goto L97
	}
L97:
	;
	v451 = *(*int32)(unsafe.Add(mBase, _c_F_AbortTransaction[45]))
	v452 = int32(_a_F_AbortTransaction_10)
	v453 = *(*int32)(unsafe.Add(mBase, _c_F_AbortTransaction[46]))
	v454 = int32(2)
	v457 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v451+v453<<(uint(v454)%32)))) = v457
	v460 = *(*int32)(unsafe.Add(mBase, _c_F_AbortTransaction[47]))
	v462 = *(*int32)(unsafe.Add(mBase, _c_F_AbortTransaction[46]))
	*(*int32)(unsafe.Add(mBase, uint32(v460+v462<<(uint(v454)%32)))) = v457
	v469 = int32(_a_F_AbortTransaction_11)
	*(*int32)(unsafe.Add(mBase, _c_F_AbortTransaction[48])) = v469
	*(*int32)(unsafe.Add(mBase, _c_F_AbortTransaction[49])) = v469
	*(*int32)(unsafe.Add(mBase, _c_F_AbortTransaction[50])) = v457
	*(*int32)(unsafe.Add(mBase, _c_F_AbortTransaction[51])) = v457
	goto L98
L98:
	;
	v481 = *(*int32)(unsafe.Add(mBase, _c_F_AbortTransaction[3]))
	F_ResourceOwnerReleaseInternal(m, v481, int32(2), int32(0), int32(1))
	mBase = m.M
	v486 = m.ExcPending
	if v486 != 0 {
		goto L4
	} else {
		goto L99
	}
L99:
	;
	v488 = *(*int32)(unsafe.Add(mBase, _c_F_AbortTransaction[3]))
	F_ResourceOwnerReleaseInternal(m, v488, int32(3), int32(0), int32(1))
	mBase = m.M
	v493 = m.ExcPending
	if v493 != 0 {
		goto L4
	} else {
		goto L100
	}
L100:
	;
	F_smgrDoPendingDeletes(m, int32(0))
	mBase = m.M
	v496 = m.ExcPending
	if v496 != 0 {
		goto L4
	} else {
		goto L101
	}
L101:
	;
	F_AtEOXact_GUC(m, int32(0), int32(1))
	mBase = m.M
	v500 = m.ExcPending
	if v500 != 0 {
		goto L4
	} else {
		goto L102
	}
L102:
	;
	F_AtEOXact_SPI(m, int32(0))
	mBase = m.M
	v503 = m.ExcPending
	if v503 != 0 {
		goto L4
	} else {
		goto L103
	}
L103:
	;
	v505 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_AbortTransaction[52])) = v505
	*(*int32)(unsafe.Add(mBase, _c_F_AbortTransaction[53])) = v505
	goto L104
L104:
	;
	F_AtEOXact_on_commit_actions(m, int32(0))
	mBase = m.M
	v512 = m.ExcPending
	if v512 != 0 {
		goto L4
	} else {
		goto L105
	}
L105:
	;
	v515 = base.B2i32(v201 == int32(5))
	F_AtEOXact_Namespace(m, int32(0), v515)
	mBase = m.M
	v517 = m.ExcPending
	if v517 != 0 {
		goto L4
	} else {
		goto L106
	}
L106:
	;
	F_smgrdestroyall(m)
	mBase = m.M
	v519 = m.ExcPending
	if v519 != 0 {
		goto L4
	} else {
		goto L107
	}
L107:
	;
	F_AtEOXact_Files(m, int32(0))
	mBase = m.M
	v522 = m.ExcPending
	if v522 != 0 {
		goto L4
	} else {
		goto L108
	}
L108:
	;
	v524 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_AbortTransaction[54])) = v524
	*(*int32)(unsafe.Add(mBase, _c_F_AbortTransaction[55])) = v524
	*(*int32)(unsafe.Add(mBase, _c_F_AbortTransaction[56])) = v524
	*(*int32)(unsafe.Add(mBase, _c_F_AbortTransaction[57])) = v524
	goto L109
L109:
	;
	F_AtEOXact_HashTables(m, int32(0))
	mBase = m.M
	v537 = m.ExcPending
	if v537 != 0 {
		goto L4
	} else {
		goto L110
	}
L110:
	;
	F_AtEOXact_RI(m)
	mBase = m.M
	v539 = m.ExcPending
	if v539 != 0 {
		goto L4
	} else {
		goto L111
	}
L111:
	;
	F_AtEOXact_PgStat(m, int32(0), v515)
	mBase = m.M
	v542 = m.ExcPending
	if v542 != 0 {
		goto L4
	} else {
		goto L112
	}
L112:
	;
	goto L114
L113:
	;
	F_AtEOXact_LogicalRepWorkers(m, int32(0))
	mBase = m.M
	v565 = m.ExcPending
	if v565 != 0 {
		goto L4
	} else {
		goto L118
	}
L114:
	;
	v561 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_AbortTransaction[58])) = uint8(v561)
	goto L113
L118:
	;
	F_AtEOXact_LogicalCtl(m)
	mBase = m.M
	v567 = m.ExcPending
	if v567 != 0 {
		goto L4
	} else {
		goto L119
	}
L119:
	;
	v571 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_AbortTransaction[10])))
	if v571 != int32(1) {
		goto L121
	} else {
		goto L122
	}
L120:
	;
	goto L78
L121:
	;
	goto L120
L122:
	;
	v575 = *(*int32)(unsafe.Add(mBase, _c_F_AbortTransaction[9]))
	if v575 == int32(0) {
		goto L121
	} else {
		goto L123
	}
L123:
	;
	v578 = int32(_a_F_AbortTransaction_1)
	v580 = *(*int32)(unsafe.Add(mBase, _c_F_AbortTransaction[11]))
	v581 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_AbortTransaction[11])) = v580 + v581
	v584 = *(*int32)(unsafe.Add(mBase, uint32(v575)))
	*(*int32)(unsafe.Add(mBase, uint32(v575))) = v584 + v581
	v588 = int32(0)
	v590 = int32(_a_F_AbortTransaction_12)
	v591 = base.AtomicRmwOr32(m, v588, v590, v588)
	*(*int64)(unsafe.Add(mBase, uint32(v575)+24)) = int64(0)
	v596 = base.AtomicRmwOr32(m, v588, v590, v588)
	v597 = *(*int32)(unsafe.Add(mBase, uint32(v575)))
	*(*int32)(unsafe.Add(mBase, uint32(v575))) = v597 + v581
	v603 = *(*int32)(unsafe.Add(mBase, _c_F_AbortTransaction[11]))
	*(*int32)(unsafe.Add(mBase, _c_F_AbortTransaction[11])) = v603 - v581
	goto L121
}
func F_AssignTransactionId(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v25 int64
	_ = v25
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v122 int32
	_ = v122
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v138 int64
	_ = v138
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v197 int64
	_ = v197
	var v201 int32
	_ = v201
	var v209 int64
	_ = v209
	var v214 int32
	_ = v214
	var v216 int64
	_ = v216
	var v218 int32
	_ = v218
	var v219 int64
	_ = v219
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v239 int64
	_ = v239
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	var v268 int32
	_ = v268
	v2 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v14 = *(*int32)(unsafe.Add(mBase, _c_F_AssignTransactionId[0]))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+72))
	if v15 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L9
	} else {
		goto L50
	}
L2:
	;
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+76)))
	if v16 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v18 = *(*int32)(unsafe.Add(mBase, _c_F_AssignTransactionId[1]))
	if int32(0) <= v18 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	if v21 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v160 = int32(_a_F_AssignTransactionId_0)
	v161 = *(*int32)(unsafe.Add(mBase, _c_F_AssignTransactionId[2]))
	v163 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	*(*int32)(unsafe.Add(mBase, _c_F_AssignTransactionId[2])) = v163
	v165 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v166 = m.G0
	v168 = v166 - int32(16)
	m.G0 = v168
	*(*int32)(unsafe.Add(mBase, uint32(v168)+12)) = int32(17104896)
	*(*int64)(unsafe.Add(mBase, uint32(v168)+4)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v168))) = v165
	v176 = int32(0)
	v178 = F_LockAcquire(m, v168, int32(7), v176, v176)
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L9
	} else {
		goto L37
	}
L6:
	;
	v25 = F_GetNewTransactionId(m, int32(0))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	if v70 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L9:
	;
	return
L10:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v25
	*(*int64)(unsafe.Add(mBase, _c_F_AssignTransactionId[3])) = v25
	v30 = m.G0
	v32 = v30 - int32(16)
	m.G0 = v32
	v35 = *(*int32)(unsafe.Add(mBase, _c_F_AssignTransactionId[4]))
	if v35 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v36 = base.I32_wrap_i64(v25)
	v38 = *(*int32)(unsafe.Add(mBase, _c_F_AssignTransactionId[5]))
	v42 = F_LWLockAcquire(m, v38+int32(3584), int32(0))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L9
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	m.G0 = v32 + int32(16)
	v155 = v2
	goto L5
L14:
	;
	v45 = *(*int32)(unsafe.Add(mBase, _c_F_AssignTransactionId[4]))
	*(*int32)(unsafe.Add(mBase, uint32(v45)+96)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v32)+12)) = v36
	v49 = *(*int32)(unsafe.Add(mBase, _c_F_AssignTransactionId[6]))
	v55 = F_hash_search(m, v49, v32+int32(12), int32(1), v32+int32(11))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L9
	} else {
		goto L15
	}
L15:
	;
	v58 = *(*int32)(unsafe.Add(mBase, _c_F_AssignTransactionId[4]))
	*(*int32)(unsafe.Add(mBase, uint32(v55)+4)) = v58
	v61 = *(*int32)(unsafe.Add(mBase, _c_F_AssignTransactionId[5]))
	F_LWLockRelease(m, v61+int32(3584))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L9
	} else {
		goto L16
	}
L16:
	;
	goto L13
L17:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v75 = F_palloc_mul(m, int32(4), v74)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L9
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v132 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_AssignTransactionId[7])))
	v134 = *(*int32)(unsafe.Add(mBase, _c_F_AssignTransactionId[8]))
	v136 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_AssignTransactionId[9])))
	v138 = F_GetNewTransactionId(m, int32(1))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L9
	} else {
		goto L35
	}
L20:
	;
	v78 = v2
	v80 = v21
	goto L21
L21:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
	if v85 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	if v95 != 0 {
		goto L27
	} else {
		goto L28
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v75+v78<<(uint(int32(2))%32)))) = v80
	v93 = v78 + int32(1)
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v80)+80))
	if v94 != 0 {
		v78 = v93
		v80 = v94
		goto L21
	} else {
		goto L26
	}
L24:
	;
	v95 = v78
	goto L25
L25:
	;
	goto L22
L26:
	;
	v95 = v93
	goto L25
L27:
	;
	v98 = v95
	goto L30
L28:
	;
	goto L29
L29:
	;
	F_pfree(m, v75)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L9
	} else {
		goto L34
	}
L30:
	;
	v106 = v98 - int32(1)
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v75+v106<<(uint(int32(2))%32))))
	F_AssignTransactionId(m, v110)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L9
	} else {
		goto L32
	}
L31:
	;
	goto L29
L32:
	;
	if v106 != 0 {
		v98 = v106
		goto L30
	} else {
		goto L33
	}
L33:
	;
	goto L31
L34:
	;
	goto L19
L35:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v138
	v142 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
	F_SubTransSetParent(m, base.I32_wrap_i64(v138), v143)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L9
	} else {
		goto L36
	}
L36:
	;
	v146 = int32(1)
	v155 = (v132 | base.B2i32(v146 < v134)) & (v136 ^ v146)
	goto L5
L37:
	;
	m.G0 = v168 + int32(16)
	*(*int32)(unsafe.Add(mBase, _c_F_AssignTransactionId[2])) = v161
	if v21 == int32(0) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	m.G0 = v11 + int32(16)
	return
L39:
	;
	v188 = *(*int32)(unsafe.Add(mBase, _c_F_AssignTransactionId[8]))
	if v188 <= int32(0) {
		goto L38
	} else {
		goto L40
	}
L40:
	;
	v191 = int32(_a_F_AssignTransactionId_1)
	v192 = *(*int32)(unsafe.Add(mBase, _c_F_AssignTransactionId[10]))
	v197 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*uint32)(unsafe.Add(mBase, uint32(v192<<(uint(int32(2))%32))+uint32(_c_F_AssignTransactionId[11]))) = uint32(v197)
	v201 = v192 + int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_AssignTransactionId[10])) = v201
	if (v155^int32(-1))&base.B2i32(v201 < int32(64)) != 0 {
		goto L38
	} else {
		goto L41
	}
L41:
	;
	v209 = int64(*(*uint32)(unsafe.Add(mBase, _c_F_AssignTransactionId[3])))
	if v209 == int64(0) {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	F_AssignTransactionId(m, int32(_a_F_AssignTransactionId_2))
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L9
	} else {
		goto L45
	}
L43:
	;
	v219 = v209
	v220 = v201
	goto L44
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v220
	*(*uint32)(unsafe.Add(mBase, uint32(v11)+8)) = uint32(v219)
	F_XLogBeginInsert(m)
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L9
	} else {
		goto L46
	}
L45:
	;
	v216 = *(*int64)(unsafe.Add(mBase, _c_F_AssignTransactionId[3]))
	v218 = *(*int32)(unsafe.Add(mBase, _c_F_AssignTransactionId[10]))
	v219 = v216
	v220 = v218
	goto L44
L46:
	;
	v225 = int32(8)
	F_XLogRegisterData(m, v11+v225, v225)
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L9
	} else {
		goto L47
	}
L47:
	;
	v232 = *(*int32)(unsafe.Add(mBase, _c_F_AssignTransactionId[10]))
	F_XLogRegisterData(m, int32(_a_F_AssignTransactionId_3), v232<<(uint(int32(2))%32))
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L9
	} else {
		goto L48
	}
L48:
	;
	v239 = F_XLogInsert(m, int32(1), int32(80))
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L9
	} else {
		goto L49
	}
L49:
	;
	v242 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_AssignTransactionId[9])) = uint8(v242)
	*(*int32)(unsafe.Add(mBase, _c_F_AssignTransactionId[10])) = int32(0)
	goto L38
L50:
	;
	F_errcode(m, int32(322))
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L9
	} else {
		goto L51
	}
L51:
	;
	F_errmsg(m, int32(_a_F_AssignTransactionId_4), int32(0))
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L9
	} else {
		goto L52
	}
L52:
	;
	F_errfinish(m, int32(_a_F_AssignTransactionId_5), int32(654), int32(_a_F_AssignTransactionId_6))
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L9
	} else {
		goto L53
	}
L53:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_CleanupTransaction(m *base.Module) {
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
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
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
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v117 int64
	_ = v117
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v156 int32
	_ = v156
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v10 = *(*int32)(unsafe.Add(mBase, _c_F_CleanupTransaction[0]))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+20))
	if v11 == int32(4) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v14 = m.G0
	v16 = v14 - int32(32)
	m.G0 = v16
	v19 = v16 + int32(12)
	v21 = *(*int32)(unsafe.Add(mBase, _c_F_CleanupTransaction[1]))
	F_hash_seq_init(m, v19, v21)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L4
	} else {
		goto L44
	}
L4:
	;
	return
L5:
	;
	v24 = F_hash_seq_search(m, v19)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	if v24 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v26 = v24
	goto L10
L8:
	;
	goto L9
L9:
	;
	m.G0 = v16 + int32(32)
	F_AtEOXact_Snapshot(m, int32(0), int32(1))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L4
	} else {
		goto L31
	}
L10:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v26)+64))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+80))
	if v31 == int32(3) {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	goto L9
L12:
	;
	v65 = F_hash_seq_search(m, v16+int32(12))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L4
	} else {
		goto L29
	}
L13:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v30)+20))
	if v34 == int32(0) {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+85)))
	if v37 != 0 {
		goto L12
	} else {
		goto L15
	}
L15:
	;
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+84)))
	if v38 == int32(1) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v41 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v30)+84)) = uint8(v41)
	goto L18
L17:
	;
	goto L18
L18:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v30)+16))
	if v43 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v46 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L4
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	F_PortalDrop(m, v30, int32(0))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L4
	} else {
		goto L28
	}
L22:
	;
	if v46 != 0 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v48
	F_errmsg_internal(m, int32(_a_F_CleanupTransaction_0), v16)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L4
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+16)) = int32(0)
	goto L21
L26:
	;
	F_errfinish(m, int32(_a_F_CleanupTransaction_1), int32(903), int32(_a_F_CleanupTransaction_2))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L4
	} else {
		goto L27
	}
L27:
	;
	goto L25
L28:
	;
	goto L12
L29:
	;
	if v65 != 0 {
		v26 = v65
		goto L10
	} else {
		goto L30
	}
L30:
	;
	goto L11
L31:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CleanupTransaction[2])) = int32(0)
	v82 = *(*int32)(unsafe.Add(mBase, _c_F_CleanupTransaction[3]))
	if v82 != 0 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	F_ResourceOwnerDelete(m, v82)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L4
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	v85 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+40)) = v85
	*(*int32)(unsafe.Add(mBase, _c_F_CleanupTransaction[3])) = v85
	*(*int32)(unsafe.Add(mBase, _c_F_CleanupTransaction[4])) = v85
	v95 = *(*int32)(unsafe.Add(mBase, _c_F_CleanupTransaction[0]))
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v95)+44))
	*(*int32)(unsafe.Add(mBase, _c_F_CleanupTransaction[5])) = v96
	v99 = *(*int32)(unsafe.Add(mBase, _c_F_CleanupTransaction[6]))
	if v99 != 0 {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	goto L34
L36:
	;
	F_MemoryContextReset(m, v99)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L4
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	v103 = *(*int32)(unsafe.Add(mBase, _c_F_CleanupTransaction[7]))
	if v103 != 0 {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	goto L38
L40:
	;
	F_MemoryContextReset(m, v103)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L4
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	v107 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_CleanupTransaction[8])) = v107
	*(*int32)(unsafe.Add(mBase, uint32(v95)+36)) = v107
	*(*uint8)(unsafe.Add(mBase, uint32(v10)+76)) = uint8(v107)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+72)) = v107
	*(*int32)(unsafe.Add(mBase, uint32(v10)+56)) = v107
	v117 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v10)+48)) = v117
	*(*int64)(unsafe.Add(mBase, uint32(v10)+28)) = v117
	*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v107
	*(*int64)(unsafe.Add(mBase, uint32(v10))) = v117
	*(*int64)(unsafe.Add(mBase, _c_F_CleanupTransaction[9])) = v117
	*(*int32)(unsafe.Add(mBase, _c_F_CleanupTransaction[10])) = v107
	*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v107
	m.G0 = v7 + int32(16)
	return
L43:
	;
	goto L42
L44:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v10)+20))
	if base.Ui32(v140) <= base.Ui32(int32(5)) {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v140<<(uint(int32(2))%32))+uint32(_c_F_CleanupTransaction[11])))
	v147 = v145
	goto L47
L46:
	;
	v147 = int32(_a_F_CleanupTransaction_3)
	goto L47
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v147
	F_errmsg_internal(m, int32(_a_F_CleanupTransaction_4), v7)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L4
	} else {
		goto L48
	}
L48:
	;
	F_errfinish(m, int32(_a_F_CleanupTransaction_5), int32(3071), int32(_a_F_CleanupTransaction_6))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L4
	} else {
		goto L49
	}
L49:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_LockApplyTransactionForSession(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	v3 = l2
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = int32(267)
	*(*uint16)(unsafe.Add(mBase, uint32(v8)+14)) = uint16(v10)
	*(*uint16)(unsafe.Add(mBase, uint32(v8)+12)) = uint16(v3)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = l0
	v16 = *(*int32)(unsafe.Add(mBase, _c_F_LockApplyTransactionForSession[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v16
	v20 = F_LockAcquire(m, v8, l3, int32(1), int32(0))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		return
	} else {
		m.G0 = v8 + int32(16)
		return
	}
}
func F_PrepareTransactionBlock(m *base.Module, l0 int32) int32 {
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
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	v2 = int32(0)
	v5 = F_EndTransactionBlock(m, v2)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		if v5 == int32(0) {
			v31 = v2
			return v31
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransactionBlock[0]))
			v14 = v12
			for {
				v16 = *(*int32)(unsafe.Add(mBase, uint32(v14)+80))
				if v16 != 0 {
					v14 = v16
					continue
				} else {
					break
				}
				break
			}
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
			if v18 != int32(6) {
				v31 = int32(0)
				return v31
			} else {
				v23 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransactionBlock[1]))
				v24 = F_MemoryContextStrdup(m, v23, l0)
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransactionBlock[2])) = v24
					*(*int32)(unsafe.Add(mBase, uint32(v14)+24)) = int32(10)
					v31 = int32(1)
					return v31
				}
			}
		}
	}
}
func F_RequireTransactionBlock(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_RequireTransactionBlock[0]))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+24))
	if base.B2i32(l0 == int32(0))|base.B2i32(base.Ui32(int32(1)) < base.Ui32(v12)) != 0 {
		m.G0 = v6 + int32(16)
		return
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
		if int32(1) < v16 {
			m.G0 = v6 + int32(16)
			return
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return
			} else {
				F_errcode(m, int32(16908610))
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v6))) = l1
					F_errmsg(m, int32(_a_F_RequireTransactionBlock_0), v6)
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_RequireTransactionBlock_1), int32(3803), int32(_a_F_RequireTransactionBlock_2))
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
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
func F_RestoreTransactionCharacteristics(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, _c_F_RestoreTransactionCharacteristics[0])) = v3
	v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)))
	*(*uint8)(unsafe.Add(mBase, _c_F_RestoreTransactionCharacteristics[1])) = uint8(v6)
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+5)))
	*(*uint8)(unsafe.Add(mBase, _c_F_RestoreTransactionCharacteristics[2])) = uint8(v9)
	return
}
func F_RestoreTransactionSnapshot(m *base.Module, l0 int32, l1 int32) {
	var v6 int32
	_ = v6
	F_SetTransactionSnapshot(m, l0, int32(0), int32(-1), l1)
	v6 = m.ExcPending
	if v6 != 0 {
		return
	} else {
		return
	}
}
func F_StartTransaction(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v21 float64
	_ = v21
	var v27 int32
	_ = v27
	var v30 int64
	_ = v30
	var v31 int64
	_ = v31
	var v32 int64
	_ = v32
	var v53 float64
	_ = v53
	var v55 float64
	_ = v55
	var v57 int32
	_ = v57
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v197 int64
	_ = v197
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v209 int64
	_ = v209
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v223 int32
	_ = v223
	var v228 int64
	_ = v228
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v239 int64
	_ = v239
	var v240 int64
	_ = v240
	var v248 int64
	_ = v248
	var v252 int64
	_ = v252
	var v255 int32
	_ = v255
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v287 int32
	_ = v287
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v300 int32
	_ = v300
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v310 int32
	_ = v310
	var v314 int32
	_ = v314
	var v319 int32
	_ = v319
	var v327 int32
	_ = v327
	var v338 int32
	_ = v338
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v351 int32
	_ = v351
	var v354 int32
	_ = v354
	var v367 int32
	_ = v367
	var v371 int32
	_ = v371
	var v375 int32
	_ = v375
	var v381 int32
	_ = v381
	var v384 int32
	_ = v384
	var v388 int32
	_ = v388
	var v390 int32
	_ = v390
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v11 = int32(_a_F_StartTransaction_0)
	*(*int32)(unsafe.Add(mBase, _c_F_StartTransaction[0])) = v11
	*(*int32)(unsafe.Add(mBase, _c_F_StartTransaction[1])) = int32(1)
	*(*int64)(unsafe.Add(mBase, _c_F_StartTransaction[2])) = int64(0)
	v21 = *(*float64)(unsafe.Add(mBase, _c_F_StartTransaction[3]))
	if base.F64_eq(v21, float64(0)) != 0 {
		v57 = int32(0)
		goto L1
	} else {
		goto L2
	}
L1:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_StartTransaction[6])) = int64(0)
	*(*int64)(unsafe.Add(mBase, _c_F_StartTransaction[7])) = int64(4294967297)
	*(*uint8)(unsafe.Add(mBase, _c_F_StartTransaction[8])) = uint8(v57)
	*(*int32)(unsafe.Add(mBase, _c_F_StartTransaction[9])) = int32(0)
	v72 = *(*int32)(unsafe.Add(mBase, _c_F_StartTransaction[10]))
	*(*int32)(unsafe.Add(mBase, _c_F_StartTransaction[11])) = v72
	v75 = *(*int32)(unsafe.Add(mBase, _c_F_StartTransaction[12]))
	*(*int32)(unsafe.Add(mBase, _c_F_StartTransaction[13])) = v75
	goto L5
L2:
	;
	if base.F64_eq(v21, float64(1)) != 0 {
		v57 = int32(1)
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v27 = int32(_a_F_StartTransaction_1)
	v30 = *(*int64)(unsafe.Add(mBase, _c_F_StartTransaction[4]))
	v31 = *(*int64)(unsafe.Add(mBase, _c_F_StartTransaction[5]))
	v32 = v30 ^ v31
	*(*int64)(unsafe.Add(mBase, _c_F_StartTransaction[5])) = base.I64_rotl(v32, int64(37))
	*(*int64)(unsafe.Add(mBase, _c_F_StartTransaction[4])) = v32<<(uint(int64(16))%64) ^ base.I64_rotl(v30, int64(24)) ^ v32
	v53 = F_scalbn(m, base.F64_convert_i64_u(int64(base.Ui64(base.I64_rotl(v30*int64(5), int64(7))*int64(9))>>(uint(int64(12))%64))), int32(-52))
	mBase = m.M
	goto L4
L4:
	;
	v55 = *(*float64)(unsafe.Add(mBase, _c_F_StartTransaction[3]))
	v57 = base.F64_le(v53, v55)
	goto L1
L5:
	;
	v80 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartTransaction[14])))
	if v80 == int32(1) {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	*(*uint8)(unsafe.Add(mBase, _c_F_StartTransaction[16])) = uint8(v90)
	v93 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_StartTransaction[17])) = v93
	*(*int32)(unsafe.Add(mBase, _c_F_StartTransaction[18])) = v93
	v100 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartTransaction[19])))
	*(*uint8)(unsafe.Add(mBase, _c_F_StartTransaction[20])) = uint8(v100)
	v104 = *(*int32)(unsafe.Add(mBase, _c_F_StartTransaction[21]))
	*(*int32)(unsafe.Add(mBase, _c_F_StartTransaction[22])) = v104
	v109 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartTransaction[23])))
	if v90 != 0 {
		goto L10
	} else {
		goto L11
	}
L7:
	;
	v85 = *(*int32)(unsafe.Add(mBase, _c_F_StartTransaction[15]))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)+308))
	v88 = base.B2i32(v86 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _c_F_StartTransaction[14])) = uint8(v88)
	v90 = v88
	goto L9
L8:
	;
	v90 = int32(0)
	goto L9
L9:
	;
	goto L6
L10:
	;
	v110 = v93
	goto L12
L11:
	;
	v110 = v109
	goto L12
L12:
	;
	*(*uint8)(unsafe.Add(mBase, _c_F_StartTransaction[24])) = uint8(v110)
	v113 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_StartTransaction[25])) = uint8(v113)
	*(*int32)(unsafe.Add(mBase, _c_F_StartTransaction[26])) = v113
	*(*int32)(unsafe.Add(mBase, _c_F_StartTransaction[27])) = v113
	*(*uint8)(unsafe.Add(mBase, _c_F_StartTransaction[28])) = uint8(v113)
	*(*int32)(unsafe.Add(mBase, _c_F_StartTransaction[29])) = v113
	*(*uint8)(unsafe.Add(mBase, _c_F_StartTransaction[30])) = uint8(v113)
	v131 = *(*int32)(unsafe.Add(mBase, _c_F_StartTransaction[0]))
	v133 = *(*int32)(unsafe.Add(mBase, _c_F_StartTransaction[31]))
	*(*int32)(unsafe.Add(mBase, uint32(v131)+44)) = v133
	v136 = *(*int32)(unsafe.Add(mBase, _c_F_StartTransaction[32]))
	if v136 == v113 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v141 = *(*int32)(unsafe.Add(mBase, _c_F_StartTransaction[33]))
	v143 = int32(_a_F_StartTransaction_2)
	v146 = F_AllocSetContextCreateInternal(m, v141, int32(_a_F_StartTransaction_3), v143, v143, v143)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	goto L15
L15:
	;
	v150 = *(*int32)(unsafe.Add(mBase, _c_F_StartTransaction[34]))
	if v150 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	return
L17:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_StartTransaction[32])) = v146
	goto L15
L18:
	;
	v155 = *(*int32)(unsafe.Add(mBase, _c_F_StartTransaction[33]))
	v160 = F_AllocSetContextCreateInternal(m, v155, int32(_a_F_StartTransaction_4), int32(0), int32(_a_F_StartTransaction_5), int32(_a_F_StartTransaction_6))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L16
	} else {
		goto L21
	}
L19:
	;
	v163 = v150
	goto L20
L20:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_StartTransaction[35])) = v163
	*(*int32)(unsafe.Add(mBase, uint32(v131)+36)) = v163
	*(*int32)(unsafe.Add(mBase, _c_F_StartTransaction[31])) = v163
	v170 = *(*int32)(unsafe.Add(mBase, _c_F_StartTransaction[0]))
	v173 = F_ResourceOwnerCreate(m, int32(0), int32(_a_F_StartTransaction_7))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L16
	} else {
		goto L22
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_StartTransaction[34])) = v160
	v163 = v160
	goto L20
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v170)+40)) = v173
	*(*int32)(unsafe.Add(mBase, _c_F_StartTransaction[36])) = v173
	*(*int32)(unsafe.Add(mBase, _c_F_StartTransaction[37])) = v173
	*(*int32)(unsafe.Add(mBase, _c_F_StartTransaction[38])) = v173
	v183 = *(*int32)(unsafe.Add(mBase, _c_F_StartTransaction[39]))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v183
	v186 = int32(_a_F_StartTransaction_8)
	v187 = int32(1)
	v189 = *(*int32)(unsafe.Add(mBase, _c_F_StartTransaction[40]))
	if base.Ui32(v189) <= base.Ui32(v187) {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v192
	v197 = *(*int64)(unsafe.Add(mBase, uint32(v8)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v8))) = v197
	F_VirtualXactLockTableInsert(m, v8)
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L16
	} else {
		goto L27
	}
L24:
	;
	v192 = v187
	goto L26
L25:
	;
	v192 = v189
	goto L26
L26:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_StartTransaction[40])) = v192 + int32(1)
	goto L23
L27:
	;
	v202 = *(*int32)(unsafe.Add(mBase, _c_F_StartTransaction[41]))
	*(*int32)(unsafe.Add(mBase, uint32(v202)+44)) = v192
	v205 = *(*int32)(unsafe.Add(mBase, _c_F_StartTransaction[42]))
	if int32(0) <= v205 {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	v255 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartTransaction[46])))
	if v255 != int32(1) {
		goto L41
	} else {
		goto L42
	}
L29:
	;
	v209 = *(*int64)(unsafe.Add(mBase, _c_F_StartTransaction[43]))
	v252 = v209
	goto L28
L30:
	;
	goto L31
L31:
	;
	v210 = int32(0)
	v212 = *(*int32)(unsafe.Add(mBase, _c_F_StartTransaction[44]))
	if v212 == v210 {
		v223 = v210
		goto L32
	} else {
		goto L33
	}
L32:
	;
	if v223 == int32(0) {
		goto L36
	} else {
		goto L37
	}
L33:
	;
	v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v212)+40)))
	if v215 != 0 {
		v223 = v210
		goto L32
	} else {
		goto L34
	}
L34:
	;
	v217 = *(*int32)(unsafe.Add(mBase, _c_F_StartTransaction[0]))
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v217)+28))
	goto L35
L35:
	;
	v223 = base.B2i32(int32(1) < v218) ^ int32(1)
	goto L32
L36:
	;
	v228 = *(*int64)(unsafe.Add(mBase, _c_F_StartTransaction[45]))
	*(*int64)(unsafe.Add(mBase, _c_F_StartTransaction[43])) = v228
	v252 = v228
	goto L28
L37:
	;
	goto L38
L38:
	;
	v234 = m.G0
	v235 = int32(16)
	v236 = v234 - v235
	m.G0 = v236
	F_gettimeofday(m, v236)
	mBase = m.M
	v239 = *(*int64)(unsafe.Add(mBase, uint32(v236)))
	v240 = int64(*(*int32)(unsafe.Add(mBase, uint32(v236)+8)))
	m.G0 = v236 + v235
	v248 = v240 + v239*int64(1000000) - int64(946684800000000)
	goto L39
L39:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_StartTransaction[43])) = v248
	v252 = v248
	goto L28
L40:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_StartTransaction[49])) = int64(0)
	v295 = m.G0
	v297 = v295 - int32(16)
	m.G0 = v297
	v300 = *(*int32)(unsafe.Add(mBase, _c_F_StartTransaction[50]))
	if v300 == int32(0) {
		goto L44
	} else {
		goto L45
	}
L41:
	;
	goto L40
L42:
	;
	v259 = *(*int32)(unsafe.Add(mBase, _c_F_StartTransaction[47]))
	if v259 == int32(0) {
		goto L41
	} else {
		goto L43
	}
L43:
	;
	v262 = int32(_a_F_StartTransaction_9)
	v264 = *(*int32)(unsafe.Add(mBase, _c_F_StartTransaction[48]))
	v265 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_StartTransaction[48])) = v264 + v265
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v259)))
	*(*int32)(unsafe.Add(mBase, uint32(v259))) = v268 + v265
	v272 = int32(0)
	v274 = int32(_a_F_StartTransaction_10)
	v275 = base.AtomicRmwOr32(m, v272, v274, v272)
	*(*int64)(unsafe.Add(mBase, uint32(v259)+24)) = v252
	v280 = base.AtomicRmwOr32(m, v272, v274, v272)
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v259)))
	*(*int32)(unsafe.Add(mBase, uint32(v259))) = v281 + v265
	v287 = *(*int32)(unsafe.Add(mBase, _c_F_StartTransaction[48]))
	*(*int32)(unsafe.Add(mBase, _c_F_StartTransaction[48])) = v287 - v265
	goto L41
L44:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_StartTransaction[50])) = int32(1)
	m.G0 = v297 + int32(16)
	F_ReceiveSharedInvalidMessages(m)
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L16
	} else {
		goto L50
	}
L45:
	;
	v305 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L16
	} else {
		goto L46
	}
L46:
	;
	if v305 == int32(0) {
		goto L44
	} else {
		goto L47
	}
L47:
	;
	v310 = *(*int32)(unsafe.Add(mBase, _c_F_StartTransaction[50]))
	*(*int32)(unsafe.Add(mBase, uint32(v297))) = v310
	F_errmsg_internal(m, int32(_a_F_StartTransaction_12), v297)
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L16
	} else {
		goto L48
	}
L48:
	;
	F_errfinish(m, int32(_a_F_StartTransaction_13), int32(2131), int32(_a_F_StartTransaction_14))
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L16
	} else {
		goto L49
	}
L49:
	;
	goto L44
L50:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_StartTransaction[51])) = int32(-1)
	*(*int32)(unsafe.Add(mBase, _c_F_StartTransaction[52])) = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_StartTransaction[1])) = int32(2)
	v338 = *(*int32)(unsafe.Add(mBase, _c_F_StartTransaction[53]))
	if int32(0) < v338 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	F_enable_timeout_after(m, int32(8), v338)
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L16
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	v344 = int32(10)
	goto L57
L54:
	;
	goto L53
L55:
	;
	if v384 != 0 {
		goto L68
	} else {
		goto L69
	}
L56:
	;
	goto L55
L57:
	;
	v351 = *(*int32)(unsafe.Add(mBase, _c_F_StartTransaction[54]))
	v354 = *(*int32)(unsafe.Add(mBase, uint32(v351<<(uint(int32(2))%32))+uint32(_c_F_StartTransaction[55])))
	goto L60
L58:
	;
	v367 = int32(0)
	goto L65
L60:
	;
	goto L61
L61:
	;
	if int32(0)|base.B2i32(v354 == int32(15)) != 0 {
		goto L58
	} else {
		goto L63
	}
L63:
	;
	if v354 <= v344 {
		v384 = int32(1)
		goto L56
	} else {
		goto L64
	}
L64:
	;
	goto L58
L65:
	;
	v371 = *(*int32)(unsafe.Add(mBase, _c_F_StartTransaction[56]))
	if v371 != int32(2) {
		v384 = v367
		goto L56
	} else {
		goto L66
	}
L66:
	;
	v375 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartTransaction[57])))
	if v375&int32(1) != 0 {
		v384 = v367
		goto L56
	} else {
		goto L67
	}
L67:
	;
	v381 = *(*int32)(unsafe.Add(mBase, _c_F_StartTransaction[58]))
	v384 = int32(0) | base.B2i32(v381 <= v344)
	goto L56
L68:
	;
	v388 = *(*int32)(unsafe.Add(mBase, _c_F_StartTransaction[0]))
	F_ShowTransactionStateRec(m, int32(_a_F_StartTransaction_11), v388)
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L16
	} else {
		goto L71
	}
L69:
	;
	goto L70
L70:
	;
	m.G0 = v8 + int32(16)
	return
L71:
	;
	goto L70
}
func F_TransactionIdAsyncCommitTree(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64) {
	var v7 int32
	_ = v7
	F_TransactionIdSetTreeStatus(m, l0, l1, l2, int32(1), l3)
	v7 = m.ExcPending
	if v7 != 0 {
		return
	} else {
		return
	}
}
func F_check_transaction_deferrable(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	v4 = int32(1)
	v6 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_check_transaction_deferrable[0])))
	if v6 != 0 {
		v34 = v4
		return v34
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, _c_F_check_transaction_deferrable[1]))
		v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+28))
		if int32(1) < v9 {
			v18 = int32(_a_F_check_transaction_deferrable_0)
			*(*int32)(unsafe.Add(mBase, _c_F_check_transaction_deferrable[2])) = int32(16777538)
			v24 = *(*int32)(unsafe.Add(mBase, _c_F_check_transaction_deferrable[3]))
			*(*int32)(unsafe.Add(mBase, _c_F_check_transaction_deferrable[4])) = v24
			v29 = F_format_elog_string(m, v18, int32(0))
			mBase = m.M
			v32 = m.ExcPending
			if v32 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, _c_F_check_transaction_deferrable[5])) = v29
				v34 = int32(0)
				return v34
			}
		} else {
			v14 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_check_transaction_deferrable[6])))
			if v14 != int32(1) {
				v34 = v4
				return v34
			} else {
				v18 = int32(_a_F_check_transaction_deferrable_1)
				*(*int32)(unsafe.Add(mBase, _c_F_check_transaction_deferrable[2])) = int32(16777538)
				v24 = *(*int32)(unsafe.Add(mBase, _c_F_check_transaction_deferrable[3]))
				*(*int32)(unsafe.Add(mBase, _c_F_check_transaction_deferrable[4])) = v24
				v29 = F_format_elog_string(m, v18, int32(0))
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, _c_F_check_transaction_deferrable[5])) = v29
					v34 = int32(0)
					return v34
				}
			}
		}
	}
}
func F_check_transaction_isolation(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	v4 = int32(1)
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v7 = *(*int32)(unsafe.Add(mBase, _c_F_check_transaction_isolation[0]))
	if v5 == v7 {
		v84 = v4
		return v84
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, _c_F_check_transaction_isolation[1]))
		v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+20))
		if base.B2i32(v11 == int32(2)) == int32(0) {
			v84 = v4
			return v84
		} else {
			v17 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_check_transaction_isolation[2])))
			if v17&int32(1) != 0 {
				v84 = v4
				return v84
			} else {
				v21 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_check_transaction_isolation[3])))
				if v21 == int32(1) {
					*(*int32)(unsafe.Add(mBase, _c_F_check_transaction_isolation[4])) = int32(16777538)
					v73 = int32(_a_F_check_transaction_isolation_0)
					v74 = int32(_a_F_check_transaction_isolation_1)
					v77 = *(*int32)(unsafe.Add(mBase, _c_F_check_transaction_isolation[5]))
					*(*int32)(unsafe.Add(mBase, _c_F_check_transaction_isolation[6])) = v77
					v81 = F_format_elog_string(m, v73, int32(0))
					mBase = m.M
					v82 = m.ExcPending
					if v82 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v74))) = v81
						v84 = int32(0)
						return v84
					}
				} else {
					v30 = *(*int32)(unsafe.Add(mBase, _c_F_check_transaction_isolation[1]))
					v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+28))
					if int32(1) < v31 {
						*(*int32)(unsafe.Add(mBase, _c_F_check_transaction_isolation[4])) = int32(16777538)
						v73 = int32(_a_F_check_transaction_isolation_2)
						v74 = int32(_a_F_check_transaction_isolation_1)
						v77 = *(*int32)(unsafe.Add(mBase, _c_F_check_transaction_isolation[5]))
						*(*int32)(unsafe.Add(mBase, _c_F_check_transaction_isolation[6])) = v77
						v81 = F_format_elog_string(m, v73, int32(0))
						mBase = m.M
						v82 = m.ExcPending
						if v82 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v74))) = v81
							v84 = int32(0)
							return v84
						}
					} else {
						if v5 != int32(3) {
							v84 = v4
							return v84
						} else {
							v43 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_check_transaction_isolation[7])))
							if v43 == int32(1) {
								v48 = *(*int32)(unsafe.Add(mBase, _c_F_check_transaction_isolation[8]))
								v49 = *(*int32)(unsafe.Add(mBase, uint32(v48)+308))
								v51 = base.B2i32(v49 != int32(2))
								*(*uint8)(unsafe.Add(mBase, _c_F_check_transaction_isolation[7])) = uint8(v51)
								v53 = v51
							} else {
								v53 = int32(0)
							}
							if v53 == int32(0) {
								v84 = v4
								return v84
							} else {
								*(*int32)(unsafe.Add(mBase, _c_F_check_transaction_isolation[4])) = int32(1088)
								v60 = *(*int32)(unsafe.Add(mBase, _c_F_check_transaction_isolation[5]))
								*(*int32)(unsafe.Add(mBase, _c_F_check_transaction_isolation[6])) = v60
								v66 = F_format_elog_string(m, int32(_a_F_check_transaction_isolation_3), int32(0))
								mBase = m.M
								v69 = m.ExcPending
								if v69 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, _c_F_check_transaction_isolation[9])) = v66
									v73 = int32(_a_F_check_transaction_isolation_4)
									v74 = int32(_a_F_check_transaction_isolation_5)
									v77 = *(*int32)(unsafe.Add(mBase, _c_F_check_transaction_isolation[5]))
									*(*int32)(unsafe.Add(mBase, _c_F_check_transaction_isolation[6])) = v77
									v81 = F_format_elog_string(m, v73, int32(0))
									mBase = m.M
									v82 = m.ExcPending
									if v82 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v74))) = v81
										v84 = int32(0)
										return v84
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
