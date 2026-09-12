package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_AbortTransaction(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v198 int32
	_ = v198
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v242 int32
	_ = v242
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v274 int32
	_ = v274
	var v280 int32
	_ = v280
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v330 int64
	_ = v330
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v346 int32
	_ = v346
	var v350 int32
	_ = v350
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v355 int32
	_ = v355
	var v358 int32
	_ = v358
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v378 int32
	_ = v378
	var v380 int32
	_ = v380
	var v383 int32
	_ = v383
	var v385 int32
	_ = v385
	var v388 int32
	_ = v388
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v396 int32
	_ = v396
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v408 int32
	_ = v408
	var v420 int32
	_ = v420
	var v425 int32
	_ = v425
	var v427 int32
	_ = v427
	var v432 int32
	_ = v432
	var v435 int32
	_ = v435
	var v439 int32
	_ = v439
	var v442 int32
	_ = v442
	var v444 int32
	_ = v444
	var v451 int32
	_ = v451
	var v454 int32
	_ = v454
	var v456 int32
	_ = v456
	var v458 int32
	_ = v458
	var v461 int32
	_ = v461
	var v463 int32
	_ = v463
	var v476 int32
	_ = v476
	var v479 int32
	_ = v479
	var v482 int32
	_ = v482
	var v485 int32
	_ = v485
	var v489 int32
	_ = v489
	var v493 int32
	_ = v493
	var v496 int32
	_ = v496
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v502 int32
	_ = v502
	var v513 int32
	_ = v513
	var v522 int32
	_ = v522
	var v524 int32
	_ = v524
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = int32(4556732)
	v11 = *(*int32)(unsafe.Add(mBase, _consts[171]))
	*(*int32)(unsafe.Add(mBase, _consts[171])) = v11 + int32(1)
	v16 = *(*int32)(unsafe.Add(mBase, _consts[25]))
	v18 = *(*int32)(unsafe.Add(mBase, _consts[223]))
	if int32(0) < v18 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_disable_timeout(m, int32(8))
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
	v26 = *(*int32)(unsafe.Add(mBase, _consts[214]))
	*(*int32)(unsafe.Add(mBase, _consts[182])) = v26
	v30 = *(*int32)(unsafe.Add(mBase, _consts[211]))
	v32 = *(*int32)(unsafe.Add(mBase, _consts[36]))
	if v30 != 0 {
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
	v33 = v30
	goto L8
L7:
	;
	v33 = v32
	goto L8
L8:
	;
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v33
	F_LWLockReleaseAll(m)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L4
	} else {
		goto L9
	}
L9:
	;
	v38 = *(*int32)(unsafe.Add(mBase, _consts[157]))
	v39 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v38))) = v39
	v43 = *(*int32)(unsafe.Add(mBase, _consts[55]))
	if v43 == v39 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	F_pgaio_error_cleanup(m)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L4
	} else {
		goto L15
	}
L11:
	;
	goto L10
L12:
	;
	v47 = int32(*(*uint8)(unsafe.Add(mBase, _consts[56])))
	if v47 != int32(1) {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v43)+220))
	if v50 == int32(0) {
		goto L11
	} else {
		goto L14
	}
L14:
	;
	v53 = int32(4556740)
	v55 = *(*int32)(unsafe.Add(mBase, _consts[26]))
	v56 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[26])) = v55 + v56
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
	*(*int32)(unsafe.Add(mBase, uint32(v43))) = v59 + v56
	v63 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v43)+220)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v43)+224)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v43))) = v59 + int32(2)
	v73 = *(*int32)(unsafe.Add(mBase, _consts[26]))
	*(*int32)(unsafe.Add(mBase, _consts[26])) = v73 - v56
	goto L11
L15:
	;
	F_UnlockBuffers(m)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L4
	} else {
		goto L16
	}
L16:
	;
	v81 = int32(0)
	v87 = *(*int32)(unsafe.Add(mBase, _consts[130]))
	if v87 <= v81 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	F_ConditionVariableCancelSleep(m)
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L4
	} else {
		goto L30
	}
L18:
	;
	*(*int64)(unsafe.Add(mBase, _consts[230])) = int64(0)
	*(*int32)(unsafe.Add(mBase, _consts[231])) = int32(4457904)
	v198 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[130])) = v198
	*(*int32)(unsafe.Add(mBase, _consts[232])) = v198
	*(*uint8)(unsafe.Add(mBase, _consts[94])) = uint8(v198)
	*(*uint8)(unsafe.Add(mBase, _consts[233])) = uint8(v198)
	goto L17
L19:
	;
	v91 = v87 & int32(7)
	v93 = *(*int32)(unsafe.Add(mBase, _consts[132]))
	if base.Ui32(int32(8)) <= base.Ui32(v87) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v99 = v81
	v100 = int32(0)
	goto L23
L21:
	;
	v163 = v81
	goto L22
L22:
	;
	if v91 == int32(0) {
		goto L18
	} else {
		goto L26
	}
L23:
	;
	v104 = int32(8260)
	v107 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v93+v99*v104))) = uint8(v107)
	*(*uint8)(unsafe.Add(mBase, uint32(v93+(v99|int32(1))*v104))) = uint8(v107)
	*(*uint8)(unsafe.Add(mBase, uint32(v93+(v99|int32(2))*v104))) = uint8(v107)
	*(*uint8)(unsafe.Add(mBase, uint32(v93+(v99|int32(3))*v104))) = uint8(v107)
	*(*uint8)(unsafe.Add(mBase, uint32(v93+(v99|int32(4))*v104))) = uint8(v107)
	*(*uint8)(unsafe.Add(mBase, uint32(v93+(v99|int32(5))*v104))) = uint8(v107)
	*(*uint8)(unsafe.Add(mBase, uint32(v93+(v99|int32(6))*v104))) = uint8(v107)
	*(*uint8)(unsafe.Add(mBase, uint32(v93+(v99|int32(7))*v104))) = uint8(v107)
	v158 = int32(8)
	v159 = v99 + v158
	v161 = v100 + v158
	if v161 != v87&int32(2147483640) {
		v99 = v159
		v100 = v161
		goto L23
	} else {
		goto L25
	}
