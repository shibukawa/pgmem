package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_sparsevec_eq(m *base.Module, l0 int32) int32 {
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
	return base.B2i32(v157 == int32(0))
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
func F_sparsevec_inner_product(m *base.Module, l0 int32) int32 {
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
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v62 float32
	_ = v62
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v85 float32
	_ = v85
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v92 float32
	_ = v92
	var v94 float32
	_ = v94
	var v97 float32
	_ = v97
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v107 float32
	_ = v107
	var v109 int32
	_ = v109
	var v128 float64
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v148 int32
	_ = v148
	var v153 int32
	_ = v153
	v2 = int32(0)
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
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
	if v32 <= int32(0) {
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
	v138 = m.ExcPending
	if v138 != 0 {
		goto L1
	} else {
		goto L29
	}
L7:
	;
	v128 = float64(0)
	goto L9
L8:
	;
	v36 = int32(16)
	v37 = v27 + v36
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v27)+8))
	v39 = int32(2)
	v43 = v22 + v36
	v49 = v2
	v51 = v2
	v62 = float32(0)
	goto L10
L9:
	;
	v129 = F_Float8GetDatum(m, v128)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L1
	} else {
		goto L28
	}
L10:
	;
	if v38 < v49 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v128 = base.F64_promote_f32(v107)
	goto L9
L12:
	;
	v64 = v49
	goto L14
L13:
	;
	v64 = v38
	goto L14
L14:
	;
	v66 = v51 << (uint(int32(2)) % 32)
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v66+v43)))
	v71 = v49
	v72 = v49
	v85 = v62
	goto L15
L15:
	;
	if v71 != v64 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v109 = v51 + int32(1)
	if v109 != v32 {
		v49 = v104
		v51 = v109
		v62 = v107
		goto L10
	} else {
		goto L27
	}
L17:
	;
	v88 = v71 << (uint(int32(2)) % 32)
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v37+v88)))
	if v90 == v69 {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	v104 = v72
	v107 = v85
	goto L19
L19:
	;
	goto L16
L20:
	;
	v92 = *(*float32)(unsafe.Add(mBase, uint32(v43+v32<<(uint(v39)%32)+v66)))
	v94 = *(*float32)(unsafe.Add(mBase, uint32(v37+v38<<(uint(v39)%32)+v88)))
	v97 = base.F32_add(base.F32_mul(v92, v94), v85)
	goto L22
L21:
	;
	v97 = v85
	goto L22
L22:
	;
	v99 = v71 + int32(1)
	if v69 < v90 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v101 = v72
	goto L25
L24:
	;
	v101 = v99
	goto L25
L25:
	;
	if v90 < v69 {
		v71 = v99
		v72 = v101
		v85 = v97
		goto L15
	} else {
		goto L26
	}
L26:
	;
	v104 = v101
	v107 = v97
	goto L19
L27:
	;
	goto L11
L28:
	;
	m.G0 = v19 + int32(16)
	return v129
L29:
	;
	F_errcode(m, int32(130))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = v143
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v142
	F_errmsg(m, int32(496667), v19)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	F_errfinish(m, int32(521057), int32(50), int32(158859))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_sparsevec_l2_normalize(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v14 float64
	_ = v14
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
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v59 int32
	_ = v59
	var v66 int32
	_ = v66
	var v72 float64
	_ = v72
	var v77 int32
	_ = v77
	var v78 float32
	_ = v78
	var v79 float64
	_ = v79
	var v81 float32
	_ = v81
	var v82 float64
	_ = v82
	var v84 float32
	_ = v84
	var v85 float64
	_ = v85
	var v87 float32
	_ = v87
	var v88 float64
	_ = v88
	var v93 float64
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v112 float64
	_ = v112
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v128 float64
	_ = v128
	var v134 float32
	_ = v134
	var v135 float64
	_ = v135
	var v137 float64
	_ = v137
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v156 float64
	_ = v156
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v194 float32
	_ = v194
	var v197 float32
	_ = v197
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v248 int32
	_ = v248
	var v260 int32
	_ = v260
	var v262 float32
	_ = v262
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v300 int32
	_ = v300
	var v305 int32
	_ = v305
	var v309 int32
	_ = v309
	var v314 int32
	_ = v314
	v2 = int32(0)
	v14 = float64(0)
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v18 = F_pg_detoast_datum(m, v17)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
	v26 = F_mul_size(m, int32(4), v25)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v28 = F_add_size(m, int32(16), v26)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v31 = F_mul_size(m, int32(4), v25)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v33 = F_add_size(m, v28, v31)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v35 = F_palloc0(m, v33)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v35)+8)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v35)+4)) = v22
	*(*int32)(unsafe.Add(mBase, uint32(v35))) = v33 << (uint(int32(2)) % 32)
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
	if v42 <= int32(0) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	return v35
L9:
	;
	goto L10
L10:
	;
	v47 = v18 + int32(16)
	v49 = v25 << (uint(int32(2)) % 32)
	v50 = v47 + v49
	v52 = v42 & int32(3)
	if base.Ui32(v42) < base.Ui32(int32(4)) {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	if v52 != 0 {
		goto L18
	} else {
		goto L19
	}
L12:
	;
	v99 = int32(0)
	v112 = v14
	goto L11
L13:
	;
	goto L14
L14:
	;
	v59 = int32(0)
	v66 = v2
	v72 = v14
	goto L15
L15:
	;
	v77 = v50 + v59<<(uint(int32(2))%32)
	v78 = *(*float32)(unsafe.Add(mBase, uint32(v77)+12))
	v79 = base.F64_promote_f32(v78)
	v81 = *(*float32)(unsafe.Add(mBase, uint32(v77)+8))
	v82 = base.F64_promote_f32(v81)
	v84 = *(*float32)(unsafe.Add(mBase, uint32(v77)+4))
	v85 = base.F64_promote_f32(v84)
	v87 = *(*float32)(unsafe.Add(mBase, uint32(v77)))
	v88 = base.F64_promote_f32(v87)
	v93 = base.F64_add(base.F64_mul(v79, v79), base.F64_add(base.F64_mul(v82, v82), base.F64_add(base.F64_mul(v85, v85), base.F64_add(base.F64_mul(v88, v88), v72))))
	v94 = int32(4)
	v95 = v59 + v94
	v97 = v66 + v94
	if v97 != v42&int32(2147483644) {
		v59 = v95
		v66 = v97
		v72 = v93
		goto L15
	} else {
		goto L17
	}
L16:
	;
	v99 = v95
	v112 = v93
	goto L11
L17:
	;
	goto L16
L18:
	;
	v115 = v99
	v119 = v2
	v128 = v112
	goto L21
L19:
	;
	v156 = v112
	goto L20
L20:
	;
	if base.F64_gt(v156, float64(0)) == int32(0) {
		goto L24
	} else {
		goto L25
	}
L21:
	;
	v134 = *(*float32)(unsafe.Add(mBase, uint32(v50+v115<<(uint(int32(2))%32))))
	v135 = base.F64_promote_f32(v134)
	v137 = base.F64_add(base.F64_mul(v135, v135), v128)
	v138 = int32(1)
	v141 = v119 + v138
	if v141 != v52 {
		v115 = v115 + v138
		v119 = v141
		v128 = v137
		goto L21
	} else {
		goto L23
	}
L22:
	;
	v156 = v137
	goto L20
L23:
	;
	goto L22
L24:
	;
	return v35
L25:
	;
	goto L26
L26:
	;
	v165 = v35 + int32(16)
	v166 = v165 + v49
	v168 = int32(0)
	v171 = v168
	v174 = v168
	goto L28
L27:
	;
	if v204 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L28:
	;
	v187 = v171 << (uint(int32(2)) % 32)
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v187+v47)))
	*(*int32)(unsafe.Add(mBase, uint32(v165+v187))) = v190
	v194 = *(*float32)(unsafe.Add(mBase, uint32(v187+v50)))
	v197 = base.F32_demote_f64(base.F64_div(base.F64_promote_f32(v194), base.F64_sqrt(v156)))
	*(*float32)(unsafe.Add(mBase, uint32(v187+v166))) = v197
	if base.F32_eq(base.F32_abs(v197), math.Float32frombits(uint32(0x7f800000))) != 0 {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	F_float_overflow_error(m)
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L1
	} else {
		goto L33
	}
L30:
	;
	goto L29
L31:
	;
	v204 = v174 + base.F32_eq(v197, float32(0))
	v206 = v171 + int32(1)
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
	if v206 < v207 {
		v171 = v206
		v174 = v204
		goto L28
	} else {
		goto L32
	}
L32:
	;
	goto L27
L33:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L34:
	;
	return v35
L35:
	;
	goto L36
L36:
	;
	v216 = v25 - v204
	v217 = F_mul_size(m, int32(4), v216)
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	v219 = F_add_size(m, int32(16), v217)
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	v222 = F_mul_size(m, int32(4), v216)
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	v224 = F_add_size(m, v219, v222)
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	v226 = F_palloc0(m, v224)
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v226)+8)) = v216
	*(*int32)(unsafe.Add(mBase, uint32(v226)+4)) = v22
	*(*int32)(unsafe.Add(mBase, uint32(v226))) = v224 << (uint(int32(2)) % 32)
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v35)+8))
	if int32(0) < v233 {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L1
	} else {
		goto L54
	}
L43:
	;
	v237 = v226 + int32(16)
	v241 = int32(0)
	v243 = v241
	v244 = v241
	v248 = v233
	goto L46
L44:
	;
	goto L45
L45:
	;
	F_pfree(m, v35)
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L1
	} else {
		goto L53
	}
L46:
	;
	v260 = v243 << (uint(int32(2)) % 32)
	v262 = *(*float32)(unsafe.Add(mBase, uint32(v166+v260)))
	if base.F32_ne(v262, float32(0)) != 0 {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	goto L45
L48:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v226)+8))
	if v265 <= v244 {
		goto L42
	} else {
		goto L51
	}
L49:
	;
	v278 = v244
	v279 = v248
	goto L50
L50:
	;
	v281 = v243 + int32(1)
	if v281 < v279 {
		v243 = v281
		v244 = v278
		v248 = v279
		goto L46
	} else {
		goto L52
	}
L51:
	;
	v268 = v244 << (uint(int32(2)) % 32)
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v260+v165)))
	*(*int32)(unsafe.Add(mBase, uint32(v237+v268))) = v271
	*(*float32)(unsafe.Add(mBase, uint32(v268+(v237+v216<<(uint(int32(2))%32))))) = v262
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v35)+8))
	v278 = v244 + int32(1)
	v279 = v275
	goto L50
L52:
	;
	goto L47
L53:
	;
	return v226
