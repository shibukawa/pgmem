package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ExtractSetVariableArgs(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	switch v3 {
	case 0:
		v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v6 = F_flatten_set_variable_args(m, v4, v5)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			return v6
		}
	default:
		v16 = int32(0)
		return v16
	case 2:
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v12 = int32(0)
		v14 = F_GetConfigOptionByName(m, v11, v12, v12)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			v16 = v14
			return v16
		}
	}
}
func F_flatten_set_variable_args(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v90 int64
	_ = v90
	var v91 int64
	_ = v91
	var v92 int32
	_ = v92
	var v93 int64
	_ = v93
	var v94 int32
	_ = v94
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v187 int64
	_ = v187
	var v188 int64
	_ = v188
	var v189 int32
	_ = v189
	var v190 int64
	_ = v190
	var v191 int32
	_ = v191
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v232 int32
	_ = v232
	var v241 int32
	_ = v241
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v258 int32
	_ = v258
	var v263 int32
	_ = v263
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
	var v299 int32
	_ = v299
	var v302 int32
	_ = v302
	var v306 int32
	_ = v306
	var v311 int32
	_ = v311
	var v315 int32
	_ = v315
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v330 int32
	_ = v330
	var v335 int32
	_ = v335
	v3 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(144)
	m.G0 = v11
	if l1 == v3 {
		v241 = v3
		goto L5
	} else {
		goto L6
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L9
	} else {
		goto L86
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L9
	} else {
		goto L82
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L9
	} else {
		goto L79
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L9
	} else {
		goto L75
	}
L5:
	;
	m.G0 = v11 + int32(144)
	return v241
L6:
	;
	v19 = F_find_option(m, l0, int32(0), int32(1), int32(19))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	F_initStringInfo(m, v11+int32(128))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L9
	} else {
		goto L18
	}
L8:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v45 != int32(1) {
		goto L4
	} else {
		goto L17
	}
L9:
	;
	return int32(0)
L10:
	;
	if v19 == int32(0) {
		v44 = v3
		goto L8
	} else {
		goto L11
	}
L11:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v19)+20))
	if v25&int32(1) == int32(0) {
		v44 = v25
		goto L8
	} else {
		goto L12
	}
L12:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v30 != int32(1) {
		v49 = v25
		goto L7
	} else {
		goto L13
	}
L13:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
	if v35 != int32(72) {
		v49 = v25
		goto L7
	} else {
		goto L14
	}
L14:
	;
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+12)))
	if v38 != int32(1) {
		v49 = v25
		goto L7
	} else {
		goto L15
	}
L15:
	;
	v42 = F_pstrdup(m, int32(_a_F_flatten_set_variable_args_0))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L9
	} else {
		goto L16
	}
L16:
	;
	v241 = v42
	goto L5
L17:
	;
	v49 = v44
	goto L7
L18:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v54 <= int32(0) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v11)+128))
	v241 = v232
	goto L5
L20:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)))
	if v59 == int32(73) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v58)+8))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v58)+4))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
	v65 = v63
	v66 = v62
	v67 = v64
	goto L23
L22:
	;
	v65 = v58
	v66 = v3
	v67 = v59
	goto L23
L23:
	;
	if v67 != int32(72) {
		v267 = v65
		goto L3
	} else {
		goto L24
	}
L24:
	;
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65)+12)))
	if v70 != 0 {
		goto L2
	} else {
		goto L25
	}
L25:
	;
	v72 = v49 & int32(2)
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v65)+4))
	switch v73 - int32(473) {
	case 0:
		goto L27
	case 1:
		goto L28
	default:
		v315 = v65
		goto L1
	case 3:
		goto L29
	}
L26:
	;
	v128 = int32(1)
	v129 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v129 <= v128 {
		goto L19
	} else {
		goto L45
	}
L27:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v65)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+80)) = v118
	F_appendStringInfo(m, v11+int32(128), int32(_a_F_flatten_set_variable_args_1), v11+int32(80))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L9
	} else {
		goto L44
	}
L28:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v65)+8))
	F_appendStringInfoString(m, v11+int32(128), v115)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L9
	} else {
		goto L43
	}
L29:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v65)+8))
	if v66 != 0 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	F_typenameTypeIdAndMod(m, int32(0), v66, v11+int32(124), v11+int32(120))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L9
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	if v72 != 0 {
		goto L37
	} else {
		goto L38
	}
L33:
	;
	v85 = int32(0)
	v90 = int64(*(*int32)(unsafe.Add(mBase, uint32(v11)+120)))
	v91 = F_DirectFunctionCall3Coll(m, int32(631), v85, base.I64_extend_i32_u(v76), int64(0), v90)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L9
	} else {
		goto L34
	}
L34:
	;
	v93 = F_DirectFunctionCall1Coll(m, int32(1401), v85, v91)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L9
	} else {
		goto L35
	}
L35:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v11)+96)) = uint32(v93)
	F_appendStringInfo(m, v11+int32(128), int32(_a_F_flatten_set_variable_args_2), v11+int32(96))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L9
	} else {
		goto L36
	}
