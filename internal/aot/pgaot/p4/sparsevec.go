package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_sparsevec_cosine_distance(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v16 float32
	_ = v16
	var v19 float64
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
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
	var v33 float64
	_ = v33
	var v35 int32
	_ = v35
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
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v64 int32
	_ = v64
	var v69 float32
	_ = v69
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v96 float32
	_ = v96
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
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v150 float32
	_ = v150
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
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v188 int32
	_ = v188
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v209 float32
	_ = v209
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v229 float32
	_ = v229
	var v237 float32
	_ = v237
	var v239 float32
	_ = v239
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v260 float32
	_ = v260
	var v286 float64
	_ = v286
	var v287 float64
	_ = v287
	var v291 int32
	_ = v291
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v318 float32
	_ = v318
	var v324 int32
	_ = v324
	var v325 float32
	_ = v325
	var v327 float32
	_ = v327
	var v329 float32
	_ = v329
	var v331 float32
	_ = v331
	var v336 float32
	_ = v336
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v342 int32
	_ = v342
	var v358 float32
	_ = v358
	var v362 int32
	_ = v362
	var v365 int32
	_ = v365
	var v378 float32
	_ = v378
	var v385 float32
	_ = v385
	var v387 float32
	_ = v387
	var v388 int32
	_ = v388
	var v391 int32
	_ = v391
	var v409 float32
	_ = v409
	var v434 float64
	_ = v434
	var v437 float64
	_ = v437
	var v445 float64
	_ = v445
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	v2 = int32(0)
	v16 = float32(0)
	v19 = float64(0)
	v21 = m.G0
	v23 = v21 - int32(16)
	m.G0 = v23
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v26 = F_pg_detoast_datum(m, v25)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v31 = F_pg_detoast_datum(m, v30)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v33 = float64(1)
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	if v35 == v36 {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	v447 = F_Float8GetDatum(m, base.F64_sub(v33, v445))
	mBase = m.M
	v448 = m.ExcPending
	if v448 != 0 {
		goto L1
	} else {
		goto L63
	}
L5:
	;
	if int32(0) < v40 {
		goto L45
	} else {
		goto L46
	}
L6:
	;
	if v126 != 0 {
		goto L39
	} else {
		goto L40
	}
L7:
	;
	v39 = v31 + int32(16)
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v31)+8))
	v43 = v39 + v40<<(uint(int32(2))%32)
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v26)+8))
	if v45 <= int32(0) {
		v286 = v19
		v287 = float64(0)
		goto L5
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L1
	} else {
		goto L35
	}
L10:
	;
	v49 = v26 + int32(16)
	v52 = v49 + v45<<(uint(int32(2))%32)
	v55 = v2
	v64 = v2
	v69 = v16
	goto L11
L11:
	;
	if v40 < v55 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v126 = v45 & int32(3)
	if base.Ui32(v45) < base.Ui32(int32(4)) {
		goto L29
	} else {
		goto L30
	}
L13:
	;
	v74 = v55
	goto L15
L14:
	;
	v74 = v40
	goto L15
L15:
	;
	v76 = v64 << (uint(int32(2)) % 32)
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v76+v49)))
	v80 = v55
	v82 = v55
	v96 = v69
	goto L16
L16:
	;
	if v80 != v74 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v123 = v64 + int32(1)
	if v123 != v45 {
		v55 = v118
		v64 = v123
		v69 = v121
		goto L11
	} else {
		goto L28
	}
L18:
	;
	v102 = v80 << (uint(int32(2)) % 32)
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v39+v102)))
	if v104 == v79 {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	v118 = v82
	v121 = v96
	goto L20
L20:
	;
	goto L17
L21:
	;
	v106 = *(*float32)(unsafe.Add(mBase, uint32(v52+v76)))
	v108 = *(*float32)(unsafe.Add(mBase, uint32(v102+v43)))
	v111 = base.F32_add(base.F32_mul(v106, v108), v96)
	goto L23
L22:
	;
	v111 = v96
	goto L23
L23:
	;
	v113 = v80 + int32(1)
	if v79 < v104 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v115 = v82
	goto L26
L25:
	;
	v115 = v113
	goto L26
L26:
	;
	if v104 < v79 {
		v80 = v113
		v82 = v115
		v96 = v111
		goto L16
	} else {
		goto L27
	}
L27:
	;
	v118 = v115
	v121 = v111
	goto L20
L28:
	;
	goto L12
L29:
	;
	v194 = int32(0)
	v209 = v16
	goto L6
L30:
	;
	goto L31
L31:
	;
	v133 = int32(0)
	v135 = v133
	v137 = v133
	v150 = v16
	goto L32
L32:
	;
	v157 = v52 + v135<<(uint(int32(2))%32)
	v158 = *(*float32)(unsafe.Add(mBase, uint32(v157)+12))
	v160 = *(*float32)(unsafe.Add(mBase, uint32(v157)+8))
	v162 = *(*float32)(unsafe.Add(mBase, uint32(v157)+4))
	v164 = *(*float32)(unsafe.Add(mBase, uint32(v157)))
	v169 = base.F32_add(base.F32_mul(v158, v158), base.F32_add(base.F32_mul(v160, v160), base.F32_add(base.F32_mul(v162, v162), base.F32_add(base.F32_mul(v164, v164), v150))))
	v170 = int32(4)
	v171 = v135 + v170
	v173 = v137 + v170
	if v173 != v45&int32(2147483644) {
		v135 = v171
		v137 = v173
		v150 = v169
		goto L32
	} else {
		goto L34
	}
L33:
	;
	v194 = v171
	v209 = v169
	goto L6
L34:
	;
	goto L33
L35:
	;
	F_errcode(m, int32(130))
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+4)) = v183
	*(*int32)(unsafe.Add(mBase, uint32(v23))) = v182
	F_errmsg(m, int32(499677), v23)
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	F_errfinish(m, int32(524189), int32(50), int32(159834))
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L39:
	;
	v214 = v194
	v217 = int32(0)
	v229 = v209
	goto L42
