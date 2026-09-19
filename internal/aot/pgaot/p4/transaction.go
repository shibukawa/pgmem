package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_AbortTransaction(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v174 int32
	_ = v174
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v248 int32
	_ = v248
	var v254 int32
	_ = v254
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v303 int32
	_ = v303
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v339 int32
	_ = v339
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v360 int32
	_ = v360
	var v364 int32
	_ = v364
	var v366 int32
	_ = v366
	var v378 int32
	_ = v378
	var v385 int32
	_ = v385
	var v387 int32
	_ = v387
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v394 int64
	_ = v394
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v403 int32
	_ = v403
	var v405 int32
	_ = v405
	var v410 int32
	_ = v410
	var v419 int32
	_ = v419
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v424 int32
	_ = v424
	var v427 int32
	_ = v427
	var v436 int32
	_ = v436
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v441 int32
	_ = v441
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v457 int32
	_ = v457
	var v459 int32
	_ = v459
	var v462 int32
	_ = v462
	var v464 int32
	_ = v464
	var v467 int32
	_ = v467
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v475 int32
	_ = v475
	var v478 int32
	_ = v478
	var v480 int32
	_ = v480
	var v487 int32
	_ = v487
	var v499 int32
	_ = v499
	var v504 int32
	_ = v504
	var v506 int32
	_ = v506
	var v511 int32
	_ = v511
	var v514 int32
	_ = v514
	var v518 int32
	_ = v518
	var v521 int32
	_ = v521
	var v523 int32
	_ = v523
	var v530 int32
	_ = v530
	var v533 int32
	_ = v533
	var v535 int32
	_ = v535
	var v537 int32
	_ = v537
	var v540 int32
	_ = v540
	var v542 int32
	_ = v542
	var v555 int32
	_ = v555
	var v558 int32
	_ = v558
	var v577 int32
	_ = v577
	var v581 int32
	_ = v581
	var v585 int32
	_ = v585
	var v589 int32
	_ = v589
	var v592 int32
	_ = v592
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v598 int32
	_ = v598
	var v609 int32
	_ = v609
	var v623 int32
	_ = v623
	var v625 int32
	_ = v625
	v1 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v14 = int32(_a_F_AbortTransaction_0)
	v16 = *(*int32)(unsafe.Add(mBase, _c_F_AbortTransaction[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_AbortTransaction[0])) = v16 + int32(1)
	v21 = *(*int32)(unsafe.Add(mBase, _c_F_AbortTransaction[1]))
	v23 = *(*int32)(unsafe.Add(mBase, _c_F_AbortTransaction[2]))
	if v1 < v23 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_disable_timeout(m, int32(8))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v31 = *(*int32)(unsafe.Add(mBase, _c_F_AbortTransaction[3]))
	*(*int32)(unsafe.Add(mBase, _c_F_AbortTransaction[4])) = v31
	v35 = *(*int32)(unsafe.Add(mBase, _c_F_AbortTransaction[5]))
	v37 = *(*int32)(unsafe.Add(mBase, _c_F_AbortTransaction[6]))
	if v35 != 0 {
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
	v38 = v35
	goto L8
L7:
	;
	v38 = v37
	goto L8
L8:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_AbortTransaction[7])) = v38
	F_LWLockReleaseAll(m)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L4
	} else {
		goto L9
	}
L9:
	;
	v43 = *(*int32)(unsafe.Add(mBase, _c_F_AbortTransaction[8]))
	v44 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v43))) = v44
	v48 = *(*int32)(unsafe.Add(mBase, _c_F_AbortTransaction[9]))
	if v48 == v44 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	F_pgaio_error_cleanup(m)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L4
	} else {
		goto L15
	}
L11:
	;
	goto L10
L12:
	;
	v52 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_AbortTransaction[10])))
	if v52&int32(1) == int32(0) {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v48)+220))
	if v57 == int32(0) {
		goto L11
	} else {
		goto L14
	}
L14:
	;
	v60 = int32(_a_F_AbortTransaction_1)
	v62 = *(*int32)(unsafe.Add(mBase, _c_F_AbortTransaction[11]))
	v63 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_AbortTransaction[11])) = v62 + v63
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
	*(*int32)(unsafe.Add(mBase, uint32(v48))) = v66 + v63
	v70 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+220)) = v70
	*(*int32)(unsafe.Add(mBase, uint32(v48)+224)) = v70
	*(*int32)(unsafe.Add(mBase, uint32(v48))) = v66 + int32(2)
	v80 = *(*int32)(unsafe.Add(mBase, _c_F_AbortTransaction[11]))
	*(*int32)(unsafe.Add(mBase, _c_F_AbortTransaction[11])) = v80 - v63
	goto L11
L15:
	;
	F_UnlockBuffers(m)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L4
	} else {
		goto L16
	}
L16:
	;
	v88 = int32(0)
	v95 = *(*int32)(unsafe.Add(mBase, _c_F_AbortTransaction[12]))
	if v95 <= v88 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	F_ConditionVariableCancelSleep(m)
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L4
	} else {
		goto L30
	}
L18:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_AbortTransaction[13])) = int64(0)
	*(*int32)(unsafe.Add(mBase, _c_F_AbortTransaction[14])) = int32(_a_F_AbortTransaction_2)
	v174 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_AbortTransaction[12])) = v174
	*(*int32)(unsafe.Add(mBase, _c_F_AbortTransaction[15])) = v174
	*(*uint8)(unsafe.Add(mBase, _c_F_AbortTransaction[16])) = uint8(v174)
	*(*uint8)(unsafe.Add(mBase, _c_F_AbortTransaction[17])) = uint8(v174)
	goto L17
L19:
	;
	v99 = v95 & int32(7)
	v101 = *(*int32)(unsafe.Add(mBase, _c_F_AbortTransaction[18]))
	if base.Ui32(int32(8)) <= base.Ui32(v95) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v107 = v88
	v111 = v88
	goto L23
L21:
	;
	v139 = v88
	goto L22
L22:
	;
	v145 = int32(0)
	v146 = v139
	goto L27
L23:
	;
	v114 = v101 + v107*int32(_a_F_AbortTransaction_3)
	v115 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v114)+uint32(_c_F_AbortTransaction[19]))) = uint8(v115)
	*(*uint8)(unsafe.Add(mBase, uint32(v114)+uint32(_c_F_AbortTransaction[20]))) = uint8(v115)
	*(*uint8)(unsafe.Add(mBase, uint32(v114)+uint32(_c_F_AbortTransaction[21]))) = uint8(v115)
	*(*uint8)(unsafe.Add(mBase, uint32(v114)+uint32(_c_F_AbortTransaction[22]))) = uint8(v115)
	*(*uint8)(unsafe.Add(mBase, uint32(v114)+uint32(_c_F_AbortTransaction[23]))) = uint8(v115)
	*(*uint8)(unsafe.Add(mBase, uint32(v114)+uint32(_c_F_AbortTransaction[24]))) = uint8(v115)
	*(*uint8)(unsafe.Add(mBase, uint32(v114)+uint32(_c_F_AbortTransaction[25]))) = uint8(v115)
	*(*uint8)(unsafe.Add(mBase, uint32(v114))) = uint8(v115)
	v131 = int32(8)
	v132 = v107 + v131
	v134 = v111 + v131
	if v134 != v95&int32(2147483640) {
		v107 = v132
		v111 = v134
		goto L23
	} else {
		goto L25
	}
L24:
	;
	if v99 == int32(0) {
		goto L18
	} else {
		goto L26
	}
L25:
	;
	goto L24
