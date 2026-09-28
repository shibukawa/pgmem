package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ExpandColumnRefStar(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
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
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
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
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
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
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v236 int32
	_ = v236
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v274 int64
	_ = v274
	var v280 int32
	_ = v280
	var v285 int32
	_ = v285
	var v294 int32
	_ = v294
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v318 int32
	_ = v318
	v4 = int32(0)
	v11 = m.G0
	v13 = v11 - int32(48)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v15 == v4 {
		v91 = v4
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v13 + int32(48)
	return v318
L2:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	if v92 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L3:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	if v18 != int32(1) {
		v91 = v18
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v22 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L13
	} else {
		goto L18
	}
L6:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	if v25 <= int32(0) {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v28 = int32(0)
	v31 = v28
	v33 = v28
	v34 = v25
	v38 = v4
	goto L8
L8:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v40+v31<<(uint(int32(2))%32))))
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+21)))
	if v45 != 0 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	if v57&int32(1) != 0 {
		v318 = v55
		goto L1
	} else {
		goto L17
	}
L10:
	;
	v48 = F_expandNSItemAttrs(m, l0, v44, int32(0), v21)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	v55 = v33
	v56 = v34
	v57 = v38
	goto L12
L12:
	;
	v59 = v31 + int32(1)
	if v59 < v56 {
		v31 = v59
		v33 = v55
		v34 = v56
		v38 = v57
		goto L8
	} else {
		goto L16
	}
L13:
	;
	return int32(0)
L14:
	;
	v52 = F_list_concat(m, v33, v48)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	v55 = v52
	v56 = v54
	v57 = int32(1)
	goto L12
L16:
	;
	goto L9
L17:
	;
	goto L5
L18:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L13
	} else {
		goto L19
	}
L19:
	;
	F_errmsg(m, int32(_a_F_ExpandColumnRefStar_0), int32(0))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L13
	} else {
		goto L20
	}
L20:
	;
	F_parser_errposition(m, l0, v21)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L13
	} else {
		goto L21
	}
L21:
	;
	F_errfinish(m, int32(_a_F_ExpandColumnRefStar_1), int32(1332), int32(_a_F_ExpandColumnRefStar_2))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L13
	} else {
		goto L22
	}
L22:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L23:
	;
	v102 = int32(2)
	v103 = int32(0)
	switch v91 - v102 {
	case 0:
		goto L32
	case 1:
		goto L31
	case 2:
		goto L30
	default:
		v164 = v103
		v165 = v102
		v166 = v4
		v167 = v4
		goto L28
	}
L24:
	;
	v95 = m.T0[v92].(func(*base.Module, int32, int32) int32)(m, l0, l1)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L13
	} else {
		goto L25
	}
L25:
	;
	if v95 == int32(0) {
		goto L23
	} else {
		goto L26
	}
L26:
	;
	v99 = F_ExpandRowReference(m, l0, v95, l2)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L13
	} else {
		goto L27
	}
L27:
	;
	v318 = v99
	goto L1
L28:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v168 != 0 {
		goto L47
	} else {
		goto L48
	}
L29:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v152)))
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v156)+4))
	v158 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v161 = F_refnameNamespaceItem(m, l0, v154, v157, v158, v13+int32(44))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L13
	} else {
		goto L44
	}
L30:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v113)))
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v114)+4))
	v117 = *(*int32)(unsafe.Add(mBase, _c_F_ExpandColumnRefStar[0]))
	v118 = F_get_database_name(m, v117)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L13
	} else {
		goto L33
	}
L31:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v108)))
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v111)+4))
	v152 = v108 + int32(4)
	v154 = v112
	goto L29
L32:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	v152 = v106
	v154 = int32(0)
	goto L29
L33:
	;
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115))))
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118))))
	if base.B2i32(v122 == int32(0))|base.B2i32(v122 != v125) != 0 {
		v143 = v122
		v144 = v125
		goto L35
	} else {
		goto L36
	}
L34:
	;
	if v143-v144 != 0 {
		goto L41
	} else {
		goto L42
	}
L35:
	;
	goto L34
L36:
	;
	v128 = v115
	v129 = v118
	goto L37
L37:
	;
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129)+1)))
	v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128)+1)))
	if v133 == int32(0) {
		v143 = v133
		v144 = v132
		goto L35
	} else {
		goto L39
	}
L38:
	;
	v143 = v133
	v144 = v132
	goto L35
L39:
	;
	v136 = int32(1)
	if v133 == v132 {
		v128 = v128 + v136
		v129 = v129 + v136
		goto L37
	} else {
		goto L40
	}
L40:
	;
	goto L38
L41:
	;
	v164 = v103
	v165 = int32(1)
	v166 = v4
	v167 = v4
	goto L28
L42:
	;
	goto L43
L43:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v147)+4))
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v150)+4))
	v152 = v147 + int32(8)
	v154 = v151
	goto L29
L44:
	;
	v164 = v161
	v165 = int32(0)
	v166 = v154
	v167 = v157
	goto L28
L45:
	;
	v261 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v13)+44))
	if l2 != 0 {
		goto L85
	} else {
		goto L86
	}
L46:
	;
	if v165 != 0 {
		goto L67
	} else {
		goto L68
	}
L47:
	;
	if v164 == int32(0) {
		goto L50
	} else {
		goto L51
	}
L48:
	;
	goto L49
L49:
	;
	if v164 != 0 {
		goto L45
	} else {
		goto L64
	}
L50:
	;
	v172 = m.T0[v168].(func(*base.Module, int32, int32, int32) int32)(m, l0, l1, int32(0))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L13
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v164)+4))
	v179 = m.T0[v168].(func(*base.Module, int32, int32, int32) int32)(m, l0, l1, v178)
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L13
	} else {
		goto L56
	}
L53:
	;
	if v172 == int32(0) {
		goto L46
	} else {
		goto L54
	}
L54:
	;
	v176 = F_ExpandRowReference(m, l0, v172, l2)
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L13
	} else {
		goto L55
	}
L55:
	;
	v318 = v176
	goto L1
L56:
	;
	if v179 == int32(0) {
		goto L45
	} else {
		goto L57
	}
L57:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L13
	} else {
		goto L58
	}
L58:
	;
	F_errcode(m, int32(33583236))
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L13
	} else {
		goto L59
	}
L59:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v191 = F_NameListToString(m, v190)
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L13
	} else {
		goto L60
	}
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+32)) = v191
	F_errmsg(m, int32(_a_F_ExpandColumnRefStar_3), v13+int32(32))
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L13
	} else {
		goto L61
	}
L61:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	F_parser_errposition(m, l0, v199)
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L13
	} else {
		goto L62
	}
L62:
	;
	F_errfinish(m, int32(_a_F_ExpandColumnRefStar_1), int32(1244), int32(_a_F_ExpandColumnRefStar_4))
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L13
	} else {
		goto L63
	}
L63:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L64:
	;
	goto L46
L65:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L13
	} else {
		goto L79
	}
L66:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L13
	} else {
		goto L73
	}
L67:
	;
	if v165-int32(2) != 0 {
		goto L66
	} else {
		goto L70
	}
L68:
	;
	goto L69
L69:
	;
	v210 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v211 = F_makeRangeVar(m, v166, v167, v210)
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L13
	} else {
		goto L71
	}
L70:
	;
	goto L65
L71:
	;
	F_errorMissingRTE(m, l0, v211)
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L13
	} else {
		goto L72
	}
L72:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L73:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L13
	} else {
		goto L74
	}
L74:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v223 = F_NameListToString(m, v222)
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L13
	} else {
		goto L75
	}
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v223
	F_errmsg(m, int32(_a_F_ExpandColumnRefStar_5), v13)
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L13
	} else {
		goto L76
	}
L76:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	F_parser_errposition(m, l0, v229)
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L13
	} else {
		goto L77
	}
L77:
	;
	F_errfinish(m, int32(_a_F_ExpandColumnRefStar_1), int32(1265), int32(_a_F_ExpandColumnRefStar_4))
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L13
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
	F_errcode(m, int32(16801924))
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L13
	} else {
		goto L80
	}
L80:
	;
	v244 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v245 = F_NameListToString(m, v244)
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L13
	} else {
		goto L81
	}
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v245
	F_errmsg(m, int32(_a_F_ExpandColumnRefStar_6), v13+int32(16))
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L13
	} else {
		goto L82
	}
L82:
	;
	v253 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	F_parser_errposition(m, l0, v253)
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L13
	} else {
		goto L83
	}
L83:
	;
	F_errfinish(m, int32(_a_F_ExpandColumnRefStar_1), int32(1272), int32(_a_F_ExpandColumnRefStar_4))
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L13
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
	v263 = F_expandNSItemAttrs(m, l0, v164, v262, v261)
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L13
	} else {
		goto L88
	}
L86:
	;
	goto L87
L87:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v164)+12))
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v164)+4))
	v267 = int32(0)
	v269 = F_expandNSItemVars(m, l0, v164, v262, v261, v267)
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L13
	} else {
		goto L89
	}
L88:
	;
	v318 = v263
	goto L1
L89:
	;
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v266)+12))
	if v271 == int32(0) {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v274 = *(*int64)(unsafe.Add(mBase, uint32(v265)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v265)+16)) = v274 | int64(2)
	goto L92
L91:
	;
	goto L92
L92:
	;
	if v269 == int32(0) {
		v318 = v267
		goto L1
	} else {
		goto L93
	}
L93:
	;
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v269)+4))
	if int32(0) < v280 {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	v285 = int32(0)
	goto L97
L95:
	;
	goto L96
L96:
	;
	v318 = v269
	goto L1
L97:
	;
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v269)+12))
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v294+v285<<(uint(int32(2))%32))))
	F_markVarForSelectPriv(m, l0, v298)
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L13
	} else {
		goto L99
	}
L98:
	;
	goto L96
L99:
	;
	v302 = v285 + int32(1)
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v269)+4))
	if v302 < v303 {
		v285 = v302
		goto L97
	} else {
		goto L100
	}
