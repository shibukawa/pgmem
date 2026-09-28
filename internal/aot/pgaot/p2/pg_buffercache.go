package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_pg_buffercache_evict(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int64
	_ = v51
	var v53 int64
	_ = v53
	var v63 int64
	_ = v63
	var v72 int64
	_ = v72
	var v87 int32
	_ = v87
	var v88 int64
	_ = v88
	var v91 int64
	_ = v91
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v123 int32
	_ = v123
	var v125 int64
	_ = v125
	var v127 int64
	_ = v127
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v147 int64
	_ = v147
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int64
	_ = v157
	var v158 int32
	_ = v158
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v189 int32
	_ = v189
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v207 int32
	_ = v207
	v2 = int32(0)
	v7 = m.G0
	v9 = v7 + int32(-64)
	m.G0 = v9
	*(*uint16)(unsafe.Add(mBase, uint32(v9)+30)) = uint16(v2)
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v17 = F_get_call_result_type(m, l0, v2, v7+int32(-4))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L3
	} else {
		goto L48
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L3
	} else {
		goto L44
	}
L3:
	;
	return int64(0)
L4:
	;
	if v17 == int32(1) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v23 = F_superuser(m)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L3
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L3
	} else {
		goto L41
	}
L8:
	;
	if v23 == int32(0) {
		goto L2
	} else {
		goto L9
	}
L9:
	;
	if v13 <= int32(0) {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	v30 = *(*int32)(unsafe.Add(mBase, _c_F_pg_buffercache_evict[0]))
	if v30 < v13 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v34 = m.G0
	v36 = v34 - int32(32)
	m.G0 = v36
	v39 = *(*int32)(unsafe.Add(mBase, _c_F_pg_buffercache_evict[1]))
	F_ResourceOwnerEnlarge(m, v39)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L3
	} else {
		goto L12
	}
L12:
	;
	F_ReservePrivateRefCountEntry(m)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L3
	} else {
		goto L13
	}
L13:
	;
	v45 = *(*int32)(unsafe.Add(mBase, _c_F_pg_buffercache_evict[2]))
	v48 = v45 + v13*int32(56)
	v50 = v48 - int32(32)
	v51 = int64(4194304)
	v53 = base.AtomicRmwOr64(m, v50, int32(0), v51)
	if v53&v51 != int64(0) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v63 = v53
	goto L17
L15:
	;
	goto L16
L16:
	;
	v140 = F_EvictUnpinnedBufferInternal(m, v48-int32(56), v7+int32(-35))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L3
	} else {
		goto L38
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+28)) = int32(_a_F_pg_buffercache_evict_0)
	*(*int32)(unsafe.Add(mBase, uint32(v36)+24)) = int32(_a_F_pg_buffercache_evict_1)
	*(*int32)(unsafe.Add(mBase, uint32(v36)+20)) = int32(_a_F_pg_buffercache_evict_2)
	*(*int32)(unsafe.Add(mBase, uint32(v36)+16)) = int32(0)
	v72 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v36)+8)) = v72
	if v63&int64(4194304) != v72 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	goto L16
L19:
	;
	goto L22
L20:
	;
	goto L21
L21:
	;
	v105 = int32(_a_F_pg_buffercache_evict_3)
	v106 = *(*int32)(unsafe.Add(mBase, _c_F_pg_buffercache_evict[3]))
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v36+int32(8))+8))
	if v108 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L22:
	;
	F_perform_spin_delay(m, v36+int32(8))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L3
	} else {
		goto L24
	}
L23:
	;
	goto L21
L24:
	;
	v88 = int64(0)
	v91 = base.AtomicRmwCmpxchg64(m, v50, int32(0), v88, v88)
	if v91&int64(4194304) != v88 {
		goto L22
	} else {
		goto L25
	}
L25:
	;
	goto L23
L26:
	;
	v125 = int64(4194304)
	v127 = base.AtomicRmwOr64(m, v50, int32(0), v125)
	if v127&v125 != int64(0) {
		v63 = v127
		goto L17
	} else {
		goto L37
	}
L27:
	;
	goto L26
L28:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_pg_buffercache_evict[3])) = v123
	goto L27
