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
	v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
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
	F_errfinish(m, int32(_a_F_ExpandColumnRefStar_1), int32(1331), int32(_a_F_ExpandColumnRefStar_2))
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
	v168 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
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
	F_errfinish(m, int32(_a_F_ExpandColumnRefStar_1), int32(1243), int32(_a_F_ExpandColumnRefStar_4))
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
	F_errfinish(m, int32(_a_F_ExpandColumnRefStar_1), int32(1264), int32(_a_F_ExpandColumnRefStar_4))
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
	F_errfinish(m, int32(_a_F_ExpandColumnRefStar_1), int32(1271), int32(_a_F_ExpandColumnRefStar_4))
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
func F_has_column_privilege_id_attnum(m *base.Module, l0 int32) int32 {
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
	var v48 int32
	_ = v48
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v15 = F_pg_detoast_datum_packed(m, v14)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		v20 = *(*int32)(unsafe.Add(mBase, _c_F_has_column_privilege_id_attnum[0]))
		v22 = F_convert_any_priv_string(m, v15, int32(_a_F_has_column_privilege_id_attnum_0))
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return int32(0)
		} else {
			v24 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)) = uint8(v24)
			if v13&int32(_a_F_has_column_privilege_id_attnum_1) == v24 {
				v43 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v43)
				v48 = int32(0)
				m.G0 = v10 + int32(16)
				return v48
			} else {
				v32 = v10 + int32(15)
				v33 = F_pg_attribute_aclcheck_ext(m, v12, base.I32_extend16_s(v13), v20, v22, v32)
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return int32(0)
				} else {
					if v33 != 0 {
						v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)))
						if v35 != 0 {
							v43 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v43)
							v48 = int32(0)
							m.G0 = v10 + int32(16)
							return v48
						} else {
							v36 = F_pg_class_aclcheck_ext(m, v12, v20, v22, v32)
							mBase = m.M
							v37 = m.ExcPending
							if v37 != 0 {
								return int32(0)
							} else {
								if v36 != 0 {
									v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)))
									if v39 == int32(0) {
									} else {
										v43 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v43)
									}
									v48 = int32(0)
								} else {
									v48 = int32(1)
								}
								m.G0 = v10 + int32(16)
								return v48
							}
						}
					} else {
						v48 = int32(1)
						m.G0 = v10 + int32(16)
						return v48
					}
				}
			}
		}
	}
}
func F_has_column_privilege_id_id_name(m *base.Module, l0 int32) int32 {
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
	var v48 int32
	_ = v48
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v15 = F_pg_detoast_datum_packed(m, v14)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
		v20 = F_pg_detoast_datum_packed(m, v19)
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return int32(0)
		} else {
			v22 = F_convert_column_name(m, v13, v15)
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int32(0)
			} else {
				v25 = F_convert_any_priv_string(m, v20, int32(_a_F_has_column_privilege_id_id_name_0))
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int32(0)
				} else {
					v27 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)) = uint8(v27)
					if v22 == v27 {
						v43 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v43)
						v48 = int32(0)
						m.G0 = v10 + int32(16)
						return v48
					} else {
						v32 = v10 + int32(15)
						v33 = F_pg_attribute_aclcheck_ext(m, v13, v22, v12, v25, v32)
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
							return int32(0)
						} else {
							if v33 != 0 {
								v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)))
								if v35 != 0 {
									v43 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v43)
									v48 = int32(0)
									m.G0 = v10 + int32(16)
									return v48
								} else {
									v36 = F_pg_class_aclcheck_ext(m, v13, v12, v25, v32)
									mBase = m.M
									v37 = m.ExcPending
									if v37 != 0 {
										return int32(0)
									} else {
										if v36 != 0 {
											v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)))
											if v39 == int32(0) {
											} else {
												v43 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v43)
											}
											v48 = int32(0)
										} else {
											v48 = int32(1)
										}
										m.G0 = v10 + int32(16)
										return v48
									}
								}
							} else {
								v48 = int32(1)
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
func F_has_column_privilege_id_name(m *base.Module, l0 int32) int32 {
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
	var v49 int32
	_ = v49
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v14 = F_pg_detoast_datum_packed(m, v13)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
		v19 = F_pg_detoast_datum_packed(m, v18)
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int32(0)
		} else {
			v22 = *(*int32)(unsafe.Add(mBase, _c_F_has_column_privilege_id_name[0]))
			v23 = F_convert_column_name(m, v12, v14)
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int32(0)
			} else {
				v26 = F_convert_any_priv_string(m, v19, int32(_a_F_has_column_privilege_id_name_0))
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return int32(0)
				} else {
					v28 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)) = uint8(v28)
					if v23 == v28 {
						v44 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v44)
						v49 = int32(0)
						m.G0 = v10 + int32(16)
						return v49
					} else {
						v33 = v10 + int32(15)
						v34 = F_pg_attribute_aclcheck_ext(m, v12, v23, v22, v26, v33)
						mBase = m.M
						v35 = m.ExcPending
						if v35 != 0 {
							return int32(0)
						} else {
							if v34 != 0 {
								v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)))
								if v36 != 0 {
									v44 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v44)
									v49 = int32(0)
									m.G0 = v10 + int32(16)
									return v49
								} else {
									v37 = F_pg_class_aclcheck_ext(m, v12, v22, v26, v33)
									mBase = m.M
									v38 = m.ExcPending
									if v38 != 0 {
										return int32(0)
									} else {
										if v37 != 0 {
											v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)))
											if v40 == int32(0) {
											} else {
												v44 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v44)
											}
											v49 = int32(0)
										} else {
											v49 = int32(1)
										}
										m.G0 = v10 + int32(16)
										return v49
									}
								}
							} else {
								v49 = int32(1)
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
func F_has_column_privilege_id_name_attnum(m *base.Module, l0 int32) int32 {
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
	var v43 int32
	_ = v43
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
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v14 = F_pg_detoast_datum_packed(m, v13)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
		v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
		v20 = F_pg_detoast_datum_packed(m, v19)
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return int32(0)
		} else {
			v22 = F_textToQualifiedNameList(m, v14)
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int32(0)
			} else {
				v24 = F_makeRangeVarFromNameList(m, v22)
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					v26 = int32(0)
					v30 = F_RangeVarGetRelidExtended(m, v24, v26, v26, v26, v26)
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return int32(0)
					} else {
						v33 = F_convert_any_priv_string(m, v20, int32(_a_F_has_column_privilege_id_name_attnum_0))
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
							return int32(0)
						} else {
							v35 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)) = uint8(v35)
							if v18&int32(_a_F_has_column_privilege_id_name_attnum_1) == v35 {
								v54 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v54)
								v59 = int32(0)
								m.G0 = v10 + int32(16)
								return v59
							} else {
								v43 = v10 + int32(15)
								v44 = F_pg_attribute_aclcheck_ext(m, v30, base.I32_extend16_s(v18), v12, v33, v43)
								mBase = m.M
								v45 = m.ExcPending
								if v45 != 0 {
									return int32(0)
								} else {
									if v44 != 0 {
										v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)))
										if v46 != 0 {
											v54 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v54)
											v59 = int32(0)
											m.G0 = v10 + int32(16)
											return v59
										} else {
											v47 = F_pg_class_aclcheck_ext(m, v30, v12, v33, v43)
											mBase = m.M
											v48 = m.ExcPending
											if v48 != 0 {
												return int32(0)
											} else {
												if v47 != 0 {
													v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)))
													if v50 == int32(0) {
													} else {
														v54 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v54)
													}
													v59 = int32(0)
												} else {
													v59 = int32(1)
												}
												m.G0 = v10 + int32(16)
												return v59
											}
										}
									} else {
										v59 = int32(1)
										m.G0 = v10 + int32(16)
										return v59
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
func F_has_column_privilege_id_name_name(m *base.Module, l0 int32) int32 {
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
	var v60 int32
	_ = v60
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v14 = F_pg_detoast_datum_packed(m, v13)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
		v19 = F_pg_detoast_datum_packed(m, v18)
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int32(0)
		} else {
			v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
			v22 = F_pg_detoast_datum_packed(m, v21)
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int32(0)
			} else {
				v24 = F_textToQualifiedNameList(m, v14)
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					v26 = F_makeRangeVarFromNameList(m, v24)
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return int32(0)
					} else {
						v28 = int32(0)
						v32 = F_RangeVarGetRelidExtended(m, v26, v28, v28, v28, v28)
						mBase = m.M
						v33 = m.ExcPending
						if v33 != 0 {
							return int32(0)
						} else {
							v34 = F_convert_column_name(m, v32, v19)
							mBase = m.M
							v35 = m.ExcPending
							if v35 != 0 {
								return int32(0)
							} else {
								v37 = F_convert_any_priv_string(m, v22, int32(_a_F_has_column_privilege_id_name_name_0))
								mBase = m.M
								v38 = m.ExcPending
								if v38 != 0 {
									return int32(0)
								} else {
									v39 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)) = uint8(v39)
									if v34 == v39 {
										v55 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v55)
										v60 = int32(0)
										m.G0 = v10 + int32(16)
										return v60
									} else {
										v44 = v10 + int32(15)
										v45 = F_pg_attribute_aclcheck_ext(m, v32, v34, v12, v37, v44)
										mBase = m.M
										v46 = m.ExcPending
										if v46 != 0 {
											return int32(0)
										} else {
											if v45 != 0 {
												v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)))
												if v47 != 0 {
													v55 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v55)
													v60 = int32(0)
													m.G0 = v10 + int32(16)
													return v60
												} else {
													v48 = F_pg_class_aclcheck_ext(m, v32, v12, v37, v44)
													mBase = m.M
													v49 = m.ExcPending
													if v49 != 0 {
														return int32(0)
													} else {
														if v48 != 0 {
															v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)))
															if v51 == int32(0) {
															} else {
																v55 = int32(1)
																*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v55)
															}
															v60 = int32(0)
														} else {
															v60 = int32(1)
														}
														m.G0 = v10 + int32(16)
														return v60
													}
												}
											} else {
												v60 = int32(1)
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
func F_has_column_privilege_name_name(m *base.Module, l0 int32) int32 {
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
	var v61 int32
	_ = v61
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = F_pg_detoast_datum_packed(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v18 = F_pg_detoast_datum_packed(m, v17)
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int32(0)
		} else {
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
			v21 = F_pg_detoast_datum_packed(m, v20)
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int32(0)
			} else {
				v24 = *(*int32)(unsafe.Add(mBase, _c_F_has_column_privilege_name_name[0]))
				v25 = F_textToQualifiedNameList(m, v13)
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int32(0)
				} else {
					v27 = F_makeRangeVarFromNameList(m, v25)
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return int32(0)
					} else {
						v29 = int32(0)
						v33 = F_RangeVarGetRelidExtended(m, v27, v29, v29, v29, v29)
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
							return int32(0)
						} else {
							v35 = F_convert_column_name(m, v33, v18)
							mBase = m.M
							v36 = m.ExcPending
							if v36 != 0 {
								return int32(0)
							} else {
								v38 = F_convert_any_priv_string(m, v21, int32(_a_F_has_column_privilege_name_name_0))
								mBase = m.M
								v39 = m.ExcPending
								if v39 != 0 {
									return int32(0)
								} else {
									v40 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)) = uint8(v40)
									if v35 == v40 {
										v56 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v56)
										v61 = int32(0)
										m.G0 = v10 + int32(16)
										return v61
									} else {
										v45 = v10 + int32(15)
										v46 = F_pg_attribute_aclcheck_ext(m, v33, v35, v24, v38, v45)
										mBase = m.M
										v47 = m.ExcPending
										if v47 != 0 {
											return int32(0)
										} else {
											if v46 != 0 {
												v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)))
												if v48 != 0 {
													v56 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v56)
													v61 = int32(0)
													m.G0 = v10 + int32(16)
													return v61
												} else {
													v49 = F_pg_class_aclcheck_ext(m, v33, v24, v38, v45)
													mBase = m.M
													v50 = m.ExcPending
													if v50 != 0 {
														return int32(0)
													} else {
														if v49 != 0 {
															v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)))
															if v52 == int32(0) {
															} else {
																v56 = int32(1)
																*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v56)
															}
															v61 = int32(0)
														} else {
															v61 = int32(1)
														}
														m.G0 = v10 + int32(16)
														return v61
													}
												}
											} else {
												v61 = int32(1)
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
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v296 int32
	_ = v296
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v332 int32
	_ = v332
	var v335 int32
	_ = v335
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v355 int32
	_ = v355
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v363 int32
	_ = v363
	var v366 int32
	_ = v366
	var v371 int32
	_ = v371
	var v378 int32
	_ = v378
	var v383 int32
	_ = v383
	var v390 int32
	_ = v390
	var v392 int32
	_ = v392
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v404 int32
	_ = v404
	var v407 int32
	_ = v407
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v415 int32
	_ = v415
	var v420 int32
	_ = v420
	var v423 int32
	_ = v423
	var v430 int32
	_ = v430
	var v436 int32
	_ = v436
	var v439 int32
	_ = v439
	var v441 int32
	_ = v441
	var v445 int32
	_ = v445
	var v448 int32
	_ = v448
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v456 int32
	_ = v456
	var v461 int32
	_ = v461
	var v464 int32
	_ = v464
	var v471 int32
	_ = v471
	var v476 int32
	_ = v476
	var v482 int32
	_ = v482
	var v485 int32
	_ = v485
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v493 int32
	_ = v493
	var v498 int32
	_ = v498
	var v501 int32
	_ = v501
	var v506 int32
	_ = v506
	var v512 int32
	_ = v512
	var v515 int32
	_ = v515
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v523 int32
	_ = v523
	var v528 int32
	_ = v528
	var v531 int32
	_ = v531
	var v533 int32
	_ = v533
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v550 int32
	_ = v550
	var v563 int32
	_ = v563
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v575 int32
	_ = v575
	var v578 int32
	_ = v578
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v587 int32
	_ = v587
	var v594 int32
	_ = v594
	var v596 int32
	_ = v596
	var v605 int32
	_ = v605
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v615 int32
	_ = v615
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v623 int32
	_ = v623
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v629 int32
	_ = v629
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v637 int32
	_ = v637
	var v648 int32
	_ = v648
	var v657 int32
	_ = v657
	var v659 int32
	_ = v659
	var v667 int32
	_ = v667
	var v673 int32
	_ = v673
	var v676 int32
	_ = v676
	var v681 int32
	_ = v681
	var v682 int32
	_ = v682
	var v686 int32
	_ = v686
	var v687 int32
	_ = v687
	var v690 int32
	_ = v690
	var v708 int32
	_ = v708
	var v711 int32
	_ = v711
	var v713 int32
	_ = v713
	var v715 int32
	_ = v715
	var v718 int32
	_ = v718
	var v719 int32
	_ = v719
	var v722 int32
	_ = v722
	var v731 int32
	_ = v731
	var v733 int32
	_ = v733
	var v734 int32
	_ = v734
	var v737 int32
	_ = v737
	var v738 int32
	_ = v738
	var v740 int32
	_ = v740
	var v742 int32
	_ = v742
	var v745 int32
	_ = v745
	var v749 int32
	_ = v749
	var v750 int32
	_ = v750
	var v753 int32
	_ = v753
	var v759 int32
	_ = v759
	var v762 int32
	_ = v762
	var v765 int32
	_ = v765
	var v770 int32
	_ = v770
	var v773 int32
	_ = v773
	var v776 int32
	_ = v776
	var v779 int32
	_ = v779
	var v782 int32
	_ = v782
	var v783 int32
	_ = v783
	var v784 int32
	_ = v784
	var v790 int32
	_ = v790
	var v791 int32
	_ = v791
	var v793 int32
	_ = v793
	var v794 int32
	_ = v794
	var v795 int32
	_ = v795
	var v801 int32
	_ = v801
	var v804 int32
	_ = v804
	var v809 int32
	_ = v809
	var v812 int32
	_ = v812
	var v815 int32
	_ = v815
	var v816 int32
	_ = v816
	var v819 int32
	_ = v819
	var v820 int32
	_ = v820
	var v823 int32
	_ = v823
	var v830 int32
	_ = v830
	var v831 int32
	_ = v831
	var v834 int32
	_ = v834
	var v835 int32
	_ = v835
	var v837 int32
	_ = v837
	var v844 int32
	_ = v844
	var v847 int32
	_ = v847
	var v850 int32
	_ = v850
	var v851 int32
	_ = v851
	var v852 int32
	_ = v852
	var v854 int32
	_ = v854
	var v855 int32
	_ = v855
	var v856 int32
	_ = v856
	var v857 int32
	_ = v857
	var v859 int32
	_ = v859
	var v861 int32
	_ = v861
	var v864 int32
	_ = v864
	var v865 int32
	_ = v865
	var v867 int32
	_ = v867
	var v871 int32
	_ = v871
	var v872 int32
	_ = v872
	var v880 int32
	_ = v880
	var v885 int32
	_ = v885
	var v888 int32
	_ = v888
	var v889 int32
	_ = v889
	var v890 int32
	_ = v890
	var v891 int32
	_ = v891
	var v898 int32
	_ = v898
	var v899 int32
	_ = v899
	var v900 int32
	_ = v900
	var v902 int32
	_ = v902
	var v907 int32
	_ = v907
	var v908 int32
	_ = v908
	var v913 int32
	_ = v913
	var v915 int32
	_ = v915
	var v920 int32
	_ = v920
	var v923 int32
	_ = v923
	var v924 int32
	_ = v924
	var v930 int32
	_ = v930
	var v933 int32
	_ = v933
	var v937 int32
	_ = v937
	var v938 int32
	_ = v938
	var v939 int32
	_ = v939
	var v941 int32
	_ = v941
	var v946 int32
	_ = v946
	var v947 int32
	_ = v947
	var v950 int32
	_ = v950
	var v951 int32
	_ = v951
	var v954 int32
	_ = v954
	var v955 int32
	_ = v955
	var v956 int32
	_ = v956
	var v962 int32
	_ = v962
	var v963 int32
	_ = v963
	var v966 int32
	_ = v966
	var v967 int32
	_ = v967
	var v968 int32
	_ = v968
	var v973 int32
	_ = v973
	var v977 int32
	_ = v977
	var v982 int32
	_ = v982
	var v983 int32
	_ = v983
	var v986 int32
	_ = v986
	var v987 int32
	_ = v987
	var v988 int32
	_ = v988
	var v994 int32
	_ = v994
	var v995 int32
	_ = v995
	var v997 int32
	_ = v997
	var v998 int32
	_ = v998
	var v999 int32
	_ = v999
	var v1004 int32
	_ = v1004
	var v1005 int32
	_ = v1005
	var v1009 int32
	_ = v1009
	var v1014 int32
	_ = v1014
	var v1015 int32
	_ = v1015
	var v1016 int32
	_ = v1016
	var v1017 int32
	_ = v1017
	var v1019 int32
	_ = v1019
	var v1020 int32
	_ = v1020
	var v1021 int32
	_ = v1021
	var v1022 int32
	_ = v1022
	var v1023 int32
	_ = v1023
	var v1025 int32
	_ = v1025
	var v1027 int32
	_ = v1027
	var v1041 int32
	_ = v1041
	var v1042 int32
	_ = v1042
	var v1047 int32
	_ = v1047
	var v1050 int32
	_ = v1050
	var v1051 int32
	_ = v1051
	var v1052 int32
	_ = v1052
	var v1053 int32
	_ = v1053
	var v1060 int32
	_ = v1060
	var v1061 int32
	_ = v1061
	var v1062 int32
	_ = v1062
	var v1064 int32
	_ = v1064
	var v1069 int32
	_ = v1069
	var v1073 int32
	_ = v1073
	var v1076 int32
	_ = v1076
	var v1080 int32
	_ = v1080
	var v1081 int32
	_ = v1081
	var v1082 int32
	_ = v1082
	var v1083 int32
	_ = v1083
	var v1085 int32
	_ = v1085
	var v1090 int32
	_ = v1090
	var v1094 int32
	_ = v1094
	var v1097 int32
	_ = v1097
	var v1098 int32
	_ = v1098
	var v1099 int32
	_ = v1099
	var v1100 int32
	_ = v1100
	var v1106 int32
	_ = v1106
	var v1107 int32
	_ = v1107
	var v1108 int32
	_ = v1108
	var v1109 int32
	_ = v1109
	var v1111 int32
	_ = v1111
	var v1116 int32
	_ = v1116
	var v1120 int32
	_ = v1120
	var v1121 int32
	_ = v1121
	var v1127 int32
	_ = v1127
	var v1132 int32
	_ = v1132
	var v1136 int32
	_ = v1136
	var v1139 int32
	_ = v1139
	var v1143 int32
	_ = v1143
	var v1144 int32
	_ = v1144
	var v1145 int32
	_ = v1145
	var v1147 int32
	_ = v1147
	var v1152 int32
	_ = v1152
	var v1156 int32
	_ = v1156
	var v1159 int32
	_ = v1159
	var v1163 int32
	_ = v1163
	var v1164 int32
	_ = v1164
	var v1165 int32
	_ = v1165
	var v1167 int32
	_ = v1167
	var v1172 int32
	_ = v1172
	var v1176 int32
	_ = v1176
	var v1179 int32
	_ = v1179
	var v1183 int32
	_ = v1183
	var v1184 int32
	_ = v1184
	var v1185 int32
	_ = v1185
	var v1187 int32
	_ = v1187
	var v1192 int32
	_ = v1192
	var v1196 int32
	_ = v1196
	var v1199 int32
	_ = v1199
	var v1203 int32
	_ = v1203
	var v1204 int32
	_ = v1204
	var v1205 int32
	_ = v1205
	var v1207 int32
	_ = v1207
	var v1212 int32
	_ = v1212
	var v1216 int32
	_ = v1216
	var v1219 int32
	_ = v1219
	var v1223 int32
	_ = v1223
	var v1224 int32
	_ = v1224
	var v1225 int32
	_ = v1225
	var v1227 int32
	_ = v1227
	var v1232 int32
	_ = v1232
	var v1236 int32
	_ = v1236
	var v1239 int32
	_ = v1239
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
	var v1256 int32
	_ = v1256
	var v1259 int32
	_ = v1259
	var v1263 int32
	_ = v1263
	var v1264 int32
	_ = v1264
	var v1265 int32
	_ = v1265
	var v1267 int32
	_ = v1267
	var v1272 int32
	_ = v1272
	var v1276 int32
	_ = v1276
	var v1279 int32
	_ = v1279
	var v1283 int32
	_ = v1283
	var v1284 int32
	_ = v1284
	var v1285 int32
	_ = v1285
	var v1287 int32
	_ = v1287
	var v1292 int32
	_ = v1292
	var v1296 int32
	_ = v1296
	var v1299 int32
	_ = v1299
	var v1303 int32
	_ = v1303
	var v1304 int32
	_ = v1304
	var v1305 int32
	_ = v1305
	var v1307 int32
	_ = v1307
	var v1312 int32
	_ = v1312
	var v1316 int32
	_ = v1316
	var v1319 int32
	_ = v1319
	var v1323 int32
	_ = v1323
	var v1324 int32
	_ = v1324
	var v1325 int32
	_ = v1325
	var v1327 int32
	_ = v1327
	var v1332 int32
	_ = v1332
	var v1342 int32
	_ = v1342
	var v1345 int32
	_ = v1345
	var v1349 int32
	_ = v1349
	var v1354 int32
	_ = v1354
	var v1358 int32
	_ = v1358
	var v1361 int32
	_ = v1361
	var v1362 int32
	_ = v1362
	var v1363 int32
	_ = v1363
	var v1364 int32
	_ = v1364
	var v1371 int32
	_ = v1371
	var v1372 int32
	_ = v1372
	var v1373 int32
	_ = v1373
	var v1375 int32
	_ = v1375
	var v1380 int32
	_ = v1380
	var v1384 int32
	_ = v1384
	var v1387 int32
	_ = v1387
	var v1388 int32
	_ = v1388
	var v1394 int32
	_ = v1394
	var v1399 int32
	_ = v1399
	var v1403 int32
	_ = v1403
	var v1404 int32
	_ = v1404
	var v1405 int32
	_ = v1405
	var v1412 int32
	_ = v1412
	var v1417 int32
	_ = v1417
	var v1421 int32
	_ = v1421
	var v1424 int32
	_ = v1424
	var v1425 int32
	_ = v1425
	var v1431 int32
	_ = v1431
	var v1436 int32
	_ = v1436
	var v1440 int32
	_ = v1440
	var v1443 int32
	_ = v1443
	var v1444 int32
	_ = v1444
	var v1445 int32
	_ = v1445
	var v1446 int32
	_ = v1446
	var v1453 int32
	_ = v1453
	var v1454 int32
	_ = v1454
	var v1455 int32
	_ = v1455
	var v1457 int32
	_ = v1457
	var v1462 int32
	_ = v1462
	var v1466 int32
	_ = v1466
	var v1469 int32
	_ = v1469
	var v1473 int32
	_ = v1473
	var v1478 int32
	_ = v1478
	var v1482 int32
	_ = v1482
	var v1485 int32
	_ = v1485
	var v1489 int32
	_ = v1489
	var v1494 int32
	_ = v1494
	var v1498 int32
	_ = v1498
	var v1501 int32
	_ = v1501
	var v1502 int32
	_ = v1502
	var v1503 int32
	_ = v1503
	var v1504 int32
	_ = v1504
	var v1511 int32
	_ = v1511
	var v1512 int32
	_ = v1512
	var v1513 int32
	_ = v1513
	var v1515 int32
	_ = v1515
	var v1520 int32
	_ = v1520
	var v1524 int32
	_ = v1524
	var v1527 int32
	_ = v1527
	var v1531 int32
	_ = v1531
	var v1536 int32
	_ = v1536
	var v1540 int32
	_ = v1540
	var v1543 int32
	_ = v1543
	var v1544 int32
	_ = v1544
	var v1545 int32
	_ = v1545
	var v1546 int32
	_ = v1546
	var v1553 int32
	_ = v1553
	var v1554 int32
	_ = v1554
	var v1555 int32
	_ = v1555
	var v1557 int32
	_ = v1557
	var v1562 int32
	_ = v1562
	var v1566 int32
	_ = v1566
	var v1569 int32
	_ = v1569
	var v1570 int32
	_ = v1570
	var v1571 int32
	_ = v1571
	var v1572 int32
	_ = v1572
	var v1579 int32
	_ = v1579
	var v1580 int32
	_ = v1580
	var v1581 int32
	_ = v1581
	var v1583 int32
	_ = v1583
	var v1588 int32
	_ = v1588
	var v1592 int32
	_ = v1592
	var v1595 int32
	_ = v1595
	var v1599 int32
	_ = v1599
	var v1600 int32
	_ = v1600
	var v1601 int32
	_ = v1601
	var v1603 int32
	_ = v1603
	var v1608 int32
	_ = v1608
	var v1612 int32
	_ = v1612
	var v1615 int32
	_ = v1615
	var v1619 int32
	_ = v1619
	var v1620 int32
	_ = v1620
	var v1621 int32
	_ = v1621
	var v1623 int32
	_ = v1623
	var v1628 int32
	_ = v1628
	var v1632 int32
	_ = v1632
	var v1635 int32
	_ = v1635
	var v1636 int32
	_ = v1636
	var v1637 int32
	_ = v1637
	var v1638 int32
	_ = v1638
	var v1645 int32
	_ = v1645
	var v1646 int32
	_ = v1646
	var v1647 int32
	_ = v1647
	var v1649 int32
	_ = v1649
	var v1654 int32
	_ = v1654
	var v1658 int32
	_ = v1658
	var v1661 int32
	_ = v1661
	var v1662 int32
	_ = v1662
	var v1663 int32
	_ = v1663
	var v1664 int32
	_ = v1664
	var v1671 int32
	_ = v1671
	var v1672 int32
	_ = v1672
	var v1673 int32
	_ = v1673
	var v1675 int32
	_ = v1675
	var v1680 int32
	_ = v1680
	var v1684 int32
	_ = v1684
	var v1687 int32
	_ = v1687
	var v1688 int32
	_ = v1688
	var v1689 int32
	_ = v1689
	var v1690 int32
	_ = v1690
	var v1697 int32
	_ = v1697
	var v1698 int32
	_ = v1698
	var v1699 int32
	_ = v1699
	var v1701 int32
	_ = v1701
	var v1706 int32
	_ = v1706
	var v1719 int32
	_ = v1719
	var v1724 int32
	_ = v1724
	var v1733 int32
	_ = v1733
	var v1739 int32
	_ = v1739
	var v1742 int32
	_ = v1742
	var v1743 int32
	_ = v1743
	var v1752 int32
	_ = v1752
	var v1754 int32
	_ = v1754
	var v1755 int32
	_ = v1755
	var v1756 int32
	_ = v1756
	var v1757 int32
	_ = v1757
	var v1758 int32
	_ = v1758
	var v1759 int32
	_ = v1759
	var v1760 int32
	_ = v1760
	var v1761 int32
	_ = v1761
	var v1781 int32
	_ = v1781
	var v1783 int32
	_ = v1783
	var v1784 int32
	_ = v1784
	var v1787 int32
	_ = v1787
	var v1789 int32
	_ = v1789
	var v1790 int32
	_ = v1790
	var v1796 int32
	_ = v1796
	var v1797 int32
	_ = v1797
	var v1800 int32
	_ = v1800
	var v1805 int32
	_ = v1805
	var v1806 int32
	_ = v1806
	var v1808 int32
	_ = v1808
	var v1809 int32
	_ = v1809
	var v1810 int32
	_ = v1810
	v3 = int32(0)
	v19 = m.G0
	v21 = v19 - int32(272)
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
		goto L40
	} else {
		goto L41
	}
