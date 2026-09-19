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
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v54 int32
	_ = v54
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
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 float64
	_ = v118
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v152 int32
	_ = v152
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
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
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v228 int32
	_ = v228
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v241 int32
	_ = v241
	var v246 int32
	_ = v246
	var v267 int32
	_ = v267
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v278 int32
	_ = v278
	var v283 int32
	_ = v283
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v297 int32
	_ = v297
	var v312 int32
	_ = v312
	var v315 int32
	_ = v315
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
	v36 = v8
	goto L4
L2:
	;
	v315 = l1
	goto L3
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v315 << (uint(int32(2)) % 32)
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
	v315 = v297
	goto L3
L6:
	;
	v312 = v36 + int32(1)
	if v312 != l3 {
		v25 = v297
		v36 = v312
		goto L4
	} else {
		goto L87
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v94))) = int32(0)
	v297 = v25
	goto L6
L8:
	;
	v288 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+4)))
	if v288 != 0 {
		goto L84
	} else {
		goto L85
	}
L9:
	;
	if l4 == int32(0) {
		v297 = v25
		goto L6
	} else {
		goto L80
	}
L10:
	;
	v41 = l2 + v36<<(uint(int32(4))%32)
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
	v54 = int32(0)
	goto L11
L11:
	;
	v62 = l5 + v54<<(uint(int32(4))%32)
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
	v94 = v93 + l0
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v62)+12))
	if int32(0) < v95 {
		goto L24
	} else {
		goto L25
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
	v91 = v54 + int32(1)
	if l6 != v91 {
		v54 = v91
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
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+4)))
	*(*uint8)(unsafe.Add(mBase, uint32(l0+v95))) = uint8(v99)
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
	v102 = v101
	goto L26
L25:
	;
	v102 = v42
	goto L26
L26:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v102)+20))
	switch v103 {
	case 0:
		goto L8
	case 1:
		goto L31
	case 2:
		goto L30
	case 3:
		goto L29
	case 4:
		goto L28
	default:
		goto L27
	}
L27:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L51
	} else {
		goto L77
	}
L28:
	;
	v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+4)))
	if v128 == int32(1) {
		goto L44
	} else {
		goto L45
	}
L29:
	;
	v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+4)))
	if v124 != 0 {
		goto L38
	} else {
		goto L39
	}
L30:
	;
	v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+4)))
	if v116 != 0 {
		goto L35
	} else {
		goto L36
	}
L31:
	;
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+4)))
	if v108 != 0 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v109 = v41 + int32(8)
	goto L34
L33:
	;
	v109 = v102 + int32(24)
	goto L34
L34:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v109)))
	*(*int32)(unsafe.Add(mBase, uint32(v94))) = v110
	v297 = v25
	goto L6
L35:
	;
	v117 = v41 + int32(8)
	goto L37
L36:
	;
	v117 = v102 + int32(24)
	goto L37
L37:
	;
	v118 = *(*float64)(unsafe.Add(mBase, uint32(v117)))
	*(*float64)(unsafe.Add(mBase, uint32(v94))) = v118
	v297 = v25
	goto L6
L38:
	;
	v125 = v41 + int32(8)
	goto L40
L39:
	;
	v125 = v102 + int32(28)
	goto L40
L40:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v125)))
	*(*int32)(unsafe.Add(mBase, uint32(v94))) = v126
	v297 = v25
	goto L6
L41:
	;
	v152 = l0 + v25
	if (v137^v152)&int32(3) != 0 {
		goto L59
	} else {
		goto L60
	}
L42:
	;
	v146 = m.T0[v144].(func(*base.Module, int32, int32) int32)(m, v143, l0+v25)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L51
	} else {
		goto L52
	}
L43:
	;
	v139 = int32(0)
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v102)+36))
	if v140 == v139 {
		goto L7
	} else {
		goto L50
	}
