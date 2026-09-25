package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_AbortSubTransaction(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v169 int32
	_ = v169
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v197 int32
	_ = v197
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v286 int32
	_ = v286
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v297 int32
	_ = v297
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v325 int32
	_ = v325
	var v328 int32
	_ = v328
	var v331 int32
	_ = v331
	var v334 int32
	_ = v334
	var v336 int32
	_ = v336
	var v340 int32
	_ = v340
	var v346 int32
	_ = v346
	var v349 int32
	_ = v349
	var v352 int32
	_ = v352
	var v354 int32
	_ = v354
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v370 int32
	_ = v370
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v383 int32
	_ = v383
	var v389 int32
	_ = v389
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v394 int32
	_ = v394
	var v402 int32
	_ = v402
	var v404 int32
	_ = v404
	var v407 int32
	_ = v407
	var v409 int32
	_ = v409
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v415 int32
	_ = v415
	var v417 int32
	_ = v417
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v423 int32
	_ = v423
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v429 int32
	_ = v429
	var v432 int32
	_ = v432
	var v435 int32
	_ = v435
	var v437 int32
	_ = v437
	var v439 int32
	_ = v439
	var v441 int32
	_ = v441
	var v443 int32
	_ = v443
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v449 int32
	_ = v449
	var v451 int32
	_ = v451
	var v455 int32
	_ = v455
	var v460 int32
	_ = v460
	var v463 int32
	_ = v463
	var v475 int32
	_ = v475
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v483 int32
	_ = v483
	var v485 int32
	_ = v485
	var v487 int32
	_ = v487
	var v489 int32
	_ = v489
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v495 int32
	_ = v495
	var v498 int32
	_ = v498
	var v503 int32
	_ = v503
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v515 int32
	_ = v515
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v520 int32
	_ = v520
	var v531 int32
	_ = v531
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v536 int32
	_ = v536
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v551 int32
	_ = v551
	var v554 int32
	_ = v554
	var v556 int32
	_ = v556
	var v560 int32
	_ = v560
	var v577 int32
	_ = v577
	var v579 int32
	_ = v579
	var v581 int32
	_ = v581
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v14 = *(*int32)(unsafe.Add(mBase, _c_F_AbortSubTransaction[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_AbortSubTransaction[1])) = v14
	v16 = int32(_a_F_AbortSubTransaction_0)
	v18 = *(*int32)(unsafe.Add(mBase, _c_F_AbortSubTransaction[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_AbortSubTransaction[2])) = v18 + int32(1)
	v24 = *(*int32)(unsafe.Add(mBase, _c_F_AbortSubTransaction[3]))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+40))
	*(*int32)(unsafe.Add(mBase, _c_F_AbortSubTransaction[4])) = v25
	F_LWLockReleaseAll(m)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v30 = *(*int32)(unsafe.Add(mBase, _c_F_AbortSubTransaction[5]))
	v31 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v30))) = v31
	v35 = *(*int32)(unsafe.Add(mBase, _c_F_AbortSubTransaction[6]))
	if v35 == v31 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	F_pgaio_error_cleanup(m)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L1
	} else {
		goto L8
	}
L4:
	;
	goto L3
L5:
	;
	v39 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_AbortSubTransaction[7])))
	if v39&int32(1) == int32(0) {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v35)+220))
	if v44 == int32(0) {
		goto L4
	} else {
		goto L7
	}
L7:
	;
	v47 = int32(_a_F_AbortSubTransaction_1)
	v49 = *(*int32)(unsafe.Add(mBase, _c_F_AbortSubTransaction[8]))
	v50 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_AbortSubTransaction[8])) = v49 + v50
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	*(*int32)(unsafe.Add(mBase, uint32(v35))) = v53 + v50
	v57 = int32(0)
	v59 = int32(_a_F_AbortSubTransaction_2)
	v60 = base.AtomicRmwOr32(m, v57, v59, v57)
	*(*int32)(unsafe.Add(mBase, uint32(v35)+220)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v35)+224)) = v57
	v68 = base.AtomicRmwOr32(m, v57, v59, v57)
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	*(*int32)(unsafe.Add(mBase, uint32(v35))) = v69 + v50
	v75 = *(*int32)(unsafe.Add(mBase, _c_F_AbortSubTransaction[8]))
	*(*int32)(unsafe.Add(mBase, _c_F_AbortSubTransaction[8])) = v75 - v50
	goto L4
L8:
	;
	F_UnlockBuffers(m)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	v83 = int32(0)
	v90 = *(*int32)(unsafe.Add(mBase, _c_F_AbortSubTransaction[9]))
	if v90 <= v83 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	F_ConditionVariableCancelSleep(m)
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L1
	} else {
		goto L23
	}
L11:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_AbortSubTransaction[10])) = int64(0)
	*(*int32)(unsafe.Add(mBase, _c_F_AbortSubTransaction[11])) = int32(_a_F_AbortSubTransaction_3)
	v169 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_AbortSubTransaction[9])) = v169
	*(*int32)(unsafe.Add(mBase, _c_F_AbortSubTransaction[12])) = v169
	*(*uint8)(unsafe.Add(mBase, _c_F_AbortSubTransaction[13])) = uint8(v169)
	*(*uint8)(unsafe.Add(mBase, _c_F_AbortSubTransaction[14])) = uint8(v169)
	goto L10
L12:
	;
	v94 = v90 & int32(7)
	v96 = *(*int32)(unsafe.Add(mBase, _c_F_AbortSubTransaction[15]))
	if base.Ui32(int32(8)) <= base.Ui32(v90) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v102 = v83
	v106 = v83
	goto L16