L36:
	;
	goto L26
L37:
	;
	v105 = F_quote_identifier(m, v76)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L9
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	F_appendStringInfoString(m, v11+int32(128), v76)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L9
	} else {
		goto L42
	}
L40:
	;
	F_appendStringInfoString(m, v11+int32(128), v105)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L9
	} else {
		goto L41
	}
L41:
	;
	goto L26
L42:
	;
	goto L26
L43:
	;
	goto L26
L44:
	;
	goto L26
L45:
	;
	v137 = v128
	goto L46
L46:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v140+v137<<(uint(int32(2))%32))))
	F_appendStringInfoString(m, v11+int32(128), int32(_a_F_flatten_set_variable_args_3))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L9
	} else {
		goto L48
	}
L47:
	;
	goto L19
L48:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v144)))
	if v150 != int32(73) {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	if v159 != int32(72) {
		v267 = v157
		goto L3
	} else {
		goto L53
	}
L50:
	;
	v157 = v144
	v158 = int32(0)
	v159 = v150
	goto L49
L51:
	;
	goto L52
L52:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v144)+8))
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v144)+4))
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v155)))
	v157 = v155
	v158 = v154
	v159 = v156
	goto L49
L53:
	;
	v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157)+12)))
	if v162 == int32(1) {
		goto L2
	} else {
		goto L54
	}
L54:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v157)+4))
	switch v165 - int32(473) {
	case 0:
		goto L56
	case 1:
		goto L58
	default:
		v315 = v157
		goto L1
	case 3:
		goto L57
	}
L55:
	;
	v221 = v137 + int32(1)
	v222 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v221 < v222 {
		v137 = v221
		goto L46
	} else {
		goto L74
	}
L56:
	;
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v157)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v210
	F_appendStringInfo(m, v11+int32(128), int32(_a_F_flatten_set_variable_args_1), v11+int32(32))
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L9
	} else {
		goto L73
	}
L57:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v157)+8))
	if v158 != 0 {
		goto L60
	} else {
		goto L61
	}
L58:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v157)+8))
	F_appendStringInfoString(m, v11+int32(128), v170)
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L9
	} else {
		goto L59
	}
L59:
	;
	goto L55
L60:
	;
	F_typenameTypeIdAndMod(m, int32(0), v158, v11+int32(124), v11+int32(120))
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L9
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	if v72 != 0 {
		goto L67
	} else {
		goto L68
	}
L63:
	;
	v182 = int32(0)
	v187 = int64(*(*int32)(unsafe.Add(mBase, uint32(v11)+120)))
	v188 = F_DirectFunctionCall3Coll(m, int32(631), v182, base.I64_extend_i32_u(v173), int64(0), v187)
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L9
	} else {
		goto L64
	}
L64:
	;
	v190 = F_DirectFunctionCall1Coll(m, int32(1401), v182, v188)
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L9
	} else {
		goto L65
	}
L65:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v11)+48)) = uint32(v190)
	F_appendStringInfo(m, v11+int32(128), int32(_a_F_flatten_set_variable_args_2), v11+int32(48))
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L9
	} else {
		goto L66
	}
L66:
	;
	goto L55
L67:
	;
	v202 = F_quote_identifier(m, v173)
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L9
	} else {
		goto L70
	}
L68:
	;
	goto L69
L69:
	;
	F_appendStringInfoString(m, v11+int32(128), v173)
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L9
	} else {
		goto L72
	}
L70:
	;
	F_appendStringInfoString(m, v11+int32(128), v202)
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L9
	} else {
		goto L71
	}
L71:
	;
	goto L55
L72:
	;
	goto L55
L73:
	;
	goto L55
L74:
	;
	goto L47
L75:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L9
	} else {
		goto L76
	}
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+112)) = l0
	F_errmsg(m, int32(_a_F_flatten_set_variable_args_4), v11+int32(112))
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L9
	} else {
		goto L77
	}
L77:
	;
	F_errfinish(m, int32(_a_F_flatten_set_variable_args_5), int32(236), int32(_a_F_flatten_set_variable_args_6))
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L9
	} else {
		goto L78
	}
L78:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L79:
	;
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v267)))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+64)) = v276
	F_errmsg_internal(m, int32(_a_F_flatten_set_variable_args_7), v11-int32(-64))
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L9
	} else {
		goto L80
	}
L80:
	;
	F_errfinish(m, int32(_a_F_flatten_set_variable_args_5), int32(265), int32(_a_F_flatten_set_variable_args_6))
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L9
	} else {
		goto L81
	}
L81:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L82:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L9
	} else {
		goto L83
	}
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = l0
	F_errmsg(m, int32(_a_F_flatten_set_variable_args_8), v11)
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L9
	} else {
		goto L84
	}
