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
	return base.B2i32(v155 == int32(0))
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
	var v53 int32
	_ = v53
	var v62 float32
	_ = v62
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
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
	v53 = v2
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
	v66 = v53 << (uint(int32(2)) % 32)
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v66+v43)))
	v70 = v49
	v72 = v49
	v85 = v62
	goto L15
L15:
	;
	if v70 != v64 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v109 = v53 + int32(1)
	if v109 != v32 {
		v49 = v104
		v53 = v109
		v62 = v107
		goto L10
	} else {
		goto L27
	}
L17:
	;
	v88 = v70 << (uint(int32(2)) % 32)
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
	v99 = v70 + int32(1)
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
		v70 = v99
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
	F_errmsg(m, int32(_a_F_sparsevec_inner_product_0), v19)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	F_errfinish(m, int32(_a_F_sparsevec_inner_product_1), int32(50), int32(_a_F_sparsevec_inner_product_2))
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
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v70 int32
	_ = v70
	var v72 float64
	_ = v72
	var v76 int32
	_ = v76
	var v77 float32
	_ = v77
	var v78 float64
	_ = v78
	var v80 float32
	_ = v80
	var v81 float64
	_ = v81
	var v83 float32
	_ = v83
	var v84 float64
	_ = v84
	var v86 float32
	_ = v86
	var v87 float64
	_ = v87
	var v92 float64
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v114 float64
	_ = v114
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v130 float64
	_ = v130
	var v135 float32
	_ = v135
	var v136 float64
	_ = v136
	var v138 float64
	_ = v138
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v158 float64
	_ = v158
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v195 float32
	_ = v195
	var v198 float32
	_ = v198
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
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
	var v228 int32
	_ = v228
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v261 int32
	_ = v261
	var v263 float32
	_ = v263
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v272 int32
	_ = v272
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v301 int32
	_ = v301
	var v306 int32
	_ = v306
	var v310 int32
	_ = v310
	var v315 int32
	_ = v315
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
	v53 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v42) {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	if base.F64_gt(v158, float64(0)) == int32(0) {
		goto L22
	} else {
		goto L23
	}
L12:
	;
	v58 = v53
	v70 = v2
	v72 = v14
	goto L15
L13:
	;
	v100 = v53
	v114 = v14
	goto L14
L14:
	;
	v116 = v100
	v118 = v2
	v130 = v114
	goto L19
L15:
	;
	v76 = v50 + v58<<(uint(int32(2))%32)
	v77 = *(*float32)(unsafe.Add(mBase, uint32(v76)+12))
	v78 = base.F64_promote_f32(v77)
	v80 = *(*float32)(unsafe.Add(mBase, uint32(v76)+8))
	v81 = base.F64_promote_f32(v80)
	v83 = *(*float32)(unsafe.Add(mBase, uint32(v76)+4))
	v84 = base.F64_promote_f32(v83)
	v86 = *(*float32)(unsafe.Add(mBase, uint32(v76)))
	v87 = base.F64_promote_f32(v86)
	v92 = base.F64_add(base.F64_mul(v78, v78), base.F64_add(base.F64_mul(v81, v81), base.F64_add(base.F64_mul(v84, v84), base.F64_add(base.F64_mul(v87, v87), v72))))
	v93 = int32(4)
	v94 = v58 + v93
	v96 = v70 + v93
	if v96 != v42&int32(2147483644) {
		v58 = v94
		v70 = v96
		v72 = v92
		goto L15
	} else {
		goto L17
	}
L16:
	;
	if v52 == int32(0) {
		v158 = v92
		goto L11
	} else {
		goto L18
	}
L17:
	;
	goto L16
L18:
	;
	v100 = v94
	v114 = v92
	goto L14
L19:
	;
	v135 = *(*float32)(unsafe.Add(mBase, uint32(v50+v116<<(uint(int32(2))%32))))
	v136 = base.F64_promote_f32(v135)
	v138 = base.F64_add(base.F64_mul(v136, v136), v130)
	v139 = int32(1)
	v142 = v118 + v139
	if v142 != v52 {
		v116 = v116 + v139
		v118 = v142
		v130 = v138
		goto L19
	} else {
		goto L21
	}
L20:
	;
	v158 = v138
	goto L11
L21:
	;
	goto L20
L22:
	;
	return v35
L23:
	;
	goto L24
L24:
	;
	v166 = v35 + int32(16)
	v167 = v166 + v49
	v169 = int32(0)
	v172 = v169
	v173 = v169
	goto L26
L25:
	;
	if v205 == int32(0) {
		goto L32
	} else {
		goto L33
	}
L26:
	;
	v188 = v172 << (uint(int32(2)) % 32)
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v188+v47)))
	*(*int32)(unsafe.Add(mBase, uint32(v166+v188))) = v191
	v195 = *(*float32)(unsafe.Add(mBase, uint32(v188+v50)))
	v198 = base.F32_demote_f64(base.F64_div(base.F64_promote_f32(v195), base.F64_sqrt(v158)))
	*(*float32)(unsafe.Add(mBase, uint32(v188+v167))) = v198
	if base.F32_eq(base.F32_abs(v198), math.Float32frombits(uint32(0x7f800000))) != 0 {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	F_float_overflow_error(m)
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L1
	} else {
		goto L31
	}
L28:
	;
	goto L27
L29:
	;
	v205 = v173 + base.F32_eq(v198, float32(0))
	v207 = v172 + int32(1)
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
	if v207 < v208 {
		v172 = v207
		v173 = v205
		goto L26
	} else {
		goto L30
	}
L30:
	;
	goto L25
L31:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L32:
	;
	return v35
L33:
	;
	goto L34
L34:
	;
	v217 = v25 - v205
	v218 = F_mul_size(m, int32(4), v217)
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	v220 = F_add_size(m, int32(16), v218)
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	v223 = F_mul_size(m, int32(4), v217)
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	v225 = F_add_size(m, v220, v223)
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	v227 = F_palloc0(m, v225)
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v227)+8)) = v217
	*(*int32)(unsafe.Add(mBase, uint32(v227)+4)) = v22
	*(*int32)(unsafe.Add(mBase, uint32(v227))) = v225 << (uint(int32(2)) % 32)
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v35)+8))
	if int32(0) < v234 {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L1
	} else {
		goto L52
	}
L41:
	;
	v238 = v227 + int32(16)
	v242 = int32(0)
	v244 = v242
	v245 = v242
	v247 = v234
	goto L44
L42:
	;
	goto L43
L43:
	;
	F_pfree(m, v35)
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L1
	} else {
		goto L51
	}
L44:
	;
	v261 = v244 << (uint(int32(2)) % 32)
	v263 = *(*float32)(unsafe.Add(mBase, uint32(v167+v261)))
	if base.F32_ne(v263, float32(0)) != 0 {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	goto L43
L46:
	;
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v227)+8))
	if v266 <= v245 {
		goto L40
	} else {
		goto L49
	}
L47:
	;
	v279 = v245
	v280 = v247
	goto L48
L48:
	;
	v282 = v244 + int32(1)
	if v282 < v280 {
		v244 = v282
		v245 = v279
		v247 = v280
		goto L44
	} else {
		goto L50
	}
L49:
	;
	v269 = v245 << (uint(int32(2)) % 32)
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v166+v261)))
	*(*int32)(unsafe.Add(mBase, uint32(v238+v269))) = v272
	*(*float32)(unsafe.Add(mBase, uint32(v269+(v238+v217<<(uint(int32(2))%32))))) = v263
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v35)+8))
	v279 = v245 + int32(1)
	v280 = v276
	goto L48
L50:
	;
	goto L45
L51:
	;
	return v227