L54:
	;
	F_errmsg_internal(m, int32(471885), int32(0))
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	F_errfinish(m, int32(521057), int32(1129), int32(356130))
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_sparsevec_out(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 float32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v120 int32
	_ = v120
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v138 int32
	_ = v138
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v160 int32
	_ = v160
	var v165 int64
	_ = v165
	var v167 int64
	_ = v167
	var v168 int64
	_ = v168
	var v170 int64
	_ = v170
	var v172 int32
	_ = v172
	var v174 int64
	_ = v174
	var v175 int64
	_ = v175
	var v177 int32
	_ = v177
	var v184 int32
	_ = v184
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v196 int64
	_ = v196
	var v200 int32
	_ = v200
	var v201 int64
	_ = v201
	var v203 int32
	_ = v203
	var v211 int32
	_ = v211
	var v212 int64
	_ = v212
	var v216 int32
	_ = v216
	var v217 int64
	_ = v217
	var v219 int32
	_ = v219
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v245 int64
	_ = v245
	var v249 int64
	_ = v249
	var v251 int32
	_ = v251
	var v254 int64
	_ = v254
	var v256 int32
	_ = v256
	var v267 int32
	_ = v267
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v282 int32
	_ = v282
	var v287 int32
	_ = v287
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v325 int32
	_ = v325
	var v329 int32
	_ = v329
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v352 int32
	_ = v352
	var v373 int32
	_ = v373
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v380 int64
	_ = v380
	var v382 int64
	_ = v382
	var v383 int64
	_ = v383
	var v385 int64
	_ = v385
	var v387 int32
	_ = v387
	var v389 int64
	_ = v389
	var v390 int64
	_ = v390
	var v392 int32
	_ = v392
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v407 int32
	_ = v407
	var v409 int32
	_ = v409
	var v410 int64
	_ = v410
	var v414 int32
	_ = v414
	var v415 int64
	_ = v415
	var v417 int32
	_ = v417
	var v425 int32
	_ = v425
	var v426 int64
	_ = v426
	var v430 int32
	_ = v430
	var v431 int64
	_ = v431
	var v433 int32
	_ = v433
	var v441 int32
	_ = v441
	var v443 int32
	_ = v443
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v449 int32
	_ = v449
	var v452 int32
	_ = v452
	var v457 int64
	_ = v457
	var v461 int64
	_ = v461
	var v463 int32
	_ = v463
	var v466 int64
	_ = v466
	var v468 int32
	_ = v468
	var v481 int32
	_ = v481
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v496 int32
	_ = v496
	var v502 int32
	_ = v502
	var v512 int32
	_ = v512
	var v514 int32
	_ = v514
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v520 int32
	_ = v520
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v532 int32
	_ = v532
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v541 int32
	_ = v541
	var v543 int32
	_ = v543
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v561 int32
	_ = v561
	var v566 int32
	_ = v566
	var v568 int32
	_ = v568
	var v570 int32
	_ = v570
	var v576 int32
	_ = v576
	var v578 int32
	_ = v578
	var v580 int32
	_ = v580
	var v585 int32
	_ = v585
	var v593 int32
	_ = v593
	var v595 int32
	_ = v595
	var v608 int32
	_ = v608
	var v610 int32
	_ = v610
	var v613 int32
	_ = v613
	var v614 int32
	_ = v614
	var v616 int32
	_ = v616
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v628 int32
	_ = v628
	var v632 int32
	_ = v632
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v636 int32
	_ = v636
	var v650 int32
	_ = v650
	var v651 int32
	_ = v651
	var v652 int32
	_ = v652
	var v654 int32
	_ = v654
	var v656 int32
	_ = v656
	var v663 int32
	_ = v663
	var v664 int32
	_ = v664
	var v666 int32
	_ = v666
	var v672 int32
	_ = v672
	var v688 int32
	_ = v688
	var v690 int32
	_ = v690
	var v694 int32
	_ = v694
	var v705 int32
	_ = v705
	var v706 int32
	_ = v706
	var v707 int32
	_ = v707
	var v713 int32
	_ = v713
	var v723 int32
	_ = v723
	var v752 int32
	_ = v752
	var v755 int32
	_ = v755
	var v765 int32
	_ = v765
	var v772 int32
	_ = v772
	var v773 int32
	_ = v773
	var v774 int32
	_ = v774
	var v777 int32
	_ = v777
	var v780 int32
	_ = v780
	var v785 int32
	_ = v785
	var v786 int32
	_ = v786
	var v798 int32
	_ = v798
	var v803 int32
	_ = v803
	var v822 int32
	_ = v822
	var v829 int32
	_ = v829
	var v835 int32
	_ = v835
	var v845 int32
	_ = v845
	var v852 int32
	_ = v852
	var v853 int32
	_ = v853
	var v869 int32
	_ = v869
	var v873 int32
	_ = v873
	var v876 int32
	_ = v876
	var v877 int32
	_ = v877
	var v878 int32
	_ = v878
	var v879 int32
	_ = v879
	var v883 int32
	_ = v883
	var v894 int32
	_ = v894
	var v897 int32
	_ = v897
	var v902 int32
	_ = v902
	var v904 int32
	_ = v904
	var v926 int32
	_ = v926
	var v928 int32
	_ = v928
	var v929 int32
	_ = v929
	var v939 int32
	_ = v939
	var v943 int32
	_ = v943
	var v944 int32
	_ = v944
	var v951 int32
	_ = v951
	var v954 int32
	_ = v954
	var v958 int32
	_ = v958
	var v962 int32
	_ = v962
	var v963 int32
	_ = v963
	var v964 int32
	_ = v964
	var v965 int32
	_ = v965
	var v970 int32
	_ = v970
	var v974 int32
	_ = v974
	var v975 int32
	_ = v975
	var v976 int32
	_ = v976
	var v977 int32
	_ = v977
	var v982 int32
	_ = v982
	var v983 int32
	_ = v983
	var v987 int32
	_ = v987
	var v992 int32
	_ = v992
	var v996 int32
	_ = v996
	var v997 int64
	_ = v997
	var v999 int32
	_ = v999
	var v1001 int32
	_ = v1001
	var v1008 int32
	_ = v1008
	var v1009 int32
	_ = v1009
	var v1025 int32
	_ = v1025
	var v1026 int32
	_ = v1026
	var v1029 int32
	_ = v1029
	var v1032 int32
	_ = v1032
	var v1033 int32
	_ = v1033
	var v1034 int32
	_ = v1034
	var v1035 int32
	_ = v1035
	var v1039 int32
	_ = v1039
	var v1050 int32
	_ = v1050
	var v1053 int32
	_ = v1053
	var v1059 int32
	_ = v1059
	var v1060 int32
	_ = v1060
	var v1080 int32
	_ = v1080
	var v1082 int32
	_ = v1082
	var v1084 int32
	_ = v1084
	var v1085 int32
	_ = v1085
	var v1095 int32
	_ = v1095
	var v1099 int32
	_ = v1099
	var v1100 int32
	_ = v1100
	var v1112 int32
	_ = v1112
	var v1116 int32
	_ = v1116
	var v1118 int32
	_ = v1118
	var v1123 int32
	_ = v1123
	var v1126 int32
	_ = v1126
	var v1129 int32
	_ = v1129
	var v1132 int32
	_ = v1132
	var v1137 int32
	_ = v1137
	var v1140 int32
	_ = v1140
	var v1143 int32
	_ = v1143
	var v1147 int32
	_ = v1147
	var v1155 int32
	_ = v1155
	var v1158 int32
	_ = v1158
	var v1181 int32
	_ = v1181
	var v1182 int32
	_ = v1182
	var v1183 int32
	_ = v1183
	var v1187 int32
	_ = v1187
	var v1190 int32
	_ = v1190
	var v1194 int32
	_ = v1194
	var v1197 int32
	_ = v1197
	var v1199 int32
	_ = v1199
	var v1200 int32
	_ = v1200
	var v1201 int32
	_ = v1201
	var v1203 int32
	_ = v1203
	var v1208 int32
	_ = v1208
	var v1213 int32
	_ = v1213
	var v1214 int32
	_ = v1214
	var v1216 int32
	_ = v1216
	var v1217 int32
	_ = v1217
	var v1219 int32
	_ = v1219
	var v1221 int32
	_ = v1221
	var v1222 int32
	_ = v1222
	var v1225 float32
	_ = v1225
	var v1227 int32
	_ = v1227
	var v1228 int32
	_ = v1228
	var v1245 int32
	_ = v1245
	var v1247 int32
	_ = v1247
	var v1250 int32
	_ = v1250
	var v1251 int32
	_ = v1251
	var v1256 int32
	_ = v1256
	var v1260 int32
	_ = v1260
	var v1263 int32
	_ = v1263
	var v1268 int32
	_ = v1268
	var v1272 int32
	_ = v1272
	var v1279 int32
	_ = v1279
	var v1280 int32
	_ = v1280
	var v1286 int32
	_ = v1286
	var v1291 int32
	_ = v1291
	var v1293 int32
	_ = v1293
	var v1304 int32
	_ = v1304
	var v1310 int32
	_ = v1310
	var v1313 int32
	_ = v1313
	var v1314 int32
	_ = v1314
	var v1316 int32
	_ = v1316
	var v1320 int32
	_ = v1320
	var v1326 int32
	_ = v1326
	var v1331 int64
	_ = v1331
	var v1333 int64
	_ = v1333
	var v1334 int64
	_ = v1334
	var v1336 int64
	_ = v1336
	var v1338 int32
	_ = v1338
	var v1340 int64
	_ = v1340
	var v1341 int64
	_ = v1341
	var v1343 int32
	_ = v1343
	var v1350 int32
	_ = v1350
	var v1355 int32
	_ = v1355
	var v1356 int32
	_ = v1356
	var v1359 int32
	_ = v1359
	var v1361 int32
	_ = v1361
	var v1362 int64
	_ = v1362
	var v1366 int32
	_ = v1366
	var v1367 int64
	_ = v1367
	var v1369 int32
	_ = v1369
	var v1377 int32
	_ = v1377
	var v1378 int64
	_ = v1378
	var v1382 int32
	_ = v1382
	var v1383 int64
	_ = v1383
	var v1385 int32
	_ = v1385
	var v1393 int32
	_ = v1393
	var v1394 int32
	_ = v1394
	var v1400 int32
	_ = v1400
	var v1401 int32
	_ = v1401
	var v1403 int32
	_ = v1403
	var v1406 int32
	_ = v1406
	var v1411 int64
	_ = v1411
	var v1415 int64
	_ = v1415
	var v1417 int32
	_ = v1417
	var v1420 int64
	_ = v1420
	var v1422 int32
	_ = v1422
	var v1433 int32
	_ = v1433
	var v1441 int32
	_ = v1441
	var v1442 int32
	_ = v1442
	var v1448 int32
	_ = v1448
	var v1453 int32
	_ = v1453
	var v1458 int32
	_ = v1458
	var v1460 int32
	_ = v1460
	var v1476 int32
	_ = v1476
	var v1477 int32
	_ = v1477
	var v1478 int32
	_ = v1478
	var v1480 int32
	_ = v1480
	var v1484 int32
	_ = v1484
	var v1486 int32
	_ = v1486
	var v1491 int32
	_ = v1491
	var v1495 int32
	_ = v1495
	var v1509 int32
	_ = v1509
	var v1510 int32
	_ = v1510
	var v1511 int32
	_ = v1511
	var v1513 int32
	_ = v1513
	var v1518 int32
	_ = v1518
	var v1539 int32
	_ = v1539
	var v1541 int32
	_ = v1541
	var v1542 int32
	_ = v1542
	var v1546 int64
	_ = v1546
	var v1548 int64
	_ = v1548
	var v1549 int64
	_ = v1549
	var v1551 int64
	_ = v1551
	var v1553 int32
	_ = v1553
	var v1555 int64
	_ = v1555
	var v1556 int64
	_ = v1556
	var v1558 int32
	_ = v1558
	var v1569 int32
	_ = v1569
	var v1570 int32
	_ = v1570
	var v1573 int32
	_ = v1573
	var v1575 int32
	_ = v1575
	var v1576 int64
	_ = v1576
	var v1580 int32
	_ = v1580
	var v1581 int64
	_ = v1581
	var v1583 int32
	_ = v1583
	var v1591 int32
	_ = v1591
	var v1592 int64
	_ = v1592
	var v1596 int32
	_ = v1596
	var v1597 int64
	_ = v1597
	var v1599 int32
	_ = v1599
	var v1607 int32
	_ = v1607
	var v1609 int32
	_ = v1609
	var v1612 int32
	_ = v1612
	var v1613 int32
	_ = v1613
	var v1615 int32
	_ = v1615
	var v1618 int32
	_ = v1618
	var v1623 int64
	_ = v1623
	var v1627 int64
	_ = v1627
	var v1629 int32
	_ = v1629
	var v1632 int64
	_ = v1632
	var v1634 int32
	_ = v1634
	var v1647 int32
	_ = v1647
	var v1655 int32
	_ = v1655
	var v1656 int32
	_ = v1656
	var v1662 int32
	_ = v1662
	var v1668 int32
	_ = v1668
	var v1678 int32
	_ = v1678
	var v1680 int32
	_ = v1680
	var v1683 int32
	_ = v1683
	var v1684 int32
	_ = v1684
	var v1686 int32
	_ = v1686
	var v1694 int32
	_ = v1694
	var v1695 int32
	_ = v1695
	var v1696 int32
	_ = v1696
	var v1698 int32
	_ = v1698
	var v1704 int32
	_ = v1704
	var v1705 int32
	_ = v1705
	var v1707 int32
	_ = v1707
	var v1709 int32
	_ = v1709
	var v1711 int32
	_ = v1711
	var v1712 int32
	_ = v1712
	var v1722 int32
	_ = v1722
	var v1723 int32
	_ = v1723
	var v1724 int32
	_ = v1724
	var v1727 int32
	_ = v1727
	var v1732 int32
	_ = v1732
	var v1734 int32
	_ = v1734
	var v1736 int32
	_ = v1736
	var v1742 int32
	_ = v1742
	var v1744 int32
	_ = v1744
	var v1746 int32
	_ = v1746
	var v1751 int32
	_ = v1751
	var v1759 int32
	_ = v1759
	var v1761 int32
	_ = v1761
	var v1774 int32
	_ = v1774
	var v1776 int32
	_ = v1776
	var v1779 int32
	_ = v1779
	var v1780 int32
	_ = v1780
	var v1782 int32
	_ = v1782
	var v1790 int32
	_ = v1790
	var v1791 int32
	_ = v1791
	var v1792 int32
	_ = v1792
	var v1794 int32
	_ = v1794
	var v1798 int32
	_ = v1798
	var v1799 int32
	_ = v1799
	var v1800 int32
	_ = v1800
	var v1802 int32
	_ = v1802
	var v1816 int32
	_ = v1816
	var v1817 int32
	_ = v1817
	var v1818 int32
	_ = v1818
	var v1820 int32
	_ = v1820
	var v1822 int32
	_ = v1822
	var v1829 int32
	_ = v1829
	var v1830 int32
	_ = v1830
	var v1832 int32
	_ = v1832
	var v1838 int32
	_ = v1838
	var v1854 int32
	_ = v1854
	var v1856 int32
	_ = v1856
	var v1860 int32
	_ = v1860
	var v1871 int32
	_ = v1871
	var v1872 int32
	_ = v1872
	var v1873 int32
	_ = v1873
	var v1879 int32
	_ = v1879
	var v1889 int32
	_ = v1889
	var v1918 int32
	_ = v1918
	var v1921 int32
	_ = v1921
	var v1931 int32
	_ = v1931
	var v1938 int32
	_ = v1938
	var v1939 int32
	_ = v1939
	var v1940 int32
	_ = v1940
	var v1943 int32
	_ = v1943
	var v1946 int32
	_ = v1946
	var v1951 int32
	_ = v1951
	var v1952 int32
	_ = v1952
	var v1964 int32
	_ = v1964
	var v1969 int32
	_ = v1969
	var v1988 int32
	_ = v1988
	var v1995 int32
	_ = v1995
	var v2001 int32
	_ = v2001
	var v2011 int32
	_ = v2011
	var v2018 int32
	_ = v2018
	var v2019 int32
	_ = v2019
	var v2035 int32
	_ = v2035
	var v2039 int32
	_ = v2039
	var v2042 int32
	_ = v2042
	var v2043 int32
	_ = v2043
	var v2044 int32
	_ = v2044
	var v2045 int32
	_ = v2045
	var v2049 int32
	_ = v2049
	var v2060 int32
	_ = v2060
	var v2063 int32
	_ = v2063
	var v2068 int32
	_ = v2068
	var v2070 int32
	_ = v2070
	var v2092 int32
	_ = v2092
	var v2094 int32
	_ = v2094
	var v2095 int32
	_ = v2095
	var v2105 int32
	_ = v2105
	var v2109 int32
	_ = v2109
	var v2110 int32
	_ = v2110
	var v2117 int32
	_ = v2117
	var v2120 int32
	_ = v2120
	var v2124 int32
	_ = v2124
	var v2128 int32
	_ = v2128
	var v2129 int32
	_ = v2129
	var v2130 int32
	_ = v2130
	var v2131 int32
	_ = v2131
	var v2136 int32
	_ = v2136
	var v2140 int32
	_ = v2140
	var v2141 int32
	_ = v2141
	var v2142 int32
	_ = v2142
	var v2143 int32
	_ = v2143
	var v2148 int32
	_ = v2148
	var v2149 int32
	_ = v2149
	var v2153 int32
	_ = v2153
	var v2158 int32
	_ = v2158
	var v2162 int32
	_ = v2162
	var v2163 int64
	_ = v2163
	var v2165 int32
	_ = v2165
	var v2167 int32
	_ = v2167
	var v2174 int32
	_ = v2174
	var v2175 int32
	_ = v2175
	var v2191 int32
	_ = v2191
	var v2192 int32
	_ = v2192
	var v2195 int32
	_ = v2195
	var v2198 int32
	_ = v2198
	var v2199 int32
	_ = v2199
	var v2200 int32
	_ = v2200
	var v2201 int32
	_ = v2201
	var v2205 int32
	_ = v2205
	var v2216 int32
	_ = v2216
	var v2219 int32
	_ = v2219
	var v2225 int32
	_ = v2225
	var v2226 int32
	_ = v2226
	var v2246 int32
	_ = v2246
	var v2248 int32
	_ = v2248
	var v2250 int32
	_ = v2250
	var v2251 int32
	_ = v2251
	var v2261 int32
	_ = v2261
	var v2265 int32
	_ = v2265
	var v2266 int32
	_ = v2266
	var v2278 int32
	_ = v2278
	var v2282 int32
	_ = v2282
	var v2284 int32
	_ = v2284
	var v2289 int32
	_ = v2289
	var v2292 int32
	_ = v2292
	var v2295 int32
	_ = v2295
	var v2298 int32
	_ = v2298
	var v2303 int32
	_ = v2303
	var v2306 int32
	_ = v2306
	var v2309 int32
	_ = v2309
	var v2313 int32
	_ = v2313
	var v2321 int32
	_ = v2321
	var v2324 int32
	_ = v2324
	var v2347 int32
	_ = v2347
	var v2348 int32
	_ = v2348
	var v2350 int32
	_ = v2350
	var v2351 int32
	_ = v2351
	var v2354 int32
	_ = v2354
	var v2361 int32
	_ = v2361
	var v2363 int32
	_ = v2363
	var v2365 int32
	_ = v2365
	var v2370 int32
	_ = v2370
	var v2375 int32
	_ = v2375
	var v2376 int32
	_ = v2376
	var v2378 int32
	_ = v2378
	var v2379 int32
	_ = v2379
	var v2381 int32
	_ = v2381
	var v2384 int32
	_ = v2384
	var v2386 int32
	_ = v2386
	var v2389 int32
	_ = v2389
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = F_pg_detoast_datum(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	v16 = F_mul_size(m, int32(28), v15)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v19 = F_add_size(m, v16, int32(15))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v21 = F_palloc(m, v19)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v23 = int32(123)
	*(*uint8)(unsafe.Add(mBase, uint32(v21))) = uint8(v23)
	v25 = int32(1)
	v27 = v21 + v25
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	if v28 <= int32(0) {
		v2354 = v27
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v2361 = int32(12157)
	*(*uint16)(unsafe.Add(mBase, uint32(v2354))) = uint16(v2361)
	v2363 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	v2365 = v2354 + int32(2)
	if int32(0) <= v2363 {
		goto L355
	} else {
		goto L356
	}
L7:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
	v33 = v31 + int32(1)
	if int32(0) <= v33 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v51 = v47 + v27
	v52 = int32(58)
	*(*uint8)(unsafe.Add(mBase, uint32(v51))) = uint8(v52)
	v55 = v11 + int32(16)
	v58 = v55 + v15<<(uint(int32(2))%32)
	v59 = *(*float32)(unsafe.Add(mBase, uint32(v58)))
	v61 = v51 + int32(1)
	v62 = int32(0)
	v79 = base.I32_reinterpret_f32(v59)
	v81 = v79 & int32(8388607)
	v84 = int32(255)
	v85 = int32(base.Ui32(v79)>>(uint(int32(23))%32)) & v84
	if v81|v85 != 0 {
		goto L13
	} else {
		goto L14
	}
L9:
	;
	v43 = v33
	v44 = int32(0)
	goto L11
L10:
	;
	v38 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(v27))) = uint8(v38)
	v43 = int32(0) - v33
	v44 = int32(1)
	goto L11
L11:
	;
	v46 = F_pg_ultoa_n(m, v43, v27+v44)
	mBase = m.M
	v47 = v46 + v44
	v49 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v27+v47))) = uint8(v49)
	goto L8