L29:
	;
	if int32(999) < v106 {
		goto L27
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	if v106 < int32(11) {
		goto L27
	} else {
		goto L36
	}
L32:
	;
	v113 = int32(900)
	if v113 <= v106 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v116 = v113
	goto L35
L34:
	;
	v116 = v106
	goto L35
L35:
	;
	v123 = v116 + int32(100)
	goto L28
L36:
	;
	v123 = v106 - int32(1)
	goto L28
L37:
	;
	goto L18
L38:
	;
	m.G0 = v36 + int32(32)
	*(*int64)(unsafe.Add(mBase, uint32(v9)+32)) = base.I64_extend_i32_u(v140)
	v147 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v9)+29)))
	*(*int64)(unsafe.Add(mBase, uint32(v9)+40)) = v147
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v9)+60))
	v154 = F_heap_form_tuple(m, v149, v7+int32(-32), v7+int32(-34))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L3
	} else {
		goto L39
	}
L39:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v154)+16))
	v157 = F_HeapTupleHeaderGetDatum(m, v156)
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L3
	} else {
		goto L40
	}
L40:
	;
	m.G0 = v9 - int32(-64)
	return v157
L41:
	;
	F_errmsg_internal(m, int32(_a_F_pg_buffercache_evict_4), int32(0))
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L3
	} else {
		goto L42
	}
L42:
	;
	F_errfinish(m, int32(_a_F_pg_buffercache_evict_5), int32(710), int32(_a_F_pg_buffercache_evict_6))
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L3
	} else {
		goto L43
	}
L43:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L44:
	;
	F_errcode(m, int32(16797828))
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L3
	} else {
		goto L45
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = int32(_a_F_pg_buffercache_evict_6)
	F_errmsg(m, int32(_a_F_pg_buffercache_evict_7), v7+int32(-48))
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L3
	} else {
		goto L46
	}
L46:
	;
	F_errfinish(m, int32(_a_F_pg_buffercache_evict_5), int32(691), int32(_a_F_pg_buffercache_evict_8))
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L3
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
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v13
	F_errmsg_internal(m, int32(_a_F_pg_buffercache_evict_9), v9)
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L3
	} else {
		goto L49
	}
L49:
	;
	F_errfinish(m, int32(_a_F_pg_buffercache_evict_5), int32(715), int32(_a_F_pg_buffercache_evict_6))
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L3
	} else {
		goto L50
	}
L50:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pg_buffercache_os_pages_internal(m *base.Module, l0 int32, l1 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
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
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v118 int32
	_ = v118
	var v125 int32
	_ = v125
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v196 int32
	_ = v196
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v210 int32
	_ = v210
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v249 int32
	_ = v249
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v261 int32
	_ = v261
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v300 int32
	_ = v300
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v311 int32
	_ = v311
	var v312 int64
	_ = v312
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v317 int64
	_ = v317
	var v318 int32
	_ = v318
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v328 int32
	_ = v328
	var v333 int32
	_ = v333
	var v341 int64
	_ = v341
	var v344 int32
	_ = v344
	var v350 int32
	_ = v350
	var v352 int32
	_ = v352
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v367 int32
	_ = v367
	var v377 int32
	_ = v377
	var v396 int64
	_ = v396
	var v402 int32
	_ = v402
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v422 int64
	_ = v422
	var v423 int64
	_ = v423
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v430 int32
	_ = v430
	var v431 int64
	_ = v431
	var v432 int32
	_ = v432
	var v435 int64
	_ = v435
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v444 int32
	_ = v444
	var v447 int32
	_ = v447
	var v450 int32
	_ = v450
	var v452 int32
	_ = v452
	var v454 int64
	_ = v454
	var v457 int32
	_ = v457
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v465 int64
	_ = v465
	var v466 int32
	_ = v466
	var v467 int64
	_ = v467
	var v471 int32
	_ = v471
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v479 int32
	_ = v479
	var v485 int64
	_ = v485
	var v493 int32
	_ = v493
	var v497 int32
	_ = v497
	var v502 int32
	_ = v502
	var v506 int32
	_ = v506
	var v510 int32
	_ = v510
	var v515 int32
	_ = v515
	var v519 int32
	_ = v519
	var v523 int32
	_ = v523
	var v528 int32
	_ = v528
	v2 = l1
	v17 = m.G0
	v19 = v17 + int32(-64)
	m.G0 = v19
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+16))
	if v22 != 0 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v519 = m.ExcPending
	if v519 != 0 {
		goto L34
	} else {
		goto L116
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v506 = m.ExcPending
	if v506 != 0 {
		goto L34
	} else {
		goto L113
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v493 = m.ExcPending
	if v493 != 0 {
		goto L34
	} else {
		goto L110
	}
L4:
	;
	v420 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v421 = *(*int32)(unsafe.Add(mBase, uint32(v420)+16))
	goto L97
L5:
	;
	if v2 != 0 {
		goto L3
	} else {
		goto L6
	}
L6:
	;
	v23 = m.G0
	v25 = v23 - int32(16)
	m.G0 = v25
	v28 = int32(*(*int16)(unsafe.Add(mBase, _c_F_pg_buffercache_os_pages_internal[0])))
	if v28 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+12)) = v54
	v57 = *(*int32)(unsafe.Add(mBase, _c_F_pg_buffercache_os_pages_internal[1]))
	if v57 == int32(1) {
		goto L21
	} else {
		goto L22
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_pg_buffercache_os_pages_internal[2])) = int32(28)
	v54 = int32(-1)
	goto L7
