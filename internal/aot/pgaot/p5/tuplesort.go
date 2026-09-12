package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_tuplesort_estimate_shared(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	v3 = F_mul_size(m, int32(8), l0)
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v8 = F_add_size(m, v3, int32(72))
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			return (v8 + int32(7)) & int32(-8)
		}
	}
}
func F_tuplesort_getgintuple(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	v3 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = int32(4515248)
	v11 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v13
	v16 = F_tuplesort_gettuple_common(m, l0, int32(1), v8)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return int32(0)
	} else {
		if v16 == int32(0) {
			*(*int32)(unsafe.Add(mBase, _consts[9])) = v11
			v32 = v3
		} else {
			*(*int32)(unsafe.Add(mBase, _consts[9])) = v11
			v26 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
			if v26 == int32(0) {
				v32 = v3
			} else {
				v29 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
				*(*int32)(unsafe.Add(mBase, uint32(l1))) = v29
				v32 = v26
			}
		}
		m.G0 = v8 + int32(16)
		return v32
	}
}
func F_tuplesort_getheaptuple(m *base.Module, l0 int32) int32 {
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
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = int32(4515248)
	v9 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v11
	v14 = F_tuplesort_gettuple_common(m, l0, int32(1), v6)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, _consts[9])) = v9
		v20 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
		m.G0 = v6 + int32(16)
		if v14 != 0 {
			v25 = v20
		} else {
			v25 = int32(0)
		}
		return v25
	}
}
func F_tuplesort_gettuple_common(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int64
	_ = v31
	var v33 int64
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v78 int64
	_ = v78
	var v80 int64
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v222 int64
	_ = v222
	var v224 int64
	_ = v224
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v244 int32
	_ = v244
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v296 int32
	_ = v296
	var v297 int64
	_ = v297
	var v299 int64
	_ = v299
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v308 int32
	_ = v308
	var v320 int32
	_ = v320
	var v321 int64
	_ = v321
	var v323 int64
	_ = v323
	var v337 int32
	_ = v337
	var v342 int32
	_ = v342
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v351 int32
	_ = v351
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v366 int32
	_ = v366
	var v374 int32
	_ = v374
	var v376 int32
	_ = v376
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v399 int32
	_ = v399
	var v400 int64
	_ = v400
	var v402 int64
	_ = v402
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v407 int32
	_ = v407
	var v410 int32
	_ = v410
	var v423 int32
	_ = v423
	var v424 int64
	_ = v424
	var v426 int64
	_ = v426
	var v432 int32
	_ = v432
	var v436 int32
	_ = v436
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v445 int32
	_ = v445
	var v463 int32
	_ = v463
	var v467 int32
	_ = v467
	var v472 int32
	_ = v472
	var v476 int32
	_ = v476
	var v480 int32
	_ = v480
	var v485 int32
	_ = v485
	var v489 int32
	_ = v489
	var v493 int32
	_ = v493
	var v498 int32
	_ = v498
	v4 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(32)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	switch v17 - int32(3) {
	case 0:
		goto L9
	case 1:
		goto L8
	case 2:
		goto L7
	default:
		goto L6
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v489 = m.ExcPending
	if v489 != 0 {
		goto L18
	} else {
		goto L145
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v476 = m.ExcPending
	if v476 != 0 {
		goto L18
	} else {
		goto L142
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v463 = m.ExcPending
	if v463 != 0 {
		goto L18
	} else {
		goto L139
	}
L4:
	;
	m.G0 = v15 + int32(32)
	return v445
L5:
	;
	v442 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+208)) = uint8(v442)
	v445 = v96
	goto L4
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v432 = m.ExcPending
	if v432 != 0 {
		goto L18
	} else {
		goto L136
	}
L7:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
	if v199 != 0 {
		goto L71
	} else {
		goto L72
	}
L8:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
	if v83 != 0 {
		goto L28
	} else {
		goto L29
	}
L9:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
	if l1 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	if v20 < v21 {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	goto L12
L12:
	;
	v58 = int32(0)
	if v20 <= v58 {
		v445 = v58
		goto L4
	} else {
		goto L22
	}
L13:
	;
	v23 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v20 + v23
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v30 = v27 + v20<<(uint(int32(4))%32)
	v31 = *(*int64)(unsafe.Add(mBase, uint32(v30)))
	*(*int64)(unsafe.Add(mBase, uint32(l2))) = v31
	v33 = *(*int64)(unsafe.Add(mBase, uint32(v30)+8))
	*(*int64)(unsafe.Add(mBase, uint32(l2)+8)) = v33
	v445 = v23
	goto L4
L14:
	;
	goto L15
L15:
	;
	v35 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+208)) = uint8(v35)
	v37 = int32(0)
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+68)))
	if v38 != v35 {
		v445 = v37
		goto L4
	} else {
		goto L16
	}
