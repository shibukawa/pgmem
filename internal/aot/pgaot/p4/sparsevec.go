package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_sparsevec_cosine_distance(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v16 float32
	_ = v16
	var v19 float64
	_ = v19
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v67 float32
	_ = v67
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v95 float32
	_ = v95
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v106 float32
	_ = v106
	var v108 float32
	_ = v108
	var v111 float32
	_ = v111
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v121 float32
	_ = v121
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v149 float32
	_ = v149
	var v157 int32
	_ = v157
	var v158 float32
	_ = v158
	var v160 float32
	_ = v160
	var v162 float32
	_ = v162
	var v164 float32
	_ = v164
	var v169 float32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v192 float32
	_ = v192
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v213 float32
	_ = v213
	var v222 float32
	_ = v222
	var v224 float32
	_ = v224
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v245 float32
	_ = v245
	var v271 float64
	_ = v271
	var v273 float64
	_ = v273
	var v277 int32
	_ = v277
	var v278 float32
	_ = v278
	var v279 int32
	_ = v279
	var v286 int32
	_ = v286
	var v289 int32
	_ = v289
	var v302 float32
	_ = v302
	var v309 int32
	_ = v309
	var v310 float32
	_ = v310
	var v312 float32
	_ = v312
	var v314 float32
	_ = v314
	var v316 float32
	_ = v316
	var v321 float32
	_ = v321
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v329 int32
	_ = v329
	var v345 float32
	_ = v345
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v366 float32
	_ = v366
	var v374 float32
	_ = v374
	var v376 float32
	_ = v376
	var v377 int32
	_ = v377
	var v380 int32
	_ = v380
	var v398 float32
	_ = v398
	var v423 float64
	_ = v423
	var v428 float64
	_ = v428
	var v432 float64
	_ = v432
	var v440 float64
	_ = v440
	var v447 int32
	_ = v447
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v457 int32
	_ = v457
	var v462 int32
	_ = v462
	v2 = int32(0)
	v16 = float32(0)
	v19 = float64(0)
	v22 = m.G0
	v24 = v22 - int32(16)
	m.G0 = v24
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v27 = F_pg_detoast_datum(m, v26)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v32 = F_pg_detoast_datum(m, v31)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v32)+4))
	if v34 == v35 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v38 = v32 + int32(16)
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v32)+8))
	v42 = v38 + v39<<(uint(int32(2))%32)
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v27)+8))
	if int32(0) < v43 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v447 = m.ExcPending
	if v447 != 0 {
		goto L1
	} else {
		goto L56
	}
L7:
	;
	v47 = v27 + int32(16)
	v50 = v47 + v43<<(uint(int32(2))%32)
	v54 = v2
	v57 = v2
	v67 = v16
	goto L10
L8:
	;
	v271 = v19
	v273 = v19
	goto L9
L9:
	;
	if int32(0) < v39 {
		goto L39
	} else {
		goto L40
	}
L10:
	;
	if v39 < v54 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v126 = v43 & int32(3)
	v127 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v43) {
		goto L29
	} else {
		goto L30
	}
L12:
	;
	v73 = v54
	goto L14
L13:
	;
	v73 = v39
	goto L14
L14:
	;
	v75 = v57 << (uint(int32(2)) % 32)
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v75+v47)))
	v79 = v54
	v82 = v54
	v95 = v67
	goto L15
L15:
	;
	if v79 != v73 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v123 = v57 + int32(1)
	if v123 != v43 {
		v54 = v118
		v57 = v123
		v67 = v121
		goto L10
	} else {
		goto L27
	}
L17:
	;
	v102 = v79 << (uint(int32(2)) % 32)
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v38+v102)))
	if v104 == v78 {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	v118 = v82
	v121 = v95
	goto L19
L19:
	;
	goto L16
L20:
	;
	v106 = *(*float32)(unsafe.Add(mBase, uint32(v50+v75)))
	v108 = *(*float32)(unsafe.Add(mBase, uint32(v42+v102)))
	v111 = base.F32_add(base.F32_mul(v106, v108), v95)
	goto L22
L21:
	;
	v111 = v95
	goto L22
L22:
	;
	v113 = v79 + int32(1)
	if v78 < v104 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v115 = v82
	goto L25
L24:
	;
	v115 = v113
	goto L25
L25:
	;
	if v104 < v78 {
		v79 = v113
		v82 = v115
		v95 = v111
		goto L15
	} else {
		goto L26
	}
L26:
	;
	v118 = v115
	v121 = v111
	goto L19
L27:
	;
	goto L11
L28:
	;
	v271 = base.F64_promote_f32(v245)
	v273 = base.F64_promote_f32(v121)
	goto L9
L29:
	;
	v134 = v127
	v137 = int32(0)
	v149 = v16
	goto L32
L30:
	;
	v177 = v127
	v192 = v16
	goto L31
L31:
	;
	v198 = v177
	v199 = v127
	v213 = v192
	goto L36
L32:
	;
	v157 = v50 + v134<<(uint(int32(2))%32)
	v158 = *(*float32)(unsafe.Add(mBase, uint32(v157)+12))
	v160 = *(*float32)(unsafe.Add(mBase, uint32(v157)+8))
	v162 = *(*float32)(unsafe.Add(mBase, uint32(v157)+4))
	v164 = *(*float32)(unsafe.Add(mBase, uint32(v157)))
	v169 = base.F32_add(base.F32_mul(v158, v158), base.F32_add(base.F32_mul(v160, v160), base.F32_add(base.F32_mul(v162, v162), base.F32_add(base.F32_mul(v164, v164), v149))))
	v170 = int32(4)
	v171 = v134 + v170
	v173 = v137 + v170
	if v173 != v43&int32(2147483644) {
		v134 = v171
		v137 = v173
		v149 = v169
		goto L32
	} else {
		goto L34
	}
L33:
	;
	if v126 == int32(0) {
		v245 = v169
		goto L28
	} else {
		goto L35
	}
L34:
	;
	goto L33
L35:
	;
	v177 = v171
	v192 = v169
	goto L31
L36:
	;
	v222 = *(*float32)(unsafe.Add(mBase, uint32(v50+v198<<(uint(int32(2))%32))))
	v224 = base.F32_add(base.F32_mul(v222, v222), v213)
	v225 = int32(1)
	v228 = v199 + v225
	if v228 != v126 {
		v198 = v198 + v225
		v199 = v228
		v213 = v224
		goto L36
	} else {
		goto L38
	}
L37:
	;
	v245 = v224
	goto L28
L38:
	;
	goto L37
L39:
	;
	v277 = v39 & int32(3)
	v278 = float32(0)
	v279 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v39) {
		goto L43
	} else {
		goto L44
	}
