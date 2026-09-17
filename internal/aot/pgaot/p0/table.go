package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ExecAllocTableSlot(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	v4 = F_MakeTupleTableSlot(m, l1, l2)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v9 = F_lappend(m, v8, v4)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0))) = v9
			return v4
		}
	}
}
func F_ExecMakeTableFunctionResult(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
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
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
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
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v93 int32
	_ = v93
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v148 int32
	_ = v148
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v162 int64
	_ = v162
	var v164 int32
	_ = v164
	var v181 int32
	_ = v181
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v248 int64
	_ = v248
	var v250 int64
	_ = v250
	var v251 int64
	_ = v251
	var v252 int64
	_ = v252
	var v256 int64
	_ = v256
	var v257 int64
	_ = v257
	var v260 int64
	_ = v260
	var v262 int64
	_ = v262
	var v267 int64
	_ = v267
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v286 int32
	_ = v286
	var v291 int32
	_ = v291
	var v299 int32
	_ = v299
	var v302 int32
	_ = v302
	var v306 int32
	_ = v306
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v388 int32
	_ = v388
	var v392 int32
	_ = v392
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v404 int32
	_ = v404
	var v407 int32
	_ = v407
	var v411 int32
	_ = v411
	var v416 int32
	_ = v416
	var v420 int32
	_ = v420
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v428 int32
	_ = v428
	var v433 int32
	_ = v433
	var v438 int32
	_ = v438
	var v441 int32
	_ = v441
	var v445 int32
	_ = v445
	var v450 int32
	_ = v450
	var v462 int32
	_ = v462
	var v466 int32
	_ = v466
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
	var v476 int32
	_ = v476
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
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
	var v499 int32
	_ = v499
	var v503 int32
	_ = v503
	var v507 int32
	_ = v507
	v6 = int32(0)
	v16 = m.G0
	v18 = v16 - int32(96)
	m.G0 = v18
	F_MemoryContextReset(m, l2)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v24 = int32(_a_F_ExecMakeTableFunctionResult_0)
	v25 = *(*int32)(unsafe.Add(mBase, _c_F_ExecMakeTableFunctionResult[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecMakeTableFunctionResult[0])) = l2
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v29 = F_exprType(m, v28)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v31 = F_type_is_rowtype(m, v29)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v18)+56)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+48)) = int32(1)
	if l4 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v39 = int32(15)
	goto L7
L6:
	;
	v39 = int32(11)
	goto L7
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+44)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v18)+40)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v18)+36)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v18)+32)) = int32(383)
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v45 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
	v52 = v46<<(uint(int32(3))%32) + int32(20)
	goto L10
L9:
	;
	v52 = int32(20)
	goto L10
L10:
	;
	v53 = F_palloc(m, v52)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v55 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	v466 = *(*int32)(unsafe.Add(mBase, uint32(v18)+56))
	if v466 != 0 {
		goto L112
	} else {
		goto L113
	}
L13:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecMakeTableFunctionResult[0])) = v186
	v189 = v53 + int32(16)
	v191 = int32(0)
	v199 = int32(1)
	v201 = v191
	v202 = v191
	goto L33
L14:
	;
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+57)))
	v59 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v53)+4)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v53))) = l0 + int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v53)+8)) = v18 + int32(32)
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)+12))
	*(*uint8)(unsafe.Add(mBase, uint32(v53)+16)) = uint8(v59)
	*(*int32)(unsafe.Add(mBase, uint32(v53)+12)) = v68
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v72 != 0 {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	goto L16
L16:
	;
	v162 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v53))) = v162
	v164 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v53)+18)) = uint16(v164)
	*(*int64)(unsafe.Add(mBase, uint32(v53)+8)) = v162
	*(*uint8)(unsafe.Add(mBase, uint32(v53)+16)) = uint8(v164)
	v181 = v6
	goto L13
L17:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v72)+4))
	v75 = v73
	goto L19
L18:
	;
	v75 = int32(0)
	goto L19
L19:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v53)+18)) = uint16(v75)
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v77 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+26)))
	if v133 != int32(1) {
		v181 = v58
		goto L13
	} else {
		goto L27
	}
