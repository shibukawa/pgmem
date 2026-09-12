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
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v190 int32
	_ = v190
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v218 int32
	_ = v218
	var v230 int32
	_ = v230
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v274 int32
	_ = v274
	var v279 int32
	_ = v279
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v294 int32
	_ = v294
	var v306 int32
	_ = v306
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v317 int32
	_ = v317
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v345 int32
	_ = v345
	var v348 int32
	_ = v348
	var v351 int32
	_ = v351
	var v354 int32
	_ = v354
	var v356 int32
	_ = v356
	var v360 int32
	_ = v360
	var v366 int32
	_ = v366
	var v369 int32
	_ = v369
	var v372 int32
	_ = v372
	var v374 int32
	_ = v374
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v390 int32
	_ = v390
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v403 int32
	_ = v403
	var v409 int32
	_ = v409
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v414 int32
	_ = v414
	var v422 int32
	_ = v422
	var v424 int32
	_ = v424
	var v427 int32
	_ = v427
	var v429 int32
	_ = v429
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v435 int32
	_ = v435
	var v437 int32
	_ = v437
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v443 int32
	_ = v443
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v449 int32
	_ = v449
	var v452 int32
	_ = v452
	var v455 int32
	_ = v455
	var v457 int32
	_ = v457
	var v459 int32
	_ = v459
	var v461 int32
	_ = v461
	var v463 int32
	_ = v463
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v469 int32
	_ = v469
	var v471 int32
	_ = v471
	var v475 int32
	_ = v475
	var v480 int32
	_ = v480
	var v483 int32
	_ = v483
	var v495 int32
	_ = v495
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v503 int32
	_ = v503
	var v505 int32
	_ = v505
	var v507 int32
	_ = v507
	var v509 int32
	_ = v509
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v515 int32
	_ = v515
	var v518 int32
	_ = v518
	var v523 int32
	_ = v523
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v535 int32
	_ = v535
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v540 int32
	_ = v540
	var v551 int32
	_ = v551
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v556 int32
	_ = v556
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v571 int32
	_ = v571
	var v574 int32
	_ = v574
	var v576 int32
	_ = v576
	var v580 int32
	_ = v580
	var v597 int32
	_ = v597
	var v599 int32
	_ = v599
	var v601 int32
	_ = v601
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v14 = *(*int32)(unsafe.Add(mBase, _consts[202]))
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v14
	v16 = int32(4465212)
	v18 = *(*int32)(unsafe.Add(mBase, _consts[162]))
	*(*int32)(unsafe.Add(mBase, _consts[162])) = v18 + int32(1)
	v24 = *(*int32)(unsafe.Add(mBase, _consts[25]))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+40))
	*(*int32)(unsafe.Add(mBase, _consts[173])) = v25
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
	v30 = *(*int32)(unsafe.Add(mBase, _consts[148]))
	v31 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v30))) = v31
	v35 = *(*int32)(unsafe.Add(mBase, _consts[55]))
	if v35 == v31 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	F_pgaio_error_cleanup(m)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L1
	} else {
		goto L8
	}
L4:
	;
	goto L3
L5:
	;
	v39 = int32(*(*uint8)(unsafe.Add(mBase, _consts[56])))
	if v39 != int32(1) {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v35)+220))
	if v42 == int32(0) {
		goto L4
	} else {
		goto L7
	}
L7:
	;
	v45 = int32(4465220)
	v47 = *(*int32)(unsafe.Add(mBase, _consts[26]))
	v48 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[26])) = v47 + v48
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	*(*int32)(unsafe.Add(mBase, uint32(v35))) = v51 + v48
	v55 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v35)+220)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v35)+224)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v35))) = v51 + int32(2)
	v65 = *(*int32)(unsafe.Add(mBase, _consts[26]))
	*(*int32)(unsafe.Add(mBase, _consts[26])) = v65 - v48
	goto L4
L8:
	;
	F_UnlockBuffers(m)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	v73 = int32(0)
	v79 = *(*int32)(unsafe.Add(mBase, _consts[130]))
	if v79 <= v73 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	F_ConditionVariableCancelSleep(m)
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L1
	} else {
		goto L23
	}
L11:
	;
	*(*int64)(unsafe.Add(mBase, _consts[221])) = int64(0)
	*(*int32)(unsafe.Add(mBase, _consts[222])) = int32(4366400)
	v190 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[130])) = v190
	*(*int32)(unsafe.Add(mBase, _consts[223])) = v190
	*(*uint8)(unsafe.Add(mBase, _consts[94])) = uint8(v190)
	*(*uint8)(unsafe.Add(mBase, _consts[224])) = uint8(v190)
	goto L10
L12:
	;
	v83 = v79 & int32(7)
	v85 = *(*int32)(unsafe.Add(mBase, _consts[132]))
	if base.Ui32(int32(8)) <= base.Ui32(v79) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v91 = v73
	v92 = int32(0)
	goto L16
L14:
	;
	v155 = v73
	goto L15