L40:
	;
	v423 = v19
	goto L41
L41:
	;
	m.G0 = v24 + int32(16)
	v428 = float64(1)
	v432 = base.F64_div(v273, base.F64_sqrt(base.F64_mul(v423, v271)))
	if base.F64_gt(v432, v428) != 0 {
		v440 = v428
		goto L53
	} else {
		goto L54
	}
L42:
	;
	v423 = base.F64_promote_f32(v398)
	goto L41
L43:
	;
	v286 = v279
	v289 = int32(0)
	v302 = v278
	goto L46
L44:
	;
	v329 = v279
	v345 = v278
	goto L45
L45:
	;
	v350 = v329
	v351 = v279
	v366 = v345
	goto L50
L46:
	;
	v309 = v42 + v286<<(uint(int32(2))%32)
	v310 = *(*float32)(unsafe.Add(mBase, uint32(v309)+12))
	v312 = *(*float32)(unsafe.Add(mBase, uint32(v309)+8))
	v314 = *(*float32)(unsafe.Add(mBase, uint32(v309)+4))
	v316 = *(*float32)(unsafe.Add(mBase, uint32(v309)))
	v321 = base.F32_add(base.F32_mul(v310, v310), base.F32_add(base.F32_mul(v312, v312), base.F32_add(base.F32_mul(v314, v314), base.F32_add(base.F32_mul(v316, v316), v302))))
	v322 = int32(4)
	v323 = v286 + v322
	v325 = v289 + v322
	if v325 != v39&int32(2147483644) {
		v286 = v323
		v289 = v325
		v302 = v321
		goto L46
	} else {
		goto L48
	}
L47:
	;
	if v277 == int32(0) {
		v398 = v321
		goto L42
	} else {
		goto L49
	}
L48:
	;
	goto L47
L49:
	;
	v329 = v323
	v345 = v321
	goto L45
L50:
	;
	v374 = *(*float32)(unsafe.Add(mBase, uint32(v42+v350<<(uint(int32(2))%32))))
	v376 = base.F32_add(base.F32_mul(v374, v374), v366)
	v377 = int32(1)
	v380 = v351 + v377
	if v380 != v277 {
		v350 = v350 + v377
		v351 = v380
		v366 = v376
		goto L50
	} else {
		goto L52
	}
L51:
	;
	v398 = v376
	goto L42
L52:
	;
	goto L51
L53:
	;
	return base.I64_reinterpret_f64(base.F64_sub(v428, v440))
L54:
	;
	if base.F64_lt(v432, float64(-1)) == int32(0) {
		v440 = v432
		goto L53
	} else {
		goto L55
	}
L55:
	;
	v440 = float64(-1)
	goto L53
L56:
	;
	F_errcode(m, int32(130))
	mBase = m.M
	v450 = m.ExcPending
	if v450 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	v451 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
	v452 = *(*int32)(unsafe.Add(mBase, uint32(v32)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+4)) = v452
	*(*int32)(unsafe.Add(mBase, uint32(v24))) = v451
	F_errmsg(m, int32(_a_F_sparsevec_cosine_distance_0), v24)
	mBase = m.M
	v457 = m.ExcPending
	if v457 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	F_errfinish(m, int32(_a_F_sparsevec_cosine_distance_1), int32(50), int32(_a_F_sparsevec_cosine_distance_2))
	mBase = m.M
	v462 = m.ExcPending
	if v462 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_sparsevec_gt(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v64 float32
	_ = v64
	var v67 int32
	_ = v67
	var v74 float32
	_ = v74
	var v77 int32
	_ = v77
	var v79 float32
	_ = v79
	var v81 float32
	_ = v81
	var v87 int32
	_ = v87
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v114 float32
	_ = v114
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v130 float32
	_ = v130
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v155 int32
	_ = v155
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = F_pg_detoast_datum(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v8 = F_pg_detoast_datum(m, v7)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v23 = int32(16)
	v24 = v8 + v23
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
	v26 = int32(2)
	v28 = v24 + v25<<(uint(v26)%32)
	v30 = v3 + v23
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v3)+8))
	v34 = v30 + v31<<(uint(v26)%32)
	if v31 < v25 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	return base.I64_extend_i32_u(base.B2i32(int32(0) < v155))
L5:
	;
	v36 = v31
	goto L7
L6:
	;
	v36 = v25
	goto L7
L7:
	;
	if int32(0) < v36 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v42 = int32(0)
	goto L11
L9:
	;
	goto L10
L10:
	;
	if v25 <= v31 {
		goto L32
	} else {
		goto L33
	}
L11:
	;
	v55 = v42 << (uint(int32(2)) % 32)
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v30+v55)))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v55+v24)))
	if v57 < v59 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	goto L10
L13:
	;
	v64 = *(*float32)(unsafe.Add(mBase, uint32(v55+v34)))
	if base.F32_lt(v64, float32(0)) != 0 {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	goto L15
L15:
	;
	if v59 < v57 {
		goto L19
	} else {
		goto L20
	}
L16:
	;
	v67 = int32(-1)
	goto L18
L17:
	;
	v67 = int32(1)
	goto L18
L18:
	;
	v155 = v67
	goto L4
L19:
	;
	v74 = *(*float32)(unsafe.Add(mBase, uint32(v28+v42<<(uint(int32(2))%32))))
	if base.F32_lt(v74, float32(0)) != 0 {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	goto L21
L21:
	;
	v79 = *(*float32)(unsafe.Add(mBase, uint32(v55+v34)))
	v81 = *(*float32)(unsafe.Add(mBase, uint32(v55+v28)))
	if base.F32_lt(v79, v81) != 0 {
		goto L25
	} else {
		goto L26
	}
L22:
	;
	v77 = int32(1)
	goto L24
L23:
	;
	v77 = int32(-1)
	goto L24
L24:
	;
	v155 = v77
	goto L4
L25:
	;
	v155 = int32(-1)
	goto L4
L26:
	;
	goto L27
L27:
	;
	if base.F32_gt(v79, v81) != 0 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v155 = int32(1)
	goto L4
L29:
	;
	goto L30
L30:
	;
	v87 = v42 + int32(1)
	if v87 != v36 {
		v42 = v87
		goto L11
	} else {
		goto L31
	}
L31:
	;
	goto L12
L32:
	;
	if v31 <= v25 {
		goto L39
	} else {
		goto L40
	}
L33:
	;
	v106 = v31 << (uint(int32(2)) % 32)
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v24+v106)))
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v3)+4))
	if v109 <= v108 {
		goto L32
	} else {
		goto L34
	}