L100:
	;
	goto L98
}
func F_has_column_privilege_id_attnum(m *base.Module, l0 int32) int64 {
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
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int64
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v45 int64
	_ = v45
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+40)))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v15 = F_pg_detoast_datum_packed(m, v14)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int64(0)
	} else {
		v20 = *(*int32)(unsafe.Add(mBase, _c_F_has_column_privilege_id_attnum[0]))
		v22 = F_convert_any_priv_string(m, v15, int32(_a_F_has_column_privilege_id_attnum_0))
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return int64(0)
		} else {
			v24 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)) = uint8(v24)
			if v12 == v24 {
				v40 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v40)
				v45 = int64(0)
				m.G0 = v10 + int32(16)
				return v45
			} else {
				v29 = v10 + int32(15)
				v30 = F_pg_attribute_aclcheck_ext(m, v13, v12, v20, v22, v29)
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return int64(0)
				} else {
					if v30 != 0 {
						v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)))
						if v32 != 0 {
							v40 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v40)
							v45 = int64(0)
							m.G0 = v10 + int32(16)
							return v45
						} else {
							v33 = F_pg_class_aclcheck_ext(m, v13, v20, v22, v29)
							mBase = m.M
							v34 = m.ExcPending
							if v34 != 0 {
								return int64(0)
							} else {
								if v33 != 0 {
									v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)))
									if v36 == int32(0) {
									} else {
										v40 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v40)
									}
									v45 = int64(0)
								} else {
									v45 = int64(1)
								}
								m.G0 = v10 + int32(16)
								return v45
							}
						}
					} else {
						v45 = int64(1)
						m.G0 = v10 + int32(16)
						return v45
					}
				}
			}
		}
	}
}
func F_has_column_privilege_id_id_name(m *base.Module, l0 int32) int64 {
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
	var v18 int32
	_ = v18
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
	var v25 int64
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
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
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v48 int64
	_ = v48
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v15 = F_pg_detoast_datum_packed(m, v14)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int64(0)
	} else {
		v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
		v20 = F_pg_detoast_datum_packed(m, v19)
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return int64(0)
		} else {
			v22 = F_convert_column_name(m, v13, v15)
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int64(0)
			} else {
				v25 = F_convert_any_priv_string(m, v20, int32(_a_F_has_column_privilege_id_id_name_0))
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int64(0)
				} else {
					v27 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)) = uint8(v27)
					if v22 == v27 {
						v43 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v43)
						v48 = int64(0)
						m.G0 = v10 + int32(16)
						return v48
					} else {
						v32 = v10 + int32(15)
						v33 = F_pg_attribute_aclcheck_ext(m, v13, v22, v12, v25, v32)
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
							return int64(0)
						} else {
							if v33 != 0 {
								v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)))
								if v35 != 0 {
									v43 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v43)
									v48 = int64(0)
									m.G0 = v10 + int32(16)
									return v48
								} else {
									v36 = F_pg_class_aclcheck_ext(m, v13, v12, v25, v32)
									mBase = m.M
									v37 = m.ExcPending
									if v37 != 0 {
										return int64(0)
									} else {
										if v36 != 0 {
											v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)))
											if v39 == int32(0) {
											} else {
												v43 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v43)
											}
											v48 = int64(0)
										} else {
											v48 = int64(1)
										}
										m.G0 = v10 + int32(16)
										return v48
									}
								}
							} else {
								v48 = int64(1)
								m.G0 = v10 + int32(16)
								return v48
							}
						}
					}
				}
			}
		}
	}
}
func F_has_column_privilege_id_name(m *base.Module, l0 int32) int64 {
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
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int64
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v49 int64
	_ = v49
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v14 = F_pg_detoast_datum_packed(m, v13)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int64(0)
	} else {
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
		v19 = F_pg_detoast_datum_packed(m, v18)
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int64(0)
		} else {
			v22 = *(*int32)(unsafe.Add(mBase, _c_F_has_column_privilege_id_name[0]))
			v23 = F_convert_column_name(m, v12, v14)
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int64(0)
			} else {
				v26 = F_convert_any_priv_string(m, v19, int32(_a_F_has_column_privilege_id_name_0))
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return int64(0)
				} else {
					v28 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)) = uint8(v28)
					if v23 == v28 {
						v44 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v44)
						v49 = int64(0)
						m.G0 = v10 + int32(16)
						return v49
					} else {
						v33 = v10 + int32(15)
						v34 = F_pg_attribute_aclcheck_ext(m, v12, v23, v22, v26, v33)
						mBase = m.M
						v35 = m.ExcPending
						if v35 != 0 {
							return int64(0)
						} else {
							if v34 != 0 {
								v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)))
								if v36 != 0 {
									v44 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v44)
									v49 = int64(0)
									m.G0 = v10 + int32(16)
									return v49
								} else {
									v37 = F_pg_class_aclcheck_ext(m, v12, v22, v26, v33)
									mBase = m.M
									v38 = m.ExcPending
									if v38 != 0 {
										return int64(0)
									} else {
										if v37 != 0 {
											v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)))
											if v40 == int32(0) {
											} else {
												v44 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v44)
											}
											v49 = int64(0)
										} else {
											v49 = int64(1)
										}
										m.G0 = v10 + int32(16)
										return v49
									}
								}
							} else {
								v49 = int64(1)
								m.G0 = v10 + int32(16)
								return v49
							}
						}
					}
				}
			}
		}
	}
}
func F_has_column_privilege_id_name_attnum(m *base.Module, l0 int32) int64 {
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
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
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
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int64
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
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
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v56 int64
	_ = v56
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v14 = F_pg_detoast_datum_packed(m, v13)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int64(0)
	} else {
		v18 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+56)))
		v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
		v20 = F_pg_detoast_datum_packed(m, v19)
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return int64(0)
		} else {
			v22 = F_textToQualifiedNameList(m, v14)
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int64(0)
			} else {
				v24 = F_makeRangeVarFromNameList(m, v22)
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int64(0)
				} else {
					v26 = int32(0)
					v30 = F_RangeVarGetRelidExtended(m, v24, v26, v26, v26, v26)
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return int64(0)
					} else {
						v33 = F_convert_any_priv_string(m, v20, int32(_a_F_has_column_privilege_id_name_attnum_0))
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
							return int64(0)
						} else {
							v35 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)) = uint8(v35)
							if v18 == v35 {
								v51 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v51)
								v56 = int64(0)
								m.G0 = v10 + int32(16)
								return v56
							} else {
								v40 = v10 + int32(15)
								v41 = F_pg_attribute_aclcheck_ext(m, v30, v18, v12, v33, v40)
								mBase = m.M
								v42 = m.ExcPending
								if v42 != 0 {
									return int64(0)
								} else {
									if v41 != 0 {
										v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)))
										if v43 != 0 {
											v51 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v51)
											v56 = int64(0)
											m.G0 = v10 + int32(16)
											return v56
										} else {
											v44 = F_pg_class_aclcheck_ext(m, v30, v12, v33, v40)
											mBase = m.M
											v45 = m.ExcPending
											if v45 != 0 {
												return int64(0)
											} else {
												if v44 != 0 {
													v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)))
													if v47 == int32(0) {
													} else {
														v51 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v51)
													}
													v56 = int64(0)
												} else {
													v56 = int64(1)
												}
												m.G0 = v10 + int32(16)
												return v56
											}
										}
									} else {
										v56 = int64(1)
										m.G0 = v10 + int32(16)
										return v56
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_has_column_privilege_id_name_name(m *base.Module, l0 int32) int64 {
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
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
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
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int64
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v60 int64
	_ = v60
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v14 = F_pg_detoast_datum_packed(m, v13)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int64(0)
	} else {
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
		v19 = F_pg_detoast_datum_packed(m, v18)
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int64(0)
		} else {
			v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
			v22 = F_pg_detoast_datum_packed(m, v21)
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int64(0)
			} else {
				v24 = F_textToQualifiedNameList(m, v14)
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int64(0)
				} else {
					v26 = F_makeRangeVarFromNameList(m, v24)
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return int64(0)
					} else {
						v28 = int32(0)
						v32 = F_RangeVarGetRelidExtended(m, v26, v28, v28, v28, v28)
						mBase = m.M
						v33 = m.ExcPending
						if v33 != 0 {
							return int64(0)
						} else {
							v34 = F_convert_column_name(m, v32, v19)
							mBase = m.M
							v35 = m.ExcPending
							if v35 != 0 {
								return int64(0)
							} else {
								v37 = F_convert_any_priv_string(m, v22, int32(_a_F_has_column_privilege_id_name_name_0))
								mBase = m.M
								v38 = m.ExcPending
								if v38 != 0 {
									return int64(0)
								} else {
									v39 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)) = uint8(v39)
									if v34 == v39 {
										v55 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v55)
										v60 = int64(0)
										m.G0 = v10 + int32(16)
										return v60
									} else {
										v44 = v10 + int32(15)
										v45 = F_pg_attribute_aclcheck_ext(m, v32, v34, v12, v37, v44)
										mBase = m.M
										v46 = m.ExcPending
										if v46 != 0 {
											return int64(0)
										} else {
											if v45 != 0 {
												v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)))
												if v47 != 0 {
													v55 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v55)
													v60 = int64(0)
													m.G0 = v10 + int32(16)
													return v60
												} else {
													v48 = F_pg_class_aclcheck_ext(m, v32, v12, v37, v44)
													mBase = m.M
													v49 = m.ExcPending
													if v49 != 0 {
														return int64(0)
													} else {
														if v48 != 0 {
															v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)))
															if v51 == int32(0) {
															} else {
																v55 = int32(1)
																*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v55)
															}
															v60 = int64(0)
														} else {
															v60 = int64(1)
														}
														m.G0 = v10 + int32(16)
														return v60
													}
												}
											} else {
												v60 = int64(1)
												m.G0 = v10 + int32(16)
												return v60
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_has_column_privilege_name_name(m *base.Module, l0 int32) int64 {
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
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
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
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int64
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v61 int64
	_ = v61
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v13 = F_pg_detoast_datum_packed(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int64(0)
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v18 = F_pg_detoast_datum_packed(m, v17)
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int64(0)
		} else {
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
			v21 = F_pg_detoast_datum_packed(m, v20)
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int64(0)
			} else {
				v24 = *(*int32)(unsafe.Add(mBase, _c_F_has_column_privilege_name_name[0]))
				v25 = F_textToQualifiedNameList(m, v13)
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int64(0)
				} else {
					v27 = F_makeRangeVarFromNameList(m, v25)
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return int64(0)
					} else {
						v29 = int32(0)
						v33 = F_RangeVarGetRelidExtended(m, v27, v29, v29, v29, v29)
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
							return int64(0)
						} else {
							v35 = F_convert_column_name(m, v33, v18)
							mBase = m.M
							v36 = m.ExcPending
							if v36 != 0 {
								return int64(0)
							} else {
								v38 = F_convert_any_priv_string(m, v21, int32(_a_F_has_column_privilege_name_name_0))
								mBase = m.M
								v39 = m.ExcPending
								if v39 != 0 {
									return int64(0)
								} else {
									v40 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)) = uint8(v40)
									if v35 == v40 {
										v56 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v56)
										v61 = int64(0)
										m.G0 = v10 + int32(16)
										return v61
									} else {
										v45 = v10 + int32(15)
										v46 = F_pg_attribute_aclcheck_ext(m, v33, v35, v24, v38, v45)
										mBase = m.M
										v47 = m.ExcPending
										if v47 != 0 {
											return int64(0)
										} else {
											if v46 != 0 {
												v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)))
												if v48 != 0 {
													v56 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v56)
													v61 = int64(0)
													m.G0 = v10 + int32(16)
													return v61
												} else {
													v49 = F_pg_class_aclcheck_ext(m, v33, v24, v38, v45)
													mBase = m.M
													v50 = m.ExcPending
													if v50 != 0 {
														return int64(0)
													} else {
														if v49 != 0 {
															v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)))
															if v52 == int32(0) {
															} else {
																v56 = int32(1)
																*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v56)
															}
															v61 = int64(0)
														} else {
															v61 = int64(1)
														}
														m.G0 = v10 + int32(16)
														return v61
													}
												}
											} else {
												v61 = int64(1)
												m.G0 = v10 + int32(16)
												return v61
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_transformColumnDefinition(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
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
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v247 int32
	_ = v247
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v276 int32
	_ = v276
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v290 int32
	_ = v290
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v367 int32
	_ = v367
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v391 int32
	_ = v391
	var v393 int32
	_ = v393
	var v395 int32
	_ = v395
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v407 int32
	_ = v407
	var v410 int32
	_ = v410
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v419 int32
	_ = v419
	var v424 int32
	_ = v424
	var v428 int32
	_ = v428
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v445 int32
	_ = v445
	var v450 int32
	_ = v450
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v471 int32
	_ = v471
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v502 int32
	_ = v502
	var v509 int32
	_ = v509
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v535 int32
	_ = v535
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v542 int32
	_ = v542
	var v546 int32
	_ = v546
	var v552 int32
	_ = v552
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v572 int32
	_ = v572
	var v574 int32
	_ = v574
	var v578 int32
	_ = v578
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v586 int32
	_ = v586
	var v589 int32
	_ = v589
	var v592 int32
	_ = v592
	var v595 int32
	_ = v595
	var v598 int32
	_ = v598
	var v601 int32
	_ = v601
	var v604 int32
	_ = v604
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v615 int32
	_ = v615
	var v616 int32
	_ = v616
	var v618 int32
	_ = v618
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v626 int32
	_ = v626
	var v629 int32
	_ = v629
	var v634 int32
	_ = v634
	var v637 int32
	_ = v637
	var v640 int32
	_ = v640
	var v641 int32
	_ = v641
	var v644 int32
	_ = v644
	var v645 int32
	_ = v645
	var v648 int32
	_ = v648
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v659 int32
	_ = v659
	var v660 int32
	_ = v660
	var v662 int32
	_ = v662
	var v667 int32
	_ = v667
	var v670 int32
	_ = v670
	var v673 int32
	_ = v673
	var v674 int32
	_ = v674
	var v675 int32
	_ = v675
	var v677 int32
	_ = v677
	var v678 int32
	_ = v678
	var v679 int32
	_ = v679
	var v680 int32
	_ = v680
	var v682 int32
	_ = v682
	var v684 int32
	_ = v684
	var v687 int32
	_ = v687
	var v688 int32
	_ = v688
	var v690 int32
	_ = v690
	var v694 int32
	_ = v694
	var v695 int32
	_ = v695
	var v701 int32
	_ = v701
	var v706 int32
	_ = v706
	var v709 int32
	_ = v709
	var v710 int32
	_ = v710
	var v711 int32
	_ = v711
	var v712 int32
	_ = v712
	var v719 int32
	_ = v719
	var v720 int32
	_ = v720
	var v721 int32
	_ = v721
	var v723 int32
	_ = v723
	var v728 int32
	_ = v728
	var v729 int32
	_ = v729
	var v734 int32
	_ = v734
	var v736 int32
	_ = v736
	var v739 int32
	_ = v739
	var v742 int32
	_ = v742
	var v743 int32
	_ = v743
	var v749 int32
	_ = v749
	var v752 int32
	_ = v752
	var v756 int32
	_ = v756
	var v757 int32
	_ = v757
	var v758 int32
	_ = v758
	var v760 int32
	_ = v760
	var v765 int32
	_ = v765
	var v766 int32
	_ = v766
	var v769 int32
	_ = v769
	var v770 int32
	_ = v770
	var v773 int32
	_ = v773
	var v774 int32
	_ = v774
	var v775 int32
	_ = v775
	var v781 int32
	_ = v781
	var v782 int32
	_ = v782
	var v785 int32
	_ = v785
	var v786 int32
	_ = v786
	var v787 int32
	_ = v787
	var v792 int32
	_ = v792
	var v796 int32
	_ = v796
	var v801 int32
	_ = v801
	var v802 int32
	_ = v802
	var v805 int32
	_ = v805
	var v806 int32
	_ = v806
	var v807 int32
	_ = v807
	var v813 int32
	_ = v813
	var v814 int32
	_ = v814
	var v816 int32
	_ = v816
	var v817 int32
	_ = v817
	var v818 int32
	_ = v818
	var v823 int32
	_ = v823
	var v824 int32
	_ = v824
	var v828 int32
	_ = v828
	var v833 int32
	_ = v833
	var v834 int32
	_ = v834
	var v835 int32
	_ = v835
	var v836 int32
	_ = v836
	var v838 int32
	_ = v838
	var v839 int32
	_ = v839
	var v840 int32
	_ = v840
	var v842 int32
	_ = v842
	var v843 int32
	_ = v843
	var v845 int32
	_ = v845
	var v856 int32
	_ = v856
	var v857 int32
	_ = v857
	var v862 int32
	_ = v862
	var v865 int32
	_ = v865
	var v866 int32
	_ = v866
	var v867 int32
	_ = v867
	var v868 int32
	_ = v868
	var v875 int32
	_ = v875
	var v876 int32
	_ = v876
	var v877 int32
	_ = v877
	var v879 int32
	_ = v879
	var v884 int32
	_ = v884
	var v888 int32
	_ = v888
	var v891 int32
	_ = v891
	var v895 int32
	_ = v895
	var v900 int32
	_ = v900
	var v904 int32
	_ = v904
	var v907 int32
	_ = v907
	var v908 int32
	_ = v908
	var v909 int32
	_ = v909
	var v910 int32
	_ = v910
	var v917 int32
	_ = v917
	var v918 int32
	_ = v918
	var v919 int32
	_ = v919
	var v921 int32
	_ = v921
	var v926 int32
	_ = v926
	var v930 int32
	_ = v930
	var v933 int32
	_ = v933
	var v934 int32
	_ = v934
	var v940 int32
	_ = v940
	var v945 int32
	_ = v945
	var v949 int32
	_ = v949
	var v950 int32
	_ = v950
	var v951 int32
	_ = v951
	var v958 int32
	_ = v958
	var v963 int32
	_ = v963
	var v967 int32
	_ = v967
	var v970 int32
	_ = v970
	var v971 int32
	_ = v971
	var v977 int32
	_ = v977
	var v982 int32
	_ = v982
	var v986 int32
	_ = v986
	var v989 int32
	_ = v989
	var v990 int32
	_ = v990
	var v991 int32
	_ = v991
	var v992 int32
	_ = v992
	var v999 int32
	_ = v999
	var v1000 int32
	_ = v1000
	var v1001 int32
	_ = v1001
	var v1003 int32
	_ = v1003
	var v1008 int32
	_ = v1008
	var v1012 int32
	_ = v1012
	var v1015 int32
	_ = v1015
	var v1019 int32
	_ = v1019
	var v1024 int32
	_ = v1024
	var v1028 int32
	_ = v1028
	var v1031 int32
	_ = v1031
	var v1035 int32
	_ = v1035
	var v1040 int32
	_ = v1040
	var v1044 int32
	_ = v1044
	var v1047 int32
	_ = v1047
	var v1048 int32
	_ = v1048
	var v1049 int32
	_ = v1049
	var v1050 int32
	_ = v1050
	var v1057 int32
	_ = v1057
	var v1058 int32
	_ = v1058
	var v1059 int32
	_ = v1059
	var v1061 int32
	_ = v1061
	var v1066 int32
	_ = v1066
	var v1070 int32
	_ = v1070
	var v1073 int32
	_ = v1073
	var v1077 int32
	_ = v1077
	var v1082 int32
	_ = v1082
	var v1086 int32
	_ = v1086
	var v1089 int32
	_ = v1089
	var v1090 int32
	_ = v1090
	var v1091 int32
	_ = v1091
	var v1092 int32
	_ = v1092
	var v1099 int32
	_ = v1099
	var v1100 int32
	_ = v1100
	var v1101 int32
	_ = v1101
	var v1103 int32
	_ = v1103
	var v1108 int32
	_ = v1108
	var v1112 int32
	_ = v1112
	var v1115 int32
	_ = v1115
	var v1116 int32
	_ = v1116
	var v1117 int32
	_ = v1117
	var v1118 int32
	_ = v1118
	var v1125 int32
	_ = v1125
	var v1126 int32
	_ = v1126
	var v1127 int32
	_ = v1127
	var v1129 int32
	_ = v1129
	var v1134 int32
	_ = v1134
	var v1138 int32
	_ = v1138
	var v1141 int32
	_ = v1141
	var v1145 int32
	_ = v1145
	var v1146 int32
	_ = v1146
	var v1147 int32
	_ = v1147
	var v1149 int32
	_ = v1149
	var v1154 int32
	_ = v1154
	var v1158 int32
	_ = v1158
	var v1161 int32
	_ = v1161
	var v1165 int32
	_ = v1165
	var v1166 int32
	_ = v1166
	var v1167 int32
	_ = v1167
	var v1169 int32
	_ = v1169
	var v1174 int32
	_ = v1174
	var v1178 int32
	_ = v1178
	var v1181 int32
	_ = v1181
	var v1182 int32
	_ = v1182
	var v1183 int32
	_ = v1183
	var v1184 int32
	_ = v1184
	var v1191 int32
	_ = v1191
	var v1192 int32
	_ = v1192
	var v1193 int32
	_ = v1193
	var v1195 int32
	_ = v1195
	var v1200 int32
	_ = v1200
	var v1204 int32
	_ = v1204
	var v1207 int32
	_ = v1207
	var v1208 int32
	_ = v1208
	var v1209 int32
	_ = v1209
	var v1210 int32
	_ = v1210
	var v1217 int32
	_ = v1217
	var v1218 int32
	_ = v1218
	var v1219 int32
	_ = v1219
	var v1221 int32
	_ = v1221
	var v1226 int32
	_ = v1226
	var v1230 int32
	_ = v1230
	var v1233 int32
	_ = v1233
	var v1234 int32
	_ = v1234
	var v1235 int32
	_ = v1235
	var v1236 int32
	_ = v1236
	var v1243 int32
	_ = v1243
	var v1244 int32
	_ = v1244
	var v1245 int32
	_ = v1245
	var v1247 int32
	_ = v1247
	var v1252 int32
	_ = v1252
	var v1261 int32
	_ = v1261
	var v1267 int32
	_ = v1267
	var v1275 int32
	_ = v1275
	var v1280 int32
	_ = v1280
	var v1283 int32
	_ = v1283
	var v1288 int32
	_ = v1288
	var v1294 int32
	_ = v1294
	var v1296 int32
	_ = v1296
	var v1297 int32
	_ = v1297
	var v1298 int32
	_ = v1298
	var v1299 int32
	_ = v1299
	var v1300 int32
	_ = v1300
	var v1301 int32
	_ = v1301
	var v1302 int32
	_ = v1302
	var v1303 int32
	_ = v1303
	var v1323 int32
	_ = v1323
	var v1325 int32
	_ = v1325
	var v1326 int32
	_ = v1326
	var v1329 int32
	_ = v1329
	var v1331 int32
	_ = v1331
	var v1332 int32
	_ = v1332
	var v1338 int32
	_ = v1338
	var v1339 int32
	_ = v1339
	var v1342 int32
	_ = v1342
	var v1347 int32
	_ = v1347
	var v1348 int32
	_ = v1348
	var v1350 int32
	_ = v1350
	var v1351 int32
	_ = v1351
	var v1352 int32
	_ = v1352
	v3 = int32(0)
	v19 = m.G0
	v21 = v19 - int32(256)
	m.G0 = v21
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v24 = F_lappend(m, v23, l1)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v24
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v27 != 0 {
		goto L12
	} else {
		goto L13
	}
L3:
	;
	v1323 = *(*int32)(unsafe.Add(mBase, uint32(l1)+60))
	if v1323 != 0 {
		goto L342
	} else {
		goto L343
	}
L4:
	;
	v1294 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1283))) = uint8(v1294)
	v1296 = *(*int32)(unsafe.Add(mBase, uint32(v1280)))
	v1297 = F_makeString(m, v1296)
	mBase = m.M
	v1298 = m.ExcPending
	if v1298 != 0 {
		goto L1
	} else {
		goto L339
	}