L84:
	;
	F_errfinish(m, int32(_a_F_flatten_set_variable_args_5), int32(272), int32(_a_F_flatten_set_variable_args_6))
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L9
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
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v315)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v324
	F_errmsg_internal(m, int32(_a_F_flatten_set_variable_args_7), v11+int32(16))
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L9
	} else {
		goto L87
	}
L87:
	;
	F_errfinish(m, int32(_a_F_flatten_set_variable_args_5), int32(328), int32(_a_F_flatten_set_variable_args_6))
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L9
	} else {
		goto L88
	}
L88:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_get_variable_numdistinct(m *base.Module, l0 int32, l1 int32) float64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 float64
	_ = v4
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 float32
	_ = v12
	var v14 float32
	_ = v14
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v41 float64
	_ = v41
	var v42 float64
	_ = v42
	var v46 int32
	_ = v46
	var v47 float64
	_ = v47
	var v50 float64
	_ = v50
	var v59 float64
	_ = v59
	var v63 float64
	_ = v63
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v72 float64
	_ = v72
	var v75 int32
	_ = v75
	var v82 float64
	_ = v82
	var v83 float64
	_ = v83
	var v92 float64
	_ = v92
	var v96 float64
	_ = v96
	var v100 float64
	_ = v100
	var v109 float64
	_ = v109
	var v113 float64
	_ = v113
	var v115 int32
	_ = v115
	v3 = int32(0)
	v4 = float64(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v3)
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v8 != 0 {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+16))
		v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+22)))
		v11 = v9 + v10
		v12 = *(*float32)(unsafe.Add(mBase, uint32(v11)+8))
		v14 = *(*float32)(unsafe.Add(mBase, uint32(v11)+16))
		v41 = base.F64_promote_f32(v14)
		v42 = base.F64_promote_f32(v12)
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		if v16 == int32(16) {
			v41 = float64(2)
			v42 = v4
		} else {
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			if v20 == int32(0) {
				v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				if v27 == int32(0) {
					v41 = float64(0)
					v42 = v4
				} else {
					v30 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
					if v30 != int32(6) {
						v41 = float64(0)
						v42 = v4
					} else {
						v34 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v27)+8)))
						switch v34 - int32(_a_F_get_variable_numdistinct_0) {
						case 0:
							v41 = float64(1)
							v42 = v4
						default:
							v41 = float64(0)
							v42 = v4
						case 5:
							v41 = float64(-1)
							v42 = v4
						}
					}
				}
			} else {
				v23 = *(*int32)(unsafe.Add(mBase, uint32(v20)+84))
				if v23 != int32(5) {
					v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					if v27 == int32(0) {
						v41 = float64(0)
						v42 = v4
					} else {
						v30 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
						if v30 != int32(6) {
							v41 = float64(0)
							v42 = v4
						} else {
							v34 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v27)+8)))
							switch v34 - int32(_a_F_get_variable_numdistinct_0) {
							case 0:
								v41 = float64(1)
								v42 = v4
							default:
								v41 = float64(0)
								v42 = v4
							case 5:
								v41 = float64(-1)
								v42 = v4
							}
						}
					}
				} else {
					v41 = float64(-1)
					v42 = v4
				}
			}
		}
	}
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
	if v46 != 0 {
		v47 = base.F64_neg(base.F64_sub(float64(1), v42))
	} else {
		v47 = v41
	}
	if base.F64_gt(v47, float64(0)) != 0 {
		v50 = float64(1e+100)
		if base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v47)&int64(9223372036854775807)))|base.F64_gt(v47, v50) != 0 {
			v63 = v50
		} else {
			v59 = float64(1)
			if base.F64_le(v47, v59) != 0 {
				v63 = v59
			} else {
				v63 = base.F64_nearest(v47)
			}
		}
		return v63
	} else {
		v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		if v65 == int32(0) {
			v68 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v68)
			return float64(200)
		} else {
			v72 = *(*float64)(unsafe.Add(mBase, uint32(v65)+128))
			if base.F64_le(v72, float64(0)) != 0 {
				v75 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v75)
				return float64(200)
			} else {
				if base.F64_lt(v47, float64(0)) != 0 {
					v82 = base.F64_mul(v72, base.F64_neg(v47))
					v83 = float64(1e+100)
					if base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v82)&int64(9223372036854775807)))|base.F64_gt(v82, v83) != 0 {
						v96 = v83
					} else {
						v92 = float64(1)
						if base.F64_le(v82, v92) != 0 {
							v96 = v92
						} else {
							v96 = base.F64_nearest(v82)
						}
					}
					return v96
				} else {
					if base.F64_lt(v72, float64(200)) != 0 {
						v100 = float64(1e+100)
						if base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v72)&int64(9223372036854775807)))|base.F64_gt(v72, v100) != 0 {
							v113 = v100
						} else {
							v109 = float64(1)
							if base.F64_le(v72, v109) != 0 {
								v113 = v109
							} else {
								v113 = base.F64_nearest(v72)
							}
						}
						return v113
					} else {
						v115 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v115)
						return float64(200)
					}
				}
			}
		}
	}
}
