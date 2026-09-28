package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_AbortSubTransaction(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v170 int32
	_ = v170
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v258 int32
	_ = v258
	var v263 int32
	_ = v263
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v290 int32
	_ = v290
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v328 int32
	_ = v328
	var v331 int32
	_ = v331
	var v334 int32
	_ = v334
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v343 int32
	_ = v343
	var v348 int32
	_ = v348
	var v351 int32
	_ = v351
	var v354 int32
	_ = v354
	var v356 int32
	_ = v356
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v384 int32
	_ = v384
	var v389 int32
	_ = v389
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v394 int32
	_ = v394
	var v401 int32
	_ = v401
	var v403 int32
	_ = v403
	var v406 int32
	_ = v406
	var v408 int32
	_ = v408
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v414 int32
	_ = v414
	var v416 int32
	_ = v416
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v422 int32
	_ = v422
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v428 int32
	_ = v428
	var v431 int32
	_ = v431
	var v434 int32
	_ = v434
	var v436 int32
	_ = v436
	var v438 int32
	_ = v438
	var v440 int32
	_ = v440
	var v442 int32
	_ = v442
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v448 int32
	_ = v448
	var v450 int32
	_ = v450
	var v454 int32
	_ = v454
	var v459 int32
	_ = v459
	var v462 int32
	_ = v462
	var v474 int32
	_ = v474
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v482 int32
	_ = v482
	var v484 int32
	_ = v484
	var v486 int32
	_ = v486
	var v488 int32
	_ = v488
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v493 int32
	_ = v493
	var v496 int32
	_ = v496
	var v500 int32
	_ = v500
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v512 int32
	_ = v512
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v517 int32
	_ = v517
	var v527 int32
	_ = v527
	var v531 int32
	_ = v531
	var v534 int32
	_ = v534
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v543 int32
	_ = v543
	var v546 int32
	_ = v546
	var v572 int32
	_ = v572
	var v574 int32
	_ = v574
	var v576 int32
	_ = v576
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v13 = *(*int32)(unsafe.Add(mBase, _c_F_AbortSubTransaction[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_AbortSubTransaction[1])) = v13
	v15 = int32(_a_F_AbortSubTransaction_0)
	v17 = *(*int32)(unsafe.Add(mBase, _c_F_AbortSubTransaction[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_AbortSubTransaction[2])) = v17 + int32(1)
	v23 = *(*int32)(unsafe.Add(mBase, _c_F_AbortSubTransaction[3]))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+40))
	*(*int32)(unsafe.Add(mBase, _c_F_AbortSubTransaction[4])) = v24
	F_LWLockReleaseAll(m)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	F_WaitLSNCleanup(m)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v31 = *(*int32)(unsafe.Add(mBase, _c_F_AbortSubTransaction[5]))
	v32 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v31))) = v32
	v36 = *(*int32)(unsafe.Add(mBase, _c_F_AbortSubTransaction[6]))
	if v36 == v32 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	F_pgaio_error_cleanup(m)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L1
	} else {
		goto L9
	}
L5:
	;
	goto L4
L6:
	;
	v40 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_AbortSubTransaction[7])))
	if v40&int32(1) == int32(0) {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v36)+220))
	if v45 == int32(0) {
		goto L5
	} else {
		goto L8
	}
L8:
	;
	v48 = int32(_a_F_AbortSubTransaction_1)
	v50 = *(*int32)(unsafe.Add(mBase, _c_F_AbortSubTransaction[8]))
	v51 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_AbortSubTransaction[8])) = v50 + v51
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
	*(*int32)(unsafe.Add(mBase, uint32(v36))) = v54 + v51
	v58 = int32(0)
	v60 = int32(_a_F_AbortSubTransaction_2)
	v61 = base.AtomicRmwOr32(m, v58, v60, v58)
	*(*int32)(unsafe.Add(mBase, uint32(v36)+220)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v36)+224)) = v58
	v69 = base.AtomicRmwOr32(m, v58, v60, v58)
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
	*(*int32)(unsafe.Add(mBase, uint32(v36))) = v70 + v51
	v76 = *(*int32)(unsafe.Add(mBase, _c_F_AbortSubTransaction[8]))
	*(*int32)(unsafe.Add(mBase, _c_F_AbortSubTransaction[8])) = v76 - v51
	goto L5
L9:
	;
	F_UnlockBuffers(m)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	v84 = int32(0)
	v91 = *(*int32)(unsafe.Add(mBase, _c_F_AbortSubTransaction[9]))
	if v91 <= v84 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	F_ConditionVariableCancelSleep(m)
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L1
	} else {
		goto L24
	}
L12:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_AbortSubTransaction[10])) = int64(0)
	*(*int32)(unsafe.Add(mBase, _c_F_AbortSubTransaction[11])) = int32(_a_F_AbortSubTransaction_3)
	v170 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_AbortSubTransaction[9])) = v170
	*(*int32)(unsafe.Add(mBase, _c_F_AbortSubTransaction[12])) = v170
	*(*uint8)(unsafe.Add(mBase, _c_F_AbortSubTransaction[13])) = uint8(v170)
	*(*uint8)(unsafe.Add(mBase, _c_F_AbortSubTransaction[14])) = uint8(v170)
	goto L11
L13:
	;
	v95 = v91 & int32(7)
	v97 = *(*int32)(unsafe.Add(mBase, _c_F_AbortSubTransaction[15]))
	if base.Ui32(int32(8)) <= base.Ui32(v91) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v103 = v84
	v107 = v84
	goto L17
L15:
	;
	v135 = v84
	goto L16
L16:
	;
	v141 = int32(0)
	v142 = v135
	goto L21
L17:
	;
	v110 = v97 + v103*int32(_a_F_AbortSubTransaction_4)
	v111 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v110)+uint32(_c_F_AbortSubTransaction[16]))) = uint8(v111)
	*(*uint8)(unsafe.Add(mBase, uint32(v110)+uint32(_c_F_AbortSubTransaction[17]))) = uint8(v111)
	*(*uint8)(unsafe.Add(mBase, uint32(v110)+uint32(_c_F_AbortSubTransaction[18]))) = uint8(v111)
	*(*uint8)(unsafe.Add(mBase, uint32(v110)+uint32(_c_F_AbortSubTransaction[19]))) = uint8(v111)
	*(*uint8)(unsafe.Add(mBase, uint32(v110)+uint32(_c_F_AbortSubTransaction[20]))) = uint8(v111)
	*(*uint8)(unsafe.Add(mBase, uint32(v110)+uint32(_c_F_AbortSubTransaction[21]))) = uint8(v111)
	*(*uint8)(unsafe.Add(mBase, uint32(v110)+uint32(_c_F_AbortSubTransaction[22]))) = uint8(v111)
	*(*uint8)(unsafe.Add(mBase, uint32(v110))) = uint8(v111)
	v127 = int32(8)
	v128 = v103 + v127
	v130 = v107 + v127
	if v130 != v91&int32(2147483640) {
		v103 = v128
		v107 = v130
		goto L17
	} else {
		goto L19
	}
L18:
	;
	if v95 == int32(0) {
		goto L12
	} else {
		goto L20
	}