L24:
	;
	v163 = v159
	goto L22
L25:
	;
	goto L24
L26:
	;
	v171 = v163
	v172 = int32(0)
	goto L27
L27:
	;
	v179 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v93+v171*int32(8260)))) = uint8(v179)
	v181 = int32(1)
	v184 = v172 + v181
	if v184 != v91 {
		v171 = v171 + v181
		v172 = v184
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
	v212 = m.ExcPending
	if v212 != 0 {
		goto L4
	} else {
		goto L31
	}
L31:
	;
	F_reschedule_timeouts(m)
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L4
	} else {
		goto L32
	}
L32:
	;
	F_sigprocmask(m, int32(4469288), int32(0))
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L4
	} else {
		goto L33
	}
L33:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v16)+24))
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v16)+20))
	switch v220 - int32(2) {
	case 0, 3:
		goto L34
	default:
		goto L35
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+20)) = int32(4)
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v16)+60))
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v16)+64))
	*(*int32)(unsafe.Add(mBase, _consts[196])) = v252
	*(*int32)(unsafe.Add(mBase, _consts[168])) = v251
	goto L43
L35:
	;
	v225 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L4
	} else {
		goto L36
	}
L36:
	;
	if v225 == int32(0) {
		goto L34
	} else {
		goto L37
	}
L37:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v16)+20))
	if base.Ui32(v229) <= base.Ui32(int32(5)) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v229<<(uint(int32(2))%32))+uint32(_consts[229])))
	v238 = v236
	goto L40
L39:
	;
	v238 = int32(570185)
	goto L40
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v238
	F_errmsg_internal(m, int32(370000), v7)
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L4
	} else {
		goto L41
	}
L41:
	;
	F_errfinish(m, int32(517290), int32(2877), int32(270867))
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L4
	} else {
		goto L42
	}
L42:
	;
	goto L34
L43:
	;
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v16)+28))
	v259 = *(*int32)(unsafe.Add(mBase, _consts[234]))
	if v257 <= v259 {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	v274 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _consts[116])) = uint8(v274)
	*(*int32)(unsafe.Add(mBase, _consts[115])) = v274
	goto L48
L45:
	;
	v262 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[113])) = v262
	*(*int32)(unsafe.Add(mBase, _consts[235])) = v262
	*(*int32)(unsafe.Add(mBase, _consts[114])) = v262
	*(*int32)(unsafe.Add(mBase, _consts[234])) = v262
	goto L47
L46:
	;
	goto L47
L47:
	;
	goto L44
L48:
	;
	v280 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _consts[236])) = uint8(v280)
	*(*int32)(unsafe.Add(mBase, _consts[237])) = v280
	F_AtEOXact_Parallel(m, v280)
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L4
	} else {
		goto L49
	}
L49:
	;
	v289 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+76)) = uint8(v289)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+72)) = v289
	F_AfterTriggerEndXact(m)
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L4
	} else {
		goto L50
	}
L50:
	;
	F_AtAbort_Portals(m)
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L4
	} else {
		goto L51
	}
L51:
	;
	v299 = base.B2i32(v219 == int32(5))
	F_smgrDoPendingSyncs(m, int32(0), v299)
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L4
	} else {
		goto L52
	}
L52:
	;
	F_AtEOXact_LargeObject(m, int32(0))
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L4
	} else {
		goto L53
	}
L53:
	;
	v306 = int32(*(*uint8)(unsafe.Add(mBase, _consts[238])))
	if v306 == int32(0) {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v314 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[239])) = v314
	*(*int32)(unsafe.Add(mBase, _consts[240])) = v314
	F_AtEOXact_RelationMap(m, v314, v299)
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L4
	} else {
		goto L58
	}
L55:
	;
	v310 = *(*int32)(unsafe.Add(mBase, _consts[241]))
	if v310 != 0 {
		goto L54
	} else {
		goto L56
	}
L56:
	;
	F_asyncQueueUnregister(m)
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L4
	} else {
		goto L57
	}
L57:
	;
	goto L54
L58:
	;
	F_AtAbort_Twophase(m)
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L4
	} else {
		goto L59
	}
L59:
	;
	if v299 == int32(0) {
		goto L61
	} else {
		goto L62
	}
L60:
	;
	v335 = *(*int32)(unsafe.Add(mBase, _consts[120]))
	F_ProcArrayEndTransaction(m, v335, v333)
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L4
	} else {
		goto L66
	}
L61:
	;
	v327 = F_RecordTransactionAbort(m, int32(0))
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L4
	} else {
		goto L64
	}
L62:
	;
	goto L63
L63:
	;
	v330 = *(*int64)(unsafe.Add(mBase, _consts[175]))
	F_XLogSetAsyncXactLSN(m, v330)
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L4
	} else {
		goto L65
	}
L64:
	;
	v333 = v327
	goto L60
L65:
	;
	v333 = v280
	goto L60
L66:
	;
	v339 = *(*int32)(unsafe.Add(mBase, _consts[214]))
	if v339 != 0 {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v341 = *(*int32)(unsafe.Add(mBase, _consts[242]))
	if v219 == int32(5) {
		goto L71
	} else {
		goto L72
	}
L68:
	;
	goto L69
L69:
	;
	v522 = int32(4556732)
	v524 = *(*int32)(unsafe.Add(mBase, _consts[171]))
	*(*int32)(unsafe.Add(mBase, _consts[171])) = v524 - int32(1)
	m.G0 = v7 + int32(16)
	return
L70:
	;
	v373 = *(*int32)(unsafe.Add(mBase, _consts[214]))
	v374 = int32(1)
	F_ResourceOwnerRelease(m, v373, v374, int32(0), v374)
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L4
	} else {
		goto L84
	}
L71:
	;
	if v341 == int32(0) {
		goto L70
	} else {
		goto L74
	}
L72:
	;
	goto L73
L73:
	;
	if v341 == int32(0) {
		goto L70
	} else {
		goto L79
	}
L74:
	;
	v346 = v341
	goto L75
