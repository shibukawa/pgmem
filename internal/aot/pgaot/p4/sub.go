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
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v161 int32
	_ = v161
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v189 int32
	_ = v189
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v246 int32
	_ = v246
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v278 int32
	_ = v278
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v289 int32
	_ = v289
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v317 int32
	_ = v317
	var v320 int32
	_ = v320
	var v323 int32
	_ = v323
	var v326 int32
	_ = v326
	var v328 int32
	_ = v328
	var v332 int32
	_ = v332
	var v338 int32
	_ = v338
	var v341 int32
	_ = v341
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v362 int32
	_ = v362
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v375 int32
	_ = v375
	var v381 int32
	_ = v381
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v386 int32
	_ = v386
	var v394 int32
	_ = v394
	var v396 int32
	_ = v396
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v407 int32
	_ = v407
	var v409 int32
	_ = v409
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v415 int32
	_ = v415
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v421 int32
	_ = v421
	var v424 int32
	_ = v424
	var v427 int32
	_ = v427
	var v429 int32
	_ = v429
	var v431 int32
	_ = v431
	var v433 int32
	_ = v433
	var v435 int32
	_ = v435
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v441 int32
	_ = v441
	var v443 int32
	_ = v443
	var v447 int32
	_ = v447
	var v452 int32
	_ = v452
	var v455 int32
	_ = v455
	var v467 int32
	_ = v467
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v475 int32
	_ = v475
	var v477 int32
	_ = v477
	var v479 int32
	_ = v479
	var v481 int32
	_ = v481
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v487 int32
	_ = v487
	var v490 int32
	_ = v490
	var v495 int32
	_ = v495
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v507 int32
	_ = v507
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v512 int32
	_ = v512
	var v523 int32
	_ = v523
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v528 int32
	_ = v528
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v543 int32
	_ = v543
	var v546 int32
	_ = v546
	var v548 int32
	_ = v548
	var v552 int32
	_ = v552
	var v569 int32
	_ = v569
	var v571 int32
	_ = v571
	var v573 int32
	_ = v573
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
	v72 = m.ExcPending
	if v72 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v35)+220)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v35)+224)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v35))) = v53 + int32(2)
	v67 = *(*int32)(unsafe.Add(mBase, _c_F_AbortSubTransaction[8]))
	*(*int32)(unsafe.Add(mBase, _c_F_AbortSubTransaction[8])) = v67 - v50
	goto L4
L8:
	;
	F_UnlockBuffers(m)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	v75 = int32(0)
	v82 = *(*int32)(unsafe.Add(mBase, _c_F_AbortSubTransaction[9]))
	if v82 <= v75 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	F_ConditionVariableCancelSleep(m)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L1
	} else {
		goto L23
	}
L11:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_AbortSubTransaction[10])) = int64(0)
	*(*int32)(unsafe.Add(mBase, _c_F_AbortSubTransaction[11])) = int32(_a_F_AbortSubTransaction_2)
	v161 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_AbortSubTransaction[9])) = v161
	*(*int32)(unsafe.Add(mBase, _c_F_AbortSubTransaction[12])) = v161
	*(*uint8)(unsafe.Add(mBase, _c_F_AbortSubTransaction[13])) = uint8(v161)
	*(*uint8)(unsafe.Add(mBase, _c_F_AbortSubTransaction[14])) = uint8(v161)
	goto L10
L12:
	;
	v86 = v82 & int32(7)
	v88 = *(*int32)(unsafe.Add(mBase, _c_F_AbortSubTransaction[15]))
	if base.Ui32(int32(8)) <= base.Ui32(v82) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v94 = v75
	v98 = v75
	goto L16
L14:
	;
	v126 = v75
	goto L15
L15:
	;
	v132 = int32(0)
	v133 = v126
	goto L20