L44:
	;
	v136 = v41 + int32(8)
	goto L46
L45:
	;
	v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102)+28)))
	if v133 != 0 {
		goto L43
	} else {
		goto L47
	}
L46:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v136)))
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v102)+36))
	if v138 != 0 {
		v143 = v137
		v144 = v138
		goto L42
	} else {
		goto L48
	}
L47:
	;
	v136 = v102 + int32(40)
	goto L46
L48:
	;
	if v137 != 0 {
		goto L41
	} else {
		goto L49
	}
L49:
	;
	goto L7
L50:
	;
	v143 = v139
	v144 = v140
	goto L42
L51:
	;
	return
L52:
	;
	if v146 != 0 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v94))) = v25
	v297 = v25 + v146
	goto L6
L54:
	;
	goto L55
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v94))) = int32(0)
	v297 = v25
	goto L6
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v94))) = v25
	v228 = F_strlen(m, v137)
	mBase = m.M
	v297 = v228 + v25 + int32(1)
	goto L6
L57:
	;
	goto L56
L58:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v207))) = uint8(v206)
	if v206&int32(255) == int32(0) {
		goto L57
	} else {
		goto L73
	}
L59:
	;
	v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v137))))
	v205 = v137
	v206 = v158
	v207 = v152
	goto L58
L60:
	;
	goto L61
L61:
	;
	if v137&int32(3) != 0 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v162 = v137
	v164 = v152
	goto L65
L63:
	;
	v176 = v137
	v178 = v152
	goto L64
L64:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v176)))
	v183 = int32(-2139062144)
	if (int32(16843008)-v180|v180)&v183 != v183 {
		v205 = v176
		v206 = v180
		v207 = v178
		goto L58
	} else {
		goto L69
	}
L65:
	;
	v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162))))
	*(*uint8)(unsafe.Add(mBase, uint32(v164))) = uint8(v165)
	if v165 == int32(0) {
		goto L57
	} else {
		goto L67
	}
L66:
	;
	v176 = v172
	v178 = v170
	goto L64
L67:
	;
	v169 = int32(1)
	v170 = v164 + v169
	v172 = v162 + v169
	if v172&int32(3) != 0 {
		v162 = v172
		v164 = v170
		goto L65
	} else {
		goto L68
	}
L68:
	;
	goto L66
L69:
	;
	v188 = v176
	v189 = v180
	v190 = v178
	goto L70
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v190))) = v189
	v192 = int32(4)
	v193 = v190 + v192
	v195 = v188 + v192
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v188)+4))
	v200 = int32(-2139062144)
	if (int32(16843008)-v197|v197)&v200 == v200 {
		v188 = v195
		v189 = v197
		v190 = v193
		goto L70
	} else {
		goto L72
	}
L71:
	;
	v205 = v195
	v206 = v197
	v207 = v193
	goto L58
L72:
	;
	goto L71
L73:
	;
	v214 = v205
	v216 = v207
	goto L74
L74:
	;
	v217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v214)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v216)+1)) = uint8(v217)
	v219 = int32(1)
	if v217 != 0 {
		v214 = v214 + v219
		v216 = v216 + v219
		goto L74
	} else {
		goto L76
	}
L75:
	;
	goto L57
L76:
	;
	goto L75
L77:
	;
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v236)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v237
	F_errmsg_internal(m, int32(_a_F_fillRelOptions_0), v18)
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L51
	} else {
		goto L78
	}
L78:
	;
	F_errfinish(m, int32(_a_F_fillRelOptions_1), int32(1850), int32(_a_F_fillRelOptions_2))
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L51
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
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L51
	} else {
		goto L81
	}
L81:
	;
	v271 = *(*int32)(unsafe.Add(mBase, uint32(l2+v36<<(uint(int32(4))%32))))
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v271)))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v272
	F_errmsg_internal(m, int32(_a_F_fillRelOptions_3), v18+int32(16))
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L51
	} else {
		goto L82
	}
