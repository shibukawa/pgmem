package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ExecAllocTableSlot(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	v5 = F_MakeTupleTableSlot(m, l1, l2, int32(0))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v10 = F_lappend(m, v9, v5)
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0))) = v10
			return v5
		}
	}
}
func F_ExecMakeTableFunctionResult(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v93 int32
	_ = v93
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v113 int64
	_ = v113
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v150 int32
	_ = v150
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v166 int64
	_ = v166
	var v168 int32
	_ = v168
	var v185 int32
	_ = v185
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int64
	_ = v234
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v254 int64
	_ = v254
	var v256 int64
	_ = v256
	var v257 int64
	_ = v257
	var v258 int64
	_ = v258
	var v262 int64
	_ = v262
	var v263 int64
	_ = v263
	var v266 int64
	_ = v266
	var v268 int64
	_ = v268
	var v273 int64
	_ = v273
	var v285 int32
	_ = v285
	var v286 int64
	_ = v286
	var v287 int32
	_ = v287
	var v293 int32
	_ = v293
	var v298 int32
	_ = v298
	var v306 int32
	_ = v306
	var v309 int32
	_ = v309
	var v313 int32
	_ = v313
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v353 int32
	_ = v353
	var v357 int32
	_ = v357
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v373 int32
	_ = v373
	var v376 int32
	_ = v376
	var v378 int32
	_ = v378
	var v381 int32
	_ = v381
	var v384 int32
	_ = v384
	var v387 int32
	_ = v387
	var v390 int32
	_ = v390
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v398 int32
	_ = v398
	var v401 int32
	_ = v401
	var v407 int32
	_ = v407
	var v413 int32
	_ = v413
	var v415 int32
	_ = v415
	var v421 int32
	_ = v421
	var v428 int32
	_ = v428
	var v432 int32
	_ = v432
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
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
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v464 int32
	_ = v464
	var v466 int32
	_ = v466
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v482 int32
	_ = v482
	var v486 int32
	_ = v486
	var v488 int32
	_ = v488
	var v490 int32
	_ = v490
	var v498 int32
	_ = v498
	var v501 int32
	_ = v501
	var v505 int32
	_ = v505
	var v510 int32
	_ = v510
	var v514 int32
	_ = v514
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v522 int32
	_ = v522
	var v527 int32
	_ = v527
	var v532 int32
	_ = v532
	var v535 int32
	_ = v535
	var v539 int32
	_ = v539
	var v544 int32
	_ = v544
	var v556 int32
	_ = v556
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v565 int32
	_ = v565
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v579 int32
	_ = v579
	var v584 int32
	_ = v584
	var v588 int32
	_ = v588
	var v592 int32
	_ = v592
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v598 int32
	_ = v598
	var v602 int32
	_ = v602
	v6 = int32(0)
	v17 = m.G0
	v19 = v17 - int32(112)
	m.G0 = v19
	F_MemoryContextReset(m, l2)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v25 = int32(_a_F_ExecMakeTableFunctionResult_0)
	v26 = *(*int32)(unsafe.Add(mBase, _c_F_ExecMakeTableFunctionResult[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecMakeTableFunctionResult[0])) = l2
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v30 = F_exprType(m, v29)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v32 = F_type_is_rowtype(m, v30)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v19)+72)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+64)) = int32(1)
	if l4 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v40 = int32(15)
	goto L7
L6:
	;
	v40 = int32(11)
	goto L7
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+60)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v19)+56)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v19)+52)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v19)+48)) = int32(389)
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v46 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	v53 = v47<<(uint(int32(4))%32) + int32(24)
	goto L10
L9:
	;
	v53 = int32(24)
	goto L10
L10:
	;
	v54 = F_palloc(m, v53)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v56 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	v561 = *(*int32)(unsafe.Add(mBase, uint32(v19)+72))
	if v561 != 0 {
		goto L131
	} else {
		goto L132
	}
L13:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecMakeTableFunctionResult[0])) = v191
	v194 = v54 + int32(16)
	v196 = int32(0)
	v204 = int32(1)
	v205 = v196
	v207 = v196
	goto L33
L14:
	;
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+57)))
	v60 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v54)+4)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v54))) = l0 + int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v54)+8)) = v19 + int32(48)
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)+12))
	*(*uint8)(unsafe.Add(mBase, uint32(v54)+16)) = uint8(v60)
	*(*int32)(unsafe.Add(mBase, uint32(v54)+12)) = v69
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v73 != 0 {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	goto L16
L16:
	;
	v166 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v54))) = v166
	v168 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v54)+18)) = uint16(v168)
	*(*int64)(unsafe.Add(mBase, uint32(v54)+8)) = v166
	*(*uint8)(unsafe.Add(mBase, uint32(v54)+16)) = uint8(v168)
	v185 = v6
	goto L13
L17:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v73)+4))
	v76 = v74
	goto L19
L18:
	;
	v76 = int32(0)
	goto L19