L52:
	;
	F_errmsg_internal(m, int32(_a_F_sparsevec_l2_normalize_0), int32(0))
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	F_errfinish(m, int32(_a_F_sparsevec_l2_normalize_1), int32(1129), int32(_a_F_sparsevec_l2_normalize_2))
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
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
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v160 int32
	_ = v160
	var v163 int64
	_ = v163
	var v165 int64
	_ = v165
	var v166 int64
	_ = v166
	var v168 int64
	_ = v168
	var v170 int32
	_ = v170
	var v172 int64
	_ = v172
	var v173 int64
	_ = v173
	var v175 int32
	_ = v175
	var v182 int32
	_ = v182
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v194 int64
	_ = v194
	var v198 int32
	_ = v198
	var v199 int64
	_ = v199
	var v201 int32
	_ = v201
	var v209 int32
	_ = v209
	var v210 int64
	_ = v210
	var v214 int32
	_ = v214
	var v215 int64
	_ = v215
	var v217 int32
	_ = v217
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v239 int64
	_ = v239
	var v243 int64
	_ = v243
	var v245 int32
	_ = v245
	var v248 int64
	_ = v248
	var v250 int32
	_ = v250
	var v261 int32
	_ = v261
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v277 int32
	_ = v277
	var v283 int32
	_ = v283
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v314 int32
	_ = v314
	var v316 int32
	_ = v316
	var v322 int32
	_ = v322
	var v325 int32
	_ = v325
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v343 int32
	_ = v343
	var v349 int32
	_ = v349
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v376 int64
	_ = v376
	var v378 int64
	_ = v378
	var v379 int64
	_ = v379
	var v381 int64
	_ = v381
	var v383 int32
	_ = v383
	var v385 int64
	_ = v385
	var v386 int64
	_ = v386
	var v388 int32
	_ = v388
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v403 int32
	_ = v403
	var v405 int32
	_ = v405
	var v406 int64
	_ = v406
	var v410 int32
	_ = v410
	var v411 int64
	_ = v411
	var v413 int32
	_ = v413
	var v421 int32
	_ = v421
	var v422 int64
	_ = v422
	var v426 int32
	_ = v426
	var v427 int64
	_ = v427
	var v429 int32
	_ = v429
	var v437 int32
	_ = v437
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v443 int32
	_ = v443
	var v446 int32
	_ = v446
	var v449 int64
	_ = v449
	var v453 int64
	_ = v453
	var v455 int32
	_ = v455
	var v458 int64
	_ = v458
	var v460 int32
	_ = v460
	var v473 int32
	_ = v473
	var v481 int32
	_ = v481
	var v483 int32
	_ = v483
	var v490 int32
	_ = v490
	var v494 int32
	_ = v494
	var v506 int32
	_ = v506
	var v509 int32
	_ = v509
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v527 int32
	_ = v527
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v539 int32
	_ = v539
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v556 int32
	_ = v556
	var v561 int32
	_ = v561
	var v563 int32
	_ = v563
	var v565 int32
	_ = v565
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v588 int32
	_ = v588
	var v590 int32
	_ = v590
	var v602 int32
	_ = v602
	var v605 int32
	_ = v605
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v623 int32
	_ = v623
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v629 int32
	_ = v629
	var v631 int32
	_ = v631
	var v645 int32
	_ = v645
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v649 int32
	_ = v649
	var v651 int32
	_ = v651
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v664 int32
	_ = v664
	var v665 int32
	_ = v665
	var v683 int32
	_ = v683
	var v684 int32
	_ = v684
	var v688 int32
	_ = v688
	var v700 int32
	_ = v700
	var v701 int32
	_ = v701
	var v702 int32
	_ = v702
	var v709 int32
	_ = v709
	var v714 int32
	_ = v714
	var v747 int32
	_ = v747
	var v751 int32
	_ = v751
	var v756 int32
	_ = v756
	var v767 int32
	_ = v767
	var v768 int32
	_ = v768
	var v771 int32
	_ = v771
	var v774 int32
	_ = v774
	var v775 int32
	_ = v775
	var v780 int32
	_ = v780
	var v781 int32
	_ = v781
	var v794 int32
	_ = v794
	var v796 int32
	_ = v796
	var v817 int32
	_ = v817
	var v823 int32
	_ = v823
	var v828 int32
	_ = v828
	var v840 int32
	_ = v840
	var v847 int32
	_ = v847
	var v848 int32
	_ = v848
	var v864 int32
	_ = v864
	var v868 int32
	_ = v868
	var v871 int32
	_ = v871
	var v872 int32
	_ = v872
	var v873 int32
	_ = v873
	var v874 int32
	_ = v874
	var v876 int32
	_ = v876
	var v885 int32
	_ = v885
	var v888 int32
	_ = v888
	var v893 int32
	_ = v893
	var v894 int32
	_ = v894
	var v917 int32
	_ = v917
	var v919 int32
	_ = v919
	var v920 int32
	_ = v920
	var v928 int32
	_ = v928
	var v932 int32
	_ = v932
	var v933 int32
	_ = v933
	var v935 int32
	_ = v935
	var v936 int32
	_ = v936
	var v943 int32
	_ = v943
	var v944 int32
	_ = v944
	var v946 int32
	_ = v946
	var v950 int32
	_ = v950
	var v954 int32
	_ = v954
	var v959 int32
	_ = v959
	var v960 int32
	_ = v960
	var v961 int32
	_ = v961
	var v962 int32
	_ = v962
	var v967 int32
	_ = v967
	var v968 int32
	_ = v968
	var v972 int32
	_ = v972
	var v977 int32
	_ = v977
	var v981 int32
	_ = v981
	var v982 int64
	_ = v982
	var v984 int32
	_ = v984
	var v986 int32
	_ = v986
	var v993 int32
	_ = v993
	var v994 int32
	_ = v994
	var v1010 int32
	_ = v1010
	var v1011 int32
	_ = v1011
	var v1014 int32
	_ = v1014
	var v1017 int32
	_ = v1017
	var v1018 int32
	_ = v1018
	var v1019 int32
	_ = v1019
	var v1020 int32
	_ = v1020
	var v1022 int32
	_ = v1022
	var v1031 int32
	_ = v1031
	var v1034 int32
	_ = v1034
	var v1039 int32
	_ = v1039
	var v1040 int32
	_ = v1040
	var v1061 int32
	_ = v1061
	var v1063 int32
	_ = v1063
	var v1065 int32
	_ = v1065
	var v1066 int32
	_ = v1066
	var v1074 int32
	_ = v1074
	var v1078 int32
	_ = v1078
	var v1079 int32
	_ = v1079
	var v1089 int32
	_ = v1089
	var v1093 int32
	_ = v1093
	var v1095 int32
	_ = v1095
	var v1100 int32
	_ = v1100
	var v1103 int32
	_ = v1103
	var v1106 int32
	_ = v1106
	var v1109 int32
	_ = v1109
	var v1114 int32
	_ = v1114
	var v1117 int32
	_ = v1117
	var v1120 int32
	_ = v1120
	var v1124 int32
	_ = v1124
	var v1132 int32
	_ = v1132
	var v1135 int32
	_ = v1135
	var v1158 int32
	_ = v1158
	var v1159 int32
	_ = v1159
	var v1160 int32
	_ = v1160
	var v1164 int32
	_ = v1164
	var v1167 int32
	_ = v1167
	var v1171 int32
	_ = v1171
	var v1174 int32
	_ = v1174
	var v1176 int32
	_ = v1176
	var v1177 int32
	_ = v1177
	var v1178 int32
	_ = v1178
	var v1180 int32
	_ = v1180
	var v1185 int32
	_ = v1185
	var v1190 int32
	_ = v1190
	var v1191 int32
	_ = v1191
	var v1193 int32
	_ = v1193
	var v1194 int32
	_ = v1194
	var v1196 int32
	_ = v1196
	var v1198 int32
	_ = v1198
	var v1199 int32
	_ = v1199
	var v1202 float32
	_ = v1202
	var v1204 int32
	_ = v1204
	var v1205 int32
	_ = v1205
	var v1222 int32
	_ = v1222
	var v1224 int32
	_ = v1224
	var v1227 int32
	_ = v1227
	var v1228 int32
	_ = v1228
	var v1233 int32
	_ = v1233
	var v1237 int32
	_ = v1237
	var v1240 int32
	_ = v1240
	var v1245 int32
	_ = v1245
	var v1249 int32
	_ = v1249
	var v1256 int32
	_ = v1256
	var v1257 int32
	_ = v1257
	var v1263 int32
	_ = v1263
	var v1268 int32
	_ = v1268
	var v1270 int32
	_ = v1270
	var v1280 int32
	_ = v1280
	var v1283 int32
	_ = v1283
	var v1286 int32
	_ = v1286
	var v1291 int32
	_ = v1291
	var v1293 int32
	_ = v1293
	var v1297 int32
	_ = v1297
	var v1303 int32
	_ = v1303
	var v1306 int64
	_ = v1306
	var v1308 int64
	_ = v1308
	var v1309 int64
	_ = v1309
	var v1311 int64
	_ = v1311
	var v1313 int32
	_ = v1313
	var v1315 int64
	_ = v1315
	var v1316 int64
	_ = v1316
	var v1318 int32
	_ = v1318
	var v1325 int32
	_ = v1325
	var v1330 int32
	_ = v1330
	var v1331 int32
	_ = v1331
	var v1334 int32
	_ = v1334
	var v1336 int32
	_ = v1336
	var v1337 int64
	_ = v1337
	var v1341 int32
	_ = v1341
	var v1342 int64
	_ = v1342
	var v1344 int32
	_ = v1344
	var v1352 int32
	_ = v1352
	var v1353 int64
	_ = v1353
	var v1357 int32
	_ = v1357
	var v1358 int64
	_ = v1358
	var v1360 int32
	_ = v1360
	var v1368 int32
	_ = v1368
	var v1369 int32
	_ = v1369
	var v1373 int32
	_ = v1373
	var v1374 int32
	_ = v1374
	var v1376 int32
	_ = v1376
	var v1379 int32
	_ = v1379
	var v1382 int64
	_ = v1382
	var v1386 int64
	_ = v1386
	var v1388 int32
	_ = v1388
	var v1391 int64
	_ = v1391
	var v1393 int32
	_ = v1393
	var v1404 int32
	_ = v1404
	var v1412 int32
	_ = v1412
	var v1413 int32
	_ = v1413
	var v1420 int32
	_ = v1420
	var v1426 int32
	_ = v1426
	var v1432 int32
	_ = v1432
	var v1433 int32
	_ = v1433
	var v1449 int32
	_ = v1449
	var v1450 int32
	_ = v1450
	var v1451 int32
	_ = v1451
	var v1453 int32
	_ = v1453
	var v1457 int32
	_ = v1457
	var v1459 int32
	_ = v1459
	var v1465 int32
	_ = v1465
	var v1468 int32
	_ = v1468
	var v1482 int32
	_ = v1482
	var v1483 int32
	_ = v1483
	var v1484 int32
	_ = v1484
	var v1486 int32
	_ = v1486
	var v1492 int32
	_ = v1492
	var v1512 int32
	_ = v1512
	var v1514 int32
	_ = v1514
	var v1515 int32
	_ = v1515
	var v1519 int64
	_ = v1519
	var v1521 int64
	_ = v1521
	var v1522 int64
	_ = v1522
	var v1524 int64
	_ = v1524
	var v1526 int32
	_ = v1526
	var v1528 int64
	_ = v1528
	var v1529 int64
	_ = v1529
	var v1531 int32
	_ = v1531
	var v1542 int32
	_ = v1542
	var v1543 int32
	_ = v1543
	var v1546 int32
	_ = v1546
	var v1548 int32
	_ = v1548
	var v1549 int64
	_ = v1549
	var v1553 int32
	_ = v1553
	var v1554 int64
	_ = v1554
	var v1556 int32
	_ = v1556
	var v1564 int32
	_ = v1564
	var v1565 int64
	_ = v1565
	var v1569 int32
	_ = v1569
	var v1570 int64
	_ = v1570
	var v1572 int32
	_ = v1572
	var v1580 int32
	_ = v1580
	var v1582 int32
	_ = v1582
	var v1583 int32
	_ = v1583
	var v1584 int32
	_ = v1584
	var v1586 int32
	_ = v1586
	var v1589 int32
	_ = v1589
	var v1592 int64
	_ = v1592
	var v1596 int64
	_ = v1596
	var v1598 int32
	_ = v1598
	var v1601 int64
	_ = v1601
	var v1603 int32
	_ = v1603
	var v1616 int32
	_ = v1616
	var v1624 int32
	_ = v1624
	var v1626 int32
	_ = v1626
	var v1633 int32
	_ = v1633
	var v1637 int32
	_ = v1637
	var v1649 int32
	_ = v1649
	var v1652 int32
	_ = v1652
	var v1654 int32
	_ = v1654
	var v1655 int32
	_ = v1655
	var v1656 int32
	_ = v1656
	var v1666 int32
	_ = v1666
	var v1667 int32
	_ = v1667
	var v1668 int32
	_ = v1668
	var v1670 int32
	_ = v1670
	var v1676 int32
	_ = v1676
	var v1677 int32
	_ = v1677
	var v1678 int32
	_ = v1678
	var v1679 int32
	_ = v1679
	var v1680 int32
	_ = v1680
	var v1682 int32
	_ = v1682
	var v1694 int32
	_ = v1694
	var v1695 int32
	_ = v1695
	var v1696 int32
	_ = v1696
	var v1699 int32
	_ = v1699
	var v1704 int32
	_ = v1704
	var v1706 int32
	_ = v1706
	var v1708 int32
	_ = v1708
	var v1714 int32
	_ = v1714
	var v1715 int32
	_ = v1715
	var v1720 int32
	_ = v1720
	var v1721 int32
	_ = v1721
	var v1731 int32
	_ = v1731
	var v1733 int32
	_ = v1733
	var v1745 int32
	_ = v1745
	var v1748 int32
	_ = v1748
	var v1750 int32
	_ = v1750
	var v1751 int32
	_ = v1751
	var v1752 int32
	_ = v1752
	var v1762 int32
	_ = v1762
	var v1763 int32
	_ = v1763
	var v1764 int32
	_ = v1764
	var v1766 int32
	_ = v1766
	var v1770 int32
	_ = v1770
	var v1771 int32
	_ = v1771
	var v1772 int32
	_ = v1772
	var v1774 int32
	_ = v1774
	var v1788 int32
	_ = v1788
	var v1789 int32
	_ = v1789
	var v1790 int32
	_ = v1790
	var v1792 int32
	_ = v1792
	var v1794 int32
	_ = v1794
	var v1801 int32
	_ = v1801
	var v1802 int32
	_ = v1802
	var v1807 int32
	_ = v1807
	var v1808 int32
	_ = v1808
	var v1826 int32
	_ = v1826
	var v1827 int32
	_ = v1827
	var v1831 int32
	_ = v1831
	var v1843 int32
	_ = v1843
	var v1844 int32
	_ = v1844
	var v1845 int32
	_ = v1845
	var v1852 int32
	_ = v1852
	var v1857 int32
	_ = v1857
	var v1890 int32
	_ = v1890
	var v1894 int32
	_ = v1894
	var v1899 int32
	_ = v1899
	var v1910 int32
	_ = v1910
	var v1911 int32
	_ = v1911
	var v1914 int32
	_ = v1914
	var v1917 int32
	_ = v1917
	var v1918 int32
	_ = v1918
	var v1923 int32
	_ = v1923
	var v1924 int32
	_ = v1924
	var v1937 int32
	_ = v1937
	var v1939 int32
	_ = v1939
	var v1960 int32
	_ = v1960
	var v1966 int32
	_ = v1966
	var v1971 int32
	_ = v1971
	var v1983 int32
	_ = v1983
	var v1990 int32
	_ = v1990
	var v1991 int32
	_ = v1991
	var v2007 int32
	_ = v2007
	var v2011 int32
	_ = v2011
	var v2014 int32
	_ = v2014
	var v2015 int32
	_ = v2015
	var v2016 int32
	_ = v2016
	var v2017 int32
	_ = v2017
	var v2019 int32
	_ = v2019
	var v2028 int32
	_ = v2028
	var v2031 int32
	_ = v2031
	var v2036 int32
	_ = v2036
	var v2037 int32
	_ = v2037
	var v2060 int32
	_ = v2060
	var v2062 int32
	_ = v2062
	var v2063 int32
	_ = v2063
	var v2071 int32
	_ = v2071
	var v2075 int32
	_ = v2075
	var v2076 int32
	_ = v2076
	var v2078 int32
	_ = v2078
	var v2079 int32
	_ = v2079
	var v2086 int32
	_ = v2086
	var v2087 int32
	_ = v2087
	var v2089 int32
	_ = v2089
	var v2093 int32
	_ = v2093
	var v2097 int32
	_ = v2097
	var v2102 int32
	_ = v2102
	var v2103 int32
	_ = v2103
	var v2104 int32
	_ = v2104
	var v2105 int32
	_ = v2105
	var v2110 int32
	_ = v2110
	var v2111 int32
	_ = v2111
	var v2115 int32
	_ = v2115
	var v2120 int32
	_ = v2120
	var v2124 int32
	_ = v2124
	var v2125 int64
	_ = v2125
	var v2127 int32
	_ = v2127
	var v2129 int32
	_ = v2129
	var v2136 int32
	_ = v2136
	var v2137 int32
	_ = v2137
	var v2153 int32
	_ = v2153
	var v2154 int32
	_ = v2154
	var v2157 int32
	_ = v2157
	var v2160 int32
	_ = v2160
	var v2161 int32
	_ = v2161
	var v2162 int32
	_ = v2162
	var v2163 int32
	_ = v2163
	var v2165 int32
	_ = v2165
	var v2174 int32
	_ = v2174
	var v2177 int32
	_ = v2177
	var v2182 int32
	_ = v2182
	var v2183 int32
	_ = v2183
	var v2204 int32
	_ = v2204
	var v2206 int32
	_ = v2206
	var v2208 int32
	_ = v2208
	var v2209 int32
	_ = v2209
	var v2217 int32
	_ = v2217
	var v2221 int32
	_ = v2221
	var v2222 int32
	_ = v2222
	var v2232 int32
	_ = v2232
	var v2236 int32
	_ = v2236
	var v2238 int32
	_ = v2238
	var v2243 int32
	_ = v2243
	var v2246 int32
	_ = v2246
	var v2249 int32
	_ = v2249
	var v2252 int32
	_ = v2252
	var v2257 int32
	_ = v2257
	var v2260 int32
	_ = v2260
	var v2263 int32
	_ = v2263
	var v2267 int32
	_ = v2267
	var v2275 int32
	_ = v2275
	var v2278 int32
	_ = v2278
	var v2301 int32
	_ = v2301
	var v2302 int32
	_ = v2302
	var v2304 int32
	_ = v2304
	var v2305 int32
	_ = v2305
	var v2308 int32
	_ = v2308
	var v2315 int32
	_ = v2315
	var v2317 int32
	_ = v2317
	var v2319 int32
	_ = v2319
	var v2324 int32
	_ = v2324
	var v2329 int32
	_ = v2329
	var v2330 int32
	_ = v2330
	var v2332 int32
	_ = v2332
	var v2333 int32
	_ = v2333
	var v2335 int32
	_ = v2335
	var v2338 int32
	_ = v2338
	var v2340 int32
	_ = v2340
	var v2343 int32
	_ = v2343
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
		v2308 = v27
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v2315 = int32(_a_F_sparsevec_out_0)
	*(*uint16)(unsafe.Add(mBase, uint32(v2308))) = uint16(v2315)
	v2317 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	v2319 = v2308 + int32(2)
	if int32(0) <= v2317 {
		goto L347
	} else {
		goto L348
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
	v81 = v79 & int32(_a_F_sparsevec_out_1)
	v84 = int32(255)
	v85 = int32(base.Ui32(v79)>>(uint(int32(23))%32)) & v84
	if v85|v81 != 0 {
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
	v1159 = v1158 + v61
	v1160 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	if v1160 <= int32(1) {
		v2308 = v1159
		goto L6
	} else {
		goto L175
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
	v94 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_sparsevec_out[0])))
	*(*uint8)(unsafe.Add(mBase, uint32(v61)+2)) = uint8(v94)
	v97 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_sparsevec_out[1])))
	*(*uint16)(unsafe.Add(mBase, uint32(v61))) = uint16(v97)
	v1158 = int32(3)
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
	v1158 = v113
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
	v1158 = v120
	goto L12