L19:
	;
	goto L18
L20:
	;
	v135 = v128
	goto L16
L21:
	;
	v150 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v97+v142*int32(_a_F_AbortSubTransaction_4)))) = uint8(v150)
	v152 = int32(1)
	v155 = v141 + v152
	if v155 != v95 {
		v141 = v155
		v142 = v142 + v152
		goto L21
	} else {
		goto L23
	}
L22:
	;
	goto L12
L23:
	;
	goto L22
L24:
	;
	F_LockErrorCleanup(m)
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	F_reschedule_timeouts(m)
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	F_pgmem_sigprocmask(m, int32(_a_F_AbortSubTransaction_5), int32(0))
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	v191 = int32(10)
	goto L30
L28:
	;
	if v231 != 0 {
		goto L41
	} else {
		goto L42
	}
L29:
	;
	goto L28
L30:
	;
	v198 = *(*int32)(unsafe.Add(mBase, _c_F_AbortSubTransaction[23]))
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v198<<(uint(int32(2))%32))+uint32(_c_F_AbortSubTransaction[24])))
	goto L33
L31:
	;
	v214 = int32(0)
	goto L38
L33:
	;
	goto L34
L34:
	;
	if int32(0)|base.B2i32(v201 == int32(15)) != 0 {
		goto L31
	} else {
		goto L36
	}
L36:
	;
	if v201 <= v191 {
		v231 = int32(1)
		goto L29
	} else {
		goto L37
	}
L37:
	;
	goto L31
L38:
	;
	v218 = *(*int32)(unsafe.Add(mBase, _c_F_AbortSubTransaction[25]))
	if v218 != int32(2) {
		v231 = v214
		goto L29
	} else {
		goto L39
	}
L39:
	;
	v222 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_AbortSubTransaction[26])))
	if v222&int32(1) != 0 {
		v231 = v214
		goto L29
	} else {
		goto L40
	}
L40:
	;
	v228 = *(*int32)(unsafe.Add(mBase, _c_F_AbortSubTransaction[27]))
	v231 = int32(0) | base.B2i32(v228 <= v191)
	goto L29
L41:
	;
	v235 = *(*int32)(unsafe.Add(mBase, _c_F_AbortSubTransaction[3]))
	F_ShowTransactionStateRec(m, int32(_a_F_AbortSubTransaction_6), v235)
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L1
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v23)+20))
	if v238 == int32(2) {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	goto L43
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+20)) = int32(4)
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v23)+60))
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v23)+64))
	*(*int32)(unsafe.Add(mBase, _c_F_AbortSubTransaction[28])) = v268
	*(*int32)(unsafe.Add(mBase, _c_F_AbortSubTransaction[29])) = v267
	goto L54
L46:
	;
	v243 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	if v243 == int32(0) {
		goto L45
	} else {
		goto L48
	}
L48:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v23)+20))
	if base.Ui32(v247) <= base.Ui32(int32(5)) {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v247<<(uint(int32(2))%32))+uint32(_c_F_AbortSubTransaction[30])))
	v254 = v252
	goto L51
L50:
	;
	v254 = int32(_a_F_AbortSubTransaction_7)
	goto L51
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v254
	F_errmsg_internal(m, int32(_a_F_AbortSubTransaction_8), v9)
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	F_errfinish(m, int32(_a_F_AbortSubTransaction_9), int32(_a_F_AbortSubTransaction_10), int32(_a_F_AbortSubTransaction_6))
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	goto L45
L54:
	;
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v23)+28))
	v275 = *(*int32)(unsafe.Add(mBase, _c_F_AbortSubTransaction[31]))
	if v273 <= v275 {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	v290 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_AbortSubTransaction[32])) = uint8(v290)
	*(*int32)(unsafe.Add(mBase, _c_F_AbortSubTransaction[33])) = v290
	goto L59
L56:
	;
	v278 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_AbortSubTransaction[34])) = v278
	*(*int32)(unsafe.Add(mBase, _c_F_AbortSubTransaction[35])) = v278
	*(*int32)(unsafe.Add(mBase, _c_F_AbortSubTransaction[36])) = v278
	*(*int32)(unsafe.Add(mBase, _c_F_AbortSubTransaction[31])) = v278
	goto L58
L57:
	;
	goto L58
L58:
	;
	goto L55
L59:
	;
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	F_AtEOSubXact_Parallel(m, int32(0), v296)
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+72)) = int32(0)
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v23)+40))
	if v301 != 0 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	F_AfterTriggerEndSubXact(m, int32(0))
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L1
	} else {
		goto L64
	}
L62:
	;
	goto L63
L63:
	;
	v572 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+68)))
	*(*uint8)(unsafe.Add(mBase, _c_F_AbortSubTransaction[37])) = uint8(v572)
	v574 = int32(_a_F_AbortSubTransaction_0)
	v576 = *(*int32)(unsafe.Add(mBase, _c_F_AbortSubTransaction[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_AbortSubTransaction[2])) = v576 - int32(1)
	m.G0 = v9 + int32(16)
	return
L64:
	;
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v23)+80))
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v305)+8))
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v23)+40))
	F_AtSubAbort_Portals(m, v308, v306, v309)
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v23)+80))
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v314)+8))
	F_AtEOSubXact_LargeObject(m, int32(0), v313, v315)
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	v319 = *(*int32)(unsafe.Add(mBase, _c_F_AbortSubTransaction[3]))
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v319)+28))
	goto L67
L67:
	;
	goto L68
L68:
	;
	v328 = *(*int32)(unsafe.Add(mBase, _c_F_AbortSubTransaction[38]))
	if v328 == int32(0) {
		goto L70
	} else {
		goto L71
	}
L69:
	;
	v339 = *(*int32)(unsafe.Add(mBase, _c_F_AbortSubTransaction[39]))
	if v339 == int32(0) {
		goto L74
	} else {
		goto L75
	}
L70:
	;
	goto L69
L71:
	;
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v328)))
	if v331 < v320 {
		goto L70
	} else {
		goto L72
	}
L72:
	;
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v328)+8))
	*(*int32)(unsafe.Add(mBase, _c_F_AbortSubTransaction[38])) = v334
	F_pfree(m, v328)
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L1
	} else {
		goto L73
	}
L73:
	;
	goto L68
L74:
	;
	v364 = F_RecordTransactionAbort(m, int32(1))
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L1
	} else {
		goto L81
	}
L75:
	;
	v343 = v339
	goto L76
L76:
	;
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v343)))
	if v348 < v320 {
		goto L74
	} else {
		goto L78
	}
L77:
	;
	goto L74
L78:
	;
	v351 = *(*int32)(unsafe.Add(mBase, uint32(v343)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_AbortSubTransaction[39])) = v351
	F_pfree(m, v343)
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L1
	} else {
		goto L79
	}
L79:
	;
	v356 = *(*int32)(unsafe.Add(mBase, _c_F_AbortSubTransaction[39]))
	if v356 != 0 {
		v343 = v356
		goto L76
	} else {
		goto L80
	}
L80:
	;
	goto L77
