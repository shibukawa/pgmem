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
	v10 = int32(_a_F_tuplesort_getgintuple_0)
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_tuplesort_getgintuple[0]))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, _c_F_tuplesort_getgintuple[0])) = v13
	v16 = F_tuplesort_gettuple_common(m, l0, int32(1), v8)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return int32(0)
	} else {
		if v16 == int32(0) {
			*(*int32)(unsafe.Add(mBase, _c_F_tuplesort_getgintuple[0])) = v11
			v32 = v3
		} else {
			*(*int32)(unsafe.Add(mBase, _c_F_tuplesort_getgintuple[0])) = v11
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
	v8 = int32(_a_F_tuplesort_getheaptuple_0)
	v9 = *(*int32)(unsafe.Add(mBase, _c_F_tuplesort_getheaptuple[0]))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, _c_F_tuplesort_getheaptuple[0])) = v11
	v14 = F_tuplesort_gettuple_common(m, l0, int32(1), v6)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, _c_F_tuplesort_getheaptuple[0])) = v9
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
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
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
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
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
	var v481 int32
	_ = v481
	var v485 int32
	_ = v485
	var v490 int32
	_ = v490
	v4 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(32)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	switch v17 - int32(3) {
	case 0:
		goto L8
	case 1:
		goto L7
	case 2:
		goto L6
	default:
		goto L5
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v481 = m.ExcPending
	if v481 != 0 {
		goto L17
	} else {
		goto L141
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v463 = m.ExcPending
	if v463 != 0 {
		goto L17
	} else {
		goto L138
	}
L3:
	;
	m.G0 = v15 + int32(32)
	return v445
L4:
	;
	v442 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+208)) = uint8(v442)
	v445 = v96
	goto L3
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v432 = m.ExcPending
	if v432 != 0 {
		goto L17
	} else {
		goto L135
	}
L6:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
	if v199 != 0 {
		goto L70
	} else {
		goto L71
	}
L7:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
	if v83 != 0 {
		goto L27
	} else {
		goto L28
	}
L8:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
	if l1 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	if v20 < v21 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	goto L11
L11:
	;
	v58 = int32(0)
	if v20 <= v58 {
		v445 = v58
		goto L3
	} else {
		goto L21
	}
L12:
	;
	v23 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v20 + v23
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v30 = v27 + v20<<(uint(int32(4))%32)
	v31 = *(*int64)(unsafe.Add(mBase, uint32(v30)+8))
	*(*int64)(unsafe.Add(mBase, uint32(l2)+8)) = v31
	v33 = *(*int64)(unsafe.Add(mBase, uint32(v30)))
	*(*int64)(unsafe.Add(mBase, uint32(l2))) = v33
	v445 = v23
	goto L3
L13:
	;
	goto L14
L14:
	;
	v35 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+208)) = uint8(v35)
	v37 = int32(0)
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+68)))
	if v38 != v35 {
		v445 = v37
		goto L3
	} else {
		goto L15
	}
L15:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	if v20 < v41 {
		v445 = v37
		goto L3
	} else {
		goto L16
	}
L16:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	return int32(0)
L18:
	;
	F_errmsg_internal(m, int32(_a_F_tuplesort_gettuple_common_0), int32(0))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L17
	} else {
		goto L19
	}
L19:
	;
	F_errfinish(m, int32(_a_F_tuplesort_gettuple_common_1), int32(1498), int32(_a_F_tuplesort_gettuple_common_2))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L17
	} else {
		goto L20
	}
L20:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L21:
	;
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+208)))
	if v61 == int32(1) {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v77 = v72 + v71<<(uint(int32(4))%32) - int32(16)
	v78 = *(*int64)(unsafe.Add(mBase, uint32(v77)+8))
	*(*int64)(unsafe.Add(mBase, uint32(l2)+8)) = v78
	v80 = *(*int64)(unsafe.Add(mBase, uint32(v77)))
	*(*int64)(unsafe.Add(mBase, uint32(l2))) = v80
	v445 = int32(1)
	goto L3
L23:
	;
	v64 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+208)) = uint8(v64)
	v71 = v20
	goto L22