L21:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v77)+4))
	if v80 <= int32(0) {
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
	v102 = v53 + int32(20) + v93<<(uint(int32(3))%32)
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v77)+12))
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v103+v93<<(uint(int32(2))%32))))
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v107)+20))
	v111 = m.T0[v110].(func(*base.Module, int32, int32, int32) int32)(m, v107, l1, v102+int32(4))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L1
	} else {
		goto L25
	}
L24:
	;
	goto L20
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v102))) = v111
	v115 = v93 + int32(1)
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v77)+4))
	if v115 < v116 {
		v93 = v115
		goto L23
	} else {
		goto L26
	}
L26:
	;
	goto L24
L27:
	;
	v136 = int32(0)
	v137 = int32(*(*int16)(unsafe.Add(mBase, uint32(v53)+18)))
	if v137 <= v136 {
		v181 = v58
		goto L13
	} else {
		goto L28
	}
L28:
	;
	v148 = v136
	goto L29
L29:
	;
	v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53+v148<<(uint(int32(3))%32))+24)))
	if v158 != 0 {
		v462 = v58
		goto L12
	} else {
		goto L31
	}
L30:
	;
	v181 = v58
	goto L13
L31:
	;
	v160 = v148 + int32(1)
	if v137 != v160 {
		v148 = v160
		goto L29
	} else {
		goto L32
	}
L32:
	;
	goto L30
L33:
	;
	v209 = *(*int32)(unsafe.Add(mBase, _c_F_ExecMakeTableFunctionResult[1]))
	if v209 != 0 {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v438 = m.ExcPending
	if v438 != 0 {
		goto L1
	} else {
		goto L108
	}
L35:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L1
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	F_MemoryContextReset(m, v212)
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L1
	} else {
		goto L39
	}
L38:
	;
	goto L37
L39:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v215 == int32(0) {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v18)+48))
	if v286 != int32(1) {
		goto L57
	} else {
		goto L58
	}
L41:
	;
	v219 = v18 - int32(-64)
	F_pgstat_init_function_usage(m, v53, v219)
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L1
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v215)+20))
	v280 = m.T0[v279].(func(*base.Module, int32, int32, int32) int32)(m, v215, l1, v189)
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L1
	} else {
		goto L53
	}
L44:
	;
	v222 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v53)+16)) = uint8(v222)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+52)) = v222
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v226)))
	v228 = m.T0[v227].(func(*base.Module, int32) int32)(m, v53)
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = v228
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v18)+52))
	v240 = m.G0
	v242 = v240 - int32(16)
	m.G0 = v242
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v219)))
	if v244 != 0 {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	goto L40
L47:
	;
	F___clock_gettime(m, int32(1), v242)
	mBase = m.M
	v247 = int32(_a_F_ExecMakeTableFunctionResult_1)
	v248 = *(*int64)(unsafe.Add(mBase, _c_F_ExecMakeTableFunctionResult[2]))
	v250 = *(*int64)(unsafe.Add(mBase, uint32(v219)+16))
	v251 = int64(*(*int32)(unsafe.Add(mBase, uint32(v242)+8)))
	v252 = *(*int64)(unsafe.Add(mBase, uint32(v242)))
	v256 = *(*int64)(unsafe.Add(mBase, uint32(v219)+24))
	v257 = v251 + v252*int64(1000000000) - v256
	*(*int64)(unsafe.Add(mBase, _c_F_ExecMakeTableFunctionResult[2])) = v250 + v257
	v260 = *(*int64)(unsafe.Add(mBase, uint32(v219)+8))
	if v231 != int32(1) {
		goto L50
	} else {
		goto L51
	}
L48:
	;
	goto L49
L49:
	;
	m.G0 = v242 + int32(16)
	goto L46
L50:
	;
	v262 = *(*int64)(unsafe.Add(mBase, uint32(v244)))
	*(*int64)(unsafe.Add(mBase, uint32(v244))) = v262 + int64(1)
	goto L52
L51:
	;
	goto L52
L52:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v244)+8)) = v260 + v257
	v267 = *(*int64)(unsafe.Add(mBase, uint32(v244)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v244)+16)) = v267 + (v257 - v248 + v250)
	goto L49
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+52)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = v280
	goto L40
L54:
	;
	goto L34
L55:
	;
	v199 = int32(0)
	v201 = v395
	v202 = v343
	goto L33
L56:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v420 = m.ExcPending
	if v420 != 0 {
		goto L1
	} else {
		goto L104
	}
