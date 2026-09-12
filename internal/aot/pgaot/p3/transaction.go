package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_GetNewTransactionId(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int64
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v178 int32
	_ = v178
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v193 int32
	_ = v193
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v223 int32
	_ = v223
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
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
	var v258 int32
	_ = v258
	var v259 int64
	_ = v259
	var v261 int32
	_ = v261
	var v267 int64
	_ = v267
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v277 int32
	_ = v277
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v295 int64
	_ = v295
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v305 int32
	_ = v305
	var v308 int64
	_ = v308
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v315 int32
	_ = v315
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v351 int64
	_ = v351
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v363 int32
	_ = v363
	var v366 int64
	_ = v366
	var v367 int32
	_ = v367
	var v369 int32
	_ = v369
	var v378 int32
	_ = v378
	var v380 int32
	_ = v380
	var v381 int64
	_ = v381
	var v382 int64
	_ = v382
	var v390 int64
	_ = v390
	var v394 int64
	_ = v394
	var v397 int64
	_ = v397
	var v402 int32
	_ = v402
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v429 int32
	_ = v429
	var v432 int32
	_ = v432
	var v434 int32
	_ = v434
	var v437 int32
	_ = v437
	var v445 int32
	_ = v445
	var v449 int32
	_ = v449
	var v458 int64
	_ = v458
	v11 = m.G0
	v13 = v11 + int32(-64)
	m.G0 = v13
	v19 = *(*int32)(unsafe.Add(mBase, _consts[72]))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+72))
	if v20 != 0 {
		goto L7
	} else {
		goto L8
	}
L1:
	;
	m.G0 = v13 - int32(-64)
	return v458
L2:
	;
	v268 = m.G0
	v270 = v268 - int32(16)
	m.G0 = v270
	if v261&int32(32767) != 0 {
		goto L76
	} else {
		goto L77
	}
L3:
	;
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v91))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v70)) == int32(0) {
		goto L59
	} else {
		goto L60
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v131
	F_errmsg(m, int32(713244), v11+int32(-48))
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L21
	} else {
		goto L54
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L21
	} else {
		goto L51
	}
L6:
	;
	if v22&int32(1) == int32(0) {
		goto L10
	} else {
		goto L11
	}
L7:
	;
	v22 = int32(1)
	goto L9
L8:
	;
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+76)))
	v22 = v21
	goto L9
L9:
	;
	goto L6
L10:
	;
	v28 = *(*int32)(unsafe.Add(mBase, _consts[84]))
	if v28 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	goto L12
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L21
	} else {
		goto L48
	}
L13:
	;
	v32 = *(*int32)(unsafe.Add(mBase, _consts[137]))
	v33 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+36)) = v33
	v36 = *(*int32)(unsafe.Add(mBase, _consts[150]))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)+4))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v32)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v37+v38<<(uint(int32(2))%32)))) = v33
	v458 = int64(1)
	goto L1
L14:
	;
	goto L15
L15:
	;
	v47 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1])))
	if v47 == int32(1) {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	if v57 != 0 {
		goto L5
	} else {
		goto L20
	}
L17:
	;
	v52 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)+316))
	v55 = base.B2i32(v53 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _consts[1])) = uint8(v55)
	v57 = v55
	goto L19
L18:
	;
	v57 = int32(0)
	goto L19
L19:
	;
	goto L16
L20:
	;
	v59 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	v63 = F_LWLockAcquire(m, v59+int32(384), int32(0))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	return int64(0)
L22:
	;
	v68 = *(*int32)(unsafe.Add(mBase, _consts[87]))
	v69 = *(*int64)(unsafe.Add(mBase, uint32(v68)+8))
	v70 = base.I32_wrap_i64(v69)
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v68)+20))
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v71))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v70)) == int32(0) {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	if v83 == int32(0) {
		v261 = v70
		v267 = v69
		goto L2
	} else {
		goto L27
	}
L24:
	;
	v83 = base.B2i32(base.Ui32(v71) <= base.Ui32(v70))
	goto L23
L25:
	;
	goto L26
L26:
	;
	v83 = base.B2i32(int32(0) <= v70-v71)
	goto L23
L27:
	;
	v87 = *(*int32)(unsafe.Add(mBase, _consts[87]))
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v87)+36))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v87)+32))
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v87)+28))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v87)+24))
	v93 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	F_LWLockRelease(m, v93+int32(384))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L21
	} else {
		goto L28
	}
L28:
	;
	v99 = int32(*(*uint8)(unsafe.Add(mBase, _consts[101])))
	if v70&int32(65535) != 0 {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v90))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v70)) == int32(0) {
		goto L37
	} else {
		goto L38
	}
