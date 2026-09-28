package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_lseg_center(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 float64
	_ = v15
	var v16 float64
	_ = v16
	var v17 float64
	_ = v17
	var v19 float64
	_ = v19
	var v31 float64
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 float64
	_ = v35
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v43 float64
	_ = v43
	var v46 float64
	_ = v46
	var v53 float64
	_ = v53
	var v54 int32
	_ = v54
	var v55 float64
	_ = v55
	var v60 float64
	_ = v60
	var v61 int32
	_ = v61
	var v62 float64
	_ = v62
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 float64
	_ = v71
	var v72 float64
	_ = v72
	var v73 float64
	_ = v73
	var v75 float64
	_ = v75
	var v87 float64
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 float64
	_ = v91
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v99 float64
	_ = v99
	var v102 float64
	_ = v102
	var v109 float64
	_ = v109
	var v110 int32
	_ = v110
	var v111 float64
	_ = v111
	var v116 float64
	_ = v116
	var v117 int32
	_ = v117
	var v118 float64
	_ = v118
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v133 int32
	_ = v133
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v10 = F_palloc(m, int32(16))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v15 = *(*float64)(unsafe.Add(mBase, uint32(v8)))
	v16 = *(*float64)(unsafe.Add(mBase, uint32(v8)+16))
	v17 = base.F64_add(v15, v16)
	v19 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_ne(base.F64_abs(v17), v19)|base.F64_eq(base.F64_abs(v15), v19)|base.F64_eq(base.F64_abs(v16), v19) == int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v31 = F_float_overflow_error_ext(m, v14)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L1
	} else {
		goto L6
	}
L4:
	;
	v34 = v14
	v35 = v17
	goto L5
L5:
	;
	if v34 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v34 = v33
	v35 = v31
	goto L5
L7:
	;
	v133 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v133)
	return int64(0)
L8:
	;
	v43 = math.Float64frombits(uint64(0x7ff0000000000000))
	v46 = base.F64_mul(v35, float64(0.5))
	if base.F64_eq(base.F64_abs(v35), v43)|base.F64_ne(base.F64_abs(v46), v43) == int32(0) {
		goto L13
	} else {
		goto L14
	}
L9:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
	if v38 != int32(453) {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+4)))
	if v41 != 0 {
		goto L7
	} else {
		goto L11
	}
L11:
	;
	goto L8
L12:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v10))) = v62
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v64 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L13:
	;
	v53 = F_float_overflow_error_ext(m, v34)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L1
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v55 = float64(0)
	if base.F64_eq(v35, v55)|base.F64_ne(v46, v55) != 0 {
		v62 = v46
		goto L12
	} else {
		goto L17
	}
L16:
	;
	v62 = v53
	goto L12
L17:
	;
	v60 = F_float_underflow_error_ext(m, v34)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v62 = v60
	goto L12
L19:
	;
	v71 = *(*float64)(unsafe.Add(mBase, uint32(v8)+8))
	v72 = *(*float64)(unsafe.Add(mBase, uint32(v8)+24))
	v73 = base.F64_add(v71, v72)
	v75 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_ne(base.F64_abs(v73), v75)|base.F64_eq(base.F64_abs(v71), v75)|base.F64_eq(base.F64_abs(v72), v75) == int32(0) {
		goto L23
	} else {
		goto L24
	}
L20:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
	if v67 != int32(453) {
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64)+4)))
	if v70 != 0 {
		goto L7
	} else {
		goto L22
	}
L22:
	;
	goto L19
L23:
	;
	v87 = F_float_overflow_error_ext(m, v64)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L1
	} else {
		goto L26
	}
L24:
	;
	v90 = v64
	v91 = v73
	goto L25
L25:
	;
	if v90 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v90 = v89
	v91 = v87
	goto L25