L19:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v54)+18)) = uint16(v76)
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v78 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+26)))
	if v136 != int32(1) {
		v185 = v59
		goto L13
	} else {
		goto L27
	}
L21:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v78)+4))
	if v81 <= int32(0) {
		goto L20
	} else {
		goto L22
	}
L22:
	;
	v93 = v6
	goto L23
L23:
	;
	v104 = v54 + int32(24) + v93<<(uint(int32(4))%32)
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v78)+12))
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v105+v93<<(uint(int32(2))%32))))
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v109)+24))
	v113 = m.T0[v112].(func(*base.Module, int32, int32, int32) int64)(m, v109, l1, v104+int32(8))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L1
	} else {
		goto L25
	}
L24:
	;
	goto L20
L25:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v104))) = v113
	v117 = v93 + int32(1)
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v78)+4))
	if v117 < v118 {
		v93 = v117
		goto L23
	} else {
		goto L26
	}
L26:
	;
	goto L24
L27:
	;
	v139 = int32(0)
	v140 = int32(*(*int16)(unsafe.Add(mBase, uint32(v54)+18)))
	if v140 <= v139 {
		v185 = v59
		goto L13
	} else {
		goto L28
	}
L28:
	;
	v150 = v139
	goto L29
L29:
	;
	v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54+v150<<(uint(int32(4))%32))+32)))
	if v162 != 0 {
		v556 = v59
		goto L12
	} else {
		goto L31
	}
L30:
	;
	v185 = v59
	goto L13
L31:
	;
	v164 = v150 + int32(1)
	if v140 != v164 {
		v150 = v164
		goto L29
	} else {
		goto L32
	}
L32:
	;
	goto L30
L33:
	;
	v215 = *(*int32)(unsafe.Add(mBase, _c_F_ExecMakeTableFunctionResult[1]))
	if v215 != 0 {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v532 = m.ExcPending
	if v532 != 0 {
		goto L1
	} else {
		goto L127
	}
L35:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L1
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	F_MemoryContextReset(m, v218)
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L1
	} else {
		goto L39
	}
L38:
	;
	goto L37
L39:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v221 == int32(0) {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v19)+64))
	if v293 != int32(1) {
		goto L57
	} else {
		goto L58
	}
L41:
	;
	v225 = v19 + int32(80)
	F_pgstat_init_function_usage(m, v54, v225)
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L1
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v221)+24))
	v286 = m.T0[v285].(func(*base.Module, int32, int32, int32) int64)(m, v221, l1, v194)
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L1
	} else {
		goto L53
	}
L44:
	;
	v228 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v54)+16)) = uint8(v228)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+68)) = v228
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v232)))
	v234 = m.T0[v233].(func(*base.Module, int32) int64)(m, v54)
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v19)+16)) = v234
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v19)+68))
	v246 = m.G0
	v248 = v246 - int32(16)
	m.G0 = v248
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v225)))
	if v250 != 0 {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	goto L40
L47:
	;
	F___clock_gettime(m, int32(1), v248)
	mBase = m.M
	v253 = int32(_a_F_ExecMakeTableFunctionResult_1)
	v254 = *(*int64)(unsafe.Add(mBase, _c_F_ExecMakeTableFunctionResult[2]))
	v256 = *(*int64)(unsafe.Add(mBase, uint32(v225)+16))
	v257 = int64(*(*int32)(unsafe.Add(mBase, uint32(v248)+8)))
	v258 = *(*int64)(unsafe.Add(mBase, uint32(v248)))
	v262 = *(*int64)(unsafe.Add(mBase, uint32(v225)+24))
	v263 = v257 + v258*int64(1000000000) - v262
	*(*int64)(unsafe.Add(mBase, _c_F_ExecMakeTableFunctionResult[2])) = v256 + v263
	v266 = *(*int64)(unsafe.Add(mBase, uint32(v225)+8))
	if v237 != int32(1) {
		goto L50
	} else {
		goto L51
	}
L48:
	;
	goto L49
L49:
	;
	m.G0 = v248 + int32(16)
	goto L46
L50:
	;
	v268 = *(*int64)(unsafe.Add(mBase, uint32(v250)))
	*(*int64)(unsafe.Add(mBase, uint32(v250))) = v268 + int64(1)
	goto L52
L51:
	;
	goto L52
L52:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v250)+8)) = v266 + v263
	v273 = *(*int64)(unsafe.Add(mBase, uint32(v250)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v250)+16)) = v273 + (v263 - v254 + v256)
	goto L49
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+68)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v19)+16)) = v286
	goto L40
L54:
	;
	goto L34
L55:
	;
	v204 = int32(0)
	v205 = v488
	v207 = v437
	goto L33
L56:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v514 = m.ExcPending
	if v514 != 0 {
		goto L1
	} else {
		goto L123
	}