L14:
	;
	v134 = v83
	goto L15
L15:
	;
	v140 = int32(0)
	v141 = v134
	goto L20
L16:
	;
	v109 = v96 + v102*int32(_a_F_AbortSubTransaction_4)
	v110 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v109)+uint32(_c_F_AbortSubTransaction[16]))) = uint8(v110)
	*(*uint8)(unsafe.Add(mBase, uint32(v109)+uint32(_c_F_AbortSubTransaction[17]))) = uint8(v110)
	*(*uint8)(unsafe.Add(mBase, uint32(v109)+uint32(_c_F_AbortSubTransaction[18]))) = uint8(v110)
	*(*uint8)(unsafe.Add(mBase, uint32(v109)+uint32(_c_F_AbortSubTransaction[19]))) = uint8(v110)
	*(*uint8)(unsafe.Add(mBase, uint32(v109)+uint32(_c_F_AbortSubTransaction[20]))) = uint8(v110)
	*(*uint8)(unsafe.Add(mBase, uint32(v109)+uint32(_c_F_AbortSubTransaction[21]))) = uint8(v110)
	*(*uint8)(unsafe.Add(mBase, uint32(v109)+uint32(_c_F_AbortSubTransaction[22]))) = uint8(v110)
	*(*uint8)(unsafe.Add(mBase, uint32(v109))) = uint8(v110)
	v126 = int32(8)
	v127 = v102 + v126
	v129 = v106 + v126
	if v129 != v90&int32(2147483640) {
		v102 = v127
		v106 = v129
		goto L16
	} else {
		goto L18
	}
L17:
	;
	if v94 == int32(0) {
		goto L11
	} else {
		goto L19
	}
L18:
	;
	goto L17
L19:
	;
	v134 = v127
	goto L15
L20:
	;
	v149 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v96+v141*int32(_a_F_AbortSubTransaction_4)))) = uint8(v149)
	v151 = int32(1)
	v154 = v140 + v151
	if v154 != v94 {
		v140 = v154
		v141 = v141 + v151
		goto L20
	} else {
		goto L22
	}
L21:
	;
	goto L11
L22:
	;
	goto L21
L23:
	;
	F_LockErrorCleanup(m)
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	F_reschedule_timeouts(m)
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	F_pgmem_sigprocmask(m, int32(_a_F_AbortSubTransaction_5), int32(0))
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	v190 = int32(10)
	goto L29
L27:
	;
	if v227 != 0 {
		goto L40
	} else {
		goto L41
	}
L28:
	;
	goto L27
L29:
	;
	v197 = *(*int32)(unsafe.Add(mBase, _c_F_AbortSubTransaction[23]))
	goto L32
L30:
	;
	v210 = int32(0)
	goto L37
L32:
	;
	goto L33
L33:
	;
	if int32(0)|base.B2i32(v197 == int32(15)) != 0 {
		goto L30
	} else {
		goto L35
	}
L35:
	;
	if v197 <= v190 {
		v227 = int32(1)
		goto L28
	} else {
		goto L36
	}
L36:
	;
	goto L30
L37:
	;
	v214 = *(*int32)(unsafe.Add(mBase, _c_F_AbortSubTransaction[24]))
	if v214 != int32(2) {
		v227 = v210
		goto L28
	} else {
		goto L38
	}
L38:
	;
	v218 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_AbortSubTransaction[25])))
	if v218&int32(1) != 0 {
		v227 = v210
		goto L28
	} else {
		goto L39
	}
L39:
	;
	v224 = *(*int32)(unsafe.Add(mBase, _c_F_AbortSubTransaction[26]))
	v227 = int32(0) | base.B2i32(v224 <= v190)
	goto L28
L40:
	;
	v231 = *(*int32)(unsafe.Add(mBase, _c_F_AbortSubTransaction[3]))
	F_ShowTransactionStateRec(m, int32(_a_F_AbortSubTransaction_6), v231)
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L1
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
	if v234 == int32(2) {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	goto L42
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+20)) = int32(4)
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v24)+60))
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v24)+64))
	*(*int32)(unsafe.Add(mBase, _c_F_AbortSubTransaction[27])) = v264
	*(*int32)(unsafe.Add(mBase, _c_F_AbortSubTransaction[28])) = v263
	goto L53
L45:
	;
	v239 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	if v239 == int32(0) {
		goto L44
	} else {
		goto L47
	}
L47:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
	if base.Ui32(v243) <= base.Ui32(int32(5)) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v243<<(uint(int32(2))%32))+uint32(_c_F_AbortSubTransaction[29])))
	v250 = v248
	goto L50
L49:
	;
	v250 = int32(_a_F_AbortSubTransaction_7)
	goto L50
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v250
	F_errmsg_internal(m, int32(_a_F_AbortSubTransaction_8), v10)
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	F_errfinish(m, int32(_a_F_AbortSubTransaction_9), int32(_a_F_AbortSubTransaction_10), int32(_a_F_AbortSubTransaction_6))
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	goto L44
L53:
	;
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v24)+28))
	v271 = *(*int32)(unsafe.Add(mBase, _c_F_AbortSubTransaction[30]))
	if v269 <= v271 {
		goto L55
	} else {
		goto L56
	}
L54:
	;
	v286 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_AbortSubTransaction[31])) = uint8(v286)
	*(*int32)(unsafe.Add(mBase, _c_F_AbortSubTransaction[32])) = v286
	goto L58
