package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_SpecialHyphen(m *base.Module, l0 int32) {
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
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v3)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v3))) = v4 - v5
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v8)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v9 - v10
	return
}
func F_resolve_special_varno(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v257 int32
	_ = v257
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v268 int32
	_ = v268
	var v273 int32
	_ = v273
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v284 int32
	_ = v284
	var v289 int32
	_ = v289
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v300 int32
	_ = v300
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v316 int32
	_ = v316
	v10 = m.G0
	v12 = v10 - int32(144)
	m.G0 = v12
	F_check_stack_depth(m)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v16 != int32(6) {
		v306 = l0
		goto L4
	} else {
		goto L5
	}
L3:
	;
	m.G0 = v12 + int32(144)
	return
L4:
	;
	m.T0[l2].(func(*base.Module, int32, int32, int32))(m, v306, l1, l3)
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L1
	} else {
		goto L92
	}
L5:
	;
	v19 = l0
	goto L10
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L1
	} else {
		goto L89
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L1
	} else {
		goto L86
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L1
	} else {
		goto L83
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L1
	} else {
		goto L80
	}
L10:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+12))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v19)+28))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v29+v30<<(uint(int32(2))%32))))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	if v35 != int32(-3) {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	if v35 <= int32(0) {
		goto L9
	} else {
		goto L77
	}
L12:
	;
	goto L11
L13:
	;
	switch v35 + int32(2) {
	case 0:
		goto L17
	case 1:
		goto L16
	default:
		goto L12
	}
L14:
	;
	goto L15
L15:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v34)+64))
	if v184 == int32(0) {
		goto L9
	} else {
		goto L60
	}
L16:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v34)+60))
	if v118 == int32(0) {
		goto L9
	} else {
		goto L41
	}
L17:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v34)+56))
	if v40 == int32(0) {
		goto L9
	} else {
		goto L18
	}
L18:
	;
	v43 = int32(*(*int16)(unsafe.Add(mBase, uint32(v19)+8)))
	if v40 != 0 {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	if v81 == int32(0) {
		goto L8
	} else {
		goto L32
	}
L20:
	;
	goto L19
L21:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
	if v47 <= int32(0) {
		v81 = int32(0)
		goto L20
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v81 = int32(0)
	goto L20
L24:
	;
	v50 = int32(0)
	if v50 < v47 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v53 = v47
	goto L27
L26:
	;
	v53 = v50
	goto L27
L27:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v40)+12))
	v58 = int32(0)
	goto L28
L28:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v54+v58<<(uint(int32(2))%32))))
	v67 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v66)+8)))
	if v67 == v43&int32(_a_F_resolve_special_varno_0) {
		v81 = v66
		goto L20
	} else {
		goto L30
	}
L29:
	;
	goto L23
L30:
	;
	v70 = v58 + int32(1)
	if v70 != v53 {
		v58 = v70
		goto L28
	} else {
		goto L31
	}
L31:
	;
	goto L29
L32:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v34)+40))
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
	if v87&int32(-2) == int32(334) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v86)+72))
	v93 = F_bms_union(m, v85, v92)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L1
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v34)+48))
	v98 = v12 - int32(-64)
	base.MemoryCopy(m, v98, v34, int32(80))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v34)+40))
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v34)+44))
	v103 = F_lcons(m, v101, v102)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L1
	} else {
		goto L37
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = v93
	goto L35
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+44)) = v103
	F_set_deparse_plan(m, v34, v96)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v81)+4))
	F_resolve_special_varno(m, v108, l1, l2, l3)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v34)+44))
	v112 = F_list_delete_first(m, v111)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	base.MemoryCopy(m, v34, v98, int32(80))
	*(*int32)(unsafe.Add(mBase, uint32(v34)+44)) = v112
	*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = v85
	goto L3
L41:
	;
	v121 = int32(*(*int16)(unsafe.Add(mBase, uint32(v19)+8)))
	if v118 != 0 {
		goto L44
	} else {
		goto L45
	}
L42:
	;
	if v159 == int32(0) {
		goto L7
	} else {
		goto L55
	}
L43:
	;
	goto L42
L44:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v118)+4))
	if v125 <= int32(0) {
		v159 = int32(0)
		goto L43
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	v159 = int32(0)
	goto L43
L47:
	;
	v128 = int32(0)
	if v128 < v125 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v131 = v125
	goto L50
L49:
	;
	v131 = v128
	goto L50
L50:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v118)+12))
	v136 = int32(0)
	goto L51
L51:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v132+v136<<(uint(int32(2))%32))))
	v145 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v144)+8)))
	if v145 == v121&int32(_a_F_resolve_special_varno_0) {
		v159 = v144
		goto L43
	} else {
		goto L53
	}