L81:
	;
	v366 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	if v366 != 0 {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v368 = *(*int32)(unsafe.Add(mBase, _c_F_AbortSubTransaction[3]))
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v368)+48))
	if v369 != 0 {
		goto L85
	} else {
		goto L86
	}
L83:
	;
	goto L84
L84:
	;
	v379 = *(*int32)(unsafe.Add(mBase, _c_F_AbortSubTransaction[40]))
	if v379 != 0 {
		goto L89
	} else {
		goto L90
	}
L85:
	;
	F_pfree(m, v369)
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
		goto L1
	} else {
		goto L88
	}
L86:
	;
	goto L87
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v368)+56)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v368)+48)) = int64(0)
	goto L84
L88:
	;
	goto L87
L89:
	;
	v380 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v23)+80))
	v382 = *(*int32)(unsafe.Add(mBase, uint32(v381)+8))
	v384 = v379
	goto L92
L90:
	;
	goto L91
L91:
	;
	v401 = *(*int32)(unsafe.Add(mBase, uint32(v23)+40))
	v403 = int32(0)
	F_ResourceOwnerReleaseInternal(m, v401, int32(1), v403, v403)
	mBase = m.M
	v406 = m.ExcPending
	if v406 != 0 {
		goto L1
	} else {
		goto L96
	}
L92:
	;
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v384)))
	v391 = *(*int32)(unsafe.Add(mBase, uint32(v384)+8))
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v384)+4))
	m.T0[v392].(func(*base.Module, int32, int32, int32, int32))(m, int32(2), v380, v382, v391)
	mBase = m.M
	v394 = m.ExcPending
	if v394 != 0 {
		goto L1
	} else {
		goto L94
	}
L93:
	;
	goto L91
L94:
	;
	if v389 != 0 {
		v384 = v389
		goto L92
	} else {
		goto L95
	}
L95:
	;
	goto L93
L96:
	;
	F_AtEOXact_Aio(m)
	mBase = m.M
	v408 = m.ExcPending
	if v408 != 0 {
		goto L1
	} else {
		goto L97
	}
L97:
	;
	v410 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	v411 = *(*int32)(unsafe.Add(mBase, uint32(v23)+80))
	v412 = *(*int32)(unsafe.Add(mBase, uint32(v411)+8))
	F_AtEOSubXact_RelationCache(m, int32(0), v410, v412)
	mBase = m.M
	v414 = m.ExcPending
	if v414 != 0 {
		goto L1
	} else {
		goto L98
	}
L98:
	;
	F_AtEOXact_TypeCache(m)
	mBase = m.M
	v416 = m.ExcPending
	if v416 != 0 {
		goto L1
	} else {
		goto L99
	}
L99:
	;
	F_AtEOSubXact_Inval(m, int32(0))
	mBase = m.M
	v419 = m.ExcPending
	if v419 != 0 {
		goto L1
	} else {
		goto L100
	}
L100:
	;
	v420 = *(*int32)(unsafe.Add(mBase, uint32(v23)+40))
	v422 = int32(0)
	F_ResourceOwnerReleaseInternal(m, v420, int32(2), v422, v422)
	mBase = m.M
	v425 = m.ExcPending
	if v425 != 0 {
		goto L1
	} else {
		goto L101
	}
L101:
	;
	v426 = *(*int32)(unsafe.Add(mBase, uint32(v23)+40))
	v428 = int32(0)
	F_ResourceOwnerReleaseInternal(m, v426, int32(3), v428, v428)
	mBase = m.M
	v431 = m.ExcPending
	if v431 != 0 {
		goto L1
	} else {
		goto L102
	}
L102:
	;
	F_smgrDoPendingDeletes(m, int32(0))
	mBase = m.M
	v434 = m.ExcPending
	if v434 != 0 {
		goto L1
	} else {
		goto L103
	}
L103:
	;
	v436 = *(*int32)(unsafe.Add(mBase, uint32(v23)+32))
	F_AtEOXact_GUC(m, int32(0), v436)
	mBase = m.M
	v438 = m.ExcPending
	if v438 != 0 {
		goto L1
	} else {
		goto L104
	}
L104:
	;
	v440 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	F_AtEOSubXact_SPI(m, int32(0), v440)
	mBase = m.M
	v442 = m.ExcPending
	if v442 != 0 {
		goto L1
	} else {
		goto L105
	}
L105:
	;
	v444 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	v445 = *(*int32)(unsafe.Add(mBase, uint32(v23)+80))
	v446 = *(*int32)(unsafe.Add(mBase, uint32(v445)+8))
	F_AtEOSubXact_on_commit_actions(m, int32(0), v444, v446)
	mBase = m.M
	v448 = m.ExcPending
	if v448 != 0 {
		goto L1
	} else {
		goto L106
	}
L106:
	;
	v450 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	v454 = *(*int32)(unsafe.Add(mBase, _c_F_AbortSubTransaction[41]))
	if v450 == v454 {
		goto L108
	} else {
		goto L109
	}
L107:
	;
	v478 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	v479 = *(*int32)(unsafe.Add(mBase, uint32(v23)+80))
	v480 = *(*int32)(unsafe.Add(mBase, uint32(v479)+8))
	F_AtEOSubXact_Files(m, int32(0), v478, v480)
	mBase = m.M
	v482 = m.ExcPending
	if v482 != 0 {
		goto L1
	} else {
		goto L114
	}
L108:
	;
	goto L112
L109:
	;
	goto L110
L110:
	;
	goto L107
L112:
	;
	goto L113
L113:
	;
	v459 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_AbortSubTransaction[42])) = uint8(v459)
	v462 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_AbortSubTransaction[43])) = v462
	*(*int32)(unsafe.Add(mBase, _c_F_AbortSubTransaction[41])) = v462
	*(*int32)(unsafe.Add(mBase, _c_F_AbortSubTransaction[44])) = v462
	*(*uint8)(unsafe.Add(mBase, _c_F_AbortSubTransaction[45])) = uint8(v462)
	v474 = *(*int32)(unsafe.Add(mBase, _c_F_AbortSubTransaction[46]))
	*(*int32)(unsafe.Add(mBase, uint32(v474)+28)) = v462
	goto L110
L114:
	;
	v484 = *(*int32)(unsafe.Add(mBase, uint32(v23)+28))
	F_AtEOSubXact_HashTables(m, int32(0), v484)
	mBase = m.M
	v486 = m.ExcPending
	if v486 != 0 {
		goto L1
	} else {
		goto L115
	}
L115:
	;
	v488 = *(*int32)(unsafe.Add(mBase, uint32(v23)+28))
	F_AtEOSubXact_PgStat(m, int32(0), v488)
	mBase = m.M
	v490 = m.ExcPending
	if v490 != 0 {
		goto L1
	} else {
		goto L116
	}
L116:
	;
	v491 = *(*int32)(unsafe.Add(mBase, uint32(v23)+28))
	v493 = *(*int32)(unsafe.Add(mBase, _c_F_AbortSubTransaction[47]))
	if v493 != 0 {
		goto L118
	} else {
		goto L119
	}
L117:
	;
	goto L63
L118:
	;
	v496 = v493
	goto L121
L119:
	;
	goto L120
