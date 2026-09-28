package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_lseg_closept_point(m *base.Module, l0 int32, l1 int32, l2 int32) float64 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 float64
	_ = v16
	var v19 int32
	_ = v19
	var v27 float64
	_ = v27
	var v35 float64
	_ = v35
	var v40 float64
	_ = v40
	var v41 float64
	_ = v41
	var v42 float64
	_ = v42
	var v44 float64
	_ = v44
	var v53 float64
	_ = v53
	var v54 int32
	_ = v54
	var v55 float64
	_ = v55
	var v61 float64
	_ = v61
	var v62 int32
	_ = v62
	var v63 float64
	_ = v63
	var v65 float64
	_ = v65
	var v67 float64
	_ = v67
	var v79 float64
	_ = v79
	var v80 int32
	_ = v80
	var v81 float64
	_ = v81
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v99 float64
	_ = v99
	var v100 int32
	_ = v100
	var v102 float64
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v106 int64
	_ = v106
	var v108 int64
	_ = v108
	var v111 int64
	_ = v111
	var v113 int64
	_ = v113
	var v115 float64
	_ = v115
	var v116 float64
	_ = v116
	var v117 float64
	_ = v117
	var v119 float64
	_ = v119
	var v130 float64
	_ = v130
	var v131 int32
	_ = v131
	var v132 float64
	_ = v132
	var v133 float64
	_ = v133
	var v134 float64
	_ = v134
	var v135 float64
	_ = v135
	var v137 float64
	_ = v137
	var v148 float64
	_ = v148
	var v149 int32
	_ = v149
	var v150 float64
	_ = v150
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v163 float64
	_ = v163
	var v164 float64
	_ = v164
	var v167 int32
	_ = v167
	var v168 float64
	_ = v168
	var v169 int64
	_ = v169
	var v171 int64
	_ = v171
	var v174 float64
	_ = v174
	var v177 int64
	_ = v177
	var v179 int64
	_ = v179
	var v190 float64
	_ = v190
	var v198 float64
	_ = v198
	var v203 float64
	_ = v203
	var v204 float64
	_ = v204
	var v205 float64
	_ = v205
	var v214 float64
	_ = v214
	var v215 float64
	_ = v215
	var v217 float64
	_ = v217
	var v219 float64
	_ = v219
	var v226 float64
	_ = v226
	v10 = m.G0
	v12 = v10 - int32(48)
	m.G0 = v12
	v15 = l1 + int32(16)
	v16 = F_point_invsl(m, l1, v15)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v93 = v12 + int32(8)
	v94 = F_lseg_interpt_line(m, v12+int32(32), l1, v93)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L2
	} else {
		goto L22
	}
L2:
	;
	return float64(0)
L3:
	;
	if base.F64_eq(base.F64_abs(v16), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v12)+16)) = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v12)+8)) = int64(-4616189618054758400)
	v27 = *(*float64)(unsafe.Add(mBase, uint32(l2)))
	*(*float64)(unsafe.Add(mBase, uint32(v12)+24)) = v27
	goto L1
L5:
	;
	goto L6
L6:
	;
	if base.F64_eq(v16, float64(0)) != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v12)+16)) = int64(-4616189618054758400)
	*(*int64)(unsafe.Add(mBase, uint32(v12)+8)) = int64(0)
	v35 = *(*float64)(unsafe.Add(mBase, uint32(l2)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v12)+24)) = v35
	goto L1
L8:
	;
	goto L9
L9:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v12)+16)) = int64(-4616189618054758400)
	*(*float64)(unsafe.Add(mBase, uint32(v12)+8)) = v16
	v40 = *(*float64)(unsafe.Add(mBase, uint32(l2)+8))
	v41 = *(*float64)(unsafe.Add(mBase, uint32(l2)))
	v42 = base.F64_mul(v16, v41)
	v44 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_ne(base.F64_abs(v42), v44)|base.F64_eq(base.F64_abs(v41), v44) == int32(0) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v65 = math.Float64frombits(uint64(0x7ff0000000000000))
	v67 = base.F64_sub(v40, v63)
	if base.F64_eq(base.F64_abs(v40), v65)|base.F64_ne(base.F64_abs(v67), v65)|base.F64_eq(base.F64_abs(v63), v65) == int32(0) {
		goto L17
	} else {
		goto L18
	}
L11:
	;
	v53 = F_float_overflow_error_ext(m, int32(0))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L2
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	v55 = float64(0)
	if base.F64_eq(v41, v55)|base.F64_ne(v42, v55) != 0 {
		v63 = v42
		goto L10
	} else {
		goto L15
	}
L14:
	;
	v63 = v53
	goto L10
L15:
	;
	v61 = F_float_underflow_error_ext(m, int32(0))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L2
	} else {
		goto L16
	}
L16:
	;
	v63 = v61
	goto L10
L17:
	;
	v79 = F_float_overflow_error_ext(m, int32(0))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L2
	} else {
		goto L20
	}
L18:
	;
	v81 = v67
	goto L19
L19:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v12)+24)) = v81
	if base.F64_ne(v81, float64(0)) != 0 {
		goto L1
	} else {
		goto L21
	}
L20:
	;
	v81 = v79
	goto L19
L21:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v12)+24)) = int64(0)
	goto L1
L22:
	;
	if v94 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v99 = F_line_closept_point(m, int32(0), v93, l1)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L2
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	if l0 != 0 {
		goto L31
	} else {
		goto L32
	}
L26:
	;
	v102 = F_line_closept_point(m, int32(0), v93, v15)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L2
	} else {
		goto L27
	}
L27:
	;
	if base.F64_lt(v99, v102) != 0 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v105 = l1
	goto L30
L29:
	;
	v105 = v15
	goto L30
L30:
	;
	v106 = *(*int64)(unsafe.Add(mBase, uint32(v105)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v12)+40)) = v106
	v108 = *(*int64)(unsafe.Add(mBase, uint32(v105)))
	*(*int64)(unsafe.Add(mBase, uint32(v12)+32)) = v108
	goto L25
L31:
	;
	v111 = *(*int64)(unsafe.Add(mBase, uint32(v12)+40))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v111
	v113 = *(*int64)(unsafe.Add(mBase, uint32(v12)+32))
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v113
	goto L33
L32:
	;
	goto L33
L33:
	;
	v115 = *(*float64)(unsafe.Add(mBase, uint32(v12)+32))
	v116 = *(*float64)(unsafe.Add(mBase, uint32(l2)))
	v117 = base.F64_sub(v115, v116)
	v119 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_ne(base.F64_abs(v117), v119)|base.F64_eq(base.F64_abs(v115), v119)|base.F64_eq(base.F64_abs(v116), v119) != 0 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v132 = v117
	goto L36
L35:
	;
	v130 = F_float_overflow_error_ext(m, int32(0))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L2
	} else {
		goto L37
	}
L36:
	;
	v133 = *(*float64)(unsafe.Add(mBase, uint32(v12)+40))
	v134 = *(*float64)(unsafe.Add(mBase, uint32(l2)+8))
	v135 = base.F64_sub(v133, v134)
	v137 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_ne(base.F64_abs(v135), v137)|base.F64_eq(base.F64_abs(v133), v137)|base.F64_eq(base.F64_abs(v134), v137) != 0 {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	v132 = v130
	goto L36
L38:
	;
	v150 = v135
	goto L40
L39:
	;
	v148 = F_float_overflow_error_ext(m, int32(0))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L2
	} else {
		goto L41
	}
L40:
	;
	v159 = m.G0
	v161 = v159 - int32(32)
	m.G0 = v161
	v163 = base.F64_abs(v132)
	v164 = base.F64_abs(v150)
	v167 = base.B2i32(base.Ui64(base.I64_reinterpret_f64(v163)) < base.Ui64(base.I64_reinterpret_f64(v164)))
	if base.Ui64(base.I64_reinterpret_f64(v163)) < base.Ui64(base.I64_reinterpret_f64(v164)) {
		goto L44
	} else {
		goto L45
	}
L41:
	;
	v150 = v148
	goto L40
L42:
	;
	m.G0 = v12 + int32(48)
	return v226