L52:
	;
	goto L46
L53:
	;
	v148 = v136 + int32(1)
	if v148 != v131 {
		v136 = v148
		goto L51
	} else {
		goto L54
	}
L54:
	;
	goto L52
L55:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v34)+52))
	v165 = v12 - int32(-64)
	base.MemoryCopy(m, v165, v34, int32(80))
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v34)+40))
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v34)+44))
	v170 = F_lcons(m, v168, v169)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+44)) = v170
	F_set_deparse_plan(m, v34, v163)
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v159)+4))
	F_resolve_special_varno(m, v175, l1, l2, l3)
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v34)+44))
	v179 = F_list_delete_first(m, v178)
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	base.MemoryCopy(m, v34, v165, int32(80))
	*(*int32)(unsafe.Add(mBase, uint32(v34)+44)) = v179
	goto L3
L60:
	;
	v187 = int32(*(*int16)(unsafe.Add(mBase, uint32(v19)+8)))
	if v184 != 0 {
		goto L63
	} else {
		goto L64
	}
L61:
	;
	if v225 == int32(0) {
		goto L6
	} else {
		goto L74
	}
L62:
	;
	goto L61
L63:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v184)+4))
	if v191 <= int32(0) {
		v225 = int32(0)
		goto L62
	} else {
		goto L66
	}
L64:
	;
	goto L65
L65:
	;
	v225 = int32(0)
	goto L62
L66:
	;
	v194 = int32(0)
	if v194 < v191 {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v197 = v191
	goto L69
L68:
	;
	v197 = v194
	goto L69
L69:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v184)+12))
	v202 = int32(0)
	goto L70
L70:
	;
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v198+v202<<(uint(int32(2))%32))))
	v211 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v210)+8)))
	if v211 == v187&int32(_a_F_resolve_special_varno_0) {
		v225 = v210
		goto L62
	} else {
		goto L72
	}
L71:
	;
	goto L65
L72:
	;
	v214 = v202 + int32(1)
	if v214 != v197 {
		v202 = v214
		goto L70
	} else {
		goto L73
	}
L73:
	;
	goto L71
L74:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v225)+4))
	F_check_stack_depth(m)
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L1
	} else {
		goto L75
	}
L75:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v229)))
	if v232 == int32(6) {
		v19 = v229
		goto L10
	} else {
		goto L76
	}
L76:
	;
	v306 = v229
	goto L4
L77:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
	if v237 == int32(0) {
		goto L9
	} else {
		goto L78
	}
L78:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v237)+4))
	if v35 <= v240 {
		v306 = v19
		goto L4
	} else {
		goto L79
	}
L79:
	;
	goto L9
L80:
	;
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v248
	F_errmsg_internal(m, int32(_a_F_resolve_special_varno_1), v12)
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	F_errfinish(m, int32(_a_F_resolve_special_varno_2), int32(_a_F_resolve_special_varno_3), int32(_a_F_resolve_special_varno_4))
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L1
	} else {
		goto L82
	}
L82:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L83:
	;
	v262 = int32(*(*int16)(unsafe.Add(mBase, uint32(v19)+8)))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v262
	F_errmsg_internal(m, int32(_a_F_resolve_special_varno_5), v12+int32(16))
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L1
	} else {
		goto L84
	}
L84:
	;
	F_errfinish(m, int32(_a_F_resolve_special_varno_2), int32(_a_F_resolve_special_varno_6), int32(_a_F_resolve_special_varno_4))
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L1
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
	v278 = int32(*(*int16)(unsafe.Add(mBase, uint32(v19)+8)))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v278
	F_errmsg_internal(m, int32(_a_F_resolve_special_varno_7), v12+int32(32))
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L1
	} else {
		goto L87
	}
L87:
	;
	F_errfinish(m, int32(_a_F_resolve_special_varno_2), int32(_a_F_resolve_special_varno_8), int32(_a_F_resolve_special_varno_4))
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L1
	} else {
		goto L88
	}
L88:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L89:
	;
	v294 = int32(*(*int16)(unsafe.Add(mBase, uint32(v19)+8)))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = v294
	F_errmsg_internal(m, int32(_a_F_resolve_special_varno_9), v12+int32(48))
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L1
	} else {
		goto L90
	}
L90:
	;
	F_errfinish(m, int32(_a_F_resolve_special_varno_2), int32(_a_F_resolve_special_varno_10), int32(_a_F_resolve_special_varno_4))
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L1
	} else {
		goto L91
	}
L91:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L92:
	;
	goto L3
}