L40:
	;
	v260 = v209
	goto L41
L41:
	;
	v286 = base.F64_promote_f32(v260)
	v287 = base.F64_promote_f32(v121)
	goto L5
L42:
	;
	v237 = *(*float32)(unsafe.Add(mBase, uint32(v52+v214<<(uint(int32(2))%32))))
	v239 = base.F32_add(base.F32_mul(v237, v237), v229)
	v240 = int32(1)
	v243 = v217 + v240
	if v243 != v126 {
		v214 = v214 + v240
		v217 = v243
		v229 = v239
		goto L42
	} else {
		goto L44
	}
L43:
	;
	v260 = v239
	goto L41
L44:
	;
	goto L43
L45:
	;
	v291 = v40 & int32(3)
	if base.Ui32(v40) < base.Ui32(int32(4)) {
		goto L49
	} else {
		goto L50
	}
L46:
	;
	v434 = v19
	goto L47
L47:
	;
	v437 = base.F64_div(v287, base.F64_sqrt(base.F64_mul(v434, v286)))
	if base.F64_gt(v437, float64(1)) != 0 {
		v445 = v33
		goto L4
	} else {
		goto L61
	}
L48:
	;
	if v291 != 0 {
		goto L55
	} else {
		goto L56
	}
L49:
	;
	v342 = int32(0)
	v358 = float32(0)
	goto L48
L50:
	;
	goto L51
L51:
	;
	v300 = int32(0)
	v302 = v300
	v304 = v300
	v318 = float32(0)
	goto L52
L52:
	;
	v324 = v43 + v302<<(uint(int32(2))%32)
	v325 = *(*float32)(unsafe.Add(mBase, uint32(v324)+12))
	v327 = *(*float32)(unsafe.Add(mBase, uint32(v324)+8))
	v329 = *(*float32)(unsafe.Add(mBase, uint32(v324)+4))
	v331 = *(*float32)(unsafe.Add(mBase, uint32(v324)))
	v336 = base.F32_add(base.F32_mul(v325, v325), base.F32_add(base.F32_mul(v327, v327), base.F32_add(base.F32_mul(v329, v329), base.F32_add(base.F32_mul(v331, v331), v318))))
	v337 = int32(4)
	v338 = v302 + v337
	v340 = v304 + v337
	if v340 != v40&int32(2147483644) {
		v302 = v338
		v304 = v340
		v318 = v336
		goto L52
	} else {
		goto L54
	}
L53:
	;
	v342 = v338
	v358 = v336
	goto L48
L54:
	;
	goto L53
L55:
	;
	v362 = v342
	v365 = int32(0)
	v378 = v358
	goto L58
L56:
	;
	v409 = v358
	goto L57
L57:
	;
	v434 = base.F64_promote_f32(v409)
	goto L47
L58:
	;
	v385 = *(*float32)(unsafe.Add(mBase, uint32(v43+v362<<(uint(int32(2))%32))))
	v387 = base.F32_add(base.F32_mul(v385, v385), v378)
	v388 = int32(1)
	v391 = v365 + v388
	if v391 != v291 {
		v362 = v362 + v388
		v365 = v391
		v378 = v387
		goto L58
	} else {
		goto L60
	}
L59:
	;
	v409 = v387
	goto L57
L60:
	;
	goto L59
L61:
	;
	if base.F64_lt(v437, float64(-1)) == int32(0) {
		v445 = v437
		goto L4
	} else {
		goto L62
	}
L62:
	;
	v445 = float64(-1)
	goto L4
L63:
	;
	m.G0 = v23 + int32(16)
	return v447
}
func F_sparsevec_gt(m *base.Module, l0 int32) int32 {
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
	var v41 int32
	_ = v41
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v66 float32
	_ = v66
	var v69 int32
	_ = v69
	var v76 float32
	_ = v76
	var v79 int32
	_ = v79
	var v81 float32
	_ = v81
	var v83 float32
	_ = v83
	var v89 int32
	_ = v89
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v116 float32
	_ = v116
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v132 float32
	_ = v132
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v157 int32
	_ = v157
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
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
	return int32(0)
L2:
	;
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
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
	return base.B2i32(int32(0) < v157)
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
	v41 = int32(0)
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
	v55 = v41 << (uint(int32(2)) % 32)
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
	v66 = *(*float32)(unsafe.Add(mBase, uint32(v34+v41<<(uint(int32(2))%32))))
	if base.F32_lt(v66, float32(0)) != 0 {
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
	v69 = int32(-1)
	goto L18
L17:
	;
	v69 = int32(1)
	goto L18
L18:
	;
	v157 = v69
	goto L4
L19:
	;
	v76 = *(*float32)(unsafe.Add(mBase, uint32(v28+v41<<(uint(int32(2))%32))))
	if base.F32_lt(v76, float32(0)) != 0 {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	goto L21
L21:
	;
	v81 = *(*float32)(unsafe.Add(mBase, uint32(v55+v34)))
	v83 = *(*float32)(unsafe.Add(mBase, uint32(v55+v28)))
	if base.F32_lt(v81, v83) != 0 {
		goto L25
	} else {
		goto L26
	}
L22:
	;
	v79 = int32(1)
	goto L24
L23:
	;
	v79 = int32(-1)
	goto L24
L24:
	;
	v157 = v79
	goto L4
L25:
	;
	v157 = int32(-1)
	goto L4
L26:
	;
	goto L27
L27:
	;
	if base.F32_gt(v81, v83) != 0 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v157 = int32(1)
	goto L4
L29:
	;
	goto L30
L30:
	;
	v89 = v41 + int32(1)
	if v89 != v36 {
		v41 = v89
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
	v108 = v31 << (uint(int32(2)) % 32)
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v24+v108)))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v3)+4))
	if v111 <= v110 {
		goto L32
	} else {
		goto L34
	}