L15:
	;
	if v83 == int32(0) {
		goto L11
	} else {
		goto L19
	}
L16:
	;
	v96 = int32(8260)
	v99 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v85+v91*v96))) = uint8(v99)
	*(*uint8)(unsafe.Add(mBase, uint32(v85+(v91|int32(1))*v96))) = uint8(v99)
	*(*uint8)(unsafe.Add(mBase, uint32(v85+(v91|int32(2))*v96))) = uint8(v99)
	*(*uint8)(unsafe.Add(mBase, uint32(v85+(v91|int32(3))*v96))) = uint8(v99)
	*(*uint8)(unsafe.Add(mBase, uint32(v85+(v91|int32(4))*v96))) = uint8(v99)
	*(*uint8)(unsafe.Add(mBase, uint32(v85+(v91|int32(5))*v96))) = uint8(v99)
	*(*uint8)(unsafe.Add(mBase, uint32(v85+(v91|int32(6))*v96))) = uint8(v99)
	*(*uint8)(unsafe.Add(mBase, uint32(v85+(v91|int32(7))*v96))) = uint8(v99)
	v150 = int32(8)
	v151 = v91 + v150
	v153 = v92 + v150
	if v153 != v79&int32(2147483640) {
		v91 = v151
		v92 = v153
		goto L16
	} else {
		goto L18
	}
L17:
	;
	v155 = v151
	goto L15
L18:
	;
	goto L17
L19:
	;
	v163 = v155
	v164 = int32(0)
	goto L20
L20:
	;
	v171 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v85+v163*int32(8260)))) = uint8(v171)
	v173 = int32(1)
	v176 = v164 + v173
	if v176 != v83 {
		v163 = v163 + v173
		v164 = v176
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
	v204 = m.ExcPending
	if v204 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	F_reschedule_timeouts(m)
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	F_sigprocmask(m, int32(4377784), int32(0))
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	v211 = int32(10)
	goto L29
L27:
	;
	if v245 != 0 {
		goto L41
	} else {
		goto L42
	}
L28:
	;
	goto L27
L29:
	;
	v218 = *(*int32)(unsafe.Add(mBase, _consts[215]))
	goto L32
L30:
	;
	v230 = int32(0)
	goto L38
L32:
	;
	goto L33
L33:
	;
	goto L35
L35:
	;
	if v218 == int32(15) {
		goto L30
	} else {
		goto L36
	}
L36:
	;
	if v218 <= v211 {
		v245 = int32(1)
		goto L28
	} else {
		goto L37
	}
L37:
	;
	goto L30
L38:
	;
	v234 = *(*int32)(unsafe.Add(mBase, _consts[216]))
	if v234 != int32(2) {
		v245 = v230
		goto L28
	} else {
		goto L39
	}
L39:
	;
	v238 = int32(*(*uint8)(unsafe.Add(mBase, _consts[217])))
	if v238 != 0 {
		v245 = v230
		goto L28
	} else {
		goto L40
	}
L40:
	;
	v242 = *(*int32)(unsafe.Add(mBase, _consts[218]))
	v245 = int32(0) | base.B2i32(v242 <= v211)
	goto L28
L41:
	;
	v249 = *(*int32)(unsafe.Add(mBase, _consts[25]))
	F_ShowTransactionStateRec(m, int32(254133), v249)
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L1
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
	if v252 == int32(2) {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	goto L43
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+20)) = int32(4)
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v24)+60))
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v24)+64))
	*(*int32)(unsafe.Add(mBase, _consts[187])) = v284
	*(*int32)(unsafe.Add(mBase, _consts[159])) = v283
	goto L54
L46:
	;
	v257 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	if v257 == int32(0) {
		goto L45
	} else {
		goto L48
	}
L48:
	;
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
	if base.Ui32(v261) <= base.Ui32(int32(5)) {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v261<<(uint(int32(2))%32))+uint32(_consts[220])))
	v270 = v268
	goto L51
L50:
	;
	v270 = int32(534390)
	goto L51
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v270
	F_errmsg_internal(m, int32(346356), v10)
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	F_errfinish(m, int32(485028), int32(5283), int32(254133))
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	goto L45
L54:
	;
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v24)+28))
	v291 = *(*int32)(unsafe.Add(mBase, _consts[225]))
	if v289 <= v291 {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	v306 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _consts[116])) = uint8(v306)
	*(*int32)(unsafe.Add(mBase, _consts[115])) = v306
	goto L59
L56:
	;
	v294 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[113])) = v294
	*(*int32)(unsafe.Add(mBase, _consts[226])) = v294
	*(*int32)(unsafe.Add(mBase, _consts[114])) = v294
	*(*int32)(unsafe.Add(mBase, _consts[225])) = v294
	goto L58
L57:
	;
	goto L58
L58:
	;
	goto L55
L59:
	;
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
	F_AtEOSubXact_Parallel(m, int32(0), v312)
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+72)) = int32(0)
	v317 = *(*int32)(unsafe.Add(mBase, uint32(v24)+40))
	if v317 != 0 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	F_AfterTriggerEndSubXact(m, int32(0))
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L1
	} else {
		goto L64
	}
