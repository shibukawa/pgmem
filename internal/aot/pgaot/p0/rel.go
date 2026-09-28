package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_fillRelOptions(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v25 int32
	_ = v25
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v53 int32
	_ = v53
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 float64
	_ = v117
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v227 int32
	_ = v227
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v240 int32
	_ = v240
	var v245 int32
	_ = v245
	var v266 int32
	_ = v266
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v277 int32
	_ = v277
	var v282 int32
	_ = v282
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v296 int32
	_ = v296
	var v311 int32
	_ = v311
	var v314 int32
	_ = v314
	v8 = int32(0)
	v16 = m.G0
	v18 = v16 - int32(32)
	m.G0 = v18
	if v8 < l3 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v25 = l1
	v35 = v8
	goto L4
L2:
	;
	v314 = l1
	goto L3
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v314 << (uint(int32(2)) % 32)
	m.G0 = v18 + int32(32)
	return
L4:
	;
	if l6 <= int32(0) {
		goto L9
	} else {
		goto L10
	}
L5:
	;
	v314 = v296
	goto L3
L6:
	;
	v311 = v35 + int32(1)
	if v311 != l3 {
		v25 = v296
		v35 = v311
		goto L4
	} else {
		goto L88
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v94))) = int32(0)
	v296 = v25
	goto L6
L8:
	;
	v287 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+4)))
	if v287 != 0 {
		goto L85
	} else {
		goto L86
	}
L9:
	;
	if l4 == int32(0) {
		v296 = v25
		goto L6
	} else {
		goto L81
	}
L10:
	;
	v41 = l2 + v35<<(uint(int32(4))%32)
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
	v53 = int32(0)
	goto L11
L11:
	;
	v62 = l5 + v53*int32(12)
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)))
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43))))
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63))))
	if base.B2i32(v66 == int32(0))|base.B2i32(v66 != v69) != 0 {
		v87 = v66
		v88 = v69
		goto L14
	} else {
		goto L15
	}
L12:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v62)+8))
	v94 = l0 + v93
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v42)+20))
	switch v95 {
	case 0:
		goto L8
	case 1:
		goto L29
	case 2:
		goto L28
	case 3:
		goto L27
	case 4:
		goto L26
	case 5:
		goto L25
	default:
		goto L24
	}
L13:
	;
	if v87-v88 != 0 {
		goto L20
	} else {
		goto L21
	}
L14:
	;
	goto L13
L15:
	;
	v72 = v43
	v73 = v63
	goto L16
L16:
	;
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73)+1)))
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72)+1)))
	if v77 == int32(0) {
		v87 = v77
		v88 = v76
		goto L14
	} else {
		goto L18
	}
L17:
	;
	v87 = v77
	v88 = v76
	goto L14
L18:
	;
	v80 = int32(1)
	if v77 == v76 {
		v72 = v72 + v80
		v73 = v73 + v80
		goto L16
	} else {
		goto L19
	}
L19:
	;
	goto L17
L20:
	;
	v91 = v53 + int32(1)
	if l6 != v91 {
		v53 = v91
		goto L11
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	goto L12
L23:
	;
	goto L9
L24:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L52
	} else {
		goto L78
	}
L25:
	;
	v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+4)))
	if v127 == int32(1) {
		goto L45
	} else {
		goto L46
	}
L26:
	;
	v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+4)))
	if v123 != 0 {
		goto L39
	} else {
		goto L40
	}
L27:
	;
	v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+4)))
	if v115 != 0 {
		goto L36
	} else {
		goto L37
	}
L28:
	;
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+4)))
	if v107 != 0 {
		goto L33
	} else {
		goto L34
	}
L29:
	;
	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+4)))
	if v96 == int32(1) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v41)+8))
	v101 = v99
	goto L32
L31:
	;
	v101 = int32(-1)
	goto L32
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v94))) = v101
	v296 = v25
	goto L6
L33:
	;
	v108 = v41 + int32(8)
	goto L35
L34:
	;
	v108 = v42 + int32(24)
	goto L35
L35:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v108)))
	*(*int32)(unsafe.Add(mBase, uint32(v94))) = v109
	v296 = v25
	goto L6
L36:
	;
	v116 = v41 + int32(8)
	goto L38
L37:
	;
	v116 = v42 + int32(24)
	goto L38
L38:
	;
	v117 = *(*float64)(unsafe.Add(mBase, uint32(v116)))
	*(*float64)(unsafe.Add(mBase, uint32(v94))) = v117
	v296 = v25
	goto L6
L39:
	;
	v124 = v41 + int32(8)
	goto L41
L40:
	;
	v124 = v42 + int32(28)
	goto L41
L41:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v124)))
	*(*int32)(unsafe.Add(mBase, uint32(v94))) = v125
	v296 = v25
	goto L6