L55:
	;
	v274 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_AbortSubTransaction[33])) = v274
	*(*int32)(unsafe.Add(mBase, _c_F_AbortSubTransaction[34])) = v274
	*(*int32)(unsafe.Add(mBase, _c_F_AbortSubTransaction[35])) = v274
	*(*int32)(unsafe.Add(mBase, _c_F_AbortSubTransaction[30])) = v274
	goto L57
L56:
	;
	goto L57
L57:
	;
	goto L54
L58:
	;
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
	F_AtEOSubXact_Parallel(m, int32(0), v292)
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+72)) = int32(0)
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v24)+40))
	if v297 != 0 {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	F_AfterTriggerEndSubXact(m, int32(0))
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L1
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	v577 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+68)))
	*(*uint8)(unsafe.Add(mBase, _c_F_AbortSubTransaction[36])) = uint8(v577)
	v579 = int32(_a_F_AbortSubTransaction_0)
	v581 = *(*int32)(unsafe.Add(mBase, _c_F_AbortSubTransaction[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_AbortSubTransaction[2])) = v581 - int32(1)
	m.G0 = v10 + int32(16)
	return
L63:
	;
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v24)+80))
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v301)+8))
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v24)+40))
	F_AtSubAbort_Portals(m, v304, v302, v305)
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v24)+80))
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v310)+8))
	F_AtEOSubXact_LargeObject(m, int32(0), v309, v311)
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	v315 = *(*int32)(unsafe.Add(mBase, _c_F_AbortSubTransaction[3]))
	v316 = *(*int32)(unsafe.Add(mBase, uint32(v315)+28))
	goto L66
L66:
	;
	goto L67
L67:
	;
	v325 = *(*int32)(unsafe.Add(mBase, _c_F_AbortSubTransaction[37]))
	if v325 == int32(0) {
		goto L69
	} else {
		goto L70
	}
L68:
	;
	v336 = *(*int32)(unsafe.Add(mBase, _c_F_AbortSubTransaction[38]))
	if v336 == int32(0) {
		goto L73
	} else {
		goto L74
	}
L69:
	;
	goto L68
L70:
	;
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v325)))
	if v328 < v316 {
		goto L69
	} else {
		goto L71
	}
L71:
	;
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v325)+8))
	*(*int32)(unsafe.Add(mBase, _c_F_AbortSubTransaction[37])) = v331
	F_pfree(m, v325)
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	goto L67
L73:
	;
	v363 = F_RecordTransactionAbort(m, int32(1))
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L1
	} else {
		goto L80
	}
L74:
	;
	v340 = v336
	goto L75
L75:
	;
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v340)))
	if v346 < v316 {
		goto L73
	} else {
		goto L77
	}
L76:
	;
	goto L73
L77:
	;
	v349 = *(*int32)(unsafe.Add(mBase, uint32(v340)+12))
	*(*int32)(unsafe.Add(mBase, _c_F_AbortSubTransaction[38])) = v349
	F_pfree(m, v340)
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L1
	} else {
		goto L78
	}
L78:
	;
	v354 = *(*int32)(unsafe.Add(mBase, _c_F_AbortSubTransaction[38]))
	if v354 != 0 {
		v340 = v354
		goto L75
	} else {
		goto L79
	}
L79:
	;
	goto L76
L80:
	;
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	if v365 != 0 {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	v367 = *(*int32)(unsafe.Add(mBase, _c_F_AbortSubTransaction[3]))
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v367)+48))
	if v368 != 0 {
		goto L84
	} else {
		goto L85
	}
L82:
	;
	goto L83
L83:
	;
	v378 = *(*int32)(unsafe.Add(mBase, _c_F_AbortSubTransaction[39]))
	if v378 != 0 {
		goto L88
	} else {
		goto L89
	}
L84:
	;
	F_pfree(m, v368)
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
		goto L1
	} else {
		goto L87
	}
L85:
	;
	goto L86
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v367)+56)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v367)+48)) = int64(0)
	goto L83
L87:
	;
	goto L86
L88:
	;
	v379 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
	v380 = *(*int32)(unsafe.Add(mBase, uint32(v24)+80))
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v380)+8))
	v383 = v378
	goto L91
L89:
	;
	goto L90
L90:
	;
	v402 = *(*int32)(unsafe.Add(mBase, uint32(v24)+40))
	v404 = int32(0)
	F_ResourceOwnerReleaseInternal(m, v402, int32(1), v404, v404)
	mBase = m.M
	v407 = m.ExcPending
	if v407 != 0 {
		goto L1
	} else {
		goto L95
	}
L91:
	;
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v383)))
	v391 = *(*int32)(unsafe.Add(mBase, uint32(v383)+8))
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v383)+4))
	m.T0[v392].(func(*base.Module, int32, int32, int32, int32))(m, int32(2), v379, v381, v391)
	mBase = m.M
	v394 = m.ExcPending
	if v394 != 0 {
		goto L1
	} else {
		goto L93
	}
L92:
	;
	goto L90
L93:
	;
	if v389 != 0 {
		v383 = v389
		goto L91
	} else {
		goto L94
	}
L94:
	;
	goto L92
L95:
	;
	F_AtEOXact_Aio(m)
	mBase = m.M
	v409 = m.ExcPending
	if v409 != 0 {
		goto L1
	} else {
		goto L96
	}
L96:
	;
	v411 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
	v412 = *(*int32)(unsafe.Add(mBase, uint32(v24)+80))
	v413 = *(*int32)(unsafe.Add(mBase, uint32(v412)+8))
	F_AtEOSubXact_RelationCache(m, int32(0), v411, v413)
	mBase = m.M
	v415 = m.ExcPending
	if v415 != 0 {
		goto L1
	} else {
		goto L97
	}