L57:
	;
	if v293 != int32(2) {
		goto L56
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v19)+68))
	if v319 == int32(2) {
		v556 = v185
		goto L12
	} else {
		goto L66
	}
L60:
	;
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v19)+68))
	if v204&base.B2i32(v298 == int32(0))&v185 != 0 {
		v556 = v185
		goto L12
	} else {
		goto L61
	}
L61:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	F_errcode(m, int32(33686083))
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	F_errmsg(m, int32(_a_F_ExecMakeTableFunctionResult_2), int32(0))
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	F_errfinish(m, int32(_a_F_ExecMakeTableFunctionResult_3), int32(375), int32(_a_F_ExecMakeTableFunctionResult_4))
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L66:
	;
	if v204 != 0 {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v322 = int32(_a_F_ExecMakeTableFunctionResult_0)
	v323 = *(*int32)(unsafe.Add(mBase, _c_F_ExecMakeTableFunctionResult[0]))
	v325 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecMakeTableFunctionResult[0])) = v325
	v329 = *(*int32)(unsafe.Add(mBase, _c_F_ExecMakeTableFunctionResult[3]))
	v330 = F_tuplestore_begin_heap(m, l4, int32(0), v329)
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L1
	} else {
		goto L70
	}
L68:
	;
	v436 = v205
	v437 = v207
	goto L69
L69:
	;
	if v32 != 0 {
		goto L96
	} else {
		goto L97
	}
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+72)) = v330
	if v32 == int32(0) {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v336 = F_CreateTemplateTupleDesc(m, int32(1))
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L1
	} else {
		goto L74
	}
L72:
	;
	v432 = v205
	goto L73
L73:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecMakeTableFunctionResult[0])) = v323
	v436 = v432
	v437 = v330
	goto L69
L74:
	;
	F_TupleDescInitEntry(m, v336, int32(1), int32(_a_F_ExecMakeTableFunctionResult_5), v30, int32(-1), int32(0))
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L1
	} else {
		goto L75
	}
L75:
	;
	v344 = int32(0)
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v336)))
	if v344 < v353 {
		goto L77
	} else {
		goto L78
	}
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+76)) = v336
	v432 = v336
	goto L73
L77:
	;
	v357 = v336 + int32(28)
	v364 = v344
	v365 = v353
	v367 = v344
	goto L81
L78:
	;
	v421 = v344
	v428 = v353
	goto L79
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v336)+20)) = v428
	*(*int32)(unsafe.Add(mBase, uint32(v336)+16)) = v421
	goto L76
L80:
	;
	v421 = v415
	v428 = v394
	goto L79
L81:
	;
	v373 = v357 + v353<<(uint(int32(3))%32) + v364*int32(100)
	v376 = v357 + v364<<(uint(int32(3))%32)
	if v353 != v365 {
		v394 = v365
		goto L83
	} else {
		goto L84
	}
L82:
	;
	v415 = v353
	goto L80
L83:
	;
	v395 = int32(*(*int16)(unsafe.Add(mBase, uint32(v376)+2)))
	if v395 <= int32(0) {
		v415 = v364
		goto L80
	} else {
		goto L91
	}
L84:
	;
	v378 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v376)+7)))
	if v378 != int32(118) {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v394 = v364
	goto L83
L86:
	;
	v381 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v376)+4)))
	if v381 != int32(1) {
		goto L85
	} else {
		goto L87
	}
L87:
	;
	v384 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v376)+6)))
	if v384&int32(6) != 0 {
		goto L85
	} else {
		goto L88
	}
L88:
	;
	v387 = int32(*(*int16)(unsafe.Add(mBase, uint32(v376)+2)))
	if v387 <= int32(0) {
		goto L85
	} else {
		goto L89
	}
L89:
	;
	v390 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v373)+90)))
	if v390 != int32(118) {
		v394 = v353
		goto L83
	} else {
		goto L90
	}
L90:
	;
	goto L85
L91:
	;
	v398 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v373)+90)))
	if v398 == int32(118) {
		v415 = v364
		goto L80
	} else {
		goto L92
	}
L92:
	;
	v401 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v376)+5)))
	v407 = (v367 + v401 - int32(1)) & (int32(0) - v401)
	if int32(_a_F_ExecMakeTableFunctionResult_6) < v407 {
		v415 = v364
		goto L80
	} else {
		goto L93
	}
L93:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v376))) = uint16(v407)
	v413 = v364 + int32(1)
	if v413 != v353 {
		v364 = v413
		v365 = v394
		v367 = v407 + v395
		goto L81
	} else {
		goto L94
	}
L94:
	;
	goto L82
L95:
	;
	v490 = *(*int32)(unsafe.Add(mBase, uint32(v19)+68))
	if v490 != int32(1) {
		v556 = v185
		goto L12
	} else {
		goto L117
	}
L96:
	;
	v438 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v194))))
	if v438 == int32(0) {
		goto L99
	} else {
		goto L100
	}
L97:
	;
	goto L98
