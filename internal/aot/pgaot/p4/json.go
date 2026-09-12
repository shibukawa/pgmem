package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_GetJsonPathVar(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
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
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v131 int32
	_ = v131
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
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v241 int32
	_ = v241
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v271 int32
	_ = v271
	var v276 int32
	_ = v276
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v296 int32
	_ = v296
	var v301 int32
	_ = v301
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v331 int32
	_ = v331
	var v334 int32
	_ = v334
	var v336 int32
	_ = v336
	var v345 int32
	_ = v345
	var v348 int32
	_ = v348
	var v352 int32
	_ = v352
	var v357 int32
	_ = v357
	var v366 int64
	_ = v366
	var v368 int32
	_ = v368
	var v370 int64
	_ = v370
	var v389 int32
	_ = v389
	var v391 int32
	_ = v391
	v6 = int32(0)
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	if l0 == v6 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v391
	m.G0 = v13 + int32(16)
	return v389
L2:
	;
	v389 = int32(0)
	v391 = int32(-1)
	goto L1
L3:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v18 <= int32(0) {
		v389 = v6
		v391 = int32(-1)
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v21 = int32(0)
	if v21 < v18 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v24 = v18
	goto L7
L6:
	;
	v24 = v21
	goto L7
L7:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v32 = v6
	v34 = int32(1)
	goto L8
L8:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v25+v32<<(uint(int32(2))%32))))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
	if l2 == v41 {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	v96 = F_palloc(m, int32(20))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L30
	} else {
		goto L31
	}
L10:
	;
	goto L9
L11:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
	if l2 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L12:
	;
	goto L13
L13:
	;
	v90 = int32(1)
	v93 = v32 + v90
	if v24 != v93 {
		v32 = v93
		v34 = v34 + v90
		goto L8
	} else {
		goto L29
	}
L14:
	;
	if v87 == int32(0) {
		goto L10
	} else {
		goto L28
	}
L15:
	;
	v87 = int32(0)
	goto L14
L16:
	;
	goto L17
L17:
	;
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43))))
	if v49 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v50 = v43
	v51 = l1
	v52 = l2
	v53 = v49
	goto L22
L19:
	;
	v75 = l1
	v79 = int32(0)
	goto L20
L20:
	;
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75))))
	v87 = v79 - v80
	goto L14
L21:
	;
	v75 = v70
	v79 = v72
	goto L20
L22:
	;
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51))))
	if v53 != v55 {
		v70 = v51
		v72 = v53
		goto L21
	} else {
		goto L24
	}
L23:
	;
	v70 = v64
	v72 = int32(0)
	goto L21
L24:
	;
	if v55 == int32(0) {
		v70 = v51
		v72 = v53
		goto L21
	} else {
		goto L25
	}
L25:
	;
	v60 = v52 - int32(1)
	if v60 == int32(0) {
		v70 = v51
		v72 = v53
		goto L21
	} else {
		goto L26
	}
L26:
	;
	v63 = int32(1)
	v64 = v51 + v63
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50)+1)))
	if v65 != 0 {
		v50 = v50 + v63
		v51 = v64
		v52 = v60
		v53 = v65
		goto L22
	} else {
		goto L27
	}
L27:
	;
	goto L23
L28:
	;
	goto L13
L29:
	;
	goto L2
L30:
	;
	return int32(0)
L31:
	;
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+20)))
	if v100 == int32(1) {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	v366 = *(*int64)(unsafe.Add(mBase, uint32(v96)))
	*(*int64)(unsafe.Add(mBase, uint32(l3))) = v366
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v96)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l3)+16)) = v368
	v370 = *(*int64)(unsafe.Add(mBase, uint32(v96)+8))
	*(*int64)(unsafe.Add(mBase, uint32(l3)+8)) = v370
	v389 = v96
	v391 = v34
	goto L1
L33:
	;
	v103 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v103
	*(*int32)(unsafe.Add(mBase, uint32(v96))) = v103
	goto L32
L34:
	;
	goto L35
L35:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v40)+16))
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v40)+8))
	if v108 <= int32(1042) {
		goto L47
	} else {
		goto L48
	}
L36:
	;
	v309 = F_pg_detoast_datum(m, v107)
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L30
	} else {
		goto L113
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v96))) = int32(3)
	*(*uint8)(unsafe.Add(mBase, uint32(v96)+4)) = uint8(base.B2i32(v107 != int32(0)))
	goto L32
L38:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L30
	} else {
		goto L108
	}
