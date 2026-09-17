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
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v201 int32
	_ = v201
	var v210 int32
	_ = v210
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v227 int32
	_ = v227
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v250 int32
	_ = v250
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v271 int32
	_ = v271
	var v276 int32
	_ = v276
	v3 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(128)
	m.G0 = v10
	if l1 != 0 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L8
	} else {
		goto L77
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L8
	} else {
		goto L74
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L8
	} else {
		goto L70
	}
L4:
	;
	v15 = F_find_option(m, l0, int32(0), int32(1), int32(19))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L8
	} else {
		goto L9
	}
L5:
	;
	v210 = int32(0)
	goto L6
L6:
	;
	m.G0 = v10 + int32(128)
	return v210
L7:
	;
	F_initStringInfo(m, v10+int32(112))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L8
	} else {
		goto L15
	}
L8:
	;
	return int32(0)
L9:
	;
	if v15 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v15)+20))
	if v19&int32(1) != 0 {
		v26 = v19
		goto L7
	} else {
		goto L13
	}
L11:
	;
	v22 = v3
	goto L12
L12:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v23 != int32(1) {
		goto L3
	} else {
		goto L14
	}
L13:
	;
	v22 = v19
	goto L12
L14:
	;
	v26 = v22
	goto L7
L15:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v31 <= int32(0) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v10)+112))
	v210 = v201
	goto L6
L17:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	if v36 == int32(73) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v35)+8))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
	v42 = v40
	v43 = v39
	v44 = v41
	goto L20
L19:
	;
	v42 = v35
	v43 = v3
	v44 = v36
	goto L20
L20:
	;
	if v44 != int32(72) {
		v233 = v42
		goto L2
	} else {
		goto L21
	}
L21:
	;
	v48 = v26 & int32(2)
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
	switch v49 - int32(465) {
	case 0:
		goto L23
	case 1:
		goto L24
	default:
		v256 = v42
		goto L1
	case 3:
		goto L25
	}
L22:
	;
	v103 = int32(1)
	v104 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v104 <= v103 {
		goto L16
	} else {
		goto L41
	}
L23:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v42)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+64)) = v93
	F_appendStringInfo(m, v10+int32(112), int32(_a_F_flatten_set_variable_args_0), v10-int32(-64))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L8
	} else {
		goto L40
	}
L24:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v42)+8))
	F_appendStringInfoString(m, v10+int32(112), v90)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L8
	} else {
		goto L39
	}
L25:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v42)+8))
	if v43 != 0 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	F_typenameTypeIdAndMod(m, int32(0), v43, v10+int32(108), v10+int32(104))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L8
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	if v48 != 0 {
		goto L33
	} else {
		goto L34
	}
L29:
	;
	v61 = int32(0)
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v10)+104))
	v66 = F_DirectFunctionCall3Coll(m, int32(585), v61, v52, v61, v65)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L8
	} else {
		goto L30
	}
L30:
	;
	v68 = F_DirectFunctionCall1Coll(m, int32(1273), v61, v66)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L8
	} else {
		goto L31
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+80)) = v68
	F_appendStringInfo(m, v10+int32(112), int32(_a_F_flatten_set_variable_args_1), v10+int32(80))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L8
	} else {
		goto L32
	}
L32:
	;
	goto L22
L33:
	;
	v80 = F_quote_identifier(m, v52)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L8
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	F_appendStringInfoString(m, v10+int32(112), v52)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L8
	} else {
		goto L38
	}
L36:
	;
	F_appendStringInfoString(m, v10+int32(112), v80)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L8
	} else {
		goto L37
	}
L37:
	;
	goto L22
L38:
	;
	goto L22
L39:
	;
	goto L22
L40:
	;
	goto L22
L41:
	;
	v112 = v103
	goto L42
L42:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v114+v112<<(uint(int32(2))%32))))
	F_appendStringInfoString(m, v10+int32(112), int32(_a_F_flatten_set_variable_args_2))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L8
	} else {
		goto L44
	}
L43:
	;
	goto L16