L16:
	;
	v101 = v88 + v94*int32(_a_F_AbortSubTransaction_3)
	v102 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v101)+uint32(_c_F_AbortSubTransaction[16]))) = uint8(v102)
	*(*uint8)(unsafe.Add(mBase, uint32(v101)+uint32(_c_F_AbortSubTransaction[17]))) = uint8(v102)
	*(*uint8)(unsafe.Add(mBase, uint32(v101)+uint32(_c_F_AbortSubTransaction[18]))) = uint8(v102)
	*(*uint8)(unsafe.Add(mBase, uint32(v101)+uint32(_c_F_AbortSubTransaction[19]))) = uint8(v102)
	*(*uint8)(unsafe.Add(mBase, uint32(v101)+uint32(_c_F_AbortSubTransaction[20]))) = uint8(v102)
	*(*uint8)(unsafe.Add(mBase, uint32(v101)+uint32(_c_F_AbortSubTransaction[21]))) = uint8(v102)
	*(*uint8)(unsafe.Add(mBase, uint32(v101)+uint32(_c_F_AbortSubTransaction[22]))) = uint8(v102)
	*(*uint8)(unsafe.Add(mBase, uint32(v101))) = uint8(v102)
	v118 = int32(8)
	v119 = v94 + v118
	v121 = v98 + v118
	if v121 != v82&int32(2147483640) {
		v94 = v119
		v98 = v121
		goto L16
	} else {
		goto L18
	}
L17:
	;
	if v86 == int32(0) {
		goto L11
	} else {
		goto L19
	}
L18:
	;
	goto L17
L19:
	;
	v126 = v119
	goto L15
L20:
	;
	v141 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v88+v133*int32(_a_F_AbortSubTransaction_3)))) = uint8(v141)
	v143 = int32(1)
	v146 = v132 + v143
	if v146 != v86 {
		v132 = v146
		v133 = v133 + v143
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
	v175 = m.ExcPending
	if v175 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	F_reschedule_timeouts(m)
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	F_pgmem_sigprocmask(m, int32(_a_F_AbortSubTransaction_4), int32(0))
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	v182 = int32(10)
	goto L29
L27:
	;
	if v219 != 0 {
		goto L40
	} else {
		goto L41
	}
L28:
	;
	goto L27
L29:
	;
	v189 = *(*int32)(unsafe.Add(mBase, _c_F_AbortSubTransaction[23]))
	goto L32
L30:
	;
	v202 = int32(0)
	goto L37
L32:
	;
	goto L33
L33:
	;
	if int32(0)|base.B2i32(v189 == int32(15)) != 0 {
		goto L30
	} else {
		goto L35
	}
L35:
	;
	if v189 <= v182 {
		v219 = int32(1)
		goto L28
	} else {
		goto L36
	}
L36:
	;
	goto L30
L37:
	;
	v206 = *(*int32)(unsafe.Add(mBase, _c_F_AbortSubTransaction[24]))
	if v206 != int32(2) {
		v219 = v202
		goto L28
	} else {
		goto L38
	}
L38:
	;
	v210 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_AbortSubTransaction[25])))
	if v210&int32(1) != 0 {
		v219 = v202
		goto L28
	} else {
		goto L39
	}
L39:
	;
	v216 = *(*int32)(unsafe.Add(mBase, _c_F_AbortSubTransaction[26]))
	v219 = int32(0) | base.B2i32(v216 <= v182)
	goto L28
L40:
	;
	v223 = *(*int32)(unsafe.Add(mBase, _c_F_AbortSubTransaction[3]))
	F_ShowTransactionStateRec(m, int32(_a_F_AbortSubTransaction_5), v223)
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L1
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
	if v226 == int32(2) {
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
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v24)+60))
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v24)+64))
	*(*int32)(unsafe.Add(mBase, _c_F_AbortSubTransaction[27])) = v256
	*(*int32)(unsafe.Add(mBase, _c_F_AbortSubTransaction[28])) = v255
	goto L53
L45:
	;
	v231 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	if v231 == int32(0) {
		goto L44
	} else {
		goto L47
	}
L47:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
	if base.Ui32(v235) <= base.Ui32(int32(5)) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v235<<(uint(int32(2))%32))+uint32(_c_F_AbortSubTransaction[29])))
	v242 = v240
	goto L50
L49:
	;
	v242 = int32(_a_F_AbortSubTransaction_6)
	goto L50
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v242
	F_errmsg_internal(m, int32(_a_F_AbortSubTransaction_7), v10)
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	F_errfinish(m, int32(_a_F_AbortSubTransaction_8), int32(_a_F_AbortSubTransaction_9), int32(_a_F_AbortSubTransaction_5))
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	goto L44
L53:
	;
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v24)+28))
	v263 = *(*int32)(unsafe.Add(mBase, _c_F_AbortSubTransaction[30]))
	if v261 <= v263 {
		goto L55
	} else {
		goto L56
	}
L54:
	;
	v278 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_AbortSubTransaction[31])) = uint8(v278)
	*(*int32)(unsafe.Add(mBase, _c_F_AbortSubTransaction[32])) = v278
	goto L58