L9:
	;
	goto L10
L10:
	;
	if int32(-2) < v28 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v54 = v52
	goto L7
L12:
	;
	v52 = v28
	goto L11
L13:
	;
	switch v28&int32(255) - int32(1) {
	case 0:
		v52 = int32(_a_F_pg_buffercache_os_pages_internal_0)
		goto L11
	case 1:
		goto L20
	case 2:
		goto L19
	case 3:
		goto L18
	case 4, 10:
		goto L17
	case 5, 6:
		goto L16
	case 7, 8:
		goto L15
	case 9:
		goto L14
	default:
		goto L12
	}
L14:
	;
	v54 = int32(0)
	goto L7
L15:
	;
	v47 = m.Env.Emscripten_get_heap_max(m)
	mBase = m.M
	v54 = int32(base.Ui32(v47) >> (uint(int32(16)) % 32))
	goto L7
L16:
	;
	v54 = int32(1)
	goto L7
L17:
	;
	v54 = int32(2147483647)
	goto L7
L18:
	;
	v54 = int32(_a_F_pg_buffercache_os_pages_internal_1)
	goto L7
L19:
	;
	v54 = int32(_a_F_pg_buffercache_os_pages_internal_2)
	goto L7
L20:
	;
	v54 = int32(_a_F_pg_buffercache_os_pages_internal_3)
	goto L7
L21:
	;
	v61 = v25 + int32(12)
	v64 = *(*int32)(unsafe.Add(mBase, _c_F_pg_buffercache_os_pages_internal[3]))
	if v64 != 0 {
		goto L25
	} else {
		goto L26
	}
L22:
	;
	v82 = v54
	goto L23
L23:
	;
	m.G0 = v25 + int32(16)
	v86 = F_init_MultiFuncCall(m, l0)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L34
	} else {
		goto L35
	}
L24:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
	v82 = v81
	goto L23
L25:
	;
	v68 = v64 << (uint(int32(10)) % 32)
	goto L27
L26:
	;
	v68 = int32(_a_F_pg_buffercache_os_pages_internal_4)
	goto L27
L27:
	;
	if v68 != 0 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	goto L30
L29:
	;
	goto L30
L30:
	;
	if v61 != 0 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v61))) = v68
	goto L33
L32:
	;
	goto L33
L33:
	;
	goto L24
L34:
	;
	return int64(0)
L35:
	;
	v90 = int32(_a_F_pg_buffercache_os_pages_internal_5)
	v91 = *(*int32)(unsafe.Add(mBase, _c_F_pg_buffercache_os_pages_internal[4]))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v86)+24))
	*(*int32)(unsafe.Add(mBase, _c_F_pg_buffercache_os_pages_internal[4])) = v93
	v96 = F_palloc(m, int32(12))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L34
	} else {
		goto L36
	}
L36:
	;
	v101 = F_get_call_result_type(m, l0, int32(0), v17+int32(-4))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L34
	} else {
		goto L37
	}
L37:
	;
	if v101 != int32(1) {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v19)+60))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v105)))
	if v106 != int32(3) {
		goto L2
	} else {
		goto L39
	}
L39:
	;
	v110 = F_CreateTemplateTupleDesc(m, int32(3))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L34
	} else {
		goto L40
	}
L40:
	;
	F_TupleDescInitEntry(m, v110, int32(1), int32(_a_F_pg_buffercache_os_pages_internal_6), int32(23), int32(-1), int32(0))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L34
	} else {
		goto L41
	}