L34:
	;
	v768 = int32(0)
	if v79 < v768 {
		goto L102
	} else {
		goto L103
	}
L35:
	;
	if base.Ui32(int32(_a_F_sparsevec_out_2)) < base.Ui32(v709) {
		v751 = v709
		v756 = v714
		v767 = int32(8)
		goto L34
	} else {
		goto L93
	}
L36:
	;
	v137 = v81 << (uint(int32(2)) % 32)
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
	v709 = int32(base.Ui32(v81|int32(_a_F_sparsevec_out_3)) >> (uint(v127) % 32))
	v714 = v62
	goto L35
L39:
	;
	v140 = v137 | int32(33554432)
	goto L41
L40:
	;
	v140 = v137
	goto L41
L41:
	;
	v143 = int32(2)
	v148 = v140 + (base.B2i32(v81 != int32(0)) | base.B2i32(base.Ui32(v85) < base.Ui32(v143)) ^ int32(-1))
	v150 = v140 | v143
	if v85 != 0 {
		goto L45
	} else {
		goto L46
	}
L42:
	;
	v701 = v684 + v688
	v702 = v683 + v700
	if base.Ui32(v702) <= base.Ui32(int32(99999999)) {
		v709 = v702
		v714 = v701
		goto L35
	} else {
		goto L92
	}
L43:
	;
	v619 = int32(0)
	v620 = int32(10)
	v621 = base.I32_div_u_s(v605, v620)
	v623 = base.I32_div_u_s(v609, v620)
	if base.Ui32(v623) < base.Ui32(v621) {
		goto L86
	} else {
		goto L87
	}