L57:
	;
	if v286 != int32(2) {
		goto L56
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v18)+52))
	if v312 == int32(2) {
		v462 = v181
		goto L12
	} else {
		goto L66
	}
L60:
	;
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v18)+52))
	if v199&base.B2i32(v291 == int32(0))&v181 != 0 {
		v462 = v181
		goto L12
	} else {
		goto L61
	}
L61:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	F_errcode(m, int32(33686083))
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	F_errmsg(m, int32(_a_F_ExecMakeTableFunctionResult_2), int32(0))
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	F_errfinish(m, int32(_a_F_ExecMakeTableFunctionResult_3), int32(373), int32(_a_F_ExecMakeTableFunctionResult_4))
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
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
	if v199 != 0 {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v315 = int32(_a_F_ExecMakeTableFunctionResult_0)
	v316 = *(*int32)(unsafe.Add(mBase, _c_F_ExecMakeTableFunctionResult[0]))
	v318 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecMakeTableFunctionResult[0])) = v318
	v322 = *(*int32)(unsafe.Add(mBase, _c_F_ExecMakeTableFunctionResult[3]))
	v323 = F_tuplestore_begin_heap(m, l4, int32(0), v322)
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L1
	} else {
		goto L70
	}
L68:
	;
	v342 = v201
	v343 = v202
	goto L69
L69:
	;
	if v31 != 0 {
		goto L77
	} else {
		goto L78
	}
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+56)) = v323
	if v31 == int32(0) {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v329 = F_CreateTemplateTupleDesc(m, int32(1))
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L1
	} else {
		goto L74
	}
L72:
	;
	v338 = v201
	goto L73
L73:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecMakeTableFunctionResult[0])) = v316
	v342 = v338
	v343 = v323
	goto L69
L74:
	;
	F_TupleDescInitEntry(m, v329, int32(1), int32(_a_F_ExecMakeTableFunctionResult_5), v29, int32(-1), int32(0))
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L1
	} else {
		goto L75
	}
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+60)) = v329
	v338 = v329
	goto L73
L76:
	;
	v396 = *(*int32)(unsafe.Add(mBase, uint32(v18)+52))
	if v396 != int32(1) {
		v462 = v181
		goto L12
	} else {
		goto L98
	}
L77:
	;
	v344 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v189))))
	if v344 == int32(0) {
		goto L80
	} else {
		goto L81
	}
L78:
	;
	goto L79
L79:
	;
	F_tuplestore_putvalues(m, v343, v342, v18+int32(8), v189)
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L1
	} else {
		goto L97
	}
L80:
	;
	v347 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
	v348 = F_pg_detoast_datum(m, v347)
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L1
	} else {
		goto L83
	}
L81:
	;
	goto L82
L82:
	;
	v381 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v382 = F_palloc(m, v381)
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L1
	} else {
		goto L92
	}
L83:
	;
	if v342 == int32(0) {
		goto L85
	} else {
		goto L86
	}
L84:
	;
	v372 = *(*int32)(unsafe.Add(mBase, uint32(v348)))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+28)) = v348
	*(*int32)(unsafe.Add(mBase, uint32(v18)+12)) = int32(base.Ui32(v372) >> (uint(int32(2)) % 32))
	F_tuplestore_puttuple(m, v343, v18+int32(12))
	mBase = m.M
	v380 = m.ExcPending
	if v380 != 0 {
		goto L1
	} else {
		goto L91
	}
L85:
	;
	v352 = int32(_a_F_ExecMakeTableFunctionResult_0)
	v353 = *(*int32)(unsafe.Add(mBase, _c_F_ExecMakeTableFunctionResult[0]))
	v355 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecMakeTableFunctionResult[0])) = v355
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v348)+8))
	v358 = *(*int32)(unsafe.Add(mBase, uint32(v348)+4))
	v359 = F_lookup_rowtype_tupdesc_copy(m, v357, v358)
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L1
	} else {
		goto L88
	}
L86:
	;
	goto L87
L87:
	;
	v364 = *(*int32)(unsafe.Add(mBase, uint32(v348)+8))
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v342)+4))
	if v364 != v365 {
		goto L54
	} else {
		goto L89
	}