L75:
	;
	v350 = *(*int32)(unsafe.Add(mBase, uint32(v346)))
	v352 = *(*int32)(unsafe.Add(mBase, uint32(v346)+8))
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v346)+4))
	m.T0[v353].(func(*base.Module, int32, int32))(m, int32(3), v352)
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L4
	} else {
		goto L77
	}
L76:
	;
	goto L70
L77:
	;
	if v350 != 0 {
		v346 = v350
		goto L75
	} else {
		goto L78
	}
L78:
	;
	goto L76
L79:
	;
	v358 = v341
	goto L80
L80:
	;
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v358)))
	v364 = *(*int32)(unsafe.Add(mBase, uint32(v358)+8))
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v358)+4))
	m.T0[v365].(func(*base.Module, int32, int32))(m, int32(2), v364)
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L4
	} else {
		goto L82
	}
L81:
	;
	goto L70
L82:
	;
	if v362 != 0 {
		v358 = v362
		goto L80
	} else {
		goto L83
	}
L83:
	;
	goto L81
L84:
	;
	F_AtEOXact_Aio(m)
	mBase = m.M
	v380 = m.ExcPending
	if v380 != 0 {
		goto L4
	} else {
		goto L85
	}
L85:
	;
	F_AtEOXact_RelationCache(m, int32(0))
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L4
	} else {
		goto L86
	}
L86:
	;
	F_AtEOXact_TypeCache(m)
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L4
	} else {
		goto L87
	}
L87:
	;
	F_AtEOXact_Inval(m, int32(0))
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L4
	} else {
		goto L88
	}
L88:
	;
	v390 = *(*int32)(unsafe.Add(mBase, _consts[152]))
	v391 = int32(4164656)
	v392 = *(*int32)(unsafe.Add(mBase, _consts[166]))
	v393 = int32(2)
	v396 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v390+v392<<(uint(v393)%32)))) = v396
	v399 = *(*int32)(unsafe.Add(mBase, _consts[243]))
	v401 = *(*int32)(unsafe.Add(mBase, _consts[166]))
	*(*int32)(unsafe.Add(mBase, uint32(v399+v401<<(uint(v393)%32)))) = v396
	v408 = int32(4155296)
	*(*int32)(unsafe.Add(mBase, _consts[244])) = v408
	*(*int32)(unsafe.Add(mBase, _consts[245])) = v408
	*(*int32)(unsafe.Add(mBase, _consts[246])) = v396
	*(*int32)(unsafe.Add(mBase, _consts[247])) = v396
	goto L89
L89:
	;
	v420 = *(*int32)(unsafe.Add(mBase, _consts[214]))
	F_ResourceOwnerRelease(m, v420, int32(2), int32(0), int32(1))
	mBase = m.M
	v425 = m.ExcPending
	if v425 != 0 {
		goto L4
	} else {
		goto L90
	}
L90:
	;
	v427 = *(*int32)(unsafe.Add(mBase, _consts[214]))
	F_ResourceOwnerRelease(m, v427, int32(3), int32(0), int32(1))
	mBase = m.M
	v432 = m.ExcPending
	if v432 != 0 {
		goto L4
	} else {
		goto L91
	}
L91:
	;
	F_smgrDoPendingDeletes(m, int32(0))
	mBase = m.M
	v435 = m.ExcPending
	if v435 != 0 {
		goto L4
	} else {
		goto L92
	}
L92:
	;
	F_AtEOXact_GUC(m, int32(0), int32(1))
	mBase = m.M
	v439 = m.ExcPending
	if v439 != 0 {
		goto L4
	} else {
		goto L93
	}
L93:
	;
	F_AtEOXact_SPI(m, int32(0))
	mBase = m.M
	v442 = m.ExcPending
	if v442 != 0 {
		goto L4
	} else {
		goto L94
	}
L94:
	;
	v444 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[248])) = v444
	*(*int32)(unsafe.Add(mBase, _consts[249])) = v444
	goto L95
L95:
	;
	F_AtEOXact_on_commit_actions(m, int32(0))
	mBase = m.M
	v451 = m.ExcPending
	if v451 != 0 {
		goto L4
	} else {
		goto L96
	}
L96:
	;
	v454 = base.B2i32(v219 == int32(5))
	F_AtEOXact_Namespace(m, int32(0), v454)
	mBase = m.M
	v456 = m.ExcPending
	if v456 != 0 {
		goto L4
	} else {
		goto L97
	}
L97:
	;
	F_smgrdestroyall(m)
	mBase = m.M
	v458 = m.ExcPending
	if v458 != 0 {
		goto L4
	} else {
		goto L98
	}
L98:
	;
	F_AtEOXact_Files(m, int32(0))
	mBase = m.M
	v461 = m.ExcPending
	if v461 != 0 {
		goto L4
	} else {
		goto L99
	}
L99:
	;
	v463 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[93])) = v463
	*(*int32)(unsafe.Add(mBase, _consts[250])) = v463
	*(*int32)(unsafe.Add(mBase, _consts[251])) = v463
	*(*int32)(unsafe.Add(mBase, _consts[252])) = v463
	goto L100
L100:
	;
	F_AtEOXact_HashTables(m, int32(0))
	mBase = m.M
	v476 = m.ExcPending
	if v476 != 0 {
		goto L4
	} else {
		goto L101
	}
L101:
	;
	F_AtEOXact_PgStat(m, int32(0), v454)
	mBase = m.M
	v479 = m.ExcPending
	if v479 != 0 {
		goto L4
	} else {
		goto L102
	}
L102:
	;
	F_AtEOXact_ApplyLauncher(m, int32(0))
	mBase = m.M
	v482 = m.ExcPending
	if v482 != 0 {
		goto L4
	} else {
		goto L103
	}
L103:
	;
	F_AtEOXact_LogicalRepWorkers(m, int32(0))
	mBase = m.M
	v485 = m.ExcPending
	if v485 != 0 {
		goto L4
	} else {
		goto L104
	}
L104:
	;
	v489 = int32(*(*uint8)(unsafe.Add(mBase, _consts[56])))
	if v489 != int32(1) {
		goto L106
	} else {
		goto L107
	}