L120:
	;
	v527 = *(*int32)(unsafe.Add(mBase, _c_F_AbortSubTransaction[48]))
	if v527 == int32(0) {
		goto L130
	} else {
		goto L131
	}
L121:
	;
	v500 = *(*int32)(unsafe.Add(mBase, uint32(v496)+4))
	if v500 < v491 {
		goto L117
	} else {
		goto L123
	}
L122:
	;
	goto L120
L123:
	;
	v502 = *(*int32)(unsafe.Add(mBase, uint32(v496)+8))
	v503 = *(*int32)(unsafe.Add(mBase, uint32(v496)))
	v504 = *(*int32)(unsafe.Add(mBase, uint32(v503)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v503)+44)) = v504 - int32(1)
	v508 = *(*int32)(unsafe.Add(mBase, uint32(v496)))
	v509 = *(*int32)(unsafe.Add(mBase, uint32(v508)+44))
	if v509 != 0 {
		v515 = v496
		goto L124
	} else {
		goto L125
	}
L124:
	;
	F_pfree(m, v515)
	mBase = m.M
	v517 = m.ExcPending
	if v517 != 0 {
		goto L1
	} else {
		goto L128
	}
L125:
	;
	v510 = *(*int32)(unsafe.Add(mBase, uint32(v508)+48))
	if v510 != 0 {
		v515 = v496
		goto L124
	} else {
		goto L126
	}
L126:
	;
	F_pfree(m, v508)
	mBase = m.M
	v512 = m.ExcPending
	if v512 != 0 {
		goto L1
	} else {
		goto L127
	}
L127:
	;
	v514 = *(*int32)(unsafe.Add(mBase, _c_F_AbortSubTransaction[47]))
	v515 = v514
	goto L124
L128:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_AbortSubTransaction[47])) = v502
	if v502 != 0 {
		v496 = v502
		goto L121
	} else {
		goto L129
	}
L129:
	;
	goto L122
L130:
	;
	v531 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_AbortSubTransaction[49])) = v531
	v534 = *(*int32)(unsafe.Add(mBase, _c_F_AbortSubTransaction[46]))
	*(*int32)(unsafe.Add(mBase, uint32(v534)+52)) = v531
	goto L117
L131:
	;
	goto L132
L132:
	;
	v538 = *(*int32)(unsafe.Add(mBase, _c_F_AbortSubTransaction[46]))
	v539 = *(*int32)(unsafe.Add(mBase, uint32(v538)+52))
	v540 = int32(3)
	v543 = *(*int32)(unsafe.Add(mBase, _c_F_AbortSubTransaction[48]))
	v546 = *(*int32)(unsafe.Add(mBase, uint32(v543-int32(48))))
	if base.B2i32(base.Ui32(v539) < base.Ui32(v540))|base.B2i32(base.Ui32(v546) < base.Ui32(v540)) == int32(0) {
		goto L134
	} else {
		goto L135
	}
L133:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_AbortSubTransaction[49])) = v546
	*(*int32)(unsafe.Add(mBase, uint32(v538)+52)) = v546
	goto L117
L134:
	;
	if v539-v546 < int32(0) {
		goto L133
	} else {
		goto L137
	}
L135:
	;
	goto L136
L136:
	;
	if base.Ui32(v546) <= base.Ui32(v539) {
		goto L117
	} else {
		goto L138
	}
L137:
	;
	goto L117