L34:
	;
	v114 = *(*float32)(unsafe.Add(mBase, uint32(v106+v28)))
	if base.F32_lt(v114, float32(0)) != 0 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v117 = int32(1)
	goto L37
L36:
	;
	v117 = int32(-1)
	goto L37
L37:
	;
	v155 = v117
	goto L4
L38:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v3)+4))
	if v136 < v134 {
		goto L46
	} else {
		goto L47
	}
L39:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	v134 = v120
	goto L38
L40:
	;
	goto L41
L41:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	v123 = v36 << (uint(int32(2)) % 32)
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v30+v123)))
	if v121 <= v125 {
		v134 = v121
		goto L38
	} else {
		goto L42
	}
L42:
	;
	v130 = *(*float32)(unsafe.Add(mBase, uint32(v123+v34)))
	if base.F32_lt(v130, float32(0)) != 0 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v133 = int32(-1)
	goto L45
L44:
	;
	v133 = int32(1)
	goto L45
L45:
	;
	v155 = v133
	goto L4
L46:
	;
	v155 = int32(-1)
	goto L4
L47:
	;
	goto L48
L48:
	;
	v155 = base.B2i32(v134 < v136)
	goto L4
}
func F_sparsevec_in(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int64
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v57 int32
	_ = v57
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v111 int32
	_ = v111
	var v121 int32
	_ = v121
	var v128 int64
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v146 int32
	_ = v146
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v192 float32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v212 int32
	_ = v212
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v250 int32
	_ = v250
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v267 int32
	_ = v267
	var v272 int32
	_ = v272
	var v289 int32
	_ = v289
	var v292 int32
	_ = v292
	var v298 int32
	_ = v298
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v311 int32
	_ = v311
	var v318 int32
	_ = v318
	var v332 int32
	_ = v332
	var v342 int32
	_ = v342
	var v346 int32
	_ = v346
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v373 int64
	_ = v373
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v377 int32
	_ = v377
	var v391 int32
	_ = v391
	var v401 int32
	_ = v401
	var v403 int32
	_ = v403
	var v411 int32
	_ = v411
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v430 int32
	_ = v430
	var v434 int32
	_ = v434
	var v438 int32
	_ = v438
	var v451 int32
	_ = v451
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v459 float32
	_ = v459
	var v462 int32
	_ = v462
	var v464 int32
	_ = v464
	var v486 int32
	_ = v486
	var v489 int32
	_ = v489
	var v494 int32
	_ = v494
	var v499 int32
	_ = v499
	var v503 int32
	_ = v503
	var v506 int32
	_ = v506
	var v512 int32
	_ = v512
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v521 int32
	_ = v521
	var v525 int32
	_ = v525
	var v528 int32
	_ = v528
	var v534 int32
	_ = v534
	var v539 int32
	_ = v539
	var v543 int32
	_ = v543
	var v546 int32
	_ = v546
	var v552 int32
	_ = v552
	var v557 int32
	_ = v557
	var v561 int32
	_ = v561
	var v564 int32
	_ = v564
	var v570 int32
	_ = v570
	var v575 int32
	_ = v575
	var v579 int32
	_ = v579
	var v582 int32
	_ = v582
	var v588 int32
	_ = v588
	var v593 int32
	_ = v593
	var v597 int32
	_ = v597
	var v600 int32
	_ = v600
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v609 int32
	_ = v609
	var v614 int32
	_ = v614
	var v618 int32
	_ = v618
	var v621 int32
	_ = v621
	var v627 int32
	_ = v627
	var v630 int32
	_ = v630
	var v631 int32
	_ = v631
	var v636 int32
	_ = v636
	var v640 int32
	_ = v640
	var v643 int32
	_ = v643
	var v649 int32
	_ = v649
	var v654 int32
	_ = v654
	var v658 int32
	_ = v658
	var v661 int32
	_ = v661
	var v667 int32
	_ = v667
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	var v676 int32
	_ = v676
	var v680 int32
	_ = v680
	var v683 int32
	_ = v683
	var v690 int32
	_ = v690
	var v695 int32
	_ = v695
	var v698 int32
	_ = v698
	v2 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(208)
	m.G0 = v15
	v17 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v20 = v18
	v24 = int32(1)
	goto L1
L1:
	;
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
	if v32 != int32(44) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v20 = v20 + int32(1)
	v24 = v698
	goto L1
L4:
	;
	if v32 != 0 {
		v698 = v24
		goto L3
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	v698 = v24 + int32(1)
	goto L3
L7:
	;
	if v24 < int32(_a_F_sparsevec_in_0) {
		goto L18
	} else {
		goto L19
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v680 = m.ExcPending
	if v680 != 0 {
		goto L21
	} else {
		goto L160
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v658 = m.ExcPending
	if v658 != 0 {
		goto L21
	} else {
		goto L155
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v640 = m.ExcPending
	if v640 != 0 {
		goto L21
	} else {
		goto L151
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v618 = m.ExcPending
	if v618 != 0 {
		goto L21
	} else {
		goto L146
	}
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v597 = m.ExcPending
	if v597 != 0 {
		goto L21
	} else {
		goto L141
	}
L13:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v579 = m.ExcPending
	if v579 != 0 {
		goto L21
	} else {
		goto L137
	}
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v561 = m.ExcPending
	if v561 != 0 {
		goto L21
	} else {
		goto L133
	}
L15:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v543 = m.ExcPending
	if v543 != 0 {
		goto L21
	} else {
		goto L129
	}
L16:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v525 = m.ExcPending
	if v525 != 0 {
		goto L21
	} else {
		goto L125
	}
L17:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v503 = m.ExcPending
	if v503 != 0 {
		goto L21
	} else {
		goto L120
	}
L18:
	;
	v37 = base.I32_wrap_i64(v17)
	v39 = F_palloc_mul(m, int32(8), v24)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	goto L20
L20:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v486 = m.ExcPending
	if v486 != 0 {
		goto L21
	} else {
		goto L116
	}
L21:
	;
	return int64(0)
L22:
	;
	v44 = v18
	goto L23
L23:
	;
	v57 = int32(*(*int8)(unsafe.Add(mBase, uint32(v44))))
	goto L25
L24:
	;
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44))))
	if v67 != int32(123) {
		goto L17
	} else {
		goto L27
	}
L25:
	;
	if base.B2i32(v57 == int32(32))|base.B2i32(base.Ui32((v57-int32(9))&int32(255)) < base.Ui32(int32(5))) != 0 {
		v44 = v44 + int32(1)
		goto L23
	} else {
		goto L26
	}
L26:
	;
	goto L24
L27:
	;
	v70 = v44
	goto L28
L28:
	;
	v83 = v70 + int32(1)
	v84 = int32(*(*int8)(unsafe.Add(mBase, uint32(v70)+1)))
	goto L30
L29:
	;
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83))))
	if v94 != int32(125) {
		goto L33
	} else {
		goto L34
	}