L26:
	;
	v139 = v132
	goto L22
L27:
	;
	v154 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v101+v146*int32(_a_F_AbortTransaction_3)))) = uint8(v154)
	v156 = int32(1)
	v159 = v145 + v156
	if v159 != v99 {
		v145 = v159
		v146 = v146 + v156
		goto L27
	} else {
		goto L29
	}
L28:
	;
	goto L18
L29:
	;
	goto L28
L30:
	;
	F_LockErrorCleanup(m)
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L4
	} else {
		goto L31
	}
L31:
	;
	F_reschedule_timeouts(m)
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L4
	} else {
		goto L32
	}
L32:
	;
	F_pgmem_sigprocmask(m, int32(_a_F_AbortTransaction_4), int32(0))
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L4
	} else {
		goto L33
	}
L33:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v21)+24))
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v21)+20))
	switch v196 - int32(2) {
	case 0, 3:
		goto L34
	default:
		goto L35
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+20)) = int32(4)
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v21)+60))
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v21)+64))
	*(*int32)(unsafe.Add(mBase, _c_F_AbortTransaction[26])) = v226
	*(*int32)(unsafe.Add(mBase, _c_F_AbortTransaction[27])) = v225
	goto L43
L35:
	;
	v201 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L4
	} else {
		goto L36
	}
L36:
	;
	if v201 == int32(0) {
		goto L34
	} else {
		goto L37
	}
L37:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v21)+20))
	if base.Ui32(v205) <= base.Ui32(int32(5)) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v205<<(uint(int32(2))%32))+uint32(_c_F_AbortTransaction[28])))
	v212 = v210
	goto L40
L39:
	;
	v212 = int32(_a_F_AbortTransaction_5)
	goto L40
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v212
	F_errmsg_internal(m, int32(_a_F_AbortTransaction_6), v12)
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L4
	} else {
		goto L41
	}
L41:
	;
	F_errfinish(m, int32(_a_F_AbortTransaction_7), int32(2877), int32(_a_F_AbortTransaction_8))
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L4
	} else {
		goto L42
	}
L42:
	;
	goto L34
L43:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v21)+28))
	v233 = *(*int32)(unsafe.Add(mBase, _c_F_AbortTransaction[29]))
	if v231 <= v233 {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	v248 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_AbortTransaction[30])) = uint8(v248)
	*(*int32)(unsafe.Add(mBase, _c_F_AbortTransaction[31])) = v248
	goto L48
L45:
	;
	v236 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_AbortTransaction[32])) = v236
	*(*int32)(unsafe.Add(mBase, _c_F_AbortTransaction[33])) = v236
	*(*int32)(unsafe.Add(mBase, _c_F_AbortTransaction[34])) = v236
	*(*int32)(unsafe.Add(mBase, _c_F_AbortTransaction[29])) = v236
	goto L47
L46:
	;
	goto L47
L47:
	;
	goto L44
L48:
	;
	v254 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_AbortTransaction[35])) = uint8(v254)
	*(*int32)(unsafe.Add(mBase, _c_F_AbortTransaction[36])) = v254
	F_AtEOXact_Parallel(m, v254)
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L4
	} else {
		goto L49
	}
L49:
	;
	v262 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+76)) = uint8(v262)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+72)) = v262
	F_AfterTriggerEndXact(m)
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L4
	} else {
		goto L50
	}
L50:
	;
	F_AtAbort_Portals(m)
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L4
	} else {
		goto L51
	}
L51:
	;
	v272 = base.B2i32(v195 == int32(5))
	F_smgrDoPendingSyncs(m, int32(0), v272)
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L4
	} else {
		goto L52
	}
L52:
	;
	F_AtEOXact_LargeObject(m, int32(0))
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L4
	} else {
		goto L53
	}
L53:
	;
	v279 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_AbortTransaction[37])))
	if v279 == int32(0) {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v378 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_AbortTransaction[38])) = v378
	*(*int32)(unsafe.Add(mBase, _c_F_AbortTransaction[39])) = v378
	F_AtEOXact_RelationMap(m, v378, v272)
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L4
	} else {
		goto L69
	}
L55:
	;
	v283 = *(*int32)(unsafe.Add(mBase, _c_F_AbortTransaction[40]))
	if v283 != 0 {
		goto L54
	} else {
		goto L56
	}
L56:
	;
	v285 = *(*int32)(unsafe.Add(mBase, _c_F_AbortTransaction[41]))
	v289 = F_LWLockAcquire(m, v285+int32(3456), int32(0))
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L4
	} else {
		goto L57
	}
L57:
	;
	v292 = *(*int32)(unsafe.Add(mBase, _c_F_AbortTransaction[42]))
	v294 = v292 + int32(56)
	v295 = int32(_a_F_AbortTransaction_9)
	v296 = *(*int32)(unsafe.Add(mBase, _c_F_AbortTransaction[43]))
	v297 = int32(5)
	*(*int32)(unsafe.Add(mBase, uint32(v294+v296<<(uint(v297)%32)))) = int32(-1)
	v303 = *(*int32)(unsafe.Add(mBase, _c_F_AbortTransaction[43]))
	*(*int32)(unsafe.Add(mBase, uint32(v292+v303<<(uint(v297)%32))+60)) = int32(0)
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v292)+40))
	v311 = *(*int32)(unsafe.Add(mBase, _c_F_AbortTransaction[43]))
	if v309 == v311 {
		goto L59
	} else {
		goto L60
	}
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v292+v343<<(uint(int32(5))%32)-int32(-64)))) = int32(-1)
	v360 = *(*int32)(unsafe.Add(mBase, _c_F_AbortTransaction[41]))
	F_LWLockRelease(m, v360+int32(3456))
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L4
	} else {
		goto L68
	}
L59:
	;
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v292+v309<<(uint(int32(5))%32)-int32(-64))))
	*(*int32)(unsafe.Add(mBase, uint32(v292)+40)) = v318
	v343 = v309
	goto L58
L60:
	;
	goto L61
L61:
	;
	v320 = v309
	goto L62
L62:
	;
	if v320 == int32(-1) {
		goto L64
	} else {
		goto L65
	}
L63:
	;
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v294+v311<<(uint(int32(5))%32))+8))
	*(*int32)(unsafe.Add(mBase, uint32(v333)+8)) = v339
	v342 = *(*int32)(unsafe.Add(mBase, _c_F_AbortTransaction[43]))
	v343 = v342
	goto L58
L64:
	;
	v343 = v311
	goto L58
L65:
	;
	goto L66
L66:
	;
	v333 = v294 + v320<<(uint(int32(5))%32)
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v333)+8))
	if v334 != v311 {
		v320 = v334
		goto L62
	} else {
		goto L67
	}
L67:
	;
	goto L63
L68:
	;
	v366 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_AbortTransaction[37])) = uint8(v366)
	goto L54
L69:
	;
	F_AtAbort_Twophase(m)
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L4
	} else {
		goto L70
	}
L70:
	;
	if v272 == int32(0) {
		goto L72
	} else {
		goto L73
	}
L71:
	;
	v399 = *(*int32)(unsafe.Add(mBase, _c_F_AbortTransaction[44]))
	F_ProcArrayEndTransaction(m, v399, v397)
	mBase = m.M
	v401 = m.ExcPending
	if v401 != 0 {
		goto L4
	} else {
		goto L77
	}
L72:
	;
	v391 = F_RecordTransactionAbort(m, int32(0))
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L4
	} else {
		goto L75
	}
L73:
	;
	goto L74