L138:
	;
	goto L133
}
func F_CheckSubDeadTupleRetention(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	if l3 != 0 {
		if l0 != 0 {
			v13 = *(*int32)(unsafe.Add(mBase, _c_F_CheckSubDeadTupleRetention[0]))
			if v13 <= int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v100 = m.ExcPending
				if v100 != 0 {
					return
				} else {
					F_errcode(m, int32(325))
					mBase = m.M
					v103 = m.ExcPending
					if v103 != 0 {
						return
					} else {
						F_errmsg(m, int32(_a_F_CheckSubDeadTupleRetention_0), int32(0))
						mBase = m.M
						v107 = m.ExcPending
						if v107 != 0 {
							return
						} else {
							F_errhint(m, int32(_a_F_CheckSubDeadTupleRetention_1), int32(0))
							mBase = m.M
							v111 = m.ExcPending
							if v111 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_CheckSubDeadTupleRetention_2), int32(3293), int32(_a_F_CheckSubDeadTupleRetention_3))
								mBase = m.M
								v116 = m.ExcPending
								if v116 != 0 {
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
			} else {
				if l0 == int32(0) {
					v47 = int32(0)
					if base.B2i32(l1 == v47)|base.B2i32(l4 == v47) != 0 {
						m.G0 = v10 + int32(32)
						return
					} else {
						v53 = F_errstart(m, l2, int32(0))
						mBase = m.M
						v54 = m.ExcPending
						if v54 != 0 {
							return
						} else {
							if v53 == int32(0) {
								m.G0 = v10 + int32(32)
								return
							} else {
								F_errcode(m, int32(50856066))
								mBase = m.M
								v59 = m.ExcPending
								if v59 != 0 {
									return
								} else {
									F_errmsg(m, int32(_a_F_CheckSubDeadTupleRetention_4), int32(0))
									mBase = m.M
									v63 = m.ExcPending
									if v63 != 0 {
										return
									} else {
										v64 = int32(3308)
										if l2 < int32(19) {
											v88 = v64
											F_errfinish(m, int32(_a_F_CheckSubDeadTupleRetention_2), v88, int32(_a_F_CheckSubDeadTupleRetention_3))
											mBase = m.M
											v92 = m.ExcPending
											if v92 != 0 {
												return
											} else {
												m.G0 = v10 + int32(32)
												return
											}
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v10))) = int32(_a_F_CheckSubDeadTupleRetention_5)
											F_errhint(m, int32(_a_F_CheckSubDeadTupleRetention_6), v10)
											mBase = m.M
											v71 = m.ExcPending
											if v71 != 0 {
												return
											} else {
												v88 = v64
												F_errfinish(m, int32(_a_F_CheckSubDeadTupleRetention_2), v88, int32(_a_F_CheckSubDeadTupleRetention_3))
												mBase = m.M
												v92 = m.ExcPending
												if v92 != 0 {
													return
												} else {
													m.G0 = v10 + int32(32)
													return
												}
											}
										}
									}
								}
							}
						}
					}
				} else {
					v19 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CheckSubDeadTupleRetention[1])))
					if v19&int32(1) != 0 {
						v47 = int32(0)
						if base.B2i32(l1 == v47)|base.B2i32(l4 == v47) != 0 {
							m.G0 = v10 + int32(32)
							return
						} else {
							v53 = F_errstart(m, l2, int32(0))
							mBase = m.M
							v54 = m.ExcPending
							if v54 != 0 {
								return
							} else {
								if v53 == int32(0) {
									m.G0 = v10 + int32(32)
									return
								} else {
									F_errcode(m, int32(50856066))
									mBase = m.M
									v59 = m.ExcPending
									if v59 != 0 {
										return
									} else {
										F_errmsg(m, int32(_a_F_CheckSubDeadTupleRetention_4), int32(0))
										mBase = m.M
										v63 = m.ExcPending
										if v63 != 0 {
											return
										} else {
											v64 = int32(3308)
											if l2 < int32(19) {
												v88 = v64
												F_errfinish(m, int32(_a_F_CheckSubDeadTupleRetention_2), v88, int32(_a_F_CheckSubDeadTupleRetention_3))
												mBase = m.M
												v92 = m.ExcPending
												if v92 != 0 {
													return
												} else {
													m.G0 = v10 + int32(32)
													return
												}
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v10))) = int32(_a_F_CheckSubDeadTupleRetention_5)
												F_errhint(m, int32(_a_F_CheckSubDeadTupleRetention_6), v10)
												mBase = m.M
												v71 = m.ExcPending
												if v71 != 0 {
													return
												} else {
													v88 = v64
													F_errfinish(m, int32(_a_F_CheckSubDeadTupleRetention_2), v88, int32(_a_F_CheckSubDeadTupleRetention_3))
													mBase = m.M
													v92 = m.ExcPending
													if v92 != 0 {
														return
													} else {
														m.G0 = v10 + int32(32)
														return
													}
												}
											}
										}
									}
								}
							}
						}
					} else {
						v24 = F_errstart(m, int32(19), int32(0))
						mBase = m.M
						v25 = m.ExcPending
						if v25 != 0 {
							return
						} else {
							if v24 == int32(0) {
								v47 = int32(0)
								if base.B2i32(l1 == v47)|base.B2i32(l4 == v47) != 0 {
									m.G0 = v10 + int32(32)
									return
								} else {
									v53 = F_errstart(m, l2, int32(0))
									mBase = m.M
									v54 = m.ExcPending
									if v54 != 0 {
										return
									} else {
										if v53 == int32(0) {
											m.G0 = v10 + int32(32)
											return
										} else {
											F_errcode(m, int32(50856066))
											mBase = m.M
											v59 = m.ExcPending
											if v59 != 0 {
												return
											} else {
												F_errmsg(m, int32(_a_F_CheckSubDeadTupleRetention_4), int32(0))
												mBase = m.M
												v63 = m.ExcPending
												if v63 != 0 {
													return
												} else {
													v64 = int32(3308)
													if l2 < int32(19) {
														v88 = v64
														F_errfinish(m, int32(_a_F_CheckSubDeadTupleRetention_2), v88, int32(_a_F_CheckSubDeadTupleRetention_3))
														mBase = m.M
														v92 = m.ExcPending
														if v92 != 0 {
															return
														} else {
															m.G0 = v10 + int32(32)
															return
														}
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v10))) = int32(_a_F_CheckSubDeadTupleRetention_5)
														F_errhint(m, int32(_a_F_CheckSubDeadTupleRetention_6), v10)
														mBase = m.M
														v71 = m.ExcPending
														if v71 != 0 {
															return
														} else {
															v88 = v64
															F_errfinish(m, int32(_a_F_CheckSubDeadTupleRetention_2), v88, int32(_a_F_CheckSubDeadTupleRetention_3))
															mBase = m.M
															v92 = m.ExcPending
															if v92 != 0 {
																return
															} else {
																m.G0 = v10 + int32(32)
																return
															}
														}
													}
												}
											}
										}
									}
								}
							} else {
								F_errcode(m, int32(50856066))
								mBase = m.M
								v30 = m.ExcPending
								if v30 != 0 {
									return
								} else {
									F_errmsg(m, int32(_a_F_CheckSubDeadTupleRetention_7), int32(0))
									mBase = m.M
									v34 = m.ExcPending
									if v34 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = int32(_a_F_CheckSubDeadTupleRetention_8)
										F_errhint(m, int32(_a_F_CheckSubDeadTupleRetention_9), v10+int32(16))
										mBase = m.M
										v41 = m.ExcPending
										if v41 != 0 {
											return
										} else {
											F_errfinish(m, int32(_a_F_CheckSubDeadTupleRetention_2), int32(3300), int32(_a_F_CheckSubDeadTupleRetention_3))
											mBase = m.M
											v46 = m.ExcPending
											if v46 != 0 {
												return
											} else {
												v47 = int32(0)
												if base.B2i32(l1 == v47)|base.B2i32(l4 == v47) != 0 {
													m.G0 = v10 + int32(32)
													return
												} else {
													v53 = F_errstart(m, l2, int32(0))
													mBase = m.M
													v54 = m.ExcPending
													if v54 != 0 {
														return
													} else {
														if v53 == int32(0) {
															m.G0 = v10 + int32(32)
															return
														} else {
															F_errcode(m, int32(50856066))
															mBase = m.M
															v59 = m.ExcPending
															if v59 != 0 {
																return
															} else {
																F_errmsg(m, int32(_a_F_CheckSubDeadTupleRetention_4), int32(0))
																mBase = m.M
																v63 = m.ExcPending
																if v63 != 0 {
																	return
																} else {
																	v64 = int32(3308)
																	if l2 < int32(19) {
																		v88 = v64
																		F_errfinish(m, int32(_a_F_CheckSubDeadTupleRetention_2), v88, int32(_a_F_CheckSubDeadTupleRetention_3))
																		mBase = m.M
																		v92 = m.ExcPending
																		if v92 != 0 {
																			return
																		} else {
																			m.G0 = v10 + int32(32)
																			return
																		}
																	} else {
																		*(*int32)(unsafe.Add(mBase, uint32(v10))) = int32(_a_F_CheckSubDeadTupleRetention_5)
																		F_errhint(m, int32(_a_F_CheckSubDeadTupleRetention_6), v10)
																		mBase = m.M
																		v71 = m.ExcPending
																		if v71 != 0 {
																			return
																		} else {
																			v88 = v64
																			F_errfinish(m, int32(_a_F_CheckSubDeadTupleRetention_2), v88, int32(_a_F_CheckSubDeadTupleRetention_3))
																			mBase = m.M
																			v92 = m.ExcPending
																			if v92 != 0 {
																				return
																			} else {
																				m.G0 = v10 + int32(32)
																				return
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
									}
								}
							}
						}
					}
				}
			}
		} else {
			if l0 == int32(0) {
				v47 = int32(0)
				if base.B2i32(l1 == v47)|base.B2i32(l4 == v47) != 0 {
					m.G0 = v10 + int32(32)
					return
				} else {
					v53 = F_errstart(m, l2, int32(0))
					mBase = m.M
					v54 = m.ExcPending
					if v54 != 0 {
						return
					} else {
						if v53 == int32(0) {
							m.G0 = v10 + int32(32)
							return
						} else {
							F_errcode(m, int32(50856066))
							mBase = m.M
							v59 = m.ExcPending
							if v59 != 0 {
								return
							} else {
								F_errmsg(m, int32(_a_F_CheckSubDeadTupleRetention_4), int32(0))
								mBase = m.M
								v63 = m.ExcPending
								if v63 != 0 {
									return
								} else {
									v64 = int32(3308)
									if l2 < int32(19) {
										v88 = v64
										F_errfinish(m, int32(_a_F_CheckSubDeadTupleRetention_2), v88, int32(_a_F_CheckSubDeadTupleRetention_3))
										mBase = m.M
										v92 = m.ExcPending
										if v92 != 0 {
											return
										} else {
											m.G0 = v10 + int32(32)
											return
										}
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v10))) = int32(_a_F_CheckSubDeadTupleRetention_5)
										F_errhint(m, int32(_a_F_CheckSubDeadTupleRetention_6), v10)
										mBase = m.M
										v71 = m.ExcPending
										if v71 != 0 {
											return
										} else {
											v88 = v64
											F_errfinish(m, int32(_a_F_CheckSubDeadTupleRetention_2), v88, int32(_a_F_CheckSubDeadTupleRetention_3))
											mBase = m.M
											v92 = m.ExcPending
											if v92 != 0 {
												return
											} else {
												m.G0 = v10 + int32(32)
												return
											}
										}
									}
								}
							}
						}
					}
				}
			} else {
				v19 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CheckSubDeadTupleRetention[1])))
				if v19&int32(1) != 0 {
					v47 = int32(0)
					if base.B2i32(l1 == v47)|base.B2i32(l4 == v47) != 0 {
						m.G0 = v10 + int32(32)
						return
					} else {
						v53 = F_errstart(m, l2, int32(0))
						mBase = m.M
						v54 = m.ExcPending
						if v54 != 0 {
							return
						} else {
							if v53 == int32(0) {
								m.G0 = v10 + int32(32)
								return
							} else {
								F_errcode(m, int32(50856066))
								mBase = m.M
								v59 = m.ExcPending
								if v59 != 0 {
									return
								} else {
									F_errmsg(m, int32(_a_F_CheckSubDeadTupleRetention_4), int32(0))
									mBase = m.M
									v63 = m.ExcPending
									if v63 != 0 {
										return
									} else {
										v64 = int32(3308)
										if l2 < int32(19) {
											v88 = v64
											F_errfinish(m, int32(_a_F_CheckSubDeadTupleRetention_2), v88, int32(_a_F_CheckSubDeadTupleRetention_3))
											mBase = m.M
											v92 = m.ExcPending
											if v92 != 0 {
												return
											} else {
												m.G0 = v10 + int32(32)
												return
											}
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v10))) = int32(_a_F_CheckSubDeadTupleRetention_5)
											F_errhint(m, int32(_a_F_CheckSubDeadTupleRetention_6), v10)
											mBase = m.M
											v71 = m.ExcPending
											if v71 != 0 {
												return
											} else {
												v88 = v64
												F_errfinish(m, int32(_a_F_CheckSubDeadTupleRetention_2), v88, int32(_a_F_CheckSubDeadTupleRetention_3))
												mBase = m.M
												v92 = m.ExcPending
												if v92 != 0 {
													return
												} else {
													m.G0 = v10 + int32(32)
													return
												}
											}
										}
									}
								}
							}
						}
					}
				} else {
					v24 = F_errstart(m, int32(19), int32(0))
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return
					} else {
						if v24 == int32(0) {
							v47 = int32(0)
							if base.B2i32(l1 == v47)|base.B2i32(l4 == v47) != 0 {
								m.G0 = v10 + int32(32)
								return
							} else {
								v53 = F_errstart(m, l2, int32(0))
								mBase = m.M
								v54 = m.ExcPending
								if v54 != 0 {
									return
								} else {
									if v53 == int32(0) {
										m.G0 = v10 + int32(32)
										return
									} else {
										F_errcode(m, int32(50856066))
										mBase = m.M
										v59 = m.ExcPending
										if v59 != 0 {
											return
										} else {
											F_errmsg(m, int32(_a_F_CheckSubDeadTupleRetention_4), int32(0))
											mBase = m.M
											v63 = m.ExcPending
											if v63 != 0 {
												return
											} else {
												v64 = int32(3308)
												if l2 < int32(19) {
													v88 = v64
													F_errfinish(m, int32(_a_F_CheckSubDeadTupleRetention_2), v88, int32(_a_F_CheckSubDeadTupleRetention_3))
													mBase = m.M
													v92 = m.ExcPending
													if v92 != 0 {
														return
													} else {
														m.G0 = v10 + int32(32)
														return
													}
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v10))) = int32(_a_F_CheckSubDeadTupleRetention_5)
													F_errhint(m, int32(_a_F_CheckSubDeadTupleRetention_6), v10)
													mBase = m.M
													v71 = m.ExcPending
													if v71 != 0 {
														return
													} else {
														v88 = v64
														F_errfinish(m, int32(_a_F_CheckSubDeadTupleRetention_2), v88, int32(_a_F_CheckSubDeadTupleRetention_3))
														mBase = m.M
														v92 = m.ExcPending
														if v92 != 0 {
															return
														} else {
															m.G0 = v10 + int32(32)
															return
														}
													}
												}
											}
										}
									}
								}
							}
						} else {
							F_errcode(m, int32(50856066))
							mBase = m.M
							v30 = m.ExcPending
							if v30 != 0 {
								return
							} else {
								F_errmsg(m, int32(_a_F_CheckSubDeadTupleRetention_7), int32(0))
								mBase = m.M
								v34 = m.ExcPending
								if v34 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = int32(_a_F_CheckSubDeadTupleRetention_8)
									F_errhint(m, int32(_a_F_CheckSubDeadTupleRetention_9), v10+int32(16))
									mBase = m.M
									v41 = m.ExcPending
									if v41 != 0 {
										return
									} else {
										F_errfinish(m, int32(_a_F_CheckSubDeadTupleRetention_2), int32(3300), int32(_a_F_CheckSubDeadTupleRetention_3))
										mBase = m.M
										v46 = m.ExcPending
										if v46 != 0 {
											return
										} else {
											v47 = int32(0)
											if base.B2i32(l1 == v47)|base.B2i32(l4 == v47) != 0 {
												m.G0 = v10 + int32(32)
												return
											} else {
												v53 = F_errstart(m, l2, int32(0))
												mBase = m.M
												v54 = m.ExcPending
												if v54 != 0 {
													return
												} else {
													if v53 == int32(0) {
														m.G0 = v10 + int32(32)
														return
													} else {
														F_errcode(m, int32(50856066))
														mBase = m.M
														v59 = m.ExcPending
														if v59 != 0 {
															return
														} else {
															F_errmsg(m, int32(_a_F_CheckSubDeadTupleRetention_4), int32(0))
															mBase = m.M
															v63 = m.ExcPending
															if v63 != 0 {
																return
															} else {
																v64 = int32(3308)
																if l2 < int32(19) {
																	v88 = v64
																	F_errfinish(m, int32(_a_F_CheckSubDeadTupleRetention_2), v88, int32(_a_F_CheckSubDeadTupleRetention_3))
																	mBase = m.M
																	v92 = m.ExcPending
																	if v92 != 0 {
																		return
																	} else {
																		m.G0 = v10 + int32(32)
																		return
																	}
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(v10))) = int32(_a_F_CheckSubDeadTupleRetention_5)
																	F_errhint(m, int32(_a_F_CheckSubDeadTupleRetention_6), v10)
																	mBase = m.M
																	v71 = m.ExcPending
																	if v71 != 0 {
																		return
																	} else {
																		v88 = v64
																		F_errfinish(m, int32(_a_F_CheckSubDeadTupleRetention_2), v88, int32(_a_F_CheckSubDeadTupleRetention_3))
																		mBase = m.M
																		v92 = m.ExcPending
																		if v92 != 0 {
																			return
																		} else {
																			m.G0 = v10 + int32(32)
																			return
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
								}
							}
						}
					}
				}
			}
		}
	} else {
		if l5 == int32(0) {
			m.G0 = v10 + int32(32)
			return
		} else {
			v76 = F_errstart(m, int32(18), int32(0))
			mBase = m.M
			v77 = m.ExcPending
			if v77 != 0 {
				return
			} else {
				if v76 == int32(0) {
					m.G0 = v10 + int32(32)
					return
				} else {
					F_errcode(m, int32(50856066))
					mBase = m.M
					v82 = m.ExcPending
					if v82 != 0 {
						return
					} else {
						F_errmsg(m, int32(_a_F_CheckSubDeadTupleRetention_10), int32(0))
						mBase = m.M
						v86 = m.ExcPending
						if v86 != 0 {
							return
						} else {
							v88 = int32(3314)
							F_errfinish(m, int32(_a_F_CheckSubDeadTupleRetention_2), v88, int32(_a_F_CheckSubDeadTupleRetention_3))
							mBase = m.M
							v92 = m.ExcPending
							if v92 != 0 {
								return
							} else {
								m.G0 = v10 + int32(32)
								return
							}
						}
					}
				}
			}
		}
	}
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
	var v22 int32
	_ = v22
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_CleanupSubTransaction[0]))
	v12 = int32(10)
	goto L3