L44:
	;
	v523 = int32(0)
	v524 = int32(10)
	v525 = base.I32_div_u_s(v509, v524)
	v527 = base.I32_div_u_s(v513, v524)
	if base.Ui32(v525) <= base.Ui32(v527) {
		goto L80
	} else {
		goto L81
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
	v160 = int32(base.Ui32(v154*int32(_a_F_sparsevec_out_4)) >> (uint(int32(18)) % 32))
	v163 = *(*int64)(unsafe.Add(mBase, uint32(v160<<(uint(int32(3))%32))+uint32(_c_F_sparsevec_out[2])))
	v165 = v163 & int64(4294967295)
	v166 = base.I64_extend_i32_u(v148)
	v168 = int64(32)
	v170 = base.I32_wrap_i64(int64(base.Ui64(v165*v166) >> (uint(v168) % 64)))
	v172 = int64(base.Ui64(v163) >> (uint(v168) % 64))
	v173 = v166 * v172
	v175 = v170 + base.I32_wrap_i64(v173)
	v182 = v160 - v154
	v187 = v182 + int32(base.Ui32(v160*int32(_a_F_sparsevec_out_5))>>(uint(int32(19))%32))
	v188 = int32(5) - v187
	v191 = v187 + int32(27)
	v193 = (base.B2i32(base.Ui32(v175) < base.Ui32(v170))+base.I32_wrap_i64(int64(base.Ui64(v173)>>(uint(v168)%64))))<<(uint(v188)%32) | int32(base.Ui32(v175)>>(uint(v191)%32))
	v194 = base.I64_extend_i32_u(v150)
	v198 = base.I32_wrap_i64(int64(base.Ui64(v165*v194) >> (uint(v168) % 64)))
	v199 = v194 * v172
	v201 = v198 + base.I32_wrap_i64(v199)
	v209 = (base.B2i32(base.Ui32(v201) < base.Ui32(v198))+base.I32_wrap_i64(int64(base.Ui64(v199)>>(uint(v168)%64))))<<(uint(v188)%32) | int32(base.Ui32(v201)>>(uint(v191)%32))
	v210 = base.I64_extend_i32_u(v140)
	v214 = base.I32_wrap_i64(int64(base.Ui64(v165*v210) >> (uint(v168) % 64)))
	v215 = v210 * v172
	v217 = v214 + base.I32_wrap_i64(v215)
	v225 = (base.B2i32(base.Ui32(v217) < base.Ui32(v214))+base.I32_wrap_i64(int64(base.Ui64(v215)>>(uint(v168)%64))))<<(uint(v188)%32) | int32(base.Ui32(v217)>>(uint(v191)%32))
	v226 = int32(0)
	if v160 != 0 {
		goto L51
	} else {
		goto L52
	}
L49:
	;
	goto L50
L50:
	;
	v369 = v154 * int32(-732923)
	v371 = int32(base.Ui32(v369) >> (uint(int32(20)) % 32))
	v372 = v154 + v371
	v376 = *(*int64)(unsafe.Add(mBase, uint32(int32(_a_F_sparsevec_out_6)-v372<<(uint(int32(3))%32))))
	v378 = v376 & int64(4294967295)
	v379 = base.I64_extend_i32_u(v148)
	v381 = int64(32)
	v383 = base.I32_wrap_i64(int64(base.Ui64(v378*v379) >> (uint(v381) % 64)))
	v385 = int64(base.Ui64(v376) >> (uint(v381) % 64))
	v386 = v379 * v385
	v388 = v383 + base.I32_wrap_i64(v386)
	v399 = v371 - int32(base.Ui32(v372*int32(-1217359))>>(uint(int32(19))%32))
	v400 = int32(4) - v399
	v403 = v399 + int32(28)
	v405 = (base.B2i32(base.Ui32(v388) < base.Ui32(v383))+base.I32_wrap_i64(int64(base.Ui64(v386)>>(uint(v381)%64))))<<(uint(v400)%32) | int32(base.Ui32(v388)>>(uint(v403)%32))
	v406 = base.I64_extend_i32_u(v140)
	v410 = base.I32_wrap_i64(int64(base.Ui64(v378*v406) >> (uint(v381) % 64)))
	v411 = v406 * v385
	v413 = v410 + base.I32_wrap_i64(v411)
	v421 = (base.B2i32(base.Ui32(v413) < base.Ui32(v410))+base.I32_wrap_i64(int64(base.Ui64(v411)>>(uint(v381)%64))))<<(uint(v400)%32) | int32(base.Ui32(v413)>>(uint(v403)%32))
	v422 = base.I64_extend_i32_u(v150)
	v426 = base.I32_wrap_i64(int64(base.Ui64(v378*v422) >> (uint(v381) % 64)))
	v427 = v385 * v422
	v429 = v426 + base.I32_wrap_i64(v427)
	v437 = (base.B2i32(base.Ui32(v429) < base.Ui32(v426))+base.I32_wrap_i64(int64(base.Ui64(v427)>>(uint(v381)%64))))<<(uint(v400)%32) | int32(base.Ui32(v429)>>(uint(v403)%32))
	v439 = v437 - int32(1)
	if v371 != 0 {
		goto L72
	} else {
		goto L73
	}
L51:
	;
	v230 = int32(10)
	v231 = base.I32_div_u_s(v209-int32(1), v230)
	v233 = base.I32_div_u_s(v193, v230)
	if base.Ui32(v231) <= base.Ui32(v233) {
		goto L54
	} else {
		goto L55
	}
L52:
	;
	v277 = v226
	goto L53
L53:
	;
	v283 = base.I32_rem_u_s(v140, int32(5))
	if v283 == int32(0) {
		goto L58
	} else {
		goto L59
	}
L54:
	;
	v236 = v160 - int32(1)
	v239 = *(*int64)(unsafe.Add(mBase, uint32(v236<<(uint(int32(3))%32))+uint32(_c_F_sparsevec_out[2])))
	v243 = int64(32)
	v245 = base.I32_wrap_i64(int64(base.Ui64(v239&int64(4294967295)*v210) >> (uint(v243) % 64)))
	v248 = int64(base.Ui64(v239)>>(uint(v243)%64)) * v210
	v250 = v245 + base.I32_wrap_i64(v248)
	v261 = v182 + int32(base.Ui32(v236*int32(_a_F_sparsevec_out_5))>>(uint(int32(19))%32))
	v269 = base.I32_rem_u_s((base.B2i32(base.Ui32(v250) < base.Ui32(v245))+base.I32_wrap_i64(int64(base.Ui64(v248)>>(uint(v243)%64))))<<(uint(int32(6)-v261)%32)|int32(base.Ui32(v250)>>(uint(v261+int32(26))%32)), int32(10))
	v270 = v269
	goto L56
L55:
	;
	v270 = v226
	goto L56
L56:
	;
	if base.Ui32(int32(33)) < base.Ui32(v154) {
		v602 = v225
		v605 = v209
		v607 = v160
		v608 = v270
		v609 = v193
		goto L43
	} else {
		goto L57
	}
L57:
	;
	v277 = v270
	goto L53
L58:
	;
	v289 = v140
	v290 = v226
	goto L61
L59:
	;
	goto L60
L60:
	;
	v314 = int32(0)
	v316 = base.I32_rem_u_s(v150, int32(5))
	if v316 == v314 {
		goto L65
	} else {
		goto L66
	}
L61:
	;
	v306 = v290 + int32(1)
	v307 = int32(5)
	v308 = base.I32_div_u_s(v289, v307)
	v310 = base.I32_rem_u_s(v308, v307)
	if v310 == int32(0) {
		v289 = v308
		v290 = v306
		goto L61
	} else {
		goto L63
	}
L62:
	;
	if base.Ui32(v306) < base.Ui32(v160) {
		v602 = v225
		v605 = v209
		v607 = v160
		v608 = v277
		v609 = v193
		goto L43
	} else {
		goto L64
	}
L63:
	;
	goto L62
L64:
	;
	v506 = v225
	v509 = v209
	v511 = v160
	v512 = v277
	v513 = v193
	goto L44
L65:
	;
	v322 = v314
	v325 = v150
	goto L68
L66:
	;
	v349 = v314
	goto L67
L67:
	;
	v602 = v225
	v605 = v209 - base.B2i32(base.Ui32(v160) <= base.Ui32(v349))
	v607 = v160
	v608 = v277
	v609 = v193
	goto L43
L68:
	;
	v339 = v322 + int32(1)
	v340 = int32(5)
	v341 = base.I32_div_u_s(v325, v340)
	v343 = base.I32_rem_u_s(v341, v340)
	if v343 == int32(0) {
		v322 = v339
		v325 = v341
		goto L68
	} else {
		goto L70
	}
L69:
	;
	v349 = v339
	goto L67
L70:
	;
	goto L69
L71:
	;
	v494 = int32(-1)
	if v140&(v494<<(uint(v371-int32(1))%32)^v494)|base.B2i32(base.Ui32(int32(32505855)) < base.Ui32(v369)) != 0 {
		v602 = v421
		v605 = v437
		v607 = v372
		v608 = v483
		v609 = v405
		goto L43
	} else {
		goto L79
	}
L72:
	;
	v440 = int32(10)
	v441 = base.I32_div_u_s(v439, v440)
	v443 = base.I32_div_u_s(v405, v440)
	if base.Ui32(v441) <= base.Ui32(v443) {
		goto L75
	} else {
		goto L76
	}
L73:
	;
	v490 = v62
	goto L74
L74:
	;
	v506 = v421
	v509 = v439
	v511 = v372
	v512 = v490
	v513 = v405
	goto L44
L75:
	;
	v446 = int32(1) - v372
	v449 = *(*int64)(unsafe.Add(mBase, uint32(v446<<(uint(int32(3))%32))+uint32(_c_F_sparsevec_out[3])))
	v453 = int64(32)
	v455 = base.I32_wrap_i64(int64(base.Ui64(v449&int64(4294967295)*v406) >> (uint(v453) % 64)))
	v458 = int64(base.Ui64(v449)>>(uint(v453)%64)) * v406
	v460 = v455 + base.I32_wrap_i64(v458)
	v473 = v371 + (int32(base.Ui32(v446*int32(_a_F_sparsevec_out_5))>>(uint(int32(19))%32)) ^ int32(-1))
	v481 = base.I32_rem_u_s((base.B2i32(base.Ui32(v460) < base.Ui32(v455))+base.I32_wrap_i64(int64(base.Ui64(v458)>>(uint(v453)%64))))<<(uint(int32(4)-v473)%32)|int32(base.Ui32(v460)>>(uint(v473+int32(28))%32)), int32(10))
	v483 = v481
	goto L77
L76:
	;
	v483 = v62
	goto L77
L77:
	;
	if v371 != int32(1) {
		goto L71
	} else {
		goto L78
	}
L78:
	;
	v490 = v483
	goto L74
L79:
	;
	v506 = v421
	v509 = v437
	v511 = v372
	v512 = v483
	v513 = v405
	goto L44
L80:
	;
	v571 = v506
	v572 = v523
	v577 = v512
	v578 = v513
	v588 = int32(0)
	goto L82
L81:
	;
	v533 = v506
	v534 = v523
	v535 = v525
	v536 = int32(1)
	v537 = v527
	v539 = v512
	goto L83
L82:
	;
	v590 = v577 & int32(255)
	v683 = v571
	v684 = v572
	v688 = v511
	v700 = (v588|base.B2i32(v590 != int32(5))|v571)&base.B2i32(base.Ui32(int32(4)) < base.Ui32(v590)) | base.B2i32(v571 == v578)
	goto L42
L83:
	;
	v551 = v534 + int32(1)
	v552 = int32(10)
	v553 = base.I32_div_u_s(v533, v552)
	v556 = v533 - v553*v552
	v561 = v536 & base.B2i32(v539&int32(255) == int32(0))
	v563 = base.I32_div_u_s(v535, v552)
	v565 = base.I32_div_u_s(v537, v552)
	if base.Ui32(v565) < base.Ui32(v563) {
		v533 = v553
		v534 = v551
		v535 = v563
		v536 = v561
		v537 = v565
		v539 = v556
		goto L83
	} else {
		goto L85
	}
L84:
	;
	v571 = v553
	v572 = v551
	v577 = v556
	v578 = v537
	v588 = v561 ^ int32(1)
	goto L82
L85:
	;
	goto L84
L86:
	;
	v627 = v602
	v628 = v619
	v629 = v621
	v631 = v623
	goto L89
L87:
	;
	v658 = v602
	v659 = v619
	v664 = v608
	v665 = v609
	goto L88
L88:
	;
	v683 = v658
	v684 = v659
	v688 = v607
	v700 = base.B2i32(v658 == v665) | base.B2i32(base.Ui32(int32(4)) < base.Ui32(v664&int32(255)))
	goto L42
L89:
	;
	v645 = v628 + int32(1)
	v646 = int32(10)
	v647 = base.I32_div_u_s(v627, v646)
	v649 = base.I32_div_u_s(v629, v646)
	v651 = base.I32_div_u_s(v631, v646)
	if base.Ui32(v651) < base.Ui32(v649) {
		v627 = v647
		v628 = v645
		v629 = v649
		v631 = v651
		goto L89
	} else {
		goto L91
	}
L90:
	;
	v658 = v647
	v659 = v645
	v664 = v627 - v647*int32(10)
	v665 = v631
	goto L88
L91:
	;
	goto L90
L92:
	;
	v751 = v702
	v756 = v701
	v767 = int32(9)
	goto L34
L93:
	;
	if base.Ui32(int32(_a_F_sparsevec_out_7)) < base.Ui32(v709) {
		v751 = v709
		v756 = v714
		v767 = int32(7)
		goto L34
	} else {
		goto L94
	}
L94:
	;
	if base.Ui32(int32(_a_F_sparsevec_out_8)) < base.Ui32(v709) {
		v751 = v709
		v756 = v714
		v767 = int32(6)
		goto L34
	} else {
		goto L95
	}
L95:
	;
	if base.Ui32(int32(_a_F_sparsevec_out_9)) < base.Ui32(v709) {
		v751 = v709
		v756 = v714
		v767 = int32(5)
		goto L34
	} else {
		goto L96
	}
L96:
	;
	if base.Ui32(int32(999)) < base.Ui32(v709) {
		v751 = v709
		v756 = v714
		v767 = int32(4)
		goto L34
	} else {
		goto L97
	}
L97:
	;
	if base.Ui32(int32(99)) < base.Ui32(v709) {
		v751 = v709
		v756 = v714
		v767 = int32(3)
		goto L34
	} else {
		goto L98
	}
L98:
	;
	if base.Ui32(int32(9)) < base.Ui32(v709) {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	v747 = int32(2)
	goto L101
L100:
	;
	v747 = int32(1)
	goto L101
L101:
	;
	v751 = v709
	v756 = v714
	v767 = v747
	goto L34
L102:
	;
	v771 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(v61))) = uint8(v771)
	v774 = int32(1)
	goto L104