L97:
	;
	F_AtEOXact_TypeCache(m)
	mBase = m.M
	v417 = m.ExcPending
	if v417 != 0 {
		goto L1
	} else {
		goto L98
	}
L98:
	;
	F_AtEOSubXact_Inval(m, int32(0))
	mBase = m.M
	v420 = m.ExcPending
	if v420 != 0 {
		goto L1
	} else {
		goto L99
	}
L99:
	;
	v421 = *(*int32)(unsafe.Add(mBase, uint32(v24)+40))
	v423 = int32(0)
	F_ResourceOwnerReleaseInternal(m, v421, int32(2), v423, v423)
	mBase = m.M
	v426 = m.ExcPending
	if v426 != 0 {
		goto L1
	} else {
		goto L100
	}
L100:
	;
	v427 = *(*int32)(unsafe.Add(mBase, uint32(v24)+40))
	v429 = int32(0)
	F_ResourceOwnerReleaseInternal(m, v427, int32(3), v429, v429)
	mBase = m.M
	v432 = m.ExcPending
	if v432 != 0 {
		goto L1
	} else {
		goto L101
	}
L101:
	;
	F_smgrDoPendingDeletes(m, int32(0))
	mBase = m.M
	v435 = m.ExcPending
	if v435 != 0 {
		goto L1
	} else {
		goto L102
	}
L102:
	;
	v437 = *(*int32)(unsafe.Add(mBase, uint32(v24)+32))
	F_AtEOXact_GUC(m, int32(0), v437)
	mBase = m.M
	v439 = m.ExcPending
	if v439 != 0 {
		goto L1
	} else {
		goto L103
	}
L103:
	;
	v441 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
	F_AtEOSubXact_SPI(m, int32(0), v441)
	mBase = m.M
	v443 = m.ExcPending
	if v443 != 0 {
		goto L1
	} else {
		goto L104
	}
L104:
	;
	v445 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
	v446 = *(*int32)(unsafe.Add(mBase, uint32(v24)+80))
	v447 = *(*int32)(unsafe.Add(mBase, uint32(v446)+8))
	F_AtEOSubXact_on_commit_actions(m, int32(0), v445, v447)
	mBase = m.M
	v449 = m.ExcPending
	if v449 != 0 {
		goto L1
	} else {
		goto L105
	}
L105:
	;
	v451 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
	v455 = *(*int32)(unsafe.Add(mBase, _c_F_AbortSubTransaction[40]))
	if v451 == v455 {
		goto L107
	} else {
		goto L108
	}
L106:
	;
	v479 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
	v480 = *(*int32)(unsafe.Add(mBase, uint32(v24)+80))
	v481 = *(*int32)(unsafe.Add(mBase, uint32(v480)+8))
	F_AtEOSubXact_Files(m, int32(0), v479, v481)
	mBase = m.M
	v483 = m.ExcPending
	if v483 != 0 {
		goto L1
	} else {
		goto L113
	}
L107:
	;
	goto L111
L108:
	;
	goto L109
L109:
	;
	goto L106
L111:
	;
	goto L112
L112:
	;
	v460 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_AbortSubTransaction[41])) = uint8(v460)
	v463 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_AbortSubTransaction[42])) = v463
	*(*int32)(unsafe.Add(mBase, _c_F_AbortSubTransaction[40])) = v463
	*(*int32)(unsafe.Add(mBase, _c_F_AbortSubTransaction[43])) = v463
	*(*uint8)(unsafe.Add(mBase, _c_F_AbortSubTransaction[44])) = uint8(v463)
	v475 = *(*int32)(unsafe.Add(mBase, _c_F_AbortSubTransaction[45]))
	*(*int32)(unsafe.Add(mBase, uint32(v475)+68)) = v463
	goto L109
L113:
	;
	v485 = *(*int32)(unsafe.Add(mBase, uint32(v24)+28))
	F_AtEOSubXact_HashTables(m, int32(0), v485)
	mBase = m.M
	v487 = m.ExcPending
	if v487 != 0 {
		goto L1
	} else {
		goto L114
	}
L114:
	;
	v489 = *(*int32)(unsafe.Add(mBase, uint32(v24)+28))
	F_AtEOSubXact_PgStat(m, int32(0), v489)
	mBase = m.M
	v491 = m.ExcPending
	if v491 != 0 {
		goto L1
	} else {
		goto L115
	}
L115:
	;
	v492 = *(*int32)(unsafe.Add(mBase, uint32(v24)+28))
	v495 = *(*int32)(unsafe.Add(mBase, _c_F_AbortSubTransaction[46]))
	if v495 != 0 {
		goto L117
	} else {
		goto L118
	}
L116:
	;
	goto L62
L117:
	;
	v498 = v495
	goto L120
L118:
	;
	goto L119
L119:
	;
	v531 = *(*int32)(unsafe.Add(mBase, _c_F_AbortSubTransaction[47]))
	if v531 != 0 {
		goto L129
	} else {
		goto L130
	}
L120:
	;
	v503 = *(*int32)(unsafe.Add(mBase, uint32(v498)+4))
	if v503 < v492 {
		goto L116
	} else {
		goto L122
	}
L121:
	;
	goto L119