L1:
	;
	if v52 != 0 {
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
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v19<<(uint(int32(2))%32))+uint32(_c_F_CleanupSubTransaction[2])))
	goto L6
L4:
	;
	v35 = int32(0)
	goto L11
L6:
	;
	goto L7
L7:
	;
	if int32(0)|base.B2i32(v22 == int32(15)) != 0 {
		goto L4
	} else {
		goto L9
	}
L9:
	;
	if v22 <= v12 {
		v52 = int32(1)
		goto L2
	} else {
		goto L10
	}
L10:
	;
	goto L4
L11:
	;
	v39 = *(*int32)(unsafe.Add(mBase, _c_F_CleanupSubTransaction[3]))
	if v39 != int32(2) {
		v52 = v35
		goto L2
	} else {
		goto L12
	}
L12:
	;
	v43 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CleanupSubTransaction[4])))
	if v43&int32(1) != 0 {
		v52 = v35
		goto L2
	} else {
		goto L13
	}
L13:
	;
	v49 = *(*int32)(unsafe.Add(mBase, _c_F_CleanupSubTransaction[5]))
	v52 = int32(0) | base.B2i32(v49 <= v12)
	goto L2
L14:
	;
	v56 = *(*int32)(unsafe.Add(mBase, _c_F_CleanupSubTransaction[0]))
	F_ShowTransactionStateRec(m, int32(_a_F_CleanupSubTransaction_0), v56)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	goto L16