L62:
	;
	goto L63
L63:
	;
	v597 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+68)))
	*(*uint8)(unsafe.Add(mBase, _consts[198])) = uint8(v597)
	v599 = int32(4465212)
	v601 = *(*int32)(unsafe.Add(mBase, _consts[162]))
	*(*int32)(unsafe.Add(mBase, _consts[162])) = v601 - int32(1)
	m.G0 = v10 + int32(16)
	return
L64:
	;
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v24)+80))
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v321)+8))
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v24)+40))
	F_AtSubAbort_Portals(m, v324, v322, v325)
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
	v330 = *(*int32)(unsafe.Add(mBase, uint32(v24)+80))
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v330)+8))
	F_AtEOSubXact_LargeObject(m, int32(0), v329, v331)
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	v335 = *(*int32)(unsafe.Add(mBase, _consts[25]))
	v336 = *(*int32)(unsafe.Add(mBase, uint32(v335)+28))
	goto L67
L67:
	;
	goto L68
L68:
	;
	v345 = *(*int32)(unsafe.Add(mBase, _consts[231]))
	if v345 == int32(0) {
		goto L70
	} else {
		goto L71
	}
L69:
	;
	v356 = *(*int32)(unsafe.Add(mBase, _consts[230]))
	if v356 == int32(0) {
		goto L74
	} else {
		goto L75
	}
L70:
	;
	goto L69
L71:
	;
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v345)))
	if v348 < v336 {
		goto L70
	} else {
		goto L72
	}
L72:
	;
	v351 = *(*int32)(unsafe.Add(mBase, uint32(v345)+8))
	*(*int32)(unsafe.Add(mBase, _consts[231])) = v351
	F_pfree(m, v345)
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L1
	} else {
		goto L73
	}
L73:
	;
	goto L68
L74:
	;
	v383 = F_RecordTransactionAbort(m, int32(1))
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L1
	} else {
		goto L81
	}
L75:
	;
	v360 = v356
	goto L76
L76:
	;
	v366 = *(*int32)(unsafe.Add(mBase, uint32(v360)))
	if v366 < v336 {
		goto L74
	} else {
		goto L78
	}
L77:
	;
	goto L74
L78:
	;
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v360)+12))
	*(*int32)(unsafe.Add(mBase, _consts[230])) = v369
	F_pfree(m, v360)
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L1
	} else {
		goto L79
	}
L79:
	;
	v374 = *(*int32)(unsafe.Add(mBase, _consts[230]))
	if v374 != 0 {
		v360 = v374
		goto L76
	} else {
		goto L80
	}
L80:
	;
	goto L77
L81:
	;
	v385 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	if v385 != 0 {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v387 = *(*int32)(unsafe.Add(mBase, _consts[25]))
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v387)+48))
	if v388 != 0 {
		goto L85
	} else {
		goto L86
	}
L83:
	;
	goto L84
L84:
	;
	v398 = *(*int32)(unsafe.Add(mBase, _consts[247]))
	if v398 != 0 {
		goto L89
	} else {
		goto L90
	}
L85:
	;
	F_pfree(m, v388)
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L1
	} else {
		goto L88
	}
L86:
	;
	goto L87
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v387)+56)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v387)+48)) = int64(0)
	goto L84
L88:
	;
	goto L87
L89:
	;
	v399 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
	v400 = *(*int32)(unsafe.Add(mBase, uint32(v24)+80))
	v401 = *(*int32)(unsafe.Add(mBase, uint32(v400)+8))
	v403 = v398
	goto L92
L90:
	;
	goto L91
L91:
	;
	v422 = *(*int32)(unsafe.Add(mBase, uint32(v24)+40))
	v424 = int32(0)
	F_ResourceOwnerRelease(m, v422, int32(1), v424, v424)
	mBase = m.M
	v427 = m.ExcPending
	if v427 != 0 {
		goto L1
	} else {
		goto L96
	}
L92:
	;
	v409 = *(*int32)(unsafe.Add(mBase, uint32(v403)))
	v411 = *(*int32)(unsafe.Add(mBase, uint32(v403)+8))
	v412 = *(*int32)(unsafe.Add(mBase, uint32(v403)+4))
	m.T0[v412].(func(*base.Module, int32, int32, int32, int32))(m, int32(2), v399, v401, v411)
	mBase = m.M
	v414 = m.ExcPending
	if v414 != 0 {
		goto L1
	} else {
		goto L94
	}
L93:
	;
	goto L91