L74:
	;
	v394 = *(*int64)(unsafe.Add(mBase, _c_F_AbortTransaction[45]))
	F_XLogSetAsyncXactLSN(m, v394)
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L4
	} else {
		goto L76
	}
L75:
	;
	v397 = v391
	goto L71
L76:
	;
	v397 = v1
	goto L71
L77:
	;
	v403 = *(*int32)(unsafe.Add(mBase, _c_F_AbortTransaction[3]))
	if v403 != 0 {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v405 = *(*int32)(unsafe.Add(mBase, _c_F_AbortTransaction[46]))
	if v195 == int32(5) {
		goto L82
	} else {
		goto L83
	}
L79:
	;
	goto L80
L80:
	;
	v623 = int32(_a_F_AbortTransaction_0)
	v625 = *(*int32)(unsafe.Add(mBase, _c_F_AbortTransaction[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_AbortTransaction[0])) = v625 - int32(1)
	m.G0 = v12 + int32(16)
	return
L81:
	;
	v452 = *(*int32)(unsafe.Add(mBase, _c_F_AbortTransaction[3]))
	v453 = int32(1)
	F_ResourceOwnerReleaseInternal(m, v452, v453, int32(0), v453)
	mBase = m.M
	v457 = m.ExcPending
	if v457 != 0 {
		goto L4
	} else {
		goto L95
	}
L82:
	;
	if v405 == int32(0) {
		goto L81
	} else {
		goto L85
	}
L83:
	;
	goto L84
L84:
	;
	if v405 == int32(0) {
		goto L81
	} else {
		goto L90
	}
L85:
	;
	v410 = v405
	goto L86
L86:
	;
	v419 = *(*int32)(unsafe.Add(mBase, uint32(v410)))
	v421 = *(*int32)(unsafe.Add(mBase, uint32(v410)+8))
	v422 = *(*int32)(unsafe.Add(mBase, uint32(v410)+4))
	m.T0[v422].(func(*base.Module, int32, int32))(m, int32(3), v421)
	mBase = m.M
	v424 = m.ExcPending
	if v424 != 0 {
		goto L4
	} else {
		goto L88
	}
L87:
	;
	goto L81
L88:
	;
	if v419 != 0 {
		v410 = v419
		goto L86
	} else {
		goto L89
	}
L89:
	;
	goto L87
L90:
	;
	v427 = v405
	goto L91
L91:
	;
	v436 = *(*int32)(unsafe.Add(mBase, uint32(v427)))
	v438 = *(*int32)(unsafe.Add(mBase, uint32(v427)+8))
	v439 = *(*int32)(unsafe.Add(mBase, uint32(v427)+4))
	m.T0[v439].(func(*base.Module, int32, int32))(m, int32(2), v438)
	mBase = m.M
	v441 = m.ExcPending
	if v441 != 0 {
		goto L4
	} else {
		goto L93
	}
L92:
	;
	goto L81
L93:
	;
	if v436 != 0 {
		v427 = v436
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
	v459 = m.ExcPending
	if v459 != 0 {
		goto L4
	} else {
		goto L96
	}
L96:
	;
	F_AtEOXact_RelationCache(m, int32(0))
	mBase = m.M
	v462 = m.ExcPending
	if v462 != 0 {
		goto L4
	} else {
		goto L97
	}
L97:
	;
	F_AtEOXact_TypeCache(m)
	mBase = m.M
	v464 = m.ExcPending
	if v464 != 0 {
		goto L4
	} else {
		goto L98
	}
L98:
	;
	F_AtEOXact_Inval(m, int32(0))
	mBase = m.M
	v467 = m.ExcPending
	if v467 != 0 {
		goto L4
	} else {
		goto L99
	}
L99:
	;
	v469 = *(*int32)(unsafe.Add(mBase, _c_F_AbortTransaction[47]))
	v470 = int32(_a_F_AbortTransaction_9)
	v471 = *(*int32)(unsafe.Add(mBase, _c_F_AbortTransaction[43]))
	v472 = int32(2)
	v475 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v469+v471<<(uint(v472)%32)))) = v475
	v478 = *(*int32)(unsafe.Add(mBase, _c_F_AbortTransaction[48]))
	v480 = *(*int32)(unsafe.Add(mBase, _c_F_AbortTransaction[43]))
	*(*int32)(unsafe.Add(mBase, uint32(v478+v480<<(uint(v472)%32)))) = v475
	v487 = int32(_a_F_AbortTransaction_10)
	*(*int32)(unsafe.Add(mBase, _c_F_AbortTransaction[49])) = v487
	*(*int32)(unsafe.Add(mBase, _c_F_AbortTransaction[50])) = v487
	*(*int32)(unsafe.Add(mBase, _c_F_AbortTransaction[51])) = v475
	*(*int32)(unsafe.Add(mBase, _c_F_AbortTransaction[52])) = v475
	goto L100
L100:
	;
	v499 = *(*int32)(unsafe.Add(mBase, _c_F_AbortTransaction[3]))
	F_ResourceOwnerReleaseInternal(m, v499, int32(2), int32(0), int32(1))
	mBase = m.M
	v504 = m.ExcPending
	if v504 != 0 {
		goto L4
	} else {
		goto L101
	}
L101:
	;
	v506 = *(*int32)(unsafe.Add(mBase, _c_F_AbortTransaction[3]))
	F_ResourceOwnerReleaseInternal(m, v506, int32(3), int32(0), int32(1))
	mBase = m.M
	v511 = m.ExcPending
	if v511 != 0 {
		goto L4
	} else {
		goto L102
	}
L102:
	;
	F_smgrDoPendingDeletes(m, int32(0))
	mBase = m.M
	v514 = m.ExcPending
	if v514 != 0 {
		goto L4
	} else {
		goto L103
	}
L103:
	;
	F_AtEOXact_GUC(m, int32(0), int32(1))
	mBase = m.M
	v518 = m.ExcPending
	if v518 != 0 {
		goto L4
	} else {
		goto L104
	}
L104:
	;
	F_AtEOXact_SPI(m, int32(0))
	mBase = m.M
	v521 = m.ExcPending
	if v521 != 0 {
		goto L4
	} else {
		goto L105
	}
L105:
	;
	v523 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_AbortTransaction[53])) = v523
	*(*int32)(unsafe.Add(mBase, _c_F_AbortTransaction[54])) = v523
	goto L106
L106:
	;
	F_AtEOXact_on_commit_actions(m, int32(0))
	mBase = m.M
	v530 = m.ExcPending
	if v530 != 0 {
		goto L4
	} else {
		goto L107
	}
L107:
	;
	v533 = base.B2i32(v195 == int32(5))
	F_AtEOXact_Namespace(m, int32(0), v533)
	mBase = m.M
	v535 = m.ExcPending
	if v535 != 0 {
		goto L4
	} else {
		goto L108
	}
L108:
	;
	F_smgrdestroyall(m)
	mBase = m.M
	v537 = m.ExcPending
	if v537 != 0 {
		goto L4
	} else {
		goto L109
	}
L109:
	;
	F_AtEOXact_Files(m, int32(0))
	mBase = m.M
	v540 = m.ExcPending
	if v540 != 0 {
		goto L4
	} else {
		goto L110
	}
L110:
	;
	v542 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_AbortTransaction[55])) = v542
	*(*int32)(unsafe.Add(mBase, _c_F_AbortTransaction[56])) = v542
	*(*int32)(unsafe.Add(mBase, _c_F_AbortTransaction[57])) = v542
	*(*int32)(unsafe.Add(mBase, _c_F_AbortTransaction[58])) = v542
	goto L111