L24:
	;
	goto L25
L25:
	;
	v66 = int32(1)
	v67 = v20 - v66
	*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v67
	if v20 == v66 {
		v445 = v58
		goto L3
	} else {
		goto L26
	}
L26:
	;
	v71 = v67
	goto L22
L27:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+148))
	if base.Ui32(v83) < base.Ui32(v84) {
		goto L31
	} else {
		goto L32
	}
L28:
	;
	goto L29
L29:
	;
	v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+208)))
	if l1 != 0 {
		goto L35
	} else {
		goto L36
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+164)) = int32(0)
	goto L29
L31:
	;
	F_pfree(m, v83)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L17
	} else {
		goto L34
	}
L32:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	if base.Ui32(v86) <= base.Ui32(v83) {
		goto L31
	} else {
		goto L33
	}
L33:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	*(*int32)(unsafe.Add(mBase, uint32(v83))) = v88
	*(*int32)(unsafe.Add(mBase, uint32(l0)+156)) = v83
	goto L30
L34:
	;
	goto L30
L35:
	;
	v96 = int32(0)
	if v95&int32(1) != 0 {
		v445 = v96
		goto L3
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	if v95&int32(1) != 0 {
		goto L45
	} else {
		goto L46
	}
L38:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	v103 = F_LogicalTapeRead(m, v99, v15+int32(12), int32(4))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L17
	} else {
		goto L39
	}
L39:
	;
	if v103 != int32(4) {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	if v107 == int32(0) {
		goto L4
	} else {
		goto L41
	}
L41:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	m.T0[v111].(func(*base.Module, int32, int32, int32, int32))(m, l0, l2, v110, v107)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L17
	} else {
		goto L42
	}
L42:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+164)) = v114
	v445 = int32(1)
	goto L3
L43:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	v186 = F_getlen(m, v185)
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L17
	} else {
		goto L66
	}
L44:
	;
	v179 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+208)) = uint8(v179)
	goto L43
L45:
	;
	v122 = F_LogicalTapeBackspace(m, v117, int32(8))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L17
	} else {
		goto L49
	}
L46:
	;
	goto L47
L47:
	;
	v137 = int32(0)
	v139 = F_LogicalTapeBackspace(m, v117, int32(4))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L17
	} else {
		goto L55
	}
L48:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L17
	} else {
		goto L50
	}
L49:
	;
	switch v122 {
	case 0:
		v445 = int32(0)
		goto L3
	default:
		goto L48
	case 8:
		goto L44
	}
L50:
	;
	F_errmsg_internal(m, int32(_a_F_tuplesort_gettuple_common_3), int32(0))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L17
	} else {
		goto L51
	}
L51:
	;
	F_errfinish(m, int32(_a_F_tuplesort_gettuple_common_1), int32(1581), int32(_a_F_tuplesort_gettuple_common_2))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L17
	} else {
		goto L52
	}
L52:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L53:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	v155 = F_getlen(m, v154)
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L17
	} else {
		goto L59
	}
L54:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L17
	} else {
		goto L56
	}
L55:
	;
	switch v139 {
	case 0:
		v445 = v137
		goto L3
	default:
		goto L54
	case 4:
		goto L53
	}
L56:
	;
	F_errmsg_internal(m, int32(_a_F_tuplesort_gettuple_common_3), int32(0))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L17
	} else {
		goto L57
	}
L57:
	;
	F_errfinish(m, int32(_a_F_tuplesort_gettuple_common_1), int32(1595), int32(_a_F_tuplesort_gettuple_common_2))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L17
	} else {
		goto L58
	}
L58:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L59:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	v159 = v155 + int32(8)
	v160 = F_LogicalTapeBackspace(m, v157, v159)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L17
	} else {
		goto L60
	}