L44:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v118)))
	if v124 != int32(73) {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	if v133 != int32(72) {
		v233 = v131
		goto L2
	} else {
		goto L49
	}
L46:
	;
	v131 = v118
	v132 = int32(0)
	v133 = v124
	goto L45
L47:
	;
	goto L48
L48:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v118)+8))
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v118)+4))
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v129)))
	v131 = v129
	v132 = v128
	v133 = v130
	goto L45
L49:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v131)+4))
	switch v136 - int32(465) {
	case 0:
		goto L51
	case 1:
		goto L53
	default:
		v256 = v131
		goto L1
	case 3:
		goto L52
	}
L50:
	;
	v191 = v112 + int32(1)
	v192 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v191 < v192 {
		v112 = v191
		goto L42
	} else {
		goto L69
	}
L51:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v131)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v180
	F_appendStringInfo(m, v10+int32(112), int32(_a_F_flatten_set_variable_args_0), v10+int32(16))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L8
	} else {
		goto L68
	}
L52:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v131)+8))
	if v132 != 0 {
		goto L55
	} else {
		goto L56
	}
L53:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v131)+8))
	F_appendStringInfoString(m, v10+int32(112), v141)
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L8
	} else {
		goto L54
	}
L54:
	;
	goto L50
L55:
	;
	F_typenameTypeIdAndMod(m, int32(0), v132, v10+int32(108), v10+int32(104))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L8
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	if v48 != 0 {
		goto L62
	} else {
		goto L63
	}
L58:
	;
	v153 = int32(0)
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v10)+104))
	v158 = F_DirectFunctionCall3Coll(m, int32(585), v153, v144, v153, v157)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L8
	} else {
		goto L59
	}
L59:
	;
	v160 = F_DirectFunctionCall1Coll(m, int32(1273), v153, v158)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L8
	} else {
		goto L60
	}
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = v160
	F_appendStringInfo(m, v10+int32(112), int32(_a_F_flatten_set_variable_args_1), v10+int32(32))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L8
	} else {
		goto L61
	}
L61:
	;
	goto L50
L62:
	;
	v172 = F_quote_identifier(m, v144)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L8
	} else {
		goto L65
	}
L63:
	;
	goto L64
L64:
	;
	F_appendStringInfoString(m, v10+int32(112), v144)
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L8
	} else {
		goto L67
	}
L65:
	;
	F_appendStringInfoString(m, v10+int32(112), v172)
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L8
	} else {
		goto L66
	}
L66:
	;
	goto L50
L67:
	;
	goto L50
L68:
	;
	goto L50
L69:
	;
	goto L43
L70:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L8
	} else {
		goto L71
	}
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+96)) = l0
	F_errmsg(m, int32(_a_F_flatten_set_variable_args_3), v10+int32(96))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L8
	} else {
		goto L72
	}
L72:
	;
	F_errfinish(m, int32(_a_F_flatten_set_variable_args_4), int32(218), int32(_a_F_flatten_set_variable_args_5))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L8
	} else {
		goto L73
	}
L73:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L74:
	;
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v233)))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+48)) = v244
	F_errmsg_internal(m, int32(_a_F_flatten_set_variable_args_6), v10+int32(48))
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L8
	} else {
		goto L75
	}
L75:
	;
	F_errfinish(m, int32(_a_F_flatten_set_variable_args_4), int32(246), int32(_a_F_flatten_set_variable_args_5))
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L8
	} else {
		goto L76
	}
L76:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L77:
	;
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v256)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v267
	F_errmsg_internal(m, int32(_a_F_flatten_set_variable_args_6), v10)
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L8
	} else {
		goto L78
	}
L78:
	;
	F_errfinish(m, int32(_a_F_flatten_set_variable_args_4), int32(300), int32(_a_F_flatten_set_variable_args_5))
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L8
	} else {
		goto L79
	}
L79:
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
				v23 = *(*int32)(unsafe.Add(mBase, uint32(v20)+76))
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
			v72 = *(*float64)(unsafe.Add(mBase, uint32(v65)+120))
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