L111:
	;
	F_AtEOXact_HashTables(m, int32(0))
	mBase = m.M
	v555 = m.ExcPending
	if v555 != 0 {
		goto L4
	} else {
		goto L112
	}
L112:
	;
	F_AtEOXact_PgStat(m, int32(0), v533)
	mBase = m.M
	v558 = m.ExcPending
	if v558 != 0 {
		goto L4
	} else {
		goto L113
	}
L113:
	;
	goto L115
L114:
	;
	F_AtEOXact_LogicalRepWorkers(m, int32(0))
	mBase = m.M
	v581 = m.ExcPending
	if v581 != 0 {
		goto L4
	} else {
		goto L119
	}
L115:
	;
	v577 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_AbortTransaction[59])) = uint8(v577)
	goto L114
L119:
	;
	v585 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_AbortTransaction[10])))
	if v585 != int32(1) {
		goto L121
	} else {
		goto L122
	}
L120:
	;
	goto L80
L121:
	;
	goto L120
L122:
	;
	v589 = *(*int32)(unsafe.Add(mBase, _c_F_AbortTransaction[9]))
	if v589 == int32(0) {
		goto L121
	} else {
		goto L123
	}
L123:
	;
	v592 = int32(_a_F_AbortTransaction_1)
	v594 = *(*int32)(unsafe.Add(mBase, _c_F_AbortTransaction[11]))
	v595 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_AbortTransaction[11])) = v594 + v595
	v598 = *(*int32)(unsafe.Add(mBase, uint32(v589)))
	*(*int32)(unsafe.Add(mBase, uint32(v589))) = v598 + v595
	*(*int64)(unsafe.Add(mBase, uint32(v589)+24)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v589))) = v598 + int32(2)
	v609 = *(*int32)(unsafe.Add(mBase, _c_F_AbortTransaction[11]))
	*(*int32)(unsafe.Add(mBase, _c_F_AbortTransaction[11])) = v609 - v595
	goto L121
}
func F_AssignTransactionId(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v52 int32
	_ = v52
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v103 int32
	_ = v103
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v160 int32
	_ = v160
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v180 int64
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v222 int32
	_ = v222
	var v226 int32
	_ = v226
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v268 int32
	_ = v268
	var v272 int32
	_ = v272
	var v277 int32
	_ = v277
	var v281 int32
	_ = v281
	var v285 int32
	_ = v285
	var v290 int32
	_ = v290
	var v294 int32
	_ = v294
	var v298 int32
	_ = v298
	var v303 int32
	_ = v303
	var v309 int32
	_ = v309
	var v313 int32
	_ = v313
	var v318 int32
	_ = v318
	var v330 int32
	_ = v330
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v348 int32
	_ = v348
	var v360 int32
	_ = v360
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v367 int32
	_ = v367
	var v371 int32
	_ = v371
	var v377 int32
	_ = v377
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v384 int32
	_ = v384
	var v385 int64
	_ = v385
	var v387 int32
	_ = v387
	var v393 int64
	_ = v393
	var v394 int32
	_ = v394
	var v396 int32
	_ = v396
	var v403 int32
	_ = v403
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v410 int32
	_ = v410
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v416 int32
	_ = v416
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v421 int64
	_ = v421
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v431 int32
	_ = v431
	var v434 int64
	_ = v434
	var v435 int32
	_ = v435
	var v437 int32
	_ = v437
	var v441 int32
	_ = v441
	var v444 int32
	_ = v444
	var v446 int32
	_ = v446
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v461 int32
	_ = v461
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v470 int32
	_ = v470
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v475 int32
	_ = v475
	var v477 int64
	_ = v477
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v489 int32
	_ = v489
	var v492 int64
	_ = v492
	var v493 int32
	_ = v493
	var v495 int32
	_ = v495
	var v503 int32
	_ = v503
	var v505 int32
	_ = v505
	var v506 int64
	_ = v506
	var v507 int64
	_ = v507
	var v515 int64
	_ = v515
	var v519 int64
	_ = v519
	var v522 int64
	_ = v522
	var v527 int32
	_ = v527
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v554 int32
	_ = v554
	var v557 int32
	_ = v557
	var v559 int32
	_ = v559
	var v562 int32
	_ = v562
	var v570 int32
	_ = v570
	var v574 int32
	_ = v574
	var v582 int64
	_ = v582
	var v591 int32
	_ = v591
	var v593 int32
	_ = v593
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v599 int32
	_ = v599
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v606 int32
	_ = v606
	var v610 int32
	_ = v610
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v619 int32
	_ = v619
	var v622 int32
	_ = v622
	var v626 int32
	_ = v626
	var v632 int32
	_ = v632
	var v633 int32
	_ = v633
	var v635 int32
	_ = v635
	var v638 int32
	_ = v638
	var v639 int32
	_ = v639
	var v641 int32
	_ = v641
	var v643 int32
	_ = v643
	var v644 int32
	_ = v644
	var v646 int32
	_ = v646
	var v654 int32
	_ = v654
	var v656 int32
	_ = v656
	var v657 int32
	_ = v657
	var v666 int32
	_ = v666
	var v669 int32
	_ = v669
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	var v675 int64
	_ = v675
	var v679 int32
	_ = v679
	var v691 int64
	_ = v691
	var v696 int32
	_ = v696
	var v698 int64
	_ = v698
	var v700 int32
	_ = v700
	var v701 int64
	_ = v701
	var v702 int32
	_ = v702
	var v706 int32
	_ = v706
	var v707 int32
	_ = v707
	var v711 int32
	_ = v711
	var v714 int32
	_ = v714
	var v718 int32
	_ = v718
	var v721 int64
	_ = v721
	var v722 int32
	_ = v722
	var v724 int32
	_ = v724
	var v738 int32
	_ = v738
	var v741 int32
	_ = v741
	var v745 int32
	_ = v745
	var v750 int32
	_ = v750
	v16 = m.G0
	v18 = v16 - int32(16)
	m.G0 = v18
	v21 = *(*int32)(unsafe.Add(mBase, _c_F_AssignTransactionId[0]))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+72))
	if v22 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v738 = m.ExcPending
	if v738 != 0 {
		goto L8
	} else {
		goto L167
	}
L2:
	;
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+76)))
	if v23 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v25 = *(*int32)(unsafe.Add(mBase, _c_F_AssignTransactionId[1]))
	if int32(0) <= v25 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	if v28 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v120 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_AssignTransactionId[2])))
	v122 = *(*int32)(unsafe.Add(mBase, _c_F_AssignTransactionId[3]))
	v125 = m.G0
	v127 = v125 + int32(-64)
	m.G0 = v127
	v131 = *(*int32)(unsafe.Add(mBase, _c_F_AssignTransactionId[0]))
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v131)+72))
	if v132 != 0 {
		goto L30
	} else {
		goto L31
	}
L6:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	if v31 != 0 {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v35 = F_palloc(m, v32<<(uint(int32(2))%32))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	return
L9:
	;
	v38 = int32(0)
	v40 = v28
	goto L10
L10:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
	if v52 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	if v62 != 0 {
		goto L16
	} else {
		goto L17
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v35+v38<<(uint(int32(2))%32)))) = v40
	v60 = v38 + int32(1)
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v40)+80))
	if v61 != 0 {
		v38 = v60
		v40 = v61
		goto L10
	} else {
		goto L15
	}
L13:
	;
	v62 = v38
	goto L14
L14:
	;
	goto L11
L15:
	;
	v62 = v60
	goto L14
L16:
	;
	v65 = v62
	goto L19