L88:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecMakeTableFunctionResult[0])) = v353
	*(*int32)(unsafe.Add(mBase, uint32(v18)+60)) = v359
	v371 = v359
	goto L84
L89:
	;
	v367 = *(*int32)(unsafe.Add(mBase, uint32(v348)+4))
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v342)+8))
	if v367 != v368 {
		goto L54
	} else {
		goto L90
	}
L90:
	;
	v371 = v342
	goto L84
L91:
	;
	v395 = v371
	goto L76
L92:
	;
	if v381 != 0 {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	base.MemoryFill(m, v382, int32(1), v381)
	goto L95
L94:
	;
	goto L95
L95:
	;
	F_tuplestore_putvalues(m, v343, l3, int32(0), v382)
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L1
	} else {
		goto L96
	}
L96:
	;
	v395 = v342
	goto L76
L97:
	;
	v395 = v342
	goto L76
L98:
	;
	if v181&int32(1) != 0 {
		goto L55
	} else {
		goto L99
	}
L99:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v404 = m.ExcPending
	if v404 != 0 {
		goto L1
	} else {
		goto L100
	}
L100:
	;
	F_errcode(m, int32(33686083))
	mBase = m.M
	v407 = m.ExcPending
	if v407 != 0 {
		goto L1
	} else {
		goto L101
	}
L101:
	;
	F_errmsg(m, int32(_a_F_ExecMakeTableFunctionResult_6), int32(0))
	mBase = m.M
	v411 = m.ExcPending
	if v411 != 0 {
		goto L1
	} else {
		goto L102
	}
L102:
	;
	F_errfinish(m, int32(_a_F_ExecMakeTableFunctionResult_3), int32(365), int32(_a_F_ExecMakeTableFunctionResult_4))
	mBase = m.M
	v416 = m.ExcPending
	if v416 != 0 {
		goto L1
	} else {
		goto L103
	}
L103:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L104:
	;
	F_errcode(m, int32(33686083))
	mBase = m.M
	v423 = m.ExcPending
	if v423 != 0 {
		goto L1
	} else {
		goto L105
	}
L105:
	;
	v424 = *(*int32)(unsafe.Add(mBase, uint32(v18)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v424
	F_errmsg(m, int32(_a_F_ExecMakeTableFunctionResult_7), v18)
	mBase = m.M
	v428 = m.ExcPending
	if v428 != 0 {
		goto L1
	} else {
		goto L106
	}
L106:
	;
	F_errfinish(m, int32(_a_F_ExecMakeTableFunctionResult_3), int32(381), int32(_a_F_ExecMakeTableFunctionResult_4))
	mBase = m.M
	v433 = m.ExcPending
	if v433 != 0 {
		goto L1
	} else {
		goto L107
	}
L107:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L108:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v441 = m.ExcPending
	if v441 != 0 {
		goto L1
	} else {
		goto L109
	}
L109:
	;
	F_errmsg(m, int32(_a_F_ExecMakeTableFunctionResult_8), int32(0))
	mBase = m.M
	v445 = m.ExcPending
	if v445 != 0 {
		goto L1
	} else {
		goto L110
	}
L110:
	;
	F_errfinish(m, int32(_a_F_ExecMakeTableFunctionResult_3), int32(315), int32(_a_F_ExecMakeTableFunctionResult_4))
	mBase = m.M
	v450 = m.ExcPending
	if v450 != 0 {
		goto L1
	} else {
		goto L111
	}
L111:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L112:
	;
	v493 = *(*int32)(unsafe.Add(mBase, uint32(v18)+60))
	if v493 == int32(0) {
		goto L121
	} else {
		goto L122
	}
L113:
	;
	v467 = int32(_a_F_ExecMakeTableFunctionResult_0)
	v468 = *(*int32)(unsafe.Add(mBase, _c_F_ExecMakeTableFunctionResult[0]))
	v470 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecMakeTableFunctionResult[0])) = v470
	v474 = *(*int32)(unsafe.Add(mBase, _c_F_ExecMakeTableFunctionResult[3]))
	v475 = F_tuplestore_begin_heap(m, l4, int32(0), v474)
	mBase = m.M
	v476 = m.ExcPending
	if v476 != 0 {
		goto L1
	} else {
		goto L114
	}
L114:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecMakeTableFunctionResult[0])) = v468
	*(*int32)(unsafe.Add(mBase, uint32(v18)+56)) = v475
	if v462&int32(1) != 0 {
		goto L112
	} else {
		goto L115
	}