L30:
	;
	if v99&int32(1) == int32(0) {
		goto L3
	} else {
		goto L35
	}
L31:
	;
	if v99&int32(1) == int32(0) {
		goto L30
	} else {
		goto L32
	}
L32:
	;
	F_SendPostmasterSignal(m, int32(4))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L21
	} else {
		goto L33
	}
L33:
	;
	v110 = int32(*(*uint8)(unsafe.Add(mBase, _consts[101])))
	if v110&int32(1) != 0 {
		goto L29
	} else {
		goto L34
	}
L34:
	;
	goto L3
L35:
	;
	goto L29
L36:
	;
	if v128 == int32(0) {
		goto L3
	} else {
		goto L40
	}
L37:
	;
	v128 = base.B2i32(base.Ui32(v90) <= base.Ui32(v70))
	goto L36
L38:
	;
	goto L39
L39:
	;
	v128 = base.B2i32(int32(0) <= v70-v90)
	goto L36
L40:
	;
	v131 = F_get_database_name(m, v88)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L21
	} else {
		goto L41
	}
L41:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L21
	} else {
		goto L42
	}
L42:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L21
	} else {
		goto L43
	}
L43:
	;
	if v131 != 0 {
		goto L4
	} else {
		goto L44
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v88
	F_errmsg(m, int32(56306), v13)
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L21
	} else {
		goto L45
	}
L45:
	;
	F_errhint(m, int32(584736), int32(0))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L21
	} else {
		goto L46
	}
L46:
	;
	F_errfinish(m, int32(495499), int32(166), int32(465190))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L21
	} else {
		goto L47
	}
L47:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L48:
	;
	F_errmsg_internal(m, int32(260415), int32(0))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L21
	} else {
		goto L49
	}
L49:
	;
	F_errfinish(m, int32(495499), int32(87), int32(465190))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L21
	} else {
		goto L50
	}
L50:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L51:
	;
	F_errmsg_internal(m, int32(14609), int32(0))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L21
	} else {
		goto L52
	}
L52:
	;
	F_errfinish(m, int32(495499), int32(103), int32(465190))
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L21
	} else {
		goto L53
	}
L53:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L54:
	;
	F_errhint(m, int32(584736), int32(0))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L21
	} else {
		goto L55
	}
L55:
	;
	F_errfinish(m, int32(495499), int32(159), int32(465190))
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L21
	} else {
		goto L56
	}
L56:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L57:
	;
	v251 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	v255 = F_LWLockAcquire(m, v251+int32(384), int32(0))
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L21
	} else {
		goto L75
	}
L58:
	;
	if v205 == int32(0) {
		goto L57
	} else {
		goto L62
	}
L59:
	;
	v205 = base.B2i32(base.Ui32(v91) <= base.Ui32(v70))
	goto L58
L60:
	;
	goto L61
L61:
	;
	v205 = base.B2i32(int32(0) <= v70-v91)
	goto L58
L62:
	;
	v208 = F_get_database_name(m, v88)
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L21
	} else {
		goto L63
	}
L63:
	;
	v212 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L21
	} else {
		goto L64
	}
L64:
	;
	if v208 != 0 {
		goto L66
	} else {
		goto L67
	}
L65:
	;
	F_errhint(m, v238, int32(0))
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L21
	} else {
		goto L73
	}
L66:
	;
	if v212 == int32(0) {
		goto L57
	} else {
		goto L69
	}
L67:
	;
	goto L68
L68:
	;
	if v212 == int32(0) {
		goto L57
	} else {
		goto L71
	}
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+48)) = v208
	*(*int32)(unsafe.Add(mBase, uint32(v13)+52)) = v89 - v70
	F_errmsg(m, int32(141953), v11+int32(-16))
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L21
	} else {
		goto L70
	}
L70:
	;
	v238 = int32(584539)
	v239 = int32(179)
	goto L65
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+32)) = v88
	*(*int32)(unsafe.Add(mBase, uint32(v13)+36)) = v89 - v70
	F_errmsg(m, int32(141892), v11+int32(-32))
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L21
	} else {
		goto L72
	}
L72:
	;
	v238 = int32(584353)
	v239 = int32(186)
	goto L65
L73:
	;
	F_errfinish(m, int32(495499), v239, int32(465190))
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L21
	} else {
		goto L74
	}
L74:
	;
	goto L57