L55:
	;
	v266 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_AbortSubTransaction[33])) = v266
	*(*int32)(unsafe.Add(mBase, _c_F_AbortSubTransaction[34])) = v266
	*(*int32)(unsafe.Add(mBase, _c_F_AbortSubTransaction[35])) = v266
	*(*int32)(unsafe.Add(mBase, _c_F_AbortSubTransaction[30])) = v266
	goto L57
L56:
	;
	goto L57
L57:
	;
	goto L54
L58:
	;
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
	F_AtEOSubXact_Parallel(m, int32(0), v284)
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+72)) = int32(0)
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v24)+40))
	if v289 != 0 {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	F_AfterTriggerEndSubXact(m, int32(0))
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L1
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	v569 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+68)))
	*(*uint8)(unsafe.Add(mBase, _c_F_AbortSubTransaction[36])) = uint8(v569)
	v571 = int32(_a_F_AbortSubTransaction_0)
	v573 = *(*int32)(unsafe.Add(mBase, _c_F_AbortSubTransaction[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_AbortSubTransaction[2])) = v573 - int32(1)
	m.G0 = v10 + int32(16)
	return
L63:
	;
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v24)+80))
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v293)+8))
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v24)+40))
	F_AtSubAbort_Portals(m, v296, v294, v297)
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v24)+80))
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v302)+8))
	F_AtEOSubXact_LargeObject(m, int32(0), v301, v303)
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	v307 = *(*int32)(unsafe.Add(mBase, _c_F_AbortSubTransaction[3]))
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v307)+28))
	goto L66
L66:
	;
	goto L67
L67:
	;
	v317 = *(*int32)(unsafe.Add(mBase, _c_F_AbortSubTransaction[37]))
	if v317 == int32(0) {
		goto L69
	} else {
		goto L70
	}
L68:
	;
	v328 = *(*int32)(unsafe.Add(mBase, _c_F_AbortSubTransaction[38]))
	if v328 == int32(0) {
		goto L73
	} else {
		goto L74
	}
L69:
	;
	goto L68
L70:
	;
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v317)))
	if v320 < v308 {
		goto L69
	} else {
		goto L71
	}
L71:
	;
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v317)+8))
	*(*int32)(unsafe.Add(mBase, _c_F_AbortSubTransaction[37])) = v323
	F_pfree(m, v317)
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	goto L67
L73:
	;
	v355 = F_RecordTransactionAbort(m, int32(1))
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L1
	} else {
		goto L80
	}
L74:
	;
	v332 = v328
	goto L75
L75:
	;
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v332)))
	if v338 < v308 {
		goto L73
	} else {
		goto L77
	}
L76:
	;
	goto L73
L77:
	;
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v332)+12))
	*(*int32)(unsafe.Add(mBase, _c_F_AbortSubTransaction[38])) = v341
	F_pfree(m, v332)
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L1
	} else {
		goto L78
	}
L78:
	;
	v346 = *(*int32)(unsafe.Add(mBase, _c_F_AbortSubTransaction[38]))
	if v346 != 0 {
		v332 = v346
		goto L75
	} else {
		goto L79
	}
L79:
	;
	goto L76
L80:
	;
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	if v357 != 0 {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	v359 = *(*int32)(unsafe.Add(mBase, _c_F_AbortSubTransaction[3]))
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v359)+48))
	if v360 != 0 {
		goto L84
	} else {
		goto L85
	}
L82:
	;
	goto L83
L83:
	;
	v370 = *(*int32)(unsafe.Add(mBase, _c_F_AbortSubTransaction[39]))
	if v370 != 0 {
		goto L88
	} else {
		goto L89
	}
L84:
	;
	F_pfree(m, v360)
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L1
	} else {
		goto L87
	}
L85:
	;
	goto L86
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v359)+56)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v359)+48)) = int64(0)
	goto L83
L87:
	;
	goto L86
L88:
	;
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
	v372 = *(*int32)(unsafe.Add(mBase, uint32(v24)+80))
	v373 = *(*int32)(unsafe.Add(mBase, uint32(v372)+8))
	v375 = v370
	goto L91
L89:
	;
	goto L90
L90:
	;
	v394 = *(*int32)(unsafe.Add(mBase, uint32(v24)+40))
	v396 = int32(0)
	F_ResourceOwnerReleaseInternal(m, v394, int32(1), v396, v396)
	mBase = m.M
	v399 = m.ExcPending
	if v399 != 0 {
		goto L1
	} else {
		goto L95
	}