L34:
	;
	v116 = *(*float32)(unsafe.Add(mBase, uint32(v108+v28)))
	if base.F32_lt(v116, float32(0)) != 0 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v119 = int32(1)
	goto L37
L36:
	;
	v119 = int32(-1)
	goto L37
L37:
	;
	v157 = v119
	goto L4
L38:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v3)+4))
	if v138 < v137 {
		goto L46
	} else {
		goto L47
	}
L39:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	v137 = v122
	goto L38
L40:
	;
	goto L41
L41:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	v125 = v36 << (uint(int32(2)) % 32)
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v30+v125)))
	if v123 <= v127 {
		v137 = v123
		goto L38
	} else {
		goto L42
	}
L42:
	;
	v132 = *(*float32)(unsafe.Add(mBase, uint32(v125+v34)))
	if base.F32_lt(v132, float32(0)) != 0 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v135 = int32(-1)
	goto L45
L44:
	;
	v135 = int32(1)
	goto L45
L45:
	;
	v157 = v135
	goto L4
L46:
	;
	v157 = int32(-1)
	goto L4
L47:
	;
	goto L48
L48:
	;
	v157 = base.B2i32(v137 < v138)
	goto L4
}
func F_sparsevec_in(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v55 int32
	_ = v55
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v107 int32
	_ = v107
	var v117 int32
	_ = v117
	var v124 int64
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v132 int32
	_ = v132
	var v141 int32
	_ = v141
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v186 float32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v206 int32
	_ = v206
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v243 int32
	_ = v243
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v260 int32
	_ = v260
	var v265 int32
	_ = v265
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v290 int32
	_ = v290
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v303 int32
	_ = v303
	var v309 int32
	_ = v309
	var v322 int32
	_ = v322
	var v332 int32
	_ = v332
	var v336 int32
	_ = v336
	var v346 int32
	_ = v346
	var v348 int32
	_ = v348
	var v362 int64
	_ = v362
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v366 int32
	_ = v366
	var v379 int32
	_ = v379
	var v389 int32
	_ = v389
	var v391 int32
	_ = v391
	var v399 int32
	_ = v399
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v418 int32
	_ = v418
	var v422 int32
	_ = v422
	var v426 int32
	_ = v426
	var v438 int32
	_ = v438
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v446 float32
	_ = v446
	var v449 int32
	_ = v449
	var v451 int32
	_ = v451
	var v471 int32
	_ = v471
	var v474 int32
	_ = v474
	var v479 int32
	_ = v479
	var v484 int32
	_ = v484
	var v488 int32
	_ = v488
	var v491 int32
	_ = v491
	var v497 int32
	_ = v497
	var v501 int32
	_ = v501
	var v506 int32
	_ = v506
	var v510 int32
	_ = v510
	var v513 int32
	_ = v513
	var v519 int32
	_ = v519
	var v524 int32
	_ = v524
	var v528 int32
	_ = v528
	var v531 int32
	_ = v531
	var v537 int32
	_ = v537
	var v542 int32
	_ = v542
	var v546 int32
	_ = v546
	var v549 int32
	_ = v549
	var v555 int32
	_ = v555
	var v560 int32
	_ = v560
	var v564 int32
	_ = v564
	var v567 int32
	_ = v567
	var v573 int32
	_ = v573
	var v578 int32
	_ = v578
	var v582 int32
	_ = v582
	var v585 int32
	_ = v585
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v594 int32
	_ = v594
	var v599 int32
	_ = v599
	var v603 int32
	_ = v603
	var v606 int32
	_ = v606
	var v612 int32
	_ = v612
	var v616 int32
	_ = v616
	var v621 int32
	_ = v621
	var v625 int32
	_ = v625
	var v628 int32
	_ = v628
	var v634 int32
	_ = v634
	var v639 int32
	_ = v639
	var v643 int32
	_ = v643
	var v646 int32
	_ = v646
	var v652 int32
	_ = v652
	var v656 int32
	_ = v656
	var v661 int32
	_ = v661
	var v665 int32
	_ = v665
	var v668 int32
	_ = v668
	var v675 int32
	_ = v675
	var v680 int32
	_ = v680
	var v683 int32
	_ = v683
	v2 = int32(0)
	v12 = m.G0
	v14 = v12 - int32(208)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v19 = v17
	v26 = int32(1)
	goto L1
L1:
	;
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
	if v30 != int32(44) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v19 = v19 + int32(1)
	v26 = v683
	goto L1
L4:
	;
	if v30 != 0 {
		v683 = v26
		goto L3
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	v683 = v26 + int32(1)
	goto L3
L7:
	;
	if v26 < int32(16001) {
		goto L18
	} else {
		goto L19
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v665 = m.ExcPending
	if v665 != 0 {
		goto L21
	} else {
		goto L161
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v643 = m.ExcPending
	if v643 != 0 {
		goto L21
	} else {
		goto L156
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v625 = m.ExcPending
	if v625 != 0 {
		goto L21
	} else {
		goto L152
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v603 = m.ExcPending
	if v603 != 0 {
		goto L21
	} else {
		goto L147
	}
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v582 = m.ExcPending
	if v582 != 0 {
		goto L21
	} else {
		goto L142
	}
L13:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v564 = m.ExcPending
	if v564 != 0 {
		goto L21
	} else {
		goto L138
	}
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v546 = m.ExcPending
	if v546 != 0 {
		goto L21
	} else {
		goto L134
	}
L15:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v528 = m.ExcPending
	if v528 != 0 {
		goto L21
	} else {
		goto L130
	}
L16:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v510 = m.ExcPending
	if v510 != 0 {
		goto L21
	} else {
		goto L126
	}
L17:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v488 = m.ExcPending
	if v488 != 0 {
		goto L21
	} else {
		goto L121
	}
L18:
	;
	v36 = F_mul_size(m, int32(8), v26)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
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
	v471 = m.ExcPending
	if v471 != 0 {
		goto L21
	} else {
		goto L117
	}
L21:
	;
	return int32(0)
L22:
	;
	v40 = F_palloc(m, v36)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L21
	} else {
		goto L23
	}
L23:
	;
	v43 = v17
	goto L24
L24:
	;
	v55 = int32(*(*int8)(unsafe.Add(mBase, uint32(v43))))
	goto L26
L25:
	;
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43))))
	if v65 != int32(123) {
		goto L17
	} else {
		goto L28
	}
L26:
	;
	if base.B2i32(v55 == int32(32))|base.B2i32(base.Ui32((v55-int32(9))&int32(255)) < base.Ui32(int32(5))) != 0 {
		v43 = v43 + int32(1)
		goto L24
	} else {
		goto L27
	}
L27:
	;
	goto L25
L28:
	;
	v68 = v43
	goto L29
L29:
	;
	v80 = v68 + int32(1)
	v81 = int32(*(*int8)(unsafe.Add(mBase, uint32(v68)+1)))
	goto L31
L30:
	;
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80))))
	if v91 != int32(125) {
		goto L34
	} else {
		goto L35
	}