L5:
	;
	v546 = *(*int32)(unsafe.Add(mBase, uint32(v539)+4))
	if v546 <= int32(0) {
		goto L125
	} else {
		goto L126
	}
L6:
	;
	v532 = l1 + int32(4)
	v533 = v509
	v535 = l1 + int32(19)
	v539 = v257
	v540 = l0 + int32(32)
	v542 = v3
	goto L5
L7:
	;
	v471 = *(*int32)(unsafe.Add(mBase, uint32(v257)+12))
	v477 = v458
	v478 = v459
	v479 = v270
	goto L119
L8:
	;
	if v269 == int32(0) {
		v509 = v321
		goto L6
	} else {
		goto L118
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v428 = m.ExcPending
	if v428 != 0 {
		goto L1
	} else {
		goto L112
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v407 = m.ExcPending
	if v407 != 0 {
		goto L1
	} else {
		goto L107
	}
L11:
	;
	v327 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v327)+8))
	v329 = int32(0)
	F_generateSerialExtraStmts(m, l0, l1, v328, v329, v329, v329, v21+int32(252), v21+int32(248))
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L1
	} else {
		goto L95
	}
L12:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
	if v28 == int32(0) {
		v225 = v3
		v226 = v27
		goto L15
	} else {
		goto L16
	}