L27:
	;
	v99 = math.Float64frombits(uint64(0x7ff0000000000000))
	v102 = base.F64_mul(v91, float64(0.5))
	if base.F64_eq(base.F64_abs(v91), v99)|base.F64_ne(base.F64_abs(v102), v99) == int32(0) {
		goto L32
	} else {
		goto L33
	}
L28:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v90)))
	if v94 != int32(453) {
		goto L27
	} else {
		goto L29
	}
L29:
	;
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90)+4)))
	if v97 != 0 {
		goto L7
	} else {
		goto L30
	}
L30:
	;
	goto L27
L31:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v10)+8)) = v118
	v120 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v120 == int32(0) {
		goto L38
	} else {
		goto L39
	}
L32:
	;
	v109 = F_float_overflow_error_ext(m, v90)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L1
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	v111 = float64(0)
	if base.F64_eq(v91, v111)|base.F64_ne(v102, v111) != 0 {
		v118 = v102
		goto L31
	} else {
		goto L36
	}
L35:
	;
	v118 = v109
	goto L31
L36:
	;
	v116 = F_float_underflow_error_ext(m, v90)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	v118 = v116
	goto L31
L38:
	;
	return base.I64_extend_i32_u(v10)
L39:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v120)))
	if v123 != int32(453) {
		goto L38
	} else {
		goto L40
	}
L40:
	;
	v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v120)+4)))
	if v126 != 0 {
		goto L7
	} else {
		goto L41
	}
L41:
	;
	goto L38
}
func F_lseg_construct(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 float64
	_ = v11
	var v13 float64
	_ = v13
	var v15 float64
	_ = v15
	var v17 float64
	_ = v17
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v7 = F_palloc(m, int32(32))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int64(0)
	} else {
		v11 = *(*float64)(unsafe.Add(mBase, uint32(v5)))
		*(*float64)(unsafe.Add(mBase, uint32(v7))) = v11
		v13 = *(*float64)(unsafe.Add(mBase, uint32(v5)+8))
		*(*float64)(unsafe.Add(mBase, uint32(v7)+8)) = v13
		v15 = *(*float64)(unsafe.Add(mBase, uint32(v4)))
		*(*float64)(unsafe.Add(mBase, uint32(v7)+16)) = v15
		v17 = *(*float64)(unsafe.Add(mBase, uint32(v4)+8))
		*(*float64)(unsafe.Add(mBase, uint32(v7)+24)) = v17
		return base.I64_extend_i32_u(v7)
	}
}
func F_lseg_inside_poly(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v23 int64
	_ = v23
	var v25 int64
	_ = v25
	var v27 int64
	_ = v27
	var v29 int64
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v40 int64
	_ = v40
	var v42 int64
	_ = v42
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int64
	_ = v71
	var v73 int64
	_ = v73
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v118 int64
	_ = v118
	var v120 int64
	_ = v120
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v134 int32
	_ = v134
	var v141 float64
	_ = v141
	var v142 float64
	_ = v142
	var v143 float64
	_ = v143
	var v145 float64
	_ = v145
	var v158 float64
	_ = v158
	var v159 int32
	_ = v159
	var v160 float64
	_ = v160
	var v162 float64
	_ = v162
	var v165 float64
	_ = v165
	var v173 float64
	_ = v173
	var v174 int32
	_ = v174
	var v175 float64
	_ = v175
	var v181 float64
	_ = v181
	var v182 int32
	_ = v182
	var v183 float64
	_ = v183
	var v185 float64
	_ = v185
	var v186 float64
	_ = v186
	var v187 float64
	_ = v187
	var v189 float64
	_ = v189
	var v202 float64
	_ = v202
	var v203 int32
	_ = v203
	var v204 float64
	_ = v204
	var v206 float64
	_ = v206
	var v209 float64
	_ = v209
	var v217 float64
	_ = v217
	var v218 int32
	_ = v218
	var v219 float64
	_ = v219
	var v225 float64
	_ = v225
	var v226 int32
	_ = v226
	var v227 float64
	_ = v227
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v236 int32
	_ = v236
	v5 = int32(0)
	v15 = m.G0
	v17 = v15 - int32(80)
	m.G0 = v17
	F_check_stack_depth(m)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v23 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+24)) = v23
	v25 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+16)) = v25
	v27 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+32)) = v27
	v29 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+40)) = v29
	v32 = l2 + int32(40)
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if l3 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v34 = l3
	goto L5