L12:
	;
	v1182 = v1181 + v61
	v1183 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	if v1183 <= int32(1) {
		v2354 = v1182
		goto L6
	} else {
		goto L179
	}
L13:
	;
	v90 = base.B2i32(v85 != v84)
	goto L15
L14:
	;
	v90 = v62
	goto L15
L15:
	;
	if v90 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	if v81 != 0 {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	goto L18
L18:
	;
	if base.Ui32(int32(23)) < base.Ui32(v85-int32(127)) {
		goto L36
	} else {
		goto L37
	}
L19:
	;
	v94 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1460])))
	*(*uint8)(unsafe.Add(mBase, uint32(v61)+2)) = uint8(v94)
	v97 = int32(*(*uint16)(unsafe.Add(mBase, _consts[1461])))
	*(*uint16)(unsafe.Add(mBase, uint32(v61))) = uint16(v97)
	v1181 = int32(3)
	goto L12
L20:
	;
	goto L21
L21:
	;
	if v79 < int32(0) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v102 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(v61))) = uint8(v102)
	goto L24
L23:
	;
	goto L24
L24:
	;
	v106 = v61 + int32(base.Ui32(v79)>>(uint(int32(31))%32))
	if v85 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v106))) = int64(8751735898823355977)
	if v79 < int32(0) {
		goto L28
	} else {
		goto L29
	}
L26:
	;
	goto L27
L27:
	;
	v114 = int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v106))) = uint8(v114)
	if v79 < int32(0) {
		goto L31
	} else {
		goto L32
	}
L28:
	;
	v113 = int32(9)
	goto L30
L29:
	;
	v113 = int32(8)
	goto L30
L30:
	;
	v1181 = v113
	goto L12
L31:
	;
	v120 = int32(2)
	goto L33
L32:
	;
	v120 = int32(1)
	goto L33
L33:
	;
	v1181 = v120
	goto L12
L34:
	;
	v773 = v772 + v765
	v774 = int32(0)
	if v79 < v774 {
		goto L104
	} else {
		goto L105
	}
L35:
	;
	if base.Ui32(int32(9999999)) < base.Ui32(v713) {
		v755 = v713
		v765 = v723
		v772 = int32(8)
		goto L34
	} else {
		goto L95
	}
L36:
	;
	v138 = int32(2)
	v144 = v81 << (uint(v138) % 32)
	if v85 != 0 {
		goto L39
	} else {
		goto L40
	}
L37:
	;
	v125 = int32(-1)
	v127 = int32(150) - v85
	if v81&(v125<<(uint(v127)%32)^v125) != 0 {
		goto L36
	} else {
		goto L38
	}
L38:
	;
	v713 = int32(base.Ui32(v81|int32(8388608)) >> (uint(v127) % 32))
	v723 = v62
	goto L35
L39:
	;
	v147 = v144 | int32(33554432)
	goto L41
L40:
	;
	v147 = v144
	goto L41
L41:
	;
	v148 = base.B2i32(v81 != int32(0)) | base.B2i32(base.Ui32(v85) < base.Ui32(v138)) ^ int32(-1) + v147
	v150 = v147 | int32(2)
	if v85 != 0 {
		goto L45
	} else {
		goto L46
	}
L42:
	;
	v706 = v688 + v694
	v707 = v705 + v690
	if base.Ui32(v707) <= base.Ui32(int32(99999999)) {
		v713 = v707
		v723 = v706
		goto L35
	} else {
		goto L94
	}
