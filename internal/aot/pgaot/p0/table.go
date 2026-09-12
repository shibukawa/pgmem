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
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v87 int32
	_ = v87
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
	var v144 int32
	_ = v144
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v164 int64
	_ = v164
	var v166 int32
	_ = v166
	var v183 int32
	_ = v183
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v252 int64
	_ = v252
	var v254 int64
	_ = v254
	var v255 int64
	_ = v255
	var v256 int64
	_ = v256
	var v260 int64
	_ = v260
	var v261 int64
	_ = v261
	var v264 int64
	_ = v264
	var v266 int64
	_ = v266
	var v271 int64
	_ = v271
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v290 int32
	_ = v290
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v303 int32
	_ = v303
	var v312 int32
	_ = v312
	var v315 int32
	_ = v315
	var v319 int32
	_ = v319
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v335 int32
	_ = v335
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v353 int32
	_ = v353
	var v355 int32
	_ = v355
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v372 int32
	_ = v372
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v404 int32
	_ = v404
	var v406 int32
	_ = v406
	var v410 int32
	_ = v410
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v422 int32
	_ = v422
	var v425 int32
	_ = v425
	var v429 int32
	_ = v429
	var v434 int32
	_ = v434
	var v438 int32
	_ = v438
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v446 int32
	_ = v446
	var v451 int32
	_ = v451
	var v456 int32
	_ = v456
	var v459 int32
	_ = v459
	var v463 int32
	_ = v463
	var v468 int32
	_ = v468
	var v480 int32
	_ = v480
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v488 int32
	_ = v488
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v506 int32
	_ = v506
	var v508 int32
	_ = v508
	var v511 int32
	_ = v511
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v521 int32
	_ = v521
	var v525 int32
	_ = v525
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
	v24 = int32(4470400)
	v25 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = l2
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
	v484 = *(*int32)(unsafe.Add(mBase, uint32(v18)+56))
	if v484 != 0 {
		goto L113
	} else {
		goto L114
	}
L13:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v188
	v191 = v53 + int32(16)
	v192 = int32(0)
	v202 = v192
	v203 = v192
	v204 = v192
	goto L33
L14:
	;
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+57)))
	v59 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v53)+4)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v53))) = l0 + int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v53)+8)) = v18 + int32(32)
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)+12))
	*(*uint8)(unsafe.Add(mBase, uint32(v53)+16)) = uint8(v59)
	*(*int32)(unsafe.Add(mBase, uint32(v53)+12)) = v69
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
	v164 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v53))) = v164
	v166 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v53)+18)) = uint16(v166)
	*(*int64)(unsafe.Add(mBase, uint32(v53)+8)) = v164
	*(*uint8)(unsafe.Add(mBase, uint32(v53)+16)) = uint8(v166)
	v183 = int32(0)
	goto L13
L17:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v73)+4))
	v75 = v74
	goto L19
L18:
	;
	v75 = v59
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
		v183 = v58
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
	v87 = v59
	goto L23
L23:
	;
	v102 = v53 + int32(20) + v87<<(uint(int32(3))%32)
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v77)+12))
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v103+v87<<(uint(int32(2))%32))))
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
	v115 = v87 + int32(1)
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v77)+4))
	if v115 < v116 {
		v87 = v115
		goto L23
	} else {
		goto L26
	}
L26:
	;
	goto L24
L27:
	;
	v136 = int32(*(*int16)(unsafe.Add(mBase, uint32(v53)+18)))
	if v136 <= int32(0) {
		v183 = v58
		goto L13
	} else {
		goto L28
	}
L28:
	;
	v144 = int32(0)
	goto L29
L29:
	;
	v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53+int32(24)+v144<<(uint(int32(3))%32)))))
	if v160 != 0 {
		v480 = v58
		goto L12
	} else {
		goto L31
	}
L30:
	;
	v183 = v58
	goto L13
L31:
	;
	v162 = v144 + int32(1)
	if v136 != v162 {
		v144 = v162
		goto L29
	} else {
		goto L32
	}
L32:
	;
	goto L30
L33:
	;
	v211 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	if v211 != 0 {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v456 = m.ExcPending
	if v456 != 0 {
		goto L1
	} else {
		goto L109
	}
L35:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L1
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	F_MemoryContextReset(m, v214)
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L1
	} else {
		goto L39
	}