L60:
	;
	if v160 == v155+int32(4) {
		v445 = v137
		goto L3
	} else {
		goto L61
	}
L61:
	;
	if v159 == v160 {
		goto L43
	} else {
		goto L62
	}
L62:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L17
	} else {
		goto L63
	}
L63:
	;
	F_errmsg_internal(m, int32(_a_F_tuplesort_gettuple_common_4), int32(0))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L17
	} else {
		goto L64
	}
L64:
	;
	F_errfinish(m, int32(_a_F_tuplesort_gettuple_common_1), int32(1615), int32(_a_F_tuplesort_gettuple_common_2))
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L17
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
	v188 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	v189 = F_LogicalTapeBackspace(m, v188, v186)
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L17
	} else {
		goto L67
	}
L67:
	;
	if v189 != v186 {
		goto L2
	} else {
		goto L68
	}
L68:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	v193 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	m.T0[v193].(func(*base.Module, int32, int32, int32, int32))(m, l0, l2, v192, v186)
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L17
	} else {
		goto L69
	}
L69:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+164)) = v196
	v445 = int32(1)
	goto L3
L70:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(l0)+148))
	if base.Ui32(v199) < base.Ui32(v200) {
		goto L74
	} else {
		goto L75
	}
L71:
	;
	goto L72
L72:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	if v211 <= int32(0) {
		goto L78
	} else {
		goto L79
	}
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+164)) = int32(0)
	goto L72
L74:
	;
	F_pfree(m, v199)
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L17
	} else {
		goto L77
	}
L75:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	if base.Ui32(v202) <= base.Ui32(v199) {
		goto L74
	} else {
		goto L76
	}
L76:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	*(*int32)(unsafe.Add(mBase, uint32(v199))) = v204
	*(*int32)(unsafe.Add(mBase, uint32(l0)+156)) = v199
	goto L73
L77:
	;
	goto L73
L78:
	;
	v445 = int32(0)
	goto L3
L79:
	;
	goto L80
L80:
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
		goto L17
	} else {
		goto L81
	}
L81:
	;
	if v230 != int32(4) {
		goto L1
	} else {
		goto L82
	}
L82:
	;
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v15)+28))
	if v234 == int32(0) {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v238 = int32(1)
	v239 = v237 - v238
	*(*int32)(unsafe.Add(mBase, uint32(l0)+136)) = v239
	if int32(0) < v239 {
		goto L86
	} else {
		goto L87
	}
L84:
	;
	goto L85
L85:
	;
	v345 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	m.T0[v345].(func(*base.Module, int32, int32, int32, int32))(m, l0, v15+int32(12), v221, v234)
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L17
	} else {
		goto L112
	}
L86:
	;
	v244 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v248 = *(*int32)(unsafe.Add(mBase, _c_F_tuplesort_gettuple_common[0]))
	if v248 != 0 {
		goto L89
	} else {
		goto L90
	}
L87:
	;
	goto L88
L88:
	;
	v337 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+180)) = v337 - int32(1)
	F_LogicalTapeClose(m, v221)
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L17
	} else {
		goto L111
	}
L89:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L17
	} else {
		goto L92
	}
L90:
	;
	v252 = v239
	goto L91
L91:
	;
	v253 = v239<<(uint(int32(4))%32) + v244
	if base.Ui32(v252) < base.Ui32(int32(2)) {
		goto L94
	} else {
		goto L95
	}
L92:
	;
	v251 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v252 = v251
	goto L91
L93:
	;
	v320 = v244 + v308<<(uint(int32(4))%32)
	v321 = *(*int64)(unsafe.Add(mBase, uint32(v253)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v320)+8)) = v321
	v323 = *(*int64)(unsafe.Add(mBase, uint32(v253)))
	*(*int64)(unsafe.Add(mBase, uint32(v320))) = v323
	goto L88
L94:
	;
	v308 = int32(0)
	goto L93
L95:
	;
	goto L96