L41:
	;
	F_TupleDescInitEntry(m, v110, int32(2), int32(_a_F_pg_buffercache_os_pages_internal_7), int32(20), int32(-1), int32(0))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L34
	} else {
		goto L42
	}
L42:
	;
	F_TupleDescInitEntry(m, v110, int32(3), int32(_a_F_pg_buffercache_os_pages_internal_8), int32(23), int32(-1), int32(0))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L34
	} else {
		goto L43
	}
L43:
	;
	v133 = int32(0)
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v110)))
	if v133 < v142 {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	v220 = F_BlessTupleDesc(m, v110)
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L34
	} else {
		goto L63
	}
L45:
	;
	v146 = v110 + int32(28)
	v153 = v133
	v154 = v142
	v156 = v133
	goto L49
L46:
	;
	v210 = v133
	v217 = v142
	goto L47
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v110)+20)) = v217
	*(*int32)(unsafe.Add(mBase, uint32(v110)+16)) = v210
	goto L44
L48:
	;
	v210 = v204
	v217 = v183
	goto L47
L49:
	;
	v162 = v146 + v142<<(uint(int32(3))%32) + v153*int32(100)
	v165 = v146 + v153<<(uint(int32(3))%32)
	if v142 != v154 {
		v183 = v154
		goto L51
	} else {
		goto L52
	}
L50:
	;
	v204 = v142
	goto L48
L51:
	;
	v184 = int32(*(*int16)(unsafe.Add(mBase, uint32(v165)+2)))
	if v184 <= int32(0) {
		v204 = v153
		goto L48
	} else {
		goto L59
	}
L52:
	;
	v167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+7)))
	if v167 != int32(118) {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v183 = v153
	goto L51
L54:
	;
	v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+4)))
	if v170 != int32(1) {
		goto L53
	} else {
		goto L55
	}
L55:
	;
	v173 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+6)))
	if v173&int32(6) != 0 {
		goto L53
	} else {
		goto L56
	}
L56:
	;
	v176 = int32(*(*int16)(unsafe.Add(mBase, uint32(v165)+2)))
	if v176 <= int32(0) {
		goto L53
	} else {
		goto L57
	}
L57:
	;
	v179 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162)+90)))
	if v179 != int32(118) {
		v183 = v142
		goto L51
	} else {
		goto L58
	}
L58:
	;
	goto L53
L59:
	;
	v187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162)+90)))
	if v187 == int32(118) {
		v204 = v153
		goto L48
	} else {
		goto L60
	}
L60:
	;
	v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+5)))
	v196 = (v156 + v190 - int32(1)) & (int32(0) - v190)
	if int32(_a_F_pg_buffercache_os_pages_internal_9) < v196 {
		v204 = v153
		goto L48
	} else {
		goto L61
	}
L61:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v165))) = uint16(v196)
	v202 = v153 + int32(1)
	if v202 != v142 {
		v153 = v202
		v154 = v183
		v156 = v196 + v184
		goto L49
	} else {
		goto L62
	}
L62:
	;
	goto L50
L63:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v96)+4)) = uint8(v2)
	*(*int32)(unsafe.Add(mBase, uint32(v96))) = v220
	v225 = *(*int32)(unsafe.Add(mBase, _c_F_pg_buffercache_os_pages_internal[4]))
	if base.Ui32(v82) <= base.Ui32(int32(_a_F_pg_buffercache_os_pages_internal_10)) {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v231 = base.I32_div_u_s(int32(_a_F_pg_buffercache_os_pages_internal_10), v82&int32(_a_F_pg_buffercache_os_pages_internal_11))
	v232 = int32(24)
	v237 = v231*v232 + v232
	goto L66
L65:
	;
	v237 = int32(48)
	goto L66
L66:
	;
	v239 = *(*int32)(unsafe.Add(mBase, _c_F_pg_buffercache_os_pages_internal[5]))
	v241 = F_MemoryContextAllocHuge(m, v225, v237*v239)
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L34
	} else {
		goto L67
	}
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v96)+8)) = v241
	*(*int32)(unsafe.Add(mBase, _c_F_pg_buffercache_os_pages_internal[4])) = v91
	if v2 == int32(0) {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v268 = *(*int32)(unsafe.Add(mBase, _c_F_pg_buffercache_os_pages_internal[6]))
	v270 = *(*int32)(unsafe.Add(mBase, _c_F_pg_buffercache_os_pages_internal[5]))
	if v270 <= int32(0) {
		goto L75
	} else {
		goto L76
	}
L69:
	;
	v249 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_buffercache_os_pages_internal[7])))
	if v249&int32(1) != 0 {
		goto L68
	} else {
		goto L70
	}