L3:
	;
	v1781 = *(*int32)(unsafe.Add(mBase, uint32(l1)+60))
	if v1781 != 0 {
		goto L472
	} else {
		goto L473
	}
L4:
	;
	v1752 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1743))) = uint8(v1752)
	v1754 = *(*int32)(unsafe.Add(mBase, uint32(v1739)))
	v1755 = F_makeString(m, v1754)
	mBase = m.M
	v1756 = m.ExcPending
	if v1756 != 0 {
		goto L1
	} else {
		goto L469
	}
L5:
	;
	if v1724 == int32(0) {
		goto L3
	} else {
		goto L466
	}
L6:
	;
	v1719 = int32(0)
	v1724 = v328
	goto L5
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1684 = m.ExcPending
	if v1684 != 0 {
		goto L1
	} else {
		goto L461
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1658 = m.ExcPending
	if v1658 != 0 {
		goto L1
	} else {
		goto L456
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1632 = m.ExcPending
	if v1632 != 0 {
		goto L1
	} else {
		goto L451
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1612 = m.ExcPending
	if v1612 != 0 {
		goto L1
	} else {
		goto L446
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1592 = m.ExcPending
	if v1592 != 0 {
		goto L1
	} else {
		goto L441
	}
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1566 = m.ExcPending
	if v1566 != 0 {
		goto L1
	} else {
		goto L436
	}
L13:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1540 = m.ExcPending
	if v1540 != 0 {
		goto L1
	} else {
		goto L431
	}
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1524 = m.ExcPending
	if v1524 != 0 {
		goto L1
	} else {
		goto L427
	}
L15:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1498 = m.ExcPending
	if v1498 != 0 {
		goto L1
	} else {
		goto L422
	}
L16:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1482 = m.ExcPending
	if v1482 != 0 {
		goto L1
	} else {
		goto L418
	}
L17:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1466 = m.ExcPending
	if v1466 != 0 {
		goto L1
	} else {
		goto L414
	}
L18:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1440 = m.ExcPending
	if v1440 != 0 {
		goto L1
	} else {
		goto L409
	}
L19:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1421 = m.ExcPending
	if v1421 != 0 {
		goto L1
	} else {
		goto L405
	}
L20:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1403 = m.ExcPending
	if v1403 != 0 {
		goto L1
	} else {
		goto L402
	}
L21:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1384 = m.ExcPending
	if v1384 != 0 {
		goto L1
	} else {
		goto L398
	}
L22:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1358 = m.ExcPending
	if v1358 != 0 {
		goto L1
	} else {
		goto L393
	}
L23:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1342 = m.ExcPending
	if v1342 != 0 {
		goto L1
	} else {
		goto L389
	}
L24:
	;
	v1739 = l1 + int32(4)
	v1742 = l0 + int32(32)
	v1743 = l1 + int32(19)
	goto L4
L25:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1316 = m.ExcPending
	if v1316 != 0 {
		goto L1
	} else {
		goto L384
	}
L26:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1296 = m.ExcPending
	if v1296 != 0 {
		goto L1
	} else {
		goto L379
	}
L27:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1276 = m.ExcPending
	if v1276 != 0 {
		goto L1
	} else {
		goto L374
	}
L28:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1256 = m.ExcPending
	if v1256 != 0 {
		goto L1
	} else {
		goto L369
	}
L29:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1236 = m.ExcPending
	if v1236 != 0 {
		goto L1
	} else {
		goto L364
	}
L30:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1216 = m.ExcPending
	if v1216 != 0 {
		goto L1
	} else {
		goto L359
	}
L31:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1196 = m.ExcPending
	if v1196 != 0 {
		goto L1
	} else {
		goto L354
	}
L32:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1176 = m.ExcPending
	if v1176 != 0 {
		goto L1
	} else {
		goto L349
	}
L33:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1156 = m.ExcPending
	if v1156 != 0 {
		goto L1
	} else {
		goto L344
	}
L34:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1136 = m.ExcPending
	if v1136 != 0 {
		goto L1
	} else {
		goto L339
	}
L35:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1120 = m.ExcPending
	if v1120 != 0 {
		goto L1
	} else {
		goto L336
	}
L36:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1094 = m.ExcPending
	if v1094 != 0 {
		goto L1
	} else {
		goto L330
	}
L37:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1073 = m.ExcPending
	if v1073 != 0 {
		goto L1
	} else {
		goto L325
	}
L38:
	;
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v327)+4))
	if int32(0) < v332 {
		goto L117
	} else {
		goto L118
	}
L39:
	;
	v256 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v256)+8))
	v258 = int32(0)
	F_generateSerialExtraStmts(m, l0, l1, v257, v258, v258, v258, v21+int32(268), v21+int32(264))
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L1
	} else {
		goto L106
	}