L115:
	;
	v482 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v483 = F_palloc(m, v482)
	mBase = m.M
	v484 = m.ExcPending
	if v484 != 0 {
		goto L1
	} else {
		goto L116
	}
L116:
	;
	if v482 != 0 {
		goto L117
	} else {
		goto L118
	}
L117:
	;
	base.MemoryFill(m, v483, int32(1), v482)
	goto L119
L118:
	;
	goto L119
L119:
	;
	F_tuplestore_putvalues(m, v475, l3, int32(0), v483)
	mBase = m.M
	v489 = m.ExcPending
	if v489 != 0 {
		goto L1
	} else {
		goto L120
	}
L120:
	;
	goto L112
L121:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecMakeTableFunctionResult[0])) = v25
	v507 = *(*int32)(unsafe.Add(mBase, uint32(v18)+56))
	m.G0 = v18 + int32(96)
	return v507
L122:
	;
	F_tupledesc_match(m, l3, v493)
	mBase = m.M
	v497 = m.ExcPending
	if v497 != 0 {
		goto L1
	} else {
		goto L123
	}
L123:
	;
	v498 = *(*int32)(unsafe.Add(mBase, uint32(v18)+60))
	v499 = *(*int32)(unsafe.Add(mBase, uint32(v498)+12))
	if v499 != int32(-1) {
		goto L121
	} else {
		goto L124
	}
L124:
	;
	F_FreeTupleDesc(m, v498)
	mBase = m.M
	v503 = m.ExcPending
	if v503 != 0 {
		goto L1
	} else {
		goto L125
	}
L125:
	;
	goto L121
}
func F_ExecTableFuncScan(m *base.Module, l0 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = F_ExecScan(m, l0, int32(762), int32(763))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_FetchTableStates(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v67 int64
	_ = v67
	var v69 int64
	_ = v69
	var v71 int64
	_ = v71
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v92 int32
	_ = v92
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
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v129 int32
	_ = v129
	var v135 int32
	_ = v135
	var v141 int32
	_ = v141
	v2 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0))) = uint8(v2)
	v9 = *(*int32)(unsafe.Add(mBase, _c_F_FetchTableStates[0]))
	if v9 == int32(2) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v141 & int32(1)
L2:
	;
	v13 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_FetchTableStates[1])))
	v141 = v13
	goto L1
L3:
	;
	goto L4
L4:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_FetchTableStates[0])) = int32(1)
	v18 = *(*int32)(unsafe.Add(mBase, _c_F_FetchTableStates[2]))
	F_list_free_deep(m, v18)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	return int32(0)
L6:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_FetchTableStates[2])) = int32(0)
	v27 = *(*int32)(unsafe.Add(mBase, _c_F_FetchTableStates[3]))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+20))
	goto L7
L7:
	;
	if base.B2i32(v28 == int32(2)) == int32(0) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	F_StartTransactionCommand(m)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L5
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	v38 = *(*int32)(unsafe.Add(mBase, _c_F_FetchTableStates[4]))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
	v41 = F_GetSubscriptionRelations(m, v39, int32(1))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L5
	} else {
		goto L12
	}
L11:
	;
	v35 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0))) = uint8(v35)
	goto L10
L12:
	;
	v43 = int32(_a_F_FetchTableStates_0)
	v44 = *(*int32)(unsafe.Add(mBase, _c_F_FetchTableStates[5]))
	v47 = *(*int32)(unsafe.Add(mBase, _c_F_FetchTableStates[6]))
	*(*int32)(unsafe.Add(mBase, _c_F_FetchTableStates[5])) = v47
	if v41 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_FetchTableStates[5])) = v44
	v92 = *(*int32)(unsafe.Add(mBase, _c_F_FetchTableStates[2]))
	if v92 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L14:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v41)+4))
	if v51 <= int32(0) {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v57 = v2
	goto L16
L16:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v41)+12))
	v61 = F_palloc(m, int32(24))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L5
	} else {
		goto L18
	}
L17:
	;
	goto L13