L16:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	if v20 < v41 {
		v445 = v37
		goto L4
	} else {
		goto L17
	}
L17:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	return int32(0)
L19:
	;
	F_errmsg_internal(m, int32(78765), int32(0))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L18
	} else {
		goto L20
	}
L20:
	;
	F_errfinish(m, int32(492667), int32(1498), int32(245806))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L18
	} else {
		goto L21
	}
L21:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L22:
	;
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+208)))
	if v61 == int32(1) {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v77 = v72 + v71<<(uint(int32(4))%32) - int32(16)
	v78 = *(*int64)(unsafe.Add(mBase, uint32(v77)))
	*(*int64)(unsafe.Add(mBase, uint32(l2))) = v78
	v80 = *(*int64)(unsafe.Add(mBase, uint32(v77)+8))
	*(*int64)(unsafe.Add(mBase, uint32(l2)+8)) = v80
	v445 = int32(1)
	goto L4
L24:
	;
	v64 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+208)) = uint8(v64)
	v71 = v20
	goto L23
L25:
	;
	goto L26
L26:
	;
	v66 = int32(1)
	v67 = v20 - v66
	*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v67
	if v20 == v66 {
		v445 = v58
		goto L4
	} else {
		goto L27
	}
L27:
	;
	v71 = v67
	goto L23
L28:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+148))
	if base.Ui32(v83) < base.Ui32(v84) {
		goto L32
	} else {
		goto L33
	}
L29:
	;
	goto L30
L30:
	;
	v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+208)))
	if l1 != 0 {
		goto L36
	} else {
		goto L37
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+164)) = int32(0)
	goto L30
L32:
	;
	F_pfree(m, v83)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L18
	} else {
		goto L35
	}
L33:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	if base.Ui32(v86) <= base.Ui32(v83) {
		goto L32
	} else {
		goto L34
	}
L34:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	*(*int32)(unsafe.Add(mBase, uint32(v83))) = v88
	*(*int32)(unsafe.Add(mBase, uint32(l0)+156)) = v83
	goto L31
L35:
	;
	goto L31
L36:
	;
	v96 = int32(0)
	if v95&int32(1) != 0 {
		v445 = v96
		goto L4
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	if v95&int32(1) != 0 {
		goto L46
	} else {
		goto L47
	}
L39:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	v103 = F_LogicalTapeRead(m, v99, v15+int32(12), int32(4))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L18
	} else {
		goto L40
	}
L40:
	;
	if v103 != int32(4) {
		goto L3
	} else {
		goto L41
	}
L41:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	if v107 == int32(0) {
		goto L5
	} else {
		goto L42
	}
L42:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	m.T0[v111].(func(*base.Module, int32, int32, int32, int32))(m, l0, l2, v110, v107)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L18
	} else {
		goto L43
	}
L43:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+164)) = v114
	v445 = int32(1)
	goto L4
L44:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	v186 = F_getlen(m, v185)
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L18
	} else {
		goto L67
	}