L40:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
	if v28 == int32(0) {
		v225 = v27
		v226 = v3
		goto L43
	} else {
		goto L44
	}
L41:
	;
	goto L42
L42:
	;
	v253 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	if v253 == int32(0) {
		goto L3
	} else {
		goto L105
	}
L43:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v229 = F_typenameType(m, v227, v225, int32(0))
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L1
	} else {
		goto L97
	}
L44:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	if v31 != int32(1) {
		v225 = v27
		v226 = v3
		goto L43
	} else {
		goto L45
	}
L45:
	;
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+13)))
	if v34 != 0 {
		v225 = v27
		v226 = v3
		goto L43
	} else {
		goto L46
	}
L46:
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
		goto L49
	} else {
		goto L50
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+4)) = int32(0)
	v219 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v219)+8)) = v216
	v222 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v222)+24))
	if v223 != 0 {
		goto L37
	} else {
		goto L96
	}
L48:
	;
	if v63-v64 == int32(0) {
		v216 = v35
		goto L47
	} else {
		goto L55
	}
L49:
	;
	goto L48
L50:
	;
	v48 = v38
	v49 = v39
	goto L51
L51:
	;
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49)+1)))
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+1)))
	if v53 == int32(0) {
		v63 = v53
		v64 = v52
		goto L49
	} else {
		goto L53
	}
L52:
	;
	v63 = v53
	v64 = v52
	goto L49
L53:
	;
	v56 = int32(1)
	if v53 == v52 {
		v48 = v48 + v56
		v49 = v49 + v56
		goto L51
	} else {
		goto L54
	}
L54:
	;
	goto L52
L55:
	;
	v68 = int32(_a_F_transformColumnDefinition_1)
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38))))
	v74 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_transformColumnDefinition[1])))
	if base.B2i32(v71 == int32(0))|base.B2i32(v71 != v74) != 0 {
		v92 = v71
		v93 = v74
		goto L57
	} else {
		goto L58
	}
L56:
	;
	if v92-v93 == int32(0) {
		v216 = v35
		goto L47
	} else {
		goto L63
	}
L57:
	;
	goto L56
L58:
	;
	v77 = v38
	v78 = v68
	goto L59
L59:
	;
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78)+1)))
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77)+1)))
	if v82 == int32(0) {
		v92 = v82
		v93 = v81
		goto L57
	} else {
		goto L61
	}
L60:
	;
	v92 = v82
	v93 = v81
	goto L57
L61:
	;
	v85 = int32(1)
	if v82 == v81 {
		v77 = v77 + v85
		v78 = v78 + v85
		goto L59
	} else {
		goto L62
	}
L62:
	;
	goto L60
L63:
	;
	v97 = int32(23)
	v98 = int32(_a_F_transformColumnDefinition_2)
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38))))
	v104 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_transformColumnDefinition[2])))
	if base.B2i32(v101 == int32(0))|base.B2i32(v101 != v104) != 0 {
		v122 = v101
		v123 = v104
		goto L65
	} else {
		goto L66
	}
L64:
	;
	if v122-v123 == int32(0) {
		v216 = v97
		goto L47
	} else {
		goto L71
	}
L65:
	;
	goto L64
L66:
	;
	v107 = v38
	v108 = v98
	goto L67
L67:
	;
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108)+1)))
	v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+1)))
	if v112 == int32(0) {
		v122 = v112
		v123 = v111
		goto L65
	} else {
		goto L69
	}
L68:
	;
	v122 = v112
	v123 = v111
	goto L65
L69:
	;
	v115 = int32(1)
	if v112 == v111 {
		v107 = v107 + v115
		v108 = v108 + v115
		goto L67
	} else {
		goto L70
	}
L70:
	;
	goto L68
L71:
	;
	v127 = int32(_a_F_transformColumnDefinition_3)
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38))))
	v133 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_transformColumnDefinition[3])))
	if base.B2i32(v130 == int32(0))|base.B2i32(v130 != v133) != 0 {
		v151 = v130
		v152 = v133
		goto L73
	} else {
		goto L74
	}
L72:
	;
	if v151-v152 == int32(0) {
		v216 = v97
		goto L47
	} else {
		goto L79
	}
L73:
	;
	goto L72
L74:
	;
	v136 = v38
	v137 = v127
	goto L75
L75:
	;
	v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v137)+1)))
	v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v136)+1)))
	if v141 == int32(0) {
		v151 = v141
		v152 = v140
		goto L73
	} else {
		goto L77
	}
L76:
	;
	v151 = v141
	v152 = v140
	goto L73
L77:
	;
	v144 = int32(1)
	if v141 == v140 {
		v136 = v136 + v144
		v137 = v137 + v144
		goto L75
	} else {
		goto L78
	}
L78:
	;
	goto L76
L79:
	;
	v156 = int32(20)
	v157 = int32(_a_F_transformColumnDefinition_4)
	v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38))))
	v163 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_transformColumnDefinition[4])))
	if base.B2i32(v160 == int32(0))|base.B2i32(v160 != v163) != 0 {
		v181 = v160
		v182 = v163
		goto L81
	} else {
		goto L82
	}
L80:
	;
	if v181-v182 == int32(0) {
		v216 = v156
		goto L47
	} else {
		goto L87
	}
L81:
	;
	goto L80
L82:
	;
	v166 = v38
	v167 = v157
	goto L83
L83:
	;
	v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+1)))
	v171 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166)+1)))
	if v171 == int32(0) {
		v181 = v171
		v182 = v170
		goto L81
	} else {
		goto L85
	}
L84:
	;
	v181 = v171
	v182 = v170
	goto L81
L85:
	;
	v174 = int32(1)
	if v171 == v170 {
		v166 = v166 + v174
		v167 = v167 + v174
		goto L83
	} else {
		goto L86
	}
L86:
	;
	goto L84
L87:
	;
	v186 = int32(_a_F_transformColumnDefinition_5)
	v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38))))
	v192 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_transformColumnDefinition[5])))
	if base.B2i32(v189 == int32(0))|base.B2i32(v189 != v192) != 0 {
		v210 = v189
		v211 = v192
		goto L89
	} else {
		goto L90
	}
L88:
	;
	if v210-v211 == int32(0) {
		v216 = v156
		goto L47
	} else {
		goto L95
	}
L89:
	;
	goto L88
L90:
	;
	v195 = v38
	v196 = v186
	goto L91