L39:
	;
	if v108 == int32(114) {
		goto L36
	} else {
		goto L107
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v96)+8)) = v146
	*(*int32)(unsafe.Add(mBase, uint32(v96))) = int32(18)
	v250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v143))))
	if v250 == int32(1) {
		goto L95
	} else {
		goto L96
	}
L41:
	;
	v204 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v96))) = v204
	v207 = v107 + v204
	v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107))))
	if v210&v204 != 0 {
		goto L80
	} else {
		goto L81
	}
L42:
	;
	v197 = F_DirectFunctionCall1Coll(m, int32(1417), int32(0), v107)
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L30
	} else {
		goto L78
	}
L43:
	;
	v188 = F_DirectFunctionCall1Coll(m, int32(1416), int32(0), v107)
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L30
	} else {
		goto L76
	}
L44:
	;
	v179 = F_DirectFunctionCall1Coll(m, int32(1415), int32(0), v107)
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L30
	} else {
		goto L74
	}
L45:
	;
	v170 = F_DirectFunctionCall1Coll(m, int32(1414), int32(0), v107)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L30
	} else {
		goto L72
	}
L46:
	;
	v161 = F_DirectFunctionCall1Coll(m, int32(1413), int32(0), v107)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L30
	} else {
		goto L70
	}
L47:
	;
	switch v108 - int32(16) {
	case 0:
		goto L37
	case 1, 2, 3, 6, 8:
		goto L38
	case 4:
		goto L44
	case 5:
		goto L46
	case 7:
		goto L45
	case 9:
		goto L41
	default:
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	if v108 <= int32(1183) {
		goto L53
	} else {
		goto L54
	}
L50:
	;
	switch v108 - int32(700) {
	case 0:
		goto L43
	case 1:
		goto L42
	default:
		goto L39
	}
L51:
	;
	if v108 != int32(1700) {
		goto L62
	} else {
		goto L63
	}
L52:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v40)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v96)+16)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v96)+12)) = v131
	*(*int32)(unsafe.Add(mBase, uint32(v96)+8)) = v108
	*(*int32)(unsafe.Add(mBase, uint32(v96)+4)) = v107
	*(*int32)(unsafe.Add(mBase, uint32(v96))) = int32(32)
	goto L32
L53:
	;
	if base.Ui32(v108-int32(1082)) < base.Ui32(int32(2)) {
		goto L52
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	if int32(1699) < v108 {
		goto L51
	} else {
		goto L59
	}
L56:
	;
	if v108 == int32(1043) {
		goto L41
	} else {
		goto L57
	}
L57:
	;
	if v108 == int32(1114) {
		goto L52
	} else {
		goto L58
	}
L58:
	;
	goto L38
L59:
	;
	if v108 == int32(1184) {
		goto L52
	} else {
		goto L60
	}
L60:
	;
	if v108 != int32(1266) {
		goto L38
	} else {
		goto L61
	}
L61:
	;
	goto L52
L62:
	;
	if v108 != int32(3802) {
		goto L38
	} else {
		goto L65
	}
L63:
	;
	goto L64
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v96))) = int32(2)
	v156 = F_pg_detoast_datum(m, v107)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L30
	} else {
		goto L69
	}
L65:
	;
	v143 = F_pg_detoast_datum(m, v107)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L30
	} else {
		goto L66
	}
L66:
	;
	v146 = v143 + int32(4)
	v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v143)+7)))
	if v147&int32(16) == int32(0) {
		goto L40
	} else {
		goto L67
	}
L67:
	;
	v152 = F_JsonbExtractScalar(m, v146, v96)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L30
	} else {
		goto L68
	}
L68:
	;
	goto L32
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v96)+4)) = v156
	goto L32
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v96))) = int32(2)
	v165 = F_pg_detoast_datum(m, v161)
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L30
	} else {
		goto L71
	}
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v96)+4)) = v165
	goto L32
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v96))) = int32(2)
	v174 = F_pg_detoast_datum(m, v170)
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L30
	} else {
		goto L73
	}
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v96)+4)) = v174
	goto L32
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v96))) = int32(2)
	v183 = F_pg_detoast_datum(m, v179)
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L30
	} else {
		goto L75
	}
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v96)+4)) = v183
	goto L32
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v96))) = int32(2)
	v192 = F_pg_detoast_datum(m, v188)
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L30
	} else {
		goto L77
	}
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v96)+4)) = v192
	goto L32
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v96))) = int32(2)
	v201 = F_pg_detoast_datum(m, v197)
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L30
	} else {
		goto L79
	}
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v96)+4)) = v201
	goto L32