L43:
	;
	v624 = int32(0)
	v625 = int32(10)
	v626 = base.I32_div_u_s(v614, v625)
	v628 = base.I32_div_u_s(v616, v625)
	if base.Ui32(v628) < base.Ui32(v626) {
		goto L88
	} else {
		goto L89
	}
L44:
	;
	v528 = int32(0)
	v529 = int32(10)
	v530 = base.I32_div_u_s(v518, v529)
	v532 = base.I32_div_u_s(v520, v529)
	if base.Ui32(v530) <= base.Ui32(v532) {
		goto L82
	} else {
		goto L83
	}
L45:
	;
	v154 = v85 - int32(152)
	goto L47
L46:
	;
	v154 = int32(-151)
	goto L47
L47:
	;
	if int32(0) <= v154 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v160 = int32(base.Ui32(v154*int32(78913)) >> (uint(int32(18)) % 32))
	v165 = *(*int64)(unsafe.Add(mBase, uint32(v160<<(uint(int32(3))%32))+uint32(_consts[1462])))
	v167 = v165 & int64(4294967295)
	v168 = base.I64_extend_i32_u(v148)
	v170 = int64(32)
	v172 = base.I32_wrap_i64(int64(base.Ui64(v167*v168) >> (uint(v170) % 64)))
	v174 = int64(base.Ui64(v165) >> (uint(v170) % 64))
	v175 = v174 * v168
	v177 = v172 + base.I32_wrap_i64(v175)
	v184 = v160 - v154
	v189 = v184 + int32(base.Ui32(v160*int32(1217359))>>(uint(int32(19))%32))
	v190 = int32(5) - v189
	v193 = v189 + int32(27)
	v195 = (base.B2i32(base.Ui32(v177) < base.Ui32(v172))+base.I32_wrap_i64(int64(base.Ui64(v175)>>(uint(v170)%64))))<<(uint(v190)%32) | int32(base.Ui32(v177)>>(uint(v193)%32))
	v196 = base.I64_extend_i32_u(v150)
	v200 = base.I32_wrap_i64(int64(base.Ui64(v167*v196) >> (uint(v170) % 64)))
	v201 = v196 * v174
	v203 = v200 + base.I32_wrap_i64(v201)
	v211 = (base.B2i32(base.Ui32(v203) < base.Ui32(v200))+base.I32_wrap_i64(int64(base.Ui64(v201)>>(uint(v170)%64))))<<(uint(v190)%32) | int32(base.Ui32(v203)>>(uint(v193)%32))
	v212 = base.I64_extend_i32_u(v147)
	v216 = base.I32_wrap_i64(int64(base.Ui64(v167*v212) >> (uint(v170) % 64)))
	v217 = v212 * v174
	v219 = v216 + base.I32_wrap_i64(v217)
	v227 = (base.B2i32(base.Ui32(v219) < base.Ui32(v216))+base.I32_wrap_i64(int64(base.Ui64(v217)>>(uint(v170)%64))))<<(uint(v190)%32) | int32(base.Ui32(v219)>>(uint(v193)%32))
	v228 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v154) {
		goto L51
	} else {
		goto L52
	}
L49:
	;
	goto L50
L50:
	;
	v373 = v154 * int32(-732923)
	v375 = int32(base.Ui32(v373) >> (uint(int32(20)) % 32))
	v376 = v375 + v154
	v380 = *(*int64)(unsafe.Add(mBase, uint32(int32(1872416)-v376<<(uint(int32(3))%32))))
	v382 = v380 & int64(4294967295)
	v383 = base.I64_extend_i32_u(v148)
	v385 = int64(32)
	v387 = base.I32_wrap_i64(int64(base.Ui64(v382*v383) >> (uint(v385) % 64)))
	v389 = int64(base.Ui64(v380) >> (uint(v385) % 64))
	v390 = v389 * v383
	v392 = v387 + base.I32_wrap_i64(v390)
	v403 = v375 - int32(base.Ui32(v376*int32(-1217359))>>(uint(int32(19))%32))
	v404 = int32(4) - v403
	v407 = v403 + int32(28)
	v409 = (base.B2i32(base.Ui32(v392) < base.Ui32(v387))+base.I32_wrap_i64(int64(base.Ui64(v390)>>(uint(v385)%64))))<<(uint(v404)%32) | int32(base.Ui32(v392)>>(uint(v407)%32))
	v410 = base.I64_extend_i32_u(v147)
	v414 = base.I32_wrap_i64(int64(base.Ui64(v382*v410) >> (uint(v385) % 64)))
	v415 = v410 * v389
	v417 = v414 + base.I32_wrap_i64(v415)
	v425 = (base.B2i32(base.Ui32(v417) < base.Ui32(v414))+base.I32_wrap_i64(int64(base.Ui64(v415)>>(uint(v385)%64))))<<(uint(v404)%32) | int32(base.Ui32(v417)>>(uint(v407)%32))
	v426 = base.I64_extend_i32_u(v150)
	v430 = base.I32_wrap_i64(int64(base.Ui64(v382*v426) >> (uint(v385) % 64)))
	v431 = v389 * v426
	v433 = v430 + base.I32_wrap_i64(v431)
	v441 = (base.B2i32(base.Ui32(v433) < base.Ui32(v430))+base.I32_wrap_i64(int64(base.Ui64(v431)>>(uint(v385)%64))))<<(uint(v404)%32) | int32(base.Ui32(v433)>>(uint(v407)%32))
	v443 = v441 - int32(1)
	if base.Ui32(int32(1048576)) <= base.Ui32(v373) {
		goto L72
	} else {
		goto L73
	}
L51:
	;
	v234 = int32(10)
	v235 = base.I32_div_u_s(v211-int32(1), v234)
	v237 = base.I32_div_u_s(v195, v234)
	if base.Ui32(v235) <= base.Ui32(v237) {
		goto L54
	} else {
		goto L55
	}
L52:
	;
	v282 = v228
	goto L53
L53:
	;
	v287 = base.I32_rem_u_s(v147, int32(5))
	if v287 == int32(0) {
		goto L58
	} else {
		goto L59
	}
L54:
	;
	v240 = v160 - int32(1)
	v245 = *(*int64)(unsafe.Add(mBase, uint32(v240<<(uint(int32(3))%32))+uint32(_consts[1462])))
	v249 = int64(32)
	v251 = base.I32_wrap_i64(int64(base.Ui64(v245&int64(4294967295)*v212) >> (uint(v249) % 64)))
	v254 = int64(base.Ui64(v245)>>(uint(v249)%64)) * v212
	v256 = v251 + base.I32_wrap_i64(v254)
	v267 = v184 + int32(base.Ui32(v240*int32(1217359))>>(uint(int32(19))%32))
	v275 = base.I32_rem_u_s((base.B2i32(base.Ui32(v256) < base.Ui32(v251))+base.I32_wrap_i64(int64(base.Ui64(v254)>>(uint(v249)%64))))<<(uint(int32(6)-v267)%32)|int32(base.Ui32(v256)>>(uint(v267+int32(26))%32)), int32(10))
	v276 = v275
	goto L56
L55:
	;
	v276 = v228
	goto L56
L56:
	;
	if base.Ui32(int32(33)) < base.Ui32(v154) {
		v608 = v227
		v610 = v276
		v613 = v160
		v614 = v211
		v616 = v195
		goto L43
	} else {
		goto L57
	}
L57:
	;
	v282 = v276
	goto L53
L58:
	;
	v292 = v147
	v294 = v228
	goto L61
L59:
	;
	goto L60
L60:
	;
	v318 = int32(0)
	v320 = base.I32_rem_u_s(v150, int32(5))
	if v320 == v318 {
		goto L65
	} else {
		goto L66
	}
L61:
	;
	v310 = v294 + int32(1)
	v311 = int32(5)
	v312 = base.I32_div_u_s(v292, v311)
	v314 = base.I32_rem_u_s(v312, v311)
	if v314 == int32(0) {
		v292 = v312
		v294 = v310
		goto L61
	} else {
		goto L63
	}
L62:
	;
	if base.Ui32(v310) < base.Ui32(v160) {
		v608 = v227
		v610 = v282
		v613 = v160
		v614 = v211
		v616 = v195
		goto L43
	} else {
		goto L64
	}
L63:
	;
	goto L62
L64:
	;
	v512 = v227
	v514 = v282
	v517 = v160
	v518 = v211
	v520 = v195
	goto L44
L65:
	;
	v325 = v318
	v329 = v150
	goto L68
L66:
	;
	v352 = v318
	goto L67
L67:
	;
	v608 = v227
	v610 = v282
	v613 = v160
	v614 = v211 - base.B2i32(base.Ui32(v160) <= base.Ui32(v352))
	v616 = v195
	goto L43
L68:
	;
	v343 = v325 + int32(1)
	v344 = int32(5)
	v345 = base.I32_div_u_s(v329, v344)
	v347 = base.I32_rem_u_s(v345, v344)
	if v347 == int32(0) {
		v325 = v343
		v329 = v345
		goto L68
	} else {
		goto L70
	}
L69:
	;
	v352 = v343
	goto L67
L70:
	;
	goto L69
L71:
	;
	if base.Ui32(int32(32505855)) < base.Ui32(v373) {
		v608 = v425
		v610 = v490
		v613 = v376
		v614 = v441
		v616 = v409
		goto L43
	} else {
		goto L79
	}
L72:
	;
	v446 = int32(10)
	v447 = base.I32_div_u_s(v443, v446)
	v449 = base.I32_div_u_s(v409, v446)
	if base.Ui32(v447) <= base.Ui32(v449) {
		goto L75
	} else {
		goto L76
	}
L73:
	;
	v496 = v62
	goto L74
L74:
	;
	v512 = v425
	v514 = v496
	v517 = v376
	v518 = v443
	v520 = v409
	goto L44
L75:
	;
	v452 = int32(1) - v376
	v457 = *(*int64)(unsafe.Add(mBase, uint32(v452<<(uint(int32(3))%32))+uint32(_consts[1463])))
	v461 = int64(32)
	v463 = base.I32_wrap_i64(int64(base.Ui64(v457&int64(4294967295)*v410) >> (uint(v461) % 64)))
	v466 = int64(base.Ui64(v457)>>(uint(v461)%64)) * v410
	v468 = v463 + base.I32_wrap_i64(v466)
	v481 = v375 + (int32(base.Ui32(v452*int32(1217359))>>(uint(int32(19))%32)) ^ int32(-1))
	v489 = base.I32_rem_u_s((base.B2i32(base.Ui32(v468) < base.Ui32(v463))+base.I32_wrap_i64(int64(base.Ui64(v466)>>(uint(v461)%64))))<<(uint(int32(4)-v481)%32)|int32(base.Ui32(v468)>>(uint(v481+int32(28))%32)), int32(10))
	v490 = v489
	goto L77
L76:
	;
	v490 = v62
	goto L77
L77:
	;
	if base.Ui32(int32(2097151)) < base.Ui32(v373) {
		goto L71
	} else {
		goto L78
	}
L78:
	;
	v496 = v490
	goto L74
L79:
	;
	v502 = int32(-1)
	if v147&(v502<<(uint(v375-int32(1))%32)^v502) != 0 {
		v608 = v425
		v610 = v490
		v613 = v376
		v614 = v441
		v616 = v409
		goto L43
	} else {
		goto L80
	}
L80:
	;
	v512 = v425
	v514 = v490
	v517 = v376
	v518 = v441
	v520 = v409
	goto L44
L81:
	;
	v595 = v580 & int32(255)
	v688 = v576
	v690 = v578
	v694 = v517
	v705 = (v593|base.B2i32(v595 != int32(5))|v578)&base.B2i32(base.Ui32(int32(4)) < base.Ui32(v595)) | base.B2i32(v578 == v585)
	goto L42
L82:
	;
	v576 = v528
	v578 = v512
	v580 = v514
	v585 = v520
	v593 = int32(0)
	goto L81