L91:
	;
	v199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v196)+1)))
	v200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v195)+1)))
	if v200 == int32(0) {
		v210 = v200
		v211 = v199
		goto L89
	} else {
		goto L93
	}
L92:
	;
	v210 = v200
	v211 = v199
	goto L89
L93:
	;
	v203 = int32(1)
	if v200 == v199 {
		v195 = v195 + v203
		v196 = v196 + v203
		goto L91
	} else {
		goto L94
	}
L94:
	;
	goto L92
L95:
	;
	v225 = v27
	v226 = int32(0)
	goto L43
L96:
	;
	v225 = v222
	v226 = int32(1)
	goto L43
L97:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	if v231 != 0 {
		goto L98
	} else {
		goto L99
	}
L98:
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
		goto L101
	}
L99:
	;
	goto L100
L100:
	;
	F_ReleaseCatCache(m, v229)
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L1
	} else {
		goto L103
	}
L101:
	;
	v239 = v232 + v233
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v239)+144))
	if v240 == int32(0) {
		goto L36
	} else {
		goto L102
	}
L102:
	;
	goto L100
L103:
	;
	if v226 != 0 {
		goto L39
	} else {
		goto L104
	}
L104:
	;
	goto L42
L105:
	;
	v327 = v253
	v328 = v3
	goto L38
L106:
	;
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v21)+268))
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v21)+264))
	v269 = F_quote_qualified_identifier(m, v267, v268)
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L1
	} else {
		goto L107
	}
L107:
	;
	v272 = F_palloc0(m, int32(20))
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L1
	} else {
		goto L108
	}
L108:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v272)+16)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v272)+8)) = v269
	*(*int64)(unsafe.Add(mBase, uint32(v272))) = int64(2010044694600)
	v280 = F_palloc0(m, int32(16))
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L1
	} else {
		goto L109
	}
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v280))) = int32(73)
	v285 = F_SystemTypeName(m, int32(_a_F_transformColumnDefinition_6))
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L1
	} else {
		goto L110
	}
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v280)+12)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v280)+4)) = v272
	*(*int32)(unsafe.Add(mBase, uint32(v280)+8)) = v285
	v292 = F_SystemFuncName(m, int32(_a_F_transformColumnDefinition_7))
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L1
	} else {
		goto L111
	}
L111:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+236)) = v280
	*(*int32)(unsafe.Add(mBase, uint32(v21)+260)) = v280
	v296 = int32(1)
	v300 = F_list_make1_impl(m, v296, v21+int32(236))
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L1
	} else {
		goto L112
	}
L112:
	;
	v304 = F_makeFuncCall(m, v292, v300, int32(0), int32(-1))
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L1
	} else {
		goto L113
	}
L113:
	;
	v307 = F_palloc0(m, int32(108))
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L1
	} else {
		goto L114
	}
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v307)+104)) = int32(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v307))) = int64(8589934753)
	*(*int32)(unsafe.Add(mBase, uint32(v307)+24)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v307)+20)) = v304
	v316 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v317 = F_lappend(m, v316, v307)
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L1
	} else {
		goto L115
	}
L115:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+56)) = v317
	if v317 == int32(0) {
		goto L24
	} else {
		goto L116
	}
L116:
	;
	v327 = v317
	v328 = v296
	goto L38
L117:
	;
	v335 = int32(0)
	v342 = v335
	v343 = v3
	v346 = v3
	v347 = v3
	v348 = v335
	goto L120
L118:
	;
	v563 = v327
	goto L119
L119:
	;
	if v328 != 0 {
		v708 = int32(1)
		goto L182
	} else {
		goto L183
	}
L120:
	;
	v355 = *(*int32)(unsafe.Add(mBase, uint32(v327)+12))
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v355+v348<<(uint(int32(2))%32))))
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v359)))
	if v360 != int32(161) {
		goto L35
	} else {
		goto L122
	}
L121:
	;
	v550 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(56))))
	v563 = v550
	goto L119
L122:
	;
	v363 = int32(0)
	v366 = *(*int32)(unsafe.Add(mBase, uint32(v359)+4))
	switch v366 - int32(10) {
	case 0:
		goto L132
	case 1:
		goto L131
	case 2:
		goto L130
	case 3:
		goto L129
	case 4:
		goto L128
	case 5:
		goto L127
	default:
		v542 = v363
		v543 = v363
		v544 = v359
		v545 = v363
		goto L123
	}
L123:
	;
	v547 = v348 + int32(1)
	v548 = *(*int32)(unsafe.Add(mBase, uint32(v327)+4))
	if v547 < v548 {
		v342 = v542
		v343 = v543
		v346 = v544
		v347 = v545
		v348 = v547
		goto L120
	} else {
		goto L181
	}
L124:
	;
	v542 = v539
	v543 = v540
	v544 = v346
	v545 = v541
	goto L123
L125:
	;
	v539 = v538
	v540 = v537
	v541 = v347
	goto L124
L126:
	;
	v537 = v536
	v538 = v342
	goto L125
L127:
	;
	if v346 == int32(0) {
		goto L173
	} else {
		goto L174
	}
L128:
	;
	if v346 == int32(0) {
		goto L164
	} else {
		goto L165
	}
L129:
	;
	if v346 == int32(0) {
		goto L28
	} else {
		goto L160
	}
L130:
	;
	if v346 == int32(0) {
		goto L30
	} else {
		goto L148
	}
L131:
	;
	if v346 == int32(0) {
		goto L32
	} else {
		goto L136
	}
L132:
	;
	if v346 == int32(0) {
		goto L34
	} else {
		goto L133
	}
L133:
	;
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v346)+4))
	if base.Ui32(int32(4)) <= base.Ui32(v371-int32(6)) {
		goto L34
	} else {
		goto L134
	}
L134:
	;
	if v347&int32(1) != 0 {
		goto L33
	} else {
		goto L135
	}
L135:
	;
	v378 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v346)+12)) = uint8(v378)
	v539 = v342
	v540 = v343
	v541 = v378
	goto L124
L136:
	;
	v383 = *(*int32)(unsafe.Add(mBase, uint32(v346)+4))
	if base.Ui32(int32(4)) <= base.Ui32(v383-int32(6)) {
		goto L32
	} else {
		goto L137
	}
L137:
	;
	if v347&int32(1) != 0 {
		goto L31
	} else {
		goto L138
	}
L138:
	;
	v390 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v346)+12)) = uint8(v390)
	v392 = int32(1)
	if v342&v392 == v390 {
		goto L139
	} else {
		goto L140
	}
L139:
	;
	v539 = v363
	v540 = v343
	v541 = v392
	goto L124
L140:
	;
	goto L141
L141:
	;
	v397 = int32(1)
	v398 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v346)+13)))
	if v398 != v397 {
		v542 = v397
		v543 = v343
		v544 = v346
		v545 = v392
		goto L123
	} else {
		goto L142
	}
L142:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v404 = m.ExcPending
	if v404 != 0 {
		goto L1
	} else {
		goto L143
	}
L143:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v407 = m.ExcPending
	if v407 != 0 {
		goto L1
	} else {
		goto L144
	}
L144:
	;
	F_errmsg(m, int32(_a_F_transformColumnDefinition_8), int32(0))
	mBase = m.M
	v411 = m.ExcPending
	if v411 != 0 {
		goto L1
	} else {
		goto L145
	}
L145:
	;
	v412 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v413 = *(*int32)(unsafe.Add(mBase, uint32(v359)+104))
	F_parser_errposition(m, v412, v413)
	mBase = m.M
	v415 = m.ExcPending
	if v415 != 0 {
		goto L1
	} else {
		goto L146
	}
L146:
	;
	F_errfinish(m, int32(_a_F_transformColumnDefinition_9), int32(3947), int32(_a_F_transformColumnDefinition_10))
	mBase = m.M
	v420 = m.ExcPending
	if v420 != 0 {
		goto L1
	} else {
		goto L147
	}
L147:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L148:
	;
	v423 = *(*int32)(unsafe.Add(mBase, uint32(v346)+4))
	if base.Ui32(int32(4)) <= base.Ui32(v423-int32(6)) {
		goto L30
	} else {
		goto L149
	}
L149:
	;
	if v342&int32(1) != 0 {
		goto L29
	} else {
		goto L150
	}
L150:
	;
	v430 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v346)+13)) = uint8(v430)
	if v347&v430 == int32(0) {
		goto L151
	} else {
		goto L152
	}
L151:
	;
	v436 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v346)+12)) = uint8(v436)
	v539 = v436
	v540 = v343
	v541 = v363
	goto L124
L152:
	;
	goto L153
L153:
	;
	v439 = int32(1)
	v441 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v346)+12)))
	if v441 != 0 {
		v542 = v439
		v543 = v343
		v544 = v346
		v545 = v439
		goto L123
	} else {
		goto L154
	}
L154:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v445 = m.ExcPending
	if v445 != 0 {
		goto L1
	} else {
		goto L155
	}
L155:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v448 = m.ExcPending
	if v448 != 0 {
		goto L1
	} else {
		goto L156
	}
L156:
	;
	F_errmsg(m, int32(_a_F_transformColumnDefinition_8), int32(0))
	mBase = m.M
	v452 = m.ExcPending
	if v452 != 0 {
		goto L1
	} else {
		goto L157
	}
L157:
	;
	v453 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v454 = *(*int32)(unsafe.Add(mBase, uint32(v359)+104))
	F_parser_errposition(m, v453, v454)
	mBase = m.M
	v456 = m.ExcPending
	if v456 != 0 {
		goto L1
	} else {
		goto L158
	}
L158:
	;
	F_errfinish(m, int32(_a_F_transformColumnDefinition_9), int32(3973), int32(_a_F_transformColumnDefinition_10))
	mBase = m.M
	v461 = m.ExcPending
	if v461 != 0 {
		goto L1
	} else {
		goto L159
	}
L159:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L160:
	;
	v464 = *(*int32)(unsafe.Add(mBase, uint32(v346)+4))
	if base.Ui32(int32(4)) <= base.Ui32(v464-int32(6)) {
		goto L28
	} else {
		goto L161
	}
L161:
	;
	if v342&int32(1) != 0 {
		goto L27
	} else {
		goto L162
	}
L162:
	;
	v471 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v346)+13)) = uint8(v471)
	v537 = v343
	v538 = int32(1)
	goto L125
L163:
	;
	if v343&int32(1) != 0 {
		goto L26
	} else {
		goto L171
	}
L164:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v482 = m.ExcPending
	if v482 != 0 {
		goto L1
	} else {
		goto L166
	}
L165:
	;
	v476 = *(*int32)(unsafe.Add(mBase, uint32(v346)+4))
	switch v476 - int32(5) {
	case 0, 4:
		goto L163
	default:
		goto L164
	}
L166:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v485 = m.ExcPending
	if v485 != 0 {
		goto L1
	} else {
		goto L167
	}
L167:
	;
	F_errmsg(m, int32(_a_F_transformColumnDefinition_11), int32(0))
	mBase = m.M
	v489 = m.ExcPending
	if v489 != 0 {
		goto L1
	} else {
		goto L168
	}
L168:
	;
	v490 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v491 = *(*int32)(unsafe.Add(mBase, uint32(v359)+104))
	F_parser_errposition(m, v490, v491)
	mBase = m.M
	v493 = m.ExcPending
	if v493 != 0 {
		goto L1
	} else {
		goto L169
	}
L169:
	;
	F_errfinish(m, int32(_a_F_transformColumnDefinition_9), int32(3998), int32(_a_F_transformColumnDefinition_10))
	mBase = m.M
	v498 = m.ExcPending
	if v498 != 0 {
		goto L1
	} else {
		goto L170
	}
L170:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L171:
	;
	v501 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v346)+14)) = uint8(v501)
	v536 = v501
	goto L126
L172:
	;
	if v343&int32(1) != 0 {
		goto L25
	} else {
		goto L180
	}
L173:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v512 = m.ExcPending
	if v512 != 0 {
		goto L1
	} else {
		goto L175
	}
L174:
	;
	v506 = *(*int32)(unsafe.Add(mBase, uint32(v346)+4))
	switch v506 - int32(5) {
	case 0, 4:
		goto L172
	default:
		goto L173
	}
L175:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v515 = m.ExcPending
	if v515 != 0 {
		goto L1
	} else {
		goto L176
	}
L176:
	;
	F_errmsg(m, int32(_a_F_transformColumnDefinition_12), int32(0))
	mBase = m.M
	v519 = m.ExcPending
	if v519 != 0 {
		goto L1
	} else {
		goto L177
	}
L177:
	;
	v520 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v521 = *(*int32)(unsafe.Add(mBase, uint32(v359)+104))
	F_parser_errposition(m, v520, v521)
	mBase = m.M
	v523 = m.ExcPending
	if v523 != 0 {
		goto L1
	} else {
		goto L178
	}
L178:
	;
	F_errfinish(m, int32(_a_F_transformColumnDefinition_9), int32(4015), int32(_a_F_transformColumnDefinition_10))
	mBase = m.M
	v528 = m.ExcPending
	if v528 != 0 {
		goto L1
	} else {
		goto L179
	}
L179:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L180:
	;
	v531 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v346)+16)) = uint8(v531)
	v533 = int32(256)
	*(*uint16)(unsafe.Add(mBase, uint32(v346)+14)) = uint16(v533)
	v536 = int32(1)
	goto L126
L181:
	;
	goto L121
L182:
	;
	v711 = l0 + int32(32)
	v713 = l1 + int32(19)
	v715 = l1 + int32(4)
	if v563 == int32(0) {
		goto L6
	} else {
		goto L209
	}