L105:
	;
	goto L69
L106:
	;
	goto L105
L107:
	;
	v493 = *(*int32)(unsafe.Add(mBase, _consts[55]))
	if v493 == int32(0) {
		goto L106
	} else {
		goto L108
	}
L108:
	;
	v496 = int32(4556740)
	v498 = *(*int32)(unsafe.Add(mBase, _consts[26]))
	v499 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[26])) = v498 + v499
	v502 = *(*int32)(unsafe.Add(mBase, uint32(v493)))
	*(*int32)(unsafe.Add(mBase, uint32(v493))) = v502 + v499
	*(*int64)(unsafe.Add(mBase, uint32(v493)+24)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v493))) = v502 + int32(2)
	v513 = *(*int32)(unsafe.Add(mBase, _consts[26]))
	*(*int32)(unsafe.Add(mBase, _consts[26])) = v513 - v499
	goto L106
}
func F_AssignTransactionId(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
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
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
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
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v123 int32
	_ = v123
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v137 int64
	_ = v137
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v153 int32
	_ = v153
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v195 int64
	_ = v195
	var v199 int32
	_ = v199
	var v207 int64
	_ = v207
	var v212 int32
	_ = v212
	var v214 int64
	_ = v214
	var v216 int32
	_ = v216
	var v217 int64
	_ = v217
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v234 int32
	_ = v234
	var v237 int64
	_ = v237
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v260 int32
	_ = v260
	var v265 int32
	_ = v265
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v14 = *(*int32)(unsafe.Add(mBase, _consts[25]))
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
	v253 = m.ExcPending
	if v253 != 0 {
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
	v18 = *(*int32)(unsafe.Add(mBase, _consts[181]))
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
	v158 = int32(4562132)
	v159 = *(*int32)(unsafe.Add(mBase, _consts[182]))
	v161 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	*(*int32)(unsafe.Add(mBase, _consts[182])) = v161
	v163 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v164 = m.G0
	v166 = v164 - int32(16)
	m.G0 = v166
	*(*int32)(unsafe.Add(mBase, uint32(v166)+12)) = int32(17104896)
	*(*int64)(unsafe.Add(mBase, uint32(v166)+4)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v166))) = v163
	v174 = int32(0)
	v176 = F_LockAcquire(m, v166, int32(7), v174, v174)
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
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
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	if v69 == int32(0) {
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
	*(*int64)(unsafe.Add(mBase, _consts[84])) = v25
	v30 = base.I32_wrap_i64(v25)
	v31 = m.G0
	v33 = v31 - int32(16)
	m.G0 = v33
	v36 = *(*int32)(unsafe.Add(mBase, _consts[85]))
	if v36 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v38 = *(*int32)(unsafe.Add(mBase, _consts[86]))
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
	m.G0 = v33 + int32(16)
	v153 = int32(0)
	goto L5
L14:
	;
	v45 = *(*int32)(unsafe.Add(mBase, _consts[85]))
	*(*int32)(unsafe.Add(mBase, uint32(v45)+96)) = v30
	*(*int32)(unsafe.Add(mBase, uint32(v33)+12)) = v30
	v49 = *(*int32)(unsafe.Add(mBase, _consts[87]))
	v55 = F_hash_search(m, v49, v33+int32(12), int32(1), v33+int32(11))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L9
	} else {
		goto L15
	}
L15:
	;
	v58 = *(*int32)(unsafe.Add(mBase, _consts[85]))
	*(*int32)(unsafe.Add(mBase, uint32(v55)+4)) = v58
	v61 = *(*int32)(unsafe.Add(mBase, _consts[86]))
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
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v76 = F_palloc(m, v73<<(uint(int32(2))%32))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L9
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v133 = int32(*(*uint8)(unsafe.Add(mBase, _consts[183])))
	v135 = *(*int32)(unsafe.Add(mBase, _consts[27]))
	v137 = F_GetNewTransactionId(m, int32(1))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L9
	} else {
		goto L35
	}
L20:
	;
	v79 = int32(0)
	v81 = v21
	goto L21
L21:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v81)))
	if v86 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	if v96 != 0 {
		goto L27
	} else {
		goto L28
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v76+v79<<(uint(int32(2))%32)))) = v81
	v94 = v79 + int32(1)
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v81)+80))
	if v95 != 0 {
		v79 = v94
		v81 = v95
		goto L21
	} else {
		goto L26
	}
L24:
	;
	v96 = v79
	goto L25
L25:
	;
	goto L22
L26:
	;
	v96 = v94
	goto L25
L27:
	;
	v99 = v96
	goto L30
L28:
	;
	goto L29
L29:
	;
	F_pfree(m, v76)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L9
	} else {
		goto L34
	}
L30:
	;
	v107 = v99 - int32(1)
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v76+v107<<(uint(int32(2))%32))))
	F_AssignTransactionId(m, v111)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L9
	} else {
		goto L32
	}
L31:
	;
	goto L29
L32:
	;
	if v107 != 0 {
		v99 = v107
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
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v137
	v141 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v141)))
	F_SubTransSetParent(m, base.I32_wrap_i64(v137), v142)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L9
	} else {
		goto L36
	}
L36:
	;
	v145 = int32(1)
	v153 = (v133 ^ v145) & base.B2i32(v145 < v135)
	goto L5
L37:
	;
	m.G0 = v166 + int32(16)
	*(*int32)(unsafe.Add(mBase, _consts[182])) = v159
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
	v186 = *(*int32)(unsafe.Add(mBase, _consts[27]))
	if v186 <= int32(0) {
		goto L38
	} else {
		goto L40
	}
L40:
	;
	v189 = int32(4457552)
	v190 = *(*int32)(unsafe.Add(mBase, _consts[184]))
	v195 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*uint32)(unsafe.Add(mBase, uint32(v190<<(uint(int32(2))%32))+uint32(_consts[185]))) = uint32(v195)
	v199 = v190 + int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[184])) = v199
	if (v153^int32(-1))&base.B2i32(v199 < int32(64)) != 0 {
		goto L38
	} else {
		goto L41
	}