L43:
	;
	m.G0 = v161 + int32(32)
	goto L42
L44:
	;
	v168 = v163
	goto L46
L45:
	;
	v168 = v164
	goto L46
L46:
	;
	v169 = base.I64_reinterpret_f64(v168)
	v171 = int64(base.Ui64(v169) >> (uint(int64(52)) % 64))
	if v171 == int64(2047) {
		v226 = v168
		goto L43
	} else {
		goto L47
	}
L47:
	;
	if base.Ui64(base.I64_reinterpret_f64(v163)) < base.Ui64(base.I64_reinterpret_f64(v164)) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v174 = v164
	goto L50
L49:
	;
	v174 = v163
	goto L50
L50:
	;
	if v169 == int64(0) {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v226 = v174
	goto L43
L52:
	;
	v177 = base.I64_reinterpret_f64(v174)
	v179 = int64(base.Ui64(v177) >> (uint(int64(52)) % 64))
	if v179 == int64(2047) {
		goto L51
	} else {
		goto L53
	}
L53:
	;
	if int32(65) <= base.I32_wrap_i64(v179)-base.I32_wrap_i64(v171) {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v226 = base.F64_add(v163, v164)
	goto L43
L55:
	;
	goto L56
L56:
	;
	if base.Ui64(int64(6908521828386340864)) <= base.Ui64(v177) {
		goto L58
	} else {
		goto L59
	}
L57:
	;
	F_sq(m, v161+int32(24), v161+int32(16), v203)
	mBase = m.M
	F_sq(m, v161+int32(8), v161, v204)
	mBase = m.M
	v214 = *(*float64)(unsafe.Add(mBase, uint32(v161)))
	v215 = *(*float64)(unsafe.Add(mBase, uint32(v161)+16))
	v217 = *(*float64)(unsafe.Add(mBase, uint32(v161)+8))
	v219 = *(*float64)(unsafe.Add(mBase, uint32(v161)+24))
	v226 = base.F64_mul(v205, base.F64_sqrt(base.F64_add(base.F64_add(base.F64_add(v214, v215), v217), v219)))
	goto L43
L58:
	;
	v190 = float64(1.90109156629516e-211)
	v203 = base.F64_mul(v174, v190)
	v204 = base.F64_mul(v168, v190)
	v205 = float64(5.260135901548374e+210)
	goto L57
L59:
	;
	goto L60
L60:
	;
	if base.Ui64(int64(2580562586483294207)) < base.Ui64(v169) {
		v203 = v174
		v204 = v168
		v205 = float64(1)
		goto L57
	} else {
		goto L61
	}
L61:
	;
	v198 = float64(5.260135901548374e+210)
	v203 = base.F64_mul(v174, v198)
	v204 = base.F64_mul(v168, v198)
	v205 = float64(1.90109156629516e-211)
	goto L57
}
func F_lseg_gt(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 float64
	_ = v11
	var v12 float64
	_ = v12
	var v13 float64
	_ = v13
	var v15 float64
	_ = v15
	var v28 float64
	_ = v28
	var v31 int32
	_ = v31
	var v32 float64
	_ = v32
	var v33 float64
	_ = v33
	var v34 float64
	_ = v34
	var v35 float64
	_ = v35
	var v37 float64
	_ = v37
	var v50 float64
	_ = v50
	var v51 int32
	_ = v51
	var v52 float64
	_ = v52
	var v53 float64
	_ = v53
	var v54 float64
	_ = v54
	var v55 float64
	_ = v55
	var v57 float64
	_ = v57
	var v70 float64
	_ = v70
	var v71 int32
	_ = v71
	var v72 float64
	_ = v72
	var v73 float64
	_ = v73
	var v74 float64
	_ = v74
	var v75 float64
	_ = v75
	var v77 float64
	_ = v77
	var v90 float64
	_ = v90
	var v91 int32
	_ = v91
	var v92 float64
	_ = v92
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v105 float64
	_ = v105
	var v106 float64
	_ = v106
	var v109 int32
	_ = v109
	var v110 float64
	_ = v110
	var v111 int64
	_ = v111
	var v113 int64
	_ = v113
	var v116 float64
	_ = v116
	var v119 int64
	_ = v119
	var v121 int64
	_ = v121
	var v132 float64
	_ = v132
	var v140 float64
	_ = v140
	var v145 float64
	_ = v145
	var v146 float64
	_ = v146
	var v147 float64
	_ = v147
	var v156 float64
	_ = v156
	var v157 float64
	_ = v157
	var v159 float64
	_ = v159
	var v161 float64
	_ = v161
	var v168 float64
	_ = v168
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v186 float64
	_ = v186
	var v187 float64
	_ = v187
	var v190 int32
	_ = v190
	var v191 float64
	_ = v191
	var v192 int64
	_ = v192
	var v194 int64
	_ = v194
	var v197 float64
	_ = v197
	var v200 int64
	_ = v200
	var v202 int64
	_ = v202
	var v213 float64
	_ = v213
	var v221 float64
	_ = v221
	var v226 float64
	_ = v226
	var v227 float64
	_ = v227
	var v228 float64
	_ = v228
	var v237 float64
	_ = v237
	var v238 float64
	_ = v238
	var v240 float64
	_ = v240
	var v242 float64
	_ = v242
	var v249 float64
	_ = v249
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v11 = *(*float64)(unsafe.Add(mBase, uint32(v10)))
	v12 = *(*float64)(unsafe.Add(mBase, uint32(v10)+16))
	v13 = base.F64_sub(v11, v12)
	v15 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_ne(base.F64_abs(v13), v15)|base.F64_eq(base.F64_abs(v11), v15)|base.F64_eq(base.F64_abs(v12), v15) == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v28 = F_float_overflow_error_ext(m, int32(0))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v32 = v13
	goto L3
L3:
	;
	v33 = *(*float64)(unsafe.Add(mBase, uint32(v10)+8))
	v34 = *(*float64)(unsafe.Add(mBase, uint32(v10)+24))
	v35 = base.F64_sub(v33, v34)
	v37 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_ne(base.F64_abs(v35), v37)|base.F64_eq(base.F64_abs(v33), v37)|base.F64_eq(base.F64_abs(v34), v37) == int32(0) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	return int64(0)
L5:
	;
	v32 = v28
	goto L3
L6:
	;
	v50 = F_float_overflow_error_ext(m, int32(0))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L4
	} else {
		goto L9
	}
L7:
	;
	v52 = v35
	goto L8
L8:
	;
	v53 = *(*float64)(unsafe.Add(mBase, uint32(v9)))
	v54 = *(*float64)(unsafe.Add(mBase, uint32(v9)+16))
	v55 = base.F64_sub(v53, v54)
	v57 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_ne(base.F64_abs(v55), v57)|base.F64_eq(base.F64_abs(v53), v57)|base.F64_eq(base.F64_abs(v54), v57) == int32(0) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v52 = v50
	goto L8
L10:
	;
	v70 = F_float_overflow_error_ext(m, int32(0))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L4
	} else {
		goto L13
	}
L11:
	;
	v72 = v55
	goto L12
L12:
	;
	v73 = *(*float64)(unsafe.Add(mBase, uint32(v9)+8))
	v74 = *(*float64)(unsafe.Add(mBase, uint32(v9)+24))
	v75 = base.F64_sub(v73, v74)
	v77 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_ne(base.F64_abs(v75), v77)|base.F64_eq(base.F64_abs(v73), v77)|base.F64_eq(base.F64_abs(v74), v77) == int32(0) {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	v72 = v70
	goto L12
L14:
	;
	v90 = F_float_overflow_error_ext(m, int32(0))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L4
	} else {
		goto L17
	}
L15:
	;
	v92 = v75
	goto L16
L16:
	;
	v101 = m.G0
	v103 = v101 - int32(32)
	m.G0 = v103
	v105 = base.F64_abs(v32)
	v106 = base.F64_abs(v52)
	v109 = base.B2i32(base.Ui64(base.I64_reinterpret_f64(v105)) < base.Ui64(base.I64_reinterpret_f64(v106)))
	if base.Ui64(base.I64_reinterpret_f64(v105)) < base.Ui64(base.I64_reinterpret_f64(v106)) {
		goto L20
	} else {
		goto L21
	}