L122:
	;
	v505 = *(*int32)(unsafe.Add(mBase, uint32(v498)+8))
	v506 = *(*int32)(unsafe.Add(mBase, uint32(v498)))
	v507 = *(*int32)(unsafe.Add(mBase, uint32(v506)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v506)+44)) = v507 - int32(1)
	v511 = *(*int32)(unsafe.Add(mBase, uint32(v498)))
	v512 = *(*int32)(unsafe.Add(mBase, uint32(v511)+44))
	if v512 != 0 {
		v518 = v498
		goto L123
	} else {
		goto L124
	}
L123:
	;
	F_pfree(m, v518)
	mBase = m.M
	v520 = m.ExcPending
	if v520 != 0 {
		goto L1
	} else {
		goto L127
	}
L124:
	;
	v513 = *(*int32)(unsafe.Add(mBase, uint32(v511)+48))
	if v513 != 0 {
		v518 = v498
		goto L123
	} else {
		goto L125
	}
L125:
	;
	F_pfree(m, v511)
	mBase = m.M
	v515 = m.ExcPending
	if v515 != 0 {
		goto L1
	} else {
		goto L126
	}
L126:
	;
	v517 = *(*int32)(unsafe.Add(mBase, _c_F_AbortSubTransaction[46]))
	v518 = v517
	goto L123
L127:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_AbortSubTransaction[46])) = v505
	if v505 != 0 {
		v498 = v505
		goto L120
	} else {
		goto L128
	}
L128:
	;
	goto L121
L129:
	;
	v533 = *(*int32)(unsafe.Add(mBase, _c_F_AbortSubTransaction[45]))
	v534 = *(*int32)(unsafe.Add(mBase, uint32(v533)+40))
	v536 = *(*int32)(unsafe.Add(mBase, _c_F_AbortSubTransaction[47]))
	v538 = v536 - int32(48)
	v539 = *(*int32)(unsafe.Add(mBase, uint32(v538)))
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v539))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v534)) == int32(0) {
		goto L133
	} else {
		goto L134
	}
L130:
	;
	v556 = int32(0)
	goto L131
L131:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_AbortSubTransaction[48])) = v556
	v560 = *(*int32)(unsafe.Add(mBase, _c_F_AbortSubTransaction[45]))
	*(*int32)(unsafe.Add(mBase, uint32(v560)+40)) = v556
	goto L116
L132:
	;
	if v551 == int32(0) {
		goto L116
	} else {
		goto L136
	}
L133:
	;
	v551 = base.B2i32(base.Ui32(v534) < base.Ui32(v539))
	goto L132
L134:
	;
	goto L135
L135:
	;
	v551 = int32(base.Ui32(v534-v539) >> (uint(int32(31)) % 32))
	goto L132
L136:
	;
	v554 = *(*int32)(unsafe.Add(mBase, uint32(v538)))
	v556 = v554
	goto L131
}
func F_CleanupSubTransaction(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v19 int32
	_ = v19
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v173 int32
	_ = v173
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_CleanupSubTransaction[0]))
	v12 = int32(10)
	goto L3
L1:
	;
	if v49 != 0 {
		goto L14
	} else {
		goto L15
	}
L2:
	;
	goto L1
L3:
	;
	v19 = *(*int32)(unsafe.Add(mBase, _c_F_CleanupSubTransaction[1]))
	goto L6
L4:
	;
	v32 = int32(0)
	goto L11
L6:
	;
	goto L7
L7:
	;
	if int32(0)|base.B2i32(v19 == int32(15)) != 0 {
		goto L4
	} else {
		goto L9
	}
L9:
	;
	if v19 <= v12 {
		v49 = int32(1)
		goto L2
	} else {
		goto L10
	}
L10:
	;
	goto L4
L11:
	;
	v36 = *(*int32)(unsafe.Add(mBase, _c_F_CleanupSubTransaction[2]))
	if v36 != int32(2) {
		v49 = v32
		goto L2
	} else {
		goto L12
	}
L12:
	;
	v40 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CleanupSubTransaction[3])))
	if v40&int32(1) != 0 {
		v49 = v32
		goto L2
	} else {
		goto L13
	}
L13:
	;
	v46 = *(*int32)(unsafe.Add(mBase, _c_F_CleanupSubTransaction[4]))
	v49 = int32(0) | base.B2i32(v46 <= v12)
	goto L2
L14:
	;
	v53 = *(*int32)(unsafe.Add(mBase, _c_F_CleanupSubTransaction[0]))
	F_ShowTransactionStateRec(m, int32(_a_F_CleanupSubTransaction_0), v53)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	goto L16
L16:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v11)+20))
	if v56 == int32(4) {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	return
L18:
	;
	goto L16
L19:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	v84 = m.G0
	v86 = v84 - int32(32)
	m.G0 = v86
	v89 = v86 + int32(12)
	v91 = *(*int32)(unsafe.Add(mBase, _c_F_CleanupSubTransaction[5]))
	F_hash_seq_init(m, v89, v91)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L17
	} else {
		goto L28
	}
L20:
	;
	v61 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L17
	} else {
		goto L21
	}
L21:
	;
	if v61 == int32(0) {
		goto L19
	} else {
		goto L22
	}
L22:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v11)+20))
	if base.Ui32(v65) <= base.Ui32(int32(5)) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v65<<(uint(int32(2))%32))+uint32(_c_F_CleanupSubTransaction[6])))
	v72 = v70
	goto L25
L24:
	;
	v72 = int32(_a_F_CleanupSubTransaction_1)
	goto L25
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v72
	F_errmsg_internal(m, int32(_a_F_CleanupSubTransaction_2), v8)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L17
	} else {
		goto L26
	}