L91:
	;
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v375)))
	v383 = *(*int32)(unsafe.Add(mBase, uint32(v375)+8))
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v375)+4))
	m.T0[v384].(func(*base.Module, int32, int32, int32, int32))(m, int32(2), v371, v373, v383)
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L1
	} else {
		goto L93
	}
L92:
	;
	goto L90
L93:
	;
	if v381 != 0 {
		v375 = v381
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
	v401 = m.ExcPending
	if v401 != 0 {
		goto L1
	} else {
		goto L96
	}
L96:
	;
	v403 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
	v404 = *(*int32)(unsafe.Add(mBase, uint32(v24)+80))
	v405 = *(*int32)(unsafe.Add(mBase, uint32(v404)+8))
	F_AtEOSubXact_RelationCache(m, int32(0), v403, v405)
	mBase = m.M
	v407 = m.ExcPending
	if v407 != 0 {
		goto L1
	} else {
		goto L97
	}
L97:
	;
	F_AtEOXact_TypeCache(m)
	mBase = m.M
	v409 = m.ExcPending
	if v409 != 0 {
		goto L1
	} else {
		goto L98
	}
L98:
	;
	F_AtEOSubXact_Inval(m, int32(0))
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L1
	} else {
		goto L99
	}
L99:
	;
	v413 = *(*int32)(unsafe.Add(mBase, uint32(v24)+40))
	v415 = int32(0)
	F_ResourceOwnerReleaseInternal(m, v413, int32(2), v415, v415)
	mBase = m.M
	v418 = m.ExcPending
	if v418 != 0 {
		goto L1
	} else {
		goto L100
	}
L100:
	;
	v419 = *(*int32)(unsafe.Add(mBase, uint32(v24)+40))
	v421 = int32(0)
	F_ResourceOwnerReleaseInternal(m, v419, int32(3), v421, v421)
	mBase = m.M
	v424 = m.ExcPending
	if v424 != 0 {
		goto L1
	} else {
		goto L101
	}
L101:
	;
	F_smgrDoPendingDeletes(m, int32(0))
	mBase = m.M
	v427 = m.ExcPending
	if v427 != 0 {
		goto L1
	} else {
		goto L102
	}
L102:
	;
	v429 = *(*int32)(unsafe.Add(mBase, uint32(v24)+32))
	F_AtEOXact_GUC(m, int32(0), v429)
	mBase = m.M
	v431 = m.ExcPending
	if v431 != 0 {
		goto L1
	} else {
		goto L103
	}
L103:
	;
	v433 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
	F_AtEOSubXact_SPI(m, int32(0), v433)
	mBase = m.M
	v435 = m.ExcPending
	if v435 != 0 {
		goto L1
	} else {
		goto L104
	}
L104:
	;
	v437 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
	v438 = *(*int32)(unsafe.Add(mBase, uint32(v24)+80))
	v439 = *(*int32)(unsafe.Add(mBase, uint32(v438)+8))
	F_AtEOSubXact_on_commit_actions(m, int32(0), v437, v439)
	mBase = m.M
	v441 = m.ExcPending
	if v441 != 0 {
		goto L1
	} else {
		goto L105
	}
L105:
	;
	v443 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
	v447 = *(*int32)(unsafe.Add(mBase, _c_F_AbortSubTransaction[40]))
	if v443 == v447 {
		goto L107
	} else {
		goto L108
	}
L106:
	;
	v471 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
	v472 = *(*int32)(unsafe.Add(mBase, uint32(v24)+80))
	v473 = *(*int32)(unsafe.Add(mBase, uint32(v472)+8))
	F_AtEOSubXact_Files(m, int32(0), v471, v473)
	mBase = m.M
	v475 = m.ExcPending
	if v475 != 0 {
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
	v452 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_AbortSubTransaction[41])) = uint8(v452)
	v455 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_AbortSubTransaction[42])) = v455
	*(*int32)(unsafe.Add(mBase, _c_F_AbortSubTransaction[40])) = v455
	*(*int32)(unsafe.Add(mBase, _c_F_AbortSubTransaction[43])) = v455
	*(*uint8)(unsafe.Add(mBase, _c_F_AbortSubTransaction[44])) = uint8(v455)
	v467 = *(*int32)(unsafe.Add(mBase, _c_F_AbortSubTransaction[45]))
	*(*int32)(unsafe.Add(mBase, uint32(v467)+68)) = v455
	goto L109