L13:
	;
	goto L14
L14:
	;
	v253 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v254 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	F_transformConstraintAttrs(m, v253, v254)
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L1
	} else {
		goto L77
	}
L15:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v229 = F_typenameType(m, v227, v226, int32(0))
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L1
	} else {
		goto L69
	}
L16:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	if v31 != int32(1) {
		v225 = v3
		v226 = v27
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+13)))
	if v34 != 0 {
		v225 = v3
		v226 = v27
		goto L15
	} else {
		goto L18
	}
L18:
	;
	v35 = int32(21)
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v28)+12))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	v39 = int32(_a_F_transformColumnDefinition_0)
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38))))
	v45 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_transformColumnDefinition[0])))
	if base.B2i32(v42 == int32(0))|base.B2i32(v42 != v45) != 0 {
		v63 = v42
		v64 = v45
		goto L21
	} else {
		goto L22
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+4)) = int32(0)
	v219 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v219)+8)) = v216
	v222 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v222)+24))
	if v223 != 0 {
		goto L10
	} else {
		goto L68
	}
L20:
	;
	if v63-v64 == int32(0) {
		v216 = v35
		goto L19
	} else {
		goto L27
	}
L21:
	;
	goto L20
L22:
	;
	v48 = v38
	v49 = v39
	goto L23
L23:
	;
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49)+1)))
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+1)))
	if v53 == int32(0) {
		v63 = v53
		v64 = v52
		goto L21
	} else {
		goto L25
	}
L24:
	;
	v63 = v53
	v64 = v52
	goto L21
L25:
	;
	v56 = int32(1)
	if v53 == v52 {
		v48 = v48 + v56
		v49 = v49 + v56
		goto L23
	} else {
		goto L26
	}
L26:
	;
	goto L24
L27:
	;
	v68 = int32(_a_F_transformColumnDefinition_1)
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38))))
	v74 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_transformColumnDefinition[1])))
	if base.B2i32(v71 == int32(0))|base.B2i32(v71 != v74) != 0 {
		v92 = v71
		v93 = v74
		goto L29
	} else {
		goto L30
	}
L28:
	;
	if v92-v93 == int32(0) {
		v216 = v35
		goto L19
	} else {
		goto L35
	}
L29:
	;
	goto L28
L30:
	;
	v77 = v38
	v78 = v68
	goto L31
L31:
	;
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78)+1)))
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77)+1)))
	if v82 == int32(0) {
		v92 = v82
		v93 = v81
		goto L29
	} else {
		goto L33
	}
L32:
	;
	v92 = v82
	v93 = v81
	goto L29
L33:
	;
	v85 = int32(1)
	if v82 == v81 {
		v77 = v77 + v85
		v78 = v78 + v85
		goto L31
	} else {
		goto L34
	}
L34:
	;
	goto L32
L35:
	;
	v97 = int32(23)
	v98 = int32(_a_F_transformColumnDefinition_2)
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38))))
	v104 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_transformColumnDefinition[2])))
	if base.B2i32(v101 == int32(0))|base.B2i32(v101 != v104) != 0 {
		v122 = v101
		v123 = v104
		goto L37
	} else {
		goto L38
	}
L36:
	;
	if v122-v123 == int32(0) {
		v216 = v97
		goto L19
	} else {
		goto L43
	}
L37:
	;
	goto L36
L38:
	;
	v107 = v38
	v108 = v98
	goto L39
L39:
	;
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108)+1)))
	v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+1)))
	if v112 == int32(0) {
		v122 = v112
		v123 = v111
		goto L37
	} else {
		goto L41
	}
L40:
	;
	v122 = v112
	v123 = v111
	goto L37
L41:
	;
	v115 = int32(1)
	if v112 == v111 {
		v107 = v107 + v115
		v108 = v108 + v115
		goto L39
	} else {
		goto L42
	}
L42:
	;
	goto L40
L43:
	;
	v127 = int32(_a_F_transformColumnDefinition_3)
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38))))
	v133 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_transformColumnDefinition[3])))
	if base.B2i32(v130 == int32(0))|base.B2i32(v130 != v133) != 0 {
		v151 = v130
		v152 = v133
		goto L45
	} else {
		goto L46
	}
L44:
	;
	if v151-v152 == int32(0) {
		v216 = v97
		goto L19
	} else {
		goto L51
	}
L45:
	;
	goto L44
L46:
	;
	v136 = v38
	v137 = v127
	goto L47
L47:
	;
	v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v137)+1)))
	v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v136)+1)))
	if v141 == int32(0) {
		v151 = v141
		v152 = v140
		goto L45
	} else {
		goto L49
	}
L48:
	;
	v151 = v141
	v152 = v140
	goto L45
L49:
	;
	v144 = int32(1)
	if v141 == v140 {
		v136 = v136 + v144
		v137 = v137 + v144
		goto L47
	} else {
		goto L50
	}
L50:
	;
	goto L48
L51:
	;
	v156 = int32(20)
	v157 = int32(_a_F_transformColumnDefinition_4)
	v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38))))
	v163 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_transformColumnDefinition[4])))
	if base.B2i32(v160 == int32(0))|base.B2i32(v160 != v163) != 0 {
		v181 = v160
		v182 = v163
		goto L53
	} else {
		goto L54
	}
L52:
	;
	if v181-v182 == int32(0) {
		v216 = v156
		goto L19
	} else {
		goto L59
	}
L53:
	;
	goto L52
L54:
	;
	v166 = v38
	v167 = v157
	goto L55
L55:
	;
	v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+1)))
	v171 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166)+1)))
	if v171 == int32(0) {
		v181 = v171
		v182 = v170
		goto L53
	} else {
		goto L57
	}
L56:
	;
	v181 = v171
	v182 = v170
	goto L53
L57:
	;
	v174 = int32(1)
	if v171 == v170 {
		v166 = v166 + v174
		v167 = v167 + v174
		goto L55
	} else {
		goto L58
	}
L58:
	;
	goto L56
L59:
	;
	v186 = int32(_a_F_transformColumnDefinition_5)
	v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38))))
	v192 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_transformColumnDefinition[5])))
	if base.B2i32(v189 == int32(0))|base.B2i32(v189 != v192) != 0 {
		v210 = v189
		v211 = v192
		goto L61
	} else {
		goto L62
	}
L60:
	;
	if v210-v211 == int32(0) {
		v216 = v156
		goto L19
	} else {
		goto L67
	}
L61:
	;
	goto L60
L62:
	;
	v195 = v38
	v196 = v186
	goto L63
L63:
	;
	v199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v196)+1)))
	v200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v195)+1)))
	if v200 == int32(0) {
		v210 = v200
		v211 = v199
		goto L61
	} else {
		goto L65
	}
L64:
	;
	v210 = v200
	v211 = v199
	goto L61
L65:
	;
	v203 = int32(1)
	if v200 == v199 {
		v195 = v195 + v203
		v196 = v196 + v203
		goto L63
	} else {
		goto L66
	}
L66:
	;
	goto L64
L67:
	;
	v225 = int32(0)
	v226 = v27
	goto L15
L68:
	;
	v225 = int32(1)
	v226 = v222
	goto L15
L69:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	if v231 != 0 {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v229)+16))
	v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v232)+22)))
	v234 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v231)+8))
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v231)+12))
	v237 = F_LookupCollation(m, v234, v235, v236)
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L1
	} else {
		goto L73
	}
L71:
	;
	goto L72
L72:
	;
	F_ReleaseCatCache(m, v229)
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L1
	} else {
		goto L75
	}
L73:
	;
	v239 = v232 + v233
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v239)+144))
	if v240 == int32(0) {
		goto L9
	} else {
		goto L74
	}
L74:
	;
	goto L72
L75:
	;
	if v225 != 0 {
		goto L11
	} else {
		goto L76
	}
L76:
	;
	goto L14
L77:
	;
	v257 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	if v257 == int32(0) {
		goto L3
	} else {
		goto L78
	}
L78:
	;
	v260 = int32(0)
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v257)+4))
	if v261 <= v260 {
		v509 = v260
		goto L6
	} else {
		goto L79
	}
L79:
	;
	v264 = int32(0)
	if v264 < v261 {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v267 = v261
	goto L82
L81:
	;
	v267 = v264
	goto L82
L82:
	;
	v269 = v267 & int32(3)
	v270 = int32(0)
	if v261 < int32(4) {
		v458 = v260
		v459 = v270
		goto L7
	} else {
		goto L83
	}
L83:
	;
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v257)+12))
	v282 = v260
	v283 = v270
	v290 = v3
	goto L84
L84:
	;
	v297 = v276 + v283<<(uint(int32(2))%32)
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v297)))
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v298)+4))
	switch v299 - int32(3) {
	case 0, 3:
		goto L87
	default:
		v303 = v282
		goto L86
	}
L85:
	;
	goto L8
L86:
	;
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v297)+4))
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v304)+4))
	switch v305 - int32(3) {
	case 0, 3:
		goto L89
	default:
		v309 = v303
		goto L88
	}
L87:
	;
	v303 = int32(1)
	goto L86
L88:
	;
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v297)+8))
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v310)+4))
	switch v311 - int32(3) {
	case 0, 3:
		goto L91
	default:
		v315 = v309
		goto L90
	}
L89:
	;
	v309 = int32(1)
	goto L88
L90:
	;
	v316 = *(*int32)(unsafe.Add(mBase, uint32(v297)+12))
	v317 = *(*int32)(unsafe.Add(mBase, uint32(v316)+4))
	switch v317 - int32(3) {
	case 0, 3:
		goto L93
	default:
		v321 = v315
		goto L92
	}
L91:
	;
	v315 = int32(1)
	goto L90
L92:
	;
	v322 = int32(4)
	v323 = v283 + v322
	v325 = v290 + v322
	if v267&int32(2147483644) != v325 {
		v282 = v321
		v283 = v323
		v290 = v325
		goto L84
	} else {
		goto L94
	}
L93:
	;
	v321 = int32(1)
	goto L92
L94:
	;
	goto L85
L95:
	;
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v21)+252))
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v21)+248))
	v340 = F_quote_qualified_identifier(m, v338, v339)
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		goto L1
	} else {
		goto L96
	}
L96:
	;
	v343 = F_palloc0(m, int32(20))
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L1
	} else {
		goto L97
	}
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v343)+16)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v343)+8)) = v340
	*(*int64)(unsafe.Add(mBase, uint32(v343))) = int64(2044404432968)
	v351 = F_palloc0(m, int32(16))
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L1
	} else {
		goto L98
	}
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v351))) = int32(73)
	v356 = F_SystemTypeName(m, int32(_a_F_transformColumnDefinition_6))
	mBase = m.M
	v357 = m.ExcPending
	if v357 != 0 {
		goto L1
	} else {
		goto L99
	}
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v351)+12)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v351)+4)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v351)+8)) = v356
	v363 = F_SystemFuncName(m, int32(_a_F_transformColumnDefinition_7))
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L1
	} else {
		goto L100
	}
L100:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+220)) = v351
	*(*int32)(unsafe.Add(mBase, uint32(v21)+244)) = v351
	v367 = int32(1)
	v371 = F_list_make1_impl(m, v367, v21+int32(220))
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L1
	} else {
		goto L101
	}
L101:
	;
	v375 = F_makeFuncCall(m, v363, v371, int32(0), int32(-1))
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L1
	} else {
		goto L102
	}
L102:
	;
	v378 = F_palloc0(m, int32(108))
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L1
	} else {
		goto L103
	}