L31:
	;
	if base.B2i32(v81 == int32(32))|base.B2i32(base.Ui32((v81-int32(9))&int32(255)) < base.Ui32(int32(5))) != 0 {
		v68 = v80
		goto L29
	} else {
		goto L32
	}
L32:
	;
	goto L30
L33:
	;
	v309 = v298
	goto L86
L34:
	;
	if v26 != 0 {
		goto L37
	} else {
		goto L38
	}
L35:
	;
	goto L36
L36:
	;
	v298 = v68 + int32(2)
	v303 = v2
	goto L33
L37:
	;
	v94 = v80
	v99 = v2
	goto L40
L38:
	;
	goto L39
L39:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L21
	} else {
		goto L82
	}
L40:
	;
	v107 = int32(*(*int8)(unsafe.Add(mBase, uint32(v94))))
	goto L42
L41:
	;
	goto L39
L42:
	;
	if base.B2i32(v107 == int32(32))|base.B2i32(base.Ui32((v107-int32(9))&int32(255)) < base.Ui32(int32(5))) != 0 {
		v94 = v94 + int32(1)
		goto L40
	} else {
		goto L43
	}
L43:
	;
	v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94))))
	if v117 == int32(0) {
		goto L16
	} else {
		goto L44
	}
L44:
	;
	v124 = F_strtox_2(m, v94, v14+int32(204), int32(10), int64(2147483648))
	mBase = m.M
	v125 = base.I32_wrap_i64(v124)
	goto L45
L45:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v14)+204))
	if v126 == v94 {
		goto L15
	} else {
		goto L46
	}
L46:
	;
	v132 = v126
	goto L47
L47:
	;
	v141 = int32(*(*int8)(unsafe.Add(mBase, uint32(v132))))
	goto L49
L48:
	;
	v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v132))))
	if v151 != int32(58) {
		goto L14
	} else {
		goto L51
	}
L49:
	;
	if base.B2i32(v141 == int32(32))|base.B2i32(base.Ui32((v141-int32(9))&int32(255)) < base.Ui32(int32(5))) != 0 {
		v132 = v132 + int32(1)
		goto L47
	} else {
		goto L50
	}
L50:
	;
	goto L48
L51:
	;
	v154 = int32(-2147483647)
	if v125 <= v154 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v157 = v154
	goto L54
L53:
	;
	v157 = v125
	goto L54
L54:
	;
	v158 = v132
	goto L55
L55:
	;
	v169 = int32(*(*int8)(unsafe.Add(mBase, uint32(v158)+1)))
	v171 = v158 + int32(1)
	goto L57
L56:
	;
	*(*int32)(unsafe.Add(mBase, _consts[140])) = int32(0)
	v186 = F_strtof(m, v171, v14+int32(204))
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L21
	} else {
		goto L59
	}
L57:
	;
	if base.B2i32(v169 == int32(32))|base.B2i32(base.Ui32((v169-int32(9))&int32(255)) < base.Ui32(int32(5))) != 0 {
		v158 = v171
		goto L55
	} else {
		goto L58
	}
L58:
	;
	goto L56
L59:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v14)+204))
	if v188 == v171 {
		goto L13
	} else {
		goto L60
	}
L60:
	;
	v191 = *(*int32)(unsafe.Add(mBase, _consts[140]))
	if v191 == int32(68) {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v196 = base.I32_reinterpret_f32(v186) & int32(2147483647)
	v197 = int32(0)
	if base.B2i32(v196 != v197)&base.B2i32(v196 != int32(2139095040)) == v197 {
		goto L12
	} else {
		goto L64
	}
L62:
	;
	goto L63
L63:
	;
	F_CheckElement_2(m, v186)
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L21
	} else {
		goto L65
	}
L64:
	;
	goto L63
L65:
	;
	if base.F32_ne(v186, float32(0)) != 0 {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v211 = v40 + v99<<(uint(int32(3))%32)
	*(*float32)(unsafe.Add(mBase, uint32(v211)+4)) = v186
	v213 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v211))) = v157 - v213
	v219 = v99 + v213
	goto L68