L98:
	;
	F_tuplestore_putvalues(m, v437, v436, v19+int32(16), v194)
	mBase = m.M
	v486 = m.ExcPending
	if v486 != 0 {
		goto L1
	} else {
		goto L116
	}
L99:
	;
	v441 = *(*int32)(unsafe.Add(mBase, uint32(v19)+16))
	v442 = F_pg_detoast_datum(m, v441)
	mBase = m.M
	v443 = m.ExcPending
	if v443 != 0 {
		goto L1
	} else {
		goto L102
	}
L100:
	;
	goto L101
L101:
	;
	v475 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v476 = F_palloc(m, v475)
	mBase = m.M
	v477 = m.ExcPending
	if v477 != 0 {
		goto L1
	} else {
		goto L111
	}
L102:
	;
	if v436 == int32(0) {
		goto L104
	} else {
		goto L105
	}
L103:
	;
	v466 = *(*int32)(unsafe.Add(mBase, uint32(v442)))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+44)) = v442
	*(*int32)(unsafe.Add(mBase, uint32(v19)+28)) = int32(base.Ui32(v466) >> (uint(int32(2)) % 32))
	F_tuplestore_puttuple(m, v437, v19+int32(28))
	mBase = m.M
	v474 = m.ExcPending
	if v474 != 0 {
		goto L1
	} else {
		goto L110
	}
L104:
	;
	v446 = int32(_a_F_ExecMakeTableFunctionResult_0)
	v447 = *(*int32)(unsafe.Add(mBase, _c_F_ExecMakeTableFunctionResult[0]))
	v449 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecMakeTableFunctionResult[0])) = v449
	v451 = *(*int32)(unsafe.Add(mBase, uint32(v442)+8))
	v452 = *(*int32)(unsafe.Add(mBase, uint32(v442)+4))
	v453 = F_lookup_rowtype_tupdesc_copy(m, v451, v452)
	mBase = m.M
	v454 = m.ExcPending
	if v454 != 0 {
		goto L1
	} else {
		goto L107
	}
L105:
	;
	goto L106
L106:
	;
	v458 = *(*int32)(unsafe.Add(mBase, uint32(v442)+8))
	v459 = *(*int32)(unsafe.Add(mBase, uint32(v436)+4))
	if v458 != v459 {
		goto L54
	} else {
		goto L108
	}
L107:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecMakeTableFunctionResult[0])) = v447
	*(*int32)(unsafe.Add(mBase, uint32(v19)+76)) = v453
	v464 = v453
	goto L103
L108:
	;
	v461 = *(*int32)(unsafe.Add(mBase, uint32(v442)+4))
	v462 = *(*int32)(unsafe.Add(mBase, uint32(v436)+8))
	if v461 != v462 {
		goto L54
	} else {
		goto L109
	}
L109:
	;
	v464 = v436
	goto L103
L110:
	;
	v488 = v464
	goto L95
L111:
	;
	if v475 != 0 {
		goto L112
	} else {
		goto L113
	}
L112:
	;
	base.MemoryFill(m, v476, int32(1), v475)
	goto L114
L113:
	;
	goto L114
L114:
	;
	F_tuplestore_putvalues(m, v437, l3, int32(0), v476)
	mBase = m.M
	v482 = m.ExcPending
	if v482 != 0 {
		goto L1
	} else {
		goto L115
	}
L115:
	;
	v488 = v436
	goto L95
L116:
	;
	v488 = v436
	goto L95
L117:
	;
	if v185&int32(1) != 0 {
		goto L55
	} else {
		goto L118
	}
L118:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v498 = m.ExcPending
	if v498 != 0 {
		goto L1
	} else {
		goto L119
	}
L119:
	;
	F_errcode(m, int32(33686083))
	mBase = m.M
	v501 = m.ExcPending
	if v501 != 0 {
		goto L1
	} else {
		goto L120
	}
L120:
	;
	F_errmsg(m, int32(_a_F_ExecMakeTableFunctionResult_7), int32(0))
	mBase = m.M
	v505 = m.ExcPending
	if v505 != 0 {
		goto L1
	} else {
		goto L121
	}
L121:
	;
	F_errfinish(m, int32(_a_F_ExecMakeTableFunctionResult_3), int32(367), int32(_a_F_ExecMakeTableFunctionResult_4))
	mBase = m.M
	v510 = m.ExcPending
	if v510 != 0 {
		goto L1
	} else {
		goto L122
	}
L122:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L123:
	;
	F_errcode(m, int32(33686083))
	mBase = m.M
	v517 = m.ExcPending
	if v517 != 0 {
		goto L1
	} else {
		goto L124
	}