L26:
	;
	F_errfinish(m, int32(_a_F_CleanupSubTransaction_3), int32(_a_F_CleanupSubTransaction_4), int32(_a_F_CleanupSubTransaction_0))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L17
	} else {
		goto L27
	}
L27:
	;
	goto L19
L28:
	;
	v94 = F_hash_seq_search(m, v89)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L17
	} else {
		goto L29
	}
L29:
	;
	if v94 != 0 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v97 = v94
	goto L33
L31:
	;
	goto L32
L32:
	;
	m.G0 = v86 + int32(32)
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v11)+80))
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v142)+40))
	*(*int32)(unsafe.Add(mBase, _c_F_CleanupSubTransaction[7])) = v143
	*(*int32)(unsafe.Add(mBase, _c_F_CleanupSubTransaction[8])) = v143
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v11)+40))
	if v147 != 0 {
		goto L53
	} else {
		goto L54
	}
L33:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v97)+64))
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v101)+20))
	if v83 == v102 {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	goto L32
L35:
	;
	v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101)+84)))
	if v104 == int32(1) {
		goto L38
	} else {
		goto L39
	}
L36:
	;
	goto L37
L37:
	;
	v131 = F_hash_seq_search(m, v86+int32(12))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L17
	} else {
		goto L51
	}
L38:
	;
	v107 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v101)+84)) = uint8(v107)
	goto L40
L39:
	;
	goto L40
L40:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v101)+16))
	if v109 != 0 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v112 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L17
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	F_PortalDrop(m, v101, int32(0))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L17
	} else {
		goto L50
	}
L44:
	;
	if v112 != 0 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v101)))
	*(*int32)(unsafe.Add(mBase, uint32(v86))) = v114
	F_errmsg_internal(m, int32(_a_F_CleanupSubTransaction_5), v86)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L17
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v101)+16)) = int32(0)
	goto L43
L48:
	;
	F_errfinish(m, int32(_a_F_CleanupSubTransaction_6), int32(1120), int32(_a_F_CleanupSubTransaction_7))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L17
	} else {
		goto L49
	}
L49:
	;
	goto L47
L50:
	;
	goto L37
L51:
	;
	if v131 != 0 {
		v97 = v131
		goto L33
	} else {
		goto L52
	}
L52:
	;
	goto L34
L53:
	;
	F_ResourceOwnerDelete(m, v147)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L17
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+40)) = int32(0)
	v154 = *(*int32)(unsafe.Add(mBase, _c_F_CleanupSubTransaction[0]))
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v154)+44))
	*(*int32)(unsafe.Add(mBase, _c_F_CleanupSubTransaction[9])) = v155
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v154)+80))
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v158)+36))
	*(*int32)(unsafe.Add(mBase, _c_F_CleanupSubTransaction[10])) = v159
	v162 = *(*int32)(unsafe.Add(mBase, _c_F_CleanupSubTransaction[11]))
	if v162 != 0 {
		goto L57
	} else {
		goto L58
	}
L56:
	;
	goto L55
L57:
	;
	F_MemoryContextReset(m, v162)
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L17
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v154)+36))
	if v165 != 0 {
		goto L61
	} else {
		goto L62
	}
L60:
	;
	goto L59
L61:
	;
	F_MemoryContextDelete(m, v165)
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L17
	} else {
		goto L64
	}
L62:
	;
	goto L63
L63:
	;
	v168 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v154)+36)) = v168
	*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = v168
	F_PopTransaction(m)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L17
	} else {
		goto L65
	}
L64:
	;
	goto L63
L65:
	;
	m.G0 = v8 + int32(16)
	return
}
func F_StartSubTransaction(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v127 int64
	_ = v127
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v166 int32
	_ = v166
	var v173 int32
	_ = v173
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_StartSubTransaction[0]))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
	if v13 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = int32(1)
	v43 = *(*int32)(unsafe.Add(mBase, _c_F_StartSubTransaction[0]))
	v45 = *(*int32)(unsafe.Add(mBase, _c_F_StartSubTransaction[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v43)+44)) = v45
	v47 = int32(_a_F_StartSubTransaction_0)
	v49 = *(*int32)(unsafe.Add(mBase, _c_F_StartSubTransaction[2]))
	v54 = F_AllocSetContextCreateInternal(m, v49, int32(_a_F_StartSubTransaction_1), int32(0), int32(_a_F_StartSubTransaction_2), int32(_a_F_StartSubTransaction_3))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L3
	} else {
		goto L11
	}
L2:
	;
	v18 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return
L4:
	;
	if v18 == int32(0) {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
	if base.Ui32(v22) <= base.Ui32(int32(5)) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v22<<(uint(int32(2))%32))+uint32(_c_F_StartSubTransaction[3])))
	v29 = v27
	goto L8
L7:
	;
	v29 = int32(_a_F_StartSubTransaction_4)
	goto L8
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v29
	F_errmsg_internal(m, int32(_a_F_StartSubTransaction_5), v9)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L3
	} else {
		goto L9
	}
L9:
	;
	F_errfinish(m, int32(_a_F_StartSubTransaction_6), int32(_a_F_StartSubTransaction_7), int32(_a_F_StartSubTransaction_8))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L3
	} else {
		goto L10
	}
L10:
	;
	goto L1