L94:
	;
	if v409 != 0 {
		v403 = v409
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
	v429 = m.ExcPending
	if v429 != 0 {
		goto L1
	} else {
		goto L97
	}
L97:
	;
	v431 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
	v432 = *(*int32)(unsafe.Add(mBase, uint32(v24)+80))
	v433 = *(*int32)(unsafe.Add(mBase, uint32(v432)+8))
	F_AtEOSubXact_RelationCache(m, int32(0), v431, v433)
	mBase = m.M
	v435 = m.ExcPending
	if v435 != 0 {
		goto L1
	} else {
		goto L98
	}
L98:
	;
	F_AtEOXact_TypeCache(m)
	mBase = m.M
	v437 = m.ExcPending
	if v437 != 0 {
		goto L1
	} else {
		goto L99
	}
L99:
	;
	F_AtEOSubXact_Inval(m, int32(0))
	mBase = m.M
	v440 = m.ExcPending
	if v440 != 0 {
		goto L1
	} else {
		goto L100
	}
L100:
	;
	v441 = *(*int32)(unsafe.Add(mBase, uint32(v24)+40))
	v443 = int32(0)
	F_ResourceOwnerRelease(m, v441, int32(2), v443, v443)
	mBase = m.M
	v446 = m.ExcPending
	if v446 != 0 {
		goto L1
	} else {
		goto L101
	}
L101:
	;
	v447 = *(*int32)(unsafe.Add(mBase, uint32(v24)+40))
	v449 = int32(0)
	F_ResourceOwnerRelease(m, v447, int32(3), v449, v449)
	mBase = m.M
	v452 = m.ExcPending
	if v452 != 0 {
		goto L1
	} else {
		goto L102
	}
L102:
	;
	F_smgrDoPendingDeletes(m, int32(0))
	mBase = m.M
	v455 = m.ExcPending
	if v455 != 0 {
		goto L1
	} else {
		goto L103
	}
L103:
	;
	v457 = *(*int32)(unsafe.Add(mBase, uint32(v24)+32))
	F_AtEOXact_GUC(m, int32(0), v457)
	mBase = m.M
	v459 = m.ExcPending
	if v459 != 0 {
		goto L1
	} else {
		goto L104
	}
L104:
	;
	v461 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
	F_AtEOSubXact_SPI(m, int32(0), v461)
	mBase = m.M
	v463 = m.ExcPending
	if v463 != 0 {
		goto L1
	} else {
		goto L105
	}
L105:
	;
	v465 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
	v466 = *(*int32)(unsafe.Add(mBase, uint32(v24)+80))
	v467 = *(*int32)(unsafe.Add(mBase, uint32(v466)+8))
	F_AtEOSubXact_on_commit_actions(m, int32(0), v465, v467)
	mBase = m.M
	v469 = m.ExcPending
	if v469 != 0 {
		goto L1
	} else {
		goto L106
	}
L106:
	;
	v471 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
	v475 = *(*int32)(unsafe.Add(mBase, _consts[249]))
	if v471 == v475 {
		goto L108
	} else {
		goto L109
	}
L107:
	;
	v499 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
	v500 = *(*int32)(unsafe.Add(mBase, uint32(v24)+80))
	v501 = *(*int32)(unsafe.Add(mBase, uint32(v500)+8))
	F_AtEOSubXact_Files(m, int32(0), v499, v501)
	mBase = m.M
	v503 = m.ExcPending
	if v503 != 0 {
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
	v480 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[250])) = uint8(v480)
	v483 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[251])) = v483
	*(*int32)(unsafe.Add(mBase, _consts[249])) = v483
	*(*int32)(unsafe.Add(mBase, _consts[252])) = v483
	*(*uint8)(unsafe.Add(mBase, _consts[253])) = uint8(v483)
	v495 = *(*int32)(unsafe.Add(mBase, _consts[120]))
	*(*int32)(unsafe.Add(mBase, uint32(v495)+68)) = v483
	goto L110
L114:
	;
	v505 = *(*int32)(unsafe.Add(mBase, uint32(v24)+28))
	F_AtEOSubXact_HashTables(m, int32(0), v505)
	mBase = m.M
	v507 = m.ExcPending
	if v507 != 0 {
		goto L1
	} else {
		goto L115
	}
L115:
	;
	v509 = *(*int32)(unsafe.Add(mBase, uint32(v24)+28))
	F_AtEOSubXact_PgStat(m, int32(0), v509)
	mBase = m.M
	v511 = m.ExcPending
	if v511 != 0 {
		goto L1
	} else {
		goto L116
	}
L116:
	;
	v512 = *(*int32)(unsafe.Add(mBase, uint32(v24)+28))
	v515 = *(*int32)(unsafe.Add(mBase, _consts[254]))
	if v515 != 0 {
		goto L118
	} else {
		goto L119
	}
L117:
	;
	goto L63
L118:
	;
	v518 = v515
	goto L121
L119:
	;
	goto L120
L120:
	;
	v551 = *(*int32)(unsafe.Add(mBase, _consts[255]))
	if v551 != 0 {
		goto L130
	} else {
		goto L131
	}
L121:
	;
	v523 = *(*int32)(unsafe.Add(mBase, uint32(v518)+4))
	if v523 < v512 {
		goto L117
	} else {
		goto L123
	}
L122:
	;
	goto L120