L124:
	;
	v518 = *(*int32)(unsafe.Add(mBase, uint32(v19)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v518
	F_errmsg(m, int32(_a_F_ExecMakeTableFunctionResult_8), v19)
	mBase = m.M
	v522 = m.ExcPending
	if v522 != 0 {
		goto L1
	} else {
		goto L125
	}
L125:
	;
	F_errfinish(m, int32(_a_F_ExecMakeTableFunctionResult_3), int32(383), int32(_a_F_ExecMakeTableFunctionResult_4))
	mBase = m.M
	v527 = m.ExcPending
	if v527 != 0 {
		goto L1
	} else {
		goto L126
	}
L126:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L127:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v535 = m.ExcPending
	if v535 != 0 {
		goto L1
	} else {
		goto L128
	}
L128:
	;
	F_errmsg(m, int32(_a_F_ExecMakeTableFunctionResult_9), int32(0))
	mBase = m.M
	v539 = m.ExcPending
	if v539 != 0 {
		goto L1
	} else {
		goto L129
	}
L129:
	;
	F_errfinish(m, int32(_a_F_ExecMakeTableFunctionResult_3), int32(317), int32(_a_F_ExecMakeTableFunctionResult_4))
	mBase = m.M
	v544 = m.ExcPending
	if v544 != 0 {
		goto L1
	} else {
		goto L130
	}
L130:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L131:
	;
	v588 = *(*int32)(unsafe.Add(mBase, uint32(v19)+76))
	if v588 == int32(0) {
		goto L140
	} else {
		goto L141
	}
L132:
	;
	v562 = int32(_a_F_ExecMakeTableFunctionResult_0)
	v563 = *(*int32)(unsafe.Add(mBase, _c_F_ExecMakeTableFunctionResult[0]))
	v565 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecMakeTableFunctionResult[0])) = v565
	v569 = *(*int32)(unsafe.Add(mBase, _c_F_ExecMakeTableFunctionResult[3]))
	v570 = F_tuplestore_begin_heap(m, l4, int32(0), v569)
	mBase = m.M
	v571 = m.ExcPending
	if v571 != 0 {
		goto L1
	} else {
		goto L133
	}
L133:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecMakeTableFunctionResult[0])) = v563
	*(*int32)(unsafe.Add(mBase, uint32(v19)+72)) = v570
	if v556&int32(1) != 0 {
		goto L131
	} else {
		goto L134
	}
L134:
	;
	v577 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v578 = F_palloc(m, v577)
	mBase = m.M
	v579 = m.ExcPending
	if v579 != 0 {
		goto L1
	} else {
		goto L135
	}
L135:
	;
	if v577 != 0 {
		goto L136
	} else {
		goto L137
	}
L136:
	;
	base.MemoryFill(m, v578, int32(1), v577)
	goto L138
L137:
	;
	goto L138
L138:
	;
	F_tuplestore_putvalues(m, v570, l3, int32(0), v578)
	mBase = m.M
	v584 = m.ExcPending
	if v584 != 0 {
		goto L1
	} else {
		goto L139
	}
L139:
	;
	goto L131
L140:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecMakeTableFunctionResult[0])) = v26
	v602 = *(*int32)(unsafe.Add(mBase, uint32(v19)+72))
	m.G0 = v19 + int32(112)
	return v602
L141:
	;
	F_tupledesc_match(m, l3, v588)
	mBase = m.M
	v592 = m.ExcPending
	if v592 != 0 {
		goto L1
	} else {
		goto L142
	}
L142:
	;
	v593 = *(*int32)(unsafe.Add(mBase, uint32(v19)+76))
	v594 = *(*int32)(unsafe.Add(mBase, uint32(v593)+12))
	if v594 != int32(-1) {
		goto L140
	} else {
		goto L143
	}
L143:
	;
	F_FreeTupleDesc(m, v593)
	mBase = m.M
	v598 = m.ExcPending
	if v598 != 0 {
		goto L1
	} else {
		goto L144
	}
L144:
	;
	goto L140
}
func F_ExecTableFuncScan(m *base.Module, l0 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = F_ExecScan(m, l0, int32(808), int32(809))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_LockTableRecurse(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v62 int64
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	v4 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v14 = F_find_all_inheritors(m, l0, v4, v4)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v10 + int32(16)
	return
L2:
	;
	return
L3:
	;
	if v14 == int32(0) {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	if v18 <= int32(0) {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v27 = v4
	goto L6
L6:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v28+v27<<(uint(int32(2))%32))))
	if v32 == l0 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	goto L1
L8:
	;
	v71 = v27 + int32(1)
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	if v71 < v72 {
		v27 = v71
		goto L6
	} else {
		goto L26
	}
L9:
	;
	if l2 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v62 = int64(0)
	v65 = F_SearchSysCacheExists(m, int32(57), base.I64_extend_i32_u(v32), v62, v62, v62)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L2
	} else {
		goto L23
	}
L11:
	;
	F_LockRelationOid(m, v32, l1)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L2
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	v38 = F_ConditionalLockRelationOid(m, v32, l1)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L2
	} else {
		goto L15
	}
L14:
	;
	goto L10