L67:
	;
	v219 = v99
	goto L68
L68:
	;
	v220 = v188
	goto L69
L69:
	;
	v232 = v220 + int32(1)
	v233 = int32(*(*int8)(unsafe.Add(mBase, uint32(v220))))
	goto L71
L70:
	;
	v243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v220))))
	if v243 != int32(44) {
		goto L73
	} else {
		goto L74
	}
L71:
	;
	if base.B2i32(v233 == int32(32))|base.B2i32(base.Ui32((v233-int32(9))&int32(255)) < base.Ui32(int32(5))) != 0 {
		v220 = v232
		goto L69
	} else {
		goto L72
	}
L72:
	;
	goto L70
L73:
	;
	if v243 == int32(125) {
		v298 = v232
		v303 = v219
		goto L33
	} else {
		goto L76
	}
L74:
	;
	goto L75
L75:
	;
	if v219 != v26 {
		v94 = v232
		v99 = v219
		goto L40
	} else {
		goto L81
	}
L76:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L21
	} else {
		goto L77
	}
L77:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L21
	} else {
		goto L78
	}
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+160)) = v17
	F_errmsg(m, int32(760545), v14+int32(160))
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L21
	} else {
		goto L79
	}
L79:
	;
	F_errfinish(m, int32(524189), int32(345), int32(292893))
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L21
	} else {
		goto L80
	}
L80:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L81:
	;
	goto L41
L82:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L21
	} else {
		goto L83
	}
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+80)) = v17
	F_errmsg(m, int32(758642), v14+int32(80))
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L21
	} else {
		goto L84
	}
L84:
	;
	F_errfinish(m, int32(524189), int32(262), int32(292893))
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L21
	} else {
		goto L85
	}
L85:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L86:
	;
	v322 = int32(*(*int8)(unsafe.Add(mBase, uint32(v309))))
	goto L88
L87:
	;
	v332 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v309))))
	if v332 != int32(47) {
		goto L11
	} else {
		goto L90
	}
L88:
	;
	if base.B2i32(v322 == int32(32))|base.B2i32(base.Ui32((v322-int32(9))&int32(255)) < base.Ui32(int32(5))) != 0 {
		v309 = v309 + int32(1)
		goto L86
	} else {
		goto L89
	}
L89:
	;
	goto L87
L90:
	;
	v336 = v309
	goto L91
L91:
	;
	v346 = int32(*(*int8)(unsafe.Add(mBase, uint32(v336)+1)))
	v348 = v336 + int32(1)
	goto L93
L92:
	;
	v362 = F_strtox_2(m, v348, v14+int32(204), int32(10), int64(2147483648))
	mBase = m.M
	v363 = base.I32_wrap_i64(v362)
	goto L95
L93:
	;
	if base.B2i32(v346 == int32(32))|base.B2i32(base.Ui32((v346-int32(9))&int32(255)) < base.Ui32(int32(5))) != 0 {
		v336 = v348
		goto L91
	} else {
		goto L94
	}
L94:
	;
	goto L92
L95:
	;
	v364 = *(*int32)(unsafe.Add(mBase, uint32(v14)+204))
	if v364 == v348 {
		goto L10
	} else {
		goto L96
	}
L96:
	;
	v366 = v364
	goto L97
L97:
	;
	v379 = int32(*(*int8)(unsafe.Add(mBase, uint32(v366))))
	goto L99
L98:
	;
	v389 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v366))))
	if v389 != 0 {
		goto L9
	} else {
		goto L101
	}
L99:
	;
	if base.B2i32(v379 == int32(32))|base.B2i32(base.Ui32((v379-int32(9))&int32(255)) < base.Ui32(int32(5))) != 0 {
		v366 = v366 + int32(1)
		goto L97
	} else {
		goto L100
	}
L100:
	;
	goto L98
L101:
	;
	F_CheckDim_2(m, v363)
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L21
	} else {
		goto L102
	}
L102:
	;
	if base.B2i32(v16 != int32(-1))&base.B2i32(v363 != v16) != 0 {
		goto L8
	} else {
		goto L103
	}
L103:
	;
	F_pg_qsort(m, v40, v303, int32(8), int32(7684))
	mBase = m.M
	v399 = m.ExcPending
	if v399 != 0 {
		goto L21
	} else {
		goto L104
	}
L104:
	;
	v402 = F_mul_size(m, int32(4), v303)
	mBase = m.M
	v403 = m.ExcPending
	if v403 != 0 {
		goto L21
	} else {
		goto L105
	}
L105:
	;
	v404 = F_add_size(m, int32(16), v402)
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
		goto L21
	} else {
		goto L106
	}
L106:
	;
	v407 = F_mul_size(m, int32(4), v303)
	mBase = m.M
	v408 = m.ExcPending
	if v408 != 0 {
		goto L21
	} else {
		goto L107
	}
L107:
	;
	v409 = F_add_size(m, v404, v407)
	mBase = m.M
	v410 = m.ExcPending
	if v410 != 0 {
		goto L21
	} else {
		goto L108
	}
L108:
	;
	v411 = F_palloc0(m, v409)
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L21
	} else {
		goto L109
	}
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v411)+8)) = v303
	*(*int32)(unsafe.Add(mBase, uint32(v411)+4)) = v363
	*(*int32)(unsafe.Add(mBase, uint32(v411))) = v409 << (uint(int32(2)) % 32)
	v418 = int32(0)
	if v418 < v303 {
		goto L110
	} else {
		goto L111
	}
L110:
	;
	v422 = v411 + int32(16)
	v426 = v418
	goto L113
L111:
	;
	goto L112
L112:
	;
	m.G0 = v14 + int32(208)
	return v411