L42:
	;
	v151 = l0 + v25
	if (v136^v151)&int32(3) != 0 {
		goto L60
	} else {
		goto L61
	}
L43:
	;
	v145 = m.T0[v143].(func(*base.Module, int32, int32) int32)(m, v142, l0+v25)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L52
	} else {
		goto L53
	}
L44:
	;
	v138 = int32(0)
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v42)+36))
	if v139 == v138 {
		goto L7
	} else {
		goto L51
	}
L45:
	;
	v135 = v41 + int32(8)
	goto L47
L46:
	;
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+28)))
	if v132 != 0 {
		goto L44
	} else {
		goto L48
	}
L47:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v135)))
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v42)+36))
	if v137 != 0 {
		v142 = v136
		v143 = v137
		goto L43
	} else {
		goto L49
	}
L48:
	;
	v135 = v42 + int32(40)
	goto L47
L49:
	;
	if v136 != 0 {
		goto L42
	} else {
		goto L50
	}
L50:
	;
	goto L7
L51:
	;
	v142 = v138
	v143 = v139
	goto L43
L52:
	;
	return
L53:
	;
	if v145 != 0 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v94))) = v25
	v296 = v25 + v145
	goto L6
L55:
	;
	goto L56
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v94))) = int32(0)
	v296 = v25
	goto L6
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v94))) = v25
	v227 = F_strlen(m, v136)
	mBase = m.M
	v296 = v227 + v25 + int32(1)
	goto L6
L58:
	;
	goto L57
L59:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v206))) = uint8(v205)
	if v205&int32(255) == int32(0) {
		goto L58
	} else {
		goto L74
	}
L60:
	;
	v157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v136))))
	v204 = v136
	v205 = v157
	v206 = v151
	goto L59
L61:
	;
	goto L62
L62:
	;
	if v136&int32(3) != 0 {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v161 = v136
	v163 = v151
	goto L66
L64:
	;
	v175 = v136
	v177 = v151
	goto L65
L65:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v175)))
	v182 = int32(-2139062144)
	if (int32(16843008)-v179|v179)&v182 != v182 {
		v204 = v175
		v205 = v179
		v206 = v177
		goto L59
	} else {
		goto L70
	}
L66:
	;
	v164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161))))
	*(*uint8)(unsafe.Add(mBase, uint32(v163))) = uint8(v164)
	if v164 == int32(0) {
		goto L58
	} else {
		goto L68
	}
L67:
	;
	v175 = v171
	v177 = v169
	goto L65
L68:
	;
	v168 = int32(1)
	v169 = v163 + v168
	v171 = v161 + v168
	if v171&int32(3) != 0 {
		v161 = v171
		v163 = v169
		goto L66
	} else {
		goto L69
	}
L69:
	;
	goto L67
L70:
	;
	v187 = v175
	v188 = v179
	v189 = v177
	goto L71
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v189))) = v188
	v191 = int32(4)
	v192 = v189 + v191
	v194 = v187 + v191
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v187)+4))
	v199 = int32(-2139062144)
	if (int32(16843008)-v196|v196)&v199 == v199 {
		v187 = v194
		v188 = v196
		v189 = v192
		goto L71
	} else {
		goto L73
	}
L72:
	;
	v204 = v194
	v205 = v196
	v206 = v192
	goto L59
L73:
	;
	goto L72
L74:
	;
	v213 = v204
	v215 = v206
	goto L75
L75:
	;
	v216 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v213)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v215)+1)) = uint8(v216)
	v218 = int32(1)
	if v216 != 0 {
		v213 = v213 + v218
		v215 = v215 + v218
		goto L75
	} else {
		goto L77
	}
L76:
	;
	goto L58
L77:
	;
	goto L76
L78:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v235)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v236
	F_errmsg_internal(m, int32(_a_F_fillRelOptions_0), v18)
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L52
	} else {
		goto L79
	}
L79:
	;
	F_errfinish(m, int32(_a_F_fillRelOptions_1), int32(1956), int32(_a_F_fillRelOptions_2))
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L52
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
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L52
	} else {
		goto L82
	}
L82:
	;
	v270 = *(*int32)(unsafe.Add(mBase, uint32(l2+v35<<(uint(int32(4))%32))))
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v270)))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v271
	F_errmsg_internal(m, int32(_a_F_fillRelOptions_3), v18+int32(16))
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L52
	} else {
		goto L83
	}
L83:
	;
	F_errfinish(m, int32(_a_F_fillRelOptions_1), int32(1965), int32(_a_F_fillRelOptions_2))
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L52
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
	v288 = v41 + int32(8)
	goto L87
L86:
	;
	v288 = v42 + int32(24)
	goto L87
L87:
	;
	v289 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v288))))
	*(*uint8)(unsafe.Add(mBase, uint32(v94))) = uint8(v289)
	v296 = v25
	goto L6