L30:
	;
	if base.B2i32(v84 == int32(32))|base.B2i32(base.Ui32((v84-int32(9))&int32(255)) < base.Ui32(int32(5))) != 0 {
		v70 = v83
		goto L28
	} else {
		goto L31
	}
L31:
	;
	goto L29
L32:
	;
	v318 = v306
	goto L85
L33:
	;
	if v24 != 0 {
		goto L36
	} else {
		goto L37
	}
L34:
	;
	goto L35
L35:
	;
	v306 = v70 + int32(2)
	v311 = v2
	goto L32
L36:
	;
	v97 = v83
	v102 = v2
	goto L39
L37:
	;
	goto L38
L38:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L21
	} else {
		goto L81
	}
L39:
	;
	v111 = int32(*(*int8)(unsafe.Add(mBase, uint32(v97))))
	goto L41
L40:
	;
	goto L38
L41:
	;
	if base.B2i32(v111 == int32(32))|base.B2i32(base.Ui32((v111-int32(9))&int32(255)) < base.Ui32(int32(5))) != 0 {
		v97 = v97 + int32(1)
		goto L39
	} else {
		goto L42
	}
L42:
	;
	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97))))
	if v121 == int32(0) {
		goto L16
	} else {
		goto L43
	}
L43:
	;
	v128 = F_strtox_2(m, v97, v15+int32(204), int32(10), int64(2147483648))
	mBase = m.M
	v129 = base.I32_wrap_i64(v128)
	goto L44
L44:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v15)+204))
	if v97 == v130 {
		goto L15
	} else {
		goto L45
	}
L45:
	;
	v133 = v130
	goto L46
L46:
	;
	v146 = int32(*(*int8)(unsafe.Add(mBase, uint32(v133))))
	goto L48
L47:
	;
	v156 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v133))))
	if v156 != int32(58) {
		goto L14
	} else {
		goto L50
	}
L48:
	;
	if base.B2i32(v146 == int32(32))|base.B2i32(base.Ui32((v146-int32(9))&int32(255)) < base.Ui32(int32(5))) != 0 {
		v133 = v133 + int32(1)
		goto L46
	} else {
		goto L49
	}
L49:
	;
	goto L47
L50:
	;
	v159 = int32(-2147483647)
	if v129 <= v159 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v162 = v159
	goto L53
L52:
	;
	v162 = v129
	goto L53
L53:
	;
	v163 = v133
	goto L54
L54:
	;
	v175 = int32(*(*int8)(unsafe.Add(mBase, uint32(v163)+1)))
	v177 = v163 + int32(1)
	goto L56
L55:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_sparsevec_in[0])) = int32(0)
	v192 = F_strtof(m, v177, v15+int32(204))
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L21
	} else {
		goto L58
	}
L56:
	;
	if base.B2i32(v175 == int32(32))|base.B2i32(base.Ui32((v175-int32(9))&int32(255)) < base.Ui32(int32(5))) != 0 {
		v163 = v177
		goto L54
	} else {
		goto L57
	}
L57:
	;
	goto L55
L58:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v15)+204))
	if v194 == v177 {
		goto L13
	} else {
		goto L59
	}
L59:
	;
	v197 = *(*int32)(unsafe.Add(mBase, _c_F_sparsevec_in[0]))
	if v197 == int32(68) {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v202 = base.I32_reinterpret_f32(v192) & int32(2147483647)
	v203 = int32(0)
	if base.B2i32(v202 != v203)&base.B2i32(v202 != int32(2139095040)) == v203 {
		goto L12
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	F_CheckElement_2(m, v192)
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L21
	} else {
		goto L64
	}
L63:
	;
	goto L62
L64:
	;
	if base.F32_ne(v192, float32(0)) != 0 {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v217 = v39 + v102<<(uint(int32(3))%32)
	*(*float32)(unsafe.Add(mBase, uint32(v217)+4)) = v192
	v219 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v217))) = v162 - v219
	v225 = v102 + v219
	goto L67
L66:
	;
	v225 = v102
	goto L67
L67:
	;
	v226 = v194
	goto L68
L68:
	;
	v239 = v226 + int32(1)
	v240 = int32(*(*int8)(unsafe.Add(mBase, uint32(v226))))
	goto L70
L69:
	;
	v250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v226))))
	if v250 != int32(44) {
		goto L72
	} else {
		goto L73
	}
L70:
	;
	if base.B2i32(v240 == int32(32))|base.B2i32(base.Ui32((v240-int32(9))&int32(255)) < base.Ui32(int32(5))) != 0 {
		v226 = v239
		goto L68
	} else {
		goto L71
	}
L71:
	;
	goto L69
L72:
	;
	if v250 == int32(125) {
		v306 = v239
		v311 = v225
		goto L32
	} else {
		goto L75
	}
L73:
	;
	goto L74
L74:
	;
	if v24 != v225 {
		v97 = v239
		v102 = v225
		goto L39
	} else {
		goto L80
	}
L75:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L21
	} else {
		goto L76
	}
L76:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L21
	} else {
		goto L77
	}
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+160)) = v18
	F_errmsg(m, int32(_a_F_sparsevec_in_1), v15+int32(160))
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L21
	} else {
		goto L78
	}
L78:
	;
	F_errfinish(m, int32(_a_F_sparsevec_in_2), int32(345), int32(_a_F_sparsevec_in_3))
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L21
	} else {
		goto L79
	}
L79:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L80:
	;
	goto L40
L81:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L21
	} else {
		goto L82
	}
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+80)) = v18
	F_errmsg(m, int32(_a_F_sparsevec_in_4), v15+int32(80))
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L21
	} else {
		goto L83
	}