L83:
	;
	goto L84
L84:
	;
	v538 = v528
	v539 = v512
	v541 = v514
	v543 = v530
	v545 = v532
	v546 = int32(1)
	goto L85
L85:
	;
	v556 = v538 + int32(1)
	v557 = int32(10)
	v558 = base.I32_div_u_s(v539, v557)
	v561 = v539 - v558*v557
	v566 = v546 & base.B2i32(v541&int32(255) == int32(0))
	v568 = base.I32_div_u_s(v543, v557)
	v570 = base.I32_div_u_s(v545, v557)
	if base.Ui32(v570) < base.Ui32(v568) {
		v538 = v556
		v539 = v558
		v541 = v561
		v543 = v568
		v545 = v570
		v546 = v566
		goto L85
	} else {
		goto L87
	}
L86:
	;
	v576 = v556
	v578 = v558
	v580 = v561
	v585 = v545
	v593 = v566 ^ int32(1)
	goto L81
L87:
	;
	goto L86
L88:
	;
	v632 = v624
	v633 = v608
	v634 = v626
	v636 = v628
	goto L91
L89:
	;
	v663 = v624
	v664 = v608
	v666 = v610
	v672 = v616
	goto L90
L90:
	;
	v688 = v663
	v690 = v664
	v694 = v613
	v705 = base.B2i32(v664 == v672) | base.B2i32(base.Ui32(int32(4)) < base.Ui32(v666&int32(255)))
	goto L42
L91:
	;
	v650 = v632 + int32(1)
	v651 = int32(10)
	v652 = base.I32_div_u_s(v633, v651)
	v654 = base.I32_div_u_s(v634, v651)
	v656 = base.I32_div_u_s(v636, v651)
	if base.Ui32(v656) < base.Ui32(v654) {
		v632 = v650
		v633 = v652
		v634 = v654
		v636 = v656
		goto L91
	} else {
		goto L93
	}
L92:
	;
	v663 = v650
	v664 = v652
	v666 = v633 - v652*int32(10)
	v672 = v636
	goto L90
L93:
	;
	goto L92
L94:
	;
	v755 = v707
	v765 = v706
	v772 = int32(9)
	goto L34
L95:
	;
	if base.Ui32(int32(999999)) < base.Ui32(v713) {
		v755 = v713
		v765 = v723
		v772 = int32(7)
		goto L34
	} else {
		goto L96
	}
L96:
	;
	if base.Ui32(int32(99999)) < base.Ui32(v713) {
		v755 = v713
		v765 = v723
		v772 = int32(6)
		goto L34
	} else {
		goto L97
	}
L97:
	;
	if base.Ui32(int32(9999)) < base.Ui32(v713) {
		v755 = v713
		v765 = v723
		v772 = int32(5)
		goto L34
	} else {
		goto L98
	}
L98:
	;
	if base.Ui32(int32(999)) < base.Ui32(v713) {
		v755 = v713
		v765 = v723
		v772 = int32(4)
		goto L34
	} else {
		goto L99
	}
L99:
	;
	if base.Ui32(int32(99)) < base.Ui32(v713) {
		v755 = v713
		v765 = v723
		v772 = int32(3)
		goto L34
	} else {
		goto L100
	}
L100:
	;
	if base.Ui32(int32(9)) < base.Ui32(v713) {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	v752 = int32(2)
	goto L103
L102:
	;
	v752 = int32(1)
	goto L103
L103:
	;
	v755 = v713
	v765 = v723
	v772 = v752
	goto L34
L104:
	;
	v777 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(v61))) = uint8(v777)
	v780 = int32(1)
	goto L106
L105:
	;
	v780 = v774
	goto L106
L106:
	;
	if base.Ui32(v773+int32(3)) <= base.Ui32(int32(9)) {
		goto L109
	} else {
		goto L110
	}
L107:
	;
	v1001 = int32(0)
	if base.Ui32(v755) < base.Ui32(int32(10000)) {
		goto L149
	} else {
		goto L150
	}
L108:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v785))) = v997
	v999 = v996
	goto L107
L109:
	;
	v785 = v61 + v780
	v786 = int32(0)
	if v773 <= v786 {
		goto L112
	} else {
		goto L113
	}
L110:
	;
	goto L111
L111:
	;
	if v765 != 0 {
		goto L117
	} else {
		goto L118
	}
L112:
	;
	v996 = int32(2) - v773
	v997 = int64(3472328296227679792)
	goto L108
L113:
	;
	goto L114
L114:
	;
	if int32(0) <= v765 {
		v996 = v786
		v997 = int64(3472328296227680304)
		goto L108
	} else {
		goto L115
	}
L115:
	;
	v999 = int32(1)
	goto L107
L116:
	;
	v845 = int32(0)
	if base.Ui32(v829) < base.Ui32(int32(10000)) {
		goto L125
	} else {
		goto L126
	}
L117:
	;
	v829 = v755
	v835 = v772
	goto L116
L118:
	;
	goto L119
L119:
	;
	v798 = v755
	v803 = v772
	goto L120
L120:
	;
	if v798&int32(1) != 0 {
		v829 = v798
		v835 = v803
		goto L116
	} else {
		goto L122
	}
L121:
	;
	v829 = v798
	v835 = v803
	goto L116
L122:
	;
	v822 = base.I32_div_u_s(v798, int32(10))
	if int32(0)-v798 == v822*int32(-10) {
		v798 = v822
		v803 = v803 - int32(1)
		goto L120
	} else {
		goto L123
	}
L123:
	;
	goto L121
L124:
	;
	if base.Ui32(v904) < base.Ui32(int32(100)) {
		goto L132
	} else {
		goto L133
	}
L125:
	;
	v902 = v845
	v904 = v829
	goto L124
L126:
	;
	goto L127
L127:
	;
	v852 = v845
	v853 = v829
	goto L128
L128:
	;
	v869 = v61 + v780 + v835 - v852
	v873 = base.I32_div_u_s(v853, int32(10000))
	v876 = v873*int32(-10000) + v853
	v877 = int32(100)
	v878 = base.I32_div_u_s(v876, v877)
	v879 = int32(1)
	v883 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v878<<(uint(v879)%32))+uint32(_consts[1464]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v869-int32(3)))) = uint16(v883)
	v894 = int32(*(*uint16)(unsafe.Add(mBase, uint32((v876-v878*v877)<<(uint(v879)%32))+uint32(_consts[1464]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v869-v879))) = uint16(v894)
	v897 = v852 + int32(4)
	if base.Ui32(int32(99999999)) < base.Ui32(v853) {
		v852 = v897
		v853 = v873
		goto L128
	} else {
		goto L130
	}
L129:
	;
	v902 = v897
	v904 = v873
	goto L124
L130:
	;
	goto L129
L131:
	;
	if base.Ui32(int32(10)) <= base.Ui32(v944) {
		goto L136
	} else {
		goto L137
	}
L132:
	;
	v943 = v902
	v944 = v904
	goto L131
L133:
	;
	goto L134
L134:
	;
	v926 = int32(65535)
	v928 = int32(100)
	v929 = base.I32_div_u_s(v904&v926, v928)
	v939 = int32(*(*uint16)(unsafe.Add(mBase, uint32((v904-v929*v928)&v926<<(uint(int32(1))%32))+uint32(_consts[1464]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v61+v780+v835+(v902^int32(-1))))) = uint16(v939)
	v943 = v902 | int32(2)
	v944 = v929
	goto L131
L135:
	;
	v963 = int32(1)
	v964 = v773 - v963
	v965 = v61 + v780
	*(*uint8)(unsafe.Add(mBase, uint32(v965))) = uint8(v962)
	if base.Ui32(int32(2)) <= base.Ui32(v835) {
		goto L139
	} else {
		goto L140
	}
L136:
	;
	v951 = v944 << (uint(int32(1)) % 32)
	v954 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v951)+uint32(_consts[1465]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v61+(v780+v835-v943)))) = uint8(v954)
	v958 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v951)+uint32(_consts[1464]))))
	v962 = v958
	goto L135
L137:
	;
	goto L138
L138:
	;
	v962 = v944 | int32(48)
	goto L135
L139:
	;
	v970 = int32(46)
	*(*uint8)(unsafe.Add(mBase, uint32(v965)+1)) = uint8(v970)
	v974 = v835 + int32(1)
	goto L141
L140:
	;
	v974 = v963
	goto L141
L141:
	;
	v975 = v974 + v780
	v976 = v61 + v975
	v977 = int32(101)
	*(*uint8)(unsafe.Add(mBase, uint32(v976))) = uint8(v977)
	v982 = base.B2i32(v964 < int32(0))
	if v964 < int32(0) {
		goto L142
	} else {
		goto L143
	}
L142:
	;
	v983 = int32(45)
	goto L144
L143:
	;
	v983 = int32(43)
	goto L144
L144:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v976)+1)) = uint8(v983)
	if v964 < int32(0) {
		goto L145
	} else {
		goto L146
	}
L145:
	;
	v987 = int32(1) - v773
	goto L147
L146:
	;
	v987 = v964
	goto L147
L147:
	;
	v992 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v987<<(uint(int32(1))%32))+uint32(_consts[1464]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v976)+2)) = uint16(v992)
	v1181 = v975 + int32(4)
	goto L12
L148:
	;
	if base.Ui32(v1060) < base.Ui32(int32(100)) {
		goto L156
	} else {
		goto L157
	}
L149:
	;
	v1059 = v1001
	v1060 = v755
	goto L148
L150:
	;
	goto L151
L151:
	;
	v1008 = v755
	v1009 = v1001
	goto L152
L152:
	;
	v1025 = v785 + v999 + v772 - v1009
	v1026 = int32(4)
	v1029 = base.I32_div_u_s(v1008, int32(10000))
	v1032 = v1029*int32(-10000) + v1008
	v1033 = int32(100)
	v1034 = base.I32_div_u_s(v1032, v1033)
	v1035 = int32(1)
	v1039 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1034<<(uint(v1035)%32))+uint32(_consts[1464]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v1025-v1026))) = uint16(v1039)
	v1050 = int32(*(*uint16)(unsafe.Add(mBase, uint32((v1032-v1034*v1033)<<(uint(v1035)%32))+uint32(_consts[1464]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v1025-int32(2)))) = uint16(v1050)
	v1053 = v1009 + v1026
	if base.Ui32(int32(99999999)) < base.Ui32(v1008) {
		v1008 = v1029
		v1009 = v1053
		goto L152
	} else {
		goto L154
	}
L153:
	;
	v1059 = v1053
	v1060 = v1029
	goto L148
L154:
	;
	goto L153
L155:
	;
	if base.Ui32(int32(10)) <= base.Ui32(v1099) {
		goto L160
	} else {
		goto L161
	}
L156:
	;
	v1099 = v1060
	v1100 = v1059
	goto L155
L157:
	;
	goto L158
L158:
	;
	v1080 = int32(2)
	v1082 = int32(65535)
	v1084 = int32(100)
	v1085 = base.I32_div_u_s(v1060&v1082, v1084)
	v1095 = int32(*(*uint16)(unsafe.Add(mBase, uint32((v1060-v1085*v1084)&v1082<<(uint(int32(1))%32))+uint32(_consts[1464]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v785+v999+v772-v1059-v1080))) = uint16(v1095)
	v1099 = v1085
	v1100 = v1059 | v1080
	goto L155
L159:
	;
	v1118 = int32(1)
	if v999 == v1118 {
		goto L164
	} else {
		goto L165
	}
L160:
	;
	v1112 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1099<<(uint(int32(1))%32))+uint32(_consts[1464]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v785+v999+v772-v1100-int32(2)))) = uint16(v1112)
	goto L159
L161:
	;
	goto L162
L162:
	;
	v1116 = v1099 | int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v785+v999))) = uint8(v1116)
	goto L159
L163:
	;
	v1181 = v1158 + int32(base.Ui32(v79)>>(uint(int32(31))%32))
	goto L12