L113:
	;
	v438 = v426 << (uint(int32(2)) % 32)
	v442 = v40 + v426<<(uint(int32(3))%32)
	v443 = *(*int32)(unsafe.Add(mBase, uint32(v442)))
	*(*int32)(unsafe.Add(mBase, uint32(v422+v438))) = v443
	v446 = *(*float32)(unsafe.Add(mBase, uint32(v442)+4))
	*(*float32)(unsafe.Add(mBase, uint32(v438+(v422+v303<<(uint(int32(2))%32))))) = v446
	F_CheckIndex(m, v422, v426, v363)
	mBase = m.M
	v449 = m.ExcPending
	if v449 != 0 {
		goto L21
	} else {
		goto L115
	}
L114:
	;
	goto L112
L115:
	;
	v451 = v426 + int32(1)
	if v451 != v303 {
		v426 = v451
		goto L113
	} else {
		goto L116
	}
L116:
	;
	goto L114
L117:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v474 = m.ExcPending
	if v474 != 0 {
		goto L21
	} else {
		goto L118
	}
L118:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = int32(16000)
	F_errmsg(m, int32(130868), v14)
	mBase = m.M
	v479 = m.ExcPending
	if v479 != 0 {
		goto L21
	} else {
		goto L119
	}
L119:
	;
	F_errfinish(m, int32(524189), int32(230), int32(292893))
	mBase = m.M
	v484 = m.ExcPending
	if v484 != 0 {
		goto L21
	} else {
		goto L120
	}
L120:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L121:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v491 = m.ExcPending
	if v491 != 0 {
		goto L21
	} else {
		goto L122
	}
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+192)) = v17
	F_errmsg(m, int32(760545), v14+int32(192))
	mBase = m.M
	v497 = m.ExcPending
	if v497 != 0 {
		goto L21
	} else {
		goto L123
	}
L123:
	;
	F_errdetail(m, int32(691991), int32(0))
	mBase = m.M
	v501 = m.ExcPending
	if v501 != 0 {
		goto L21
	} else {
		goto L124
	}
L124:
	;
	F_errfinish(m, int32(524189), int32(243), int32(292893))
	mBase = m.M
	v506 = m.ExcPending
	if v506 != 0 {
		goto L21
	} else {
		goto L125
	}
L125:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L126:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v513 = m.ExcPending
	if v513 != 0 {
		goto L21
	} else {
		goto L127
	}
L127:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+96)) = v17
	F_errmsg(m, int32(760545), v14+int32(96))
	mBase = m.M
	v519 = m.ExcPending
	if v519 != 0 {
		goto L21
	} else {
		goto L128
	}
L128:
	;
	F_errfinish(m, int32(524189), int32(271), int32(292893))
	mBase = m.M
	v524 = m.ExcPending
	if v524 != 0 {
		goto L21
	} else {
		goto L129
	}
L129:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L130:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v531 = m.ExcPending
	if v531 != 0 {
		goto L21
	} else {
		goto L131
	}
L131:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+112)) = v17
	F_errmsg(m, int32(760545), v14+int32(112))
	mBase = m.M
	v537 = m.ExcPending
	if v537 != 0 {
		goto L21
	} else {
		goto L132
	}
L132:
	;
	F_errfinish(m, int32(524189), int32(279), int32(292893))
	mBase = m.M
	v542 = m.ExcPending
	if v542 != 0 {
		goto L21
	} else {
		goto L133
	}
L133:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L134:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v549 = m.ExcPending
	if v549 != 0 {
		goto L21
	} else {
		goto L135
	}
L135:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+176)) = v17
	F_errmsg(m, int32(760545), v14+int32(176))
	mBase = m.M
	v555 = m.ExcPending
	if v555 != 0 {
		goto L21
	} else {
		goto L136
	}
L136:
	;
	F_errfinish(m, int32(524189), int32(295), int32(292893))
	mBase = m.M
	v560 = m.ExcPending
	if v560 != 0 {
		goto L21
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
	F_errcode(m, int32(33685634))
	mBase = m.M
	v567 = m.ExcPending
	if v567 != 0 {
		goto L21
	} else {
		goto L139
	}
L139:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+128)) = v17
	F_errmsg(m, int32(760545), v14+int32(128))
	mBase = m.M
	v573 = m.ExcPending
	if v573 != 0 {
		goto L21
	} else {
		goto L140
	}
L140:
	;
	F_errfinish(m, int32(524189), int32(311), int32(292893))
	mBase = m.M
	v578 = m.ExcPending
	if v578 != 0 {
		goto L21
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
	F_errcode(m, int32(50331778))
	mBase = m.M
	v585 = m.ExcPending
	if v585 != 0 {
		goto L21
	} else {
		goto L143
	}
L143:
	;
	v587 = F_pnstrdup(m, v171, v188-v171)
	mBase = m.M
	v588 = m.ExcPending
	if v588 != 0 {
		goto L21
	} else {
		goto L144
	}
L144:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+144)) = v587
	F_errmsg(m, int32(513709), v14+int32(144))
	mBase = m.M
	v594 = m.ExcPending
	if v594 != 0 {
		goto L21
	} else {
		goto L145
	}
L145:
	;
	F_errfinish(m, int32(524189), int32(317), int32(292893))
	mBase = m.M
	v599 = m.ExcPending
	if v599 != 0 {
		goto L21
	} else {
		goto L146
	}
L146:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L147:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v606 = m.ExcPending
	if v606 != 0 {
		goto L21
	} else {
		goto L148
	}
L148:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+64)) = v17
	F_errmsg(m, int32(760545), v14-int32(-64))
	mBase = m.M
	v612 = m.ExcPending
	if v612 != 0 {
		goto L21
	} else {
		goto L149
	}