L80:
	;
	v213 = v207
	goto L82
L81:
	;
	v213 = v107 + int32(4)
	goto L82
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v96)+8)) = v213
	v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107))))
	if v215 == int32(1) {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v218 = int32(4)
	v220 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207))))
	if v220&int32(254) == int32(2) {
		goto L86
	} else {
		goto L87
	}
L84:
	;
	goto L85
L85:
	;
	if v215&int32(1) != 0 {
		goto L92
	} else {
		goto L93
	}
L86:
	;
	v229 = v218
	goto L88
L87:
	;
	v229 = base.B2i32(v220 == int32(18)) << (uint(v218) % 32)
	goto L88
L88:
	;
	if v220 == int32(1) {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	v232 = v218
	goto L91
L90:
	;
	v232 = v229
	goto L91
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v96)+4)) = v232
	goto L32
L92:
	;
	v236 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v96)+4)) = int32(base.Ui32(v215)>>(uint(v236)%32)) - v236
	goto L32
L93:
	;
	goto L94
L94:
	;
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v107)))
	*(*int32)(unsafe.Add(mBase, uint32(v96)+4)) = int32(base.Ui32(v241)>>(uint(int32(2))%32)) - int32(4)
	goto L32
L95:
	;
	v253 = int32(4)
	v255 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v143)+1)))
	if v255&int32(254) == int32(2) {
		goto L98
	} else {
		goto L99
	}
L96:
	;
	goto L97
L97:
	;
	if v250&int32(1) != 0 {
		goto L104
	} else {
		goto L105
	}
L98:
	;
	v264 = v253
	goto L100
L99:
	;
	v264 = base.B2i32(v255 == int32(18)) << (uint(v253) % 32)
	goto L100
L100:
	;
	if v255 == int32(1) {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	v267 = v253
	goto L103
L102:
	;
	v267 = v264
	goto L103
L103:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v96)+4)) = v267
	goto L32
L104:
	;
	v271 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v96)+4)) = int32(base.Ui32(v250)>>(uint(v271)%32)) - v271
	goto L32
L105:
	;
	goto L106
L106:
	;
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v143)))
	*(*int32)(unsafe.Add(mBase, uint32(v96)+4)) = int32(base.Ui32(v276)>>(uint(int32(2))%32)) - int32(4)
	goto L32
L107:
	;
	goto L38
L108:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L30
	} else {
		goto L109
	}
L109:
	;
	v291 = F_format_type_be(m, v108)
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L30
	} else {
		goto L110
	}
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v291
	F_errmsg(m, int32(322265), v13)
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L30
	} else {
		goto L111
	}
L111:
	;
	F_errfinish(m, int32(501542), int32(3124), int32(287270))
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L30
	} else {
		goto L112
	}
L112:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L113:
	;
	v311 = F_text_to_cstring(m, v309)
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L30
	} else {
		goto L114
	}
L114:
	;
	v313 = F_DirectFunctionCall1Coll(m, int32(486), int32(0), v311)
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L30
	} else {
		goto L115
	}
L115:
	;
	v315 = F_pg_detoast_datum(m, v313)
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L30
	} else {
		goto L116
	}
L116:
	;
	F_pfree(m, v311)
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L30
	} else {
		goto L117
	}
L117:
	;
	v319 = F_pg_detoast_datum(m, v315)
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L30
	} else {
		goto L118
	}
L118:
	;
	v322 = v319 + int32(4)
	v323 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v319)+7)))
	if v323&int32(16) != 0 {
		goto L119
	} else {
		goto L120
	}
L119:
	;
	v326 = F_JsonbExtractScalar(m, v322, v96)
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L30
	} else {
		goto L122
	}
L120:
	;
	goto L121
L121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v96)+8)) = v322
	*(*int32)(unsafe.Add(mBase, uint32(v96))) = int32(18)
	v331 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v319))))
	if v331 == int32(1) {
		goto L123
	} else {
		goto L124
	}
L122:
	;
	goto L32
L123:
	;
	v334 = int32(4)
	v336 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v319)+1)))
	if v336&int32(254) == int32(2) {
		goto L126
	} else {
		goto L127
	}
L124:
	;
	goto L125
L125:
	;
	if v331&int32(1) != 0 {
		goto L132
	} else {
		goto L133
	}
L126:
	;
	v345 = v334
	goto L128
