package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_generic_desc(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+68))
	if v13 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v12)+64))
	v15 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v14))))
	v16 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v14)+2)))
	v18 = v16 + int32(4)
	if base.Ui32(v18) < base.Ui32(v13) {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	m.G0 = v10 + int32(32)
	return
L4:
	;
	v23 = v14 + v18
	v26 = v15
	v27 = v16
	goto L7
L5:
	;
	v46 = v15
	v47 = v16
	goto L6
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v46
	F_appendStringInfo(m, l0, int32(_a_F_generic_desc_0), v10)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L9
	} else {
		goto L12
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v27
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v26
	F_appendStringInfo(m, l0, int32(_a_F_generic_desc_1), v10+int32(16))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v46 = v36
	v47 = v37
	goto L6
L9:
	;
	return
L10:
	;
	v36 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23))))
	v37 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+2)))
	v40 = v23 + v37 + int32(4)
	if base.Ui32(v40) < base.Ui32(v14+v13) {
		v23 = v40
		v26 = v36
		v27 = v37
		goto L7
	} else {
		goto L11
	}
L11:
	;
	goto L8
L12:
	;
	goto L3
}
func F_transformGenericOptions(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v152 int32
	_ = v152
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v189 int32
	_ = v189
	var v194 int32
	_ = v194
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v309 int32
	_ = v309
	var v314 int32
	_ = v314
	v11 = m.G0
	v13 = v11 - int32(96)
	m.G0 = v13
	v15 = F_untransformRelOptions(m, l1)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	if l2 == int32(0) {
		v202 = v15
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v205 = int32(0)
	if v202 == v205 {
		v285 = v205
		goto L58
	} else {
		goto L59
	}
L4:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v21 <= int32(0) {
		v202 = v15
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v31 = v15
	v32 = int32(0)
	goto L9
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L1
	} else {
		goto L54
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L1
	} else {
		goto L50
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L1
	} else {
		goto L46
	}
L9:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v34+v32<<(uint(int32(2))%32))))
	if v31 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L1
	} else {
		goto L42
	}
L11:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v38)+16))
	switch v105 {
	case 0, 2:
		goto L33
	case 1:
		goto L34
	case 3:
		goto L35
	default:
		goto L6
	}
L12:
	;
	v100 = int32(0)
	goto L11
L13:
	;
	goto L14
L14:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	v43 = int32(0)
	if v43 < v42 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v46 = v42
	goto L17
L16:
	;
	v46 = v43
	goto L17
L17:
	;
	v49 = int32(0)
	goto L18
L18:
	;
	if v49 == v46 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v100 = v65
	goto L11
L20:
	;
	v100 = int32(0)
	goto L11
L21:
	;
	goto L22
L22:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v31)+12))
	v65 = v49<<(uint(int32(2))%32) + v64
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v66)+8))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v38)+8))
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67))))
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68))))
	if base.B2i32(v71 == int32(0))|base.B2i32(v71 != v74) != 0 {
		v92 = v71
		v93 = v74
		goto L24
	} else {
		goto L25
	}
L23:
	;
	if v92-v93 != 0 {
		v49 = v49 + int32(1)
		goto L18
	} else {
		goto L30
	}
L24:
	;
	goto L23
L25:
	;
	v77 = v67
	v78 = v68
	goto L26
L26:
	;
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78)+1)))
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77)+1)))
	if v82 == int32(0) {
		v92 = v82
		v93 = v81
		goto L24
	} else {
		goto L28
	}
L27:
	;
	v92 = v82
	v93 = v81
	goto L24
L28:
	;
	v85 = int32(1)
	if v82 == v81 {
		v77 = v77 + v85
		v78 = v78 + v85
		goto L26
	} else {
		goto L29
	}
L29:
	;
	goto L27
L30:
	;
	goto L19
L31:
	;
	goto L10
L32:
	;
	v117 = v32 + int32(1)
	v118 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v117 < v118 {
		v31 = v115
		v32 = v117
		goto L9
	} else {
		goto L41
	}
L33:
	;
	if v100 != 0 {
		goto L7
	} else {
		goto L39
	}
L34:
	;
	if v100 == int32(0) {
		goto L8
	} else {
		goto L38
	}
L35:
	;
	if v100 == int32(0) {
		goto L31
	} else {
		goto L36
	}
L36:
	;
	v108 = F_list_delete_cell(m, v31, v100)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	v115 = v108
	goto L32
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v100))) = v38
	v115 = v31
	goto L32
L39:
	;
	v113 = F_lappend(m, v31, v38)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	v115 = v113
	goto L32
L41:
	;
	v202 = v115
	goto L3