L75:
	;
	v258 = *(*int32)(unsafe.Add(mBase, _consts[87]))
	v259 = *(*int64)(unsafe.Add(mBase, uint32(v258)+8))
	v261 = base.I32_wrap_i64(v259)
	v267 = v259
	goto L2
L76:
	;
	v277 = base.B2i32(v261 != int32(3))
	goto L78
L77:
	;
	v277 = int32(0)
	goto L78
L78:
	;
	if v277 == int32(0) {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v281 = *(*int32)(unsafe.Add(mBase, _consts[81]))
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v281)+28))
	v284 = int32(base.Ui32(v261) >> (uint(int32(15)) % 32))
	v286 = int32(*(*uint16)(unsafe.Add(mBase, _consts[83])))
	v287 = base.I32_rem_u_s(v284, v286)
	v290 = v282 + v287<<(uint(int32(7))%32)
	v292 = F_LWLockAcquire(m, v290, int32(0))
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L21
	} else {
		goto L82
	}
L80:
	;
	goto L81
L81:
	;
	v315 = int32(16)
	m.G0 = v270 + v315
	v318 = m.G0
	v320 = v318 - v315
	m.G0 = v320
	v323 = *(*int32)(unsafe.Add(mBase, _consts[85]))
	v324 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v323)+24)))
	if v324 != int32(1) {
		goto L88
	} else {
		goto L89
	}
L82:
	;
	v295 = base.I64_extend_i32_u(v284)
	v296 = F_SimpleLruZeroPage(m, int32(4409776), v295)
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L21
	} else {
		goto L83
	}
L83:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v270)+8)) = v295
	F_XLogBeginInsert(m)
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L21
	} else {
		goto L84
	}
L84:
	;
	v301 = int32(8)
	F_XLogRegisterData(m, v270+v301, v301)
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L21
	} else {
		goto L85
	}
L85:
	;
	v308 = F_XLogInsert(m, int32(3), int32(0))
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L21
	} else {
		goto L86
	}
L86:
	;
	F_LWLockRelease(m, v290)
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L21
	} else {
		goto L87
	}
L87:
	;
	goto L81
L88:
	;
	m.G0 = v320 + int32(16)
	F_ExtendSUBTRANS(m, v261)
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L21
	} else {
		goto L103
	}
L89:
	;
	v330 = int32(819)
	v331 = base.I32_div_u_s(v261, v330)
	if v261-v331*v330 != 0 {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v335 = base.B2i32(v261 != int32(3))
	goto L92
L91:
	;
	v335 = int32(0)
	goto L92
L92:
	;
	if v335 != 0 {
		goto L88
	} else {
		goto L93
	}
L93:
	;
	v337 = *(*int32)(unsafe.Add(mBase, _consts[86]))
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v337)+28))
	v340 = int32(*(*uint16)(unsafe.Add(mBase, _consts[88])))
	v341 = base.I32_rem_u_s(v331, v340)
	v344 = v338 + v341<<(uint(int32(7))%32)
	v346 = F_LWLockAcquire(m, v344, int32(0))
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L21
	} else {
		goto L94
	}
L94:
	;
	v349 = int32(*(*uint8)(unsafe.Add(mBase, _consts[82])))
	v351 = base.I64_extend_i32_u(v331)
	v352 = F_SimpleLruZeroPage(m, int32(4409860), v351)
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L21
	} else {
		goto L95
	}
L95:
	;
	if v349 == int32(0) {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v320)+8)) = v351
	F_XLogBeginInsert(m)
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L21
	} else {
		goto L99
	}
L97:
	;
	goto L98
L98:
	;
	F_LWLockRelease(m, v344)
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L21
	} else {
		goto L102
	}
L99:
	;
	v359 = int32(8)
	F_XLogRegisterData(m, v320+v359, v359)
	mBase = m.M
	v363 = m.ExcPending
	if v363 != 0 {
		goto L21
	} else {
		goto L100
	}
L100:
	;
	v366 = F_XLogInsert(m, int32(18), int32(0))
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L21
	} else {
		goto L101
	}
L101:
	;
	goto L98
L102:
	;
	goto L88
L103:
	;
	v380 = *(*int32)(unsafe.Add(mBase, _consts[87]))
	v381 = *(*int64)(unsafe.Add(mBase, uint32(v380)+8))
	v382 = int64(1)
	v390 = v381 + v382
	if base.Ui32(base.I32_wrap_i64(v390)) < base.Ui32(int32(3)) {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	v394 = v381 + (v382-v381)&int64(4294967295) + int64(2)
	goto L106
L105:
	;
	v394 = v390
	goto L106
L106:
	;
	if base.Ui64(int64(2)) < base.Ui64(v390) {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	v397 = v394
	goto L109
L108:
	;
	v397 = v390
	goto L109
L109:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v380)+8)) = v397
	if l0 == int32(0) {
		goto L111
	} else {
		goto L112
	}