L18:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v59+v57<<(uint(int32(2))%32))))
	v67 = *(*int64)(unsafe.Add(mBase, uint32(v66)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v61)+16)) = v67
	v69 = *(*int64)(unsafe.Add(mBase, uint32(v66)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v61)+8)) = v69
	v71 = *(*int64)(unsafe.Add(mBase, uint32(v66)))
	*(*int64)(unsafe.Add(mBase, uint32(v61))) = v71
	v73 = int32(_a_F_FetchTableStates_1)
	v75 = *(*int32)(unsafe.Add(mBase, _c_F_FetchTableStates[2]))
	v76 = F_lappend(m, v75, v61)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L5
	} else {
		goto L19
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_FetchTableStates[2])) = v76
	v80 = v57 + int32(1)
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v41)+4))
	if v80 < v81 {
		v57 = v80
		goto L16
	} else {
		goto L20
	}
L20:
	;
	goto L17
L21:
	;
	v96 = *(*int32)(unsafe.Add(mBase, _c_F_FetchTableStates[4]))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v96)))
	v98 = m.G0
	v100 = v98 - int32(48)
	m.G0 = v100
	v104 = F_table_open(m, int32(_a_F_FetchTableStates_2), int32(1))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L5
	} else {
		goto L24
	}
L22:
	;
	v129 = int32(1)
	goto L23
L23:
	;
	*(*uint8)(unsafe.Add(mBase, _c_F_FetchTableStates[1])) = uint8(v129)
	v135 = *(*int32)(unsafe.Add(mBase, _c_F_FetchTableStates[0]))
	if v135 != int32(1) {
		v141 = v129
		goto L1
	} else {
		goto L30
	}
L24:
	;
	F_ScanKeyInit(m, v100, int32(1), int32(3), int32(184), v97)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L5
	} else {
		goto L25
	}
L25:
	;
	v111 = int32(0)
	v115 = F_systable_beginscan(m, v104, v111, v111, v111, int32(1), v100)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L5
	} else {
		goto L26
	}
L26:
	;
	v117 = F_systable_getnext(m, v115)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L5
	} else {
		goto L27
	}
L27:
	;
	F_systable_endscan(m, v115)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L5
	} else {
		goto L28
	}
L28:
	;
	F_relation_close(m, v104, int32(1))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L5
	} else {
		goto L29
	}
L29:
	;
	m.G0 = v100 + int32(48)
	v129 = base.B2i32(v117 != int32(0))
	goto L23