L41:
	;
	v207 = int64(*(*uint32)(unsafe.Add(mBase, _consts[84])))
	if v207 == int64(0) {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	F_AssignTransactionId(m, int32(4457208))
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L9
	} else {
		goto L45
	}
L43:
	;
	v217 = v207
	v218 = v199
	goto L44
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v218
	*(*uint32)(unsafe.Add(mBase, uint32(v11)+8)) = uint32(v217)
	F_XLogBeginInsert(m)
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L9
	} else {
		goto L46
	}
L45:
	;
	v214 = *(*int64)(unsafe.Add(mBase, _consts[84]))
	v216 = *(*int32)(unsafe.Add(mBase, _consts[184]))
	v217 = v214
	v218 = v216
	goto L44
L46:
	;
	v223 = int32(8)
	F_XLogRegisterData(m, v11+v223, v223)
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L9
	} else {
		goto L47
	}
L47:
	;
	v230 = *(*int32)(unsafe.Add(mBase, _consts[184]))
	F_XLogRegisterData(m, int32(4457296), v230<<(uint(int32(2))%32))
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L9
	} else {
		goto L48
	}
L48:
	;
	v237 = F_XLogInsert(m, int32(1), int32(80))
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L9
	} else {
		goto L49
	}
L49:
	;
	v240 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[183])) = uint8(v240)
	*(*int32)(unsafe.Add(mBase, _consts[184])) = int32(0)
	goto L38
L50:
	;
	F_errcode(m, int32(322))
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L9
	} else {
		goto L51
	}
L51:
	;
	F_errmsg(m, int32(273550), int32(0))
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L9
	} else {
		goto L52
	}
L52:
	;
	F_errfinish(m, int32(517290), int32(652), int32(487310))
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
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
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v119 int64
	_ = v119
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v10 = *(*int32)(unsafe.Add(mBase, _consts[25]))
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
	v21 = *(*int32)(unsafe.Add(mBase, _consts[228]))
	F_hash_seq_init(m, v16+int32(12), v21)
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
	v141 = m.ExcPending
	if v141 != 0 {
		goto L4
	} else {
		goto L44
	}
L4:
	;
	return
L5:
	;
	v26 = F_hash_seq_search(m, v16+int32(12))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	if v26 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v28 = v26
	goto L10
L8:
	;
	goto L9
L9:
	;
	m.G0 = v16 + int32(32)
	F_AtEOXact_Snapshot(m, int32(0), int32(1))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L4
	} else {
		goto L31
	}
L10:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v28)+64))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+80))
	if v33 == int32(3) {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	goto L9
L12:
	;
	v67 = F_hash_seq_search(m, v16+int32(12))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L4
	} else {
		goto L29
	}
L13:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v32)+20))
	if v36 == int32(0) {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+85)))
	if v39 != 0 {
		goto L12
	} else {
		goto L15
	}
L15:
	;
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+84)))
	if v40 == int32(1) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v43 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+84)) = uint8(v43)
	goto L18
L17:
	;
	goto L18
L18:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v32)+16))
	if v45 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v48 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L4
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	F_PortalDrop(m, v32, int32(0))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L4
	} else {
		goto L28
	}
L22:
	;
	if v48 != 0 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v50
	F_errmsg_internal(m, int32(746707), v16)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L4
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+16)) = int32(0)
	goto L21
L26:
	;
	F_errfinish(m, int32(522107), int32(901), int32(163197))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
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
	if v67 != 0 {
		v28 = v67
		goto L10
	} else {
		goto L30
	}
L30:
	;
	goto L11
L31:
	;
	*(*int32)(unsafe.Add(mBase, _consts[182])) = int32(0)
	v84 = *(*int32)(unsafe.Add(mBase, _consts[214]))
	if v84 != 0 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	F_ResourceOwnerDelete(m, v84)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L4
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	v87 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+40)) = v87
	*(*int32)(unsafe.Add(mBase, _consts[214])) = v87
	*(*int32)(unsafe.Add(mBase, _consts[213])) = v87
	v97 = *(*int32)(unsafe.Add(mBase, _consts[25]))
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v97)+44))
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v98
	v101 = *(*int32)(unsafe.Add(mBase, _consts[211]))
	if v101 != 0 {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	goto L34
L36:
	;
	F_MemoryContextReset(m, v101)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L4
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	v105 = *(*int32)(unsafe.Add(mBase, _consts[153]))
	if v105 != 0 {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	goto L38
L40:
	;
	F_MemoryContextReset(m, v105)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L4
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	v109 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[212])) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v97)+36)) = v109
	*(*uint8)(unsafe.Add(mBase, uint32(v10)+76)) = uint8(v109)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+72)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v10)+56)) = v109
	v119 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v10)+48)) = v119
	*(*int64)(unsafe.Add(mBase, uint32(v10)+28)) = v119
	*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v109
	*(*int64)(unsafe.Add(mBase, uint32(v10))) = v119
	*(*int64)(unsafe.Add(mBase, _consts[84])) = v119
	*(*int32)(unsafe.Add(mBase, _consts[91])) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v109
	m.G0 = v7 + int32(16)
	return
L43:
	;
	goto L42
L44:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v10)+20))
	if base.Ui32(v142) <= base.Ui32(int32(5)) {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v142<<(uint(int32(2))%32))+uint32(_consts[229])))
	v151 = v149
	goto L47
L46:
	;
	v151 = int32(570185)
	goto L47
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v151
	F_errmsg_internal(m, int32(197813), v7)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L4
	} else {
		goto L48
	}