L110:
	;
	v445 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	F_LWLockRelease(m, v445+int32(384))
	mBase = m.M
	v449 = m.ExcPending
	if v449 != 0 {
		goto L21
	} else {
		goto L117
	}
L111:
	;
	v402 = *(*int32)(unsafe.Add(mBase, _consts[137]))
	*(*int32)(unsafe.Add(mBase, uint32(v402)+36)) = v261
	v405 = *(*int32)(unsafe.Add(mBase, _consts[150]))
	v406 = *(*int32)(unsafe.Add(mBase, uint32(v405)+4))
	v407 = *(*int32)(unsafe.Add(mBase, uint32(v402)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v406+v407<<(uint(int32(2))%32)))) = v261
	goto L110
L112:
	;
	goto L113
L113:
	;
	v413 = *(*int32)(unsafe.Add(mBase, _consts[150]))
	v414 = *(*int32)(unsafe.Add(mBase, uint32(v413)+8))
	v416 = *(*int32)(unsafe.Add(mBase, _consts[137]))
	v417 = *(*int32)(unsafe.Add(mBase, uint32(v416)+48))
	v420 = v414 + v417<<(uint(int32(1))%32)
	v421 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v416)+276)))
	if base.Ui32(v421) <= base.Ui32(int32(63)) {
		goto L114
	} else {
		goto L115
	}
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v416+v421<<(uint(int32(2))%32))+280)) = v261
	v429 = v421 + int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v420))) = uint8(v429)
	v432 = *(*int32)(unsafe.Add(mBase, _consts[137]))
	*(*uint8)(unsafe.Add(mBase, uint32(v432)+276)) = uint8(v429)
	goto L110
L115:
	;
	goto L116
L116:
	;
	v434 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v420)+1)) = uint8(v434)
	v437 = *(*int32)(unsafe.Add(mBase, _consts[137]))
	*(*uint8)(unsafe.Add(mBase, uint32(v437)+277)) = uint8(v434)
	goto L110
L117:
	;
	v458 = v267
	goto L1
}
func F_TransactionIdFollowsOrEquals(m *base.Module, l0 int32, l1 int32) int32 {
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(l1))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(l0)) == int32(0) {
		return base.B2i32(base.Ui32(l1) <= base.Ui32(l0))
	} else {
		return base.B2i32(int32(0) <= l0-l1)
	}
}
func F_TransactionIdGetStatus(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v36 int64
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	v10 = int32(base.Ui32(l0) >> (uint(int32(15)) % 32))
	v12 = F_SimpleLruReadPage_ReadOnly(m, int32(4409776), base.I64_extend_i32_u(v10), l0)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v18 = int32(2)
		v19 = int32(base.Ui32(l0&int32(32764)) >> (uint(v18) % 32))
		v20 = int32(4409776)
		v21 = *(*int32)(unsafe.Add(mBase, _consts[81]))
		v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
		v26 = *(*int32)(unsafe.Add(mBase, uint32(v22+v12<<(uint(v18)%32))))
		v28 = int32(*(*int8)(unsafe.Add(mBase, uint32(v19+v26))))
		v29 = *(*int32)(unsafe.Add(mBase, uint32(v21)+36))
		v36 = *(*int64)(unsafe.Add(mBase, uint32(v29+(v19&int32(8184)|v12<<(uint(int32(13))%32)))))
		*(*int64)(unsafe.Add(mBase, uint32(l1))) = v36
		v39 = *(*int32)(unsafe.Add(mBase, _consts[81]))
		v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+28))
		v42 = int32(*(*uint16)(unsafe.Add(mBase, _consts[83])))
		v43 = base.I32_rem_u_s(v10, v42)
		F_LWLockRelease(m, v40+v43<<(uint(int32(7))%32))
		mBase = m.M
		v48 = m.ExcPending
		if v48 != 0 {
			return int32(0)
		} else {
			return v28 >> (uint(l0<<(uint(int32(1))%32)&int32(6)) % 32) & int32(3)
		}
	}
}
func F_TransactionIdIsCurrentTransactionId(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v126 int32
	_ = v126
	if base.Ui32(l0) < base.Ui32(int32(3)) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	goto L3
L3:
	;
	v13 = *(*int32)(unsafe.Add(mBase, _consts[80]))
	if v13 == l0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return int32(1)
L5:
	;
	goto L6
L6:
	;
	v18 = *(*int32)(unsafe.Add(mBase, _consts[145]))
	if v18 <= int32(0) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	return v126
L8:
	;
	v22 = *(*int32)(unsafe.Add(mBase, _consts[72]))
	if v22 == int32(0) {
		v126 = int32(0)
		goto L7
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	v95 = *(*int32)(unsafe.Add(mBase, _consts[146]))
	v97 = int32(0)
	v99 = v18 - int32(1)
	goto L34
L11:
	;
	v27 = v22
	goto L12
L12:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v27)+20))
	if v32 == int32(4) {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	v126 = int32(0)
	goto L7
L14:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v27)+80))
	if v90 != 0 {
		v27 = v90
		goto L12
	} else {
		goto L33
	}