L45:
	;
	v179 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+208)) = uint8(v179)
	goto L44
L46:
	;
	v122 = F_LogicalTapeBackspace(m, v117, int32(8))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L18
	} else {
		goto L50
	}
L47:
	;
	goto L48
L48:
	;
	v137 = int32(0)
	v139 = F_LogicalTapeBackspace(m, v117, int32(4))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L18
	} else {
		goto L56
	}
L49:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L18
	} else {
		goto L51
	}
L50:
	;
	switch v122 {
	case 0:
		v445 = int32(0)
		goto L4
	default:
		goto L49
	case 8:
		goto L45
	}
L51:
	;
	F_errmsg_internal(m, int32(249534), int32(0))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L18
	} else {
		goto L52
	}
L52:
	;
	F_errfinish(m, int32(492667), int32(1581), int32(245806))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L18
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
	v154 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	v155 = F_getlen(m, v154)
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L18
	} else {
		goto L60
	}
L55:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L18
	} else {
		goto L57
	}
L56:
	;
	switch v139 {
	case 0:
		v445 = v137
		goto L4
	default:
		goto L55
	case 4:
		goto L54
	}
L57:
	;
	F_errmsg_internal(m, int32(249534), int32(0))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L18
	} else {
		goto L58
	}
L58:
	;
	F_errfinish(m, int32(492667), int32(1595), int32(245806))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L18
	} else {
		goto L59
	}
L59:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L60:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	v159 = v155 + int32(8)
	v160 = F_LogicalTapeBackspace(m, v157, v159)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L18
	} else {
		goto L61
	}
L61:
	;
	if v160 == v155+int32(4) {
		v445 = v137
		goto L4
	} else {
		goto L62
	}
L62:
	;
	if v159 == v160 {
		goto L44
	} else {
		goto L63
	}
L63:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L18
	} else {
		goto L64
	}
L64:
	;
	F_errmsg_internal(m, int32(284919), int32(0))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L18
	} else {
		goto L65
	}
L65:
	;
	F_errfinish(m, int32(492667), int32(1615), int32(245806))
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L18
	} else {
		goto L66
	}
L66:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L67:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	v189 = F_LogicalTapeBackspace(m, v188, v186)
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L18
	} else {
		goto L68
	}
L68:
	;
	if v189 != v186 {
		goto L2
	} else {
		goto L69
	}
L69:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	v193 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	m.T0[v193].(func(*base.Module, int32, int32, int32, int32))(m, l0, l2, v192, v186)
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L18
	} else {
		goto L70
	}
L70:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+164)) = v196
	v445 = int32(1)
	goto L4
L71:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(l0)+148))
	if base.Ui32(v199) < base.Ui32(v200) {
		goto L75
	} else {
		goto L76
	}
L72:
	;
	goto L73
L73:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	if v211 <= int32(0) {
		goto L79
	} else {
		goto L80
	}
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+164)) = int32(0)
	goto L73
L75:
	;
	F_pfree(m, v199)
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L18
	} else {
		goto L78
	}
L76:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	if base.Ui32(v202) <= base.Ui32(v199) {
		goto L75
	} else {
		goto L77
	}
L77:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	*(*int32)(unsafe.Add(mBase, uint32(v199))) = v204
	*(*int32)(unsafe.Add(mBase, uint32(l0)+156)) = v199
	goto L74
L78:
	;
	goto L74
L79:
	;
	v445 = int32(0)
	goto L4
L80:
	;
	goto L81
L81:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(l0)+172))
	v216 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v216)+12))
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v215+v217<<(uint(int32(2))%32))))
	v222 = *(*int64)(unsafe.Add(mBase, uint32(v216)+8))
	*(*int64)(unsafe.Add(mBase, uint32(l2)+8)) = v222
	v224 = *(*int64)(unsafe.Add(mBase, uint32(v216)))
	*(*int64)(unsafe.Add(mBase, uint32(l2))) = v224
	*(*uint32)(unsafe.Add(mBase, uint32(l0)+164)) = uint32(v224)
	v230 = F_LogicalTapeRead(m, v221, v15+int32(28), int32(4))
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L18
	} else {
		goto L82
	}