L15:
	;
	if v38 != 0 {
		goto L10
	} else {
		goto L16
	}
L16:
	;
	v40 = F_get_rel_name(m, v32)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L2
	} else {
		goto L17
	}
L17:
	;
	if v40 == int32(0) {
		goto L8
	} else {
		goto L18
	}
L18:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L2
	} else {
		goto L19
	}
L19:
	;
	F_errcode(m, int32(50463045))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L2
	} else {
		goto L20
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v40
	F_errmsg(m, int32(_a_F_LockTableRecurse_0), v10)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L2
	} else {
		goto L21
	}
L21:
	;
	F_errfinish(m, int32(_a_F_LockTableRecurse_1), int32(144), int32(_a_F_LockTableRecurse_2))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L2
	} else {
		goto L22
	}
L22:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L23:
	;
	if v65 != 0 {
		goto L8
	} else {
		goto L24
	}
L24:
	;
	F_UnlockRelationOid(m, v32, l1)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L2
	} else {
		goto L25
	}
L25:
	;
	goto L8
L26:
	;
	goto L7
}
func F_TableToProcessComparator(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 float64
	_ = v7
	var v8 int32
	_ = v8
	var v9 float64
	_ = v9
	var v12 int32
	_ = v12
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v7 = *(*float64)(unsafe.Add(mBase, uint32(v6)+8))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v9 = *(*float64)(unsafe.Add(mBase, uint32(v8)+8))
	if base.F64_lt(v7, v9) != 0 {
		v12 = int32(-1)
	} else {
		v12 = base.F64_gt(v7, v9)
	}
	return v12
}
func F__equalDropTableSpaceStmt(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	v3 = int32(0)
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v7 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return v42
L2:
	;
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+8)))
	v42 = base.B2i32(v39 == v40)
	goto L1
L3:
	;
	if v6 == int32(0) {
		v42 = v3
		goto L1
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	if v6 != v7 {
		v42 = v3
		goto L1
	} else {
		goto L15
	}
L6:
	;
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6))))
	if base.B2i32(v12 == int32(0))|base.B2i32(v12 != v15) != 0 {
		v33 = v12
		v34 = v15
		goto L8
	} else {
		goto L9
	}
L7:
	;
	if v33-v34 == int32(0) {
		goto L2
	} else {
		goto L14
	}
L8:
	;
	goto L7
L9:
	;
	v18 = v7
	v19 = v6
	goto L10
L10:
	;
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+1)))
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+1)))
	if v23 == int32(0) {
		v33 = v23
		v34 = v22
		goto L8
	} else {
		goto L12
	}
L11:
	;
	v33 = v23
	v34 = v22
	goto L8
L12:
	;
	v26 = int32(1)
	if v23 == v22 {
		v18 = v18 + v26
		v19 = v19 + v26
		goto L10
	} else {
		goto L13
	}
L13:
	;
	goto L11
L14:
	;
	v42 = v3
	goto L1