L127:
	;
	v345 = base.B2i32(v336 == int32(18)) << (uint(v334) % 32)
	goto L128
L128:
	;
	if v336 == int32(1) {
		goto L129
	} else {
		goto L130
	}
L129:
	;
	v348 = v334
	goto L131
L130:
	;
	v348 = v345
	goto L131
L131:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v96)+4)) = v348
	goto L32
L132:
	;
	v352 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v96)+4)) = int32(base.Ui32(v331)>>(uint(v352)%32)) - v352
	goto L32
L133:
	;
	goto L134
L134:
	;
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v319)))
	*(*int32)(unsafe.Add(mBase, uint32(v96)+4)) = int32(base.Ui32(v357)>>(uint(int32(2))%32)) - int32(4)
	goto L32
}
func F_JsonTableFetchRow(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	v3 = F_GetJsonTableExecContext(m, l0, int32(32260))
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(v3)+4))
		v8 = F_JsonTablePlanNextRow(m, v7)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			return v8
		}
	}
}
func F__equalJsonIsPredicate(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	v3 = int32(0)
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v6 = F_equal(m, v4, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		if v6 == int32(0) {
			v24 = v3
			return v24
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
			v14 = F_equal(m, v12, v13)
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				if v14 == int32(0) {
					v24 = v3
				} else {
					v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
					if v18 != v19 {
						v24 = v3
					} else {
						v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)))
						v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+16)))
						v24 = base.B2i32(v21 == v22)
					}
				}
				return v24
			}
		}
	}
}
func F__equalJsonObjectConstructor(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	v3 = int32(0)
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v6 = F_equal(m, v4, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		if v6 == int32(0) {
			v24 = v3
			return v24
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
			v14 = F_equal(m, v12, v13)
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				if v14 == int32(0) {
					v24 = v3
				} else {
					v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)))
					v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+12)))
					if v18 != v19 {
						v24 = v3
					} else {
						v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+13)))
						v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+13)))
						v24 = base.B2i32(v21 == v22)
					}
				}
				return v24
			}
		}
	}
}
func F_get_json_agg_constructor(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	F_initStringInfo(m, v10+int32(16))
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
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v17 == int32(1) {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+25)))
	if v32 == int32(1) {
		goto L9
	} else {
		goto L10
	}
L4:
	;
	F_appendStringInfoString(m, v10+int32(16), v26)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L8
	}
L5:
	;
	switch v16 - int32(1) {
	case 0, 2:
		v26 = int32(534551)
		goto L4
	default:
		goto L3
	}
L6:
	;
	goto L7
L7:
	;
	switch v16 - int32(2) {
	case 0, 2:
		v26 = int32(534567)
		goto L4
	default:
		goto L3
	}
L8:
	;
	goto L3
L9:
	;
	F_appendStringInfoString(m, v10+int32(16), int32(523796))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L1
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if base.Ui32(int32(2)) <= base.Ui32(v40-int32(5)) {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	goto L11
L13:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	F_get_json_returning(m, v45, v10+int32(16), int32(1))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L1
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
	switch v52 - int32(9) {
	case 0:
		goto L18
	default:
		goto L19
	case 2:
		goto L20
	}
L16:
	;
	goto L15
L17:
	;
	m.G0 = v10 + int32(32)
	return
L18:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
	F_get_agg_expr_helper(m, v51, l1, v51, l2, v73, l3)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L1
	} else {
		goto L25
	}
L19:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L1
	} else {
		goto L22
	}
L20:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
	F_get_windowfunc_expr_helper(m, v51, l1, l2, v55, l3)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	goto L17
L22:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)))
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v63
	F_errmsg_internal(m, int32(487269), v10)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	F_errfinish(m, int32(495391), int32(11786), int32(208495))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L25:
	;
	goto L17
}
func F_get_json_table_nested_columns(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	switch v11 - int32(50) {
	case 0:
		goto L4
	case 1:
		goto L5
	default:
		goto L1
	}
L1:
	;
	m.G0 = v9 + int32(16)
	return
L2:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	F_appendStringInfoChar(m, v55, int32(32))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L6
	} else {
		goto L16
	}
L3:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	F_appendStringInfoChar(m, v45, int32(44))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L6
	} else {
		goto L15
	}
L4:
	;
	if l4 == int32(0) {
		v50 = l1
		goto L2
	} else {
		goto L14
	}
L5:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	F_get_json_table_nested_columns(m, l0, v14, l2, l3, l4)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	return
L7:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	switch v18 - int32(50) {
	case 0:
		v40 = v17
		goto L3
	case 1:
		goto L8
	default:
		goto L1
	}