L38:
	;
	goto L37
L39:
	;
	v217 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v217 == int32(0) {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v18)+48))
	if v290 != int32(1) {
		goto L57
	} else {
		goto L58
	}
L41:
	;
	F_pgstat_init_function_usage(m, v53, v18-int32(-64))
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L1
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v217)+20))
	v284 = m.T0[v283].(func(*base.Module, int32, int32, int32) int32)(m, v217, l1, v191)
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L1
	} else {
		goto L53
	}
L44:
	;
	v224 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v53)+16)) = uint8(v224)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+52)) = v224
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v228)))
	v230 = m.T0[v229].(func(*base.Module, int32) int32)(m, v53)
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = v230
	v234 = v18 - int32(-64)
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v18)+52))
	v244 = m.G0
	v246 = v244 - int32(16)
	m.G0 = v246
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v234)))
	if v248 != 0 {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	goto L40
L47:
	;
	F___clock_gettime(m, int32(1), v246)
	mBase = m.M
	v251 = int32(4450176)
	v252 = *(*int64)(unsafe.Add(mBase, _consts[265]))
	v254 = *(*int64)(unsafe.Add(mBase, uint32(v234)+16))
	v255 = int64(*(*int32)(unsafe.Add(mBase, uint32(v246)+8)))
	v256 = *(*int64)(unsafe.Add(mBase, uint32(v246)))
	v260 = *(*int64)(unsafe.Add(mBase, uint32(v234)+24))
	v261 = v255 + v256*int64(1000000000) - v260
	*(*int64)(unsafe.Add(mBase, _consts[265])) = v254 + v261
	v264 = *(*int64)(unsafe.Add(mBase, uint32(v234)+8))
	if v235 != int32(1) {
		goto L50
	} else {
		goto L51
	}
L48:
	;
	goto L49
L49:
	;
	m.G0 = v246 + int32(16)
	goto L46
L50:
	;
	v266 = *(*int64)(unsafe.Add(mBase, uint32(v248)))
	*(*int64)(unsafe.Add(mBase, uint32(v248))) = v266 + int64(1)
	goto L52
L51:
	;
	goto L52
L52:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v248)+8)) = v264 + v261
	v271 = *(*int64)(unsafe.Add(mBase, uint32(v248)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v248)+16)) = v271 + (v261 - v252 + v254)
	goto L49
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+52)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = v284
	goto L40
L54:
	;
	goto L34
L55:
	;
	v202 = int32(1)
	v203 = v413
	v204 = v360
	goto L33
L56:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v438 = m.ExcPending
	if v438 != 0 {
		goto L1
	} else {
		goto L105
	}
L57:
	;
	if v290 != int32(2) {
		goto L56
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v18)+52))
	if v325 == int32(2) {
		v480 = v183
		goto L12
	} else {
		goto L69
	}
L60:
	;
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v18)+52))
	v296 = int32(0)
	if (v202|base.B2i32(v295 != v296))&int32(1) == v296 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v303 = int32(1)
	if (v183^v303)&v303 == int32(0) {
		v480 = v183
		goto L12
	} else {
		goto L64
	}
L62:
	;
	goto L63
L63:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L1
	} else {
		goto L65
	}
L64:
	;
	goto L63
L65:
	;
	F_errcode(m, int32(33686083))
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	F_errmsg(m, int32(431316), int32(0))
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	F_errfinish(m, int32(491150), int32(373), int32(96619))
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L1
	} else {
		goto L68
	}
L68:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L69:
	;
	if v202&int32(1) == int32(0) {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v332 = int32(4470400)
	v333 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v335 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v335
	v339 = *(*int32)(unsafe.Add(mBase, _consts[326]))
	v340 = F_tuplestore_begin_heap(m, l4, int32(0), v339)
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		goto L1
	} else {
		goto L73
	}
L71:
	;
	v359 = v203
	v360 = v204
	goto L72
L72:
	;
	if v31 != 0 {
		goto L80
	} else {
		goto L81
	}
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+56)) = v340
	if v31 == int32(0) {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v346 = F_CreateTemplateTupleDesc(m, int32(1))
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L1
	} else {
		goto L77
	}