L15:
	;
	goto L2
}
func F__equalTableSampleClause(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	v3 = int32(0)
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v4 != v5 {
		v19 = v3
		return v19
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
		v9 = F_equal(m, v7, v8)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			if v9 == int32(0) {
				v19 = v3
				return v19
			} else {
				v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
				v17 = F_equal(m, v15, v16)
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return int32(0)
				} else {
					v19 = v17
					return v19
				}
			}
		}
	}
}
func F_table_beginscan_parallel_tidrange(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	v4 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+12)) = uint8(v4)
	v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+13)))
	if v8 != 0 {
		v20 = int32(_a_F_table_beginscan_parallel_tidrange_0)
		v21 = int32(272)
		v23 = *(*int32)(unsafe.Add(mBase, _c_F_table_beginscan_parallel_tidrange[0]))
		if v23 == int32(0) {
			v43 = int32(0)
			v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
			v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)+8))
			v48 = m.T0[v47].(func(*base.Module, int32, int32, int32, int32, int32, int32) int32)(m, l0, v20, v43, v43, l1, l2|v21)
			mBase = m.M
			v49 = m.ExcPending
			if v49 != 0 {
				return int32(0)
			} else {
				return v48
			}
		} else {
			v27 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_table_beginscan_parallel_tidrange[1])))
			if v27&int32(1) != 0 {
				v43 = int32(0)
				v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
				v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)+8))
				v48 = m.T0[v47].(func(*base.Module, int32, int32, int32, int32, int32, int32) int32)(m, l0, v20, v43, v43, l1, l2|v21)
				mBase = m.M
				v49 = m.ExcPending
				if v49 != 0 {
					return int32(0)
				} else {
					return v48
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return int32(0)
				} else {
					F_errmsg_internal(m, int32(_a_F_table_beginscan_parallel_tidrange_1), int32(0))
					mBase = m.M
					v37 = m.ExcPending
					if v37 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_table_beginscan_parallel_tidrange_2), int32(931), int32(_a_F_table_beginscan_parallel_tidrange_3))
						mBase = m.M
						v42 = m.ExcPending
						if v42 != 0 {
							return int32(0)
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
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
		v13 = F_RestoreSnapshot(m, l1+v11)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			v17 = F_RegisterSnapshot(m, v13)
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				v20 = v13
				v21 = int32(784)
				v23 = *(*int32)(unsafe.Add(mBase, _c_F_table_beginscan_parallel_tidrange[0]))
				if v23 == int32(0) {
					v43 = int32(0)
					v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
					v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)+8))
					v48 = m.T0[v47].(func(*base.Module, int32, int32, int32, int32, int32, int32) int32)(m, l0, v20, v43, v43, l1, l2|v21)
					mBase = m.M
					v49 = m.ExcPending
					if v49 != 0 {
						return int32(0)
					} else {
						return v48
					}
				} else {
					v27 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_table_beginscan_parallel_tidrange[1])))
					if v27&int32(1) != 0 {
						v43 = int32(0)
						v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
						v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)+8))
						v48 = m.T0[v47].(func(*base.Module, int32, int32, int32, int32, int32, int32) int32)(m, l0, v20, v43, v43, l1, l2|v21)
						mBase = m.M
						v49 = m.ExcPending
						if v49 != 0 {
							return int32(0)
						} else {
							return v48
						}
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v33 = m.ExcPending
						if v33 != 0 {
							return int32(0)
						} else {
							F_errmsg_internal(m, int32(_a_F_table_beginscan_parallel_tidrange_1), int32(0))
							mBase = m.M
							v37 = m.ExcPending
							if v37 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_table_beginscan_parallel_tidrange_2), int32(931), int32(_a_F_table_beginscan_parallel_tidrange_3))
								mBase = m.M
								v42 = m.ExcPending
								if v42 != 0 {
									return int32(0)
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
}
func F_table_block_parallelscan_estimate(m *base.Module, l0 int32) int32 {
	return int32(48)
}
func F_table_index_fetch_tuple_check(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
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
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	v5 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	*(*uint8)(unsafe.Add(mBase, uint32(v9)+15)) = uint8(v5)
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	if v13 != 0 {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
		v15 = m.T0[v14].(func(*base.Module, int32) int32)(m, l0)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			v26 = v15
			v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
			v28 = F_MakeSingleTupleTableSlot(m, v27, v26)
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return int32(0)
			} else {
				v31 = *(*int32)(unsafe.Add(mBase, _c_F_table_index_fetch_tuple_check[0]))
				if v31 == int32(0) {
					v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
					v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)+44))
					v54 = m.T0[v53].(func(*base.Module, int32, int32) int32)(m, l0, int32(0))
					mBase = m.M
					v55 = m.ExcPending
					if v55 != 0 {
						return int32(0)
					} else {
						v58 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
						v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)+188))
						v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)+56))
						v61 = m.T0[v60].(func(*base.Module, int32, int32, int32, int32, int32, int32) int32)(m, v54, l1, l2, v28, v9+int32(15), l3)
						mBase = m.M
						v62 = m.ExcPending
						if v62 != 0 {
							return int32(0)
						} else {
							v63 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
							v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)+188))
							v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+52))
							m.T0[v65].(func(*base.Module, int32))(m, v54)
							mBase = m.M
							v67 = m.ExcPending
							if v67 != 0 {
								return int32(0)
							} else {
								F_ExecDropSingleTupleTableSlot(m, v28)
								mBase = m.M
								v69 = m.ExcPending
								if v69 != 0 {
									return int32(0)
								} else {
									m.G0 = v9 + int32(16)
									return v61
								}
							}
						}
					}
				} else {
					v35 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_table_index_fetch_tuple_check[1])))
					if v35&int32(1) != 0 {
						v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
						v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)+44))
						v54 = m.T0[v53].(func(*base.Module, int32, int32) int32)(m, l0, int32(0))
						mBase = m.M
						v55 = m.ExcPending
						if v55 != 0 {
							return int32(0)
						} else {
							v58 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
							v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)+188))
							v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)+56))
							v61 = m.T0[v60].(func(*base.Module, int32, int32, int32, int32, int32, int32) int32)(m, v54, l1, l2, v28, v9+int32(15), l3)
							mBase = m.M
							v62 = m.ExcPending
							if v62 != 0 {
								return int32(0)
							} else {
								v63 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
								v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)+188))
								v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+52))
								m.T0[v65].(func(*base.Module, int32))(m, v54)
								mBase = m.M
								v67 = m.ExcPending
								if v67 != 0 {
									return int32(0)
								} else {
									F_ExecDropSingleTupleTableSlot(m, v28)
									mBase = m.M
									v69 = m.ExcPending
									if v69 != 0 {
										return int32(0)
									} else {
										m.G0 = v9 + int32(16)
										return v61
									}
								}
							}
						}
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v41 = m.ExcPending
						if v41 != 0 {
							return int32(0)
						} else {
							F_errmsg_internal(m, int32(_a_F_table_index_fetch_tuple_check_0), int32(0))
							mBase = m.M
							v45 = m.ExcPending
							if v45 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_table_index_fetch_tuple_check_1), int32(1256), int32(_a_F_table_index_fetch_tuple_check_2))
								mBase = m.M
								v50 = m.ExcPending
								if v50 != 0 {
									return int32(0)
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
	} else {
		v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
		v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+119)))
		if v22 == int32(102) {
			v25 = int32(_a_F_table_index_fetch_tuple_check_3)
		} else {
			v25 = int32(_a_F_table_index_fetch_tuple_check_4)
		}
		v26 = v25
		v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
		v28 = F_MakeSingleTupleTableSlot(m, v27, v26)
		mBase = m.M
		v29 = m.ExcPending
		if v29 != 0 {
			return int32(0)
		} else {
			v31 = *(*int32)(unsafe.Add(mBase, _c_F_table_index_fetch_tuple_check[0]))
			if v31 == int32(0) {
				v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
				v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)+44))
				v54 = m.T0[v53].(func(*base.Module, int32, int32) int32)(m, l0, int32(0))
				mBase = m.M
				v55 = m.ExcPending
				if v55 != 0 {
					return int32(0)
				} else {
					v58 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
					v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)+188))
					v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)+56))
					v61 = m.T0[v60].(func(*base.Module, int32, int32, int32, int32, int32, int32) int32)(m, v54, l1, l2, v28, v9+int32(15), l3)
					mBase = m.M
					v62 = m.ExcPending
					if v62 != 0 {
						return int32(0)
					} else {
						v63 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
						v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)+188))
						v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+52))
						m.T0[v65].(func(*base.Module, int32))(m, v54)
						mBase = m.M
						v67 = m.ExcPending
						if v67 != 0 {
							return int32(0)
						} else {
							F_ExecDropSingleTupleTableSlot(m, v28)
							mBase = m.M
							v69 = m.ExcPending
							if v69 != 0 {
								return int32(0)
							} else {
								m.G0 = v9 + int32(16)
								return v61
							}
						}
					}
				}
			} else {
				v35 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_table_index_fetch_tuple_check[1])))
				if v35&int32(1) != 0 {
					v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
					v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)+44))
					v54 = m.T0[v53].(func(*base.Module, int32, int32) int32)(m, l0, int32(0))
					mBase = m.M
					v55 = m.ExcPending
					if v55 != 0 {
						return int32(0)
					} else {
						v58 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
						v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)+188))
						v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)+56))
						v61 = m.T0[v60].(func(*base.Module, int32, int32, int32, int32, int32, int32) int32)(m, v54, l1, l2, v28, v9+int32(15), l3)
						mBase = m.M
						v62 = m.ExcPending
						if v62 != 0 {
							return int32(0)
						} else {
							v63 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
							v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)+188))
							v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+52))
							m.T0[v65].(func(*base.Module, int32))(m, v54)
							mBase = m.M
							v67 = m.ExcPending
							if v67 != 0 {
								return int32(0)
							} else {
								F_ExecDropSingleTupleTableSlot(m, v28)
								mBase = m.M
								v69 = m.ExcPending
								if v69 != 0 {
									return int32(0)
								} else {
									m.G0 = v9 + int32(16)
									return v61
								}
							}
						}
					}
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v41 = m.ExcPending
					if v41 != 0 {
						return int32(0)
					} else {
						F_errmsg_internal(m, int32(_a_F_table_index_fetch_tuple_check_0), int32(0))
						mBase = m.M
						v45 = m.ExcPending
						if v45 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_table_index_fetch_tuple_check_1), int32(1256), int32(_a_F_table_index_fetch_tuple_check_2))
							mBase = m.M
							v50 = m.ExcPending
							if v50 != 0 {
								return int32(0)
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
func F_table_parallelscan_estimate(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v3 != 0 {
		v12 = int32(0)
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+32))
		v15 = m.T0[v14].(func(*base.Module, int32) int32)(m, l0)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			v17 = F_add_size(m, v12, v15)
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				return v17
			}
		}
	} else {
		v6 = F_EstimateSnapshotSpace(m, l1)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			v10 = F_add_size(m, int32(0), v6)
			mBase = m.M
			v11 = m.ExcPending
			if v11 != 0 {
				return int32(0)
			} else {
				v12 = v10
				v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
				v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+32))
				v15 = m.T0[v14].(func(*base.Module, int32) int32)(m, l0)
				mBase = m.M
				v16 = m.ExcPending
				if v16 != 0 {
					return int32(0)
				} else {
					v17 = F_add_size(m, v12, v15)
					mBase = m.M
					v18 = m.ExcPending
					if v18 != 0 {
						return int32(0)
					} else {
						return v17
					}
				}
			}
		}
	}
}