L183:
	;
	if v563 == int32(0) {
		goto L3
	} else {
		goto L184
	}
L184:
	;
	v571 = int32(0)
	v572 = *(*int32)(unsafe.Add(mBase, uint32(v563)+4))
	if v572 <= v571 {
		v708 = v571
		goto L182
	} else {
		goto L185
	}
L185:
	;
	v575 = int32(0)
	if v575 < v572 {
		goto L186
	} else {
		goto L187
	}
L186:
	;
	v578 = v572
	goto L188
L187:
	;
	v578 = v575
	goto L188
L188:
	;
	v580 = v578 & int32(3)
	v581 = int32(0)
	if int32(4) <= v572 {
		goto L189
	} else {
		goto L190
	}
L189:
	;
	v587 = *(*int32)(unsafe.Add(mBase, uint32(v563)+12))
	v594 = int32(0)
	v596 = v581
	v605 = v571
	goto L192
L190:
	;
	v648 = v581
	v657 = v571
	goto L191
L191:
	;
	v659 = *(*int32)(unsafe.Add(mBase, uint32(v563)+12))
	v667 = v648
	v673 = v581
	v676 = v657
	goto L204
L192:
	;
	v609 = v587 + v596<<(uint(int32(2))%32)
	v610 = *(*int32)(unsafe.Add(mBase, uint32(v609)))
	v611 = *(*int32)(unsafe.Add(mBase, uint32(v610)+4))
	switch v611 - int32(3) {
	case 0, 3:
		goto L195
	default:
		v615 = v605
		goto L194
	}
L193:
	;
	if v580 == int32(0) {
		v708 = v633
		goto L182
	} else {
		goto L203
	}
L194:
	;
	v616 = *(*int32)(unsafe.Add(mBase, uint32(v609)+4))
	v617 = *(*int32)(unsafe.Add(mBase, uint32(v616)+4))
	switch v617 - int32(3) {
	case 0, 3:
		goto L197
	default:
		v621 = v615
		goto L196
	}
L195:
	;
	v615 = int32(1)
	goto L194
L196:
	;
	v622 = *(*int32)(unsafe.Add(mBase, uint32(v609)+8))
	v623 = *(*int32)(unsafe.Add(mBase, uint32(v622)+4))
	switch v623 - int32(3) {
	case 0, 3:
		goto L199
	default:
		v627 = v621
		goto L198
	}
L197:
	;
	v621 = int32(1)
	goto L196
L198:
	;
	v628 = *(*int32)(unsafe.Add(mBase, uint32(v609)+12))
	v629 = *(*int32)(unsafe.Add(mBase, uint32(v628)+4))
	switch v629 - int32(3) {
	case 0, 3:
		goto L201
	default:
		v633 = v627
		goto L200
	}
L199:
	;
	v627 = int32(1)
	goto L198
L200:
	;
	v634 = int32(4)
	v635 = v596 + v634
	v637 = v594 + v634
	if v637 != v578&int32(2147483644) {
		v594 = v637
		v596 = v635
		v605 = v633
		goto L192
	} else {
		goto L202
	}
L201:
	;
	v633 = int32(1)
	goto L200
L202:
	;
	goto L193
L203:
	;
	v648 = v635
	v657 = v633
	goto L191
L204:
	;
	v681 = *(*int32)(unsafe.Add(mBase, uint32(v659+v667<<(uint(int32(2))%32))))
	v682 = *(*int32)(unsafe.Add(mBase, uint32(v681)+4))
	switch v682 - int32(3) {
	case 0, 3:
		goto L207
	default:
		v686 = v676
		goto L206
	}
L205:
	;
	v708 = v686
	goto L182
L206:
	;
	v687 = int32(1)
	v690 = v673 + v687
	if v690 != v580 {
		v667 = v667 + v687
		v673 = v690
		v676 = v686
		goto L204
	} else {
		goto L208
	}
L207:
	;
	v686 = int32(1)
	goto L206
L208:
	;
	goto L205
L209:
	;
	v718 = int32(0)
	v719 = *(*int32)(unsafe.Add(mBase, uint32(v563)+4))
	if v719 <= v718 {
		goto L6
	} else {
		goto L210
	}
L210:
	;
	v722 = int32(0)
	v731 = v718
	v733 = v722
	v734 = v722
	v737 = v722
	v738 = v722
	v740 = v722
	v742 = v328
	goto L211
L211:
	;
	v745 = *(*int32)(unsafe.Add(mBase, uint32(v563)+12))
	v749 = *(*int32)(unsafe.Add(mBase, uint32(v745+v734<<(uint(int32(2))%32))))
	v750 = *(*int32)(unsafe.Add(mBase, uint32(v749)+4))
	switch v750 {
	case 0:
		goto L225
	case 1:
		goto L224
	case 2:
		goto L223
	case 3:
		goto L222
	case 4:
		goto L221
	case 5:
		goto L214
	case 6:
		goto L220
	case 7:
		goto L219
	case 8:
		goto L217
	case 9:
		goto L216
	case 10, 11, 12, 13, 14, 15:
		v1019 = v731
		v1020 = v733
		v1021 = v737
		v1022 = v738
		v1023 = v740
		v1025 = v742
		goto L213
	default:
		goto L215
	}
L212:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1047 = m.ExcPending
	if v1047 != 0 {
		goto L1
	} else {
		goto L320
	}
L213:
	;
	v1027 = int32(1)
	if v1023&v1027&(v1020&v1027) != 0 {
		goto L9
	} else {
		goto L314
	}
L214:
	;
	v1015 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1016 = F_lappend(m, v1015, v749)
	mBase = m.M
	v1017 = m.ExcPending
	if v1017 != 0 {
		goto L1
	} else {
		goto L313
	}
L215:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1004 = m.ExcPending
	if v1004 != 0 {
		goto L1
	} else {
		goto L310
	}
L216:
	;
	v983 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
	if v983 == int32(1) {
		goto L10
	} else {
		goto L306
	}
L217:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v973 = m.ExcPending
	if v973 != 0 {
		goto L1
	} else {
		goto L303
	}
L218:
	;
	v951 = *(*int32)(unsafe.Add(mBase, uint32(v749)+32))
	if v951 == int32(0) {
		goto L297
	} else {
		goto L298
	}
L219:
	;
	v947 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
	if v947 == int32(1) {
		goto L11
	} else {
		goto L296
	}
L220:
	;
	if v737&int32(1) != 0 {
		goto L286
	} else {
		goto L287
	}
L221:
	;
	v908 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+68)))
	if v908 == int32(1) {
		goto L14
	} else {
		goto L284
	}
L222:
	;
	v847 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+68)))
	if v847 == int32(1) {
		goto L17
	} else {
		goto L267
	}
L223:
	;
	if v740&int32(1) != 0 {
		goto L18
	} else {
		goto L266
	}
L224:
	;
	v762 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+60)))
	if v762 == int32(1) {
		goto L232
	} else {
		goto L233
	}
L225:
	;
	if v737&int32(1) != 0 {
		goto L227
	} else {
		goto L228
	}
L226:
	;
	v759 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v713))) = uint8(v759)
	v1019 = v731
	v1020 = v733
	v1021 = int32(1)
	v1022 = v738
	v1023 = v740
	v1025 = v742
	goto L213
L227:
	;
	v753 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v713))))
	if (v753|v742)&int32(1) == int32(0) {
		goto L226
	} else {
		goto L230
	}
L228:
	;
	goto L229
L229:
	;
	if v742 != 0 {
		goto L7
	} else {
		goto L231
	}
L230:
	;
	goto L7
L231:
	;
	goto L226
L232:
	;
	v765 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v749)+17)))
	if v765 == int32(1) {
		goto L23
	} else {
		goto L235
	}
L233:
	;
	goto L234
L234:
	;
	if v737&int32(1) != 0 {
		goto L236
	} else {
		goto L237
	}
L235:
	;
	goto L234
L236:
	;
	v770 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v713))))
	if v770 == int32(0) {
		goto L22
	} else {
		goto L239
	}
L237:
	;
	goto L238
L238:
	;
	if v708 != 0 {
		goto L240
	} else {
		goto L241
	}
L239:
	;
	goto L238
L240:
	;
	v773 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v749)+17)))
	if v773 == int32(1) {
		goto L21
	} else {
		goto L243
	}
L241:
	;
	goto L242
L242:
	;
	v776 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v713))))
	if v776 == int32(0) {
		goto L244
	} else {
		goto L245
	}
L243:
	;
	goto L242
L244:
	;
	v779 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v713))) = uint8(v779)
	v782 = *(*int32)(unsafe.Add(mBase, uint32(v715)))
	v783 = F_makeString(m, v782)
	mBase = m.M
	v784 = m.ExcPending
	if v784 != 0 {
		goto L1
	} else {
		goto L247
	}
L245:
	;
	goto L246
L246:
	;
	if v731 == int32(0) {
		goto L250
	} else {
		goto L251
	}
L247:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+124)) = v783
	*(*int32)(unsafe.Add(mBase, uint32(v21)+256)) = v783
	v790 = F_list_make1_impl(m, int32(1), v21+int32(124))
	mBase = m.M
	v791 = m.ExcPending
	if v791 != 0 {
		goto L1
	} else {
		goto L248
	}
L248:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v749)+32)) = v790
	v793 = *(*int32)(unsafe.Add(mBase, uint32(v711)))
	v794 = F_lappend(m, v793, v749)
	mBase = m.M
	v795 = m.ExcPending
	if v795 != 0 {
		goto L1
	} else {
		goto L249
	}
L249:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v711))) = v794
	v1019 = v749
	v1020 = v733
	v1021 = v779
	v1022 = v738
	v1023 = v740
	v1025 = int32(0)
	goto L213
L250:
	;
	v1019 = int32(0)
	v1020 = v733
	v1021 = v737
	v1022 = v738
	v1023 = v740
	v1025 = v742
	goto L213
L251:
	;
	goto L252
L252:
	;
	v801 = *(*int32)(unsafe.Add(mBase, uint32(v749)+8))
	if v801 == int32(0) {
		goto L253
	} else {
		goto L254
	}
L253:
	;
	v834 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v731)+17)))
	v835 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v749)+17)))
	if v834 != v835 {
		goto L19
	} else {
		goto L264
	}
L254:
	;
	v804 = *(*int32)(unsafe.Add(mBase, uint32(v731)+8))
	if v804 == int32(0) {
		goto L253
	} else {
		goto L255
	}
L255:
	;
	v809 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v804))))
	v812 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v801))))
	if base.B2i32(v809 == int32(0))|base.B2i32(v809 != v812) != 0 {
		v830 = v809
		v831 = v812
		goto L257
	} else {
		goto L258
	}
L256:
	;
	if v830-v831 != 0 {
		goto L20
	} else {
		goto L263
	}
L257:
	;
	goto L256
L258:
	;
	v815 = v804
	v816 = v801
	goto L259
L259:
	;
	v819 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v816)+1)))
	v820 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v815)+1)))
	if v820 == int32(0) {
		v830 = v820
		v831 = v819
		goto L257
	} else {
		goto L261
	}
L260:
	;
	v830 = v820
	v831 = v819
	goto L257
L261:
	;
	v823 = int32(1)
	if v820 == v819 {
		v815 = v815 + v823
		v816 = v816 + v823
		goto L259
	} else {
		goto L262
	}
L262:
	;
	goto L260
L263:
	;
	goto L253
L264:
	;
	v837 = *(*int32)(unsafe.Add(mBase, uint32(v731)+8))
	if v837|base.B2i32(v801 == int32(0)) != 0 {
		v1019 = v731
		v1020 = v733
		v1021 = v737
		v1022 = v738
		v1023 = v740
		v1025 = v742
		goto L213
	} else {
		goto L265
	}
L265:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v731)+8)) = v801
	v1019 = v731
	v1020 = v733
	v1021 = v737
	v1022 = v738
	v1023 = v740
	v1025 = v742
	goto L213
L266:
	;
	v844 = *(*int32)(unsafe.Add(mBase, uint32(v749)+20))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v844
	v1019 = v731
	v1020 = v733
	v1021 = v737
	v1022 = v738
	v1023 = int32(1)
	v1025 = v742
	goto L213
L267:
	;
	v850 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	if v850 != 0 {
		goto L16
	} else {
		goto L268
	}
L268:
	;
	v851 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v852 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v854 = F_typenameType(m, v851, v852, int32(0))
	mBase = m.M
	v855 = m.ExcPending
	if v855 != 0 {
		goto L1
	} else {
		goto L269
	}
L269:
	;
	v856 = *(*int32)(unsafe.Add(mBase, uint32(v854)+16))
	v857 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v856)+22)))
	v859 = *(*int32)(unsafe.Add(mBase, uint32(v856+v857)))
	F_ReleaseCatCache(m, v854)
	mBase = m.M
	v861 = m.ExcPending
	if v861 != 0 {
		goto L1
	} else {
		goto L270
	}
L270:
	;
	if v733&int32(1) != 0 {
		goto L15
	} else {
		goto L271
	}
L271:
	;
	v864 = int32(1)
	v865 = *(*int32)(unsafe.Add(mBase, uint32(v749)+48))
	v867 = int32(0)
	F_generateSerialExtraStmts(m, l0, l1, v859, v865, v864, v867, v867, v867)
	mBase = m.M
	v871 = m.ExcPending
	if v871 != 0 {
		goto L1
	} else {
		goto L272
	}
L272:
	;
	v872 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v749)+28)))
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+36)) = uint8(v872)
	if v737&int32(1) == int32(0) {
		goto L273
	} else {
		goto L274
	}
L273:
	;
	v1019 = v731
	v1020 = v864
	v1021 = int32(0)
	v1022 = v738
	v1023 = v740
	v1025 = int32(1)
	goto L213