L11:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_StartSubTransaction[2])) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v43)+36)) = v54
	*(*int32)(unsafe.Add(mBase, _c_F_StartSubTransaction[1])) = v54
	v61 = *(*int32)(unsafe.Add(mBase, _c_F_StartSubTransaction[0]))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)+80))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)+40))
	v65 = F_ResourceOwnerCreate(m, v63, int32(_a_F_StartSubTransaction_9))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L3
	} else {
		goto L12
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v61)+40)) = v65
	*(*int32)(unsafe.Add(mBase, _c_F_StartSubTransaction[4])) = v65
	*(*int32)(unsafe.Add(mBase, _c_F_StartSubTransaction[5])) = v65
	v73 = *(*int32)(unsafe.Add(mBase, _c_F_StartSubTransaction[0]))
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v73)+28))
	goto L14
L13:
	;
	v116 = v74 * int32(24)
	*(*int32)(unsafe.Add(mBase, uint32(v111+v116))) = int32(0)
	v120 = int32(_a_F_StartSubTransaction_10)
	v121 = *(*int32)(unsafe.Add(mBase, _c_F_StartSubTransaction[6]))
	v122 = v121 + v116
	v124 = *(*int32)(unsafe.Add(mBase, _c_F_StartSubTransaction[7]))
	*(*int32)(unsafe.Add(mBase, uint32(v122)+12)) = v124
	v127 = *(*int64)(unsafe.Add(mBase, _c_F_StartSubTransaction[8]))
	*(*int64)(unsafe.Add(mBase, uint32(v122)+4)) = v127
	v130 = *(*int32)(unsafe.Add(mBase, _c_F_StartSubTransaction[6]))
	v133 = *(*int32)(unsafe.Add(mBase, _c_F_StartSubTransaction[9]))
	*(*int32)(unsafe.Add(mBase, uint32(v130+v116)+16)) = v133
	v136 = *(*int32)(unsafe.Add(mBase, _c_F_StartSubTransaction[6]))
	v139 = *(*int32)(unsafe.Add(mBase, _c_F_StartSubTransaction[10]))
	*(*int32)(unsafe.Add(mBase, uint32(v136+v116)+20)) = v139
	*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = int32(2)
	v144 = *(*int32)(unsafe.Add(mBase, _c_F_StartSubTransaction[11]))
	if v144 != 0 {
		goto L27
	} else {
		goto L28
	}
L14:
	;
	v76 = *(*int32)(unsafe.Add(mBase, _c_F_StartSubTransaction[12]))
	if v76 <= v74 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v78 = v76
	goto L18
L16:
	;
	goto L17
L17:
	;
	v108 = *(*int32)(unsafe.Add(mBase, _c_F_StartSubTransaction[6]))
	v111 = v108
	goto L13
L18:
	;
	if v78 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_StartSubTransaction[12])) = v100
	*(*int32)(unsafe.Add(mBase, _c_F_StartSubTransaction[6])) = v101
	if v100 <= v74 {
		v78 = v100
		goto L18
	} else {
		goto L26
	}
L21:
	;
	v88 = *(*int32)(unsafe.Add(mBase, _c_F_StartSubTransaction[13]))
	v90 = F_MemoryContextAlloc(m, v88, int32(192))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L3
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v93 = *(*int32)(unsafe.Add(mBase, _c_F_StartSubTransaction[6]))
	v96 = F_repalloc(m, v93, v78*int32(48))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L3
	} else {
		goto L25
	}
L24:
	;
	v100 = int32(8)
	v101 = v90
	goto L20
L25:
	;
	v100 = v78 << (uint(int32(1)) % 32)
	v101 = v96
	goto L20
L26:
	;
	v111 = v101
	goto L13
L27:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v12)+80))
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v146)+8))
	v148 = v144
	goto L30
L28:
	;
	goto L29
L29:
	;
	v166 = int32(10)
	goto L36
L30:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v148)+8))
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v148)+4))
	m.T0[v157].(func(*base.Module, int32, int32, int32, int32))(m, int32(0), v145, v147, v156)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L3
	} else {
		goto L32
	}
L31:
	;
	goto L29
L32:
	;
	if v154 != 0 {
		v148 = v154
		goto L30
	} else {
		goto L33
	}
L33:
	;
	goto L31
L34:
	;
	if v203 != 0 {
		goto L47
	} else {
		goto L48
	}
L35:
	;
	goto L34
L36:
	;
	v173 = *(*int32)(unsafe.Add(mBase, _c_F_StartSubTransaction[14]))
	goto L39
L37:
	;
	v186 = int32(0)
	goto L44
L39:
	;
	goto L40
L40:
	;
	if int32(0)|base.B2i32(v173 == int32(15)) != 0 {
		goto L37
	} else {
		goto L42
	}
L42:
	;
	if v173 <= v166 {
		v203 = int32(1)
		goto L35
	} else {
		goto L43
	}
L43:
	;
	goto L37
L44:
	;
	v190 = *(*int32)(unsafe.Add(mBase, _c_F_StartSubTransaction[15]))
	if v190 != int32(2) {
		v203 = v186
		goto L35
	} else {
		goto L45
	}
L45:
	;
	v194 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartSubTransaction[16])))
	if v194&int32(1) != 0 {
		v203 = v186
		goto L35
	} else {
		goto L46
	}
L46:
	;
	v200 = *(*int32)(unsafe.Add(mBase, _c_F_StartSubTransaction[17]))
	v203 = int32(0) | base.B2i32(v200 <= v166)
	goto L35
L47:
	;
	v207 = *(*int32)(unsafe.Add(mBase, _c_F_StartSubTransaction[0]))
	F_ShowTransactionStateRec(m, int32(_a_F_StartSubTransaction_8), v207)
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L3
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	m.G0 = v9 + int32(16)
	return