L123:
	;
	v525 = *(*int32)(unsafe.Add(mBase, uint32(v518)+8))
	v526 = *(*int32)(unsafe.Add(mBase, uint32(v518)))
	v527 = *(*int32)(unsafe.Add(mBase, uint32(v526)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v526)+44)) = v527 - int32(1)
	v531 = *(*int32)(unsafe.Add(mBase, uint32(v518)))
	v532 = *(*int32)(unsafe.Add(mBase, uint32(v531)+44))
	if v532 != 0 {
		v538 = v518
		goto L124
	} else {
		goto L125
	}
L124:
	;
	F_pfree(m, v538)
	mBase = m.M
	v540 = m.ExcPending
	if v540 != 0 {
		goto L1
	} else {
		goto L128
	}
L125:
	;
	v533 = *(*int32)(unsafe.Add(mBase, uint32(v531)+48))
	if v533 != 0 {
		v538 = v518
		goto L124
	} else {
		goto L126
	}
L126:
	;
	F_pfree(m, v531)
	mBase = m.M
	v535 = m.ExcPending
	if v535 != 0 {
		goto L1
	} else {
		goto L127
	}
L127:
	;
	v537 = *(*int32)(unsafe.Add(mBase, _consts[254]))
	v538 = v537
	goto L124
L128:
	;
	*(*int32)(unsafe.Add(mBase, _consts[254])) = v525
	if v525 != 0 {
		v518 = v525
		goto L121
	} else {
		goto L129
	}
L129:
	;
	goto L122
L130:
	;
	v553 = *(*int32)(unsafe.Add(mBase, _consts[120]))
	v554 = *(*int32)(unsafe.Add(mBase, uint32(v553)+40))
	v556 = *(*int32)(unsafe.Add(mBase, _consts[255]))
	v558 = v556 - int32(48)
	v559 = *(*int32)(unsafe.Add(mBase, uint32(v558)))
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v559))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v554)) == int32(0) {
		goto L134
	} else {
		goto L135
	}
L131:
	;
	v576 = int32(0)
	goto L132
L132:
	;
	*(*int32)(unsafe.Add(mBase, _consts[83])) = v576
	v580 = *(*int32)(unsafe.Add(mBase, _consts[120]))
	*(*int32)(unsafe.Add(mBase, uint32(v580)+40)) = v576
	goto L117
L133:
	;
	if v571 == int32(0) {
		goto L117
	} else {
		goto L137
	}
L134:
	;
	v571 = base.B2i32(base.Ui32(v554) < base.Ui32(v559))
	goto L133
L135:
	;
	goto L136
L136:
	;
	v571 = int32(base.Ui32(v554-v559) >> (uint(int32(31)) % 32))
	goto L133
L137:
	;
	v574 = *(*int32)(unsafe.Add(mBase, uint32(v558)))
	v576 = v574
	goto L132
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
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v174 int32
	_ = v174
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v11 = *(*int32)(unsafe.Add(mBase, _consts[25]))
	v12 = int32(10)
	goto L3
L1:
	;
	if v46 != 0 {
		goto L15
	} else {
		goto L16
	}
L2:
	;
	goto L1
L3:
	;
	v19 = *(*int32)(unsafe.Add(mBase, _consts[215]))
	goto L6
L4:
	;
	v31 = int32(0)
	goto L12
L6:
	;
	goto L7
L7:
	;
	goto L9
L9:
	;
	if v19 == int32(15) {
		goto L4
	} else {
		goto L10
	}
L10:
	;
	if v19 <= v12 {
		v46 = int32(1)
		goto L2
	} else {
		goto L11
	}
L11:
	;
	goto L4
L12:
	;
	v35 = *(*int32)(unsafe.Add(mBase, _consts[216]))
	if v35 != int32(2) {
		v46 = v31
		goto L2
	} else {
		goto L13
	}
L13:
	;
	v39 = int32(*(*uint8)(unsafe.Add(mBase, _consts[217])))
	if v39 != 0 {
		v46 = v31
		goto L2
	} else {
		goto L14
	}
L14:
	;
	v43 = *(*int32)(unsafe.Add(mBase, _consts[218]))
	v46 = int32(0) | base.B2i32(v43 <= v12)
	goto L2
L15:
	;
	v50 = *(*int32)(unsafe.Add(mBase, _consts[25]))
	F_ShowTransactionStateRec(m, int32(254234), v50)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	goto L17
L17:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v11)+20))
	if v53 == int32(4) {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	return
L19:
	;
	goto L17
L20:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	v83 = m.G0
	v85 = v83 - int32(32)
	m.G0 = v85
	v90 = *(*int32)(unsafe.Add(mBase, _consts[219]))
	F_hash_seq_init(m, v85+int32(12), v90)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L18
	} else {
		goto L29
	}
L21:
	;
	v58 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L18
	} else {
		goto L22
	}
L22:
	;
	if v58 == int32(0) {
		goto L20
	} else {
		goto L23
	}
L23:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v11)+20))
	if base.Ui32(v62) <= base.Ui32(int32(5)) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v62<<(uint(int32(2))%32))+uint32(_consts[220])))
	v71 = v69
	goto L26