L113:
	;
	v477 = *(*int32)(unsafe.Add(mBase, uint32(v24)+28))
	F_AtEOSubXact_HashTables(m, int32(0), v477)
	mBase = m.M
	v479 = m.ExcPending
	if v479 != 0 {
		goto L1
	} else {
		goto L114
	}
L114:
	;
	v481 = *(*int32)(unsafe.Add(mBase, uint32(v24)+28))
	F_AtEOSubXact_PgStat(m, int32(0), v481)
	mBase = m.M
	v483 = m.ExcPending
	if v483 != 0 {
		goto L1
	} else {
		goto L115
	}
L115:
	;
	v484 = *(*int32)(unsafe.Add(mBase, uint32(v24)+28))
	v487 = *(*int32)(unsafe.Add(mBase, _c_F_AbortSubTransaction[46]))
	if v487 != 0 {
		goto L117
	} else {
		goto L118
	}
L116:
	;
	goto L62
L117:
	;
	v490 = v487
	goto L120
L118:
	;
	goto L119
L119:
	;
	v523 = *(*int32)(unsafe.Add(mBase, _c_F_AbortSubTransaction[47]))
	if v523 != 0 {
		goto L129
	} else {
		goto L130
	}
L120:
	;
	v495 = *(*int32)(unsafe.Add(mBase, uint32(v490)+4))
	if v495 < v484 {
		goto L116
	} else {
		goto L122
	}
L121:
	;
	goto L119
L122:
	;
	v497 = *(*int32)(unsafe.Add(mBase, uint32(v490)+8))
	v498 = *(*int32)(unsafe.Add(mBase, uint32(v490)))
	v499 = *(*int32)(unsafe.Add(mBase, uint32(v498)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v498)+44)) = v499 - int32(1)
	v503 = *(*int32)(unsafe.Add(mBase, uint32(v490)))
	v504 = *(*int32)(unsafe.Add(mBase, uint32(v503)+44))
	if v504 != 0 {
		v510 = v490
		goto L123
	} else {
		goto L124
	}
L123:
	;
	F_pfree(m, v510)
	mBase = m.M
	v512 = m.ExcPending
	if v512 != 0 {
		goto L1
	} else {
		goto L127
	}
L124:
	;
	v505 = *(*int32)(unsafe.Add(mBase, uint32(v503)+48))
	if v505 != 0 {
		v510 = v490
		goto L123
	} else {
		goto L125
	}
L125:
	;
	F_pfree(m, v503)
	mBase = m.M
	v507 = m.ExcPending
	if v507 != 0 {
		goto L1
	} else {
		goto L126
	}
L126:
	;
	v509 = *(*int32)(unsafe.Add(mBase, _c_F_AbortSubTransaction[46]))
	v510 = v509
	goto L123
L127:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_AbortSubTransaction[46])) = v497
	if v497 != 0 {
		v490 = v497
		goto L120
	} else {
		goto L128
	}
L128:
	;
	goto L121
L129:
	;
	v525 = *(*int32)(unsafe.Add(mBase, _c_F_AbortSubTransaction[45]))
	v526 = *(*int32)(unsafe.Add(mBase, uint32(v525)+40))
	v528 = *(*int32)(unsafe.Add(mBase, _c_F_AbortSubTransaction[47]))
	v530 = v528 - int32(48)
	v531 = *(*int32)(unsafe.Add(mBase, uint32(v530)))
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v531))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v526)) == int32(0) {
		goto L133
	} else {
		goto L134
	}
L130:
	;
	v548 = int32(0)
	goto L131
L131:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_AbortSubTransaction[48])) = v548
	v552 = *(*int32)(unsafe.Add(mBase, _c_F_AbortSubTransaction[45]))
	*(*int32)(unsafe.Add(mBase, uint32(v552)+40)) = v548
	goto L116
L132:
	;
	if v543 == int32(0) {
		goto L116
	} else {
		goto L136
	}
L133:
	;
	v543 = base.B2i32(base.Ui32(v526) < base.Ui32(v531))
	goto L132
L134:
	;
	goto L135
L135:
	;
	v543 = int32(base.Ui32(v526-v531) >> (uint(int32(31)) % 32))
	goto L132
L136:
	;
	v546 = *(*int32)(unsafe.Add(mBase, uint32(v530)))
	v548 = v546
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