L103:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v378)+104)) = int32(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v378))) = int64(8589934753)
	*(*int32)(unsafe.Add(mBase, uint32(v378)+24)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v378)+20)) = v375
	v387 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v388 = F_lappend(m, v387, v378)
	mBase = m.M
	v389 = m.ExcPending
	if v389 != 0 {
		goto L1
	} else {
		goto L104
	}
L104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+56)) = v388
	v391 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_transformConstraintAttrs(m, v391, v388)
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L1
	} else {
		goto L105
	}
L105:
	;
	v395 = l0 + int32(32)
	v397 = l1 + int32(19)
	v399 = l1 + int32(4)
	v401 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	if v401 == int32(0) {
		v1280 = v399
		v1283 = v397
		v1288 = v395
		goto L4
	} else {
		goto L106
	}
L106:
	;
	v532 = v399
	v533 = v367
	v535 = v397
	v539 = v401
	v540 = v395
	v542 = int32(1)
	goto L5
L107:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v410 = m.ExcPending
	if v410 != 0 {
		goto L1
	} else {
		goto L108
	}
L108:
	;
	F_errmsg(m, int32(_a_F_transformColumnDefinition_8), int32(0))
	mBase = m.M
	v414 = m.ExcPending
	if v414 != 0 {
		goto L1
	} else {
		goto L109
	}
L109:
	;
	v415 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v416 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v417 = *(*int32)(unsafe.Add(mBase, uint32(v416)+28))
	F_parser_errposition(m, v415, v417)
	mBase = m.M
	v419 = m.ExcPending
	if v419 != 0 {
		goto L1
	} else {
		goto L110
	}
L110:
	;
	F_errfinish(m, int32(_a_F_transformColumnDefinition_9), int32(650), int32(_a_F_transformColumnDefinition_10))
	mBase = m.M
	v424 = m.ExcPending
	if v424 != 0 {
		goto L1
	} else {
		goto L111
	}
L111:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L112:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v431 = m.ExcPending
	if v431 != 0 {
		goto L1
	} else {
		goto L113
	}
L113:
	;
	v432 = *(*int32)(unsafe.Add(mBase, uint32(v239)))
	v433 = F_format_type_be(m, v432)
	mBase = m.M
	v434 = m.ExcPending
	if v434 != 0 {
		goto L1
	} else {
		goto L114
	}
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+224)) = v433
	F_errmsg(m, int32(_a_F_transformColumnDefinition_11), v21+int32(224))
	mBase = m.M
	v440 = m.ExcPending
	if v440 != 0 {
		goto L1
	} else {
		goto L115
	}
L115:
	;
	v441 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v442 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v443 = *(*int32)(unsafe.Add(mBase, uint32(v442)+12))
	F_parser_errposition(m, v441, v443)
	mBase = m.M
	v445 = m.ExcPending
	if v445 != 0 {
		goto L1
	} else {
		goto L116
	}
L116:
	;
	F_errfinish(m, int32(_a_F_transformColumnDefinition_9), int32(_a_F_transformColumnDefinition_12), int32(_a_F_transformColumnDefinition_13))
	mBase = m.M
	v450 = m.ExcPending
	if v450 != 0 {
		goto L1
	} else {
		goto L117
	}
L117:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L118:
	;
	v458 = v321
	v459 = v323
	goto L7
L119:
	;
	v493 = *(*int32)(unsafe.Add(mBase, uint32(v471+v478<<(uint(int32(2))%32))))
	v494 = *(*int32)(unsafe.Add(mBase, uint32(v493)+4))
	switch v494 - int32(3) {
	case 0, 3:
		goto L122
	default:
		v498 = v477
		goto L121
	}
L120:
	;
	v509 = v498
	goto L6
L121:
	;
	v499 = int32(1)
	v502 = v479 + v499
	if v502 != v269 {
		v477 = v498
		v478 = v478 + v499
		v479 = v502
		goto L119
	} else {
		goto L123
	}
L122:
	;
	v498 = int32(1)
	goto L121
L123:
	;
	goto L120
L124:
	;
	if v1267 == int32(0) {
		goto L3
	} else {
		goto L336
	}
L125:
	;
	v1261 = int32(0)
	v1267 = v542
	goto L124
L126:
	;
	goto L127
L127:
	;
	v552 = int32(0)
	v561 = v552
	v562 = v552
	v564 = v552
	v565 = v3
	v569 = v552
	v570 = v542
	v572 = v3
	goto L145
L128:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1230 = m.ExcPending
	if v1230 != 0 {
		goto L1
	} else {
		goto L331
	}
L129:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1204 = m.ExcPending
	if v1204 != 0 {
		goto L1
	} else {
		goto L326
	}
L130:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1178 = m.ExcPending
	if v1178 != 0 {
		goto L1
	} else {
		goto L321
	}
L131:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1158 = m.ExcPending
	if v1158 != 0 {
		goto L1
	} else {
		goto L316
	}
L132:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1138 = m.ExcPending
	if v1138 != 0 {
		goto L1
	} else {
		goto L311
	}
L133:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1112 = m.ExcPending
	if v1112 != 0 {
		goto L1
	} else {
		goto L306
	}
L134:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1086 = m.ExcPending
	if v1086 != 0 {
		goto L1
	} else {
		goto L301
	}
L135:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1070 = m.ExcPending
	if v1070 != 0 {
		goto L1
	} else {
		goto L297
	}
L136:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1044 = m.ExcPending
	if v1044 != 0 {
		goto L1
	} else {
		goto L292
	}
L137:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1028 = m.ExcPending
	if v1028 != 0 {
		goto L1
	} else {
		goto L288
	}
L138:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1012 = m.ExcPending
	if v1012 != 0 {
		goto L1
	} else {
		goto L284
	}
L139:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v986 = m.ExcPending
	if v986 != 0 {
		goto L1
	} else {
		goto L279
	}
L140:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v967 = m.ExcPending
	if v967 != 0 {
		goto L1
	} else {
		goto L275
	}
L141:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v949 = m.ExcPending
	if v949 != 0 {
		goto L1
	} else {
		goto L272
	}
L142:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v930 = m.ExcPending
	if v930 != 0 {
		goto L1
	} else {
		goto L268
	}
L143:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v904 = m.ExcPending
	if v904 != 0 {
		goto L1
	} else {
		goto L263
	}
L144:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v888 = m.ExcPending
	if v888 != 0 {
		goto L1
	} else {
		goto L259
	}
L145:
	;
	v574 = *(*int32)(unsafe.Add(mBase, uint32(v539)+12))
	v578 = *(*int32)(unsafe.Add(mBase, uint32(v574+v562<<(uint(int32(2))%32))))
	v579 = *(*int32)(unsafe.Add(mBase, uint32(v578)+4))
	switch v579 {
	case 0:
		goto L159
	case 1:
		goto L158
	case 2:
		goto L157
	case 3:
		goto L156
	case 4:
		goto L155
	case 5:
		goto L148
	case 6:
		goto L154
	case 7:
		goto L153
	case 8:
		goto L151
	case 9:
		goto L150
	case 10, 11, 12, 13, 14, 15:
		v838 = v561
		v839 = v564
		v840 = v565
		v842 = v569
		v843 = v570
		v845 = v572
		goto L147
	default:
		goto L149
	}
L146:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v862 = m.ExcPending
	if v862 != 0 {
		goto L1
	} else {
		goto L254
	}
L147:
	;
	if v845&(v842&int32(1)) != 0 {
		goto L130
	} else {
		goto L248
	}
L148:
	;
	v834 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v835 = F_lappend(m, v834, v578)
	mBase = m.M
	v836 = m.ExcPending
	if v836 != 0 {
		goto L1
	} else {
		goto L247
	}
L149:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v823 = m.ExcPending
	if v823 != 0 {
		goto L1
	} else {
		goto L244
	}
L150:
	;
	v802 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
	if v802 == int32(1) {
		goto L131
	} else {
		goto L240
	}
L151:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v792 = m.ExcPending
	if v792 != 0 {
		goto L1
	} else {
		goto L237
	}
L152:
	;
	v770 = *(*int32)(unsafe.Add(mBase, uint32(v578)+32))
	if v770 == int32(0) {
		goto L231
	} else {
		goto L232
	}
L153:
	;
	v766 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
	if v766 == int32(1) {
		goto L132
	} else {
		goto L230
	}
L154:
	;
	if v564 != 0 {
		goto L220
	} else {
		goto L221
	}
L155:
	;
	v729 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+68)))
	if v729 == int32(1) {
		goto L135
	} else {
		goto L218
	}
L156:
	;
	v670 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+68)))
	if v670 == int32(1) {
		goto L138
	} else {
		goto L201
	}
L157:
	;
	if v572 != 0 {
		goto L139
	} else {
		goto L200
	}
L158:
	;
	v589 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+60)))
	if v589 == int32(1) {
		goto L166
	} else {
		goto L167
	}
L159:
	;
	if v564 != 0 {
		goto L161
	} else {
		goto L162
	}
L160:
	;
	v586 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v535))) = uint8(v586)
	v838 = v561
	v839 = int32(1)
	v840 = v565
	v842 = v569
	v843 = v570
	v845 = v572
	goto L147
L161:
	;
	v580 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v535))))
	if (v580|v570)&int32(1) == int32(0) {
		goto L160
	} else {
		goto L164
	}
L162:
	;
	goto L163
L163:
	;
	if v570 != 0 {
		goto L128
	} else {
		goto L165
	}
L164:
	;
	goto L128
L165:
	;
	goto L160
L166:
	;
	v592 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v578)+17)))
	if v592 == int32(1) {
		goto L144
	} else {
		goto L169
	}
L167:
	;
	goto L168
L168:
	;
	if v564 != 0 {
		goto L170
	} else {
		goto L171
	}
L169:
	;
	goto L168
L170:
	;
	v595 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v535))))
	if v595 == int32(0) {
		goto L143
	} else {
		goto L173
	}
L171:
	;
	goto L172
L172:
	;
	if v533&int32(1) != 0 {
		goto L174
	} else {
		goto L175
	}
L173:
	;
	goto L172
L174:
	;
	v598 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v578)+17)))
	if v598 == int32(1) {
		goto L142
	} else {
		goto L177
	}
L175:
	;
	goto L176
L176:
	;
	v601 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v535))))
	if v601 == int32(0) {
		goto L178
	} else {
		goto L179
	}
L177:
	;
	goto L176
L178:
	;
	v604 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v535))) = uint8(v604)
	v607 = *(*int32)(unsafe.Add(mBase, uint32(v532)))
	v608 = F_makeString(m, v607)
	mBase = m.M
	v609 = m.ExcPending
	if v609 != 0 {
		goto L1
	} else {
		goto L181
	}
L179:
	;
	goto L180
L180:
	;
	if v565 == int32(0) {
		goto L184
	} else {
		goto L185
	}
L181:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+124)) = v608
	*(*int32)(unsafe.Add(mBase, uint32(v21)+240)) = v608
	v615 = F_list_make1_impl(m, int32(1), v21+int32(124))
	mBase = m.M
	v616 = m.ExcPending
	if v616 != 0 {
		goto L1
	} else {
		goto L182
	}
L182:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v578)+32)) = v615
	v618 = *(*int32)(unsafe.Add(mBase, uint32(v540)))
	v619 = F_lappend(m, v618, v578)
	mBase = m.M
	v620 = m.ExcPending
	if v620 != 0 {
		goto L1
	} else {
		goto L183
	}
L183:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v540))) = v619
	v838 = v561
	v839 = v604
	v840 = v578
	v842 = v569
	v843 = int32(0)
	v845 = v572
	goto L147
L184:
	;
	v838 = v561
	v839 = v564
	v840 = int32(0)
	v842 = v569
	v843 = v570
	v845 = v572
	goto L147
L185:
	;
	goto L186