L96:
	;
	v261 = int32(1)
	v262 = v4
	v268 = v4
	goto L97
L97:
	;
	v271 = v268 + int32(2)
	if base.Ui32(v252) <= base.Ui32(v271) {
		goto L99
	} else {
		goto L100
	}
L98:
	;
	v308 = v285
	goto L93
L99:
	;
	v285 = v261
	goto L101
L100:
	;
	v273 = int32(4)
	v279 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v280 = m.T0[v279].(func(*base.Module, int32, int32, int32) int32)(m, v244+v261<<(uint(v273)%32), v244+v271<<(uint(v273)%32), l0)
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L17
	} else {
		goto L102
	}
L101:
	;
	v288 = v244 + v285<<(uint(int32(4))%32)
	v289 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v290 = m.T0[v289].(func(*base.Module, int32, int32, int32) int32)(m, v253, v288, l0)
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L17
	} else {
		goto L106
	}
L102:
	;
	if int32(0) < v280 {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	v284 = v271
	goto L105
L104:
	;
	v284 = v261
	goto L105
L105:
	;
	v285 = v284
	goto L101
L106:
	;
	if v290 <= int32(0) {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	v308 = v262
	goto L93
L108:
	;
	goto L109
L109:
	;
	v296 = v244 + v262<<(uint(int32(4))%32)
	v297 = *(*int64)(unsafe.Add(mBase, uint32(v288)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v296)+8)) = v297
	v299 = *(*int64)(unsafe.Add(mBase, uint32(v288)))
	*(*int64)(unsafe.Add(mBase, uint32(v296))) = v299
	v301 = int32(1)
	v302 = v285 << (uint(v301) % 32)
	v304 = v302 | v301
	if base.Ui32(v304) < base.Ui32(v252) {
		v261 = v304
		v262 = v285
		v268 = v302
		goto L97
	} else {
		goto L110
	}
L110:
	;
	goto L98
L111:
	;
	v445 = v238
	goto L3
L112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+24)) = v217
	v349 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v351 = *(*int32)(unsafe.Add(mBase, _c_F_tuplesort_gettuple_common[0]))
	if v351 != 0 {
		goto L113
	} else {
		goto L114
	}
L113:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L17
	} else {
		goto L116
	}
L114:
	;
	goto L115
L115:
	;
	v354 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	if base.Ui32(v354) < base.Ui32(int32(2)) {
		goto L118
	} else {
		goto L119
	}
L116:
	;
	goto L115
L117:
	;
	v423 = v349 + v410<<(uint(int32(4))%32)
	v424 = *(*int64)(unsafe.Add(mBase, uint32(v15)+20))
	*(*int64)(unsafe.Add(mBase, uint32(v423)+8)) = v424
	v426 = *(*int64)(unsafe.Add(mBase, uint32(v15)+12))
	*(*int64)(unsafe.Add(mBase, uint32(v423))) = v426
	v445 = int32(1)
	goto L3
L118:
	;
	v410 = int32(0)
	goto L117
L119:
	;
	goto L120
L120:
	;
	v361 = int32(1)
	v362 = v4
	v363 = v4
	goto L121
L121:
	;
	v374 = v363 + int32(2)
	if base.Ui32(v354) <= base.Ui32(v374) {
		goto L123
	} else {
		goto L124
	}
L122:
	;
	v410 = v388
	goto L117
L123:
	;
	v388 = v361
	goto L125
L124:
	;
	v376 = int32(4)
	v382 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v383 = m.T0[v382].(func(*base.Module, int32, int32, int32) int32)(m, v349+v361<<(uint(v376)%32), v349+v374<<(uint(v376)%32), l0)
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L17
	} else {
		goto L126
	}
L125:
	;
	v391 = v349 + v388<<(uint(int32(4))%32)
	v392 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v393 = m.T0[v392].(func(*base.Module, int32, int32, int32) int32)(m, v15+int32(12), v391, l0)
	mBase = m.M
	v394 = m.ExcPending
	if v394 != 0 {
		goto L17
	} else {
		goto L130
	}