L17:
	;
	goto L18
L18:
	;
	F_pfree(m, v35)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L8
	} else {
		goto L23
	}
L19:
	;
	v80 = v65 - int32(1)
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v35+v80<<(uint(int32(2))%32))))
	F_AssignTransactionId(m, v84)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L8
	} else {
		goto L21
	}
L20:
	;
	goto L18
L21:
	;
	if v80 != 0 {
		v65 = v80
		goto L19
	} else {
		goto L22
	}
L22:
	;
	goto L20
L23:
	;
	goto L5
L24:
	;
	m.G0 = v127 - int32(-64)
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v582
	if v28 == int32(0) {
		goto L144
	} else {
		goto L145
	}
L25:
	;
	v394 = m.G0
	v396 = v394 - int32(16)
	m.G0 = v396
	if v387&int32(_a_F_AssignTransactionId_0) != 0 {
		goto L101
	} else {
		goto L102
	}
L26:
	;
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v202))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v181)) == int32(0) {
		goto L84
	} else {
		goto L85
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v127)+16)) = v256
	F_errmsg(m, int32(_a_F_AssignTransactionId_1), v125+int32(-48))
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L8
	} else {
		goto L79
	}
L28:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L8
	} else {
		goto L76
	}
L29:
	;
	if v135&int32(1) == int32(0) {
		goto L33
	} else {
		goto L34
	}
L30:
	;
	v135 = int32(1)
	goto L32
L31:
	;
	v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v131)+76)))
	v135 = v134
	goto L32
L32:
	;
	goto L29
L33:
	;
	v141 = *(*int32)(unsafe.Add(mBase, _c_F_AssignTransactionId[4]))
	if v141 == int32(0) {
		goto L36
	} else {
		goto L37
	}
L34:
	;
	goto L35
L35:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L8
	} else {
		goto L73
	}
L36:
	;
	v145 = *(*int32)(unsafe.Add(mBase, _c_F_AssignTransactionId[5]))
	v146 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v145)+36)) = v146
	v149 = *(*int32)(unsafe.Add(mBase, _c_F_AssignTransactionId[6]))
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v149)+4))
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v145)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v150+v151<<(uint(int32(2))%32)))) = v146
	v582 = int64(1)
	goto L24
L37:
	;
	goto L38
L38:
	;
	v160 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_AssignTransactionId[7])))
	if v160 == int32(1) {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	if v170 != 0 {
		goto L28
	} else {
		goto L43
	}
L40:
	;
	v165 = *(*int32)(unsafe.Add(mBase, _c_F_AssignTransactionId[8]))
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v165)+316))
	v168 = base.B2i32(v166 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _c_F_AssignTransactionId[7])) = uint8(v168)
	v170 = v168
	goto L42
L41:
	;
	v170 = int32(0)
	goto L42
L42:
	;
	goto L39
L43:
	;
	v172 = *(*int32)(unsafe.Add(mBase, _c_F_AssignTransactionId[9]))
	v176 = F_LWLockAcquire(m, v172+int32(384), int32(0))
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L8
	} else {
		goto L44
	}
L44:
	;
	v179 = *(*int32)(unsafe.Add(mBase, _c_F_AssignTransactionId[10]))
	v180 = *(*int64)(unsafe.Add(mBase, uint32(v179)+8))
	v181 = base.I32_wrap_i64(v180)
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v179)+20))
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v182))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v181)) == int32(0) {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	if v194 == int32(0) {
		v387 = v181
		v393 = v180
		goto L25
	} else {
		goto L49
	}
L46:
	;
	v194 = base.B2i32(base.Ui32(v182) <= base.Ui32(v181))
	goto L45
L47:
	;
	goto L48
L48:
	;
	v194 = base.B2i32(int32(0) <= v181-v182)
	goto L45
L49:
	;
	v198 = *(*int32)(unsafe.Add(mBase, _c_F_AssignTransactionId[10]))
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v198)+36))
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v198)+32))
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v198)+28))
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v198)+24))
	v204 = *(*int32)(unsafe.Add(mBase, _c_F_AssignTransactionId[9]))
	F_LWLockRelease(m, v204+int32(384))
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L8
	} else {
		goto L50
	}
L50:
	;
	v210 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_AssignTransactionId[11])))
	v213 = int32(0)
	if base.B2i32(v210&int32(1) == v213)|v181&int32(_a_F_AssignTransactionId_2) == v213 {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v201))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v181)) == int32(0) {
		goto L62
	} else {
		goto L63
	}
L52:
	;
	v222 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_AssignTransactionId[11])))
	if v222 == int32(1) {
		goto L56
	} else {
		goto L57
	}
L53:
	;
	goto L54
L54:
	;
	if v210&int32(1) == int32(0) {
		goto L26
	} else {
		goto L60
	}
L55:
	;
	v237 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_AssignTransactionId[11])))
	if v237 != 0 {
		goto L51
	} else {
		goto L59
	}
L56:
	;
	v226 = *(*int32)(unsafe.Add(mBase, _c_F_AssignTransactionId[12]))
	*(*int32)(unsafe.Add(mBase, uint32(v226+int32(16)))) = int32(1)
	v233 = *(*int32)(unsafe.Add(mBase, _c_F_AssignTransactionId[13]))
	v235 = F_pgmem_kill(m, v233, int32(10))
	mBase = m.M
	goto L58
L57:
	;
	goto L58
L58:
	;
	goto L55
L59:
	;
	goto L26
L60:
	;
	goto L51
L61:
	;
	if v253 == int32(0) {
		goto L26
	} else {
		goto L65
	}
L62:
	;
	v253 = base.B2i32(base.Ui32(v201) <= base.Ui32(v181))
	goto L61
L63:
	;
	goto L64
L64:
	;
	v253 = base.B2i32(int32(0) <= v181-v201)
	goto L61
L65:
	;
	v256 = F_get_database_name(m, v199)
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L8
	} else {
		goto L66
	}
L66:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L8
	} else {
		goto L67
	}
L67:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L8
	} else {
		goto L68
	}
L68:
	;
	if v256 != 0 {
		goto L27
	} else {
		goto L69
	}
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v127))) = v199
	F_errmsg(m, int32(_a_F_AssignTransactionId_3), v127)
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L8
	} else {
		goto L70
	}
L70:
	;
	F_errhint(m, int32(_a_F_AssignTransactionId_4), int32(0))
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L8
	} else {
		goto L71
	}
L71:
	;
	F_errfinish(m, int32(_a_F_AssignTransactionId_5), int32(166), int32(_a_F_AssignTransactionId_6))
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L8
	} else {
		goto L72
	}
L72:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L73:
	;
	F_errmsg_internal(m, int32(_a_F_AssignTransactionId_7), int32(0))
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L8
	} else {
		goto L74
	}
L74:
	;
	F_errfinish(m, int32(_a_F_AssignTransactionId_5), int32(87), int32(_a_F_AssignTransactionId_6))
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L8
	} else {
		goto L75
	}
L75:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L76:
	;
	F_errmsg_internal(m, int32(_a_F_AssignTransactionId_8), int32(0))
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L8
	} else {
		goto L77
	}
L77:
	;
	F_errfinish(m, int32(_a_F_AssignTransactionId_5), int32(103), int32(_a_F_AssignTransactionId_6))
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L8
	} else {
		goto L78
	}
L78:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L79:
	;
	F_errhint(m, int32(_a_F_AssignTransactionId_4), int32(0))
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L8
	} else {
		goto L80
	}