L8:
	;
	v22 = v17
	goto L9
L9:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	F_get_json_table_nested_columns(m, l0, v27, l2, l3, int32(1))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L6
	} else {
		goto L11
	}
L10:
	;
	if v32 == int32(50) {
		v40 = v31
		goto L3
	} else {
		goto L13
	}
L11:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
	if v32 == int32(51) {
		v22 = v31
		goto L9
	} else {
		goto L12
	}
L12:
	;
	goto L10
L13:
	;
	goto L1
L14:
	;
	v40 = l1
	goto L3
L15:
	;
	v50 = v40
	goto L2
L16:
	;
	v60 = int32(0)
	F_appendContextKeyword(m, l2, int32(746766), v60, v60, v60)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L6
	} else {
		goto L17
	}
L17:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)+4))
	F_get_const_expr(m, v66, l2, int32(-1))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L6
	} else {
		goto L18
	}
L18:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v71)+8))
	v73 = F_quote_identifier(m, v72)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L6
	} else {
		goto L19
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v73
	F_appendStringInfo(m, v70, int32(198553), v9)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L6
	} else {
		goto L20
	}
L20:
	;
	F_get_json_table_columns(m, l0, v50, l2, l3)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L6
	} else {
		goto L21
	}
L21:
	;
	goto L1
}
func F_json_array_elements(m *base.Module, l0 int32) int32 {
	var v7 int32
	_ = v7
	F_elements_worker(m, l0, int32(123446), int32(0))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return int32(0)
	}
}
func F_json_array_elements_text(m *base.Module, l0 int32) int32 {
	var v7 int32
	_ = v7
	F_elements_worker(m, l0, int32(63362), int32(1))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return int32(0)
	}
}
func F_json_build_array_noargs(m *base.Module, l0 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = F_cstring_to_text_with_len(m, int32(509225), int32(2))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_json_each(m *base.Module, l0 int32) int32 {
	var v6 int32
	_ = v6
	F_each_worker(m, l0, int32(0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return int32(0)
	}
}
func F_json_in(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	v5 = m.G0
	v7 = v5 - int32(80)
	m.G0 = v7
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = F_cstring_to_text(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		F_makeJsonLexContext(m, v7+int32(12), v12, int32(0))
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v23 = F_pg_parse_json_or_errsave(m, v7+int32(12), int32(1859584), v22)
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int32(0)
			} else {
				if v23 == int32(0) {
					v27 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v27)
					v29 = int32(0)
				} else {
					v29 = v12
				}
				m.G0 = v7 + int32(80)
				return v29
			}
		}
	}
}
func F_json_lex_number(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v274 int32
	_ = v274
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v298 int32
	_ = v298
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v307 int32
	_ = v307
	var v312 int32
	_ = v312
	var v321 int32
	_ = v321
	var v327 int32
	_ = v327
	v11 = int32(1)
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v13 = l1 - v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if base.Ui32(v14) <= base.Ui32(v13) {
		v67 = l1
		v70 = v13
		v75 = v11
		goto L2
	} else {
		goto L3
	}
L1:
	;
	if l3 != 0 {
		goto L76
	} else {
		goto L77
	}
L2:
	;
	if base.Ui32(v14) <= base.Ui32(v70) {
		v131 = v67
		v134 = v70
		v139 = v75
		goto L15
	} else {
		goto L16
	}
L3:
	;
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v16 == int32(48) {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v67 = v56
	v70 = v59
	v75 = int32(0)
	goto L2
L5:
	;
	v19 = int32(1)
	v56 = l1 + v19
	v59 = v13 + v19
	goto L4
L6:
	;
	goto L7
L7:
	;
	if base.Ui32(int32(8)) < base.Ui32((v16-int32(49))&int32(255)) {
		v67 = l1
		v70 = v13
		v75 = v11
		goto L2
	} else {
		goto L8
	}
L8:
	;
	v33 = l1
	v36 = v13
	goto L9
L9:
	;
	if v36 == v14-int32(1) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v56 = v47
	v59 = v45
	goto L4
L11:
	;
	v292 = v12 + v14
	v293 = v14
	v294 = int32(0)
	goto L1
L12:
	;
	goto L13
L13:
	;
	v44 = int32(1)
	v45 = v36 + v44
	v47 = v33 + v44
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47))))
	if base.Ui32((v48-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v33 = v47
		v36 = v45
		goto L9
	} else {
		goto L14
	}
L14:
	;
	goto L10
L15:
	;
	if base.Ui32(v14) <= base.Ui32(v134) {
		goto L35
	} else {
		goto L36
	}
L16:
	;
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67))))
	if v77 != int32(46) {
		v131 = v67
		v134 = v70
		v139 = v75
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v80 = int32(1)
	v82 = v67 + v80
	v84 = v70 + v80
	if v14 == v84 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v292 = v82
	v293 = v14
	v294 = v80
	goto L1
L19:
	;
	goto L20
L20:
	;
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82))))
	if base.Ui32((v86-int32(58))&int32(255)) < base.Ui32(int32(246)) {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v131 = v121
	v134 = v125
	v139 = v129
	goto L15
L22:
	;
	v121 = v82
	v125 = v84
	v129 = int32(1)
	goto L21
L23:
	;
	goto L24
L24:
	;
	v95 = v70 + int32(2)
	if base.Ui32(v95) < base.Ui32(v14) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v97 = v14
	goto L27
L26:
	;
	v97 = v95
	goto L27
L27:
	;
	v99 = v82
	v103 = v84
	goto L28
L28:
	;
	v108 = int32(1)
	v109 = v99 + v108
	v111 = v103 + v108
	if base.Ui32(v14) <= base.Ui32(v111) {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	v121 = v109
	v125 = v111
	v129 = v75
	goto L21
L30:
	;
	v292 = v109
	v293 = v97
	v294 = v75
	goto L1
L31:
	;
	goto L32
L32:
	;
	v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v109))))
	if base.Ui32((v113-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v99 = v109
		v103 = v111
		goto L28
	} else {
		goto L33
	}
L33:
	;
	goto L29
L34:
	;
	if base.Ui32(v14) <= base.Ui32(v210) {
		goto L58
	} else {
		goto L59
	}
L35:
	;
	v210 = v134
	v211 = v131
	v215 = v139
	goto L34
L36:
	;
	goto L37
L37:
	;
	v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v131))))
	if v141|int32(32) != int32(101) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v210 = v134
	v211 = v131
	v215 = v139
	goto L34