L16:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v11)+20))
	if v59 == int32(4) {
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
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	v87 = m.G0
	v89 = v87 - int32(32)
	m.G0 = v89
	v92 = v89 + int32(12)
	v94 = *(*int32)(unsafe.Add(mBase, _c_F_CleanupSubTransaction[6]))
	F_hash_seq_init(m, v92, v94)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L17
	} else {
		goto L28
	}
L20:
	;
	v64 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L17
	} else {
		goto L21
	}
L21:
	;
	if v64 == int32(0) {
		goto L19
	} else {
		goto L22
	}
L22:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v11)+20))
	if base.Ui32(v68) <= base.Ui32(int32(5)) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v68<<(uint(int32(2))%32))+uint32(_c_F_CleanupSubTransaction[7])))
	v75 = v73
	goto L25
L24:
	;
	v75 = int32(_a_F_CleanupSubTransaction_1)
	goto L25
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v75
	F_errmsg_internal(m, int32(_a_F_CleanupSubTransaction_2), v8)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L17
	} else {
		goto L26
	}
L26:
	;
	F_errfinish(m, int32(_a_F_CleanupSubTransaction_3), int32(_a_F_CleanupSubTransaction_4), int32(_a_F_CleanupSubTransaction_0))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L17
	} else {
		goto L27
	}
L27:
	;
	goto L19
L28:
	;
	v97 = F_hash_seq_search(m, v92)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L17
	} else {
		goto L29
	}
L29:
	;
	if v97 != 0 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v100 = v97
	goto L33
L31:
	;
	goto L32
L32:
	;
	m.G0 = v89 + int32(32)
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v11)+80))
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v145)+40))
	*(*int32)(unsafe.Add(mBase, _c_F_CleanupSubTransaction[8])) = v146
	*(*int32)(unsafe.Add(mBase, _c_F_CleanupSubTransaction[9])) = v146
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v11)+40))
	if v150 != 0 {
		goto L53
	} else {
		goto L54
	}
L33:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v100)+64))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v104)+20))
	if v86 == v105 {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	goto L32
L35:
	;
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104)+84)))
	if v107 == int32(1) {
		goto L38
	} else {
		goto L39
	}
L36:
	;
	goto L37
L37:
	;
	v134 = F_hash_seq_search(m, v89+int32(12))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L17
	} else {
		goto L51
	}
L38:
	;
	v110 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v104)+84)) = uint8(v110)
	goto L40
L39:
	;
	goto L40
L40:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v104)+16))
	if v112 != 0 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v115 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L17
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	F_PortalDrop(m, v104, int32(0))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L17
	} else {
		goto L50
	}
L44:
	;
	if v115 != 0 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v104)))
	*(*int32)(unsafe.Add(mBase, uint32(v89))) = v117
	F_errmsg_internal(m, int32(_a_F_CleanupSubTransaction_5), v89)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L17
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v104)+16)) = int32(0)
	goto L43
L48:
	;
	F_errfinish(m, int32(_a_F_CleanupSubTransaction_6), int32(1122), int32(_a_F_CleanupSubTransaction_7))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
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
	if v134 != 0 {
		v100 = v134
		goto L33
	} else {
		goto L52
	}
L52:
	;
	goto L34
L53:
	;
	F_ResourceOwnerDelete(m, v150)
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
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
	v157 = *(*int32)(unsafe.Add(mBase, _c_F_CleanupSubTransaction[0]))
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v157)+44))
	*(*int32)(unsafe.Add(mBase, _c_F_CleanupSubTransaction[10])) = v158
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v157)+80))
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v161)+36))
	*(*int32)(unsafe.Add(mBase, _c_F_CleanupSubTransaction[11])) = v162
	v165 = *(*int32)(unsafe.Add(mBase, _c_F_CleanupSubTransaction[12]))
	if v165 != 0 {
		goto L57
	} else {
		goto L58
	}
L56:
	;
	goto L55
L57:
	;
	F_MemoryContextReset(m, v165)
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L17
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v157)+36))
	if v168 != 0 {
		goto L61
	} else {
		goto L62
	}