L274:
	;
	goto L275
L275:
	;
	v880 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v713))))
	if v880 != 0 {
		goto L276
	} else {
		goto L277
	}
L276:
	;
	v1019 = v731
	v1020 = v864
	v1021 = int32(1)
	v1022 = v738
	v1023 = v740
	v1025 = v742
	goto L213
L277:
	;
	goto L278
L278:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v885 = m.ExcPending
	if v885 != 0 {
		goto L1
	} else {
		goto L279
	}
L279:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v888 = m.ExcPending
	if v888 != 0 {
		goto L1
	} else {
		goto L280
	}
L280:
	;
	v889 = *(*int32)(unsafe.Add(mBase, uint32(v715)))
	v890 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v891 = *(*int32)(unsafe.Add(mBase, uint32(v890)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+180)) = v891
	*(*int32)(unsafe.Add(mBase, uint32(v21)+176)) = v889
	F_errmsg(m, int32(_a_F_transformColumnDefinition_13), v21+int32(176))
	mBase = m.M
	v898 = m.ExcPending
	if v898 != 0 {
		goto L1
	} else {
		goto L281
	}
L281:
	;
	v899 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v900 = *(*int32)(unsafe.Add(mBase, uint32(v749)+104))
	F_parser_errposition(m, v899, v900)
	mBase = m.M
	v902 = m.ExcPending
	if v902 != 0 {
		goto L1
	} else {
		goto L282
	}
L282:
	;
	F_errfinish(m, int32(_a_F_transformColumnDefinition_9), int32(876), int32(_a_F_transformColumnDefinition_14))
	mBase = m.M
	v907 = m.ExcPending
	if v907 != 0 {
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
	if v738&int32(1) != 0 {
		goto L13
	} else {
		goto L285
	}
L285:
	;
	v913 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v749)+29)))
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+44)) = uint8(v913)
	v915 = *(*int32)(unsafe.Add(mBase, uint32(v749)+20))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v915
	v1019 = v731
	v1020 = v733
	v1021 = v737
	v1022 = int32(1)
	v1023 = v740
	v1025 = v742
	goto L213
L286:
	;
	v920 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v713))))
	if v920 == int32(0) {
		goto L12
	} else {
		goto L289
	}
L287:
	;
	goto L288
L288:
	;
	v923 = int32(1)
	v924 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
	if v924 != v923 {
		v950 = v923
		goto L218
	} else {
		goto L290
	}
L289:
	;
	goto L288
L290:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v930 = m.ExcPending
	if v930 != 0 {
		goto L1
	} else {
		goto L291
	}
L291:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v933 = m.ExcPending
	if v933 != 0 {
		goto L1
	} else {
		goto L292
	}
L292:
	;
	F_errmsg(m, int32(_a_F_transformColumnDefinition_15), int32(0))
	mBase = m.M
	v937 = m.ExcPending
	if v937 != 0 {
		goto L1
	} else {
		goto L293
	}
L293:
	;
	v938 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v939 = *(*int32)(unsafe.Add(mBase, uint32(v749)+104))
	F_parser_errposition(m, v938, v939)
	mBase = m.M
	v941 = m.ExcPending
	if v941 != 0 {
		goto L1
	} else {
		goto L294
	}
L294:
	;
	F_errfinish(m, int32(_a_F_transformColumnDefinition_9), int32(917), int32(_a_F_transformColumnDefinition_14))
	mBase = m.M
	v946 = m.ExcPending
	if v946 != 0 {
		goto L1
	} else {
		goto L295
	}
L295:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L296:
	;
	v950 = v742
	goto L218
L297:
	;
	v954 = *(*int32)(unsafe.Add(mBase, uint32(v715)))
	v955 = F_makeString(m, v954)
	mBase = m.M
	v956 = m.ExcPending
	if v956 != 0 {
		goto L1
	} else {
		goto L300
	}
L298:
	;
	goto L299
L299:
	;
	v966 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v967 = F_lappend(m, v966, v749)
	mBase = m.M
	v968 = m.ExcPending
	if v968 != 0 {
		goto L1
	} else {
		goto L302
	}
L300:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+204)) = v955
	*(*int32)(unsafe.Add(mBase, uint32(v21)+252)) = v955
	v962 = F_list_make1_impl(m, int32(1), v21+int32(204))
	mBase = m.M
	v963 = m.ExcPending
	if v963 != 0 {
		goto L1
	} else {
		goto L301
	}
L301:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v749)+32)) = v962
	goto L299
L302:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v967
	v1019 = v731
	v1020 = v733
	v1021 = v737
	v1022 = v738
	v1023 = v740
	v1025 = v950
	goto L213
L303:
	;
	F_errmsg_internal(m, int32(_a_F_transformColumnDefinition_16), int32(0))
	mBase = m.M
	v977 = m.ExcPending
	if v977 != 0 {
		goto L1
	} else {
		goto L304
	}
L304:
	;
	F_errfinish(m, int32(_a_F_transformColumnDefinition_9), int32(934), int32(_a_F_transformColumnDefinition_14))
	mBase = m.M
	v982 = m.ExcPending
	if v982 != 0 {
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
	v986 = *(*int32)(unsafe.Add(mBase, uint32(v715)))
	v987 = F_makeString(m, v986)
	mBase = m.M
	v988 = m.ExcPending
	if v988 != 0 {
		goto L1
	} else {
		goto L307
	}
L307:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+220)) = v987
	*(*int32)(unsafe.Add(mBase, uint32(v21)+248)) = v987
	v994 = F_list_make1_impl(m, int32(1), v21+int32(220))
	mBase = m.M
	v995 = m.ExcPending
	if v995 != 0 {
		goto L1
	} else {
		goto L308
	}
L308:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v749)+76)) = v994
	v997 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v998 = F_lappend(m, v997, v749)
	mBase = m.M
	v999 = m.ExcPending
	if v999 != 0 {
		goto L1
	} else {
		goto L309
	}
L309:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v998
	v1019 = v731
	v1020 = v733
	v1021 = v737
	v1022 = v738
	v1023 = v740
	v1025 = v742
	goto L213
L310:
	;
	v1005 = *(*int32)(unsafe.Add(mBase, uint32(v749)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = v1005
	F_errmsg_internal(m, int32(_a_F_transformColumnDefinition_17), v21)
	mBase = m.M
	v1009 = m.ExcPending
	if v1009 != 0 {
		goto L1
	} else {
		goto L311
	}
L311:
	;
	F_errfinish(m, int32(_a_F_transformColumnDefinition_9), int32(964), int32(_a_F_transformColumnDefinition_14))
	mBase = m.M
	v1014 = m.ExcPending
	if v1014 != 0 {
		goto L1
	} else {
		goto L312
	}
L312:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L313:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v1016
	v1019 = v731
	v1020 = v733
	v1021 = v737
	v1022 = v738
	v1023 = v740
	v1025 = v742
	goto L213
L314:
	;
	if v1022&v1023&int32(1) != 0 {
		goto L8
	} else {
		goto L315
	}
L315:
	;
	if v1020&v1022&int32(1) == int32(0) {
		goto L316
	} else {
		goto L317
	}
L316:
	;
	v1041 = v734 + int32(1)
	v1042 = *(*int32)(unsafe.Add(mBase, uint32(v563)+4))
	if v1042 <= v1041 {
		v1719 = v1021
		v1724 = v1025
		goto L5
	} else {
		goto L319
	}
L317:
	;
	goto L318
L318:
	;
	goto L212
L319:
	;
	v731 = v1019
	v733 = v1020
	v734 = v1041
	v737 = v1021
	v738 = v1022
	v740 = v1023
	v742 = v1025
	goto L211
L320:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v1050 = m.ExcPending
	if v1050 != 0 {
		goto L1
	} else {
		goto L321
	}
L321:
	;
	v1051 = *(*int32)(unsafe.Add(mBase, uint32(v715)))
	v1052 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1053 = *(*int32)(unsafe.Add(mBase, uint32(v1052)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+68)) = v1053
	*(*int32)(unsafe.Add(mBase, uint32(v21)+64)) = v1051
	F_errmsg(m, int32(_a_F_transformColumnDefinition_18), v21-int32(-64))
	mBase = m.M
	v1060 = m.ExcPending
	if v1060 != 0 {
		goto L1
	} else {
		goto L322
	}
L322:
	;
	v1061 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1062 = *(*int32)(unsafe.Add(mBase, uint32(v749)+104))
	F_parser_errposition(m, v1061, v1062)
	mBase = m.M
	v1064 = m.ExcPending
	if v1064 != 0 {
		goto L1
	} else {
		goto L323
	}
L323:
	;
	F_errfinish(m, int32(_a_F_transformColumnDefinition_9), int32(990), int32(_a_F_transformColumnDefinition_14))
	mBase = m.M
	v1069 = m.ExcPending
	if v1069 != 0 {
		goto L1
	} else {
		goto L324
	}
L324:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L325:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v1076 = m.ExcPending
	if v1076 != 0 {
		goto L1
	} else {
		goto L326
	}
L326:
	;
	F_errmsg(m, int32(_a_F_transformColumnDefinition_19), int32(0))
	mBase = m.M
	v1080 = m.ExcPending
	if v1080 != 0 {
		goto L1
	} else {
		goto L327
	}
L327:
	;
	v1081 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1082 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v1083 = *(*int32)(unsafe.Add(mBase, uint32(v1082)+28))
	F_parser_errposition(m, v1081, v1083)
	mBase = m.M
	v1085 = m.ExcPending
	if v1085 != 0 {
		goto L1
	} else {
		goto L328
	}
L328:
	;
	F_errfinish(m, int32(_a_F_transformColumnDefinition_9), int32(644), int32(_a_F_transformColumnDefinition_14))
	mBase = m.M
	v1090 = m.ExcPending
	if v1090 != 0 {
		goto L1
	} else {
		goto L329
	}
L329:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L330:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v1097 = m.ExcPending
	if v1097 != 0 {
		goto L1
	} else {
		goto L331
	}
L331:
	;
	v1098 = *(*int32)(unsafe.Add(mBase, uint32(v239)))
	v1099 = F_format_type_be(m, v1098)
	mBase = m.M
	v1100 = m.ExcPending
	if v1100 != 0 {
		goto L1
	} else {
		goto L332
	}
L332:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+240)) = v1099
	F_errmsg(m, int32(_a_F_transformColumnDefinition_20), v21+int32(240))
	mBase = m.M
	v1106 = m.ExcPending
	if v1106 != 0 {
		goto L1
	} else {
		goto L333
	}
L333:
	;
	v1107 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1108 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v1109 = *(*int32)(unsafe.Add(mBase, uint32(v1108)+12))
	F_parser_errposition(m, v1107, v1109)
	mBase = m.M
	v1111 = m.ExcPending
	if v1111 != 0 {
		goto L1
	} else {
		goto L334
	}
L334:
	;
	F_errfinish(m, int32(_a_F_transformColumnDefinition_9), int32(4067), int32(_a_F_transformColumnDefinition_21))
	mBase = m.M
	v1116 = m.ExcPending
	if v1116 != 0 {
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
	v1121 = *(*int32)(unsafe.Add(mBase, uint32(v359)))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+224)) = v1121
	F_errmsg_internal(m, int32(_a_F_transformColumnDefinition_22), v21+int32(224))
	mBase = m.M
	v1127 = m.ExcPending
	if v1127 != 0 {
		goto L1
	} else {
		goto L337
	}
L337:
	;
	F_errfinish(m, int32(_a_F_transformColumnDefinition_9), int32(3911), int32(_a_F_transformColumnDefinition_10))
	mBase = m.M
	v1132 = m.ExcPending
	if v1132 != 0 {
		goto L1
	} else {
		goto L338
	}
L338:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L339:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v1139 = m.ExcPending
	if v1139 != 0 {
		goto L1
	} else {
		goto L340
	}
L340:
	;
	F_errmsg(m, int32(_a_F_transformColumnDefinition_23), int32(0))
	mBase = m.M
	v1143 = m.ExcPending
	if v1143 != 0 {
		goto L1
	} else {
		goto L341
	}
L341:
	;
	v1144 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1145 = *(*int32)(unsafe.Add(mBase, uint32(v359)+104))
	F_parser_errposition(m, v1144, v1145)
	mBase = m.M
	v1147 = m.ExcPending
	if v1147 != 0 {
		goto L1
	} else {
		goto L342
	}
L342:
	;
	F_errfinish(m, int32(_a_F_transformColumnDefinition_9), int32(3919), int32(_a_F_transformColumnDefinition_10))
	mBase = m.M
	v1152 = m.ExcPending
	if v1152 != 0 {
		goto L1
	} else {
		goto L343
	}
L343:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L344:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v1159 = m.ExcPending
	if v1159 != 0 {
		goto L1
	} else {
		goto L345
	}
L345:
	;
	F_errmsg(m, int32(_a_F_transformColumnDefinition_24), int32(0))
	mBase = m.M
	v1163 = m.ExcPending
	if v1163 != 0 {
		goto L1
	} else {
		goto L346
	}
L346:
	;
	v1164 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1165 = *(*int32)(unsafe.Add(mBase, uint32(v359)+104))
	F_parser_errposition(m, v1164, v1165)
	mBase = m.M
	v1167 = m.ExcPending
	if v1167 != 0 {
		goto L1
	} else {
		goto L347
	}
L347:
	;
	F_errfinish(m, int32(_a_F_transformColumnDefinition_9), int32(3924), int32(_a_F_transformColumnDefinition_10))
	mBase = m.M
	v1172 = m.ExcPending
	if v1172 != 0 {
		goto L1
	} else {
		goto L348
	}