L80:
	;
	F_errfinish(m, int32(_a_F_AssignTransactionId_5), int32(159), int32(_a_F_AssignTransactionId_6))
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L8
	} else {
		goto L81
	}
L81:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L82:
	;
	v377 = *(*int32)(unsafe.Add(mBase, _c_F_AssignTransactionId[9]))
	v381 = F_LWLockAcquire(m, v377+int32(384), int32(0))
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L8
	} else {
		goto L100
	}
L83:
	;
	if v330 == int32(0) {
		goto L82
	} else {
		goto L87
	}
L84:
	;
	v330 = base.B2i32(base.Ui32(v202) <= base.Ui32(v181))
	goto L83
L85:
	;
	goto L86
L86:
	;
	v330 = base.B2i32(int32(0) <= v181-v202)
	goto L83
L87:
	;
	v333 = F_get_database_name(m, v199)
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L8
	} else {
		goto L88
	}
L88:
	;
	v337 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L8
	} else {
		goto L89
	}
L89:
	;
	if v333 != 0 {
		goto L91
	} else {
		goto L92
	}
L90:
	;
	F_errhint(m, v363, int32(0))
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L8
	} else {
		goto L98
	}
L91:
	;
	if v337 == int32(0) {
		goto L82
	} else {
		goto L94
	}
L92:
	;
	goto L93
L93:
	;
	if v337 == int32(0) {
		goto L82
	} else {
		goto L96
	}
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v127)+48)) = v333
	*(*int32)(unsafe.Add(mBase, uint32(v127)+52)) = v200 - v181
	F_errmsg(m, int32(_a_F_AssignTransactionId_9), v125+int32(-16))
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L8
	} else {
		goto L95
	}
L95:
	;
	v363 = int32(_a_F_AssignTransactionId_10)
	v364 = int32(179)
	goto L90
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v127)+32)) = v199
	*(*int32)(unsafe.Add(mBase, uint32(v127)+36)) = v200 - v181
	F_errmsg(m, int32(_a_F_AssignTransactionId_11), v125+int32(-32))
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L8
	} else {
		goto L97
	}
L97:
	;
	v363 = int32(_a_F_AssignTransactionId_12)
	v364 = int32(186)
	goto L90
L98:
	;
	F_errfinish(m, int32(_a_F_AssignTransactionId_5), v364, int32(_a_F_AssignTransactionId_6))
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
		goto L8
	} else {
		goto L99
	}
L99:
	;
	goto L82
L100:
	;
	v384 = *(*int32)(unsafe.Add(mBase, _c_F_AssignTransactionId[10]))
	v385 = *(*int64)(unsafe.Add(mBase, uint32(v384)+8))
	v387 = base.I32_wrap_i64(v385)
	v393 = v385
	goto L25
L101:
	;
	v403 = base.B2i32(v387 != int32(3))
	goto L103
L102:
	;
	v403 = int32(0)
	goto L103
L103:
	;
	if v403 == int32(0) {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	v407 = *(*int32)(unsafe.Add(mBase, _c_F_AssignTransactionId[14]))
	v408 = *(*int32)(unsafe.Add(mBase, uint32(v407)+28))
	v410 = int32(base.Ui32(v387) >> (uint(int32(15)) % 32))
	v412 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_AssignTransactionId[15])))
	v413 = base.I32_rem_u_s(v410, v412)
	v416 = v408 + v413<<(uint(int32(7))%32)
	v418 = F_LWLockAcquire(m, v416, int32(0))
	mBase = m.M
	v419 = m.ExcPending
	if v419 != 0 {
		goto L8
	} else {
		goto L107
	}
L105:
	;
	goto L106
L106:
	;
	v441 = int32(16)
	m.G0 = v396 + v441
	v444 = m.G0
	v446 = v444 - v441
	m.G0 = v446
	v449 = *(*int32)(unsafe.Add(mBase, _c_F_AssignTransactionId[16]))
	v450 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v449)+24)))
	if v450 != int32(1) {
		goto L113
	} else {
		goto L114
	}
L107:
	;
	v421 = base.I64_extend_i32_u(v410)
	v422 = F_SimpleLruZeroPage(m, int32(_a_F_AssignTransactionId_13), v421)
	mBase = m.M
	v423 = m.ExcPending
	if v423 != 0 {
		goto L8
	} else {
		goto L108
	}
L108:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v396)+8)) = v421
	F_XLogBeginInsert(m)
	mBase = m.M
	v426 = m.ExcPending
	if v426 != 0 {
		goto L8
	} else {
		goto L109
	}
L109:
	;
	v427 = int32(8)
	F_XLogRegisterData(m, v396+v427, v427)
	mBase = m.M
	v431 = m.ExcPending
	if v431 != 0 {
		goto L8
	} else {
		goto L110
	}
L110:
	;
	v434 = F_XLogInsert(m, int32(3), int32(0))
	mBase = m.M
	v435 = m.ExcPending
	if v435 != 0 {
		goto L8
	} else {
		goto L111
	}
L111:
	;
	F_LWLockRelease(m, v416)
	mBase = m.M
	v437 = m.ExcPending
	if v437 != 0 {
		goto L8
	} else {
		goto L112
	}
L112:
	;
	goto L106
L113:
	;
	m.G0 = v446 + int32(16)
	F_ExtendSUBTRANS(m, v387)
	mBase = m.M
	v503 = m.ExcPending
	if v503 != 0 {
		goto L8
	} else {
		goto L128
	}
L114:
	;
	v456 = int32(819)
	v457 = base.I32_div_u_s(v387, v456)
	if v387-v457*v456 != 0 {
		goto L115
	} else {
		goto L116
	}
L115:
	;
	v461 = base.B2i32(v387 != int32(3))
	goto L117
L116:
	;
	v461 = int32(0)
	goto L117
L117:
	;
	if v461 != 0 {
		goto L113
	} else {
		goto L118
	}
L118:
	;
	v463 = *(*int32)(unsafe.Add(mBase, _c_F_AssignTransactionId[17]))
	v464 = *(*int32)(unsafe.Add(mBase, uint32(v463)+28))
	v466 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_AssignTransactionId[18])))
	v467 = base.I32_rem_u_s(v457, v466)
	v470 = v464 + v467<<(uint(int32(7))%32)
	v472 = F_LWLockAcquire(m, v470, int32(0))
	mBase = m.M
	v473 = m.ExcPending
	if v473 != 0 {
		goto L8
	} else {
		goto L119
	}
L119:
	;
	v475 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_AssignTransactionId[19])))
	v477 = base.I64_extend_i32_u(v457)
	v478 = F_SimpleLruZeroPage(m, int32(_a_F_AssignTransactionId_14), v477)
	mBase = m.M
	v479 = m.ExcPending
	if v479 != 0 {
		goto L8
	} else {
		goto L120
	}
L120:
	;
	if v475 == int32(0) {
		goto L121
	} else {
		goto L122
	}
L121:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v446)+8)) = v477
	F_XLogBeginInsert(m)
	mBase = m.M
	v484 = m.ExcPending
	if v484 != 0 {
		goto L8
	} else {
		goto L124
	}
L122:
	;
	goto L123
L123:
	;
	F_LWLockRelease(m, v470)
	mBase = m.M
	v495 = m.ExcPending
	if v495 != 0 {
		goto L8
	} else {
		goto L127
	}
L124:
	;
	v485 = int32(8)
	F_XLogRegisterData(m, v446+v485, v485)
	mBase = m.M
	v489 = m.ExcPending
	if v489 != 0 {
		goto L8
	} else {
		goto L125
	}