L30:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_FetchTableStates[0])) = int32(2)
	v141 = v129
	goto L1
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
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
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
	v70 = v27 + int32(1)
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	if v70 < v71 {
		v27 = v70
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
	v61 = int32(0)
	v64 = F_SearchSysCacheExists(m, int32(57), v32, v61, v61, v61)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
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
	if v64 != 0 {
		goto L8
	} else {
		goto L24
	}
L24:
	;
	F_UnlockRelationOid(m, v32, l1)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
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
func F_table_block_parallelscan_estimate(m *base.Module, l0 int32) int32 {
	return int32(40)
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
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
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
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
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
			v28 = F_MakeTupleTableSlot(m, v27, v26)
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return int32(0)
			} else {
				v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
				v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+44))
				v32 = m.T0[v31].(func(*base.Module, int32) int32)(m, l0)
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return int32(0)
				} else {
					v35 = *(*int32)(unsafe.Add(mBase, _c_F_table_index_fetch_tuple_check[0]))
					if v35 == int32(0) {
						v57 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
						v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)+188))
						v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)+56))
						v60 = m.T0[v59].(func(*base.Module, int32, int32, int32, int32, int32, int32) int32)(m, v32, l1, l2, v28, v9+int32(15), l3)
						mBase = m.M
						v61 = m.ExcPending
						if v61 != 0 {
							return int32(0)
						} else {
							v62 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
							v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)+188))
							v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)+52))
							m.T0[v64].(func(*base.Module, int32))(m, v32)
							mBase = m.M
							v66 = m.ExcPending
							if v66 != 0 {
								return int32(0)
							} else {
								F_ExecDropSingleTupleTableSlot(m, v28)
								mBase = m.M
								v68 = m.ExcPending
								if v68 != 0 {
									return int32(0)
								} else {
									m.G0 = v9 + int32(16)
									return v60
								}
							}
						}
					} else {
						v39 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_table_index_fetch_tuple_check[1])))
						if v39&int32(1) != 0 {
							v57 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
							v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)+188))
							v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)+56))
							v60 = m.T0[v59].(func(*base.Module, int32, int32, int32, int32, int32, int32) int32)(m, v32, l1, l2, v28, v9+int32(15), l3)
							mBase = m.M
							v61 = m.ExcPending
							if v61 != 0 {
								return int32(0)
							} else {
								v62 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
								v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)+188))
								v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)+52))
								m.T0[v64].(func(*base.Module, int32))(m, v32)
								mBase = m.M
								v66 = m.ExcPending
								if v66 != 0 {
									return int32(0)
								} else {
									F_ExecDropSingleTupleTableSlot(m, v28)
									mBase = m.M
									v68 = m.ExcPending
									if v68 != 0 {
										return int32(0)
									} else {
										m.G0 = v9 + int32(16)
										return v60
									}
								}
							}
						} else {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v45 = m.ExcPending
							if v45 != 0 {
								return int32(0)
							} else {
								F_errmsg_internal(m, int32(_a_F_table_index_fetch_tuple_check_0), int32(0))
								mBase = m.M
								v49 = m.ExcPending
								if v49 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_table_index_fetch_tuple_check_1), int32(1218), int32(_a_F_table_index_fetch_tuple_check_2))
									mBase = m.M
									v54 = m.ExcPending
									if v54 != 0 {
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
		v28 = F_MakeTupleTableSlot(m, v27, v26)
		mBase = m.M
		v29 = m.ExcPending
		if v29 != 0 {
			return int32(0)
		} else {
			v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
			v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+44))
			v32 = m.T0[v31].(func(*base.Module, int32) int32)(m, l0)
			mBase = m.M
			v33 = m.ExcPending
			if v33 != 0 {
				return int32(0)
			} else {
				v35 = *(*int32)(unsafe.Add(mBase, _c_F_table_index_fetch_tuple_check[0]))
				if v35 == int32(0) {
					v57 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
					v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)+188))
					v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)+56))
					v60 = m.T0[v59].(func(*base.Module, int32, int32, int32, int32, int32, int32) int32)(m, v32, l1, l2, v28, v9+int32(15), l3)
					mBase = m.M
					v61 = m.ExcPending
					if v61 != 0 {
						return int32(0)
					} else {
						v62 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
						v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)+188))
						v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)+52))
						m.T0[v64].(func(*base.Module, int32))(m, v32)
						mBase = m.M
						v66 = m.ExcPending
						if v66 != 0 {
							return int32(0)
						} else {
							F_ExecDropSingleTupleTableSlot(m, v28)
							mBase = m.M
							v68 = m.ExcPending
							if v68 != 0 {
								return int32(0)
							} else {
								m.G0 = v9 + int32(16)
								return v60
							}
						}
					}
				} else {
					v39 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_table_index_fetch_tuple_check[1])))
					if v39&int32(1) != 0 {
						v57 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
						v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)+188))
						v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)+56))
						v60 = m.T0[v59].(func(*base.Module, int32, int32, int32, int32, int32, int32) int32)(m, v32, l1, l2, v28, v9+int32(15), l3)
						mBase = m.M
						v61 = m.ExcPending
						if v61 != 0 {
							return int32(0)
						} else {
							v62 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
							v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)+188))
							v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)+52))
							m.T0[v64].(func(*base.Module, int32))(m, v32)
							mBase = m.M
							v66 = m.ExcPending
							if v66 != 0 {
								return int32(0)
							} else {
								F_ExecDropSingleTupleTableSlot(m, v28)
								mBase = m.M
								v68 = m.ExcPending
								if v68 != 0 {
									return int32(0)
								} else {
									m.G0 = v9 + int32(16)
									return v60
								}
							}
						}
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v45 = m.ExcPending
						if v45 != 0 {
							return int32(0)
						} else {
							F_errmsg_internal(m, int32(_a_F_table_index_fetch_tuple_check_0), int32(0))
							mBase = m.M
							v49 = m.ExcPending
							if v49 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_table_index_fetch_tuple_check_1), int32(1218), int32(_a_F_table_index_fetch_tuple_check_2))
								mBase = m.M
								v54 = m.ExcPending
								if v54 != 0 {
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
func F_table_parallelscan_estimate(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
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
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	switch v4 {
	case 0, 5:
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
	default:
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
	}
}