L48:
	;
	F_errfinish(m, int32(517290), int32(3018), int32(270944))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
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
	v16 = *(*int32)(unsafe.Add(mBase, _consts[100]))
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
			v12 = *(*int32)(unsafe.Add(mBase, _consts[25]))
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
				v23 = *(*int32)(unsafe.Add(mBase, _consts[153]))
				v24 = F_MemoryContextStrdup(m, v23, l0)
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, _consts[265])) = v24
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
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
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
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v10 = *(*int32)(unsafe.Add(mBase, _consts[25]))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+24))
	if base.Ui32(int32(1)) < base.Ui32(v11) {
		m.G0 = v7 + int32(16)
		return
	} else {
		if l0 == int32(0) {
			m.G0 = v7 + int32(16)
			return
		} else {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v10)+28))
			if int32(1) < v16 {
				m.G0 = v7 + int32(16)
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
						*(*int32)(unsafe.Add(mBase, uint32(v7))) = l1
						F_errmsg(m, int32(163659), v7)
						mBase = m.M
						v29 = m.ExcPending
						if v29 != 0 {
							return
						} else {
							F_errfinish(m, int32(517290), int32(3749), int32(332723))
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
	*(*int32)(unsafe.Add(mBase, _consts[134])) = v3
	v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)))
	*(*uint8)(unsafe.Add(mBase, _consts[207])) = uint8(v6)
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+5)))
	*(*uint8)(unsafe.Add(mBase, _consts[204])) = uint8(v9)
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
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v210 int64
	_ = v210
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v224 int32
	_ = v224
	var v229 int64
	_ = v229
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v240 int64
	_ = v240
	var v241 int64
	_ = v241
	var v249 int64
	_ = v249
	var v253 int64
	_ = v253
	var v256 int32
	_ = v256
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v280 int32
	_ = v280
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v303 int32
	_ = v303
	var v307 int32
	_ = v307
	var v312 int32
	_ = v312
	var v320 int32
	_ = v320
	var v331 int32
	_ = v331
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v344 int32
	_ = v344
	var v356 int32
	_ = v356
	var v360 int32
	_ = v360
	var v364 int32
	_ = v364
	var v368 int32
	_ = v368
	var v371 int32
	_ = v371
	var v375 int32
	_ = v375
	var v377 int32
	_ = v377
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v11 = int32(4457208)
	*(*int32)(unsafe.Add(mBase, _consts[25])) = v11
	*(*int32)(unsafe.Add(mBase, _consts[186])) = int32(1)
	*(*int64)(unsafe.Add(mBase, _consts[187])) = int64(0)
	v21 = *(*float64)(unsafe.Add(mBase, _consts[188]))
	if base.F64_eq(v21, float64(0)) != 0 {
		v57 = int32(0)
		goto L1
	} else {
		goto L2
	}
L1:
	;
	*(*int64)(unsafe.Add(mBase, _consts[191])) = int64(0)
	*(*int64)(unsafe.Add(mBase, _consts[192])) = int64(4294967297)
	*(*uint8)(unsafe.Add(mBase, _consts[193])) = uint8(v57)
	*(*int32)(unsafe.Add(mBase, _consts[194])) = int32(0)
	v72 = *(*int32)(unsafe.Add(mBase, _consts[168]))
	*(*int32)(unsafe.Add(mBase, _consts[195])) = v72
	v75 = *(*int32)(unsafe.Add(mBase, _consts[196]))
	*(*int32)(unsafe.Add(mBase, _consts[197])) = v75
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
	v27 = int32(4645584)
	v30 = *(*int64)(unsafe.Add(mBase, _consts[189]))
	v31 = *(*int64)(unsafe.Add(mBase, _consts[190]))
	v32 = v30 ^ v31
	*(*int64)(unsafe.Add(mBase, _consts[190])) = base.I64_rotl(v32, int64(37))
	*(*int64)(unsafe.Add(mBase, _consts[189])) = v32<<(uint(int64(16))%64) ^ base.I64_rotl(v30, int64(24)) ^ v32
	v53 = F_ldexp(m, base.F64_convert_i64_u(int64(base.Ui64(base.I64_rotl(v30*int64(5), int64(7))*int64(9))>>(uint(int64(12))%64))), int32(-52))
	mBase = m.M
	goto L4
L4:
	;
	v55 = *(*float64)(unsafe.Add(mBase, _consts[188]))
	v57 = base.F64_le(v53, v55)
	goto L1
L5:
	;
	v80 = int32(*(*uint8)(unsafe.Add(mBase, _consts[198])))
	if v80 == int32(1) {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	*(*uint8)(unsafe.Add(mBase, _consts[200])) = uint8(v90)
	v93 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[201])) = v93
	*(*int32)(unsafe.Add(mBase, _consts[202])) = v93
	v100 = int32(*(*uint8)(unsafe.Add(mBase, _consts[203])))
	*(*uint8)(unsafe.Add(mBase, _consts[204])) = uint8(v100)
	v104 = *(*int32)(unsafe.Add(mBase, _consts[205]))
	*(*int32)(unsafe.Add(mBase, _consts[134])) = v104
	v109 = int32(*(*uint8)(unsafe.Add(mBase, _consts[206])))
	if v90 != 0 {
		goto L10
	} else {
		goto L11
	}
L7:
	;
	v85 = *(*int32)(unsafe.Add(mBase, _consts[199]))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)+316))
	v88 = base.B2i32(v86 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _consts[198])) = uint8(v88)
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
	*(*uint8)(unsafe.Add(mBase, _consts[207])) = uint8(v110)
	v113 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _consts[208])) = uint8(v113)
	*(*int32)(unsafe.Add(mBase, _consts[173])) = v113
	*(*int32)(unsafe.Add(mBase, _consts[209])) = v113
	*(*uint8)(unsafe.Add(mBase, _consts[210])) = uint8(v113)
	*(*int32)(unsafe.Add(mBase, _consts[184])) = v113
	*(*uint8)(unsafe.Add(mBase, _consts[183])) = uint8(v113)
	v131 = *(*int32)(unsafe.Add(mBase, _consts[25]))
	v133 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	*(*int32)(unsafe.Add(mBase, uint32(v131)+44)) = v133
	v136 = *(*int32)(unsafe.Add(mBase, _consts[211]))
	if v136 == v113 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v141 = *(*int32)(unsafe.Add(mBase, _consts[36]))
	v143 = int32(32768)
	v146 = F_AllocSetContextCreateInternal(m, v141, int32(67879), v143, v143, v143)
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
	v150 = *(*int32)(unsafe.Add(mBase, _consts[153]))
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
	*(*int32)(unsafe.Add(mBase, _consts[211])) = v146
	goto L15