L60:
	;
	goto L59
L61:
	;
	F_MemoryContextDelete(m, v168)
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L17
	} else {
		goto L64
	}
L62:
	;
	goto L63
L63:
	;
	v171 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v157)+36)) = v171
	*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = v171
	F_PopTransaction(m)
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
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
	var v176 int32
	_ = v176
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
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
	if v206 != 0 {
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
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v173<<(uint(int32(2))%32))+uint32(_c_F_StartSubTransaction[15])))
	goto L39
L37:
	;
	v189 = int32(0)
	goto L44
L39:
	;
	goto L40
L40:
	;
	if int32(0)|base.B2i32(v176 == int32(15)) != 0 {
		goto L37
	} else {
		goto L42
	}
L42:
	;
	if v176 <= v166 {
		v206 = int32(1)
		goto L35
	} else {
		goto L43
	}
L43:
	;
	goto L37
L44:
	;
	v193 = *(*int32)(unsafe.Add(mBase, _c_F_StartSubTransaction[16]))
	if v193 != int32(2) {
		v206 = v189
		goto L35
	} else {
		goto L45
	}
L45:
	;
	v197 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartSubTransaction[17])))
	if v197&int32(1) != 0 {
		v206 = v189
		goto L35
	} else {
		goto L46
	}
L46:
	;
	v203 = *(*int32)(unsafe.Add(mBase, _c_F_StartSubTransaction[18]))
	v206 = int32(0) | base.B2i32(v203 <= v166)
	goto L35
L47:
	;
	v210 = *(*int32)(unsafe.Add(mBase, _c_F_StartSubTransaction[0]))
	F_ShowTransactionStateRec(m, int32(_a_F_StartSubTransaction_8), v210)
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
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
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v94 int32
	_ = v94
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	v2 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	if l0 == v2 {
		v85 = v2
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L15
	} else {
		goto L22
	}
L2:
	;
	m.G0 = v9 + int32(16)
	return v85
L3:
	;
	v14 = *(*int32)(unsafe.Add(mBase, _c_F_SubTransGetTopmostTransaction[0]))
	v15 = l0
	v18 = v14
	goto L4
L4:
	;
	v21 = int32(3)
	v22 = base.B2i32(base.Ui32(v15) < base.Ui32(v21))
	if v22|base.B2i32(base.Ui32(v18) < base.Ui32(v21)) == int32(0) {
		goto L9
	} else {
		goto L10
	}
L5:
	;
	v85 = v15
	goto L2
L6:
	;
	if v81 != 0 {
		v15 = v81
		v18 = v82
		goto L4
	} else {
		goto L21
	}
L7:
	;
	if base.Ui32(v15) <= base.Ui32(v77) {
		v94 = v77
		goto L1
	} else {
		goto L20
	}
L8:
	;
	v39 = int32(base.Ui32(v15) >> (uint(int32(11)) % 32))
	v43 = F_SimpleLruReadPage_ReadOnly(m, int32(_a_F_SubTransGetTopmostTransaction_0), base.I64_extend_i32_u(v39), v9+int32(12))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L15
	} else {
		goto L16
	}
L9:
	;
	if v15-v18 < int32(0) {
		v85 = v15
		goto L2
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	if base.Ui32(v15) < base.Ui32(v18) {
		v85 = v15
		goto L2
	} else {
		goto L13
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v15
	goto L8
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v15
	if base.Ui32(int32(2)) < base.Ui32(v15) {
		goto L8
	} else {
		goto L14
	}
L14:
	;
	v77 = int32(0)
	v78 = v18
	goto L7
L15:
	;
	return int32(0)
L16:
	;
	v48 = *(*int32)(unsafe.Add(mBase, _c_F_SubTransGetTopmostTransaction[1]))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
	v50 = int32(2)
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v49+v43<<(uint(v50)%32))))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v53+v15&int32(2047)<<(uint(v50)%32))))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v48)+28))
	v62 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_SubTransGetTopmostTransaction[2])))
	v63 = base.I32_rem_u_s(v39, v62)
	F_LWLockRelease(m, v60+v63<<(uint(int32(7))%32))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v70 = *(*int32)(unsafe.Add(mBase, _c_F_SubTransGetTopmostTransaction[0]))
	if v22|base.B2i32(base.Ui32(v59) < base.Ui32(int32(3))) != 0 {
		v77 = v59
		v78 = v70
		goto L7
	} else {
		goto L18
	}
L18:
	;
	if v59-v15 < int32(0) {
		v81 = v59
		v82 = v70
		goto L6
	} else {
		goto L19
	}
L19:
	;
	v94 = v59
	goto L1
L20:
	;
	v81 = v77
	v82 = v78
	goto L6
L21:
	;
	goto L5
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v94
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v15
	F_errmsg_internal(m, int32(_a_F_SubTransGetTopmostTransaction_1), v9)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L15
	} else {
		goto L23
	}
L23:
	;
	F_errfinish(m, int32(_a_F_SubTransGetTopmostTransaction_2), int32(192), int32(_a_F_SubTransGetTopmostTransaction_3))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L15
	} else {
		goto L24
	}
L24:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_SubTransPagePrecedes(m *base.Module, l0 int64, l1 int64) int32 {
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	v5 = int32(11)
	v10 = base.I32_wrap_i64(l0)<<(uint(v5)%32) - base.I32_wrap_i64(l1)<<(uint(v5)%32)
	return int32(base.Ui32(v10&(v10-int32(2047))) >> (uint(int32(31)) % 32))
}
func F_SubTransSetParent(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = l0
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_SubTransSetParent[0]))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+28))
	v15 = int32(base.Ui32(l0) >> (uint(int32(11)) % 32))
	v17 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_SubTransSetParent[1])))
	v18 = base.I32_rem_u_s(v15, v17)
	v21 = v13 + v18<<(uint(int32(7))%32)
	v23 = F_LWLockAcquire(m, v21, int32(0))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		return
	} else {
		v30 = F_SimpleLruReadPage(m, int32(_a_F_SubTransSetParent_0), base.I64_extend_i32_u(v15), int32(1), v8+int32(12))
		mBase = m.M
		v31 = m.ExcPending
		if v31 != 0 {
			return
		} else {
			v33 = *(*int32)(unsafe.Add(mBase, _c_F_SubTransSetParent[0]))
			v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
			v35 = int32(2)
			v38 = *(*int32)(unsafe.Add(mBase, uint32(v34+v30<<(uint(v35)%32))))
			v43 = v38 + l0&int32(2047)<<(uint(v35)%32)
			v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
			if l1 != v44 {
				*(*int32)(unsafe.Add(mBase, uint32(v43))) = l1
				v48 = *(*int32)(unsafe.Add(mBase, _c_F_SubTransSetParent[0]))
				v49 = *(*int32)(unsafe.Add(mBase, uint32(v48)+12))
				v51 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(v49+v30))) = uint8(v51)
			} else {
			}
			F_LWLockRelease(m, v21)
			mBase = m.M
			v54 = m.ExcPending
			if v54 != 0 {
				return
			} else {
				m.G0 = v8 + int32(16)
				return
			}
		}
	}
}