L164:
	;
	if v773&int32(4) != 0 {
		goto L167
	} else {
		goto L168
	}
L165:
	;
	goto L166
L166:
	;
	if v765 < int32(0) {
		goto L176
	} else {
		goto L177
	}
L167:
	;
	v1123 = *(*int32)(unsafe.Add(mBase, uint32(v785)+1))
	*(*int32)(unsafe.Add(mBase, uint32(v785))) = v1123
	v1126 = int32(5)
	goto L169
L168:
	;
	v1126 = v1118
	goto L169
L169:
	;
	if v773&int32(2) != 0 {
		goto L170
	} else {
		goto L171
	}
L170:
	;
	v1129 = v785 + v1126
	v1132 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1129))))
	*(*uint16)(unsafe.Add(mBase, uint32(v1129-int32(1)))) = uint16(v1132)
	v1137 = v1126 | int32(2)
	goto L172
L171:
	;
	v1137 = v1126
	goto L172
L172:
	;
	if v773&int32(1) != 0 {
		goto L173
	} else {
		goto L174
	}
L173:
	;
	v1140 = v785 + v1137
	v1143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1140))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1140-int32(1)))) = uint8(v1143)
	goto L175
L174:
	;
	goto L175
L175:
	;
	v1147 = int32(46)
	*(*uint8)(unsafe.Add(mBase, uint32(v785+v773))) = uint8(v1147)
	v1158 = v772 + int32(1)
	goto L163
L176:
	;
	v1155 = int32(2) - v765
	goto L178
L177:
	;
	v1155 = v773
	goto L178
L178:
	;
	v1158 = v1155
	goto L163
L179:
	;
	v1187 = v1182
	v1190 = v25
	goto L180
L180:
	;
	v1194 = int32(44)
	*(*uint8)(unsafe.Add(mBase, uint32(v1187))) = uint8(v1194)
	v1197 = v1190 << (uint(int32(2)) % 32)
	v1199 = *(*int32)(unsafe.Add(mBase, uint32(v55+v1197)))
	v1200 = int32(1)
	v1201 = v1199 + v1200
	v1203 = v1187 + v1200
	if int32(0) <= v1201 {
		goto L183
	} else {
		goto L184
	}
L181:
	;
	v2354 = v2348
	goto L6
L182:
	;
	v1221 = v1217 + v1203
	v1222 = int32(58)
	*(*uint8)(unsafe.Add(mBase, uint32(v1221))) = uint8(v1222)
	v1225 = *(*float32)(unsafe.Add(mBase, uint32(v1197+v58)))
	v1227 = v1221 + int32(1)
	v1228 = int32(0)
	v1245 = base.I32_reinterpret_f32(v1225)
	v1247 = v1245 & int32(8388607)
	v1250 = int32(255)
	v1251 = int32(base.Ui32(v1245)>>(uint(int32(23))%32)) & v1250
	if v1247|v1251 != 0 {
		goto L187
	} else {
		goto L188
	}
L183:
	;
	v1213 = v1201
	v1214 = int32(0)
	goto L185
L184:
	;
	v1208 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(v1203))) = uint8(v1208)
	v1213 = int32(0) - v1201
	v1214 = int32(1)
	goto L185
L185:
	;
	v1216 = F_pg_ultoa_n(m, v1213, v1203+v1214)
	mBase = m.M
	v1217 = v1216 + v1214
	v1219 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1203+v1217))) = uint8(v1219)
	goto L182
L186:
	;
	v2348 = v2347 + v1227
	v2350 = v1190 + int32(1)
	v2351 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	if v2350 < v2351 {
		v1187 = v2348
		v1190 = v2350
		goto L180
	} else {
		goto L353
	}
L187:
	;
	v1256 = base.B2i32(v1251 != v1250)
	goto L189
L188:
	;
	v1256 = v1228
	goto L189
L189:
	;
	if v1256 == int32(0) {
		goto L190
	} else {
		goto L191
	}
L190:
	;
	if v1247 != 0 {
		goto L193
	} else {
		goto L194
	}
L191:
	;
	goto L192
L192:
	;
	if base.Ui32(int32(23)) < base.Ui32(v1251-int32(127)) {
		goto L210
	} else {
		goto L211
	}
L193:
	;
	v1260 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1460])))
	*(*uint8)(unsafe.Add(mBase, uint32(v1227)+2)) = uint8(v1260)
	v1263 = int32(*(*uint16)(unsafe.Add(mBase, _consts[1461])))
	*(*uint16)(unsafe.Add(mBase, uint32(v1227))) = uint16(v1263)
	v2347 = int32(3)
	goto L186
L194:
	;
	goto L195
L195:
	;
	if v1245 < int32(0) {
		goto L196
	} else {
		goto L197
	}
L196:
	;
	v1268 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(v1227))) = uint8(v1268)
	goto L198
L197:
	;
	goto L198
L198:
	;
	v1272 = v1227 + int32(base.Ui32(v1245)>>(uint(int32(31))%32))
	if v1251 != 0 {
		goto L199
	} else {
		goto L200
	}
L199:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1272))) = int64(8751735898823355977)
	if v1245 < int32(0) {
		goto L202
	} else {
		goto L203
	}
L200:
	;
	goto L201
L201:
	;
	v1280 = int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v1272))) = uint8(v1280)
	if v1245 < int32(0) {
		goto L205
	} else {
		goto L206
	}
L202:
	;
	v1279 = int32(9)
	goto L204
L203:
	;
	v1279 = int32(8)
	goto L204
L204:
	;
	v2347 = v1279
	goto L186
L205:
	;
	v1286 = int32(2)
	goto L207
L206:
	;
	v1286 = int32(1)
	goto L207
L207:
	;
	v2347 = v1286
	goto L186
L208:
	;
	v1939 = v1938 + v1931
	v1940 = int32(0)
	if v1245 < v1940 {
		goto L278
	} else {
		goto L279
	}
L209:
	;
	if base.Ui32(int32(9999999)) < base.Ui32(v1879) {
		v1921 = v1879
		v1931 = v1889
		v1938 = int32(8)
		goto L208
	} else {
		goto L269
	}
L210:
	;
	v1304 = int32(2)
	v1310 = v1247 << (uint(v1304) % 32)
	if v1251 != 0 {
		goto L213
	} else {
		goto L214
	}
L211:
	;
	v1291 = int32(-1)
	v1293 = int32(150) - v1251
	if v1247&(v1291<<(uint(v1293)%32)^v1291) != 0 {
		goto L210
	} else {
		goto L212
	}
L212:
	;
	v1879 = int32(base.Ui32(v1247|int32(8388608)) >> (uint(v1293) % 32))
	v1889 = v1228
	goto L209
L213:
	;
	v1313 = v1310 | int32(33554432)
	goto L215
L214:
	;
	v1313 = v1310
	goto L215
L215:
	;
	v1314 = base.B2i32(v1247 != int32(0)) | base.B2i32(base.Ui32(v1251) < base.Ui32(v1304)) ^ int32(-1) + v1313
	v1316 = v1313 | int32(2)
	if v1251 != 0 {
		goto L219
	} else {
		goto L220
	}
L216:
	;
	v1872 = v1854 + v1860
	v1873 = v1871 + v1856
	if base.Ui32(v1873) <= base.Ui32(int32(99999999)) {
		v1879 = v1873
		v1889 = v1872
		goto L209
	} else {
		goto L268
	}
L217:
	;
	v1790 = int32(0)
	v1791 = int32(10)
	v1792 = base.I32_div_u_s(v1780, v1791)
	v1794 = base.I32_div_u_s(v1782, v1791)
	if base.Ui32(v1794) < base.Ui32(v1792) {
		goto L262
	} else {
		goto L263
	}
L218:
	;
	v1694 = int32(0)
	v1695 = int32(10)
	v1696 = base.I32_div_u_s(v1684, v1695)
	v1698 = base.I32_div_u_s(v1686, v1695)
	if base.Ui32(v1696) <= base.Ui32(v1698) {
		goto L256
	} else {
		goto L257
	}
L219:
	;
	v1320 = v1251 - int32(152)
	goto L221
L220:
	;
	v1320 = int32(-151)
	goto L221
L221:
	;
	if int32(0) <= v1320 {
		goto L222
	} else {
		goto L223
	}
L222:
	;
	v1326 = int32(base.Ui32(v1320*int32(78913)) >> (uint(int32(18)) % 32))
	v1331 = *(*int64)(unsafe.Add(mBase, uint32(v1326<<(uint(int32(3))%32))+uint32(_consts[1462])))
	v1333 = v1331 & int64(4294967295)
	v1334 = base.I64_extend_i32_u(v1314)
	v1336 = int64(32)
	v1338 = base.I32_wrap_i64(int64(base.Ui64(v1333*v1334) >> (uint(v1336) % 64)))
	v1340 = int64(base.Ui64(v1331) >> (uint(v1336) % 64))
	v1341 = v1340 * v1334
	v1343 = v1338 + base.I32_wrap_i64(v1341)
	v1350 = v1326 - v1320
	v1355 = v1350 + int32(base.Ui32(v1326*int32(1217359))>>(uint(int32(19))%32))
	v1356 = int32(5) - v1355
	v1359 = v1355 + int32(27)
	v1361 = (base.B2i32(base.Ui32(v1343) < base.Ui32(v1338))+base.I32_wrap_i64(int64(base.Ui64(v1341)>>(uint(v1336)%64))))<<(uint(v1356)%32) | int32(base.Ui32(v1343)>>(uint(v1359)%32))
	v1362 = base.I64_extend_i32_u(v1316)
	v1366 = base.I32_wrap_i64(int64(base.Ui64(v1333*v1362) >> (uint(v1336) % 64)))
	v1367 = v1362 * v1340
	v1369 = v1366 + base.I32_wrap_i64(v1367)
	v1377 = (base.B2i32(base.Ui32(v1369) < base.Ui32(v1366))+base.I32_wrap_i64(int64(base.Ui64(v1367)>>(uint(v1336)%64))))<<(uint(v1356)%32) | int32(base.Ui32(v1369)>>(uint(v1359)%32))
	v1378 = base.I64_extend_i32_u(v1313)
	v1382 = base.I32_wrap_i64(int64(base.Ui64(v1333*v1378) >> (uint(v1336) % 64)))
	v1383 = v1378 * v1340
	v1385 = v1382 + base.I32_wrap_i64(v1383)
	v1393 = (base.B2i32(base.Ui32(v1385) < base.Ui32(v1382))+base.I32_wrap_i64(int64(base.Ui64(v1383)>>(uint(v1336)%64))))<<(uint(v1356)%32) | int32(base.Ui32(v1385)>>(uint(v1359)%32))
	v1394 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v1320) {
		goto L225
	} else {
		goto L226
	}
L223:
	;
	goto L224