L15:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
	if v35 == int32(0) {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v38 = int32(1)
	if l0 == v35 {
		v126 = v38
		goto L7
	} else {
		goto L17
	}
L17:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v27)+52))
	v42 = v40 - int32(1)
	if v42 < int32(0) {
		goto L14
	} else {
		goto L18
	}
L18:
	;
	v47 = int32(0)
	v49 = v42
	goto L19
L19:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v27)+48))
	v55 = int32(2)
	v56 = base.I32_div_s(v49-v47, v55)
	v57 = v56 + v47
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v53+v57<<(uint(v55)%32))))
	if v61 == l0 {
		v126 = v38
		goto L7
	} else {
		goto L21
	}
L20:
	;
	goto L14
L21:
	;
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(l0))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v61)) == int32(0) {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	if v76 != 0 {
		goto L26
	} else {
		goto L27
	}
L23:
	;
	v76 = base.B2i32(base.Ui32(v61) < base.Ui32(l0))
	goto L22
L24:
	;
	goto L25
L25:
	;
	v76 = int32(base.Ui32(v61-l0) >> (uint(int32(31)) % 32))
	goto L22
L26:
	;
	v77 = v57 + int32(1)
	goto L28
L27:
	;
	v77 = v47
	goto L28
L28:
	;
	if v76 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v80 = v49
	goto L31
L30:
	;
	v80 = v57 - int32(1)
	goto L31
L31:
	;
	if v77 <= v80 {
		v47 = v77
		v49 = v80
		goto L19
	} else {
		goto L32
	}
L32:
	;
	goto L20
L33:
	;
	goto L13
L34:
	;
	v104 = int32(2)
	v105 = base.I32_div_s(v99-v97, v104)
	v106 = v105 + v97
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v95+v106<<(uint(v104)%32))))
	v111 = base.B2i32(v110 == l0)
	if v110 == l0 {
		v126 = v111
		goto L7
	} else {
		goto L36
	}
L35:
	;
	v126 = v111
	goto L7
L36:
	;
	v114 = base.B2i32(base.Ui32(v110) < base.Ui32(l0))
	if base.Ui32(v110) < base.Ui32(l0) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v115 = v106 + int32(1)
	goto L39
L38:
	;
	v115 = v97
	goto L39