L126:
	;
	if int32(0) < v383 {
		goto L127
	} else {
		goto L128
	}
L127:
	;
	v387 = v374
	goto L129
L128:
	;
	v387 = v361
	goto L129
L129:
	;
	v388 = v387
	goto L125
L130:
	;
	if v393 <= int32(0) {
		goto L131
	} else {
		goto L132
	}
L131:
	;
	v410 = v362
	goto L117
L132:
	;
	goto L133
L133:
	;
	v399 = v349 + v362<<(uint(int32(4))%32)
	v400 = *(*int64)(unsafe.Add(mBase, uint32(v391)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v399)+8)) = v400
	v402 = *(*int64)(unsafe.Add(mBase, uint32(v391)))
	*(*int64)(unsafe.Add(mBase, uint32(v399))) = v402
	v404 = int32(1)
	v405 = v388 << (uint(v404) % 32)
	v407 = v405 | v404
	if base.Ui32(v407) < base.Ui32(v354) {
		v361 = v407
		v362 = v388
		v363 = v405
		goto L121
	} else {
		goto L134
	}
L134:
	;
	goto L122
L135:
	;
	F_errmsg_internal(m, int32(_a_F_tuplesort_gettuple_common_5), int32(0))
	mBase = m.M
	v436 = m.ExcPending
	if v436 != 0 {
		goto L17
	} else {
		goto L136
	}
L136:
	;
	F_errfinish(m, int32(_a_F_tuplesort_gettuple_common_1), int32(1698), int32(_a_F_tuplesort_gettuple_common_2))
	mBase = m.M
	v441 = m.ExcPending
	if v441 != 0 {
		goto L17
	} else {
		goto L137
	}
L137:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L138:
	;
	F_errmsg_internal(m, int32(_a_F_tuplesort_gettuple_common_4), int32(0))
	mBase = m.M
	v467 = m.ExcPending
	if v467 != 0 {
		goto L17
	} else {
		goto L139
	}
L139:
	;
	F_errfinish(m, int32(_a_F_tuplesort_gettuple_common_1), int32(1628), int32(_a_F_tuplesort_gettuple_common_2))
	mBase = m.M
	v472 = m.ExcPending
	if v472 != 0 {
		goto L17
	} else {
		goto L140
	}
L140:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L141:
	;
	F_errmsg_internal(m, int32(_a_F_tuplesort_gettuple_common_6), int32(0))
	mBase = m.M
	v485 = m.ExcPending
	if v485 != 0 {
		goto L17
	} else {
		goto L142
	}
L142:
	;
	F_errfinish(m, int32(_a_F_tuplesort_gettuple_common_1), int32(2862), int32(_a_F_tuplesort_gettuple_common_7))
	mBase = m.M
	v490 = m.ExcPending
	if v490 != 0 {
		goto L17
	} else {
		goto L143
	}
L143:
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
	v12 = int32(_a_F_tuplesort_gettupleslot_0)
	v13 = *(*int32)(unsafe.Add(mBase, _c_F_tuplesort_gettupleslot[0]))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, _c_F_tuplesort_gettupleslot[0])) = v15
	v17 = F_tuplesort_gettuple_common(m, l0, l1, v10)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		return int32(0)
	} else {
		if v17 == int32(0) {
			*(*int32)(unsafe.Add(mBase, _c_F_tuplesort_gettupleslot[0])) = v13
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
			*(*int32)(unsafe.Add(mBase, _c_F_tuplesort_gettupleslot[0])) = v13
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
	var v16 int32
	_ = v16
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = int64(0)
	F_SharedFileSetInit(m, l0+int32(12), l2)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = l1
		if l1 <= int32(0) {
		} else {
			v16 = l1 << (uint(int32(3)) % 32)
			if v16 == int32(0) {
			} else {
				base.MemoryFill(m, l0+int32(72), int32(0), v16)
			}
		}
		return
	}
}