L82:
	;
	if v230 != int32(4) {
		goto L1
	} else {
		goto L83
	}
L83:
	;
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v15)+28))
	if v234 == int32(0) {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v238 = int32(1)
	v239 = v237 - v238
	*(*int32)(unsafe.Add(mBase, uint32(l0)+136)) = v239
	if int32(0) < v239 {
		goto L87
	} else {
		goto L88
	}
L85:
	;
	goto L86
L86:
	;
	v345 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	m.T0[v345].(func(*base.Module, int32, int32, int32, int32))(m, l0, v15+int32(12), v221, v234)
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L18
	} else {
		goto L113
	}
L87:
	;
	v244 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v248 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	if v248 != 0 {
		goto L90
	} else {
		goto L91
	}
L88:
	;
	goto L89
L89:
	;
	v337 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+180)) = v337 - int32(1)
	F_LogicalTapeClose(m, v221)
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L18
	} else {
		goto L112
	}
L90:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L18
	} else {
		goto L93
	}
L91:
	;
	v252 = v239
	goto L92
L92:
	;
	v253 = v239<<(uint(int32(4))%32) + v244
	if base.Ui32(v252) < base.Ui32(int32(2)) {
		goto L95
	} else {
		goto L96
	}
L93:
	;
	v251 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v252 = v251
	goto L92
L94:
	;
	v320 = v244 + v308<<(uint(int32(4))%32)
	v321 = *(*int64)(unsafe.Add(mBase, uint32(v253)))
	*(*int64)(unsafe.Add(mBase, uint32(v320))) = v321
	v323 = *(*int64)(unsafe.Add(mBase, uint32(v253)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v320)+8)) = v323
	goto L89
L95:
	;
	v308 = int32(0)
	goto L94
L96:
	;
	goto L97
L97:
	;
	v262 = int32(1)
	v265 = v4
	v268 = v4
	goto L98
L98:
	;
	v271 = v265 + int32(2)
	if base.Ui32(v252) <= base.Ui32(v271) {
		goto L100
	} else {
		goto L101
	}
L99:
	;
	v308 = v285
	goto L94
L100:
	;
	v285 = v262
	goto L102
L101:
	;
	v273 = int32(4)
	v279 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v280 = m.T0[v279].(func(*base.Module, int32, int32, int32) int32)(m, v244+v262<<(uint(v273)%32), v244+v271<<(uint(v273)%32), l0)
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L18
	} else {
		goto L103
	}
L102:
	;
	v288 = v244 + v285<<(uint(int32(4))%32)
	v289 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v290 = m.T0[v289].(func(*base.Module, int32, int32, int32) int32)(m, v253, v288, l0)
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L18
	} else {
		goto L107
	}
L103:
	;
	if int32(0) < v280 {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	v284 = v271
	goto L106
L105:
	;
	v284 = v262
	goto L106
L106:
	;
	v285 = v284
	goto L102
L107:
	;
	if v290 <= int32(0) {
		goto L108
	} else {
		goto L109
	}
L108:
	;
	v308 = v268
	goto L94
L109:
	;
	goto L110
L110:
	;
	v296 = v244 + v268<<(uint(int32(4))%32)
	v297 = *(*int64)(unsafe.Add(mBase, uint32(v288)))
	*(*int64)(unsafe.Add(mBase, uint32(v296))) = v297
	v299 = *(*int64)(unsafe.Add(mBase, uint32(v288)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v296)+8)) = v299
	v301 = int32(1)
	v302 = v285 << (uint(v301) % 32)
	v304 = v302 | v301
	if base.Ui32(v304) < base.Ui32(v252) {
		v262 = v304
		v265 = v302
		v268 = v285
		goto L98
	} else {
		goto L111
	}