L39:
	;
	if base.Ui32(v110) < base.Ui32(l0) {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v118 = v99
	goto L42
L41:
	;
	v118 = v106 - int32(1)
	goto L42
L42:
	;
	if v115 <= v118 {
		v97 = v115
		v99 = v118
		goto L34
	} else {
		goto L43
	}
L43:
	;
	goto L35
}
func F_TransactionIdSetPageStatusInternal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int64, l5 int64) {
	mBase := m.M
	_ = mBase
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v35 int32
	_ = v35
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
	var v85 int64
	_ = v85
	var v90 int32
	_ = v90
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v154 int32
	_ = v154
	var v155 int64
	_ = v155
	var v184 int32
	_ = v184
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v241 int32
	_ = v241
	var v242 int64
	_ = v242
	var v247 int32
	_ = v247
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	v17 = base.B2i32(l4 == int64(0))
	v18 = F_SimpleLruReadPage(m, int32(4409776), l5, v17, l0)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return
	} else {
		if l0 == int32(0) {
		} else {
			if l3 != int32(1) {
			} else {
				if l1 <= int32(0) {
				} else {
					v35 = int32(0)
					for {
						v43 = *(*int32)(unsafe.Add(mBase, _consts[81]))
						v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
						v45 = int32(2)
						v48 = *(*int32)(unsafe.Add(mBase, uint32(v44+v18<<(uint(v45)%32))))
						v52 = *(*int32)(unsafe.Add(mBase, uint32(l2+v35<<(uint(v45)%32))))
						v54 = v52 & int32(32767)
						v57 = v48 + int32(base.Ui32(v54)>>(uint(v45)%32))
						v58 = int32(*(*int8)(unsafe.Add(mBase, uint32(v57))))
						v59 = int32(1)
						v62 = v52 << (uint(v59) % 32) & int32(6)
						v64 = int32(*(*uint8)(unsafe.Add(mBase, _consts[82])))
						if v64 == v59 {
							if v58>>(uint(v62)%32)&int32(3) == int32(1) {
							} else {
								v74 = v58 | int32(3)<<(uint(v62)%32)
								*(*uint8)(unsafe.Add(mBase, uint32(v57))) = uint8(v74)
								if l4 == int64(0) {
								} else {
									v77 = *(*int32)(unsafe.Add(mBase, _consts[81]))
									v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)+36))
									v84 = v78 + (int32(base.Ui32(v54)>>(uint(int32(5))%32))|v18<<(uint(int32(10))%32))<<(uint(int32(3))%32)
									v85 = *(*int64)(unsafe.Add(mBase, uint32(v84)))
									if base.Ui64(l4) <= base.Ui64(v85) {
									} else {
										*(*int64)(unsafe.Add(mBase, uint32(v84))) = l4
									}
								}
							}
						} else {
							v74 = v58 | int32(3)<<(uint(v62)%32)
							*(*uint8)(unsafe.Add(mBase, uint32(v57))) = uint8(v74)
							if l4 == int64(0) {
							} else {
								v77 = *(*int32)(unsafe.Add(mBase, _consts[81]))
								v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)+36))
								v84 = v78 + (int32(base.Ui32(v54)>>(uint(int32(5))%32))|v18<<(uint(int32(10))%32))<<(uint(int32(3))%32)
								v85 = *(*int64)(unsafe.Add(mBase, uint32(v84)))
								if base.Ui64(l4) <= base.Ui64(v85) {
								} else {
									*(*int64)(unsafe.Add(mBase, uint32(v84))) = l4
								}
							}
						}
						v90 = v35 + int32(1)
						if v90 != l1 {
							v35 = v90
							continue
						} else {
							break
						}
						break
					}
				}
			}
			v107 = *(*int32)(unsafe.Add(mBase, _consts[81]))
			v108 = *(*int32)(unsafe.Add(mBase, uint32(v107)+4))
			v109 = int32(2)
			v112 = *(*int32)(unsafe.Add(mBase, uint32(v108+v18<<(uint(v109)%32))))
			v114 = l0 & int32(32767)
			v117 = v112 + int32(base.Ui32(v114)>>(uint(v109)%32))
			v118 = int32(*(*int8)(unsafe.Add(mBase, uint32(v117))))
			v122 = l0 << (uint(int32(1)) % 32) & int32(6)
			if l3 != int32(3) {
				v142 = v118&(int32(3)<<(uint(v122)%32)^int32(-1)) | l3<<(uint(v122)%32)
				*(*uint8)(unsafe.Add(mBase, uint32(v117))) = uint8(v142)
				if l4 == int64(0) {
				} else {
					v145 = *(*int32)(unsafe.Add(mBase, _consts[81]))
					v146 = *(*int32)(unsafe.Add(mBase, uint32(v145)+36))
					v154 = v146 + (int32(base.Ui32(v114)>>(uint(int32(2))%32))&int32(8184) | v18<<(uint(int32(13))%32))
					v155 = *(*int64)(unsafe.Add(mBase, uint32(v154)))
					if base.Ui64(l4) <= base.Ui64(v155) {
					} else {
						*(*int64)(unsafe.Add(mBase, uint32(v154))) = l4
					}
				}
			} else {
				v126 = int32(*(*uint8)(unsafe.Add(mBase, _consts[82])))
				if v126&int32(1) == int32(0) {
					v142 = v118&(int32(3)<<(uint(v122)%32)^int32(-1)) | l3<<(uint(v122)%32)
					*(*uint8)(unsafe.Add(mBase, uint32(v117))) = uint8(v142)
					if l4 == int64(0) {
					} else {
						v145 = *(*int32)(unsafe.Add(mBase, _consts[81]))
						v146 = *(*int32)(unsafe.Add(mBase, uint32(v145)+36))
						v154 = v146 + (int32(base.Ui32(v114)>>(uint(int32(2))%32))&int32(8184) | v18<<(uint(int32(13))%32))
						v155 = *(*int64)(unsafe.Add(mBase, uint32(v154)))
						if base.Ui64(l4) <= base.Ui64(v155) {
						} else {
							*(*int64)(unsafe.Add(mBase, uint32(v154))) = l4
						}
					}
				} else {
					if v118>>(uint(v122)%32)&int32(3) == int32(1) {
					} else {
						v142 = v118&(int32(3)<<(uint(v122)%32)^int32(-1)) | l3<<(uint(v122)%32)
						*(*uint8)(unsafe.Add(mBase, uint32(v117))) = uint8(v142)
						if l4 == int64(0) {
						} else {
							v145 = *(*int32)(unsafe.Add(mBase, _consts[81]))
							v146 = *(*int32)(unsafe.Add(mBase, uint32(v145)+36))
							v154 = v146 + (int32(base.Ui32(v114)>>(uint(int32(2))%32))&int32(8184) | v18<<(uint(int32(13))%32))
							v155 = *(*int64)(unsafe.Add(mBase, uint32(v154)))
							if base.Ui64(l4) <= base.Ui64(v155) {
							} else {
								*(*int64)(unsafe.Add(mBase, uint32(v154))) = l4
							}
						}
					}
				}
			}
		}
		if int32(0) < l1 {
			v184 = int32(0)
			for {
				v192 = *(*int32)(unsafe.Add(mBase, _consts[81]))
				v193 = *(*int32)(unsafe.Add(mBase, uint32(v192)+4))
				v194 = int32(2)
				v197 = *(*int32)(unsafe.Add(mBase, uint32(v193+v18<<(uint(v194)%32))))
				v201 = *(*int32)(unsafe.Add(mBase, uint32(l2+v184<<(uint(v194)%32))))
				v203 = v201 & int32(32767)
				v206 = v197 + int32(base.Ui32(v203)>>(uint(v194)%32))
				v207 = int32(*(*int8)(unsafe.Add(mBase, uint32(v206))))
				v211 = v201 << (uint(int32(1)) % 32) & int32(6)
				if l3 != int32(3) {
					v231 = v207&(int32(3)<<(uint(v211)%32)^int32(-1)) | l3<<(uint(v211)%32)
					*(*uint8)(unsafe.Add(mBase, uint32(v206))) = uint8(v231)
					if l4 == int64(0) {
					} else {
						v234 = *(*int32)(unsafe.Add(mBase, _consts[81]))
						v235 = *(*int32)(unsafe.Add(mBase, uint32(v234)+36))
						v241 = v235 + (int32(base.Ui32(v203)>>(uint(int32(5))%32))|v18<<(uint(int32(10))%32))<<(uint(int32(3))%32)
						v242 = *(*int64)(unsafe.Add(mBase, uint32(v241)))
						if base.Ui64(l4) <= base.Ui64(v242) {
						} else {
							*(*int64)(unsafe.Add(mBase, uint32(v241))) = l4
						}
					}
				} else {
					v215 = int32(*(*uint8)(unsafe.Add(mBase, _consts[82])))
					if v215&int32(1) == int32(0) {
						v231 = v207&(int32(3)<<(uint(v211)%32)^int32(-1)) | l3<<(uint(v211)%32)
						*(*uint8)(unsafe.Add(mBase, uint32(v206))) = uint8(v231)
						if l4 == int64(0) {
						} else {
							v234 = *(*int32)(unsafe.Add(mBase, _consts[81]))
							v235 = *(*int32)(unsafe.Add(mBase, uint32(v234)+36))
							v241 = v235 + (int32(base.Ui32(v203)>>(uint(int32(5))%32))|v18<<(uint(int32(10))%32))<<(uint(int32(3))%32)
							v242 = *(*int64)(unsafe.Add(mBase, uint32(v241)))
							if base.Ui64(l4) <= base.Ui64(v242) {
							} else {
								*(*int64)(unsafe.Add(mBase, uint32(v241))) = l4
							}
						}
					} else {
						if v207>>(uint(v211)%32)&int32(3) == int32(1) {
						} else {
							v231 = v207&(int32(3)<<(uint(v211)%32)^int32(-1)) | l3<<(uint(v211)%32)
							*(*uint8)(unsafe.Add(mBase, uint32(v206))) = uint8(v231)
							if l4 == int64(0) {
							} else {
								v234 = *(*int32)(unsafe.Add(mBase, _consts[81]))
								v235 = *(*int32)(unsafe.Add(mBase, uint32(v234)+36))
								v241 = v235 + (int32(base.Ui32(v203)>>(uint(int32(5))%32))|v18<<(uint(int32(10))%32))<<(uint(int32(3))%32)
								v242 = *(*int64)(unsafe.Add(mBase, uint32(v241)))
								if base.Ui64(l4) <= base.Ui64(v242) {
								} else {
									*(*int64)(unsafe.Add(mBase, uint32(v241))) = l4
								}
							}
						}
					}
				}
				v247 = v184 + int32(1)
				if v247 != l1 {
					v184 = v247
					continue
				} else {
					break
				}
				break
			}
		} else {
		}
		v264 = *(*int32)(unsafe.Add(mBase, _consts[81]))
		v265 = *(*int32)(unsafe.Add(mBase, uint32(v264)+12))
		v267 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(v265+v18))) = uint8(v267)
		return
	}
}
func F_TransactionTimeoutHandler(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	v2 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[1124])) = v2
	*(*int32)(unsafe.Add(mBase, _consts[8])) = v2
	v8 = *(*int32)(unsafe.Add(mBase, _consts[311]))
	F_SetLatch(m, v8)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return
	} else {
		return
	}
}
func F_UnlockApplyTransactionForSession(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
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
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	v3 = l2
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = int32(267)
	*(*uint16)(unsafe.Add(mBase, uint32(v8)+14)) = uint16(v10)
	*(*uint16)(unsafe.Add(mBase, uint32(v8)+12)) = uint16(v3)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = l0
	v16 = *(*int32)(unsafe.Add(mBase, _consts[128]))
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v16
	v19 = F_LockRelease(m, v8, l3, int32(1))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		return
	} else {
		m.G0 = v8 + int32(16)
		return
	}
}
func F_check_transaction_read_only(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	v4 = int32(1)
	v5 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v5 != 0 {
		v63 = v4
		return v63
	} else {
		v7 = int32(*(*uint8)(unsafe.Add(mBase, _consts[181])))
		if v7 != int32(1) {
			v63 = v4
			return v63
		} else {
			v11 = *(*int32)(unsafe.Add(mBase, _consts[72]))
			v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+20))
			if base.B2i32(v12 == int32(2)) == int32(0) {
				v63 = v4
				return v63
			} else {
				v18 = int32(*(*uint8)(unsafe.Add(mBase, _consts[526])))
				if v18 != 0 {
					v63 = v4
					return v63
				} else {
					v19 = int32(16777538)
					v22 = *(*int32)(unsafe.Add(mBase, _consts[72]))
					v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+28))
					if int32(1) < v23 {
						v46 = v19
						v47 = int32(256478)
						*(*int32)(unsafe.Add(mBase, _consts[527])) = v46
						v52 = *(*int32)(unsafe.Add(mBase, _consts[40]))
						*(*int32)(unsafe.Add(mBase, _consts[506])) = v52
						v57 = F_format_elog_string(m, v47, int32(0))
						mBase = m.M
						v60 = m.ExcPending
						if v60 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, _consts[525])) = v57
							v63 = int32(0)
							return v63
						}
					} else {
						v28 = int32(*(*uint8)(unsafe.Add(mBase, _consts[528])))
						if v28 != 0 {
							v46 = v19
							v47 = int32(16035)
							*(*int32)(unsafe.Add(mBase, _consts[527])) = v46
							v52 = *(*int32)(unsafe.Add(mBase, _consts[40]))
							*(*int32)(unsafe.Add(mBase, _consts[506])) = v52
							v57 = F_format_elog_string(m, v47, int32(0))
							mBase = m.M
							v60 = m.ExcPending
							if v60 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, _consts[525])) = v57
								v63 = int32(0)
								return v63
							}
						} else {
							v31 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1])))
							if v31 == int32(1) {
								v36 = *(*int32)(unsafe.Add(mBase, _consts[2]))
								v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)+316))
								v39 = base.B2i32(v37 != int32(2))
								*(*uint8)(unsafe.Add(mBase, _consts[1])) = uint8(v39)
								v41 = v39
							} else {
								v41 = int32(0)
							}
							if v41 == int32(0) {
								v63 = v4
								return v63
							} else {
								v46 = int32(1088)
								v47 = int32(14723)
								*(*int32)(unsafe.Add(mBase, _consts[527])) = v46
								v52 = *(*int32)(unsafe.Add(mBase, _consts[40]))
								*(*int32)(unsafe.Add(mBase, _consts[506])) = v52
								v57 = F_format_elog_string(m, v47, int32(0))
								mBase = m.M
								v60 = m.ExcPending
								if v60 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, _consts[525])) = v57
									v63 = int32(0)
									return v63
								}
							}
						}
					}
				}
			}
		}
	}
}