L75:
	;
	v355 = v203
	goto L76
L76:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v333
	v359 = v355
	v360 = v340
	goto L72
L77:
	;
	F_TupleDescInitEntry(m, v346, int32(1), int32(270389), v29, int32(-1), int32(0))
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L1
	} else {
		goto L78
	}
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+60)) = v346
	v355 = v346
	goto L76
L79:
	;
	v414 = *(*int32)(unsafe.Add(mBase, uint32(v18)+52))
	if v414 != int32(1) {
		v480 = v183
		goto L12
	} else {
		goto L99
	}
L80:
	;
	v361 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v191))))
	if v361 == int32(0) {
		goto L83
	} else {
		goto L84
	}
L81:
	;
	goto L82
L82:
	;
	F_tuplestore_putvalues(m, v360, v359, v18+int32(8), v191)
	mBase = m.M
	v410 = m.ExcPending
	if v410 != 0 {
		goto L1
	} else {
		goto L98
	}
L83:
	;
	v364 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
	v365 = F_pg_detoast_datum(m, v364)
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L1
	} else {
		goto L86
	}
L84:
	;
	goto L85
L85:
	;
	v399 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v400 = F_palloc(m, v399)
	mBase = m.M
	v401 = m.ExcPending
	if v401 != 0 {
		goto L1
	} else {
		goto L95
	}
L86:
	;
	if v359 == int32(0) {
		goto L88
	} else {
		goto L89
	}
L87:
	;
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v365)))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+28)) = v365
	*(*int32)(unsafe.Add(mBase, uint32(v18)+12)) = int32(base.Ui32(v389) >> (uint(int32(2)) % 32))
	F_tuplestore_puttuple(m, v360, v18+int32(12))
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
		goto L1
	} else {
		goto L94
	}
L88:
	;
	v369 = int32(4470400)
	v370 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v372 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v372
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v365)+8))
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v365)+4))
	v376 = F_lookup_rowtype_tupdesc_copy(m, v374, v375)
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L1
	} else {
		goto L91
	}
L89:
	;
	goto L90
L90:
	;
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v365)+8))
	v382 = *(*int32)(unsafe.Add(mBase, uint32(v359)+4))
	if v381 != v382 {
		goto L54
	} else {
		goto L92
	}
L91:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v370
	*(*int32)(unsafe.Add(mBase, uint32(v18)+60)) = v376
	v388 = v376
	goto L87
L92:
	;
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v365)+4))
	v385 = *(*int32)(unsafe.Add(mBase, uint32(v359)+8))
	if v384 != v385 {
		goto L54
	} else {
		goto L93
	}
L93:
	;
	v388 = v359
	goto L87
L94:
	;
	v413 = v388
	goto L79
L95:
	;
	v404 = F__emscripten_memset_bulkmem(m, v400, base.I32_extend8_s(int32(1)), v399)
	mBase = m.M
	goto L96
L96:
	;
	F_tuplestore_putvalues(m, v360, l3, int32(0), v404)
	mBase = m.M
	v406 = m.ExcPending
	if v406 != 0 {
		goto L1
	} else {
		goto L97
	}
L97:
	;
	v413 = v359
	goto L79
L98:
	;
	v413 = v359
	goto L79
L99:
	;
	if v183&int32(1) != 0 {
		goto L55
	} else {
		goto L100
	}
L100:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v422 = m.ExcPending
	if v422 != 0 {
		goto L1
	} else {
		goto L101
	}
L101:
	;
	F_errcode(m, int32(33686083))
	mBase = m.M
	v425 = m.ExcPending
	if v425 != 0 {
		goto L1
	} else {
		goto L102
	}
L102:
	;
	F_errmsg(m, int32(431251), int32(0))
	mBase = m.M
	v429 = m.ExcPending
	if v429 != 0 {
		goto L1
	} else {
		goto L103
	}
L103:
	;
	F_errfinish(m, int32(491150), int32(365), int32(96619))
	mBase = m.M
	v434 = m.ExcPending
	if v434 != 0 {
		goto L1
	} else {
		goto L104
	}
L104:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L105:
	;
	F_errcode(m, int32(33686083))
	mBase = m.M
	v441 = m.ExcPending
	if v441 != 0 {
		goto L1
	} else {
		goto L106
	}