L25:
	;
	v71 = int32(534390)
	goto L26
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v71
	F_errmsg_internal(m, int32(346471), v8)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L18
	} else {
		goto L27
	}
L27:
	;
	F_errfinish(m, int32(485028), int32(5391), int32(254234))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L18
	} else {
		goto L28
	}
L28:
	;
	goto L20
L29:
	;
	v95 = F_hash_seq_search(m, v85+int32(12))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L18
	} else {
		goto L30
	}
L30:
	;
	if v95 != 0 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v97 = v95
	goto L34
L32:
	;
	goto L33
L33:
	;
	m.G0 = v85 + int32(32)
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v11)+80))
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v143)+40))
	*(*int32)(unsafe.Add(mBase, _consts[204])) = v144
	*(*int32)(unsafe.Add(mBase, _consts[173])) = v144
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v11)+40))
	if v148 != 0 {
		goto L54
	} else {
		goto L55
	}
L34:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v97)+64))
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v102)+20))
	if v82 == v103 {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	goto L33
L36:
	;
	v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102)+84)))
	if v105 == int32(1) {
		goto L39
	} else {
		goto L40
	}
L37:
	;
	goto L38
L38:
	;
	v132 = F_hash_seq_search(m, v85+int32(12))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L18
	} else {
		goto L52
	}
L39:
	;
	v108 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v102)+84)) = uint8(v108)
	goto L41
L40:
	;
	goto L41
L41:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v102)+16))
	if v110 != 0 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v113 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L18
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	F_PortalDrop(m, v102, int32(0))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L18
	} else {
		goto L51
	}
L45:
	;
	if v113 != 0 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v102)))
	*(*int32)(unsafe.Add(mBase, uint32(v85))) = v115
	F_errmsg_internal(m, int32(686971), v85)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L18
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v102)+16)) = int32(0)
	goto L44
L49:
	;
	F_errfinish(m, int32(488652), int32(1120), int32(151026))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L18
	} else {
		goto L50
	}
L50:
	;
	goto L48
L51:
	;
	goto L38
L52:
	;
	if v132 != 0 {
		v97 = v132
		goto L34
	} else {
		goto L53
	}
L53:
	;
	goto L35
L54:
	;
	F_ResourceOwnerDelete(m, v148)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L18
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+40)) = int32(0)
	v155 = *(*int32)(unsafe.Add(mBase, _consts[25]))
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v155)+44))
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v156
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v155)+80))
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v159)+36))
	*(*int32)(unsafe.Add(mBase, _consts[203])) = v160
	v163 = *(*int32)(unsafe.Add(mBase, _consts[202]))
	if v163 != 0 {
		goto L58
	} else {
		goto L59
	}
L57:
	;
	goto L56
L58:
	;
	F_MemoryContextReset(m, v163)
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L18
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v155)+36))
	if v166 != 0 {
		goto L62
	} else {
		goto L63
	}
L61:
	;
	goto L60
L62:
	;
	F_MemoryContextDelete(m, v166)
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L18
	} else {
		goto L65
	}
L63:
	;
	goto L64
L64:
	;
	v169 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v155)+36)) = v169
	*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = v169
	F_PopTransaction(m)
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L18
	} else {
		goto L66
	}
L65:
	;
	goto L64
L66:
	;
	m.G0 = v8 + int32(16)
	return
}
func F_StartSubTransaction(m *base.Module) {
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
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
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
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v126 int64
	_ = v126
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v163 int32
	_ = v163
	var v170 int32
	_ = v170
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v11 = *(*int32)(unsafe.Add(mBase, _consts[25]))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+20))
	if v12 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = int32(1)
	v44 = *(*int32)(unsafe.Add(mBase, _consts[25]))
	v46 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+44)) = v46
	v48 = int32(4470588)
	v50 = *(*int32)(unsafe.Add(mBase, _consts[203]))
	v55 = F_AllocSetContextCreateInternal(m, v50, int32(61420), int32(0), int32(8192), int32(8388608))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L3
	} else {
		goto L11
	}
L2:
	;
	v17 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return
L4:
	;
	if v17 == int32(0) {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v11)+20))
	if base.Ui32(v21) <= base.Ui32(int32(5)) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v21<<(uint(int32(2))%32))+uint32(_consts[220])))
	v30 = v28
	goto L8
L7:
	;
	v30 = int32(534390)
	goto L8
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v30
	F_errmsg_internal(m, int32(346394), v8)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L3
	} else {
		goto L9
	}
L9:
	;
	F_errfinish(m, int32(485028), int32(5073), int32(254153))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L3
	} else {
		goto L10
	}
L10:
	;
	goto L1