L83:
	;
	F_errfinish(m, int32(_a_F_sparsevec_in_2), int32(262), int32(_a_F_sparsevec_in_3))
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L21
	} else {
		goto L84
	}
L84:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L85:
	;
	v332 = int32(*(*int8)(unsafe.Add(mBase, uint32(v318))))
	goto L87
L86:
	;
	v342 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v318))))
	if v342 != int32(47) {
		goto L11
	} else {
		goto L89
	}
L87:
	;
	if base.B2i32(v332 == int32(32))|base.B2i32(base.Ui32((v332-int32(9))&int32(255)) < base.Ui32(int32(5))) != 0 {
		v318 = v318 + int32(1)
		goto L85
	} else {
		goto L88
	}
L88:
	;
	goto L86
L89:
	;
	v346 = v318
	goto L90
L90:
	;
	v357 = int32(*(*int8)(unsafe.Add(mBase, uint32(v346)+1)))
	v359 = v346 + int32(1)
	goto L92
L91:
	;
	v373 = F_strtox_2(m, v359, v15+int32(204), int32(10), int64(2147483648))
	mBase = m.M
	v374 = base.I32_wrap_i64(v373)
	goto L94
L92:
	;
	if base.B2i32(v357 == int32(32))|base.B2i32(base.Ui32((v357-int32(9))&int32(255)) < base.Ui32(int32(5))) != 0 {
		v346 = v359
		goto L90
	} else {
		goto L93
	}
L93:
	;
	goto L91
L94:
	;
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v15)+204))
	if v359 == v375 {
		goto L10
	} else {
		goto L95
	}
L95:
	;
	v377 = v375
	goto L96
L96:
	;
	v391 = int32(*(*int8)(unsafe.Add(mBase, uint32(v377))))
	goto L98
L97:
	;
	v401 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v377))))
	if v401 != 0 {
		goto L9
	} else {
		goto L100
	}
L98:
	;
	if base.B2i32(v391 == int32(32))|base.B2i32(base.Ui32((v391-int32(9))&int32(255)) < base.Ui32(int32(5))) != 0 {
		v377 = v377 + int32(1)
		goto L96
	} else {
		goto L99
	}
L99:
	;
	goto L97
L100:
	;
	F_CheckDim_2(m, v374)
	mBase = m.M
	v403 = m.ExcPending
	if v403 != 0 {
		goto L21
	} else {
		goto L101
	}
L101:
	;
	if base.B2i32(v37 != int32(-1))&base.B2i32(v374 != v37) != 0 {
		goto L8
	} else {
		goto L102
	}
L102:
	;
	F_pg_qsort(m, v39, v311, int32(8), int32(_a_F_sparsevec_in_5))
	mBase = m.M
	v411 = m.ExcPending
	if v411 != 0 {
		goto L21
	} else {
		goto L103
	}
L103:
	;
	v414 = F_mul_size(m, int32(4), v311)
	mBase = m.M
	v415 = m.ExcPending
	if v415 != 0 {
		goto L21
	} else {
		goto L104
	}
L104:
	;
	v416 = F_add_size(m, int32(16), v414)
	mBase = m.M
	v417 = m.ExcPending
	if v417 != 0 {
		goto L21
	} else {
		goto L105
	}
L105:
	;
	v419 = F_mul_size(m, int32(4), v311)
	mBase = m.M
	v420 = m.ExcPending
	if v420 != 0 {
		goto L21
	} else {
		goto L106
	}
L106:
	;
	v421 = F_add_size(m, v416, v419)
	mBase = m.M
	v422 = m.ExcPending
	if v422 != 0 {
		goto L21
	} else {
		goto L107
	}
L107:
	;
	v423 = F_palloc0(m, v421)
	mBase = m.M
	v424 = m.ExcPending
	if v424 != 0 {
		goto L21
	} else {
		goto L108
	}
L108:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v423)+8)) = v311
	*(*int32)(unsafe.Add(mBase, uint32(v423)+4)) = v374
	*(*int32)(unsafe.Add(mBase, uint32(v423))) = v421 << (uint(int32(2)) % 32)
	v430 = int32(0)
	if v430 < v311 {
		goto L109
	} else {
		goto L110
	}
L109:
	;
	v434 = v423 + int32(16)
	v438 = v430
	goto L112
L110:
	;
	goto L111
L111:
	;
	m.G0 = v15 + int32(208)
	return base.I64_extend_i32_u(v423)
L112:
	;
	v451 = v438 << (uint(int32(2)) % 32)
	v455 = v39 + v438<<(uint(int32(3))%32)
	v456 = *(*int32)(unsafe.Add(mBase, uint32(v455)))
	*(*int32)(unsafe.Add(mBase, uint32(v434+v451))) = v456
	v459 = *(*float32)(unsafe.Add(mBase, uint32(v455)+4))
	*(*float32)(unsafe.Add(mBase, uint32(v451+(v434+v311<<(uint(int32(2))%32))))) = v459
	F_CheckIndex(m, v434, v438, v374)
	mBase = m.M
	v462 = m.ExcPending
	if v462 != 0 {
		goto L21
	} else {
		goto L114
	}
L113:
	;
	goto L111
L114:
	;
	v464 = v438 + int32(1)
	if v464 != v311 {
		v438 = v464
		goto L112
	} else {
		goto L115
	}
L115:
	;
	goto L113
L116:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v489 = m.ExcPending
	if v489 != 0 {
		goto L21
	} else {
		goto L117
	}
L117:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = int32(_a_F_sparsevec_in_6)
	F_errmsg(m, int32(_a_F_sparsevec_in_7), v15)
	mBase = m.M
	v494 = m.ExcPending
	if v494 != 0 {
		goto L21
	} else {
		goto L118
	}
L118:
	;
	F_errfinish(m, int32(_a_F_sparsevec_in_2), int32(230), int32(_a_F_sparsevec_in_3))
	mBase = m.M
	v499 = m.ExcPending
	if v499 != 0 {
		goto L21
	} else {
		goto L119
	}
L119:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L120:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v506 = m.ExcPending
	if v506 != 0 {
		goto L21
	} else {
		goto L121
	}
L121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+192)) = v18
	F_errmsg(m, int32(_a_F_sparsevec_in_1), v15+int32(192))
	mBase = m.M
	v512 = m.ExcPending
	if v512 != 0 {
		goto L21
	} else {
		goto L122
	}