L186:
	;
	v626 = *(*int32)(unsafe.Add(mBase, uint32(v578)+8))
	if v626 == int32(0) {
		goto L187
	} else {
		goto L188
	}
L187:
	;
	v659 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v565)+17)))
	v660 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v578)+17)))
	if v659 != v660 {
		goto L140
	} else {
		goto L198
	}
L188:
	;
	v629 = *(*int32)(unsafe.Add(mBase, uint32(v565)+8))
	if v629 == int32(0) {
		goto L187
	} else {
		goto L189
	}
L189:
	;
	v634 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v629))))
	v637 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v626))))
	if base.B2i32(v634 == int32(0))|base.B2i32(v634 != v637) != 0 {
		v655 = v634
		v656 = v637
		goto L191
	} else {
		goto L192
	}
L190:
	;
	if v655-v656 != 0 {
		goto L141
	} else {
		goto L197
	}
L191:
	;
	goto L190
L192:
	;
	v640 = v629
	v641 = v626
	goto L193
L193:
	;
	v644 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v641)+1)))
	v645 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v640)+1)))
	if v645 == int32(0) {
		v655 = v645
		v656 = v644
		goto L191
	} else {
		goto L195
	}
L194:
	;
	v655 = v645
	v656 = v644
	goto L191
L195:
	;
	v648 = int32(1)
	if v645 == v644 {
		v640 = v640 + v648
		v641 = v641 + v648
		goto L193
	} else {
		goto L196
	}
L196:
	;
	goto L194
L197:
	;
	goto L187
L198:
	;
	v662 = *(*int32)(unsafe.Add(mBase, uint32(v565)+8))
	if v662|base.B2i32(v626 == int32(0)) != 0 {
		v838 = v561
		v839 = v564
		v840 = v565
		v842 = v569
		v843 = v570
		v845 = v572
		goto L147
	} else {
		goto L199
	}
L199:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v565)+8)) = v626
	v838 = v561
	v839 = v564
	v840 = v565
	v842 = v569
	v843 = v570
	v845 = v572
	goto L147
L200:
	;
	v667 = *(*int32)(unsafe.Add(mBase, uint32(v578)+20))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v667
	v838 = v561
	v839 = v564
	v840 = v565
	v842 = v569
	v843 = v570
	v845 = int32(1)
	goto L147
L201:
	;
	v673 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	if v673 != 0 {
		goto L137
	} else {
		goto L202
	}
L202:
	;
	v674 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v675 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v677 = F_typenameType(m, v674, v675, int32(0))
	mBase = m.M
	v678 = m.ExcPending
	if v678 != 0 {
		goto L1
	} else {
		goto L203
	}
L203:
	;
	v679 = *(*int32)(unsafe.Add(mBase, uint32(v677)+16))
	v680 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v679)+22)))
	v682 = *(*int32)(unsafe.Add(mBase, uint32(v679+v680)))
	F_ReleaseCatCache(m, v677)
	mBase = m.M
	v684 = m.ExcPending
	if v684 != 0 {
		goto L1
	} else {
		goto L204
	}
L204:
	;
	if v569&int32(1) != 0 {
		goto L136
	} else {
		goto L205
	}
L205:
	;
	v687 = int32(1)
	v688 = *(*int32)(unsafe.Add(mBase, uint32(v578)+48))
	v690 = int32(0)
	F_generateSerialExtraStmts(m, l0, l1, v682, v688, v687, v690, v690, v690)
	mBase = m.M
	v694 = m.ExcPending
	if v694 != 0 {
		goto L1
	} else {
		goto L206
	}
L206:
	;
	v695 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v578)+28)))
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+36)) = uint8(v695)
	if v564 == int32(0) {
		goto L207
	} else {
		goto L208
	}
L207:
	;
	v838 = v561
	v839 = int32(0)
	v840 = v565
	v842 = v687
	v843 = int32(1)
	v845 = v572
	goto L147
L208:
	;
	goto L209
L209:
	;
	v701 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v535))))
	if v701 != 0 {
		goto L210
	} else {
		goto L211
	}
L210:
	;
	v838 = v561
	v839 = int32(1)
	v840 = v565
	v842 = v687
	v843 = v570
	v845 = v572
	goto L147
L211:
	;
	goto L212
L212:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v706 = m.ExcPending
	if v706 != 0 {
		goto L1
	} else {
		goto L213
	}
L213:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v709 = m.ExcPending
	if v709 != 0 {
		goto L1
	} else {
		goto L214
	}
L214:
	;
	v710 = *(*int32)(unsafe.Add(mBase, uint32(v532)))
	v711 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v712 = *(*int32)(unsafe.Add(mBase, uint32(v711)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+180)) = v712
	*(*int32)(unsafe.Add(mBase, uint32(v21)+176)) = v710
	F_errmsg(m, int32(_a_F_transformColumnDefinition_14), v21+int32(176))
	mBase = m.M
	v719 = m.ExcPending
	if v719 != 0 {
		goto L1
	} else {
		goto L215
	}
L215:
	;
	v720 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v721 = *(*int32)(unsafe.Add(mBase, uint32(v578)+104))
	F_parser_errposition(m, v720, v721)
	mBase = m.M
	v723 = m.ExcPending
	if v723 != 0 {
		goto L1
	} else {
		goto L216
	}
L216:
	;
	F_errfinish(m, int32(_a_F_transformColumnDefinition_9), int32(882), int32(_a_F_transformColumnDefinition_10))
	mBase = m.M
	v728 = m.ExcPending
	if v728 != 0 {
		goto L1
	} else {
		goto L217
	}
L217:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L218:
	;
	if v561&int32(1) != 0 {
		goto L134
	} else {
		goto L219
	}
L219:
	;
	v734 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v578)+29)))
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+44)) = uint8(v734)
	v736 = *(*int32)(unsafe.Add(mBase, uint32(v578)+20))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v736
	v838 = int32(1)
	v839 = v564
	v840 = v565
	v842 = v569
	v843 = v570
	v845 = v572
	goto L147
L220:
	;
	v739 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v535))))
	if v739 == int32(0) {
		goto L133
	} else {
		goto L223
	}
L221:
	;
	goto L222
L222:
	;
	v742 = int32(1)
	v743 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
	if v743 != v742 {
		v769 = v742
		goto L152
	} else {
		goto L224
	}
L223:
	;
	goto L222
L224:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v749 = m.ExcPending
	if v749 != 0 {
		goto L1
	} else {
		goto L225
	}
L225:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v752 = m.ExcPending
	if v752 != 0 {
		goto L1
	} else {
		goto L226
	}
L226:
	;
	F_errmsg(m, int32(_a_F_transformColumnDefinition_15), int32(0))
	mBase = m.M
	v756 = m.ExcPending
	if v756 != 0 {
		goto L1
	} else {
		goto L227
	}
L227:
	;
	v757 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v758 = *(*int32)(unsafe.Add(mBase, uint32(v578)+104))
	F_parser_errposition(m, v757, v758)
	mBase = m.M
	v760 = m.ExcPending
	if v760 != 0 {
		goto L1
	} else {
		goto L228
	}
L228:
	;
	F_errfinish(m, int32(_a_F_transformColumnDefinition_9), int32(923), int32(_a_F_transformColumnDefinition_10))
	mBase = m.M
	v765 = m.ExcPending
	if v765 != 0 {
		goto L1
	} else {
		goto L229
	}
L229:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L230:
	;
	v769 = v570
	goto L152
L231:
	;
	v773 = *(*int32)(unsafe.Add(mBase, uint32(v532)))
	v774 = F_makeString(m, v773)
	mBase = m.M
	v775 = m.ExcPending
	if v775 != 0 {
		goto L1
	} else {
		goto L234
	}
L232:
	;
	goto L233
L233:
	;
	v785 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v786 = F_lappend(m, v785, v578)
	mBase = m.M
	v787 = m.ExcPending
	if v787 != 0 {
		goto L1
	} else {
		goto L236
	}
L234:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+204)) = v774
	*(*int32)(unsafe.Add(mBase, uint32(v21)+236)) = v774
	v781 = F_list_make1_impl(m, int32(1), v21+int32(204))
	mBase = m.M
	v782 = m.ExcPending
	if v782 != 0 {
		goto L1
	} else {
		goto L235
	}
L235:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v578)+32)) = v781
	goto L233
L236:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v786
	v838 = v561
	v839 = v564
	v840 = v565
	v842 = v569
	v843 = v769
	v845 = v572
	goto L147
L237:
	;
	F_errmsg_internal(m, int32(_a_F_transformColumnDefinition_16), int32(0))
	mBase = m.M
	v796 = m.ExcPending
	if v796 != 0 {
		goto L1
	} else {
		goto L238
	}
L238:
	;
	F_errfinish(m, int32(_a_F_transformColumnDefinition_9), int32(940), int32(_a_F_transformColumnDefinition_10))
	mBase = m.M
	v801 = m.ExcPending
	if v801 != 0 {
		goto L1
	} else {
		goto L239
	}
L239:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L240:
	;
	v805 = *(*int32)(unsafe.Add(mBase, uint32(v532)))
	v806 = F_makeString(m, v805)
	mBase = m.M
	v807 = m.ExcPending
	if v807 != 0 {
		goto L1
	} else {
		goto L241
	}
L241:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+216)) = v806
	*(*int32)(unsafe.Add(mBase, uint32(v21)+232)) = v806
	v813 = F_list_make1_impl(m, int32(1), v21+int32(216))
	mBase = m.M
	v814 = m.ExcPending
	if v814 != 0 {
		goto L1
	} else {
		goto L242
	}
L242:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v578)+76)) = v813
	v816 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v817 = F_lappend(m, v816, v578)
	mBase = m.M
	v818 = m.ExcPending
	if v818 != 0 {
		goto L1
	} else {
		goto L243
	}
L243:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v817
	v838 = v561
	v839 = v564
	v840 = v565
	v842 = v569
	v843 = v570
	v845 = v572
	goto L147