L111:
	;
	goto L99
L112:
	;
	v445 = v238
	goto L4
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+24)) = v217
	v349 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v351 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	if v351 != 0 {
		goto L114
	} else {
		goto L115
	}
L114:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L18
	} else {
		goto L117
	}
L115:
	;
	goto L116
L116:
	;
	v354 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	if base.Ui32(v354) < base.Ui32(int32(2)) {
		goto L119
	} else {
		goto L120
	}
L117:
	;
	goto L116
L118:
	;
	v423 = v349 + v410<<(uint(int32(4))%32)
	v424 = *(*int64)(unsafe.Add(mBase, uint32(v15)+12))
	*(*int64)(unsafe.Add(mBase, uint32(v423))) = v424
	v426 = *(*int64)(unsafe.Add(mBase, uint32(v15)+20))
	*(*int64)(unsafe.Add(mBase, uint32(v423)+8)) = v426
	v445 = int32(1)
	goto L4
L119:
	;
	v410 = int32(0)
	goto L118
L120:
	;
	goto L121
L121:
	;
	v362 = int32(1)
	v363 = v4
	v366 = v4
	goto L122
L122:
	;
	v374 = v363 + int32(2)
	if base.Ui32(v354) <= base.Ui32(v374) {
		goto L124
	} else {
		goto L125
	}
L123:
	;
	v410 = v388
	goto L118
L124:
	;
	v388 = v362
	goto L126
L125:
	;
	v376 = int32(4)
	v382 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v383 = m.T0[v382].(func(*base.Module, int32, int32, int32) int32)(m, v349+v362<<(uint(v376)%32), v349+v374<<(uint(v376)%32), l0)
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L18
	} else {
		goto L127
	}
L126:
	;
	v391 = v349 + v388<<(uint(int32(4))%32)
	v392 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v393 = m.T0[v392].(func(*base.Module, int32, int32, int32) int32)(m, v15+int32(12), v391, l0)
	mBase = m.M
	v394 = m.ExcPending
	if v394 != 0 {
		goto L18
	} else {
		goto L131
	}
L127:
	;
	if int32(0) < v383 {
		goto L128
	} else {
		goto L129
	}
L128:
	;
	v387 = v374
	goto L130
L129:
	;
	v387 = v362
	goto L130
L130:
	;
	v388 = v387
	goto L126
L131:
	;
	if v393 <= int32(0) {
		goto L132
	} else {
		goto L133
	}
L132:
	;
	v410 = v366
	goto L118
L133:
	;
	goto L134
L134:
	;
	v399 = v349 + v366<<(uint(int32(4))%32)
	v400 = *(*int64)(unsafe.Add(mBase, uint32(v391)))
	*(*int64)(unsafe.Add(mBase, uint32(v399))) = v400
	v402 = *(*int64)(unsafe.Add(mBase, uint32(v391)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v399)+8)) = v402
	v404 = int32(1)
	v405 = v388 << (uint(v404) % 32)
	v407 = v405 | v404
	if base.Ui32(v407) < base.Ui32(v354) {
		v362 = v407
		v363 = v405
		v366 = v388
		goto L122
	} else {
		goto L135
	}
L135:
	;
	goto L123
L136:
	;
	F_errmsg_internal(m, int32(351842), int32(0))
	mBase = m.M
	v436 = m.ExcPending
	if v436 != 0 {
		goto L18
	} else {
		goto L137
	}
L137:
	;
	F_errfinish(m, int32(492667), int32(1698), int32(245806))
	mBase = m.M
	v441 = m.ExcPending
	if v441 != 0 {
		goto L18
	} else {
		goto L138
	}