L4:
	;
	v34 = v33
	goto L5
L5:
	;
	v39 = v32 + v34<<(uint(int32(4))%32) - int32(16)
	v40 = *(*int64)(unsafe.Add(mBase, uint32(v39)))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+48)) = v40
	v42 = *(*int64)(unsafe.Add(mBase, uint32(v39)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+56)) = v42
	if v33 <= l3 {
		goto L9
	} else {
		goto L10
	}
L6:
	;
	m.G0 = v17 + int32(80)
	return v236
L7:
	;
	v236 = int32(1)
	goto L6
L8:
	;
	if base.B2i32(v125 == int32(0))|v134 != 0 {
		v236 = v125
		goto L6
	} else {
		goto L39
	}
L9:
	;
	v125 = int32(1)
	v134 = v5
	goto L8
L10:
	;
	goto L11
L11:
	;
	v47 = v17 + int32(32)
	v49 = v17 - int32(-64)
	v53 = l3
	v60 = v5
	goto L12
L12:
	;
	v65 = *(*int32)(unsafe.Add(mBase, _c_F_lseg_inside_poly[0]))
	if v65 != 0 {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	v125 = v114
	v134 = v117
	goto L8
L14:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L1
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v70 = v32 + v53<<(uint(int32(4))%32)
	v71 = *(*int64)(unsafe.Add(mBase, uint32(v70)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v49)+8)) = v71
	v73 = *(*int64)(unsafe.Add(mBase, uint32(v70)))
	*(*int64)(unsafe.Add(mBase, uint32(v49))) = v73
	v76 = v17 + int32(48)
	v78 = v17 + int32(16)
	v79 = F_lseg_contain_point(m, v76, v78)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L1
	} else {
		goto L18
	}
L17:
	;
	goto L16
L18:
	;
	v81 = F_lseg_contain_point(m, v76, v47)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	if v79 != 0 {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	v118 = *(*int64)(unsafe.Add(mBase, uint32(v49)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+56)) = v118
	v120 = *(*int64)(unsafe.Add(mBase, uint32(v49)))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+48)) = v120
	v122 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v122 <= v115 {
		v125 = v114
		v134 = v117
		goto L8
	} else {
		goto L37
	}