L18:
	;
	v155 = *(*int32)(unsafe.Add(mBase, _consts[36]))
	v160 = F_AllocSetContextCreateInternal(m, v155, int32(68067), int32(0), int32(8192), int32(8388608))
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
	*(*int32)(unsafe.Add(mBase, _consts[212])) = v163
	*(*int32)(unsafe.Add(mBase, uint32(v131)+36)) = v163
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v163
	v170 = *(*int32)(unsafe.Add(mBase, _consts[25]))
	v173 = F_ResourceOwnerCreate(m, int32(0), int32(270963))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L16
	} else {
		goto L22
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, _consts[153])) = v160
	v163 = v160
	goto L20
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v170)+40)) = v173
	*(*int32)(unsafe.Add(mBase, _consts[213])) = v173
	*(*int32)(unsafe.Add(mBase, _consts[214])) = v173
	*(*int32)(unsafe.Add(mBase, _consts[182])) = v173
	v183 = *(*int32)(unsafe.Add(mBase, _consts[166]))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v183
	v186 = int32(4479444)
	v187 = int32(1)
	v189 = *(*int32)(unsafe.Add(mBase, _consts[215]))
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
	*(*int32)(unsafe.Add(mBase, _consts[215])) = v192 + int32(1)
	goto L23
L27:
	;
	v202 = *(*int32)(unsafe.Add(mBase, _consts[120]))
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v202)+56)) = v203
	v206 = *(*int32)(unsafe.Add(mBase, _consts[181]))
	if int32(0) <= v206 {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	v256 = int32(*(*uint8)(unsafe.Add(mBase, _consts[56])))
	if v256 != int32(1) {
		goto L41
	} else {
		goto L42
	}
L29:
	;
	v210 = *(*int64)(unsafe.Add(mBase, _consts[216]))
	v253 = v210
	goto L28
L30:
	;
	goto L31
L31:
	;
	v211 = int32(0)
	v213 = *(*int32)(unsafe.Add(mBase, _consts[217]))
	if v213 == v211 {
		v224 = v211
		goto L32
	} else {
		goto L33
	}
L32:
	;
	if v224 == int32(0) {
		goto L36
	} else {
		goto L37
	}
L33:
	;
	v216 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v213)+40)))
	if v216 != 0 {
		v224 = v211
		goto L32
	} else {
		goto L34
	}
L34:
	;
	v218 = *(*int32)(unsafe.Add(mBase, _consts[25]))
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v218)+28))
	goto L35
L35:
	;
	v224 = base.B2i32(int32(1) < v219) ^ int32(1)
	goto L32
L36:
	;
	v229 = *(*int64)(unsafe.Add(mBase, _consts[218]))
	*(*int64)(unsafe.Add(mBase, _consts[216])) = v229
	v253 = v229
	goto L28
L37:
	;
	goto L38
L38:
	;
	v235 = m.G0
	v236 = int32(16)
	v237 = v235 - v236
	m.G0 = v237
	F___gettimeofday(m, v237)
	mBase = m.M
	v240 = *(*int64)(unsafe.Add(mBase, uint32(v237)))
	v241 = int64(*(*int32)(unsafe.Add(mBase, uint32(v237)+8)))
	m.G0 = v237 + v236
	v249 = v241 + v240*int64(1000000) - int64(946684800000000)
	goto L39
L39:
	;
	*(*int64)(unsafe.Add(mBase, _consts[216])) = v249
	v253 = v249
	goto L28
L40:
	;
	*(*int64)(unsafe.Add(mBase, _consts[219])) = int64(0)
	v288 = m.G0
	v290 = v288 - int32(16)
	m.G0 = v290
	v293 = *(*int32)(unsafe.Add(mBase, _consts[220]))
	if v293 == int32(0) {
		goto L44
	} else {
		goto L45
	}
L41:
	;
	goto L40
L42:
	;
	v260 = *(*int32)(unsafe.Add(mBase, _consts[55]))
	if v260 == int32(0) {
		goto L41
	} else {
		goto L43
	}
L43:
	;
	v263 = int32(4556740)
	v265 = *(*int32)(unsafe.Add(mBase, _consts[26]))
	v266 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[26])) = v265 + v266
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v260)))
	*(*int32)(unsafe.Add(mBase, uint32(v260))) = v269 + v266
	*(*int64)(unsafe.Add(mBase, uint32(v260)+24)) = v253
	*(*int32)(unsafe.Add(mBase, uint32(v260))) = v269 + int32(2)
	v280 = *(*int32)(unsafe.Add(mBase, _consts[26]))
	*(*int32)(unsafe.Add(mBase, _consts[26])) = v280 - v266
	goto L41
L44:
	;
	*(*int32)(unsafe.Add(mBase, _consts[220])) = int32(1)
	m.G0 = v290 + int32(16)
	F_ReceiveSharedInvalidMessages(m)
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L16
	} else {
		goto L50
	}
L45:
	;
	v298 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L16
	} else {
		goto L46
	}
L46:
	;
	if v298 == int32(0) {
		goto L44
	} else {
		goto L47
	}
L47:
	;
	v303 = *(*int32)(unsafe.Add(mBase, _consts[220]))
	*(*int32)(unsafe.Add(mBase, uint32(v290))) = v303
	F_errmsg_internal(m, int32(89143), v290)
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L16
	} else {
		goto L48
	}
L48:
	;
	F_errfinish(m, int32(525050), int32(2224), int32(570953))
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L16
	} else {
		goto L49
	}
L49:
	;
	goto L44
L50:
	;
	*(*int32)(unsafe.Add(mBase, _consts[221])) = int32(-1)
	*(*int32)(unsafe.Add(mBase, _consts[222])) = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[186])) = int32(2)
	v331 = *(*int32)(unsafe.Add(mBase, _consts[223]))
	if int32(0) < v331 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	F_enable_timeout_after(m, int32(8), v331)
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L16
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	v337 = int32(10)
	goto L57
L54:
	;
	goto L53