L39:
	;
	goto L40
L40:
	;
	v146 = int32(1)
	v148 = v131 + v146
	v150 = v134 + v146
	if base.Ui32(v14) <= base.Ui32(v150) {
		v159 = v148
		v160 = v150
		goto L41
	} else {
		goto L42
	}
L41:
	;
	if v14 == v160 {
		v292 = v159
		v293 = v14
		v294 = v146
		goto L1
	} else {
		goto L44
	}
L42:
	;
	v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148))))
	switch v152 - int32(43) {
	case 0, 2:
		goto L43
	default:
		v159 = v148
		v160 = v150
		goto L41
	}
L43:
	;
	v155 = int32(2)
	v159 = v131 + v155
	v160 = v134 + v155
	goto L41
L44:
	;
	v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v159))))
	if base.Ui32((v162-int32(58))&int32(255)) < base.Ui32(int32(246)) {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	v210 = v204
	v211 = v201
	v215 = v205
	goto L34
L46:
	;
	v201 = v159
	v204 = v160
	v205 = int32(1)
	goto L45
L47:
	;
	goto L48
L48:
	;
	v171 = v160 + int32(1)
	if base.Ui32(v171) < base.Ui32(v14) {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v173 = v14
	goto L51
L50:
	;
	v173 = v171
	goto L51
L51:
	;
	v179 = v159
	v182 = v160
	goto L52
L52:
	;
	v184 = int32(1)
	v185 = v179 + v184
	v187 = v182 + v184
	if base.Ui32(v14) <= base.Ui32(v187) {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	v201 = v185
	v204 = v187
	v205 = v139
	goto L45
L54:
	;
	v292 = v185
	v293 = v173
	v294 = v139
	goto L1
L55:
	;
	goto L56
L56:
	;
	v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v185))))
	if base.Ui32((v189-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v179 = v185
		v182 = v187
		goto L52
	} else {
		goto L57
	}
L57:
	;
	goto L53
L58:
	;
	v292 = v211
	v293 = v210
	v294 = v215
	goto L1
L59:
	;
	v217 = int32(*(*int8)(unsafe.Add(mBase, uint32(v211))))
	if base.Ui32((v217-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v238 = int32(1)
	v240 = v210 + v238
	if v14 != v240 {
		goto L65
	} else {
		goto L66
	}
L61:
	;
	if base.Ui32((v217&int32(-33)-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L60
	} else {
		goto L62
	}
L62:
	;
	if v217 == int32(95) {
		goto L60
	} else {
		goto L63
	}
L63:
	;
	if int32(0) <= v217 {
		goto L58
	} else {
		goto L64
	}
L64:
	;
	goto L60
L65:
	;
	v246 = v240
	v247 = v211
	goto L68
L66:
	;
	goto L67
L67:
	;
	v292 = v211 + (v14 - v210)
	v293 = v14
	v294 = v238
	goto L1
L68:
	;
	v253 = v247 + int32(1)
	v254 = int32(*(*int8)(unsafe.Add(mBase, uint32(v253))))
	if base.Ui32((v254-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		goto L70
	} else {
		goto L71
	}
L69:
	;
	goto L67
L70:
	;
	v274 = v246 + int32(1)
	if v274 != v14 {
		v246 = v274
		v247 = v253
		goto L68
	} else {
		goto L75
	}
L71:
	;
	if base.Ui32((v254&int32(-33)-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L70
	} else {
		goto L72
	}
L72:
	;
	if v254 == int32(95) {
		goto L70
	} else {
		goto L73
	}
L73:
	;
	if v254 < int32(0) {
		goto L70
	} else {
		goto L74
	}
L74:
	;
	v292 = v253
	v293 = v246
	v294 = v238
	goto L1
L75:
	;
	goto L69
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v293
	goto L78
L77:
	;
	goto L78
L78:
	;
	v298 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v298 != int32(1) {
		goto L80
	} else {
		goto L81
	}
L79:
	;
	return v327
L80:
	;
	if l2 != 0 {
		goto L88
	} else {
		goto L89
	}
L81:
	;
	v301 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v302 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v301)+1)))
	if v302 != 0 {
		goto L80
	} else {
		goto L82
	}
L82:
	;
	v303 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if base.Ui32(v293) < base.Ui32(v303) {
		goto L80
	} else {
		goto L83
	}
L83:
	;
	v307 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	F_appendBinaryStringInfo(m, v301+int32(4), v307, v292-v307)
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	return int32(0)
L85:
	;
	if l2 == int32(0) {
		v327 = int32(1)
		goto L79
	} else {
		goto L86
	}
L86:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v294)
	return int32(1)
L87:
	;
	v327 = int32(0)
	goto L79
L88:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v294)
	goto L87