L11:
	;
	*(*int32)(unsafe.Add(mBase, _consts[203])) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v44)+36)) = v55
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v55
	v62 = *(*int32)(unsafe.Add(mBase, _consts[25]))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)+80))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)+40))
	v66 = F_ResourceOwnerCreate(m, v64, int32(254269))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L3
	} else {
		goto L12
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v62)+40)) = v66
	*(*int32)(unsafe.Add(mBase, _consts[173])) = v66
	*(*int32)(unsafe.Add(mBase, _consts[204])) = v66
	v74 = *(*int32)(unsafe.Add(mBase, _consts[25]))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)+28))
	goto L14
L13:
	;
	v115 = v75 * int32(24)
	*(*int32)(unsafe.Add(mBase, uint32(v110+v115))) = int32(0)
	v119 = int32(4367660)
	v120 = *(*int32)(unsafe.Add(mBase, _consts[244]))
	v121 = v120 + v115
	v123 = *(*int32)(unsafe.Add(mBase, _consts[245]))
	*(*int32)(unsafe.Add(mBase, uint32(v121)+12)) = v123
	v126 = *(*int64)(unsafe.Add(mBase, _consts[246]))
	*(*int64)(unsafe.Add(mBase, uint32(v121)+4)) = v126
	v129 = *(*int32)(unsafe.Add(mBase, _consts[244]))
	v132 = *(*int32)(unsafe.Add(mBase, _consts[212]))
	*(*int32)(unsafe.Add(mBase, uint32(v129+v115)+16)) = v132
	v135 = *(*int32)(unsafe.Add(mBase, _consts[244]))
	v138 = *(*int32)(unsafe.Add(mBase, _consts[213]))
	*(*int32)(unsafe.Add(mBase, uint32(v135+v115)+20)) = v138
	*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = int32(2)
	v143 = *(*int32)(unsafe.Add(mBase, _consts[247]))
	if v143 != 0 {
		goto L27
	} else {
		goto L28
	}
L14:
	;
	v77 = *(*int32)(unsafe.Add(mBase, _consts[248]))
	if v77 <= v75 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v79 = v77
	goto L18
L16:
	;
	goto L17
L17:
	;
	v108 = *(*int32)(unsafe.Add(mBase, _consts[244]))
	v110 = v108
	goto L13
L18:
	;
	if v79 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, _consts[248])) = v100
	*(*int32)(unsafe.Add(mBase, _consts[244])) = v101
	if v100 <= v75 {
		v79 = v100
		goto L18
	} else {
		goto L26
	}
L21:
	;
	v88 = *(*int32)(unsafe.Add(mBase, _consts[144]))
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
	v93 = *(*int32)(unsafe.Add(mBase, _consts[244]))
	v96 = F_repalloc(m, v93, v79*int32(48))
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
	v100 = v79 << (uint(int32(1)) % 32)
	v101 = v96
	goto L20
L26:
	;
	v110 = v101
	goto L13
L27:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v11)+80))
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v145)+8))
	v147 = v143
	goto L30
L28:
	;
	goto L29
L29:
	;
	v163 = int32(10)
	goto L36
L30:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v147)))
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v147)+8))
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v147)+4))
	m.T0[v155].(func(*base.Module, int32, int32, int32, int32))(m, int32(0), v144, v146, v154)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L3
	} else {
		goto L32
	}
L31:
	;
	goto L29
L32:
	;
	if v152 != 0 {
		v147 = v152
		goto L30
	} else {
		goto L33
	}
L33:
	;
	goto L31
L34:
	;
	if v197 != 0 {
		goto L48
	} else {
		goto L49
	}
L35:
	;
	goto L34
L36:
	;
	v170 = *(*int32)(unsafe.Add(mBase, _consts[215]))
	goto L39
L37:
	;
	v182 = int32(0)
	goto L45
L39:
	;
	goto L40
L40:
	;
	goto L42
L42:
	;
	if v170 == int32(15) {
		goto L37
	} else {
		goto L43
	}
L43:
	;
	if v170 <= v163 {
		v197 = int32(1)
		goto L35
	} else {
		goto L44
	}
L44:
	;
	goto L37
L45:
	;
	v186 = *(*int32)(unsafe.Add(mBase, _consts[216]))
	if v186 != int32(2) {
		v197 = v182
		goto L35
	} else {
		goto L46
	}
L46:
	;
	v190 = int32(*(*uint8)(unsafe.Add(mBase, _consts[217])))
	if v190 != 0 {
		v197 = v182
		goto L35
	} else {
		goto L47
	}
L47:
	;
	v194 = *(*int32)(unsafe.Add(mBase, _consts[218]))
	v197 = int32(0) | base.B2i32(v194 <= v163)
	goto L35
L48:
	;
	v201 = *(*int32)(unsafe.Add(mBase, _consts[25]))
	F_ShowTransactionStateRec(m, int32(254153), v201)
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L3
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	m.G0 = v8 + int32(16)
	return
L51:
	;
	goto L50
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
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v8 + int32(16)
	return v111
L2:
	;
	v111 = v2
	goto L1
L3:
	;
	goto L4
L4:
	;
	v12 = l0
	goto L6
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L16
	} else {
		goto L30
	}