L122:
	;
	v515 = F_errdetail(m, int32(_a_F_sparsevec_in_8), int32(0))
	mBase = m.M
	v516 = m.ExcPending
	if v516 != 0 {
		goto L21
	} else {
		goto L123
	}
L123:
	;
	F_errfinish(m, int32(_a_F_sparsevec_in_2), int32(243), int32(_a_F_sparsevec_in_3))
	mBase = m.M
	v521 = m.ExcPending
	if v521 != 0 {
		goto L21
	} else {
		goto L124
	}
L124:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L125:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v528 = m.ExcPending
	if v528 != 0 {
		goto L21
	} else {
		goto L126
	}
L126:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+96)) = v18
	F_errmsg(m, int32(_a_F_sparsevec_in_1), v15+int32(96))
	mBase = m.M
	v534 = m.ExcPending
	if v534 != 0 {
		goto L21
	} else {
		goto L127
	}
L127:
	;
	F_errfinish(m, int32(_a_F_sparsevec_in_2), int32(271), int32(_a_F_sparsevec_in_3))
	mBase = m.M
	v539 = m.ExcPending
	if v539 != 0 {
		goto L21
	} else {
		goto L128
	}
L128:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L129:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v546 = m.ExcPending
	if v546 != 0 {
		goto L21
	} else {
		goto L130
	}
L130:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+112)) = v18
	F_errmsg(m, int32(_a_F_sparsevec_in_1), v15+int32(112))
	mBase = m.M
	v552 = m.ExcPending
	if v552 != 0 {
		goto L21
	} else {
		goto L131
	}
L131:
	;
	F_errfinish(m, int32(_a_F_sparsevec_in_2), int32(279), int32(_a_F_sparsevec_in_3))
	mBase = m.M
	v557 = m.ExcPending
	if v557 != 0 {
		goto L21
	} else {
		goto L132
	}
L132:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L133:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v564 = m.ExcPending
	if v564 != 0 {
		goto L21
	} else {
		goto L134
	}
L134:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+176)) = v18
	F_errmsg(m, int32(_a_F_sparsevec_in_1), v15+int32(176))
	mBase = m.M
	v570 = m.ExcPending
	if v570 != 0 {
		goto L21
	} else {
		goto L135
	}
L135:
	;
	F_errfinish(m, int32(_a_F_sparsevec_in_2), int32(295), int32(_a_F_sparsevec_in_3))
	mBase = m.M
	v575 = m.ExcPending
	if v575 != 0 {
		goto L21
	} else {
		goto L136
	}
L136:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L137:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v582 = m.ExcPending
	if v582 != 0 {
		goto L21
	} else {
		goto L138
	}
L138:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+128)) = v18
	F_errmsg(m, int32(_a_F_sparsevec_in_1), v15+int32(128))
	mBase = m.M
	v588 = m.ExcPending
	if v588 != 0 {
		goto L21
	} else {
		goto L139
	}
L139:
	;
	F_errfinish(m, int32(_a_F_sparsevec_in_2), int32(311), int32(_a_F_sparsevec_in_3))
	mBase = m.M
	v593 = m.ExcPending
	if v593 != 0 {
		goto L21
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
	F_errcode(m, int32(50331778))
	mBase = m.M
	v600 = m.ExcPending
	if v600 != 0 {
		goto L21
	} else {
		goto L142
	}
L142:
	;
	v602 = F_pnstrdup(m, v177, v194-v177)
	mBase = m.M
	v603 = m.ExcPending
	if v603 != 0 {
		goto L21
	} else {
		goto L143
	}
L143:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+144)) = v602
	F_errmsg(m, int32(_a_F_sparsevec_in_9), v15+int32(144))
	mBase = m.M
	v609 = m.ExcPending
	if v609 != 0 {
		goto L21
	} else {
		goto L144
	}
L144:
	;
	F_errfinish(m, int32(_a_F_sparsevec_in_2), int32(317), int32(_a_F_sparsevec_in_3))
	mBase = m.M
	v614 = m.ExcPending
	if v614 != 0 {
		goto L21
	} else {
		goto L145
	}
L145:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L146:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v621 = m.ExcPending
	if v621 != 0 {
		goto L21
	} else {
		goto L147
	}
L147:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+64)) = v18
	F_errmsg(m, int32(_a_F_sparsevec_in_1), v15-int32(-64))
	mBase = m.M
	v627 = m.ExcPending
	if v627 != 0 {
		goto L21
	} else {
		goto L148
	}
L148:
	;
	v630 = F_errdetail(m, int32(_a_F_sparsevec_in_10), int32(0))
	mBase = m.M
	v631 = m.ExcPending
	if v631 != 0 {
		goto L21
	} else {
		goto L149
	}
L149:
	;
	F_errfinish(m, int32(_a_F_sparsevec_in_2), int32(356), int32(_a_F_sparsevec_in_3))
	mBase = m.M
	v636 = m.ExcPending
	if v636 != 0 {
		goto L21
	} else {
		goto L150
	}
L150:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L151:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v643 = m.ExcPending
	if v643 != 0 {
		goto L21
	} else {
		goto L152
	}
L152:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v18
	F_errmsg(m, int32(_a_F_sparsevec_in_1), v15+int32(16))
	mBase = m.M
	v649 = m.ExcPending
	if v649 != 0 {
		goto L21
	} else {
		goto L153
	}
L153:
	;
	F_errfinish(m, int32(_a_F_sparsevec_in_2), int32(369), int32(_a_F_sparsevec_in_3))
	mBase = m.M
	v654 = m.ExcPending
	if v654 != 0 {
		goto L21
	} else {
		goto L154
	}
L154:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L155:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v661 = m.ExcPending
	if v661 != 0 {
		goto L21
	} else {
		goto L156
	}
L156:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+48)) = v18
	F_errmsg(m, int32(_a_F_sparsevec_in_1), v15+int32(48))
	mBase = m.M
	v667 = m.ExcPending
	if v667 != 0 {
		goto L21
	} else {
		goto L157
	}
L157:
	;
	v670 = F_errdetail(m, int32(_a_F_sparsevec_in_11), int32(0))
	mBase = m.M
	v671 = m.ExcPending
	if v671 != 0 {
		goto L21
	} else {
		goto L158
	}
L158:
	;
	F_errfinish(m, int32(_a_F_sparsevec_in_2), int32(387), int32(_a_F_sparsevec_in_3))
	mBase = m.M
	v676 = m.ExcPending
	if v676 != 0 {
		goto L21
	} else {
		goto L159
	}
