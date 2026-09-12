package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_transformWindowFuncCall(m *base.Module, l0 int32, l1 int32, l2 int32) {
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
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v191 int32
	_ = v191
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
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
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v274 int32
	_ = v274
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v291 int32
	_ = v291
	var v295 int32
	_ = v295
	var v303 int32
	_ = v303
	var v311 int32
	_ = v311
	var v314 int32
	_ = v314
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v328 int32
	_ = v328
	v4 = int32(0)
	v11 = m.G0
	v13 = v11 - int32(48)
	m.G0 = v13
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+93)))
	if v15 == int32(1) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L5
	} else {
		goto L108
	}
L2:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v19 = F_contain_windowfuncs(m, v18)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	switch v22 - int32(2) {
	case 0, 1:
		v42 = int32(147802)
		goto L10
	case 2, 4, 5, 6, 13, 14, 15, 17, 20, 21, 22, 23, 24, 25, 42:
		goto L9
	case 3:
		goto L27
	case 7, 8, 9, 10, 11:
		goto L25
	default:
		goto L8
	case 16:
		goto L24
	case 26, 27:
		goto L23
	case 28, 29:
		goto L22
	case 30:
		goto L21
	case 31:
		goto L19
	case 32:
		goto L20
	case 33:
		goto L18
	case 34:
		goto L17
	case 35:
		goto L16
	case 36:
		goto L26
	case 37:
		goto L15
	case 38:
		goto L14
	case 39:
		goto L13
	case 40:
		goto L12
	case 41:
		goto L11
	}
L5:
	;
	return
L6:
	;
	if v19 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	goto L4
L8:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v97 != 0 {
		goto L43
	} else {
		goto L44
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L5
	} else {
		goto L33
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L5
	} else {
		goto L28
	}
L11:
	;
	v42 = int32(154308)
	goto L10
L12:
	;
	v42 = int32(148569)
	goto L10
L13:
	;
	v42 = int32(130263)
	goto L10
L14:
	;
	v42 = int32(152842)
	goto L10
L15:
	;
	v42 = int32(444348)
	goto L10
L16:
	;
	v42 = int32(148023)
	goto L10
L17:
	;
	v42 = int32(141329)
	goto L10
L18:
	;
	v42 = int32(154689)
	goto L10
L19:
	;
	v42 = int32(170348)
	goto L10
L20:
	;
	v42 = int32(153932)
	goto L10
L21:
	;
	v42 = int32(153378)
	goto L10
L22:
	;
	v42 = int32(155105)
	goto L10
L23:
	;
	v42 = int32(127347)
	goto L10
L24:
	;
	v42 = int32(148324)
	goto L10
L25:
	;
	v42 = int32(147373)
	goto L10
L26:
	;
	v42 = int32(153145)
	goto L10
L27:
	;
	v42 = int32(556013)
	goto L10
L28:
	;
	F_errcode(m, int32(655492))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L5
	} else {
		goto L29
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v42
	F_errmsg_internal(m, int32(216470), v13+int32(16))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L5
	} else {
		goto L30
	}
L30:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	F_parser_errposition(m, l0, v56)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L5
	} else {
		goto L31
	}
L31:
	;
	F_errfinish(m, int32(521728), int32(1039), int32(319115))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L5
	} else {
		goto L32
	}
L32:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L33:
	;
	F_errcode(m, int32(655492))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L5
	} else {
		goto L34
	}
L34:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	if base.Ui32(v71) <= base.Ui32(int32(44)) {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+32)) = v81
	F_errmsg(m, int32(195462), v13+int32(32))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L5
	} else {
		goto L39
	}
L36:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v71<<(uint(int32(2))%32))+uint32(_consts[299])))
	v81 = v80
	goto L38
L37:
	;
	v81 = int32(445222)
	goto L38
L38:
	;
	goto L35
L39:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	F_parser_errposition(m, l0, v88)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L5
	} else {
		goto L40
	}
L40:
	;
	F_errfinish(m, int32(521728), int32(1046), int32(319115))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L5
	} else {
		goto L41
	}
L41:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+32)) = v295
	v303 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+93)) = uint8(v303)
	m.G0 = v13 + int32(48)
	return
L43:
	;
	if v96 == int32(0) {
		goto L46
	} else {
		goto L47
	}
L44:
	;
	goto L45
L45:
	;
	if v96 != 0 {
		goto L72
	} else {
		goto L73
	}
L46:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L5
	} else {
		goto L67
	}
L47:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v96)+4))
	if v100 <= int32(0) {
		goto L46
	} else {
		goto L48
	}
L48:
	;
	v103 = int32(0)
	if v103 < v100 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v106 = v100
	goto L51
L50:
	;
	v106 = v103
	goto L51
L51:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v96)+12))
	v112 = v4
	goto L52
L52:
	;
	v119 = v112 + int32(1)
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v107+v112<<(uint(int32(2))%32))))
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v123)+4))
	if v124 != 0 {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	goto L46
L54:
	;
	v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97))))
	v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v124))))
	if v128 == int32(0) {
		v147 = v127
		v148 = v128
		goto L58
	} else {
		goto L59
	}
L55:
	;
	goto L56
L56:
	;
	if v106 != v119 {
		v112 = v119
		goto L52
	} else {
		goto L66
	}
L57:
	;
	if v148-v147 == int32(0) {
		v295 = v119
		goto L42
	} else {
		goto L65
	}
L58:
	;
	goto L57