L6:
	;
	v18 = *(*int32)(unsafe.Add(mBase, _consts[83]))
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v18))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v12)) == int32(0) {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	v79 = int32(0)
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v12))&int32(0) == v79 {
		goto L26
	} else {
		goto L27
	}
L8:
	;
	if v30 != 0 {
		v111 = v12
		goto L1
	} else {
		goto L12
	}
L9:
	;
	v30 = base.B2i32(base.Ui32(v12) < base.Ui32(v18))
	goto L8
L10:
	;
	goto L11
L11:
	;
	v30 = int32(base.Ui32(v12-v18) >> (uint(int32(31)) % 32))
	goto L8
L12:
	;
	if base.Ui32(int32(3)) <= base.Ui32(v12) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v35 = int32(base.Ui32(v12) >> (uint(int32(11)) % 32))
	v37 = F_SimpleLruReadPage_ReadOnly(m, int32(4365464), base.I64_extend_i32_u(v35), v12)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	goto L15
L15:
	;
	goto L7
L16:
	;
	return int32(0)
L17:
	;
	v42 = *(*int32)(unsafe.Add(mBase, _consts[154]))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
	v44 = int32(2)
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v43+v37<<(uint(v44)%32))))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v47+v12&int32(2047)<<(uint(v44)%32))))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v42)+28))
	v56 = int32(*(*uint16)(unsafe.Add(mBase, _consts[155])))
	v57 = base.I32_rem_u_s(v35, v56)
	F_LWLockRelease(m, v54+v57<<(uint(int32(7))%32))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L16
	} else {
		goto L18
	}
L18:
	;
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v12))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v53)) == int32(0) {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	if v74 == int32(0) {
		v93 = v53
		goto L5
	} else {
		goto L23
	}
L20:
	;
	v74 = base.B2i32(base.Ui32(v53) < base.Ui32(v12))
	goto L19
L21:
	;
	goto L22
L22:
	;
	v74 = int32(base.Ui32(v53-v12) >> (uint(int32(31)) % 32))
	goto L19
L23:
	;
	if v53 == int32(0) {
		v111 = v12
		goto L1
	} else {
		goto L24
	}
L24:
	;
	v12 = v53
	goto L6
L25:
	;
	if v92 != 0 {
		v111 = v12
		goto L1
	} else {
		goto L29
	}
L26:
	;
	v92 = base.B2i32(base.Ui32(v79) < base.Ui32(v12))
	goto L25
L27:
	;
	goto L28
L28:
	;
	v92 = int32(base.Ui32(v79-v12) >> (uint(int32(31)) % 32))
	goto L25
L29:
	;
	v93 = v79
	goto L5
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v12
	F_errmsg_internal(m, int32(52934), v8)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L16
	} else {
		goto L31
	}
L31:
	;
	F_errfinish(m, int32(485414), int32(185), int32(253942))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L16
	} else {
		goto L32
	}
L32:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_SubTransPagePrecedes(m *base.Module, l0 int64, l1 int64) int32 {
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	v3 = int32(0)
	v7 = int32(11)
	v9 = int32(4)
	v10 = base.I32_wrap_i64(l0)<<(uint(v7)%32) | v9
	v13 = base.I32_wrap_i64(l1) << (uint(v7) % 32)
	v15 = v13 | v9
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v15))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v10)) == v3 {
		v27 = base.B2i32(base.Ui32(v10) < base.Ui32(v15))
	} else {
		v27 = int32(base.Ui32(v10-v15) >> (uint(int32(31)) % 32))
	}
	if v27 != 0 {
		v29 = v13 + int32(2051)
		if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v29))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v10)) == int32(0) {
			v41 = base.B2i32(base.Ui32(v10) < base.Ui32(v29))
		} else {
			v41 = int32(base.Ui32(v10-v29) >> (uint(int32(31)) % 32))
		}
		v42 = v41
	} else {
		v42 = v3
	}
	return v42
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
	v6 = *(*int32)(unsafe.Add(mBase, _consts[154]))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+28))
	v9 = int32(base.Ui32(l0) >> (uint(int32(11)) % 32))
	v11 = int32(*(*uint16)(unsafe.Add(mBase, _consts[155])))
	v12 = base.I32_rem_u_s(v9, v11)
	v15 = v7 + v12<<(uint(int32(7))%32)
	v17 = F_LWLockAcquire(m, v15, int32(0))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return
	} else {
		v22 = F_SimpleLruReadPage(m, int32(4365464), base.I64_extend_i32_u(v9), int32(1), l0)
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return
		} else {
			v25 = *(*int32)(unsafe.Add(mBase, _consts[154]))
			v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
			v27 = int32(2)
			v30 = *(*int32)(unsafe.Add(mBase, uint32(v26+v22<<(uint(v27)%32))))
			v35 = v30 + l0&int32(2047)<<(uint(v27)%32)
			v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
			if l1 != v36 {
				*(*int32)(unsafe.Add(mBase, uint32(v35))) = l1
				v40 = *(*int32)(unsafe.Add(mBase, _consts[154]))
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