L159:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L160:
	;
	F_errcode(m, int32(130))
	mBase = m.M
	v683 = m.ExcPending
	if v683 != 0 {
		goto L21
	} else {
		goto L161
	}
L161:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+36)) = v374
	*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v37
	F_errmsg(m, int32(_a_F_sparsevec_in_12), v15+int32(32))
	mBase = m.M
	v690 = m.ExcPending
	if v690 != 0 {
		goto L21
	} else {
		goto L162
	}
L162:
	;
	F_errfinish(m, int32(_a_F_sparsevec_in_2), int32(62), int32(_a_F_sparsevec_in_13))
	mBase = m.M
	v695 = m.ExcPending
	if v695 != 0 {
		goto L21
	} else {
		goto L163
	}
L163:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_sparsevec_l1_distance(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v16 float32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
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
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v55 int32
	_ = v55
	var v61 float32
	_ = v61
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v84 float32
	_ = v84
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v90 float32
	_ = v90
	var v92 float32
	_ = v92
	var v98 float32
	_ = v98
	var v101 float32
	_ = v101
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v123 float32
	_ = v123
	var v126 float32
	_ = v126
	var v129 float32
	_ = v129
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v148 float32
	_ = v148
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v171 float32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v177 float32
	_ = v177
	var v179 float32
	_ = v179
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v198 float32
	_ = v198
	var v202 int32
	_ = v202
	var v217 float32
	_ = v217
	var v220 int32
	_ = v220
	var v221 float32
	_ = v221
	var v223 float32
	_ = v223
	var v225 float32
	_ = v225
	var v227 float32
	_ = v227
	var v232 float32
	_ = v232
	var v234 int32
	_ = v234
	var v251 float32
	_ = v251
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v271 int32
	_ = v271
	var v276 int32
	_ = v276
	v2 = int32(0)
	v16 = float32(0)
	v17 = m.G0
	v19 = v17 - int32(16)
	m.G0 = v19
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v22 = F_pg_detoast_datum(m, v21)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v27 = F_pg_detoast_datum(m, v26)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
	if v29 == v30 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v33 = v27 + int32(16)
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v27)+8))
	v37 = v33 + v34<<(uint(int32(2))%32)
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
	if int32(0) < v38 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L1
	} else {
		goto L43
	}
L7:
	;
	v42 = v22 + int32(16)
	v45 = v42 + v38<<(uint(int32(2))%32)
	v47 = v2
	v55 = v2
	v61 = v16
	goto L10
L8:
	;
	v134 = v2
	v148 = v16
	goto L9
L9:
	;
	if v34 <= v134 {
		v251 = v148
		goto L30
	} else {
		goto L31
	}