L348:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L349:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v1179 = m.ExcPending
	if v1179 != 0 {
		goto L1
	} else {
		goto L350
	}
L350:
	;
	F_errmsg(m, int32(_a_F_transformColumnDefinition_25), int32(0))
	mBase = m.M
	v1183 = m.ExcPending
	if v1183 != 0 {
		goto L1
	} else {
		goto L351
	}
L351:
	;
	v1184 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1185 = *(*int32)(unsafe.Add(mBase, uint32(v359)+104))
	F_parser_errposition(m, v1184, v1185)
	mBase = m.M
	v1187 = m.ExcPending
	if v1187 != 0 {
		goto L1
	} else {
		goto L352
	}
L352:
	;
	F_errfinish(m, int32(_a_F_transformColumnDefinition_9), int32(3934), int32(_a_F_transformColumnDefinition_10))
	mBase = m.M
	v1192 = m.ExcPending
	if v1192 != 0 {
		goto L1
	} else {
		goto L353
	}
L353:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L354:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v1199 = m.ExcPending
	if v1199 != 0 {
		goto L1
	} else {
		goto L355
	}
L355:
	;
	F_errmsg(m, int32(_a_F_transformColumnDefinition_24), int32(0))
	mBase = m.M
	v1203 = m.ExcPending
	if v1203 != 0 {
		goto L1
	} else {
		goto L356
	}
L356:
	;
	v1204 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1205 = *(*int32)(unsafe.Add(mBase, uint32(v359)+104))
	F_parser_errposition(m, v1204, v1205)
	mBase = m.M
	v1207 = m.ExcPending
	if v1207 != 0 {
		goto L1
	} else {
		goto L357
	}
L357:
	;
	F_errfinish(m, int32(_a_F_transformColumnDefinition_9), int32(3939), int32(_a_F_transformColumnDefinition_10))
	mBase = m.M
	v1212 = m.ExcPending
	if v1212 != 0 {
		goto L1
	} else {
		goto L358
	}
L358:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L359:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v1219 = m.ExcPending
	if v1219 != 0 {
		goto L1
	} else {
		goto L360
	}
L360:
	;
	F_errmsg(m, int32(_a_F_transformColumnDefinition_26), int32(0))
	mBase = m.M
	v1223 = m.ExcPending
	if v1223 != 0 {
		goto L1
	} else {
		goto L361
	}
L361:
	;
	v1224 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1225 = *(*int32)(unsafe.Add(mBase, uint32(v359)+104))
	F_parser_errposition(m, v1224, v1225)
	mBase = m.M
	v1227 = m.ExcPending
	if v1227 != 0 {
		goto L1
	} else {
		goto L362
	}
L362:
	;
	F_errfinish(m, int32(_a_F_transformColumnDefinition_9), int32(3955), int32(_a_F_transformColumnDefinition_10))
	mBase = m.M
	v1232 = m.ExcPending
	if v1232 != 0 {
		goto L1
	} else {
		goto L363
	}
L363:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L364:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v1239 = m.ExcPending
	if v1239 != 0 {
		goto L1
	} else {
		goto L365
	}
L365:
	;
	F_errmsg(m, int32(_a_F_transformColumnDefinition_27), int32(0))
	mBase = m.M
	v1243 = m.ExcPending
	if v1243 != 0 {
		goto L1
	} else {
		goto L366
	}
L366:
	;
	v1244 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1245 = *(*int32)(unsafe.Add(mBase, uint32(v359)+104))
	F_parser_errposition(m, v1244, v1245)
	mBase = m.M
	v1247 = m.ExcPending
	if v1247 != 0 {
		goto L1
	} else {
		goto L367
	}
L367:
	;
	F_errfinish(m, int32(_a_F_transformColumnDefinition_9), int32(3960), int32(_a_F_transformColumnDefinition_10))
	mBase = m.M
	v1252 = m.ExcPending
	if v1252 != 0 {
		goto L1
	} else {
		goto L368
	}
L368:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L369:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v1259 = m.ExcPending
	if v1259 != 0 {
		goto L1
	} else {
		goto L370
	}
L370:
	;
	F_errmsg(m, int32(_a_F_transformColumnDefinition_28), int32(0))
	mBase = m.M
	v1263 = m.ExcPending
	if v1263 != 0 {
		goto L1
	} else {
		goto L371
	}
L371:
	;
	v1264 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1265 = *(*int32)(unsafe.Add(mBase, uint32(v359)+104))
	F_parser_errposition(m, v1264, v1265)
	mBase = m.M
	v1267 = m.ExcPending
	if v1267 != 0 {
		goto L1
	} else {
		goto L372
	}
L372:
	;
	F_errfinish(m, int32(_a_F_transformColumnDefinition_9), int32(3981), int32(_a_F_transformColumnDefinition_10))
	mBase = m.M
	v1272 = m.ExcPending
	if v1272 != 0 {
		goto L1
	} else {
		goto L373
	}
L373:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L374:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v1279 = m.ExcPending
	if v1279 != 0 {
		goto L1
	} else {
		goto L375
	}
L375:
	;
	F_errmsg(m, int32(_a_F_transformColumnDefinition_27), int32(0))
	mBase = m.M
	v1283 = m.ExcPending
	if v1283 != 0 {
		goto L1
	} else {
		goto L376
	}
L376:
	;
	v1284 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1285 = *(*int32)(unsafe.Add(mBase, uint32(v359)+104))
	F_parser_errposition(m, v1284, v1285)
	mBase = m.M
	v1287 = m.ExcPending
	if v1287 != 0 {
		goto L1
	} else {
		goto L377
	}
L377:
	;
	F_errfinish(m, int32(_a_F_transformColumnDefinition_9), int32(3986), int32(_a_F_transformColumnDefinition_10))
	mBase = m.M
	v1292 = m.ExcPending
	if v1292 != 0 {
		goto L1
	} else {
		goto L378
	}
L378:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L379:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v1299 = m.ExcPending
	if v1299 != 0 {
		goto L1
	} else {
		goto L380
	}
L380:
	;
	F_errmsg(m, int32(_a_F_transformColumnDefinition_29), int32(0))
	mBase = m.M
	v1303 = m.ExcPending
	if v1303 != 0 {
		goto L1
	} else {
		goto L381
	}
L381:
	;
	v1304 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1305 = *(*int32)(unsafe.Add(mBase, uint32(v359)+104))
	F_parser_errposition(m, v1304, v1305)
	mBase = m.M
	v1307 = m.ExcPending
	if v1307 != 0 {
		goto L1
	} else {
		goto L382
	}
L382:
	;
	F_errfinish(m, int32(_a_F_transformColumnDefinition_9), int32(4003), int32(_a_F_transformColumnDefinition_10))
	mBase = m.M
	v1312 = m.ExcPending
	if v1312 != 0 {
		goto L1
	} else {
		goto L383
	}
L383:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L384:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v1319 = m.ExcPending
	if v1319 != 0 {
		goto L1
	} else {
		goto L385
	}
L385:
	;
	F_errmsg(m, int32(_a_F_transformColumnDefinition_29), int32(0))
	mBase = m.M
	v1323 = m.ExcPending
	if v1323 != 0 {
		goto L1
	} else {
		goto L386
	}
L386:
	;
	v1324 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1325 = *(*int32)(unsafe.Add(mBase, uint32(v359)+104))
	F_parser_errposition(m, v1324, v1325)
	mBase = m.M
	v1327 = m.ExcPending
	if v1327 != 0 {
		goto L1
	} else {
		goto L387
	}
L387:
	;
	F_errfinish(m, int32(_a_F_transformColumnDefinition_9), int32(4020), int32(_a_F_transformColumnDefinition_10))
	mBase = m.M
	v1332 = m.ExcPending
	if v1332 != 0 {
		goto L1
	} else {
		goto L388
	}
L388:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L389:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v1345 = m.ExcPending
	if v1345 != 0 {
		goto L1
	} else {
		goto L390
	}
L390:
	;
	F_errmsg(m, int32(_a_F_transformColumnDefinition_30), int32(0))
	mBase = m.M
	v1349 = m.ExcPending
	if v1349 != 0 {
		goto L1
	} else {
		goto L391
	}
L391:
	;
	F_errfinish(m, int32(_a_F_transformColumnDefinition_9), int32(759), int32(_a_F_transformColumnDefinition_14))
	mBase = m.M
	v1354 = m.ExcPending
	if v1354 != 0 {
		goto L1
	} else {
		goto L392
	}
L392:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L393:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v1361 = m.ExcPending
	if v1361 != 0 {
		goto L1
	} else {
		goto L394
	}
L394:
	;
	v1362 = *(*int32)(unsafe.Add(mBase, uint32(v715)))
	v1363 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1364 = *(*int32)(unsafe.Add(mBase, uint32(v1363)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+132)) = v1364
	*(*int32)(unsafe.Add(mBase, uint32(v21)+128)) = v1362
	F_errmsg(m, int32(_a_F_transformColumnDefinition_13), v21+int32(128))
	mBase = m.M
	v1371 = m.ExcPending
	if v1371 != 0 {
		goto L1
	} else {
		goto L395
	}
L395:
	;
	v1372 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1373 = *(*int32)(unsafe.Add(mBase, uint32(v749)+104))
	F_parser_errposition(m, v1372, v1373)
	mBase = m.M
	v1375 = m.ExcPending
	if v1375 != 0 {
		goto L1
	} else {
		goto L396
	}
L396:
	;
	F_errfinish(m, int32(_a_F_transformColumnDefinition_9), int32(768), int32(_a_F_transformColumnDefinition_14))
	mBase = m.M
	v1380 = m.ExcPending
	if v1380 != 0 {
		goto L1
	} else {
		goto L397
	}
L397:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L398:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v1387 = m.ExcPending
	if v1387 != 0 {
		goto L1
	} else {
		goto L399
	}
L399:
	;
	v1388 = *(*int32)(unsafe.Add(mBase, uint32(v715)))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+80)) = v1388
	F_errmsg(m, int32(_a_F_transformColumnDefinition_31), v21+int32(80))
	mBase = m.M
	v1394 = m.ExcPending
	if v1394 != 0 {
		goto L1
	} else {
		goto L400
	}
L400:
	;
	F_errfinish(m, int32(_a_F_transformColumnDefinition_9), int32(774), int32(_a_F_transformColumnDefinition_14))
	mBase = m.M
	v1399 = m.ExcPending
	if v1399 != 0 {
		goto L1
	} else {
		goto L401
	}
L401:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L402:
	;
	v1404 = *(*int32)(unsafe.Add(mBase, uint32(v731)+8))
	v1405 = *(*int32)(unsafe.Add(mBase, uint32(v749)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+116)) = v1405
	*(*int32)(unsafe.Add(mBase, uint32(v21)+112)) = v1404
	F_errmsg_internal(m, int32(_a_F_transformColumnDefinition_32), v21+int32(112))
	mBase = m.M
	v1412 = m.ExcPending
	if v1412 != 0 {
		goto L1
	} else {
		goto L403
	}
L403:
	;
	F_errfinish(m, int32(_a_F_transformColumnDefinition_9), int32(803), int32(_a_F_transformColumnDefinition_14))
	mBase = m.M
	v1417 = m.ExcPending
	if v1417 != 0 {
		goto L1
	} else {
		goto L404
	}
L404:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L405:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v1424 = m.ExcPending
	if v1424 != 0 {
		goto L1
	} else {
		goto L406
	}
L406:
	;
	v1425 = *(*int32)(unsafe.Add(mBase, uint32(v715)))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+96)) = v1425
	F_errmsg(m, int32(_a_F_transformColumnDefinition_31), v21+int32(96))
	mBase = m.M
	v1431 = m.ExcPending
	if v1431 != 0 {
		goto L1
	} else {
		goto L407
	}
L407:
	;
	F_errfinish(m, int32(_a_F_transformColumnDefinition_9), int32(809), int32(_a_F_transformColumnDefinition_14))
	mBase = m.M
	v1436 = m.ExcPending
	if v1436 != 0 {
		goto L1
	} else {
		goto L408
	}
L408:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L409:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v1443 = m.ExcPending
	if v1443 != 0 {
		goto L1
	} else {
		goto L410
	}
L410:
	;
	v1444 = *(*int32)(unsafe.Add(mBase, uint32(v715)))
	v1445 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1446 = *(*int32)(unsafe.Add(mBase, uint32(v1445)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+148)) = v1446
	*(*int32)(unsafe.Add(mBase, uint32(v21)+144)) = v1444
	F_errmsg(m, int32(_a_F_transformColumnDefinition_33), v21+int32(144))
	mBase = m.M
	v1453 = m.ExcPending
	if v1453 != 0 {
		goto L1
	} else {
		goto L411
	}
L411:
	;
	v1454 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1455 = *(*int32)(unsafe.Add(mBase, uint32(v749)+104))
	F_parser_errposition(m, v1454, v1455)
	mBase = m.M
	v1457 = m.ExcPending
	if v1457 != 0 {
		goto L1
	} else {
		goto L412
	}
L412:
	;
	F_errfinish(m, int32(_a_F_transformColumnDefinition_9), int32(824), int32(_a_F_transformColumnDefinition_14))
	mBase = m.M
	v1462 = m.ExcPending
	if v1462 != 0 {
		goto L1
	} else {
		goto L413
	}
L413:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L414:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v1469 = m.ExcPending
	if v1469 != 0 {
		goto L1
	} else {
		goto L415
	}
L415:
	;
	F_errmsg(m, int32(_a_F_transformColumnDefinition_34), int32(0))
	mBase = m.M
	v1473 = m.ExcPending
	if v1473 != 0 {
		goto L1
	} else {
		goto L416
	}