L224:
	;
	v1539 = v1320 * int32(-732923)
	v1541 = int32(base.Ui32(v1539) >> (uint(int32(20)) % 32))
	v1542 = v1541 + v1320
	v1546 = *(*int64)(unsafe.Add(mBase, uint32(int32(1872416)-v1542<<(uint(int32(3))%32))))
	v1548 = v1546 & int64(4294967295)
	v1549 = base.I64_extend_i32_u(v1314)
	v1551 = int64(32)
	v1553 = base.I32_wrap_i64(int64(base.Ui64(v1548*v1549) >> (uint(v1551) % 64)))
	v1555 = int64(base.Ui64(v1546) >> (uint(v1551) % 64))
	v1556 = v1555 * v1549
	v1558 = v1553 + base.I32_wrap_i64(v1556)
	v1569 = v1541 - int32(base.Ui32(v1542*int32(-1217359))>>(uint(int32(19))%32))
	v1570 = int32(4) - v1569
	v1573 = v1569 + int32(28)
	v1575 = (base.B2i32(base.Ui32(v1558) < base.Ui32(v1553))+base.I32_wrap_i64(int64(base.Ui64(v1556)>>(uint(v1551)%64))))<<(uint(v1570)%32) | int32(base.Ui32(v1558)>>(uint(v1573)%32))
	v1576 = base.I64_extend_i32_u(v1313)
	v1580 = base.I32_wrap_i64(int64(base.Ui64(v1548*v1576) >> (uint(v1551) % 64)))
	v1581 = v1576 * v1555
	v1583 = v1580 + base.I32_wrap_i64(v1581)
	v1591 = (base.B2i32(base.Ui32(v1583) < base.Ui32(v1580))+base.I32_wrap_i64(int64(base.Ui64(v1581)>>(uint(v1551)%64))))<<(uint(v1570)%32) | int32(base.Ui32(v1583)>>(uint(v1573)%32))
	v1592 = base.I64_extend_i32_u(v1316)
	v1596 = base.I32_wrap_i64(int64(base.Ui64(v1548*v1592) >> (uint(v1551) % 64)))
	v1597 = v1555 * v1592
	v1599 = v1596 + base.I32_wrap_i64(v1597)
	v1607 = (base.B2i32(base.Ui32(v1599) < base.Ui32(v1596))+base.I32_wrap_i64(int64(base.Ui64(v1597)>>(uint(v1551)%64))))<<(uint(v1570)%32) | int32(base.Ui32(v1599)>>(uint(v1573)%32))
	v1609 = v1607 - int32(1)
	if base.Ui32(int32(1048576)) <= base.Ui32(v1539) {
		goto L246
	} else {
		goto L247
	}
L225:
	;
	v1400 = int32(10)
	v1401 = base.I32_div_u_s(v1377-int32(1), v1400)
	v1403 = base.I32_div_u_s(v1361, v1400)
	if base.Ui32(v1401) <= base.Ui32(v1403) {
		goto L228
	} else {
		goto L229
	}
L226:
	;
	v1448 = v1394
	goto L227
L227:
	;
	v1453 = base.I32_rem_u_s(v1313, int32(5))
	if v1453 == int32(0) {
		goto L232
	} else {
		goto L233
	}
L228:
	;
	v1406 = v1326 - int32(1)
	v1411 = *(*int64)(unsafe.Add(mBase, uint32(v1406<<(uint(int32(3))%32))+uint32(_consts[1462])))
	v1415 = int64(32)
	v1417 = base.I32_wrap_i64(int64(base.Ui64(v1411&int64(4294967295)*v1378) >> (uint(v1415) % 64)))
	v1420 = int64(base.Ui64(v1411)>>(uint(v1415)%64)) * v1378
	v1422 = v1417 + base.I32_wrap_i64(v1420)
	v1433 = v1350 + int32(base.Ui32(v1406*int32(1217359))>>(uint(int32(19))%32))
	v1441 = base.I32_rem_u_s((base.B2i32(base.Ui32(v1422) < base.Ui32(v1417))+base.I32_wrap_i64(int64(base.Ui64(v1420)>>(uint(v1415)%64))))<<(uint(int32(6)-v1433)%32)|int32(base.Ui32(v1422)>>(uint(v1433+int32(26))%32)), int32(10))
	v1442 = v1441
	goto L230
L229:
	;
	v1442 = v1394
	goto L230
L230:
	;
	if base.Ui32(int32(33)) < base.Ui32(v1320) {
		v1774 = v1393
		v1776 = v1442
		v1779 = v1326
		v1780 = v1377
		v1782 = v1361
		goto L217
	} else {
		goto L231
	}
L231:
	;
	v1448 = v1442
	goto L227
L232:
	;
	v1458 = v1313
	v1460 = v1394
	goto L235
L233:
	;
	goto L234
L234:
	;
	v1484 = int32(0)
	v1486 = base.I32_rem_u_s(v1316, int32(5))
	if v1486 == v1484 {
		goto L239
	} else {
		goto L240
	}
L235:
	;
	v1476 = v1460 + int32(1)
	v1477 = int32(5)
	v1478 = base.I32_div_u_s(v1458, v1477)
	v1480 = base.I32_rem_u_s(v1478, v1477)
	if v1480 == int32(0) {
		v1458 = v1478
		v1460 = v1476
		goto L235
	} else {
		goto L237
	}
L236:
	;
	if base.Ui32(v1476) < base.Ui32(v1326) {
		v1774 = v1393
		v1776 = v1448
		v1779 = v1326
		v1780 = v1377
		v1782 = v1361
		goto L217
	} else {
		goto L238
	}
L237:
	;
	goto L236
L238:
	;
	v1678 = v1393
	v1680 = v1448
	v1683 = v1326
	v1684 = v1377
	v1686 = v1361
	goto L218
L239:
	;
	v1491 = v1484
	v1495 = v1316
	goto L242
L240:
	;
	v1518 = v1484
	goto L241
L241:
	;
	v1774 = v1393
	v1776 = v1448
	v1779 = v1326
	v1780 = v1377 - base.B2i32(base.Ui32(v1326) <= base.Ui32(v1518))
	v1782 = v1361
	goto L217
L242:
	;
	v1509 = v1491 + int32(1)
	v1510 = int32(5)
	v1511 = base.I32_div_u_s(v1495, v1510)
	v1513 = base.I32_rem_u_s(v1511, v1510)
	if v1513 == int32(0) {
		v1491 = v1509
		v1495 = v1511
		goto L242
	} else {
		goto L244
	}
L243:
	;
	v1518 = v1509
	goto L241
L244:
	;
	goto L243
L245:
	;
	if base.Ui32(int32(32505855)) < base.Ui32(v1539) {
		v1774 = v1591
		v1776 = v1656
		v1779 = v1542
		v1780 = v1607
		v1782 = v1575
		goto L217
	} else {
		goto L253
	}
L246:
	;
	v1612 = int32(10)
	v1613 = base.I32_div_u_s(v1609, v1612)
	v1615 = base.I32_div_u_s(v1575, v1612)
	if base.Ui32(v1613) <= base.Ui32(v1615) {
		goto L249
	} else {
		goto L250
	}
L247:
	;
	v1662 = v1228
	goto L248
L248:
	;
	v1678 = v1591
	v1680 = v1662
	v1683 = v1542
	v1684 = v1609
	v1686 = v1575
	goto L218
L249:
	;
	v1618 = int32(1) - v1542
	v1623 = *(*int64)(unsafe.Add(mBase, uint32(v1618<<(uint(int32(3))%32))+uint32(_consts[1463])))
	v1627 = int64(32)
	v1629 = base.I32_wrap_i64(int64(base.Ui64(v1623&int64(4294967295)*v1576) >> (uint(v1627) % 64)))
	v1632 = int64(base.Ui64(v1623)>>(uint(v1627)%64)) * v1576
	v1634 = v1629 + base.I32_wrap_i64(v1632)
	v1647 = v1541 + (int32(base.Ui32(v1618*int32(1217359))>>(uint(int32(19))%32)) ^ int32(-1))
	v1655 = base.I32_rem_u_s((base.B2i32(base.Ui32(v1634) < base.Ui32(v1629))+base.I32_wrap_i64(int64(base.Ui64(v1632)>>(uint(v1627)%64))))<<(uint(int32(4)-v1647)%32)|int32(base.Ui32(v1634)>>(uint(v1647+int32(28))%32)), int32(10))
	v1656 = v1655
	goto L251
L250:
	;
	v1656 = v1228
	goto L251
L251:
	;
	if base.Ui32(int32(2097151)) < base.Ui32(v1539) {
		goto L245
	} else {
		goto L252
	}
L252:
	;
	v1662 = v1656
	goto L248
L253:
	;
	v1668 = int32(-1)
	if v1313&(v1668<<(uint(v1541-int32(1))%32)^v1668) != 0 {
		v1774 = v1591
		v1776 = v1656
		v1779 = v1542
		v1780 = v1607
		v1782 = v1575
		goto L217
	} else {
		goto L254
	}
L254:
	;
	v1678 = v1591
	v1680 = v1656
	v1683 = v1542
	v1684 = v1607
	v1686 = v1575
	goto L218
L255:
	;
	v1761 = v1746 & int32(255)
	v1854 = v1742
	v1856 = v1744
	v1860 = v1683
	v1871 = (v1759|base.B2i32(v1761 != int32(5))|v1744)&base.B2i32(base.Ui32(int32(4)) < base.Ui32(v1761)) | base.B2i32(v1744 == v1751)
	goto L216
L256:
	;
	v1742 = v1694
	v1744 = v1678
	v1746 = v1680
	v1751 = v1686
	v1759 = int32(0)
	goto L255
L257:
	;
	goto L258
L258:
	;
	v1704 = v1694
	v1705 = v1678
	v1707 = v1680
	v1709 = v1696
	v1711 = v1698
	v1712 = int32(1)
	goto L259
L259:
	;
	v1722 = v1704 + int32(1)
	v1723 = int32(10)
	v1724 = base.I32_div_u_s(v1705, v1723)
	v1727 = v1705 - v1724*v1723
	v1732 = v1712 & base.B2i32(v1707&int32(255) == int32(0))
	v1734 = base.I32_div_u_s(v1709, v1723)
	v1736 = base.I32_div_u_s(v1711, v1723)
	if base.Ui32(v1736) < base.Ui32(v1734) {
		v1704 = v1722
		v1705 = v1724
		v1707 = v1727
		v1709 = v1734
		v1711 = v1736
		v1712 = v1732
		goto L259
	} else {
		goto L261
	}
L260:
	;
	v1742 = v1722
	v1744 = v1724
	v1746 = v1727
	v1751 = v1711
	v1759 = v1732 ^ int32(1)
	goto L255
L261:
	;
	goto L260
L262:
	;
	v1798 = v1790
	v1799 = v1774
	v1800 = v1792
	v1802 = v1794
	goto L265
L263:
	;
	v1829 = v1790
	v1830 = v1774
	v1832 = v1776
	v1838 = v1782
	goto L264
L264:
	;
	v1854 = v1829
	v1856 = v1830
	v1860 = v1779
	v1871 = base.B2i32(v1830 == v1838) | base.B2i32(base.Ui32(int32(4)) < base.Ui32(v1832&int32(255)))
	goto L216
L265:
	;
	v1816 = v1798 + int32(1)
	v1817 = int32(10)
	v1818 = base.I32_div_u_s(v1799, v1817)
	v1820 = base.I32_div_u_s(v1800, v1817)
	v1822 = base.I32_div_u_s(v1802, v1817)
	if base.Ui32(v1822) < base.Ui32(v1820) {
		v1798 = v1816
		v1799 = v1818
		v1800 = v1820
		v1802 = v1822
		goto L265
	} else {
		goto L267
	}
L266:
	;
	v1829 = v1816
	v1830 = v1818
	v1832 = v1799 - v1818*int32(10)
	v1838 = v1802
	goto L264
L267:
	;
	goto L266
L268:
	;
	v1921 = v1873
	v1931 = v1872
	v1938 = int32(9)
	goto L208
L269:
	;
	if base.Ui32(int32(999999)) < base.Ui32(v1879) {
		v1921 = v1879
		v1931 = v1889
		v1938 = int32(7)
		goto L208
	} else {
		goto L270
	}
L270:
	;
	if base.Ui32(int32(99999)) < base.Ui32(v1879) {
		v1921 = v1879
		v1931 = v1889
		v1938 = int32(6)
		goto L208
	} else {
		goto L271
	}
L271:
	;
	if base.Ui32(int32(9999)) < base.Ui32(v1879) {
		v1921 = v1879
		v1931 = v1889
		v1938 = int32(5)
		goto L208
	} else {
		goto L272
	}
L272:
	;
	if base.Ui32(int32(999)) < base.Ui32(v1879) {
		v1921 = v1879
		v1931 = v1889
		v1938 = int32(4)
		goto L208
	} else {
		goto L273
	}
L273:
	;
	if base.Ui32(int32(99)) < base.Ui32(v1879) {
		v1921 = v1879
		v1931 = v1889
		v1938 = int32(3)
		goto L208
	} else {
		goto L274
	}