L10:
	;
	v63 = v55 << (uint(int32(2)) % 32)
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v42+v63)))
	if v34 <= v47 {
		v109 = v47
		v111 = int32(-1)
		v123 = v61
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v134 = v109
	v148 = v129
	goto L9
L12:
	;
	if v65 != v111 {
		goto L26
	} else {
		goto L27
	}
L13:
	;
	v69 = v47
	v70 = v47
	v84 = v61
	goto L14
L14:
	;
	v86 = v69 << (uint(int32(2)) % 32)
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v33+v86)))
	if v88 == v65 {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	v109 = v105
	v111 = v88
	v123 = v101
	goto L12
L16:
	;
	v103 = v69 + int32(1)
	if v65 < v88 {
		goto L21
	} else {
		goto L22
	}
L17:
	;
	v90 = *(*float32)(unsafe.Add(mBase, uint32(v45+v63)))
	v92 = *(*float32)(unsafe.Add(mBase, uint32(v37+v86)))
	v101 = base.F32_add(base.F32_abs(base.F32_sub(v90, v92)), v84)
	goto L16
L18:
	;
	goto L19
L19:
	;
	if v65 <= v88 {
		v101 = v84
		goto L16
	} else {
		goto L20
	}
L20:
	;
	v98 = *(*float32)(unsafe.Add(mBase, uint32(v37+v86)))
	v101 = base.F32_add(base.F32_abs(v98), v84)
	goto L16
L21:
	;
	v105 = v70
	goto L23
L22:
	;
	v105 = v103
	goto L23
L23:
	;
	if v65 <= v88 {
		v109 = v105
		v111 = v88
		v123 = v101
		goto L12
	} else {
		goto L24
	}
L24:
	;
	if v103 != v34 {
		v69 = v103
		v70 = v105
		v84 = v101
		goto L14
	} else {
		goto L25
	}
L25:
	;
	goto L15
L26:
	;
	v126 = *(*float32)(unsafe.Add(mBase, uint32(v45+v63)))
	v129 = base.F32_add(base.F32_abs(v126), v123)
	goto L28
L27:
	;
	v129 = v123
	goto L28
L28:
	;
	v131 = v55 + int32(1)
	if v131 != v38 {
		v47 = v109
		v55 = v131
		v61 = v129
		goto L10
	} else {
		goto L29
	}
L29:
	;
	goto L11
L30:
	;
	m.G0 = v19 + int32(16)
	return base.I64_reinterpret_f64(base.F64_promote_f32(v251))
L31:
	;
	v152 = (v34 - v134) & int32(3)
	if v152 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	if base.Ui32(int32(-4)) < base.Ui32(v134-v34) {
		v251 = v198
		goto L30
	} else {
		goto L39
	}
L33:
	;
	v183 = v134
	v198 = v148
	goto L32
L34:
	;
	goto L35
L35:
	;
	v156 = v134
	v159 = int32(0)
	v171 = v148
	goto L36
L36:
	;
	v172 = int32(1)
	v173 = v156 + v172
	v177 = *(*float32)(unsafe.Add(mBase, uint32(v37+v156<<(uint(int32(2))%32))))
	v179 = base.F32_add(base.F32_abs(v177), v171)
	v181 = v159 + v172
	if v181 != v152 {
		v156 = v173
		v159 = v181
		v171 = v179
		goto L36
	} else {
		goto L38
	}
L37:
	;
	v183 = v173
	v198 = v179
	goto L32
L38:
	;
	goto L37
L39:
	;
	v202 = v183
	v217 = v198
	goto L40
L40:
	;
	v220 = v37 + v202<<(uint(int32(2))%32)
	v221 = *(*float32)(unsafe.Add(mBase, uint32(v220)+12))
	v223 = *(*float32)(unsafe.Add(mBase, uint32(v220)+8))
	v225 = *(*float32)(unsafe.Add(mBase, uint32(v220)+4))
	v227 = *(*float32)(unsafe.Add(mBase, uint32(v220)))
	v232 = base.F32_add(base.F32_abs(v221), base.F32_add(base.F32_abs(v223), base.F32_add(base.F32_abs(v225), base.F32_add(base.F32_abs(v227), v217))))
	v234 = v202 + int32(4)
	if v234 != v34 {
		v202 = v234
		v217 = v232
		goto L40
	} else {
		goto L42
	}
L41:
	;
	v251 = v232
	goto L30
L42:
	;
	goto L41
L43:
	;
	F_errcode(m, int32(130))
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = v266
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v265
	F_errmsg(m, int32(_a_F_sparsevec_l1_distance_0), v19)
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	F_errfinish(m, int32(_a_F_sparsevec_l1_distance_1), int32(50), int32(_a_F_sparsevec_l1_distance_2))
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_sparsevec_to_vector(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v82 float32
	_ = v82
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v92 float32
	_ = v92
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v102 float32
	_ = v102
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v112 float32
	_ = v112
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v133 int32
	_ = v133
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v153 float32
	_ = v153
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v188 int32
	_ = v188
	var v193 int32
	_ = v193
	v2 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(16)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v18 = F_pg_detoast_datum(m, v17)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		return int64(0)
	} else {
		v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v23 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
		v24 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
		F_CheckDim_3(m, v24)
		mBase = m.M
		v26 = m.ExcPending
		if v26 != 0 {
			return int64(0)
		} else {
			if base.B2i32(v22 != int32(-1))&base.B2i32(v24 != v22) == int32(0) {
				v35 = F_mul_size(m, int32(4), v24)
				mBase = m.M
				v36 = m.ExcPending
				if v36 != 0 {
					return int64(0)
				} else {
					v37 = F_add_size(m, int32(8), v35)
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
						return int64(0)
					} else {
						v39 = F_palloc0(m, v37)
						mBase = m.M
						v40 = m.ExcPending
						if v40 != 0 {
							return int64(0)
						} else {
							*(*uint16)(unsafe.Add(mBase, uint32(v39)+4)) = uint16(v24)
							*(*int32)(unsafe.Add(mBase, uint32(v39))) = v37 << (uint(int32(2)) % 32)
							v45 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
							if v45 <= int32(0) {
							} else {
								v49 = v18 + int32(16)
								v52 = v49 + v23<<(uint(int32(2))%32)
								v54 = v45 & int32(3)
								v56 = v39 + int32(8)
								v57 = int32(0)
								if base.Ui32(int32(4)) <= base.Ui32(v45) {
									v62 = v57
									v72 = v2
									for {
										v74 = int32(2)
										v75 = v62 << (uint(v74) % 32)
										v77 = *(*int32)(unsafe.Add(mBase, uint32(v49+v75)))
										v82 = *(*float32)(unsafe.Add(mBase, uint32(v75+v52)))
										*(*float32)(unsafe.Add(mBase, uint32(v56+v77<<(uint(v74)%32)))) = v82
										v84 = int32(4)
										v85 = v75 | v84
										v87 = *(*int32)(unsafe.Add(mBase, uint32(v49+v85)))
										v92 = *(*float32)(unsafe.Add(mBase, uint32(v52+v85)))
										*(*float32)(unsafe.Add(mBase, uint32(v56+v87<<(uint(v74)%32)))) = v92
										v95 = v75 | int32(8)
										v97 = *(*int32)(unsafe.Add(mBase, uint32(v49+v95)))
										v102 = *(*float32)(unsafe.Add(mBase, uint32(v52+v95)))
										*(*float32)(unsafe.Add(mBase, uint32(v56+v97<<(uint(v74)%32)))) = v102
										v105 = v75 | int32(12)
										v107 = *(*int32)(unsafe.Add(mBase, uint32(v49+v105)))
										v112 = *(*float32)(unsafe.Add(mBase, uint32(v105+v52)))
										*(*float32)(unsafe.Add(mBase, uint32(v56+v107<<(uint(v74)%32)))) = v112
										v115 = v62 + v84
										v117 = v72 + v84
										if v117 != v45&int32(2147483644) {
											v62 = v115
											v72 = v117
											continue
										} else {
											break
										}
										break
									}
									if v54 == int32(0) {
									} else {
										v121 = v115
										v133 = v121
										v144 = v2
										for {
											v145 = int32(2)
											v146 = v133 << (uint(v145) % 32)
											v148 = *(*int32)(unsafe.Add(mBase, uint32(v49+v146)))
											v153 = *(*float32)(unsafe.Add(mBase, uint32(v146+v52)))
											*(*float32)(unsafe.Add(mBase, uint32(v56+v148<<(uint(v145)%32)))) = v153
											v155 = int32(1)
											v158 = v144 + v155
											if v158 != v54 {
												v133 = v133 + v155
												v144 = v158
												continue
											} else {
												break
											}
											break
										}
									}
								} else {
									v121 = v57
									v133 = v121
									v144 = v2
									for {
										v145 = int32(2)
										v146 = v133 << (uint(v145) % 32)
										v148 = *(*int32)(unsafe.Add(mBase, uint32(v49+v146)))
										v153 = *(*float32)(unsafe.Add(mBase, uint32(v146+v52)))
										*(*float32)(unsafe.Add(mBase, uint32(v56+v148<<(uint(v145)%32)))) = v153
										v155 = int32(1)
										v158 = v144 + v155
										if v158 != v54 {
											v133 = v133 + v155
											v144 = v158
											continue
										} else {
											break
										}
										break
									}
								}
							}
							m.G0 = v15 + int32(16)
							return base.I64_extend_i32_u(v39)
						}
					}
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v180 = m.ExcPending
				if v180 != 0 {
					return int64(0)
				} else {
					F_errcode(m, int32(130))
					mBase = m.M
					v183 = m.ExcPending
					if v183 != 0 {
						return int64(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v24
						*(*int32)(unsafe.Add(mBase, uint32(v15))) = v22
						F_errmsg(m, int32(_a_F_sparsevec_to_vector_0), v15)
						mBase = m.M
						v188 = m.ExcPending
						if v188 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(_a_F_sparsevec_to_vector_1), int32(88), int32(_a_F_sparsevec_to_vector_2))
							mBase = m.M
							v193 = m.ExcPending
							if v193 != 0 {
								return int64(0)
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