L138:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L139:
	;
	F_errmsg_internal(m, int32(371777), int32(0))
	mBase = m.M
	v467 = m.ExcPending
	if v467 != 0 {
		goto L18
	} else {
		goto L140
	}
L140:
	;
	F_errfinish(m, int32(492667), int32(2862), int32(281794))
	mBase = m.M
	v472 = m.ExcPending
	if v472 != 0 {
		goto L18
	} else {
		goto L141
	}
L141:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L142:
	;
	F_errmsg_internal(m, int32(284919), int32(0))
	mBase = m.M
	v480 = m.ExcPending
	if v480 != 0 {
		goto L18
	} else {
		goto L143
	}
L143:
	;
	F_errfinish(m, int32(492667), int32(1628), int32(245806))
	mBase = m.M
	v485 = m.ExcPending
	if v485 != 0 {
		goto L18
	} else {
		goto L144
	}
L144:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L145:
	;
	F_errmsg_internal(m, int32(371777), int32(0))
	mBase = m.M
	v493 = m.ExcPending
	if v493 != 0 {
		goto L18
	} else {
		goto L146
	}
L146:
	;
	F_errfinish(m, int32(492667), int32(2862), int32(281794))
	mBase = m.M
	v498 = m.ExcPending
	if v498 != 0 {
		goto L18
	} else {
		goto L147
	}
L147:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_tuplesort_gettupleslot(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = int32(4515248)
	v13 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v15
	v17 = F_tuplesort_gettuple_common(m, l0, l1, v10)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		return int32(0)
	} else {
		if v17 == int32(0) {
			*(*int32)(unsafe.Add(mBase, _consts[9])) = v13
			*(*int32)(unsafe.Add(mBase, uint32(v10))) = int32(0)
			v49 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
			v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)+12))
			m.T0[v50].(func(*base.Module, int32))(m, l3)
			mBase = m.M
			v52 = m.ExcPending
			if v52 != 0 {
				return int32(0)
			} else {
				v55 = int32(0)
				m.G0 = v10 + int32(16)
				return v55
			}
		} else {
			*(*int32)(unsafe.Add(mBase, _consts[9])) = v13
			v29 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
			if v29 == int32(0) {
				v49 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
				v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)+12))
				m.T0[v50].(func(*base.Module, int32))(m, l3)
				mBase = m.M
				v52 = m.ExcPending
				if v52 != 0 {
					return int32(0)
				} else {
					v55 = int32(0)
					m.G0 = v10 + int32(16)
					return v55
				}
			} else {
				if l4 == int32(0) {
				} else {
					v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
					v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)+24))
					if v35 == int32(0) {
					} else {
						v38 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
						*(*int32)(unsafe.Add(mBase, uint32(l4))) = v38
					}
				}
				if l2 != 0 {
					v41 = F_heap_copy_minimal_tuple(m, v29, int32(0))
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v10))) = v41
						v44 = v41
						v45 = F_ExecStoreMinimalTuple(m, v44, l3, l2)
						mBase = m.M
						v46 = m.ExcPending
						if v46 != 0 {
							return int32(0)
						} else {
							v55 = int32(1)
							m.G0 = v10 + int32(16)
							return v55
						}
					}
				} else {
					v44 = v29
					v45 = F_ExecStoreMinimalTuple(m, v44, l3, l2)
					mBase = m.M
					v46 = m.ExcPending
					if v46 != 0 {
						return int32(0)
					} else {
						v55 = int32(1)
						m.G0 = v10 + int32(16)
						return v55
					}
				}
			}
		}
	}
}
func F_tuplesort_initialize_shared(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v21 int32
	_ = v21
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = int64(0)
	F_SharedFileSetInit(m, l0+int32(12), l2)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = l1
		if int32(0) < l1 {
			v21 = F__emscripten_memset_bulkmem(m, l0+int32(72), base.I32_extend8_s(int32(0)), l1<<(uint(int32(3))%32))
			mBase = m.M
		} else {
		}
		return
	}
}