L70:
	;
	v254 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L34
	} else {
		goto L71
	}
L71:
	;
	if v254 == int32(0) {
		goto L68
	} else {
		goto L72
	}
L72:
	;
	F_errmsg_internal(m, int32(_a_F_pg_buffercache_os_pages_internal_12), int32(0))
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L34
	} else {
		goto L73
	}
L73:
	;
	F_errfinish(m, int32(_a_F_pg_buffercache_os_pages_internal_13), int32(431), int32(_a_F_pg_buffercache_os_pages_internal_14))
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L34
	} else {
		goto L74
	}
L74:
	;
	goto L68
L75:
	;
	v396 = int64(0)
	goto L77
L76:
	;
	v274 = int32(0)
	v275 = v274 - v82
	v284 = int32(0)
	v285 = v274
	goto L78
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v86)+16)) = v96
	*(*int64)(unsafe.Add(mBase, uint32(v86)+8)) = v396
	if v2 == int32(0) {
		goto L4
	} else {
		goto L95
	}
L78:
	;
	v295 = *(*int32)(unsafe.Add(mBase, _c_F_pg_buffercache_os_pages_internal[6]))
	v297 = v284 + int32(1)
	v300 = v295 + v297<<(uint(int32(13))%32)
	v304 = *(*int32)(unsafe.Add(mBase, _c_F_pg_buffercache_os_pages_internal[8]))
	if v304 != 0 {
		goto L80
	} else {
		goto L81
	}
L79:
	;
	v396 = base.I64_extend_i32_s(v367)
	goto L77
L80:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L34
	} else {
		goto L83
	}
L81:
	;
	goto L82
L82:
	;
	v308 = *(*int32)(unsafe.Add(mBase, _c_F_pg_buffercache_os_pages_internal[9]))
	v311 = v308 + v284*int32(56)
	v312 = F_LockBufHdr(m, v311)
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L34
	} else {
		goto L84
	}
L83:
	;
	goto L82
L84:
	;
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v311)+20))
	v317 = base.AtomicRmwSub64(m, v311, int32(24), int64(4194304))
	v318 = (v300 + int32(-8192)) & v275
	if base.Ui32(v318) < base.Ui32(v300) {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v323 = base.I32_div_u_s(v318-v268&v275, v82)
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v96)+8))
	v328 = v318
	v333 = v285
	v341 = base.I64_extend_i32_s(v323)
	goto L88
L86:
	;
	v367 = v285
	goto L87
L87:
	;
	v377 = *(*int32)(unsafe.Add(mBase, _c_F_pg_buffercache_os_pages_internal[5]))
	if v297 < v377 {
		v284 = v297
		v285 = v367
		goto L78
	} else {
		goto L94
	}
L88:
	;
	v344 = v325 + v333*int32(24)
	*(*int64)(unsafe.Add(mBase, uint32(v344)+8)) = v341
	*(*int32)(unsafe.Add(mBase, uint32(v344))) = v314 + int32(1)
	if v2 != 0 {
		goto L90
	} else {
		goto L91
	}
L89:
	;
	v367 = v357
	goto L87
L90:
	;
	v350 = *(*int32)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v341)<<(uint(int32(2))%32))))
	v352 = v350
	goto L92
L91:
	;
	v352 = int32(-1)
	goto L92
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v344)+16)) = v352
	v357 = v333 + int32(1)
	v358 = v328 + v82
	if base.Ui32(v358) < base.Ui32(v300) {
		v328 = v358
		v333 = v357
		v341 = v341 + int64(1)
		goto L88
	} else {
		goto L93
	}
L93:
	;
	goto L89
L94:
	;
	goto L79
L95:
	;
	v402 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_pg_buffercache_os_pages_internal[7])) = uint8(v402)
	goto L4
L96:
	;
	m.G0 = v19 - int32(-64)
	return v485