L88:
	;
	goto L5
}
func F_get_rel_data_width(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int64
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int64
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int64
	_ = v66
	var v69 int32
	_ = v69
	var v77 int64
	_ = v77
	var v78 int64
	_ = v78
	var v81 int64
	_ = v81
	v7 = int64(0)
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v9 = int32(*(*int16)(unsafe.Add(mBase, uint32(v8)+120)))
	if int32(0) < v9 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v15 = v8
	v17 = int32(1)
	v19 = v7
	goto L4
L2:
	;
	v77 = v7
	goto L3
L3:
	;
	v78 = int64(1073741823)
	if v78 <= v77 {
		goto L22
	} else {
		goto L23
	}
L4:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	v27 = v20 + v21<<(uint(int32(3))%32) + v17*int32(100)
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+19)))
	if v28 != 0 {
		v63 = v15
		v66 = v19
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v77 = v66
	goto L3
L6:
	;
	v69 = int32(*(*int16)(unsafe.Add(mBase, uint32(v63)+120)))
	if v17 < v69 {
		v15 = v63
		v17 = v17 + int32(1)
		v19 = v66
		goto L4
	} else {
		goto L20
	}
L7:
	;
	if l1 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v42 = F_get_attavgwidth(m, v40, base.I32_extend16_s(v17))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l1+v17<<(uint(int32(2))%32))))
	if v34 <= int32(0) {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v63 = v15
	v66 = v19 + base.I64_extend_i32_u(v34)
	goto L6
L11:
	;
	return int32(0)
L12:
	;
	if v42 <= int32(0) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v49 = v27 - int32(72)
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)+68))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v49)+76))
	v52 = F_get_typavgwidth(m, v50, v51)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L11
	} else {
		goto L16
	}
L14:
	;
	v54 = v42
	goto L15
L15:
	;
	if l1 != 0 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v54 = v52
	goto L15
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1+v17<<(uint(int32(2))%32)))) = v54
	goto L19
L18:
	;
	goto L19
L19:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v63 = v62
	v66 = v19 + base.I64_extend_i32_s(v54)
	goto L6
L20:
	;
	goto L5
L21:
	;
	return base.I32_wrap_i64(v81)
L22:
	;
	v81 = v78
	goto L24
L23:
	;
	v81 = v77
	goto L24
L24:
	;
	goto L21
}
func F_get_rel_name(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14286(m, l0, int32(57))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_rel_supports_distinctness(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v68 int32
	_ = v68
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v5 != 0 {
		v68 = int32(0)
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v68
L2:
	;
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)+84))
	switch v6 {
	case 0:
		goto L5
	case 1:
		goto L4
	default:
		goto L3
	}
L3:
	;
	v68 = int32(0)
	goto L1
L4:
	;
	v39 = int32(1)
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v40+v41<<(uint(int32(2))%32))))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)+36))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)+120))
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46)+38)))
	if v48 == v39 {
		goto L18
	} else {
		goto L19
	}
L5:
	;
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)+116))
	if v7 == int32(0) {
		goto L3
	} else {
		goto L6
	}
L6:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	if v10 <= int32(0) {
		goto L3
	} else {
		goto L7
	}
L7:
	;
	v13 = int32(0)
	if v13 < v10 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v17 = v10
	goto L10
L9:
	;
	v17 = v13
	goto L10
L10:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
	v20 = v13
	goto L11
L11:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v18+v20<<(uint(int32(2))%32))))
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+101)))
	if v27 != int32(1) {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	goto L3
L13:
	;
	v37 = v20 + int32(1)
	if v37 != v17 {
		v20 = v37
		goto L11
	} else {
		goto L17
	}
L14:
	;
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+103)))
	if v30 != int32(1) {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v26)+88))
	if v33 != 0 {
		goto L13
	} else {
		goto L16
	}
L16:
	;
	return int32(1)
L17:
	;
	goto L12
L18:
	;
	if v47 == int32(0) {
		goto L3
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	if v47 != 0 {
		v68 = v39
		goto L1
	} else {
		goto L23
	}
L21:
	;
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46)+40)))
	if v53 == int32(1) {
		goto L3
	} else {
		goto L22
	}
L22:
	;
	v68 = v39
	goto L1
L23:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v46)+100))
	if v56 != 0 {
		v68 = v39
		goto L1
	} else {
		goto L24
	}
L24:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v46)+108))
	if v57 != 0 {
		v68 = v39
		goto L1
	} else {
		goto L25
	}
L25:
	;
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46)+36)))
	if v58 != 0 {
		v68 = v39
		goto L1
	} else {
		goto L26
	}
L26:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v46)+112))
	if v59 != 0 {
		v68 = v39
		goto L1
	} else {
		goto L27
	}
L27:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v46)+144))
	if v60 != 0 {
		v68 = v39
		goto L1
	} else {
		goto L28
	}
L28:
	;
	goto L3
}