L89:
	;
	goto L90
L90:
	;
	v321 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v321
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v292
	if v294 != 0 {
		v327 = int32(15)
		goto L79
	} else {
		goto L91
	}
L91:
	;
	goto L87
}
func F_json_manifest_array_end(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v26 int32
	_ = v26
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	switch v8 - int32(6) {
	case 0, 4:
		*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(2)
		m.G0 = v6 + int32(16)
		return int32(0)
	default:
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+20))
		*(*int32)(unsafe.Add(mBase, uint32(v6))) = int32(428496)
		m.T0[v19].(func(*base.Module, int32, int32, int32))(m, v18, int32(199958), v6)
		mBase = m.M
		v26 = m.ExcPending
		if v26 != 0 {
			return int32(0)
		} else {
			base.Wasm_trap_unreachable()
			for {
			}
		}
	}
}
func F_json_to_record(m *base.Module, l0 int32) int32 {
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v4 = int32(0)
	v6 = F_populate_record_worker(m, l0, int32(421660), int32(1), v4, v4)
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return v6
	}
}
func F_json_to_recordset(m *base.Module, l0 int32) int32 {
	var v8 int32
	_ = v8
	F_populate_recordset_worker(m, l0, int32(106405), int32(1), int32(0))
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return int32(0)
	}
}
func F_json_typeof(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	v3 = m.G0
	v5 = v3 - int32(80)
	m.G0 = v5
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = F_pg_detoast_datum_packed(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		F_makeJsonLexContext(m, v5+int32(12), v10, int32(0))
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			v19 = F_json_lex(m, v5+int32(12))
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int32(0)
			} else {
				if v19 != 0 {
					F_json_errsave_error(m, v19, v5+int32(12), int32(0))
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return int32(0)
					} else {
						v26 = *(*int32)(unsafe.Add(mBase, uint32(v5)+40))
						v28 = v26 - int32(1)
						if base.Ui32(int32(11)) <= base.Ui32(v28) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v51 = m.ExcPending
							if v51 != 0 {
								return int32(0)
							} else {
								v52 = *(*int32)(unsafe.Add(mBase, uint32(v5)+40))
								*(*int32)(unsafe.Add(mBase, uint32(v5))) = v52
								F_errmsg_internal(m, int32(484504), v5)
								mBase = m.M
								v56 = m.ExcPending
								if v56 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(497693), int32(1909), int32(339490))
									mBase = m.M
									v61 = m.ExcPending
									if v61 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						} else {
							if int32(base.Ui32(int32(1815))>>(uint(v28)%32))&int32(1) == int32(0) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v51 = m.ExcPending
								if v51 != 0 {
									return int32(0)
								} else {
									v52 = *(*int32)(unsafe.Add(mBase, uint32(v5)+40))
									*(*int32)(unsafe.Add(mBase, uint32(v5))) = v52
									F_errmsg_internal(m, int32(484504), v5)
									mBase = m.M
									v56 = m.ExcPending
									if v56 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(497693), int32(1909), int32(339490))
										mBase = m.M
										v61 = m.ExcPending
										if v61 != 0 {
											return int32(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							} else {
								v41 = *(*int32)(unsafe.Add(mBase, uint32(v28<<(uint(int32(2))%32))+uint32(_consts[1304])))
								v42 = F_cstring_to_text(m, v41)
								mBase = m.M
								v43 = m.ExcPending
								if v43 != 0 {
									return int32(0)
								} else {
									m.G0 = v5 + int32(80)
									return v42
								}
							}
						}
					}
				} else {
					v26 = *(*int32)(unsafe.Add(mBase, uint32(v5)+40))
					v28 = v26 - int32(1)
					if base.Ui32(int32(11)) <= base.Ui32(v28) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v51 = m.ExcPending
						if v51 != 0 {
							return int32(0)
						} else {
							v52 = *(*int32)(unsafe.Add(mBase, uint32(v5)+40))
							*(*int32)(unsafe.Add(mBase, uint32(v5))) = v52
							F_errmsg_internal(m, int32(484504), v5)
							mBase = m.M
							v56 = m.ExcPending
							if v56 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(497693), int32(1909), int32(339490))
								mBase = m.M
								v61 = m.ExcPending
								if v61 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						if int32(base.Ui32(int32(1815))>>(uint(v28)%32))&int32(1) == int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v51 = m.ExcPending
							if v51 != 0 {
								return int32(0)
							} else {
								v52 = *(*int32)(unsafe.Add(mBase, uint32(v5)+40))
								*(*int32)(unsafe.Add(mBase, uint32(v5))) = v52
								F_errmsg_internal(m, int32(484504), v5)
								mBase = m.M
								v56 = m.ExcPending
								if v56 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(497693), int32(1909), int32(339490))
									mBase = m.M
									v61 = m.ExcPending
									if v61 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						} else {
							v41 = *(*int32)(unsafe.Add(mBase, uint32(v28<<(uint(int32(2))%32))+uint32(_consts[1304])))
							v42 = F_cstring_to_text(m, v41)
							mBase = m.M
							v43 = m.ExcPending
							if v43 != 0 {
								return int32(0)
							} else {
								m.G0 = v5 + int32(80)
								return v42
							}
						}
					}
				}
			}
		}
	}
}
func F_json_unique_object_end(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	v2 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)))
	if v2 == int32(1) {
		v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v6
		F_pfree(m, v5)
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			return int32(0)
		}
	} else {
		return int32(0)
	}
}
func F_json_unique_object_field_start(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)))
	if v9 != int32(1) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v7 + int32(16)
	return int32(0)
L2:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = l1
	v15 = F_strlen(m, l1)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = v13
	*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v15
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v24 = F_hash_search(m, v18, v7+int32(4), int32(1), v7+int32(3))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return int32(0)
L4:
	;
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+3)))
	if v28 != int32(1) {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v31 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v31)
	goto L6
L6:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v37 == int32(0) {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v40
	F_pfree(m, v37)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L3
	} else {
		goto L9
	}
L9:
	;
	goto L6
}
func F_makeJsonTablePathSpec(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	v8 = F_palloc0(m, int32(20))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v8))) = int32(123)
		v15 = F_palloc0(m, int32(20))
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = l2
			*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = l0
			*(*int64)(unsafe.Add(mBase, uint32(v15))) = int64(2010044694600)
			*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v15
			if l1 != 0 {
				v22 = F_pstrdup(m, l1)
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v22
					*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = l2
					*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = l3
					return v8
				}
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = l2
				*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = l3
				return v8
			}
		}
	}
}