L50:
	;
	goto L49
}
func F_SubTransGetTopmostTransaction(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
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
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	v2 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	if l0 == v2 {
		v111 = v2
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v8 + int32(16)
	return v111
L2:
	;
	v12 = l0
	goto L4
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L14
	} else {
		goto L28
	}
L4:
	;
	v18 = *(*int32)(unsafe.Add(mBase, _c_F_SubTransGetTopmostTransaction[0]))
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v18))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v12)) == int32(0) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	v79 = int32(0)
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v12))&int32(0) == v79 {
		goto L24
	} else {
		goto L25
	}
L6:
	;
	if v30 != 0 {
		v111 = v12
		goto L1
	} else {
		goto L10
	}
L7:
	;
	v30 = base.B2i32(base.Ui32(v12) < base.Ui32(v18))
	goto L6
L8:
	;
	goto L9
L9:
	;
	v30 = int32(base.Ui32(v12-v18) >> (uint(int32(31)) % 32))
	goto L6
L10:
	;
	if base.Ui32(int32(3)) <= base.Ui32(v12) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v35 = int32(base.Ui32(v12) >> (uint(int32(11)) % 32))
	v37 = F_SimpleLruReadPage_ReadOnly(m, int32(_a_F_SubTransGetTopmostTransaction_0), base.I64_extend_i32_u(v35), v12)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	goto L13
L13:
	;
	goto L5
L14:
	;
	return int32(0)
L15:
	;
	v42 = *(*int32)(unsafe.Add(mBase, _c_F_SubTransGetTopmostTransaction[1]))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
	v44 = int32(2)
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v43+v37<<(uint(v44)%32))))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v47+v12&int32(2047)<<(uint(v44)%32))))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v42)+28))
	v56 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_SubTransGetTopmostTransaction[2])))
	v57 = base.I32_rem_u_s(v35, v56)
	F_LWLockRelease(m, v54+v57<<(uint(int32(7))%32))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v12))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v53)) == int32(0) {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	if v74 == int32(0) {
		v93 = v53
		goto L3
	} else {
		goto L21
	}
L18:
	;
	v74 = base.B2i32(base.Ui32(v53) < base.Ui32(v12))
	goto L17
L19:
	;
	goto L20
L20:
	;
	v74 = int32(base.Ui32(v53-v12) >> (uint(int32(31)) % 32))
	goto L17
L21:
	;
	if v53 == int32(0) {
		v111 = v12
		goto L1
	} else {
		goto L22
	}
L22:
	;
	v12 = v53
	goto L4
L23:
	;
	if v92 != 0 {
		v111 = v12
		goto L1
	} else {
		goto L27
	}
L24:
	;
	v92 = base.B2i32(base.Ui32(v79) < base.Ui32(v12))
	goto L23
L25:
	;
	goto L26
L26:
	;
	v92 = int32(base.Ui32(v79-v12) >> (uint(int32(31)) % 32))
	goto L23
L27:
	;
	v93 = v79
	goto L3
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v12
	F_errmsg_internal(m, int32(_a_F_SubTransGetTopmostTransaction_1), v8)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L14
	} else {
		goto L29
	}
L29:
	;
	F_errfinish(m, int32(_a_F_SubTransGetTopmostTransaction_2), int32(185), int32(_a_F_SubTransGetTopmostTransaction_3))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L14
	} else {
		goto L30
	}
L30:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_SubTransPagePrecedes(m *base.Module, l0 int64, l1 int64) int32 {
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	v4 = int32(11)
	v9 = int32(4)
	v10 = base.I32_wrap_i64(l0)<<(uint(v4)%32) | v9
	v12 = base.I32_wrap_i64(l1) << (uint(v4) % 32)
	v15 = F_TransactionIdPrecedes(m, v10, v12|v9)
	if v15 != 0 {
		v17 = F_TransactionIdPrecedes(m, v10, v12+int32(2051))
		v19 = v17
	} else {
		v19 = int32(0)
	}
	return v19
}
func F_SubTransSetParent(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
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
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	v6 = *(*int32)(unsafe.Add(mBase, _c_F_SubTransSetParent[0]))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+28))
	v9 = int32(base.Ui32(l0) >> (uint(int32(11)) % 32))
	v11 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_SubTransSetParent[1])))
	v12 = base.I32_rem_u_s(v9, v11)
	v15 = v7 + v12<<(uint(int32(7))%32)
	v17 = F_LWLockAcquire(m, v15, int32(0))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return
	} else {
		v22 = F_SimpleLruReadPage(m, int32(_a_F_SubTransSetParent_0), base.I64_extend_i32_u(v9), int32(1), l0)
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return
		} else {
			v25 = *(*int32)(unsafe.Add(mBase, _c_F_SubTransSetParent[0]))
			v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
			v27 = int32(2)
			v30 = *(*int32)(unsafe.Add(mBase, uint32(v26+v22<<(uint(v27)%32))))
			v35 = v30 + l0&int32(2047)<<(uint(v27)%32)
			v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
			if l1 != v36 {
				*(*int32)(unsafe.Add(mBase, uint32(v35))) = l1
				v40 = *(*int32)(unsafe.Add(mBase, _c_F_SubTransSetParent[0]))
				v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+12))
				v43 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(v41+v22))) = uint8(v43)
			} else {
			}
			F_LWLockRelease(m, v15)
			mBase = m.M
			v46 = m.ExcPending
			if v46 != 0 {
				return
			} else {
				return
			}
		}
	}
}