L244:
	;
	v824 = *(*int32)(unsafe.Add(mBase, uint32(v578)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = v824
	F_errmsg_internal(m, int32(_a_F_transformColumnDefinition_17), v21)
	mBase = m.M
	v828 = m.ExcPending
	if v828 != 0 {
		goto L1
	} else {
		goto L245
	}
L245:
	;
	F_errfinish(m, int32(_a_F_transformColumnDefinition_9), int32(970), int32(_a_F_transformColumnDefinition_10))
	mBase = m.M
	v833 = m.ExcPending
	if v833 != 0 {
		goto L1
	} else {
		goto L246
	}
L246:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L247:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v835
	v838 = v561
	v839 = v564
	v840 = v565
	v842 = v569
	v843 = v570
	v845 = v572
	goto L147
L248:
	;
	if v838&v845 != 0 {
		goto L129
	} else {
		goto L249
	}
L249:
	;
	if v838&v842&int32(1) == int32(0) {
		goto L250
	} else {
		goto L251
	}
L250:
	;
	v856 = v562 + int32(1)
	v857 = *(*int32)(unsafe.Add(mBase, uint32(v539)+4))
	if v857 <= v856 {
		v1261 = v839
		v1267 = v843
		goto L124
	} else {
		goto L253
	}
L251:
	;
	goto L252
L252:
	;
	goto L146
L253:
	;
	v561 = v838
	v562 = v856
	v564 = v839
	v565 = v840
	v569 = v842
	v570 = v843
	v572 = v845
	goto L145
L254:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v865 = m.ExcPending
	if v865 != 0 {
		goto L1
	} else {
		goto L255
	}
L255:
	;
	v866 = *(*int32)(unsafe.Add(mBase, uint32(v532)))
	v867 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v868 = *(*int32)(unsafe.Add(mBase, uint32(v867)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+68)) = v868
	*(*int32)(unsafe.Add(mBase, uint32(v21)+64)) = v866
	F_errmsg(m, int32(_a_F_transformColumnDefinition_18), v21-int32(-64))
	mBase = m.M
	v875 = m.ExcPending
	if v875 != 0 {
		goto L1
	} else {
		goto L256
	}
L256:
	;
	v876 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v877 = *(*int32)(unsafe.Add(mBase, uint32(v578)+104))
	F_parser_errposition(m, v876, v877)
	mBase = m.M
	v879 = m.ExcPending
	if v879 != 0 {
		goto L1
	} else {
		goto L257
	}
L257:
	;
	F_errfinish(m, int32(_a_F_transformColumnDefinition_9), int32(996), int32(_a_F_transformColumnDefinition_10))
	mBase = m.M
	v884 = m.ExcPending
	if v884 != 0 {
		goto L1
	} else {
		goto L258
	}
L258:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L259:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v891 = m.ExcPending
	if v891 != 0 {
		goto L1
	} else {
		goto L260
	}
L260:
	;
	F_errmsg(m, int32(_a_F_transformColumnDefinition_19), int32(0))
	mBase = m.M
	v895 = m.ExcPending
	if v895 != 0 {
		goto L1
	} else {
		goto L261
	}
L261:
	;
	F_errfinish(m, int32(_a_F_transformColumnDefinition_9), int32(765), int32(_a_F_transformColumnDefinition_10))
	mBase = m.M
	v900 = m.ExcPending
	if v900 != 0 {
		goto L1
	} else {
		goto L262
	}
L262:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L263:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v907 = m.ExcPending
	if v907 != 0 {
		goto L1
	} else {
		goto L264
	}
L264:
	;
	v908 = *(*int32)(unsafe.Add(mBase, uint32(v532)))
	v909 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v910 = *(*int32)(unsafe.Add(mBase, uint32(v909)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+132)) = v910
	*(*int32)(unsafe.Add(mBase, uint32(v21)+128)) = v908
	F_errmsg(m, int32(_a_F_transformColumnDefinition_14), v21+int32(128))
	mBase = m.M
	v917 = m.ExcPending
	if v917 != 0 {
		goto L1
	} else {
		goto L265
	}
L265:
	;
	v918 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v919 = *(*int32)(unsafe.Add(mBase, uint32(v578)+104))
	F_parser_errposition(m, v918, v919)
	mBase = m.M
	v921 = m.ExcPending
	if v921 != 0 {
		goto L1
	} else {
		goto L266
	}
L266:
	;
	F_errfinish(m, int32(_a_F_transformColumnDefinition_9), int32(774), int32(_a_F_transformColumnDefinition_10))
	mBase = m.M
	v926 = m.ExcPending
	if v926 != 0 {
		goto L1
	} else {
		goto L267
	}
L267:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L268:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v933 = m.ExcPending
	if v933 != 0 {
		goto L1
	} else {
		goto L269
	}
L269:
	;
	v934 = *(*int32)(unsafe.Add(mBase, uint32(v532)))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+80)) = v934
	F_errmsg(m, int32(_a_F_transformColumnDefinition_20), v21+int32(80))
	mBase = m.M
	v940 = m.ExcPending
	if v940 != 0 {
		goto L1
	} else {
		goto L270
	}
L270:
	;
	F_errfinish(m, int32(_a_F_transformColumnDefinition_9), int32(780), int32(_a_F_transformColumnDefinition_10))
	mBase = m.M
	v945 = m.ExcPending
	if v945 != 0 {
		goto L1
	} else {
		goto L271
	}
L271:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L272:
	;
	v950 = *(*int32)(unsafe.Add(mBase, uint32(v565)+8))
	v951 = *(*int32)(unsafe.Add(mBase, uint32(v578)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+116)) = v951
	*(*int32)(unsafe.Add(mBase, uint32(v21)+112)) = v950
	F_errmsg_internal(m, int32(_a_F_transformColumnDefinition_21), v21+int32(112))
	mBase = m.M
	v958 = m.ExcPending
	if v958 != 0 {
		goto L1
	} else {
		goto L273
	}
L273:
	;
	F_errfinish(m, int32(_a_F_transformColumnDefinition_9), int32(809), int32(_a_F_transformColumnDefinition_10))
	mBase = m.M
	v963 = m.ExcPending
	if v963 != 0 {
		goto L1
	} else {
		goto L274
	}
L274:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L275:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v970 = m.ExcPending
	if v970 != 0 {
		goto L1
	} else {
		goto L276
	}
L276:
	;
	v971 = *(*int32)(unsafe.Add(mBase, uint32(v532)))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+96)) = v971
	F_errmsg(m, int32(_a_F_transformColumnDefinition_20), v21+int32(96))
	mBase = m.M
	v977 = m.ExcPending
	if v977 != 0 {
		goto L1
	} else {
		goto L277
	}
L277:
	;
	F_errfinish(m, int32(_a_F_transformColumnDefinition_9), int32(815), int32(_a_F_transformColumnDefinition_10))
	mBase = m.M
	v982 = m.ExcPending
	if v982 != 0 {
		goto L1
	} else {
		goto L278
	}
L278:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L279:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v989 = m.ExcPending
	if v989 != 0 {
		goto L1
	} else {
		goto L280
	}
L280:
	;
	v990 = *(*int32)(unsafe.Add(mBase, uint32(v532)))
	v991 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v992 = *(*int32)(unsafe.Add(mBase, uint32(v991)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+148)) = v992
	*(*int32)(unsafe.Add(mBase, uint32(v21)+144)) = v990
	F_errmsg(m, int32(_a_F_transformColumnDefinition_22), v21+int32(144))
	mBase = m.M
	v999 = m.ExcPending
	if v999 != 0 {
		goto L1
	} else {
		goto L281
	}
L281:
	;
	v1000 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1001 = *(*int32)(unsafe.Add(mBase, uint32(v578)+104))
	F_parser_errposition(m, v1000, v1001)
	mBase = m.M
	v1003 = m.ExcPending
	if v1003 != 0 {
		goto L1
	} else {
		goto L282
	}
L282:
	;
	F_errfinish(m, int32(_a_F_transformColumnDefinition_9), int32(830), int32(_a_F_transformColumnDefinition_10))
	mBase = m.M
	v1008 = m.ExcPending
	if v1008 != 0 {
		goto L1
	} else {
		goto L283
	}
L283:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L284:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v1015 = m.ExcPending
	if v1015 != 0 {
		goto L1
	} else {
		goto L285
	}
L285:
	;
	F_errmsg(m, int32(_a_F_transformColumnDefinition_23), int32(0))
	mBase = m.M
	v1019 = m.ExcPending
	if v1019 != 0 {
		goto L1
	} else {
		goto L286
	}
L286:
	;
	F_errfinish(m, int32(_a_F_transformColumnDefinition_9), int32(844), int32(_a_F_transformColumnDefinition_10))
	mBase = m.M
	v1024 = m.ExcPending
	if v1024 != 0 {
		goto L1
	} else {
		goto L287
	}
L287:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L288:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v1031 = m.ExcPending
	if v1031 != 0 {
		goto L1
	} else {
		goto L289
	}
L289:
	;
	F_errmsg(m, int32(_a_F_transformColumnDefinition_24), int32(0))
	mBase = m.M
	v1035 = m.ExcPending
	if v1035 != 0 {
		goto L1
	} else {
		goto L290
	}
L290:
	;
	F_errfinish(m, int32(_a_F_transformColumnDefinition_9), int32(848), int32(_a_F_transformColumnDefinition_10))
	mBase = m.M
	v1040 = m.ExcPending
	if v1040 != 0 {
		goto L1
	} else {
		goto L291
	}
L291:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L292:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v1047 = m.ExcPending
	if v1047 != 0 {
		goto L1
	} else {
		goto L293
	}
L293:
	;
	v1048 = *(*int32)(unsafe.Add(mBase, uint32(v532)))
	v1049 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1050 = *(*int32)(unsafe.Add(mBase, uint32(v1049)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+164)) = v1050
	*(*int32)(unsafe.Add(mBase, uint32(v21)+160)) = v1048
	F_errmsg(m, int32(_a_F_transformColumnDefinition_25), v21+int32(160))
	mBase = m.M
	v1057 = m.ExcPending
	if v1057 != 0 {
		goto L1
	} else {
		goto L294
	}
L294:
	;
	v1058 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1059 = *(*int32)(unsafe.Add(mBase, uint32(v578)+104))
	F_parser_errposition(m, v1058, v1059)
	mBase = m.M
	v1061 = m.ExcPending
	if v1061 != 0 {
		goto L1
	} else {
		goto L295
	}
L295:
	;
	F_errfinish(m, int32(_a_F_transformColumnDefinition_9), int32(860), int32(_a_F_transformColumnDefinition_10))
	mBase = m.M
	v1066 = m.ExcPending
	if v1066 != 0 {
		goto L1
	} else {
		goto L296
	}
L296:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L297:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v1073 = m.ExcPending
	if v1073 != 0 {
		goto L1
	} else {
		goto L298
	}
L298:
	;
	F_errmsg(m, int32(_a_F_transformColumnDefinition_26), int32(0))
	mBase = m.M
	v1077 = m.ExcPending
	if v1077 != 0 {
		goto L1
	} else {
		goto L299
	}
L299:
	;
	F_errfinish(m, int32(_a_F_transformColumnDefinition_9), int32(890), int32(_a_F_transformColumnDefinition_10))
	mBase = m.M
	v1082 = m.ExcPending
	if v1082 != 0 {
		goto L1
	} else {
		goto L300
	}
L300:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L301:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v1089 = m.ExcPending
	if v1089 != 0 {
		goto L1
	} else {
		goto L302
	}
L302:
	;
	v1090 = *(*int32)(unsafe.Add(mBase, uint32(v532)))
	v1091 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1092 = *(*int32)(unsafe.Add(mBase, uint32(v1091)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+196)) = v1092
	*(*int32)(unsafe.Add(mBase, uint32(v21)+192)) = v1090
	F_errmsg(m, int32(_a_F_transformColumnDefinition_27), v21+int32(192))
	mBase = m.M
	v1099 = m.ExcPending
	if v1099 != 0 {
		goto L1
	} else {
		goto L303
	}
L303:
	;
	v1100 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1101 = *(*int32)(unsafe.Add(mBase, uint32(v578)+104))
	F_parser_errposition(m, v1100, v1101)
	mBase = m.M
	v1103 = m.ExcPending
	if v1103 != 0 {
		goto L1
	} else {
		goto L304
	}
L304:
	;
	F_errfinish(m, int32(_a_F_transformColumnDefinition_9), int32(897), int32(_a_F_transformColumnDefinition_10))
	mBase = m.M
	v1108 = m.ExcPending
	if v1108 != 0 {
		goto L1
	} else {
		goto L305
	}
L305:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L306:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v1115 = m.ExcPending
	if v1115 != 0 {
		goto L1
	} else {
		goto L307
	}
L307:
	;
	v1116 = *(*int32)(unsafe.Add(mBase, uint32(v532)))
	v1117 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1118 = *(*int32)(unsafe.Add(mBase, uint32(v1117)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+212)) = v1118
	*(*int32)(unsafe.Add(mBase, uint32(v21)+208)) = v1116
	F_errmsg(m, int32(_a_F_transformColumnDefinition_14), v21+int32(208))
	mBase = m.M
	v1125 = m.ExcPending
	if v1125 != 0 {
		goto L1
	} else {
		goto L308
	}
L308:
	;
	v1126 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1127 = *(*int32)(unsafe.Add(mBase, uint32(v578)+104))
	F_parser_errposition(m, v1126, v1127)
	mBase = m.M
	v1129 = m.ExcPending
	if v1129 != 0 {
		goto L1
	} else {
		goto L309
	}