L106:
	;
	v442 = *(*int32)(unsafe.Add(mBase, uint32(v18)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v442
	F_errmsg(m, int32(478236), v18)
	mBase = m.M
	v446 = m.ExcPending
	if v446 != 0 {
		goto L1
	} else {
		goto L107
	}
L107:
	;
	F_errfinish(m, int32(491150), int32(381), int32(96619))
	mBase = m.M
	v451 = m.ExcPending
	if v451 != 0 {
		goto L1
	} else {
		goto L108
	}
L108:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L109:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v459 = m.ExcPending
	if v459 != 0 {
		goto L1
	} else {
		goto L110
	}
L110:
	;
	F_errmsg(m, int32(360877), int32(0))
	mBase = m.M
	v463 = m.ExcPending
	if v463 != 0 {
		goto L1
	} else {
		goto L111
	}
L111:
	;
	F_errfinish(m, int32(491150), int32(315), int32(96619))
	mBase = m.M
	v468 = m.ExcPending
	if v468 != 0 {
		goto L1
	} else {
		goto L112
	}
L112:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L113:
	;
	v511 = *(*int32)(unsafe.Add(mBase, uint32(v18)+60))
	if v511 == int32(0) {
		goto L120
	} else {
		goto L121
	}
L114:
	;
	v485 = int32(4470400)
	v486 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v488 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v488
	v492 = *(*int32)(unsafe.Add(mBase, _consts[326]))
	v493 = F_tuplestore_begin_heap(m, l4, int32(0), v492)
	mBase = m.M
	v494 = m.ExcPending
	if v494 != 0 {
		goto L1
	} else {
		goto L115
	}
L115:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v486
	*(*int32)(unsafe.Add(mBase, uint32(v18)+56)) = v493
	if v480&int32(1) != 0 {
		goto L113
	} else {
		goto L116
	}
L116:
	;
	v501 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v502 = F_palloc(m, v501)
	mBase = m.M
	v503 = m.ExcPending
	if v503 != 0 {
		goto L1
	} else {
		goto L117
	}
L117:
	;
	v506 = F__emscripten_memset_bulkmem(m, v502, base.I32_extend8_s(int32(1)), v501)
	mBase = m.M
	goto L118
L118:
	;
	F_tuplestore_putvalues(m, v493, l3, int32(0), v506)
	mBase = m.M
	v508 = m.ExcPending
	if v508 != 0 {
		goto L1
	} else {
		goto L119
	}
L119:
	;
	goto L113
L120:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v25
	v525 = *(*int32)(unsafe.Add(mBase, uint32(v18)+56))
	m.G0 = v18 + int32(96)
	return v525
L121:
	;
	F_tupledesc_match(m, l3, v511)
	mBase = m.M
	v515 = m.ExcPending
	if v515 != 0 {
		goto L1
	} else {
		goto L122
	}
L122:
	;
	v516 = *(*int32)(unsafe.Add(mBase, uint32(v18)+60))
	v517 = *(*int32)(unsafe.Add(mBase, uint32(v516)+12))
	if v517 != int32(-1) {
		goto L120
	} else {
		goto L123
	}
L123:
	;
	F_FreeTupleDesc(m, v516)
	mBase = m.M
	v521 = m.ExcPending
	if v521 != 0 {
		goto L1
	} else {
		goto L124
	}
L124:
	;
	goto L120
}
func F_ExecTableFuncScan(m *base.Module, l0 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = F_ExecScan(m, l0, int32(761), int32(762))
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
	var v136 int32
	_ = v136
	var v142 int32
	_ = v142
	v2 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0))) = uint8(v2)
	v9 = *(*int32)(unsafe.Add(mBase, _consts[534]))
	if v9 == int32(2) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v142 & int32(1)
L2:
	;
	v13 = int32(*(*uint8)(unsafe.Add(mBase, _consts[535])))
	v142 = v13
	goto L1
L3:
	;
	goto L4