L103:
	;
	v774 = v768
	goto L104
L104:
	;
	v775 = v767 + v756
	if base.Ui32(v775+int32(3)) <= base.Ui32(int32(9)) {
		goto L107
	} else {
		goto L108
	}
L105:
	;
	v986 = int32(0)
	if base.Ui32(int32(_a_F_sparsevec_out_10)) <= base.Ui32(v751) {
		goto L145
	} else {
		goto L146
	}
L106:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v780))) = v982
	v984 = v981
	goto L105
L107:
	;
	v780 = v61 + v774
	v781 = int32(0)
	if v775 <= v781 {
		goto L110
	} else {
		goto L111
	}
L108:
	;
	goto L109
L109:
	;
	if v756 != 0 {
		goto L115
	} else {
		goto L116
	}
L110:
	;
	v981 = int32(2) - v775
	v982 = int64(3472328296227679792)
	goto L106
L111:
	;
	goto L112
L112:
	;
	if int32(0) <= v756 {
		v981 = v781
		v982 = int64(3472328296227680304)
		goto L106
	} else {
		goto L113
	}
L113:
	;
	v984 = int32(1)
	goto L105
L114:
	;
	v840 = int32(0)
	if base.Ui32(int32(_a_F_sparsevec_out_10)) <= base.Ui32(v823) {
		goto L122
	} else {
		goto L123
	}
L115:
	;
	v823 = v751
	v828 = v767
	goto L114
L116:
	;
	goto L117
L117:
	;
	v794 = v751
	v796 = v767
	goto L118
L118:
	;
	if v794&int32(1) != 0 {
		v823 = v794
		v828 = v796
		goto L114
	} else {
		goto L120
	}
L119:
	;
	v823 = v794
	v828 = v796
	goto L114
L120:
	;
	v817 = base.I32_div_u_s(v794, int32(10))
	if int32(0)-v794 == v817*int32(-10) {
		v794 = v817
		v796 = v796 - int32(1)
		goto L118
	} else {
		goto L121
	}
L121:
	;
	goto L119
L122:
	;
	v847 = v823
	v848 = v840
	goto L125
L123:
	;
	v893 = v823
	v894 = v840
	goto L124
L124:
	;
	if base.Ui32(v893) < base.Ui32(int32(100)) {
		goto L129
	} else {
		goto L130
	}
L125:
	;
	v864 = v61 + v774 + v828 - v848
	v868 = base.I32_div_u_s(v847, int32(_a_F_sparsevec_out_10))
	v871 = v847 + v868*int32(-10000)
	v872 = int32(100)
	v873 = base.I32_div_u_s(v871, v872)
	v874 = int32(1)
	v876 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v873<<(uint(v874)%32))+uint32(_c_F_sparsevec_out[4]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v864-int32(3)))) = uint16(v876)
	v885 = int32(*(*uint16)(unsafe.Add(mBase, uint32((v871-v873*v872)<<(uint(v874)%32))+uint32(_c_F_sparsevec_out[4]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v864-v874))) = uint16(v885)
	v888 = v848 + int32(4)
	if base.Ui32(int32(99999999)) < base.Ui32(v847) {
		v847 = v868
		v848 = v888
		goto L125
	} else {
		goto L127
	}
L126:
	;
	v893 = v868
	v894 = v888
	goto L124
L127:
	;
	goto L126
L128:
	;
	v935 = v775 - int32(1)
	v936 = v61 + v774
	if base.Ui32(int32(10)) <= base.Ui32(v933) {
		goto L133
	} else {
		goto L134
	}
L129:
	;
	v932 = v894
	v933 = v893
	goto L128
L130:
	;
	goto L131
L131:
	;
	v917 = int32(_a_F_sparsevec_out_11)
	v919 = int32(100)
	v920 = base.I32_div_u_s(v893&v917, v919)
	v928 = int32(*(*uint16)(unsafe.Add(mBase, uint32((v893-v920*v919)&v917<<(uint(int32(1))%32))+uint32(_c_F_sparsevec_out[4]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v61+v774+v828+(v894^int32(-1))))) = uint16(v928)
	v932 = v894 | int32(2)
	v933 = v920
	goto L128
L132:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v936))) = uint8(v950)
	if base.Ui32(int32(2)) <= base.Ui32(v828) {
		goto L136
	} else {
		goto L137
	}
L133:
	;
	v943 = v933 << (uint(int32(1)) % 32)
	v944 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v943)+uint32(_c_F_sparsevec_out[5]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v61+(v774+v828-v932)))) = uint8(v944)
	v946 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v943)+uint32(_c_F_sparsevec_out[4]))))
	v950 = v946
	goto L132
L134:
	;
	goto L135
L135:
	;
	v950 = v933 | int32(48)
	goto L132
L136:
	;
	v954 = int32(46)
	*(*uint8)(unsafe.Add(mBase, uint32(v936)+1)) = uint8(v954)
	v959 = v828 + int32(1)
	goto L138
L137:
	;
	v959 = int32(1)
	goto L138
L138:
	;
	v960 = v959 + v774
	v961 = v61 + v960
	v962 = int32(101)
	*(*uint8)(unsafe.Add(mBase, uint32(v961))) = uint8(v962)
	v967 = base.B2i32(v935 < int32(0))
	if v935 < int32(0) {
		goto L139
	} else {
		goto L140
	}
L139:
	;
	v968 = int32(45)
	goto L141
L140:
	;
	v968 = int32(43)
	goto L141
L141:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v961)+1)) = uint8(v968)
	if v935 < int32(0) {
		goto L142
	} else {
		goto L143
	}
L142:
	;
	v972 = int32(1) - v775
	goto L144
L143:
	;
	v972 = v935
	goto L144
L144:
	;
	v977 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v972<<(uint(int32(1))%32))+uint32(_c_F_sparsevec_out[4]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v961)+2)) = uint16(v977)
	v1158 = v960 + int32(4)
	goto L12
L145:
	;
	v993 = v986
	v994 = v751
	goto L148
L146:
	;
	v1039 = v986
	v1040 = v751
	goto L147
L147:
	;
	if base.Ui32(v1040) < base.Ui32(int32(100)) {
		goto L152
	} else {
		goto L153
	}
L148:
	;
	v1010 = v984 + v780 + v767 - v993
	v1011 = int32(4)
	v1014 = base.I32_div_u_s(v994, int32(_a_F_sparsevec_out_10))
	v1017 = v994 + v1014*int32(-10000)
	v1018 = int32(100)
	v1019 = base.I32_div_u_s(v1017, v1018)
	v1020 = int32(1)
	v1022 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1019<<(uint(v1020)%32))+uint32(_c_F_sparsevec_out[4]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v1010-v1011))) = uint16(v1022)
	v1031 = int32(*(*uint16)(unsafe.Add(mBase, uint32((v1017-v1019*v1018)<<(uint(v1020)%32))+uint32(_c_F_sparsevec_out[4]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v1010-int32(2)))) = uint16(v1031)
	v1034 = v993 + v1011
	if base.Ui32(int32(99999999)) < base.Ui32(v994) {
		v993 = v1034
		v994 = v1014
		goto L148
	} else {
		goto L150
	}
L149:
	;
	v1039 = v1034
	v1040 = v1014
	goto L147
L150:
	;
	goto L149
L151:
	;
	if base.Ui32(int32(10)) <= base.Ui32(v1079) {
		goto L156
	} else {
		goto L157
	}
L152:
	;
	v1078 = v1039
	v1079 = v1040
	goto L151
L153:
	;
	goto L154
L154:
	;
	v1061 = int32(2)
	v1063 = int32(_a_F_sparsevec_out_11)
	v1065 = int32(100)
	v1066 = base.I32_div_u_s(v1040&v1063, v1065)
	v1074 = int32(*(*uint16)(unsafe.Add(mBase, uint32((v1040-v1066*v1065)&v1063<<(uint(int32(1))%32))+uint32(_c_F_sparsevec_out[4]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v984+v780+v767-v1039-v1061))) = uint16(v1074)
	v1078 = v1039 | v1061
	v1079 = v1066
	goto L151
L155:
	;
	v1095 = int32(1)
	if v984 == v1095 {
		goto L160
	} else {
		goto L161
	}
L156:
	;
	v1089 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1079<<(uint(int32(1))%32))+uint32(_c_F_sparsevec_out[4]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v984+v780+v767-v1078-int32(2)))) = uint16(v1089)
	goto L155
L157:
	;
	goto L158
L158:
	;
	v1093 = v1079 | int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v984+v780))) = uint8(v1093)
	goto L155
L159:
	;
	v1158 = v1135 + int32(base.Ui32(v79)>>(uint(int32(31))%32))
	goto L12
L160:
	;
	if v775&int32(4) != 0 {
		goto L163
	} else {
		goto L164
	}
L161:
	;
	goto L162
L162:
	;
	if v756 < int32(0) {
		goto L172
	} else {
		goto L173
	}
L163:
	;
	v1100 = *(*int32)(unsafe.Add(mBase, uint32(v780)+1))
	*(*int32)(unsafe.Add(mBase, uint32(v780))) = v1100
	v1103 = int32(5)
	goto L165
L164:
	;
	v1103 = v1095
	goto L165
L165:
	;
	if v775&int32(2) != 0 {
		goto L166
	} else {
		goto L167
	}
L166:
	;
	v1106 = v1103 + v780
	v1109 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1106))))
	*(*uint16)(unsafe.Add(mBase, uint32(v1106-int32(1)))) = uint16(v1109)
	v1114 = v1103 | int32(2)
	goto L168
L167:
	;
	v1114 = v1103
	goto L168
L168:
	;
	if v775&int32(1) != 0 {
		goto L169
	} else {
		goto L170
	}
L169:
	;
	v1117 = v1114 + v780
	v1120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1117))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1117-int32(1)))) = uint8(v1120)
	goto L171
L170:
	;
	goto L171
L171:
	;
	v1124 = int32(46)
	*(*uint8)(unsafe.Add(mBase, uint32(v780+v775))) = uint8(v1124)
	v1135 = v767 + int32(1)
	goto L159
L172:
	;
	v1132 = int32(2) - v756
	goto L174
L173:
	;
	v1132 = v775
	goto L174
L174:
	;
	v1135 = v1132
	goto L159
L175:
	;
	v1164 = v1159
	v1167 = v25
	goto L176
L176:
	;
	v1171 = int32(44)
	*(*uint8)(unsafe.Add(mBase, uint32(v1164))) = uint8(v1171)
	v1174 = v1167 << (uint(int32(2)) % 32)
	v1176 = *(*int32)(unsafe.Add(mBase, uint32(v55+v1174)))
	v1177 = int32(1)
	v1178 = v1176 + v1177
	v1180 = v1164 + v1177
	if int32(0) <= v1178 {
		goto L179
	} else {
		goto L180
	}
L177:
	;
	v2308 = v2302
	goto L6