L309:
	;
	F_errfinish(m, int32(_a_F_transformColumnDefinition_9), int32(915), int32(_a_F_transformColumnDefinition_10))
	mBase = m.M
	v1134 = m.ExcPending
	if v1134 != 0 {
		goto L1
	} else {
		goto L310
	}
L310:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L311:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v1141 = m.ExcPending
	if v1141 != 0 {
		goto L1
	} else {
		goto L312
	}
L312:
	;
	F_errmsg(m, int32(_a_F_transformColumnDefinition_28), int32(0))
	mBase = m.M
	v1145 = m.ExcPending
	if v1145 != 0 {
		goto L1
	} else {
		goto L313
	}
L313:
	;
	v1146 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1147 = *(*int32)(unsafe.Add(mBase, uint32(v578)+104))
	F_parser_errposition(m, v1146, v1147)
	mBase = m.M
	v1149 = m.ExcPending
	if v1149 != 0 {
		goto L1
	} else {
		goto L314
	}
L314:
	;
	F_errfinish(m, int32(_a_F_transformColumnDefinition_9), int32(932), int32(_a_F_transformColumnDefinition_10))
	mBase = m.M
	v1154 = m.ExcPending
	if v1154 != 0 {
		goto L1
	} else {
		goto L315
	}
L315:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L316:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v1161 = m.ExcPending
	if v1161 != 0 {
		goto L1
	} else {
		goto L317
	}
L317:
	;
	F_errmsg(m, int32(_a_F_transformColumnDefinition_29), int32(0))
	mBase = m.M
	v1165 = m.ExcPending
	if v1165 != 0 {
		goto L1
	} else {
		goto L318
	}
L318:
	;
	v1166 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1167 = *(*int32)(unsafe.Add(mBase, uint32(v578)+104))
	F_parser_errposition(m, v1166, v1167)
	mBase = m.M
	v1169 = m.ExcPending
	if v1169 != 0 {
		goto L1
	} else {
		goto L319
	}
L319:
	;
	F_errfinish(m, int32(_a_F_transformColumnDefinition_9), int32(949), int32(_a_F_transformColumnDefinition_10))
	mBase = m.M
	v1174 = m.ExcPending
	if v1174 != 0 {
		goto L1
	} else {
		goto L320
	}
L320:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L321:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v1181 = m.ExcPending
	if v1181 != 0 {
		goto L1
	} else {
		goto L322
	}
L322:
	;
	v1182 = *(*int32)(unsafe.Add(mBase, uint32(v532)))
	v1183 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1184 = *(*int32)(unsafe.Add(mBase, uint32(v1183)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+36)) = v1184
	*(*int32)(unsafe.Add(mBase, uint32(v21)+32)) = v1182
	F_errmsg(m, int32(_a_F_transformColumnDefinition_30), v21+int32(32))
	mBase = m.M
	v1191 = m.ExcPending
	if v1191 != 0 {
		goto L1
	} else {
		goto L323
	}
L323:
	;
	v1192 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1193 = *(*int32)(unsafe.Add(mBase, uint32(v578)+104))
	F_parser_errposition(m, v1192, v1193)
	mBase = m.M
	v1195 = m.ExcPending
	if v1195 != 0 {
		goto L1
	} else {
		goto L324
	}
L324:
	;
	F_errfinish(m, int32(_a_F_transformColumnDefinition_9), int32(980), int32(_a_F_transformColumnDefinition_10))
	mBase = m.M
	v1200 = m.ExcPending
	if v1200 != 0 {
		goto L1
	} else {
		goto L325
	}
L325:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L326:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v1207 = m.ExcPending
	if v1207 != 0 {
		goto L1
	} else {
		goto L327
	}
L327:
	;
	v1208 = *(*int32)(unsafe.Add(mBase, uint32(v532)))
	v1209 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1210 = *(*int32)(unsafe.Add(mBase, uint32(v1209)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+52)) = v1210
	*(*int32)(unsafe.Add(mBase, uint32(v21)+48)) = v1208
	F_errmsg(m, int32(_a_F_transformColumnDefinition_31), v21+int32(48))
	mBase = m.M
	v1217 = m.ExcPending
	if v1217 != 0 {
		goto L1
	} else {
		goto L328
	}
L328:
	;
	v1218 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1219 = *(*int32)(unsafe.Add(mBase, uint32(v578)+104))
	F_parser_errposition(m, v1218, v1219)
	mBase = m.M
	v1221 = m.ExcPending
	if v1221 != 0 {
		goto L1
	} else {
		goto L329
	}
L329:
	;
	F_errfinish(m, int32(_a_F_transformColumnDefinition_9), int32(988), int32(_a_F_transformColumnDefinition_10))
	mBase = m.M
	v1226 = m.ExcPending
	if v1226 != 0 {
		goto L1
	} else {
		goto L330
	}
L330:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L331:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v1233 = m.ExcPending
	if v1233 != 0 {
		goto L1
	} else {
		goto L332
	}
L332:
	;
	v1234 = *(*int32)(unsafe.Add(mBase, uint32(v532)))
	v1235 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1236 = *(*int32)(unsafe.Add(mBase, uint32(v1235)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+20)) = v1236
	*(*int32)(unsafe.Add(mBase, uint32(v21)+16)) = v1234
	F_errmsg(m, int32(_a_F_transformColumnDefinition_14), v21+int32(16))
	mBase = m.M
	v1243 = m.ExcPending
	if v1243 != 0 {
		goto L1
	} else {
		goto L333
	}
L333:
	;
	v1244 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1245 = *(*int32)(unsafe.Add(mBase, uint32(v578)+104))
	F_parser_errposition(m, v1244, v1245)
	mBase = m.M
	v1247 = m.ExcPending
	if v1247 != 0 {
		goto L1
	} else {
		goto L334
	}
L334:
	;
	F_errfinish(m, int32(_a_F_transformColumnDefinition_9), int32(756), int32(_a_F_transformColumnDefinition_10))
	mBase = m.M
	v1252 = m.ExcPending
	if v1252 != 0 {
		goto L1
	} else {
		goto L335
	}
L335:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L336:
	;
	if v1261 == int32(0) {
		v1280 = v532
		v1283 = v535
		v1288 = v540
		goto L4
	} else {
		goto L337
	}
L337:
	;
	v1275 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v535))))
	if v1275 != 0 {
		goto L3
	} else {
		goto L338
	}
L338:
	;
	v1280 = v532
	v1283 = v535
	v1288 = v540
	goto L4
L339:
	;
	v1299 = F_makeNotNullConstraint(m, v1297)
	mBase = m.M
	v1300 = m.ExcPending
	if v1300 != 0 {
		goto L1
	} else {
		goto L340
	}
L340:
	;
	v1301 = *(*int32)(unsafe.Add(mBase, uint32(v1288)))
	v1302 = F_lappend(m, v1301, v1299)
	mBase = m.M
	v1303 = m.ExcPending
	if v1303 != 0 {
		goto L1
	} else {
		goto L341
	}
L341:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1288))) = v1302
	goto L3
L342:
	;
	v1325 = F_palloc0(m, int32(32))
	mBase = m.M
	v1326 = m.ExcPending
	if v1326 != 0 {
		goto L1
	} else {
		goto L345
	}
L343:
	;
	goto L344
L344:
	;
	m.G0 = v21 + int32(256)
	return
L345:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1325))) = int64(107374182547)
	v1329 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1325)+8)) = v1329
	v1331 = *(*int32)(unsafe.Add(mBase, uint32(l1)+60))
	v1332 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1325)+28)) = uint8(v1332)
	*(*int32)(unsafe.Add(mBase, uint32(v1325)+24)) = v1332
	*(*int32)(unsafe.Add(mBase, uint32(v1325)+20)) = v1331
	v1338 = F_palloc0(m, int32(20))
	mBase = m.M
	v1339 = m.ExcPending
	if v1339 != 0 {
		goto L1
	} else {
		goto L346
	}
L346:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1338))) = int32(146)
	v1342 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v1338)+8)) = int64(77309411328)
	*(*int32)(unsafe.Add(mBase, uint32(v1338)+4)) = v1342
	v1347 = F_lappend(m, int32(0), v1325)
	mBase = m.M
	v1348 = m.ExcPending
	if v1348 != 0 {
		goto L1
	} else {
		goto L347
	}
L347:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1338)+8)) = v1347
	v1350 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v1351 = F_lappend(m, v1350, v1338)
	mBase = m.M
	v1352 = m.ExcPending
	if v1352 != 0 {
		goto L1
	} else {
		goto L348
	}
L348:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v1351
	goto L344
}
func F_transformColumnNameList(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v75 int32
	_ = v75
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
	v6 = int32(0)
	v12 = m.G0
	v14 = v12 - int32(32)
	m.G0 = v14
	if l1 == v6 {
		v75 = v6
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L9
	} else {
		goto L30
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L9
	} else {
		goto L26
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L9
	} else {
		goto L22
	}
L4:
	;
	m.G0 = v14 + int32(32)
	return v75
L5:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v18 <= int32(0) {
		v75 = v6
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v28 = v6
	goto L7
L7:
	;
	v33 = v28 << (uint(int32(2)) % 32)
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v33+v34)))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)+4))
	v38 = F_SearchSysCacheAttName(m, l0, v37)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v75 = v65
	goto L4
L9:
	;
	return int32(0)
L10:
	;
	if v38 == int32(0) {
		goto L3
	} else {
		goto L11
	}
L11:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v38)+16))
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+22)))
	v46 = v44 + v45
	v47 = int32(*(*int16)(unsafe.Add(mBase, uint32(v46)+74)))
	if v47 < int32(0) {
		goto L2
	} else {
		goto L12
	}
L12:
	;
	if v28 == int32(32) {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(l2+v28<<(uint(int32(1))%32)))) = uint16(v47)
	if l3 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v46)+68))
	*(*int32)(unsafe.Add(mBase, uint32(l3+v33))) = v57
	goto L16
L15:
	;
	goto L16
L16:
	;
	if l4 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v46)+96))
	*(*int32)(unsafe.Add(mBase, uint32(l4+v33))) = v60
	goto L19
L18:
	;
	goto L19
L19:
	;
	F_ReleaseCatCache(m, v38)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L9
	} else {
		goto L20
	}
L20:
	;
	v65 = v28 + int32(1)
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v65 < v66 {
		v28 = v65
		goto L7
	} else {
		goto L21
	}
L21:
	;
	goto L8
L22:
	;
	F_errcode(m, int32(50360452))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L9
	} else {
		goto L23
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v37
	F_errmsg(m, int32(_a_F_transformColumnNameList_0), v14)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L9
	} else {
		goto L24
	}
L24:
	;
	F_errfinish(m, int32(_a_F_transformColumnNameList_1), int32(_a_F_transformColumnNameList_2), int32(_a_F_transformColumnNameList_3))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L9
	} else {
		goto L25
	}
L25:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L26:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L9
	} else {
		goto L27
	}
L27:
	;
	F_errmsg(m, int32(_a_F_transformColumnNameList_4), int32(0))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L9
	} else {
		goto L28
	}
L28:
	;
	F_errfinish(m, int32(_a_F_transformColumnNameList_1), int32(_a_F_transformColumnNameList_5), int32(_a_F_transformColumnNameList_3))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L9
	} else {
		goto L29
	}
L29:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L30:
	;
	F_errcode(m, int32(17039621))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L9
	} else {
		goto L31
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = int32(32)
	F_errmsg(m, int32(_a_F_transformColumnNameList_6), v14+int32(16))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L9
	} else {
		goto L32
	}
L32:
	;
	F_errfinish(m, int32(_a_F_transformColumnNameList_1), int32(_a_F_transformColumnNameList_7), int32(_a_F_transformColumnNameList_3))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L9
	} else {
		goto L33
	}
L33:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