L82:
	;
	F_errfinish(m, int32(_a_F_fillRelOptions_1), int32(1859), int32(_a_F_fillRelOptions_2))
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L51
	} else {
		goto L83
	}
L83:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L84:
	;
	v289 = v41 + int32(8)
	goto L86
L85:
	;
	v289 = v102 + int32(24)
	goto L86
L86:
	;
	v290 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v289))))
	*(*uint8)(unsafe.Add(mBase, uint32(v94))) = uint8(v290)
	v297 = v25
	goto L6
L87:
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
	v27 = v20 + v21<<(uint(int32(4))%32) + v17*int32(100)
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+11)))
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
	v49 = v27 - int32(80)
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
	v3 = Fn13897(m, l0, int32(57))
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
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
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
	var v55 int32
	_ = v55
	var v69 int32
	_ = v69
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v5 != 0 {
		v69 = int32(0)
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v69
L2:
	;
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
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
	v69 = int32(0)
	goto L1
L4:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v39+v40<<(uint(int32(2))%32))))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)+36))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)+120))
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45)+38)))
	if v47 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L5:
	;
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)+108))
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
	v50 = int32(1)
	if v46 != 0 {
		v69 = v50
		goto L1
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	if v46 == int32(0) {
		goto L3
	} else {
		goto L27
	}
L21:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v45)+100))
	if v51 != 0 {
		v69 = v50
		goto L1
	} else {
		goto L22
	}
L22:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v45)+108))
	if v52 != 0 {
		v69 = v50
		goto L1
	} else {
		goto L23
	}
L23:
	;
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45)+36)))
	if v53 != 0 {
		v69 = v50
		goto L1
	} else {
		goto L24
	}
L24:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v45)+112))
	if v54 != 0 {
		v69 = v50
		goto L1
	} else {
		goto L25
	}
L25:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v45)+144))
	if v55 == int32(0) {
		goto L3
	} else {
		goto L26
	}
L26:
	;
	v69 = v50
	goto L1
L27:
	;
	return int32(1)
}
func F_remove_rel_from_joinlist(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	v4 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	if l0 == v4 {
		v76 = v4
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v11 + int32(16)
	return v76
L2:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v15 <= int32(0) {
		v76 = v4
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v22 = v4
	v24 = v4
	goto L4
L4:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v26+v24<<(uint(int32(2))%32))))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	if v31 != int32(1) {
		goto L9
	} else {
		goto L10
	}
L5:
	;
	v76 = v67
	goto L1
L6:
	;
	v69 = v24 + int32(1)
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v69 < v70 {
		v22 = v67
		v24 = v69
		goto L4
	} else {
		goto L22
	}
L7:
	;
	v64 = F_lappend(m, v22, v30)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L14
	} else {
		goto L21
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L14
	} else {
		goto L18
	}
L9:
	;
	if v31 != int32(63) {
		goto L8
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v42 = F_remove_rel_from_joinlist(m, v30, l1, l2)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	if v36 != l1 {
		goto L7
	} else {
		goto L13
	}
L13:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v38 + int32(1)
	v67 = v22
	goto L6
L14:
	;
	return int32(0)
L15:
	;
	if v42 == int32(0) {
		v67 = v22
		goto L6
	} else {
		goto L16
	}
L16:
	;
	v48 = F_lappend(m, v22, v42)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L14
	} else {
		goto L17
	}
L17:
	;
	v67 = v48
	goto L6
L18:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v54
	F_errmsg_internal(m, int32(_a_F_remove_rel_from_joinlist_0), v11)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L14
	} else {
		goto L19
	}
L19:
	;
	F_errfinish(m, int32(_a_F_remove_rel_from_joinlist_1), int32(822), int32(_a_F_remove_rel_from_joinlist_2))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L14
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
	v67 = v64
	goto L6
L22:
	;
	goto L5
}