L149:
	;
	F_errdetail(m, int32(604977), int32(0))
	mBase = m.M
	v616 = m.ExcPending
	if v616 != 0 {
		goto L21
	} else {
		goto L150
	}
L150:
	;
	F_errfinish(m, int32(524189), int32(356), int32(292893))
	mBase = m.M
	v621 = m.ExcPending
	if v621 != 0 {
		goto L21
	} else {
		goto L151
	}
L151:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L152:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v628 = m.ExcPending
	if v628 != 0 {
		goto L21
	} else {
		goto L153
	}
L153:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v17
	F_errmsg(m, int32(760545), v14+int32(16))
	mBase = m.M
	v634 = m.ExcPending
	if v634 != 0 {
		goto L21
	} else {
		goto L154
	}
L154:
	;
	F_errfinish(m, int32(524189), int32(369), int32(292893))
	mBase = m.M
	v639 = m.ExcPending
	if v639 != 0 {
		goto L21
	} else {
		goto L155
	}
L155:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L156:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v646 = m.ExcPending
	if v646 != 0 {
		goto L21
	} else {
		goto L157
	}
L157:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = v17
	F_errmsg(m, int32(760545), v14+int32(48))
	mBase = m.M
	v652 = m.ExcPending
	if v652 != 0 {
		goto L21
	} else {
		goto L158
	}
L158:
	;
	F_errdetail(m, int32(654004), int32(0))
	mBase = m.M
	v656 = m.ExcPending
	if v656 != 0 {
		goto L21
	} else {
		goto L159
	}
L159:
	;
	F_errfinish(m, int32(524189), int32(387), int32(292893))
	mBase = m.M
	v661 = m.ExcPending
	if v661 != 0 {
		goto L21
	} else {
		goto L160
	}
L160:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L161:
	;
	F_errcode(m, int32(130))
	mBase = m.M
	v668 = m.ExcPending
	if v668 != 0 {
		goto L21
	} else {
		goto L162
	}
L162:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+36)) = v363
	*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v16
	F_errmsg(m, int32(488679), v14+int32(32))
	mBase = m.M
	v675 = m.ExcPending
	if v675 != 0 {
		goto L21
	} else {
		goto L163
	}
L163:
	;
	F_errfinish(m, int32(524189), int32(62), int32(302055))
	mBase = m.M
	v680 = m.ExcPending
	if v680 != 0 {
		goto L21
	} else {
		goto L164
	}
L164:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_sparsevec_l1_distance(m *base.Module, l0 int32) int32 {
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
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v61 float32
	_ = v61
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
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
	var v110 int32
	_ = v110
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
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v167 float32
	_ = v167
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v190 float32
	_ = v190
	var v194 float32
	_ = v194
	var v196 float32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v217 float32
	_ = v217
	var v227 int32
	_ = v227
	var v242 float32
	_ = v242
	var v244 int32
	_ = v244
	var v246 float32
	_ = v246
	var v249 float32
	_ = v249
	var v252 float32
	_ = v252
	var v255 float32
	_ = v255
	var v260 float32
	_ = v260
	var v262 int32
	_ = v262
	var v279 float32
	_ = v279
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	v2 = int32(0)
	v16 = float32(0)
	v17 = m.G0
	v19 = v17 - int32(16)
	m.G0 = v19
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
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
	return int32(0)
L2:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
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
		goto L5
	} else {
		goto L6
	}
L4:
	;
	if v34 <= v154 {
		v279 = v167
		goto L35
	} else {
		goto L36
	}
L5:
	;
	v33 = v27 + int32(16)
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v27)+8))
	v37 = v33 + v34<<(uint(int32(2))%32)
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
	if v38 <= int32(0) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L1
	} else {
		goto L31
	}
L8:
	;
	v154 = v2
	v167 = v16
	goto L4
L9:
	;
	goto L10
L10:
	;
	v42 = v22 + int32(16)
	v45 = v42 + v38<<(uint(int32(2))%32)
	v48 = v2
	v54 = v2
	v61 = v16
	goto L11
L11:
	;
	v63 = v54 << (uint(int32(2)) % 32)
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v42+v63)))
	if v34 <= v48 {
		v110 = v48
		v111 = int32(-1)
		v123 = v61
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v154 = v110
	v167 = v129
	goto L4
L13:
	;
	if v65 != v111 {
		goto L27
	} else {
		goto L28
	}
L14:
	;
	v69 = v48
	v71 = v48
	v84 = v61
	goto L15
L15:
	;
	v86 = v69 << (uint(int32(2)) % 32)
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v33+v86)))
	if v88 == v65 {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	v110 = v105
	v111 = v88
	v123 = v101
	goto L13
L17:
	;
	v103 = v69 + int32(1)
	if v65 < v88 {
		goto L22
	} else {
		goto L23
	}
L18:
	;
	v90 = *(*float32)(unsafe.Add(mBase, uint32(v45+v63)))
	v92 = *(*float32)(unsafe.Add(mBase, uint32(v37+v86)))
	v101 = base.F32_add(base.F32_abs(base.F32_sub(v90, v92)), v84)
	goto L17
L19:
	;
	goto L20
L20:
	;
	if v65 <= v88 {
		v101 = v84
		goto L17
	} else {
		goto L21
	}
L21:
	;
	v98 = *(*float32)(unsafe.Add(mBase, uint32(v37+v86)))
	v101 = base.F32_add(base.F32_abs(v98), v84)
	goto L17
L22:
	;
	v105 = v71
	goto L24
L23:
	;
	v105 = v103
	goto L24
L24:
	;
	if v65 <= v88 {
		v110 = v105
		v111 = v88
		v123 = v101
		goto L13
	} else {
		goto L25
	}
L25:
	;
	if v103 < v34 {
		v69 = v103
		v71 = v105
		v84 = v101
		goto L15
	} else {
		goto L26
	}
L26:
	;
	goto L16
L27:
	;
	v126 = *(*float32)(unsafe.Add(mBase, uint32(v45+v63)))
	v129 = base.F32_add(base.F32_abs(v126), v123)
	goto L29
L28:
	;
	v129 = v123
	goto L29
L29:
	;
	v131 = v54 + int32(1)
	if v38 != v131 {
		v48 = v110
		v54 = v131
		v61 = v129
		goto L11
	} else {
		goto L30
	}
L30:
	;
	goto L12
L31:
	;
	F_errcode(m, int32(130))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = v141
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v140
	F_errmsg(m, int32(499677), v19)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	F_errfinish(m, int32(524189), int32(50), int32(159834))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L35:
	;
	v281 = F_Float8GetDatum(m, base.F64_promote_f32(v279))
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L1
	} else {
		goto L48
	}