L55:
	;
	if v371 != 0 {
		goto L69
	} else {
		goto L70
	}
L56:
	;
	goto L55
L57:
	;
	v344 = *(*int32)(unsafe.Add(mBase, _consts[224]))
	goto L60
L58:
	;
	v356 = int32(0)
	goto L66
L60:
	;
	goto L61
L61:
	;
	goto L63
L63:
	;
	if v344 == int32(15) {
		goto L58
	} else {
		goto L64
	}
L64:
	;
	if v344 <= v337 {
		v371 = int32(1)
		goto L56
	} else {
		goto L65
	}
L65:
	;
	goto L58
L66:
	;
	v360 = *(*int32)(unsafe.Add(mBase, _consts[225]))
	if v360 != int32(2) {
		v371 = v356
		goto L56
	} else {
		goto L67
	}
L67:
	;
	v364 = int32(*(*uint8)(unsafe.Add(mBase, _consts[226])))
	if v364 != 0 {
		v371 = v356
		goto L56
	} else {
		goto L68
	}
L68:
	;
	v368 = *(*int32)(unsafe.Add(mBase, _consts[227]))
	v371 = int32(0) | base.B2i32(v368 <= v337)
	goto L56
L69:
	;
	v375 = *(*int32)(unsafe.Add(mBase, _consts[25]))
	F_ShowTransactionStateRec(m, int32(270884), v375)
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L16
	} else {
		goto L72
	}
L70:
	;
	goto L71
L71:
	;
	m.G0 = v8 + int32(16)
	return
L72:
	;
	goto L71
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
	v6 = int32(*(*uint8)(unsafe.Add(mBase, _consts[490])))
	if v6 != 0 {
		v34 = v4
		return v34
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, _consts[25]))
		v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+28))
		if int32(1) < v9 {
			v18 = int32(268623)
			*(*int32)(unsafe.Add(mBase, _consts[492])) = int32(16777538)
			v24 = *(*int32)(unsafe.Add(mBase, _consts[140]))
			*(*int32)(unsafe.Add(mBase, _consts[141])) = v24
			v29 = F_format_elog_string(m, v18, int32(0))
			mBase = m.M
			v32 = m.ExcPending
			if v32 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, _consts[493])) = v29
				v34 = int32(0)
				return v34
			}
		} else {
			v14 = int32(*(*uint8)(unsafe.Add(mBase, _consts[491])))
			if v14 != int32(1) {
				v34 = v4
				return v34
			} else {
				v18 = int32(16416)
				*(*int32)(unsafe.Add(mBase, _consts[492])) = int32(16777538)
				v24 = *(*int32)(unsafe.Add(mBase, _consts[140]))
				*(*int32)(unsafe.Add(mBase, _consts[141])) = v24
				v29 = F_format_elog_string(m, v18, int32(0))
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, _consts[493])) = v29
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
	var v19 int32
	_ = v19
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	v4 = int32(1)
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v7 = *(*int32)(unsafe.Add(mBase, _consts[134]))
	if v5 == v7 {
		v83 = v4
		return v83
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, _consts[25]))
		v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+20))
		if base.B2i32(v11 == int32(2)) == int32(0) {
			v83 = v4
			return v83
		} else {
			v17 = int32(*(*uint8)(unsafe.Add(mBase, _consts[490])))
			if v17 != 0 {
				v83 = v4
				return v83
			} else {
				v19 = int32(*(*uint8)(unsafe.Add(mBase, _consts[491])))
				if v19 == int32(1) {
					*(*int32)(unsafe.Add(mBase, _consts[492])) = int32(16777538)
					v71 = int32(16352)
					v72 = int32(4560664)
					v75 = *(*int32)(unsafe.Add(mBase, _consts[140]))
					*(*int32)(unsafe.Add(mBase, _consts[141])) = v75
					v79 = F_format_elog_string(m, v71, int32(0))
					mBase = m.M
					v80 = m.ExcPending
					if v80 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v72))) = v79
						v83 = int32(0)
						return v83
					}
				} else {
					v28 = *(*int32)(unsafe.Add(mBase, _consts[25]))
					v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+28))
					if int32(1) < v29 {
						*(*int32)(unsafe.Add(mBase, _consts[492])) = int32(16777538)
						v71 = int32(268739)
						v72 = int32(4560664)
						v75 = *(*int32)(unsafe.Add(mBase, _consts[140]))
						*(*int32)(unsafe.Add(mBase, _consts[141])) = v75
						v79 = F_format_elog_string(m, v71, int32(0))
						mBase = m.M
						v80 = m.ExcPending
						if v80 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v72))) = v79
							v83 = int32(0)
							return v83
						}
					} else {
						if v5 != int32(3) {
							v83 = v4
							return v83
						} else {
							v41 = int32(*(*uint8)(unsafe.Add(mBase, _consts[198])))
							if v41 == int32(1) {
								v46 = *(*int32)(unsafe.Add(mBase, _consts[199]))
								v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)+316))
								v49 = base.B2i32(v47 != int32(2))
								*(*uint8)(unsafe.Add(mBase, _consts[198])) = uint8(v49)
								v51 = v49
							} else {
								v51 = int32(0)
							}
							if v51 == int32(0) {
								v83 = v4
								return v83
							} else {
								*(*int32)(unsafe.Add(mBase, _consts[492])) = int32(1088)
								v58 = *(*int32)(unsafe.Add(mBase, _consts[140]))
								*(*int32)(unsafe.Add(mBase, _consts[141])) = v58
								v64 = F_format_elog_string(m, int32(23959), int32(0))
								mBase = m.M
								v67 = m.ExcPending
								if v67 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, _consts[493])) = v64
									v71 = int32(681565)
									v72 = int32(4560672)
									v75 = *(*int32)(unsafe.Add(mBase, _consts[140]))
									*(*int32)(unsafe.Add(mBase, _consts[141])) = v75
									v79 = F_format_elog_string(m, v71, int32(0))
									mBase = m.M
									v80 = m.ExcPending
									if v80 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v72))) = v79
										v83 = int32(0)
										return v83
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