L178:
	;
	v1198 = v1194 + v1180
	v1199 = int32(58)
	*(*uint8)(unsafe.Add(mBase, uint32(v1198))) = uint8(v1199)
	v1202 = *(*float32)(unsafe.Add(mBase, uint32(v58+v1174)))
	v1204 = v1198 + int32(1)
	v1205 = int32(0)
	v1222 = base.I32_reinterpret_f32(v1202)
	v1224 = v1222 & int32(_a_F_sparsevec_out_1)
	v1227 = int32(255)
	v1228 = int32(base.Ui32(v1222)>>(uint(int32(23))%32)) & v1227
	if v1228|v1224 != 0 {
		goto L183
	} else {
		goto L184
	}
L179:
	;
	v1190 = v1178
	v1191 = int32(0)
	goto L181
L180:
	;
	v1185 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(v1180))) = uint8(v1185)
	v1190 = int32(0) - v1178
	v1191 = int32(1)
	goto L181
L181:
	;
	v1193 = F_pg_ultoa_n(m, v1190, v1180+v1191)
	mBase = m.M
	v1194 = v1193 + v1191
	v1196 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1180+v1194))) = uint8(v1196)
	goto L178
L182:
	;
	v2302 = v2301 + v1204
	v2304 = v1167 + int32(1)
	v2305 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	if v2304 < v2305 {
		v1164 = v2302
		v1167 = v2304
		goto L176
	} else {
		goto L345
	}
L183:
	;
	v1233 = base.B2i32(v1228 != v1227)
	goto L185
L184:
	;
	v1233 = v1205
	goto L185
L185:
	;
	if v1233 == int32(0) {
		goto L186
	} else {
		goto L187
	}
L186:
	;
	if v1224 != 0 {
		goto L189
	} else {
		goto L190
	}
L187:
	;
	goto L188
L188:
	;
	if base.Ui32(int32(23)) < base.Ui32(v1228-int32(127)) {
		goto L206
	} else {
		goto L207
	}
L189:
	;
	v1237 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_sparsevec_out[0])))
	*(*uint8)(unsafe.Add(mBase, uint32(v1204)+2)) = uint8(v1237)
	v1240 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_sparsevec_out[1])))
	*(*uint16)(unsafe.Add(mBase, uint32(v1204))) = uint16(v1240)
	v2301 = int32(3)
	goto L182
L190:
	;
	goto L191
L191:
	;
	if v1222 < int32(0) {
		goto L192
	} else {
		goto L193
	}
L192:
	;
	v1245 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(v1204))) = uint8(v1245)
	goto L194
L193:
	;
	goto L194
L194:
	;
	v1249 = v1204 + int32(base.Ui32(v1222)>>(uint(int32(31))%32))
	if v1228 != 0 {
		goto L195
	} else {
		goto L196
	}
L195:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1249))) = int64(8751735898823355977)
	if v1222 < int32(0) {
		goto L198
	} else {
		goto L199
	}
L196:
	;
	goto L197
L197:
	;
	v1257 = int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v1249))) = uint8(v1257)
	if v1222 < int32(0) {
		goto L201
	} else {
		goto L202
	}
L198:
	;
	v1256 = int32(9)
	goto L200
L199:
	;
	v1256 = int32(8)
	goto L200
L200:
	;
	v2301 = v1256
	goto L182
L201:
	;
	v1263 = int32(2)
	goto L203
L202:
	;
	v1263 = int32(1)
	goto L203
L203:
	;
	v2301 = v1263
	goto L182
L204:
	;
	v1911 = int32(0)
	if v1222 < v1911 {
		goto L272
	} else {
		goto L273
	}
L205:
	;
	if base.Ui32(int32(_a_F_sparsevec_out_2)) < base.Ui32(v1852) {
		v1894 = v1852
		v1899 = v1857
		v1910 = int32(8)
		goto L204
	} else {
		goto L263
	}
L206:
	;
	v1280 = v1224 << (uint(int32(2)) % 32)
	if v1228 != 0 {
		goto L209
	} else {
		goto L210
	}
L207:
	;
	v1268 = int32(-1)
	v1270 = int32(150) - v1228
	if v1224&(v1268<<(uint(v1270)%32)^v1268) != 0 {
		goto L206
	} else {
		goto L208
	}
L208:
	;
	v1852 = int32(base.Ui32(v1224|int32(_a_F_sparsevec_out_3)) >> (uint(v1270) % 32))
	v1857 = v1205
	goto L205
L209:
	;
	v1283 = v1280 | int32(33554432)
	goto L211
L210:
	;
	v1283 = v1280
	goto L211
L211:
	;
	v1286 = int32(2)
	v1291 = v1283 + (base.B2i32(v1224 != int32(0)) | base.B2i32(base.Ui32(v1228) < base.Ui32(v1286)) ^ int32(-1))
	v1293 = v1283 | v1286
	if v1228 != 0 {
		goto L215
	} else {
		goto L216
	}
L212:
	;
	v1844 = v1827 + v1831
	v1845 = v1826 + v1843
	if base.Ui32(v1845) <= base.Ui32(int32(99999999)) {
		v1852 = v1845
		v1857 = v1844
		goto L205
	} else {
		goto L262
	}
L213:
	;
	v1762 = int32(0)
	v1763 = int32(10)
	v1764 = base.I32_div_u_s(v1748, v1763)
	v1766 = base.I32_div_u_s(v1752, v1763)
	if base.Ui32(v1766) < base.Ui32(v1764) {
		goto L256
	} else {
		goto L257
	}
L214:
	;
	v1666 = int32(0)
	v1667 = int32(10)
	v1668 = base.I32_div_u_s(v1652, v1667)
	v1670 = base.I32_div_u_s(v1656, v1667)
	if base.Ui32(v1668) <= base.Ui32(v1670) {
		goto L250
	} else {
		goto L251
	}
L215:
	;
	v1297 = v1228 - int32(152)
	goto L217
L216:
	;
	v1297 = int32(-151)
	goto L217
L217:
	;
	if int32(0) <= v1297 {
		goto L218
	} else {
		goto L219
	}
L218:
	;
	v1303 = int32(base.Ui32(v1297*int32(_a_F_sparsevec_out_4)) >> (uint(int32(18)) % 32))
	v1306 = *(*int64)(unsafe.Add(mBase, uint32(v1303<<(uint(int32(3))%32))+uint32(_c_F_sparsevec_out[2])))
	v1308 = v1306 & int64(4294967295)
	v1309 = base.I64_extend_i32_u(v1291)
	v1311 = int64(32)
	v1313 = base.I32_wrap_i64(int64(base.Ui64(v1308*v1309) >> (uint(v1311) % 64)))
	v1315 = int64(base.Ui64(v1306) >> (uint(v1311) % 64))
	v1316 = v1309 * v1315
	v1318 = v1313 + base.I32_wrap_i64(v1316)
	v1325 = v1303 - v1297
	v1330 = v1325 + int32(base.Ui32(v1303*int32(_a_F_sparsevec_out_5))>>(uint(int32(19))%32))
	v1331 = int32(5) - v1330
	v1334 = v1330 + int32(27)
	v1336 = (base.B2i32(base.Ui32(v1318) < base.Ui32(v1313))+base.I32_wrap_i64(int64(base.Ui64(v1316)>>(uint(v1311)%64))))<<(uint(v1331)%32) | int32(base.Ui32(v1318)>>(uint(v1334)%32))
	v1337 = base.I64_extend_i32_u(v1293)
	v1341 = base.I32_wrap_i64(int64(base.Ui64(v1308*v1337) >> (uint(v1311) % 64)))
	v1342 = v1337 * v1315
	v1344 = v1341 + base.I32_wrap_i64(v1342)
	v1352 = (base.B2i32(base.Ui32(v1344) < base.Ui32(v1341))+base.I32_wrap_i64(int64(base.Ui64(v1342)>>(uint(v1311)%64))))<<(uint(v1331)%32) | int32(base.Ui32(v1344)>>(uint(v1334)%32))
	v1353 = base.I64_extend_i32_u(v1283)
	v1357 = base.I32_wrap_i64(int64(base.Ui64(v1308*v1353) >> (uint(v1311) % 64)))
	v1358 = v1353 * v1315
	v1360 = v1357 + base.I32_wrap_i64(v1358)
	v1368 = (base.B2i32(base.Ui32(v1360) < base.Ui32(v1357))+base.I32_wrap_i64(int64(base.Ui64(v1358)>>(uint(v1311)%64))))<<(uint(v1331)%32) | int32(base.Ui32(v1360)>>(uint(v1334)%32))
	v1369 = int32(0)
	if v1303 != 0 {
		goto L221
	} else {
		goto L222
	}
L219:
	;
	goto L220
L220:
	;
	v1512 = v1297 * int32(-732923)
	v1514 = int32(base.Ui32(v1512) >> (uint(int32(20)) % 32))
	v1515 = v1297 + v1514
	v1519 = *(*int64)(unsafe.Add(mBase, uint32(int32(_a_F_sparsevec_out_6)-v1515<<(uint(int32(3))%32))))
	v1521 = v1519 & int64(4294967295)
	v1522 = base.I64_extend_i32_u(v1291)
	v1524 = int64(32)
	v1526 = base.I32_wrap_i64(int64(base.Ui64(v1521*v1522) >> (uint(v1524) % 64)))
	v1528 = int64(base.Ui64(v1519) >> (uint(v1524) % 64))
	v1529 = v1522 * v1528
	v1531 = v1526 + base.I32_wrap_i64(v1529)
	v1542 = v1514 - int32(base.Ui32(v1515*int32(-1217359))>>(uint(int32(19))%32))
	v1543 = int32(4) - v1542
	v1546 = v1542 + int32(28)
	v1548 = (base.B2i32(base.Ui32(v1531) < base.Ui32(v1526))+base.I32_wrap_i64(int64(base.Ui64(v1529)>>(uint(v1524)%64))))<<(uint(v1543)%32) | int32(base.Ui32(v1531)>>(uint(v1546)%32))
	v1549 = base.I64_extend_i32_u(v1283)
	v1553 = base.I32_wrap_i64(int64(base.Ui64(v1521*v1549) >> (uint(v1524) % 64)))
	v1554 = v1549 * v1528
	v1556 = v1553 + base.I32_wrap_i64(v1554)
	v1564 = (base.B2i32(base.Ui32(v1556) < base.Ui32(v1553))+base.I32_wrap_i64(int64(base.Ui64(v1554)>>(uint(v1524)%64))))<<(uint(v1543)%32) | int32(base.Ui32(v1556)>>(uint(v1546)%32))
	v1565 = base.I64_extend_i32_u(v1293)
	v1569 = base.I32_wrap_i64(int64(base.Ui64(v1521*v1565) >> (uint(v1524) % 64)))
	v1570 = v1528 * v1565
	v1572 = v1569 + base.I32_wrap_i64(v1570)
	v1580 = (base.B2i32(base.Ui32(v1572) < base.Ui32(v1569))+base.I32_wrap_i64(int64(base.Ui64(v1570)>>(uint(v1524)%64))))<<(uint(v1543)%32) | int32(base.Ui32(v1572)>>(uint(v1546)%32))
	v1582 = v1580 - int32(1)
	if v1514 != 0 {
		goto L242
	} else {
		goto L243
	}
L221:
	;
	v1373 = int32(10)
	v1374 = base.I32_div_u_s(v1352-int32(1), v1373)
	v1376 = base.I32_div_u_s(v1336, v1373)
	if base.Ui32(v1374) <= base.Ui32(v1376) {
		goto L224
	} else {
		goto L225
	}