L17:
	;
	v92 = v90
	goto L16
L18:
	;
	v182 = m.G0
	v184 = v182 - int32(32)
	m.G0 = v184
	v186 = base.F64_abs(v72)
	v187 = base.F64_abs(v92)
	v190 = base.B2i32(base.Ui64(base.I64_reinterpret_f64(v186)) < base.Ui64(base.I64_reinterpret_f64(v187)))
	if base.Ui64(base.I64_reinterpret_f64(v186)) < base.Ui64(base.I64_reinterpret_f64(v187)) {
		goto L40
	} else {
		goto L41
	}
L19:
	;
	m.G0 = v103 + int32(32)
	goto L18
L20:
	;
	v110 = v105
	goto L22
L21:
	;
	v110 = v106
	goto L22
L22:
	;
	v111 = base.I64_reinterpret_f64(v110)
	v113 = int64(base.Ui64(v111) >> (uint(int64(52)) % 64))
	if v113 == int64(2047) {
		v168 = v110
		goto L19
	} else {
		goto L23
	}
L23:
	;
	if base.Ui64(base.I64_reinterpret_f64(v105)) < base.Ui64(base.I64_reinterpret_f64(v106)) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v116 = v106
	goto L26
L25:
	;
	v116 = v105
	goto L26
L26:
	;
	if v111 == int64(0) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v168 = v116
	goto L19
L28:
	;
	v119 = base.I64_reinterpret_f64(v116)
	v121 = int64(base.Ui64(v119) >> (uint(int64(52)) % 64))
	if v121 == int64(2047) {
		goto L27
	} else {
		goto L29
	}
L29:
	;
	if int32(65) <= base.I32_wrap_i64(v121)-base.I32_wrap_i64(v113) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v168 = base.F64_add(v105, v106)
	goto L19
L31:
	;
	goto L32