L4:
	;
	*(*int32)(unsafe.Add(mBase, _consts[534])) = int32(1)
	v18 = *(*int32)(unsafe.Add(mBase, _consts[536]))
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
	*(*int32)(unsafe.Add(mBase, _consts[536])) = int32(0)
	v27 = *(*int32)(unsafe.Add(mBase, _consts[37]))
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
	v38 = *(*int32)(unsafe.Add(mBase, _consts[533]))
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
	v43 = int32(4470400)
	v44 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v47 = *(*int32)(unsafe.Add(mBase, _consts[207]))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v47
	if v41 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v44
	v92 = *(*int32)(unsafe.Add(mBase, _consts[536]))
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
	v67 = *(*int64)(unsafe.Add(mBase, uint32(v66)))
	*(*int64)(unsafe.Add(mBase, uint32(v61))) = v67
	v69 = *(*int64)(unsafe.Add(mBase, uint32(v66)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v61)+16)) = v69
	v71 = *(*int64)(unsafe.Add(mBase, uint32(v66)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v61)+8)) = v71
	v73 = int32(4380468)
	v75 = *(*int32)(unsafe.Add(mBase, _consts[536]))
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
	*(*int32)(unsafe.Add(mBase, _consts[536])) = v76
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
	v96 = *(*int32)(unsafe.Add(mBase, _consts[533]))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v96)))
	v98 = m.G0
	v100 = v98 - int32(48)
	m.G0 = v100
	v104 = F_table_open(m, int32(6102), int32(1))
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
	*(*uint8)(unsafe.Add(mBase, _consts[535])) = uint8(v129)
	v136 = *(*int32)(unsafe.Add(mBase, _consts[534]))
	if v136 != int32(1) {
		v142 = v129
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
	F_sequence_close(m, v104, int32(1))
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
	*(*int32)(unsafe.Add(mBase, _consts[534])) = int32(2)
	v142 = v129
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
	var v26 int32
	_ = v26
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
	v26 = v4
	goto L6
L6:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v28+v26<<(uint(int32(2))%32))))
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
	v70 = v26 + int32(1)
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	if v70 < v71 {
		v26 = v70
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
	F_errmsg(m, int32(679628), v10)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L2
	} else {
		goto L21
	}
L21:
	;
	F_errfinish(m, int32(485567), int32(144), int32(354391))
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
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
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
	return v41
L2:
	;
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+8)))
	v41 = base.B2i32(v38 == v39)
	goto L1
L3:
	;
	if v6 == int32(0) {
		v41 = v3
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
		v41 = v3
		goto L1
	} else {
		goto L16
	}
L6:
	;
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6))))
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
	if v13 == int32(0) {
		v32 = v12
		v33 = v13
		goto L8
	} else {
		goto L9
	}
L7:
	;
	if v33-v32 == int32(0) {
		goto L2
	} else {
		goto L15
	}
L8:
	;
	goto L7
L9:
	;
	if v12 != v13 {
		v32 = v12
		v33 = v13
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v17 = v7
	v18 = v6
	goto L11
L11:
	;
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+1)))
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+1)))
	if v22 == int32(0) {
		v32 = v21
		v33 = v22
		goto L8
	} else {
		goto L13
	}
L12:
	;
	v32 = v21
	v33 = v22
	goto L8
L13:
	;
	v25 = int32(1)
	if v21 == v22 {
		v17 = v17 + v25
		v18 = v18 + v25
		goto L11
	} else {
		goto L14
	}
L14:
	;
	goto L12
L15:
	;
	v41 = v3
	goto L1
L16:
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
			v28 = F_MakeSingleTupleTableSlot(m, v27, v26)
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
					v35 = *(*int32)(unsafe.Add(mBase, _consts[25]))
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
						v39 = int32(*(*uint8)(unsafe.Add(mBase, _consts[26])))
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
								F_errmsg_internal(m, int32(330635), int32(0))
								mBase = m.M
								v49 = m.ExcPending
								if v49 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(320974), int32(1218), int32(376834))
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
			v25 = int32(1591688)
		} else {
			v25 = int32(1591636)
		}
		v26 = v25
		v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
		v28 = F_MakeSingleTupleTableSlot(m, v27, v26)
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
				v35 = *(*int32)(unsafe.Add(mBase, _consts[25]))
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
					v39 = int32(*(*uint8)(unsafe.Add(mBase, _consts[26])))
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
							F_errmsg_internal(m, int32(330635), int32(0))
							mBase = m.M
							v49 = m.ExcPending
							if v49 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(320974), int32(1218), int32(376834))
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