L222:
	;
	v1420 = v1369
	goto L223
L223:
	;
	v1426 = base.I32_rem_u_s(v1283, int32(5))
	if v1426 == int32(0) {
		goto L228
	} else {
		goto L229
	}
L224:
	;
	v1379 = v1303 - int32(1)
	v1382 = *(*int64)(unsafe.Add(mBase, uint32(v1379<<(uint(int32(3))%32))+uint32(_c_F_sparsevec_out[2])))
	v1386 = int64(32)
	v1388 = base.I32_wrap_i64(int64(base.Ui64(v1382&int64(4294967295)*v1353) >> (uint(v1386) % 64)))
	v1391 = int64(base.Ui64(v1382)>>(uint(v1386)%64)) * v1353
	v1393 = v1388 + base.I32_wrap_i64(v1391)
	v1404 = v1325 + int32(base.Ui32(v1379*int32(_a_F_sparsevec_out_5))>>(uint(int32(19))%32))
	v1412 = base.I32_rem_u_s((base.B2i32(base.Ui32(v1393) < base.Ui32(v1388))+base.I32_wrap_i64(int64(base.Ui64(v1391)>>(uint(v1386)%64))))<<(uint(int32(6)-v1404)%32)|int32(base.Ui32(v1393)>>(uint(v1404+int32(26))%32)), int32(10))
	v1413 = v1412
	goto L226
L225:
	;
	v1413 = v1369
	goto L226
L226:
	;
	if base.Ui32(int32(33)) < base.Ui32(v1297) {
		v1745 = v1368
		v1748 = v1352
		v1750 = v1303
		v1751 = v1413
		v1752 = v1336
		goto L213
	} else {
		goto L227
	}
L227:
	;
	v1420 = v1413
	goto L223
L228:
	;
	v1432 = v1283
	v1433 = v1369
	goto L231
L229:
	;
	goto L230
L230:
	;
	v1457 = int32(0)
	v1459 = base.I32_rem_u_s(v1293, int32(5))
	if v1459 == v1457 {
		goto L235
	} else {
		goto L236
	}
L231:
	;
	v1449 = v1433 + int32(1)
	v1450 = int32(5)
	v1451 = base.I32_div_u_s(v1432, v1450)
	v1453 = base.I32_rem_u_s(v1451, v1450)
	if v1453 == int32(0) {
		v1432 = v1451
		v1433 = v1449
		goto L231
	} else {
		goto L233
	}
L232:
	;
	if base.Ui32(v1449) < base.Ui32(v1303) {
		v1745 = v1368
		v1748 = v1352
		v1750 = v1303
		v1751 = v1420
		v1752 = v1336
		goto L213
	} else {
		goto L234
	}
L233:
	;
	goto L232
L234:
	;
	v1649 = v1368
	v1652 = v1352
	v1654 = v1303
	v1655 = v1420
	v1656 = v1336
	goto L214
L235:
	;
	v1465 = v1457
	v1468 = v1293
	goto L238
L236:
	;
	v1492 = v1457
	goto L237
L237:
	;
	v1745 = v1368
	v1748 = v1352 - base.B2i32(base.Ui32(v1303) <= base.Ui32(v1492))
	v1750 = v1303
	v1751 = v1420
	v1752 = v1336
	goto L213
L238:
	;
	v1482 = v1465 + int32(1)
	v1483 = int32(5)
	v1484 = base.I32_div_u_s(v1468, v1483)
	v1486 = base.I32_rem_u_s(v1484, v1483)
	if v1486 == int32(0) {
		v1465 = v1482
		v1468 = v1484
		goto L238
	} else {
		goto L240
	}
L239:
	;
	v1492 = v1482
	goto L237
L240:
	;
	goto L239
L241:
	;
	v1637 = int32(-1)
	if v1283&(v1637<<(uint(v1514-int32(1))%32)^v1637)|base.B2i32(base.Ui32(int32(32505855)) < base.Ui32(v1512)) != 0 {
		v1745 = v1564
		v1748 = v1580
		v1750 = v1515
		v1751 = v1626
		v1752 = v1548
		goto L213
	} else {
		goto L249
	}
L242:
	;
	v1583 = int32(10)
	v1584 = base.I32_div_u_s(v1582, v1583)
	v1586 = base.I32_div_u_s(v1548, v1583)
	if base.Ui32(v1584) <= base.Ui32(v1586) {
		goto L245
	} else {
		goto L246
	}
L243:
	;
	v1633 = v1205
	goto L244
L244:
	;
	v1649 = v1564
	v1652 = v1582
	v1654 = v1515
	v1655 = v1633
	v1656 = v1548
	goto L214
L245:
	;
	v1589 = int32(1) - v1515
	v1592 = *(*int64)(unsafe.Add(mBase, uint32(v1589<<(uint(int32(3))%32))+uint32(_c_F_sparsevec_out[3])))
	v1596 = int64(32)
	v1598 = base.I32_wrap_i64(int64(base.Ui64(v1592&int64(4294967295)*v1549) >> (uint(v1596) % 64)))
	v1601 = int64(base.Ui64(v1592)>>(uint(v1596)%64)) * v1549
	v1603 = v1598 + base.I32_wrap_i64(v1601)
	v1616 = v1514 + (int32(base.Ui32(v1589*int32(_a_F_sparsevec_out_5))>>(uint(int32(19))%32)) ^ int32(-1))
	v1624 = base.I32_rem_u_s((base.B2i32(base.Ui32(v1603) < base.Ui32(v1598))+base.I32_wrap_i64(int64(base.Ui64(v1601)>>(uint(v1596)%64))))<<(uint(int32(4)-v1616)%32)|int32(base.Ui32(v1603)>>(uint(v1616+int32(28))%32)), int32(10))
	v1626 = v1624
	goto L247
L246:
	;
	v1626 = v1205
	goto L247
L247:
	;
	if v1514 != int32(1) {
		goto L241
	} else {
		goto L248
	}
L248:
	;
	v1633 = v1626
	goto L244
L249:
	;
	v1649 = v1564
	v1652 = v1580
	v1654 = v1515
	v1655 = v1626
	v1656 = v1548
	goto L214
L250:
	;
	v1714 = v1649
	v1715 = v1666
	v1720 = v1655
	v1721 = v1656
	v1731 = int32(0)
	goto L252
L251:
	;
	v1676 = v1649
	v1677 = v1666
	v1678 = v1668
	v1679 = int32(1)
	v1680 = v1670
	v1682 = v1655
	goto L253
L252:
	;
	v1733 = v1720 & int32(255)
	v1826 = v1714
	v1827 = v1715
	v1831 = v1654
	v1843 = (v1731|base.B2i32(v1733 != int32(5))|v1714)&base.B2i32(base.Ui32(int32(4)) < base.Ui32(v1733)) | base.B2i32(v1714 == v1721)
	goto L212
L253:
	;
	v1694 = v1677 + int32(1)
	v1695 = int32(10)
	v1696 = base.I32_div_u_s(v1676, v1695)
	v1699 = v1676 - v1696*v1695
	v1704 = v1679 & base.B2i32(v1682&int32(255) == int32(0))
	v1706 = base.I32_div_u_s(v1678, v1695)
	v1708 = base.I32_div_u_s(v1680, v1695)
	if base.Ui32(v1708) < base.Ui32(v1706) {
		v1676 = v1696
		v1677 = v1694
		v1678 = v1706
		v1679 = v1704
		v1680 = v1708
		v1682 = v1699
		goto L253
	} else {
		goto L255
	}
L254:
	;
	v1714 = v1696
	v1715 = v1694
	v1720 = v1699
	v1721 = v1680
	v1731 = v1704 ^ int32(1)
	goto L252
L255:
	;
	goto L254
L256:
	;
	v1770 = v1745
	v1771 = v1762
	v1772 = v1764
	v1774 = v1766
	goto L259
L257:
	;
	v1801 = v1745
	v1802 = v1762
	v1807 = v1751
	v1808 = v1752
	goto L258
L258:
	;
	v1826 = v1801
	v1827 = v1802
	v1831 = v1750
	v1843 = base.B2i32(v1801 == v1808) | base.B2i32(base.Ui32(int32(4)) < base.Ui32(v1807&int32(255)))
	goto L212
L259:
	;
	v1788 = v1771 + int32(1)
	v1789 = int32(10)
	v1790 = base.I32_div_u_s(v1770, v1789)
	v1792 = base.I32_div_u_s(v1772, v1789)
	v1794 = base.I32_div_u_s(v1774, v1789)
	if base.Ui32(v1794) < base.Ui32(v1792) {
		v1770 = v1790
		v1771 = v1788
		v1772 = v1792
		v1774 = v1794
		goto L259
	} else {
		goto L261
	}
L260:
	;
	v1801 = v1790
	v1802 = v1788
	v1807 = v1770 - v1790*int32(10)
	v1808 = v1774
	goto L258
L261:
	;
	goto L260
L262:
	;
	v1894 = v1845
	v1899 = v1844
	v1910 = int32(9)
	goto L204
L263:
	;
	if base.Ui32(int32(_a_F_sparsevec_out_7)) < base.Ui32(v1852) {
		v1894 = v1852
		v1899 = v1857
		v1910 = int32(7)
		goto L204
	} else {
		goto L264
	}
L264:
	;
	if base.Ui32(int32(_a_F_sparsevec_out_8)) < base.Ui32(v1852) {
		v1894 = v1852
		v1899 = v1857
		v1910 = int32(6)
		goto L204
	} else {
		goto L265
	}
L265:
	;
	if base.Ui32(int32(_a_F_sparsevec_out_9)) < base.Ui32(v1852) {
		v1894 = v1852
		v1899 = v1857
		v1910 = int32(5)
		goto L204
	} else {
		goto L266
	}
L266:
	;
	if base.Ui32(int32(999)) < base.Ui32(v1852) {
		v1894 = v1852
		v1899 = v1857
		v1910 = int32(4)
		goto L204
	} else {
		goto L267
	}
L267:
	;
	if base.Ui32(int32(99)) < base.Ui32(v1852) {
		v1894 = v1852
		v1899 = v1857
		v1910 = int32(3)
		goto L204
	} else {
		goto L268
	}
L268:
	;
	if base.Ui32(int32(9)) < base.Ui32(v1852) {
		goto L269
	} else {
		goto L270
	}
L269:
	;
	v1890 = int32(2)
	goto L271
L270:
	;
	v1890 = int32(1)
	goto L271
L271:
	;
	v1894 = v1852
	v1899 = v1857
	v1910 = v1890
	goto L204
L272:
	;
	v1914 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(v1204))) = uint8(v1914)
	v1917 = int32(1)
	goto L274
L273:
	;
	v1917 = v1911
	goto L274
L274:
	;
	v1918 = v1910 + v1899
	if base.Ui32(v1918+int32(3)) <= base.Ui32(int32(9)) {
		goto L277
	} else {
		goto L278
	}
L275:
	;
	v2129 = int32(0)
	if base.Ui32(int32(_a_F_sparsevec_out_10)) <= base.Ui32(v1894) {
		goto L315
	} else {
		goto L316
	}
L276:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1923))) = v2125
	v2127 = v2124
	goto L275