L59:
	;
	if v127 != v128 {
		v147 = v127
		v148 = v128
		goto L58
	} else {
		goto L60
	}
L60:
	;
	v132 = v124
	v133 = v97
	goto L61
L61:
	;
	v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v133)+1)))
	v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v132)+1)))
	if v137 == int32(0) {
		v147 = v136
		v148 = v137
		goto L58
	} else {
		goto L63
	}
L62:
	;
	v147 = v136
	v148 = v137
	goto L58
L63:
	;
	v140 = int32(1)
	if v136 == v137 {
		v132 = v132 + v140
		v133 = v133 + v140
		goto L61
	} else {
		goto L64
	}
L64:
	;
	goto L62
L65:
	;
	goto L56
L66:
	;
	goto L53
L67:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L5
	} else {
		goto L68
	}
L68:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v170
	F_errmsg(m, int32(76472), v13)
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L5
	} else {
		goto L69
	}
L69:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(l2)+32))
	F_parser_errposition(m, l0, v175)
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L5
	} else {
		goto L70
	}
L70:
	;
	F_errfinish(m, int32(521728), int32(1079), int32(319115))
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L5
	} else {
		goto L71
	}
L71:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L72:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v96)+4))
	if int32(0) < v184 {
		goto L75
	} else {
		goto L76
	}
L73:
	;
	v285 = v4
	goto L74
L74:
	;
	v286 = F_lappend(m, v285, l2)
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L5
	} else {
		goto L106
	}
L75:
	;
	v191 = v4
	goto L78
L76:
	;
	goto L77
L77:
	;
	v274 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v285 = v274
	goto L74
L78:
	;
	v198 = v191 + int32(1)
	v199 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v96)+12))
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v200+v191<<(uint(int32(2))%32))))
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v204)+8))
	if v205 != 0 {
		goto L82
	} else {
		goto L83
	}
L79:
	;
	goto L77
L80:
	;
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v96)+4))
	if v198 < v262 {
		v191 = v198
		goto L78
	} else {
		goto L105
	}
L81:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v204)+12))
	v236 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v237 = F_equal(m, v235, v236)
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L5
	} else {
		goto L96
	}
L82:
	;
	if v199 == int32(0) {
		goto L80
	} else {
		goto L85
	}
L83:
	;
	goto L84
L84:
	;
	if v199 != 0 {
		goto L80
	} else {
		goto L95
	}
L85:
	;
	v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v199))))
	v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205))))
	if v211 == int32(0) {
		v230 = v210
		v231 = v211
		goto L87
	} else {
		goto L88
	}
L86:
	;
	if v231-v230 == int32(0) {
		goto L81
	} else {
		goto L94
	}
L87:
	;
	goto L86
L88:
	;
	if v210 != v211 {
		v230 = v210
		v231 = v211
		goto L87
	} else {
		goto L89
	}
L89:
	;
	v215 = v205
	v216 = v199
	goto L90
L90:
	;
	v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v216)+1)))
	v220 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v215)+1)))
	if v220 == int32(0) {
		v230 = v219
		v231 = v220
		goto L87
	} else {
		goto L92
	}
L91:
	;
	v230 = v219
	v231 = v220
	goto L87
L92:
	;
	v223 = int32(1)
	if v219 == v220 {
		v215 = v215 + v223
		v216 = v216 + v223
		goto L90
	} else {
		goto L93
	}
L93:
	;
	goto L91
L94:
	;
	goto L80
L95:
	;
	goto L81
L96:
	;
	if v237 == int32(0) {
		goto L80
	} else {
		goto L97
	}
L97:
	;
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v204)+16))
	v242 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v243 = F_equal(m, v241, v242)
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L5
	} else {
		goto L98
	}
L98:
	;
	if v243 == int32(0) {
		goto L80
	} else {
		goto L99
	}
L99:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v204)+20))
	v248 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	if v247 != v248 {
		goto L80
	} else {
		goto L100
	}
L100:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v204)+24))
	v251 = *(*int32)(unsafe.Add(mBase, uint32(l2)+24))
	v252 = F_equal(m, v250, v251)
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L5
	} else {
		goto L101
	}
L101:
	;
	if v252 == int32(0) {
		goto L80
	} else {
		goto L102
	}
L102:
	;
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v204)+28))
	v257 = *(*int32)(unsafe.Add(mBase, uint32(l2)+28))
	v258 = F_equal(m, v256, v257)
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L5
	} else {
		goto L103
	}
L103:
	;
	if v258 == int32(0) {
		goto L80
	} else {
		goto L104
	}
L104:
	;
	v295 = v198
	goto L42
L105:
	;
	goto L79
L106:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = v286
	if v286 == int32(0) {
		v295 = int32(0)
		goto L42
	} else {
		goto L107
	}
L107:
	;
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v286)+4))
	v295 = v291
	goto L42
L108:
	;
	F_errcode(m, int32(655492))
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L5
	} else {
		goto L109
	}
L109:
	;
	F_errmsg(m, int32(461507), int32(0))
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L5
	} else {
		goto L110
	}
L110:
	;
	v319 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v320 = F_locate_windowfunc(m, v319)
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L5
	} else {
		goto L111
	}
L111:
	;
	F_parser_errposition(m, l0, v320)
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L5
	} else {
		goto L112
	}
L112:
	;
	F_errfinish(m, int32(521728), int32(898), int32(319115))
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L5
	} else {
		goto L113
	}
L113:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