L21:
	;
	if v81 != 0 {
		goto L7
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	if v81 != 0 {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	v84 = v53 + int32(1)
	v85 = F_touched_lseg_inside_poly(m, v78, v47, v76, l2, v84)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	v114 = v85
	v115 = v84
	v117 = v60
	goto L20
L26:
	;
	v92 = v53 + int32(1)
	v93 = F_touched_lseg_inside_poly(m, v47, v17+int32(16), v17+int32(48), l2, v92)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L1
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	v95 = int32(1)
	v97 = v53 + v95
	v99 = v17 + int32(16)
	v102 = F_lseg_interpt_lseg(m, v17, v99, v17+int32(48))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L1
	} else {
		goto L30
	}
L29:
	;
	v114 = v93
	v115 = v92
	v117 = v60
	goto L20
L30:
	;
	if v102 == int32(0) {
		v114 = v95
		v115 = v97
		v117 = v60
		goto L20
	} else {
		goto L31
	}
L31:
	;
	v106 = F_lseg_inside_poly(m, v99, v17, l2, v97)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	if v106 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v236 = int32(0)
	goto L6
L34:
	;
	goto L35
L35:
	;
	v112 = F_lseg_inside_poly(m, v47, v17, l2, v97)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	v114 = v112
	v115 = v97
	v117 = int32(1)
	goto L20
L37:
	;
	if v114 != 0 {
		v53 = v115
		v60 = v117
		goto L12
	} else {
		goto L38
	}
L38:
	;
	goto L13
L39:
	;
	v141 = *(*float64)(unsafe.Add(mBase, uint32(v17)+16))
	v142 = *(*float64)(unsafe.Add(mBase, uint32(v17)+32))
	v143 = base.F64_add(v141, v142)
	v145 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_ne(base.F64_abs(v143), v145)|base.F64_eq(base.F64_abs(v141), v145)|base.F64_eq(base.F64_abs(v142), v145) == int32(0) {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v158 = F_float_overflow_error_ext(m, int32(0))
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L1
	} else {
		goto L43
	}
L41:
	;
	v160 = v143
	goto L42
L42:
	;
	v162 = math.Float64frombits(uint64(0x7ff0000000000000))
	v165 = base.F64_mul(v160, float64(0.5))
	if base.F64_eq(base.F64_abs(v160), v162)|base.F64_ne(base.F64_abs(v165), v162) == int32(0) {
		goto L45
	} else {
		goto L46
	}
L43:
	;
	v160 = v158
	goto L42
L44:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v17))) = v183
	v185 = *(*float64)(unsafe.Add(mBase, uint32(v17)+24))
	v186 = *(*float64)(unsafe.Add(mBase, uint32(v17)+40))
	v187 = base.F64_add(v185, v186)
	v189 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_ne(base.F64_abs(v187), v189)|base.F64_eq(base.F64_abs(v185), v189)|base.F64_eq(base.F64_abs(v186), v189) == int32(0) {
		goto L51
	} else {
		goto L52
	}
L45:
	;
	v173 = F_float_overflow_error_ext(m, int32(0))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L1
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	v175 = float64(0)
	if base.F64_eq(v160, v175)|base.F64_ne(v165, v175) != 0 {
		v183 = v165
		goto L44
	} else {
		goto L49
	}
L48:
	;
	v183 = v173
	goto L44
L49:
	;
	v181 = F_float_underflow_error_ext(m, int32(0))
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	v183 = v181
	goto L44
L51:
	;
	v202 = F_float_overflow_error_ext(m, int32(0))
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L1
	} else {
		goto L54
	}
L52:
	;
	v204 = v187
	goto L53
L53:
	;
	v206 = math.Float64frombits(uint64(0x7ff0000000000000))
	v209 = base.F64_mul(v204, float64(0.5))
	if base.F64_eq(base.F64_abs(v204), v206)|base.F64_ne(base.F64_abs(v209), v206) == int32(0) {
		goto L56
	} else {
		goto L57
	}
L54:
	;
	v204 = v202
	goto L53
L55:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v17)+8)) = v227
	v229 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v230 = F_point_inside(m, v17, v229, v32)
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L1
	} else {
		goto L62
	}
L56:
	;
	v217 = F_float_overflow_error_ext(m, int32(0))
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L1
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	v219 = float64(0)
	if base.F64_eq(v204, v219)|base.F64_ne(v209, v219) != 0 {
		v227 = v209
		goto L55
	} else {
		goto L60
	}
L59:
	;
	v227 = v217
	goto L55
L60:
	;
	v225 = F_float_underflow_error_ext(m, int32(0))
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	v227 = v225
	goto L55
L62:
	;
	v236 = base.B2i32(v230 != int32(0))
	goto L6
}
func F_lseg_parallel(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 float64
	_ = v9
	var v12 int32
	_ = v12
	var v15 float64
	_ = v15
	var v16 int32
	_ = v16
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v9 = F_point_sl(m, v6, v6+int32(16))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int64(0)
	} else {
		v15 = F_point_sl(m, v5, v5+int32(16))
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int64(0)
		} else {
			return base.I64_extend_i32_u(base.F64_eq(v9, v15) | base.F64_le(base.F64_abs(base.F64_sub(v9, v15)), float64(1e-06)))
		}
	}
}