L42:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v38)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+48)) = v127
	F_errmsg(m, int32(_a_F_transformGenericOptions_0), v13+int32(48))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	F_errfinish(m, int32(_a_F_transformGenericOptions_1), int32(160), int32(_a_F_transformGenericOptions_2))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L46:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v38)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+64)) = v146
	F_errmsg(m, int32(_a_F_transformGenericOptions_0), v13-int32(-64))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	F_errfinish(m, int32(_a_F_transformGenericOptions_1), int32(169), int32(_a_F_transformGenericOptions_2))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L50:
	;
	F_errcode(m, int32(_a_F_transformGenericOptions_3))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v38)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+80)) = v165
	F_errmsg(m, int32(_a_F_transformGenericOptions_4), v13+int32(80))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	F_errfinish(m, int32(_a_F_transformGenericOptions_1), int32(179), int32(_a_F_transformGenericOptions_2))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L54:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v38)+16))
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v38)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+36)) = v182
	*(*int32)(unsafe.Add(mBase, uint32(v13)+32)) = v181
	F_errmsg_internal(m, int32(_a_F_transformGenericOptions_5), v13+int32(32))
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	F_errfinish(m, int32(_a_F_transformGenericOptions_1), int32(185), int32(_a_F_transformGenericOptions_2))
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L57:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L1
	} else {
		goto L83
	}
L58:
	;
	if l3 != 0 {
		goto L75
	} else {
		goto L76
	}
L59:
	;
	v208 = int32(0)
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v202)+4))
	if v209 <= v208 {
		v285 = v208
		goto L58
	} else {
		goto L60
	}
L60:
	;
	v218 = v208
	v221 = int32(0)
	goto L61
L61:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v202)+12))
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v223+v218<<(uint(int32(2))%32))))
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v227)+8))
	v229 = F_defGetString(m, v227)
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L1
	} else {
		goto L63
	}
L62:
	;
	v268 = int32(0)
	if v262 == v268 {
		v285 = v268
		goto L58
	} else {
		goto L73
	}
L63:
	;
	v231 = int32(61)
	v232 = F___strchrnul(m, v228, v231)
	mBase = m.M
	v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v232))))
	if v234 == v231 {
		goto L65
	} else {
		goto L66
	}
L64:
	;
	if v238 != 0 {
		goto L57
	} else {
		goto L68
	}
L65:
	;
	v238 = v232
	goto L67
L66:
	;
	v238 = int32(0)
	goto L67
L67:
	;
	goto L64
L68:
	;
	v239 = F_strlen(m, v228)
	mBase = m.M
	v240 = F_strlen(m, v229)
	mBase = m.M
	v241 = v239 + v240
	v244 = F_palloc(m, v241+int32(6))
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L1
	} else {
		goto L69
	}
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v244))) = v241<<(uint(int32(2))%32) + int32(20)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v229
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v228
	v256 = F_pg_sprintf(m, v244+int32(4), int32(_a_F_transformGenericOptions_6), v13)
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L1
	} else {
		goto L70
	}
L70:
	;
	v261 = *(*int32)(unsafe.Add(mBase, _c_F_transformGenericOptions[0]))
	v262 = F_accumArrayResult(m, v221, v244, int32(0), int32(25), v261)
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L1
	} else {
		goto L71
	}
L71:
	;
	v265 = v218 + int32(1)
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v202)+4))
	if v265 < v266 {
		v218 = v265
		v221 = v262
		goto L61
	} else {
		goto L72
	}
L72:
	;
	goto L62
L73:
	;
	v272 = *(*int32)(unsafe.Add(mBase, _c_F_transformGenericOptions[0]))
	v273 = F_makeArrayResult(m, v262, v272)
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L1
	} else {
		goto L74
	}
L74:
	;
	v285 = v273
	goto L58
L75:
	;
	if v285 != 0 {
		goto L78
	} else {
		goto L79
	}
L76:
	;
	goto L77
L77:
	;
	m.G0 = v13 + int32(96)
	return v285
L78:
	;
	v290 = v285
	goto L80
L79:
	;
	v288 = F_construct_empty_array(m, int32(25))
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L1
	} else {
		goto L81
	}
L80:
	;
	v291 = F_OidFunctionCall2Coll(m, l3, int32(0), v290, l0)
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L1
	} else {
		goto L82
	}
L81:
	;
	v290 = v288
	goto L80
L82:
	;
	goto L77
L83:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L1
	} else {
		goto L84
	}
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v228
	F_errmsg(m, int32(_a_F_transformGenericOptions_7), v13+int32(16))
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	F_errfinish(m, int32(_a_F_transformGenericOptions_1), int32(87), int32(_a_F_transformGenericOptions_8))
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L1
	} else {
		goto L86
	}
L86:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