L277:
	;
	v1923 = v1204 + v1917
	v1924 = int32(0)
	if v1918 <= v1924 {
		goto L280
	} else {
		goto L281
	}
L278:
	;
	goto L279
L279:
	;
	if v1899 != 0 {
		goto L285
	} else {
		goto L286
	}
L280:
	;
	v2124 = int32(2) - v1918
	v2125 = int64(3472328296227679792)
	goto L276
L281:
	;
	goto L282
L282:
	;
	if int32(0) <= v1899 {
		v2124 = v1924
		v2125 = int64(3472328296227680304)
		goto L276
	} else {
		goto L283
	}
L283:
	;
	v2127 = int32(1)
	goto L275
L284:
	;
	v1983 = int32(0)
	if base.Ui32(int32(_a_F_sparsevec_out_10)) <= base.Ui32(v1966) {
		goto L292
	} else {
		goto L293
	}
L285:
	;
	v1966 = v1894
	v1971 = v1910
	goto L284
L286:
	;
	goto L287
L287:
	;
	v1937 = v1894
	v1939 = v1910
	goto L288
L288:
	;
	if v1937&int32(1) != 0 {
		v1966 = v1937
		v1971 = v1939
		goto L284
	} else {
		goto L290
	}
L289:
	;
	v1966 = v1937
	v1971 = v1939
	goto L284
L290:
	;
	v1960 = base.I32_div_u_s(v1937, int32(10))
	if int32(0)-v1937 == v1960*int32(-10) {
		v1937 = v1960
		v1939 = v1939 - int32(1)
		goto L288
	} else {
		goto L291
	}
L291:
	;
	goto L289
L292:
	;
	v1990 = v1966
	v1991 = v1983
	goto L295
L293:
	;
	v2036 = v1966
	v2037 = v1983
	goto L294
L294:
	;
	if base.Ui32(v2036) < base.Ui32(int32(100)) {
		goto L299
	} else {
		goto L300
	}
L295:
	;
	v2007 = v1204 + v1917 + v1971 - v1991
	v2011 = base.I32_div_u_s(v1990, int32(_a_F_sparsevec_out_10))
	v2014 = v1990 + v2011*int32(-10000)
	v2015 = int32(100)
	v2016 = base.I32_div_u_s(v2014, v2015)
	v2017 = int32(1)
	v2019 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2016<<(uint(v2017)%32))+uint32(_c_F_sparsevec_out[4]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v2007-int32(3)))) = uint16(v2019)
	v2028 = int32(*(*uint16)(unsafe.Add(mBase, uint32((v2014-v2016*v2015)<<(uint(v2017)%32))+uint32(_c_F_sparsevec_out[4]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v2007-v2017))) = uint16(v2028)
	v2031 = v1991 + int32(4)
	if base.Ui32(int32(99999999)) < base.Ui32(v1990) {
		v1990 = v2011
		v1991 = v2031
		goto L295
	} else {
		goto L297
	}
L296:
	;
	v2036 = v2011
	v2037 = v2031
	goto L294
L297:
	;
	goto L296
L298:
	;
	v2078 = v1918 - int32(1)
	v2079 = v1204 + v1917
	if base.Ui32(int32(10)) <= base.Ui32(v2076) {
		goto L303
	} else {
		goto L304
	}
L299:
	;
	v2075 = v2037
	v2076 = v2036
	goto L298
L300:
	;
	goto L301
L301:
	;
	v2060 = int32(_a_F_sparsevec_out_11)
	v2062 = int32(100)
	v2063 = base.I32_div_u_s(v2036&v2060, v2062)
	v2071 = int32(*(*uint16)(unsafe.Add(mBase, uint32((v2036-v2063*v2062)&v2060<<(uint(int32(1))%32))+uint32(_c_F_sparsevec_out[4]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v1204+v1917+v1971+(v2037^int32(-1))))) = uint16(v2071)
	v2075 = v2037 | int32(2)
	v2076 = v2063
	goto L298
L302:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v2079))) = uint8(v2093)
	if base.Ui32(int32(2)) <= base.Ui32(v1971) {
		goto L306
	} else {
		goto L307
	}
L303:
	;
	v2086 = v2076 << (uint(int32(1)) % 32)
	v2087 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2086)+uint32(_c_F_sparsevec_out[5]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1204+(v1917+v1971-v2075)))) = uint8(v2087)
	v2089 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2086)+uint32(_c_F_sparsevec_out[4]))))
	v2093 = v2089
	goto L302
L304:
	;
	goto L305
L305:
	;
	v2093 = v2076 | int32(48)
	goto L302
L306:
	;
	v2097 = int32(46)
	*(*uint8)(unsafe.Add(mBase, uint32(v2079)+1)) = uint8(v2097)
	v2102 = v1971 + int32(1)
	goto L308
L307:
	;
	v2102 = int32(1)
	goto L308
L308:
	;
	v2103 = v2102 + v1917
	v2104 = v1204 + v2103
	v2105 = int32(101)
	*(*uint8)(unsafe.Add(mBase, uint32(v2104))) = uint8(v2105)
	v2110 = base.B2i32(v2078 < int32(0))
	if v2078 < int32(0) {
		goto L309
	} else {
		goto L310
	}
L309:
	;
	v2111 = int32(45)
	goto L311
L310:
	;
	v2111 = int32(43)
	goto L311
L311:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v2104)+1)) = uint8(v2111)
	if v2078 < int32(0) {
		goto L312
	} else {
		goto L313
	}
L312:
	;
	v2115 = int32(1) - v1918
	goto L314
L313:
	;
	v2115 = v2078
	goto L314
L314:
	;
	v2120 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2115<<(uint(int32(1))%32))+uint32(_c_F_sparsevec_out[4]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v2104)+2)) = uint16(v2120)
	v2301 = v2103 + int32(4)
	goto L182
L315:
	;
	v2136 = v2129
	v2137 = v1894
	goto L318
L316:
	;
	v2182 = v2129
	v2183 = v1894
	goto L317
L317:
	;
	if base.Ui32(v2183) < base.Ui32(int32(100)) {
		goto L322
	} else {
		goto L323
	}
L318:
	;
	v2153 = v2127 + v1923 + v1910 - v2136
	v2154 = int32(4)
	v2157 = base.I32_div_u_s(v2137, int32(_a_F_sparsevec_out_10))
	v2160 = v2137 + v2157*int32(-10000)
	v2161 = int32(100)
	v2162 = base.I32_div_u_s(v2160, v2161)
	v2163 = int32(1)
	v2165 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2162<<(uint(v2163)%32))+uint32(_c_F_sparsevec_out[4]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v2153-v2154))) = uint16(v2165)
	v2174 = int32(*(*uint16)(unsafe.Add(mBase, uint32((v2160-v2162*v2161)<<(uint(v2163)%32))+uint32(_c_F_sparsevec_out[4]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v2153-int32(2)))) = uint16(v2174)
	v2177 = v2136 + v2154
	if base.Ui32(int32(99999999)) < base.Ui32(v2137) {
		v2136 = v2177
		v2137 = v2157
		goto L318
	} else {
		goto L320
	}
L319:
	;
	v2182 = v2177
	v2183 = v2157
	goto L317
L320:
	;
	goto L319
L321:
	;
	if base.Ui32(int32(10)) <= base.Ui32(v2222) {
		goto L326
	} else {
		goto L327
	}
L322:
	;
	v2221 = v2182
	v2222 = v2183
	goto L321
L323:
	;
	goto L324
L324:
	;
	v2204 = int32(2)
	v2206 = int32(_a_F_sparsevec_out_11)
	v2208 = int32(100)
	v2209 = base.I32_div_u_s(v2183&v2206, v2208)
	v2217 = int32(*(*uint16)(unsafe.Add(mBase, uint32((v2183-v2209*v2208)&v2206<<(uint(int32(1))%32))+uint32(_c_F_sparsevec_out[4]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v2127+v1923+v1910-v2182-v2204))) = uint16(v2217)
	v2221 = v2182 | v2204
	v2222 = v2209
	goto L321
L325:
	;
	v2238 = int32(1)
	if v2127 == v2238 {
		goto L330
	} else {
		goto L331
	}
L326:
	;
	v2232 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2222<<(uint(int32(1))%32))+uint32(_c_F_sparsevec_out[4]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v2127+v1923+v1910-v2221-int32(2)))) = uint16(v2232)
	goto L325
L327:
	;
	goto L328
L328:
	;
	v2236 = v2222 | int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v2127+v1923))) = uint8(v2236)
	goto L325
L329:
	;
	v2301 = v2278 + int32(base.Ui32(v1222)>>(uint(int32(31))%32))
	goto L182
L330:
	;
	if v1918&int32(4) != 0 {
		goto L333
	} else {
		goto L334
	}
L331:
	;
	goto L332
L332:
	;
	if v1899 < int32(0) {
		goto L342
	} else {
		goto L343
	}
L333:
	;
	v2243 = *(*int32)(unsafe.Add(mBase, uint32(v1923)+1))
	*(*int32)(unsafe.Add(mBase, uint32(v1923))) = v2243
	v2246 = int32(5)
	goto L335
L334:
	;
	v2246 = v2238
	goto L335
L335:
	;
	if v1918&int32(2) != 0 {
		goto L336
	} else {
		goto L337
	}
L336:
	;
	v2249 = v2246 + v1923
	v2252 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2249))))
	*(*uint16)(unsafe.Add(mBase, uint32(v2249-int32(1)))) = uint16(v2252)
	v2257 = v2246 | int32(2)
	goto L338
L337:
	;
	v2257 = v2246
	goto L338
L338:
	;
	if v1918&int32(1) != 0 {
		goto L339
	} else {
		goto L340
	}
L339:
	;
	v2260 = v2257 + v1923
	v2263 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2260))))
	*(*uint8)(unsafe.Add(mBase, uint32(v2260-int32(1)))) = uint8(v2263)
	goto L341
L340:
	;
	goto L341
L341:
	;
	v2267 = int32(46)
	*(*uint8)(unsafe.Add(mBase, uint32(v1923+v1918))) = uint8(v2267)
	v2278 = v1910 + int32(1)
	goto L329
L342:
	;
	v2275 = int32(2) - v1899
	goto L344
L343:
	;
	v2275 = v1918
	goto L344
L344:
	;
	v2278 = v2275
	goto L329
L345:
	;
	goto L177
L346:
	;
	v2338 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v2333+v2319))) = uint8(v2338)
	v2340 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v2340 != v11 {
		goto L350
	} else {
		goto L351
	}
L347:
	;
	v2329 = v2317
	v2330 = int32(0)
	goto L349
L348:
	;
	v2324 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(v2319))) = uint8(v2324)
	v2329 = int32(0) - v2317
	v2330 = int32(1)
	goto L349
L349:
	;
	v2332 = F_pg_ultoa_n(m, v2329, v2319+v2330)
	mBase = m.M
	v2333 = v2332 + v2330
	v2335 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v2319+v2333))) = uint8(v2335)
	goto L346
L350:
	;
	F_pfree(m, v11)
	mBase = m.M
	v2343 = m.ExcPending
	if v2343 != 0 {
		goto L1
	} else {
		goto L353
	}
L351:
	;
	goto L352
L352:
	;
	return v21
L353:
	;
	goto L352
}