L416:
	;
	F_errfinish(m, int32(_a_F_transformColumnDefinition_9), int32(838), int32(_a_F_transformColumnDefinition_14))
	mBase = m.M
	v1478 = m.ExcPending
	if v1478 != 0 {
		goto L1
	} else {
		goto L417
	}
L417:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L418:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v1485 = m.ExcPending
	if v1485 != 0 {
		goto L1
	} else {
		goto L419
	}
L419:
	;
	F_errmsg(m, int32(_a_F_transformColumnDefinition_35), int32(0))
	mBase = m.M
	v1489 = m.ExcPending
	if v1489 != 0 {
		goto L1
	} else {
		goto L420
	}
L420:
	;
	F_errfinish(m, int32(_a_F_transformColumnDefinition_9), int32(842), int32(_a_F_transformColumnDefinition_14))
	mBase = m.M
	v1494 = m.ExcPending
	if v1494 != 0 {
		goto L1
	} else {
		goto L421
	}
L421:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L422:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v1501 = m.ExcPending
	if v1501 != 0 {
		goto L1
	} else {
		goto L423
	}
L423:
	;
	v1502 = *(*int32)(unsafe.Add(mBase, uint32(v715)))
	v1503 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1504 = *(*int32)(unsafe.Add(mBase, uint32(v1503)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+164)) = v1504
	*(*int32)(unsafe.Add(mBase, uint32(v21)+160)) = v1502
	F_errmsg(m, int32(_a_F_transformColumnDefinition_36), v21+int32(160))
	mBase = m.M
	v1511 = m.ExcPending
	if v1511 != 0 {
		goto L1
	} else {
		goto L424
	}
L424:
	;
	v1512 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1513 = *(*int32)(unsafe.Add(mBase, uint32(v749)+104))
	F_parser_errposition(m, v1512, v1513)
	mBase = m.M
	v1515 = m.ExcPending
	if v1515 != 0 {
		goto L1
	} else {
		goto L425
	}
L425:
	;
	F_errfinish(m, int32(_a_F_transformColumnDefinition_9), int32(854), int32(_a_F_transformColumnDefinition_14))
	mBase = m.M
	v1520 = m.ExcPending
	if v1520 != 0 {
		goto L1
	} else {
		goto L426
	}
L426:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L427:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v1527 = m.ExcPending
	if v1527 != 0 {
		goto L1
	} else {
		goto L428
	}
L428:
	;
	F_errmsg(m, int32(_a_F_transformColumnDefinition_37), int32(0))
	mBase = m.M
	v1531 = m.ExcPending
	if v1531 != 0 {
		goto L1
	} else {
		goto L429
	}
L429:
	;
	F_errfinish(m, int32(_a_F_transformColumnDefinition_9), int32(884), int32(_a_F_transformColumnDefinition_14))
	mBase = m.M
	v1536 = m.ExcPending
	if v1536 != 0 {
		goto L1
	} else {
		goto L430
	}
L430:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L431:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v1543 = m.ExcPending
	if v1543 != 0 {
		goto L1
	} else {
		goto L432
	}
L432:
	;
	v1544 = *(*int32)(unsafe.Add(mBase, uint32(v715)))
	v1545 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1546 = *(*int32)(unsafe.Add(mBase, uint32(v1545)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+196)) = v1546
	*(*int32)(unsafe.Add(mBase, uint32(v21)+192)) = v1544
	F_errmsg(m, int32(_a_F_transformColumnDefinition_38), v21+int32(192))
	mBase = m.M
	v1553 = m.ExcPending
	if v1553 != 0 {
		goto L1
	} else {
		goto L433
	}
L433:
	;
	v1554 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1555 = *(*int32)(unsafe.Add(mBase, uint32(v749)+104))
	F_parser_errposition(m, v1554, v1555)
	mBase = m.M
	v1557 = m.ExcPending
	if v1557 != 0 {
		goto L1
	} else {
		goto L434
	}
L434:
	;
	F_errfinish(m, int32(_a_F_transformColumnDefinition_9), int32(891), int32(_a_F_transformColumnDefinition_14))
	mBase = m.M
	v1562 = m.ExcPending
	if v1562 != 0 {
		goto L1
	} else {
		goto L435
	}
L435:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L436:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v1569 = m.ExcPending
	if v1569 != 0 {
		goto L1
	} else {
		goto L437
	}
L437:
	;
	v1570 = *(*int32)(unsafe.Add(mBase, uint32(v715)))
	v1571 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1572 = *(*int32)(unsafe.Add(mBase, uint32(v1571)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+212)) = v1572
	*(*int32)(unsafe.Add(mBase, uint32(v21)+208)) = v1570
	F_errmsg(m, int32(_a_F_transformColumnDefinition_13), v21+int32(208))
	mBase = m.M
	v1579 = m.ExcPending
	if v1579 != 0 {
		goto L1
	} else {
		goto L438
	}
L438:
	;
	v1580 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1581 = *(*int32)(unsafe.Add(mBase, uint32(v749)+104))
	F_parser_errposition(m, v1580, v1581)
	mBase = m.M
	v1583 = m.ExcPending
	if v1583 != 0 {
		goto L1
	} else {
		goto L439
	}
L439:
	;
	F_errfinish(m, int32(_a_F_transformColumnDefinition_9), int32(909), int32(_a_F_transformColumnDefinition_14))
	mBase = m.M
	v1588 = m.ExcPending
	if v1588 != 0 {
		goto L1
	} else {
		goto L440
	}
L440:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L441:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v1595 = m.ExcPending
	if v1595 != 0 {
		goto L1
	} else {
		goto L442
	}
L442:
	;
	F_errmsg(m, int32(_a_F_transformColumnDefinition_39), int32(0))
	mBase = m.M
	v1599 = m.ExcPending
	if v1599 != 0 {
		goto L1
	} else {
		goto L443
	}
L443:
	;
	v1600 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1601 = *(*int32)(unsafe.Add(mBase, uint32(v749)+104))
	F_parser_errposition(m, v1600, v1601)
	mBase = m.M
	v1603 = m.ExcPending
	if v1603 != 0 {
		goto L1
	} else {
		goto L444
	}
L444:
	;
	F_errfinish(m, int32(_a_F_transformColumnDefinition_9), int32(926), int32(_a_F_transformColumnDefinition_14))
	mBase = m.M
	v1608 = m.ExcPending
	if v1608 != 0 {
		goto L1
	} else {
		goto L445
	}
L445:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L446:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v1615 = m.ExcPending
	if v1615 != 0 {
		goto L1
	} else {
		goto L447
	}
L447:
	;
	F_errmsg(m, int32(_a_F_transformColumnDefinition_40), int32(0))
	mBase = m.M
	v1619 = m.ExcPending
	if v1619 != 0 {
		goto L1
	} else {
		goto L448
	}
L448:
	;
	v1620 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1621 = *(*int32)(unsafe.Add(mBase, uint32(v749)+104))
	F_parser_errposition(m, v1620, v1621)
	mBase = m.M
	v1623 = m.ExcPending
	if v1623 != 0 {
		goto L1
	} else {
		goto L449
	}
L449:
	;
	F_errfinish(m, int32(_a_F_transformColumnDefinition_9), int32(943), int32(_a_F_transformColumnDefinition_14))
	mBase = m.M
	v1628 = m.ExcPending
	if v1628 != 0 {
		goto L1
	} else {
		goto L450
	}
L450:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L451:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v1635 = m.ExcPending
	if v1635 != 0 {
		goto L1
	} else {
		goto L452
	}
L452:
	;
	v1636 = *(*int32)(unsafe.Add(mBase, uint32(v715)))
	v1637 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1638 = *(*int32)(unsafe.Add(mBase, uint32(v1637)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+36)) = v1638
	*(*int32)(unsafe.Add(mBase, uint32(v21)+32)) = v1636
	F_errmsg(m, int32(_a_F_transformColumnDefinition_41), v21+int32(32))
	mBase = m.M
	v1645 = m.ExcPending
	if v1645 != 0 {
		goto L1
	} else {
		goto L453
	}
L453:
	;
	v1646 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1647 = *(*int32)(unsafe.Add(mBase, uint32(v749)+104))
	F_parser_errposition(m, v1646, v1647)
	mBase = m.M
	v1649 = m.ExcPending
	if v1649 != 0 {
		goto L1
	} else {
		goto L454
	}
L454:
	;
	F_errfinish(m, int32(_a_F_transformColumnDefinition_9), int32(974), int32(_a_F_transformColumnDefinition_14))
	mBase = m.M
	v1654 = m.ExcPending
	if v1654 != 0 {
		goto L1
	} else {
		goto L455
	}
L455:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L456:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v1661 = m.ExcPending
	if v1661 != 0 {
		goto L1
	} else {
		goto L457
	}
L457:
	;
	v1662 = *(*int32)(unsafe.Add(mBase, uint32(v715)))
	v1663 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1664 = *(*int32)(unsafe.Add(mBase, uint32(v1663)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+52)) = v1664
	*(*int32)(unsafe.Add(mBase, uint32(v21)+48)) = v1662
	F_errmsg(m, int32(_a_F_transformColumnDefinition_42), v21+int32(48))
	mBase = m.M
	v1671 = m.ExcPending
	if v1671 != 0 {
		goto L1
	} else {
		goto L458
	}
L458:
	;
	v1672 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1673 = *(*int32)(unsafe.Add(mBase, uint32(v749)+104))
	F_parser_errposition(m, v1672, v1673)
	mBase = m.M
	v1675 = m.ExcPending
	if v1675 != 0 {
		goto L1
	} else {
		goto L459
	}
L459:
	;
	F_errfinish(m, int32(_a_F_transformColumnDefinition_9), int32(982), int32(_a_F_transformColumnDefinition_14))
	mBase = m.M
	v1680 = m.ExcPending
	if v1680 != 0 {
		goto L1
	} else {
		goto L460
	}
L460:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L461:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v1687 = m.ExcPending
	if v1687 != 0 {
		goto L1
	} else {
		goto L462
	}
L462:
	;
	v1688 = *(*int32)(unsafe.Add(mBase, uint32(v715)))
	v1689 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1690 = *(*int32)(unsafe.Add(mBase, uint32(v1689)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+20)) = v1690
	*(*int32)(unsafe.Add(mBase, uint32(v21)+16)) = v1688
	F_errmsg(m, int32(_a_F_transformColumnDefinition_13), v21+int32(16))
	mBase = m.M
	v1697 = m.ExcPending
	if v1697 != 0 {
		goto L1
	} else {
		goto L463
	}
L463:
	;
	v1698 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1699 = *(*int32)(unsafe.Add(mBase, uint32(v749)+104))
	F_parser_errposition(m, v1698, v1699)
	mBase = m.M
	v1701 = m.ExcPending
	if v1701 != 0 {
		goto L1
	} else {
		goto L464
	}
L464:
	;
	F_errfinish(m, int32(_a_F_transformColumnDefinition_9), int32(750), int32(_a_F_transformColumnDefinition_14))
	mBase = m.M
	v1706 = m.ExcPending
	if v1706 != 0 {
		goto L1
	} else {
		goto L465
	}
L465:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L466:
	;
	if v1719&int32(1) == int32(0) {
		v1739 = v715
		v1742 = v711
		v1743 = v713
		goto L4
	} else {
		goto L467
	}
L467:
	;
	v1733 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v713))))
	if v1733 != 0 {
		goto L3
	} else {
		goto L468
	}
L468:
	;
	v1739 = v715
	v1742 = v711
	v1743 = v713
	goto L4
L469:
	;
	v1757 = F_makeNotNullConstraint(m, v1755)
	mBase = m.M
	v1758 = m.ExcPending
	if v1758 != 0 {
		goto L1
	} else {
		goto L470
	}
L470:
	;
	v1759 = *(*int32)(unsafe.Add(mBase, uint32(v1742)))
	v1760 = F_lappend(m, v1759, v1757)
	mBase = m.M
	v1761 = m.ExcPending
	if v1761 != 0 {
		goto L1
	} else {
		goto L471
	}
L471:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1742))) = v1760
	goto L3
L472:
	;
	v1783 = F_palloc0(m, int32(32))
	mBase = m.M
	v1784 = m.ExcPending
	if v1784 != 0 {
		goto L1
	} else {
		goto L475
	}
L473:
	;
	goto L474
L474:
	;
	m.G0 = v21 + int32(272)
	return
L475:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1783))) = int64(107374182547)
	v1787 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1783)+8)) = v1787
	v1789 = *(*int32)(unsafe.Add(mBase, uint32(l1)+60))
	v1790 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1783)+28)) = uint8(v1790)
	*(*int32)(unsafe.Add(mBase, uint32(v1783)+24)) = v1790
	*(*int32)(unsafe.Add(mBase, uint32(v1783)+20)) = v1789
	v1796 = F_palloc0(m, int32(20))
	mBase = m.M
	v1797 = m.ExcPending
	if v1797 != 0 {
		goto L1
	} else {
		goto L476
	}
L476:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1796))) = int32(146)
	v1800 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v1796)+8)) = int64(77309411328)
	*(*int32)(unsafe.Add(mBase, uint32(v1796)+4)) = v1800
	v1805 = F_lappend(m, int32(0), v1783)
	mBase = m.M
	v1806 = m.ExcPending
	if v1806 != 0 {
		goto L1
	} else {
		goto L477
	}
L477:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1796)+8)) = v1805
	v1808 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v1809 = F_lappend(m, v1808, v1796)
	mBase = m.M
	v1810 = m.ExcPending
	if v1810 != 0 {
		goto L1
	} else {
		goto L478
	}
L478:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v1809
	goto L474
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