L125:
	;
	v492 = F_XLogInsert(m, int32(18), int32(0))
	mBase = m.M
	v493 = m.ExcPending
	if v493 != 0 {
		goto L8
	} else {
		goto L126
	}
L126:
	;
	goto L123
L127:
	;
	goto L113
L128:
	;
	v505 = *(*int32)(unsafe.Add(mBase, _c_F_AssignTransactionId[10]))
	v506 = *(*int64)(unsafe.Add(mBase, uint32(v505)+8))
	v507 = int64(1)
	v515 = v506 + v507
	if base.Ui32(base.I32_wrap_i64(v515)) < base.Ui32(int32(3)) {
		goto L129
	} else {
		goto L130
	}
L129:
	;
	v519 = v506 + (v507-v506)&int64(4294967295) + int64(2)
	goto L131
L130:
	;
	v519 = v515
	goto L131
L131:
	;
	if base.Ui64(int64(2)) < base.Ui64(v515) {
		goto L132
	} else {
		goto L133
	}
L132:
	;
	v522 = v519
	goto L134
L133:
	;
	v522 = v515
	goto L134
L134:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v505)+8)) = v522
	if base.B2i32(v28 != int32(0)) == int32(0) {
		goto L136
	} else {
		goto L137
	}
L135:
	;
	v570 = *(*int32)(unsafe.Add(mBase, _c_F_AssignTransactionId[9]))
	F_LWLockRelease(m, v570+int32(384))
	mBase = m.M
	v574 = m.ExcPending
	if v574 != 0 {
		goto L8
	} else {
		goto L142
	}
L136:
	;
	v527 = *(*int32)(unsafe.Add(mBase, _c_F_AssignTransactionId[5]))
	*(*int32)(unsafe.Add(mBase, uint32(v527)+36)) = v387
	v530 = *(*int32)(unsafe.Add(mBase, _c_F_AssignTransactionId[6]))
	v531 = *(*int32)(unsafe.Add(mBase, uint32(v530)+4))
	v532 = *(*int32)(unsafe.Add(mBase, uint32(v527)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v531+v532<<(uint(int32(2))%32)))) = v387
	goto L135
L137:
	;
	goto L138
L138:
	;
	v538 = *(*int32)(unsafe.Add(mBase, _c_F_AssignTransactionId[6]))
	v539 = *(*int32)(unsafe.Add(mBase, uint32(v538)+8))
	v541 = *(*int32)(unsafe.Add(mBase, _c_F_AssignTransactionId[5]))
	v542 = *(*int32)(unsafe.Add(mBase, uint32(v541)+48))
	v545 = v539 + v542<<(uint(int32(1))%32)
	v546 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v541)+276)))
	if base.Ui32(v546) <= base.Ui32(int32(63)) {
		goto L139
	} else {
		goto L140
	}
L139:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v541+v546<<(uint(int32(2))%32))+280)) = v387
	v554 = v546 + int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v545))) = uint8(v554)
	v557 = *(*int32)(unsafe.Add(mBase, _c_F_AssignTransactionId[5]))
	*(*uint8)(unsafe.Add(mBase, uint32(v557)+276)) = uint8(v554)
	goto L135
L140:
	;
	goto L141
L141:
	;
	v559 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v545)+1)) = uint8(v559)
	v562 = *(*int32)(unsafe.Add(mBase, _c_F_AssignTransactionId[5]))
	*(*uint8)(unsafe.Add(mBase, uint32(v562)+277)) = uint8(v559)
	goto L135
L142:
	;
	v582 = v393
	goto L24
L143:
	;
	v638 = int32(_a_F_AssignTransactionId_15)
	v639 = *(*int32)(unsafe.Add(mBase, _c_F_AssignTransactionId[20]))
	v641 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	*(*int32)(unsafe.Add(mBase, _c_F_AssignTransactionId[20])) = v641
	v643 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v644 = m.G0
	v646 = v644 - int32(16)
	m.G0 = v646
	*(*int32)(unsafe.Add(mBase, uint32(v646)+12)) = int32(17104896)
	*(*int64)(unsafe.Add(mBase, uint32(v646)+4)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v646))) = v643
	v654 = int32(0)
	v656 = F_LockAcquire(m, v646, int32(7), v654, v654)
	mBase = m.M
	v657 = m.ExcPending
	if v657 != 0 {
		goto L8
	} else {
		goto L154
	}
L144:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_AssignTransactionId[21])) = v582
	v591 = m.G0
	v593 = v591 - int32(16)
	m.G0 = v593
	v596 = *(*int32)(unsafe.Add(mBase, _c_F_AssignTransactionId[22]))
	if v596 != 0 {
		goto L147
	} else {
		goto L148
	}
L145:
	;
	goto L146
L146:
	;
	v632 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	v633 = *(*int32)(unsafe.Add(mBase, uint32(v632)))
	F_SubTransSetParent(m, base.I32_wrap_i64(v582), v633)
	mBase = m.M
	v635 = m.ExcPending
	if v635 != 0 {
		goto L8
	} else {
		goto L153
	}
L147:
	;
	v597 = base.I32_wrap_i64(v582)
	v599 = *(*int32)(unsafe.Add(mBase, _c_F_AssignTransactionId[9]))
	v603 = F_LWLockAcquire(m, v599+int32(3584), int32(0))
	mBase = m.M
	v604 = m.ExcPending
	if v604 != 0 {
		goto L8
	} else {
		goto L150
	}
L148:
	;
	goto L149
L149:
	;
	m.G0 = v593 + int32(16)
	goto L143
L150:
	;
	v606 = *(*int32)(unsafe.Add(mBase, _c_F_AssignTransactionId[22]))
	*(*int32)(unsafe.Add(mBase, uint32(v606)+96)) = v597
	*(*int32)(unsafe.Add(mBase, uint32(v593)+12)) = v597
	v610 = *(*int32)(unsafe.Add(mBase, _c_F_AssignTransactionId[23]))
	v616 = F_hash_search(m, v610, v593+int32(12), int32(1), v593+int32(11))
	mBase = m.M
	v617 = m.ExcPending
	if v617 != 0 {
		goto L8
	} else {
		goto L151
	}
L151:
	;
	v619 = *(*int32)(unsafe.Add(mBase, _c_F_AssignTransactionId[22]))
	*(*int32)(unsafe.Add(mBase, uint32(v616)+4)) = v619
	v622 = *(*int32)(unsafe.Add(mBase, _c_F_AssignTransactionId[9]))
	F_LWLockRelease(m, v622+int32(3584))
	mBase = m.M
	v626 = m.ExcPending
	if v626 != 0 {
		goto L8
	} else {
		goto L152
	}
L152:
	;
	goto L149
L153:
	;
	goto L143
L154:
	;
	m.G0 = v646 + int32(16)
	*(*int32)(unsafe.Add(mBase, _c_F_AssignTransactionId[20])) = v639
	if v28 == int32(0) {
		goto L155
	} else {
		goto L156
	}
L155:
	;
	m.G0 = v18 + int32(16)
	return
L156:
	;
	v666 = *(*int32)(unsafe.Add(mBase, _c_F_AssignTransactionId[3]))
	if v666 <= int32(0) {
		goto L155
	} else {
		goto L157
	}
L157:
	;
	v669 = int32(_a_F_AssignTransactionId_16)
	v670 = *(*int32)(unsafe.Add(mBase, _c_F_AssignTransactionId[24]))
	v671 = int32(2)
	v675 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*uint32)(unsafe.Add(mBase, uint32(v670<<(uint(v671)%32))+uint32(_c_F_AssignTransactionId[25]))) = uint32(v675)
	v679 = v670 + int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_AssignTransactionId[24])) = v679
	if (base.B2i32(v28 == int32(0))|base.B2i32(v122 < v671)|v120)&base.B2i32(v679 < int32(64)) != 0 {
		goto L155
	} else {
		goto L158
	}