L274:
	;
	if base.Ui32(int32(9)) < base.Ui32(v1879) {
		goto L275
	} else {
		goto L276
	}
L275:
	;
	v1918 = int32(2)
	goto L277
L276:
	;
	v1918 = int32(1)
	goto L277
L277:
	;
	v1921 = v1879
	v1931 = v1889
	v1938 = v1918
	goto L208
L278:
	;
	v1943 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(v1227))) = uint8(v1943)
	v1946 = int32(1)
	goto L280
L279:
	;
	v1946 = v1940
	goto L280
L280:
	;
	if base.Ui32(v1939+int32(3)) <= base.Ui32(int32(9)) {
		goto L283
	} else {
		goto L284
	}
L281:
	;
	v2167 = int32(0)
	if base.Ui32(v1921) < base.Ui32(int32(10000)) {
		goto L323
	} else {
		goto L324
	}
L282:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1951))) = v2163
	v2165 = v2162
	goto L281
L283:
	;
	v1951 = v1227 + v1946
	v1952 = int32(0)
	if v1939 <= v1952 {
		goto L286
	} else {
		goto L287
	}
L284:
	;
	goto L285
L285:
	;
	if v1931 != 0 {
		goto L291
	} else {
		goto L292
	}
L286:
	;
	v2162 = int32(2) - v1939
	v2163 = int64(3472328296227679792)
	goto L282
L287:
	;
	goto L288
L288:
	;
	if int32(0) <= v1931 {
		v2162 = v1952
		v2163 = int64(3472328296227680304)
		goto L282
	} else {
		goto L289
	}
L289:
	;
	v2165 = int32(1)
	goto L281
L290:
	;
	v2011 = int32(0)
	if base.Ui32(v1995) < base.Ui32(int32(10000)) {
		goto L299
	} else {
		goto L300
	}
L291:
	;
	v1995 = v1921
	v2001 = v1938
	goto L290
L292:
	;
	goto L293
L293:
	;
	v1964 = v1921
	v1969 = v1938
	goto L294
L294:
	;
	if v1964&int32(1) != 0 {
		v1995 = v1964
		v2001 = v1969
		goto L290
	} else {
		goto L296
	}
L295:
	;
	v1995 = v1964
	v2001 = v1969
	goto L290
L296:
	;
	v1988 = base.I32_div_u_s(v1964, int32(10))
	if int32(0)-v1964 == v1988*int32(-10) {
		v1964 = v1988
		v1969 = v1969 - int32(1)
		goto L294
	} else {
		goto L297
	}
L297:
	;
	goto L295
L298:
	;
	if base.Ui32(v2070) < base.Ui32(int32(100)) {
		goto L306
	} else {
		goto L307
	}
L299:
	;
	v2068 = v2011
	v2070 = v1995
	goto L298
L300:
	;
	goto L301
L301:
	;
	v2018 = v2011
	v2019 = v1995
	goto L302
L302:
	;
	v2035 = v1227 + v1946 + v2001 - v2018
	v2039 = base.I32_div_u_s(v2019, int32(10000))
	v2042 = v2039*int32(-10000) + v2019
	v2043 = int32(100)
	v2044 = base.I32_div_u_s(v2042, v2043)
	v2045 = int32(1)
	v2049 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2044<<(uint(v2045)%32))+uint32(_consts[1464]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v2035-int32(3)))) = uint16(v2049)
	v2060 = int32(*(*uint16)(unsafe.Add(mBase, uint32((v2042-v2044*v2043)<<(uint(v2045)%32))+uint32(_consts[1464]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v2035-v2045))) = uint16(v2060)
	v2063 = v2018 + int32(4)
	if base.Ui32(int32(99999999)) < base.Ui32(v2019) {
		v2018 = v2063
		v2019 = v2039
		goto L302
	} else {
		goto L304
	}
L303:
	;
	v2068 = v2063
	v2070 = v2039
	goto L298
L304:
	;
	goto L303
L305:
	;
	if base.Ui32(int32(10)) <= base.Ui32(v2110) {
		goto L310
	} else {
		goto L311
	}
L306:
	;
	v2109 = v2068
	v2110 = v2070
	goto L305
L307:
	;
	goto L308
L308:
	;
	v2092 = int32(65535)
	v2094 = int32(100)
	v2095 = base.I32_div_u_s(v2070&v2092, v2094)
	v2105 = int32(*(*uint16)(unsafe.Add(mBase, uint32((v2070-v2095*v2094)&v2092<<(uint(int32(1))%32))+uint32(_consts[1464]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v1227+v1946+v2001+(v2068^int32(-1))))) = uint16(v2105)
	v2109 = v2068 | int32(2)
	v2110 = v2095
	goto L305
L309:
	;
	v2129 = int32(1)
	v2130 = v1939 - v2129
	v2131 = v1227 + v1946
	*(*uint8)(unsafe.Add(mBase, uint32(v2131))) = uint8(v2128)
	if base.Ui32(int32(2)) <= base.Ui32(v2001) {
		goto L313
	} else {
		goto L314
	}
L310:
	;
	v2117 = v2110 << (uint(int32(1)) % 32)
	v2120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2117)+uint32(_consts[1465]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1227+(v1946+v2001-v2109)))) = uint8(v2120)
	v2124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2117)+uint32(_consts[1464]))))
	v2128 = v2124
	goto L309
L311:
	;
	goto L312
L312:
	;
	v2128 = v2110 | int32(48)
	goto L309
L313:
	;
	v2136 = int32(46)
	*(*uint8)(unsafe.Add(mBase, uint32(v2131)+1)) = uint8(v2136)
	v2140 = v2001 + int32(1)
	goto L315
L314:
	;
	v2140 = v2129
	goto L315
L315:
	;
	v2141 = v2140 + v1946
	v2142 = v1227 + v2141
	v2143 = int32(101)
	*(*uint8)(unsafe.Add(mBase, uint32(v2142))) = uint8(v2143)
	v2148 = base.B2i32(v2130 < int32(0))
	if v2130 < int32(0) {
		goto L316
	} else {
		goto L317
	}
L316:
	;
	v2149 = int32(45)
	goto L318
L317:
	;
	v2149 = int32(43)
	goto L318
L318:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v2142)+1)) = uint8(v2149)
	if v2130 < int32(0) {
		goto L319
	} else {
		goto L320
	}
L319:
	;
	v2153 = int32(1) - v1939
	goto L321
L320:
	;
	v2153 = v2130
	goto L321
L321:
	;
	v2158 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2153<<(uint(int32(1))%32))+uint32(_consts[1464]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v2142)+2)) = uint16(v2158)
	v2347 = v2141 + int32(4)
	goto L186
L322:
	;
	if base.Ui32(v2226) < base.Ui32(int32(100)) {
		goto L330
	} else {
		goto L331
	}
L323:
	;
	v2225 = v2167
	v2226 = v1921
	goto L322
L324:
	;
	goto L325
L325:
	;
	v2174 = v1921
	v2175 = v2167
	goto L326
L326:
	;
	v2191 = v1951 + v2165 + v1938 - v2175
	v2192 = int32(4)
	v2195 = base.I32_div_u_s(v2174, int32(10000))
	v2198 = v2195*int32(-10000) + v2174
	v2199 = int32(100)
	v2200 = base.I32_div_u_s(v2198, v2199)
	v2201 = int32(1)
	v2205 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2200<<(uint(v2201)%32))+uint32(_consts[1464]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v2191-v2192))) = uint16(v2205)
	v2216 = int32(*(*uint16)(unsafe.Add(mBase, uint32((v2198-v2200*v2199)<<(uint(v2201)%32))+uint32(_consts[1464]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v2191-int32(2)))) = uint16(v2216)
	v2219 = v2175 + v2192
	if base.Ui32(int32(99999999)) < base.Ui32(v2174) {
		v2174 = v2195
		v2175 = v2219
		goto L326
	} else {
		goto L328
	}
L327:
	;
	v2225 = v2219
	v2226 = v2195
	goto L322
L328:
	;
	goto L327
L329:
	;
	if base.Ui32(int32(10)) <= base.Ui32(v2265) {
		goto L334
	} else {
		goto L335
	}
L330:
	;
	v2265 = v2226
	v2266 = v2225
	goto L329
L331:
	;
	goto L332
L332:
	;
	v2246 = int32(2)
	v2248 = int32(65535)
	v2250 = int32(100)
	v2251 = base.I32_div_u_s(v2226&v2248, v2250)
	v2261 = int32(*(*uint16)(unsafe.Add(mBase, uint32((v2226-v2251*v2250)&v2248<<(uint(int32(1))%32))+uint32(_consts[1464]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v1951+v2165+v1938-v2225-v2246))) = uint16(v2261)
	v2265 = v2251
	v2266 = v2225 | v2246
	goto L329
L333:
	;
	v2284 = int32(1)
	if v2165 == v2284 {
		goto L338
	} else {
		goto L339
	}
L334:
	;
	v2278 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2265<<(uint(int32(1))%32))+uint32(_consts[1464]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v1951+v2165+v1938-v2266-int32(2)))) = uint16(v2278)
	goto L333
L335:
	;
	goto L336
L336:
	;
	v2282 = v2265 | int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v1951+v2165))) = uint8(v2282)
	goto L333
L337:
	;
	v2347 = v2324 + int32(base.Ui32(v1245)>>(uint(int32(31))%32))
	goto L186
L338:
	;
	if v1939&int32(4) != 0 {
		goto L341
	} else {
		goto L342
	}
L339:
	;
	goto L340
L340:
	;
	if v1931 < int32(0) {
		goto L350
	} else {
		goto L351
	}
L341:
	;
	v2289 = *(*int32)(unsafe.Add(mBase, uint32(v1951)+1))
	*(*int32)(unsafe.Add(mBase, uint32(v1951))) = v2289
	v2292 = int32(5)
	goto L343
L342:
	;
	v2292 = v2284
	goto L343
L343:
	;
	if v1939&int32(2) != 0 {
		goto L344
	} else {
		goto L345
	}
L344:
	;
	v2295 = v1951 + v2292
	v2298 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2295))))
	*(*uint16)(unsafe.Add(mBase, uint32(v2295-int32(1)))) = uint16(v2298)
	v2303 = v2292 | int32(2)
	goto L346
L345:
	;
	v2303 = v2292
	goto L346
L346:
	;
	if v1939&int32(1) != 0 {
		goto L347
	} else {
		goto L348
	}
L347:
	;
	v2306 = v1951 + v2303
	v2309 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2306))))
	*(*uint8)(unsafe.Add(mBase, uint32(v2306-int32(1)))) = uint8(v2309)
	goto L349
L348:
	;
	goto L349
L349:
	;
	v2313 = int32(46)
	*(*uint8)(unsafe.Add(mBase, uint32(v1951+v1939))) = uint8(v2313)
	v2324 = v1938 + int32(1)
	goto L337
L350:
	;
	v2321 = int32(2) - v1931
	goto L352
L351:
	;
	v2321 = v1939
	goto L352
L352:
	;
	v2324 = v2321
	goto L337
L353:
	;
	goto L181
L354:
	;
	v2384 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v2379+v2365))) = uint8(v2384)
	v2386 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v2386 != v11 {
		goto L358
	} else {
		goto L359
	}
L355:
	;
	v2375 = v2363
	v2376 = int32(0)
	goto L357
L356:
	;
	v2370 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(v2365))) = uint8(v2370)
	v2375 = int32(0) - v2363
	v2376 = int32(1)
	goto L357
L357:
	;
	v2378 = F_pg_ultoa_n(m, v2375, v2365+v2376)
	mBase = m.M
	v2379 = v2378 + v2376
	v2381 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v2365+v2379))) = uint8(v2381)
	goto L354
L358:
	;
	F_pfree(m, v11)
	mBase = m.M
	v2389 = m.ExcPending
	if v2389 != 0 {
		goto L1
	} else {
		goto L361
	}
L359:
	;
	goto L360
L360:
	;
	return v21
L361:
	;
	goto L360
}