L36:
	;
	v171 = (v34 - v154) & int32(3)
	if v171 == int32(0) {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	if base.Ui32(int32(-4)) < base.Ui32(v154-v34) {
		v279 = v217
		goto L35
	} else {
		goto L44
	}
L38:
	;
	v202 = v154
	v217 = v167
	goto L37
L39:
	;
	goto L40
L40:
	;
	v176 = v154
	v178 = int32(0)
	v190 = v167
	goto L41
L41:
	;
	v194 = *(*float32)(unsafe.Add(mBase, uint32(v37+v176<<(uint(int32(2))%32))))
	v196 = base.F32_add(base.F32_abs(v194), v190)
	v197 = int32(1)
	v198 = v176 + v197
	v200 = v178 + v197
	if v200 != v171 {
		v176 = v198
		v178 = v200
		v190 = v196
		goto L41
	} else {
		goto L43
	}
L42:
	;
	v202 = v198
	v217 = v196
	goto L37
L43:
	;
	goto L42
L44:
	;
	v227 = v202
	v242 = v217
	goto L45
L45:
	;
	v244 = v227 << (uint(int32(2)) % 32)
	v246 = *(*float32)(unsafe.Add(mBase, uint32(v37+int32(12)+v244)))
	v249 = *(*float32)(unsafe.Add(mBase, uint32(v244+(v37+int32(8)))))
	v252 = *(*float32)(unsafe.Add(mBase, uint32(v244+(v37+int32(4)))))
	v255 = *(*float32)(unsafe.Add(mBase, uint32(v244+v37)))
	v260 = base.F32_add(base.F32_abs(v246), base.F32_add(base.F32_abs(v249), base.F32_add(base.F32_abs(v252), base.F32_add(base.F32_abs(v255), v242))))
	v262 = v227 + int32(4)
	if v262 != v34 {
		v227 = v262
		v242 = v260
		goto L45
	} else {
		goto L47
	}
L46:
	;
	v279 = v260
	goto L35
L47:
	;
	goto L46
L48:
	;
	m.G0 = v19 + int32(16)
	return v281
}
func F_sparsevec_to_vector(m *base.Module, l0 int32) int32 {
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
	var v63 int32
	_ = v63
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
	var v120 int32
	_ = v120
	var v134 int32
	_ = v134
	var v142 int32
	_ = v142
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
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v187 int32
	_ = v187
	var v192 int32
	_ = v192
	v2 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(16)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v18 = F_pg_detoast_datum(m, v17)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		return int32(0)
	} else {
		v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v23 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
		v24 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
		F_CheckDim_3(m, v24)
		mBase = m.M
		v26 = m.ExcPending
		if v26 != 0 {
			return int32(0)
		} else {
			if base.B2i32(v22 != int32(-1))&base.B2i32(v24 != v22) == int32(0) {
				v35 = F_mul_size(m, int32(4), v24)
				mBase = m.M
				v36 = m.ExcPending
				if v36 != 0 {
					return int32(0)
				} else {
					v37 = F_add_size(m, int32(8), v35)
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
						return int32(0)
					} else {
						v39 = F_palloc0(m, v37)
						mBase = m.M
						v40 = m.ExcPending
						if v40 != 0 {
							return int32(0)
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
									v63 = v57
									v72 = v2
									for {
										v74 = int32(2)
										v75 = v63 << (uint(v74) % 32)
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
										v115 = v63 + v84
										v117 = v72 + v84
										if v117 != v45&int32(2147483644) {
											v63 = v115
											v72 = v117
											continue
										} else {
											break
										}
										break
									}
									v120 = v115
								} else {
									v120 = v57
								}
								if v54 == int32(0) {
								} else {
									v134 = v120
									v142 = v2
									for {
										v145 = int32(2)
										v146 = v134 << (uint(v145) % 32)
										v148 = *(*int32)(unsafe.Add(mBase, uint32(v49+v146)))
										v153 = *(*float32)(unsafe.Add(mBase, uint32(v146+v52)))
										*(*float32)(unsafe.Add(mBase, uint32(v56+v148<<(uint(v145)%32)))) = v153
										v155 = int32(1)
										v158 = v142 + v155
										if v158 != v54 {
											v134 = v134 + v155
											v142 = v158
											continue
										} else {
											break
										}
										break
									}
								}
							}
							m.G0 = v15 + int32(16)
							return v39
						}
					}
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v179 = m.ExcPending
				if v179 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(130))
					mBase = m.M
					v182 = m.ExcPending
					if v182 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v24
						*(*int32)(unsafe.Add(mBase, uint32(v15))) = v22
						F_errmsg(m, int32(488679), v15)
						mBase = m.M
						v187 = m.ExcPending
						if v187 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(518027), int32(88), int32(302055))
							mBase = m.M
							v192 = m.ExcPending
							if v192 != 0 {
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