L158:
	;
	v691 = int64(*(*uint32)(unsafe.Add(mBase, _c_F_AssignTransactionId[21])))
	if v691 == int64(0) {
		goto L159
	} else {
		goto L160
	}
L159:
	;
	F_AssignTransactionId(m, int32(_a_F_AssignTransactionId_17))
	mBase = m.M
	v696 = m.ExcPending
	if v696 != 0 {
		goto L8
	} else {
		goto L162
	}
L160:
	;
	v701 = v691
	v702 = v679
	goto L161
L161:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+12)) = v702
	*(*uint32)(unsafe.Add(mBase, uint32(v18)+8)) = uint32(v701)
	F_XLogBeginInsert(m)
	mBase = m.M
	v706 = m.ExcPending
	if v706 != 0 {
		goto L8
	} else {
		goto L163
	}
L162:
	;
	v698 = *(*int64)(unsafe.Add(mBase, _c_F_AssignTransactionId[21]))
	v700 = *(*int32)(unsafe.Add(mBase, _c_F_AssignTransactionId[24]))
	v701 = v698
	v702 = v700
	goto L161
L163:
	;
	v707 = int32(8)
	F_XLogRegisterData(m, v18+v707, v707)
	mBase = m.M
	v711 = m.ExcPending
	if v711 != 0 {
		goto L8
	} else {
		goto L164
	}
L164:
	;
	v714 = *(*int32)(unsafe.Add(mBase, _c_F_AssignTransactionId[24]))
	F_XLogRegisterData(m, int32(_a_F_AssignTransactionId_18), v714<<(uint(int32(2))%32))
	mBase = m.M
	v718 = m.ExcPending
	if v718 != 0 {
		goto L8
	} else {
		goto L165
	}
L165:
	;
	v721 = F_XLogInsert(m, int32(1), int32(80))
	mBase = m.M
	v722 = m.ExcPending
	if v722 != 0 {
		goto L8
	} else {
		goto L166
	}
L166:
	;
	v724 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_AssignTransactionId[2])) = uint8(v724)
	*(*int32)(unsafe.Add(mBase, _c_F_AssignTransactionId[24])) = int32(0)
	goto L155
L167:
	;
	F_errcode(m, int32(322))
	mBase = m.M
	v741 = m.ExcPending
	if v741 != 0 {
		goto L8
	} else {
		goto L168
	}
L168:
	;
	F_errmsg(m, int32(_a_F_AssignTransactionId_19), int32(0))
	mBase = m.M
	v745 = m.ExcPending
	if v745 != 0 {
		goto L8
	} else {
		goto L169
	}
L169:
	;
	F_errfinish(m, int32(_a_F_AssignTransactionId_20), int32(652), int32(_a_F_AssignTransactionId_21))
	mBase = m.M
	v750 = m.ExcPending
	if v750 != 0 {
		goto L8
	} else {
		goto L170
	}
L170:
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
	F_errfinish(m, int32(_a_F_CleanupTransaction_1), int32(901), int32(_a_F_CleanupTransaction_2))
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
	F_errfinish(m, int32(_a_F_CleanupTransaction_5), int32(3018), int32(_a_F_CleanupTransaction_6))
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
						F_errfinish(m, int32(_a_F_RequireTransactionBlock_1), int32(3749), int32(_a_F_RequireTransactionBlock_2))
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
	var v279 int32
	_ = v279
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v292 int32
	_ = v292
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v302 int32
	_ = v302
	var v306 int32
	_ = v306
	var v311 int32
	_ = v311
	var v319 int32
	_ = v319
	var v330 int32
	_ = v330
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v343 int32
	_ = v343
	var v356 int32
	_ = v356
	var v360 int32
	_ = v360
	var v364 int32
	_ = v364
	var v370 int32
	_ = v370
	var v373 int32
	_ = v373
	var v377 int32
	_ = v377
	var v379 int32
	_ = v379
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
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)+316))
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
	*(*int32)(unsafe.Add(mBase, uint32(v202)+56)) = v192
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
	v287 = m.G0
	v289 = v287 - int32(16)
	m.G0 = v289
	v292 = *(*int32)(unsafe.Add(mBase, _c_F_StartTransaction[50]))
	if v292 == int32(0) {
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
	*(*int64)(unsafe.Add(mBase, uint32(v259)+24)) = v252
	*(*int32)(unsafe.Add(mBase, uint32(v259))) = v268 + int32(2)
	v279 = *(*int32)(unsafe.Add(mBase, _c_F_StartTransaction[48]))
	*(*int32)(unsafe.Add(mBase, _c_F_StartTransaction[48])) = v279 - v265
	goto L41
L44:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_StartTransaction[50])) = int32(1)
	m.G0 = v289 + int32(16)
	F_ReceiveSharedInvalidMessages(m)
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L16
	} else {
		goto L50
	}
L45:
	;
	v297 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L16
	} else {
		goto L46
	}
L46:
	;
	if v297 == int32(0) {
		goto L44
	} else {
		goto L47
	}
L47:
	;
	v302 = *(*int32)(unsafe.Add(mBase, _c_F_StartTransaction[50]))
	*(*int32)(unsafe.Add(mBase, uint32(v289))) = v302
	F_errmsg_internal(m, int32(_a_F_StartTransaction_11), v289)
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L16
	} else {
		goto L48
	}
L48:
	;
	F_errfinish(m, int32(_a_F_StartTransaction_12), int32(2224), int32(_a_F_StartTransaction_13))
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
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
	v330 = *(*int32)(unsafe.Add(mBase, _c_F_StartTransaction[53]))
	if int32(0) < v330 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	F_enable_timeout_after(m, int32(8), v330)
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L16
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	v336 = int32(10)
	goto L57
L54:
	;
	goto L53
L55:
	;
	if v373 != 0 {
		goto L68
	} else {
		goto L69
	}
L56:
	;
	goto L55
L57:
	;
	v343 = *(*int32)(unsafe.Add(mBase, _c_F_StartTransaction[54]))
	goto L60
L58:
	;
	v356 = int32(0)
	goto L65
L60:
	;
	goto L61
L61:
	;
	if int32(0)|base.B2i32(v343 == int32(15)) != 0 {
		goto L58
	} else {
		goto L63
	}
L63:
	;
	if v343 <= v336 {
		v373 = int32(1)
		goto L56
	} else {
		goto L64
	}
L64:
	;
	goto L58
L65:
	;
	v360 = *(*int32)(unsafe.Add(mBase, _c_F_StartTransaction[55]))
	if v360 != int32(2) {
		v373 = v356
		goto L56
	} else {
		goto L66
	}
L66:
	;
	v364 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartTransaction[56])))
	if v364&int32(1) != 0 {
		v373 = v356
		goto L56
	} else {
		goto L67
	}
L67:
	;
	v370 = *(*int32)(unsafe.Add(mBase, _c_F_StartTransaction[57]))
	v373 = int32(0) | base.B2i32(v370 <= v336)
	goto L56
L68:
	;
	v377 = *(*int32)(unsafe.Add(mBase, _c_F_StartTransaction[0]))
	F_ShowTransactionStateRec(m, int32(_a_F_StartTransaction_10), v377)
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
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
								v49 = *(*int32)(unsafe.Add(mBase, uint32(v48)+316))
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