L32:
	;
	if base.Ui64(int64(6908521828386340864)) <= base.Ui64(v119) {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	F_sq(m, v103+int32(24), v103+int32(16), v145)
	mBase = m.M
	F_sq(m, v103+int32(8), v103, v146)
	mBase = m.M
	v156 = *(*float64)(unsafe.Add(mBase, uint32(v103)))
	v157 = *(*float64)(unsafe.Add(mBase, uint32(v103)+16))
	v159 = *(*float64)(unsafe.Add(mBase, uint32(v103)+8))
	v161 = *(*float64)(unsafe.Add(mBase, uint32(v103)+24))
	v168 = base.F64_mul(v147, base.F64_sqrt(base.F64_add(base.F64_add(base.F64_add(v156, v157), v159), v161)))
	goto L19
L34:
	;
	v132 = float64(1.90109156629516e-211)
	v145 = base.F64_mul(v116, v132)
	v146 = base.F64_mul(v110, v132)
	v147 = float64(5.260135901548374e+210)
	goto L33
L35:
	;
	goto L36
L36:
	;
	if base.Ui64(int64(2580562586483294207)) < base.Ui64(v111) {
		v145 = v116
		v146 = v110
		v147 = float64(1)
		goto L33
	} else {
		goto L37
	}
L37:
	;
	v140 = float64(5.260135901548374e+210)
	v145 = base.F64_mul(v116, v140)
	v146 = base.F64_mul(v110, v140)
	v147 = float64(1.90109156629516e-211)
	goto L33
L38:
	;
	return base.I64_extend_i32_u(base.F64_gt(v168, base.F64_add(v249, float64(1e-06))))
L39:
	;
	m.G0 = v184 + int32(32)
	goto L38
L40:
	;
	v191 = v186
	goto L42
L41:
	;
	v191 = v187
	goto L42
L42:
	;
	v192 = base.I64_reinterpret_f64(v191)
	v194 = int64(base.Ui64(v192) >> (uint(int64(52)) % 64))
	if v194 == int64(2047) {
		v249 = v191
		goto L39
	} else {
		goto L43
	}
L43:
	;
	if base.Ui64(base.I64_reinterpret_f64(v186)) < base.Ui64(base.I64_reinterpret_f64(v187)) {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v197 = v187
	goto L46
L45:
	;
	v197 = v186
	goto L46
L46:
	;
	if v192 == int64(0) {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v249 = v197
	goto L39
L48:
	;
	v200 = base.I64_reinterpret_f64(v197)
	v202 = int64(base.Ui64(v200) >> (uint(int64(52)) % 64))
	if v202 == int64(2047) {
		goto L47
	} else {
		goto L49
	}
L49:
	;
	if int32(65) <= base.I32_wrap_i64(v202)-base.I32_wrap_i64(v194) {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v249 = base.F64_add(v186, v187)
	goto L39
L51:
	;
	goto L52
L52:
	;
	if base.Ui64(int64(6908521828386340864)) <= base.Ui64(v200) {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	F_sq(m, v184+int32(24), v184+int32(16), v226)
	mBase = m.M
	F_sq(m, v184+int32(8), v184, v227)
	mBase = m.M
	v237 = *(*float64)(unsafe.Add(mBase, uint32(v184)))
	v238 = *(*float64)(unsafe.Add(mBase, uint32(v184)+16))
	v240 = *(*float64)(unsafe.Add(mBase, uint32(v184)+8))
	v242 = *(*float64)(unsafe.Add(mBase, uint32(v184)+24))
	v249 = base.F64_mul(v228, base.F64_sqrt(base.F64_add(base.F64_add(base.F64_add(v237, v238), v240), v242)))
	goto L39
L54:
	;
	v213 = float64(1.90109156629516e-211)
	v226 = base.F64_mul(v197, v213)
	v227 = base.F64_mul(v191, v213)
	v228 = float64(5.260135901548374e+210)
	goto L53
L55:
	;
	goto L56
L56:
	;
	if base.Ui64(int64(2580562586483294207)) < base.Ui64(v192) {
		v226 = v197
		v227 = v191
		v228 = float64(1)
		goto L53
	} else {
		goto L57
	}
L57:
	;
	v221 = float64(5.260135901548374e+210)
	v226 = base.F64_mul(v197, v221)
	v227 = base.F64_mul(v191, v221)
	v228 = float64(1.90109156629516e-211)
	goto L53
}
func F_lseg_intersect(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v5 = F_lseg_interpt_lseg(m, int32(0), v3, v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int64(0)
	} else {
		return base.I64_extend_i32_u(v5)
	}
}
func F_lseg_le(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 float64
	_ = v11
	var v12 float64
	_ = v12
	var v13 float64
	_ = v13
	var v15 float64
	_ = v15
	var v28 float64
	_ = v28
	var v31 int32
	_ = v31
	var v32 float64
	_ = v32
	var v33 float64
	_ = v33
	var v34 float64
	_ = v34
	var v35 float64
	_ = v35
	var v37 float64
	_ = v37
	var v50 float64
	_ = v50
	var v51 int32
	_ = v51
	var v52 float64
	_ = v52
	var v53 float64
	_ = v53
	var v54 float64
	_ = v54
	var v55 float64
	_ = v55
	var v57 float64
	_ = v57
	var v70 float64
	_ = v70
	var v71 int32
	_ = v71
	var v72 float64
	_ = v72
	var v73 float64
	_ = v73
	var v74 float64
	_ = v74
	var v75 float64
	_ = v75
	var v77 float64
	_ = v77
	var v90 float64
	_ = v90
	var v91 int32
	_ = v91
	var v92 float64
	_ = v92
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v105 float64
	_ = v105
	var v106 float64
	_ = v106
	var v109 int32
	_ = v109
	var v110 float64
	_ = v110
	var v111 int64
	_ = v111
	var v113 int64
	_ = v113
	var v116 float64
	_ = v116
	var v119 int64
	_ = v119
	var v121 int64
	_ = v121
	var v132 float64
	_ = v132
	var v140 float64
	_ = v140
	var v145 float64
	_ = v145
	var v146 float64
	_ = v146
	var v147 float64
	_ = v147
	var v156 float64
	_ = v156
	var v157 float64
	_ = v157
	var v159 float64
	_ = v159
	var v161 float64
	_ = v161
	var v168 float64
	_ = v168
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v186 float64
	_ = v186
	var v187 float64
	_ = v187
	var v190 int32
	_ = v190
	var v191 float64
	_ = v191
	var v192 int64
	_ = v192
	var v194 int64
	_ = v194
	var v197 float64
	_ = v197
	var v200 int64
	_ = v200
	var v202 int64
	_ = v202
	var v213 float64
	_ = v213
	var v221 float64
	_ = v221
	var v226 float64
	_ = v226
	var v227 float64
	_ = v227
	var v228 float64
	_ = v228
	var v237 float64
	_ = v237
	var v238 float64
	_ = v238
	var v240 float64
	_ = v240
	var v242 float64
	_ = v242
	var v249 float64
	_ = v249
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v11 = *(*float64)(unsafe.Add(mBase, uint32(v10)))
	v12 = *(*float64)(unsafe.Add(mBase, uint32(v10)+16))
	v13 = base.F64_sub(v11, v12)
	v15 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_ne(base.F64_abs(v13), v15)|base.F64_eq(base.F64_abs(v11), v15)|base.F64_eq(base.F64_abs(v12), v15) == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v28 = F_float_overflow_error_ext(m, int32(0))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v32 = v13
	goto L3
L3:
	;
	v33 = *(*float64)(unsafe.Add(mBase, uint32(v10)+8))
	v34 = *(*float64)(unsafe.Add(mBase, uint32(v10)+24))
	v35 = base.F64_sub(v33, v34)
	v37 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_ne(base.F64_abs(v35), v37)|base.F64_eq(base.F64_abs(v33), v37)|base.F64_eq(base.F64_abs(v34), v37) == int32(0) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	return int64(0)
L5:
	;
	v32 = v28
	goto L3
L6:
	;
	v50 = F_float_overflow_error_ext(m, int32(0))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L4
	} else {
		goto L9
	}
L7:
	;
	v52 = v35
	goto L8
L8:
	;
	v53 = *(*float64)(unsafe.Add(mBase, uint32(v9)))
	v54 = *(*float64)(unsafe.Add(mBase, uint32(v9)+16))
	v55 = base.F64_sub(v53, v54)
	v57 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_ne(base.F64_abs(v55), v57)|base.F64_eq(base.F64_abs(v53), v57)|base.F64_eq(base.F64_abs(v54), v57) == int32(0) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v52 = v50
	goto L8
L10:
	;
	v70 = F_float_overflow_error_ext(m, int32(0))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L4
	} else {
		goto L13
	}
L11:
	;
	v72 = v55
	goto L12
L12:
	;
	v73 = *(*float64)(unsafe.Add(mBase, uint32(v9)+8))
	v74 = *(*float64)(unsafe.Add(mBase, uint32(v9)+24))
	v75 = base.F64_sub(v73, v74)
	v77 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_ne(base.F64_abs(v75), v77)|base.F64_eq(base.F64_abs(v73), v77)|base.F64_eq(base.F64_abs(v74), v77) == int32(0) {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	v72 = v70
	goto L12
L14:
	;
	v90 = F_float_overflow_error_ext(m, int32(0))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L4
	} else {
		goto L17
	}
L15:
	;
	v92 = v75
	goto L16
L16:
	;
	v101 = m.G0
	v103 = v101 - int32(32)
	m.G0 = v103
	v105 = base.F64_abs(v32)
	v106 = base.F64_abs(v52)
	v109 = base.B2i32(base.Ui64(base.I64_reinterpret_f64(v105)) < base.Ui64(base.I64_reinterpret_f64(v106)))
	if base.Ui64(base.I64_reinterpret_f64(v105)) < base.Ui64(base.I64_reinterpret_f64(v106)) {
		goto L20
	} else {
		goto L21
	}
L17:
	;
	v92 = v90
	goto L16
L18:
	;
	v182 = m.G0
	v184 = v182 - int32(32)
	m.G0 = v184
	v186 = base.F64_abs(v72)
	v187 = base.F64_abs(v92)
	v190 = base.B2i32(base.Ui64(base.I64_reinterpret_f64(v186)) < base.Ui64(base.I64_reinterpret_f64(v187)))
	if base.Ui64(base.I64_reinterpret_f64(v186)) < base.Ui64(base.I64_reinterpret_f64(v187)) {
		goto L40
	} else {
		goto L41
	}
L19:
	;
	m.G0 = v103 + int32(32)
	goto L18
L20:
	;
	v110 = v105
	goto L22
L21:
	;
	v110 = v106
	goto L22
L22:
	;
	v111 = base.I64_reinterpret_f64(v110)
	v113 = int64(base.Ui64(v111) >> (uint(int64(52)) % 64))
	if v113 == int64(2047) {
		v168 = v110
		goto L19
	} else {
		goto L23
	}
L23:
	;
	if base.Ui64(base.I64_reinterpret_f64(v105)) < base.Ui64(base.I64_reinterpret_f64(v106)) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v116 = v106
	goto L26
L25:
	;
	v116 = v105
	goto L26
L26:
	;
	if v111 == int64(0) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v168 = v116
	goto L19
L28:
	;
	v119 = base.I64_reinterpret_f64(v116)
	v121 = int64(base.Ui64(v119) >> (uint(int64(52)) % 64))
	if v121 == int64(2047) {
		goto L27
	} else {
		goto L29
	}
L29:
	;
	if int32(65) <= base.I32_wrap_i64(v121)-base.I32_wrap_i64(v113) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v168 = base.F64_add(v105, v106)
	goto L19
L31:
	;
	goto L32
L32:
	;
	if base.Ui64(int64(6908521828386340864)) <= base.Ui64(v119) {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	F_sq(m, v103+int32(24), v103+int32(16), v145)
	mBase = m.M
	F_sq(m, v103+int32(8), v103, v146)
	mBase = m.M
	v156 = *(*float64)(unsafe.Add(mBase, uint32(v103)))
	v157 = *(*float64)(unsafe.Add(mBase, uint32(v103)+16))
	v159 = *(*float64)(unsafe.Add(mBase, uint32(v103)+8))
	v161 = *(*float64)(unsafe.Add(mBase, uint32(v103)+24))
	v168 = base.F64_mul(v147, base.F64_sqrt(base.F64_add(base.F64_add(base.F64_add(v156, v157), v159), v161)))
	goto L19
L34:
	;
	v132 = float64(1.90109156629516e-211)
	v145 = base.F64_mul(v116, v132)
	v146 = base.F64_mul(v110, v132)
	v147 = float64(5.260135901548374e+210)
	goto L33
L35:
	;
	goto L36
L36:
	;
	if base.Ui64(int64(2580562586483294207)) < base.Ui64(v111) {
		v145 = v116
		v146 = v110
		v147 = float64(1)
		goto L33
	} else {
		goto L37
	}
L37:
	;
	v140 = float64(5.260135901548374e+210)
	v145 = base.F64_mul(v116, v140)
	v146 = base.F64_mul(v110, v140)
	v147 = float64(1.90109156629516e-211)
	goto L33
L38:
	;
	return base.I64_extend_i32_u(base.F64_le(v168, base.F64_add(v249, float64(1e-06))))
L39:
	;
	m.G0 = v184 + int32(32)
	goto L38
L40:
	;
	v191 = v186
	goto L42
L41:
	;
	v191 = v187
	goto L42
L42:
	;
	v192 = base.I64_reinterpret_f64(v191)
	v194 = int64(base.Ui64(v192) >> (uint(int64(52)) % 64))
	if v194 == int64(2047) {
		v249 = v191
		goto L39
	} else {
		goto L43
	}
L43:
	;
	if base.Ui64(base.I64_reinterpret_f64(v186)) < base.Ui64(base.I64_reinterpret_f64(v187)) {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v197 = v187
	goto L46
L45:
	;
	v197 = v186
	goto L46
L46:
	;
	if v192 == int64(0) {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v249 = v197
	goto L39
L48:
	;
	v200 = base.I64_reinterpret_f64(v197)
	v202 = int64(base.Ui64(v200) >> (uint(int64(52)) % 64))
	if v202 == int64(2047) {
		goto L47
	} else {
		goto L49
	}
L49:
	;
	if int32(65) <= base.I32_wrap_i64(v202)-base.I32_wrap_i64(v194) {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v249 = base.F64_add(v186, v187)
	goto L39
L51:
	;
	goto L52
L52:
	;
	if base.Ui64(int64(6908521828386340864)) <= base.Ui64(v200) {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	F_sq(m, v184+int32(24), v184+int32(16), v226)
	mBase = m.M
	F_sq(m, v184+int32(8), v184, v227)
	mBase = m.M
	v237 = *(*float64)(unsafe.Add(mBase, uint32(v184)))
	v238 = *(*float64)(unsafe.Add(mBase, uint32(v184)+16))
	v240 = *(*float64)(unsafe.Add(mBase, uint32(v184)+8))
	v242 = *(*float64)(unsafe.Add(mBase, uint32(v184)+24))
	v249 = base.F64_mul(v228, base.F64_sqrt(base.F64_add(base.F64_add(base.F64_add(v237, v238), v240), v242)))
	goto L39
L54:
	;
	v213 = float64(1.90109156629516e-211)
	v226 = base.F64_mul(v197, v213)
	v227 = base.F64_mul(v191, v213)
	v228 = float64(5.260135901548374e+210)
	goto L53
L55:
	;
	goto L56
L56:
	;
	if base.Ui64(int64(2580562586483294207)) < base.Ui64(v192) {
		v226 = v197
		v227 = v191
		v228 = float64(1)
		goto L53
	} else {
		goto L57
	}
L57:
	;
	v221 = float64(5.260135901548374e+210)
	v226 = base.F64_mul(v197, v221)
	v227 = base.F64_mul(v191, v221)
	v228 = float64(1.90109156629516e-211)
	goto L53
}
func F_lseg_ne(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 float64
	_ = v13
	var v19 float64
	_ = v19
	var v21 int64
	_ = v21
	var v22 int64
	_ = v22
	var v23 float64
	_ = v23
	var v26 int64
	_ = v26
	var v31 int64
	_ = v31
	var v34 float64
	_ = v34
	var v41 int64
	_ = v41
	var v48 float64
	_ = v48
	var v69 int32
	_ = v69
	var v75 float64
	_ = v75
	var v80 int64
	_ = v80
	var v82 float64
	_ = v82
	var v85 int64
	_ = v85
	var v88 int64
	_ = v88
	var v104 float64
	_ = v104
	var v110 float64
	_ = v110
	var v112 int64
	_ = v112
	var v113 int64
	_ = v113
	var v114 float64
	_ = v114
	var v117 int64
	_ = v117
	var v122 int32
	_ = v122
	var v125 float64
	_ = v125
	var v133 int64
	_ = v133
	var v138 float64
	_ = v138
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v165 float64
	_ = v165
	var v169 int64
	_ = v169
	var v171 float64
	_ = v171
	var v174 int64
	_ = v174
	var v190 int32
	_ = v190
	var v200 int64
	_ = v200
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v13 = *(*float64)(unsafe.Add(mBase, uint32(v12)))
	if base.Ui64(base.I64_reinterpret_f64(v13)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
		v19 = *(*float64)(unsafe.Add(mBase, uint32(v11)))
		v21 = int64(9223372036854775807)
		v22 = base.I64_reinterpret_f64(v19) & v21
		v23 = *(*float64)(unsafe.Add(mBase, uint32(v12)+8))
		v26 = base.I64_reinterpret_f64(v23) & v21
		if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v26) {
			v69 = base.B2i32(base.Ui64(v22) < base.Ui64(int64(9218868437227405313)))
			if base.B2i32(v69 == int32(0))|base.F64_ne(v13, v19) != 0 {
				v200 = int64(1)
				return v200
			} else {
				v75 = v23
				v80 = v26
				v82 = *(*float64)(unsafe.Add(mBase, uint32(v11)+8))
				v85 = base.I64_reinterpret_f64(v82) & int64(9223372036854775807)
				if base.Ui64(v80) <= base.Ui64(int64(9218868437227405312)) {
					v88 = int64(1)
					if base.F64_ne(v82, v75) != 0 {
						v200 = v88
					} else {
						if base.Ui64(v85) < base.Ui64(int64(9218868437227405313)) {
							v104 = *(*float64)(unsafe.Add(mBase, uint32(v12)+16))
							if base.Ui64(base.I64_reinterpret_f64(v104)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
								v110 = *(*float64)(unsafe.Add(mBase, uint32(v11)+16))
								v112 = int64(9223372036854775807)
								v113 = base.I64_reinterpret_f64(v110) & v112
								v114 = *(*float64)(unsafe.Add(mBase, uint32(v12)+24))
								v117 = base.I64_reinterpret_f64(v114) & v112
								if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v117) {
									v156 = base.B2i32(base.Ui64(v113) < base.Ui64(int64(9218868437227405313)))
									v159 = int32(0)
									if base.B2i32(v156 == v159)|base.F64_ne(v104, v110) != 0 {
										v190 = v159
									} else {
										v165 = v114
										v169 = v117
										v171 = *(*float64)(unsafe.Add(mBase, uint32(v11)+24))
										v174 = base.I64_reinterpret_f64(v171) & int64(9223372036854775807)
										if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v169) {
											v190 = base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v174))
										} else {
											v190 = base.B2i32(base.Ui64(v174) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v171, v165)
										}
									}
								} else {
									v122 = int32(0)
									if base.Ui64(int64(9218868437227405312)) < base.Ui64(v113) {
										v190 = v122
									} else {
										v125 = *(*float64)(unsafe.Add(mBase, uint32(v11)+24))
										if base.Ui64(base.I64_reinterpret_f64(v125)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
											if base.B2i32(base.F64_le(base.F64_abs(base.F64_sub(v104, v110)), float64(1e-06)) == int32(0))&base.F64_ne(v104, v110) != 0 {
												v190 = v122
											} else {
												v190 = base.F64_eq(v114, v125) | base.F64_le(base.F64_abs(base.F64_sub(v114, v125)), float64(1e-06))
											}
										} else {
											v156 = int32(1)
											v159 = int32(0)
											if base.B2i32(v156 == v159)|base.F64_ne(v104, v110) != 0 {
												v190 = v159
											} else {
												v165 = v114
												v169 = v117
												v171 = *(*float64)(unsafe.Add(mBase, uint32(v11)+24))
												v174 = base.I64_reinterpret_f64(v171) & int64(9223372036854775807)
												if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v169) {
													v190 = base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v174))
												} else {
													v190 = base.B2i32(base.Ui64(v174) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v171, v165)
												}
											}
										}
									}
								}
							} else {
								v133 = *(*int64)(unsafe.Add(mBase, uint32(v11)+16))
								if base.Ui64(v133&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
									v190 = int32(0)
								} else {
									v138 = *(*float64)(unsafe.Add(mBase, uint32(v12)+24))
									v165 = v138
									v169 = base.I64_reinterpret_f64(v138) & int64(9223372036854775807)
									v171 = *(*float64)(unsafe.Add(mBase, uint32(v11)+24))
									v174 = base.I64_reinterpret_f64(v171) & int64(9223372036854775807)
									if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v169) {
										v190 = base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v174))
									} else {
										v190 = base.B2i32(base.Ui64(v174) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v171, v165)
									}
								}
							}
							v200 = base.I64_extend_i32_u(base.B2i32(v190 == int32(0)))
						} else {
							v200 = v88
						}
					}
					return v200
				} else {
					if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v85) {
						v104 = *(*float64)(unsafe.Add(mBase, uint32(v12)+16))
						if base.Ui64(base.I64_reinterpret_f64(v104)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
							v110 = *(*float64)(unsafe.Add(mBase, uint32(v11)+16))
							v112 = int64(9223372036854775807)
							v113 = base.I64_reinterpret_f64(v110) & v112
							v114 = *(*float64)(unsafe.Add(mBase, uint32(v12)+24))
							v117 = base.I64_reinterpret_f64(v114) & v112
							if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v117) {
								v156 = base.B2i32(base.Ui64(v113) < base.Ui64(int64(9218868437227405313)))
								v159 = int32(0)
								if base.B2i32(v156 == v159)|base.F64_ne(v104, v110) != 0 {
									v190 = v159
								} else {
									v165 = v114
									v169 = v117
									v171 = *(*float64)(unsafe.Add(mBase, uint32(v11)+24))
									v174 = base.I64_reinterpret_f64(v171) & int64(9223372036854775807)
									if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v169) {
										v190 = base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v174))
									} else {
										v190 = base.B2i32(base.Ui64(v174) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v171, v165)
									}
								}
							} else {
								v122 = int32(0)
								if base.Ui64(int64(9218868437227405312)) < base.Ui64(v113) {
									v190 = v122
								} else {
									v125 = *(*float64)(unsafe.Add(mBase, uint32(v11)+24))
									if base.Ui64(base.I64_reinterpret_f64(v125)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
										if base.B2i32(base.F64_le(base.F64_abs(base.F64_sub(v104, v110)), float64(1e-06)) == int32(0))&base.F64_ne(v104, v110) != 0 {
											v190 = v122
										} else {
											v190 = base.F64_eq(v114, v125) | base.F64_le(base.F64_abs(base.F64_sub(v114, v125)), float64(1e-06))
										}
									} else {
										v156 = int32(1)
										v159 = int32(0)
										if base.B2i32(v156 == v159)|base.F64_ne(v104, v110) != 0 {
											v190 = v159
										} else {
											v165 = v114
											v169 = v117
											v171 = *(*float64)(unsafe.Add(mBase, uint32(v11)+24))
											v174 = base.I64_reinterpret_f64(v171) & int64(9223372036854775807)
											if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v169) {
												v190 = base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v174))
											} else {
												v190 = base.B2i32(base.Ui64(v174) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v171, v165)
											}
										}
									}
								}
							}
						} else {
							v133 = *(*int64)(unsafe.Add(mBase, uint32(v11)+16))
							if base.Ui64(v133&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
								v190 = int32(0)
							} else {
								v138 = *(*float64)(unsafe.Add(mBase, uint32(v12)+24))
								v165 = v138
								v169 = base.I64_reinterpret_f64(v138) & int64(9223372036854775807)
								v171 = *(*float64)(unsafe.Add(mBase, uint32(v11)+24))
								v174 = base.I64_reinterpret_f64(v171) & int64(9223372036854775807)
								if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v169) {
									v190 = base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v174))
								} else {
									v190 = base.B2i32(base.Ui64(v174) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v171, v165)
								}
							}
						}
						v200 = base.I64_extend_i32_u(base.B2i32(v190 == int32(0)))
						return v200
					} else {
						return int64(1)
					}
				}
			}
		} else {
			v31 = int64(1)
			if base.Ui64(int64(9218868437227405312)) < base.Ui64(v22) {
				v200 = v31
				return v200
			} else {
				v34 = *(*float64)(unsafe.Add(mBase, uint32(v11)+8))
				if base.Ui64(base.I64_reinterpret_f64(v34)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
					if base.B2i32(base.F64_le(base.F64_abs(base.F64_sub(v13, v19)), float64(1e-06)) == int32(0))&base.F64_ne(v13, v19) != 0 {
						v200 = v31
					} else {
						if base.F64_eq(v23, v34) != 0 {
							v104 = *(*float64)(unsafe.Add(mBase, uint32(v12)+16))
							if base.Ui64(base.I64_reinterpret_f64(v104)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
								v110 = *(*float64)(unsafe.Add(mBase, uint32(v11)+16))
								v112 = int64(9223372036854775807)
								v113 = base.I64_reinterpret_f64(v110) & v112
								v114 = *(*float64)(unsafe.Add(mBase, uint32(v12)+24))
								v117 = base.I64_reinterpret_f64(v114) & v112
								if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v117) {
									v156 = base.B2i32(base.Ui64(v113) < base.Ui64(int64(9218868437227405313)))
									v159 = int32(0)
									if base.B2i32(v156 == v159)|base.F64_ne(v104, v110) != 0 {
										v190 = v159
									} else {
										v165 = v114
										v169 = v117
										v171 = *(*float64)(unsafe.Add(mBase, uint32(v11)+24))
										v174 = base.I64_reinterpret_f64(v171) & int64(9223372036854775807)
										if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v169) {
											v190 = base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v174))
										} else {
											v190 = base.B2i32(base.Ui64(v174) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v171, v165)
										}
									}
								} else {
									v122 = int32(0)
									if base.Ui64(int64(9218868437227405312)) < base.Ui64(v113) {
										v190 = v122
									} else {
										v125 = *(*float64)(unsafe.Add(mBase, uint32(v11)+24))
										if base.Ui64(base.I64_reinterpret_f64(v125)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
											if base.B2i32(base.F64_le(base.F64_abs(base.F64_sub(v104, v110)), float64(1e-06)) == int32(0))&base.F64_ne(v104, v110) != 0 {
												v190 = v122
											} else {
												v190 = base.F64_eq(v114, v125) | base.F64_le(base.F64_abs(base.F64_sub(v114, v125)), float64(1e-06))
											}
										} else {
											v156 = int32(1)
											v159 = int32(0)
											if base.B2i32(v156 == v159)|base.F64_ne(v104, v110) != 0 {
												v190 = v159
											} else {
												v165 = v114
												v169 = v117
												v171 = *(*float64)(unsafe.Add(mBase, uint32(v11)+24))
												v174 = base.I64_reinterpret_f64(v171) & int64(9223372036854775807)
												if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v169) {
													v190 = base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v174))
												} else {
													v190 = base.B2i32(base.Ui64(v174) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v171, v165)
												}
											}
										}
									}
								}
							} else {
								v133 = *(*int64)(unsafe.Add(mBase, uint32(v11)+16))
								if base.Ui64(v133&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
									v190 = int32(0)
								} else {
									v138 = *(*float64)(unsafe.Add(mBase, uint32(v12)+24))
									v165 = v138
									v169 = base.I64_reinterpret_f64(v138) & int64(9223372036854775807)
									v171 = *(*float64)(unsafe.Add(mBase, uint32(v11)+24))
									v174 = base.I64_reinterpret_f64(v171) & int64(9223372036854775807)
									if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v169) {
										v190 = base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v174))
									} else {
										v190 = base.B2i32(base.Ui64(v174) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v171, v165)
									}
								}
							}
							v200 = base.I64_extend_i32_u(base.B2i32(v190 == int32(0)))
						} else {
							if base.F64_le(base.F64_abs(base.F64_sub(v23, v34)), float64(1e-06)) == int32(0) {
								v200 = v31
							} else {
								v104 = *(*float64)(unsafe.Add(mBase, uint32(v12)+16))
								if base.Ui64(base.I64_reinterpret_f64(v104)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
									v110 = *(*float64)(unsafe.Add(mBase, uint32(v11)+16))
									v112 = int64(9223372036854775807)
									v113 = base.I64_reinterpret_f64(v110) & v112
									v114 = *(*float64)(unsafe.Add(mBase, uint32(v12)+24))
									v117 = base.I64_reinterpret_f64(v114) & v112
									if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v117) {
										v156 = base.B2i32(base.Ui64(v113) < base.Ui64(int64(9218868437227405313)))
										v159 = int32(0)
										if base.B2i32(v156 == v159)|base.F64_ne(v104, v110) != 0 {
											v190 = v159
										} else {
											v165 = v114
											v169 = v117
											v171 = *(*float64)(unsafe.Add(mBase, uint32(v11)+24))
											v174 = base.I64_reinterpret_f64(v171) & int64(9223372036854775807)
											if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v169) {
												v190 = base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v174))
											} else {
												v190 = base.B2i32(base.Ui64(v174) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v171, v165)
											}
										}
									} else {
										v122 = int32(0)
										if base.Ui64(int64(9218868437227405312)) < base.Ui64(v113) {
											v190 = v122
										} else {
											v125 = *(*float64)(unsafe.Add(mBase, uint32(v11)+24))
											if base.Ui64(base.I64_reinterpret_f64(v125)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
												if base.B2i32(base.F64_le(base.F64_abs(base.F64_sub(v104, v110)), float64(1e-06)) == int32(0))&base.F64_ne(v104, v110) != 0 {
													v190 = v122
												} else {
													v190 = base.F64_eq(v114, v125) | base.F64_le(base.F64_abs(base.F64_sub(v114, v125)), float64(1e-06))
												}
											} else {
												v156 = int32(1)
												v159 = int32(0)
												if base.B2i32(v156 == v159)|base.F64_ne(v104, v110) != 0 {
													v190 = v159
												} else {
													v165 = v114
													v169 = v117
													v171 = *(*float64)(unsafe.Add(mBase, uint32(v11)+24))
													v174 = base.I64_reinterpret_f64(v171) & int64(9223372036854775807)
													if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v169) {
														v190 = base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v174))
													} else {
														v190 = base.B2i32(base.Ui64(v174) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v171, v165)
													}
												}
											}
										}
									}
								} else {
									v133 = *(*int64)(unsafe.Add(mBase, uint32(v11)+16))
									if base.Ui64(v133&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
										v190 = int32(0)
									} else {
										v138 = *(*float64)(unsafe.Add(mBase, uint32(v12)+24))
										v165 = v138
										v169 = base.I64_reinterpret_f64(v138) & int64(9223372036854775807)
										v171 = *(*float64)(unsafe.Add(mBase, uint32(v11)+24))
										v174 = base.I64_reinterpret_f64(v171) & int64(9223372036854775807)
										if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v169) {
											v190 = base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v174))
										} else {
											v190 = base.B2i32(base.Ui64(v174) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v171, v165)
										}
									}
								}
								v200 = base.I64_extend_i32_u(base.B2i32(v190 == int32(0)))
							}
						}
					}
					return v200
				} else {
					v69 = int32(1)
					if base.B2i32(v69 == int32(0))|base.F64_ne(v13, v19) != 0 {
						v200 = int64(1)
						return v200
					} else {
						v75 = v23
						v80 = v26
						v82 = *(*float64)(unsafe.Add(mBase, uint32(v11)+8))
						v85 = base.I64_reinterpret_f64(v82) & int64(9223372036854775807)
						if base.Ui64(v80) <= base.Ui64(int64(9218868437227405312)) {
							v88 = int64(1)
							if base.F64_ne(v82, v75) != 0 {
								v200 = v88
							} else {
								if base.Ui64(v85) < base.Ui64(int64(9218868437227405313)) {
									v104 = *(*float64)(unsafe.Add(mBase, uint32(v12)+16))
									if base.Ui64(base.I64_reinterpret_f64(v104)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
										v110 = *(*float64)(unsafe.Add(mBase, uint32(v11)+16))
										v112 = int64(9223372036854775807)
										v113 = base.I64_reinterpret_f64(v110) & v112
										v114 = *(*float64)(unsafe.Add(mBase, uint32(v12)+24))
										v117 = base.I64_reinterpret_f64(v114) & v112
										if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v117) {
											v156 = base.B2i32(base.Ui64(v113) < base.Ui64(int64(9218868437227405313)))
											v159 = int32(0)
											if base.B2i32(v156 == v159)|base.F64_ne(v104, v110) != 0 {
												v190 = v159
											} else {
												v165 = v114
												v169 = v117
												v171 = *(*float64)(unsafe.Add(mBase, uint32(v11)+24))
												v174 = base.I64_reinterpret_f64(v171) & int64(9223372036854775807)
												if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v169) {
													v190 = base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v174))
												} else {
													v190 = base.B2i32(base.Ui64(v174) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v171, v165)
												}
											}
										} else {
											v122 = int32(0)
											if base.Ui64(int64(9218868437227405312)) < base.Ui64(v113) {
												v190 = v122
											} else {
												v125 = *(*float64)(unsafe.Add(mBase, uint32(v11)+24))
												if base.Ui64(base.I64_reinterpret_f64(v125)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
													if base.B2i32(base.F64_le(base.F64_abs(base.F64_sub(v104, v110)), float64(1e-06)) == int32(0))&base.F64_ne(v104, v110) != 0 {
														v190 = v122
													} else {
														v190 = base.F64_eq(v114, v125) | base.F64_le(base.F64_abs(base.F64_sub(v114, v125)), float64(1e-06))
													}
												} else {
													v156 = int32(1)
													v159 = int32(0)
													if base.B2i32(v156 == v159)|base.F64_ne(v104, v110) != 0 {
														v190 = v159
													} else {
														v165 = v114
														v169 = v117
														v171 = *(*float64)(unsafe.Add(mBase, uint32(v11)+24))
														v174 = base.I64_reinterpret_f64(v171) & int64(9223372036854775807)
														if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v169) {
															v190 = base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v174))
														} else {
															v190 = base.B2i32(base.Ui64(v174) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v171, v165)
														}
													}
												}
											}
										}
									} else {
										v133 = *(*int64)(unsafe.Add(mBase, uint32(v11)+16))
										if base.Ui64(v133&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
											v190 = int32(0)
										} else {
											v138 = *(*float64)(unsafe.Add(mBase, uint32(v12)+24))
											v165 = v138
											v169 = base.I64_reinterpret_f64(v138) & int64(9223372036854775807)
											v171 = *(*float64)(unsafe.Add(mBase, uint32(v11)+24))
											v174 = base.I64_reinterpret_f64(v171) & int64(9223372036854775807)
											if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v169) {
												v190 = base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v174))
											} else {
												v190 = base.B2i32(base.Ui64(v174) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v171, v165)
											}
										}
									}
									v200 = base.I64_extend_i32_u(base.B2i32(v190 == int32(0)))
								} else {
									v200 = v88
								}
							}
							return v200
						} else {
							if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v85) {
								v104 = *(*float64)(unsafe.Add(mBase, uint32(v12)+16))
								if base.Ui64(base.I64_reinterpret_f64(v104)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
									v110 = *(*float64)(unsafe.Add(mBase, uint32(v11)+16))
									v112 = int64(9223372036854775807)
									v113 = base.I64_reinterpret_f64(v110) & v112
									v114 = *(*float64)(unsafe.Add(mBase, uint32(v12)+24))
									v117 = base.I64_reinterpret_f64(v114) & v112
									if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v117) {
										v156 = base.B2i32(base.Ui64(v113) < base.Ui64(int64(9218868437227405313)))
										v159 = int32(0)
										if base.B2i32(v156 == v159)|base.F64_ne(v104, v110) != 0 {
											v190 = v159
										} else {
											v165 = v114
											v169 = v117
											v171 = *(*float64)(unsafe.Add(mBase, uint32(v11)+24))
											v174 = base.I64_reinterpret_f64(v171) & int64(9223372036854775807)
											if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v169) {
												v190 = base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v174))
											} else {
												v190 = base.B2i32(base.Ui64(v174) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v171, v165)
											}
										}
									} else {
										v122 = int32(0)
										if base.Ui64(int64(9218868437227405312)) < base.Ui64(v113) {
											v190 = v122
										} else {
											v125 = *(*float64)(unsafe.Add(mBase, uint32(v11)+24))
											if base.Ui64(base.I64_reinterpret_f64(v125)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
												if base.B2i32(base.F64_le(base.F64_abs(base.F64_sub(v104, v110)), float64(1e-06)) == int32(0))&base.F64_ne(v104, v110) != 0 {
													v190 = v122
												} else {
													v190 = base.F64_eq(v114, v125) | base.F64_le(base.F64_abs(base.F64_sub(v114, v125)), float64(1e-06))
												}
											} else {
												v156 = int32(1)
												v159 = int32(0)
												if base.B2i32(v156 == v159)|base.F64_ne(v104, v110) != 0 {
													v190 = v159
												} else {
													v165 = v114
													v169 = v117
													v171 = *(*float64)(unsafe.Add(mBase, uint32(v11)+24))
													v174 = base.I64_reinterpret_f64(v171) & int64(9223372036854775807)
													if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v169) {
														v190 = base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v174))
													} else {
														v190 = base.B2i32(base.Ui64(v174) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v171, v165)
													}
												}
											}
										}
									}
								} else {
									v133 = *(*int64)(unsafe.Add(mBase, uint32(v11)+16))
									if base.Ui64(v133&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
										v190 = int32(0)
									} else {
										v138 = *(*float64)(unsafe.Add(mBase, uint32(v12)+24))
										v165 = v138
										v169 = base.I64_reinterpret_f64(v138) & int64(9223372036854775807)
										v171 = *(*float64)(unsafe.Add(mBase, uint32(v11)+24))
										v174 = base.I64_reinterpret_f64(v171) & int64(9223372036854775807)
										if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v169) {
											v190 = base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v174))
										} else {
											v190 = base.B2i32(base.Ui64(v174) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v171, v165)
										}
									}
								}
								v200 = base.I64_extend_i32_u(base.B2i32(v190 == int32(0)))
								return v200
							} else {
								return int64(1)
							}
						}
					}
				}
			}
		}
	} else {
		v41 = *(*int64)(unsafe.Add(mBase, uint32(v11)))
		if base.Ui64(v41&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
			return int64(1)
		} else {
			v48 = *(*float64)(unsafe.Add(mBase, uint32(v12)+8))
			v75 = v48
			v80 = base.I64_reinterpret_f64(v48) & int64(9223372036854775807)
			v82 = *(*float64)(unsafe.Add(mBase, uint32(v11)+8))
			v85 = base.I64_reinterpret_f64(v82) & int64(9223372036854775807)
			if base.Ui64(v80) <= base.Ui64(int64(9218868437227405312)) {
				v88 = int64(1)
				if base.F64_ne(v82, v75) != 0 {
					v200 = v88
				} else {
					if base.Ui64(v85) < base.Ui64(int64(9218868437227405313)) {
						v104 = *(*float64)(unsafe.Add(mBase, uint32(v12)+16))
						if base.Ui64(base.I64_reinterpret_f64(v104)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
							v110 = *(*float64)(unsafe.Add(mBase, uint32(v11)+16))
							v112 = int64(9223372036854775807)
							v113 = base.I64_reinterpret_f64(v110) & v112
							v114 = *(*float64)(unsafe.Add(mBase, uint32(v12)+24))
							v117 = base.I64_reinterpret_f64(v114) & v112
							if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v117) {
								v156 = base.B2i32(base.Ui64(v113) < base.Ui64(int64(9218868437227405313)))
								v159 = int32(0)
								if base.B2i32(v156 == v159)|base.F64_ne(v104, v110) != 0 {
									v190 = v159
								} else {
									v165 = v114
									v169 = v117
									v171 = *(*float64)(unsafe.Add(mBase, uint32(v11)+24))
									v174 = base.I64_reinterpret_f64(v171) & int64(9223372036854775807)
									if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v169) {
										v190 = base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v174))
									} else {
										v190 = base.B2i32(base.Ui64(v174) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v171, v165)
									}
								}
							} else {
								v122 = int32(0)
								if base.Ui64(int64(9218868437227405312)) < base.Ui64(v113) {
									v190 = v122
								} else {
									v125 = *(*float64)(unsafe.Add(mBase, uint32(v11)+24))
									if base.Ui64(base.I64_reinterpret_f64(v125)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
										if base.B2i32(base.F64_le(base.F64_abs(base.F64_sub(v104, v110)), float64(1e-06)) == int32(0))&base.F64_ne(v104, v110) != 0 {
											v190 = v122
										} else {
											v190 = base.F64_eq(v114, v125) | base.F64_le(base.F64_abs(base.F64_sub(v114, v125)), float64(1e-06))
										}
									} else {
										v156 = int32(1)
										v159 = int32(0)
										if base.B2i32(v156 == v159)|base.F64_ne(v104, v110) != 0 {
											v190 = v159
										} else {
											v165 = v114
											v169 = v117
											v171 = *(*float64)(unsafe.Add(mBase, uint32(v11)+24))
											v174 = base.I64_reinterpret_f64(v171) & int64(9223372036854775807)
											if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v169) {
												v190 = base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v174))
											} else {
												v190 = base.B2i32(base.Ui64(v174) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v171, v165)
											}
										}
									}
								}
							}
						} else {
							v133 = *(*int64)(unsafe.Add(mBase, uint32(v11)+16))
							if base.Ui64(v133&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
								v190 = int32(0)
							} else {
								v138 = *(*float64)(unsafe.Add(mBase, uint32(v12)+24))
								v165 = v138
								v169 = base.I64_reinterpret_f64(v138) & int64(9223372036854775807)
								v171 = *(*float64)(unsafe.Add(mBase, uint32(v11)+24))
								v174 = base.I64_reinterpret_f64(v171) & int64(9223372036854775807)
								if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v169) {
									v190 = base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v174))
								} else {
									v190 = base.B2i32(base.Ui64(v174) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v171, v165)
								}
							}
						}
						v200 = base.I64_extend_i32_u(base.B2i32(v190 == int32(0)))
					} else {
						v200 = v88
					}
				}
				return v200
			} else {
				if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v85) {
					v104 = *(*float64)(unsafe.Add(mBase, uint32(v12)+16))
					if base.Ui64(base.I64_reinterpret_f64(v104)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
						v110 = *(*float64)(unsafe.Add(mBase, uint32(v11)+16))
						v112 = int64(9223372036854775807)
						v113 = base.I64_reinterpret_f64(v110) & v112
						v114 = *(*float64)(unsafe.Add(mBase, uint32(v12)+24))
						v117 = base.I64_reinterpret_f64(v114) & v112
						if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v117) {
							v156 = base.B2i32(base.Ui64(v113) < base.Ui64(int64(9218868437227405313)))
							v159 = int32(0)
							if base.B2i32(v156 == v159)|base.F64_ne(v104, v110) != 0 {
								v190 = v159
							} else {
								v165 = v114
								v169 = v117
								v171 = *(*float64)(unsafe.Add(mBase, uint32(v11)+24))
								v174 = base.I64_reinterpret_f64(v171) & int64(9223372036854775807)
								if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v169) {
									v190 = base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v174))
								} else {
									v190 = base.B2i32(base.Ui64(v174) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v171, v165)
								}
							}
						} else {
							v122 = int32(0)
							if base.Ui64(int64(9218868437227405312)) < base.Ui64(v113) {
								v190 = v122
							} else {
								v125 = *(*float64)(unsafe.Add(mBase, uint32(v11)+24))
								if base.Ui64(base.I64_reinterpret_f64(v125)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
									if base.B2i32(base.F64_le(base.F64_abs(base.F64_sub(v104, v110)), float64(1e-06)) == int32(0))&base.F64_ne(v104, v110) != 0 {
										v190 = v122
									} else {
										v190 = base.F64_eq(v114, v125) | base.F64_le(base.F64_abs(base.F64_sub(v114, v125)), float64(1e-06))
									}
								} else {
									v156 = int32(1)
									v159 = int32(0)
									if base.B2i32(v156 == v159)|base.F64_ne(v104, v110) != 0 {
										v190 = v159
									} else {
										v165 = v114
										v169 = v117
										v171 = *(*float64)(unsafe.Add(mBase, uint32(v11)+24))
										v174 = base.I64_reinterpret_f64(v171) & int64(9223372036854775807)
										if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v169) {
											v190 = base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v174))
										} else {
											v190 = base.B2i32(base.Ui64(v174) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v171, v165)
										}
									}
								}
							}
						}
					} else {
						v133 = *(*int64)(unsafe.Add(mBase, uint32(v11)+16))
						if base.Ui64(v133&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
							v190 = int32(0)
						} else {
							v138 = *(*float64)(unsafe.Add(mBase, uint32(v12)+24))
							v165 = v138
							v169 = base.I64_reinterpret_f64(v138) & int64(9223372036854775807)
							v171 = *(*float64)(unsafe.Add(mBase, uint32(v11)+24))
							v174 = base.I64_reinterpret_f64(v171) & int64(9223372036854775807)
							if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v169) {
								v190 = base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v174))
							} else {
								v190 = base.B2i32(base.Ui64(v174) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v171, v165)
							}
						}
					}
					v200 = base.I64_extend_i32_u(base.B2i32(v190 == int32(0)))
					return v200
				} else {
					return int64(1)
				}
			}
		}
	}
}