L97:
	;
	v422 = *(*int64)(unsafe.Add(mBase, uint32(v421)))
	v423 = *(*int64)(unsafe.Add(mBase, uint32(v421)+8))
	if base.Ui64(v422) < base.Ui64(v423) {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v425 = *(*int32)(unsafe.Add(mBase, uint32(v421)+16))
	v426 = *(*int32)(unsafe.Add(mBase, uint32(v425)+8))
	v430 = v426 + base.I32_wrap_i64(v422)*int32(24)
	v431 = int64(*(*int32)(unsafe.Add(mBase, uint32(v430))))
	v432 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+29)) = uint8(v432)
	*(*int64)(unsafe.Add(mBase, uint32(v19)+32)) = v431
	v435 = *(*int64)(unsafe.Add(mBase, uint32(v430)+8))
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+30)) = uint8(v432)
	*(*int64)(unsafe.Add(mBase, uint32(v19)+40)) = v435
	v440 = int32(1)
	v441 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v425)+4)))
	if v441 == v440 {
		goto L101
	} else {
		goto L102
	}
L99:
	;
	goto L100
L100:
	;
	F_end_MultiFuncCall(m, l0)
	mBase = m.M
	v475 = m.ExcPending
	if v475 != 0 {
		goto L34
	} else {
		goto L109
	}
L101:
	;
	v444 = *(*int32)(unsafe.Add(mBase, uint32(v430)+16))
	v447 = int32(0)
	if v447 < v444 {
		goto L104
	} else {
		goto L105
	}
L102:
	;
	v452 = v440
	v454 = int64(0)
	goto L103
L103:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+31)) = uint8(v452)
	*(*int64)(unsafe.Add(mBase, uint32(v19)+48)) = v454
	v457 = *(*int32)(unsafe.Add(mBase, uint32(v425)))
	v462 = F_heap_form_tuple(m, v457, v17+int32(-32), v17+int32(-35))
	mBase = m.M
	v463 = m.ExcPending
	if v463 != 0 {
		goto L34
	} else {
		goto L107
	}
L104:
	;
	v450 = v444
	goto L106
L105:
	;
	v450 = v447
	goto L106
L106:
	;
	v452 = int32(base.Ui32(v444) >> (uint(int32(31)) % 32))
	v454 = base.I64_extend_i32_u(v450)
	goto L103
L107:
	;
	v464 = *(*int32)(unsafe.Add(mBase, uint32(v462)+16))
	v465 = F_HeapTupleHeaderGetDatum(m, v464)
	mBase = m.M
	v466 = m.ExcPending
	if v466 != 0 {
		goto L34
	} else {
		goto L108
	}
L108:
	;
	v467 = *(*int64)(unsafe.Add(mBase, uint32(v421)))
	*(*int64)(unsafe.Add(mBase, uint32(v421))) = v467 + int64(1)
	v471 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v471)+20)) = int32(1)
	v485 = v465
	goto L96
L109:
	;
	v476 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v476)+20)) = int32(2)
	v479 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v479)
	v485 = int64(0)
	goto L96
L110:
	;
	F_errmsg_internal(m, int32(_a_F_pg_buffercache_os_pages_internal_15), int32(0))
	mBase = m.M
	v497 = m.ExcPending
	if v497 != 0 {
		goto L34
	} else {
		goto L111
	}
L111:
	;
	F_errfinish(m, int32(_a_F_pg_buffercache_os_pages_internal_13), int32(305), int32(_a_F_pg_buffercache_os_pages_internal_14))
	mBase = m.M
	v502 = m.ExcPending
	if v502 != 0 {
		goto L34
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
	F_errmsg_internal(m, int32(_a_F_pg_buffercache_os_pages_internal_16), int32(0))
	mBase = m.M
	v510 = m.ExcPending
	if v510 != 0 {
		goto L34
	} else {
		goto L114
	}
L114:
	;
	F_errfinish(m, int32(_a_F_pg_buffercache_os_pages_internal_13), int32(397), int32(_a_F_pg_buffercache_os_pages_internal_14))
	mBase = m.M
	v515 = m.ExcPending
	if v515 != 0 {
		goto L34
	} else {
		goto L115
	}
L115:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L116:
	;
	F_errmsg_internal(m, int32(_a_F_pg_buffercache_os_pages_internal_17), int32(0))
	mBase = m.M
	v523 = m.ExcPending
	if v523 != 0 {
		goto L34
	} else {
		goto L117
	}
L117:
	;
	F_errfinish(m, int32(_a_F_pg_buffercache_os_pages_internal_13), int32(394), int32(_a_F_pg_buffercache_os_pages_internal_14))
	mBase = m.M
	v528 = m.ExcPending
	if v528 != 0 {
		goto L34
	} else {
		goto L118
	}
L118:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
